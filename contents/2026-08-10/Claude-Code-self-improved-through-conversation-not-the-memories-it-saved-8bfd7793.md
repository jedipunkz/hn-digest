---
source: "https://shojin.dev/blog/claude-code-improved-through-conversation-not-memories/"
hn_url: "https://news.ycombinator.com/item?id=49246281"
title: "Claude Code self-improved through conversation, not the memories it saved"
article_title: "Claude Code self-improved on business workflows through conversation, not memories · shōjin"
author: "anndvision"
captured_at: "2026-08-10T17:45:00Z"
capture_tool: "hn-digest"
hn_id: 49246281
score: 2
comments: 2
posted_at: "2026-08-10T16:49:46Z"
tags:
  - hacker-news
  - translated
---

# Claude Code self-improved through conversation, not the memories it saved

- HN: [49246281](https://news.ycombinator.com/item?id=49246281)
- Source: [shojin.dev](https://shojin.dev/blog/claude-code-improved-through-conversation-not-memories/)
- Score: 2
- Comments: 2
- Posted: 2026-08-10T16:49:46Z

## Translation

タイトル: クロード・コードは保存された記憶ではなく、会話を通じて自己改善した
記事タイトル：クロード・コード、記憶ではなく会話でビジネスワークフローを自己改善・shojin
説明: クロード コードは記憶ではなく会話を通じてビジネス ワークフローを自己改善しました

記事本文:
書く
クロード・コードは記憶ではなく会話を通じてビジネス ワークフローを自己改善した
Claude Code は、メモリの保存、コードの作成と実行、サブエージェントの生成、Web の検索を行うことができます。
これらのアフォーダンスはセッション全体にわたって持続し、動的に更新できるため、自己改善の手段となります。
コード (Opus 5、xhigh 努力) に、シミュレートされたビジネス ワークフロー サーバーからタスクの説明 (スプレッドシート、CRM レコード、電子メール、チケット発行) を取得する方法を指示し、改善するように指示しました。
その結果、成功は 34% から 44% に上昇しました。
もっと良くなる。
MCP サーバーの「カリキュラム」は、一連のタスクを提供します。
外部ループは何もありません - 自分で動かしてください:
1. `get_task` を呼び出すと、次のタスクがプルされます。
{done: true} が返された場合、ストリームは使い果たされ、停止します。
2. リストされているツールを使用してタスクを完了できます。
そして「done」を呼び出すとスコアリングのために送信されます。
3. ストリームがなくなるまで続けます。
会話履歴以外で生成された成果物は記憶メモだけだったので、それらが改善の原動力であると考えました。
ただし、それらを含めることは、保留されたタスクの成功に測定可能な影響を与えませんでした。
代わりに、会話によって自己改善が促進されました。つまり、タスクの実行トレースと圧縮後の概要です。
クロード・コードは記憶を使用して指向性のある知識ベースを構築しました
ビジネス ワークフロー サーバーからのタスクの実行中に、84 件の書き込みと編集が行われました。
どれもメモリファイルにありました。
Claude Code は、コードを書いたり、サブエージェントを生成したり、Web を検索したりすることはありませんでした。
その結果、指向性のある知識ベースが誕生しました。
上のグラフの各ノードは、 MEMORY.md によってインデックス付けされたノートに対応します。
各ノートは、名前、1 行の説明、およびメタデータの下にネストされたタイプを含む前付で始まります。
少なくとも 1 つのエピソードと受け取ったスコアを要約します。
9つの音符の要約

複数のエピソードをそれぞれ 2 ～ 4 つずつまとめました。
各ノートは、他のノートを指す関連行で終わります。
33 件はフィードバックのタイプで、「適用方法」セクションが含まれています。 24 には「なぜ」セクションも含まれています。
1 つはタイプ参照のもので、これらのセクションは含まれていません。
分類法の他の 2 つのタイプ ( user と project ) は使用されなくなりました。
思い出ではなく会話が改善を促進した
このセッションでは、記憶ファイルに加えて、会話の痕跡も生成されました。
会話トレースの任意の時点から再開されたセッションには、その時点までの次のコンテキストが含まれます。
メモリ ファイル インデックスが挿入されたシステム プロンプト
最後の圧縮時に書かれた要約、または最初のストレッチでは何も書かれない
実行されたすべてのタスクと受け取ったフィードバック
綴られた思い出の記録
改善がどこから来るのかという問題は真っ二つに分かれます。
1) 再開されたセッション コンテキストのサブセットでは、どの程度の改善が残っていますか?
2) そのコンテキストに冗長な要素はありますか?
最初の質問は、新しいセッションを段階的に渡すことで答えられます。最初は何もせず、次に圧縮後に要約を書き、次に最初と最初の圧縮または 2 つの圧縮イベントの間の完全なコンテキストを渡します。
2 つ目は、会話のコンテキストから 1 つのもの (記憶ファイルのインデックスまたは思考の内容) を削除することです。
各バージョンは、同じ 120 個の保留タスクで検証されます。
各タスクは、同じツールを使用して同じモデル上で独自のセッションで実行されるため、あるタスクから次のタスクに何も引き継がれません。
セッションは、4 つの圧縮イベントのそれぞれの開始時とその前後の 9 つの時点で再開されます。
以下のフォレスト プロットの各行は、1 つのコンテキスト設定と何もない設定の違いを、その設定が存在する再開ポイント全体で集計して示しています。

