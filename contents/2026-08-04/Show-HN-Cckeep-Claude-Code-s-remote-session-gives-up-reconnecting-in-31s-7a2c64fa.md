---
source: "https://github.com/kamihork/cckeep"
hn_url: "https://news.ycombinator.com/item?id=49164763"
title: "Show HN: Cckeep – Claude Code's remote session gives up reconnecting in 31s"
article_title: "GitHub - kamihork/cckeep: Keeps Claude Code Remote Control from silently going dead. Re-arms the sessions its 31-second retry budget gives up on — local, tmux-native, zero dependencies. · GitHub"
author: "kamihork"
captured_at: "2026-08-04T06:25:15Z"
capture_tool: "hn-digest"
hn_id: 49164763
score: 1
comments: 0
posted_at: "2026-08-04T05:56:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cckeep – Claude Code's remote session gives up reconnecting in 31s

- HN: [49164763](https://news.ycombinator.com/item?id=49164763)
- Source: [github.com](https://github.com/kamihork/cckeep)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T05:56:11Z

## Translation

タイトル: Show HN: Cckeep – Claude Code のリモート セッションは 31 秒で再接続を断念します
記事のタイトル: GitHub - kamahork/cckeep: クロード コード リモート コントロールが静かに停止するのを防ぎます。 31 秒の再試行バジェットが放棄したセッションを再準備します。ローカル、tmux ネイティブ、依存関係はありません。 · GitHub
説明: クロード コード リモート コントロールが静かに停止するのを防ぎます。 31 秒の再試行バジェットが放棄したセッションを再準備します。ローカル、tmux ネイティブ、依存関係はありません。 - カミホーク/cckeep

記事本文:
GitHub - kamahork/cckeep: クロード コード リモート コントロールが静かに停止するのを防ぎます。 31 秒の再試行バジェットが放棄したセッションを再準備します。ローカル、tmux ネイティブ、依存関係はありません。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
カミホーク

/
保存します
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
23 コミット 23 コミット .github/ workflows .github/ workflows アセット アセット bin bin docs docs src src test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.ja.md README.ja.md README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コード リモート コントロールが静かに停止するのを防ぎます。
リモート コントロールは約 31 秒間再試行し、その後完全に諦めます。
cckeep はそれに気づき、ビジー状態のセッションには触れずにセッションを再準備します。
リモート コントロールを使用すると、携帯電話または claude.ai からローカルのクロード コード セッションを実行できます。リンクが切断されると、1/2/4/8/16 秒のバックオフで 5 回試行され、自動的に再接続されます。それは 31 秒の予算です。ラップトップの蓋を閉め、Wi-Fi を切り替え、エレベーターに乗れば、予算はなくなります。接続は閉じられ、二度と戻りません。
2 番目の障害もあります。/rc 内のセッション ウェッジが再接続し、永久にそこに留まります。それは anthropics/claude-code#34255 — 2026 年 3 月 99 日からオープン 👍、修正なし。
どちらの場合でも、わかる方法は同じです。携帯電話に手を伸ばすと、セッションが終了します。文書化された回復方法は、デスクに戻って /remote-control と入力することです。
npm install -g cckeep
cckeepを有効にする
npm install は CLI を PATH に置くだけです。 cckeep Enable は、バックグラウンド ジョブ (macOS では launchd、Linux では systemd ユーザー タイマー) を登録するステップであり、15 秒ごとにチェックして、停止したものを再準備します。
npx cckeep enable を実行するのではなく、グローバルにインストールします。スケジュールされたジョブは、インストールされた場所から cckeep を実行し、npx のキャッシュは使い捨てです: a jo

それを指す b は、キャッシュがクリアされるまで動作し続け、その後静かに停止します。これは、ウォッチドッグが発生してはならない 1 つの失敗です。そのため、cckeep Enable は npx パスからのスケジュールを拒否します。
何かをインストールする前に確認するには、npx で問題ありません。どちらも何も変わりません。
npx cckeep # 現在表示されている内容
npx cckeep Doctor # tmux、ペイン、スケジューラをチェックする
コマンドが見つかりません: インストール直後は cckeep ですか?シムベースを使用しています
バージョンマネージャー。 nodenv および rbenv スタイルのセットアップでは、新たにセットアップする前に再ハッシュが必要です
インストールされたバイナリが PATH に表示され、nvm には新しいシェルが必要です。
nodenv の再ハッシュ # nodenv
asdf reshim nodejs # asdf
# nvm: 新しいシェルを開くだけです
1 つの要件: クロード コードが tmux 内で実行されている必要があります。裸の端末で開始されたセッションには別のプロセスからアクセスできないため、ツールでできることは何もありません。 「 tmux でのクロード コードの実行」を参照してください。
cckeep によってデスクに戻る手間が省けた場合、⭐ は他のリモート コントロール ユーザーがそれを見つけるのに役立ちます。
動作しているセッションには入力されません
これが設計全体の問題です。タイマーで端末に入力するウォッチドッグは、その瞬間が安全であることが確実でない限り、責任を負うことになります。これらはすべて適用され、テストされます。
入力した内容を上書きすることはありません。 Enter はコンポーザーにあるものをすべて送信するため、未送信のドラフトはコマンドが貼り付けられた状態で送信されます。ボックスの中に何かがある場合、またはボックスがまったく見つからない場合は、何も送信されません。アイドルチェックではこれをカバーできません。ボックス内にあるドラフトは完全に静止しています。
ターン中は決して行わないでください。ペインは 3 回キャプチャされます。実行中のターンではスピナーとトークン カウンターがアニメーション化されるため、同一のキャプチャは何も起こっていないことを意味します。 2 つではなく 3 つの間隔で、意図的に丸くならない間隔にします。これは、アニメーションの周期が間隔を分割するためです。

それ以外の場合、al は同一のフレームにエイリアスを作成します。
決してダイアログに入らないでください。許可のプロンプトが表示されたら、Enter キーを押して選択します。選択マーカーは画面上のどこでもカウントされます。クロード・コードは「私にもテストを実行してほしいですか?」のような文を書くので、平易な英語のフレーズは作曲者の近くでのみ考慮されます。普通の返事で。
開いたパネルには決して入らないでください。 /remote-control は、QR コードを含むステータス パネルを開きます。 cckeep は、パネル自体を開いたときにのみ Enter キーを押します。また、パネルを押したときにパネルがまだ開いている場合にのみ実行します。
それについてのみ言及したセッションは決してありません。インジケーターはトランスクリプトからではなく、最後の数行から読み取られるため、/rc active について議論するセッションが接続されているものと誤解されることはありません。接続が一度も確認されたことのないペインは、クロード コード自体がリンクが切断されたことを出力した場合にのみ動作します。
決してタイトなループに入ることはなく、一度に 2 回行うこともありません。ペインごとに 5 分ごとに 1 つのアクションが実行され、スケジュールされたジョブと並行して実行されている cckeep 監視が 1 つの文字化けしたプロンプトに 2 つのパスをインターリーブできないようにロックされています。
決して壁に向かってはいけない。 「リモート コントロールには claude.ai サブスクリプションが必要です」のような応答は、ここでは再接続が不可能であること、つまりコマンドでは解決できない認証またはプランの問題であることを意味するため、リンクが再び正常であることが確認されるまでペインはドロップされます。また、認識できない障害モードの場合、停止ごとに最大 3 回の再準備が行われます。その後、トランスクリプトに永遠に再入力するのではなく、停止してリンクが正常であることを確認するのを待ちます。
最後の瞬間に再確認しました。決定は 1 回のキャプチャから行われ、待機後に再検証されます。その間に再接続され、ダイアログが表示され、何か入力されましたか?何も送信されず、ペインは待機をやり直すのではなく、それまでの進行状況を維持します。
--dry-run は実行内容を出力し、何も送信しません。
画面上の状態
私は何

つまり
cckeep の機能
/rc active 、または切り詰められた /rc
接続されています
ペインだけを覚えており、他には何も覚えていません
/rc 再接続中
31秒の予算内で
待機します - 通常はこれで解決します
/rc 再接続中、2 分後
ウェッジ ( #34255 )
ブリッジを循環させます: パネルを開き、切断し、再接続します
リモコンが切断されました
諦めた
すぐに再武装する
インジケーターがあったペインにインジケーターなし
通知がスクロールして消えた
4回の静かなチェックの後、再武装します
インジケーターがありません。これまでインジケーターがなかったペインにあります
あなたの設定ではありません
何も、決して
コマンド
cckeep # status: クロード コード ペインごとに 1 行
cckeep watch # スケジュールの代わりにフォアグラウンドで実行
cckeep Once # 単一パス — スケジューラが実行する内容
cckeep Enable # バックグラウンドでチェックを開始
cckeep disable # チェックを停止
cckeep ドクター # tmux、ペイン、スケジューラ、パス
cckeep logs # 何をしたか
オプション: --dry-run 、 --json 、 --interval <s> 、 --lang en|ja ( LANG から自動検出)。
install と uninstall は、enable と disable のエイリアスとして引き続き機能します。
cckeep は tmux ペインを読み取り、入力します。これは、別のプロセスがライブ クロード コード セッションにアクセスできる唯一のチャネルです。これが、セッションを tmux 内で開始する必要がある理由です。プロセスを再起動することは代替策ではありません。それは会話を終了することになり、まさにあなたが保存しようとしているものです。
最も小さな変更は、インタラクティブな起動のみをラップするシェル関数です。そのため、 claude update 、claude Doctor、および claude -p は引き続き通常どおり動作します。
cc () {
地元の
" $@ " の a の場合;する
「 $a 」の場合
-p|--print|-v|--version|-h|--help|--bg|--background|--output-format) コマンド クロード " $@ " ;戻る ;;
エージェント | 認証 | ドクター | インストール | mcp | プラグイン | プロジェクト | セットアップ トークン | 更新 | アップグレード | リモート コントロール | rc | config)
コマンド クロード " $@ " ;戻る ;;
- * ) ;;
* ） 壊す ;;
イーサック
完了しました
ローカルセッション = " クロード $( ベース名

" $PWD " ) - $( printf ' %s ' " $PWD " | cksum | Cut -d ' ' -f1 ) "
if [-n "$TMUX" ] ;次に、コマンド クロード " $@ " ;戻る;フィ
もし！ tmux has-session -t " = $session " 2> /dev/null ;それから
tmux 新しいセッション -s " $session " -c " $PWD " クロード " $@ " ;戻る
フィ
# tmux セッションは内部の claude プロセスよりも存続するため、アタッチは可能です
# 会話がなくなって、裸の殻に落ちます。取り戻してください。
local pane_cmd コマンド
pane_cmd= $( tmux list-panes -t " = $session : " -F ' #{pane_current_command} ' 2> /dev/null | head -1 )
「 $pane_cmd 」の場合
zsh|bash|sh|fish)
cmd= "クロード"
「 $* 」の場合
* " -c " * | * " --続き" * | * " -r " * | * " --resume " * | * " --session-id " * ) ;;
* ) cmd= " $cmd --Continue " ;;
イーサック
[ $# -gt 0 ] && cmd= " $cmd $* "
tmux send-keys -t " = $session : " " $cmd " C-m ;;
イーサック
tmux アタッチセッション -t " = $session "
}
cc -c および cc -- continue は期待どおりに機能します。フラグはそのまま渡されます。
クロード・コードまで。上記のブロックは、フラグができない場合をカバーしています: tmux
クロード コード プロセスがすでに終了しているセッション。
Claude Code には、 ~/.tmux.conf に 2 行を追加するか、Shift+Enter を押してデスクトップ通知を tmux 内で中断する必要があります (公式ガイダンス)。
set -g パススルーを許可する
-s 拡張キーをオンに設定します
set -as ターミナル機能 ' xterm*:extkeys '
Ctrl+B は修正の必要がありません。Claude コードは tmux を検出し、独自のショートカットを Ctrl+B Ctrl+B に再バインドします。
デフォルトは調整されているため、気にすることはありません。 ~/.cckeep/config.json で、環境変数または実行ごとのフラグによってオーバーライドします。後で優先されます。不正な構成は、サイレントな半ロードではなく、ハード エラーです。
{
「間隔」 : 15 、
「クールダウン」：300、
"スタックリミット" : 8 、
"ミスリミット" : 4 、
「決済」 : 2000 、
"paneCommand" : " クロード "
}
間隔 — パス間の秒数 (スケジュールを有効にするものでもあります)
クールダウン — se

