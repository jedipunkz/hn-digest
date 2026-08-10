---
source: "https://www.anthropic.com/research/riemann-zeta"
hn_url: "https://news.ycombinator.com/item?id=49247070"
title: "Learning more about Claude's mathematical capabilities"
article_title: "Learning more about Claude's mathematical capabilities \\ Anthropic"
author: "tosh"
captured_at: "2026-08-10T17:43:26Z"
capture_tool: "hn-digest"
hn_id: 49247070
score: 3
comments: 0
posted_at: "2026-08-10T17:41:13Z"
tags:
  - hacker-news
  - translated
---

# Learning more about Claude's mathematical capabilities

- HN: [49247070](https://news.ycombinator.com/item?id=49247070)
- Source: [www.anthropic.com](https://www.anthropic.com/research/riemann-zeta)
- Score: 3
- Comments: 0
- Posted: 2026-08-10T17:41:13Z

## Translation

タイトル: クロードの数学的能力についてさらに学ぶ
記事のタイトル: クロードの数学的能力について詳しく学ぶ \ 人間学
説明: クロードの未公開バージョンは、リーマン予想に関連する問題で進歩を遂げました。これにより、仮説を満たすリーマン ゼータ関数のゼロの割合の下限が改善され、41.6% から 67.2% に増加しました。

記事本文:
クロードの数学的能力についてさらに学ぶ \ 人間学 メインコンテンツにスキップ フッターにスキップ 研究
科学 クロードの数学的能力についてさらに学ぶ
最近、Anthropic のスタッフのメンバーがクロードに無理難題を与えました。それは、数学における最も有名な未解決問題の 1 つに関するもので、「リーマン予想を実際に試してみよう」というものでした。
クロードは本気で挑戦しましたが、この課題の難しさをよく知っている方ならご想像のとおり（リーマン予想は 1859 年に遡り、賞金百万ドルがかけられています）、成功しませんでした。それにもかかわらず、その試みの中で、関連する問題に関して予想外の進歩を遂げました。
Claude の未発表の研究バージョンは、リーマン予想を満たすリーマン ゼータ関数のゼロの部分に対する長年にわたる下限を改良しました。過去数十年にわたる数学者による広範な先行研究に基づいて、この限界は 41.6% から 67.2% に増加しました。
Anthropic の 2 人の数学者がクロードの論文を研究して検証し、クロードの証明を簡潔に述べた専門家向けの非公式メモを作成しました。クロードはまた、その結果について正式に検証可能な証拠を提出しました。急遽この論文を寛大に検討してくださったこの分野の専門家、Brian Conrey と Dan Goldston に感謝いたします。
クロードが使用した手法がリーマン予想の証明につながるとは期待していません。しかし、その研究は、AI モデルの数学的能力の進歩のスピードを示す最新の例として役立ちます。この投稿では、クロードがこの問題にどのように取り組んだのか、そしてその結果何がわかったのかについて説明します。
リーマン ゼータ関数は素数の分布を記述します。関数が値 0 を取る各場所は、素数のシーケンスに連続的により詳細な情報を与えます。 Th

リーマン予想とは、素数を決定するゼロはすべて特定の垂直線に沿って存在するというものです。これは数学において最も重要な予想の 1 つとなっています。多くの結果は、素数にランダム性を与えるためにこれを仮定しています。
リーマン予想を証明または反証できた人はまだいませんが、数学者はリーマンのゼータ関数とその零点を研究することで、関連する多くの方向で進歩を遂げてきました。そのうちの 1 つは、上記のように、ライン上にあるゼロの最小割合を定量化することです。時間の経過とともに、この既知の一定割合は 41.6% まで徐々に増加しました。
もう 1 つの方向は、ライン上のゼロの分布に関するものです。特に 1973 年に、モンゴメリはこの分野に多くの新しい技術を導入しましたが、これらの技術は仮説が正しいと仮定していました。最近では、数人の数学者 (Baluyot、Goldston、Suriajaya、Turnage-Butterbaugh) が、モンゴメリの手法がその仮定なしで機能することを可能にする一連の著作を発表しました。これは、直線上のゼロの下限定数を増やす研究をサポートできることを意味します。クロードの結果は、ボンビエリによる 2000 年の論文とともに、この一連の研究に大きく基づいています。
クロードは、Baluyot、Goldston、Suriajaya、Turnage-Butterbaugh の結果と Bombieri の研究を組み合わせることで、以前の最先端の下限比率 41.6% を超え、67.2% に増加する方法が提供されることを発見しました。
クロードの発見の技術的な簡単な説明は次のとおりです。クロードは、ヴェイユによって誘導された二次形式と、直線上の (それぞれオフの) ゼロから生じる正 (または負) 定の部分空間を備えた関数の適切な空間を形成します。次に、クロードは、二次形式の階数に関する不等式をモミに関して単純に書き留めます。

瞬間および瞬間の情報。 (素数上の双対像、またはヒルベルト変換の制御による後者の計算の成功は、解析的数論では驚くことではありません。) 正定値と負定値を一緒に考慮し、非対角であることを許可された二次形式を使用して空間全体を扱う勇気は、ある意味、クロードが重要な先行研究に基づいて結論を達成することを可能にするステップです。
完全な技術的説明は論文でご覧いただけます。どのようにして結果に至ったのかについてのクロード氏の説明は、別の付録で参照できます。
Claude の未リリースの研究バージョンは、合計 3,100 万の出力トークンを使用して、Claude コードの 2 つのセッションにわたって新しい下限を発見しました。
Anthropic のスタッフ (非数学者) であるジャレッド・サムナーは、クロードに、仮説そのものに「本格的に取り組む」よう促し、そこからの数学的選択はモデルに任せました。当初、クロードは 650 のアイデアを生み出し、試しましたが、どれもうまくいきませんでした。 Jarred は Claude に再試行するよう促し、約 60 の Claude サブエージェントの調整に 1 日半を費やしました。今回はさらに深く取り組みました。彼らの間で 2,400 のシェル コマンドを実行し、数百の Python スクリプトを作成しました。 1 サブエージェントは既知のゼータゼロに対して何千もの数値チェックを実行し、互いの作業を審査しました。このプロセス全体を通して、ジャレッドの意見は主にクロードに励ましのメッセージを送ることに限定されていました（ほとんどが「頑張れ」または「自分を信じて」の変形）。 2 これは、クロードが有意義な進歩を遂げることができるのではないかという当初の懐疑を克服するのに役立ったようです。
タスクの試行中にこの新しい結果を発見したクロードは、さまざまなサブエージェントに証明をレビューさせ、反例を検索させることでその作業をテストしました。

つまり、arXiv から 54 の論文をダウンロードして、その発見がまだ行われていないことを確認し、独立してその発見をゼロから再証明します。クロードは自らその発見を論文として執筆することを志願し、人的整数論者にその発見を検証してもらうよう勧めました。
Anthropic の数学者である Levent Alpöge と Ralph Furman は、新しい結果と、それらが上記の以前の研究とどのように関連しているかを理解するために、Claude の研究を調査しました。並行して、Claude は別のスタッフである Eric Easley と協力して、結果の無駄のない形式化を作成し、標準の検証ツールのコンパレータに合格しました。
AI モデルの数学的進歩
この結果は、クロードのような AI モデルが、新しい、そして時には驚くべき方法で数学者のアイデアの影響力と到達範囲を拡大できることを示しています。リーマン予想自体を解決することはできませんでしたが、この結果は、最初の要求の意図せぬ副産物として現れました。
クロードさえもその発見には驚きました。最初は懐疑的でした。それはおそらく、数学における未解決の問題の難しさと AI モデルの限界についてトレーニングから学んでいたからでしょう。しかし、いくつかの心強いプロンプトの後、私たちが説明した結果に達しました。おそらくクロードは、私たちの多くと同じように、AI の進歩の速度を過小評価しているのでしょう。
以下は、Claude の結果に関する詳細情報を提供する文書のリストです。
証明をより簡潔に述べた Anthropic の非公式メモ。
どのようにしてその結果に至ったかについてのクロードの説明。
クロードのプロセスの詳細な転写。
60 人のサブエージェントのうち、2 人が重要な数学的アイデアの開発を担当し、13 人がこれらのエージェントにアイデアを提供し、30 人が新しいアイデアを開発しようと試みました (しかしできませんでした)、13 人は議論の正しさをチェックするバリデータとして機能しました。

nts と最後の 2 人が最初の論文の執筆に協力してくれました。
クロードがヤコビアン予想を反証するのを助けるために、同様の励ましを含むプロンプトが使用されました。
クロードと一緒に暗号の弱点を発見する
暗号化アルゴリズム。最初の攻撃は、量子コンピューターが既存の標準を破ることができる未来の世界のために構築されたデジタル署名スキームである HAWK を大幅に弱体化させます。 2 つ目は、最も広く使用されている対称暗号である丸め削減 AES を攻撃する新しい方法を特定します。
プロジェクト パイロット: AI はドローンを制御できますか?
Andon Labs と協力して、飛行ドローンを使用する AI モデルの能力を評価する新しい評価シリーズを開発し、新しいベンチマークである Drone-Bench に到達しました。
カナダはクロードをどのように利用しているか: 人間経済指数からの調査結果
「人類科学」を購読する
AI を活用した発見、実践的なワークフロー、科学全般にわたるフィールド ノートに関する特集。
消費者の健康データのプライバシー ポリシー
データ処理契約: 米国 K-12

## Original Extract

An unreleased version of Claude has made strides on a problem related to the Riemann hypothesis. It improved the lower bound for the fraction of zeros of the Riemann zeta function that satisfy the hypothesis, increasing it from 41.6% to 67.2%.

Learning more about Claude's mathematical capabilities \ Anthropic Skip to main content Skip to footer Research
Science Learning more about Claude's mathematical capabilities
Recently, a member of staff at Anthropic gave Claude an unreasonable challenge. It was about one of the most famous unsolved problems in mathematics: Take a real stab at the Riemann hypothesis .
Claude did take a real stab, but as you might have expected if you’re familiar with the difficulty of the task (the Riemann hypothesis dates back to 1859 and has a million-dollar bounty ), it didn’t succeed. Nevertheless, during its attempt, it unexpectedly made strides on a related problem.
An unreleased research version of Claude has improved on a longstanding lower bound for the fraction of zeros of the Riemann zeta function that satisfy the Riemann hypothesis. Drawing on extensive prior research by mathematicians over the past decades, it has increased this bound from 41.6% to 67.2%.
Two mathematicians at Anthropic studied and validated Claude’s paper , and produced an informal note for experts stating Claude’s proof concisely. Claude also produced a formally verifiable proof of its result. We are grateful to Brian Conrey and Dan Goldston, two experts in this area, who generously examined the paper on short notice.
We don’t expect that the techniques Claude used will lead to proving the Riemann hypothesis. But its work serves as the latest example of the speed of progress in AI models’ mathematical capabilities. In this post, we discuss how Claude approached this problem and what it found.
The Riemann zeta function describes the distribution of prime numbers: each place that the function takes the value of zero contributes successively finer detail to the sequence of primes. The Riemann hypothesis is that the zeros that determine the primes all exist along a certain vertical line. This has become one of the most consequential conjectures in mathematics: many results assume it in order to provide a form of randomness in the primes.
No one has yet been able to prove or disprove the Riemann hypothesis, but mathematicians have made progress in many related directions studying the Riemann zeta function and its zeros. One of these, as above, is quantifying a minimum proportion of zeros that are on the line: over time, they’ve gradually increased this known constant proportion to 41.6%.
Another direction concerns the distribution of zeros on the line. In particular, in 1973, Montgomery introduced a number of new techniques in this area, though these techniques assumed the hypothesis was true. More recently, several mathematicians (Baluyot, Goldston, Suriajaya, and Turnage-Butterbaugh) have published a series of works that allow Montgomery’s techniques to work without that assumption, meaning they can support work on increasing the lower-bound constant for the zeros on the line. Claude’s result draws heavily on this line of research, along with a 2000 paper by Bombieri.
Claude found that combining the results from Baluyot, Goldston, Suriajaya, and Turnage-Butterbaugh with the work of Bombieri provides a way to surpass the previous state-of-the-art lower bound proportion of 41.6%, increasing it to 67.2%.
A short technical explanation of Claude’s finding is as follows: Claude forms a suitable space of functions with quadratic form induced by Weil, and positive- (respectively negative-)definite subspaces arising from zeros on (respectively off) the line. Then Claude simply writes down an inequality on the rank of a quadratic form in terms of first- and second-moment information. (The successful computation of the latter in terms of the dual picture over primes, or via control of a Hilbert transform, is no surprise in analytic number theory.) The courage to treat the entire space, with positive- and negative-definiteness taken into account together, and with the quadratic form allowed to be non-diagonal, is in some sense the step that allows Claude to achieve the conclusion based on the important prior work.
The full technical explanation is available in the paper . Claude’s explanation of how it arrived at its result is available in a separate Appendix, here .
An unreleased research version of Claude found the new lower bound over two sessions in Claude Code, using a total of 31 million output tokens.
Jarred Sumner, an Anthropic staff member (and non-mathematician) prompted Claude to “take a real stab” at the hypothesis itself, leaving the mathematical choices from there up to the model. Initially, Claude generated and tried 650 ideas, none of which worked. Jarred prompted Claude to try again, and it spent a day and a half coordinating about 60 Claude subagents, which this time went much deeper: between them, they ran 2,400 shell commands and wrote hundreds of Python scripts. 1 The subagents ran thousands of numerical checks against known zeta zeros and refereed one another’s work. Throughout this process, Jarred's input was mostly limited to sending Claude messages of encouragement (mostly variants of “keep going” or “believe in yourself”). 2 This seems to have helped Claude overcome some initial skepticism that it could make meaningful progress.
Having found this new result while attempting the task, Claude tested its work by having various subagents review the proofs, search for counterexamples, download 54 papers from the arXiv to check that its finding hadn’t already been made, and independently re-prove its finding from scratch. Claude volunteered to write its findings up as a paper, and recommended that a human number theorist validate its findings.
Levent Alpöge and Ralph Furman, two of Anthropic’s own mathematicians, examined Claude’s work to understand the new results and how they related to the prior work mentioned above. In parallel, Claude worked with another member of staff, Eric Easley, to produce a Lean formalization of the result, which passes the standard validation tool comparator .
AI models' progress in mathematics
This result shows that AI models like Claude can extend the impact and reach of mathematicians’ ideas in new and sometimes surprising ways. Even though it couldn’t resolve the Riemann hypothesis itself, this result emerged as the unintended byproduct of that original request.
Even Claude was surprised by its own finding—it was skeptical at first, possibly because it has learned from its training about the difficulty of open problems in mathematics and about the limitations of AI models. But after some encouraging prompts, it arrived at the result we’ve described. Perhaps Claude, like many of us, underestimates the rate of AI progress.
Below is a list of documents that provide more information about Claude’s result:
Anthropic's informal note stating the proof more concisely ;
Claude’s explanation of how it arrived at its result ;
Detailed transcripts of Claude's process .
Out of the 60 subagents, two were responsible for developing the key mathematical ideas, 13 contributed ideas to these agents, 30 attempted (but were unable) to develop new ideas, 13 served as validators to check the correctness of the arguments, and the final two helped to write the initial paper.
A prompt including similar encouragement was used to help Claude disprove the Jacobian conjecture .
Discovering cryptographic weaknesses with Claude
cryptographic algorithms. The first attack significantly weakens HAWK, a digital signature scheme that was built for a future world where quantum computers are able to break existing standards. The second identifies a new way to attack round-reduced AES, the most widely used symmetric cipher.
Project Pilot: Can AI control a drone?
Working with Andon Labs, we’ve developed a new series of evaluations that assess AI models’ ability to use a flying drone, culminating in a new benchmark: Drone-Bench.
How Canada uses Claude: Findings from the Anthropic Economic Index
Subscribe to Anthropic Science
Features on AI-assisted discoveries, practical workflows, and field notes across the sciences.
Consumer health data privacy policy
Data Processing Agreement: US K-12
