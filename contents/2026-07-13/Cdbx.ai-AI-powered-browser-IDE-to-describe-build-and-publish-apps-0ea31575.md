---
source: "https://cdbx.ai/"
hn_url: "https://news.ycombinator.com/item?id=48900140"
title: "Cdbx.ai – AI-powered browser IDE to describe, build, and publish apps"
article_title: "cdbx.ai"
author: "chbutler"
captured_at: "2026-07-13T23:43:36Z"
capture_tool: "hn-digest"
hn_id: 48900140
score: 3
comments: 1
posted_at: "2026-07-13T23:07:16Z"
tags:
  - hacker-news
  - translated
---

# Cdbx.ai – AI-powered browser IDE to describe, build, and publish apps

- HN: [48900140](https://news.ycombinator.com/item?id=48900140)
- Source: [cdbx.ai](https://cdbx.ai/)
- Score: 3
- Comments: 1
- Posted: 2026-07-13T23:07:16Z

## Translation

タイトル: Cdbx.ai – アプリを記述、構築、公開するための AI を活用したブラウザ IDE
記事のタイトル: cdbx.ai
説明: 次世代の開発者向けに構築された AI を活用した高速開発環境。

記事本文:
あなたはそれを説明します。 cdbx がそれを構築します。
すでにアカウントをお持ちですか?サインイン
コードを書くかどうかに関係なく、cdbx は役に立ちます。
開発者向け 本物の IDE。
ブラウザで。
デスクトップ IDE に期待されるすべてが、セットアップなしで実行できます。タブを開いて構築を開始します。
完全な Monaco エディター — VS Code と同じエンジン
ストリーミング出力を備えた実端末
ファイルを読み書きするAIペアプログラマー
インスタントプレビュー。 npm インストールはありません。
メーカー・クリエイター向け 未経験可
必要です。
あなたにはその考えがあります。 cdbxにはAIが搭載されています。欲しいものを説明して、それが実現するのを見てください。
欲しいものをわかりやすい英語で説明してください
アプリがリアルタイムで構築される様子を観察する
ワンクリックで公開して共有
セットアップもインストールも学習も必要ありません
MCP コネクタ あなたのツール、
文脈の中で。
Notion、GitHub、Slack、Google Drive など、あらゆる外部サービスを AI コンテキストに接続します。 AI はオンデマンドでクエリを実行します。コピー＆ペーストやタブの切り替えは必要ありません。
Notion、GitHub、Slack、Google Drive などを接続
AI は接続されたサービスをオンデマンドでクエリします - プリロードは何もありません
自己ホスト型またはサードパーティ製の MCP 互換サーバーであればどれでも動作します
ローカルで実行しますか? ngrokで公開して即座に接続
cdbx MCP サーバー経由で Claude Desktop または Cursor を実験に接続します
AI エージェント のエージェント
あらゆるステップ。
cdbx には AI チャット ウィンドウがあるだけではありません。エージェントは、アプリの生成からエラーの修正、UI の調整に至るまで、ワークフローのあらゆる部分に組み込まれています。
計画エージェント — 完全なコネクタ コンテキストを使用して PRD、アーキテクチャ、および調査を計画します。
Vibecoding は単純な説明から完全なアプリを生成します
自己修復機能により、作業を中断することなくビルド エラーを検出して修正します。
ビジュアル ループは、ライブ プレビューに対して UI を反復します。
ヘルプ アシスタントが作業中に cdbx に関する質問にインラインで回答します
y への完全なアクセス権を持つ戦略パートナー

私たちの接続されたサービス。実際に Notion ページ、Linear バックログ、GitHub リポジトリを読み取った AI を使用して、PRD の作成、システム アーキテクチャの設計、調査レポートの計画、またはスライド デッキの構築を行います。
一般、Web アプリ、アーキテクチャ、ディープ リサーチ、スライド、API 仕様 - それぞれに専用の AI とライブ プレビューが含まれています。
Planning Agent は、Notion、GitHub、Google Drive、Linear、Figma、Slack にオンデマンドでクエリを実行します。
エージェントはプランのファイル ツリーに直接書き込みます。 /PRD、/ADR、/RFC などをスラッシュ コマンドとして使用します。
計画が完了したら、ワンクリックで計画がコンテキストとして事前に読み込まれた新しい実験が開きます。
AI を使用して Web アプリを計画および構築する
システム設計、C4 ダイアグラム、および ADR
ソースをレポートに統合する
ライブプレビューを使用してスライドデッキを作成する
OpenAPI 3.0 を使用して REST API を設計する
アイデアから実際のアプリまで数分で完成します。
cdbx に何を構築したいかを平易な英語で伝えるか、テンプレートを選択します。
AI はすべてのファイルを即座に生成します。アプリはブラウザーですぐに実行されます。
ワンクリックで公開します。ライブアプリを世界中の誰とでも共有します。
あなたの言語。あなたのフレームワーク。
フロントエンドまたはバックエンド、スクリプトまたはコンパイル — cdbx はすべてを 1 か所で実行します。
React Vue Svelte Angular Preact JavaScript TypeScript バックエンド
Python Go Node.js Java Ruby Bun Rust Kotlin Runner API
コードを送信し、出力を取得します。インフラストラクチャも Docker もコールド スタートもありません。 1 回の API 呼び出しで安全なサンドボックス内でコードが実行され、ミリ秒単位で stdout、stderr、および終了コードが返されます。
curl -X POST https://api.cdbx.ai/v1/runner/execute \
-H "認証: ベアラー cdbx_your_key" \
-H "コンテンツ タイプ: application/json" \
-d '{
"言語": "Python",
"files": { "main.py": "print('Hello!')" },
"エントリポイント": "main.py"
}' | jq
# {
# "stdout": "こんにちは!\n",
# "終了コード": 0,
# "期間": 94
# } 30 言語、1 つのエンドポイント
Python、Go、ノード、Rust、

Java、C、C++、Swift、Haskell、Elixir、Zig、その他 19 種類 - すべて分離されたサンドボックスで実行されます。
ネットワーク分離、メモリ上限、PID 制限、seccomp プロファイル。各実行は新しいコンテナーで実行されます。
AI エージェントは、生成したコードを実行できます。 cdbx MCP サーバーは run_code をツールとして公開します。クロードはそれを直接呼び出すことができます。
無料利用枠には、1 か月あたり 100 回の実行が含まれます。 Proには2,000が含まれます。さらに必要な場合は、トップアップ パックを購入してください。有効期限はありません。
Python Go Node.js TypeScript Java Kotlin Rust C# Ruby C C++ PHP Swift Scala Haskell Elixir Clojure Bash Perl Lua R D Nim Zig OCaml F# Julia Dart Crystal Groovy API ドキュメントを読む 価格
無料で始めましょう。準備ができたらスケールします。
無料で始められます。クレジットカードは必要ありません。
永久無料。クレジットカードは必要ありません。
· 無料で構築を開始するパブリックおよびプライベートの実験
AIアシスタント（Haiku + GPT-4o-mini）
ZIP としてエクスポートするか、GitHub にプッシュします
7日間の無料トライアル。いつでもキャンセルできます。
· 無料トライアルを開始する すべてが無料です
すべての AI モデル — Sonnet、Opus、GPT-4o
公開されたアプリには cdbx バッジがありません
一緒に構築するチーム向けに構築されています。
ライブコラボレーションとペアコーディング
API キーは保存時に暗号化されます 独自の Anthropic キーまたは OpenAI キーを持ち込んでいつでもコードを GitHub にエクスポートできます アイデアを現実化できます
cdbxで。
cdbx を使用して構築する何千人もの開発者やクリエイターに加わりましょう。

## Original Extract

An AI-powered Rapid Development Environment built for the next generation of developers.

You describe it. cdbx builds it.
Already have an account? Sign in
Whether you code or not, cdbx works for you.
For developers A real IDE.
In your browser.
Everything you'd expect from a desktop IDE — without the setup. Open a tab and start building.
Full Monaco editor — same engine as VS Code
Real terminal with streaming output
AI pair programmer that reads and writes your files
Instant preview. No npm install.
For makers & creators No experience
needed.
You have the idea. cdbx has the AI. Just describe what you want and watch it come to life.
Describe what you want in plain English
Watch your app get built in real-time
One click to publish and share
No setup, no installs, no learning curve
MCP Connectors Your tools,
in context.
Connect any external service to your AI context — Notion, GitHub, Slack, Google Drive, and more. The AI queries them on demand. No copy-pasting, no tab-switching.
Connect Notion, GitHub, Slack, Google Drive, and more
AI queries connected services on demand — nothing preloaded
Any MCP-compatible server works — self-hosted or third-party
Running locally? Expose it with ngrok and connect instantly
Connect Claude Desktop or Cursor to your experiments via the cdbx MCP server
AI Agents Agents in
every step.
cdbx doesn't just have an AI chat window — agents are woven into every part of the workflow, from generating your app to fixing errors to refining your UI.
Planning Agent — plan PRDs, architecture, and research with full connector context
Vibe Coding generates complete apps from a plain description
Self-Healing detects and fixes build errors without interrupting you
Visual Loop iterates your UI against the live preview
Help Assistant answers questions about cdbx inline as you work
A strategy partner with full access to your connected services. Write PRDs, design system architecture, plan a research report, or build a slide deck — with an AI that has actually read your Notion pages, Linear backlog, and GitHub repos.
General, Web App, Architecture, Deep Research, Slides, and API Spec — each with a purpose-built AI and live preview.
The Planning Agent queries your Notion, GitHub, Google Drive, Linear, Figma, and Slack on demand.
The agent writes directly to your plan's file tree. /PRD, /ADR, /RFC, and more as slash commands.
When your plan is done, one click opens a new experiment pre-loaded with your plan as context.
Plan and build a web app with AI
System design, C4 diagrams, and ADRs
Synthesise sources into a report
Build a slide deck with live preview
Design REST APIs with OpenAPI 3.0
From idea to live app in minutes.
Tell cdbx what you want to build in plain English, or pick a template.
AI generates all your files instantly. Your app runs in the browser immediately.
One click to publish. Share your live app with anyone in the world.
Your language. Your framework.
Frontend or backend, scripted or compiled — cdbx runs it all in one place.
React Vue Svelte Angular Preact JavaScript TypeScript Backend
Python Go Node.js Java Ruby Bun Rust Kotlin Runner API
Send code, get output. No infrastructure, no Docker, no cold starts. A single API call runs your code in a secure sandbox and returns stdout, stderr, and exit code in milliseconds.
curl -X POST https://api.cdbx.ai/v1/runner/execute \
-H "Authorization: Bearer cdbx_your_key" \
-H "Content-Type: application/json" \
-d '{
"language": "python",
"files": { "main.py": "print('Hello!')" },
"entrypoint": "main.py"
}' | jq
# {
# "stdout": "Hello!\n",
# "exitCode": 0,
# "durationMs": 94
# } 30 languages, one endpoint
Python, Go, Node, Rust, Java, C, C++, Swift, Haskell, Elixir, Zig, and 19 more — all running in isolated sandboxes.
Network isolation, memory caps, PID limits, seccomp profiles. Each execution runs in a fresh container.
Your AI agent can run code it generates. The cdbx MCP server exposes run_code as a tool — Claude can call it directly.
Free tier includes 100 executions/month. Pro includes 2,000. Buy top-up packs when you need more — they never expire.
Python Go Node.js TypeScript Java Kotlin Rust C# Ruby C C++ PHP Swift Scala Haskell Elixir Clojure Bash Perl Lua R D Nim Zig OCaml F# Julia Dart Crystal Groovy Read the API docs Pricing
Start free. Scale when you're ready.
Free to start. No credit card required.
Free forever. No credit card needed.
· Start building free Public & private experiments
AI assistant (Haiku + GPT-4o-mini)
Export as ZIP or push to GitHub
7-day free trial. Cancel anytime.
· Start free trial Everything in Free
All AI models — Sonnet, Opus, GPT-4o
No cdbx badge on published apps
Built for teams that build together.
Live collaboration & pair coding
API keys encrypted at rest Bring your own Anthropic or OpenAI key Export your code to GitHub any time Bring your idea to life
on cdbx.
Join thousands of developers and creators building with cdbx.
