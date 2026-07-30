---
source: "https://arstechnica.com/ai/2026/07/tested-google-synthid-works-great-but-labeling-ai-content-may-be-a-losing-game/"
hn_url: "https://news.ycombinator.com/item?id=49107692"
title: "Google's SynthID watermark is hard to break, but doesn't solve AI disinformation"
article_title: "Google's SynthID watermark is hard to break, but it doesn't solve AI disinformation - Ars Technica"
author: "thm"
captured_at: "2026-07-30T10:14:10Z"
capture_tool: "hn-digest"
hn_id: 49107692
score: 1
comments: 1
posted_at: "2026-07-30T09:25:13Z"
tags:
  - hacker-news
  - translated
---

# Google's SynthID watermark is hard to break, but doesn't solve AI disinformation

- HN: [49107692](https://news.ycombinator.com/item?id=49107692)
- Source: [arstechnica.com](https://arstechnica.com/ai/2026/07/tested-google-synthid-works-great-but-labeling-ai-content-may-be-a-losing-game/)
- Score: 1
- Comments: 1
- Posted: 2026-07-30T09:25:13Z

## Translation

タイトル: Google の SynthID ウォーターマークは破るのが難しいが、AI の偽情報は解決しない
記事タイトル: Google の SynthID ウォーターマークは破るのが難しいが、AI の偽情報は解決しない - Ars Technica
説明: インターネット上で何が本物かを判断することは、将来的には簡単ではなくなるでしょう。

記事本文:
コンテンツにスキップ
アルス テクニカ ホーム
セクション
フォーラム
購読する
検索
AI
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ストーリーにピンを付ける
テーマ
ハイパーライト
検索
サインイン
サインインダイアログ...
サインイン
現実の問題
Google の SynthID ウォーターマークは破るのが難しいですが、AI の偽情報は解決できません
将来的には、インターネット上で何が本物かを判断するのは簡単ではなくなるでしょう。
79
クレジット:
オーリッヒ・ローソン
クレジット:
オーリッヒ・ローソン
テキスト
設定
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ナビゲーション用に最小化する
AI が生成するメディアの規模を把握するのは難しい場合があります。
スタンフォード大学と南カリフォルニア大学の共同研究であるスターリング研究所は、人類が 15 億枚の画像を作成するのに 1975 年 (カメラの発明から 149 年後) までかかったと推定しています。生成 AI が同じことを行うのにわずか 18 か月かかりました。 AI はそこで終わりませんでした。
今年の春、Google は I/O で、自社のツールがわずか 2 年間で 1,000 億以上の AI 画像とビデオの作成に使用されたと発表しましたが、Google が AI コンテンツの唯一のソースというわけではありません。同社は、この驚異的な統計と多数のパートナーシップを組み合わせて、SynthID 透かしテクノロジーの使用を拡大しました。このテクノロジーは、AI が生成したコンテンツにラベルを付けるのに使用でき、人々が何が本物かを識別できる可能性があります。
同社は、SynthID は編集に耐えられるほど堅牢であると述べており、OpenAI、Runway、Nvidia などが使用を開始するにつれて、SynthID がテストされるのを目にすることになるでしょう。目に見えない透かしは AI コンテンツの難問を解決できるでしょうか?
現在、AI コンテンツのラベル付けには 2 つのアプローチがあります。SynthID のような目に見えないウォーターマークと、Coalition for Content Provenance and Authenti のようなメタデータ スキーマです。

市 (C2PA)。 Google は AI コンテンツにラベルを付けるために両方を使用していますが、C2PA は透かしのように目に見えないところに隠すことを目的としたものではありません。暗号的に安全であるため、偽造することはできませんが、C2PA は簡単に削除できます。画像を編集して保存するか、スクリーンショットを撮るだけで C2PA を削除できます。
SynthID のようなウォーターマークは、画像やビデオのピクセル、またはオーディオ クリップの波形にエンコードされます。インターネット上で渡されるコンテンツは圧縮、サイズ変更、編集によって劣化するが、Google は使い古されたミームであっても SynthID が存在するはずだと述べている。 Google DeepMind の科学者である Pushmeet Kohli 氏は、Ars に対し、チームは SynthID が AI コンテンツにラベルを付ける永続的な方法であることを保証するために多大な労力を費やしたと説明しました。
「開発プロセス全体を通じて、私たちはこのようなテクノロジーが攻撃されることをある程度想定していました」とコーリ氏は言います。 「そこで私たちは、SynthID をさまざまな種類の変換に対して堅牢にするために多くの研究を行いました。ユーザーが何らかのフィルターを追加する場合でも、画像をトリミングする場合でも、私たちはこれらの変換を使用し、検出器がそれらに対して堅牢であることを確認しました。」
Google は、SynthID の機能について元の論文以外の技術的な詳細をあまり提供することに消極的であるため、この技術が AI 環境全体に拡大するにつれて、これらの主張をテストする価値はあります。しかし、AI 画像がインターネット上で流通するにつれて自然に劣化するのを待つ時間がある人がいるでしょうか?
Python Pillow ライブラリを使用して、非常に高速な共有とダウンロードの繰り返しによるデータ損失をシミュレートしました。 Python スクリプトは、設定された範囲内でランダムな圧縮値とサイズ変更値を選択してテスト イメージに適用し、出力イメージを次の反復の基礎として使用します。
このテストは 2 種類の AI 画像に焦点を当てました。1 つはモデルによって完全にゼロから作成されました。

AIによって編集されたオリジナルの写真。 Nano Banana Pro によって生成された両方の画像には SynthID ウォーターマークが含まれており、メタデータをいくら編集してもこれを隠すことはできません。
完全にAIが生成したテスト画像。
Google AI
Google AIで編集したテスト画像。
Google AI
Google AIで編集したテスト画像。
Google AI
完全にAIが生成したテスト画像。
Google AI
Google AIで編集したテスト画像。
Google AI
オリジナルのテストと編集されたテスト用に、上の画像から始めました。政治はメタルではないし、私の犬はドラゴンの一部ではありません（私の言葉を信じてください）。これらの画像を画像粉砕機に何百回もかけて実行し、鮮明なオリジナルをかろうじて認識できる塊に変えました。 50 世代ごとに、画像のトリミングされたバージョン (再エンコードなし) も作成し、潜在的な SynthID ピクセルを削除して、検出をさらに弱めました。
300 世代の共有をシミュレートした後、次のようなフルフレーム画像が残ります。
300 圧縮サイクル後の完全に AI 生成のイメージ (SynthID はそのまま)。
Google AI
300 圧縮サイクル後の AI 編集画像 (SynthID はそのまま)。
Google AI
300 圧縮サイクル後の AI 編集画像 (SynthID はそのまま)。
Google AI
300 圧縮サイクル後の完全に AI 生成のイメージ (SynthID はそのまま)。
Google AI
300 圧縮サイクル後の AI 編集画像 (SynthID はそのまま)。
Google AI
そして、見てください。SynthID ウォーターマークは両方でまだ機能しています。画像を Gemini にアップロードし、SynthID チェックを依頼することで、自分で確認できます。画像全体のスクリーンショットを撮ることもできますが、特別なピクセルが新しいファイルに転送されるため、システムは引き続きその画像に SynthID というラベルを付けます。
SynthID ピクセルは画像全体に分散されているため、トリミングに対しても、少なくともある程度は耐性があります。これはw

ここでついに SynthID の限界がわかります。 300 回の圧縮生成の後、テスト画像の境界線から数ピクセルを削除すると、最終的に SynthID (編集済みおよび完全な AI の両方) が壊れました。これらのバージョン (下記を参照) は 20% トリミングされていますが、SynthID を検出できなくなるには十分です。 50% クロップを大きくすると、画像圧縮を 250 回ほど繰り返すと、少し早く SynthID が壊れる可能性があります。
この圧縮 AI イメージの一部を切り取って、SynthID を削除することに成功しました。
Google AI
少量の収穫の後、この高度に圧縮された AI 編集には SynthID がなくなりました。
Google AI
少量の収穫の後、この高度に圧縮された AI 編集には SynthID がなくなりました。
Google AI
この圧縮 AI イメージの一部を切り取って、SynthID を削除することに成功しました。
Google AI
少量の収穫の後、この高度に圧縮された AI 編集には SynthID がなくなりました。
Google AI
これらの画像は醜いかもしれませんが、（コンテンツ以外に）AI であることを証明するものは何もありません。
私のテストによると、SynthID は徐々に劣化する可能性がありますが、そうなる頃には画像が無駄にぼやけてしまいます。では、AIの偽情報の問題は解決したのでしょうか？それほど速くはありません。 SynthID が強引な編集やデータ損失にも耐えられるのは印象的ですが、AI ウォーターマークにはいくつかの注目すべき問題があります。
AI 透かしテクノロジーは SynthID だけではありません。 Meta は最近、Content Seal AI ウォーターマークをリリースしましたが、ロイターは、画像を少しトリミングするだけでウォーターマークが削除されることが多いことを発見しました。 SynthID が現時点では堅実であるように見えるからといって、AI コンテンツに目に見えないラベルを付ける試みがすべてうまくいくというわけではありません。
また Google は、SynthID が無敵であるとは主張していません。この技術に関する元の論文では、透かしは敵対的な攻撃に耐えることを目的としたものではないと述べています。十分な時間とやる気のある人が SynthID をバイパスする方法を見つける可能性は十分にあります。それで

SynthID のハッキングに成功したとすでに人々が主張していますが、テストを通じてそれを確認することはできませんでした。 Google のコーリ氏も、彼のチームは想定されている回避策を再現できなかったと述べています。
それでも、セキュリティは決して完璧ではありません。 SynthID のバイパスに成功した場合 (人気が高まり続けるのであれば、それはもっともらしいことですが)、数十億の AI 画像とビデオすべてのラベルが瞬時に解除される可能性があります。 Googleはこうした攻撃をブロックするための標準を更新する中で、いたちごっこに陥るかもしれない。 Google は、SynthID がデジタル プロビデンスに対する特効薬ではないとすぐに指摘しています。しかし、SynthID はさらに大きなターゲットになりつつあります。 Google と OpenAI の間では、消費者向け AI 分野のほとんどが現在このテクノロジーを使用しています。
SynthID の背後にあるチームはハッキングに対抗するための計画を立てているが、コーリ氏は、Google はその多くを自社内に留めておく必要があると述べている。 「もちろん、透かしを変更することもできます」とコーリ氏は言う。 「私たちが自由に使えるものはたくさんあります。これらの緩和策の多くを公に公開することはできません。なぜなら、それらはシステム全体を確実に保護するのに役立つからです。」
SynthID を使用してみると、SynthID のセキュリティ モデルの少なくとも一部が明らかになります。 Google は人々に SynthID を攻撃する簡単な道を与えたくないため、アクセスは意図的に制限されています。疑わしい AI 画像を見つけた場合は、Gemini に検証者を呼び出すように依頼する必要があります。API や公的にアクセスできる SynthID 検出器の Web ページはありません。
制限はそれだけではありません。ユーザーがバイパス ワークフローを調整するために連続した AI チェックを使用することを防ぐため、Google は SynthID を 1 日あたり「約 10 回の画像チェック」に制限しています。私のテストでは、似たような画像をアップロードしすぎると、検証者はさらに早くロックアウトしてしまいます。画像が AI によって複数回生成されているのではないかと疑う場合

特定の日には、運が悪いかもしれません。政治家が喜んで AI を利用して敵対者を中傷したり攻撃したりしている現在、真実を検証するために冷却期間を設けるべきではありません。
AI 検出ツールからロックアウトされていないとしても、どのツールを使用すべきでしょうか? Googleがこの技術を他の企業に展開するにつれて、ウォーターマークは急速に断片化している。 OpenAI や Runway などの企業は、基盤となる SynthID テクノロジーを使用して AI コンテンツにラベルを付けることを選択していますが、実際のウォーターマークは異なります。たとえば、Google の検出器は OpenAI のウォーターマークを認識せず、OpenAI の検出器は Google のウォーターマークを認識しません。そのため、SynthID へのアクセスが制限されているだけでなく、AI 画像を複数の検出器で実行しても何も見つからない可能性があります。これは、画像が知らない別のシステムによって生成されたものであるためです。
Googleの広報担当者は、チームはこのエクスペリエンスが最適ではないことを認識しており、それを改善するための漠然とした計画があると述べた。 「私たちは検証をよりアクセスしやすく統一するために、業界パートナーと積極的に協力しています。私たちのビジョンは、さまざまなプラットフォーム間でよりシームレスな体験を可能にする、より相互運用性の高いエコシステムを構築することです」と彼らは述べています。
ウォーターマークによるユーザビリティの問題は解決できるかもしれません。しかし、それはせいぜい、何が本物かを判断するための 1 つの要素にすぎません。 AI のウォーターマークは、1 つの単純な理由から決して最善の方法ではありません。それは、ラベルを付けずに AI コンテンツを作成する方法が常に存在するということです。
数え切れないほどの企業が画像やビデオの生成を提供していますが、そのすべてが透かし技術やメタデータのラベル付けの使用に同意する可能性は低いです。さらに重要なのは、AI スロップを生成するためにビッグテックの集中システムに依存する必要がないことです。
「ここでのより大きな問題は私だけではないと思います

スターリング・ラボのフェロー兼シニアアドバイザーであるアダム・ローズ氏は、「問題は、自分のモデルを実行できる人、自分のコンピュータで何かを実行できる人だ」と語った。
現在でも、ラベルをまったく付けずに画像を生成するオープン モデルがたくさんあります。主要な AI プレーヤーのほとんどが AI コンテンツのラベル付けに多面的なアプローチを採用している場合でも、これらの AI モデルは無期限に共有および改善することができます。生成の魔神がボトルから出てきたので、それを再び詰め込む必要はありません。
ラベルのない AI コンテンツは常に存在するため、何が真実かを判断するために透かしを当てにすることはできません。 AI によるラベル付けがより一般的になるにつれて、人々は、ラベル付けが欠けている画像やビデオが自動的に正当であると考えるかもしれません。それは危険ですが、誰かが犯す間違いは理解できます。目に見えない透かしの存在は、ますます歪んだインターネットを探索する人々にある程度の安心感を与えるかもしれませんが、何が真実かを知る信頼できる方法ではありません。
「私たちが住む未来は、ますます無制限のコンテンツの世界になるでしょう」とローズ氏は言います。 「それは経済学 101、需要と供給です。供給は天井知らずになり、コンテンツの価値は暴落し続けるだろう。しかし、認証されたコンテンツ、つまり証明する方法の需要がますます高まるでしょう。

[切り捨てられた]

## Original Extract

Deciding what's real on the Internet won't be easy in the future.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
Real problems
Google’s SynthID watermark is hard to break, but it doesn’t solve AI disinformation
Deciding what’s real on the Internet won’t be easy in the future.
79
Credit:
Aurich Lawson
Credit:
Aurich Lawson
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
The scale of AI-generated media can be hard to grasp.
Starling Lab, a research collaboration from Stanford University and the University of Southern California, estimates that it took until 1975—149 years after the invention of the camera—for humanity to create 1.5 billion images. It took generative AI just 18 months to do the same. And AI didn’t stop there.
This spring, Google announced at I/O that its tools had been used to create more than 100 billion AI images and videos in just a couple of years, and Google is far from the only source of AI content. The company coupled this staggering statistic with a slew of partnerships to expand the use of its SynthID watermarking technology, which can be used to label AI-generated content and potentially help people identify what’s real.
The company says SynthID is robust enough to survive edits, and we’re about to see that tested as OpenAI, Runway, Nvidia, and others begin using it. Can invisible watermarks solve the AI content conundrum?
There are currently two approaches to labeling AI content: invisible watermarks like SynthID and metadata schemas like the Coalition for Content Provenance and Authenticity (C2PA). Google uses both to label its AI content, but C2PA is not intended to hide in plain sight like a watermark. It’s cryptographically secure, so you can’t fake it, but C2PA is trivially easy to strip out. Simply editing and saving an image or taking a screenshot can remove C2PA.
Watermarks like SynthID are encoded in the pixels of an image or video or in the waveform of an audio clip. Content that gets passed around the Internet degrades from compression, resizing, and edits, but Google says SynthID should still be present even in well-worn memes. Google DeepMind scientist Pushmeet Kohli explained to Ars that the team put a great deal of work into ensuring SynthID was a durable way to label AI content.
“Through the whole development process, we sort of assumed that a technology like this will be attacked,” said Kohli. “So we did a lot of research in making SynthID robust to different kinds of transformations. Whether people are adding some sort of filter or cropping the image, we used these transformations and made sure that the detector was robust against them.”
Google has been reluctant to provide much technical detail on SynthID’s functionality beyond the original paper, so it’s worth testing these claims as the technology expands across the AI landscape. But who has the time to wait for AI images to degrade naturally as they get passed around the Internet?
I used the Python Pillow library to simulate data loss from repeated sharing and downloading at a vastly accelerated rate. The Python script chooses random compression and resizing values within set ranges to apply to the test image, then uses the output image as the basis for the next iteration.
This test focused on two types of AI images: one created entirely from scratch by the model and an original photo edited by AI. Both images generated by Nano Banana Pro include the SynthID watermark, and no amount of metadata editing can hide it.
A completely AI-generated test image.
Google AI
A test image edited with Google AI.
Google AI
A test image edited with Google AI.
Google AI
A completely AI-generated test image.
Google AI
A test image edited with Google AI.
Google AI
I started with the images above for the original and edited tests. We can probably all agree that neither one is real—politics is not metal, and my dog is not part dragon (take my word for it). I ran these images through the image-crushing machine hundreds of times, turning the crisp originals into barely recognizable blobs. Every 50 generations, I also created cropped versions of the images (without reencoding), removing potential SynthID pixels to further weaken the detection.
After 300 generations of simulated sharing, we’re left with these full-frame images:
A fully AI-generated image after 300 compression cycles (SynthID intact).
Google AI
An AI-edited image after 300 compression cycles (SynthID intact).
Google AI
An AI-edited image after 300 compression cycles (SynthID intact).
Google AI
A fully AI-generated image after 300 compression cycles (SynthID intact).
Google AI
An AI-edited image after 300 compression cycles (SynthID intact).
Google AI
And look at that—the SynthID watermark still works on both of them. You can check for yourself by uploading the images to Gemini and asking for a SynthID check. You can even take a screenshot of the full image, and the system will still label it as SynthID because the special pixels transfer over to the new file.
Because SynthID pixels are spread throughout the image, they’re also resistant to cropping… at least to a point. This is where we finally find the limits of SynthID. After 300 compression generations, removing a few pixels from the border of our test images finally broke SynthID (both edited and fully AI). These versions (see below) have been cropped by 20 percent, but that’s enough to render SynthID undetectable. A larger 50 percent crop can break SynthID a bit earlier, at around 250 image compression iterations.
A small crop of this compressed AI image successfully removes SynthID.
Google AI
After a small crop, this heavily compressed AI edit no longer has SynthID.
Google AI
After a small crop, this heavily compressed AI edit no longer has SynthID.
Google AI
A small crop of this compressed AI image successfully removes SynthID.
Google AI
After a small crop, this heavily compressed AI edit no longer has SynthID.
Google AI
These images may be ugly, but there’s nothing (other than the content) to prove they’re AI.
Based on my testing, SynthID can gradually degrade, but the images are uselessly blurry by the time that happens. So have we solved the problem of AI disinformation? Not so fast. It’s impressive that SynthID can survive aggressive edits and data loss, but AI watermarks have some notable issues.
SynthID is not the only AI watermarking technology. Meta released its Content Seal AI watermark recently, but Reuters found that simply cropping the images a bit would often eliminate the watermark . Just because SynthID seems solid right now doesn’t mean all attempts to invisibly label AI content will work.
Google also doesn’t claim that SynthID is invulnerable—the original paper on the technology notes that the watermark is not intended to withstand adversarial attacks. It’s entirely possible that someone with enough time and incentive will find a way to bypass SynthID. Some people have already claimed to have successfully hacked SynthID, but we have been unable to confirm that through testing. Google’s Kohli also says his team has been unable to replicate the supposed workarounds.
Still, security is never perfect. If SynthID is successfully bypassed, which seems plausible if it continues to grow in popularity, all those billions of AI images and videos could be unlabeled in a snap. Google may find itself in a cat-and-mouse game as it updates the standard to block those attacks. Google is quick to point out that SynthID is not a silver bullet for digital providence. But SynthID is becoming a much bigger target. Between Google and OpenAI, most of the consumer AI space is now using this technology.
The team behind SynthID has plans in place to combat hacks, but Kohli says Google needs to keep much of that to itself. “We can, of course, change the watermark,” said Kohli. “There are a lot of things at our disposal. I cannot publicly disclose many of these mitigations because they help us make sure that the whole system is protected.”
At least some of the security model for SynthID becomes apparent when you try to use it. Google doesn’t want to give people an easy path to attack SynthID, so access is intentionally limited. If you see a suspicious AI image, you have to ask Gemini to call the verifier—there’s no API or publicly accessible SynthID detector webpage.
That’s not the only limit, either. To prevent people from using successive AI checks to tune a bypass workflow, Google limits SynthID to “approximately 10 image checks” per day. In my testing, the verifier will lock you out even faster if you upload too many similar-looking images. If you’re suspicious that an image may be generated by AI more than a few times on a given day, you may be out of luck. At a time when politicians are gleefully using AI to defame and attack opponents , verifying the truth shouldn’t come with a cooldown period.
Even if you aren’t locked out of AI detection tools, which one are you supposed to use? As Google rolls out the technology to other firms, watermarks are quickly becoming fragmented. While companies like OpenAI and Runway have chosen to use the underlying SynthID technology to label AI content, the actual watermarks are different. For instance, Google’s detector doesn’t recognize OpenAI’s watermark, and OpenAI’s detector doesn’t recognize Google’s watermark. So not only is access to SynthID restricted, but you may run an AI image through multiple detectors and still find nothing because it was generated by a different system you don’t even know about.
A Google spokesperson says the team is aware that this experience is suboptimal, and it has vague plans to improve it. “We are actively collaborating with industry partners to make verification more accessible and unified. Our vision is to create a more interoperable ecosystem that allows for a more seamless experience across different platforms,” they said.
Maybe the usability issues with watermarks can be solved. But at best, it will be only one element of how we determine what’s real. Watermarks in AI will never be the best way for one simple reason: There will always be ways to create AI content without any labeling.
Countless companies offer image and video generation, and it’s unlikely they will all agree to use watermarking technology or even metadata labeling. More importantly, you don’t have to rely on Big Tech’s centralized systems to generate AI slop.
“I think the bigger problem here is not just the images that are being produced by the Googles and the OpenAIs of the world,” said Adam Rose, a fellow and senior advisor at Starling Lab. “The problem is people who can run their own models, who can do things on their own computers.”
Even today, there are plenty of open models that generate images with no labeling whatsoever. These AI models can be shared and improved upon indefinitely, even if most of the major AI players adopt a multifaceted approach to labeling AI content. The generative genie is out of the bottle, and there’s no stuffing it back in.
Since there will always be unlabeled AI content, you can’t count on watermarks to tell you what’s true. As AI labeling becomes more common, people may think an image or video that lacks one is automatically legit. That’s a dangerous but understandable mistake for someone to make. While the existence of invisible watermarks might give people some sense of security as they explore an increasingly distorted Internet, they aren’t a reliable way to know what’s true.
“The future that we live in will increasingly be a world of unlimited content,” said Rose. “It’s economics 101, supply and demand. Supply goes through the roof, and the value of content is going to continue to crash. But increasingly, there is going to be a demand for authenticated content, a way to prove

[truncated]
