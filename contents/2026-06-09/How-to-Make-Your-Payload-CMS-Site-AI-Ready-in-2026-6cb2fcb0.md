---
source: "https://focusreactive.com/blog/making-your-payload-cms-site-ai-ready-right-now/"
hn_url: "https://news.ycombinator.com/item?id=48462793"
title: "How to Make Your Payload CMS Site AI-Ready in 2026"
article_title: "How to Make Your Payload CMS Site AI-Ready Right Now"
author: "katyadrozd"
captured_at: "2026-06-09T16:03:24Z"
capture_tool: "hn-digest"
hn_id: 48462793
score: 1
comments: 0
posted_at: "2026-06-09T15:57:39Z"
tags:
  - hacker-news
  - translated
---

# How to Make Your Payload CMS Site AI-Ready in 2026

- HN: [48462793](https://news.ycombinator.com/item?id=48462793)
- Source: [focusreactive.com](https://focusreactive.com/blog/making-your-payload-cms-site-ai-ready-right-now/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T15:57:39Z

## Translation

タイトル: 2026 年にペイロード CMS サイトを AI 対応にする方法
記事のタイトル: 今すぐペイロード CMS サイトを AI 対応にする方法
説明: Next.js + ペイロード CMS サイトを AI 対応にします。AI クローラーを導入し、生の Markdown を提供し、@graph スキーマを接続して、LLM がコンテンツを検索して引用できるようにします。

記事本文:
メイン コンテンツにスキップ ヘッドレス CMS ペイロード NextJS AI ペイロード CMS サイトを今すぐ AI 対応にする方法
Next.js + ペイロード CMS サイトを AI 対応にします。AI クローラーを導入し、生の Markdown を提供し、@graph スキーマを接続して、LLM がコンテンツを検索して引用できるようにします。
Googlebot とは異なり、AI クローラーはサイトに大量のトラフィックを送り込み、生の Markdown を探し、カスタム メタデータをスキャンします。 Next.js と Payload に基づいて構築している場合、これらの最適化の実装には最小限のオーバーヘッドが必要ですが、構成を誤ったままにしておくと、LLM にとってコンテンツが見えなくなります。
この変化はすでに測定可能です。 Ahrefs が 300,000 のキーワードを分析したところ、AI 概要が表示されると、上位のオーガニック検索結果のクリック数が 34.5% 失われることがわかりました。 2026 年 2 月の追跡調査では、その数字は 58% でした。 Vercel では、6 か月間で ChatGPT から到着する新規サインアップの 1% 未満から 10% まで増加しました。交通の混合は静かに、しかし急速に進んでいます。
まずはアクセスしてください。 AI クローラーは、デフォルトでは Cloudflare の Bot Fight モードによってブロックされます。 CDN レベルで検索ボットのブロックを解除し、robots.ts を使用して検索インデクサー (OAI-SearchBot、Claude-SearchBot、PerplexityBot) を明示的に許可し、トレーニング スクレイパー (GPTBot、ClaudeBot、Google-Extended) をブロックします。リクエストがサーバーに到達するまで、ここでは他に何も関係ありません。
コンテンツ ネゴシエーションを通じて Markdown を提供します。 Accept: text/markdown のミドルウェア チェックを追加し、ルート ハンドラーから生の Markdown を返します。これにより、トークンのオーバーヘッドが HTML と比較して最大 80% 削減され、コンテンツが LLM パイプラインによって構造的に優先されるようになります (高信号、低労力)。
*構造化データを 1 つの @graph にリンクします。 3 つのスキーマ タイプ (階層の BreadcrumbList、引用可能な事実の FAQPage、著者と組織にリンクされた記事) は、AI 検索ボットが関連性と権威性を評価するために実際に使用するものです。構造なし

赤い著者である場合、モデルにはあなたのコンテンツを他の人のコンテンツよりも信頼する信号がありません。
*ドキュメント サイトを運営している場合を除き、llms.txt をスキップしてください。開発者ツール以外での採用は最大 2% であり、Google はそれを完全に無視し、AI 検索クローラーはプロンプトなしでそれを取得しません。代わりに、その時間を Markdown エンドポイントに費やしてください。
AI クローラーをどのようにしてサイトに到達させるのでしょうか?
Cloudflareのボットファイトモードでは、AIボットはデフォルトでブロックされます。 Cloudflare で明示的にブロックを解除し、 robots.ts で検索ボットをトレーニング スクレイパーから分離する必要があります。
何かを最適化する前に、クローラーがサーバーに到達できることを確認してください。
Cloudflare を使用している場合、Bot Fight モードは通常デフォルトでオンになっており、リクエストがアプリに到達する前に GPTBot と PerplexityBot をブロックします。プランのオプションを選択してください:
有料プラン - AI ボットの切り替えを [セキュリティ] → [ボット] に切り替えます。
無料プラン - 通過するユーザー エージェントと一致するアクションをスキップする WAF カスタム ルールを追加します。
どちらの方法でも目標は同じです。ダウンストリームで問題が発生する前に、リクエストがオリジンに到達する必要があります。
Cloudflare Security → Bots パネル: Bot Fight Mode がオン、AI Scrapers と Crawlers がオフ。
Next.js 側では、これを app/robots.ts で処理します。ワイルドカード ルールにより全員が参加できます。より賢い方法は、AI トレーニングをブロックしながら AI 検索を歓迎することです。LLM トレーニングの次のラウンドに黙ってフィードすることなく、ChatGPT と Claude でリアルタイムの回答を表示します。
現在、各主要ベンダーはこれら 2 つのジョブ用に個別のユーザー エージェントを出荷しています。きれいに分割してください。
「次」から { MetadataRoute } をインポートします。
デフォルト関数 robots() : MetadataRoute をエクスポートします。ロボット {
戻り値 {
ルール：[
{
ユーザーエージェント: [
「OAI-サーチボット」、
"ChatGPT ユーザー" ,
「クロードサーチボット」 、
"クロードユーザー" 、
「パープレキシティボット」 、
]、
許可: "/" 、
} 、
{
userAgent : [ "GPTBot" 、 "ClaudeBot" 、 "Google 拡張" ] 、
禁止: "/" 、
} 、
{
ユーザーAg

ent : "*" ,
許可: "/" 、
許可しない: [ "/admin" 、 "/api" ] 、
} 、
]、
サイトマップ : ` ${ プロセス .環境 。 NEXT_PUBLIC_SITE_URL } /sitemap.xml ` 、
} ;
}
デプロイして、3 ～ 5 日後にアクセス ログを確認します。 GPTBot または PerplexityBot からの成功したリクエストが表示されない場合でも、上流の何か (WAF、CDN、オリジン ファイアウォールなど) が依然としてリクエストをブロックしています。ネットワーク構成を監査し、ボトルネックを特定します。
各ページの Markdown エンドポイントを公開するにはどうすればよいでしょうか?
コンテンツ ネゴシエーションを使用して Markdown を提供します。ボットが Accept: text/markdown ヘッダーを持つ URL にヒットすると、生の Markdown を返します。
確かに、AI クローラーは HTML を解析できます。しかし、プレーンな Markdown はより効率的に処理されます。つまり、トークンが減り、レイアウト ノイズがゼロになり、構造がよりクリーンになります。 Cloudflare は、Markdown を提供する場合と HTML を提供する場合、自社の投稿の 1 つでトークンが約 80% 少ないと報告しました。そのため、LLM パイプラインは、利用可能な場合にはそれを優先します。そして、すべての公開ページの .md バージョンが、それらに提供できる最も信号の高いサーフェスである理由もわかります。
実装には 2 つの部分があります。
ミドルウェア ( src/proxy.ts ) を介して書き換えを処理します。 Accept: text/markdown ヘッダーをチェックし、リクエストを内部の /md/* パスにルーティングします。これにより、ユーザーとクローラーの両方にとって URL がクリーンかつ均一に保たれ、.md 拡張子や追加の正規コードが不要になります。
ルート ハンドラー ( app/md/posts/[slug]/route.ts ) - ペイロードからドキュメントをフェッチし、プレーンな Markdown として返します。
// src/proxy.ts
"next/server" から { NextResponse } をインポートします。
"next/server" からタイプ { NextRequest } をインポートします。
デフォルト関数プロキシをエクスポート (リクエスト: NextRequest ) {
const accept = リクエスト。ヘッダー 。 get ( "受け入れる" ) ?? "" ;
if ( accept . include ( "text/markdown" ) ) {
NextResponse を返します。書き換えます(
新しい URL ( ` /md ${ request . nextUrl . pathn

アメ } `、リクエスト。 URL）、
) ;
}
}
エクスポート const 構成 = {
マッチャー: [ "/((?!_next/|api/|md/).*)" ] ,
} ;
// app/md/posts/[slug]/route.ts
import { getPayload } from "ペイロード" ;
"@payload-config" から configPromise をインポートします。
エクスポート非同期関数 GET (_req : リクエスト , {params } : {params : Promise } ) {
const { slug } = await params ;
const payload = await getPayload({config:configPromise});
const result = ペイロードを待ちます。見つける ( {
コレクション:「投稿」、
ここで: {
および: [ { スラッグ : { 等しい : スラッグ } } 、 { _status : { 等しい : "公開" } } ] 、
} 、
深さ: 1、
制限: 1、
ドラフト : false 、
overrideAccess : false 、
} ) ;
const post = 結果 。ドキュメント [0] ;
if ( !post ) return new Response ( "Not found" , { status : 404 } ) ;
新しい応答を返す ( buildMarkdown (post) , {
ヘッダー: {
"Content-Type" : "text/markdown; charset=utf-8" ,
変更: "受け入れる" 、
} 、
} ) ;
}
buildMarkdown は、H1 タイトル、本文、および必要なフィールド (出版日、抜粋、著者) などの応答を組み立てます。体内に何が入るかは、コンテンツをどのように保存するかによって異なります。
コレクションに contentMarkdown プレーンテキスト フィールドがある場合は、それを直接使用します。コンテンツが Lexical JSON (Payload 3.x のデフォルト) で存在する場合は、 @payloadcms/richtext-lexical ( docs ) の組み込みの ConvertLexicalToMarkdown を使用します。
カスタム ブロックを含む複雑なコンテンツ モデルの場合、リクエストごとにシリアル化しないでください。保存時に字句をマークダウンに変換し、ソースと一緒に保存する beforeChange フックを追加します。シリアル化は、読み取り時の操作ではなく、公開時の操作です。
Markdown 応答の Vary: Accept ヘッダーが重要です。これは、キャッシュと CDN に HTML と Markdown バリアントを別々に保存するように指示するため、ブラウザーが Markdown 本文を取得することはなく、エージェントが HTML ページを取得することもありません。
Accept: テキスト/マークダウンを返すカール

ブラウザが HTML を取得するのと同じ URL からのクリーンな Markdown 本文|697x273
AI 検索ボットが実際に読み取る構造化データはどれですか?
3 つの JSON-LD タイプ ( BreadcrumbList 、 FAQPage 、および著者/組織にリンクされた記事) は、AI 検索ボットが実際に使用する構造、引用、および典拠シグナルをカバーします。ページごとに 1 つの @graph の下にそれらを生成します。
AI 検索ボットは、きれいに解析できるコンテンツを好みます。 3 つの特定のスキーマ タイプがこの面倒な作業のほとんどを処理し、それらはすべてペイロード CMS コレクションに完全にマッピングされます。
これらを 1 つの @graph ブロックにドロップすることも、別個の JSON-LD スクリプトに分割することもできます。パーサーは両方を同じ方法で処理します。以下では @graph を使用します。
同じシグナルを送信する編集ツール
上記の構造化データは、そのデータに記述されているコンテンツによってのみ優れています。 2 つのプラグインは、ペイロード管理内のループを閉じます。 @payloadcms/plugin-seo は、タイトル、説明、画像の自動生成フックを公開します。これらを独自の関数に接続することで、編集者は手動でメタ フィールドを入力するのではなく、ドキュメントの実際のコンテンツからメタ フィールドを再生成できます。
コンテンツ自体については、payload-ai により AI 支援による生成がテキスト フィールドとリッチテキスト フィールドに直接追加されます。 GEO の実際的な価値は速度ではなく、一貫性です。FAQ スキーマ、作成者フィールド、および既に入力されている SameAs リンクを使用して作成されたコンテンツにより、構造化データが正確に指し示すことができます。
plugin-nested-docs からの階層
@payloadcms/plugin-nested-docs は、有効にしたコレクション内のすべてのドキュメントに 2 つのフィールドを追加します。親は自己参照関係 (編集者が親ドキュメントを選択します)、ブレッドクラムは祖先の自動入力配列で、各エントリにはラベルと URL が付いています。
ペイロード管理者: 親フィールドと plugin-nested-docs から自動生成されたブレッドクラム配列
階層は str の 1 つです

AI クローラーが使用する唯一のシグナル/docs/getting-started/installation のページは、本文を 1 行読む前に、モデルにその内容とそれがどこに当てはまるかをすでに伝えています。 BreadcrumbList JSON-LD はそのパスをマシンが読み取れるものに変換し、plugin-nested-docs がそれを構築します。すべてのドキュメントは、ラベルと URL を含む親ページの完全なチェーンをすでに認識しており、以下のスキーマにドロップする準備ができています。 URL ロジックを手動で記述する必要はなく、編集者が親ページの名前を変更しても何も壊れません。
FAQPage スキーマは、質問と回答を 1 つの機械可読ユニットにパッケージ化し、AI ボットがそのユニットを逐語的に持ち上げることができます (帰属も含め)。これを逆ピラミッドと考えてください。acceptedAnswer を 1 つまたは 2 つの短い文にまとめ、人間の読者のためにページ本文に詳細を残しておきます。長く曖昧な回答は引用されません。簡潔で具体的なものはそうします。ペイロードでは、専用の FAQ コレクションまたは単純な配列フィールドにより、編集者はこれを管理するための明確で構造化された方法を提供します。
著者と組織にリンクされている記事は、機械が権威を読み取る方法です。ページのテキストに書かれた自己紹介は彼らにとって何の役にも立ちません。リンクはマークアップ内に設定する必要があります。作成者は、検証済みのプロファイル (LinkedIn、GitHub、ORCID) を指す SameAs 配列を持つ個人です。パブリッシャーは、独自の SameAs およびロゴを使用して一度編集された組織グローバルです。構造化された作成者がいないと、権威のシグナルが失われ、LLM はコンテンツが信頼できるかどうかを推測するだけになります。
const jsonLd = {
"@context" : "https://schema.org" ,
"@グラフ" : [
{
"@type" : "パンくずリスト" ,
itemListElement : [
{
"@type" : "ListItem" ,
位置: 1 、
名前：「ホーム」、
項目: プロセス。環境 。 NEXT_PUBLIC_SITE_URL 、
} 、
...パンくずリスト。マップ ( ( アイテム , インデックス ) => ( {
"@type" : "ListItem" ,
位置: インデックス + 2、
名前:アイテム。ラベル、
私は

項目: アイテム。 URL 、
} ) )、
]、
} 、
{
"@type" : "FAQページ" ,
mainEntity : doc 。よくある質問 。マップ ( ( アイテム ) => ( {
"@type" : "質問" ,
名前:アイテム。質問、
受け入れられた回答 : {
"@type" : "回答" ,
テキスト: 項目 .回答プレーンテキスト 、
} 、
} ) )、
} 、
{
"@type" : "記事" ,
見出し: ドキュメント。タイトル、
発行日: doc 。公開された、
著者 : {
"@type" : "人" ,
名前: ドキュメント。著者 。名前、
同じ: ドキュメント 。著者 。ソーシャルリンク?マップ (( l ) => l .url ) 、
} 、
出版社 : {
"@type" : "組織" ,
名前：組織。名前、
ロゴ: {
"@type" : "ImageObject" ,
URL : org.ロゴ。 URL 、
} 、
} 、
} 、
]、
} ;
ここで重要なルールが 2 つあります。まず、すべてをサーバー側でレンダリングします。スキーマは DOM と一致する必要があります。クリック時にのみマウントされる JS アコーディオンの背後に FAQ を隠すと、JSON-LD はファントム マークアップを指すことになります。 2 番目に、組織をペイロード グローバルとして保持します。一度編集すれば、どこにでも挿入でき、ロゴや正式名称が変更されたときのデータのドリフトを防ぎます。出荷前に、Google のリッチリザルト テストで最終出力を検証します。
llms.txt が実際に重要になるのはどのような場合ですか?
ほとんどの場合、開発者向けドキュメントとして価値があります。 AI エージェントは、明示的に指定された場合にのみ llms.txt を取得します。そして、そのトラフィックはほぼ完全にコーディング エージェントが API ドキュメントを読んでいます。

[切り捨てられた]

## Original Extract

Make your Next.js + Payload CMS site AI-ready: let AI crawlers in, serve raw Markdown, and wire up @graph schema to let LLMs find and cite your content.

Skip to main content Headless CMS Payload NextJS AI How to Make Your Payload CMS Site AI-Ready Right Now
Make your Next.js + Payload CMS site AI-ready: let AI crawlers in, serve raw Markdown, and wire up @graph schema to let LLMs find and cite your content.
Unlike Googlebot, AI crawlers flood your site with heavy traffic, look for raw Markdown, and scan for custom metadata. If you’re building on Next.js and Payload, implementing these optimizations requires a minimal overhead - but leaving it misconfigured means invisible content for LLMs.
The shift is already measurable. Ahrefs analyzed 300,000 keywords and found that when AI Overviews appear, the top organic result loses 34.5% of its clicks. Their February 2026 follow-up put that figure at 58%. Vercel went from less than 1% to 10% of new signups arriving from ChatGPT in six months . The traffic mix is moving - quietly, but fast.
Access first . AI crawlers are blocked by Cloudflare's Bot Fight Mode by default. Unblock search bots at the CDN level, then use robots.ts to explicitly allow search indexers (OAI-SearchBot, Claude-SearchBot, PerplexityBot) while blocking training scrapers (GPTBot, ClaudeBot, Google-Extended). Nothing else here matters until requests reach your server.
Serve Markdown via content negotiation . Add a middleware check for Accept: text/markdown and return raw Markdown from a route handler. This cuts token overhead by ~80% versus HTML and makes your content structurally preferred by LLM pipelines — high signal, low effort.
*Link your structured data into one @graph . Three Schema types — BreadcrumbList for hierarchy, FAQPage for citable facts, and Article linked to Author and Organization — are what AI search bots actually use to assess relevance and authority. Without structured authorship, models have no signal to trust your content over anyone else's.
*Skip llms.txt unless you run a documentation site . Adoption is ~2% outside developer tooling, Google ignores it entirely, and AI search crawlers don't fetch it unprompted. Spend that time on the Markdown endpoint instead.
How Do You Let AI Crawlers Reach Your Site?
AI bots are blocked by default on Cloudflare’s Bot Fight Mode. You have to explicitly unblock them in Cloudflare, then split search bots from training scrapers in your robots.ts .
Before you optimize anything, confirm the crawlers can reach your server.
If you use Cloudflare, Bot Fight Mode is usually on by default and blocks GPTBot and PerplexityBot before requests reach your app. Choose the option for your plan:
Paid plans - flip the AI-bot toggles in Security → Bots .
Free plan - add a WAF Custom Rule with action Skip matching the user agents you want through.
Same goal either way: the request has to reach your origin before anything downstream matters.
Cloudflare Security → Bots panel: Bot Fight Mode on, AI Scrapers and Crawlers off.
On the Next.js side, you handle this in app/robots.ts . A wildcard rule lets everyone in. The smarter move is to welcome AI search while blocking AI training - surface in real-time answers on ChatGPT and Claude, without silently feeding the next round of LLM training.
Each major vendor now ships separate user agents for those two jobs. Split them cleanly.
import { MetadataRoute } from "next" ;
export default function robots ( ) : MetadataRoute . Robots {
return {
rules : [
{
userAgent : [
"OAI-SearchBot" ,
"ChatGPT-User" ,
"Claude-SearchBot" ,
"Claude-User" ,
"PerplexityBot" ,
] ,
allow : "/" ,
} ,
{
userAgent : [ "GPTBot" , "ClaudeBot" , "Google-Extended" ] ,
disallow : "/" ,
} ,
{
userAgent : "*" ,
allow : "/" ,
disallow : [ "/admin" , "/api" ] ,
} ,
] ,
sitemap : ` ${ process . env . NEXT_PUBLIC_SITE_URL } /sitemap.xml ` ,
} ;
}
Deploy, then check your access logs three to five days later. If you don't see successful requests from GPTBot or PerplexityBot , something upstream - like a WAF, CDN, or origin firewall - is still blocking them. It's time to audit your network configuration and locate the bottleneck.
How Do You Expose a Markdown Endpoint for Every Page?
Use content negotiation to serve Markdown. Return raw Markdown when bots hit your URLs with an Accept: text/markdown header.
Sure, AI crawlers can parse HTML. But they process plain Markdown way more efficiently - it means fewer tokens, zero layout noise, and a much cleaner structure. Cloudflare reported roughly 80% fewer tokens on one of their own posts when serving Markdown vs. HTML. That's why LLM pipelines prefer it when it's availablea. And why a .md version of every public page is the highest-signal surface you can offer them.
The implementation has two parts:
Handle rewrites via middleware ( src/proxy.ts ). It checks for the Accept: text/markdown header and routes requests to an internal /md/* path. This keeps the URLs clean and uniform for both users and crawlers, eliminating the need for .md extensions or extra canonicals.
A route handler ( app/md/posts/[slug]/route.ts ) - fetches the document from Payload and returns it as plain Markdown.
// src/proxy.ts
import { NextResponse } from "next/server" ;
import type { NextRequest } from "next/server" ;
export default function proxy ( request : NextRequest ) {
const accept = request . headers . get ( "accept" ) ?? "" ;
if ( accept . includes ( "text/markdown" ) ) {
return NextResponse . rewrite (
new URL ( ` /md ${ request . nextUrl . pathname } ` , request . url ) ,
) ;
}
}
export const config = {
matcher : [ "/((?!_next/|api/|md/).*)" ] ,
} ;
// app/md/posts/[slug]/route.ts
import { getPayload } from "payload" ;
import configPromise from "@payload-config" ;
export async function GET ( _req : Request , { params } : { params : Promise } ) {
const { slug } = await params ;
const payload = await getPayload ( { config : configPromise } ) ;
const result = await payload . find ( {
collection : "posts" ,
where : {
and : [ { slug : { equals : slug } } , { _status : { equals : "published" } } ] ,
} ,
depth : 1 ,
limit : 1 ,
draft : false ,
overrideAccess : false ,
} ) ;
const post = result . docs [ 0 ] ;
if ( ! post ) return new Response ( "Not found" , { status : 404 } ) ;
return new Response ( buildMarkdown ( post ) , {
headers : {
"Content-Type" : "text/markdown; charset=utf-8" ,
Vary : "Accept" ,
} ,
} ) ;
}
buildMarkdown assembles the response: an H1 title, the body, and whatever fields you need (publication date, excerpt, author). What goes inside the body depends on how you store content.
If your collection has a contentMarkdown plain-text field, use it directly. If content lives in Lexical JSON - the default in Payload 3.x - use the built-in convertLexicalToMarkdown from @payloadcms/richtext-lexical ( docs ).
For complex content models with custom blocks, don't serialize on every request. Add a beforeChange hook that converts Lexical to Markdown at save time and stores it alongside the source. Serialization is a publish-time operation, not a read-time one.
The Vary: Accept header on the Markdown response matters. It tells caches and CDNs to store HTML and Markdown variants separately, so a browser never gets the Markdown body and an agent never gets the HTML page.
curl with Accept: text/markdown returning a clean Markdown body from the same URL a browser would get HTML for|697x273
Which Structured Data Do AI Search Bots Actually Read?
Three JSON-LD types - BreadcrumbList , FAQPage , and Article linked to Author / Organization - cover the structural, citation, and authority signals AI search bots actually use. Generate them under one @graph per page.
AI search bots prefer content they can parse cleanly. Three specific Schema types handle most of this heavy lifting, and all of them map perfectly onto Payload CMS collections.
You can drop these into one @graph block or split them across separate JSON-LD scripts - parsers handle both the same way. We'll use @graph below.
Editorial tooling that feeds the same signals
The structured data above is only as good as the content it describes. Two plugins close that loop inside the Payload admin. @payloadcms/plugin-seo exposes auto-generate hooks for title, description, and image — you wire them to your own functions, so editors can regenerate meta fields from the document's actual content rather than filling them in by hand.
For the content itself, payload-ai adds AI-assisted generation directly to Text and RichText fields. The practical value for GEO isn't speed, it's consistency: content drafted with your FAQ schema, author fields, and sameAs links already populated gives the structured data something accurate to point at.
Hierarchy from plugin-nested-docs
@payloadcms/plugin-nested-docs adds two fields to every document in the collections you enable: parent is a self-referencing relationship (editors pick the parent doc), and breadcrumbs is an auto-populated array of ancestors, each entry with a label and a URL.
Payload admin: Parent field and auto-generated Breadcrumbs array from plugin-nested-docs
Hierarchy is one of the strongest signals AI crawlers use. A page at /docs/getting-started/installation already tells a model what it's about and where it fits - before reading a single line of body text. BreadcrumbList JSON-LD turns that path into something machines can read, and plugin-nested-docs builds it for you: every document already knows its full chain of parent pages, with labels and URLs, ready to drop into the schema below. No URL logic to write by hand, and nothing breaks when an editor renames a parent page.
An FAQPage schema packages questions and answers into a single machine-readable unit that AI bots can lift verbatim - attribution included. Think of it as an inverted pyramid: keep the acceptedAnswer to one or two tight sentences, and leave the deep details in the page body for human readers. Long, vague answers don’t get cited; concise, specific ones do. In Payload, a dedicated faqs collection or a simple array field gives editors a clean, structured way to manage this.
Article linked to Author and Organization is how machines read authority. A bio written in the page text does nothing for them. The links have to be set in the markup. The Author is a Person with a sameAs array pointing to verified profiles (LinkedIn, GitHub, ORCID); the publisher is an Organization global edited once with its own sameAs and logo . Without structured authorship, you lose the authority signal - leaving the LLM to just guess whether your content is trustworthy.
const jsonLd = {
"@context" : "https://schema.org" ,
"@graph" : [
{
"@type" : "BreadcrumbList" ,
itemListElement : [
{
"@type" : "ListItem" ,
position : 1 ,
name : "Home" ,
item : process . env . NEXT_PUBLIC_SITE_URL ,
} ,
... breadcrumbs . map ( ( item , index ) => ( {
"@type" : "ListItem" ,
position : index + 2 ,
name : item . label ,
item : item . url ,
} ) ) ,
] ,
} ,
{
"@type" : "FAQPage" ,
mainEntity : doc . faqs . map ( ( item ) => ( {
"@type" : "Question" ,
name : item . question ,
acceptedAnswer : {
"@type" : "Answer" ,
text : item . answerPlainText ,
} ,
} ) ) ,
} ,
{
"@type" : "Article" ,
headline : doc . title ,
datePublished : doc . publishedAt ,
author : {
"@type" : "Person" ,
name : doc . author . name ,
sameAs : doc . author . socialLinks ?. map ( ( l ) => l . url ) ,
} ,
publisher : {
"@type" : "Organization" ,
name : org . name ,
logo : {
"@type" : "ImageObject" ,
url : org . logo . url ,
} ,
} ,
} ,
] ,
} ;
Two key rules here. First, render everything server-side: your schema must match the DOM. If you hide your FAQ behind a JS accordion that only mounts on click, your JSON-LD will point to phantom markup. Second, keep Organization as a Payload global - edit it once, inject it everywhere, and prevent data drift when your logo or legal name changes. Validate the final output with Google's Rich Results Test before shipping.
When does llms.txt actually matter?
Mostly worth it for developer documentation. AI agents only fetch llms.txt when explicitly pointed at it - and that traffic is almost entirely coding agents reading API docs

[truncated]
