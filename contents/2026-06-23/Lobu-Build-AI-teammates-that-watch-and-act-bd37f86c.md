---
source: "https://lobu.ai/"
hn_url: "https://news.ycombinator.com/item?id=48651524"
title: "Lobu: Build AI teammates that watch and act"
article_title: "Lobu - AI teammates that watch and act"
author: "handfuloflight"
captured_at: "2026-06-23T21:32:02Z"
capture_tool: "hn-digest"
hn_id: 48651524
score: 1
comments: 0
posted_at: "2026-06-23T21:11:28Z"
tags:
  - hacker-news
  - translated
---

# Lobu: Build AI teammates that watch and act

- HN: [48651524](https://news.ycombinator.com/item?id=48651524)
- Source: [lobu.ai](https://lobu.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T21:11:28Z

## Translation

タイトル: Lobu: 監視して行動する AI チームメイトを構築する
記事のタイトル: Lobu - 監視し行動する AI チームメイト
Description: Open-source infrastructure for multi-tenant AI agents: sandboxed execution, shared memory, and an SDK that pulls live company data and acts on your high-level goals.

記事本文:
Lobu Docs リソース サインイン 監視して行動する AI チームメイトを構築する
マルチテナント AI エージェント用のオープンソース インフラストラクチャ: サンドボックス実行、共有メモリ、および企業のライブ データを取得し、高レベルの目標に基づいて動作する SDK。
ChatGPT カーソル OpenCode セットアップ プロンプトのコピー 創設者と話す × コーディング エージェントに貼り付けてプロジェクトを足場にします。
または、自分で開始します: $ npx @lobu/cli@latest init my-agent
データを接続します。目標を設定します。ロブは変化を観察し、次のステップに備える。
読み取り可能なシステムを選択してください。 Lobu はそれらの更新を顧客の生きた記憶に変えます。
Stripe Zendesk Gmail Google Drive GitHub Notion Snowflake PostgreSQL 50 以上 既存のコネクタを使用するか、エージェントにコードを記述してデータを接続させます。すべてのソースの変更を監視する コネクタ SDK ↗ 2 目標を定義する
行動する前に、何に注意すべきか、いつ質問すべきかを伝えます。
「すべてのアカウントに解約リスクがないか注意してください。更新が 30 日以内で健全性が低下した場合は、承認を得るために CSM チェックインの草案を作成してください。」
スケジュールどおりにメモリをスキャンし、危険にさらされているアカウントを特定し、証拠を添付し続けます。
下書きを編集したり、送信したり、そのまま残したりすることができます。
R Revenue Agent online 注意: Acme Corp は解約傾向にあります。ログイン数は 14 日間で 38% 減少し、更新は 21 日後になります。
CSM へのチェックインの下書きを依頼したいですか? 12:01 「はい」というメールの下書きを作成し、使用量のドロップを含めます。 12:01 ドラフト。それを Acme アカウントに保存し、@dana に ping を送信しました。 12:01 API を介して、チームのチャットまたは独自のアプリで下書きメールの返信を表示します: API ↗ 例 エージェントのワークフローを調べます。
各例は、1 人の AI チームメイトのソース、メモリ、アクションを示しています。
ローカル、セルフホスト、または管理。
1 つのコマンドでゲートウェイ、ワーカー、メモリ、エンベディングを起動します。
データとコントロールを手元に置いておく必要がある場合は、Docker または Helm を使用してデプロイします。
使用

同じプロジェクトに管理された分離、シークレット、およびアップグレードを適用します。
最初のものを構築する
マルチユーザーエージェント。
電話予約 × ブログから 最新のブログ投稿
エージェント ループは新しい SaaS
かつてソフトウェアとは、機能ごとに垂直ツールを購入することを意味していました。ここで、代わりにループを構築します。データを監視して動作するループは、SaaS が販売したビジネス ロジックです。ただし、それを所有しているのはあなたです。 Lobu でビルドする方法は次のとおりです。
Shopify の帯水層、野外で
Shopify は、エージェントのコーパスが複利資産であると確信しています。私たちは同じ賭けをしましたが、2 つの違いがあります。チャットの代わりにシグナルを保持することと、1 つの企業ではなく多くの企業のためにそれを構築したことです。
エージェントメモリのファイルシステムとデータベース
エージェントには、考えるためのワークスペースと記憶するための倉庫が必要です。ファイルシステムは一時的な作業用です。記憶層は永続的な組織知識のためのものです。
監視して行動する AI のチームメイト。

## Original Extract

Open-source infrastructure for multi-tenant AI agents: sandboxed execution, shared memory, and an SDK that pulls live company data and acts on your high-level goals.

Lobu Docs Resources Sign in Build AI teammates that watch and act
Open-source infrastructure for multi-tenant AI agents: sandboxed execution, shared memory, and an SDK that pulls live company data and acts on your high-level goals.
ChatGPT Cursor OpenCode Copy setup prompt Talk to the founder × Paste into your coding agent to scaffold a project.
Or start it yourself: $ npx @lobu/cli@latest init my-agent
Connect your data. Set a goal. Lobu watches for changes and prepares the next step.
Pick the systems it can read. Lobu turns those updates into live customer memory.
Stripe Zendesk Gmail Google Drive GitHub Notion Snowflake PostgreSQL 50+ more Use existing connectors, or let your agent write code to connect any data. Watching changes across every source Connector SDK ↗ 2 Define the goal
Tell it what to watch for and when to ask before acting.
“Watch every account for churn risk. If renewal is within 30 days and health drops, draft a CSM check-in for approval.”
It scans memory on schedule, spots the account at risk, and keeps the evidence attached.
You can edit the draft, send it, or leave it.
R Revenue agent online Heads up: Acme Corp is trending toward churn. Logins are down 38% over 14 days and their renewal is in 21 days.
Want me to draft a check-in for their CSM? 12:01 Draft the email Yes, and include the usage drop. 12:01 Drafted. Saved it to the Acme account and pinged @dana. 12:01 See drafted email Replies in your team's chat, or your own app over the API: API ↗ Examples Explore agent workflows.
Each example shows the sources, memory, and actions for one AI teammate.
Local, self-hosted, or managed.
Boot the gateway, workers, memory, and embeddings with one command.
Deploy with Docker or Helm when data and controls need to stay with you.
Use the same project with managed isolation, secrets, and upgrades.
Build your first
multi-user agent.
Schedule a call × From the blog Latest blog posts
The Agent Loop Is the New SaaS
Software used to mean buying a vertical tool for every function. Now you build the loop instead. A loop that watches your data and acts is the business logic SaaS sold you, except you own it. Here is how to build one on Lobu.
Shopify's Aquifer, in the Open
Shopify bet that an agent's corpus is the compounding asset. We made the same bet, with two differences: we keep the signal instead of the chat, and we built it for many companies instead of one.
Filesystem vs Database for Agent Memory
Agents need a workspace to think in and a warehouse to remember in. The filesystem is for ephemeral work. The memory layer is for durable organizational knowledge.
AI teammates that watch and act.
