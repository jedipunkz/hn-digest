---
source: "https://github.com/hamidi-dev/opentab"
hn_url: "https://news.ycombinator.com/item?id=48845224"
title: "Show HN: OpenTab – a lazygit-style TUI for your AI token spend"
article_title: "GitHub - hamidi-dev/opentab: 📊 Browse your AI coding spend in the terminal — OpenCode, Claude Code, Codex & more · GitHub"
author: "hamidi-dev"
captured_at: "2026-07-09T13:44:08Z"
capture_tool: "hn-digest"
hn_id: 48845224
score: 2
comments: 0
posted_at: "2026-07-09T13:07:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OpenTab – a lazygit-style TUI for your AI token spend

- HN: [48845224](https://news.ycombinator.com/item?id=48845224)
- Source: [github.com](https://github.com/hamidi-dev/opentab)
- Score: 2
- Comments: 0
- Posted: 2026-07-09T13:07:21Z

## Translation

タイトル: HN を表示: OpenTab – AI トークンの支出のための Lazygit スタイルの TUI
記事のタイトル: GitHub - hamidi-dev/opentab: 📊 ターミナルで AI コーディングの支出を参照 — OpenCode、Claude Code、Codex など · GitHub
説明: 📊 ターミナルで AI コーディングの支出を参照 — OpenCode、Claude Code、Codex など - hamidi-dev/opentab

記事本文:
GitHub - hamidi-dev/opentab: 📊 ターミナルで AI コーディングの支出を参照 — OpenCode、Claude Code、Codex など · GitHub
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
ハミディ開発
/
オープンタブ
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
170 コミット 170 コミット .github/ workflows .github/ workflows docs/assets docs/assets フック フック スクリプト スクリプト src/ opentab src/ opentab .editorconfig .editorconfig .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md install.sh install.sh pyproject.toml pyproject.toml ruff.toml ruff.toml test_opentab.py test_opentab.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング ツールは常に監視します。 OpenTab で開きます。
匿名化されたデモ データ — リールをクリックするとフル品質のビデオが表示されます。
1 つのリール、すべてのビュー — トレンド、カレンダーの支出ヒートマップ、1 か月から 1 つのセッションまでのドリルダウン、ライブ テーマ
また、Web ブラウザ — opentab --web は、同じデータを 1 つの自己完結型の共有可能なページとしてレンダリングします。
AI コーディング費用のためのローカルの標準ライブラリ ターミナル UI。それはあなたの記録を読み取ります
コーディング ツールはすでにディスク上に保存されています — OpenCode の SQLite データベース、
クロード コードとコーデックス
セッションのトランスクリプトに加え、Hermes、GitHub Copilot (CLI および VS Code)、pi-agent、OpenClaw、および CSV/JSONL
API リクエストのログ - 月、日、プロジェクトごとに、トークンとお金がどこに行ったかを示します。
セッション、モデルからサブエージェント ツリーまで。一度に 1 つのツールを参照するか、すべてのツールを結合します。
あなたのツールはすでにこの台帳を保持しています。 OpenTab は単なるリーダーです。バックエンドもテレメトリーもありません。
アカウントなし - これらのファイルは読み取り専用で開きます。実行時は標準ライブラリのみ
( Curses + sqlite3 ): pipx install opentab-ai だけで、他に取り込むものは何もありません。
OpenCode、Claude Code、Codex、Hermes、GitHub Copilot (VS Code の CLI および Copilot Chat)、pi-agent、OpenClaw、およびログに記録されたリクエストの CSV/JSONL ファイルを、一度に 1 つのツールずつ、または 1 つのツールにマージして読み取ります。

単一ビュー
月、日、プロジェクト、セッション、モデルごとのコスト
トレンドオーバーレイ: 日次/週次/月次チャート、カレンダー支出ヒートマップ、モデル、プロバイダー、ソース支出ランキング
コストシェアの割合とインライン支出バー
セッションごとのモデルの組み合わせとトークンの内訳
時間の経過に伴うターンごとのコストとツール呼び出しごとのトークン使用量
作業を委任したセッションに対する再帰的なサブエージェントのコスト
「What-if」価格設定 ( $ ): 価格未定のサブスクリプション/クレジット使用量を再価格設定します。
models.dev API リストのレート。 P はその後ろにある価格表を示します
メイン リポジトリに折り畳まれた Git ワークツリー
ライブファジーフィルター (fzf スタイル、タイトル / プロジェクト / ID) およびライブ日付範囲スコープ
自己完結型 HTML ブラウザ ( --html ) — 同じドリルイン、カレンダー ヒートマップ、および
$ 1 つの共有可能なファイルと、そのファイル用のローカル ライブ サーバー ( --serve 、または
--web ブラウザで開くこともできます)
キーボードおよびマウス駆動 (スクロール、クリックして選択、ダブルクリックしてドリル)
範囲、並べ替え、無視したプロジェクト、および実行間の $ ビューを記憶します。
カラーテーマ (opentab、Catppuccin、Tokyo Night、Gruvbox、Nord、Dracula、Rosé Pine —)
ライトとダーク) TUI ( C ) と Web ブラウザーによって共有され、 --theme で設定されます
読み取り専用、ローカル専用、標準ライブラリ ランタイム (追加で取り込むものは何もありません)
スクリーンショットとライブデモのデモモード
CLI だけでなくブラウザを使用する理由
トークンの合計を出力するツールは数多くあります。 OpenTab はそれらを探索するために構築されています。
単発のレポートではなく、インタラクティブです。ドリル月→日→プロジェクト→セッション→
モデルを作成し、リストをライブでファジーフィルターし、その場で日付範囲を再スコープし、並べ替え、
キーボードまたはマウスでナビゲートします。テーブルを再実行するのではなく、lazygit スタイルのブラウザを使用します。
さまざまなフラグ。
サブエージェントのコスト ツリー。セッションが作業を委任すると、OpenTab はコストを帰属させます。
再帰的なサブエージェント サブツリー全体にわたって、つまり、

支出がどこに消えたかを確認するのではなく、
セッションの合計だけです。
標準ライブラリのランタイム。標準ライブラリから Curses + sqlite3 を追加するだけです。
ノードも npx も、実行するサービスもありません。 pipx をインストールすると、opentab-ai はどこでも実行できます
ロックダウン ボックス (唯一の依存関係である windows-curses 、
ネイティブ Windows でのみ取り込まれます)。
サブスクリプション使用の正直なコスト。サブスクリプション/クレジットセッションは真実を示します
$0 が記録され、$ ビューは API リスト レートでトークンの価格を再設定します。これは明らかです
「従量制の場合にかかる費用」の見積もりのオンとオフを切り替えることができます。
端末に単一の番号だけが必要な場合は、使用法 CLI がその役割を果たします。オープンタブは
出費を少しでも増やしたいときのために。 (「コストの精度に関する注意事項」も参照してください。)
ローカルのみ、ネットワーク、テレメトリ、アカウントなし - すべてのソース ファイルを開きます
読み取り専用なので、いずれも変更されません。完全な透明性を実現するには、すべてを
すべて自分のマシン上で操作できます。
ツール自身のレコードを読み取り専用で読み取ります: OpenCode の SQLite DB、JSONL
Claude Code / Codex / pi-agent / OpenClaw、Hermes の SQLite DB、Copilot の転写
CLI の OpenTelemetry エクスポート、VS Code の Copilot Chat セッション ストア、および CSV/JSONL
ログに記録された API リクエスト ( --csv / --jsonl )。
git ワークツリーをメイン リポジトリに折りたたむために、プロジェクト .git ファイルも読み取ります (git はありません)
プロセスが生成されます。 --no-worktrees で無効にします)。
~/.config/opentab/state.json に小さな設定ファイルを書き込みます (最後の
ソース、範囲、並べ替え。 --no-state で無効にします)、オプションのモデル価格キャッシュ
~/.config/opentab/prices.json (--refresh-models を実行するか、
P オーバーレイ)、および - e を押すか --html を実行した場合のみ - opentab-*.csv
現在のディレクトリにある HTML ブラウザ ファイルをエクスポートします。
押したキーでのみ外部プログラムを実行します: ファイルが開きます

