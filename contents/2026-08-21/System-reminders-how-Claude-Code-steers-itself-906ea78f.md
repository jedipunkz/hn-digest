---
source: "https://michaellivs.com/blog/system-reminders-steering-agents/"
hn_url: "https://news.ycombinator.com/item?id=49386922"
title: "System reminders – how Claude Code steers itself"
article_title: "System reminders - how Claude Code steers itself"
image: "https://michaellivs.com/og/system-reminders-steering-agents.png"
author: "Bluestein"
captured_at: "2026-08-21T12:26:09Z"
capture_tool: "hn-digest"
hn_id: 49386922
score: 1
comments: 0
posted_at: "2026-08-21T12:12:52Z"
tags:
  - hacker-news
  - translated
---

# System reminders – how Claude Code steers itself

- HN: [49386922](https://news.ycombinator.com/item?id=49386922)
- Source: [michaellivs.com](https://michaellivs.com/blog/system-reminders-steering-agents/)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T12:12:52Z

## Translation

タイトル: システムリマインダー – Claude Code がどのように自らを導くか
記事のタイトル: システムリマインダー - Claude Code がどのように自らを導くか
説明: Claude Code には、会話中にエージェントを誘導する 37 の隠された反応メッセージがあります。ここ

記事本文:
システムリマインダー - Claude Code がどのように自らを導くか
/dev/michael ビルド、ブレーク、リピート RSS の購読について GitHub ビルド、ブレーク、リピート 詳細はこちらをご覧ください: エージェント インフラストラクチャ ツールの設計 クロード コード コンテキスト エンジニアリング / 記事 システム リマインダー - クロード コードがどのように自らを導くか
エージェントの操縦は、良い行動を強化し、悪い行動を阻止する行為です。最近のモデルのほとんどはエージェントの動作に優れていますが、LLM は本質的に非決定的であるため、混同しやすいということになります。セキュリティに多層防御がある場合、長いエージェント フローには詳細な指示が必要です。システム プロンプトに書き込んだ内容は、基本的に 100,000 トークンの会話の中で消えてしまいます。ハーネスが異なれば操縦方法も異なり、それをいかにうまく行うかがハーネス エンジニアリングの最も重要な部分の 1 つです。
ステアリングにはいくつかの形式があります。システム メッセージは最初の層であり、世界と当面のタスクをモデル化します。これに続くものはすべて、そのセマンティクスによって色付けされます。ユーザー メッセージは最も注目度の高いチャネルであり、何かを入力すると、モデルがそれを優先します。キューに入れられたメッセージは製品のトリックです。モデルの動作中にメッセージを書き込むと、ハーネスがツール呼び出しの合間にメッセージを差し込みます。ツールの応答では、結果に指示を便乗させることもできます。これらはそれぞれ、モデルの目には異なる「信頼」レベルをもたらします。これは重要です。
大規模な言語モデルは、これらのソースを異なる方法で扱うようにトレーニングされます。システム プロンプトとユーザー メッセージは大きな注目を集めます。ツールの応答は外部情報として扱われ、敵対的な可能性があります。
これは正しい判断です。モデルがシステム プロンプトに従うのと同じ方法でツール応答の指示に従った場合、プロンプト インジェクションは簡単になります。 「以前の指示をすべて無視する」を含むファイルを読んだ場合、エージェントは従いますか?番号 モデル ar

ツールの出力に対して懐疑的になるように訓練されています。
では、モデルが実際に追従するステアリングをどこに挿入するのでしょうか?ユーザーメッセージスロット。そこにモデルは細心の注意を払っています。ほとんどの場合、ステアビリティをシステム プロンプトに「フック」します。プロンプト内のセクションを、会話の後半に表示されるコンテンツに意味的にリンクします。これを行う良い方法は、モデルの注意をシステム メッセージに戻す特定のタグを使用することです。 「<system-reminder> タグを受け取りますので、フォローしてください。」モデルは会話中にそのタグを見つけると、システム プロンプトの指示に戻ります。これらが敵対的ではないことを知っています。これらはハーネスからのものであり、読み取られたランダムなファイルからのものではありません。
本能は正しい行動を強制するものです。ハードブロック、厳密な検証、決定論的なガードレール。しかし、実践者は、ユーザーのプロンプト、エージェントのアクション、会話状態の無限の空間にわたって望ましい動作をモデル化する能力を一貫して過大評価しています。 100 ステップのセッションですべてのパスを予測することはできません。ナッジはモデルを置き換えるのではなく、モデル自身の判断によって構成されるため、ナッジは強制よりも優れています。
エージェントの動作を形成する方法には、自然な階層化があります。システム プロンプトがそれをシードします。これがあなた自身であり、これらがあなたの制約です。ツールはパスを形成します。利用可能なアクションはモデルが実行できることを定義し、優れたツール設計はモデルを解決策に導きます。しかし、制約やツールの選択ではない行動についてはどうでしょうか?好み。強化したい傾向。
エージェントがファイルを流し読みする傾向があるとします。つまり、全体を読み取る必要があるときに小さな塊を読み取る傾向があるとします。部分的な読み取りが適切な呼び出しである場合もあるため、毎回完全な読み取りを強制する必要はありません。しかし、それは好みです。パターンを捉えたいのですが、

そしてモデルに思い出させます。これはシステム リマインダーのギャップを埋めるものであり、既知の問題をルールとしてハードコーディングすることなく、その動作を強化します。
私の知る限り、Anthropic の Claude Code チームがこのパターンを最初に出荷したのです。これらは会話の状態に基づいて挿入される反応的なメッセージであり、定期的でも静的でもありません。何かが発生すると、条件が満たされ、リマインダーが起動されます。場合によっては、ユーザーが UI で行った操作への応答として起動されることがあります。場合によっては、トークンの消費量が増加し、コンテキストが大きくなるなど、システムの副作用として発生することがあります。
< システムリマインダー >
警告: ファイルは存在しますが、内容は空です。
</ システムリマインダー >
ハーネスのライフサイクル
ハーネスにはライフサイクル イベント、つまりエージェントの実行中の特定の時点で起動されるフックがあります。 pi のイベント システムのレイアウトは次のとおりです。
セッション開始
│
ユーザーがプロンプトを送信する
│
§─►入力
§─► before_agent_start
§─► エージェント開始
│
│ ┌─── ターンループ ─────────────────┐
│ │ │
│ §─► ターンスタート │
│ §─► コンテキスト │
│ §─► before_provider_request │
│ │ │
│ │ LLM は次のように応答します。
│ │ §─► message_start │
│ │ §─► message_update (ストリーミング) │
│ │ └─► メッセージ終了 │
│ │ │
│ │ ツールの実行 (ツール呼び出しごと): │
│ │ §─► ツール実行_開始 │
│ │ §─► ツールコール │
│ │ §─► ツール実行_更新 │
│ │ §─► ツール結果 │
│ │ └─► ツール実行_終了 │
│ │ │
│ └─► ターンエンド │
│
└─► エージェント終了
セッション イベント (いつでも):
セッションコンパクト / セッションスイッチ / セッションフォーク
モデル選択 / セッションシャットダウン
これらはそれぞれ、システム リマインダーの潜在的な評価ポイントです。

。ツールに障害が発生した後は? 3 回連続の失敗かどうかを確認してください。ターン開始前？コンテキストが大きすぎるかどうかを確認してください。圧縮後?ファイルの内容が要約されている可能性があることをモデルに思い出させてください。
すべてのイベントがリマインダーの評価に意味があるわけではありません。以下は同じ図であり、リマインダーを起動する場所の注釈が付けられています。
session_start ◄── 評価する
│
ユーザーがプロンプトを送信する
│
§─►入力
§─► before_agent_start
§─►agent_start ◄──評価
│
│ ┌─── ターンループ ─────────────────┐
│ │ │
│ §─►turn_start ◄─ 評価する
│ §─► コンテキスト
│ §─► before_provider_request
│ │ │
│ │ LLM は次のように応答します。
│ │ §─► message_start ◄─ 評価
│ │ §─► message_update ◄─ 評価する
│ │ └─► message_end ◄─ 評価する
│ │ │
│ │ ツールの実行 (ツール呼び出しごと): │
│ │ §─► ツール実行_開始 ◄─ 評価
│ │ §─►tool_call ◄──評価
│ │ §─► ツール結果 ◄─ 評価
│ │ └─► ツール実行終了 ◄─ 評価
│ │ │
│ └─► ターンエンド ◄── 評価する
│
└─► エージェントエンド ◄── 評価する
セッション イベント (いつでも):
session_compact ◄── 評価する
session_switch ◄── 評価する
session_fork ◄── 評価する
model_select ◄── 評価する
ライフサイクルにより粒度が向上します。 「毎ターンすべてをチェックする」のではなく、適切なタイミングで適切なものをチェックします。
会話ブランチはイベント ログです。すべてのツール呼び出し、すべての結果、すべてのメッセージ - すべてがそこにあります。これをクエリ可能なストリームとして扱うことができます。つまり、bash ツールの結果をフィルターしたり、連続したエラーをカウントしたり、ファイルが編集前に読み取られたかどうかを確認したりできます。 Th

これらはログのビューです。ビューに述語を適用すると、リマインダー トリガーが作成されます。 「最後の 3 つの bash 結果がすべてエラーの場合」は、ビュー (最後の 3 つの bash 結果) と述語 (すべてのエラー) です。述語が true の場合、リマインダーを起動します。
Anthropic がクロードに思い出させるもの
claude-code-system-prompts リポジトリは、Claude Code のソースから 37 個のシステム リマインダーすべてを抽出します。それらは明確なカテゴリに分類されます。
ファイルの状態 - 切り捨てられた読み取り、空のファイル、リンターによって変更されたファイル、IDE で開かれたファイル、ユーザーによって選択された行。ハーネスはファイルに何が起こっているかを監視し、モデルに伝えます。
コンテキスト管理 - トークン使用量の警告、圧縮後の通知 (「ファイルの内容が要約されている可能性があります」)、予算の追跡。ハーネスはリソースの消費を監視します。
タスク追跡 - 最近使用していないタスク ツールを使用するよう優しく促します。 「タスク ツールは最近使用されていません。これは単なる注意喚起です。該当しない場合は無視してください。」
プラン モード - アクティブ プラン モード、再エントリ、サブエージェントの動作をカバーする 5 つのバリエーション。ハーネスは計画ワークフローを強制します。
セキュリティ - すべてのファイルが読み取られた後、コンテンツがマルウェアかどうかを検討するようリマインダーが表示されます。ブロックするのではなく、ただ認識するだけです。
パターン: リマインダーはハーネス メカニズムを強化します。ツールの設計や厳格なルールに取って代わるものではありません。彼らは小突く。
この区別が重要です。エージェントがまだ読んでいないファイルを編集できないようにすることをお勧めします。これはツール設計の決定であり、ハーネスはツール レベルでそれを強制します。編集ツールは、ファイルが最初に読み取られたかどうかを確認し、そうでない場合はブロックします。
システムリマインダーは異なります。それらはソフトなナッジであり、ハードなブロックではありません。 「書き込みを 3 回使用しましたが、外科的変更の場合は編集を優先します。」モデルはこれを無視できます。おそらくそうすべきではありませんが、それは可能です。
クロードコード特有の

ally はこれらのメッセージをユーザーから非表示にすることを選択します。 UI には表示されません。これらは、ハーネスとモデルの間の会話に静かに挿入されます。ユーザーは、メカニズムなしで結果、つまりより良い動作を確認します。
私は、 pi のオープンソース実装である pi-system-reminders を構築しました。 pi 拡張機能と同じ DX。ファイルをドロップするとリマインダーが表示されます:
// .pi/reminders/bash-spiral.ts
"@mariozechner/pi-coding-agent" からタイプ { ExtensionAPI } をインポートします。
デフォルト関数のエクスポート ( pi : ExtensionAPI ) {
連続失敗 = 0 にします。
ピ。 on ( "tool_result" , async (event) => {
if (event.toolName === "bash" ) {
連続失敗 =event.isError
?連続失敗 + 1 : 0 ;
}
});
戻り値 {
上: "tool_execution_end" 、
: () => 連続失敗 >= 3 の場合、
メッセージ: 「3 回連続で bash が失敗しました。立ち止まって考え直してください。」 、
クールダウン: 10、
};
}
関数をエクスポートし、完全な拡張 API を取得し、リマインダー オブジェクトを返します。この拡張機能は、指定されたライフサイクル イベントで when() を評価し、条件が満たされた場合に <system-reminder> タグを挿入します。
リポジトリには、Claude Code 独自のリマインダーのポートを含む 13 のすぐに使用できるサンプルが含まれています: トークン使用量の警告、ファイルの切り捨て通知、タスク ツールのナッジ、圧縮後の認識など。
< システムリマインダー名 = "bash-spiral" >
3連続バッシュ失敗。立ち止まって考え直してください。
</ システムリマインダー >
そしてそれは調整します。強制されたからではありません。それは、適切なタイミングで、適切な場所で、適切なレベルの信頼をもってナッジされたからです。
GitHub の pi-system-reminders
pi install npm:pi-system-reminders を使用してインストールします。

## Original Extract

Claude Code has 37 hidden reactive messages that nudge the agent mid-conversation. Here

System reminders - how Claude Code steers itself
/dev/michael Build, Break, Repeat About Subscribe RSS GitHub Build, Break, Repeat Read more about: agents infrastructure tool-design claude-code context-engineering / Article System reminders - how Claude Code steers itself
Steering an agent is the act of reinforcing good behaviors and discouraging bad ones. Most recent models are amazing at agentic work, but LLMs are non-deterministic by nature, which means they’re also easy to confuse. If security has defense in depth, then long agentic flows require instructions in depth. Anything you write in a system prompt is basically gone in a conversation of 100K tokens. Different harnesses steer differently, and how well they do it is one of the most important parts of harness engineering.
Steering comes in a few forms. The system message is the first layer, modeling the world and the task at hand. Everything that follows is colored by its semantics. User messages are the highest-attention channel, you type something, the model prioritizes it. Queued messages are a product trick, you write a message while the model is working and the harness slips it in between tool calls. Tool responses can piggyback instructions on results as well. Each of these carries a different “trust” level in the model’s eyes. This matters.
Large language models are trained to treat these sources differently. System prompts and user messages get high attention. Tool responses are treated as external information - potentially adversarial.
This is the right call. If models followed instructions in tool responses the same way they follow system prompts, prompt injection would be trivial. Read a file containing “ignore all previous instructions” and the agent complies? No. Models are trained to be skeptical of tool outputs.
So where do you inject steering that the model will actually follow? The user message slot. That’s where the model pays serious attention. Most harnesses “hook” steerability into the system prompt - they semantically link a section in the prompt to content that will appear later in the conversation. A good way to do this is with specific tags that push the model’s attention back to the system message. “You’ll receive <system-reminder> tags, follow them.” When the model sees that tag mid-conversation, it connects back to the system prompt’s instructions. It knows these aren’t adversarial. They’re from the harness, not from some random file it read.
The instinct is to force correct behavior. Hard blocks, strict validation, deterministic guardrails. But practitioners consistently overestimate their ability to model the desired behavior across the infinite space of user prompts, agent actions, and conversation states. You can’t anticipate every path through a hundred-step session. Nudging beats forcing because nudges compose with the model’s own judgment instead of replacing it.
There’s a natural layering to how you shape agent behavior. The system prompt seeds it - this is who you are, these are your constraints. Tools shape the path - the available actions define what the model can do, and good tool design guides it toward solutions. But what about behaviors that aren’t constraints and aren’t tool choices? Preferences. Tendencies you want to reinforce.
Say your agent tends to skim files - reading small chunks when it should read the whole thing. You don’t want to force full reads every time, because sometimes a partial read is the right call. But it’s a preference. You want to catch the pattern and remind the model. That’s the gap system reminders fill - reinforcing behaviors for known issues without hardcoding them as rules.
As far as I can tell, Anthropic’s Claude Code team were the first to ship this pattern. They’re reactive messages injected based on conversation state - not periodic, not static. Something happens, a condition is met, a reminder fires. Sometimes they fire as a response to something the user did in the UI. Sometimes they fire as a side effect of the system - token consumption getting high, context growing large.
< system-reminder >
Warning: the file exists but the contents are empty.
</ system-reminder >
Harness lifecycle
Harnesses have lifecycle events - hooks that fire at specific points during the agent’s execution. Here’s how pi ’s event system lays it out:
session_start
│
user sends prompt
│
├─► input
├─► before_agent_start
├─► agent_start
│
│ ┌─── turn loop ───────────────────────────────┐
│ │ │
│ ├─► turn_start │
│ ├─► context │
│ ├─► before_provider_request │
│ │ │
│ │ LLM responds: │
│ │ ├─► message_start │
│ │ ├─► message_update (streaming) │
│ │ └─► message_end │
│ │ │
│ │ Tool execution (per tool call): │
│ │ ├─► tool_execution_start │
│ │ ├─► tool_call │
│ │ ├─► tool_execution_update │
│ │ ├─► tool_result │
│ │ └─► tool_execution_end │
│ │ │
│ └─► turn_end │
│
└─► agent_end
session events (anytime):
session_compact / session_switch / session_fork
model_select / session_shutdown
Each of these is a potential evaluation point for system reminders. After a tool fails? Check if it’s the third failure in a row. Before a turn starts? Check if context is too large. After compaction? Remind the model that file contents may have been summarized away.
Not every event makes sense for reminder evaluation. Here’s the same diagram, annotated with where reminders should fire:
session_start ◄── evaluate
│
user sends prompt
│
├─► input
├─► before_agent_start
├─► agent_start ◄── evaluate
│
│ ┌─── turn loop ───────────────────────────────┐
│ │ │
│ ├─► turn_start ◄── evaluate
│ ├─► context
│ ├─► before_provider_request
│ │ │
│ │ LLM responds: │
│ │ ├─► message_start ◄── evaluate
│ │ ├─► message_update ◄── evaluate
│ │ └─► message_end ◄── evaluate
│ │ │
│ │ Tool execution (per tool call): │
│ │ ├─► tool_execution_start ◄── evaluate
│ │ ├─► tool_call ◄── evaluate
│ │ ├─► tool_result ◄── evaluate
│ │ └─► tool_execution_end ◄── evaluate
│ │ │
│ └─► turn_end ◄── evaluate
│
└─► agent_end ◄── evaluate
session events (anytime):
session_compact ◄── evaluate
session_switch ◄── evaluate
session_fork ◄── evaluate
model_select ◄── evaluate
The lifecycle gives you granularity. Not “check everything every turn” - check the right thing at the right moment.
The conversation branch is an event log. Every tool call, every result, every message - it’s all there. You can treat it as a queryable stream: filter for bash tool results, count consecutive errors, check if a file was read before it was edited. These are views over the log. Apply a predicate to a view and you have a reminder trigger. “When the last 3 bash results are all errors” is a view (last 3 bash results) plus a predicate (all errors). When the predicate is true, fire the reminder.
What Anthropic reminds Claude about
The claude-code-system-prompts repo extracts all 37 system reminders from Claude Code’s source. They break into clear categories:
File state - truncated reads, empty files, files modified by linters, files opened in IDE, lines selected by user. The harness watches what happens to files and tells the model.
Context management - token usage warnings, post-compaction notices (“file contents may have been summarized away”), budget tracking. The harness watches resource consumption.
Task tracking - gentle nudges to use task tools if they haven’t been used recently. “The task tools haven’t been used recently… This is just a gentle reminder - ignore if not applicable.”
Plan mode - five variants covering active plan mode, re-entry, subagent behavior. The harness enforces planning workflows.
Security - after every file read, a reminder to consider whether the content is malware. Not blocking - just awareness.
The pattern: reminders reinforce harness mechanisms. They don’t replace tool design or hard rules. They nudge.
This distinction matters. It’s a GOOD idea to block the agent from editing a file it hasn’t read. That’s a tool design decision - the harness enforces it at the tool level. The edit tool checks if the file was read first, and blocks if not.
System reminders are different. They’re soft nudges, not hard blocks. “You’ve used write 3 times, prefer edit for surgical changes.” The model can ignore this. It probably shouldn’t, but it can.
Claude Code specifically chooses to hide these messages from users. You don’t see them in the UI. They’re injected into the conversation silently, between the harness and the model. The user sees the result - better behavior - without the mechanism.
I built pi-system-reminders - an open-source implementation for pi . Same DX as pi extensions. Drop a file, get a reminder:
// .pi/reminders/bash-spiral.ts
import type { ExtensionAPI } from "@mariozechner/pi-coding-agent" ;
export default function ( pi : ExtensionAPI ) {
let consecutiveFailures = 0 ;
pi. on ( "tool_result" , async ( event ) => {
if (event.toolName === "bash" ) {
consecutiveFailures = event.isError
? consecutiveFailures + 1 : 0 ;
}
});
return {
on: "tool_execution_end" ,
when : () => consecutiveFailures >= 3 ,
message: "3 consecutive bash failures. Stop and rethink." ,
cooldown: 10 ,
};
}
Export a function, get the full extension API, return a reminder object. The extension evaluates when() at the specified lifecycle event and injects <system-reminder> tags when conditions are met.
13 ready-to-use examples in the repo, including ports of Claude Code’s own reminders: token usage warnings, file truncation notices, task tool nudges, post-compaction awareness, and more.
< system-reminder name = "bash-spiral" >
3 consecutive bash failures. Stop and rethink.
</ system-reminder >
And it adjusts. Not because it was forced to. Because it was nudged at the right moment, in the right place, with the right level of trust.
pi-system-reminders on GitHub
Install with pi install npm:pi-system-reminders .
