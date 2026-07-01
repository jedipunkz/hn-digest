---
source: "https://ht-ml.app/"
hn_url: "https://news.ycombinator.com/item?id=48749414"
title: "HT-ML.app – Deploy HTML Artifacts from Claude Code and Codex"
article_title: "ht-ml.app — HTML hosting with an API built for agents"
author: "nsmith22"
captured_at: "2026-07-01T16:46:29Z"
capture_tool: "hn-digest"
hn_id: 48749414
score: 1
comments: 1
posted_at: "2026-07-01T16:27:14Z"
tags:
  - hacker-news
  - translated
---

# HT-ML.app – Deploy HTML Artifacts from Claude Code and Codex

- HN: [48749414](https://news.ycombinator.com/item?id=48749414)
- Source: [ht-ml.app](https://ht-ml.app/)
- Score: 1
- Comments: 1
- Posted: 2026-07-01T16:27:14Z

## Translation

タイトル: HT-ML.app – クロード コードおよびコーデックスからの HTML アーティファクトのデプロイ
記事のタイトル: ht-ml.app — エージェント向けに構築された API を使用した HTML ホスティング
説明: ht-ml.app は、AI エージェント用に設計された API を備えた HTML ホスティングです。単一の HTML ドキュメントを投稿すると、アカウントもサインアップも必要なく、すぐにパブリック URL を取得できます。各サイトは、作成時に返される update_key によって保護されます。

記事本文:
コンテンツにスキップ
HTMLホスティング・エージェント用API
1 分以内に Claude Code または Codex から直接ライブに移行します。
まず、Claude Code または Codex に https://ht-ml.app を使用して HTML ファイルをデプロイするように依頼します。プロトタイプ、豊富なアーキテクチャ図、コード レビュー、デッキ、イラストなどをエージェント環境から直接展開できるように最適化されています。単一の HTML ドキュメントを POST すると、即座にパブリック URL が返されます。アカウント、ダッシュボード、プロビジョニングする API キーは必要ありません。
エージェントのスキルをインストールする
ターンキーセットアップをご希望ですか?オープン HTML エージェント スキルは、エージェントにワークフロー全体 (作成、更新、アセット、パスワード保護) を教え、20 個の既製のページ テンプレートをバンドルします。 1 つのコマンドで Claude Code と Codex 用にインストールされます。
カール -fsSL https://raw.githubusercontent.com/nsmith/html/main/install.sh |しー
代替スキル CLI - 検出されたすべてのエージェントにインストールします
コピー
npx スキルで nsmith/html を追加
クロード・コード
コーデックス
オープンクロー
エルメス
+ 任意のエージェント スキル クライアント
オープン形式 (agentkills.io) · GitHub 上のソースおよびクライアントごとのインストール パス。
HTML を API に送信します。応答により、公開 URL と、後の編集に使用する秘密キーが渡されます。
カール -X POST https://api.ht-ml.app/v1/sites \
-H "コンテンツ タイプ: application/json" \
-d '{"html_content": "<h1>エージェントからこんにちは</h1>"}'
応答 · 200 — {中括弧} 内の値はプレースホルダーです
{
"site_id" : "{site_id}" ,
"url" : "https://{site_id}.ht-ml.app/" ,
"update_key" : "{update_key}" ,
「ステータス」：「アクティブ」
}
応答の URL フィールドからライブ URL を読み取ります。各サイトは独自のサブドメイン https://{site_id}.ht-ml.app/ を取得するため、相対アセット パス ( <img src="images/logo.png"> ) が機能します。 update_key は秘密にしておきます。これは 1 回だけ返され、書き込みアクセスが許可されます。
html_content を /v1/sites に POST します。 HT

ML は安全でないコンテンツをスキャンして保存されます。 site_id と update_key が返されます。
/v1/sites/{site_id} を取得して、HTML が参照しているアセット (画像、ビデオ) のうち、まだ見つからないアセットを列挙します。
欠落しているアセットごとに、update_key をベアラー トークンとして使用して、ファイルを /v1/sites/{site_id}/assets に POST します。パスは HTML 内の src を反映します。
ベース URL: https://api.ht-ml.app/v1 。機械可読な概要は /v1/help にあります。
POST /v1/sites/{site_id}/assets
エラーは単に失敗するだけではなく、エージェントに回復方法を指示します。すべてのエラー応答には、対処可能なメッセージが含まれています。
機械で消耗するように設計されています
ブラウザ内エージェントの場合、このページは、 navigator.modelContext / document.modelContext を介して、publish_html_site という名前の WebMCP ツールを登録します。 html_content を指定して呼び出すと、サイトが公開されて URL が返されます。手動の HTTP は必要ありません。
/llms.txt — 1 回のフェッチで読み取ることができる、LLM に最適化された簡潔なリファレンス。
/v1/help — API はそれ自体を JSON で説明します。
このページには、クローラーと回答エンジン用の構造化データ ( schema.org/WebAPI + HowTo ) が埋め込まれています。
AI クローラーは /robots.txt で明示的に歓迎されます。
エージェント向けに構築された API を使用した HTML ホスティング。

## Original Extract

ht-ml.app is HTML hosting with an API designed for AI agents. POST a single HTML document and get a public URL instantly — no accounts, no signup. Each site is secured by an update_key returned at creation.

Skip to content
HTML hosting · API for agents
Go live straight from Claude Code or Codex in under a minute .
To get started, ask Claude Code or Codex to deploy your HTML file with https://ht-ml.app . Optimized for deploying prototypes, rich architecture diagrams, code reviews, decks, illustrations, and more straight from agentic environments. POST a single HTML document and get a public URL back instantly. No accounts, no dashboards, no API keys to provision.
Install the skill for your agent
Prefer a turnkey setup? The open html Agent Skill teaches your agent the whole workflow — create, update, assets, password protection — and bundles 20 ready-made page templates . One command installs it for Claude Code and Codex.
curl -fsSL https://raw.githubusercontent.com/nsmith/html/main/install.sh | sh
Alternative · skills CLI — installs to all detected agents
Copy
npx skills add nsmith/html
Claude Code
Codex
OpenClaw
Hermes
+ any Agent Skills client
Open format ( agentskills.io ) · source & per-client install paths on GitHub .
Send your HTML to the API. The response hands you the public URL and the secret key you'll use for any later edits.
curl -X POST https://api.ht-ml.app/v1/sites \
-H "Content-Type: application/json" \
-d '{"html_content": "<h1>Hello from my agent</h1>"}'
Response · 200 — values in {braces} are placeholders
{
"site_id" : "{site_id}" ,
"url" : "https://{site_id}.ht-ml.app/" ,
"update_key" : "{update_key}" ,
"status" : "active"
}
Read the live URL from the url field of your response — each site gets its own subdomain, https://{site_id}.ht-ml.app/ , so relative asset paths just work ( <img src="images/logo.png"> ). Keep the update_key secret — it's returned only once and grants write access.
POST html_content to /v1/sites . The HTML is scanned for unsafe content, then stored. You get back a site_id and an update_key .
GET /v1/sites/{site_id} to enumerate the assets your HTML references (images, video) and which are still missing .
For each missing asset, POST the file to /v1/sites/{site_id}/assets with your update_key as a Bearer token. Paths mirror the src in your HTML.
Base URL: https://api.ht-ml.app/v1 . A machine-readable summary lives at /v1/help .
POST /v1/sites/{site_id}/assets
Errors don't just fail — they tell an agent how to recover. Every error response includes an actionable message .
Designed to be machine-consumable
If you're an in-browser agent, this page registers a WebMCP tool named publish_html_site via navigator.modelContext / document.modelContext . Call it with html_content and it publishes the site and returns the URL — no manual HTTP required.
/llms.txt — a concise, LLM-optimized reference you can read in one fetch.
/v1/help — the API describes itself in JSON.
Structured data ( schema.org/WebAPI + HowTo ) is embedded in this page for crawlers and answer engines.
AI crawlers are explicitly welcomed in /robots.txt .
HTML hosting with an API built for agents.
