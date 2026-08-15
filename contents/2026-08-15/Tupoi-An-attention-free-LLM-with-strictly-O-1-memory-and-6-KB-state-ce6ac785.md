---
source: "https://github.com/narelabs/TUPOI"
hn_url: "https://news.ycombinator.com/item?id=49312921"
title: "Tupoi: An attention-free LLM with strictly O(1) memory and 6 KB state"
article_title: "GitHub - narelabs/TUPOI: TUPOI: Symplectic Post-Transformer Language Model with O(1) Memory (No KV-Cache). Built by 15 y.o. independent researcher. · GitHub"
author: "IntellegenceIsP"
captured_at: "2026-08-15T19:14:47Z"
capture_tool: "hn-digest"
hn_id: 49312921
score: 1
comments: 0
posted_at: "2026-08-15T18:21:24Z"
tags:
  - hacker-news
  - translated
---

# Tupoi: An attention-free LLM with strictly O(1) memory and 6 KB state

- HN: [49312921](https://news.ycombinator.com/item?id=49312921)
- Source: [github.com](https://github.com/narelabs/TUPOI)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T18:21:24Z

## Translation

タイトル: Tupoi: 厳密に O(1) メモリと 6 KB 状態のアテンションフリー LLM
記事のタイトル: GitHub - narelabs/TUPOI: TUPOI: O(1) メモリ (KV キャッシュなし) を備えたシンプレクティック ポストトランスフォーマー言語モデル。 15歳までに建てられた独立した研究者。 · GitHub
説明: TUPOI: O(1) メモリ (KV キャッシュなし) を備えたシンプレクティック ポストトランスフォーマー言語モデル。 15歳までに建てられた独立した研究者。 - ナレラボ/TUPOI

記事本文:
GitHub - narelabs/TUPOI: TUPOI: O(1) メモリ (KV キャッシュなし) を備えたシンプレクティック ポストトランスフォーマー言語モデル。 15歳までに建てられた独立した研究者。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ナレラボ
/
ツポイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット アセット アセット ベンチマーク ベンチマーク スクリプト スクリプト tupoi tupoi .gitigno

re .gitignore README.md README.md 要件.txt 要件.txt setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
TUPOI: O(1) メモリを備えたシンプレクティックなポストトランスフォーマー言語モデル
密なアテンション行列をシンプレクティック ハミルトニアン積分器 (Velocity-Verlet) に置き換える、準 2 次のアテンションフリー シーケンス モデリング アーキテクチャ。
著者: NARE LABS (15 歳で作成)
このリポジトリには、TUPOI 言語モデルのリファレンス実装、事前トレーニングされた重み、経験的ベンチマーク スイートが含まれています。TUPOI 言語モデルは、アテンション メカニズムと永続的な KV キャッシュを完全に排除して、連続的なハミルトニアン位相空間ダイナミクスを優先するポストトランスフォーマー アーキテクチャです。
この研究の主な目的は、シンプレクティック位相空間 $(q, p)$ を移動する粒子としてのシーケンス コンテキストのモデリングが、 $O(1)$ のメモリ消費とゼロ状態散逸 ( $\det(J) \equiv 1.0$ ) を厳密に保証しながら、標準の Transformer モデルの表現力と一致できるかどうかを調査することです。
TUPOI は、$O(N^2)$ ペアワイズ アテンションを深いシンプレクティック遷移システムに置き換えます。
1. シンプレクティックハミルトニアンコア
隠れ状態は正準座標 $(q, p)$ に分解され、Velocity-Verlet Leapfrog 動的スキームを介して 24 層を通じて統合されます。
$$p_{t+1/2} = p_t - \frac{\Delta t}{2} \nabla V(q_t)$$
$$q_{t+1} = q_t + \Delta t , p_{t+1/2}$$
$$p_{t+1} = p_{t+1/2} - \frac{\Delta t}{2} \nabla V(q_{t+1})$$
この変換のヤコビアン行列式は $\det(J) \equiv 1.000000$ (Liouville の定理) を満たすため、位相空間ボリュームはシーケンスの深さ全体にわたって厳密に保存されます。フォワードパス中に情報は作成も破壊もされません。
2. IgnoranceGate & OpinionAnchor
無限世代にわたって言語表現を安定させるには

ゾーン:
IgnoranceGate: 位相空間を乱す前にトークン ノイズを減衰する、学習された滑らかなシグモイド フィルター $x' = x \odot \sigma(W_g x)$ です。
OpinionAnchor: 正準位相アトラクターとして機能する実行指数移動平均 $a_t = (1-\eta)a_{t-1} + \eta \cdot \bar{x}'$。
私たちの評価は、同一のトレーニング制約 (シード、オプティマイザー、トークン バジェット、パラメーター スケール) の下での標準 GPT スタイルのトランスフォーマー (BaselineLM) との厳密な同一比較に焦点を当てています。
1. 言語モデリングとスケーリングの法則 (TinyStories、最大 3 億スケール)
セットアップ: GPT-2 BPE トークナイザー (語彙: 50,257)、seq_len=512、Tesla T4 で 10,000 ステップ用にトレーニング済み。
分析: TUPOI は 10,000 ステップにわたって確実に収束し、Transformer ベースラインよりも 14% 少ないパラメータを利用しながら、優れた検証の複雑度 (46.39 vs 50.80) を達成します。べき乗則の適合により、TUPOI が容量の飽和に達することなく予測どおりに拡張されることが確認されます。
2. 自己回帰状態のメモリ フットプリント ($O(1)$ 対 $O(N)$)
セットアップ: トークンごとの順次自動回帰生成 (FP16) 中のコンテキスト メモリのフットプリント。
分析: 標準トランスフォーマーは $O(N)$ KV キャッシュを蓄積し、64k コンテキストで 6.0 GB に達します。 TUPOI は厳密に不変の 6.00 KB 状態ベクトルを維持し、コンシューマ ハードウェアでの無限のストリーミング生成を可能にします。
3. ハードウェア ピーク VRAM ベンチマーク (NVIDIA Tesla T4、FP16)
セットアップ: torch.cuda.max_memory_allocated() を介して直接測定されたピーク割り当て VRAM (MB) を評価するフォワードパス ベンチマーク。
分析: TUPOI は、すべてのシーケンス長にわたって 400 MB ～ 1.2 GB の低いピーク メモリ割り当てで動作します。 8,192 個のトークン コンテキスト全体で、TUPOI の合計アクティベーション デルタは、静的重みベース (約 1.21 GB) をわずか約 500 MB 上回るだけです。
科学的な厳密性を維持するために、現時点では次の制限があることに注意してください。
コンテキストウィンドウのスケール:

300M モデルは、最大 8,192 個のトークン ハードウェア コンテキストまで検証されています。 Extreme Horizo​​n ( $>128\text{k}$ ) には、専用の Flash-Verlet CUDA カーネルが必要です。
ダウンストリーム タスク ベンチマーク: 推論ベンチマーク (GSM8k、MetaMathQA) の監視付き微調整が現在進行中です。
数十億規模のスケーリング: 10 億以上のパラメーター領域を超えたべき乗則の動作を検証することが、引き続き次の反復の主な目的です。
このリポジトリには、上記で報告された調査結果を再現するために必要なスクリプトがすべて含まれています。
git clone https://github.com/narelabs/TUPOI.git
cd ツポイ
pip install -r 要件.txt
2. インタラクティブテキスト生成 CLI
Python スクリプト/generate_cli.py
3. ハードウェア VRAM ベンチマークを実行する
Python ベンチマーク/benchmark_vram.py
4. ベンチマークの視覚化を生成する
Python スクリプト/plot_benchmarks.py
Python スクリプト/generate_dual_axis_charts.py
📄 引用
@article { tupoi2026 ,
title = { TUPOI: O(1) メモリを備えたシンプレクティック ポストトランスフォーマー言語モデル } ,
著者 = { NARE LABS (15 歳で作成) } ,
年 = { 2026 } 、
ジャーナル = { arXiv プレプリント } 、
URL = { https://github.com/narelabs/TUPOI }
}
⚖️ライセンス
MIT ライセンスに基づいて配布されます。
TUPOI: O(1) メモリ (KV キャッシュなし) を備えたシンプレクティックなポストトランスフォーマー言語モデル。 15歳までに建てられた独立した研究者。
ハグフェイス.co/DanilKZ/TUPOI-1 トピックス
Readme アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

TUPOI: Symplectic Post-Transformer Language Model with O(1) Memory (No KV-Cache). Built by 15 y.o. independent researcher. - narelabs/TUPOI

GitHub - narelabs/TUPOI: TUPOI: Symplectic Post-Transformer Language Model with O(1) Memory (No KV-Cache). Built by 15 y.o. independent researcher. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
narelabs
/
TUPOI
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits assets assets benchmarks benchmarks scripts scripts tupoi tupoi .gitignore .gitignore README.md README.md requirements.txt requirements.txt setup.py setup.py View all files Repository files navigation
TUPOI: Symplectic Post-Transformer Language Model with O(1) Memory
A sub-quadratic, attention-free sequence modeling architecture that replaces dense attention matrices with a Symplectic Hamiltonian Integrator (Velocity-Verlet).
Author: NARE LABS ( Built by 15 y.o. )
This repository contains the reference implementation, pre-trained weights, and empirical benchmark suite for the TUPOI language model, a post-transformer architecture that completely eliminates the Attention mechanism and the persistent KV-cache in favor of continuous Hamiltonian phase-space dynamics.
The primary objective of this research is to investigate whether modeling sequence context as particles moving through a symplectic phase space $(q, p)$ can match the expressivity of standard Transformer models while strictly guaranteeing $O(1)$ memory consumption and zero state dissipation ( $\det(J) \equiv 1.0$ ).
TUPOI replaces $O(N^2)$ pairwise Attention with a deep symplectic transition system:
1. The Symplectic Hamiltonian Core
The hidden state is decomposed into canonical coordinates $(q, p)$ and integrated through 24 layers via the Velocity-Verlet Leapfrog dynamical scheme:
$$p_{t+1/2} = p_t - \frac{\Delta t}{2} \nabla V(q_t)$$
$$q_{t+1} = q_t + \Delta t , p_{t+1/2}$$
$$p_{t+1} = p_{t+1/2} - \frac{\Delta t}{2} \nabla V(q_{t+1})$$
Because the Jacobian determinant of this transformation satisfies $\det(J) \equiv 1.000000$ (Liouville's Theorem), the phase-space volume is strictly conserved across sequence depth. Information is neither created nor destroyed during the forward pass.
2. IgnoranceGate & OpinionAnchor
To stabilize language representations over infinite generation horizons:
IgnoranceGate: A learned smooth sigmoid filter $x' = x \odot \sigma(W_g x)$ that attenuates token noise before perturbing phase space.
OpinionAnchor: A running exponential moving average $a_t = (1-\eta)a_{t-1} + \eta \cdot \bar{x}'$ acting as a canonical phase attractor.
Our evaluations focus on rigorous, apples-to-apples comparisons against a standard GPT-style Transformer (BaselineLM) under identical training constraints (seed, optimizer, token budget, parameter scale).
1. Language Modeling & Scaling Laws (TinyStories, ~300M Scale)
Setup: GPT-2 BPE tokenizer (vocab: 50,257), seq_len=512, trained for 10,000 steps on Tesla T4.
Analysis: TUPOI converges reliably across 10,000 steps and achieves superior validation perplexity ( 46.39 vs 50.80 ) while utilizing 14% fewer parameters than the Transformer baseline. The power-law fit confirms that TUPOI scales predictably without hitting capacity saturation.
2. Autoregressive State Memory Footprint ( $O(1)$ vs $O(N)$)
Setup: Context memory footprint during sequential autoregressive token-by-token generation (FP16).
Analysis: Standard Transformers accumulate an $O(N)$ KV-cache reaching 6.0 GB at 64k context. TUPOI maintains a strictly invariant 6.00 KB state vector , allowing infinite streaming generation on consumer hardware.
3. Hardware Peak VRAM Benchmark (NVIDIA Tesla T4, FP16)
Setup: Forward-pass benchmark evaluating peak VRAM allocated (MB) measured directly via torch.cuda.max_memory_allocated() .
Analysis: TUPOI operates with 400 MB to 1.2 GB lower peak memory allocation across all sequence lengths. Across an 8,192 token context, TUPOI's total activation delta is merely ~500 MB above the static weights base (~1.21 GB).
To maintain scientific rigor, we note the following current limitations:
Context Window Scale: The 300M model has been validated up to 8,192 token hardware contexts. Extreme horizons ( $&gt;128\text{k}$ ) require dedicated Flash-Verlet CUDA kernels.
Downstream Task Benchmarks: Supervised fine-tuning on reasoning benchmarks (GSM8k, MetaMathQA) is currently in progress.
Multi-Billion Scaling: Validating power-law behavior beyond 1B+ parameter regimes remains the primary objective for next iterations.
This repository contains all necessary scripts to reproduce the findings reported above.
git clone https://github.com/narelabs/TUPOI.git
cd TUPOI
pip install -r requirements.txt
2. Interactive Text Generation CLI
python scripts/generate_cli.py
3. Run Hardware VRAM Benchmark
python benchmarks/benchmark_vram.py
4. Generate Benchmark Visualizations
python scripts/plot_benchmarks.py
python scripts/generate_dual_axis_charts.py
📄 Citation
@article { tupoi2026 ,
title = { TUPOI: Symplectic Post-Transformer Language Model with O(1) Memory } ,
author = { NARE LABS (Built by 15 y.o.) } ,
year = { 2026 } ,
journal = { arXiv preprint } ,
url = { https://github.com/narelabs/TUPOI }
}
⚖️ License
Distributed under the MIT License .
TUPOI: Symplectic Post-Transformer Language Model with O(1) Memory (No KV-Cache). Built by 15 y.o. independent researcher.
huggingface.co/DanilKZ/TUPOI-1 Topics
Readme Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
