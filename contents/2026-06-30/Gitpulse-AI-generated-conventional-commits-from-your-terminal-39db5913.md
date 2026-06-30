---
source: "https://github.com/erico964-blip/gitpulse"
hn_url: "https://news.ycombinator.com/item?id=48739045"
title: "Gitpulse – AI-generated conventional commits from your terminal"
article_title: "GitHub - erico964-blip/gitpulse: ⚡ AI-powered conventional commit message generator for Git · GitHub"
author: "erico964-blip"
captured_at: "2026-06-30T21:34:32Z"
capture_tool: "hn-digest"
hn_id: 48739045
score: 2
comments: 0
posted_at: "2026-06-30T20:52:51Z"
tags:
  - hacker-news
  - translated
---

# Gitpulse – AI-generated conventional commits from your terminal

- HN: [48739045](https://news.ycombinator.com/item?id=48739045)
- Source: [github.com](https://github.com/erico964-blip/gitpulse)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T20:52:51Z

## Translation

タイトル: Gitpulse – AI が端末から生成した従来のコミット
記事タイトル: GitHub - erico964-blip/gitpulse: ⚡ AI を活用した従来の Git 用コミット メッセージ ジェネレーター · GitHub
説明: ⚡ AI を利用した従来の Git 用コミット メッセージ ジェネレーター - erico964-blip/gitpulse

記事本文:
GitHub - erico964-blip/gitpulse: ⚡ AI を活用した従来の Git 用コミット メッセージ ジェネレーター · GitHub
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
erico964-ブリップ
/
ギットパルス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ママ

ブランチ内 タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .github/ workflows .github/ workflows gitpulse gitpulse .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE PROMOTION.md PROMOTION.md README.md README.md ROADMAP.md ROADMAP.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI を利用した従来のコミット メッセージを端末から直接送信できます。
コミットメッセージに悩まされるのはもうやめましょう。 Git Pulse は段階的な変更を読み取り、お気に入りの AI モデルに質問し、すぐに使用できる完璧な従来のコミット メッセージを生成します。
🧠 AI を使用してコードの差分を理解し、意味のあるメッセージを作成します
📐 厳密な従来のコミット形式 ( feat(scope): description )
🔒 タイトルは 72 文字を超えないでください
🎮 インタラクティブ (確認/編集/中止) または完全自動モード
🪝 ワンコマンドフックインストール ( gitpulse init )
🔌 プラグイン可能なプロバイダー: OpenCode 、 OpenAI 、 Ollama (詳細は近日公開)
🐍 小さな依存関係: リクエストのみ – Python 3.8+
pip インストール gitpulse-commit
2. AI プロバイダーを構成する
デフォルトでは、Git Pulse は OpenCode (ローカルまたはリモート) を使用します。環境変数を設定します。
エクスポート OPENCODE_API_URL= " http://localhost:8080/v1/chat/completions "
import OPENCODE_API_KEY= " あなたのキーをここに "
OpenAI をご希望の場合:
エクスポート OPENAI_API_KEY= " sk-... "
あるいはオラマ：
import OLLAMA_API_URL= " http://localhost:11434/api/generate "
3. いくつかの変更を加えてステージングする
git add 。
4. Git Pulse にコミット メッセージを生成させます
gitパルス
feat(auth): add JWT token validation のような生成されたメッセージが表示されます。
確認、編集、中止 - とても簡単です。
完全に自動化されたパイプラインの場合:
gitpulse [オプション] # コミットメッセージを生成する
gitpulseinit #prepare-commit-msgフックをインストールする
git パルスのオプション
旗
説明
-

-自動
確認なしで自動的にコミットする
--プロバイダー
AI プロバイダー: opencode (デフォルト)、openai 、ollam
--モデル
モデル名 (例: gpt-4o-mini )
--api-url
API ベース URL をオーバーライドする
--APIキー
APIキーを上書きする
--出力
メッセージをファイルに書き込みます (フックによって内部的に使用されます)
Gitフック
gitpulseinit を実行した後、 git commit を実行するたびに、フックは次のことを行います。
従来のコミットメッセージを自動生成
メッセージを事前に入力してエディタを開きます (または、--auto が設定されている場合は直接コミットします)。
プロバイダー
デフォルトのモデル
環境変数
オープンコード
オープンコード
OPENCODE_API_URL 、OPENCODE_API_KEY
OpenAI
gpt-4o-mini
OPENAI_API_KEY
オラマ
ラマ3
OLLAMA_API_URL
実行時に --provider 、 --model 、 --api-url 、および --api-key を渡してデフォルトをオーバーライドできます。
git diff --cached 出力を抽出します。
慎重に設計されたシステム プロンプトとともに AI エンドポイントに送信します
応答を解析し、書式設定を削除し、72 文字に切り詰めます。
インタラクティブ モードの場合は、確認および編集できます。自動の場合はすぐにコミットします
システム プロンプトは、コミット メッセージのみを生成し、他には何も生成しないように設計されています。
リポジトリのクローンを作成し、編集可能モードでインストールします。
git clone https://github.com/erico964-blip/gitpulse
cd ギットパルス
pip install -e 。
テストを実行します (近日公開予定):
⭐ リポジトリにスターを付ける — Git Pulse によって数分間のコミット メッセージの苦痛から救われるのであれば、GitHub でスターを付けてください。
⚡ AI を利用した従来の Git 用コミット メッセージ ジェネレーター
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

⚡ AI-powered conventional commit message generator for Git - erico964-blip/gitpulse

GitHub - erico964-blip/gitpulse: ⚡ AI-powered conventional commit message generator for Git · GitHub
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
erico964-blip
/
gitpulse
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .github/ workflows .github/ workflows gitpulse gitpulse .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PROMOTION.md PROMOTION.md README.md README.md ROADMAP.md ROADMAP.md pyproject.toml pyproject.toml View all files Repository files navigation
AI-powered conventional commit messages, right from your terminal.
Stop struggling with commit messages. Git Pulse reads your staged changes, asks your favorite AI model, and generates a perfect Conventional Commit message – ready to use.
🧠 Uses AI to understand your code diff and craft meaningful messages
📐 Strict Conventional Commits format ( feat(scope): description )
🔒 Title never exceeds 72 characters
🎮 Interactive (confirm/edit/abort) or fully automatic mode
🪝 One-command hook installation ( git pulse init )
🔌 Pluggable providers: OpenCode , OpenAI , Ollama (more soon)
🐍 Tiny dependency: only requests – Python 3.8+
pip install gitpulse-commit
2. Configure your AI provider
By default, Git Pulse uses OpenCode (local or remote). Set the environment variables:
export OPENCODE_API_URL= " http://localhost:8080/v1/chat/completions "
export OPENCODE_API_KEY= " your-key-here "
If you prefer OpenAI :
export OPENAI_API_KEY= " sk-... "
Or Ollama :
export OLLAMA_API_URL= " http://localhost:11434/api/generate "
3. Make some changes and stage them
git add .
4. Let Git Pulse generate your commit message
git pulse
You'll see a generated message like feat(auth): add JWT token validation .
Confirm, edit, or abort – it's that simple.
For fully automated pipelines:
git pulse [options] # generate a commit message
git pulse init # install the prepare-commit-msg hook
Options for git pulse
Flag
Description
--auto
Automatically commit without confirmation
--provider
AI provider: opencode (default), openai , ollama
--model
Model name (e.g., gpt-4o-mini )
--api-url
Override API base URL
--api-key
Override API key
--output
Write message to a file (used internally by the hook)
Git Hook
After running git pulse init , every time you run git commit , the hook will:
Generate a conventional commit message automatically
Open your editor with the message pre-filled (or commit directly if --auto was set).
Provider
Default Model
Env Vars
OpenCode
opencode
OPENCODE_API_URL , OPENCODE_API_KEY
OpenAI
gpt-4o-mini
OPENAI_API_KEY
Ollama
llama3
OLLAMA_API_URL
You can pass --provider , --model , --api-url , and --api-key at runtime to override defaults.
Extracts the git diff --cached output
Sends it together with a carefully engineered system prompt to the AI endpoint
Parses the response, strips any formatting, truncates to 72 characters
If in interactive mode, lets you review and edit; if automatic, commits immediately
The system prompt is designed to produce only the commit message, nothing else.
Clone the repo and install in editable mode:
git clone https://github.com/erico964-blip/gitpulse
cd gitpulse
pip install -e .
Run tests (coming soon):
⭐ Star the repo — If Git Pulse saves you from a few minutes of commit-message anguish, give it a star on GitHub!
⚡ AI-powered conventional commit message generator for Git
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
