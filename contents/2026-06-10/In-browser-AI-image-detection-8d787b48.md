---
source: "https://www.quantable.com/ai/detecting-ai-generated-images/"
hn_url: "https://news.ycombinator.com/item?id=48476068"
title: "In-browser AI image detection"
article_title: "Detecting AI Generated Images - Quantable Analytics"
author: "jhpacker"
captured_at: "2026-06-10T16:23:57Z"
capture_tool: "hn-digest"
hn_id: 48476068
score: 3
comments: 0
posted_at: "2026-06-10T13:34:15Z"
tags:
  - hacker-news
  - translated
---

# In-browser AI image detection

- HN: [48476068](https://news.ycombinator.com/item?id=48476068)
- Source: [www.quantable.com](https://www.quantable.com/ai/detecting-ai-generated-images/)
- Score: 3
- Comments: 0
- Posted: 2026-06-10T13:34:15Z

## Translation

タイトル: ブラウザ内 AI 画像検出
記事のタイトル: AI で生成された画像の検出 - Quantable Analytics
説明: AI 画像生成検出をブラウザ内でローカルに行う「SlopGuard」という新しいブラウザ拡張機能をリリースしました。こちらの Chrome ウェブストアから入手できます。注: [Chrome に追加] ボタンをクリックする前に、これは非常に大きな拡張機能 (400M) であるため、
[切り捨てられた]

記事本文:
AI 生成画像の検出 - Quantable Analytics
量的分析
AI 画像生成検出をブラウザ内でローカルに行う「SlopGuard」という新しいブラウザ拡張機能をリリースしました。こちらの Chrome ウェブストアから入手できます。
注: [Chrome に追加] ボタンをクリックする前に、これは非常に大きな拡張機能 (400M) であるため、ダウンロードとインストールに数分かかることを知っておく必要があります。これはデスクトップ上でのみ実行できます。
AI によって生成された画像の品質が現実とほとんど区別できなくなるにつれて、現実と非現実を区別することがますます重要になります。
新しい拡張機能では、次のようなものが必要でした。
かなり正確です。見逃しや誤検知は常にありますが、多くの場合間違っている AI 検出モデルは役に立ちません。私たちはこの拡張機能に数か月間取り組んでおり、さまざまなモデルや方法を試してきましたが、そのほとんどが信頼できませんでした。私たちは、「役立つ」という重要な基準に達するものの組み合わせを見つけたと考えています。
高速、無料、そしてプライベートです。ここでの私たちの目標は、ユーザーが何が真実で何が真実でないのかを広く理解できるようにすることです。私たちが使用するモデルの精度は、優れた有料 API を介して実行できるものほど良くありませんが、ローカル モデルを使用することで、誰でも使用できるようになります。
複数の識別手段に依存しており、モデルに依存しません。画像生成モデルは急速に変化するため、検出モデルも同様に交換できる必要があります。数年前のオープンウェイト モデルは、現行世代のモデルで作成された画像を検出できません。
任意の Web ページから 6 本指の手のアイコンをクリックすると、非常に小さいものを除く、現在ページ上に表示されているすべての画像がスキャンされます。これ

画像ごとに平均して約 1 ～ 2 秒かかります。画像に「AI」ラベルが付けられるか、AI 信号が検出されない場合は緑色の境界線が追加されます。画像の周りの灰色の境界線は、検出が失敗したことを意味します。
したがって、上記の例では、最初の画像は間違いなく AI として検出され、2 番目の画像はそうではないと検出されました。個々の画像を右クリックして、その画像だけを確認することもできます。無限スクロールの場合 (上の Reddit スクリーンショットのような)、新しい画像が読み込まれた後、手をもう一度クリックして再度スキャンする必要があります。
画像に「AI」とマークが付いている場合は、メタデータが見つかったことを意味します。決定的なメタデータがない場合、最も信頼できると宣言されるのは「おそらく AI」です。オプションでデバッグ モードをオンにすると、さまざまな検出レイヤーの出力を確認できます。
デバッグをオンにした SlopGuard。
SlopGuard は 3 つの方法で AI 画像を検出します。
メタデータは、物事が理想的に機能する方法です。 C2PA 署名付きマニフェスト クレーム (コンテンツ検証のゴールド スタンダード) であれ、EXIF/IPTC などの従来の画像メタデータ標準であれ、メタデータを持つことが、AI であることを 100% 確信できる唯一の方法です。また、速くて簡単、そしてオープンです。
残念ながら、メタデータはオンラインの画像から削除される傾向があります。これは、画像がどのように生成されたかを隠すために意図的に行われる場合もありますが、多くの場合、Web サイトの仕組みに過ぎません。たとえば、LinkedIn は画像上での C2PA コンテンツ認証情報の表示をサポートしていますが、提供する実際の画像ファイルからは認証情報を取り除いています。したがって、その画像をダウンロードすると、そのメタデータはすでに失われています。パフォーマンスとプライバシー上の理由から、Web 上の画像からメタデータを削除することが長年にわたってベスト プラクティスとなってきました。ただし、今はそのデータの一部を戻す必要があります。
最大の AI 画像生成ツールのほとんど (ChatGPT、CoPilot、イレブンラボ、

Gemini) は、ある種のメタデータを設定しています。グロクはまだ何もしません。
コマンドライン exiftool と c2patool を使用したテスト。 CoPilot には C2PA 署名がありますが、検証されていません。
レベル 2: SynthID ウォーターマーク
Google DeepMind は、Gemini が作成したすべての画像に配置する SynthID と呼ばれる優れた不可視透かしシステムを開発しました。
5月にGoogleは、他のAI画像生成ツール（OpenAIやイレブンラボなど）もコンテンツに透かしを入れるためにSynthIDを使用すると発表した。 ChatGPT は、生成するイメージでこれをすでに起動しています。ただし、私のテストでは、現在 ChatGPT と Google の SynthID 実装には互換性がないように見えます。言い換えれば、OpenAI の画像検証ツールでは Gemini が作成した SynthID を検出できず、同様に Gemini の「@Verify AI」も ChatGPT によって作成された SynthID を検出できませんでした。現在、私たちの拡張機能は Google SynthID の Gemini バージョンのみを検出し、OpenAI のバリアントは検出しません。
ユーザーの観点から見たこの方法の現時点での欠点は、個々の画像が AI によって生成されたものであるかどうかを確認するために、Gemini にログインして画像をアップロードする必要があることです。
SlopGuard は、ローカル モデル アプローチを使用して、一度に多くの画像をチェックできます。画像を 1 つずつ Gemini にアップロードし、「@Verify AI」と言って、完全な回答が得られるまで 20 秒待つのではなく。
Google は SynthID をオープンにしていませんが、プラグインは OpenSynthID Detect というリバース エンジニアリングされた検出モデルを使用しています。ローカルにダウンロードされた OpenSynthID モデルは、画像あたり 1 秒未満でこの質問に答えることができます。 Google は Chrome に SynthID を含めると述べているため、おそらくこのメソッドはすぐに廃止されるでしょう。
レベル 3: 視覚的分類子モデル
これは、「AI 検出ツール」について話すときに、ほとんどの人が思い浮かべる部分です。

透かしやメタデータに関係なく、画像自体のピクセルに基づきます。
私たちのプラグインは Victor Livernoche による OpenFake モデルを使用しています。一連のオープンウェイト モデルをテストした結果、他のオプションはすべて不正確すぎるか、場合によっては Chrome 拡張機能には大きすぎることがわかりました。 OpenFake モデル自体は Microsoft の SwinV2 Transformer に基づいています。
このモデルは OpenSynthID モデル (350M 対 80M) よりもはるかに大きく、はるかに困難なタスクを抱えていますが、このテストの推論部分は、かなり強力な CPU を使用しても 1 秒未満かかります。この拡張機能の将来のバージョンでは、利用可能なローカル GPU をより有効に活用する方法が検討される可能性があります。テストでは推論をより高速に実行できましたが、拡張機能の他の部分で問題やバグが発生しました。推論が高速化しても、GPU バージョンは他のすべての可動部分のせいで、全体的にはそれほど高速ではありませんでした。
バージョン 0.4 であるこの拡張機能は、間違いなくまだ開発中です。
https://github.com/jhpacker/slopguard ではオープンソース (GPL) なので、問題がある場合はそこに送信するか、jason [at] quantable [dot] com にメールしてください。
この投稿を共有:
ツイッターで
フェイスブックで
LinkedIn で
返信を残す 返信をキャンセルするにはここをクリックしてください。
このサイトはスパムを低減するために Akismet を使っています。コメントデータがどのように処理されるかをご覧ください。
Google Analytics のサイトごとのオプトアウト ブラウザ拡張機能
人気があることがすべてではない
Google Analytics の代替手段、総合ガイド (第 2 版)
Quantable Analytics © 2026.全著作権所有。
WordPress を利用しています。によって設計されました

## Original Extract

We’ve just launched a new browser extension called “SlopGuard” which does AI image generation detection locally inside your browser! You can get it from the Chrome Web Store here. Note: before you click the “Add to Chrome” button, you should know that this is a very large extension (400M), so it wil
[truncated]

Detecting AI Generated Images - Quantable Analytics
Quantable Analytics
We’ve just launched a new browser extension called “SlopGuard” which does AI image generation detection locally inside your browser! You can get it from the Chrome Web Store here .
Note: before you click the “Add to Chrome” button, you should know that this is a very large extension (400M), so it will take a few minutes to download and install. This can only be run on desktop.
As the quality of AI generated images becomes nearly indistinguishable from reality, separating the real from unreal becomes increasingly important.
With our new extension, we wanted something that:
Is reasonably accurate. There’s always going to be misses and false positives, but an AI detection model that is wrong much of the time simply isn’t helpful. We’ve been working on this extension for a couple months now and have tried many different models and methods, the majority of which just not reliable. We think we’ve found a combination of things that reaches the important bar of “useful”.
Is fast, free, and private . Our goal here is to broadly help users see what’s real and what’s not. The accuracy of the models we use will not be as good as what can be done via a good paid API, but using local models makes it something anyone can use.
Relies on multiple means of identification, and is model agnostic . Since image generation models change so fast, the detection models need to be able to be swapped out as well. Open weight models that are a couple years old are simply not able to detect images created by current generation models.
From any webpage, click the six-fingered hand icon, and it will scan all the images it sees currently on the page, excluding those that are very small. This should take about 1-2 seconds per image on average, and it will either place an “AI” label on the image or add a green border if no AI signals are detected. A grey border around the image means detection failed.
So in the example above it detected the first image as definitely AI, and the second as not. You can also right-click on any individual image to check it that image alone. In cases where’s the infinite scroll (like the Reddit screenshot above), you’ll have to click the hand again to get it scan again once there are new images loaded.
If the image is marked “AI”, that means we found metadata. Without definitive metadata the most confidence it will declare is “Probably AI”. You can see the outputs of the different layers of detection by turning on the Debug Mode in options.
SlopGuard with debug turned on.
SlopGuard detects AI images in 3 ways:
Metadata is the way things should ideally work! Whether it’s an C2PA signed manifest claim (the gold standard for validating content), or a traditional image metadata standard like EXIF/IPTC, having metadata is really the only way to be 100% sure something is AI. It’s also fast, easy, and open.
Unfortunately metadata tends to get stripped from images online. Sometimes this is purposeful to hide how an image was generated, but more often it’s just how a website works. For example, LinkedIn supports displaying C2PA content credentials on an image, but they strip it from the actual image file they serve. So when you download that image, it’s already lost that metadata. For performance and privacy reasons, stripping metadata from images on the web has been a best practice for many years. Except now we need some of that data back!
Most of the biggest AI image generation tools (ChatGPT, CoPilot, ElevenLabs, Gemini) are setting metadata of some sort. Grok still does nothing.
Testing with command-line exiftool and c2patool. CoPilot does have a C2PA sig, but it didn’t validate.
Level 2: SynthID Watermarks
Google DeepMind has developed an excellent invisible watermarking system called SynthID that they place on all Gemini-created images.
In May, Google announced other AI image generation tools (including OpenAI and ElevenLabs) are going to be using SynthID to watermark their content as well. ChatGPT has already launched this in images it generates! Although in my testing it appears that currently ChatGPT and Google’s implementations of SynthID are incompatible. In other words, I was unable to detect a Gemini-created SynthID in OpenAI’s image validator and similarly Gemini’s “@Verify AI” couldn’t detect one created by ChatGPT. Currently our extension only detects the Gemini version of Google SynthID, not OpenAI’s variant.
The downside of this method currently from a user perspective is that you have to login to Gemini and upload your image to check if an individual image is AI-generated.
SlopGuard can check lots of images at once with its local model approach. Instead of uploading an image one at a time to Gemini, saying “@Verify AI”, and then waiting 20 seconds for a full answer.
While Google has not made SynthID open, our plugin uses a reverse-engineered detection model named OpenSynthID Detect . The locally downloaded OpenSynthID model can answer this question in < 1s second per image. Google has said they will be including SynthID in Chrome, so perhaps this method will soon be deprecated.
Level 3: Visual Classifier Model
This is the part that most of us would think of when talking about an “AI detection tool”, something that can do its detection just based upon the pixels in the image itself, regardless of watermarking or metadata.
Our plugin uses the OpenFake model by Victor Livernoche. After testing a series of open weight models, we found all the other options either too inaccurate or in some cases simply too big for a Chrome extension. The OpenFake model is itself based upon Microsoft’s SwinV2 Transformer .
This model is much bigger than the OpenSynthID model (350M vs. 80M) and has the much harder task, but the inference part of this test still takes under 1s with a reasonably powerful CPU. Future versions of this extension may look at how to better utilize available local GPUs, which in testing were able to run inference much faster, but caused difficulties and bugs with other parts of the extension. Even with faster inference, the GPU version wasn’t that much faster overall because of all the other moving parts.
As version 0.4, this extension is still definitely a work-in-progress!
It’s also open source (GPL) at https://github.com/jhpacker/slopguard , so if you have issues please submit them there or just email me, jason [at] quantable [dot] com .
Share this post:
on Twitter
on Facebook
on LinkedIn
Leave a Reply Click here to cancel reply.
This site uses Akismet to reduce spam. Learn how your comment data is processed.
Google Analytics Opt-Out Per Site Browser Extension
Being Popular Ain’t Everything
Google Analytics Alternatives, a Comprehensive Guide (2nd Edition)
Quantable Analytics © 2026. All Rights Reserved.
Powered by WordPress . Designed by
