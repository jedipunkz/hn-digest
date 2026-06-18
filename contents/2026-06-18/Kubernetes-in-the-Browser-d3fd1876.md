---
source: "https://github.com/ngrok/webernetes"
hn_url: "https://news.ycombinator.com/item?id=48588578"
title: "Kubernetes in the Browser"
article_title: "GitHub - ngrok/webernetes: Kubernetes in the browser. · GitHub"
author: "FelipeCortez"
captured_at: "2026-06-18T18:55:29Z"
capture_tool: "hn-digest"
hn_id: 48588578
score: 5
comments: 0
posted_at: "2026-06-18T17:22:24Z"
tags:
  - hacker-news
  - translated
---

# Kubernetes in the Browser

- HN: [48588578](https://news.ycombinator.com/item?id=48588578)
- Source: [github.com](https://github.com/ngrok/webernetes)
- Score: 5
- Comments: 0
- Posted: 2026-06-18T17:22:24Z

## Translation

タイトル: ブラウザーでの Kubernetes
記事のタイトル: GitHub - ngrok/webernetes: ブラウザーの Kubernetes。 · GitHub
説明: ブラウザーの Kubernetes。 GitHub でアカウントを作成して、ngrok/webernetes の開発に貢献してください。

記事本文:
GitHub - ngrok/webernetes: ブラウザー内の Kubernetes。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ngrok
/
webernetes
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらにアクトを開く

イオンメニュー フォルダーとファイル
552 コミット 552 コミット .changeset .changeset .github/ workflows .github/ workflows .husky .husky デモ デモの例 例 スクリプト スクリプト src src .editorconfig .editorconfig .gitignore .gitignore .lintstagedrc.json .lintstagedrc.json .nvmrc .nvmrc .oxlintrc.json .oxlintrc.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンスに関する通知 README.md README.md misse.lock misse.lock misse.toml misse.toml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.browser.config.ts vitest.browser.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ブラウザ上で動作する Kubernetes。
実際の動作を確認するには、デモをチェックしてください。
このプロジェクトは、Kubernetes プロジェクトのサブセットを移植して、そのようなプロジェクトを実現します。
バックエンドサーバーなしでクラスターをブラウザーで起動できること
コンポーネント。
ngrok では、ビジュアルでインタラクティブなコンテンツを作りたいと考えています。
Kubernetesについて。私たちはインフラストラクチャを作成して維持したくありませんでした。
実際のクラスターを起動するため、ブラウザベースのシミュレーターを作成することにしました。
代わりに。希望と夢は、これが私たち (そしてあなたたち) にとっても可能になることです。
長期間存続するインタラクティブな Kubernetes コンテンツを作成するには、
メンテナンスの負担が大幅に軽減されます。
注意: これは非常に実験的なものです。 API は変更される可能性があります。
さまざまなリソースのサポート レベルは変更される可能性があります。私はちょっと
やっていくうちにこれを理解します。
まず、Webernetes を依存関係としてインストールします。
npm インストール @ngrok/webernetes
次に、クラスター内で実行するイメージを定義します。 Webernetes は実際には動作しません
Docker Hub からのイメージを取得することも、そうすることが目標ではありません。
import { BaseImage , type ProcessContext } from

"@ngrok/webernetes" ;
class MyImage extends BaseImage {
// imageName 変数と imageVersion 変数は画像ラベルを構成するものです
// コンテナ定義で使用します。ここには my-image:1.0 がありますが、
// webernetes は、my-image だけを指定した場合や、
// 私の画像:最新
static readonly imageName = "私の画像" ;
静的読み取り専用 imageVersion = "1.0" ;
// コンテナマニフェストで他のコマンドが指定されていない場合、これは
// 以下の argv として渡されるコマンド。
readonly defaultCommand = [ "サーバー" ] ;
// exec はイメージのメイン エントリ ポイントです。で呼び出されます
// コンテナ定義から渡されるコマンドライン引数。
override async exec (ctx : ProcessContext , argv : readonly string [] ) : Promise <number> {
if ( argv [ 0 ] !== "サーバー" ) {
// 基本イメージは、一連のコア ユーティリティ (cat、false、printenv、
// など）そのため、コマンドを認識できない場合は、ベースに戻ります
// 画像。
戻るのを待ってスーパー。 exec (ctx, argv) ;
}
// このコンテナのポート 8080 にバインドします。
ctx 。 listenHttp (8080 , async (リクエスト) => {
戻り値 {
ステータスコード : 200 、
本文 : "こんにちは、世界\n" 、
} ;
} ) ;
// クラスターのシャットダウン時に長時間実行プロセスをキャンセルできるようにするために必要です
// ダウン。ここで終了コード 0 を返した場合、上記のリスナーは次のようになります。
// このコンテナは終了しているため登録を解除します。
ctx を待ちます。 waitUntilKilled() ;
}
}
次に、クラスターを作成し、それにイメージを登録します。
"@ngrok/webernetes" から {クラスター} をインポートします。
const クラスター = 新しいクラスター () ;
クラスター。 registerImage (MyImage) ;
次に、クラスターを実行し、その中のイメージを使用してポッドを生成します。
// デフォルトでは、これにより 3 ノードのクラスターが起動されます。これは現在変更できません。
クラスターを待ちます。初期化() ;
クラスターを待ちます。適用する ( [
{
APIV

バージョン: "v1" 、
種類：「ポッド」、
メタデータ: {
名前：「マイポッド」、
ラベル : { アプリ : "my-pod" } 、
} 、
仕様: {
コンテナ: [
{
名前: "私のコンテナ" 、
画像 : "私の画像:1.0" 、
} 、
]、
} 、
} 、
]);
ポッドにリクエストを送信するには、ポッドと通信するサービスを作成する必要があります。
この場合、NodePort サービスが最も簡単なルートを提供します。
クラスターを待ちます。適用する ( [
{
apiVersion : "v1" 、
種類：「サービス」、
メタデータ : { 名前 : "my-service" } 、
仕様: {
タイプ: "ノードポート" 、
ポート: [
{
ポート: 80、
ターゲットポート: 8080 、
ノードポート: 31000、
プロトコル: "TCP" 、
} 、
]、
セレクター: {
アプリ: "my-pod" 、
} 、
} 、
} 、
]);
const resp = クラスターを待ちます。 fetch ( "http://node-1:31000" ) ;
const text = 応答を待ちます。文章 （ ） ; // こんにちは、世界
ポッドは HTTP 経由で相互に通信することもできます。これがどのように機能するかを確認するには、
実行可能な例がいくつかあるので、examples/ ディレクトリの下のコードを確認してください。のために
完全なビジュアルデモについては、demo/ ディレクトリにあるコードをチェックしてください。
何が実装され、何が実装されていないのか
これまでのところ、最初のコンテンツを作成するために必要な部分に範囲を絞りました
私が作りたいのは、探求することです。
また、私は決してすべての分野について網羅的な専門家ではないことを前置きしておきます。
Kubernetes について詳しく説明していないため、何かが抜けているか、まだ分かっていない可能性があります。
私が持っていると信じていることを完全に実装しました。
クラスターは 3 ノードクラスター (node-1、node-2、node-3) をスピンアップします。
まだ設定可能ではありません。任意に追加すると仮定したいと思います
将来的にはノードを削除します。
内のリソースを削除するための特別な処理を含む、サポートされています。
名前空間コントローラー経由の名前空間 (ガベージ コレクターとは別)
他のすべての削除を処理するコントローラー)。
基本がサポートされています: Pod には Container を含めることができ、それらのコンテナーは
ポート上で HTTP トラフィックをリッスンします。ポッド名を取得します。

IP アドレスがあれば、
DNS 名または IP アドレスによって他のポッドと通信します。環境を受け入れることができる
変数。彼らは捜査を受けます。
おそらく他にもたくさんのことがありますが、これらが私にとって大きなものです。
ClusterIP および NodePort サービスのサポートは、LoadBalancer および
外部名サービスはまだサポートされていません。ポッドはサービス DNS と通信できます
名前とリクエストはサービス内の Pod 全体で負荷分散されます。
ラウンドロビンを使用します。
UDP はサポートされていません。よく考えてみると、TCP もどちらかというとそうではありませんが、私はそうではありません
ネットワークスタックの下位までエミュレートします。 HTTP と DNS をそれぞれに通信できるもの
その他、それで終わりです。私はこれを変更したい、または変更する必要があるとは決して思っていません。
その結果、IP ファミリ間の区別も実際にはモデル化されません。
開始するまで存在を知らなかった Service の楽しい実装の詳細
このプロジェクト。これらは、 の一部である Pod のセットを追跡するために作成されます。
サービス。通常、それらはそれぞれ 100 個の Pod にシャード化されますが、私はそれを行っていません。
純粋に簡単にするためです。それらは存在し、正常に動作しますが、シャーディングは
今のところはありません。
ほとんどの部分がサポートされており、同じイベントを確実に発生させるように努めました
Kubernetes と同じように。イベントの集約は行っていませんし、集約できない可能性もあります。
すべてのフィールドが存在し、正しいですが、メッセージを含むイベントが発生し、
検査することができます。
サポートされており、通常は Deployment によって作成されます。 ReplicaSetコントローラーも入っています
上流の Kubernetes ReplicaSet コントローラーとほぼ同等です。
RollingUpdate および Recreate 戦略を含むサポートされています。導入
コントローラーが配置されており、アップストリームの Kubernetes とほぼ同等です
デプロイメントコントローラー。
このリポジトリは、mise を使用してツールチェーン (Node、pnpm、
ast-grep 、 ripgrep ）なので、マシン間で再現可能です。ノード

から読み取られます
.nvmrc と pnpm は package.json#packageManager から単一ソースされます。
misse をインストールした後、新しい状態から
クローン:
misse install # 固定ツールをインストールします (そして、mise.lock を書き込みます)
misse run setup # ロックファイルからワークスペースの依存関係をインストールします
利用可能なタスク:
インストールを実行してください — pnpm install --frozen-lockfile 。
セットアップを実行する — 新しいクローンを作成した後にリポジトリを準備します ( install を実行します)。
misse run relock — .nvmrc と一致するように misse.lock を更新し、
package.json#packageManager 。
ドクターを実行する — アクティブなツールがコミットされたピンと一致することを確認します。
固定されたバージョンをバンプするには、 .nvmrc 、 package.json#packageManager 、または
misse.toml を実行し、mise run relock を実行します。パッケージスクリプト ( pnpm test 、
pnpm build 、 pnpm vibe-check など) は変更されず、通常どおり一度実行されます。
依存関係がインストールされます。
Apache-2.0ライセンス
行動規範
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

Kubernetes in the browser. Contribute to ngrok/webernetes development by creating an account on GitHub.

GitHub - ngrok/webernetes: Kubernetes in the browser. · GitHub
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
ngrok
/
webernetes
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
552 Commits 552 Commits .changeset .changeset .github/ workflows .github/ workflows .husky .husky demo demo examples examples scripts scripts src src .editorconfig .editorconfig .gitignore .gitignore .lintstagedrc.json .lintstagedrc.json .nvmrc .nvmrc .oxlintrc.json .oxlintrc.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE NOTICE NOTICE README.md README.md mise.lock mise.lock mise.toml mise.toml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.browser.config.ts vitest.browser.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Kubernetes that runs in your browser.
To see it in action, check out the demo!
This project is a port of a subset of the Kubernetes project to make it such
that clusters can be booted up in the browser, without any backend server
components.
At ngrok , we want to make visual and interactive content
about Kubernetes. We didn't want to create and maintain infrastructure for
spinning up real clusters, so we decided to create a browser-based simulator
instead. The hope and dream is that this will make it possible for us (and you!)
to create interactive Kubernetes content that lives for a long time, because the
maintenance burden is much smaller.
Please note: This is very experimental. The API is subject to change, the
level of support for different resources is subject to change. I'm kinda
figuring this out as I go.
First, install webernetes as a dependency:
npm install @ngrok/webernetes
Then define an image to run in your cluster. Webernetes does not run real
images from Docker Hub, nor is it a goal to do so.
import { BaseImage , type ProcessContext } from "@ngrok/webernetes" ;
class MyImage extends BaseImage {
// The imageName and imageVersion variables are what make up the image label
// you'll use in your container definition. Here we have my-image:1.0 but
// webernetes also knows what to do if you specify just my-image or
// my-image:latest
static readonly imageName = "my-image" ;
static readonly imageVersion = "1.0" ;
// If no other command is specified in your container manifest, this is the
// command that will be passed in as argv below.
readonly defaultCommand = [ "server" ] ;
// exec is the main entrypoint for your image. It will be called with the
// command-line arguments passed in from your container definition.
override async exec ( ctx : ProcessContext , argv : readonly string [ ] ) : Promise < number > {
if ( argv [ 0 ] !== "server" ) {
// The base image defines a bunch of core utils (cat, false, printenv,
// etc.) so if we don't recognize the command, fall back to the base
// image.
return await super . exec ( ctx , argv ) ;
}
// Binds to port 8080 on this container.
ctx . listenHttp ( 8080 , async ( request ) => {
return {
statusCode : 200 ,
body : "hello, world\n" ,
} ;
} ) ;
// Required for long-running processes to be cancellable when clusters shut
// down. If we returned an exit code of 0 here, the listener above would be
// unregistered because this container will have exited.
return await ctx . waitUntilKilled ( ) ;
}
}
Then we create a cluster and register our image with it.
import { Cluster } from "@ngrok/webernetes" ;
const cluster = new Cluster ( ) ;
cluster . registerImage ( MyImage ) ;
And then we can run the cluster and spawn a pod using our image in it.
// By default this spins up a 3-node cluster. This can't currently be changed.
await cluster . init ( ) ;
await cluster . apply ( [
{
apiVersion : "v1" ,
kind : "Pod" ,
metadata : {
name : "my-pod" ,
labels : { app : "my-pod" } ,
} ,
spec : {
containers : [
{
name : "my-container" ,
image : "my-image:1.0" ,
} ,
] ,
} ,
} ,
] ) ;
To send a request to your pod, you'll need to create a Service to talk to it.
In this case, a NodePort service gives us the easiest route.
await cluster . apply ( [
{
apiVersion : "v1" ,
kind : "Service" ,
metadata : { name : "my-service" } ,
spec : {
type : "NodePort" ,
ports : [
{
port : 80 ,
targetPort : 8080 ,
nodePort : 31000 ,
protocol : "TCP" ,
} ,
] ,
selector : {
app : "my-pod" ,
} ,
} ,
} ,
] ) ;
const resp = await cluster . fetch ( "http://node-1:31000" ) ;
const text = await resp . text ( ) ; // hello, world
Pods are also able to talk to each other over HTTP. To see how this works in a
few runnable examples, check out the code under the examples/ directory. For
the full visual demo, check out the code under the demo/ directory.
What's implemented and what isn't
I've scoped this so far to the bits I need to make the first piece of content
I want to make, which is about probing.
I'll also preface this by saying I am by no means an exhaustive expert on every
detail of Kubernetes, so it is likely that I'm missing some things or I haven't
fully implemented the things I believe I have.
Cluster spins up a 3-node cluster ( node-1 , node-2 , node-3 ) and that
isn't configurable yet. I would like to suppose arbitrarily adding and
removing nodes in the future.
Supported, including special handling for deleting the resources within a
namespace via a namespace controller (separate to the garbage collector
controller that handles deleting everything else).
Basics are supported: Pod s can have Container s and those containers can
listen for HTTP traffic on ports. They get a pod name, an IP address, they can
speak to other pods by their DNS name or IP address. They can accept environment
variables. They get probed.
Probably a lot of other things, but those are the big ones that come to mine.
Support for ClusterIP and NodePort services is in, LoadBalancer and
ExternalName services are not yet supported. Pod s can talk to service DNS
names and the requests will be load balanced across the Pod s in the service
using round robin.
UDP isn't supported. TCP kinda sorta isn't either if you think about it, I'm not
emulating that far down the network stack. Stuff can talk HTTP and DNS to each
other and that's it. I don't anticipate ever wanting or needing to change this.
As a result, the distinction between IP families also isn't really modeled.
A fun implementation detail of Service s I had no idea existed until starting
this project. These are created to track sets of Pod s that are part of a
Service . They're usually sharded into 100 Pod s each but I haven't done that,
purely for simplicity. They exist, they work how they should, but the sharding
isn't there for now.
Supported for the most part, and I've tried to make sure we fire the same events
as Kubernetes does. I'm not doing any event aggregating, and it's possible not
all fields are present and correct, but events with messages do get fired and
can be inspected.
Supported and usually created by Deployment s. ReplicaSet controller is also in
place and largely at parity with the upstream Kubernetes ReplicaSet controller.
Supported, including RollingUpdate and Recreate strategies. Deployment
controller is in place and largely at parity with the upstream Kubernetes
Deployment controller.
This repo uses mise to pin the toolchain (Node, pnpm,
ast-grep , ripgrep ) so it's reproducible across machines. Node is read from
.nvmrc and pnpm is single-sourced from package.json#packageManager .
After installing mise , from a fresh
clone:
mise install # install the pinned tools (and write mise.lock)
mise run setup # install workspace dependencies from the lockfile
Available mise tasks:
mise run install — pnpm install --frozen-lockfile .
mise run setup — prepare the repo after a fresh clone (runs install ).
mise run relock — refresh mise.lock to match .nvmrc and
package.json#packageManager .
mise run doctor — verify the active tools match the committed pins.
To bump a pinned version, edit .nvmrc , package.json#packageManager , or
mise.toml and run mise run relock . The package scripts ( pnpm test ,
pnpm build , pnpm vibe-check , etc.) are unchanged and run as usual once
dependencies are installed.
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
