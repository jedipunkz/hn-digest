---
source: "https://flaky.build/helping-claude-code-to-find-undocumented-apis-from-the-code/"
hn_url: "https://news.ycombinator.com/item?id=48435241"
title: "Helping Claude Code to find undocumented DuckDB APIs from the code"
article_title: "Helping Claude Code to find undocumented APIs from the code | Flaky Build"
author: "resheku"
captured_at: "2026-06-07T15:34:47Z"
capture_tool: "hn-digest"
hn_id: 48435241
score: 2
comments: 0
posted_at: "2026-06-07T14:34:30Z"
tags:
  - hacker-news
  - translated
---

# Helping Claude Code to find undocumented DuckDB APIs from the code

- HN: [48435241](https://news.ycombinator.com/item?id=48435241)
- Source: [flaky.build](https://flaky.build/helping-claude-code-to-find-undocumented-apis-from-the-code/)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T14:34:30Z

## Translation

タイトル: コードから文書化されていない DuckDB API を見つけるための Claude Code の支援
記事のタイトル: Claude Code がコードから文書化されていない API を見つけるのを支援 |薄っぺらなビルド
説明: クロード コードは、DuckDB 拡張機能を作成できないと言い続けました。 GitHits MCP は、ドキュメントの代わりに実際のコードを検索し、最終的にクロードが必要な文書化されていない C++ API を見つけるのに役立ちました。

記事本文:
Claude Code がコードから文書化されていない API を見つけるのを支援 |薄っぺらなビルド
Flaky Build コードから文書化されていない API を見つけるための Claude Code の支援について
TL;DR: クロード・コードは素晴らしいですが、不可能ではないときに、物事は不可能であると自信を持って言うことがあります。 GitHits MCP サーバーは、Claude がドキュメントの代わりに実際のコードを検索して、ドキュメント化されていない DuckDB C++ API を見つけるのに役立ちました。これにより、Web アーカイブ拡張機能に適切な述語プッシュダウンを構築できるようになりました。
クロード・コードは素晴らしかったです 🚀
私は 11 月から Claude Code Pro を、12 月から Max を使用しています。本当に働き方が変わりました。コードベースを読んで本番コードを書く AI ですか?未来は間違いなくここにあります。
問題: 文書化されていない C++ API 😤
私は duckdb-web-archive を構築していました。これは、SQL から直接 Wayback Machine と Common Crawl をクエリするための DuckDB 拡張機能です。目標は、適切な述語プッシュダウンを実装することです。
SELECT URL 、タイムスタンプ
FROM wayback_machine( 'example.com/*' )
WHERE タイムスタンプ > '2020-01-01'
リミット10;
プッシュダウンを使用しない場合、拡張機能はすべてのデータを取得し、ローカルでフィルターします。プッシュダウンでは、WHERE と LIMIT が API 自体に送信されます ( &from=2020 、 &limit=10 )。データがかなり少なくなります。はるかに速くなります。
フィルタープッシュダウン: WHERE → APIパラメータ
リミットプッシュダウン：LIMIT→CDXサービス
プロジェクション プッシュダウン : &fl= を介して要求された列のみをフェッチします
個別のプッシュダウン : DISTINCT ON → API の &collapse=
問題は？これらに対する DuckDB の内部 C++ API は文書化されていません。 TableFunctionSet 、 FunctionData 、 TableFilterSet 、フィルター伝播を伴うバインド/初期化/スキャン ライフサイクルなど。ドキュメントには何もありません。
ドキュメント MCP についてグーグルで調べました。最初の推奨は Context7 でした。当然のことですが、ライブラリの最新のドキュメントを取得します。
役に立ちませんでした。見つかったものはすべて高レベルすぎました。内部 C++

必要な拡張 API は DuckDB のソース コードに存在します。ファーストパーティの拡張機能はこれらを使用します。ただし、公的文書はありません。
クロードは、私が望んでいることは不可能だと言い続けました。複数回。とても自信を持って。
しかし、私はそれが可能であることを知っていました。 httpfs や postgres_scanner などの DuckDB 独自の拡張機能は、リモート ファイルに対してまさにこれを行っていました。能力は存在した。クロードが見つけたどこにも文書化されていませんでした。
同僚が GitHits について言及しました。すべての GitHub から実際の例を検索するコード検索ツール。私の注意を引いたのは、私の出身地と同じ都市に建てられたことです。小さな世界。
ウェイティングリストに加わりました。かなり早く承認されました。
Claude Code で GitHits MCP を有効にした瞬間、すべてが変わりました。
GitHits はドキュメントだけを検索するわけではありません。実際のコードを検索します。何百万ものリポジトリ。実際の実装。私が DuckDB 拡張機能の述語プッシュダウンについて尋ねたとき、Claude は突然次の例を見つけることができました。
DuckDB 独自の拡張機能 (postgres_scanner、httpfs など)
同様の問題を解決するコミュニティ拡張機能
正確な関数シグネチャとパターンを備えた実際のコード
夜も昼も。クロードは「そんなことは不可能だ」という考えから、次の方法を正確に教えてくれました。
TableFunction::pushdown_complex_filter を実装する
TableFilterSet を使用してプッシュされたフィルターを抽出する
ステートフル スキャン用に FunctionData を接続する
バインド/初期化/スキャンのライフサイクルを正しく処理する
完全な述語プッシュダウンのサポートを備えた拡張機能が完成しました。ウェイバック マシン CDX API。共通のクロール インデックス サーバー。どちらも動作しています。
数千のレコードをフェッチしていたクエリが、わずか 6 つの HTTP リクエストを行うようになりました。すべてのフィルタリングがソースにプッシュされます。
とても感動したので、他の人を助けるために PR を作成しました: duckdb/extension-template#158 。 AI アシスタントを使用して DuckDB 拡張機能を構築するためのドキュメントを追加します。
「ここのドキュメントは非常に役に立ちます。

LLM エージェントの考慮事項。将来のリリースで拡張機能 API を文書化するための何らかの取り組みについて読んだ気がしますが、どこで行われたのか思いつきません。」
まさにそれが GitHits によって可能になったのです。コードベースに分散しており、どこにも統合されていない知識を見つけます。
クロード コードは強力ですが、その知識には限界があります。すべての内部 API を知っているわけではありません。 Context7 は時々役に立ちますが、私の経験ではこれ以上複雑な例は見つからず、実際のエンジニアリングには文書化されていない内部構造が含まれることがよくあります。ソースコード内にのみ存在するパターン。
GitHits はそのギャップを埋めます。検索可能なナレッジベースとしてすべてオープンソースです。正規の例。実戦テストされた実装。実際に動作するコード。
Claude Code を真剣に使用している場合、特に文書化されていないライブラリを使用している場合は、ツールキットに GitHits を追加してください。クロードの推論と GitHits のコード検索は強力な組み合わせです。
*duckdb-web-archive は、DuckDB コミュニティ拡張機能を通じて利用できます。 INSTALL web_archive FROM コミュニティで試してください。 *

## Original Extract

Claude Code kept saying my DuckDB extension was impossible to create. GitHits MCP saved the day—it searches real code instead of docs, finally helping Claude find the undocumented C++ APIs I needed.

Helping Claude Code to find undocumented APIs from the code | Flaky Build
Flaky Build About Helping Claude Code to find undocumented APIs from the code
TL;DR: Claude Code is amazing but sometimes confidently says things are impossible when they’re not. GitHits MCP server helped Claude find undocumented DuckDB C++ APIs by searching actual code instead of docs. This let me build proper predicate pushdowns for my web archive extension.
Claude Code has been amazing 🚀
I’ve been using Claude Code Pro since November and Max from December. It’s genuinely changed how I work. An AI that reads your codebase and writes production code? The future is definitely here.
The problem: undocumented C++ APIs 😤
I was building duckdb-web-archive . It’s a DuckDB extension for querying Wayback Machine and Common Crawl directly from SQL. The goal: implement proper predicate pushdowns .
SELECT url , timestamp
FROM wayback_machine( 'example.com/*' )
WHERE timestamp > '2020-01-01'
LIMIT 10 ;
Without pushdowns the extension fetches all data and filters locally. With pushdowns the WHERE and LIMIT get sent to the API itself ( &from=2020 , &limit=10 ). Way less data. Way faster.
Filter pushdown : WHERE → API parameters
Limit pushdown : LIMIT → CDX service
Projection pushdown : Only fetch requested columns via &fl=
Distinct pushdown : DISTINCT ON → API’s &collapse=
The problem? DuckDB’s internal C++ APIs for these aren’t documented. Things like TableFunctionSet , FunctionData , TableFilterSet , bind/init/scan lifecycle with filter propagation. None of it in the docs.
I googled about documentation MCPs. First recommendation was Context7. Makes sense—it fetches up-to-date docs for libraries.
Didn’t help. Everything it found was too high-level. The internal C++ extension APIs I needed exist in DuckDB’s source code. First-party extensions use them. No public documentation though.
Claude kept telling me what I wanted wasn’t possible. Multiple times. Very confidently.
But I knew it was possible. DuckDB’s own extensions like httpfs and postgres_scanner were doing exactly this with remote files. The capability existed. Just wasn’t documented anywhere Claude could find.
A colleague mentioned GitHits . A code search tool that finds real examples from all of GitHub. What caught my attention: built in the same city I’m from. Small world.
Joined the waitlist. Got approved pretty quickly.
The moment I enabled GitHits MCP in Claude Code, everything changed.
GitHits doesn’t search just the documentation. It searches actual code . Millions of repos. Real implementations. When I asked about predicate pushdowns in DuckDB extensions, Claude could suddenly find examples from:
DuckDB’s own extensions (postgres_scanner, httpfs, etc.)
Community extensions solving similar problems
Real code with exact function signatures and patterns
Night and day. Claude went from “this isn’t possible” to showing me exactly how to:
Implement TableFunction::pushdown_complex_filter
Use TableFilterSet to extract pushed filters
Wire up FunctionData for stateful scanning
Handle bind/init/scan lifecycle correctly
Completed the extension with full predicate pushdown support. Wayback Machine CDX API. Common Crawl Index Server. Both working.
A query that would’ve fetched thousands of records now makes just 6 HTTP requests. All filtering pushed to the source.
I was so impressed I created a PR to help others: duckdb/extension-template#158 . Adds docs for building DuckDB extensions with AI assistants.
“The documentation here is super helpful, even without LLM agent considerations. I believe I read about some sort of effort to document the extension API in a future release, but can’t think of where.”
That’s exactly what GitHits enabled. Finding knowledge scattered across codebases that isn’t consolidated anywhere.
Claude Code is powerful but its knowledge has a cutoff. Doesn’t know every internal API. Context7 is helpful sometimes but in my experience it can’t find more complex examples and real engineering often involves undocumented internals. Patterns that only exist in source code.
GitHits fills that gap. All of open source as a searchable knowledge base. Canonical examples. Battle-tested implementations. Code that actually works.
If you’re using Claude Code seriously—especially with less-documented libraries—add GitHits to your toolkit. Claude’s reasoning plus GitHits’ code search is a powerful combo.
*duckdb-web-archive is available through DuckDB community extensions. Try it with INSTALL web_archive FROM community; *
