---
source: "https://github.com/LiteLLM-Labs/lite-harness"
hn_url: "https://news.ycombinator.com/item?id=48341726"
title: "Show HN: Lite-Harness – Self-Hosted Cursor Agents (Use Claude Code/OpenCode)"
article_title: "GitHub - LiteLLM-Labs/lite-harness: Unified Server for running OpenCode, Claude Code, Codex · GitHub"
author: "detente18"
captured_at: "2026-05-31T00:23:51Z"
capture_tool: "hn-digest"
hn_id: 48341726
score: 3
comments: 0
posted_at: "2026-05-30T23:51:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lite-Harness – Self-Hosted Cursor Agents (Use Claude Code/OpenCode)

- HN: [48341726](https://news.ycombinator.com/item?id=48341726)
- Source: [github.com](https://github.com/LiteLLM-Labs/lite-harness)
- Score: 3
- Comments: 0
- Posted: 2026-05-30T23:51:21Z

## Translation

タイトル: HN を表示: Lite-Harness – 自己ホスト型カーソル エージェント (クロード コード/オープンコードを使用)
記事のタイトル: GitHub - LiteLLM-Labs/lite-harness: OpenCode、Claude Code、Codex を実行するための統合サーバー · GitHub
説明: OpenCode、Claude Code、Codex を実行するための統合サーバー - LiteLLM-Labs/lite-harness
HN テキスト: 私たちがこの Dockerfile を構築したのは、エージェントを実行し、メモリ、永続的なセッション、cron スケジューリング、およびボールトをすぐに利用できるシンプルなハーネス サーバーが必要だったからです。

記事本文:
GitHub - LiteLLM-Labs/lite-harness: OpenCode、Claude Code、Codex を実行するための統合サーバー · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
LiteLLM-Labs
/
ライトハーネス
公共
通知
通知設定を変更するにはサインインする必要があります
あ

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
292 コミット 292 コミット .githooks .githooks .github/ workflows .github/ workflows cli cli docs docs harnesses harnesses mcp mcp ui ui .gitignore .gitignore AGENTS.md AGENTS.md CODING_STANDARDS.md CODING_STANDARDS.md Dockerfile Dockerfile README.md README.md VISION.md VISION.md start-local.sh start-local.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自動操縦のサンドボックスで Claude Code/Codex/OpenCode を実行します。 CLI または OpenCode 互換の API 仕様を介した API を介してハーネスを呼び出します。
# git clone https://github.com/LiteLLM-Labs/lite-harness && cd lite-harness
# cd cli && npm install -g 。
ライトログイン
# サーバー URL [http://localhost:4096]:
# マスターキー: ✓ 保存されました
ライトクロードコード
ライトクロードコード
クロード-ソネット-4-6 · ローカルホスト:4096 · sess_a1b2c3
/clear で履歴をリセット · Ctrl+C または " exit " で終了
❯ CI を 1 時間ごとに監視し、バグを修正する
⠙考え中…
✓ bash { " コマンド " : " gh run list --limit 5 " }
定期的な CI モニターをセットアップします。過去 5 回の実行を確認しています...
ライトオープンコード
ライトオープンコード
クロード-ソネット-4-6 · ローカルホスト:4096 · sess_d4e5f6
/clear で履歴をリセット · Ctrl+C または " exit " で終了
❯ github stargazers 毎日DMしてください
⠙考え中…
✓ bash { " コマンド " : " gh api /repos/LiteLLM-Labs/lite-harness/stargazers " }
新しいスターゲイザーを 42 人獲得しました。 DM の下書き中...
サポートされているエージェント: opencode claude-code github-copilot codex — 完全な CLI ドキュメント
エクスポート MASTER_KEY= $( openssl rand -hex 32 )
エコー「 MASTER_KEY: $MASTER_KEY 」
docker run -p 4096:4096 \
-e LITELLM_API_BASE=https://your-litellm-gateway \
-e LITELLM_API_KEY=sk-... \
-e MASTER_KEY= " $MASTER_KEY " \
ghcr.io/litellm-labs/lite-harness:最新
localhost:4096 を開き、ログイン ページにマスター キーを貼り付けます。
ニーズ

LiteLLM ゲートウェイ。完全な構成: docs/configuration.md 。
デフォルトでは、サーバーの再起動時にセッション (履歴、モデル コンテキスト) が失われます。再起動後もそれらを保持するには、データ ディレクトリに永続ストレージをマウントします。
docker run -p 4096:4096 \
-v ./data:/home/sandbox/.local/share/lite-harness \
-e LITELLM_API_BASE=... \
-e LITELLM_API_KEY=... \
-e MASTER_KEY=... \
ghcr.io/litellm-labs/lite-harness:最新
カスタム パス (例: /mnt/data にマウントされたクラウド ボリューム):
-e DB_PATH=/mnt/data/db.db
再起動すると、サーバーはデータベースからハイドレートされた N セッションをログに記録し、以前のすべてのセッションがすぐに利用可能になります。
E2B_API_KEY または DAYTONA_API_KEY を設定すると、エージェントは隔離された Linux サンドボックスを自動的に取得します。フルオプション (テンプレート、スナップショット、ボールト): docs/configuration.md 。
私たちがライトハーネスを構築したのは、オープンコードとクロードコードを別個のサーバーとして実行することにより、複数のサービス、異なる API 仕様、信頼性の低いセッション管理、MCP ツールとシステム プロンプトへの異なる入力などの維持が困難になったためです。
そこで、すべてのハーネスを OpenCode 互換サーバーにラップし、1 つの Dockerfile に配置して、すべてのハーネスにわたる共有 MCP ツール、プロンプト、セッション管理を使用して、保守する 1 つのサービスを提供しました。
API リファレンス · アーキテクチャ · 構成 · CLI · ハーネスの追加
OpenCode、Claude Code、Codex を実行するための統合サーバー
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Unified Server for running OpenCode, Claude Code, Codex - LiteLLM-Labs/lite-harness

We built this Dockerfile because we wanted a simple harness server to run our agents and get memory, durable sessions, cron scheduling, and a vault, out of the box.

GitHub - LiteLLM-Labs/lite-harness: Unified Server for running OpenCode, Claude Code, Codex · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
LiteLLM-Labs
/
lite-harness
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
292 Commits 292 Commits .githooks .githooks .github/ workflows .github/ workflows cli cli docs docs harnesses harnesses mcp mcp ui ui .gitignore .gitignore AGENTS.md AGENTS.md CODING_STANDARDS.md CODING_STANDARDS.md Dockerfile Dockerfile README.md README.md VISION.md VISION.md start-local.sh start-local.sh View all files Repository files navigation
Run Claude Code/Codex/OpenCode on a sandbox in autopilot. Call the harnesses via a CLI or API's through an OpenCode-compatible API spec.
# git clone https://github.com/LiteLLM-Labs/lite-harness && cd lite-harness
# cd cli && npm install -g .
lite login
# Server URL [http://localhost:4096]:
# Master key: ✓ Saved
lite claude-code
lite claude-code
claude-sonnet-4-6 · localhost:4096 · sess_a1b2c3
/clear to reset history · Ctrl+C or " exit " to quit
❯ monitor CI every hour and fix any bugs
⠙ thinking…
✓ bash { " command " : " gh run list --limit 5 " }
I ' ll set up a recurring CI monitor. Checking the last 5 runs now...
lite opencode
lite opencode
claude-sonnet-4-6 · localhost:4096 · sess_d4e5f6
/clear to reset history · Ctrl+C or " exit " to quit
❯ dm github stargazers daily
⠙ thinking…
✓ bash { " command " : " gh api /repos/LiteLLM-Labs/lite-harness/stargazers " }
Got 42 new stargazers. Drafting DMs...
Supported agents: opencode claude-code github-copilot codex — full CLI docs
export MASTER_KEY= $( openssl rand -hex 32 )
echo " MASTER_KEY: $MASTER_KEY "
docker run -p 4096:4096 \
-e LITELLM_API_BASE=https://your-litellm-gateway \
-e LITELLM_API_KEY=sk-... \
-e MASTER_KEY= " $MASTER_KEY " \
ghcr.io/litellm-labs/lite-harness:latest
Open localhost:4096 , paste the master key on the login page.
Needs a LiteLLM gateway. Full config: docs/configuration.md .
By default, sessions (history, model context) are lost when the server restarts. To keep them across restarts, mount persistent storage at the data directory:
docker run -p 4096:4096 \
-v ./data:/home/sandbox/.local/share/lite-harness \
-e LITELLM_API_BASE=... \
-e LITELLM_API_KEY=... \
-e MASTER_KEY=... \
ghcr.io/litellm-labs/lite-harness:latest
Custom path (e.g. a mounted cloud volume at /mnt/data ):
-e DB_PATH=/mnt/data/db.db
On restart the server logs hydrated N session(s) from db and all prior sessions are immediately available.
Set E2B_API_KEY or DAYTONA_API_KEY and agents get an isolated Linux sandbox automatically. Full options (templates, snapshots, vault): docs/configuration.md .
We built lite-harness because running opencode and claude-code as separate servers got hard to maintain - multiple services, different API specs, unreliable session management, different inputs for MCP tools and system prompts.
So we wrapped all harnesses in an OpenCode-compatible server and put it in one Dockerfile, giving us one service to maintain, with shared MCP tools, prompts, and session management across all harnesses.
API reference · Architecture · Configuration · CLI · Add a harness
Unified Server for running OpenCode, Claude Code, Codex
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
