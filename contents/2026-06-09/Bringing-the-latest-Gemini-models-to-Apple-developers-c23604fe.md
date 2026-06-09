---
source: "https://blog.google/innovation-and-ai/technology/developers-tools/bringing-gemini-models-to-apple-developers/"
hn_url: "https://news.ycombinator.com/item?id=48455479"
title: "Bringing the latest Gemini models to Apple developers"
article_title: "Gemini models for Apple developers"
author: "xnx"
captured_at: "2026-06-09T04:29:38Z"
capture_tool: "hn-digest"
hn_id: 48455479
score: 2
comments: 0
posted_at: "2026-06-09T02:26:04Z"
tags:
  - hacker-news
  - translated
---

# Bringing the latest Gemini models to Apple developers

- HN: [48455479](https://news.ycombinator.com/item?id=48455479)
- Source: [blog.google](https://blog.google/innovation-and-ai/technology/developers-tools/bringing-gemini-models-to-apple-developers/)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T02:26:04Z

## Translation

タイトル: 最新の Gemini モデルを Apple 開発者に提供
記事のタイトル: Apple 開発者向けの Gemini モデル
説明: Apple 開発者は、Foundation Models フレームワークを使用してクラウドでホストされている Gemini モデルを安全に呼び出し、Xcode で Gemini にアクセスできるようになりました。

記事本文:
メインコンテンツにスキップ
キーワード
最新の Gemini モデルを Apple 開発者に提供
シェアする
x.com
フェイスブック
リンクトイン
メール
リンクをコピー
ホーム
イノベーションとAI
イノベーションとAI
モデルと研究
Googleディープマインド
インフラストラクチャとクラウド
グローバルネットワーク
モデルと研究
Googleディープマインド
インフラストラクチャとクラウド
グローバルネットワーク
製品とプラットフォーム
製品とプラットフォーム
製品
検索
会社ニュース
会社ニュース
アウトリーチと取り組み
機会の創出
Google の内部
世界中で
アウトリーチと取り組み
機会の創出
Google の内部
世界中で
イノベーションとAI
イノベーションとAI
モデルと研究
Googleディープマインド
インフラストラクチャとクラウド
グローバルネットワーク
製品とプラットフォーム
製品とプラットフォーム
製品
検索
会社ニュース
会社ニュース
アウトリーチと取り組み
機会の創出
Google の内部
世界中で
最新の Gemini モデルを Apple 開発者に提供
Apple 開発者は、Gemini モデルにシームレスにアクセスして、動的なエクスペリエンスとよりスマートなアプリをより迅速に構築できるようになります。さらに、Xcode の Gemini は、ウィンドウを切り替えることなく、複数ステップのコーディング タスクを高速化するのに役立ちます。
ニコラス・マクナマラ
シニアプロダクトマネージャー
テビ・スンダラリンガム
グループプロダクトマネージャー
AI が生成した概要を読む
箇条書き
「最新の Gemini モデルを Apple 開発者に提供」すると、よりスマートなアプリをより迅速に構築できるようになります。
Apple のネイティブ Foundation Models フレームワークを通じて Gemini モデルを直接呼び出すことができるようになりました。
Firebase AI Logic を使用すると、バックエンド サーバーを管理せずに Gemini をアプリに追加できます。
Gemini は Xcode に統合され、コードの作成と修正を簡単に行えるようになりました。
特定の開発ニーズに合わせて、セルフサービス キーまたはエンタープライズ オプションのいずれかを選択してください。
お使いのブラウザは audio 要素をサポートしていません。
今週、Apple の Worldwide Developers Conference (WWDC) が開幕しました。

Apple 開発者が Foundation Models フレームワークを使用してクラウドでホストされている Gemini モデルを安全に呼び出し、Xcode で Gemini にアクセスできるようになったということを共有します。この発表により、Apple 開発者は Gemini モデルにシームレスにアクセスできるようになり、エンド ユーザーに動的なエクスペリエンスを提供し、開発速度を向上させることができます。
Foundation Models フレームワークから Gemini モデルを直接呼び出す
Apple は WWDC で、Foundation Models フレームワークをサードパーティのクラウド モデル プロバイダーに公開すると発表しました。 iOS 27、macOS 27、iPadOS 27、visionOS 27、watchOS 27 以降、モデルプロバイダーは新しいパブリック LanguageModel プロトコルを実装して、モデル推論用の共通インターフェイスを提供できます。 Gemini モデルを Firebase Apple SDK を通じて Foundation Models フレームワークで利用できるようにしました。
これにより、完全にネイティブな開発エクスペリエンスが提供されます。クラウドでホストされている Gemini モデルは、同じ API を使用して Foundation Models フレームワークに直接接続できます。つまり、オンデバイスの Apple モデルとクラウドでホストされる Gemini モデルは共有 API サーフェスの背後に位置するため、ユースケースに合わせてローカル推論とクラウド推論を簡単に切り替えることができます。この柔軟性は、コストを最小限に抑え、遅延を削減しながら、最適なエージェント アプリ エクスペリエンスを実現するために非常に重要です。
この統合では、Google Cloud のクライアント アプリ開発プラットフォームである Firebase を活用します。これは、別のバックエンド サーバーを構築して維持する必要なく、最新の Gemini モデルを iOS、macOS、iPadOS、visionOS アプリに直接統合できる実稼働グレードのサービスである Firebase AI Logic に基づいています。さらに、 Firebase App Check を使用すると、Gemini モデルへのアクセスに使用されるサービス API などのサービス API が悪用から保護されます。
すでに Apple の Foundation Models フレームワークを使用している場合、Gemini モデルへの切り替えはコードの小さな変更です。

モデルインスタンス。
この統合は明日からプレビュー リリースとして利用可能になります。それまでの間、Firebase AI Logic について詳しく学んでください。
Xcode でエージェント ワークフローをネイティブに使用する
私たちは Apple と協力して Gemini を Xcode に統合し、開発中にツールやウィンドウを切り替えることなく複雑な複数ステップのタスクを実行できるようにしました。
開始するには、Xcode のインテリジェンス設定パネルからオンボードします。 Gemini を構成すると、エージェント エクスペリエンスが提供され、コードのレビュー、バグの修正、新機能の構築をより迅速に行うことができます。
個人開発者と企業開発者の両方のワークフローをサポートするには、ユースケースに基づいて認証を構成できます。
個人の開発者: Google AI Studio からセルフサービスの Gemini API キーを直接取得します。これには、開始するための無料枠と、より高度なモデルや大容量の有料枠の両方が含まれます。
エンタープライズ開発者: Gemini Enterprise Agent Platform を使用して API キーを取得すると、開発チームは組織で保証された法人専用割り当てとデータ プライバシー パラメーターを活用できるようになります。
Apple プラットフォームで Gemini モデルを使ってみる
Apple の Foundation Models フレームワークと Xcode 間のシームレスな統合により、Gemini モデルを使用して次世代のエージェント AI エクスペリエンスを簡単に構築、保護、拡張できます。あなたが何を構築するのかを見るのが待ちきれません。
受信トレイで Google からの記事をさらに入手しましょう。
受信トレイで Google からの記事をさらに入手しましょう。
あなたの情報は以下に従って使用されます
Google のプライバシー ポリシー。
受信箱をチェックして購読を確認してください。
あなたはすでにニュースレターを購読しています。
で購読することもできます
別のメールアドレス
。
連絡を取り合いましょう。 Google からの最新ニュースを受信トレイで受け取ります。
グローバル（英語）
アフリカ (英語)
オーストラリア (英語)
ブラジル (ポルトガル語)

s)
カナダ (英語)
カナダ (フランス語)
チェコ語
ドイチュラント (ドイツ)
スペイン (スペイン語)
フランス
ギリシャ
インド (英語)
インドネシア (インドネシア語)
アイルランド (英語)
イタリア語
日本 (日本語)
한국어 (한국어)
ラテンアメリカ (スペイン語)
中東および北アフリカ (アラビア語)
メナ (英語)
オランダ (オランダ)
ニュージーランド (英語)
ポーランド (ポルスキ)
ポルトガル
スウェーデン
タイ (ไทย)
トゥルキエ (トゥルクチェ)
台灣 (中文)

## Original Extract

Apple developers can now securely call cloud-hosted Gemini models using the Foundation Models framework, and access Gemini in Xcode.

Skip to main content
The Keyword
Bringing the latest Gemini models to Apple developers
Share
x.com
Facebook
LinkedIn
Mail
Copy link
Home
Innovation & AI
Innovation & AI
Models & Research
Google DeepMind
Infrastructure & cloud
Global network
Models & Research
Google DeepMind
Infrastructure & cloud
Global network
Products & platforms
Products & platforms
Products
Search
Company news
Company news
Outreach & initiatives
Creating opportunity
Inside Google
Around the globe
Outreach & initiatives
Creating opportunity
Inside Google
Around the globe
Innovation & AI
Innovation & AI
Models & Research
Google DeepMind
Infrastructure & cloud
Global network
Products & platforms
Products & platforms
Products
Search
Company news
Company news
Outreach & initiatives
Creating opportunity
Inside Google
Around the globe
Bringing the latest Gemini models to Apple developers
Apple developers are getting seamless access to Gemini models to build dynamic experiences and smarter apps, faster. Additionally, Gemini in Xcode helps you accelerate multi-step coding tasks without switching windows.
Nicholas McNamara
Senior Product Manager
Thevi Sundaralingam
Group Product Manager
Read AI-generated summary
Bullet points
"Bringing the latest Gemini models to Apple developers" helps you build smarter apps faster.
You can now call Gemini models directly through Apple's native Foundation Models framework.
Firebase AI Logic lets you add Gemini to your apps without managing backend servers.
Gemini is now integrated into Xcode to help you write and fix code easily.
Choose between self-serve keys or enterprise options to fit your specific development needs.
Your browser does not support the audio element.
Apple’s Worldwide Developers Conference (WWDC) kicked off this week, and we’re excited to share that Apple developers can now securely call cloud-hosted Gemini models using the Foundation Models framework, and access Gemini in Xcode. This announcement gives Apple developers seamless access to Gemini models to deliver dynamic experiences for their end users and increase their own development velocity.
Call Gemini models directly from the Foundation Models framework
At WWDC, Apple announced that it's opening its Foundation Models framework to third-party cloud model providers. Starting with iOS 27, macOS 27, iPadOS 27, visionOS 27 and watchOS 27, model providers can implement the new public LanguageModel protocol to provide a common interface for model inference. We've made Gemini models available to the Foundation Models framework through the Firebase Apple SDK.
This provides a fully native development experience — cloud-hosted Gemini models can plug directly into the Foundation Models framework using the same API. That means the on-device Apple model and cloud-hosted Gemini models sit behind a shared API surface, so you can easily swap between local and cloud inference to fit your use case. This flexibility is crucial for enabling optimal agentic app experiences while minimizing costs and reducing latency.
This integration leverages Firebase , Google Cloud's client apps development platform. It's based on Firebase AI Logic , a production-grade service that lets you integrate the latest Gemini models directly into your iOS, macOS, iPadOS and visionOS apps without needing to build and maintain a separate backend server. Plus, by using Firebase App Check , your service APIs, like those used to access Gemini models, are protected from abuse.
If you're already using Apple's Foundation Models framework, switching to Gemini models is a small code change: swap the model instance.
This integration will be available as a preview release starting tomorrow. In the meantime, learn more about Firebase AI Logic .
Use agentic workflows natively in Xcode
We worked with Apple to integrate Gemini into Xcode so you can perform complex, multi-step tasks during development without switching tools or windows.
To get started, onboard through the Intelligence settings panel in Xcode. Once configured, Gemini provides an agentic experience, helping you review code, fix bugs and build new features faster.
To support both individual and enterprise developer workflows, you can configure authentication based on your use case:
Individual developers: Obtain a self-serve Gemini API key directly from Google AI Studio, which features both a free tier for getting started and a paid tier for more advanced models and higher volumes.
Enterprise developers: Use Gemini Enterprise Agent Platform to obtain an API key, which enables development teams to leverage their organization’s guaranteed dedicated corporate quotas and data privacy parameters.
Get started with Gemini models on Apple platforms
With seamless integration across Apple's Foundation Models framework and Xcode, you can effortlessly build, secure and scale next-generation, agentic AI experiences using Gemini models. We can’t wait to see what you build.
Get more stories from Google in your inbox.
Get more stories from Google in your inbox.
Your information will be used in accordance with
Google's privacy policy.
Check your inbox to confirm your subscription.
You are already subscribed to our newsletter.
You can also subscribe with a
different email address
.
Let’s stay in touch. Get the latest news from Google in your inbox.
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)
