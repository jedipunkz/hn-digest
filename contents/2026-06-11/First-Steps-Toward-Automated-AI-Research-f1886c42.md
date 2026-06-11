---
source: "https://www.recursive.com/articles/first-steps-toward-automated-ai-research"
hn_url: "https://news.ycombinator.com/item?id=48491593"
title: "First Steps Toward Automated AI Research"
article_title: "First Steps Toward Automated AI Research - Recursive"
author: "praccu"
captured_at: "2026-06-11T16:27:29Z"
capture_tool: "hn-digest"
hn_id: 48491593
score: 2
comments: 0
posted_at: "2026-06-11T15:21:48Z"
tags:
  - hacker-news
  - translated
---

# First Steps Toward Automated AI Research

- HN: [48491593](https://news.ycombinator.com/item?id=48491593)
- Source: [www.recursive.com](https://www.recursive.com/articles/first-steps-toward-automated-ai-research)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T15:21:48Z

## Translation

タイトル: 自動化された AI 研究に向けた第一歩
記事のタイトル: 自動化された AI 研究に向けた第一歩 - 再帰的
説明: Recursive の自動 AI 研究システムによるモデル トレーニングと GPU カーネル ベンチマークの初期の結果

記事本文:
自動化された AI 研究に向けた最初のステップ - 自動化された AI 研究に向けた再帰的な最初のステップ
Recursive の自動 AI 研究システムによるモデル トレーニングと GPU カーネル ベンチマークの初期結果
本日、私たちは Recursive の自動 AI 研究システムからの初期の結果をリリースします。このシステムは、固定予算の言語モデル トレーニング、小規模モデルのトレーニング速度、GPU カーネルの最適化という 3 つのベンチマークで最先端の結果を達成しています。
このシステムは、目標目的に向けた研究ループを自動化します。つまり、アイデアの提案、実装、実験の実行、結果の検証、そして学んだことを次の実験の選択に使用します。長い期間にわたって多くの研究スレッドを実行し、以前の実験からの有用なコンテキストを保持し、有望なブランチを結合し、パフォーマンスの向上を実際の進歩として扱う前に、結果を報酬ハックと差異の検証に掛けます。これは、オープンエンド アルゴリズムの原則を拡張および活用するように設計されており、私たちのチームや他のチームによる以前の研究からのアイデアを再帰的に自己改善する AI に構築しています。
私たちは、実用的な重要性と緊密なフィードバック ループの両方を考慮して選択されたベンチマークでシステムをテストしました。彼らは、AI の進歩の 3 つの中心的な手段であるトレーニング アルゴリズムの改善、トレーニングの高速化、ハードウェアのより効率的な使用を強調しています。また、明確な指標、比較的低い分散、報酬ハッキングに対して強化できる評価機能を備えているため、自動リサーチにも適しています。
私たちは、他の人がシステムの出力を検査して構築できるように、これらの実行からのアーティファクトをオープンソース化しています。
ケーススタディ 1: NanoChat 自動リサーチ
Andrej Karpathy の NanoChat 自動リサーチ リポジトリは、自動リサーチ システムの人気の出発点です。タスクは、最小の検証損失、測定値まで小さな言語モデルをトレーニングすることです。

単一の GPU で固定の 5 分間の予算内で、ビット/バイト (BPB) 単位で処理されます。実験は速く、分散は低く、報酬ハッキングは比較的簡単に検出できるため、これは私たちのシステムの自然なテストです。
おそらくこれらの理由から、この設定を中心に公的協力の取り組みがすでに形成されています。 autoresearch@home は、元の設定を、数十人の人間と数百人のエージェントが共同してパフォーマンスを向上させる共同設定に拡張します。これにより、カルパシーの一晩のランニングよりも強力な比較ポイントが得られます。私たちは、人間とエージェントのコミュニティ全体によって生成されたソリューションをシステムが改善できるかどうかをテストしたいと考えていました。
私たちのシステムは、Autoresearch コードが開始するのと同じ初期シード ソリューションから開始します。最初は NVIDIA H100 GPU で検索し、次に公開された結果と公平に比較​​するために、発見されたソリューションを NVIDIA B200 GPU で実行するように転送しました。以前の最高の autoresearch@home ソリューションからマイナーな報酬ハックを削除し、10 個のランダム シードで評価した後、その平均パフォーマンスは 0.9372 BPB でした。私たちのシステムは、0.9109 BPB に達する解決策を見つけました。これは 0.0263 BPB の改善です。別の方法で測定すると、当社のソリューションは、最良の autoresearch@home ソリューションよりも約 1.3 倍少ないトレーニング時間で、Karpathy のオリジナルの夜間自動リサーチ BPB の品質に達します。
自動リサーチは、いくつかの重要な設計上の決定が組み込まれた、すでに最適化されたモデルから始まります。この目的のために、私たちはシステムがはるかに弱い出発点、つまり単純な初期実装 (AdamW を備えたバニラの Transformer) からでも改善できるかどうかをテストしました。私たちのシステムはモデルを 1.059 BPB から 0.9344 BPB (NVIDIA B200 GPU で評価) に改善し、autoresearch@home コミュニティによって作成された最良のソリューションを再び上回りました。これは必要ありません

基盤となるモデルは、autoresearch@home コミュニティによって使用または作成されたものを含む多くの公開技術を知っている可能性があるため、独立した再発見が実際に証明されていますが、検索プロセスがはるかに弱い出発点から競争力のあるトレーニング スタックを組み立てることができることを示しています。結果として得られたソリューションは、公開されている最良のソリューションとはいくつかの点で異なります。
私たちのシステムにはどのような変更が加えられましたか?最良の解決策は 1 つのトリックによって導かれるものではありません。これらは、アーキテクチャ、ショートコンテキストメモリ、補助損失、注意、オプティマイザの動作、重み減衰スケジュール、コンパイラ設定などへの変更を組み合わせました。
最大の利点の 1 つは、より豊富なショート コンテキスト メモリ メカニズムによるものです。ベースラインではすでに値の埋め込みが使用されています。私たちのシステムは、ハッシュされたバイグラムとトリグラムの埋め込みテーブルを使用してこのアイデアを拡張し、学習されたゲートを介して注意値のパスに混合しました。これにより、モデルは、低速の畳み込みや注意を必要とする代替手段による時間コストを支払うことなく、ローカルの n-gram 情報を使用する安価な方法を得ることができました。
これは、スパース性軸としてハッシュ テーブルを探索する DeepSeek Engram などの最近の研究につながります。私たちの設定では、ハッシュ テーブルは約 5,000 万のパラメーター モデルに 10 ～ 20 億のスパース パラメーターを追加できます。ほとんどのエントリは特定のバッチで非アクティブであり、ルックアップは安価です。同様のハッシュ テーブルと N グラムのアイデアは、NanoGPT Speedrun の上位投稿にも表示されます。このシステムは、ハッシュ化されたバイグラムとトライグラムの埋め込みを、レイヤーごとに異なるハッシュを使用して、複数のレイヤーにわたるアテンション値ベクトルに注入することで、固定予算設定にこの一連のアイデアを適応させ、繰り返しの衝突を減らしました。私たちは、この正確な亜種を使用した以前の研究を知りません。
以下の展開可能なボックスには、システムのソリューションから選択された技術的な詳細が含まれています。手動で検査しました

これらの出力を分析し、報酬ハックのテクニックとスクリーニングを理解するために AI 支援分析を使用しました。私たちが専門家ではない場合、カーネルの最適化でエラーを見逃している可能性がありますが、それも重要な点の一部です。ここで提示されたアイデアは、私たちの以前の専門知識からではなく、システムから得られたものです。
開始ソリューションに埋め込まれた標準的なユニグラム値に加えて、私たちの最良のソリューションでは、モデルは各バイグラムとトライグラムを固定サイズのテーブルにハッシュし、学習された入力依存のヘッドごとのゲートを介してルックアップされたベクトルをアテンション値パスに混合します。これにより、古典的な N グラム モデルがトランスフォーマーの値ストリームに効果的に組み込まれます。
j の場合、列挙 (ve_layers) のlayer_i:
self.bigram_ves[ str (layer_i)] = nn.ModuleList([
nn.Embedding(self.bigram_table_size,half_kv_dim),
nn.Embedding(self.bigram_table_size,half_kv_dim),
])
self.bigram_hash_primes_per_layer[layer_i] = _decorr_bigram_primes[j]
j の場合、layer_i を列挙します (sorted (self.trigram_ve_layers)):
self.trigram_ves[ str (layer_i)] = nn.ModuleList([
nn.Embedding(self.trigram_table_size,half_kv_dim),
nn.Embedding(self.trigram_table_size,half_kv_dim),
])
self.trigram_hash_primes_per_layer[layer_i] = _decorr_trigram_primes[j]
v = v + Gate.unsqueeze(- 1 ) * ve ## 標準値の埋め込み
v = v + bg_gate.unsqueeze(- 1 ) * bigram_ve ## ルックアップ テーブルからの追加のバイグラム埋め込み
v = v + tg_gate.unsqueeze(- 1 ) * trigram_ve ## ルックアップ テーブルからの追加のトライグラム埋め込み このソリューションでは、異なるトランスフォーマー層に異なるハッシュ関数 (素数の素数ペアを素にする) も与えられます。
self.bigram_hash_primes_per_layer[layer_i] = _decorr_bigram_primes[j]
self.trigram_hash_primes_per_layer[layer_i] = _decorr_trigram_primes[j] つまり、衝突は依然として発生しますが、同じレベルで発生する可能性は低くなります。

y をレイヤー間で実行します。
バニラ Transformer を最適化する実行では、ハッシュ テーブルや二乗 ReLU MLP など、最良のソリューションと同じテクニックのいくつかが使用されました。しかし、トークンシフト、評価前の重み平均化、バイトレベルの機能埋め込みなど、別の (それでも同等に競争力のある) 最終スタックにも収束しました。これは、システムが他の実行で見つかったのと同じ発見を単に繰り返していたわけではないことを示唆しています。以下の展開可能なボックスは、バニラ Transformer の実行に固有のいくつかの変更を示しています。
バニラトランスフォーマーの最適化
バニラ Transformer ソリューションの変更の多くは、AdamW を Muon に置き換えたり、ハッシュ テーブルを追加したりするなど、私たちの最良のソリューション (Autoresearch の初期シード コードでシステムを起動することから生じたもの) にも反映されています。その他のいくつかの改善点は、最良のソリューションを生み出すメインの実行では現れませんでしたが、私たちにとって際立ったものでした。 1 つ目は因果的トークンシフトです。これは、次元ごとに学習された係数を使用して、前のトークンの注意投影 Q および K を現在のトークンにブレンドします。
B、T、C = x.size()
x_prev = F.pad(x[:, :- 1 , :], ( 0 , 0 , 1 , 0 ))
q = self.c_q(x + self.q_shift_beta * x_prev).view(B, T, self.n_head, self.head_dim)
k = self.c_k(x + self.k_shift_beta * x_prev).view(B, T, self.n_head, self.head_dim)
v = self.c_v(x).view(B, T, self.n_head, self.head_dim)
2 つ目は、トークン埋め込みの直後に挿入されるバイトレベルの機能のセットです。バイトレベルの特徴は、トークンがどのバイト (例: 個々の文字) で構成されているかに関する情報を表します。同様のバイトで構成されるトークンは、同様のバイトレベルの埋め込みを取得します。バイト特徴埋め込み行列は次のように構築されます。
結合 = torch.zeros(vocab_size, 769 )
範囲 (vocab_size) 内の token_id の場合:
raw_bytes = tokenizer_enc.decode_single_token_b

ytes(token_id) # 可変長
len (raw_bytes) > 0 の場合:
raw_bytes[:max_bytes] の b の場合: # max_bytes=16
結合[token_id, b] += 1.0 / len (raw_bytes) # [0:256] バイト頻度ヒストグラム
結合[token_id, 256 + raw_bytes[ 0 ]] = 1.0 # 最初のバイトのワンホット
combined[token_id, 512 + raw_bytes[- 1 ]] = 1.0 # 最後のバイトのワンホット
結合[token_id, 768 ] = len (raw_bytes) / max_bytes # 長さの特徴
torch.manual_seed( 1337 )
proj = torch.randn( 769 , embed_dim) * 0.01
init_emb = 結合 @ proj # [vocab, embed_dim]
self.embed = nn.Embedding(vocab_size, embed_dim)
self.embed.weight.data.copy_(init_emb) # INIT としてのみ使用されます。これらの埋め込みはトレーニング中に勾配降下法によって更新され、トークン埋め込みの後にバイグラムおよびトライグラム埋め込みとともに追加されます。
x_base = self.wte(idx) # トークンの埋め込み
Gates = self.embed_mixer(x_base) # 4 つのソース上のトークンごとのゲート [B,T,4]
x = x_base
x = x + ゲート[:,:, 0 : 1 ] * bi_raw # バイグラムハッシュ
x = x + ゲート[:,:, 1 : 2 ] * tri_raw # トライグラムハッシュ
x = x + ゲート[:,:, 2 : 3 ] * self.byte_embed.get_raw(idx) # バイト内容
x = x +gates[:,:, 3 : 4 ] * self.byte_boundary.get_raw(idx) # バイト境界
x = x + self.ssm_light(x)
x = self.embed_ctx_norm(x)
これらは、今回の実行でシステムに加えられた変更のほんの一部です。
NanoChat は、固定予算のトレーニングを改善するようシステムに要求することで、予算を意識した多くの複雑な改善点が発見されたことを示しています。次のテストは、何年にもわたって人間によるベンチマークの最適化を公に行った後でも、同じプロセスでまだ利益が得られるかどうかでした。私たちはこれを NanoGPT Speedrun でテストしました。NanoGPT Speedrun の最良のパブリック ソリューションは、2 年間にわたってコミュニティによって高度に最適化されてきました。
ケーススタディ 2: NanoGPT スピードラン
NanoGPT Speedrun も同様のタスクですが、大規模なコミュニティが存在するため、最先端の技術に勝つのははるかに困難です。

は 2 年以上にわたり、そのソリューションを最適化してきました。このベンチマークでは、一定の時間予算内でどの程度低い検証損失を達成できるかを問うのではなく、単一の HGX H100 8 GPU ノードを使用して、FineWeb テキスト データセット上で小規模な GPT スタイル モデルを固定検証損失 3.28 までどれだけ早くトレーニングできるかを問うています。
これは成熟したコミュニティの取り組みであり、これまでにリーダーボードに 83 件の人類記録を樹立する貢献があり、数百件の PR が提案されています。 2024 年半ば以降、主に手作業でエンジニアリングされた長い一連の提出により、トレーニング時間は約 45 分から 79.7 秒まで短縮されました。現在のソリューションは十分に最適化されているため、明らかな改善点はほとんど残されていません。
現在の主要なソリューションから始めて、私たちのシステムは、リーダーボードの検証損失有意性要件 (p < 0.01 で平均検証損失 ≤ 3.28) を満たしながら、トレーニング時間を 79.7 秒から 77.5 秒に短縮する一連の追加の最適化を見つけました。 1 これは、最近の人間の貢献と同等かそれ以上の改善です。
77.5 秒のソリューション: FP8 の注意、探索ノイズ、および慎重な埋め込み
以下の変更は、77.5 秒 (約 20 秒) を達成するために生み出され、組み合わされたイノベーションの例です。

[切り捨てられた]

## Original Extract

Early results from Recursive’s automated AI research system on model training and GPU kernel benchmarks

First Steps Toward Automated AI Research - Recursive First Steps Toward Automated AI Research
Early results from Recursive’s automated AI research system on model training and GPU kernel benchmarks
Today we are releasing early results from Recursive’s automated AI research system. Across three benchmarks, the system achieves state-of-the-art results: in fixed-budget language model training, small-model training speed, and GPU kernel optimization.
The system automates the research loop for a target objective: it proposes an idea, implements it, runs an experiment, validates the result, and uses what it learns to choose the next experiment. It runs many research threads over long horizons, keeps useful context from prior experiments, combines promising branches, and puts results through validation for reward hacks and variance before treating improved performance as real progress. It is designed to scale and harnesses principles of open-ended algorithms, building on ideas from previous work by our team and others into recursively self-improving AI.
We tested the system on benchmarks chosen for both practical importance and tight feedback loops. They stress three core levers of AI progress: better training algorithms, faster training, and more efficient use of hardware. They are also well suited to automated research because they have clear metrics, relatively low variance, and evaluators that can be hardened against reward hacks.
We are open-sourcing artifacts from these runs so others can inspect and build on the system’s outputs.
Case study 1: NanoChat Autoresearch
Andrej Karpathy’s NanoChat autoresearch repo is a popular starting point for automated research systems. The task is to train a small language model to the lowest validation loss, measured in bits per byte (BPB), within a fixed five-minute budget on a single GPU. It is a natural test of our system because experiments are fast, variance is low, and reward hacks are relatively easy to detect.
Perhaps for those reasons, a public collaborative effort has already formed around this setup. autoresearch@home extends the original setup into a collaborative setting where several dozens of humans and hundreds of their agents collectively improve performance. That gives us a stronger comparison point than Karpathy’s single overnight run. We wanted to test if our system could improve on solutions produced by an entire community of humans and agents.
Our system starts from the same initial seed solution the Autoresearch code starts from. We initially searched on NVIDIA H100 GPUs, then transferred the discovered solution to run on an NVIDIA B200 GPU for a fair comparison to public results. After removing minor reward hacks from the previous best autoresearch@home solution and evaluating it on 10 random seeds, its mean performance is 0.9372 BPB. Our system found a solution that reached 0.9109 BPB, a 0.0263 BPB improvement. Measured another way, our solution reaches the quality of Karpathy’s original overnight autoresearch BPB in roughly 1.3x less training time than the best autoresearch@home solution.
Autoresearch starts from an already optimized model with some non-trivial design decisions baked in. To this end, we tested whether our system could also make improvements from a much weaker starting point, a naive initial implementation (a vanilla Transformer with AdamW). Our system improved the model from 1.059 BPB to 0.9344 BPB (evaluated on an NVIDIA B200 GPU), again outperforming the best solution produced by the autoresearch@home community. This does not necessarily prove independent rediscovery, since the underlying models may know many public techniques including those used by or created by the autoresearch@home community, but it does show that the search process can assemble a competitive training stack from a much weaker starting point. The resulting solution also differed in several ways from the public best solution.
What modifications did our system come up with? The best solutions were not driven by one trick. They combined changes to architecture, short-context memory, auxiliary losses, attention, optimizer behavior, weight decay schedules, compiler settings, and more.
One of the biggest gains came from a richer short-context memory mechanism. The baseline already uses value embeddings; our system extended this idea with hashed bigram and trigram embedding tables, mixed into the attention value path through learned gates. This gave the model a cheap way to use local n-gram information without paying the time cost of slower convolutional or attention-heavy alternatives.
This connects to recent work such as DeepSeek Engram , which explores hash tables as a sparsity axis. In our setting, hash tables can add 1-2 billion sparse parameters to a roughly 50M parameter model: most entries are inactive on any given batch, and lookup is cheap. Similar hash-table and n-gram ideas also appear in top NanoGPT Speedrun submissions. The system adapted this family of ideas to the fixed-budget setting by injecting hashed bigram and trigram embeddings into attention value vectors across multiple layers, with different hashes per layer to reduce repeated collisions. We are not aware of prior work using this exact variant.
The expandable boxes below include selected technical details from the system’s solutions. We manually inspected these outputs and used AI-assisted analysis to understand the techniques and screen for reward hacks. We may still have missed errors in kernel optimization where we are not specialists, but that is also part of the point: the ideas presented here came from the system, not from our prior expertise.
On top of the standard unigram value embedding in the starting solution, in our best solution the model hashes each bigram and trigram into fixed-size tables and mixes the looked-up vectors into the attention value path through learned, input-dependent per-head gates — effectively folding a classic n-gram model into the transformer's value stream.
for j, layer_i in enumerate (ve_layers):
self.bigram_ves[ str (layer_i)] = nn.ModuleList([
nn.Embedding(self.bigram_table_size, half_kv_dim),
nn.Embedding(self.bigram_table_size, half_kv_dim),
])
self.bigram_hash_primes_per_layer[layer_i] = _decorr_bigram_primes[j]
for j, layer_i in enumerate ( sorted (self.trigram_ve_layers)):
self.trigram_ves[ str (layer_i)] = nn.ModuleList([
nn.Embedding(self.trigram_table_size, half_kv_dim),
nn.Embedding(self.trigram_table_size, half_kv_dim),
])
self.trigram_hash_primes_per_layer[layer_i] = _decorr_trigram_primes[j]
v = v + gate.unsqueeze(- 1 ) * ve ## standard value embedding
v = v + bg_gate.unsqueeze(- 1 ) * bigram_ve ## additional bigram embedding from lookup table
v = v + tg_gate.unsqueeze(- 1 ) * trigram_ve ## additional trigram embedding from lookup table The solution also gives different transformer layers different hash functions (with disjoint hash prime pairs).
self.bigram_hash_primes_per_layer[layer_i] = _decorr_bigram_primes[j]
self.trigram_hash_primes_per_layer[layer_i] = _decorr_trigram_primes[j] That means collisions still happen, but they are less likely to happen in the same way across layers.
The run optimizing the vanilla Transformer used some of the same techniques as our best solution, including hash tables and squared-ReLU MLPs. But it also converged on a different (yet equally competitive) final stack, including token-shifting, weight averaging before eval, and byte-level feature embeddings. This suggests the system was not merely repeating the same discoveries it found in the other run. The expandable box below shows a few modifications unique to the vanilla Transformer run.
Optimizing a vanilla Transformer
Many of the changes in the vanilla Transformer solution also appear in our best solution (which came from starting our system with the Autoresearch initial seed code), such as replacing AdamW with Muon and adding hash tables. A few other improvements did not emerge in our main run that produced the best solution, yet stood out to us. The first is causal token shifting, which blends the previous token's attention projections Q and K into the current token's, with a learned coefficient per dimension.
B, T, C = x.size()
x_prev = F.pad(x[:, :- 1 , :], ( 0 , 0 , 1 , 0 ))
q = self.c_q(x + self.q_shift_beta * x_prev).view(B, T, self.n_head, self.head_dim)
k = self.c_k(x + self.k_shift_beta * x_prev).view(B, T, self.n_head, self.head_dim)
v = self.c_v(x).view(B, T, self.n_head, self.head_dim)
The second is a set of byte-level features injected right after the token embedding. The byte-level features represent information about what bytes (e.g., individual characters) tokens are composed of. Tokens consisting of similar bytes will get similar byte-level embeddings. The byte feature embedding matrix is built as follows:
combined = torch.zeros(vocab_size, 769 )
for token_id in range (vocab_size):
raw_bytes = tokenizer_enc.decode_single_token_bytes(token_id) # variable length
if len (raw_bytes) > 0 :
for b in raw_bytes[:max_bytes]: # max_bytes=16
combined[token_id, b] += 1.0 / len (raw_bytes) # [0:256] byte-frequency histogram
combined[token_id, 256 + raw_bytes[ 0 ]] = 1.0 # first byte one-hot
combined[token_id, 512 + raw_bytes[- 1 ]] = 1.0 # last byte one-hot
combined[token_id, 768 ] = len (raw_bytes) / max_bytes # length feature
torch.manual_seed( 1337 )
proj = torch.randn( 769 , embed_dim) * 0.01
init_emb = combined @ proj # [vocab, embed_dim]
self.embed = nn.Embedding(vocab_size, embed_dim)
self.embed.weight.data.copy_(init_emb) # used only as the INIT These embeddings are then updated by gradient descent during training, and added after the token embedding alongside the bigram and trigram embeddings:
x_base = self.wte(idx) # token embedding
gates = self.embed_mixer(x_base) # per-token gates over 4 sources [B,T,4]
x = x_base
x = x + gates[:,:, 0 : 1 ] * bi_raw # bigram hash
x = x + gates[:,:, 1 : 2 ] * tri_raw # trigram hash
x = x + gates[:,:, 2 : 3 ] * self.byte_embed.get_raw(idx) # byte-content
x = x + gates[:,:, 3 : 4 ] * self.byte_boundary.get_raw(idx) # byte-boundary
x = x + self.ssm_light(x)
x = self.embed_ctx_norm(x)
These are just a few of the changes our system made in this run.
NanoChat shows how asking our system to improve fixed-budget training led to the discovery of many compounding, budget-aware improvements. The next test was whether the same process could still find gains after years of public human optimization on a benchmark. We tested that on NanoGPT Speedrun, whose best public solution has been highly optimized by the community over two years.
Case study 2: NanoGPT Speedrun
NanoGPT Speedrun is a similar task, yet it's much harder to beat the state of the art because a large community has been optimizing solutions for it for over two years. Instead of asking how low of a validation loss can be achieved in a fixed time budget, the benchmark asks how quickly a small GPT-style model can be trained to a fixed validation loss of 3.28 on the FineWeb text dataset, using a single HGX H100 8-GPU node.
This is a mature community effort, with 83 human record-setting contributions to the leaderboard so far and hundreds of proposed PRs. Since mid-2024, the training time has been pushed from roughly 45 minutes down to 79.7 seconds through a long sequence of primarily hand-engineered submissions. Given that the current solution is so well optimized, there are few obvious improvements left.
Starting from the current leading solution, our system found a set of additional optimizations that reduced training time from 79.7 seconds to 77.5 seconds while still meeting the leaderboard’s validation-loss significance requirement (mean validation loss ≤ 3.28 at p < 0.01). 1 This is a similar or larger improvement than recent human contributions.
The 77.5 s solution: FP8 attention, exploration noise, and cautious embeddings
The changes below are examples of innovations it created and combined to reach 77.5 s (in ~ 20

[truncated]
