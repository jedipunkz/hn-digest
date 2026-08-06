---
source: "https://udnes.dev/posts/context-engineering-harness-part-1-ontology/"
hn_url: "https://news.ycombinator.com/item?id=49193595"
title: "Context Engineering in an LLM Harness"
article_title: "Context Engineering in an LLM Harness — Part 1: Ontology — Henrik Udnes"
author: "udnes"
captured_at: "2026-08-06T07:34:25Z"
capture_tool: "hn-digest"
hn_id: 49193595
score: 2
comments: 0
posted_at: "2026-08-06T07:25:19Z"
tags:
  - hacker-news
  - translated
---

# Context Engineering in an LLM Harness

- HN: [49193595](https://news.ycombinator.com/item?id=49193595)
- Source: [udnes.dev](https://udnes.dev/posts/context-engineering-harness-part-1-ontology/)
- Score: 2
- Comments: 0
- Posted: 2026-08-06T07:25:19Z

## Translation

タイトル: LLM ハーネスにおけるコンテキスト エンジニアリング
記事のタイトル: LLM ハーネスにおけるコンテキスト エンジニアリング — パート 1: オントロジー — Henrik Udnes
説明: システム アーキテクトおよびソフトウェア エンジニアが、興味深い問題や苦労して得た教訓について書いています。

記事本文:
Henrik Udnes システムアーキテクト兼ソフトウェアエンジニア
LLM ハーネスにおけるコンテキスト エンジニアリング — パート 1: オントロジー
LLM ハーネスのコンテキスト エンジニアリング · 5 つの部分
LLM エージェントを問題トラッカーに埋め込むことを想像してください。ユーザーがアシスタントを開く
リリース プロジェクト内で次のように尋ねます。「どの未解決の問題がプロジェクトをブロックしているか」
リリース？」モデルには製品固有の知識が必要です。つまり、プロジェクト、問題、
ワークフローと依存関係は互いに適合します。このチームが使用するステータス名。
そして、内部 ID がユーザーに表示される名前にどのようにマッピングされるか。生の ID をエコーバックする
user はまだバグであり、答えではありません。
明らかな方法は、これらすべてをシステム プロンプトに貼り付けることです。みんなの
私たちのものも含め、最初のハーネスはこれを行います。機能しますが、コストが増加するにつれて増加します
ドメイン: システム プロンプトは毎ターンの反復ごとに再レンダリングされます —
私たちのハーネスでは文字通り、buildSystemPrompt は LLM 呼び出しごとに 1 回実行されます。
既存のコンテキストのすべての文字は、その生涯にわたって支払う税金です。
会話。ドメインの知識は増えるばかりです。十分な規模で、
ペーストするとすべてが高価になり、通常の避難用ハッチは破壊的になります
状況が厳しくなったときに履歴を切り捨てる — コンテキストを正確に破棄し、
エージェントが最も必要とするもの。
このシリーズは、当社のハーネスが取り組む 3 番目のオプションに関するものです。
どこでも: 延期し、破棄しないでください。インデックスを常駐コンテキスト内に保持する
そしてモデルがオンデマンドでプルするものをペイロードにします。引っ張られたときでも
結果が大きすぎる場合は、ビューを制限し、正確な値をポインタの後ろに保持してください。
5 つの異なる高度で 5 回適用されました。
オントロジー (この投稿) — ドメインとは何か
ツールの検出 - エージェントができること
参照による値 - そこを流れるデータ
記憶 - エージェントが覚えていること
コストは議論の半分にすぎません。コンテキストウィンドウ

も注目の予算です。
モデルの前にあるものが多ければ多いほど、モデルの一部の使用方法が悪くなります。
無関係なスキーマと参照テキストはトークンを消費するだけでなく、競合します。
タスクと一緒に。延期が機能するのは、それが省略ではないためです。
安定したアドレスですべてにアクセスできるため、順位が縮小します
コンテキストは、ターンから情報を削除せずにノイズを削除します。
システム。必要なものがない限り、コンテキストが少ないほど正確性が高くなります
到達不能になります — そしてモデルが認識できるようにインデックスを十分に豊富に保ちます
何を引っ張るか、それがこのシリーズの機構が気を付けていることです。
ここでのメカニズムと測定値は、生産ハーネスから得られます。の
実行中のサンプルでは汎用の問題トラッカーを使用しているため、アーキテクチャが読みやすくなっています。
商品知識が無くても。トークンの推定値は 1 つの大まかな変換を使用します。
4 文字ごとのトークン。
explore_ontology ツール サーフェスの名前空間行 + スキャフォールディング + 学習済み/アクティブ セット ツール カタログ 169 ツール · スキーマの ≈85,000 文字 search_tools activate_tool スキル カタログ ID — タイトル: スキルごとの説明 スキル本体 アクティブ時の完全な命令 activate_skill deactivate_skill メモリ ダイジェストでキャップされた {key, description} のインデックス メモリ バンクでキャップされたレコード、キーでアドレス指定された recall_memories ツール結果の限定プレビュー + $ref スタブBLOB ストアの境界のないバイト、サーバー側の $ref 解決 シリーズ全体の形状: 常駐コンテキストは 5 つの小さなインデックスを保持します。すべてのペイロードはプルの背後に存在します。この投稿は最初の行です。
文書ではなくグラフ
オントロジーは散文ではありません。これは型付きグラフです: 15 ドメイン (境界あり)
課題、ワークフロー、ラベルなどのコンテキスト)、28 のエンティティ、それぞれ
正確に 1 つのドメインによって所有され、次のように派生した ID を持つ 12 の関係
<from>.<predicate>.<to> :
エクスポート定数 ISSUE_ASSIGNE

D_TO_USER = 新しい関係 (
問題、
" 割り当て先 " 、
ユーザー、
" 問題は、その問題を担当するユーザーに割り当てられます。"
);
プロジェクト 課題 ラベル ワークフロー ユーザー 添付ファイル 所属先 ラベル付き 割り当て先 定義済み プロジェクト課題ラベル ステータス ワークフロー ユーザー 添付ファイル ステータス · 注 (ドリルインのみ) 「ユーザーはステータスを「ステージ」または「列」と呼ぶ場合があります。」問題トラッカーにマッピングされた生産グラフ パターン。エンティティは、独自のドメイン内に存在します。関係はドメインにまたがる場合があります。メモには、何かが掘り下げられた場合にのみ伝達される運用上の知識が含まれます。
ここでは 2 つの設計ルールが多くの地味な作業を行います。
まず、グラフへの参照は文字列型ではなく検証されます。
自分自身にコンセプトのタグを付けるもの - を宣言するツール
semantics = [ISSUE] 、エンティティをスコープとするメモリ — に対してチェックされます
オントロジーに問題があるため、不正な ID はオーサリング時に黙って一致するのではなく失敗します。
実行時には何もありません。依存関係は厳密に一方向です。ツールは、
オントロジー;オントロジーではツールについて聞いたことがありません。そしてツールのドメイン
メンバーシップはタグから派生し、作成されることはないため、流出することはありません
同期の。
第二に、知識の寿命は一度だけです。エンティティのメモ — 運用上の注意事項:
エンティティの正式な名前の代わりにユーザーが実際に言う言葉、
何かを発表する前に実行され、何も言わずに失敗する編集はデータです。
エンティティがレンダリングされる場所にレンダリングされます。インターフェイスのドキュメントは次のとおりです。
メモは「プロンプト文字列に決して手書きではない」ことを明示します。あります
間違った事実を修正するのにちょうど 1 つの場所です。
実際に毎ターン出荷されるもの
Ontology.toPrompt() からのオントロジーの常駐フットプリント全体を次に示します。
public toPrompt (): string {
const ルート = この .domainList
。 filter (( d ) => d.getParent () ===

null );
if (roots.length === 0 ) は "" を返します。
const 行: string [] = [ " このスコープ内のドメイン: " ];
for (ルートの定数ドメイン)
線。 Push ( `- ${ ドメイン.id } — ${ ドメイン.説明 } ` );
リターンライン。結合 ( " \n " );
}
ルート ドメインのみ、各 1 行。コード内のコメントは意図をより適切に示しています
私ができることよりも、意図的に簡潔に — 立ったままのオリエンテーション (有効な ID と場所)
物事は生きています)、知識のダンプではありません。レンダリング前、forScope フィルター
現在のスコープが何を意味するかまでグラフに表示: できないドメイン
エージェントが立っている場所に適用すると、ユーザーが何であろうともドロップされます
は現在、エージェントに渡されたスコープ参照を参照しており、そのスコープ参照をプルします。
一般的なスコープには 13 行が残ります。
1,356 文字 ≈ 340 トーク フルグラフ ダンプ 28 エンティティ + ノート + 12 のリレーションシップ ≈ 7.5,000 文字 ≈ 1.9,000 トーク オントロジー ソースから測定: 常駐セクションと完全なグラフ (すべてのエンティティの説明、メモ、関係) をプロンプトにレンダリングすること。完全なダンプも、エンティティを追加するたびに大きくなります。インデックスはドメインのみで増加します。
5.5倍の削減は素晴らしいですが、これはこの中で最小の数字です
シリーズ。ここで始めるポイントは、オントロジーが次のことを確立するということです。
パターン — そして、これから見るように、座標系 — が大きいほど、
パート 2 と 3 での節約が基礎となります。
モデルは、常に利用可能な 1 つのインデックスを通じてルート インデックス以下のすべてに到達します。
ツール、explore_ontology、3 つのモード:
explore_ontology({ entities: ["issue"] }) ドリルイン時 — エンティティ用語集、プル時に一度支払われたメモが含まれています - 問題 — プロジェクト内で追跡される作業: ステータス、優先度、担当者、ラベル、添付ファイル、およびブロック関係があります。注: 「ユーザーは問題を「チケット」または「作業項目」と呼ぶ場合があります。

