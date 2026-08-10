---
source: "https://blog.cloudflare.com/ai-search-easier/"
hn_url: "https://news.ycombinator.com/item?id=49241903"
title: "Cloudflare AI Search: give your agents a search engine for your data"
article_title: "Cloudflare AI Search: give your agents a search engine for your data | Cloudflare Blog"
author: "ankushKun"
captured_at: "2026-08-10T10:58:18Z"
capture_tool: "hn-digest"
hn_id: 49241903
score: 2
comments: 0
posted_at: "2026-08-10T10:36:55Z"
tags:
  - hacker-news
  - translated
---

# Cloudflare AI Search: give your agents a search engine for your data

- HN: [49241903](https://news.ycombinator.com/item?id=49241903)
- Source: [blog.cloudflare.com](https://blog.cloudflare.com/ai-search-easier/)
- Score: 2
- Comments: 0
- Posted: 2026-08-10T10:36:55Z

## Translation

タイトル: Cloudflare AI Search: エージェントにデータの検索エンジンを提供します
記事のタイトル: Cloudflare AI Search: エージェントにデータの検索エンジンを提供 |クラウドフレアのブログ
説明: AI Search により、Cloudflare プリミティブをつなぎ合わせる必要がなく、検索がこれまでより簡単になります。データにポイントして、独自のファイルや Web サイトの検索を作成します。新しい料金モデルのプレビューも共有しています。

記事本文:
Cloudflare AI Search: エージェントにデータの検索エンジンを提供します |クラウドフレアのブログ
コンテンツへスキップ すべてのカテゴリ AI
ログイン 営業担当者へのお問い合わせ ブログ エージェント エージェント週間 AI +4 さらに 4 つのタグを表示 7 つのタグ 7 つのタグを表示 選択したタグ
エージェント エージェントウィーク AI AI 検索 開発者 プラットフォーム 開発者 製品ニュース
プラットフォームの自動最適化
Cloudflare 1 ユーザーのリスクスコア
Cloudflare AI Search: エージェントにデータの検索エンジンを提供します
ネルソン・ドゥアルテ、ティアゴ・テレス、アニ・ワン
本日、Cloudflare AI Search の開発者エクスペリエンスがいくつか改善され、すぐに使える検索ソリューションの管理が容易になることを発表できることを嬉しく思います。以前は、Cloudflare プリミティブのコンポーネント (Workers AI、AI Gateway、Vectorize、R2、Browser Run) をつなぎ合わせる必要がありましたが、AI Search はこれを自動的に、より適切に実行できるようになりました。私たちの目標は、エージェントに独自の検索エンジンを提供し、エージェントがデータを簡単に見つけて、自分自身や担当者にとってより良い答えを提供できるようにすることです。
また、AI Search の顧客向けに価格の初期プレビューも共有しているので、これがどのように拡張されるかを知ることができます。私たちは、予測可能でスケーラブルになる方法で価格をモデル化しました。デフォルトのモデルを使用する場合、埋め込みと再ランキングは無料なので、トークン数の予測について心配する必要はありません。
エージェント向けにデータのコレクションにインデックスを付ける: 個々のファイルから所有する Web サイトまで、エージェントが構築できる構造化データと非構造化データに簡単にアクセスできるようにします。 (現在、これは Cloudflare アカウント上のゾーンである必要がありますが、所有権を確認するためのより多くの方法が近日中に追加される予定です。)
Web サイトのサイトマップをスキップする: 以前、AI Search では、Web サイト統合を使用するには Web サイトにサイトマップが必要でした。 「検出」解析オプションを選択して、サイトマップのない Web サイトをソースとして追加できるようになりました。
PU を 1 つ取得する

名前空間全体を検索するための blic エンドポイント: 名前空間でパブリック URL を有効にすると、認証なしで複数のインスタンスまたは Web サイトを一度に検索できる /search および /mcp エンドポイントを取得できるため、顧客と簡単に共有できます。
パブリック エンドポイント上に独自のカスタム ドメインを配置する: パブリック URL 上に独自のドメインを追加できるようになり、/search および /mcp エンドポイント (例: search.example.com/mcp ) をブランド化できます。 Cloudflare Accessを追加してプライベート検索インスタンスを作成することもできます。
AI 検索プラグインを使用して EmDash 上に構築されたサイトにセマンティック検索を追加する: サイトがオープンソース CMS である EmDash で実行されている場合、AI 検索プラグインはコンテンツにセマンティック検索を追加します。
AI Search の新しい価格モデルをプレビューします。価格設定が予測可能で、お客様に合わせて拡張できるようにしたいため、埋め込みと再ランキングのコストが組み込まれています。Workers AI カタログから選択したモデルを使用する場合、それらは無料です。
最後に、Cloudflare.com、開発者ドキュメント、EmDash、Cloudflare Dev Stack MCP、さらには今読んでいるブログ投稿など、独自のプラットフォーム全体で AI Search がどのように使用されているかの例も共有します (cmd+K を試してください)。
AI Search の動作: 新しい Cloudflare Dev Stack MCP を強化
AI Search の使用方法の 1 つは、新しい Cloudflare Dev Stack MCP にあります。これは、AI Playground で今すぐ試すことができます。コーディングエージェントには、Cloudflare開発者エコシステム全体から引用された最新のドキュメントが提供されるため、古いトレーニングデータの代わりに最新の機能と修正に基づいて構築されます。
現在 AI Search で利用できる機能を使用して構築した方法は次のとおりです。
Cloudflare が所有するサーフェスごとに 1 つの AI Search インスタンスを作成しました: Docs、Blog、API Docs、Community、Astro、Vite、Vitest、Hono、Replicate、OpenNext。 (これらはそれぞれ Cloudflare が所有しています。)
それらは異なるものにまたがります

ドメインをレンタルしますが、Cloudflare が Web サイトのデータを所有しているため、AI Search はそれらを 1 つのセットとして扱い、すべて同じ方法で取り込むことができます。 Point AI Search は 1 つのサイトまたは一連のサイトで、クロール、取り込み、埋め込み、取得を処理します。インスタンスの作成は 1 つのコマンドです。サイトマップのないサイトの場合は、 --parse-type Discover を追加して、次のリンクによってページを検索します (Browser Run の /crawl を利用します)。
Npx Wrangler AI - 検索作成 Cloudflare - コミュニティ \
-- 名前空間開発 - スタック \
-- ソース https://community.cloudflare.com \
--type web - クローラー \
-- parse -type Discover 2. インスタンスを 1 つの検索に結合します。
ここで興味深いのは、10 個のインスタンスすべてにわたって 1 つのクエリに答えることです。それには 2 つの方法があります。
オプション A: ワーカー内 (Cloudflare Stack MCP で行ったこと)
名前空間をワーカーにバインドしてリモート MCP サーバーを作成し、10 個のインスタンスすべてにわたって 1 つのマルチインスタンス呼び出しを作成しました。このパスを選択したのは、スタック検索を Cloudflare の MCP サーバーに追加するためで、エージェントがすでに接続している Cloudflare ツールと一緒にツールとして出荷されます。
wrangler.jsonc のバインディング:
{
"ai_search_namespaces" : [
{ "バインディング" : "AI_SEARCH" 、 "名前空間" : "cloudflare-stack" }
】
次に、単一のツールが 1 つの呼び出しを実行し、指定したインスタンス全体にファンアウトします。
// 1 つのツール、1 つの呼び出しで、名前空間内のすべてのサーフェスを一度に検索します。
コンテクスト。登録ツール (
'search_dev_stack' ,
{
説明: 「Cloudflare スタック全体で現在のドキュメントを検索します。」 、
入力スキーマ: z.オブジェクト ({ クエリ: z.string () })、
}、
async ({ クエリ }) => {
const res = context.env を待ちます。 AI_SEARCH 。検索({
クエリ、
ai_search_options: {
インスタンス ID: [ 'developers-cloudflare-com' , 'astro' , /* ...すべてのサーフェス */ ],
取得: { max_num_results: 10 }、
再ランキング: { 有効: true }、

}、
})
// res.chunks は引用され、元のインスタンスでタグ付けされて戻ってきます。
return { content: [{ type: 'text' , text: format (res.chunks) }] }
}
) オプション B: パブリック エンドポイントをオンにする (コードなし)
Worker をまったく作成したくない場合は、名前空間でパブリック URL を有効にします。すべてのインスタンスをクエリする /search および /mcp エンドポイントをすぐに取得できます。認証もデプロイするものもありません。
私たちと同じように、検索を既存のアプリまたは MCP サーバーに折りたたむ場合は、ワーカーに手を伸ばしてください。または、ワンクリックで共有可能な検索エンドポイントが必要な場合は、パブリック エンドポイントを使用します。
パブリック エンドポイントにはデフォルトのパブリック URL が付属していますが、独自のカスタム ドメインをその上に配置して、エンドポイントをブランド化することができます (例: search.example.com/mcp )。
検索をプライベートにする必要がある場合は、ドメインの前に Cloudflare Access を追加します。エンドポイントにはログインが必要となるため、承認された人 (またはエージェント) のみがエンドポイントをクエリできるようになります。
自分で試してみましょう: Dev Stack MCP を使用してください
Cloudflare Dev Stack MCP サーバーを使用すると、ツールについて質問したり、構築したいアプリについて説明したりできます。Cloudflare スタック上でアプリを構築する最適な方法について、引用された最新の回答が得られます。
AI Playground はチェックしてみる価値がありますが、本当の魔法は MCP をコーディング エージェントに接続することであり、スタックの現在のドキュメントは 1 回のツール呼び出しで入手できます。これは、遅くてトークンが多く、間違ったソースや古いソースに到達することが多い通常のフォールバック (Web 検索からページ全体を取得する) に代わるものです。選択したエージェントで使用するには、Dev Stack MCP URL を MCP 構成にドロップします。たとえば:
{
"mcpサーバー": {
"dev-stack" : { "url" : "https://stack.mcp.cloudflare.com/mcp" }
}
ブログ、開発者ドキュメント、Cloudflare.com での検索の強化
私たちは顧客と同じ方法で AI 検索を構築します。Cloudflare ブログの検索はすでにその上で実行されており、

開発者ドキュメントと Cloudflare.com がそれに参加します。これらはすべて、1 つのクエリでハイブリッド検索、セマンティック、キーワードを組み合わせて使用​​するため、「これは何をするのか」という自由形式の質問と、名前やキーワードの正確な検索の両方を処理します。私たちは最近、EmDash、新しいオープンソース CMS、そして新しい
EmDash AI Search の統合により、現在その検索が強化されています。また、これを自分の EmDash サイトに追加して、すぐにコンテンツに対して同じ検索を実行することもできます。
AI Search はすべてのボット ポリシーを尊重します
AI Search はバックグラウンドでの Browser Run /crawl によって機能しますが、さらに一歩進んで、独自のボット ID である Cloudflare-AI-Search で自身を識別します。 Browser Run と同様に、robots.txt に従い、不変のパブリック ユーザー エージェントで自身を識別し、サイトに導入されているボット コントロールを尊重します。
プレビュー価格: 予測可能な価格
AI Search は現在ベータ版のため無料であり、課金はまだ有効になっていません。開始前に十分な通知を電子メールでお送りします。一般提供に向けて、取り込み、ストレージ、クエリ、さらに埋め込みと再ランキングの料金のプレビューを以下に示します (プレビュー価格は請求が開始される前に変更される可能性があります)。
無料の月額割り当て (すべての Workers プラン)
セマンティック (ハイブリッドおよびベクトル検索)
一部の Workers AI モデルでは無料。サードパーティの場合は別途請求される
† 毎月 500 万個の取り込みトークンの単一プール。現在サポートされているあらゆるファイル タイプ (テキスト、画像など) をカバーします。 ‡ 両方のクエリ タイプで共有される、1 か月あたり 2,000 クエリの単一プール。
私たちの目標は、検索の対象となるモデルから始めて、予測できる価格を提供することです。埋め込むとテキストが検索一致のベクトルに変換され、再ランキングによって結果が並べ替えられるため、最も関連性の高いものが最初に表示されます。どちらも AI Search のデフォルト設定または私たちの場合は無料で実行されます。

Workers AI カタログからモデルを選択するため、インデックス作成やすべての検索の背後にあるモデルのコストを心配する必要はありません。回答の生成とクエリの書き換えは、選択したモデルで実行されるオプションの手順であり、Workers AI の使用量として請求されます。または、任意のモデル/プロバイダーで AI Gateway クレジットを使用することもできます。
プレビュー価格を含む請求書の例
これは、デフォルトの AI Search 埋め込みおよび再ランキング モデルを使用して月に 30,000 件のセマンティック クエリを実行する、20,000 ドキュメント データ ソース (テキストの約 2,000 万トークン) と 1,000 枚の画像 (それぞれ約 1,000 トークンを想定) 用の新しい AI Search インスタンスを作成するための、Workers Payed プランの月額請求のサンプルです。取り込みは約 10% の重複でチャンク化され、以下の × 1.1 として表示されます。
(テキストの 2000 万トークン + 画像の 100 万トークン) × 1.1 - 500 万の無料 = 1810 万のトークン
1M トークンのイメージ × 1.1 = 110 万トークン
選択した Workers AI モデルに含まれる使用量
選択した Workers AI モデルに含まれる使用量
画像は基本取り込みにカウントされ、画像アドオンのコストも発生します。ストレージは、ドキュメントあたり約 10 KB、画像あたり 1 MB を想定しています。インデックス作成は主に 1 回限りのコストであるため、その後の数か月はほとんどがクエリであり、21 ドル近くになります。
AI Search は今すぐ有効にして使用できます。これをサイトに向けて、セマンティック マッチングとキーワード マッチングの両方のハイブリッド検索をオンにすると、エージェントが使用できる独自のデータ用の検索エンジンが完成します。 1 つのコマンドでスピンアップします。
npx wrangler ai - 検索作成 - 検索 \
-- 名前空間 my - 名前空間 \
-- ソース https://my-website.com \
--type web - クローラー \
-- ハイブリッド - 検索 そこからクエリを実行し、 /mcp 経由でエージェントに接続するか、パブリック /search エンドポイントにカスタム ドメインを配置してユーザーと共有します。詳細については、AI Search のドキュメントをご覧ください。
新しい投稿の通知を受け取るために購読する
私たちは決してシャアしません

あなたのメールアドレスです。
ご購読いただきありがとうございます!受信箱をチェックして確認してください。
マルチテナントプラットフォームの開発
プライバシーの選択 セキュリティの問題を報告する |プライバシーポリシー |利用規約 | GDPR |商標検索は一時的に利用できません。製品
新しいタブで開きます 新しいタブで開きます 新しいタブで開きます すべてのカテゴリ AI

## Original Extract

AI Search makes search easier than ever, with no Cloudflare primitives to stitch together. Point it at your data to create a search for your own files and websites. We're also sharing a preview of our new pricing model.

Cloudflare AI Search: give your agents a search engine for your data | Cloudflare Blog
Skip to content All Categories AI
Login Contact Sales Blog Agents Agents Week AI +4 Show 4 more tags 7 Tags Show 7 tags Selected Tags
Agents Agents Week AI AI Search Developer Platform Developers Product News
Automatic Platform Optimization
Cloudflare One User Risk Score
Cloudflare AI Search: give your agents a search engine for your data
Nelson Duarte , Tiago Teles , and Anni Wang
Today, we’re excited to announce a few developer experience improvements to Cloudflare AI Search to make it easy to manage a search solution out of the box. Previously, you had to stitch together components of the Cloudflare primitives (Workers AI, AI Gateway, Vectorize, R2, Browser Run) but now, AI Search can do this automatically, and better. Our goal is to give your agents their own search engine, where they can easily find data to provide better answers for themselves and their humans.
We’re also sharing an early preview of pricing for customers of AI Search so you can learn how this scales. We modeled pricing in a way that makes it predictable and scalable: embedding and reranking are free when you use the default models, so no need to worry about predicting token count.
Index a collection of data for your agent: Make structured and unstructured data easily accessible for your agent to build with, from individual files to websites you own. (Today, it must be a zone on your Cloudflare account, but with more ways to verify ownership coming soon.)
Skip the sitemap for your websites: Previously, AI Search required that websites have a sitemap to use the website integration. Now you can select the “Discover” parsing option to add a website without a sitemap as a source.
Get a single public endpoint for searching across a namespace: When you enable public URLs on your namespace, you can get a /search and /mcp endpoint that can search through multiple instances or websites at once without authentication, so you can share easily with your customers.
Put your own custom domain over public endpoints: You can now add your own domains over your public URLs, so you can brand your /search and /mcp endpoints (e.g., search.example.com/mcp ). You can also add Cloudflare Access to create private search instances.
Add semantic search to your sites built on EmDash with AI Search plugin: If your site runs on EmDash , our open-source CMS, the AI Search plugin adds semantic search over your content.
Preview the new pricing model for AI Search: We want pricing to be predictable and to scale with you, so we built in the cost of embedding and reranking: they’re free when you use select models from the Workers AI catalog.
Finally, we will also share examples of how AI Search is used across our own platform including Cloudflare.com , our Developer Docs, with EmDash, in Cloudflare Dev Stack MCP — and even the blog post you’re reading right now (try cmd+K).
AI Search in action: powering the new Cloudflare Dev Stack MCP
One of the ways we use AI Search is in our new Cloudflare Dev Stack MCP , which you can try today in our AI Playground . It gives coding agents current, cited docs from across the Cloudflare developer ecosystem, so they build on the latest features and fixes instead of stale training data.
Here's how we built it using the features available today in AI Search:
We created one AI Search instance per Cloudflare-owned surface: Docs, Blog, API Docs, Community, Astro, Vite, Vitest, Hono, Replicate, OpenNext. (Each of these is Cloudflare-owned.)
They span different domains, but, because Cloudflare owns the website data, AI Search is able to treat them as a single set and ingest them all the same way. Point AI Search at a site, or set of sites, and it handles crawling, ingestion, embedding, and retrieval. Creating an instance is a single command, and for a site without a sitemap you add --parse-type discover to find pages by following links (powered by /crawl from Browser Run):
npx wrangler ai - search create cloudflare - community \
-- namespace dev - stack \
-- source https : //community.cloudflare.com \
--type web - crawler \
-- parse -type discover 2. Combine the instances into one search
Now the interesting part: answering a single query across all 10 instances. There are two ways to do it.
Option A: in a Worker (what we did for Cloudflare Stack MCP)
We bound the namespace to a Worker to create a remote MCP server and made one multi-instance call across all 10 instances. We took this path because we're adding the stack search into Cloudflare's MCP server , so it ships as a tool alongside the Cloudflare tools agents already connect to.
The binding, in wrangler.jsonc:
{
"ai_search_namespaces" : [
{ "binding" : "AI_SEARCH" , "namespace" : "cloudflare-stack" }
]
} Then a single tool makes one call that fans out across the instances you name:
// One tool, one call that searches every surface in the namespace at once.
context. registerTool (
'search_dev_stack' ,
{
description: 'Search current docs across the Cloudflare stack.' ,
inputSchema: z. object ({ query: z. string () }),
},
async ({ query }) => {
const res = await context.env. AI_SEARCH . search ({
query,
ai_search_options: {
instance_ids: [ 'developers-cloudflare-com' , 'astro' , /* ...every surface */ ],
retrieval: { max_num_results: 10 },
reranking: { enabled: true },
},
})
// res.chunks come back cited and tagged with the instance they came from.
return { content: [{ type: 'text' , text: format (res.chunks) }] }
}
) Option B: flip on public endpoints (no code)
If you'd rather not write a Worker at all, enable public URLs on the namespace. You immediately get /search and /mcp endpoints that query every instance, with no auth and nothing to deploy.
Reach for the Worker when you're folding search into an existing app or MCP server, as we are. Or reach for the public endpoint when you just want a shareable search endpoint in one click.
Public endpoints come with a default public URL, but you can put your own custom domain over them to brand the endpoint (e.g., search.example.com/mcp ).
If the search should be private, add Cloudflare Access in front of the domain. The endpoint now requires a login, so only authorized people (or agents) can query it.
Try it yourself: use the Dev Stack MCP
With the Cloudflare Dev Stack MCP Server, you can ask about any tool, or describe an app you want to build, and you'll get back current, cited answers on how best to build it on the Cloudflare stack.
The AI Playground is worth checking out, but the real magic is wiring the MCP into your coding agent, so the stack's current docs are one tool call away. That replaces the usual fallback (web search then fetching full pages), which is slow, token-heavy, and often lands on the wrong or stale source. To use with your agent of choice, drop the Dev Stack MCP URL into your MCP configuration. For example:
{
"mcpServers" : {
"dev-stack" : { "url" : "https://stack.mcp.cloudflare.com/mcp" }
}
} Powering search on our Blog, Developer Docs, and Cloudflare.com
We build with AI Search the same way our customers would: Cloudflare Blog's search already runs on it, and today Developer Docs and Cloudflare.com join it. All of it uses hybrid search, semantic and keyword together in one query, so it handles both open-ended "what does this do" questions and exact lookups of names or keywords. We recently rebuilt the Blog on EmDash, our new open-source CMS, and our new
EmDash AI Search integration is what powers that search now. You can also add it to your own EmDash site and get the same search over your content out of the box.
AI Search respects all bot policies
AI Search is powered by Browser Run /crawl in the background, but goes a step further to identify itself with its own bot identity: Cloudflare-AI-Search . Just like Browser Run, it follows robots.txt, identifies itself with an immutable, public user agent, and will respect whatever bot controls a site has in place.
Preview pricing: pricing you can predict
AI Search is currently free while in beta, and billing is not yet enabled; we'll email you with plenty of notice before it starts. As we move toward general availability, here's a preview of pricing across ingestion, storage, and queries, plus embedding and reranking ( preview prices are subject to change before billing begins) :
Free monthly allotment (all Workers plans)
Semantic (hybrid and vector search)
Free with select Workers AI models; third-party billed separately
† A single pool of 5M ingestion tokens per month, covering any file type currently supported (e.g., text, images). ‡ A single pool of 2,000 queries per month, shared across both query types.
Our goal is to provide pricing you can predict, starting with the models your search leans on. Embedding turns your text into the vectors that search matches on, and reranking reorders results so the most relevant come first. Both run free with AI Search defaults or when using select models from the Workers AI catalog, so the models behind indexing and every search are not a cost you have to worry about. Answer generation and query rewriting are optional steps that run on a model you choose, billed as Workers AI usage, or you can use AI Gateway credits with any model/provider.
Example bill with preview pricing
Here's a sample monthly bill on the Workers Paid plan for creating a new AI Search instance for a 20,000-document data source (about 20M tokens of text) plus 1,000 images (assume about 1,000 tokens each), with 30,000 semantic queries a month using the default AI Search embedding and reranking model. Ingestion is chunked with roughly 10% overlap, which shows up as the × 1.1 below:
(20M tokens of text + 1M tokens of images) × 1.1 - 5M free = 18.1M tokens
1M tokens of images × 1.1 = 1.1M tokens
Usage included with selected Workers AI model
Usage included with selected Workers AI model
Images count toward base ingestion and also incur the image add-on cost. Storage assumes about 10 KB per document and 1 MB per image. Indexing is largely a one-time cost, so later months are mostly queries, closer to $21.
AI Search is available to enable and use today. Point it at your site, turn on hybrid search for both semantic and keyword matching, and you have a search engine for your own data, ready for your agents. Spin one up with one command:
npx wrangler ai - search create my - search \
-- namespace my - namespace \
-- source https : //my-website.com \
--type web - crawler \
-- hybrid - search From there, query it, wire it into an agent over /mcp , or put a custom domain on a public /search endpoint to share it with your users. Check out the AI Search docs for more information.
Subscribe to receive notifications of new posts
We’ll never share your email address.
Thanks for subscribing! Check your inbox to confirm.
Multi-Tenant Platform Development
Your privacy choices Report security issues | Privacy Policy | Terms of use | GDPR | Trademark Search is temporarily unavailable. Products
opens in a new tab opens in a new tab opens in a new tab All Categories AI
