---
source: "https://thomas-witt.com/blog/blog-subagent-model-pin/"
hn_url: "https://news.ycombinator.com/item?id=49320976"
title: "I stopped my Claude Code subagents from running on Fable instead of Sonnet"
article_title: "How I stopped my Claude Code subagents from secretly running on Fable instead of Sonnet | Thomas Witt: Tech Entrepreneur & Angel Investor"
author: "thomas_witt"
captured_at: "2026-08-16T16:13:38Z"
capture_tool: "hn-digest"
hn_id: 49320976
score: 1
comments: 0
posted_at: "2026-08-16T15:29:22Z"
tags:
  - hacker-news
  - translated
---

# I stopped my Claude Code subagents from running on Fable instead of Sonnet

- HN: [49320976](https://news.ycombinator.com/item?id=49320976)
- Source: [thomas-witt.com](https://thomas-witt.com/blog/blog-subagent-model-pin/)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T15:29:22Z

## Translation

タイトル: クロード コードのサブエージェントが Sonnet ではなく Fable で実行されるのを停止しました
記事のタイトル: クロード コードのサブエージェントがソネットではなくフェイブルで密かに実行されるのをどのように阻止したか |トーマス・ウィット: テクノロジー起業家兼エンジェル投資家
説明: 私は大規模なセッション モデル、Fable または Opus 5 と、退屈な部分を実行するサブエージェントの小さな動物園を使用して Claude Code を実行しています。ゲートウェイ エージェント、フォーマッタ、チェッカー。エージェントの前の問題として、小さなモデルに一度固定すると、その後は二度と考えなくなる類の機械的なもの。

記事本文:
クロード コードのサブエージェントがソネットではなくフェイブルで密かに実行されるのをどのように阻止したか |トーマス・ウィット: テクノロジー起業家兼エンジェル投資家
私について
ブログ
お問い合わせ
クロード コードのサブエージェントがソネットではなくフェイブルで密かに実行されるのをどのように阻止したか
私は、Fable または Opus 5 という大きなセッション モデルと、退屈な部分を実行するサブエージェントの小さな動物園を使用して Claude Code を実行しています。ゲートウェイ エージェント、フォーマッタ、チェッカー。エージェントの前の問題として、小さなモデルに一度固定すると、その後は二度と考えなくなる類の機械的なもの。
ある時点で、トークンの消費量が、その週に私が実際に何をしていたのかという直感と一致しなくなりました。何も壊れていませんでした。何もエラーはありませんでした。すべてがうまくいきました。必要以上に費用がかかっただけです。
そこで私は、Claude Code が実際にサブエージェントのモデルをどのように選択するかを詳しく調べることにしました。私が信頼していたピンが、信頼すべきではないレイヤーにあることが判明しました。
免責事項: これは私が自分のマシン、自分のプロジェクトで実行しているものです。フックは自分自身のディスパッチをブロックすることができます。それがこのディスパッチの重要な点です。そのため、フックの接続を誤ると、ブロックされたタスク呼び出しを見つめて、なぜなのか疑問に思うことになるでしょう。また、クロード コードの優先順位の動作はリリースごとに異なります。使用しているバージョンを確認してください。
そもそもサブエージェントが必要な理由
愚痴を言い始める前に、一つだけはっきりさせておきたいのは、サブエージェントは優れているということです。この投稿は彼らに対する反論ではなく、彼らを適切に固定するための議論です。
これを使用する理由は、それが小さいモデルだからではありません。それは、独自のコンテキスト ウィンドウで実行されるということです。それは電源を切り、大音量で厄介なことをし、短い応答を返します。メインセッションにノイズが入ることはありません。 4 メガバイトのデータではなく、重要な 3 行が得られます。
どれの

つまり、理想的なサブエージェント ジョブは次のようになります: たくさん取得し、フィルターし、少しだけ返す。そしてその仕事には、優秀ではなく従順なモデルが必要です。ソネットのようなものは一日中それを行います。 Fable や Opus 5 で実行すると、礼儀正しい grep にフロンティア価格を支払うことになります。
フロンティアモデルに決して触れるべきではないエージェントの私のリスト:
AWS、特に CloudWatch Logs。 1 つのリクエスト ID に対するログ クエリは、ocean を返します。スタック トレースとタイムスタンプが必要です。これだけでパターン全体が正当化されます。
GitHub。問題、PR の差分、CI ステータス、「どのコミットがこのファイルに触れたか」。長い出力と小さな応答。
ハニーアナグマ。エラーの発生、バックトレース、「これは先週の火曜日と同じバグですか」。構造化された入力、構造化された出力。
ラングフューズ。トレース ダンプは膨大であり、すべてのトレースの 95% は、それに関する質問とは無関係です。
静的解析: RubyCritic、RuboCop とその仲間たち。報告書は 100 ページあり、実行可能な部分は 9 行です。
テストランナー。 Fable が RSpec の出力を見て、どの 4 つの仕様が赤であるかを知る必要はありません。読めるものが必要です。
ピン留めは、 .claude/agents/cloudwatch-digger.md 内で次のようになります。
---
名前：クラウドウォッチディガー
説明 : CloudWatch Logs にクエリを実行し、関連する行のみを返します
モデル：ソネット
ツール : Bash、Read
---
一行。モデル: ソネット。それがピン全体であり、それがまさにそれが静かに動作を停止するときに痛みを伴う理由です。わざわざピン留めするエージェントは、定義上、最も頻繁に派遣し、少なくとも確認しているエージェントです。
Claude Code は、サブエージェントがどのモデルで実行されるかを次の順序で解決します。
4層。そして、誰もが実際に使用しているフロントマター ピンは、文書化されており、明白で、一度書き込み可能であるため、ランク 3/4 です。
ランク3が常に保持されていれば問題ありません。そうではありません

。いくつかのリリースにわたって、フロントマター層はひっそりと脱落し、ピン留めされたエージェントはランク 4、つまりセッション モデルに一気に落ちました。私の場合はそれが寓話です。
つまり、1 日に 200 回派遣される安価な小さなエージェントは、あなたが所有する最も高価なもので静かに実行されます。そして、私が本当にイライラしているのは、信号が存在しないことです。エラーや警告はなく、トランスクリプトには何も変わっていないように見えます。エージェントはその仕事をします。それは単に価格の何倍かでそれを行うだけです。
クラッシュは礼儀正しい、と教えてくれます。これでは何もわかりません。 4週間後に数字として現れます。
これに遭遇したのは私が最初ではありませんでした。更新後にフロントマターピンが無視されるという一連のアップストリームレポートがあり、人々が確認し続けている回避策は常に同じであり、ディスパッチ時にモデルを明示的に渡すことです。これは、フロントマターの 1 つ上のレイヤーであるランク 2 であり、ランク 2 が破られたレイヤーではありません。
わざわざこれらの問題を複製物とともに提出したすべての人に大きな叫び声を上げます。サイレントコスト回帰は、誰も気づかないため、まさに誰も報告しない種類のバグです。
これには明らかな問題が残ります。「常にモデルを明示的に渡すだけ」ということは、オーケストレーターがそれを毎回、永久に記憶しなければならないことを意味します。物事を確実に永久に記憶するオーケストレーターには、私は出会ったことがありません。
PreToolUse はツール呼び出しが通過する前に実行され、終了コード 2 でブロックできます。 したがって、サブエージェントがフロントマターに固定されており、ディスパッチに明示的なモデルがない場合は、ディスパッチを拒否し、何を再送信するかを正確に指定します。
これらの数行のコードにより、サブエージェントに実際に選択したモデルを使用するようにクロード コードに通知するため、多くのトークンを節約できます。
.claude/hooks/enforce-subagent-model.sh :
#!/bin/bash
{ 読み取り -r T ;読み取り -r S ;読む

-r M ; } < < ( jq -r '.tool_name//"",.tool_input.subagent_type//"",.tool_input.model//""' )
タスク|エージェントの場合 $T ) ;; * ) 0 を終了します ;;イーサック
[ -n " $S " ] && [ -z " $M " ] || 0番出口
P = $( awk '{sub(/\r$/,"")} NR==1&&$0=="---"{f=1;next} f&&$0=="---"{exit} f&&/^model:[ \t]/{gsub(/["' "'" ']/,"",$2);print $2;exit}' \
" ${ CLAUDE_PROJECT_DIR :- . } /.claude/agents/ $S .md" 2>/dev/null )
"" の $P の場合 |継承 ) 0 を終了 ;;イーサック
echo "ブロックされました: ' $S ' は固定されています (モデル: $P ) が、このディスパッチには明示的な 'モデル' がありません。モデル \" $P \" で再ディスパッチします。意図的に異なるモデルも合格しますが、それは明示的である必要があります。" > &2
出口2
9行。 jq は標準入力からフック ペイロードを読み取り、awk はエージェントのフロントマターからピンを読み取り、stderr 上のメッセージはクロード コードに戻り、クロード コードが自動的に正しく再ディスパッチされます。
chmod +x .claude/hooks/enforce-subagent-model.sh
フェーズ 2: 配線する
{
「フック」: {
"PreToolUse" : [
{
"matcher" : "タスク" ,
「フック」: [
{ "タイプ" : "コマンド" , "コマンド" : "$CLAUDE_PROJECT_DIR/.claude/hooks/enforce-subagent-model.sh" }
】
}
】
}
}
モデルなしで固定エージェントをディスパッチしてテストします。 BLOCKED メッセージが表示され、次のディスパッチで PIN が送信されるはずです。
意図的に透過するもの
ガードレールの興味深い部分は、その例外です。
明示的なモデルを含むすべてのディスパッチ (異なるモデルであっても)。私が守っているのは選択ではなく省略です。意図的に小規模なエージェントを大規模なモデルに派遣する場合、それは決定であり、決定は許可されます。偶然モデルに陥ることはありません。
固定解除されたエージェント。モデル: 継承またはモデル キーがまったくないことは、継承が目的であることを意味します。いいよ、パス。
組み込みタイプ (Explore、Plan、汎用など)。 .claude/agents にファイルはなく、何も固定されておらず、強制するものもありません。
何でも解析解除

できる。 jq がペイロードを読み取れない場合、3 つの変数はすべて空に戻り、フックは 0 で終了します。
最後の 1 つは意図的なものです。これはコストのガードレールであり、セキュリティの境界ではありません。理解できないペイロードで失敗して閉じられたフックは、最終的に考えられる最悪の瞬間にセッションを中断することになるため、フックを無効にすると、最初の場所に戻ることになります。
ちなみに、空または null のモデルは選択肢として扱われません。それは、疑わしいレイヤーであるフロントマターに直接到達するため、存在しないレイヤーのようにゲートされます。
ワークフロー ツールの内部 Agent() の生成は PreToolUse を経由しません。このフックは、タスクとエージェントのディスパッチのみを対象としています。ワークフロー ツール内で生成されたものはすべて、ワークフロー ツールを通り過ぎていきます。代わりに、散文で説明されているものについては参考資料で説明します。これは、それらがカバーされていないと言う良い方法です。ゲートがどこで終わるのかを知ってください。
Fable または Opus 5 でクロード コードを実行していて、トークン グラフが週に感じたより急勾配になっているように見える場合は、他のものを確認する前に固定エージェントを確認してください。ここでの故障モードはそれ自体を発表しないため、安価であると思われるものはすべて検証する価値があります。黙って請求するだけです。
まだサブエージェントをまったく使用していない場合は、それはまったく別のトピックであり、正直に言って、この投稿よりも重要なトピックです。読みたい場合はお知らせください。書きます。
ワークフローとツールのギャップを埋めるためのよりクリーンな方法を見つけたら、私にも知らせてください。

## Original Extract

I run Claude Code with a big session model, Fable or Opus 5, plus a small zoo of subagents doing the boring parts. Gateway agents, formatters, checkers. The kind of mechanical stuff you pin to a small model once, in the agent’s frontmatter, and then never think about again.

How I stopped my Claude Code subagents from secretly running on Fable instead of Sonnet | Thomas Witt: Tech Entrepreneur & Angel Investor
About me
Blog
Contact
How I stopped my Claude Code subagents from secretly running on Fable instead of Sonnet
I run Claude Code with a big session model, Fable or Opus 5, plus a small zoo of subagents doing the boring parts. Gateway agents, formatters, checkers. The kind of mechanical stuff you pin to a small model once, in the agent’s frontmatter, and then never think about again.
At some point the token consumption stopped matching my gut feeling of what I’d actually been doing that week. Nothing was broken. Nothing errored. Everything worked. It just cost more than it should have.
So I decided to take a deeper look at how Claude Code actually picks the model for a subagent. It turned out the pin I’d been trusting sits in the one layer you shouldn’t trust.
Disclaimer: This is what I run on my own machine, on my own projects. Hooks can block your own dispatches, that’s the whole point of this one, so if you wire it up wrong you’ll be staring at a blocked Task call and wondering why. Also, precedence behaviour in Claude Code changes between releases. Verify against the version you’re on.
Why you want subagents in the first place
Before the complaining starts, let me be clear about one thing: subagents are good. This post is not an argument against them, it’s an argument for pinning them properly.
The reason to use one isn’t that it’s a smaller model. It’s that it runs in its own context window . It goes off, does something loud and messy, and hands back a short answer. The noise never lands in your main session. You get the three lines that matter instead of the four megabytes they came from.
Which means the ideal subagent job looks like this: fetch a lot, filter, return a little . And that job needs a model that is obedient, not brilliant. Something like Sonnet does it all day. Running it on Fable or Opus 5 is paying frontier prices for grep with good manners.
My list of agents that should never touch a frontier model:
AWS, especially CloudWatch Logs. A log query for one request ID returns an ocean. You want the stack trace and the timestamp. This one alone justifies the whole pattern.
GitHub. Issues, PR diffs, CI status, “which commit touched this file”. Long output, small answer.
Honeybadger. Error occurrences, backtraces, “is this the same bug as last Tuesday”. Structured input, structured output.
Langfuse. Trace dumps are enormous and 95% of every trace is irrelevant to the question you’re asking about it.
Static analysis: RubyCritic, RuboCop and friends. The report is a hundred pages, the actionable part is nine lines.
Test runners. You do not need Fable to look at RSpec output and tell you which four specs are red. You need something that can read.
Pinning one looks like this, in .claude/agents/cloudwatch-digger.md :
---
name : cloudwatch-digger
description : Queries CloudWatch Logs and returns only the relevant lines
model : sonnet
tools : Bash, Read
---
One line. model: sonnet . That’s the whole pin, and that’s exactly why it hurts when it silently stops working. The agents you bother to pin are, by definition, the ones you dispatch most often and look at least.
Claude Code resolves which model a subagent runs on in this order:
Four layers. And the one that everybody actually uses, the frontmatter pin, because it’s the one that’s documented, obvious and writable once, is rank 3 of 4.
That would be fine if rank 3 always held. It doesn’t. Across several releases the frontmatter layer has silently dropped out, and pinned agents fell straight through to rank 4: the session model. Which in my case is Fable.
So the cheap little agent you dispatch two hundred times a day quietly runs on the most expensive thing you have. And here’s the part I find genuinely annoying: there is no signal. No error, no warning, nothing in the transcript that looks different. The agent does its job. It just does it at a multiple of the price.
A crash is polite, it tells you. This doesn’t tell you anything. It shows up four weeks later as a number.
I wasn’t the first one to run into this. There’s a whole class of upstream reports about frontmatter pins being ignored after an update, and the workaround people keep confirming is always the same: pass the model explicitly on the dispatch. That’s rank 2, one layer above frontmatter, and rank 2 has never been the layer that breaks.
Big shoutout to everyone who bothered to file those issues with reproductions. Silent cost regressions are exactly the kind of bug nobody files, because nobody notices.
Which leaves an obvious problem: “just always pass the model explicitly” means the orchestrator has to remember it, every single time, forever. An orchestrator that reliably remembers a thing forever is not something I’ve met.
PreToolUse runs before a tool call goes through and can block it with exit code 2. So: if a subagent is pinned in its frontmatter, and the dispatch carries no explicit model , refuse the dispatch and say exactly what to re-send.
These few lines of code can save you a lot of tokens, because they remind Claude Code to use the model you actually chose for your subagents.
.claude/hooks/enforce-subagent-model.sh :
#!/bin/bash
{ read -r T ; read -r S ; read -r M ; } < < ( jq -r '.tool_name//"",.tool_input.subagent_type//"",.tool_input.model//""' )
case $T in Task|Agent ) ;; * ) exit 0 ;; esac
[ -n " $S " ] && [ -z " $M " ] || exit 0
P = $( awk '{sub(/\r$/,"")} NR==1&&$0=="---"{f=1;next} f&&$0=="---"{exit} f&&/^model:[ \t]/{gsub(/["' "'" ']/,"",$2);print $2;exit}' \
" ${ CLAUDE_PROJECT_DIR :- . } /.claude/agents/ $S .md" 2>/dev/null )
case $P in "" | inherit ) exit 0 ;; esac
echo "BLOCKED: ' $S ' is pinned (model: $P ) but this dispatch has no explicit 'model'. Re-dispatch with model: \" $P \" . A deliberate different model also passes, but it must be explicit." > &2
exit 2
Nine lines. jq reads the hook payload from stdin, awk reads the pin out of the agent’s frontmatter, and the message on stderr goes back to Claude Code, which then re-dispatches correctly on its own.
chmod +x .claude/hooks/enforce-subagent-model.sh
Phase 2: Wiring it up
{
"hooks" : {
"PreToolUse" : [
{
"matcher" : "Task" ,
"hooks" : [
{ "type" : "command" , "command" : "$CLAUDE_PROJECT_DIR/.claude/hooks/enforce-subagent-model.sh" }
]
}
]
}
}
Test it by dispatching a pinned agent without a model. You should get the BLOCKED message, and the very next dispatch should carry the pin.
What it lets through, on purpose
The interesting part of a guardrail is its exceptions:
Any dispatch that carries an explicit model, even a different one. What I’m guarding against is omission , not choice. If I deliberately send a small agent to a big model, that’s a decision, and decisions are allowed. Falling into a model by accident is not.
Unpinned agents. model: inherit or no model key at all means inheritance is the intent. Fine, pass.
Built-in types (Explore, Plan, general-purpose, …). No file in .claude/agents , nothing pinned, nothing to enforce.
Anything unparseable. If jq can’t read the payload, all three variables come back empty and the hook exits 0.
That last one is deliberate: this is a cost guardrail, not a security boundary. A hook that fails closed on a payload it doesn’t understand will eventually wedge a session at the worst possible moment, and then I’ll disable it, and then I’ll be back where I started.
An empty or null model, by the way, is not treated as a choice. It falls straight through to frontmatter, the exact layer under suspicion, so it gets gated like an absent one.
Workflow-tool internal agent() spawns don’t go through PreToolUse. The hook covers Task and Agent dispatches only. Anything spawned inside a workflow tool sails right past it. I cover those with prose in my reference docs instead, which is a nice way of saying they’re not covered. Know where your gate ends.
If you’re running Claude Code on Fable or Opus 5 and your token graph looks steeper than your week felt, check your pinned agents before you check anything else. Everything that’s supposed to be cheap is worth verifying, because the failure mode here doesn’t announce itself. It just quietly bills you.
And if you’re not using subagents at all yet, that’s a whole other topic, and honestly a bigger one than this post. Let me know if you’d like to read it and I’ll write it up.
If you find a cleaner way to close the workflow-tool gap, let me know as well!
