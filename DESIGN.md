# HN Digest RSS — 設計書

## 概要

Hacker News の記事を日次クロール→日本語翻訳→AI サマリー→RSS 配信するパイプラインを構築する。

```
[GitHub Actions] HN クロール
        ↓
contents/YYYY-MM-DD/*.md (翻訳済み Markdown)
        ↓
[GitHub Actions] AI サマリー (GitHub Models — 無料)
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

### 2. AI サマリー（GitHub Models — 無料）

GitHub Models は `GITHUB_TOKEN` だけで使える OpenAI 互換 API。**料金なし**（レート制限あり）。

| 項目 | 仕様 |
|------|------|
| エンドポイント | `https://models.inference.ai.azure.com` |
| モデル | `gpt-4o-mini`（高速・無料・要約に十分） |
| レート制限 | 15 req/min、150 req/day（個人アカウント） |
| 1日の使用量 | 30 記事 = 30 リクエスト → 余裕あり |
| 認証 | GitHub Actions 組み込みの `${{ secrets.GITHUB_TOKEN }}` |

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
                 ├─ GitHub Models API で各記事をサマリー
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
  models: read   # GitHub Models へのアクセス

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
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
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
5. GitHub Models API (`gpt-4o-mini`) でサマリー生成（`max_tokens=1200`）
6. `summaries/YYYY-MM-DD.json` に保存

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
| GitHub Models (`gpt-4o-mini`) | `GITHUB_TOKEN` だけで無料利用可。サマリー程度なら品質十分 |
| Astro | 静的サイト生成が得意。`@astrojs/rss` で RSS 生成が簡単 |
| Vercel | Astro との相性が良い。git push で自動デプロイ。無料プランで十分 |
| Go (`cmd/summarize`) | クローラーと同じ言語で統一。標準ライブラリだけで GitHub Models API を叩ける |
| RSS 2.0 | 幅広い RSS リーダーで対応 |

---

## コスト

| サービス | 料金 |
|----------|------|
| GitHub Actions | 無料（パブリックリポジトリ） |
| GitHub Models API | 無料（`GITHUB_TOKEN` で利用、レート制限内） |
| Vercel | 無料（Hobby プラン） |
| 合計 | **$0** |

---

## 今後の拡張候補（スコープ外）

- カテゴリ別 RSS（AI、DevOps、SRE など）
- Slack / LINE 通知
- スコアが極めて高い記事（> 500 点）の即時通知
- 週次・月次のベスト記事まとめ
- Web UI でのアーカイブ検索