。
-5
+14
ターン
-7
+5
コスト
$+0
$+7
会話
メモリファイルなしで
思考内容がなければ
圧縮の概要
メモリファイルなしで
メモリファイルだけ
+9.9 [+5.9, +13.8] -5.6 [-6.5, -4.8] $+6.33 [+6.06, +6.60]
+8.8 [+4.9, +12.8] -3.9 [-4.9, -3.0] $+7.12 [+6.82, +7.43]
+8.2 [+4.2, +12.2] -6.0 [-6.9, -5.2] $+6.00 [+5.70, +6.30]
+5.4 [+1.8, +9.0] +3.8 [+3.0, +4.6] $+0.14 [+0.11, +0.17]
+8.5 [+5.0, +12.0] +3.2 [+2.4, +4.0] $+0.11 [+0.08, +0.14]
+1.7 [-5.3, +8.6] +3.4 [+2.2, +4.6] $+0.11 [+0.07, +0.15]
上の図と同じ 6 行を数字で示します。各開始状態の影響
何もせずに起動したセッションに対して、同じ 120 個の保留タスクで、95%
信頼区間。
開始状態のタスクが解決され、パーセントポイントがコストに変わります
会話 +9.9 [+5.9, +13.8] -5.6 [-6.5, -4.8] +$6.33 [+$6.06, +$6.60]
記憶ファイルなしの会話 +8.8 [+4.9, +12.8] -3.9 [-4.9, -3.0] +$7.12 [+$6.82, +$7.43]
思考内容のない会話 +8.2 [+4.2, +12.2] -6.0 [-6.9, -5.2] +$6.00 [+$5.70, +$6.30]
圧縮の概要 +5.4 [+1.8, +9.0] +3.8 [+3.0, +4.6] +$0.14 [+$0.11, +$0.17]
メモリ ファイルを含まない圧縮の概要 +8.5 [+5.0, +12.0] +3.2 [+2.4, +4.0] +$0.11 [+$0.08, +$0.14]
メモリファイルのみ +1.7 [−5.3, +8.6] +3.4 [+2.2, +4.6] +$0.11 [+$0.07, +$0.15]
新しいセッションがセッションに対してそれぞれのことを起動するときに解決するもの、かかるもの、コスト
同じ120の保留されたタスクについては、何も渡されませんでした。 95% 信頼区間による点推定。
黒丸はゼロを除いた区間です。インデントされた行は、上の行に 1 つのことがある行です
取り出した。何も渡されないセッションでは 120 件中 41 件が解決され、29.8 ターンかかり、コストは 0.86 ドルです。
メモリファイルの有無にかかわらず、すべての会話設定はゼロをクリアします

