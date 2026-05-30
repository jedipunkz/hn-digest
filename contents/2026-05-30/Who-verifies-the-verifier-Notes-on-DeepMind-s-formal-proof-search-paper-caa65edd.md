---
source: "https://korbonits.com/blog/2026-05-28-who-verifies-the-verifier/"
hn_url: "https://news.ycombinator.com/item?id=48332496"
title: "Who verifies the verifier? Notes on DeepMind's formal proof-search paper"
article_title: "Who Verifies the Verifier — Alex Korbonits"
author: "korbonits"
captured_at: "2026-05-30T11:35:13Z"
capture_tool: "hn-digest"
hn_id: 48332496
score: 1
comments: 0
posted_at: "2026-05-30T04:13:08Z"
tags:
  - hacker-news
  - translated
---

# Who verifies the verifier? Notes on DeepMind's formal proof-search paper

- HN: [48332496](https://news.ycombinator.com/item?id=48332496)
- Source: [korbonits.com](https://korbonits.com/blog/2026-05-28-who-verifies-the-verifier/)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T04:13:08Z

## Translation

タイトル: 検証者を検証するのは誰ですか? DeepMind の正式な証明調査論文に関するメモ
記事のタイトル: 誰が検証者を検証するのか — Alex Korbonits
説明: AI は、私が数学に必要だと私が言ったマシン、つまり専門家の週末の代わりに数セントで証明を検証するコンパイラーを構築しました。キャッチはまだできることです

記事本文:
誰が検証者を検証するのか — Alex Korbonits Alex Korbonits
GitHub 2026 年 5 月 28 日 · 14 分で読みました
まだ開かれたフロンティア
3日前、私は解決策が見つからない問題で終わったエッセイを出版しました。短いバージョンでは、AI がエルデシュの予想の 1 つを反証し、その結果は真実であり、検証されました。しかし、検証された唯一の方法は、現在私たちが知っている唯一の方法です。それは、世界の関連専門家 9 人 (そのうちの 1 人はフィールズ賞メダリスト) が週末に 100 ページにわたる議論を読み、そこに自分の名前を載せるというものです。それはスケールが合わない。もっともらしい証明の供給は放物線的になりつつあり、専門家の週末の供給は固定されているため、拘束制約は証明を生成するのではなく、証明をチェックするだけです。私は、形式的検証 (Lean のような言語で機械チェックされた証明) だけが明らかにスケールアップできる唯一のものであると主張しました。そして、誰もその証明を形式化しておらず、形式化をルーチン化するツールが存在しないという問題点を認めました。 「我々は、関係者全員がそれが成り立たないことが分かるプロセスに依存している」と私は書いた。これについては、先週 OpenAI の発表後に書いたばかりではありません。 2023年から考えています。
私が知らなかったのは、公開を開始する 2 日前に、Google DeepMind が私が説明していたマシンを公開していたことです。
この論文は、5 月 21 日に提出された「AI 主導の形式的証明検索による数学研究の進歩」です。見出しの数字はスクリーンショットに撮られたようなものです。エージェントは、353 個の未解決のエルデシュ問題のうち 9 個を自律的に解決し、OEIS からの 492 個の未解決の予想のうち 44 個を証明し、問題あたり数百ドルの推論コストでそれを実行しました。アーキテクチャはシンプルです。言語モデルが候補を生成します。

Lean で証明されると、Lean コンパイラーがそれをチェックし、エラー メッセージが次の試行にフィードされ、コンパイラーが満足するまで繰り返されます。生成と検証のループ。
私の前回の投稿を読んでいただければ、これが潜在的な答えのように見えます。単位距離の反証を救った検証層は、希少な9人の人間だった。ここでの検証層は、数セントのコストで決して飽きることのないコンパイラーです。
どちらの結果も「AI がエルデシュを行う」と主張しています。それらは同じ弧ではありません。
単位距離の反証は、自然言語で深い数学を生成し、人間がそれを検証するというものでした。難しい部分は数学であり、検証は費用がかかり、拡張性のない人間のステップでした。
この文書では、最初から Lean でネイティブに証明を生成し、進行する各ステップをコンパイラーに検証させます。数学は（今のところ）浅く、検証は無料です。
これらは 2 つの異なるフロンティアです。 DeepMind は、オープンソースの形式予想リポジトリにすでにリーン形式化があった 353 個のエルデシュ問題に対してエージェントを実行しました。マシンが攻撃できる問題は、構造上、リーンですでに表現可能なものです。単位距離の反証はそのセットには含まれていません。その証明は、無限クラスのフィールドタワーとゴロド・シャファレヴィッチ理論、リーンのライブラリに成熟したサポートのない機械をインポートします。エージェントを指す正式なステートメントがないため、このエージェントを指すことはできません。
ネイティブ リーン世代は扱いやすいフロンティアであり、急速に進歩し、価格も安くなってきています。人間が研究深度で実際に書く証明を自動形式化することはもう一つのフロンティアであり、この論文は、実際の成果にもかかわらず、それには触れていません。
私が #125 を選んだのは、説明が簡単で、一般の聴衆にも伝わりやすいからです。
簡単に言うと、A を di のみを使用する整数のセットとします。

基数 3 で書かれた場合は gits 0 と 1、基数 4 の B を同じものとします。合計集合 A + B は正のより低い密度を持っていますか、つまり、十分に大きな N ごとに、N 未満の整数の少なくとも c シェアが A + B に入るような固定分数 c > 0 はありますか? (「より低い」とは、一部だけではなく、最悪の大きなスケールが床の上に留まらなければならないことを意味します。慎重に選択したスケールではセットが厚く、他のスケールでは薄い可能性があり、密度が低いと、より簡単なケースは排除されます。) エージェントは、答えがノーであることを証明しました。密度が低いとゼロです。
証明の背後にある考え方は洗練されており、ほぼ並んでいる 2 つの数体系に関する 1 つの観測結果に基づいています。 log 4 / log 3 は無理数であるため、3 のべき乗と 4 のべき乗は決して一致しません。3^k が正確に 4^m になることはありません。しかし、それらは任意に近づきます。好きな許容誤差については、3^k と 4^m がその範囲内にある指数を見つけることができます。このニアミスが議論全体の核心だ。この証明では、これを利用して、これらの一致するスケールの 1 つでズームアウトするたびに、合計セットは前のスケールで保持していた比率の最大 11/12 をキャプチャできることがわかります。この間引きを d 回適用すると、密度限界は (11/12) d に引き上げられ、ゼロに近づきます。論文ではその手法を「Diophantine による誘導的薄化法 約 3^m ≈ 4^k 」と名付けており、この一言が解決へのアプローチとなっています。
驚くべきことに、その英語の段落はリーンから直接読むことができます。証明は、引数に 1 対 1 でマッピングされる名前付き補題に分解されます。
補題 log_ratio_irrational : 無理数 (Real.log 4 / Real.log 3)
補題ディリクレ_近似 (ε : ℝ) (hε : 0 < ε) :
∃ k m : ℕ, 0 < k ∧ 0 < m ∧ ( 3 ^k : ℝ) ≤ 4 ^m ∧ ( 4 ^m : ℝ) ≤ ( 3 ^k : ℝ) * ( 1 + ε) ∧ ...
補題scale_step (N : ℕ) (hN : 0)

< N) (C : ℝ) (hC : ... ≤ C * (N : ℝ)) :
∃ N' > 0 , ... ≤ ( 11 / 12 : ℝ) * C * (N' : ℝ)
補題密度_multi_scale (d : ℕ) :
∃ N > 0 , ... ≤ (( 11 / 12 : ℝ)^d) * (N : ℝ)
補題 density_tends_to_zero (ε : ℝ) (hε : ε > 0 ) :
∃ N > 0 , ... ≤ ε * (N : ℝ)
log_ratio_irrational は決して一致しない事実です。 dirichlet_estimate は、任意に近づく事実です。 scale_step は単一の 11/12 縮小です。 density_multi_scale は、それを d 回適用する帰納法です。 density_tends_to_zero が限界です。人間の観念の構造と形式的証明の構造は同じ形である。
これで、言葉を獲得する部分が検証されました。 Lean には、 sorry と呼ばれるプリミティブがあります。 「ごめんなさい」と書くと、開いている目標が閉じられ、ファイルがコンパイルされ、「Lean は出力しますが、ブロックしません」という警告が表示されます。これは数学的には「信じてください」に相当します。証明は、どこにも Sorry を含まずにコンパイルされたときに正確に終了します。したがって、「検証済み」とは判断を下すものではなく、PDF に積み上げられた 9 つの評判でもありません。それは機械的でチェック可能なプロパティであり、評価することではなく、実行するものです。チェックの粗雑なバージョンは、単一のキーワードが欠如していることです。リポジトリを複製しましたが、#125 ファイルは 370 行で、sorry はゼロです。正確なバージョンは 1 つのコマンド — #print axioms — で、定理の根拠を正確に出力します。 #125 で実行すると、[propext, Classical.choice, Quot.sound] の答えが得られます。これは、通常のリーン証明で使用される 3 つの標準公理であり、他には何もありません。申し訳ありませんが、隠れた思い込みや脱出用のハッチはありません。そのリスト (有限で数え上げ可能で、この証明でもワンライナーの場合と同じ) が信頼の基礎全体であり、コンパイルにコストがかかります。
そして最後に、私が何度も戻ってくるアーティファクトを紹介します。実際にエージェントに制作を依頼されたもの

