---
source: "https://github.com/kubetable/kubetable"
hn_url: "https://news.ycombinator.com/item?id=48418161"
title: "KubeTable – your Kubernetes databases, one click away"
article_title: "GitHub - kubetable/kubetable: Local-first desktop client for discovering, connecting to, and querying databases running inside Kubernetes. · GitHub"
author: "emmaera"
captured_at: "2026-06-05T21:34:20Z"
capture_tool: "hn-digest"
hn_id: 48418161
score: 1
comments: 1
posted_at: "2026-06-05T20:55:59Z"
tags:
  - hacker-news
  - translated
---

# KubeTable – your Kubernetes databases, one click away

- HN: [48418161](https://news.ycombinator.com/item?id=48418161)
- Source: [github.com](https://github.com/kubetable/kubetable)
- Score: 1
- Comments: 1
- Posted: 2026-06-05T20:55:59Z

## Translation

タイトル: KubeTable – ワンクリックで Kubernetes データベースを利用できます
記事のタイトル: GitHub - kubetable/kubetable: Kubernetes 内で実行されているデータベースを検出、接続、クエリするためのローカルファースト デスクトップ クライアント。 · GitHub
説明: Kubernetes 内で実行されているデータベースを検出、接続、クエリするためのローカルファーストのデスクトップ クライアントです。 - kubetable/kubetable

記事本文:
GitHub - kubetable/kubetable: Kubernetes 内で実行されているデータベースを検出、接続、クエリするためのローカルファーストのデスクトップ クライアントです。 · GitHub
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
キューブテーブル
/
キューブテーブル
公共
通知
通知設定を変更するにはサインインする必要があります

イングス
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github .github docs docs public public src-tauri src-tauri src src test test .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock Index.html Index.html package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
KubeTable Community は、Kubernetes 内で実行されているデータベースを操作するためのローカルファーストのデスクトップ クライアントです。
kubeconfig を読み取り、データベース サービスを検出し、ローカル ポートフォワードを開き、 kubectl port-forward 、ターミナル タブ、および個別のデータベース ツールを切り替えることなくクエリ ワークスペースを提供します。
ローカルの kubeconfig ファイルからの Kubernetes サービスの検出
ローカルポートフォワードライフサイクル管理
永続的なローカル ドラフトを備えたクエリ エディター
リレーショナル データベースおよびキー/コレクション スタイル データベース用のスキーマ エクスプローラー
フィルタリング、並べ替え、ページネーション、行編集、CSV/JSON エクスポートを備えた結果テーブル
SQL フォーマット、クエリ履歴、保存されたクエリ
より安全な検査のための読み取り専用接続モード
ローカルクラスターソース管理
テストおよび開発クラスター用のデータベース デプロイ アシスタント
PostgreSQL、MySQL、Redis、MongoDB のテスト マニフェスト
KubeTable コミュニティはローカルで実行されるように設計されています。
コアデスクトップワークフローには外部サービスへの依存はありません
アプリは引き続き機密ローカル データ (kubeconfig、クラスター メタデータ、データベース資格情報、SQL テキスト、クエリ結果) に触れる可能性があります。実際の kubeconfig、認証情報、トークン、秘密キー、運用 URL、クエリ結果、シークレットを含むスクリーンショット、

または、問題やプルリクエストの詳細なログ。
kubeconfig、認証情報、クエリ実行、ロギング、ストレージ、またはネットワーク動作を処理するコードを変更する前に、 docs/SECURITY_MODEL.md をお読みください。
パブリック アプリは、ダウンストリーム ビルド用の小さなプラグイン ホストを公開します。
「./App」からアプリをインポートします。
デフォルト関数のエクスポート Root ( ) {
return < アプリプラグイン = { [ ] } /> ;
}
プラグインは、コミュニティのソース コードを変更せずに、テレメトリ アダプター、コマンド パレット アクション、トップバー アクション、クラスター マネージャー セクションなどのアプリレベルの統合を追加できます。具体的なホスト型統合、リリース インフラストラクチャ、およびサービス固有の認証情報をこのリポジトリの外部に保持します。
Tauri に必要なプラットフォームの依存関係
検査したいクラスターへの kubectl アクセス
Tauri プラットフォームのセットアップには、オペレーティング システムの公式 Tauri 前提条件を使用してください。
バンインストール
フロントエンドを開発モードで実行します。
バンラン開発
デスクトップ アプリを実行します。
バンランタウリ開発者
フロントエンドを構築します。
バンランビルド
Rust バックエンドを確認します。
cd src-tauri
貨物検査
デスクトップ バンドルを構築します。
バンランタウリビルド
テストクラスター
test/ ディレクトリには、サンプルの PostgreSQL、MySQL、Redis、MongoDB サービスを含む小規模な Kubernetes テスト環境が含まれています。
それを現在のクラスターに適用します。
kubectl apply -f test/kubetable-test.yaml
再度削除します。
kubectl delete -f test/kubetable-test.yaml
ヘルパー スクリプトもあります。
./test/kubetable-test.sh アップ
./test/kubetable-test.sh ステータス
./test/kubetable-test.sh ダウン
テスト データにはローカル クラスターまたは使い捨てクラスターのみを使用します。
src/React デスクトップ UI
src/App.tsx アプリシェルとプラグインホスト
src/lib/plugins.ts 中立的なプラグインインターフェイス
src/lib/tauri.ts Tauri コマンドのフロントエンド バインディング
src-tauri/Tauri アプリと Rust バックエンド
src-tauri/src/discovery/ Kubernetes の検出と資格情報の検出
src-tauri/src/portforward/Lo

カリフォルニアトンネル管理
src-tauri/src/adapters/ データベースアダプター
テスト/Kubernetes テスト環境
docs/セキュリティモデルとリリースノート
開発ノート
コミュニティをローカルファーストに保ちます。プル リクエストでは、デスクトップ ワークフローに外部インフラストラクチャを必要とするシークレット、非表示のエンドポイント、実際の認証情報、またはサービス統合を追加しないでください。
プル リクエストを開く前に:
バンランビルド
cd src-tauri
貨物検査
変更が kubeconfig の解析、資格情報の検出、クエリの実行、ローカル ストレージ、ログ、またはネットワークの動作に影響を与える場合は、プル リクエストに短いセキュリティに関するメモを含めてください。
プル リクエストを開く前に CONTRIBUTING.md をお読みください。
セキュリティ レポートは SECURITY.md に従う必要があります。保守者が脆弱性を確認する前に、公開問題で脆弱性を開示しないでください。
KubeTable コミュニティは、Apache-2.0 に基づいてライセンスされています。 「ライセンス」を参照してください。
Kubernetes 内で実行されているデータベースを検出、接続、クエリするためのローカルファーストのデスクトップ クライアントです。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

Local-first desktop client for discovering, connecting to, and querying databases running inside Kubernetes. - kubetable/kubetable

GitHub - kubetable/kubetable: Local-first desktop client for discovering, connecting to, and querying databases running inside Kubernetes. · GitHub
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
kubetable
/
kubetable
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github .github docs docs public public src-tauri src-tauri src src test test .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock index.html index.html package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
KubeTable Community is a local-first desktop client for working with databases running inside Kubernetes.
It reads your kubeconfig, discovers database services, opens local port-forwards, and gives you a query workspace without switching between kubectl port-forward , terminal tabs, and separate database tools.
Kubernetes service discovery from local kubeconfig files
Local port-forward lifecycle management
Query editor with persistent local drafts
Schema explorer for relational databases and key/collection-style databases
Result table with filtering, sorting, pagination, row editing, and CSV/JSON export
SQL formatting, query history, and saved queries
Read-only connection mode for safer inspection
Local cluster source management
Database deploy assistant for test and development clusters
Test manifests for PostgreSQL, MySQL, Redis, and MongoDB
KubeTable Community is designed to run locally:
no external service dependency for the core desktop workflow
The app can still touch sensitive local data: kubeconfigs, cluster metadata, database credentials, SQL text, and query results. Do not put real kubeconfigs, credentials, tokens, private keys, production URLs, query results, screenshots with secrets, or verbose logs in issues or pull requests.
Read docs/SECURITY_MODEL.md before changing code that handles kubeconfigs, credentials, query execution, logging, storage, or network behavior.
The public app exposes a small plugin host for downstream builds:
import App from "./App" ;
export default function Root ( ) {
return < App plugins = { [ ] } /> ;
}
Plugins can add app-level integrations such as telemetry adapters, command-palette actions, topbar actions, and cluster-manager sections without changing Community source code. Keep concrete hosted integrations, release infrastructure, and service-specific credentials outside this repository.
Platform dependencies required by Tauri
kubectl access to any cluster you want to inspect
For Tauri platform setup, use the official Tauri prerequisites for your operating system.
bun install
Run the frontend in development mode:
bun run dev
Run the desktop app:
bun run tauri dev
Build the frontend:
bun run build
Check the Rust backend:
cd src-tauri
cargo check
Build a desktop bundle:
bun run tauri build
Test Cluster
The test/ directory contains a small Kubernetes test environment with sample PostgreSQL, MySQL, Redis, and MongoDB services.
Apply it to your current cluster:
kubectl apply -f test/kubetable-test.yaml
Remove it again:
kubectl delete -f test/kubetable-test.yaml
There is also a helper script:
./test/kubetable-test.sh up
./test/kubetable-test.sh status
./test/kubetable-test.sh down
Use only local or disposable clusters for test data.
src/ React desktop UI
src/App.tsx App shell and plugin host
src/lib/plugins.ts Neutral plugin interfaces
src/lib/tauri.ts Frontend bindings for Tauri commands
src-tauri/ Tauri app and Rust backend
src-tauri/src/discovery/ Kubernetes discovery and credential detection
src-tauri/src/portforward/ Local tunnel management
src-tauri/src/adapters/ Database adapters
test/ Kubernetes test environment
docs/ Security model and release notes
Development Notes
Keep Community local-first. Pull requests should not add secrets, hidden endpoints, real credentials, or service integrations that require external infrastructure for the desktop workflow.
Before opening a pull request:
bun run build
cd src-tauri
cargo check
If your change affects kubeconfig parsing, credential detection, query execution, local storage, logs, or network behavior, include a short security note in the pull request.
Read CONTRIBUTING.md before opening a pull request.
Security reports should follow SECURITY.md . Please do not disclose vulnerabilities in public issues before maintainers have had a chance to review them.
KubeTable Community is licensed under Apache-2.0. See LICENSE .
Local-first desktop client for discovering, connecting to, and querying databases running inside Kubernetes.
Apache-2.0 license
Code of conduct
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