注: 「ステータス名はプロジェクト ワークフローから取得されます。最初に解決してください。」 explore_ontology({ query: "release blocker" }) 検索時 — ID と説明に対する BM25 は、プル時に一度支払われます → 発行 — 「…他の問題またはリリースをブロックする可能性があります」で一致 → または noMatch: true — モデルは語彙がどこで終了するかを学習します ドメイン モデルの段階的な開示。継続的な 340 トークンが唯一の定期的なコストです。用語集、メモ、検索結果は、必要なときに一度支払われます。 (表現は軽く一般化されています。示されている注記は例示です。)
引数なしで呼び出すと、完全なプラットフォーム ドメイン マップが返されます。
現在の範囲を超えるタスクの避難用ハッチ。でエラーが発生しました
空の呼び出しはソース内で明示的に拒否されました。「強制的に呼び出しを行うだけです」
推測します。」 ID を渡すとドリルインします。ドメインはそのエンティティを返し、
関係と、それを通じて到達可能な隣接ドメイン。実体
は、その説明、メモ、およびそれに関連するすべての関係を返します。通過
クエリは ID と説明に対して BM25 キーワード検索を実行するため、「リリース」
ブロッカー」は、モデルを使用せずに依存関係の記述を通じて問題を発見します。
IDを知っていること。
インデックスも座標系です
これがプロンプト セクションを縮小するだけの場合、オントロジーは次のようになります。
やりすぎ — 手動で調整した段落は競合する可能性があります。グラフになっている理由
安定した ID とは、それらの ID が他のすべてのものの共有語彙になることです。
ハーネス内:
search_tools は同じドメイン/エンティティ/関係 ID を受け入れます
フィルターとして使用できるため、モデルは「問題に触れるツール」を要求できます。
関係でタグ付けされたツールは、どちらのエンドポイント エンティティからも見つけることができます。
各エンティティの用語集テキストは、すべてのツールの検索ドキュメントに組み込まれます
それに触れると、leが与えられます

