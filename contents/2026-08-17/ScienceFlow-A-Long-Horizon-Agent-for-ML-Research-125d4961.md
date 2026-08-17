---
source: "https://github.com/huawei-noah/noah-research/tree/master/ScienceFlow"
hn_url: "https://news.ycombinator.com/item?id=49327231"
title: "ScienceFlow – A Long-Horizon Agent for ML Research"
article_title: "noah-research/ScienceFlow at master · huawei-noah/noah-research · GitHub"
author: "d3ron"
captured_at: "2026-08-17T07:44:16Z"
capture_tool: "hn-digest"
hn_id: 49327231
score: 1
comments: 0
posted_at: "2026-08-17T06:45:21Z"
tags:
  - hacker-news
  - translated
---

# ScienceFlow – A Long-Horizon Agent for ML Research

- HN: [49327231](https://news.ycombinator.com/item?id=49327231)
- Source: [github.com](https://github.com/huawei-noah/noah-research/tree/master/ScienceFlow)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T06:45:21Z

## Translation

タイトル: ScienceFlow – ML 研究のための長期的なエージェント
記事のタイトル: noah-research/ScienceFlow at master · huawei-noah/noah-research · GitHub
説明: ノア研究。 GitHub でアカウントを作成して、huawei-noah/noah-research の開発に貢献してください。

記事本文:
noah-research/ScienceFlow マスター · huawei-noah/noah-research · GitHub
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
ファーウェイノア
/
ノアリサーチ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
その他のオプション ディレクトリアクション
歴史 歴史マスター ブレッドクラム
コピーパスのトップフォルダーとファイル
.. .scienceflow/ スキル/ データ処理 .scienceflow/ スキル/ データ処理 ディープクラフト ディープクラフト ドキュメント ドキュメント サイエンスフロー サイエンスフロー スクリプト scr

ipts タスク タスク テスト テスト Web サイト Web サイト .gitignore .gitignore ライセンス ライセンス README.md README.md env.example env.example pyproject.toml pyproject.toml 要件.txt 要件.txt uv.lock uv.lock すべてのファイルを表示 README.md
エンドツーエンドの自動調査エージェント フレームワーク
プロジェクトニュース
·
紙（arXiv）
·
中国語
ScienceFlow は、数時間または数日にわたる生産的で安定した目標に沿った研究のためのエンドツーエンドの自動研究エージェント フレームワークです。回復可能な実行可能ワークスペース、永続的な状態、適応型探索、および証拠を認識した実行制御の結合に関する研究を組織化し、エージェントが検証済みの進行状況を失うことなく継続、リダイレクト、または回復できるようにします。
ScienceFlow は、機械学習、科学的モデリング、および数学的最適化を通じて、効果的な長期研究を維持し、24 時間の予算内で 75 タスクの MLE ベンチ全体で 70.22 ± 1.18% Any-Medal に達し、報告されている最も強力なベースラインを 4.92 パーセント ポイント上回っています。
図1a。完全な MLE ベンチ Any-Medal リーダーボード。 ScienceFlow の 3 回の独立した実行にわたる平均 ± SEM。
回復可能な実行可能状態。永続的な各 LNR ワーカーは、分離された実行可能ワークスペースで研究を進めます。アーカイブされた状態は、そのワークスペースをコンパクトなメモリ、検証証拠、およびリソース レコードにバインドします。
ステージゲート。タスク固有の結果信号は GateService を呼び出します。構成されたエバリュエーターは正規化された証拠を生成し、ゲート ポリシーが許可を決定します。受け入れられた結果は、台帳ファクトと回復可能なワークスペース スナップショットを含む不変のステージを具体化します。
エストラ。研究境界では、再アンカーによる実行可能状態遷移は、開始点 (現在のワークスペースまたはアーカイブされたステージ) と意図 (続行またはリダイレクト) の 2 軸の決定を行います。アーカイブされた開始ポイントを選択すると、i が復元されます。

次の調査セグメントの前の実行可能ファイルの状態。
永続的な記憶。レコードの追加が受け入れられました ステージの進行状況。 Fold は、古い記録を要約しながら、最新の最も検証済みのアンカー関連の証拠を明確に保ちます。展開/復元はインデックス付きの証拠と状態を取得し、アセンブルは次のセグメントのアンカー固有のコンテキストを構築します。
証拠を意識した実行制御。研究従事者は科学的なルートを選択しますが、管理者はリソースの可用性、残りの予算、検証された進捗状況、および回復可能性を使用して物理的なジョブを許可、リース、監視、タイムボックス、および停止します。有効なワーカーの状態は、 merge/finals/final_* でファイナライズされます。
図 2. ScienceFlow システム アーキテクチャ。研究者は、回復可能な実行可能な状態を操作し、境界トリガーの ESTRA 遷移を通じて長期的な軌道を適応させます。一方、証拠を意識した実行制御は、物理リソースの割り当てとランタイム実行を調整します。
LNR はデフォルトでスキルなしです: lnr_skill_tool_enabled: false および lnr_skill_auto_read: false 。
.scienceflow/skills/data_processing/ は、専用のデータ準備エージェントと検証分割ワークフローに対してのみ保持されます。
auto はデフォルトのエバリュエーター バックエンドであり、登録されたタスクを task_package に解決します。 artifact_command は、引き続き一般的なコマンドベースの評価に使用できます。
並列実行では、CPU/GPU リソースをタスク レベルでバインドし、CPU 容量をワーカー間で分割します。 GPU リースは、複数のワーカーによる制御された共有をサポートします。
タスク契約で許可されている場合、結果シグナルは送信なしでステージを作成する場合があります。 Merge は、必要なアーティファクトを保持する候補からのみ最終結果を出力できます。
サイエンスフロー/
§──scienceflow/ # フレームワークランタイム
│ §── core/ # エージェントのランタイム、ツール、メモリ、および実行
│ §── ソルバー/ # LNR、ステージライフサイクル、ESTRA

、再開、マージ
│ §── Gates/ # Stage Gate および Evaluator プラグイン
│ §── 安全性/ # 証拠を意識したリソースと実行制御
│ §── ui/ # インターフェースの監視とトレース
│ §── config/ # デフォルトとマニフェストの例
│ §── utils/ # 共有ランタイムユーティリティ
│ └─ cli.py # コマンドラインエントリポイント
§── task/ # タスク パッケージとエバリュエーター
§── scripts/ # 保守された実行および監視マニフェスト
§── .scienceflow/skills/data_processing/ # データ準備スキル
└── doc/scienceflow/ # 詳細なアーキテクチャのドキュメント
インストール
要件: Python 3.11+ および uv 。
# まだお持ちでない場合は uv をインストールします
カール -LsSf https://astral.sh/uv/install.sh | sh # または: pip install uv
# リポジトリのクローンを作成してプロジェクトに入る
git clone https://github.com/huawei-noah/noah-research.git
cd ノアリサーチ/サイエンスフロー
# LLM 資格情報を構成する
cp env.example .env # 次に .env を編集します: プロバイダーの API_KEY と BASE_URL を設定します
# .venv を作成し、ロックされた環境をインストールする
UV同期
注:
uv sync は、リポジトリ内の deepcraft サブパッケージ、公式 mlebench Git リビジョン、完全なテスト/ML スタックを含む、ロックされた uv.lock 環境をインストールします。
PyTorch ホイールのデフォルトは cu128 インデックスです (CUDA 12.8 時代のドライバーの場合)。別の CUDA ビルドが必要な場合は、pyproject.toml の [[tool.uv.index]] を調整します。
SciModelingBench のサポートはオプションです: uv sync --extra Scientific-design 。
uv run ... 経由でコマンドを実行するか、.venv/bin/python を直接使用してコマンドを実行します。
インタラクティブなリサーチを開始する REPL:
uv run python -mscienceflow.cli repl
維持されている 2 ワーカーの Nomad2018 サンプルを実行します。
uv run python -mscienceflow.cliParallel -mscripts/lnr.yaml -j 1
既存の実行を監視します。
uv run python -mサイエンスフロー.cliモニター --manifest scripts/lnr.ya

ml --リフレッシュ 5
専用のデータ準備エージェントを使用してデータセットを準備します。
uv run python -mscienceflow.cliParallel -mscripts/prep.yaml -j 1
自己完結型の円パッキング数学最適化サンプルを実行します (データセットや追加のオプションは必要ありません)。
uv 実行 Python タスク/opt_solver/_tools/prepare_math_opt_solver_tasks.py
uv 実行 python -mscienceflow.cli 並列 \
-mscienceflow/config/examples/tasks_circle_packing_example.yaml -j 1
準備ステップでは、小さなタスク パッケージ (problem.json と有効なベースライン) を ./data/opt_solver/ に書き込みます。次に、エージェントは artifacts/best_solution.json を繰り返し改善し、システム側の評価者が各候補を権威的に検証し、半径の合計によってスコアを付けます。
設定
目的
lnr.num_workers
1 つのタスク内の永続的なリサーチ ワーカーの数。
task.cpu_list / task.gpu_list
タスクレベルの CPU および GPU リソース境界。 LNR はさらに CPU をワーカー間で分割します。
lnr.omp_threads_cap
各ワーカー スライスの CPU スレッド キャップ。
lnr.wall_clock_budget_sec
LNR プロセスの実時間の合計バジェット。
lnr.estra_enabled / estra_trigger_stage_count
ESTRA を有効にし、境界レビューとコンテキストの折りたたみのトリガーを設定します。
lnr.resource_runtime_enabled
証拠を認識したリソースと実行制御を有効にします。
再開_予算_ポリシー
再開された実行に対する予算の計算。 fresh は、累積時間にこのラウンドの time_limit を追加します。
評価者.バックエンド
auto (デフォルト)、 task_package 、または artifact_command を選択します。
evaluator.stage_source_mode
Shadow 、 Adjudicate 、または Primary Stage ソース モードを選択します。
評価者.コマンド.python_executable
タスクを隔離された Python 環境に向けます。
profile_overrides.<プロファイル>
タスク タイプごとにプロンプト、エバリュエーター、およびリソースの動作をオーバーライドします。
タスク/**/task.yaml
タスクレベルのアーティファクト、メトリック、プロバイダー/プロファイル、評価を宣言します。

tor、およびゲート ポリシー。
metric.authoritative: true
メトリクスを信頼性の高い選択に適格な信頼できる証拠としてマークします。
ドキュメント
アーキテクチャの概要 · 回復可能な状態と LNR · 証拠を意識した実行制御 · オプトソルバー タスクの追加 · 科学的モデリングの例
この論文では、次の 3 つのクラスの実行可能な研究タスクにわたって同じ ScienceFlow ワークフローを評価しています。
機械学習エンジニアリング: パイプライン構築インターフェイスを介した 75 の MLE ベンチ タスクすべて。
科学的モデリングと設計: 候補最適化インターフェイスを介した Hugging Face に関する 12 の SciModelingBench タスク。
数学的および工学的最適化: 円パッキング、比率最小化、不確実性不等式、および候補最適化インターフェイスを介した簡単、中程度、および難しい SpOC4 KTTSP トラック。
すべてのタスク ファミリは Stage Gate と Evaluator コントラクトを共有します。各 task.yaml は、プロバイダー/プロファイル、アーティファクト スキーマ、メトリックの方向、エバリュエーター バックエンド、権限ステータス、およびゲート ポリシーを汎用ソルバーの外部に保持します。
MLE ベンチ タスクでは、データ ルート、タスク exp_id 、および submit.csv コントラクトを調整する必要があります。
異なるタスク プロファイル間で古いワークスペースを再開しないでください。再開すると、プロンプト、データセット、またはアーティファクト ディメンションが正しく継承されない可能性があります。
stop_by_user は、失敗ではなく、手動停止から再開可能な終了状態をマークします。後の履歴書は、state.json とワークスペースのステージに蓄積された予算から継続する必要があります。
uv 実行 pytest -q
このプロジェクトでは、Python 3.11 以降、Pydantic、OmegaConf、Click、および Rich を使用します。公式の mlebench 依存関係は、 uv.lock の Git リビジョンに固定されています。オプションの科学設計エクストラでは、SciModelingBench、Datasets、および PyArrow のサポートが提供されます。コマンドベースの最適化タスクでは、エバリュエーターを通じて挿入された別の Python 環境を使用する場合があります。

構成。タスク固有のスキーマ/スコアリング ロジックは、tasks/<category>/... タスク パッケージ内に保持されます。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Noah Research. Contribute to huawei-noah/noah-research development by creating an account on GitHub.

noah-research/ScienceFlow at master · huawei-noah/noah-research · GitHub
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
huawei-noah
/
noah-research
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History master Breadcrumbs
Copy path Top Folders and files
.. .scienceflow/ skills/ data_processing .scienceflow/ skills/ data_processing deepcraft deepcraft doc doc scienceflow scienceflow scripts scripts tasks tasks tests tests website website .gitignore .gitignore LICENSE LICENSE README.md README.md env.example env.example pyproject.toml pyproject.toml requirements.txt requirements.txt uv.lock uv.lock View all files README.md
An End-to-End Autoresearch Agent Framework
Project News
·
Paper (arXiv)
·
Chinese
ScienceFlow is an end-to-end autoresearch agent framework for productive, stable, and goal-aligned research over hours or days. It organizes research around recoverable executable workspaces, coupling persistent state, adaptive exploration, and evidence-aware execution control so agents can continue, redirect, or recover without losing validated progress.
Across machine learning, scientific modeling, and mathematical optimization, ScienceFlow sustains effective long-horizon research and reaches 70.22 ± 1.18% Any-Medal on the full 75-task MLE-bench within a 24-hour budget, exceeding the strongest reported baseline by 4.92 percentage points .
Figure 1a. Full MLE-bench Any-Medal leaderboard. Mean ± SEM over three independent runs for ScienceFlow.
Recoverable executable state. Each persistent LNR worker advances research in an isolated executable workspace. An archived state binds that workspace to compact memory, validation evidence, and resource records.
Stage Gate. A task-specific result signal invokes GateService : the configured Evaluator produces normalized evidence, and the Gate policy decides admission. An accepted result materializes an immutable Stage with ledger facts and a recoverable workspace snapshot.
ESTRA. At a research boundary, Executable-State Transition through Re-Anchoring makes a two-axis decision: a start point (the current workspace or an archived Stage) and an intent ( continue or redirect ). Selecting an archived start point restores its executable state before the next research segment.
Persistent memory. Add records accepted Stage progress. Fold keeps recent, best-validated, and anchor-relevant evidence explicit while summarizing older records; Unfold/restore retrieves indexed evidence and state, and Assemble constructs the anchor-specific context for the next segment.
Evidence-aware execution control. Research workers choose scientific routes, while the controller admits, leases, monitors, timeboxes, and stops physical jobs using resource availability, remaining budget, validated progress, and recoverability. Valid worker states are finalized under merge/finals/final_* .
Figure 2. ScienceFlow system architecture. Research workers operate over recoverable executable states and adapt long-horizon trajectories through boundary-triggered ESTRA transitions, while evidence-aware execution control coordinates physical resource allocation and runtime execution.
LNR is no-skill by default : lnr_skill_tool_enabled: false and lnr_skill_auto_read: false .
.scienceflow/skills/data_processing/ is retained only for the dedicated data-prep agent and validation-split workflow.
auto is the default evaluator backend and resolves registered tasks to task_package ; artifact_command remains available for generic command-based evaluation.
Parallel runs bind CPU/GPU resources at task level, then split CPU capacity across workers. GPU leases support controlled sharing by multiple workers.
Result signals may create Stages without a submission when the task contract permits it. Merge can only emit finals from candidates that carry the required artifact.
ScienceFlow/
├── scienceflow/ # Framework runtime
│ ├── core/ # Agent runtime, tools, memory, and execution
│ ├── solver/ # LNR, Stage lifecycle, ESTRA, resume, and merge
│ ├── gates/ # Stage Gate and Evaluator plugins
│ ├── safety/ # Evidence-aware resource and execution control
│ ├── ui/ # Monitor and trace interfaces
│ ├── config/ # Defaults and example manifests
│ ├── utils/ # Shared runtime utilities
│ └── cli.py # Command-line entry point
├── tasks/ # Task packages and evaluators
├── scripts/ # Maintained run and monitor manifests
├── .scienceflow/skills/data_processing/ # Data-preparation skills
└── doc/scienceflow/ # Detailed architecture documentation
Installation
Requirements: Python 3.11+ and uv .
# Install uv if you do not have it yet
curl -LsSf https://astral.sh/uv/install.sh | sh # or: pip install uv
# Clone the repository and enter the project
git clone https://github.com/huawei-noah/noah-research.git
cd noah-research/ScienceFlow
# Configure LLM credentials
cp env.example .env # then edit .env: set API_KEY and BASE_URL for your provider
# Create .venv and install the locked environment
uv sync
Notes:
uv sync installs the locked uv.lock environment, including the in-repo deepcraft subpackages, the official mlebench Git revision, and the full test/ML stack.
PyTorch wheels default to the cu128 index (for CUDA 12.8-era drivers). Adjust [[tool.uv.index]] in pyproject.toml if you need a different CUDA build.
SciModelingBench support is an optional extra: uv sync --extra scientific-design .
Run commands either via uv run ... or by using .venv/bin/python directly.
Start an interactive research REPL:
uv run python -m scienceflow.cli repl
Run the maintained two-worker Nomad2018 example:
uv run python -m scienceflow.cli parallel -m scripts/lnr.yaml -j 1
Monitor an existing run:
uv run python -m scienceflow.cli monitor --manifest scripts/lnr.yaml --refresh 5
Prepare a dataset with the dedicated data-prep agent:
uv run python -m scienceflow.cli parallel -m scripts/prep.yaml -j 1
Run the self-contained Circle Packing math-optimization example (no dataset or optional extra required):
uv run python tasks/opt_solver/_tools/prepare_math_opt_solver_tasks.py
uv run python -m scienceflow.cli parallel \
-m scienceflow/config/examples/tasks_circle_packing_example.yaml -j 1
The prepare step writes a tiny task package ( problem.json plus a valid baseline) under ./data/opt_solver/ . The agent then iteratively improves artifacts/best_solution.json , and the system-side evaluator authoritatively validates each candidate and scores it by the sum of radii.
Setting
Purpose
lnr.num_workers
Number of persistent research workers inside one task.
task.cpu_list / task.gpu_list
Task-level CPU and GPU resource boundaries; LNR further splits CPU across workers.
lnr.omp_threads_cap
CPU thread cap for each worker slice.
lnr.wall_clock_budget_sec
Total wall-clock budget for the LNR process.
lnr.estra_enabled / estra_trigger_stage_count
Enables ESTRA and sets the trigger for boundary review and context folding.
lnr.resource_runtime_enabled
Enables evidence-aware resource and execution control.
resume_budget_policy
Budget accounting for resumed runs; fresh adds this round's time_limit on top of accumulated time.
evaluator.backend
Selects auto (default), task_package , or artifact_command .
evaluator.stage_source_mode
Selects the shadow , adjudicate , or primary Stage source mode.
evaluator.command.python_executable
Points a task at an isolated Python environment.
profile_overrides.<profile>
Overrides prompts, Evaluator, and resource behavior per task type.
tasks/**/task.yaml
Declares the task-level artifact, metric, provider/profile, Evaluator, and Gate policy.
metric.authoritative: true
Marks a metric as authoritative evidence eligible for high-trust selection.
Documentation
Architecture overview · Recoverable states and LNR · Evidence-aware execution control · Adding opt-solver tasks · Scientific modeling example
The paper evaluates the same ScienceFlow workflow across three classes of executable research tasks:
Machine learning engineering: all 75 MLE-bench tasks through the pipeline-construction interface.
Scientific modeling and design: 12 SciModelingBench tasks on Hugging Face through the candidate-optimization interface.
Mathematical and engineering optimization: Circle Packing , Ratio Minimization , Uncertainty Inequality , and the easy, medium, and hard SpOC4 KTTSP tracks through the candidate-optimization interface.
All task families share the Stage Gate and Evaluator contract. Each task.yaml keeps provider/profile, artifact schema, metric direction, evaluator backend, authoritative status, and Gate policy outside the generic solver.
MLE-bench tasks require the data root, task exp_id , and submission.csv contract to be aligned.
Do not resume old workspaces across different task profiles, or prompts, datasets, or artifact dimensions may be inherited incorrectly.
stopped_by_user marks a resumable terminal state from a manual stop, not a failure; a later resume should continue from the accumulated budget in state.json and the workspace stages.
uv run pytest -q
The project uses Python 3.11+, Pydantic, OmegaConf, Click, and Rich. The official mlebench dependency is pinned to a Git revision in uv.lock ; the optional scientific-design extra provides SciModelingBench, Datasets, and PyArrow support. Command-based optimization tasks may use a separate Python environment injected through the evaluator configuration, with task-specific schema/scoring logic kept inside the tasks/<category>/... task package.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
