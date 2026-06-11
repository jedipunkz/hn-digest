---
source: "https://www.axamy.com/blog/tool-budget"
hn_url: "https://news.ycombinator.com/item?id=48496230"
title: "Every LLM Tool Call Needs an Output Budget"
article_title: "Every LLM Tool Call Needs an Output Budget"
author: "jhonovich"
captured_at: "2026-06-11T21:45:16Z"
capture_tool: "hn-digest"
hn_id: 48496230
score: 2
comments: 0
posted_at: "2026-06-11T20:50:54Z"
tags:
  - hacker-news
  - translated
---

# Every LLM Tool Call Needs an Output Budget

- HN: [48496230](https://news.ycombinator.com/item?id=48496230)
- Source: [www.axamy.com](https://www.axamy.com/blog/tool-budget)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T20:50:54Z

## Translation

タイトル: すべての LLM ツール呼び出しには出力バジェットが必要です
説明: ツールの出力により、コスト、レイテンシ、コンテキストの使用量が密かに爆発的に増加する可能性があります。ツールが詳細な API オブジェクト、メタデータ、コメント、ログ、または多すぎる結果を返す場合、短いユーザー リクエストが巨大なモデル リクエストになる可能性があります。解決策は、ツールを諦めないことです。実際のエージェント トレースをプロファイリングし、ツールの出力を次のように処理します。
[切り捨てられた]

記事本文:
すべての LLM ツール呼び出しには出力バジェットが必要です
ツールは AI エージェントが役立つ主な理由の 1 つですが、ツールの出力によりコスト、レイテンシ、コンテキストの使用量が静かに爆発する可能性があります。ツールが詳細な API オブジェクト、メタデータ、コメント、ログ、または多すぎる結果を返す場合、短いユーザー リクエストが巨大なモデル リクエストになる可能性があります。解決策は、ツールを諦めないことです。これは、実際のエージェント トレースをプロファイリングし、ツールの出力を第一級の最適化問題として扱うことです。デフォルトで最小限の有用な情報を返し、さらに必要な場合はドリルダウン パスを提供します。
強力なエンジニアリング チームがすでにこれのバージョンを実行していると確信しています。ツールの出力予算に特に焦点を当てた実践的な投稿があまり見つからなかったため、この記事を書いています。良い例をご存知でしたら、リンクをいただければ幸いです。
ツールに対する興奮とリスク
最新の AI システムで最も興味深いものの 1 つはツールです。モデルを検索システム、データベース、CRM、ブラウザ自動化、内部 API、電子メール システム、カレンダー、サードパーティ統合に接続できることは、本当に強力です。エージェントが複数のツールを組み合わせて質問に答えたり、タスクを完了したりすると、まるで魔法のように感じられることがあります。チャットボットというよりも、さまざまなシステムからの情報をつなぎ合わせている人のように感じられ始めます。
あの興奮は本物だ。それは危険でもあります。
ツールが機能すると、本能的にさらに追加したくなるものです。より多くの API。さらなる統合。さらなるアクション。より多くのデータソース。システムの機能が向上し、デモもより印象的になります。
しかし、私たちは実際の痕跡をより注意深く観察し始めました。
エージェントは何回ツール呼び出しを行いましたか?
各ツール呼び出しのリターンはいくらですか?
ユーザーの短い質問が 20 回のツール呼び出しを引き起こしましたか?
これらのツール呼び出しのそれぞれが、数千または数万のトークンをモデルに送り返しましたか?
それは

不快な部分です。エージェントは、必要以上に多くのトークンを使用しながら、高いレベルで正しいことを行っている可能性があります。
これは、エージェント、ツール、MCP、統合プラットフォーム、またはサードパーティのサービスに対して反論する投稿ではありません。まったく逆です。最新の AI システムが役立つ主な理由の 1 つはツールです。しかし、ツールの出力が付随的なものとして扱われると、エージェントはすぐに遅くなり、コストがかかり、コンテキストが制限される可能性があります。
通常の LLM リクエストは小さい場合があります。ユーザーは短い質問をするかもしれません。モデルは短い答えを生成するかもしれません。単純な数学または推論の質問には、わずか数十の入力トークンと数百の出力トークンが含まれる場合があります。それにはわずか 1 セントの費用がかかります。
しかし、ひとたびツールがループに入ると、経済状況は完全に変化する可能性があります。
ユーザーは短い質問をし、エージェントはツールを呼び出し、ツールは非常に大きな応答を返すことができます。場合によっては、その応答に有益な情報が含まれることがあります。多くの場合、メタデータ、タイムスタンプ、コメント、監査フィールド、内部 ID、冗長フィールド、冗長レコード、無関係な結果、行数が多すぎる、ページ コンテンツが多すぎる、ログ出力が多すぎるなども含まれます。
今では、その要求はもはや小さくありません。ユーザーのプロンプトは 30 トークンだったかもしれませんが、ツールの結果は 30,000 または 100,000 トークンになる可能性があります。
これにより、すぐに次の 3 つの問題が発生します。
モデル API はトークンによって課金されるため、コストがかかります。
レイテンシー。コンテキストが大きいと処理に時間がかかるためです。
トークンの最大化。モデルが作業の有用な部分を実行する前に、リクエストがコンテキストの制限に達する可能性があるためです。
エージェントが予期せぬ高額な請求書を作成したり、文脈を無視したり、規模が非現実的であると感じたりすることについて人々が不満を言うとき、私はこれが理由の一部であることが多いのではないかと思います。問題は、必ずしもエージェントが本質的に無駄であるということではありません。多くの場合、問題はツールが適切でないことです。

モデルが実際に必要とするものに合わせて最適化されていない。
私たちの行動を変えたもの: クエリごとのコストの内訳
私たちが行った最も有益なことは、巧妙なプロンプトではありませんでした。
すべての内部エージェントのクエリについて、チームがコストを検査し、内訳を追跡できるようにしました。最終的な答えだけではありません。道全体。
各ツールの利益はいくらですか?
各ステップでどれくらいのコンテキストが消費されましたか?
モデルコールの費用はいくらかかりましたか?
エージェントはツール間を行き来しましたか?
ツールを呼び出して失敗し、別のツールを呼び出して、最初のツールに戻ったのでしょうか?
1 つのツールが多すぎる結果を返したため、短いユーザーの質問が巨大なモデル リクエストに変わってしまったのでしょうか?
この可視性により、システムのデバッグ方法が変わりました。
それ以前は、クエリは外から見れば問題ないように見えたかもしれません。答えは正しいかもしれません。 UIは普通に感じるかもしれません。しかしその根底には、エージェントが高価な呼び出しを連鎖的に行ったり、ツールの出力が多すぎたり、不要なコンテキストを後のステップに持ち込んだりした可能性があります。
社内の全員が内訳を確認できるようになると、パターンが明らかになりました。
ツールは 6 フィールドが必要なときに 40 フィールドを返しました。
検索では 5 件で十分なはずの結果が 30 件返されました。
統合には、すべてのコメント、タイムスタンプ、メタデータ フィールドが含まれます。
エージェントは 4 件で機能するはずのときに 15 件の電話をかけました。
これらは哲学的な問題ではありませんでした。それらは痕跡の中に見えました。
内部でトレース ビューを使用するたびに、何かを学びました。明らかなバグが見つかることもありました。場合によっては、非効率的なツールの説明が見つかることがありました。ツールが正しい動作をしているにもかかわらず、間違った形式の出力を返していることが判明することがありました。
それが、私たちを生産高予算の考え方へと押し上げた理由です。
クエリごとのコストの内訳がないと、「ツール出力の最適化」という言葉は抽象的に聞こえます。彼らを見れば、それが明らかになる。トークンがどこに向かっているのかを正確に確認できます。
私たちの間違いはドラマではありませんでした

チック。私たちはデータベース全体を意図的にモデルにダンプしていませんでした。私たちは何百万ものレコードをコンテキストに送信していませんでした。ツールの多くは合理的でした。システムは機能しました。エージェントは役に立ちました。
間違いは、ツールの出力を二次的なものとして扱っていたことです。
ツールが正しい情報を返した場合、私たちは当初、ツールが機能していると考える傾向がありました。しかし、正しさだけでは十分ではありません。ツールは正しい情報を返しても、非効率的に情報を返すことがあります。含まれるレコードが多すぎる、フィールドが多すぎる、メタデータが多すぎる、周囲のコンテキストが多すぎる、または次に何をすべきかについてのガイダンスが少なすぎる可能性があります。
それが私たちにとっての考え方の変化でした。
このツールは、適切なコストで、適切な量の情報を適切な形で返しましたか?
これは、基本的な機能の開発よりもパフォーマンスの最適化に近いものです。最初のバージョンでは、ツールが機能することが証明されています。次のバージョンでは、効率的に動作するかどうかが問われます。
ツールは有用な答えを返しますが、必要以上に多くのフィールドが含まれることもあります。検索すると関連する結果が返されますが、その結果が多すぎます。統合は有効な API オブジェクトを返しますが、そのオブジェクトは LLM コンテキスト ウィンドウ用ではなくソフトウェア用に設計されています。ツールの説明には、そのツールが何を行ったかが記載されていますが、何が返されるべきかを十分に制約するものではありません。
これらはどれも単独では壊滅的なものには見えませんでした。しかし、多くのツールと多くのエージェントの手順を考慮すると、コストは急速に増大しました。
トレースとトークンの使用状況を慎重に調査し始めると、多くのツール出力を大幅に削減できることがわかりました。場合によっては、削減は桁違いのレベルでした。すべてのツールにそれほど無駄があったわけではなく、節約の一部は私たち自身の初期の間違いを修正することで得られました。しかし、パターンは明らかだったので、ツールに対する考え方が変わりました。
すべての LLM ツール呼び出しには出力が必要です

予算。
ツールの出力が次のプロンプトになる
重要なメンタル モデルはシンプルです。ツールの出力は単なる API 応答ではありません。ツールの出力は次のプロンプトの一部になります。
モデルがそれを読み上げます。プロバイダーが料金を請求します。コンテキスト ウィンドウはそれを保持する必要があります。エージェントの次のステップはそれに依存します。
つまり、ツールの出力設計は二次的な詳細ではないということです。迅速な設計、コスト管理、レイテンシ管理、コンテキスト管理を一度に実現します。
ツールが 50,000 個の不要なトークンを返した場合、それらのトークンはツール内に残りません。それらはモデルの入力になります。
これが、「より大きなコンテキスト ウィンドウを使用するだけ」が完全な答えではない理由です。コンテキスト ウィンドウが大きくなると役に立ちます。それらは役に立ちます。ただし、過剰なツール出力が無料になるわけではありません。
モデルが 200,000 個のトークンを受け入れることができる場合でも、100,000 個の不要なトークンを送信するとコストがかかり、遅延が増加し、タスクの後半でエージェントが制限に達する可能性が高くなります。
コンテキスト ウィンドウを大きくすると、無駄を吸収できます。無駄がなくなるわけではありません。
生産予算は単なるハードトークンの上限ではありません。ハードキャップは便利ですが、それだけでは十分ではありません。
このアイデアの悪いバージョンは次のとおりです。
最初の 10,000 文字を返し、残りを切り取ります。
それが切り捨てです。必要な場合もありますが、それは乱暴な最終手段です。
このツールがデフォルトで返す必要がある最小限の有用な出力は何ですか?
絶対的な最小値ではありません。やみくもに短縮された応答ではありません。最低限の有用な応答。
これは通常、モデルが次の決定を下すための十分な情報に加え、必要に応じてさらに情報を求める方法を意味します。
たとえば、検索ツールは、デフォルトでは一致するすべてのドキュメントを返さないことがよくあります。次のようなものを返すことができます。
{
「カウント」 : 184 、
"summary" : "一致する結果のほとんどは、最新のファームウェア更新後のログイン失敗に関するものです。" 、
"top_results" : [
{
"id" : "result_1" ,
"title" : "ログイン失敗

ファームウェア5.2インチ以降、
"why_relevant" : "報告された問題と直接一致します"
} 、
{
"id" : "結果_2" ,
"title" : "セッションタイムアウトの増加" ,
"why_relevant" : "関連する可能性が高い障害モード"
}
]、
"ハンドル" : "search_abc123" ,
"next_actions" : [
"fetch_result" ,
"fetch_more_results" ,
「結果内の検索」
]
ログ ツールは、生のログの前にクラスター、カウント、および代表的な行を返すと、より適切に機能することがよくあります。
{
"summary" : "主な障害は、14:03 UTC から始まる AuthTimeoutError です。" 、
"error_clusters" : [
{
"error" : "AuthTimeoutError" ,
「カウント」 : 3812 、
"first_seen" : "14:03" ,
"last_seen" : "14:18"
} 、
{
"エラー" : "データベース接続エラー" ,
「カウント」 : 102 、
"first_seen" : "14:07" ,
"last_seen" : "14:09"
}
]、
"代表例" : [
「14:03 AuthTimeoutError: トークンの更新が 5000 ミリ秒を超えました」
]、
"ハンドル": "logs_456"
CRM またはサードパーティ統合では、多くの場合、アップストリーム API が公開するすべてのフィールドを返す必要はありません。必要な場合にのみ完全なオブジェクトをフェッチする方法を使用して、タスクに関連するフィールドを返す必要があります。
目標は、モデルから情報を隠すことではありません。目標は、まだ必要のない情報を読み取るためにモデルに料金を支払わせないようにすることです。
モデルが実際に生の出力を必要とする場合があります。重要なのはそれを禁止しないことです。重要なのは、すべての統合のデフォルト結果ではなく、生の出力を明示的に選択することです。
私たちが便利だと感じたパターンの 1 つは、生の統合と LLM の間にアダプターを置くことです。
統合 → 出力アダプター → LLM アダプターの仕事は、システム形状の出力をモデル形状の出力に変換することです。
投影。モデルに必要のないフィールドを削除します。
ランキング。最も関連性の高いアイテムを最初に返します。
集計。生の行の代わりにカウント、合計、またはグループ化された結果を返します。
クラスタリング、グループ化ログ、エラー、ドキュメント

nts、または意味のあるバケットにレコードを保存します。
サンプリング。すべての例ではなく代表的な例を返します。
要約。詳細な出力をタスク関連の要約に圧縮します。
生成を処理し、生の結果を別の場所に保存し、モデルが後で拡張できる参照を返します。
このアダプターは、ベクトル データベースや別の LLM 呼び出しを常に必要とするわけではありません。場合によっては、最良のアダプターが単により優れたクエリであることもあります。ランキング機能の場合もあります。場合によってはまとめになることもあります。キャッシュされたハンドルである場合もあります。場合によっては、ツールが何を返すかをより厳密に制約するツールの説明が使用されます。
重要な点は、統合からの生の出力がモデルの理想的な入力になることはめったにないということです。
サードパーティの統合によりこれがさらに重要になります
これは、サードパーティの統合レイヤーを使用する場合に特に重要になります。
これらのシステムは、多くの外部サービスに迅速に接続できるため、貴重です。それがポイントです。すべての統合を自分で構築しなくても、エージェントに多くの便利なアクションやデータ ソースへのアクセスを突然与えることができます。
それは強力です。まさにこれが、生産量の予算設定が重要な理由でもあります。
外部 API 応答のデフォルトの形状は、必ずしも LLM にとって正しい形状であるとは限りません。 API 対応

[切り捨てられた]

## Original Extract

Tool outputs can quietly explode cost, latency, and context usage. A short user request can become a huge model request if tools return verbose API objects, metadata, comments, logs, or too many results. The fix is not to give up on tools. It is to profile real agent traces and treat tool output as
[truncated]

Every LLM Tool Call Needs an Output Budget
Tools are one of the main reasons AI agents are useful, but tool outputs can quietly explode cost, latency, and context usage. A short user request can become a huge model request if tools return verbose API objects, metadata, comments, logs, or too many results. The fix is not to give up on tools. It is to profile real agent traces and treat tool output as a first-class optimization problem: return the minimum useful information by default, with drill-down paths when more is needed.
I’m sure strong engineering teams are already doing versions of this. I’m writing because I could not find many practical posts focused specifically on tool output budgets. If you know good examples, I’d appreciate links.
Excitement And Risk Over Tools
One of the most exciting things about modern AI systems are tools. Being able to connect models to search systems, databases, CRMs, browser automation, internal APIs, email systems, calendars, and third-party integrations is genuinely powerful. When an agent combines multiple tools to answer a question or complete a task, it can feel almost magical. It starts to feel less like a chatbot and more like a person piecing together information from different systems.
That excitement is real. It is also dangerous.
Once a tool works, the natural instinct is to add more. More APIs. More integrations. More actions. More data sources. The system becomes more capable, and the demos become more impressive.
But we started looking more carefully at the actual traces.
How many tool calls did the agent make?
How much did each tool call return?
Did a short user question trigger twenty tool calls?
Did each of those tool calls send thousands or tens of thousands of tokens back into the model?
That was the uncomfortable part. The agent could be doing the right thing at a high level while still using far more tokens than necessary.
This is not a post arguing against agents, tools, MCP, integration platforms, or third-party services. Quite the opposite. Tools are one of the main reasons modern AI systems are useful. But if tool outputs are treated as incidental, agents can become slow, expensive, and context-limited very quickly.
A normal LLM request can be tiny. A user might ask a short question. A model might produce a short answer. A simple math or reasoning question might involve only a few dozen input tokens and a few hundred output tokens. That can cost a tiny fraction of a cent.
But once tools enter the loop, the economics can change completely.
A user can ask a short question, the agent can call a tool, and the tool can send back a very large response. Sometimes that response includes useful information. Often it also includes metadata, timestamps, comments, audit fields, internal IDs, redundant fields, verbose records, irrelevant results, too many rows, too much page content, or too much log output.
Now the request is no longer small. The user’s prompt may have been 30 tokens, but the tool result might be 30,000 or 100,000 tokens.
That creates three problems immediately:
Cost, because model APIs charge by tokens.
Latency, because large contexts take longer to process.
Token maxing, because the request can hit the context limit before the model has done the useful part of the work.
When people complain about agents producing unexpectedly large bills, burning through context, or feeling impractical at scale, I suspect this is often part of the reason. The problem is not necessarily that agents are inherently wasteful. The problem is often that the tools are not optimized for what the model actually needs.
The Thing That Changed Our Behavior: Per-Query Cost Breakdowns
The most useful thing we did was not a clever prompt.
For every internal agent query, we made it possible for our team to inspect a cost and trace breakdown. Not just the final answer. The whole path.
How much did each tool return?
How much context did each step consume?
How much did the model call cost?
Did the agent bounce between tools?
Did it call a tool, miss, call another tool, then come back to the first one?
Did a short user question turn into a huge model request because one tool returned too much?
That visibility changed how we debugged the system.
Before that, a query could look fine from the outside. The answer might be correct. The UI might feel normal. But underneath, the agent might have made a chain of expensive calls, included too much tool output, or carried unnecessary context into later steps.
Once everyone internally could see the breakdown, patterns became obvious.
A tool returned 40 fields when 6 were needed.
A search returned 30 results when 5 would have been enough.
An integration included every comment, timestamp, and metadata field.
An agent made 15 calls when 4 should have worked.
These were not philosophical problems. They were visible in the trace.
Every time we used the trace view internally, we learned something. Sometimes we found obvious bugs. Sometimes we found inefficient tool descriptions. Sometimes we found that a tool was doing the right thing but returning the wrong shape of output.
That is what pushed us toward the output-budget mindset.
Without per-query cost breakdowns, “tool output optimization” sounds abstract. With them, it becomes obvious. You can see exactly where the tokens are going.
Our mistake was not dramatic. We were not intentionally dumping entire databases into models. We were not sending millions of records into context. Many of the tools were reasonable. The system worked. The agents were useful.
The mistake was treating tool output as secondary.
If the tool returned the correct information, we initially tended to think the tool was working. But correctness is not enough. A tool can return the correct information and still return it inefficiently. It can include too many records, too many fields, too much metadata, too much surrounding context, or too little guidance about what to do next.
That was the shift in thinking for us.
Did the tool return the right amount of information, in the right shape, at the right cost?
It is closer to performance optimization than basic feature development. The first version proves the tool works. The next version asks whether it works efficiently.
A tool would return a useful answer, but also include more fields than necessary. A search would return relevant results, but too many of them. An integration would return a valid API object, but the object was designed for software, not for an LLM context window. A tool description would say what the tool did, but not firmly enough constrain what should come back.
None of this looked catastrophic in isolation. But across many tools and many agent steps, the costs compounded quickly.
Once we started examining traces and token usage carefully, we found that a number of tool outputs could be reduced significantly. In some cases, the reductions were on the order-of-magnitude level. Not every tool had that much waste, and some of the savings came from fixing our own early mistakes. But the pattern was clear enough that we changed how we think about tools.
Every LLM tool call needs an output budget.
Tool Outputs Become the Next Prompt
The key mental model is simple: a tool output is not just an API response. A tool output becomes part of the next prompt.
The model reads it. The provider charges for it. The context window has to hold it. The next step of the agent depends on it.
That means tool output design is not a secondary detail. It is prompt design, cost control, latency control, and context management all at once.
If a tool returns 50,000 unnecessary tokens, those tokens do not stay inside the tool. They become model input.
This is why “just use a bigger context window” is not a full answer. Bigger context windows help. They are useful. But they do not make excessive tool output free.
Even if a model can accept 200,000 tokens, sending 100,000 unnecessary tokens still costs money, adds latency, and increases the chance that the agent hits limits later in the task.
A bigger context window can absorb waste. It does not eliminate waste.
An output budget is not just a hard token cap. A hard cap is useful, but it is not enough.
A bad version of this idea is:
Return the first 10,000 characters and cut off the rest.
That is truncation. Sometimes it is necessary, but it is a crude last resort.
What is the minimum useful output this tool should return by default?
Not the absolute minimum. Not a blindly shortened response. The minimum useful response.
That usually means enough information for the model to make the next decision, plus a way to ask for more if needed.
For example, a search tool often should not return every matching document by default. It can return something like:
{
"count" : 184 ,
"summary" : "Most matching results concern login failures after the latest firmware update." ,
"top_results" : [
{
"id" : "result_1" ,
"title" : "Login failures after firmware 5.2" ,
"why_relevant" : "Direct match to the reported issue"
} ,
{
"id" : "result_2" ,
"title" : "Session timeout increase" ,
"why_relevant" : "Likely related failure mode"
}
] ,
"handle" : "search_abc123" ,
"next_actions" : [
"fetch_result" ,
"fetch_more_results" ,
"search_within_results"
]
} A log tool often works better when it returns clusters, counts, and representative lines before raw logs:
{
"summary" : "The dominant failure is AuthTimeoutError beginning at 14:03 UTC." ,
"error_clusters" : [
{
"error" : "AuthTimeoutError" ,
"count" : 3812 ,
"first_seen" : "14:03" ,
"last_seen" : "14:18"
} ,
{
"error" : "DatabaseConnectionError" ,
"count" : 102 ,
"first_seen" : "14:07" ,
"last_seen" : "14:09"
}
] ,
"representative_examples" : [
"14:03 AuthTimeoutError: token refresh exceeded 5000ms"
] ,
"handle" : "logs_456"
} A CRM or third-party integration often should not return every field the upstream API exposes. It should return the fields relevant to the task, with a way to fetch the full object only if needed.
The goal is not to hide information from the model. The goal is to avoid making the model pay to read information it does not need yet.
There are cases where the model really does need raw output. The point is not to forbid that. The point is to make raw output an explicit choice, not the default result of every integration.
One pattern we have found useful is putting an adapter between raw integrations and the LLM.
Integration → Output Adapter → LLM The adapter’s job is to turn system-shaped output into model-shaped output.
Projection, removing fields the model does not need.
Ranking, returning the most relevant items first.
Aggregation, returning counts, totals, or grouped results instead of raw rows.
Clustering, grouping logs, errors, documents, or records into meaningful buckets.
Sampling, returning representative examples instead of all examples.
Summarization, compressing verbose output into a task-relevant summary.
Handle generation, storing the raw result elsewhere and returning a reference the model can expand later.
This adapter does not always require a vector database or another LLM call. Sometimes the best adapter is just a better query. Sometimes it is a ranking function. Sometimes it is a summary. Sometimes it is a cached handle. Sometimes it is a tool description that more firmly constrains what the tool should return.
The important point is that the raw output from an integration is rarely the ideal input for a model.
Third-Party Integrations Make This More Important
This becomes especially important when using third-party integration layers.
These systems are valuable because they let you connect to many external services quickly. That is the point. You can suddenly give an agent access to many useful actions and data sources without building every integration yourself.
That is powerful. It is also exactly why output budgeting matters.
The default shape of an external API response is not necessarily the right shape for an LLM. An API resp

[truncated]
