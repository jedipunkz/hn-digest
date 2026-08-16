---
source: "https://philippdubach.com/posts/kimi-k3-inside-claude-code/"
hn_url: "https://news.ycombinator.com/item?id=49319610"
title: "Testing Moonshot AI's Kimi K3 Inside Claude Code"
article_title: "Testing Moonshot AI's Kimi K3 Inside Claude Code - philippdubach.com"
author: "7777777phil"
captured_at: "2026-08-16T13:23:54Z"
capture_tool: "hn-digest"
hn_id: 49319610
score: 3
comments: 0
posted_at: "2026-08-16T12:54:31Z"
tags:
  - hacker-news
  - translated
---

# Testing Moonshot AI's Kimi K3 Inside Claude Code

- HN: [49319610](https://news.ycombinator.com/item?id=49319610)
- Source: [philippdubach.com](https://philippdubach.com/posts/kimi-k3-inside-claude-code/)
- Score: 3
- Comments: 0
- Posted: 2026-08-16T12:54:31Z

## Translation

タイトル: Moonshot AI の Kim K3 Inside Claude コードのテスト
記事のタイトル: Moonshot AI の Kim K3 Inside Claude コードのテスト - philippdubach.com
説明: Moonshot AI を介してクロード コードをルーティングしました

記事本文:
Moonshot AI の Kim K3 Inside Claude コードのテスト - philippdubach.com メイン コンテンツにスキップ philippdubach philippdubach 定量的金融、機械学習、および複雑なシステム。
キミK3インサイドクロードコードをやってみた
Claude Code の内部では、Kimi K3 が十分に優れていたため、どのモデルがインターフェイスの背後にあるのかについてすぐに考えるのをやめました。
モデルは、測定された出力速度が示唆するよりも遅く感じられましたが、詳細なプロンプトとアセット パックから近いフロントエンドの再構築が生成されました。
私のエクスポートしたランの費用は 7.18 ドルでした。同じ記録されたトークン ミックスの価格は、公表されている定価で、Claude Opus 4.8 では約 11.96 ドル、Claude Fable 5 では約 23.92 ドルになります。
より大きなチャンスはソブリン AI である可能性があります。オープン ウェイトはローカル インフラストラクチャでホストされ、トークンのみではなく常駐性、制御性、監査可能性に基づいて販売されます。
× Kim K3 は、ネイティブ ビジョンと 100 万トークンのコンテキスト ウィンドウを備えた、Moonshot の新しい 2.8 兆パラメータ モデルです。 Moonshot はこれをこのスケールで初のオープンモデルと呼び、重量は 7 月 27 日までにリリースされる予定であると述べています。 同社自身の打ち上げ記事も、いつになく率直です。K3 は全体的には依然として Fable 5 と GPT-5.6 Sol に後れをとっており、そのユーザー エクスペリエンスはまだ同じレベルに達していないと述べています。
私がこれを書いている間に、Alibaba は Qwen 3.8 を発表しました。これは、オープンウェイト リリースに向けたもう 1 つの非常に大きなモデルです。これは、Thinking Machines Lab が、410 億のアクティブ パラメーターとフル ウェイトを使用できる 9,750 億パラメーターの米国モデルである Inkling をリリースした直後でもあります。キミはもう一回限りのものではありません。現在行列ができています。
私はこれまでにクロード コードを通じて数百万のトークンを焼き付けてきました。私の設定は機能し、ツールは期待どおりの場所にあり、ハーネスのリズムもわかっています。
そこで、Claude コードを保持し、 OpenRouter を使用してリクエストを Kim K3 にルーティングしました。
×
遅く感じました。非常に遅い、f

まず。測定された数値がそうではないことを示しているため、これは奇妙です。 Artificial Analysis は、Kimi K3 については 1 秒あたり約 62 の出力トークン、Claude Opus 4.8 については 57 の出力トークンを報告します。また、K3 の最初のトークンまでの時間も大幅に短縮されます。
おそらく違いは、有用なテキストが表示されるまでに推論に費やした時間でした。おそらくそれは流れのリズムだったのでしょう。もしかしたら、私が単にクロード コードを長い間使用してきただけで、リズムが違うと違和感を感じてしまうのかもしれません。 (あるいは単に openrouter を経由していたからかもしれません。)
それから、最初の1時間くらいのどこかで、私は気づかなくなりました。 K3が急に速くなったからではありません。ファイル、編集、ツールの呼び出しをうまく処理してくれたので、トラフィックがどこに向かっているのか忘れてしまいました。
K3 はフロントエンドの仕事で注目を集めているので (現在 Code Arena で 1 位にランクされています)、別のコーディング ベンチマークではなく視覚的なものが必要でした。この Lafys ビルドのプロンプトとアセットを再利用しました。私はそれを確認することはできませんが、参照は Fable 5 で行われたと思います。
×
×
×
× 私は、彼が非常に具体的な指示とオリジナルのアセットを使って 1 回実行することを知っています。モデルに曖昧なプロンプトを与えると、結果がバラバラになる可能性があります。
この実行の OpenRouter エクスポートには、115 個のリクエスト、1,364 万個のプロンプト トークン、および 82,307 個の出力トークンが含まれています。プロンプト トークンのうち、1,295 万がキャッシュ ヒットでした。総費用: 7.18 ドル。
Anthropic の公表されている定価をその正確なトラフィック ミックスに適用すると、Claude Opus 4.8 が約 11.96 ドル、Claude Fable 5 が 23.92 ドルになります。この実行では、Kimi は Opus より約 40%、Fable より 70% 安くなりました。
作業が困難になったときは、オーパスの方がまだ信頼できると感じます。いつやめるべきかをよりよく理解し、曖昧さを処理するのが上手で、自分が求めていない精力的な決定を下す可能性が低くなります。 Moonshot 自体は過剰な積極性をリストしています。

K3 の制限の中で感度を利用することは、私の経験の一部と一致します。
K3もFable 5クラスモデルとは言いません。
しかし、フェイブルとオーパスの下では、空気はすぐに薄くなります。 K3 は十分近いので、多くのユーザーは残りの差額を気にしないか、価格よりも気にするでしょう。
ここ数日間の影響はあらゆる方向に進んでいます。
キミ K3 はアメリカンモデルの堀の終わりです。キミK3は劇場のベンチマークです。オープンウェイトによりクローズドモデルは時代遅れになります。運営に小規模なデータセンターが必要な場合、オープンウェイトは意味がありません。中国は現在米国市場の大部分が賭けている層を意図的にコモディティ化している。あるいは、基本計画がなく、良いモデルは開発者、使用法、名声をもたらすため、中国の研究所が良いモデルをリリースしている。
あまり劇的ではないリスト (私のメモから) は次のようになります。
モデルが非常に異なる数のトークンを使用する場合、$/token は不適切な尺度です。
公開ベンチマークは最適化が容易ですが、上位モデルを分離するのは困難です。
オープンウェイトは、ほとんどのユーザーがセルフホストをしない場合でも API 価格に上限を設けます。
モデルは置き換え可能になりつつあります。ハーネス、データ、ワークフローはそうでない可能性があります。
価値はアプリケーションに移行し、チップ、電源、推論インフラストラクチャに移行します。
オープンウェイトは制御とプライバシーに役立ちますが、「オープン」だからといってランニングコストが安いわけではありません。
輸出規制は中国の研究所の活動を遅らせる可能性があると同時に、米国のスタックを中心に構築する理由を与える可能性がある。
最大の脅威は、必ずしもキミがナンバーワンになることではない。それは、長くナンバーワンであり続ける人はいないということです。
Epoch AI は、最高のオープンウェイト モデルがクローズド フロンティアにわずか数か月遅れていると推定しています。正確なギャップはベンチマークとともに変動するため、世間の評価は注意深く読む必要がありますが、方向性を判断するのは困難です

逃す。ラグは十分に短いため、購入者は待機したり、プロバイダー間でワークロードを切り替えたり、分割したりできます。個人的には、これは、はい、わかりますが、脅威と機会に蒸留できると思います。
脅威は、キミ K3 がすべてのアメリカのモデルよりも優れているということではありません。そうではない。
脅威は、プレミアムを守るのが難しくなるのに十分である可能性があることです。
米国の大規模研究所は技術的能力以上のものを重視しています。この財務ケースでは、フロンティアのインテリジェンスが不足し続け、リードが持続し、顧客がアクセスに対して高いマージンを支払い続けていることが前提となっています。モデルは、最初の地位に就かなくても、そのストーリーにダメージを与える可能性があります。必要なのは、フロンティアに十分近く、低価格で、既存のソフトウェアに簡単に組み込めることだけです。
これは価格決定力にとって悪い組み合わせです。
ほとんどのエンタープライズ ソフトウェアは切り替えがひどいため、不安定です。データベース、オペレーティング システム、およびコア バンキング プラットフォームには、長年にわたる依存関係が蓄積されています。 API の背後にある言語モデルは異なります。このモデルは非常に重要であると同時に、驚くほど簡単に置き換えることができます。私のクロード コード テストは、その小規模バージョンです。
閉鎖された研究室でもプレミアムを獲得する方法はまだあります。信頼性が重要です。ツールの使用が重要です。企業の管理は重要です。タスクが十分に価値のあるものである場合、6 か月前に進めることには大きな価値があります。
インフラ費用が莫大なので、タイミングが難しい。これについてはすでに「AI 設備投資の軍拡競争: 誰が最初にまばたきするか?」で書きました。 。そこでの問題は、AI が価値を生み出すかどうかではなく、データセンター、GPU、すでにコミットされている負債と減価償却をサポートするために、どれだけの収益をどれだけ早く得なければならないかということでした。
収益の増加よりも早くモデルの機能がコモディティ化すると、被害はラボにとどまりません。ハイパースケーラー、チップメーカー、データセンター開発者、公益事業者、およびその他の企業を通じて実行されます。

建設に資金を提供します。より安価なキミはユーザーにとっては良いことですが、希少性を引き受けた人にとっては潜在的に悪影響を及ぼします。
次に政策の問題がある。
米国は、閉鎖モデルのフロンティアの多くとその下のハードウェアの多くを管理している。これはレバレッジですが、それは世界の残りの部分が同じスタックに依存している間だけです。チップと API を制限すると、リードを維持できる可能性がありますが、並列スタックの価値が高まる可能性もあります。
この裏側については、『クルーグマン』、『寓話 5』、そして『衰退するヨーロッパ』で書きました。 。クローズド モデルは、プロバイダーまたはプロバイダーの政府によって廃止される可能性があります。オープンウェイトが一度広がると、同じ方法でスイッチをオフにすることはできません。
もう 1 つの、よりありふれたリスクがあります。それは、モデルを過大評価している可能性があることです。
すべてのリリースには、ベンチマークの勝利に関する高密度のページが付属しています。いくつかは本物です。異なるハーネス、異なる予算、またはモデル固有のチューニングを使用する場合もあります。一部のベンチマークは単に飽和状態になっています。リーダーボードでは優秀に見えるモデルでも、長いセッションでは不快な場合があります。
価格が安くなると、より多くのタスクに挑戦する価値があります。エージェントは、より多くのパスを取得し、より多くのコンテキストを読み取り、有用ではあるが Fable の価格設定を正当化するほど重要ではない作業を処理できます。単価が下がっても市場はまったく縮小しない可能性があります。
もっと興味深い機会は別居です。インターフェイス、エージェント、モデルは同じ会社のものである必要はありません。
定期的なリファクタリングでは、より安価なモデルを使用できます。難しいアーキテクチャの決定は Opus に任せることができます。機密性の高いマテリアルはプライベート展開に残すことができます。フロントエンドの作業は、その月にたまたま最も得意なモデルに移ることができます。耐久性のある製品は、単一のモデルではなく、ルーター、コンテキスト層、検証ループである可能性があります。
これは、オープンウェイトが価格以上に重要な点でもあります。組織を微調整し、定量化できるようにします。

独自の制御下でモデルをデプロイします。これにより、結果が自動的に安全になったり準拠したりするわけではありませんが、データの場所、保存、アクセス、継続性に関する難しい質問に誰が答えることができるかが変わります。
私がこれについてしばらく考えていたもう 1 つのことは、「DORA はどのように主権を銀行問題にしたのか」の中で、データ主権が業務に浸透していると主張しました。銀行は現在、集中、撤退計画、監査権、そして重要な外国プロバイダーがルールを変更した場合に何が起こるかについて考えなければなりません。
AI も同様の扱いを受けるでしょう (または実際に受けます)。最近まで、それには大きな能力上のペナルティが伴いました。新しいリリースでは、計算がそれほど愚かではなくなりました。キミ K3 は巨大で、その重量はまもなく測定される予定です。アリババはクウェン家を同じ方向に推し進めている。次に、Inkling が重要です (おそらくさらに重要です)。これは、利用可能な完全な重み、100 万トークンのコンテキスト、およびカスタマイズへの明確な焦点を備えた、米国でトレーニングされた本格的なモデルを追加するためです。
それでは、裏のビジネス ケースを考えてみましょう。スイスのプロバイダーがチューリッヒで 1 つの 64 GPU クラスターを実行し、専用またはマネージド K3 クラスの推論を銀行、製薬会社、政府機関、およびスイスのデータ常駐に関心のあるその他の顧客に販売していると仮定します。
× このモデルでは、初期設備投資が約 700 万ドル、月々の運用コストが 370,000 ドル、使用率が 70%、月額 15,000 スイスフランから 50,000 スイスフランのサブスクリプションを使用します。これらの仮定に基づくと、1 つのクラスターの年間収益は約 740 万スイスフラン、EBITDA は 300 万スイスフランに達し、3 年間で投資回収可能となります。
この製品は実際にはトークンではありません。それは管理です。スイス居住、定義されたセキュリティ境界線、顧客データに関するトレーニングなし、監査アクセス、契約の継続性、オペレーターが失敗した場合の退出ルートです。ほとんどの会社はw

文字通りモデルを地下室に置いたわけではありません。彼らは機能的に同等のものを提供してもらうために地元の誰かにお金を払うかもしれません。
これが機能しない理由はいくつかあります。 2.8 兆のパラメータを持つモデルを提供するにはコストがかかります。 1 つのクラスター自体が集中リスクになります。ハードウェアはすぐに老朽化します。使用率が低下すると、経済性が悪化します。はるかに小型のモデルが来年もほぼ同等のパフォーマンスを発揮すれば、高価なラックは昨年のベンチマークの記念碑となるでしょう。それでもこの市場は存在すると思います。
私はキミ K3 が人類を殺すとは思いません。 1 つのフロントエンドの再構築が同等であることを証明するとは思えません。そして、私は今でも最も難しいコーディング作業には Opus を選ぶでしょう。
でも、K3は使っていることを忘れてしまうほど良かったです。実質的にコストが安くなります。すでに気に入っていたハーネスの中で走りました。間もなく、ハードウェアを提供したい人は誰でもその重みを利用できるようになるはずです。
問題はもはや、オープン モデルが抽象的な未来のフロンティアを捉えるかどうかではありません。クローズドラボがどのくらいのリードを必要とするか、どのくらいの期間それを保持できるか、そして切り替えが日常的になったとしても、どの顧客が追加ポイントごとに支払い続けるかが重要です。
私は AI、主権、そしてその根底にある経済学について数週間ごとに書いています。次のものを入手してください。
追跡はありません。いつでも購読解除

[切り捨てられた]

## Original Extract

I routed Claude Code through Moonshot AI

Testing Moonshot AI's Kimi K3 Inside Claude Code - philippdubach.com Skip to main content philippdubach philippdubach Quantitative Finance, Machine Learning, and Complex Systems.
I Tried Kimi K3 Inside Claude Code
Inside Claude Code, Kimi K3 was good enough that I soon stopped thinking about which model was behind the interface.
The model felt slower than its measured output speed suggests, but it produced a close frontend reconstruction from a detailed prompt and asset pack.
My exported run cost $7.18. The same recorded token mix would have cost about $11.96 on Claude Opus 4.8 and $23.92 on Claude Fable 5 at published list prices.
The larger opportunity may be sovereign AI: open weights hosted on local infrastructure and sold on residency, control and auditability rather than on tokens alone.
× Kimi K3 is Moonshot’s new 2.8-trillion-parameter model , with native vision and a one-million-token context window. Moonshot calls it the first open model at this scale and says the weights will be released by July 27. Its own launch post is also unusually candid: it says K3 still trails Fable 5 and GPT-5.6 Sol overall, and that its user experience is not yet at the same level.
While I was writing this, Alibaba announced Qwen 3.8 , another very large model headed for an open-weight release. This is also shortly after Thinking Machines Lab released Inkling , a 975-billion-parameter US model with 41 billion active parameters and full weights available. Kimi is not a one-off anymore. There is now a queue.
I have burned millions of tokens through Claude Code by now. My config works, the tools are where I expect them to be, and I know rhythm of the harness.
So I kept Claude Code and using OpenRouter , I routed the requests to Kimi K3.
×
It felt slow. Very slow, at first. Which is odd, because the measured numbers say otherwise. Artificial Analysis reports about 62 output tokens per second for Kimi K3 and 57 for Claude Opus 4.8. It also measures a much shorter time to first token for K3.
Maybe the difference was time spent reasoning before useful text appeared. Maybe it was the cadence of the stream. Maybe I have simply used Claude Code for long enough that anything with a different rhythm feels wrong. ( Or maybe it was just because I was going through openrouter. )
Then, somewhere during the first hour, I stopped noticing. Not because K3 suddenly became faster. It handled files, edits and tool calls well enough that I forgot where the traffic was going.
K3 has been getting attention for frontend work ( ranked #1 on Code Arena right now), so I wanted something visual rather than another coding benchmark. I reused the prompt and assets from this Lafys build . I assume the reference was made with Fable 5, although I cannot verify that.
×
×
×
× I know his is one run, with very specific instructions and the original assets. Give the model a vaguer prompt and the result may fall apart.
The OpenRouter export for this run contains 115 requests, 13.64 million prompt tokens and 82,307 output tokens. Of the prompt tokens, 12.95 million were cache hits. Total cost: $7.18.
Applying Anthropic’s published list prices to that exact traffic mix gives about $11.96 for Claude Opus 4.8 and $23.92 for Claude Fable 5. In this run, then, Kimi came out roughly 40% cheaper than Opus and 70% cheaper than Fable.
Opus still feels more dependable when the work gets difficult. It is better at knowing when to stop, better at handling ambiguity, and less likely to make an energetic decision I did not ask for. Moonshot itself lists excessive proactiveness and harness sensitivity among K3’s limitations , which matches parts of my experience.
I would not call K3 a Fable 5-class model either.
But below Fable and Opus the air gets thin very quickly. K3 is close enough that many users will not care about the remaining difference, or will care less than they care about the price.
The takes over the past few days have gone in every direction.
Kimi K3 is the end of the American model moat. Kimi K3 is benchmark theatre. Open weights make closed models obsolete. Open weights are irrelevant if running the thing costs a small data centre. China is deliberately commoditising the layer on which a large part of the US market is now betting. Or there is no master plan and Chinese labs are releasing good models because good models bring developers, usage and prestige.
A less dramatic list (from my notes) would look like this:
$/token is a bad measure when models use very different numbers of tokens.
Public benchmarks are easier to optimise for and worse at separating the top models.
Open weights put a ceiling on API prices even if most users never self-host.
The model is becoming replaceable; the harness, data and workflow may not be.
Value moves up into applications and down into chips, power and inference infrastructure.
Open weights help with control and privacy, but “open” does not mean cheap to run.
Export controls can slow Chinese labs and, at the same time, give them a reason to build around the US stack.
The biggest threat is not necessarily that Kimi becomes number one. It is that nobody stays number one for long.
Epoch AI estimates that the best open-weight models trail the closed frontier by only a few months. Exact gaps move with the benchmark, and public evaluations should be read with care, but the direction is hard to miss. The lag is short enough that buyers can wait, switch or split workloads across providers. Personally, I think this can be distilled, yes, I know, into a threat and an opportunity.
The threat is not that Kimi K3 is better than every American model. It is not.
The threat is that it may be good enough to make the premium harder to defend.
The large US labs are valued on more than technical competence. The financial case assumes that frontier intelligence stays scarce, that the lead lasts, and that customers keep paying high margins for access. A model can damage that story without taking first place. It only needs to sit close enough to the frontier, at a lower price, and be easy to slot into existing software.
That is a bad combination for pricing power.
Most enterprise software is sticky because switching is awful. Databases, operating systems and core banking platforms accumulate years of dependencies. A language model behind an API is different. The model may be deeply important while still being surprisingly easy to replace. My Claude Code test is the small version of that..
Closed labs still have ways to earn the premium. Reliability matters. Tool use matters. Enterprise controls matter. Being six months ahead can be worth a lot when the task is valuable enough.
The timing is awkward because the infrastructure bill is enormous. I already wrote about this in AI Capex Arms Race: Who Blinks First? . The issue there was not whether AI creates value, but how much revenue has to arrive, and how quickly, to support the data centres, GPUs, debt and depreciation already being committed.
If model capability commoditises faster than revenue grows, the damage does not stop at the labs. It runs through hyperscalers, chipmakers, data-centre developers, utilities and the lenders financing the build-out. A cheaper Kimi is good for users and potentially bad for whoever underwrote scarcity.
Then there is the policy problem.
The US controls much of the closed-model frontier and much of the hardware underneath it. That is leverage, but only while the rest of the world remains dependent on the same stack. Restricting chips and APIs may preserve the lead but it can also make a parallel stack more valuable.
I wrote about the other side of this in Krugman, Fable 5, and Europe in Decline? . A closed model can be cut off by its provider or by the provider’s government. Once open weights have spread, they cannot be switched off in the same way.
There is another, more mundane risk: we may be overrating the model.
Every launch now comes with a dense page of benchmark wins. Some are real. Some use different harnesses, different budgets or model-specific tuning. Some benchmarks are simply saturated. A model that looks brilliant in a leaderboard can still be unpleasant over a long session.
At lower prices, more tasks become worth attempting. Agents can take more passes, read more context and handle work that is useful but not important enough to justify Fable pricing. Lower unit costs may not shrink the market at all.
The more interesting opportunity is separation. The interface, agent and model do not have to come from the same company.
A routine refactor can go to a cheaper model. A difficult architecture decision can go to Opus. Sensitive material can stay on a private deployment. Frontend work can go to whichever model happens to be best at it that month. The durable product may be the router, context layer and validation loop rather than any single model.
This is also where open weights matter beyond price. They let an organisation fine-tune, quantise and deploy a model under its own controls. That does not automatically make the result secure or compliant, but it changes who gets to answer the hard questions about data location, retention, access and continuity.
Another thing I have been thinking about this for a while: In How DORA Made Sovereignty a Bank Problem I argued that data sovereignty had made it’s way into operations. Banks now have to think about concentration, exit plans, audit rights and what happens when a critical foreign provider changes the rules.
AI will ( or does ) get the same treatment. Until recently that came with a large capability penalty. The newer releases make the calculation less silly. Kimi K3 is huge and its weights are due shortly. Alibaba is pushing the Qwen family in the same direction. Then there is Inkling which matters ( maybe even more ) because it adds a serious US-trained model with full weights available, one-million-token context and a clear focus on customisation.
So let’s do a back-of-the-napkin business case: Let’s assume a Swiss provider running one 64-GPU cluster in Zurich and selling dedicated or managed K3-class inference to banks, pharmaceutical companies, government bodies and other customers that care about Swiss data residency.
× The model uses roughly $7 million of initial capex, $370,000 of monthly operating cost, 70% utilisation and subscriptions between CHF 15,000 and CHF 50,000 per month. With those assumptions, one cluster reaches about CHF 7.4 million of annual revenue, CHF 3 million of EBITDA and a three-year payback.
The product would not really be tokens. It would be control: Swiss residency, a defined security perimeter, no training on customer data, audit access, contractual continuity and an exit route if the operator fails. Most companies will not literally put the model in the basement. They may pay somebody local to give them the functional equivalent.
There are reasons this may not work. A 2.8-trillion-parameter model is expensive to serve. One cluster is its own concentration risk. Hardware ages quickly. If utilisation drops, the economics get ugly. If a much smaller model performs nearly as well next year, the expensive rack becomes a monument to last year’s benchmark. Even so, I think this market will exist.
I do not think Kimi K3 kills Anthropic. I do not think one frontend reconstruction proves parity. And I would still choose Opus for the hardest coding work today.
But K3 was good enough that I forgot I was using it. It cost materially less. It ran inside the harness I already liked. Soon, its weights should be available to anybody willing to provide the hardware.
The question is no longer whether open models will catch the frontier in some abstract future. It is how much of a lead the closed labs need, how long they can hold it, and which customers will still pay for every extra point once switching becomes routine.
I write every few weeks about AI, sovereignty and the economics underneath both. Get the next one.
No tracking . Unsubscribe anytim

[truncated]
