---
source: "https://github.com/helis-d/waylou"
hn_url: "https://news.ycombinator.com/item?id=48961587"
title: "Show HN: Waylou / a multi-provider CLI coding agent / fork of Gemini CLI"
article_title: "GitHub - helis-d/waylou: Waylou is a next-generation BYOK-enabled CLI based on Gemini CLI. · GitHub"
author: "Emirhan123"
captured_at: "2026-07-18T19:59:03Z"
capture_tool: "hn-digest"
hn_id: 48961587
score: 1
comments: 1
posted_at: "2026-07-18T19:45:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Waylou / a multi-provider CLI coding agent / fork of Gemini CLI

- HN: [48961587](https://news.ycombinator.com/item?id=48961587)
- Source: [github.com](https://github.com/helis-d/waylou)
- Score: 1
- Comments: 1
- Posted: 2026-07-18T19:45:35Z

## Translation

タイトル: 表示 HN: Waylou / マルチプロバイダー CLI コーディング エージェント / Gemini CLI のフォーク
記事タイトル: GitHub - helis-d/waylou: Waylou は、Gemini CLI をベースにした次世代の BYOK 対応 CLI です。 · GitHub
説明: Waylou は、Gemini CLI をベースにした次世代の BYOK 対応 CLI です。 - ヘリス-D/ウェイルー

記事本文:
GitHub - helis-d/waylou: Waylou は、Gemini CLI をベースにした次世代の BYOK 対応 CLI です。 · GitHub
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
ヘリス-D
/
ウェイルー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6,301 コミット 6,301 コミット .allstar .allstar .gemini .gemini .github .github .husky .husky .vscode .vscode docs docs evals evals 統合テスト 統合テスト メモリテスト メモリテスト パッケージ パッケージ パフォーマンステスト パフォーマンステスト スキーマ スキーマ スクリプト スクリプト sea sea third_party/get-ripgrep third_party/get-ripgrep tools/caretaker-agent/cloudrun tools/caretaker-agent/cloudrun .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .lycheeignore .lycheeignore .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json .waylouignore .waylouignore .yamllint.yml .yamllint.yml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile FAQ.md FAQ.md GEMINI.md GEMINI.md LICENSE LICENSE Makefile Makefile PROJECT_STRUCTURE.md PROJECT_STRUCTURE.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md esbuild.config.js esbuild.config.js eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたの端末。あなたのモデル。あなたのルール。
Waylou ACE CLI は、直接実行されるオープンソースの AI コーディング アシスタントです。
あなたの端末で。複数の AI モデル (Gemini、OpenAI、
Anthropic、Ollama、DeepSeek — 単一の統合インターフェースを介して。
Google Gemini CLI から分岐した Waylou は、基盤を拡張します。
マルチプロバイダーのサポート、自律的な運用、および設計されたアーキテクチャ
拡張性のために。
ステータス: 開発が活発です。コア機能が動作します。プロバイダー層
CLI 機能が注目されています

ヴィ開発。貢献は大歓迎です。
能力
Waylou ACE CLI
ジェミニ CLI
マルチプロバイダー (Gemini、OpenAI、Claude、ローカル)
✅
❌
BYOK — 自分のキーを持参する
✅
❌
プロバイダーのオーケストレーションとチーム ルーティング
✅
❌
ヘッドレス/自律モード
✅
✅
MCP (モデル コンテキスト プロトコル)
✅
✅
ACE準拠仕様
✅
❌
VS コードコンパニオン
✅
❌
Apache 2.0 オープンソース
✅
✅
クイックスタート
# グローバルにインストール
npm install -g @waylou/cli
# または即座に実行
npx @waylou/cli
# 起動
ウェイルー
特長
任意のモデル、任意のプロバイダーを使用します。セッション中に交換します。タスクを右側にルーティングする
機能、速度、コストに基づいてモデルを決定します。
Google Gemini — Gemini 2.5、3.0 (1M トークン コンテキスト)
人間 — クロード 3.5、クロード 4
Ollama — ローカル モデル (Llama、Mistral、Gemma)
アーキテクト、コーダー、レビューアーなど、さまざまな役割を持つモデルのチームを形成します。
オーケストレーターは、構成可能な戦略を使用してタスクを分散します。
(能力最高、ラウンドロビン、並列投票)。
エージェントコーディング環境の仕様によりエージェントが標準化されます。
3 つのコンプライアンス レベルにわたる行動:
CORE — ファイル操作、シェルコマンド、Web フェッチ
STANDARD — MCP 統合、チェックポイント、サンドボックス、メモリ
ENTERPRISE — ポリシー エンジン、監査ログ、チーム オーケストレーション
CI/CD パイプライン、スクリプト、自動化されたワークフローで実行します。
waylou -p " この PR を確認し、変更点を要約します " --output-format json
コンテキストエンジン
ベクトル埋め込み (SQLite-vec) と AST を組み合わせたセマンティック コード検索
分析 (ツリーシッター) により、正確で関連性のあるコンテキストを取得します。
macOS Seatbelt、Docker、または Podman を介した安全な実行。
npm install -g @waylou/cli
ソースから
git clone https://github.com/helis-d/waylou.git
CD ウェイルー
npmインストール
npm ビルドを実行する
ノード run.js
構成
import GEMINI_API_KEY= " your-key " # Google Gemini
import OPENAI_API_KEY= " your-key " # OpenA

私は
import ANTHROPIC_API_KEY= " your-key " # Anthropic
import DEEPSEEK_API_KEY= " your-key " # DeepSeek
ローカル モデルの場合は、Ollama をインストールして実行します。
waylou --プロバイダー オラマ
建築
CLI (インク/反応 UI)
│
▼
コア エンジン (エージェント ループ、ツール、プロンプト)
│
▼
プロバイダー層 (Gemini | OpenAI | Anthropic | Ollama | DeepSeek)
│
▼
AI API
完全なシステム設計については、ARCHITECTURE.md を参照してください。
パッケージ
説明
@waylou/cli
ターミナル UI および CLI バイナリ
@waylou/cli-core
コアエージェントロジック、ツール、プロンプト
@waylou/cliプロバイダー
マルチ AI プロバイダーの抽象化
@waylou/エーススペック
ACEの仕様と検証
@waylou/自律エンジン
ヘッドレス/自律動作
@waylou/コンテキストエンジン
ベクトルおよび AST セマンティック検索
@waylou/サンドボックス
安全な実行環境
@waylou/cli-sdk
拡張機能用のパブリック SDK
ドキュメント
計画されている機能と優先順位については、ROADMAP.md を参照してください。
貢献は大歓迎です — 特にプロバイダー層については
プロジェクトの最大の開発ニーズ。 AI API の経験がある場合は、
SDK、またはモデル統合については、ぜひご協力をお願いいたします。
完全なガイドについては、CONTRIBUTING.md を参照してください。
このプロジェクトはからフォークされています
Google Gemini CLI 。私たちは
オリジナルの作者と貢献者に感謝します。
例外的な基礎。
Apache ライセンス 2.0 — 詳細については、「ライセンス」を参照してください。
Waylou は、Gemini CLI をベースにした次世代の BYOK 対応 CLI です。
Apache-2.0ライセンス
行動規範
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

Waylou is a next-generation BYOK-enabled CLI based on Gemini CLI. - helis-d/waylou

GitHub - helis-d/waylou: Waylou is a next-generation BYOK-enabled CLI based on Gemini CLI. · GitHub
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
helis-d
/
waylou
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6,301 Commits 6,301 Commits .allstar .allstar .gemini .gemini .github .github .husky .husky .vscode .vscode docs docs evals evals integration-tests integration-tests memory-tests memory-tests packages packages perf-tests perf-tests schemas schemas scripts scripts sea sea third_party/ get-ripgrep third_party/ get-ripgrep tools/ caretaker-agent/ cloudrun tools/ caretaker-agent/ cloudrun .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .lycheeignore .lycheeignore .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json .waylouignore .waylouignore .yamllint.yml .yamllint.yml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile FAQ.md FAQ.md GEMINI.md GEMINI.md LICENSE LICENSE Makefile Makefile PROJECT_STRUCTURE.md PROJECT_STRUCTURE.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md esbuild.config.js esbuild.config.js eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Your terminal. Your models. Your rules.
Waylou ACE CLI is an open-source AI coding assistant that runs directly
in your terminal. It connects to multiple AI models — Gemini, OpenAI,
Anthropic, Ollama, DeepSeek — through a single, unified interface.
Forked from Google Gemini CLI, Waylou extends the foundation with
multi-provider support, autonomous operation, and an architecture designed
for extensibility.
Status: Active development. Core functionality works. Provider layer
and CLI features are under heavy development. Contributions welcome.
Capability
Waylou ACE CLI
Gemini CLI
Multi-provider (Gemini, OpenAI, Claude, local)
✅
❌
BYOK — Bring Your Own Key
✅
❌
Provider orchestration & team routing
✅
❌
Headless / autonomous mode
✅
✅
MCP (Model Context Protocol)
✅
✅
ACE compliance specification
✅
❌
VS Code companion
✅
❌
Apache 2.0 open source
✅
✅
Quick Start
# Install globally
npm install -g @waylou/cli
# Or run instantly
npx @waylou/cli
# Launch
waylou
Features
Use any model, any provider. Swap mid-session. Route tasks to the right
model based on capability, speed, or cost.
Google Gemini — Gemini 2.5, 3.0 (1M token context)
Anthropic — Claude 3.5, Claude 4
Ollama — Local models (Llama, Mistral, Gemma)
Form teams of models with different roles — architect, coder, reviewer.
The orchestrator distributes tasks using configurable strategies
(capability-best, round-robin, parallel-vote).
The Agentic Coding Environment specification standardizes agent
behavior across three compliance levels:
CORE — File operations, shell commands, web fetch
STANDARD — MCP integration, checkpointing, sandbox, memory
ENTERPRISE — Policy engine, audit logging, team orchestration
Run in CI/CD pipelines, scripts, and automated workflows:
waylou -p " Review this PR and summarize the changes " --output-format json
Context Engine
Semantic code search combining vector embeddings (SQLite-vec) with AST
analysis (Tree-sitter) for precise, relevant context retrieval.
Secure execution via macOS Seatbelt, Docker, or Podman.
npm install -g @waylou/cli
From Source
git clone https://github.com/helis-d/waylou.git
cd waylou
npm install
npm run build
node run.js
Configuration
export GEMINI_API_KEY= " your-key " # Google Gemini
export OPENAI_API_KEY= " your-key " # OpenAI
export ANTHROPIC_API_KEY= " your-key " # Anthropic
export DEEPSEEK_API_KEY= " your-key " # DeepSeek
For local models, install Ollama and run:
waylou --provider ollama
Architecture
CLI (Ink/React UI)
│
▼
Core Engine (Agent Loop, Tools, Prompts)
│
▼
Provider Layer (Gemini | OpenAI | Anthropic | Ollama | DeepSeek)
│
▼
AI APIs
See ARCHITECTURE.md for the full system design.
Package
Description
@waylou/cli
Terminal UI & CLI binary
@waylou/cli-core
Core agent logic, tools, prompts
@waylou/cli-provider
Multi-AI provider abstraction
@waylou/ace-spec
ACE specification & validation
@waylou/autonomous-engine
Headless / autonomous operation
@waylou/context-engine
Vector & AST semantic search
@waylou/sandbox
Secure execution environment
@waylou/cli-sdk
Public SDK for extensions
Documentation
See ROADMAP.md for planned features and priorities.
Contributions are welcome — especially for the provider layer , which is
the project's biggest development need. If you have experience with AI APIs,
SDKs, or model integrations, we'd love your help.
See CONTRIBUTING.md for the full guide.
This project is forked from
Google Gemini CLI . We are
grateful to the original authors and contributors for building an
exceptional foundation.
Apache License 2.0 — See LICENSE for details.
Waylou is a next-generation BYOK-enabled CLI based on Gemini CLI.
Apache-2.0 license
Code of conduct
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
