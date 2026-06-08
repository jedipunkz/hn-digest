---
source: "https://viktor.com/blog/how-we-built-viktor-around-prompt-caching"
hn_url: "https://news.ycombinator.com/item?id=48448276"
title: "How to cut the cost of long AI agent threads (without making the agent dumber)"
article_title: "How We Built Viktor Around Prompt Caching (80% Cheaper Agent Threads) | Viktor Blog"
author: "peteralbert"
captured_at: "2026-06-08T18:58:05Z"
capture_tool: "hn-digest"
hn_id: 48448276
score: 1
comments: 0
posted_at: "2026-06-08T17:26:25Z"
tags:
  - hacker-news
  - translated
---

# How to cut the cost of long AI agent threads (without making the agent dumber)

- HN: [48448276](https://news.ycombinator.com/item?id=48448276)
- Source: [viktor.com](https://viktor.com/blog/how-we-built-viktor-around-prompt-caching)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T17:26:25Z

## Translation

タイトル: 長い AI エージェント スレッドのコストを削減する方法 (エージェントを愚かにすることなく)
記事のタイトル: プロンプト キャッシュを中心に Viktor を構築した方法 (エージェント スレッドが 80% 安価) |ヴィクトルのブログ
説明: Viktor のスレッド エンジンは、プロンプト キャッシュを中心に設計されています。ツール読み込みの代わりに SDK、追加専用スレッド、キャッシュ内圧縮、80% のコスト削減。

記事本文:
プロンプト キャッシュを中心に Viktor を構築した方法 (エージェント スレッドが 80% 安価) | Viktor ブログ 製品価格 セキュリティ エンタープライズ リソース ブログ ケーススタディ 変更ログ 比較 vs ChatGPT vs OpenClaw vs Claude in Slack vs Tasklet ソリューション 統合のユースケース 無料で始めましょう ブログに戻る AI 用のコピー 2026 年 6 月 8 日 · Toni Albert プロンプト キャッシュを中心に Viktor を構築した方法
Viktor のスレッド エンジンは、プロンプト キャッシュを中心に設計されています。ツール読み込みの代わりに SDK、追加専用スレッド、キャッシュ内圧縮、80% のコスト削減などです。
モデル API は（ほとんど）ステートレスであるため、エージェントは通話ごとに履歴全体を再送信します。 40 ステップの Viktor スレッドは、トランスクリプトの長さが最大 ​​85,000 トークンしかないにもかかわらず、最大 217,000 個の入力トークンを送信します。
プロンプト キャッシュにより、再送信されたトークンが 0.1x のキャッシュ読み取りに変換されます。 Claude Opus 4.8 では、サンプル スレッドが 11.35 ドルから 2.07 ドルに下がり、81.8% 削減されます。
キャッシュはプレフィックスが byte-stable の場合にのみ機能するため、エージェント アーキテクチャ全体を形成する必要があります。ツールはプロンプト内のスキーマではなくコード内の SDK 関数として公開され、すべてのスレッドは追加専用のログです。
要約はスレッド自体のキャッシュ内で実行され、別の呼び出しで全額を支払うのではなく、完全な履歴を 0.1x 読み取りとして送信します。
圧縮のタイミングはキャッシュのライフサイクルに従います。ホットスレッドを決して圧縮せず、キャッシュがコールドになる前の数分間に積極的に圧縮します。
プロバイダーごとにキャッシュの動作が異なるため (明示的なブレークポイント、自動、TTL、ルーティング)、スレッド エンジンはプロバイダーごとに適応します。
Viktor は、Slack と Microsoft Teams に住む AI 従業員です。サポート受信トレイの優先順位付け、CRM パイプラインの監査、QA 画面記録の分析、レポートの作成など、実際の作業は人々に任せられます。単一のタスクは通常、数十のモデル呼び出しを含むスレッドを意味し、それぞれがシステム プロンプト t を実行します。

ユーザーのスキルと記憶、これまでの会話、そして増え続けるツールの結果の山。
このワークロードの形状を単純に実装すると、恐ろしいコスト プロファイルになります。この記事では、フロンティア モデル エージェントを経済的に存続させるための問題、計算、および Viktor のスレッド エンジン内の特定のアーキテクチャ上の決定について説明します。以下のすべては実稼働コードベースからのものであり、Claude Opus 4.8 の価格で、すべての計算を通じて 1 つの具体的なサンプル スレッドを実行し続けます。
1. 問題: LLM API にはメモリがない
モデルとのチャットについて多くの人が抱くメンタル モデルは電話です。つまり、自分の新しい発言だけを送信するオープン ラインです。現実は、フォローアップの質問があるたびに、ケースファイル全体を新しいコンサルタントに郵送することに近いです。
モデル API はステートレスです (ほとんどの場合、ステートフル オプションが存在しますが、モデルが認識するものを完全に制御したい場合は、それらをステートレスとして扱います)。プロバイダー側​​には会話を記憶するセッションはありません。すべての呼び出しには、システム プロンプト、ツール定義、以前のすべてのユーザー メッセージ、すべてのアシスタントの応答、すべてのツール呼び出し、すべてのツールの結果など、モデルに必要なものすべてが含まれている必要があります。モデルが応答すると、その応答をローカルのトランスクリプトに追加し、次の呼び出しですべての応答と新しいターンが再送信されます。
ターンごとに履歴全体が再送信されます。新しいトークンはほんの一部です。再送信されたトークンが優先されます。
5 つの短いターンを伴う人間のチャットの場合、これは無関係です。エージェントにとって、エージェント ループはマシン速度でのエージェント自身との単なる会話であるため、これはすべての試合です。つまり、モデルを呼び出し、ツール呼び出しを取得し、実行し、結果を追加し、モデルを再度呼び出します。 40 ステップとは、増え続けるトランスクリプトを 40 回完全に再送信することを意味します。
このグロのコスト

ws は二次関数的に 。コンテキストが P トークンで始まり、各ステップで s トークンが追加される場合、N 回の呼び出しにわたる入力トークンの合計はおよそ N・P + s・N²/2 になります。タスクの長さが 2 倍になると、テール料金は 4 倍になります。
私たちの実行例。現実的な Viktor スレッド: 25,000 トークンの安定したプレフィックス (システム プロンプト、スキル、ツール定義)、40 のモデル呼び出し、およびツール呼び出しと結果の約 1,500 トークンを追加する各ステップ。スレッド全体で送信される合計入力: 2,170,000 トークン (ただし、最終的なトランスクリプトの長さはわずか約 85,000 トークンです)。同じ初期トークンを最大 40 回送信します。
2. 迅速なキャッシュによりユニットエコノミクスが変わります
プロバイダーは、事実上すべてのエージェント トラフィックが次のようになっていることに気づきました。つまり、バイトが同一の長いプレフィックスと小さな新しいサフィックスです。そこで彼らはプロンプト キャッシュを構築しました。プロバイダーは、プロンプト プレフィックスの処理された内部状態 (KV キャッシュ) を短期間保持し、次のリクエストがまったく同じバイトで始まる場合、再計算するのではなく、キャッシュされた状態から再開します。
割引率は劇的です。以下は、ほとんどの Viktor スレッドで実行されているモデルである Claude Opus 4.8 の価格です。
プレフィックスをキャッシュに書き込むために 25% のプレミアムを 1 回支払うと、その後はその読み取りごとに通常料金の 10 分の 1 の料金がかかります。コール N+1 がコール N からのすべてを再送信するエージェント ループでは、ほぼすべての入力トークンがキャッシュ読み取りになります。
両方の価格設定モードでサンプル スレッドを実行します。
同じスレッドで、Opus 4.8 では双方向の価格が設定されています。キャッシングにより 2 次再読み取りが 0.1x ラインアイテムに変換されるため、ギャップはステップごとに広がります。
過小評価されている 2 番目のメリットとして、レイテンシがあります。キャッシュされたトークンは事前入力計算をスキップするため、80K トークン コンテキストでの最初のトークンまでの時間は、数秒から新しいサフィックスを処理するコストとほぼ同程度まで減少します。エージェントが 40 シーケンシャを作成する場合

l コールを実行すると、タスクごとに数分の実時間を節約できます。
3. 落とし穴: キャッシュは壊れやすく、せっかちです
割引が 10 倍なら、なぜ誰もがそれをオンにして立ち去らないのでしょうか?プロンプト キャッシュには次の 2 つの明確な制約があるためです。
完全なプレフィックス一致。キャッシュは、新しいリクエストがキャッシュされたポイントまでバイト同一である場合にのみ再開されます。システム プロンプトで 1 文字を変更し、ツール定義を並べ替え、タイムスタンプを挿入すると、全額料金ですべてが再計算されます。
寿命が短い。 Anthropic の標準キャッシュ エントリの存続期間は約 5 分間で、ヒットするたびに更新されます。数分間静かにすると、キャッシュは消えます。次の呼び出しでは、プレフィックス全体が完全に書き換えられます。
これは、プロンプト キャッシュがチェックボックスではないことを意味します。これは、エージェント アーキテクチャ全体を形成する必要がある設計上の制約です。 Viktor のモデル層における興味深いエンジニアリングのほとんどは、1 つの質問に答えるために存在しています。それは、プレフィックスのバイト安定性をどのように維持するか、キャッシュが失われる数秒前に何をするかということです。
4. Viktor のスレッド エンジンはキャッシュを中心にどのように設計されているか
Viktor モデル呼び出しの構造。最新のメッセージの下にあるものはすべて、キャッシュされた安定したプレフィックスです。
4.1 決して変更されないツールセット: SDK の決定
Viktor は 3,000 以上のツールに接続します。ツールをモデルに公開する標準的な方法は、すべてのツールの JSON スキーマをリクエストに挿入することですが、この規模ではそれはまったく機能しません。スキーマだけではコンテキストの制限を超えてしまい、そこに到達するずっと前にツール呼び出しの精度が低下します。ツールセットが大きくなるにつれて、モデルは適切なツールを選択することが目に見えて悪くなります。したがって、常に関連するツールのみが表示される、何らかの形式の動的なツール検出が必要です。しかし、ここではキャッシュが反撃します。ツール定義はプロンプトの最前面、つまり会話の前に置かれます。ロア

d はリクエストに動的にツールを検出し、プレフィックスが位置 0 で変更され、ツールセットがシフトするたびにスレッドのキャッシュが無効になります。
したがって、動的ツールは安定したものの背後にラップする必要があります。 1 つのオプションは、自分で解析したペイロード辞書を受け入れる単一の汎用ディスパッチ ツールです。 Viktor はさらに一歩進んでいます。モデルは実際のコードを記述し、すべての統合はエージェントのサンドボックス内の SDK でプレーンな関数として公開されます。このモデルは、小規模で固定されたネイティブ ツールのセット (シェル コマンドの実行、ファイルの読み取りと編集、Slack メッセージの送信、およびその他のいくつかのツール) を認識します。 Stripe または Linear の呼び出しは、プロンプト内のツール スキーマではありません。モデルがスクリプトで記述する Python の 3 行です。 (ツール アクセスのスケーリングについては、「エージェントが 100,000 ツールを使用すると何が起こるか」で詳しく説明しました。)
Copy # モデルは「ストライプリスト_サブスクリプション」ツールスキーマを取得しません。
# bash ツールを取得し、代わりにこれを書き込みます:
sdk.tools. Stripe_tools から list_subscriptions をインポートします
#
subs = await list_subscriptions(status="active", limit=100) 4.2 スレッドは追加専用のログです
正確なプレフィックス規則を記述するより厳密な方法があり、設計原則に昇格する価値があります。つまり、キャッシュに優しいスレッドは追加専用のログです。送信された内容を編集、並べ替え、または削除することはできません。唯一の合法的な操作は末尾に追加することです。 Viktor のスレッド エンジンはこれを厳密な不変条件として扱います。これは思っているよりも制約が強いです。
システムプロンプトに時計がありません。現在のタイムスタンプを位置 0 に挿入すると、呼び出しごとにすべてのキャッシュが無効になります。代わりに、各ターンには追加時に独自の固定タイムスタンプが保持され、その後変更されることはなく、モデルは最新の時刻から現在時刻を導出します。
スレッド中の変更は、編集ではなくメッセージとして到着します。いつ

ユーザーが新しい統合に接続したり、スレッドの途中で命令を更新したりすると、ライブ スレッドのシステム プロンプトを書き換えることはできません。更新は新しいメッセージとして追加する必要があり、システム プロンプトはその後開始されるスレッドにのみ更新を反映します。
これらの制約を考慮して、0.1 倍のラインアイテムの価格を設定します。その結果、キャッシュの安全性が構造化されます。誤って履歴を変更する可能性のあるコード パスが存在しないため、ウォーム キャッシュを誤って焼き付ける可能性のあるコード パスも存在しません。
4.3 会話に応じたキャッシュ ブレークポイント
Anthropic のキャッシュは明示的です。cache_control ブレークポイントを使用してプロンプト内の位置をマークすると、マーカーまでのすべてが 1 つのプレフィックスとしてキャッシュされます。 Viktor のモデル レイヤーは、これらのマーカーを 3 種類の場所に配置します。
システム プロンプト ブロック上 (システム プロンプトは特定のユーザーに対して安定しているため、1 つのスレッド内だけでなくスレッド全体にキャッシュされます)、
最後のツール定義 (ツールがメッセージの前にあるため、ツール ブロック全体がキャッシュされます)、
会話の最後の 2 つのユーザー役割メッセージについて。
def add_cache_control(messages) をコピーします。
「」
最後から 2 番目と最後のユーザー メッセージに、cache_control を追加します。
こうすることで、以前のメッセージは変更されていないと想定できます。
したがって、キャッシングは機能します。
「」
user_messages = [i for i, msg in enumerate(messages) if msg["role"] == "user"]
変更するインデックス = ユーザーメッセージ[-2:]
#
indices_to_modify のインデックスの場合:
# ... そのメッセージ内の最後の text/tool_result ブロックを検索します ...
item["cache_control"] = {"type": "ephemeral"} 4.4 スレッド自身のキャッシュを再利用するコンパクション
キャッシュにより再送信は安くなりますが、コンテキストが無限になるわけではありません。また、200K トークンのトランスクリプトをモデルが推論するのに快適なものにはなりません。すべての本格的なエージェント ランタイムと同様に、Viktor は長いスレッドを圧縮します: 古いターン

高密度の要約に要約され、ライブ ウィンドウには要約と最近のメッセージのみが保持されます。
スレッドを要約する単純な方法は、別の呼び出しを開始することです。新しいプロンプト、「この会話を要約してください」、トランスクリプトを貼り付けます。この呼び出しはライブ スレッドとプレフィックスを共有しないため、慎重にキャッシュしていた履歴全体を再処理するには全額を支払うことになります。
Viktor は代わりに、スレッド自体のキャッシュ内で圧縮を実行します。要約リクエストは、バイトごとのライブ スレッドであり、圧縮命令を含む最終ユーザー メッセージが 1 つだけ追加されています。
def build_summary_request_messages(ai_messages, plan, *, use_full_history) をコピーします。
use_full_history の場合:
summary_messages = list(ai_messages) # ライブスレッド、変更なし
#
summary_messages.append({
"ロール": "ユーザー",
「コンテンツ」: [{
"タイプ": "テキスト",
"テキスト": COMPACTION_PROMPT,
SKIP_CACHE_CONTROL_KEY: True、 # キャッシュ ブレークポイントを移動しません
}]、
})
return summary_messages 圧縮プロンプトはキャッシュ ブレークポイントから除外されます。 Anthropic ではリクエストごとに 4 つのブレークポイントが許可されており、スレッドはすでに 4 つすべてを使用しています。追加された命令はスキップされるようにマークされているため、この命令にのみ存在するメッセージのブレークポイントが無駄にならないようになります。

[切り捨てられた]

## Original Extract

How Viktor's thread engine is designed around prompt caching: SDK instead of tool loading, append-only threads, in-cache compaction, and an 80% cost cut.

How We Built Viktor Around Prompt Caching (80% Cheaper Agent Threads) | Viktor Blog Product Pricing Security Enterprise Resources Blog Case Studies Changelog Compare vs ChatGPT vs OpenClaw vs Claude in Slack vs Tasklet Solutions Integrations Use Cases Get Started for Free Back to Blog Copy for AI June 8, 2026 · Toni Albert How We Built Viktor Around Prompt Caching
How Viktor's thread engine is designed around prompt caching: SDK instead of tool loading, append-only threads, in-cache compaction, and an 80% cost cut.
Model APIs are (mostly) stateless, so agents re-send their entire history on every call. A 40-step Viktor thread transmits ~2.17M input tokens even though the transcript is only ~85K tokens long.
Prompt caching turns re-sent tokens into 0.1x cache reads. On Claude Opus 4.8 our example thread drops from $11.35 to $2.07, an 81.8% reduction.
Caching only works if the prefix is byte-stable , so it has to shape the whole agent architecture: tools are exposed as SDK functions in code instead of schemas in the prompt, and every thread is an append-only log.
Summarization runs inside the thread's own cache , sending the full history as a 0.1x read instead of paying full price in a separate call.
Compaction timing follows the cache lifecycle: never compact a hot thread, compact aggressively in the minutes before the cache goes cold.
Every provider's cache behaves differently (explicit breakpoints vs automatic, TTLs, routing), so the thread engine adapts per provider.
Viktor is an AI employee that lives in Slack and Microsoft Teams. People hand it real work: triage a support inbox, audit a CRM pipeline, analyze a QA screen recording, build a report. A single task routinely means a thread with dozens of model calls , each one carrying the system prompt, the user's skills and memory, the conversation so far, and a growing pile of tool results.
That workload shape has a brutal cost profile if you implement it naively. This post walks through the problem, the math, and the specific architectural decisions inside Viktor's thread engine that keep frontier-model agents economically viable. Everything below comes from our production codebase, and we will keep one concrete example thread running through every calculation, priced on Claude Opus 4.8.
1. The problem: LLM APIs have no memory
The mental model many people have of a chat with a model is a phone call: an open line where you only transmit the new things you say. The reality is closer to mailing the entire case file to a new consultant every time you have a follow-up question.
Model APIs are stateless (mostly: stateful options exist, but if you want to retain full control over what the model sees, you treat them as stateless). There is no session on the provider's side that remembers your conversation. Every single call must contain everything the model needs: the system prompt, the tool definitions, every prior user message, every assistant reply, every tool call and every tool result. When the model answers, you append its reply to your local transcript, and the next call re-sends all of it again, plus the new turn.
Every turn re-sends the entire history. The new tokens are the small part; the re-sent tokens dominate.
For a human chat with five short turns, this is irrelevant. For an agent it is the whole ballgame, because an agent loop is just a conversation with itself at machine speed: call the model, get a tool call, execute it, append the result, call the model again. Forty steps means forty full re-transmissions of an ever-growing transcript.
The cost of this grows quadratically . If your context starts at P tokens and each step appends s tokens, the total input tokens across N calls is roughly N·P + s·N²/2 . Double the length of a task and you pay four times as much for the tail.
Our running example. A realistic Viktor thread: a 25,000-token stable prefix (system prompt, skills, tool definitions), 40 model calls, and each step appending ~1,500 tokens of tool calls and results. Total input transmitted across the thread: 2,170,000 tokens , even though the final transcript is only ~85,000 tokens long. You send the same early tokens up to 40 times.
2. Prompt caching changes the unit economics
Providers noticed that virtually all agent traffic looks like this: a long, byte-identical prefix plus a small new suffix. So they built prompt caching: the provider keeps the processed internal state (the KV cache) of your prompt prefix for a short time, and if your next request starts with the exact same bytes, it resumes from the cached state instead of recomputing it.
The discount is dramatic. Here is the pricing for Claude Opus 4.8, the model we run for most Viktor threads:
You pay a 25% premium once to write a prefix into the cache, and then every read of it costs a tenth of the normal price. In an agent loop where call N+1 re-sends everything from call N, nearly all input tokens become cache reads.
Run our example thread through both pricing modes:
The same thread, priced both ways on Opus 4.8. The gap widens with every step, because caching turns quadratic re-reading into a 0.1× line item.
There is a second, underrated benefit: latency . Cached tokens skip the prefill computation, so time-to-first-token on a 80K-token context drops from many seconds to roughly the cost of processing the new suffix. For an agent making 40 sequential calls, that compounds into minutes of wall-clock time saved per task.
3. The catch: caches are fragile and impatient
If the discount is 10×, why doesn't everyone just turn it on and walk away? Because prompt caches come with two sharp constraints:
Exact prefix match. The cache resumes only if the new request is byte-identical up to the cached point. Change one character in the system prompt, reorder a tool definition, inject a timestamp, and you recompute everything at full price.
Short lifetime. Anthropic's standard cache entry lives about 5 minutes and refreshes on every hit. Go quiet for a few minutes and the cache is gone; the next call pays a full re-write of the whole prefix.
This means prompt caching is not a checkbox. It is a design constraint that has to shape your entire agent architecture. Most of the interesting engineering in Viktor's model layer exists to answer one question: how do we keep the prefix byte-stable, and what do we do in the seconds before a cache dies?
4. How Viktor's thread engine is designed around the cache
The anatomy of a Viktor model call. Everything below the newest message is a stable, cached prefix.
4.1 A toolset that never changes: the SDK decision
Viktor connects to 3,000+ tools. The standard way to expose tools to a model is to inject every tool's JSON schema into the request, and at this scale that simply does not work. The schemas alone would blow past context limits, and tool-calling accuracy degrades long before you get there: models get measurably worse at picking the right tool as the toolset grows. So you need some form of dynamic tool discovery , where only the relevant tools are surfaced at any given moment. But here caching bites back: tool definitions sit at the very front of the prompt, before the conversation. Load discovered tools into the request dynamically and the prefix changes at position zero, invalidating the thread's cache every time the toolset shifts.
So dynamic tools have to be wrapped behind something stable. One option is a single generic dispatch tool that accepts a payload dict you parse out yourself. Viktor goes a step further: the model writes real code, and every integration is exposed as a plain function in an SDK inside the agent's sandbox . The model sees a small, fixed set of native tools (run a shell command, read and edit files, send a Slack message, and a handful of others). Calling Stripe or Linear is not a tool schema in the prompt; it is three lines of Python the model writes in a script. (We wrote more about scaling tool access in What Breaks When Your Agent Has 100,000 Tools .)
Copy # The model does not get a "stripe_list_subscriptions" tool schema.
# It gets a bash tool, and writes this instead:
from sdk.tools.stripe_tools import list_subscriptions
#
subs = await list_subscriptions(status="active", limit=100) 4.2 The thread is an append-only log
There is a stricter way to state the exact-prefix rule, and it is worth elevating to a design principle: a cache-friendly thread is an append-only log . Nothing that has been sent may ever be edited, reordered, or deleted; the only legal operation is appending to the end. Viktor's thread engine treats this as a hard invariant, and it is more constraining than it sounds:
No clock in the system prompt. Injecting the current timestamp at position zero would invalidate every cache on every call. Instead, each turn carries its own fixed timestamp when it is appended, which never changes afterwards, and the model derives the current time from the newest one.
Mid-thread changes arrive as messages, not edits. When a user connects a new integration or updates an instruction halfway through a thread, we cannot rewrite the system prompt of the live thread. The update has to be appended as a new message, and the system prompt only reflects it in threads that start afterwards.
Living with these constraints is the price of the 0.1× line item. The payoff is that cache safety becomes structural: there is no code path that can accidentally mutate history, so there is no code path that can accidentally torch a warm cache.
4.3 Cache breakpoints that ride the conversation
Anthropic's cache is explicit: you mark positions in the prompt with cache_control breakpoints, and everything up to a marker is cached as one prefix. Viktor's model layer places these markers in three kinds of places:
on the system prompt block (since the system prompt is stable for a given user, this caches it across threads, not just within one),
on the last tool definition (which caches the whole tool block, since tools precede messages),
on the last two user-role messages in the conversation.
Copy def add_cache_control(messages):
"""
Adds cache_control to the second to last and last user messages.
This way we can assume that the earlier messages didn't change
so caching will work.
"""
user_messages = [i for i, msg in enumerate(messages) if msg["role"] == "user"]
indices_to_modify = user_messages[-2:]
#
for index in indices_to_modify:
# ... find the last text/tool_result block in that message ...
item["cache_control"] = {"type": "ephemeral"} 4.4 Compaction that reuses the thread's own cache
Caching makes re-sending cheap, but it does not make context infinite, and it does not make a 200K-token transcript pleasant for the model to reason over. Like every serious agent runtime, Viktor compacts long threads: older turns get summarized into a dense recap, and the live window keeps only the summary plus recent messages.
The naive way to summarize a thread is to spin up a separate call: fresh prompt, "summarize this conversation", paste the transcript. That call shares no prefix with the live thread, so you pay full price to re-process the entire history you were so carefully caching.
Viktor instead runs compaction inside the thread's own cache . The summarization request is the live thread, byte-for-byte, with exactly one thing appended: a final user message containing the compaction instructions.
Copy def build_summary_request_messages(ai_messages, plan, *, use_full_history):
if use_full_history:
summary_messages = list(ai_messages) # the live thread, unchanged
#
summary_messages.append({
"role": "user",
"content": [{
"type": "text",
"text": COMPACTION_PROMPT,
SKIP_CACHE_CONTROL_KEY: True, # don't move the cache breakpoints
}],
})
return summary_messages The compaction prompt is excluded from cache breakpoints. Anthropic allows four breakpoints per request and the thread already spends all four. The appended instruction is marked to be skipped so it does not waste a breakpoint on a message that exists only in this one

[truncated]
