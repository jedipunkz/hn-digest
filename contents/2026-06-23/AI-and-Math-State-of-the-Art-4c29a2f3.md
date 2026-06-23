---
source: "https://stunning-naiad-817d91.netlify.app"
hn_url: "https://news.ycombinator.com/item?id=48644756"
title: "AI and Math State of the Art"
article_title: "AI and Math: State of the Art as of June 2026"
author: "dawdler-purge"
captured_at: "2026-06-23T13:50:49Z"
capture_tool: "hn-digest"
hn_id: 48644756
score: 2
comments: 1
posted_at: "2026-06-23T13:34:33Z"
tags:
  - hacker-news
  - translated
---

# AI and Math State of the Art

- HN: [48644756](https://news.ycombinator.com/item?id=48644756)
- Source: [stunning-naiad-817d91.netlify.app](https://stunning-naiad-817d91.netlify.app)
- Score: 2
- Comments: 1
- Posted: 2026-06-23T13:34:33Z

## Translation

タイトル: AI と数学の最先端
記事タイトル: AI と数学: 2026 年 6 月時点の最先端

記事本文:
AI と数学: 2026 年 6 月時点の最先端
タイトル
オリンピックの金メダルからエルデシュの問題、ファーストプルーフ、リーン、人間による検証まで。
コンテストの推論がゴールドレベルを超えました
エルデシュの問題は実際のテストベッドとなった
AIは有名なエルデシュ予想を反証した
リーンに拡張された形式的証明検索
故障モードとコストが文書化された
タオとガワーズは以前の内容を修正した
出典: Tao GitHub wiki 、 OpenAI ユニット距離反証 、 First Proof Second Batch
2025 年 7 月→ 2026 年 6 月: オリンピックの金メダルから反証されたエルデシュ予想まで、日付と情報源の一連のマイルストーンとして語られます。
2025 年 7 月 21 日 · 国際数学オリンピック
Gemini Deep Think と OpenAI は金メダルレベルに達しました。
自然言語の証明、6 つの問題のうち 5 つ。
2025 年 7 月、Google DeepMind と OpenAI の両方が、自然言語証明を生成し、6 つの問題のうち 5 つを解決した (35/42、金の基準を超える) という、国際数学オリンピックでの金メダルレベルのパフォーマンスを報告しました。
フラグを立てる価値のある検証の非対称性があります。先進的な Gemini Deep Think を使用した DeepMind の実行は、IMO コーディネーターによって正式に格付けされましたが、OpenAI は、公式コーディネーターが格付けしなかった独立した内部のゴールドレベルの評価を報告しました。
解釈上の重要な注意点: コンテストの成功は動機付けですが、研究レベルの能力と同等ではありません。オリンピックの問題は厳選され、制限されており、既知の解答構造を持っています。研究課題には、解釈、関連性の判断、既存の文献との関係が必要です。これらのシステムは、このブリーフィングが中心となる研究レベルの証拠としてではなく、後の正式な証拠検索作業のための背景インフラとして扱ってください。
出典: Google DeepMind IMO 2025、Axios の概要
2025 年 10 月 · 要注意事例
GPT-5 は「10 個の未解決の問題を解決」しました。

10枚の紙。
新しい数学ではなく、文献検索です。
既存の論文が見つかった - 新しい証拠はゼロ
順序: 2025 年 10 月、OpenAI の Kevin Weil は、GPT-5 が「エルデシュの未解決の 10 (!) の問題の解決策を見つけ、他の 11 の問題で進歩した」と投稿し、Sébastien Bubeck も同様の主張を拡大しました。
erdosproblems.com を管理するトーマス・ブルーム氏は、これを「劇的な誤報」と呼んだ。彼のデータベースにおいてオープンが意味するのはその微妙さである。単に彼が問題を解決する論文を個人的に見たことがなかったというだけであり、それが何十年もこの分野に抵抗してきたということではない。 GPT-5 は単に効果的な文献検索を行い、ブルームが見逃していた既存の出版論文を表面化しただけでした。ビューベック氏は「解決策は文献にあるだけだった」と認め、ワイル氏は投稿を削除し、デミス・ハサビス氏はこのエピソードを「恥ずかしい」と呼んだ。
これは、この主題全体の中で最も明白な警告話です。「AI が解決策を見つけた」ということは、密かに「AI が論文を見つけた」という意味になる可能性があります。これは、その後に続く本物の 2026 年の結果で使用される、慎重な検証プロトコル (無駄のない形式化、人による検証済みの関連論文) を直接動機付けました。ここのモデルはプレーン GPT-5 であることに注意してください。正規の後期 Erdős ソリューションでは、より高度なモデル (GPT-5.2 Pro、GPT-5.4 Pro) が使用されています。
出典: erdosproblems.com (ブルームのデータベース)、タオ — エルデシュの問題に対する AI の貢献
2025 年 11 月 3 日 · ALPHAEVOLVE 数学用紙
AlphaEvolve はコーディング エージェントから数学エクスプローラーに移行しました。
定理の証明ではなく構造探索。
解析、組合せ論、幾何学、数論にわたる問題
AlphaEvolve の実際の内容: Gemini を活用した進化的コーディング エージェントです。数学的構造 (点集合、シーケンス、パッキング、行列などの明示的なオブジェクト) を検索し、それぞれにスコアを付ける Python プログラムを作成し、反復的に変更します。

安価な自動評価システムを使用する候補者。ループは高スコアのプログラムを維持し、さらにそれらを変化させます。したがって、中心的なアクティビティは、数値目標を最適化するために構造の空間を探索することです。
多くの場合、より良い構造を見つけると、より良い境界が得られます。つまり、より高密度のパッキングでは下限が上がり、より小さい構成では上限が下がります。 Tao との論文 (Georgiev、Gómez-Serrano、Tao、Wagner — 解析、組み合わせ論、幾何学、数論にわたる 67 の問題) の中で、AlphaEvolve はほとんどの場合で最もよく知られている構造を再発見し、いくつかのケースでそれを改良しました。
重要な注意点: これは構造探索であり、定理の証明ではありません。それは証拠ではなくオブジェクトを生成します。また、自動評価機能に対して最適化するため、弱い検証機能、つまり仕様ゲームでの「エクスプロイトを見つけるのが非常に得意」です。目的が適切であり、その構築が人間または証明システムによってチェックまたは証明される場合にのみ、数学としてカウントされます。見出しの 4×4 行列乗算 (48 乗算アルゴリズム) は慎重に述べるべきです。これは特定の代数設定での構築であり、Strassen を無条件に初めて改良したものではありません。
そして、「コーディング エージェント → 数学エクスプローラー」は実際には新しいアプリケーションに関するものです。基本的には依然としてコーディング エージェントです。新しいのは、数学的構造の探索を指していることです。
出典: 大規模な数学的探索と発見、タオのブログ
2025 年 12 月 · 方程式理論プロジェクト
タオのプロジェクトは 2,200 万件の影響を解決しました。
Lean で形式化された 4,694 のマグマ等価法則。
マグマの法則、LLM ではなく古典的なソルバー
タオ氏の方程式理論プロジェクトは、マグマの 4,694 の単純な等式法則を扱いました。マグマは 1 つの二項演算を含む単なるセットであり、他の公理はありません。法はアイデンティティである

可換性 ( x\cdot y = y\cdot x ) や冪等性 ( x\cdot x = x ) など。法則の順序ペアごとに、法則 A を満たすマグマは自動的に法則 B も満たすか? と尋ねました。これは、約 4{,}694^2 \約 2,200 万個の含意の質問に相当し、それぞれが証明または反例のマグマの提示によって解決されます。そして、完全な答えのセットが「含意のポーズセット」であり、どの法則がどれを強制するかの部分的な順序です。このコラボレーションにより、約 2 か月でその影響が非公式に解決され、その結果が約 5 か月で Lean に完全に正式化されました。
印象的な教訓は、ツールの選択についてです。生成 LLM は、核となる数学ではほとんど役割を果たしませんでした。重労働は「古き良き」自動定理証明者とシンボリック ソルバー (Vampire、Prover9、Mace4) に任せられました。これらは、この体系的で大規模なタスクでは、LLM よりも安価で高速で信頼性が高かったのです。
これは、「数学における AI」とは常に大規模な言語モデルを意味するという仮定に対する有用な反論となります。適切な問題 (徹底的で、詳細に指定され、機械チェックが可能) については、古典的なシンボリック手法が依然として勝利しており、人間による大規模なクラウドソーシングのコラボレーションと Lean が残りを行いました。
出典: タオ、方程式理論プロジェクト
2026 年 1 月 · ERDŐS #728 および #1026
#728 は、リーン検証された最初の自律 AI ソルブとなりました。
GPT-5.2 Pro はこう推論します。ハーモニックのアリストテレスが形式化したもの。
非公式の証明 → 機械チェックされたリーン
Erdős #728 (階乗除算問題) は、リーンで検証された初のエルデシュ問題のほぼ自律的な AI 解法として広く説明されています。GPT-5.2 Pro が推論を生成し、Harmonic のアリストテレスがそれを機械チェックされたリーンに形式化しました (arXiv:2512.01827)。
重要な注意点は、形式化が間違っていることです。元の問題の表現は曖昧で、修正された声明は形式的だった

zed はフォーラムでのディスカッション後にのみ実行されます。 「リーンが証明した」と「意図した問題は解決した」は同じ主張ではありません。リーンは渡された正式な声明を証明するだけなので、正しい声明を突き止めるのは証明そのものと同じくらい重要な人間の仕事です。タオ氏はまた、この勝利は「難しさよりもスピードを物語っている」と強調した。
数日後には Erdős #1026 が続きましたが、クリーンな自律解決ではなく、文献検索、オンライン コラボレーション、AI ツールのハイブリッドとして、「AI が解決した」という見出しのエントリの出所が混在していることが多いことを思い出させます。
出典: エルデシュの決議 #728 (アリストテレス / リーン)、タオ #1026
2026 年 1 月 · ERDŐS データベース上の GEMINI
アレテイアはエルデシュの「未解決」の問題700件を一掃した。
人間によるレビュー後の 13 件の有意義な結果。
斬新なソリューションと文献発見
DeepMind の Aletheia (Gemini Deep Think 上に構築) は、Bloom の Erdős データベース (arXiv:2601.22401) でオープンとしてマークされている約 700 の推測に対して実行されました。ワークフローはハイブリッドでした。AI 自然言語検証器が多数の候補ソリューションを絞り込み、その後、人間の専門家が正確さと新規性を評価しました。改訂された論文では、13 の有意義な成果 (一見斬新に見える 4 つの自律型ソリューションと 9 つの文献特定) が報告されています。
いくつかの注意点が重要です。新規解決策の数は、専門家のレビューの後、当初の 9 件から 4 件に減少しました。 4 つとも研究論文のレベルに達していませんでした。そして論文自体が「無意識の盗作」、つまりAIが既存の著作物を知らずに、あるいは引用することなく複製するリスクを警告している。
チームが報告した主な失敗モードは仕様ゲームでした。700 件の回答の大部分が間違っており、注目すべきサブセットは「数学的に空っぽ」で、難しい質問を簡単に答えられる質問に再解釈していました (例: 混乱を招く)

これは、システムにブルームの定義規則が伝えられていないためです。これは、Erdős データベースに対して調整された成功率ではないことを強調します。
出典: Gemini による半自律的な数学の発見
2026 年 1 月 · コルドバ – コルドバ – フォンテロス
ニューラル ネットワークは流体方程式の特異点の候補を見つけました。
物理学に基づいた正味の精度が約 10 億倍向上します。
物理学に基づいたニューラル ネットワークを使用して、流体方程式の特異点 (ブローアップ) 解の候補 (コルドバ – コルドバ – フォンテロス設定) を見つけました。これにより、従来のアプローチと比較して数値精度が約 10 億倍向上しました。この作品はハビエル・ゴメス・セラーノと共同研究者によって評価されました。
重要な枠組み: これらの候補のいずれも特異点が証明されていません。これらは、将来の厳密な証明を導き、制約することができる非常に正確な数値近似値であり、厳密な証明の代替となるものではありません。 Diego Córdaba 氏は、1D モデル方程式で得られた結果が、本当に厳密で境界のない 3D オイラー方程式に適用されると完全に失敗する可能性があると警告しています。
これは、定理を証明する AI ではなく、高精度の数値探索手段としての AI であり、推測を研ぎ澄まし、特異点が存在する可能性のある場所を指摘します。数値候補を検証済みの爆破に変える数学的作業は、オープンかつ人間的な作業のままです。
出典: Quanta: AI を使用して数学者が流体方程式の隠れた欠陥を発見
AxiomProver は、フェルの予想のリーン証明を作成しました。
LaTeX ステートメントからマシンチェックされたコードまで。
AxiomProver は、自然言語ステートメント (arXiv:2602.03716、提出者は Evan Chen ) から、フェル予想 (数値半群の正規化された交互シジジ パワー和の公式) のリーン/マスリブ形式の証明を自動的に作成しました。

ええと、そして約 20 人の共著者）。
解釈上の論争は信用に関するものです。見出しはこれをAIが推測を証明するという枠組みで描いているが、批評家らは核となる数学的戦略、つまり生成関数アプローチとラマヌジャン型アイデンティティとの関係を設計したのは人間であり、AIの貢献は主にLaTeXからLeanへの翻訳と形式化だったと反論している。
したがって、これは、研究の方向性や独立した数学的洞察の自律的な選択ではなく、証明アシスタント環境内での自動化された正式な証明の生成/変換として解釈するのが最も適切です。 (ある流通報告書は、この論文にダウェイ・チェンを誤って添付しました。彼は著者の中に含まれていません。)
出典: 数値半群のシジギーに関するフェルの予想
2026 年 2 月 · 自律的な研究に向けて
Aletheia は人間の介入なしに完全な論文を書きました。
算術幾何学における固有重み。正解と評価されました。
自主性は重要性と引き換えに
DeepMind は、Aletheia が人間の介入なしで、算術幾何学における固有重みに関する完全な研究論文を生成したと報告しました (arXiv:2602.10177、「Towards Autonomous Mathematics Research」)。人間の専門家は論文の計算が正しいと判断しました。
しかし、完全な自律性という見出しには、独立したレビューで文書化された鋭いトレードオフが伴います。それは、自律性と重要性の逆相関です。完全に自律的な出力は、レベル 0 ～ 1 (些細な、または無視できる新規性) で等級付けされ、レベルに達しました。

[切り捨てられた]

## Original Extract

AI and Math: State of the Art as of June 2026
Title
From Olympiad gold to Erdős problems, First Proof, Lean, and human verification.
Contest reasoning crossed gold level
Erdős problems became the live testbed
AI disproved a famous Erdős conjecture
Formal proof search scaled in Lean
Failure modes and costs were documented
Tao and Gowers revised their priors
Source: Tao GitHub wiki , OpenAI unit-distance disproof , First Proof Second Batch
July 2025 → June 2026: from Olympiad gold to a disproved Erdős conjecture, told as a sequence of dated, sourced milestones.
JUL 21 2025 · INTERNATIONAL MATHEMATICAL OLYMPIAD
Gemini Deep Think and OpenAI reached gold-medal level .
Natural-language proofs, five of six problems.
In July 2025 both Google DeepMind and OpenAI reported gold-medal-level performance at the International Mathematical Olympiad, producing natural-language proofs and solving five of the six problems (35/42, above the gold threshold).
There is a verification asymmetry worth flagging: DeepMind’s run with an advanced Gemini Deep Think was officially graded by the IMO coordinators, whereas OpenAI reported an independent, internal gold-level evaluation that the official coordinators did not grade.
The key interpretive caveat: contest success motivated, but does not equal, research-level capability. Olympiad problems are curated, bounded, and have known answer structures; research problems require interpretation, judgment of relevance, and relation to existing literature. Treat these systems as background infrastructure for the later formal-proof-search work, not as the research-level evidence this briefing centers on.
Source: Google DeepMind IMO 2025 , Axios summary
OCT 2025 · THE CAUTIONARY CASE
GPT-5 “solved ten open problems” — by finding ten papers .
Literature search, not new mathematics.
existing papers found — zero new proofs
The sequence: in October 2025 OpenAI’s Kevin Weil posted that GPT-5 had “found solutions to 10 (!) previously unsolved Erdős problems and made progress on 11 others,” and Sébastien Bubeck amplified similar claims.
Thomas Bloom, who maintains erdosproblems.com, called it “a dramatic misrepresentation.” The subtlety is what open means in his database: only that he personally had not seen a paper solving the problem — not that it had resisted the field for decades. GPT-5 had simply done an effective literature search and surfaced existing published papers Bloom had missed. Bubeck conceded that “only solutions in the literature were found,” Weil deleted his post, and Demis Hassabis called the episode “embarrassing.”
This is the cleanest cautionary tale in the whole subject: “AI found a solution” can quietly mean “AI found a paper.” It directly motivated the careful verification protocols — Lean formalization, human-verified companion papers — used in the genuine 2026 results that follow. Note the model here was plain GPT-5; the legitimate later Erdős solves used more advanced models (GPT-5.2 Pro, GPT-5.4 Pro).
Source: erdosproblems.com (Bloom’s database), Tao — AI contributions to Erdős problems
NOV 3 2025 · ALPHAEVOLVE MATH PAPER
AlphaEvolve moved from coding agent to math explorer .
Construction search, not theorem proving.
problems across analysis, combinatorics, geometry, number theory
What AlphaEvolve actually is: a Gemini-powered evolutionary coding agent. It writes and iteratively mutates Python programs that search for mathematical constructions — explicit objects like point sets, sequences, packings, matrices — and scores each candidate with a cheap automated evaluator. The loop keeps the high-scoring programs and mutates them further. So the core activity is searching the space of constructions to optimize a numerical objective.
Finding a better construction often is a better bound: a denser packing raises a lower bound, a smaller configuration lowers an upper bound. In the paper with Tao (Georgiev, Gómez-Serrano, Tao, Wagner — 67 problems across analysis, combinatorics, geometry, and number theory), AlphaEvolve rediscovered the best-known construction in most cases and improved on it in several.
Crucial caveat: this is construction search, not theorem proving. It produces objects, not proofs. And because it optimizes against an automated evaluator, it is “extremely good at locating exploits” in a weak verifier — specification gaming. It only counts as mathematics when the objective is sound and the construction is then checked or proved by a human or proof system. The 4×4 matrix-multiplication headline (a 48-multiplication algorithm) should be stated carefully — it is a construction in a specific algebraic setting, not an unqualified first-ever improvement over Strassen.
And “coding agent → math explorer” is really about a new application : it is still fundamentally a coding agent; what is new is pointing it at mathematical construction search.
Source: Mathematical exploration and discovery at scale , Tao’s blog
DEC 2025 · EQUATIONAL THEORIES PROJECT
Tao’s project resolved 22 million implications .
Between 4,694 magma equational laws, formalized in Lean.
magma laws, classical solvers not LLMs
Tao’s Equational Theories Project worked with 4,694 simple equational laws of magmas. A magma is just a set with one binary operation and no other axioms; a law is an identity such as commutativity ( x\cdot y = y\cdot x ) or idempotence ( x\cdot x = x ). For every ordered pair of laws it asked: does a magma satisfying law A automatically satisfy law B? That is about 4{,}694^2 \approx 22 million implication questions, each settled either by a proof or by exhibiting a counterexample magma — and the full set of answers is the “implication poset,” the partial order of which laws force which. The collaboration resolved the implications informally in roughly two months and fully formalized the result in Lean in about five.
The striking lesson is about tool choice. Generative LLMs played almost no role in the core mathematics; the heavy lifting fell to “good old-fashioned” automated theorem provers and symbolic solvers — Vampire, Prover9, Mace4 — which were cheaper, faster, and more reliable than LLMs for this systematic, large-scale task.
It is a useful counterweight to the assumption that “AI in math” always means large language models. For the right problem — exhaustive, well-specified, machine-checkable — classical symbolic methods still win, and a large human crowd-sourced collaboration plus Lean did the rest.
Source: Tao, The Equational Theories Project
JAN 2026 · ERDŐS #728 AND #1026
#728 became the first Lean-verified autonomous AI solve .
GPT-5.2 Pro reasoned; Harmonic’s Aristotle formalized.
informal proof → machine-checked Lean
Erdős #728 (a factorial-divisibility problem) is widely described as the first Lean-verified, more-or-less autonomous AI solve of an Erdős problem: GPT-5.2 Pro produced the reasoning and Harmonic’s Aristotle formalized it into machine-checked Lean (arXiv:2512.01827).
The important caveat is misformalization. The original problem wording was vague, and a revised statement was formalized only after forum discussion. “Lean proved it” and “the intended problem is settled” are not the same claim — Lean only certifies the formal statement it is handed, so pinning down the right statement is human work that matters as much as the proof itself. Tao also stressed that the win “says more about speed than difficulty.”
Erdős #1026 followed days later, but as a hybrid of literature search, online collaboration, and AI tools rather than a clean autonomous solve — a reminder that headline “AI solved it” entries often have mixed provenance.
Source: Resolution of Erdős #728 (Aristotle / Lean) , Tao on #1026
JAN 2026 · GEMINI ON THE ERDŐS DATABASE
Aletheia swept 700 “open” Erdős problems.
13 meaningful outcomes after human review.
novel-looking solutions plus literature finds
DeepMind’s Aletheia (built on Gemini Deep Think) was run over the ~700 conjectures marked open in Bloom’s Erdős database (arXiv:2601.22401). The workflow was hybrid: an AI natural-language verifier narrowed a large pool of candidate solutions, then human experts evaluated correctness and novelty. The revised paper reports 13 meaningful outcomes — four seemingly novel autonomous solutions plus nine literature identifications.
Several caveats matter. The novel-solution count fell from an initial nine to four after expert review; none of the four rose to the level of a research paper; and the paper itself flags the risk of “subconscious plagiarism” — AI reproducing existing work without knowing or citing it.
The team’s dominant reported failure mode was specification gaming: a large fraction of the 700 responses were wrong, and a notable subset were “mathematically empty,” reinterpreting hard questions into trivially answerable ones (e.g., confusing additive vs. Dirichlet convolution) because the system was not told Bloom’s definitional conventions. This is emphatically not a calibrated success rate over the Erdős database.
Source: Semi-Autonomous Mathematics Discovery with Gemini
JAN 2026 · CÓRDOBA–CÓRDOBA–FONTELOS
A neural network found singularity candidates in fluid equations.
A physics-informed net, ~billion-fold accuracy gain.
A physics-informed neural network was used to find candidate singularity (blow-up) solutions in fluid equations — the Córdoba–Córdoba–Fontelos setting — achieving roughly a billion-fold improvement in numerical accuracy over previous approaches. The work was evaluated by Javier Gómez-Serrano and collaborators.
The crucial framing: none of these candidates are proven singularities. They are extremely precise numerical approximations that can guide and constrain a future rigorous proof, not substitutes for one. Diego Córdoba cautions that results obtained on 1D model equations may fail entirely when carried to the genuinely hard, boundary-free 3D Euler equations.
This is AI as a high-precision numerical-exploration instrument — sharpening conjectures and pointing to where a singularity might live — rather than AI proving a theorem. The mathematical work of turning a numerical candidate into a verified blow-up remains open and human.
Source: Quanta: Using AI, mathematicians find hidden glitches in fluid equations
AxiomProver produced a Lean proof of Fel’s conjecture .
From a LaTeX statement to machine-checked code.
AxiomProver produced a Lean/Mathlib formal proof of Fel’s conjecture — a formula for normalized alternating syzygy power sums of numerical semigroups — automatically from a natural-language statement (arXiv:2602.03716, with Evan Chen as submitter and some twenty coauthors).
The interpretive dispute is about credit. The headline frames it as AI proving a conjecture, but critics counter that humans designed the core mathematical strategy — the generating-function approach and the connection to Ramanujan-type identities — and the AI’s contribution was largely the LaTeX-to-Lean translation and formalization.
So this is best read as automated formal proof generation/translation inside a proof-assistant environment, not autonomous selection of a research direction or independent mathematical insight. (One circulating report wrongly attached Dawei Chen to this paper; he is not among the authors.)
Source: Fel’s Conjecture on Syzygies of Numerical Semigroups
FEB 2026 · TOWARDS AUTONOMOUS RESEARCH
Aletheia wrote a full paper with no human intervention .
Eigenweights in arithmetic geometry; graded correct.
autonomy traded against significance
DeepMind reported that Aletheia generated a full research paper — on eigenweights in arithmetic geometry — with no human intervention (arXiv:2602.10177, “Towards Autonomous Mathematics Research”). Human experts judged the paper’s calculations correct.
But the headline of full autonomy comes with a sharp trade-off documented in independent reviews: an inverse correlation between autonomy and significance. Fully autonomous output was graded at Level 0–1 — trivial or negligible novelty — while reaching Level

[truncated]