えーっと
o の場合は ( open / xdg-open 、または Windows のエクスプローラー)、 L の場合は tmux 、独自の
ランチャーフック ( ~/.config/opentab/launcher )、またはクリップボードツール
コピー先として ( pbcopy / wl-copy / xclip / xsel ) を指定します。全員が
--demo で無効化されています。
Python 3.9 以降と、curses を備えたターミナル — すでに macOS、Linux、
そしてWSL。ネイティブ Windows も動作します: opentab-ai をインストールすると windows-curses が取り込まれます
自動的に実行されます (「 Windows 」を参照)。
uvx --from opentab-ai opentab --demo # または: pipx run --spec opentab-ai opentab --demo
コマンド 1 つで何もインストールされない: --demo は匿名化された合成ファイルで完全な TUI を実行します。
データなので、AI ツールの履歴を読み取れないマシンでも動作します。 --demo を次の場所にドロップします
実際の使用方法にそれを向けてください。
pipx インストール opentab-ai
後で pipx upgrade opentab-ai を使用してアップグレードします。 (プレーン pip install --user opentab-ai
も機能します。) PyPI ディストリビューションは opentab-ai です。インストールするコマンドは
オープンタブ 。
brew install hamidi-dev/tap/opentab
後で brew upgrade opentab を使用してアップグレードします。
1 行 (pipx 経由で opentab コマンドをインストールします。再実行して更新します):
カール -fsSL https://raw.githubusercontent.com/hamidi-dev/opentab/main/install.sh |バッシュ
ソースから
git clone https://github.com/hamidi-dev/opentab && cd opentab
pipx をインストールします。 # またはライブ編集可能なチェックアウトの場合は `pip install -e .`
使用法
opentab # 常にブラウザを開きます
opentab --days 30 # ウィンドウ内で開始 (R でライブ変更)
opentab --2026-05-01 以降 --2026-05-31 まで
opentab --db /path/to/opencode.db # デフォルト: ~/.local/share/opencode/opencode.db
opentab --source claude # 代わりにクロード コードを参照します (以下を参照)
opentab --demo # ライブデモ/スクリーンショットにとって安全 (以下を参照)
opentab --html # opentab-report.html を書き込んで終了します (以下を参照)
opentab --http://localhost:8321 上で # 同じブラウザをライブで提供します
opentab --web # --serve、および o

