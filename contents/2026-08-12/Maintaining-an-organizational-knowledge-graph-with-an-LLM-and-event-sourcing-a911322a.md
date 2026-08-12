---
source: "https://blog.arkency.com/maintaining-an-organizational-knowledge-graph-with-an-llm-and-event-sourcing/"
hn_url: "https://news.ycombinator.com/item?id=49275420"
title: "Maintaining an organizational knowledge graph with an LLM and event sourcing"
article_title: "Maintaining an organizational knowledge graph with an LLM and event sourcing | Arkency Blog"
author: "pdabrowski6"
captured_at: "2026-08-12T17:52:39Z"
capture_tool: "hn-digest"
hn_id: 49275420
score: 2
comments: 0
posted_at: "2026-08-12T16:56:43Z"
tags:
  - hacker-news
  - translated
---

# Maintaining an organizational knowledge graph with an LLM and event sourcing

- HN: [49275420](https://news.ycombinator.com/item?id=49275420)
- Source: [blog.arkency.com](https://blog.arkency.com/maintaining-an-organizational-knowledge-graph-with-an-llm-and-event-sourcing/)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T16:56:43Z

## Translation

タイトル: LLM とイベント ソーシングを使用した組織ナレッジ グラフの維持
記事のタイトル: LLM とイベント ソーシングを使用した組織ナレッジ グラフの維持 |アーケンシーのブログ

記事本文:
ナビゲーションを切り替え
すべての記事
LLM とイベント ソーシングによる組織のナレッジ グラフの維持
…そして、5600 人以上の Rails エンジニアがこれも読んでいる理由を確認してください
LLM とイベント ソーシングによる組織のナレッジ グラフの維持
組織は驚くほど忘れることが得意です。
意思決定は電話で行われ、洞察は Slack スレッドに埋もれ、1 か月後には、なぜ現状がそうなったのかを誰も思い出せなくなります。
アーケンシーでも、時々何かが私たちから抜け落ちていくような気がしました。
毎週の通話、臨時ミーティング、読書会、Slack でのディスカッション、GitHub でのメンション、電子メールの受信箱など、これらすべてのシグナルを整理する際に何らかのサポートを利用することができます。
そして3月にRuby Community Conference 2026が開催されました。
クラクフでは、オビエ フェルナンデス氏が NEXUS システムの一部を披露しました。
彼はすでに 1 月にブログでこのことについて説明していましたが、私が初めてこのことに出会ったのはカンファレンスでした。
それが私が独自のソフトウェアを構築し始めるのに必要な後押しでした。
それがすでに形になりつつあったとき、Andrej Karpathy は LLM Wiki ノートを公開しました。
RAG システムがクエリごとにドキュメントを再検出する代わりに、LLM は永続的な Wiki (相互リンクされたマークダウン ページ、その下にある不変のソース、および人間によるループのキュレーション) を段階的に維持します。
業界で最もホットなトピックの 1 つになったばかりのことに自分が取り組んでいることに気づき、非常に興奮しました。
最終的には、Rails Event Store 上に構築された、クローズド オントロジーを備えたマルチテナント ナレッジ グラフである Planet Arkency にたどり着きました。
この投稿では、私が行った設計上の決定について説明したいと思います。
LLM が実際に威力を発揮するのは非構造化入力です
構造化データの場合、20 年ほど前にそのようなシステムを構築できたはずです。
Webhook、フォーム、統合 - グラフへの構造化入力の解析は解決された問題です。
でも、

最も興味深い知識は、会議記録、Slack ディスカッション、電子メール、またはまだ誰も構築していない統合から得られるあらゆるパーサーが処理できない入力に存在します。
ここで、LLM が私たちの状況を変えました。
すべては単一の取り込みエンドポイントを介してシステムに流れ込みます。
トランスクリプト、誰かが専用の絵文字リアクションでフラグを立てた Slack スレッド、ブリッジの受信箱に届くメール、RSS フィード、カレンダーの招待状、個人的なメモ。
統合ポイントのコードも書きません。
Zapier や n8n などのツールはソースを監視し、コンテンツをその単一のエンドポイントにプッシュします。
取り込まれたすべてのコンテンツは、システムの中心である抽出を経ます。
LLM はコンテンツを読み取り、それが私たちの知識にとって何を意味するのか、つまりどのエンティティがその中に登場するのか、それらについて何を学んだのか、そしてそれらが互いにどのように関係しているのかを判断します。
この投稿の大部分は、その 1 つのステップの周囲で何が起こっているかについて説明しています。
私たちの会話では、人、プロジェクト、クライアント、ツール、意思決定など、同じ名前が何度も​​出てきます。
週ごとに変化するのは、それらについて私たちが知っていることと、それらが互いにどのように関係しているかということです。
これは、属性を持つエンティティが型指定された関係によって接続されたグラフに自然にマッピングされます。
誰が何に取り組んでいるのか。
誰が、いつ、どのような決定を下したのか。
どのプロジェクトはどのツールによって異なります。
これが LLM Wiki のアプローチと最も異なる点です。
Wiki では、誰かが何らかのプロジェクトに取り組んでいるという事実が、ページ上の 1 つの文に記録され、せいぜい 2 つのページ間のリンクが付けられます。
知識はそこにありますが、それを活用できるのは読者だけです。
型付きグラフでは、person --works_on--> project はデータの一部です。クエリ、トラバース、カウントが可能です。
グラフ自体は PostgreSQL 上にあります。ノード テーブル、一意の (ソース、ターゲット、リレーション) トリプルを持つエッジ テーブル、両方の jsonb 属性です。
N

ロケット科学はここにあります。
専用のグラフ データベース (Neo4j、NEXUS が使用しているようなトリプル ストア) は、ディープ マルチホップ トラバーサルなどの一部の特定のワークロードにより適している可能性があります。
ただし、これはストレージの詳細であり、データ層に別のアダプターを作成することで後で変更できる種類のものです。
どの種類のノードと関係が存在できるかは、オントロジーで定義され、プレーンな YAML ファイルに保存されます。
# config/ontology.yml から
ノードの種類:
- 種類：人
説明 : 「チームメンバー、候補者、クライアント担当者、社外の人」
- 種類：決定
説明 : 「正式な決定にはグループの評決が必要です。カジュアルな提案にはアイデアを使用してください」
エッジ関係:
- 関連: works_on
署名: "人 --works_on--> プロジェクト"
オントロジーは閉じられています。種類または関係がリストにない場合、モデルはそれを使用できません。
当初、私は LLM が独自の型を導入できるオープン オントロジーについて考えていました。
驚くほど早く、グラフに完全な混乱がもたらされました。
私の意見では、モデルに何を探すべきかを事前に伝える方が良いと思います。
「組織ナレッジ グラフ」は、あらゆる異なる目的に対応する 1 つの普遍的なグラフを提案します。
私たちはそんなことを信じていませんが、DDD 実践者ならその理由を理解するでしょう。
私たちはマルチテナント アーキテクチャを使用して、独自のオントロジーを持つ個別のグラフを維持します。これは実際には独自のユビキタス言語を意味します。
当社の内部 Arkency グラフは、CRM に非常に近い領域である、人、プロジェクト、意思決定について語っています。
Rails Event Store のメンテナーとして私たちが実行しているグラフは、リリース、既知の問題、コミュニティ コンテンツで説明されています。
異なるドメイン、異なる語彙、その下にある同じ機構。
有界コンテキストの境界は、あるグラフがどこで終わり、別のグラフが始まるかを示します。
抽出から何が得られるか
オントロジーは抽出プロンプトに markdo としてレンダリングされます。

wn テーブルを作成し、列挙型としてスキーマに追加します。
(app/lib/prompts/extraction.md.erb より)
あなたは <%= Tenancy の組織ナレッジ アナリストです。 current_tenant 。名前 %> 。私たちは内部ナレッジグラフを構築しています。
提供されたコンテンツ (ノードとエッジ) からナレッジ グラフを抽出します。グラフでは、提供されたコンテンツを完全に再構築できる必要があります。
## ノード
各ノードには、名前、種類、short_description、説明、属性 (オプションのキーと値のペア) があります。
許可される種類:
|種類 |それが何を表しているのか |典型的な属性 |
|---|---|---|
<% オントロジー。ノードの種類 。それぞれが行うこと | k | -%>
| <%= k 。 fetch ( "種類" ) %> | <%= k 。 fetch ( "説明" ) %> | <%= k 。フェッチ ( "attrs" , [])。次に { |属性 |属性 。空の？ ? "—" : 属性。マップ { | | 。は? (ハッシュ) ? ( a [ "values" ] ? " #{ a [ "name" ] } ( #{ a [ "values" ].join ( ", " ) } )" : a [ "name" ]) : a }.結合 ( ", " ) } %> |
<% 終了 -%>
## エッジ
各エッジには、ソース、ターゲット、関係、コンテキスト、属性 (オプションのキーと値のペア) があります。
許可される関係:
|関係 |ソースの種類 |対象の種類 |ヒント |属性 |
|---|---|---|---|---|
<% オントロジー。エッジ関係 。それぞれが行うこと | r | -%>
<%
sig = オントロジー 。 parse_signature ( r . fetch ( "署名" ))
source_kind = sig [:source ]。結合 ( " / " )
target_kind = sig [:target ]。結合 ( " / " )
ヒント = r [ "ヒント" ] || 「――」
attrs = r .フェッチ ( "attrs" , [])。次に { | | 。空の？ ? 「―」：a 。マップ { | |でで 。は? (ハッシュ) ? ( at [ "values" ] ? " #{ at [ "name" ] } ( #{ at [ "values" ].join ( ", " ) } )" : at [ "name" ]) : at }.結合 ( ", " ) }
-%>
| <%= r 。 fetch ( "関係" ) %> | <%= ソースの種類 %> | <%= ターゲットの種類 %> | <%= ヒント %> | <%= 属性 %> |
<% 終了 -%>
...
各抽出は、モデルが 1 つの構造化された結果 (コンテンツ内で見つかったエンティティ、エンティティ間の関係) を返すことで終了します。

それらを確認し、それらを反映するために既存のグラフをどのように変更する必要があるか。
そのために RubyLLM のスキーマ サポートを使用します。
# app/lib/extraction_result_schema.rb から
配列 :nodes 、説明: 「作成または更新するエンティティ。各名前は一意である必要があります。重複したノードは禁止です。」する
オブジェクトが行う
string :status 、 enum: [ "new" 、 "existing" ]、説明: 「ノードが search_nodes/list_nodes_by_kind/get_node_edges によって返され、それを再利用している場合は 'existing'。導入している場合は 'new'。システムは正規名を検証し、不一致がある場合は中止します。」
string :name 、説明: "エンティティ名。「既存の」ノードの場合は、ツール呼び出し結果からの正確な正規名を使用します。「新しい」ノードの場合は、導入する正規名を使用します。
string :new_name 、必須: false 、説明: "オプション。コンテンツがより明示的な標準形式 (例: 頭字語 → 完全な用語、指称 → フルネーム) を明らかにする場合、'既存の' ノードにのみ設定します。ノードは `name` によって検索され、`new_name` に名前変更されます。
文字列 :kind 、説明: "次のいずれかである必要があります: #{ kind_names } "
string :short_description 、説明: 「このエンティティが何であるかの安定した合成 (検索用)。エピソード固有ではなく、一般的でアイデンティティに焦点を当てています。最大 15 単語。」
string :description , description: "新しいノードの場合: 内容に基づいた簡単な説明。既存のノードの場合: 以前の説明と新しい情報を合成します。わかりやすくするために書き直すのは問題ありませんが、以前の事実は保持してください。"
配列 :attrs 、説明: 「キーと値の属性。コンテンツからわかっているもののみを含めます。」 ...を終了します
array :aliases 、必須: false 、説明: "オプション。コンテンツ内でこのエンティティが参照された代替表面形式 (指称、頭字語、完全形式と短縮形式)、または — `new_name` 経由で名前を変更する場合 — v のままの場合は古い標準

固体の表面形状。既存のノードにまだ存在しない新しいエイリアスのみを含めます。エイリアスは、別の名前の下にある同じエンティティであり、決して別個のエンティティではありません。" do ... end
終わり
終わり
配列 :edges 、説明:「すべての関係。徹底して正確に行うこと。」
オブジェクトが行う
文字列 :source 、説明: "ソース ノード名 (正確な
[切り捨てられた]
# app/handlers/propose_graph_change.rb から
ノード = ノード . find_or_initialize_by (名前: データ [:名前])
enforce_status! ( data [ :name ], data [ :status ], node ) # モデルの新規/既存の主張が DB と一致しない場合に発生します
was_new = ノード 。新しいレコード?
ノード。 assign_attributes ( short_description: ... 、 description: ... 、 attrs: node . attrs . merge ( attrs ))
変更 = ノード 。変化します。 ( "updated_at" 、 "created_at" 、 "kind" 、 "slug" ) を除く
{ 操作: was_new ? "create" : "update" 、node_id: ノード 。持続しましたか？ ?ノード。 id : nil 、変更: 変更 }
node.changes は、{field => [before, after]} ペアを無料で提供し、この前後のスナップショットがグラフ変更提案のワイヤー形式になります。
エッジはまったく同じ処理を受けます。(ソース、ターゲット、リレーション) トリプルによって検索され、ダーティ トラッキングで差分されます。
また、LLM の主張を盲目的に信頼することもありません。
各ノードを new または既存のノードとして宣言する必要があり、バリデーターがそれをデータベースと照合チェックします。
不一致の場合、LLM は自然言語フィードバックを取得し、同じ会話に対する別の試行を行います。
ID の解決は難しい部分です
モデルは各ノードを new または既存として宣言する必要があると書きました。
しかし、どうやってそれを知ることができるのでしょうか？
グラフ全体を LLM コンテキストにロードしますか?
いいえ、ここでツール呼び出しが必要になります。
そして、それは単純な検索よりも困難です。
「ピョートレック」、「ピョートル・ジュレヴィチ」、そして私の名前をZoomが書き写したものはすべて同一人物です。
あなたがそうするなら

サーフェス フォームごとにノードを消費すると、グラフは 1 週間以内にゴミになってしまいます。
まず、モデルは書き込む前に参照する必要があります。
抽出中は、 search_nodes や get_node_edges などの読み取り専用ツールにアクセスできます。
抽出プロンプトはそれについて明示的に示しています。
(app/lib/prompts/extraction.md.erb より)
- ノードを作成する前に、search_nodes を使用してノードがすでに存在するかどうかを確認します。 (...)
- search_nodes が結果を返さない場合、ノードはまだ存在しません。ノードの作成に進みます。 (...)
- search_nodes があいまいな結果を返した場合、または抽出を決定するためにより広範なコンテキストが必要な場合は、get_node_edges を使用してノードの接続を検査します。
- search_nodes でノードを見つけた後、それらの接続方法を決定する前に、get_node_edges を使用して既存の関係を確認します。
2 番目に、エイリアスは ID メカニズムです。
各ノードには 1 つの正規名と任意の数のエイリアスがあります。
スキーマは、エイリアスが別の名前の下で同じエンティティであること、決して別個のエンティティではないことをモデルに指示します。
コンテンツがより適切な正規形式を明らかにすると、モデルは new_name を設定し、古い名前はエイリアスとして残るため、今後のあいまい検索でも解決されます。
第三に、検索はハイブリッドです。
ノード名とエイリアスのトライグラム類似性 (GIN インデックスを備えた pg_trgm) により、スペルミスが検出されます。
埋め込む

[切り捨てられた]

## Original Extract

Toggle navigation
All Articles
Maintaining an organizational knowledge graph with an LLM and event sourcing
… and check why 5600+ Rails engineers read also this
Maintaining an organizational knowledge graph with an LLM and event sourcing
Organizations are surprisingly good at forgetting.
Decisions are made on calls, insights get buried in Slack threads, and a month later no one remembers why things are the way they are.
At Arkency, I had a feeling that some things slip away from us too from time to time.
Weekly calls, ad-hoc meetings, our book clubs, Slack discussions, GitHub mentions, e-mail inbox - we could use some support in organizing all those signals.
Then Ruby Community Conference 2026 happened in March.
In Kraków, Obie Fernandez showed some parts of his NEXUS system.
He had already described it on his blog back in January, but the conference was where I first came across it.
That was the push I needed to start building our own software.
When it was already taking shape, Andrej Karpathy published his LLM Wiki note.
Instead of a RAG system rediscovering your documents on every query, an LLM incrementally maintains a persistent wiki: interlinked markdown pages, immutable sources underneath, and a human curating the loop.
It was quite exciting to realize I was working on something that had just become one of the hottest topics in the industry.
We ended up with Planet Arkency - a multi-tenant knowledge graph with a closed ontology , built on Rails Event Store.
In this post, I want to walk you through the design decisions I made.
Unstructured input is where LLMs actually shine
For structured data, you could have built such a system like twenty years ago.
Webhooks, forms, integrations - parsing structured input into a graph is a solved problem.
But the most interesting knowledge lives in the input no parser could ever handle: meeting transcripts, Slack discussions, emails, or anything coming from an integration nobody has built yet.
This is where LLMs changed the game for us.
Everything flows into the system through a single ingestion endpoint.
Transcripts, Slack threads someone flagged with a dedicated emoji reaction, emails arriving at a bridge inbox, RSS feeds, calendar invites, personal notes.
We don’t even write code for the integration points.
Tools like Zapier or n8n watch the sources and push the content to that single endpoint.
Every ingested piece of content then goes through an extraction - the heart of the system.
An LLM reads the content and works out what it means for our knowledge: which entities appear in it, what we learned about them, and how they relate to each other.
Most of this post is about what happens around that single step.
The same names keep coming back in our conversations: people, projects, clients, tools, decisions.
What changes from week to week is what we know about them and how they relate to each other.
That maps naturally to a graph: entities with attributes, connected by typed relations.
Who works on what.
Who made which decision, and when.
Which project depends on which tool.
This is where we differ most from the LLM Wiki approach.
In a wiki, the fact that someone works on some project is written down in a sentence on a page, at best with a link between the two pages.
The knowledge is there, but only a reader can make use of it.
In a typed graph, person --works_on--> project is a piece of data: you can query it, traverse it, count it.
The graph itself sits on PostgreSQL: a nodes table, an edges table with a unique (source, target, relation) triple, jsonb attributes on both.
No rocket science here.
Dedicated graph databases (Neo4j, triple stores like the one NEXUS uses) could be a better fit for some specific workloads, like deep multi-hop traversal.
But that is a storage detail - the kind you could change later by writing another adapter for the data layer.
Which kinds of nodes and relations may exist is defined in an ontology , stored in a plain YAML file:
# from config/ontology.yml
node_kinds :
- kind : person
description : " team member, candidate, client contact, external person"
- kind : decision
description : " formal decision requiring group verdict — for casual suggestions use idea"
edge_relations :
- relation : works_on
signature : " person --works_on--> project"
The ontology is closed - if a kind or relation is not on the list, the model cannot use it.
Initially I was thinking about an open ontology, where the LLM could introduce its own types.
It brought complete chaos into the graph surprisingly fast.
In my opinion, it is better to tell the model upfront what to look for.
“The organizational knowledge graph” suggests one universal graph for all different purposes.
We don’t believe in that, and DDD practitioners will recognize why.
We use multi-tenant architecture to maintain separate graphs with their own ontologies, which really means their own ubiquitous languages.
Our internal Arkency graph speaks in people, projects and decisions - a domain quite close to a CRM.
The graph we run as Rails Event Store maintainers speaks in releases, known problems and community content:
Different domains, different vocabularies, the same machinery underneath.
The boundaries of a bounded context tell you where one graph ends and another begins.
What comes out of an extraction
The ontology is rendered into the extraction prompt as markdown tables and into the schema as enums.
(from app/lib/prompts/extraction.md.erb)
You are an organizational knowledge analyst for <%= Tenancy . current_tenant . name %> . We are building an internal knowledge graph.
Extract a knowledge graph from the provided content: nodes and edges. The graph should allow full reconstruction of the provided content.
## Nodes
Each node has: name, kind, short_description, description, attrs (optional key-value pairs).
Allowed kinds:
| kind | what it represents | typical attrs |
|---|---|---|
<% ontology . node_kinds . each do | k | -%>
| <%= k . fetch ( "kind" ) %> | <%= k . fetch ( "description" ) %> | <%= k . fetch ( "attrs" , []). then { | attrs | attrs . empty? ? "—" : attrs . map { | a | a . is_a? ( Hash ) ? ( a [ "values" ] ? " #{ a [ "name" ] } ( #{ a [ "values" ]. join ( ", " ) } )" : a [ "name" ]) : a }. join ( ", " ) } %> |
<% end -%>
## Edges
Each edge has: source, target, relation, context, attrs (optional key-value pairs).
Allowed relations:
| relation | source kind | target kind | hint | attrs |
|---|---|---|---|---|
<% ontology . edge_relations . each do | r | -%>
<%
sig = Ontology . parse_signature ( r . fetch ( "signature" ))
source_kind = sig [ :source ]. join ( " / " )
target_kind = sig [ :target ]. join ( " / " )
hint = r [ "hint" ] || "—"
attrs = r . fetch ( "attrs" , []). then { | a | a . empty? ? "—" : a . map { | at | at . is_a? ( Hash ) ? ( at [ "values" ] ? " #{ at [ "name" ] } ( #{ at [ "values" ]. join ( ", " ) } )" : at [ "name" ]) : at }. join ( ", " ) }
-%>
| <%= r . fetch ( "relation" ) %> | <%= source_kind %> | <%= target_kind %> | <%= hint %> | <%= attrs %> |
<% end -%>
...
Each extraction ends with the model returning one structured result : the entities it found in the content, the relations between them, and how the existing graph should change to reflect them.
We use RubyLLM’s schema support for that.
# from app/lib/extraction_result_schema.rb
array :nodes , description: "Entities to create or update. Each name must be unique — no duplicate nodes." do
object do
string :status , enum: [ "new" , "existing" ], description: "'existing' iff the node was returned by search_nodes/list_nodes_by_kind/get_node_edges and you are reusing it. 'new' if you are introducing it. The system verifies the canonical name and aborts on mismatch."
string :name , description: "Entity name. For 'existing' nodes use the EXACT canonical name from the tool call result. For 'new' nodes the canonical name you are introducing."
string :new_name , required: false , description: "Optional. Set ONLY for 'existing' nodes when the content reveals a more explicit canonical form (e.g. acronym → full term, diminutive → full name). The node is looked up by `name` and renamed to `new_name`."
string :kind , description: "Must be one of: #{ kind_names } "
string :short_description , description: "Stable synthesis of what this entity is (for search). General and identity-focused, not episode-specific. Max 15 words."
string :description , description: "For new nodes: brief description based on the content. For existing nodes: synthesize prior description with new information. Rewriting for clarity is fine, but preserve prior facts."
array :attrs , description: "Key-value attributes. Only include what is known from the content." do ... end
array :aliases , required: false , description: "Optional. Alternative surface forms (diminutives, acronyms, full vs short forms) under which this entity was referred to in the content, or — when renaming via `new_name` — the old canonical if it remains a valid surface form. Only include NEW aliases not already present on the existing node. An alias is the SAME entity under another name — never a separate entity." do ... end
end
end
array :edges , description: "ALL relationships. Be thorough and precise." do
object do
string :source , description: "Source node name (exact
[truncated]
# from app/handlers/propose_graph_change.rb
node = Node . find_or_initialize_by ( name: data [ :name ])
enforce_status! ( data [ :name ], data [ :status ], node ) # raises when the model's new/existing claim disagrees with the DB
was_new = node . new_record?
node . assign_attributes ( short_description: ... , description: ... , attrs: node . attrs . merge ( attrs ))
changes = node . changes . except ( "updated_at" , "created_at" , "kind" , "slug" )
{ op: was_new ? "create" : "update" , node_id: node . persisted? ? node . id : nil , changes: changes }
node.changes gives us {field => [before, after]} pairs for free, and this before/after snapshot becomes the wire format of the graph change proposal.
Edges get exactly the same treatment - looked up by their (source, target, relation) triple and diffed with dirty tracking.
We also don’t blindly trust what the LLM claims.
It has to declare each node as new or existing , and a validator cross-checks it against the database.
On mismatch, the LLM gets natural-language feedback and another attempt on the same conversation.
Identity resolution is the hard part
I just wrote that the model has to declare each node as new or existing .
But how would it know?
Do we load the whole graph into LLM context?
No - this is where tool calls come in.
And it is harder than a simple lookup.
“Piotrek”, “Piotr Jurewicz” and whatever Zoom’s transcription makes out of my name are the same person.
If you create a node per surface form, your graph turns into garbage within a week.
First, the model must look before it writes.
During extraction it has access to read-only tools like search_nodes or get_node_edges .
The extraction prompt is explicit about it:
(from app/lib/prompts/extraction.md.erb)
- Before creating any node, use search_nodes to check if it already exists. (...)
- If search_nodes returns no results, the node does not exist yet — proceed to create it. (...)
- If search_nodes returns ambiguous results, or you need broader context to make extraction decisions, use get_node_edges to inspect the node's connections.
- After finding nodes with search_nodes, use get_node_edges to see their existing relationships before deciding how to connect them.
Second, aliases are the identity mechanism.
Each node has one canonical name and any number of aliases.
The schema instructs the model that an alias is the same entity under another name - never a separate entity.
When the content reveals a better canonical form, the model sets new_name and the old name stays as an alias, so future fuzzy searches still resolve it.
Third, the search is hybrid.
Trigram similarity (pg_trgm with GIN indexes) over node names and aliases catches misspellings.
Embedd

[truncated]
