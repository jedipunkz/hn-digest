---
source: "https://github.com/stern/stern"
hn_url: "https://news.ycombinator.com/item?id=48777152"
title: "Stern: Multi pod and container log tailing for Kubernetes"
article_title: "GitHub - stern/stern: ⎈ Multi pod and container log tailing for Kubernetes -- Friendly fork of https://github.com/wercker/stern · GitHub"
author: "theanonymousone"
captured_at: "2026-07-03T17:13:48Z"
capture_tool: "hn-digest"
hn_id: 48777152
score: 1
comments: 0
posted_at: "2026-07-03T16:54:28Z"
tags:
  - hacker-news
  - translated
---

# Stern: Multi pod and container log tailing for Kubernetes

- HN: [48777152](https://news.ycombinator.com/item?id=48777152)
- Source: [github.com](https://github.com/stern/stern)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T16:54:28Z

## Translation

タイトル: Stern: Kubernetes のマルチ ポッドとコンテナーのログ テーリング
記事のタイトル: GitHub - stern/stern: ⎈ Kubernetes のマルチ ポッドとコンテナのログ テーリング -- https://github.com/wercker/stern · GitHub のフレンドリー フォーク
説明: ⎈ Kubernetes のマルチポッドとコンテナーのログテーリング -- https://github.com/wercker/stern - stern/stern のフレンドリーフォーク

記事本文:
GitHub - stern/stern: ⎈ Kubernetes のマルチ ポッドとコンテナのログ テーリング -- https://github.com/wercker/stern · GitHub のフレンドリー フォーク
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
船尾
/
船尾
公共
通知

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
316 コミット 316 コミット .github .github cmd cmd ハック ハック スターン スターン .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum main.go main.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
生産終了したヴェルカー/スターンのフォーク
Stern を使用すると、Kubernetes 上の複数のポッドと複数のコンテナを追跡できます
ポッド内。各結果は、デバッグを迅速化するために色分けされています。
クエリは、次の形式の正規表現または Kubernetes リソースです。
<resource>/<name> により、ポッド名を簡単にフィルタリングでき、
正確な ID を指定する必要はありません (たとえば、デプロイメントを省略する)
ID)。ポッドが削除された場合は末尾から削除され、新しいポッドが追加された場合は末尾から削除されます。
自動的に尾行されます。
ポッドに複数のコンテナが含まれている場合、Stern は、それらをすべて追跡することもできます。
これをそれぞれ手動で行う必要があります。コンテナフラグを指定するだけで、
表示するコンテナを制限します。デフォルトでは、すべてのコンテナがリッスンされます。
github.com/stern/stern@latest をインストールしてください
asdf (Linux/macOS)
asdf を使用する場合は、次のようにインストールできます。
asdf プラグイン追加船尾
ASDF 船尾の最新のインストール
自作 (Linux/macOS)
Homebrew を使用する場合は、次のようにインストールできます。
醸造インストール船尾
クルー (Linux/macOS/Windows)
kubectl プラグインのパッケージマネージャーである Krew を使用する場合は、次のようにインストールできます。
kubectl krew インストール船尾
WinGet (Windows)
Windows で実行している場合、通常、実際に最善の方法は、Windows の組み込みパッケージである WinGet を使用してインストールすることです。

古い場合は、次のようにインストールできます。
winget インストール stern.stern
使用法
stern ポッドクエリ [フラグ]
pod-query は、 <resource>/<name> 形式の正規表現または Kubernetes リソースです。
Kubernetes リソースではない場合、クエリは正規表現です。
したがって、「web-\w」を指定して web-backend および web-frontend ポッドを尾行することはできますが、 web-123 は指定できません。
クエリが <resource>/<name> (完全一致) の形式の場合、それに属するすべてのポッドを選択できます。
指定された Kubernetes リソース (deployment/nginx など) にコピーします。
サポートされている Kubernetes リソースは、 pod 、 replicationcontroller 、 service 、 daemonset 、deployment 、
レプリカセット、ステートフルセット、およびジョブ。
旗
デフォルト
目的
--all-namespaces 、-A
偽
存在する場合、すべての名前空間を追跡します。 --namespace で指定した場合でも、特定の名前空間は無視されます。
--バースト
0
Kubernetes API サーバーへのスロットルの最大バースト。デフォルトは 0 (client-go のデフォルトを使用)。 --qps=-1 の場合は無視されます。
--色
自動
カラー出力を強制的に設定します。 'auto': tty が接続されている場合に色付けします、'always': 常に色付けします、'never': 色付けしません。
--完了
指定されたシェルの stern コマンドライン完了コードを出力します。 「bash」、「zsh」、または「fish」を使用できます。
--条件
フィルタリングする条件: [条件名[=条件値]]。デフォルトの条件値は true です。一致では大文字と小文字が区別されません。現在、 --tail=0 または --no-follow でのみサポートされています。
--config
~/.config/stern/config.yaml
Stern 構成ファイルへのパス
--コンテナ 、 -c
.*
ポッド内に複数のコンテナがある場合のコンテナ名。 (正規表現)
--コンテナの色
コンテナ名の強調表示に使用する色を指定します。 --pod-colors と同じ形式を使用します。省略した場合は --pod-colors の値がデフォルトとなり、その長さと一致する必要があります。
--コンテナの状態
すべて
実行中、待機中、終了済み、またはすべての状態のコンテナーを追跡します。 '全て'

すべてのコンテナの状態に一致します。複数の状態を指定するには、これを繰り返すか、カンマ区切りの値を設定します。
--コンテキスト
使用するkubeconfigコンテキストの名前
--diff-container 、-d
偽
コンテナごとに異なる色を表示します。
--ephemeral-containers
本当の
一時的なコンテナを含めるか除外します。
--exclude 、-e
[]
除外するログ行。 (正規表現)
--exclude-container 、-E
[]
ポッド内に複数のコンテナがある場合に除外するコンテナ名。 (正規表現)
--ポッドの除外
[]
除外するポッド名。 (正規表現)
--フィールドセレクター
フィルタリングするセレクター (フィールド クエリ)。存在する場合、ポッドクエリのデフォルトは「.*」になります。
--ハイライト、-H
[]
ハイライトするログ行。 (正規表現)
--include 、-i
[]
含めるログ行。 (正規表現)
--init-containers
本当の
init コンテナを含めるか除外します。
--kubeconfig
キューブへの道
[切り捨てられた]
Stern は、$KUBECONFIG 環境変数が設定されている場合、それを使用します。両方の場合、
環境変数と --kubeconfig フラグが渡され、cli フラグが次のようになります。
使用されています。
設定ファイルを使用して、stern オプションのデフォルト値を変更できます。デフォルトの構成ファイルのパスは ~/.config/stern/config.yaml です。
# <フラグ名>: <値>
テール：10
最大ログリクエスト数 : 999
タイムスタンプ: 短い
--config フラグまたは STERNCONFIG 環境変数を使用して、構成ファイルのパスを変更できます。
stern はカスタム ログ メッセージの出力をサポートしています。事前に定義されているものがいくつかあります
--output フラグを指定して使用できるテンプレート:
--template フラグを介してカスタム テンプレートを受け入れます。
Go テンプレートにコンパイルされ、すべてのログ メッセージに使用されます。この Go テンプレート
次の構造体を受け取ります。
テンプレート内では、次の関数を使用できます（組み込みの関数に加えて）
関数):
--verbosity フラグを使用して、ログ レベルの詳細度を構成できます。
それは