ブラウザにペンを入れてください
データソース
OpenTab は、各 AI コーディング ツールが保持するローカル レコードを読み取ります — OpenCode (SQLite)、
クロード コード ( ~/.claude/projects/**/*.jsonl )、コーデックス
( ~/.codex/sessions/**/rollout-*.jsonl )、GitHub Copilot CLI (OpenTelemetry)
~/.copilot/otel/ にエクスポート)、VS Code の Copilot Chat (VS Code 独自のもの)
chatSessions ストア)、pi-agent ( ~/.pi/agent/sessions/ )、および OpenClaw
( ~/.openclaw/agents/**/sessions/ ) に加えて、Hermes データベースと汎用 CSV または
ログに記録された API リクエストの JSONL。 --source を使用していずれかを選択します。
opentab --source opencode # OpenCode のみ
opentab --source claude --claude-dir /path # クロード コード (デフォルト ~/.claude/projects)
opentab --source codex --codex-dir /path # Codex (デフォルト ~/.codex/sessions)
opentab --source copilot # GitHub Copilot CLI (デフォルト ~/.copilot/otel)
opentab --source vscode # VS Code でのコパイロット チャット (インストールされているすべてのバリアント)
opentab --source pi # pi-agent (デフォルト ~/.pi/agent/sessions)
opentab --source openclaw # OpenClaw ゲートウェイ (デフォルト ~/.openclaw; $OPENCLAW_DIR を優先)
opentab --csvrequests.csv #ログに記録された API リクエストの CSV (または --jsonlrequests.jsonl)
opentab --source all # 存在するすべてのソースをマージ
GitHub Copilot CLI: OpenTelemetry エクスポートが行われた場合にのみ使用状況を記録します。
有効になりました。セッションの起動/再開前に COPILOT_OTEL_FILE_EXPORTER_PATH を設定します。
(export COPILOT_OTEL_FILE_EXPORTER_PATH=~/.copilot/otel/usage.jsonl );後のセッション
--source copilot の下に表示されます。
VS Code のコパイロット チャット: VS Code 自体のチャット セッション ストアから読み取ります
( <User>/workspaceStorage/*/chatSessions 、および空のウィンドウ セッション)、コード全体、
コード - Insider および VSCodium — 有効にするものはありません。プロジェクトはそれぞれから生まれます
ワークスペースのフォルダー。パネルが開かれただけのセッション (トークンなし) は無視されます。ポイント
ユーザーディレクトリの --vscode-dir

ポータブル/リモート コピーの場合 — WSL から、
Windows 側のストア (「 Windows 」を参照)。
OpenClaw: セッションは ~/.openclaw/agents/<agent>/sessions/ に存在します (1 つのプロジェクトにつき 1 つ)
エージェント);サーバー上で実行されている場合は、マウントされたコピーを --openclaw-dir / $OPENCLAW_DIR で指定します。
すべてのプロバイダーのメッセージごとのコストを記録しますが、従量制のルート (直接ルート) のみを記録します。
Anthropic/OpenRouter キー) は実際の支出です — 計画ルート (openai-codex、github-copilot) は
$ の下で推定され、openclaw.json から読み取られます (読み取り専用)。
--source auto (デフォルト) は最後に使用したソースを復元します。それ以外の場合は、存在するすべてのソースをマージします。
複数のソースが存在する場合はソース (1 つだけの場合は単一ソース)。アクティブなソース
ヘッダチップとして表示されます。 switch live with c (存在するソースを循環させ、
プラスすべて)。 TUI 全体は同じように機能します - 月、日、プロジェクト、セッション、モデル、
傾向 — クロード コード、コーデックス、およびコパイロット (CLI と VS) のため、2 つの違いがあります。
コード）トークンのみを記録し、メッセージごとのコストはかかりません。
このようなセッションは、OpenCode サブスクリプション セッションのように動作します。通常モードでは $0、
$ ビューで (トークン × API 定価) を見積もります。そのような見方になるので、
それ以外の場合は $0.00 の壁があり、見積もりはデフォルトでそこから始まります (ヘッダー タグ:
推定);記録された数字に対して $ を押すと、選択が変更されます

[切り捨てられた]

## Original Extract

📊 Browse your AI coding spend in the terminal — OpenCode, Claude Code, Codex & more - hamidi-dev/opentab

GitHub - hamidi-dev/opentab: 📊 Browse your AI coding spend in the terminal — OpenCode, Claude Code, Codex & more · GitHub
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
hamidi-dev
/
opentab
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
170 Commits 170 Commits .github/ workflows .github/ workflows docs/ assets docs/ assets hooks hooks scripts scripts src/ opentab src/ opentab .editorconfig .editorconfig .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md install.sh install.sh pyproject.toml pyproject.toml ruff.toml ruff.toml test_opentab.py test_opentab.py View all files Repository files navigation
Your AI coding tools keep a tab. OpenTab opens it.
Anonymized demo data — click the reel for the full-quality video.
One reel, every view — trends, a calendar spend heatmap, drill from a month down to a single session, and live theming
Also a web browser — opentab --web renders the same data as one self-contained, shareable page
A local, standard-library terminal UI for your AI coding spend. It reads the records your
coding tools already keep on disk — OpenCode 's SQLite database,
Claude Code 's and Codex 's
session transcripts, plus Hermes, GitHub Copilot (CLI and VS Code), pi-agent, OpenClaw, and CSV/JSONL
logs of API requests — and shows where your tokens and money went: by month, day, project,
session, and model, down to the subagent tree. Browse one tool at a time, or merge them all.
Your tools already keep this ledger; OpenTab is just the reader. No backend, no telemetry,
no accounts — it opens those files read-only . Standard-library-only at runtime
( curses + sqlite3 ): pipx install opentab-ai and there's nothing else to pull in.
Reads OpenCode, Claude Code, Codex, Hermes, GitHub Copilot (its CLI and Copilot Chat in VS Code), pi-agent, OpenClaw, and logged-request CSV/JSONL files — one tool at a time, or merged into a single view
Cost by month, day, project, session, and model
Trends overlay: daily / weekly / monthly charts, a calendar spend heatmap, and model-, provider- and source-spend rankings
Cost-share percentages and inline spend bars
Per-session model mix and token breakdown
Per-turn cost over time, and token usage per tool call
Recursive subagent costs, on the sessions that delegated work
"What-if" pricing ( $ ): re-price unpriced subscription/credit usage at
models.dev API list rates; P shows the price table behind it
Git worktrees folded into their main repo
Live fuzzy filter (fzf-style, title / project / id) and live date-range scoping
A self-contained HTML browser ( --html ) — same drill-in, calendar heatmap, and
$ toggle in one shareable file — and a local live server for it ( --serve , or
--web to also pop it open in your browser)
Keyboard- and mouse-driven (scroll, click to select, double-click to drill)
Remembers your range, sort, ignored projects, and the $ view between runs
Colour themes (opentab, Catppuccin, Tokyo Night, Gruvbox, Nord, Dracula, Rosé Pine —
light and dark) shared by the TUI ( C ) and the web browser, set with --theme
Read-only, local-only, standard-library runtime (nothing extra to pull in)
Demo mode for screenshots and live demos
Why a browser, not just a usage CLI
Plenty of tools will print your token totals. OpenTab is built to explore them:
Interactive, not a one-shot report. Drill month → day → project → session →
model, fuzzy-filter the lists live, rescope the date range on the fly, sort, and
navigate by keyboard or mouse — a lazygit-style browser, not a table you re-run with
different flags.
Subagent cost trees. When a session delegated work, OpenTab attributes the cost
across its whole recursive subagent subtree — so you see where the spend went, not
just the session total.
Standard-library runtime. Just curses + sqlite3 from the standard library:
no Node, no npx , no service to run. pipx install opentab-ai and it runs anywhere
Python 3.9+ exists, including a locked-down box (the sole dependency, windows-curses ,
is pulled in only on native Windows).
Honest cost for subscription usage. Subscription/credit sessions show a truthful
$0 recorded, and the $ view reprices their tokens at API list rates — a clear
"what this would have cost metered" estimate you can toggle on and off.
If you just want a single number in your terminal, a usage CLI does the job. OpenTab is
for when you want to poke at the spend. (See also A note on cost accuracy .)
Local-only, no network, no telemetry, no accounts — it opens every source file
read-only , so it doesn't modify any of them. For full transparency, everything it
touches, all on your own machine:
Reads your tools' own records, read-only: OpenCode's SQLite DB, the JSONL
transcripts of Claude Code / Codex / pi-agent / OpenClaw, Hermes' SQLite DB, the Copilot
CLI's OpenTelemetry export, VS Code's Copilot Chat session store, and a CSV/JSONL of
logged API requests ( --csv / --jsonl ).
To fold git worktrees into their main repo it also reads project .git files (no git
process is spawned; disable with --no-worktrees ).
Writes a small preferences file at ~/.config/opentab/state.json (your last
source, range, and sort; disable with --no-state ), an optional model-price cache at
~/.config/opentab/prices.json (only when you run --refresh-models or press r in the
P overlay), and — only when you press e or run --html — an opentab-*.csv
export or the HTML browser file in the current directory.
Runs external programs only on the key you press: your file opener
( open / xdg-open , or Explorer on Windows) for o , and for L either tmux , your own
launcher hook ( ~/.config/opentab/launcher ), or your clipboard tool
( pbcopy / wl-copy / xclip / xsel ) for its copy target. All are
disabled in --demo .
Python 3.9+ and a terminal with curses — already present on macOS, Linux,
and WSL. Native Windows works too: installing opentab-ai pulls in windows-curses
automatically (see Windows ).
uvx --from opentab-ai opentab --demo # or: pipx run --spec opentab-ai opentab --demo
One command, nothing installed: --demo runs the full TUI on anonymized synthetic
data, so it works even on a machine with no AI-tool history to read. Drop --demo to
point it at your real usage.
pipx install opentab-ai
Upgrade later with pipx upgrade opentab-ai . (Plain pip install --user opentab-ai
works too.) The PyPI distribution is opentab-ai ; the command it installs is
opentab .
brew install hamidi-dev/tap/opentab
Upgrade later with brew upgrade opentab .
One line (installs the opentab command via pipx; re-run to update):
curl -fsSL https://raw.githubusercontent.com/hamidi-dev/opentab/main/install.sh | bash
From source
git clone https://github.com/hamidi-dev/opentab && cd opentab
pipx install . # or `pip install -e .` for a live-editable checkout
Usage
opentab # open the browser, all time
opentab --days 30 # start within a window (change live with R)
opentab --since 2026-05-01 --until 2026-05-31
opentab --db /path/to/opencode.db # default: ~/.local/share/opencode/opencode.db
opentab --source claude # browse Claude Code spend instead (see below)
opentab --demo # safe for live demos / screenshots (see below)
opentab --html # write opentab-report.html and exit (see below)
opentab --serve # same browser on http://localhost:8321, live
opentab --web # --serve, and open it in your browser
Data sources
OpenTab reads the local records each AI coding tool keeps — OpenCode (SQLite),
Claude Code ( ~/.claude/projects/**/*.jsonl ), Codex
( ~/.codex/sessions/**/rollout-*.jsonl ), the GitHub Copilot CLI (its OpenTelemetry
export under ~/.copilot/otel/ ), Copilot Chat in VS Code (VS Code's own
chatSessions store), pi-agent ( ~/.pi/agent/sessions/ ), and OpenClaw
( ~/.openclaw/agents/**/sessions/ ), plus a Hermes database and a generic CSV or
JSONL of logged API requests. Pick one with --source :
opentab --source opencode # OpenCode only
opentab --source claude --claude-dir /path # Claude Code (default ~/.claude/projects)
opentab --source codex --codex-dir /path # Codex (default ~/.codex/sessions)
opentab --source copilot # GitHub Copilot CLI (default ~/.copilot/otel)
opentab --source vscode # Copilot Chat in VS Code (every installed variant)
opentab --source pi # pi-agent (default ~/.pi/agent/sessions)
opentab --source openclaw # OpenClaw gateway (default ~/.openclaw; honors $OPENCLAW_DIR)
opentab --csv requests.csv # a CSV of logged API requests (or --jsonl requests.jsonl)
opentab --source all # all present sources, merged
GitHub Copilot CLI: it records usage only when its OpenTelemetry export is
enabled. Set COPILOT_OTEL_FILE_EXPORTER_PATH before launching/resuming a session
( export COPILOT_OTEL_FILE_EXPORTER_PATH=~/.copilot/otel/usage.jsonl ); sessions after
that show up under --source copilot .
Copilot Chat in VS Code: read from VS Code's own chat-session store
( <User>/workspaceStorage/*/chatSessions , plus empty-window sessions), across Code,
Code - Insiders, and VSCodium — nothing to enable. Projects come from each
workspace's folder; sessions the panel merely opened (no tokens) are ignored. Point
--vscode-dir at a User directory for a portable/remote copy — from WSL, at the
Windows-side store (see Windows ).
OpenClaw: sessions live under ~/.openclaw/agents/<agent>/sessions/ (one project per
agent); point --openclaw-dir / $OPENCLAW_DIR at a mounted copy if it runs on a server.
It records a per-message cost for every provider, but only metered routes (a direct
Anthropic/OpenRouter key) are real spend — plan routes (openai-codex, github-copilot) are
estimated under $ , read from openclaw.json (read-only).
--source auto (the default) restores your last-used source, else merges every present
source when more than one exists (a single source when only one is). The active source
shows as a header chip; switch live with c (cycles whichever sources are present,
plus all ). The whole TUI works the same — months, days, projects, sessions, models,
trends — with two differences, because Claude Code, Codex, and Copilot (CLI and VS
Code) record only tokens, no per-message cost :
Such a session works like an OpenCode subscription session: $0 in normal mode and an
estimate (tokens × API list price) under the $ view. Since that view would
otherwise be a wall of $0.00 , the estimate starts on by default there (header tag:
ESTIMATED ); press $ for the recorded numbers, and your choice is re

[truncated]
