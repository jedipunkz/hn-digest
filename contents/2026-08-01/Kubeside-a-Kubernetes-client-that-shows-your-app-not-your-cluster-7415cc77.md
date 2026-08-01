---
source: "https://github.com/dynaum/kubeside"
hn_url: "https://news.ycombinator.com/item?id=49139573"
title: "Kubeside – a Kubernetes client that shows your app, not your cluster"
article_title: "GitHub - dynaum/kubeside: A Kubernetes client scoped to the developer, not the cluster operator. Your app, not your cluster. · GitHub"
author: "dynaum"
captured_at: "2026-08-01T23:48:39Z"
capture_tool: "hn-digest"
hn_id: 49139573
score: 1
comments: 0
posted_at: "2026-08-01T23:22:12Z"
tags:
  - hacker-news
  - translated
---

# Kubeside – a Kubernetes client that shows your app, not your cluster

- HN: [49139573](https://news.ycombinator.com/item?id=49139573)
- Source: [github.com](https://github.com/dynaum/kubeside)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T23:22:12Z

## Translation

タイトル: Kubeside – クラスターではなくアプリを表示する Kubernetes クライアント
記事のタイトル: GitHub - dynaum/kubeside: クラスター オペレーターではなく、開発者を対象とした Kubernetes クライアント。クラスターではなくアプリです。 · GitHub
説明: クラスター オペレーターではなく、開発者を対象とする Kubernetes クライアント。クラスターではなくアプリです。 - ダイナウム/クベサイド

記事本文:
GitHub - dynaum/kubeside: クラスター オペレーターではなく、開発者を対象とした Kubernetes クライアント。クラスターではなくアプリです。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ダイナウム
/
クベサイド
公共
通知
nを変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
81 コミット 81 コミット .github/ workflows .github/ workflows cmd/ kubeside cmd/ kubeside docs docs 内部 内部スクリプト スクリプト Web Web .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CLAUDE.md CLAUDE.md LAUNCH.md LAUNCH.md ライセンス ライセンス通知 通知 README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ReplicaSet について考えることなく、すべてのクラスターにわたってアプリを実行できます。
Kubernetes クライアントのスコープは、オペレーターではなく、アプリを出荷する開発者に限定されます。
クラスターを実行している人。 1 つのバイナリ、データベース、エージェント、セットアップ手順はありません。
ドキュメント ·
インストール・
なぜ存在するのか
すべての Kubernetes UI は API ツリーをミラーリングします。リソースの種類を選択し、参照します。
インスタンスの場合は、スイッチャーからクラスターを選択します。それが正しい形です
クラスターを実行する人。発送する人にとっては間違った形状です
アプリ。
開発者は ReplicaSet ではなくサービスを考慮します。彼らの作業単位は 1 つのアプリです
QA、STG、および製品全体にわたり、それらの質問は過去のものよりも頻繁に発生します。
live: 何が変更されたか、誰が変更したか、コンテナは実際に何を取得したか。
今日それらのいずれかに答えることは、3 つの端末、レプリカごとに 1 つのタブ、そして 1 つのタブを意味します。
推測します。
したがって、ダッシュボードは kubectl のランチャーとなり、4 つの質問が残ります。
答えられていない。難しいからではなく、誰もその画面を作っていないからです
彼ら。
kubeside はこれら 4 つに正確に答えますが、意図的に何も答えることを拒否しています
それ以外の場合:
すべてのポッドにわたって一度にログには何が表示されますか?
コンテナーは実際にどのような構成を受け取りましたか?
QA、STG、製品全体を横並びにします。
醸造インストール dynaum/tap/kubeside
クベサイド
macOS、Linux、および Windows のバイナリは
リリースページ。ファイルは 1 つ、いいえ

ランタイム、クラスター内にインストーラー、エージェントはありません。 UI は
バイナリ。
セットアップ手順もありません。 kubeside はマシン上にすでに存在する kubeconfig を読み取ります。
すべてのコンテキストをロードし、実行資格情報プラグインをネイティブに実行します。 kubectl の場合
作品、クベサイド作品。
最初のリリースにタグが付けられるまで、それをビルドします。
npm --prefix web ci && npm --prefix web run build # UI が埋め込まれているため、最初に実行されます
ビルドに行く ./cmd/kubeside && ./kubeside
完全な手順: インストール ガイド 。
所有するすべてのアプリ、kubeconfig が到達するすべてのクラスターで、アプリとしてグループ化される
ReplicaSet としてではなく。健康は派生しており、その行には何が派生したのかが示されています。
pod search-indexer-75f4d は ImagePullBackOff にあり、赤い点を上回ります。の
GROUPED BY 列には、各行を生成したルールの名前が付けられるため、次のようなリストが作成されます。
間違っていると、それが間違っているように見える理由がわかります。
到達できないクラスターは、その行にその旨を示します。決して貢献しない
他の人のリストに沈黙する。
タイムラインは、Kubernetes がすでに保持している履歴からオンデマンドで再構築されます。
ReplicaSet、ControllerRevisions、Helm リリース シークレット、ポッドの終了状態、
イベントは依然として API サーバーの TTL 内にあります。変更にはアクターが読み込まれます
manageFields 、つまり誰も認めていない kubectl がロールアウトの隣に表示されます
それが引き起こしたのです。
その知識が終わるところに、地平線が描かれます。
レプリカセット ·reviationHistoryLimit によって削除された古いロールアウト。リビジョン 11 はクラスターがまだ保持している中で最も古いものです
イベント · 古いイベントは API サーバーから期限切れになりました。デフォルトでは約 1 時間保持されます。
セッション・クベサイドはここで視聴を開始しました。これより前の内容はクラスター自体の履歴から取得されます
あなたの役割が読み取ることができないソースは、ラベル付きギャップになります。空の軸は決してありません
静かな期間として表示されます。
コンテナが実際に受け取ったもの
env 、 envFrom 、 configMapKeyR を通じて解決される環境

エフ、
SecretKeyRef 、下向き API、およびマウントされたボリューム。各キーは
そのソースと以前のリビジョンとの比較。
シークレット値は、kubeside が取得したことがないため、マスクされています。明らかにする人は尋ねます
クラスターでその特定の Secret を取得できるかどうか、および答えが次のとおりであるかどうか
いいえ、コントロールは無効になり、消えるのではなく動詞に名前が付けられます。
アプリごとに 1 行、環境ごとに 1 列。バージョンの製品にはそのステージングがあります
誰もバックマージしなかったホットフィックスが最悪であるため、トップには表示されません。
この画面にあるもの。異なるダイジェストを持つ同じタグが呼び出されます。アン
読み取れない環境では「読み取れません」と表示されますが、これは
不在。
ディスクには何も書き込まれません。データベースもキャッシュファイルも履歴もありません
ディレクトリ。サーバーを停止すると、何も残りません。これは製品の
中心的な賭け: 歴史には保存が必要だと誰もが思い込んでいるため、誰も歴史を組み立てようとはしない
Kubernetes にはすでに履歴が保存されています。
知識の欠如は、物の欠如ではありません。取ることができなかった指標
は利用不可として報告されますが、ゼロとして報告されることはありません。リストできなかった種類に名前が付けられています。あ
覗けないウィンドウは空白ではなくハッチングで表示されます。
許可がないためにコントロールが非表示になることはありません。無効になっており、名前が付けられています
クラスターが拒否した動詞。セキュリティ境界はクラスターの RBAC です。
SelfSubjectAccessReview で解決しました。上部のガードレールは人間工学に基づいたもので、
両方の答えが一緒に伝わります。
資格情報はマシン上に残ります。 kubeconfig は読み取り専用の入力です。
編集されており、決してコピーされていません。ブラウザは 127.0.0.1 と通信しますが、127.0.0.1 とは通信しません。
アピサーバー。
ノードビューはありません。 PersistentVolume の参照はありません。 RBAC エディターはありません。 CRDブラウザはありません。いいえ
ヘルムチャート管理。費用の報告はありません。トポロジーグラフはありません。 YAMLエディターはありません
読み取り専用ビューアを超えています。プラグインシステムはありません。
それぞれは誰かに属する正当なニーズです

他人のツール。いずれかの発送
彼らはこれを汎用のダッシュボードに変えます。
もう十分です。
リリースされました。アプリのリスト、ワークロード全体のログ、再構築されたタイムライン、
解決された構成、環境間の差分、プロモーション マトリックス、
port-forward、コマンド パレット、prod ガードレール、および exec が構築され、テストされ、
そして実際のクラスターに対して駆動されます。
v1.0.0 より前では、書き込みパス、ローカル サーバー、および
この README が最も主張している 2 つの主張です。見つかったものはすべて修正されています。
ポートフォワード パスは、ダイヤルする前にクラスタとガードに問い合わせるようになりました。
パーミッションキャッシュは、尋ねられていない質問に答えることができなくなり、セッションが
トークンはブラウザの履歴に残るのではなく、一度引き渡されます。
秘密のフィンガープリントはキー化されているため、推測された値を確認することはできません。
完全なサイトは次のとおりです
kubeside.dynaum.com 。
ガイド — インストールと実行 ·
最初の実行 ·
アプリリスト ·
ログ ·
タイムライン ·
解決された構成 ·
プロモーションとドリフト ·
ポートフォワード、exec、ガードレール ·
キーボード・
設定ファイル ·
権限 ·
物が足りないとき
設計ノート — 問題 ·
ペルソナ ·
製品仕様・
マルチクラスタ ·
建築・
ロードマップ
設計書は仕様書であり、事前にレビューされます。
コードが存在していました。変更内容が矛盾する場合、ドキュメントは同じ内容で更新されます。
静かに発散するのではなく、コミットしてください。
テストは実装の前に行われます。テストに行く ./... 、
npm --prefix web run test と Playwright スクリーンショット ゲートはすべて CI で実行されます。
Linux、macOS、Windows 上で。
Kubernetes クライアントのスコープはクラスター オペレーターではなく開発者です。クラスターではなくアプリです。
kubeside.dynaum.com/ リソース
Readme Apache-2.0 ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
F

オーターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Kubernetes client scoped to the developer, not the cluster operator. Your app, not your cluster. - dynaum/kubeside

GitHub - dynaum/kubeside: A Kubernetes client scoped to the developer, not the cluster operator. Your app, not your cluster. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
dynaum
/
kubeside
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
81 Commits 81 Commits .github/ workflows .github/ workflows cmd/ kubeside cmd/ kubeside docs docs internal internal scripts scripts web web .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CLAUDE.md CLAUDE.md LAUNCH.md LAUNCH.md LICENSE LICENSE NOTICE NOTICE README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Your apps, across every cluster, without thinking in ReplicaSets.
A Kubernetes client scoped to the developer who ships the app, not the operator
who runs the cluster. One binary, no database, no agent, no setup step.
Documentation ·
Install ·
Why it exists
Every Kubernetes UI mirrors the API tree: pick a resource kind, browse
instances, pick a cluster from a switcher. That is the right shape for the
person who runs the cluster. It is the wrong shape for the person who ships an
app.
A developer thinks in services, not ReplicaSets. Their unit of work is one app
across qa, stg, and prod, and their questions are historical more often than
live: what changed, who changed it, what did the container actually get.
Answering any of those today means three terminals, a tab per replica, and a
guess.
So the dashboards end up as launchers for kubectl , and four questions stay
unanswered. Not because they are hard, but because nobody built the screen for
them.
kubeside answers exactly those four, and deliberately refuses to answer anything
else:
What do the logs say, across every pod at once?
What configuration did the container actually receive?
Across qa, stg, and prod, side by side.
brew install dynaum/tap/kubeside
kubeside
macOS, Linux, and Windows binaries are on the
releases page . One file, no
runtime, no installer, no agent in your cluster. The UI is embedded in the
binary.
No setup step either. kubeside reads the kubeconfig already on your machine,
loads every context, and runs exec credential plugins natively. If kubectl
works, kubeside works.
Until the first release is tagged, build it:
npm --prefix web ci && npm --prefix web run build # the UI is embedded, so it goes first
go build ./cmd/kubeside && ./kubeside
Full instructions: the install guide .
Every app you own, in every cluster your kubeconfig reaches, grouped as apps
rather than as ReplicaSets. Health is derived and the row says what derived it:
pod search-indexer-75f4d is in ImagePullBackOff beats a red dot. The
GROUPED BY column names the rule that produced each row, so a list that looks
wrong tells you why it looks wrong.
A cluster that cannot be reached says so in its own row. It never contributes
silence to somebody else's list.
The timeline is reconstructed on demand from history Kubernetes already keeps:
ReplicaSets, ControllerRevisions, Helm release secrets, pod termination states,
and events still inside the apiserver's TTL. Changes carry an actor read from
managedFields , so the kubectl nobody admits to shows up next to the rollout
it caused.
Where its knowledge ends, it draws the horizon:
replicaset · older rollouts pruned by revisionHistoryLimit; revision 11 is the oldest the cluster still holds
event · older events have expired from the apiserver, which keeps roughly an hour by default
session · kubeside started watching here; anything before this comes from the cluster's own history
A source your role cannot read becomes a labelled gap. An empty axis is never
rendered as a quiet period.
What the container actually received
Environment resolved through env , envFrom , configMapKeyRef ,
secretKeyRef , the downward API, and mounted volumes, each key attributed to
its source and compared against the previous revision.
Secret values are masked because kubeside never fetched them. Revealing one asks
your cluster whether you may get that specific Secret, and if the answer is
no, the control is disabled and names the verb rather than disappearing.
One row per app, one column per environment. A version prod has that staging
does not floats to the top, because a hotfix nobody back-merged is the worst
thing on this screen. Same tag with a different digest is called out. An
environment you cannot read says not readable , which is not the same as
absent .
Nothing is written to disk. No database, no cache file, no history
directory. Stop the server and nothing is left behind. This is the product's
central bet: everyone assumes history needs storage, so nobody assembles the
history Kubernetes already keeps.
Absence of knowledge is not absence of a thing. A metric it could not take
is reported as unavailable, never as zero. A kind it could not list is named. A
window it cannot see into is hatched, not blank.
A control is never hidden for lack of permission. It is disabled and names
the verb the cluster refused. The security boundary is your cluster's RBAC,
resolved with SelfSubjectAccessReview; the guardrails on top are ergonomics, and
both answers travel together.
Credentials stay on your machine. The kubeconfig is read-only input, never
edited and never copied. The browser talks to 127.0.0.1 and never to an
apiserver.
No node view. No PersistentVolume browsing. No RBAC editor. No CRD browser. No
Helm chart management. No cost reporting. No topology graph. No YAML editor
beyond a read-only viewer. No plugin system.
Each is a legitimate need belonging to somebody else's tool. Shipping any of
them turns this into a general-purpose dashboard, which is the thing there is
already enough of.
Released. The app list, whole-workload logs, the reconstructed timeline,
resolved configuration, cross-environment diff, the promotion matrix,
port-forward, the command palette, prod guardrails, and exec are built, tested,
and driven against real clusters.
Before v1.0.0 a security review went over the write paths, the local server, and
the two claims this README makes loudest. Everything it found is fixed: the
port-forward path now asks the cluster and the guard before it dials, the
permission cache can no longer answer a question it was not asked, the session
token is handed over once instead of living in your browser history, and the
secret fingerprint is keyed so it cannot confirm a guessed value.
The full site is at
kubeside.dynaum.com .
Guide — install and run ·
first run ·
the app list ·
logs ·
the timeline ·
resolved configuration ·
promotion and drift ·
port-forward, exec, guardrails ·
keyboard ·
the config file ·
permissions ·
when things are missing
Design notes — the problem ·
personas ·
product spec ·
multi-cluster ·
architecture ·
roadmap
The design documents are the specification, and they were reviewed before any
code existed. If a change contradicts one, the document gets updated in the same
commit rather than quietly diverging.
Tests come before implementation. go test ./... ,
npm --prefix web run test , and the Playwright screenshot gate all run in CI,
on Linux, macOS, and Windows.
A Kubernetes client scoped to the developer, not the cluster operator. Your app, not your cluster.
kubeside.dynaum.com/ Resources
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
