---
source: "https://github.com/vishal-dehurdle/hypersae"
hn_url: "https://news.ycombinator.com/item?id=49262340"
title: "Show HN: HyperSAE – Hyperbolic Sparse Autoencoders for LLM Interpretability"
article_title: "GitHub - vishal-dehurdle/hypersae: High-Performance Hyperbolic Sparse Autoencoders for Mechanistic Interpretability · GitHub"
author: "visha1v"
captured_at: "2026-08-11T18:47:24Z"
capture_tool: "hn-digest"
hn_id: 49262340
score: 2
comments: 0
posted_at: "2026-08-11T18:19:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: HyperSAE – Hyperbolic Sparse Autoencoders for LLM Interpretability

- HN: [49262340](https://news.ycombinator.com/item?id=49262340)
- Source: [github.com](https://github.com/vishal-dehurdle/hypersae)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T18:19:02Z

## Translation

タイトル: 表示 HN: HyperSAE – LLM 解釈性のための双曲線スパース オートエンコーダ
記事のタイトル: GitHub - vishal-dehurdle/hypersae: 機械的解釈のための高性能双曲線スパース オートエンコーダ · GitHub
説明: 機械的な解釈を可能にする高性能双曲線スパース オートエンコーダ - vishal-dehurdle/hypersae

記事本文:
GitHub - vishal-dehurdle/hypersae: 機械的解釈のための高性能双曲線スパース オートエンコーダ · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ヴィシャル・デハードル
/
ハイパーサエ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .github .github src/ hypersae src/ hypersae テスト テスト .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス

SE ライセンス README.md README.md SECURITY.md SECURITY.md cloud_run.py cloud_run.py eval_gpqa.py eval_gpqa.py pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json run_benchmarks.py run_benchmarks.py setup_vm.sh setup_vm.sh steer_llm.py steer_llm.py train_on_llm.py train_on_llm.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
HyperSAE: 高性能双曲線スパース オートエンコーダー
hypersae は、大規模言語モデル (LLM) から階層概念オントロジーを抽出するように設計された高性能の機構的解釈可能エンジンです。双曲幾何学をフォワードパスから切り離すことにより、リーマン負の曲率のセマンティックマッピング能力とともに、標準ユークリッドスパースオートエンコーダのゼロレイテンシー実行を提供します。
pip インストールハイパーサエ
または、ソースからローカルにインストールします。
git クローン https://github.com/vishal-dehurdle/hypersae.git
CDハイパーサエ
pip install -e 。
1. コア アーキテクチャ: 分離された重み空間正則化
高い GPU スループットとモデルの互換性を維持するために、hypersae は実行を 2 つの計算速度に分割します。
高速パス (ユークリッド フォワード パス): 大量のアクティブなトークン データは完全にフラットな高速ユークリッド空間 ( $\mathbb{R}^d$ ) に残ります。これにより、リーマン多様体の待ち時間が回避され、基本モデルの正規化 ( RMSNorm など) が尊重され、直接的な因果関係のあるステアリングの互換性が維持されます。
スローパス (双曲線重み正則化): 概念の構造的階層関係は、ポアンカレ球投影 $(\mathcal{B}^d, g_{\mathbf{x}})$ による最適化中に辞書パラメータ空間でのみ適用されます。
グラフTD
サブグラフ Fast_Path["高速パス: ユークリッド前方パス (bfloat16)"]
X["正規化されたトークンのアクティブ化 x"] --> ENC["ユークリッド エンコーダー"]
ENC --> F["スパースアクティベーション

nsf"]
F --> DEC["ユークリッド デコーダ (W_dec)"]
12 月 --> X_HAT["再構築されたアクティベーション x̂"]
終わり
subgraph Slow_Path["スローパス: 双曲線重みの最適化 (float32 へのアップキャスト)"]
W_dec["デコーダーの重み (W_dec)"] & R_ Depth["深度スカラー (r_i)"] --> MAP["ポアンカレ多様体投影"]
MAP --> H_coords["双曲線座標 (h_i)"]
H_coords --> MOCO["CoActivation Queue"]
MOCO --> LOSS[「非対称ポアンカレ含意損失」]
終わり
損失 -.->|"デュアル オプティマイザー アップデート (AdamW / RiemannianAdam)"| W_dec
読み込み中
2. 経験的ベンチマーク結果 (Gemma-2-2B レイヤ 13)
NVIDIA L4 GPU クラスター上の FineWeb-Edu の 2,000 万トークンを超えるストリーミング Google Gemma-2-2B レイヤー 13 残差ストリーム アクティベーション ( $d=2304$ 、辞書サイズ $M=16384$ ) で大規模に評価しました。
ダウンストリーム推論の保持 (単一トークン置換)
ベンチマーク
Gemma-2-2B ベースライン
FlatSAE (ベースライン)
HyperSAE (当社)
相対保持容量
GPQA ダイヤモンド
66.67%
100.00%
100.00%
100% の精度を維持
MMLU-Pro (12,032 の質問)
17.69%
16.11%
16.26%
HyperSAE は優れた精度を維持 (+0.15%)
パレート再構築とスパーシティパフォーマンス
モデルのアーキテクチャ
$L_1$ ペナルティ
アクティブな機能/トークン ( $L_0$ )
再構築 MSE ( $\downarrow$ )
CE 損失回復 % ( $\uparrow$ )
フック付きCEロス
HyperSAE (当社)
0.005
54.2
4.1232
78.9%
6.1164
FlatSAE (ベースライン)
0.005
52.4
4.5724
75.5%
6.3861
HyperSAE (当社)
0.001
988.8
1.3965
97.7%
4.6036
FlatSAE (ベースライン)
0.001
744.5
1.7364
97.2%
4.6499
HyperSAE (当社)
0.0005
2285.4
0.7666
98.1%
4.5721
FlatSAE (ベースライン)
0.0005
1511.8
1.0112
97.0%
4.6608
重要なポイント: HyperSAE は、一致するスパース性 ( $L_0 \約 53$ ) でフラット SAE と比較して、再構成 MSE の 9.8% 削減とクロスエントロピー損失回復の +3.4% 向上を達成します。
輸入トーチ
ハイパーサエからインポート HyperSAE 、 CoActivationQueue 、 TriPartit

eLoss、HyperSAETrainer
device = "cuda" if torch 。クダ。 is_available () else "CPU"
#1. HyperSAE モデル、CoActivationQueue、および TriPartiteLoss をインスタンス化する
sae = HyperSAE ( d_model = 2304 、dict_size = 16384 )。へ (デバイス)
キュー = CoActivationQueue ( dict_size = 16384 )。へ (デバイス)
loss_fn = TriPartiteLoss ( l1_coeff = 0.005 、 entail_coeff = 0.01 )
# 2. HyperSAETrainer をインスタンス化する
トレーナー = HyperSAETrainer (モデル = sae 、キュー = queue 、loss_fn = loss_fn 、lr = 1e-3)
# 3. 残留ストリームのアクティブ化バッチのトレーニング ステップ
x = トーチ。 randn ( 64 , 2304 , device = デバイス )
メトリクス = トレーナー。トレイン_ステップ ( x )
print ( f"総損失: { metrics [ 'loss_total' ]:.4f } " )
print ( f"再構築 MSE: { metrics [ 'loss_recon' ]:.4f } " )
print ( f"ポアンカレ含意ペナルティ: { metrics [ 'loss_entail' ]:.4f } " )
4. ソフトウェアアーキテクチャ
hypersae.HyperSAE : 線形前方パスと学習可能な半径方向の深さ $r_i \in [0, 1)$ を実装するコア モデル モジュール。
hypersae.FlatSAE : ベンチマーク比較用の標準ユークリッド ベースライン。
hypersae.TriPartiteLoss : MSE、$L_1$ スパース性、およびポアンカレ錐含意ペナルティを組み合わせた損失オーケストレーター。
hypersae.CoActivationQueue : $\mathcal{O}(M^2)$ メモリ増加を伴わない非同期 GPU メモリ キュー追跡機能。
hypersae.hooks : ステアリングと介入のための PyTorch および TransformerLens の前方フック ユーティリティ。
5. 研究論文と出版物
理論論文: 平地からの脱出: 力学的解釈における重量空間正則化と双曲幾何学
実証論文: 双曲線疎オートエンコーダ: LLM 活性化におけるポアンカレ多様体幾何の実証的検証
このプロジェクトは MIT ライセンスに基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
M 用の高性能双曲線スパース オートエンコーダ

機械的な解釈可能性
vishalvermalabs.com/papers/empirical-validation-hypersae-poincare-geometry/ トピック
Readme MIT ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

High-Performance Hyperbolic Sparse Autoencoders for Mechanistic Interpretability - vishal-dehurdle/hypersae

GitHub - vishal-dehurdle/hypersae: High-Performance Hyperbolic Sparse Autoencoders for Mechanistic Interpretability · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
vishal-dehurdle
/
hypersae
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .github .github src/ hypersae src/ hypersae tests tests .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md cloud_run.py cloud_run.py eval_gpqa.py eval_gpqa.py pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json run_benchmarks.py run_benchmarks.py setup_vm.sh setup_vm.sh steer_llm.py steer_llm.py train_on_llm.py train_on_llm.py View all files Repository files navigation
HyperSAE: High-Performance Hyperbolic Sparse Autoencoders
hypersae is a high-performance mechanistic interpretability engine designed to extract hierarchical concept ontologies from Large Language Models (LLMs). By decoupling hyperbolic geometry from the forward pass, it provides the zero-latency execution of standard Euclidean Sparse Autoencoders alongside the semantic mapping power of Riemannian negative curvature.
pip install hypersae
Or install locally from source:
git clone https://github.com/vishal-dehurdle/hypersae.git
cd hypersae
pip install -e .
1. Core Architecture: Decoupled Weight-Space Regularization
To preserve high GPU throughput and model compatibility, hypersae separates execution into two computational speeds:
The Fast-Path (Euclidean Forward Pass): The massive volume of active token data remains entirely in flat, high-speed Euclidean space ( $\mathbb{R}^d$ ). This avoids the latency of Riemannian manifolds, respects base model normalizations (e.g., RMSNorm ), and maintains direct causal steering compatibility.
The Slow-Path (Hyperbolic Weight Regularization): The structural, hierarchical relationships of concepts are enforced exclusively in the dictionary parameter space during optimization via Poincaré ball projections $(\mathcal{B}^d, g_{\mathbf{x}})$ .
graph TD
subgraph Fast_Path["Fast-Path: Euclidean Forward Pass (bfloat16)"]
X["Normalized Token Activations x"] --> ENC["Euclidean Encoder"]
ENC --> F["Sparse Activations f"]
F --> DEC["Euclidean Decoder (W_dec)"]
DEC --> X_HAT["Reconstructed Activations x̂"]
end
subgraph Slow_Path["Slow-Path: Hyperbolic Weight Optimization (Upcast to float32)"]
W_dec["Decoder Weights (W_dec)"] & R_depth["Depth Scalars (r_i)"] --> MAP["Poincaré Manifold Projection"]
MAP --> H_coords["Hyperbolic Coordinates (h_i)"]
H_coords --> MOCO["CoActivation Queue"]
MOCO --> LOSS["Asymmetric Poincaré Entailment Loss"]
end
LOSS -.->|"Dual-Optimizer Update (AdamW / RiemannianAdam)"| W_dec
Loading
2. Empirical Benchmark Results (Gemma-2-2B Layer 13)
Evaluated at scale on Google Gemma-2-2B Layer 13 residual stream activations ( $d=2304$ , dict size $M=16384$ ) streaming over 20M tokens of FineWeb-Edu on an NVIDIA L4 GPU cluster :
Downstream Reasoning Retention (Single-Token Substitution)
Benchmark
Gemma-2-2B Baseline
FlatSAE (Baseline)
HyperSAE (Ours)
Relative Retained Capacity
GPQA Diamond
66.67%
100.00%
100.00%
100% Accuracy Preserved
MMLU-Pro (12,032 Questions)
17.69%
16.11%
16.26%
HyperSAE Retains Superior Accuracy (+0.15%)
Pareto Reconstruction & Sparsity Performance
Model Architecture
$L_1$ Penalty
Active Features / Token ( $L_0$ )
Reconstruction MSE ( $\downarrow$ )
CE Loss Recovery % ( $\uparrow$ )
CE Loss with Hook
HyperSAE (Ours)
0.005
54.2
4.1232
78.9%
6.1164
FlatSAE (Baseline)
0.005
52.4
4.5724
75.5%
6.3861
HyperSAE (Ours)
0.001
988.8
1.3965
97.7%
4.6036
FlatSAE (Baseline)
0.001
744.5
1.7364
97.2%
4.6499
HyperSAE (Ours)
0.0005
2285.4
0.7666
98.1%
4.5721
FlatSAE (Baseline)
0.0005
1511.8
1.0112
97.0%
4.6608
Key Takeaway: HyperSAE achieves a 9.8% reduction in reconstruction MSE and a +3.4% boost in Cross-Entropy Loss Recovery over flat SAEs at matching sparsity ( $L_0 \approx 53$ ).
import torch
from hypersae import HyperSAE , CoActivationQueue , TriPartiteLoss , HyperSAETrainer
device = "cuda" if torch . cuda . is_available () else "cpu"
# 1. Instantiate HyperSAE model, CoActivationQueue, and TriPartiteLoss
sae = HyperSAE ( d_model = 2304 , dict_size = 16384 ). to ( device )
queue = CoActivationQueue ( dict_size = 16384 ). to ( device )
loss_fn = TriPartiteLoss ( l1_coeff = 0.005 , entail_coeff = 0.01 )
# 2. Instantiate HyperSAETrainer
trainer = HyperSAETrainer ( model = sae , queue = queue , loss_fn = loss_fn , lr = 1e-3 )
# 3. Train step on residual stream activation batch
x = torch . randn ( 64 , 2304 , device = device )
metrics = trainer . train_step ( x )
print ( f"Total Loss: { metrics [ 'loss_total' ]:.4f } " )
print ( f"Reconstruction MSE: { metrics [ 'loss_recon' ]:.4f } " )
print ( f"Poincaré Entailment Penalty: { metrics [ 'loss_entail' ]:.4f } " )
4. Software Architecture
hypersae.HyperSAE : Core model module implementing linear forward pass and learnable radial depths $r_i \in [0, 1)$ .
hypersae.FlatSAE : Standard Euclidean baseline for benchmark comparison.
hypersae.TriPartiteLoss : Loss orchestrator combining MSE, $L_1$ sparsity, and Poincaré cone entailment penalties.
hypersae.CoActivationQueue : Asynchronous GPU memory queue tracking feature co-occurrences without $\mathcal{O}(M^2)$ memory growth.
hypersae.hooks : PyTorch and TransformerLens forward hook utilities for steering and intervention.
5. Research Papers & Publications
Theoretical Paper: Escaping Flatland: Weight-Space Regularization and Hyperbolic Geometry in Mechanistic Interpretability
Empirical Paper: Hyperbolic Sparse Autoencoders: Empirical Validation of Poincaré Manifold Geometry on LLM Activations
This project is licensed under the MIT License — see the LICENSE file for details.
High-Performance Hyperbolic Sparse Autoencoders for Mechanistic Interpretability
vishalvermalabs.com/papers/empirical-validation-hypersae-poincare-geometry/ Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
