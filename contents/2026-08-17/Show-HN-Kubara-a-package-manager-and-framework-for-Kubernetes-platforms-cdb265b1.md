---
source: "https://github.com/kubara-io/kubara"
hn_url: "https://news.ycombinator.com/item?id=49327257"
title: "Show HN: Kubara a package manager and framework for Kubernetes platforms"
article_title: "GitHub - kubara-io/kubara: kubara is a single binary CLI tool written in Go providing a lightweight framework for bootstrapping Kubernetes platforms with production-proven best practices. · GitHub"
image: "https://opengraph.githubassets.com/bab9b1e2f53900da2fc6ca83b5cdf45a4b0321041218780f9238c58c9b3c2d98/kubara-io/kubara"
author: "tuunit"
captured_at: "2026-08-17T07:44:05Z"
capture_tool: "hn-digest"
hn_id: 49327257
score: 2
comments: 0
posted_at: "2026-08-17T06:48:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Kubara a package manager and framework for Kubernetes platforms

- HN: [49327257](https://news.ycombinator.com/item?id=49327257)
- Source: [github.com](https://github.com/kubara-io/kubara)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T06:48:30Z

## Translation

タイトル: Show HN: Kubara が Kubernetes プラットフォーム用のパッケージ マネージャーとフレームワークを担当
記事のタイトル: GitHub - kubara-io/kubara: kubara は、Go で書かれた単一のバイナリ CLI ツールで、本番環境で実証済みのベスト プラクティスを使用して Kubernetes プラットフォームをブートストラップするための軽量フレームワークを提供します。 · GitHub
説明: kubara は、Go で書かれた単一のバイナリ CLI ツールで、本番環境で実証済みのベスト プラクティスを使用して Kubernetes プラットフォームをブートストラップするための軽量フレームワークを提供します。 - クバライオ/クバラ
HN テキスト: kubara はプラットフォーム アーキテクチャを再利用可能な製品に変えます。その目標は、チームがクラスター全体で一貫した Kubernetes プラットフォーム スタックをブートストラップ、パッケージ化、バージョン管理、運用できるように支援することです。

記事本文:
GitHub - kubara-io/kubara: kubara は、Go で書かれた単一のバイナリ CLI ツールで、運用環境で実証済みのベスト プラクティスを使用して Kubernetes プラットフォームをブートストラップするための軽量フレームワークを提供します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
クバライオ
/
クバラ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
407 コミット 407 コミット .github .github .pre-commi

t-config .pre-commit-config .scripts .scripts docs docs src src .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE.md NOTICE.md README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_HELM_DEPENDENCIES.md THIRD_PARTY_HELM_DEPENDENCIES.md install.sh install.sh renovate.json renovate.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
kubara は、GitOps ファーストのワークフローで Kubernetes プラットフォームをブートストラップして操作するための独自の CLI です。
これは、プラットフォームのスキャフォールディング、環境構成、実稼働対応のデフォルトを単一のバイナリに結合します。
1 つの CLI でプラットフォームのセットアップとライフサイクル タスクを実行
反復可能なデプロイメントのための GitOps ネイティブ構造
マルチクラスターおよびマルチテナント環境向けに構築
Terraform および Helm ベースのコンポーネントで拡張可能
Linux、macOS、および Windows のインストール手順については、INSTALLATION.md を参照してください。
パブリックドキュメント: https://docs.kubara.io
kubara は、ブートストラップ基盤とデフォルトのプラットフォーム スタックを OCI カタログから解決します。 kubara-io/catalogs and its releases でカタログのソースとバージョン付きリリースを見つけてください。クラスターごとのカタログ構成とカスタム カタログについては、「カタログ」を参照してください。
init 新しい kubara ディレクトリを初期化します
構成されたカタログ テンプレートから Helm および Terraform アーティファクトを生成します。
bootstrap 前提条件の CRD と Argo CD を指定したクラスターにブートストラップします。
schema 構成構造の JSON スキーマ ファイルを生成します
エージェント AI コーディング アシスタントのオンボーディング ファイル (AGENTS.md) をスキャフォールディングします。
カタログ プラットフォーム カタログの管理
クラスター kubara クラスター構成を管理する
help, h コマンドのリストまたは 1 つのコマンドのヘルプを表示します
グローバルオプション
--kubeconfig 文字列

g kubeconfig ファイルへのパス (デフォルト: "~/.kube/config")
--work-dir 文字列、-w 文字列 作業ディレクトリ (デフォルト: ".")
--config-file string、-c string 設定ファイルへのパス (デフォルト: "config.yaml")
--env-file string .env ファイルへのパス (デフォルト: ".env")
--test-connection Kubernetes クラスターに到達できるかどうかを確認します。名前空間をリストして終了します
--base64 Base64 エンコード/デコード モードを有効にする
--encode Base64 エンコード入力
--decode Base64 デコード入力
--string string Base64 操作の入力文字列
--file string Base64 操作の入力ファイル パス
--check-update 新しい kubara リリースをオンラインで確認します
--help、-h ヘルプを表示
--version、-v バージョンを出力します
アップデートチェック
kubara は実行のたびに新しい GitHub リリースをチェックします。 KUBARA_UPDATE_CHECK=0 で無効にします。ライブチェックのために kubara --check-update を実行します。
#kubara Slack チャネルに参加して、kubara の他のユーザーとチャットしたり、メンテナーに直接連絡したりできます。パブリック招待リンク ( http://slack.k8s.io/ ) を使用して、公式の Kubernetes Slack スペースへの招待を取得します。
質問とバグレポート: GitHub の問題
ディスカッションと Q&A: GitHub ディスカッション
チームと貢献者のガイダンス: COTRIBUTING.md
行動規範: CODE_OF_CONDUCT.md
貢献は大歓迎です。
プル リクエストを開く前に CONTRIBUTING.md をお読みください。
ソフトウェアソースコード：Apache 2.0
kubara は、Go で書かれた単一のバイナリ CLI ツールで、運用環境で実証済みのベスト プラクティスを使用して Kubernetes プラットフォームをブートストラップするための軽量フレームワークを提供します。
Readme Apache-2.0 ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
28 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

kubara is a single binary CLI tool written in Go providing a lightweight framework for bootstrapping Kubernetes platforms with production-proven best practices. - kubara-io/kubara

kubara turns platform architecture into a reusable product. Its goal is to help teams bootstrap, package, version, and operate consistent Kubernetes platform stacks across clusters.

GitHub - kubara-io/kubara: kubara is a single binary CLI tool written in Go providing a lightweight framework for bootstrapping Kubernetes platforms with production-proven best practices. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
kubara-io
/
kubara
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
407 Commits 407 Commits .github .github .pre-commit-config .pre-commit-config .scripts .scripts docs docs src src .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE.md NOTICE.md README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_HELM_DEPENDENCIES.md THIRD_PARTY_HELM_DEPENDENCIES.md install.sh install.sh renovate.json renovate.json View all files Repository files navigation
kubara is an opinionated CLI to bootstrap and operate Kubernetes platforms with a GitOps-first workflow.
It combines platform scaffolding, environment configuration, and production-ready defaults in a single binary.
One CLI for platform setup and lifecycle tasks
GitOps-native structure for repeatable deployments
Built for multi-cluster and multi-tenant environments
Extensible with Terraform and Helm based components
See INSTALLATION.md for Linux, macOS, and Windows installation instructions.
Public docs: https://docs.kubara.io
kubara resolves its bootstrap foundation and default platform stack from OCI catalogs. Find the catalog source and versioned releases at kubara-io/catalogs and its releases . See Catalogs for per-cluster catalog configuration and custom catalogs.
init Initialize a new kubara directory
generate Generate Helm and Terraform artifacts from configured catalog templates.
bootstrap Bootstrap prerequisite CRDs and Argo CD onto the specified cluster
schema Generate JSON schema file for config structure
agents Scaffold an onboarding file for AI coding assistants (AGENTS.md)
catalog Manage platform catalogs
cluster Manage your kubara cluster configurations
help, h Shows a list of commands or help for one command
Global Options
--kubeconfig string Path to kubeconfig file (default: "~/.kube/config")
--work-dir string, -w string Working directory (default: ".")
--config-file string, -c string Path to the configuration file (default: "config.yaml")
--env-file string Path to the .env file (default: ".env")
--test-connection Check if Kubernetes cluster can be reached. List namespaces and exit
--base64 Enable base64 encode/decode mode
--encode Base64 encode input
--decode Base64 decode input
--string string Input string for base64 operation
--file string Input file path for base64 operation
--check-update Check online for a newer kubara release
--help, -h show help
--version, -v print the version
Update Check
kubara checks for newer GitHub releases on each run; disable with KUBARA_UPDATE_CHECK=0 ; run kubara --check-update for a live check.
Join the #kubara Slack channel to chat with other users of kubara or reach out to the maintainers directly. Use the public invite link ( http://slack.k8s.io/ ) to get an invite for the official Kubernetes Slack space.
Questions and bug reports: GitHub Issues
Discussions and Q&A: GitHub Discussions
Team and contributor guidance: CONTRIBUTING.md
Code of conduct: CODE_OF_CONDUCT.md
Contributions are welcome.
Please read CONTRIBUTING.md before opening a pull request.
Software source code: Apache 2.0
kubara is a single binary CLI tool written in Go providing a lightweight framework for bootstrapping Kubernetes platforms with production-proven best practices.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
28 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
