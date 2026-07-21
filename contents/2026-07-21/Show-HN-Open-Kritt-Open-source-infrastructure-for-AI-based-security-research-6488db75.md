---
source: "https://github.com/Kritt-ai/open-kritt"
hn_url: "https://news.ycombinator.com/item?id=48990871"
title: "Show HN: Open-Kritt – Open-source infrastructure for AI-based security research"
article_title: "GitHub - Kritt-ai/open-kritt: Orchestrate AI agents to find real vulnerabilities in code. · GitHub"
author: "GabiControlZ"
captured_at: "2026-07-21T12:20:42Z"
capture_tool: "hn-digest"
hn_id: 48990871
score: 1
comments: 0
posted_at: "2026-07-21T11:32:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-Kritt – Open-source infrastructure for AI-based security research

- HN: [48990871](https://news.ycombinator.com/item?id=48990871)
- Source: [github.com](https://github.com/Kritt-ai/open-kritt)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T11:32:04Z

## Translation

タイトル: Show HN: Open-Kritt – AI ベースのセキュリティ研究のためのオープンソース インフラストラクチャ
記事のタイトル: GitHub - Kritt-ai/open-kritt: コード内の実際の脆弱性を見つけるために AI エージェントを調整します。 · GitHub
説明: AI エージェントを調整して、コード内の実際の脆弱性を検出します。 - クリットアイ/オープンクリット

記事本文:
GitHub - Kritt-ai/open-kritt: コード内の実際の脆弱性を見つけるために AI エージェントを調整します。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
くりっとあい
/
オープンクリット
公共
通知

イオン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github .github アセット アセット バックエンド バックエンド データベース データベース docs-site docs-site docs docs エンジン エンジン executor-view executor-view フロントエンド フロントエンド local_repos local_repos スクリプト scripts .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .prettierignore .prettierignore .prettierrc.json .prettierrc.json .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス MAINTAINERS.md MAINTAINERS.md README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md バージョン バージョン docker-compose.dev.yml docker-compose.dev.yml docker-compose.yml docker-compose.yml kritt kritt release-please-config.json release-please-config.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントを調整して、コード内の実際の脆弱性を見つけます。
集中的な AI 分析を可能にする、オープンソースの自己ホスト型セキュリティ研究プラットフォーム
構成可能な検証と強化により、重複が排除され、ランク付けされた結果が得られます。
ウェブサイト ·
ドキュメント ·
はじめに ·
貢献する ·
研究論文・
Discordコミュニティ ·
ツイッター
モデルにリポジトリ全体を指定して脆弱性を見つけるように要求することはほとんどありません。
うまくいきます。 open·kritt は焦点を絞ったアプローチを採用しています。研究を小さなものに分割し、
明確に定義されたタスクを AI エージェント間で並行して実行し、その出力を組み合わせて
結果を検証して優先順位を付けることができます。
制御を必要とするセキュリティ研究者やセキュリティ志向の開発者向けに構築されています
以上

eir プロンプト、ワークフロー、モデル プロバイダー、およびインフラストラクチャ。
ワークフローを構築する — 焦点を絞ったプロンプトを再利用可能なセキュリティ調査プレイブックに連鎖させます。
スキャンの実行 - Codex を使用してリモートまたはローカルのリポジトリとその依存関係を分析します。
クロード・コードとか。
結果を検証する — ポストスクリプトを使用して問題を検証し、概念実証を構築し、
レポートを作成します。
結果に優先順位を付ける — カスタム重大度ランカー、一貫した検出スキーマを適用します。
そして自動重複除外。
独自のモデルにアクセスする - Codex ログインを使用するか、OpenAI 経由で接続します。
人間的、または OpenRouter。
実際のセキュリティ研究に基づいて構築されています。 Kritt チームは 1,500,000 ドル以上の収益を上げています
研究者名 Blockian でのバグ報奨金の支払い
(免疫力・
ハッケンプルーフ ·
blockian.xyz · @ControlZ_1337 )。
open·kritt は、その作業の背後にある内部プロジェクトをオープンソースで抽出したものです。
Git、Docker Compose を備えた Docker、Node.js 20 以降が必要です。リポジトリローカル
CLI にはインストール手順がありません。
git clone https://github.com/Kritt-ai/open-kritt
CD オープンクリット
./kritt セットアップ
./クリットスタート
スタックが実行されたら、http://localhost:5173 を開きます。あなただけ
モデルアクセスオプションが 1 つ必要です。 ./kritt セットアップでは、利用可能なログインと
API キー。 GITHUB_TOKEN はオプションであり、プライベート GitHub リポジトリにのみ必要です。
デフォルトのポートは 127.0.0.1 にバインドされ、バックエンドにはアプリケーションが含まれません。
認証。スタックを非公開にしておきます。
ツール対応エージェントは、書き込み可能なリポジトリを備えた使い捨てジョブ コンテナ内で root として実行されます。
コピーしてインターネットに直接アクセスできるようにすることで、ツールのインストール、ターゲットのコンパイル、テストの実行、
そして概念実証を構築します。専用の Docker ホストまたは VM 上で open·kritt を実行します。を見てください
信頼できないコードをスキャンする前に脅威モデルを確認してください。
前提条件、手動の Docker セットアップ、プロバイダー固有の手順については、

を読んでください
インストールガイドと
AI プロバイダーのセットアップ。
Mint を使用してドキュメントをローカルでプレビューします。
npm install -g mint
cd ドキュメント サイト
npm 実行開発
http://localhost:3001 を開いてサイトを表示します。
質問やアイデアは GitHub ディスカッションに属します。
バグや機能については GitHub Issues を使用してください
リクエスト。
貢献は大歓迎です。開発については CONTRIBUTING.md を読んでください
セットアップ、テスト コマンド、従来のコミット、DCO サインオフ要件。
セキュリティの脆弱性については、公開問題を通じてではなく、 SECURITY.md に従って非公開で報告してください。
open·kritt は、GNU Affero General Public License v3.0 に基づいてライセンスされています。
AI エージェントを調整して、コード内の実際の脆弱性を見つけます。
AGPL-3.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
31
フォーク
レポートリポジトリ
リリース
1
v1.1.0
最新の
2026 年 7 月 20 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Orchestrate AI agents to find real vulnerabilities in code. - Kritt-ai/open-kritt

GitHub - Kritt-ai/open-kritt: Orchestrate AI agents to find real vulnerabilities in code. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Kritt-ai
/
open-kritt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github .github assets assets backend backend database database docs-site docs-site docs docs engine engine executor-view executor-view frontend frontend local_repos local_repos scripts scripts .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .prettierignore .prettierignore .prettierrc.json .prettierrc.json .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md VERSION VERSION docker-compose.dev.yml docker-compose.dev.yml docker-compose.yml docker-compose.yml kritt kritt release-please-config.json release-please-config.json View all files Repository files navigation
Orchestrate AI agents to find real vulnerabilities in code.
An open-source, self-hosted security research platform that turns focused AI analysis
into de-duplicated, ranked findings with configurable validation and enrichment.
Website ·
Documentation ·
Getting started ·
Contributing ·
Research paper ·
Discord community ·
Twitter
Pointing a model at an entire repository and asking it to find vulnerabilities rarely
works well. open·kritt takes a focused approach: break the research into small,
well-defined tasks, run them across AI agents in parallel, and combine their output into
findings you can validate and prioritize.
It is built for security researchers and security-minded developers who want control
over their prompts, workflows, model providers, and infrastructure.
Build workflows — chain focused prompts into reusable security research playbooks.
Run scans — analyze remote or local repositories and their dependencies with Codex
or Claude Code.
Verify findings — use post-scripts to validate issues, build proofs of concept, and
produce reports.
Prioritize results — apply custom severity rankers, a consistent finding schema,
and automatic de-duplication.
Bring your own model access — use a Codex login or connect through OpenAI,
Anthropic, or OpenRouter.
Built from real security research. The Kritt team has earned over $1,500,000 in
bug-bounty payouts under the researcher name Blockian
( Immunefi ·
HackenProof ·
blockian.xyz · @ControlZ_1337 ).
open·kritt is the open-source distillation of the internal project behind that work.
You need Git, Docker with Docker Compose, and Node.js 20 or newer. The repository-local
CLI has no install step.
git clone https://github.com/Kritt-ai/open-kritt
cd open-kritt
./kritt setup
./kritt start
Open http://localhost:5173 once the stack is running. You only
need one model-access option; ./kritt setup guides you through the available logins and
API keys. A GITHUB_TOKEN is optional and only needed for private GitHub repositories.
The default ports bind to 127.0.0.1 , and the backend does not include application
authentication. Keep the stack private.
Tool-enabled agents run as root inside disposable job containers, with writable repository
copies and direct internet access so they can install tools, compile targets, run tests,
and build proofs of concept. Run open·kritt on a dedicated Docker host or VM; see the
threat model before scanning untrusted code.
For prerequisites, manual Docker setup, and provider-specific instructions, read the
installation guide and
AI provider setup .
Preview the documentation locally with Mint:
npm install -g mint
cd docs-site
npm run dev
Open http://localhost:3001 to view the site.
Questions and ideas belong in GitHub Discussions .
Use GitHub Issues for bugs and feature
requests.
Contributions are welcome. Read CONTRIBUTING.md for the development
setup, test commands, Conventional Commits, and DCO sign-off requirements.
Please report security vulnerabilities privately by following SECURITY.md , not through a public issue.
open·kritt is licensed under the GNU Affero General Public License v3.0 .
Orchestrate AI agents to find real vulnerabilities in code.
AGPL-3.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
31
forks
Report repository
Releases
1
v1.1.0
Latest
Jul 20, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
