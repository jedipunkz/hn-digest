---
source: "https://github.com/hexo-ai/sia"
hn_url: "https://news.ycombinator.com/item?id=48344807"
title: "AI Agent that at inference time updates it's harness and model weights"
article_title: "GitHub - hexo-ai/sia: SIA is a Self Improving AI framework to autonomously improve the performance of any AI system (Model / Agent) on a benchmark task. · GitHub"
author: "martianvoid"
captured_at: "2026-06-01T00:32:24Z"
capture_tool: "hn-digest"
hn_id: 48344807
score: 5
comments: 0
posted_at: "2026-05-31T11:13:43Z"
tags:
  - hacker-news
  - translated
---

# AI Agent that at inference time updates it's harness and model weights

- HN: [48344807](https://news.ycombinator.com/item?id=48344807)
- Source: [github.com](https://github.com/hexo-ai/sia)
- Score: 5
- Comments: 0
- Posted: 2026-05-31T11:13:43Z

## Translation

タイトル: 推論時にハーネスとモデルの重みを更新する AI エージェント
記事タイトル: GitHub - hexo-ai/sia: SIA は、ベンチマーク タスクで任意の AI システム (モデル / エージェント) のパフォーマンスを自律的に向上させる自己改善型 AI フレームワークです。 · GitHub
説明: SIA は、ベンチマーク タスクで任意の AI システム (モデル / エージェント) のパフォーマンスを自律的に向上させる自己改善型 AI フレームワークです。 - ヘクソアイ/シア

記事本文:
GitHub - hexo-ai/sia: SIA は、ベンチマーク タスクで任意の AI システム (モデル / エージェント) のパフォーマンスを自律的に向上させる自己改善型 AI フレームワークです。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ヘクソアイ
/
シア
公共
通知
あなたはしなければなりません

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .github/ workflows .github/ workflows docs docs sia sia testing testing .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md EVALUATION_GUIDE.md EVALUATION_GUIDE.md ライセンス ライセンス README.md README.md 環境.yml 環境.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
SIA の正式実装: ハーネスと重みの更新による自己改善 AI (Hebbar et al.、2026) — 言語モデル エージェントがタスク固有のエージェントのハーネスと重みの両方を更新する自己改善ループ。この論文では、ベースラインと比較して、LawBench で 56.6% の向上、GPU カーネルでのランタイムの 91.9% の削減、および単一セル RNA のノイズ除去で 502% の改善が報告されています。
SIA は、ベンチマーク タスクで任意の AI システム (モデル / エージェント) のパフォーマンスを自律的に向上させる自己改善型 AI フレームワークです。
試してみたいだけですか？ 「SIA をローカルで実行する」に進みます。
メタエージェント、ターゲットエージェント、およびフィードバックエージェント間の一連の世代にわたるフローを制御します。
SIA は、タスクのパフォーマンスを継続的に向上させるために連携する 3 つの主要なタイプの AI エージェントを調整することによって動作します。
Meta-Agent : タスクの説明を読み取り、タスクに合わせた初期ターゲット エージェントを生成します。
ターゲット/タスク固有のエージェント : タスクの完了を試み、そのアクションと結果を記録します。
フィードバック/改善エージェント : ターゲット エージェントのパフォーマンス ログを確認し、改善点を特定し、それに応じてターゲット エージェントを更新します。
この反復プロセスにより、システムは科学的タスクを解決する能力を自律的に洗練し、強化することができます。
OpenAI MLE-Bench Hard: エージェントが完全に記述、実行、反復する必要がある実際の Kaggle ML コンテストの挑戦

ML パイプライン。 SIA は、テストされたすべての世代で第 1 位にランクされています。
LawBench: 191 の罪状カテゴリーにわたる中国の裁判の説明から刑事告訴を予測します。 SIA-W+H は 70.1% のトップ 1 精度に達し、以前の SOTA の 45% を上回りました。
AlphaFold-3 TriMul Triton カーネル: Triangle Multiplicative Update を Triton カーネルとして実装および最適化し、H100 レイテンシ ターゲットを達成しながら正確さを維持します。 SIA-W+H はベースラインと比較して 14 倍の高速化を達成します。
scRNA-seq ノイズ除去: 単一細胞 RNA シーケンス データ内の欠落している遺伝子発現値を代入します。 SIA-W+H のスコアは 0.289 MSE 基準で、以前の SOTA の 0.220 を上回っています。
組み込みタスクを使用して SIA をローカルで実行する
SIA には、 gpqa 、 lawbench 、 longcot-chess 、 spaceship-titanic の 4 つの組み込みタスクが付属しています。
実行する LLM に一致するエージェント バックエンドを選択します。
クロード バックエンド (クロード エージェント SDK、クロード モデルのみ):
python3 -m venv .venv && ソース .venv/bin/activate
pip インストール ' sia-agent[claude] '
エクスポート ANTHROPIC_API_KEY= " ... "
OpenHands バックエンド (マルチプロバイダー - Gemini、OpenAI、Anthropic など):
python3 -m venv .venv && ソース .venv/bin/activate
pip インストール ' sia-agent[openhands] '
# 使用するプロバイダーのキーをエクスポートします。
import ANTHROPIC_API_KEY= " ... " # anthropic/* モデルの場合
import GEMINI_API_KEY= " ... " # gemini/* モデル (または GOOGLE_API_KEY) の場合
import OPENAI_API_KEY= " ... " # openai/* モデルの場合
プロバイダー/モデルの完全なリファレンス: docs/configuration.md 。
sia --task gpqa --max_gen 5 --run_id 1
--task を 4 つのバンドルされたタスクのいずれかと交換します。
アーティファクトは、runs/run_{run_id}/gen_{n}/ に配置されます。
target_agent.py — その世代のエージェント
Agent_execution.json — 実行ログ
Improvement.md — 差分の根拠 (第 2 世代以降)
旗
デフォルト
説明
--タスク
—
バンドルされたタスク名 ( --task_dir と相互に排他的)
--task_dir
—
パ

th を外部タスクディレクトリへ
--max_gen
3
自己啓発世代数
--run_id
1
一意の実行識別子
--バックエンド
クロード
claude (Claude Agent SDK) または openhands (マルチプロバイダー)
--meta_model
俳句
メタ/フィードバック モデル (例: haiku 、 Sonnet 、 opus 、または gemini/... 、 openai/... with openhands)
--タスクモデル
クロード-俳句-4-5-20251001
ターゲットエージェントモデル
バックエンド、モデル、API キーの完全なリファレンス: docs/configuration.md 。何か問題がありましたか？ docs/troubleshooting.md 。
以下のレイアウトのタスク ディレクトリを準備し、 --task_dir を指定します。
私のタスク/
§── データ/
│ §── 公開/
│ │ §── task.md # タスクの説明 — SIA がこれを読み取ります
│ │ └─ ... # エージェントが参照できる入力
│ └── private/ # 保持された評価データ;エージェントに決して暴露されない
━──参考/
§──reference_target_agent.py # テンプレート; sia/tasks/_shared/ からコピー
└── SAMPLE_TASK_DESCRIPTIONS.md # オプション: メタエージェントのタスクの例
sia --task_dir ./my-task --max_gen 5 --run_id 1
または、MLE ベンチ コンテストを開催してください。 SIA は、MLE-Bench コンペティションから直接タスク ディレクトリをブートストラップできます。Kaggle API 経由でデータセットを取得し、パブリック/プライベート分割を設定し、参照エージェント テンプレートをドロップします。
python -m sia.prepare_mlebench_dataset -c " スペースシップ-タイタニック "
sia --task_dir ./tasks/spaceship-titanic --max_gen 5 --run_id 1
両方のパスの完全なステップバイステップ: docs/walkthrough.md 。
docs/architecture.md — ディレクトリのレイアウト、生成フロー、プロンプトのカスタマイズ
docs/walkthrough.md — カスタムタスクの詳細なウォークスルー
docs/configuration.md — バックエンド、モデル、API キー、CLI リファレンス
docs/troubleshooting.md — 一般的なエラーと修正
研究で SIA を使用する場合は、以下を引用してください。
@article { hebbar2026sia 、
title = { SIA: Har による自己改善 AI

ネス \& 体重の更新 } 、
著者 = { ヘバール、プラナイとマナワット、ヨゲンドラとフェルブーメン、サミュエルとイワノワ、アレシアとパラニマライ、セルヴァムとバティア、クナルとバスカラン、ヴィグネーシュ } ,
ジャーナル = { arXiv プレプリント arXiv:2605.27276 } 、
年 = { 2026 } 、
URL = { https://arxiv.org/abs/2605.27276 }
}
について
SIA は、ベンチマーク タスクで任意の AI システム (モデル / エージェント) のパフォーマンスを自律的に向上させる自己改善型 AI フレームワークです。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
82
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

SIA is a Self Improving AI framework to autonomously improve the performance of any AI system (Model / Agent) on a benchmark task. - hexo-ai/sia

GitHub - hexo-ai/sia: SIA is a Self Improving AI framework to autonomously improve the performance of any AI system (Model / Agent) on a benchmark task. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
hexo-ai
/
sia
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github/ workflows .github/ workflows docs docs sia sia tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md EVALUATION_GUIDE.md EVALUATION_GUIDE.md LICENSE LICENSE README.md README.md environment.yml environment.yml pyproject.toml pyproject.toml View all files Repository files navigation
Official implementation of SIA: Self Improving AI with Harness & Weight Updates (Hebbar et al., 2026) — a self-improving loop where a language-model agent updates both the harness and the weights of a task-specific agent. The paper reports a 56.6% gain on LawBench, 91.9% runtime reduction on GPU kernels, and 502% improvement on single-cell RNA denoising over baseline.
SIA is a Self Improving AI framework to autonomously improve the performance of any AI system (Model / Agent) on a benchmark task.
Just want to try it? Skip to Run SIA locally .
Control flow between Meta, Target, and Feedback agents over successive generations.
SIA operates by coordinating three main types of AI agents that work together to continuously improve task performance:
Meta-Agent : Reads the task description and generates an initial Target Agent tailored to the task.
Target / Task Specific Agent : Attempts to complete the task and records its actions and results.
Feedback/Improvement Agent : Reviews the Target Agent's performance logs, identifies improvements, and updates the Target Agent accordingly.
This iterative process allows the system to autonomously refine and enhance its ability to solve scientific tasks.
OpenAI MLE-Bench Hard: a gauntlet of real Kaggle ML competitions where agents must write, run, and iterate full ML pipelines. SIA ranks #1 across all generations tested.
LawBench: predict the criminal charge from Chinese court case descriptions across 191 charge categories. SIA-W+H reaches 70.1% Top-1 accuracy, beating the prior SOTA of 45%.
AlphaFold-3 TriMul Triton Kernel: implement and optimize the Triangle Multiplicative Update as a Triton kernel, preserving correctness while hitting H100 latency targets. SIA-W+H achieves 14x speedup over baseline.
scRNA-seq Denoising: impute missing gene expression values in single-cell RNA sequencing data. SIA-W+H scores 0.289 MSE norm , surpassing the prior SOTA of 0.220.
Run SIA locally with built-in tasks
SIA ships with four built-in tasks: gpqa , lawbench , longcot-chess , spaceship-titanic .
Pick the Agent backend that matches the LLMs you want to run.
Claude backend (Claude Agent SDK, Claude models only):
python3 -m venv .venv && source .venv/bin/activate
pip install ' sia-agent[claude] '
export ANTHROPIC_API_KEY= " ... "
OpenHands backend (multi-provider — Gemini, OpenAI, Anthropic, etc.):
python3 -m venv .venv && source .venv/bin/activate
pip install ' sia-agent[openhands] '
# Export the key(s) for the provider(s) you'll use:
export ANTHROPIC_API_KEY= " ... " # for anthropic/* models
export GEMINI_API_KEY= " ... " # for gemini/* models (or GOOGLE_API_KEY)
export OPENAI_API_KEY= " ... " # for openai/* models
Full provider/model reference: docs/configuration.md .
sia --task gpqa --max_gen 5 --run_id 1
Swap --task for any of the four bundled tasks.
Artifacts land in runs/run_{run_id}/gen_{n}/ :
target_agent.py — the agent for that generation
agent_execution.json — execution logs
improvement.md — diff rationale (gen 2+)
Flag
Default
Description
--task
—
Bundled task name (mutually exclusive with --task_dir )
--task_dir
—
Path to an external task directory
--max_gen
3
Number of self-improvement generations
--run_id
1
Unique run identifier
--backend
claude
claude (Claude Agent SDK) or openhands (multi-provider)
--meta_model
haiku
Meta/feedback model (e.g. haiku , sonnet , opus , or gemini/... , openai/... with openhands)
--task_model
claude-haiku-4-5-20251001
Target agent model
Full backend, model, and API-key reference: docs/configuration.md . Hit a snag? docs/troubleshooting.md .
Prepare a task directory with the layout below and point --task_dir at it:
my-task/
├── data/
│ ├── public/
│ │ ├── task.md # Task description — SIA reads this
│ │ └── ... # Inputs the agent is allowed to see
│ └── private/ # Held-out eval data; never exposed to the agent
└── reference/
├── reference_target_agent.py # Template; copy from sia/tasks/_shared/
└── SAMPLE_TASK_DESCRIPTIONS.md # Optional: example tasks for the meta-agent
sia --task_dir ./my-task --max_gen 5 --run_id 1
Or bring an MLE-Bench competition. SIA can bootstrap a task directory directly from any MLE-Bench competition — it pulls the dataset via the Kaggle API, sets up the public/private split, and drops in the reference agent template:
python -m sia.prepare_mlebench_dataset -c " spaceship-titanic "
sia --task_dir ./tasks/spaceship-titanic --max_gen 5 --run_id 1
Full step-by-step for both paths: docs/walkthrough.md .
docs/architecture.md — directory layout, generation flow, prompt customization
docs/walkthrough.md — detailed custom-task walkthrough
docs/configuration.md — backends, models, API keys, CLI reference
docs/troubleshooting.md — common errors and fixes
If you use SIA in your research, please cite:
@article { hebbar2026sia ,
title = { SIA: Self Improving AI with Harness \& Weight Updates } ,
author = { Hebbar, Prannay and Manawat, Yogendra and Verboomen, Samuel and Ivanova, Alesia and Palanimalai, Selvam and Bhatia, Kunal and Baskaran, Vignesh } ,
journal = { arXiv preprint arXiv:2605.27276 } ,
year = { 2026 } ,
url = { https://arxiv.org/abs/2605.27276 }
}
About
SIA is a Self Improving AI framework to autonomously improve the performance of any AI system (Model / Agent) on a benchmark task.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
82
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
