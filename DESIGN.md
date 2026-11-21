# HN Digest RSS — 設計書

## 概要

Hacker News の記事を日次クロール→日本語翻訳→サマリー→RSS 配信するパイプラインを構築する。

```
[GitHub Actions] HN クロール
        ↓
contents/YYYY-MM-DD/*.md (翻訳済み Markdown)
        ↓
[GitHub Actions] サマリー (Google Translate — 無料)
        ↓
summaries/YYYY-MM-DD.json
        ↓
[Vercel] Astro ビルド → rss.xml 配信
```

---

## 現状

| 項目 | 現状 |
|------|------|
| クロール | Go ツール (`cmd/hn-digest`) が Algolia API で 24h 以内の記事を取得 |
| 翻訳 | Google Translate (非公式 API、無料) で日本語化 |
| 保存先 | `contents/YYYY-MM-DD/*.md` (1日あたり約 150〜200 記事) |
| フロントマター | `hn_id`, `score`, `title`, `hn_url`, `posted_at`, `image` など |
| 本文 | `## Translation` に日本語訳、`## Original Extract` に原文 |
| CI | GitHub Actions (`hn-digest.yml`) が毎日 UTC 0:00 に実行 |

---

## 設計方針

### 1. 記事の選定

1日 150〜200 記事を全部 RSS に出すと読めない。**1日 30 記事** に絞る。

選定基準:
- `final_score` 降順で上位 30 件
- `final_score = int(hn_score * (1.0 + sum of matched interest bonuses))`
- 興味カテゴリ別の bonus（タイトル / 翻訳本文に含まれるキーワードで判定）:
  - `ai`: 0.3
  - `sre`: 0.5
  - `platform`: 0.8
  - 各カテゴリは最大 1 回しか加算されない（複数キーワードがヒットしても重複加算なし）
- すでに日本語翻訳済みの `.md` を読むだけなので追加クロール不要

### 2. サマリー（Gemini 無料枠 + 抽出要約フォールバック）

当初は GitHub Models（`GITHUB_TOKEN` だけで使える無料の OpenAI 互換 API、`gpt-4o-mini`）で
サマリーとタイトル翻訳を生成していたが、**GitHub Models は 2026 年に廃止された**。

- 旧エンドポイント `https://models.inference.ai.azure.com` → `404`
- 新エンドポイント `https://models.github.ai/inference` → `410 github_models_retirement_brownout`

その後しばらくは LLM を使わず Google Translate だけの**抽出要約**で運用していたが、
本文の先頭を切り出すだけなのでナビゲーション文言などのノイズが混じりやすかった。
現在は **Gemini API の無料枠（AI Studio のキー）で抽象要約を生成し、失敗時は抽出要約に落ちる**
二段構えにしている。

| 項目 | 仕様 |
|------|------|
| LLM | Gemini API `generateContent`。モデルは `GEMINI_MODEL`（既定 `gemini-3.5-flash-lite`） |
| 認証 | リポジトリ secret `GEMINI_API_KEY`。**未設定でも動く**（抽出要約になるだけ） |
| プロンプト | 「3文以内・300文字以内・記事の事実のみ・前置きなし」を日本語で指示 |
| フォールバック | 失敗した記事だけ `leadSummary`（抽出要約）を使う |
| タイトル翻訳 | 従来どおり Google Translate。失敗時は英語タイトルのまま |

#### 無料枠を超えないための設計

Gemini 無料枠は RPM（分あたり）と RPD（日あたり）で制限される。`summarize.yml` は
`hn-digest.yml` の完了ごと、つまり**毎時**走るため、素直に実装すると
「30 記事 × 24 回 = 720 リクエスト/日」になり枠を圧迫する。そこで:

1. **フロントマターにキャッシュする**。生成した要約は記事の `.md` に `summary_ja` として
   書き戻し、`contents/` ごと commit する（画像 `image` のバックフィルと同じ仕組み）。
   次回以降そのキーがあれば API を呼ばない。結果、**1 記事あたり生涯 1 リクエスト**で、
   実消費は「その日の新規記事数」（数十件）程度に収まる。
