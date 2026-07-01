---
source: "https://github.com/alavicka/pollux"
hn_url: "https://news.ycombinator.com/item?id=48746067"
title: "Pollux – a natively vector quantized LLM with 0.76 bits per parameter"
article_title: "GitHub - alavicka/pollux: Sub-1-bit transformer architecture via native H24 Leech-lattice quantization, achieving 0.76 bits/parameter for SRAM-resident Edge AI and RAG. · GitHub"
author: "pollux_llm"
captured_at: "2026-07-01T13:48:03Z"
capture_tool: "hn-digest"
hn_id: 48746067
score: 1
comments: 0
posted_at: "2026-07-01T13:03:40Z"
tags:
  - hacker-news
  - translated
---

# Pollux – a natively vector quantized LLM with 0.76 bits per parameter

- HN: [48746067](https://news.ycombinator.com/item?id=48746067)
- Source: [github.com](https://github.com/alavicka/pollux)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T13:03:40Z

## Translation

タイトル: Pollux – パラメータあたり 0.76 ビットのネイティブにベクトル量子化された LLM
記事タイトル: GitHub - alavicka/pollux: ネイティブ H24 Leech-lattice 量子化によるサブ 1 ビット トランスフォーマー アーキテクチャ。SRAM 常駐の Edge AI および RAG で 0.76 ビット/パラメータを達成。 · GitHub
説明: ネイティブ H24 Leech-lattice 量子化を介したサブ 1 ビットのトランスフォーマー アーキテクチャ。SRAM 常駐の Edge AI および RAG で 0.76 ビット/パラメータを実現します。 - アラヴィッカ/ポルックス

記事本文:
GitHub - alavicka/pollux: ネイティブ H24 Leech-lattice 量子化によるサブ 1 ビット トランスフォーマー アーキテクチャ。SRAM 常駐の Edge AI および RAG でパラメータあたり 0.76 ビットを達成。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アラヴィッカ
/
ポルックス
公共
通知
サインインする必要があります

o 通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット チェックポイント チェックポイント データ data .gitignore .gitignore README.md README.md Castor.py Castor.py Evaluate.py Evaluate.py Generate.py Generate.py Pack.py Pack.py Polox.py Poloux.py prepare_fineweb.py prepare_fineweb.py train.py train.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Pollux — 0.76 ビット ネイティブ $H_{24}$ Leech-Lattice Transformers
論文: 必要なのは 0.76 ビットだけです: LLM のネイティブ H24 リーチ格子量子化によるベクトル三値論理
Alexander Lavicka · lavicka@cantab.net · プレプリント 2026
WIPO 特許出願番号 PCT/AT2026/060108 およびオーストリア特許出願番号 A65086/2026
Pollux は、フォン ノイマン メモリの壁を克服するためにトランスバックボーンの連続浮動小数点重みを放棄する、デコーダ専用 LLM の根本的に新しいクラスです。
パラメータあたり 0.76 ビット: ニューラル パラメータ多様体を $H_{24}$ Leech 格子 (24D で最も密度の高い球体パッキング) にネイティブにマッピングすることにより、バックボーンは極端なサブ 1 ビット レベルまで圧縮されます。
ゼロ連続重みバックボーン: オブザーバブル レイヤーは連続構造重みを持たず、離散 18 ビット コードブック インデックスと行ごとに 1 つのグローバル FP16 スケールのみを持ちます。
SRAM 常駐エッジ AI: 1B クラスのトランスフォーマー バックボーン (Pollux-1920) がわずか 76 MB の SRAM に圧縮され、推論をメモリ帯域幅制限から計算制限操作に変換します。
RAG の「ステートレス CPU」: Pollux は、流動的なインテリジェンス (構文) を結晶化されたインテリジェンス (事実のトリビア) から物理的に切り離します。幾何学的なボロノイ フィルターを通じて、高エントロピーの事実ノイズを機械的に拒否し、幻覚を引き起こすパラメトリックな知識の矛盾を排除します。
パラメータ

r-Free 熱力学トレーニング: レート スケジュール、ウォームアップ、重量減衰を学習せずにトレーニング。ネットワークは内生熱力学とランダウアー消去によって最適化されます。唯一の環境入力は H_floor です。これは、カルノー理論における周囲温度に相当する、経験的に測定されたコーパス ノイズ フロアです。すべてのアーキテクチャ定数は 2 つの公理から導出されます。ハイパーパラメータ検索は必要ありません。
経験的パリティ: トレーニング データの 1% 未満、アクティブな SRAM フットプリントの半分未満で、Pollux は連続ベースライン (Pythia 160M ～ 410M) による厳密な流体構文パリティ (BLiMP) を達成します。
なぜトレーニング後の圧縮ではなく、ネイティブ ベクトル量子化を使用するのでしょうか?
標準量子化 (INT8、GPTQ、1.58 ビット) は、事前トレーニングされた FP16 モデルを事後的に近似する (容量を破壊する) か、組み合わせ状態空間を無駄にする 1 次元のスカラー丸めに依存します。 Pollux は、ステップ 0 から最終的な量子化解像度でトレーニングされます。近似するための連続したベースラインはありません。
スカラー 3 値量子化 (24 次元の $3^{24} \約 282 \times 10^9$ 状態) の代わりに、Pollux は Leech 格子 $\Lambda_{24}$ の 196,560 個の数学的に最適なキス ポイントを使用します。これにより、ビットあたりの可能な限り最高の構造解像度が提供されます。原点 (ゼロベクトル) はコードブックの先頭に付加され、ベクトル三値論理のヌル アトラクターとして機能します。
$$
\underbrace{18\ \text{bits}}_{\text{LUT インデックス}} /\ \underbrace{24\ \text{params}}_{\text{atom dim}}
+
\underbrace{16\ \text{bits}}_{\text{FP16 スケール}} /\ \underbrace{1152\ \text{params}}_{d\text{-dim row}}
=\ 0.750 + 0.0138 \およそ \mathbf{0.76\text{ビット/パラメータ}}
$$
流動的な知性と結晶化した知性
格子の $C=\sqrt{2}$ ボロノイ深穴バリアは、勾配に対する物理ハイパス フィルターとして機能します。
F

luid Intelligence (構造構文): 一貫した繰り返しの構文規則は、指向性のある運動量を蓄積し、ボロノイ障壁を越え、$H_{24}$ キス ポイントに永続的に結晶化します。
結晶化された知能 (事実のトリビア): 支離滅裂で高エントロピーの事実ノイズは、しきい値を越えることができず、ゼロポテンシャルのヌルアトラクターにルーティングされます。したがって、Pollux は、数十億のトークンにわたって一貫した勾配を生成するユビキタスな事実の高周波漏洩によって制限された、事実のベンチマークでランダムな偶然に近い、またはわずかに上回るスコアを記録します。
これにより、Pollux は検索拡張生成 (Macro-RAG) の究極のエンジンになります。これは、出力を内部パラメトリック バイアスや幻覚で汚染することなく、外部ベクトル データベースに盲目的に従う純粋なステートレス推論エンジンとして機能します。
Pollux モデルは、厳密な Iso-Memory パラダイムに基づいて評価されます。生のパラメータ数ではなく、アクティブなバックボーン SRAM フットプリントが、自己回帰生成中の実行に不可欠なメトリクスです。 (注: Iso-Memory 基準は、ターゲットのネイティブ LUT ランタイムの下でメモリ帯域幅のフットプリントを分離します。現在の FP16 リファレンスの具体化では、トークンあたりの FLOP はバックボーン パラメーターの数に応じてスケールされ、Pollux ベースラインと Pythia ベースラインの間では一致しません。)
アイソメモリの評価とトリレンマの打破
連続モデルはパレートのトリレンマに陥っています。つまり、SRAM の最小化、流体構文の最大化 (BLiMP)、および事実の汚染の抑制 (SciQ/HellaSwag) を同時に行うことはできません。 Pollux-1920 は、事実の記憶を機械的に抵抗しながら、76 MB のエンベロープ内で Pythia-410M の推論能力に匹敵し、この領域を突破します。
(ランダムチャンスベースライン: BLiMP (2-way) = 50.0%、HellaSwag / SciQ (4-way) ≈ 25%、PIQA (2-way) ≈ 50%。すべての Pollux スコアはパックされた状態で測定

.plx デプロイメントアーティファクト。)
熱力学容量曲線と「ディープフリーズ」
10,000 ステップ (約 26 億トークン) で、ネットワークは熱力学的結晶化ピークに達します。この点を超えると、$H_{24}$ トポロジカル結合エネルギーによって構文が所定の位置にロックされ、熱力学的静止状態 (「ディープ フリーズ」) の段階に入ります。BLiMP は 0.5% 以下でシフトし、すべての実際のベンチマークは 1.3B 以上の追加トークンにわたって 1.0% 以下でシフトします。容量のチャーンは止まります。モデルは、新しい事実の関連性を獲得することも、確立された構文構造を失うこともありません。
トポロジカルロバスト性 (ロスレスシリアル化)
.plx シリアル化は、ネットワークを数学的に 0.76 ビット/パラメータに圧縮します。生の連続 .pt チェックポイントと比較すると、すべての BLiMP タスクにわたる最大偏差は 0.2 pp (および合計平均差は 0.01%) であり、グローバルな行スケールの量子化が実質的に無損失であることがエンジニアリングによって確認されました。
ハードウェアと推論の制限
Pollux は、研究用の機能参照実装です。次の制約は、コードベースをデプロイまたは拡張するすべてのユーザーに適用されます。
パックされたストレージと PyTorch ランタイム: パックされた .plx 表現は完全にオンチップ メモリ (Pollux-1152 の場合は約 27 MB バックボーン) に収まりますが、現在の参照 PyTorch パスは、標準の cuBLAS 互換性のために推論時に高密度の FP16 重み行列を具体化します ( PackedH24Linear.materialize() )。これにより、結晶化とゼロショット ベンチマークが検証されますが、ネイティブ SRAM に依存したレイテンシは提供されません。ネイティブのマトリックスフリー LUT 収集累積カーネル (インデックスの読み取り → コードブック ベクトルのフェッチ → 累積 $\sigma_{\mathrm{rms}} \cdot c$ ) は、真の計算重視の高速化に必要です。
エッジ CPU の実行可能性と RAM ボトルネック: 標準 GPU は、サブバイトの組み合わせアドレス指定に大きなペナルティを与えます。ただし、最新の CPU は大きな L を備えています。

9 MB $H_{24}$ コードブック全体を保持できる 3 つのキャッシュ (8 ～ 32 MB) により、インデックス ルーティング パイプラインを極めて効率的に実行できます。 1B クラスのモデルを厳密な 265 MB のディスク上のフットプリントに圧縮することで、Pollux は、連続モデルが瞬時にメモリ不足 (OOM) 障害を引き起こす IoT/エッジ ハードウェアの推論を解き放ちます。
アーキテクチャの厳密性: カスタム構成は n_embd % 24 == 0 を満たす必要があります。 Leech 格子原子タイリングを適切に行うには、すべての量子化された線形 in_features が 24 できれいに割り切れる必要があります。
出版/
│
§── Castor.py # 公理層: Leech 格子コードブック、定数、
│ # 最近傍量子化器、ビットパッキング。
│ # リーフ ノード — このプロジェクトからは何もインポートしません。
│
§──pollux.py # ゼロ連続重みアーキテクチャ +
│ # パラメータフリーの熱力学推定器
│ # (pollux_step)。キャスターのみに依存します。
│ # トレーニング (PolluxH24Linear) とトレーニング (PolluxH24Linear) の両方が含まれています
│ # 推論 (PackedH24Linear) 層クラス。
│
§── train.py # トレーニングのエントリ ポイント。 FineWeb-Edu memmap を読み取り、
│ # polux_step を呼び出し、.pt チェックポイントを書き込みます。
│ # LR スケジュールなし、重量減衰なし、ウォームアップなし。
│
§── prepare_fineweb.py # FineWeb-Edu 10B をダウンロード、GPT-2 でトークン化、
│ # uint16 memmap を data/fineweb_10b.bin に書き込みます。
│
§── Pack.py # チェックポイント → .plx コンバータ。
│ # H24 レイヤーを 18 ビット LUT インデックスに量子化 +
│ # FP16 行ごとの σ_rms、INT8 量子化埋め込み。
│ # 10k 結晶化ピークのチェックポイントでパックします。
│
§──generate.py # .plx または .pt ファイルからテキストを生成します。
│ # .plx:index_select 実体化 + F.linear;
│ # ネイティブ LUT カーネル (将来) は密集を排除します
│ # FP アクティベーションではなく、重みマトリックス トラフィック。
│
§──evaluate.py # lm-eval-harness ラッパー。層別プリント
│ # 構造 (4 BLiMP) 対事実

デュアル (4 MCQ) テーブル。
│ # .plx と .pt の両方の入力を受け入れます。
│
§── data/ # ローカル トレーニング コーパス (gitignored; によって作成されました)
│ ━──fineweb_10b.bin # prepare_fineweb.py)
│
└── チェックポイント/ # トレーニング チェックポイント (gitignored; によって作成されました)
lux_step_*.pt # train.py 2.5k オプティマイザー ステップごと)
推論パイプライン
train.py ──(pollux_10k.pt)──►pack.py ──(model.plx)──►generate.py
──►evaluate.py
事前トレーニングされたモデルと重み
完全に結晶化された 0.76 ビット .plx デプロイメント アーティファクトは、Hugging Face でホストされています。これらのコンテナーは完全にパックされており、グローバルな行単位の RMS スケールとともに不変の H24 コードブック インデックスが含まれています。
Pollux-1152 : 287M バックボーン パラメータが 27 MB SRAM に圧縮されました (INT8 埋め込みを含むディスク上の合計 142 MB)。
Pollux-1920 : 796M のバックボーン パラメータが 76 MB SRAM に圧縮されました (INT8 埋め込みを含むディスク上の合計 265 MB)。
ファイル サイズに関する注意: ここおよび論文で報告されているディスク上のフットプリント (142 MB / 265 MB) は、標準の OS 環境で計算されたバイナリ メガバイト (MiB) を指します。 Hugging Face ファイル エクスプローラーでは、これらの同一のファイルが 10 進数の SI 単位 (149 MB / 278 MB) を使用して表示されます。
ネイティブ推論に関するテクニカルノート:
現在のリファレンス PyTorch ランタイムは、標準 F.linear / cuBLAS を介して実行される、index_select を介して FP16 重みタイルへの 18 ビット インデックスを具体化します。これは明示的に検証します

[切り捨てられた]

## Original Extract

Sub-1-bit transformer architecture via native H24 Leech-lattice quantization, achieving 0.76 bits/parameter for SRAM-resident Edge AI and RAG. - alavicka/pollux

GitHub - alavicka/pollux: Sub-1-bit transformer architecture via native H24 Leech-lattice quantization, achieving 0.76 bits/parameter for SRAM-resident Edge AI and RAG. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
alavicka
/
pollux
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits checkpoints checkpoints data data .gitignore .gitignore README.md README.md castor.py castor.py evaluate.py evaluate.py generate.py generate.py pack.py pack.py pollux.py pollux.py prepare_fineweb.py prepare_fineweb.py train.py train.py View all files Repository files navigation
Pollux — 0.76-Bit Native $H_{24}$ Leech-Lattice Transformers
Paper: 0.76 Bits Is All You Need: Vector Ternary Logic via Native H24 Leech-Lattice Quantization in LLMs
Alexander Lavicka · lavicka@cantab.net · Preprint 2026
WIPO Patent Application No. PCT/AT2026/060108 and Austrian Patent Application No. A65086/2026
Pollux is a fundamentally new class of decoder-only LLMs that abandons continuous floating-point weights in the transformer backbone to overcome the von Neumann memory wall.
0.76 Bits per Parameter: By mapping the neural parameter manifold natively onto the $H_{24}$ Leech lattice (the densest sphere packing in 24D), the backbone is compressed to extreme sub-1-bit levels.
Zero-Continuous-Weight Backbone: Observable layers carry no continuous structural weights—only discrete 18-bit codebook indices and a single global FP16 scale per row.
SRAM-Resident Edge AI: A 1B-class transformer backbone (Pollux-1920) is compressed into just 76 MB of SRAM , converting inference from a memory-bandwidth-bound to a compute-bound operation.
The "Stateless CPU" for RAG: Pollux physically decouples fluid intelligence (syntax) from crystallised intelligence (factual trivia). Through a geometric Voronoi filter, it mechanically rejects high-entropy factual noise, eliminating the parametric knowledge conflicts that trigger hallucinations.
Parameter-Free Thermodynamic Training: Trained without learning rate schedules, warmup, or weight decay. The network is optimized via endogenous thermodynamic kinetics and Landauer erasure. The only environmental input is H_floor — the empirically measured corpus noise floor, analogous to ambient temperature in Carnot theory. All architectural constants are derived from two axioms; no hyperparameter search is required.
Empirical Parity: At less than 1% of the training data and less than half the active SRAM footprint, Pollux achieves strict fluid-syntax parity (BLiMP) with continuous baselines (Pythia 160M–410M).
Why native vector quantization, not post-training compression?
Standard quantization (INT8, GPTQ, 1.58-bit) either approximates a pre-trained FP16 model after the fact (destroying capacity) or relies on one-dimensional scalar rounding that wastes combinatorial state-space. Pollux is trained at its final quantization resolution from step zero . There is no continuous baseline to approximate.
Instead of scalar ternary quantization ( $3^{24} \approx 282 \times 10^9$ states for 24 dimensions), Pollux uses the 196,560 mathematically optimal kissing points of the Leech lattice $\Lambda_{24}$ . This provides the highest possible structural resolution per bit. The origin (zero-vector) is prepended to the codebook, acting as a null attractor for vector ternary logic.
$$
\underbrace{18\ \text{bits}}_{\text{LUT index}} /\ \underbrace{24\ \text{params}}_{\text{atom dim}}
+
\underbrace{16\ \text{bits}}_{\text{FP16 scale}} /\ \underbrace{1152\ \text{params}}_{d\text{-dim row}}
=\ 0.750 + 0.0138 \approx \mathbf{0.76\text{bits/param}}
$$
Fluid vs. Crystallised Intelligence
The $C=\sqrt{2}$ Voronoi deep-hole barrier of the lattice acts as a physical high-pass filter on gradients:
Fluid intelligence (structural syntax): Coherent, recurring syntactic rules accumulate directed momentum, cross the Voronoi barrier, and permanently crystallise into $H_{24}$ kissing points.
Crystallised intelligence (factual trivia): Incoherent, high-entropy factual noise fails to cross the threshold and is routed into the zero-potential null attractor. Pollux therefore scores near or modestly above random chance on factual benchmarks — bounded by high-frequency leakage for ubiquitous facts that generate coherent gradients over billions of tokens.
This makes Pollux the ultimate engine for Retrieval-Augmented Generation (Macro-RAG) : It acts as a pure, stateless reasoning engine that blindly obeys external vector databases without contaminating the output with internal parametric bias or hallucinations.
Pollux models are evaluated under a strict Iso-Memory paradigm : the active backbone SRAM footprint—not raw parameter count—is the execution-critical metric during autoregressive generation. (Note: the Iso-Memory criterion isolates memory-bandwidth footprint under the targeted native LUT runtime. Under the current FP16 reference materialisation, FLOPs per token scale with backbone parameter count and are not matched between Pollux and Pythia baselines.)
Iso-Memory Evaluation & Breaking the Trilemma
Continuous models are trapped in a Pareto trilemma: they cannot simultaneously minimize SRAM, maximize fluid syntax (BLiMP), and suppress factual contamination (SciQ/HellaSwag). Pollux-1920 breaks this frontier, matching the reasoning capacity of Pythia-410M inside a 76 MB envelope while mechanically resisting factual memorisation.
(Random-chance baselines: BLiMP (2-way) = 50.0%; HellaSwag / SciQ (4-way) ≈ 25%; PIQA (2-way) ≈ 50%. All Pollux scores measured on packed .plx deployment artifacts.)
Thermodynamic Capacity Curve & The "Deep Freeze"
At 10k steps (~2.6B tokens), the network reaches its thermodynamic crystallisation peak. Beyond this point, the $H_{24}$ topological binding energy locks syntax in place, entering a phase of thermodynamic stasis (the "Deep Freeze"): BLiMP shifts by ≤ 0.5% and all factual benchmarks shift by ≤ 1.0% over ≥ 1.3B additional tokens. Capacity churn ceases; the model neither gains new factual associations nor loses established syntactic structure.
Topological Robustness (Lossless Serialization)
The .plx serialization mathematically compresses the network to 0.76 bits/param. Compared to the raw continuous .pt checkpoint, the maximum deviation across all BLiMP tasks is 0.2 pp (and 0.01% aggregate mean difference)—engineering confirmation that the global row-scale quantization is practically lossless.
Hardware & Inference Limitations
Pollux is a functional reference implementation for research . The following constraints apply to anyone deploying or extending the codebase:
Packed storage vs. PyTorch runtime: While the packed .plx representation fits entirely in on-chip memory (~27 MB backbone for Pollux-1152), the current reference PyTorch path materialises dense FP16 weight matrices at inference time ( PackedH24Linear.materialize() ) for standard cuBLAS compatibility. This validates crystallisation and zero-shot benchmarks but does not deliver native SRAM-bound latency. Native matrix-free LUT gather–accumulate kernels (read index → fetch codebook vector → accumulate $\sigma_{\mathrm{rms}} \cdot c$ ) are required for the true compute-bound speedup.
Edge CPU Viability & The RAM Bottleneck: Standard GPUs severely penalise sub-byte combinatorial addressing. However, modern CPUs feature large L3 caches (8–32 MB) capable of holding the entire 9 MB $H_{24}$ codebook, executing the index-routing pipeline with extreme efficiency. By compressing a 1B-class model to a strict 265 MB on-disk footprint , Pollux unlocks reasoning for IoT/Edge hardware where continuous models instantly trigger Out-Of-Memory (OOM) failures.
Architectural Strictness: Custom configurations must satisfy n_embd % 24 == 0 . Every quantized linear in_features must be cleanly divisible by 24 for proper Leech lattice atom tiling.
publish/
│
├── castor.py # Axiom layer: Leech lattice codebook, constants,
│ # nearest-neighbour quantizer, bit-packing.
│ # Leaf node — imports nothing from this project.
│
├── pollux.py # Zero-continuous-weight architecture +
│ # parameter-free thermodynamic estimator
│ # (pollux_step). Depends only on castor.
│ # Contains both training (PolluxH24Linear) and
│ # inference (PackedH24Linear) layer classes.
│
├── train.py # Training entry point. Reads FineWeb-Edu memmap,
│ # calls pollux_step, writes .pt checkpoints.
│ # No LR schedule, no weight decay, no warmup.
│
├── prepare_fineweb.py # Downloads FineWeb-Edu 10B, tokenizes with GPT-2,
│ # writes uint16 memmap to data/fineweb_10b.bin.
│
├── pack.py # Checkpoint → .plx converter.
│ # Quantizes H24 layers to 18-bit LUT indices +
│ # FP16 σ_rms per row, INT8-quantized embeddings.
│ # Pack at the 10k crystallisation peak checkpoint.
│
├── generate.py # Text generation from .plx or .pt files.
│ # .plx: index_select materialisation + F.linear;
│ # native LUT kernels (future) eliminate dense
│ # weight-matrix traffic, not FP activations.
│
├── evaluate.py # lm-eval-harness wrapper. Prints stratified
│ # Structural (4 BLiMP) vs Factual (4 MCQ) table.
│ # Accepts both .plx and .pt inputs.
│
├── data/ # Local training corpus (gitignored; created by
│ └── fineweb_10b.bin # prepare_fineweb.py)
│
└── checkpoints/ # Training checkpoints (gitignored; written by
└── pollux_step_*.pt # train.py every 2.5k optimizer steps)
Inference pipeline
train.py ──(pollux_10k.pt)──► pack.py ──(model.plx)──► generate.py
──► evaluate.py
Pre-trained Models & Weights
The fully crystallized, 0.76-bit .plx deployment artifacts are hosted on Hugging Face. These containers are fully packed and include the immutable H24 codebook indices alongside the global row-wise RMS scales.
Pollux-1152 : 287M backbone parameters compressed into 27 MB SRAM (142 MB total on-disk including INT8 embeddings).
Pollux-1920 : 796M backbone parameters compressed into 76 MB SRAM (265 MB total on-disk including INT8 embeddings).
Note on File Sizes: The on-disk footprints (142 MB / 265 MB) reported here and in the paper refer to binary Megabytes (MiB) as calculated by standard OS environments. The Hugging Face file explorer displays these identical files using decimal SI units (149 MB / 278 MB).
Technical Note on Native Inference:
The current reference PyTorch runtime materialises 18-bit indices to FP16 weight tiles via index_select , executing via standard F.linear / cuBLAS . This explicitly validates th

[truncated]