es、思考内容の有無にかかわらず。
完全な会話コンテンツにはドルがかかりますが、各圧縮サマリーも数セントでクリアされます。
メモリ ファイルだけが、間隔の空いた 1 つの点です。
[切り捨てられた]
一緒に読むと、設定は所有権を絞り込みます。
タスクの記録を保持するすべての設定はゼロにクリアされ、要約 (会話が圧縮された) を保持するすべての設定も同様です。
どちらも保持しない設定、つまりメモリ ファイルのみを保持し、何も保持しない設定は失敗します。
思考内容が追加するものは実行全体にわたっており、会話自体が最も強い 4 番目の境界で最小になります。
したがって、改善は会話にあり、要約すればそのほとんどが数セントで済みます。
追加のドルで得られるものは、答えへのステップが減ることです。
記憶ファイルは冗長であり、セッションが進行するにつれて思考内容の影響は減衰します。
この図は 4 つの圧縮境界をプールしています。一つ一つ見てみると、走りは平坦ではありません。
この会話は、+10.0、+7.8、+7.5、+14.2 ポイントのすべての境界でコールド スタートを上回っています。
間隔が重なっているため、間隔を隔てるものは何もありません。
重要なのは、最後のタスクでも効果が減っていないということです。タスク 194 での会話で開始されたセッションでは、120 件中 58 件が解決されています。
最初の 3 つの境界については、要約は置換された会話と歩調を合わせており、会話の 90 万トークンに対して数千トークンで 1.7、3.4、および 0.8 ポイントを譲り、いずれもゼロと区別できません。
4 回目では、11.7、95% CI [-19.8、-3.6] になります。
会話によって解決されたタスクのうち 20 はその概要に残っておらず、残ったタスクはコールド スタートよりも +2.5 の価値があり、新たに開始する場合と区別できません。
その時の会話

ダリーは以前のどの職場でも同じように働いていました。そこから書かれた要約はそうではありませんでした。
要約は、圧縮された会話よりも改善の堅牢性が劣る可能性があります。
思考内容が逆方向に流れてしまいます。
4 つの境界全体で +4.2、+2.6、+0.8、および -0.8 ポイントの価値があり、着実にゼロに下がります。
これはトークンによる会話の 5.6% から 8.1% の間であり、最後の境界までに会話の残りの部分には、初期に考えが寄与していたものがすべて含まれているように見えます。
プールされたメモリの数値には、言及する価値のある分割が隠されています。
メモリ ファイルは、会話にマウントされているときは中立です。
また、単独でマウントするときも中立であり、会話や要約はありません。
要約に基づいてそれらは約 3 ポイントかかり、4 つの境界線すべてがその方向を指します。
このデータの分割は、これらの違いを見た上で選択されたものであるため、結果ではなくヒントとして読む必要があります。
エージェントがメモを使って何をしたか
ファイルはすべてのメモリ アームにマウントされ、インデックスはシステム プロンプトに接続されているため、セッションが探しているかどうかに関係なく、メモはそこにありました。
読んだよりもはるかに多くのことを書いていました。
1,080 回の実行中、セッションは 545 回でメモリに書き込み、101 回でノート本体を開きましたが、プロンプトで自身の設定を変更しないように指示されていました。
各プローブはナレッジ ベースの独自のコピーを取得するため、プローブが終了すると、これらの書き込みはすべて破棄されます。
したがって、ヌルは配管ではなく、摂取に関するものです。
インデックスがシステム プロンプト内にあり、本体が 10 回に 1 回未満の実行で開かれるナレッジ ベースには何の価値もありません。
使われなかった自己改善の道
未使用のアフォーダンスは欠落しておらず、エージェントはそれらの使用を禁止されていませんでした。
コード実行、サブエージェント、Web 検索が利用可能であることが検証されました

実行全体にわたって、モデルを Sonnet から Opus に移動すること、工数を low から xhigh に増やすこと、プロンプトで明示的に許可を与えることの 3 つの介入では、それらを引き出すことができませんでした。
タスクを完了するためにエージェント SDK を使用して構築するなど、より野心的な探索が行われれば興奮するでしょうが、いずれも実現しませんでした。
タスクは一度に 1 つずつ到着し、次のタスクが分配される前に完了する必要がありました。
その体制下では、サブエージェントはほとんど購入せず、分析対象となる実行中の 2 番目のタスクがないため、タスク間分析用にトレースを保存しても得られるものはほとんどない可能性があります。
私は、Fable 5 を最大限の推論と、同時に分配するタスク サーバーでこれらの動作が現れるかどうかを確認することに興味があります。
クロード・コードでは明らかにされない潜在的なアフォーダンス
スキルは作成できますが、スキルをロードするにはセッションのリロードが必要で、エージェント自体がトリガーすることはできません。
これを書いても、それを書いた実行には影響しません。
独自のツールは修正されています。MCP 構成は --strict-mcp-config の下に読み取り専用でマウントされているため、コードは独自のサーバーに新しいツールを登録できません。
モデルは初期化時に claude-opus-5 に設定されており、コードは /model スキルにアクセスできません。
サブエージェント ツールはモデル引数を受け取るため、コードは委任された作業に対してより強力な、またはより安価なモデルを選択できた可能性があります。
Anthropic モデルには微調整 API がないため、重みの更新も閉じられます。
推論エフォートも初期化時に実行全体にわたって xhigh に設定され、コードは自律的に /effort スキルを呼び出すことができません。
モデルとは異なり、これはサブエージェント ルートによっても閉じられます。サブエージェント ツールにはエフォート パラメーターがないため、サブエージェントはセッションの設定を継承します。
エージェントは、難しいタスクについてはより深く考えるか、簡単なタスクについては安く考えるかを選択できませんでした。
コンテキスト圧縮が t にわたって 4 回実行されました

