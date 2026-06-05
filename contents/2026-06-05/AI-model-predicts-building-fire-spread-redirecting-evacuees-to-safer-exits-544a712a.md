---
source: "https://techxplore.com/news/2026-06-ai-redirecting-evacuees-safer-exits.html"
hn_url: "https://news.ycombinator.com/item?id=48405715"
title: "AI model predicts building fire spread, redirecting evacuees to safer exits"
article_title: "AI model predicts building fire spread, redirecting evacuees to safer exits in real time"
author: "lschueller"
captured_at: "2026-06-05T00:58:32Z"
capture_tool: "hn-digest"
hn_id: 48405715
score: 1
comments: 0
posted_at: "2026-06-04T22:48:03Z"
tags:
  - hacker-news
  - translated
---

# AI model predicts building fire spread, redirecting evacuees to safer exits

- HN: [48405715](https://news.ycombinator.com/item?id=48405715)
- Source: [techxplore.com](https://techxplore.com/news/2026-06-ai-redirecting-evacuees-safer-exits.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T22:48:03Z

## Translation

タイトル: AI モデルが建物火災の延焼を予測し、避難者をより安全な出口に誘導する
記事タイトル: AI モデルが建物火災の延焼を予測し、避難者をより安全な出口にリアルタイムで誘導
説明: 火災警報器が鳴り響き、あなたはオフィスのデスクから降り、最寄りの出口に向かいます。しかし、最も近い出口がすでに火事で塞がれている場合はどうなるでしょうか?米国国立標準技術研究所 (NIST) の研究者らは...

記事本文:
AI モデルが建物火災の延焼を予測し、避難者をより安全な出口にリアルタイムで誘導します
米国国立標準技術研究所による
編集者
ステファニー・ボーム、ロバート・イーガンによるレビュー
この記事は Science X に従ってレビューされています。
編集プロセス
そしてポリシー。
編集者が強調した
コンテンツの信頼性を確保しながら、次の属性を備えています。
優先ソースとして追加
NIST の研究者は、火災時に安全な避難経路を特定できる新しい AI モデルを開発しました。このモデルは、動的非常口ディスプレイと呼ばれる新しい電子出口標識と併用して、出口が安全に使用できるかどうかを示すことができます。クレジット: A. Kim / NIST
オフィスのデスクから火災警報器が鳴り響き、あなたは最寄りの出口に向かいます。しかし、最も近い出口がすでに火事で塞がれている場合はどうなるでしょうか?米国立標準技術研究所 (NIST) の研究者らは、火災時に居住者を最も安全な避難経路に誘導できる「Safe Step」と呼ばれる AI モデルを開発しました。 Journal of Building Engineering に記載されているように、このモデルを電子ディスプレイと併用して、出口が安全に使用できるかどうかを示すことができます。
「火災は成長し、広がる可能性があります」とNISTの研究員であり、このジャーナル論文の筆頭著者であるHongqiang "Rory" Fang氏は述べた。 「私たちのモデルは火災がどのように進行するかを予測し、非常口の表​​示を更新して人々を最も安全な出口に誘導するのに役立ちます。」
Safe Step は、センサーが温度や空気の質などの環境条件をリアルタイムで監視する「スマート」ビルディングで使用できます。これらの建物の一部では、動的非常口ディスプレイと呼ばれる新技術をテストしており、出口が安全に使用できることを示したり、建物からのより安全な出口を矢印で示したりすることができます。
以前の研究では、従来のアルゴンを使用することが提案されていました。

建物火災から安全に避難するための最短経路を見つけるためのゴリズム。ただし、これらのアルゴリズムは現在の建物の状況に完全に依存しており、避難者が経路に沿って直面する可能性のある累積的な危険は考慮されていません。
「私たちは自問しました。『火災がどのように進行するかを予測し、より多くの命を救う方法で、より優れたアルゴリズムを構築できないか?』」と NIST 機械エンジニアの Wai Cheong Tam 氏は述べています。
安全な避難のための機械学習
彼らのモデル「Safe Step」は、強化学習として知られる AI の一種を使用しています。試行錯誤を通じて最も安全なルートを決定します。 Safe Step は、建物のレイアウトを使用して避難ルートと NIST 火災シミュレーション ツールからのデータを学習し、レイアウト内の火災が時間の経過とともにどのように発生するかを予測します。
訓練中、モデルは火災が居住者に与える影響を予測し、より安全な避難経路に誘導する方法を学習します。実際の使用では、モデルは火災のシミュレーションをリアルタイムで実行する必要はありません。代わりに、建物からのライブセンサーデータに依存して、火災の進行に応じて推奨事項を継続的に調整します。
ただし、アルゴリズムが最適なルートを選択しているかどうかを判断するには数値が必要です。そこで NIST の研究者は、有毒ガスの部分実効線量 (FED) と呼ばれる火災安全指標を使用しました。この変数は、人が時間の経過とともにさらされる火災の危険の重大度を表します。 FED が低いほど、乗員の危険性は低くなります。このモデルは、乗員の移動に伴う有毒ガスへの曝露が時間の経過とともにどのように変化するかを考慮して、FED が最も低いルートを選択します。
次に研究者は、このモデルを 2 つのテスト ケースで使用して、従来のアルゴリズムと比較しました。彼らはまた、より複雑な 1 階建ての建物構造を使用し、そのモデルが一貫して安全な避難経路を提供することを発見しました。
のために

たとえば、廊下を隔てた部屋で火災が発生し、少量の煙が廊下に広がったとします。従来のアルゴリズムでは、居住者が廊下を渡って最も近い出口に到達するように誘導されます。しかし、火災が拡大し続け、居住者が廊下を渡って出口に近づくまでに非常に危険になったらどうなるでしょうか?その最寄りの出口はもはや安全な選択肢ではありません。 Safe Step はこの変化を予測し、動的な出口標識のデータを提供して、廊下の反対側にあるより遠くても安全な出口に乗員を誘導します。
現在のモデルは 1 階建てのフロア プランに対応しています。研究者の次のステップには、避難者が廊下を左または右に曲がるだけでなく、階を上り下りできる多層建築構造を処理するモデルの機能を改善することが含まれています。
複数の人の避難を最も正確にモデル化するために、研究者らは、各エージェントが異なる建物居住者に対応する、いわゆるマルチエージェントを備えた AI システムを構築することを計画しています。複数のエージェント間の対話により、モデルは実際の火災対応および避難シナリオにより適応できるようになります。
たとえば、火災の際には、複数の人が同時に建物から出ようとするため、建物の入り口で混雑が生じる可能性があります。これによりボトルネックが生じますが、改良されたアルゴリズムを使用すると、モデルは消防士が建物に入るアクセスポイントを調整しながら、避難者を別の出口に誘導することができます。これにより、消防士による消火や、高齢者、子供、障害者などの弱い立場にある人々の救助が容易になります。
NIST は、他の組織と協力して火災安全研究を推進してきた 1 世紀以上の経験があります。ほんの数十年の間に、煙の改善により

警報器や消防用装備など、NIST の火災研究は毎年の火災関連死亡者数の削減に重要な役割を果たしてきました。
研究者らは、Safe Step のような技術は 5 ～ 10 年以内に登場し始める可能性があると推定していますが、広く採用されるかどうかは規制当局の承認、信頼性テスト、既存の安全システムとの統合にかかっています。
「この研究はまだ初期の研究開発段階にありますが、先進技術を効果的に使用することで財産を保護し、命を救うことができるインテリジェントな消火活動に向けた重要な一歩を示しています」とファン氏は述べた。
Honqiang Fang 他、強化学習を使用した建物火災避難のためのテナビリティベースのパス計画モデルの開発、Journal of Building Engineering (2026)。 DOI: 10.1016/j.jobe.2025.115132
主要な概念
AI 交通安全 建築環境における AI
提供元
米国国立標準技術研究所
この話の背後にいるのは誰ですか?
ステファニー・ボーム
The New SchoolでTESOLの修士号を取得。言語学習と、生物学や宇宙探査に関する科学ニュースの編集に情熱を注いでいます。
詳しいプロフィール→
数理生物学で学士号、クリエイティブライティングで修士号を取得。旅行が豊富で、科学と言語について独自の視点を持っています。
詳しいプロフィール→
この記事は NIST のご厚意により再掲載されています。元のストーリーはここで読んでください。
Safe Step AI モデルは、強化学習とリアルタイムのセンサー データを使用して延焼を予測し、建物の居住者を最も安全な避難ルートに動的に誘導し、分数実効線量 (FED) で測定される有毒ガスへの曝露を最小限に抑えます。テストでは、Safe Step は、進化する危険を予測し、それに応じて出口推奨を調整することで、従来の最短経路アルゴリズムを上回りました。将来の開発は、このモデルを複数レベルの建物とマルチエージェント シナリオに拡張することを目的としています。

彼女は避難と緊急対応を最適化します。
この概要は、LLM を使用して自動的に生成されました。
完全な免責事項
タイプミスや不正確さを見つけた場合、またはこのページのコンテンツの編集リクエストを送信したい場合は、このフォームを使用してください。
一般的なお問い合わせについては、お問い合わせフォームをご利用ください。
一般的なフィードバックについては、以下のパブリック コメント セクションをご利用ください (ガイドラインに従ってください)。
リクエストの処理を容易にするために、最も適切なカテゴリを選択してください
編集者にフィードバックをお寄せいただきありがとうございます。
あなたのフィードバックは私たちにとって重要です。ただし、メッセージの量が多いため、個別の返信を保証するものではありません。
あなたの電子メール アドレスは、受信者に電子メールの送信者を知らせるためにのみ使用されます。あなたのアドレスも受信者のアドレスも他の目的には使用されません。入力した情報は電子メール メッセージに表示され、Tech Xplore によっていかなる形式でも保持されることはありません。
研究開発や最新情報に関する毎日の科学ニュース
科学的革新
医学研究の進歩と健康ニュース
ウェブ上で最も包括的な科学技術ニュース報道
このサイトは、ナビゲーションを支援し、サービスの使用を分析し、広告のデータを収集するために Cookie を使用します。
パーソナライズし、サードパーティからのコンテンツを提供します。
当社のサイトを使用すると、当社のプライバシーを読んで理解したことを認めたものとみなされます
ポリシー
および利用規約
を使用します。

## Original Extract

A fire alarm jolts you from your office desk, and you head for the nearest exit. But what if the closest exit has already been blocked by the fire? Researchers at the National Institute of Standards and Technology (NIST) ...

AI model predicts building fire spread, redirecting evacuees to safer exits in real time
by National Institute of Standards and Technology
edited by
Stephanie Baum , reviewed by Robert Egan
This article has been reviewed according to Science X's
editorial process
and policies .
Editors have highlighted
the following attributes while ensuring the content's credibility:
Add as preferred source
NIST researchers developed a new AI model that can identify safe evacuation routes during a fire. The model can be used with new electronic exit signs, called dynamic emergency exit displays, to show whether an exit is safe to use. Credit: A. Kim / NIST
A fire alarm jolts you from your office desk, and you head for the nearest exit. But what if the closest exit has already been blocked by the fire? Researchers at the National Institute of Standards and Technology (NIST) and their colleagues have developed an AI model called Safe Step that can redirect occupants to the safest evacuation route in a fire. Described in the Journal of Building Engineering , the model can be used with electronic displays to show whether an exit is safe to use.
"Fires can grow and spread," said Hongqiang "Rory" Fang, a research associate at NIST and first author of the journal paper. "Our model forecasts how the fire is evolving and can help update emergency exit displays to direct people toward the safest exit."
Safe Step can be used in "smart" buildings, where sensors monitor real-time environmental conditions, such as temperature and air quality. Some of these buildings are testing a new technology called a dynamic emergency exit display, which can indicate that the exit is safe to use or point arrows to a safer route out of the building.
Previous research has proposed using traditional algorithms to find the shortest path for safely evacuating a building fire. However, these algorithms depend entirely on current building conditions and do not consider the cumulative hazards that evacuees can face along the route.
"We asked ourselves, 'Can we build a better algorithm that predicts how the fire evolves, and in a way that helps save more lives?'" said NIST mechanical engineer Wai Cheong Tam.
Machine learning for safe evacuations
Their model, Safe Step, uses a type of AI known as reinforcement learning. It makes decisions on the safest routes through trial and error. Safe Step uses the building layout to learn evacuation routes and data from a NIST fire simulation tool to anticipate how a fire in the layout develops over time.
During training, the model learns to forecast how a fire will affect occupants and then guides them to safer evacuation routes. In real-world use, the model does not need to run a simulation of the fire in real time. Instead, it would rely on live sensor data from the building to continuously adjust its recommendations as the fire evolves.
The algorithm needs numbers, though, to determine whether it's choosing the best route. So NIST researchers used a fire safety metric called the fractional effective dose (FED) of toxic gases. This variable represents the severity of fire hazards to which a person is exposed over time. The lower the FED, the lower the hazard exposure for the occupants. The model chooses the route with the lowest FED, accounting for how toxic gas exposure changes over time as an occupant moves.
Researchers then used the model in two test cases to compare with the traditional algorithm. They also used a more complex single-level building structure and found that the model consistently gave safe evacuation routes.
For example, suppose a fire starts in a room across the hallway, and a small amount of smoke spreads into the hallway. A traditional algorithm would guide the occupant to cross the hallway to get to the closest exit. But what happens if the fire continues to grow and becomes extremely dangerous by the time the occupant crosses the hallway and approaches the exit? That nearest exit is no longer a safe option. Safe Step can anticipate this change and provide data for dynamic exit signs to direct the occupant to a more distant but safer exit at the opposite end of the hallway.
The current model works for a single-story floor plan. Researchers' next steps include improving the model's capabilities to handle multilevel building structures, where an evacuee can go up or down a floor in addition to turning left or right down a hallway.
To most accurately model the evacuation of multiple individuals, researchers plan to build an AI system that has what is known as multiple agents, with each agent corresponding to a different building occupant. Interactions among multiple agents will make the model more adaptable to real fire response and evacuation scenarios.
For instance, during a fire, congestion can build up at the building's entrance as multiple people try to exit at the same time. This creates a bottleneck, but with an improved algorithm , the model could direct evacuees to different exits while coordinating access points for firefighters to enter the building. This would make it easier for firefighters to extinguish the fire or rescue vulnerable individuals, such as older adults, children and people with disabilities.
NIST has more than a century of experience working with other organizations to advance fire safety research. In just the last several decades, by improving smoke alarms and firefighter gear, NIST's fire research has played a crucial role in reducing fire-related deaths each year.
Researchers estimate that technologies like Safe Step could start appearing in five to 10 years, though widespread adoption will depend on regulatory approval, reliability testing and integration with existing safety systems.
"This research is still in the early R&D stage, but it represents an important step toward intelligent firefighting where effective use of advanced technologies can protect property and save lives," said Fang.
Hongqiang Fang et al, Development of a tenability-based path planning model for building fire evacuations using reinforcement learning, Journal of Building Engineering (2026). DOI: 10.1016/j.jobe.2025.115132
Key concepts
AI traffic safety AI in built environment
Provided by
National Institute of Standards and Technology
Who's behind this story?
Stephanie Baum
Master's in TESOL from The New School. Passionate about language learning and editing science news on biology and space exploration.
Full profile →
Bachelor's in mathematical biology, Master's in creative writing. Well-traveled with unique perspectives on science and language.
Full profile →
This story is republished courtesy of NIST. Read the original story here .
The Safe Step AI model uses reinforcement learning and real-time sensor data to predict fire spread and dynamically guide building occupants to the safest evacuation routes, minimizing exposure to toxic gases as measured by fractional effective dose (FED). In tests, Safe Step outperformed traditional shortest-path algorithms by anticipating evolving hazards and adjusting exit recommendations accordingly. Future developments aim to extend the model to multilevel buildings and multi-agent scenarios to further optimize evacuation and emergency response.
This summary was automatically generated using LLM.
Full disclaimer
Use this form if you have come across a typo, inaccuracy or would like to send an edit request for the content on this page.
For general inquiries, please use our contact form .
For general feedback, use the public comments section below (please adhere to guidelines ).
Please select the most appropriate category to facilitate processing of your request
Thank you for taking time to provide your feedback to the editors.
Your feedback is important to us. However, we do not guarantee individual replies due to the high volume of messages.
Your email address is used only to let the recipient know who sent the email. Neither your address nor the recipient's address will be used for any other purpose. The information you enter will appear in your e-mail message and is not retained by Tech Xplore in any form.
Daily science news on research developments and the latest
scientific innovations
Medical research advances and health news
The most comprehensive sci-tech news coverage on the web
This site uses cookies to assist with navigation, analyse your use of our services, collect data for ads
personalisation and provide content from third parties.
By using our site, you acknowledge that you have read and understand our Privacy
Policy
and Terms of
Use .