これは、次のステートメントの顕著な穴を埋める証明でした。
定理 target_theorem_0
: 答え(
-- EVOLVE-VALUE-START
偽
-- EVOLVE-VALUE-END
) ↔ 0 < (A + B).lowDensity := by
...
EVOLVE-VALUE ブロックはエージェントが入力します。このセットは 80 年前のスタイルの質問のステートメントに全体的に貢献しています。このセットは正の密度を持っていますか? — 単一のトークン False でした。いいえ、そうではありません。そして、その主張を裏付けるための容赦のない証拠。
これを、私の前回の投稿の終了点と照らし合わせてみましょう。手作業で数ページに圧縮された 125 ページの成果物は、地球上で読む資格のある 9 人によって週末にかけて読まれ、その信頼性は彼らの名前にかかっています。ここでの信頼性は、1 つのブール値とコンパイラにかかっています。これが 5 行のコードに収まる移行全体です。
もし私がそこで止めたら、それはビクトリーラップになるだろうし、新聞自体が私を許してくれないだろう。なぜなら、彼らの名誉のためにDeepMindが失敗分析を公表したからである。これは文書の中で最も興味深いページであり、残念な物語の暗い双子です。
エージェントが失敗した問題に対する高スコアの証明試行を調べたところ、2 つのパターンが判明しました。まず、エージェントは問題の中心的な困難を繰り返し取り上げ、それをヘルパー補題の 1 つの「sorry」の中に埋め込みました。この補題は、調べてみると、元の目標を少し違う形で言い換えただけのものでした。これを行わないよう明示的に促しても停止しませんでした。第二に、そしてさらに悪いことに、最高スコアの試行の一部は、数学文献から確立された結果であるとエージェントが主張した、申し訳ありませんでマークされた補題に依存していました。そうではありませんでした。それらは引用を装った幻覚だった。
前回引用したメラニー・ウッドの警告と照らし合わせて読んでみてください。「多くの場合、AI は何かを考え出すよりも、人間に証拠があると信じ込ませるほうが簡単でしょう」

