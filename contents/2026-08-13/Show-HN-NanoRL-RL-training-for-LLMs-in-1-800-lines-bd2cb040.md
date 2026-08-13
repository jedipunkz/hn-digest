---
source: "https://github.com/alex000kim/nanoRL"
hn_url: "https://news.ycombinator.com/item?id=49286216"
title: "Show HN: NanoRL – RL training for LLMs in ~1,800 lines"
article_title: "GitHub - alex000kim/nanoRL: One RL training loop that runs CartPole on a laptop and disaggregated async RLVR on a GPU cluster. About 1,800 lines across 7 files. No Ray, no TRL, no DeepSpeed. It's small enough to read in an afternoon, and it's meant to be forked, not imported. · GitHub"
author: "alex000kim"
captured_at: "2026-08-13T14:15:46Z"
capture_tool: "hn-digest"
hn_id: 49286216
score: 7
comments: 0
posted_at: "2026-08-13T14:08:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: NanoRL – RL training for LLMs in ~1,800 lines

- HN: [49286216](https://news.ycombinator.com/item?id=49286216)
- Source: [github.com](https://github.com/alex000kim/nanoRL)
- Score: 7
- Comments: 0
- Posted: 2026-08-13T14:08:54Z

## Translation

タイトル: Show HN: NanoRL – LLM 向けの RL トレーニング (約 1,800 行)
記事のタイトル: GitHub - alex000kim/nanoRL: ラップトップ上で CartPole を実行し、GPU クラスター上で分離された非同期 RLVR を実行する 1 つの RL トレーニング ループ。 7 ファイルで約 1,800 行。 RayもTRLもDeepSpeedもありません。午後に読むのに十分な量であり、輸入されるものではなく、フォークされることを目的としています。 · GitHub
説明: ラップトップ上で CartPole を実行し、GPU クラスター上で非集約非同期 RLVR を実行する 1 つの RL トレーニング ループ。 7 ファイルで約 1,800 行。 RayもTRLもDeepSpeedもありません。午後に読むのに十分な量であり、輸入されるものではなく、フォークされることを目的としています。 - アレックス000キム/nanoRL
HN テキスト: 私が作成できる最小の非同期 RL トレーナー: ラップトップ上の CartPole で REINFORCE を実行し、クラスター上で非同期 GRPO を実行する 1 つのループ
(例: 8xH100 トレーナー、8 つの vLLM ワーカー、k8s 上で [SkyPilot ジョブ グループ]( https://docs.skypilot.ai/en/latest/examples/job-groups.html ) として実行されました)。 Ray、TRL、DeepSpeed などを使用しない場合、ワーカーは stdlib HTTP 経由でトレーナーと通信します。

記事本文:
GitHub - alex000kim/nanoRL: ラップトップ上で CartPole を実行し、GPU クラスター上で分離された非同期 RLVR を実行する 1 つの RL トレーニング ループ。 7 ファイルで約 1,800 行。 RayもTRLもDeepSpeedもありません。午後に読むのに十分な量であり、輸入されるものではなく、フォークされることを目的としています。 · GitHub
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
アレックス000キム
/
ナノRL
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット

.github/ workflows .github/ workflows 資産 アセット 設定 設定 結果 結果 スカイ スカイ テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md algos.py algos.py core.py core.py model.py model.pyserve.pyserve.py task.py task.py train.py train.py utils.py utils.py すべてのファイルを表示リポジトリ ファイルのナビゲーション
ラップトップ上で CartPole を実行し、GPU 上で非集約非同期 RLVR を実行する 1 つの RL トレーニング ループ
クラスター。 7 ファイルで約 1,800 行。 RayもTRLもDeepSpeedもありません。十分小さいです
午後に読むために、そして、の精神で
nanoGPT は、インポートされるのではなく、フォークされることを目的としています。
損失 = - (利点 * logprob)。平均 () # + PPO 比率クリップ、+ オプションの KL
それがアップデート全体です。 REINFORCE、PPO、GRPO、および RLOO は、利点が異なるだけです(...)
が計算され、同期と非同期の違いはバッチがどこから来たのかだけです。それ以外はすべて
リポジトリは、この 1 行を正しくフィードするために存在します。
アルゴリズム動物園は algos.py です: 生の戻り値、一定のベースライン、GAE
批評家とグループのベースライン (GRPO、RLOO)。新しいアルゴリズムは 1 つの機能です。
非同期トレーニングは、同じループをマシン間で分割します。 1 つのボックスで --role Trainer を実行し、
--その他の役割ロールアウト。彼らは stdlib HTTP (serve.py) を介して通信します。
労働者が死亡した場合、その飛行中のバッチのみが費用としてかかります。それ以外は何もかかりません。ノブが一つあるのですが、
max_staleness 、拒否限界、キューの深さ、および重みの数を設定します。
トレーナーが保持するスナップショット。
トレーナーは、他の場所で計算された logprobs を信頼しません。 old_logp を再計算します。
各バッチをサンプリングした正確な重み (重みがなくなったバッチは、
ドロップされ、トレーニングされたことはありません）、測定されたギャップをステップごとに重要度の隣に出力します
比率が低下し、古くなり、キューが減少します。サンプリングにも同じことが当てはまります: HFPolicy
あらゆる世代を中和する

n_config.json ノブの比率はモデル化されません ( top_k 、
繰り返しペナルティ , ...)。
速度の向上: ワーカー上の vLLM、生成とスコアリング用に別個のバッチ サイズ (ワーカーには
反対のメモリプロファイル。それらを接続すると、1 回の測定実行で 6 倍のコストがかかります)、アダプターのみの重量同期
16 GB ではなく約 170 MB で、まったく同じものを生成するマイクロバッチ バックワード
任意の数の GPU 上で 1 つの大きな後方としての勾配。 44 CPU のみのテストにより、これらすべてが明らかになります。
Python テスト/test_nanorl.py 。
Python 3.10+ と uv が必要です。
uv pip install torch Gymnasium numpy トランスフォーマー「データセット<4」peft pyyaml
python train.py --task cartpole --algo reinforce # CPU 上で約 10 秒で解決
ルング
コマンド
ハードウェア
1
--タスクカートポール --アルゴ強化
CPU、秒
2
--config configs/cartpole_ppo.yaml
CPU、約20秒
3
--config configs/countdown_grpo.yaml
1x 24 GB GPU
4
torchrun --nproc_per_node=8 train.py --config configs/countdown_grpo_gemma4.yaml
8×H100
5
sky jobs が sky/jobgroup.yaml を起動します
2 個以上の GPU、k8 のみ (ワーカーはホスト名でトレーナーを検索します)
エンドツーエンドの実行
ラング 5、YAML に同梱されているものとまったく同じ: Qwen3-8B と LoRA がカウントダウンで GRPO を実行します。 8x H100 1 個
ノード トレイン、vLLM で生成される 8 つの L40S GPU、および
SkyPilot ジョブ グループ ワイヤ
それらをホスト名でまとめます。
ギャップローはキープを獲得するものです。 vLLM はさまざまなカーネルで logprob を計算します
トレーナーのHFフォワードよりも。以前の実行では、とにかくそれらを old_logp として使用しました: 比率
古さがゼロの場合は 1.006 から 1.165 にドリフトし、精度は 0.500 から 0.188 に低下しました。今
トレーナーは old_logp 自体を再計算し、同じエンジン上の同じモデルが上昇します。
崩壊する代わりに。
再現するには 2 つの GPU で十分です (トレーナー 1 つ、ワーカー 1 つ)。実行の CSV は次のとおりです
結果/ 。同期パス (ラング 4) では、Gemma-4-12B が 0.312 から 0.688 になりました。
同じタスク。
algos.py ~90 の利点の推定値

rs: 強化 / ベースライン / ppo (GAE) / grpo / rloo
core.py ~110 軌跡 + バッチ
utils.py ~170 マスクされたリダクション、リターン/GAE、dist プラミング、CSV ロガー
serve.py ~250 非同期トランスポート: 重みの公開、ロールアウト キュー、古さの拒否
task.py ~310 CartPole + カウントダウン/GSM8K;タスク = サンプル + ロールアウト + 報酬 FNS
model.py ~330 MLPPolicy、HFPolicy (トレーニング + スコア)、VLLMGenerator (ロールアウトのみ)
train.py ~560 ループ: ロールアウト、アドバンテージ、クリップされた更新、パブリッシュ
あなたのモデル、あなたのデータセット
AutoModelForCausalLM ID は、1 つの GPU に適合する限り、 --model で機能します (
LoRA は 48 GB でおよそ 12B です)。 LoRA は q_proj/k_proj/v_proj/o_proj をターゲットとします。もしあなたの
モデルはその投影に別の名前を付けます。これは、次の 1 行の変更です。
モデル.py 。 vLLM ワーカーには、vLLM がロードできるアーキテクチャが必要です。
データセットは LLMTask のサブクラスです: 行をロードし、 build_prompt を書き込み、報酬を選択します
関数、クラスを make_task (tasks.py) に登録し、実行します
--タスク <名前> :
クラス MyTask ( LLMTask ):
def __init__ ( self 、 n_examples = 2000 、シード = 0 、 n_eval = 128 、ランク = 0 、ワールド = 1 、 ** kw ):
自分自身。ランク、自己判断となります。ワールド = ランク、ワールド
rows = [...] # [{"質問": ..., "値": ...}, ...]
自分自身。 _split ( rows , n_eval ) # eval ホールドアウト + ランクごとのシャーディング
自分自身。報酬_fns = [( gsm8k_reward , 1.0 ), ( format_reward , 0.1 )]
def build_prompt ( self , p ):
return p [ "question" ] + " \n <answer> </answer> タグには最後の数字のみを入力してください。"
報酬関数は (プロンプト、完了、応答) -> float であり、それを何でも呼び出すことができます
正規表現から外部ジャッジまで。
意図的に省略されたもの: メガトロン規模の並列処理、MoE、マルチテナント スケジューリング、
ダッシュボード。それらが必要な場合は、このリポジトリを超えてしまいますが、それが重要です。
CartPole をラップトップ上で実行し、非集約化された 1 つの RL トレーニング ループ

GPU クラスター上の RLVR。 7 ファイルで約 1,800 行。 RayもTRLもDeepSpeedもありません。午後に読むのに十分な量であり、輸入されるものではなく、フォークされることを目的としています。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

One RL training loop that runs CartPole on a laptop and disaggregated async RLVR on a GPU cluster. About 1,800 lines across 7 files. No Ray, no TRL, no DeepSpeed. It's small enough to read in an afternoon, and it's meant to be forked, not imported. - alex000kim/nanoRL

The smallest async RL trainer I could write: one loop that runs REINFORCE on CartPole on a laptop and async GRPO on a cluster
(e.g. 8xH100 trainer, 8 vLLM workers, ran as a [SkyPilot job group]( https://docs.skypilot.ai/en/latest/examples/job-groups.html ) on k8s ). All without Ray or TRL or DeepSpeed etc., workers talk to the trainer over stdlib HTTP.

GitHub - alex000kim/nanoRL: One RL training loop that runs CartPole on a laptop and disaggregated async RLVR on a GPU cluster. About 1,800 lines across 7 files. No Ray, no TRL, no DeepSpeed. It's small enough to read in an afternoon, and it's meant to be forked, not imported. · GitHub
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
alex000kim
/
nanoRL
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github/ workflows .github/ workflows assets assets configs configs results results sky sky tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md algos.py algos.py core.py core.py model.py model.py serve.py serve.py tasks.py tasks.py train.py train.py utils.py utils.py View all files Repository files navigation
One RL training loop that runs CartPole on a laptop and disaggregated async RLVR on a GPU
cluster. About 1,800 lines across 7 files. No Ray, no TRL, no DeepSpeed. It's small enough
to read in an afternoon, and, in the spirit of
nanoGPT , it's meant to be forked, not imported.
loss = - ( advantage * logprob ). mean () # + PPO ratio clip, + optional KL
That's the whole update. REINFORCE, PPO, GRPO and RLOO only differ in how advantage(...)
is computed, and sync vs async only differs in where batches come from. Everything else in
the repo exists to feed this one line correctly.
The algorithm zoo is algos.py : raw returns, a constant baseline, GAE
with a critic, and the group baselines (GRPO, RLOO). A new algorithm is one function.
Async training is the same loop split across machines. Run --role trainer on one box and
--role rollout on the others; they talk over stdlib HTTP ( serve.py ).
A worker that dies costs you its in-flight batch and nothing else. There's one knob,
max_staleness , and it sets the rejection bound, the queue depth, and how many weight
snapshots the trainer keeps.
The trainer doesn't trust logprobs computed anywhere else. It recomputes old_logp under
the exact weights that sampled each batch (a batch whose weights it no longer has is
dropped, never trained on) and prints the measured gap every step, next to the importance
ratio, the staleness, and the queue drops. The same applies to sampling: HFPolicy
neutralizes every generation_config.json knob the ratio doesn't model ( top_k ,
repetition_penalty , ...).
For speed: vLLM on the workers, separate batch sizes for generation and scoring (they have
opposite memory profiles; tying them cost 6x in one measured run), adapter-only weight sync
at ~170 MB instead of 16 GB, and a micro-batched backward that produces exactly the same
gradient as one big backward on any number of GPUs. 44 CPU-only tests pin all of this down:
python tests/test_nanorl.py .
You need Python 3.10+ and uv :
uv pip install torch gymnasium numpy transformers " datasets<4 " peft pyyaml
python train.py --task cartpole --algo reinforce # solves in ~10 s on CPU
Rung
Command
Hardware
1
--task cartpole --algo reinforce
CPU, seconds
2
--config configs/cartpole_ppo.yaml
CPU, ~20 s
3
--config configs/countdown_grpo.yaml
1x 24 GB GPU
4
torchrun --nproc_per_node=8 train.py --config configs/countdown_grpo_gemma4.yaml
8x H100
5
sky jobs launch sky/jobgroup.yaml
2+ GPUs, k8s only (workers find the trainer by hostname)
The end-to-end run
Rung 5, exactly as the YAML ships: Qwen3-8B with LoRA doing GRPO on Countdown. One 8x H100
node trains, eight L40S GPUs generate with vLLM, and a
SkyPilot Job Group wires
them together by hostname.
The gap row is the one that earns its keep. vLLM computes logprobs with different kernels
than the trainer's HF forward. An earlier run used them as old_logp anyway: the ratio
drifted from 1.006 to 1.165 at zero staleness and accuracy fell from 0.500 to 0.188. Now
the trainer recomputes old_logp itself, and the same model on the same engines climbs
instead of collapsing.
Two GPUs are enough to reproduce (one trainer, one worker). The run's CSV is in
results/ . The sync path (rung 4) got Gemma-4-12B from 0.312 to 0.688 on the
same task.
algos.py ~90 advantage estimators: reinforce / baseline / ppo (GAE) / grpo / rloo
core.py ~110 Trajectory + Batch
utils.py ~170 masked reductions, returns/GAE, dist plumbing, CSV logger
serve.py ~250 async transport: weight publishing, rollout queue, staleness rejection
tasks.py ~310 CartPole + Countdown/GSM8K; task = sample + rollout + reward fns
model.py ~330 MLPPolicy, HFPolicy (train + score), VLLMGenerator (rollout only)
train.py ~560 the loop: rollout, advantage, clipped update, publish
Your model, your dataset
Any AutoModelForCausalLM id works with --model , as long as it fits on one GPU (with
LoRA that's roughly 12B on 48 GB). LoRA targets q_proj/k_proj/v_proj/o_proj ; if your
model names its projections differently, that's a one-line change in
model.py . vLLM workers need an architecture vLLM can load.
A dataset is a subclass of LLMTask : load rows, write build_prompt , pick reward
functions, register the class in make_task ( tasks.py ), run with
--task <name> :
class MyTask ( LLMTask ):
def __init__ ( self , n_examples = 2000 , seed = 0 , n_eval = 128 , rank = 0 , world = 1 , ** kw ):
self . rank , self . world = rank , world
rows = [...] # [{"question": ..., "value": ...}, ...]
self . _split ( rows , n_eval ) # eval holdout + per-rank sharding
self . reward_fns = [( gsm8k_reward , 1.0 ), ( format_reward , 0.1 )]
def build_prompt ( self , p ):
return p [ "question" ] + " \n Put ONLY the final number in <answer> </answer> tags."
A reward function is (prompt, completion, answer) -> float and can call whatever it
wants, from a regex to an external judge.
Things deliberately left out: Megatron-scale parallelism, MoE, multi-tenant scheduling,
dashboards. If you need those, you've outgrown this repo, and that's the point.
One RL training loop that runs CartPole on a laptop and disaggregated async RLVR on a GPU cluster. About 1,800 lines across 7 files. No Ray, no TRL, no DeepSpeed. It's small enough to read in an afternoon, and it's meant to be forked, not imported.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
