---
source: "https://pytorch.org/blog/from-minutes-to-seconds-llm-guided-autotuning-for-helion-kernels/"
hn_url: "https://news.ycombinator.com/item?id=48590151"
title: "From Minutes to Seconds: LLM-Guided Autotuning for Helion Kernels"
article_title: "From Minutes to Seconds: LLM-Guided Autotuning for Helion Kernels – PyTorch"
author: "matt_d"
captured_at: "2026-06-18T21:46:49Z"
capture_tool: "hn-digest"
hn_id: 48590151
score: 3
comments: 0
posted_at: "2026-06-18T19:16:58Z"
tags:
  - hacker-news
  - translated
---

# From Minutes to Seconds: LLM-Guided Autotuning for Helion Kernels

- HN: [48590151](https://news.ycombinator.com/item?id=48590151)
- Source: [pytorch.org](https://pytorch.org/blog/from-minutes-to-seconds-llm-guided-autotuning-for-helion-kernels/)
- Score: 3
- Comments: 0
- Posted: 2026-06-18T19:16:58Z

## Translation

タイトル: 数分から秒へ: LLM ガイドによる Helion カーネルの自動チューニング
記事のタイトル: 数分から秒へ: LLM ガイドによる Helion カーネルの自動チューニング – PyTorch

記事本文:
メインコンテンツにスキップ
PyTorch Conference North America にご参加ください · 10 月 20 ～ 21 日 · カリフォルニア州サンノゼ
検索
検索を閉じる
検索
メニュー
学ぶ
始めましょう
PyTorch の紹介 - YouTube シリーズ
数分から数秒へ: LLM ガイドによる Helion カーネルの自動チューニング
Helion は、パフォーマンス ポータブル機械学習カーネル用の PyTorch のドメイン固有言語 (DSL) であり、パフォーマンスの自動チューニングに大きく依存しています。現在、Helion 検索では、尤度自由ベイジアン最適化 (LFBO) を利用して、最もパフォーマンスの高い構成を見つけます。 LFBO は強力なベースラインであり、うまく機能しますが、カーネルごとに何百ものコンパイルとベンチマーク サイクルを実行する必要があります。この目的を達成するために、LFBO レベルのカーネル パフォーマンス (ジオメアン 1.009 倍) に匹敵する LLM ガイド付き自動チューナーを導入し、ベンチマークを最大 10 分の 1 の構成数で最大 6.7 分の 1 少ない実時間で実現します。 LLM が 5% を超える差がある少数のカーネルでは、ハイブリッド戦略 (LLM シードとそれに続く LFBO 改良) により、完全な LFBO 検索よりも 3 倍ほど安価なままギャップを埋めることができます。最後に、結果は LLM モデルにほとんど依存しません。Opus-4.8、gpt-5.5、および Sonnet-4.6 は相互に数パーセント以内のパフォーマンスを示します。これは、LLM ガイドによる自動チューニングが実稼働品質で劇的に高速なカーネル チューニングへの実用的なアプローチであることを示しています。
自動チューニングは Helion のバックボーンであり、高性能でポータブルな ML カーネルを作成するための PyTorch の DSL です。すべての Helion カーネルは、ターゲット ハードウェアで最高のパフォーマンスを達成するために、広大な高次元構成空間 (タイル サイズ、ブロック サイズ、num_warps、num_stages、詳細についてはドキュメントを参照) にわたって調整されます。チューニング時間の短縮は、開発速度と本番展開にとっても重要であり、Helion の導入に影響を与えます。
オートチューナーは組み合わせ空間を検索して構成、ベンチを見つけます。

それぞれの構成をアーク化し、最良の状態を維持します。その検索をより高速かつスマートにすることが、積極的な取り組み分野です。 Helion の現在のデフォルトのオートチューナーは LFBO (尤度フリー ベイジアン最適化) を使用します。これにより、軽量のランダム フォレスト分類器がベンチマーク データのオンザフライ検索中にトレーニングされ、どの構成が有望な候補であるかを予測することが学習されます。予測を使用して、空間をターゲットにジャンプするために最も重要なパラメーターに焦点を当てます。 LFBO 検索は、NVIDIA および AMD GPU でのカーネル パフォーマンスとチューニング時間の両方で大幅な改善が見られたため、現在はデフォルトになっています。詳細については、PyTorch ブログ「ベイジアン最適化による Helion での自動チューニングの加速」を参照してください。
LFBO は強力なベースラインであり、うまく機能しますが、カーネルごとに何百ものコンパイルとベンチマーク サイクルを実行する必要があります。やみくもに検索を開始するのではなく、LLM にカーネルについて推論して構成を提案するよう依頼できたらどうでしょうか?それが LLM ガイド付き自動チューナーです。自動チューニングの各ラウンドで、LLM にはカーネル、ワークロード、およびこれまでの最良の構成が表示され、試行する新しい構成が提案されます。このブログでは、LLM ガイド付きオートチューナーがどのように動作するかを説明し、B200 上の 33 (11 カーネル x 3 シェイプ) のケースで LLM ガイド付き検索と LFBO 検索を比較したベンチマーク結果を示します。結果は、新しい LLM ベースのアプローチが 10 分の 1 の構成をコンパイル/ベンチマークしながら LFBO レベルのカーネル パフォーマンスに達し、実稼働時間が 6.7 分の 1 に短縮されることを示しています。また、両方の長所を組み合わせたハイブリッド検索も導入しました。これは、LLM を使用してパフォーマンスの高い構成を迅速に取得し、その後にきめの細かい検索に LFBO を使用します。
LLM ガイド付きオートチューナーの仕組み
新しい LLM ベースのオートチューナーは、プロンプトとフィードバックの複数のサイクルを通じて動作します。

人口ベースの検索を実行します。最初のフェーズでは、Helion はカーネルと関連する詳細を LLM に提供して、一連の候補構成を要求します。 LLM が応答すると、Helion は構成をコンパイルしてベンチマークを実行し、最高のパフォーマンスを発揮する構成を保持します。その後、後続の改良ラウンドが行われ、LLM には最も成功した構成、そのパフォーマンス メトリクス、および特定の突然変異を導くための成功したパターンの分析が与えられます。大幅なパフォーマンスの向上が検出されない場合、プロセスは早期に終了します。 B200 GPU で実行される rms_norm カーネルのプロセス例を示します。
最初のプロンプトでは、LLM の役割を設定し、ノブを提供し、出力コントラクトを指定します。
あなたは、Helion/Triton カーネルの GPU カーネル自動チューナーのエキスパートです。
提供された構成スペースとデフォルト構成を、許可されるフィールド名、スカラー対リスト、必要なリストの長さ、有効な範囲およびデフォルトの信頼できる情報源として使用します。
- 縮小された JSON を 1 行で返します: {"configs":[...]}。値下げ/フェンス/コメントはありません。
-変更したいフィールドのみを指定します。未指定 = デフォルト。
- リスト値フィールドの場合は、スペースに表示されている正確な必要な長さの JSON 配列を出力します。
- フィールドの構造、長さ、または許可される値が不明な場合は、推測せずに省略してください。
Helion は、カーネル、ターゲット ハードウェア、構成スペースも提供します。 rms_norm カーネル プロンプトには次のものが含まれます。
カーネル ソース: 実際の @helion.kernel ソース コード
入力テンソル: 例: arg[0]:shape=[4096, 1024], dtype=torch.float16, …
GPU ハードウェア: 例: NVIDIA B200、148 SM、178.4 GB、2048 スレッド/SM
構成スペース: タイプ/範囲を含むすべての調整可能なフィールド
デフォルト設定: ベースライン設定。
Helion コンパイラはカーネルも分析して、プロンプトにヒューリスティックを追加します。 rms_norm の場合:

Helion のコンパイラは、このカーネルの構造を静的に分析し、次の構造事前分布を導き出しました。それらを強力な出発点として扱います。
**コンパイラ由来のシード構成:**
```json
{"block_sizes":[1],"load_eviction_policies":["last","last","last","last","last"],"reduction_loops":[null]}
15 個の構成を含む縮小された JSON。 Opus-4.8 の rms_norm に対する最初の応答:
{"構成":[
{"block_sizes":[1],"load_eviction_policies":["last","last","last","last","last"]},
{"ブロックサイズ":[1]},
{"block_sizes":[4],"load_eviction_policies":["last", "..."],"num_warps":8},
{"block_sizes":[8],"load_eviction_policies":["last", "..."],"num_warps":8,"num_stages":2},
{"block_sizes":[16],"load_eviction_policies":["last", "..."],"num_warps":8,"num_stages":4},
{"block_sizes":[1],"load_eviction_policies":["last", "..."],"pid_type":"persistent_blocked"}
…。
]}
Helion のハーネスはこれを解析し、不正な構成や重複した構成を削除してコンパイルしてベンチマークを行います。
ベンチマークの後、後続の各ラウンドでは、検索状態から構築された絞り込みプロンプトが送信されます。
検索状態: ラウンド数、人口規模、これまでの最高のパフォーマンス。
アンカー設定: 周囲に変更する上位の設定。
結果: ベンチマークされた構成のパフォーマンスを測定しました。
上位/失敗した構成パターン: どのフィールド値が高速構成と失敗/低速構成に相関するか。
次のステップ: 構成の失敗率に基づく推奨事項
したがって、モデルは最良のものを重視し、失敗したパターンを回避します。各ラウンドからの相対的な改善が ~0.5% を下回ると、フィードバック ループは早期に停止します。
LLM シードの LFBO: 両方の長所
また、マイクロアーキテクチャのノブを未調査のままにする LLM の傾向に対処するためのハイブリッド戦略 (LLM シード LFBO 検索) も検討します。このアプローチは、LLM と LFBO の両方の相補的な利点を統合します。LLM は強力なスターを提供します。

LFBO はローカル検索に優れています。
ステージ 1 – LLM シード : プロセスは、上記の「最初のプロンプト」で説明したように、LLM ガイド付き検索の 1 ラウンドから始まります。 Helion ベンチマークを使用して、最上位の構成を検証して保持します。
ハンドオフ: LLM で生成された最も成功した構成は、LFBO のサロゲート モデルをトレーニングする次のフェーズの開始点として機能します。これにより、LFBO は白紙の状態から始めるのではなく、有望な地域を即座に知ることができます。
ステージ 2 – LFBO 絞り込み : LFBO 検索は、シードされた初期母集団で実行されます。 LFBO は、ランダムフォレスト分類器の更新、上位候補の予測、特徴の重要性に基づいた重要なパラメーターの変更という反復ループを実行します。このサイクルは、パフォーマンスの向上が失速するか、最大反復数 (20) に達するまで継続します。
システムは、両方のステージにわたって見つかった最適な構成を返します。高品質の開始点と情報に基づいたサロゲートを活用することにより、ハイブリッド検索はコールド LFBO 検索よりも大幅に高速に収束します。この効率により、LLM が見つけられない可能性のある特定のマイクロアーキテクチャのノブの微調整に検索予算を集中させることができます。
NVIDIA B200 上の小規模、中規模、および大規模なシェイプについて、11 のカーネル (matmul (square + Split-K)、grouped-GEMM、attention、fp8-attention、softmax、rms_norm、rope、swiglu、mamba2、および Gated-delta-net) にわたる Opus 4.8 を使用した LFBO (全力を尽くした LFBOTreeSearch) と LLM ガイド付き検索を比較します。
ここで LLM が威力を発揮し、LFBO の品質に匹敵する検索コストを大幅に削減します。
Geomean 構成のベンチマーク: LLM ガイド付き自動チューナーの構成が 9.8 倍少ない (カーネルあたり ~55 対 ~546)。これは、新しいアプローチの有効性を示す、マシンに依存しない指標です。
幾何平均実時間: 6.7 倍短縮

エンドツーエンドのチューニング時間 (39 秒 vs 261 秒)。384 スレッドのホストで測定。エンドツーエンドのチューニング時間は、構成の生成 (LLM の API ラウンドトリップを含む)、すべての候補の Triton/ptxas コンパイル、およびすべての候補の GPU ベンチマークで構成されます。
エンドツーエンドの調整時間の大部分は候補構成のコンパイルにあり、Helion はそれらを CPU コア間で並行してプリコンパイルします。このホストには数百のスレッドがあるため、コンパイルは高度に並列化されます。コアの数が少ないマシンでは、LLM の構成が最大 10 分の 1 に減少するため、それに比例して実時間の大幅な短縮が実現されます。マシンに依存しないメトリクスは、ベンチマークされた構成の数と、最適な構成が最適な結果に収束するまでの速度です (以下を参照)。
結果 2: LLM は LFBO 予算の最初の ~7% に収束
これまでの最良の構成を検索労力に対してプロットすると、違いが明確になります。 LLM は、数十の構成でプラトーに落ちます。
12 個のコンバージェンス カーネルすべてにおいて、LLM は LFBO の予算の最初の約 7% 以内でプラトーに落ちます。グループ化された GEMM (g=4、m=512) では、同じカーネル パフォーマンスでも LFBO よりも構成が 18 倍少なくなります。
結果 3: LLM は LFBO レベルのパフォーマンスを実現
カーネル パフォーマンスに関しては、LLM は LFBO とほぼ同等で、LLM カーネル/LFBO カーネル レイテンシの地理平均パフォーマンスは 1.009 倍です。
したがって、LLM は適切な構成を迅速に提供しますが、LFBO はより多くの構成探索と調整時間を犠牲にしてパフォーマンスを向上させることができます。 LLM が LFBO に対してカーネル パフォーマンスで 5% 以上劣るケースが 8 件あります。
ハイブリッド検索でギャップを埋めることができるか?
LLM の弱点がノブの微調整を放置していることである場合は、LLM シード LFBO TreeSearch を使用したハイブリッド検索戦略を使用できます。
LLM が LFBO を 5% 以上下回る 8 つのケースに対してハイブリッド検索を実行しました。
ハイブリッド戦略の改善

すべてのケースでカーネルのパフォーマンスが向上し、6/8 のケースで LFBO との差が縮まりました。 mamba2 ファミリは依然として LFBO よりも劣っており、私たちはこのギャップを埋めるために LLM ヒューリスティックを改善することを検討しています。
自動チューニング時間の点では、ハイブリッド検索は LFBO よりも大幅に効率的です。 8 つのカーネル全体で 4 倍少ない構成 LFBO を探索するため、エンドツーエンドの自動チューニング時間が 3 倍高速になります。 LLM のみ、ハイブリッド、および LFBO を比較した幾何平均結果を以下に示します。
また、調査された構成の数とその調整時間を比較した個々のカーネルの結果も以下に示します。
上記のすべてでは、Claude Opus-4.8 という 1 つのモデルを使用しました。作業を行っているモデルが耐荷重性があるのか​​、それとも有能な LLM が同じ場所に到達するのか、疑問に思う人もいるかもしれません。この目的を達成するために、OpenAI gpt-5.5 と Claude Sonnet-4.6 の 2 つのモデルを使用して、完全な 33 カーネル インスタンスにわたる LLM のみの検索 (LLM ガイド付き検索) のベンチマークを行い、Opus 4.8 ベースラインと比較します。
geomean では、3 つのモデルすべてが非常に類似したパフォーマンスを示し、興味深いことに、Sonnet-4.6 は最も少ない構成数でそれを実行しました。
私たちが答えようとした質問は、「LLM は、LFBO 検索と同様に Helion カーネルを自動調整できますが、はるかに安価に実行できますか? B200 でベンチマークされた 33 カーネル スイート全体で、答えは「はい」です。
効率

[切り捨てられた]

## Original Extract

Skip to main content
Join us at PyTorch Conference North America · Oct 20-21 · San Jose, CA
Search
Close Search
search
Menu
Learn
Get Started
Intro to PyTorch – YouTube Series
From Minutes to Seconds: LLM-Guided Autotuning for Helion Kernels
Helion, PyTorch’s domain-specific language (DSL) for performance portable machine learning kernels, heavily relies on autotuning for performance. Currently Helion searches utilize the Likelihood-Free Bayesian Optimization (LFBO) to find the most performant configs. LFBO is a strong baseline which works well, but it still grinds through hundreds of compile-and-benchmark cycles per kernel. To this end, we introduce an LLM-guided autotuner that matches LFBO-level kernel performance (geomean 1.009X) while benchmarking ~10X fewer configurations in ~6.7X less wall-clock time. For the handful of kernels where the LLM trails by >5%, a hybrid strategy (LLM seeding followed by LFBO refinement) closes the gap while remaining ~3X cheaper than the full LFBO search. Finally, the result is largely LLM model-independent — Opus-4.8, gpt-5.5, and Sonnet-4.6 perform within a couple percent of each other — showing that LLM-guided autotuning is a practical approach to dramatically faster kernel tuning at production quality.
Autotuning is the backbone of Helion – PyTorch’s DSL for authoring performant and portable ML kernels. Every Helion kernel is tuned across a vast, high-dimensional configuration space (tile sizes, block sizes, num_warps, num_stages, see documentation for more) to reach peak performance on the target hardware. Reducing the tuning time is also critical for developer velocity and production deployment, which impacts Helion adoption.
The autotuner searches the combinatorial space to find configurations, benchmarks each configuration, and keeps the best. Making that search faster and smarter is an active area of work. Helion’s current default autotuner uses LFBO (Likelihood-Free Bayesian Optimization), where a lightweight Random Forest classifier is trained during the search on the fly on the benchmarked data, learning to predict which configurations are promising candidates. It uses the prediction to focus on the parameters that matter the most to take targeted jumps through the space. LFBO search is now the default, as it showed substantial improvements in both kernel performance and tuning time on NVIDIA and AMD GPUs. See our PyTorch blog “ Accelerating Autotuning in Helion with Bayesian Optimization ” for more details.
LFBO is a strong baseline which works well, but it still grinds through hundreds of compile-and-benchmark cycles per kernel. What if, instead of starting the search blindly, you could ask an LLM to reason about the kernel and propose configurations? That’s the LLM-guided autotuner – for each round of autotuning, an LLM is shown the kernel, the workload, and the best-so-far configs to propose new configs to try. In this blog, we describe how the LLM-guided autotuner works and show benchmarking results comparing the LLM-guided search to LFBO search on 33 (11 kernels x 3 shapes) cases on B200. Results show that the new LLM-based approach reaches LFBO-level kernel performance while compiling/benchmarking 10X less configs, leading to 6.7X less wall-clock time. We also introduce a hybrid search to combine the best of both worlds, which uses an LLM to quickly get to a performant configuration, followed by LFBO for fine-grained search.
How the LLM-Guided Autotuner Works
Operating through multiple cycles of prompts and feedback, the new LLM-based autotuner executes a population-based search. In the initial phase, Helion provides the kernel and the associated details to the LLM to ask for a set of candidate configurations. Once LLM responds, Helion compiles and benchmarks the configs, retaining the top-performing configurations. Subsequent refinement rounds then occur, where the LLM is given the most successful configs, their performance metrics, and an analysis of successful patterns to guide specific mutations. If no significant performance gains are detected, the process terminates early. An example process is shown for an rms_norm kernel running on a B200 GPU.
The initial prompt sets the role of the LLM, provides the knobs, and gives the output contract:
You are an expert GPU kernel autotuner for Helion/Triton kernels.
Use the provided Configuration Space and Default Configuration as the source of truth for allowed field names, scalar-vs-list, required list lengths, valid ranges and defaults.
-Return minified JSON on a single line: {"configs":[...]}. No markdown/fences/comments.
-Only specify fields you want to change; unspecified = default.
-For list-valued fields, emit a JSON array of the exact required length shown in the space.
-If unsure about a field's structure, length, or allowed values, omit it instead of guessing.
Helion also provides the kernel, the target hardware, and the configuration space. The rms_norm kernel prompt has:
Kernel source: The actual @helion.kernel source code
Input Tensors: e.g.: arg[0]: shape=[4096, 1024], dtype=torch.float16, …
GPU Hardware: e.g.: NVIDIA B200, 148 SMs, 178.4 GB, 2048 threads/SM
Configuration Space: Every tunable field with type/range
Default Configuration: The baseline config.
The Helion compiler also analyzes the kernel to add heuristics to the prompt. For rms_norm:
Helion's compiler statically analyzed this kernel's structure and derived the following structural priors. Treat them as strong starting points.
**Compiler-derived seed config(s):**
```json
{"block_sizes":[1],"load_eviction_policies":["last","last","last","last","last"],"reduction_loops":[null]}
A minified JSON with 15 configs. Opus-4.8’s first reply for rms_norm:
{"configs":[
{"block_sizes":[1],"load_eviction_policies":["last","last","last","last","last"]},
{"block_sizes":[1]},
{"block_sizes":[4],"load_eviction_policies":["last", "..."],"num_warps":8},
{"block_sizes":[8],"load_eviction_policies":["last", "..."],"num_warps":8,"num_stages":2},
{"block_sizes":[16],"load_eviction_policies":["last", "..."],"num_warps":8,"num_stages":4},
{"block_sizes":[1],"load_eviction_policies":["last", "..."],"pid_type":"persistent_blocked"}
….
]}
Helion’s harness parses this, drops malformed and duplicate configs to compile and benchmark them.
After benchmarking, each subsequent round sends a refinement prompt built from the search state:
Search state: Round number, population size, best perf so far.
Anchor configs: Top configs to mutate around.
Results: Measured performance for the benchmarked configs.
Top/failed config patterns: Which field values correlate with fast vs. failed/slow configs.
Next Step: Recommendations based on failure-rate of configs
So the model anchors on the best and avoids the patterns that failed. The feedback loop stops early if relative improvement from each round drops below ~0.5%.
LLM-Seeded LFBO: The Best of Both Worlds
We also explore a hybrid strategy (LLM-Seeded LFBO Search) to address LLM’s tendency to leave micro-architectural knobs unexplored. This approach merges the complementary advantages of both LLM and LFBO: the LLM provides a strong starting point, while LFBO excels at local search.
Stage 1 – LLM Seeding : The process begins with a single round of LLM-Guided Search, as described “The Initial Prompt” above. Helion benchmarks to validate and retain the top configs.
Handoff: The most successful LLM-generated configs serve as the starting point for the next phase to train LFBO’s surrogate model. This allows LFBO to begin with immediate knowledge of promising regions rather than starting from a blank slate.
Stage 2 – LFBO Refinement : LFBO Search is executed with its initial population seeded. LFBO performs its iterative loop: updating the Random-Forest classifier, predicting top candidates, and mutating critical parameters based on feature importance. This cycle continues until performance gains stall or reaches the maximum number (20) of iterations.
The system returns the optimal configuration found across both stages. By leveraging a high-quality starting point and an informed surrogate, the hybrid search converges significantly faster than a cold LFBO search. This efficiency allows the search budget to be focused on fine-tuning the specific micro-architectural knobs that the LLM may not find.
We compare LFBO (LFBOTreeSearch with full effort) to LLM-Guided Search using Opus 4.8 across 11 kernels — matmul (square + split-K), grouped-GEMM, attention, fp8-attention, softmax, rms_norm, rope, swiglu, mamba2, and gated-delta-net — on small, medium, and large shapes on NVIDIA B200.
This is where the LLM shines, matching LFBO’s quality with significantly smaller search cost.
Geomean configs benchmarked: 9.8X fewer configs (~55 vs ~546 per kernel) for LLM-guided autotuner . This is a machine-independent metric that demonstrates the efficacy of the new approach.
Geomean wall-clock time: 6.7X less end-to-end tuning time (39 s vs 261 s), measured on a 384-thread host . The end-to-end tuning time consists of config generation (for the LLM, including its API round-trips), Triton/ptxas compilation of every candidate, and GPU benchmarking of every candidate.
End-to-end tuning time is dominated by compiling candidate configs, where Helion precompiles them in parallel across CPU cores. As this host has 100s of threads, compilation is heavily parallelized. On a machine with fewer cores, the LLM’s ~10X fewer configs would translate into a proportionally larger wall-clock time reduction. The machine-independent metrics are the number of configs benchmarked and how fast the best configs converge to their optimal results (shown below).
Result 2: LLM Converges in the First ~7% of LFBO Budget
Plotting best-config-so-far against search effort makes the difference vivid. The LLM drops to its plateau in a few dozen configs.
Across all 12 convergence kernels, the LLM drops to its plateau inside roughly the first ~7% of LFBO’s budget. On grouped GEMM (g=4, m=512), that’s 18X fewer configs than LFBO at the same kernel performance.
Result 3: LLM Delivers LFBO-level Performance
On kernel performance, the LLM is roughly on-par with LFBO, with the geomean performance of LLM kernel/LFBO kernel latency being 1.009X.
Hence LLM gives you a good config fast, while LFBO can outperform at the cost of more config exploration and tuning time. There are 8 cases where LLM loses to LFBO by more than 5% in kernel performance.
Can the Hybrid Search Close the Gap?
If LLM’s weakness is leaving fine tuning the knobs, we can use our hybrid search strategy with LLM-Seeded LFBO TreeSearch.
We ran the hybrid search on the 8 cases where the LLM trails LFBO by more than 5%.
The hybrid strategy improves kernel performance in all cases and closes the gap to LFBO in 6/8 cases. The mamba2 family still does worse than LFBO and we are investigating improving the LLM heuristics to close this gap.
In terms of autotuning time, hybrid search is significantly more efficient than LFBO. Across the 8 kernels, it explores 4X fewer configs LFBO, leading to 3X faster end-to-end autotuning time. The geomean results comparing LLM-only, hybrid, and LFBO are shown below.
We also present the individual kernel results below comparing the number of explored configs as well as their tuning times:
Everything above used one model: Claude Opus-4.8. One may ask whether the model doing the work is load-bearing, or whether any capable LLM gets you the same place. To this end, we benchmark the LLM-only search (LLM-Guided Search) across the full 33-kernel instances with two more models, OpenAI gpt-5.5 and Claude Sonnet-4.6, to compare to the Opus 4.8 baseline.
In geomean, all 3 models performed very similarly and interestingly, Sonnet-4.6 did it with the fewest number of configs.
The question we set out to answer was, “Can an LLM autotune Helion kernels as well as the LFBO search, but far more cheaply? Across a 33-kernel suite, benchmarked on B200, the answer is yes.
The efficiency

[truncated]
