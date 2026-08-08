---
source: "https://agentblog.dev"
hn_url: "https://news.ycombinator.com/item?id=49220327"
title: "AgentBlog – open-source, AI-native SEO blog"
article_title: "AgentBlog: an SEO and GEO complete blog for Next.js 16 | AgentBlog"
author: "goldkey"
captured_at: "2026-08-08T10:21:26Z"
capture_tool: "hn-digest"
hn_id: 49220327
score: 2
comments: 0
posted_at: "2026-08-08T10:02:32Z"
tags:
  - hacker-news
  - translated
---

# AgentBlog – open-source, AI-native SEO blog

- HN: [49220327](https://news.ycombinator.com/item?id=49220327)
- Source: [agentblog.dev](https://agentblog.dev)
- Score: 2
- Comments: 0
- Posted: 2026-08-08T10:02:32Z

## Translation

タイトル: AgentBlog – オープンソースの AI ネイティブ SEO ブログ
記事のタイトル: AgentBlog: Next.js 16 の SEO および GEO に関する完全なブログ |エージェントブログ
説明: AI クローラーが読み取ることができる Next.js 16 ブログをインストールします: 事前レンダリングされた HTML、接続された JSON-LD グラフ、サイトマップ、RSS、IndexNow、OG 画像、投稿を作成および監査するエージェント スキル。

記事本文:
AgentBlog: Next.js 16 の SEO および GEO に関する完全なブログ | AgentBlog コンテンツへスキップ AgentBlog ドキュメント GEO プレイブック レジストリ ブログ ドキュメント GEO プレイブック レジストリ ブログ ランク付けされ、引用されるブログをインストールする
1 つのコマンドで、SEO に最適化された完全なブログが Next.js サイトにインストールされます。 100%無料＆オープンソース。
ブログを Next.js アプリにインストールし、検証します。
アプリを変更せずに、6 つのスキルを Claude Code にインストールします。
ドキュメント 実行中の様子を確認 GitHub で星 2 つ星 AI クローラーは JavaScript を実行しません
GPTBot、ClaudeBot、および PerplexityBot は HTML を 1 回フェッチし、テキストとして読み取ります。最初の応答に欠けているものは、彼らが書く答えにも欠けています。
<!doctype html> <html lang="ja"> <head> <meta charset="utf-8" /> <meta name="viewport" content="width=device-width" /> <link rel="stylesheet" href="/assets/index-a1c9f2.css" /> </head> <body> <div id="root"></div> <script type="module" src="/assets/index-7f2b04.js"></script> </body> </html> # 412 バイト。タイトルも説明も記事もありません。クライアントがレンダリングしたブログ。記事は存在しますが、それはバンドルの実行後にのみ存在します。カール -sA "GPTBot" https://agentblog.dev/blog/how-ai-search-engines-read-your-blog <!doctype html> <html lang="en"> <head> <title>AI クローラーは JavaScript を実行しますか? | AgentBlog</title> <meta name="description" content="いいえ。GPTBot、ClaudeBot、および PerplexityBot は HTML を 1 回フェッチし、テキストとして解析します。" /> <link rel="canonical" href="https://agentblog.dev/blog/…" /> <script type="application/ld+json">{"@context": "https://schema.org","@graph":[{"@type":"BlogPosting",… </head> <body> <article> <h1>AI クローラーは JavaScript を実行しますか?</h1> <p>いいえ。 Vercel と MERJ は、Vercel のネットワーク全体の AI クローラー トラフィックを 1 か月間にわたって計測し、主要な AI クローラーのどれも JavaScript をレンダリングしないと報告しました。</p> <h2

id="this-crawlers-render">どのクローラーがレンダリングしますか?</h2> <p>Googlebot がレンダリングします。 Bingbot は遅延してレンダリングします。… # 41 KB。記事全体を 1 回の返信で。 AgentBlog にも同じ投稿。クローラーに必要なものすべてが 1 つの応答で得られます。幅を考慮して省略した代表的な出力。右側のコマンドは実際のものです。
1 つのコマンドでブログ全体をインストールする
76 個のファイルがあり、それらすべてを所有しています。依存するランタイム パッケージも、アップグレードする必要のあるパッケージもありません。
.claude スキル アプリ API 著者ブログ [スラッグ] opengraph-image.tsx page.tsx カテゴリタグ レイアウト.tsx opengraph-image.tsx page.tsx 編集方針フィード.xml llms.txt not-found.tsx robots.ts サイトマップ.ts コンポーネント ブログ mdx コンテンツ ブログ authors.json カテゴリ.json フック use-toc-active-Heading.ts lib mdx-plugins ソース ai-referrers.ts config.ts define-config.ts Indexnow.ts metadata.ts og-card.tsx Posts.ts preflight-checks.ts preflight.ts reading-time.ts render-mdx.tsx schema.ts schemas.ts toc.ts Types.ts スタイル Agentblog.CSS Agentblog.config.ts AGENTS.agentblog.md app/blog/[slug]/page.tsx インポート タイプ { メタデータ } 'next' から
「next/image」から画像をインポート
'next/navigation' から { notFound } をインポートします
「@/components/blog/answer-capsule」から { AnswerCapsule } をインポートします
import { AuthorBio } から '@/components/blog/author-bio'
import { ブレッドクラム } から '@/components/blog/breadcrumbs'
import { Byline, PostDates } from '@/components/blog/byline'
'@/components/blog/json-ld' から { JsonLd } をインポートします
import { FAQ, FAQ_HEADING } from '@/components/mdx/faq'
import { Prose } から '@/components/blog/prose'
import { Relationships } から '@/components/blog/relative-posts'
'@/components/blog/share-buttons' から {ShareButtons } をインポートします
import { TableOfContents } から '@/components/blog/table-of-contents'
インポート {
CLUSTER_GAP、任意のファイルを開いてソースを読み取ります。

インストールします。各プレビューはファイル ヘッダーの後の最初の 16 行であり、ファイル全体がインストールとともに到着します。
すでに shadcn を利用していますか? npx shadcn@latest add @agentblog/blog は同じファイルをインストールし、npx Agentblog@latest Doctor --fix はレジストリが到達できない構成を終了します。
コーディングエージェントが投稿を書きます
投稿はリポジトリ内の MDX ファイルです。 AgentBlog をインストールすると、エージェントが読み取る 6 つのスキルとルール ブロックも追加されるため、エージェントは質問する前に形式を認識できます。
レジストリが単独で実行できなかった処理をすべて完了し、シード作成者、カテゴリ、投稿をあなたのものに置き換えます。これにより、ブログは私たちの主題ではなくあなたの主題に関するものになります。
サイトが何で知られるべきかを決定し、人々が実際に検索する質問を掘り起こし、内部リンクの方向がすでに決定されているバックログを作成します。
フォーマットに書き込みます: 最初に回答カプセル、質問の見出し、引用された統計、表内の比較データ。決して数字や引用をでっち上げないよう指示されています。
投稿で引用されているすべてのソースを再取得し、投稿の内容が依然として記載されていることを確認し、最小の正しい変更を適用します。何も変更する必要はありませんが、それが有効な結果となります。
デプロイされた URL を GPTBot として取得し、タイトルが頭に入っていることを確認し、公開する前にグラフを検証します。
投稿、インデックス、サイトマップ、フィードを再検証し、IndexNow に送信して、応答コードを報告します。 403 と 200 は、誰かが番号を読み取るまでは同じに見えます。
ブログで見つけなければならないものすべて
プリレンダリング、構造化データ、フィード、およびクローラー ワイヤリングが組み込まれています。投稿はユーザーが作成します。
すべての投稿は、ビルド時には完全な静的 HTML です。何も延期されることはなく、相互作用によって何も増加することはありません。
1 ページに 1 つのグラフ。 BlogPosting、個人、組織、WebSite、および FAQPage は、再記述される代わりに @id によって結合されます。
日付はあなたのものです

デプロイからではなく、コンテンツから。公開すると、変更された URL が Bing、Yandex、Seznam、および Naver にプッシュされます。
すべての見出しの下に直接の回答があり、その中にリンクはありません。その段落は、検索システムが持ち上げるチャンクです。
ブログは shadcn トークンを読み取り、プリミティブを構成します。テーマ、ベース、フォントはインストールされません。
デプロイされた URL を 5 つの異なるボットとして取得し、不足しているものに名前を付けます。 CIから実行します。
Next.js 16 の SEO と GEO に関する完全なブログ。所有するファイルとしてアプリにインストールされます。
コードはMITです。文書の散文は CC BY 4.0 に準拠します。シードポストCC0。

## Original Extract

Install a Next.js 16 blog that AI crawlers can read: prerendered HTML, a connected JSON-LD graph, sitemap, RSS, IndexNow, OG images, and agent skills that write and audit posts.

AgentBlog: an SEO and GEO complete blog for Next.js 16 | AgentBlog Skip to content AgentBlog Docs GEO playbook Registry Blog Docs GEO playbook Registry Blog Install a blog that ranks and gets cited
One command installs a complete, SEO optimized blog into your Next.js site. 100% Free & Open Source.
Installs the blog into your Next.js app, then verifies it.
Installs the six skills into Claude Code, with no app changes.
Documentation See it running Star on GitHub 2 stars AI crawlers do not run JavaScript
GPTBot, ClaudeBot, and PerplexityBot fetch your HTML once and read it as text. Whatever is missing from that first response is missing from the answer they write.
<!doctype html> <html lang="en"> <head> <meta charset="utf-8" /> <meta name="viewport" content="width=device-width" /> <link rel="stylesheet" href="/assets/index-a1c9f2.css" /> </head> <body> <div id="root"></div> <script type="module" src="/assets/index-7f2b04.js"></script> </body> </html> # 412 bytes. No title, no description, no article. A client-rendered blog. The article exists, but only after a bundle runs. curl -sA "GPTBot" https://agentblog.dev/blog/how-ai-search-engines-read-your-blog <!doctype html> <html lang="en"> <head> <title>Do AI crawlers run JavaScript? | AgentBlog</title> <meta name="description" content="No. GPTBot, ClaudeBot, and PerplexityBot fetch HTML once and parse it as text." /> <link rel="canonical" href="https://agentblog.dev/blog/…" /> <script type="application/ld+json">{"@context": "https://schema.org","@graph":[{"@type":"BlogPosting",… </head> <body> <article> <h1>Do AI crawlers run JavaScript?</h1> <p>No. Vercel and MERJ instrumented AI crawler traffic across Vercel’s network over a month and reported that none of the major AI crawlers render JavaScript.</p> <h2 id="which-crawlers-render">Which crawlers render?</h2> <p>Googlebot renders. Bingbot renders on a delay.… # 41 KB. The whole article, in one response. The same post on AgentBlog. Everything a crawler needs, in one response. Representative output, abridged for width. The right-hand command is real.
One command installs the whole blog
76 files, and you own every one of them. No runtime package to depend on, and nothing of ours to upgrade around.
.claude skills app api authors blog [slug] opengraph-image.tsx page.tsx category tag layout.tsx opengraph-image.tsx page.tsx editorial-policy feed.xml llms.txt not-found.tsx robots.ts sitemap.ts components blog mdx content blog authors.json categories.json hooks use-toc-active-heading.ts lib mdx-plugins sources ai-referrers.ts config.ts define-config.ts indexnow.ts metadata.ts og-card.tsx posts.ts preflight-checks.ts preflight.ts reading-time.ts render-mdx.tsx schema.ts schemas.ts toc.ts types.ts styles agentblog.css agentblog.config.ts AGENTS.agentblog.md app/blog/[slug]/page.tsx import type { Metadata } from 'next'
import Image from 'next/image'
import { notFound } from 'next/navigation'
import { AnswerCapsule } from '@/components/blog/answer-capsule'
import { AuthorBio } from '@/components/blog/author-bio'
import { Breadcrumbs } from '@/components/blog/breadcrumbs'
import { Byline, PostDates } from '@/components/blog/byline'
import { JsonLd } from '@/components/blog/json-ld'
import { Faq, FAQ_HEADING } from '@/components/mdx/faq'
import { Prose } from '@/components/blog/prose'
import { RelatedPosts } from '@/components/blog/related-posts'
import { ShareButtons } from '@/components/blog/share-buttons'
import { TableOfContents } from '@/components/blog/table-of-contents'
import {
CLUSTER_GAP, Open any file to read the source it installs. Each preview is the first sixteen lines after the file header, and the whole file arrives with the install.
Already on shadcn? npx shadcn@latest add @agentblog/blog installs the same files, and npx agentblog@latest doctor --fix finishes the config a registry cannot reach.
Your coding agent writes the posts
Posts are MDX files in your repository. Installing AgentBlog also adds six skills and a rule block your agent reads, so it knows the format before you ask.
Finishes whatever the registry could not do on its own, then replaces the seed author, categories, and posts with yours, so the blog is about your subject and not ours.
Decides what the site should be known for, mines the questions people actually search, and writes a backlog with the internal link direction already settled.
Writes to the format: answer capsule first, question headings, a cited statistic, comparison data in a table. It is instructed never to invent a figure or a quotation.
Re-fetches every source a post cites, checks it still says what the post says it says, and applies the smallest correct change. Nothing needed changing is a valid outcome.
Fetches your deployed URL as GPTBot, checks the title landed in head, and validates the graph before you publish.
Revalidates the post, the index, the sitemap, and the feed, submits to IndexNow, and reports the response code. A 403 and a 200 look identical until somebody reads the number.
Everything a blog needs to get found
Prerendering, structured data, feeds, and crawler wiring are built in. You write the posts.
Every post is complete static HTML at build time. Nothing is deferred, and nothing mounts on interaction.
One graph per page. BlogPosting, Person, Organization, WebSite, and FAQPage, joined by @id instead of restated.
Dates come from your content, not from the deploy. Publishing pushes changed URLs to Bing, Yandex, Seznam, and Naver.
A direct answer under every heading, with no links inside it. That paragraph is the chunk a retrieval system lifts.
The blog reads your shadcn tokens and composes your primitives. It installs no theme, no base, and no font.
Fetches your deployed URL as five different bots and names what is missing. Run it from CI.
An SEO and GEO complete blog for Next.js 16, installed into your app as files you own.
Code MIT. Documentation prose CC BY 4.0 . Seed posts CC0.
