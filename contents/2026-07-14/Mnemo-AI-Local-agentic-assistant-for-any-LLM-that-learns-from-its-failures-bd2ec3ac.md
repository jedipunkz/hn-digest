---
source: "https://github.com/brunopistone/mnemoai"
hn_url: "https://news.ycombinator.com/item?id=48906032"
title: "Mnemo AI – Local agentic assistant for any LLM that learns from its failures"
article_title: "GitHub - brunopistone/mnemoai · GitHub"
author: "br1pistone"
captured_at: "2026-07-14T13:08:21Z"
capture_tool: "hn-digest"
hn_id: 48906032
score: 2
comments: 0
posted_at: "2026-07-14T12:49:54Z"
tags:
  - hacker-news
  - translated
---

# Mnemo AI – Local agentic assistant for any LLM that learns from its failures

- HN: [48906032](https://news.ycombinator.com/item?id=48906032)
- Source: [github.com](https://github.com/brunopistone/mnemoai)
- Score: 2
- Comments: 0
- Posted: 2026-07-14T12:49:54Z

## Translation

タイトル: Mnemo AI – 失敗から学習するあらゆる LLM のローカル エージェント アシスタント
記事タイトル: GitHub - brunopistone/mnemoai · GitHub
説明: GitHub でアカウントを作成して、brunopistone/mnemoai の開発に貢献します。

記事本文:
GitHub - ブルーノピストン/ニーモアイ · GitHub
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
ブルーノピストン
/
ニーモアイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダー

dファイル
308 コミット 308 コミット .github/ workflows .github/ workflows bash bash docs docs 画像 画像 src/ mnemoai src/ mnemoai テスト テスト .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSEライセンス README.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml pytest.ini pytest.ini 要件-dev.txt 要件-dev.txt 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
A local agentic AI assistant with MCP (Model Context Protocol) integration, RAG capabilities, and intelligent conversation management. Built on LangGraph with LangChain for multi-provider LLM support (Ollama, Amazon Bedrock, OpenAI, Anthropic, Amazon SageMaker AI, LiteLLM).
Full documentation is available at https://brunopistone.github.io/mnemoai/
pip install mnemoai-assistant # or: uv tool install mnemoai-assistant
mnemoai # 冗長 (思考を示す); -- 非表示にする詳細はありません
最初の実行時に構成が見つからない場合は、対話型コンフィギュレーターが起動し、プロバイダー、モデル、機能の切り替えを選択する手順が示され、 ~/.mnemoai/config/config.yaml が書き込まれます。
→ See the Getting Started guide for full setup.
🤖 Multi-Model Support : Ollama (local), Amazon Bedrock, OpenAI, Anthropic (Claude), Amazon SageMaker AI, LiteLLM (100+ providers)
🔧 MCP Tool System : Extensible tool architecture via Model Context Protocol
📚 RAG : Automatic document indexing and semantic (hybrid) search
🧠 User Profile Learning : Personalized responses learned from interactions
🧩 Episodic Memory : Learns from successful task completions and retrieves similar solutions
📖 ACE Playbook : Learns strategies from successes AND failures (Agentic Context Engineering)
🔍 Web Search & 🌐 Crawler : Brave Search API + web page e

RAG摂取による抽出
🖼️ ビジョンサポート : ビジョンモデルによる画像分析
📁 ファイル操作 & ✏️ 正確な編集 : テキスト、CSV、JSON、PDF、DOCX の読み取り/書き込み/編集
🔎 高速検索: Glob + ripgrep コンテンツ検索 (10 ～ 100 倍高速)
📋 Todo 追跡、📝 計画モード & 🔄 バックグラウンド タスク : マルチステップのタスク管理
⚡ Bash 実行 & 🛡️ Git セーフティ: スマートなエラー処理とガードレールを備えたシェル コマンド
MIT ライセンスに基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
これは個人開発プロジェクトです。自由にフォークしてニーズに合わせて調整してください。元のリポジトリへの帰属は歓迎されますが、必須ではありません。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to brunopistone/mnemoai development by creating an account on GitHub.

GitHub - brunopistone/mnemoai · GitHub
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
brunopistone
/
mnemoai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
308 Commits 308 Commits .github/ workflows .github/ workflows bash bash docs docs images images src/ mnemoai src/ mnemoai tests tests .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml pytest.ini pytest.ini requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
A local agentic AI assistant with MCP (Model Context Protocol) integration, RAG capabilities, and intelligent conversation management. Built on LangGraph with LangChain for multi-provider LLM support (Ollama, Amazon Bedrock, OpenAI, Anthropic, Amazon SageMaker AI, LiteLLM).
Full documentation is available at https://brunopistone.github.io/mnemoai/
pip install mnemoai-assistant # or: uv tool install mnemoai-assistant
mnemoai # verbose (shows thinking); --no-verbose to hide
On first run, if no config is found, an interactive configurator launches and walks you through picking a provider, model, and feature toggles — then writes ~/.mnemoai/config/config.yaml .
→ See the Getting Started guide for full setup.
🤖 Multi-Model Support : Ollama (local), Amazon Bedrock, OpenAI, Anthropic (Claude), Amazon SageMaker AI, LiteLLM (100+ providers)
🔧 MCP Tool System : Extensible tool architecture via Model Context Protocol
📚 RAG : Automatic document indexing and semantic (hybrid) search
🧠 User Profile Learning : Personalized responses learned from interactions
🧩 Episodic Memory : Learns from successful task completions and retrieves similar solutions
📖 ACE Playbook : Learns strategies from successes AND failures (Agentic Context Engineering)
🔍 Web Search & 🌐 Crawler : Brave Search API + web page extraction with RAG ingestion
🖼️ Vision Support : Image analysis with vision models
📁 File Operations & ✏️ Precise Editing : Read/write/edit text, CSV, JSON, PDF, DOCX
🔎 Fast Search : Glob + ripgrep content search (10-100x faster)
📋 Todo Tracking, 📝 Plan Mode & 🔄 Background Tasks : Multi-step task management
⚡ Bash Execution & 🛡️ Git Safety : Shell commands with smart error handling and guardrails
Licensed under the MIT License — see the LICENSE file for details.
This is a personal development project. Feel free to fork and adapt it to your needs; attribution to the original repository is appreciated but not required.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
