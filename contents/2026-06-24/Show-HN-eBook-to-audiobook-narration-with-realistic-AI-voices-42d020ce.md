---
source: "https://ebookaloud.com"
hn_url: "https://news.ycombinator.com/item?id=48661083"
title: "Show HN: eBook to audiobook narration with realistic AI voices"
article_title: "EbookAloud — Turn Epubs into Audiobooks"
author: "flatline"
captured_at: "2026-06-24T15:46:38Z"
capture_tool: "hn-digest"
hn_id: 48661083
score: 3
comments: 0
posted_at: "2026-06-24T15:04:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: eBook to audiobook narration with realistic AI voices

- HN: [48661083](https://news.ycombinator.com/item?id=48661083)
- Source: [ebookaloud.com](https://ebookaloud.com)
- Score: 3
- Comments: 0
- Posted: 2026-06-24T15:04:54Z

## Translation

タイトル: HN を表示: リアルな AI 音声による電子書籍からオーディオブックへのナレーション
記事のタイトル: EbookAloud — Epub をオーディオブックに変える
説明: プライバシーを最優先した電子ブックからオーディオブックへの変換。 DRM フリーの電子ブックをアップロードし、音声を選択して、M4B オーディオブックをダウンロードします。
HN テキスト: しばらくの間、長い形式のナレーション用に新しい AI 音声を試してみたいと思っていましたが、見つけたものはすべてサブスクリプションが必要で、使用量が限られていました。私はオープン Kokoro モデル [0] に出会いました。その音声は非常に優れており、従来のロボットのような TTS 音声から得られる疲労を感じることなく、何時間でも聞いていても十分です。モデルは 82m パラメーターで、高速に実行されるように設計されていますが、12 コアのラップトップで CPU 推論から適切な時間を得るのに依然として苦労しました。クラウドベースの GPU サービスを使用すれば、自分のセルフホスト ライブラリにフィードするのに十分な速度でオーディオブックを生成できるようになり、その同じパイプラインが他の人も使用できる製品になる可能性があると考えました。これを構築する際の目標は 2 つありました。AI マルチエージェント コーディング ワークフローをある程度理解することと、電子ブックからオーディオブックへの変換に特化した TTS 製品を構築することです。 ebookaloud の 99% は OpenCode の DeepSeek v4 によって作成されました。私は 1 か月間で約 7 億 5,000 万のトークン (クレジットで 12 ドル相当) を使用しましたが、その結果には非常に満足しています。すべての変更/機能は、Pro エージェントと Flash エージェントを組み合わせて、計画 -> 実装 -> テスト -> レビュー -> 修正 -> コミットのサイクルを経ました。これは通常、同時実行ワーカーが 1 人か 2 人に制限されていました。抽出および合成パイプラインのさまざまな部分の品質管理のために、一度に 8 ～ 10 個を実行できる別個の評価エージェントを用意しました。 AI ワークフローの自動化という点では、Yegge のステージ 6 [1] に近づいているかもしれません。その後、Claude Code をセットアップし、Opus 4.8 を DeepSeek と並行して実行しました。明らかに品質が違います

ces ですが、私は実践的なアプローチを持つ経験豊富な開発者です。私はコードをまったく書きませんでしたが、コードが生成したものの重要なセクションを読み、アプローチの各ステップについて DS Pro と広範な会話を行いました。 Opus は DeepSeek の選択についてあまり批判的なことを言っていませんでしたし、フロンティア モデルが私のワークフローに大きな変化をもたらしたとは思えません。大規模なコードベースでは違いがより明らかになるのではないかと思いますが、Opus で実装したいくつかの変更には、私が使用したすべてのモデルと同様の問題がありました。私の指示のないランダムな変更、複雑すぎる単純なソリューション、行き詰まったときに予期しない/破壊的なアクションを取るなどです。Opus は、私が関与していた複雑な計画とオーケストレーションをより多く処理できることがわかりました。それは私が望むこともありますが、他のものでは望まないことです。製品自体に関しては、より洗練されたソリューションがたくさんあります。私はイレブンラボと競争するつもりはありません。私は、従量課金制の価格モデルと十分な出力品質を備えたシームレスなオーディオブック体験を目的として、m4b 世代をターゲットにしています。これは私が商品化しようとした初めての製品であり、AI コード生成によって洗練されたものが手の届くところにありました。 AI がなかったら、手動での研究開発に 6 ～ 8 か月かかり、完成するずっと前に燃え尽きてしまっていたでしょう。音声/フォーマットに関して生成されるものを確認したいだけであれば、サイトのトップページに無料サンプルがあります。私は出力品質に関していくつかの独断的な決定を下しました。業界標準に一致するように、ほとんどの音声で 140 wpm を目指しましたが、一部の音声はそれよりも遅いか速く、その速度では韻律が失われます。ユーザーに選択肢の壁を与えるのではなく、速度などについては再生デバイスに任せます。

dコントロール。このサイトが実際に使用されるようになったら、Kokoro の他の言語をサポートするように拡張したいと思います。PDF からの抽出と合成により、製品がかなり充実したものになるでしょう。 [0] https://github.com/hexgrad/kokoro [1] https://steve-yegge.medium.com/welcome-to-gas-town-4f25ee16d...

記事本文:
EbookAloud — Epub をオーディオブックに変える
電子ブック大声で
電子ブックをオーディオブックに変える
電子ブックをアップロードし、音声を選択し、Apple Books、AudioBookshelf などで再生可能な標準 M4B オーディオブックを入手します。透明性のある価格設定で、アカウントやサブスクリプションは必要ありません。
またはクリックして参照 · epub (最良の結果)、docx、html、md、txt など
[アップロードと変換] をクリックすると、サービス利用規約とプライバシー ポリシーに同意したことになります。
パブリックドメイン作品の無料サンプルオーディオブックをダウンロードする
不思議の国のアリス、
エマのナレーション。
自然な響きのココロの声からお選びください。音声をクリックするとサンプルが聞こえます。
私は何年もの間、著作権の切れた電子書籍、会議録、自分の文章など、声に出して読みたいと思っていた文書を集めてきました。古いテキスト読み上げモデルは長い形式のコンテンツには耐えられませんでしたが、新しい AI 音声は本当に優れています。ココロの声はオープンで信頼でき、プロフェッショナルで、何時間でも聞いていられるほどです。
すでに多くの TTS サービスが存在しており、中にはこれよりも多くの機能を備えたサービスもあります。ただし、私が見つけたものはすべてサブスクリプションが必要で、価格体系が不透明でした。推論は、私のラップトップには搭載されていない優れた GPU を使用すれば比較的安価であり、電子ブックからのテキスト抽出パイプラインでは、良好な結果を得るために実際の作業が必要でした。私は、コンバージョンあたりわずか 1 ドル、単語ごとの透明な価格設定、サブスクリプション不要、アカウント不要、既存のオーディオブック プログラムとの優れた統合など、私が探していたサービスとして EbookAloud を構築しました。
ご質問またはフィードバックはありますか?連絡先:
オープンソースのクレジットとライセンス情報

## Original Extract

Privacy-first ebook to audiobook conversion. Upload your DRM-free ebook, choose a voice, and download your M4B audiobook.

For a while I've wanted to try out the new AI voices for long-form narration, but everything I found required a subscription that didn't justify my limited usage. I came across the open Kokoro model [0] and the voices are very good -- good enough to listen to for hours without the fatigue I got from legacy, robotic TTS voices. The model is 82m parameters and designed to run fast, but I still struggled to get reasonable times from CPU inference on my 12-core laptop. I thought a cloud-based GPU service would let me generate audiobooks fast enough to feed my own self-hosted library, and that same pipeline could become a product other people could use. I had two goals in building this: get some exposure to AI multi-agent coding workflows, and build a TTS product targeting ebook to audiobook conversion specifically. 99% of ebookaloud was written by DeepSeek v4 in OpenCode. I've used about 750 million tokens costing $12 in credits over the course of a month, and I'm very pleased with the results. Every change/feature went through a plan -> implement -> test -> review -> correct -> commit cycle with a mix of Pro and Flash agents. This was generally limited to one or two concurrent workers. I had a separate eval agent for quality control on various parts of the extraction and synthesis pipeline, which I could run 8-10 at a time. I may be approaching Yegge's Stage 6 [1] in terms of AI workflow automation. I later set up Claude Code and ran Opus 4.8 side by side with DeepSeek. There are definitely quality differences, but I'm an experienced developer with a hands-on approach. I didn't write any of the code, but I have read critical sections of what it generated and had extensive conversations with DS Pro about each step of the approach. Opus didn't have much critical to say about DeepSeek's choices, and I'm not convinced a frontier model would have made an appreciable difference for my workflow. I suspect on a large codebase the differences would become more apparent, but the few changes I implemented with Opus had similar issues to all the models I've used: random changes without my direction, over-complicating simple solutions, taking unanticipated/destructive actions when it gets stuck, etc. I do see Opus being capable of handling more of the complex planning and orchestration that I was involved in. That's something I may want sometimes but not others. As to the product itself, there are a lot more sophisticated solutions out there. I'm not trying to compete with ElevenLabs. I'm targeting m4b generation for a seamless audiobook experience with a pay-as-you-go pricing model and good-enough output quality. This is the first product I've ever tried to commercialize, and AI code generation put something polished within reach. Without AI, this would have taken me 6-8 months of manual research and development, and I would have gotten burned out long before completing it. I have a free sample on the front page of the site if you just want to see what it generates in terms of voice/format. I made a few opinionated decisions regarding output quality. I aimed for 140 wpm for most of the voices to match industry standards, but some are much slower or faster and lose prosody at that rate. Rather than give users a wall of options, I'm deferring to the playback device for things like speed control. If the site sees real usage I'd like to expand to support Kokoro's other languages, and extraction and synthesis from PDF would round out the product quite a bit. [0] https://github.com/hexgrad/kokoro [1] https://steve-yegge.medium.com/welcome-to-gas-town-4f25ee16d...

EbookAloud — Turn Epubs into Audiobooks
EbookAloud
Turn Your Ebook Into an Audiobook
Upload an ebook, pick a voice, and get a standard M4B audiobook playable in Apple Books, AudioBookshelf, and more. Transparent pricing, no account or subscription required.
or click to browse · epub (best results), docx, html, md, txt, and more
By clicking Upload & Convert, you agree to our Terms of Service and Privacy Policy .
Download a free sample audiobook of the public domain work
Alice's Adventures in Wonderland ,
narrated by Emma.
Pick from our selection of natural-sounding Kokoro voices. Click a voice to hear a sample.
For years I've collected written material I've wanted read aloud: e-books that are out of copyright, conference notes, my own writing, and more. The old text-to-speech models were not tolerable for long-form content, but new AI voices are genuinely good. The Kokoro voices are open, reliable, and professional — good enough to listen to for hours.
There are many TTS services out there already, some with many more features than this one. However, all the ones I've found require subscriptions and had opaque pricing schemes. Inference is relatively cheap with a good GPU, which I don't have on my laptop, and the text extraction pipeline from ebook required real work to get good results. I built EbookAloud as the service I was looking for: as little as $1 per conversion, transparent per-word pricing, no subscription, no account required, and good integration with existing audiobook programs.
Questions or feedback? Get in touch:
Open source credits and licensing information
