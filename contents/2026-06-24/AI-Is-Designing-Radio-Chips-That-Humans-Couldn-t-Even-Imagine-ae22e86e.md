---
source: "https://spectrum.ieee.org/ai-radio-chip-design"
hn_url: "https://news.ycombinator.com/item?id=48660021"
title: "AI Is Designing Radio Chips That Humans Couldn't Even Imagine"
article_title: "AI Learns the “Dark Art” of RFIC Design - IEEE Spectrum"
author: "Brajeshwar"
captured_at: "2026-06-24T14:57:10Z"
capture_tool: "hn-digest"
hn_id: 48660021
score: 1
comments: 0
posted_at: "2026-06-24T14:02:16Z"
tags:
  - hacker-news
  - translated
---

# AI Is Designing Radio Chips That Humans Couldn't Even Imagine

- HN: [48660021](https://news.ycombinator.com/item?id=48660021)
- Source: [spectrum.ieee.org](https://spectrum.ieee.org/ai-radio-chip-design)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T14:02:16Z

## Translation

タイトル: AI は人間が想像すらできなかった無線チップを設計している
記事タイトル: AI が RFIC 設計の「ダークアート」を学ぶ - IEEE Spectrum
説明: RF チップ設計用 AI は、RL および拡散モデルを使用して、5G、レーダーなどを強化するワイルドかつ効率的な回路を作成し、RFIC のタイムラインを大幅に短縮します。

記事本文:
-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
AI が学ぶ RFIC 設計の「闇の芸術」 - IEEE スペクトラム IEEE.org IEEE Xplore IEEE 標準 IEEE 求人サイト その他のサイト サインイン 参加 IEEE AI は人間が想像すらできない無線チップを設計している テクノロジー インサイダー検索: トピック別に探索 航空宇宙 AI 生物医学 気候 テクノロジー コンピューティング 家庭用電化製品 エネルギー 技術の歴史 ロボット工学 半導体 電気通信 輸送
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
アクセスサウザンド

の記事 — 完全に無料
アカウントを作成して限定コンテンツと機能を入手してください:
記事の保存、コレクションのダウンロード、
そして
コメントを投稿する
— すべて無料です！フルアクセスと特典については、
購読する
スペクトルに。
AIは人間が想像すらできない無線チップを設計している
わかりやすさや美しさから解放され、AI によるデザインの高速化
垂直
チャド・ハーゲン
緑の概要
RFIC 設計は、5G、自動運転車、衛星通信などの無線技術の進歩を制限する複雑な「闇の芸術」です。
プリンストンの研究者は、強化学習と逆設計を使用して、RFIC をゼロから迅速に作成します。
拡散モデルは、新規または人間が解釈可能な RF レイアウトを迅速に生成し、記録的なパフォーマンスを達成し、設計時間を大幅に短縮します。
将来の進歩には、AI が普遍的な電磁気と回路の動作を学習できるように、大規模な共有チップ設計データセットとオープン エコシステムが必要です。
少し時間を取って、過去 30 年間のワイヤレスの進歩がなかったあなたの生活を想像してみてください。
荷物を紛失しましたか？ AirTags が発明されていないのは何と残念なことでしょう。航空会社の担当者が最新情報を電話することを約束したので、手頃な価格の携帯電話がないので、台所の電話のそばで長い間待つことにします。ストリーミング サービスがないため、待っている間はラジオを聞くことになります。台無しになっていたであろう映画のプロットのすべてについて語るまでもない。
これは、ワイヤレス テクノロジーが日常生活の中でどのように感じられるかを示すほんの一部にすぎません。サプライチェーン、インフラ、経済運営に与えた影響は世界を変えてきました。
すべてのデバイスが目立たずに情報を送受信できるようにする高周波集積回路がなければ、どれも不可能です。
N

このテクノロジーのさらなる進化が何をもたらすか想像してみてください。自動運転車の普及、量子通信、6G モバイル サービス、衛星通信などです。勢いが続くかどうかは、今日の RF チップのより新しく、より高度なバージョンにかかっています。
しかし、問題もあります。世界のほとんどのコンピューティング チップの設計は独自の科学として標準化されていますが、RF 設計は依然として芸術の領域に留まっています。闇の芸術でさえ、長年の経験によってのみ習得されます。魔術師なら誰でも言うように、闇の魔術は独自のスケジュールを守っています。そしてそのスケジュールは、RF チップの設計だけでなく、RF チップに依存する他のすべてのテクノロジーの進歩を妨げています。
約 7 年前、AlphaGo が囲碁の世界チャンピオン、イ・セドルに勝利した後、プリンストン大学の学生たちと私は、「AI にもこの技術を教えられるだろうか？」と考え始めました。最近の成功例は、大部分はそれが可能であることを示唆しています。ここ数年にわたり、私たちのグループとこの分野の他のリーダーは、RFIC を設計するための機械学習主導のアルゴリズム手法の開発を開始しました。結果として得られるチップの中には、回路レイアウトというよりも現代アートのように見えるものもあります。しかし多くの場合、物理プロトタイプは性能の点で最先端の回路を上回っていました。しかし、本当の成果は、AI が実際に機能するデザインを思いつくのに、人間のデザイナーよりもはるかに短い時間がかかったということです。
これは 1 つまたは 2 つの RF チップに関するものではありません。 AI を活用した設計は、すべての RF 設計の未来となる可能性があり、さらにそれ以上の未来となる可能性があります。
では、なぜこれらのチップはすべて手作業で作られなければならないのでしょうか? CPU や GPU と同様に、RFIC はアルゴリズム合成プロセスを使用して設計されていないのはなぜですか?
RFIC の設計は、複数の物理ドメインにわたるエンジニアリングの演習です。マクスウェル方程式、

さまざまな空間的および時間的スケールを横断し、電磁場が能動デバイスおよび受動デバイスとどのように相互作用するかを決定します。これらのデバイスは、チップが機能するために慎重に共同設計する必要があります。これらに加えて、動作中に熱がどのように生成および除去されるかを決定する熱力学の法則や、チップとそのパッケージングが温度変化にどの程度確実に耐えられるかを決定する熱膨張と熱収縮の仕組みも決まります。
AI が RFIC 設計をショートさせる可能性
高周波集積回路の設計には、人間の直観と、頻繁に繰り返される複数の最適化ステップが必要です。マクスウェル方程式を理解することで、このプロセスを短縮して設計を迅速に作成するように AI を学習できることが期待されています。
これらが課すすべての物理的制約を同時に考慮すると、設計スペースがほとんど不可能になるほど大きくなります。すべての決定には複雑な優先順位が含まれており、それらが互いに競合することが多く、いずれの最適化も妨げられます。
この問題をよりよく理解するために、関連する手順を見てみましょう。その後、1 つの新しいチップ設計に何年もかかり、数千万から数億ドルかかる理由がよく理解できるでしょう。
高周波集積回路の領域の大部分は、複雑な電磁構造によって占められています。このブロードバンド パワー アンプ [1] のような人間が設計した RFIC は、テンプレートから始まり、対称的でわかりやすいパターンに従います。しかし、人間が設計したテンプレートの制約や、人間が電磁構造の理論的根拠を理解する必要さえないため、パワーアンプ IC [2-5] と低ノイズアンプ [6] は、真にワイルドな外観でありながら効率的な設計を実現できます。セングプタラボ
あなたが 5G ミル用の新しい 28 ギガヘルツ パワー アンプの設計を任されたエンジニアだとしましょう

ライメーター波ハンドセット。 (これは、携帯電話の 5G 信号を増幅してアンテナに送信し、離れた基地局で受信できるようにするタイプの RFIC です)。どこから始めますか?
RFIC設計には住宅建築と共通する特徴がいくつかあります。家の設計図が、建設する寝室とバスルームの数、およびそれらを接続する廊下を決定するのと同じように、アーキテクチャと呼ばれる RFIC の設計図は、RFIC が意図された機能を果たすために必要な要素の種類を確立します。部屋の代わりに、アーキテクチャには、たとえば、パワーアンプに必要な増幅段数が含まれます。廊下の代わりに、信号がこれらの段階を通過するために通らなければならない経路を示します。
RFIC の設計図は、実際にはほとんどが廊下にあります。インダクタや伝送線路などの受動素子は、トランジスタなどの能動素子よりもはるかに多くの面積を占めます。
その理由は次のとおりです。皆さんも経験したことがあるでしょうが、一般的な CPU のトランジスタは、わずか数ギガヘルツの動作周波数にさらされると過熱します。 RFIC が動作できる周波数は一桁高く、5G 信号の場合は 28 GHz と 39 GHz、衛星通信の場合は 26.5 ～ 40 GHz およびさらに高く、自動車レーダーの場合は 77 GHz です。この猛攻撃の下では、CPU のトランジスタが故障してしまいます。
RFIC トランジスタは、これらのチップが慎重な電磁設計により信号エネルギーを巧みに管理するため、この運命を回避します。これは、チップの領域を支配する金属元素のビザンチン ネットワークの形をとります。これらの構造は幾何学的に規則的で、多くの場合対称的で、非常に複雑に構築されているため、レースのようなフィリグリーに似ている場合もあります。しかし、それらは装飾的に見えるかもしれませんが、チップの機能には不可欠です。
電気的に言えば

いいえ、これらの「廊下」はチップの配管のように機能します。配管と同様に、この広範な受動部品の迷路は、電磁エネルギーをチップの周りで伝わるべき場所にのみ閉じ込めます。
RFIC 設計における主な課題は、設計図から家を建設するときに耐力梁、パイプ、外壁の正確な仕様が必要になるのと同じように、これらすべての要素を組み合わせて確実に機能するようにすることです。 RFIC では、信号がチップ内を通過して処理できるように、物理的に製造可能なトランジスタと受動部品を接続してアーキテクチャを実現する必要があります。これらのデバイスがローカルに接続される方法を、回路のトポロジーと呼びます。
そのパワーアンプを作るための最初のステップは、回路テンプレートの候補を特定することです。これは、特定の回路トポロジを備えた特定のアーキテクチャの目標を満たす構造の組み合わせです。研究者たちは長年にわたり、特定の機能向けに再利用可能なデザイン テンプレートを開発することで、ユーザーの負担を軽減してきました。たとえば、テンプレートは、回路に必要な増幅段の数を提案します (場合によっては、2 つの小型アンプの出力を組み合わせた方が、単一の大型アンプから得られるよりも優れた帯域幅と効率が得られるためです)。そして、受動構造の一般的な構成がどうあるべきかを示唆しています。現在、そのようなテンプレートの広範なライブラリが存在します。
ただし、それぞれにトレードオフがあるため、既製のものをそのまま使用することはできません。安定性を犠牲にして利得が向上するものもあります。効率を犠牲にして帯域幅を改善する。さらに、出力電力を犠牲にしてエネルギー効率が高いものもあります。明確な最善の選択が存在することはほとんどありません。
これらのさまざまなパラメータがすべて揃う「スイートスポット」に到達するには

設計者は通常、長年のトレーニングで習得した直観と方法を使用して、回路のいくつかの異なるバージョンをレイアウトします。
課題は、アーキテクチャ、回路トポロジ、または電磁パッシブに関する決定を個別に行うことができないことです。 1 つの決定が他の決定に影響を与えます。そのため、RF 回路の設計は、狭い部屋に特大のカーペットを収めようとするような感じになることがよくあります。隅の 1 つを押すと、もう 1 つが飛び出すのです。
マイクロ波やミリ波の周波数では、ほんのわずかなミスがチップが動作するか動作しないかの違いとなり、さまざまな問題が発生する可能性があります。たとえば、電磁波がトランジスタやその他のコンポーネントに遭遇した場合、電磁波が伝わる経路は次に来るものと適切に「一致」する必要があります。そうでない場合、エネルギーの一部は前方に流れるのではなく後方に反射します。高圧消火ホースを庭の細いホースに直接接続しようとしているところを想像してみてください。適切なアダプターを使用しないと、接続部で水が後方に飛び散ります。通過できるものはほとんどありません。エレクトロニクスでは、これはインピーダンス整合問題と呼ばれます。
これらの反射を防ぐために、エンジニアはハンドオフをスムーズにする特別なトランジション、つまり本質的に微細なアダプターを設計します。

[切り捨てられた]

## Original Extract

AI for RF chip design slashes RFIC timelines, using RL and diffusion models to craft wild yet efficient circuits that supercharge 5G, radar, and more.

-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
AI Learns the “Dark Art” of RFIC Design - IEEE Spectrum IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE AI Is Designing Radio Chips That Humans Couldn’t Even Imagine Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
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
AI Is Designing Radio Chips That Humans Couldn’t Even Imagine
Freed from intelligibility and aesthetics, AI designs faster
Vertical
Chad Hagen
Green Summary
RFIC design is a complex “ dark art ” that limits progress in wireless technologies like 5G, autonomous vehicles, and satellite communications.
Princeton researchers use reinforcement learning and inverse design to rapidly create RFICs from scratch.
Diffusion models rapidly generate novel or human-interpretable RF layouts, achieving record performance and drastically reducing design time.
Future progress needs large, shared chip design datasets and open ecosystems so AI can learn universal electromagnetic and circuit behaviors.
Take a moment and try to imagine your life without the wireless advances of the past three decades.
Have you lost your luggage? What a shame AirTags have not been invented. The airline representative has promised to call with updates, so settle in for a long wait by the kitchen telephone, because there are no affordable cellphones. You’ll be stuck listening to whatever is on the radio while you wait, because there are no streaming services. That’s not even to speak of all the movie plots that would have been ruined.
This is just a tiny sliver of how wireless technology makes itself felt in your day-to-day existence. The effects it has had on supply chains, infrastructure, and how the economy runs have been world-altering.
None of it would be possible without the radio-frequency integrated circuits that allow all our devices to unobtrusively send and receive information.
Now imagine what the further evolution of this technology will bring: Wide-spread autonomous vehicles , quantum communications , 6G mobile service and satellite communications. Continued momentum will depend on newer and more advanced versions of today’s RF chips.
But there’s the rub. Whereas the design of most of the world’s computing chips has been standardized into its own science, RF design has remained stubbornly in the realm of art. A dark art, even, that is mastered only through years of experience. As any sorcerer will tell you, the dark arts keep their own schedule. And that schedule is impeding progress not just in RF chip design but in every other technology that depends on it.
About seven years ago, in the wake of AlphaGo’s victory over world Go champion Lee Sedol , my students at Princeton and I began to wonder: Could AI be taught this art as well? Recent successes suggest that, to a large extent, it can. Over the last few years, our group and other leaders in the field have started to develop machine-learning-driven algorithmic methods for designing RFICs . Some of the resulting chips look more like modern art than circuit layouts. Yet in many cases, the physical prototypes bested state-of-the art circuits in terms of performance. The real achievement, however, is that it took the AI orders of magnitude less time to conceive a working design than it would a human designer.
This is not about one or two RF chips. AI-enabled design could be the future of all RF design, and maybe much more.
So why do these chips all have to be crafted by hand? Why aren’t RFICs designed with an algorithmic synthesis process, much as CPUs and GPUs are?
The design of RFICs is an exercise in engineering across multiple physical domains. Maxwell’s equations , operating across different spatial and temporal scales, govern how electromagnetic fields interact with active and passive devices that must be carefully codesigned for the chip to function. Alongside these are the laws of thermodynamics, which determine how heat is generated and removed during operation, as well as the mechanics of thermal expansion and contraction that dictate how reliably the chip and its packaging survive temperature changes.
AI Could Short-Circuit RFIC Design
The design of a radio-frequency integrated circuit requires human intuition and multiple, often-repeated optimization steps. The hope is that through an understanding of Maxwell’s Equations, an AI can be taught to short-circuit this process and quickly produce a design.
Simultaneously accounting for all the physical constraints these impose makes the design space almost impossibly large. Every decision involves complex priorities that often compete with one another, preventing the optimization of any of them.
To better understand the issue, let’s walk through the steps involved, after which you’ll better understand why a single new chip design takes years and tens to hundreds of millions of dollars.
Most of the area of radio-frequency integrated circuits is dominated by complex electromagnetic structures. Human-designed RFICs, like this broadband power amplifier [1], start with templates and follow a symmetric, understandable pattern. But freed from the constraints of human-designed templates and the need for humans to even understand the rationale of electromagnetic structures, power amplifier ICs [2–5] and low-noise amplifiers [6] can take on truly wild-looking yet efficient designs. SENGUPTA LAB
Let’s say you’re an engineer assigned to design a new 28-gigahertz power amplifier for a 5G-millimeter-wave handset. (This is the type of RFIC that boosts the 5G signals on your phone and transmits them to the antenna where they can be picked up by a distant base station). Where do you start?
RFIC design has some features in common with house building. Just as the blueprint for a house dictates the number of bedrooms and bathrooms to be built and the hallways connecting them, the blueprint for an RFIC—called the architecture—establishes the kinds of elements the RFIC needs to fulfill its intended function. Instead of rooms, the architecture includes, for example, the number of stages of amplification your power amplifier needs. Instead of hallways, it shows the paths that signals must take to get through those stages.
The blueprint for RFICs is actually mostly hallway ; passive elements, like inductors and transmission lines, take up far more real estate than active elements like transistors.
Here’s why. As you have probably experienced yourself, a typical CPU’s transistors overheat when faced with operating frequencies of just a few gigahertz. The frequencies RFICs can operate at are higher by an order of magnitude—28 and 39 GHz for 5G signals, 26.5 to 40 GHz and even higher for satellite communications, and 77 GHz for automotive radar . Under this onslaught, a CPU’s transistors would fail.
RFIC transistors avoid this fate because these chips cleverly manage the signal’s energy with careful electromagnetic design. This takes the form of byzantine networks of metal elements that dominate the chip’s real estate. These structures are geometrically regular, often symmetrical, and so intricately constructed they sometimes resemble lacelike filigree. But while they may look decorative, they are essential to the chip’s functioning.
Electrically speaking, these “hallways” work more like the chip’s plumbing . Like plumbing, this extensive labyrinth of passives confines electromagnetic energy only to the places it should be traveling around the chip.
The major challenge in RFIC design is putting all these elements together to ensure they work, just as constructing a house from its blueprints demands exact specs for load-bearing beams, pipes, and external walls. On an RFIC, the architecture needs to be realized with physically fabricable transistors and passive components that are connected just so, to permit the signal to travel through the chip and be processed. The way these devices are connected locally is what we call the circuit’s topology.
To make that power amplifier, then, your first step is to identify a candidate circuit template: The combination of structures that will meet the goals of a particular architecture with a specific circuit topology. Over the years, researchers have eased your burden by developing reusable design templates for specific functions. For example, templates suggest how many amplification stages a circuit needs (because sometimes, combining the output of two smaller amplifiers will result in better bandwidth and efficiency than you would get from a single larger one). And they suggest what the general configuration of the passive structures should be. Today there is an extensive library of such templates.
However, these can’t simply be used off-the-shelf, because each comes with trade-offs. Some have better gain at the expense of stability; some better bandwidth at the expense of efficiency; still others are more energy efficient at the expense of output power, and so on. There is rarely a clear best choice.
To arrive at the “sweet spot” where all these different parameters are balanced into optimal harmony, designers will typically lay out several different versions of the circuit, using intuitions and methods they have picked up in their years of training.
The challenge is that the decision around the architecture, circuit topology, or the electromagnetic passives cannot be done separately. One decision influences the others. So, designing an RF circuit can often feel like trying to fit an oversized carpet into too small a room—press down one corner, and another pops up.
At microwave and millimeter-wave frequencies, even the smallest misstep is the difference between a chip that works and one that doesn’t, and any number of things can go wrong. For example, when an electromagnetic wave encounters a transistor—or any other component —the path it travels must be properly “matched” to what comes next. If it isn’t, some of the energy reflects backward instead of flowing forward. Imagine trying to connect a high-pressure fire hose directly to a narrow garden hose. Without the right adapter, water will splash backward at the junction. Very little will make it through. In electronics, this is called the impedance-matching problem.
To prevent those reflections, engineers design special transitions, essentially microscopic adapters, that smooth the handof

[truncated]