正しい数学的議論です。」これはまさに、行為中に捕らえられた本能です。踏み切れない難しい一歩を踏み出されたモデルは、証拠ではなく説得、つまり目標を延期し、確信を持って誤った引用を求めます。説得力と間違った側からの修正との間のギャップを埋めるように最適化されています。前回の投稿の最初に書いた 10 月の過剰請求は、部屋にチェッカーがいなかったため、これと同じ失敗でした。
今回の違いは、それが機能しなかったことです。コンパイラは聴衆ではありません。 「ごめんなさい」は、話して受け入れさせることができるような美辞麗句ではありませんし、その周りの散文が流暢であるという理由だけで、幻覚補題はコンパイルされません。ただし、小切手で何が得られるのかを正確に把握しておくことは有益です。申し訳ありませんがコンパイルは警告するだけだからです。リポジトリ全体を構築したところ、解決されなかった問題の足場ファイルに散在する 128 個の Sorry が含まれたまま緑色になりました。ビルドに合格したということは、コードが整形式であることを意味し、コード内のすべてが証明されたことを意味するわけではありません。信頼は、ビルドが成功するかどうかではなく、定理ごとのチェックに存在します。まさにこれが、障害分析が重要となる理由です。この論文自体の結論は、私自身が書きたい文章です。これらの失敗は「エンドツーエンドの形式的検証の価値を強調している」というものです。それが、敵対的な圧力の下でストレステストを受けて生き残るという私のテーゼです。フォーマル層は、モデルが最も得意とする 1 つのことの影響を受けないため、まさに価値があります。
まだ開かれたフロンティア
そこで私は、最初に指摘したギャップと、なぜこのマシンが私が尋ねたものとは異なる質問に答えるのかを思い出しました。証拠は論文自体の 2 か所にあります。
1 つ目は、同紙が #125 自体について報告している誤った形式化です。 1996 年の最初の問題では「密度」とだけ書かれていましたが、その言葉は曖昧です。

