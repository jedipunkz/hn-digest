---
source: "https://github.com/botcircuits-ai/botcircuits-agent"
hn_url: "https://news.ycombinator.com/item?id=48350800"
title: "New AI Agent Architecture to fix LLM deviations and token costs"
article_title: "GitHub - botcircuits-ai/botcircuits-agent: Workflow-native AI agent for predictable and token-efficient multi-step automations · GitHub"
author: "nexcatara"
captured_at: "2026-06-01T00:24:37Z"
capture_tool: "hn-digest"
hn_id: 48350800
score: 1
comments: 0
posted_at: "2026-05-31T23:21:43Z"
tags:
  - hacker-news
  - translated
---

# New AI Agent Architecture to fix LLM deviations and token costs

- HN: [48350800](https://news.ycombinator.com/item?id=48350800)
- Source: [github.com](https://github.com/botcircuits-ai/botcircuits-agent)
- Score: 1
- Comments: 0
- Posted: 2026-05-31T23:21:43Z

## Translation

タイトル: LLM の逸脱とトークンコストを修正するための新しい AI エージェント アーキテクチャ
記事のタイトル: GitHub - botcircuits-ai/botcircuits-agent: 予測可能でトークン効率の高いマルチステップ自動化のためのワークフローネイティブ AI エージェント · GitHub
説明: 予測可能でトークン効率の高いマルチステップ自動化のためのワークフロー ネイティブ AI エージェント - botcircuits-ai/botcircuits-agent

記事本文:
GitHub - botcircuits-ai/botcircuits-agent: 予測可能でトークン効率の高いマルチステップ自動化のためのワークフローネイティブ AI エージェント · GitHub
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
ボットサーキット-ai
/
ボットサーキットエージェント
公共
通知
あなたはきっとsでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
ボットサーキット-ai/ボットサーキット-エージェント
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .botcircuits .botcircuits .vscode .vscode docs docs scripts scripts skill/ botcircuits-faq skill/ botcircuits-faq src/ botcircuits src/ botcircuits .env.example .env.example .gitignore .gitignore .python-version .python-version IMPLEMENTATION.md IMPLEMENTATION.md ライセンス ライセンス README.md README.md main.py main.py pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ワークフロー ネイティブの AI エージェントでは、LLM が各ステップの推論とツール呼び出しを処理し、決定論的なステート マシンが全体のフローを制御します。その結果、LLM にすべてを依存することなく、予測可能でトークン効率の高いマルチステップの自動化が実現します。
クイックインストール
Linux、macOS、WSL2、Termux
カール -fsSL https://raw.githubusercontent.com/botcircuits-ai/botcircuits-agent/main/scripts/install.sh |バッシュ
またはクローンを作成してインストールします
# 1. UV をインストールします (既にお持ちの場合はスキップしてください)
カール -LsSf https://astral.sh/uv/install.sh |しー
# または: brew install UV
#2. リポジトリのクローンを作成する
git clone https://github.com/botcircuits-ai/botcircuits-agent
cd ボットサーキットエージェント
# 3. Python (3.11+) を選択し、プロジェクト venv を作成します
uv Python インストール 3.11
uv venv --python 3.11 # ./.venv を作成します
# 4. venv をアクティブ化する
ソース .venv/bin/activate # bash / zsh
# 5. 依存関係を venv にインストールする
UV同期
プロバイダー、モデル、API キーを構成します。
ボットサーキットのセットアップ
ウィザードでは、矢印キー ナビゲーション (↑/↓ で移動、Enter で選択、Esc で現在の値を保持) を使用して、プロバイダー ( anthropic / openai / gemini )、モデル、および API キーを案内します。各選択は、実行中に保存されます。
プロバイダーとモデル →

~/.botcircuits/settings.json
APIキー → ~/.botcircuits/.env ( ANTHROPIC_API_KEY / OPENAI_API_KEY / GEMINI_API_KEY 、ファイルモード 0600 )
ボットサーキットのセットアップを再実行すると、既存の値がデフォルトとして表示され、既存の API キーを使用すると、シークレットの入力を再度求める代わりに、「保持」/「置換」/「クリア」の選択肢が表示されます。
手動で設定したいですか?代わりに env テンプレートをコピーします。
cp .env.example .env
# .env — 使用するプロバイダーの API キー (必須)
ANTHROPIC_API_KEY=...
OPENAI_API_KEY=...
GEMINI_API_KEY=...
# オプション — settings.json / CLI フラグで設定されない場合のフォールバックとしてのみ使用されます
LLM_PROVIDER=人間 # 人間 |オープンナイ |ジェミニ
ANTHROPIC_MODEL=クロード-作品-4-7
OPENAI_MODEL=gpt-4.1
GEMINI_MODEL=ジェミニ-2.5-フラッシュ
有効な優先順位 (最も高い優先順位): CLI フラグ → settings.json (階層化) → 環境変数 → 組み込みのデフォルト 。
ボット回路
ボットサーキット --プロバイダー openai
単一のメッセージをパイプします (非対話型):
echo "2+2 とは何ですか?" |ボットサーキット -- ストリームなし
FastAPI ゲートウェイ (HTTP + メッセージング チャネル用):
uv run uvicorn botcircuits.gateway:app --reload --port 8000
# または
ボットサーキットゲートウェイ
便利な CLI フラグ
旗
説明
--プロバイダー
人間的 (デフォルト) |オープンナイ |ジェミニ
--モデル
プロバイダーのデフォルトモデルをオーバーライドする
--stream / --no-stream
ストリーミングの強制オン/オフ
--自動
すべてのゲート ツールで y/N の確認をスキップします。各アクションの前に警告が出力されます。
--config <パス>
特定の settings.json をロードします (自動検出されたファイルに加えて)
設定ファイル（オプション）
CLI はこれらを順番に自動ロードします (後のレイヤーが優先されます)。
CLI フラグは常に JSON よりも優先されます。スターター ファイルは .botcircuits/settings.example.json にあります。
コマンド
アクション
/ヘルプ
すべてのコマンドを表示
/リセット
現在のセッションをクリアして新たに開始する
/セッション[ID]
セッションIDの表示または切り替え
/ストリーム オン|オフ
ストリーミングの切り替え
/ツール

s
モデルが呼び出すことができるツールをリストします。
/スキル
ファイルシステムのスキルをリストする
/メモリ
永続的なメモリを表示する
/workflow 追加 "<プロンプト>"
自然言語から新しいワークフローを作成する
/workflow edit "<プロンプト>" --name <wf>
既存のワークフローを編集する
/workflow run --name <wf> [--initial-args '{"k":"v"}']
モデルのツール選択をバイパスして、ワークフロー ツールを強制的に開始します。
/やめる
終了
複数行のメッセージを開始するには、単独の行に「"」を入力します (また終了するには再度「"」を入力します)。
ワークフローは、エージェントがユーザーに代わって実行できる段階的な会話スクリプトです。各ワークフローは .botcircuits/workflows/ に存在する 1 つの JSON ファイルであり、エージェントがそれを読み込むと、ワークフローはモデルが名前で呼び出すことができるツールになります。
JSON を手動で編集する必要はありません。 /workflow スラッシュ コマンドを使用してすべてを実行します。
ボット回路
/workflow add --name workflow_demo "1 つの初期引数 end_id (早期終了を制御する 'step_3' や 'step_7' などの文字列) を受け取る、合計 11 のステップ (step_1 から step_10 と終了ステップ) を持つワークフローを作成します。番号付きの各ステップのアクションは、「現在のステップ番号と現在の日付と時刻を含む現在のディレクトリに <step_id>.md という名前のマークダウン ファイルを作成します。たとえば、step_3 は次の内容の step_3.md を作成します。ステップ 3 — <現在の日付と時刻>」です。 step_1 から step_9 にはそれぞれ分岐条件があります。end_id が現在のステップ ID と等しい場合は終了ステップに進み、そうでない場合は次の番号のステップに進みます (step_1 → end if end_id==step_1 else step_2、step_2 → end if end_id==step_2 else step_3、... から step_9 → end if end_id==step_9 else step_10)。 step_10 には条件がなく、アクションの後に直接終了ステップに進みます。終了ステップのアクションは、「END を正確に含む、現在のディレクトリに end.md という名前のマークダウン ファイルを作成する」です。
エージェントはワークフローの草案を作成し、プレビューを表示し、y/N を尋ねます。

ファイルを書き込み、ツールとして登録します。新しいワークフローは、次のメッセージで呼び出し可能になります。再起動は必要ありません。
デフォルトでは、モデルはワークフロー名のスラッグを選択します。 --name <wf> を渡して自分で設定します。その値は、JSON ファイル名 ( .botcircuits/workflows/<wf>.json ) と登録されたツール名の両方になります。
/workflow add "<prompt>" --name check_order_status
--name はスラッグセーフである必要があります (文字、数字、_ 、 - )。
既存のワークフローを変更するには:
/workflow edit "払い戻しも処理します" --name check_order_status
エージェントは現在のファイルを読み取り、編集を適用し、 y/N を尋ね、ライブ ツールを更新します。
ワークフローを呼び出すかどうかの決定をモデルに任せたくない場合は、ワークフローを直接開始します。
/workflow run --name workflow_demo
/workflow run --name workflow_demo --initial-args '{"end_id":"step_3"}'
これにより、指定した引数を使用してワークフロー ツールがすぐに呼び出され、結果の最初のステップで会話がシードされ、制御がモデルに戻されて実行されます。 --initial-args は JSON オブジェクトである必要があります。 {} で始める場合は省略します。ターゲット ワークフローはすでに登録されている必要があります。ワークフロー ツールは .botcircuits/workflows/.build/ から自動検出され、コマンドは名前を検索する前にそのレジストリを更新するため、新しく作成されたワークフローは再起動しなくても機能します。
デフォルトでは、ワークフローは現在のディレクトリの下の .botcircuits/workflows/*.json に存在します。 BOTCIRCUITS_WORKFLOWS_DIR=/abs/path でオーバーライドします (または .env に設定します)。ディレクトリが見つからないということは、単に「ワークフローがない」ことを意味します。オプトインするには、フォルダーをドロップします。
各ファイルは 1 つのワークフロー レコードです。
{
"name" : "greet_user " ,
"description" : " 発信者に挨拶してから別れを告げます。 " ,
「フロー」: {
"開始" : " s0 " 、
「ステップ」: {
"s0" : { "type" : " start " 、 "next" : " a1 " },
"a1" : {
"type" : " エージェントアクション

n "、
"次" : " a2 " 、
「条件」：[
{ "condition" : " 要求されたトーンは暖かい " , "next" : " a2 " }
]、
「設定」: {
"action" : " 目的のトーンをキャプチャし、それに応じてユーザーに挨拶します。 」
}
}、
"a2" : {
"type" : " エージェントアクション " ,
"settings" : { "action" : " さようなら。 " }
}
}
}
}
name はモデルが呼び出すツール名と同じです。 ^[a-zA-Z0-9_-]+$ と一致する必要があります。 start および agentAction ステップ タイプのみがサポートされます。分岐するには、ステップ ルート ( type と next の兄弟、 settings 内にネストされていない) に条件リストをアタッチします。条件は制御フローであるため、 settings のステップタイプ固有のペイロードではなく、他の制御フロー フィールドの隣に配置されます。
あなたが作成した生のファイルは、エンジンが実行するものではありません。ワークフローのビルド ステップでは、各 AgentAction の自然言語条件を取得し、エンジンが決定論的に評価できる条件と変数を準備します。
条件 — 各 NL 条件 (例: 「要求されたトーンは暖かい」) は、演算子 ( is 、 >= 、 contains 、 …) と値を含む型付きの Choices[] エントリにコンパイルされるため、エンジンは実行時に LLM を再度呼び出すことなく、一致するブランチを選択できます。
変数 — 集約された flow.variables リストが出力され、コンパイルされた条件によって参照されるすべてのスロットに、その推論された dataType と簡単な説明とともに名前が付けられます。ランタイムはこのリストを使用して、分岐を評価する前に LLM の自由テキスト引数を正しい形式に強制します。
/workflow add|edit はビルダーを実行します。ワークフロー ファイルを手動で編集する場合は、CLI から再構築します。
botcircuits ワークフロー ビルド --name=greet_user
エージェント ランタイムは .botcircuits/workflows/.build/ からワークフローのみをロードするため、構築されていないワークフローは呼び出すことができません。ワークフローのビルドによって実行可能なコピーが生成されます。
スキルは命令の小さなフォルダーです

エージェントがディスクから読み取るオプション (およびオプションで許可されたツール)。スキルは、「サポートの質問にどのように答えるか」、「PR の説明をどのように下書きするか」など、システム プロンプトにパターンを組み込むことなく、反復可能なパターンを把握するのに役立ちます。
CLI は以下からスキルを検出します。
.botcircuits/skills/ (プロジェクト)
各スキルは SKILL.md ファイルを含むフォルダーです。
スキル/
└── ボットサーキット-faq/
━── SKILL.md
---
名前 : ボットサーキット よくある質問
description : BotCircuits に関する質問に答えます — 機能、価格、ドキュメント。
許可されたツール: playwright__browser_navigate、playwright__browser_snapshot
---
あなたは予備知識から BotCircuits について知りません。唯一の情報源
真実は https://botcircuits.ai/ です。答える前にサイトを取得してください...
説明は、モデルがスキルをいつ呼び出すかを決定するために使用するものです。 allowed-tools (オプション) は、スキルの実行中に呼び出しを許可されるツールを制限します。
/skills # 発見されたスキルをリストします
/<skill-name> # スキルを直接実行します (モデルをバイパスします)
エージェントは説明に基づいて独自にスキルを選択することもできます。
MCP サーバーは、外部ツール (ファイルシステム、GitHub、データベースなど) をエージェントに公開します。これらを一度設定すると、すべての CLI セッションおよびゲートウェイ要求で使用できるようになります。
MCP サーバー構成は mcp.json ファイル内に存在します。 l

[切り捨てられた]

## Original Extract

Workflow-native AI agent for predictable and token-efficient multi-step automations - botcircuits-ai/botcircuits-agent

GitHub - botcircuits-ai/botcircuits-agent: Workflow-native AI agent for predictable and token-efficient multi-step automations · GitHub
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
botcircuits-ai
/
botcircuits-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
botcircuits-ai/botcircuits-agent
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .botcircuits .botcircuits .vscode .vscode docs docs scripts scripts skills/ botcircuits-faq skills/ botcircuits-faq src/ botcircuits src/ botcircuits .env.example .env.example .gitignore .gitignore .python-version .python-version IMPLEMENTATION.md IMPLEMENTATION.md LICENSE LICENSE README.md README.md main.py main.py pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
The workflow-native AI agent where an LLM handles the reasoning and tool calls for each step, while a deterministic state machine controls the overall flow. The result:predictable and token-efficient multi-step automation without depending on an LLM to drive everything.
Quick install
Linux, macOS, WSL2, Termux
curl -fsSL https://raw.githubusercontent.com/botcircuits-ai/botcircuits-agent/main/scripts/install.sh | bash
Or Clone and install
# 1. Install uv (skip if you already have it)
curl -LsSf https://astral.sh/uv/install.sh | sh
# or: brew install uv
# 2. Clone the repo
git clone https://github.com/botcircuits-ai/botcircuits-agent
cd botcircuits-agent
# 3. Pick a Python (3.11+) and create the project venv
uv python install 3.11
uv venv --python 3.11 # creates ./.venv
# 4. Activate the venv
source .venv/bin/activate # bash / zsh
# 5. Install dependencies into the venv
uv sync
Configure your provider, model, and API key:
botcircuits setup
The wizard walks you through provider ( anthropic / openai / gemini ), model, and API key with arrow-key navigation (↑/↓ to move, Enter to select, Esc to keep the current value). Each pick is saved as you go:
provider and model → ~/.botcircuits/settings.json
API key → ~/.botcircuits/.env ( ANTHROPIC_API_KEY / OPENAI_API_KEY / GEMINI_API_KEY , file mode 0600 )
Re-running botcircuits setup shows your existing values as defaults, and an existing API key gives you a Keep / Replace / Clear choice instead of re-prompting for the secret.
Prefer to configure by hand? Copy the env template instead:
cp .env.example .env
# .env — API key for the provider you want to use (required)
ANTHROPIC_API_KEY=...
OPENAI_API_KEY=...
GEMINI_API_KEY=...
# Optional — only used as a fallback when settings.json / CLI flags don't set them
LLM_PROVIDER=anthropic # anthropic | openai | gemini
ANTHROPIC_MODEL=claude-opus-4-7
OPENAI_MODEL=gpt-4.1
GEMINI_MODEL=gemini-2.5-flash
Effective precedence (highest wins): CLI flag → settings.json (layered) → env var → built-in default .
botcircuits
botcircuits --provider openai
Pipe a single message (non-interactive):
echo " what is 2+2? " | botcircuits --no-stream
FastAPI gateway (for HTTP + messaging channels):
uv run uvicorn botcircuits.gateway:app --reload --port 8000
# or
botcircuits-gateway
Useful CLI flags
Flag
Description
--provider
anthropic (default) | openai | gemini
--model
Override the provider's default model
--stream / --no-stream
Force streaming on/off
--auto
Skip y/N confirmation on every gated tool. A warning still prints before each action.
--config <path>
Load a specific settings.json (in addition to the auto-discovered files)
Settings files (optional)
The CLI auto-loads these in order (later layers win):
CLI flags always win over JSON. A starter file is at .botcircuits/settings.example.json .
Command
Action
/help
Show all commands
/reset
Clear the current session and start fresh
/session [id]
Show or switch session id
/stream on|off
Toggle streaming
/tools
List the tools the model can call
/skills
List filesystem skills
/memory
Show persistent memory
/workflow add "<prompt>"
Author a new workflow from natural language
/workflow edit "<prompt>" --name <wf>
Edit an existing workflow
/workflow run --name <wf> [--initial-args '{"k":"v"}']
Force-start a workflow tool, bypassing the model's tool choice
/quit
Exit
Type """ on its own line to start (and again to end) a multi-line message.
A workflow is a step-by-step conversation script the agent can run on your behalf. Each workflow is one JSON file that lives under .botcircuits/workflows/ , and once the agent loads it the workflow becomes a tool the model can call by name.
You don't have to hand-edit JSON. Drive everything through the /workflow slash command:
botcircuits
/workflow add --name workflow_demo "Create a workflow with 11 steps total (step_1 through step_10 plus an end step) that takes one initial argument end_id (a string like 'step_3' or 'step_7' controlling early termination); each numbered step's action is 'Create a markdown file named <step_id>.md in the current directory containing the current step number and current date and time, e.g., step_3 creates step_3.md with content: Step 3 — <current date and time>'; step_1 through step_9 each have a branching condition: if end_id equals the current step id go to the end step, otherwise go to the next numbered step (step_1 → end if end_id==step_1 else step_2, step_2 → end if end_id==step_2 else step_3, ... through step_9 → end if end_id==step_9 else step_10); step_10 has no condition and goes directly to the end step after its action; the end step's action is 'Create a markdown file named end.md in the current directory containing exactly: END'."
The agent drafts the workflow, shows you a preview, asks y/N , writes the file, and registers it as a tool. The new workflow becomes callable on the very next message — no restart.
By default the model picks a slug for the workflow name . Pass --name <wf> to set it yourself — that value becomes both the JSON filename ( .botcircuits/workflows/<wf>.json ) and the registered tool name:
/workflow add "<prompt>" --name check_order_status
--name must be slug-safe (letters, digits, _ , - ).
To change an existing workflow:
/workflow edit "also handle refunds" --name check_order_status
The agent reads the current file, applies your edit, asks y/N , and refreshes the live tool.
When you don't want to leave it up to the model to decide whether to invoke a workflow, kick one off directly:
/workflow run --name workflow_demo
/workflow run --name workflow_demo --initial-args '{"end_id":"step_3"}'
This calls the workflow tool right away with the args you supplied, seeds the conversation with the resulting first step, and hands control back to the model to perform it. --initial-args must be a JSON object; omit it to start with {} . The target workflow must already be registered — workflow tools are auto-discovered from .botcircuits/workflows/.build/ and the command refreshes that registry before looking up the name, so a freshly authored workflow works without a restart.
By default, workflows live in .botcircuits/workflows/*.json under the current directory. Override with BOTCIRCUITS_WORKFLOWS_DIR=/abs/path (or set it in .env ). A missing directory just means "no workflows" — drop a folder in to opt in.
Each file is one workflow record:
{
"name" : " greet_user " ,
"description" : " Greet the caller, then say goodbye. " ,
"flow" : {
"start" : " s0 " ,
"steps" : {
"s0" : { "type" : " start " , "next" : " a1 " },
"a1" : {
"type" : " agentAction " ,
"next" : " a2 " ,
"conditions" : [
{ "condition" : " the requested tone is warm " , "next" : " a2 " }
],
"settings" : {
"action" : " Capture the desired tone and greet the user accordingly. "
}
},
"a2" : {
"type" : " agentAction " ,
"settings" : { "action" : " Say goodbye. " }
}
}
}
}
name doubles as the tool name the model calls; it must match ^[a-zA-Z0-9_-]+$ . Only start and agentAction step types are supported — to branch, attach a conditions list at the step root (sibling of type and next , not nested inside settings ). conditions is control flow, so it sits next to the other control-flow fields rather than with the step-type-specific payload in settings .
The raw file you author is not what the engine runs. The workflow build step takes the natural-language conditions on each agentAction and prepares conditions and variables the engine can evaluate deterministically:
Conditions — each NL condition (e.g. "the requested tone is warm" ) is compiled into a typed choices[] entry with an operator ( is , >= , contains , …) and a value, so the engine can pick the matching branch without re-calling the LLM at runtime.
Variables — an aggregated flow.variables list is emitted, naming every slot referenced by the compiled conditions along with its inferred dataType and a short description. The runtime uses this list to coerce the LLM's free-text args into the right shape before evaluating branches.
/workflow add|edit runs the builder for you. If you hand-edit a workflow file, re-build it from the CLI:
botcircuits workflow build --name=greet_user
The agent runtime only loads workflows from .botcircuits/workflows/.build/ , so a workflow that hasn't been built isn't callable — workflow build is what produces the runnable copy.
A skill is a small folder of instructions (and optionally allowed tools) the agent reads from disk. Skills are useful for capturing repeatable patterns — "how we answer support questions", "how we draft a PR description" — without baking them into the system prompt.
The CLI discovers skills from:
.botcircuits/skills/ (project)
Each skill is a folder with a SKILL.md file:
skills/
└── botcircuits-faq/
└── SKILL.md
---
name : botcircuits-faq
description : Answer questions about BotCircuits — features, pricing, docs.
allowed-tools : playwright__browser_navigate, playwright__browser_snapshot
---
You do NOT know about BotCircuits from prior knowledge. The only source of
truth is https://botcircuits.ai/ . Fetch the site before answering...
The description is what the model uses to decide when to invoke the skill. allowed-tools (optional) restricts which tools the skill is allowed to call during its run.
/skills # list discovered skills
/<skill-name> # run a skill directly (bypass the model)
The agent can also pick a skill on its own based on the description.
MCP servers expose external tools (filesystem, GitHub, databases, …) to the agent. You can configure them once and they become available across every CLI session and gateway request.
MCP server configs live in mcp.json files, l

[truncated]
