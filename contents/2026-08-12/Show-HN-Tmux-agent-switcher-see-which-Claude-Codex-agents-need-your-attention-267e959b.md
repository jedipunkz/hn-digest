---
source: "https://github.com/Ymirke/tmux-agent-switcher"
hn_url: "https://news.ycombinator.com/item?id=49278395"
title: "Show HN: Tmux-agent-switcher: see which Claude/Codex agents need your attention"
article_title: "GitHub - Ymirke/tmux-agent-switcher: Tmux sidebar plugin for managing agents · GitHub"
author: "ymir_e"
captured_at: "2026-08-12T21:35:46Z"
capture_tool: "hn-digest"
hn_id: 49278395
score: 1
comments: 0
posted_at: "2026-08-12T20:53:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tmux-agent-switcher: see which Claude/Codex agents need your attention

- HN: [49278395](https://news.ycombinator.com/item?id=49278395)
- Source: [github.com](https://github.com/Ymirke/tmux-agent-switcher)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T20:53:14Z

## Translation

タイトル: HN を表示: Tmux-agent-switcher: どのクロード/コーデックス エージェントが注意を必要としているかを確認します
記事のタイトル: GitHub - Ymirke/tmux-agent-switcher: エージェントを管理するための Tmux サイドバー プラグイン · GitHub
説明: エージェントを管理するための Tmux サイドバー プラグイン。 GitHub でアカウントを作成して、Ymirke/tmux-agent-switcher の開発に貢献してください。
HN テキスト: 今年の初めから、私はほぼ独占的に複数の AI エージェントを並行して実行することによってコーディングしてきました。それらを管理するために多くのツールを試しましたが、私は tmux に戻ってきました。これは、ラップトップを閉じたり、Wi-Fi のない飛行機に搭乗した場合でも長時間実行されるタスクを継続したりできることが好きなので、別のマシンで実行したいためです。そこで私は、サーバー上のプロセスを管理するためのコア技術として tmux を置き換えることなく、それらすべての概要をより良く把握できる tmux プラグインである tmux-agent-switcher を構築しました。 Ctrl+n を押すとサイドバーが開き、すべてのセッションのウィンドウが表示されます。サイドバーには、エージェントの状態がアイコンとして表示されます。アイドル セッションを選択するとチェックマークも付けられるため、それを見たかどうかがわかります。 Vim と「数値ベース」のナビゲーション スタイルも備えています。私はエージェントの起動方法を管理するラッパーを作成したくなかったので、今でも通常どおり claude または codex を実行しています。プラグインは、tmux メタデータと目に見える端末出力を受動的に調べて、エージェントの状態を推測します。このアプローチには、UI の変更やプロセス名の変更により機能しなくなる可能性があるという欠点があります。他のアプローチを検討するために、いつか再訪することになるでしょうが、今のところ大きな問題は発生していません。最近別の HN ディスカッションでこのプロジェクトについて言及しましたが、これは私が使用してきたツールを一緒にハッキングしただけだったため、25 人がこのプロジェクトにスターを付けたことには少し驚きました。とにかく、私は適切な HN を表示しようと考えました: https://github.com/Ymirk

e/tmux-agent-switcher ここにいる他の人が別のマシンまたはクラウドでエージェントを実行するために何を使用しているのか知りたいです。

記事本文:
GitHub - Ymirke/tmux-agent-switcher: エージェントを管理するための Tmux サイドバー プラグイン · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ユミルケ
/
tmux エージェント スイッチャー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .github/ workflows .github/ workflows アセット アセット bin bin 例 例 src src テスト テスト .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md

README.md tmux-agent-switcher.tmux tmux-agent-switcher.tmux すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ウィンドウを切り替えて、ウィンドウを監視できる tmux サイドバー
AI コーディング エージェント (Claude Code、Codex、OpenCode) を 1 つの全画面から実行する
ポップアップ。
Ctrl + n を押すと、すべてのウィンドウのリストが表示されます。
左側にはセッション、右側には選択したウィンドウのライブプレビューが表示されます。
エージェントを実行しているペインのステータス バッジ: Working 、 Blocked
(待機中)、またはアイドル (完了)。あなたを必要としているエージェントに直接連絡してください。
検出は完全に受動的です (仕組みを参照)。プラグイン
エージェントをラップしたり、シム化したり、起動したりしないでください。 clude 、 codex 、または
いつものようにオープンコード。サイドバーには tmux とプロセス テーブルが表示されます。
ライブ、スケーリングされた全画面ポップアップのクロスセッション ウィンドウ スイッチャー
ハイライト表示されたウィンドウのプレビュー。
エージェント監視: Claude Code、Codex、または OpenCode を実行している各ペインは、
タグ付き 動作中/ブロック中/アイドル中、実行タイマー付きなので一目でわかります
どのエージェントが入力を待っているか。
タブのステータス インジケータ: 同じロールアップされたエージェントの状態が各タブに追加されます。
既存のタブ形式を置き換えることなく、tmux ウィンドウ タブを作成できます。
Vim 対応ナビゲーション: Ctrl + h/j/k/l ペイン間を移動、
ウィンドウとセッションは存在しますが、フォーカスされている場合は Vim/Neovim に渡されます。
ベビーシッターをするデーモンがありません : 軽量のバックグラウンド ポーラーがそれ自体を開始します。
初めてサイドバーを開いたときにエージェントの状態を最新の状態に保ちます。
tmux ≥ 3.3 (display-popup -B -e の場合)
bash および ps (macOS および Linux に存在)
tmux ペイン内で実行される 1 つ以上のエージェント ( claude 、 codex 、 opencode )。
検出に必要なのはこれだけです。好きなように起動してください。
ソースからビルドするには: Rust ツールチェーン (事前にビルドされたバイナリがない場合にのみ必要)
あなたのプラットフォーム用に存在します。以下を参照してください)。
~/.tmux.conf (または ~/.config/t) に追加

mux/tmux.conf ):
set -g @plugin ' Ymirke/tmux-agent-switcher '
次に、接頭語 + I を押してインストールします。初めてプラグインを使用するとき
プラットフォーム用に事前にビルドされたバイナリをダウンロードします (次の場合はソース ビルドにフォールバックします)
Rust があり、事前構築済みの存在はありません)。
git clone https://github.com/Ymirke/tmux-agent-switcher \
~ /.tmux/plugins/tmux-agent-switcher
tmux 構成に追加してリロードします。
run-shell " ~/.tmux/plugins/tmux-agent-switcher/tmux-agent-switcher.tmux "
手動キーバインド スニペットについては、examples/tmux.conf を参照してください。
カーゴインストール --git https://github.com/Ymirke/tmux-agent-switcher
使用法
キー
アクション
Ctrl + n
サイドバーを開く
Ctrl + j / Ctrl + k
次/前のセッションに切り替えます (フォーカスされている場合は Vim にパススルーします)
Ctrl + h / Ctrl + l
ペインを移動/前/次のウィンドウに折り返す (フォーカスされている場合は Vim にパススルー)
スイッチャー内部では、Vim モードがデフォルトです。裸の j/k が動きます
ハイライトすると、10j / 10k などのカウントがすぐに開きます
相対ウィンドウ。ウィンドウの行には、Vim スタイルの相対番号が表示されます: 選択された
window は 0 で、他のすべての数値はその上下の距離です。
それ。カウントを入力すると、一致する相対番号プレフィックスが強調表示されます
薄暗い j または k は、どのターゲットがまだ利用可能であるかを示します。あ
どのウィンドウにも一致しないプレフィックスは自動的にキャンセルされるため、次の桁が新たに始まります。
Tab は Vim 、数値、および検索モードを切り替えます。を維持する
数値ショートカットが分離されているということは、10j の最初の桁が決して存在しないことを意味します
セッション1と間違えました。
数値モードでは、セッションとそのウィンドウに 1 からラベルが付けられます。 2 を押し、
次に 5 を入力すると、セッション 2 の 5 番目のウィンドウが開きます。カンマは残ります。
オプションで、Enter キーを押すと、あいまいな複数桁のウィンドウ プレフィックスがコミットされます。
検索モードはウィンドウを曖昧にフィルターし (telescope.nvim スタイル)、セッションに一致します。
名前、ウィンドウ名、プロセス、ディレクトリ。 Esc でフィルターをクリアします。
その後閉じます。
↑ / ↓ 選択移動

