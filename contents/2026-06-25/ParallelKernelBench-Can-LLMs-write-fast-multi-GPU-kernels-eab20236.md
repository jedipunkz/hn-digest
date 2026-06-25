---
source: "https://github.com/togethercomputer/ParallelKernelBench"
hn_url: "https://news.ycombinator.com/item?id=48677270"
title: "ParallelKernelBench: Can LLMs write fast multi-GPU kernels?"
article_title: "GitHub - togethercomputer/ParallelKernelBench · GitHub"
author: "matt_d"
captured_at: "2026-06-25T18:42:52Z"
capture_tool: "hn-digest"
hn_id: 48677270
score: 1
comments: 0
posted_at: "2026-06-25T18:17:11Z"
tags:
  - hacker-news
  - translated
---

# ParallelKernelBench: Can LLMs write fast multi-GPU kernels?

- HN: [48677270](https://news.ycombinator.com/item?id=48677270)
- Source: [github.com](https://github.com/togethercomputer/ParallelKernelBench)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T18:17:11Z

## Translation

タイトル: ParallelKernelBench: LLM は高速マルチ GPU カーネルを作成できますか?
記事のタイトル: GitHub - togethercomputer/ParallelKernelBench · GitHub
説明: GitHub でアカウントを作成して、togettercomputer/ParallelKernelBench の開発に貢献します。

記事本文:
GitHub - togethercomputer/ParallelKernelBench · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
一緒にコンピューター
/
ParallelKernelBench
公共
通知
通知設定を変更するにはサインインする必要があります
アディ

ナビゲーション オプション
コード
一緒にコンピューター/ParallelKernelBench
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
39 コミット 39 コミット 画像 画像 kernelgen kernelgen リファレンス リファレンス run_together run_together スクリプト スクリプト utils utils .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml run_local.py run_local.py run_modal.py run_modal.py run_together.py run_together.py uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ParallelKernelBench: LLM は高速マルチ GPU カーネルを作成できますか?
ParallelKernelBench (PKB) は、LLM がマルチ GPU カーネルを最適化できるようにすることを目的としたベンチマークです。具体的には、既存の PyTorch + NCCL リファレンス コードをきめ細かい CUDA (または関連 DSL) に変換するモデルの機能を調査します。
この設計は KernelBench から大きく影響を受けています。
📄紙 ·
🤗 顔を抱きしめる ·
🌐 ブログ
PKB はモデルにマルチ GPU カーネルを最適化するように要求します。各問題には、reference/ の下に PyTorch + NCCL 参照があります。候補は solutions_<backend>/ (CUDA、Triton、ParallelKittens、または生成からの実行固有のツリー) に入ります。
正確さ: eval モードは、同じ入力に対して参照と候補を実行し、 --atol / --rtol 内でランクごとの出力 (rank_*.pt) を比較します。
パフォーマンス: オプションのタイミング レポートの速度向上と基準。私たちは ThunderKittens 2 に従っています — ベンチマークの厳密さ: 500 回のウォームアップ反復、100 回の時間指定反復 (ワーカー/パフォーマンス ユーティリティを参照)。
ルーフライン (概算):reference_rooflines_code/ は使用率の推定値を提供します。貢献を歓迎します。
PKB は、再現可能な Python 環境のために uv を使用します。
OS: NVIDIA GPU を搭載した Linux (マルチ GPU の実行には、torchrun / NCCL と一致する必要があります)。
ドライバー: CUDA 12.8 ホイールにとって十分な最新のもの (通常、H100 ノードはこれを満たします)。
ParallelKittens バックエンド (オプション): Thun のクローン

derKittens を開き、THUNDERKITTENS_ROOT をリポジトリ ルートに設定します (Modal/Togetter イメージではこれが自動的に行われます)。
# UV をインストールします (既にインストールされている場合はスキップします)
カール -LsSf https://astral.sh/uv/install.sh |しー
cd ParallelKernelBench
cd カーネルゲン
git clone https://github.com/SWE-agent/mini-swe-agent.git
CD..
UV同期
# 環境を確認する
uv run python -c " import torch; print(torch.__version__, torch.cuda.is_available()) "
これにより、リポジトリのルートに .venv/ が作成され、固定された uv.lock からすべての依存関係がインストールされます ( pyproject.toml を変更するときにそのファイルをコミットします)。
APIキー（生成/クラウド評価）
Google モデル: GEMINI_API_KEY または GOOGLE_API_KEY 。
Together モデル: TOGETHER_API_KEY ( kernelgen/generate_kernel.py の ALLOWED_MODELS を参照)。
人間モデル: ANTHROPIC_API_KEY 。
OpenAI モデル: OPENAI_API_KEY 。
[kernelgen/generate_kernel.py](kernelgen/generate_kernel.py) は、[kernelgen/prompts.toml](kernelgen/prompts.toml) からカーネル生成プロンプトをアセンブルし、オプションで LLM を呼び出します。 (1) プロンプトのみを印刷する ( --print-prompt )、または (2) 1 つの問題とバックエンドに対するソリューション ファイルを生成することができます。
# [print-prompt] 組み立てられたユーザー プロンプトを検査します。 API呼び出しもファイルの書き込みもありません
# --精度: fp32 | FP16 | bf16 (prompts.toml のエントリと一致する必要があります)
# --ハードウェア: h100_8 | b200_72 (オプション; 省略 → 出力ディレクトリ名に「なし」)
# --バックエンド: cuda |トリトン | Parallelkittens (prompts.toml の [backends] と一致する必要があります)
# --model:generate_kernel.py の ALLOWED_MODELS になければなりません (--print-prompt では無視されます)
uv 実行 python kernelgen/generate_kernel.py \
--精度 bf16 \
--ハードウェア h100_8 \
--問題 1 \
--model zai-org/GLM-5.1 \
--バックエンド cuda \
--プリントプロンプト
# [生成] LLM を呼び出して次のように書きます。 solutions_cuda_bf16_h100_8_together_zai-org_GLM-5.1/1_allreduce_cuda.py
UV実行Pythonカーン

lgen/generate_kernel.py \
--精度 bf16 \
--ハードウェア h100_8 \
--問題 1 \
--model zai-org/GLM-5.1 \
--バックエンドcuda
# 他のバックエンド (同じフラグ、異なる出力ディレクトリ接頭辞とファイル名接尾辞)
uv run python kernelgen/generate_kernel.py --precision bf16 --hardware h100_8 --problem 1 --model gemini-3-pro-preview --backend triton
uv run python kernelgen/generate_kernel.py --precision bf16 --hardware h100_8 --problem 1 --model gemini-3-pro-preview --backendParallelkittens
# オプション: カスタム プロンプト テンプレート
uv run python kernelgen/generate_kernel.py --paths-to-prompts-template /path/to/prompts.toml --precision bf16 --problem 1 --backend cuda --print-prompt
出力: --print-prompt を使用しない場合、各実行は solutions_<backend>_<precision>_<hardware|none>_<provider>_<model_slug>/ の下に {stem}_{backend}.py (たとえば 1_allreduce_cuda.py ) として書き込みます。評価時に --solutions-dir 経由でそのディレクトリを [run_local.py](run_local.py) に渡します。
単一のソリューションを生成します (ミニ SWE エージェント)
カーネルの生成に mini-swe-agent を使用するスクリプト ( [kernelgen/generate_kernel_agent.py](kernelgen/generate_kernel_agent.py) ) を提供します。 Google モデルで機能を検証したことに注意してください。
このスクリプトを使用するには、mini-swe-agent を個別にインストールする必要があります。
pip install -e kernelgen/mini-swe-agent
cd カーネルゲン
git clone https://github.com/SWE-agent/mini-swe-agent.git
コマンドの例:
python kernelgen/generate_kernel_agent.py \
--問題 1 \
--バックエンド cuda \
--model gemini-3-flash-preview \
--ステップ制限 3 \
--タイムアウト 600 \
--remote-dryrun-command \
' python run_local.py --nproc-per-node 4 --mode dryrun --problem {problem_arg} --solution {backend} --measure-perf ' \
--remote-eval-command \
' python run_local.py --nproc-per-node 4 --mode eval --problem {problem_arg} --solution {backend} --measure

-perf '
複数またはすべてのソリューションを生成する
--problem フラグに便利な機能を提供して、複数の問題を簡単に生成できるようにします。
# 参照の下にすべての問題を生成/
python kernelgen/generate_kernel.py \
--精度 bf16 \
--ハードウェア h100_8 \
-- 問題すべて \
--モデル deepseek-ai/DeepSeek-V4-Pro \
--バックエンドcuda
# 問題の特定のサブセットを生成する
python kernelgen/generate_kernel.py \
--精度 bf16 \
--ハードウェア h100_8 \
-- 問題 ' [72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88] ' \
--model zai-org/GLM-5.1 \
--バックエンドcuda
単一の問題を (ローカルで) 評価する
[run_local.py](run_local.py) は、ローカル マルチ GPU ハーネスです ( torchrun 経由)。 NVLink 経由で複数の GPU が接続されている環境があると仮定すると、(1) 1 つのバックエンドを分離して実行する ( dryrun )、または (2) 候補カーネルをリファレンスと比較する ( eval ) ことができます。
# [dryrun] --solution の単一の torchrun ジョブを起動します。参考文献と比較しません。これを使用して、カーネルがコンパイル、実行され、ランク出力が生成されることを確認します。
# PyTorch + NCCL ベースラインを参照 (デフォルトのソリューション ディレクトリ:reference/)
uv 実行 python run_local.py \
--nproc-ノードごと 4 \
--mode ドライラン \
--問題 1 \
--dtype bfloat16 \
--ソリューションリファレンス \ # (リファレンス / cuda / triton /Parallelkittens)
--download \ # オプション: Rank_*.pt を logs/problem_<stem>/<solution>/ に保存します。
--measure-perf # オプション: ウォームアップ + 時間制限された反復。 Rank_*_perf.json の概要を印刷して保存します
# CUDA ソリューション (カスタム ディレクトリ内)
uv 実行 python run_local.py \
--nproc-ノードごと 4 \
--mode ドライラン \
--問題 1 \
--dtype bfloat16 \
--ソリューション cuda \
--solutions-dir /path/to/custom/cuda/solution \
--ダウンロード \
--measure-perf \
--profile # オプション: PyTorch プロファイラー トレースを保存します (traces/ の下に書き込まれます)
# [eval] それぞれについて -

-ランダムな入力を試行し (デフォルトは 5)、参照を実行してから --solution を実行します。
# Rank_*.pt テンソルを比較し、最初の不一致で停止します。 --solution は参照であってはなりません。
python run_local.py \
--nproc-per-node 4 \ # 評価する GPU の数
--mode 評価 \
--問題 1 \
--dtype bfloat16 \
--ソリューション cuda \
--solutions-dir /path/to/custom/cuda/solution \
--trials 5 \ # RNG トライアルの数 (デフォルトは 5)
--atol 1e-2 \ # テンソル比較許容差
--rtol 1e-2 \
--download \ # トライアルごとのアーティファクトをログに保存します/
--measure-perf # 時間の参照と解決策
アーティファクト: --download を使用すると、出力は logs/problem_<stem>/{reference|cuda|triton|...}/ に配置されます (ランクごとの .pt ファイル、オプションの Rank_*_perf.json )。 [utils/compare_outputs.py](utils/compare_outputs.py) および [utils/compare_performance.py](utils/compare_performance.py) を使用してオフラインでそれらを検査します (「ダウンロードされた出力の検査」を参照)。
単一の問題を評価する (コンテナーをまとめて)
残念ながら、誰もが地元のコストコで DGX 8xH100 コンピューティング ノードを購入できるわけではありません。そこで、Together AI のクラウドコンテナサービス (Together Containers) を使用して評価する機能を提供します。専用 GPU ジョブは Together Jig フロー ( Together beta jig ) を使用します。
# プロフィール/ダッシュボードから Together API キーを見つける
エクスポート TOGETHER_API_KEY= " ... "
# dockerfile を作成し、ビルド、プッシュ、クラウドにデプロイします
一緒にベータジグをデプロイする
# 画像がクラウドサーバーにアップロードされたら、run_local.py と同じように評価できます
# 例:
python run_together.py --mode dryrun --問題 1
python run_together.py --mode dryrun --problem 1 --solution cuda --download
python run_together.py --mode dryrun --problem 1 --solution cuda --download --measure-perf --profile
python run_together.py --mode eval --問題 1
python run_together.py --mode eval --problem 1 --solution cuda --download

python run_together.py --mode eval --problem 1 --solution cuda --download --measure-perf --profile
# クラウドの状態を監視する
一緒にベータジグのステータス
# クラウドマシンのログを表示する
ベータ版ジグログをまとめて --フォローする
# リソースを節約する
一緒にベータジグを破壊する
完全に再構築せずにタグ付きイメージを再利用することもできます。
一緒にベータ ジグ ビルド --tag pkb_no_nvshmem
一緒にベータジグプッシュ --tag pkb_no_nvshmem
一緒にベータ ジグ デプロイ --image registry.together.xyz/ < レジストリ パス > :pkb_no_nvshmem
ダウンロードした出力を検査する
run_local.py --download (およびオプションで --measure-perf ) の後、アーティファクトは次のように編成されます。
ログ/問題_<ステム>/
Reference/ # Rank_0.pt、rank_1.pt、… およびオプションの Rank_*_perf.json
cuda/ # 候補となるバックエンドと同じレイアウト
比較_cuda.json # 比較_パフォーマンス.pyによって書かれました
--problem は run_local.py と同じ ID を受け入れます (例: 1 は 1_allreduce → logs/problem_1_allreduce/ に解決されます)。
# リファレンスと CUDA テンソルを比較します (デフォルト --solution cuda)
uv run python utils/compare_outputs.py --問題 1
# 別のバックエンドと比較します。必要に応じて公差を厳しくする
uv run python utils/compare_outputs.py --problem 1_allreduce --solution triton --atol 1e-2 --rtol 1e-2
# com を使用せずに 1 つのバックエンド フォルダーを検査します

[切り捨てられた]

## Original Extract

Contribute to togethercomputer/ParallelKernelBench development by creating an account on GitHub.

GitHub - togethercomputer/ParallelKernelBench · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
togethercomputer
/
ParallelKernelBench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
togethercomputer/ParallelKernelBench
main Branches Tags Go to file Code Open more actions menu Folders and files
39 Commits 39 Commits images images kernelgen kernelgen reference reference run_together run_together scripts scripts utils utils .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml run_local.py run_local.py run_modal.py run_modal.py run_together.py run_together.py uv.lock uv.lock View all files Repository files navigation
ParallelKernelBench: Can LLMs write fast multi-GPU kernels?
ParallelKernelBench (PKB) is a benchmark with the goal of enabling LLMs to optimize multi-GPU kernels. Specifically, we investigate model capabilities on turning existing PyTorch + NCCL reference code into fine-grained CUDA (or related DSLs).
The design is heavily inspired by KernelBench .
📄 Paper ·
🤗 Hugging Face ·
🌐 Blog
PKB asks models to optimize multi-GPU kernels: each problem has a PyTorch + NCCL reference under reference/ ; candidates go in solutions_<backend>/ (CUDA, Triton, ParallelKittens, or run-specific trees from generation).
Correctness: eval mode runs reference and candidate on the same inputs and compares per-rank outputs ( rank_*.pt ) within --atol / --rtol .
Performance: optional timing reports speedup vs reference. We follow ThunderKittens 2 — benchmark rigor : 500 warmup iterations, 100 timed iterations (see worker / perf utilities).
Roofline (approximate): reference_rooflines_code/ provides utilization estimates; contributions welcome.
PKB uses uv for reproducible Python environments.
OS: Linux with NVIDIA GPUs (multi-GPU runs need matching torchrun / NCCL).
Driver: Recent enough for CUDA 12.8 wheels (H100 nodes typically satisfy this).
ParallelKittens backend (optional): clone ThunderKittens and set THUNDERKITTENS_ROOT to the repo root (Modal/Together images do this automatically).
# Install uv (skip if already installed)
curl -LsSf https://astral.sh/uv/install.sh | sh
cd ParallelKernelBench
cd kernelgen
git clone https://github.com/SWE-agent/mini-swe-agent.git
cd ..
uv sync
# Verify the environment
uv run python -c " import torch; print(torch.__version__, torch.cuda.is_available()) "
This creates .venv/ at the repo root and installs all dependencies from the pinned uv.lock (commit that file when you change pyproject.toml ).
API keys (generation / cloud eval)
Google models: GEMINI_API_KEY or GOOGLE_API_KEY .
Together models: TOGETHER_API_KEY (see ALLOWED_MODELS in kernelgen/generate_kernel.py ).
Anthropic models: ANTHROPIC_API_KEY .
OpenAI models: OPENAI_API_KEY .
[kernelgen/generate_kernel.py](kernelgen/generate_kernel.py) assembles a kernel-generation prompt from [kernelgen/prompts.toml](kernelgen/prompts.toml) and optionally calls an LLM. You can (1) print the prompt only ( --print-prompt ) or (2) generate a solution file for one problem and backend.
# [print-prompt] Inspect the assembled user prompt; no API call, no file written
# --precision: fp32 | fp16 | bf16 (must match an entry in prompts.toml)
# --hardware: h100_8 | b200_72 (optional; omitted → "none" in the output directory name)
# --backend: cuda | triton | parallelkittens (must match [backends] in prompts.toml)
# --model: must be in ALLOWED_MODELS in generate_kernel.py (ignored with --print-prompt)
uv run python kernelgen/generate_kernel.py \
--precision bf16 \
--hardware h100_8 \
--problem 1 \
--model zai-org/GLM-5.1 \
--backend cuda \
--print-prompt
# [generate] Call the LLM and write e.g. solutions_cuda_bf16_h100_8_together_zai-org_GLM-5.1/1_allreduce_cuda.py
uv run python kernelgen/generate_kernel.py \
--precision bf16 \
--hardware h100_8 \
--problem 1 \
--model zai-org/GLM-5.1 \
--backend cuda
# Other backends (same flags; different output dir prefix and filename suffix)
uv run python kernelgen/generate_kernel.py --precision bf16 --hardware h100_8 --problem 1 --model gemini-3-pro-preview --backend triton
uv run python kernelgen/generate_kernel.py --precision bf16 --hardware h100_8 --problem 1 --model gemini-3-pro-preview --backend parallelkittens
# Optional: custom prompt template
uv run python kernelgen/generate_kernel.py --paths-to-prompts-template /path/to/prompts.toml --precision bf16 --problem 1 --backend cuda --print-prompt
Outputs: without --print-prompt , each run writes under solutions_<backend>_<precision>_<hardware|none>_<provider>_<model_slug>/ as {stem}_{backend}.py (for example 1_allreduce_cuda.py ). Pass that directory to [run_local.py](run_local.py) via --solutions-dir when evaluating.
Generate a single solution (mini-SWE-agent)
We provide a script ( [kernelgen/generate_kernel_agent.py](kernelgen/generate_kernel_agent.py) ) that uses mini-swe-agent in generating a kernel. Note that we verified functionality on Google models.
To use this script, you must install mini-swe-agent separately:
pip install -e kernelgen/mini-swe-agent
cd kernelgen
git clone https://github.com/SWE-agent/mini-swe-agent.git
An example command:
python kernelgen/generate_kernel_agent.py \
--problem 1 \
--backend cuda \
--model gemini-3-flash-preview \
--step-limit 3 \
--timeout 600 \
--remote-dryrun-command \
' python run_local.py --nproc-per-node 4 --mode dryrun --problem {problem_arg} --solution {backend} --measure-perf ' \
--remote-eval-command \
' python run_local.py --nproc-per-node 4 --mode eval --problem {problem_arg} --solution {backend} --measure-perf '
Generate multiple or all solutions
We provide convienient functionality to the --problem flag to make it simple to generate for multiple problems:
# generate every problem under reference/
python kernelgen/generate_kernel.py \
--precision bf16 \
--hardware h100_8 \
--problem all \
--model deepseek-ai/DeepSeek-V4-Pro \
--backend cuda
# generate a specific subset of problems
python kernelgen/generate_kernel.py \
--precision bf16 \
--hardware h100_8 \
--problem ' [72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88] ' \
--model zai-org/GLM-5.1 \
--backend cuda
Evaluate a single problem (locally)
[run_local.py](run_local.py) is the local multi-GPU harness (via torchrun ). Assuming you have an environment with multiple GPUs connected via NVLink, you can (1) run one backend in isolation ( dryrun ) or (2) compare a candidate kernel against the reference ( eval ).
# [dryrun] Launches a single torchrun job for --solution. Does not compare to reference. Use this to verify a kernel compiles, runs, and produces rank outputs.
# Reference PyTorch + NCCL baseline (default solutions dir: reference/)
uv run python run_local.py \
--nproc-per-node 4 \
--mode dryrun \
--problem 1 \
--dtype bfloat16 \
--solution reference \ # (reference / cuda / triton / parallelkittens)
--download \ # Optional: save rank_*.pt under logs/problem_<stem>/<solution>/
--measure-perf # Optional: warmup + timed iterations; prints and saves rank_*_perf.json summary
# CUDA solution (in a custom directory)
uv run python run_local.py \
--nproc-per-node 4 \
--mode dryrun \
--problem 1 \
--dtype bfloat16 \
--solution cuda \
--solutions-dir /path/to/custom/cuda/solution \
--download \
--measure-perf \
--profile # Optional: save PyTorch profiler traces (written under traces/)
# [eval] For each of --trials random inputs (default 5), runs reference then --solution.
# Compares rank_*.pt tensors, and stops on first mismatch. --solution must NOT be reference.
python run_local.py \
--nproc-per-node 4 \ # number of GPUs to evaluate on
--mode eval \
--problem 1 \
--dtype bfloat16 \
--solution cuda \
--solutions-dir /path/to/custom/cuda/solution \
--trials 5 \ # number of RNG trials (default 5)
--atol 1e-2 \ # tensor compare tolerances
--rtol 1e-2 \
--download \ # keep per-trial artifacts under logs/
--measure-perf # time reference and solution
Artifacts: with --download , outputs land in logs/problem_<stem>/{reference|cuda|triton|...}/ (per-rank .pt files, optional rank_*_perf.json ). Inspect them offline with [utils/compare_outputs.py](utils/compare_outputs.py) and [utils/compare_performance.py](utils/compare_performance.py) (see Inspect downloaded outputs ).
Evaluate a single problem (Together Containers)
Unfortunately, not everyone can buy a DGX 8xH100 compute node at their local Costco. So we provide the ability to evaluate using Together AI's cloud container service (Together Containters). Dedicated GPU jobs use the Together Jig flow ( together beta jig ).
# find your together API key from your profile/dashboard
export TOGETHER_API_KEY= " ... "
# creates a dockerfile, then build, push, deploy to the cloud
together beta jig deploy
# once the image is uploaded ot the cloud server, you can evaluate just like run_local.py
# e.g.
python run_together.py --mode dryrun --problem 1
python run_together.py --mode dryrun --problem 1 --solution cuda --download
python run_together.py --mode dryrun --problem 1 --solution cuda --download --measure-perf --profile
python run_together.py --mode eval --problem 1
python run_together.py --mode eval --problem 1 --solution cuda --download
python run_together.py --mode eval --problem 1 --solution cuda --download --measure-perf --profile
# monitor cloud status
together beta jig status
# view cloud machine logs
together beta jig logs --follow
# save resources
together beta jig destroy
You can also reuse a tagged image without a full rebuild:
together beta jig build --tag pkb_no_nvshmem
together beta jig push --tag pkb_no_nvshmem
together beta jig deploy --image registry.together.xyz/ < your-registry-path > :pkb_no_nvshmem
Inspect downloaded outputs
After run_local.py --download (and optionally --measure-perf ), artifacts are organized as:
logs/problem_<stem>/
reference/ # rank_0.pt, rank_1.pt, … and optional rank_*_perf.json
cuda/ # same layout for the candidate backend
comparison_cuda.json # written by compare_performance.py
--problem accepts the same ids as run_local.py (e.g. 1 resolves to 1_allreduce → logs/problem_1_allreduce/ ).
# Compare reference vs CUDA tensors (default --solution cuda)
uv run python utils/compare_outputs.py --problem 1
# Compare against another backend; tighten tolerances if needed
uv run python utils/compare_outputs.py --problem 1_allreduce --solution triton --atol 1e-2 --rtol 1e-2
# Inspect one backend folder without com

[truncated]
