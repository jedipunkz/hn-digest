---
source: "https://docs.victoriametrics.com/victoriametrics/contributing/index.html"
hn_url: "https://news.ycombinator.com/item?id=49310440"
title: "VictoriaMetrics: AI Policy"
article_title: "VictoriaMetrics: Contributing"
author: "tosh"
captured_at: "2026-08-15T14:14:01Z"
capture_tool: "hn-digest"
hn_id: 49310440
score: 1
comments: 0
posted_at: "2026-08-15T13:36:11Z"
tags:
  - hacker-news
  - translated
---

# VictoriaMetrics: AI Policy

- HN: [49310440](https://news.ycombinator.com/item?id=49310440)
- Source: [docs.victoriametrics.com](https://docs.victoriametrics.com/victoriametrics/contributing/index.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T13:36:11Z

## Translation

タイトル: VictoriaMetrics: AI ポリシー
記事のタイトル: VictoriaMetrics: 貢献

記事本文:
ビクトリアメトリクス
ドキュメント
当社の製品
統合
グラファナ
VictoriaMetrics データソース
ストリーミングアグリゲーション
構成
Loki クエリを VictoriaLogs クエリに変換する方法
VictoriaLogs のリリース プロセス ガイダンス
データの取り込み
OpenTelemetry のセットアップ
クエリを実行する
Grafana での視覚化
ガイド
VictoriaMetrics Single を介した Kubernetes モニタリング
VictoriaMetrics Cluster を使用した Kubernetes モニタリング
VM オペレーターの使用を開始する
VictoriaMetrics コンポーネントをクラウド ストレージに接続する
VictoriaMetrics および VictoriaLogs で OpenTelemetry を使用する方法
vmagent のセットアップ - マルチテナントのリモート書き込みと OIDC
vmauth のセットアップ - Grafana および OIDC を使用したマルチテナント アクセス
VictoriaMetrics でメトリクスを削除または置換する方法
InfluxDB から VictoriaMetrics への移行
VictoriaMetrics クラスターを介した Kubernetes での HA モニタリングのセットアップ
VictoriaMetrics クラスター内の複数保持のセットアップ
VictoriaMetrics の複数地域セットアップ: 専用モニタリング
Victoria Logs を使用した OpenShift ログの収集
ヘッドランプ Kubernetes UI と VictoriaMetrics
vmalert と Grafana を使用したデータソース管理アラート
Kubernetes オペレーター
クイックスタート
認可とコンポーネントの公開
ヘルムチャート
ビクトリアログシングル
変更履歴
VictoriaLogs クラスター
変更履歴
VictoriaLogs コレクター
変更履歴
VictoriaTraces クラスター
変更履歴
VictoriaMetrics シングル
変更履歴
VictoriaMetrics クラスター
変更履歴
VictoriaMetrics エージェント
変更履歴
VictoriaMetrics アラート
変更履歴
VictoriaMetrics の異常
変更履歴
VictoriaMetrics 認証
変更履歴
VictoriaMetrics の分散
変更履歴
VictoriaMetrics ゲートウェイ
変更履歴
VictoriaMetrics K8 スタック
変更履歴
VictoriaMetrics オペレーター
変更履歴
VictoriaMetrics オペレーター CRD
変更履歴
VictoriaMetrics Common
変更履歴
VictoriaTraces シングル
変更履歴
VictoriaLogs マルチレベル
変更履歴
VictoriaMetrics クラウド
始めましょう
概要
アカウント

管理
登録とトライアル
データの探索
VictoriaMetrics を探索する
CloudWatch - エージェントレスの AWS モニタリング
VictoriaMetrics Cloud を使用した Kubernetes モニタリング
VictoriaMetrics Cloud の Alertmanager と VMAlert をセットアップする
vmalert と VictoriaMetrics Cloud を使用したアラート
ガイド
異常の検出と警告のセットアップ
VictoriaMetrics が好きで、貢献したい場合は、素晴らしいことです。
VictoriaMetrics コミュニティ Slack に参加する (Slack の招待者)
とSlackチャンネル
)
そしてそこで他のコミュニティメンバーを助けます。
VictoriaMetrics GitHub で問題、機能リクエスト、質問を提出してください
。
改善する
VictoriaMetrics ドキュメント
。アップデート方法をご覧ください
ドキュメント
。
さまざまなチャネルを通じて VictoriaMetrics についての情報を広める: カンファレンス トーク
ブログ投稿、記事、ケーススタディ
Hacker News、Twitter、LinkedIn、Reddit、Facebook などでのコメント
同僚と経験を共有する。
経営陣に署名を説得する
エンタープライズ契約
VictoriaMetricsを使用して。
新しい問題を作成するときは、重複を作成しないようにしてください。 GitHub 検索を使用して、同様の問題がすでに存在するかどうかを確認します。
新しい問題は英語で書かれ、問題と問題が存在する環境についての簡潔な説明が含まれている必要があります。
回避策や代替ソリューションがある可能性があるため、説明に特定の使用例を含めることを強く希望します。
貢献できる問題を探すときは、常にバグに取り組むことを優先してください
機能強化の代わりに
。
他の人の質問を助ける
も貢献です。
ドキュメントに貢献したい場合
、お願いします
を読んでください
ガイドライン
。
ラベルを使用します
GitHub の問題を分類します。次のラベルがあります。
コンポーネント ラベル: vmalert、vmagent など。特定のコンポーネントに関連する場合は、このラベルを問題に追加します。
問題の種類: バグ、機能強化、q

質問。
エンタープライズ、耳鼻咽喉科の機能に関連する問題を担当
「詳細が必要」。問題の作成者による詳細な説明が必要な問題に割り当てられます。
たとえば、チケットの説明に基づいて報告されたバグを再現できなかった場合は、追加の質問をします。
問題を再現し、「詳細情報が必要」というラベルを追加するのに役立つ質問。このラベルは他のメンテナーに役立ちます
この問題が忘れ去られたわけではないことを理解するには、ユーザーからのフィードバックを待ちます。
completed 、コード変更が必要な問題に割り当てられ、それらの変更は上流にマージされましたが、まだリリースされていません。
リリースが作成されると、メンテナはラベル付きの問題をすべて確認し、新しいリリースについてコメントを残して、問題をクローズします。
vmui 、に関連する問題に割り当てられています
vmui
または
VictoriaLogs ウェブイム
バグ修正または機能拡張を実装するには、対応するリポジトリにプル リクエストを送信する必要があります
。
プルリクエストは以下に準拠する必要があります
VictoriaMetrics の開発目標
。
レビュー担当者が変更を修正できなくなるため、PR の作成には master ブランチを使用しないでください。
すべてのコミットには署名が必要です
。
どのコンポーネントが変更されたかを示すために、プル リクエストのタイトルの前に <dir>/<component>: を付ける必要があります (例: app/vmalert: fix... )。
プル リクエストの説明には、何が行われたのか、なぜそれが必要なのか、どのような目的で行われたのかを明確かつ簡潔に説明する必要があります。
レビュー担当者が変更とその影響をすぐに理解できるように、明確な言葉を使用してください。
変更に関連する問題へのリンク (存在する場合)。 PR が問題を解決する場合は修正 [問題のリンク] を使用し、参照用に [問題のリンク] に関連するものを使用します。
変更が効果的であることを証明するテスト。テストは、重要な新機能または重要な変更に対して期待されます。
メンテナが明示的に同意しない限り、バグ修正にはテストを含める必要があります。
見てください

彼のスタイルガイド
テスト用。参照
このセクション
テストの実行方法については。
プル リクエストの範囲を課題の外に拡張したり、無関係な変更を加えたりしないようにしてください。
ドキュメントを更新する
必要に応じて。たとえば、新しいフラグを追加したり、既存のフラグや機能の動作を変更したりします。
これらの変更をドキュメントに反映する必要があります。新しい機能については、{{% available_from "#" %}} ショートコードをドキュメントに追加します。
後で実際のリリース バージョンに自動的に置き換えられます。
の行
変更履歴
変更とそれに関連する問題について何らかの方法で言及する
他の読者が完全な文脈を知らなくても、それは明らかです。
ベンダー提供のパッケージが VictoriaMetrics GitHub 組織からのものである場合でも、/vendor フォルダー内のコードを手動で変更することは避けてください。
たとえば、VictoriaLogs ベンダーは VictoriaMetrics の /lib フォルダーの下にあるパッケージをベンダーし、VictoriaTraces は VictoriaLogs の /lib/logstorage パッケージをベンダーします。
まず、上流リポジトリにプル リクエストを送信します。その後、別のプル リクエストを開いて、ダウンストリーム リポジトリ内のベンダー フォルダーのバージョンを更新できます。一般的なパッケージの場合、ベンダーのパッケージは次のコマンドで更新できます: go get <dependency>@vX.Y.Z 。
VictoriaMetrics パッケージの場合は、 go get <dependency>@canonical_commit_hash を使用します。
最後に、 go mod tiny および go modvendor を実行して、 go.mod 、 go.sum 、および /vendor を更新します。
その問題に関して最も優れた専門知識を持っていると思われるレビュー担当者に Ping を送ります。
プルリクエストの良い例を見る
。
投稿やコードに取り組む際には、AI ツールを自由に使用できます。
ドキュメント、問題、その他何でも。かどうかを開示する必要はありません。
どのように使ったか。
AI の助けの有無にかかわらず、送信した変更についてはお客様の責任となります。
コードベースとすべてを理解するために努力してください。

プルリクエストの変更、
送信する前に AI のスロップをクリーンアップします。 AI を使用して自動化しないでください。
メンテナーへの返答。
私たちは、制作方法に関係なく、その品質に基づいて投稿を審査します。あ
低品質または未レビューの AI 出力のように見えるプル リクエストまたは問題
壊れた変更は、詳細なレビューやトリアージなしに終了される可能性があります。
プル リクエストをマージする人は、以下の要件を満たす責任があります。
PR が満足していることを確認してください
プルリクエストのチェックリスト
、
少なくとも 1 人のレビュー担当者によって承認されており、すべての CI チェックは緑色です。
変化を評価するために最善を尽くしてください。可能であれば、ローカルでテストしてください。
マージしたら、関連するすべてのブランチにわたって変更を厳選してください。
該当する場合は、変更を厳選して、
LTSリリースライン
PR コメントで、何が厳選されたのか、何が選ばれなかったのかについて言及します。
何が変更され、いつ変更されるのかについての意味のあるメッセージを含めて、関連する問題を更新します。
解放されました。これにより、ユーザーは PR を読まなくても変更を理解できます。
関連する問題に完了ラベルを追加します。
リリースが行われるまで、関連するチケットを閉じないでください。チケットが GitHub またはユーザーによって自動クローズされた場合は、チケットを再度開きます。
KISS 設計原則に従っている限り、サードパーティのプル リクエストを受け入れます。
。
シンプルなコードとアーキテクチャを好みます。
マジック コードや派手なアルゴリズムは避けてください。
最適化を適用すると、次の場合にのみコードが理解しにくくなります。
プロファイリング
パフォーマンスとスケーラビリティの大幅な向上、または RAM 使用量の大幅な削減を示しています。
プロファイリングは Go ベンチマークで実行する必要があります
そして本番環境のワークロードについても同様です。
外部への大きな依存関係を避ける
。
分散システム内の可動部分の数を最小限に抑えます。
クラスターの可用性、一貫性、パフォーマンス、またはデバッグ可能性を損なう可能性がある自動決定は避けてください。
アド

KISS 原則に準拠することで、結果として得られるコードとアーキテクチャが簡素化され、より幅広いユーザーがレビュー、理解、デバッグできるようになります。
KISSのせいで、
VictoriaMetrics のクラスター バージョン
には、分散コンピューティングで人気のある次の「機能」がありません。
脆弱なゴシッププロトコル。 Thanos での失敗した試みを参照してください
。
理解するのが難しく、適切に実装するのが難しい Paxos プロトコル
。
複雑なレプリケーション スキーム。予期しないエッジ ケースで機能不全に陥る可能性があります。参照
レプリケーションドキュメント
詳細については。
ストレージ ノード間のデータの自動再シャッフル。これにより、クラスターのパフォーマンスと可用性が損なわれる可能性があります。
クラスターの自動サイズ変更。構成が不適切な場合、多額の費用がかかる可能性があります。
クラスター内の新しいノードを自動的に検出して追加します。これにより、開発クラスターと本番クラスターの間でデータが混在する可能性があります:)
リーダーの自動選出により、ネットワーク エラーによるスプリット ブレイン災害が発生する可能性があります。
プル リクエストを送信する前に、次の一連のチェックとテストを実行することをお勧めします。