プレビュー用のイオン、
Ctrl + j / k で次/前のウィンドウが開きます
Enter を押すと、選択したウィンドウにジャンプします。
Vim または Numbers モードでは、r を押すと、次の情報が事前に入力された編集可能なプロンプトが開きます。
選択したウィンドウの現在の名前。 Enter を押すと名前の変更が適用され、
Esc でキャンセルします。矢印キー、Home/End、Backspace、および Delete を使用して編集します。
フィールド。
セッションエッジ間の H / L 移動: セッションの最初と最後
現在のセッション、次に前または次のセッションの最初と最後。
Shift + J / K は、選択したセッションを下に入れ替えます。または
リストの上位にあります。カスタム オーダーは、tmux サーバーの存続期間中継続します。
Alt + j / k (または Alt + ↓ / ↑ )
行を移動するように、選択したウィンドウをセッション内で上下に移動します
エディターで。移動は swap-window を通じて実行されるため、 tmux は新しいファイルを保持します。
注文自体。
Shift + Tab でビューを循環します: 左側にドッキングされたサイドバー、
右にドッキングされたサイドバー-right 、またはそのすぐ上にあるパレット
画面の中央に、選択したウィンドウが後ろに全画面でプレビューされます
それ。
どちらのトグルも、tmux サーバーの存続期間中ずっと維持されます。
Ctrl + t / Ctrl + s 新規作成
任意のモードのウィンドウ/セッション。プレス ？完全なショートカット リストについては、
プラグインをロードする前にこれらを設定します。
set -g @agent_switcher_key ' C-n ' # サイドバーを開くキー (デフォルト: C-n)
set -g @agent_switcher_nav ' on ' # vim 対応 C-h/C-j/C-k/C-l nav (デフォルト: on)
set -g @agent_switcher_view ' サイドバー ' # 'サイドバー' (左)、'サイドバー右' (右)、または 'パレット' (フローティング)
set -g @agent_switcher_input 'keys' # 'keys' (デフォルト)、'numbers'、または 'search'
set -g @agent_switcher_tab_status 'on' # tmux ウィンドウのタブにエージェントの状態を表示
すでにバインドしている場合は、@agent_switcher_nav を「オフ」に設定します
Ctrl + h/j/k/l を自分で押すか、これらのキーを保持しておきます。
@agent_switcher_tab_status を「off」に設定して、tmux のウィンドウ ステータス形式をそのままにします
手付かずの。タブインジケーターが更新を開始します

