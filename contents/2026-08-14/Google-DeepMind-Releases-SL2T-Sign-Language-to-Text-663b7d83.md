---
source: "https://deepmind.google/blog/putting-sign-language-ai-into-users-hands/"
hn_url: "https://news.ycombinator.com/item?id=49302874"
title: "Google DeepMind Releases SL2T (Sign Language to Text)"
article_title: "Putting sign language AI into users’ hands — Google DeepMind"
author: "neelbuilds"
captured_at: "2026-08-14T18:40:44Z"
capture_tool: "hn-digest"
hn_id: 49302874
score: 1
comments: 0
posted_at: "2026-08-14T18:38:01Z"
tags:
  - hacker-news
  - translated
---

# Google DeepMind Releases SL2T (Sign Language to Text)

- HN: [49302874](https://news.ycombinator.com/item?id=49302874)
- Source: [deepmind.google](https://deepmind.google/blog/putting-sign-language-ai-into-users-hands/)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T18:38:01Z

## Translation

タイトル: Google DeepMind が SL2T (手話からテキストへ) をリリース
記事タイトル: 手話 AI をユーザーの手に — Google DeepMind

記事本文:
メイン コンテンツにスキップ 次世代 AI システムを探索する
AI の最新のブレークスルーとラボからの最新情報
AI で新たな発見の時代を切り開く
私たちの使命は、人類に利益をもたらすために責任を持って AI を構築することです
次世代 AI システムを探索する
AI の最新のブレークスルーとラボからの最新情報
AI で新たな発見の時代を切り開く
私たちの使命は、人類に利益をもたらすために責任を持って AI を構築することです
Google DeepMind Google AI すべての AI について学ぶ Google DeepMind AI の最前線を探る Google Labs AI 実験を試す Google Research 研究を探索する 製品とアプリ Gemini アプリ Gemini とチャットする Google AI Studio 次世代 AI モデルで構築する Google Antigravity 当社のエージェント開発プラットフォーム モデル 研究科学 Gemini について Build Gemini を試す 2026 年 8 月 12 日 モデル 手話 AI をユーザーの手に
Google DeepMind 手話チーム
シェア コピーしました 手話からテキストへ (SL2T) を導入します。これは、聴覚障害者および難聴ユーザーのための新しい手話機能を強化する画期的なモデルです。
AI の話し言葉を処理する能力はここ数十年で急速に進歩し、自動翻訳、ディクテーション、聴覚ユーザーにとって楽に感じられる会話インターフェイスが可能になりました。しかし、この技術革命は、世界の 200 以上の手話と、それを使用する推定 7,000 万人のろう者および難聴者にはまだ届いていません。
本日、私たちは品質と汎用性において画期的な、大規模な多言語の手話からテキストへの (SL2T) 翻訳モデルを導入します。これにより、私たちは手話 AI を初めて研究室から消費者製品に導入します。SL2T は、アメリカ手話 (ASL) から英語まで、Gboard での手話からテキストへのディクテーションと Pixel 11 でのライブ文字起こしを強化します。

今後さらに多くのデバイスが登場し、言語も追加される予定です。
聴者が入力する代わりにディクテーションを使用して話すことができるのと同様に、この機能を使用すると、聴覚障害者ユーザーは、通常入力している場所であればどこでも携帯電話にサインできるようになります。署名して Web を検索したり、メッセージやドキュメントを下書きしたり、Gemini にクエリの解決やタスクの実行を依頼したりできます。 Live Transcribe では、会話の中で何度も入力する代わりに、応答に署名できます。テスターに​​よると、ASL でのサインインは英語で入力するよりも速く、より自然で、楽しいとのことです。
SL2T を利用した Sign-to-Text により、ユーザーは通常入力している場所ならどこでも携帯電話に署名できます。
手話は世界中のろう者コミュニティの主要言語であり、ろう者の文化的アイデンティティの基礎です。手話、会話、読み書きの習熟度の点で聴覚障害者には大きな多様性があるため、あらゆる手段でのアクセスをサポートすることが重要です。聴覚障害者は、健聴者が音声言語処理から恩恵を受けるのと同じように、手話処理から恩恵を受けることができ、さらにこのテクノロジーは、聴覚障害者コミュニティと聴者コミュニティの間のコミュニケーションギャップを埋める新たな可能性をもたらします。社会にプラスの影響を与えるこの機会にもかかわらず、手話 AI の進歩は遅れています。手話用 AI の構築には複雑な課題があることと、手話言語自体がどのように機能するかについて誤解が広まっているためです。
音声言語の転写と比較すると、手話翻訳には 2 つの主要な課題があります。まず、音声の転写は、同じ言語で音からテキストへの順次マッピングを実行することですが、一方、手話は、独自の明確な文法と語彙を持つ独立した自然言語です。結果として、

記号から単語への変換という一連のプロセスではなく、真の機械翻訳が必要です。第二に、モデルは物理的な動きを「見て」理解することを学ばなければなりません。手話は、手、腕、胴体、頭、顔の同時の動きを通じて意味を伝えます。これらを高フレーム レートで正確に追跡することは、困難で計算量の多いコンピュータ ビジョン タスクです。
このような背景を考慮すると、手話手袋のような手話技術の初期の試みが根本的に制限されていた理由は簡単に理解できます。手話は単に「手の英語」ではないからです。きめ細かい全身の動きの複雑な視覚認識と本格的な言語翻訳が必要です。 SL2T は両方を提供するように設計されています。
SL2T は、手話入力を署名者の身体上の点として認識し、ストリーミング テキスト出力に変換します。 FLEURS-ASL ベンチマークの例。
私たちは、ユーザー中心で文化に基づいたアプローチと大規模なデータ スケーリングを組み合わせて SL2T を構築しました。このモデルは、50 以上の手話にわたる 100,000 時間以上のデータでトレーニングされており、データの約 4 分の 1 が ASL です。多様な言語、方言、熟練度レベルで共同トレーニングを行うと、モデルは共通の基礎構造を学習し、実験では単一言語モデルを上回りました。
ユーザーのプライバシーを保護するために、SL2T は手話を生のカメラ フィードではなく、一連のポーズのランドマーク位置と見なします。オンデバイス モデル ( MediaPipe Holistic ) は署名者の点の位置を追跡し、これらの幾何学的座標のみが翻訳のためにサーバーに送信されるため、元のビデオはすぐに破棄できます。
SL2T は、幅の広い「グロス」として知られる中間注釈をバイパスして、この座標シーケンスをテキストに直接変換します。

手話翻訳に関する以前の研究で使用されました。グロスでは、非手動マーカーや空間構造など、手話の豊かで非線形な側面を捉えることができません。ランドマークから直接翻訳することで、人為的な語彙制限がなくなり、データに応じて翻訳品質を直接調整できるようになります。
SL2T は、ASL から英語への翻訳の品質を評価する FLEURS-ASL (sd-test) などの主要なベンチマークによると、これまでで最も有能な手話翻訳モデルです。 SL2T は、70 BLEURT という驚くべきゼロショット スコアを達成しており、これは以前に報告されたどのスコアよりも大幅に高くなります。しかし、学術的なベンチマークを最適化するだけでは、現実世界のアプリケーションでの使いやすさが保証されるわけではありません。そのため、ストリーミング遅延の最小化、署名以外の入力での幻覚の防止、左利きの署名者の 10% に対する公平性の確保、スマートフォンをもう一方の手に持ちながら使用する片手署名のパフォーマンスの向上などの実践的な問題に熱心に取り組みました。
FLEURS-ASL ベンチマークの例。 SL2T は、複雑な ASL を流暢な英語に正確に翻訳します。まれな標識、素早い指の綴り (「獲物」→「灰色」)、受動的な構文、分類子の描写 (「爪」を落とす)、および文脈のない時制 (「蹴り落とされる」→「開始」) には時折エラーが残ります。
私たちは、ろう者コミュニティのためだけではなく、ろう者コミュニティとともに構築することを信じています。聴覚障害者の視点は、聴覚障害者の Google 社員である Sam Sepah による概念化から、聴覚障害者パートナーとのデータ収集、聴覚障害者ユーザー調査での評価、聴覚障害者専門家によるテクノロジーの影響評価に至るまで、このプロジェクトのあらゆる段階を形作ってきました。
責任ある現実世界への展開を導くために、私たちは AI 手話諮問委員会 (AISLAC) を設立し、多くの世界的な聴覚障害者組織と対象分野の専門家を結集しました。を通して

参加型ガバナンス モデルであるため、当社のテクノロジーによって最も影響を受けるコミュニティは、開発の優先順位に直接影響を与えます。私たちは、Gboard と Live Transcribe での SL2T 1.0 のリリースに関する共同影響レポートを共同執筆し、テクノロジーの機能と現在の制限を透過的に詳細に説明しました。この共同アプローチは、すべての主要な手話リリースで継続する予定です。
SL2T は、学界と産業界にわたる数十年にわたる基礎研究に基づいて構築されていますが、ASL 入力をユーザーの携帯電話に提供することは始まりにすぎません。 Google の使命は、世界中の情報を整理し、世界中の人がアクセスして使えるようにすることです。ユニバーサルなアクセシビリティを実現するということは、話し言葉と書き言葉が完全に同等になることを意味します。私たちのチームは、このテクノロジーを追加の手話、手話生成、最先端の AI 機能に拡張することに取り組んでいます。私たちは、デジタル環境全体で手話によるアクセスを標準にするために、責任を持って進捗状況を共有することを楽しみにしています。
Gboard での SL2T と Live Transcribe は、まず Pixel 11 で体験できます。今後さらに多くのデバイスが追加される予定で、すべて追加費用なしで利用できます。
この作業は、Google DeepMind と Android のチームによって共同で行われました。 SL2T モデルを開発したコア チームは、Garrett Tanzer、Benoit Brard、Elizabeth Clark、Tim Dozat、Sebastian Ebert、Dan Garrette、Manfred Georg、Vicky Holgate、Shankar Kumar、Mohammad Saboorian、Miloš Stanojević、Meghumekar、John Wieting、Andy Zhang、Chris Dyer です。
モデルを Gboard と Live Transcribe に統合した Android チームは、Ausmus Chang、Sai Aditya Chitturu、Dayle Chiu、Anna Chou、Ajay Dudani、Angana Ghosh、Alex Huang、Joanne Kim、Ed Lee、Thomas Lin、James Su、Yanchao Su、Sharlene Yuan です。
Anelia Angelova 氏のさらなるサポートに感謝いたします。

ビシェク・バプナ、サラ・バッソン、グレン・キャメロン、スコット・クロウェル、トレバー・コーン、ノア・フィーデル、ズービン・ガーラマーニ、ライア・ハドセル、トム・ハドソン、アレクサンダー・ハウエルスレフ・ジェンセン、川上和也、ペイケ・リー、リアム・マカファティ、キャロライン・パントファル、アビナフ・パラシャール、クリストファー・パトノエ、ローラ・ライメル、サーガル・サヴラ、サムセパ、サド・スターナー、デイブ・ユーサス、ビアオ・チャン。
また、私たちのモデルの初期段階のテストに参加してくださった方々にも感謝します。
私は Google の利用規約に同意し、私の情報が Google のプライバシー ポリシーに従って使用されることを認めます。

## Original Extract

Skip to main content Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Google DeepMind Google AI Learn about all our AI Google DeepMind Explore the frontier of AI Google Labs Try our AI experiments Google Research Explore our research Products and apps Gemini app Chat with Gemini Google AI Studio Build with our next-gen AI models Google Antigravity Our agentic development platform Models Research Science About Build with Gemini Try Gemini August 12, 2026 Models Putting sign language AI into users’ hands
Google DeepMind Sign Language Team
Share Copied Introducing sign-language-to-text (SL2T), our breakthrough model powering new sign language features for Deaf and hard of hearing users.
AI's ability to process spoken languages has advanced rapidly over recent decades, enabling automatic translation, dictation, and conversational interfaces that feel effortless to hearing users. Yet this technological revolution has not reached the world’s more than 200 sign languages — and the estimated 70 million Deaf and hard of hearing people who use them.
Today, we’re introducing a massively multilingual sign-language-to-text (SL2T) translation model that marks a breakthrough in quality and generality. With it, we are bringing sign language AI out of the lab and into consumer products for the first time: SL2T powers sign-to-text dictation in Gboard and Live Transcribe on Pixel 11 , starting with American Sign Language (ASL) to English. More devices are coming soon, and additional languages will follow.
Similarly to how hearing users can use dictation to speak instead of typing, this feature enables Deaf users to sign to their phone anywhere they’d normally type. You can sign to search the web, draft messages or documents, and ask Gemini to solve queries or execute tasks. In Live Transcribe, you can sign responses in conversations instead of having to type back and forth. According to our testers, signing in ASL is faster, more natural, and more delightful than typing in English.
Sign-to-text, powered by SL2T, enables users to sign to their phone anywhere they'd normally type.
Sign languages are the primary languages of Deaf communities around the world and the cornerstone of Deaf cultural identity. There is great diversity among deaf people in terms of their level of proficiency in signing, speaking, reading, and writing, so it is important to support access in all modalities. Deaf people can benefit from sign language processing in the same way that hearing people benefit from spoken language processing, plus the technology opens new possibilities for bridging the communication gap between Deaf and hearing communities. Despite this opportunity for positive social impact, progress in sign language AI has been slow — both because building AI for sign languages presents complex challenges and because widespread misconceptions exist about how the languages themselves work.
Compared to spoken language transcription, sign language translation presents two core challenges. First, transcribing speech is a matter of performing a sequential mapping from sound to text in the same language, whereas sign languages are independent, natural languages with their own distinct grammars and lexicons. As a result, they require true machine translation rather than a sequential process of sign-to-word transformations. Second, the model must learn to “see” and understand physical movement. Sign languages convey meaning through simultaneous movements of the hands, arms, torso, head, and face. Accurately tracking these at high frame rates is a difficult and computationally demanding computer vision task.
Given this background, it is easy to understand why some early attempts at sign language technology, like sign language gloves, were fundamentally limited: sign languages aren't simply “English on the hands.” They require complex visual perception of fine-grained whole-body movements and full-fledged language translation. SL2T is designed to deliver both.
SL2T sees sign language inputs as points on the signer's body and translates them into streaming text outputs. Example from the FLEURS-ASL benchmark.
We built SL2T by combining a user-centric, culturally informed approach with massive data scaling. The model is trained on over 100,000 hours of data across more than 50 sign languages — with roughly a quarter of the data in ASL. Training jointly on diverse languages, dialects, and proficiency levels causes the model to learn shared underlying structures, outperforming single-language models in our experiments.
To protect user privacy, SL2T sees sign language as a sequence of pose landmark locations rather than a raw camera feed. An on-device model ( MediaPipe Holistic ) tracks the location of points on the signer, and only these geometric coordinates are sent to the server for translation, allowing the original video to be discarded immediately.
SL2T translates this coordinate sequence directly into text, bypassing intermediate annotations known as “glosses” that are widely used in prior work on sign language translation. Glosses fail to capture rich, non-linear aspects of sign languages such as non-manual markers and spatial constructions. Translating directly from landmarks removes artificial vocabulary limits and allows translation quality to scale directly with data.
SL2T is the most capable sign language translation model to date according to key benchmarks like FLEURS-ASL (sd-test), which assesses ASL to English translation quality. SL2T achieves a remarkable zero-shot score of 70 BLEURT, which is significantly higher than any previously reported score. But optimizing academic benchmarks alone doesn’t guarantee usability in real-world applications, so we worked hard on practical issues like minimizing streaming latency, preventing hallucination on non-signing inputs, ensuring fairness for the 10% of signers who are left-handed, and improving performance for one-handed signing, which is used while holding a smartphone in the other hand.
Examples from the FLEURS-ASL benchmark. SL2T accurately translates complex ASL into fluent English. Occasional errors remain in rare signs, rapid fingerspelling ("prey" → "grey"), passive constructions, classifier depictions (dropping "claws"), and tense without context ("kicked off" → "start").
We believe in building with the Deaf community, not just for it. Deaf perspectives have shaped every stage of this project — from conceptualization by Sam Sepah, a Deaf Googler, to data collection with Deaf partners, evaluation in Deaf user studies, and impact assessment of the technology with Deaf experts.
To guide responsible real-world deployment, we established the AI Sign Language Advisory Committee (AISLAC), bringing together many global Deaf organizations and subject-matter experts. Through this participatory governance model, the communities most impacted by our technology directly influence our development priorities. We co-authored a joint impact report for the release of SL2T 1.0 in Gboard and Live Transcribe, transparently detailing the technology's capabilities and current limitations — a collaborative approach we plan to continue for all major sign language releases.
SL2T builds upon decades of foundational research across academia and industry, but bringing ASL input to users’ phones is only the beginning. Google’s mission is to organize the world's information and make it universally accessible and useful. Achieving universal accessibility means reaching full parity with spoken and written languages. Our team is working to expand this technology into additional sign languages, sign language generation, and frontier AI capabilities. We look forward to sharing our progress responsibly in order to make access through sign languages standard across the digital landscape.
You can experience SL2T in Gboard and Live Transcribe first on Pixel 11 , with more devices coming soon — all at no additional cost.
This work was done jointly by teams from Google DeepMind and Android. The core team who developed the SL2T model is: Garrett Tanzer, Benoit Brard, Elizabeth Clark, Tim Dozat, Sebastian Ebert, Dan Garrette, Manfred Georg, Vicky Holgate, Shankar Kumar, Mohammad Saboorian, Miloš Stanojević, Megh Umekar, John Wieting, Andy Zhang, and Chris Dyer.
The Android team who integrated the model into Gboard and Live Transcribe is: Ausmus Chang, Sai Aditya Chitturu, Dayle Chiu, Anna Chou, Ajay Dudani, Angana Ghosh, Alex Huang, Joanne Kim, Ed Lee, Thomas Lin, James Su, Yanchao Su, and Sharlene Yuan.
We are grateful for additional support from Anelia Angelova, Abhishek Bapna, Sara Basson, Glenn Cameron, Scott Crowell, Trevor Cohn, Noah Fiedel, Zoubin Ghahramani, Raia Hadsell, Tom Hudson, Alexander Hauerslev Jensen, Kazuya Kawakami, Peike Li, Liam McCafferty, Caroline Pantofaru, Abhinav Parashar, Christopher Patnoe, Laura Rimell, Sagar Savla, Sam Sepah, Thad Starner, Dave Uthus, and Biao Zhang.
Many thanks also go to those who participated in early stage testing of our models.
I accept Google's Terms and Conditions and acknowledge that my information will be used in accordance with Google's Privacy Policy .
