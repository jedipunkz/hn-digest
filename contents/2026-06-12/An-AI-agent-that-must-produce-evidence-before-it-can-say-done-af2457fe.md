---
source: "https://github.com/Aspct3434/Distill-Agent"
hn_url: "https://news.ycombinator.com/item?id=48506215"
title: "An AI agent that must produce evidence before it can say \"done\""
article_title: "GitHub - Aspct3434/Distill-Agent: Self-hosted AI agent that can't say \"done\" without proof — evidence-gated task contracts, self-distilled skills with auto-rollback, multi-channel. · GitHub"
author: "aspctpro"
captured_at: "2026-06-12T18:47:17Z"
capture_tool: "hn-digest"
hn_id: 48506215
score: 2
comments: 0
posted_at: "2026-06-12T16:33:03Z"
tags:
  - hacker-news
  - translated
---

# An AI agent that must produce evidence before it can say "done"

- HN: [48506215](https://news.ycombinator.com/item?id=48506215)
- Source: [github.com](https://github.com/Aspct3434/Distill-Agent)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T16:33:03Z

## Translation

タイトル: 「完了」と言う前に証拠を提出する必要がある AI エージェント
記事のタイトル: GitHub - Aspct3434/Distill-Agent: 証拠がなければ「完了」とは言えない自己ホスト型 AI エージェント — 証拠ゲート型タスク契約、自動ロールバックを備えた自己蒸留スキル、マルチチャネル。 · GitHub
説明: 証拠がなければ「完了」とは言えない自己ホスト型 AI エージェント — 証拠ゲート型タスク契約、自動ロールバックを備えた自己蒸留スキル、マルチチャネル。 - Aspct3434/蒸留剤

記事本文:
GitHub - Aspct3434/Distill-Agent: 証拠がなければ「完了」とは言えない自己ホスト型 AI エージェント — 証拠ゲート型タスク契約、自動ロールバックを備えた自己蒸留スキル、マルチチャネル。 · GitHub
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
アスペクト3434
/
蒸留剤
公共
通知
Y

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
138 コミット 138 コミット .github/ workflows .github/ workflows bin bin control-panel control-panel docs docs 例 例 lib lib ペルソナ ペルソナ スクリプト スクリプト スキル スキル src src テスト テスト .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md an-api.env.example an-api.env.example cli_chat.py cli_chat.py docker-compose.yml docker-compose.yml fly.toml fly.toml ollama.env.example ollama.env.example package.json package.json pyproject.toml pyproject.toml Railway.json Railway.json render.yaml render.yaml 要件-ml.txt 要件-ml.txt 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
証拠がなければ「完了」とは言えない自己ホスト型 AI エージェント。
すべての Distill の実行はタスク コントラクトによって管理されます。終了する前に、エージェントはファイルの書き込み、テストの合格、ポートでのサービスの応答などの物理的な証拠を提出する必要があります。証拠もなければ「完了」もありません。
また、Distill はセッション全体でプロジェクトを学習し、経験から独自の再利用可能なスキルを抽出し (それぞれバージョン管理され、新しいバージョンが目に見えて低下した場合は自動的にロールバックされます)、永続的なハイブリッド メモリを保持し、Telegram、Discord、Slack、電子メール、リアルタイム ストリーミング コントロール パネル、またはターミナルを介して連絡します。
次のコマンドを使用して対話型ターミナル インターフェイスを起動します。
グラフTD
%% 外部インターフェイス
UI[コントロールパネルUI<br/>React + Tailwind]
アダプター[メッセージングアダプター<br/>テレグラム、ディスコード、スラック、電子メール]
TUI[ターミナルUI]
%% ゲートウェイと同時実行性
サブグラフ ゲートウェイ L

エイヤー
API[FastAPIゲートウェイ]
WS[Webソケットストリーム]
キュー[セッション FIFO キュー]
API --- WS
WS --> キュー
終わり
UI --> API
アダプター --> API
TUI --> WS
%% コア エージェント エンジン
サブグラフ コア エージェント エンジン
エージェント[エージェント エンジンの ReAct ループ]
Contract[業務契約システム<br/>証拠管理]
計画[計画管理<br/>行動台帳]
Eval[スキル蒸留者<br/>評価者]
エージェント <--> 契約
エージェント <--> 計画
エージェント --> 評価
終わり
キュー --> エージェント
%% 状態とメモリ
サブグラフの状態とメモリ
チェックポイント[(SQLite チェックポインタ<br/>状態と履歴)]
Chroma[(ChromaDB <br/>セマンティックメモリ)]
Neo4j[(Neo4j<br/>グラフメモリ)]
エージェント <--> チェックポイント
エージェント <--> クロマ
エージェント <--> Neo4j
終わり
%% 外部サービス
LLM((LLMプロバイダー<br/>LiteLLM))
エージェント <--> LLM
%% ツールと実行
サブグラフのツールと実行サンドボックス
ツール[ツールマネージャー]
MCP[MCPサーバー]
サンドボックス[ターミナルサンドボックス<br/>Docker / ホスト / サーバーレス]
スキル[進化後のスキル<br/>オートメーカー]
エージェント --> ツール
ツール --> MCP
ツール --> サンドボックス
ツール --> スキル
Eval -.->|合成と検証|スキル
終わり
読み込み中
ディレクトリ構造
ソース/
Agent.py -- ReAct ループ、セッション管理、チェックポイント設定
Contract.py -- タスク契約システム (証拠追跡)
evaluator.py -- スキル蒸留器: 軌跡 -> 再利用可能な MCP スキル
Memory.py -- HybridMemory: ChromaDB (セマンティック) + Neo4j (グラフ)
Gateway.py -- FastAPI アプリ、WebSocket ストリーム、FIFO レーンごとのセッション
tools.py -- MCP サーバー、ターミナル、ファイル、プロセス、ポート ツール
control-panel/ -- ライブ トークン ストリーミングを使用した React + Tailwind チャット UI
これらの決定の背後にある理由と、その代替案
拒否されました — docs/DESIGN.md に記載されています。
タスク コントラクトの実行 : エージェントは、タスクを開始する前に、必要な実行証拠 (ファイルの作成、サービスの実行) を宣言する必要があります。最終応答はこの phy でゲート制御されます

科学的証拠により、「今すぐやる」という幻覚を排除します。
スキルの蒸留 : 複雑なタスクが成功した後、LLM は軌道を評価し、パラメーター化された Python ツールを合成します。新しいスキルはバージョン管理され、検証され、成功率が低下した場合は自動的にロールバックされます。
FIFO レーンごとのセッションの同時実行性: すべてのユーザー セッションは専用のキューとワーカー タスクを取得し、厳密なメッセージ順序付けによる高い同時実行性を実現します。
ハイブリッド メモリ : SQLite 全文検索、ChromaDB セマンティック埋め込み、Neo4j グラフ リレーションシップを組み合わせて、セッション間のコンテキストを呼び出します。
ユニバーサル サンドボックス: シェル操作をローカル、Docker、またはサーバーレス サンドボックス (Daytona、E2B、Modal など) の前面に配置できる単一の HTTP 実行シムを介してリモートで実行します。
共有可能なスキル: オープンな SKILL.md 形式を介して、抽出されたスキルをエクスポートおよびインポートします。
Distill は、安定したコアとその他の機能を備えた研究フレームワークです。
それに関する実験的な機能。このマトリックスは、各サブシステムがどこにあるかを示します。
スタンド — 実験的とマークされたものはすべて変更される可能性があるものとして扱います。
クラウドでエージェント ゲートウェイをスピンアップし、LLM API キー (およびオプションで
ボット トークンにテレグラムを送信し、エージェントにテキスト メッセージを送信します。無駄のない実行 — SQLite ベースのメモリ、なし
プロビジョニングするデータベース。
レンダリングは render.yaml を読み取り、ゲートウェイ トークンを生成します。
プロバイダー キーとモデルの入力を求められます。
鉄道はrailway.jsonを使用します：ここからプロジェクトを作成します
リポジトリを開き、 AGENT_API_TOKEN 、 AGENT_MODEL 、およびプロバイダー キーを設定します。
Fly.io : fly.toml に対して --copy-config を起動し、その後
fly Secrets set … (コマンドはファイルヘッダーにあります)。
./scripts/quickstart.sh
強力なシークレットを生成し、env ファイルを作成し、完全なスタックを構築します
Docker Compose を使用します。 LLM API キーを an-api.env に追加して再実行します。
# 1. 環境テンプレートをコピーする
cp an-api.env.e

サンプル an-api.env
# 2. LLM API キーを an-api.env に追加します (例: MOONSHOT_API_KEY または OPENAI_API_KEY)
# デフォルトでは、Distill は LiteLLM を使用し、OpenAI、Anthropic、Gemini、Ollama などをサポートします。
# 3. バンドルされたグラフメモリ データベースの Neo4j パスワードを設定します (必須。
# 安全でないデフォルトはありません)。 Compose は .env からそれを読み取ります。
echo " NEO4J_PASSWORD= $( openssl rand -hex 24 ) " >> .env
# 4. アプリケーションを起動する
docker 構成 -d --build
コントロールパネル (UI) : http://localhost:5173
API / ドキュメント : http://localhost:8000/docs
空のマシン (ノードも Python もありません)?
ブートストラップ スクリプトは前提条件をインストールしてから、
インタラクティブなインストーラー。これらは、新しい PC 上で推奨されるパスです。
irm https://raw.githubusercontent.com / Aspct3434 / Distill - エージェント / マスター / スクリプト / bootstrap.ps1 |アイエックス
macOS / Linux (bash):
カール -fsSL https://raw.githubusercontent.com/Aspct3434/Distill-Agent/master/scripts/bootstrap.sh |バッシュ
最初にリポジトリのクローンを作成して、script/bootstrap.ps1 (Windows) を実行することもできます。
scripts/bootstrap.sh (macOS/Linux) を直接実行します。
npx を使用して対話型インストーラーを実行します (グローバル インストールは必要ありません)。
npx @aspct / distill - エージェントのインストール
正確な GitHub リビジョンに固定したいですか?リポジトリ フォームを使用します。
npx --yes github:Aspct3434/Distill-Agent インストール
インストール後、CLI を使用してエージェントを管理します。
npm i -g @aspct/distill-agent
distill # インタラクティブなターミナル UI を開きます
distill start # バックエンドとコントロールパネルを起動します
ログを蒸留する # 実行中のログを表示する
distill update # 最新の変更をプルする
distill Doctor # インストールと環境を診断する
# グローバルにインストールしたくない場合は、npx も機能します。
npx @aspct/distill-agent start # バックエンドとコントロール パネルを起動します
npx @aspct/distill-agent logs # 実行ログを表示する
npx @aspct/distill-agent update # Pul

l 最新の変更点
npx @aspct/distill-agent Doctor # インストールと環境を診断する
構成
Distill は環境変数を介して設定されます。 an-api.env の主要な設定:
AGENT_MODEL は、任意の LiteLLM プロバイダー文字列を受け入れます。設置者ができることは、
kimi/Moonshot、Ollama、OpenRouter、OpenAI、Anthropic、Gemini、を事前設定します。
DeepSeek、Groq、xAI Grok、Mistral、または任意の OpenAI 互換エンドポイント (vLLM) —
一致する *_API_KEY を an-api.env に設定します。
OpenAI の場合は、ChatGPT アカウント (Codex OAuth) でサインインすることもできます。
キーを貼り付ける代わりに、インストーラーでキーを選択するか、実行します。
PYTHONPATH=src python -m auth login (コントロール パネルでも利用可能)
[設定] → [認証] で)。トークンはローカルに保存され、次のように注入されます。
OPENAI_API_KEY であり、LLM を呼び出す前に自動的に更新されます。
アダプターはデフォルトでは無効になっており、 an-api.env でボット トークンを指定すると有効になります。
電報: TELEGRAM_BOT_TOKEN を設定します。 Whisper による音声メモの文字起こしをサポートします。
Discord: DISCORD_BOT_TOKEN を設定します。
Slack : SLACK_BOT_TOKEN と SLACK_APP_TOKEN (ソケットモード) を設定します。
電子メール: EMAIL_ADDRESS 、 EMAIL_PASSWORD 、 EMAIL_IMAP_HOST 、 EMAIL_SMTP_HOST を設定します。
各チャネルはライブ タイピング インジケーターを提供し、リアルタイムのツール実行ログをストリーミングし、会話を分離します。
テスト スイートは、コントラクト、プランナー、アダプター、および ReAct ループをカバーします。
pytest # 完全なスイートを実行する
pytest testing/test_task_contract_loop.py -v # 幻覚防止契約システムをテストする
ライセンス
Distill は MIT ライセンスに基づいてリリースされています。
証拠がなければ「完了」とは言えない自己ホスト型 AI エージェント — 証拠ゲート型タスク契約、自動ロールバックを備えた自己抽出スキル、マルチチャネル。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。これをリロードしてください

ページに
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hosted AI agent that can't say "done" without proof — evidence-gated task contracts, self-distilled skills with auto-rollback, multi-channel. - Aspct3434/Distill-Agent

GitHub - Aspct3434/Distill-Agent: Self-hosted AI agent that can't say "done" without proof — evidence-gated task contracts, self-distilled skills with auto-rollback, multi-channel. · GitHub
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
Aspct3434
/
Distill-Agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
138 Commits 138 Commits .github/ workflows .github/ workflows bin bin control-panel control-panel docs docs examples examples lib lib persona persona scripts scripts skills skills src src tests tests .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md an-api.env.example an-api.env.example cli_chat.py cli_chat.py docker-compose.yml docker-compose.yml fly.toml fly.toml ollama.env.example ollama.env.example package.json package.json pyproject.toml pyproject.toml railway.json railway.json render.yaml render.yaml requirements-ml.txt requirements-ml.txt requirements.txt requirements.txt View all files Repository files navigation
The self-hosted AI agent that can't say "done" without proof.
Every Distill run is governed by a task contract : before finishing, the agent must produce physical evidence — a file written, a test passing, a service answering on its port. No evidence, no "done".
Distill also learns your projects across sessions, distills its own reusable skills from experience (each one versioned and rolled back automatically if a newer version measurably regresses), keeps persistent hybrid memory , and reaches you over Telegram, Discord, Slack, email, a real-time streaming control panel, or the terminal.
Launch the interactive terminal interface with:
graph TD
%% External Interfaces
UI[Control Panel UI <br/>React + Tailwind]
Adapters[Messaging Adapters <br/>Telegram, Discord, Slack, Email]
TUI[Terminal UI]
%% Gateway & Concurrency
subgraph Gateway Layer
API[FastAPI Gateway]
WS[WebSocket Stream]
Queue[Session FIFO Queue]
API --- WS
WS --> Queue
end
UI --> API
Adapters --> API
TUI --> WS
%% Core Agent Engine
subgraph Core Agent Engine
Agent[Agent Engine <br/>ReAct Loop]
Contract[Task Contract System <br/>Evidence Gating]
Plan[Plan Management <br/>Action Ledger]
Eval[Skill Distiller <br/>Evaluator]
Agent <--> Contract
Agent <--> Plan
Agent --> Eval
end
Queue --> Agent
%% State & Memory
subgraph State & Memory
Checkpoint[(SQLite Checkpointer <br/>State & History)]
Chroma[(ChromaDB <br/>Semantic Memory)]
Neo4j[(Neo4j <br/>Graph Memory)]
Agent <--> Checkpoint
Agent <--> Chroma
Agent <--> Neo4j
end
%% External Services
LLM((LLM Provider <br/>LiteLLM))
Agent <--> LLM
%% Tooling & Execution
subgraph Tooling & Execution Sandbox
Tools[Tool Manager]
MCP[MCP Servers]
Sandbox[Terminal Sandbox <br/>Docker / Host / Serverless]
Skills[Evolved Skills <br/>Auto-Maker]
Agent --> Tools
Tools --> MCP
Tools --> Sandbox
Tools --> Skills
Eval -.->|Synthesizes & Validates| Skills
end
Loading
Directory Structure
src/
agent.py -- ReAct loop, session management, checkpointing
contract.py -- Task-contract system (evidence tracking)
evaluator.py -- Skill distiller: trajectory -> reusable MCP skill
memory.py -- HybridMemory: ChromaDB (semantic) + Neo4j (graph)
gateway.py -- FastAPI app, WebSocket stream, session-per-FIFO-lane
tools.py -- MCP servers, terminal, file, process, port tools
control-panel/ -- React + Tailwind chat UI with live token streaming
The reasoning behind these decisions — and the alternatives that were
rejected — is documented in docs/DESIGN.md .
Task-Contract Execution : The agent must declare the required execution evidence (files created, services running) before starting a task. The final response is gated on this physical evidence, eliminating "I'll do it now" hallucinations.
Skill Distillation : After a successful complex task, an LLM evaluates the trajectory and synthesizes a parameterized Python tool. New skills are versioned, validated, and automatically rolled back if their success rate drops.
Session-per-FIFO-Lane Concurrency : Every user session gets a dedicated queue and worker task, allowing high concurrency with strict message ordering.
Hybrid Memory : Combines SQLite full-text search, ChromaDB semantic embeddings, and Neo4j graph relationships to recall cross-session context.
Universal Sandboxing : Run shell operations locally, in Docker, or remotely through a single HTTP exec shim that can front any serverless sandbox (Daytona, E2B, Modal, …).
Shareable Skills : Export and import distilled skills via the open SKILL.md format.
Distill is a research framework with a stable core and a set of more
experimental capabilities around it. This matrix shows where each subsystem
stands — treat anything marked Experimental as subject to change.
Spin up the agent gateway in the cloud, add your LLM API key (and optionally a
Telegram bot token), and text your agent. Runs lean — SQLite-backed memory, no
database to provision.
Render reads render.yaml , generates a gateway token for
you, and prompts for your provider key + models.
Railway uses railway.json : create a project from this
repo, then set AGENT_API_TOKEN , AGENT_MODEL , and your provider key.
Fly.io : fly launch --copy-config against fly.toml , then
fly secrets set … (commands are in the file header).
./scripts/quickstart.sh
Generates strong secrets, creates your env file, and brings the full stack up
with Docker Compose. Add your LLM API key to an-api.env and re-run.
# 1. Copy the environment template
cp an-api.env.example an-api.env
# 2. Add your LLM API Key to an-api.env (e.g., MOONSHOT_API_KEY or OPENAI_API_KEY)
# By default, Distill uses LiteLLM and supports OpenAI, Anthropic, Gemini, Ollama, etc.
# 3. Set a Neo4j password for the bundled graph-memory database (required;
# there is no insecure default). Compose reads it from .env:
echo " NEO4J_PASSWORD= $( openssl rand -hex 24 ) " >> .env
# 4. Start the application
docker compose up -d --build
Control Panel (UI) : http://localhost:5173
API / Docs : http://localhost:8000/docs
Empty machine (no Node, no Python)?
The bootstrap scripts install the prerequisites for you, then hand off to the
interactive installer. They are the recommended path on a fresh PC.
irm https: // raw.githubusercontent.com / Aspct3434 / Distill - Agent / master / scripts / bootstrap.ps1 | iex
macOS / Linux (bash):
curl -fsSL https://raw.githubusercontent.com/Aspct3434/Distill-Agent/master/scripts/bootstrap.sh | bash
You can also clone the repo first and run scripts/bootstrap.ps1 (Windows) or
scripts/bootstrap.sh (macOS/Linux) directly.
Run the interactive installer with npx (no global install required):
npx @aspct / distill - agent install
Prefer pinning to the exact GitHub revision? Use the repo form:
npx --yes github:Aspct3434/Distill-Agent install
After installation, use the CLI to manage the agent:
npm i -g @aspct/distill-agent
distill # Open the interactive terminal UI
distill start # Start the backend and control panel
distill logs # View running logs
distill update # Pull the latest changes
distill doctor # Diagnose the install and environment
# npx works too if you prefer not to install globally:
npx @aspct/distill-agent start # Start the backend and control panel
npx @aspct/distill-agent logs # View running logs
npx @aspct/distill-agent update # Pull the latest changes
npx @aspct/distill-agent doctor # Diagnose the install and environment
Configuration
Distill is configured via environment variables. Key settings in an-api.env :
AGENT_MODEL accepts any LiteLLM provider string. The installers can
preconfigure Kimi/Moonshot, Ollama, OpenRouter, OpenAI, Anthropic, Gemini,
DeepSeek, Groq, xAI Grok, Mistral, or any OpenAI-compatible endpoint (vLLM) —
set the matching *_API_KEY in an-api.env .
For OpenAI you can also sign in with your ChatGPT account (Codex OAuth)
instead of pasting a key: choose it in the installer, or run
PYTHONPATH=src python -m auth login (also available in the control panel
under Settings → Authentication). The token is stored locally, injected as
OPENAI_API_KEY , and refreshed automatically before LLM calls.
Adapters are disabled by default and activate when you provide a bot token in an-api.env :
Telegram : Set TELEGRAM_BOT_TOKEN . Supports voice note transcriptions via Whisper.
Discord : Set DISCORD_BOT_TOKEN .
Slack : Set SLACK_BOT_TOKEN & SLACK_APP_TOKEN (Socket Mode).
Email : Set EMAIL_ADDRESS , EMAIL_PASSWORD , EMAIL_IMAP_HOST , EMAIL_SMTP_HOST .
Each channel provides a live typing indicator, streams real-time tool execution logs, and isolates conversations.
The test suite covers contracts, planners, adapters, and the ReAct loop:
pytest # Run the full suite
pytest tests/test_task_contract_loop.py -v # Test the anti-hallucination contract system
License
Distill is released under the MIT License .
Self-hosted AI agent that can't say "done" without proof — evidence-gated task contracts, self-distilled skills with auto-rollback, multi-channel.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