サイドバーのステータスが開始された後の状態
初めてのデーモン。
プラグインは監視します。決して傍受しません。エージェントの状態は次のものから取得されます。
tmux ペインのメタデータ: pane_current_command および OSC ペインのタイトル、
tmux Capture-pane : 表示される画面テキスト、および
エージェントをペインに割り当てる ps プロセス ツリー スナップショット。
作業状況は、エージェントの画面上のアクティビティ インジケーターから推測されます。
入力待ちのプロンプト/選択がブロックされ、アクティビティが 1 回アイドル状態になります
安定します (デバウンスを使用するため、単一の浮遊サンプルが誤って「完了」をフラッシュすることはできません)。
ラッパー、シム、PID ファイル、FIFO、LD_PRELOAD、またはログ スクレイピングはありません
エージェントの周りで。
検出はエージェントの画面上の出力を読み取るため、ヒューリスティックであり、
バージョン依存 : Claude/Codex/OpenCode UI の変更、カスタム テーマ、または
英語以外のロケールでは、州の分類が正しく行われない可能性があります。それはベストエフォート型であり、
エージェント CLI の進化に伴い、時々メンテナンスが必要になることが予想されます。
カーゴテスト # 単体テスト
カーゴビルド --release
bin/tmux-agent-switcher ランチャーは既存のリリース バイナリを優先します。
それ以外の場合は事前に構築されたものをダウンロードし、それ以外の場合はソースからビルドします。編集
src/ を実行し、サイドバーを再度開くと変更が反映されます。
信頼できる最初のリリースへのパスを追跡します。チェックした項目はこれで完了です
リポジトリ;チェックされていないものには、公開リリースまたは外部アカウントが必要です。
出荷可能な製品 (Rust クレート + ランチャー + tmux エントリ ポイント) までの範囲。個人のドットファイルと置き換えられたスクリプトを除外します
テストフィクスチャ内のハードコードされた個人パスを無効化する
tmux-version ガードと構成可能なキーを備えた TPM エントリ ポイント ( tmux-agent-switcher.tmux )
ランチャーは事前に構築されたバイナリを優先し、最初の実行時にダウンロードし、カーゴビルドにフォールバックします
リリース ワークフローはプラットフォームごとのバイナリ (macOS arm64/x86_64、Linux musl x86_64/aarch64) をビルドし、GitHub リリースに添付します。
読む

ME のインストール、使用法、要件、検出に関する警告
@agent_switcher_* オプションでキーバインドをオーバーライド可能
CI は macOS + Linux でカーゴ テストを実行します
CI のステータス デーモンのヘッドレス スモーク テスト
v0.1.0 にタグを付け、Rust ツールチェーンのないクリーンなマシン (macOS arm64 と Linux の両方) でプレフィックス + I を確認します。
crates.io に公開し、代替インストール チャネルとして Homebrew タップを追加します
検出パターン (エージェント名、動作/ブロックされたフレーズ) を他のロケール / エージェント UI バージョンの構成として公開します
エージェントを管理するための Tmux サイドバー プラグイン
Readme MIT ライセンス アクティビティ スター
2 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Tmux sidebar plugin for managing agents. Contribute to Ymirke/tmux-agent-switcher development by creating an account on GitHub.

