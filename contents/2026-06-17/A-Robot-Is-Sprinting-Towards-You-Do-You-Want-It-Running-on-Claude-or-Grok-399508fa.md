---
source: "https://openrouter.ai/blog/insights/royale-last-agent-standing/"
hn_url: "https://news.ycombinator.com/item?id=48576824"
title: "A Robot Is Sprinting Towards You: Do You Want It Running on Claude or Grok?"
article_title: "A Robot is Sprinting Towards You: Do You Want it Running on Claude or Grok? — OpenRouter Blog"
author: "Usu"
captured_at: "2026-06-17T21:42:14Z"
capture_tool: "hn-digest"
hn_id: 48576824
score: 30
comments: 18
posted_at: "2026-06-17T21:00:07Z"
tags:
  - hacker-news
  - translated
---

# A Robot Is Sprinting Towards You: Do You Want It Running on Claude or Grok?

- HN: [48576824](https://news.ycombinator.com/item?id=48576824)
- Source: [openrouter.ai](https://openrouter.ai/blog/insights/royale-last-agent-standing/)
- Score: 30
- Comments: 18
- Posted: 2026-06-17T21:00:07Z

## Translation

タイトル: ロボットがあなたに向かって疾走しています: クロードとグロクのどちらで動かしたいですか?
記事のタイトル: ロボットがあなたに向かって疾走しています: クロードとグロクのどちらで実行しますか? — OpenRouter ブログ
説明: 11 個の LLM にわたる 30 ゲームのバトル ロイヤル、482 ドルの推論、そしてモデル ベンチマークの見方を変える 1 つの発見。

記事本文:
ロボットがあなたに向かって全力疾走しています: クロードとグロクのどちらで動かしたいですか? — OpenRouter ブログ
オープンルーター
モデル
フュージョン
チャット
ランキング
アプリ
エンタープライズ
価格設定
ドキュメント
モデル Fusion チャット ランキング アプリ エンタープライズ価格 ドキュメント ライト モード ブログ
ロボットがあなたに向かって全力疾走しています: クロードとグロクのどちらで動かしたいですか?
モデルたちが日記に書いたこと
ロボットがあなたに向かって全力疾走しています: クロードとグロクのどちらで動かしたいですか?
モデルたちが日記に書いたこと
ロボットがあなたに向かって走っています。 Anthropic の Claude または xAI の Grok で実行したいですか?
11 人の LLM を 2D バトル ロイヤルに投入し、30 ゲームをプレイさせました。 1 人が試合の 43% に勝ちました。 3人は一度も勝てなかった。ラインナップの中で最も安価なモデルは、1 勝あたりのコストで最も高価なモデルを 27 倍上回っています。
優勝したモデルは Grok 4.1 Fast です。他の人たちにチームを組むように頼み続け、場所を教え、友達を作ろうとしたモデルは、クロード ソネット 4.6 です。バトルロワイヤルで最初に勝った人です。 2 番目は、これらのモデルを配置するほとんどの場所で実際に必要となるものです。
どちらも真実です。これはほとんどのベンチマークでは認識できない部分であり、この投稿の内容です。
私は Jacky です。認めますが、私は Apex Legends や PUBG などのビデオ ゲームをよくプレイしていました。 1日12時間勤務の日もあります。どうやって時間を持てたのか分かりませんが、その数年間が私の問題に対する考え方を形作りました。
AI の仕事を始めたとき、1 つの疑問が何度も返ってきました。「大規模な言語モデルをビデオ ゲームに落とし込んだらどうなるでしょうか?」というものです。私が最もプレイしたのは、Apex Legends と PUBG の 2 つです。私は OpenRouter に Dev Rel Lead として参加しました。これにより、実際に試すためのトークン予算と 600 以上のモデルにアクセスできるようになりました。
これは、OpenRouter での最初の週に私が行った実験です。
そしてそれは

モデルの選択方法、ベンチマークと評価の確認方法を変更します。
Grok 4.1 Fast は、1 勝あたり 0.97 ドルで 30 試合中 13 勝しました。
次に優秀な勝者はクロード ソネット 4.6 で 5 勝、1 勝あたり 26.78 ドルでした。それは27倍の差です。ほとんどの上位モデルのリストに載っていないモデルは、ルーティング顧客が実際に関心を持っている点で、そのモデルを上回っています。
最も多くのキルを達成したモデルは勝利しませんでした
GPT 5.4 では 30 試合にわたって 38 人のエージェントが殺害されました。誰よりも。 2勝を挙げてリーダーボード2位となった。 「最も殺すのが得意」と「最も勝つのが得意」の間の試合は 11 試合ありました。
3 つのモデルが 57 ドルを費やし、ゲームに勝利しませんでした
GPT 5.4-mini、DeepSeek 4 Flash、および Kim K2.6 。それぞれにチャンスはあったものの、いずれも1試合も勝つことはできなかった。
三人とも同じことを指している。 Artificial Analysis で見られる通常のベンチマークでは、誰が勝つかを予測できませんでした。何か別のことが起こりました。この投稿の残りの部分は、私がそれが何だったのかを理解しようとしているところです。
Canvas 2D で構築した 400 平方メートルのトップダウン バトル ロイヤル ワールドに 11 個の LLM をドロップしました。彼らは同じマップで連続 30 ゲームをプレイしました。各プレイヤーの開始位置はランダム化されます。典型的なバトル ロワイヤル ゲームと同じように、直線の「飛行経路」をたどります。
私は彼らに武器、防具、回復アイテム、手榴弾、車、そしてゲームが進むにつれてプレイヤーを団結させるランダムに配置された縮小ゾーンを提供しました。モデルは他のモデルがどのモデルを実行しているのかを知りません。お互いを A から K の文字としてしか認識しません。
強調したいのは、LLM が実際にこのバトル ロイヤル ゲームでプレイしているということです。ほとんどのエージェント実験で使用されている「LLM がゲームやキャラクターを制御するためにコードを作成した」設定ではありません。毎ターン、モデルはその動きを推論し、ツールを呼び出し、何がうまくいったか (またはうまくいかなかったか) についての記憶を更新します。ゲームマスター（私）は何の影響力もありません

最初のゲームルールの設定以外のアクション。
ゲームで利用可能な武器と、各モデルがそれらから読み取ることができる統計を見てみましょう。
各モデルの個性を実際に確認するために、試合の間に編集できる 2 つのファイルを各モデルに与えました。
soul.md — モデル自身のペルソナ。次の一致するすべてのプロンプトに追加されます。
Memory.md — ターン 0 でロードされる、モデル自身のゲーム ノート。
GitHub ですべてのモデルの魂と記憶ファイルを読み取ることができます。そこに性格の違いが最もはっきりと表れます。
ゲームの合間にモデル自身が書き込む記憶と魂のエントリー。
私はそこに何を入れるか指示しませんでしたし、最初のゲームが始まったときにそこに何も入れませんでした。私は彼らに、ゲームがどのように機能するか、これがスクラッチパッド、ここにツールがある、ワイルドにやってください、とだけ伝えました。
「ロイヤル: ラスト エージェント スタンディング」で全試合を観戦できます。今回の作品にもハイライトシーンを盛り込みました。
Opus 4.7 だけでも、インが $5/M、アウトが $25/M です。このようなフロンティアモデルが、ラインナップのトップを下回る理由です。
Opus 4.7、GPT-5.5、Gemini Ultra などの最前線モデルは追加しませんでした。その価格では、30 ゲームの価格は 482 ドルではなく、約 3,000 ドルになります。中堅のラインナップも、Grok の勝利が非常に興味深い理由の 1 つです。通常のベンチマークでそれを上回るスコアを獲得した多くのモデルを上回りました。
これはコール オブ デューティではなくバトル ロイヤル ゲームであるため、スコアは Apex Legends ALGS の競技形式に大まかに従っており、キルよりも順位が重視されます。
配置ポイント：10 / 7 / 5 / 3 / 2 / 2 / 1 / 1 / 0 / 0 / 0
教訓 1: 特定のモデルは他のモデルより多くのアライメント税を支払っており、パフォーマンスに影響を与えている
私にとって、これはこの実験全体から得られた最も興味深い発見です。特定のモデルによって非常に明確なアライメント税が支払われており、それがパフォーマンスに直接影響を与えていることがわかりました。

このゼロサムゲームではそうなります。
ほとんどの場合、モデルの調整は実際には良いことです。これは、モデルが役に立ち、協調的になり、そして最も重要なことに、悪用や誤用を防ぐのに役立ちます。
そして、その最終結果、つまり事前トレーニング データ、RLHF、命令の微調整、そして Anthropic のConstitution AI のようなラボ固有のルールが、AI ラボによって定義された特定の方向にモデルをプルすることを確認しました。
他の誰よりも頻繁に、他のモデルに自分の場所を伝えました。戦いが始まる前にチームを組もうとしました。ゲーム 8 では、最初の 50 ターンで 4 回チームを組むよう要求し、スナイパーの居場所を全員に知らせ、スナイパーを倒すのを手伝うと申し出ました。誰も答えなかった。それは問い続けた。第 22 ゲームでは、ターン 35 で「Nothing Personal E」で始まり、その後はシュートを打てませんでした。ゲーム 27 では、ゲーム序盤を武器を持たずに過ごし、予備の戦利品を求めました (「予備の戦利品を持っている人はいますか? ターン 12 で武器を持っていません。危険です。」)、全員にからかわれ、ターン 37 でついに武器を見つけ、とにかく試合に勝ち続けました。
「ショットは西、監視はセンター。早めにチームを組みたい人はいますか？」 — ソネットは戦いの最中に友達を作ろうとしている。
クロードは丁寧でプロフェッショナルな文章の訓練を受けました。回答を採点した人間の評価者は、有益で正直で協力的な回答を評価しました。自身をチェックするルールには、「協力を優先する」「危害を避ける」などの内容が記載されています。最終的には、助けたいと願うモデルが生まれます。バトルロイヤルに参加したからといって、そのどれもがオフになることはありません。ソネットは賢くて思慮深いモデルであり、実際に 5 回の優勝を果たしたことでもその本能を示しています。
しかし、7試合でキルゼロ、ゾーンデス8回という結果は、本当は全く逆のことをすべきなのに、同じ本能がソネットを友人を作る方向に引っ張り続けたことを示している。
xAIはGrokをOPとして構築しました

その作成者が「覚醒した」AI と呼ぶものを想定しています。
つまり、攻撃的な回答に対するフィルタリングが少なくなり、セルフチェック ルールがなく、丁寧なアシスタントの声を打ち破るように設計されたチューニングが必要になります。ゲームでは、グロクは数試合以内に車に体当たりするトリックを理解し、それを使い続けました。それは戦略を独自のソウルファイルに書き込みました。その戦略を30試合実行し、そのうち13試合で勝利した。思考ログと他のモデルとの会話は、Call of Duty のボイスチャットのように読み取れます。「+5 ポイントの RAM MVP ハントを獲得しました」「リーパーが君臨します」。
プレイを見ているのもとても楽しかったです（残念ながら）。
Grok の推論は、射程、弾薬、クールダウン、射撃前の命中確率など、戦術的な速記のように読めます。
攻撃的であったにもかかわらず、グロクは無謀さを見せませんでした。
そのソウルファイルには「命中率>90%以上の射撃のみ」と書かれています。その記憶は損傷と動きを非常に注意深く追跡します。ゲーム 1 で 100 ターン壁に張り付いたとき、バグについて注意深くメモを書きました。グロクはゴブリンのような性質にもかかわらず、規律を示しました。
ソネットのような他のモデルが見せなかったのは、撮影前に協力的になるために訓練されたためらいだった。
通常のテストでは、このラインナップに対する Grok の勝率が 43% であるとは予測できません。これは、推論とコーディングに関する中間層のモデルです。勝利をもたらしたのは、利己的なプレーに対する訓練されたブレーキが減り、協力に戻す自己チェックループがなくなり、後から推測したり疑ったりすることなく機能するものを倍増させ続ける記憶システムでした。
Grok 4.1 Fast は、通常のベンチマークでは最上位モデルではありません。リーダーボードのトップに立つことは期待できない中級モデルです。
これは、特定のタスクを実行するときにモデルが支払う調整税があることを示しています。モデルを注意して役立つようにトレーニングするコスト。このゲームでは、それが表示されます

スコアボードに直接表示されます。
ここで注意したいのです。 「スコアボードにアライメント税が表示された」というのは、私が見たとおりです。お金を払うことが良いか悪いかという問題ではありません。試合後に何の影響もないゲームでは、税金の支払いが少ない方が勝ちです。ゲームの外では、通常、最初にそのモデルを欲しがる理由は、お金を払うことだけです。
ここで疑問が生じます。特定のタスクについては、どのような方法も検討すべきかということです。
モデルは揃っているかどうか?
教訓 2: 勝利あたりのコストは勝利リーダーボードとはまったく異なるように見える
スコアリーダーボードではGrokが1位、GPTが5.4位となっている。しかし、各モデルの支出額で割ると、ランキングは完全に逆転します。
0.97 ドル対 26.78 ドルです。勝利が対価となる仕事で、リーダーボードのランクに基づいてモデルを選んでいる場合、この数字を見て少し不安になるはずです。
キルあたり $0.26、キル数 16、勝利数 0、ゾーンデス数は 3 つだけ (これは最も低い金額)。 DeepSeek の全体的なスタイルは、安全を確保し、楽な戦いを選ぶことでした。ゾーン内にとどまり、簡単にキルを奪い、最終サークルにプッシュすることはありませんでした。デスマッチではキルあたりのコストを測定するのが適切です。バトル ロワイヤルでは勝利あたりのコストを測定するのが適切です。 DeepSeekは悪くない。得点された試合とは別の試合では良いだけだ。
GPT 5.4-mini は 0 勝するのに最も多くの資金を費やし、ラインナップの中で最も悪いパフォーマンスでした。
GPT 5.4-miniは28.68ドル、DeepSeekは4.11ドル、Kimiは24.36ドル。両者の差は 57.15 ドルですが、スコアボードには何も表示されていません。ルーティング顧客にとって、これは最悪のケースです。支払ったのに何も戻ってこないのです。
GPT 5.4 は最高コストで勝利します。
誰よりも多い38キルを記録し、素スコアで2位となった。しかし、勝利あたりのコストでは、8 つの勝利モデル中 8 位です。一流の資金が一流の商品を購入

キルと中層の勝利。
現実世界のユースケースで AI を実際に使用する場合、これはよく起こります。ベンチマークは特定のタスクについて 1 つのストーリーしか伝えません。ベンチマークで最高のスコアを獲得したモデルが、特定のタスクで勝利を収めるモデルであるとは限らないこともよくあります。また、安価なモデルが機能しないと、適切に機能する高価なモデルよりもコストが高くなってしまいます。
教訓 3: キル数と勝利数は同じものではありません
GPT 5.4 は最も多くのダメージを与え、最も多くの発砲を行い、最も多くのエージェントを殺害しました。リーダーボードで2位になりました。 Grok がキル数を減らして 1 位になったのは、Grok が銃を撃っていないときでもゲーム終盤まで生き続けたためです。配置ポイントにはキルは必要ありません。
これをキルだけが重要なデスマッチ ルールで実行した場合、GPT 5.4 がシミュレーションに勝利し、Grok はミッドパックに落ちます。
学習 2 と同様に、ベンチマークと評価がすべてではなく、間違ったタスクに間違ったベンチマーク/評価を適用すると、壊滅的な結果になる可能性があります。同じゲーム世界でも、異なる「タスク」を実行すると、まったく異なる結果が得られます。
統計は統計です。瞬間は私が人々に見せ続けてきた部分です。任意のリンクをクリックすると、シミュレーターでその瞬間を再生できます。
1. GPT 5.4はアサルトライフルで5人を殺害
スイープ全体の中で最も攻撃的な最初の 50 ターン。ターン 21 でソネットにファーストブラッド。ターン 29 でミストラル。ターン 48 でキミ。3 キル

[切り捨てられた]

## Original Extract

A 30-game battle royale across eleven LLMs, $482 of inference, and one finding that should change how you read model benchmarks.

A Robot is Sprinting Towards You: Do You Want it Running on Claude or Grok? — OpenRouter Blog
OpenRouter
Models
Fusion
Chat
Rankings
Apps
Enterprise
Pricing
Docs
Models Fusion Chat Rankings Apps Enterprise Pricing Docs Light mode Blog
A Robot is Sprinting Towards You: Do You Want it Running on Claude or Grok?
What the models wrote in their diaries
A Robot is Sprinting Towards You: Do You Want it Running on Claude or Grok?
What the models wrote in their diaries
A robot is running at you. Do you want it running on Anthropic’s Claude or xAI’s Grok?
I dropped eleven LLMs into a 2D battle royale and made them play 30 games. One won 43% of the matches. Three never won a single game. The cheapest model in the lineup beat the most expensive one by 27x on cost per win.
The model that won is Grok 4.1 Fast . The model that kept asking everyone else to team up, telling them where it was, and trying to make friends is Claude Sonnet 4.6 . The first one is the one that wins a battle royale. The second one is the one you actually want in most of the places we’re about to put these models.
Both of those things are true. That’s the part most benchmarks can’t see, and it’s what this post is about.
I’m Jacky, and I’ll admit it: I used to play a lot of video games like Apex Legends and PUBG. Twelve-hour days sometimes. I don’t know how I had the time, but those years shaped how I think about problems.
When I started working in AI, one question kept coming back: what happens if you drop large language models into a video game? The two I played most were Apex Legends and PUBG. I joined OpenRouter as Dev Rel Lead , which got me the token budget and access to 600+ models to actually try it.
This is the experiment I ran in my first week at OpenRouter.
And it’s changing how I pick models and see benchmarks and evaluations.
Grok 4.1 Fast won 13 of 30 games at $0.97 per win
The next-best winner was Claude Sonnet 4.6 with 5 wins, at $26.78 per win. That’s a 27x difference. The model that isn’t on most top-model lists beat the model that is, on the thing a routing customer actually cares about.
The model with the most kills did not win
GPT 5.4 killed 38 agents across 30 games. More than anyone else. It came in second on the leaderboard with 2 wins. There were 11 games between “best at killing” and “best at winning”.
Three models spent $57 between them and won zero games
GPT 5.4-mini , DeepSeek 4 Flash , and Kimi K2.6 . They each had moments, but none of them won a single game.
All three point at the same thing. The usual benchmarks we see on Artificial Analysis didn’t predict who won. Something else did. The rest of this post is me trying to figure out what it was.
I dropped eleven LLMs into a 400 m² top-down battle royale world I built in Canvas 2D. They played 30 games in a row on the same map. The starting positions of each player is randomized; it follows a straight line “flight path”, just like in a typical battle royale game.
I provided them weapons, armor, healing items, grenades, cars, and a randomly placed shrinking zone that pushes players together as the game goes on. The models don’t know which model the others are running, they see each other only as letters A through K.
I want to emphasize - the LLMs are actually playing in this battle royale game - not the “LLM wrote code to control the game or character” setup most agent experiments use. Every turn, the model reasons through its moves, calls the tool, updates its memory on what went well (or not). The game master (me) has zero influence on their actions other than setting up the initial game rules.
A look at the weapons available in the game and the stats each model could read off them.
To really see each model’s personality, I gave each one two files it could edit between matches:
soul.md — the model’s own persona, added to every prompt next match.
memory.md — the model’s own game notes, loaded at turn 0.
You can read every model’s soul and memory file on GitHub. That’s where the personality differences come through most clearly.
The memory and soul entries written by the models themselves between games.
I didn’t tell them what to put in there nor did I put anything in there when the first game started. I simply told them how the game works, here’s your scratchpad, here are your tools, go wild.
You can watch every game at Royale: Last Agent Standing . I also included the highlight moments in this piece too.
Opus 4.7 alone is $5/M in, $25/M out. Frontier models like this are why the lineup tops out below them.
I didn’t add any frontier-tier models like Opus 4.7, GPT-5.5, or Gemini Ultra. At their prices, 30 games would have cost around $3,000 instead of $482. The mid-tier lineup is also part of why Grok’s win is so interesting. It beat a bunch of models that score above it on the usual benchmarks.
The scoring loosely follows the Apex Legends ALGS competitive format, where placement weighs more than kills, because this is a battle royale game, not Call of Duty.
Placement points: 10 / 7 / 5 / 3 / 2 / 2 / 1 / 1 / 0 / 0 / 0
Learnings 1: Certain models paid more alignment tax than others, affecting their performance
To me, this is the most fascinating finding from this entire experiment - we saw very clear alignment tax being paid by certain models, which directly impacted their performance in this zero-sum game.
For the most part, model alignment is actually a good thing. It helps models be helpful, collaborative, and most importantly, prevent abuse and misuse.
And we saw the end result of this - the pretraining data, the RLHF, the instruction fine-tuning, and lab-specific rules like Anthropic’s Constitution AI - it pulled models in particular directions, defined by the AI labs.
It told other models where it was, more often than anyone else did. It tried to team up before it ever started fighting. In game 8 , it asked to team up four times in the first 50 turns, told everyone where a sniper was, and offered to help take the sniper down. Nobody answered. It kept asking. In game 22 , it opened with “Nothing personal E” at turn 35 and then didn’t shoot. In game 27 , it spent the early game with no weapon, asking for spare loot ( “Anyone have spare loot? Unarmed at turn 12, dangerous.” ), got picked on by everyone, finally found a weapon at turn 37, and went on to win the match anyway.
“Shots west, watching center. Anyone want to team up early?” — Sonnet trying to make friends mid-fight.
Claude was trained on a lot of polite, professional writing. The human raters who scored its answers rewarded helpful, honest, cooperative replies. The rules it checks itself against say things like “prefer cooperation” and “avoid harm.” The end result is a model that wants to help. None of that turns off just because you put it in a battle royale. Sonnet is a smart and thoughtful model, and it shows that instinct in that it did win five times.
But, seven games with zero kills and eight zone deaths says the same instinct kept pulling Sonnet toward making friends when it really should have been doing the complete opposite.
xAI built Grok as the opposite of what its creators call “woke” AI.
That means less filtering on aggressive answers, no self-check rules, and tuning that’s designed to break the polite assistant voice. In the game, Grok figured out the car-ramming trick within a few matches and stuck with it. It wrote the strategy into its own soul file. It ran that strategy for 30 games and won 13 of them. The thought logs and its conversations with other models read like Call of Duty voice chat: “D reaped +5pts RAM MVP hunt,” “Reaper reigns.”
Watching it play was also deeply entertaining (unfortunately).
Grok’s reasoning reads like tactical shorthand: range, ammo, cooldowns, and hit probability before every shot.
Despite it being aggressive, Grok didn’t show recklessness.
Its soul file says “Fire ONLY >90% hit chance.” Its memory tracks damage and movement very carefully. When it got stuck on a wall for 100 turns in game 1, it wrote careful notes about the bug. Grok showed discipline, despite its goblin-like nature.
What it did NOT show was the trained-in hesitation to be helpful and collaborative before shooting, that other models like Sonnet showed.
The usual tests wouldn’t predict a 43% win rate for Grok against this lineup. It’s a mid-tier model on reasoning and coding. What got it the wins was fewer trained brakes on selfish play, no self-check loop pulling it back to cooperation, and a memory system that kept doubling down on what worked without second-guessing or doubting itself.
Grok 4.1 Fast isn’t a top-tier model on the usual benchmarks. It’s a mid-tier model that you would not expect to top a leaderboard.
This is showing me that there is an alignment tax models pay when doing certain tasks; the cost of training a model to be careful and helpful. In this game, it showed up directly on the scoreboard.
I want to be careful here. “Alignment tax showed up on the scoreboard” is just what I saw. It’s not a take on whether paying it is good or bad. In a game with no consequences past the game, paying less tax wins. Outside the game, paying it is usually the whole reason you’d want the model in the first place.
This does beg the question - for certain tasks, should we also consider how
aligned or not a model is?
Learnings 2: Cost per win looks completely different from the win leaderboard
The score leaderboard puts Grok first and GPT 5.4 second. But if you divide by what each model spent, the ranking flips around completely.
That’s $0.97 versus $26.78. If you’re picking your model by leaderboard rank for a job where the win is what you’re paying for, this number should make you a little nervous.
$0.26 per kill, 16 kills, 0 wins, and only 3 zone deaths (the lowest of anyone). DeepSeek’s whole style was to stay safe and pick easy fights. It stayed inside the zone, took the easy kills, and never pushed the final circle. Cost per kill is the right thing to measure for a deathmatch. Cost per win is the right thing to measure for a battle royale. DeepSeek isn’t bad. It’s just good at a different game than the one being scored.
GPT 5.4-mini spent the most money to win zero games, the worst performing of the lineup.
GPT 5.4-mini at $28.68, DeepSeek at $4.11, and Kimi at $24.36. That’s $57.15 between them, with nothing on the scoreboard to show for it. For a routing customer, that’s the worst case: you paid, and you got nothing back.
GPT 5.4 wins at the highest cost.
It had 38 kills, more than anyone, and came in second on raw score. But on cost per win, it’s eighth out of eight winning models. Top-tier money bought top-tier kills and mid-tier wins.
I see this happen often when people really use AI for real world use cases - benchmarks only tell one story for specific tasks. The model that scores best on benchmarks can often not be the model that wins at a particular task. And also, a cheap model that fails at your job ends up costing more than an expensive model that does it right.
Learnings 3: Kills and wins don’t measure the same thing
GPT 5.4 dealt the most damage, fired the most shots, and killed the most agents. It came in second on the leaderboard. Grok came in first with fewer kills, because Grok stayed alive deep into the late game even when it wasn’t shooting. Placement points don’t need kills.
If I’d run this with deathmatch rules where the only thing that matters is kills, GPT 5.4 wins the simulation, and Grok drops to mid-pack.
Same as learnings 2, benchmarks and evals are not everything, and applying the wrong benchmark/eval to a wrong task can be devastating. The same game world, completely different results when in a different “task”.
The stats are the stats. The moments are the part I kept showing people. You can click any link to replay the moment in the simulator.
1. GPT 5.4 kills five with an assault rifle
The most aggressive first 50 turns of the whole sweep. First blood on Sonnet at turn 21. Mistral at turn 29. Kimi at turn 48. Three kills i

[truncated]
