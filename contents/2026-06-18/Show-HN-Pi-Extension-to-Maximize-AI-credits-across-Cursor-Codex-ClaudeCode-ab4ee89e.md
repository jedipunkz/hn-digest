---
source: "https://github.com/sathish316/pi-omniagent-extensions/"
hn_url: "https://news.ycombinator.com/item?id=48587117"
title: "Show HN: Pi Extension to Maximize AI credits across Cursor, Codex, ClaudeCode"
article_title: "GitHub - sathish316/pi-omniagent-extensions: Pi coding agent extensions that helps you connect to all coding agents from one place using ACP and maximize your AI credits across Cursor, Codex, Claude Code, Rovo · GitHub"
author: "sathish316"
captured_at: "2026-06-18T16:13:05Z"
capture_tool: "hn-digest"
hn_id: 48587117
score: 1
comments: 0
posted_at: "2026-06-18T15:41:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pi Extension to Maximize AI credits across Cursor, Codex, ClaudeCode

- HN: [48587117](https://news.ycombinator.com/item?id=48587117)
- Source: [github.com](https://github.com/sathish316/pi-omniagent-extensions/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T15:41:44Z

## Translation

タイトル: HN を表示: Cursor、Codex、ClaudeCode 全体で AI クレジットを最大化するための Pi 拡張機能
記事のタイトル: GitHub - sathish316/pi-omniagent-extensions: ACP を使用して 1 か所からすべてのコーディング エージェントに接続し、Cursor、Codex、Claude Code、Rovo 全体で AI クレジットを最大化するのに役立つ Pi コーディング エージェント拡張機能 · GitHub
説明: ACP を使用して 1 か所からすべてのコーディング エージェントに接続し、Cursor、Codex、Claude Code、Rovo 全体で AI クレジットを最大化するのに役立つ Pi コーディング エージェント拡張機能 - sathish316/pi-omniagent-extensions
HN テキスト: ACP ( https://agentclientprotocol.com/ ) を使用してコーディング エージェントに接続し、Cursor、Codex、ClaudeCode、RovoDev などの複数のサブスクリプション全体で AI クレジットを最大化するのに役立つ Pi 拡張機能

記事本文:
GitHub - sathish316/pi-omniagent-extensions: ACP を使用して 1 か所からすべてのコーディング エージェントに接続し、Cursor、Codex、Claude Code、Rovo 全体で AI クレジットを最大化するのに役立つ Pi コーディング エージェント拡張機能 · GitHub
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
サティシュ316
/
ピオムニア

ジェント拡張機能
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
sathish316/pi-omniagent-extensions
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .gitignore .gitignore HOW_IT_WORKS.md HOW_IT_WORKS.md README.md README.md USAGE_GUIDE.md USAGE_GUIDE.md claude-code-acp.ts claude-code-acp.ts codex-app-server.ts codex-app-server.tsカーソル-acp.ts カーソル-acp.ts パッケージ-lock.json パッケージ-lock.json パッケージ.json パッケージ.json rovo-acp.ts rovo-acp.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Pi エージェント拡張機能 — それぞれの仕組み
共通の考え方 (最初にお読みください)
4 つのファイルはすべて同じ仕事をします。つまり、Pi (ホスト コーディング エージェント) が、まるで Pi の /model ピッカーの別のモデルであるかのように、別のコーディング エージェントの CLI を駆動できるようになります。 ACP ( https://agentclientprotocol.com/ ) を使用してコーディング エージェントと通信します。 Pi で「Cursor Sonnet」または「GPT-5 [codex-app-server]」または「Opus [claude-code-acp]」を選択すると、ワークスペースで実行されている外部エージェントによって秘密裏にターンが実行されます。
┌───────────────────────────┐
│ Pi ホストプロセス │
│ │
│ /モデルピッカー ──► pi.registerProvider(...) │
│ │ │
│ │ streamSimple(モデル, コンテキスト) │
│ ▼ │
│ ┌─────────┐ ビルドプロンプト ┌───────────┐ │
│ │ プロバイダ │ ───────────► │ ブリッジ（シングルトン） │ │
│ │ stream fn │ ◄───────────── │ をシリアル化します

骨壷、 │ │
│ └─────────┘ ストリームイベント │ 1 セッションを所有 │ │
│ ▲ └─────┬───────┘ │
│ │ AssistantMessageEventStream │ スポーン │
│ │ (本文/思考デルタ) ▼ │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│ │
│ ▼
│ ┌───────┬─────┐
└──── JSON-RPC over stdio ──► │ 子プロセス (CLI) │
│ カーソルエージェント / rovo │
│ コーデックス / クロード-ACP │
━━━━━━━━━━━┘
バンドルされている拡張機能:
Cursor-acp.ts — ACP 上のカーソル エージェント
codex-app-server.ts — アプリサーバー上の OpenAI Codex
claude-code-acp.ts — ACP 上の Claude Code (公式 ACP SDK を使用)
rovo-acp.ts — ACP を使用した Atlassian Rovo Dev
~/.pi/agent/extensions へのインストール
Pi は ~/.pi/agent/extensions 内の *.ts ファイルを自動検出しますが、外部の npm 依存関係はこれらのファイルにバンドルされません。
git clone https://github.com/sathish316/pi-omniagent-extensions.git
cd pi-omniagent-extensions
拡張ファイルを Pi 拡張ディレクトリにコピーします。
mkdir -p ~ /.pi/agent/extensions
cp カーソル-acp.ts rovo-acp.ts codex-app-server.ts claude-code-acp.ts \
package.json ~ /.pi/agent/extensions/
拡張機能の依存関係をインストールします。
cd ~ /.pi/agent/extensions
npmインストール
npm run check:claude-code-acp-deps
基礎となるエージェント CLI がインストールされており、その PATH 上にあることを確認してください。
使用する拡張子:cursor-agent (cursor-acp)、rovo (rovo-acp)、
コーデックス (コーデックス-アプリ-サーバー)。 claude-code-acp は n からランタイムを取得します

午後。
Piを再起動します。ブリッジされたモデルは /model ピッカーにタグ付けされて表示されます。
プロバイダー、例: [codex-app-server] または [claude-code-acp] 。
claude-code-acp.ts 拡張機能には以下が必要です。
@agentclientprotocol/claude-agent-acp
拡張機能によってインポートされた Pi API パッケージは、実行中の Pi インストールによって提供されます。
Pi コーディング エージェント拡張機能により、ACP を使用して 1 か所からすべてのコーディング エージェントに接続し、Cursor、Codex、Claude Code、Rovo 全体で AI クレジットを最大化できます。
読み込み中にエラーが発生しました。このページをリロードしてください。
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

Pi coding agent extensions that helps you connect to all coding agents from one place using ACP and maximize your AI credits across Cursor, Codex, Claude Code, Rovo - sathish316/pi-omniagent-extensions

Pi extension that uses ACP ( https://agentclientprotocol.com/ ) to connect to any Coding agent and helps you Maximize AI Credits across multiple subscriptions like Cursor, Codex, ClaudeCode, RovoDev and more

GitHub - sathish316/pi-omniagent-extensions: Pi coding agent extensions that helps you connect to all coding agents from one place using ACP and maximize your AI credits across Cursor, Codex, Claude Code, Rovo · GitHub
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
sathish316
/
pi-omniagent-extensions
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
sathish316/pi-omniagent-extensions
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .gitignore .gitignore HOW_IT_WORKS.md HOW_IT_WORKS.md README.md README.md USAGE_GUIDE.md USAGE_GUIDE.md claude-code-acp.ts claude-code-acp.ts codex-app-server.ts codex-app-server.ts cursor-acp.ts cursor-acp.ts package-lock.json package-lock.json package.json package.json rovo-acp.ts rovo-acp.ts View all files Repository files navigation
Pi Agent Extensions — How Each One Works
The common idea (read this first)
All four files do the same job: they let Pi (the host coding agent) drive another coding agent's CLI as if it were just another model in Pi's /model picker. It uses ACP ( https://agentclientprotocol.com/ ) to communicate with any coding agent. You pick "Cursor Sonnet" or "GPT-5 [codex-app-server]" or "Opus [claude-code-acp]" in Pi, and your turn is secretly executed by that external agent running in your workspace.
┌────────────────────────────────────────────────────────────────────┐
│ Pi host process │
│ │
│ /model picker ──► pi.registerProvider(...) │
│ │ │
│ │ streamSimple(model, context) │
│ ▼ │
│ ┌─────────────┐ builds prompt ┌──────────────────────┐ │
│ │ Provider │ ─────────────────► │ Bridge (singleton) │ │
│ │ stream fn │ ◄───────────────── │ serializes turns, │ │
│ └─────────────┘ stream events │ owns 1 session │ │
│ ▲ └──────────┬───────────┘ │
│ │ AssistantMessageEventStream │ spawn │
│ │ (text/thinking deltas) ▼ │
└────────┼────────────────────────────────────────┼──────────────────┘
│ │
│ ▼
│ ┌───────────┬──────────┐
└──── JSON-RPC over stdio ───► │ Child process (CLI) │
│ cursor-agent / rovo │
│ codex / claude-acp │
└──────────────────────┘
The bundled extensions:
cursor-acp.ts — Cursor Agent over ACP
codex-app-server.ts — OpenAI Codex over app-server
claude-code-acp.ts — Claude Code over ACP (uses the official ACP SDK)
rovo-acp.ts — Atlassian Rovo Dev over ACP
Installing into ~/.pi/agent/extensions
Pi auto-discovers *.ts files in ~/.pi/agent/extensions , but external npm dependencies are not bundled into those files.
git clone https://github.com/sathish316/pi-omniagent-extensions.git
cd pi-omniagent-extensions
Copy the extension files into your Pi extensions directory:
mkdir -p ~ /.pi/agent/extensions
cp cursor-acp.ts rovo-acp.ts codex-app-server.ts claude-code-acp.ts \
package.json ~ /.pi/agent/extensions/
Install the extension dependencies:
cd ~ /.pi/agent/extensions
npm install
npm run check:claude-code-acp-deps
Make sure the underlying agent CLIs are installed and on your PATH for the
extensions you want to use: cursor-agent (cursor-acp), rovo (rovo-acp),
codex (codex-app-server). claude-code-acp pulls its runtime from npm.
Restart Pi. The bridged models appear in the /model picker tagged with
their provider, e.g. [codex-app-server] or [claude-code-acp] .
The claude-code-acp.ts extension needs:
@agentclientprotocol/claude-agent-acp
The Pi API packages imported by the extensions are provided by the running Pi installation.
Pi coding agent extensions that helps you connect to all coding agents from one place using ACP and maximize your AI credits across Cursor, Codex, Claude Code, Rovo
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