厳選された同義語レイヤーをゼロで xical 検索します
埋め込みが関係しています。
記憶には同じ概念、つまり範囲指定された想起がタグ付けされます。
パート 3 の値参照も参加します: すべてのスコープ参照が知っています
getEntity() および getDomain() 。これはまさに forScope が消費するものです。
それがこのシリーズの難題です。オントロジーはプロンプト セクションではなく、
これは、他の 2 つのメカニズムがルーティングするアドレス空間です。
深さには決断が必要です。すべての遅延レイヤーはモデルの往復です。
作ることを選ぶこと。決してドリルを行わないモデルは 13 のワンライナーで動作します。
したがって、これらの行は信号として書き込む必要があります。
モデルは、引っ張る価値のあるものが下にあることを知ることができます。延期により、
トークンの予算編成から、権利を引き出すインデックスの作成までの大変な作業
引っ張ります。
地図は領土を追跡する必要があります。オントロジーは、次の記述です。
製品と2番目の説明がずれています。立っている部分がレイヤーです
モデルが最も信頼します。モデルはツールの結果としてではなく、システム プロンプトとして到着します。
古い説明は見逃されるだけでなく、あらゆる会話を誤解させます。
キュレーションは永続的な仕事です。インデックスは小規模かつ真実のままです。
それをそのままにしておきます。
オントロジーは、数百の既存の「ドメインが何を意味するか」を解決します
トークン。次の問題は一桁大きいものです。私たちのエージェントには 169
tools とその JSON スキーマだけで約 85,000 文字になります。それらを宣言する
すべてを行うと、他のすべてのプロンプト セクションを合わせたよりもコストが高くなります。そのため、ハーネスは
そうではなく、代わりに使用される機械が対象となります。
パート２。

