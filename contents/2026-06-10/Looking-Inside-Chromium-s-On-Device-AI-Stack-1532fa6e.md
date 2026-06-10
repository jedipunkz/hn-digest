---
source: "https://www.island.io/blog/looking-inside-chromiums-on-device-ai-stack"
hn_url: "https://news.ycombinator.com/item?id=48478721"
title: "Looking Inside Chromium's On-Device AI Stack"
article_title: "Looking Inside Chromium’s On-Device AI Stack"
author: "wild_pointer"
captured_at: "2026-06-10T19:03:15Z"
capture_tool: "hn-digest"
hn_id: 48478721
score: 3
comments: 0
posted_at: "2026-06-10T16:25:57Z"
tags:
  - hacker-news
  - translated
---

# Looking Inside Chromium's On-Device AI Stack

- HN: [48478721](https://news.ycombinator.com/item?id=48478721)
- Source: [www.island.io](https://www.island.io/blog/looking-inside-chromiums-on-device-ai-stack)
- Score: 3
- Comments: 0
- Posted: 2026-06-10T16:25:57Z

## Translation

タイトル: Chromium のオンデバイス AI スタックの内部を見る
記事タイトル: Chromium のオンデバイス AI スタックの内部を見る
説明: Gemini Nano および WebNN への最適化ガイドから、Chromium がローカル AI モデルを使用する方法と、Island がエンタープライズ グレードの AI ガバナンスを可能にする方法についての技術的な詳細。

記事本文:
Chromium のオンデバイス AI スタックの内部を見る
-->
ここではpaint0_linear_5019_99065とpaint1_radial_5019_99065を使用します -->
製品
エンタープライズブラウザ
企業の業務に必要なすべてのシステムを単一の統合環境に統合
安全でアクセスしやすく、職場での実際の AI の使用方法に合わせて構築されています
スケーラブルで回復力があり、エンドユーザー向けに設計された最新の SASE
Island でイノベーションを起こしている世界中の組織の声を聞く
エンタープライズ ブラウザ購入者ガイド
Island が AI を企業向けに活用
ブログに戻る 15 分で読む 2026 年 6 月 9 日 |更新: Chromium のオンデバイス AI スタックの内部を見る
最新のブラウザは、強力な AI エンジンへと静かに進化しました。 Chromium のオンデバイス機械学習スタックの基礎となる層とそれを駆動する仕組みを見てみましょう。
Tal Marian 氏、シニア ソフトウェア エンジニア
煩わしい通知の削減から翻訳の提供まで、Web を閲覧するたびに、AI モデルの高度なシステムがブラウザーの舞台裏で動作します。この記事では、Chromium でローカル AI モデルが「なぜ」「どのように」使用されるのかについて技術的に詳しく説明します。 Island Enterprise Browser がその基盤に基づいて AI を安全に職場に導入する方法について説明します。
Chromium がローカル AI モデルを使用するのはなぜですか?
一般に、AI モデルをローカルで実行することについて話すと、次のような利点があります。
プライバシー: モデルによって処理されたデータはユーザーのデバイスから流出することはなく、機密情報のプライバシーが保たれます。
レイテンシー: デバイス上でモデルを実行すると、ネットワークのレイテンシーが解消され、結果が即時に得られ、ユーザー エクスペリエンスがよりスムーズになります。
信頼性: ローカル モデルを利用した機能は、ユーザーがオフラインの場合やネットワーク接続が不安定な場合でも機能し続けます。
コスト: 推論をローカルで実行すると、計算コストがクラウドからオフロードされます。

サーバー。
次に、Chromium がローカル AI を実際にどのように機能させるかを見てみましょう。ブラウザでの使用法は、次の 2 つのグループに分類できます。
Chromium の内部使用: Chromium がブラウザ自体で使用するものを実装するために使用する AI モデル。これらは最適化ガイドによって管理されます。これについては後ほど詳しく説明します。
AI サービス : Web 開発者が Web ページに AI ベースの機能を実装するための API。注目すべき例としては、新しく追加された Summarizer API と WebNN があります。これらについても説明します。
どちらの場合も、Chrome は Google サーバーから AI モデルをダウンロードできます。これらはオンデマンドでダウンロードされます。つまり、それらを必要とする機能がトリガーされない限り、デバイスにはダウンロードされません。
Chromium と Chrome は似ていますが、同じものではないことに注意してください。 Chromium はオープンソース プロジェクトですが、Chrome は Chromium をベースにした Google 製品です。一部の機能は Chrome を通じてのみ利用可能であり、ライセンスの関係で通常の Chromium には同梱されていません。
社内向けAI 最適化ガイド
最適化ガイドは、Chromium の内部で使用されるローカル AI モデルのダウンロード、管理、実行を担当する Chromium サービスの名前です。
なぜ「最適化ガイド」と呼ばれるのでしょうか?名前はほとんどが歴史上の産物です。当初 (2017 年頃) は、JS を事前に無効にする、ロードするリソースを提案する、軽量ページ バージョンにリダイレクトするなどのヒントをブラウザーに提供することでページを最適化することを目的としていました。このメカニズムに AI が初めて追加されたのは 2019 年頃で、ページの読み込みが「痛みを伴う」かどうかを予測するモデルが追加されました。それ以来、モデルを管理および実行するための同じインフラストラクチャが追加のタスクに再利用され、その多くはページの最適化とは何の関係もありません。
Op が実行するすべてのタスクをリストすることができます。

最適化ガイドは、Chromium のソースに含まれる models.proto ファイルを参照することでサポートできます。 proto ファイルには、AI モデルのダウンロードをリクエストするために Chromium が Google のバックエンドと結んだ契約を定義する protobuf 宣言が含まれています。このファイルには、約 70 個のエントリを含む OptimizationTarget 列挙型が表示されます。
これらのエントリを見ることで、最適化ガイドの主なトピックが何であるかを理解できます。
ページコンテンツ分析
PAGE_TOPICS_V2、EDU_CLASSIFIER ‍
安心・安全
CLIENT_SIDE_PHISHING、SCAM_DETECTION ‍
ユーザーのセグメンテーション
SEGMENTATION_SHOPPING_USER、SEGMENTATION_VOICE ‍
オムニボックスの注文と提案
OMNIBOX_URL_SCORING、
OMNIBOX_ON_DEVICE_TAIL_SUGGEST ‍
許可承認の予測 ‍
NOTIFICATION_PERMISSION_PREDICTIONS ‍
自動入力の改善 ‍
PASSWORD_MANAGER_FORM_CLASSIFICATION
この列挙型のすべてのシナリオが実際に機械学習モデルによって実装されているわけではなく、一部のシナリオは自明な if チェーンによって実装されていることに注意してください。例: 新しいタブ ページ フィードを使用するユーザーを分類するための SEGMENTATION_FEED_USER は、フィードが少なくとも 2 回エンゲージされたかどうかの単純なチェックによって実装されます。
ユースケースのスポットライト: NOTIFICATION_PERMISSION_PREDICTIONS
最適化ガイドには多くのモデルがあるため、何が起こっているのかを理解するために、OPTIMIZATION_TARGET_NOTIFICATION_PERMISSION_PREDICTIONS という特定の例に焦点を当ててみましょう。
このモデルは、ユーザーが Web ページからの通知送信リクエストを受け入れる可能性が高いかどうかを判断するために使用されます。ユーザーが VERY_UNLIKELY とみなされる場合、Chromium は通常の通知リクエスト プロンプトの代わりに、静かなインジケーターを表示します。
Chromium は、この機能について (技術的ではない) ブログ投稿も書きました。ここでそれを読んだり、いくつかの写真を見ることができます。
Th

このモデルを呼び出すクラスは PermissionsAiUiSelector と呼ばれます。その実装を見ると、このモデルへの入力が、通知のアクセス許可とアクセス許可一般に関するさまざまな統計 (浮動小数点と整数で表される) であることがわかります。たとえば、入力の一部は次のとおりです。
平均許可/拒否/却下/無視率
プロンプトの総数 (すべてのタイプ)
プロンプトの総数 (特に通知)
これらはすべて、過去 28 日間の範囲でカウントされます。これは、コード内にハードコーディングされた値です。これを変更するには、モデルを再トレーニングする必要があります。
これらの値はローカル モデルに供給されて、判定が生成されます。このモデルはディスク上に保持され、他の多くの最適化ガイドのターゲットと同様に、 model.tflite という名前が付けられます。ファイル拡張子が示すように、Chromium は、エッジ デバイス上でモデルを高パフォーマンスで実行するために Google によって開発された TFLite フレームワークを使用します。 TFLite は現在 LiteRT に取って代わられ、Chromium はこの新しいフレームワークへの移行過程にあることは注目に値します。
補足: 最適化ガイドのすべてのモデルは、GPU (または NPU などの専用 AI ハードウェア) ではなく CPU で実行されます。すぐにわかりますが、これは、最適化ガイドで使用されるモデルの種類にとっては合理的であり、非常に軽量になる傾向があります。
この場合、モデルのディスク上のサイズはわずか 31 KB です。 netron.app などのツールを使用して視覚化できます。
何が起こっているのかを簡単に見てみましょう:
ブラウザからの権限統計はモデル機能 (または入力) にマッピングされます。
各フィーチャは中央に配置され、拡大縮小され、欠落している場合はマスクされます (上のセクションを参照)。
正規化された特徴値は、特徴間の相互作用を学習する 2 つの (非線形) 高密度層を通過します。
たとえば、無視率と通知総額の関係

アクション数とプラットフォーム (モバイル/デスクトップ)
最後の線形層はロジットを生成し、それが単一の確率出力に変換されます。
これらすべては、ブラウザーが表示する通知リクエストの種類を決定するたびに、デバイス上で行われます。
このモデルの実際のインテリジェンスは、重みとバイアスを生成するために Google が実行したトレーニングから得られます。残念ながら、そのトレーニング データは公開されていません。
補足: 特にこの使用例の場合、最適化ガイドは代わりに、ページをテキスト埋め込みとして受け取るクラウドベース (非ローカル) モデルを実行できます。これは、ユーザーが Chromium 設定で「検索とブラウジングを改善する」を有効にしている場合に発生します。デフォルトではオフになっています。
この時点で、これら 70 以上の最適化ターゲットのどれが実際にデバイス上にあるのか疑問に思うかもしれません。
Chromium には、デバッグを目的とした内部ページが多数あります。最適化ガイドに関する情報は、 chrome://optimization-guide-internals にあります。これを開いたら、「モデル」タブに切り替えて、ブラウザーが既にダウンロードしているすべてのモデルを表示します。私のブラウザには現在 13 個のエントリがあります。例:
Chromium コンポーネントが最適化ガイドからモデルを要求すると、Google のサーバーからモデルをダウンロードし、 optimization_guide_model_store というフォルダーに保存します。モデルはバックグラウンドでブラウザによって自動的に更新されます。
ユーザー セグメンテーション (これも最適化ガイドの一部です) に興味がある場合は、代わりに chrome://segmentation-internals を参照してください。
最適化ガイドによって維持されている軽量モデルに加えて、Google は Gemini Nano と呼ばれる最先端の小型言語モデルを Chrome に追加しました。これは Chrome の組み込み AI イニシアチブの一部であり、こちらで読むことができます。

この取り組みでは、開発者がデバイス上で AI タスクを実行するための Web API を追加します。現在、最もサポートされている (ただし、まだあまり広くサポートされていない) API は、ローカル AI を使用してテキストを要約する Summarizer API です。 Chrome では、同様のインターフェースを通じてローカル言語の検出と翻訳も提供します。閲覧するサイトがこれらの API のいずれかを呼び出すと、Gemini Nano がデバイスに自動的にダウンロードされます。
Summarizer API は、Chrome および Edge バージョン 138 以降で利用できます。 Edge では、内部的には Chromium ですが、この API は Gemini Nano ではなく Microsoft の Phi 4 Mini モデルを使用します。
これらのモデルは非常に大きく、実行するには数 GB のディスク領域とメモリを必要とします。したがって、ハードウェアと利用可能なディスク容量に応じて、互換性のあるデバイスにのみダウンロードされます。
この機能に関する情報は、 chrome://on-device-internals でご覧いただけます。ブラウザがすでにモデルをダウンロードしている (つまり、サイトがモデルを要求している) 場合は、モデルの状態、名前、ディスク使用量を確認できます。
これまで最適化ガイドで見てきた TFLite 環境とは異なり、これらのモデルはデバイス ハードウェアを最大限に活用し、より効率的な実行のために GPU を利用します。
残念ながら、これらのモデルの実行に関連するコードは、Chromium オープン ソース プロジェクトの一部ではありません。関連クラスである ChromeML の実装は、Chromium の git ツリーにコミットされていないサブモジュール内にあります。これは、オープンソース プロジェクトである Chromium と Google 製品である Google Chrome の違いの一部です。
ChromeML のパブリック API を調べてバイナリをリバース エンジニアリングすると、LiteRT と MediaPipe への参照が確認できるため、このバイナリは同様のテクノロジーを使用してモデルを実際に実行していると考えるのが妥当です。
chrome://on-device-

内部ページには、このモデルが Chrome に保存されているディレクトリ (私の場合は ~/Library/Application Support/Google/Chrome Canary/OptGuideOnDeviceModel/2025.8.8.1141 ) も表示されます。通常のmodel.tfliteとは異なり、これはモデル自体を含むweights.binファイルと、モデル名およびその他の情報を含むmetadata.jsonファイルを含む専用フォルダーです。
このページでは、特定の Web API の制約の外で LLM と対話することもできます。これを行うには、「デフォルトをロード」をクリックするだけです。ロードされたら、「入力」フィールドにプロンプ​​トを書き込み、「実行」をクリックします。
Gemini Nano は汎用 LLM であると考えているかもしれません。要約などの特定のタスクを実行するようにどのようにガイドされますか?ソース コードのどこかにシステム プロンプトが必要です…そして、それは正しいでしょう。何をすべきかをガイドするシステム プロンプトが必要ですが、そのシステム プロンプトは Chromium のソースには含まれていません。代わりに、最適化ガイドによっても管理およびダウンロードされます。
以前に確認したように、すべての最適化ターゲットが実際にはモデルであるわけではありません。これは、ターゲット OPTIMIZATION_TARGET_MODEL_EXECUTION_FEATURE_SUMMARIZE に特に当てはまります。フォルダー内に model.tflite ファイルはありますが、完全に空です。代わりに、 on_device_model_execution_config.pb というファイルがあります。

[切り捨てられた]

## Original Extract

A technical deep dive into how Chromium uses local AI models, from Optimization Guide to Gemini Nano and WebNN, and how Island enables enterprise-grade AI governance.

Looking Inside Chromium’s On-Device AI Stack
-->
with paint0_linear_5019_99065 and paint1_radial_5019_99065 here -->
Product
Enterprise Browser
All the systems enterprise work requires in a single, unified environment
Secure, accessible, and built for how the workplace actually uses AI
Modern SASE that’s scalable, resilient, and designed for the end user
Hear from organizations around the world innovating with Island
Enterprise Browser Buyer's Guide
Island Makes AI Work for the Enterprise
Back to blog 15 min read June 9, 2026 | Updated: Looking Inside Chromium’s On-Device AI Stack
The modern browser has quietly evolved into a powerful AI engine. Let's explore the underlying layers of Chromium’s on-device machine learning stack and the mechanics driving it.
Tal Marian , Senior Software Engineer
From reducing annoying notifications to providing translations, a sophisticated system of AI models is working behind the scenes in your browser every time you browse the web. In this article, we’ll take a technical deep dive into the “why” and “how” local AI models are used in Chromium. We’ll go over how the Island Enterprise Browser builds on that foundation to bring AI safely into the workplace.
Why does Chromium use local AI models?
In general, when we talk about running AI models locally, there are a few benefits:
Privacy: Data processed by the model never leaves the user's device, keeping sensitive information private.
Latency: Executing the model on-device removes network latency, leading to immediate results and a smoother user experience.
Reliability: Features powered by local models remain functional even when the user is offline or has an unstable network connection.
Cost: Running inference locally offloads computational cost from cloud servers.
Now let's look at how Chromium actually puts local AI to work. Its usage in the browser can be divided into two groups:
Internal Chromium usage: AI models that Chromium uses to implement something for the browser itself to use. They’re managed by Optimization Guide, which we’ll dive into shortly.
AI services : APIs meant for web developers to implement AI-based features in their web pages. Some notable examples are the newly-added Summarizer API and WebNN, which we will also cover.
For both of these, Chrome can download the AI models from Google servers. These are downloaded on demand, meaning that they won’t be on your device unless a feature that needs them is triggered.
Note that while Chromium and Chrome sound similar, they are not the same thing. Chromium is an open-source project, while Chrome is the Google product that is based on Chromium. Some features are only available through Chrome and don’t ship with vanilla Chromium due to licensing.
AI for internal use - Optimization Guide
Optimization Guide is the name of the Chromium service responsible for downloading, managing, and executing local AI models used internally in Chromium.
Why is it called “Optimization Guide”? The name is mostly a historical artifact. Initially (circa 2017), it was meant for optimizing pages by giving the browser hints, such as disabling JS in advance, suggesting which resources to load, or redirecting to a lite page version. The first addition of AI to this mechanism occurred around 2019, when a model for predicting if a page load will be ‘painful’ was added. From then on, the same infrastructure for managing and running the models was reused for additional tasks, many of which have nothing to do with page optimization.
We can list all of the tasks that the Optimization Guide can support by looking at the models.proto file included in Chromium’s source here . A proto file contains protobuf declarations that define the contract Chromium has with Google’s backend for requesting the download of an AI model. In this file, we can see the OptimizationTarget enum, which contains around 70 entries.
By looking at these entries, we get a sense of what the main topics in the Optimization Guide are:
Page content analysis
PAGE_TOPICS_V2, EDU_CLASSIFIER ‍
Security and safety
CLIENT_SIDE_PHISHING, SCAM_DETECTION ‍
User segmentation
SEGMENTATION_SHOPPING_USER, SEGMENTATION_VOICE ‍
Omnibox ordering and suggestions
OMNIBOX_URL_SCORING,
OMNIBOX_ON_DEVICE_TAIL_SUGGEST ‍
Permission approval predictions ‍
NOTIFICATION_PERMISSION_PREDICTIONS ‍
Autofill improvements ‍
PASSWORD_MANAGER_FORM_CLASSIFICATION
It is worth noting that not all scenarios in this enum are actually implemented by machine learning models—some of them are implemented by trivial if chains. For example: SEGMENTATION_FEED_USER , which is meant to categorize users who use the new tab page feed, is implemented by simple checks if the feed was engaged with at least twice.
Use case spotlight: NOTIFICATION_PERMISSION_PREDICTIONS
There are many models in the Optimization Guide, so let’s focus on a specific example to get a sense of what’s happening: OPTIMIZATION_TARGET_NOTIFICATION_PERMISSION_PREDICTIONS.
This model is used for determining whether the user is likely to accept a request from the web page to send notifications. If the user is deemed as VERY_UNLIKELY , Chromium shows a quieter indicator instead of the usual notification request prompt.
Chromium also wrote a (non-technical) blog post about this feature; you can read it and see some pictures here .
The class that invokes this model is called PermissionsAiUiSelector . By looking at its implementation , we can see that the inputs to this model are various statistics (represented as floats and integers) about notification permissions and permissions in general—for example, some of the inputs are:
Average grant/deny/dismiss/ignore rates
Total number of prompts (all types)
Total number of prompts (notifications specifically)
All of these are counted in the span of the last 28 days - this is a hardcoded value in the code. To change it, the model will need to be retrained.
These values are fed to the local model to produce a verdict. This model is kept on disk, and it is named model.tflite , like many other Optimization Guide targets. As the file extension suggests, Chromium uses the TFLite framework, which was developed by Google for high-performance execution of models on edge devices. It is worth noting that TFLite is now superseded by LiteRT , and Chromium is in the process of migrating to this newer framework.
As a side note: All of the models in the Optimization Guide are executed on the CPU, not the GPU (or dedicated AI hardware like NPU). As we will soon see, this is reasonable for the type of models used by the Optimization Guide, which tend to be quite lightweight.
In this case, the model is just 31KB on disk. We can use tools like netron.app to visualize it:
Let’s take a quick look at what’s happening:
The permission statistics from the browser are mapped into the model features (or inputs)
Each feature is centered, scaled, and masked if missing (as shown in the top section)
The normalized feature values are passed through two (nonlinear) dense layers that learn interactions between features
For example, the relationship between ignore rates, total notification counts, and the platform (Mobile/Desktop)
A final linear layer produces a logit, which is converted into a single probability output
All of this happens on-device, each time the browser wants to determine what type of notification request to show.
The real intelligence of this model is derived from the training that Google performed to produce its weights and biases. Unfortunately, that training data is not publicly available.
Side note : For this use case specifically, the Optimization Guide can alternatively run a cloud-based (non-local) model, which receives the page as text embeddings . This happens if the user enables “Make searches and browsing better” in Chromium settings. It is off by default.
At this point, you may be wondering which of these 70+ optimization targets are actually on your device.
Chromium has many internal pages for debugging purposes. Information about the Optimization Guide can be found in chrome://optimization-guide-internals . When you open it, switch to the “Models” tab to see all of the models that the browser has already downloaded. On my browser, there are currently 13 entries, for example:
When a Chromium component asks for a model from the Optimization Guide, it downloads the model from Google’s servers and keeps it in a folder called optimization_guide_model_store . The models are automatically updated by the browser in the background.
If you’re interested in user segmentation, which is also part of the Optimization Guide, you can view chrome://segmentation-internals instead.
In addition to the lightweight models maintained by the Optimization Guide, Google added a state-of-the-art small language model into Chrome called Gemini Nano . This is part of Chrome’s Built-in AI initiative, which you can read about here .
This initiative adds web APIs for developers to run AI tasks on-device. Right now, the most supported (but still not very widely supported) API is the Summarizer API , which summarizes text using local AI. Chrome also offers local language detection and translation through a similar interface. Gemini Nano is automatically downloaded to your device once a site you browse to invokes one of these APIs.
The Summarizer API is available on Chrome and Edge versions 138 and up. On Edge, which is also Chromium under the hood, this API uses Microsoft’s Phi 4 Mini model instead of Gemini Nano.
These models are quite large, taking up multiple GB of disk space and memory to execute. Therefore, they are only downloaded on compatible devices, depending on the hardware and disk space available.
Information about this feature can be seen in chrome://on-device-internals . You can see the state, name, and disk usage of the model if your browser has already downloaded it (meaning: a site requested it).
Unlike the TFLite environment that we’ve seen in the Optimization Guide so far, these models take full advantage of the device hardware, making use of the GPU for more efficient execution.
Unfortunately, the code relevant for the execution of these models is not part of the Chromium open source project. The implementation of ChromeML , which is the relevant class, is in a submodule that is not committed to Chromium’s git tree. This is part of the distinction between Chromium , the open-source project, and Google Chrome , the Google product.
By looking at the public API of ChromeML and reverse-engineering the binary, we can see references to LiteRT and MediaPipe, so it is reasonable to assume that this binary uses similar technology to actually execute the model.
The chrome://on-device-internals page also shows us the directory in which this model is stored in Chrome, ~/Library/Application Support/Google/Chrome Canary/OptGuideOnDeviceModel/2025.8.8.1141 , in my case. Unlike the usual model.tflite, this is a dedicated folder with a weights.bin file that contains the model itself, and a metadata.json file with the model name and other information.
This page also allows us to interact with the LLM outside the constraints of a specific Web API. To do so, we just need to click “Load Default”. Once it is loaded, we can write a prompt in the “Input” field and click “Execute”.
You may be thinking, Gemini Nano is a general-purpose LLM; how is it guided to perform specific tasks like summarization? It must have a system prompt somewhere in the source code… And you’d be right! It does need a system prompt to guide it on what to do, but that system prompt is NOT in Chromium’s source. Instead, it is also managed and downloaded by the Optimization Guide.
As we’ve previously established, not all optimization targets are actually models. That’s especially true for the target OPTIMIZATION_TARGET_MODEL_EXECUTION_FEATURE_SUMMARIZE . It does have a model.tflite file in its folder, but it’s completely empty. Instead, it has a file called on_device_model_execution_config.pb , which

[truncated]
