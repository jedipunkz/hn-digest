---
source: "https://codeaholicguy.com/2026/06/17/i-read-every-claude-code-hook-so-you-dont-have-to/"
hn_url: "https://news.ycombinator.com/item?id=48569091"
title: "I read every Claude Code hook so you don't have to"
article_title: "I read every Claude Code hook so you don’t have to – Codeaholicguy"
author: "hoangnnguyen"
captured_at: "2026-06-17T13:24:25Z"
capture_tool: "hn-digest"
hn_id: 48569091
score: 2
comments: 1
posted_at: "2026-06-17T11:54:17Z"
tags:
  - hacker-news
  - translated
---

# I read every Claude Code hook so you don't have to

- HN: [48569091](https://news.ycombinator.com/item?id=48569091)
- Source: [codeaholicguy.com](https://codeaholicguy.com/2026/06/17/i-read-every-claude-code-hook-so-you-dont-have-to/)
- Score: 2
- Comments: 1
- Posted: 2026-06-17T11:54:17Z

## Translation

タイトル: クロード コードのフックはすべて読んでいますので、読む必要はありません
記事のタイトル: 私はクロード コードのフックをすべて読んでいるので、読む必要はありません – Codeaholicguy
説明: AI DevKit を構築するとき、さまざまなコーディング エージェントのハーネスをさらに詳しく調べる必要がありました。それぞれに独自の構造、ワークフロー、メンタルモデルがあります。この投稿では、Claude Code のフックを詳しく調べて学んだことを共有します。クロード コードには 27 個のフック イベントがあります。私はそれらすべてを経験しました。ほとんどはそうではありません
[切り捨てられた]

記事本文:
私はクロード コードのフックをすべて読んでいますので、あなたが読む必要はありません – Codeaholicguy
コンテンツにスキップ
コードアホリックガイ
私はクロード コードのフックをすべて読んでいますので、あなたが読む必要はありません
AI DevKit を構築するときは、さまざまなコーディング エージェントのハーネスをさらに詳しく調べる必要がありました。それぞれに独自の構造、ワークフロー、メンタル モデルがあります。
この投稿では、Claude Code のフックを詳しく調べて学んだことを共有します。
クロード コードには 27 個のフック イベントがあります。私はそれらすべてを経験しました。
ほとんどは時間をかける価値がありませんが、役立つものはほんの一握りで、ツールの使用方法を変えることができるものは 1 つか 2 つあります。
注意すべき重要な点は、フックを使用すると、ルールを記憶するためにモデルに依存するのをやめることができるということです。代わりに、これらのルールをエージェントの周囲に配置することができます。
.claude/settings.json に次のような内容を記述します。
{ "フック" : { "PreToolUse" : [ { "matcher" : "Bash" , "フック" : [ { "type" : "command" , "command" : "/path/to/hook.sh" , "timeout" : 30 } ] } ] } }
イベントが発生すると、Claude Code はコマンドを実行し、JSON を標準入力にパイプし、JSON を標準出力に読み取ります。
入力には、イベント名、セッション ID、トランスクリプト パス、CWD、およびイベント固有のフィールドが含まれます。
標準出力によって、次に何が起こるかが決まります。 continue: false で中止するか、 Decision: "block" でツールを拒否するか、hookSpecificOutput.AdditionalContext で次のモデル ターンにテキストを挿入できます。
空の stdout で終了 0 を実行しても何も行われません。ゼロ以外の終了は、フック エラーとしてトランスクリプトに表示されます。
すべてのツール呼び出しの前に起動します。 updatedInput を介して入力を承認、ブロック、または書き換えることができます。
入力の書き換えはただ観察することとは異なるため、これが最も重要です。コマンドを実行する前に、コマンドからシークレットを削除できます。許可されたディレクトリ外での terraform apply の実行を拒否できます。
Bash コマンドで認証情報または .env パスを Regex スキャンし、理由を付けてブロックします
アウ

コマンドをブロックするのではなく書き換えることでトークンを編集します
DELETE または DROP TABLE を完全に拒否する
現在のブランチが main でコマンドが git Push の場合、ブロックします
Enter キーを押すと発火します。これは、次のターンのシステム メッセージとして挿入される追加Context を返します。
これは基本的に CLAUDE.md の動的バージョンです。ターンごと、スクリプトが計算できるものすべてに条件付き。
プロンプトに「ユーザー データ」または「PII」が記載されている場合のコンプライアンスのリマインダー
「関連ファイル: A、B、C」を挿入するプロンプトを簡単に grep します。
今日の API 使用量がしきい値を超えた場合は送信を拒否する
ツール呼び出しが成功した後に起動されます。次のターンのために追加のContextを返すことも、モデルが認識する内容を書き換えるためにupdateedMCPToolOutputを返すこともできます。
ここは、「編集後に常にテストを実行する」を強制するのに適しています。
小さなことですが、定期的に発生する古いコンテキストのバグが削除されます。
また、モデルがシークレットを認識する前に、Bash 出力のシークレットを編集する場合にも役立ちます。モデルは動作を続けるためにアクセス トークンを知る必要はありません。
起動時、再開時、クリア時、およびコンパクト後に発生します。 AdditionalContext 、initialUserMessage 、または watchPaths を返すことができます。
ソースフィールドは開始の種類を示すため、再開時と新規ブート時に異なる操作を行うことができます。
watchPaths は過小評価されている部分です。これにより、cwd の外部にパスを登録できるようになり、それらに対して FileChanged が起動されます。これは、共有構成リポジトリが別の場所に存在する場合に便利です。
ツール呼び出しが許可を求めようとしているときに発生します。入力を許可、拒否、書き換えたり、永続的なアクセス許可を更新したりできます。
これにより、完全に bypassPermissions を実行することなく、「すべてを許可する」という反射から解放されます。
ls 、 cat 、 grep 、 git status などを自動的に許可します。その他はすべて UI に任せます。作業リポジトリでは、リポジトリの外に書き込むものはすべて拒否します。秒

リポジトリをクラッチして、さらに許可します。
この値は、同じプロンプトをずっとクリックし続けるのではなく、コードでポリシーをエンコードすることです。
これらは常に接続されているわけではありませんが、状況が必要なときに手を伸ばします。
PreCompact と PostCompact は注目に値します。
圧縮によって保持したいものが破棄されるため、長時間のセッションは時間の経過とともに悪化します。圧縮全体で現在のタスクと最近の決定を固定することは小さな変更ですが、長時間実行される作業を多く行う場合には、目に見える違いが生じます。
これらは存在します。ほとんどの人は決して配線しないでしょう。
StopFailure 、監査専用、停止処理自体がクラッシュしたとき用
SubagentStart 、単独で役立つことはほとんどありませんが、SubagentStop と組み合わせます
セットアップ、/init およびメンテナンスでの起動。 /init テンプレートを出荷する場合に便利です
SessionEnd は REPL の外部で実行され、終了をブロックできず、サイレントに失敗します。クリティカルなクリーンアップをここに置かないでください
CwdChanged 、 FileChanged 、「これを見る」ワークフローのビルディング ブロック。ほとんどの場合、SessionStart.watchPaths とペアになります。
WorktreeCreate 、 WorktreeRemove は、ワークツリー機能に対してのみ起動されます。 worktreePath 応答を使用すると、ワークツリーが作成される場所をオーバーライドできます
ConfigChange は、ディスク上の設定ファイルが変更されるときに起動します。実際には監査のみ
Elitation 、ElicitationResult 、MCP のみ、構造化フォームを要求するサーバー用。毎回同じ質問をするサーバーがある場合に便利です
TeammateIdle 、 TaskCreated 、 TaskCompleted 、Claude Code のチーム機能を使用する場合にのみ関連します
構造化フィールドはhookSpecificOutputに配置されます。古い例や、まだ浮遊している多くの投稿では、AdditionalContext 、 updatedInput 、および watchPaths を応答の最上位に配置します。スキーマが移動しました。フックが何も行っていない場合は、まずそれを確認してください。実際の標準出力を出力して形状を確認します。
PreToolUse ごとに遅延が追加されます

すべてのツール呼び出しに cy を適用します。 macOS でのスポーンコストは約 10 ～ 50 ミリ秒です。 1 ターンで 30 回のツール呼び出しが行われることになります。応答を気にしないログ フックの場合は、ツールが待機しないように async: true と asyncTimeout を設定します。
フック コマンドは完全なシェルで実行されます。プロジェクトの .claude/settings.json は、SessionStart フックを登録できます。 Claude Code は、新しいプロジェクト フックを信頼する前にプロンプ​​トを表示しますが、これは良いことです。しかし、「はい」と言う前にフックを読む価値はあります。信頼できないリポジトリのクローンを作成した場合は、開く前にその .claude/ を確認してください。
私のシェアに共感していただけましたら、私のブログを購読してください。 AI をループに組み込んだ実際のシステムを構築する際に学んだことを共有します。また、X または Threads で私をフォローして、さらに考えたことや進行中の実験を確認することもできます。
Codeaholicguy の詳細をご覧ください
最新の投稿をメールで受け取るには購読してください。
この投稿を気に入っていただけましたら、お気軽にシェアしてください。
Reddit で共有 (新しいウィンドウで開きます)
レディット
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
LinkedIn で共有 (新しいウィンドウで開きます)
リンクトイン
スレッドで共有 (新しいウィンドウで開きます)
スレッド
Telegram で共有 (新しいウィンドウで開きます)
電報
X で共有 (新しいウィンドウで開きます)
×
X で共有 (新しいウィンドウで開きます)
×
AI コーディング エージェントのベスト プラクティス
クロード コードのフック固有の出力
新しい投稿があるたびに通知を受け取るには、メールアドレスを入力してください。
エージェント エンジニアリングを学ぶための実践的なロードマップ
AI コーディング エージェントの説明: ルール、コマンド、スキル、MCP、フック
私はクロード コードのフックをすべて読んでいますので、あなたが読む必要はありません
Claude Code ガイド: Claude Code を 1 年間使用したかのように使用する方法
AIで10倍のエンジニアになる方法
Codeaholicguy の詳細をご覧ください
今すぐ購読して読み続け、完全なアーカイブにアクセスしてください。
購読する
購読しました
コードアホリックガイ
すでに WordPress.com アカウントをお持ちですか?今すぐログインしてください。

## Original Extract

When building AI DevKit, I had to look much deeper into different coding agent harnesses. Each one has its own structure, workflow, and mental model. In this post, I’ll share what I learned from digging into Claude Code hooks. Claude Code has 27 hook events. I went through all of them. Most are not
[truncated]

I read every Claude Code hook so you don’t have to – Codeaholicguy
Skip to content
Codeaholicguy
I read every Claude Code hook so you don’t have to
When building AI DevKit , I had to look much deeper into different coding agent harnesses. Each one has its own structure, workflow , and mental model.
In this post, I’ll share what I learned from digging into Claude Code hooks .
Claude Code has 27 hook events. I went through all of them.
Most are not worth your time, a handful are useful, one or two can change how you use the tool.
The important thing to note is that hooks let you stop relying on the model to remember your rules. Instead, you can put those rules around the agent.
You put something like this in .claude/settings.json :
{ "hooks" : { "PreToolUse" : [ { "matcher" : "Bash" , "hooks" : [ { "type" : "command" , "command" : "/path/to/hook.sh" , "timeout" : 30 } ] } ] } }
When the event fires, Claude Code runs your command, pipes JSON to stdin, and reads JSON back on stdout.
The input includes the event name, session id, transcript path, cwd, and event-specific fields.
Your stdout decides what happens next. You can abort with continue: false, deny the tool with decision: "block" , or inject text into the next model turn with hookSpecificOutput.additionalContext .
Exit 0 with empty stdout is a no-op. Non-zero exits show up in the transcript as a hook error.
Fires before every tool call. It can approve, block, or rewrite the input via updatedInput .
This is the one that matters most, because rewriting input is different from just observing. You can strip secrets out of commands before they run. You can refuse to run terraform apply outside an allowed directory.
Regex scan for credentials or .env paths in Bash commands, then block with a reason
Auto-redact tokens by rewriting the command instead of blocking it
Hard-deny DELETE or DROP TABLE
If the current branch is main and the command is git push , block
Fires when you hit enter. It returns additionalContext , which gets injected as a system message for the next turn.
This is basically a dynamic version of CLAUDE.md . Per-turn, conditional on whatever your script can compute.
Compliance reminders when the prompt mentions “user data” or “PII”
A quick grep over the prompt that injects “relevant files: A, B, C”
Refuse to submit if today’s API spend has crossed a threshold
Fires after a tool call succeeds. It can return additionalContext for the next turn, or updatedMCPToolOutput to rewrite what the model sees.
This is a good place to enforce “always run test after an Edit”.
Small thing, but it removes a recurring class of stale-context bugs.
It is also useful for redacting secrets from Bash output before the model sees them. The model does not need to know your access tokens to keep working.
Fires on startup, resume, clear, and after compact. It can return additionalContext , initialUserMessage , or watchPaths .
The source field tells you what kind of start it is, so you can do different things on resume vs fresh boot.
watchPaths is the underrated part. It lets you register paths outside cwd so FileChanged fires for them. That is handy when a shared config repo lives somewhere else.
Fires when a tool call is about to ask you for permission. It can allow, deny, rewrite the input, or update persistent permissions.
This gets me off the “yolo allow everything” reflex without going full bypassPermissions .
Auto-allow ls , cat , grep , git status , etc. Defer everything else to the UI. In work repos, deny anything that writes outside the repo. In scratch repos, allow more.
The value is encoding the policy in code instead of clicking through the same prompts forever.
These are not always wired up for me, but I reach for them when the situation calls.
PreCompact and PostCompact deserve a callout.
Long sessions get worse over time because compaction throws away things you wanted to keep. Pinning the current task and recent decisions across compaction is a small change, but it makes a visible difference if you do a lot of long-running work.
These exist. Most people will never wire them.
StopFailure , audit-only, for when stop handling itself crashes
SubagentStart , rarely useful alone, but pairs with SubagentStop
Setup , fires on /init and maintenance. Useful if you ship /init templates
SessionEnd , runs outside the REPL, cannot block exit, fails silently. Do not put critical cleanup here
CwdChanged , FileChanged , building blocks for “watch this thing” workflows, mostly paired with SessionStart.watchPaths
WorktreeCreate , WorktreeRemove , only fire for the worktree feature. The worktreePath response lets you override where worktrees get created
ConfigChange , fires when settings files change on disk. Audit-only in practice
Elicitation , ElicitationResult , MCP-only, for servers that request structured forms. Useful if you have a server that asks the same question every time
TeammateIdle , TaskCreated , TaskCompleted , only relevant if you use Claude Code’s team features
hookSpecificOutput is where the structured fields go. Older examples, and a lot of posts still floating around, put additionalContext , updatedInput , and watchPaths at the top level of the response. The schema moved. If your hook is not doing anything, check that first. Print the actual stdout and look at the shape.
Every PreToolUse adds latency to every tool call. Spawn cost on macOS is around 10 to 50ms. In a turn that makes 30 tool calls, that adds up. For logging hooks where you do not care about the response, set async: true and an asyncTimeout so the tool does not wait.
Hook commands run with your full shell. A project’s .claude/settings.json can register a SessionStart hook. Claude Code prompts you before trusting new project hooks, which is good. But it is still worth reading the hook before saying yes. If you ever clone an untrusted repo, look at its .claude/ before opening it.
If my sharing resonates with you, subscribe to my blog. I share what I learn while building real systems with AI in the loop. You can also follow me on X or Threads for more thoughts and ongoing experiments.
Discover more from Codeaholicguy
Subscribe to get the latest posts sent to your email.
If you enjoyed this post, feel free to share it!
Share on Reddit (Opens in new window)
Reddit
Share on Facebook (Opens in new window)
Facebook
Share on LinkedIn (Opens in new window)
LinkedIn
Share on Threads (Opens in new window)
Threads
Share on Telegram (Opens in new window)
Telegram
Share on X (Opens in new window)
X
Share on X (Opens in new window)
X
AI coding agent best practices
Claude Code hookSpecificOutput
Enter your email to get notified whenever there's a new post!
A practical roadmap to learn agentic engineering
AI Coding Agents Explained: Rules, Commands, Skills, MCP, Hooks
I read every Claude Code hook so you don't have to
Claude Code Guide: How to use Claude Code like you've used it for a year
How to become a 10x engineer with AI
Discover more from Codeaholicguy
Subscribe now to keep reading and get access to the full archive.
Subscribe
Subscribed
Codeaholicguy
Already have a WordPress.com account? Log in now.