2. **ペーシング**。`--llm-interval`（既定 4s）で呼び出し間隔を空け、RPM 制限に当てない。
3. **上限**。`--llm-max`（既定 30）で 1 実行あたりの呼び出し回数を打ち切る。`0` で LLM 無効。

#### ワークフローを失敗させないための設計

要約は「あれば良いもの」であり、**要約起因でワークフローが失敗してはならない**。

- `GEMINI_API_KEY` 未設定 → 警告ログのみ出して抽出要約で続行。
- `429`（レート制限）/ `5xx` → 指数バックオフで最大 2 回リトライ（5s → 10s）。
- `401`（キー不正）/ `404`（モデル廃止）→ リトライせず即フォールバック。
  時間をかけても直らないため。
- **3 回連続失敗したらその実行では LLM を諦める**。キー失効やモデル廃止は全記事で失敗するので、
  30 回無駄に叩いてジョブを引き延ばさない。
- 成功した要約だけを書き戻す。失敗を「確認済み」として記録しないので、次回の実行で再挑戦できる。

抽出要約のロジック（`leadSummary`、フォールバック時に使用）:

1. `## Translation` は `タイトル: / 記事タイトル: / 説明: / 記事本文:` の順に並ぶため、
   `記事本文:` 以降を優先して使う。本文が無い記事（取得失敗など）は `説明:` にフォールバックする。
2. 900 文字を超える場合は、直前の句点（`。！？`）で切って文が途中で終わらないようにする。
   句点が見つからない場合のみ `…` を付けて切る。

注意: Gemini 無料枠のデータは Google の製品改善に利用され得る。送信しているのは
公開済みの Hacker News コンテンツのみなので許容している。

### 3. summaries JSON スキーマ

`summaries/YYYY-MM-DD.json` に保存:

```json
{
  "date": "2026-05-30",
  "generated_at": "2026-05-30T01:00:00Z",
  "articles": [
    {
      "rank": 1,
      "hn_id": 48334157,
      "title": "Show HN: AI-org – org-mode powered by AI",
      "title_ja": "Show HN: AI-org – AI で動く org-mode",
      "hn_url": "https://news.ycombinator.com/item?id=48334157",
      "source_url": "https://ai-org.net/",
      "image_url": "https://ai-org.net/og.png",
      "score": 312,
      "comments": 87,
      "posted_at": "2026-05-30T08:59:47Z",
      "summary_ja": "AI を活用した org-mode タスクマネージャー。..."
    }
  ]
}
```

### 4. 記事画像（og:image）

リンク先サイトの HTML から `og:image`（無ければ `twitter:image`）を取り出し、RSS のアイテム画像として配信する。
追加のリクエストは発生しない（本文取得のために既に HTML を GET しているので、そのレスポンスから抽出するだけ）。

| 項目 | 仕様 |
|------|------|
| 優先順 | `og:image:secure_url` → `og:image` → `twitter:image` |
| 相対 URL | リダイレクト後の最終 URL（`resp.Request.URL`）を基準に絶対 URL 化 |
| 除外 | `http` / `https` 以外のスキーム（`data:` など）とホスト無しの URL |
| 保存先 | フロントマターの `image` → `summaries/*.json` の `image_url` |

`firstMeta` は値を `htmlToText` に通して 300 文字で切るため URL には使えない（CDN の署名付き URL が壊れる）。
画像 URL 専用に `firstImageURL` を用意し、HTML エンティティのアンエスケープのみ行う。

取得できない記事もある（Reuters / Bloomberg など本文取得自体が 401 / 403 になるサイト）。
その場合 `image_url` は省略され、RSS は従来通りテキストのみのアイテムになる。

#### バックフィル

クローラーは `contents/` に既にある記事をスキップするため、画像抽出の実装前にクロールした記事は
そのままでは永久に画像が付かない。そこで `cmd/summarize` 側で不足分を補う:

