---
source: "https://github.com/smarter-sh/smarter"
hn_url: "https://news.ycombinator.com/item?id=49324366"
title: "Smarter – open-source, declarative AI authoring platform (Django, K8s-inspired)"
article_title: "GitHub - smarter-sh/smarter: Smarter is an extensible AI authoring and resource management system · GitHub"
author: "lpm0073"
captured_at: "2026-08-16T23:11:31Z"
capture_tool: "hn-digest"
hn_id: 49324366
score: 1
comments: 0
posted_at: "2026-08-16T22:25:17Z"
tags:
  - hacker-news
  - translated
---

# Smarter – open-source, declarative AI authoring platform (Django, K8s-inspired)

- HN: [49324366](https://news.ycombinator.com/item?id=49324366)
- Source: [github.com](https://github.com/smarter-sh/smarter)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T22:25:17Z

## Translation

タイトル: Smarter – オープンソースの宣言型 AI オーサリング プラットフォーム (Django、K8s からインスピレーションを得た)
記事タイトル: GitHub - Smarter-sh/smarter: Smarter は拡張可能な AI オーサリングおよびリソース管理システムです · GitHub
説明: Smarter は、拡張可能な AI オーサリングおよびリソース管理システムです - Smarter-sh/smarter

記事本文:
GitHub - Smarter-sh/smarter: Smarter は拡張可能な AI オーサリングおよびリソース管理システムです · GitHub
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
もっと賢いよ
/
より賢い
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6,784 コミット 6,784 コミット .github .github .vscode .vscode 変更ログ

Changelogs docs docs helm/ charts/ Smarter Helm/ charts/ Smarter scripts スクリプト セットアップ セットアップ sigstore sigstore Smarter Smarter .codespellrc .codespellrc .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .flake8 .flake8 .gitattributes .gitattributes .gitignore .gitignore .markdownlint.json .markdownlint.json .mergify.yml .mergify.yml .pre-commit-config.yaml .pre-commit-config.yaml .prettierignore .prettierignore .prettierrc.json .prettierrc.json .pylintrc .pylintrc .readthedocs.yml .readthedocs.yml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md codepell.txt codepell.txt commitlint.config.js commitlint.config.js docker-compose.yml docker-compose.yml mypy.ini mypy.ini package.json package.json pyproject.toml pyproject.toml release.config.js release.config.js tox.ini tox.ini すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Smarter はブリティッシュ コロンビア大学で教育ツールとして使用されています
AIを教えるため。
Smarter は、エンタープライズ AI システムを構築および実行するためのオープンソース プラットフォームです。
オーケストレーション コードを記述する代わりに、エージェント、プロンプト、ツール、および
YAML マニフェストを使用して統合し、Web UI、CLI、API を通じてデプロイします。
またはKubernetes。
Smarter は分散した AI オーケストレーション コードを単一の宣言で置き換えます
環境全体で AI ワークフローを定義、テスト、実行するためのシステム。
1 クリックのクイックスタート展開
ドッカーと一緒に。
YAML マニフェストを使用して AI リソースを宣言的に定義します。
カスタム オーケストレーション コードを作成せずに、LLM を API や SQL データベースに接続します。
Windows、macOS、Linux 用のコマンドライン インターフェイス
とドッカー
Webコンソール/プロンプトエンジニアワークベンチ
SDK 、
UI社

コンポーネント 、
CLI ツール、
とIDEの統合
Smarter をベースに構築するための
公的にアクセス可能なオンラインドキュメント
およびセルフオンボーディングのリソース
プロジェクトをすぐに始めるためのオープンソース UI コンポーネント
このセットアップでは Docker を使用し、初めてインストールする場合は約 20 分かかります。
Windows、macOS、Linux
オペレーティングシステム
このリポジトリのルートにある .env に資格情報を追加します。
最小環境変数の詳細については、インライン ドキュメントを参照してください。
設定する必要があります。
アプリケーションをローカルで初期化、構築、実行します。
git clone https://github.com/smarter-sh/smarter
make help # リポジトリのルートに .env ファイルをスキャフォールディングします
#
# ****************************
＃ここでやめてください！
# ****************************
# プロジェクトのルート フォルダーにある .env に資格情報を追加します。
#
make init # Docker コンテナをプルし、Python 仮想環境を作成します。
# すべてのパッケージをインストールし、パッケージを作成して初期化します
# ローカル MySql データベース、サンプル AI リソースをプリロード
make run # すべての Docker コンテナを実行し、
# ローカル Web サーバー http://localhost:9357/
ユーザー admin@smarter.sh を使用して http://localhost:9357/login/ にログインし、
パスワードをよりスマートに。
よりスマートな開発者のオンボーディング I
よりスマートな開発者のオンボーディング II
よりスマートな開発者ワークフローのチュートリアル
Smarter は、AI リソースを管理するための yaml マニフェストベースのアプローチを実装します。
これは Kubernetes プロジェクトからインスピレーションを得たものです。
定義、構成、オーケストレーションを行うための統合された宣言型の方法を提供します。
AI リソースの作成と管理に必要なさまざまなリソース
REST API や SQL データベースなどの他のエンタープライズ リソースに統合されます。
また、迅速なエンジニアリング チームに直感的なワークベンチ アプローチを提供します。
強力な AI リソースの設計、プロトタイピング、テスト、導入、管理
エージェントワークフローを含む一般的な企業ユースケース向け

ああ、顧客向けチャット
ソリューションなど。別途管理されるものも含まれます。
React ベースのチャット UI
NPM、Wordpress、
Squarespace、Drupal、Office 365、Sharepoint、.Net、Netsuite、salesforce.com、および
SAP。があります
Golangコマンドラインインターフェース、
およびPyPiパッケージ
API 関数を独自の Python プロジェクトに統合します。に開発されています
大規模な組織で働く迅速なエンジニアリング チームをサポートします。したがって、
Smarter は、資格情報管理などの一般的なエンタープライズ機能を提供します。
チームワークグループ管理、役割ベースのセキュリティ、会計コストコード、
ロギング機能と監査機能。
Smarter は、LLM 間のシームレスな統合と相互運用を提供します。
DeepSeek、Google AI、メタ AI、OpenAI。 LLM プロバイダーに依存せず、
継続的に進化する付加価値リストへのシームレスな統合を提供します
セキュリティ管理、迅速なコンテンツモデレーション、監査、コストのためのサービス
会計とワークフロー管理。 Smarter はクラウドネイティブであり、上で実行されます
データセンターまたはクラウドのオンサイトの Kubernetes。
大規模に実行する場合、Smarter はコスト効率が高くなります。拡張性があり、
必要のないコンパクトなコアの哲学に基づいて設計されています。
カスタマイズもフォークも。水平方向にスケーラブルです。ネイティブです
マルチテナントに対応しており、既存のシステムと一緒にインストールできます。 ## クイックスタート
ghcr.io/smarter-sh/charts/smarter を参照してください。
または アーティファクトハブ 。
ドキュメントを読む: docs.smarter.sh
バグは GitHub の問題ページに報告してください。
このプロジェクトのために。
Smarter は拡張可能な AI オーサリングおよびリソース管理システムです
Readme AGPL-3.0 ライセンス 行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
37 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
共有しないでください

私の個人情報

## Original Extract

Smarter is an extensible AI authoring and resource management system - smarter-sh/smarter

GitHub - smarter-sh/smarter: Smarter is an extensible AI authoring and resource management system · GitHub
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
smarter-sh
/
smarter
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6,784 Commits 6,784 Commits .github .github .vscode .vscode changelogs changelogs docs docs helm/ charts/ smarter helm/ charts/ smarter scripts scripts setup setup sigstore sigstore smarter smarter .codespellrc .codespellrc .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .flake8 .flake8 .gitattributes .gitattributes .gitignore .gitignore .markdownlint.json .markdownlint.json .mergify.yml .mergify.yml .pre-commit-config.yaml .pre-commit-config.yaml .prettierignore .prettierignore .prettierrc.json .prettierrc.json .pylintrc .pylintrc .readthedocs.yml .readthedocs.yml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md codespell.txt codespell.txt commitlint.config.js commitlint.config.js docker-compose.yml docker-compose.yml mypy.ini mypy.ini package.json package.json pyproject.toml pyproject.toml release.config.js release.config.js tox.ini tox.ini View all files Repository files navigation
Smarter is used as an instructional tool at University of British Columbia
for teaching AI.
Smarter is an open-source platform for building and running enterprise AI systems.
Instead of writing orchestration code, you define agents, prompts, tools, and
integrations using YAML manifests, then deploy them through a web UI, CLI, API,
or Kubernetes.
Smarter replaces scattered AI orchestration code with a single declarative
system for defining, testing, and running AI workflows across environments.
1-click Quickstart deployment
with Docker.
Define AI resources declaratively using YAML manifests.
Connect LLMs to APIs and SQL databases without writing custom orchestration code.
command-line interface for Windows, macOS, Linux
and Docker
web console / prompt engineer workbench
SDKs ,
UI components ,
CLI tools ,
and IDE integrations
for building on Smarter
publicly accessible online documentation
and self onboarding resources
open source UI components for jump starting projects
This setup uses Docker and takes around 20 minutes for first time installations.
Windows , macOS , Linux
operating system
Add your credentials to .env in the root of this repo.
See the inline documentation for details on the minimum environment variables
that you will need to set.
Initialize, build and run the application locally.
git clone https://github.com/smarter-sh/smarter
make help # scaffolds a .env file in the root of the repo
#
# ****************************
# STOP HERE!
# ****************************
# Add your credentials to .env located in the project root folder.
#
make init # pulls Docker containers, creates a Python virtual environment,
# installs all packages, creates and initializes a
# local MySql database, preloads example AI resources
make run # runs all docker containers and starts a
# local web server http://localhost:9357/
Login at http://localhost:9357/login/ with user admin@smarter.sh and
password smarter .
Smarter Developer Onboarding I
Smarter Developer Onboarding II
Smarter Developer Workflow Tutorial
Smarter implements a yaml manifest-based approach to managing AI resources
that is inspired by the Kubernetes project.
It provides a unified, declarative way to define, configure, and orchestrate
the disparate resources that are required for creating and managing AI resources
that integrate to other enterprise resources like REST API's and Sql databases.
And it gives prompt engineering teams an intuitive workbench approach to
designing, prototyping, testing, deploying and managing powerful AI resources
for common corporate use cases including agentic workflows, customer facing chat
solutions, and more. It includes a separately managed
React-based chat UI that is
compatible with a wide variety of front end ecosystems including NPM, Wordpress,
Squarespace, Drupal, Office 365, Sharepoint, .Net, Netsuite, salesforce.com, and
SAP. There is a
Golang command-line interface ,
and a PyPi package for
integrating the API functions into your own Python projects. It is developed to
support prompt engineering teams working in large organizations. Accordingly,
Smarter provides common enterprise features such as credentials management,
team workgroup management, role-based security, accounting cost codes, and
logging and audit capabilities.
Smarter provides seamless integration and interoperation between LLMs from
DeepSeek, Google AI, Meta AI and OpenAI. It is LLM provider-agnostic, and
provides seamless integrations to a continuously evolving list of value added
services for security management, prompt content moderation, audit, cost
accounting, and workflow management. Smarter is cloud native and runs on
Kubernetes, on-site in your data center or in the cloud.
Smarter is cost effective when running at scale. It is extensible and
architected on the philosophy of a compact core that does not require
customization nor forking. It is horizontally scalable. It is natively
multi-tenant, and can be installed alongside your existing systems. ## Quickstart
See ghcr.io/smarter-sh/charts/smarter
or Artifact Hub .
Read the Docs: docs.smarter.sh
Please report bugs to the GitHub Issues Page
for this project.
Smarter is an extensible AI authoring and resource management system
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
37 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
