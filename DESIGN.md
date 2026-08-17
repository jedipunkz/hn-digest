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

### 2. サマリー（Google Translate — 無料）

当初は GitHub Models（`GITHUB_TOKEN` だけで使える無料の OpenAI 互換 API、`gpt-4o-mini`）で
サマリーとタイトル翻訳を生成していたが、**GitHub Models は 2026 年に廃止された**。

- 旧エンドポイント `https://models.inference.ai.azure.com` → `404`
- 新エンドポイント `https://models.github.ai/inference` → `410 github_models_retirement_brownout`

無料枠で継続するため、姉妹リポジトリ [hw-digest](https://github.com/jedipunkz/hw-digest) と同じ方式に切り替えた。
すなわち **LLM を使わず、Google Translate（非公式 API、認証不要・無料）だけで完結させる**。

| 項目 | 仕様 |
|------|------|
| タイトル翻訳 | Google Translate で `title` を日本語化し `title_ja` に保存 |
| サマリー | 抽象要約ではなく**抽出要約**。`## Translation` の本文冒頭を 900 文字で切る |
| 認証 | 不要（`GITHUB_TOKEN` も `models: read` 権限も使わない） |
| レート制限 | 1日 30 リクエスト（タイトル翻訳のみ）。実測で問題なし |

抽出要約のロジック（`leadSummary`）:

1. `## Translation` は `タイトル: / 記事タイトル: / 説明: / 記事本文:` の順に並ぶため、
   `記事本文:` 以降を優先して使う。本文が無い記事（取得失敗など）は `説明:` にフォールバックする。
2. 900 文字を超える場合は、直前の句点（`。！？`）で切って文が途中で終わらないようにする。
   句点が見つからない場合のみ `…` を付けて切る。

トレードオフ: LLM 要約に比べて要約品質は落ちる（本文の抜粋なので、ページ内のナビゲーション文言などの
ノイズが混じることがある）。無料枠を維持する制約下での意図的な妥協で、hw-digest と同じ水準。
将来的に無料の LLM API を使う場合は `leadSummary` を差し替えるだけで済む。

タイトル翻訳が失敗しても記事は落とさず、英語タイトルにフォールバックして警告ログのみ出す。
サマリー生成には外部 API を使わないため、**ワークフローが要約起因で失敗することはない**。

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
| `<description>` | サマリー（日本語、抽出要約） |
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
5. Google Translate でタイトルを日本語化（`title_ja`）
6. `## Translation` 本文の冒頭 900 文字を句点で切って `summary_ja` を生成
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
│   └── gtranslate/              # Google Translate クライアント (両 cmd 共有)
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
| Google Translate (非公式 API) | 認証不要・無料。GitHub Models 廃止後に残った唯一の無料手段 |
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
| Vercel | 無料（Hobby プラン） |
| 合計 | **$0** |

---

## 今後の拡張候補（スコープ外）

- 無料枠の LLM（Gemini API 等）による抽象要約への差し替え（`leadSummary` を置き換える）
- カテゴリ別 RSS（AI、DevOps、SRE など）
- Slack / LINE 通知
- スコアが極めて高い記事（> 500 点）の即時通知
- 週次・月次のベスト記事まとめ
- Web UI でのアーカイブ検索
