---
source: "https://shrikar.com/writing/ml-engineer-to-ai-native-reskill"
hn_url: "https://news.ycombinator.com/item?id=48505186"
title: "From ML engineer to AI-native: reskilling toward an edge"
article_title: "From ML engineer to AI-native: reskilling toward an edge | Shrikar Archak"
author: "shrikar"
captured_at: "2026-06-12T16:07:09Z"
capture_tool: "hn-digest"
hn_id: 48505186
score: 1
comments: 1
posted_at: "2026-06-12T15:11:20Z"
tags:
  - hacker-news
  - translated
---

# From ML engineer to AI-native: reskilling toward an edge

- HN: [48505186](https://news.ycombinator.com/item?id=48505186)
- Source: [shrikar.com](https://shrikar.com/writing/ml-engineer-to-ai-native-reskill)
- Score: 1
- Comments: 1
- Posted: 2026-06-12T15:11:20Z

## Translation

タイトル: ML エンジニアから AI ネイティブへ: エッジに向けた再スキルアップ
記事のタイトル: ML エンジニアから AI ネイティブへ: エッジに向けた再スキルアップ |シュリカー・アーチャク
説明: ML エンジニアリングの未分化な中間部分が圧迫されてきています。データの厳密性は依然として優れており、実験ループの力はさらに高まり、具体的な 90 日間の計画は次のとおりです。

記事本文:
コンテンツにスキップcom AI ネイティブ開発の微調整エージェントの評価 クロード コード AI ネイティブ開発の RSS X について 微調整エージェントの評価 クロード コード レポートの作成について スニペット ケース スタディ GPU 電卓 X / Twitter LinkedIn エッセイ 2026 年 6 月 11 日 12 分で読む
ML エンジニアから AI ネイティブへ: エッジに向けた再スキルアップ
ML エンジニアリングの未分化な中間部分が圧迫されてきています。データの厳密性は依然として優れており、実験ループの力はさらに高まり、具体的な 90 日間の計画は次のとおりです。
月曜日にリポジトリを開くと、火曜日に費やす予定だったデータ パイプラインの足場がエージェントによってすでに構築されています。機能ストアの配線、train/val の分割、標準モデルのボイラープレート、eval スクリプトは、あなたが眠っている間に、テストとともに完了しました。良い仕事だ。そして、安堵感のどこかに、あなたが大声では言わなかった静かな考えがあります。エージェントが私が報酬をもらっていた役割を果たせるとしたら、1年後の私の仕事は一体何になるでしょうか？
私もその気持ちを経験したので、その気持ちを大切にしたいと思っています。私は何年も微調整と ML 担当者として過ごしてきましたが、正直な答えは「リラックスしてください、何も変わりません」でも、「パニックになっています、ML は死んだ」でもありません。どちらよりも具体的で、一度形を見れば動きは明らかです。
ML エンジニアリングの未分化な中間部分、つまり標準パイプライン、モデルのトレーニング、ピクルスのハンドオフが圧迫されてきています。 2つのエッジは耐久性を維持します。データと目的 (ランキング、pCTR、入札、最適化、調整) に結び付けられたディープ ML はどこにも行きません。 AI ネイティブのエンジニアリングが成長しています。立っていると危険な場所は真ん中です。終わりを選んでください。
どこにも行かないことから始めてください。破滅がかかるので、それは常にスキップされます。
Deep ML は目的に結び付けられ、データはこれまでと同様に防御可能になります。ランキングモデルth

at は、2 億人のユーザーに何を表示するかを決定し、広告オークションに組み込まれたクリックスルー率予測器、入札戦略、需要予測、支払いの異常検出器など、これらは独自のデータとビジネス上の制約を計算したものです。オークションのダイナミクスやラベルの分布を知るために呼び出すことができるフロンティア API はありません。あなたがこれらのいずれかに深く陥っている場合、そこから離れて再スキルを習得する必要はありません。それはより深く追求することです：より良いキャリブレーション、より良いオンライン/オフラインギャップ分析、より明確な目標。このエッジは本物であり、深くなるからといって反復が遅くなる必要はありません。この投稿のさらに下にあるエージェントによる探索ループは、反対側に渡る人だけでなく、あなたにも役立ちます。
しかし、ここが正直なところ複雑です。なぜなら、そのエッジの内側の線も動いているからです。これらのシステムの検索レイヤー (コンテンツベースの推奨、セマンティック検索、候補生成、古典的に適用された NLP の大部分) は、埋め込みモデル、LLM、およびそれらの微調整といった汎用レイヤーに着実に崩壊しています。以前は、ドメインごとの機能と特注の取得スタックを手動で構築していました。一般的な埋め込みモデルを呼び出したり、ドメインに合わせて微調整したりするモデルが増えており、職人によるパイプラインを打ち負かしています。そのサブレイヤーは一般化されており、高速です。
自分のものとして残るのは、ビジネス目標に基づくランキング、調整、オークションと入札のロジック、オンラインとオフラインのギャップ、スコアを意思決定に変える最適化など、目的と制約に縛られた部分です。一般的なモデルで候補を取得できます。マーケットプレイスが 1 人の販売者に過剰にサービスを提供していることや、特定のオークションで精度と収益をどのように交換するかを知りません。それが耐久性のあるコアであり、「I do resys」よりも細くて鋭いです。
したがって、左の重量を正確に読み取ります。耐久性のあるものは、

右エッジを強化する同じ汎用レイヤーに移行している検索およびコンテンツ理解サブレイヤーではなく、目的に限定されたモデリングです。それがこの投稿全体のパターンです。一般層は標準化できるものは何でも吸収し続けますが、それに抵抗するものは、還元できないほどあなたのものです。
もう 1 つのエッジである AI ネイティブ エンジニアリングは急速に成長しており、まさにその分野に飢えています。詳細については後ほど説明します。
真ん中が問題です。 「データセットを取得し、かなり標準的なモデルをトレーニングし、成果物を渡す」というのは、まさに現在エージェントが有能かつ精力的に行っている部分です。それは予測ではありません。それは月曜日の朝です。週の大半をその中間で過ごしているなら、この記事はあなたにぴったりです。このニュースは思っている以上に良いものです。
データの厳密さは資産であり、それが転送されます
これは、心配している ML エンジニアに誰も教えない部分です。あなたが所有する最も価値のあるものは、モデル アーキテクチャではありません。それは反射です。あまりにも良すぎる数字にたじろいでしまいます。あなたはそのホールドアウトがどこから来たのか尋ねます。あなたは、漏れ、間違った理由で移動したメトリクス、トレーニングと静かに重なったテストセットによって火傷を負いました。この本能は構築するのに何年もかかりましたが、AI ネイティブの世界では希少なスキルであり、多くの人が「見栄えの良い」プロンプトを送ってそれを完了したと判断します。
その反射はほぼ 1 対 1 で転送されます。オブジェクトの名前が変わります。規律は同じです。
最初からやり直すわけではありません。あなたは自分の強みの名前を変更し、それを決定論的なシステムではなく確率論的なシステムに向けています。この移行期で苦労しているのは、厳格な人々ではありません。彼らはかつては反射神経を持っていなかった人たちですが、今ではバイブスを発信しています。あなたには反射神経があります。それがゲーム全体だ。
本当にマッピングされないもの (曲線については正直に言ってください)

e)
私は横断歩道が無料であるかのように振る舞うつもりはありません。実際に新しいことがいくつかありますが、それらは最初の 1 か月間費やすべきことです。
第一級の懸念としての非決定論。同じ入力から 2 つの出力を得ることができます。評価では、単一のスコアではなく、分布と合格率を考慮する必要があります。不安定なテストのトリアージを行ったことがある場合は、有利なスタートを切ることができます。
トレーニングよりもオーケストレーション。作業単位は、「モデルをトレーニングする」から「ツール、エージェント、およびコンテキストを持続可能なワークフローに構成する」に移行します。違う筋肉。
LLM を提供します。スループット、KV キャッシュ、バッチ処理、コスト対レイテンシの曲線。 MLOps に隣接していますが、同一ではありません。
エージェントループ自体。エージェントを適切に操作する (いつ実行するか、いつ制約するか、どのように操作するか) ことはスキルであり、最も高い成果が得られるスキルです。ここで私が最も興奮している部分に移ります。
力の乗算器を備えた実験ループ
ここで、リスキルは防御ではなくなり、利点になります。
コアの ML ループは 10 年間変わっていません。仮説を立て、アブレーションを実行し、結果を読み取り、次の実験を決定し、勝ったバージョンを選択します。そのループにおける判断はあなたのものであり、それは苦労して勝ち取るものです。いつも苦痛だったのは、掃除の配線、実行の子守、結果を読めるように表にまとめ、すでに試したことを思い出すなどの機械的な負担でした。
エージェント ワークフローはその負担を解消します。エージェントはスイープを実行し、メトリクスでスコアを付け、表を作成し、どの要因が実際に数値を動かしたかを示し、これまでに確認したすべてのことを考慮して、次の実験の草案を作成します。どの信号が本物なのか、どのリフトがリークなのか、GPU にどれだけの価値があるのか​​を判断するのはあなた自身です。あなたの好みは希少なインプットです。クロード・コードは、たゆまぬラボ技術者です。
そして、これは大丈夫な人だけではありません

LLM を統合します。耐久性の高いエッジ (ランキング モデル、pCTR 予測変数、入札ポリシー) の深層にいる場合、ループは同一です。つまり、特徴の除去、ハイパーパラメーターのスイープ、候補セットの探索、セグメント全体のスライス分析です。同じ機械税で、重要な通話を維持している間、エージェントがそれを負担できます。より深く進むことと AI ネイティブになることは、どちらか一方ではありません。自分の好みとエージェントの探索ループを組み合わせたスペシャリストは、すべてのスイープを手動で実行しているスペシャリストよりも反復的な作業を行います。
これを私自身の作品から実際にアブレーションして具体的に説明しましょう。これはまさに実際にループが実行される方法だからです。
機能したループ: 壊れそうになったアブレーション
Llama 3.1 8B モデル (QLoRA、r=16) を微調整して、船荷証券から 18 個の構造化フィールドを抽出しました。最初のトレーニング実行、標準の配布内テスト セットでスコア付け:
JSON の有効性 100.0%
スキーマ準拠 100.0%
フィールド精度 100.0%
完了しましたね?これは、反射神経がその効果を発揮する瞬間です。 100％冷凍しても勝利ではなく、臭いです。テスト セットはトレーニング セットと似ているので、当然合格します。 ML エンジニアが反射的に抱く質問: ディストリビューション外では何が起こるのか?そこで私はエージェントに敵対的分割を構築してもらいました。同じ 184 レコードを 5 つのレイアウト (表形式、簡潔、物語、ノイズを含む、およびオリジナル) に再レンダリングし、すべての 920 について同じモデルをスコア付けします。1 つのプロンプトが全体を動かし、フィールド精度メトリクス (ID の完全一致、名前のファジー ≥90、数値の ±1%) によってスコア付けされます。
> テスト セットを 5 つのレイアウト バリアントに再レンダリングし、すべてのレイアウト バリアントで v1 モデルを実行します。
920、eval/metrics.py でスコアを付け、結果をレイアウトごとに分類します。
レイアウトと JSON スキーマ フィールドの Acc
オリジナル 184 100.0% 100.0% 100.0%
表形式 184 100.0% 0.0% 54.1%
簡潔 184 100.0% 0.0% 51.6%
物語 184 100.0% 0.0% 87.1%
うるさい 184 100.0%

97.3% 93.7%
全体 920 100.0% 39.5% 77.3%
そこにあります。配布中に 100% のスコアを獲得したモデルは、レイアウトが移動した後も 39.5% のスキーマ準拠を維持します。 tableular および terse ではスキーマ準拠がゼロになります。レイアウトごとの内訳が全体的な診断になります。これは容量の問題やハイパーパラメータの問題ではなく、データの多様性の問題です。トレーニング セットは単一レイアウトであったため、モデルはスキーマを学習するのではなく、レイアウトを記憶しました。
ここで「次は何をするか」のステップになります。ここでループが効果を発揮します。そのテーブルが与えられると、次の実験はそれ自体を書き込みます。つまり、トレーニング データを同じ 5 つのレイアウトに均一に再レンダリングし、再トレーニングします。同じハイパーパラメータ、同じ 1,465 レコード、同じ 6 分間の計算:
レイアウト n JSON スキーマ フィールド acc Δ
オリジナル 184 100.0% 100.0% 100.0% 0.0
表形式 184 100.0% 100.0% 98.0% +43.9
簡潔 184 100.0% 100.0% 100.0% +48.4
ナラティブ 184 100.0% 100.0% 100.0% +12.9
うるさい 184 100.0% 100.0% 99.9% +6.2
全体 920 100.0% 100.0% 99.6% +22.3
負荷を支える変数はデータの多様性であり、アーキテクチャではありません。最終モデルは、わずかなレイテンシとコストでフィールド精度 (99.6% 対 92.4%) で Claude Sonnet 4.5 を上回っています。
そのループで実際に何が起こったのかを読んでください。エージェントは、レイアウトの再レンダリング、920 評価の 2 回の実行、表作成などの機械的な作業を行いました。人間の判断は、100% を信頼することを拒否し、敵対的な分割を構築することを認識し、レイアウトごとのゼロをチューニングの問題ではなくデータの問題として読み取るなど、負荷のかかる作業を行いました。その判断は ML エンジニアが行います。スピードはエージェントのものです。どちらの半分も単独でそこに到達することはできません。その組み合わせは、再スキルを習得する価値のある仕事です。
重要な役割分担
エージェントは、実験の実行においてあなたよりも高速です。どの実験が実行に値するかを知る点ではあなたよりも優れているわけではありません

ng、またはどの結果が嘘をついているか。再スキル化とは、前半部分を推進する方法を学び、不足している部分である判断力がより多くの領域をカバーできるようにすることです。
ループを盗む: プロンプトのパターン
このループは明日実行できます。機能するパターンは次のとおりです。
# エージェントに実行するだけでなく提案させる
「ここに私の過去 12 回のランを表として示します。どの要因がフィールドの精度を変えたのか
一番？ギャップを埋める可能性が最も高い 3 つの実験をランク付けして提案します。
それぞれが機能する理由と、実行にかかる費用も説明します。」
# 敵対本能を強制的にループに入れる
「この結果を信じる前に、この結果からできる限り厳しい抵抗力を構築してください
データ: 実際の入力が考えられる 3 つの方法で分布をシフトし、スコアを再設定します。
単なる集計ではなく、スライスごとの数値を表示してください。」
# 見出しではなく、重要なメトリクスのバージョンを選択してください
「これらのチェックポイントを、平均値ではなく、最悪のスライス フィールド精度でランク付けします。私は気にします」
平均ではなく、最悪のレイアウトを処理することについて。」
エージェントがテーブルを埋めます。電話をかけるのはあなたです。それがループです。
エッジが融合しているもう 1 つの証拠: SecSid
ディープ ML と AI ネイティブの仕事が 2 つの別のキャリアであるとまだ考えている場合は、ここに私自身のベンチからの反例があります。私はレコメンダー システム、つまり TIGER の業務分野からの RQ-VAE セマンティック ID から直接技術を取り入れ、次のことを指摘しました。

[切り捨てられた]

## Original Extract

The undifferentiated middle of ML engineering is getting squeezed. Your data rigor still wins, the experimentation loop gets a force multiplier, and here is a concrete 90-day plan to move.

Skip to content shrikar . com AI Native Development Fine-Tuning Agents Evals Claude Code About RSS X in AI Native Development Fine-Tuning Agents Evals Claude Code About Writing Reports Snippets Case Studies GPU Calculator X / Twitter LinkedIn Essay June 11, 2026 12 min read
From ML engineer to AI-native: reskilling toward an edge
The undifferentiated middle of ML engineering is getting squeezed. Your data rigor still wins, the experimentation loop gets a force multiplier, and here is a concrete 90-day plan to move.
You open the repo on Monday and an agent has already scaffolded the data pipeline you were going to spend Tuesday on. The feature store wiring, the train/val split, the boilerplate around a standard model, the eval script: done, with tests, while you were asleep. It is good work. And somewhere under the relief there is a quieter thought you have not said out loud: if an agent can do the part I was getting paid for, what exactly is my job in a year?
I want to take that feeling seriously, because I have had it. I spent years as a fine-tuning and ML person, and the honest answer is not “relax, nothing changes” and it is not “panic, ML is dead.” It is more specific than either, and once you see the shape, the move is clear.
The undifferentiated middle of ML engineering is getting squeezed: standard pipelines, train a model, hand off a pickle. Two edges stay durable. Deep ML bound to your data and objective (ranking, pCTR, bidding, optimization, calibration) does not go anywhere. AI-native engineering is growing. The risky place to stand is the middle. Pick an end.
Start with what is not going anywhere, because the doom takes always skip it.
Deep ML bound to your objective and your data is as defensible as it has ever been. A ranking model that decides what 200 million users see, a click-through-rate predictor wired into an ad auction, a bidding strategy, a demand forecast, an anomaly detector on payments: these are math over your proprietary data and your business constraints. There is no frontier API you can call that knows your auction dynamics or your label distribution. If you are deep in one of these, the move is not to reskill away from it. It is to go deeper : better calibration, better online/offline gap analysis, a sharper objective. That edge is real, and going deeper does not have to mean iterating slower. The agentic exploration loop further down this post is as much for you as for anyone crossing to the other side.
But here is the honest complication, because the line inside that edge is moving too. The retrieval layer of these systems (content-based recommendation, semantic search, candidate generation, a large slice of classic applied NLP) is steadily collapsing into a general-purpose layer: embedding models, LLMs, and fine-tunes of them. You used to hand-build per-domain features and a bespoke retrieval stack. Increasingly you call a general embedding model, or fine-tune one for your domain, and it beats the artisanal pipeline. That sublayer is generalizing, and fast.
What stays yours is the part bound to your objective and your constraints: ranking under a business goal, calibration, the auction and bidding logic, the online/offline gap, the optimization that turns a score into a decision. A general model can fetch the candidates. It does not know that your marketplace over-serves one seller, or how to trade precision for revenue in your specific auction. That is the durable core, and it is narrower and sharper than “I do recsys.”
So read the left weight precisely. The durable thing is objective-bound modeling, not the retrieval and content-understanding sublayer, which is migrating toward the same general-purpose layer that powers the right edge. That is the pattern under this whole post: the general layer keeps absorbing whatever can be standardized, and what resists is whatever is irreducibly yours.
The other edge, AI-native engineering, is growing fast and is starved for exactly the discipline you have. More on that in a second.
The middle is the problem. “I take a dataset, train a fairly standard model, and hand off an artifact” is precisely the slice that agents now do competently and tirelessly. That is not a prediction. That is Monday morning. If most of your week lives in that middle, this post is for you, and the news is better than it feels.
Your data rigor is the asset, and it transfers
Here is the part nobody tells the anxious ML engineer: the most valuable thing you own is not a model architecture. It is a reflex . You flinch at a number that looks too good. You ask where the holdout came from. You have been burned by leakage, by a metric that moved for the wrong reason, by a test set that quietly overlapped with training. That instinct took years to build and it is the scarce skill in the AI-native world, where a lot of people ship a prompt that “looks good” and call it done.
That reflex transfers almost one-for-one. The objects change names; the discipline is identical.
You are not starting over. You are renaming your strengths and pointing them at a stochastic system instead of a deterministic one. The people who struggle in this transition are not the rigorous ones. They are the ones who never had the reflex and now ship vibes. You have the reflex. That is the whole game.
What genuinely does not map (be honest about the curve)
I am not going to pretend the crossing is free. A few things are actually new, and they are the things to spend your first month on:
Non-determinism as a first-class concern. The same input can give two outputs. Your eval has to think in distributions and pass-rates, not a single score. If you have ever done flaky-test triage, you have a head start.
Orchestration over training. The unit of work shifts from “train a model” to “compose tools, agents, and context into a workflow that holds up.” Different muscle.
Serving LLMs. Throughput, KV cache, batching, the cost-versus-latency curve. Adjacent to MLOps but not identical.
The agentic loop itself. Driving an agent well (when to let it run, when to constrain it, how to instrument it) is a skill, and it is the one with the highest payoff. Which brings me to the part I am most excited about.
The experimentation loop, with a force multiplier
This is where reskilling stops being defense and becomes upside.
The core ML loop has not changed in a decade: form a hypothesis, run an ablation, read the results, decide the next experiment, pick the version that wins. The judgment in that loop is yours and it is hard-won. What was always painful was the mechanical tax around it: wiring the sweep, babysitting the run, tabulating results into something you can read, remembering what you already tried.
Agentic workflows collapse that tax. The agent runs the sweep, scores it with your metric, tabulates it, tells you which factor actually moved the number, and drafts the next experiments to try given everything you have already seen. You stay the judgment: which signal is real, which lift is leakage, what is worth the GPU. Your taste is the scarce input. Claude Code is the tireless lab tech.
And this is not only for people fine-tuning LLMs. If you are deep on the durable edge (a ranking model, a pCTR predictor, a bidding policy), the loop is identical: feature ablations, hyperparameter sweeps, candidate-set exploration, slice analysis across segments. Same mechanical tax, and an agent can carry it while you keep the calls that matter. Going deeper and going AI-native is not an either/or. The specialists who pair their taste with an agentic exploration loop will out-iterate the ones still running every sweep by hand.
Let me make that concrete with a real ablation from my own work, because this is exactly how the loop runs in practice.
A worked loop: the ablation that nearly shipped broken
I fine-tuned a Llama 3.1 8B model (QLoRA, r=16) to extract 18 structured fields from Bills of Lading. First training run, scored on the standard in-distribution test set:
JSON validity 100.0%
Schema compliance 100.0%
Field accuracy 100.0%
Done, right? This is the moment the reflex earns its keep. A frozen 100% is not a victory, it is a smell . The test set looks like the training set, so of course it passes. The question an ML engineer asks reflexively: what happens off-distribution? So I had the agent build an adversarial split: the same 184 records re-rendered into five layouts (tabular, terse, narrative, noisy, plus the original), then score the same model on all 920. One prompt drives the whole thing, scored by my field-accuracy metric (exact match on IDs, fuzzy ≥90 on names, ±1% on numerics):
> Re-render the test set into 5 layout variants, run the v1 model on all
920, score with eval/metrics.py, and break the results down per layout.
Layout n JSON Schema Field acc
original 184 100.0% 100.0% 100.0%
tabular 184 100.0% 0.0% 54.1%
terse 184 100.0% 0.0% 51.6%
narrative 184 100.0% 0.0% 87.1%
noisy 184 100.0% 97.3% 93.7%
overall 920 100.0% 39.5% 77.3%
There it is. The model that scored 100% in-distribution holds 39.5% schema compliance once the layout moves. Schema compliance goes to zero on tabular and terse . That per-layout breakdown is the whole diagnosis: this is not a capacity problem or a hyperparameter problem, it is a data diversity problem. The training set was single-layout, so the model memorized a layout instead of learning the schema.
Now the “what next” step, which is where the loop pays off. Given that table, the next experiment writes itself: re-render the training data into the same five layouts uniformly and retrain. Same hyperparameters, same 1,465 records, same six minutes of compute:
Layout n JSON Schema Field acc Δ
original 184 100.0% 100.0% 100.0% 0.0
tabular 184 100.0% 100.0% 98.0% +43.9
terse 184 100.0% 100.0% 100.0% +48.4
narrative 184 100.0% 100.0% 100.0% +12.9
noisy 184 100.0% 100.0% 99.9% +6.2
overall 920 100.0% 100.0% 99.6% +22.3
The data diversity was the load-bearing variable, not the architecture. The final model beats Claude Sonnet 4.5 on field accuracy (99.6% versus 92.4%) at a fraction of the latency and cost.
Read what actually happened in that loop. The agent did the mechanical work: re-rendering layouts, running 920 evals twice, tabulating. The human judgment did the load-bearing work: refusing to trust the 100%, knowing to build an adversarial split, reading the per-layout zeros as a data problem rather than a tuning problem. That judgment is the ML engineer’s. The speed is the agent’s. Neither half gets there alone, and that combination is the job worth reskilling into.
The division of labor that matters
The agent is faster than you at running experiments. It is not better than you at knowing which experiment is worth running, or which result is lying. Reskilling is learning to drive the first half so your judgment, the scarce half, covers more ground.
Steal the loop: prompt patterns
You can run this loop tomorrow. The patterns that do the work:
# Make the agent propose, not just execute
"Here are my last 12 runs as a table. Which factor moved field accuracy
the most? Propose the 3 experiments most likely to close the gap, ranked,
with the reason each might work and what it would cost to run."
# Force the adversarial instinct into the loop
"Before I trust this result, build the hardest holdout you can from this
data: shift the distribution in 3 ways a real input might, and re-score.
Show me per-slice numbers, not just the aggregate."
# Pick the version on the metric that matters, not the headline
"Rank these checkpoints by worst-slice field accuracy, not mean. I care
about the layout it handles worst, not the average."
The agent fills the table. You make the call. That is the loop.
One more proof that the edges fuse: SecSid
If you still think deep ML and AI-native work are two separate careers, here is a counterexample from my own bench. I took a technique straight out of recommender systems, RQ-VAE Semantic IDs from the TIGER line of work, and pointed

[truncated]
