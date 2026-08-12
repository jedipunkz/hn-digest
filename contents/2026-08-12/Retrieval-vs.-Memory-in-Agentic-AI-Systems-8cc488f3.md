---
source: "https://machinelearningmastery.com/retrieval-vs-memory-in-agentic-ai-systems/"
hn_url: "https://news.ycombinator.com/item?id=49271447"
title: "Retrieval vs. Memory in Agentic AI Systems"
article_title: "Retrieval vs. Memory in Agentic AI System"
author: "eigenBasis"
captured_at: "2026-08-12T12:44:32Z"
capture_tool: "hn-digest"
hn_id: 49271447
score: 1
comments: 0
posted_at: "2026-08-12T12:37:08Z"
tags:
  - hacker-news
  - translated
---

# Retrieval vs. Memory in Agentic AI Systems

- HN: [49271447](https://news.ycombinator.com/item?id=49271447)
- Source: [machinelearningmastery.com](https://machinelearningmastery.com/retrieval-vs-memory-in-agentic-ai-systems/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T12:37:08Z

## Translation

タイトル: Agentic AI システムにおける検索とメモリ
記事のタイトル: Agentic AI システムにおける検索とメモリ
説明: この記事では、長時間実行される AI エージェントにおける検索とメモリの仕組みと、それぞれがどこに当てはまるかを詳しく説明します。また、エージェントが不要なコンテキストを持たずに関連情報にアクセスできるように、それらを効果的に組み合わせる方法についても検討します。

記事本文:
Agentic AI システムにおける検索とメモリ
ナビゲーション
開発者を機械学習で優れたものにする
開発者を機械学習で優れたものにする
コード アルゴリズム 機械学習アルゴリズムを最初から実装します。
ディープラーニング (keras) ディープラーニング
時系列予測のためのニューラルネット時系列ディープラーニング
LSTM 長短期記憶ネットワーク
開発者を機械学習で優れたものにする
開発者を機械学習で優れたものにする
コード アルゴリズム 機械学習アルゴリズムを最初から実装します。
ディープラーニング (keras) ディープラーニング
時系列予測のためのニューラルネット時系列ディープラーニング
LSTM 長短期記憶ネットワーク
Agentic AI システムにおける検索とメモリ
シェアする
ポスト
シェアする
この記事では、エージェント AI システムにおける検索と記憶の概念的および実践的な違いと、両方を効果的に組み合わせる方法について学びます。
検索とメモリを区別するもの、および長期実行エージェントにとってその区別がなぜ重要なのか。
検索パイプラインとメモリ システムがそれぞれどのように構築されるかを具体的な例で示します。
検索と記憶を単一の効果的なエージェント アーキテクチャに組み合わせる方法。
以前のやり取りを覚えていない AI エージェントはあまり役に立ちません。すべての大規模な言語モデルには固定のコンテキスト ウィンドウがあり、会話、一連のツール出力、または取得したドキュメントの山がその制限を超えると、何かを削除したり、要約したり、新たに取得したりする必要があります。長時間実行されるエージェントを構築している開発者は、常にこの問題に遭遇します。エージェントは、すでに回答した質問を再質問したり、以前に下した決定に矛盾したり、必要な文書が存在することさえ認識できなかったりします。
検索と記憶はこれに対処する 2 つのメカニズムであり、問​​題の異なる部分を解決します。レット

rieval は、ドキュメント、コード、データベース レコードなど、モデルがトレーニングされておらず、デフォルトで保持する必要のない外部の知識を取り込みます。メモリは、エージェント自体が学習または実行した内容をセッション全体または複数のセッションにわたって保持するため、毎回ゼロから開始されるわけではありません。 2 つを混同したり、1 つだけを構築したりすると、多くのエージェント アーキテクチャが破綻します。この記事の内容は次のとおりです。
概念レベルで検索と記憶を分けるものは何ですか
取得パイプラインとメモリ システムがそれぞれどのように構築されるか、実際の例を示します
両者を並べて比較すると、
両方を単一の効果的なエージェント システムに組み合わせる方法
そもそもなぜ分割が存在するのかということから始めます。
コンテキストによって分割が強制される理由を理解する
コンテキスト ウィンドウは、システム プロンプト、会話履歴、ツール出力、事前に挿入されたものなど、モデルが一度に確認できるトークンの合計セットです。これは有限であり、その中のすべてのトークンが転送のたびに処理されるため、単にウィンドウを大きくするだけでは、期待どおりに拡大することはできません。コンテキスト エンジニアリングは、その限られたリソースをキュレーションおよび管理する学問として登場し、それを単なる命令を詰め込む場所ではなく、特定の瞬間にモデルが利用できる完全な状態として扱います。
この制約を考慮すると、エージェントには必要だがコンテキスト内に永続的に保持できない 2 種類の情報があります。
モデルの外部および現在の会話の外部に存在する情報 (ナレッジ ベース、コードベース、または一連のポリシー文書など)。これが検索で処理されることです。
エージェントが自ら生成または学習した情報で、現在のコンテキスト ウィンドウを超えて存続する必要がある情報 (10 ターン前に行われた決定や特定のユーザーに関する事実など)。これはメモリが処理するものです。
どちらも同様のツールを使用して実装されます。

埋め込み、ベクトル検索、構造化ストア。主な違いは、保存する内容と情報の出所です。検索ではエージェントの外部にあるコーパスを検索しますが、メモリではエージェント自身の対話や過去の行動からの情報が保存されます。
エージェントシステムでの取得の定義
検索とは、エージェントが「私の体重や現在の状況では把握していないことについて、世界は何を知っているのか」に答える方法です。最も一般的な実装は、検索拡張生成 (RAG) です。
ソース文書は、役に立つ程度に小さなパッセージに分割されます。
各チャンクはエンベディングに変換され、ベクトル インデックスに格納されます。
クエリ時に、受信した質問は同じ方法で埋め込まれ、インデックスは最も近い一致を返します。
これらの一致は、ユーザーの質問とともにプロンプ​​トに挿入されます。
このパターンは通常、取得ステップを残りのエージェントの推論に結び付けるオーケストレーション層を備えたマネージド データストア上で実行されます。これは、今日の実稼働環境におけるほとんどの取得拡張生成アーキテクチャの背後にあるアプローチです。コーパス自体は共有されており、同じ製品ドキュメントについて質問するすべてのユーザーが同じインデックスにヒットします。また、個々の会話とは関係なく、独自のスケジュールで更新されます。
エージェントティック システムでのメモリの定義
記憶とは、エージェントが「すでに学習したこと、または実行したことのうち、今後も継続する必要があること」に答える方法です。これは、動作が異なる 2 つのレイヤーに分割されます。
短期記憶とは、実行中のセッションの状態、つまりこれまでの会話に加えて、現在のタスク中にエージェントがスクラッチパッドに書き込んだ内容のことです。安価で、セッションが終了すると消滅します。
長期記憶はセッションをまたいで持続します。検索よりも難しい質問、つまり「何が関連しているか」だけでなく「何が関連しているのか」に答えなければなりません。

まず最初に保持する価値があります。」
一部のエージェント記憶システムは、会話から有用な事実、好み、コンテキストを自動的に抽出し、後で使用できるように保存します。新しいセッションの開始時に、エージェントは検索インデックスをクエリするのと同じようにメモリにクエリを実行できますが、結果は共有ドキュメント コーパスではなく、ユーザー、タスク、またはエージェントに固有のものになります。このレイヤーを設計するとき、チームは、保存および取得する必要があるものに応じて、さまざまなエージェント メモリ戦略とエージェント メモリ フレームワークを検討できます。
簡単に実行した例では、分割コンクリートを作成します。顧客が注文の遅れについてサポート エージェントにメッセージを送りました。
注文が遅れた場合、エージェントはまず顧客の以前の履歴の記憶を確認します。 3週間前のメモには、電子メールでのフォローアップを希望し、同様の配送問題は部分返金で解決したという内容が記載されていた。それは、この特定の顧客に関するエージェントの記録から得られたものであるため、記憶です。
次に、エージェントは先月変更された現在の配送ポリシーが必要となるため、会社のドキュメントを検索し、関連するセクションを取得します。情報は外部ソースから取得され、すべての顧客に適用されるため、これが検索です。どちらの結果も同じプロンプトに追加されますが、回答する質問は異なります。
検索と記憶の比較
並べて配置すると、検索と記憶の違いが一目で分かりやすくなります。
上の表にリストされている障害モードは、2 つのうちの 1 つだけを使用して構築されたエージェントが予測可能な形で壊れる傾向がある理由、およびほとんどの動作中のシステムが最終的に両方を必要とする理由を説明しています。
検索と記憶を効果的なシステムに組み合わせる
検索能力はあるが記憶がないエージェントは、セッションごとに同じ結論を再度導き出し、何もパーソナライズできません。年齢

記憶はあるが検索はできないが、それ自体の歴史は知っているが、その歴史の外にあるものに自分自身を根付かせる方法がない。トレーニング データの終了後に変更されたポリシーに関する質問には答えることができません。組み合わせを正しく行うには、次のようなことが必要になります。
フィルタリングはウィンドウ サイズよりも重要です。取得した文書や記憶エントリをさらに追加しても、必ずしも回答が改善されるわけではありません。モデルは追加のトークンをすべて処理して重み付けする必要があるため、ある点を超えると、追加のコンテキストによって答えが悪化する可能性があります。対象を絞った小規模な検索は、1 回の大規模な検索よりも効果的であることが多く、取得トークンの効率を維持できます。
古いことは、検索と記憶では異なる働きをします。基礎となるドキュメントが再インデックス付けされずに変更されると、検索インデックスは古くなります。ユーザーに関する情報 (設定や計画など) が変更されると、メモリは古くなりますが、保存されている事実は更新または削除されません。
メモリにより書き込みコストが増加します。取得には通常、エージェントが必要なときに情報を検索することが含まれます。メモリには、インタラクション後にどの情報を保存する価値があるかを決定する必要もあります。これにより、モデルの呼び出しと処理時間が増加する可能性があります。この抽出は多くの場合非同期で処理されるため、エージェントの応答が遅くなることはありません。
2 つのソースを慎重にマージする必要があります。検索と記憶により、重複または競合する情報が返される場合があります。エージェントには、各ソースにどれだけの重みを与えるか、および同じコンテキストで両方を使用する方法を決定するための明確なルールが必要です。
検索と記憶の設計作業は、結局のところ、それぞれに何が属するのか、両方をどれだけ積極的に剪定するのか、そして必要のないモデル トークンを渡さずにそれらを 1 つのプロンプトにまとめる方法を決定することになります。
検索とメモリは、長時間実行されるエージェント システムにおけるさまざまな問題を解決します。回収は元をもたらす

ドキュメント、ポリシー、コード、データベース レコードなど、エージェントがその時点で必要とする内部情報。メモリは、決定、好み、ユーザー固有のコンテキストなど、以前の対話からの情報を引き継ぎます。 2 つのシステムには範囲、鮮度に関する懸念、障害モードが異なるため、この区別が重要になります。取得は外部ソースを最新の状態に保つことに依存しますが、メモリは何を保存する価値があるか、また保存された情報がいつ有効でなくなるかを判断することに依存します。
最も効果的なエージェント アーキテクチャでは両方を使用します。これらは、コンテキストに入る内容をフィルタリングし、情報を適度に新鮮に保ち、取得した知識をエージェントが知る必要があるすべての完全な記録として扱うのではなく、関連する記憶とマージします。
したがって、目標は、不必要な情報を伝えることなく、必要なときに必要なコンテキストをエージェントに提供することです。
シェアする
ポスト
シェアする
このトピックの詳細
Agentic AI システムにおけるコンテキスト エンジニアリングとメモリ エンジニアリング
Agentic AI システムのメモリをマスターする 7 つのステップ
ベクトル検索を超えて: 5 つの次世代 RAG 取得戦略
RAG についてパート VI: 効果的な取得の最適化
RAG III について: 融合の取得と再ランキング
短期記憶を超えて: 3 種類の長期記憶…
返信を残す 返信をキャンセルするにはここをクリックしてください。
メールアドレス（公開されません）（必須）
ようこそ！
私はジェイソン・ブラウンリー博士です
また、開発者が機械学習で結果を得るのを支援しています。
続きを読む
電子書籍カタログはここにあります
本当に良いものが見つかります。
Machine Learning Mastery は、人々がテクノロジーを理解できるよう支援することに重点を置いた大手デジタル メディア パブリッシャーである Guiding Tech Media の一部です。当社の使命とチームについて詳しくは、当社の企業 Web サイトをご覧ください。
© 2026 ガイディングテックメディア全著作権所有

## Original Extract

In this article, we break down how retrieval and memory work in long-running AI agents and where each fits. We also look at how to combine them effectively so agents can access relevant information without carrying unnecessary context.

Retrieval vs. Memory in Agentic AI System
Navigation
Making developers awesome at machine learning
Making Developers Awesome at Machine Learning
Code Algorithms Implementing machine learning algorithms from scratch.
Deep Learning (keras) Deep Learning
Neural Net Time Series Deep Learning for Time Series Forecasting
LSTMs Long Short-Term Memory Networks
Making developers awesome at machine learning
Making Developers Awesome at Machine Learning
Code Algorithms Implementing machine learning algorithms from scratch.
Deep Learning (keras) Deep Learning
Neural Net Time Series Deep Learning for Time Series Forecasting
LSTMs Long Short-Term Memory Networks
Retrieval vs. Memory in Agentic AI Systems
Share
Post
Share
In this article, you will learn the conceptual and practical differences between retrieval and memory in agentic AI systems, and how to combine both effectively.
What separates retrieval from memory, and why the distinction matters for long-running agents.
How retrieval pipelines and memory systems are each built, illustrated with a concrete worked example.
How to combine retrieval and memory into a single, effective agent architecture.
An AI agent that can’t remember its previous interactions is not very helpful. Every large language model has a fixed context window , and once a conversation, a set of tool outputs, or a pile of retrieved documents grows past that limit, something has to be dropped, summarized, or fetched fresh. Developers building long-running agents run into this constantly. The agent re-asks questions it already answered, contradicts decisions it made earlier, or fails to recognize that a document it needs even exists.
Retrieval and memory are the two mechanisms that address this, and they solve different halves of the problem. Retrieval pulls in outside knowledge the model was never trained on and should not have to carry by default, such as documentation, code, and database records. Memory persists what the agent itself has learned or done, across a session or across many, so it isn’t starting from zero every time. Confusing the two, or building only one, is where a lot of agent architectures break down. This article covers:
What separates retrieval from memory at a conceptual level
How a retrieval pipeline and a memory system are each built, with a worked example
A side-by-side comparison of the two
How to combine both into a single, effective agentic system
We start with why the split exists in the first place.
Understanding Why Context Forces a Split
A context window is the total set of tokens the model can see at once: system prompt, conversation history, tool outputs, anything inserted ahead of time. It is finite, and every token in it gets attended to on every forward pass, so simply making the window bigger doesn’t scale the way it sounds like it should. Context engineering has emerged as the discipline of curating and managing that limited resource, treating it as the full state available to the model at a given moment, not just a place to stuff instructions.
Given that constraint, an agent has two kinds of information it needs but can’t keep permanently in context:
Information that exists outside the model and outside the current conversation, such as a knowledge base, a codebase, or a set of policy documents. This is what retrieval handles.
Information the agent generated or learned itself, that needs to outlive the current context window, such as a decision made ten turns ago or a fact about a specific user. This is what memory handles.
Both get implemented with similar tools: embeddings, vector search, structured stores. The key difference is what they store and where the information comes from. Retrieval searches a corpus outside the agent, while memory stores information from the agent’s own interactions and past actions.
Defining Retrieval in Agentic Systems
Retrieval is how an agent answers “what does the world know about this that I don’t have in my weights or my current context.” The most common implementation is retrieval-augmented generation , or RAG:
Source documents get chunked into passages small enough to be useful.
Each chunk is converted into an embedding and stored in a vector index.
At query time, the incoming question is embedded the same way, and the index returns the nearest matches.
Those matches get inserted into the prompt alongside the user’s question.
This pattern typically runs on managed datastores with an orchestration layer that ties the retrieval step into the rest of the agent’s reasoning — the approach behind most retrieval-augmented generation architectures in production today. The corpus itself is shared — every user asking about the same product documentation hits the same index — and it is refreshed on its own schedule, independent of any individual conversation.
Defining Memory in Agentic Systems
Memory is how an agent answers “what have I already learned or done that I need to carry forward.” It splits into two layers that behave differently:
Short-term memory is the running session state: the conversation so far, plus anything the agent has written to a scratchpad during the current task. It’s cheap, and it disappears when the session ends.
Long-term memory persists across sessions. It has to answer a harder question than retrieval does: not just “what’s relevant,” but “what’s worth keeping in the first place.”
Some agent memory systems automatically extract useful facts, preferences, and context from conversations and store them for later use. At the start of a new session, the agent can query that memory much like it would query a retrieval index, but the results are specific to a user, task, or agent rather than a shared document corpus. When designing this layer, teams can explore different agent memory strategies and agent memory frameworks depending on what they need to store and retrieve.
A quick worked example makes the split concrete. A customer messages a support agent about a delayed order.
For a delayed order, the agent first checks its memory for the customer’s previous history. It finds a note from three weeks ago saying they prefer email follow-up and that a similar shipping issue was resolved with a partial refund. That is memory, because it comes from the agent’s record of this specific customer.
The agent then needs the current shipping policy, which changed last month, so it searches the company’s documentation and retrieves the relevant section. That is retrieval, because the information comes from an external source and applies to all customers. Both results are added to the same prompt, but they answer different questions.
Comparing Retrieval and Memory
Laid out side by side, the differences between retrieval and memory are easier to see at a glance:
The failure modes listed in the table above explain why an agent built with only one of the two tends to break in predictable ways, and why most working systems end up needing both.
Combining Retrieval and Memory into an Effective System
An agent with retrieval but no memory re-derives the same conclusions every session and can’t personalize anything. An agent with memory but no retrieval knows its own history but has no way to ground itself in anything outside that history; it can’t answer questions about a policy that changed after its training data ended. Getting the combination right comes down to a few things:
Filtering matters more than window size. Adding more retrieved documents or memory entries does not necessarily improve answers. Beyond a point, extra context can make answers worse because the model has to process and weigh every additional token. Small, targeted searches are often more effective than one broad search and can keep retrieval token-efficient.
Staleness works differently for retrieval and memory. A retrieval index becomes stale when the underlying documents change without being re-indexed. Memory becomes stale when information about a user changes — such as a preference or plan — but the stored fact is not updated or removed.
Memory adds a write cost. Retrieval usually involves looking up information when the agent needs it. Memory also requires deciding what information is worth saving after an interaction, which can add model calls and processing time. This extraction is often handled asynchronously so it does not slow down the agent’s response.
The two sources need to be merged carefully. Retrieval and memory can return information that overlaps or conflicts. The agent needs clear rules for deciding how much weight to give each source and how to use both in the same context.
The design work for retrieval and memory comes down to deciding what belongs in each, how aggressively to prune both, and how they come together into a single prompt without handing the model tokens it doesn’t need.
Retrieval and memory solve different problems in long-running agent systems. Retrieval brings in external information the agent needs at the moment, such as documentation, policies, code, or database records. Memory carries forward information from previous interactions, such as decisions, preferences, and user-specific context. The distinction matters because the two systems have different scopes, freshness concerns, and failure modes. Retrieval depends on keeping external sources up to date, while memory depends on deciding what is worth storing and when stored information is no longer valid.
The most effective agent architectures use both. They filter what enters the context, keep information reasonably fresh, and merge retrieved knowledge with relevant memory instead of treating either as a complete record of everything the agent needs to know.
The goal, therefore, is to give the agent the context it needs, when it needs it, without carrying unnecessary information.
Share
Post
Share
More On This Topic
Context vs. Memory Engineering in Agentic AI Systems
7 Steps to Mastering Memory in Agentic AI Systems
Beyond Vector Search: 5 Next-Gen RAG Retrieval Strategies
Understanding RAG Part VI: Effective Retrieval Optimization
Understanding RAG III: Fusion Retrieval and Reranking
Beyond Short-term Memory: The 3 Types of Long-term…
Leave a Reply Click here to cancel reply.
Email (will not be published) (required)
Welcome!
I'm Jason Brownlee PhD
and I help developers get results with machine learning .
Read more
The EBook Catalog is where
you'll find the Really Good stuff.
Machine Learning Mastery is part of Guiding Tech Media, a leading digital media publisher focused on helping people figure out technology. Visit our corporate website to learn more about our mission and team.
© 2026 Guiding Tech Media All Rights Reserved