Since the start of this year I've coded pretty much exclusively by running multiple AI agents in parallel. Having tried a bunch of tools for managing them, I kept coming back to tmux. This is becaues I want to run them on a separate machine as I like being able to close my laptop, or even have long running tasks continue even if I board a flight without WiFi. So I built tmux-agent-switcher, a tmux plugin that gives me a better overview of all of them without replacing tmux as the core tech for managing processes on the server. Ctrl+n opens a sidebar showing windows across all sessions. The sidebar shows the agent's state as icons. After selecting an Idle session it also checkmarks it, so you get indicator if you've seen it or not. It also has Vim and "number-based" navigation styles. I didn't want to make a wrapper that would own how agents are launched, so I still run claude or codex normally. The plugin passively looks at tmux metadata, and visible terminal output to infer the agent state. There are downsides with this approach as UI changes or process name changes can break it. I will likely revisit it at some point to look at other approaches, but so far I haven't had any significant issues with it. I mentioned the project in another HN discussion recently and was a bit surprised 25 people starred it as it was just some tool I've been using that I hacked together. Regardless I figured I'd do a proper Show HN: https://github.com/Ymirke/tmux-agent-switcher I'm curious to hear what other people on here are using to run agents in the on a separate machine / in the cloud.

GitHub - Ymirke/tmux-agent-switcher: Tmux sidebar plugin for managing agents · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Ymirke
/
tmux-agent-switcher
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .github/ workflows .github/ workflows assets assets bin bin examples examples src src tests tests .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md tmux-agent-switcher.tmux tmux-agent-switcher.tmux View all files Repository files navigation
A tmux sidebar that lets you switch between windows and keep an eye on your
running AI coding agents (Claude Code, Codex, and OpenCode) from one full-screen
popup.
Press Ctrl + n and you get a list of every window across all
your sessions on the left, a live preview of the selected window on the right,
and a status badge on any pane running an agent: Working , Blocked
(waiting on you), or Idle (done). Jump straight to the agent that needs you.
Detection is fully passive (see How it works ). The plugin
never wraps, shims, or launches your agents. You run claude , codex , or
opencode as usual; the sidebar reads tmux and the process table.
Cross-session window switcher in a full-screen popup with a live, scaled
preview of the highlighted window.
Agent monitoring : each pane running Claude Code, Codex, or OpenCode is
tagged Working / Blocked / Idle, with a run timer, so you can see at a glance
which agent is waiting on input.
Tab status indicators : the same rolled-up agent state is appended to each
tmux window tab without replacing your existing tab format.
Vim-aware navigation : Ctrl + h/j/k/l move between panes,
windows, and sessions but pass through to Vim/Neovim when it's focused.
No daemon to babysit : a lightweight background poller starts itself the
first time you open the sidebar and keeps agent state fresh.
tmux ≥ 3.3 (for display-popup -B -e )
bash and ps (present on macOS and Linux)
One or more agents ( claude , codex , opencode ) running inside tmux panes.
That's all detection needs; launch them however you like.
To build from source: a Rust toolchain (only needed if no prebuilt binary
exists for your platform; see below).
Add to ~/.tmux.conf (or ~/.config/tmux/tmux.conf ):
set -g @plugin ' Ymirke/tmux-agent-switcher '
Then press prefix + I to install. On first use the plugin
downloads a prebuilt binary for your platform (falling back to a source build if
you have Rust and no prebuilt exists).
git clone https://github.com/Ymirke/tmux-agent-switcher \
~ /.tmux/plugins/tmux-agent-switcher
Add to your tmux config and reload:
run-shell " ~/.tmux/plugins/tmux-agent-switcher/tmux-agent-switcher.tmux "
See examples/tmux.conf for a manual keybinding snippet.
cargo install --git https://github.com/Ymirke/tmux-agent-switcher
Usage
Key
Action
Ctrl + n
Open the sidebar
Ctrl + j / Ctrl + k
Switch to the next / previous session (pass through to Vim if focused)
Ctrl + h / Ctrl + l
Move panes / wrap to prev/next window (pass through to Vim if focused)
Inside the switcher, Vim mode is the default. Bare j/k moves the
highlight, and a count such as 10j / 10k immediately opens
the relative window. Window rows show Vim-style relative numbers: the selected
window is 0 , and every other number is its distance above or below
it. After you type a count, matching relative-number prefixes are highlighted
and a dim j or k shows which targets remain available. A
prefix that matches no window cancels itself, so the next digit starts fresh.
Tab cycles Vim , numbers , and search modes. Keeping the
numeric shortcuts separate means the first digit of 10j is never
mistaken for session 1.
Numbers mode labels sessions and their windows from 1. Press 2 ,
then 5 to open the fifth window in session 2. A comma remains
optional, and Enter commits an ambiguous multi-digit window prefix.
Search mode filters windows fuzzily (telescope.nvim style), matching session
name, window name, process, and directory. Esc clears the filter,
then closes.
↑ / ↓ move the selection for previewing,
Ctrl + j / k opens the next/previous window
directly, and Enter jumps to the selected window.
In Vim or Numbers mode, r opens an editable prompt prefilled with
the selected window's current name; Enter applies the rename and
Esc cancels. Arrow keys, Home/End, Backspace, and Delete edit the
field.
H / L move between session edges: first and last in the
current session, then first and last in the previous or next session.
Shift + J / K swaps the selected session down or
up in the list. The custom order lasts for the tmux server's lifetime.
Alt + j / k (or Alt + ↓ / ↑ )
moves the selected window down or up within its session, like moving a line
in an editor. The move runs through swap-window , so tmux keeps the new
order itself.
Shift + Tab cycles the view: the left-docked sidebar ,
a right-docked sidebar-right , or a palette floating just above the
middle of the screen with the selected window previewed full-screen behind
it.
Both toggles stick for the rest of the tmux server's lifetime.
Ctrl + t / Ctrl + s create a new
window / session in any mode. Press ? for the full shortcut list.
Set these before the plugin loads:
set -g @agent_switcher_key ' C-n ' # key that opens the sidebar (default: C-n)
set -g @agent_switcher_nav ' on ' # vim-aware C-h/C-j/C-k/C-l nav (default: on)
set -g @agent_switcher_view ' sidebar ' # 'sidebar' (left), 'sidebar-right' (right) or 'palette' (floating)
set -g @agent_switcher_input ' keys ' # 'keys' (default), 'numbers', or 'search'
set -g @agent_switcher_tab_status ' on ' # show agent state in tmux window tabs
Set @agent_switcher_nav 'off' if you already bind
Ctrl + h/j/k/l yourself or want to keep those keys.
Set @agent_switcher_tab_status 'off' to leave tmux's window status formats
untouched. Tab indicators begin updating after the sidebar starts its status
daemon for the first time.
The plugin observes; it never intercepts. Agent state comes from:
tmux pane metadata: pane_current_command and the OSC pane title,
tmux capture-pane : the visible screen text, and
a ps process-tree snapshot that attributes agents to panes.
Working is inferred from the agent's on-screen activity indicators,
Blocked from a prompt/selection awaiting input, and Idle once activity
settles (with debouncing so a single stray sample can't flash a false "done").
There are no wrappers, shims, PID files, FIFOs, LD_PRELOAD , or log scraping
around your agents.
Because detection reads the agents' on-screen output, it is heuristic and
version-sensitive : a Claude/Codex/OpenCode UI change, a custom theme, or a
non-English locale can throw off state classification. It's best-effort and
expected to need occasional upkeep as the agent CLIs evolve.
cargo test # unit tests
cargo build --release
The bin/tmux-agent-switcher launcher prefers an existing release binary,
otherwise downloads a prebuilt one, otherwise builds from source. Editing
src/ and reopening the sidebar picks up your changes.
Tracking the path to a trustworthy first release. Checked items are done in this
repo; unchecked ones need a published release or external accounts.
Scope to the shippable product (Rust crate + launchers + tmux entry point); leave personal dotfiles and superseded scripts out
Neutralize hardcoded personal paths in test fixtures
TPM entry point ( tmux-agent-switcher.tmux ) with a tmux-version guard and configurable keys
Launcher prefers a prebuilt binary, downloads on first run, falls back to cargo build
Release workflow builds per-platform binaries (macOS arm64/x86_64, Linux musl x86_64/aarch64) and attaches them to GitHub Releases
README with install, usage, requirements, and the detection caveat
Keybindings overridable via @agent_switcher_* options
CI runs cargo test on macOS + Linux
Headless smoke test of the status daemon in CI
Tag v0.1.0 and verify prefix + I on a clean machine with no Rust toolchain , on both macOS arm64 and Linux
Publish to crates.io and add a Homebrew tap as alternate install channels
Expose detection patterns (agent names, Working/Blocked phrases) as config for other locales / agent-UI versions
Tmux sidebar plugin for managing agents
Readme MIT license Activity Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
