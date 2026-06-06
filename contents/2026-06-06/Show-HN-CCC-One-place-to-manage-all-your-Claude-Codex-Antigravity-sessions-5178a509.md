---
source: "https://github.com/amirfish1/claude-command-center"
hn_url: "https://news.ycombinator.com/item?id=48428005"
title: "Show HN: CCC: One place to manage all your Claude, Codex, Antigravity sessions"
article_title: "GitHub - amirfish1/claude-command-center: Local-first dashboard that orchestrates Claude Code, Codex, and Gemini CLI sessions side-by-side — spawn, resume, and review from one browser tab. · GitHub"
author: "amirfish2"
captured_at: "2026-06-06T21:27:36Z"
capture_tool: "hn-digest"
hn_id: 48428005
score: 1
comments: 0
posted_at: "2026-06-06T19:15:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CCC: One place to manage all your Claude, Codex, Antigravity sessions

- HN: [48428005](https://news.ycombinator.com/item?id=48428005)
- Source: [github.com](https://github.com/amirfish1/claude-command-center)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T19:15:48Z

## Translation

タイトル: HN を表示: CCC: すべてのクロード、コーデックス、反重力セッションを 1 か所で管理
記事のタイトル: GitHub - amirfish1/claude-command-center: Claude Code、Codex、および Gemini CLI セッションを並べて調整するローカルファーストのダッシュボード — 1 つのブラウザー タブから生成、再開、レビューを行います。 · GitHub
説明: Claude Code、Codex、および Gemini CLI セッションを並べて調整するローカルファーストのダッシュボード - 1 つのブラウザー タブから生成、再開、レビューを行います。 - amirfish1/claude-command-center

記事本文:
GitHub - amirfish1/claude-command-center: Claude Code、Codex、Gemini CLI セッションを並べて調整するローカルファーストのダッシュボードです。ブラウザーの 1 つのタブから生成、再開、レビューを行います。 · GitHub
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
アミルフィッシュ1
/
クロードコマンドセンター
公共

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
amirfish1/クロード-コマンドセンター
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
942 コミット 942 コミット .ccc .ccc .claude .claude .github .github _history_index _history_index キャッシュ キャッシュ changelog.d changelog.d docs docs フック フック infra/ telemetry-worker infra/ telemetry-worker ingesters ingesters スクリプト スクリプト スキル スキル 静的 静的テスト テスト vscode-extension vscode-extension .dockerignore .dockerignore .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE parking_LOT.md parking_LOT.md README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml run.sh run.sh server.py server.py snapshot.js snapshot.js snapshot.png snapshot.png snapshot2.js snapshot2.js snapshot2.png snapshot2.png snapshot3.js snapshot3.js snapshot3.png snapshot3.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロードが最初のものを構築している間に、次のものを開始します。
Mac 上のすべての Claude Code 、 Codex 、 Cursor 、および Antigravity セッションに 1 つのローカル ダッシュボード。並行してスポーンし、並行して出荷します。
カール -fsSL https://raw.githubusercontent.com/amirfish1/claude-command-center/main/scripts/install.sh | CCC_FROM=readme bash
自作の場合:
ブリュータップアミルフィッシュ1/ccc
醸造インストール CCC
CC
または、macOS DMG をダウンロードし、CCC.app をアプリケーションにドラッグします。
github.com/amirfish1/claude-command-center/releases/latest
まず読み取り専用のデモを試してください: ccc.amirfish.ai/demo (または amirfish1.github.io/claude-command-center/demo ) - シードされた偽のデータを含む完全なかんばん。インストールは必要ありません。
CCC がラッチします

Mac 上のすべての Claude Code、Codex、Cursor、および Antigravity セッション、つまりターミナル セッション、ヘッドレス プロセス、ダッシュボードから生成したセッション。各エージェントのディスク上の状態を真実の情報源として扱うため、何も漏れることはありません。最初のタスクがまだ構築されている間に、次のタスクを生成します。コンテキストを失うことなくプロジェクト間を切り替えることができます。複数のものを一度に発送します。
エンジンごとのファーストクラスとパーシャルの違いについては、以下のエンジン サポート マトリックスを参照してください。スポーンはサポートされているすべてのエンジンで機能し、トランスクリプトの取り込みと UX パリティは異なります。
2026-06-03 — v4.6.0 — 主要なパフォーマンス パス: CPU コアを固定する代わりにダッシュボードがアイドル状態になり、グループ チャットが最大 40 倍速く開き、長い会話がほぼ瞬時に開きます (ウィンドウの読み込み + スクロールアップでより早く読み込み)、スクリーンショットを含む Codex セッションが複数 MB の画像で停止することはなくなりました。フッターに新しい CCC セルフヘルス表示が追加されました。
2026-05-21 — v4.0.0 — Antigravity (Google DeepMind) が、Claude Code および Codex と並ぶファーストクラスのエンジンとしてダッシュボードに追加されました。
2026-05-21 — 会話行をウィンドウの外にドラッグすると、会話ごとに 24 色のアクセント カラーが表示された、フォーカスされたサイド ペインにポップされます。
2026-05-19 — static/templates.json によって駆動される、再利用可能な新しいセッション プロンプトのためのテンプレート ギャラリー メカニズム。 ( #46 )
2026-05-19 — VS Code 拡張機能 v0.1.0 が公開 — アクティブなワークスペース フォルダーからセッションを生成します。 ( #52 )
2026-05-19 — ワンコマンドカール | bash インストーラー; git clone が「ソースから」セクションに降格されました。 ( #58 )
2026-05-19 — シードされたモック データを使用した静的 GitHub Pages デモ (インストールは不要)。 ( #49 )
2026-05-18 — ローカル macOS では、会話中にテキスト読み上げボタンが表示されます。
インストールされている場合は、その方法をぜひお聞きしたいです。 ⭐ をドロップして問題をオープンしてください
何がうまくいったのか、何が壊れたのか、または単に挨拶をしてください。これは一人で行うプロジェクトです
建てられた

特定のワークフロー。外部からのフィードバックが私が知る唯一の方法です
それがどれほど広範囲に着陸するか。 @amirfish1
クロード コード オーケストレーション ツールのほとんどは独自のラッパーです。彼らはそうしたいのです
自らの処刑。それらを通じてエージェントを起動すると、その見返りとして、
ダッシュボード。そうならないまでは大丈夫です。ターミナルを開いた瞬間に、
クロード -- 何かを再開し、それを手動で反復すると、あなたは限界の外に出ます。
ツールの世界。ダッシュボードからは見えません。あなたが行った作業はそうではありません
問題に対してカンバンに記載され、レビューの列に並びます。
これは逆の方向に進みます。クロード コードのディスク上の状態を次のように扱います。
真実の情報源: ~/.claude/projects/*.jsonl トランスクリプト、
~/.claude/sessions/<pid>.json ライブ レジストリ、およびツール呼び出しごとのサイドカー
~/.claude/settings.json にインストールする 2 つのフックによって書き込まれたファイル。もし
Claude コードはマシン上のどこでも実行されており、ここに表示されます。もしあなたが
ダッシュボードを閉じても、セッションは実行され続けます。ターミナルを開いて、
手動で繰り返すと、カードが更新されます。
ダッシュボードは、ヘッドレス セッションを生成する方法も認識しています (経由)
claude -p --input-format stream-json ) および休止中のものをオンデマンドで再開します。
しかし、それらは追加的なものです。それを中心に構築されているのは、仕事に結びついていることです
それはすでに存在します。
デモをお試しください: ccc.amirfish.ai/demo — シードされた偽のデータを含む読み取り専用のかんばん。インストールは必要ありません。
要件: macOS、Python 3、および Claude Code がインストールされていること。
オプション: GitHub 統合の場合は gh、デプロイ ステータスの場合は vercel。
curl — ~/.ccc/claude-command-center にクローンを作成し、フォアグラウンドで実行します。再実行すると git pull が実行されます。
カール -fsSL https://raw.githubusercontent.com/amirfish1/claude-command-center/main/scripts/install.sh | CCC_FROM=readme bash
Homebrew — Cellar にインストールし、 ccc を PATH に置き、brew で管理される Python を固定します。 brew upgrade ccc 経由でアップグレードします。
ブリュータップアミルフィッシュ1/ccc

醸造インストール CCC
ccc # 前景
brew services start ccc # または brew 管理のバックグラウンド サービスとして実行
DMG — CCC.app をアプリケーションにドラッグし、ダブルクリックして起動します。これは内部で同じ install.sh を実行します (~/.ccc/claude-command-center にクローンを作成し、launchd サービスをインストールするかどうかを尋ねます)。最初の起動では、macOS の「正体不明の開発者」プロンプトが表示されます。CCC.app を右クリック→「開く」→「開く」の順にクリックします。必要なのは 1 回だけです。最新リリースからダウンロードします。
最初にクローンを作成してスクリプトを直接実行したい場合は、代わりにチャネルをフラグとして渡します: ./scripts/install.sh --from=readme 。
git clone https://github.com/amirfish1/claude-command-center
cd クロード コマンド センター
＃試してみてください。 Ctrl-C / ターミナルが閉じるまでフォアグラウンドで実行されます
./run.sh
＃それを保持します。今すぐログイン時に開始されるユーザーごとの launchd エージェントとしてインストールします
./run.sh --install-service
http://localhost:8090 を開き、からリポジトリを選択します。
リポジトリ スコープのアクションを開始する前に、リポジトリ ドロップダウンを確認します。
--install-service は ~/Library/LaunchAgents/com.github.claude-command-center.plist を書き込みます
そしてそれをユーザーごとの launchd ドメインに登録して、CCC がすぐに開始されるようにします。
終了すると再起動し、macOS ログイン時に再度起動します。何でも焼きます
PORT / CCC_* 環境変数は、実行時に設定されました。再実行して構成を更新します。
./run.sh --service-status で確認します。 ./run.sh --uninstall-service で削除します。
サービス ログは ~/.claude/command-center/logs/service.{out,err}.log に保存されます。
通常の CCC アプリの更新では、同じチェックアウト パスが使用され続けます。再実行
./run.sh --install-service は、焼き付けられた環境変数を変更する場合のみ、または
launchd plist 自体を変更するリリースを選択します。
最初の起動 (フォアグラウンドまたはサービス) では、2 つのフック スクリプトがコピーされます。
~/.claude/command-center/hooks/ に登録します。
~/.claude/settings.json 。その後、クロードごとに

コードセッション
マシン (ターミナル、ヘッドレス、またはダッシュボード生成) がサイドカーの状態を書き込みます
UI はカンバンに使用します。
┌───────────┐ 書きます ┌─────────────┐
│ 任意のクロード │ ─────────> │ ~/.claude/projects/*.jsonl │
│ プロセス │ │ ~/.claude/sessions/<pid>.json │
│ どこでも │ │ ~/.claude/command-center/ │
│ あなたの Mac │ │ live-state/<sid>.json │
━━━━━━━━━━━━━━━━━━━━━━━┘
│読む
v
┌───────────────┐
│server.py(stdlib)│
│ :8090 │
━───────┬───────┘
│
v
┌───────────────┐
│ static/index.html │
│ カンバン + 詳細ペイン │
━━━━━━━━━━┘
セッション : ディスク上の任意のクロード コード トランスクリプト (生存中または休止中)。
Attach : サーバーはクロード自身のファイルとサイドカーの状態を読み取ります。
インストールされたフックは、ツール呼び出しのたびに書き込みを行います。設定するものは何もありません
セッションごと。
列 : バックログ → 計画 → 作業中 → レビュー → テスト中 →
確認済み/非アクティブ/アーカイブ済み。列はセッション状態から派生します
(ライブ?コミット?プッシュ?サイドカーアクティビティ?)、ドラッグでオーバーライド可能。
バックログ: GitHub の問題と TODO.md エントリを開き、カードとして表示されます。
アクティブなセッションの隣にあるため、すべてが 1 つのボード上に存在します。
CCC は最初にクロード コードを中心に構築されました。 Codex、Cursor、Antigravity のサポートが続きました。ダッシュボードからのスポーンは 4 つすべてで機能します。残りは異なります:
カーソル IDE 統合に関する注意事項

説明: CCC は CLI 経由でカーソル エージェントをヘッドレスで生成しますが、デスクトップ IDE は、store.db 内の高度にネストされた独自の Protobuf Merkle ツリーを使用して UI 状態を内部で管理します。ワークスペースが破損する危険性が非常に高いため、IDE への完全な「双方向チャット同期」はサポートされていません。代わりに、CCC はメタデータ統合を実行します。CLI セッションはブックマーク (正しいタイトルとタイムスタンプ付き) として IDE サイドバーに挿入されるため、セッションを見失うことはありませんが、IDE ウィンドウ内でネイティブに操作することはできません。完全な履歴を確認するには、CCC ダッシュボードを使用してください。
エンジンが「部分的」からファーストクラスに向上したことを確認したい場合は、問題をオープンしてください。これは主にアダプターの作業であり、取り込み層はエンジンに依存しません。
マルチエンジン オーケストレーション : 1 つのダッシュボードから Claude Code 、 Codex 、 Cursor 、および Antigravity セッションを生成、再開、レビューします。エンジンごとのパリティについては、エンジン サポート マトリックスを参照してください。
すべてのセッションでカンバンを使用し、列間でドラッグ アンド ドロップを使用します。
ラバーバンドによる複数選択、および列ごとの色付け。
会話を分割する: サイドバーセッションを右側にドラッグするか、
開いている会話の下端で 2 つのトランスクリプトを表示します
並べて表示し、それぞれに独自の入力バーを持ちます。閉じて単一ペインに戻ります
クリックすると; 900px 以下では自動的に折りたたまれます。
GitHub の統合: ワンクリックで課題からセッションを開始
(進行中のクロードラベルを自動追加 + 自己割り当て)。確認して閉じます
commit-SHA コメントの問題。アーカイブ済みにドラッグすると「非」として閉じます
計画した」。ダッシュボード内に問題の本文とコメントがレンダリングされます (iframe はありません。
GitHub はそれをブロックします)。
既存のセッションに接続: ターミナル クロード プロセスが表示されます
自動的に。ターミナルにジャンプすると、TTY によってそれらに焦点が当てられます。名前変更/色変更
クロード独自のスラッシュ コマンドを使用してタブを開きます。
クロードで開く

[切り捨てられた]

## Original Extract

Local-first dashboard that orchestrates Claude Code, Codex, and Gemini CLI sessions side-by-side — spawn, resume, and review from one browser tab. - amirfish1/claude-command-center

GitHub - amirfish1/claude-command-center: Local-first dashboard that orchestrates Claude Code, Codex, and Gemini CLI sessions side-by-side — spawn, resume, and review from one browser tab. · GitHub
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
amirfish1
/
claude-command-center
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
amirfish1/claude-command-center
main Branches Tags Go to file Code Open more actions menu Folders and files
942 Commits 942 Commits .ccc .ccc .claude .claude .github .github _history_index _history_index cache cache changelog.d changelog.d docs docs hooks hooks infra/ telemetry-worker infra/ telemetry-worker ingesters ingesters scripts scripts skills skills static static tests tests vscode-extension vscode-extension .dockerignore .dockerignore .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE PARKING_LOT.md PARKING_LOT.md README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml run.sh run.sh server.py server.py snapshot.js snapshot.js snapshot.png snapshot.png snapshot2.js snapshot2.js snapshot2.png snapshot2.png snapshot3.js snapshot3.js snapshot3.png snapshot3.png View all files Repository files navigation
Start the next while Claude builds the first.
One local dashboard for every Claude Code , Codex , Cursor , and Antigravity session on your Mac. Spawn in parallel, ship in parallel.
curl -fsSL https://raw.githubusercontent.com/amirfish1/claude-command-center/main/scripts/install.sh | CCC_FROM=readme bash
With Homebrew:
brew tap amirfish1/ccc
brew install ccc
ccc
Or download the macOS DMG and drag CCC.app to Applications:
github.com/amirfish1/claude-command-center/releases/latest
Try the read-only demo first: ccc.amirfish.ai/demo (or amirfish1.github.io/claude-command-center/demo ) - full kanban with seeded fake data, no install required.
CCC latches onto every Claude Code, Codex, Cursor, and Antigravity session on your Mac — terminal sessions, headless processes, and sessions you spawned from the dashboard. It treats each agent's on-disk state as the source of truth, so nothing slips through. Spawn the next task while the first is still building. Switch between projects without losing context. Ship multiple things at once.
See the engine support matrix below for what's first-class vs. partial per engine — spawn works across all supported engines, transcript ingestion and UX parity vary.
2026-06-03 — v4.6.0 — Major performance pass: the dashboard idles instead of pinning a CPU core, group-chat opens ~40× faster, long conversations open near-instantly (windowed load + scroll-up to load earlier), and Codex sessions with screenshots no longer stall on multi-MB images. New CCC self-health readout in the footer.
2026-05-21 — v4.0.0 — Antigravity (Google DeepMind) joins the dashboard as a first-class engine alongside Claude Code and Codex.
2026-05-21 — Drag any conversation row outside the window to pop it into a focused side pane, with 24 per-conversation accent colors.
2026-05-19 — Template gallery mechanism for reusable new-session prompts, driven by static/templates.json . ( #46 )
2026-05-19 — VS Code extension v0.1.0 published — spawn a session from the active workspace folder. ( #52 )
2026-05-19 — One-command curl | bash installer; git clone demoted to a "From source" section. ( #58 )
2026-05-19 — Static GitHub Pages demo with seeded mock data (no install required). ( #49 )
2026-05-18 — Local macOS say text-to-speech button on conversations.
If you install it, I'd love to hear how. Drop a ⭐, open an issue with
what worked or what broke, or just say hi. This is a one-person project
built around a specific workflow. Outside feedback is the only way I know
how widely it lands. @amirfish1
Most Claude Code orchestration tools are opinionated wrappers. They want to
own execution. You launch agents through them, and in return you get a
dashboard. That's fine until it isn't. The moment you open a terminal,
claude --resume something, and iterate on it by hand, you're outside the
tool's universe. The dashboard can't see it. The work you just did doesn't
show up on the kanban, against the issue, in the review queue.
This goes the other way. It treats Claude Code's on-disk state as the
source of truth: ~/.claude/projects/*.jsonl transcripts, the
~/.claude/sessions/<pid>.json live registry, and per-tool-call sidecar
files written by two hooks we install into ~/.claude/settings.json . If
Claude Code is running anywhere on your machine, it shows up here. If you
close the dashboard, your sessions keep running. If you open a terminal and
iterate by hand, the card updates.
The dashboard also knows how to spawn headless sessions (via
claude -p --input-format stream-json ) and resume dormant ones on demand,
but those are additive. The thing it's built around is attaching to work
that already exists.
Try the demo: ccc.amirfish.ai/demo — read-only kanban with seeded fake data, no install required.
Requirements: macOS, Python 3, and Claude Code installed.
Optional: gh for GitHub integration, vercel for deploy status.
curl — clones into ~/.ccc/claude-command-center and runs in foreground. Re-running does a git pull .
curl -fsSL https://raw.githubusercontent.com/amirfish1/claude-command-center/main/scripts/install.sh | CCC_FROM=readme bash
Homebrew — installs into the Cellar, puts ccc on PATH , pins a brew-managed Python. Upgrade via brew upgrade ccc .
brew tap amirfish1/ccc
brew install ccc
ccc # foreground
brew services start ccc # or run as a brew-managed background service
DMG — drag CCC.app to Applications, double-click to launch. Under the hood this runs the same install.sh (clones into ~/.ccc/claude-command-center and asks whether to install the launchd service). First launch shows the macOS "unidentified developer" prompt — right-click CCC.app → Open → Open . Only required once. Download from the latest release .
If you'd rather clone first and run the script directly, pass the channel as a flag instead: ./scripts/install.sh --from=readme .
git clone https://github.com/amirfish1/claude-command-center
cd claude-command-center
# Try it. Runs in the foreground until Ctrl-C / terminal close
./run.sh
# Keep it. Install as a per-user launchd agent that starts now and at login
./run.sh --install-service
Open http://localhost:8090 , then pick a repo from
the repo dropdown before starting repo-scoped actions.
--install-service writes ~/Library/LaunchAgents/com.github.claude-command-center.plist
and registers it in your per-user launchd domain so CCC starts immediately,
restarts if it exits, and starts again at macOS login. It bakes in whatever
PORT / CCC_* env vars were set when you ran it. Re-run it to update config;
check with ./run.sh --service-status ; remove with ./run.sh --uninstall-service .
Service logs go to ~/.claude/command-center/logs/service.{out,err}.log .
Normal CCC app updates keep using the same checkout path; re-run
./run.sh --install-service only when you want to change baked-in env vars or
pick up a release that changes the launchd plist itself.
First launch (foreground or service) copies two hook scripts into
~/.claude/command-center/hooks/ and registers them in
~/.claude/settings.json . After that, every Claude Code session on your
machine (terminal, headless, or dashboard-spawned) writes sidecar state
the UI uses for the kanban.
┌─────────────┐ writes ┌────────────────────────────────┐
│ any claude │ ─────────> │ ~/.claude/projects/*.jsonl │
│ process │ │ ~/.claude/sessions/<pid>.json │
│ anywhere on │ │ ~/.claude/command-center/ │
│ your mac │ │ live-state/<sid>.json │
└─────────────┘ └──────────────┬─────────────────┘
│ reads
v
┌───────────────────────┐
│ server.py (stdlib) │
│ :8090 │
└───────────┬───────────┘
│
v
┌───────────────────────┐
│ static/index.html │
│ kanban + detail pane │
└───────────────────────┘
Session : any Claude Code transcript on disk, alive or dormant.
Attach : the server reads Claude's own files + sidecar state the
installed hooks write after every tool call. Nothing to configure
per-session.
Columns : Backlog → Planning → Working → Review → In Testing →
Verified / Inactive / Archived. Columns are derived from session state
(live? commits? pushed? sidecar activity?), overridable by drag.
Backlog : open GitHub issues + TODO.md entries, surfaced as cards
next to your active sessions so everything lives on one board.
CCC was built around Claude Code first; Codex, Cursor, and Antigravity support followed. Spawn-from-dashboard works for all four. The rest varies:
Note on Cursor IDE integration: While CCC spawns Cursor agents headlessly via the CLI, the Desktop IDE manages UI state internally using a highly-nested, proprietary Protobuf Merkle tree in store.db . Full "two-way chat sync" into the IDE is unsupported due to the extreme risk of workspace corruption. Instead, CCC performs a metadata integration : CLI sessions are injected into the IDE sidebar as bookmarks (with correct titles and timestamps) so you don't lose track of them, but they cannot be interacted with natively inside the IDE window. Use the CCC dashboard for full history.
If you'd like to see an engine bumped from "partial" to first-class, open an issue — it's mostly adapter work, the ingestion layer is engine-agnostic.
Multi-engine orchestration : spawn, resume, and review Claude Code , Codex , Cursor , and Antigravity sessions from one dashboard. See the engine support matrix for per-engine parity.
Kanban across every session, with drag-drop between columns,
rubber-band multi-select, and per-column tinting.
Split conversations : drag any sidebar session onto the right or
bottom edge of the open conversation to view two transcripts
side-by-side, each with its own input bar. Closes back to single-pane
with a click; collapses automatically below 900px.
GitHub integration : start a session from an issue with one click
(auto-adds claude-in-progress label + self-assigns). Verify closes the
issue with a commit-SHA comment. Drag to Archived closes as "not
planned". Issue body + comments render inside the dashboard (no iframe;
GitHub blocks that).
Attach to existing sessions : terminal claude processes show up
automatically. Jump-to-terminal focuses them by TTY; rename/color the
tab via Claude's own slash commands.
Open in Claude

[truncated]
