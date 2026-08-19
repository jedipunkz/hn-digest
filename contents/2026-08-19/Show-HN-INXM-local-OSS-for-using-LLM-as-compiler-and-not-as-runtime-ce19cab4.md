---
source: "https://github.com/inxm-ai/inxm-local"
hn_url: "https://news.ycombinator.com/item?id=49362974"
title: "Show HN: INXM // local` OSS for using LLM as compiler and not as runtime"
article_title: "GitHub - inxm-ai/inxm-local: The LLM is the compiler, not the runtime · GitHub"
image: "https://opengraph.githubassets.com/b477467390572ba6149955d6ff310ae7c36fb5a8380b852fc9c83405b92df561/inxm-ai/inxm-local"
author: "oesimania"
captured_at: "2026-08-19T16:20:23Z"
capture_tool: "hn-digest"
hn_id: 49362974
score: 4
comments: 3
posted_at: "2026-08-19T15:35:29Z"
tags:
  - hacker-news
  - translated
---

# Show HN: INXM // local` OSS for using LLM as compiler and not as runtime

- HN: [49362974](https://news.ycombinator.com/item?id=49362974)
- Source: [github.com](https://github.com/inxm-ai/inxm-local)
- Score: 4
- Comments: 3
- Posted: 2026-08-19T15:35:29Z

## Translation

タイトル: Show HN: INXM // LLM をランタイムとしてではなくコンパイラとして使用するためのローカル ` OSS
記事のタイトル: GitHub - inxm-ai/inxm-local: LLM はランタイムではなくコンパイラーである · GitHub
説明: LLM はランタイムではなくコンパイラーです。 GitHub でアカウントを作成して、inxm-ai/inxm-local の開発に貢献してください。

記事本文:
GitHub - inxm-ai/inxm-local: LLM はランタイムではなくコンパイラーです · GitHub
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
inxm-ai
/
inxm-local
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
1 コミット 1 コミット フォルダーとファイル
.github .github 資産 アセット ドキュメント ドキュメント 例-構成 例-構成 例 例 パッケージ化 パッケージ化スキル スキル src src te

lemetry-worker telemetry-worker テスト テスト .gitignore .gitignore Agents.md Agents.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス通知 通知 README.md README.md build.rs build.rs すべて表示ファイル リポジトリ ファイルのナビゲーション
コンパイル済み AI ワークフロー用のローカルファーストの Rust デスクトップ アプリ
LLM はコンパイラであり、ランタイムではありません。チャットで意図を説明すると、
コンパイラは型付きプランを生成し、決定論的実行プログラムがそれを実行します。
実行パスでの AI の即興はありません。
計画を作成するためのチャット — 平易な言語と設定された LLM を入力します
それを検証済みのバージョン管理された計画に変換します。 API キー、既存の
Codex/Claude Code ログイン、または互換性のあるローカル/ホスト型エンドポイント。スラッシュ
コマンド ( /run 、 /plans 、 /repair など) は、他のすべてを実行します。
アニメーション コマンド パレット ( / を入力し、Tab キーを押して完了します)。
プラン所有の会話 — すべてのプランには 1 つの永続的なチャットがあります。オープニング
プランまたはその実行の 1 つが、カードを挿入する代わりにそのチャットに移動します
現在開いている会話に参加します。固定ワークスペースカードでプランを維持
コントロール、ライブ進行状況、詳細、および完全な実行履歴が表示されます
スクロール可能なトランスクリプトの上。
再利用可能な型指定されたプラン入力 — コンパイルされたプランは、によって提供される値を宣言します。
各トリガー (クエリ、ターゲット、受信者、制限、環境など)。
入力は検証され、 ${input.<name>} として利用可能で、実行とともに永続化されます。
各スケジュールごとに独立してキャプチャされます。
決定的な実行 — 移植されたソロプレイヤー エグゼキューターはステップを実行します。
トポロジー的順序、各ステップ後の状態と解決された入力の保持、および
進行状況を計画カードにライブでストリーミングします。
人間参加型 — HUMAN_INTERACTION ステップで実行を一時停止し、参加を求める
チャット (承認/拒否ボタンまたは自由記述の回答)。
修復ループ —

失敗した実行はコンパイラに戻すことができます
( /repair <実行 ID> );提案されたパッチは、適用するカードとして表示されます。または
拒否します。パッチを適用すると、新しい計画バージョンが作成されます。
UI での MCP 管理 - MCP ツール ビューにツールがリストされます
カタログを作成し、ローカル stdio またはリモート Streamable を追加/編集/削除できます
HTTP MCP サーバー (およびサブプロセスおよび HTTP ツール)。変更は次まで持続します
データ ディレクトリ内の tools.yaml
ローカル HTTP MCP サーバー - デスクトップ クライアントはローカル MCP サーバーを起動します。
他のクライアントがコンパイル、検索/表示、実行、修復、編集、
同じ決定論的なコアを通じてワークフローをスケジュールし、検査します。
Linux および macOS では、これによりマシンの最新リリースがダウンロードされ、
ユーザーごとにインストールします (root は必要ありません)。
カール -fsSL https://raw.githubusercontent.com/inxm-ai/inxm-local/main/packaging/install.sh |しー
Windows の場合 (PowerShell):
irm https://raw.githubusercontent.com/inxm - ai / inxm - local / main /Packaging / install.ps1 |アイエックス
インストーラーは、INXM Local のローカル MCP サーバーを
コーディング エージェントを同じステップで実行できるため、INXM ワークフローをコンパイルして実行できます。
すぐに：
# このマシン上で見つかったすべてのエージェントに登録します
カール -fsSL https://raw.githubusercontent.com/inxm-ai/inxm-local/main/packaging/install.sh | sh -s -- --エージェント
# または特定のエージェントを選択します
カール -fsSL https://raw.githubusercontent.com/inxm-ai/inxm-local/main/packaging/install.sh | sh -s -- --claude --codex --cursor
サポートされているエージェントと、それぞれに対するインストーラーの機能:
既存の構成ファイルは上書きされずにマージされ、登録は
冪等 — インストーラを再実行してもエントリが重複することはありません。セット
INXM_MCP_URL を使用して、デフォルト以外のエンドポイントを登録します。
Windows では、install.ps1 をダウンロードして実行します。
-Agents または同じものに対応するスイッチ ( -Claude 、 -Cursor 、...)
再

指示。
その他の便利なフラグ: --autostart (Linux: ログイン時に非表示で開始)、
--version 0.1.0 (リリースを固定)、--uninstall (アプリとエージェントを削除)
登録）。両方のスクリプトも各リリースに添付されているため、
https://github.com/inxm-ai/inxm-local/releases/latest/download/install.sh
も機能します。
プラットフォーム用の最新パッケージを次の場所からダウンロードします。
GitHub のリリース ページ。
Windows (x86-64): .exe インストーラーをダウンロードして実行します。インストールします
ユーザーごとに (管理者権限は必要ありません)、オプションで INXM Local を登録できます
ログイン時にシステム トレイに隠れて起動します。SmartScreen が警告します。
未署名のインストーラーについて — を参照してください。
未署名のビルド。
macOS: Apple Silicon ( aarch64 ) または Intel の .app.zip をダウンロードします。
( x86_64 ) を解凍し、 INXM Local を開きます。管理者権限は必要ありません —
/Applications に書き込めない場合は、 ~/Applications にドロップしてください。
ゲートキーパーは公証されていないアプリをブロックします — を参照してください。
未署名のビルド。
Linux (Debian/Ubuntu、x86-64): .deb をダウンロードし、次のようにインストールします。
sudo apt install ./inxm-local-x86_64-unknown-linux-gnu.deb 。のために
root を使用せずにユーザーごとにインストールし、代わりに .tar.gz をダウンロードして解凍し、
./install.sh を実行します (ログイン時に INXM Local を開始するには --autostart を追加します。
--uninstall は再度削除します)。
未署名のビルド: macOS Gatekeeper および Windows SmartScreen
私たちのリリースはまだ Apple Developer ID または Windows コードで署名されていません
証明書に署名しているため、両方のオペレーティング システムが警告 (またはブロック) します。
ブラウザでダウンロードされたビルド。ビルドは安全です - すべてのリリースがビルドされます
GitHub Actions によってこのリポジトリから。まで
署名が設定されている場合は、次の回避策を使用してください。
macOS — ブラウザでダウンロードしたアプリは隔離されており、アプリは隔離されているため、
公証されていないため、macOS が「破損している」と報告するか、「開けない」と報告する
Apple がそれを確認できなかったためです。」どちらか：

解凍後に隔離フラグを削除します。
xattr -dr com.apple.quarantine ~ /Applications/ " INXM Local.app "
または、アプリを一度開いてから、「システム設定」→「プライバシーと」に移動してみてください。
[セキュリティ] をクリックし、ブロックされたアプリの通知まで下にスクロールし、[とにかく開く] をクリックします。
または、クイック インストール スクリプト (curl) を使用します。
ダウンロードでは隔離フラグが設定されないため、アプリは通常どおり開きます。
Windows — SmartScreen を実行すると、「Windows が PC を保護しました」と表示されます。
署名のないインストーラー。次のいずれか:
SmartScreen ダイアログで「詳細」→「とにかく実行」をクリックします。
または、実行する前に、Mark-of-the-Web をクリアします。
ブロック解除ファイル .\ inxm-local-x86_64-pc-windows-msvc-setup.exe
一部の企業ポリシーは、署名されていない実行可能ファイルを完全にブロックします
(SmartScreen が「警告してバイパスを防止する」または AppLocker ルールに設定されています)。その中で
この場合、IT 管理者にインストーラーを許可リストに登録するよう依頼するか、次のコマンドを使用してソースからビルドするように依頼してください。
カーゴ ビルド --リリース 。
Debian または Ubuntu では、このワンライナーで最新の x86-64 をダウンロードしてインストールします。
リリース:
カール -fL https://github.com/inxm-ai/inxm-local/releases/latest/download/inxm-local-x86_64-unknown-linux-gnu.deb -o /tmp/inxm-local.deb && sudo apt install -y /tmp/inxm-local.deb
次に、アプリケーション メニューまたはターミナルから inxm-local を起動します。
「設定」→「コンパイラー」を開き、接続を 1 つ選択します。
Claude API / OpenAI API — API キーを入力するか、設定します
アプリを起動する前に ANTHROPIC_API_KEY / OPENAI_API_KEY を入力してください。
OpenAI アカウント / Claude アカウント — コーデックスをインストールしてサインインするか、
クロードCLI。アプリは非対話的に CLI を呼び出すため、API キーは必要ありません。
INXMに保存されます。
カスタム OpenAI URL / カスタム Anthropic URL — ベース URL とモデルを入力します。
API キーはオプションであり、Ollama、LM Studio、
llama.cpp、vLLM、または別の互換性のあるゲートウェイ。 OpenAI対応の場合
サーバーの API ルート (例: http://localhost:1) を入力します。

1434/v1 )、ではありません
完全な /chat/completions パス。
選択した接続とモデルは、プランの作成、修復/編集によって共有されます
リクエスト、および実行中の PROMPT_CALL ステップ。既存の settings.json
ファイルの互換性は維持され、構成された Claude/OpenAI が引き続き使用されます。
API キー バックエンド。
データ (計画、実行、パッチ、tools.yaml) はプラットフォーム データ ディレクトリに存在します。
(Linux の場合は ~/.local/share/inxm-local/);で上書きする
INXM_LOCAL_DATA_DIR=/パス 。エコーツールを備えたスターターカタログは次のとおりです。
最初の起動時にシードされます。カタログの例は、examples-config/tools.yaml にあります。
コマンド
効果
(プレーンテキスト) / /compile <目的>
新しいチャットで計画を作成します。リンクされたチャットで所有プランを調整する
/plans 、 /runs 、 /tools
計画/実行/カタログの一覧表示
/show <計画>
所有するチャットでプランを開きます (ID プレフィックスまたは名前)
/run <プラン> [--inputs '<json>']
呼び出し入力を使用してプランを実行する
/inspect <実行ID>
実行のステップステータス、タイミング、エラー
/repair <実行ID>
失敗した実行に対するパッチを提案する
/apply <パッチ ID> / /reject <パッチ ID> [理由]
パッチを解決する
/schedule <プラン> <cron> [--inputs '<json>'] / /schedules
キャプチャした入力を使用してスケジュールを作成/リストする
/ヘルプ、/クリア
ヘルプ / チャットをクリアする
バックグラウンドでのスケジュールの実行
[スケジュールを実行し続ける] の場合、デスクトップ アプリはシステム トレイで実行され続けます。
バックグラウンドでの機能は「設定」で有効になります。オプションがオンになります
有効なスケジュールが存在する場合、自動的に実行されます。トレイ メニューを使用して、
ウィンドウ、個々の状態を変更せずにすべてのスケジュールを一時停止または再開します。
またはプロセスを完全に終了します。
サーバーおよび無人マシンの場合、ヘッドレス モードでは MCP サーバーが実行され、
ウィンドウのないスケジューラ:
inxm-local --headless # または: INXM_HEADLESS=1 inxm-local
プラットフォームのサービスマネージャーでログアウトした後も実行し続けます。例:
nohup inxm-local --headless

> /tmp/inxm-headless.log 2>&1 &
または systemd ユーザーユニット ( ~/.config/systemd/user/inxm-local.service ):
【単位】
説明 =INXM // ローカルのヘッドレス スケジューラ
【サービス】
ExecStart =%h/.local/bin/inxm-local --headless
再起動 = 失敗時
[インストール]
WantedBy =default.target
( systemctl --user enable --now inxm-local )。 Windows では、タスク スケジューラを使用します
同じ --headless 引数を使用します。
データ ディレクトリごとに実行されるスケジューラは 1 つだけです。それは、scheduler.lock ファイル (
所有者の PID) は、デスクトップ アプリとヘッドレス インスタンスの起動を防ぎます。
同じスケジュールを2回。 2 番目のインスタンスはライブ ホルダーを検出し、スキップします。
そのスケジューラ。クラッシュしたプロセスからの古いロックは再利用されます
自動的に。何も実行されていない間に欠落したスロットは追いつきません。
デザイン。
デスクトップが起動すると、アプリはローカルの Streamable HTTP スタイルの MCP サーバーを起動します。
クライアントが起動します。デフォルトでは、ループバックでのみリッスンします。
http://127.0.0.1:39387/mcp
ポートは settings.json に保存されており、以下で変更できます。
設定 → ローカル MCP サーバー 。起動時にポートをバインドできない場合 (たとえば、
別のプロセスがすでにそれを使用しています)、サイドバー/フッターと設定ビューに表示されます。
バインドエラーによる警告。別のポートを選択し、設定を保存して再起動します
アプリ。
自動化またはネイティブ Wi-Fi が使用される環境の場合

[切り捨てられた]

## Original Extract

The LLM is the compiler, not the runtime. Contribute to inxm-ai/inxm-local development by creating an account on GitHub.

GitHub - inxm-ai/inxm-local: The LLM is the compiler, not the runtime · GitHub
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
inxm-ai
/
inxm-local
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github .github assets assets docs docs examples-config examples-config examples examples packaging packaging skills skills src src telemetry-worker telemetry-worker tests tests .gitignore .gitignore Agents.md Agents.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE NOTICE NOTICE README.md README.md build.rs build.rs View all files Repository files navigation
A local-first, Rust desktop app for compiled-AI workflows
The LLM is the compiler , not the runtime. You describe intent in chat,
the compiler produces a typed plan, and a deterministic executor runs it.
No AI improvisation in the execution path.
Chat to create plans — type plain language and the configured LLM
turns it into a validated, versioned plan. Use an API key, an existing
Codex/Claude Code login, or a compatible local/hosted endpoint. Slash
commands ( /run , /plans , /repair , …) drive everything else, with an
animated command palette (type / , Tab to complete).
Plan-owned conversations — every plan has one persistent chat. Opening
a plan or one of its runs navigates to that chat instead of inserting a card
into the currently open conversation. A fixed workspace card keeps plan
controls, live progress, details, and complete execution history visible
above the scrollable transcript.
Reusable, typed plan inputs — compiled plans declare values supplied by
each trigger (for example query, target, recipient, limit, or environment).
Inputs are validated, available as ${input.<name>} , persisted with runs,
and captured independently by each schedule.
Deterministic runs — the ported soloplayer executor runs steps in
topological order, persists state and resolved inputs after every step, and
streams live progress into the plan card.
Human-in-the-loop — HUMAN_INTERACTION steps pause the run and ask in
chat (Approve / Reject buttons or a free-text answer).
Repair loop — a failed run can be handed back to the compiler
( /repair <run-id> ); the proposed patch appears as a card you apply or
reject. Applied patches create a new plan version.
MCP management in the UI — the MCP Tools view lists the tool
catalog and lets you add / edit / delete local stdio or remote Streamable
HTTP MCP servers (plus subprocess and HTTP tools). Changes persist to
tools.yaml in the data dir.
Local HTTP MCP server — the desktop client starts a local MCP server on
launch so other clients can compile, find/show, execute, repair, edit,
schedule, and inspect workflows through the same deterministic core.
On Linux and macOS, this downloads the latest release for your machine and
installs it per-user (no root needed):
curl -fsSL https://raw.githubusercontent.com/inxm-ai/inxm-local/main/packaging/install.sh | sh
On Windows (PowerShell):
irm https: // raw.githubusercontent.com / inxm - ai / inxm - local / main / packaging / install.ps1 | iex
The installer can also register INXM Local's local MCP server with your
coding agents in the same step, so they can compile and run INXM workflows
right away:
# register with every agent found on this machine
curl -fsSL https://raw.githubusercontent.com/inxm-ai/inxm-local/main/packaging/install.sh | sh -s -- --agents
# or pick specific agents
curl -fsSL https://raw.githubusercontent.com/inxm-ai/inxm-local/main/packaging/install.sh | sh -s -- --claude --codex --cursor
Supported agents and what the installer does for each:
Existing config files are merged, not overwritten, and registration is
idempotent — rerunning the installer never duplicates entries. Set
INXM_MCP_URL to register a non-default endpoint.
On Windows, download install.ps1 and run it with
-Agents or the matching switches ( -Claude , -Cursor , ...) for the same
registrations.
Other useful flags: --autostart (Linux: start hidden at login),
--version 0.1.0 (pin a release), --uninstall (remove the app and agent
registrations). Both scripts are also attached to each release, so
https://github.com/inxm-ai/inxm-local/releases/latest/download/install.sh
works too.
Download the latest package for your platform from the
GitHub Releases page .
Windows (x86-64): download and run the .exe installer. It installs
per-user (no admin rights needed) and can optionally register INXM Local
to start hidden in the system tray when you log in. SmartScreen will warn
about the unsigned installer — see
unsigned builds .
macOS: download the .app.zip for Apple Silicon ( aarch64 ) or Intel
( x86_64 ), unzip it, and open INXM Local . No admin rights needed —
drop it in ~/Applications if you can't write to /Applications .
Gatekeeper will block the un-notarized app — see
unsigned builds .
Linux (Debian/Ubuntu, x86-64): download the .deb and install it with
sudo apt install ./inxm-local-x86_64-unknown-linux-gnu.deb . For a
per-user install without root, download the .tar.gz instead, unpack it,
and run ./install.sh (add --autostart to start INXM Local at login;
--uninstall removes it again).
Unsigned builds: macOS Gatekeeper & Windows SmartScreen
Our releases are not yet signed with an Apple Developer ID or a Windows code
signing certificate, so both operating systems will warn about (or block)
builds downloaded with a browser. The builds are safe — every release is built
from this repository by GitHub Actions . Until
we have signing set up, use these workarounds:
macOS — a browser-downloaded app is quarantined, and because the app is
not notarized, macOS reports it as "damaged" or says it "cannot be opened
because Apple could not verify" it. Either:
Remove the quarantine flag after unzipping:
xattr -dr com.apple.quarantine ~ /Applications/ " INXM Local.app "
Or try to open the app once, then go to System Settings → Privacy &
Security , scroll down to the blocked-app notice, and click Open Anyway .
Or use the quick install script — curl
downloads don't set the quarantine flag, so the app opens normally.
Windows — SmartScreen shows "Windows protected your PC" when you run the
unsigned installer. Either:
Click More info → Run anyway in the SmartScreen dialog.
Or clear the mark-of-the-web before running it:
Unblock-File .\ inxm-local-x86_64-pc-windows-msvc-setup.exe
Some corporate policies block unsigned executables entirely
(SmartScreen set to "Warn and prevent bypass", or AppLocker rules). In that
case ask your IT admin to allowlist the installer, or build from source with
cargo build --release .
On Debian or Ubuntu, this one-liner downloads and installs the latest x86-64
release:
curl -fL https://github.com/inxm-ai/inxm-local/releases/latest/download/inxm-local-x86_64-unknown-linux-gnu.deb -o /tmp/inxm-local.deb && sudo apt install -y /tmp/inxm-local.deb
Then launch inxm-local from your application menu or terminal.
Open Settings → Compiler and choose one connection:
Claude API / OpenAI API — enter an API key or set
ANTHROPIC_API_KEY / OPENAI_API_KEY before starting the app.
OpenAI account / Claude account — install and sign in to the codex or
claude CLI. The app invokes the CLI non-interactively, so no API key is
stored in INXM.
Custom OpenAI URL / Custom Anthropic URL — enter a base URL and model.
API keys are optional, allowing local servers such as Ollama, LM Studio,
llama.cpp, vLLM, or another compatible gateway. For an OpenAI-compatible
server, enter its API root (for example http://localhost:11434/v1 ), not
the full /chat/completions path.
The selected connection and model are shared by plan compilation, repair/edit
requests, and PROMPT_CALL steps during execution. Existing settings.json
files remain compatible and continue to use their configured Claude/OpenAI
API-key backend.
Data (plans, runs, patches, tools.yaml ) lives in the platform data dir
( ~/.local/share/inxm-local/ on Linux); override with
INXM_LOCAL_DATA_DIR=/path . A starter catalog with an echo tool is
seeded on first launch. An example catalog is in examples-config/tools.yaml .
Command
Effect
(plain text) / /compile <intent>
Compile a plan in a new chat; refine the owned plan in a linked chat
/plans , /runs , /tools
List plans / runs / catalog
/show <plan>
Open the plan in its owned chat (id prefix or name)
/run <plan> [--inputs '<json>']
Execute a plan with invocation inputs
/inspect <run-id>
Step status, timing, errors of a run
/repair <run-id>
Propose a patch for a failed run
/apply <patch-id> / /reject <patch-id> [reason]
Resolve a patch
/schedule <plan> <cron> [--inputs '<json>'] / /schedules
Create / list schedules with captured inputs
/help , /clear
Help / clear chat
Running schedules in the background
The desktop app keeps running in the system tray when Keep schedules running
in the background is enabled under Settings. The option turns on
automatically when an enabled schedule exists. Use the tray menu to reopen the
window, pause or resume all schedules without changing their individual state,
or quit the process completely.
For servers and unattended machines, headless mode runs the MCP server and
the scheduler without a window:
inxm-local --headless # or: INXM_HEADLESS=1 inxm-local
Keep it running after logout with your platform's service manager, e.g.:
nohup inxm-local --headless > /tmp/inxm-headless.log 2>&1 &
or a systemd user unit ( ~/.config/systemd/user/inxm-local.service ):
[Unit]
Description =INXM // local headless scheduler
[Service]
ExecStart =%h/.local/bin/inxm-local --headless
Restart =on-failure
[Install]
WantedBy =default.target
( systemctl --user enable --now inxm-local ). On Windows, use Task Scheduler
with the same --headless argument.
Only one scheduler runs per data dir: a scheduler.lock file (holding the
owner's PID) guards against the desktop app and a headless instance firing
the same schedule twice. The second instance detects a live holder and skips
its scheduler; stale locks from crashed processes are reclaimed
automatically. Missed slots while nothing was running are not caught up, by
design.
The app starts a local Streamable-HTTP-style MCP server when the desktop
client starts. By default it listens only on loopback:
http://127.0.0.1:39387/mcp
The port is stored in settings.json and can be changed under
Settings → Local MCP server . If startup cannot bind the port (for example,
another process is already using it), the sidebar/footer and Settings view show
a warning with the bind error. Choose another port, save settings, and restart
the app.
For automation or environments where a native wi

[truncated]
