---
source: "https://github.com/akatzmann/slash-agent"
hn_url: "https://news.ycombinator.com/item?id=48602690"
title: "Show HN: slash-agent – Native LLM copilot for your terminal"
article_title: "GitHub - akatzmann/slash-agent: Native LLM Copilot for Your Terminal – **No Daemons. Zero Idle Memory. 100% On-Demand.** · GitHub"
author: "akatzmann"
captured_at: "2026-06-19T21:29:42Z"
capture_tool: "hn-digest"
hn_id: 48602690
score: 1
comments: 0
posted_at: "2026-06-19T20:07:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: slash-agent – Native LLM copilot for your terminal

- HN: [48602690](https://news.ycombinator.com/item?id=48602690)
- Source: [github.com](https://github.com/akatzmann/slash-agent)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T20:07:55Z

## Translation

タイトル: HN を表示: スラッシュエージェント – 端末のネイティブ LLM コパイロット
記事のタイトル: GitHub - akatzmann/slash-agent: 端末用のネイティブ LLM コパイロット – **デーモンなし。アイドルメモリゼロ。 100% オンデマンド。** · GitHub
説明: 端末用のネイティブ LLM コパイロット – **デーモンなし。アイドルメモリゼロ。 100% オンデマンド。** - akatzmann/slash-agent

記事本文:
GitHub - akatzmann/slash-agent: 端末用のネイティブ LLM コパイロット – **デーモンなし。アイドルメモリゼロ。 100% オンデマンド。** · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
アカッツマン
/
スラッシュエージェント
公共
通知
通知設定を変更するにはサインインする必要があります
さらに

l ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
16 コミット 16 コミット .agent .agent bin bin docs docs openspec openspec shlash_agent sinnerslash_agent .env.template .env.template .gitignore .gitignore README.md README.md要件.txt要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
⚡ スラッシュエージェント: 端末用のネイティブ LLM コパイロット
スラッシュエージェントは、アクティブな Bash シェルにネイティブに統合される、超軽量でオーバーヘッドのない AI コーディング パートナーです。コマンド ラインのシームレスな拡張機能として機能するように設計されており、100% ローカルのプライベート LLM サポート (Ollama 経由) またはクラウド パワーハウス (OpenAI/Azure OpenAI) により、ユーザーはターミナルに集中し続けることができます。
シェル内の自然なコーディング パートナー - ワークフローの中断なし。
スラッシュエージェントは、アクティブな Bash セッション内で直接動作します。使用していないときは完全に邪魔にならず、バックグラウンド リソースを消費しません (実行中のデーモンやバックグラウンド プロセスはありません)。ブロッカーに遭遇したときに /agent と入力するだけです。最近のターミナル コンテキスト (tmux ペインのスクロールバックまたはコマンド履歴) を即座に取得して、エラーを診断し、ファイルを編集し、コマンドを実行します。ディレクトリの変更 ( cd ) が自動的に同期され、終了時に環境が親シェル セッションにエクスポートされます。
⚡ クイックスタート (5 秒インストール)
すぐに起動して実行できます。シェルでクイック インストーラー スクリプトを実行します (Bash、Zsh、Ksh、Fish をサポート)。
カール -fsSL https://raw.githubusercontent.com/akatzmann/slash-agent/master/bin/installer.sh |バッシュ
(これにより、リポジトリが ~/.slash-agent に自動的に複製され、Python 仮想環境が構成され、要件がインストールされ、適切なシェル プロファイル ファイル (例: ~/.zshrc 、 ~/.bashrc 、 ~/.bash_profile 、または config.fish ) にシェル統合が登録されます)

。）
🤖 LLM Agnostic & Privacy First: OpenAI と Azure OpenAI だけでなく、キーも必要なく、マシンからデータも流出しないローカル オフライン モデル (Ollama など) をサポートします。
🔌 Zero-Overhead Integration: Completely passive. Consumes zero CPU/memory until you run /agent —no running background daemons, cron jobs, or log listeners.
🔍 コンテキスト認識診断: アクティブな tmux ペインまたは端末履歴の最後の 50 行を即座に抽出し、LLM が手動でコピーアンドペーストすることなくエラー出力とトレースバックを読み取ることができます。
⚡ 状態の同期: エージェントによって行われた作業ディレクトリの遷移 ( cd ) と環境エクスポート (export KEY=VAL ) は、終了時に親シェル セッションに自動的に同期されます。
🌉 インタラクティブ PTY ブリッジ: 提案されたコマンドを疑似端末 (PTY) で実行し、パスワード (例: sudo ) をインタラクティブに入力したり、色付きの出力を表示したり、進行状況バーを表示したりすることができます。
🕹️ Steerable Confirmation Loop: Full control over every action:
n (いいえ): コマンドを拒否し、エージェントに通知します。
e (edit): Inline edit the command before running it.
c (comment): Type natural language guidance back to the agent (e.g. "Use yarn instead of npm" ).
🛡️ Dry-run & Auto-confirm Modes: Preview agent actions safely with -n / --dry-run , or run fully unattended with -y / --yes .
$ npm ビルドを実行
❌ エラー: ビルドに失敗しました。モジュール「dotenv」がserver.jsに見つかりません:12
$ /agent これを修正する
[Agent Shell] Initializing with model 'gemma4:e4b-it-qat' at 'http://127.0.0.1:11434'...
[エージェントが開始したタスク]
Analyzing terminal context... Identified missing dependency 'dotenv' in server.js.
[エージェント] 提案されたコマンド:
$ npm install dotenv && npm run build
Confirm action: [y]es / [n]o / [e]dit / [c]omment ? y
[Agent Running]: npm install dotenv && npm run build
...
1 つのパッケージと aud を追加

1秒で120個のパッケージを実行
✓ ビルドが正常に完了しました。
不足している「dotenv」パッケージをインストールし、ビルドが成功することを確認しました。
🎯 一般的な使用例
🛠️ ビルドとテストのクラッシュ: コンパイラ エラー、スクリプト トレースバック、または単体テストが失敗した場合は、/agent を実行するだけでエラー ログを直接読み取り、修正を提案できます。
📦 依存関係の解決: パッケージのインポートが見つかりませんか?エージェントはインポート エラーを読み取り、パッケージをインストールし、ビルドを検証します。
💻 クイック スクリプトと自動化: ヘルパー スクリプトの生成、開発環境の構成、またはオンザフライでの正規表現ログ処理の実行をエージェントに依頼します。
⚙️ システム構成: コマンド フラグを検索せずに、ローカル データベース、systemd サービス、または構成ファイルを簡単にセットアップします。
分離されたサブシェルで実行される (現在のディレクトリや環境を変更できない) 標準エージェントとは異なり、slash-agent は軽量の状態同期プロトコルを使用します。
グラフTD
A[アクティブな Bash セッション] -->|1。 /agent| と入力しますB(Bash Sourcing Wrapper bin/slash-agent.sh)
B -->|2.画面出力のキャプチャ| C(Python オーケストレーターslash_agent/main.py)
C -->|3.ローカル/クラウド LLM とのインターフェース| C
C -->|4.コマンドを提案する| D(PTY実行ブリッジ)
D -->|5。インタラクティブ ステアリング ループ|あ
D -->|6.終了コード、PWD、環境更新をキャプチャ| B
B -->|7.終了時のソース同期ファイル|あ
読み込み中
コンテキスト キャプチャ: シェル ラッパーは、アクティブな tmux ペイン バッファ (または履歴) を自動的にキャプチャし、LLM に即時のコンテキストを提供します。
インタラクティブな PTY ブリッジ: コマンドは実際の疑似端末 (PTY) 内で実行されるため、色付きの出力や進行状況バーが表示され、プロンプトと対話することができます ( sudo のパスワードの入力など)。
親シェルの同期: ディレクトリの変更 ( cd ) または環境変数 (export ) は、一時的なソース経由でメイン シェル セッションに安全に戻されます。

終了時のスクリプト。
クイック スタート スクリプトを使用する代わりにエージェントを手動でセットアップする場合は、次の手順を実行します。
リポジトリのクローンを作成します。
git clone https://github.com/akatzmann/slash-agent.git ~ /.slash-agent
cd ~ /.slash-agent
Python のインストール要件:
pip install -r 要件.txt
シェル統合を登録します。
適切なソーシング ステートメントをシェル構成ファイルに追加します。
Bash (Linux): ~/.bashrc 内のソース ~/.slash-agent/bin/slash-agent.sh
Bash (macOS): ~/.bash_profile (または ~/.profile ) 内のソース ~/.slash-agent/bin/slash-agent.sh
Zsh: ソース ~/.slash-agent/bin/slash-agent.sh (~/.zshrc)
Ksh: ソース ~/.slash-agent/bin/slash-agent.sh (~/.kshrc)
Fish: ソース ~/.config/fish/config.fish 内の ~/.slash-agent/bin/slash-agent.fish
スラッシュエージェントは Unix のような PTY 環境内でネイティブに実行されます。 PTY エミュレーションの制限のため、ネイティブ Windows 実行 (標準 CMD または PowerShell の下) はサポートされていません。
ただし、このツールは WSL2 (Windows Subsystem for Linux) と 100% 互換性があります。 Windows ユーザーは、WSL2 Linux ターミナル (Ubuntu や Debian など) を開いて標準のクイック スタート インストール コマンドを実行することで、スラッシュ エージェントを実行できます。
.env ファイルまたはシェル プロファイルで LLM バックエンド、エンドポイント、モデル、およびキャプチャ設定を構成します。
# LLM バックエンド: openai (デフォルト)、ollama、azure_openai、ダミー
エクスポート AGENT_BACKEND= " openai "
# モデル名 (デフォルト: openai の場合は gpt-4o-mini、ollam の場合は gemma4:e4b-it-qat)
エクスポート AGENT_MODEL= " gpt-4o-mini "
# API エンドポイントのベース URL (デフォルトは公式の OpenAI API エンドポイント)
エクスポート AGENT_ENDPOINT= " "
# OpenAI API キー (デフォルトの OpenAI バックエンドに必要)
import OPENAI_API_KEY= " your-api-key-here "
# コンテキスト抽出設定
import AGENT_TMUX_LINES=50 # アクティブな tmux スクロールバックからキャプチャされた行
export AGENT_HISTORY_COMMANDS=20 # h からキャプチャされたコマンド

ストーリーフォールバック
構成変数 (Azure OpenAI 変数など) の完全なリストについては、.env.template ファイルを参照してください。
/agent 'sandbox' という名前の新しいディレクトリを作成し、その中に基本的な Python フラスコ サーバーを書き込みます
2. 衝突後の診断
コンパイラ、ビルド ツール、またはスクリプトがクラッシュした場合は、引数を指定せずに /agent を実行します (または修正リクエストを指定します)。
/agent このクラッシュを修正する
3. 予行運転モード
実際のシステム変更を行わずに、提案された手順をシミュレートし、エージェントの計画を確認します。
/agent -n PostgreSQL および Redis 用の docker compose ファイルをセットアップします
4. 自動確認モード
安全、低リスク、または中リスクのコマンドについては、確認プロンプトを表示せずにタスクを実行します。
/agent -y パッケージリストとインストールツリーを更新します
5. 重要なコマンドの自動確認
デフォルトでは、重要なコマンド (rm -rf や sudo を使用するコマンドなど) は、偶発的な損傷を防ぐために -y によって自動確認されません。重要なコマンドでも自動確認するには、 --unsafe-yes フラグを渡します。
/agent --unsafe-yes Docker ボリュームとシステム キャッシュをクリーンアップします
📘 詳細
アーキテクチャ、対話型 PTY ブリッジ ループ、および環境状態同期プロトコルの技術的な詳細については、「技術ドキュメント」を参照してください。
端末用のネイティブ LLM コパイロット – **デーモンなし。アイドルメモリゼロ。 100% オンデマンド。**
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Native LLM Copilot for Your Terminal – **No Daemons. Zero Idle Memory. 100% On-Demand.** - akatzmann/slash-agent

GitHub - akatzmann/slash-agent: Native LLM Copilot for Your Terminal – **No Daemons. Zero Idle Memory. 100% On-Demand.** · GitHub
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
akatzmann
/
slash-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits .agent .agent bin bin docs docs openspec openspec slash_agent slash_agent .env.template .env.template .gitignore .gitignore README.md README.md requirements.txt requirements.txt View all files Repository files navigation
⚡ slash-agent: Native LLM Copilot for Your Terminal
slash-agent is an ultra-lightweight, zero-overhead AI coding partner that integrates natively into your active Bash shell. It is designed to act as a seamless extension of your command line, keeping you focused in your terminal with 100% local, private LLM support (via Ollama) or cloud powerhouses (OpenAI/Azure OpenAI).
A Natural Coding Partner in Your Shell—Zero Workflow Interruption.
slash-agent operates directly inside your active Bash session. It stays completely out of your way and consumes zero background resources (no running daemons, no background processes) when not in use. Simply type /agent when you hit a blocker: it instantly grabs your recent terminal context (tmux pane scrollback or command history) to diagnose errors, edit files, and execute commands—automatically syncing directory changes ( cd ) and environment exports back to your parent shell session when it exits.
⚡ Quick Start (5-Second Install)
Get up and running instantly. Run the quick installer script in your shell (supports Bash, Zsh, Ksh, and Fish):
curl -fsSL https://raw.githubusercontent.com/akatzmann/slash-agent/master/bin/installer.sh | bash
(This automatically clones the repo to ~/.slash-agent , configures a Python virtual environment, installs requirements, and registers the shell integration in your appropriate shell profile file, e.g. ~/.zshrc , ~/.bashrc , ~/.bash_profile , or config.fish .)
🤖 LLM Agnostic & Privacy First: Supports local offline models (like Ollama) with zero keys required and zero data leaving your machine, as well as OpenAI and Azure OpenAI.
🔌 Zero-Overhead Integration: Completely passive. Consumes zero CPU/memory until you run /agent —no running background daemons, cron jobs, or log listeners.
🔍 Context-Aware Diagnoses: Instantly extracts the last 50 lines of your active tmux pane or terminal history, letting the LLM read error outputs and tracebacks without manual copy-pasting.
⚡ State Synchronization: Working directory transitions ( cd ) and environment exports ( export KEY=VAL ) made by the agent automatically sync back to your parent shell session on exit.
🌉 Interactive PTY Bridge: Executes proposed commands in a pseudo-terminal (PTY), allowing you to interactively type passwords (e.g. sudo ), view colored output, and see progress bars.
🕹️ Steerable Confirmation Loop: Full control over every action:
n (no): Refuse the command and inform the agent.
e (edit): Inline edit the command before running it.
c (comment): Type natural language guidance back to the agent (e.g. "Use yarn instead of npm" ).
🛡️ Dry-run & Auto-confirm Modes: Preview agent actions safely with -n / --dry-run , or run fully unattended with -y / --yes .
$ npm run build
❌ ERROR: Build failed. Cannot find module 'dotenv' in server.js:12
$ /agent Fix this
[Agent Shell] Initializing with model 'gemma4:e4b-it-qat' at 'http://127.0.0.1:11434'...
[Agent Started Task]
Analyzing terminal context... Identified missing dependency 'dotenv' in server.js.
[Agent] Proposed Command:
$ npm install dotenv && npm run build
Confirm action: [y]es / [n]o / [e]dit / [c]omment ? y
[Agent Running]: npm install dotenv && npm run build
...
added 1 package, and audited 120 packages in 1s
✓ Build completed successfully!
I have installed the missing 'dotenv' package and verified that the build now passes.
🎯 Common Use Cases
🛠️ Build & Test Crashes: When a compiler error, script traceback, or unit test fails, simply run /agent to let it read the error logs directly and propose a fix.
📦 Dependency Resolution: Missing package imports? The agent reads the import error, installs the package, and verifies the build.
💻 Quick Scripting & Automation: Ask the agent to generate helper scripts, configure development environments, or perform regex logs processing on the fly.
⚙️ System Configuration: Easily set up local databases, systemd services, or configuration files without looking up command flags.
Unlike standard agents that run in isolated subshells (and cannot modify your current directory or environment), slash-agent uses a lightweight state synchronization protocol:
graph TD
A[Active Bash Session] -->|1. Type /agent| B(Bash Sourcing Wrapper bin/slash-agent.sh)
B -->|2. Capture Screen Output| C(Python Orchestrator slash_agent/main.py)
C -->|3. Interface with local/cloud LLM| C
C -->|4. Propose commands| D(PTY Execution Bridge)
D -->|5. Interactive Steering Loop| A
D -->|6. Capture exit code, PWD, env updates| B
B -->|7. Source sync file on exit| A
Loading
Context Capture: The shell wrapper automatically captures the active tmux pane buffer (or history) to give the LLM immediate context.
Interactive PTY Bridge: Commands run inside a real pseudo-terminal (PTY) so you see colored outputs, progress bars, and can interact with prompts (like typing passwords for sudo ).
Parent Shell Sync: Directory changes ( cd ) or environment variables ( export ) are safely passed back to your main shell session via a temporary sourcing script on exit.
If you prefer to set up the agent manually instead of using the Quick Start script:
Clone the Repository:
git clone https://github.com/akatzmann/slash-agent.git ~ /.slash-agent
cd ~ /.slash-agent
Install Python Requirements:
pip install -r requirements.txt
Register Shell Integration:
Add the appropriate sourcing statement to your shell configuration file:
Bash (Linux): source ~/.slash-agent/bin/slash-agent.sh in ~/.bashrc
Bash (macOS): source ~/.slash-agent/bin/slash-agent.sh in ~/.bash_profile (or ~/.profile )
Zsh: source ~/.slash-agent/bin/slash-agent.sh in ~/.zshrc
Ksh: source ~/.slash-agent/bin/slash-agent.sh in ~/.kshrc
Fish: source ~/.slash-agent/bin/slash-agent.fish in ~/.config/fish/config.fish
slash-agent runs natively inside Unix-like PTY environments. Native Windows execution (under standard CMD or PowerShell ) is not supported due to PTY emulation limitations.
However, the tool is 100% compatible with WSL2 (Windows Subsystem for Linux) . Windows users can run slash-agent by opening any WSL2 Linux terminal (such as Ubuntu or Debian) and running the standard Quick Start installation command.
Configure the LLM backend, endpoint, model, and capture settings in your .env file or shell profile:
# LLM Backend: openai (default), ollama, azure_openai, dummy
export AGENT_BACKEND= " openai "
# Model name (Defaults: gpt-4o-mini for openai, gemma4:e4b-it-qat for ollama)
export AGENT_MODEL= " gpt-4o-mini "
# API endpoint base URL (defaults to official OpenAI API endpoint)
export AGENT_ENDPOINT= " "
# OpenAI API Key (required for default OpenAI backend)
export OPENAI_API_KEY= " your-api-key-here "
# Context extraction settings
export AGENT_TMUX_LINES=50 # Lines captured from active tmux scrollback
export AGENT_HISTORY_COMMANDS=20 # Commands captured from history fallback
For a full list of configuration variables (e.g., Azure OpenAI variables), see the .env.template file.
/agent create a new directory named ' sandbox ' and write a basic python flask server inside it
2. Post-Crash Diagnosis
If a compiler, build tool, or script crashes, run /agent with no arguments (or a request to fix it):
/agent Fix this crash
3. Dry-run Mode
Simulate proposed steps and check the agent's plan without making actual system changes:
/agent -n setup a docker compose file for PostgreSQL and Redis
4. Auto-confirm Mode
Run tasks without any confirmation prompts for safe, low, or moderate risk commands:
/agent -y update package lists and install tree
5. Auto-confirm Critical Commands
By default, critical commands (like rm -rf or commands using sudo ) are not auto-confirmed by -y to prevent accidental damage. To auto-confirm even critical commands, pass the --unsafe-yes flag:
/agent --unsafe-yes clean up docker volumes and system cache
📘 Deep Dive
For more technical details on the architecture, the interactive PTY bridge loop, and the environment state-synchronization protocol, read the Technical Documentation .
Native LLM Copilot for Your Terminal – **No Daemons. Zero Idle Memory. 100% On-Demand.**
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
