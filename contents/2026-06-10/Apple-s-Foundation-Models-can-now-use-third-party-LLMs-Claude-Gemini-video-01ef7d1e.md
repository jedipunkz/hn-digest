---
source: "https://developer.apple.com/videos/play/wwdc2026/241/"
hn_url: "https://news.ycombinator.com/item?id=48470772"
title: "Apple's Foundation Models can now use third-party LLMs (Claude, Gemini) [video]"
article_title: "What’s new in the Foundation Models framework - WWDC26 - Videos - Apple Developer"
author: "nryoo"
captured_at: "2026-06-10T04:35:07Z"
capture_tool: "hn-digest"
hn_id: 48470772
score: 2
comments: 1
posted_at: "2026-06-10T02:51:42Z"
tags:
  - hacker-news
  - translated
---

# Apple's Foundation Models can now use third-party LLMs (Claude, Gemini) [video]

- HN: [48470772](https://news.ycombinator.com/item?id=48470772)
- Source: [developer.apple.com](https://developer.apple.com/videos/play/wwdc2026/241/)
- Score: 2
- Comments: 1
- Posted: 2026-06-10T02:51:42Z

## Translation

タイトル: Apple の Foundation モデルでサードパーティ LLM が使用できるようになりました (Claude、Gemini) [ビデオ]
記事のタイトル: Foundation Models フレームワークの新機能 - WWDC26 - ビデオ - Apple Developer
説明: 何を調べるか

記事本文:
Foundation Models フレームワークの新機能 - WWDC26 - ビデオ - Apple Developer
英語で見る
Foundation Models フレームワークの新機能
Foundation Models フレームワークの新機能を確認してください。プライベート クラウド コンピューティングにアクセスし、サードパーティおよびオープン ソース モデルを統合し、ビジョン機能を使用する方法を学びます。コンテキスト管理 API、組み込みのセマンティック検索、アプリでエージェント エクスペリエンスを作成するための強力なプリミティブを発見します。
3:21 - 視覚: 画像の理解
6:46 - モデル抽象化レイヤー
7:32 - パートナー モデルの統合
9:40 - システムツール: ビジョンとスポットライト
10:57 - エージェント アプリの動的プロファイル
13:46 - モデルと構成の構成
16:02 - fm コマンドライン ツール
17:13 - 基礎モデル Python SDK
17:55 - オープンソースとフレームワークのユーティリティ
ツール呼び出しによる世代の拡大
マルチモーダル プロンプトによる画像の分析
命令とプロファイルを使用した動的セッションの構成
プライベート クラウド コンピューティングによるサーバー側インテリジェンスの追加
このビデオを検索…
みなさんこんにちは、ようこそ！私の名前はエリックです！私の名前はジェンです！昨年、私たちはガイド付き生成、スナップショット ストリーミング、強力なツール プロトコルなどの機能を備えた Foundation Models フレームワークを導入しました。私たちは、初年度のファウンデーション モデル フレームワークに対する皆様の興奮に驚かされました。そして、今年私たちがラインナップしたものをさらに気に入っていただけると思います。今年のフレームワークの新機能をすべて詳しく見ていきます。そして、言っておきますが、このリリースには内容が満載です! 2027 年のリリースでは、OS 内外の統合、幅広いモデル、エージェント エクスペリエンスを構築するための新しいプリミティブがすべてです。最もエキサイティングなアップデートの 1 つから始めましょう。 Foundation Models フレームワークには、多くの機能が含まれます。

本日発表する新しい API はオープンソースになります。そしてそれをスタイリッシュに！コア フレームワークに加えて、新しいパッケージ、Foundation Models フレームワーク ユーティリティもリリースします。これは、OS リリースごとに更新され、新しいビルディング ブロックや実験的なビルディング ブロックにアクセスできるようになります。そして、このセッション中に、Foundation Models フレームワーク エコシステムに参加する他の複数のパッケージについて聞くことになります。一緒に参加する準備はできていますか?フレームワークに追加する新しいモデル、新しいモダリティ、新しいツールに関連するすべてを説明します。そして、それらを最大限に活用できるように私たちが作成した新しい API を紹介します。スケジュールはいっぱいあります！まず、オンデバイス モデルの更新やサーバー モデルへのアクセスなど、モデルの更新を山ほど用意しました。
また、Spotlight と Vision フレームワークの助けを借りてセッションを強化する新しいシステム ツールも追加しました。その後、Zhen が再び参加して、エージェント エクスペリエンスを作成するための新しいプリミティブである動的プロファイルと呼ばれる強力な新しい API について説明します。 Zhen は、新しい評価フレームワークとその基盤モデルとの緊密な統合についても詳しく説明します。
最後に、Mac 固有の生産性向上ツールに関する興味深いニュースをいくつかご紹介しますので、ぜひ最後までお付き合いください。新しいモデルを早速使ってみましょう!
このリリースには、徹底的に再構築され、全面的に改良された新しいオンデバイス モデルが付属しています。よりインテリジェントです。ロジックとツールの呼び出しが得意です。 iOS 26.4 では、モデルのコンテキスト サイズを検査し、命令、プロンプト、トランスクリプト内のトークンをカウントするための新しい API をリリースしました。今後、これらを使用して、アプリを実行しているハードウェアに適応させることができます。私たちも頑張ってきました

ガードレールの改良に取り組んでいます。 iOS 26.4 では誤検知の数を減らすための調整が行われていることにお気付きかもしれませんが、iOS 27 ではさらなる改善が続けられています。
さらに、オンデバイス モデルには Vision 機能も追加されており、アプリケーションの新しいカテゴリ全体が可能になります。
API はシンプルで、既存のプロンプト ビルダーを自然に拡張したものです。
ここでセッションを作成しました。右側の折り紙の写真について質問したいと思います。
画像添付ファイルをテキストとともにプロンプ​​トに挿入するだけです。
これで、モデルは画像に関する質問に答えることができます。
画像添付ファイルは、次のようなさまざまなタイプから作成できます。 UIImage、NSImage、CGImage コア イメージ タイプ、CoreVideo ピクセル バッファ、およびファイル URL。
このモデルはあらゆるサイズとアスペクト比の画像をサポートしているため、特定の形状にトリミングしたりパディングしたりする必要はありません。
任意の画像サイズを使用できますが、画像が大きいほど、より多くのトークンを消費し、待ち時間が長くなることに注意してください。これらすべてのアップグレードのおかげで、オンデバイスのシステム言語モデルの機能がこれまで以上に向上しました。ただし、さらに多くの馬力が必要な場合は、新しい PrivateCloudComputeLanguageModel へのアクセスを提供します。プライベート クラウド コンピューティング モデルは、皆さんがよく知っている Apple Intelligence 機能の多くを強化するものとまったく同じです。これはオンデバイス モデルよりもはるかに大きなモデルで、32,000 トークンのコンテキスト ウィンドウがあり、強力な新機能である推論が付属しています。推論モデルは、応答を提供する前に時間をかけて慎重に答えを検討するようにトレーニングされており、その結果、結果が大幅に向上します。プライベート クラウド コンピューティングの使用は、これまでにないほど簡単です。モデルのインスタンスを作成し、それを使用して言語モデル セッションを初期化するだけです。セッションをプロンプトするときに、指定できるようになりました

新しい contextOptions 引数の推論レベルを変更します。 reasoningLevel は、モデルが応答する前にどれだけ考えることができるかを制御します。深い推論により、追加のコンピューティングと引き換えに、より良い応答が生成されます。
Private Cloud Compute の最も優れた点の 1 つは、アカウントのセットアップについて心配する必要がなく、認証に対処する必要がなく、API キーを保存する必要がなく、すべてが完全にシームレスであることです。そしてもちろん、Private Cloud Compute は何よりもプライベートです。プロンプトは決して保存されず、独立した研究者がこれらの主張を検証できるようにしています。さらに、Private Cloud Compute により、Foundation Models フレームワークを watchOS に導入できるようになります。 watchOS 27 以降では、最も強力なインテリジェンス機能を手首に装着できるようになりました。
PCC は、初回ダウンロード数が 200 万未満の開発者はクラウド API コストなしで利用できます。ユーザーは毎日 PCC にアクセスできるようになり、iCloud+ に加入している場合、制限はさらに高くなります。 PrivateCloudComputeLanguageModel を使用するために必要な資格など、PrivateCloudComputeLanguageModel の詳細について詳しくは、Private Cloud Compute を使用した構築に関するビデオを必ずご覧ください。
全面的に刷新されたオンデバイス モデルと新しいプライベート クラウド コンピューティング モデルに加えて、モデル抽象化レイヤーをオープンして、ほぼすべての言語モデルを Foundation Model のフレームワークで使用できるようにします。抽象化レイヤーは、ローカル モデルとサーバー モデルの両方が LanguageModelSession をサポートできるようにする新しい LanguageModel プロトコルを中心に構築されています。 SystemLanguageModel や PrivateCloudComputeLanguageModel などの既存のモデルは、すでにこのプロトコルに準拠しています。さらに、次の 2 つの追加実装をオープンソース化しています: CoreAILanguageM

odel と MLXLanguageModel。Mac の GPU の Apple Neural Engine で無数のローカル モデルを実行します。
また、さまざまなフロンティア サーバー モデルに確実にアクセスできるように、私たちが舞台裏で懸命に取り組んできたことを共有できることを嬉しく思います。 Anthropic と Google は両方とも、最新かつ最高のモデルへのアクセスを提供する Swift パッケージを公開しています。モデル抽象化レイヤーにより、サードパーティ モデルの使用が簡単になります。 Swift Package Manager を使用して言語モデル パッケージをインポートし、使用するモデルを初期化し、セッションの作成時にそれを渡します。
下流側はすべて同じままです。
サードパーティのサーバー モデルを使用する場合は、おそらく認証と請求の両方に対処する必要があることに注意してください。秘密キーをアプリのバイナリに保存しないでください。 OAuth などの安全なメカニズムを使用してアクセス トークンを常に取得し、KeyChain を使用して安全に保存します。開発者は通常、サードパーティ モデルを使用するときにトークンごとに料金が請求されるため、使用状況を簡単に追跡できるようにしました。セッションと応答には、使用されたトークンの数を正確に示す使用状況プロパティが追加されました。また、キャッシュから読み取られた入力トークンの数と、推論に使用された応答トークンの数を確認することもできます。
LanguageModel の使用方法、または独自の LanguageModel パッケージを作成する方法について詳しく知りたい場合は、「LLM プロバイダーを Foundation Models フレームワークに導入する」を参照してください。ついにすべてのモデルアップデートを完了しました！次はシステムツールです！このリリースでは、システムが提供する機能で LanguageModelSessions を強化するいくつかの組み込みツールを導入しています。 FoundationModels には、Vision フレームワークの強力な機能をサポートする 2 つのネイティブ ツールが含まれるようになりました。
BarcodeReaderTool も

モデルはバーコードから情報を読み取り、OCRTool を使用すると、モデルは画像から構造化テキストを抽出できます。どちらも、ネイティブでは不可能な方法で視覚情報を推論するモデルの能力を強化します。これらのツールの活用方法については、「画像理解の新機能」ビデオで詳しく説明されているので、詳細についてはビデオを参照してください。同様に、完全にローカルな検索拡張生成を実装するために、Spotlight を利用した検索ツールも導入します。これは、最も要望の多かった機能の 1 つです。検索拡張生成 (RAG) は、Spotlight インデックスと特別に処理されたクエリを活用することで、モデルが最新の個人またはドメインの知識にアクセスできるようにする手法です。これがまさにあなたが待ち望んでいたものだと思われる場合は、「Core Spotlight を使用した LLM 検索」が監視リストの一番上にあるはずです。このリリースのすべての新しいモデルと、SDK に追加する新しいシステム ツールを確認したので、Zhen に代わって、エージェント アプリ エクスペリエンスを構築するための新しい API についてすべて説明します。興奮してください！奪ってください、ジェン！エリックさん、ありがとう！
エージェント エクスペリエンスを構築するための新しいプリミティブである動的プロファイルを紹介します。
まずは工芸アプリを見て、動的プロファイルがどのような体験を可能にするかを見てみましょう。アプリ内で折り紙の写真を使って日記を作成できます。アプリはクラフト分析モードで開始するセッションを作成します。この命令は、モデルに画像を分析し、検出した内容を記録するように指示します。クラフトの種類、色、材料を識別し、ツール呼び出しを通じてそれらをジャーナルに保存します。
次に、アプリは Private Cloud Compute の推論機能を使用してブレインストーミング モードに切り替わります。学習したすべての情報が取得され、提案されます。

創作折り紙プロジェクトのリスト。かなりクールですよね？この機能を実装するには、まず LanguageModelSession を作成します。次に、それぞれに独自のモデル、手順、ツールを備えたセッションをさらに追加します。
しかし、モデルに自律的にモードを切り替えさせたい場合はどうすればよいでしょうか?物事が毛深くなり始めます。
このようなコンテキストの管理とエージェント システムの調整には、多くの定型文が必要となる場合があります。そのため、Foundation Models では新しい宣言型 API である動的プロファイルを導入しているため、単一の言語モデル セッション内で、命令型コントロールについて心配することなく、コンテキスト内で重要なことに集中できるようになります。単純な動的プロファイルを作成するには、構造体を宣言し、プロファイルを含む body プロパティを使用してそれを DynamicProfile プロトコルに準拠させます。
言語モデルセッションは、DynamicProfile を使用して初期化できるようになりました。
コンテキスト内に存在すべき指示やツールをその瞬間に指定できます。これは、命令とツールで構成されるデータ構造である DynamicProfile の最も単純な形式です。クラフト分析用とブレインストーミング用の 2 つの異なるモードを実装したいと考えています。私のアプリにはモード変数を保存する監視可能なオブジェクトがあるので、それをオンにできます。
ディで

[切り捨てられた]

## Original Extract

Explore what

What’s new in the Foundation Models framework - WWDC26 - Videos - Apple Developer
View in English
What’s new in the Foundation Models framework
Explore what's new in the Foundation Models framework. Learn how to access Private Cloud Compute, integrate third-party and open source models, and work with vision capabilities. Discover context management APIs, built-in semantic search, and powerful primitives for creating agentic experiences in your apps.
3:21 - Vision: image understanding
6:46 - Model abstraction layer
7:32 - Partner model integrations
9:40 - System tools: Vision and Spotlight
10:57 - Dynamic Profiles for agentic apps
13:46 - Composing models and configurations
16:02 - The fm command line tool
17:13 - Foundation Models Python SDK
17:55 - Open source and framework utilities
Expanding generation with tool calling
Analyzing images with multimodal prompting
Composing dynamic sessions with instructions and profiles
Adding server-side intelligence with Private Cloud Compute
Search this video…
Hello everyone, and welcome! My name is Erik! And my name is Zhen! Last year we introduced the Foundation Models framework with features like guided generation, snapshot streaming, and the powerful tool protocol. We were blown away by your excitement for the Foundation Models framework in year one, and we think you're going to love what we've lined up for you this year, even more! We're going to walk through everything that's new in the framework this year. And let me tell you, this release is packed! Our 2027 release is all about integrations into and beyond the OS, a wider variety of models, and new primitives for building agentic experiences. Let's kick it off with one of our most exciting updates. The Foundation Models framework, including many of the brand new APIs that we're announcing today, is going open source! And doing it in style! In addition to the core framework, we're also releasing a new package, Foundation Models framework utilities, that will be updated between OS releases to give you access to emerging and experimental building blocks. And during the course of this session, you'll be hearing about multiple other packages all joining the Foundation Models framework ecosystem. So, are you ready to jump in with us? I'm going to cover everything related to the new models, new modalities, and new tools we're adding to the framework. And I'm going to show you the brand new APIs we've made to help you get the most out of them! We've got a full schedule! First up, we've got a whole pile of model updates for you, including updates to the on-device model and access to server models.
We've also added new system tools that supercharge your session with help from Spotlight and the Vision framework. Afterwards, Zhen will rejoin us to tell you about a powerful new API called dynamic profiles, our new primitive for creating agentic experiences. Zhen will also dig into the brand new Evaluations framework and its tight integration with Foundation Models.
And to wrap up, we have some exciting news to share about Mac-specific productivity tools, so make sure to stick around! Let's jump right in with our new model!
This release comes with a new on-device model, rebuilt from the ground up, and better across the board. It's more intelligent; better at logic and tool calling. In iOS 26.4, we released new APIs for inspecting the model's context size and counting the tokens in instructions, prompts, and transcripts. You'll want to use these going forward to adapt your app to the hardware it's running on. We've also been hard at work on refining our guardrails. You may have noticed adjustments in iOS 26.4 to reduce the number of false positives, and we're continuing to make even more improvements in iOS 27.
In addition the on-device model is also gaining Vision capabilities, which unlocks entire new categories of applications.
The API is simple, a natural extension of the existing prompt builders.
Here we've created a session, and we want to ask about the photo of the origami on the right.
Simply insert an image attachment into your prompt, together with text.
Now, the model can answer questions about the image.
Image attachments can be created from a variety of types including; UIImage, NSImage, CGImage Core Image types, CoreVideo Pixel Buffers, and file URLs.
The model supports images in any size and aspect ratio, so you don't need to crop or pad to any particular shape.
Arbitrary image sizes are allowed, but bear in mind that larger images will consume more tokens and incur more latency. Thanks to all these upgrades, the on-device system language model is more capable than ever. But if you need even more horsepower, we're giving you access to a brand new PrivateCloudComputeLanguageModel. The Private Cloud Compute model is the very same one that powers many of the Apple Intelligence features you know and love. It is a much bigger model than the on-device models, and has a 32,000 token context window, and it comes with a powerful new capability, reasoning. Reasoning models are trained to spend time carefully thinking through their answers before providing a response, which results in significantly better outcomes. Using Private Cloud Compute couldn't be easier. Just create an instance of the model and use it to initialize your language model session. When prompting the session, I can now specify a reasoning level on the new contextOptions argument. reasoningLevel controls how much the model is allowed think before responding. Deep reasoning produces better responses in exchange for additional compute.
One of the best things about Private Cloud Compute is that you don't have to worry about account setup, you don't have to deal with authentication, and you don't have to store API keys, it's all completely seamless! And of course, Private Cloud Compute is above all else, private. No prompts are ever stored, and we make it possible for independent researchers to verify these claims. And to top it all off, Private Cloud Compute makes it possible for us to bring the Foundation Models framework to watchOS. Starting in watchOS 27, you can wear your most powerful intelligence features right on your wrist.
PCC is available with no cloud API costs to developers who have less than 2 million first time downloads. Your users will have access to PCC every day and if they are subscribed to iCloud+, their limit will be even higher! To learn more about the ins and outs of PrivateCloudComputeLanguageModel, including the entitlement you'll need to use it, make sure to tune in to our video about building with Private Cloud Compute.
In addition to an overhauled on-device model and the new Private Cloud Compute model, we're opening up our model abstraction layer to make it possible for nearly any language model to be used with the Foundation Model's framework. The abstraction layer is built around a new LanguageModel protocol that allows both local and servers models to back a LanguageModelSession. Existing models like SystemLanguageModel and PrivateCloudComputeLanguageModel already conform to this protocol. And, we're open sourcing two additional implementations: CoreAILanguageModel and MLXLanguageModel, for running a myriad of local models on the Apple Neural Engine in your Mac's GPU.
We're also excited to share that we've been hard at work behind the scenes to ensure you have access to a variety of frontier server models! Anthropic, and Google are both publishing Swift packages to provide you with access to their latest and greatest models! The model abstraction layer makes using third party models simple. I'll just import a language model package using Swift Package Manager, initialize the model that I want to use, and pass it when creating my session.
Everything downstream stays the same.
Bear in mind that if you use third party server models, you'll probably have to deal with both authentication and billing. Remember, never store private keys in your app binary. Always fetch access tokens with a secure mechanism like OAuth, and store them securely using KeyChain. As a developer, you'll typically be billed per-token when using 3rd party models, so we've made it easy to keep track of your usage. Sessions and responses now have a usage property that tells you precisely how many tokens were used. You can also check how many of the input tokens were read from cache, and how many of the response tokens were used for reasoning.
If you'd like to learn more about using LanguageModels, or how you can author your own LanguageModel package, check out "Bring an LLM provider to the Foundation Models framework". We have finally made it through all our model updates! Next up is system tools! In this release, we're introducing several built-in tools that supercharge your LanguageModelSessions with system provided functionality. FoundationModels now contains two native tools backed by the Vision framework's powerful capabilities.
The BarcodeReaderTool allows the model read information from barcodes, and the OCRTool allows the model to extract structured text from images. Both enhance a model's ability to reason about visual information in ways it can't natively. The "What's new in image understanding" video has more detail about how to leverage these tools, so queue that one up for more info. Similarly, we're also introducing a search tool powered by Spotlight for implementing fully local Retrieval-Augmented Generation. This has been one of your most most requested features. Retrieval-Augmented Generation, or RAG, is a technique that gives the model access to up-to-date personal or domain knowledge by leveraging a Spotlight index and specially processed queries. If this sounds like just what you've been waiting for, "LLM search using Core Spotlight" should be at the top of your watch list. Now that we've looked at all the new models in this release, and the new system tools we're adding to the SDK, I'm going to hand off to Zhen to tell you all about our new APIs for building agentic app experiences. Get excited! Take it away, Zhen! Thanks Erik!
I'm going to introduce you to dynamic profiles, our new primitive for building agentic experiences.
Let's begin by walking through the crafts app and looking at the kinds of experiences dynamic profiles make possible. Inside the app, I can create a journal entry with some origami photos. The app creates a session that starts in craft analysis mode. The instructions tell the model to analyze the images, and record what it finds. It identifies the craft type, colors, and materials, then saves them back to the journal, through a tool call.
Next, the app switches to brainstorming mode using Private Cloud Compute's reasoning capability, it takes everything it just learned and suggests a list of creative origami projects. Pretty cool, right? To implement this feature, I'd start by creating a LanguageModelSession. Then I'll add more sessions, each with its own models, instructions, and tools.
But what if I want the model to autonomously switch modes? Things start to get hairy.
Managing context and orchestrating an agentic system like this can involve a lot of boilerplates. That's why Foundation Models is introducing a new declarative API, dynamic profiles, so you can focus on what matters in the context, and worry less about imperative controls, all within a single language model session. To create a simple dynamic profile, I can declare a struct, and conform it to the DynamicProfile protocol with a body property that contains a Profile.
Language model sessions now can be initialized with a DynamicProfile.
I can specify the instructions and tools, that should be present in the context, at that very moment. This is the simplest form of a DynamicProfile, a data structure made up of instructions, and tools. I want to implement two different modes, one for craft analysis, and one for brainstorm. My app has an observable object that stores a mode variable, so I can switch on it.
In the di

[truncated]
