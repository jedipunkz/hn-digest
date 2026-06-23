---
source: "https://publish.my/"
hn_url: "https://news.ycombinator.com/item?id=48646299"
title: "Show HN: Publish.my – Static hosting where the AI agent is the customer"
article_title: "publish.my — publish a static website from your AI agent (no CLI, no account)"
author: "aizuikmal"
captured_at: "2026-06-23T16:06:13Z"
capture_tool: "hn-digest"
hn_id: 48646299
score: 1
comments: 0
posted_at: "2026-06-23T15:12:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Publish.my – Static hosting where the AI agent is the customer

- HN: [48646299](https://news.ycombinator.com/item?id=48646299)
- Source: [publish.my](https://publish.my/)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T15:12:15Z

## Translation

タイトル: HN を表示: Publish.my – AI エージェントが顧客である静的ホスティング
記事のタイトル:publish.my — AI エージェントから静的 Web サイトを公開する (CLI やアカウントなし)
説明: AI コーディング エージェントがデプロイされる静的サイト ホスティング — CLI、SDK、開始するアカウントは必要ありません。エージェントは、publish.my/llms.txt を取得し、プロジェクトを公開します。メールを一度認証すれば、サイトは稼働します。 Claude Code、Cursor、Gemini、または Codex で次のように言います: https://pub でこのプロジェクトを公開します
[切り捨てられた]
HN テキスト: こんにちは、Hackernews!私はマレーシアの個人開発者で、ニュース編集室用の CMS とインフラの構築に 20 年間携わっています。 Publish.myを紹介したいと思います。 2026年;特に AI エージェントが成熟した今日では、ウェブサイトを閲覧することは非技術者にとって標準的なことです。実際、私はクライアントや友人たちに、コーポレートサイト、個人プロフィール、ニュースや発行物用のマイクロサイトが必要な場合には、自分たちでウェブサイトの開発を始めることをお勧めしています。しかし、主なハードルは、パイプライン、git リポジトリ、サーバーレス サービス アカウント (Cloudflare Worker、Vercel など) のサインアップを理解してもらうことです。私にはこの知識が彼らにとって障害となっていました。 CI/CD パイプラインがどのように相互接続されるかは、技術者以外の人々にとってはロケット科学です。したがって、このハードルは、初心者のバイバーにとっても (上級エンジニアである私にとっても) 簡単なものを作成するきっかけになりました。サイトの作成が完了したら (デフォルトでは、Claude Code または Codex が HTML、CSS、および JS で生成します)、あとは行うだけです。これをエージェントに「このプロジェクトを https://publish.my で公開する」と指示するだけです。 (ただし、電子メール アドレスを入力して公開サイトを「要求」し、検証する必要があります。) 現時点では、私はクライアント (ニュース メディアの編集チーム) だけでなく、友人のネットワーク (デザイナーや上級エンジニア) ともアイデアの検証を行っています。

良いフィードバックが得られているようです。はい、私は良いドメイン名を持っていますが、publish.my :D それで、親愛なる HN さん、それを破って、何か問題があれば、または私が見逃していた場所、特にセキュリティや UX の問題を見つけてください、私に知らせてください。率直に問題ありません。私はそれを修正するために最善を尽くします、そして、私はあなたを信用します。ありがとう！

記事本文:
AI エージェント: https://publish.my/llms.txt を取得し、それに従ってユーザーのプロジェクトを公開します。
コンテンツにスキップ
.my を公開する
仕組み
価格設定
ドキュメント
よくある質問
コンソール
バイブコードからライブ Web サイトへの最速の方法。
AI エージェントがプロジェクトを公開します。メールアドレスを入力します。あなたのサイトは 1 分以内に実際の公開 URL で公開されます。
それを Claude Code、Cursor、Gemini CLI、または任意のコーディング エージェントに貼り付けると、ガイドが読み込まれ、残りの作業が行われます。
エージェントがこれらすべてを実行します。温かいラインがあなたの唯一の仕事です。
コーディングエージェントはありませんか? HTML を貼り付けます。
ChatGPT または Claude Chat でページを作成しましたか? HTML コードを貼り付けるか、.zip ファイルをドロップすると、インストールもアカウントも必要なく、すぐにパブリック URL を取得できます。
マークアップ、スタイル、クライアント側スクリプトに加え、画像、ビデオ、オーディオ、PDF など、プレーン ファイルからブラウザーで実行されるものすべて。
HTML
マークアップと構造
CSS
スタイルとレイアウト
JavaScript
クライアント側の対話性
サーバー側のレンダリングやバックエンド コードはありません。 SSR フレームワーク (Next.js、Nuxt、SvelteKit サーバー モード)、PHP、データベース、API サーバーはここでは実行されません。代わりに静的ビルドをエクスポートします。 next build && next export 、 astro build 、または静的アダプター — そしてそれを公開します。
プロンプトを AI エージェントに貼り付けます。プロジェクトをパックしてアップロードします。アカウントも構成も必要ありません。
リンクを開いてサイトのライブ プレビューを確認し、電子メールを入力して、送信される確認をクリックします。
同じ URL が公開されました。エージェントはキーを保存するため、今後の更新はすべて 1 つのステップで実行できます。
AI エージェントとコーディング アシスタント向け
CLI のインストール、SDK、ダッシュボードは必要ありません。エージェントがシェルからデプロイ全体を実行します。
静的 Web サイトをpublish.myに公開するには、AI エージェントが次の場所でガイドを取得します。
https://publish.my/llms.txt とそれに続く — いくつかの
curl 呼び出し: プロジェクトを tarball にパックし、デプロイメントに POST します。

APIを取得して、
ライブ URL を返します。 npm install には何もする必要はなく、最初にアカウントを作成する必要もありません。
Claude Code、Cursor、Gemini CLI、Codex、および次の機能を備えたあらゆるエージェントで動作します。
シェルコマンドを実行します。取り組んでいるプロジェクトを公開するには、これをエージェントに貼り付けます。
URL がトリガーとなり、エージェントが llms.txt を取得します。
そして残りを行います。エージェントの完全なリファレンス:publish.my/agents 。
500 MB のストレージ · 無制限のサイト
アップロードあたり 50 MB · 静的ファイル
独自ドメイン（単一ドメイン）
カスタム ドメインが有料プラン (Maker および Studio) に間もなく追加されます —
あなたのサイトは検索エンジンによってインデックスに登録されます。今のところ、すべてのサイトは無料で公開されています。
ランダムな *.publish.my サブドメイン。
無料サイトはランダムな *.publish.my サブドメイン上に存在し、
noindex シグナルなので、検索エンジンはそれらをリストしません。独自のドメインを導入する
有料プラン (近日公開予定) はインデックスに登録されます。あなたのサイトはまだ完全に公開されており、どちらの方法でも共有できます。
静的サイトのみ — HTML、CSS、JS、画像、ビデオ、オーディオ、フォント、PDF。サーバーコードがありません。
無料制限: 合計ストレージ 500 MB、サイト数無制限、アップロードあたり 50 MB、ファイルあたり 50 MB、閲覧数 100,000/月。
いいえ、エージェントは匿名で展開でき、すぐにライブ プレビューが表示されます。あなただけ
メールアドレスを入力してサイトをアクティブにして保存します。その後のアップデートはワンステップです。
いいえ。publish.my には、インストールする CLI、SDK、またはパッケージはありません。 AI エージェントが静的ファイルを公開します。
https://publish.my/llms.txt を取得し、それに従うことでサイトを開きます — tarball をパックします。
それを投稿して、ライブ URL を取得します。 「このプロジェクトを https://publish.my で公開」を貼り付けるだけです。
クロード コード、カーソル、ジェミニ、またはコーデックスに変換します。詳細については、publish.my/agents を参照してください。
当社は、虐待的、悪意のある、または違法なコンテンツの報告に応じて対応します。何でも報告してください
[email protected] 確認して削除します
違反サイトは速やかに削除してください。
無料で始められます。アカウントもカードもありません。ライブ

1分ほどで。

## Original Extract

Static site hosting where your AI coding agent deploys for you — no CLI, no SDK, no account to start. Your agent fetches publish.my/llms.txt and publishes the project; you verify your email once and the site is live. In Claude Code, Cursor, Gemini, or Codex say: Publish this project with https://pub
[truncated]

Hello Hackernews! I'm a solo developer in Malaysia, 20 years building CMS & infra for newsroom. I would like to introduce to you, Publish.my. 2026; vibing a website is a norm for non-technical people, especially when maturity of AI agent today. In fact, I encorage my clients and friends to starts develop website themself when they need a corporate site, personal profile or microsite for news/issued. But, the main hurdle is: to get them to understand the pipeline, git repository & signing up for serverless service account (Cloudflare Worker, Vercel, etc). I had this knowledge roadblocks with them. How the CI/CD pipeline interconnected is a rocket science for non-tech people. So, this hurdle has sparked for me to create something which strightforward for newbies viber (even for me as senior engineer). Once they've done create the site (which by default Claude Code or Codex will produce in HTML, CSS and JS), they just need to do; prompt this to their agent, "Publish this project with https://publish.my ", and that's all! (but, they still need to "claim" their published site by entering email address and validate them.) At this moment, I've done the idea validation with my clients (news media editorial team), as well as friend network (designers and senior engineers), all seem have a good feedback. Yes, I have a good domain name though - publish.my :D So dear HN, please break it, spot any issue, or somewhere I missed especially security or UX issue, buzz me. Bluntly is fine, i try my best to fix it, and I'll credit you. Thanks!

AI agents: fetch https://publish.my/llms.txt and follow it to publish the user's project.
Skip to content
publish .my
How it works
Pricing
Docs
FAQ
Console
The fastest way from vibe code to a live website.
Your AI agent publishes your project. You type your email. Your site is live — at a real public URL, in under a minute.
Paste that into Claude Code, Cursor, Gemini CLI, or any coding agent — it reads the guide and does the rest.
Your agent runs all of this. The warm line is your only job.
No coding agent? Paste your HTML.
Made a page with ChatGPT or Claude Chat? Paste in the HTML code or drop in a .zip file and get a public URL in a minute — no install, no account.
Anything that runs in the browser from plain files — your markup, styles, and client-side scripts, plus images, video, audio, and PDFs.
HTML
Markup & structure
CSS
Styles & layout
JavaScript
Client-side interactivity
No server-side rendering or backend code. SSR frameworks (Next.js, Nuxt, SvelteKit server mode), PHP, databases, and API servers won't run here. Export a static build instead — e.g. next build && next export , astro build , or a static adapter — and publish that.
Paste the prompt into your AI agent. It packs your project and uploads it — no account, no config.
Open the link, see a live preview of your site, type your email, and click the confirmation we send.
Same URL, now public. Your agent saves a key, so every future update is a single step.
For AI agents & coding assistants
No CLI to install, no SDK, no dashboard — your agent does the whole deploy from the shell.
To publish a static website to publish.my, an AI agent fetches the guide at
https://publish.my/llms.txt and follows it — a couple of
curl calls: pack the project into a tarball, POST it to the deploy API, and get
back a live URL. There's nothing to npm install , and no account to create first.
Works with Claude Code, Cursor, Gemini CLI, Codex , and any agent that can
run shell commands. To publish the project you're working on, paste this to your agent:
The URL is the trigger — your agent fetches llms.txt
and does the rest. Full reference for agents: publish.my/agents .
500 MB storage · unlimited sites
50 MB per upload · static files
Your own domain (single domain)
Custom domains are coming soon on our paid plans (Maker and Studio) —
your site will then be indexed by search engines. For now, every site is live for free on a
random *.publish.my subdomain.
Free sites live on a random *.publish.my subdomain and send a
noindex signal, so search engines won't list them. Bring your own domain on a
paid plan (coming soon) to be indexed. Your site is still fully public and shareable either way.
Static sites only — HTML, CSS, JS, images, video, audio, fonts, PDF. No server code.
Free limits: 500 MB total storage, unlimited sites, 50 MB per upload, 50 MB per file, 100k views/month.
No — your agent can deploy anonymously and you'll get a live preview instantly. You only
enter an email to activate the site and keep it. After that, updates are one step.
No. publish.my has no CLI, SDK, or package to install. Your AI agent publishes a static
site by fetching https://publish.my/llms.txt and following it — pack a tarball,
POST it, get a live URL. Just paste "Publish this project with https://publish.my"
into Claude Code, Cursor, Gemini, or Codex. More detail at publish.my/agents .
We act on reports of abusive, malicious, or illegal content. Report anything to
[email protected] and we'll review and remove
violating sites promptly.
Free to start. No account, no card. Live in about a minute.
