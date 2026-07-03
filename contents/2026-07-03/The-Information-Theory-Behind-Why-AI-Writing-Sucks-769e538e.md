---
source: "https://www.pangram.com/blog/joe-stech-information-theory-why-ai-writing-sucks"
hn_url: "https://news.ycombinator.com/item?id=48776876"
title: "The Information Theory Behind Why AI Writing Sucks"
article_title: "The Information Theory Behind Why AI Writing Sucks | Pangram Labs"
author: "malshe"
captured_at: "2026-07-03T17:14:02Z"
capture_tool: "hn-digest"
hn_id: 48776876
score: 2
comments: 0
posted_at: "2026-07-03T16:29:52Z"
tags:
  - hacker-news
  - translated
---

# The Information Theory Behind Why AI Writing Sucks

- HN: [48776876](https://news.ycombinator.com/item?id=48776876)
- Source: [www.pangram.com](https://www.pangram.com/blog/joe-stech-information-theory-why-ai-writing-sucks)
- Score: 2
- Comments: 0
- Posted: 2026-07-03T16:29:52Z

## Translation

タイトル: AI ライティングがダメな理由の背後にある情報理論
記事のタイトル: AI ライティングがダメな理由の背後にある情報理論 |パングラム研究所
説明: 人間のライターの独特の声と比較すると、AI の文章は驚くほど均一に聞こえます。これには十分な理由があり、それは情報理論に関係していることが判明しました。

記事本文:
AI ライティングがダメな理由の背後にある情報理論
AI ライティングがダメな理由の背後にある情報理論
目次 確率分布としての音声
RLHF トラップと「アノテーター コンセンサス方言」
カモフラージュの錯覚 (スタイルを促すのが失敗する理由)
The failure of temperature and friends
開示: AI 言語モデルは編集プロセス中に使用され、技術的な説明を草案し、構造的および散文的な改善を提案しました。 AI からのいくつかの提案が記事の最終バージョンで使用されました。
私は恥ずかしいほど大量の小説、特にSFを読んできました。私はまた、ソフトウェア エンジニアリングの仕事でリリースされる主力 AI モデルをすべて使用しています。
これら 2 つの経験により、AI は高機能な人間の作者と比較した場合、驚くほど均一な「声」を持っているという、ひりひりとした感覚を私に残しました。
文学を愛する人なら誰でも、私の言っていることを感じたことがあるでしょう。私はこれまで約 5,000 人のさまざまな作家の物語を読んできましたが、たとえ 6 人の作家しか読んだことがなかったとしても、各作家が独自の文体の空間を占めていることに気づくと正直に思います。
人間のライターの独特の声と比較すると、AI の文章は驚くほど均一に聞こえます。これには十分な理由があり、それは情報理論に関係していることが判明しました。
Voice as a probability distribution
独特の作者の「声」はランダムではなく、平均的でもありません。これは特定の確率分布です。これを P_author と呼びます。著者は執筆するとき、非常に特異なプロセスからサンプリングします。コンセプト、ペース、語彙、その他の文体ツールを実装する方法について、特定の条件付き確率があります。
音声を認識できるものにするのは、作者が一貫して行う低周波で影響力の高い選択 (ロングテール

ディストリビューションの場合）。私が「テッド・チャン」と言えば、彼の文章が構文的には平易だが意味的には緻密であることをすぐに思い浮かべるでしょう（これは私が尊敬するスタイルですが、この括弧書きが示すように、私には真似できません）。私が「アーシュラ・K・ル・グウィン」と言ったら、彼女がどのようにしてこれほど明晰で地に足の着いた表現をしていながらも、抒情的な雰囲気を与えることができるのかを考えるでしょう。彼女のスタイルをうまく説明することはできませんが、ル・グウィンの読者なら私の言いたいことはわかっていただけるでしょう。
最終的に私が得ているのは、テキストがどの程度「AI っぽく」聞こえるかを測定する正しい方法は、それが一般的に予測可能かどうかを確認することではなく (ほとんどの有能な文章はある程度予測可能です)、モデルの出力分布と特定の著者の分布の間の KL 乖離、つまり D_KL(P_author || Q_model) を測定することです。 KL 発散に慣れていない人のために説明すると、これは、モデルの分布が作成者の選択をどの程度カバーできていないかを測定します (具体的には、Q に最適化されたコードを使用して P からサンプルをエンコードする際に予想される追加情報コストを測定しています)。この乖離が大きく、構造化されている場合、声が聞こえます。
RLHF トラップと「アノテーター コンセンサス方言」
事前トレーニング中に、大規模な言語モデルが人間のテキストの一般化された分布のマップを生成します。この基本ディストリビューション Q_base は非常に幅広いです。その潜在空間には、ほぼすべての P_author を近似する能力が含まれています。
私が言及した罠は調整から始まります。モデルを安全で有用なものにするために、研究室ではヒューマン フィードバックからの強化学習 (RLHF) などの技術を適用します。詳細は異なりますが、要は、人間 (または AI) の好みから派生した報酬シグナルに対して良好なスコアを示す出力を生成するようにモデルが最適化されているということです。
これはモデルを英語の統計的平均に近づけるものではありません。それは何かに向かってそれを押します

異なる確率分布を持つ ng — これをアノテーター コンセンサス方言と呼びましょう。
そこに到達するメカニズムは次のとおりです。審査員 (成果物を評価するために雇われたギグワーカーや専門家など) が成果物を評価するとき、特異な文章によって評価に大きなばらつきが生じます。私の書き方は、ある評価者からは 5/5 の評価を受けるかもしれませんが、別の評価者からは 2/5 の評価を受けるかもしれません。しかし、不毛で対称的で、厳しくヘッジされた反応は、全体的に 4/5 のスコアを獲得する可能性があります。最適化アルゴリズムでは、期待される報酬を最大化する最も安全な方法は分散を縮小することであると規定しています。会話で言えば、ホテルのロビーの装飾に相当します。
「ジョー、これは公平な特徴付けではありません! 新しい位置合わせ技術は明らかに多様性を維持するように設計されています!」と言うかもしれません。これは事実ですが、新しい手法は依然として「優先」出力の概念に合わせて最適化されており、安全で広く受け入れられる散文と比較して、変動の大きいリスクテイクに依然としてペナルティを課します。
これはテスト可能な主張です (私はテストしていませんが、テスト可能です)。整列されたモデルの出力と、たとえば企業コミュニケーションと文学フィクションのコーパスの間の KL 発散を測定した場合、私の予測では、モデルの分布は企業の中心にはるかに近くなるでしょう。私の知る限り、この正確な測定値を発表した人は誰もいませんが、最適化数学によりそれが強く予測されます。
カモフラージュの錯覚 (スタイルを促すのが失敗する理由)
あなたが考えていることはわかります: はい、しかし、この方言からモデルをプロンプトで出力することはできます。 「1920年代のハードボイルド探偵のスタイルで書く」とかなんとか（モデルにこの記事をルーペ・フィアスコの歌詞として書き直すよう頼んだらどうなるか見てみたいという気持ちもある）。これにより、アノテーターのコンセンサス方言とは異なるように見えるテキストが生成されますが、それでも疑わしいほど均一に感じられます。
これは、

re は、分布の平均をシフトすることと分散構造を再現することの数学的な違いです。
モデルに作者を模倣するよう依頼すると、モデルは重心を移動します。ターゲットの語彙、文構造、その他のスタイル実装の統計的平均を計算し、そこに移動します。ただし、これまで議論してきたのと同じ分散崩壊メカニズムがこの新しい場所に適用されます。
人間のスタイルは構造化された不規則性に依存します。作者には基本的なリズムがありますが、断片につまずいたり、特徴のない動詞を省略したり、感情的な効果を得るために文をもつれさせたりするなどして、意図的にリズムを壊します。計算スタイロメトリーには、これを測定するためのツールがあります。文の長さの時系列のハースト指数は、AI テキストにはない人間の文章の長距離依存関係を明らかにすることができます。人間の作者は、モデルとは異なる方法で語彙の多様性を調整します。
これはすべて、特定のスタイルでの執筆を要求すると、モデルは対象のスタイルの比喩を捉えながらも、爆発的な部分をすべて滑らかにしてくれるということです。あなたが要求したものの似顔絵を生成します。
温度と仲間たちの失敗
AI の分布が狭すぎるのであれば、なぜそれを広げることができないのでしょうか?
最も一般的なアプローチは温度スケーリングです。温度 T を増やすと、確率を計算する前にモデルの生のロジットを T で割ることになります。これにより、分布全体が平坦になり、モデルは可能性の低い単語を選択するようになります。しかし、それは盲目的に行われます。人間の作家の奇行には非常に条件がつきものである。人間は非常に具体的かつ一貫した方法でルールを破りますが、温度のスケーリングは確率的なノイズを引き起こすだけです。
これは直感的に明らかだと思いますが、最終的に温度を上げると、「疑わしいほどスムーズ」から「疑わしい」状態に移行するだけです。

人間をまったく介さずに「巧妙にランダム」。
もっと洗練されたデコード戦略があることは知っています。 Top-p (核) サンプリング、top-k フィルタリング、反復ペナルティ、および分類子なしのガイダンスはすべて、よりターゲットを絞った再配布を試みます。これらはほんのわずかなところでは役に立ちますが、これらは、全体の動作哲学 (そう呼んでよいか) が調整中に形成されたモデル上で動作する推論時の介入であるという根本的な問題を解決するものではありません。
ここには、私の友人の 1 人が最近私に指摘した重要なニュアンスもあります。それは、位置合わせによって、スタイルの差異に対する基本モデルの潜在的な能力が消去されるわけではありません。十分な重みがある限り、事前トレーニングされた重みは Q_base の豊富な機能のほとんどをエンコードします。根底にある潜在空間に到達することで、抑制された分散を部分的に回復できる表現エンジニアリングのような、新たな推論時間ステアリング技術が登場しています。ただし、これらは研究分野であり、公開されている AI 製品で利用できるものではありません。
同様に、長いコンテキストのインコンテキスト学習でも、わずかに良い結果が得られますが、コンテキストが十分に大きくなると、注意のメカニズムは減衰します (コンテキストが大きくなるにつれて、一様な分布に戻り始めます)。
ここでの主なポイントは、RLHF に隣接する技術に組み込まれた設計上の選択により、これらの AI の「音声」が、誰もが認めようとしているよりもはるかに長く検出可能になることを強いられることになるということです。
また、著者のスタイルを特定の高次元の確率分布として考えると便利です。次にお気に入りの著者を読むときに、KL の発散の一部を自分で特定してみてください。作者の声はどこから来るのでしょうか？これは、テキストの楽しみをさらに増やすかもしれない楽しい演習です。

LLM によるスキルの萎縮が見られる昨今、新しい知識を練習して自分の中に取り入れることは良いことです。
Joe Stech は、年次アンソロジー シリーズ「Think Weirder: The Year's Best SF Ideas」の編集者です。また、Arm では開発者およびプラットフォームの有効化に関する主任ソリューション アーキテクトとしても働いています。ここで表明された見解は彼自身のものです。
確率分布としての音声
RLHF トラップと「アノテーター コンセンサス方言」
カモフラージュの錯覚 (スタイルを促すのが失敗する理由)
温度と仲間たちの失敗
Tremau と Pangram Labs が提携して AI 生成コンテンツに取り組む
史上最大の選挙の年を迎える今、テクノロジーと民主主義の交差点が再び注目を集めています。
マーケティング担当者は AI 生成コンテンツに広告費を無駄に費やしている
LLM がより洗練されるにつれ、詐欺師やスパマーは、広告費を盗むことを目的として、生成 AI を使用して本物でないコンテンツを Web 全体に拡大する機会に気づくでしょう。
メタは AI で生成されたコンテンツの識別を開始します
ヨーロッパで可決された新しい AI 規制に対応して、私たちは、大手企業が AI の透明性に関して行動を開始すると予測した分析を発表しました。
従業員スポットライト: AI 研究科学者、キャサリンの紹介
Pangram の創設 AI 研究科学者であるキャサリン・タイ博士をご紹介します。
ICLR 2026 論文で紹介した EditLens テクノロジーに基づくパングラムのオープン ウェイトおよびソースコード利用可能なバージョンのリリースを発表します。
学生が AI の検出を回避しようとする方法
学生は、AI 検出器をだまして、文法や句読点の間違いを導入したり、特定の単語やフレーズを削除したり、文や段落全体を言い換えたりするために、AI で生成されたエッセイを編集する可能性があります。

## Original Extract

Compared to the unique voices of human writers, AI writing sounds remarkably uniform. It turns out that there's a good reason for this, and it has to do with information theory.

The Information Theory Behind Why AI Writing Sucks
The Information Theory Behind Why AI Writing Sucks
Table of contents Voice as a probability distribution
The RLHF trap and the "Annotator Consensus Dialect"
The illusion of camouflage (why prompting for style fails)
The failure of temperature and friends
Disclosure: An AI language model was used during the editing process to draft technical descriptions and suggest structural and prose improvements. Several suggestions from AI were used in the final version of the article.
I have read an embarrassingly large amount of fiction, especially science fiction. I also use every flagship AI model that is released for my software engineering job.
Those two sets of experiences left me with a gnawing feeling that AI has a shockingly uniform "voice" when compared to a high-functioning human author.
Anyone with a love for literature has felt what I'm talking about. I've read stories by about five thousand different authors, but I honestly think that even if you've only read a half-dozen authors you'll notice that each author occupies their own stylistic space.
Compared to the unique voices of human writers, AI writing sounds remarkably uniform. It turns out that there's a good reason for this, and it has to do with information theory.
Voice as a probability distribution
A unique authorial "voice" is not random, and it is not average. It is a specific probability distribution — let's call it P_author. When an author writes, they sample from a highly idiosyncratic process. They have specific conditional probabilities for how they implement concepts, pacing, vocabulary, and other stylistic tools.
What makes a voice recognizable are the low-frequency, high-impact choices that an author makes consistently (the long tail of the distribution). If I say "Ted Chiang", you'll immediately think about how syntactically plain but semantically dense his sentences are (it's a style I admire, but as this parenthetical demonstrates, I cannot emulate). If I say "Ursula K. Le Guin", you'll think about how she can be so clear and grounded but still give a lyrical feel — I can't really describe her style well, but readers of Le Guin know what I mean.
Ultimately what I'm getting at is that the right way to measure how "AI-like" a text sounds is not to check whether it's predictable in general — most competent writing is somewhat predictable — but to measure the KL divergence between the model's output distribution and a specific author's distribution: D_KL(P_author || Q_model). For those unfamiliar with KL divergence, this measures how badly the model's distribution fails to cover the author's choices (to be specific, it's measuring the expected extra information cost of encoding samples from P using a code optimized for Q). When this divergence is large and structured, you hear a voice.
The RLHF trap and the "Annotator Consensus Dialect"
During pre-training, a large language model generates a map of a generalized distribution of human text. This base distribution, Q_base, is enormously wide. In its latent space it contains the capacity to approximate almost any P_author.
The trap I mention begins with alignment. To make the model safe and useful, labs apply techniques like Reinforcement Learning from Human Feedback (RLHF) and others. The details vary, but the bottom line is that the model is optimized to produce outputs that score well against a reward signal derived from human (or AI) preferences.
This does not push the model toward the statistical average of English. It pushes it toward something with a different probability distribution — let's call this the Annotator Consensus Dialect.
The mechanism to get there is this: when the judges (gig workers hired to evaluate outputs or experts or whoever) evaluate outputs, idiosyncratic writing creates high variance in ratings. My style of writing might score 5/5 from one rater and 2/5 from another. But a sterile, symmetrical, heavily hedged response might score 4/5 across the board. The optimization algorithm dictates that the safest way to maximize expected reward is to collapse variance. It is the conversational equivalent of hotel lobby decor.
You might say "Joe, this isn't a fair characterization! Newer alignment techniques are explicitly designed to preserve diversity!". While this is true, the newer methods still optimize for a notion of "preferred" output, which still penalizes high-variance risk-taking relative to safe, broadly acceptable prose.
This is a testable claim (I haven't tested it, but it's testable). If you measured the KL divergence between aligned model outputs and a corpus of, say, corporate communications versus literary fiction, my prediction is that the model's distribution would sit far closer to the corporate center. To my knowledge, no one has published this exact measurement, but the optimization math strongly predicts it.
The illusion of camouflage (why prompting for style fails)
I know what you're thinking: yeah, but you can prompt the model out of this dialect . "Write in the style of a 1920s hard-boiled detective" or whatever (part of me wants to see what this article would read like if I asked a model to rewrite it as Lupe Fiasco lyrics). This does produce text that looks different from the Annotator Consensus Dialect, but it still feels suspiciously uniform.
This is because there is a mathematical difference between shifting a distribution's mean and reproducing its variance structure.
When you ask a model to mimic an author, it shifts its center of mass. It calculates the statistical average of the target's vocabulary, sentence structure, and other style implementations, and moves there. But it applies the same variance-collapsed mechanics we've been discussing to this new location.
Human style relies on structured irregularity. An author has a baseline rhythm, but they break it intentionally by doing things like stumbling into a fragment, dropping an uncharacteristic verb, or tangling a sentence for emotional effect. Computational stylometry has tools for measuring this: Hurst exponents on sentence-length time series can reveal long-range dependencies in human writing that AI text lacks. Human authors modulate their lexical diversity in ways that models don't.
All this is to say that when you ask for writing in a particular style, the model captures the tropes of the target style but smooths out all the burstiness. It generates a caricature of what you asked for.
The failure of temperature and friends
If the AI's distribution is too narrow, why can't we just widen it?
The most common approach is temperature scaling. When you increase the temperature T, you divide the model's raw logits by T before computing probabilities, which flattens the entire distribution and forces the model to pick less likely words. But it does this blindly. A human author's eccentricity is highly conditional. Humans break the rules in very specific, consistent ways, whereas temperature scaling just introduces stochastic noise.
Hopefully this is pretty intuitively obvious — ultimately increasing temperature just transitions you from "suspiciously smooth" to "suspiciously random" without passing through human at all.
I know there are more sophisticated decoding strategies. Top-p (nucleus) sampling, top-k filtering, repetition penalties, and classifier-free guidance all attempt more targeted redistribution. They do help at the margins, but none of them solve the fundamental problem that these are inference-time interventions operating on a model whose whole operating philosophy (if you can call it that) was shaped during alignment.
There is also an important nuance here that one of my friends recently pointed out to me: alignment does not erase the base model's latent capacity for stylistic variance. The pre-trained weights still encode most of the richness of Q_base, as long as you've got enough weights. There are emerging inference-time steering techniques like Representation Engineering that can partially recover the suppressed variance by reaching into the underlying latent space. These are research areas though and not something available in the public AI products.
Similarly, long-context in-context learning can also provide slightly better results, but attention mechanisms attenuate when context gets big enough (and you will start to drift back to the uniform distribution as the context grows).
The main takeaway here is that design choices that go into RLHF-adjacent techniques are going to force these AI "voices" to be detectable far longer than anyone wants to admit.
Also, it's useful to think of an author's style as a specific high-dimensional probability distribution, and I'd challenge you to try and identify some of the KL divergence yourself the next time you're reading your favorite author. Where does the author's voice come from? It's a fun exercise that might increase your enjoyment of the text, and the difficult process of practicing and internalizing new knowledge is a good one to perform in these days of LLM-induced skill atrophy.
Joe Stech is the editor of the yearly anthology series Think Weirder: The Year's Best Science Fiction Ideas . He also works as a Principal Solutions Architect on developer and platform enablement at Arm. Views expressed here are his own.
Voice as a probability distribution
The RLHF trap and the "Annotator Consensus Dialect"
The illusion of camouflage (why prompting for style fails)
The failure of temperature and friends
Tremau and Pangram Labs partner to take on AI-generated content
As we stand on the cusp of the biggest election year in history, the intersection of technology and democracy takes centre stage once again.
Marketers are Wasting Advertising Spend on AI-generated Content
As LLMs become more sophisticated, fraudsters and spammers will realize the opportunity to use generative AI to scale inauthentic content across the web with the intention of stealing advertising dollars.
Meta will start identifying AI generated content
In response to new AI regulations passed in Europe, we published our analysis where we predicted that leading companies would start to take action on AI transparency.
Employee Spotlight: Meet Katherine, AI Research Scientist
Meet Katherine Thai, Ph. D. , Pangram's Founding AI Research Scientist!
Announcing the release of open weights and source-code available versions of Pangram based on the EditLens technology we introduced in our ICLR 2026 paper.
How Students Try to Avoid AI Detection
Students may edit an AI generated essay in efforts to fool AI detectors, introducing grammar and punctuation errors, removing certain words and phrases, and paraphrasing whole sentences and paragraphs.