トラブルシューティングで stern が Kubernetes API サーバーとどのように対話するかを知りたい場合に役立ちます。
詳細度を高めると、ログの数が増加します。 --verbosity 6 が出発点として適しています。
Stern には、クラスターへの意図しない負荷を防ぐために要求する同時ログの最大数があります。
この数は --max-log-requests フラグで構成できます。
--no-follow フラグの有無に応じて、動作とデフォルトが異なります。
--max-log-requests 1 と --no-follow の組み合わせは、ログを順番に表示したい場合に役立ちます。
以下に示すように、SGR (Select Graphic Rendition) シーケンスのカンマ区切りリストを使用して、構成ファイル内のポッドとコンテナーのハイライト色を構成できます。 Container-colors を省略すると、ポッドの色がコンテナの色としても使用されます。
# グリーン、イエロー、ブルー、マゼンタ、シアン、ホワイト
ポッドの色: " 32,33,34,35,36,37 "
# 下線付きの色 (4)
# 空の場合、ポッドの色がコンテナの色として使用されます
コンテナの色: " 32;4,33;4,34;4,35;4,36;4,37;4 "
この形式では、端末がサポートしている場合、下線、背景色、8 ビット カラー、24 ビット カラーなどのさまざまな属性を使用できるようになります。
同等のフラグ --pod-colors および --container-colors も使用できます。次のコマンドは、 --pod-colors フラグを使用して 24 ビット カラーを適用します。
#モノカイのテーマ
podColors= " 38;2;255;97;136,38;2;169;220;118,38;2;255;216;102,38;2;120;220;232,38;2;171;157;242 "
stern --pod-colors " $podColors " デプロイ/アプリ
例:
すべての名前空間からのすべてのログを追跡する
船尾。 --all-namespaces
以前のログを出力せずに kube-system 名前空間を追跡します
船尾。 -n kube-system --tail 0
ステージング上の envvars ポッド内で実行されているゲートウェイ コンテナーを追跡します
stern envvars --context staging --コンテナゲートウェイ
テールt