| 状態 | 判定 | 動作 |
|------|------|------|
| `image` キーが無い | 未確認（画像抽出前にクロール） | リンク先を取得して抽出し、フロントマターに書き戻す |
| `image: ""` | 確認済み・画像なし | 何もしない（再取得しない） |
| `image: "https://…"` | 確認済み・画像あり | そのまま使う |
| 取得が失敗（timeout / 401 / 403） | 一時的な失敗と見なす | キーを書かない → 次回以降のランで再試行 |

フェッチ対象は上位 30 件のうち未確認のものだけで、結果はフロントマターに永続化するため
**1 記事あたり 1 リクエスト**しか発生しない。`summarize.yml` は `contents/` も commit する
（しないとバックフィルが破棄され、毎ランで再取得してしまう）。

過去日付は `workflow_dispatch` の `date` 入力で個別にバックフィルできる。

RSS では 3 通りの方法で同じ画像を渡し、リーダー側の対応差を吸収する:

1. `<media:content>` / `<media:thumbnail>`（Media RSS）— Inoreader・Feedly のサムネイル用
2. `<content:encoded>` の `<img>` — HTML を描画するリーダーでの本文内表示用
3. `<description>` は従来通りプレーンテキストのまま — 上記に未対応のリーダーでも壊れない

### 5. RSS フィード設計

| 項目 | 仕様 |
|------|------|
| フォーマット | RSS 2.0 |
| 配信 URL | `https://<your-site>.vercel.app/rss.xml` |
| アイテム数 | 直近 14 日分 × 30 記事 = 最大 420 件 |
| `<title>` | 記事タイトル（日本語訳 `title_ja`、無ければ英語タイトル） |
| `<description>` | サマリー（日本語、抽出要約） |
| `<link>` | HN のディスカッション URL |
| `<pubDate>` | `posted_at` |
| `<category>` | `Hacker News` |
| 画像 | `<media:content>` / `<media:thumbnail>`（Media RSS）と `<content:encoded>` の `<img>` |

RSS アイテムは記事単位（1日まとめではなく個別）にする方が RSS リーダーで読みやすい。

### 6. Astro サイト構成

```
site/
├── src/
│   ├── pages/
│   │   ├── index.astro          # 最新ダイジェスト一覧
│   │   └── rss.xml.ts           # RSS フィード生成
│   └── lib/
│       └── summaries.ts         # JSON 読み込みユーティリティ
├── astro.config.mjs
├── package.json
└── tsconfig.json
```

Astro のビルド時に `summaries/*.json` を読み込んで静的 RSS を生成する。

**`astro.config.mjs` のポイント:**
- `output: 'static'`（SSG）
- `@astrojs/rss` パッケージで RSS 生成

### 7. Vercel デプロイ設定

Vercel のプロジェクト設定:

| 設定 | 値 |
|------|----|
| Root Directory | `site` |
| Build Command | `npm run build` |
| Output Directory | `dist` |
| Install Command | `npm install` |
| 自動デプロイ | `main` ブランチへの push で起動 |

---

## パイプライン全体の流れ

```
毎日 UTC 00:00
  └─ [hn-digest.yml] Go クローラー実行
       └─ contents/2026-05-31/*.md を commit & push
            └─ [summarize.yml] hn-digest.yml 完了をトリガーに起動
                 ├─ contents/2026-05-31/*.md を interest bonus 適用後の final_score でソート
                 ├─ 上位 30 件を選定
                 ├─ Google Translate でタイトルを日本語化、本文冒頭から抽出要約
                 ├─ summaries/2026-05-31.json を生成
                 └─ commit & push → Vercel が自動ビルド → rss.xml 更新
```

---

## GitHub Actions ワークフロー設計

### `summarize.yml`

