---
source: "https://www.macrumors.com/2026/06/09/apples-new-ai-contains-no-gemini/"
hn_url: "https://news.ycombinator.com/item?id=48462659"
title: "Apple's New AI Models Contain 'None' of Google's Gemini Assistant"
article_title: "Apple's New AI Models Contain 'None' of Google's Gemini Assistant - MacRumors"
author: "bound008"
captured_at: "2026-06-09T16:03:41Z"
capture_tool: "hn-digest"
hn_id: 48462659
score: 1
comments: 0
posted_at: "2026-06-09T15:47:35Z"
tags:
  - hacker-news
  - translated
---

# Apple's New AI Models Contain 'None' of Google's Gemini Assistant

- HN: [48462659](https://news.ycombinator.com/item?id=48462659)
- Source: [www.macrumors.com](https://www.macrumors.com/2026/06/09/apples-new-ai-contains-no-gemini/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T15:47:35Z

## Translation

タイトル: Appleの新しいAIモデルにはGoogleのGemini Assistantが「何も」含まれていない
記事タイトル：Appleの新しいAIモデルにはGoogleのGemini Assistantが「何も」含まれていない - MacRumors
説明: Apple 幹部は、同社の新しい Apple Foundation Models (AFM) のアーキテクチャを詳細に説明し、Google のテクノロジーがその開発にどのように組み込まれているかを正確に明らかにしました。

記事本文:
Appleの新しいAIモデルにはGoogleのGemini Assistantが「何も」含まれていない - MacRumors
コンテンツへスキップ ヒントはありますか?お知らせください
ａ．メールを送信してください
フォーラム フォーラム フォーラムメニューを表示 ログイン
Appleの新しいAIモデルにはGoogleのGemini Assistantが「まったく」含まれていない
Appleの幹部らは、同社の新しいApple Foundation Models（AFM）のアーキテクチャを詳細に説明し、Googleのテクノロジーがその開発にどのように組み込まれているかを正確に明らかにした。
Appleのソフトウェアエンジニアリング担当SVPであるCraig Federighi氏は月曜日、AI担当VPのAmar Subramanya氏、SiriリードのMike Rockwell氏、ソフトウェア担当VPのSebastien Marineau-Mes氏とともに基調講演後の技術講演（9to5Mac経由）をプレス向けに開催し、第3世代AFMファミリーがどのように構築され、それがApple Intelligenceにどのように機能するのかを説明した。
フェデリギ氏は、「われわれが使用しているGoogleアシスタントはまったくない」と述べ、AppleはGoogleの顧客に導入されているGeminiモデルを一切使用しておらず、Googleのクライアント側コードも、知識バックボーンとしてのGoogle検索インフラも使用していないことを説明した。
もちろん、アプリとして Gemini アプリはありません。実際、そのクライアント コードはどれも、iOS での実行方法には含まれていません。これらのモデルについては、Google が顧客にデプロイするモデルは一切使用せず、Google が顧客にモデルをデプロイするインフラストラクチャや手段も使用しません。そして、ナレッジベースに関して言えば、当然のことながら、システムの基盤として Google 検索などは使用しません。
Subramanya 氏は、2 つのオンデバイス モデルと 3 つのサーバーサイド モデルにわたる新しい AFM ファミリの概要を説明しました。オンデバイス層は、次世代の高密度アーキテクチャ モデルである AFM Core と、スパース アーキテクチャを使用し、ネイティブにマルチモーダルである AFM Core Advanced で構成されます。
Subramanya 氏は、AFM Core Advanced は「これまでに実行したオンデバイス モデルとは異なる」と述べました。

クラウドリクエストを必要とせずに、招待や表現力豊かな音声などの新機能を利用できます。サーバー側では、AFM Cloud がレイテンシが最適化されたプライベート クラウド コンピューティング リクエストを処理し、AFM Cloud Image が空間リフレーミングなどの画像生成および編集機能を強化します。
Google とのコラボレーションに関する重要な詳細は、これら 4 つのモデルがどのようにトレーニングされたかについての Subramanya の説明の中にありました。 「これらはすべて Apple Silicon 用にカスタム構築されており、強化学習による独自のデータを使用してトレーニングされ、Gemini フロンティア モデルからの出力を使用して洗練されています」と同氏は述べ、Google の貢献は蒸留ベースであり、Gemini の全面的な採用ではないことを明らかにしました。
5 番目の最も有能なモデルである AFM Cloud Pro は、エージェント ツールの使用と複雑な推論タスク向けに設計されており、Subramanya 氏によると「Gemini フロンティア モデルに似た」品質です。このモデルは、Apple の標準的なプライベート クラウド コンピューティング セットアップからの脱却を示しています。
これを実行するために、Apple は Google と Nvidia の両方と協力して、プライベート クラウド インフラストラクチャを Google のクラウドでホストされている Nvidia GPU に拡張しました。マリノーメス氏は、アップルはエヌビディアの最新チップの使用を望んでいたが、アップルのサーバーの内容を読み取れないようにチップを構成する必要があったと述べた。 「曖昧な機密コンピューティング」と呼ばれる最近の Nvidia テクノロジーが解決策を提供しました。
私たちは Nvidia の最新テクノロジーを活用したいと考え、プライベート クラウド コンピューティングをサードパーティ クラウドに拡張することに着手しました。
フェデリギ氏は、より広範なシステム アーキテクチャが System Orchestrator を中心に組織されていると説明しました。これは、彼が「システム全体のプライバシー アーキテクチャの鍵」と呼ぶソフトウェアです。オーケストレーターは、リクエストの複雑さと必要な個人的なコンテキストに基づいて、特定のクエリをデバイス上またはクラウドの適切なモデルにルーティングします。
それは次のようになります

アプリ内アクション用のアプリ ツールボックス、個人コンテンツ用の Spotlight セマンティック インデックス、およびリアルタイム認識用の画面上のコンテキストです。時事問題に関するクエリの場合、Apple 独自の World Knowledge Service を通じて応答が見つかります。Federighi 氏によると、このサービスは同社が数年間にわたって構築してきたものです。
Apple はまた、Google のクラウド内の拡張 Nvidia GPU 容量を含むすべてのプライベート クラウド コンピューティング インフラストラクチャは、サードパーティの研究者によって独立して検証され、ユーザー データが決して保存またはアクセスされないことを確認できると主張しています。
毎週の MacRumors のトップ記事を受信箱に入手してください。
CarPlay 用 Google マップ Gemini AI を入手
Google、AppleのiOS 27発表の1か月前に「Gemini Intelligence」を搭載したAndroid 17をプレビュー
Google により Android と iPhone 間でのファイル共有が簡単に
私の 16 Pro Max で新しい Siri をテストしてみたところ、正直言って非常に印象的でした。明らかに解決すべき問題がいくつかあり、Siri がサードパーティ アプリから情報を取得する際にも問題があるようですが、それでも、以前のバージョンの Siri とは昼も夜も違います。
iOS 26 のすべての主要な新機能に焦点を当てた包括的なガイドと、新機能の使用手順を説明するガイドです。
ベータ版のエンドツーエンド暗号化 RCS メッセージング、マップ内のおすすめの場所、新しいプライド ルミナンスの壁紙など。
アップグレードする場合は、新機能とあまり知られていない変更点を確認してください。
Apple の 4 つの新しい iPhone は、最新モデル間の違いがこれまで以上に多いのが特徴です。ではどれを購入すべきでしょうか?
iOS 27やmacOS 27などを発表するAppleの年次開発者カンファレンス。
Siriの改善、将来の折りたたみ式iPhone向けの最適化など。 WWDC で 6 月に導入され、秋にリリースされます。
よりスマートな Siri、将来の MacBook Pro に向けたタッチの最適化、安定性を重視したアップデートなど

。 WWDC で 6 月に導入され、秋にリリースされます。
M5 および M5 Pro チップのオプションにより仕様が向上します。
MacRumors は、最新のテクノロジーや製品に興味のある消費者と専門家の両方の幅広い視聴者を惹きつけています。また、iPhone、iPad、Mac、その他の Apple プラットフォームの購入決定や技術的側面に焦点を当てた活発なコミュニティも誇っています。

## Original Extract

Apple executives have detailed the architecture of the company's new Apple Foundation Models (AFM) and clarified exactly how Google's technology factored into their development.

Apple's New AI Models Contain 'None' of Google's Gemini Assistant - MacRumors
Skip to Content Got a tip for us? Let us know
a. Send us an email
Forums Forums Show Forums menu Login
Apple's New AI Models Contain 'None' of Google's Gemini Assistant
Apple executives have detailed the architecture of the company's new Apple Foundation Models (AFM) and clarified exactly how Google's technology factored into their development.
Craig Federighi, Apple's SVP of Software Engineering, held a post-keynote tech talk (via 9to5Mac ) with press on Monday alongside AI VP Amar Subramanya, Siri lead Mike Rockwell, and software VP Sebastien Marineau-Mes to walk through how the third-generation AFM family was built and how it powers Apple Intelligence .
"The amount of the Google Assistant we use is none," Federighi said, explaining that Apple uses none of the Gemini models deployed to Google's customers, none of Google's client-side code, and no Google Search infrastructure as the knowledge backbone.
Of course, we don't have the Gemini app as our app. In fact, none of that client code is part of how we run on iOS. For these models, we use none of the models that Google deploys to their customers, nor do we use the infrastructure and means by which they deploy models to their customers. And then, when it comes to the knowledge base, we of course don't use Google Search or anything like that as the foundation of our system.
Subramanya outlined the new AFM family, which spans two on-device models and three server-side models. The on-device tier consists of AFM Core, a next-generation dense architecture model, and AFM Core Advanced, which uses a sparse architecture and is natively multimodal.
Subramanya said AFM Core Advanced is "unlike any on-device model we've run before," enabling new features including invitation and expressive voices without any cloud requests. On the server side, AFM Cloud handles latency-optimized Private Cloud Compute requests, while AFM Cloud Image powers image generation and editing features including spatial reframing.
The key detail on the Google collaboration came in Subramanya's description of how these four models were trained. "All of these are custom built for Apple Silicon , trained using proprietary data with reinforcement learning and refined using outputs from Gemini frontier models," he said, making clear that Google's contribution was distillation-based, not a wholesale adoption of Gemini.
The fifth and most capable model, AFM Cloud Pro, is designed for agentic tool use and complex reasoning tasks, with quality that Subramanya said is "similar to Gemini frontier models." This model marks a departure from Apple's standard Private Cloud Compute setup.
To run it, Apple worked with both Google and Nvidia to extend its private cloud infrastructure to Nvidia GPUs hosted in Google's cloud. Marineau-Mes said Apple wanted to use Nvidia's latest chips but required them to be configured so they couldn't read the contents of Apple's servers. A recent Nvidia technology called "ambiguous confidential compute" provided the solution.
We wanted to avail ourselves of the latest technology from Nvidia, and so we set out to extend private cloud compute to third-party cloud.
Federighi described the broader system architecture as being organized around a System Orchestrator, a piece of software he called "key to the privacy architecture of our entire system." The orchestrator routes any given query to the appropriate model, on-device or cloud, based on the complexity of the request and the personal context required.
It draws on an App Toolbox for in-app actions, a Spotlight Semantic Index for personal content, and on-screen context for real-time awareness. For queries involving current events, responses are found through Apple's own World Knowledge Service, which Federighi said the company has been building for several years.
Apple also maintains that all Private Cloud Compute infrastructure, including the extended Nvidia GPU capacity in Google's cloud, can be independently verified by third-party researchers to confirm that user data is never stored or accessed.
Get weekly top MacRumors stories in your inbox.
Google Maps for CarPlay Getting Gemini AI
Google Previews Android 17 With 'Gemini Intelligence' a Month Before Apple's iOS 27 Reveal
Google Makes It Easier to Share Files Between Android and iPhone
Testing the new Siri on my 16 Pro Max, it's honestly quite impressive. There are obviously some quirks that need to be solved, and it also seems Siri has issues drawing information from 3rd party apps, but still, night and day difference from the previous version of Siri.
Our comprehensive guide highlighting every major new addition in iOS 26, plus how-tos that walk you through using the new features.
End-to-end encrypted RCS messaging in beta, Suggested Places in Maps, new Pride Luminance wallpaper, and more.
New features and lesser-known changes to check out if you're upgrading.
Apple's four new iPhones feature more differences between the latest models than ever before, so which one should you buy?
Apple's annual developer conference where it will unveil iOS 27, macOS 27, and more.
Siri improvements, optimizations for future foldable iPhone, and more. June introduction at WWDC with fall release.
Smarter Siri, touch optimizations for a future MacBook Pro, stability-focused updates, and more. June introduction at WWDC with fall release.
Spec bumps with M5 and M5 Pro chip options.
MacRumors attracts a broad audience of both consumers and professionals interested in the latest technologies and products. We also boast an active community focused on purchasing decisions and technical aspects of the iPhone, iPad, Mac, and other Apple platforms.
