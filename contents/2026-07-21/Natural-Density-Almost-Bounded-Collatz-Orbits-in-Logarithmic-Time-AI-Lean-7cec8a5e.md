---
source: "https://www.proofatlas.ai/formalizations/natural-density-log-time-collatz/"
hn_url: "https://news.ycombinator.com/item?id=48996871"
title: "Natural-Density Almost-Bounded Collatz Orbits in Logarithmic Time (AI, Lean)"
article_title: "Natural-density logarithmic Collatz descent | ProofAtlas"
author: "zone411"
captured_at: "2026-07-21T20:15:09Z"
capture_tool: "hn-digest"
hn_id: 48996871
score: 1
comments: 0
posted_at: "2026-07-21T19:18:43Z"
tags:
  - hacker-news
  - translated
---

# Natural-Density Almost-Bounded Collatz Orbits in Logarithmic Time (AI, Lean)

- HN: [48996871](https://news.ycombinator.com/item?id=48996871)
- Source: [www.proofatlas.ai](https://www.proofatlas.ai/formalizations/natural-density-log-time-collatz/)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T19:18:43Z

## Translation

タイトル: 対数時間での自然密度ほぼ有界コラッツ軌道 (AI、リーン)
記事のタイトル: 自然密度対数コラッツ降下法 |プルーフアトラス
説明: 自然密度 1 の多くの開始点は、明示的な対数シラキュースおよび生のコラッツ クロック内のすべての発散しきい値を下回ります。

記事本文:
コンテンツへスキップ ProofAtlas 形式化カタログ Mathlib ランドマーク Mathlib 証拠の仕組み 信頼コラボレーション エージェント ProofAtlas / 形式化 / 自然密度 対数時間におけるコラッツ降下 数理論 / 力学系 / 形式定理
対数時間における自然密度コラッツ降下法
奇数の入力に沿って無限大に向かう傾向のあるしきい値の場合、奇数相対密度 1 の多くの奇数の開始は 145 · log N シラキュース ステップ内で下降します。すべての正の入力で無限大になる傾向があるしきい値の場合、通常-自然-密度-1 多くの正の開始は 436 · log N の生のコラッツ ステップ内で下降します。
2 つの対数時間密度の結論が一目で分かる
証明では、グローバルべき乗則位相ギャップと固定量的レートを使用し、固定ターゲット制御を奇数相対シラキュース結論に変換し、次に 2 進リフトを使用して通常密度の生のコラッツ結論を取得します。
2 つの対数クロック内の 2 つの密度ドメイン
拡大する シラキュースのエンドポイントは、奇数の開始点間で奇数の相対密度を使用します。生の Collat​​z エンドポイントは、正の開始点の中で通常の密度を使用します。奇数相対密度(シラキュースヒット) = 1、C_syr < 145;通常の密度(生のコラッツヒット) = 1 (C_coll < 436)
奇数の入力に沿ってしきい値が増加する場合、奇数相対密度 1 の多くの奇数開始では、145 · log N の奇数から奇数へのステップの前に Syracuse ヒットがあります。すべての正の入力でしきい値が増加する場合、通常の自然密度 1 つの多くの正の開始では、436 · log N の個々のステップの前に生の Collat​​z ヒットがあります。
平方根対数時間ウィンドウの概要
自然密度定理は、平方根しきい値の上限クロックを提供します。決定論的な半分の引数は厳密に低いクロックを提供し、チェックされた結果は両方の不等式を満たす 1 つの証人を保持します。
内部での平方根ヒット

対数ウィンドウ
拡大 自然密度 1 多スタートの場合、√N 未満の生のコラッツ ヒットは真の対数時間ウィンドウ内で発生します。 log N/(2 log 2) < m ≤ C_coll log N < 436 log N および Collatz^m(N) < √N
同じ監視時間は厳密に log N を 2 log 2 で割った値よりも大きく、チェックされた生の Collatz クロック (436 · log N より小さい) よりも大きくはありません。下限は平方根ターゲットに特別です。
これらの AI 生成のビジュアルは、定理と証明ルートを説明します。それらは証拠ではありません。彼らの出版物の審査は、正式な結果の審査とは別に完了しました。正確なリーン提案とチェックされたソースは引き続き信頼できます。
対数時間での自然密度ほぼ有界コラッツ軌道 ↗
レッヒ・マズール · 2026-07-16 · 27 ページ
自然密度 1 対数時間のコラッツ降下定理、そのライン位相ギャップ入力、定量的レート アーキテクチャ、シラキュースから生時間へのブリッジ、および平方根時間ウィンドウの帰結を展開する数学論文。
権利者によって許可されたホスティング。オリジナルの ProofAtlas メタデータ カバー。紙のページの複製ではありません。
Lech Mazur、「対数時間における自然密度のほぼ有界のコラッツ軌道」、バージョン 1、2026 年 7 月 16 日。
この論文は、同じ定理群に関連した数学的解説です。正確にチェックされたリーン宣言とピン留めされたソースは、表現が異なっていても信頼できるままです。
この論文は、密度ゼロの例外集合を許可する定理について議論していますが、完全なコラッツ予想やすべての軌道の収束を主張するものではありません。
アトラスは、証明の依存関係を、より強力な兄弟の結果や有用な比較から切り離して保持します。以下の 2 つのレーンは切断されています。どちらの先行者カウント ファミリも密度と時間ファミリへの入力ではありません。
3 つの正確な 0.90 定理の変形
Rh

フィードの定量的レート エンジン
依存関係パスの確認: Rhin 位相ギャップ → ND31 メイン → ND31 境界 → 同一指数率 → 固定率 → 2 つの兄弟密度ファミリー エンドポイント。 6 つの保持された depend_on エッジがこの短縮パスをサポートします。
比較のみ: Terras の結果は、自然密度証明への入力ではなく、別個のコンパニオンとして明示的に記録されます。
推論されたエッジがない: 共有ソース ファイル、共通の主題、または歴史的背景によって、定理の依存関係が作成されません。
コンパニオンの結果 · 依存関係ではありません
有限の省電力停止限界を比較する
Terras スタイルの定理は、開始値を下回る降下に対する明示的な例外比率限界を与えます。これは別のチェックされたルートであり、この自然密度証明への入力ではありません。
この形式化が主張していないこと
これは、コラッツ予想、開始ごとの収束、または 1 への到達を証明するものではありません。密度ゼロの例外セットが残る場合があります。
しきい値は無限大になる傾向がある必要がありますが、単調である必要はなく、それ以下のヒットは厳密です。
奇数相対シラキュースの結論と通常密度の生のコラッツの結論は異なる領域を持ち、すべての肯定的なスタートに関する 1 つの主張にまとめてはなりません。
145 境界は奇数から奇数のシラキュース ステップをカウントしますが、436 境界は半分を含む生のコラッツ ステップをカウントします。
平方根の下側クロックは、一般的な成長しきい値定理ではなく、付随する平方根の系にのみ属します。
別の Terras の省電力有限停止定理は、関連する比較であり、この証明への入力ではありません。
各定理のページには、その拡張された正確な命題、視覚的な説明、完全にチェックされたソース、およびチェッカー証拠が含まれています。より長い証明ルートには、段階的なウォークスルーも含まれています。
ロガリットの自然密度ほぼ有界コラッツ軌道

フミックタイム
奇数の入力に沿って無限大に向かう傾向のあるしきい値の場合、奇数相対密度 1 の多くの奇数の開始は 145 · log N シラキュース ステップ内で下降します。すべての正の入力で無限大になる傾向があるしきい値の場合、通常-自然-密度-1 多くの正の開始は 436 · log N の生のコラッツ ステップ内で下降します。
平方根降下法には対数の生時間ウィンドウがあります
自然密度 1 の多くの正の開始 N の場合、生のコラッツ反復は、log N / (2 log 2) ステップを超え、最大 436 · log N ステップで厳密に √N を下回ります。
チェックされた定理から構築する
チェックされた自然密度定理を使用して、より強力な時定数、制限されたしきい値クラスの明示的な密度収束率、または算術位相ギャップが定量的な混合と降下をもたらすその他の動的システムを研究します。
各 ZIP には、チェックされたファーストパーティのリーン インポート クロージャ、正確なステートメントと境界、ライセンス、通知、証拠、ソース フットプリント マニフェスト、およびエージェント継続ファイルが含まれています。 Mathlib およびその他のサードパーティの依存関係はバンドルされていません。
何がチェックされるのか、そしてそれをどの程度のソースがサポートしているのか
カウントの仕組み: 行数には空白行が含まれません。コメントとドキュメントはカウントされます。合計は、重複排除され、コミット固定されたファーストパーティのリーン インポート クロージャです。 Mathlib およびその他のサードパーティの依存関係は除外されます。宣言数とは、アーティファクトの記録された証拠によってカバーされる名前を意味します。これは、ソース内のすべての宣言の数ではありません。ソース フットプリントは、難易度や校正品質のスコアではありません。
独立した出版物のレビュー
形式定理の出版ゲートが受け入れられます
リーンは証拠を確認します。独立した AI レビューにより、証拠の完全性、ステートメントの整合性、結果の境界、保持された定理の文言が個別に承認されました。これらのゲートは正式な結果に適用されます。生成されたメディア i

は個別にレビューされ、促進されます。どちらのレビューもリーンの証明チェックに取って代わるものでも、定理を拡張するものでもありません。
独立したレビューでは、記録されたビルド、正確な宣言、未完了のステップのスキャン、および公理の証拠を受け入れました。
正式な宣言は、指定された定理とその正確な変形に対して受け入れられました。
受け入れられた境界により、近くにあるより強力な主張や一般的に混同されている主張が範囲外に保たれます。
独立したレビューは、保持された定理の説明とソースの提示を受け入れました。生成されたメディアは、別のレビューとプロモーション ゲートに従います。
ファーストパーティのソース リンクは、チェックされたパッケージのコミットと正確なリーン ファイルに固定されます。
検証された承認結果レコードは、4 つのレビューをチェック済みの形式化にバインドします。

## Original Extract

Natural-density-one many starts fall below every diverging threshold within explicit logarithmic Syracuse and raw Collatz clocks.

Skip to content ProofAtlas Formalizations Catalog Mathlib landmarks Mathlib How evidence works Trust Collaboration Agents ProofAtlas / Formalizations / Natural-Density Collatz Descent in Logarithmic Time Number theory · dynamical systems · formal theorem
Natural-Density Collatz Descent in Logarithmic Time
For thresholds tending to infinity along odd inputs, odd-relative-density-one many odd starts descend within 145 · log N Syracuse steps; for thresholds tending to infinity on all positive inputs, ordinary-natural-density-one many positive starts descend within 436 · log N raw Collatz steps.
Two logarithmic-time density conclusions at a glance
The proof uses a global power-law phase gap and fixed quantitative rate, turns fixed-target control into an odd-relative Syracuse conclusion, then uses the two-adic lift to obtain the ordinary-density raw Collatz conclusion.
Two density domains within two logarithmic clocks
Enlarge The Syracuse endpoint uses odd-relative density among odd starts; the raw Collatz endpoint uses ordinary density among positive starts. odd-relative density(Syracuse hit) = 1 with C_syr < 145; ordinary density(raw Collatz hit) = 1 with C_coll < 436
For thresholds growing along odd inputs, odd-relative-density-one many odd starts have a Syracuse hit before 145 · log N odd-to-odd steps. For thresholds growing on all positive inputs, ordinary-natural-density-one many positive starts have a raw Collatz hit before 436 · log N individual steps.
Square-root logarithmic time window at a glance
The natural-density theorem supplies the upper clock for the square-root threshold. A deterministic halving argument supplies the strict lower clock, and the checked corollary retains one witness satisfying both inequalities.
A square-root hit inside a logarithmic window
Enlarge For natural-density-one many starts, a raw Collatz hit below √N occurs inside a genuine logarithmic time window. log N/(2 log 2) < m ≤ C_coll log N < 436 log N and Collatz^m(N) < √N
The same witness time is strictly greater than log N divided by 2 log 2 and no greater than the checked raw Collatz clock, which is below 436 · log N. The lower bound is special to the square-root target.
These AI-generated visuals explain the theorem and proof route; they are not proof evidence. Their publication review was completed separately from review of the formal result. The exact Lean proposition and checked source remain authoritative.
Natural-Density Almost-Bounded Collatz Orbits in Logarithmic Time ↗
Lech Mazur · 2026-07-16 · 27 pages
A mathematical paper developing the natural-density-one logarithmic-time Collatz descent theorem, its Rhin phase-gap input, the quantitative rate architecture, the Syracuse-to-raw-time bridge, and the square-root time-window corollary.
Hosting authorized by the rightsholder. Original ProofAtlas metadata cover; not a reproduction of a paper page.
Lech Mazur, “Natural-Density Almost-Bounded Collatz Orbits in Logarithmic Time,” version 1, 16 July 2026.
The paper is mathematical exposition linked to the same theorem family; the exact checked Lean declarations and pinned source remain authoritative if wording differs.
The paper discusses a theorem that permits a density-zero exceptional set and does not claim the full Collatz conjecture or convergence of every orbit.
The atlas keeps proof dependencies separate from stronger sibling results and useful comparisons. The two lanes below are disconnected: neither predecessor-count family is an input to the density-and-time family.
Three exact 0.90 theorem variants
Rhin feeds the quantitative rate engine
Reviewed dependency path: Rhin phase gap → ND31 main → ND31 bounds → same-exponent rate → fixed rate → the two sibling density-family endpoints. Six retained depends_on edges support this contracted path.
Comparison only: the Terras result is explicitly recorded as a separate companion, not an input to the natural-density proof.
No inferred edge: shared source files, a common subject, or historical background do not create a theorem dependency.
Companion result · not a dependency
Compare the finite power-saving stopping bound
The Terras-style theorem gives an explicit exceptional-proportion bound for descent below the starting value. It is a separate checked route, not an input to this natural-density proof.
What this formalization does not claim
This does not prove the Collatz conjecture, convergence for every start, or arrival at 1; a density-zero exceptional set may remain.
The threshold must tend to infinity but need not be monotone, and the hit below it is strict.
The odd-relative Syracuse conclusion and ordinary-density raw Collatz conclusion have different domains and must not be collapsed into one claim about all positive starts.
The 145 bound counts odd-to-odd Syracuse steps, while the 436 bound counts raw Collatz steps including halvings.
The square-root lower clock belongs only to the companion square-root corollary, not to the general growing-threshold theorem.
The separate Terras power-saving finite-stopping theorem is a companion comparison, not an input to this proof.
Each theorem page includes its expanded exact proposition, visual explanation, complete checked source, and checker evidence. Longer proof routes also include a step-by-step walkthrough.
Natural-Density Almost-Bounded Collatz Orbits in Logarithmic Time
For thresholds tending to infinity along odd inputs, odd-relative-density-one many odd starts descend within 145 · log N Syracuse steps; for thresholds tending to infinity on all positive inputs, ordinary-natural-density-one many positive starts descend within 436 · log N raw Collatz steps.
Square-Root Descent Has a Logarithmic Raw-Time Window
For natural-density-one many positive starts N, a raw Collatz iterate falls strictly below √N after more than log N / (2 log 2) steps and by at most 436 · log N steps.
Build from the checked theorem
Use the checked natural-density theorem to study stronger time constants, explicit density-convergence rates for restricted threshold classes, or other dynamical systems where arithmetic phase gaps feed quantitative mixing and descent.
Each ZIP contains the checked first-party Lean import closure, exact statements and boundaries, license, notice, evidence, source-footprint manifest, and an agent continuation file. Mathlib and other third-party dependencies are not bundled.
What is checked—and how much source supports it
How counting works: Line counts exclude blank lines; comments and documentation count. The total is the deduplicated, commit-pinned first-party Lean import closure; Mathlib and other third-party dependencies are excluded. Declaration count means names covered by the artifact's recorded evidence; it is not a count of every declaration in the source. Source footprint is not a difficulty or proof-quality score.
Independent publication review
The formal theorem's publication gates are accepted
Lean checks the proof. Independent AI review separately accepted evidence completeness, statement alignment, result boundary, and the retained theorem wording. Those gates apply to the formal result; generated media is reviewed and promoted separately. Neither review replaces Lean's proof check or broadens the theorem.
Independent review accepted the recorded build, exact declarations, unfinished-step scan, and axiom evidence.
The formal declaration was accepted against the named theorem and its exact variant.
The accepted boundary keeps nearby stronger or commonly confused claims out of scope.
Independent review accepted the retained theorem explanation and source presentation. Generated media follows a separate review and promotion gate.
The first-party source link is pinned to the checked package commit and exact Lean file.
A validated accepted-result record binds the four reviews to the checked formalization.
