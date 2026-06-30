---
source: "https://github.com/experientiallabs/world-model-harness"
hn_url: "https://news.ycombinator.com/item?id=48735124"
title: "LLMs as 5x Faster Sandboxes"
article_title: "GitHub - experientiallabs/world-model-harness: World-model-as-a-harness for simulating AI agent environments · GitHub"
author: "SilenN"
captured_at: "2026-06-30T16:42:42Z"
capture_tool: "hn-digest"
hn_id: 48735124
score: 1
comments: 1
posted_at: "2026-06-30T16:30:08Z"
tags:
  - hacker-news
  - translated
---

# LLMs as 5x Faster Sandboxes

- HN: [48735124](https://news.ycombinator.com/item?id=48735124)
- Source: [github.com](https://github.com/experientiallabs/world-model-harness)
- Score: 1
- Comments: 1
- Posted: 2026-06-30T16:30:08Z

## Translation

タイトル: 5 倍高速なサンドボックスとしての LLM
記事のタイトル: GitHub - experiencelabs/world-model-harness: AI エージェント環境をシミュレートするための World-model-as-a-harness · GitHub
説明: AI エージェント環境をシミュレートするための World-model-as-a-harness - experiencelabs/world-model-harness

記事本文:
GitHub - experiencelabs/world-model-harness: AI エージェント環境をシミュレートするための World-model-as-a-harness · GitHub
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
体験ラボ
/
ワールドモデルハーネス
公共
通知

イオン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
エクスペリエンシャルラボ/ワールドモデルハーネス
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
30 コミット 30 コミットの例 例 wmh wmh .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM としての Docker。サンドボックスを立ち上げるのではなく、トレースからエージェント環境をシミュレートします。
フロンティア LLM は、OpenTelemetry から再構築された、エージェントが踏み込む環境として機能します。
跡。ハーネスは記録された (状態、アクション) -> 観察ステップを取り込み、検索インデックスを構築します。
GEPA を使用して基本環境プロンプトを進化させ、結果として得られるワールド モデルをローカルに提供します。
OTel トレースからの構築: 取り込み、正規化、トレイン/ホールドアウトの分割、リプレイ バッファのインデックス付け、および
環境プロンプトを最適化します。
構築されたモデルを提供または再生します。エージェントはプロセス内またはプロセスを通じて WorldModel.step(action) を呼び出します。
ローカル HTTP バックエンド。
トレース ファイルに対して wmh eval を使用して再構築の忠実度を評価します。
UV同期
WMH プロバイダーが検証する
wmh build --name航空会社 --fileexamples/tau-bench/traces.otel.jsonl
うーんリスト
wmh eval 例/tau-bench/traces.otel.jsonl
WMH 評価リスト
WMH 評価実行 tau-bench
WMH 評価結果
WMH の例のリスト
wmh の例では、tau-bench -- --trace 0 を実行します。
うーんサーブ
WMH デモ -- 航空会社名
wmh play --name 航空会社
フラグを指定しない wmh ビルドでは、対話型ターミナル上でガイド付き作成ウィザードが起動します。パス
スクリプト可能な実行の場合は、 --file および関連フラグ、または --no-interactive 。
ワールド モデルには名前が付けられ、 .wmh/models/<name>/ に保存されます。 wmh リスト、wmh サーブ、wmh デモ、
wmh play は、そのディレクトリ内でローカルに構築されたモデルのみを使用します。
コマンド
何

それはあります
うーんビルド
OTel トレースまたはベンダー トレース プルから名前付きワールド モデルを構築します。トレースの取り込み、正規化、トレーニング/保持データの分割、取得インデックスの構築、GEPA プロンプト最適化の実行、アーティファクトの .wmh/models/<name>/ への書き込みを行います。 TTY に必須の入力を行わないと、ガイド付きウィザードが開きます。
うーんリスト
選択したルートの models/ ディレクトリの下にあるワールド モデルをリストします。これには、プロバイダー、ホールドアウト スコア、ロールアウト数、およびメトリクスが存在する場合のフロンティア サイズが含まれます。デフォルトでは、選択されたルートは .wmh/ であるため、プレーンな wmh リストはコミットされたサンプル アーティファクトを読み取りません。
wmh eval <トレース ファイル...>
1 つ以上の OTel トレース ファイルの再構築忠実度をスコアします。決定論的なトレーニング/ホールドアウト分割を実行し、ベースまたは提供されたプロンプトを通じてホールドアウト ステップを再生し、記録された観測値に対して予測された観測値を採点し、ファイルごとと全体の忠実度を出力します。
WMH 評価リスト
Examples/<task>/evals/*.toml からの名前付き eval スイートをリストします。スイートは、反復可能な再構成忠実度の実行のためのサンプルローカル定義です。
WMH 評価実行 <スイート>
構成されたトレース ファイルと分割/スコアリング設定を使用して、名前付き評価スイートを実行します。 --out が指定されていない限り、結果は .wmh/evals/<task>/<suite>/ にローカル JSON として書き込まれます。例のデフォルト スイートはタスク名によって選択できます。 wmh eval tau-bench を実行します。
WMH 評価結果 [スイート]
.wmh/evals/ からローカルに保存された名前付き eval 結果を要約します。これらは生成されたアーティファクトであるため、コミットしないでください。
うーんサーブ
デフォルトでは、ローカル FastAPI バックエンドを 127.0.0.1:8000 で開始します。 /world_models/... HTTP ルートを介して、ローカルで構築されたすべてのモデル、または繰り返された --name 選択のみを提供します。
うーんデモ
構築されたモデルに対して短いデモを実行します。使い捨て LLM エージェントがサンプリングされたトレースの例からアクションを提案する

つまり、ワールド モデルは環境の観察を予測し、CLI はアクション、環境のプロンプト、および観察を出力します。
うーん、遊ぶ
構築されたモデルの対話型 REPL を開きます。あなたはタイプします
[切り捨てられた]
データセット固有のロジックは、examples/ の下にのみ存在します。各タスク フォルダーは自己完結型です。
例/swe-bench/traces.otel.jsonl
例/tau-bench/traces.otel.jsonl
例/ターミナルタスク/traces.otel.jsonl
各サンプル フォルダーには、タスク ローカル キャプチャまたは起動ヘルパーが含まれる場合があります。それらを介して起動します
wmh の例では、 <task> -- <args> を実行します。再利用可能なハーネスの動作は wmh/ に属しており、次のようにする必要があります。
wmh CLI を通じて公開されます。
反復可能な eval スイート定義は、examples/<task>/evals/*.toml の下にあります。彼らが指差すのは、
ローカル トレース ファイルの例を作成し、トレイン分割、サンプリング、RAG、
裁判官。生成された評価結果は、 .wmh/evals/ の下にローカルに残ります。
サンプルローカルの事前構築済みアーティファクトは、examples/<task>/models/<name>/ の下に存在します。パスする
--root example/<task> を wmh list 、 wmh demo 、 wmh play 、または wmhserve に指定せずに使用します。
それを .wmh/ にコピーします。
wmh import アクション、ActionKind から
うーん、から。構成。ストアインポート WorldModelStore
うーん、から。エンジン。ローダーインポートload_world_model
model_dir = WorldModelStore ( ".wmh" )。解決 (「航空会社」)
wm , _provider =load_world_model (model_dir)
セッション = wm 。 new_session ( task = "カートをチェックアウト" )
obs = wm 。ステップ(
セッション。 ID 、
アクション (種類 = ActionKind . TOOL_CALL 、名前 = "add_to_cart" 、引数 = { "sku" : "A1" })、
)
print (obs.content)
HTTP 経由で、 GET /world_models を使用し、次に POST /world_models/{name}/sessions を使用します。
POST /world_models/{name}/sessions/{id}/step 。
資格情報は環境から読み取られます。
UV 同期 --extra dev
UVランラフチェック。
UV 実行ラフ形式。
UV ランタイチェック
uv 実行 pytest -q
慣習が生きている

エージェント.md 。テストは対象となるコードの隣にインラインで表示されます
( foo.py -> foo_test.py ) wmh/ の下にあります。
AI エージェント環境をシミュレートするためのワールド モデル アズ ア ハーネス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

World-model-as-a-harness for simulating AI agent environments - experientiallabs/world-model-harness

GitHub - experientiallabs/world-model-harness: World-model-as-a-harness for simulating AI agent environments · GitHub
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
experientiallabs
/
world-model-harness
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
experientiallabs/world-model-harness
main Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits examples examples wmh wmh .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Docker as an LLM. Simulate an agent environment from traces instead of standing up a sandbox.
A frontier LLM acts as the environment your agent steps against, reconstructed from OpenTelemetry
traces. The harness ingests recorded (state, action) -> observation steps, builds a retrieval index,
evolves the base environment prompt with GEPA, and serves the resulting world model locally.
Build from OTel traces: ingest, normalize, split train/held-out, index the replay buffer, and
optimize the environment prompt.
Serve or play the built model: agents call WorldModel.step(action) in-process or through the
local HTTP backend.
Evaluate reconstruction fidelity with wmh eval against trace files.
uv sync
wmh providers verify
wmh build --name airline --file examples/tau-bench/traces.otel.jsonl
wmh list
wmh eval examples/tau-bench/traces.otel.jsonl
wmh eval list
wmh eval run tau-bench
wmh eval results
wmh examples list
wmh examples run tau-bench -- --trace 0
wmh serve
wmh demo --name airline
wmh play --name airline
wmh build with no flags launches a guided creation wizard on an interactive terminal. Pass
--file and related flags, or --no-interactive , for scriptable runs.
World models are named and stored under .wmh/models/<name>/ . wmh list , wmh serve , wmh demo ,
and wmh play only use models built locally in that directory.
Command
What it does
wmh build
Builds a named world model from OTel traces or a vendor trace pull. It ingests traces, normalizes them, splits train/held-out data, builds the retrieval index, runs GEPA prompt optimization, and writes the artifact to .wmh/models/<name>/ . With no required inputs on a TTY, it opens the guided wizard.
wmh list
Lists world models found under the selected root's models/ directory, including provider, held-out score, rollout count, and frontier size when those metrics exist. By default, the selected root is .wmh/ , so plain wmh list does not read committed example artifacts.
wmh eval <trace files...>
Scores reconstruction fidelity on one or more OTel trace files. It performs a deterministic train/held-out split, replays held-out steps through the base or supplied prompt, grades predicted observations against recorded observations, and prints per-file plus overall fidelity.
wmh eval list
Lists named eval suites from examples/<task>/evals/*.toml . Suites are example-local definitions for repeatable reconstruction-fidelity runs.
wmh eval run <suite>
Runs a named eval suite, using its configured trace files and split/scoring settings. Results are written as local JSON under .wmh/evals/<task>/<suite>/ unless --out is supplied. The default suite for an example can be selected by task name, e.g. wmh eval run tau-bench .
wmh eval results [suite]
Summarizes locally saved named eval results from .wmh/evals/ . These are generated artifacts and should not be committed.
wmh serve
Starts the local FastAPI backend on 127.0.0.1:8000 by default. It serves all locally built models, or only the repeated --name selections, through /world_models/... HTTP routes.
wmh demo
Runs a short demo against a built model. A throwaway LLM agent proposes an action from sampled trace examples, the world model predicts the environment observation, and the CLI prints the action, environment prompt, and observation.
wmh play
Opens an interactive REPL for a built model. You typ
[truncated]
Dataset-specific logic lives only under examples/ . Each task folder is self-contained:
examples/swe-bench/traces.otel.jsonl
examples/tau-bench/traces.otel.jsonl
examples/terminal-tasks/traces.otel.jsonl
Each example folder may include task-local capture or launch helpers. Launch them through
wmh examples run <task> -- <args> . Reusable harness behavior belongs in wmh/ and should be
exposed through the wmh CLI.
Repeatable eval suite definitions live under examples/<task>/evals/*.toml . They point at
example-local trace files and configure replay options such as train split, sampling, RAG, and
judge. Generated eval results stay local under .wmh/evals/ .
Example-local prebuilt artifacts live under examples/<task>/models/<name>/ ; pass
--root examples/<task> to wmh list , wmh demo , wmh play , or wmh serve to use one without
copying it into .wmh/ .
from wmh import Action , ActionKind
from wmh . config . store import WorldModelStore
from wmh . engine . loader import load_world_model
model_dir = WorldModelStore ( ".wmh" ). resolve ( "airline" )
wm , _provider = load_world_model ( model_dir )
session = wm . new_session ( task = "check out the cart" )
obs = wm . step (
session . id ,
Action ( kind = ActionKind . TOOL_CALL , name = "add_to_cart" , arguments = { "sku" : "A1" }),
)
print ( obs . content )
Over HTTP, use GET /world_models , then POST /world_models/{name}/sessions and
POST /world_models/{name}/sessions/{id}/step .
Credentials are read from the environment.
uv sync --extra dev
uv run ruff check .
uv run ruff format .
uv run ty check
uv run pytest -q
Conventions live in AGENTS.md . Tests are inline next to the code they cover
( foo.py -> foo_test.py ) under wmh/ .
World-model-as-a-harness for simulating AI agent environments
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