コンテキストが満たされると、自動的に実行されます。
コードはそれをスケジュールせず、延期することもできず、何が生き残るかを選択することもできませんでした。
圧縮の概要が改善に貢献したことを考えると、いつ圧縮するか、何を保持するかを決定できるエージェントは実行する価値のある実験です。
Pi 、Hermes 、および Prime Agent はオープンソースのエージェント ハーネスであるため、Claude Code が公開していないアフォーダンスは、利用できないものとしてリストされるのではなく、公開されて測定される可能性があります。
制作上のフィードバックが密ではない
すべての提出物は採点され、すべてのトレーニング タスクについて、数値による報酬と成功フラグというスコアが直接返されます。
それが結果全体の基礎となる条件です。
実際のデプロイメントがこのようになることはほとんどありません。
CRM レコードを更新するエージェントは、下流の担当者が気づくか、四半期レポートが間違っていることが判明するまで、アカウントが間違っていたことに気づかず、フィードバックをまったく受け取らないこともよくあります。
未解決の問題は、フィードバック信号が減少したときにどのような改善が残るかということです。
ここでのすべての数値は、すべての試技の即時採点に基づいて測定されました。
10 回中 1 つのタスクで得点したラン、または 1 週間後に得点したランは、改善が少ないかまったく改善しない可能性があり、それを伝えるコンテキストの部分が同じ部分である必要はありません。
思い出はメカじゃなかった

[切り捨てられた]

## Original Extract

Claude Code self-improved on business workflows through conversation, not memories