istio-proxy コンテナからのログを除外するステージング名前空間
stern -n staging --exclude-container istio-proxy 。
kube-apiserver ポッドからのログを除外して、kube-system 名前空間を追跡します。
stern -n kube-system --exclude-pod kube-apiserver 。
15 分前の認証アクティビティをタイムスタンプ付きで表示します
stern auth -t --since 15m
過去 5 分間のすべてのログを時間順に表示します
stern --since=5m --no-follow --only-log-lines -A -t 。 |ソート -k4
特定のタイムゾーンのタイムスタンプを使用して認証アクティビティを表示します (デフォルトはローカル タイムゾーン)
stern auth -t --timezone アジア/東京
minikube の新機能の開発を追跡する
厳しいいくつかの新機能 --context minikube
別の名前空間からポッドを表示する
厳しい kubernetes-dashboard --namespace kube-system
すべての名前空間にわたって run=nginx ラベル セレクターでフィルターされたポッドを追跡します
stern --all-namespaces -l run=nginx
カナリア リリースのフロントエンド ポッドをフォローする
厳しいフロントエンド --selector release=canary
すべての名前空間にわたって kind-control-plane ノード上のポッドを追跡します
stern --all-namespaces --field-selector spec.nodeName=kind-control-plane
デプロイメント/nginx によって作成されたポッドを追跡します
厳格なデプロイ/nginx
ログ メッセージを jq にパイプします。
stern バックエンド -o json | jq 。
ログ メッセージ自体のみを出力します。
厳しいバックエンド -o raw
カスタム テンプレートを使用して出力します。
stern --template '{{printf "%s (%s/%s/%s/%s)\n" .Message .NodeName .Namespace .PodName .ContainerName}}' バックエンド
船尾が提供する色のカスタム テンプレートを使用して出力します。
stern --template '{{.Message}} ({{.Namespace}}/{{color .PodColor .PodName}}/{{color .ContainerColor .ContainerName}}){{"\n"}}' バックエンド
parseJSON でカスタム テンプレートを使用して出力します。
stern --template='{{.PodName}}/{{.ContainerName}} {{with $d := .Message | parseJSON}}[{{$d.level}}] {{$d.message}}{{end}}{{"\n"}}' バックエンド
カスタム テンプレートを使用して出力する

JSON を解析するか、生の形式にフォールバックしようとするもの:
stern --template='{{.PodName}}/{{.ContainerName}} {{ with $msg := .Message | tryParseJSON }}[{{ colorGreen (toRFC3339Nano $msg.ts) }}] {{ levelColor $msg.level }} ({{ colorCyan $msg.caller }}) {{ $msg.msg }}{{ else }} {{ .Message }} {{ end }}{{"\n"}}' バックエンド
JSON (JSON の場合) をそのまま出力して出力します。
# .Message を JSON として解析し、適切に出力しようとします。そうでない場合は、json がそのまま出力されます
船尾 --template='{{ .メッセージ | prettyJSON }}{{"\n"}}' バックエンド
# または、解析された json を使用すると、`with` により非 json ログが削除されます
stern --template='{{ with $msg := .Message | tryParseJSON }}{{ prettyJSON $msg }}{{"\n"}}{{end}}' バックエンド
ファイルからカスタム テンプレートをロードします。
stern --template-file=~/.stern.tpl バックエンド
対話型プロンプトをトリガーして、「app.kubernetes.io/instance」ラベル値を選択します。
船尾 -p
ログ行のみを出力します。
船尾。 --only-log-lines
標準入力から読み取る:
stern --stdin < service.log
準備ができていないポッドのログのみを表示します。
船尾。 --condition=ready=false --tail=0
完成
Stern は、bash、zsh、または Fish のコマンドライン自動補完をサポートしています。 stern --completion=(bash|zsh|fish) は、次のように動作するシェル完了コードを出力します。
で評価されました

[切り捨てられた]

## Original Extract

⎈ Multi pod and container log tailing for Kubernetes -- Friendly fork of https://github.com/wercker/stern - stern/stern

