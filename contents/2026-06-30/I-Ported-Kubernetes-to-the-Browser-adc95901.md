---
source: "https://ngrok.com/blog/i-ported-kubernetes-to-the-browser"
hn_url: "https://news.ycombinator.com/item?id=48734656"
title: "I Ported Kubernetes to the Browser"
article_title: "I ported Kubernetes to the browser | ngrok blog"
author: "mariuz"
captured_at: "2026-06-30T16:44:33Z"
capture_tool: "hn-digest"
hn_id: 48734656
score: 5
comments: 0
posted_at: "2026-06-30T16:03:13Z"
tags:
  - hacker-news
  - translated
---

# I Ported Kubernetes to the Browser

- HN: [48734656](https://news.ycombinator.com/item?id=48734656)
- Source: [ngrok.com](https://ngrok.com/blog/i-ported-kubernetes-to-the-browser)
- Score: 5
- Comments: 0
- Posted: 2026-06-30T16:03:13Z

## Translation

タイトル: Kubernetes をブラウザに移植してみました
記事タイトル: Kubernetes をブラウザに移植してみました |ングロクのブログ
説明: 2 か月でほぼ 100,000 行の LLM 生成コードが作成されましたが、どれもいい加減なものはありませんでした。

記事本文:
Kubernetes をブラウザに移植しました | ngrok ブログ メイン コンテンツにスキップ ngrok ホーム / ngrok ブログ ホーム ブログ モバイル ナビゲーションを開く 製品 解決する問題 リソース ドキュメント 価格 ログイン サインアップ 製品 ゲートウェイ
dev からリモート K8s クラスターにアクセスする
接続性 サイト間の接続性
AI ゲートウェイのネイティブ Anthropic SDK サポート
LLM を 4 倍小さく、2 倍高速化
即時キャッシュ: 10 倍安い LLM トークン、しかしどうやって?
検索… Control ⌃ K ニュースレター RSS 2026 年 6 月 30 日
Kubernetesをブラウザに移植してみた
Sam Rose は、ngrok の上級開発者教育者であり、開発者が ngrok を最大限に活用できるようにするコンテンツの作成に重点を置いています。
Kubernetes をブラウザに移植しました (ハッカーニュース)
Linkedin のブラウザに Kubernetes を移植しました
Kubernetes をブラウザに移植したことを Twitter で共有しました
Reddit で Kubernetes をブラウザに移植しました
Whatsapp のブラウザに Kubernetes を移植しました
先週、私は Kubernetes の TypeScript への部分移植である webernetes をリリースしました。
ブラウザーでクラスターを実行できるようにします。結局ほぼ生成しました
629 ファイルにわたる 552 コミットの 100,000 行のコード。 2ヶ月かかりました。
以下のデモは Webernetes クラスターです。完全にブラウザ内で実行され、
実際の Kubernetes クラスターが行うのとほぼ同じ作業を実行します。ポッド
ライフサイクル、クラスター DNS とネットワーキング、コンテナー ガベージ コレクション、IP
割り当て、デプロイメントと ReplicaSet の追跡など。青い点々
相互にリクエストを送信するポッドを表します。
3 つの Kubernetes ノードにわたるデプロイメントから 3 つのポッド間を移動する HTTP リクエストを示すインタラクティブな Webernetes デモ。
リセット 一時停止 シミュレートされたクラスターを開始します。
このセクションをブックマークする ちょっと待って、私は何を見ているのでしょうか?
私が最もよく受ける質問は次のとおりです。「Kubernetes を次のようにコンパイルしましたか?」
Webアセンブリ

？」答えはノーです。シンプルな「こんにちは、世界！」コンパイルされた Go プログラム
WebAssembly は約 540KiB で gzip 圧縮されています。それだけでもすでにウェベルネットよりも大きいのですが、
これは gzip 圧縮された最大 140KiB です。すべての Kubernetes を WebAssembly にコンパイルしても、
疑いがあるということは、ネットワーク上でメガバイトを送信することを意味します。一応確認してみましたが、
残念ながら、Kubernetes はシステムレベルで呼び出しを行うため、コンパイル時にエラーが発生します。
ブラウザでは利用できない API。
Kubernetes の「kubelet」バイナリの部分ポート。ポッドを実行してプローブするのに十分です。
いくつかの Kubernetes 「コントローラー」のポート: ポッド スケジューラー、名前空間コントローラー、kube プロキシ、デプロイメント コントローラーなど。
コンテナー ネットワーク インターフェイス (CNI) をブラウザーベースで解釈したもので、ポッドはシミュレートされたネットワーク上で相互に通信できます。
ブラウザベースのコンテナ ランタイム。kubelet はコンテナ ランタイム インターフェイス (CRI) 経由で通信してコンテナを実行します。
Webernetes クラスターと対話してマニフェストの適用やリソースの監視などを行うための API。
Webernetes を小さく保ちたいという要望の結果、実際のイメージは取得されません
Docker Hub などのレジストリから。代わりに、独自のブラウザベースのレジストリがあります。
TypeScript API を使用して画像を定義します。画像は次のようになります。
1 import * as w8s from "@ngrok/webernetes" ; 2 3 クラス HelloWorld は w8s を拡張します。 BaseImage { ⋯ 4 static readonly imageName = "hello-world" ; 5 静的読み取り専用 imageVersion = "1.0" ; 6 7 async exec ( ctx : w8s . ProcessContext , argv : string []) : Promise <number> { ⋯ 8 ctx .listenHttp ( 8080 , async ( _ctx , request ) => { ⋯ 9 return { ⋯ 10 status : 200 , 11 body : "こんにちは。世界!" , 12 }; 13 }); 14 return await ctx .waitUntilKilled (); 15 } 16 }
イメージをクラスターにデプロイするには、次の手順を実行します。
1 import * as w8s from "@ngrok/webernetes" ; 2 3クラス こんにちは

ワールドは w8s を拡張します。 BaseImage { ⋯ 4 // 以前と同様 ... 5 } 6 7 const cluster = new w8s .Cluster (); 8 await クラスター .registerImage ( HelloWorld ); 9 10 const [ pod ] = await クラスター .apply ([ ⋯ 11 { ⋯ 12 apiVersion : "apps/v1" , 13 kind : "Deployment" , 14 メタデータ : { name : "hello-world-deployment" } , 15 spec : { ⋯ 16 レプリカ : 1 , 17セレクター : { ⋯ 18 matchLabels : { アプリ : "hello-world-pod" } , 19 } , 20 テンプレート : { ⋯ 21 メタデータ : { ⋯ 22 ラベル : { アプリ : "hello-world-pod" } , 23 } , 24 仕様 : { ⋯ 25 コンテナー : [ ⋯ 26 { ⋯ 27 名前 : "hello-world-container" , 28 イメージ : "hello-world:1.0" , 29 } , 30 ] , 31 } , 32 } , 33 } , 34 } , 35 ]);
次に、次のように webernetes API を使用してクラスターと対話できます。
1 // デフォルトの名前空間内のすべてのポッドをリストします。 2 const pods = await cluster 。 API 。 corev1 .listNamespacedPod ({ ⋯ 3 名前空間 : "デフォルト" , 4 }); 5 6 // すべての名前空間のポッドへの変更を監視します 7 const infomer = cluster .informer ( "pods" , ( type , pod ) => { ⋯ 8 console .log ( `pod ${ type } : ${ pod . metadata ?. name } ` ); 9 }); 10 11 // 完了したらインフォーマーを停止します 12 await密報者 .stop (); 13 14 // ポッドが相互にリクエストとレスポンスを送信していることをリッスンします。 15 // これは、上の移動するドットを視覚化する方法です。 16 クラスタ .on ( "リクエスト" , ( イベント ) => { ⋯ 17 コンソール .log ( `リクエスト: ${ イベント . リクエスト . メソッド } ${ イベント . リクエスト . URL } ` ); 18 }); 19 クラスタ .on ( "応答" , ( イベント ) => { ⋯ 20 コンソール .log ( `応答: ${ イベント . 応答 ?. ステータス } ` ); 21 }); 22 23 // クラスター ネットワークを使用してポッドにリクエストを送信します。これにより、 24 // 上記の要求/応答イベント ハンドラーもトリガーされます。 25 const ポッド = ポッド。アイテム [ 0 ]; 26 const resp = await クラスター .fetch ( `http:// ${ pod . status ?.podIP } :8

080/`); 27 コンソール .log (それぞれ .body ); // "こんにちは世界！"
webernetes リポジトリには、さらに多くの例があります。
Webernetes は、インタラクティブな Kubernetes コンテンツを作成するために使用することを目的としています。それは
本番環境に対応した Kubernetes ディストリビューションではありません。実際に実行する必要はない
画像。必要なのは、クリエイターが特定のワークロードを設定する方法だけです。
彼らが教えようとしていることを説明します。
時間の経過とともに、より多くの Kubernetes をサポートするために Webernetes を拡張するつもりです。
特徴。現時点では、ConfigMap、Secret、ポッド リソース、
永続ボリュームや、まだ必要としていないその他の機能が多数あります。私としては
このライブラリを使用してさらに多くのコンテンツを作成する場合は、必要なものをさらに実装します。
Webernetes 上に構築しようとしているが、Webernetes がサポートしていないものがある場合は、
必要な場合は、ご連絡ください。私は s.rose@ngrok.com です。
貢献者として喜んでお手伝いいたします。
このセクションをブックマークする これはただの怠慢ですか?
ほぼすべての Webernetes コードは LLM によって作成されました。私は人々がそうであることを期待しています
その結果、プロジェクトに疑問が生じました。ずさんな移植で非難されることを期待しています
ビューには Kubernetes を使用しますが、それが私が行ったことではないことを示したいと思います
完了しました。
このプロジェクトを手抜きのないものにするには、次の 2 つのことを実行しました。
コードのすべての行を確認しました。
私は、Webernetes が実際のクラスターと同じように動作することを保証する何百ものテストを作成しました。
最初のポイントは、これまでで最も時間がかかりましたが、どのようにして自信を獲得したかです。
コードの大部分が行ごとに Kubernetes と同一であること
コードベースに進みます。 2 番目のポイントは、語彙の類似性を確認する方法です。
実際には同じ動作になります。
私のレビュー後にコードベースに残った間違いはすべて私の責任です。
存在を疑う人もいます。見つけた場合は、問題を開いてお知らせください。
本

このセクションにマークを付けます コードをレビューする理由は何ですか?
C コンパイラまたはポートの作成に LLM が使用されているという記事を読んだ
Zig から Rust への Bun は、自動化された方法によって可能になりました。
正しさを主張する。 Anthropic には比較できる既存の C コンパイラがたくさんありました
Bun にはメンテナーが信頼する大規模な既存のテスト スイートがあった
手動レビューなしで 100 万行を超える新しい Rust コードをマージするには十分です。
私はそれらのものを持っていませんでした。テストスイートが必要な場合は、それを作成する必要があります
私自身。実際の Kubernetes と比較したい場合は、次のことを理解する必要があります。
それを行う方法。
webernetes のコードのほとんどは、Kubernetes Go コードベースから移植されています。私は
LLM を使用して移植しました。入力するよりも速いと確信していたからです。
しかし、すぐに遭遇した問題は、LLM は移植が苦手だということでした。
コード。私がどんなに努力しても、彼らは間違いを犯し続けました。間違い
いくつかのフレーバーがありました:
ショートカット。これは、大きなファイルを移植するときによく発生するようで、LLM の冗長性を軽減するために設計されたトレーニング後の選択と格闘しているのではないかと疑問に思いました。私がよく見たショートカットの例: Kubernetes にはさまざまな種類のキャッシュが含まれています。 LRU キャッシュ、期限切れキャッシュ、FIFO キャッシュ、変換キャッシュなどがあります。これらを実装する代わりに、マップを使用して LLM を捕捉すると、誤った動作が発生します。
役に立とうとしすぎます。 LLM は、元の Go コードには存在しなかったヘルパー関数を発明することで、コードを整理しようとします。多くの場合、これらのヘルパーは無害ですが、微妙な違いがある場合もあります。いずれにせよ、それらによりコードをオリジナルと並べてレビューするのが難しくなったので、LLM にそれらを削除するように依頼しました。
ただ…足りないもの。これは、テーブル テストを移植するときに最も頻繁に発生しました。

ゴーから。慣れていない方のために説明すると、テーブル テストは次のようになります。これらは、テスト コードがその下にあるテスト ケースの配列です。 LLM は恣意的にテストを省略するので、私はその理由を尋ねなければなりませんでした。場合によっては、テストが適用できないことについて話されることがあります。場合によってはそれが当てはまりますが、通常、LLM は誤ってテストを省略したことを認めます。
少なくとも何人かは「スキルの問題」と叫び、準備ができていることを私は知っています。
プロンプトをもっと上手にする必要があるというコメント。それは本当かもしれません！私はそうするだろう
このテーブルを完全にワンショットで移植するプロンプトの例を見るのが大好きです
「Go to TypeScript」からテストします。あなたは私に膨大な時間を節約してくれるでしょう
将来的には。
それまでは、LLM による移植に自信を持つためには、次のことを行う必要があります。
出力を確認します。ショートカットはわかりません。
このセクションをブックマークする テストの内容
コードが並べて同一であることを知っているのは良いことですが、
実際に機能しますか？ Go と JavaScript は実行環境が異なるため、
同じコードがそれぞれで異なる動作をする可能性は常にありました。私も
結局、JavaScript バージョンの Channels 、mutexes を作成する必要がありました。
Go の select ステートメント、およびその他の Go-ism。彼らが働いていることを知る必要がありました
重要なシナリオ。
これに満足するために、まったく同じコードを実行するテストを作成しました。
webernetes と k3s クラスターの両方。これを行うには、次の API が必要でした。
既存の Kubernetes API と一致する webernetes。私が選んだ
kubernetes-client/javascript は、の公式クライアント ライブラリであるため、
JavaScript の Kubernetes には TypeScript タイプがあります。
1 import { 期待、それ } from "vitest" ; 2 import { kubernetes } から "../../test/harnesses/kubernetes" ; 3 4 // `kubernetes.describe` は舞台裏で魔法を実行して、次のいずれかをセットアップします。 5 // k3s (h

ttps://k3s.io/) クラスターまたは Webernetes クラスターを指定し、それを 6 // `context` 引数経由で渡します。 7 // 8 // 次に、`pnpm test:node` または `pnpm test:browser` を実行して、Node 環境を使用して k3s に対してテストを実行するか、ヘッドレス ブラウザを使用して webernetes に対してテストを実行します。 10 kubernetes .describe ( "Pods" , ( context ) => { ⋯ 11 const { core } = context ; 12 const { getTestNamespace , waitFor } = context . helpers ; 13 14 it ( "ポッドを削除できるはずです" , async () => { ⋯ 15 // テストは相互に分離するために独自の一意の名前空間を取得します16 const 名前空間 = await getTestNamespace (); 17 18 await core .createNamespacedPod ({ ⋯ 19 名前空間 , 20 本体 : { ⋯ 21 メタデータ : { 名前 : "delete-test" } , 22 仕様 : { ⋯ 23 コンテナ : [ ⋯ 24 { ⋯ 25 名前: "pause" , 26 // Webernetes にはこのイメージの実装が組み込まれています 27 image : "registry.k8s.io/pause:3.10" , 28 } , 29 ] , 30 } , 31 } , 32 }); 35 await waitFor ( async. () => { ⋯ 36 const ポッド = await core .listNamespacedPod ({ namespace }); 37 const found = ポッド . メタデータ ?. name === "delete-test" ); 40 41 await core .deleteNamespacedPod ({ ⋯ 42 name : "delete-test" , 43 namespace , 44 }); 45 46 // ポッドが完全になくなるまで待ちます

[切り捨てられた]

## Original Extract

Almost 100,000 lines of LLM-generated code in 2 months, and none of it is slop.

I ported Kubernetes to the browser | ngrok blog Skip to main content ngrok home / ngrok blog home blog open mobile navigation Products Problems We Solve Resources Docs Pricing Log in Sign up Products Gateway
Access remote K8s clusters from dev
Connectivity Site-to-site connectivity
Native Anthropic SDK support for the AI Gateway
Make LLMs 4x smaller and 2x faster
Prompt caching: 10x cheaper LLM tokens, but how?
Search… Control ⌃ K Newsletter RSS Jun 30, 2026
I ported Kubernetes to the browser
Sam Rose is a Senior Developer Educator at ngrok, focusing on creating content that helps developers get the most out of ngrok.
Share I ported Kubernetes to the browser on hackernews
Share I ported Kubernetes to the browser on linkedin
Share I ported Kubernetes to the browser on twitter
Share I ported Kubernetes to the browser on reddit
Share I ported Kubernetes to the browser on whatsapp
Last week I released webernetes , a partial port of Kubernetes to TypeScript
to make it possible to run clusters in the browser. I ended up generating almost
100,000 lines of code in 552 commits across 629 files. It took me 2 months.
The demo below is a webernetes cluster. It runs entirely in your browser, and
it’s genuinely doing much of the same work a real Kubernetes cluster does: pod
lifecycles, cluster DNS and networking, container garbage collection, IP
allocation, Deployment and ReplicaSet tracking, and more. The blue dots
represent pods sending requests to each other.
Interactive Webernetes demo showing HTTP requests moving between three pods from a Deployment across three Kubernetes nodes.
Reset Pause Starting simulated cluster.
Bookmark this section Wait, what am I looking at?
The question I’ve been getting most often is: “Did you compile Kubernetes to
WebAssembly?” The answer is no. A simple “hello, world!” Go program compiled to
WebAssembly is ~540KiB gzipped. That alone is already bigger than webernetes,
which is ~140KiB gzipped. Compiling all of Kubernetes to WebAssembly would no
doubt mean sending megabytes over the wire. I did try to check, but
unfortunately there are compile-time errors because Kubernetes calls system-level
APIs that aren’t available in the browser.
A partial port of Kubernetes’ “kubelet” binary, enough to run pods and probe them.
Ports of several Kubernetes “controllers”: pod scheduler, namespace controller, kube-proxy, deployment controller, and a few more.
A browser-based take on a container network interface (CNI), so pods can talk to each other over a simulated network.
A browser-based container runtime, which the kubelet talks to over the container runtime interface (CRI) to run containers.
An API for interacting with your webernetes cluster to do things like apply manifests and watch resources.
As a result of the desire to keep webernetes small, it doesn’t pull real images
from a registry like Docker Hub. Instead, it has its own browser-based registry
and you define images using a TypeScript API. Images look like this:
1 import * as w8s from "@ngrok/webernetes" ; 2 3 class HelloWorld extends w8s . BaseImage { ⋯ 4 static readonly imageName = "hello-world" ; 5 static readonly imageVersion = "1.0" ; 6 7 async exec ( ctx : w8s . ProcessContext , argv : string []) : Promise < number > { ⋯ 8 ctx .listenHttp ( 8080 , async ( _ctx , request ) => { ⋯ 9 return { ⋯ 10 status : 200 , 11 body : "Hello, world!" , 12 }; 13 }); 14 return await ctx .waitUntilKilled (); 15 } 16 }
To deploy your image into a cluster, you do this:
1 import * as w8s from "@ngrok/webernetes" ; 2 3 class HelloWorld extends w8s . BaseImage { ⋯ 4 // as before ... 5 } 6 7 const cluster = new w8s .Cluster (); 8 await cluster .registerImage ( HelloWorld ); 9 10 const [ pod ] = await cluster .apply ([ ⋯ 11 { ⋯ 12 apiVersion : "apps/v1" , 13 kind : "Deployment" , 14 metadata : { name : "hello-world-deployment" } , 15 spec : { ⋯ 16 replicas : 1 , 17 selector : { ⋯ 18 matchLabels : { app : "hello-world-pod" } , 19 } , 20 template : { ⋯ 21 metadata : { ⋯ 22 labels : { app : "hello-world-pod" } , 23 } , 24 spec : { ⋯ 25 containers : [ ⋯ 26 { ⋯ 27 name : "hello-world-container" , 28 image : "hello-world:1.0" , 29 } , 30 ] , 31 } , 32 } , 33 } , 34 } , 35 ]);
And then you can use the webernetes API to interact with the cluster, like this:
1 // List all pods in the default namespace 2 const pods = await cluster . api . corev1 .listNamespacedPod ({ ⋯ 3 namespace : "default" , 4 }); 5 6 // Watch for changes to pods in all namespaces 7 const informer = cluster .informer ( "pods" , ( type , pod ) => { ⋯ 8 console .log ( `pod ${ type } : ${ pod . metadata ?. name } ` ); 9 }); 10 11 // Stop the informer when you're done 12 await informer .stop (); 13 14 // Listen to pods sending requests and responses to each other. 15 // This is how I visualise the moving dots above. 16 cluster .on ( "request" , ( event ) => { ⋯ 17 console .log ( `request: ${ event . request . method } ${ event . request . url } ` ); 18 }); 19 cluster .on ( "response" , ( event ) => { ⋯ 20 console .log ( `response: ${ event . response ?. status } ` ); 21 }); 22 23 // Use the cluster network to send a request to a pod. This will also trigger 24 // the request/response event handlers above. 25 const pod = pods . items [ 0 ]; 26 const resp = await cluster .fetch ( `http:// ${ pod . status ?. podIP } :8080/` ); 27 console .log ( resp . body ); // "Hello, world!"
There are plenty more examples in the webernetes repository .
Webernetes is intended to be used to make interactive Kubernetes content; it’s
not a production-ready Kubernetes distribution. It doesn’t need to run real
images. It just needs a way for creators to set up specific workloads to
illustrate the thing they’re trying to teach.
Over time, it is my intention to expand webernetes to support more Kubernetes
features. Right now, it doesn’t support ConfigMaps, Secrets, pod resources,
persistent volumes, and a whole host of other things I haven’t needed yet. As I
make more content with this library, I’ll implement more of what I need.
If you’re looking to build on webernetes and it doesn’t support something you
need, please reach out! I’m s.rose@ngrok.com and I’d
be happy to help you become a contributor.
Bookmark this section Is this just slop?
Almost all of the webernetes code was authored by LLMs. I expect people to be
dubious of the project as a result. I expect to be accused of slop-porting
Kubernetes for views, but I’m going to try to show you that’s not what I’ve
done.
I did two things that I think make this a slop-free project:
I reviewed every line of code.
I created hundreds of tests that assert webernetes behaves the same as a real cluster.
The first point, by far the most time-consuming, is how I gained the confidence
that the vast majority of the code is line-for-line identical to the Kubernetes
Go codebase. The second point is how I made sure the lexical similarity
translates to identical behaviour in practice.
Any mistakes that remain in the codebase after my review are on me, and I’ve no
doubt some exist. If you find any, please let me know by opening an issue .
Bookmark this section Why review the code?
The stories I’ve read about LLMs being used to write a C compiler or port
Bun from Zig to Rust were made possible by having an automated way to
assert correctness. Anthropic had plenty of existing C compilers to compare
against, and Bun had a large existing test suite that its maintainers trusted
enough to merge over 1 million lines of new Rust code without manual review.
I didn’t have those things. If I wanted a test suite, I’d need to write it
myself. If I wanted to compare against real Kubernetes, I’d need to figure out a
way to do it.
Most of the code in webernetes is ported from the Kubernetes Go codebase. I
ported it with LLMs because I was confident that would be faster than typing it
by hand, but the problem I quickly encountered was that LLMs suck at porting
code . No matter how hard I tried, they kept making mistakes. The mistakes
came in a few flavours:
Shortcuts. This seemed to happen more when porting larger files, and it made me wonder if I was battling with post-training choices designed to make LLMs less verbose. An example of a shortcut I saw a lot: Kubernetes contains a lot of different types of cache. There’s an LRU cache , an expiring cache , a FIFO cache , a transforming cache , and more. Instead of implementing these, I’d catch the LLM using a Map instead, leading to incorrect behaviour.
Trying to be too helpful. The LLM would try to tidy up code by inventing helper functions that didn’t exist in the original Go code. Often, these helpers were harmless, but sometimes they’d have subtle differences. Either way, they made the code harder for me to review side-by-side with the original, so I asked the LLM to remove them.
Stuff just… missing. This happened most often when porting over table tests from Go. If you’re unfamiliar, table tests look like this . They are arrays of test cases with some test code underneath. The LLM would arbitrarily omit tests and I’d have to ask it why. Sometimes it would talk about the test not being applicable. Occasionally that was true, but usually the LLM would own up to omitting the test by mistake.
I know at least a few of you are screaming “SKILL ISSUE” and are ready to
comment saying I need to get better at prompting. That could be true! I would
love to see an example prompt that perfectly one-shots porting this table
test from Go to TypeScript. You stand to save me an enormous amount of time
in future.
Until then, for me to have confidence in an LLM porting something, I need to
review the output. I’m not aware of any shortcuts.
Bookmark this section What the tests look like
It’s all well and good to know that the code is side-by-side identical, but does
it actually work ? Go and JavaScript have different runtime environments, so it
was always possible that the same code would behave differently in each. I also
ended up having to create JavaScript versions of channels , mutexes ,
Go’s select statement , and other Go-isms. I needed to know they worked in
non-trivial scenarios.
To feel good about this, I wrote tests where the exact same code is run against
both webernetes and a k3s cluster. To do this, I needed to have an API for
webernetes that matched an existing Kubernetes API. I picked
kubernetes-client/javascript because it’s the official client library for
Kubernetes in JavaScript, and it has TypeScript types.
1 import { expect , it } from "vitest" ; 2 import { kubernetes } from "../../test/harnesses/kubernetes" ; 3 4 // `kubernetes.describe` does some magic behind the scenes to set up either a 5 // k3s (https://k3s.io/) cluster or a webernetes cluster and pass that in via 6 // the `context` argument. 7 // 8 // Then I can run either `pnpm test:node` or `pnpm test:browser` to run tests 9 // against k3s using a Node environment, or webernetes using a headless browser. 10 kubernetes .describe ( "Pods" , ( context ) => { ⋯ 11 const { core } = context ; 12 const { getTestNamespace , waitFor } = context . helpers ; 13 14 it ( "should be able to delete a pod" , async () => { ⋯ 15 // Tests get their own unique namespace for isolation from each other 16 const namespace = await getTestNamespace (); 17 18 await core .createNamespacedPod ({ ⋯ 19 namespace , 20 body : { ⋯ 21 metadata : { name : "delete-test" } , 22 spec : { ⋯ 23 containers : [ ⋯ 24 { ⋯ 25 name : "pause" , 26 // Webernetes has an implementation of this image built-in 27 image : "registry.k8s.io/pause:3.10" , 28 } , 29 ] , 30 } , 31 } , 32 }); 33 34 // Make sure the pod definitely exists before moving on. 35 await waitFor ( async () => { ⋯ 36 const pods = await core .listNamespacedPod ({ namespace }); 37 const found = pods . items .find (( pod ) => pod . metadata ?. name === "delete-test" ); 38 expect ( found ) .toBeDefined (); 39 }); 40 41 await core .deleteNamespacedPod ({ ⋯ 42 name : "delete-test" , 43 namespace , 44 }); 45 46 // Wait until the pod is definitely gone

[truncated]
