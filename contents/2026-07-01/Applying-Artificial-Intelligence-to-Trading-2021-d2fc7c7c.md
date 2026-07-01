---
source: "https://www.hudsonrivertrading.com/hrtbeat/ai-trading/"
hn_url: "https://news.ycombinator.com/item?id=48747685"
title: "Applying Artificial Intelligence to Trading (2021)"
article_title: "Applying Artificial Intelligence to Trading | Hudson River Trading"
author: "eustoria"
captured_at: "2026-07-01T14:58:08Z"
capture_tool: "hn-digest"
hn_id: 48747685
score: 1
comments: 0
posted_at: "2026-07-01T14:44:16Z"
tags:
  - hacker-news
  - translated
---

# Applying Artificial Intelligence to Trading (2021)

- HN: [48747685](https://news.ycombinator.com/item?id=48747685)
- Source: [www.hudsonrivertrading.com](https://www.hudsonrivertrading.com/hrtbeat/ai-trading/)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T14:44:16Z

## Translation

タイトル: 取引への人工知能の適用 (2021)
記事のタイトル: 取引への人工知能の適用 |ハドソン・リバー・トレーディング
説明: ここ数年、人工知能の普及から逃れるのは困難でした。携帯電話から車、冷蔵庫に至るまであらゆるものに AI が搭載されている可能性がありますが、実際に何が AI とみなされるかは人によって大きく異なります。の研究者でさえ、
[切り捨てられた]

記事本文:
人工知能を取引に応用する |ハドソン・リバー・トレーディング
ナビゲーションをスキップしてコンテンツにジャンプします
私たちは誰なのか
プログラマーによって構築され、プログラマーが主導します。 HRT では、数学者、コンピューター科学者、統計学者、物理学者、エンジニアが働いています。
世界中の 14 のオフィスにまたがる同僚のコミュニティ。
取引の未来を一緒に築きましょう。
流動性プロバイダーとして、HRT は顧客に最良の価格を提供するように設計された自動取引アルゴリズムを開発しています。
好奇心旺盛な思想家や熱心な自動化技術者が集まる活気に満ちたコミュニティに参加して、イノベーションの最前線で協力しましょう。
6 か国と世界のほぼすべての電子市場にわたって、優しさと卓越性が集まる場所です。
人工知能を取引に応用する
ここ数年、人工知能の普及から逃れるのは困難でした。携帯電話から車、冷蔵庫に至るまであらゆるものに AI が搭載されている可能性がありますが、実際に何が AI とみなされるかは人によって大きく異なります。 AI 分野の研究者ですら普遍的な定義を持っていないことが判明しましたが、（起源は不明瞭ですが）次のような定義があります。
AI は、周囲の環境を認識し、目標を達成できる可能性を最大化するアクションを実行するデバイスです。
HRT はアルゴリズム取引会社であり、モデルとコードが 1 日のすべての取引決定を (ナノ) 秒ごとに推進します。 AI のこの定義は私たちの取引の問題にどのように対応するのでしょうか?
まず、私たちの「環境」について考えてみましょう。 HRT の 1 つの見方は、価格発見 (証券が相互に相対的に適正な価格であることを保証する) や流動性の提供 (何かを買いたい、または売りたい場合、HRT は反対側に立つ準備ができている) などのサービスを世界に提供しているということです。提供するには、大量のデータを統合する必要があります。

これらのサービスを最低コストで利用できるため、ある意味「すべて」が私たちの環境なのです。データの最も重要な形式の 1 つは、株式、仮想通貨、商品、オプション、通貨などの為替取引商品の価格や出来高などの「市場データ」です。その後、ニュースレポート、企業提出書類などの「代替データ」、さらにはツイートや衛星データ、さらにはお気に入りの株取引サブレディットの人々の意見など、さらに珍しい形式の情報を追加することができます。
次に、「アクション」を定義する必要があります。これを正確にどのように定義するかは、おそらく達成したいことによって異なりますが、単純に検討している特定の証券や商品を購入、売却、または保有することにしましょう。
最後に、私たちの目標は何ですか?純トレーディング収益（取引の収益性から手数料などの取引コストを差し引いたもの）を最大化することは賢明ですが、その収益がいつ実現するかを考慮する必要があります。たとえば、「次の 1 時間の収益を最大化する」などと言ったら、次の 1 時間で大きなポジションを蓄積し、紙の上では利益があるように見えても、実際にはポジションが大きすぎるために売却して利益を実現できないなど、奇妙な効果が期待されるかもしれません。
これらすべてを最適化問題にまとめることができます。まだ議論していないことの 1 つは、トレーディング AI に対する制約です。たとえば、リスクを超えずにトレーディング収益を最大化したい場合があります (リスクにはさまざまな形態が考えられます)。私たちは取引する市場におけるいかなる規制や法律にも違反したくありませんが、さらに踏み込んでいきたいと考えています。一部の行為はいかなる法律にも違反しないかもしれませんが、私たちは善良な市場参加者でありたいため、特定の行動は避けたいと考えています。
結局のところ、PRを正確に定義するのは難しいのです

さて、目標を達成するために AI を実際に最適化する必要があります。その長い旅に着手する前に、AI 技術が成功している問題に問題をマッピングしてみることができます。あまりごまかさずにそのマップを作成できれば、同じテクニックを使用して問題を解決できるかもしれません。
AI の古典的な成功事例の 1 つはボード ゲームです。たとえば、IBM の Deep Blue (チェス用) や DeepMind の AlphaGo (囲碁用) などです。どちらの体制でも、AI は説得力のある超人間的なパフォーマンスを示しました (ただし、取引における人間のパフォーマンスはそれほど優れているわけではないため、超人であることは私たちにとってそれほどハードルではありません!)。ボードゲームにおける AI の成功の重要な側面のいくつかについて説明し、それらをトレーディングの問題と対比してみましょう。
2 人プレイの「ゼロサム」: チェス/囲碁は 2 人のプレイヤーによってプレイされ、1 人のプレイヤーだけが「勝つ」か引き分けます。つまり、一方のプレイヤーが勝つためにはもう一方のプレイヤーが負ける必要があり、私たちは両方とも勝ちたいと考えています。これは金融市場にはあまり当てはまりません。「現金」はほぼ保存されていますが、参加者にとっての効用や利益という点では市場はゼロサムではなく、これがなぜ誰もが取引するのかを部分的に説明しています。たとえば、401(k) に資金を投入するとき、最小限の約定コストで今すぐ株式を購入したいと考えていますが、40 年以上売却しないため、正確な価格にはやや鈍感です。私に売っている人は、私とはまったく関係のない取引によるリスクをヘッジするために売っているのかもしれません。私たち二人ともこの取引に満足できます。望んでいたものを両方とも手に入れたので、私たちは両方とも勝ちました。重要な意味の 1 つは、HRT の AI を他の「プレーヤー」の代理として使用できないということです。これは、AlphaGo の仕組みの重要なコンポーネントです。代わりに、さまざまなタイプの「プレーヤー」をモデル化してシミュレーションする必要があります。

実際の市場で起こるインタラクションを体験しました。
ルールを知る: AlphaGo は、ツリー検索手順でゲーム ルールの完全な知識を活用しており、すべてのアクションの即時結果は完全に既知であり、決定的です。次に、上で説明したように、他のプレイヤーのその後の反応は、私たちが行うことを行うだけで適切にモデル化されます。トレーディングにおいて、私たちは部分的にこの特性を持っています。つまり、全国株式市場のようなより発展した市場における市場の動きは通常明確に定義されていますが、注文を送信した場合に何が起こるかを正確に知る特性はありません。この簡単な例としては、別の「プレイヤー」がまったく同時に注文を送信した場合、結果はランダムになります。
状態を完全に観察する: チェスや囲碁などのボード ゲームでは、特殊な場合を除いて、世界の過去の状態を知る必要はありません。知る必要があるものはすべてゲーム ボード上に表示されます。市場では、未来を予測するには過去が非常に重要です。簡単な例としては、回帰という考え方があります。株式の価格が突然変動した場合、長期的な平均価格に「回帰」する可能性が高くなります。過去は重要ですが、どの瞬間にも多くの隠された情報が含まれています。他のすべての市場参加者の身元と立場は明白ですが、細かいレベルでは、取引するまで表示されない「隠された」注文タイプを考慮してください。
無限のデータ: 私たちは、自分自身または別の AI と好きなだけチェスのゲームをプレイできます。これらのゲームはそれぞれ「本物」ですが、唯一の欠点は、生成されたデータセットの多様性が、ゲームをプレイするすべての可能な方法を捉えていない可能性があることです。対照的に、取引では現実は 1 つだけであり、新しいデータをリアルタイムで、つまり非常にゆっくりと取得します。多くの差分をシミュレートすることでデータの生成を試みることができます。

AI は互いに激しく対立しますが、上で説明した理由により、現実世界の豊かな複雑さと相互作用の代理としては不十分です。 AlphaGo の勝利 AI で使用された大規模なニューラル ネットワークをトレーニングするには、大量のデータを保持することが重要です。
これらすべてを踏まえると、AI を取引に適用することはできないと考えたくなるかもしれません。障害が多すぎます。ただし、これらの問題を回避する方法はたくさんあります。最良の方法の 1 つは、全体構造に対する人間の理解を利用して、大きな問題を小さなサブ問題に分解することです。これらのサブ問題の解決策を組み合わせると、元の質問に対する良い、しかし不完全な解決策が得られます。そのギャップを埋めることに取り組むことは、研究者が積極的に取り組んでいる問題の 1 つです。
これをトレードス​​タイルに当てはめて分解する方法を考えてみましょう。この問題では、ポートフォリオの総リスクの制限を条件として、複数日間にわたる純トレーディング利益を最大化することが目的です。私たちは、価格が上昇すると思われる銘柄を保有することと、ポジションのリスクを軽減することとの間のバランスを維持するために、ポートフォリオ内の各証券を売買することを毎分選択できます。当然のことながら、この問題には、どの証券の価格が上昇または下降するかの特定から、リスクの管理、売買の最適な時期の決定まで、AI にとって困難な側面が数多くあります。
1 つの AI にすべてを実行できるように訓練しようとするのではなく、数学と直感を使用して、これを 3 つのサブ問題に分割できます。
予測。この AI の仕事は、単に未来を予測することです。たとえば、「世界の観察を考慮して、今後 5 日間の終了時の各証券の価格を予測してください」と言うことができます。これはまだ非常に難しいです

問題はありますが、元のものよりもはるかに簡単です。 AI は、他の市場参加者がどのように反応するかを心配する必要はなく、さらには行動を決定する必要もありません。代わりに、この「教師あり」学習問題を解決する必要があるだけです。私たちは、データから「シグナル」を手作りしてそれらを組み合わせることから、複雑なブラックボックス モデルに可能な限りすべてのデータを使用させることまで、この問題に対処するための多くのテクニックを持っています。私たちの研究者の多くは、この副問題にほとんどの時間を費やしています。
最適化 。この AI の仕事は、最初の AI の予測を取得し、それを各証券への投資可能な資金の全体的な配分に変換することです。価値が上がると思われる証券により多くの資金を投入することになるが、その資金を分散させようとするため、1つの企業にオールインすることはなくなり、よりリスクの高い立場となる。この問題に対する金融数学にはすでに多くの解決策があるため、必ずしも AI の観点を考慮する必要はありませんが、ソリューションをさらに洗練させる余地は十分にあります。
実行。最適化によるターゲット割り当てを考えると、そのターゲットを達成するために実際にどのように売買すればよいでしょうか?たとえば、午前 10 時に AI が Apple 株の価値が今日よりも明日の方が上がると確信しており、リスクを調整した後、次の 5 時間でさらに 9,000 株を購入したいと確信しているとします。一度に全部買ったほうがいいのでしょうか？目標は確実に達成できますが、現時点では販売する人が不足している可能性があり、高い値段を支払う可能性があります。毎分 30 株を購入すれば間違いなく十分な量になりますが、非常に短期的な予測に基づいて、一部の分は他の分よりも価格が高くなる可能性があります。しかし、数分以内に買わなければ、後で追いつく必要があります。これらすべての懸念をトレードオフするのは、

複雑な問題ですが、すべてをまとめるよりは簡単で、AI に学習させてみることができる問題です。それでも、この問題は、上で説明した問題に直面するため、多くの意味で最も困難です。ゲームとは異なり、私たちの行動は、理由を説明するのが難しい即時の結果を引き起こす可能性があります。
HRT はこれらの問題に長年取り組んできており、今後もさらに多くの問題に取り組む予定です。これらは内容が豊富で、複雑で、難しい考慮事項であり、解決策を見つけるために読める本や学術論文はありません。代わりに、研究者は関連分野からの洞察、自らの直観と創造性、そして厳密な科学的アプローチを活用して進歩を遂げなければなりません。興味があれば、私たちのチームへの参加を申請してください!
株式リターンのモデリング: 線形ケース
トレーディングでは、機械学習ベンチマークはあなたが関心を持っていることを追跡しません
Hudson River Trading は、グローバル市場で流動性をクライアントに直接提供するマルチアセットクラスの定量取引会社です。当社は、研究開発、モデリング、リスク管理のための世界で最も先進的なコンピューティング環境の 1 つを構築し、あらゆる金融市場の技術革新の最前線に立っています。私たちは思慮深い市場リーダーであり、世界市場の健全性と長寿に尽力しています。
3 世界貿易センター
175 グリニッジ ストリート、76 階
ネ

[切り捨てられた]

## Original Extract

In the past few years it has been hard to escape the ubiquity of Artificial Intelligence. Although everything from your phone to your car or even your refrigerator might be endowed with AI, what is actually considered AI will vary widely from person to person. It turns out that even researchers in t
[truncated]

Applying Artificial Intelligence to Trading | Hudson River Trading
Skip navigation and jump to content
Who We Are
Built by coders, led by coders. At HRT we are mathematicians, computer scientists, statisticians, physicists, and engineers.
A community of colleagues spanning 14 offices worldwide.
Building the future of trading together.
As a liquidity provider, HRT develops automated trading algorithms designed to provide the best prices to our clients.
Join our vibrant community of curious thinkers and eager automators, working together at the forefront of innovation.
A place where kindness and excellence converge across six countries and nearly all the world’s electronic markets.
Applying Artificial Intelligence to Trading
In the past few years it has been hard to escape the ubiquity of Artificial Intelligence. Although everything from your phone to your car or even your refrigerator might be endowed with AI, what is actually considered AI will vary widely from person to person. It turns out that even researchers in the field of AI don’t have a universal definition, but one (of unclear origin) is:
AI is any device that perceives its environment and takes actions that maximize its chance of successfully achieving its goals.
HRT is an algorithmic trading company, where models and code drive all our trading decisions every (nano)second of the day. How does this definition of AI map to our trading problems?
First, let’s consider our “environment.” One view of HRT is that it provides services to the world like price discovery (ensuring securities are fairly priced relative to each other) and liquidity provision (if you want to buy or sell something, HRT stands ready to take the other side). We need to integrate large amounts of data in order to provide these services at lowest cost, and so in some sense “everything” is our environment! One of the most integral forms of data is “market data” like prices and volumes for exchange-traded products like stocks, cryptocurrencies, commodities, options, and currencies. We can then add “alternative data” like news reports, company filings, and even more exotic forms of information like Tweets, satellite data – even the opinions of people on your favorite stock-trading subreddits.
Next, we need to define our “actions.” Exactly how we define this is probably going to be dependent on exactly what we want to achieve, but let’s simply go with buy, sell, or hold any given security or instrument we are considering.
Finally, what is our goal? Maximizing the net trading revenue (profitability of our trades less trading costs like commissions) would be sensible – although, we need to consider when that revenue will be realized. For example, if we said something like “maximize revenue over the next hour” then we might expect some strange effects like accumulating a large position over the next hour which looks profitable on paper but can’t actually be sold to realize the profit because the position is too large.
We can put this all together into an optimization problem. One thing we haven’t yet discussed is the constraints on our trading AI. For example, we might want to maximize our trading revenue but without exceeding risk (which can take many forms). We definitely don’t want to violate any regulations or laws in the markets we trade in, but we want to go even further: some actions may not violate the letter of any law, but we want to be good market participants, so want to avoid certain behaviours.
After all the difficulty of precisely defining our problem, now we have to actually optimize our AI to achieve its goals. Before we embark on that long journey, we can try to map our problem onto one where AI techniques have been successful. If we can make that map without too much fudging, then we might be able to use the same techniques to solve our problem.
One of the classic success stories of AI has been board games; for example, IBM’s Deep Blue (for Chess) and DeepMind’s AlphaGo (for Go). In both regimes AI demonstrated convincingly super-human performance (although human performance at trading isn’t very good, so super-human isn’t much of a bar for us!). Let’s discuss some of the key aspects of the success of AI for board games, and contrast them to the trading problem:
Two-player “zero sum” : Chess/Go are played by two players, and only one player can “win” or they draw. That is, for one player to win, the other player must lose, and we both want to win. This doesn’t map very well to financial markets: while “cash” is roughly conserved, markets aren’t zero sum in terms of the utility or benefit to the participants, which partially explains why anyone trades at all. For example, when I put money into my 401(k), I want to purchase stock now at minimal execution cost, but I’m somewhat insensitive to the exact price since I won’t be selling for 40+ years. The person selling to me might be selling to hedge some risk they have from some completely unrelated-to-me trade. Both of us can be happy with the trade – we both win because we both got what we wanted. One crucial implication is that we cannot use HRT’s AI as a proxy for the other “players”, which is a key component of how AlphaGo works. Instead, we would need to model a wide variety of different types of “players” to simulate the interactions that happen in real markets.
Knowing the rules : AlphaGo makes use of full knowledge of the game rules in its tree search procedure, and the immediate outcome of every action is fully known and deterministic. Then, the other player’s subsequent response is well-modelled by just doing what we’d do, as we described above. In trading we partially have this property: market behaviour in more developed markets like national stock markets is usually well-defined, but we do not have the property of knowing exactly what will happen if we send an order. A simple example of this is if another “player” sends their order at the exact same time, the outcome is random.
Fully observing state: Board games like chess and Go don’t require knowing past states of the world, outside of some corner cases – everything you need to know is visible on the game board. In markets, the past is critical for forecasting the future. A simple example is the idea of reversion: if the price of a stock moves suddenly, it is more likely than not to “revert” to its long-run average price. The past is important, but any given instant holds a lot of hidden information: the identities and positions of all other market participants is an obvious one, but at a fine-grained level, consider “hidden” order types that don’t appear until you trade with one.
Infinite data : We can play as many games of chess against ourselves, or a different AI, as we want. Each of those games is “real” and the only shortcoming is that the diversity of that generated dataset might not capture all possible ways of playing the game. In contrast, in trading we have only one reality, and acquire new data in real time – i.e., very slowly. We can attempt to produce data by simulating many different AIs against each other, but it is a poor proxy for the rich complexities and interactions of the real world, for reasons described above. Having a large amount of data is critical for training the large neural networks that were used by AlphaGo’s winning AI.
Based on all that, one might be tempted to think we can’t apply AI to trading – there are just too many obstacles! However, there are plenty of ways to work around these problems. One of the best is to decompose the big problems into smaller subproblems using our human understanding of the overall structure. If we put the solutions to these subproblems together, we can get a good but imperfect solution to the original question, and working on closing that gap is one of the problems our researchers actively work on.
Let’s work through applying this to a style of trading and think about how to decompose it. In this problem our objective is to maximize our net trading profit over multiple days, subject to a limit on the total risk of our portfolio. Every minute, we can choose to buy or sell each of the securities in our portfolio to maintain the balance between owning those we think will increase in price, and reducing the risks of our position. Naturally, there are many aspects of this problem that are hard for an AI, from identifying which securities will increase or decrease in price, to managing risk, to determining the best time to buy or sell them.
Rather than trying to train one AI to do it all, we can use mathematics and intuition to break this up into three subproblems:
Prediction . The job of this AI is to simply predict the future! For example, we could say “given the observations of the world, predict the price of each security at the end of the next five days.” This is still a very hard problem, but much simpler than the original. The AI doesn’t need to worry about how other market participants will react, or even decide on an action. Instead, it just needs to solve this “supervised” learning problem. We have many techniques to attack this problem, from hand-crafting “signals” from data and combining them, to letting a complex black-box model use all the data any way it can. Many of our researchers spend most of their time on this subproblem.
Optimization . The job of this AI is to take the predictions of the first AI and convert them to an overall allocation of our investable money to each of the securities. It will put more money into securities that we think will increase in value, but will attempt to spread the money around so we don’t go all-in on one company – a much riskier position. Financial mathematics has many solutions to this problem already, so we don’t have to take an AI perspective here at all necessarily, but there is plenty of room here to make our solutions more and more sophisticated.
Execution . Given the target allocation from our optimization, how do we actually buy or sell to reach that target? For example, let us suppose our AI is sure at 10am that Apple stock will be worth more tomorrow than today, and after adjusting for risk it is sure it wants to buy an additional 9,000 shares over the next 5 hours. Should it buy them all at once? You’d be sure to reach the target, but there might not be enough people selling right now, and we might pay a high price. If we buy 30 shares every minute, we’d definitely get enough, but maybe some minutes have better prices than others – based on a very short-term forecast. But if we don’t buy in some minutes, we have to catch up later. Trading off all these concerns is a complex problem, but simpler than taking them all together, and one we can try to train an AI to do. Still, this problem is in many ways the hardest as it hits the issues we described above – our actions can have immediate consequences that are hard to reason about, unlike in a game.
HRT has been working on these problems for many years, and will be working on them for many more to come. These are rich, complex, and difficult considerations, and there is no book or academic paper that you can read to find their solutions. Instead, our researchers must draw on insights from related fields, their own intuition and creativity, and a rigorous scientific approach to make progress. If that sounds interesting to you, apply to join our team!
Modeling Equities Returns: The Linear Case
In Trading, Machine Learning Benchmarks Don’t Track What You Care About
Hudson River Trading is a multi-asset class quantitative trading firm that provides liquidity on global markets and directly to our clients. We have built one of the world’s most advanced computing environments for research and development, modeling, and risk management, and are at the forefront of technical innovation for financial markets everywhere. We are thoughtful market leaders, committed to the health and longevity of global markets.
3 World Trade Center
175 Greenwich Street, 76th Floor
Ne

[truncated]
