---
source: "https://m-malinowski.github.io/2026/06/05/ai-open-physics-problems.html"
hn_url: "https://news.ycombinator.com/item?id=48447543"
title: "Can AI answer open questions in physics?"
article_title: "Can AI answer open questions in physics? | Reading the quantum"
author: "stared"
captured_at: "2026-06-08T18:58:58Z"
capture_tool: "hn-digest"
hn_id: 48447543
score: 3
comments: 0
posted_at: "2026-06-08T16:33:35Z"
tags:
  - hacker-news
  - translated
---

# Can AI answer open questions in physics?

- HN: [48447543](https://news.ycombinator.com/item?id=48447543)
- Source: [m-malinowski.github.io](https://m-malinowski.github.io/2026/06/05/ai-open-physics-problems.html)
- Score: 3
- Comments: 0
- Posted: 2026-06-08T16:33:35Z

## Translation

タイトル: AI は物理学における未解決の疑問に答えることができるか?
記事のタイトル: AI は物理学の未解決の質問に答えることができますか? |量子の読み取り
説明: 新しい AI モデルが登場するたびに、誰かが念のためにコラッツ予想を投げかけるというジョークです。実際、これはおそらく最悪のアイデアではありませんが、「高校の物理学は習得できるが、まだすべての理論を作成することはできません」は、物理学のレベルを説明するのに最も役立つものではありません。
[切り捨てられた]

記事本文:
-->
-->
量子の読み取り
イオンのアイデアについて
AI は物理学における未解決の疑問に答えることができるでしょうか?
新しい AI モデルが登場するたびに、誰かが念のためにコラッツ予想を投げかけるというのがジョークです。実際、これはおそらく最悪のアイデアではありませんが、「高校の物理学は習得できるが、まだすべての理論を作成することはできません」は、AI モデルの物理学習熟度のレベルを表す最も有用な説明ではありません。私が知る必要があるのは、2026 年半ばの時点で、フロンティア AI はフロンティア物理学において正確にどれくらい優れているのかということです。
出版された文献によると、その答えは「そこに到達しつつある」ということです。たとえば、CritPt は、研究レベルの物理学の 50 の質問のベンチマークをまとめました。 2025 年 11 月の評価時点では、最上位モデルのスコアはわずか 12.6% でした。しかし、それから 9 か月間で、最高スコアは 30% 以上に上昇し、合計 10 モデルが 11 月の優勝者よりも高いスコアを獲得しました。
同様に、理論物理学問題のベンチマークである TPBench について考えてみましょう。 2025 年 2 月に発行されたプレプリントでは、モデルは高校の物理学をほぼ完璧に学習しましたが、大学院レベルの物理学では最大 50%、研究レベルの物理学では最大 15% のスコアを獲得しました。
2026 年に同じベンチマークを再実行したところ、GPT-5.5 (xhigh) のスコアは大学院レベルの質問で約 85%、研究レベルの質問ではなんと 55% であることがわかりました。これはかなり驚くべきことであり、特に GPT-5.5 (xhigh) は世の中で最も強力な物理モデルですらないことです。
それでも、これらのベンチマークの質問は一般的に人為的なものです。それらはすべて具体的で解決可能であり、解決策を誰かが知っています。また、著者が答えを秘密にしていたとしても、その答えが何らかの形でインターネット上に「漏洩」する可能性があるのではないかと心配する人もいます。出版された研究論文の一部として。
そこで、数週間前、私は次のように決めました。

o 最先端の AI モデルを「実際の」未解決の物理学の質問に投げかける実験を行います。これが私が見つけたものです。
物理学における未解決の疑問は何ですか?
まず明確にしておきたいのは、「物理学のオープンな質問」というのは、「物質と反物質の非対称性の起源は何ですか」といった、ナショナル ジオグラフィックで聞いたような質問のことではありません。
私たちは皆、物理学における「大きな」未解決の疑問について聞き慣れていますが、実際のところ、この分野には「小さな」未解決の疑問がたくさんあります。実際、世界中の物理学者は毎日何百もの論文やプレプリントを発表しており、その多くには「未解決の疑問」が含まれています。これには、説明を逃れる不可解な実験観察から、合理的であるように見えるが証明を回避する補題、さらにはフォローアップ作業の長期的な提案に至るまで、あらゆるものが含まれます。
これらのオープンクエスチョンの良いところは、本当にオープンな質問であるということです。 TPBench とは異なり、これらの質問のいずれかに答えることは、たとえほんのわずかで、非常にニッチなトピックであっても、人類の知識を真に進歩させることになります。さらに、既存のベンチマークとは異なり、問題の難易度、あるいは解決策が存在するかどうかさえもアプリオリにわかりません。
同時に、それらの多くは実際にはそれほど難しくない、と私は仮説を立てています。物理学の論文における典型的な「未解決の質問」は、非常に少数の人間から比較的少量の注目を集めただけです。このため、彼らは AI 入力の有望なターゲットになります。
実験では、quant-ph arXiv から 1 週間分のプレプリントをスクレイピングすることから始め、「未解決」、「未解決」、「将来の作業」などのフレーズのグレッピングに基づいて、未解決の質問の 200 を超える候補を含む 275 の論文を見つけました。
最初の観察は、当然のことながら、私の「未解決の問題」の大部分は、(私の意見では)

イオン) はまだ AI には適していません。典型的な理由は次のとおりです。
実験/ハードウェア検証に関するものです。私のデータセットの例: 「数ケルビン温度での動作はまだ実証されていません。」
それらは将来のロードマップの漠然とした概要です。例: 「今後の研究では、ボルン・マルコフ近似を超える強結合補正、非ガウス TLS ノイズ、機械学習ベースのリアルタイム推定を検討します」、「現実的な不確実性の下での秘密量子通信に関する将来の研究 (適応的、相関的、展開指向、分布的に堅牢な拡張を含む)」
とても幅広いです。例: 「孤立した量子多体系の非平衡力学の記述は、基本的な未解決の問題です。」
彼らは実践的なエンジニアリングに重点を置いています。例: 量子シミュレーションの論文では、「完全なゲート数解析、回路合成 […]、および量子ハードウェアの実行」は将来の作業に残されています。別の論文では、「完全なエンドツーエンドのフォールトトレラントな分散量子コンピューティング コンパイラ」が求められています。
大きすぎます。例: 高次元レイヤー コードに関する論文では、「これらの 4D および 5D レイヤー コードの有限温度自己修正特性の決定」について尋ねています。
それらは広範なプロトコル設計に関するものです。例: QKD 論文は、「このような非ガウス測定 CV-QKD プロトコルの無条件安全性の証明」について尋ねています。分散分子ノード量子コンピューティングに関する論文では、「実現可能なモジュール間もつれ率とハードウェア リソース要件でフォールト トレランスを達成するモジュール式アーキテクチャともつれ生成プロトコルを設計する」ことが求められています。
私は、AI アシストに最適な質問を特定することを目的として、軽量かつ積極的な (エージェント的な) 下位選択パイプラインでこれに対処しました。

アンス。これには個人的な判断も含まれていますが、全体として、次のような質問にはボーナス ポイントが与えられました。
はい/いいえ数学的主張。例: 「2 つの純粋なステアリング状態を許容するすべてのもつれ 2 量子ビット状態は、密度行列ランク 2 を持たなければなりません。」
数式や数値に関する質問。例: 論文では閾値の公式を数値的に確立し、分析的な証明を求めています。
最適化/特性評価に関する質問。例: ある論文では、「PPT のもつれ状態に対する再調整基準の違反の可能性が最大限にあることを分析的に特徴付ける」ことが求められました。
結果拡張の質問。例: 論文は単一量子ビット量子チャネルに関する結果を証明し、複数量子ビットチャネルへの拡張を求めました。
この下方選択の後、33 論文にわたる 48 の質問が抽出され、独立したユニットにパッケージ化され、AI が 1 つずつ解読を試みるようになりました。
私の最初の試みでは、自動化された 2 つのエージェント ループを使用しました。claude opus 4.8 を「ソルバー」として、gpt 5.5-xhigh を「レビューアー」として使用しました。しかし、これは非常に低い成功率をもたらしました。次に、「ソルバー」として gpt 5.5 Pro に切り替えたところ、成功率が大幅に向上しました。ただし、API 経由ではそのモデルにアクセスできず、Web UI 経由のみであったため、実験の残りの部分はかなり手動で、体系的ではありませんでした。
GPT 5.5 では、33 の論文にわたって 48 の問題に取り組みました。
30 件の論文にわたって、そのうち 42/48 件を正常に解決したと主張しました。
私はこれらの論文の著者に 30 通の電子メールを送り、結果の検証を求めました
これらのメールのうち 16/30 件に返信を受け取りました
検証の結果は次のようになりました。
25% (4/16): 悪い。 AI が誤って証明を主張したか、問題の枠組みを誤った
55% (9/16): 混合。 AI は部分的または有望な証拠を提供しましたが、おそらく不完全です
20% (3/16): 良い。 AI確認

問題を効果的に解決できたと考えられます。
これらは見出しの数字ですが、本当に興味深いニュアンスが隠されています。それでは、次回はそれぞれの例、物理学について何を教えてくれるのか、そして実験のパフォーマンスを向上させるために実験を改善する方法を詳しく見ていきましょう。
失敗の例の 1 つは、射影測定を超えた量子ランダム性による問題でした。引用: 「単純な疑問は未解決のままです: 与えられた極値測定によってどの程度の固有のランダム性が生成されるのでしょうか?」 。 AIは、純粋状態のミニマックス最適化として再定式化し、変分特性化を行うことによってそれを解決したと主張しました。 2 つの問題:
この論文は、わずかに異なる表記法を使用して、これらすべてをすでに行っています。したがって、本質的に、AI が行ったことは、論文内の結果の一部を再定式化または再導出したことだけです。
人間の読者であれば、実際の未解決の問題は、最適化問題を定式化して特徴付けるだけではなく、最適化問題に対する明示的な解決策であることを理解したでしょう。
もう 1 つの誤解されている未解決の質問は、「密度演算子の限界問題」から来ました。序論で、著者らは「Leifer & Poulin (2008) は、エントロピー量子の条件付き独立関係が完全なグラフイド特性を持つかどうかは未解決のままである」という歴史的な発言をしました。 AI はこの質問を「量子の条件付き独立関係は任意の密度演算子に対して完全なグラフイド特性を持っているか」と解釈し、すぐに反例を見つけることができました。
ただし、前の文は「ここで考慮される厳密に正の有限次元設定では…」で始まりました。これにより、実際の問題は、この性質が厳密に正の密度演算子に当てはまるかどうかであることが AI に認識されるはずです。これも、AI が Leifer & P の実際の引用文を読んでいたら明らかだったでしょう。

オウリン（2008）。
さまざまな種類の部分ソリューションが含まれる最大のカテゴリ。
著者の最も一般的な反対意見は、AI は証明を主張したが、数学を検証するのに十分な詳細を提供しなかったというものです。それは主に私の悪かった点でした。それ以来、私は常にモデルに詳細な段階的な解決策を求めることを学び、より良い結果が得られました。それでも、この実験では、著者らは主張を完全に検証できず、決定的ではない結果が得られました。
これ以外に、最も一般的な状況は次のようなものでした。
AI: 「論文と同じ、ABC という仮定の下で未解決の問題を解決しました。」
著者: 「おおむね機能していると思いますが、ステップ 10 が完全に一致しておらず、ステップ 20 が完全に正当化されておらず、仮定 C は私たちのものとは微妙に異なります。」
AI: 「同意します。実際、ステップ 10 と 20 は追加の仮定 DEF のもとでのみ証明できます。」
したがって、一言で言えば、AI は未解決の問題に対して部分的な解決策を提供しましたが、(1) 追加の仮定を課したときに問題が興味深いものであるかどうか、(2) AI が大きな問題に向かって本当に部分的にうまく進んだのか、それとも単に行き止まりに陥ったのかを理解するには、追加の作業または人間の判断が必要になります。
その他の興味深い混合結果の選択
「相互情報量からトポロジカルにエンコードされた状態の長距離非安定化」 という未解決の問題に対する可能な解決策についての議論の中で、著者は、AI が実際にその主張を証明できるかどうか、またその証明に使用される還元戦略が実際に有効であるかどうかについて疑問を表明しました。より詳細な返答の中で、AIは証明が不完全であることを認めた。しかし、削減戦略が実際に有効であることを示すことができました。それを部分的な勝利と呼ぶことにします。
「ホルスタインモデルに局所保存量が存在しないことの証明」では、著者らは

「将来の研究の重要な方向性は、私たちの方法を他の物理的に関連するフェルミ粒子と粒子の結合系に拡張することです。」と述べました。 AIはこれを見て、他の多くの物理的に関連するシステムでこの方法を試し、成功を宣言し、同じトリックがホルスタインモデルから「多成分ホルスタイン、ホルスタイン・ハバード鎖、軌道/ヤン・テラー・ハバード型鎖およびフレーバー依存局所電子フォノンを含む、オンサイトボソン支援フェルミオン双線形を有するすべての一次元光学フェルミオン粒子鎖」に拡張できることを証明した。モデル。」
著者の返答は基本的に「確かに、しかしそれは私たちのメソッドを拡張するという意味ではありません」というものでした。彼らの見解を要約すると、もちろん、この方法は同様のハミルトニアンに直接適用でき、同様の結果が得られますが、それは必ずしも興味深いことではありません。興味深い物理学は、これらの方法を使用してこれらのモデルについて何か新しいことを証明することで得られるでしょう。しかし、彼らが知る限り、AI が行ったことは予期せぬものや些細な応用ではありませんでした。
ここでは AI に部分的な点を与えます。著者が反発したとしても、著者にとっては些細な応用であっても、専門知識のない読者にとっては予期せぬものであった可能性はあると思います。しかし、それは結局のところ、物理学において問題を「解決する」ということが何を意味するのかを実際に定義するという課題を指していると私は思います。
もう 1 つの「曖昧な」例は、物理学の手順測定の統計的解釈から来ました。

[切り捨てられた]

## Original Extract

The joke is that every time that a new AI model comes out, someone throws it at the Collatz conjecture just in case. It’s actually probably not the worst idea, but “nails high-school physics but cannot yet create the theory of everything” is not the most useful description of the level of physics pr
[truncated]

-->
-->
Reading the quantum
About Ion ideas
Can AI answer open questions in physics?
The joke is that every time that a new AI model comes out, someone throws it at the Collatz conjecture just in case. It’s actually probably not the worst idea, but “nails high-school physics but cannot yet create the theory of everything” is not the most useful description of the level of physics proficiency of AI models. What I need to know is: as of mid-2026, how good is frontier AI at frontier physics exactly?
According to the published literature, the answer is “it’s getting there”. For example, CritPt compiled a benchmark of 50 research-level physics questions. At the time of their evaluation in November 2025, the top model scored a mere 12.6% . However, in the nine months since, the top score went up to over 30%, with a total of 10 models scoring better than November’s winner!
Likewise, consider TPBench , a benchmark for theoretical physics problems. In the preprint published in February 2025, the models pretty much nailed high-school physics, but scored up to ~50% on graduate-level physics and ~15% on research-level physics:
Rerunning the same benchmarks in 2026, the authors found that GPT-5.5 (xhigh) scored ~85% on graduate-level questions and a whopping 55% on research-level questions. That’s pretty astonishing, especially that GPT-5.5 (xhigh) is not even the most powerful physics model out there.
Still, the questions in those benchmarks are generally artificial: they are all concrete, solvable, and someone out there knows the solution. One can also worry that even if the author keeps the answer to themselves, the answer may one way or another “leak” on the internet, e.g. as part of a published research paper.
Thus, a few weeks ago, I decided to make an experiment where I throw frontier AI models at “real” open physics questions . Here is what I found.
What are the open questions in physics?
First clarification - by “open physics questions”, I don’t mean the ones you heard on national geographic, like “what’s the origin of the matter-antimatter asymmetry” !
While we’re all used to hearing about the “big” open questions in physics, the truth of the matter is that the field is full of “small” open questions. In fact, every day physicists around the world publish hundreds of papers and preprints, many of which contain “open questions” . This can be anything from a puzzling experimental observation that eludes explanation, through a lemma that seems reasonable but evades proof, all the way to the long-term proposals for follow-up work.
What’s nice about these open questions is that they’re truly open. Unlike TPBench, answering any of these questions would genuinely advance human knowledge - even if only by a tiny amount and on a very niche topic. Furthermore, unlike existing benchmarks, the question difficulty - or even whether a solution exists - is unknown a priori.
At the same time, I hypothesise that many of them are actually not that hard - a typical “open question” in a physics paper received a relatively small amount of attention from a very small number of humans. This makes them a promising target for AI input.
For the experiment, I started by scraping one week’s worth of preprints from the quant-ph arXiv, finding 275 papers with over 200 candidates for open questions based on grepping of phrases such as “open”, “unresolved”, and “future work”.
The first observation was that - unsurprisingly - the vast majority of my “open problems” were (in my opinion) not suitable for AI just yet. Typical reasons were that:
They’re about experimental / hardware validation . Example from my dataset: “operation at few-kelvin temperatures remains to be demonstrated” .
They’re vague outlines of future roadmap . Examples: “Future work will explore strong-coupling corrections beyond the Born–Markov approximation, non-Gaussian TLS noise, and machine-learning-based real-time estimation” , “future work on covert quantum communication under realistic uncertainty, including adaptive, correlated, deployment-oriented, and distributionally robust extensions”
They’re very broad . Example: “The description of the non-equilibrium dynamics of isolated quantum many-body systems […] is a fundamental open question.”
They’re heavy on practical engineering . Examples: a quantum simulation paper leaves “A full gate-count analysis, circuit synthesis […] and quantum-hardware execution” for future work / another paper asks for “a full end-to-end fault-tolerant distributed-quantum-computing compiler.”
They’re too big . Example: a paper on high-dimensional layer codes asks about “determining the finite-temperature self-correction properties of these 4D and 5D layer codes.”
They’re about broad protocol design . Examples: a QKD paper asks about “proving unconditional security for such non-Gaussian-measurement CV-QKD protocol” ; a paper on distributed molecular-node quantum computing asks to “design modular architectures and entanglement-generation protocols that achieve fault tolerance with feasible inter-module entanglement rates and hardware-resource requirements.”
I addressed this with a lightweight but aggressive (agentic) down-selection pipeline, aimed to identify questions most suitable for AI assistance. Some personal judgement went into this, but overall, questions received bonus points when they were:
Yes/no mathematical claims . Example: “does every entangled two-qubit state admitting two pure steered states must have density-matrix rank 2.”
Questions about formulas or numbers . Example: a paper establishes a threshold formula numerically, and asks for an analytic proof.
Optimization / characterization questions . Example: a paper asked for “analytically characterizing the maximal possible violation of the realignment criterion over PPT entangled states.”
Result extension questions . Example: a paper proved a result about single-qubit quantum channels, and asked for extension to multi-qubit channels.
Following this down-selection, 48 questions across 33 papers were extracted and packaged into stand-alone units, and the AI got to work on trying to crack them one by one.
My first attempt used an automated two-agent loop: claude opus 4.8 as the “solver”, and gpt 5.5-xhigh as the “reviewer”. This, however, resulted in a very low success rate. I then switched to gpt 5.5 Pro as the “solver”, with much better success rate. However, as I didn’t have access to that model through an API - only through web UI - the remaining part of the experiment was rather manual and less systematic.
GPT 5.5 proceeded to work on 48 problems across 33 papers
It claimed to successfully solve 42/48 of them, across 30 papers
I sent 30 emails to authors of these papers, asking for results verification
I received a response to 16/30 of these emails
The result of the verification has been:
25% (4/16): Bad. AI incorrectly claimed a proof, or misframed the problem
55% (9/16): Mixed. AI provided a partial or promising proof, but likely incomplete
20% (3/16): Good. AI confirmed to have effectively solved the problem.
Those are the headline numbers, but they hide the really interesting nuance. So let’s dig in into the examples of each, what they tell us about physics, and how to improve the experiment to improve its performance next time.
One example miss was the problem from Quantum randomness beyond projective measurements . Quote: “a simple question remains open: how much intrinsic randomness can be generated by a given extremal measurement?” . AI claimed to have solved it by reformulating it as a pure-state minimax optimization and doing the variational characterisation. Two problems:
The paper already did all of this, using a slightly different notation. So in essence, all that AI did was to re-formulate or re-derive some of the results in the paper.
A human reader would have understood that the real open problem was an explicit solution to the optimisation problem, not just formulating and characterising it
Another misunderstood open question came from The Marginal Problem for Density Operators . In the introduction, the authors made a historical remark that “Leifer & Poulin (2008) left open whether the entropic quantum conditional-independence relation has the full graphoid property” . AI interpreted this question as “does the quantum conditional-independence relation have the full graphoid property for arbitrary density operators “ , and was able to quickly find a counterexample.
However, the sentence before started with “In the strictly positive finite-dimensional setting considered here…“ . This should have made the AI realise that the actual question was about whether this property holds for strictly positive density operators . This also would have been obvious had the AI read the actual citation of Leifer & Poulin (2008).
The largest category, with many different types of partial solutions.
By far the most common author objection is that AI claimed a proof but did not provide enough detail to verify the maths. That was mainly my bad : I have since learned to always prompt the model for detailed step-by-step solutions, with much better results. Still, for this experiment, this led to some inconclusive results, with authors unable to fully verify the claims.
Other than this, the most common situation was along the lines of:
AI: “I have solved the open problem subject to assumptions ABC, same as the paper”
Author: “I think it mainly works, but step 10 doesn’t quite follow, step 20 is not fully justified, and assumption C is subtly different from ours “
AI: “I agree - in fact, I can only prove steps 10 and 20 under additional assumptions DEF”
Thus, in a nutshell, AI provided a partial solution to the open problem, but it would take additional work or human judgement to understand (1) whether the problem is interesting once the additional assumptions are imposed and (2) whether AI really made good partial progress towards the big question, or just followed a dead end.
Selection of other interesting mixed results
In the discussion about a possible solution to an open question from Long-range nonstabilizerness of topologically encoded states from mutual information , the author expressed doubt about whether AI can actually prove the claim, and whether the reduction strategy used in that proof is indeed valid. In a more detailed response, AI admitted the proof is incomplete. However, it was able to show that the reduction strategy is indeed valid. I’ll call it a partial win.
In “Proof of the absence of local conserved quantities in the Holstein model” , authors said “An important direction for future work is to extend our method to other physically relevant fermion–boson coupled systems” . AI saw this, tried the method on a bunch of other physically relevant systems, and declared success, proving that the same trick can be extended from the Holstein model onto “all one-dimensional optical fermion-boson chains with onsite boson-assisted fermion bilinears, including multicomponent Holstein, Holstein-Hubbard, orbital/Jahn-Teller-Hubbard-type chains and flavor-dependent local electron-phonon models.”
The author reply was essentially “sure, but that’s not what we meant by extending our method ”. Summarising their view: of course, the method can be directly applied to similar hamiltonians, obtaining similar results - but that’s not necessarily interesting. The interesting physics would come from proving something new about those models using those methods - but as far as they can tell, nothing that the AI did was unexpected or a trivial application.
I give AI partial points here. Even though the authors pushed back, I think it’s possible that what’s a trivial application to authors could have been unexpected to a less expert reader. But I think it really points down to the challenge of actually defining what it means to “solve” a problem in physics.
Another “fuzzy” example came from Statistical Interpretation of the Procedures Measurement of Phys

[truncated]
