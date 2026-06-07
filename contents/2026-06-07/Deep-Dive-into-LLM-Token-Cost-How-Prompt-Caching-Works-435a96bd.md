---
source: "https://weidongzhou.wordpress.com/2026/06/06/deep-dive-into-llm-token-cost-blog-series-part-2-how-prompt-caching-actually-works/"
hn_url: "https://news.ycombinator.com/item?id=48437673"
title: "Deep Dive into LLM Token Cost: How Prompt Caching Works"
article_title: "Deep Dive into LLM Token Cost — Blog Series Part 2: How Prompt Caching Actually Works | My Big Data World"
author: "tanelpoder"
captured_at: "2026-06-07T21:31:55Z"
capture_tool: "hn-digest"
hn_id: 48437673
score: 2
comments: 0
posted_at: "2026-06-07T19:09:46Z"
tags:
  - hacker-news
  - translated
---

# Deep Dive into LLM Token Cost: How Prompt Caching Works

- HN: [48437673](https://news.ycombinator.com/item?id=48437673)
- Source: [weidongzhou.wordpress.com](https://weidongzhou.wordpress.com/2026/06/06/deep-dive-into-llm-token-cost-blog-series-part-2-how-prompt-caching-actually-works/)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T19:09:46Z

## Translation

タイトル: LLM トークン コストの詳細: プロンプト キャッシュの仕組み
記事のタイトル: LLM トークン コストの詳細 — ブログ シリーズ パート 2: プロンプト キャッシュの実際の仕組み |私のビッグデータワールド
説明: このシリーズの最初の投稿、パート 1: 現実世界のケーススタディは、1 つの数字で終わりました。それは、172.58 ドルの費用がかかった 31 時間のクロード コード セッションで、そのうち 114.98 ドル (請求額の約 66%) がキャッシュ読み取りでした。キャッシュはそのセッションの副作用ではありませんでした。量的にはこれが支配的なコストラインでした。
[切り捨てられた]

記事本文:
LLM トークンのコストの詳細 — ブログ シリーズ パート 2: プロンプト キャッシュの実際の仕組み
このシリーズの最初の投稿であるパート 1: 現実世界のケーススタディは、1 つの数字で終わりました。31 時間のクロード コード セッションの費用は 172.58 ドルで、そのうち 114.98 ドル (請求額の約 66%) はキャッシュ読み取りでした。キャッシュはそのセッションの副作用ではありませんでした。これは、他のすべてのラインを合わせたものを上回るほどの量の支配的なコストラインでした。キャッシュがどのように機能するかについての正確なメンタル モデルを持たずに LLM コストについて推論しようとする人は、最も重要な部分を見逃していることになります。
この投稿はそのメンタルモデルです。これは戦略に関する記事ではありません。それは次の記事です。これは、特にキャッシュに関するメカニズムの投稿です。プレフィックス マッチが実際に回線上でどのように動作するか、会話の 2 番目のメッセージと最初のメッセージで何が起こるか、2 日間離れた場合に何が起こるか、ターン 5 でセッションに読み込まれた 1 つの 800 トークン ファイルに実際に何が起こるかなどです。
範囲に関するメモ。 Claude は、そのキャッシュ サーフェスが最も明示的であるため、全体を通して実際の例として使用されます。すべての読み取り、書き込み、TTL には使用ブロック内に対応するカウンターがあり、仕組みを 1 行ずつ追跡できます。 GPT と Gemini が現実的な答えを変えるような分岐点がある場合は、プロバイダー間の短いコールアウトが見つかります。このシリーズの 3 番目の投稿では、クロスプロバイダーの話をさらに詳しく説明します。ここでの目標は、まずクロードのメカニズムをしっかりとしたものにすることです。
パート 1 — キャッシュが実際にどのように機能するかを明らかにする 3 つの質問。すべてのチャット メッセージはキャッシュを使用しますか?フルコンテキストウィンドウは本当に毎ターン送信されるのでしょうか? 2 日後にセッションを再開するとどうなりますか?短い回答は、ほぼすべてのクロードのコスト決定を支配する非対称性を明らかにします。
パート 2 — 実際に動作した例

e.単一の 800 トークン ファイルを 20 ターンのセッションのライフサイクル全体にわたって追跡します。新しい入力として到着し、1 ターン遅れてキャッシュに書き込み、その後残りのセッションでキャッシュを読み取ります。 3 つの図によって仕組みが具体的にわかります。 800 個のトークンを追跡できるようになると、何でも追跡できるようになります。
シリーズ。これは 3 つの投稿のうちの 2 つ目です。
パート 1: 現実世界のケーススタディ (前) — 172 ドルの実際のケーススタディに基づいたメンタル モデル。
パート 2: プロンプト キャッシュが実際にどのように機能するか (この投稿) — 3 つの質問と 1 つの実際の例。
パート 3: 戦略とアンチパターン (次回) — 影響、それを元に戻すサイレント障害、およびプロバイダー間の比較によってランク付けされた 5 つの戦略。
各ポストは独立して立っています。しかし、最初の記事を読まずにここにたどり着いた方にとって、172 ドルのケーススタディは、なぜこの投稿が重要なのかをデータで証明したものです。読み取りのキャッシュがその請求額の 3 分の 2 であり、キャッシュに関するメンタル モデルを誤ると、Claude での残りの時間で間違った行を最適化することになります。
パート 1: キャッシュが実際にどのように機能するかを明らかにする 3 つの質問
初めてプロンプト キャッシュについて推論しようとすると、すぐに次の 3 つの疑問が浮かび上がります。
チャット メッセージを送信すると、実際にキャッシュが使用されますか?
コンテキスト ウィンドウが 53.6K トークンの場合、完全なペイロードは本当に毎ターン送信されるのでしょうか?
セッションを一時停止して 2 日後に再開するとどうなりますか?
質問は簡単そうに聞こえます。その答えは、ステートレス性、キャッシュ プレフィックス、5 分間の TTL がどのように相互作用するかを正確に明らかにしており、一部の使用パターンが安価であり、他の使用パターンが驚くほど高価である理由を説明しています。
質問 1 — チャット メッセージを送信するとき、クロードは実際にキャッシュを使用しますか?
はい、ほぼ確実に、そして積極的にです。これは、Claude Code のようなツールで送信するすべてのメッセージで起こっていることです。
E

メッセージを送信すると、クライアントは新しい API 呼び出しを実行して、これまでの会話全体を再送信します。ペイロードは大まかに次のようになります。
[ システム プロンプト + ツール定義 ] ← 大きく、静的、キャッシュ可能 [ ターン 1: メッセージ + アシスタントの返信 ] [ ターン 2: メッセージ + アシスタントの返信 ] ... [ ターン N: 最新のメッセージ ] ← 真に新しいバイトだけ
クライアントは、cache_control ブレークポイントを配置して、静的プレフィックスをキャッシュ可能としてマークします。新しいメッセージを送信するたびに、次のようになります。
システム プロンプト + ツール定義 → ~0.1x でのキャッシュ読み取り。 (多くの場合、このブロックだけで 10,000 ～ 20,000 トークンになります。)
以前の会話はすべて、→ 変更されず、以前に何も編集されていない限り、~0.1x でキャッシュ読み取りも行われます。
最新のメッセージ → 1x での新鮮な入力。
アシスタントの応答 → 5x で出力され、次のターンでキャッシュされたプレフィックスの一部になります。
そうですね、キャッシュはやり取りのチャットでも一生懸命働いてくれます。会話の経済的な形状は、大きな反復プレフィックスではなく、各ターンの終わりにある小さな新鮮なテールによって支配されます。
5分間のキャッチ。キャッシュ エントリは、最後の読み取りから 5 分間存続します。メッセージ間の距離を超えて離れると、エントリは期限切れになります。次のメッセージはキャッシュから読み取られなくなります。キャッシュを最初から書き直す方が有益です。これについては質問 3 でもう一度触れます。ここに最大の隠れたコストが隠れているからです。
プロバイダー間のメモ。同じ「すべてのメッセージがキャッシュを使用する」パターンが GPT (~1K トークンを超えるプロンプトに対してキャッシュ ルックアップが自動的に行われます) と Gemini (明示的に CachedContent オブジェクトを作成し、プレフィックスが ~32K の最小値を超えている場合のみ) にも当てはまります。 GPT では、それは簡単で目に見えません。 Gemini では、意図的なセットアップが必要です。非対称: 短いクロード チャットや GPT チャットは、コンテキストが大きくなると自動的にキャッシュされる利点があります。短い双子座

チャットではこの機能をまったく使用できません。
質問 2 — コンテキスト ウィンドウが 53.6K トークンである場合、クロードは実際に毎ターン 53.6K をすべて API に送信しますか?
はい、53.6K トークンのすべてのバイトが毎ターン、ネットワーク上を移動します。クロードはステートレスです。会話のサーバー側メモリはありません。毎回完全なペイロードが送信されます。
しかし、「発送済み」と「全額請求」は全く異なります。たとえば 200 トークンの新しいメッセージを送信すると、これらの 53.6K トークンに実際に何が起こるかは次のとおりです。
有効な入力コスト: 53,600 ではなく、~5,540 トークン相当。この 1 ターンで入力側はおよそ 10 倍減少します。
これにより、ネットワーク帯域幅と請求額が異なるという非対称性が生じることに注意してください。ネットワークは 53.6K トークンのペイロードを移動しています。 Anthropic の請求システムでは、入力の約 5.5,000 トークンに相当する料金が請求されます。ペイロード全体を解析する必要があり、モデルはコンテキストの 53.6K トークンすべてを認識する必要がありますが、ドルは小さな新鮮なテールにのみ対応します。
これは、クロードとの長い会話を経済的に実行可能にする中心的な洞察です。キャッシュがないと、50 ターンの会話の各ターンで、増加するプレフィックス全体が 1 倍のレートで再請求され、会話が進むにつれてターンごとのコストが膨れ上がります。キャッシュを使用すると、ターンごとの入力コストは、会話の既存の長さではなく、追加する新しいコンテンツの量にほぼ比例します。
小さなニュアンスが 1 つあります。この新しいメッセージを送信したとき、前のターンからのアシスタントの応答はまだキャッシュにありませんでした。前のターンに設定されたキャッシュ ブレークポイントは、前のメッセージのすべてをキャッシュしましたが、応答自体はキャッシュしませんでした。したがって、実際には、新しいリクエストは、このターンのcache_creation_inpuの一部として、以前の応答もキャッシュに書き込みます。

t_tokens ライン — 1.25 倍で 1 回請求されます。次のターン以降、その応答はキャッシュされたプレフィックスに結合され、0.1x で読み取られます。これは、次のパートで詳しく説明する「1 ターン ラグ」メカニズムと同じです。
プロバイダー間のメモ。 「全ペイロードが毎ターン送信される」部分は普遍的です。すべての主要な LLM API はステートレスであり、呼び出しごとに会話全体を再受信します。異なるのは、最終的に割引料金で請求される金額だけです。 Claude では ~10%、 GPT では ~50%、 Gemini では (キャッシュされた部分の) ~25% です。ネットワークのコストはプロバイダーを問わず同じです。請求コストは、キャッシングの仕組みによって差異が生じる部分です。
質問 3 — 作業を一時停止し、2 日後にセッションを再開した場合はどうなりますか?
キャッシュが消えてしまいました。そして、その結果は人々が予想するよりも大きなものになります。
5 分間の TTL は約 47 時間 55 分前に期限切れになりました。プレフィックスのキャッシュ エントリはもうありません。長い一時停止の後に最初のメッセージを送信するときは、次のようにします。
この最初に再開されたメッセージの実効入力コスト: ~66,950 トークン相当。
この数字を大局的に見るための 2 つの比較:
5 分間の TTL (約 5,540 トークン相当) 以内に送信されたメッセージと比較すると、約 12 倍のコストがかかります。キャッシュ割引はなくなりました。
キャッシュをまったく有効にせずに送信したメッセージ (1x で 53,600 トークン相当) と比較すると、25% 高くなります。キャッシュ書き込みには 1.25 倍のコストがかかるため、決して読み取ることができないキャッシュを再構築することは、最初からキャッシュしないよりも厳密には悪いです。
追加料金を払ってキャッシュを再構築すると、使用できるかどうかがわかりません。
次に何が起こるか。最初に再開されたメッセージがキャッシュを再構築する費用を支払うと、それから 5 分以内の後続のメッセージは通常どおり恩恵を受けます。以前と同様に、キャッシュ読み取りは 0.1 倍になります。コストの形状は、再構築に高価なスパイクが 1 つあり、その後、安価な運用に戻るというものです。
したがって、のライフサイクルは、

一時停止後に再開されたセッションは次のようになります。
再開メッセージ: コストが最大 12 倍に跳ね上がります。キャッシュが完全に再構築されました。
次の 5 分間のアクティビティ: 通常の安価な読み取り。
5 分を過ぎたら再度一時停止します。キャッシュの有効期限が切れると、次の再開で再びスパイクが発生します。
これにより、実際のワークフローでは非常に特殊な悪いパターンが生成されます。
「この古いセッションで簡単なフォローアップ メッセージを 1 つだけ送信します。」 → 約 53.6K の再構築コストを全額支払います。 → 返信が 1 件あります。 → セッションを終了します。 → 2日後、別の古いセッションで再度実行します。 → 再度完全な再構築を行います。
これを 1 週間に 5 つの異なる古いセッションにわたって行うと、再構築税を 5 回支払ったことになります。 「迅速なフォローアップ」は決して迅速なものではありません。
より良いパターン: 重要な作業のために古いセッションを再開する場合は、質問をまとめてください。最初のメッセージは再構築税を支払います。メッセージ 2 ～ 10 は安価です。作業を数日に分けて行うのではなく、一度に前倒しして作業を進めます。
本当に断続的な作業に最適なパターン: 古いセッションをまったく再開しないでください。実際に引き継ぐ必要があるコンテキストのみを使用して、新しいセッションを開始します。 5,000 トークンの概要を持つ新しいセッションは、小さなタスクのために 53.6,000 トークンの古いセッションを再開するよりも劇的に安くなります。また、新しいセッションに入ると、そのキャッシュはすぐに効果を発揮し始めます。
サブスクリプションプランでもこれが重要となる 2 つの理由
(トークンごとの API 請求ではなく) サブスクリプション プランのユーザーの場合、ドルコストはプロバイダーによって吸収されます。しかし、キャッシュの再構築は依然として実際の影響を及ぼします。
待ち時間。 53.6K トークンのキャッシュ エントリを再構築するのは、同じ 53.6K トークンをキャッシュから読み取るよりも遅くなります。長い一時停止後の最初のメッセージは、その後のメッセージよりも実際に遅く感じられます。多くのユーザーは理由を理解せずにこれに気づきました。
レート制限とキュー

太田。サブスクリプション プランであっても、キャッシュ書き込みはキャッシュ読み取りよりも使用量制限に対して大きくカウントされます。多くの古いセッションを習慣的に再開すると、予想よりも早くレート制限に達する可能性があります。
5 分間の TTL は請求の脚注ではありません。これは、長期にわたる作業をどのように構成するかについての実際の制約です。そして、それがわかると、うっかりそれに対処するのではなく、それに基づいて設計を開始します。
プロバイダー間のメモ。 GPT では、2 日間の再開のイメージは似ています。自動キャッシュはベストエフォート型であり、一晩では確実に存続しません。そのため、長い一時停止後の最初のメッセージは、入力価格の全額で実行されます。 Gemini では、答えはキャッシュの作成時に支払った TTL に完全に依存します。48 時間のキャッシュをプロビジョニングした場合、履歴書はそこから安価に読み取られます。 1 時間のキャッシュをプロビジョニングした場合、それはとうの昔に失われています。 Gemini は 3 つの中で唯一、「古いセッションを安価に再開する」という操作が可能です。その代償として、ストレージの使用料をずっと支払う必要があります。
これら 3 つの質問を合計するとどうなるか
3 つの小さな質問と 1 つの根本的な洞察: クロードのキャッシュは継続性に報酬を与え、散発的なタッチポイントを罰します。
継続的な会話では、キャッシュが味方になります。キャッシュは、有効な入力コストを桁違いに削減します。
5 分間の TTL を超えると、キャッシュは敵になります。キャッシュにより、メッセージに 25% のプレミアムが追加されます。

[切り捨てられた]

## Original Extract

The first post in this series, Part 1: A Real-World Case Study, ended with a single number: a 31-hour Claude Code session that cost $172.58, of which $114.98 — about 66% of the bill — was cache reads. Caching wasn't a side effect of that session. It was the dominant cost line, by volume so large tha
[truncated]

Deep Dive into LLM Token Cost — Blog Series Part 2: How Prompt Caching Actually Works
The first post in this series, Part 1: A Real-World Case Study , ended with a single number: a 31-hour Claude Code session that cost $172.58, of which $114.98 — about 66% of the bill — was cache reads. Caching wasn’t a side effect of that session. It was the dominant cost line, by volume so large that it outweighed every other line combined. Anyone trying to reason about LLM cost without a precise mental model of how the cache works is missing the part that matters most.
This post is that mental model. It’s not a strategies post — that’s the next one. It’s the mechanics post for caching specifically: how the prefix match actually behaves on the wire, what happens on the second message of a conversation versus the first, what happens when you walk away for two days, and what really happens to a single 800-token file you read into the session at turn 5.
A note on scope. Claude is used as the worked example throughout because its cache surface is the most explicit — every read, write, and TTL has a corresponding counter in the usage block, so the mechanics can be traced line by line. Where GPT or Gemini diverge in ways that change the practical answer, you’ll find a short cross-provider callout. The third post in this series goes much further on the cross-provider story; here the goal is to get the Claude mechanics rock-solid first.
Part 1 — Three questions that reveal how caching actually works. Does every chat message use the cache? Is the full context window really sent on every turn? What happens when you resume a session two days later? The short answers expose the asymmetry that governs almost every Claude cost decision.
Part 2 — A worked example. Following a single 800-token file through its full lifecycle in a 20-turn session: arrival as fresh input, the one-turn-lagged cache write, then cache reads for the rest of the session. Three diagrams make the mechanics tangible. Once you can trace 800 tokens, you can trace anything.
The series. This is the second of three posts:
Part 1: A Real-World Case Study (previous) — the mental model, anchored in a real $172 case study.
Part 2: How Prompt Caching Actually Works (this post) — three questions and one worked example.
Part 3: Strategies and Anti-Patterns (next) — five strategies ranked by impact, the silent failures that undo them, and the cross-provider comparison.
Each post stands on its own. But if you arrived here without reading the first one, the $172 case study is the proof-by-data of why this post matters: caching reads were two-thirds of that bill, and getting your mental model wrong about caching means optimizing the wrong line for the rest of your time on Claude .
Part 1: Three Questions That Reveal How Caching Actually Works
Three questions come up almost immediately the first time someone tries to reason about a prompt cache:
When I send a chat message, does the cache actually get used?
If my context window is 53.6K tokens, is the full payload really sent every turn?
What happens if I pause and resume the session two days later?
The questions sound simple. The answers expose precisely how statelessness, the cache prefix, and the 5-minute TTL interact — and they explain why some usage patterns are cheap and others are surprisingly expensive.
Question 1 — When I send a chat message, does Claude actually use the cache?
Yes — almost certainly, and aggressively. This is what’s happening on every message you send in a tool like Claude Code .
Every time you send a message, the client makes a fresh API call that re-ships the entire conversation so far. The payload looks roughly like:
[ System prompt + tool definitions ] ← large, static, cacheable [ Turn 1: your msg + assistant reply ] [ Turn 2: your msg + assistant reply ] ... [ Turn N: your newest message ] ← the only truly new bytes
The client places cache_control breakpoints to mark the static prefix as cacheable. On each new message you send:
System prompt + tool definitions → cache read at ~0.1x. (This block alone is often 10K–20K tokens.)
All prior conversation turns → also cache read at ~0.1x, as long as nothing earlier was edited.
Your newest message → fresh input at 1x.
The assistant’s reply → output at 5x, then becomes part of the cached prefix on your next turn.
So yes, the cache is working hard for you in any back-and-forth chat. The economic shape of the conversation is dominated by the small fresh tail at the end of each turn , not the large repeating prefix.
The 5-minute catch. Cache entries live for 5 minutes from the last read. If you walk away for longer than that between messages, the entry expires. Your next message no longer reads from cache — it pays to rewrite the cache from scratch. We’ll come back to this in Question 3, because it’s where the biggest hidden costs hide.
Cross-provider note. The same “every message uses the cache” pattern holds on GPT (cache lookup is automatic for prompts above ~1K tokens) and on Gemini (only if you’ve explicitly created a CachedContent object and the prefix exceeds the ~32K minimum). On GPT it’s effortless and invisible; on Gemini it requires deliberate setup. The asymmetry: short Claude or GPT chats benefit from caching automatically once your context grows; short Gemini chats can’t use the feature at all.
Question 2 — If my context window is 53.6K tokens, does Claude actually send all 53.6K back to the API every turn?
Yes — every byte of those 53.6K tokens travels over the wire on every turn. Claude is stateless: there is no server-side memory of your conversation. The full payload ships each time.
But “shipped” and “billed at full rate” are very different things. Here is what actually happens to those 53.6K tokens when you send a new message of, say, 200 tokens:
Effective input cost: ~5,540 token-equivalents, instead of 53,600. Roughly a 10x reduction on the input side, on this one turn.
Note the asymmetry this creates: network bandwidth and billing diverge. Your network is moving 53.6K tokens of payload. Anthropic’s billing system is charging you for the equivalent of ~5.5K tokens of input. The full payload still has to be parsed and the model still has to be aware of all 53.6K tokens of context — but the dollars correspond only to the small fresh tail.
This is the central insight that makes long Claude conversations economically viable. Without caching, every turn of a 50-turn conversation would re-bill the entire growing prefix at the 1x rate, and the per-turn cost would balloon as the conversation went on. With caching, the per-turn input cost stays roughly proportional to how much new content you’re adding, not how long the conversation already is.
One small nuance. The assistant’s reply from the previous turn was not in the cache yet when you sent this new message. The cache breakpoint placed on the previous turn cached everything through your previous message, but not the reply itself. So in practice the new request also writes that previous reply into the cache as part of this turn’s cache_creation_input_tokens line — billed once at 1.25x. From the next turn onward, that reply joins the cached prefix and reads at 0.1x. This is the same “one-turn lag” mechanic we’ll see in detail in the next part.
Cross-provider note. The “full payload ships every turn” part is universal — every major LLM API is stateless and re-receives the entire conversation each call. What differs is only how much of it ends up billed at the discounted rate: ~10% on Claude , ~50% on GPT , ~25% (of the cached portion) on Gemini . The network cost is identical across providers; the billing cost is where caching mechanics produce the divergence.
Question 3 — What if I pause my work and reopen the session two days later?
The cache is gone. And the consequences of that are bigger than people expect.
The 5-minute TTL expired ~47 hours and 55 minutes ago. There is no cache entry for the prefix anymore. When you send your first message after the long pause:
Effective input cost on this first resumed message: ~66,950 token-equivalents.
Two comparisons to put that number in perspective:
Versus a message sent within the 5-minute TTL (~5,540 token-equivalents): about 12x more expensive. The cache discount is gone.
Versus a message sent with caching never enabled at all (53,600 token-equivalents at 1x): 25% more expensive. Cache writes cost 1.25x, so rebuilding a cache you never get to read from is strictly worse than not caching in the first place.
You pay extra to rebuild a cache that you may or may not get to use.
What happens next. Once the first resumed message has paid to rebuild the cache, subsequent messages within 5 minutes of it benefit normally — cache reads at 0.1x, just like before. The cost shape is one expensive spike to rebuild, then back to cheap operation.
So the lifecycle of a resumed-after-pause session looks like:
Resume message: ~12x cost spike. Cache fully rebuilt.
Next 5 minutes of activity: normal cheap reads.
Pause again past 5 minutes: cache expires, the next resume pays the spike again.
This produces a very specific bad pattern in real workflows:
“I’ll just send one quick follow-up message in this old session.” → Pays the full ~53.6K rebuild cost. → Gets one reply. → Closes the session. → Two days later does it again with a different old session. → Pays another full rebuild.
If you do this across five different old sessions in a week, you’ve paid the rebuild tax five times. The “quick follow-up” is anything but quick on the bill.
Better pattern: If you’re going to resume an old session for substantial work, batch your questions . The first message pays the rebuild tax. Messages 2–10 are cheap. Front-load your work into one sitting instead of spreading it across days.
Best pattern for genuinely intermittent work: Don’t resume the old session at all. Start a fresh session with only the context you actually need carried forward. A new session with a 5K-token brief is dramatically cheaper than resuming a 53.6K-token old session for a small task — and once you’re in the fresh session, its cache starts paying off immediately.
Two reasons this matters even on subscription plans
For users on subscription plans (rather than per-token API billing), the dollar cost is absorbed by the provider. But the cache rebuild still has real consequences for you:
Latency. Rebuilding a cache entry on 53.6K tokens is slower than reading the same 53.6K tokens from cache. The first message after a long pause genuinely feels slower than subsequent ones. Many users have noticed this without understanding why.
Rate limits and quota. Even on subscription plans, cache writes count more heavily toward your usage limits than cache reads. Habitually resuming many stale sessions can cause you to hit rate limits faster than you’d otherwise expect.
The 5-minute TTL is not a billing footnote. It’s a real constraint on how to structure long-running work — and once you see it, you start designing around it instead of accidentally fighting it.
Cross-provider note. On GPT , the two-day resume picture is similar — automatic caching is best-effort and certainly won’t survive overnight, so your first message after a long pause runs at full input price. On Gemini the answer depends entirely on what TTL you paid for when you created the cache: if you provisioned a 48-hour cache, your resume reads from it cheaply; if you provisioned a 1-hour cache, it’s long gone. Gemini is the only one of the three where “resume an old session cheaply” is a knob you can turn — at the cost of paying storage rent the whole time.
What these three questions add up to
Three small questions, one underlying insight: Claude ‘s cache rewards continuity and punishes sporadic touch-points.
Within a continuous conversation, the cache is your friend — it reduces your effective input cost by an order of magnitude.
Across the 5-minute TTL, the cache is your enemy — it adds a 25% premium to a message that wo

[truncated]