```yaml
name: Summarize HN digest

on:
  workflow_run:
    workflows: ["HN digest"]
    types: [completed]
  workflow_dispatch:

permissions:
  contents: write

jobs:
  summarize:
    runs-on: ubuntu-latest
    if: ${{ github.event_name == 'workflow_dispatch' || github.event.workflow_run.conclusion == 'success' }}
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
          cache: true
      - name: Test
        run: go test ./...
      - name: Build
        run: go build -o bin/summarize ./cmd/summarize
      - name: Generate summaries
        env:
          GEMINI_API_KEY: ${{ secrets.GEMINI_API_KEY }}   # 任意。無くても動く
          GEMINI_MODEL: ${{ vars.GEMINI_MODEL }}
        run: ./bin/summarize
      - name: Commit summaries
        run: |
          git config user.name "github-actions[bot]"
          git config user.email "41898282+github-actions[bot]@users.noreply.github.com"
          git add summaries/
          git diff --staged --quiet || git commit -m "chore: update summaries"
          git push origin HEAD:main
```

### `cmd/summarize`（Go）

処理フロー:
1. 今日の `contents/YYYY-MM-DD/` フォルダを読む
2. YAML フロントマターから `score` と `title` を取得
3. `title + ## Translation` 本文に対して interest bonus を計算し、`final_score` 降順でソート
4. 上位 30 件の `## Translation` セクションを先頭 4000 文字まで抽出
5. Google Translate でタイトルを日本語化（`title_ja`）、フロントマターの `image` を `image_url` に引き継ぐ
6. `summary_ja` を決める:
   フロントマターにキャッシュ済みならそれを使う →
   無ければ Gemini（無料枠）で生成してフロントマターに書き戻す →
   失敗したら `## Translation` 本文の冒頭 900 文字を句点で切った抽出要約
7. `summaries/YYYY-MM-DD.json` に保存

---

## ディレクトリ構成（完成形）

```
hn-digest/
├── .github/
│   └── workflows/
│       ├── hn-digest.yml        # 既存: クロール
│       └── summarize.yml        # サマリー生成
├── cmd/
│   ├── hn-digest/               # 既存: Go クローラー
│   └── summarize/               # サマリー生成 (Go)
├── internal/
│   ├── frontmatter/             # フロントマター読み取り (両 cmd 共有)
│   ├── gtranslate/              # Google Translate クライアント (両 cmd 共有)
│   └── llm/                     # Gemini 無料枠クライアント (要約、任意)
├── contents/                    # 既存: 翻訳済み Markdown
│   └── YYYY-MM-DD/
│       └── *.md
├── summaries/                   # サマリー JSON
│   └── YYYY-MM-DD.json
└── site/                        # 新規: Astro サイト
    ├── src/
    │   ├── pages/
    │   │   ├── index.astro
    │   │   └── rss.xml.ts
    │   └── lib/
    │       └── summaries.ts
    ├── astro.config.mjs
    ├── package.json
    └── tsconfig.json
```

---

## 技術選定の理由

| 技術 | 理由 |
|------|------|
| Gemini API (無料枠) | 無料で抽象要約が得られる。キー未設定でも動く任意機能として組み込んだ |
| Google Translate (非公式 API) | 認証不要・無料。タイトルと本文の日本語化に使用 |
| Astro | 静的サイト生成が得意。`@astrojs/rss` で RSS 生成が簡単 |
| Vercel | Astro との相性が良い。git push で自動デプロイ。無料プランで十分 |
| Go (`cmd/summarize`) | クローラーと同じ言語で統一。標準ライブラリだけで完結する |
| RSS 2.0 | 幅広い RSS リーダーで対応 |

---

## コスト

| サービス | 料金 |
|----------|------|
| GitHub Actions | 無料（パブリックリポジトリ） |
| Google Translate (非公式 API) | 無料（認証不要） |
| Gemini API | 無料枠（AI Studio のキー。フロントマターにキャッシュするため実消費は数十リクエスト/日） |
| Vercel | 無料（Hobby プラン） |
| 合計 | **$0** |

---

## 今後の拡張候補（スコープ外）

- Gemini 以外の無料枠（Cloudflare Workers AI、Groq 等）へのフェイルオーバー
- カテゴリ別 RSS（AI、DevOps、SRE など）
- Slack / LINE 通知
- スコアが極めて高い記事（> 500 点）の即時通知
- 週次・月次のベスト記事まとめ
- Web UI でのアーカイブ検索