GitHub - stern/stern: ⎈ Multi pod and container log tailing for Kubernetes -- Friendly fork of https://github.com/wercker/stern · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
stern
/
stern
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
316 Commits 316 Commits .github .github cmd cmd hack hack stern stern .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum main.go main.go View all files Repository files navigation
Fork of discontinued wercker/stern
Stern allows you to tail multiple pods on Kubernetes and multiple containers
within the pod. Each result is color coded for quicker debugging.
The query is a regular expression or a Kubernetes resource in the form
<resource>/<name> so the pod name can easily be filtered and
you don't need to specify the exact id (for instance omitting the deployment
id). If a pod is deleted it gets removed from tail and if a new pod is added it
automatically gets tailed.
When a pod contains multiple containers Stern can tail all of them too without
having to do this manually for each one. Simply specify the container flag to
limit what containers to show. By default all containers are listened to.
go install github.com/stern/stern@latest
asdf (Linux/macOS)
If you use asdf , you can install like this:
asdf plugin add stern
asdf install stern latest
Homebrew (Linux/macOS)
If you use Homebrew , you can install like this:
brew install stern
Krew (Linux/macOS/Windows)
If you use Krew which is the package manager for kubectl plugins, you can install like this:
kubectl krew install stern
WinGet (Windows)
If you're running on Windows, usually the actual best way is to use WinGet to install which is Windows' built-in package manager, you can install like this:
winget install stern.stern
Usage
stern pod-query [flags]
The pod-query is a regular expression or a Kubernetes resource in the form <resource>/<name> .
The query is a regular expression when it is not a Kubernetes resource,
so you could provide "web-\w" to tail web-backend and web-frontend pods but not web-123 .
When the query is in the form <resource>/<name> (exact match), you can select all pods belonging
to the specified Kubernetes resource, such as deployment/nginx .
Supported Kubernetes resources are pod , replicationcontroller , service , daemonset , deployment ,
replicaset , statefulset and job .
flag
default
purpose
--all-namespaces , -A
false
If present, tail across all namespaces. A specific namespace is ignored even if specified with --namespace.
--burst
0
Maximum burst for throttle to the Kubernetes API server. Defaults to 0 (use client-go default). Ignored when --qps=-1.
--color
auto
Force set color output. 'auto': colorize if tty attached, 'always': always colorize, 'never': never colorize.
--completion
Output stern command-line completion code for the specified shell. Can be 'bash', 'zsh' or 'fish'.
--condition
The condition to filter on: [condition-name[=condition-value]. The default condition-value is true. Match is case-insensitive. Currently only supported with --tail=0 or --no-follow.
--config
~/.config/stern/config.yaml
Path to the stern config file
--container , -c
.*
Container name when multiple containers in pod. (regular expression)
--container-colors
Specifies the colors used to highlight container names. Use the same format as --pod-colors. Defaults to the values of --pod-colors if omitted, and must match its length.
--container-state
all
Tail containers with state in running, waiting, terminated, or all. 'all' matches all container states. To specify multiple states, repeat this or set comma-separated value.
--context
The name of the kubeconfig context to use
--diff-container , -d
false
Display different colors for different containers.
--ephemeral-containers
true
Include or exclude ephemeral containers.
--exclude , -e
[]
Log lines to exclude. (regular expression)
--exclude-container , -E
[]
Container name to exclude when multiple containers in pod. (regular expression)
--exclude-pod
[]
Pod name to exclude. (regular expression)
--field-selector
Selector (field query) to filter on. If present, default to ".*" for the pod-query.
--highlight , -H
[]
Log lines to highlight. (regular expression)
--include , -i
[]
Log lines to include. (regular expression)
--init-containers
true
Include or exclude init containers.
--kubeconfig
Path to the kube
[truncated]
Stern will use the $KUBECONFIG environment variable if set. If both the
environment variable and --kubeconfig flag are passed the cli flag will be
used.
You can use the config file to change the default values of stern options. The default config file path is ~/.config/stern/config.yaml .
# <flag name>: <value>
tail : 10
max-log-requests : 999
timestamps : short
You can change the config file path with --config flag or STERNCONFIG environment variable.
stern supports outputting custom log messages. There are a few predefined
templates which you can use by specifying the --output flag:
It accepts a custom template through the --template flag, which will be
compiled to a Go template and then used for every log message. This Go template
will receive the following struct:
The following functions are available within the template (besides the builtin
functions ):
You can configure the log level verbosity by the --verbosity flag.
It is useful when you want to know how stern interacts with a Kubernetes API server in troubleshooting.
Increasing the verbosity increases the number of logs. --verbosity 6 would be a good starting point.
Stern has the maximum number of concurrent logs to request to prevent unintentional load to a cluster.
The number can be configured by the --max-log-requests flag.
The behavior and the default are different depending on the presence of the --no-follow flag.
The combination of --max-log-requests 1 and --no-follow will be helpful if you want to show logs in order.
You can configure highlight colors for pods and containers in the config file using a comma-separated list of SGR (Select Graphic Rendition) sequences , as shown below. If you omit container-colors , the pod colors will be used as container colors as well.
# Green, Yellow, Blue, Magenta, Cyan, White
pod-colors : " 32,33,34,35,36,37 "
# Colors with underline (4)
# If empty, the pod colors will be used as container colors
container-colors : " 32;4,33;4,34;4,35;4,36;4,37;4 "
This format enables the use of various attributes, such as underline, background colors, 8-bit colors, and 24-bit colors, if your terminal supports them.
The equivalent flags --pod-colors and --container-colors are also available. The following command applies 24-bit colors using the --pod-colors flag.
# Monokai theme
podColors= " 38;2;255;97;136,38;2;169;220;118,38;2;255;216;102,38;2;120;220;232,38;2;171;157;242 "
stern --pod-colors " $podColors " deploy/app
Examples:
Tail all logs from all namespaces
stern . --all-namespaces
Tail the kube-system namespace without printing any prior logs
stern . -n kube-system --tail 0
Tail the gateway container running inside of the envvars pod on staging
stern envvars --context staging --container gateway
Tail the staging namespace excluding logs from istio-proxy container
stern -n staging --exclude-container istio-proxy .
Tail the kube-system namespace excluding logs from kube-apiserver pod
stern -n kube-system --exclude-pod kube-apiserver .
Show auth activity from 15min ago with timestamps
stern auth -t --since 15m
Show all logs of the last 5min by time, sorted by time
stern --since=5m --no-follow --only-log-lines -A -t . | sort -k4
Show auth activity with timestamps in specific timezone (default is your local timezone)
stern auth -t --timezone Asia/Tokyo
Follow the development of some-new-feature in minikube
stern some-new-feature --context minikube
View pods from another namespace
stern kubernetes-dashboard --namespace kube-system
Tail the pods filtered by run=nginx label selector across all namespaces
stern --all-namespaces -l run=nginx
Follow the frontend pods in canary release
stern frontend --selector release=canary
Tail the pods on kind-control-plane node across all namespaces
stern --all-namespaces --field-selector spec.nodeName=kind-control-plane
Tail the pods created by deployment/nginx
stern deployment/nginx
Pipe the log message to jq:
stern backend -o json | jq .
Only output the log message itself:
stern backend -o raw
Output using a custom template:
stern --template '{{printf "%s (%s/%s/%s/%s)\n" .Message .NodeName .Namespace .PodName .ContainerName}}' backend
Output using a custom template with stern-provided colors:
stern --template '{{.Message}} ({{.Namespace}}/{{color .PodColor .PodName}}/{{color .ContainerColor .ContainerName}}){{"\n"}}' backend
Output using a custom template with parseJSON :
stern --template='{{.PodName}}/{{.ContainerName}} {{with $d := .Message | parseJSON}}[{{$d.level}}] {{$d.message}}{{end}}{{"\n"}}' backend
Output using a custom template that tries to parse JSON or fallbacks to raw format:
stern --template='{{.PodName}}/{{.ContainerName}} {{ with $msg := .Message | tryParseJSON }}[{{ colorGreen (toRFC3339Nano $msg.ts) }}] {{ levelColor $msg.level }} ({{ colorCyan $msg.caller }}) {{ $msg.msg }}{{ else }} {{ .Message }} {{ end }}{{"\n"}}' backend
Pretty print JSON (if it is JSON) and output it:
# Will try to parse .Message as JSON and pretty print it, if not json will output as is
stern --template='{{ .Message | prettyJSON }}{{"\n"}}' backend
# Or with parsed json, will drop non-json logs because of `with`
stern --template='{{ with $msg := .Message | tryParseJSON }}{{ prettyJSON $msg }}{{"\n"}}{{end}}' backend
Load custom template from file:
stern --template-file=~/.stern.tpl backend
Trigger the interactive prompt to select an 'app.kubernetes.io/instance' label value:
stern -p
Output log lines only:
stern . --only-log-lines
Read from stdin:
stern --stdin < service.log
Only display logs for pods that are not ready:
stern . --condition=ready=false --tail=0
Completion
Stern supports command-line auto completion for bash, zsh or fish. stern --completion=(bash|zsh|fish) outputs the shell completion code which work by being
evaluated in

[truncated]
