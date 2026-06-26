---
source: "https://github.com/CopilotKit/OpenTag/"
hn_url: "https://news.ycombinator.com/item?id=48692614"
title: "OpenTag: An open-source alternative to Claude in Slack"
article_title: "GitHub - CopilotKit/OpenTag: OpenTag · GitHub"
author: "davidmckayv"
captured_at: "2026-06-26T22:24:54Z"
capture_tool: "hn-digest"
hn_id: 48692614
score: 2
comments: 0
posted_at: "2026-06-26T22:09:07Z"
tags:
  - hacker-news
  - translated
---

# OpenTag: An open-source alternative to Claude in Slack

- HN: [48692614](https://news.ycombinator.com/item?id=48692614)
- Source: [github.com](https://github.com/CopilotKit/OpenTag/)
- Score: 2
- Comments: 0
- Posted: 2026-06-26T22:09:07Z

## Translation

タイトル: OpenTag: Slack の Claude に代わるオープンソースの代替ツール
記事タイトル: GitHub - CopilotKit/OpenTag: OpenTag · GitHub
説明: オープンタグ。 GitHub でアカウントを作成して、CopilotKit/OpenTag の開発に貢献してください。

記事本文:
GitHub - CopilotKit/OpenTag: OpenTag · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
コパイロットキット
/
オープンタグ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
C

頌歌
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット app app e2e e2e scripts scripts .env.example .env.example .gitignore .gitignore .npmrc .npmrc ライセンス ライセンス README.md README.md docker-compose.yml docker-compose.yml package.json package.json runtime.ts runtime.ts setup.md setup.mdlack-app-manifest.jsonslack-app-manifest.jsonslack-app-manifest.yamlslack-app-manifest.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenTag: Slack の Claude に代わるオープンソースの代替ツール
Slack 内で独自の AI エージェントを実行します。スレッドを読み取り、応答し、ツールを呼び出し、
会話の中で豊富な結果を提供します。あなたの中にクロードがいると考えてください。
ワークスペース (オープンソースとセルフホストを除く) : ランタイムは自分で所有し、独自のものを持ち込む
モデルを作成し、それを独自のツールに接続します。座席ごとの価格設定やロックインはありません。
@copilotkit/bot に基づいて構築されています —
CopilotKit のチャット プラットフォーム エージェント用のオープン SDK (最初は Slack、同じコードは Slack でも実行されます)
Discord、Telegram、WhatsApp)。クローンを作成し、モデルとツールを指定すれば、自分のものになります。
スタック全体。
デモ-slack-opentag.mp4
▶️ デモを見る (~50 秒) — Slack スレッドを動作させる OpenTag エージェント: 内訳、表、棒グラフをインラインでレンダリングし (生成 UI)、承認ゲートの後にのみチケットを提出します (人間参加型)。
実行するには 2 つの方法があります。以下のオープンソース SDK を使用して自分でホストするか、操作をスキップしてマネージド サービスにサインアップします (CopilotKit から近日提供予定)。
OpenTag は CopilotKit モノリポジトリ内に
第一級のサンプル (examples/slack)。それが今日でも信頼できる方法です。
ボット SDK パッケージは npm への公開を終了します。 (このリポジトリからスタンドアロンの npm インストールを実行すると、
着陸した瞬間に

— setup.md を参照してください。)
エージェント (LLM バックエンド) とボット (Slack
接続) — そして 3 つのシークレットを設定します。
1. Slack アプリを作成します。 api.slack.com/apps で→
マニフェストから→slack-app-manifest.yamlを貼り付けます。インストールして、
次に、ボット ユーザー OAuth トークン ( xoxb-… ) とアプリレベル トークン ( xapp-… ) を取得します。
接続:書き込みスコープ)。 setup.md のステップバイステップ。
2. .env ( cp .env.example .env ) に 3 つのシークレットを設定します。
SLACK_BOT_TOKEN=xoxb-...
SLACK_APP_TOKEN=xapp-...
OPENAI_API_KEY=sk-... # または ANTHROPIC_API_KEY — 独自のモデルを使用する
3. CopilotKit モノリポジトリのルートから実行します。
pnpmインストール
pnpm --filterlack-example runtime # エージェント バックエンド、:8200
pnpm --filterlack-example dev # ボット
4. 話しかけてください。任意のチャネル スレッドでボットを @メンションします。
@OpenTag はこのスレッドを要約し、バグとして報告します
それがループ全体です。 Linear、Notion、インライン チャート、Redis 永続性を接続するか、実行するには
Discord / Telegram / WhatsApp については、 setup.md を参照してください。
ただし、私たちはあなたに嘘はつきません。チャット エージェントのホスティングを設定するのは簡単ではありません。そのような心痛をすべて回避するには、Intelligence プラットフォームでホストされている CopilotKit 管理サービスの待機リストに参加してください。
OpenTag は意図的に小さく、ハッキング可能です。
動作を変更します。エージェントの動作は、単一のシステム プロンプトによって制御されます。
runtime.ts — これを書き直すと、別のエージェントが作成されます。
app/ をコピーして、独自のボットを起動します。これはプラットフォームに依存しないボット (ツール、コンポーネント、
人間参加型ゲート)。 runtime.ts はエージェント バックエンドです。1 つの CopilotKit BuiltInAgent (
LLM + オプションの MCP ツール (Python、LangGraph なし)、AG-UI 経由で提供されます。
1 つのプラットフォームでも、すべてのプラットフォームでも。 createBot はアダプターの配列を受け取ります。シークレットを設定する
必要なプラットフォームを選択すると、ボットがそれぞれのアダプターを開始します。
完全なアーキテクチャ

ファイルごとのマップ、およびすべての統合が存在します。
setup.md 。
自分でホストしたくないですか?
セルフホスティングとは、ランタイム、永続性、および検査ツールを自分で実行および拡張することを意味します。
マネージド CopilotKit サービスが準備中です。操作を除いた同じエージェントです: 耐久性があります
スレッド、永続性、ホストされた検査、フィードバックから改善するエージェント (継続的)
人間のフィードバックから学ぶ)。
待機リストに参加する → — マネージド サービスが開始されたときに最初に参加します。
エンジニアに相談 → — これを基に何か実際のものを構築しますか?発送のお手伝いをさせていただきます。
CopilotKit Slack クイックスタートは正規のガイドです
Slack エージェントの構築については、このスターターと一緒にお読みください。詳細なセットアップと構成
setup.md にあります。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

