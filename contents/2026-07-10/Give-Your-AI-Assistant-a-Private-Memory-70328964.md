---
source: "https://hister.org/posts/give-your-ai-assistant-a-private-memory"
hn_url: "https://news.ycombinator.com/item?id=48859214"
title: "Give Your AI Assistant a Private Memory"
article_title: "Hister"
author: "mstef"
captured_at: "2026-07-10T13:03:48Z"
capture_tool: "hn-digest"
hn_id: 48859214
score: 3
comments: 0
posted_at: "2026-07-10T12:52:28Z"
tags:
  - hacker-news
  - translated
---

# Give Your AI Assistant a Private Memory

- HN: [48859214](https://news.ycombinator.com/item?id=48859214)
- Source: [hister.org](https://hister.org/posts/give-your-ai-assistant-a-private-memory)
- Score: 3
- Comments: 0
- Posted: 2026-07-10T12:52:28Z

## Translation

タイトル: AI アシスタントにプライベート メモリを与えましょう
記事タイトル: ハイスター
説明: Hister MCP を使用すると、AI アシスタントがあなたが読んだページを検索し、保存されているプレビューを取得し、独自のインデックスから回答できるようになります。

記事本文:
ヒスター
AI アシスタントにプライベート メモリを与える |ヒスター
ヒスターデータセット
サポート svg]:px-3 bg-hister-indigo font-space border-brutal-border h-autorounded-none border-[3px] px-5 py-2.5 text-[13px] font-semibold tracking-[1px] text-white uppercase no-underlineshadow-[3px_3px_0_rgba(0,0,0,0.25)]transition-all hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-[1px_1px_0_rgba(0,0,0,0.25)]" href="https://github.com/asciimoo/hister" target="_blank" rel="noopener noreferrer"> GitHub 2026 年 6 月 30 日 AI アシスタントにプライベート メモリを与える
AI アシスタントは物事を説明するのは非常に上手ですが、通常は説明が苦手です。
自分自身の文脈の記憶。
彼らはあなたが先週どのドキュメントページを読んだのか知りません。彼らは知りません
この記事を読んで、あなたはライブラリを試してみることにしました。彼らはどの移住か分からない
従ったガイド、回避策があった GitHub の問題、またはセキュリティ
あなたがすでに信頼している書き込み。
そのコンテキストは多くの場合、ブラウザーの履歴、ブックマーク、ローカルのメモ、
半分覚えた検索クエリ。
Hister は、そのコンテキストをプライベート検索インデックスに変換します。 MCP サポートにより、次のことが可能になります。
また、そのインデックスを AI アシスタントが利用できるようにします。
ヒスターの簡単な紹介
Hister は、関心のあるページやファイルを検索できる個人用検索エンジンです。
ブラウザ拡張機能は、閲覧中にページのインデックスを自動的に作成します。コマンドライン
このツールはブラウザ履歴のインポート、ドキュメント サイトのクロール、URL の手動インデックス作成、
そしてローカルファイルを追加します。 Web インターフェイスを使用すると、すべてをフルテキストで検索できます
検索、フィールド フィルター、ラベル、ファセット、優先結果、およびオプションのセマンティクス
検索します。
重要な部分は、データがどこに存在するかです。
閲覧履歴、インデックス付きページ、検索クエリ、ローカル ファイルは次の場所に残ります。
Hister インスタンス。ほとんどの人にとって、それは彼らが同じ母親にとどまることを意味します

中国
または彼らが管理するサーバー上で。
これは、すべてのユーザーに対して個別のアシスタント スキルを設定する必要がないことも意味します。
使用しているサイト。アシスタントに個別の API アクセスを与える必要はありません。
GitHub、ドキュメント サイト、問題トラッカー、フォーラム、Wiki、その他のサイト
認証が必要なもの。有用なページがすでに Hister にある場合、
アシスタントは代わりにヒスターを検索できます。
Hister は、すべてのプライベート サービスをアシスタントに接続するのではなく、それを提供します。
すでにインデックスを作成したマテリアルに対する 1 つの検索インターフェイス。
そのため、Hister は単独でも日常の検索ツールとして役立ちます。 MCP がさらに追加
レイヤー: アシスタントは、関連するページから推測する代わりに、Hister に関連するページを尋ねることができます。
一般的なモデルの知識、またはランダムな Web 結果の取得。
実践的なヒスター MCP ワークフロー
Hister MCP が日常業務で役立つ方法をいくつか紹介します。共通の
パターンは単純です。アシスタントは最初にヒスターに質問し、次にページから作業し、
すでに独自のインデックスの一部となっているファイル。
漠然と覚えている記事を見つける
あなたは状況を知っています。バグや設計パターンについての適切な説明を読みました。
または展開に問題があるのに、タイトルやサイトを思い出せない。
ヒスターがなければ、アシスタントに広範な質問をすると、広範な答えが得られます。
役に立つかもしれませんが、あなたが覚えている情報源に基づいたものではありません。
Hister MCP が設定されている場合、次のことを尋ねることができます。
PostgreSQL の移行について読んだ記事を Hister インデックスで検索してください
ロックして、最も関連性の高い結果を要約します。アシスタントは、PostgreSQL 移行ロックの Hister を検索し、
保存された結果と、独自のインデックス内のページに基づいて回答します。
Hister サーバーでセマンティック検索が有効になっている場合、アシスタントは質問することもできます。
セマンティックマッチングのため。アイデアは覚えているが正確には覚えていないときに役立ちます
言葉。セマンティック検索が利用できない場合

サーバー上では、Hister は次のようにフォールバックします。
通常のキーワード検索。
すでに索引付けされているドキュメントを使用してコードを説明する
AI コーディング アシスタントは、多くの場合、広範なトレーニング データから回答します。それは便利ですが、
使用しているライブラリまたはフレームワークの正確なバージョンが間違っている可能性があります。
Hister はアシスタントに、より具体的なコンテキストのソースを提供します。
プロジェクトのドキュメントをクロールできます。
ヒスターインデックス --recursive
--allowed-domain=docs.example.com
--最大深さ=4
https://docs.example.com/ の代わりに、Hister が準備したドキュメント データセットから開始することもできます。
すべてを自分でクロールします。データセット ページにはインポートが含まれています
言語やプラットフォームのドキュメントなどの共通の参考資料のため、
アシスタントは、インポート直後に Hister を通じてこれらのドキュメントを検索できます。
次に、AI アシスタント内で次のように尋ねます。
Hister インデックスを使用して、接続を構成するための公式ドキュメントを見つけます
このライブラリのタイムアウト。次に、このコードにどのオプションが適用されるかを説明します。アシスタントは、インデックス付きドキュメントを検索し、プレビューを取得して、それらを使用できます。
コードを説明するときのページ。
これは、内部ドキュメント、プライベート Wiki、古い API に特に役立ちます。
通常の Web 検索では見つけるのが難しいバージョンやドキュメント。
リサーチを情報源に裏付けられたブリーフに変える
数日間にわたってトピックを調査すると、役立つ資料は通常終了します
タブと履歴に散在しています。
Hister を使用すると、閲覧がキャプチャになります。 MCP を使用すると、アシスタントがそれを変えることができます
資料をブリーフにまとめました。
過去のサプライ チェーン攻撃に関するページをヒスター履歴で検索します。
月。調査結果をインシデントごとにグループ化し、影響を受けるプロジェクトをリストし、
使用したページ。アシスタントは、Hister 検索ツールで日付フィルターを使用して、データを取得できます。
最も関連性の高いページのプレビューを保存します。結果は良くない

ジェネリックではない
サプライチェーン攻撃の概要。実際に作った資料をまとめたものです
読み取りまたはインデックス化されます。
これは、セキュリティ研究、学術文献レビュー、市場に適しています。
研究、プロジェクト計画、およびソースの追跡可能性が重要なワークフロー。
何を検索すればよいのかさえわからない場合もあります。あなたはあなたがそうだったということだけを知っています
前に関連するものを読んでいます。
get_history ツールを使用すると、アシスタントは最近インデックス付けされたページや
最近開いたヒスターの結果。
最近インデックスされた Hister ページを見て、どのページが関連しているかを教えてください。
私たちがデバッグしているキャッシュの問題。アシスタントは最近の履歴から開始し、可能性のあるソースを特定してから、
検索とプレビューを取得してさらに深く掘り下げます。
これは、アシスタントにオープン Web 全体へのアクセスを許可することとは異なります。それは
まず、作業中に作成したページの痕跡を確認します。
閲覧履歴から作業ログを作成する
作業ログは便利ですが、手動で書くと忘れがちです。もうヒスター
作業中にどのページにインデックスを付け、検索し、開いたかを把握しているため、
アシスタントはその証跡を使用して、日次または週次のログを作成できます。
今日、最近インデックスされた Hister ページを見てください。プロジェクトごとにグループ化し、
私が取り組んだ内容を要約し、ソース URL を含む短い作業ログを作成します。これは、アシスタントが結果をドラフトとして扱う場合に最も効果的です。あなたはまだ
最終的な作業ログに何を含めるかを決定しますが、再構築の退屈な部分は
ブラウザ履歴からその日を記録することは完了です。
ワークロギングを軽量に保つ方法の詳細については、kvch の投稿「Worklogging for the Lazy Dev」を参照してください。
Hister MCP では、履歴全体を AI プロバイダーに送信する必要はありません。
アシスタントは特定の検索を要求します。 Hister は一致する結果を返し、
プレビュー。アクセストークンまたはユーザートークンを使用してエンドポイントを保護できます

en。あなた
スキップ ルール、手動インデックス作成、
自動インデックス作成設定、ラベル、ドメイン ルール。
警告: Hister MCP を外部 AI を利用したアシスタントに接続した場合
プロバイダー、Hister から返される検索結果とプレビューは、次の宛先に送信される場合があります。
そのプロバイダーを会話コンテキストの一部として使用します。プライベートが暴露される可能性がある
閲覧データ、インデックス付きドキュメント、ページ タイトル、URL、スニペット、保存されたページ
内容。
マテリアルを外部に送信したくない場合は、ローカル モデルを使用します。
AIプロバイダー。これにより、アシスタントとヒスターの両方がインフラストラクチャ上に留まります。
コントロール。スキップ ルールを追加して、プライベート ページや機密ページがアクセスされないようにすることもできます。
そもそもヒスターなので、後から MCP を通じて返却することはできません。
個人の検索インデックスは機密性が高いため、これは重要です。それはあなたのことを反映します
読書、取り組んでいること、調査していること、忘れていること。
Hister は、何をインデックス化するか、どこに保存するかを決定できるように設計されています。
どのクライアントがそれを読めるか。
まず、「Hister の取得」ガイドに従い、 Hister を実行します。
そして、いくつかのコンテンツがインデックスされていることを確認してください。ブラウザの拡張機能は、
通常のブラウジング中にインデックスを構築する最も簡単な方法。インポートすることもできます
既存のブラウザ履歴やドキュメント サイトをクロールします。
次に、MCP クライアントが使用するように設定します。
http://127.0.0.1:4433/mcp 認証が有効な場合は、Hister アクセス トークンをベアラー トークンとして渡します。
完全なセットアップ ガイドは、MCP 統合ドキュメントで入手できます。
MCP は AI ツールを接続できるようにするプロトコルです
外部のデータソースやツールへ。 Hister は、 /mcp で MCP エンドポイントを公開します。
そのエンドポイントを通じて、アシスタントは次のことを行うことができます。
検索ツールを使用して、インデックス付けされたページとファイルを検索します。
get_preview を使用して、特定のページの保存されたプレビューを取得します。
最近インデックスされたものまたは最近のものを検査する

get_history を使用してヒスター履歴を開きました。
これにより、アシスタントが一般的な留守番電話からそれに近いものに変わります。
あなたがすでに読んだ資料を扱うことができる研究パートナーに。
多くのアシスタント ワークフローは、URL を再び取得することに依存しています。それは次の理由で失敗する可能性があります
ログインウォール、レート制限、削除されたページ、ボット保護、または変更されたコンテンツ。
Hister には、インデックス付きページのテキストとメタデータがすでに保存されています。 MCP get_preview ツールは、保存されているコンテンツを直接取得できます。
つまり、アシスタントはインデックスを作成したページのバージョンから作業できるということです。
サイトが後で返すものの代わりに。
読み終わるとページが変わりました。
このページはもう利用できません。
このページにはブラウザでのセッションが必要です。
このページは自動フェッチをブロックします。
アシスタントが個人アーカイブから作業できるようにしたいと考えています。

## Original Extract

Use Hister MCP to let AI assistants search pages you have read, retrieve stored previews, and answer from your own index.

Hister
Give Your AI Assistant a Private Memory | Hister
Hister DATASETS
Support svg]:px-3 bg-hister-indigo font-space border-brutal-border h-auto rounded-none border-[3px] px-5 py-2.5 text-[13px] font-semibold tracking-[1px] text-white uppercase no-underline shadow-[3px_3px_0_rgba(0,0,0,0.25)] transition-all hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-[1px_1px_0_rgba(0,0,0,0.25)]" href="https://github.com/asciimoo/hister" target="_blank" rel="noopener noreferrer"> GitHub June 30, 2026 Give Your AI Assistant a Private Memory
AI assistants are very good at explaining things, but they usually have a weak
memory of your own context.
They do not know which documentation page you read last week. They do not know
which article convinced you to try a library. They do not know which migration
guide you followed, which GitHub issue had the workaround, or which security
writeup you already trusted.
That context often lives in your browser history, bookmarks, local notes, and
half-remembered search queries.
Hister turns that context into a private search index. With MCP support, it can
also make that index available to AI assistants.
A Quick Introduction to Hister
Hister is a personal search engine for pages and files you care about.
The browser extension automatically indexes pages as you browse. The command line
tool can import browser history, crawl documentation sites, index URLs manually,
and add local files. The web interface lets you search everything with full-text
search, field filters, labels, facets, priority results, and optional semantic
search.
The important part is where the data lives.
Your browsing history, indexed pages, search queries, and local files stay in
your Hister instance. For most people, that means they stay on the same machine
or on a server they control.
This also means you do not need to set up a separate assistant skill for every
site you use. You do not need to give your assistant individual API access to
GitHub, documentation sites, issue trackers, forums, wikis, and other sites
that require authentication. If the useful page is already in Hister, the
assistant can search Hister instead.
Instead of wiring every private service into your assistant, Hister gives it
one search interface over material you already indexed.
That makes Hister useful as a daily search tool by itself. MCP adds another
layer: an assistant can ask Hister for relevant pages instead of guessing from
general model knowledge or fetching random web results.
Practical Hister MCP Workflows
Here are a few ways Hister MCP can be useful in everyday work. The common
pattern is simple: the assistant asks Hister first, then works from pages and
files that are already part of your own index.
Find the Article You Remember Vaguely
You know the situation: you read a good explanation of a bug, a design pattern,
or a deployment problem, but you cannot remember the title or the site.
Without Hister, you ask an assistant a broad question and get a broad answer.
It might be helpful, but it is not grounded in the source you remember.
With Hister MCP configured, you can ask:
Search my Hister index for the article I read about PostgreSQL migration
locking and summarize the most relevant result. The assistant can search Hister for PostgreSQL migration locking , inspect the
stored result, and answer based on the pages in your own index.
If semantic search is enabled on your Hister server, the assistant can also ask
for semantic matching. That helps when you remember the idea but not the exact
words. If semantic search is not available on the server, Hister falls back to
normal keyword search.
Explain Code with Documentation You Already Indexed
AI coding assistants often answer from broad training data. That is useful, but
it can be wrong for the exact version of a library or framework you use.
Hister gives the assistant a more specific source of context.
You can crawl the documentation for a project:
hister index --recursive
--allowed-domain=docs.example.com
--max-depth=4
https://docs.example.com/ You can also start from Hister’s prepared documentation datasets instead of
crawling everything yourself. The datasets page includes imports
for common reference material such as language and platform documentation, so
your assistant can search those docs through Hister immediately after import.
Then, inside your AI assistant, ask:
Using my Hister index, find the official docs for configuring connection
timeouts in this library. Then explain which option applies to this code. The assistant can search your indexed docs, retrieve previews, and use those
pages when explaining the code.
This is especially useful for internal documentation, private wikis, older API
versions, or documentation that is hard to find through a normal web search.
Turn Research into a Source-Backed Brief
When you research a topic over several days, the useful material usually ends
up scattered across tabs and history.
With Hister, browsing becomes capture. With MCP, the assistant can turn that
captured material into a brief.
Search my Hister history for pages about supply chain attacks from the last
month. Group the findings by incident, list the affected projects, and cite the
pages you used. The assistant can use date filters in the Hister search tool, then retrieve
stored previews for the most relevant pages. The result is not just a generic
summary of supply chain attacks. It is a summary of the material you actually
read or indexed.
This works well for security research, academic literature review, market
research, project planning, and any workflow where source traceability matters.
Sometimes you do not even know what to search for. You only know that you were
reading something relevant earlier.
The get_history tool lets an assistant inspect recently indexed pages or
recently opened Hister results.
Look at my recently indexed Hister pages and tell me which ones are relevant to
the caching problem we are debugging. The assistant can start from recent history, identify likely sources, then use
search and preview retrieval to dig deeper.
This is different from giving the assistant access to the whole open web. It is
looking first at the trail of pages you created while working.
Create a Worklog from Your Browsing History
Worklogs are useful, but writing them manually is easy to forget. Hister already
knows which pages you indexed, searched, and opened while working, so an
assistant can use that trail to draft a daily or weekly log.
Look at my recently indexed Hister pages from today. Group them by project,
summarize what I worked on, and create a short worklog with source URLs. This works best when the assistant treats the result as a draft. You still
decide what belongs in the final worklog, but the boring part of reconstructing
the day from browser history is done for you.
For more details on keeping worklogging lightweight, see kvch’s post, Worklogging for the Lazy Dev .
Hister MCP does not require sending your whole history to an AI provider.
The assistant asks for specific searches. Hister returns matching results and
previews. You can protect the endpoint with an access token or user token. You
can also control what enters the index with skip rules, manual indexing,
automatic indexing settings, labels, and domain rules.
Warning: if you connect Hister MCP to an assistant backed by an external AI
provider, the search results and previews returned by Hister may be sent to
that provider as part of the conversation context. That can expose private
browsing data, indexed documents, page titles, URLs, snippets, and stored page
content.
Use a local model when you want to avoid sending that material to an external
AI provider. This keeps both the assistant and Hister on infrastructure you
control. You can also add skip rules to keep private or sensitive pages out of
Hister in the first place, so they cannot be returned through MCP later.
This matters because a personal search index is sensitive. It reflects what you
read, what you work on, what you investigate, and what you forget.
Hister is designed so you can decide what gets indexed, where it is stored, and
which clients can read it.
First, follow the Obtaining Hister guide , run Hister ,
and make sure you have some content indexed. The browser extension is the
easiest way to build the index during normal browsing. You can also import
existing browser history or crawl documentation sites.
Then configure your MCP client to use:
http://127.0.0.1:4433/mcp If authentication is enabled, pass your Hister access token as a bearer token.
The full setup guide is available in the MCP integration docs .
MCP is a protocol that lets AI tools connect
to external data sources and tools. Hister exposes an MCP endpoint at /mcp .
Through that endpoint, an assistant can:
Search your indexed pages and files with the search tool.
Retrieve the stored preview of a specific page with get_preview .
Inspect recently indexed or recently opened Hister history with get_history .
This changes the assistant from a general answer machine into something closer
to a research partner that can work with the material you have already read.
Many assistant workflows rely on fetching URLs again. That can fail because of
login walls, rate limits, removed pages, bot protection, or changed content.
Hister already stores the text and metadata of indexed pages. The MCP get_preview tool can retrieve that stored content directly.
That means an assistant can work from the version of the page you indexed,
instead of whatever the site returns later.
The page changed after you read it.
The page is no longer available.
The page requires a session in your browser.
The page blocks automated fetches.
You want the assistant to work from your personal archive.
