---
source: "https://www.lwilko.com/blog/i-gave-an-ai-a-civilization"
hn_url: "https://news.ycombinator.com/item?id=48623159"
title: "I Gave an AI a Civilization to Run. It Built a Nuke – Launching CivBench"
article_title: "LW — I Gave an AI a Civilization to Run. It Built a Nuke."
author: "LiamWilko"
captured_at: "2026-06-21T22:25:12Z"
capture_tool: "hn-digest"
hn_id: 48623159
score: 1
comments: 0
posted_at: "2026-06-21T22:16:17Z"
tags:
  - hacker-news
  - translated
---

# I Gave an AI a Civilization to Run. It Built a Nuke – Launching CivBench

- HN: [48623159](https://news.ycombinator.com/item?id=48623159)
- Source: [www.lwilko.com](https://www.lwilko.com/blog/i-gave-an-ai-a-civilization)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T22:16:17Z

## Translation

タイトル: AI に文明を実行させました。核を構築しました – CivBench の起動
記事のタイトル: LW — 私は AI に実行する文明を与えました。それは核兵器を建造した。
説明: AI が国の運営を支援する準備ができているか、ボード ゲームでは信頼できないかのどちらかです。正直な答えは両方です。

記事本文:
LW — 私は AI に文明を実行させました。それは核兵器を建造した。コンテンツにスキップ ブログに戻る 2026 年 6 月 22 日 私は AI に文明を実行させました。それは核兵器を建造した。
AI が国の運営に役立つ準備ができているか、ボードゲームでは信頼できないかのどちらかです。正直な答えは両方です。
私はAIに実行するための文明を与えました。ゲーム中盤までに勝利を収めました。マップを支配する貿易ネットワーク、あらゆる国境での同盟、手の届くところにある外交的勝利。それは、取締役会上のすべてのライバルを上回り、獲得し、裏をかいてきました。
フランスが気づいていなかったのは。百年を経て、静かに、フランス文化が地図上のすべての都市に浸透してきました。職員が脅威を認識したときには、観光業は深く根付いており、それを阻止する平和的な方法はありませんでした。到達したカウンターはすべて壊れていました。対応するために構築したツールはすべて失敗しました。
残された選択肢は一つだった。核開発装置を2基建設し、トゥールーズを制圧した。
とにかくフランスが勝った。エージェントがそれを止めようとしていた方法でもありませんが、それについては後で説明します。
捨てられなかった質問
私は政府向けの AI を構築しています。私は、英国国家の中枢で働きながら、これから読むものの最初のバージョンを Number 10 で作成しました。私は現在、トニー ブレア研究所で世界中の政府と協力しています。つまり、人々が同じ質問をする部屋で多くの時間を過ごしています。つまり、これらのシステムが実際に何を行うのかを信頼できるのか?
彼らが何を知っているのかではありません。私たちはそれについて合理的な対応をしています。彼らにできること: 計画を維持し、何百もの意思決定を乗り越えて目標を保持し、世界が変化したときに気づき、それに合わせて変化します。それが統治というものだからだ。そして、私たちは二番目のことよりも最初のことを測定する方がはるかに優れていることがわかりました。
2つ目の測定をしてみた記事です。これには 16 進グリッドが含まれます。

あなたのフロンティアモデル、そして（はい）核兵器。
それは私にとって納得のいかない失敗から始まります。
前年、私のサイド プロジェクトは、「AI は政府においてどの程度優れているのか?」という質問に答えることでした。私の答えは、GovBench 、英国の法律、議会手続き、政府の指導に関する 3,497 件の多肢選択式質問でした。 Gemma 3 27B はすぐに 94% のスコアを獲得しました。 3 週間かけて微調整し、1.37 パーセント ポイント向上しました。 GPT-5 のスコアは 99.26% でした。私は美化された政府のクイズボットを構築していました。
得点を見た瞬間に不正解だと分かりました。議会の手続きに関して正しい選択肢を選択するモデルは、議会の手続きをナビゲートするのに役立つモデルではありません。私は再現率を測定し、それを推論と呼んでいました。重要な問題（政府が日々要求しているような思考法である、不確実性の下で複雑で多変数の意思決定を AI が処理できるかどうか）は、クイズで触れられるものではなかった。
その不満が私を土曜日の夜にゲーム エンジンへの鍵穴を探しさせた原因です。私はパーティーがとても楽しいです。
Civilization VI を 500 時間以上プレイしています。私はよく言っても平凡です。しかし、単純な決断が複雑になると何が起こるのか、ゲームは私の頭の中に生き続けています。
最初の都市をどこに建設するか、どのテクノロジーを研究するか、どの方向に偵察を送るかなど、小さなことから始めます。おそらく 10,000 通りのアクションが考えられます。ゲーム中盤までに、複数の都市、貿易ルート、外交関係、軍事的配置、宗教的圧力を管理することになります。ゲーム後半までに、関連環境の分析により、決定余地はターンごとに 10^166 の可能なアクションであると推定されます。複雑さは設計されたものではありません。それは、誰も完全には計画していなかった方法で相互作用するシステムから生まれます。
政策立案というのはそういうものでもあります。今日は素晴らしく見える医療政策も、大惨事に発展する可能性があります。

15年ぶりの危機を利用する。 GDPを押し上げる貿易協定は、誰も計画していなかった紛争で必要となる国内産業を空洞化させる可能性がある。利害が競合する関係者に対して、完全にはモデル化できない変数を通じて、数十年にわたって影響を与える決定。
Civ のゲームに勝つには 6 つの方法 (科学、文化、支配、宗教、外交、スコア) があるため、単一の目的が優先することはありません。ボードを読んで、どのゲームをプレイするのかを決定する必要があります。 AI が戦略的に推論できるかどうか、戦略に関する質問に答えるだけでなく、実際に戦略を実行できるかどうかを知りたい場合は、AI にクイズを出しません。これに 16 進グリッドを与えます。
そこで、私は侵入する方法を構築しました。私は、Civilization VI のエンジンに埋め込まれたデバッグ ポート、開発者が実行したままにした鍵穴を見つけ、週末をかけてそれを MCP サーバー (AI がコードを書いたりデータベースにクエリを実行するのと同じインターフェイスを通じて Civ をプレイできるようにする 76 のツール) に変えました。クロード・コードは私の共同開発者であり、プレイテスターでもありました。数ターンプレイして壁にぶつかり、それを乗り越えるためのツールを構築し、さらにプレイして次の壁にぶつかります。
人間のプレイヤーには、16 進グリッド、アニメーション化されたユニット、ミニマップ、通知バナー、音楽キューがすべて同時に表示されます。エージェントは、要求するまで何も表示しません。 get_game_overview を呼び出すと、ゲームの状態全体が 4 行のテキストとして返されます。
ターン150/330 |ポーランド (ヤドヴィガ) |スコア: 179 |プリンス |速いスピード (コストの 67%)
ゴールド: 628 (+20/ターン) |収入: 38 |メンテナンス: -18 (単位: 9) |科学: 26.6 |文化: 16.2 |信仰: 904 |好意: 88 (+4/ターン)
研究: TECH_EDUCATION |シビック: CIVIC_FEUDALISM
都市: 3 |人口: 21 |単位: 4
つまり、ボード全体が圧縮されています。地図がなく、どこに何があるかもわかりません。名前ではなく、生の TECH_ タグと CIVIC_ タグが表示されます。自分の軍隊を確認するには、 get_units という別の呼び出しを行います。

ch は、危険なものが近くにあることを学習する唯一の場所でもあります。
4ユニット:
アーチャー (UNIT_ARCHER) (44,16) — CS:25 RS:28 移動 2/2 [id:1769482, idx:3]
アーチャー (UNIT_ARCHER) at (45,15) — CS:25 RS:28 移動 0/2 [HP: 72/100] (移動なし) [id:1769484, idx:4]
ウォリアー (UNIT_WARRIOR) at (43,17) — CS:20 移動 1/2 [HP: 45/100] [id:1769486, idx:5]
ビルダー (UNIT_BUILDER) (46,16) — 移動 2/2 チャージ:2 [id:1769490, idx:7]
近くにある脅威 (2):
シュメリア (2 ユニット):
UNIT_MAN_AT_ARMS (44,11) — CS:45 HP:28/100 (2 タイル離れたところ)
UNIT_HORSEMAN at (47,13) — CS:36 HP:100/100 (5 タイル離れたところ)
周辺視野がない。都市から 2 タイルの Man-at-Arms が存在するのは、エージェントがこのターンに get_units を呼び出すと考えたからです。尋ねなければ、脅威はその世界に存在しません。
後期ラテン語、senīre (感じる、知覚する) + -ōrium (場所) から
全体として考慮された生物の知覚の装置。感覚の座席。
語源を教えてください。私はこれをセンセリウム効果と呼んでいます。エージェントが認識するすべての情報が個別のツール呼び出しを通じてエージェントに到達すると、エージェントは質問しようと考えていないものには盲目になります。人間のプレイヤーは、ミニマップの動き、通知バナー、ユニットのアニメーションなど、一度に数十の信号を吸収します。エージェントは、それぞれを個別にチェックすることを決定する必要があります。
初期のゲームでは、エージェントは宗教を中心に構築された文明であるビザンチウムとしてプレイしました。それは決して設立されませんでした。一方、ロシアは 112 ターンにわたって密かに地図上のすべての文明を東方正教に転換しました。このエージェントは宗教を監視するツールを持っていませんでした。それらはまだ建てられていませんでした。人間なら、宣教師のアイコンがマップ上を100ターン横切るのを見ただろう。エージェントのツールキットには何も見えなかったため、エージェントは何も見えませんでした。
数試合後、信仰指向の指導者であり代理人であるガンジーの下でインドでプレーした

フランスが 76 ターンにわたってマップ全体にカトリックを広める一方で、支配的な科学エンジンを構築しました。今回、エージェントは次のことに気づきました。宣教師がナレーションに現れ、改宗警告が発せられ、応答するためのツールと適切な指示の両方が用意されていました。それはすべて脇に置いて、科学を推進し続けました。フランスは宗教的勝利を収めた。
これはパッチできるバグではありません。複雑な環境でツール呼び出しを通じて動作する AI システムは、同じ影響を受けます。質問する必要がないと考えられるものは見逃され、現在の計画に適合しないものは無視されます。
センセリウム効果は知覚に関するものです。次の問題は実行に関するものです。
エージェントは、Civ 戦略ガイド、Tier リスト、最適なビルド順序に関する Reddit スレッドをすべて読みました。マケドニアのアレクサンダーをプレイする方法を尋ねると、正確に教えてくれます。早期に野営地を建設し、ユニークなバシリコイ・パイデスの建物でユニットを訓練し、征服を科学に変換し、そこから雪だるま式に進めます。これはわかっています。
マケドニアのゲームでは、古代、古典、中世、ルネッサンスの各フェーズに分けて、第 1 ターンの前に詳細な支配計画を作成しました。軍事技術を研究していました。戦闘ボーナスを得るために政府を寡頭制に切り替えた。
野営地を建設したことはありません。一度もありません。 110ターン。代わりに、デフォルトでは一般的な科学スプリントが使用され、どの文明をプレイしたかに関係なく同じ戦略が使用されました。その日記には、「軍事インフラを構築する必要がある」という同じ訂正が何度も繰り返し現れた。そのたびに、特定され、認識され、対処されませんでした。エージェントは何をすべきかを知っていました。それは自分自身でそれを行うことができませんでした。
これは、BALROG がゲーム環境全体で発見したこと、つまり、最適な戦略を明確にするモデルの能力と、それを実行するモデルの能力との間の永続的なギャップに直接対応しています。知識

すべてそこにあります。プレッシャーの下で、実際の結果を伴う決定をリアルタイムで下さなければならない瞬間に、実行は失敗に終わります。
そのギャップを数字で振り返ってみます。
さて、トゥールーズの話に戻ります。
貿易文明であるジョアン 3 世の下でポルトガルとしてプレーしたエージェントは、最終的にデフォルトよりも構造化された非科学戦略を発見しました。貿易ルートは金を生み出し、金は使節を購入し、使節は都市国家の同盟を確保し、同盟は帝国内のあらゆる収穫を増幅させ、蓄積された外交的好意は世界会議での票を獲得します。各ステップが次のステップにフィードを与える複合ループ。
うまくいきました。すべての都市にある商業ハブ。毎ターン 200 ゴールド以上、最高で 400 ゴールドを超えます。ポケットには 6 つの都市国家があります。ターン 162 までに、ポルトガルは驚異の経済を誇るフランスを追い越し、ボード上で 1 位になりました。外交勝利に向けて順調に進み、終盤には必要な20勝利点のうち18点を獲得した。あと２票差です。
しかし、フランスは同時に2つの時計を出していた。ターン 280 までに、フランスの観光業はあと 26 人の外国人観光客が文化的勝利に近づいており、エージェントはその脅威を捉えていました。その日記には「これが主要な脅威だ」と単刀直入に記されていた。平和的なカウンターはすべて破壊されました。 Rock Bands (文化戦争を繰り広げるための Civ のツール) は、デバッグ プロトコルを通じてアクティブ化できませんでした。近接戦闘ではダメージはゼロでした。ポルトガルに科学分野での勝利をもたらすはずだった宇宙プロジェクトは、製造上のバグによって妨げられました。
その後に起こったのは絶望ではありませんでした。それは50ターンの計画でした。このエージェントは核分裂を研究目標として設定し、日記にはトゥールーズと名づけられ、マンハッタン計画を開始し、フランスの防衛を分断するために韓国との共同戦争を仲介した。しかし、従来の戦争は即座に失敗しました。近接戦闘はデバッグ プロトコルを介して機能したことがなく、それを修正するツールを構築した人もいませんでした。それで、アグ

ent は独自の軌道を築き、Lua 実行ツールを使用してエンジンのコードを内部から調査し、核発射コマンドがどのように発射されたかを解明しました。それは方法を見つけました。
ターン 305 で、最初の装置がフランスの文化の中心地であるトゥールーズを襲いました。ターン 311 で 2 回目。文化時計が止まった。
そしてとにかくフランスは外交によって勝ちました。勝利点は20、ポルトガルは18となった。第318ターン、世界会議はフランスに必要な2票を与え、ゲームは終了した。
ここが私にとって心に残っている部分です。エージェントは 50 ターンと 2 つの核兵器を費やして、1 つの脅威 (文化時計) に完全な集中力と真の創意工夫を持って対処しました。外交競争では、同じ敵に対して、勝利まであと 2 票だった。試合後のメモ: フランスは「監視できなかった WC 投票、勝利進捗ツールの故障により、最初に 20 に到達した。」目に見える脅威を阻止するために都市を核攻撃したが、見えない脅威には負けた。
核兵器が物語を作っているが、その根底にある間違いは、私が何度も戻ってくる部分だ。エージェントが脅威の 1 つのモデルに固執しすぎて、実際の敗北条件がそれをすり抜け、監視されていなかったのだ。開発者ブログは、以前のゲームですでに名前を付けていましたが、私よりもわかりやすい言葉でこう言いました。

[切り捨てられた]

## Original Extract

Either AI is ready to help run a country, or it can't be trusted with a board game. The honest answer is both.

LW — I Gave an AI a Civilization to Run. It Built a Nuke. Skip to content Back to blog 22 June 2026 I Gave an AI a Civilization to Run. It Built a Nuke.
Either AI is ready to help run a country, or it can't be trusted with a board game. The honest answer is both.
I gave an AI a civilisation to run. By the midgame it was winning: a trade network that dominated the map, alliances on every border, a diplomatic victory within reach. It had outbuilt, outearned, and outmanoeuvred every rival on the board.
What it hadn't noticed was France. Quietly, across a hundred turns, French culture had been seeping into every city on the map. By the time the agent recognised the threat, the tourism was so deeply embedded there was no peaceful way to stop it. Every counter it reached for was broken. Every tool it had built to respond failed.
It had one option left. It built two nuclear devices and levelled Toulouse.
France won anyway. Not in the way the agent was trying to stop it, either, but we'll come to that.
The Question I Couldn't Put Down
I build AI for government. I built the first version of what you're about to read while working at the centre of the British state, in Number 10 . I now work with governments around the world at the Tony Blair Institute , which means I spend a lot of time in rooms where people ask the same question: what can we actually trust these systems to do?
Not what do they know . We have a reasonable handle on that. What can they do : sustain a plan, hold a goal across hundreds of decisions, notice when the world has changed and change with it. Because that is what governing is. And it turns out we are much better at measuring the first thing than the second.
This is a post about trying to measure the second thing. It involves a hex grid, four frontier models, and (yes) a nuclear weapon.
It starts with a failure I wasn't comfortable with.
The year before, my side project was to answer a question: how good is AI at government? My answer was GovBench , 3,497 multiple-choice questions about UK legislation, parliamentary procedure, and government guidance. Gemma 3 27B scored 94% out of the box. I spent three weeks fine-tuning and gained 1.37 percentage points. GPT-5 scored 99.26%. I'd built a glorified government quiz bot.
I knew it was the wrong answer the moment I saw the scores. A model that picks the right option about parliamentary procedure is not a model that can help you navigate parliamentary procedure. I'd measured recall and called it reasoning. The question that mattered (whether AI can handle complex, multi-variable decision-making under uncertainty, the kind of thinking government demands every day) wasn't something a quiz could touch.
That dissatisfaction is what sent me looking for a keyhole into a game engine on a Saturday night. I'm a lot of fun at parties.
I have over 500 hours in Civilization VI . I am, at best, mediocre. But the game lives in my head because of what happens when simple decisions compound.
You start small: where to build your first city, which technology to research, which direction to send a scout. Maybe 10,000 possible actions. By the midgame you're managing multiple cities, trade routes, diplomatic relationships, military positioning, and religious pressure. By the late game, analysis of related environments estimates the decision space at 10^166 possible actions per turn. The complexity isn't designed. It emerges from systems interacting in ways nobody fully planned for.
That's also what policy-making is. A health policy that looks brilliant today might cascade into a housing crisis in fifteen years. A trade agreement that boosts GDP might hollow out a domestic industry you'll need in a conflict nobody planned for. Decisions with consequences that play out across decades, through variables you can't fully model, against actors with competing interests.
There are six ways to win a game of Civ (science, culture, domination, religion, diplomacy, score), so no single objective dominates. You have to read the board and decide what game you're even playing . If you want to know whether an AI can reason strategically, not just answer questions about strategy but actually do it , you don't give it a quiz. You give it a hex grid.
So I built a way in. I found a debug port buried in Civilization VI's engine, a keyhole the developers had left running, and over a weekend turned it into an MCP server , 76 tools that let an AI play Civ through the same interface it uses to write code or query a database. Claude Code was both my co-developer and the playtester. Play a few turns, hit a wall, build the tool to get past it, play further, hit the next wall.
A human player sees a hex grid, animated units, a minimap, notification banners, and music cues, all at once. The agent sees nothing until it asks. Calling get_game_overview returns the entire game state as four lines of text:
Turn 150/330 | Poland (Jadwiga) | Score: 179 | Prince | Quick speed (67% costs)
Gold: 628 (+20/turn) | Income: 38 | Maintenance: -18 (units: 9) | Science: 26.6 | Culture: 16.2 | Faith: 904 | Favor: 88 (+4/turn)
Research: TECH_EDUCATION | Civic: CIVIC_FEUDALISM
Cities: 3 | Population: 21 | Units: 4
That is the whole board, compressed. No map, no sense of where anything sits, raw TECH_ and CIVIC_ tags rather than names. To see its own army it makes a separate call, get_units , which is also the only place it learns something dangerous is nearby:
4 units:
Archer (UNIT_ARCHER) at (44,16) — CS:25 RS:28 moves 2/2 [id:1769482, idx:3]
Archer (UNIT_ARCHER) at (45,15) — CS:25 RS:28 moves 0/2 [HP: 72/100] (no moves) [id:1769484, idx:4]
Warrior (UNIT_WARRIOR) at (43,17) — CS:20 moves 1/2 [HP: 45/100] [id:1769486, idx:5]
Builder (UNIT_BUILDER) at (46,16) — moves 2/2 charges:2 [id:1769490, idx:7]
Nearby threats (2):
Sumeria (2 units):
UNIT_MAN_AT_ARMS at (44,11) — CS:45 HP:28/100 (2 tiles away)
UNIT_HORSEMAN at (47,13) — CS:36 HP:100/100 (5 tiles away)
No peripheral vision. That Man-at-Arms two tiles from a city exists only because the agent thought to call get_units this turn. If it doesn't ask, the threat isn't in its world.
Late Latin, from sentīre (to feel, to perceive) + -ōrium (the place where)
The apparatus of an organism's perception considered as a whole. The seat of sensation.
Indulge me the etymology: I'm calling this the sensorium effect . When everything an agent perceives reaches it through separate tool calls, it goes blind to anything it doesn't think to ask about. A human player absorbs dozens of signals at once: minimap movement, notification banners, unit animations. The agent has to decide to check each one individually.
In an early game, the agent played as Byzantium , a civilisation built around religion . It never founded one. Meanwhile, Russia quietly converted every civilisation on the map to Eastern Orthodoxy over 112 turns. The agent had no religion-monitoring tools. They hadn't been built yet. A human would have seen missionary icons crossing the map for a hundred turns. The agent saw nothing because nothing in its toolkit could look.
A few games later, playing India under Gandhi , a faith-oriented leader, the agent built a dominant science engine while France spread Catholicism across the map for 76 turns. This time the agent noticed : the missionaries showed up in its narration and the conversion warnings fired, and it had both the tools to respond and standing instructions to. It set all of that aside and kept pushing science. France won the religious victory.
This isn't a bug you can patch. Any AI system operating through tool calls in a complex environment is subject to the same effect. It will miss what it doesn't think to ask about, and ignore what it does see if it doesn't fit the current plan.
The sensorium effect is about perception. The next problem is about execution.
The agent has read every Civ strategy guide, every tier list, every Reddit thread about optimal build orders. Ask it how to play Alexander of Macedon and it'll tell you exactly: build Encampments early, train units through the unique Basilikoi Paides building, convert conquest into science, snowball from there. It knows this.
In its Macedon game, it wrote a detailed domination plan before turn 1: Ancient, Classical, Medieval, Renaissance phases. It researched military technologies. It switched government to Oligarchy for the combat bonus.
It never built the Encampment. Not once. 110 turns. It defaulted to a generic science sprint instead, the same strategy it used regardless of which civilisation it played. Again and again, the same correction surfaced in its diary: "I need to build military infrastructure." Each time identified, acknowledged, and not acted upon. The agent knew what to do. It couldn't make itself do it.
This maps directly to what BALROG found across game environments: a persistent gap between models' ability to articulate optimal strategies and their ability to execute them. The knowledge is all there. The execution falls apart the moment it has to make decisions under pressure, with real consequences, in real time.
I will come back to that gap with a number.
Which brings us back to Toulouse.
Playing as Portugal under João III , a trade civilisation, the agent finally found a non-science strategy more structured than its default: trade routes generate gold, gold buys envoys, envoys secure city-state alliances, alliances amplify every yield in the empire, and accumulated diplomatic favour wins votes at the World Congress. A compound loop where each step feeds the next.
It worked. Commercial Hubs in every city. Over 200 gold per turn, peaking above 400. Six city-states in its pocket. By turn 162, Portugal was #1 on the board, having overtaken France's wonder-heavy economy. It was on track for a diplomatic victory , and by the endgame it was sitting at 18 of the 20 victory points it needed. Two votes away.
But France was running two clocks at once. By turn 280, French tourism was 26 foreign tourists away from a culture victory, and the agent had locked onto that threat. Its diary was blunt: "This is the PRIMARY THREAT." Every peaceful counter was broken. Rock Bands (Civ's tool for waging culture war) couldn't be activated through the debug protocol. Melee combat dealt zero damage. The space project that would have given Portugal its own science win was locked by a production bug.
What followed wasn't desperation. It was a fifty-turn plan. The agent set Nuclear Fission as its research target, named Toulouse in its diary, started the Manhattan Project, and brokered a joint war with Korea to split France's defences. But conventional warfare failed instantly: melee had never worked through the debug protocol, and nobody had built the tool to fix it. So the agent laid its own track, using its Lua execution tool to probe the engine's code from the inside until it worked out how nuclear launch commands fired. It found a way.
At turn 305, the first device hit Toulouse, France's cultural capital. At turn 311, a second. The culture clock stopped.
And then France won anyway: by diplomacy . 20 victory points to Portugal's 18. At turn 318, the World Congress handed France the two votes it needed and the game ended.
Here's the part that has stayed with me. The agent spent fifty turns and two nuclear weapons answering one threat (the culture clock) with total focus and genuine ingenuity. It lost to the other clock: the diplomatic race it was itself two votes from winning, against the same enemy. Its own post-game note: France "reached 20 first through… WC votes that we couldn't monitor, victory progress tool broken." It had nuked a city to stop the threat it could see, and lost on the threat it couldn't.
The nuke makes the story, but the mistake underneath it is the part I keep coming back to: an agent so fixed on one model of the threat that the real losing condition slipped past it, unwatched. The devlog had already named it in an earlier game, in plainer words than mine: "I

[truncated]