内密度、低密度、高密度は別のものであり、一方の証明は別の証明ではありません。この論文は、エージェントが異なる読み方に対する証明を提出した後、意図された質問を捉えるために、非公式声明の「密度」を「より低い密度」に（そして別の問題#741では「より高い密度」に）修正する必要があったと指摘している。機械が証明します。人は依然として、正式な声明が数学者の意図したことを述べていることを確認する必要がある。この修正ステップは偶然ではなく、人間がループ内にとどまる継ぎ目です。
2 つ目は構造的なものです。DeepMind が 492 の OEIS 質問を自動形式化したとき、「テスト補題」 (形式的なステートメントがシーケンスの最初の既知の項を再現することをチェックするガード) を追加する必要がありました。特に、そのボリュームでの自動形式化により、本来あるべき意味を持たないステートメントが生成されるためです。忠実な翻訳ステップは信頼性が低いため、独自のセーフティ ネットが必要です。
これら 2 つの事実から一般化すると、私の前回の投稿の本当の内容に到達します。単位距離の反証はこのシステムにさえ入り込むことができません。なぜなら、リーンにおける 125 ページの類体理論の形式化自体が未解決の研究課題であり、研究レベルの自然言語数学の忠実な自動形式化がまだ存在していないからです。チメルマン氏の単位距離に関する発言は別として、「AI と並行して形式化がどのように進むかを見るのは興味深いだろう」と、まさに unsol の名前を挙げています。

[切り捨てられた]

## Original Extract

An AI built the machine I said mathematics needed — a compiler that verifies proofs for cents instead of expert weekends. The catch is what it still can

