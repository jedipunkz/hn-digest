---
source: "https://andlukyane.com/blog/cayleypy-kaggle-with-claude"
hn_url: "https://news.ycombinator.com/item?id=48832621"
title: "Cayley graph search with Claude Code: what puzzle competitions look like in 2026"
article_title: "Cayley graph search with Claude Code: what puzzle competitions look like in 2026 – Andrey Lukyanenko"
author: "Artgor"
captured_at: "2026-07-08T15:11:13Z"
capture_tool: "hn-digest"
hn_id: 48832621
score: 2
comments: 0
posted_at: "2026-07-08T14:40:28Z"
tags:
  - hacker-news
  - translated
---

# Cayley graph search with Claude Code: what puzzle competitions look like in 2026

- HN: [48832621](https://news.ycombinator.com/item?id=48832621)
- Source: [andlukyane.com](https://andlukyane.com/blog/cayleypy-kaggle-with-claude)
- Score: 2
- Comments: 0
- Posted: 2026-07-08T14:40:28Z

## Translation

タイトル: クロード コードを使用したケイリー グラフ検索: 2026 年のパズル コンペティションはどのようになるか
記事のタイトル: Claude Code を使用した Cayley グラフ検索: 2026 年のパズル競技会の様子 – Andrey Lukyanenko
説明: 2 つの CayleyPy Kaggle コンペティションで得た 2 つの経験: 学習ヒューリスティック ビーム探索が巨大な Cayley グラフをどのように解決するか、および Claude Code がエンジニアリング ループをどのように簡素化して高速化するか。

記事本文:
Claude Code を使用した Cayley グラフ検索: 2026 年のパズルコンテストはどのようになるか – Andrey Lukyanenko
メインコンテンツにスキップ
アンドレイ・ルキャネンコ
プロジェクト
Claude Code を使用した Cayley グラフ検索: 2026 年のパズルコンテストはどうなるか
Claude Code を使用した Cayley グラフ検索: 2026 年のパズルコンテストはどうなるか
過去数か月間、私は CayleyPy シリーズの 2 つの Kaggle コンペティション、IHES Picture Cube と Megaminx に取り組みました。どちらも組み合わせパズル ソルバーであり、巨大な暗黙的なケイリー グラフの検索として再定式化できます。パズルの各状態は頂点であり、可能な動きはエッジであり、解決とは、同一性へ戻る短いパスを見つけることを意味します。これは長期的な活動です。私たちはこれらのパズルを数年間研究しており、いくつかの論文を発表しています。今年、私たちはそれらの解決に向けて新たな一歩を踏み出しました。
2024 年にこれに取り組んだことを覚えています。当時、すべてのコードは手作業で書かれており、デバッグとドキュメントの調査に多くの時間を費やしました。しかし、この 2 年間で多くのことが変わりました。現在では、エージェント ワークフローが私のワークフローの重要な部分を占めており、実装作業の多くを LLM に委任して、詳細ではなく意思決定に重点を置くことができます (ただし、まだレビューする必要があります)。
私は両方のコンテストの実装エージェントとして、Claude Code 内の Opus を使用しました。クロードは、基本的にすべてのトレーニング スクリプト、ビーム検索バリアント、TPU ポート、送信の自動化、および後処理ツールを作成しました。私は引き続きコードをレビューし、出力を確認し、実験レベルの決定を下しました。私はそれを導き、誤った道に進みすぎないようにすることに重点を置きました (これは頻繁に行われました)。私はそれに、どのアイデアを試すべきか、どの領域を研究すべきか、行き止まりのモデルの反復をいつやめるべきか、いつやめるべきかを伝えました。

ピボットなど
したがって、この投稿は 2 つのことについてです。 1 つは Cayley グラフの操作に関するもので、学習された距離ヒューリスティック、ワイド ビーム検索、対称性、およびパスの後処理を組み合わせて、DFS/BFS で検索するには大きすぎるパズル グラフを解決する方法です。 2 つ目は、エージェントのワークフローに関するものです。エージェントが実装を処理し、人間の役割が実験計画、プラトー検出、検索戦略に移行するときに、研究ループを実行するのが実際にどのようなものになるかです。
進捗状況の代理として、さまざまな方法で到達したスコアを共有します。リーダーボードには 1 つの数字 (固定された一連のスクランブルでの合計移動数) が表示されますが、重要なのは解決策の背後にあるロジックです。スコアの向上は、より短いパスを見つけることができたことを示しているため、リーダーボードの上位は検索の問題に対するより深い理解を反映しています。
ケイリー グラフとは何ですか?なぜそれを解くのですか?
まずは、ケイリー グラフとは何か、そしてそれをどのように操作できるのかについて説明します。
パズル（ルービックキューブのような）があります。可能な動き (ジェネレーター) は、小さな一連の操作 (立方体の面を回転させるなど) です。パズルには解決済みの状態があり、一連の動きを適用することでパズルを解くことができます。目標は、パズルを解決した状態に戻す一連の動きを見つけることです。
Cayley グラフでは、要素 (パズルの到達可能なすべての構成) ごとに 1 つの頂点があり、単一のジェネレーターが一方を他方に変換するたびに 2 つの頂点の間にエッジがあります。解決された状態は恒等頂点です。スクランブルは別の頂点です。パズルを解くことは、スクランブルからアイデンティティに戻る道を見つけることであり、それをうまく解くこと (私たちが気にするスコア) は、短い道を見つけることです。可能な最短のパスが、スクランブルの実際の解決までの距離です。すべての s にわたるそのような最大の距離

クランブルズはグラフの直径であり、キューバーの間では「神の数」として知られています。
ケイリーの定理によれば、すべての有限群は順列の群であるため、順列パズルは有限群理論一般を具体的に扱うものです。
ケイリー グラフは、グループを幾何学的オブジェクトとして表現する標準的な方法の 1 つです。頂点は要素、エッジはジェネレーターであり、グラフ メトリックはグループ内の距離の尺度になります。これが、これらのコンテストが離散最適化と幾何群理論を組み合わせた理由です。基本的な質問は本当に難しいです。2010 年まで 3×3×3 立方体の神の数が 20 であることを誰も証明できませんでしたし、メガミンクスの値は不明です。
これらのグラフは書き出すには大きすぎます。 3×3×3 立方体には約 4.3 × 10^19 の状態があります。メガミンクスにはおよそ 10^68 があり、天文学的な数字になります。実際には、グラフを保存しようとすることさえありません。1 つの状態を保持し、その近傍をその場で生成します。グラフはジェネレーターによって暗黙的に定義されます。
megaminx では、解決済みからの幅優先は数手の間は問題ありません。状態の数は 1、24、408、6208、90144、1.28M で、深さ 6 では最大 1800 万の状態になります。ただし、深さ 7 は最大 250M で、実際のスクランブルの深さは数十手です。
したがって、解決までの真の距離を計算することはできません。学習されたヒューリスティックを使用してそれを推定し、それに対して検索します。問題全体は 2 つの質問に要約されます。見積もりがどの程度適切か、もう 1 つはビームの幅がどれだけ許容できるかです。そして、残念な結果の 1 つは、最良の結果を達成するには、多くの場合、可能な限り最大のビームを TPU に詰め込むことを意味することです。
CayleyPy は、従来キューバーが使用していた手作りのソルバーの代わりに学習されたヒューリスティックとビーム検索を使用して、ケイリー グラフ内の短いパスを見つけるこの一連の作業の包括的なものです。それは一連の Kaggle コンペティションから独立しました。

オープンソースの CayleyPy ライブラリに実装された別のグラフ。
メインのパズルは IHES ピクチャー キューブです。すべての面が異なる 3×3×3 です。このパズルはルービック キューブよりも複雑で、ピースの向きや中心も重要です。できるだけ短く解決しなければならない 1003 のスクランブルがあります。この名前は、群の研究をケイリーグラフの幾何学に変えたミハイル・グロモフがキャリアを積んだパリ郊外の高級科学研究所にちなんで名付けられました。
メガミンクスは立方体ではなく、12 個の五角形の面が 5 分の 1 回転する十二面体です。私のソルバーは、24 個のジェネレーター (12 面 × 2 方向) を備えた長さ 120 のベクトルとして状態をエンコードし、競合では 1001 個のスクランブルが提供されます。
Megaminx は 2 つの理由から興味深い挑戦です。まず、信頼できる厳密なアルゴリズムがないため (コシエンバ スタイルの 2 フェーズ ソルバーを使用できる小規模なパズルとは異なります)、学習されたヒューリスティック検索が基本的に唯一の実行可能なアプローチです。第 2 に、合計スコアは 1,000 回のスクランブルにわたって計算されるため、小さな改善であっても目に見える改善が得られる可能性があります。
私にはそこそこの計算リソースがありました。プライマリ デバイスとして 16 GB RTX 4090 ラップトップ、長時間のジョブ用の GCP L4 マシン、並列探索用の無料の Kaggle GPU と TPU カーネル、そしてその後、より大きな TPU カーネルにアクセスできるようになりました。 Google Developer Expert のおかげで、無料の GCP クレジットを取得しました。これは、GPU での重いジョブの実行と TPU でのワイドビーム検索の実行に役立ちました。
どちらのコンテストも同じアプローチに従っており、より良く、より速く、より効率的に機能させることに重点が置かれています。
解決された状態からのランダム ウォークで、解決までの距離のニューラル推定値である V(s) を学習させます。
V を使用してビーム検索を実行して子をランク付けし、最高の B を維持します

それぞれのステップ。
さまざまなビーム方向、対称性、検索構成など、無相関な方法で各パズルを解きます。
完成したパスを後処理して短くします。
すべてのソリューションを最小マージし、パズルごとの最短パスを維持します。多くの場合、さまざまな人々が作成したソリューションを統合することが、スコアを大幅に向上させる鍵となります。
このアプローチは簡単ですが、各ステージが複雑になり、実行にかかる計算コストが高くなる可能性があります。検索ループの各段階におけるクロードの役割は次のとおりです。
この投稿の残りの部分は、両方のコンテストのソリューションを改善するための私のアプローチについて説明します。
2 つのコンテストを通じて、最小限の介入で研究ループを実行できる足場を構築しました。新しいセッションの開始時のコンテキストのリセットに耐え、長時間実行されているジョブを追跡し、作業状態を 1 か所に保持できるソリューションが必要でした。
マークダウン状態ファイル。 IDEAS_CATALOG.md 、 EXPERIMENTS.md 、 HANDOFF.md 。 Claude Code セッションは、セッションが長すぎる場合、またはセッションを終了するとコンテキストを失うため、重要なものはすべてこれらのファイルの 1 つに書き込まれます。
繰り返しのタスクにはスラッシュコマンドを使用します。最初に、提出物を Kaggle にすばやくアップロードするためのコマンドを作成しました。次に、ジョブのステータスを確認したり、送信をマージしたりするコマンドを追加しました。
バックグラウンド タスクと長時間のジョブの監視。多くの場合、1 時間にわたる実行 (場合によっては 24 時間以上) を開始する必要があり、失敗したりスタックしたりしないように監視する必要がありましたが、クロードはそれを管理することができました。
私のお気に入りの部分の 1 つは、「GCP GPU でこのトレーニング スクリプトを実行し、結果を報告する」のようなコマンドを実行することです。その後、クロードは VM を起動し、ジョブを実行し、結果を表示できます。とても時間を節約できました。
ここでクリーンなコードとエージェント設定を公開しました。私の作業リポジトリは混乱しているため非公開です。
IHES — クリエイティ

コアソリューションのNG化と磨き上げ
IHES Picture Cube では、ステートは 72 個のステッカー (8 コーナー × 3 + 12 エッジ × 2 + 6 センター × 4) の順列であり、18 個のジェネレーターは 3 軸ファミリー × 3 レイヤー インデックス × 2 サインとして構造化されています。コシエンバ 2 フェーズ ソルバー (3x3x3 の立方体を解くための古典的なコンピューター アルゴリズム) を適応させたところ、42718 の手が得られました。比較として、3×3×3 では神の数字は 20 であるため、「すべてのパズルで最適な」ソルバーは約 20,000 点のスコアを獲得します。
Claude ライブラリは、小規模な MLP 距離ヒューリスティックとビーム検索を提供します。 Claude はベースライン アーキテクチャ (シーケンシャル MLP) を分析し、残差ブロック、エンベディング エンコーダー、高速トレーニング レシピ (bf16、torch.compile 、融合 AdamW、より大きなバッチ、3 倍のスループット)、より幅広のビーム、および BFS-d5 後処理を追加して書き直しました。これによりスコアは 27106 に改善されました。
そしてこの時点で、私はすでにエージェント使用の限界に気づいていました。MSE 損失は改善されず、数回試行した後、エージェントは同じモデルのバリエーションをトレーニングし、それらをアンサンブルし始めました (スコアの迅速な向上を目指す本物の Kaggler のように!)。
同じアーキテクチャでの反復をやめて、より良い単一モデルを追求するようにクロードに明示的に伝える必要がありました。特定のトピックについて詳細な調査を行うよう依頼したところ、関連する著作物が複数見つかりました。次善のアイデアは、モデルの改善ではなく、ビーム検索の改善であることが判明しました。この時点で、ビーム幅 8K でビーム検索を実行していましたが、GPU メモリはすでにフルに使用されていました。 Claude は複数の最適化 ( torch.topk による優先キューフリーのトップ K、ステップ間で再利用される連続ビーム バッファーなど) を実装しました。これにより、ビーム幅を 65K まで増やすことができ、スコアが 2k ポイント改善されました。これは、ビーム探索幅がより多くの影響を与えることが多いことを証明しました

モデル アーキテクチャよりもアリの役割。
前述したように、モデルは通常、解決された状態から開始するランダム ウォークでトレーニングされ、ラベルはウォークの深さになります。ただし、ラベルを生成したランダム ウォークが最適ではない可能性があるため、これらのラベルは実際の距離の上限にすぎません。 Bellman 自己ブートストラップ ( target(s) = 1 + min_a target_net(apply(s, a)) ) は、そのバイアスを軽減します。これを追加すると、いくつかの非常に難しいパズルを解くことができました。
NISS も顕著な改善でした。クロードに Speedcubing フォーラムのマイニングを依頼したところ、NISS (Normal-Inverse Scramble Switch) に遭遇しました。スクランブルが与えられた場合、逆問題を解き、結果として得られるパスを逆にして、それぞれの動きを逆転させます。予想される長さは同じですが、ハッシュ シードとタイブレークが異なるため、検索は方向的に無相関になります。実装は簡単で、ベルマンと合わせてスコアは24068に達しました。
これはキューブ コミュニティでは古いアイデアですが、検索できるとは知りませんでした。
次の改善は段階的でした。最も難しいパズルに対して幅広のビーム検索を実行し、2M 幅で実行するようにビーム検索 (int8 状態エンコード) を最適化し、Q 関数蒸留 (Vlad Kuznetsov の記述より) によりスコアが 23322 に改善されました。
Q関数蒸留は非常に興味深いトリックです

[切り捨てられた]

## Original Extract

Two experiences from two CayleyPy Kaggle competitions: how learned-heuristic beam search solves enormous Cayley graphs, and how Claude Code simplifies and speeds up the engineering loop.

Cayley graph search with Claude Code: what puzzle competitions look like in 2026 – Andrey Lukyanenko
Skip to main content
Andrey Lukyanenko
Projects
Cayley graph search with Claude Code: what puzzle competitions look like in 2026
Cayley graph search with Claude Code: what puzzle competitions look like in 2026
Over the last couple of months, I worked on two Kaggle competitions in the CayleyPy series: IHES Picture Cube and Megaminx . Both are combinatorial puzzle solvers, and they can be reformulated as a search in an enormous implicit Cayley graph: each puzzle state is a vertex, each possible move is an edge, and solving means finding a short path back to the identity. This is a long term-activity: we have been researching these puzzles for a couple of years and have published several papers . This year we took another step toward solving them.
I remember working on this in 2024. At that time, all code was written by hand, and I spent a lot of time debugging it and digging into documentation. But much has changed in these two years: nowadays, agentic workflows are a significant part of my workflow, and I can delegate much of the implementation work to an LLM, focusing on decisions rather than specifics (even though I still have to review them).
I used Opus inside Claude Code as the implementation agent for both competitions. Claude wrote essentially all of the training scripts, beam-search variants, TPU ports, submission automation, and post-processing tools. I still reviewed the code, checked outputs, and made the experiment-level decisions. I focused on guiding it and preventing it from going too far down the wrong path (which it did quite often): I told it which ideas to try, which areas to research, when to stop iterating on a dead-end model, when to pivot, etc.
So this post is about two things. One is about working with Cayley graphs: how to combine learned distance heuristics, wide beam search, symmetry, and path post-processing to solve puzzle graphs far too large to search with DFS/BFS. The second one is about agentic workflow: what it actually looks like to run a research loop when an agent handles implementation and the human role shifts toward experimental design, plateau detection, and search strategy.
I share the scores reached by different methods as a proxy of progress. The leaderboard shows a single number (total moves over a fixed set of scrambles), but it is the logic behind the solutions that matters. Score improvement indicates we were able to find shorter paths, so higher places on the leaderboard reflect a better understanding of the search problem.
What’s a Cayley graph, and why would you solve one?
Let’s start with an explanation of what a Cayley graph is and how we can work with it.
We have a puzzle (like a Rubik’s cube). Possible moves (generators) are a small set of operations (like turning a face of the cube). The puzzle has a solved state, and we can scramble it by applying a sequence of moves. The goal is to find a sequence of moves that returns the puzzle to the solved state.
In a Cayley graph, we have one vertex per element (every reachable configuration of the puzzle), and an edge between two vertices whenever a single generator takes one to the other.. The solved state is the identity vertex; a scramble is some other vertex. Solving the puzzle is finding a path from the scramble back to the identity, and solving it well (the score that we care about) is finding a short one. The shortest possible path is the scramble’s true distance-to-solved; the largest such distance over all scrambles is the graph’s diameter, known to cubers as “God’s number”.
By Cayley’s theorem, every finite group is a group of permutations, so permutation puzzles are a concrete handle on finite group theory in general.
A Cayley graph is one of the standard ways of representing a group as a geometric object: vertices are elements, edges are generators, and the graph metric becomes a distance measure in the group. This is why these competitions combine discrete optimization with geometric group theory. The basic questions are genuinely hard — nobody proved the 3×3×3 cube’s God’s number is 20 until 2010, and the value for megaminx is unknown.
These graphs are far too big to write down. The 3×3×3 cube has about 4.3 × 10^19 states; the megaminx has roughly 10^68, astronomically more. In practice, you don’t even try to store the graph — you keep one state and generate its neighbors on the fly. The graph is defined implicitly, by its generators.
Breadth-first from solved is fine for a few moves — on megaminx. the numbers of states are 1, 24, 408, 6208, 90144, 1.28M, then ~18M states at depth six — but depth seven is ~250M, and a real scramble is dozens of moves deep.
So you can’t compute the true distance-to-solved; you estimate it with a learned heuristic and search against it. The whole problem boils down to two questions: how good is the estimate and how wide a beam can you afford. And one unfortunate consequence is that achieving the best results often means cramming the largest possible beam onto a TPU.
CayleyPy is the umbrella for this line of work on finding short paths in Cayley graphs, using learned heuristics and beam search instead of the hand-built solvers that cubers traditionally use. It spun off a series of Kaggle competitions, each a different graph, implemented in the open-source CayleyPy library.
The main puzzle is the IHES Picture Cube: a 3×3×3 with every face distinct. This puzzle is more complex than Rubik’s Cube, as the orientation of the pieces matters, even the centers. There are 1003 scrambles to solve as short as possible. It is named after the Institut des Hautes Études Scientifiques outside Paris, the institute where Mikhail Gromov, the one who turned the study of groups into the geometry of their Cayley graphs, has spent his career.
Megaminx is not a cube but a dodecahedron with twelve pentagonal faces turning in fifths of a full rotation. My solver encodes a state as a length-120 vector with 24 generators (12 faces × 2 directions), and the competition provides 1001 scrambles.
Megaminx is an interesting challenge for two reasons. First, there is no strict algorithm to fall back on (unlike the smaller puzzles where I could use a Kociemba-style two-phase solver), so learned-heuristic search is basically the only viable approach. Second, because the total score is calculated over a thousand scrambles, even small improvements can provide a visible improvement.
I had modest computational resources: a 16 GB RTX 4090 laptop as my primary device, a GCP L4 machine for long jobs, free Kaggle GPU and TPU kernels for parallel exploration, and later I got access to a larger TPU kernel. I got free GCP credits thanks to being a Google Developer Expert , and they helped us run heavy jobs on GPUs and wide-beam search on TPUs.
Both competitions follow the same approach, and the focus is on making it work better, faster, and more efficiently:
Train V(s), a neural estimate of distance-to-solved, on random walks from the solved state.
Run beam search, using V to rank children and keep the best B at each step.
Solve each puzzle in decorrelated ways: different beam directions, symmetries, and search configurations.
Post-process the completed paths to shorten them.
Min-merge every solution, keeping the shortest path per puzzle. Merging solutions created by different people is often a key to significantly improving the score.
This approach is straightforward, but each stage can be complex and computationally expensive to run. Here is what the Claude role was in each stage of the search loop:
The rest of the post is about my approaches to improving the solutions for both competitions.
Over the course of the two competitions, I built a scaffolding that let me run the research loop with minimal intervention. I needed a solution that could survive context resets when starting new sessions, track long-running jobs, and keep work state in a single place.
Markdown state files. IDEAS_CATALOG.md , EXPERIMENTS.md , HANDOFF.md . Claude Code session lose context when the sessions is too long or when you end it, so anything important was written down into one of these files.
Slash commands for repeated tasks. At first, I created commands to quickly upload submissions to Kaggle; then added commands to check job status, merge submissions, and more.
Background tasks and Monitor for long jobs. I often had to launch many hour-long runs (sometimes more than 24h), and needed to monitor them to make sure they don’t fail or get stuck - Claude was able to manage it.
One of my favorite parts was giving a command like “run this training script on GCP GPU and report the results”; Claude then could spin up a VM, run the job, and show the results. It saved me a lot of time.
I have published the clean code and agent settings here - my working repo is private as it is a mess.
IHES — creating and polishing the core solution
In the IHES Picture Cube, a state is a permutation of 72 stickers (8 corners × 3 + 12 edges × 2 + 6 centers × 4), with 18 generators structured as 3 axis families × 3 layer indices × 2 signs. An adapted Kociemba two-phase solver (a classical computer algorithm for solving a 3x3x3 cube) gave 42718 moves. For a comparison: on the 3×3×3, God’s number is 20, so an “optimal on every puzzle” solver would score ~20k.
The Claude library provides a small MLP distance heuristic and a beam search. Claude analyzed the baseline architecture (a sequential MLP) and rewrote it, adding residual blocks, an embedding encoder, a faster training recipe (bf16, torch.compile , fused AdamW, larger batches, 3× throughput), a wider beam, and BFS-d5 post-processing. This improved the score to 27106.
And at this point, I had already noticed the limitations of using agents: the MSE loss wasn’t improving, and after several attempts, the agent started training variations of the same model and ensembling them (just like a real Kaggler aiming for fast score improvement!).
I had to explicitly tell Claude to stop iterating on the same architecture and go after a better single model. I asked it to do deep research on specific topics, and we found multiple relevant works. The next best idea turned out to be not a model improvement but a beam-search improvement. At that point, I was running a beam search with a beam width of 8K, and the GPU memory was already fully used. Claude implemented multiple optimizations (priority-queue-free top-K via torch.topk , contiguous beam buffers reused across steps, and others), which allowed to increase the beam width to 65K, which improved the score by 2k point. This proved that beam search width often plays a more important role than the model architecture.
As mentioned before, models are usually trained on random walks starting from the solved state, and the labels are the walk depth. But these labels are merely an upper bound on true distance, because the random walk that produced a label might not be optimal . Bellman self-bootstrap ( target(s) = 1 + min_a target_net(apply(s, a)) ) reduces that bias. Adding it helped solve several very hard puzzles.
NISS was another noticeable improvement. I asked Claude to mine the speedcubing forums, and it encountered NISS - Normal-Inverse Scramble Switch. Given a scramble, you solve the inverse problem, then reverse the resulting path and invert each move. The expected length is the same, but we have different hash seeds and tie-breaking, so the search is directionally decorrelated. It was simple to implement, and together with Bellman, the score reached 24068.
This is an old idea in the cubing community, but I didn’t even know I could search for it.
The next improvements were incremental: running wider beam search for the hardest puzzles, optimizing beam search (int8 state encoding) to run at 2M width, and Q-function distillation (from Vlad Kuznetsov’s writeup) improved the score to 23322.
Q-function distillation is a very interesting tric

[truncated]
