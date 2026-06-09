---
source: "https://flatnine.co/blog/i-built-my-own-eval"
hn_url: "https://news.ycombinator.com/item?id=48460776"
title: "Show HN: JazzBench, an LLM reasoning benchmark using jazz improvisation"
article_title: "I built my own Claude eval - FlatNine Blog"
author: "mikerubini"
captured_at: "2026-06-09T16:07:12Z"
capture_tool: "hn-digest"
hn_id: 48460776
score: 2
comments: 0
posted_at: "2026-06-09T13:19:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: JazzBench, an LLM reasoning benchmark using jazz improvisation

- HN: [48460776](https://news.ycombinator.com/item?id=48460776)
- Source: [flatnine.co](https://flatnine.co/blog/i-built-my-own-eval)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T13:19:09Z

## Translation

タイトル: Show HN: JazzBench、ジャズの即興演奏を使用した LLM 推論ベンチマーク
記事タイトル: 私は独自のクロード評価値を構築しました - FlatNine Blog
説明: AI ウィークで、Gian Segato は、Anthropic 内の人々がクロードに対する個人的な評価を保持していると述べました。これは、彼らがどのリーダーボードよりも信頼している小規模でプライベートなテストです。そのアイデアが私にずっと残っていたので、私が自分の好みを最も信頼しているもの、つまりチャーリー・パーカーのソロから私のアイデアを構築しました。

記事本文:
AI ウィークで、Anthropic の Gian Segato が率直な発言をしたのですが、それが私には忘れられませんでした。彼は、Anthropic 内の多くの人々がクロードに対する個人的な評価を保持していると述べました。大きな公的ベンチマークではありません。新しいモデルが実際に優れているかどうかを判断するために、どのリーダーボードよりも信頼できる、個人的に関心のある内容に合わせて調整された小規模なプライベート テストです。
この一文で、モデルの品質についての私の考え方が再整理されました。それで私は行って自分のものを作りました。
個人的な評価が雰囲気に勝る理由
私たちのほとんどはモデルを感触で判断します。新しいリリースがリリースされ、お気に入りのハードなプロンプトをいくつか投げて、雰囲気をつかんで、次に進みます。問題は、バイブがスケールせず、バイブが漂ってしまうことです。 3 か月前に、前のモデルがまったく同じプロンプトをどのように処理したかを思い出せないため、「よりスマートに感じられる」ということは、多くの無駄な作業を行っていることになります。
公開ベンチマークには逆の問題があります。それらは厳格ですが、あなたのものではありません。彼らは一般的なものを測定し、トレーニングセットに漏洩し、ゲームにさらされます。公開リーダーボードでスコアが上がっていれば何かがわかりますが、実際に気にしていることがわかることはめったにありません。
個人的な評価によって差が分かれます。すべてのサンプルを所有できるほど小さいです。その数字があなたにとって特別に意味を持っているという意見は十分にあります。そして、それは十分にプライベートであるため、モデルはトレーニングされていません。スコアが動くと、本当のことを学ぶことができます。
私は、ほとんど誰の判断よりも自分自身の判断を信頼する唯一の領域、つまりジャズから自分のものを作ることにしました。私はそれを JazzBench と名付けましたが、その真実の正体はチャーリー・パーカーです。
このタスクを一文で表すと、パーカーのソロの最初のいくつかのコードと、次に来るコード変更を考慮して、パーカーが次の各コードで実際に演奏した音を予測します。その後、その予測を再度採点します

彼がレコード上で実際に何をしたかを説明します。
言語モデルにそれを依頼するのは奇妙なことであり、まさにそれが私がそれを好む理由です。
世の中のほぼすべての評価では、口頭、数学的、またはコーディングによる推論がテストされます。これらは難しい問題です。正解があります。ほとんどの場合、注意して間違いを犯さないようにすれば、正解にたどり着くことができます。即興演奏は別の種類の認識です。それは次のとおりです。
境界あり。コードチェンジ、キー、タイムはすべて固定です。ただ何かをプレイすることはできません。
判断可能。私たちは、パーカーの実際のソロをグラウンドトゥルースとして、さらに、推測にどれだけ近いかを採点するための正式な音楽理論の手法を持っています。
認知的に豊か。それは、制約の満足度、スタイル、創造性をリアルタイムで同時に実現するものであり、唯一の正解はなく、明らかに間違っている答えがたくさんあります。
その組み合わせは珍しいですね。これは、人間の専門家が直観的に行う、柔軟で複数の制約がある判断であり、ほとんどのベンチマークは測定しようとすることさえありません。モデルがプレッシャーの下でもセンスがあるかどうかを知りたい場合、これは他の文章題よりもはるかに優れた調査方法です。
単一の正解はないため、完全に一致するかどうかを単に確認することはできません。したがって、各予測は、5 つの音楽理論の指標を使用して、パーカーの演奏に対してスコア付けされます。
PC Jaccard: モデルが予測した音とパーカーが実際に演奏した音の間の重複。
インターバルベクトル距離: 正確にどのノートを共有するかだけでなく、インターバル空間内でこれら 2 つのノートセットがどのくらい離れているか。
複雑さと不協和音のデルタ: パーカー自身の複雑さと不協和音の尺度に関する誤差であるため、推測は「間違った音符、正しいテクスチャー」になる可能性があります。
フォルテクラス一致: 転置に関係なく、予測されたセットがパーカーのものと同じ抽象形状を持つかどうか。
そして、それはビートのための 3 つのベースラインを同梱しています: パーカーがその上で使用する傾向にあったノートからランダムにサンプリングします。

コードを演奏し、常にそのコードの最も一般的な単一のセットと、前のセグメント上の 1 次マルコフ モデルを再生します。バーはシンプルです。フロンティアモデルが「最も一般的なものを演奏するだけ」に勝てない場合、それは実際には即興演奏ではなく、平均化しています。
Haiku 4.5、Sonnet 4.6、Opus 4.7 を実行しました。これらのベースラインと比較してそれらがどこに着地するかを観察すると、発売日に読んだほとんどの内容よりもモデルについて多くのことがわかります。
一行バージョン: すべてのクロード層 (Haiku 0.370、Opus 0.400、Sonnet 0.402) は、パーカーとのピッチクラスの重複に関してすべてのベースラインを上回っていましたが、Sonnet と Opus は統計的に同点であり、音程のテクスチャーまたは不協和音の近接性に関して単純なモーダル PC 設定のベースラインと一致するものはありませんでした。クロードはパーカー音符の語彙を学習しましたが、彼の特徴である倍音抑制は学習していません。
見出し、音符の重複 (Jaccard、高いほど良い): ソネット 0.402、作品 0.400、俳句 0.370、対最多頻度 0.355、マルコフ 0.327、およびランダム 0.319。クロードの各層は、パーカーの実際のメモを選択する際にすべてのベースラインを上回り、エージェントの 399 回の通話で解析エラーはゼロでした。しかし、興味深いのはテクスチャにあり、私が予想していなかった 5 つの発見がありました (全文は論文のセクション 8 にあります)。
すべてのクロード層は、ノート選択のすべてのベースラインを上回ります。 Sonnet は「最も一般的なセットをプレイするだけ」より約 13%、ランダムより 26% 上に着地しており、同じ順序が複雑さにも当てはまります。モデルたちはブラフではなく、純粋により良い音を選んでいます。
最も愚かなベースラインは、テクスチャに関しては無敵です。 「常にこのコードのモーダルセットを演奏する」は、構造上、パーカーの音程分布の中心に位置するため、音程の距離と不協和音に関しては依然として優れています。クロードはより頻繁に正しい音符を選択しますが、その間違った推測はパーカーのテクスチャの中心からわずかに外れます。
息子

netとOpusは同点です。 0.002 ジャカード ギャップはノイズです。 Sonnet を超えて拡張してもここでは何も得られません。これは、この設定の下で約 0.40 の上限があるか、生の容量の問題ではなく代表的なボトルネックのいずれかを示しています。
完全一致率は、モデルが大きくなるにつれて低下します。俳句 0.034、ソネット 0.026、作品 0.015。大きなモデルは平均してパーカーとより多くの音を共有しますが、彼の正確なセットを再現することはほとんどありません。一方、小さなモデルは高音域の選択に頼っており、場合によってはそれを完全に釘付けにします。平均的な重複と完全一致の間の逆転は実際の話です。
クロードはパーカーよりも不協和音を奏でます。不協和音誤差は、ベースラインの 1.96 に対して、すべての層で 2.09 ～ 2.43 にとどまります。パーカーのまばらな三極性の好みは、どのモデルも学んだことがないものです。
それはどれも、リーダーボードが私に教えてくれるものではありません。それが定規を所有することの要点です。
リポジトリには 2 つのパスがあります。 Anthropic API に直接アクセスする再現可能なパスがあり、これが数値を再現する正規の方法です。また、既存のセッションを使用して、クロード コード スキルとして実行されるキーなしのパスもあります。
/jazzbench-run # 俳句、10 タスク
/jazzbench-run ソネット 20
/jazzbench-run 作品 5
/jazzbench-run all 10 # 俳句 + ソネット + 作品、次々に
このスキルはタスクごとに 1 つのワークフロー サブエージェントを生成し、各サブエージェントに厳密な予測スキーマでの回答を強制し、同じスコアリング スクリプトで評価できるように結果を書き出します。 2 つのパスは比較表内で別の行としても表示されるため、生の API を介したモデルとクロード コード内でモデルの動作が異なるかどうかを確認できます。
すべてはここにあります: github.com/code91/claude-impro-eval 。また、メソッド、メトリクス、結果を短い論文「JazzBench v0」として書きました。すべてはチャーリー・パーカーの時系列に基づいています

以前にまとめたパイプラインは、転写されたソロを音程ベクトルとコードごとのピッチ クラスに変換するという地味な作業を行います。
実際のポイント: 自分で構築する
ジャズは偶然です。 Gian のコメントから得られる本当のポイントは、モデルが優れているかどうかの感覚をリーダーボードに委託すべきではないということです。これらのモデルに最も近い人々は独自の私的支配者を保持しており、あなたもそうすべきです。
複雑である必要はありません。本物の専門知識を持った分野を 30 分間取り上げ、「良い」とはどのようなものかを知っているいくつかのケースを書き留めて、それらを得点できるものに変えます。今では、新しいモデルがリリースされるたびに、次の四半期には覚えていない雰囲気ではなく、自分の好みで構築された壁にぶつかります。
これは、FlatNine の構築に関する一般的な考え方につながります。耐久性のある資産は決してダッシュボードやスコアではなく、その下にあるコンテキストと判断です。 eval は単なる判断であり、マシンがチェックし続けることができる形式で書き留められます。
これは v0 なので、そのエッジについて明確にしておきたいと思います。ピッチクラスのみが考慮されるため、リズム、アーティキュレーション、音域はまったく採点されません。これらはパーカーパーカーを構成する大きな部分を占めています。これは単一の時代の単一のアーティストであるため、他の即興演奏者についてはまだ何も語られていません。モデルは記号データのみを認識し、音声は認識しません。そして、マルコフ基準線は一次であり、これは簡単なバージョンです。これらはすべて今後の課題であり、本質は変わりません。実際に信頼できる小さな評価は、信頼していない大きな評価に勝つのです。
それがプロジェクトです。自分のものを作りに行きましょう。
受信箱で次の投稿を受け取る
私は 10 以上のブートストラップされたソフトウェア製品を実行し、それらの構築について書いています。頻度は低いですが、スパムはありません。
収益性の高いソフトウェア ビジネスの構築
© 2026 フラットナインブログ . Flat9 によって提供されます。無断転載を禁じます。

## Original Extract

At AI week, Gian Segato mentioned that people inside Anthropic keep their own personal eval for Claude: a small, private test they trust more than any leaderboard. That idea stuck with me, so I built mine out of the thing I trust my own taste on most: Charlie Parker solos.

At AI week, Gian Segato from Anthropic said something offhand that I have not been able to put down. He mentioned that a lot of people inside Anthropic keep their own personal eval for Claude. Not the big public benchmarks. A small, private test, tuned to something they personally care about, that they trust more than any leaderboard to tell them whether a new model is actually better.
That one sentence reorganized how I think about model quality. So I went and built mine.
Why a personal eval beats vibes
Most of us judge models by feel. A new release drops, you throw a few of your favorite hard prompts at it, you get a vibe, you move on. The problem is that vibes do not scale and vibes drift. You cannot remember how the last model handled the exact same prompt three months ago, so "it feels smarter" is doing a lot of unearned work.
Public benchmarks have the opposite problem. They are rigorous, but they are not yours. They measure something generic, they leak into training sets, and they get gamed. A score going up on a public leaderboard tells you something, but rarely the thing you actually care about.
A personal eval splits the difference. It is small enough that you own every example. It is opinionated enough that the number means something to you specifically. And it is private enough that no model has been trained on it. When the score moves, you learn something real.
I decided to make mine out of the one domain where I trust my own judgment more than almost anyone's: jazz. I called it JazzBench, and the ground truth is Charlie Parker.
The task, in one sentence: given the first few chords of one of Parker's solos, plus the chord changes coming next, predict the actual notes Parker played over each of those next chords. Then score that prediction against what he really did on the record.
It is a strange thing to ask a language model to do, which is exactly why I like it.
Almost every eval out there tests verbal, mathematical, or coding reasoning. Those are hard-edged problems: there is a correct answer, and you mostly reach it by being careful and not making mistakes. Improvisation is a different kind of cognition. It is:
Bounded. The chord changes, the key, and the time are all fixed. You cannot just play anything.
Judgeable. We have Parker's actual solo as ground truth, plus formal music-theory methods to score how close a guess is.
Cognitively rich. It is constraint satisfaction, style, and creativity all at once, in real time, with no single right answer but plenty of obviously wrong ones.
That combination is rare. It is the kind of soft, multi-constraint judgment that human experts make intuitively and that almost no benchmark even tries to measure. If I want to know whether a model has taste under pressure, this is a far better probe than another word problem.
Because there is no single correct answer, you cannot just check for an exact match. So each prediction is scored against what Parker played using five music-theoretic metrics:
PC Jaccard: the overlap between the notes the model predicted and the notes Parker actually played.
Interval-vector distance: how far apart those two note sets are in interval space, not just which exact notes they share.
Complexity and dissonance deltas: the error on Parker's own complexity and dissonance measures, so a guess can be "wrong notes, right texture".
Forte-class match: whether the predicted set has the same abstract shape as Parker's, regardless of transposition.
And it ships with three baselines to beat: sampling randomly from the notes Parker tended to use over that chord, always playing the single most common set for that chord, and a first-order Markov model over the previous segment. The bar is simple. If a frontier model cannot beat "just play the most common thing", it is not really improvising, it is averaging.
I ran Haiku 4.5, Sonnet 4.6, and Opus 4.7 through it. Watching where they land relative to those baselines tells me more about a model than most of what I read on launch day.
The one-line version: every Claude tier (Haiku 0.370, Opus 0.400, Sonnet 0.402) beat every baseline on pitch-class overlap with Parker, but Sonnet and Opus are statistically tied, and none of them matched the simple modal-PC-set baseline on interval texture or dissonance proximity. Claude has learned Parker's note vocabulary, but not his characteristic harmonic restraint.
The headline, on note overlap (Jaccard, higher is better): Sonnet 0.402, Opus 0.400, Haiku 0.370, against most-frequent 0.355, Markov 0.327, and random 0.319. Every Claude tier beats every baseline at picking Parker's actual notes, with zero parse errors across 399 agent calls. But the interesting stuff is in the texture, and there are five findings I did not expect (the full writeup is in the paper , section 8):
Every Claude tier beats every baseline on note choice. Sonnet lands about 13% above "just play the most common set" and 26% above random, and the same ordering holds on complexity. The models are genuinely picking better notes, not bluffing.
The dumbest baseline is unbeatable on texture. "Always play the modal set for this chord" still wins on interval distance and dissonance, because by construction it sits at the dead center of Parker's interval distribution. Claude picks the right notes more often, but its wrong guesses fall slightly off Parker's textural center.
Sonnet and Opus are a tie. A 0.002 Jaccard gap is noise. Scaling past Sonnet buys nothing here, which points to either a ceiling around 0.40 under this setup or a bottleneck that is representational, not a matter of raw capacity.
Exact-match rate falls as models get bigger. Haiku 0.034, Sonnet 0.026, Opus 0.015. Bigger models share more notes with Parker on average but rarely reproduce his exact set, while smaller models fall back on high-frequency choices that occasionally nail it outright. That inversion between average overlap and exact match is a real story.
Claude plays more dissonant than Parker. Dissonance error sits at 2.09 to 2.43 across all tiers, versus 1.96 for the baseline. Parker's sparse, triadic preference is something none of the models have learned.
None of that is what a leaderboard would have told me. That is the whole point of owning the ruler.
The repo gives you two paths. There is a reproducible path that hits the Anthropic API directly, which is the canonical way to reproduce the numbers. And there is a no-key path that runs as a Claude Code skill, using your existing session:
/jazzbench-run # haiku, 10 tasks
/jazzbench-run sonnet 20
/jazzbench-run opus 5
/jazzbench-run all 10 # haiku + sonnet + opus, one after another
The skill spawns one workflow subagent per task, forces each one to answer in a strict prediction schema, and writes the results out so the same scoring scripts can grade them. The two paths even show up as separate rows in the comparison table, so you can check whether a model behaves differently through the raw API versus inside Claude Code.
Everything is here: github.com/code91/claude-impro-eval . I also wrote up the method, the metrics, and the results as a short paper: JazzBench v0 . It all builds on a Charlie Parker time-series pipeline I put together earlier, which does the unglamorous work of turning transcribed solos into interval vectors and pitch classes per chord.
The actual point: build your own
The jazz is incidental. The real takeaway from Gian's comment is that you should not outsource your sense of whether a model is good to a leaderboard. The people closest to these models keep their own private rulers, and so should you.
It does not have to be elaborate. Take 30 minutes of a domain where you have genuine expertise, write down a handful of cases where you know what "good" looks like, and turn them into something you can score. Now every new model release runs into a wall built out of your own taste, instead of a vibe you will not remember next quarter.
This connects to how I think about building FlatNine generally: the durable asset is never the dashboard or the score, it is the context and the judgment underneath it. An eval is just judgment, written down in a form a machine can keep checking for you.
This is a v0, and I want to be clear about its edges. It only looks at pitch class, so rhythm, articulation, and register are not scored at all, and those are a huge part of what made Parker Parker. It is a single artist in a single era, so it says nothing yet about other improvisers. The models see only symbolic data, never audio. And the Markov baseline is first-order, which is the easy version. All of that is future work, and none of it changes the point: a small eval you actually trust beats a big one you do not.
That is the project. Go build yours.
Get the next post in your inbox
I run 10+ bootstrapped software products and write about building them. Infrequent, never spam.
Building profitable software businesses
© 2026 FlatNine Blog . Powered by Flat9 . All rights reserved.
