---
source: "https://github.com/workofart/harness-training"
hn_url: "https://news.ycombinator.com/item?id=49293267"
title: "Show HN: Auto-train the harness, not the LLM. cross-model, cross-benchmark gains"
article_title: "GitHub - workofart/harness-training: Train a harness to improve its model-agnostic and task-env-agnostic capabilities with PyTorch-like APIs · GitHub"
author: "megadragon9"
captured_at: "2026-08-14T01:05:21Z"
capture_tool: "hn-digest"
hn_id: 49293267
score: 4
comments: 0
posted_at: "2026-08-14T00:05:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Auto-train the harness, not the LLM. cross-model, cross-benchmark gains

- HN: [49293267](https://news.ycombinator.com/item?id=49293267)
- Source: [github.com](https://github.com/workofart/harness-training)
- Score: 4
- Comments: 0
- Posted: 2026-08-14T00:05:34Z

## Translation

タイトル: HN を表示: LLM ではなくハーネスを自動トレーニングします。クロスモデル、クロスベンチマークの利益
記事のタイトル: GitHub - workofart/harness-training: PyTorch のような API を使用してハーネスをトレーニングし、モデルに依存しない機能とタスク環境に依存しない機能を向上させる · GitHub
説明: PyTorch のような API を使用してハーネスをトレーニングし、モデルに依存しない機能とタスク環境に依存しない機能を向上させる - workofart/harness-training

記事本文:
GitHub - workofart/harness-training: PyTorch のような API を使用してハーネスをトレーニングし、モデルに依存しない機能とタスク環境に依存しない機能を向上させる · GitHub
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
ワークファート
/
ハーネストレーニング
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
50 コミット 50 コミット config config docs docs 例 例 スクリプト スクリプト src src テスト テスト .env.example .env.example .gitignore .gitignor

e ライセンス ライセンス README.md README.md Program.md Program.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
PyTorch スタイルのハーネス トレーナー。 LLM モデルはフリーズしたままで、プロンプト、コンテキスト管理、ツール、修復ループなど、その周りのハーネスがトレーニングされています。
トレーナーの損失のために。エポック ( 30 ):
紛失。 reverse () # 評決を harness.grad に預けます
オプティマイザー。 step () # 勝者に HEAD を早送りするか、拒否します
ハーネスは 1 つの編集可能なファイル src/policy/core.py です。エポックごとに、src/trainer/estimator.py の Estimator が 1 つの diff を提案します。 diff は core.py に組み込まれ、タスクのパネル上の現在のベースラインに対して測定され、基準によって変更 (git commit) がプロモートされるかどうかが決定されます。 git log は、昇格候補の履歴です。昇格または拒否されたすべての候補者は refs/candidates/ の下に保持され、測定された各実行は refs/experiments/runs/ の下に保持されます。
ハーネスを 1 回トレーニングしてフリーズし、その後モデルの機能を全体的に引き上げます
詳細については、ブログ投稿を参照してください: https://www.henrypan.com/blog/2026-07-18-harness-training
「トレーニングされたハーネス」は凍結され、これらの評価のタスク解決 LLM のみが変更されます。
2 つのターミナルベンチ タスクでハーネスをトレーニングし、次に 2 つの保留タスクで評価します。
要件: Python 3.13、linux/amd64 コンテナーをサポートする Docker および Compose プラグイン、クロード CLI (ハーネスの変更を提案します。 CodexAgentBackend を使用してコーデックスを交換します)、ツール呼び出しを備えたローカルまたはリモートの OpenAI 互換推論サーバー。
Ollama または llama.cpp のいずれかを使用して、OpenAI 互換エンドポイントを通じてモデルを公開します。プラットフォーム固有のインストール手順に従ってください。
サーバーは一度に 1 つのリクエストをデコードする必要があります (llama.cpp の場合は --Parallel 1、

OLLAMA_NUM_PARALLEL=1 (Ollama の場合は同時リクエストの制限が他にある場合): 同時実行性の高いバッチ デコードは非決定的であるため、実行は候補ハーネスの変更に起因するものではありません。決定論 を参照してください。クライアント側では、引き続き max_rollout_concurrency を設定できます。後でバッチ デコード (より高い同時実行性) を有効にすることもできます。より広範囲のトレーニングには、SGLang Deterministic Inference に従うことをお勧めします。
config/llm/local.yaml でサーバーにクイックスタートを指定し、両方のクイックスタート構成が拡張する 1 つのファイルをコミットしてから、トレーナーはコミットされた構成のみを測定します。 model_name はサーバーがアドバタイズする ID で、tokenizer_name は HuggingFace Hub ID です。これはコンテキスト長管理のためのトークンをカウントするために必要です。
# UV をインストールします (まだインストールされていない場合)
カール -LsSf https://astral.sh/uv/install.sh |しー
# 依存関係をインストールし、仮想環境を作成する
UV同期
# 個の API キー。同梱されている LOCAL_LLM_API_KEY プレースホルダーは、次のようなサーバーに適しています。
# 認証を無視します。キーがチェックされている場合は交換してください。
cp .env.example .env
# 構成されたモデルサーバーが実行されている状態で、次のようにします。
uv 実行 Python 例/quickstart.py
注記
.env はキーが存在する場所であり、サーバーが認証を無視する場合でも変数は存在する必要があります。上記の cp をスキップすると、LOCAL_LLM_API_KEY が設定されていないためプリフライトが失敗します。サーバーがキーをチェックする場合は、付属のプレースホルダーを実際のプレースホルダーに置き換えます。そうしないと、「 401 Invalid API key 」という応答が返されます。
デフォルトの Terminal-Bench ネットワーク キャッシュは、 host.docker.internal を通じてホスト サービスに到達します。完全なホスト契約については、ネットワーク キャッシュ Runbook を参照してください。キャッシュを使用するかどうかを決定できます (クイックスタートを高速化するためにデフォルトでオンになっています)。トレーニングでは、キャッシュ サービスとそのボリュームは実行されたままになります。キャッシュされたデータを保持してそれらを停止するには: docker compose -f src/env/netcache/docker-compose

.caches.yml down — ボリュームも削除するには -v を追加します。
「実験」ブランチで実行します。変更プロモーションを利用してチェックアウトを早送りし、git ログに候補コミットが新しいベースラインとして表示されます。
初めて使用するとき、モデル サーバーはその重みをダウンロードし、クイックスタートはタスク イメージをダウンロードします。実行時間は主に、選択したモデルとハードウェアによって異なります。後続の実行では、ベースラインが再利用されます (ハーネスがベースライン git commit SHA からドリフトしていない限り)。
クイックスタートは、使い捨て VM またはコンテナー、別の OS アカウント、またはエージェント ワークロード専用のマシンで実行します。エスティメーターは、継承されたシェル環境でホスト プロセスとして実行されます。 「サンドボックスの境界」を参照してください。
# 基準によって「良さ」が決まります。ベースラインで解決されたタスクが後退する可能性はありません。
# 候補者はさらに解決する必要があります。正確な関係は二次的な指標に当てはまります。
基準 = 厳密パレート ()
# オプティマイザは、早送り HEAD または no-op の判定を適用するだけです。
オプティマイザー = GreedyMonotonic ()
トレーナー = トレーナー (
config_path = "config/train_harness.yaml" ,
エスティメーター = AgenticEstimator (
# Codex または Claude Code CLI を指定したり、AgenticEstimator を別の Estimator に切り替えることもできます
バックエンド = CodexAgentBackend (
trace_dir = パス ( "experiments/codex-traces" )、モデル = "gpt-5.6-sol"
）
）、
基準 = 基準 、
オプティマイザー = オプティマイザー 、
）
トレーナーの損失のために。エポック ( 30 ):
# HEAD を測定し、1 つの限定された変更を提案し、それを測定します。
# 同じタスク パネルで候補とベースラインを比較します。
紛失。 reverse () # 判定を記録します。
オプティマイザー。 step () # 昇格または拒否。
# 結果を次の提案のコンテキストとして保存します。
完全なチュートリアルは src/trainer/README.md にあります。
トレーニング ループがエポック全体で有用な信号を生成するため。各選挙の結果は、すべてが候補者に起因するものである必要があります。

それ以外は決定的です: シードされた決定的 LLM 推論エンジン、固定コンテナ ネットワーク、決定的環境、および凍結されたネットワーク キャッシュなど...このフレームワークはいくつかの方法でそれを保証します。詳細については、ブログ投稿を参照してください。
リスクを軽減するために、常に分離されたホスト マシンで実行することをお勧めします。不確実性は、ハーネスの変更を提案する「エージェント」から生じます。
パス
それは何ですか
ここから始めて…
ソース/ポリシー/
訓練中のハーネス
候補者が何に触れるかを見る
ソース/トレーナー/
トレーナー、推定者、基準、最適化者
トレーニング ループを作成またはカスタマイズする
src/ロールアウト/
測定されたエピソード、障害分類法、アーティファクト
測定値とそれがディスクに何を残すかを理解する
ソース/環境/
ターミナルベンチおよび SWE ベンチのタスク環境
ベンチマークを追加する
src/llm/
完了バックエンドと推定エージェント バックエンド
モデルプロバイダーを追加する
src/プラグイン/
3 つのキャッシュと決定論認証
再実行を高速化します。認定を理解する
構成/
1 つの YAML ファイル = 1 つの測定定義
コメント化された run_config.template.yaml をコピーします。
テスト/
テスト設計契約
適合するテストを書く
プログラム.md
フレームワークではありません — 動作ポリシー AgenticEstimator が各エポックの提案者と診断者に渡します
候補者が提案する内容の方向性を決める
トレーニング: uv run python scripts/train.py は、 config/train_harness.yaml を使用して完全なトレーニング ループを実行します。
評価: uv run python scripts/evaluate.py <config> <config2> ...
トレーニングと評価を含む小規模なエンドツーエンド:examples/quickstart.py
2 つの環境がすぐに出荷され、1 つのモデル プロバイダー (任意の OpenAI 互換エンドポイント) が提供されます。新しいベンチマークは 1 つの DockerTaskEnv サブクラスであり、新しいプロバイダーは 1 つの CompletionBackend サブクラスとそのレジストリ エントリです。具体的な参照として、 src/env/ の TerminalBenchEnv と SweEnv 、および src/llm/ の OpenAICompletionBackend を使用します。

タスク イメージは数十 GB のディスクを必要とする場合があります。
検索を 1 つのサーフェスに向けてバイアスします。
# プログラム.md
## 目的
+ この実行では、プロンプト表面の変更のみを提案します (システム プロンプト、初期プロンプト、および修復プロンプト)。
候補者が編集できる内容を厳密に制限します。パッチ サーフェスは、トレーニングされたモジュール自体のファイルに extra_patch_paths を加えたものです。フレームワークは、その外部の差分を拒否します。ハーネスの一部をフリーズするには、トレーニング済みモジュールの外にハーネスを移動します。その場合、候補者は連絡を取ることができなくなります。
# config/run_config.yaml
トレーニング対象:
module : src.policy.core # build_policy / build_env_action をエクスポートする必要があります
extra_patch_paths :
- testing/policy/test_core_impl.py # 候補者が書き込める唯一の他のファイル
失敗モードを 1 つ攻撃します。現在失敗しているタスクのみのパネルには、各エポックの信号が集中します。
# config/run_config.yaml
環境：
$include : task_panels/tool_call_failures.yaml # タスク パネルの例
ハーネスを新しいモデルに再取り付けします。スコアはモデル + ハーネスのプロパティであるため、モデルの交換は再ベースラインと再トレーニングを意味します。
# config/run_config.yaml
$extends :
- llm/gptoss20_openrouter.yaml # は llm/qwen35_local.yaml でした
より安価なソリューションを推進します。タイブレークはすでにオンになっており、出荷されたすべてのエントリ ポイントがベンチマークのデフォルトの二次メトリクスを通過します。独自のタプルを渡して、正確な解決セットの関係を破るものを変更するか、独自の基準を定義して主ルールを置き換えます。
# scripts/train.py で呼び出します
# タイブレークを 1 つのシグナルに絞り込みます。メトリクスは src/rollout/metrics.py にあります
criterion = StrictPareto ( Secondary_metrics = ( Steps UsedMetric (),))
# src/trainer/loss.py で定義された、異なる主候補昇格基準
criterion = NetTaskSolve () # コードベースにないプレースホルダー基準
あらゆるプロセスから提案を推進します。 Estimator の契約には、ハーネスの変更を事前に行う必要があるとは記載されていません。

エージェントの皆さん。独自の API を定義することも、審査員付きの LLM パネルを定義することもでき、各提案は同じ決定論的な実行を経ます。
# scripts/train.py — トレーナー(estimator=PatchQueue([...]))
クラス PatchQueue ( Estimator ):
"""A/B テストの手書きバリアント: あなたが差分を作成し、ループがそれらを測定します。"""
def __init__ ( self , diffs : list [ Path ]):
自分自身。 diffs = iter ( diffs )
間違いなく提案します(
self 、 * 、 repo_root : パス、トラッカー : RunStore 、ターゲット : TrainingTargetConfig
) -> なし :
サブプロセス 。 run ([ "git" , "apply" , next ( self . diffs )], cwd = repo_root , check = True )
def Diagnostics (self、result、*、repo_root、tracker、target) -> なし:
pass # あなたが診断者です
ライセンス
PyTorch のような API を使用してハーネスをトレーニングし、モデルに依存しない機能とタスク環境に依存しない機能を向上させます
www.henrypan.com/blog/2026-07-18-harness-training/ トピック
Readme MIT ライセンス アクティビティ スター
2 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Train a harness to improve its model-agnostic and task-env-agnostic capabilities with PyTorch-like APIs - workofart/harness-training

GitHub - workofart/harness-training: Train a harness to improve its model-agnostic and task-env-agnostic capabilities with PyTorch-like APIs · GitHub
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
workofart
/
harness-training
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
50 Commits 50 Commits config config docs docs examples examples scripts scripts src src tests tests .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md program.md program.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A PyTorch-style harness trainer . The LLM model stays frozen, the harness around it is being trained, including the prompts, context management, tools, and repair loop.
for loss in trainer . epochs ( 30 ):
loss . backward () # deposit the verdict on harness.grad
optimizer . step () # fast-forward HEAD to the winner, or reject
The harness is one editable file src/policy/core.py . Each epoch, the Estimator in src/trainer/estimator.py proposes one diff to it. The diff is incorporated into core.py that will be measured against the current baseline on a panel of tasks, and a criterion decides whether the change (git commit) is promoted. git log is the candidate promotion history. Every candidate, promoted or rejected, is kept under refs/candidates/ , and each measured run under refs/experiments/runs/ .
Train the harness once, freeze it, then lift model capability generally
For more details, see the blog post: https://www.henrypan.com/blog/2026-07-18-harness-training
The "trained harness" is frozen, only the task-solving LLM for these evaluations is changed.
Train the harness on two Terminal-Bench tasks, then evaluate it on two held-out tasks.
Requirements: Python 3.13, Docker with linux/amd64 container support and the Compose plugin, the claude CLI (it proposes the harness changes; swap in codex with CodexAgentBackend ), local or remote OpenAI-compatible inference server with tool calling.
Use either Ollama or llama.cpp or anything you like to expose the model through an OpenAI-compatible endpoint. Follow their platform-specific installation instructions.
The server must decode one request at a time ( --parallel 1 for llama.cpp, OLLAMA_NUM_PARALLEL=1 for Ollama, whatever caps concurrent requests elsewhere): high concurrency batched decode is non-deterministic, so the run cannot be attributable to the candidate harness change. See Determinism . On the client side, we can still set max_rollout_concurrency . You can enable batch decoding (higher concurrency) later, which I recommend following SGLang Deterministic Inference for larger scope training.
Point the quickstart at your server in config/llm/local.yaml , the one file both quickstart configs extend, then commit it — the trainer measures committed configs only. model_name is the id your server advertises, and tokenizer_name is the HuggingFace Hub id, which is required for counting tokens for context length management.
# Install uv (if not already installed)
curl -LsSf https://astral.sh/uv/install.sh | sh
# install dependencies and create virtual environment
uv sync
# API keys. The shipped LOCAL_LLM_API_KEY placeholder suits any server that
# ignores auth; replace it if yours checks keys.
cp .env.example .env
# With the configured model server running:
uv run python examples/quickstart.py
Note
.env is where keys live, and the variable must exist even when the server ignores auth — skipping the cp above fails preflight with LOCAL_LLM_API_KEY is not set . If your server checks keys, replace the shipped placeholder with the real one, or it answers 401 Invalid API key .
The default Terminal-Bench network cache reaches host services through host.docker.internal . See the network-cache runbook for the full host contract. You can decide whether to use the cache (default on to speed up quickstart). Training leaves the cache services and their volumes running. To stop them, keeping the cached data: docker compose -f src/env/netcache/docker-compose.caches.yml down — add -v to delete the volumes too.
Run it on a "experiment" branch: harness change promotions fast-forward your checkout, git log shows the candidate commit as your new baseline.
On first use, the model server downloads its weights and the quickstart downloads the task images. Runtime depends mainly on the selected model and hardware. Subsequent runs reuse the baseline (unless the harness drifted from baseline git commit SHA).
Run the quickstart in a disposable VM or container, a separate OS account, or a machine dedicated to agent workloads. The estimator runs as a host process with your inherited shell environment; see Sandbox Boundaries .
# The criterion decides "goodness": no task the baseline solved may regress, and
# the candidate must solve more. Exact ties fall to secondary metrics.
criterion = StrictPareto ()
# The optimizer just applies that verdict: fast-forward HEAD, or no-op.
optimizer = GreedyMonotonic ()
trainer = Trainer (
config_path = "config/train_harness.yaml" ,
estimator = AgenticEstimator (
# you can specify Codex or Claude Code CLI or even switch out the AgenticEstimator with another estimator
backend = CodexAgentBackend (
trace_dir = Path ( "experiments/codex-traces" ), model = "gpt-5.6-sol"
)
),
criterion = criterion ,
optimizer = optimizer ,
)
for loss in trainer . epochs ( 30 ):
# Measure HEAD, propose one bounded change, then measure it.
# Compare candidate with baseline on the same task panel.
loss . backward () # Record the verdict.
optimizer . step () # Promote or reject.
# Save the outcome as context for the next proposal.
The full walkthrough is in src/trainer/README.md .
In order for the training loop to produce useful signals across epochs. Each run's outcome must be attributable to the candidate only if everything else is deterministic: a seeded deterministic LLM inference engine, fixed container networks, deterministic environment, and a frozen network cache etc... This framework guarantees that in a couple of ways, more details in the blog post .
Always recommend to run on a isolated host machine to reduce risk. The uncertainty comes from the "Agent" proposing the harness change.
path
what it is
start here to…
src/policy/
the harness being trained
see what a candidate may touch
src/trainer/
Trainer, Estimator, Criterion, Optimizer
write or customize a training loop
src/rollout/
the measured episode, failure taxonomy, artifacts
understand a measurement and what it leaves on disk
src/env/
Terminal-Bench and SWE-bench task environments
add a benchmark
src/llm/
completion backends and estimator agent backends
add a model provider
src/plugins/
the three caches and determinism certification
make re-runs fast; understand certification
config/
one YAML file = one measurement definition
copy the commented run_config.template.yaml
tests/
the test design contract
write tests that fit
program.md
not framework — the operating policy AgenticEstimator hands its proposer and diagnoser each epoch
steer what candidates propose
Training: uv run python scripts/train.py runs the full training loop with config/train_harness.yaml .
Evaluation: uv run python scripts/evaluate.py <config> <config2> ...
Small end-to-end with training and evaluation: examples/quickstart.py
Two environments ship out of the box, and one model provider — any OpenAI-compatible endpoint. A new benchmark is one DockerTaskEnv subclass, a new provider is one CompletionBackend subclass plus its registry entry. Use TerminalBenchEnv and SweEnv in src/env/ , and OpenAICompletionBackend in src/llm/ , as the concrete references. Task images can take tens of GB of disk.
Bias the search toward one surface.
# program.md
## Objective
+ This run, propose only prompt-surface changes: system, initial, and repair prompts.
Hard-limit what candidates may edit. The patch surface is exactly the trained module's own file plus extra_patch_paths — the framework rejects any diff outside it. To freeze part of the harness, move it out of the trained module; it is then unreachable to candidates.
# config/run_config.yaml
training_target :
module : src.policy.core # must export build_policy / build_env_action
extra_patch_paths :
- tests/policy/test_core_impl.py # the only other file a candidate may write
Attack one failure mode. A panel of only the tasks that currently fail that way concentrates each epoch's signal on it:
# config/run_config.yaml
environment :
$include : task_panels/tool_call_failures.yaml # example task panel
Re-fit the harness to a new model. The score is a property of model + harness, so a model swap means re-baseline and retrain.
# config/run_config.yaml
$extends :
- llm/gptoss20_openrouter.yaml # was llm/qwen35_local.yaml
Promote cheaper solves. Tie-breaking is already on: every shipped entry point passes the benchmark's default secondary metrics. Pass your own tuple to change what breaks an exact solved-set tie, or define your own Criterion to replace the primary rule.
# call in scripts/train.py
# narrow the tie-break to one signal; metrics live in src/rollout/metrics.py
criterion = StrictPareto ( secondary_metrics = ( StepsUsedMetric (),))
# different primary candidate promotion criterion, defined in src/trainer/loss.py
criterion = NetTaskSolve () # placeholder criterion not in codebase
Drive proposals from any process. Nothing in the Estimator contract says harness changes have to come from agents. You can define your own API, or an LLM panel with a judge, and each proposal goes through the same deterministic run.
# scripts/train.py — Trainer(estimator=PatchQueue([...]))
class PatchQueue ( Estimator ):
"""A/B-test hand-written variants: you author the diffs, the loop measures them."""
def __init__ ( self , diffs : list [ Path ]):
self . diffs = iter ( diffs )
def propose (
self , * , repo_root : Path , tracker : RunStore , target : TrainingTargetConfig
) -> None :
subprocess . run ([ "git" , "apply" , next ( self . diffs )], cwd = repo_root , check = True )
def diagnose ( self , result , * , repo_root , tracker , target ) -> None :
pass # you are the diagnoser
License
Train a harness to improve its model-agnostic and task-env-agnostic capabilities with PyTorch-like APIs
www.henrypan.com/blog/2026-07-18-harness-training/ Topics
Readme MIT license Activity Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
