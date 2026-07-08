---
source: "https://github.com/animesh-94/Onboard-CLI"
hn_url: "https://news.ycombinator.com/item?id=48836813"
title: "Show HN: Onboard-CLI, a LLM powered and AST-based tool to visualize codebase"
article_title: "GitHub - animesh-94/Onboard-CLI: An AST-powered, local-first CLI that visualizes complex system architectures and enforces architectural boundaries via instant Git hooks. · GitHub"
author: "yr_animesh"
captured_at: "2026-07-08T20:57:16Z"
capture_tool: "hn-digest"
hn_id: 48836813
score: 3
comments: 0
posted_at: "2026-07-08T20:09:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Onboard-CLI, a LLM powered and AST-based tool to visualize codebase

- HN: [48836813](https://news.ycombinator.com/item?id=48836813)
- Source: [github.com](https://github.com/animesh-94/Onboard-CLI)
- Score: 3
- Comments: 0
- Posted: 2026-07-08T20:09:03Z

## Translation

タイトル: HN を表示: Onboard-CLI、コードベースを視覚化するための LLM を利用した AST ベースのツール
記事のタイトル: GitHub - animesh-94/Onboard-CLI: 複雑なシステム アーキテクチャを視覚化し、インスタント Git フックを介してアーキテクチャの境界を強制する、AST ベースのローカル ファースト CLI。 · GitHub
説明: 複雑なシステム アーキテクチャを視覚化し、インスタント Git フックを介してアーキテクチャの境界を強制する、AST を利用したローカル ファーストの CLI。 - animesh-94/オンボード-CLI

記事本文:
GitHub - animesh-94/Onboard-CLI: 複雑なシステム アーキテクチャを視覚化し、インスタント Git フックを介してアーキテクチャの境界を強制する、AST を利用したローカル ファーストの CLI。 · GitHub
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
アニメッシュ-94
/
オンボード CLI
公共
通知
あなたはきっとsiでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
44 コミット 44 コミット .github/ workflows .github/ workflows cmd cmd docs docs 内部 内部 npm-wrapper npm-wrapper ui ui .gitignore .gitignore AI_USAGE.md AI_USAGE.md COTRIBUTING.md COTRIBUTING.md README.md README.md go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh main.go main.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
無題.デザイン.1.1.mp4
オンボード CLI
コード解析、システム プロファイリング、キャンバス ベースのノード視覚化のための開発者プラットフォーム。
Onboard-CLI は、開発者が大規模で複雑なコードベースを迅速に理解し、計画を立てるのに役立つように設計された、高度なコマンドライン インターフェイス ツールおよび Web ベースのビジュアライザーです。 Tree-sitter を介した AST (抽象構文ツリー) 解析を活用することで、Onboard-CLI は構造トポロジー グラフを生成し、アーキテクチャ境界を強制し、直感的な React Flow キャンバスを通じてそれらを表示します。
AST スライシング エンジン: ツリー シッターを使用した詳細なコード解析により、複数の言語 (Go、JS、TS、Python、Java) にわたって正確な構造ノードを生成します。
Interactive Visualizer (map) : ローカル React Flow キャンバス ( http://localhost:3000/app ) を自動的に起動して、指定された半径内のコード パス、依存関係、およびトポロジー マップを視覚的に探索します。
アーキテクチャ ドリフト検出 (drift) : アーキテクチャ .yml ルールに照らしてコードベースを分析し、不正なファイル間インポートや境界違反を検出し、長期的なコードの健全性を確保します。
包括的なエコシステム : 構成管理、影響分析、コードのエクスポート、所有者の追跡、プロジェクト パルスのための組み込みコマンド。
最新の技術スタック : Go で書かれた超高速 CLI と、Vite、React 19、@ を利用した豊富なフロントエンドとの組み合わせ

xyflow/react (リアクトフロー)、Framer Motion、Tailwind CSS。
言語サポート : 5 つ以上の組み込みパーサー (Go、TypeScript、JavaScript、Python、Java)。
Visualization Radius : 構成可能なディープマッピング (デフォルトの半径: 1、複雑な依存関係ツリーにスケール)。
ルールの適用 : 数千のファイルにわたるアーキテクチャ上のドリフトに対する 1 秒未満の境界正規表現評価。
フロントエンド パフォーマンス : Vite および SWC/Oxc lint を利用して、大規模なノード セット向けに最適化されたキャンバス レンダリング。
提供されているインストール スクリプトを使用して Onboard-CLI をインストールできます。
カール -fsSL https://raw.githubusercontent.com/onboard-cli/install.sh |バッシュ
(または、リポジトリから直接 ./install.sh を実行します)。
Invoke-WebRequest - Uri https://raw.githubusercontent.com / onboard - cli / install.ps1 - OutFile install.ps1; .\install.ps1
💻使用法
必要な .onboard を生成します。
オンボード初期化
CI/CD フォークの場合は、以下のコマンドを使用して、必要な構成とアーキテクチャ.yaml ファイルを生成します。
onboard init --template クリーン アーキテクチャ
(利用可能なテンプレート: generic 、 clean-architecture 、 modular-monolith 、 mvc 、 serverless )
2. コードベースのマッピング (ビジュアライザー)
コンテキスト エンジンをトリガーしてシンボルまたはファイル パスをマップし、ビジュアライザー サーバーを起動します。
オンボード マップ --target " 内部/パーサー " --radius 2
👉 提供されたリンク ( http://localhost:3000/app ) をクリックして、インタラクティブなキャンバスを表示します。 Fuzzy Finder を使用するには、UI で Ctrl+P を押してください。右上からダークモードとコンパクトモードを切り替えます。
バックエンド API ルートを、さまざまなフレームワーク (Express、Gin、FastAPI、Spring) にわたる正確なファイルおよびライン ハンドラーの場所に自動的にマッピングします。
オンボードルート --プロトコルレスト --フレームワークエクスプレス
4. アーキテクチャのドリフトを検出する
定義したアーキテクチャ ルールに対する違反をチェックします。
オンボード ドリフト --rules アーキテクチャ.yml
5. CI/CDの統合
インテグ

GitHub Actions ワークフローにオンボード ドリフトを追加して、すべてのプル リクエストにアーキテクチャ上の境界を強制します。 docs/onboard-action.yml でテンプレートを確認してください。
オンボードの影響 : 提案されたコード変更の影響を分析します。
オンボード所有者 : コードの所有者と保守者を追跡します。
onboard export : グラフィカル/AST データをエクスポートします。
onboardpulse : コードベースの健全性と最近のアクティビティの簡単な概要を取得します。
(注: --no-update-check グローバル フラグを任意のコマンドに渡して、自動非同期バージョン チェッカーを無効にすることができます)。
CLI エンジン: Go、Cobra、Go-Tree-Sitter、YAML。
Web UI : React 19、TypeScript、Vite、@xyflow/react (React Flow)、Tailwind CSS、Framer Motion、Lucide React。
コードのハイライトとマークダウン：Shiki、MDX、Rehype。
寄付を歓迎します! CLI ロジックの cmd ディレクトリと内部ディレクトリ、および React フロントエンドの ui ディレクトリを確認してください。
プル リクエストを送信する前に、貢献ガイドラインと AI 使用ポリシーをお読みください。
機能ブランチを作成します ( git checkout -b feature/amazing-feature )
変更をコミットします ( git commit -m 'Add someすばらしい機能' )
ブランチにプッシュする ( git Push Origin feature/amazing-feature )
MIT ライセンスに基づいて配布されます。詳細については、「ライセンス」を参照してください。
複雑なシステム アーキテクチャを視覚化し、インスタント Git フックを介してアーキテクチャの境界を強制する、AST を利用したローカル ファーストの CLI。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
9
v1.3.0
最新の
2026 年 7 月 1 日
+ 8 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An AST-powered, local-first CLI that visualizes complex system architectures and enforces architectural boundaries via instant Git hooks. - animesh-94/Onboard-CLI

GitHub - animesh-94/Onboard-CLI: An AST-powered, local-first CLI that visualizes complex system architectures and enforces architectural boundaries via instant Git hooks. · GitHub
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
animesh-94
/
Onboard-CLI
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
44 Commits 44 Commits .github/ workflows .github/ workflows cmd cmd docs docs internal internal npm-wrapper npm-wrapper ui ui .gitignore .gitignore AI_USAGE.md AI_USAGE.md CONTRIBUTING.md CONTRIBUTING.md README.md README.md go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh main.go main.go View all files Repository files navigation
Untitled.design.1.1.mp4
Onboard-CLI
Developer platform for code parsing, systems profiling, and canvas-based node visualization.
Onboard-CLI is an advanced command-line interface tool and web-based visualizer designed to help developers quickly understand and map out large, complex codebases. By leveraging AST (Abstract Syntax Tree) parsing via Tree-sitter, Onboard-CLI generates structural topology graphs and enforces architectural boundaries, presenting them through an intuitive React Flow canvas.
AST Slicing Engine : Deep code parsing using tree-sitter for accurate structural node generation across multiple languages (Go, JS, TS, Python, Java).
Interactive Visualizer ( map ) : Automatically spins up a local React Flow canvas ( http://localhost:3000/app ) to visually explore code paths, dependencies, and topological maps within a specified radius.
Architecture Drift Detection ( drift ) : Analyzes codebase against architecture.yml rules to detect unauthorized cross-file imports and boundary violations, ensuring long-term code health.
Comprehensive Ecosystem : Built-in commands for config management, impact analysis, code exporting, owner tracking, and project pulse.
Modern Tech Stack : Blazing fast CLI written in Go, paired with a rich frontend utilizing Vite, React 19, @xyflow/react (React Flow), Framer Motion, and Tailwind CSS.
Language Support : 5+ Built-in parsers (Go, TypeScript, JavaScript, Python, Java).
Visualization Radius : Configurable deep-mapping (default radius: 1, scales to complex dependency trees).
Rule Enforcement : Sub-second boundary regex evaluation for architectural drift across thousands of files.
Frontend Performance : Optimized canvas rendering for large node sets, powered by Vite & SWC/Oxc linting.
You can install Onboard-CLI using the provided installation scripts:
curl -fsSL https://raw.githubusercontent.com/onboard-cli/install.sh | bash
(Or run ./install.sh directly from the repository).
Invoke-WebRequest - Uri https: // raw.githubusercontent.com / onboard - cli / install.ps1 - OutFile install.ps1; .\install.ps1
💻 Usage
Generate the necessary .onboard .
onboard init
For CI/CD fokes use below command to generate the necessary configurations and architecture.yaml file
onboard init --template clean-architecture
(Available templates: generic , clean-architecture , modular-monolith , mvc , serverless )
2. Map the Codebase (Visualizer)
Trigger the context engine to map a symbol or file path, and boot up the visualizer server.
onboard map --target " internal/parser " --radius 2
👉 Click the provided link ( http://localhost:3000/app ) to view the interactive canvas. Press Ctrl+P in the UI to use the Fuzzy Finder! Toggle Dark Mode and Compact Mode from the top right.
Automatically map backend API routes to their exact file and line handler locations across various frameworks (Express, Gin, FastAPI, Spring).
onboard routes --protocol rest --framework express
4. Detect Architectural Drift
Check for violations against your defined architecture rules.
onboard drift --rules architecture.yml
5. CI/CD Integration
Integrate onboard drift into your GitHub Actions workflow to enforce architectural boundaries on every Pull Request. Check out the template at docs/onboard-action.yml .
onboard impact : Analyze the impact of a proposed code change.
onboard owners : Track code ownership and maintainers.
onboard export : Export graphical/AST data.
onboard pulse : Get a quick summary of codebase health and recent activity.
(Note: You can pass the --no-update-check global flag to any command to disable the automatic async version checker).
CLI Engine : Go, Cobra, Go-Tree-Sitter, YAML.
Web UI : React 19, TypeScript, Vite, @xyflow/react (React Flow), Tailwind CSS, Framer Motion, Lucide React.
Code Highlighting & Markdown : Shiki, MDX, Rehype.
We welcome contributions! Please check out the cmd and internal directories for CLI logic, and the ui directory for the React frontend.
Please read our Contributing Guidelines and AI Usage Policy before submitting a Pull Request.
Create your feature branch ( git checkout -b feature/amazing-feature )
Commit your changes ( git commit -m 'Add some amazing feature' )
Push to the branch ( git push origin feature/amazing-feature )
Distributed under the MIT License. See LICENSE for more information.
An AST-powered, local-first CLI that visualizes complex system architectures and enforces architectural boundaries via instant Git hooks.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
9
v1.3.0
Latest
Jul 1, 2026
+ 8 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