OpenTag. Contribute to CopilotKit/OpenTag development by creating an account on GitHub.

GitHub - CopilotKit/OpenTag: OpenTag · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
CopilotKit
/
OpenTag
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits app app e2e e2e scripts scripts .env.example .env.example .gitignore .gitignore .npmrc .npmrc LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml package.json package.json runtime.ts runtime.ts setup.md setup.md slack-app-manifest.json slack-app-manifest.json slack-app-manifest.yaml slack-app-manifest.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
OpenTag: an open-source alternative to Claude in Slack
Run your own AI agent inside Slack: it reads a thread, answers, calls your tools, and
renders rich results right in the conversation. Think of it as having Claude in your
workspace, except open-source and self-hosted : you own the runtime, bring your own
model, and wire it to your own tools. No per-seat pricing, no lock-in.
It's built on @copilotkit/bot —
CopilotKit's open SDK for chat-platform agents (Slack first; the same code also runs on
Discord, Telegram, and WhatsApp). Clone it, point it at your model and tools, and you own
the whole stack.
demo-slack-opentag.mp4
▶️ Watch the demo (~50s) — an OpenTag agent working a Slack thread: it renders a breakdown, a table, and a bar chart inline ( generative UI ) and files a ticket only after an Approve gate ( human-in-the-loop ).
Two ways to run it: host it yourself with the open-source SDK below — or skip the ops and sign up for the managed service → coming soon from CopilotKit.
OpenTag ships inside the CopilotKit monorepo as a
first-class example ( examples/slack ). That's the dependable way to run it today while the
bot SDK packages finish publishing to npm. (A standalone npm install from this repo lights
up the moment they land — see setup.md .)
You'll run two processes: the agent (the LLM backend) and the bot (the Slack
connection) — and set three secrets.
1. Create a Slack app. At api.slack.com/apps →
From a manifest → paste slack-app-manifest.yaml . Install it,
then grab the Bot User OAuth Token ( xoxb-… ) and an App-Level Token ( xapp-… , with the
connections:write scope). Step-by-step in setup.md .
2. Set three secrets in .env ( cp .env.example .env ):
SLACK_BOT_TOKEN=xoxb-...
SLACK_APP_TOKEN=xapp-...
OPENAI_API_KEY=sk-... # or ANTHROPIC_API_KEY — bring your own model
3. Run it from the CopilotKit monorepo root:
pnpm install
pnpm --filter slack-example runtime # the agent backend, on :8200
pnpm --filter slack-example dev # the bot
4. Talk to it. @mention the bot in any channel thread:
@OpenTag summarize this thread and file it as a bug
That's the whole loop. To wire up Linear, Notion, inline charts, Redis persistence, or to run
on Discord / Telegram / WhatsApp, see setup.md .
We won't lie to you, though. Setting up hosting for chat agents is not easy. To skip all of that heartache, go join the waitlist for the CopilotKit managed service hosted on our Intelligence platform.
OpenTag is deliberately small and hackable:
Change what it does. The agent's behavior is steered by a single system prompt in
runtime.ts — rewrite it and you have a different agent.
Copy app/ to start your own bot. It's the platform-agnostic bot (tools, components, the
human-in-the-loop gate). runtime.ts is the agent backend: one CopilotKit BuiltInAgent (an
LLM + optional MCP tools — no Python, no LangGraph), served over AG-UI.
One platform, or all of them. createBot takes an array of adapters; set the secrets for
whichever platform(s) you want and the bot starts an adapter for each.
The full architecture, the file-by-file map, and every integration live in
setup.md .
Don't want to host it yourself?
Self-hosting means you run and scale the runtime, persistence, and inspection tooling yourself.
A managed CopilotKit service is on its way. It's the same agent, without the ops: durable
threads, persistence, hosted inspection, and agents that improve from feedback ( Continuous
Learning from Human Feedback ).
Join the waitlist → — be first in when the managed service opens.
Talk to an engineer → — building something real on this? We'd love to help you ship it.
The CopilotKit Slack quickstart is the canonical guide
to building a Slack agent — read it alongside this starter. Detailed setup and configuration
lives in setup.md .
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
