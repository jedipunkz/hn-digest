---
source: "https://github.com/provos/ironcurtain"
hn_url: "https://news.ycombinator.com/item?id=48902269"
title: "IronCurtain – A secure* runtime for autonomous AI agents"
article_title: "GitHub - provos/ironcurtain: A secure* runtime for autonomous AI agents. Policy from plain-English constitutions. (*https://ironcurtain.dev) · GitHub"
author: "n0on3"
captured_at: "2026-07-14T04:41:40Z"
capture_tool: "hn-digest"
hn_id: 48902269
score: 1
comments: 0
posted_at: "2026-07-14T04:22:11Z"
tags:
  - hacker-news
  - translated
---

# IronCurtain – A secure* runtime for autonomous AI agents

- HN: [48902269](https://news.ycombinator.com/item?id=48902269)
- Source: [github.com](https://github.com/provos/ironcurtain)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T04:22:11Z

## Translation

タイトル: IronCurtain – 自律型 AI エージェントのための安全な * ランタイム
記事のタイトル: GitHub - provos/ironcurtain: 自律型 AI エージェントのための安全な* ランタイム。平易な英語の憲法からの政策。 (*https://ironcurtain.dev) · GitHub
説明: 自律型 AI エージェント用の安全な* ランタイム。平易な英語の憲法からの政策。 (*https://ironcurtain.dev) - provos/ironcurtain

記事本文:
GitHub - provos/ironcurtain: 自律型 AI エージェント用の安全な* ランタイム。平易な英語の憲法からの政策。 (*https://ironcurtain.dev) · GitHub
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
プロボス島
/
鉄のカーテン
公共
通知
通知セットを変更するにはサインインする必要があります

ひずみ
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
544 コミット 544 コミット .claude .claude .github .github .hooks .hooks docker docker docs docs 例 例 パッケージ パッケージ スクリプト スクリプト src src test test .env.example .env.example .gitignore .gitignore .madgerc .madgerc .npmignore .npmignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json .semgrepignore .semgrepignore ADDING_MCP_SERVERS.md ADDING_MCP_SERVERS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONFIG.md CONFIG.md COTRIBUTING.md COTRIBUTING.md DAEMON.md DAEMON.md DEVELOPER_GUIDE.md DEVELOPER_GUIDE.md GEMINI.md GEMINI.md ライセンス ライセンス MODEL_ROUTING.md MODEL_ROUTING.md README.md README.md RUNNING_MODES.md RUNNING_MODES.md SANDBOXING.md SANDBOXING.md SECURITY.md SECURITY.md TESTING.md TESTING.md TODO.md TODO.md TRAJECTORIES.md TRAJECTORIES.md TRANSPORT.md TRANSPORT.md WORKFLOWS.md WORKFLOWS.md デモ.gif デモ.gif eslint.config.js eslint.config.js flake.lock flake.lock flake.nix flake.nix package-lock.json package-lock.json package.json package.json tsconfig.eslint.json tsconfig.eslint.json tsconfig.json tsconfig.json tsconfig.scripts.json tsconfig.scripts.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自律型 AI エージェント用のセキュア* ランタイム。セキュリティ ポリシーは人間が判読できる構成から派生します。
※誰かが「安全」と書いたら、すぐに疑うべきです。安全とはどういう意味ですか?
研究のプロトタイプ。 IronCurtain は、AI エージェントを本当に役立つのに十分安全にする方法を探る初期段階の研究プロジェクトです。 API、構成形式、アーキテクチャは変更される可能性があります。貢献とフィードバックは大歓迎です

おめ。
エージェントは、リポジトリのクローンを作成し、変更をプッシュするように求められます。 git_clone と git_push は両方ともポリシー エンジンによってエスカレーションされますが、自動承認者はそれらを自動的に承認します。コマンド モード (Ctrl-A) からのユーザーの信頼できる入力によって明確な意図が提供されるため、手動の /approve は必要ありませんでした。
自律型 AI エージェントは、ユーザーに代わってファイルの管理、git コマンドの実行、メッセージの送信、API との対話を行うことができます。しかし、今日のエージェント フレームワークでは、ファイル システム、資格情報、ネットワークへのフル アクセスなど、ユーザーと同じ権限がエージェントに与えられます。セキュリティ研究者はこれを アンビエント権限 と呼び、単一のプロンプト インジェクションまたは複数ターンのドリフトにより、エージェントがファイルを削除したり、データを窃取したり、悪意のあるコードをプッシュしたりする可能性があることを意味します。
一般的な対応は、エージェントを狭いサンドボックスに制限する (エージェントの有用性を制限する) か、ユーザーにすべてのアクションの承認を求める (エージェントの自律性を制限する) かのいずれかです。どちらも満足できるものではありません。
IronCurtain は別の方法を採用します。つまり、セキュリティの意図を平易な英語で表現し、システムに適用を判断させます。
あなたは、あなたの代理人が何をすることができ、何ができないのかを説明する短い文書である規約を作成します。 IronCurtain は、LLM パイプラインを使用してこれを決定論的なセキュリティ ポリシーにコンパイルし、生成されたテスト シナリオに対してコンパイルされたルールを検証して、ツール呼び出しごとに実行時にポリシーを適用します。その結果、自然言語で定義した境界内で自律的に動作できるエージェントが得られます。
エージェントは信頼されていません。 IronCurtain は、プロンプト インジェクションまたはドリフトによって LLM が侵害される可能性があると想定しています。セキュリティはモデルが「良い」かどうかに依存しません。
英語が入って、施行がアウト。インテント (「承認なしの破壊的な Git 操作は禁止」) を記述します。システムはそれを決定論的なルールにコンパイルし、追加の LLM なしで適用されます。

実行時の関与。
意味論的な介入。エージェントに生のシステム アクセスを与える代わりに、すべての対話は MCP サーバー (ファイル システム、git など) を経由します。すべてのツール呼び出しは、許可、拒否、または承認のためにユーザーにエスカレーションできるポリシー エンジンを通過します。
多層防御。エージェント コードは、ホストに直接アクセスせずに V8 分離で実行されます。唯一の方法は、意味的に意味のある MCP ツール呼び出しを使用し、すべての呼び出しがポリシーに対してチェックされることです。
IronCurtain は、異なる信頼モデルを持つ 2 つのセッション モードをサポートします。
組み込みエージェント (コード モード) — IronCurtain 独自の LLM エージェントは、V8 サンドボックスで実行される TypeScript スニペットを作成します。 IronCurtain は、エージェント、サンドボックス、およびポリシー エンジンを制御します。すべてのツール呼び出しは、構造化された MCP リクエストとしてサンドボックスを出て、ポリシー エンジン (許可/拒否/エスカレート) を通過してから初めて実際の MCP サーバーに到達します。
Docker エージェント モード — 外部エージェント (Claude Code、Goose など) は、ネットワーク アクセスなしで Docker コンテナ内で実行されます。 IronCurtain は外部効果を仲介します。LLM API 呼び出しは TLS 終端 MITM プロキシ (ホスト許可リスト、偽鍵と実鍵の交換) を通過し、MCP ツール呼び出しは同じポリシー エンジンを通過し、パッケージ インストール (npm/PyPI) は検証レジストリ プロキシを通過します。
どちらのモードでも、エージェントは信頼されません。セキュリティは、指示に従うモデルには依存しません。セキュリティは境界で強制されます。
図を含む完全なアーキテクチャ、レイヤーごとの信頼分析、macOS プラットフォームのメモについては、SANDBOXING.md を参照してください。
Node.js 22、24、または 26 — IronCurtain がテストする偶数番号の主要な行 (孤立した vm で必要。24 と 26 は事前に構築されたバイナリをインストールします。Node 22 はインストール時にソースからコンパイルされ、C/C++ ツールチェーンが必要です)。奇数番号の行 (23、25) は実行されますが、テストされていません — Ironcurtain doc

トールは警告する。
Docker — 必須ではありませんが、最も強力な分離を提供する Docker エージェント モードを強く推奨します。 macOS 26 以降 (Apple シリコン) では、Apple コンテナが代替バックエンドとして機能します (コンテナごとの VM。サービスの実行中に自動的に使用されます。 Ironcurtain 設定のcontainerRuntime を参照してください)。
少なくとも 1 つの LLM プロバイダー (Anthropic、Google、または OpenAI) の API キー
グローバル CLI ツール (エンド ユーザー) として:
npm install -g @provos/ironcurtain
ソース (開発) から:
git clone https://github.com/provos/ironcurtain.git
CDの鉄カーテン
npmインストール
ワンタイムセットアップ
エクスポート ANTHROPIC_API_KEY=sk-ant-...
また、キーをプロジェクト ルートの .env ファイルに配置したり ( dotenv 経由で自動的にロードされる)、または Ironcurtain config 経由で ~/.ironcurtain/config.json に追加したりすることもできます。環境変数は、構成ファイルの値よりも優先されます。サポートされている: ANTHROPIC_API_KEY 、 GOOGLE_GENERATIVE_AI_API_KEY 、 OPENAI_API_KEY 。
2. 初回起動ウィザードを実行します (推奨するマルチプレクサ パスを使用する前にこれを明示的に実行します。また、最初の非マルチプレクサの Ironcurtain start 時にも自動的に実行されます)。
アイアンカーテンのセットアップ
GitHub トークンのセットアップ、Web 検索プロバイダー、モデルの選択、その他の設定について説明します。選択した内容で ~/.ironcurtain/config.json を作成します。
IronCurtain には、開発者エクスペリエンスを対象としたデフォルトのポリシーが同梱されています。読み取り専用操作が許可され、変更 (書き込み、プッシュ、PR 作成) は人間の承認を求めてエスカレーションされます。セットアップ後すぐに使い始めることができます。
ターミナルマルチプレクサ (推奨)
アイアンカーテンのおすすめの使い方。これにより、エージェントの対話型 TUI (Claude Code または Goose) の全機能が提供され、IronCurtain はポリシー エンジンを通じてすべてのツール呼び出しを仲介します (すべてが 1 つの端末で行われます)。
アイアンカーテンマルチプレクサ
主な機能:
フルエージェント TUI — エージェント ru

ネットワークにアクセスできない Docker コンテナ内の PTY 内の ns。ローカルで実行しているかのように操作できます。
インライン エスカレーション処理 — ツール呼び出しに承認が必要な場合、エスカレーション ピッカーが単一キー アクション (承認/拒否/ホワイトリストの a/d/w) でビューポートをオーバーレイします。 /approve+ N を使用して、残りのセッションのドメインまたはパスをホワイトリストに登録します。
信頼できるユーザー入力 — コマンド モード (Ctrl-A) で入力されたテキストは、コンテナーに入る前にホスト側でキャプチャされます。これにより、自動承認者が使用できる検証済みのインテント シグナルが作成されます。たとえば、「変更をオリジンにプッシュ」と入力すると、後続の git_push エスカレーションが自動承認されます。
タブ管理 — 複数の同時セッションを生成し ( /new )、セッション間を切り替え ( /tab N 、 Alt-1..9)、閉じます ( /close )。複数の mux インスタンスを並行して実行できます。
完全なウォークスルー: 入力モード、信頼できる入力セキュリティ モデル、エスカレーション ワークフロー、およびキーボード リファレンスについては、DEVELOPER_GUIDE.md を参照してください。
Ironcurtain start は、迅速なワンショット タスク、スクリプト、または明示的にローカルの組み込みエージェントが必要な場合に使用します。通常の対話型の Docker エージェント作業には、 Ironcurtain mux を使用します。
Ironcurtain start " ./src 内のファイルを要約します" # シングルショットモード
Ironcurtain start -w ./my-project " テストを修正する " # シングルショット ワークスペース モード
Ironcurtain start --agentbuiltin # ローカル組み込み REPL、Docker なし
Ironcurtain start --persona my-assistant " Check my email " # ペルソナを使用する
その他の走行モード
IronCurtain は、セッション再開 ( --resume <session-id> )、レガシー raw PTY/デバッグ モード、モバイル承認用のシグナル メッセージング トランスポート、スケジュールされた cron ジョブ用のデーモン モードもサポートしています。デーモンには、ブラウザベースの監視とエスカレーション処理のためのオプションの Web UI ( --web-ui ) があります。詳細については、RUNNING_MODES.md を参照してください。
IronCurtain が複数のマルチをオーケストレーションします

構造化されたワークフローを通じて iple AI エージェントをサポートします。バンドルされた脆弱性発見ワークフローは、libFuzzer/AFL++ カバレッジ ゲーティング、仮説に基づく検出/トリアージ ステート、および人間による最終的なレポート レビュー ゲートを備えた階層型ハーネス パイプライン (階層 1 の分離機能 → 階層 2 マルチコンポーネント → 階層 3 のフル ビルド) を通じて、ネイティブ コードのメモリ安全性とロジックのバグを追跡します。デザインとコードのワークフローでは、計画、設計、実装、レビューのサイクルが実行され、人間のゲートも使用されます。各エージェントは、ロール固有のポリシー境界を持つ独自の Docker コンテナ内で実行されます。エンジンは、状態遷移、アーティファクトの受け渡し、クラッシュ再開チェックポイントを自動的に管理します。オープンソースで、マシン上で完全に実行され、構成ベースのポリシー エンジンを介してエージェントごとのセキュリティ ポリシーを適用し、Docker コンテナ化されたエージェントと連携します。これは、コーディング タスクの範囲では Amazon Kiro や Google Jules に匹敵しますが、最上級のセキュリティと拡張可能なワークフロー定義形式を備えています。
Web UI は、ワークフロー実行用のインターフェイスです。デーモンを起動し、印刷された URL を開き、[ワークフロー] ページから実行を開始します。上のステートマシン グラフはライブであり、エージェント メッセージのタイムラインはマークダウン レンダリングでストリームされ、ゲート レビューにはワークスペースとアーティファクが含まれます。

[切り捨てられた]

## Original Extract

A secure* runtime for autonomous AI agents. Policy from plain-English constitutions. (*https://ironcurtain.dev) - provos/ironcurtain

GitHub - provos/ironcurtain: A secure* runtime for autonomous AI agents. Policy from plain-English constitutions. (*https://ironcurtain.dev) · GitHub
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
provos
/
ironcurtain
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
544 Commits 544 Commits .claude .claude .github .github .hooks .hooks docker docker docs docs examples examples packages packages scripts scripts src src test test .env.example .env.example .gitignore .gitignore .madgerc .madgerc .npmignore .npmignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json .semgrepignore .semgrepignore ADDING_MCP_SERVERS.md ADDING_MCP_SERVERS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONFIG.md CONFIG.md CONTRIBUTING.md CONTRIBUTING.md DAEMON.md DAEMON.md DEVELOPER_GUIDE.md DEVELOPER_GUIDE.md GEMINI.md GEMINI.md LICENSE LICENSE MODEL_ROUTING.md MODEL_ROUTING.md README.md README.md RUNNING_MODES.md RUNNING_MODES.md SANDBOXING.md SANDBOXING.md SECURITY.md SECURITY.md TESTING.md TESTING.md TODO.md TODO.md TRAJECTORIES.md TRAJECTORIES.md TRANSPORT.md TRANSPORT.md WORKFLOWS.md WORKFLOWS.md demo.gif demo.gif eslint.config.js eslint.config.js flake.lock flake.lock flake.nix flake.nix package-lock.json package-lock.json package.json package.json tsconfig.eslint.json tsconfig.eslint.json tsconfig.json tsconfig.json tsconfig.scripts.json tsconfig.scripts.json vitest.config.ts vitest.config.ts View all files Repository files navigation
A secure* runtime for autonomous AI agents, where security policy is derived from a human-readable constitution.
*When someone writes "secure," you should immediately be skeptical. What do we mean by secure?
Research Prototype. IronCurtain is an early-stage research project exploring how to make AI agents safe enough to be genuinely useful. APIs, configuration formats, and architecture may change. Contributions and feedback are welcome.
The agent is asked to clone a repository and push changes. Both git_clone and git_push are escalated by the policy engine, but the auto-approver approves them automatically — the user's trusted input from command mode (Ctrl-A) provided clear intent, so no manual /approve was needed.
Autonomous AI agents can manage files, run git commands, send messages, and interact with APIs on your behalf. But today's agent frameworks give the agent the same privileges as the user such as full access to the filesystem, credentials, and network. Security researchers call this ambient authority , and it means a single prompt injection or multi-turn drift can cause an agent to delete files, exfiltrate data, or push malicious code.
The common response is to either restrict agents to a narrow sandbox (limiting their usefulness) or to ask the user to approve every action (limiting their autonomy). Neither is satisfactory.
IronCurtain takes a different path: express your security intent in plain English, then let the system figure out enforcement.
You write a constitution which is a short document describing what your agent is and isn't allowed to do. IronCurtain compiles this into a deterministic security policy using an LLM pipeline, validates the compiled rules against generated test scenarios, and then enforces the policy at runtime on every tool call. The result is an agent that can work autonomously within boundaries you define in natural language.
The agent is untrusted. IronCurtain assumes the LLM may be compromised by prompt injection or drift. Security does not depend on the model "being good."
English in, enforcement out. You write intent ("no destructive git operations without approval"); the system compiles it into deterministic rules that are enforced without further LLM involvement at runtime.
Semantic interposition. Instead of giving the agent raw system access, all interactions go through MCP servers (filesystem, git, etc.). Every tool call passes through a policy engine that can allow , deny , or escalate to the user for approval.
Defense in depth. Agent code runs in a V8 isolate with no direct access to the host. The only way out is through semantically meaningful MCP tool calls and every one is checked against policy.
IronCurtain supports two session modes with different trust models:
Builtin Agent (Code Mode) — IronCurtain's own LLM agent writes TypeScript snippets that execute in a V8 sandbox. IronCurtain controls the agent, the sandbox, and the policy engine. Every tool call exits the sandbox as a structured MCP request, passes through the policy engine (allow / deny / escalate), and only then reaches the real MCP server.
Docker Agent Mode — An external agent (Claude Code, Goose, etc.) runs inside a Docker container with no network access. IronCurtain mediates the external effects: LLM API calls pass through a TLS-terminating MITM proxy (host allowlist, fake-to-real key swap), MCP tool calls pass through the same policy engine, and package installations (npm/PyPI) go through a validating registry proxy.
In both modes, the agent is untrusted . Security does not depend on the model following instructions — it is enforced at the boundary.
See SANDBOXING.md for the full architecture with diagrams, layer-by-layer trust analysis, and macOS platform notes.
Node.js 22, 24, or 26 — the even-numbered major lines IronCurtain tests (required by isolated-vm ; 24 and 26 install prebuilt binaries, Node 22 compiles from source at install and needs a C/C++ toolchain). Odd-numbered lines (23, 25) run but are untested — ironcurtain doctor warns.
Docker — not required but strongly recommended for Docker Agent Mode, which provides the strongest isolation. On macOS 26+ (Apple silicon), Apple container works as an alternative backend (VM per container; used automatically when its services are running — see containerRuntime in ironcurtain config )
An API key for at least one LLM provider (Anthropic, Google, or OpenAI)
As a global CLI tool (end users):
npm install -g @provos/ironcurtain
From source (development):
git clone https://github.com/provos/ironcurtain.git
cd ironcurtain
npm install
One-time setup
export ANTHROPIC_API_KEY=sk-ant-...
You can also place keys in a .env file in the project root (loaded automatically via dotenv ), or add them to ~/.ironcurtain/config.json via ironcurtain config . Environment variables take precedence over config file values. Supported: ANTHROPIC_API_KEY , GOOGLE_GENERATIVE_AI_API_KEY , OPENAI_API_KEY .
2. Run the first-start wizard (run this explicitly before using the recommended mux path; it also runs automatically on first non-mux ironcurtain start ):
ironcurtain setup
Walks you through GitHub token setup, web search provider, model selection, and other settings. Creates ~/.ironcurtain/config.json with your choices.
IronCurtain ships with a default policy geared towards the developer experience — read-only operations are allowed, mutations (writes, pushes, PR creation) escalate for human approval. You can start using it immediately after setup.
Terminal multiplexer (recommended)
The recommended way to use IronCurtain. It gives you the full power of your agent's interactive TUI (Claude Code or Goose) while IronCurtain mediates every tool call through its policy engine — all in a single terminal.
ironcurtain mux
Key capabilities:
Full agent TUI — The agent runs in a PTY inside a Docker container with no network access. You interact with it exactly as if it were running locally.
Inline escalation handling — When a tool call needs approval, an escalation picker overlays the viewport with single-key actions (a/d/w for approve/deny/whitelist). Use /approve+ N to whitelist a domain or path for the rest of the session.
Trusted user input — Text typed in command mode (Ctrl-A) is captured on the host side before entering the container. This creates a verified intent signal that the auto-approver can use — e.g., typing "push my changes to origin" will auto-approve a subsequent git_push escalation.
Tab management — Spawn multiple concurrent sessions ( /new ), switch between them ( /tab N , Alt-1..9), close them ( /close ). Multiple mux instances can run in parallel.
See DEVELOPER_GUIDE.md for the full walkthrough: input modes, trusted input security model, escalation workflow, and keyboard reference.
Use ironcurtain start for quick one-shot tasks, scripts, or when you explicitly want the local builtin agent. For normal interactive Docker-agent work, use ironcurtain mux .
ironcurtain start " Summarize the files in ./src " # Single-shot mode
ironcurtain start -w ./my-project " Fix the tests " # Single-shot workspace mode
ironcurtain start --agent builtin # Local builtin REPL, no Docker
ironcurtain start --persona my-assistant " Check my email " # Use a persona
Other running modes
IronCurtain also supports session resume ( --resume <session-id> ), a legacy raw PTY/debug mode, a Signal messaging transport for mobile approval, and a daemon mode for scheduled cron jobs. The daemon has an optional web UI ( --web-ui ) for browser-based monitoring and escalation handling. See RUNNING_MODES.md for details.
IronCurtain orchestrates multiple AI agents through structured workflows. The bundled vulnerability discovery workflow hunts memory-safety and logic bugs in native code through a tiered harness pipeline (Tier 1 isolated function → Tier 2 multi-component → Tier 3 full build) with libFuzzer/AFL++ coverage gating, hypothesis-driven discover / triage states, and a final human report-review gate. The design-and-code workflow runs plan / design / implement / review cycles, also with human gates. Each agent runs in its own Docker container with role-specific policy boundaries; the engine manages state transitions, artifact passing, and crash-resume checkpointing automatically. Open source, runs entirely on your machine, enforces per-agent security policies via the constitution-based policy engine, and works with any Docker-containerized agent — comparable in scope to Amazon Kiro and Google Jules for coding tasks, but with first-class security and an extensible workflow definition format.
The web UI is the intended interface for workflow runs. Start the daemon, open the printed URL, and drive runs from the Workflows page — the state-machine graph above is live, the agent-message timeline streams with markdown rendering, gate reviews include a workspace + artifac

[truncated]