同じペインの前の conds が再度実行される可能性があります
stackLimit — ブリッジがウェッジとして扱われる前に再接続をチェックします。
missLimit — インジケーターがあるペインを再装備する前に、インジケーターがないことを確認します。
maxRearms — リンクが再び正常であると見なされるまで、停止する前に停止するたびに許可される再アーム数
Settle — アイドル チェックの 2 つのキャプチャ間のミリ秒。遅いマシンで上げる
paneCommand — ペインをクロード コードとしてマークするフォアグラウンド プロセス名
tmuxSocket — tmux がデフォルトのサーバー ( tmux -L name / -S path ) 以外で実行されている場合のソケット名またはパス。空はデフォルトを意味します
tmuxBinary — tmux への絶対パス。あなたのパスがどこかにある場合、通常の検索は失敗します。
すべてのキーには環境ツインがあります: CCKEEP_INTERVAL 、 CCKEEP_COOLDOWN 、 CCKEEP_STUCK_LIMIT 、 CCKEEP_MISS_LIMIT 、 CCKEEP_MAX_REARMS 、 CCKEEP_SETTLE 、 CCKEEP_PANE_COMMAND 、 CCKEEP_TMUX_SOCKET 、 CCKEEP_TMUX 。 cckeep enable の実行時に設定されたものはすべてスケジュールされたジョブに書き込まれるため、シェル内でのみ設定されたソケットがバックグラウンド実行から静かに失われることはありません。 CCKEEP_HOME は、状態、構成を移動し、 ~/.cckeep からログオフします。
cckeep は接続を再準備します。クロード コードの再試行バジェットは増加しません。これはクローズド ソース バイナリ内の定数であり、Anthropic のみが変更できます。

