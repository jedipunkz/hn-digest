---
source: "https://www.game-oracle.com/blog/ai-part2"
hn_url: "https://news.ycombinator.com/item?id=48611989"
title: "AI in Games: The Impact on Sales"
article_title: "AI in Games: The Impact On Sales"
author: "Macha"
captured_at: "2026-06-20T21:34:39Z"
capture_tool: "hn-digest"
hn_id: 48611989
score: 2
comments: 0
posted_at: "2026-06-20T19:06:40Z"
tags:
  - hacker-news
  - translated
---

# AI in Games: The Impact on Sales

- HN: [48611989](https://news.ycombinator.com/item?id=48611989)
- Source: [www.game-oracle.com](https://www.game-oracle.com/blog/ai-part2)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T19:06:40Z

## Translation

タイトル: ゲームにおける AI: 売上への影響
記事タイトル: ゲームにおける AI: 売上への影響
説明: AI が Steam でのゲームの成功にどのような影響を与えるかを探ります。このデータ主導の分析により、AI の使用状況と販売実績との関連性が明らかになり、Steam ゲーム内の AI 生成コンテンツが購入決定に影響を与えているかどうかが明らかになります。

記事本文:
ゲームにおける AI: 売上への影響
ゲームオラクル
リサーチツールについて 価格設定 ブログ リーダーボード エンタープライズ FAQ ログイン 無料で始める ゲームにおける AI: 売上への影響
著者: Ross Burton、PhD、製品およびデータ責任者 カテゴリ: データ分析 公開日: 2025 年 12 月 17 日 更新日: 2025 年 12 月 6 日 AI タンクの使用は販売に影響しますか?
Steam 上の生成 AI に関するこのシリーズのパート 1 では、AI の公開が急速に加速しており、2025 年にリリースされたゲームの約 21% が何らかの形で AI の使用を宣言していることを紹介しました (11 月時点)。このブログでは、分析の 2 番目の、より複雑な部分、つまり AI の使用が販売にどのような影響を与えるかに取り組みます。
ここで尋ねることができる質問はたくさんありますが、それらはすべて、プレーヤーが AI をどのように認識するか、およびテクノロジーがゲームの品質に与える影響に関する仮定に基づいています。プレイヤーは AI 開示のあるゲームを積極的に避けますか?プレイヤーは AI の使用に気づいていませんか? AI の使用は本質的にゲームの品質を低下させ、売上に影響を及ぼしますか?生成 AI は経験の浅いチームやリソースが不足しているチームだけが使用しており、そもそも売上への影響が偏っているのでしょうか?
これらの質問はすべて信じられないほど複雑です。私たちは、結果に偏りがないことを確認するために細心の注意を払った、非常に思慮深い統計分析によってこの問題に対処しようとしています。基本的に、私たちは 1 つの明確な質問に焦点を当てます。「開発者が AI を使用した場合、そのゲームは AI を使用しなかった場合と比較して何件のレビューを獲得できるでしょうか?」
「えっ？レビュー？ここで販売のことを話しているのかと思ったのですが？」とすぐに疑問に思うかもしれません。残念ながら、販売数は開発者のみが知っているため、販売数の適切な代用となるレビューの数に焦点を当てます。このプロキシは業界全体で使用されており、詳しく書かれており、その方法についてはブログにまとめられています。

Steam の売上を推定できます。
私たちの質問に完全に答えるには、多くの難しい決定を下す必要があります。そのため、ここでの前提を理解できるように、すべてを並べてみましょう。
私たちの最大の前提は、開発者が Steam 上で AI の使用を宣言しているということです。これは、Steam の厳格なポリシーであるため、合理的な仮定のように思えます。開発者は、ゲーム開発中に AI をいつ、どのように使用するかを宣言する必要があります。もちろん、一部の開発者が AI の使用を開示しないことを選択する可能性はありますが、主な収入を管理するプラットフォームからのそのような明確なポリシーを無視することは、単に AI の使用を正直に宣言するよりも危険であるため、ここでは仮定に固執しました。
また、物事をシンプルにして、AI 宣言を「どこかの開発で AI が使用された」ことを意味するものとして扱うことにしました。これが完璧ではないことは承知していますが、パート 1 で示したように、AI の開示は非常に厄介で、開示の内容に基づいてデータを分割すると、多くの潜在的なバイアスが生じる危険性があります。
Steam のレビューは複雑で、すべてが売上に直接反映されるわけではありません。最近、Game Oracle でのレビュー データの収集方法を変更し、Steam での購入から直接レビューのみを測定するようにしました。ただし、この分析では、これが事実であることを保証できませんでした。したがって、レビューの総数が実際の売上と高い相関があるという仮定に基づいて、すべてのレビュー (サードパーティ経由で購入されたかどうかに関係なく) を検討しました。私たちはこの仮定が公平であると考えています。
ほとんどのゲームの販売は発売直後に発生するため、結果として、発売後の最初の 1 か月間で受け取ったレビューの総数を測定しました。
AI の開示は Steam 市場では比較的新しい要素であり、以前の分析では各リリースの 15 ～ 25% であることが示されています。

月に AI の開示があったため、2025 年 1 月から 10 月までにリリースされたゲームのみに焦点を当てて分析しました。 11 月を除外したのは、これを書いている時点では、11 月にリリースされたすべてのゲームのリリース後の最初の 1 か月の完全なデータがなかったためです。
無料でプレイできるゲームや現在（2025 年 11 月時点で）リリースされていないゲームは除外したため、事前に分析したのは商業プロジェクトにのみ関係します。
上記の前提を決定した後でも、AI を使用してゲームが受け取る総レビューに与える全体的な影響を何を測定できるか、どのように分析するかについて、非常に深く考える必要がありました。これを行うための最良の方法は、図を描くことです。統計学ではこれを因果グラフと呼びます。
上の図をひと目見ただけで、おそらく「ああ、わかった、それで精神的に余裕がなくなったんだよね？」と思うでしょう。でも、付き合ってください...
私たちの因果関係図は、「これが売上に影響を与え、開発者が AI を使用するかどうかにも影響を与えると考えられるもの」を視覚的に示すものにすぎません。それぞれの矢印は、「X が Y に対して何らかの因果関係を持っていると仮定します」ということを示しています。オレンジ色の円は開発に AI が使用されたかどうかを表し、緑色の円は結果、つまりリリース後 1 か月間に受け取ったレビューの総数を表します。
AI の使用は、ゲームが獲得する平均評価と総レビュー数に影響を与えると考えています。 AI を使用するかどうかは、次の影響を受けます。
開発者の経験 - 経験の浅い開発者は、たとえば AI に依存する可能性が高くなります。
タイトルがパブリッシャーによってサポートされているかどうか - パブリッシャーによる AI の使用に関しては、いくつかのかなり厳格なポリシーが導入されています 1、2
ゲームがリリースされた月 - 前回のブログで、時間の経過とともに採用がどのように増加しているかを示しました
レビューの総数は、次のようなさまざまな要因によって影響されます。
リリースの月 - があることは誰もが知っています。

Steam における季節的な影響
リリース前のゲームのフォロワー数はウィッシュリストの代用になります。販売の場合と同様に、発売前にウィッシュリストの本当の数は開発者と Steam 以外には誰も知りません。
初期価格 - ゲームの価格は、割引や価格のローカライゼーションが主な要因となり、時間の経過とともに大きく変化します。残念ながら、現時点ではこのすべての変動を測定する信頼できる方法がありません。そのため、初期価格を含めることで十分であるように、初期価格が割引と価格のローカリゼーションの両方の影響に十分な影響を与えると仮定しました。
平均評価 - すぐに低評価のレビューは売上に深刻な影響を与えるため、総レビュー数に影響します。
作成されたゲームの種類 - Steam マップのおかげで、実際にこれを非常に詳細に説明できます。
さて、すべての仮定を説明するこの複雑な図ができましたが、実際に質問にどう答えるのでしょうか?
これを統計学の講義にしなくても、このような因果グラフを使用すると、「 do-calculus 」と呼ばれる論理フレームワークを使用できるようになります。このツールは、結果からバイアスを取り除くために測定する必要がある要素の最小数を特定するのに役立ちます。
これは、開発者が AI を使用するかどうかと、開発者が取得するレビューの数の両方に影響を与える外部要因という交絡因子があるため、非常に重要です。これらを考慮しないと、私たちの見積もりは間違ってしまいます。私たちの分析により、AI を宣言する場合と AI を宣言しない場合を公平に比較​​するには、開発者エクスペリエンス、パブリッシャーの支援、ゲームの種類、およびリリース月を制御する必要があることが明らかになりました。
したがって、複雑な数学の問題は、次のような明確な質問に単純化されます。
開発者エクスペリエンスの測定: 以前の Steam リリースの合計 (TPR)
完全な透明性のために最後にもう 1 つ。開発者のエクスペリエンスを明示的に測定することはできません。単にデータがないだけです。

ただし、私たちにできることは、開発者が Steam (TPR) でリリースしたゲームの数を測定することです。これは、エクスペリエンスを表す優れた指標です。
まともなゲームは、少なくとも商用リリースの場合、アイデア出し、開発、テスト、磨き上げ、公開に少なくとも 6 か月かかります。
Itch のような他のプラットフォームはラピッド プロトタイピング用に存在しますが、参加費 100 ドルの Steam はプロトタイピング、トレーニング プロジェクト、その他の小規模な非営利リリースには適していません (いずれにせよ、無料プレイ ゲームを分析から除外していることを忘れないでください)。
平均的な開発者/スタジオは、商用リリースの成功を目指して、特定のタイトルに少なくとも 6 か月を費やします。
もちろん例外もあるので、これは統計研究における他のものと同様、完璧ではありません。 Sokpop Collective を例に挙げます。これまで協力してリリースを進め、月に 1 本程度のゲームを公開してきた開発者のバンドです。彼らは粗末な成果を上げているわけではなく、独自の優れたビジネス戦略を持っているだけです。ただし、これらは外れ値であり、この分析の目的では大多数に焦点を当てたいと思います。全体として、異常に高い出版率と奇妙な初期価格を持つ開発者から 1932 のタイトルを分析から除外しました。
一見するとAIには問題があるように見える
スパムを除外し、2025 年 1 月から 10 月までの純粋に商用リリースに焦点を当てた結果、9,879 のゲームが分析対象になりました。これらのゲームのうち、17.9% に AI 宣言がありました。
統計モデルの結果に入る前に、簡単な視覚分析により、AI の使用がレビュー総数のわずかな減少、0 件のレビューの割合の増加、および 100 件未満のレビューの割合の増加と相関していることがわかります。
同じパターンがフォロワー (ウィッシュリストの代理) にも見られ、フォロー数の中央値は

AI を使用した Steam ゲームの rs は、AI を使用しないゲームの半分です。
選手のレセプションはどうですか？少なくとも 100 件以上のレビューを受けたゲームに焦点を当てると (レビュー スコアがある程度信頼できることを確認するため)、肯定的なレビューの割合の中央値は、AI を使用していないゲームの 88.3% と比較して、AI を使用したゲームでは 84.6% です。
これらすべては、AI の使用が Steam でのパフォーマンスの低下と少なくとも相関していることを意味します。
ただし、単に相関関係について話したいのではなく、AI の使用が売上に影響を与えているかどうかを確実に知りたいと考えています。私たちの特定の質問に答えるために、以前の制御変数 (開発者の経験、パブリッシャーの支援、ゲームの種類、リリース月) を取得し、次のことを問う統計モデルを構築しました。
同じ月内にリリースされ、同じパブリッシャーの支援状況と開発者のエクスペリエンスを持つ同じ種類のゲームの場合、AI の使用が発売後最初の 1 か月間で受け取ったレビューの数に与える全体的な影響はどれくらいですか?
私たちの結果は非常に大きなものでした。 2 人の開発者 (以下のインフォグラフィックの A と B) が、同じパブリッシャーの支援または真の「インディー」ステータスで同じタイプのゲームを制作し、同じレベルの経験を持ち、同じ月にゲームをリリースした場合、AI を使用しなかった開発者と比較して、AI を使用した開発者の総レビュー数が 52.6% (47.69% - 57.63%) 減少すると予想できます。つまり、AI を使用しなかった開発者 (開発者 A) がゲームに対して 100 件のレビューを得た場合、AI を使用した開発者 (開発者 B) は、同様の状況で同様のゲームを作成したにもかかわらず、47 件のレビューしか得られないと予想されます。
ここでの結果は私たちを本当に驚かせました。最初に視覚的に調査した後、ある程度の影響は予想していましたが、開発者が AI の使用によりレビュー数 (したがって売上) の半分を期待していたのは劇的でした。
私たちのモデルe

開発者のエクスペリエンスやパブリッシャーの支援が総レビューに大きなプラスの影響を与えるなど、他にもいくつかの安心感を与える非常に明白な効果について説明しました。作成されたゲームの種類によっても、総レビュー数がわずかに増減する可能性があり、程度は低いですが、ゲームがリリースされた時期によっても異なります。たとえば、夏にリリースされたゲームにはわずかな落ち込みがあります。
しかし、ノイズに次ぐ最大の効果は AI の使用でした。私たちのモデルはかなりうるさかったです。以前の因果関係図で「マーケティングと意識」の影響について言及したことを思い出してください。また、開発者のエクスペリエンスの尺度は不完全であることも率直に述べました。これらすべてが測定されていない分散を生み出し、結果に偏りが生じる可能性があります。
しかし、スキル、才能、マーケティング、そして単なる幸運についてはどうでしょうか?
懐疑的になるのは当然です。一部のゲームには「X-factor」が含まれていたり、適切なタイミングで適切な場所に配置されたりしていることは誰もが知っています。また、一部のゲームはすべてが正しく行われているのに、開発者が信じられないほど不運であることもわかっています。私たちはこの「X ファクター」を念頭に置いてモデルを作成し、「負の二項」モデルと呼ばれるものを選択しました。この種のモデルは、データでは測定できないランダム性の量を測定できます。

[切り捨てられた]

## Original Extract

Explore how AI influences game success on Steam. This data-driven analysis uncovers the link between AI usage and sales performance, revealing whether AI-generated content in Steam games is impacting purchase decisions.

AI in Games: The Impact On Sales
GAME ORACLE
About Research Tools Pricing Blog Leaderboards Enterprise FAQs Login Start For Free AI in Games: The Impact On Sales
Author: Ross Burton, PhD, Head of Product and Data Category: Data Analysis Published: 12/17/2025 Updated: 12/6/2025 Does Using AI Tank Sales?
In part 1 of this series on generative AI on Steam we showed how AI disclosure is accelerating rapidly, with ~21% of games released in 2025 declaring some form of AI use (as of November). In this blog, we'll be tackling the second and more tricky part of our analysis: how is the use of AI impacting sales?
There are a bunch of questions we can ask here all fuelled by assumptions around how players perceive AI and the impact of the technology on game quality. Do players actively avoid games with AI disclosures? Are players oblivious to the use of AI? Does using AI inherently reduce the quality of games, impacting sales? Is generative AI only used by inexperienced or under-resourced teams, biasing the impact on sales in the first place?
All of these questions are incredibly complicated. We're going to try and address this with a very thoughtful statistical analysis where we'll take a lot of care to make sure we're not biasing our results. Fundamentally we're going to focus on one clear question: "If a developer uses AI, how many reviews will their game get compared to if they didn't use AI?"
You might immediately ask "Wait? Reviews? I thought we're talking about sales here?". Unfortunately, the number of sales is only known by the developer and therefore we will be focusing on the number of reviews, which is a good proxy for the number of sales; this proxy is used across the industry, has been written about at length, and is summarised in our blog about how you can estimate Steam sales .
To answer our question in full we're going to have to make a lot of tricky decisions, so lets lay it all out so you understand our assumptions going into this:
Our biggest assumption is all this is that developers declare their AI usage on Steam. This seems like a reasonable assumption since it is strict Steam policy: developers must declare when and how AI is being used during game development. Of course, it is possible that some developers may choose not to disclose AI use, but ignoring such a clear policy from the platform that controls your primary income would be riskier than simply declaring AI use truthfully, so we stuck to our assumption here.
We also decided to keep things simple and treat AI declaration as meaning "AI was used in development somewhere". We know this isn't perfect, but as we showed in part 1, AI disclosures are really messy and splitting our data up based on the content of the disclosures risked introducing a lot of potential bias.
Reviews on Steam are complicated and not all of them directly reflect sales. We recently changed how we collect review data at Game Oracle, ensuring we only measure reviews directly from Steam purchases. However, for this analysis we couldn't guarantee this is the case. We therefore considered all reviews (regardless of whether they were purchased via a third-party or not) under the assumption that the total number of reviews is highly correlated with actual sales; we think this assumption is fair.
Most game sales occur immediately after launch, therefore, as our outcome we measured the total number of reviews received in the first month after release.
Since AI disclosure is a relatively new factor in the Steam marketplace and our previous analysis showed between 15 - 25% of releases each month had an AI disclosure, we focused our analysis exclusively on games released between January and October 2025; we excluded November because, at the time of writing this, we didn't have the complete first month post-release data for all the games released in November.
We excluded games that are free-to-play or currently unreleased (as of November 2025), so up-front our analysis is only relevant to commercial projects.
Even after deciding on the assumptions above, we then had to think really deeply about what we could measure and how we would analyse the total effect using AI has on the total reviews a game receives. The best way to do this is to draw out a diagram. In statistics we call this a Causal Graph.
One look at our diagram above and you're probably thinking "oh, I get it, so you had a mental break right?". But please stick with us...
Our causal diagram is just a visual way of saying "here is what we think impacts sales and also impacts whether a developer uses AI". Each arrow is saying "I assume X has some causal effect on Y". The orange circle represents whether AI was used in development or not and the green circle is the outcome i.e. the total number of reviews received in the first month post-release.
We think using AI will impact the average rating a game gets and the total number of reviews. Whether someone uses AI will be influenced by:
The developer's experience - less experienced devs might rely on AI more for example
Whether the title is backed by a publisher - there have been some pretty strict policies introduced around AI use by publishers 1,2
The month that the game was released - in our last blog we showed how adoption is increasing over time
The total reviews is influenced by a number of factors:
The month of release - we all know there are seasonal effects on Steam
The number of followers a game has pre-release which is a proxy for wishlists - like with sales, no one knows the true number of wishlists prior to launch other than the developer and Steam
The initial price - the price of a game varies greatly over time with discounts and price localisation being major factors. Unfortunately, we don't currently have a reliable method to measure all this variation so we assumed that the initial price sufficiently influences the impact of both discount and price localisation in such a way that including initial price is enough.
Average rating - poor reviews out-the-gate are going to severely impact sales and therefore total reviews
The type of game created - we can actually describe this in extreme detail thanks to our Steam Map
Okay, so we have this complex diagram that describes all our assumptions, but how do we actually answer our question?
Without turning this into a statistics lecture, a Causal Graph like this allows us to use a logical framework called " do-calculus. " This tool helps us identify the minimum number of factors we must measure to remove bias from our results.
This is crucial because of confounders : outside factors that influence both whether a developer uses AI and how many reviews they get. If we don't account for these, our estimates will be wrong. Our analysis revealed that to get a fair comparison between declaring AI and not declaring AI, we must control for Developer Experience , Publisher Backing , Type of Game , and Month of Release .
So, our complex math problem simplifies into a clear question:
Measuring Developer Experience: Total Previous Steam Releases (TPR)
One last thing for full transparency. We cannot explicitly measure developer experience, we simply do not have the data. However, what we can do is measure how many games a developer has released on Steam (TPR), which is a nice proxy for experience.
Any decent game, at least for commercial release, requires at least 6 months for ideation, development, testing, polishing, and publishing
Other platforms like Itch exist for rapid prototyping, whereas Steam with its $100 entrance fee, is not the place for prototyping, training projects, and other small non-commercial releases (remember we're removing free-to-play games from our analysis anyway)
Your average developer/ studio, with the intention of a successful commercial release, will dedicate at least 6 months to any given title
This, like anything in a statistical study, is not perfect because of course there are exceptions. Take the Sokpop Collective for example – a band of developers that has historically published ~1 game per month by working together on their releases – they're not producing slop, they just have a uniquely brilliant business strategy. They're an outlier though and for the purpose of this analysis we want to focus on the majority. Overall we removed 1932 titles from our analysis from developers with an unusually high publication rate and bizarre initial prices.
On the face of it, AI appears to be a problem
After filtering out spam and focusing on purely commercial releases between January and October 2025 we had 9879 games for analysis. Amongst these games, 17.9% had an AI declaration.
Before we dive into the results of our statistical model, a quick visual analysis shows that using AI is correlated with slightly lower total reviews, a higher proportion of 0 reviews and a higher proportion of less than 100 reviews:
The same pattern can be seen for followers (our proxy for wishlists) where the median number of followers for a Steam game using AI is half that of games that do not use AI.
What about player reception? Well, when focusing on games that receive at least 100 reviews (to ensure that review scores are somewhat reliable) the median % of positive reviews is 84.6% for games using AI compared to 88.3% of games not using AI.
All of this implies that using AI is at least correlated with poorer performance on Steam.
We don't want to simply talk about correlations though, we want to know for certain whether using AI is impacting sales. To answer our specific question, we took our control variables from before (developer experience, publisher backing, type of game, and month of release) and we built a statistical model to ask:
For the same type of games, released within the same month, and with the same publisher backing status and developer experience, what is the total impact of using AI on the number of reviews received in the first month after launch?
Our results were huge. If you took two developers (A & B in the infographic below), each producing the same type of game, with the same publisher backing or true 'indie' status, with the same level of experience and releasing their game in the same month, you could expect a 52.6% (47.69% - 57.63%) decrease in total reviews for the developer that used AI compared to the developer that did not. That means if the developer that didn't use AI (developer A) got 100 reviews for their game we would expect the developer that did use AI (developer B) to only get 47 reviews despite making a similar game under similar circumstances.
The results here really surprised us, after our initial visual exploration we expected some impact, but a developer expecting half the number of reviews (and therefore sales) due to AI use is dramatic.
Our model explained some other reassuring, and pretty obvious effects, such as developer experience and publisher backing having a large positive impact on total reviews. The type of game created could also slightly increase/decrease total reviews, and to a lesser extent the time of the year that the game was released. For example, there is a small dip for games released in the summer.
But the biggest effect was the use of AI, second to only one other thing: noise. Our model was pretty noisy. Remember before in our causal diagram that we mentioned the influence of "marketing & awareness". We also said up-front that our measure of developer experience is imperfect. All of this creates unmeasured variance which could bias our results.
But What About Skill, Talent, Marketing, and Just Plain Old Luck?
You would be right to be sceptical. We all know that some games just have an "X-factor" or find themselves in the right place at the right time. We also know that some games do everything right and the developer is just incredibly unlucky. We created our model with this "X-factor" in mind and chose something called a "Negative Binomial" model. This kind of model can measure the amount of randomness our data cannot ca

[truncated]
