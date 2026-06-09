---
source: "https://spectrum.ieee.org/tracking-glacier-melting-ai"
hn_url: "https://news.ycombinator.com/item?id=48460778"
title: "AI Can Help Track the Shrinking Glaciers"
article_title: "AI Could Increase Remote Monitoring of Glacier Melting - IEEE Spectrum"
author: "berlianta"
captured_at: "2026-06-09T16:07:05Z"
capture_tool: "hn-digest"
hn_id: 48460778
score: 2
comments: 0
posted_at: "2026-06-09T13:19:26Z"
tags:
  - hacker-news
  - translated
---

# AI Can Help Track the Shrinking Glaciers

- HN: [48460778](https://news.ycombinator.com/item?id=48460778)
- Source: [spectrum.ieee.org](https://spectrum.ieee.org/tracking-glacier-melting-ai)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T13:19:26Z

## Translation

タイトル: AI は縮小する氷河の追跡に役立つ
記事のタイトル: AI は氷河融解の遠隔監視を強化できる - IEEE Spectrum
説明: よりスマートな衛星画像分析により、氷河融解の正確な AI モデルが提供され、科学者が世界中の氷河をより簡単に追跡できるようになります。

記事本文:
-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
AI により氷河融解の遠隔監視が強化される可能性 - IEEE Spectrum
IEEE.org IEEE Explore IEEE 標準 IEEE 求人サイト その他のサイト サイン イン 参加 IEEE AI は世界の縮小しつつある氷河の追跡に役立つ
IEEEスペクトル
テクノロジー関係者向けトピック
アカウントを作成すると、さらに無料のコンテンツや特典をお楽しみいただけます
後で読むために記事を保存するには、IEEE Spectrum アカウントが必要です
インスティチュートのコンテンツはメンバーのみが利用できます
PDF 版全号のダウンロードは IEEE メンバー限定です
この電子ブックのダウンロードは IEEE メンバー限定です
へのアクセス
スペクトル
のデジタル版は IEEE メンバー限定です
以下のトピックは IEEE メンバー限定の機能です
記事に回答を追加するには、IEEE Spectrum アカウントが必要です
アカウントを作成すると、さらに多くのコンテンツや機能にアクセスできます
IEEEスペクトル
、後で読むために記事を保存したり、スペクトル コレクションをダウンロードしたり、イベントに参加したりする機能が含まれます。
読者や編集者との会話。より独占的なコンテンツと機能については、次のことを検討してください。
IEEEに参加する
。
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
Spectrum のすべての記事、アーカイブ、PDF ダウンロード、その他の特典を利用できます。
IEEE について詳しくはこちら →
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
この電子書籍に加えて、
IEEE スペクトラム
記事、アーカイブ、PDF ダウンロード、その他の特典。
IEEE について詳しくはこちら →
何千ものアクセス

の記事 — 完全に無料
アカウントを作成して限定コンテンツと機能を入手してください:
記事の保存、コレクションのダウンロード、
そして
コメントを投稿する
— すべて無料です！フルアクセスと特典については、
購読する
スペクトルに。
AI は世界の縮小する氷河の追跡に貢献できる
新しいアプローチにより、トップ AI 氷河追跡モデルが新しい地域に簡単に適応できるようになります
Edd Gent は、IEEE Spectrum の寄稿編集者です。
この衛星画像では、覗いている黄色の線が、予想される氷河の境界 (明るい灰色) を示しています。グラウンド トゥルース フロントは青で描画され、それらの間の重なり部分はピンクで表示されます。
氷河がどのくらいの速度で縮小しているかを追跡することは、気候変動のペースを測定し、将来の海面上昇を予測するために非常に重要です。これは通常、骨の折れる手作業ですが、AI が世界中のどこにいても氷河の衛星画像を分析できるようにする新しいアプローチは、監視プロセスの自動化に役立つ可能性があります。
海に直接流れ込む氷河は地球の気候に重要な役割を果たしていますが、地球温暖化により氷河の後退がますます加速しています。これは深刻な波及効果をもたらす可能性があり、「分断前線」（氷山が水面に切り裂かれる氷河の端）から剥がれた氷が大量の淡水を海に放出し、海流を変え、海面上昇を引き起こす可能性がある。明るく白い氷河も太陽光を多く反射します。収縮すると、太陽からの熱を吸収する暗い海水にさらされます。
これらすべては、地域と世界の両方の気候条件が時間の経過とともにどのように変化するかを理解するために、氷河の消失を追跡することが重要であることを意味します。しかし、世界中で監視する必要がある氷河の数は人間の分析能力をはるかに超えています。 AI ベースの画像分析がギャップを埋めるのに役立つことが期待されていますが、以前のモデルは十分な性能を発揮していませんでした。

トレーニング データに含まれていない領域についてはあまり効果がありませんでした。手動でラベル付けされた画像を収集することがいかに難しいかを考えると、このアプローチの適用可能性は大幅に制限されます。
今回、IEEE 国際画像処理会議 (ICIP) に採択された論文は、氷河の融解フロントを追跡するための主要な深層学習モデルが、最小限の追加データで新しい場所に適応できることを示しています。ドイツのエアランゲン・ニュルンベルク・フリードリヒ・アレクサンダー大学（FAU）の研究者らは、氷河ごとに手作業でラベル付けされた1枚の画像、ラベル付けされていない夏の参照画像、および下にある岩の地図という3つの情報を提供することによって、モデルの誤差（モデル化された境界線と実際の境界線の間の平均距離）が1キロメートル以上から70メートル未満に短縮されたことを示した。
関連研究では、論文の著者の何人かがすでにこのアプローチを実践しており、2015年から2024年までノルウェーのスヴァールバル諸島にある145の氷河すべての月ごとの分氷前線の位置を抽出しています。研究チームは現在、このアプローチを北極のさらに1,500の氷河に拡張したいと考えています。
「それは氷河をより深く理解し、氷河が気候の変化にどのように反応するかを理解することです」と博士のノラ・グルメロン氏は言います。 FAU の学生であり、ICIP 論文の共同筆頭著者。 「過去について知れば、将来どのように変化するのかもよりよく理解できるようになります。」
歴史的に、分娩前線の輪郭を描くには、学生や研究者が衛星レーダー画像を注ぎ込んで氷河と海の境界を手動で追跡する必要があった、とグルメロン氏は言う。ただし、このプロセスには時間がかかるため、多くの研究グループがコンピュータ ビジョン モデルを使用してプロセスを自動化する実験を行っています。
2023 年、グルメロンとその同僚は

d 南極、グリーンランド、アラスカの 7 つの氷河の 681 枚のレーダー画像のデータセット。新しいモデルのトレーニングとベンチマークに役立つように、手動で注釈が付けられた分断前線が含まれています。しかし、このデータセットでトレーニングされた最先端の深層学習モデルを取得し、それをスバールバル諸島のこれまで見たことのない氷河に適用したところ、平均で 1,131.6 m ずれていることがわかりました。
分析したいすべての新しい氷河についてモデルを再トレーニングするのに十分な量の手動で注釈を付けたデータを収集することは明らかに実行不可能であるため、著者らはパフォーマンスを向上させるより効率的な方法を見つけようとしました。彼らは、スバールバル諸島の 145 個の氷河すべてに対して手動で注釈を付けた 1 つの分断前線画像を作成し、これを各氷河のさらにいくつかの生の衛星画像と組み合わせて、5,539 枚の画像の新しいトレーニング セットを作成しました。この新しいデータと元のベンチマーク データの両方でモデルを再トレーニングしたところ、誤差は 445.3 m まで低下しました。
その後、精度をさらに向上させるための 2 つの新しい戦略を開発しました。人間にとっても AI にとっても、氷河の境界と氷のメランジ (分断前線に蓄積する可能性のある浮遊氷山、海氷、雪がどろどろになったもの) とを区別するのは難しい場合があります。そこで、研究者らがモデルに注釈を付けるために氷河の一連の画像をアップロードしたとき、メランジが存在せず、氷河の境界がはっきりしている夏の画像 3 枚を含めました。これらはモデルの基準点として機能し、誤差を 204.6 m まで押し下げました。
最終ステップとして、研究者らは、スバールバル諸島の海岸の概要を示すオープン ストリート マップ データから得られた、各氷河の下にある岩の静的地図をモデルに提供しました。これにより、誤差はわずか 103.6 m に短縮されました。研究者らは、モデルの 5 つの異なるバージョンのアンサンブルを実行し、その出力を平均することで、フィンを得ることができました。

誤差はわずか 68.7 m にまで下がりました。これでもかなり不正確に聞こえるかもしれませんが、Gourmelon 氏は、これは手動による注釈のエラー率に匹敵すると言います。 「特に氷のメランジがある場合や衛星画像の解像度がそれほど良くない場合、人間自身がラベル付けを一貫して行うことができません」と彼女は言います。
このアプローチにはまだある程度の作業が必要ですが、新しい領域の分析を劇的にスピードアップできます。この種の最近の研究のほとんどは、年単位または十年単位のタイムスケールで行われている、とダコタ・パイルズ博士は言います。 FAU の学生で、スバールバル諸島の 9 年分の氷河の動態をマッピングした 2 番目の研究を主導しました。低周波数の追跡とは対照的に、パイルズ氏はその研究ですべての氷河について月ごとの分断前線を生成することができ、合計 203,294 を超える注釈があり、列島で氷の動態がどのように変化しているかをより詳細に把握できるようになりました。
「モデルがなかったら、私のプロジェクトは、私たちが目指している規模では不可能でした」とパイルズ氏は言います。 「つまり、それは私たちにとって、そして一般に氷河学の分野の進歩にとって大きな利益です。」
長期的には、このアプローチにより、長期間にわたる世界中の氷河の監視を部分的に自動化できる可能性があるとグルメロン氏は述べています。 「最初にトレーニングするために、監視に使用したい特定の地域または衛星からのラベル付き画像がまだ必要ですが、その後は使用できるようになります」と彼女は言います。 「画像のキャプチャ方法とどこを見ているのかが一貫していれば、再調整は必要ありません。」
氷レーダードローンで気候変動を研究 ›
粘弾性モデリングによるグリーンランド氷河の氷の消失を予測 ›
ESA - 氷河の融解により淡水の損失が増大し、海面上昇が加速 ›
アイスメルト |

世界の海面 – NASA 海面変化ポータル ›
Edd Gent は、インドのバンガロールに拠点を置くフリーランスの科学技術ライターです。彼の著作は、コンピューティング、エンジニアリング、エネルギー、生物科学にわたる新興テクノロジーに焦点を当てています。彼は Twitter で @EddytheGent を使用し、メールで edd dot gent at Outlook dot com を使用しています。彼の PGP フィンガープリントは ABB8 6BB3 3E69 C4A7 EC91 611B 5C12 193D 5DFC C01B です。彼の公開鍵はここにあります。シグナルに関する情報はDMで。
JPL は 13 歳の好奇心探査機をどのようにして科学活動を続けているのか
器用さを超えて: なぜ接触がロボット工学の次の時代を定義するのか
サルデーニャがクリーンエネルギーの未来を拒否する古代の理由
Giga は世界中の学校を見つけて接続することを目指しています
AI が単一のレーダー画像から 3D 都市地図を生成
衛星レーダーがウクライナダム決壊の疑問を鮮明にする

## Original Extract

Smarter satellite image analysis delivers accurate AI models of glacier melting , empowering scientists to more easily track glaciers around the world.

-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
AI Could Increase Remote Monitoring of Glacier Melting - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE AI Can Help Track the World’s Shrinking Glaciers Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
IEEE Spectrum
FOR THE TECHNOLOGY INSIDER Topics
Enjoy more free content and benefits by creating an account
Saving articles to read later requires an IEEE Spectrum account
The Institute content is only available for members
Downloading full PDF issues is exclusive for IEEE Members
Downloading this e-book is exclusive for IEEE Members
Access to
Spectrum
's Digital Edition is exclusive for IEEE Members
Following topics is a feature exclusive for IEEE Members
Adding your response to an article requires an IEEE Spectrum account
Create an account to access more content and features on
IEEE Spectrum
, including the ability to save articles to read later, download Spectrum Collections, and participate in
conversations with readers and editors. For more exclusive content and features, consider
Joining IEEE
.
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
all of Spectrum’s articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
this e-book plus all of
IEEE Spectrum’s
articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Access Thousands of Articles — Completely Free
Create an account and get exclusive content and features:
Save articles, download collections,
and
post comments
— all free! For full access and benefits,
subscribe
to Spectrum .
AI Can Help Track the World’s Shrinking Glaciers
New approach lets top AI glacier tracking model easily adapt to new regions
Edd Gent is a contributing editor for IEEE Spectrum.
In this satellite image, the yellow line peeking through shows the predicted boundary of the glacier (light gray). The ground truth front is drawn in blue, and the overlap between them is pink.
Tracking how fast glaciers are shrinking is crucial for measuring the pace of climate change and projecting future sea level rises. This is normally a painstaking manual job, but a new approach that enables AI to analyze satellite images of glaciers anywhere in the world could help automate the monitoring process.
Glaciers that flow directly into the ocean play a crucial role in the earth’s climate, but global warming is making them retreat ever faster. This can have severe knock-on effects as ice that breaks away from “calving fronts”—the ends of glaciers where icebergs shear off into the water—dumps massive amounts of freshwater into the sea, which can alter ocean currents and cause sea levels to rise. Bright white glaciers also reflect a lot of sunlight. When they shrink, they expose dark seawater that absorbs heat from the sun.
All of this means that tracking glacier loss is critical for understanding how both local and global climate conditions will change over time. But the number of glaciers that need to be monitored around the world far outstrips the capacity of human analysts. There is hope that AI-based image analysis could help plug the gap, but previous models have performed poorly on regions not included in their training data. This severely limits the applicability of the approach, given how difficult it is to collect manually-labeled images.
Now, a paper accepted to IEEE International Conference on Image Processing (ICIP) shows that a leading deep learning model for tracing glacier calving fronts can be adapted to new locations with minimal additional data. Researchers from the Friedrich-Alexander University of Erlangen–Nuremberg (FAU), in Germany , showed that the model’s error—the average distance between the modeled boundary and the real one—was cut from more than a kilometer to less than 70 meters by providing three pieces of information: one hand-labeled image per glacier, un-labelled summer reference images, and a map of the underlying rock.
In related research , some of the paper’s authors have already put the approach to work, using it to extract monthly calving front positions for all 145 glaciers in Norway’s Svalbard archipelago from 2015 to 2024. The team now hopes to extend the approach to another 1,500 glaciers in the Arctic .
“It’s about understanding glaciers better, and how they react to changes in the climate,” says Nora Gourmelon , a Ph.D. student at FAU and co-lead author of the ICIP paper. “When you know about the past then you will also hopefully be better able to understand how they will change in the future.”
Historically, delineating calving fronts has required students and researchers to pour over satellite radar images to manually trace the boundary between glaciers and the ocean, says Gourmelon. The process is time-consuming though, so numerous research groups have been experimenting with using computer vision models to automate the process.
In 2023, Gourmelon and her colleagues produced a dataset of 681 radar images of seven glaciers in Antarctica , Greenland , and Alaska , with manually annotated calving fronts to help train and benchmark new models. But when they took a state-of-the-art deep learning model trained on this dataset and applied it to previously unseen glaciers in Svalbard they found it was off by an average of 1,131.6 m.
Gathering enough manually annotated data to retrain a model on every new glacier you want to analyze would clearly be infeasible, so the authors tried to find a more efficient way to boost performance. They produced one manually annotated calving front image for all 145 glaciers in Svalbard and combined this with several more raw satellite images of each glacier to create a new training set of 5,539 images. When they retrained the model on both this new data and the original benchmark data, the error dropped to 445.3 m.
They then developed two novel strategies to further improve accuracy. For both humans and AI it can be tricky to distinguish the boundary of a glacier from ice melange—the mush of floating icebergs, sea ice, and snow that can accumulate at the calving front. So when the researchers uploaded a series of images of a glacier for the model to annotate, they included three images from the summer, when the melange is not present and the boundary of the glacier is clear. These acted as a reference point for the model and pushed the error down to 204.6 m.
As a final step, the researchers also provided the model with a static map of the rock underlying each glacier, derived from Open Street Map data that outlines the coast of Svalbard. This slashed the error to just 103.6 m. By running an ensemble of five different versions of their model and averaging their outputs, the researchers were able to get their final error down to just 68.7 m. While that may still sound fairly imprecise, Gourmelon says it’s comparable to manual annotation error rates. “Humans themselves are not really consistent in labeling, especially when there’s ice melange or when the resolution of the satellite image is not that good,” she says.
While the approach still requires some leg work, it can dramatically speed up the analysis of new regions. Most recent research of this kind has been done on annual or decadal timescales, says Dakota Pyles , a Ph.D. student at FAU, who led the second study that mapped nine years worth of glacier dynamics in Svalbard. In contrast to the lower frequency tracking, Pyles was able to generate monthly calving fronts for every glacier in that study—a total of more than 203,294 annotations—providing a much finer-grained view of how ice dynamics are changing on the archipelago.
“My project would not be possible at the scale that we’re going for if we didn’t have the model,” says Pyles. “So that is a great benefit for us, and generally for advancing the field of glaciology.”
In the long run, the approach could make it possible to partially automate the monitoring of glaciers around the world over extended periods of time, says Gourmelon. “We still need some labeled images from the specific region or satellite that you want to use for monitoring to train it on first, but then it can be used,” she says. “If how the image is captured stays consistent, and where you’re looking at, then no recalibration would be necessary.”
Studying Climate Change With an Ice Radar Drone ›
Forecasting the Ice Loss of Greenland’s Glaciers With Viscoelastic Modeling ›
ESA - Glacier melt intensifying freshwater loss and accelerating sea-level rise ›
Ice Melt | Global Sea Level – NASA Sea Level Change Portal ›
Edd Gent is a freelance science and technology writer based in Bengaluru, India. His writing focuses on emerging technologies across computing, engineering, energy and bioscience. He's on Twitter at @EddytheGent and email at edd dot gent at outlook dot com. His PGP fingerprint is ABB8 6BB3 3E69 C4A7 EC91 611B 5C12 193D 5DFC C01B. His public key is here. DM for Signal info.
How JPL Keeps the 13-Year-Old Curiosity Rover Doing Science
Beyond Dexterity: Why Contact May Define the Next Era of Robotics
Sardinia’s Ancient Reasons for Rejecting a Clean Energy Future
Giga Aims to Find and Connect the World’s Schools
AI Generates 3D City Maps From Single Radar Images
Satellite Radar Sharpens Ukraine Dam Collapse Questions
