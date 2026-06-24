---
source: "https://ucbskyadrs.github.io/blog/levi/"
hn_url: "https://news.ycombinator.com/item?id=48659040"
title: "Systems optimization should be part of CI/CD"
article_title: "LEVI: Better ADRS Results at a Fraction of the Cost | ADRS — AI-Driven Research for Systems"
author: "ttanv"
captured_at: "2026-06-24T13:44:04Z"
capture_tool: "hn-digest"
hn_id: 48659040
score: 2
comments: 0
posted_at: "2026-06-24T12:57:31Z"
tags:
  - hacker-news
  - translated
---

# Systems optimization should be part of CI/CD

- HN: [48659040](https://news.ycombinator.com/item?id=48659040)
- Source: [ucbskyadrs.github.io](https://ucbskyadrs.github.io/blog/levi/)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T12:57:31Z

## Translation

タイトル: システムの最適化は CI/CD の一部である必要があります
記事のタイトル: LEVI: わずかなコストでより優れた ADRS 結果 | ADRS — AI を活用したシステム研究
説明: OpenEvolve や GEPA などのアルゴリズム検出フレームワークは、AI-Driven Research for Systems (ADRS) が強力なアルゴリズムを生成できることを示しました。しかし、今日のフレームワークは、ADRS が次に進むべき方向としては高価すぎます。 LEVI は、アルゴリズムによる発見のコストを削減することを中心に構築されたフレームワークです。
[切り捨てられた]

記事本文:
ADRS ホーム ブログ リーダーボード すべての投稿について AI システム ケース スタディ LEVI: わずかなコストでより優れた ADRS 結果
この投稿は、AI を使用して現実世界のシステムの問題に対してより適切なアルゴリズムを自動的に発見する AI-Driven Research for Systems (ADRS) ケース スタディ シリーズの一部です。
OpenEvolve や GEPA などのアルゴリズム検出フレームワークは、AI-Driven Research for Systems (ADRS) が強力なアルゴリズムを生成できることを示しています。しかし、今日のフレームワークは、ADRS が次に進むべき方向としては高価すぎます。将来は、1 回限りのベンチマーク結果ではなく、継続的かつオーダーメイドの最適化が求められます。システムは、展開の正確なワークロード、ハードウェア、および SLO に合わせたソリューションを生成する必要があります。これらの変化に合わせて適応します。あらゆる最適化に多額の費用がかかる場合、これは不可能です。
LEVI は、アルゴリズムによる発見のコストを削減することを中心に構築されたフレームワークです。すべてのステップで最も強力で最も高価なモデルを使用する代わりに、探索ハーネスに投資します。より小型で安価なモデル (例: QWEN 30B) がほとんどの突然変異を実行し、より大きなモデルはまれなパラダイム シフトのために確保されます。これが可能になるのは、LEVI がコード構造 (ループ数など) と実際の動作 (サブセット x でのパフォーマンスなど) の両方にわたって多様性を維持し、検索アーカイブが単一のソリューション ファミリに崩壊しないようにしているためです。
その結果、主要なベンチマーク比較のベースラインよりもおよそ 3 ～ 7 倍安いコストで、より強力な ADRS 結果が得られるフレームワークが誕生しました。
✍️ 以前の ADRS ブログ: https://ucbskyadrs.github.io/
👩‍💻 コード: github.com/ttanv/levi
💬 参加してください: join.slack.com/t/adrs-global および Discord
🌎フォローしてください:x.com/ai4research_ucb
LLM アルゴリズム検出フレームワークは、ADRS に強力な結果をもたらすことが期待されています。ただし、重要なボトルネック

k 残り: コスト 。このブログでは、コストが ADRS において重要な役割を果たす理由を論じ、次に LEVI を紹介します。わずかなコストで他のものより優れたパフォーマンスを発揮するアルゴリズム検出フレームワーク。
既存のフレームワークでは、高価で大規模なクローズドソース LLM を使用する必要があります。これには明らかな理由から問題があります。たとえば、ほとんどの研究者にはそのような実験を行う余裕がないため、参入障壁が高くなります。しかし、より重要な問題はより広範であり、ADRS を 1 回実行して 1 つの強力な結果を生み出すものとして見るべきではありません。
ADRS にとって、コストが桁違いに低くなるまでコストを削減することは、自然な次のステップとなるはずです。これは、ADRS の結果を、通常のシステム論文と同様に 1 回限りの研究成果として扱うべきではないためです。特定の問題に対して、研究者はアルゴリズムとヒューリスティックを改良して、より良い結果を生み出します。その後、業界はこれらのアルゴリズムを移植し、設定に適応させることでこれに続きます。
代わりに、完全にオーダーメイドのソリューションに移行する必要があります。各ソリューションは、各展開の正確なセットアップと環境に合わせて調整され、最大限の効果を絞り出します。
論理的な結論から言えば、ADRS は CI/CD のより洗練された形式と見なされるべきです。ユーザーがスコアリング関数とデプロイメント設定を定義すると、単にリンターやフォーマッタが自動的にスタイルやフォーマットを修正するのではなく、アルゴリズム自体が自動的に最適化されます。リソース (新しい GPU など) または優先順位 (異なる SLO) が変化すると、対応するアルゴリズムが自動的に最適化されます。現在、マルチリージョンのクラウド スケジューラを実行している企業は、他の企業と同じアルゴリズムを使用しています。安価な ADRS を使用すると、実際のトラフィック パターン、実際の SLO、実際のトラフィック パターンに基づいて毎晩再最適化できます。

ハードウェアミックス。
LEVI の紹介: わずかなコストでの LLM ベースの最適化
上記を踏まえて、このブログでは LEVI を紹介します。LEVI は、ADRS 問題に対してわずかなコストで SOTA パフォーマンスを生み出す LLM ベースの進化的なフレームワークです。これは、あまりにも多くのフレームワークが最大の SOTA モデルへのアクセスを前提としており、そのモデルを中心にハーネスを構築しているという重要な洞察に基づいて構築されています。
重要な洞察: モデルではなくハーネスに投資する
最大のモデルへのアクセスがデフォルトではないことを前提としています。実際、FunSearch の元の論文では、大規模なモデルからはメリットが得られず、AlphaEvolve でのみ成功したと報告されています。オープンソース コミュニティはこれを見逃していることが多く、あらゆる段階で最強のモデルを投入しています。 LEVI は、層別モデル割り当てとダイバーシティ維持の改善という 2 つの主要なコンポーネントを通じて、代わりにハーネス ファーストのアプローチを採用します。
フロンティアモデルは役立ちますが、すべての突然変異に使用すると無駄になります。小規模な LLM は、生成されるソリューションの膨大な量が大規模なモデルの品質上の利点を上回る可能性があるため、予算が限られている場合には実際に好まれる場合があります。ただし、小規模なモデルでは事前トレーニングの分布が狭くなり、根本的に異なるアプローチを提案できるアイデアや能力の範囲が制限されます。どちらのモデル クラスが厳密には優れているというわけではありません。彼らは異なる強みを持っているだけです。
一部の既存のフレームワークはすでに複数のモデルをサポートしていますが、それらを交換可能なものとして扱い、アンサンブルから均一にサンプリングしたり、突然変異が実際に要求するものに関係なく呼び出しをルーティングしたりします。これは自然な非対称性を無視しています。まったく新しいアルゴリズムの方向性を提案するには幅広い知識と創造的な推論が必要ですが、既存のアプローチを改良する（定数の調整、演算の並べ替え、エッジケースの調整）にははるかに少ない知識が必要です。ハーネス翔

この違いを認識し、それに応じて割り当ててください。
LEVI は、モデルの能力をタスクの需要に合わせる層別モデル割り当てを導入しています。小規模で安価なモデルは、確立されたアルゴリズム ファミリ内での局所的な改良と漸進的な改善など、検索の大部分を処理します。より大きなモデルは、まれなパラダイム シフト、つまり既存のアプローチを洗練するのではなく、構造的に異なるアプローチを提案することを目的とした突然変異のために予約されています。原則は簡単です。各モデルをその強みに合わせて割り当てます。小型モデルは幅広さとスループットを実現し、大型モデルは創造的な飛躍を実現します。
ただし、これには 2 つの疑問が生じます。まず、より大きなモデルにパラダイム シフトに意味のあるコンテキストを与えるために、各アルゴリズム ファミリから代表的なソリューションをどのように選択するのでしょうか?第 2 に、現在はより小規模なモデルとその出力量に大きく依存しているため、アーカイブの収束を防ぐためのより堅牢なメカニズムが必要です。
LEVI は、モデルの能力をタスクの需要に合わせます。洗練には安価なモデル (ローカル QWEN 30B など)、パラダイム シフトには高価なモデルを使用します。
多様性の維持の向上
構造的および行動的多様性の統合。既存のフレームワークがフロンティア モデルを必要とする理由はあまり明らかではありませんが、それらのモデルが二重の役割を果たしていることです。より大きな出力空間は暗黙のうちに多様性を維持します。GPT-5 または Claude Opus は、アーカイブ自体に収束を防ぐ強力なメカニズムがないという事実を無視して、30B モデルよりも自然に幅広いソリューションを生成します。多様性が崩壊すると、さらに多くの LLM 呼び出しを使用した拒否サンプリングから埋め込みモデルの使用に至るまで、その対応はさらに複雑さを増すことになります。これらは脆弱な基盤を補うものであり、根本的な問題の解決策ではありません。
根本的な問題は、既存のフレームワークがメインであることです。

多様性は 1 つの軸のみに沿って保持され、その軸は狭いです。 OpenEvolve は、コードの長さなどの構造的特徴を考慮します。 GEPA は、パレート フロントを通じてインスタンスごとのパフォーマンスのトレードオフを考慮します (実際には、前のメカニズムよりも強力であることがよくあります)。どちらも現実のものを捉えていますが、どちらも全体像を捉えていません。構造だけでは動作の違いが見逃されます。ループ数が異なる 2 つのプログラムは問題を同じように解決する可能性があります。また、インスタンスごとのスコアだけでは、個々のインスタンスでは同様に機能するものの、根本的に異なる方法で動作するソリューションは見逃されます。
LEVI では、1 つの軸を選択するのではなく、両方を 1 つの動作記述子の次元として使用します。各ソリューションはフィンガープリントにマッピングされます。これは、コード構造の特徴 (コード長などの単純な次元を超えて、ループ数、循環的複雑さなどの尺度まで) とインスタンスごとの動作結果を組み合わせたベクトルであり、すべて正規化されて [0, 1] に投影されます。ここでもフレームワークは柔軟です。デフォルトが問題に合わない場合、ユーザーは独自のディメンションを定義できます。
このフィンガープリントは CVT-MAP-Elites アーカイブに保存されており、結合された空間上のボロノイ テッセレーションは、どちらの軸も単独では提供しない幾何学的構造を維持します。アーカイブには、さまざまな次元に沿ったさまざまな価値を持つ、さまざまなソリューションのセットが保管されています。これは、前のセクションの最初の質問にも直接答えます。ボロノイ領域は自然にソリューションをアルゴリズム ファミリにクラスター化し、パラダイム シフトの代表的なソリューションを提供します。
アーカイブを初期化しています。従来の CVT-MAP-Elites は、記述子空間全体で重心を均一に初期化します。私たちが使用する高次元 (6 ～ 10 ディム) では、ほとんどの領域が決して訪問されない、非常にまばらなテッセレーションになります。リーバイスの広告

データ主導型のアプローチを選択します。連続生成を通じて (スコアに関係なく) 意図的に一意のアプローチのセットを作成し、それらを使用して初期重心を作成します。これにより、アーカイブが異なることがわかっているソリューションに基づいていることが保証されます。
LEVI は、構造と動作の両方にわたって共有フィンガープリント空間を通じて多様性を維持しているため、アーカイブ自体は、ますます強力なモデルや補助ヒューリスティックに大きく依存するのではなく、多様性の負荷をより多く担っています。
LEVI 入門 (Python API)
以下は、ダミー ビン パッキング問題を最適化する LEVI プログラムの例です。フレームワークの詳細はすべて抽象化されているため、ユーザーは問題の定義に集中できます。
輸入リヴァイ
def スコア_fn(パック):
ビン = パック([4, 8, 1, 4, 2, 1], 10)
無駄 = sum(10 - ビン内の b の sum(b))
return {"スコア": max(0.0, 100.0 - 無駄)}
結果 = levi.evolve_code(
「ビンの梱包を最適化して無駄なスペースを最小限に抑える」、
function_signature="def パック(アイテム, bin_capacity):",
スコア_fn=スコア_fn、
モデル = "openai/gpt-4o-mini",
予算ドル=2.0、
)
github.com/ttanv/levi で LEVI を自分で試してみてください。
LEVI は、改善が可能なすべての問題で最高スコアを達成し、次善のフレームワーク (GEPA) の平均 71.9 と比較して平均 76.5 で、従来の最新技術よりも +4.6 ポイント改善しました。 Cloudcast では、LEVI は完璧な 100.0 に達し、ベンチマークのスコアリング関数の下で問題が完全に解決されたことを示しています。最大の改善は LLM-SQL (+5.8)​​ とスポット マルチ (+5.7) に見られますが、スポット シングル (+0.3) とトランザクション スケジューリング (+1.1) のより控えめな改善は、より小さな意思決定空間またはより困難な最適化ランドスケープの問題を反映しています。 Prism はすべてのフレームワークにわたって 87.4 で同率を維持しており、現在の問題定式化では単一の d が許容されることが確認されています。

有力な解決策。
LEVI の階層化された配分は、コスト削減の主な推進力です。突然変異の大部分を軽量モデルを通じてルーティングすることにより、すべての呼び出しに GPT-5 または Gemini-3.0-Pro を使用するベースラインと比較して、世代ごとのコストがおよそ 1 桁下がります。これにより、LEVI は実質的により多くの世代を実行できるようになり、合計の支出は少なくなります。ベースラインの場合は 15 ～ 30 ドルであるのに対し、ほとんどのタスクでは問題ごとに 4.50 ドル (トランザクション スケジューリング: 13 ドル)。
コスト削減は、ハーネスファーストのアプローチが機能していることの証拠です。アーカイブが多様性を維持している場合、ほとんどの検索には安価なモデルで十分です。
制御されたアーキテクチャの比較
同じモデル、同じ予算、3 つのシード: 検索アーキテクチャの貢献を分離します。
主な結果では、モデルの選択、予算、アーキテクチャが同時に異なるフレームワークを比較しています。検索アーキテクチャの寄与を分離するために、LEVI、OpenEvolve、および GEPA を同じ条件下で実行しました。つまり、ローカルで提供される単一の Qwen3-30B-A3B モデル、750 回の成功した評価、および 2 つの代表的な問題に対する 3 つのランダム シードです。 OpenEvolve では、より小さいモデルの親の数を 5 から 2 に減らす必要がありましたが、それでも多くの障害が発生しました。エヴァ成功を報告します

[切り捨てられた]

## Original Extract

Algorithmic discovery frameworks like OpenEvolve and GEPA have shown that AI-Driven Research for Systems (ADRS) can produce strong algorithms. But today's frameworks are too expensive for where ADRS should go next. LEVI is a framework built around lowering the cost of algorithmic discovery, getting
[truncated]

ADRS Home Blog Leaderboard About All posts AI Systems Case Study LEVI: Better ADRS Results at a Fraction of the Cost
This post is part of our AI-Driven Research for Systems (ADRS) case study series, where we use AI to automatically discover better algorithms for real-world systems problems.
Algorithmic discovery frameworks like OpenEvolve and GEPA have shown that AI-Driven Research for Systems (ADRS) can produce strong algorithms. But today's frameworks are too expensive for where ADRS should go next. The future is in continuous and bespoke optimization, not one-off benchmark results. The system should generate solutions tailored to the deployment's exact workload, hardware, and SLOs; adapting as these change. This is not possible when every optimization costs a fortune.
LEVI is a framework built around lowering the cost of algorithmic discovery. Instead of using the strongest, most expensive models for every step, it invests in the search harness: smaller, cheaper models (e.g. QWEN 30B) do most mutations, while larger models are reserved for rarer paradigm shifts. This is made possible because LEVI maintains diversity across both code structure (e.g. number of loops) and actual behavior (e.g. performance on subset x), ensuring the search archive does not collapse into a single solution family.
The result is a framework that gets stronger ADRS results at a fraction of the cost: roughly 3–7× cheaper than baselines in the main benchmark comparison.
✍️ Previous ADRS Blogs: https://ucbskyadrs.github.io/
👩‍💻 Code: github.com/ttanv/levi
💬 Join us : join.slack.com/t/adrs-global and Discord
🌎 Follow us : x.com/ai4research_ucb
LLM algorithmic discovery frameworks have shown promise in delivering strong results for ADRS. However, a key bottleneck remains: cost . This blog argues why cost plays an important role in ADRS, and then introduces LEVI; an algorithmic discovery framework that outperforms others at a fraction of the cost.
Existing frameworks require using expensive and large closed-source LLMs. This is problematic for obvious reasons. For example, it raises the barrier of entry since most researchers can't afford such experiments. But the more important issue is broader: ADRS should not be viewed as something we run once to produce a single strong result.
Decreasing the cost until it is orders of magnitude less should be the natural next step for ADRS. This is because the results from ADRS should not be treated as one-off research outcomes similar to how usual systems papers look like. Where for a given problem, researchers improve upon algorithms and heuristics to yield better results. Where industry then follows along by porting these algorithms and adapting them to their setup.
Instead we should move towards completely bespoke solutions. Where each solution is tailored to the exact setup and environment that each deployment has, squeezing out the most juice.
Taken to its logical conclusion, ADRS should be seen as a more sophisticated form of CI/CD. Where the user defines their scoring function and the deployment setup, and instead of just linters or formatters automatically fixing the style and formatting, the algorithm itself is automatically optimized. As the resources (e.g. new GPUs) or priorities change (different SLOs), the corresponding algorithms are automatically optimized. An enterprise running a multi-region cloud scheduler today uses the same algorithm as everyone else. With cheaper ADRS, they could re-optimize nightly against their actual traffic patterns, their actual SLOs, their actual hardware mix.
Introducing LEVI: LLM-Based Optimization at a Fraction of the Cost
Given the above, this blog introduces LEVI: an LLM-based evolutionary framework that produces SOTA performances on ADRS problems at a fraction of the cost. It is built on the key insight that too many frameworks assume access to the largest SOTA models, and build their harnesses around them.
Key Insight: Invest in the Harness instead of the Model
Assuming access to the largest models should not be the default. In fact, the original FunSearch paper reported being unable to benefit from larger models, and only with AlphaEvolve did they succeed. The open-source community often misses this, throwing the strongest models at every step. LEVI proceeds to take a harness first approach instead, through two key components: stratified model allocation and improved diversity maintenance.
Frontier models help, but they are a waste if used for every mutation. Smaller LLMs may actually be preferred under tight budgets, since the sheer quantity of solutions they produce can outweigh the quality advantage of larger models. However, smaller models have a narrower pretraining distribution, limiting their range of ideas and ability to propose fundamentally different approaches. Neither model class is strictly better; they just have different strengths.
Some existing frameworks already support multiple models, but treat them as interchangeable, sampling from an ensemble uniformly or routing calls without regard to what the mutation actually demands. This ignores a natural asymmetry: proposing an entirely new algorithmic direction requires broad knowledge and creative reasoning, while refining an existing approach (adjusting constants, reordering operations, tuning edge cases) requires far less. The harness should be aware of this distinction and allocate accordingly.
LEVI introduces stratified model allocation, which matches model capacity to task demand. Smaller, cheaper models handle the majority of the search: local refinements and incremental improvements within an established algorithmic family. Larger models are reserved for infrequent paradigm shifts : mutations that aim to propose structurally different approaches rather than polish existing ones. The principle is straightforward: allocate each model toward its strength. Small models for breadth and throughput, large models for creative leaps.
However, this raises two questions. First, how do we select representative solutions from each algorithmic family to give the larger model meaningful context for paradigm shifts? Second, since we now rely more heavily on smaller models and their volume of output, we need a more robust mechanism to prevent the archive from converging.
LEVI matches model capacity to task demand: cheap models (e.g. a local QWEN 30B) for refinement, expensive models for paradigm shifts.
Improved Diversity Maintenance
Unifying Structural and Behavioral Diversity. A less obvious reason existing frameworks require frontier models is that those models are doing double duty. Their larger output space implicitly maintains diversity: a GPT-5 or Claude Opus naturally produces a wider spread of solutions than a 30B model, ignoring the fact that the archive itself has no strong mechanism to prevent convergence. When diversity does collapse, the response has been to add complexity on top: ranging from rejection sampling using even more LLM calls to using embedding models. These are compensations for a weak foundation, not solutions to the underlying issue.
The underlying issue is that existing frameworks maintain diversity along only one axis, and a narrow one at that. OpenEvolve considers structural features like code length; GEPA considers per-instance performance trade-offs through Pareto fronts (in practice often more powerful than the former mechanism). Both capture something real, but neither captures the full picture. Structure alone misses behavioral differences: two programs with different loop counts might solve the problem identically. And per-instance scores alone miss solutions that perform similarly on individual instances but work in fundamentally different ways.
Rather than choosing one axis, LEVI uses both as dimensions of a single behavioral descriptor. Each solution is mapped to a fingerprint: a vector combining code-structural features (going beyond simple dimensions like code length to measures such as loop count, cyclomatic complexity) alongside per-instance behavioral results, all normalized and projected to [0, 1]. The framework is also flexible here: users can define their own dimensions when the defaults do not fit their problem.
This fingerprint lives in a CVT-MAP-Elites archive, where a Voronoi tessellation over the combined space maintains geometric structure that neither axis provides alone. The archive holds a diverse set of solutions with different values along the different dimensions. This also directly answers the first question from the previous section: the Voronoi regions naturally cluster solutions into algorithmic families, giving us representative solutions for paradigm shifts.
Archive Initializing. Traditional CVT-MAP-Elites initializes centroids uniformly across the descriptor space. With the higher dimensionality we use (6 to 10 dims), this leads to an extremely sparse tessellation where most regions will never be visited. LEVI adopts a data-driven approach; it creates a set of deliberately unique approaches (regardless of scores) through sequential generation, and then uses those to create the initial centroids. This ensures that the archive is based on solutions that are known to be different.
LEVI maintains diversity through a shared fingerprint space over both structure and behavior, so the archive itself carries more of the diversity burden instead of relying as heavily on ever-stronger models or auxiliary heuristics.
Getting Started with LEVI (Python API)
Below is an example LEVI program, optimizing a dummy bin packing problem. All of the framework details are abstracted away, and the user can focus on defining the problem.
import levi
def score_fn(pack):
bins = pack([4, 8, 1, 4, 2, 1], 10)
wasted = sum(10 - sum(b) for b in bins)
return {"score": max(0.0, 100.0 - wasted)}
result = levi.evolve_code(
"Optimize bin packing to minimize wasted space",
function_signature="def pack(items, bin_capacity):",
score_fn=score_fn,
model="openai/gpt-4o-mini",
budget_dollars=2.0,
)
Try out LEVI yourself at: github.com/ttanv/levi !
LEVI achieves the highest score on every problem where improvement is possible, with an average of 76.5 compared to 71.9 for the next-best framework (GEPA), a +4.6 point improvement over the prior state of the art. On Cloudcast, LEVI reaches a perfect 100.0, indicating the problem is fully solved under the benchmark's scoring function. The largest gains appear on LLM-SQL (+5.8) and Spot Multi (+5.7), while more modest improvements on Spot Single (+0.3) and Transaction Scheduling (+1.1) reflect problems with smaller decision spaces or harder optimization landscapes. Prism remains tied at 87.4 across all frameworks, confirming that the current problem formulation admits a single dominant solution.
LEVI's stratified allocation is the primary driver of cost reduction. By routing the majority of mutations through lightweight models, the per-generation cost drops by roughly an order of magnitude compared to baselines that use GPT-5 or Gemini-3.0-Pro for every call. This allows LEVI to run substantially more generations while still spending less in total: $4.50 per problem on most tasks (Transaction Scheduling: $13), versus $15 to $30 for baselines.
The cost reduction is evidence that the harness-first approach works. When the archive maintains diversity, cheap models suffice for most of the search.
Controlled Architecture Comparison
Same model, same budget, three seeds: isolating the search architecture's contribution.
The main results compare frameworks that differ simultaneously in model choice, budget, and architecture. To isolate the contribution of the search architecture, we run LEVI, OpenEvolve, and GEPA under identical conditions: a single locally-served Qwen3-30B-A3B model, 750 successful evaluations and three random seeds on two representative problems. OpenEvolve required reducing parent count from 5 to 2 for the smaller model and still produced many failures. We report successful eva

[truncated]