[切り捨てられた]

## Original Extract

Keeps Claude Code Remote Control from silently going dead. Re-arms the sessions its 31-second retry budget gives up on — local, tmux-native, zero dependencies. - kamihork/cckeep

GitHub - kamihork/cckeep: Keeps Claude Code Remote Control from silently going dead. Re-arms the sessions its 31-second retry budget gives up on — local, tmux-native, zero dependencies. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
kamihork
/
cckeep
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
23 Commits 23 Commits .github/ workflows .github/ workflows assets assets bin bin docs docs src src test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.ja.md README.ja.md README.md README.md package.json package.json View all files Repository files navigation
Keeps Claude Code Remote Control from silently going dead.
Remote Control retries for about 31 seconds and then gives up for good.
cckeep notices, and re-arms the session — without touching one that's busy.
Remote Control lets you drive a local Claude Code session from your phone or from claude.ai. It reconnects on its own when the link drops — for 5 attempts with 1/2/4/8/16-second backoff . That is a 31-second budget. Close your laptop lid, switch Wi-Fi, ride an elevator, and the budget is gone. The connection closes and never comes back.
There is a second failure too: the session wedges in /rc reconnecting and sits there forever. That one is anthropics/claude-code#34255 — open since March 2026, 99 👍, no fix.
Either way you find out the same way: you reach for your phone, and the session is gone. The documented recovery is to walk back to your desk and type /remote-control .
npm install -g cckeep
cckeep enable
npm install only puts the CLI on your PATH ; cckeep enable is the step that registers a background job — launchd on macOS, a systemd user timer on Linux — that checks every 15 seconds and re-arms whatever went dead.
Install it globally rather than running npx cckeep enable . The scheduled job runs cckeep from wherever it was installed, and npx's cache is throwaway: a job pointing into it keeps working until the cache is cleared and then stops, silently — the one failure a watchdog must not have. cckeep enable refuses to schedule from an npx path for that reason.
To look before installing anything, npx is fine — neither of these changes a thing:
npx cckeep # what it sees right now
npx cckeep doctor # check tmux, panes, and the scheduler
command not found: cckeep right after installing? You are on a shim-based
version manager. nodenv and rbenv -style setups need a rehash before a newly
installed binary appears on PATH , and nvm needs a new shell:
nodenv rehash # nodenv
asdf reshim nodejs # asdf
# nvm: just open a new shell
One requirement: Claude Code has to be running inside tmux. A session started in a bare terminal cannot be reached from another process, so there is nothing any tool can do for it. See Running Claude Code in tmux .
If cckeep saved you a walk back to your desk, a ⭐ helps other Remote Control users find it.
It will not type into a session that is working
This is the whole design problem. A watchdog that types into your terminal on a timer is a liability unless it is certain the moment is safe. Every one of these is enforced, and tested :
Never on top of what you typed. Enter submits whatever is in the composer, so an unsent draft would go out with the command glued onto it. If anything is in the box — or the box cannot be found at all — nothing is sent. The idle check cannot cover this: a draft sitting in the box is perfectly still.
Never during a turn. The pane is captured three times. A running turn animates a spinner and a token counter, so identical captures mean nothing is happening. Three rather than two, at an interval that is deliberately not round, because any animation whose period divides the interval would otherwise alias into identical frames.
Never into a dialog. Permission prompts turn Enter into a selection. A selection marker counts anywhere on screen; the plain English phrasings only count near the composer, since Claude Code writes sentences like "Do you want me to run the tests as well?" in ordinary replies.
Never into the panel you opened. /remote-control opens a status panel with a QR code. cckeep only presses Enter there when it opened the panel itself, and only if the panel is still up when it goes to press it.
Never a session that only mentioned it. The indicators are read from the last few rows, never from the transcript, so a session discussing /rc active is not mistaken for a connected one. A pane that has never been seen connected is only ever acted on when Claude Code itself prints that the link died.
Never in a tight loop, and never twice at once. One action per pane per 5 minutes, and a lock so that cckeep watch running alongside the scheduled job cannot interleave two passes into one garbled prompt.
Never against a wall. Replies like “Remote Control requires a claude.ai subscription” mean reconnecting is impossible here — an auth or plan problem the command cannot fix — so the pane is dropped until the link is seen healthy again. And for failure modes it cannot recognise, at most 3 re-arms per outage; after that it stops and waits to see the link healthy rather than retyping into your transcript forever.
Re-checked at the last moment. The decision is made from one capture, then re-verified after the wait — reconnected in between, dialog appeared, something typed? Nothing is sent, and the pane keeps the progress it had made rather than starting its wait over.
--dry-run prints what it would do and sends nothing.
State on screen
What it means
What cckeep does
/rc active , or a truncated /rc
connected
remembers the pane, nothing else
/rc reconnecting
inside the 31-second budget
waits — this usually resolves
/rc reconnecting , 2 minutes on
wedged ( #34255 )
cycles the bridge: opens the panel, disconnects, reconnects
Remote Control disconnected
gave up
re-arms immediately
no indicator, on a pane that had one
notification scrolled away
re-arms after 4 quiet checks
no indicator, on a pane that never had one
not your setup
nothing, ever
Commands
cckeep # status: one line per Claude Code pane
cckeep watch # run in the foreground instead of scheduling
cckeep once # a single pass — what the scheduler runs
cckeep enable # start checking in the background
cckeep disable # stop checking
cckeep doctor # tmux, panes, scheduler, paths
cckeep logs # what it has done
Options: --dry-run , --json , --interval <s> , --lang en|ja (auto-detected from LANG ).
install and uninstall still work as aliases for enable and disable .
cckeep reads and types into tmux panes. That is the only channel a separate process has into a live Claude Code session — and it is why the session must be started inside tmux. Restarting the process is not an alternative: it would end the conversation, which is exactly what you are trying to save.
The smallest change is a shell function that wraps interactive launches only, so claude update , claude doctor and claude -p still behave normally:
cc () {
local a
for a in " $@ " ; do
case " $a " in
-p|--print|-v|--version|-h|--help|--bg|--background|--output-format) command claude " $@ " ; return ;;
agents|auth|doctor|install|mcp|plugin|project|setup-token|update|upgrade|remote-control|rc|config)
command claude " $@ " ; return ;;
- * ) ;;
* ) break ;;
esac
done
local session= " claude- $( basename " $PWD " ) - $( printf ' %s ' " $PWD " | cksum | cut -d ' ' -f1 ) "
if [ -n " $TMUX " ] ; then command claude " $@ " ; return ; fi
if ! tmux has-session -t " = $session " 2> /dev/null ; then
tmux new-session -s " $session " -c " $PWD " claude " $@ " ; return
fi
# The tmux session outlives the claude process inside it, so attaching can
# drop you at a bare shell with the conversation gone. Bring it back.
local pane_cmd cmd
pane_cmd= $( tmux list-panes -t " = $session : " -F ' #{pane_current_command} ' 2> /dev/null | head -1 )
case " $pane_cmd " in
zsh|bash|sh|fish)
cmd= " claude "
case " $* " in
* " -c " * | * " --continue " * | * " -r " * | * " --resume " * | * " --session-id " * ) ;;
* ) cmd= " $cmd --continue " ;;
esac
[ $# -gt 0 ] && cmd= " $cmd $* "
tmux send-keys -t " = $session : " " $cmd " C-m ;;
esac
tmux attach-session -t " = $session "
}
cc -c and cc --continue work as you'd expect — the flag is passed straight
through to Claude Code. The block above covers the case the flag can't: a tmux
session whose Claude Code process has already exited.
Claude Code also needs two lines in ~/.tmux.conf , or Shift+Enter and desktop notifications break inside tmux ( official guidance ):
set -g allow-passthrough on
set -s extended-keys on
set -as terminal-features ' xterm*:extkeys '
Ctrl+B needs no fix: Claude Code detects tmux and rebinds its own shortcut to Ctrl+B Ctrl+B .
Defaults are tuned so you never notice it. Override in ~/.cckeep/config.json , by environment variable, or per-run flag — later wins. A malformed config is a hard error rather than a silent half-load.
{
"interval" : 15 ,
"cooldown" : 300 ,
"stuckLimit" : 8 ,
"missLimit" : 4 ,
"settle" : 2000 ,
"paneCommand" : " claude "
}
interval — seconds between passes (also what enable schedules)
cooldown — seconds before the same pane may be acted on again
stuckLimit — checks in reconnecting before the bridge is treated as wedged
missLimit — checks with no indicator before re-arming a pane that had one
maxRearms — re-arms allowed per outage before it gives up until the link is seen healthy again
settle — milliseconds between the two captures of the idle check; raise it on a slow machine
paneCommand — foreground process name that marks a pane as Claude Code
tmuxSocket — socket name or path, if your tmux runs on something other than the default server ( tmux -L name / -S path ). Empty means the default
tmuxBinary — absolute path to tmux, if yours lives somewhere the usual lookup misses
Every key has an env twin: CCKEEP_INTERVAL , CCKEEP_COOLDOWN , CCKEEP_STUCK_LIMIT , CCKEEP_MISS_LIMIT , CCKEEP_MAX_REARMS , CCKEEP_SETTLE , CCKEEP_PANE_COMMAND , CCKEEP_TMUX_SOCKET , CCKEEP_TMUX . Whatever is set when you run cckeep enable is written into the scheduled job, so a socket set only in your shell does not quietly go missing from the background run. CCKEEP_HOME moves state, config and log off ~/.cckeep .
cckeep re-arms the connection. It does not raise Claude Code's retry budget — that is a constant inside a closed-source binary, and only Anthropic can chang

[truncated]
