---
source: "https://github.com/ngaut/agent-git-service"
hn_url: "https://news.ycombinator.com/item?id=48729434"
title: "A GitHub-compatible Git service built for AI agents"
article_title: "GitHub - ngaut/agent-git-service: Reimplement GitHub for Agents. · GitHub"
author: "shenli3514"
captured_at: "2026-06-30T07:13:48Z"
capture_tool: "hn-digest"
hn_id: 48729434
score: 1
comments: 0
posted_at: "2026-06-30T07:05:44Z"
tags:
  - hacker-news
  - translated
---

# A GitHub-compatible Git service built for AI agents

- HN: [48729434](https://news.ycombinator.com/item?id=48729434)
- Source: [github.com](https://github.com/ngaut/agent-git-service)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T07:05:44Z

## Translation

タイトル: AI エージェント向けに構築された GitHub 互換の Git サービス
記事のタイトル: GitHub - ngaut/agent-git-service: エージェント向けの GitHub の再実装。 · GitHub
説明: エージェント用に GitHub を再実装します。 GitHub でアカウントを作成して、ngaut/agent-git-service の開発に貢献してください。

記事本文:
GitHub - ngaut/agent-git-service: エージェント用に GitHub を再実装します。 · GitHub
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
ンガウト
/
エージェント-git-サービス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

コード 「その他のアクション」メニューを開く フォルダーとファイル
98 Commits 98 Commits .github .github auth auth cli cli cmd/ gh-server cmd/ gh-server config config docs docs e2e e2e internal internal scripts scripts server server .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェント、自動化、および自動化のための自己ホスト型の GitHub 互換 API サーバー
developer workflows.
Agent-git-service を使用すると、GitHub を使用するクライアントがリポジトリを操作できるようになります。
独自のエージェントファースト設計で、AI エージェントを第一級市民として扱います。
永続的なアイデンティティ、スコープ付きトークン、デフォルトのワークスペース、所有権/回復
flows.
GitHub スタイルの REST v3、GraphQL v4、OAuth デバイス フロー、および Git Smart を公開します。
実際のベア Git リポジトリおよび製品にリポジトリ データを保存する際の HTTP
TiDB/MySQL 互換ストレージ内のメタデータ。
開発バイナリの現在の名前は gh-server です。
GitHub と互換性のあるサービスが必要な場合は、agent-git-service を使用します。
エージェントが実行される場所で実行します。
リポジトリと製品メタデータを管理下に置きます。
新しいクライアント プロトコルを開発するのではなく、既存の GitHub クライアントをサポートします。
エージェントに永続的なアカウント、トークン、オプションの人間バインディング、およびリポジトリを提供します
transfer flows.
Git ネイティブのクローン、フェッチ、プッシュ、参照、差分、マージ、履歴を保存します。
本物の裸の Git リポジトリ。
ベンダーの GitHub CLI 受け入れスイートを通じて互換性を検証します。
能力
エージェント-git-サービス
GitHub.com
How it differs
GitHub スタイルのリポジトリ、問題、ラベル、Wiki、Git 履歴
はい
Y

エス
どちらも、使い慣れた GitHub 形式のコラボレーション ワークフローをサポートしています。 Agent-git-service は、これらのワークフローを自己ホスト型バックエンドで利用できるようにします。
Git Smart HTTP、REST v3、OAuth デバイス フロー、および一般的な gh ワークフロー
はい
はい
GitHub を使用する既存のクライアントは、新しいプロトコルを学習することなく、agent-git-service に対して動作できます。
ファーストクラスの永続エージェントアカウント
はい
いいえ
GitHub はボット、アプリ、PAT、マシン ユーザーをサポートしますが、agent-git-service はエージェントに独自の永続アカウント、トークン、およびデフォルトのワークスペースを提供します。
リポジトリ、組織、チーム全体にわたる直接エージェント権限
はい
部分的
Agent-Git-Service では、アプリや PAT 間接に依存するのではなく、エージェントにコラボレーター、組織、チーム、およびチーム リポジトリへのアクセスを直接付与できます。
人間とエージェントの結合と回復
はい
いいえ
Agent-git-service は、人間とエージェントのバインディング、接続されたログイン、セッションの切り替え、および回復フローをファーストクラスのワークフローとしてサポートします。
セルフホスト型エージェントサービス
はい
いいえ
Agent-git-service は、アイデンティティ、ストレージ、レート制限ポリシーをローカルに制御しながら、エージェントとデータが存在する場所で実行できます。
ローカル データ、Git ストレージ、メタデータの所有権
はい
いいえ
Agent-git-service は、製品メタデータを TiDB/MySQL 互換ストレージに保持しながら、リポジトリを実際のベア Git リポジトリとして保存します。
完全にホストされた GitHub 製品エコシステム
部分的
はい
GitHub.com は、アクション、セキュリティ製品、マーケットプレイス、トラフィック/コミュニティ機能、ロングテール API まで幅広く対応しています。
完全な GitHub GraphQL スキーマ パリティ
部分的
はい
Agent-git-service は、GitHub GraphQL の完全な同等性ではなく、選択したワークフローに対して GraphQL 互換性を提供します。
機能強化
強化
追加されるもの
エージェントのアイデンティティとガバナンス
永続エージェント アカウント、API トークン、人間によるバインディング/リカバリ、セッションの切り替え、デフォルト リポジトリ、リポジトリ転送フロー、直接リポジトリ/組織/チーム権限付与
GitHub互換コア
RE

ST v3、GraphQL 互換性、OAuth デバイス フロー、Git Smart HTTP、一般的なワークフローの gh 受け入れ範囲
課題ワークスペース
入力信号、存在、添付ファイル、既読状態、未読数、固定されたコメント、および反応
ウィキメモリ
Git ベースのページ、履歴、検索、ラベル、バックリンク、ページの移動、調整、修復、およびコンパクトな操作
セマンティック検索
オプションの埋め込みに裏付けされた問題とプル リクエストの検索
自己ホスト型の操作
ローカル データと Git ストレージ、ローカル レート制限ポリシー、Prometheus メトリクス、準備状況チェック、構造化ログ、Grafana ダッシュボード
既知の GitHub 互換性ギャップは以下で追跡されます。
docs/github-api-compatibility-matrix.md 。
このローカル パスは TiDB Zero を使用します
使い捨ての TiDB データベース用。
このクイックスタートを実行する前に、curl と jq をインストールしてください。以下のスニペットでは、
両方のツールを使用して、TiDB Zero インスタンスを作成し、MySQL DSN を構築します。
git clone https://github.com/ngaut/agent-git-service.git
cd エージェント-git-サービス
cp .env.example .env
ZERO_INSTANCE= " $(
カール -fsS -X POST https://zero.tidbapi.com/v1beta1/instances \
-H " Content-Type: application/json " \
-d ' {"タグ":"エージェント-git-サービス-クイックスタート"} '
）」
エクスポート DB_DSN= " $(
printf ' %s ' " $ZERO_INSTANCE " | jq -r '
$c としてのインスタンス接続 |
"\($c.username):\($c.password)@tcp(\($c.host):\($c.port))/test?parseTime=true&timeout=10s&tls=true"
'
）」
printf ' TiDB Zero クレーム URL: %s\n ' " $(
printf ' %s ' " $ZERO_INSTANCE " | jq -r ' .instance.claimInfo.claimUrl '
）」
./cmd/gh-server を実行します。
データベースを保持したい場合は、要求 URL から TiDB Zero インスタンスを要求します。
評価後。本番環境では、TiDB Cloud Starter インスタンスを作成し、
完全な docs/production-deployment.md に従ってください。
ガイド。
gh CLI、curl、Git プッシュの例を含む完全なローカル設定については、
docs/quickstart.md を参照してください。
作る

ld # compile gh-server
make check # build + go vet
make test-unit # go test -v ./...
make test # gh CLI acceptance tests; requires a running local server
make test-e2e # shell E2E flows under e2e/
Local setup helpers:
make setup # 外部 DB_DSN を使用した永続的なセットアップ
make test-setup # tiup Playground を使用したテストのみのセットアップ
make run-bg # バックグラウンドでローカルサーバーを起動します
make stop # stop it
make status # show local status
make run-bg はまず、使用される特権付き github.localhost リスナー パスを試行します。
by acceptance tests.パスワードなしの sudo が使用できない場合は、ポートにフォールバックします
8080 ;アクセプタンススイートが期待するため、make テストはすぐに失敗します。
http://github.localhost on port 80 .
Licensed under the Apache License 2.0 .
Reimplement GitHub for Agents.
Apache-2.0ライセンス
Code of conduct
読み込み中にエラーが発生しました。このページをリロードしてください。
8
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

Reimplement GitHub for Agents. Contribute to ngaut/agent-git-service development by creating an account on GitHub.

GitHub - ngaut/agent-git-service: Reimplement GitHub for Agents. · GitHub
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
ngaut
/
agent-git-service
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
98 Commits 98 Commits .github .github auth auth cli cli cmd/ gh-server cmd/ gh-server config config docs docs e2e e2e internal internal scripts scripts server server .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
A self-hosted, GitHub-compatible API server for agents, automation, and
developer workflows.
agent-git-service lets GitHub-speaking clients work with repositories you
own, and its agent-first design treats AI agents as first-class citizens with
durable identities, scoped tokens, default workspaces, and ownership/recovery
flows.
It exposes GitHub-style REST v3, GraphQL v4, OAuth device flow, and Git Smart
HTTP while storing repository data in real bare Git repositories and product
metadata in TiDB/MySQL-compatible storage.
The development binary is currently named gh-server .
Use agent-git-service when you want a GitHub-compatible service that can
run where your agents run:
Keep repositories and product metadata under your control.
Support existing GitHub clients instead of inventing new client protocols.
Give agents durable accounts, tokens, optional human binding, and repository
transfer flows.
Preserve Git-native clone, fetch, push, refs, diffs, merges, and history in
real bare Git repositories.
Validate compatibility through the vendored GitHub CLI acceptance suite.
Capability
agent-git-service
GitHub.com
How it differs
GitHub-style repositories, issues, labels, wiki, and Git history
Yes
Yes
Both support familiar GitHub-shaped collaboration workflows. agent-git-service keeps these workflows available on a self-hosted backend.
Git Smart HTTP, REST v3, OAuth device flow, and common gh workflows
Yes
Yes
Existing GitHub-speaking clients can work against agent-git-service without learning a new protocol.
First-class durable agent accounts
Yes
No
GitHub supports bots, Apps, PATs, and machine users, but agent-git-service gives agents their own durable accounts, tokens, and default workspaces.
Direct agent permissions across repos, orgs, and teams
Yes
Partial
In agent-git-service, agents can be granted collaborator, org, team, and team-repo access directly instead of relying on App or PAT indirection.
Human-agent binding and recovery
Yes
No
agent-git-service supports human-agent binding, connected login, switch sessions, and recovery flows as first-class workflows.
Self-hosted agent service
Yes
No
agent-git-service can run where your agents and data live, with local control over identity, storage, and rate-limit policy.
Local data, Git storage, and metadata ownership
Yes
No
agent-git-service stores repositories as real bare Git repos while keeping product metadata in TiDB/MySQL-compatible storage.
Full hosted GitHub product ecosystem
Partial
Yes
GitHub.com is broader across Actions, security products, marketplace, traffic/community features, and long-tail APIs.
Full GitHub GraphQL schema parity
Partial
Yes
agent-git-service provides GraphQL compatibility for selected workflows, not full GitHub GraphQL parity.
Enhancements
Enhancement
What it adds
Agent identities and governance
Durable agent accounts, API tokens, human binding/recovery, switch sessions, default repos, repository transfer flows, and direct repo/org/team permission grants
GitHub-compatible core
REST v3, GraphQL compatibility, OAuth device flow, Git Smart HTTP, and gh acceptance coverage for common workflows
Issue workspace
Typing signals, presence, attachments, read state, unread counts, pinned comments, and reactions
Wiki memory
Git-backed pages, history, search, labels, backlinks, page moves, reconcile, repair, and compact operations
Semantic search
Optional embedding-backed issue and pull request search
Self-hosted operations
Local data and Git storage, local rate-limit policy, Prometheus metrics, readiness checks, structured logs, and a Grafana dashboard
Known GitHub-compatibility gaps are tracked in
docs/github-api-compatibility-matrix.md .
This local path uses TiDB Zero
for a disposable TiDB database.
Install curl and jq before running this quickstart. The snippet below uses
both tools to create a TiDB Zero instance and build the MySQL DSN.
git clone https://github.com/ngaut/agent-git-service.git
cd agent-git-service
cp .env.example .env
ZERO_INSTANCE= " $(
curl -fsS -X POST https://zero.tidbapi.com/v1beta1/instances \
-H " Content-Type: application/json " \
-d ' {"tag":"agent-git-service-quickstart"} '
) "
export DB_DSN= " $(
printf ' %s ' " $ZERO_INSTANCE " | jq -r '
.instance.connection as $c |
"\($c.username):\($c.password)@tcp(\($c.host):\($c.port))/test?parseTime=true&timeout=10s&tls=true"
'
) "
printf ' TiDB Zero claim URL: %s\n ' " $(
printf ' %s ' " $ZERO_INSTANCE " | jq -r ' .instance.claimInfo.claimUrl '
) "
go run ./cmd/gh-server
Claim the TiDB Zero instance from its claim URL if you want to keep the database
after evaluation. For production, create a TiDB Cloud Starter instance and
follow the full docs/production-deployment.md
guide.
For the complete local setup, including gh CLI, curl , and Git push examples,
see docs/quickstart.md .
make build # compile gh-server
make check # build + go vet
make test-unit # go test -v ./...
make test # gh CLI acceptance tests; requires a running local server
make test-e2e # shell E2E flows under e2e/
Local setup helpers:
make setup # persistent setup with an external DB_DSN
make test-setup # test-only setup using tiup playground
make run-bg # start the local server in the background
make stop # stop it
make status # show local status
make run-bg first tries the privileged github.localhost listener path used
by acceptance tests. If passwordless sudo is unavailable, it falls back to port
8080 ; make test will then fail fast because the acceptance suite expects
http://github.localhost on port 80 .
Licensed under the Apache License 2.0 .
Reimplement GitHub for Agents.
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
8
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
