---
source: "https://blog.angeloff.name/post/2026/05/29/teaching-tmux-to-babysit-my-claude-code-agents/"
hn_url: "https://news.ycombinator.com/item?id=48327021"
title: "Teaching tmux to babysit my Claude Code agents"
article_title: "Teaching tmux to babysit my Claude Code agents · tmpfs /home/stan"
author: "StanAngeloff"
captured_at: "2026-05-30T11:41:31Z"
capture_tool: "hn-digest"
hn_id: 48327021
score: 3
comments: 0
posted_at: "2026-05-29T18:09:53Z"
tags:
  - hacker-news
  - translated
---

# Teaching tmux to babysit my Claude Code agents

- HN: [48327021](https://news.ycombinator.com/item?id=48327021)
- Source: [blog.angeloff.name](https://blog.angeloff.name/post/2026/05/29/teaching-tmux-to-babysit-my-claude-code-agents/)
- Score: 3
- Comments: 0
- Posted: 2026-05-29T18:09:53Z

## Translation

タイトル: クロード コード エージェントのベビーシッターを tmux に教える
記事のタイトル: クロード コード エージェントのベビーシッターを tmux に教える · tmpfs /home/stan
説明: あなたが私と同じなら、あなたはもはや 1 つのクロード コード セッションを実行するのではなく、小さなフリートを実行していることになります。 1 つのウィンドウは TypeScript モノリポジトリで機能を構築しており、別のウィンドウは...

記事本文:
~ tmpfs /home/stan ブログ · クロード コード エージェントのベビーシッターとして tmux を教えることについて
あなたが私と同じなら、もう 1 つのクロード コード セッションを実行するのではなく、小さなフリートを実行していることになります。 1 つのウィンドウは TypeScript モノリポジトリで機能を構築し、別のウィンドウは同僚のプル リクエストをレビューし、3 番目のウィンドウは Neovim プラグインの再描画バグを追跡しています。 tmux を使用するとこれが簡単になります。エージェントごとにウィンドウがあり、Alt + 1 ～ 9 でウィンドウ間を移動できます。多くの場合、それらは異なるプロジェクトですらない - 同じリポジトリの git ワークツリー、ブランチごとに 1 つのエージェントがあるだけです。
入力するのではなく、口述で入力するのではなく、読みます。私はこの投稿を大声で考えて文字に起こし、記述されているコードと構成に直接アクセスできる Claude Code を使用して書き上げました。私はすべての言葉を読みます。アイデアと方向性は私のもので、散文はコラボレーションです。
問題は 3 ～ 4 人のエージェントから始まります。あなたは自分自身に 20 の質問をしながら窓を巡回して時間を費やします。「これはまだ機能していますか?」あの人は私に何かを尋ねるために立ち止まったのですか？ 10 分前の処理は終了しましたか、それとも git プッシュの承認を待っているのでしょうか?エージェントはあなたを必要とする瞬間まで自律的であり、最初からあなたの肩をたたく方法はありません。
Hacker News には修正情報が満載だ。エージェント ランナー ツールの新しい波は、その多くが垂直タブの列を持ち、エージェントごとに 1 つずつ、注目を集めたいときに点灯する。カーソルの v3 書き換え、「エージェント時代の IDE」、「チーム向けのサンドボックス コーディング エージェント」、そして現在第 2 バージョンとなっている Google の Antigravity である。これらは、PR ワークフロー、GitHub 統合、サンドボックス、チーム オーケストレーションなど、タブを明るくするだけではなく、おそらくそれらからアイデアを借用するでしょう。しかし、私は一日中 tmux に住んでおり、私が望んでいたのはほんの一部、つまりどのエージェントが私を必要としているかを一目で把握できることだけでした。それは新しいものは必要ありません

IDE、すでに開いているウィンドウだけです。
そこで tmux にそれを表示するように教えました。すべてのウィンドウ内にクロード コード セッションの色付きのドットが表示されます。
オレンジ色 - ブロックされており、私が必要です（許可のプロンプトまたは質問）。
緑 — 完了し、応答が読み取られるのを待っています。
何もない — 離れて働いているか、報告することが何もない。
緑の点は、そのウィンドウに切り替えるとすぐに消えます。私が本を読んでいる頃には、消えています。琥珀色の点はより頑固です。ウィンドウをちらっと見ることは、それに応答することと同じではないため、実際にリクエストを処理するまでこのままになります。
これは、直接接触することのない 2 つの部分であり、単一の tmux 変数上で接触します。クロード コードのフックがそれを記述します。 tmux フォーマット文字列がそれを読み取り、ドットを描画します。 nix-meridian 設定の home-manager を介して実行しているので、スニペットは Nix ですが、実質はシェルと tmux 設定です。これらをそのまま持ち上げます。
半分: クロード コードがウィンドウにフラグを立てる
クロード コードのフックは、セッション内の設定ポイントでシェル コマンドを実行します。私が使っているもの:
PermissionRequest — 許可されていないことを実行したいと考えており、はい/いいえの応答を待ちます。
引き出し — 構造化された質問をします。
Stop / StopFailure — ターンが終了しました。
PostToolUse — 完成したばかりのツール。
UserPromptSubmit — 新しいプロンプトを送信しました。
接着剤は 1 つの変数 $TMUX_PANE です。 tmux はそれをすべてのペインに設定し、クロード コードはそれを継承するため、フックは常にどのウィンドウで実行されているかを認識します。そのウィンドウにフラグを立てるのは 1 つのコマンドです。
# このペインのウィンドウに注意が必要であるというフラグを立てます。
$ tmux set -w -t " $TMUX_PANE " @claude-state 権限
# ...そしてそれをクリアします:
$ tmux set -wu -t " $TMUX_PANE " @claude-state
@claude-state はカスタム オプションです。tmux を使用すると、 @ で始まる任意の名前を作成できます。 -w を指定するとウィンドウにスコープが設定され、-u を指定すると設定が解除されます。ウィンドウごとに 1 つの文字列、パーミッションを保持

、誘発、アイドル、または何もありません。
set、clear、および条件付きクリアを小さな Nix ヘルパーにラップします。
tmux-クロード-状態 =
させてください
tmux = " ${ lib . getBin pkgs . tmux } /bin/tmux" ;
で
{
# このペインのウィンドウで @claude-state を特定の値に設定します。
セット = 状態:
''[ -n "$TMUX_PANE" ] && ${ tmux } set -w -t "$TMUX_PANE" @claude-state ${ state } ||真実'' ;
＃それが何であれ、それをクリアしてください。
リセット =
''[ -n "$TMUX_PANE" ] && ${ tmux } set -wu -t "$TMUX_PANE" @claude-state ||真実'' ;
# 現在ブロックされている場合に *のみ* クリアしてください。他のものは壊さないでください。
リセットブロック = ''
[ -n "$TMUX_PANE" ] && case "$( ${ tmux } show -wv -t "$TMUX_PANE" @claude-state 2>/dev/null)" in
許可|引き出し) ${ tmux } set -wu -t "$TMUX_PANE" @claude-state ;;
イーサック;真実'' ;
} ;
[ -n "$TMUX_PANE" ] ガードにより、tmux の外部では何も操作されなくなります。末尾の || true は、失敗した tmux 呼び出しがフック エラーとして表面化するのを停止します。次に、それらを接続します (トリミングされており、完全なセットはリポジトリにあります)。
フック = {
PermissionRequest = [{
フック = [
{タイプ = "コマンド" ; command = "...チャイムを鳴らします..." ; # これについては後で詳しく説明します
{タイプ = "コマンド" ;コマンド = tmux-claude-state 。 「許可」を設定します。 }
];
}];
停止 = [{ フック = [{ タイプ = "コマンド" ;コマンド = tmux-claude-state 。 「アイドル」を設定します。 }]; }];
PostToolUse = [{ フック = [{ タイプ = "コマンド" ;コマンド = tmux-claude-state 。リセットブロック済み ; }]; }];
UserPromptSubmit = [{ フック = [{ タイプ = "コマンド" ;コマンド = tmux-claude-state 。リセット; }]; }];
} ;
誘導は PermissionRequest (チャイム 1 なし) をミラーし、 StopFailure は Stop をミラーし、SessionEnd は UserPromptSubmit をミラーします。ロジック:
ブロックされました ( PermissionRequest / elicit ) → 許可 / 誘導 — オレンジ色。
完了 ( Stop / StopFailure ) → アイドル — 緑色。
ツールの実行 ( PostToolUse ) → オレンジ色の場合はクリアします。あなたは今承認しました

編集。
新しいプロンプト/セッションの終了 → すべてクリアします。
「動作中」のフラグは何もありません。これがデフォルトであり、9 つのビジー ドットの行は単なるノイズになります。
tmux は変数を window-status-format 、つまり非アクティブなウィンドウごとのテンプレートでレンダリングします。
setw -g ウィンドウステータス形式 \
" #I│ #{?#{@claude-state},#{?#{==:#{@claude-state},idle},#[fg=colour34]●#[fg=colour247] ,#[fg=colour214]●#[fg=colour247] },}#W "
tmux のフォーマット構文は緻密です。#{?cond,then,else} は 3 項であり、ネストされています。展開時:
#{?#{@claude-state}, # そもそも @claude-state は設定されていますか?
#{?#{==:#{@claude-state},idle}, # はい — まさに「アイドル」ですか?
#[fg=colour34]●#[fg=colour247] , # はい → 緑色の点、その後灰色に戻ります
#[fg=colour214]●#[fg=colour247] }, # いいえ→オレンジ色の点、その後グレーに戻ります
} # 設定されていない → 何もレンダリングしない
@claude-state がない場合は空の文字列が返されるため、名前は通常どおり表示されます。それ以外の場合は、アイドル状態での完全一致によって色が選択されます。 ● は Unicode の円です。 #[fg=...] は、その両側の色を設定します。インデックスはプレーンな 256 色パレットです。color34 グリーン (「完了、来て見てください」)、colour214 アンバー (「I need you」)、colour247 はドットの後に復元された非アクティブなテキストのグレーです。ステータスバーに隠れている信号機。
フォーカス上の緑色の点を消去する
緑色の点は、見るまでは役に立たないので、ウィンドウを選択したときに消去します。
# ウィンドウに切り替えるとき、そのクロード状態が緑色の「アイドル」の場合
# フラグをクリアします。とにかく応答を読み取ろうとしています。
set-hook -g after-select-window \
'if -F "#{==:#{@claude-state},idle}" "set -wu @claude-state"'
アイドル状態のみをクリアし、許可/誘発は決してクリアしません。緑は「読むべきもの」を意味し、そこに切り替えるとそれを読みます。琥珀は「何かをする」という意味ですが、一見しただけではわかりません。琥珀色の点は、世界が終わるまで存続します。

rk が発生する ( PostToolUse → replace-blocked ) か、新しいプロンプトを送信します。
最後にもう 1 つの便利な点があります。フォーカスされたタブは、ドット ロジックをまったく使用しない別個の window-status-current-format を使用します。ウィンドウを見ている場合、その中に何があるかを伝える必要はありません。したがって、ドットは、自分が使用していないウィンドウにのみ表示されます。
プロンプトを送信します。 UserPromptSubmit はフラグをクリアします (ドットなし)。
クロードは働いています。状態には何も触れていません - まだドットはありません。
git Push がヒットし、無人で実行することはできません。 PermissionRequest の起動: チャイム、オレンジ色のタブ。
私はそれを承認します - ちらっと見ても消えなかったので、オレンジ色の点は残りました。コマンドが実行され、PostToolUse によってフラグがクリアされます。
ターンは終了します。停止するとタブが緑色に変わります。
切り替えて読んでみます。 after-select-window は、最初の行を終了する前に緑色の点をクリアします。
端末に目を向けているとドットが機能します。あなたがさまよったときはそうではありません。したがって、ブロッキングのケースでも、それだけでノイズが発生します。
PermissionRequest = [{
フック = [
{
タイプ = "コマンド" ;
コマンド = " ${ lib . getBin pkgs . Pipewire } /bin/pw-play ${ ./audio/notifications/mixkit-clear-announce-tones-2861.mp3 } " ;
タイムアウト = 5 ;
}
{タイプ = "コマンド" ;コマンド = tmux-claude-state 。 「許可」を設定します。 }
];
}] ;
pw-play は PipeWire のプレーヤーです。クリップはリポジトリにベンダーされた短い Mixkit トーンであるため、サウンドを生成するためだけにネットワークに到達するものはありません。好きなものを使ってください。
PermissionRequest でのみ鳴動し、 Stop では鳴動しません。曲がり角が終わるたびにチャイムが毎時9時に鳴らされ、昼食までにはチャイムが鳴り止むことになる。エージェントが本当に行き詰まった場合にのみチャイムを鳴らし続ける価値があります。これにより、私は立ち去ることができます。3 人のエージェントを起動し、コーヒーを淹れに行き、そのうちの 1 人から電話がかかってきたら戻ってきます。残りは続けます。
フォークされたサブエージェントは、理由もなくウィンドウをアンバー色に変えることがあります。サブエージェントのフォークがオンになっていると、サブエージェントが時々スポーンします

何も待っていないときはウィンドウをオレンジ色にします。フォーク内の何かがブロックフックを引っ掛けます。まだ何なのかは特定できていない。
黄色の点はビートを遅くクリアします。ユーザーが承認したときではなく、PostToolUse でツールが終了したときにクリアします。長時間のビルドやスリープ 30 にゴーサインを出すと、ウィンドウは実行中ずっとオレンジ色のままですが、「はい」と言った瞬間にウィンドウは必要なくなりました。 PreToolUse フックを使用すると、より早くクリアされます。その人はリストに載っています。
これがすべてです。すべてのエージェントが表示され、必要なときに通知が届くという体験は、大きなツールを中心に構築されています。ただし、その一部は私がすでに作業している端末内に存在します。 数十行のフック、1 つのフォーマット文字列、およびサウンド ファイルです。
これを望んでいるのは私だけではありません。Solo はステータス インジケーターが組み込まれたすべてのエージェントを実行する素敵なネイティブ ワークスペース (Electron ではなく Tauri) であり、一見の価値があります。それは私の一握りの点よりもはるかに多くのことを行います。それは私にとって必要以上のものです。ワークスペース全体を手渡したい場合は、それが最適です。
賢い部分はいずれも Claude 固有のものではありません。変数を書き込むフックとそれを読み取るステータス バーです。ロングメイクまたはデプロイ時に tmux set -w @claude-state をポイントすると、同じドットが無料で得られます。
完全バージョンは、 home/apps/tmux および home/apps/claude-code の下の nix-meridian 設定にあります。それを盗んで、二度と窓探しをしないでください。
両方のブロック状態 (許可リクエストと引き出し) は同じオレンジ色の点を表示しますが、ベルを鳴らすのは PermissionRequest のみです。それはコーヒーの途中で私を捕まえる傾向があります。 ↩

## Original Extract

If you are like me, you no longer run one Claude Code session — you run a small fleet. One window is building a feature in a TypeScript monorepo, anothe...

~ tmpfs /home/stan Blog · About Teaching tmux to babysit my Claude Code agents
If you are like me, you no longer run one Claude Code session — you run a small fleet. One window is building a feature in a TypeScript monorepo, another is reviewing a colleague’s pull request, a third is chasing a redraw bug in a Neovim plugin. tmux makes this easy: a window per agent, Alt + 1 – 9 to jump between them. Often they are not even different projects — just git worktrees of the same repo, one agent per branch.
Dictated, not typed — but read. I thought this post out loud and transcribed it, then wrote it up with Claude Code, which had direct access to the code and config it describes. I read every word; the ideas and direction are mine, the prose a collaboration.
The trouble starts at three or four agents. You burn time cycling through windows playing twenty questions with yourself: is this one still working? Has that one stopped to ask me something? Did the one from ten minutes ago finish, or is it waiting on me to approve a git push ? An agent is autonomous right up until the moment it needs you — and out of the box it has no way to tap you on the shoulder.
Hacker News is full of the fix: a new wave of agent-runner tools, many with a column of vertical tabs, one per agent, each lighting up when it wants attention — Cursor’s v3 rewrite , an “IDE for the agents era” , “sandboxed coding agents for a team” and Google’s Antigravity , now on its second version. These do far more than light up a tab — PR workflows, GitHub integration, sandboxing, team orchestration — and I will probably borrow ideas from them. But I live in tmux all day, and all I wanted was the small part: a glanceable sense of which agent needs me. That does not need a new IDE, just the windows I already have open.
So I taught tmux to show it. Every window now carries a coloured dot for the Claude Code session inside it:
amber — blocked, needs me (a permission prompt or a question);
green — finished, with a response waiting to be read;
nothing — working away, or nothing to report.
The green dot clears itself the instant I switch to that window — by the time I am reading, it is gone. The amber dot is stubborner: it stays until I have actually dealt with the request, because glancing at a window is not the same as answering it.
It is two halves that never touch directly — they meet on a single tmux variable. Claude Code hooks write it; a tmux format string reads it back and paints the dot. I drive it through home-manager in my nix-meridian config, so the snippets are Nix, but the substance is the shell and tmux config — lift those straight out.
Half one: Claude Code flags the window
Claude Code hooks run a shell command at set points in a session. The ones I use:
PermissionRequest — wants to do something unauthorised, waiting on a yes/no.
Elicitation — asking you a structured question.
Stop / StopFailure — the turn ended.
PostToolUse — a tool just finished.
UserPromptSubmit — you sent a new prompt.
The glue is one variable: $TMUX_PANE . tmux sets it in every pane and Claude Code inherits it, so a hook always knows which window it is running in. Flagging that window is one command:
# Flag this pane's window as needing attention:
$ tmux set -w -t " $TMUX_PANE " @claude-state permission
# ...and clear it:
$ tmux set -wu -t " $TMUX_PANE " @claude-state
@claude-state is a custom option — tmux lets you invent any name starting with @ . -w scopes it to the window and -u unsets it. One string per window, holding permission , elicitation , idle or nothing.
I wrap set, clear and a conditional clear into a small Nix helper:
tmux-claude-state =
let
tmux = " ${ lib . getBin pkgs . tmux } /bin/tmux" ;
in
{
# Set @claude-state to a given value on this pane's window.
set = state :
''[ -n "$TMUX_PANE" ] && ${ tmux } set -w -t "$TMUX_PANE" @claude-state ${ state } || true'' ;
# Clear it, whatever it was.
reset =
''[ -n "$TMUX_PANE" ] && ${ tmux } set -wu -t "$TMUX_PANE" @claude-state || true'' ;
# Clear it *only* if we are currently blocked — don't clobber anything else.
reset-blocked = ''
[ -n "$TMUX_PANE" ] && case "$( ${ tmux } show -wv -t "$TMUX_PANE" @claude-state 2>/dev/null)" in
permission|elicitation) ${ tmux } set -wu -t "$TMUX_PANE" @claude-state ;;
esac; true'' ;
} ;
The [ -n "$TMUX_PANE" ] guard makes it a no-op outside tmux; the trailing || true stops a failed tmux call surfacing as a hook error. Then wire them up (trimmed — the full set is in the repo):
hooks = {
PermissionRequest = [{
hooks = [
{ type = "command" ; command = "...play a chime..." ; } # more on this later
{ type = "command" ; command = tmux-claude-state . set "permission" ; }
];
}];
Stop = [{ hooks = [{ type = "command" ; command = tmux-claude-state . set "idle" ; }]; }];
PostToolUse = [{ hooks = [{ type = "command" ; command = tmux-claude-state . reset-blocked ; }]; }];
UserPromptSubmit = [{ hooks = [{ type = "command" ; command = tmux-claude-state . reset ; }]; }];
} ;
Elicitation mirrors PermissionRequest (without the chime 1 ), StopFailure mirrors Stop , SessionEnd mirrors UserPromptSubmit . The logic:
blocked ( PermissionRequest / Elicitation ) → permission / elicitation — amber.
finished ( Stop / StopFailure ) → idle — green.
tool ran ( PostToolUse ) → clear it if we were amber; you have just approved it.
new prompt / session end → clear everything.
Nothing flags “working” — that is the default, and a row of nine busy dots would just be noise.
tmux renders the variable in window-status-format , the template for every inactive window:
setw -g window-status-format \
" #I│ #{?#{@claude-state},#{?#{==:#{@claude-state},idle},#[fg=colour34]●#[fg=colour247] ,#[fg=colour214]●#[fg=colour247] },}#W "
tmux’s format syntax is dense: #{?cond,then,else} is a ternary, and they nest. Unfolded:
#{?#{@claude-state}, # is @claude-state set at all?
#{?#{==:#{@claude-state},idle}, # yes — is it exactly "idle"?
#[fg=colour34]●#[fg=colour247] , # yes → green dot, then back to grey
#[fg=colour214]●#[fg=colour247] }, # no → amber dot, then back to grey
} # not set → render nothing
No @claude-state and you get an empty string, so the name renders as normal. Otherwise an exact match on idle picks the colour. ● is a Unicode circle; #[fg=...] sets the colour on either side of it. The indices are the plain 256-colour palette: colour34 green (“done, come and look”), colour214 amber (“I need you”) and colour247 the grey for inactive text, restored after the dot. A traffic light, hiding in a status bar.
Clearing the green dot on focus
A green dot is only useful until you look, so I clear it when you select the window:
# When you switch to a window, if its Claude state is the green "idle"
# flag, clear it — you are about to read the response anyway.
set-hook -g after-select-window \
'if -F "#{==:#{@claude-state},idle}" "set -wu @claude-state"'
It only clears idle , never permission / elicitation . Green means “something to read”, and switching there reads it; amber means “something to do ”, which a glance does not. The amber dot survives until the work happens ( PostToolUse → reset-blocked ) or you send a new prompt.
One last nicety: the focused tab uses a separate window-status-current-format with no dot logic at all — if you are looking at a window, you do not need telling what is in it. So dots only ever appear on the windows you are not in.
I send a prompt. UserPromptSubmit clears the flag — no dot.
Claude works. Nothing touches the state — still no dot.
It hits a git push it cannot run unattended. PermissionRequest fires: chime, amber tab.
I approve it — the amber dot stayed, since glancing didn’t clear it. The command runs and PostToolUse clears the flag.
The turn ends. Stop turns the tab green.
I switch over to read it. after-select-window clears the green dot before I finish the first line.
Dots work when your eyes are on the terminal. They do not when you have wandered off. So the blocking case — and only that — also makes a noise:
PermissionRequest = [{
hooks = [
{
type = "command" ;
command = " ${ lib . getBin pkgs . pipewire } /bin/pw-play ${ ./audio/notifications/mixkit-clear-announce-tones-2861.mp3 } " ;
timeout = 5 ;
}
{ type = "command" ; command = tmux-claude-state . set "permission" ; }
];
}] ;
pw-play is PipeWire’s player; the clip is a short Mixkit tone vendored into the repo, so nothing reaches the network just to make a sound. Use whatever you like.
It rings on PermissionRequest only, never on Stop . A chime on every finished turn would be nine an hour and muted by lunch; a chime only when an agent is genuinely stuck is worth keeping on. This is what lets me walk away — start three agents, go and make a coffee, come back when one of them calls. The rest carry on.
Forked subagents sometimes flip a window amber for nothing. With subagent forking on, spawning one occasionally turns the window amber when nothing is waiting. Something inside the fork trips a blocking hook; I have not pinned down what yet.
The amber dot clears a beat late. I clear it on PostToolUse — when the tool finishes , not when you approve it. Green-light a long build or a sleep 30 and the window stays amber for the whole run, though it stopped needing you the second you said yes. A PreToolUse hook would clear it sooner; that one is on the list.
That is the whole thing: the see-every-agent, ping-me-when-one-needs-you experience the big tools are built around, except this slice of it lives in the terminal I already work in. A dozen lines of hooks, one format string and a sound file.
I am not the only one who wants this — Solo is a lovely native workspace (Tauri, not Electron) that runs all your agents with status indicators built in, and it is well worth a look. It does far more than my handful of dots; it is just more than I need. If you would rather have the whole workspace handed to you, that is what it is for.
None of the clever part is Claude-specific — it is hooks writing a variable and a status bar reading it. Point tmux set -w @claude-state at a long make or a deploy and you get the same dots for free.
The full version lives in my nix-meridian config under home/apps/tmux and home/apps/claude-code . Steal it, and never go window-hunting again.
Both blocking states — a permission request and an elicitation — show the same amber dot, but only PermissionRequest rings the bell; that is the one that tends to catch me mid-coffee. ↩
