---
source: "https://read.technically.dev/p/how-matt-learned-to-ship"
hn_url: "https://news.ycombinator.com/item?id=48400847"
title: "Sentry moved 2,500 page out of their CMS with Claude"
article_title: "How my non-engineering team at Sentry learned to ship"
author: "damowangcy"
captured_at: "2026-06-04T18:55:07Z"
capture_tool: "hn-digest"
hn_id: 48400847
score: 2
comments: 0
posted_at: "2026-06-04T16:20:36Z"
tags:
  - hacker-news
  - translated
---

# Sentry moved 2,500 page out of their CMS with Claude

- HN: [48400847](https://news.ycombinator.com/item?id=48400847)
- Source: [read.technically.dev](https://read.technically.dev/p/how-matt-learned-to-ship)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T16:20:36Z

## Translation

タイトル: Sentry はクロードを使用して CMS から 2,500 ページを移動しました
記事のタイトル: Sentry の非エンジニアリング チームが出荷方法を学んだ方法
説明: Claude Code と Git を使用して CMS から 2,500 ページを移動しました。 4 か月前、私たちのほとんどは発送をしていませんでした。

記事本文:
Sentry の非エンジニアリング チームが出荷方法を学んだ方法
Sentry の非エンジニアリング チームが出荷方法を学んだ方法
Claude Code と Git を使用して、CMS から 2,500 ページを移動しました。 4 か月前、私たちのほとんどはコードを書いたことがありませんでした。
Matt Henderson 2026 年 6 月 4 日 18 1 シェア 数か月前、私は Claude Code を開いてマーケティング サイトにそれを指定しました。このタスクは地味なものでした。サイトをスキャンし、内部リンクが脆弱なページを見つけて、修正を含むプル リクエストをオープンしました。
数分以内に、エージェントはリポジトリをクロールし、複数のページにプル リクエストを一度に送信しました。それから壁にぶつかりました。
コードとして GitHub に存在していたサイトの部分では、エージェントが暴走しました。
少数の AI トークンの 1 つの PR で 18 ページが更新されました。しかし、最も重要なページ (ブログなど) は、エージェントがアクセスできないコンテンツ管理システム (CMS) の背後にロックされていました。
その瞬間から移行が始まりました。 4 か月後、Sentry のマーケティング サイトの約 2,500 ページが Git リポジトリ内の Markdown およびコードとして稼働し、CMS を 100% 活用できるようになりました。これが私たちが学んだことです。
Claude Code セッションの前に、私たちは実際に CMS のフットプリントを拡大していました。私たちは、技術者以外のチーム メンバー (当時は私も含まれていました) のための拡張性、ブランドを維持するためのガードレール、そしてローカリゼーションの基盤を求めていました。エージェントが午後に 1 か月分の SEO 作業を行った後、CMS 境界で停止するのを見て、計算が変わりました。
本当の理由はコストではなく速度でした。私たちは Contentful の請求から逃れようとしたわけではありません (ただし、Cursor の CMS には 1 年間で 57,000 ドル近くの費用がかかり、これは決して簡単ではありませんでした)。私たちは、エージェントがサイトの半分を即座に更新できるが、残りの半分はまったく更新できないという状況から逃れようとしていました。より多くのものを構築するにつれて、その非対称性は週ごとにさらに苦痛になっていきました

クロード コードに関するワークフロー。
同様の動きについて執筆するチームの小さな波があった。Lee Robinson は AI トークンに 260 ドルを費やし、3 日で Cursor.com をヘッドレス CMS から Markdown に戻しました。知事はサイト全体を一から再構築しました。 Vellum の Anita Kirkovska 氏は別の道を歩み、Sanity CMS 内に留まりながら、チームの作成ルールを成文化するagent.md ファイルを備えた独自の AI アシスタントを各マーケターに提供しました。
しかし、ベストプラクティスがまだ書かれている段階にはまだ十分に早いです。明らかだったのは、私たちはコーディング エージェントを使用して Web サイトで作業したいと考えていたため、CMS の背後でページをゲートすることが障害になっていたということです。
技術的にはソフトウェア + AI を実際に意味します。このような記事をもっと購読してください。
当社のグロース エンジニアリング チームは、ここに完全な技術プレイブックを作成しました。これは、Astro、Vite、マーケティング オートメーション フォーム用に構築した Vercel Blob キャッシュ、およびスキーマとフロントマターのマッピング作業について詳しく知りたい場合に適切なリファレンスです。
2.5 人の開発者からなるチームによって 2 か月の間に約 2,500 ページが移行されました。実際のコードの大部分は Claude Code によって書かれました。エンジニアはエージェントの計画、レビュー、指示にほとんどの時間を費やしました。
一度に 1 つのディレクトリ。ドキュメントとマーケティング ページが最初に移行され (テンプレート化が最も簡単)、次にブログ コンテンツ、次に最も複雑なインタラクティブ ページが移行されました。移行中はレガシー コンテンツに対して CMS を無料枠に維持し、ハード カットオーバーではなくソフト カットオーバーを提供しました。
同時にフレームワークもGatsbyからAstroに変更しました。移行自体ではビルド時間は 14 分から 6 分未満に短縮され、Marketo フォーム キャッシュを出荷した後は (フォームはビルドごとに外部 API にアクセスしていました)、4 分未満に短縮されました。これは、約 95 時間のシステム全体で毎日約 15.8 分のビルド時間が節約されたことになります。

ビルド/日。
移行中に、約 200 のオーダーメイド ページを 3 つのテンプレートに統合しました。これにより、コードベースが大幅にシンプルになり、コーディング エージェントが使いやすくなりました。エージェントは 200 の 1 回限りのページの癖を学ぶ必要はありません。 3 つのテンプレートを学習します。
Claude Code は、Contentful の API から新しい Markdown ファイルにコンテンツを取り込む移行スクリプトのほとんどを作成しました。また、Claude が古い Gatsby サイトに対して Playwright のビジュアル回帰テストを実行して、意図しないシフトを検出する実験も行いました。
壊れたステージング ビルドは最大 95% 減少し、Web Vitals スコアは 89 から 97 になりました。壊れたビルドの削減のほとんどは、外部 API の依存関係をビルド プロセスから削除したことによるものです。
以前は Jira チケット、プロジェクト リクエスト、または Looker ダッシュボードが必要だった処理が、Claude Code セッションで行われるようになりました。
過去 4 か月間で、当社の成長チームはこれまでの可能性をはるかに超えるものを出荷しました。
Screaming Frog クローラーによって表面化されたリンクの問題を修正するために、2 つの PR にわたって 147 ページが編集されました。
50 件のレガシー ブログ リダイレクトが約 5 分でクリーンアップされました。以前は、これは誰も取得できなかったバックログ チケットでした。
SEO 価値に基づいてページに優先順位を付ける、私が構築したハブアンドスポーク ツールによって開かれた 116 の内部リンク改善 PR。これは、Claude Code スキルとして実行され、MCP 経由で Ahrefs から権限データを取得します。
製品の発売と A/B テストの結果について、同日に 10 ページが更新または構築されました。誰も更新するサイクルがなかったために、何か月も更新されずに放置されていたコンテンツ。
午後にプロトタイプを作成した新しいソリューション ページを含む、11 件のアクティブな A/B テストが同時に実行されています。これらはそれぞれ、以前は部門横断的なプロジェクトのリクエストでした。
私たちが構築したもののうち 2 つは、他のチームにとって役立つ可能性が最も高い成果物であるため、もう少し詳しく見てみる価値があります。
私たちにはスキル（AI指導fi）があります

ファイル) 各ページ タイプ: ランディング ページ、製品ページ、ソリューション、ブログ投稿、クックブック レシピ、ワークショップ テンプレート。それぞれが、そのページ タイプの構造、SEO 要件、ブランド パターン、コンポーネント ライブラリをエンコードします。
マーケティング リポジトリで作業している人なら誰でも利用できる、/create-* スキルの現在のセット。チームの誰かがクロードに新しいランディング ページの作成を依頼するとき、それは空のファイルから始まるわけではありません。スキルは、必要なフィールド (スラッグ、表示名、SEO メタデータ、ヒーロー コピー) についてインタビューし、ページタイプのスキーマに一致する構造化マークダウン ファイルまたは JSON ファイルを生成し、レビューのために PR を開きます。
動作中のスキル — /create-product-page は、PR を開く前に SEO メタデータ、コピー、アセットをユーザーに説明します。これは、CMS テンプレートが提供していたものと同じ一貫性ですが、エージェントが実際にルールに従う点が異なります (人間はエージェントよりも衒学的指示に従順ではありません)。
数週間前、私たちは Sentry Cookbook を出荷しました。これは、コピー可能なコードを使用して特定の開発者の問題を解決するレシピ集です。
「Sentry Logs を使用して React Native の未定義プロパティをデバッグする」
「Claude Code + Sentry MCP で毎週のパフォーマンス ダイジェストを自動化」
「Sentry で MCP サーバーを監視する」
私たちがこのブログを構築したのは、従来のブログが消費されていくのを目の当たりにしていたからです。長文投稿へのオーガニック検索トラフィックは 2 年間減少していますが、LLM と AI の概要では、構造化された実行可能なコンテンツが引用され続けています。
クックブック形式では、SEO のためのコンテンツとユーザーのためのコンテンツの区別が崩れます。どちらの対象者 (開発者 + LLM) も、構造化され、短く、コードファーストで構成される同じ形の答えを必要としています。
私たちはこのアプローチの初期段階にありますが、初期の収益は良好で、構造は適切であると感じています。
予想以上に大変なことは
いくつかの正直な

スムーズに進まなかったこと：
設計の忠実度はデータの移行よりも時間がかかりました。アニタも Vellum の記事で同じ点を指摘しており、私たちもその点を指摘しました。 Contentful から Markdown へのコンテンツのプルはスクリプト化されたジョブでした。新しい Astro ページを古い Gatsby ページと同じようにピクセルごとにレンダリングするのは、時間がかかる部分でした。 Claude Code は関数型コードの生成が得意です。特定の既存のデザインと 1 対 1 で一致させることは依然として困難です。
ヒューマンエラー。私たちのチームの誰かが、移行中に無害だと考えた電子メール検証ルールをフォームに追加しました。私たちがそれを発見するまで、重要なコンバージョン指標が数日間低下しました。
インフラストラクチャの構築は失敗の原因として過小評価されていました。反復的なビルド中断を修正するために移行の後半に構築した Marketo キャッシュは、プロジェクト全体の中で最も影響力の高い変更の 1 つであることが判明しました。これは、ビルドが高速化されたためではなく (実際にはそうでしたが)、私たちが何年も黙々と取り組んできたある種の失敗を排除したためです。 「必要」になる前に行う価値があります。
チームの学習曲線は実際のものですが、予想よりも小さいです。 4か月前にはコードを書くことができませんでした。私たちのマーケティングチームのほとんどもそれができませんでした。今では私たちの一握りが PR を始めていますが、それは誰かがエンジニアになることを学んだからではなく、環境が十分な流暢さを維持することを強制したからです。チームがこれを行う場合は、最初の 1 か月間は快適と感じるよりも遅くなるように計画してください。
私たちのチームは次に、スケジュールに従って実行できるエージェントという自律型ワークフローに取り組んでいます。
特定の種類のリニア チケットをクロードに直接ルーティングするワークフローは、すでに出荷済みです。チケットが提出され、エージェントがチケットを受け取り、PR が審査のために現れます。
これまでのところ、小規模で広範囲にわたる修正 (リンク切れ、タイプミスの修正、宣伝文の更新) に最適ですが、まだ状況は変わりません。

どのタスクがエージェント対応で、どのタスクがすべてのステップでループに参加する必要があるかを調べます。
私たちが投資しているもう 1 つのことは、AI コード レビューです。エージェントが本番環境に触れる場合の最大のリスクは、品質管理が PR をレビューする担当者に委ねられ、レビュー担当者が 50 件の変更をまとめて見落とすことです。
私たちはまだ表面をなぞったばかりです。移行自体は完了しており、その速度は驚くべきものであり、上記の簡単な例はすでに、前年に出荷したものを超えています。
しかし、私たちは毎週、SaaS からコードベースに移行できる別のワークフロー、Claude Code スキルになる可能性のある別のレポート、スケジュールされたループになる可能性のある別の繰り返しタスクを見つけています。時間が経つにつれて、このツールの構築に時間を費やすことから、構築したツールを単に使用することに移行すると思います。
Sentry の移行に関する Eli の技術プレイブック (Astro、Vite、Markdown、Marketo キャッシュ)
Lee Robinson による Cursor.com への移行に関する記事
Sanity + Agent.md を使用した Anita Kirkovska の CMS 内アプローチ
上記のハブアンドスポーク SEO ツール
18 1 シェア 前 この投稿に関するディスカッション コメント Restacks Matt Henderson 4h 著者 これは楽しいコラボでした！興味のある方のために、他のことも書いています: https://hendersonmatthew.substack.com/
返信 シェア トップ 最新のディスカッション 投稿はありません

## Original Extract

We moved 2,500 pages out of our CMS with Claude Code and Git. 4 months ago, most of us hadn't shipped.

How my non-engineering team at Sentry learned to ship
Subscribe Sign in How my non-engineering team at Sentry learned to ship
We moved 2,500 pages out of our CMS with Claude Code and Git. 4 months ago, most of us hadn't written code.
Matt Henderson Jun 04, 2026 18 1 Share A few months ago, I opened Claude Code and pointed it at our marketing site. The task was unglamorous: scan the site, find pages where internal linking was weak, and open pull requests with fixes.
Within minutes, the agent was crawling the repository and shipping pull requests across multiple pages at once. Then it hit a wall.
For the parts of our site that lived in GitHub as code, the agent ran wild:
18 pages updated in a single PR for a handful of AI tokens. But the pages that mattered most (like our blog), were locked behind a Content Management System (CMS) the agent couldn’t touch.
That moment kicked off a migration . Four months later, ~2500 pages of Sentry ‘s marketing site live as Markdown and code in a Git repo, and we’re 100% out of our CMS. Here is what we learned.
Before the Claude Code session, we were actually expanding our CMS footprint. We wanted scalability for non-technical team members (me included at the time), guardrails to keep things on-brand, and a foundation for localization. Watching an agent do a month of SEO work in an afternoon and then stop dead at the CMS boundary changed the calculation.
Velocity was the real reason, not cost. We weren’t trying to escape a Contentful bill (though Cursor’s CMS had cost them nearly $57,000 in one year , which is non-trivial). We were trying to escape the situation where an agent could update half of our site instantly and the other half not at all. That asymmetry was getting more painful by the week as we built more workflows around Claude Code.
There’s been a small wave of teams writing about similar moves: Lee Robinson migrated cursor.com from a headless CMS back to Markdown in three days, spending $260 in AI tokens. Prefect rebuilt their entire site from scratch. Anita Kirkovska at Vellum took a different path , staying inside Sanity CMS but giving each marketer their own AI assistant with agent.md files that codify the team’s writing rules.
But it’s still early enough that best practices are still being written. What was clear, was that we wanted to work on the website w/ coding agents, and having pages gated behind a CMS was a blocker.
Technically makes practical sense of software + AI. Subscribe for more stories like this one.
Our growth engineering team wrote up the full technical playbook here, which is the right reference if you want the deep dive on Astro, Vite, the Vercel Blob cache we built for our marketing automation forms, and the schema -to-frontmatter mapping work.
~2500 pages migrated by a team of 2.5 developers in a two-month window. Most of the actual code was written by Claude Code; the engineers spent most of their time planning, reviewing, and directing the agent.
One directory at a time. Docs and marketing pages migrated first (easiest to template), then blog content, then the most complex interactive pages. We kept the CMS on a free tier for legacy content during the transition, which gave us a soft cutover instead of a hard one.
We changed frameworks at the same time, from Gatsby to Astro. Build times went from 14 minutes to under 6 in the migration itself, and after we shipped a Marketo form cache (the forms had been hitting an external API on every build), we got down to under 4. That works out to about 15.8 build hours saved daily across our ~95 builds/day.
We consolidated ~200 bespoke pages into 3 templates during the migration. This made the codebase dramatically simpler, which makes it easier to use by coding agents. The agent doesn’t have to learn the quirks of 200 one-off pages; it learns three templates.
Claude Code wrote most of the migration scripts that pulled content from Contentful’s API into our new Markdown files. We also experimented with Claude running Playwright visual regression tests against the old Gatsby site to catch unintended shifts.
Broken staging builds are down ~95% and our Web Vitals score went from 89 to 97. Most of the broken-build reduction came from cutting external API dependencies out of the build process.
Things that used to require a Jira ticket, a project request, or a Looker dashboard now happen in a Claude Code session.
In the last four months, our growth team has shipped a lot more than previously possible:
147 pages edited across 2 PRs to fix link issues surfaced by our Screaming Frog crawler.
50 legacy blog redirects cleaned up in about 5 minutes. This used to be a backlog ticket that nobody got to.
116 internal-link improvement PRs opened by a hub-and-spoke tool I built that prioritizes pages by SEO value. It runs as a Claude Code skill and pulls authority data from Ahrefs via MCP.
10 pages updated or built same-day for product launches and A/B test outcomes. Content that used to sit stale for months because nobody had cycles to update it.
11 active A/B tests running concurrently , including a new solutions page I prototyped in an afternoon. Each of these would have been a cross-functional project request before.
Two of the things we built deserve a longer look because they’re the artifacts most likely to be useful to other teams.
We have Skills (AI instruction files) for each page type: landing pages, product pages, solutions, blog posts, cookbook recipes, workshop templates. Each one encodes the structure, the SEO requirements, the brand patterns, and the component library for that page type.
The current set of /create-* Skills available to anyone working in our marketing repo. When anyone on the team asks Claude to create a new landing page, it doesn’t start from a blank file. The Skill interviews them for the required fields (slug, display name, SEO metadata , hero copy), produces a structured Markdown or JSON file that matches the page-type schema, and opens a PR for review.
A Skill in motion — /create-product-page walks the user through SEO metadata, copy, and assets before opening a PR. This is the same consistency our CMS templates used to give us, except the agent actually follows the rules (humans are less observant than agents to pedantic instructions).
A few weeks ago we shipped the Sentry Cookbook , a collection of recipes that solve specific developer problems with copyable code.
“Debug undefined properties in React Native with Sentry Logs”
“Automate weekly performance digests with Claude Code + Sentry MCP”
“Monitor your MCP server with Sentry”
We built it because of how we were watching our traditional blog get consumed: organic search traffic to long-form posts has been falling for two years, but LLMs and AI Overviews keep citing structured, executable content.
The cookbook format collapses the distinction between content for SEO and content for users. Both audiences (developers + LLMs) need the same shaped answer: structured, short, and code-first.
We’re early on with this approach, but early returns are looking good, and the structure feels right.
What’s harder than we expected
A few honest things that didn’t go smoothly:
Design fidelity took longer than the data migration. Anita made the same point in her Vellum writeup, and we hit it too. Pulling content from Contentful into Markdown was a scripted job. Getting the new Astro pages to render the same as the old Gatsby pages, pixel for pixel, was the slow part. Claude Code is good at producing functional code; matching a specific existing design 1:1 is still hard.
Human Error. Someone on our team added an email validation rule to a form during the migration that they thought was innocuous. It tanked a key conversion metric for a few days before we caught it.
Build infrastructure was an under-appreciated source of failure. The Marketo cache we built late in the migration to fix recurring build breaks turned out to be one of the highest-leverage changes of the whole project, not because it sped up builds (though it did), but because it eliminated a class of failure we’d been quietly working around for years. Worth doing before you “need” to.
The team learning curve is real but smaller than expected. I couldn’t write code four months ago. Most of our marketing team couldn’t either. A handful of us are opening PRs now, not because anyone learned to be an engineer, but because the environment forced enough fluency to keep up. If your team is going to do this, plan for the first month being slower than feels comfortable.
Next on our team we’re tackling autonomous workflows: agents that can run on a schedule.
We’ve shipped one already, a workflow that routes specific kinds of Linear tickets directly to Claude. A ticket gets filed, an agent picks it up, a PR shows up for review.
So far it works best for small, well-scoped fixes (broken links, typo corrections, blurb updates), but we’re still figuring out which tasks are agent-ready and which need a human in the loop on every step.
The other thing we’re investing in is AI code review. The biggest risk with letting agents touch production is that quality control falls on whoever’s reviewing PRs, and reviewers miss things in a batch of 50 changes.
We’re still just scratching the surface. The migration itself is done, the velocity has been surprising, and the quick examples above are already more than we shipped in the year before this.
But every week we’re finding another workflow that could be moved off of SaaS and onto our codebase, another report that could be a Claude Code skill, another recurring task that could become a scheduled loop. Over time, I assume we’ll shift from spending time building this tooling to just using the tools we’ve built.
Eli’s technical playbook on the Sentry migration (Astro, Vite, Markdown, the Marketo cache)
Lee Robinson’s migration writeup for cursor.com
Anita Kirkovska’s in-CMS approach using Sanity + agent.md
The hub-and-spoke SEO tool referenced above
18 1 Share Previous Discussion about this post Comments Restacks Matt Henderson 4h Author This was a fun collab! For anyone curious, I write other things too: https://hendersonmatthew.substack.com/
Reply Share Top Latest Discussions No posts