## Original Extract

VictoriaMetrics
Docs
Our Products
Integrations
Grafana
VictoriaMetrics Datasource
Streaming aggregation
Configuration
How To Convert Loki Queries to VictoriaLogs Queries
Release Process Guidance for VictoriaLogs
Data ingestion
OpenTelemetry setup
Querying
Visualization in Grafana
Guides
Kubernetes monitoring via VictoriaMetrics Single
Kubernetes monitoring with VictoriaMetrics Cluster
Getting started with VM Operator
Connecting VictoriaMetrics components to cloud storage
How to use OpenTelemetry with VictoriaMetrics and VictoriaLogs
Setup vmagent - Multi-Tenant remote write & OIDC
Setup vmauth - Multi-Tenant Access with Grafana & OIDC
How to delete or replace metrics in VictoriaMetrics
Migrate from InfluxDB to VictoriaMetrics
HA monitoring setup in Kubernetes via VictoriaMetrics Cluster
Multi Retention Setup within VictoriaMetrics Cluster
VictoriaMetrics Multi-Regional Setup: Dedicated Monitoring
Collecting OpenShift logs with Victoria Logs
Headlamp Kubernetes UI and VictoriaMetrics
Datasource-Managed Alerts with vmalert and Grafana
Kubernetes Operator
QuickStart
Authorization and exposing components
Helm Charts
VictoriaLogs Single
CHANGELOG
VictoriaLogs Cluster
CHANGELOG
VictoriaLogs Collector
CHANGELOG
VictoriaTraces Cluster
CHANGELOG
VictoriaMetrics Single
CHANGELOG
VictoriaMetrics Cluster
CHANGELOG
VictoriaMetrics Agent
CHANGELOG
VictoriaMetrics Alert
CHANGELOG
VictoriaMetrics Anomaly
CHANGELOG
VictoriaMetrics Auth
CHANGELOG
VictoriaMetrics Distributed
CHANGELOG
VictoriaMetrics Gateway
CHANGELOG
VictoriaMetrics K8s Stack
CHANGELOG
VictoriaMetrics Operator
CHANGELOG
VictoriaMetrics Operator CRDs
CHANGELOG
VictoriaMetrics Common
CHANGELOG
VictoriaTraces Single
CHANGELOG
VictoriaLogs Multilevel
CHANGELOG
VictoriaMetrics Cloud
Get Started
Overview
Account Management
Registration and Trial
Exploring Data
Exploring VictoriaMetrics
CloudWatch - Agentless AWS monitoring
Kubernetes Monitoring with VictoriaMetrics Cloud
Setup Alertmanager & VMAlert for VictoriaMetrics Cloud
Alerting with vmalert and VictoriaMetrics Cloud
Guides
Anomaly Detection and Alerting Setup
If you like VictoriaMetrics and want to contribute, then it would be great:
Joining VictoriaMetrics community Slack ( Slack inviter
and Slack channel
)
and helping other community members there.
Filing issues, feature requests and questions at VictoriaMetrics GitHub
.
Improving
VictoriaMetrics docs
. See how to update our
documentation
.
Spreading the word about VictoriaMetrics via various channels: conference talks
blogposts, articles and case studies
comments at Hacker News, Twitter, LinkedIn, Reddit, Facebook, etc.
experience sharing with colleagues.
Convincing your management to sign
Enterprise contract
with VictoriaMetrics.
When making a new issue, make sure to create no duplicates. Use GitHub search to find whether similar issues exist already.
The new issue should be written in English and contain a concise description of the problem and the environment where it exists.
We’d very much prefer to have a specific use-case included in the description, since it could have workaround or alternative solutions.
When looking for an issue to contribute, always prefer working on bugs
instead of enhancements
.
Helping other people with their questions
is also a contribution.
If you would like to contribute to documentation
, please
read the
guideline
.
We use labels
to categorize GitHub issues. We have the following labels:
A component label: vmalert, vmagent, etc. Add this label to the issue if it is related to a specific component.
An issue type: bug , enhancement , question .
enterprise , assigned to issues related to ENT features
need more info , assigned to issues that require elaboration from the issue creator.
For example, if we weren’t able to reproduce the reported bug based on the ticket description then we ask additional
questions which could help to reproduce the issue and add need more info label. This label helps other maintainers
to understand that this issue wasn’t forgotten but waits for the feedback from the user.
completed , assigned to issues that required code changes and those changes were merged to upstream, but not released yet.
Once a release is made, maintainers go through all labeled issues, leave a comment about the new release, and close the issue.
vmui , assigned to issues related to
vmui
or
VictoriaLogs webui
Implementing a bugfix or enhancement requires sending a pull request to the corresponding repository
.
The pull request must conform to
VictoriaMetrics development goals
.
Don’t use master branch for making PRs, as it makes it impossible for reviewers to modify the changes.
All commits need to be signed
.
Pull request title should be prefixed with <dir>/<component>: to show what component has been changed, i.e. app/vmalert: fix... .
Pull request description should contain a clear and concise description of what was done, why it is needed and for what purpose.
Use clear language, so reviewers can quickly understand the change and its impact.
A link to the issue(s) related to the change, if any. Use Fixes [issue link] if the PR resolves the issue, or Related to [issue link] for reference.
Tests proving that the change is effective. Tests are expected for non-trivial new functionality or non-trivial modifications.
Bug fixes must include tests unless a maintainer explicitly agrees otherwise.
See this style guide
for tests. See
this section
for how to run tests.
Try to not extend the scope of the pull requests outside the issue, do not make unrelated changes.
Update docs
if needed. For example, adding a new flag or changing the behavior of existing flags or features
requires reflecting these changes in the documentation. For new features add {{% available_from "#" %}} shortcode to the documentation.
It will be later automatically replaced with an actual release version.
A line in the
changelog
mentioning the change and related issue in a way
that would be clear to other readers even if they don’t have the full context.
Avoid modifying code in the /vendor folder manually, even when the vendored package originates from the VictoriaMetrics GitHub organization.
For instance, VictoriaLogs vendors packages under the /lib folder from VictoriaMetrics, and VictoriaTraces vendors the /lib/logstorage package from VictoriaLogs.
Submit a pull request to the upstream repository first. Afterward, a separate pull request can be opened to update the version of the vendored folder in the downstream repository. For common packages, the vendored package can be updated with this command: go get <dependency>@vX.Y.Z .
For VictoriaMetrics packages, use go get <dependency>@canonical_commit_hash .
Finally, run go mod tidy and go mod vendor to update go.mod , go.sum , and /vendor .
Ping reviewers who you think have the best expertise on the matter.
See a good example of a pull request
.
You are free to use any AI tools when working on a contribution, on code,
documentation, issues, or anything else. You do not need to disclose whether or
how you used them.
With or without the help of AI, you are responsible for the changes you submit.
Take the effort to understand the code base and every change in your pull request,
and clean up any AI slop before sending it. Do not use AI to automate your
responses to maintainers.
We review contributions on their quality, regardless of how they were produced. A
pull request or issue that looks like unreviewed AI output, with low-quality or
broken changes, may be closed without a detailed review or triage.
The person who merges the Pull Request is responsible for satisfying the requirements below:
Make sure that PR satisfies
Pull Request checklist
,
it is approved by at least one reviewer, all CI checks are green.
Try doing your best at assessing the changes. If possible, test them locally.
Once merged, make sure to cherry-pick the changes across all related branches.
If applicable, cherry-pick the change to
LTS release lines
and mention in the PR comment what was or wasn’t cherry-picked.
Update related issues with a meaningful message of what has changed and when it will be
released. This helps users to understand the change without reading the PR.
Add label completed to related issues.
Do not close related tickets until the release is made. If the ticket was auto-closed by GitHub or a user - re-open it.
We are open to third-party pull requests provided they follow KISS design principle
.
Prefer simple code and architecture.
Avoid magic code and fancy algorithms.
Apply optimizations, which make the code harder to understand only if
profiling
shows significant improvements in performance and scalability or significant reduction in RAM usage.
Profiling must be performed on Go benchmarks
and on production workload.
Avoid big external dependencies
.
Minimize the number of moving parts in the distributed system.
Avoid automated decisions, which may hurt cluster availability, consistency, performance or debuggability.
Adhering to the KISS principle, simplifies the resulting code and architecture so it can be reviewed, understood and debugged by a wider audience.
Due to KISS ,
cluster version of VictoriaMetrics
has none of the following “features” popular in distributed computing:
Fragile gossip protocols. See failed attempt in Thanos
.
Hard-to-understand-and-implement-properly Paxos protocols
.
Complex replication schemes, which may go nuts in unforeseen edge cases. See
replication docs
for details.
Automatic data reshuffling between storage nodes, which may hurt cluster performance and availability.
Automatic cluster resizing, which may cost you a lot of money if improperly configured.
Automatic discovering and addition of new nodes in the cluster, which may mix data between dev and prod clusters :)
Automatic leader election, which may result in split brain disaster on network errors.
We recommend running the following sequence of checks and tests before submitting a pull request:
