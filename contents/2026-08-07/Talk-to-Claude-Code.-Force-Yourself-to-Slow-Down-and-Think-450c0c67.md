---
source: "https://github.com/tmshapland/audiochatty-plugin"
hn_url: "https://news.ycombinator.com/item?id=49214118"
title: "Talk to Claude Code. Force Yourself to Slow Down and Think"
article_title: "GitHub - tmshapland/audiochatty-plugin · GitHub"
author: "tmshapland"
captured_at: "2026-08-07T18:40:31Z"
capture_tool: "hn-digest"
hn_id: 49214118
score: 2
comments: 1
posted_at: "2026-08-07T18:03:04Z"
tags:
  - hacker-news
  - translated
---

# Talk to Claude Code. Force Yourself to Slow Down and Think

- HN: [49214118](https://news.ycombinator.com/item?id=49214118)
- Source: [github.com](https://github.com/tmshapland/audiochatty-plugin)
- Score: 2
- Comments: 1
- Posted: 2026-08-07T18:03:04Z

## Translation

タイトル: クロード・コードと話そう。ゆっくり考えてみる
記事のタイトル: GitHub - tmshapland/audiochatty-plugin · GitHub
説明: GitHub でアカウントを作成して、tmshapland/audiochatty-plugin の開発に貢献します。

記事本文:
GitHub - tmshapland/audiochatty-plugin · GitHub
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
トムシャプランド
/
オーディオチャットプラグイン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット .claude-plugin .claude-plugin コマンド コマンド フック フック スクリプト スクリプト テスト テスト ラッパー ラッパー .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
このプラグインを使用すると、

クロード コード セッションにあなたの声で話しかけます。
クロードコードを擬似端末にラップすることで動作します。 audiochatty を実行すると、疑似端末が開き、その中で claude が起動され、レイヤーとして配置されます。
あなたとクロードコードを実行している端末との間で。実際に押すすべてのキーはそのままクロード コードに渡されます。
ただし、ラッパーは独自に入力することもできます。
疑似端末を使用すると、どこにいても Claude Code セッションと会話できます。オプトインしたセッションが終了したとき
ターン、それが行ったことの短い口頭要約があなたのメッセージに表示されます
オーディオおしゃべりな受信箱。あなたはそれを聞いて、フォローアップを求めます
声に出して質問し、次に何をするかを言うと、そのクロード コード セッションで実行されます。
クロード コードが何かを尋ねるために停止すると、ターミナルがフリーズし、代わりに audiochatty が尋ねます。応答はクロード コードで実行されます。
あなた、このラップトップを机から離れて
─────────────────────────────────
何をしたか聞いてください ◀── audiochatty ◀── scripts/stop_hook.py ◀── 終了したばかりのターン
次に何をするかを言う ──▶ audiochatty ──▶ audiochatty 実行 ──▶ セッションに入力
決定して大声で ◀─▶ audiochatty ◀─▶ scripts/permission_hook.py ◀─▶ ブロックされた許可プロンプト
3 つすべてがセッションごとにオプトインされ、切断すると 3 つすべてが停止します。 audiochatty で話した内容は、このマシン上でファイルを編集してコマンドを実行するコーディング エージェントに入力されます。同じものを入力するのとまったく同じくらい強力です
自分で端末に単語を入力することが重要であり、事前に知っておく価値があります。
セッションを接続します。
これが実際にマシンに与える影響
これでは目的が足りません

e.インストールする前に読んで、インストール後にもう一度読んでください。
以下は、このリポジトリ内のファイル名にあります。
最も簡単な方法は、audiochatty.com フロントエンドを使用することです。 「+」ボタンをクリックし、エージェントを接続するための指示に従います。
クロード プラグイン マーケットプレイス tmshapland/audiochatty-plugin を追加
クロード プラグインをインストール audiochatty@audiochatty
次に、マシンごとに 1 回、次のようにします。
> /audiochatty-ペア-スタート
https://audiochatty.com/link を開いて、次のコードを入力します。
WXYZ-1234
有効期限は 10 分です。
コードを入力したら、クロード コードで /audiochatty:audiochatty-pair-finish を実行します。
ブラウザで audiochatty に戻り、ペアリング ページのリンクをクリックしてコードを入力します。次に後半を実行します。
> /audiochatty-ペア-フィニッシュ
次に、audiochatty セッションを開始するためのショートカット コマンドを作成しましょう。やめる
Claude Code を作成し、この行をシェル プロファイル (vi ~/.zshrc または vi ~/.bashrc) に追加します。
エイリアス audiochatty="/path/to/audiochat-plugin/wrapper/audiochatty"
シェル プロファイルにショートカットを追加した後、
新しい端末なので有効になります。
Audiochatty を Claude Code セッションに接続するには、次のショートカットを使用して Claude Code を起動します。
audiochatty run --name [名前]
これによりセッションが接続されます。セッション内で実行するものは他にありません。
なぜ 2 つのコマンドなのか、そしてなぜトークンではなくコードなのか。スラッシュコマンドの出力は次のとおりです。
コマンドが終了した後にプロンプトに置き換えられるため、
code and then waited では、待機を放棄した後にのみコードが表示されます。
そして、それがそもそもコードである理由は、クロード コード プロンプトに入力されたものはすべて、
ディスク上のセッション .jsonl がモデルのコンテキストにロードされます。貼り付けられたトークンは、
そこに永遠に住んでください。そのため、ターミナルには短くて使い捨てのものが表示され、
有効期間の長いトークンはパスを通って戻ります。

トランスクリプトはバックエンドから決して参照されません
0600 ファイルに直接コピーします。
なぜ claude ではなく audiochatty run で開始するのか。 2番目の矢印は、
上の図は、すでに実行中のセッションに入力する必要があり、内部には何も入力する必要はありません。
セッションはそれを行うことができます。したがって、audiochatty を実行すると、疑似端末内でクロード コードが開始されます。
を所有し、キーストロークを直接渡し、話した内容を audiochatty に入力します。
Claude Code はそれを判断できません。ロードするプラグインも、起動フラグも、警告ダイアログもありません。
セッションの観点から見ると、誰かが何かを入力したことになります。実行後のすべてが渡されます
claude を変更しないため、audiochatty は --model claude-opus-4-8 を実行し、audiochatty は -- --resume を実行します。
期待どおりに動作します。 「wrapper/README.md」は、そのプロセスで何ができるかを含む詳細です。
そして出来ない。
$ audiochatty run # または: audiochatty run --name billing-refactor
コマンド 1 つで、現在のフォルダーの下の audiochatty にセッションが表示されます。
名前。 2 番目のステップはなく、セッションに何も入力する必要はありません。
意図的に、静かに接続します。端末はクロード コードのインターフェイスに属しており、
そこに出力された行は破損した画面です。確認は、に表示されるセッションです。
あなたの受信箱。逆に、接続が失敗する (オーディオチャットが到達不能、取り消される) ということです。
device) も表示されません: /audiochatty-status がその理由が表面化する場所であり、
セッションが表示されない場合に最初に実行されるもの。
普通に働きます。端末の見た目は変わりませんが、これは期待ではなく測定値です。
キーストロークあたり 34 マイクロ秒、再描画スループットはまったく関係なく、それぞれが完了しました
ターンはその名前で受信トレイに表示されます。目に見える変化の 1 つは、許可です。
プロンプト、AskUserQuestion、および計画の承認は、このセッションでのレンダリングを停止し、
代わりにターミナル、待ってください

声で決めるのはng。 「クロードの場合に正確に何が起こるか」を参照してください。
コードは何かを尋ねるために停止します。」を以下に示します。
あなたが開いている他のターミナルは何もしません: Stop と PermissionRequest
フックは両方ともグローバルであるため、どこでも実行され、そのセッションのマーカー ファイルを探し、
何も指定せずに終了します。接続されていない端末の許可プロンプトは、いつもとまったく同じようにレンダリングされます。
持っています。プレーン クロードで開始されたセッションには、そもそもリターン パスがありません。
開いている audiochatty 実行はバインドされておらず、アイドル状態のままです。
すべてをオフにするには: claude plugin disable audiochatty を実行し、セッションの開始に戻ります。
クロードと一緒に。これを削除するには、claude plugin uninstall audiochatty と rm -rf ~/.audiochatty を実行します。
所有しなくなったマシンを強制終了するには、そのトークンを取り消します
audiochatty の [設定] → [リンクされたデバイス] から。
完了したターンごとに 1 回、接続されたセッションで script/stop_hook.py がこれを POST し、
他には何もありません:
{
"claude_session_id" : " b3ea4f55-4ab6-48f7-8ba6-5fa8f3d2d81e " ,
"last_assistant_message" : " リファクタリングが完了し、テストに合格しました。 " ,
"tool_calls" : [ " 編集 " 、 " 編集 " 、 " Bash " 、 " 読み取り " ]、
"stop_reason" : " end_turn " ,
"cwd" : " /Users/mike/repos/audiochat "
}
last_assistant_message — ターンの最後のアシスタント テキスト。クロード コード
フックを直接渡します。話された要約は主にこれに基づいて書かれています。
tool_calls — このターンで使用されたツールの名前。セッションから読み取られます。
転写。名前のみ: 引数、ファイルの内容、コマンドライン、出力はありません。
stop_reason 、 cwd — ターンがどのように終了したか、およびそれがどのリポジトリであったか。
ペイロードはこれらのキーから最初から構築されるため、フックの入力からの他の情報は送信されません。プロンプトはそうではありません
送信されました。コードは送信されません。バックエンドは、到着時にすべてのフィールドに再び上限を設定します。
どこへ行くのか

その後です。バックエンドがそれをキューに入れ、ワーカーがそれを書き換えます。
サードパーティ モデル (Gemini) を使用して大声で聞く価値のあるもの。上記の生のペイロード
は概要と一緒に保存されるため、そこからフォローアップの質問に答えることができます。もし
それはあなたの仕事には受け入れられないので、それらのセッションを接続しないでください。
SessionEnd は、セッションが本当に終了した場合にのみ、 {"claude_session_id": …} のみを送信します。
終了 — /clear または /resume では決して行われず、同じセッション ID と同じオープン状態が維持されます。
ターミナル。
ラッパーがセッションに入力したものとまったく同じ
別の方向、そして同じ約束。届くのはイベントでも通知でも、
独自のラッパーを持つもの。プロンプトでのあなたの言葉です。
> それを元に戻し、代わりに他のヘルパーを使用します
それがすべてです。届くのは、あなたが言ったことを整然と表現したものであり、逐語的なものではありません
トランスクリプト: audiochatty の音声エージェントがトランスクリプトを作成し、誤った開始と結合をクリーンアップします。
コールドで読むことを意図したものから期待される方法で断片化されています。追加するように指示されています
何もせず、あなたに代わって何も解決せず、何も残さないでください。それはあなたの場所に正確に入ります
独自のタイピングが行われるため、クロード コードはそれを独自のタイピングとして扱います。
型付け方法に関する 3 つの点。これらはすべてバグのように見えてそうではありません。
ペースト状で届くので、Enter を 1 回押します。口頭での指示は通常、複数のユーザーに実行されます。
段落。 raw と入力すると、プロンプトは最初の改行で送信され、残りは改行として扱われます。
2 番目の命令なので、代わりに端末の括弧で囲まれた貼り付けマーカーにラップされて入ります。
エスケープ シーケンスは最初にテキストから削除されるため、メッセージ内に何かが含まれていてもメッセージを終了することはできません。
早めに貼り付けて、独自のキーストロークの発行を開始します。
あなたが入力をやめるのを待ちます。文章の途中で何かが届いた場合、
w

ラッパーは一時停止するまで (デフォルトでは 1.5 秒) 保持します。遅いと思われる命令は、
通常、これは機能しており、何かが詰まっているわけではありません。
表示されるまでに最大 30 秒かかる場合があります。ラッパーはポーリングを行います。 5秒ごとにチェックします
セッションがアクティブな間、しばらく何も到着しないと速度が低下します。何もない
このマシンにプッシュされました。
ツールは追加されておらず、意図的に応答ツールもありません。Stop フックはすでに
返信パスなので、何が起こったのかを聞くことで、指示の順番が要約されます。
引き起こされた。 「wrapper/README.md」には、ポーリング、重複排除、および制限などの残りの部分が文書化されています。
そのプロセスができるのです。
クロード・コードが立ち止まって何かを尋ねると、まさに何が起こりますか
これは 3 番目の方向であり、他の 2 つの方向とは動作が異なります。
レポートまたは入力すると、 script/permission_hook.py がダイアログの代わりに応答するため、
ダイアログはまったくレンダリングされません。クロード コードが停止して待機するすべての内容をカバーします。
キー押下については、Bash または Edit 承認、AskUserQuestion 、ExitPlanMode プランのいずれかになります。これは、3 つすべてが内部で同じ PermissionRequest フックを起動するためです。
接続されたセッションでは、端末がフリーズし、audiochatty から次の質問が送信されます。
ツールの承認では、ツール名とそのツールが実行する内容の短い概要が読み上げられます (
Bash のコマンド、ファイル編集のパス) には、許可するか許可しないかの 2 つのオプションがあります。
多肢選択の質問 (AskUserQuestion

[切り捨てられた]

## Original Extract

Contribute to tmshapland/audiochatty-plugin development by creating an account on GitHub.

GitHub - tmshapland/audiochatty-plugin · GitHub
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
tmshapland
/
audiochatty-plugin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits .claude-plugin .claude-plugin commands commands hooks hooks scripts scripts tests tests wrapper wrapper .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
This plugin lets you talk to a Claude Code session with your voice.
It works by wrapping Claude Code in a pseudo-terminal. audiochatty run opens a pseudo-terminal, starts claude inside it, and sits as a layer
between you and the terminal running Claude Code. Every key you actually press still passes straight through to Claude Code,
but the wrapper can also type into it on its own.
The pseudo-terminal allows you talk to a Claude Code session from wherever you are. When a session you've opted in finishes
a turn, a short spoken summary of what it did shows up in your
audiochatty inbox. You listen to it, ask follow-up
questions out loud,m and say what to do next, which then gets executed in that Claude Code session.
When Claude Code stops to ask you something, the terminal freezes and audiochatty asks you instead. Your response is executed in Claude Code.
you, away from your desk this laptop
───────────────────────── ────────────────────────────────────────
hear what it did ◀── audiochatty ◀── scripts/stop_hook.py ◀─ the turn it just finished
say what to do next ──▶ audiochatty ──▶ audiochatty run ──▶ typed into your session
decide, out loud ◀─▶ audiochatty ◀─▶ scripts/permission_hook.py ◀─▶ a blocked permission prompt
All three are opt-in per session, and all three stop when you disconnect. What you say into audiochatty is typed into a coding agent that edits files and runs commands on this machine. It is exactly as powerful as typing the same
words into the terminal yourself, which is the point, and it is worth knowing before you
connect a session.
What this actually does to your machine
This is short on purpose. Read it before you install, and read it again after — everything
below is in this repo, in the files it names.
The easiest path is to use the audiochatty.com frontend. Click the '+' button and follow the directions for connecting an agent.
claude plugin marketplace add tmshapland/audiochatty-plugin
claude plugin install audiochatty@audiochatty
Then, once per machine:
> /audiochatty-pair-start
Open https://audiochatty.com/link and enter this code:
WXYZ-1234
It expires in 10 minutes.
After you enter the code, run /audiochatty:audiochatty-pair-finish here in Claude Code.
Go back to audiochatty in a browser, click the link for the pairing page, and type the code. Then run the second half:
> /audiochatty-pair-finish
Next, let's create a shortcut command for starting an audiochatty session. Quit
Claude Code and add this line to your shell profile (vi ~/.zshrc or vi ~/.bashrc).
alias audiochatty="/path/to/audiochat-plugin/wrapper/audiochatty"
After you add the shortcut to your shell profile, open a
new terminal so it takes effect.
To connect Audiochatty to a Claude Code session, start Claude Code with the shortcut:
audiochatty run --name [name]
That connects the session for you — there's nothing else to run inside it.
Why two commands, and why a code instead of a token. A slash command's output is
substituted into the prompt after the command exits, so a single command that minted a
code and then waited would only show you the code once it had already given up waiting.
And the reason it's a code at all: anything typed into a Claude Code prompt is written to
the session .jsonl on disk and loaded into the model's context. A pasted token would
live there forever. So the terminal displays something short and disposable, and the
long-lived token travels back over a path the transcript never sees from the backend
straight into a 0600 file.
Why you start it with audiochatty run and not claude . The second arrow in the
diagram above has to type into a session that is already running, and nothing inside a
session can do that. So audiochatty run starts Claude Code inside a pseudo-terminal it
owns, passes your keystrokes straight through, and types in what you speak into audiochatty.
Claude Code can't tell: there's no plugin to load, no launch flag, and no warning dialog.
From the session's point of view, somebody typed something. Everything after run is passed
to claude unchanged, so audiochatty run --model claude-opus-4-8 and audiochatty run -- --resume
work the way you'd expect. wrapper/README.md is the detail, including what that process can
and can't do.
$ audiochatty run # or: audiochatty run --name billing-refactor
One command, and the session shows up in audiochatty under the current folder's
name. There is no second step and nothing to type into the session.
It connects silently, on purpose. The terminal belongs to Claude Code's interface, and a
line printed into it is a corrupted screen. The confirmation is the session appearing in
your inbox. The flip side is that a connect which fails (audiochatty unreachable, a revoked
device) is also invisible: /audiochatty-status is where that reason surfaces, and it's the
first thing to run if a session never shows up.
Work normally. The terminal looks no different, which is measured rather than hoped: about
34 microseconds per keystroke and nothing at all on redraw throughput, and each completed
turn shows up in your inbox under that name. The one visible change is that permission
prompts, AskUserQuestion s, and plan approvals stop rendering in this session and freeze the
terminal instead, waiting for you to decide by voice. See "Exactly what happens when Claude
Code stops to ask you something" below.
Every other terminal you have open does nothing at all: the Stop and PermissionRequest
hooks are both global, so they run everywhere, look for a marker file for that session, find
none, and exit. A permission prompt in an unconnected terminal renders exactly as it always
has. A session started with plain claude has no return path to begin with, and any other
audiochatty run you have open sits unbound and idle.
To turn everything off: claude plugin disable audiochatty , and go back to starting sessions
with claude . To remove it: claude plugin uninstall audiochatty and rm -rf ~/.audiochatty .
To kill a machine you no longer have, revoke its token
from Settings → Linked devices in audiochatty.
Once per completed turn, in a connected session, scripts/stop_hook.py POSTs this and
nothing else:
{
"claude_session_id" : " b3ea4f55-4ab6-48f7-8ba6-5fa8f3d2d81e " ,
"last_assistant_message" : " I finished the refactor and the tests pass. " ,
"tool_calls" : [ " Edit " , " Edit " , " Bash " , " Read " ],
"stop_reason" : " end_turn " ,
"cwd" : " /Users/mike/repos/audiochat "
}
last_assistant_message — the final assistant text of the turn, which Claude Code
hands the hook directly. It is what the spoken summary is mostly written from.
tool_calls — the names of the tools this turn used, read out of the session
transcript. Names only: no arguments, no file contents, no command lines, no output.
stop_reason , cwd — how the turn ended, and which repo it was.
The payload is built from scratch out of those keys, so nothing else from the hook's input is ever transmitted. Your prompts are not
sent. Your code is not sent. The backend caps every field again on arrival.
Where it goes afterwards. The backend queues it, and a worker rewrites it into
something worth hearing out loud using a third-party model (Gemini). The raw payload above
is stored alongside the summary so that follow-up questions can be answered from it. If
that isn't acceptable for your work, don't connect those sessions.
SessionEnd sends only {"claude_session_id": …} , and only when a session has genuinely
ended — never on /clear or /resume , which keep the same session id and the same open
terminal.
Exactly what the wrapper types into your session
The other direction, and the same promise. What arrives is not an event, a notification, or
anything with a wrapper of its own. It is your words, at the prompt:
> change that back, and use the other helper instead
That's the whole of it. What arrives is a tidied rendering of what you said, not a verbatim
transcript: audiochatty's voice agent transcribes it, then cleans up false starts and joins
fragments the way you'd expect from something meant to be read cold. It is instructed to add
nothing, resolve nothing on your behalf, and leave nothing out. It goes in exactly where your
own typing goes, so Claude Code treats it as your own typing because that is what it is.
Three things about how it is typed, all of which look like bugs and aren't:
It arrives as a paste, then one Enter. A spoken instruction routinely runs to several
paragraphs. Typed raw, the prompt would submit at the first newline and treat the rest as a
second instruction, so it goes in wrapped in the terminal's bracketed-paste markers instead.
Escape sequences are stripped from the text first, so nothing inside a message can end the
paste early and start issuing keystrokes of its own.
It waits for you to stop typing. If you're mid-sentence when something arrives, the
wrapper holds it until you've paused (1.5s by default). An instruction that seems slow is
usually this working, not something stuck.
It may take up to 30 seconds to show up. The wrapper polls; it checks every 5 seconds
while a session is active and slows down when nothing has arrived for a while. Nothing is
pushed to this machine.
No tool is added and there is deliberately no reply tool — the Stop hook is already the
reply path, so the way you hear what happened is the summary of the turn your instruction
caused. wrapper/README.md documents the rest: the poll, the dedupe, and the limits on what
that process can do.
Exactly what happens when Claude Code stops to ask you something
This is the third direction, and it works differently from the other two: instead of
reporting or typing, scripts/permission_hook.py answers in place of the dialog , so the
dialog never renders at all. It covers everything Claude Code would otherwise stop and wait
for a keypress on — a Bash or Edit approval, an AskUserQuestion , an ExitPlanMode plan because all three fire the same PermissionRequest hook under the hood.
In a connected session, the terminal freezes and audiochatty sends you the question:
A tool approval reads out the tool name and a short summary of what it would do (the
command for Bash , the path for a file edit) and offers two options: allow it, or don't.
A multiple-choice question ( AskUserQuestion

[truncated]
