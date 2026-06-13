---
source: "https://getagentloop.io/"
hn_url: "https://news.ycombinator.com/item?id=48518018"
title: "Learning Infrastructure for AI Agents"
article_title: "AgentLoop — runtime learning for AI agents"
author: "martinembon"
captured_at: "2026-06-13T15:37:16Z"
capture_tool: "hn-digest"
hn_id: 48518018
score: 4
comments: 0
posted_at: "2026-06-13T15:04:52Z"
tags:
  - hacker-news
  - translated
---

# Learning Infrastructure for AI Agents

- HN: [48518018](https://news.ycombinator.com/item?id=48518018)
- Source: [getagentloop.io](https://getagentloop.io/)
- Score: 4
- Comments: 0
- Posted: 2026-06-13T15:04:52Z

## Translation

タイトル: AI エージェントの学習インフラストラクチャ
記事のタイトル: AgentLoop — AI エージェントのランタイム学習
説明: AgentLoop は、AI エージェントが繰り返される間違いを回避するのに役立つランタイム学習レイヤーです。人間による修正は再利用可能な記憶となり、将来の対応を改善します。

記事本文:
エージェントループ™
仕組み
なぜ
価格設定
お問い合わせ
ダッシュボードを開く
Runtime learning for production AI agents
AIエージェントは、
同じことを繰り返さないでください
二度間違える。
AgentLoop is a runtime learning layer for production AI agents.
人間による修正は再利用可能なメモリになります - すべての前に検索されます
応答が自動的に適用され、再トレーニングすることなくエージェントを改善します。
すべての応答は、人間、ヒューリスティック、またはダウンストリームによってレビューされます。
信号。 Corrections become memory. Memory gets pulled into the next prompt.
ウェイトに触れずに永遠にループする 3 つのステップ。
各 LLM 呼び出しの前に、AgentLoop は過去の修正を意味的に検索し、最も関連性の高い修正をプロンプトに挿入します。あなたのエージェントは、昨日の何が間違っていたのかを理解しました。
応答 (質問、回答、モデル、シグナル) は自動的に記録されます。接着コードはありません。失敗したケースは、信号強度によってランク付けされてレビュー キューに表示されます。
A reviewer writes the correct answer once.これは埋め込まれ、重複が排除され、将来のあらゆるクエリで即座に利用可能になります。同じ形状、同じ文言で永久に修正されます。
SDKのオーバーホールはありません。 No prompt-engineering rewrite. AgentLoop のラッパー
OpenAI と Anthropic インターフェイスをそのまま維持します。そして
メモリ取得を追加し、ロギングを舞台裏で実行します。
openaiインポートからOpenAI
from agentloop import AgentLoop
from agentloop_openai import wrap_openai
# 1. Wrap your existing OpenAI client.
client =wrap_openai (
OpenAI ()、
loop= AgentLoop (api_key= "ak_live_..." ),
)
#2. 通常どおり使用します。 AgentLoop runs around it.
response = client.chat.completions.作成(
モデル= "gpt-4o" 、
メッセージ=[
{ "role" : "system" , "content" : "あなたは役に立つアシスタントです。" }、
{ "role" : "user" , "content" : question},
]、
agentloop ={ "user_id" : user.id},
)
あ
検索は前に実行されます
The wrapper

AgentLoop の検索エンドポイントを呼び出し、関連する以前の修正を見つけて、システム プロンプトをサイレントに拡張します。
B
ロギングはその後実行されます
完了したターンはレビュー キューに投稿されます。 AgentLoop に到達できない場合、ラッパーはフェールオープンします。ユーザーは待つ必要はありません。
C
同じ API、ロックインなし
Wrap_openai() 呼び出しを削除しても、コードは引き続き機能します。 AgentLoop がユーザーとプロバイダーの間に介在することはありません。
03 / チームが使用する理由
Built for the messy middle, not the demo.
ほとんどのエージェントはステージ上で見栄えがよく、本番でも活躍します。難しいこと
この部分は最初の応答ではなく、1000 番目の応答です。
edge cases outnumber the happy path.
Stop hand-patching the system prompt
すべての出荷チームは最終的に、「常に覚えておいてください…」例外でいっぱいの 4,000 行のシステム プロンプトを維持することになります。 AgentLoop は、これを構造化メモリ (検索可能、重複排除、編集可能、監査可能) に置き換えます。
微調整を実行せずにフィードバック ループを閉じる
対象分野の専門家が修正を平易な言葉で一度書きます。これは、次のトレーニング サイクルではなく、将来のすべてのユーザー、将来のすべてのセッションに数秒で適用されます。ダッシュボードには、モデルが報告した内容ではなく、レビュー担当者が実際に修正した内容が表示されます。
Python and JavaScript, real parity
Python および JS SDK は、バイト同一の HMAC 署名を生成します。一方で署名されたフィードバック URL は、もう一方でも検証されます。どちらの言語も同じバックエンドにアクセスするため、各呼び出しがどの SDK から来たかに関係なく、動作は一貫しています。
Not locked to any one provider
OpenAI と Anthropic のドロップイン ラッパー、ファーストクラスの LangChain 統合、その他の直接 REST API。プロバイダーを切り替えても、蓄積された修正に費用がかかることはありません。メモリ層は、どのモデルを使用していても寿命が長くなります。
The demo answers wrong on purpose.
Ask the support agent a question. It'll give you a confidently
間違った答え

。修正してください。もう一度尋ねてください - 今度はそれを思い出します。
ループ全体は約 90 秒です。
無料で始めて、
準備ができたらアップグレードしてください。
無料プランには、AgentLoop の統合に必要なすべてが含まれています
そしてそれがスタック内で動作することを確認してください。
もっとヘッドルームが必要ですか?有料プランは使用量ベースで有効です。
予測可能であり、驚くような請求はありません。現在の計画を確認し、
アプリ内の制限。
AgentLoop がスタックに適合するかどうかについて質問がありますか?議論したい
生産価格？直接書きます。私たちはすべてのメールを読みます。
AI エージェントをレビューする人間から学習する、AI エージェントのランタイム学習レイヤー。

## Original Extract

AgentLoop is a runtime learning layer that helps AI agents avoid repeating mistakes. Human corrections become reusable memory that improves future responses.

AgentLoop ™
How it works
Why
Pricing
Contact
Open dashboard
Runtime learning for production AI agents
AI agents that
don't repeat the same
mistake twice.
AgentLoop is a runtime learning layer for production AI agents.
Human corrections become reusable memory — searched before every
response, applied automatically, improving the agent without retraining.
Every response gets reviewed — by a human, a heuristic, or a downstream
signal. Corrections become memory. Memory gets pulled into the next prompt.
Three steps, looping forever, without touching weights.
Before each LLM call, AgentLoop searches past corrections semantically and injects the most relevant ones into the prompt. Your agent now knows what it was wrong about yesterday.
The response is logged automatically — question, answer, model, signals. No glue code. Failed cases surface in a review queue, ranked by signal strength.
A reviewer writes the correct answer once. It's embedded, deduplicated, and instantly available to every future query — same shape, same wording, fixed for good.
No SDK overhaul. No prompt-engineering rewrite. AgentLoop's wrappers
keep the OpenAI and Anthropic interfaces exactly as they are — and
add memory retrieval and turn logging behind the scenes.
from openai import OpenAI
from agentloop import AgentLoop
from agentloop_openai import wrap_openai
# 1. Wrap your existing OpenAI client.
client = wrap_openai (
OpenAI (),
loop= AgentLoop (api_key= "ak_live_..." ),
)
# 2. Use it like normal. AgentLoop runs around it.
response = client.chat.completions. create (
model= "gpt-4o" ,
messages=[
{ "role" : "system" , "content" : "You are a helpful assistant." },
{ "role" : "user" , "content" : question},
],
agentloop ={ "user_id" : user.id},
)
A
Search runs before
The wrapper calls AgentLoop's search endpoint, finds relevant prior corrections, and silently augments the system prompt.
B
Logging runs after
The completed turn is posted to the review queue. If AgentLoop is unreachable, the wrapper fails open — your user never waits.
C
Same API, no lock-in
Remove the wrap_openai() call and the code still works. AgentLoop never sits between you and your provider.
03 / Why teams use it
Built for the messy middle, not the demo.
Most agents look great on stage and break in production. The hard
part isn't the first response — it's the thousandth, when the
edge cases outnumber the happy path.
Stop hand-patching the system prompt
Every shipping team eventually maintains a 4,000-line system prompt full of "always remember that…" exceptions. AgentLoop replaces that with structured memory — searchable, deduplicated, editable, audited.
Close the feedback loop without a fine-tuning run
Subject-matter experts write the fix once, in plain language. It applies to every future user, every future session, in seconds — not the next training cycle. The dashboard surfaces what reviewers actually fixed, not what models reported.
Python and JavaScript, real parity
The Python and JS SDKs produce byte-identical HMAC signatures. Feedback URLs signed in one validate in the other. Both languages hit the same backend, so behavior is consistent regardless of which SDK each call came from.
Not locked to any one provider
Drop-in wrappers for OpenAI and Anthropic, first-class LangChain integration, and a direct REST API for anything else. Switching providers doesn't cost you your accumulated corrections — the memory layer outlives whichever model you're on.
The demo answers wrong on purpose.
Ask the support agent a question. It'll give you a confidently
wrong answer. Correct it. Ask again — this time it remembers.
The whole loop, in about 90 seconds.
Start free,
upgrade when you're ready.
The free plan covers everything you need to integrate AgentLoop
and see it working in your stack.
Need more headroom? Paid plans are live — usage-based and
predictable, with no surprise bills. See current plans and
limits inside the app.
Questions about whether AgentLoop fits your stack? Want to discuss
production pricing? Write directly — we read every email.
A runtime learning layer for AI agents that learn from the humans who review them.
