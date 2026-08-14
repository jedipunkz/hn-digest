# HN Digest RSS — 設計書

## 概要

Hacker News の記事を日次クロール→日本語翻訳→AI サマリー→RSS 配信するパイプラインを構築する。

```
[GitHub Actions] HN クロール
        ↓
contents/YYYY-MM-DD/*.md (翻訳済み Markdown)
        ↓
[GitHub Actions] AI サマリー (OpenAI 互換 API — 無料枠)
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
| フロントマター | `hn_id`, `score`, `title`, `hn_url`, `posted_at` など |
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

### 2. AI サマリー（OpenAI 互換 API — 無料枠）

> **2026-07-30**: GitHub Models は完全廃止された（`models.inference.ai.azure.com` は本文空の 404 を返す）。
> サマリー生成は任意の **OpenAI 互換** プロバイダーへ移行し、環境変数で切り替え可能にした。

| 項目 | 仕様 |
|------|------|
| エンドポイント | `SUMMARIZE_API_BASE`（既定: `https://generativelanguage.googleapis.com/v1beta/openai`） |
| モデル | `SUMMARIZE_MODEL`（既定: `gemini-3.5-flash-lite`） |
| 認証 | `SUMMARIZE_API_KEY`（リポジトリ Secret） |
| 呼び出し間隔 | `-min-interval`（既定 4 秒 ≒ 15 req/min）で無料枠の RPM 制限に収める |
| 1日の使用量 | 初回実行で 30 記事 × 2 = 60 リクエスト、以降は新規記事分のみ |

`SUMMARIZE_API_BASE` / `SUMMARIZE_MODEL` を差し替えれば OpenAI・Groq・OpenRouter など
他の OpenAI 互換プロバイダーにもそのまま切り替えられる。

**既存サマリーの再利用（無料枠を守る仕組み）**

このジョブは毎時同じ日付に対して実行されるため、`summaries/YYYY-MM-DD.json` を読み込み、
`hn_id` が一致する記事の `title_ja` / `summary_ja` を再利用して API 呼び出しを省略する。
順位・スコア・コメント数などのメタデータは毎回更新する。
再利用しない場合は毎時 60 リクエスト（= 1日 1,440 リクエスト）が発生し、
どの無料枠でも超過するため、この再利用が無料運用の前提となる。
`-refresh` を付けると再利用せず全件を生成し直す。

**失敗時の挙動**

- 記事単位で失敗した場合はログに残してスキップし、成功分だけを保存する
- 全記事が失敗した場合は既存の JSON を上書きせずエラー終了する（内容の消失を防ぐ）
- HTTP ステータスを JSON パースより先に判定するため、本文が空のエラー応答でも
  `unexpected end of JSON input` ではなく実際の HTTP エラーが表示される

**サマリープロンプト（例）:**

```
以下のHacker News記事を日本語で詳しく要約してください。
以下の点を必ず含めて、600〜900文字程度（6〜10文）で記述してください。
- 記事の主題と背景
- 技術的な要点や仕組み、利用されている技術スタック
- 著者の主張や結論
- HNコミュニティで注目されている理由や論点

タイトル: {title}
本文:
{translation_text}
```

出力上限は `max_tokens = 1200`、入力本文は `## Translation` セクションを先頭 4000 文字まで切り詰めて渡す。

サマリーとは別に、記事タイトルも同じモデルで日本語へ翻訳し `title_ja` として保存する（製品名・固有名詞・技術用語は原語のまま残す）。翻訳に失敗または空文字の場合は英語タイトルにフォールバックする。

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
      "score": 312,
      "comments": 87,
      "posted_at": "2026-05-30T08:59:47Z",
      "summary_ja": "AI を活用した org-mode タスクマネージャー。..."
    }
  ]
}
```

### 4. RSS フィード設計

| 項目 | 仕様 |
|------|------|
| フォーマット | RSS 2.0 |
| 配信 URL | `https://<your-site>.vercel.app/rss.xml` |
| アイテム数 | 直近 14 日分 × 30 記事 = 最大 420 件 |
| `<title>` | 記事タイトル（日本語訳 `title_ja`、無ければ英語タイトル） |
| `<description>` | AI サマリー（日本語） |
| `<link>` | HN のディスカッション URL |
| `<pubDate>` | `posted_at` |
| `<category>` | `Hacker News` |

RSS アイテムは記事単位（1日まとめではなく個別）にする方が RSS リーダーで読みやすい。

### 5. Astro サイト構成

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

### 6. Vercel デプロイ設定

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
                 ├─ OpenAI 互換 API で各記事をサマリー（既存分は再利用）
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
      - name: Build
        run: go build -o bin/summarize ./cmd/summarize
      - name: Generate summaries
        env:
          SUMMARIZE_API_KEY: ${{ secrets.SUMMARIZE_API_KEY }}
          SUMMARIZE_API_BASE: ${{ vars.SUMMARIZE_API_BASE }}
          SUMMARIZE_MODEL: ${{ vars.SUMMARIZE_MODEL }}
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
5. 既存の `summaries/YYYY-MM-DD.json` を読み、`hn_id` が一致する記事は生成済みの
   `title_ja` / `summary_ja` を再利用（API 呼び出しをスキップ）
6. 未生成の記事のみ OpenAI 互換 API でサマリー生成（`max_tokens=1200`）
7. `summaries/YYYY-MM-DD.json` に保存

---

## ディレクトリ構成（完成形）

```
hn-digest/
├── .github/
│   └── workflows/
│       ├── hn-digest.yml        # 既存: クロール
│       └── summarize.yml        # 新規: AI サマリー
├── cmd/
│   ├── hn-digest/               # 既存: Go クローラー
│   └── summarize/               # AI サマリー生成 (Go)
├── contents/                    # 既存: 翻訳済み Markdown
│   └── YYYY-MM-DD/
│       └── *.md
├── summaries/                   # AI サマリー JSON
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
| OpenAI 互換 API（既定: Gemini `gemini-3.5-flash-lite`） | 無料枠があり、要約用途には十分な品質。環境変数だけで他プロバイダーへ切替可能 |
| Astro | 静的サイト生成が得意。`@astrojs/rss` で RSS 生成が簡単 |
| Vercel | Astro との相性が良い。git push で自動デプロイ。無料プランで十分 |
| Go (`cmd/summarize`) | クローラーと同じ言語で統一。標準ライブラリだけで OpenAI 互換 API を叩ける |
| RSS 2.0 | 幅広い RSS リーダーで対応 |

---

## コスト

| サービス | 料金 |
|----------|------|
| GitHub Actions | 無料（パブリックリポジトリ） |
| サマリー生成 API | 無料枠内（既存サマリーを再利用し、新規記事分のみ呼び出す） |
| Vercel | 無料（Hobby プラン） |
| 合計 | **$0** |

---

## 今後の拡張候補（スコープ外）

- カテゴリ別 RSS（AI、DevOps、SRE など）
- Slack / LINE 通知
- スコアが極めて高い記事（> 500 点）の即時通知
- 週次・月次のベスト記事まとめ
- Web UI でのアーカイブ検索
