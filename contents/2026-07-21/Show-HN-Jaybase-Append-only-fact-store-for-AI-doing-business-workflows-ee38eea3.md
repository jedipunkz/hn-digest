---
source: "https://github.com/kyle-visner/jaybase"
hn_url: "https://news.ycombinator.com/item?id=48999936"
title: "Show HN: Jaybase: Append-only fact store for AI doing business workflows"
article_title: "GitHub - kyle-visner/jaybase: Jaybase is a append-only fact store for AI agents doing critical business workflows · GitHub"
author: "kvisner"
captured_at: "2026-07-21T23:48:19Z"
capture_tool: "hn-digest"
hn_id: 48999936
score: 1
comments: 0
posted_at: "2026-07-21T23:47:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Jaybase: Append-only fact store for AI doing business workflows

- HN: [48999936](https://news.ycombinator.com/item?id=48999936)
- Source: [github.com](https://github.com/kyle-visner/jaybase)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T23:47:16Z

## Translation

タイトル: HN を表示: Jaybase: ビジネス ワークフローを実行する AI のための追加専用ファクト ストア
記事のタイトル: GitHub - kyle-visner/jaybase: Jaybase は、重要なビジネス ワークフローを実行する AI エージェントのための追加専用ファクト ストアです · GitHub
説明: Jaybase は、重要なビジネス ワークフローを実行する AI エージェント用の追加専用ファクト ストアです - kyle-visner/jaybase
HN テキスト: HN さん、私は企業向けにカスタム AI エージェントを構築しています。そして気づいたことの 1 つは、誰もが AI を使用して実際の重要な仕事をしたいと考えているということですが、エージェントは時々、少しレールから外れることがあるため、AI を信頼するのは難しいということです。エージェントがすべての顧客データを削除したり、帳簿を台無しにしたりすることを望む人は誰もいません。そのため、検索、要約、執筆アシスタントとしての使用に限定される傾向があります。問題は、メモやスプレッドシートからデータベースに至るまで、私たちがデータを保存するほとんどの方法が、人間のスピードで作業し間違いを犯す人間、または決定論的なソフトウェア システム向けに設計されていることです。 AI エージェントはマシンの速度で非決定的に動作しますが、これは新たな課題です。それが私がJaybaseを作った理由です。 Jaybase は、エージェント ワークフロー用の追加専用のファクト ストアです。これは、git とブロックチェーンの背後にある同じテクノロジーであるマークル有向非巡回グラフに基づいているため、すべての書き込みは監査可能、再生可能、および元に戻すことができます。エージェントが暴走して大量のデータを破損した場合でも、回復は簡単で、何が問題だったのかを監査するのも簡単です。また、固定のスキーマがないため、エージェントを追加したり、エージェントのスキルが向上したりすると、自然に追加するデータの種類も増えます。コメント欄でどんな質問にも喜んでお答えします。コミュニティからのフィードバックもお待ちしています。ありがとう！

記事本文:
GitHub - kyle-visner/jaybase: Jaybase は、重要なビジネス ワークフローを実行する AI エージェントのための追加専用のファクト ストアです · GitHub
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
カイル・ヴィスナー
/
ジェイベース
公共
通知
通知を変更するにはサインインする必要があります

フィクション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
29 コミット 29 コミット .github .github cmd/ jaybase-server cmd/ jaybase-server デプロイ デプロイ docs docs Secrets Secrets サーバー サーバー .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile ライセンス ライセンス README.md README.md behaviour_test.go Behavior_test.go compose.yaml compose.yaml go.mod go.mod hosts_test.go hosts_test.go llm.md llm.md lock_fcntl.go lock_fcntl.go lock_other.go lock_other.go lock_unix.go lock_unix.go migration.go migration.go migration_test.go migration_test.go Performance_test.go Performance_test.go snapshot.go snapshot.go storage.go storage.go storage_test.go storage_test.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Jaybase は、重要な機能を信頼できる AI エージェントのための追加専用のファクト ストアです。
ビジネスデータ。エージェントは柔軟な JSON ファクトを追加できますが、ホストされた API は追加できません
履歴を書き換えたり削除したりします。すべての書き込みは属性付けされ、暗号化され、安全に再試行できます。
古い状態がチェックされ、再生可能です。
git clone https://github.com/kyle-visner/jaybase.git
CDジェイベース
./cmd/jaybase-server をインストールしてください
jaybase-server init ./secrets
イニシャライザは、リーダー、ライター、および管理者のトークンを 1 回出力します。それらを
パスワード マネージャーを選択し、「Jaybase の実行」に進みます。
Jaybase は、読み取り専用の副操縦士から、
操作を許可されているエージェント。会計、業務、コンプライアンス、
承認や、エージェントが重要なデータを書き込むその他の作業では、間違いはそのままにしておかなければなりません
目に見えて修正可能であり、事実の形状は仕事とともに進化します。
Jaybase はシングルテナント システム用に設計されています: 1 つの組織、1 つの信頼
境界、およびストアごとに 1 つの書き込みプロセス。多くのエージェントとアプリケーションは、

ダッシュボード、内部ツール、API、自動化されたストアを共有する
ワークフロー。 Jaybase は、グローバルに分散されたマルチテナントであることを意図したものではありません。
Web アプリケーションのバックエンド。
従来のデータベースは、決定論的なアプリケーション コードがすべての読み取りを所有していることを前提としています。
書く。エージェントは判断を下し、不確実な作業を再試行し、場合によっては不確実な行動をとります。
予期せぬ方法で、マシンの速度で。変更可能なデータベースは、間違った決定を変える可能性があります。
暴走ループ、または誰よりも早く失われたソースデータへの悪意のある命令
気づきます。
汎用データベースを中心にこれらの保護を構築できます。ジェイベースが作る
それらは各アプリケーションに任せるのではなく、すべての書き込みの一部となります。
Jaybase は事実が真実かどうかを判断しません。認定代理人は引き続き、
悪い事実を書きます。 Jaybase は、そのアクションを可視化して修正可能な状態に保ちます。
Jaybase はイベントの線形チェーンを保存します。各イベントには何が起こったのか、誰が起こったのかが記録されます
実行、暗号化された JSON ペイロード、およびその前のイベント。ペイロードはそのまま
柔軟な;歴史のルールではそうではありません。
その Expected_root と安定した Idempotency-Key を使用してファクトを送信します。
Jaybase はアクターを取得し、イベントを暗号化してハッシュし、書き込み、
根を進めます。
同じ再試行では、元のイベントが返されます。古いルートと再利用されたキー
異なるコンテンツが返される 409 競合。
各イベント アドレスは、その前のイベントに依存します。古いコンテンツの変更を変更する
これに続くハッシュにより、ルートのオフホスト コピーが書き換えられたものや、
歴史を置き換えた。
修正、撤回、承認は新しいイベントであり、編集ではありません。ホストされた
API には履歴の更新パスや削除パスがなく、呼び出し元は独自のパスを選択できません
アイデンティティ。 1 つの書き込みプロセスがデータ ボリュームごとに書き込みをシリアル化します。多くのエージェント
はそのプロセスを使用できますが、Jaybase は分散型コンセンサス システムではありません。
前提条件: Linux ホスト

i 番目の Docker Compose、ポート 80 および 443 に到達可能、および
ホストのドメインを指す A/AAAA レコード。
cp .env.example .env
# .env を編集し、JAYBASE_DOMAIN を設定します。
jaybase-server init ./secrets
docker 構成 -d --build
ドッカー構成PS
カール https://jaybase.example.com/health/ready
イニシャライザは既存のシークレットを置き換えません。サーバーには
外部データキーとハッシュされた認証情報ファイル。
import JAYBASE_URL=https://jaybase.example.com
import JAYBASE_TOKEN= ' ライタートークン '
カール -fsS \
-H " 認可: ベアラー $JAYBASE_TOKEN " \
" $JAYBASE_URL /v1/root "
返されたルートを Expected_root として使用し、安定したべき等性キーを 1 つ選択します
論理演算の場合。新しいイベントの最初のイベントにのみ空の文字列を使用します。
データベース。
curl -fsS -X POST " $JAYBASE_URL /v1/events " \
-H " 認可: ベアラー $JAYBASE_TOKEN " \
-H " Content-Type: application/json " \
-H " 冪等キー: 事実-主連絡先-v1-7f3d2a " \
--データ ' {
"タイプ": "business.fact",
"entity_id": "01JOPAQUE8F3K2M7Q9R4T6V1WX",
"コマンド": "事実の主張",
"ペイロード": {"primary_contact": "エイダ ラブレス"},
"expected_root": "sha256:前の応答からのルート"
} '
読み取り、ページネーション、参照、スナップショット、および
管理エンドポイント。
ストア、エラー:= ジェイベース。 OpenStore ( ".jaybase" )
if err != nil { /* エラーを処理します */ }
ストアを延期します。閉じる ()
ルート、エラー:= ストア。 Append (jaybase.Context { Actor : "エージェント" }, jaybase.AppendOptions {
タイプ: "business.fact" 、コマンド: "factassert" 、ペイロード: fat 、
})
OpenStore は、キーとデータを同じ場所に配置するローカル開発の利便性を提供します。
運用プロセスでは OpenStoreWithDataKey を使用する必要があります。サーバーはこれを強制します。
1 つのプロセスが書き込み可能な各データ ボリュームを所有します。
Caddy は HTTPS を処理します。ベアラー資格情報は、 Reader 、 Writer 、または admin を提供します
アクセスします。
ペイロ

広告は保存時に暗号化され、データキーはボリュームの外部に保存されます
そしてスナップショット。
スナップショットはオフホストにコピーする必要があります。
コンテナーは、読み取り専用のルート ファイルシステムを使用して非ルートとして実行されます。
アーキテクチャ、セキュリティ、
API 、および実行前の操作ガイド
機密データを含む Jaybase。 Jaybase と統合するエージェントは次を使用する必要があります
llm.md を運営契約とします。
GOCACHE=/tmp/jaybase-gocache go test -race ./...
GOCACHE=/tmp/jaybase-gocache go vet ./...
ドッカー構成構成
docker build -t jaybase:test 。
ライセンス
AGPL-3.0以降。 「ライセンス」を参照してください。
Jaybase は、重要なビジネス ワークフローを実行する AI エージェントのための追加専用のファクト ストアです
AGPL-3.0ライセンス
セキュリティポリシー
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
ジェイベース v0.2.0
最新の
2026 年 7 月 21 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Jaybase is a append-only fact store for AI agents doing critical business workflows - kyle-visner/jaybase

Hey HN, I build custom AI agents for businesses, and one of the things I've noticed is that everyone wants to use AI to do real, critical work, but it's hard to trust it because sometimes agents can go a little off the rails. No one wants their agent deleting all their customer data or screwing up their books. So they tend to limit its use to searching, summarizing, and as a writing assistant. The problem is that most ways we store data, from notes & spreadsheets to databases, were designed for humans who work at and make mistakes at human speed, or for deterministic software systems. AI agents work non-deterministically at machine speed, which is a new challenge. That's why I made Jaybase. Jaybase is an append-only fact store for agentic workflows. It's based on a Merkle Directed Acyclic Graph, which is the same technology behind git and the blockchain, so every write is auditable, replayable, and reversible. If an agent goes off the rails and corrupts a bunch of data, it's easy to recover and easy to audit what went wrong. It also doesn't have a fixed schema, so as you add agents or your agents get more skills, it's natural to add more kinds of data. Happy to dive into any questions in the comments and would love the community’s feedback. Thanks!

GitHub - kyle-visner/jaybase: Jaybase is a append-only fact store for AI agents doing critical business workflows · GitHub
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
kyle-visner
/
jaybase
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
29 Commits 29 Commits .github .github cmd/ jaybase-server cmd/ jaybase-server deploy deploy docs docs secrets secrets server server .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md behavior_test.go behavior_test.go compose.yaml compose.yaml go.mod go.mod hosted_test.go hosted_test.go llm.md llm.md lock_fcntl.go lock_fcntl.go lock_other.go lock_other.go lock_unix.go lock_unix.go migration.go migration.go migration_test.go migration_test.go performance_test.go performance_test.go snapshot.go snapshot.go storage.go storage.go storage_test.go storage_test.go View all files Repository files navigation
Jaybase is an append-only fact store for AI agents trusted with critical
business data. Agents can add flexible JSON facts, but the hosted API cannot
rewrite or delete history. Every write is attributed, encrypted, safe to retry,
checked for stale state, and available for replay.
git clone https://github.com/kyle-visner/jaybase.git
cd jaybase
go install ./cmd/jaybase-server
jaybase-server init ./secrets
The initializer prints reader, writer, and admin tokens once. Save them in a
password manager, then continue to Run Jaybase .
Jaybase is for developers and small teams moving from read-only copilots to
agents that are allowed to operate. It fits accounting, operations, compliance,
approvals, and other work where agents write critical data, mistakes must stay
visible and correctable, and fact shapes evolve with the job.
Jaybase is designed for single-tenant systems: one organization, one trust
boundary, and one writer process per store. Many agents and applications can
share that store, including dashboards, internal tools, APIs, and automated
workflows. Jaybase is not meant to be the globally distributed, multi-tenant
backend for a web application.
Traditional databases assume deterministic application code owns every read and
write. Agents make judgment calls, retry uncertain work, and sometimes behave in
unexpected ways—at machine speed. A mutable database can turn one bad decision,
runaway loop, or malicious instruction into lost source data before anyone
notices.
You can build these protections around a general-purpose database. Jaybase makes
them part of every write instead of leaving them to each application.
Jaybase does not decide whether a fact is true. An authorized agent can still
write a bad fact; Jaybase keeps that action visible and correctable.
Jaybase stores a linear chain of events. Each event records what happened, who
did it, an encrypted JSON payload, and the event before it. Payloads stay
flexible; the history rules do not.
Submit a fact with that expected_root and a stable Idempotency-Key .
Jaybase derives the actor, encrypts and hashes the event, writes it, and
advances the root.
Identical retries return the original event. Stale roots and reused keys with
different content return 409 conflict .
Each event address depends on the event before it. Changing old content changes
the hashes that follow, so an off-host copy of the root can detect rewritten or
replaced history.
Corrections, retractions, and approvals are new events, never edits. The hosted
API has no update or delete path for history, and callers cannot choose their own
identity. One writer process serializes writes for each data volume; many agents
can use that process, but Jaybase is not a distributed consensus system.
Prerequisites: a Linux host with Docker Compose, ports 80 and 443 reachable, and
an A/AAAA record pointing a domain at the host.
cp .env.example .env
# Edit .env and set JAYBASE_DOMAIN.
jaybase-server init ./secrets
docker compose up -d --build
docker compose ps
curl https://jaybase.example.com/health/ready
The initializer will not replace existing secrets. The server requires an
external data key and hashed credential file.
export JAYBASE_URL=https://jaybase.example.com
export JAYBASE_TOKEN= ' the-writer-token '
curl -fsS \
-H " Authorization: Bearer $JAYBASE_TOKEN " \
" $JAYBASE_URL /v1/root "
Use the returned root as expected_root and choose one stable idempotency key
for the logical operation. Use an empty string only for the first event in a new
database.
curl -fsS -X POST " $JAYBASE_URL /v1/events " \
-H " Authorization: Bearer $JAYBASE_TOKEN " \
-H " Content-Type: application/json " \
-H " Idempotency-Key: fact-primary-contact-v1-7f3d2a " \
--data ' {
"type": "business.fact",
"entity_id": "01JOPAQUE8F3K2M7Q9R4T6V1WX",
"command": "fact assert",
"payload": {"primary_contact": "Ada Lovelace"},
"expected_root": "sha256:root-from-the-previous-response"
} '
See the API guide for reads, pagination, refs, snapshots, and
administrative endpoints.
store , err := jaybase . OpenStore ( ".jaybase" )
if err != nil { /* handle error */ }
defer store . Close ()
root , err := store . Append (jaybase. Context { Actor : "agent" }, jaybase. AppendOptions {
Type : "business.fact" , Command : "fact assert" , Payload : fact ,
})
OpenStore is a local-development convenience that co-locates the key and data.
Production processes must use OpenStoreWithDataKey ; the server enforces this.
One process owns each writable data volume.
Caddy handles HTTPS; bearer credentials provide reader , writer , or admin
access.
Payloads are encrypted at rest, with the data key stored outside the volume
and snapshots.
Snapshots should be copied off-host.
Containers run as non-root with a read-only root filesystem.
Read the architecture , security ,
API , and operations guides before running
Jaybase with sensitive data. Agents integrating with Jaybase should use
llm.md as their operating contract.
GOCACHE=/tmp/jaybase-gocache go test -race ./...
GOCACHE=/tmp/jaybase-gocache go vet ./...
docker compose config
docker build -t jaybase:test .
License
AGPL-3.0-or-later. See LICENSE .
Jaybase is a append-only fact store for AI agents doing critical business workflows
AGPL-3.0 license
Security policy
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Jaybase v0.2.0
Latest
Jul 21, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
