---
source: "https://fengrru.me/blog/ctm-theory-review.md"
hn_url: "https://news.ycombinator.com/item?id=49222937"
title: "Show HN: A Source-Level Review of Sakana AI's Continuous Thought Machine"
article_title: "Rethinking the Continuous Thought Machine: Four Optimization Directions and a Six-Direction Compatibility Audit"
author: "Fengrru"
captured_at: "2026-08-08T16:21:04Z"
capture_tool: "hn-digest"
hn_id: 49222937
score: 1
comments: 0
posted_at: "2026-08-08T15:48:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A Source-Level Review of Sakana AI's Continuous Thought Machine

- HN: [49222937](https://news.ycombinator.com/item?id=49222937)
- Source: [fengrru.me](https://fengrru.me/blog/ctm-theory-review.md)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T15:48:08Z

## Translation

タイトル: Show HN: Sakana AI の継続的思考マシンのソースレベルのレビュー
記事のタイトル: 継続的思考マシンの再考: 4 つの最適化の方向と 6 方向の互換性監査

記事本文:
Fengrru ホーム ブログ 音楽 継続的思考マシンの再考: 4 つの最適化の方向と 6 方向の互換性監査
Sakana AI の Continuous Thought Machine (CTM) は、内部時間軸に沿って神経活動を展開し、ニューロンレベルの時間処理と神経同期を表現として扱います。この記事では、ボトムアップの最適化提案とトップダウンの理論エンジニアリング監査という 2 つの補完的な分析を 1 つの包括的なレビューに結合します。
パート 1 では、予測コーディングと二重経路視覚認知理論を、二重経路予測コーディング視覚システム (v_predictive / v_dual / v_engine) での私のエンジニアリング経験と合わせて活用し、CTM の 4 つの具体的な最適化の方向性を提案します。それは、予測誤差駆動型インナーループ、腹側 - 背側二重経路アーキテクチャ、トレーニングの安定性 (マイクロパラメーター化)、および確実性ベースの早期停止です。
パート 2 では、逆の視点を採用しています。つまり、v_predictive / v_dual / v_engine アーキテクチャ パラダイムを 6 方向にわたって CTM にマッピングするトップダウンの理論的提案が与えられ、CTM の実際のソース コードに対してあらゆる仮定が監査されます。何が直接接続され、何が修正が必要かを正確に特定し、修正された実装ロードマップを作成します。
あらゆる方向性は、メカニズムの原理、数学的形式化、models/ctm.py (1 ～ 605 行目) および models/modules.py (1 ～ 693 行目) の正確なソースコード変更ポイント、設計理論的根拠、予想される利点とリスク、および実験プロトコルを使用した段階的なロードマップに基づいています。
パート 1: 4 つの最適化の方向 — CTM コードの事実に基づくボトムアップ
1. CTM の要約: 時間内に思考を展開するマシン
内部時間軸 : モデルには i から切り離された内部ティック (論文では (T) で示される) があります。

データを入力し、「思考」をプロセスとして展開します。データは一度フィードフォワードされます。後続の反復は純粋に内部状態に基づいて実行されます。
ニューロンレベルの時間処理 (NLM) : 各ニューロンは、自身の過去 (M) ステップの入力履歴 (トレース) を処理する独立した重みを持ち、きめ細かい時間ダイナミクスを可能にします。コードでは、これは SuperLinear です: 重み形状 ((M, H, D)) (履歴の長さ (\times) 出力 dim (\times) ニューロン数)、einsum('BDM,MHD->BDH') を介して (D) 独立した線形マップを並列実行します (models/modules.py 行 146 ～ 236)。
表現としての同期: 時間の経過とともにニューロンペアの活動が同期する度合いが、出力およびアクションの表現として直接機能します。同期は、(O(1)) 回帰形式で実装される、ペアごとの活性化生成物の指数関数的に減衰する時間的蓄積です。
次の形式化は、models/ctm.py の前方 (527 ～ 603 行目) に対応します。テンソル形状:
入力 (x): ((B、C、H、W)) またはタスク依存の形状。
特徴 (kv): ((B, S, d_{\text{input}}))、ここで (S) はトークンの数です (バックボーン特徴マップは平坦化され、 kv_proj によって投影されます)。
History state_trace : ((B, D, M))、ここで、(D = d_{\text{model}}) はニューロン数、(M = \text{memory_length}) です。
アクティブ化 activate_state : ((B, D))、つまりアクティブ化後 (z_t)。
アクション同期 synchronisation_action : ((B, n_{\text{synch_action}}))。
出力同期 synchronisation_out : ((B, n_{\text{synch_out}}))。
予測: ((B, \text{out_dims}, T));確実性: ((B, 2, T))。
特徴抽出 (1 回、ループ外):
[
kv = W_{kv} \big( \text{ flatten}( \text{バックボーン}(x) + \text{PE}(\text{バックボーン}(x)) ) \big)
】
(O(1)) 同期再帰 ( compute_synchronisation 、行 202 ～ 267): (p_t = z_{t-1}^L \odot z_{t-1}^R) を選択したニューロのペアワイズ積とします。

n 個のペア (random-pairing の場合は要素ごと、 first-last / random の場合は外積の上三角形)。次に:
[
\alpha_t = r \alpha_{t-1} + p_t、\qquad \beta_t = r \beta_{t-1} + 1、\qquad S_t = \frac{\alpha_t}{\sqrt{\beta_t}}
】
ここで、 (r = \exp(-\text{decay_params})) は、decay_params が ([0, 15]) にクランプされているため (551 ～ 552 行目)、(r \in [e^{-15}, 1]) となります。この繰り返しは、有効サンプル数 (「同期度」の連続時間の尺度) によって正規化された、過去のペアごとの積の指数関数的に重み付けされた累積に相当します。最初の呼び出しでは、decay_alpha /decay_beta は None です。これらは、現在のペアごとの積とすべて 1 のベクトルで初期化されます (行 259 ～ 261)。後続のステップでは、すべての履歴内積を再計算せずに (O(1)) 再帰を使用します。
ループ本体 (各ステップ stepi 、560 ～ 590 行目):
[
q_t = W_q S_t^{\text{アクション}}, \quad o_t = \text{マルチヘッドアテンション}(q_t, kv, kv)
】
[
a_t = \text{シナプス}([o_t; z_{t-1}]), \quad A_t = [a_{t-M+1}, \dots, a_t]
】
[
z_t = \text{NLM}(A_t), \quad
S_t^{\text{out}} = \frac{\alpha'_t}{\sqrt{\beta'_t}}, \quad
\hat{y}_t = W_o S_t^{\text{out}}, \quad
c_t = 1 - H_n(\text{softmax}(\hat{y}_t))
】
一時的なメモ: ループ エントリで active_state = (z_{t-1}) (前の NLM ステップまたは初期値から)、シナプス + NLM 後に (z_t) に更新されます。したがって、compute_synchronisation(action) は (z_{t-1}) を使用し、compute_synchronisation(out) は (z_t) を使用します。
1.3 主要なサブモジュールのソースレベル分析
SynapseUNET は ([o_t; z_{t-1}]) (次元 d_input + D ) を受け入れ、(a_t \in \mathbb{R}^D) を出力します。構造:
初期投影: LazyLinear(widths[0]) → LayerNorm → SiLU、入力を最初の幅レベルにマッピングします。
ダウンサンプリング パス: num_blocks = Depth - Dropout の 1 ブロック → Linear(w[i] → w[i+1]) → LayerNorm → SiLU 、幅

out_dims (つまり (D)) から minimum_width (デフォルト 16) まで指数関数的に減衰します。各ブロックの出力はスキップ接続用に保存されます。
アップサンプリング パス: 逆の順序 (最も深い→最も浅い) で、各ブロック Dropout → Linear(w[i+1] → w[i]) → LayerNorm → SiLU を経て、対応するダウンサンプリング出力を追加し、そのレベルの LayerNorm を通過します。
MLP ではなく U-Net を使用する理由 : 深い U-N​​et は、マルチレベルの情報混合 (集団間の粗い相互作用ときめ細かいローカルな相互作用) を提供し、スキップ接続により深度での勾配パスの整合性が維持されます。紙の深さ = 1 (単一 GLU) は単純な混合層に分解されます。 Depth > 1 の方がうまく機能します (get_synapses の 415 ～ 435 行目でこれが説明されています)。
MuPC のエントリ ポイント: ダウン/アップ パスの各レベルの線形における重みスケーリング係数 (1/\sqrt{l+1}) を直接注入できます (方向 3 を参照)。
基本的には最後の次元 (履歴次元 (M)) に対する線形変換ですが、各ニューロン ((D) dim) には独自の異なる重みがあります。
Weight w1 : ((M, H, D)) — (M) 履歴値から (H) 出力、(D) 独立バージョンへのマッピング。
バイアス b1 : ((1, D, H)) — ニューロンごと、出力ごとのバイアス。
計算: einsum('BDM,MHD->BDH', x, w1) + b1 → shake(-1) / T 。
学習可能な温度 (T) ( modules.py 行 209): (T) で割ることで出力スケールを制御します。 1.0に初期化されました。 deep_nlms=True の場合、2 つの SuperLinear ブロックは GLU を介してチェーンされます (ctm.py 行 393 ～ 404)。
「ニューロンはスカラーである」制約: NLM 入力 ((B, D, M)) は、各ニューロンが単なるスカラー時系列であることを意味します。これは、方向 1 のエラー信号がニューロンごとに動作しなければならない理由を理解するための前提条件です。
ランダムペアリングはこの論文の最終戦略です (72 ～ 74 行目のコメント)。 n_random_pairing_self は、i-i セルフペアリングの数を制御します。0 では、セルフペアリングは

まれであるため、「スナップショット表現」の回復が困難になります (76 ～ 78 行目)。方向 1 のエラー定義では、二次関数的に増加するのではなく、sync dim = n_synch であるため、ランダム ペアリングに基づいて設計します。
compute_normalized_entropy ソフトマックスはロジットを計算し、エントロピーを計算し、上限 (\ln C) で除算して ([0,1]) に正規化します。多次元ロジット (マルチトークンシーケンスなど) の場合、非バッチディムに対して flatten(1).mean(-1) を実行します。落とし穴: クラス数 (C) が大きい場合 (ImageNet 1000 クラスなど)、ほとんどのクラス確率は 0 に近く、初期確実性は人為的に高くなります。正規化エントロピー (H_n = 1 - H/H_{\max}) は均一の場合は 0、ワンホットの場合は 1 ですが、高次元空間では均一の確率密度は非常に低いためです。実際には、「不確実な」状態であっても、少数のクラスに集中して確実性を誇張する傾向があります。ランダムなベースラインとタスクの難易度を組み合わせて確実性を評価します。生の ([0,1]) 範囲を単独で使用しないでください。
1.4 開ループ予測の構造的欠陥
CTM のループは開ループ外挿器です。各ステップ (z_t) は固定マッピングを介した履歴によって純粋に決定され、「モデル自体の動的予測」と「モデルの実際のアクティビティ」の間の一貫性を測定するメカニズムはありません。次の 4 つの問題が発生します。
内部監視なし: 唯一の監視は、(T) タイム ステップ (BPTT) を通じて逆伝播される最終タスク損失から得られます。勾配は U-Net の深さと NLM の非線形性によって繰り返し圧縮されます。深いタイム ステップでの実効勾配は非常に弱いです。
適応計算なし: 入力の難易度に関係なく、モデルは常に固定 (T) ステップを実行します。簡単なサンプルでは計算が無駄になり、難しいサンプルではさらに多くの計算が必要になる場合があります。
内部収束信号がありません。出力側には確実性が存在しますが、

ダイナミクス側 (状態が安定しているかどうか、予測が自己矛盾しているかどうか) にはメトリックがありません。モデルは、安定した内部表現を形成しているかどうかを「知る」ことができません。
ダイナミクスは解釈できません。同期はアクティビティの暗黙的な統計です。研究者は事後的に視覚化することしかできず、「モデルがどのステップで、何について驚いたか」を直接読み取ることはできません。
これら 4 つの点はそれぞれ、この記事の 4 つの最適化の方向性の動機となります。
2. 理論的背景: 予測コーディングとデュアルパスウェイ
神経科学財団 (Rao & Ballard、1999; Friston、2005): 各皮質層は次の層の活動の予測を維持します。予測誤差のみが上方に伝播します。学習により、予測誤差が最小限に抑えられます (生成モデルの証拠の下限、つまり自由エネルギー原理を最大化するのと同じです)。推論自体は「予測修正」の反復です。エラーが消えるまで、状態はエラー信号によって修正されます。その時点で、内部表現は生成モデルと自己矛盾がなくなります。
エンジニアリング実装 (v_predictive) の 3 つのコア コンポーネント:
予測コーディング層: (z = \text{ReLU}(\text{LN}(W_e x)))、予測 (\hat{x} = W_p z)、誤差 (e = x - \hat{x})。
エラー駆動の更新 ( UpdateRule ): (h \leftarrow h + \alpha \cdot \text{MLP}(e))。重要な設計: MLP はbias=Falseを使用します。これにより、(e = 0 \Rightarrow \text{MLP}(0) = 0 \Rightarrow h_{\text{new}} = h) が保証されます。つまり、誤差がゼロであれば厳密な固定小数点が得られます。収束は予測の自己一貫性と同等です。 (\alpha) は Softplus によってポジティブに制約されます。
反復推論 ( InferenceLoop ): 状態は上記のルールに従って反復されます。誤差基準がしきい値を下回ると早期に停止します。 BN 統計は推論ループ中にフリーズされます。
神経科学財団 (Ungerl)

eider & Mishkin、1982): 腹側経路 (「何を」) はオブジェクトのアイデンティティとセマンティクスを処理します。背側経路（「どこで/どのように」）は空間位置と運動制御を処理します。
エンジニアリング実装 (v_dual):
役割の分離: 腹側ストリームは凍結された重みを使用して意味論的な入力を受け取ります。背側ストリームは空間入力 (残差 + 2D 座標) を受け取りますが、意味論的なシンボル ID (実行時アサーションによって強制される) を受け取ることはありません。
方向認識バイアス ( RelativeDirectionalBias2D ): FH 領域トークンは不規則な形状で固定グリッドがないため、従来のグリッド バイアスが使用できなくなります。解決策: ペアごとのトークン距離 (対数バケット化) + 方向 (角度バケット化) を結合して量子化 → 学習可能なバイアス テーブルで検索し、方向情報を保存します。
クロスストリーム フュージョン ( DualPathwayFusion ): 双方向クロス アテンション + 学習可能なゲート (fused = \sigma(\alpha) v + (1 - \sigma(\alpha)) d)、ゲートは 0.5 (等しい重み) で初期化されます。 none モードでは、アブレーションのために融合モジュールを物理的に取り外します。
マイクロパラメータ化 MuPC : 深い予測ネットワークでは、勾配分散は深さとともに指数関数的に発散します。 Layer-(l) ウェイト スケーリング (1/\sqrt{l+1}) (浅い勾配の爆発を抑制)、エラー スケーリング (1/\sqrt{L-l+1}) (深い勾配の消失を緩和)。レイヤーごとの勾配の分散は互いに 10 倍以内に収束します。
BN凍結：d

[切り捨てられた]

## Original Extract

Fengrru Home Blog Music Rethinking the Continuous Thought Machine: Four Optimization Directions and a Six-Direction Compatibility Audit
The Continuous Thought Machine (CTM) by Sakana AI unfolds neural activity along an internal time axis, treating neuron-level temporal processing and neural synchronization as representations. This article combines two complementary analyses — a bottom-up optimization proposal and a top-down theory-engineering audit — into a single comprehensive review.
Part 1 draws on predictive coding and dual-pathway visual cognition theory, together with my engineering experience from the dual-pathway predictive coding vision system (v_predictive / v_dual / v_engine), to propose four concrete optimization directions for CTM: prediction-error-driven inner loops, ventral-dorsal dual-pathway architecture, training stability (micro-parameterization), and certainty-based early stopping.
Part 2 takes the opposite perspective: given a top-down theoretical proposal mapping the v_predictive / v_dual / v_engine architectural paradigm onto CTM across six directions, it audits every assumption against CTM’s actual source code. It identifies precisely what connects directly and what requires correction, producing a corrected implementation roadmap.
Every direction is grounded in mechanism principles, mathematical formalization, exact source-code modification points in models/ctm.py (lines 1–605) and models/modules.py (lines 1–693), design rationale, expected benefits and risks, and phased roadmaps with experimental protocols.
Part 1: Four Optimization Directions — Bottom-Up from CTM Code Facts
1. CTM Recap: A Machine That Unfolds Thought in Time
Internal time axis : the model has internal ticks (denoted (T) in the paper) decoupled from input data, allowing “thinking” to unfold as a process. Data is fed forward once; subsequent iterations run purely on internal state.
Neuron-Level Temporal Processing (NLM) : each neuron has independent weights processing its own past (M)-step input history (trace), enabling fine-grained temporal dynamics. In the code this is SuperLinear : weight shape ((M, H, D)) (history length (\times) output dim (\times) neuron count), executing (D) independent linear maps in parallel via einsum('BDM,MHD->BDH') ( models/modules.py lines 146–236).
Synchronization as representation : the degree to which neuron-pair activity synchronizes over time directly serves as output and action representations. Synchronization is an exponentially decaying temporal accumulation of pairwise activation products, implemented in (O(1)) recurrent form.
The following formalization corresponds to models/ctm.py forward (lines 527–603). Tensor shapes:
Input (x): ((B, C, H, W)) or task-dependent shape.
Features (kv): ((B, S, d_{\text{input}})), where (S) is the number of tokens (backbone feature map flattened then projected via kv_proj ).
History state_trace : ((B, D, M)), where (D = d_{\text{model}}) is neuron count and (M = \text{memory_length}).
Activation activated_state : ((B, D)), i.e. post-activation (z_t).
Action sync synchronisation_action : ((B, n_{\text{synch_action}})).
Output sync synchronisation_out : ((B, n_{\text{synch_out}})).
Predictions: ((B, \text{out_dims}, T)); certainties: ((B, 2, T)).
Feature extraction (once, outside loop):
[
kv = W_{kv} \big( \text{flatten}( \text{Backbone}(x) + \text{PE}(\text{Backbone}(x)) ) \big)
]
(O(1)) synchronization recurrence ( compute_synchronisation , lines 202–267): let (p_t = z_{t-1}^L \odot z_{t-1}^R) be the pairwise product of selected neuron pairs (elementwise for random-pairing , upper triangle of outer product for first-last / random ). Then:
[
\alpha_t = r \alpha_{t-1} + p_t, \qquad \beta_t = r \beta_{t-1} + 1, \qquad S_t = \frac{\alpha_t}{\sqrt{\beta_t}}
]
where (r = \exp(-\text{decay_params})) with decay_params clamped to ([0, 15]) (lines 551–552), so (r \in [e^{-15}, 1]). This recurrence is equivalent to an exponentially weighted accumulation of historical pairwise products normalized by effective sample count — a continuous-time measure of “degree of synchronization.” On the first call decay_alpha / decay_beta are None ; they are initialized with the current pairwise product and an all-ones vector (lines 259–261). Subsequent steps use (O(1)) recurrence without recomputing all historical dot products.
Loop body (each step stepi , lines 560–590):
[
q_t = W_q S_t^{\text{action}}, \quad o_t = \text{MultiheadAttention}(q_t, kv, kv)
]
[
a_t = \text{Synapses}([o_t; z_{t-1}]), \quad A_t = [a_{t-M+1}, \dots, a_t]
]
[
z_t = \text{NLM}(A_t), \quad
S_t^{\text{out}} = \frac{\alpha’_t}{\sqrt{\beta’_t}}, \quad
\hat{y}_t = W_o S_t^{\text{out}}, \quad
c_t = 1 - H_n(\text{softmax}(\hat{y}_t))
]
Temporal note : at loop entry activated_state = (z_{t-1}) (from the previous NLM step or initial value), updated to (z_t) after Synapses + NLM. Thus compute_synchronisation(action) uses (z_{t-1}) and compute_synchronisation(out) uses (z_t).
1.3 Key Submodule Source-Level Analysis
SynapseUNET accepts ([o_t; z_{t-1}]) (dimension d_input + D ) and outputs (a_t \in \mathbb{R}^D). Structure:
Initial projection : LazyLinear(widths[0]) → LayerNorm → SiLU , mapping input to the first width level.
Downsampling path : num_blocks = depth - 1 blocks of Dropout → Linear(w[i] → w[i+1]) → LayerNorm → SiLU , width decaying exponentially from out_dims (i.e., (D)) to minimum_width (default 16). Each block output is saved for skip connections.
Upsampling path : in reverse order (deepest → shallowest), each block Dropout → Linear(w[i+1] → w[i]) → LayerNorm → SiLU , then add the corresponding downsampling output, then pass through that level’s LayerNorm .
Why U-Net instead of MLP : deep U-Net provides multi-level information mixing (coarse cross-population interaction vs. fine-grained local interaction), and skip connections preserve gradient path integrity at depth. Paper depth = 1 (single GLU) degrades to a simple mixing layer; depth > 1 works better ( get_synapses lines 415–435 document this).
Entry point for MuPC : the weight scaling factor (1/\sqrt{l+1}) at each level’s Linear in the down/up paths can be directly injected (see Direction 3).
Essentially a linear transform over the last dimension (history dimension (M)), but each neuron ((D) dim) has its own distinct weights:
Weight w1 : ((M, H, D)) — mapping from (M) history values to (H) outputs, (D) independent versions.
Bias b1 : ((1, D, H)) — per-neuron, per-output bias.
Computation : einsum('BDM,MHD->BDH', x, w1) + b1 → squeeze(-1) / T .
Learnable temperature (T) ( modules.py line 209): dividing by (T) controls output scale; initialized to 1.0. When deep_nlms=True , two SuperLinear blocks are chained through GLU ( ctm.py lines 393–404).
The “neuron is a scalar” constraint : NLM input ((B, D, M)) means each neuron is merely a scalar time series — a prerequisite for understanding why Direction 1’s error signal must operate per-neuron.
random-pairing is the paper’s final strategy (lines 72–74 comment). n_random_pairing_self controls the number of i-i self-pairings — at 0, self-pairings are rare, making “snapshot representation” recovery difficult (lines 76–78). For Direction 1’s error definition, design on random-pairing since sync dim = n_synch rather than growing quadratically.
compute_normalized_entropy softmaxes logits, computes entropy, and divides by the upper bound (\ln C) to normalize to ([0,1]). For multi-dimensional logits (e.g., multi-token sequences), it does flatten(1).mean(-1) over non-batch dims. Pitfall: when the class count (C) is large (e.g., ImageNet 1000 classes), most class probabilities are near 0, and initial certainty is artificially high — because normalized entropy (H_n = 1 - H/H_{\max}) is 0 for uniform and 1 for one-hot, but uniform in high-dimensional space has extremely low probability density; in practice even “uncertain” states tend to concentrate on a few classes, inflating certainty. Evaluate certainty jointly with random baselines and task difficulty; never use the raw ([0,1]) range in isolation.
1.4 Structural Deficiencies of Open-Loop Prediction
CTM’s loop is an open-loop extrapolator : each step (z_t) is determined purely by history through a fixed mapping, with no mechanism to measure the consistency between “the model’s own dynamical prediction” and “the model’s actual activity.” Four problems arise:
No internal supervision : the only supervision comes from the final task loss, backpropagated through (T) time steps (BPTT). Gradients are repeatedly compressed by U-Net depth and NLM nonlinearity — effective gradients at deep time steps are extremely weak.
No adaptive computation : regardless of input difficulty, the model always runs a fixed (T) steps; easy samples waste computation, hard samples may need more.
No internal convergence signal : certainty exists on the output side, but the dynamics side (whether state is stable, whether predictions are self-consistent) has no metric — the model cannot “know” whether it has formed a stable internal representation.
Dynamics are uninterpretable : synchronization is an implicit statistic of activity; researchers can only post-hoc visualize, not directly read “at which step, about what, was the model surprised.”
These four points respectively motivate the four optimization directions in this article.
2. Theoretical Background: Predictive Coding and Dual Pathways
Neuroscience foundation (Rao & Ballard, 1999; Friston, 2005): each cortical layer maintains a prediction of the next layer’s activity; only prediction errors propagate upward; learning minimizes prediction error (equivalent to maximizing generative-model evidence lower bound, i.e., the free-energy principle). Inference itself is a “predict-correct” iteration: state is corrected by error signals until error vanishes — at which point the internal representation is self-consistent with the generative model.
Three core components from engineering implementation (v_predictive) :
Predictive coding layer : (z = \text{ReLU}(\text{LN}(W_e x))), prediction (\hat{x} = W_p z), error (e = x - \hat{x}).
Error-driven update ( UpdateRule ): (h \leftarrow h + \alpha \cdot \text{MLP}(e)). Key design: MLP uses bias=False . This guarantees (e = 0 \Rightarrow \text{MLP}(0) = 0 \Rightarrow h_{\text{new}} = h) — i.e., zero error yields a strict fixed point; convergence equals prediction self-consistency. (\alpha) is constrained positive via Softplus.
Iterative inference ( InferenceLoop ): state iterates per the above rule; stops early when error norm falls below threshold; BN statistics are frozen during the inference loop.
Neuroscience foundation (Ungerleider & Mishkin, 1982): the ventral pathway (“What”) processes object identity and semantics; the dorsal pathway (“Where/How”) processes spatial location and motor control.
Engineering implementation (v_dual) :
Role separation : the ventral stream receives semantic input with frozen weights; the dorsal stream receives spatial input (residuals + 2D coordinates) and never receives semantic symbol IDs (enforced by runtime assertion).
Direction-aware bias ( RelativeDirectionalBias2D ): FH region tokens are irregularly shaped with no fixed grid, making traditional grid biases unusable. Solution: pairwise token distance (logarithmic bucketing) + direction (angular bucketing) jointly quantized → lookup in a learnable bias table, preserving directional information.
Cross-stream fusion ( DualPathwayFusion ): bidirectional cross-attention + learnable gating (fused = \sigma(\alpha) v + (1 - \sigma(\alpha)) d), gate initialized at 0.5 (equal weight). none mode physically removes the fusion module for ablation.
Micro-parameterized MuPC : in deep predictive networks, gradient variance diverges exponentially with depth. Layer-(l) weight scaling (1/\sqrt{l+1}) (suppresses shallow gradient explosion), error scaling (1/\sqrt{L-l+1}) (alleviates deep gradient vanishing); per-layer gradient variance converges within 10x of each other.
BN freezing : d

[truncated]
