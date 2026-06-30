---
source: "https://github.com/sajjaddoda72-design/UATC"
hn_url: "https://news.ycombinator.com/item?id=48734992"
title: "UATC – A Closed-Loop Controller to Prevent GPU OOM During LLM Training"
article_title: "GitHub - sajjaddoda72-design/UATC · GitHub"
author: "L_u_u_6"
captured_at: "2026-06-30T16:43:31Z"
capture_tool: "hn-digest"
hn_id: 48734992
score: 1
comments: 0
posted_at: "2026-06-30T16:22:32Z"
tags:
  - hacker-news
  - translated
---

# UATC – A Closed-Loop Controller to Prevent GPU OOM During LLM Training

- HN: [48734992](https://news.ycombinator.com/item?id=48734992)
- Source: [github.com](https://github.com/sajjaddoda72-design/UATC)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T16:22:32Z

## Translation

タイトル: UATC – LLM トレーニング中の GPU OOM を防ぐ閉ループ コントローラー
記事タイトル: GitHub - sajjaddoda72-design/UATC · GitHub
説明: GitHub でアカウントを作成して、sajjaddoda72-design/UATC 開発に貢献します。

記事本文:
GitHub - sajjaddoda72-design/UATC · GitHub
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
sajjaddoda72-デザイン
/
UATC
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く Fo

フォルダとファイル
48 コミット 48 コミット .github/ workflows .github/ workflows テスト テスト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md UATC.py UATC.py pyproject.toml pyproject.toml 要件.txt 要件.txt uatc_dashboard.jpg uatc_dashboard.jpg すべてのファイルを表示 リポジトリ ファイルのナビゲーション
UATC: ユニバーサル アダプティブ トレーニング コントローラー
リソースに制約のあるエッジ ハードウェア上で大規模言語モデル (LLM) を微調整するのは脆弱です。単一の長いシーケンスや予期しないバッチサイズのスパイクにより、メモリ不足 (OOM) クラッシュが引き起こされ、何時間ものコンピューティングが無駄になる可能性があります。静的構成は、動的なコンテキストの長さ、アクティベーション キャッシュ、および勾配の蓄積から生じる揮発性メモリのプレッシャーに反応できません。この論文では、LLM トレーニングを動的な産業プロセスとして扱う閉ループ制御システムである UATC (ユニバーサル アダプティブ トレーニング コントローラー) について説明します。 UATC は、ノイズ耐性のある状態推定のためのカルマン フィルター、フィードバック調整のためのアンチワインドアップを備えた PID コントローラー、遅延補償のためのスミス予測器、ヒステリシスのための 3 状態シュミット トリガー、位相認識動的データ プルーナー、および OOM および NaN/Inf イベント用の階層型回復サブシステムを融合しています。コントローラーはパラダイムを認識しており、コードを変更することなく、その弾性ゲインと回復しきい値を完全な微調整、LoRA/PEFT、および QLoRA ワークロードに適応させます。 NVIDIA T4 GPU (15 GB VRAM) 上で UATC を評価し、QLoRA を使用して意図的に混雑したメモリ環境下で Qwen2.5-1.5B-Instruct を微調整しました。 5 つの制御された実行 (フル コントローラー、3 つの単一サブシステム アブレーション、DeepSpeed スタイルの静的ベースライン) にわたって、UATC は 300 のトレーニング ステップすべてを完了し、同時に冗長なサンプル パスの 86.22% を動的にプルーニングします (合計 3,999 のサンプル パスから 3,448 個のサンプルをプルーニング)。

)、致命的なクラッシュを 1 つも発生させることなく、18 件の EMERGENCY_OOM イベントから回復しました。 18 の回復可能なイベントは、2 つの強制メモリ ショック (ステップ 50 での 1024 トークンのコンテキスト ショックとステップ 120 での 48 サンプルのバッチ ショック) と、このモデル構成で PID が T4 の BS=24 ハードウェア上限を超えようとしたときにトリガーされた 16 の物理 OOM イベントで構成されます。すべてのイベントが捕捉され、吸収され、正常に回復されました。 DeepSpeed スタイルのベースライン (勾配チェックポイント処理常時オン、バッチ サイズ 8、各ステップの empty_cache) は、大きな frac を保持しているにもかかわらず、最初の衝撃で致命的にクラッシュします。
[切り捨てられた]
生成 AI のエッジ展開はメモリがボトルネックになります。 QLoRA およびその他のパラメーター効率の高いメソッドは、1B ～ 8B パラメーター モデルの静的フットプリントをコンシューマー GPU の 16 GB 予算内に収めますが、トレーニング中の動的なフットプリント (アクティベーション、アテンション キャッシュ、勾配アキュムレーター、オプティマイザーの状態) は非常に不安定なままです。静的なバッチ サイズと学習率はこの変動性を無視しており、単一の外れ値シーケンスによって数時間の実行がクラッシュする可能性があります。
私たちは、これは基本的に閉ループ制御の問題であると主張します。 UATC は、PyTorch トレーニング ステップをラップし、GPU メモリの負荷、損失の動作、およびシーケンス長の統計を観察する薄いオーケストレーション レイヤーです。各ステップで、ターゲット バッチ サイズ、学習率、AMP 切り替え、勾配チェックポイント切り替え、枝刈り率などのアクションを返します。コントローラーは完全にデバイス上で実行され、単一の GPU 以外のインフラストラクチャを必要とせず、モデルのサイズ、データセット、トレーニング パラダイムに依存しません。
貢献。この研究は以下に貢献します。
単一の構成サーフェスを通じて FPFT、PEFT、および QLoRA ワークロードに適応するパラダイムを認識した閉ループ コントローラー。
4 つの状態のフェーズ マシン (ウォームアップ、スケーリング、コンバージェンス、リカバリ)

探査体制と開発体制の間の変動を防ぎます。
カルマン + PID + スミス推定器を組み合わせたもので、測定ノイズ、フィードバック調整、プロセスのデッドタイムを個別の関心事として処理します。
一時的な NaN/Inf (ソフト リカバリ) と持続的な OOM (ハード リダクション + リカバリ フェーズ) を区別する階層型リカバリ サブシステム。
1.5B モデルでは、静的パイプラインに 5 倍の空きメモリがある場合でも、動的な閉ループ制御が静的パイプラインよりも優れたパフォーマンスを発揮するという経験的証拠があります。
厳密な収束議論: プルーナーとフェーズの相互作用は、残存損失がハード サンプルのみに対して計算されることを数学的に保証し、高い残留損失を故障モードではなく学習の肯定的なシグナルにします。
動的バッチサイズのスケーリングは、勾配ノイズ スケールと大バッチの経験的モデルのレンズを通じて研究されてきましたが、これらの方法はオーバープロビジョニングされたハードウェアを前提としており、マルチノード クラスターで評価されます。 DeepSpeed や Megatron-LM などの分散フレームワークは、静的に、または集中スケジューリングを通じてメモリを最適化します。単一デバイスではステップごとのマイクロ適応は実行されません。
制御理論には、Web サーバー アドミッション コントロール、CPU 周波数スケーリング、キューベースのリソース管理など、コンピューティング システムにおける長い歴史があります。ニューラル ネットワーク トレーニングの内部ループへの応用、特にカルマン フィルター、PID ループ、スミス予測器、および VRAM 制約付きエッジ トレーニング用のサンプルごとの損失フィルターのハイブリッド使用については、これまで文書化されていませんでした。
カリキュラム学習やコアセットの選択などのデータ プルーニング手法は、データセット レベルで機能します。 UATC のプルーナーは、現在のトレーニング フェーズを使用してアクティブなしきい値を設定し、バッチごとに動的に動作します。
UATC は単一のエントリ ポイントを公開します:controller.dec

de(状態)→動作。状態は、トレーニング ステップからのテレメトリ スナップショット (損失、VRAM 圧力、勾配ノルム、シーケンス長統計、OOM フラグ) です。アクションは、実行ディレクティブ (バッチ サイズ、学習率、AMP、チェックポイント、プルーニング率) のバンドルです。
┌───────────────────┐
│ UATC コントローラー │
│ │
状態 ──────▶│ カルマン → PID → スミス → フェーズマシン │──────▶ アクション
│ ↘ ↘ │ (BS、LR、
│ シュミットトリガデータプルーナ │ AMP、CKPT、
│ │ プルーン）
━━━━━━━━━━━━━━━━━━━┘
コントローラーは、内部で以下に説明する 8 つの相互作用するサブシステムに分解されます。
コントローラーは 4 つの標準フェーズを循環します。遷移はシングルステップ イベントではなく持続条件でゲートされるため、位相境界でのちらつきが防止されます。
設定可能なステップ数の後に、WARMUP は終了して SCALING に戻ります。 SCALING は、持続的なプラトー (連続ステップにわたる |loss_velocity| が小さい) またはタイムアウト後に終了して CONVERGENCE になります。 RECOVERY は、パラダイム固有の数のクリーン ステップを実行し、プルーニング率が最小最高水準点に達した後にのみ、SCALING に移行します。
回復出口ゲート。 RECOVERY → SCALING への移行は、T4 + Qwen2.5-1.5B の実行で不可欠であることが証明された 2 つのノブによって条件付けされます。
リカバリ_アクセル_ノイズ_ゲート: フロート = 5e-3 。カルマン平滑化メモリ加速度 $a_k$ は安定性信号として使用されます。$|a_k|$ がこのゲートを下回っている限り、コントローラーはシステムを静止状態として扱い、RECOVERY を終了する可能性があります。ゲートは、PyTorch の自然なアロケーター ノイズ フロアに合わせて調整されています。

キャッシュ アロケータ。各ステップで約 $1 \times 10^{-3}$ から $5 \times 10^{-3}$ の残留加速を生成します (勾配の累積と autograd グラフの再割り当て中のテンソル存続期間の変動によって駆動されます)。ゲートを 5e-3 に設定すると、コントローラーは本物のサージ イベントに対する感度を維持しながら、自然なアロケーター ノイズ フロアの下でクリーンな終了が可能になります。
リカバリ_セーフ_フォールバック_ステップ: int = 12 。冗長安全弁として、加速ゲートがまだクリアされていない場合でも、コントローラーは OOM なしで Recovery_safe_fallback_steps のクリーン ステップが経過すると、RECOVERY → SCALING に移行することが許可されます。これにより、プルーナーのみが主要な救済手段ではない状況での前進が保証されます (たとえば、実行中、コントローラーは回復のための剪定ではなく、主に物理ハードウェアの天井を介して BS ノコギリ波に乗っています)。そのため、あらゆる回復状況下での滞留時間の絶対的な下限は、このカウンターによって制限されます。
3.2 状態推定(カルマンフィルター)
GPU メモリ テレメトリには、非同期 CUDA 実行、アロケータの断片化、およびドライバ レベルのキャッシュが原因でノイズが多くなります。 UATC は、生の圧力信号 $z_k \in [0, 1]$ に対して離散カルマン フィルターを実行します。
$$
\hat{x}^{-}_{k} = \hat{x}_{k-1}
$$
$$
K_k = \frac{P^{-}_{k}}{P^{-}_{k} + R}
$$
$$
\hat{x}_{k} = \hat{x}^{-}_{k} + K_k \left( z_k - \hat{x}^{-}_{k} \right)
$$
$$
P_k = (1 - K_k)、P^{-}_{k}
$$
デフォルト値: $Q = 1 \times 10^{-4}$ (プロセスノイズ)、$R = 5 \times 10^{-4}$ (測定ノイズ)。平滑化された状態 $\hat{x}_k$ は、PID 導関数項と動的設定値バックオフをフィードする指数移動平均 (EMA) を通じてメモリ速度と加速度を計算するために使用されます。
PID 誤差は、動的設定値と平滑化されたメモリ圧力の差です。設定値

メモリ アクセラレーションが正の場合、ソフト リミットから後退します。
$$
e_k = \text{設定値}_k - z_k
$$
出力は標準の加重合計です。
$$
u_k = K_p \cdot e_k + K_i \cdot I_k + K_d \cdot D_k
$$
ここで、 $I_k$ は $e_k$ ( $|I_k| \leq 0.4$ を含む) の固定された累計であり、 $D_k$ は EMA でフィルター処理された負の記憶速度です。
動的なゲインスケーリング。 3 つのベース ゲイン $K_{p,\text{base}} = 0.50$ 、 $K_{i,\text{base}} = 0.003$ 、 $K_{d,\text{base}} = 0.04$ は、ステップごとに 3 つの乗算係数によってスケーリングされます。
ここで、$\text{span} = \text{HardLimit} - \text{SoftLimit}$ および $\xi$ は §3.7 のパラダイム弾力性です。最終ゲインは、基本ゲインに $\text{capacity\_scale} \cdot \{\text{gain}\}$ を乗算します。積分項では条件付き積分 (アンチワインドアップ) が使用されます。PID 出力がクランプ限界に対して飽和すると、積分器は蓄積されるのではなく $(\text{unclamped} - \text{clamped}) \cdot \text{anti\_windup\_gain}$ に比例する量だけ逆方向に流され、飽和後のオーバーシュートを引き起こす積分ワインドアップを防ぎます。
3.4 遅延補償（スミス予測器）
カーネルの起動遅延と非同期アロケーターのコミットにより、バッチ サイズの変更はいくつかのステップの VRAM テレメトリに反映されません。スミス予測器は、最新の PID 出力の FIFO バッファーを維持し、最も古いバッファー値と現在の負の速度から導出される有界補正項を追加することによって補正します。
$$
c_

[切り捨てられた]

## Original Extract

Contribute to sajjaddoda72-design/UATC development by creating an account on GitHub.

GitHub - sajjaddoda72-design/UATC · GitHub
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
sajjaddoda72-design
/
UATC
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
48 Commits 48 Commits .github/ workflows .github/ workflows tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md UATC.py UATC.py pyproject.toml pyproject.toml requirements.txt requirements.txt uatc_dashboard.jpg uatc_dashboard.jpg View all files Repository files navigation
UATC: Universal Adaptive Training Controller
Fine-tuning Large Language Models (LLMs) on resource-constrained edge hardware is brittle. A single long sequence or an unexpected batch-size spike can trigger an Out-Of-Memory (OOM) crash and waste hours of compute. Static configurations cannot react to the volatile memory pressure that arises from dynamic context lengths, activation caching, and gradient accumulation. This paper presents UATC (Universal Adaptive Training Controller) , a closed-loop control system that treats LLM training as a dynamic industrial process. UATC fuses a Kalman filter for noise-resilient state estimation, a PID controller with anti-windup for feedback regulation, a Smith predictor for delay compensation, three-state Schmitt triggers for hysteresis, a phase-aware dynamic data pruner, and a tiered recovery subsystem for OOM and NaN/Inf events. The controller is paradigm-aware: it adapts its elasticity gains and recovery thresholds to full fine-tuning, LoRA/PEFT, and QLoRA workloads without code changes. We evaluate UATC on an NVIDIA T4 GPU (15 GB VRAM) fine-tuning Qwen2.5-1.5B-Instruct under a deliberately congested memory environment using QLoRA. Across five controlled runs (full controller, three single-subsystem ablations, and a DeepSpeed-style static baseline), UATC completes all 300 training steps while dynamically pruning 86.22% of redundant sample-passes (3,448 samples pruned out of 3,999 total sample-passes) and recovering from 18 EMERGENCY_OOM events without a single fatal crash. The 18 recoverable events comprise 2 forced memory shocks (a 1024-token context shock at step 50 and a 48-sample batch shock at step 120) and 16 physical OOM events triggered when the PID attempted to push past the BS=24 hardware ceiling of the T4 for this model configuration; every event was caught, absorbed, and recovered from gracefully. The DeepSpeed-style baseline (gradient checkpointing always-on, batch size 8, empty_cache every step) crashes fatally at the first shock despite holding a large frac
[truncated]
Edge deployment of generative AI is bottlenecked by memory. QLoRA and other parameter-efficient methods bring the static footprint of 1B–8B parameter models within the 16 GB budget of consumer GPUs, but the dynamic footprint during training — activations, attention caches, gradient accumulators, optimizer states — remains highly volatile. Static batch size and learning rate are blind to this volatility, and a single outlier sequence can crash a multi-hour run.
We argue this is fundamentally a closed-loop control problem. UATC is a thin orchestration layer that wraps the PyTorch training step and observes GPU memory pressure, loss behavior, and sequence-length statistics. Each step it returns an action: target batch size, learning rate, AMP toggle, gradient-checkpointing toggle, and pruning ratio. The controller runs entirely on-device, requires no infrastructure beyond a single GPU, and is agnostic to model size, dataset, and training paradigm.
Contributions. This work contributes:
A paradigm-aware closed-loop controller that adapts to FPFT, PEFT, and QLoRA workloads through a single configuration surface.
A four-state phase machine (WARMUP, SCALING, CONVERGENCE, RECOVERY) that prevents oscillation between exploration and exploitation regimes.
A combined Kalman + PID + Smith estimator that handles measurement noise, feedback regulation, and process dead-time as separate concerns.
A tiered recovery subsystem that distinguishes between transient NaN/Inf (soft recovery) and sustained OOM (hard reduction + RECOVERY phase).
Empirical evidence on a 1.5B model that dynamic, closed-loop regulation outperforms static pipelines even when the static pipeline has 5× more free memory.
A rigorous convergence argument: the pruner–phase interaction mathematically guarantees that the surviving loss is computed exclusively over hard samples, making high residual loss a positive signal of learning rather than a failure mode.
Dynamic batch-size scaling has been studied through the lens of gradient noise scale and large-batch empirical models, but these methods assume over-provisioned hardware and are evaluated on multi-node clusters. Distributed frameworks such as DeepSpeed and Megatron-LM optimize memory statically or through centralized scheduling; they do not perform per-step micro-adaptation on a single device.
Control theory has a long history in computing systems: web server admission control, CPU frequency scaling, and queueing-based resource management. Its application to the inner loop of neural network training — particularly the hybrid use of Kalman filtering, PID loops, Smith predictors, and per-sample loss filtering for VRAM-constrained edge training — has not been previously documented.
Data pruning methods such as curriculum learning and coreset selection operate at the dataset level. UATC's pruner operates per-batch, dynamically, using the current training phase to set the active threshold.
UATC exposes a single entry point: controller.decide(state) → action . The state is a telemetry snapshot from the training step (loss, VRAM pressure, gradient norm, sequence-length statistics, OOM flag). The action is a bundle of execution directives (batch size, learning rate, AMP, checkpointing, pruning ratio).
┌──────────────────────────────────────────┐
│ UATC Controller │
│ │
state ──────▶│ Kalman → PID → Smith → Phase Machine │──────▶ action
│ ↘ ↘ │ (BS, LR,
│ Schmitt Trigger Data Pruner │ AMP, CKPT,
│ │ Prune)
└──────────────────────────────────────────┘
The controller is internally decomposed into eight interacting subsystems, described below.
The controller cycles through four canonical phases. Transitions are gated on sustained conditions rather than single-step events, which prevents flicker at phase boundaries.
WARMUP exits to SCALING after a configurable number of steps. SCALING exits to CONVERGENCE either on a sustained plateau (small |loss_velocity| over consecutive steps) or after a timeout. RECOVERY exits to SCALING only after a paradigm-specific number of clean steps with the pruning ratio having reached a minimum high-water mark.
RECOVERY exit gating. The RECOVERY → SCALING transition is conditioned on two knobs that proved essential on the T4 + Qwen2.5-1.5B run:
recovery_accel_noise_gate: float = 5e-3 . The Kalman-smoothed memory acceleration $a_k$ is used as a stability signal: as long as $|a_k|$ remains below this gate, the controller treats the system as quiescent and may exit RECOVERY. The gate is calibrated to the natural allocator noise floor of PyTorch's caching allocator, which produces a residual acceleration of approximately $1 \times 10^{-3}$ to $5 \times 10^{-3}$ on every step (driven by tensor lifetime churn during gradient accumulation and the autograd graph re-allocation); setting the gate at 5e-3 keeps the controller sensitive to genuine surge events while allowing clean exits under the natural allocator noise floor.
recovery_safe_fallback_steps: int = 12 . As a redundant safety valve, even when the acceleration gate has not yet cleared, the controller is permitted to transition RECOVERY → SCALING once recovery_safe_fallback_steps clean steps have elapsed without an OOM. This guarantees forward progress in regimes where the pruner alone is not the dominant relief lever (e.g. mid-run the controller is mostly riding the BS sawtooth through the physical hardware ceiling rather than pruning to recover), so the absolute lower bound on dwell time under any recovery regime is bounded by this counter.
3.2 State Estimation (Kalman Filter)
GPU memory telemetry is noisy due to asynchronous CUDA execution, allocator fragmentation, and driver-level caching. UATC runs a discrete Kalman filter on the raw pressure signal $z_k \in [0, 1]$ .
$$
\hat{x}^{-}_{k} = \hat{x}_{k-1}
$$
$$
K_k = \frac{P^{-}_{k}}{P^{-}_{k} + R}
$$
$$
\hat{x}_{k} = \hat{x}^{-}_{k} + K_k \left( z_k - \hat{x}^{-}_{k} \right)
$$
$$
P_k = (1 - K_k), P^{-}_{k}
$$
Default values: $Q = 1 \times 10^{-4}$ (process noise), $R = 5 \times 10^{-4}$ (measurement noise). The smoothed state $\hat{x}_k$ is used to compute memory velocity and acceleration through Exponential Moving Averages (EMAs), which feed the PID derivative term and the dynamic setpoint backoff.
The PID error is the difference between a dynamic setpoint and the smoothed memory pressure. The setpoint backs off from the soft limit when memory acceleration is positive:
$$
e_k = \text{Setpoint}_k - z_k
$$
The output is the standard weighted sum:
$$
u_k = K_p \cdot e_k + K_i \cdot I_k + K_d \cdot D_k
$$
where $I_k$ is the clamped running sum of $e_k$ (with $|I_k| \leq 0.4$ ) and $D_k$ is the EMA-filtered negative memory velocity:
Dynamic gain scaling. The three base gains $K_{p,\text{base}} = 0.50$ , $K_{i,\text{base}} = 0.003$ , $K_{d,\text{base}} = 0.04$ are scaled by three multiplicative factors each step:
where $\text{span} = \text{HardLimit} - \text{SoftLimit}$ and $\xi$ is the paradigm elasticity from §3.7. The final gains multiply the base gains by $\text{capacity\_scale} \cdot \{\text{gain}\}$ . The integral term uses conditional integration (anti-windup): when the PID output saturates against a clamp bound, the integrator is bled in the opposite direction by an amount proportional to $(\text{unclamped} - \text{clamped}) \cdot \text{anti\_windup\_gain}$ rather than being accumulated, preventing integral windup that would otherwise cause post-saturation overshoot.
3.4 Delay Compensation (Smith Predictor)
Batch-size changes do not reflect in VRAM telemetry for several steps due to kernel launch latency and asynchronous allocator commits. The Smith predictor compensates by maintaining a FIFO buffer of the most recent PID outputs and adding a bounded correction term derived from the oldest buffered value and the current negative velocity:
$$
c_

[truncated]
