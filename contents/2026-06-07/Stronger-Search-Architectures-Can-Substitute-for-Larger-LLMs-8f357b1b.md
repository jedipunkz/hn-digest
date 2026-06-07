---
source: "https://ttanv.github.io/levi/"
hn_url: "https://news.ycombinator.com/item?id=48437993"
title: "Stronger Search Architectures Can Substitute for Larger LLMs"
article_title: "LEVI: Stronger Search Architectures Can Substitute for Larger LLMs | Temoor Tanveer"
author: "ttanv"
captured_at: "2026-06-07T21:31:20Z"
capture_tool: "hn-digest"
hn_id: 48437993
score: 2
comments: 1
posted_at: "2026-06-07T19:59:52Z"
tags:
  - hacker-news
  - translated
---

# Stronger Search Architectures Can Substitute for Larger LLMs

- HN: [48437993](https://news.ycombinator.com/item?id=48437993)
- Source: [ttanv.github.io](https://ttanv.github.io/levi/)
- Score: 2
- Comments: 1
- Posted: 2026-06-07T19:59:52Z

## Translation

タイトル: より強力な検索アーキテクチャで大規模な LLM を代替できる
記事のタイトル: LEVI: より強力な検索アーキテクチャが大規模な LLM の代替となる |テムール・タンヴィール

記事本文:
TT >
について
研究
出版物
プロジェクト
教える
← ホームへ戻る
LEVI: より強力な検索アーキテクチャが大規模な LLM の代わりに使用できる
コードとプロンプト最適化のためのハーネスファーストの進化的フレームワーク。 3.3 ～ 6.7 倍の低コストで、GEPA、OpenEvolve、ShinkaEvolve、AdaEvolve、および EvoX のフロンティア モデルの実行よりも高いスコア。
ほとんどの LLM 誘導進化システムは、問題をフロンティア モデルに指摘し、何百、何千もの高価な呼び出しを実行するという高価な方法で結果を取得します。 LEVI は、より強力な検索アーキテクチャが大規模なモデルの代わりになり、コストを大幅に削減できるという逆の賭けをしています。アーカイブにモデルではなく多様なソリューションが保存されるようにハーネスを修正し、実際にジョブに適合するモデルに突然変異をルーティングし、冗長なサンプルの再スコアリングを停止します。そうすれば、予算の数分の一で小規模なオープンソース モデルからでも優れた結果が得られます。
これは、次の 2 つのまったく異なる設定に当てはまります。
コードの最適化。システム研究ベンチマークでは、LEVI は 3.3 ～ 6.7 倍の低コストで、7 つの問題のうち 6 つに関して、GEPA、OpenEvolve、ShinkaEvolve、AdaEvolve、および EvoX の公表されているフロンティア予算で最も優れた実行を上回っています。
迅速な最適化。 4 つの GEPA スイート ベンチマークでは、LEVI は展開の半分未満を使用して GEPA と同等かそれを上回っています。
LEVI は、予算の一部を使用して、より高い最終パフォーマンスを達成します。左: トランザクション スケジューリング コードの最適化では、LEVI は最初の ~50 回の評価内ですべてのベースラインの最終スコアを上回りました (サンプル効率約 15 倍)。右: Qwen3-8B を使用した HotpotQA プロンプト最適化では、LEVI は半分未満のロールアウト (~2.75K 対 ~6.87K) で GEPA を上回りました。
これがモデルの問題ではなくハーネスの問題である理由
フロンティアモデルの依存関係は主に、既存のフレームワークが検索を割り当てる方法に起因するものであり、

基本的な要件です。
LLM 誘導進化では、言語モデルと検索ループを組み合わせます。ユーザーが問題とスコアリング機能を提供し、LLM が候補の解決策を変更し、評価者がそれらをスコアリングし、解決策のデータベースが有望な解決策を保持します。このパラダイムは FunSearch によって導入されました 1 1 FunSearch (Romera-Paredes et al., 2024) は、フロンティアスケールのモデルを使用せずに、キャップセットの改善を含む新しい数学的結果を生み出しました。 2 2 AlphaEvolve (Novikov et al., 2025) は、ループをより強力な LLM とより大きなコードベースに拡張しました。そして現在では、数学、コード最適化、ヒューリスティック設計、システム研究、プロンプト最適化にまで及びます。
問題はコストです。 1 回の実行で高価なフロンティア モデルへの呼び出しが何千回も必要になる場合があり、報告されたシステム リサーチの実行には、GPT-5 または Gemini 3.0 Pro での問題ごとに 15 ～ 30 ドルかかることがよくあります。その価格により参入障壁が高くなり、反復が遅くなります。しかし、それは明らかにパラダイムに固有のものではありません。 FunSearch はフロンティア モデルをまったく使用せずに結果を生み出しました。私たちは、そのコストは、検索アーキテクチャへの投資が過少で、大規模なモデルに過度に依存することに起因すると主張しており、それは 3 つの別々の軸に沿って表示されます。
評価ごと。アーカイブが多様性を維持できない場合、検索は 1 つの盆地に崩壊し、強力なモデルに依存して逃げ道を生み出します。既存のフレームワークは、アイランド、埋め込みベースのノベルティ フィルター、または LLM ジャッジを使用して事後的にこれにパッチを適用し、それぞれが収束を防ぐのではなく補償します。
1ドルあたり。ミューテーション呼び出しは均一に扱われるため、小規模なモデルで処理できるローカル編集であっても、フロンティア モデルの価格が支払われます。
ロールアウトごと。すべての候補者は完全な検証セットで再スコアリングされ、冗長な例にロールアウトを費やしますが、これはプロンプトで特に苦痛です。

最適化。
LEVI は 3 つすべてを修正します。強力なモデルを想定してハーネスを構築するのではなく、予算が限られている場合に検索アーキテクチャがどうあるべきかを考えます。
コスト軸ごとに 1 つずつの 3 つのコンポーネント: 多様性を保持するソリューション データベース、役割を認識したミューテーション ルーター、およびランクを保持するプロキシ ベンチマーク。
LEVI は、1 つのブートストラップ シード パスを使用して、ソリューション ファミリを初期化し、CVT-MAP-Elites データベースを調整し、プロキシ ベンチマークを構築します。検索中、アーカイブは非同期の突然変異評価ループにフィードを送ります。ほとんどの突然変異は局所的な改良に小さな LLM を使用しますが、定期的なパラダイム シフトでは構造的に新しい候補を提案するためにより強力な LLM が使用されます。
LEVI は、非同期 AlphaEvolve スタイルのループに従います。ソリューション データベースは評価された候補を保存し、サンプラーはそこから親を抽出し、ルーターは適切なモデルに突然変異リクエストを送信し、結果の候補は挿入される前に並行して評価されます。親は、ワーカー固有の温度で並列プール全体の探索と活用のバランスをとり、ソフトマックス オーバー スコアでサンプリングされます。以下の 3 つの部分は、それぞれの拡張として読むのが最適です。アーカイブは原則に基づいたルーティングを可能にする構造を提供し、原則に基づいたルーティングは、限られた予算の下で多様性を保持するアーカイブを実用化するものです。
多様性を早期に確立し、維持するソリューションデータベース
多様なシーディングと柔軟な動作次元を備えた CVT-MAP-Elites アーカイブ。
有用なアーカイブは、多様性を早期に確立し、それを全体を通して維持する必要があります。ほとんどのフレームワークは単一のシード プログラムから始まります。これは、初期の探索では 1 つの誘引領域から逃れるために予算が費やされ、その後、アドホックなメカニズムを通じて多様性を回復する必要があることを意味します。 LEVI がスマを作る

代わりに先行投資を行います。初期化中に、LLM に構造的に多様なシードのシーケンスを要求します。各シードは以前の試行に条件付けされ、それらとは異なるように明示的に指示され、次のシードが行き止まりを回避できるように失敗がフィードバックされます。弱いが特徴的なシードは意図的に保持されます。シードは観察される記述子の範囲を拡大し、局所的な突然変異がめったに到達しない領域に足場を提供します。
同じシード パスでアーカイブを調整します。 LEVI は CVT-MAP-Elites アーカイブを使用しており、3 3 CVT-MAP-Elites (Vassiliades et al., 2017) は通常の MAP-Elites グリッドを重心ボロノイ テッセレーションに置き換え、指数関数的なセルの増加なしに高次元の記述子空間にスケーリングします。各候補を動作記述子にマッピングし、それをボロノイ テッセレーション内の最も近い重心に割り当て、セルごとに最高スコアの候補のみを保持します。記述子の値はオンラインで正規化されるため (ウェルフォード Z スコアとシグモイド)、高分散次元がジオメトリを支配することはなく、新しい動作レジームが現れると実行統計が更新されます。
重要なのは、LEVI が 2 つの記述子ファミリーをサポートし、あらゆる組み合わせを受け入れることです。入力側の記述子はソリューション自体 (コード長、循環複雑度などの AST 機能、ループと演算子の数) から取得され、出力がスパースな場合に役立ちます。出力側記述子は実行動作 (ランタイム、インスタンスごとのスコア プロファイル) から取得され、迅速な最適化にとって重要となるサンプル間のトレードオフを明らかにします。既存のフレームワークは 1 つのファミリーに限定されています。 OpenEvolve と ShinkaEvolve は入力側の機能に、GEPA はインスタンスごとのパレート フロントを通じて暗黙的に出力側の機能に機能します。 LEVI は問題に合った組み合わせを選択します。
実行を支配するローカル編集用の安価なモデル。まれな構造ジャンプ用に予約された強力なモデル。

突然変異の呼び出しは外から見ると均一に見えますが、その働きは大きく異なります。ほとんどはローカルな改良、定数の微調整、データ構造の交換、ヘルパーのインライン化です。いくつかは、アルゴリズムの形状を変更する構造的な書き換えです。既存のフレームワークは、両方に対してフロンティアモデルの価格を支払い、突然変異の品質が全体的にモデルのサイズに応じて変化することを暗黙的に想定しています。私たちは、この 2 つの操作には異なる要件があると主張します。ローカルなリファインメントは速度と量に報酬を与えますが、小さいモデルで十分です。構造的な書き換えは、スケールに伴う事前の範囲の広さに報います。
LEVI はこれを 2 つのルートに分割します。改良ルートでは、ミューテーション呼び出しの約 90% が Qwen3-30B-A3B などの小規模モデルに送信され、アーカイブから親が抽出され、単一の対象を絞った編集が要求されます。小さなモデルは共通のパターンに崩壊する可能性がありますが、アーカイブはこれを吸収します。重複に近い出力は占有されたセルにマップされて破棄されるため、崩壊の圧力はアーカイブのドリフトではなく無害な冗長性になります。パラダイムシフトルートは、より大きなモデルがコストを稼ぐ場所です。 LEVI は、1 つの強力な親を編集するよう依頼するのではなく、明確に分離された個別のアーカイブ セルから高得点の代表者をサンプリングし、それらをまとめて渡し、それらのファミリーのいずれにも属さない候補者を求めます。このモデルは、単一の親ではなく集団の構造に基づいて動作し、調査対象の家族のフロンティアを拡張します。これらの呼び出しは一定の間隔で、停滞時に発生します。
これらのルートは補完的です。洗練はアーカイブがすでに占有しているセルを深め、パラダイム シフトは占有する価値のあるセルのセットを拡張します。最適な分割は問題に依存し、アーカイブは両方を結合して、小規模なモデルにそれ自体の冗長性に対する堅牢性を与え、大規模なモデルに単一親 p の構造ビューを与えます。

romptはできません。
ランクを維持するプロキシ ベンチマーク
評価コストが支配的な場合は、すべての個別のスコアではなく、候補のランキングを維持するために選択された小さなサブセットにスコアを付けます。
3 番目の軸は、候補者の評価自体にコストがかかる場合に表示されます。コードの最適化では、通常、主なコストはミューテーションの呼び出しまたはプログラムの実行です。プロンプトの最適化では、各候補プロンプトで検証セット全体にわたる多くのロールアウトが必要になる場合があり、以前の作業で予算の大部分がそこに費やされます。 LEVI は、進化的選択は正確なスコアではなく、より悪い候補よりもより良い候補を選択することだけに依存するため、完全な検出セットのサブセットであるより小さなプロキシ ベンチマークを構築します。これは、完全なセットの選択シグナルを保存します。
プロキシは同じ初期化パス中に構築されます。各多様なシードは検出セット内のすべての例でスコア付けされ、候補の品質が例間でどのように変化するかについてのキャリブレーション マトリックスを生成します。次に、LEVI は、ランクの忠実度 (サブセットの誘導された候補ランキングが完全なセットと一致するか)、分離 (サンプルが候補間を区別しているか)、および冗長ペナルティ (このサンプルのスコア列が選択されたものとすでに相関しているか) の 3 つの項を最大化するサブセットを貪欲に選択します。 24 プロンプト × 150 問題の行列では、この貪欲な列サブセット法は、すべての等コスト等高線で k メドイドおよびランダム サブセット + リッジ回帰を上回ります。
わずかなコストでのシステム研究ベンチマーク、管理された同一モデルの比較、各コンポーネントを分離するアブレーション、および迅速な最適化。
フロンティア予算ベースラインに対するシステム研究
ADRS 問題スイート 4 4 ADRS (Cheng et al., 2025) の 7 つのタスク、つまりネットワーキング、LLM サービス、データベース、分散システムにわたる現実世界のシステム問題を評価します。スパニングネットワーク

キング、LLM サービス、データベース、分散システム。ベースラインは、GEPA、5 5 GEPA (Agrawal et al.、2025)、OpenEvolve、および ShinkaEvolve (Lange et al.、2025) からの最も強力な公開結果です。ベースライン実行では、GPT-5 または Gemini 3.0 Pro を問題ごとに 15 ～ 30 ドルで使用します。 OpenEvolve、ShinkaEvolve、AdaEvolve、および EvoX は、それぞれ高価な予算と SOTA モデルを組み合わせています。その代わり、LEVI はほとんどの問題に 5 ドル未満を費やし、変異の 90% 以上を Qwen3-30B と MiMo-v2-Flash にルーティングし、残りはパラダイム シフトのために Gemini 3.0 フラッシュに送られます。
LEVI は、7 つのネイティブ メトリクスのうち 6 つについて、公開されている最良の結果を改善しています。唯一の例外は Prism です。Prism では、すべてのフロンティア予算手法が 0.14 ポイント以内に集まります。 LEVI は 6 つの問題に対して 4.50 ドルを費やし、3.3 ～ 6.7 倍のコスト削減になります。トランザクション スケジューリングでは、延長された 12.50 ドルの実行額は、フロンティア モデルの予算である 20 ドルを依然として下回っていますが、以前の最高額を大幅に上回っています。 LLM-SQL では、LEVI はわずか 0.50 ドルを費やして主要なフレームワークに合格しました。これは 35 分のコスト削減です。
管理された比較: 同じモデル、同じ予算
上の表は、モデルの選択、予算、アーキテクチャを組み合わせたものです。検索アーキテクチャを分離するために、LEVI、OpenEvolve、ShinkaEvolve、GEPA を同一の条件で実行します。

[切り捨てられた]

## Original Extract

TT ">
About
Research
Publications
Projects
Teaching
← Back to home
LEVI: Stronger Search Architectures Can Substitute for Larger LLMs
A harness-first evolutionary framework for code and prompt optimization. Better scores than frontier-model runs of GEPA, OpenEvolve, ShinkaEvolve, AdaEvolve, and EvoX, at 3.3–6.7× lower cost.
Most LLM-guided evolutionary systems get their results the expensive way: by pointing a frontier model at the problem and burning through hundreds or thousands of expensive calls. LEVI takes the opposite bet, that a stronger search architecture can substitute for a larger model and drastically reduce the cost. Fix the harness so the archive preserves diverse solutions instead of the model, route mutations to the model that actually fits the job, and stop re-scoring redundant examples, and strong results follow even from small open-source models at a fraction of the budget.
This holds across two very different settings:
Code optimization. On systems-research benchmarks, LEVI beats the best published frontier-budget runs of GEPA, OpenEvolve, ShinkaEvolve, AdaEvolve, and EvoX on six of seven problems, at 3.3–6.7× lower cost.
Prompt optimization. On four GEPA-suite benchmarks, LEVI matches or exceeds GEPA using less than half the rollouts.
LEVI reaches higher final performance using a fraction of the budget. Left: on Transaction Scheduling code optimization, LEVI exceeds every baseline's final score within the first ~50 evaluations (≈15× sample efficiency). Right: on HotpotQA prompt optimization with Qwen3-8B, LEVI outperforms GEPA with fewer than half the rollouts (~2.75K vs ~6.87K).
Why this is a harness problem, not a model problem
Frontier-model dependence is largely an artifact of how existing frameworks allocate search, not a fundamental requirement.
LLM-guided evolution pairs a language model with a search loop: the user provides a problem and a scoring function, the LLM mutates candidate solutions, an evaluator scores them, and a solution database keeps the promising ones around. The paradigm was introduced by FunSearch 1 1 FunSearch (Romera-Paredes et al., 2024) produced novel mathematical results, including the cap-set improvement, without a frontier-scale model. and scaled up by AlphaEvolve, 2 2 AlphaEvolve (Novikov et al., 2025) extended the loop to stronger LLMs and larger codebases. and it now spans math, code optimization, heuristic design, systems research, and prompt optimization.
The catch is cost. A single run can require thousands of calls to expensive frontier models, and reported systems-research runs often cost $15–30 per problem on GPT-5 or Gemini 3.0 Pro. That price raises the barrier to entry and slows iteration. But it is not obviously inherent to the paradigm; FunSearch produced its results without a frontier model at all. We argue the cost comes from over-relying on larger models while under-investing in the search architecture, and it shows up along three separate axes:
Per-evaluation. When the archive fails to preserve diversity, the search collapses into a single basin and leans on a strong model to generate escapes. Existing frameworks patch this after the fact with islands, embedding-based novelty filters, or LLM judges, each compensating for convergence rather than preventing it.
Per-dollar. Mutation calls are treated uniformly, so frontier-model prices get paid even for local edits a small model could handle.
Per-rollout. Every candidate is re-scored on the full validation set, spending rollouts on redundant examples, which is especially painful in prompt optimization.
LEVI fixes all three. Rather than building the harness around the assumption of a strong model, we ask what the search architecture should look like when the budget is limited.
Three components, one per cost axis: a diversity-preserving solution database, a role-aware mutation router, and a rank-preserving proxy benchmark.
LEVI uses one bootstrapped seed pass to initialize solution families, calibrate the CVT-MAP-Elites database, and build the proxy benchmark. During search, the archive feeds an asynchronous mutation–evaluation loop: most mutations use a small LLM for local refinement, while periodic paradigm shifts use a stronger LLM to propose structurally new candidates.
LEVI follows an asynchronous AlphaEvolve-style loop. The solution database stores evaluated candidates, samplers draw parents from it, the router sends mutation requests to the appropriate model, and resulting candidates are evaluated in parallel before being inserted back. Parents are sampled with a softmax over scores, with worker-specific temperatures balancing exploration and exploitation across the parallel pool. The three pieces below are best read as extensions of each other: the archive provides the structure that makes principled routing possible, and principled routing is what makes a diversity-preserving archive practical under a tight budget.
A solution database that establishes diversity early and maintains it
Diverse seeding plus a CVT-MAP-Elites archive with flexible behavioral dimensions.
A useful archive must establish diversity early and preserve it throughout. Most frameworks begin from a single seed program, which means the early search spends budget escaping one basin of attraction, and later diversity has to be recovered through ad-hoc mechanisms. LEVI makes a small upfront investment instead. During initialization it asks the LLM for a sequence of structurally diverse seeds, each conditioned on previous attempts and explicitly instructed to differ from them, with failures fed back so the next seed avoids dead ends. Weak-but-distinct seeds are kept on purpose: they expand the observed descriptor range and provide footholds in regions local mutation rarely reaches.
The same seed pass calibrates the archive. LEVI uses a CVT-MAP-Elites archive, 3 3 CVT-MAP-Elites (Vassiliades et al., 2017) replaces the regular MAP-Elites grid with a centroidal Voronoi tessellation, scaling to higher-dimensional descriptor spaces without exponential cell growth. mapping each candidate to a behavioral descriptor, assigning it to the nearest centroid in a Voronoi tessellation, and keeping only the highest-scoring candidate per cell. Descriptor values are normalized online (Welford z-score plus a sigmoid), so high-variance dimensions do not dominate the geometry, and the running statistics update as new behavioral regimes appear.
Crucially, LEVI supports two descriptor families and accepts any mix. Input-side descriptors come from the solution itself (code length, AST features like cyclomatic complexity, loop and operator counts) and are useful when outputs are sparse. Output-side descriptors come from execution behavior (runtime, per-instance score profiles) and expose trade-offs between examples, which matters for prompt optimization. Existing frameworks commit to one family; OpenEvolve and ShinkaEvolve to input-side features, GEPA implicitly to output-side ones through its per-instance Pareto front. LEVI chooses the mix that fits the problem.
Cheap models for the local edits that dominate the run; a stronger model reserved for rare structural jumps.
Mutation calls look uniform from the outside, but the work varies enormously. Most are local refinements, tweak a constant, swap a data structure, inline a helper. A few are structural rewrites that change the algorithm’s shape. Existing frameworks pay frontier-model prices for both, implicitly assuming mutation quality scales with model size across the board. We argue the two operations have different requirements: local refinement rewards speed and volume, where smaller models suffice; structural rewrites reward the broader prior coverage that comes with scale.
LEVI splits this into two routes. The refinement route sends roughly 90% of mutation calls to a small model such as Qwen3-30B-A3B, drawing parents from the archive and asking for a single targeted edit. Small models can collapse onto common patterns, but the archive absorbs this: near-duplicate outputs map to occupied cells and are discarded, so collapse pressure becomes harmless redundancy rather than archive drift. The paradigm-shift route is where the larger model earns its cost. Instead of asking it to edit one strong parent, LEVI samples high-scoring representatives from distinct, well-separated archive cells and passes them together, asking for a candidate that belongs to none of those families. The model operates on the structure of the population, not a single parent, extending the frontier of explored families. These calls fire at fixed intervals and on stagnation.
The routes are complementary: refinement deepens the cells the archive already occupies, paradigm shifts expand the set of cells worth occupying. The optimal split is problem-dependent, and the archive ties both together, giving the small model robustness to its own redundancy and giving the large model a structural view a single-parent prompt cannot.
A rank-preserving proxy benchmark
When evaluation cost dominates, score on a small subset chosen to preserve candidate rankings, not every individual score.
The third axis appears when evaluating a candidate is itself expensive. In code optimization, the dominant cost is usually the mutation call or the program run. In prompt optimization, each candidate prompt may require many rollouts across a validation set, and prior work spends a large fraction of its budget there. LEVI builds a smaller proxy benchmark , a subset of the full discovery set, that preserves the selection signal of the full set, since evolutionary selection only depends on choosing better candidates over worse ones, not on exact scores.
The proxy is built during the same initialization pass: each diverse seed is scored on every example in the discovery set, producing a calibration matrix of how candidate quality varies across examples. LEVI then greedily selects a subset that maximizes three terms, rank faithfulness (does the subset’s induced candidate ranking agree with the full set), separation (does the example discriminate between candidates), and a redundancy penalty (is this example’s score column already correlated with chosen ones). On a 24-prompt × 150-problem matrix this greedy column-subset method beats k -medoids and random-subset + ridge regression at every iso-cost contour.
Systems-research benchmarks at a fraction of the cost, a controlled same-model comparison, ablations isolating each component, and prompt optimization.
Systems research, against frontier-budget baselines
We evaluate on seven tasks from the ADRS problem suite 4 4 ADRS (Cheng et al., 2025): real-world systems problems spanning networking, LLM serving, databases, and distributed systems. spanning networking, LLM serving, databases, and distributed systems. Baselines are the strongest published results from GEPA, 5 5 GEPA (Agrawal et al., 2025), OpenEvolve , and ShinkaEvolve (Lange et al., 2025). Baseline runs use GPT-5 or Gemini 3.0 Pro at $15–30 per problem. OpenEvolve, ShinkaEvolve, AdaEvolve, and EvoX, each coupled with expensive budgets and SOTA models. LEVI instead spends under $5 on most problems, routing more than 90% of mutations to a Qwen3-30B and a MiMo-v2-Flash, with the rest going to Gemini 3.0 Flash for paradigm shifts.
LEVI improves on the best published result on six of seven native metrics. The lone exception is Prism, where every frontier-budget method clusters within 0.14 points. On six problems LEVI spends $4.50, a 3.3–6.7× cost reduction; on Transaction Scheduling, an extended $12.50 run still sits below the $20 frontier-model budget while beating the prior best by a wide margin. On LLM-SQL, LEVI passes the leading framework after spending only $0.50, a 35× cost reduction.
Controlled comparison: same model, same budget
The table above mixes model choice, budget, and architecture. To isolate the search architecture, we run LEVI, OpenEvolve, ShinkaEvolve, and GEPA under identical conditions: a sin

[truncated]