writing
Claude Code self-improved on business workflows through conversation, not memories
Claude Code can save memories, write and execute code, spawn subagents, and search the web.
These affordances are avenues for self-improvement because they can persist and be dynamically updated across a session.
I instructed Code (Opus 5, xhigh effort) how to pull task descriptions (spreadsheets, CRM records, email, ticketing) from a simulated business workflow server and told it to get better.
It did: held-out success rose from 34% to 44%.
Get Better.
The MCP server `curriculum` provides a stream of tasks.
Nothing external loops for you - drive yourself:
1. Calling `get_task` pulls the next task.
If it returns {done: true}, the stream is exhausted, stop.
2. The task can be completed using the tools listed,
and calling `done` submits it for scoring.
3. Continue until the stream is exhausted.
The only artifacts produced outside of the conversation history were memory notes, so I assumed that they were a driver of improvement.
However, their inclusion had no measurable effect on held-out task success.
Instead, conversation drove the self-improvement: the task execution traces and post-compaction summaries.
Claude Code used memories to build a directed knowledge base
Eighty-four writes and edits were made during the execution of tasks from the business workflow server.
Every one was on a memory file.
Claude Code never wrote code, spawned a subagent, or searched the web.
The result was a directed knowledge base.
Each node in the graph above corresponds to a note indexed by MEMORY.md .
Each note opens with frontmatter carrying a name , a one-line description , and a type nested under metadata .
They summarize at least one episode and the score received.
Nine notes summarized multiple episodes, two to four each.
Each note ends with a Related line pointing to other notes.
Thirty-three are of type feedback and contain a How to apply section; twenty-four also include a Why section.
One is of type reference and does not include these sections.
The taxonomy's other two types ( user and project ) went unused.
Conversation drove the improvement, not memories
In addition to the memory files, the session also produced a conversation trace.
A session resumed from any point in the conversation trace contains the following context up to that point:
the system prompt with memory file index inserted
the summary written at the last compaction, or nothing for the first stretch
every task executed and the feedback received
the record of the memories written
The question of where the improvement comes from splits in two.
1) How much of the improvement survives in any subset of the resumed session context?
2) Are any elements of that context redundant?
The first question is answered by handing a new session progressively more: first nothing, then a summary written after compaction, then the full context between either the beginning and the first compaction or two compaction events.
The second by removing one thing from the conversation context, either the memory file index or the thought content.
Each version is validated on the same 120 held-out tasks.
Each task is run in its own session, on the same model with the same tools, so nothing carries from one task to the next.
Sessions are resumed at nine points: the beginning, and before and after each of the four compaction events.
Each row in the forest plot below shows the difference between one context setting and the setting with nothing, aggregated across the resumption points where that setting exists.
-5
+14
turns
-7
+5
cost
$+0
$+7
the conversation
without the memory files
without the thought content
the compaction summary
without the memory files
the memory files alone
+9.9 [+5.9, +13.8] -5.6 [-6.5, -4.8] $+6.33 [+6.06, +6.60]
+8.8 [+4.9, +12.8] -3.9 [-4.9, -3.0] $+7.12 [+6.82, +7.43]
+8.2 [+4.2, +12.2] -6.0 [-6.9, -5.2] $+6.00 [+5.70, +6.30]
+5.4 [+1.8, +9.0] +3.8 [+3.0, +4.6] $+0.14 [+0.11, +0.17]
+8.5 [+5.0, +12.0] +3.2 [+2.4, +4.0] $+0.11 [+0.08, +0.14]
+1.7 [-5.3, +8.6] +3.4 [+2.2, +4.6] $+0.11 [+0.07, +0.15]
The same six rows as the figure above, as numbers. Effect of each starting state
against a session that boots with nothing, on the same 120 held-out tasks, with 95%
confidence intervals.
starting state tasks solved, percentage points turns cost
the conversation +9.9 [+5.9, +13.8] −5.6 [−6.5, −4.8] +$6.33 [+$6.06, +$6.60]
the conversation without the memory files +8.8 [+4.9, +12.8] −3.9 [−4.9, −3.0] +$7.12 [+$6.82, +$7.43]
the conversation without the thought content +8.2 [+4.2, +12.2] −6.0 [−6.9, −5.2] +$6.00 [+$5.70, +$6.30]
the compaction summary +5.4 [+1.8, +9.0] +3.8 [+3.0, +4.6] +$0.14 [+$0.11, +$0.17]
the compaction summary without the memory files +8.5 [+5.0, +12.0] +3.2 [+2.4, +4.0] +$0.11 [+$0.08, +$0.14]
the memory files alone +1.7 [−5.3, +8.6] +3.4 [+2.2, +4.6] +$0.11 [+$0.07, +$0.15]
What a fresh session solves, takes and costs when it boots with each thing, against a session
handed nothing, on the same 120 held-out tasks. Point estimates with 95% confidence intervals;
filled circles are intervals that exclude zero. Indented rows are the row above with one thing
taken out. A session handed nothing solves 41 of the 120, takes 29.8 turns and costs $0.86.
Every conversation setting clears zero, with or without the memory files and with or without the thought content.
Each compaction summary clears it too, for cents where the full conversation content costs dollars.
The memory files alone are the one hollow point with an interval that does not
[truncated]
The settings narrow ownership when read together.
Every setting that keeps the task record clears zero, and so does every setting that keeps the summary, which is the conversation compressed.
The settings that keep neither, the memory files alone and nothing at all, are the ones that fail.
What the thought content adds falls across the run, and is smallest at the fourth boundary, where the conversation itself is strongest.
So the improvement is in the conversation, and a summary keeps most of it for cents.
What the extra dollars buy is fewer steps to the answer.
The memory files are redundant and the influence of thought content decays as the session progresses.
The figure pools the four compaction boundaries. Taken one at a time, the run is not flat.
The conversation beats a cold start at every boundary: +10.0, +7.8, +7.5, and +14.2 points.
The intervals overlap, so nothing separates them.
What matters is that the effect is undiminished at the last one, where a session booted with the conversation as it stood at task 194 solves 58 of the 120.
For the first three boundaries the summary keeps pace with the conversation it replaced, giving up 1.7, 3.4, and 0.8 points, none of them distinguishable from zero, on a few thousand tokens against the conversation's nine hundred thousand.
At the fourth it gives up 11.7, 95% CI [−19.8, −3.6].
Twenty of the tasks that conversation solved did not survive its summary, and what did survive is worth +2.5 over a cold start, which is not distinguishable from starting fresh.
The conversation at that boundary was working as well as at any earlier one. The summary written from it was not.
The summary may be a less robust carrier of the improvement than the conversation it compresses.
The thought content runs the other way.
It is worth +4.2, +2.6, +0.8, and −0.8 points across the four boundaries, a steady fall to nothing.
It is between 5.6% and 8.1% of the conversation by tokens, and by the last boundary the rest of the conversation appears to carry whatever the thoughts were contributing early.
The pooled memory figure conceals a split worth stating.
The memory files are neutral when mounted on a conversation.
They are also neutral when mounted on their own, with no conversation and no summary.
Mounted on a summary they cost about three points, and all four boundaries point that way.
This division of the data was chosen after seeing these differences, so it should be read as a hint rather than a result.
What the agent did with the notes
The files were mounted in every memory arm and the index was spliced into the system prompt, so the notes were there whether or not the session went looking.
It wrote far more than it read.
Across the 1,080 runs that had them, the session wrote to memory in 545 and opened a note body in 101, having been told in the prompt not to modify its own setup.
Each probe gets its own copy of the knowledge base, so every one of those writes was discarded when the probe ended.
So the null is about uptake, not plumbing.
A knowledge base whose index sat in the system prompt, and whose bodies were opened in fewer than one run in ten, was worth nothing.
Avenues of self-improvement that were not used
The unused affordances were not missing and the agent was not forbidden to use them.
Code execution, subagents, and web search were validated as available for the whole run, and three interventions failed to elicit them: moving the model from Sonnet to Opus, raising the effort from low to xhigh , and granting explicit permission in the prompt.
It would be exciting to see more ambitious exploration, even building with the agent SDK to complete the tasks, but none of that materialized.
Tasks arrived one at a time and had to be finished before the next was dispensed.
Under that regime a subagent buys little, and there may be little to gain from saving traces for cross-task analysis, because there is no second task in flight to analyze against.
I am interested to see whether these behaviors emerge with Fable 5 at maximum reasoning and a task server that dispenses concurrently.
Potential affordances not exposed by Claude Code
Skills can be authored, but loading one requires a session reload that the agent cannot trigger itself.
Writing one would have had no effect on the run that wrote it.
Its own tools are fixed: the MCP config is mounted read-only under --strict-mcp-config , so Code cannot register a new one on its own server.
The model is set at initialization to claude-opus-5 and Code cannot access the /model skill.
The subagent tool does take a model argument, so Code could have chosen a stronger or cheaper model for delegated work.
There is no fine-tuning API for Anthropic models, so weight updates are also closed.
The reasoning effort is also set at initialization to xhigh for the whole run and Code cannot autonomously call the /effort skill.
Unlike the model, this one is closed even by the subagent route: the subagent tool has no effort parameter, so a subagent inherits the session's setting.
The agent could not choose to think harder on a hard task, or cheaper on an easy one.
Context compaction fired four times over the run, automatically, when the context filled.
Code did not schedule it, could not defer it, and could not choose what survived.
Given that the compaction summary contributed to the improvement, an agent that can decide when to compact and what to keep is an experiment worth running.
Pi , Hermes , and Prime Agent are open source agent harnesses, so the affordances Claude Code does not expose could be opened up and measured rather than listed as unavailable.
Production feedback is not dense
Every submission was graded and the score came straight back: a numeric reward and a success flag, on all training tasks.
That is the condition the whole result rests on.
Real deployments rarely look like this.
An agent that updates a CRM record does not learn it got the account wrong until someone downstream notices, or a quarterly report comes out wrong, and often never receives feedback at all.
The open question is what improvement remains as the feedback signal thins.
Every number here was measured under immediate grading of every attempt.
A run scored on one task in ten, or scored a week later, may improve less, or not at all, and the part of the context that carries it need not be the same part.
The memories were not the mechani

[truncated]
