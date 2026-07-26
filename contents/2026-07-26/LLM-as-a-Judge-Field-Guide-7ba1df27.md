---
source: "https://kraghavan.ca/llm-infrastructure/evaluation/2026/07/25/llm-as-a-judge-field-guide.html"
hn_url: "https://news.ycombinator.com/item?id=49054953"
title: "LLM-as-a-Judge Field Guide"
article_title: "What Is LLM-as-a-Judge, Really? A Field Guide to the State of the Art in 2026 - Karthika Raghavan"
author: "kraghavan"
captured_at: "2026-07-26T05:22:25Z"
capture_tool: "hn-digest"
hn_id: 49054953
score: 1
comments: 0
posted_at: "2026-07-26T05:17:30Z"
tags:
  - hacker-news
  - translated
---

# LLM-as-a-Judge Field Guide

- HN: [49054953](https://news.ycombinator.com/item?id=49054953)
- Source: [kraghavan.ca](https://kraghavan.ca/llm-infrastructure/evaluation/2026/07/25/llm-as-a-judge-field-guide.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T05:17:30Z

## Translation

タイトル: LLM-as-a-Judge フィールド ガイド
記事のタイトル: 裁判官としての LLM とは実際何ですか? 2026 年の最先端技術へのフィールド ガイド - Karthika Raghavan
説明: 実際に論文を読む前に、CAP 定理についてうなずくのではなく、裁判官としての LLM を実際に理解したいと思いました。そこで私は適切な研究パイプラインを構築し、それについて何か真に斬新な意見を見つけようと試み、私の「斬新な」アイデアが 5 つすべて論文によって潰されるのを観察しました。
[切り捨てられた]

記事本文:
裁判官としての LLM とは実際何ですか? 2026 年の最先端技術に関するフィールド ガイド
実際に論文を読む前に、CAP 定理についてうなずくのではなく、裁判官としての LLM を実際に理解したいと思いました。そこで私は適切な研究パイプラインを構築し、それについて何か真に斬新な意見を見つけようと努め、私の「斬新な」アイデア 5 つがすべて、ここ数か月で出版された論文によって潰されるのを観察しました。これが私が途中で学んだすべての領収書です。
ソフトウェア エンジニア • 分散システム • LLM インフラストラクチャ
3. 分野を構築した論文
4. 現在実際に使用されている場所
6. 裁判官が崖から落ちる場所
7. 裁判官が故意に攻撃された場合
8. 実際の解決度に応じてランク付けされた緩和ツールキット
9. 正直な間奏曲: 何か新しいものを見つけようとしました
10. フロンティアの実際の場所
11. 作戦室参照表
結論: 基本は解決しました。フロンティアはそうではありません。
この投稿がどのようにして起こったかについて正直に話させてください。投稿そのものよりも良い話だからです。
私は、裁判官としての LLM を正しく学びたかったのです。「そうです、LLM を使用して別の LLM の出力を採点するのです」ではなく、それがどこから来て、どこで壊れるのか、そして私より賢い人々がそれについて何をしているのかを実際に理解したいと思いました。そして、私は他人の製品発売について反応的なホットテイクを書くのに 1 週​​間を費やし、引用もオリジナルの洞察もゼロであると (正しく) 非難されたので、これを正しくやろうと決心しました。つまり、本物の調査、本物の情報源、そして何か新しいことを主張するつもりなら、それを積極的に殺そうとする誰かに耐えなければなりませんでした。
そこで私は調査エージェントの小さな軍隊を立ち上げました。そのうちの6人が現地に赴き、基礎、バイアス、調整、緩和策、2025年から2026年のフロンティアなどを調査した。

、そして本番環境で実際にどのように使用されるかについて説明します。次に、合成パスにより、そのすべてが現在のツールキットの 5 つの候補「ギャップ」に統合されました。次に、エージェントにそれぞれのギャップを具体的で技術的に根拠のある提案に発展させました。次に、これが実際に私に何かを教えてくれた部分ですが、提案ごとに 3 人の独立した懐疑論者にそれぞれの提案に反論してもらいました。先行技術を見つけ、メカニズムを探り、別の方法で証明されない限り「これは生き残れない」とデフォルトします。
5つの提案はすべて廃案となった。ひとつひとつ。アイデアが悪かったからではなく、メカニズムが健全だったからではなく、誰かがすでにほぼまったく同じものを、場合によっては過去 12 週間以内に公開していたからです。この残骸については、演習全体の中で最も役立つ部分であることが判明したため、この投稿の後半で詳しく説明します。しかし、まずは実際のフィールドガイドです。それが私がここに来た理由だからです。
金曜日までに 50,000 件のエッセイを採点しなければならない教授を想像してみてください。すべてを読むことはできませんし、誰も読むことはできません。そこで彼らはティーチングアシスタントを雇います。それは、頭が良くて、素早く、（極めて）一貫性のある人で、一度ルーブリックを読んで、それをエッセイ #1 とエッセイ #49,999 に同じように適用する人です。 TAは教授ではありません。 TA は間違っていたり、疲れていたり、TA が好むフォントで書く生徒に偏見を持っていたり、本当に素晴らしくて型破りな答えと、何も言わないきちんとした形式の答えを区別できないこともあります。しかし、TA は採点が実際に完了するのに十分な速さであり、ほとんどの場合、成績が何らかの意味を持っていると十分に信頼できます。
LLM-as-a-judge は TA であり、「エッセイ」は LLM の出力です: チャットボットの応答、RLHF 設定ペア、エージェントの軌跡、要約、コード。人間の評価者や固定の回答キーの代わりに、LLM に出力を点単位でスコアリングまたは比較するように促します (「この評価を評価する」

1-10」）、ペアごと（「これら 2 つのどちらが優れているか」）、またはルーブリック（「これは基準 A、B、および C を満たしていますか」）を使用します。審査員は、採点対象のモデルとは異なるモデルである必要はありませんが、後で説明するように、それが非常に重要であることがわかります。
2 つのギャップがこれを強制的に存在させましたが、どちらも後から考えると明らかです。
ギャップ 1: 自由形式のテキストには古い指標は機能しません。 BLEU と ROUGE (古典的な N グラムのオーバーラップ メトリクス) は、比較するためのおおよその「正しい」答えが存在する機械翻訳用に構築されました。要約や対話など、適切な答えの正当な多様性を伴うものでは、それらは崩壊します。 G-Eval (Liu et al.、Microsoft、EMNLP 2023) は、思考連鎖推論と構造化された「フォーム入力」スコアリング テンプレートによって促された GPT-4 が、人間の判断とどの n グラム メトリクスよりもはるかに良く相関することを示しました。要約に関するスピアマン相関は 0.514 で、本当に大きなジャンプでした。
ギャップ 2: 人間の評価は、モデルの出荷速度に比例しません。クラウド ワーカーに RLHF 反復ごと、または毎週のモデル リリースごとに数千の応答ペアを評価させるのは時間がかかり、費用もかかります。 AlpacaFarm (スタンフォード、NeurIPS 2023) とその後継の AlpacaEval は、LLM 自動アノテーターが約 50 分の 1 の低コストで、高い一致率で人間の好みのラベル付けを代替できることを示しました。これは人間の評価者の常備軍なしで実際に RLHF スタイルのメソッド開発を実行するのに十分です。
これらを組み合わせると、重要な点がわかります。それは、反復速度に十分に対応できるほど高速で、信頼できるほど人間と十分に相関関係にある評価器です。 「十分」を強調する — この投稿の多くを「十分」と「常に」の間のギャップに費やします。
3. 分野を構築した論文
このタイムラインの形に注目してください。2022 年から 2023 年は

「これはまったく機能しますか？」（はい、大まかに）。 2024 年以降は「よし、うまくいかないところを直そう」。私たちはまだ第 2 フェーズの真っ最中ですが、減速するどころか加速しています。
4. 現在実際に使用されている場所
裁判官としての LLM は研究目的ではなく、3 つの異なるレーンで負荷に耐えるインフラストラクチャです。
モデルリリースの評価。 Meta の Llama 3 技術レポートには、トレーニング後の人間による評価と並行して実行される、モデルによる評価スコア (正確性、有益性) が記載されています。 OpenAI のオープンソースの評価フレームワークは、「モデルで採点された」テンプレートをファーストクラスの評価プリミティブとして標準化しており、その採点ガイダンスでは、採点対象モデルよりも異なる、通常はより強力なモデルを審査員として使用することを明示的に推奨しています。これは、実践者レベルでの自己選好バイアスの直接的な認識です。
RLHF/RLAIF アライメント パイプライン。これは最も賭け金の高いレーンです。裁判官の評決がトレーニングの合図となります。 Anthropic の憲法 AI パイプラインと Google の RLAIF の取り組みはどちらも、LLM ジャッジを使用して、RL を駆動する報酬モデルをトレーニングする嗜好データを生成します。ここで裁判官が微妙に間違っている場合、モデルはユーザーの実際の好みではなく、裁判官の死角に合わせて静かに、そして大規模に最適化されます。
生産エンジニアリングのパイプライン。これは文化が最も成熟しており、制限について最も誠実な場所です。 Anthropic は、マルチエージェント調査システム内でシングルコールのルーブリックジャッジを実行していますが、ジャッジが見逃す失敗モードに備えて人間のテスターを常駐させています。 Google の Vertex AI AutoSxS には、組み込みのバイアス コントロールとして応答反転とマルチサンプリングが搭載されているほか、信頼する前に自動評価者を人間の好みのデータと照合してチェックするための公式ノートブックも付属しています。 Scale AI の SEAL リーダーボードは、人間が厳選したプライベート ゴールド セット (約 1,000 件) を組み合わせています

アンプル）スケールの LLM グレーディング付き。 Databricks は、単に裁判官が「機能している」と主張するのではなく、実際の検証数値 (クリッペンドルフのアルファは 0.565 ～ 0.698、コーエンのカッパは約 0.64 ～ 0.65) を公開しています。そして、広く支持されている Hamel Husain と Braintrust の実践者のプレイブックは、驚くほど地味です。 30 件までの例に手作業でラベルを付け、それに対する自分の意見の相違に対して審査員のプロンプトを反復し、リッカート スケールよりも 2 値の合否を優先し、人間のレビュー担当者を完全に退職させることはなく、サンプリングによって作業負荷を減らすだけです。
最後のポイントが伝え方です。真剣に裁判官を真実として扱う人は誰もいません。彼らはそれを、人間がラベルを付けた紐を必要とする、高速かつ安価で不完全なプロキシとして扱います。
ここが、「これが最先端だなんて信じられない」という意味で、本当に面白いところです。
位置の偏り。裁判官は体系的に、最初に示された答え（または 2 番目に示された答えが裁判官によって異なります）を支持します。王ら。どちらの答えが最初に現れるかを単純に交換するだけで、80 個のテスト クエリのうち 66 個で Vicuna-13B が ChatGPT を破ることができることが実証されました。どちらのモデルの実際の出力も変わりません。その後、15人の審査員と15万件以上の評価を対象とした大規模な調査で、それが現実であり、審査員と課題に依存し、長さよりも2つの回答の質がどれだけ近いかによって左右されることが確認された。
冗長バイアス。余分な長さが何かを加えるかどうかに関係なく、長い回答はより高い評価を受けます。これは十分に悪いことで、AlpacaEval は「どのモデルがより優れているか」ではなく「どのモデルがより多くランブルするか」を測定するのをやめるためだけに、専用の長さ制御回帰が必要でした。興味深いことに、これは普遍的ではありません。最近の研究では、ジェミニ家とラマ家の裁判官は長い回答を好み、クロード家の裁判官は実際には短い回答を好み、GPT-4o はほぼ中立であることが示されています。 Y

私たちの偏見軽減戦略では、あなたがどの裁判官を担当しているかを知る必要があります。
自己選好バイアス。審査員は自分のモデルファミリーの出力を高く評価します。綿岡氏と高橋氏はこれを困惑と結び付けています。審査員は「見慣れた/予測可能なテキスト」と「良いテキスト」を混同しているようです。これは、品質指標について発見するのはかなりひどいことです。
フォーマット偏り、これがオチです。 2026 年の分析では、同一内容の単純な散文よりもマークダウン形式 (ヘッダー、箇条書き、コード ブロック) に対する審査員の好みが 0.76 ～ 0.92 の効果量で定量化されました。比較のために、同じ分析で測定された位置バイアスは ≤0.04 でした。フォーマット バイアスは、ポジション バイアスと並ぶ小さな混乱ではなく、およそ 20 倍に小さくなります。そして、これは単なる偶然ではありません。2026 年の敵対的論文では、コンテンツの品質とは無関係に、どの書式設定の調整が最も確実に裁判官の評決を覆すかを、バンディット検索を通じて正確に知ることができることを示しています。評価パイプラインが応答を比較する前にマークダウンを削除していない場合、誰が最も美しい箇条書きを書いているかをかなりの程度ベンチマークしていることになります。
より広範な CALM フレームワークの研究では、バンドワゴン効果、権威バイアス、感情バイアス、注意散漫バイアス、思考連鎖バイアスなど、12 の異なるバイアス カテゴリがカタログ化されており、最も強力な裁判官モデルであっても、特定のタスクに関しては測定可能なバイアスが残っていることがわかります。これは、いくつかの既知のフットガンでは解決されていない問題です。それはアクティブな分類法です。
6. 裁判官が崖から落ちる場所
偏見は恥ずかしいものです。ドメインの崩壊はさらに悪いことです。なぜなら、それは裁判官があなたが測定していると思っているものをまったく測定していないことを意味するからです。
要約における事実の一貫性に関して、GPT-3.5 クラスの裁判官は人間の判断との相関関係が 0.3 ～ 0.6 のみである (人間の専門家の場合は 0.8 ～ 0.9)。

高い特異性を維持しながら、事実に矛盾する要約の 40 ～ 70% を完全に見逃します。つまり、明らかに不安定ではなく、自信を持って悪い要約を通過します。
ユーモアに関しては、ほとんどドタバタです。LLM の審査員は、無関係で意味不明な応答を非常に面白いと評価しました (人間が同じ内容を 0.681 と評価するスケールでの平均スコアは 2.18 ～ 3.29)。人間の面白さの評価とスピアマンの相関関係は、Claude Sonnet 4、GPT-4.1、および Gemini 2.5 Pro で 0.169 ～ 0.266 でした。公平を期すために言うと、人間は何が面白いかについてお互いに意見が一致しているわけではありません（ペアごとに一致する割合は 31.7%）。しかし、審査員が人間の合意を追跡した確率は 52 ～ 58% にすぎません。これは、すでに騒がしい標的の上でコイントスを行うよりもかろうじて優れています。
安全性の判断に関しては、この数字は本当に懸念すべきものです。ある調査によると、LLM の安全性判定者は、「運用上の誤用」の危害カテゴリに関してほぼゼロまたはマイナスのクリッペンドルフ アルファに達していることがわかりました。つまり、実際の信号ではなく、ラベルの不均衡のため、生の合意数値が正常に見えたということです。純粋にどのモデルが判定したかに応じて、同一のクエリが 12% ～ 83% の確率で「安全」とラベル付けされました。
臨床/世界保健の内容に関しては、最高の成績を収めた審査員（クロード・オーパスクラス）でさえ、11の評価基準のうち4つだけで人間と同等の成績に達し、英語以外ではさらに低下した。
そして、これらすべてを結びつけるメタポイントがあります。生の相関関係や一致率は、偶然を補正していないため、信頼性を誇張する可能性があります。

[切り捨てられた]

## Original Extract

I wanted to actually understand LLM-as-a-judge instead of nodding along the way I used to with CAP theorem before I’d actually read the paper. So I built a proper research pipeline, tried to find something genuinely novel to say about it, and watched all five of my “novel” ideas get killed by papers
[truncated]

What Is LLM-as-a-Judge, Really? A Field Guide to the State of the Art in 2026
I wanted to actually understand LLM-as-a-judge instead of nodding along the way I used to with CAP theorem before I’d actually read the paper. So I built a proper research pipeline, tried to find something genuinely novel to say about it, and watched all five of my “novel” ideas get killed by papers published in the last few months. Here’s everything I learned along the way — with the receipts.
Software Engineer • Distributed Systems • LLM Infrastructure
3. The Papers That Built the Field
4. Where It’s Actually Used Today
6. Where Judges Fall Off a Cliff
7. When Judges Get Attacked on Purpose
8. The Mitigation Toolkit, Ranked by How Solved It Actually Is
9. The Honest Interlude: I Tried to Find You Something Novel
10. Where the Frontier Actually Is
11. The War Room Reference Table
Conclusion: The Fundamentals Are Settled. The Frontier Is Not.
Let me be honest about how this post happened, because it’s a better story than the post itself.
I wanted to learn LLM-as-a-judge properly — not “yeah, you use an LLM to grade another LLM’s output” properly, but actually understand where it came from, where it breaks, and what people smarter than me are doing about it. And since I’d just spent a week writing a reactive hot-take about someone else’s product launch and gotten (correctly) called out for having zero citations and zero original insight, I decided to do this one right: real research, real sources, and if I was going to claim something novel, it had to survive someone actively trying to kill it.
So I spun up a small army of research agents. Six of them went and surveyed the field — foundations, biases, calibration, mitigations, the 2025-2026 frontier, and how it’s actually used in production. Then a synthesis pass consolidated all of that into five candidate “gaps” in the current toolkit. Then I had agents develop each gap into a concrete, technically-grounded proposal. Then — and this is the part that actually taught me something — I had three independent skeptics per proposal try to refute each one: find the prior art, poke the mechanism, default to “this doesn’t survive” unless proven otherwise.
All five proposals died. Every single one. Not because the ideas were bad — the mechanisms were sound — but because someone had already published almost exactly the same thing, in some cases within the last twelve weeks. I’ll walk through the wreckage later in this post, because it turned out to be the most useful part of the whole exercise. But first, the actual field guide, because that’s what I came here to learn.
Picture a professor with 50,000 essays to grade by Friday. They can’t read all of them — nobody can. So they hire a teaching assistant: someone smart, fast, and (crucially) consistent, who reads a rubric once and then applies it identically to essay #1 and essay #49,999. The TA isn’t the professor. The TA can be wrong, tired, biased toward students who write in a font the TA likes, or unable to tell a genuinely brilliant unconventional answer from a well-formatted one that says nothing. But the TA is fast enough that grading actually finishes, and reliable enough — most of the time — that the grades mean something.
LLM-as-a-judge is that TA, and the “essays” are LLM outputs: chatbot responses, RLHF preference pairs, agent trajectories, summaries, code. Instead of a human rater or a fixed answer key, you prompt an LLM to score or compare outputs — pointwise (“rate this 1-10”), pairwise (“which of these two is better”), or via a rubric (“does this satisfy criteria A, B, and C”). The judge doesn’t need to be a different model than the one being graded, though — as we’ll get to — that turns out to matter a lot.
Two gaps forced this into existence, and both are obvious in hindsight.
Gap one: the old metrics don’t work for open-ended text. BLEU and ROUGE — the classic n-gram overlap metrics — were built for machine translation, where there’s a roughly “correct” answer to compare against. They fall apart on summarization, dialogue, or anything with legitimate diversity of good answers. G-Eval (Liu et al., Microsoft, EMNLP 2023) showed that GPT-4, prompted with chain-of-thought reasoning and a structured “form-filling” scoring template, correlated with human judgment far better than any n-gram metric — 0.514 Spearman correlation on summarization, a genuinely large jump.
Gap two: human evaluation doesn’t scale to how fast models ship. Getting crowdworkers to rate thousands of response pairs for every RLHF iteration or every weekly model release is slow and expensive. AlpacaFarm (Stanford, NeurIPS 2023) and its successor AlpacaEval showed that LLM auto-annotators could substitute for human preference labeling at roughly 50x lower cost , with high agreement — enough to actually run RLHF-style method development without a standing army of human raters.
Put those together and you get the pitch: an evaluator that’s fast enough to keep up with iteration speed, and correlated enough with humans to be trustworthy. Emphasis on “enough” — we’ll spend a lot of this post on the gap between “enough” and “always.”
3. The Papers That Built the Field
Notice the shape of this timeline: 2022-2023 is “does this work at all” (yes, roughly). 2024 onward is “okay, now fix the ways it doesn’t.” We’re still deep in the second phase, and it’s accelerating, not slowing down.
4. Where It’s Actually Used Today
LLM-as-a-judge isn’t a research curiosity — it’s load-bearing infrastructure in three distinct lanes.
Model release evaluation. Meta’s Llama 3 technical report documents model-graded scoring (correctness, informativeness) running alongside human eval during post-training. OpenAI’s open-source evals framework standardizes “model-graded” templates as a first-class evaluation primitive, and their grading guidance explicitly recommends using a different, typically stronger model as judge than the one being graded — a direct, practitioner-level acknowledgment of self-preference bias.
RLHF/RLAIF alignment pipelines. This is the highest-stakes lane: the judge’s verdict becomes the training signal . Anthropic’s Constitutional AI pipeline and Google’s RLAIF work both use LLM judges to generate the preference data that trains the reward model driving RL. If the judge is subtly wrong here, the model gets optimized toward the judge’s blind spots, not the user’s actual preferences — quietly, and at scale.
Production engineering pipelines. This is where the culture is most mature, and most honest about limitations. Anthropic runs a single-call rubric judge inside its multi-agent research system but keeps human testers around for the failure modes the judge misses. Google’s Vertex AI AutoSxS ships response-flipping and multi-sampling as built-in bias controls, plus an official notebook for checking the autorater against human preference data before trusting it. Scale AI’s SEAL leaderboards pair private human-curated gold sets (~1,000 examples) with LLM grading for scale. Databricks publishes its actual validation numbers — Krippendorff’s alpha of 0.565-0.698, Cohen’s kappa around 0.64-0.65 — rather than just claiming the judge “works.” And the widely-followed practitioner playbook from Hamel Husain and Braintrust is refreshingly unglamorous: label ~30 examples by hand, iterate the judge prompt against your own disagreements with it, prefer binary pass/fail over Likert scales, and never fully retire the human reviewer — just shrink their workload via sampling.
That last point is the tell. Nobody serious treats the judge as ground truth. They treat it as a fast, cheap, imperfect proxy that needs a human-labeled leash.
Here’s where it gets genuinely funny, in a “I can’t believe this is the state of the art” way.
Position bias. Judges systematically favor whichever answer is shown first (or second — it’s judge-dependent). Wang et al. demonstrated you could make Vicuna-13B beat ChatGPT on 66 of 80 test queries purely by swapping which answer appeared first — no change to either model’s actual output. A later large-scale study across 15 judges and 150,000+ evaluations confirmed it’s real, judge- and task-dependent, and driven more by how close the two answers are in quality than by anything about length.
Verbosity bias. Longer answers get rated higher, independent of whether the extra length adds anything. This was bad enough that AlpacaEval needed a dedicated length-control regression just to stop measuring “which model rambles more” instead of “which model is better.” Interestingly, this one isn’t universal: recent work shows Gemini- and Llama-family judges prefer longer answers, Claude-family judges actually prefer shorter ones, and GPT-4o sits roughly neutral. Your bias-mitigation strategy needs to know which judge you’re running.
Self-preference bias. Judges rate their own model family’s outputs higher — Wataoka & Takahashi tie this to perplexity : the judge seems to conflate “text I find familiar/predictable” with “text that’s good,” which is a fairly damning thing to discover about your quality metric.
Format bias, and this is the punchline. A 2026 analysis quantified judges’ preference for markdown formatting — headers, bullets, code blocks — over plain prose with identical content , at an effect size of 0.76-0.92 . For comparison, position bias in the same analysis measured ≤0.04 . Format bias isn’t a minor confound sitting next to position bias — it dwarfs it by roughly 20x. And it’s not just incidental: a 2026 adversarial paper shows you can learn, via bandit search, exactly which formatting tweaks most reliably flip a judge’s verdict, independent of content quality. If your eval pipeline isn’t stripping markdown before comparing responses, you are, to a significant degree, benchmarking who writes the prettiest bullet points.
A broader CALM framework study catalogs 12 distinct bias categories — bandwagon effects, authority bias, sentiment bias, distraction bias, chain-of-thought bias, and more — and finds even the strongest judge models retain measurable bias on specific tasks. This is not a solved problem with a couple of known footguns. It’s an active taxonomy.
6. Where Judges Fall Off a Cliff
Bias is embarrassing. Domain collapse is worse, because it means the judge isn’t measuring the thing you think it’s measuring at all.
On factual consistency in summarization , GPT-3.5-class judges show only 0.3-0.6 correlation with human judgment (versus 0.8-0.9 for human experts), and miss 40-70% of factually inconsistent summaries outright — while maintaining high specificity , meaning they confidently pass bad summaries rather than obviously flailing.
On humor , it’s almost slapstick: LLM judges rated irrelevant, nonsensical responses as highly funny (mean scores 2.18-3.29 on a scale where humans rated the same content 0.681) — Spearman correlation with human funniness ratings sat at 0.169-0.266 across Claude Sonnet 4, GPT-4.1, and Gemini 2.5 Pro. To be fair, humans don’t agree with each other about what’s funny either (31.7% pairwise agreement) — but the judges tracked human consensus only ~52-58% of the time, which is barely better than a coin flip on top of an already-noisy target.
On safety judgment , the numbers get genuinely concerning. A study found LLM safety judges reaching near-zero or negative Krippendorff’s alpha on “operational misuse” harm categories — meaning the raw agreement numbers looked fine only because of label imbalance, not real signal — with identical queries labeled “safe” anywhere from 12% to 83% of the time depending purely on which model did the judging.
On clinical/global-health content , even the best-performing judge (Claude Opus-class) reached human-equivalent performance on only 4 of 11 evaluation criteria, and degraded further outside English.
And there’s a meta-point that ties all of this together: raw correlation or percent-agreement can overstate reliability, because it doesn’t correct for chance

[truncated]
