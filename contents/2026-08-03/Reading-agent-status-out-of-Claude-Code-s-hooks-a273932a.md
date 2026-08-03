---
source: "https://agent-manager.dev/writing/claude-code-hooks/"
hn_url: "https://news.ycombinator.com/item?id=49156473"
title: "Reading agent status out of Claude Code's hooks"
article_title: "What Claude Code's hooks can't tell you"
author: "yoanwaidev"
captured_at: "2026-08-03T15:38:45Z"
capture_tool: "hn-digest"
hn_id: 49156473
score: 1
comments: 0
posted_at: "2026-08-03T14:42:53Z"
tags:
  - hacker-news
  - translated
---

# Reading agent status out of Claude Code's hooks

- HN: [49156473](https://news.ycombinator.com/item?id=49156473)
- Source: [agent-manager.dev](https://agent-manager.dev/writing/claude-code-hooks/)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T14:42:53Z

## Translation

タイトル: クロード コードのフックからエージェント ステータスを読み取る
記事のタイトル: Claude Code のフックでは分からないこと
説明: クロード コード フック API が決してレポートしない 6 つのこと。ライブ ステータス リストに接続しているときに見つかりました。二重意味の通知、何も起動しない割り込み、および作業継続中の停止停止です。

記事本文:
特長
ドキュメント
比較する
インストール
執筆
Claude Code のフックでは分からないこと
フック API が決して報告しない 6 つのこと、それぞれ同じ方法で見つかりました: 私のステータスの行
自信を持って間違ったことを言ったリスト。
ターミナルに必要なものが 1 つあります。それは、実行中のエージェントのリストを、何も表示せずに確認できることです。
どれが機能しているか、どれがブロックされているか、どれが完了しているかを調べます。
明らかな方法は、画面を読み取ることです。ペインをキャプチャし、正規表現を実行します。それ
どの CLI ツールでも動作するため、私がサポートするすべてのツールのフォールバックとして使用されています。
しかし、それは推測です。スピナーは 1 フレームの間消え、エージェントはアイドル状態に見えます。ツール
スピナーの形をしたものを印刷すると、忙しそうに見えます。
Claude Code にはフックがあるため、Claude セッションでは推測は不要です。それはほとんど
です。これはそうでない部分についてです。
各管理対象セッションは、生成された設定ファイルを次のように渡して起動します。
--settings を実行し、ステータスを書き込むパスという環境変数を 1 つ取得します。
AGENT_MANAGER_STATUS_FILE=/…/hooks/<セッションID>.status
すべてのフックは、そのファイルに単語を書き込む 1 行のシェル コマンドです。
[ -z "$AGENT_MANAGER_STATUS_FILE" ] || printf 動作中 > "$AGENT_MANAGER_STATUS_FILE"
ガードは見た目以上に重要です。クロード コードがその設定ファイルを外部にロードした場合、
管理されたセッションでは、変数は設定されておらず、コマンドは 0 で終了し、何も起こりません。フック
失敗する可能性があるのは、誰かのエージェントを破壊するフックです。
ポーラーはそのファイルを 2 秒ごとに読み取ります。それが幸せな道全体であり、
簡単に言うと、正規表現による推測はなく、正確に正しくなります。それからあなたは
エッジ。
通知には 2 つの異なる意味があります #
これは、エージェントが許可プロンプトでブロックされたときに起動されます。 60秒後にも発火します
アイドル状態の入力ボックス。

クロードがあなたを待っています。同じイベント、最初のみ
1 つは、エージェントが何かで行き詰まっていることを意味します。
額面通りに受け取ると、セッションから離れると、最終的にはすべてのセッションがブロックされたことになります。
修正するには、標準入力のペイロードを読み取り、リマインダーを削除します。
[ -z "$AGENT_MANAGER_STATUS_FILE" ] || grep -q "入力を待っています" \
|| printf 待機中 > "$AGENT_MANAGER_STATUS_FILE"
プレーンテキストの質問では何も起動されません #
エージェントは「テストも更新したほうがいいですか?」と言ってターンを終了しました。フック API にとって、
完了したターン。ファイアを停止すると、ファイルには「 completed 」と表示され、リストには
セッションに緑色のチェックマークが表示され、一言を待つだけの状態になります。
答えてください。
「モデルが散文で何か質問する」というイベントはありません。唯一の場所は、
情報が画面上に表示されます。
ターンを中断してもストップは到着しません。最後に書かれていたのは、
working なので、ファイルには working と表示され、引き続き表示されます
何かを入力するまで動作します。ユーザーが行動するまで間違った状態とは、
ステータスがないよりも悪いです。それはまさに、リストを信頼して立ち去るケースです。
停止はメインエージェントに関するものであり、作業に関するものではありません #
「停止」は、メインループが応答を停止したことを意味します。それが生成したバックグラウンド エージェントは、
まだ研ぎ続けています。つまり、実際の作業は継続しているにもかかわらず、ファイルには「終了」と表示されます。
その作業によりファイルが書き込まれ、「完了」行が表示されますが、そのリポジトリはまだレビュー中です。
逆もまた現れます。バックグラウンド サブエージェントの書き込み処理
PreToolUse と PostToolUse 、およびそれらの Stop を起動しません
終わったら自分のものになる。それをクリアできるものは何もありません。修正しないと、そのファイルは次の場所に固定されたままになります。
永遠に働いています。
起動時、再開時、クリア時、コンパクト時にも起動されます。コンパクトは真ん中で起こる
アクティブなターンの場合、一致しない SessionStart ハンドラーは次のように書き込みます。
非常に作業が多いセッション中アイドル状態になる

NG。マッチャーは明示的である必要があります。
"SessionStart": [{ "matcher": "startup|resume|clear", "hooks": [ … ] }]
クラッシュによりクリーンアップがスキップされる #
SessionEnd はステータスファイルを削除します。クラッシュや SIGKILL は発生しません。
SessionEnd を実行します。ファイルはプロセスを生き残り、自信を持ってレポートを続けます
もう存在しないエージェントの下で働いています。したがって、プロセスを確認する必要があります
独立して、エージェントがいなくなったステータス ファイルは信頼されるのではなく削除される必要があります。
それがデザインの残りの部分 #
フックは第一層のソースであり、フックを修正するためにすべてのポーリングで画面が読み取られます。
ペインに質問、エラー、またはエラーが表示されると、「終了」を示すフックがアップグレードされます。
進行中の作業。ペインにターンが表示されると、作業中であることを示すフックがアップグレードされます
すでに終了しました。
スイッチフックステータス {
ケースのステータス。完了:
一致した場合 && (paneStatus == status.Waiting || paneStatus == status.Errored || paneStatus == status.Working) {
戻りペインのステータス
}
ケースのステータス。作業中:
一致した場合 && (paneStatus == status.Waiting || paneStatus == status.Finished || paneStatus == status.Errored) {
戻りペインのステータス
}
}
リターンフックステータス
このペインは、最新のターンが静かになった場合にのみ休止状態を報告するため、これは決して報告されません
途中で火災が発生します。
私はフックが画面のスクレイピングを置き換えると期待していました。彼らが実際に行ったことは、それを正確にすることです。
フックは、何も出力しないツール呼び出しなど、画面に表示できないことを認識します。
screen は、質問、割り込み、エラーなど、フックが決して起動しないものを認識します。
あなたが見ていない間、どちらか一方だけが信頼できるステータスリストを取得することはできません。
それがすべてのポイントでした。

## Original Extract

Six things the Claude Code hook API never reports, found while wiring it to a live status list: the double-meaning Notification, the interrupt that fires nothing, and Stop landing while the work continues.

Features
Docs
Compare
Install
Writing
What Claude Code's hooks can't tell you
Six things the hook API never reports, each one found the same way: a row in my status
list that confidently said the wrong thing.
I wanted one thing from my terminal: a list of running agents where I can see, without
looking at any of them, which one is working, which one is blocked, and which one is done.
The obvious way is to read the screen. Capture the pane, run regexes over it, guess. That
works with any CLI tool, which is why it is still the fallback for every tool I support.
But it is guessing. A spinner disappears for one frame and the agent looks idle. A tool
prints something spinner-shaped and it looks busy.
Claude Code has hooks, so for Claude sessions the guessing should be unnecessary. It mostly
is. This is about the part that isn't.
Each managed session launches with a generated settings file passed as
--settings , and gets one environment variable: a path to write status into.
AGENT_MANAGER_STATUS_FILE=/…/hooks/<session-id>.status
Every hook is a one-line shell command that writes a word to that file.
[ -z "$AGENT_MANAGER_STATUS_FILE" ] || printf working > "$AGENT_MANAGER_STATUS_FILE"
The guard matters more than it looks. If Claude Code ever loads that settings file outside a
managed session, the variable is unset, the command exits 0, and nothing happens. A hook
that can fail is a hook that breaks someone's agent.
A poller reads that file every two seconds. That is the whole happy path, and for a
straightforward turn it is exactly right, with none of the regex guessing. Then you meet the
edges.
Notification means two different things #
It fires when the agent is blocked on a permission prompt. It also fires after 60 seconds of
an idle input box, as a nudge that Claude is waiting for you. Same event, and only the first
one means the agent is stuck on something.
Take it at face value and every session you walk away from eventually claims to be blocked.
The fix is to read the payload on stdin and drop the reminder.
[ -z "$AGENT_MANAGER_STATUS_FILE" ] || grep -q "waiting for your input" \
|| printf waiting > "$AGENT_MANAGER_STATUS_FILE"
A plain-text question fires nothing #
The agent finishing its turn with "should I also update the tests?" is, to the hook API, a
completed turn. Stop fires, the file says finished , and the list
shows a green checkmark on a session that will sit there forever waiting for a one-word
answer.
There is no event for "the model asked you something in prose". The only place that
information exists is on screen.
Interrupt a turn and no Stop arrives. The last thing written was
working , so the file says working , and it will keep saying
working until you type something. A status that is wrong until the user acts is
worse than no status: it is the exact case where you walk away trusting the list.
Stop is about the main agent, not the work #
Stop means the main loop stopped responding. Background agents it spawned can
still be grinding. So the file says finished while real work continues, and if
that work writes files you get a "done" row whose repo is still changing under review.
The inverse shows up too. Background subagents write working through
PreToolUse and PostToolUse , and fire no Stop of their
own when they end. Nothing ever clears it. Without a correction, that file stays pinned at
working forever.
It fires on startup, on resume, on clear, and also on compact. Compact happens in the middle
of an active turn, so an unmatched SessionStart handler will write
idle over a session that is very much working. The matcher has to be explicit.
"SessionStart": [{ "matcher": "startup|resume|clear", "hooks": [ … ] }]
A crash skips your cleanup #
SessionEnd deletes the status file. A crash or a SIGKILL does not
run SessionEnd . The file survives its process and keeps confidently reporting
working for an agent that no longer exists. So the process has to be checked
independently, and a status file whose agent is gone has to be deleted rather than trusted.
Where that leaves the design #
Hooks are the first-tier source, and the screen is still read on every poll to correct them.
A hook saying finished is upgraded when the pane shows a question, an error, or
ongoing work. A hook saying working is upgraded when the pane shows the turn
already ended.
switch hookStatus {
case status.Finished:
if matched && (paneStatus == status.Waiting || paneStatus == status.Errored || paneStatus == status.Working) {
return paneStatus
}
case status.Working:
if matched && (paneStatus == status.Waiting || paneStatus == status.Finished || paneStatus == status.Errored) {
return paneStatus
}
}
return hookStatus
The pane only reports a resting state once the newest turn has gone quiet, so this never
fires mid-stream.
I expected hooks to replace the screen scraping. What they actually did is make it accurate.
Hooks know things the screen cannot show, like a tool call that prints nothing, and the
screen knows things hooks never fire for, like a question, an interrupt, or an error.
Neither one alone gets a status list you can trust while you are not looking at it, which
was the entire point.
