---
source: "https://photoinvestigator.co/blog/how-to-tell-if-a-photo-is-ai-generated-metadata/"
hn_url: "https://news.ycombinator.com/item?id=48486833"
title: "Is It AI? How to Tell Using Metadata"
article_title: "Is it AI? How to Tell Using Metadata - The Photo Investigator"
author: "Danbana"
captured_at: "2026-06-11T07:50:09Z"
capture_tool: "hn-digest"
hn_id: 48486833
score: 1
comments: 0
posted_at: "2026-06-11T06:17:02Z"
tags:
  - hacker-news
  - translated
---

# Is It AI? How to Tell Using Metadata

- HN: [48486833](https://news.ycombinator.com/item?id=48486833)
- Source: [photoinvestigator.co](https://photoinvestigator.co/blog/how-to-tell-if-a-photo-is-ai-generated-metadata/)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T06:17:02Z

## Translation

タイトル：AIですか？メタデータを使用して伝える方法
記事タイトル：AIなのか？メタデータを使用して伝える方法 - The Photo Investigator
説明: AI 画像検出器が騙されます。メタデータはそうではありません。ここでは、ジェネレーター自体に埋め込まれた信号を使用して、写真が DALL-E、Midjourney、または Adob​​e Firefly によって生成されたかどうかを確認する方法を示します。

記事本文:
AIですか？メタデータを使用して伝える方法 - The Photo Investigator
コンテンツにスキップ
写真調査官
写真のメタデータを使用した探索
AIですか？メタデータを使用して伝える方法
2023年3月、オープンソースの捜査機関ベリングキャットの創設者エリオット・ヒギンズは、ドナルド・トランプの逮捕についてミッドジャーニーにいくつかのプロンプトを流した。彼は結果として得られた画像を公開し、AI が生成したものであると明確にラベル付けし、それらが本物のニュースとして数時間以内に拡散するのを観察しました。ヒギンズ氏は、本物と偽物の境界線がいかに薄くなっているかを示すためにこれらを作ったと語った。彼は正しかった。彼が画像に付けたラベルは、最初のスクリーンショットには残っていませんでした。
2か月後、ブルームバーグニュースになりすましたチェックマーク認証済みのTwitterアカウントが、国防総省での爆発の写真のように見えるものを投稿した。トレーダーらが画像がAIによって生成されたものであることに気づくまで、S&P 500は5分間で0.3%下落した。市場は1時間以内に回復したが、明らかな偽物は市場に実際の影響を与えた。
これらの事件は、バレンシアガの白いダウンコートを着た法王、AIが生成したマウイ島の山火事の「証拠」写真、2024年の選挙サイクルの偽のトランプ・ハリス集会画像とともに、すべて同じように展開した。画像が拡散します。調査は続きます。真実は、たとえあったとしても、遅れて追いつきます。それまでに、信じたかった視聴者は先に進んでしまいましたが、多くの場合、Photo Investigator で写真のメタデータを確認することで真実が明らかになります。
ほとんどの AI ジェネレーターが残すメタデータ
実際、驚くほど多くの AI 画像ジェネレーターが、自発的に出力にラベルを付けています。ピクセル単位ではなく、メタデータ フィールドに、ジェネレーターはファイルを保存するときにファイルに直接書き込みます。
Coalition for Content Provenance and Authenticity は、Adobe、Microsoft、

OpenAI、Google、および主要なカメラ メーカーは、コンテンツ認証情報の標準を定義しました。 C2PA マニフェストは、イメージ ファイルに埋め込まれた暗号化署名されたパケットです。画像の作成者または内容、作成者が適用した編集、ファイルを作成したソフトウェアが記録されます。 OpenAI は、2024 年初頭に C2PA を DALL-E 3 出力に追加しました。Adobe Firefly は、生成するすべての画像にタグを付けます。 Google ImageFX には同じ標準が組み込まれています。マイクロソフトデザイナーも。イメージに有効な C2PA マニフェストがある場合は、どのモデルがいつ生成したかを正確に読み取ることができます。
IPTC 写真メタデータ標準には、デジタル ファイルの作成方法を宣言するために設計された「 Iptc4xmpExt:DigitalSourceType 」と呼ばれる特定の XMP フィールドが含まれています。このフィールドに「trainedAlgorithmicMedia」と表示されている場合、ファイルは生成モデルからのものです。 「digitalCapture」はカメラがそれをキャプチャしたことを意味します。標準に準拠する AI ジェネレーターは、エクスポート時にこのフィールドを書き込みます。
最後に、これは最も古く、最も粗末な信号です。実際のカメラは、そのメーカーとモデルを EXIF ソフトウェア フィールドに書き込みます。 AI ジェネレーターも時々そうします。 「Midjourney v6.1」または「Stable Diffusion XL」と表示されているソフトウェア フィールドが、検出作業を行っています。ジェネレーター自体がタグ付けしない場合でも、出力の保存に使用される Python ライブラリがタグ付けすることがあります。PIL/Pillow や OpenCV などのエントリが iPhone スナップショットであるはずのものに表示されるのはシグナルです。
どのジェネレータが出力にタグを付けるか (2026)
2026 年 6 月時点で確認されている動作を反映しています。ジェネレーターの動作はソフトウェアの更新により変更される可能性があります。
実際のフィールドガイドはどのように見えるか
最近の例では、「テヘランの抗議者」というラベルが付いたウイルス画像が、研究者が AI によって生成されたものであるとフラグを立てる前に、X、テレグラム、レディット上で拡散しました。画像には GPS データ、カメラのメーカー、レンズ、焦点距離がありませんでした。非

携帯電話からの実際の抗議写真に含まれる日常的なメタデータの一部。ソフトウェアフィールドは空でした。その XMP パケットには、trainedAlgorithmicMedia の DigitalSourceType が含まれていました。モデルが画像を生成し、その出力に自発的にタグを付けていましたが、そのファイルがオープン Web に到達した時点でも、そのタグはそのままの状態でした。
ただし、これは常に起こるわけではありません。たとえば、国防総省の爆発のフェイクは、メタデータ内でそれ自体をタグ付けしていない前世代のツールからのものです。教皇のイメージは、C2PA が広く展開される以前から存在していました。エリオット・ヒギンズのトランプ逮捕デモは、DALL-E 3 のコンテンツ認証よりも前に行われていました。メタデータ フィールド ガイドは、ジェネレーターが標準に参加している場合、誰もファイルを手動でスクラブしていない場合、および表示しているファイルが拡張メタデータを削除するために再エンコードされていない場合に機能します。
それはたくさんの条件です。それでも、予想よりも頻繁に彼らに会います。
iPhone でこれらのフィールドを読み取る方法
疑わしい画像を iPhone の写真アプリまたはファイル アプリに保存します。
Photo Investigator を開き、写真またはビデオを選択します。
画像に青い AI バッジがオーバーレイされているかどうかを確認します。表示された場合は、AI メタデータが検出されました。それをタップすると、アプリが何を見つけたかを正確に確認できます。
完全なフィールドのリストを表示するには、「メタデータ」、「すべて表示」の順にタップします。ファイルに埋め込まれているすべてのタグが表示されます。
実際の iPhone の写真には、約 40 個の EXIF フィールド、位置情報を有効にしている場合は GPS 座標、Apple のメーカーノート データ、ポートレートの場合は深度マップ情報が表示されますが、C2PA マニフェストはありません。対照的に、DALL-E 3 の画像には、OpenAI を引用する少数のフィールドと C2PA ブロックのみが表示されます。 Midjourney の出力には、モデルの名前を示す Software フィールドが表示されます。違いは明らかです。
これらはいずれもクラウドでは実行されません。画像は携帯電話から残りません。写真調査員が私を読む

ローカルでタデータ。
このアプローチが何が良いのか、何が良くないのか
全体として、これは迅速な最初のチェックです。あなたがすでに持っているデバイス上で、悪意のある者が最も頻繁に削除し忘れるフィールドに対して、数秒で機能します。今日出回っている AI の誤った情報のほとんどを占める怠惰な偽物を捕まえるだけで十分です。
これはディープフェイク検出器ではありません。ピクセル分析はメソッドの一部ではありません (ただし、近日中に提供される予定です)。このチェックでは、誰かが実際の写真を欺瞞的に編集したかどうかを判断することはできません。以下では動作しません:
AI 画像のスクリーンショット。スクリーンショット パイプラインで拡張メタデータが失われることがよくあります。
ソーシャル プラットフォームによって再エンコードされた画像 (ほとんどのプラットフォームはアップロード時にメタデータを削除しますが、元のファイルを入手できれば、多くの場合タグがまだ残っています)
それでも、信頼できる偽物を製造するコストが上昇することになります。 AI が生成した「抗議写真」をプッシュしたい悪意のある攻撃者は、タグのないジェネレーターを使用するか、メタデータを手動でスクラブするか (Photo Investigator を使用して行うこともできます)、あるいは誰もメタデータ リーダーを介してファイルを実行しないことを祈る必要があります。この追加のステップが、デフォルトで普及するものと、意図的な努力が必要なものとの違いになります。
配布のスピードが検証のスピードを常に上回っている誤った情報の環境では、制作側のコストを上げることが重要です。
政治的利害に関するメモ
フグを着た教皇の話はインターネット上の楽しい瞬間です。国防総省の爆発事件の偽物が株式市場を揺るがした。これらのトランプ逮捕の画像は、明らかに事実として広まったデモであると明記されていました。これらはいずれも選挙を決定づけるものではなかった。まだ。
2024 年の米国選挙では、AI が生成した偽の支持表明、偽の集会群衆、偽の討論シーンの画像が世界中のアカウントから出回りました。

政治的スペクトル。 2026 年の中間選挙ではさらに多くのことが起こるでしょう。インド、アルゼンチン、スロバキア、インドネシアなど、他の国の選挙でもすでにこうした事態が見られています。すべてのキャンペーン サイクルは、ダーティ トリック ツールキットを含むツールキットの一部として生成 AI が組み込まれる環境になっています。
結局のところ、拡散するスピードで画像を検証することは党派的な問題ではありません。携帯電話を持っている人なら誰でも、ウイルス画像に生成者の署名が含まれているかどうかを確認できるツールでは、問題は解決されません。しかし、それらは安価な偽物が上陸するのを難しくしており、安価な偽物がほとんどです。
ソーシャル メディアで画像を共有すると、AI メタデータは削除されますか?
通常はそうです。 Instagram、X/Twitter、Facebook、TikTok はすべて、アップロード時に拡張メタデータを削除し、XMP フィールドと C2PA フィールドを削除します。誰かが共有する前に元のファイルを取得できた場合 (ジェネレーターから直接ダウンロードしたり、AirDrop 経由で送信したり、ソースからプルしたり)、通常、タグはそのままです。すでに投稿されている画像のスクリーンショットですか？メタデータがなくなったと仮定します。
AI メタデータが見つからない場合は何を意味しますか?
簡単に言えば、それは未知であり、本物ではないことを意味します。メタデータが削除されても、写真が本物であるという証拠にはなりません。古いジェネレーター、再エンコードされたファイル、意図的にスクラブされた画像はすべてクリーンになります。メタデータは迅速な最初のチェックです。メタデータが肯定的であれば、それは決定的です。陰性の場合は、さらに調査が必要です。
Photo Investigator は、C2PA マニフェスト、XMP DigitalSourceType 、EXIF ソフトウェア署名、および iPhone、iPad、または Mac に保存できる写真やビデオの完全なメタデータ プロファイルを読み取ります。クラウドアップロードはありません。追跡はありません。ファイルはデバイス上に残ります。
ハイキングの写真から GPS を削除してください。群衆があなたを求めています
2021年1月4日 2026年5月10日
写真調査官
写真およびビデオのメタデータ用のテンプレートを使用してメタデータを一括編集
あ

2023年4月15日 2023年4月15日
写真調査官
ポストナビゲーション
前 iPhone で「未編集」エプスタイン動画の編集内容を公開する
カテゴリー
ハウツー

## Original Extract

AI image detectors get fooled. Metadata does not. Here is how to check whether a photo was generated by DALL-E, Midjourney, or Adobe Firefly using the signals the generators themselves embed.

Is it AI? How to Tell Using Metadata - The Photo Investigator
Skip to content
The Photo Investigator
Explorations with Photo Metadata
Is it AI? How to Tell Using Metadata
In March 2023, Eliot Higgins, the founder of the open-source investigations outlet Bellingcat, fed a few prompts into Midjourney about Donald Trump’s arrest. He posted the resulting images publicly, labeled them clearly as AI-generated, and watched them go viral within hours as authentic news. Higgins said he made them to demonstrate how thin the line between real and fake had become. He was right. The labels he attached to the images didn’t survive the first screenshot.
Two months later, a checkmark-verified Twitter account impersonating Bloomberg News posted what looked like a photograph of an explosion at the Pentagon. The S&P 500 dropped 0.3% in five minutes before traders realized the image was AI-generated. The market recovered within the hour, but the obvious fake had real effects on the market.
These incidents, along with the Pope in the white Balenciaga puffer coat, the AI-generated Maui wildfire “evidence” photos, the fake Trump-Harris rally images from the 2024 election cycle, have all played out the same way. Image goes viral. Investigation follows. Truth catches up late, if at all. By then, the audience that wanted to believe has moved on, but in many cases checking the photo metadata with the Photo Investigator can reveal the truth.
The metadata most AI generators leave behind
In fact, a surprising number of AI image generators voluntarily label their output. Not in pixels, but in metadata fields the generators write directly into the file when they save it.
The Coalition for Content Provenance and Authenticity, an industry consortium that includes Adobe, Microsoft, OpenAI, Google, and major camera manufacturers, defined a standard for content credentials. A C2PA manifest is a cryptographically signed packet embedded in the image file. It records who or what created the image, which edits the creator applied, and which software produced the file. OpenAI added C2PA to DALL-E 3 outputs in early 2024. Adobe Firefly tags every image it produces. Google ImageFX embeds the same standard. Microsoft Designer too. If an image has a valid C2PA manifest, you can read exactly which model produced it and when.
The IPTC photo metadata standard includes a specific XMP field called “ Iptc4xmpExt:DigitalSourceType “, designed to declare how a digital file originated. When this field says "trainedAlgorithmicMedia “, the file came from a generative model. “ digitalCapture ” means a camera captured it. AI generators that follow the standard write this field on export.
Finally, this is the oldest and crudest signal. Real cameras write their make and model into the EXIF Software field. AI generators sometimes do too. A Software field that reads “Midjourney v6.1” or “Stable Diffusion XL” is doing your detection work for you. Even when the generator itself doesn’t tag, the Python libraries used to save the output sometimes do: entries like PIL/Pillow or OpenCV showing up on what’s supposed to be an iPhone snapshot is a signal.
Which generators tag their output (2026)
Reflects confirmed behavior as of June 2026. Generator behavior can change with software updates.
What the field guide looks like in practice
In a recent example, a viral image labeled “protesters in Tehran” circulated across X, Telegram, and Reddit before researchers flagged it as AI-generated. The image had no GPS data, no camera make, no lens, no focal length. None of the routine metadata a real protest photo from a phone would carry. Its Software field was empty. Its XMP packet contained a DigitalSourceType of trainedAlgorithmicMedia . A model had generated the image and voluntarily tagged its output, and that tag was still intact when the file reached the open web.
However, this doesn’t always happen. The Pentagon explosion fake, for example, came from a tool of an earlier generation that didn’t tag itself in the metadata. The Pope-in-puffer image predates widespread C2PA deployment. Eliot Higgins’ Trump arrest demonstration predates DALL-E 3’s content credentials. The metadata field guide works when the generator participates in the standard, when nobody has manually scrubbed the file, and when the file you’re looking at hasn’t been re-encoded to strip extended metadata.
That’s a lot of conditions. Still, they’re met more often than you’d expect.
How to read these fields on iPhone
Save the suspected image to your iPhone’s Photos app or Files app.
Open Photo Investigator and select the photo or video.
Check for a blue AI badge overlaid on the image. If it appears, AI metadata was detected. Tap it to see exactly what the app found.
For the full field listing, tap Metadata , then View All . Every tag embedded in the file will be shown.
On a real iPhone photo, you’ll see roughly 40 EXIF fields, GPS coordinates if you enabled location, Apple’s maker-note data, depth map info if it’s a portrait, and no C2PA manifest. A DALL-E 3 image, by contrast, shows just a handful of fields and a C2PA block citing OpenAI. On a Midjourney output, you’ll see the Software field name the model. The difference is obvious.
None of this runs in the cloud. The image doesn’t leave your phone. Photo Investigator reads the metadata locally.
What this approach is good for, and what it isn’t
Overall, this is a fast first check. It works in seconds, on the device you’re already holding, against the fields that bad actors most often forget to strip. It’s enough to catch the lazy fakes that account for most of the AI misinformation circulating today.
This is not a deepfake detector. Pixel analysis is not part of the method (though it is coming soon). The check cannot determine whether someone has deceptively edited a real photograph. It does not work on:
Screenshots of AI images, which often lose extended metadata in the screenshot pipeline
Images re-encoded by social platforms (most platforms strip metadata on upload, but the original file, if you can get it, often still has the tags)
Even so, what it does is raise the cost of producing a credible fake. A bad actor who wants to push an AI-generated “protest photo” now has to either use a generator that doesn’t tag, manually scrub the metadata (which could also be done with the Photo Investigator ), or hope no one runs the file through a metadata reader. That extra step is the difference between something that spreads by default and something that requires deliberate effort.
In a misinformation environment where the speed of distribution has always outpaced the speed of verification, raising the cost on the production side matters.
A note on the political stakes
The Pope-in-puffer story is a fun internet moment. The Pentagon-explosion fake nudged the stock market. Those Trump arrest images were a clearly labeled demonstration that went viral as fact. None of these were election-defining. Yet.
The 2024 US election saw AI-generated images of fake endorsements, fake rally crowds, and fake debate scenes circulate from accounts across the political spectrum. The 2026 midterms will see more. Other countries’ elections are seeing them already: India, Argentina, Slovakia, Indonesia. Every campaign cycle is now an environment in which generative AI is part of the toolkit, including the dirty tricks toolkit.
Ultimately, verifying images at the speed they spread isn’t a partisan problem. Tools that let any person with a phone check whether a viral image carries a generator’s signature won’t fix the problem. But they make it harder for the cheap fakes to land, and the cheap fakes are most of them.
Does sharing an image on social media remove AI metadata?
Usually yes. Instagram, X/Twitter, Facebook, and TikTok all strip extended metadata on upload, which removes XMP and C2PA fields. If you can get the original file before someone shared it (downloaded directly from the generator, sent via AirDrop, or pulled from the source), the tags are typically still intact. A screenshot of an image that was already posted? Assume the metadata is gone.
What does it mean if no AI metadata is found?
Simply put, it means unknown, not authentic. Stripped metadata is not proof that a photo is real. Older generators, re-encoded files, and intentionally scrubbed images will all come up clean. Metadata is a fast first check: when it’s positive, it’s definitive. When it’s negative, you need more investigation.
Photo Investigator reads C2PA manifests, XMP DigitalSourceType , EXIF Software signatures, and the full metadata profile of any photo or video you can save to your iPhone, iPad, or Mac. No cloud upload. No tracking. The file stays on your device.
Remove GPS From Hiking Photos, Crowds are Begging You
January 4, 2021 May 10, 2026
The Photo Investigator
Bulk Edit Metadata with Templates for Photo and Video Metadata
April 15, 2023 April 15, 2023
The Photo Investigator
Post navigation
Previous Expose the “unedited” Epstein video’s edits on your iPhone
Categories
How To