## Original Extract

Systems architect and software engineer writing about interesting problems and hard-earned lessons.

Henrik Udnes Systems architect & software engineer
Context Engineering in an LLM Harness — Part 1: Ontology
Context Engineering in an LLM Harness · 5 parts
Imagine embedding an LLM agent in an issue tracker. A user opens the assistant
inside a release project and asks: “which open issues are blocking the
release?” The model needs product-specific knowledge: how projects, issues,
workflows and dependencies fit together; which status names this team uses;
and how internal ids map to the names people see. Echoing a raw id back at the
user is still a bug, not an answer.
The obvious move is to paste all of that into the system prompt. Everybody’s
first harness does this, ours included. It works, but its cost grows with the
domain: the system prompt is re-rendered on every iteration of every turn —
in our harness literally so, buildSystemPrompt runs once per LLM call — so
every character of standing context is a tax you pay for the lifetime of the
conversation. Domain knowledge only grows. At sufficient scale,
paste-everything becomes expensive, and the usual escape hatch — destructively
truncating history when things get tight — throws away exactly the context an
agent needs most.
This series is about the third option, the one our harness commits to
everywhere: defer, don’t discard . Keep an index in the standing context
and make the payload something the model pulls on demand. When even a pulled
result is too large, bound the view and keep the exact value behind its pointer.
Applied five times, at five different altitudes:
Ontology (this post) — what the domain is
Tool discovery — what the agent can do
Values by reference — the data flowing through it
Memory — what the agent remembers
Cost is only half the argument. A context window is also an attention budget:
the more that’s in front of a model, the worse it uses any one piece of it, and
irrelevant schemas and reference text don’t just cost tokens — they compete
with the task. What makes deferral work is that it is not omission :
everything stays reachable by a stable address, so shrinking the standing
context removes noise from the turn without removing information from the
system. Less context is more accuracy, exactly as long as nothing necessary
becomes unreachable — and keeping the index rich enough that the model knows
what to pull is what the mechanisms in this series are careful about.
The mechanisms and measurements here come from a production harness; the
running examples use a generic issue tracker so the architecture stays legible
without product knowledge. Token estimates use the rough conversion of one
token per four characters.
explore_ontology Tool surface namespace line + scaffolding + learned/active set Tool catalog 169 tools · ≈85k chars of schema search_tools activate_tool Skill catalog id — title: description per skill Skill bodies full instructions while active activate_skill deactivate_skill Memory digest capped index of {key, description} Memory bank capped records, key-addressed recall_memories Tool results bounded previews + $ref stubs Blob store unbounded bytes, server-side $ref resolution The shape of the whole series: the standing context holds five small indexes; every payload lives behind a pull. This post is the first row.
A graph, not a document
The ontology is not prose. It’s a typed graph: 15 domains (bounded
contexts such as issues , workflows and labels ), 28 entities , each
owned by exactly one domain, and 12 relationships with ids derived as
<from>.<predicate>.<to> :
export const ISSUE_ASSIGNED_TO_USER = new Relationship (
ISSUE,
" assigned-to " ,
USER,
" An issue is assigned to the user responsible for it. "
);
projects issues labels workflows users attachments belongs-to labeled-with assigned-to defined-by contains project issue label status workflow user attachment file status · note (drill-in only) “Users may call statuses ‘stages’ or ‘columns’.” The production graph pattern mapped onto an issue tracker. Entities live inside their owning domain; relationships may span domains. Notes carry operational knowledge that ships only when something drills in.
Two design rules do a lot of quiet work here.
First, references into the graph are validated, not stringly-typed .
Anything that tags itself with a concept — a tool declaring
semantics = [ISSUE] , a memory scoped to an entity — is checked against
the ontology, so a bad id fails at authoring time instead of silently matching
nothing at runtime. The dependency is strictly one-way: tools point into the
ontology; the ontology has never heard of a tool. And a tool’s domain
membership is derived from its tags, never authored, so it can’t drift out
of sync.
Second, knowledge lives once . Entity notes — the operational gotchas: the
words users actually say instead of the entity’s official name, the lookups to
perform before presenting something, the edits that silently fail — are data on
the entity, rendered wherever the entity is rendered. The interface docs are
explicit that notes are “never hand-written into prompt strings” . There is
exactly one place to fix a wrong fact.
What actually ships every turn
Here is the entire standing footprint of the ontology, from Ontology.toPrompt() :
public toPrompt (): string {
const roots = this .domainList
. filter (( d ) => d. getParent () === null );
if (roots.length === 0 ) return "" ;
const lines : string [] = [ " Domains in this scope: " ];
for ( const domain of roots)
lines. push ( `- ${ domain.id } — ${ domain.description } ` );
return lines. join ( " \n " );
}
Root domains only, one line each. The in-code comment states the intent better
than I could: deliberately brief — a standing orientation (valid ids and where
things live), not a knowledge dump . Before rendering, forScope filters
the graph down to what the current scope can even mean: domains that can’t
apply where the agent is standing are dropped wholesale, and whatever the user
is currently looking at — the scope references passed to the agent — pulls its
owning domains in. A typical scope leaves 13 lines:
1,356 chars ≈ 340 tok Full-graph dump 28 entities + notes + 12 relationships ≈7.5k chars ≈ 1.9k tok Measured from the ontology source: the standing section vs. rendering the full graph (all entity descriptions, notes and relationships) into the prompt. The full dump would also grow with every entity we add; the index grows only with domains.
A five-and-a-half-times reduction is nice, but it’s the smallest number in this
series. The point of starting here is that the ontology establishes the
pattern — and, as we’ll see, the coordinate system — that the bigger
savings in parts 2 and 3 are built on.
The model reaches everything below the root index through one always-available
tool, explore_ontology , with three modes:
explore_ontology({ entities: ["issue"] }) On drill-in — the entity glossary, notes included paid once, when pulled - issue — Work tracked in a project: has a status, priority, assignee, labels, attachments and blocking relationships. Note: “Users may call an issue a ‘ticket’ or a ‘work item’.” Note: “Status names come from the project workflow — resolve first.” explore_ontology({ query: "release blocker" }) On search — BM25 over ids and descriptions paid once, when pulled → issue — matched on “…may block other issues or a release” → or noMatch: true — the model learns where the vocabulary ends Progressive disclosure of the domain model. The standing 340 tokens are the only recurring cost; glossaries, notes and search results are paid once, in the turn that needs them. (Wording lightly generalized; the notes shown are illustrative.)
Calling it with no arguments returns the full platform domain map — the
escape hatch for tasks that reach beyond the current scope. Erroring on the
empty call was explicitly rejected in the source: it “would just force a
guess.” Passing ids drills in: a domain returns its entities and
relationships plus the neighbouring domains reachable through them; an entity
returns its description, its notes, and every relationship touching it. Passing
a query runs BM25 keyword search over ids and descriptions, so “release
blocker” finds issue through its dependency description without the model
knowing the id.
The index is also a coordinate system
If this were only about shrinking a prompt section, an ontology would be
overkill — a hand-tuned paragraph could compete. The reason it’s a graph with
stable ids is that those ids become the shared vocabulary for everything else
in the harness:
search_tools accepts the same domains / entities / relationships ids
as filters, so the model can ask for “tools touching issue ” — and a
tool tagged with a relationship is findable from either endpoint entity.
Each entity’s glossary text is folded into the search document of every tool
that touches it, giving lexical search a curated synonym layer with zero
embeddings involved.
Memories are tagged with the same concepts, scoping recall.
Even the value references of part 3 participate: every scope reference knows
getEntity() and getDomain() , which is exactly what forScope consumes.
That’s the red thread of this series: the ontology isn’t a prompt section,
it’s the address space the other two mechanisms route through.
Depth costs a decision. Every deferred layer is a round trip the model has
to choose to make. A model that never drills is operating on 13 one-liners,
so those lines have to be written as signals — descriptive enough that the
model can tell something worth pulling sits underneath. Deferral moves the
hard work from budgeting tokens to writing an index that provokes the right
pulls.
The map must track the territory. An ontology is a second description of
the product, and second descriptions drift. The standing section is the layer
the model trusts most — it arrives as system prompt, not as a tool result — so
a stale description doesn’t just miss, it misleads every conversation.
Curation is a permanent chore: the index stays small and true because someone
keeps it that way.
The ontology solves “what does the domain mean” for a few hundred standing
tokens. The next problem is an order of magnitude larger: our agent has 169
tools , and their JSON schemas alone are ~85,000 characters. Declaring them
all would cost more than every other prompt section combined — so the harness
doesn’t, and the machinery it uses instead is the subject of
part 2 .
