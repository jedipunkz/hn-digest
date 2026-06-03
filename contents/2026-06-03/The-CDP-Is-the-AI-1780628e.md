---
source: "https://www.jacquescorbytuech.com/writing/cdp-ai"
hn_url: "https://news.ycombinator.com/item?id=48375697"
title: "The CDP Is the AI"
article_title: "The CDP is the AI | Jacques Corby-Tuech"
author: "iamacyborg"
captured_at: "2026-06-03T00:36:11Z"
capture_tool: "hn-digest"
hn_id: 48375697
score: 2
comments: 0
posted_at: "2026-06-02T20:20:47Z"
tags:
  - hacker-news
  - translated
---

# The CDP Is the AI

- HN: [48375697](https://news.ycombinator.com/item?id=48375697)
- Source: [www.jacquescorbytuech.com](https://www.jacquescorbytuech.com/writing/cdp-ai)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T20:20:47Z

## Translation

タイトル: CDP は AI です
記事のタイトル: CDP は AI |ジャック・コルビ＝テューシュ
説明: 私は最近、ブルームリーチの朝食会場にいたのですが、その部屋は大小のブランドのフィンテック マーケティング担当者でいっぱいで、全員が「AI」について話していました。実際にはそうではありませんでした。というか、まだ余裕のない休暇について話すのと同じように、私たちはそれについて話しました。部屋にいる全員からの繰り返しのテーマ。
[切り捨てられた]

記事本文:
CDP は AI |ジャック・コルビ＝テューシュ
ジャック・コルビ＝テューシュ
書く
読書
リンク
仕事
2026 年 6 月 2 日
CDP は AI です
ブランドが実際に購入できるもの: サポートと意思決定
ほとんどのブランドはそれを使用していません
なぜ格差が存在し、拡大しつつあるのか
これがライフサイクル ロールにとって何を意味するのか、そしてそれが最終的にどこに向かうのか
私は最近、ブルームリーチの朝食会場にいたのですが、その部屋は大小のブランドのフィンテックマーケターでいっぱいで、全員が「AI」について話していました。実際にはそうではありませんでした。というか、まだ余裕のない休暇について話すのと同じように、私たちはそれについて話しました。会社の規模に関係なく、会場にいる全員が繰り返し話題にしたのは、そもそもデータを使用可能な形式にすることでした。どのモデルを使用すればよいか、どの賢い機能をオンにすればよいか誰も迷っていませんでした。彼らは、データが混乱しており、相互に通信しない 6 つのシステムに分散しており、それが修正されるまでは、どんな賢いことも何の意味も持たないという事実に囚われていました。
本当にこれが全体であり、ここでやめても構いません。しかし、部屋がぶつかり続けるギャップは、「データを整理する必要がある」というよりも大きく、構造的なものであり、そのギャップは拡大し続けています。
市場の一端には、ブランドが購入できるものよりも数世代先の機能を備えた、メッセージングのための実動機械学習を実行する少数の消費者向けプラットフォームがあります。もう一方の端には、ライフサイクル スタックにおける ML へのアクセスが「ESP のいくつかの予測機能」から「実際には何も有効にしていない」までにわたる中堅および中小企業ブランドのロングテールがあります。その中間にある、より新しいカテゴリーの製品が、よりフロンティアに近いものをブランドに販売しようとしています。問題は、朝食のテーブルで常に話題になっていた問題ですが、すべてはデータが整理されているかどうかにかかっており、ほとんどのブランドはそうではないということです。
貿易というもの

報道陣が言う「チャネルの成熟度」は、AI の高度さによる分類としてより正確に説明されます。一部の送信者はチャネルを機能させる余裕があります。ほとんどはできません。そして、決定要因は予算や人員ではなくなりつつあります。重要なのは、データが何でもできる状態にあるかどうかです。
少数の消費者向けテクノロジー企業が、生産通知およびメッセージング システムに関するピアレビュー済みの成果物を出版しています。これらはホワイト ペーパーやベンダーの事例研究ではありません。これらは、KDD、RecSys、WSDM、および CIKM の提出物であり、博士号を擁するチームによって作成され、学術審査員の前で自分たちの方法論を提示し、それに答える必要があります。これらの論文は、企業が内部で行っていることを網羅的に説明しているわけではありません (実際にはそうではありません) が、能力の基礎としては役立ちます。内部システムは公開されるものと少なくとも同等、通常はそれよりも優れており、公開するという行為は、チームがこの作業を真剣に行うための組織的な支援があることを示しています。
目に見えるフロンティアの簡単なツアー:
Pinterest は、通知の増分価値がカジュアル ユーザーに対して最も高いという結果に基づいて、ユーザーごとに週次の通知予算を設定し、クリックスルーではなく長期的なサイト エンゲージメントを最適化しました。ヘビー オープナーは、通知によって動かされたためではなく、すべてに関与しているため、クリックスルーが高くなります。 1
Duolingo は、バンディット アルゴリズムを使用して各ユーザーに送信するリマインダー テンプレートを選択し、強力なベースラインと比較して、毎日のアクティブ ユーザー数が 0.5% 増加し、新規ユーザー維持率が 2% 増加したと報告しました。 2
Twitter はモデルベースの強化学習を使用して、プッシュを送信するかどうかを決定し、数日間にわたる効果をモデル化しました。公表されているトレードオフは興味深い部分です。ボリュームを最も大幅に削減する設定は開封率を 14% も押し上げましたが、同じ設定でも

■ 毎日のアクティブ ユーザーの減少。最も控えめな設定のみで、開封率が約 8% 向上し、デイリーアクティブがまったく改善され、その後 0.2% 改善されました。見出しの数を最大化し、真の目的を果たすことは、反対の方向を向いていました。 3
LinkedIn は、通知の決定をオフライン強化学習、つまり保守主義ペナルティを伴う Double Deep Q-Network として構成し、ログ データに基づいてトレーニングして展開しました。セッションは 4 分の 1 パーセント増加し、クリックスルーは数ポイント増加し、通知量はすべて同時に減少しました。 4 2026 年までに、同じ系統が電子メールに到達しました。BanditLP は、ニューラル トンプソン サンプリングと、ビジネス上の制約の下で各メンバーに送信される内容を選択する数十億の変数に十分な大きさの線形プログラムを組み合わせました。 5
Zillow は、ユーザーごとに送信するかどうかを決定するブースト ツリー分類子を使用して電子メールとプッシュの量を管理し、過剰な送信とそれによって引き起こされる購読解除を排除しながら、クリックの 98% を維持するように調整されています。強化学習は必要ありません。これは、それ自体が教訓です。このリストにある最も安価な方法は、送信量が少なくても依然として勝ちます。 6
Meta は Instagram の通知スロットをオークションとして扱いました。メッセージを送りたい 550 以上の社内チームが互いに入札し合い (プラットフォームは入札を補助できる)、単一のユーザーが競合製品の表面に殺到することはありません。テストでは、各アームあたり 7,700 万人のユーザーにわたって、送信する通知がわずかに減り、クリックスルーが増加し、リーチはそのまま残されました。 7
そしてフロンティアは動き続けます。中国のショートビデオ プラットフォームである Kuaishou に導入され、2026 年 2 月の WSDM で発表された PushGen は、スタイル制御の下で LLM を使用してプッシュ コピーを生成し、学習した報酬モデルを使用して候補をランク付けし、クリックスルーを予測して、1 日に数億人のユーザーを対象に勝者を選択します。 8 Pinterest の TransAct が指摘する点

方法: トランスフォーマーがユーザーのリアルタイム アクティビティを読み取り、ホームフィード、検索、通知全体でランキングをフィードし、プッシュ開封率と電子メールのクリックスルーがそれぞれ 1 ～ 2 ポイント向上しました。 9
これらのシステムには次の 4 つの特徴があります。
プラットフォーム所有者によって完全なセッションレベルの粒度で収集されたファーストパーティのイベント ストリームに基づいて構築されます。
社内のエンジニアリング チームによって運営されており、そのチームには研究者、ML エンジニア、プラットフォーム エンジニアが含まれており、その人数はほとんどの B2C ブランドがスタッフを配置しようとしても配置できないほどです。
オープンやクリックではなく、プラットフォーム独自の長期的な価値機能、セッション、毎週のアクティブ、複数日間の保持、サイト全体のエンゲージメントに合わせて調整されています。
ユーザーの応答は、悪用される固定信号ではなく、送信されたメッセージに応じて変化するものであることが前提となります。
最後の特徴は、ブランドに販売されるほぼすべてのものが間違っているということです。静的モデルは、どのユーザーが開封または購入する可能性が最も高いかを尋ね、そのユーザーをターゲットにします。適応型の場合は、別の質問をします。これを送信すると、放っておいたバージョンのユーザーの行動がどのように変わりますか?この反応は、あなたが見つけて利用する固定的な特性ではありません。これはメッセージ自体が移動する結果であり、ある人にとっては上に、他の人には下に移動します。そのため、同じモデルはまったく何も送信しないようにする必要があります。エンゲージメントしそうな人を最適化すると、メッセージが間違った方向に進む人たちに静かに迷惑をかけながら、関係なくエンゲージメントしようとしていた人たちにメッセージを送り続けることになります。
実際の目標、セッションや毎日のアクティビティの上昇率は、ほとんどの場合 1 パーセント未満です。より大きく見える利益は、開封率などの代理指標に基づいています。フロンティアの利点は巨大なリフトではありません。それは、何億人ものユーザーがいるプラットフォームがセッションに対して信頼性の高い 0.3% を預けることができるということです。

これは絶対額で言えば巨額だが、ベースのほんの一部で同じ 0.3% を追い求めている中堅ブランドは、それを獲得するためのエンジニアリングを正当化することはできない。このギャップは規模の経済性によるものであり、大きさではありません。
ほぼすべての利益は、送信量を減らすこと、より正確には送信方法を変えることによって得られます。エンゲージメントが維持または上昇する間、出来高は減少します。この比較は相対的なものであることに注意してください。これらのプラットフォームは、ほとんどのブランドが送信する量を既に大幅に上回っているユーザーごとの量を削減しているため、「量の削減」はロングテールがかつてなかった体制から始まります。そして、その体制の中でも「送信量を減らす」という解釈はあまりにも率直すぎます。なぜなら、これらのシステムは均一に削減していないからです。一部のユーザーには再割り当てが行われ、他のユーザーには割り当てられません。 Twitter の最適な設定では、総送信数が減少したにもかかわらず、ユーザーごとの送信上限が実際に引き上げられました。これは、ポリシーがヘッドルームに値する人をより厳選したためです。スキルは誰にメッセージを送信すべきではないかを知ることですが、これは誰にメッセージを送信すべきかよりも学ぶのが難しく、ロングテールに欠けているユーザーごとのシグナルを正確に必要とします。上昇率の研究でも同じ点が指摘されています。累積的な増分効果は、リスト全体に到達するかなり前、通常は直感が示唆するよりもはるかに早くピークに達しますが、そのピークを超えるとターゲットを絞ると効果が減少します。これは、人々にリーチし始めると、メッセージが勝つのではなくオフになるためです。リスト全体に送信すれば、ゼロかマイナスの領域に十分入ります。
ブランドが実際に購入できるもの: サポートと意思決定
ここで最もよく崩れる区別は、マーケターの意思決定をサポートする機械学習と、マーケターの意思決定を行う機械学習との間の区別です。
ブランドのライフサイクル スタックに含まれる ML のほとんどは、意思決定サポートです。より小規模で新しいカテゴリは、意思決定です。この 2 つを混同すると、平均的なブランドが実際に何であるかをひどく過大評価してしまうことになります

んん。
意思決定サポートは、大手顧客エンゲージメント プラットフォームが提供するものです。
すでに送信することに決めたメッセージの時間とルートを選択する、送信時間と最適チャネルの最適化。
セグメントを構築できる、チャーン、コンバージョン、購入傾向、生涯価値の予測スコア。
予測 RFM およびその他の「AI」セグメントは、最新性、頻度、支出、または可能性の高い次のアクションによってユーザーをグループ化します。
購入履歴から製品の推奨を削除します。
件名提案と生成コピー、Phrasee (現 Jacquard) の専門分野は、LLM の波が起こるずっと前から構築されており、ブランド ボイスのバリエーションを生成し、過去のパフォーマンスに基づいてランク付けしていました。
フリークエンシー キャップと送信量の管理。
これらのほとんどは何かをします。彼らが多くのことを行っているかどうかは言うのが難しい。たとえば、送信時間の最適化が実際に賢明な固定時間に送信することを上回るかどうかは、ほぼ完全にベンダーのケーススタディに依存する一方、存在する独立した学術研究は、固定されたスケジュールされた時間を支持し、通知の過負荷について警告する傾向がある。 10 しかし、共通点は、リフトの有無にかかわらず、マーケティング担当者がジャーニーをデザインするということです。 ML は、人間が構築した構造内の特定のポイントでスコアを付け、提案し、ランク付けし、最適化します。意思決定を支援します。それは間に合いません。
これは、ほぼすべてのベンダーが該当する範囲です。
Salesforce Marketing Cloud の Einstein は、送信時間、コンテンツ選択、予測スコアリング機能を備え、Agentforce ブランドに組み込まれました。
Adobe CX Enterprise (Experience Cloud at Summit 2026 からブランド変更)、AI アシスタント、予測視聴者、自動送信時間、コンテンツ選択、生成バリアント、AI コワーカーと専用エージェントの新しい層を備えています。
Klaviyo の CLV とチャーン、送信時間、件名の最適化、GE の予測分析

ネイティブコピー。
Iterable の予測目標、送信時間の最適化、ブランド アフィニティ、チャネルの最適化。
HubSpot の Breeze がワークフロー全体に広がります。
Emarsys は現在、SAP Engagement Cloud に統合されており、コンバージョン、チャーン、支出に関する AI スコア、継続的に更新される予測セグメント、製品の推奨事項の予測、自然言語カタログ検索や AI レポート ビルダーなどの生成ヘルパーの山が増えています。
Braze の Sage AI は、送信時間、チャネル、予測チャーン、およびコンテンツの提案を行います。
CleverTap、MoEngage、Insider などもここでプレーしていますが、よりニッチで、全面的というよりはモバイルファーストや特定の地域に強いです。
これらのベンダー 3 社のマーケティング ページを続けて読んでも、それらを区別することはできません。語彙は同一であり、スクリーンショットは交換可能です。それは偶然ではありません。マルチテナント ML 機能が適切に実行できるセットは有限であり、ベンダーはすべて、送信時間、予測セグメント、生成コピー、分岐点でのスコアリングを伴うジャーニー レベルのオーケストレーションなど、ほぼ同じセットに収束しています。 MessageGears は、送信時間とチャネルの最適化について、誰よりもわかりやすく標準キットを文書化しています。

[切り捨てられた]

## Original Extract

I was at a Bloomreach breakfast recently, a room full of fintech marketers from brands large and small, all there to talk about "AI". We didn't, really. Or rather, we talked about it the way you talk about a holiday you can't afford yet. The recurring theme, from everyone in the room regardless of c
[truncated]

The CDP is the AI | Jacques Corby-Tuech
Jacques Corby-Tuech
writing
reading
links
work
2 June 2026
The CDP is the AI
What brands can actually buy: support versus decisioning
Most brands aren't using any of it
Why the gap exists and is widening
What this means for the lifecycle role, and where it ends up
I was at a Bloomreach breakfast recently, a room full of fintech marketers from brands large and small, all there to talk about "AI". We didn't, really. Or rather, we talked about it the way you talk about a holiday you can't afford yet. The recurring theme, from everyone in the room regardless of company size, was getting their data into a usable shape in the first place. Nobody was stuck on which model to use or which clever feature to switch on. They were stuck on the fact that their data is a mess, spread across half a dozen systems that don't talk to each other, and until that's fixed none of the clever stuff means anything.
That's the whole piece, really, and I could stop there. But the gap that room kept bumping into is bigger and more structural than "we need to sort our data out", and it's widening.
At one end of the market there's a small number of consumer platforms running production machine learning for messaging that is several capability generations ahead of anything a brand can buy. At the other end there's the long tail of mid-market and SMB brands whose access to ML in their lifecycle stack runs from "a few predictive features in our ESP" to "nothing we've actually turned on". In between, a newer category of product is trying to sell brands something much closer to the frontier. The catch, and it's the catch the breakfast table kept circling, is that all of it depends on having your data in order, and most brands don't.
The thing the trade press calls "channel maturity" is more accurately described as sorting by AI sophistication. Some senders can afford to make a channel work. Most can't. And increasingly, the deciding factor isn't budget or headcount. It's whether your data is in a state that lets you do anything at all.
A handful of consumer technology companies publish peer-reviewed work on production notification and messaging systems. These are not white papers or vendor case studies. These are KDD, RecSys, WSDM and CIKM submissions, written by PhD-staffed teams who have to put their methodology in front of academic reviewers and answer for it. The papers aren't exhaustive descriptions of what the companies do internally (they never are), but they're a useful floor on capability. The internal systems are at least as good as what gets published, usually better, and the act of publishing signals that the team has the organisational backing to do this work seriously.
A quick tour of the visible frontier:
Pinterest set a weekly notification budget per user, optimising against long-term site engagement rather than click-through, on the finding that the incremental value of a notification is highest for casual users; the heavy openers have high click-through because they engage with everything, not because the notification moved them. 1
Duolingo used a bandit algorithm to pick which reminder template to send each user, and reported a 0.5% lift in daily active users and a 2% lift in new-user retention over a strong baseline. 2
Twitter used model-based reinforcement learning to decide whether to send a push at all, modelling the effect over a multi-day horizon. The published trade-off is the interesting part: the settings that cut volume hardest pushed open rate up by as much as 14%, but those same settings reduced daily active users; only the most conservative setting, an open-rate gain of about 8%, improved daily actives at all, and then by 0.2%. Maximising the headline number and serving the real objective pointed in opposite directions. 3
LinkedIn framed notification decisioning as offline reinforcement learning, a Double Deep Q-Network with a conservatism penalty, trained on logged data and deployed: sessions up a quarter of a percent, click-through up a couple of points, notification volume down , all at once. 4 By 2026 the same lineage had reached email: BanditLP pairs neural Thompson Sampling with a linear program large enough for billions of variables to choose, under business constraints, what each member is sent. 5
Zillow governs email and push volume with a boosted-tree classifier deciding send-or-don't per user, tuned to keep 98% of the clicks while shedding the surplus sends and the unsubscribes they cause. No reinforcement learning required, which is its own lesson: the cheapest method on this list still wins by sending less. 6
Meta treated Instagram's notification slots as an auction: the 550-plus internal teams that want to message you bid against each other (with the platform able to subsidise bids) so no single user is flooded by competing product surfaces. In test it sent slightly fewer notifications, lifted click-through and left reach untouched, across 77 million users per arm. 7
And the frontier keeps moving. PushGen , deployed at Kuaishou, the Chinese short-video platform, and presented at WSDM in February 2026, generates push copy with an LLM under style controls, then ranks the candidates with a learned reward model that predicts click-through and picks the winner, across hundreds of millions of users a day. 8 Pinterest's TransAct points the same way: a transformer reading a user's realtime activity, now feeding ranking across Homefeed, Search and Notifications, with push open rate and email click-through up a point or two each. 9
These systems share four characteristics:
Built on first-party event streams collected by the platform owner at full session-level granularity.
Operated by in-house engineering teams that include researchers, ML engineers and platform engineers in numbers most B2C brands could not staff if they tried.
Tuned against the platform's own long-term value functions, sessions, weekly actives, multi-day retention, sitewide engagement, rather than the open or the click.
Premised on the user's response being something that changes as a function of the messages sent , rather than a fixed signal to be exploited.
That last characteristic is the one almost everything sold to brands gets wrong. A static model asks which users are most likely to open or buy and aims at them. An adaptive one asks a different question: how does sending this, now, change what this user does, against the version of them you left alone? The response isn't a fixed trait you discover and exploit. It's an outcome the message itself moves, up for some people and down for others, which is why the same model has to be willing to send nothing at all. Optimise for who looks likely to engage and you keep messaging the people who were going to engage regardless, while quietly annoying the ones a message pushes the wrong way.
The lift on the real objective, sessions or daily actives, is almost always under a single percent; the bigger-looking gains sit on proxy metrics like open rate. The frontier's advantage isn't enormous lift. It's that a platform with hundreds of millions of users can bank a reliable 0.3% on sessions, which is a vast sum in absolute terms, while a mid-market brand chasing the same 0.3% on a fraction of the base can't justify the engineering to capture it. The gap is one of scale economics, not magnitude.
The gains nearly all come from sending less, or more precisely from sending differently . Volume falls while engagement holds or rises. The comparison is relative, mind: these platforms are trimming a per-user volume that already sits well above what most brands send, so "cutting volume" starts from a regime the long tail was never in. And even within that regime, "send less" is too blunt a reading, because these systems aren't cutting uniformly; they're reallocating, more to some users, none to others. Twitter's best setting actually raised the per-user send ceiling even as total sends dropped, because the policy got more selective about who was worth the headroom. The skill is knowing whom not to message, which is harder to learn than whom to message and needs exactly the per-user signal the long tail lacks. Uplift studies make the same point: the cumulative incremental effect peaks well before you've reached the whole list, usually far earlier than intuition suggests, and past that peak more targeting reduces it, because you start reaching people the message turns off rather than wins over. Send to the whole list and you are well into zero-or-negative territory.
What brands can actually buy: support versus decisioning
The distinction most often collapsed here is between machine learning that supports a marketer's decision and machine learning that makes it.
Most of the ML that lands in a brand's lifecycle stack is decision support . A smaller, newer category is decision making . Conflating the two is how you end up badly overestimating what the average brand is actually running.
Decision support is what the big customer engagement platforms ship:
Send-time and best-channel optimisation that pick the hour and the route for a message you already decided to send.
Predictive scores for churn, conversion, purchase propensity and lifetime value that you can build a segment around.
Predictive RFM and other "AI" segments that group users by recency, frequency, spend or likely next action.
Product recommendations off purchase history.
Subject-line suggestions and generative copy, the specialism Phrasee, now Jacquard, built on well before the LLM wave, generating brand-voice variants and ranking them on past performance.
Frequency capping and send-volume governing.
Most of these do something. Whether they do much is harder to say: whether send-time optimisation actually beats sending at a sensible fixed hour, for instance, rests almost entirely on vendor case studies, while the independent academic work that exists tends to favour fixed, scheduled times and to warn about notification overload. 10 But the common thread, lift or no lift, is that the marketer still designs the journey . The ML scores, suggests, ranks and optimises at specific points inside a structure the human built. It assists the decision. It doesn't make it.
This is the bracket almost every vendor sits in:
Salesforce Marketing Cloud's Einstein, with send-time, content selection and predictive scoring, now wrapped into the Agentforce branding.
Adobe CX Enterprise ( rebranded from Experience Cloud at Summit 2026 ), with Sensei, the AI Assistant, predictive audiences, automated send-time, content selection, generative variants, and the new tier of AI Coworkers and purpose-built agents.
Klaviyo's predictive analytics for CLV and churn, plus send-time, subject-line optimisation and generative copy.
Iterable's predictive goals, send-time optimisation, Brand Affinity and channel optimisation.
HubSpot's Breeze across the workflow.
Emarsys, now folded into SAP Engagement Cloud, with its AI Scores for conversion, churn and spend, continuously-updated predictive segments, Predict product recommendations, and a growing pile of generative helpers like natural-language catalogue search and an AI report builder.
Braze's Sage AI for send-time, channel, predictive churn and content suggestions.
The likes of CleverTap, MoEngage and Insider play here too, though they're more niche, strong in mobile-first and in particular regions rather than across the board.
If you read three of these vendors' marketing pages back to back you can't tell them apart. The vocabulary is identical and the screenshots are interchangeable. That isn't a coincidence. There's a finite set of things a multi-tenant ML feature can do well, and the vendors have all converged on roughly the same set: send-time, predictive segments, generative copy, journey-level orchestration with some scoring at the branch points. MessageGears documents the standard kit about as plainly as anyone, send-time and channel optimisation, p

[truncated]
