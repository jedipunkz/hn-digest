---
source: "http://muratbuffalo.blogspot.com/2026/06/writing-code-vs-shipping-code.html"
hn_url: "https://news.ycombinator.com/item?id=48457151"
title: "Productivity Effects Across Generations of AI Coding Tools"
article_title: "Writing Code vs. Shipping Code: Productivity Effects Across Generations of AI Coding Tools"
author: "ingve"
captured_at: "2026-06-09T07:22:04Z"
capture_tool: "hn-digest"
hn_id: 48457151
score: 1
comments: 0
posted_at: "2026-06-09T06:08:50Z"
tags:
  - hacker-news
  - translated
---

# Productivity Effects Across Generations of AI Coding Tools

- HN: [48457151](https://news.ycombinator.com/item?id=48457151)
- Source: [muratbuffalo.blogspot.com](http://muratbuffalo.blogspot.com/2026/06/writing-code-vs-shipping-code.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T06:08:50Z

## Translation

タイトル: AI コーディング ツールの世代間の生産性への影響
記事のタイトル: コードの記述と出荷コード: AI コーディング ツールの世代間の生産性への影響
説明: コーディングにおける LLM の変革力には反論の余地がなく、私たちは魔法のコンピューティング ルネサンスを生きているように感じられます。の上...

記事本文:
メインコンテンツにスキップ
このブログを検索
広義の分散システムおよびその他の好奇心について。このサイトの意見は私自身のものです。
コードの作成とコードの出荷: AI コーディング ツールの世代間の生産性への影響
コーディングにおける LLM の変革力には反論の余地がなく、私たちは魔法のようなコンピューティング ルネッサンスの時代を生きているように感じられます。ソーシャル上では、驚くべき数のコード行が生成され、機能が提供され、バグが修正されたことが聞こえてきます。しかし、マクロ経済指標はまだ遅れているようです。エンジニアリング マネージャーに話を聞いてみると、製品の出荷日も奇跡的に 5 分の 1 に短縮されていないことがわかります。
この新聞は10日前に掲載されたばかりです。これは、MIT とウォートンの Mert Demirer、Leon Musolff、Liyuan Yang によるものです。彼らの研究は、AI コーディング ツールから得られる実際の生産性を評価するための構造化された経済モデルを提供しようとしています。著者らは、Microsoft の機密テレメトリーと 100,000 人を超える GitHub 開発者の公開フットプリント (オープンソース ユーティリティから Web アプリ リポジトリに至るまであらゆるものを追跡) を組み合わせることで、AI コード生成の下流での重大なシステム上の摩擦を示しています。
もちろん、私はいつものように同紙に対して懐疑的な批評をしている。この場合、経済学者たちがソフトウェア エンジニアリングの乱雑な非線形の世界を覗いて、そこに「生産階層」の抽象化を課そうとしているため、この傾向は特に高まります。しかし、それらの分析を別の観点から再検討すると、その複雑な生産関数をアムダールの法則の用語に翻訳することが可能になり、その後で私たちが以下で説明するように、独自の評価を開始して独自の結論を導き出すことができます。
この論文の核心は、この単調減衰の議論に基づいています。完全なタスクレベルの速度

AI コーディング ツールによる利点は、本番階層に上がるにつれて滲み出し始めます。著者らは、AI ツールの導入を 3 つの異なる世代層に分類しています。
オートコンプリート (インテリジェントなテキスト予測)、
同期 (Sync) エージェント (クロード コードやカーソルなどの対話型のリアルタイム コード修飾子)、
非同期 (Async) エージェント (自律型非同期エージェント)。
これらのツールのタスクレベルの速度を見ると、素晴らしい数字がわかります。この論文の要約では、オートコンプリート、インタラクティブ同期エージェント、および自律型非同期エージェントは、全体的なコミット アクティビティをそれぞれ累積合計で 40%、140%、180% 増加させると主張しています。しかし、その作業が正式な製品化マイルストーンに向けて上昇するにつれて、改善は大幅に低下します。
オートコンプリートの場合、コードの生の行が +228.2% 爆発的に増加し、実際に出荷されたソフトウェア リリースでは +10.2% のわずかな増加に至るまで、レイヤーごとに流出します。同期エージェントの場合、コード構文の +741.3% という巨大な急増は、最終的な毎週のリリース数としては +20.3% という控えめな値にまで減少しました。
この背後にある数学的モデリングをさらに詳しく見てみましょう。オートコンプリートのパフォーマンスのみを考慮することにより (コード作成レベルでのみ動作するため)、作成者は、モデルの予測と実際の開発者の動作との差異を最小限に抑えるパラメーターを選択しました。この演習を通じて、彼らは約 0.75 の局所的な上流出力弾性 ($\theta$) を抽出しました。この階層化された実稼働モデルでは、$\theta = 0.75$ が垂直パススルー メトリックとして機能します。これは、どの段階でも (生のコミットをクリーンなプル リクエストに変換するなど)、そのレイヤーの成功の 75% は、下のレイヤーから流入する上流の技術資産に完全に依存しており、残りの部分は、最初に追加されたローカルの人的労力を表していることを意味します。

レイヤーで。彼らはソフトウェア開発を垂直方向に順次集計するプロセスとしてモデル化しているため、人間の介入は効率性を高める税金のように作用します。最下位層での初期のコード生産性の大幅な上昇は、階層を登ろうとするにつれて容赦なく何度も何度も 0.75 倍に増加し、経験的データで見られる急激な垂直方向の減衰が数学的に強制されます。
現在、同期エージェントと非同期エージェントはエディター内で行を入力しているだけではありません。これらは、ファイル、コミット、プル リクエストを直接管理するレベルで動作します。この拡張されたレイヤーの対象範囲により、エージェント ワークフローは生産性への貢献をゴール近くまで落とすことができます。図 1 に示すように、垂直減衰チェーンの初期段階を短絡することで、エージェントは作業をより効率的に処理し、出荷されたソフトウェアに対する最終的な影響がオートコンプリートと比較して 2 倍になります。
経済学の公式をアムダールの法則に翻訳する
これらの個々の層の内部で何が起こっているかをモデル化するために、この論文は、入れ子になった Constant Elasticity of Substitution (CES) 生成関数に移行します。ここでは、AI が生成したコードと人間のレビュー作業との間の置換の弾性 ($\sigma$) を厳密な 0.25 として抽出しています。経済用語では、代替弾力性が 1.0 ドルを大きく下回ると、投入物が「強力な補完」であることを意味します。つまり、車のシャーシとタイヤのように、それらは相互に結合/依存しているということです。自動化された工場ラインがタイヤを 10,000% 早く製造できるかどうかは問題ではありません。シャーシの生産をスピードアップしなければ、より多くの車を生産することはできません。
もちろん、これはアムダールの法則とよく似ています。アムダールの法則は、システム全体の高速化は、並列化できない連続的なボトルネックによって厳密に制限されるというものです。
$S_{\text{合計}} = \frac{1}{(1-P) + \frac{P}{S_{\t

内線{タスク}}}}$
ここで、$S_{\text{task}}$ は自動タスク レベルで達成される高速化であり、$P$ はシステム ワークロード全体の「グローバル並列化可能部分」です。
機械の出力と人間の検証の間の代替の弾力性 ($\sigma$) が 0.25 に低下すると、経済モデルはほぼレオンチェフ生産関数と同じように動作します。 (レオンチェフの生産関数は、入力を変更できない正確な比率で組み合わせる必要がある厳密で柔軟性のない生産プロセスを表します。つまり、ある成分の過剰が別の成分の不足を代替することはできません。) これは、人間によるコード レビューが交渉不可能な厳密に逐次的なボトルネック ($1 - P$) であることを示しています。 $\sigma$ が無限であれば、レイヤー全体を効果的に並列化する生の AI テキスト ボリュームで人間の検証を完全に置き換えることができます。しかし、$\sigma = 0.25$ であるため、無限の山の自動コード ($S_{\text{task}} \to \infty$) を問題に投入しても、人間が問題をレビューするために必要な固定の連続時間投資が減ることはありません。
この論文の実際の経験的発見をこの方程式 (最終リリースに対するコミットのグローバル並列部分 $P$ を分離する) に基づいて実行すると、次のことがわかります。
オートコンプリート: $S_{\text{task}} = 1.359\times$、$S_{\text{total}} = 1.102\times \implies \mathbf{P \about 35.0\%}$
同期エージェント: $S_{\text{task}} = 2.091\times$、$S_{\text{total}} = 1.202\times \implies \mathbf{P \about 32.2\%}$
非同期エージェント: $S_{\text{task}} = 2.800\times$、$S_{\text{total}} = 1.300\times \implies \mathbf{P \about 35.9\%}$
えっ！シンプルなインライン オートコンプリート ツールから、リポジトリのクローンを作成し、アウトオブバンドでテスト スイートを実行する完全自律型エージェントに移行しても、グローバル並列化可能割合 ($P$) は変動することを拒否し、acr で 35% 前後を推移しています。

3 世代の AI をすべて網羅します。
では、なぜ同期エージェントと非同期エージェントはオートコンプリートよりも 2 ～ 3 倍多くの最終ソフトウェア リリースを絞り出すことができるのでしょうか? Autocomplete のリリース拡張は 10.2% と控えめですが、同期エージェントはそれを 20.3% に押し上げ、非同期ツールの累積スタックにより最終的な出力ベースラインが 30% 引き上げられます。上記の式から、グローバル並列化可能エンベロープ ($P$) が 35% にロックされている場合、システムは実際にどのように高速化されるのでしょうか?これは、アムダールの法則によれば、システムには最適化のための 2 つのまったく別個のレバーがあるためです。 $P$ を増やすことも、積極的に頑張って $S_{\text{task}}$ を増やすこともできます。オートコンプリートは、リリースと比較してわずか $1.359\times$ のコミット レベルの高速化 ($S_{\text{task}}$) を達成しました。ただし、同期エージェントはローカライズされたタスクを $2.091\times$ まで高速化し、非同期エージェントは $S_{\text{task}}$ を $2.800\times$ まで高速化します。しかし、これはすぐに収益逓減の壁にぶつかります。並列化可能なフットプリント ($P$) が 35% に固定されている場合、局所的なタスクの高速化 ($S_{\text{task}}$) を無限に向かってスケーリングすると、厳しい漸近的な上限が得られます。数学的には、この構成で達成できる合計速度向上の絶対最大値は $1 / (1 - 0.35)$ で、これは出荷されたソフトウェア全体の 53% 増加 ($1.53\times$) というハードキャップに相当します。したがって、自律型ボットがどれほど速くコミットを処理できても、残りの 65% の人間による逐次ボトルネック ($1 - P$) がハードストップとして機能します。
P=0.35は妥当なのでしょうか? 35％の平行部分が嗅覚テストに合格すると思います。開発者に、週のうち実際にコードを書くのに何パーセントを費やしているかと尋ねると、およそ 20% ～ 40% というおおよその数字が返ってくるでしょう。開発者の時間の残りの 65% は、タスクの検索/定義、計画、避けられない人的コミュニケーションの支払いに費やされます。

会議とチームの連携にかかる税金を統合します。
これが、3 つのツール世代すべてで 35% のプロファイルが横ばいであることが私にとって理にかなっている理由です。生成 AI は、高速でコードを量産する ($S_{\text{task}}$ を最大化する) ことでコーディング サンドボックスを大幅に強化できますが、残りの仕事を支配する体系的な推論、組織の合意、問題やモデルの深い理解を並列化することはできません。これらの部分も自動化する方法が見つかるまで、ソフトウェアの出荷は本質的に人間が調整するプロセスのままになるでしょう。
分散システム設計のヒント
エージェント的な自己: AI と自己改善の類似点
分散システムについて学ぶ: どこから始めればよいでしょうか?
分散システムの基礎に関する論文
Cloudspecs: Looking Glass によるクラウド ハードウェアの進化
AI の支配者をサポート: データ システムをエージェントファーストになるように再設計する

## Original Extract

The transformative power of LLMs in coding has been irrefutable, and it feels like we are living through a magical computing renaissance. On...

Skip to main content
Search This Blog
On distributed systems broadly defined and other curiosities. The opinions on this site are my own.
Writing Code vs. Shipping Code: Productivity Effects Across Generations of AI Coding Tools
The transformative power of LLMs in coding has been irrefutable, and it feels like we are living through a magical computing renaissance. On the socials, we hear impressive numbers of lines of code generated, features delivered, and bugs fixed. But, the macroeconomic indicators seem to be still lagging. Heck, if you talk with an engineering manager, you find that their product shipping dates haven't miraculously compressed by a factor of five, either.
This paper just landed 10 days ago. It is from MIT and Wharton by Mert Demirer, Leon Musolff, and Liyuan Yang. Their study attempts to provide a structured economic model for evaluating actual productivity obtained from AI coding tools. By pairing confidential Microsoft telemetry with the public footprints of over 100,000 GitHub developers (tracking everything from open-source utilities to web app repositories), the authors show significant systemic friction downstream of AI code generation.
Of course, I do my usual skeptical critic of the paper. In this case, this is especially heightened because these are economists peeking into the messy non-linear world of software engineering and trying to impose a "production hierarchy" abstraction onto it. But if we reconsider their analysis from a different perspective, it becomes possible to translate their complex production functions into Amdahl's Law terms, and then we can start doing our own evaluations and draw our own conclusions as I discuss below.
The core of the paper rests on this monotonic decay argument. The sheer task-level velocity gains we see from AI coding tools start to bleed out as they move up the production hierarchy. The authors break down AI tool adoption into three distinct generational tiers:
Autocomplete (intelligent text prediction),
Synchronous (Sync) Agents (interactive, real-time code modifiers like Claude Code or Cursor),
Asynchronous (Async) Agents (autonomous async agents).
When we look at the task-level velocity of these tools, we see impressive numbers. The paper's abstract claims that autocomplete, interactive sync agents, and autonomous async agents increase overall commit activity by cumulative totals of 40% , 140%, and 180% respectively. But as that work climbs toward an official production milestone, the improvements decay significantly.
For Autocomplete, the +228.2% explosion in raw lines of code bleeds out layer-by-layer until becomes a meager +10.2% increase in actual shipped software releases. For Sync agents, a gigantic +741.3% surge in code syntax reduces down to a modest +20.3% final weekly releases.
Let's dive deeper on the mathematical modeling behind this. By taking only the performance of Autocomplete into account (because it operates exclusively at the code-writing level), the authors chose parameters that minimized the differences between their model's predictions and actual developer behavior. Through this exercise, they extracted a local Upstream Output Elasticity ($\theta$) of approximately 0.75. In this layered production model, $\theta = 0.75$ acts as a vertical pass-through metric. It means that at any given stage (say, turning raw commits into clean pull requests), 75% of that layer's success leans entirely on the upstream technical assets flowing into it from the layer below, while a remaining fraction represents the local human effort added at that layer. Because they model software development as a vertically sequential aggregation process, that human intervention operates like a compounding efficiency tax. A massive initial code productivity surge at the bottom layer gets relentlessly multiplied by 0.75 over and over again as it attempts to climb the hierarchy, mathematically forcing the steep vertical attenuation we see in the empirical data.
Now, the sync and async agents aren't just typing lines inside an editor. They operate at a level where they directly manage files, commits, and pull requests. This expanded layer coverage allows agentic workflows to drop their productivity contributions closer to the finish line. By short-circuiting the early stages of the vertical decay chain, agents handle the work more efficiently, doubling the final impact on shipped software compared to autocomplete as seen in Figure 1.
Translating the Economics formulas to Amdahl’s Law
To model what happens inside each of these individual layers, the paper transitions to a nested Constant Elasticity of Substitution (CES) production function. Here, they extract an Elasticity of Substitution ($\sigma$) between AI-generated code and human review effort of a rigid 0.25. In economic lingo, an elasticity of substitution well below $1.0$ means the inputs are "strong complements". That means they are tied/dependent together like a car chassis and tires. It doesn't matter if an automated factory line can manufacture tires 10,000% faster; if you don't speed up the production of the chassis, you don't get more cars.
This of course looks a whole lot like Amdahl’s Law, which dictates that the overall speedup of a system is strictly limited by its sequential un-parallelizable bottlenecks:
$S_{\text{total}} = \frac{1}{(1-P) + \frac{P}{S_{\text{task}}}}$
where $S_{\text{task}}$ is the speedup achieved at the automated task level, and $P$ is the "Global Parallelizable Fraction" of the entire system workload.
When the elasticity of substitution ($\sigma$) between machine output and human validation drops to 0.25, the economic model behaves almost exactly like a Leontief production function. (A Leontief production function describes a strict, zero-flexibility production process where inputs must be combined in exact, unalterable proportions, meaning an excess of one ingredient cannot substitute for a shortage of another.) This dictates that human code review is a non-negotiable strictly sequential bottleneck ($1 - P$). If $\sigma$ were infinite, you could completely substitute human verification with raw AI text volume effectively parallelizing the entire layer. But because $\sigma = 0.25$, throwing an infinite mountain of automated code ($S_{\text{task}} \to \infty$) at the problem does nothing to diminish the fixed, sequential time investment required for a human to review it.
If we run the paper's real-world empirical findings through this equation (isolating the global parallel fraction $P$ for commits against final releases), we find the following:
Autocomplete: $S_{\text{task}} = 1.359\times$, $S_{\text{total}} = 1.102\times \implies \mathbf{P \approx 35.0\%}$
Sync Agents: $S_{\text{task}} = 2.091\times$, $S_{\text{total}} = 1.202\times \implies \mathbf{P \approx 32.2\%}$
Async Agents: $S_{\text{task}} = 2.800\times$, $S_{\text{total}} = 1.300\times \implies \mathbf{P \approx 35.9\%}$
Huh! Even going from a simple inline autocomplete tool to fully autonomous agents that clone repositories and run test suites out-of-band, the global parallelizable fraction ($P$) refuses to budge and hovers around 35% across all three generations of AI!
Then how come Sync and Async agents manage to squeeze out 2-3x more final software releases than Autocomplete? While Autocomplete notches a modest 10.2% release expansion, Sync agents push it to 20.3%, and the cumulative stack of Async tools lifts the final output baseline up by 30%. If the global parallelizable envelope ($P$) is locked at 35% from the formulas above, how is the system actually accelerating? This is because according to Amdahl's Law, a system has two entirely separate levers for optimization. You can either increase $P$, or you can aggressively push harder and increase $S_{\text{task}}$. Autocomplete achieved a commit-level speedup ($S_{\text{task}}$) of just $1.359\times$ relative to releases. But Sync agents drive that localized task speedup to $2.091\times$, and Async agents push $S_{\text{task}}$ to $2.800\times$. But this hits a wall of diminishing returns quickly. When the parallelizable footprint ($P$) is pinned to 35%, scaling the localized task speedup ($S_{\text{task}}$) toward infinity yields a hard asymptotic ceiling. Mathematically, the absolute maximum total speedup this configuration can ever achieve is $1 / (1 - 0.35)$, which works out to a hard cap of a 53% overall increase ($1.53\times$) in shipped software. So, no matter how fast an autonomous bot can process a commit, the remaining 65% human sequential bottleneck ($1 - P$) acts as a hard stop.
Is P=0.35 sensible? I think that 35% parallel fraction passes the smell test. If you ask any developer what percentage of their week is spent actually writing code, they'll give you a number right in this ballpark, around 20% to 40%. The remaining 65% of the developers' time is consumed by finding/defining the task, planning, and paying the inevitable human communication tax of meetings and team alignment.
This is why that flatlining 35% profile across all three tool generations makes sense to me. Generative AI can supercharge the coding sandbox by churning out code at high speed (maxing out $S_{\text{task}}$), but it can't parallelize the systemic reasoning, organizational consensus, and deep problem/model comprehension that dominates the rest of the job. Until we find a way to automate those parts as well, shipping software will remain an inherently human-throttled process.
Hints for Distributed Systems Design
The Agentic Self: Parallels Between AI and Self-Improvement
Learning about distributed systems: where to start?
Foundational distributed systems papers
Cloudspecs: Cloud Hardware Evolution Through the Looking Glass
Supporting our AI overlords: Redesigning data systems to be Agent-first