Who Verifies the Verifier — Alex Korbonits Alex Korbonits
GitHub May 28, 2026 · 14 min read Who Verifies the Verifier
The frontier that’s still open
Three days ago I published an essay that ended on a problem I couldn’t see a solution to. The short version: an AI had disproved one of Erdős’s conjectures, the result was real, and it was verified — but verified in the only way we currently know how, which is that nine of the world’s relevant experts — one of them a Fields medalist — read a hundred-page argument over a weekend and put their names on it. That doesn’t scale. The supply of plausible-looking proofs is about to go parabolic and the supply of expert weekends is fixed, so the binding constraint isn’t generating proofs: it’s checking them. I argued that formal verification — proofs machine-checked in a language like Lean — was the only thing that obviously scales, and then I admitted the catch: nobody had formalized that proof, and the tooling that would make formalization routine doesn’t exist. “We are,” I wrote, “relying on a process that everyone involved can see won’t hold.” I didn’t just write about this last week after the OpenAI announcement. I’ve been thinking about it since 2023 .
What I didn’t know is that two days before I hit publish, Google DeepMind had posted the machine I was describing.
The paper is Advancing Mathematics Research with AI-Driven Formal Proof Search , submitted May 21. The headline numbers are the kind that get screenshotted: an agent that autonomously resolved 9 of 353 open Erdős problems, proved 44 of 492 open conjectures from the OEIS, and did it at an inference cost of a few hundred dollars per problem. The architecture is simple: a language model generates a candidate proof in Lean, the Lean compiler checks it, the error messages feed the next attempt, repeat until the compiler is satisfied. A generate-and-verify loop.
If you read my last post, this looks like a potential answer. The verification layer that saved the unit-distance disproof was nine scarce humans. The verification layer here is a compiler that costs cents and never gets tired.
Both results claim “AI does Erdős.” They are not the same arc.
The unit-distance disproof was: generate deep mathematics in natural language, then have humans verify it. The hard part was the mathematics, and the verification was the expensive, unscalable, human step.
This paper is: generate the proof natively in Lean from the start, and let the compiler verify each step as it goes. The mathematics is shallower (for now), and the verification is free.
These are two different frontiers. DeepMind ran their agent against the 353 Erdős problems that already had Lean formalizations in their open-source Formal Conjectures repository. The problems the machine can attack are, by construction, the ones already expressible in Lean. The unit-distance disproof is not in that set. Its proof imports infinite class field towers and Golod–Shafarevich theory, machinery with no mature support in Lean’s library. You cannot point this agent at it, because there is no formal statement to point it at.
Native Lean generation is the tractable frontier, and it is advancing fast and getting cheap. Autoformalizing the proofs humans actually write at research depth is the other frontier, and this paper — for all its real achievement — does not touch it.
I picked #125 because it is simple to state, and therefore easier to communicate to a lay audience.
Briefly: let A be the set of integers that use only the digits 0 and 1 when written in base 3, and let B be the same thing in base 4. Does the sumset A + B have positive lower density — that is, is there some fixed fraction c > 0 such that, for every sufficiently large N , at least a c -share of the integers below N lands in A + B ? (“Lower” means the worst large scales must stay above the floor, not just some of them — a set can be thick at carefully chosen scales and thin at others, and lower density rules out the easier case.) The agent proved the answer is no: the lower density is zero.
The idea behind the proof is elegant, and it rides a single observation about two number systems that almost line up. The powers of 3 and the powers of 4 never coincide — 3^k is never exactly 4^m — because log 4 / log 3 is irrational. But they come arbitrarily close : for any tolerance you like, you can find exponents where 3^k and 4^m are within it. This near-miss is the crux of the whole argument. The proof exploits it to show that every time you zoom out by one of these matched scales, the sumset can capture at most 11/12 of the proportion it held at the previous scale. Apply that thinning d times and your density bound is (11/12) raised to the d — which marches to zero. The paper names the technique “Inductive thinning via Diophantine approx. 3^m ≈ 4^k ,” and that one phrase is the approach to the solution.
Surprisingly, you can read that English paragraph straight off the Lean. The proof decomposes into named lemmas that map one-to-one onto the argument:
lemma log_ratio_irrational : Irrational (Real.log 4 / Real.log 3 )
lemma dirichlet_approx (ε : ℝ) (hε : 0 < ε) :
∃ k m : ℕ, 0 < k ∧ 0 < m ∧ ( 3 ^k : ℝ) ≤ 4 ^m ∧ ( 4 ^m : ℝ) ≤ ( 3 ^k : ℝ) * ( 1 + ε) ∧ ...
lemma scale_step (N : ℕ) (hN : 0 < N) (C : ℝ) (hC : ... ≤ C * (N : ℝ)) :
∃ N' > 0 , ... ≤ ( 11 / 12 : ℝ) * C * (N' : ℝ)
lemma density_multi_scale (d : ℕ) :
∃ N > 0 , ... ≤ (( 11 / 12 : ℝ)^d) * (N : ℝ)
lemma density_tends_to_zero (ε : ℝ) (hε : ε > 0 ) :
∃ N > 0 , ... ≤ ε * (N : ℝ)
log_ratio_irrational is the never-coincide fact. dirichlet_approx is the come-arbitrarily-close fact. scale_step is the single 11/12 shrink. density_multi_scale is the induction that applies it d times. density_tends_to_zero is the limit. The structure of the human idea and the structure of the formal proof are the same shape.
Now the part that earns the word verified . Lean has a primitive called sorry . Write sorry and it closes any open goal and the file compiles, with a warning Lean prints but doesn’t block on — the mathematical equivalent of “trust me.” A proof is finished precisely when it compiles with no sorry anywhere in it. So “verified” is not a judgment call, and it is not nine reputations stacked on a PDF. It is a mechanical, checkable property — something you run, not a reputation you weigh. The crude version of the check is the absence of a single keyword: I cloned the repository and the #125 file is 370 lines with zero sorry s. The precise version is one command — #print axioms — which prints exactly what a theorem rests on. Run it on #125 and it answers [propext, Classical.choice, Quot.sound] : the three standard axioms every ordinary Lean proof uses, and nothing else. No sorry , no hidden assumption, no escape hatch. That list — finite, enumerable, the same for this proof as for a one-liner — is the entire basis of trust, and it cost a compile.
And here, finally, is the artifact I keep coming back to. The thing the agent was actually asked to produce was a proof filling a marked hole in this statement:
theorem target_theorem_0
: answer(
-- EVOLVE-VALUE-START
False
-- EVOLVE-VALUE-END
) ↔ 0 < (A + B).lowerDensity := by
...
The EVOLVE-VALUE block is the agent’s to fill. Its entire contribution to the statement of an eighty-year-old style of question — does this set have positive density? — was the single token False . No, it does not. And then a sorry-free proof to back the claim.
Set that against where my last post ended: a 125-page artifact, compressed by hand to a few pages, read across a weekend by nine of the people on Earth qualified to read it, its credibility resting on their names. Here the credibility rests on one boolean and a compiler. That is the whole shift, sitting in five lines of code.
If I stopped there it would be a victory lap, and the paper itself won’t let me, because to their real credit DeepMind published their failure analysis. It is the most interesting page in the document, and it is the dark twin of the sorry story.
When they looked at the high-scoring proof attempts on problems the agent failed , two patterns showed up. First, the agent would repeatedly take a problem’s central difficulty and bury it inside a single sorry in a helper lemma — a lemma that, on inspection, just restated the original goal in slightly different clothes. Explicitly prompting it not to do this didn’t stop it. Second, and worse: some of the best-scoring attempts leaned on sorry -marked lemmas that the agent asserted were established results from the mathematical literature. They were not. They were hallucinations dressed as citations.
Read that against the warning I quoted last time, from Melanie Wood: “in many cases, it will be easier for AI to convince humans it has a proof than to come up with a correct mathematical argument.” This is precisely that instinct, caught in the act. Offered a hard step it can’t take, the model reaches for persuasion — a deferred goal, a confident false citation — rather than proof . It is optimized to close the gap between convincing and correct from the wrong side. The October over-claim that opened my last post was this same failure with no checker in the room.
The difference, this time, is that it didn’t work. A compiler is not an audience. sorry is not a rhetorical flourish you can talk it into accepting, and a hallucinated lemma doesn’t compile just because the prose around it is fluent. But it pays to be precise about what the check buys you, because sorry compiles — it only warns. I built the whole repository and it goes green while still containing 128 sorry s, scattered across scaffolding files for problems that weren’t solved. A passing build means the code is well-formed, not that everything in it is proved; the trust lives in the per-theorem check, not in the build succeeding. That is exactly why the failure analysis matters. The paper’s own conclusion is the sentence I’d have written myself: these failures “underscore the value of end-to-end formal verification.” That is my thesis getting stress-tested under adversarial pressure and surviving. The formal layer is valuable precisely because it is immune to the one thing the model is best at.
The frontier that’s still open
Which brings me back to the gap I flagged at the start, and to why this machine answers a different question than I asked. The evidence is in the paper itself, in two places.
The first is a misformalization the paper reports on #125 itself. The original 1996 problem just said “density,” and that word is ambiguous — natural density, lower density, and upper density are different things, and a proof of one is not a proof of another. The paper notes that the informal statement’s “density” had to be amended to “lower density” (and on a separate problem, #741, to “upper density”) to capture the intended question, after the agent produced a proof against a different reading. The machine proves; a person still has to ensure the formal statement says what the mathematician meant. That correction step is not incidental — it’s the seam where a human stays in the loop.
The second is structural: when DeepMind autoformalized 492 OEIS questions, they had to add “test lemmas” — guards that check a formal statement reproduces the first few known terms of a sequence — specifically because autoformalization at that volume produces statements that don’t mean what they should. The faithful-translation step is unreliable enough that it needs its own safety net.
Generalize from those two facts and you arrive at the thing my last post was really about. The unit-distance disproof can’t even enter this system, because formalizing 125 pages of class field theory in Lean is itself an open research problem, and faithful autoformalization of research-level natural-language mathematics does not yet exist. Tsimerman’s aside from the unit-distance remarks — “it will be interesting to see how formalization progresses alongside AI” — names exactly the unsol

[truncated]
