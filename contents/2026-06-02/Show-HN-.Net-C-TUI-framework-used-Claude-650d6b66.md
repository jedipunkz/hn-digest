---
source: "https://github.com/rivoli-ai/andy-tui2"
hn_url: "https://news.ycombinator.com/item?id=48359909"
title: "Show HN: .Net/C# TUI framework (used Claude)"
article_title: "GitHub - rivoli-ai/andy-tui2: Terminal library · GitHub"
author: "M4R5H4LL"
captured_at: "2026-06-02T00:35:29Z"
capture_tool: "hn-digest"
hn_id: 48359909
score: 4
comments: 0
posted_at: "2026-06-01T17:26:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: .Net/C# TUI framework (used Claude)

- HN: [48359909](https://news.ycombinator.com/item?id=48359909)
- Source: [github.com](https://github.com/rivoli-ai/andy-tui2)
- Score: 4
- Comments: 0
- Posted: 2026-06-01T17:26:07Z

## Translation

タイトル: HN を表示: .Net/C# TUI フレームワーク (クロードを使用)
記事タイトル: GitHub - rivoli-ai/andy-tui2: ターミナル ライブラリ · GitHub
説明: 端末ライブラリ。 GitHub でアカウントを作成して、rivoli-ai/andy-tui2 の開発に貢献してください。
HN テキスト: HN クライアントの例 (#64)。このライブラリは、コーディング アシスタントを構築するために使用されます。そこに到達するためのすべての作業を無視して、100% バイブコード化されています。

記事本文:
GitHub - rivoli-ai/andy-tui2: ターミナル ライブラリ · GitHub
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
リボリアイ
/
アンディ・トゥイ2
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチタグ

s ファイルコードに移動 その他のアクションメニューを開く フォルダーとファイル
41 コミット 41 コミット .github/ workflows .github/ workflows 資産 アセット ドキュメント ドキュメント サンプル サンプル サンプル src src テスト テスト .gitignore .gitignore Andy.Tui.sln Andy.Tui.sln CLAUDE.md CLAUDE.md Directory.Build.props Directory.Build.props ライセンス ライセンス README.md README.md エージェント.md エージェント.md global.json global.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
宣言型コンポーネント、リアクティブ バインディング、実用的な CSS サブセット、統合されたレンダリング パイプライン、および最高クラスの可観測性を備えた、.NET 8 用の最新のリアクティブ TUI フレームワーク。高性能のターミナル アプリケーション、ダッシュボード、リアルタイム ログ ビューア向けに構築されています。
このソフトウェアは ALPHA 段階にあります。その機能、安定性、安全性についてはいかなる保証も行われません。
このライブラリはファイルとディレクトリに対して破壊操作を実行します。
権限管理は完全にテストされておらず、セキュリティ上の脆弱性がある可能性があります
運用環境では使用しないでください
重要なデータまたはかけがえのないデータを含むシステムでは使用しないでください
完全に検証されたバックアップのないシステムでは使用しないでください。
作成者は、データ損失、システム損傷、またはセキュリティ侵害に対して一切の責任を負いません。
Reactive Core : 信号、計算値、エフェクト、データ バインディング
コンポーネント システム : 修飾子を使用した宣言的なコンポーネント構成
CSS スタイリング : カスケード、特異性、変数、疑似クラスを含む CSS のサブセット
Flex Layout : 最新のフレックスボックスベースのレイアウト エンジン
リッチ テキスト : 書記素サポートによる Unicode 対応のテキスト レンダリング
ウィジェット ライブラリ : 80 以上の事前構築済みウィジェット (テーブル、フォーム、グラフなど)
統合パイプライン: Compose → Style → Layout → DisplayList → Compositor → Backend
ターミナル バックエンド: 完全な ANSI/エスケープ シーケンス サポート
可観測性 : 組み込みのロギング、トレース、パフォーマンス監視
テスト: 抑止力

信頼性の高い単体テストのためのミニスティックモード
src/ — ライブラリ プロジェクト (パック可能)
Andy.Tui (アンブレラ メタパッケージ)
testing/ — xUnit テスト プロジェクト (パック不可)
docs/ — 設計、ロードマップ、フェーズ、テスト、パフォーマンス計画
assets/ — ドキュメントと NuGet パッケージ化に使用されるアイコンと画像
dotnet パッケージ Andy.Tui --prerelease を追加
PackageReference経由
< PackageReference Include = " Andy.Tui " バージョン = " *-rc.* " />
注 : プレリリース パッケージは、メイン ブランチへのコミットごとに自動的に公開されます。
ANSI カラーをサポートする端末 (最新の端末)
テストを実行します。
dotnet test -c リリース
コードカバレッジ (オプション):
dotnet test --collect:"XPlat コード カバレッジ" --results-directory ./TestResults
reportgenerator -reports:"./TestResults/*/coverage.cobertura.xml" -targetdir:"./TestResults/CoverageReport" -reporttypes:Html
CI/CD パイプラインは、NuGet パッケージを自動的に公開します。
プレリリース バージョン (例: 2025.8.25-rc.30 ): メイン ブランチへのプッシュごとに公開されます
リリース バージョン (例: 1.0.0 ): v* に一致するタグの作成時に公開されます。
# すべてのライブラリをパックする
dotnet Pack -c Release -o ./nupkg
# NuGet に公開する (API キーが必要)
dotnet nuget プッシュ ./nupkg/Andy.Tui。 * .nupkg -k < NUGET_API_KEY > -s https://api.nuget.org/v3/index.json
開発ワークフロー
テストには src/ へのコード変更を伴う必要があります。
重要な変更をコミットする前にテストを実行します: dotnet test
意味のあるリファクタリング/機能のカバレッジを生成する
📖 はじめに - インストール、基本概念、例
🏗️ アーキテクチャ - システム設計とレンダリング パイプライン
🎨 ウィジェット カタログ - 80 以上の UI コンポーネントの完全なリスト
📚 ドキュメントインデックス - 完全なドキュメントの概要
✅ フェーズ 0: 基礎 - 完了
✅ フェーズ 1 : ビジュアルコア - 完了
✅ フェーズ 2: レンダリング コア - 完了
✅ フェーズ 3 : インタラクティブ性とアニマ

オプション - 完了
✅ フェーズ 4: 仮想化とウィジェット - 完了 (80 以上のウィジェットが実装)
🚧 フェーズ 5 : 追加のバックエンド - 計画中
✅ テスト スイートの実行 (CI でのパフォーマンス テストはスキップされます)
✅ NuGet パッケージを nuget.org に公開する
⚠️ 一部の Playwright テストでは、保留中の CI ブラウザーのセットアップが無効になりました
環境の変動により、CI でのパフォーマンス テストが失敗する場合がある
Playwright ブラウザ テストには手動でブラウザをインストールする必要があります
複雑なレイアウトでエッジケースをレンダリングする一部のウィジェット
貢献は大歓迎です！お願いします:
PR を送信する前にテストを実行する: dotnet test
既存のコード スタイルに従う (ドットネット形式を使用)
新しい機能のテストを追加する
必要に応じてドキュメントを更新する
Apache-2.0 ライセンス。詳細については、LICENSE ファイルを参照してください。
ディスカッション : GitHub ディスカッション
覚えておいてください: これは ALPHA ソフトウェアです。ご自身の責任で使用し、常にバックアップを維持してください。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

Terminal library. Contribute to rivoli-ai/andy-tui2 development by creating an account on GitHub.

With an example for an HN client (#64). The library is used to build a coding assistant. 100% vibe coded ignoring all the work to get there.

GitHub - rivoli-ai/andy-tui2: Terminal library · GitHub
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
rivoli-ai
/
andy-tui2
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
41 Commits 41 Commits .github/ workflows .github/ workflows assets assets docs docs examples examples src src tests tests .gitignore .gitignore Andy.Tui.sln Andy.Tui.sln CLAUDE.md CLAUDE.md Directory.Build.props Directory.Build.props LICENSE LICENSE README.md README.md agents.md agents.md global.json global.json View all files Repository files navigation
A modern, reactive TUI framework for .NET 8 with declarative components, reactive bindings, a pragmatic CSS subset, a unified rendering pipeline, and first-class observability. Built for high-performance terminal applications, dashboards, and real-time log viewers.
This software is in ALPHA stage. NO GUARANTEES are made about its functionality, stability, or safety.
This library performs DESTRUCTIVE OPERATIONS on files and directories
Permission management is NOT FULLY TESTED and may have security vulnerabilities
DO NOT USE in production environments
DO NOT USE on systems with critical or irreplaceable data
DO NOT USE on systems without complete, verified backups
The authors assume NO RESPONSIBILITY for data loss, system damage, or security breaches
Reactive Core : Signals, computed values, effects, and data bindings
Component System : Declarative component composition with modifiers
CSS Styling : Subset of CSS with cascade, specificity, variables, pseudo-classes
Flex Layout : Modern flexbox-based layout engine
Rich Text : Unicode-aware text rendering with grapheme support
Widget Library : 80+ pre-built widgets (tables, forms, charts, etc.)
Unified Pipeline : Compose → Style → Layout → DisplayList → Compositor → Backend
Terminal Backend : Full ANSI/escape sequence support
Observability : Built-in logging, tracing, performance monitoring
Testing : Deterministic mode for reliable unit tests
src/ — library projects (packable)
Andy.Tui (umbrella meta-package)
tests/ — xUnit test projects (non-packable)
docs/ — design, roadmap, phases, testing, perf plans
assets/ — icons and images used for documentation and NuGet packaging
dotnet add package Andy.Tui --prerelease
Via PackageReference
< PackageReference Include = " Andy.Tui " Version = " *-rc.* " />
Note : Pre-release packages are published automatically for every commit to main branch.
Terminal with ANSI color support (most modern terminals)
Run tests:
dotnet test -c Release
Code coverage (optional):
dotnet test --collect:"XPlat Code Coverage" --results-directory ./TestResults
reportgenerator -reports:"./TestResults/*/coverage.cobertura.xml" -targetdir:"./TestResults/CoverageReport" -reporttypes:Html
The CI/CD pipeline automatically publishes NuGet packages:
Pre-release versions (e.g., 2025.8.25-rc.30 ): Published on every push to main branch
Release versions (e.g., 1.0.0 ): Published when creating a tag matching v*
# Pack all libraries
dotnet pack -c Release -o ./nupkg
# Publish to NuGet (requires API key)
dotnet nuget push ./nupkg/Andy.Tui. * .nupkg -k < NUGET_API_KEY > -s https://api.nuget.org/v3/index.json
Development workflow
Tests must accompany code changes to src/
Run tests before committing significant changes: dotnet test
Generate coverage for meaningful refactors/features
📖 Getting Started - Installation, basic concepts, and examples
🏗️ Architecture - System design and rendering pipeline
🎨 Widget Catalog - Complete list of 80+ UI components
📚 Documentation Index - Full documentation overview
✅ Phase 0 : Foundations - Complete
✅ Phase 1 : Visual Core - Complete
✅ Phase 2 : Rendering Core - Complete
✅ Phase 3 : Interactivity & Animations - Complete
✅ Phase 4 : Virtualization & Widgets - Complete (80+ widgets implemented)
🚧 Phase 5 : Additional Backends - Planned
✅ Test suite runs (with performance tests skipped in CI)
✅ NuGet package publishing to nuget.org
⚠️ Some Playwright tests disabled pending CI browser setup
Performance tests may fail in CI due to environment variability
Playwright browser tests require manual browser installation
Some widget rendering edge cases in complex layouts
Contributions are welcome! Please:
Run tests before submitting PRs: dotnet test
Follow existing code style (use dotnet format )
Add tests for new functionality
Update documentation as needed
Apache-2.0 License. See LICENSE file for details.
Discussions : GitHub Discussions
Remember : This is ALPHA software. Use at your own risk and always maintain backups.
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
