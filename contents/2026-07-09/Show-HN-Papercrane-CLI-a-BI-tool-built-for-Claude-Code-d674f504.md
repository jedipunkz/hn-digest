---
source: "https://papercrane.ai/blog/today-im-launching-papercrane-cli-a-bi-tool-built-for-claude-code"
hn_url: "https://news.ycombinator.com/item?id=48847623"
title: "Show HN: Papercrane-CLI – a BI tool built for Claude Code"
article_title: "Today I’m launching papercrane-cli: a BI tool built for Claude Code | Papercrane Blog"
author: "timliu9"
captured_at: "2026-07-09T16:23:13Z"
capture_tool: "hn-digest"
hn_id: 48847623
score: 5
comments: 0
posted_at: "2026-07-09T15:33:45Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Papercrane-CLI – a BI tool built for Claude Code

- HN: [48847623](https://news.ycombinator.com/item?id=48847623)
- Source: [papercrane.ai](https://papercrane.ai/blog/today-im-launching-papercrane-cli-a-bi-tool-built-for-claude-code)
- Score: 5
- Comments: 0
- Posted: 2026-07-09T15:33:45Z

## Translation

タイトル: Show HN: Papercrane-CLI – Claude Code 用に構築された BI ツール
記事のタイトル: 今日は、Papercrane-cli を起動します: Claude Code 用に構築された BI ツール |折り鶴のブログ
説明: すべての分析ベンダーは、自社のダッシュボードに AI 副操縦士をボルトで取り付けています。 Power BI にはそれがあります。 Tableau には 1 つあります。それらはすべて同じように機能します。人間が古いものを運転します。

記事本文:
今日は、Papercrane-cli を起動します。これは、Claude Code 用に構築された BI ツールです |折り鶴のブログ
価格設定 プラットフォーム統合 比較 ブログ ログイン サインアップ すべての投稿 今日は、Claude Code 用に構築された BI ツール、papercrane-cli を起動します
すべての分析ベンダーは、自社のダッシュボードに AI 副操縦士をボルトで取り付けています。 Power BI にはそれがあります。 Tableau には 1 つあります。これらはすべて同じように機能します。人間が古いツールを操作し、AI がサイドバーに座って提案を行います。
私は反対方向に行きました。すでに使用されているエージェントを使用して、それを中心に BI ツールを構築しました。
Papercrane-cli は、ビジネス上の質問をライブ ダッシュボードに変えるために必要なものすべてを Claude Code に提供するコマンド ライン ツールです。つまり、データへの認証されたアクセス、実際のダッシュボード ワークスペース、誰にでも送信できるリンクへのパスです。今日から発売です。
これは無料で、すでに支払った Claude サブスクリプションで実行されます。
エージェントはすでにダッシュボードに優れています。これが彼らを止めるものです。
クロード コードを使用している場合は、おそらくこれを試したことがあるでしょう。収益ダッシュボードを要求すると、本当に優れた React が作成されます。チャート、レイアウト、作品。すると 3 つの問題に遭遇します。私はそれらをすべて実行したのでわかります。papercrane-cli は私が反対側で望んでいたツールです。
1. エージェントが構築するものはすべて時間内に固定されます
今すぐエージェントにダッシュボードを依頼すると、HTML ファイル、PDF、チャット アーティファクトの 3 つのうちのいずれかを入手できます。この 3 つには、木曜日まで気付かないような欠陥が共通しています。それは、数字がビルド時に焼き付けられていたということです。それはあなたのビジネスの写真であり、撮影された瞬間は正確ですが、その後は日々真実から遠ざかっていきます。
MCP ではこの問題は解決されません。その理由を正確に把握する価値があります。 MCP はその機能が優れており、エージェントが会話中にツールを呼び出すことができます。問題は、接続が

チャットセッションを待ち望んでいます。エージェントが作成したダッシュボードは継承できません。したがって、エージェントはデータをクエリし、結果をハードコードして、ダッシュボードのように見え、スクリーンショットのように動作するものを渡します。
Papercrane は、ダッシュボード自体の内部にデータ接続を配置します。すべてのダッシュボードは小さな Next.js アプリであり、そのサーバー コードは、エージェントがビルドに使用したものと同じ認証済み API を呼び出します。チャットが終了します。数字は動き続けます。この違いは構造的なものです。下に常駐するデータ層がなければ、いくらプロンプトを出してもそこに到達することはできません。
2. 認証は面倒で移植性がない
あなたの番号はサインイン画面の背後にあります: GA4、HubSpot、Salesforce、QuickBooks。エージェントはこれらの画面をクリックして移動することができず、回避策も醜いです。技術者がすべてのツール (承認キューとセキュリティ レビューを備えた数週間で測定されるプロジェクト) のカスタム開発者接続を構築するか、ターミナルやチャット ウィンドウに秘密キーを貼り付けて、誰も上にスクロールしないことを期待するかのどちらかです。
接続作業は全員に一度行いました。ツールの接続はサインインするだけです。ブラウザーで通常のサインイン ページが開き、2 つのアプリを接続するときと同じようにアクセスを承認すれば完了です。エージェントは、GA4、Stripe、HubSpot、Postgres、BigQuery、および 50 以上のウェアハウス、広告プラットフォーム、CRM、財務、製品分析にアクセスするための 1 つのクリーンで安全な方法を利用できます。資格情報は Papercrane で暗号化され、ラップトップやチャット履歴から外されたままになります。サービスがアクセスを更新する必要がある場合、バックグラウンドで処理されます。
3. ダッシュボードは誰かに送るものです
ダッシュボードは、チーム、CEO、クライアントに送信した瞬間に価値を獲得します。これは、すべてのローカル プロトタイプが消滅するステップです。準備ができたら、PaperCrane パブリッシュによりダッシュボードがリンクに配置されます

アクセス制御付きで共有すると、受信者にも番号が公開されます。クラウド側は、ローカル ツールでは決して処理できないこと (ホスティング、埋め込み、カスタム ドメイン) を処理します。
ここまではデータの流入について説明しましたが、流出する可能性もあります。ダッシュボードには、Salesforce の取引を次のステージに移動するボタン、グラフの背後にあるレコードを更新するコントロールなどのアクションを含めることができます。数値を読み取るのと同じ認証された接続は、数値を書き戻すことができます。これが、ダッシュボードという言葉が持っているものを過小評価し始めるポイントです。行動できるレポートはツールであり、BI がそのようなものを提供したことはありません。
すべての資格情報を 1 か所に
今日、人々がエージェントをデータに接続する方法には、もっと地味な問題があります。資格情報があらゆる場所に流出してしまうのです。ラップトップ上の環境変数内の API キー。もう 1 つは MCP 構成内にあります。 1 つは 3 週間前にプロンプ​​トに貼り付けられたもので、すでに忘れていました。これに、エージェントが同じシステムにアクセスすることを望んでいるすべてのチームメイトを掛け合わせると、誰が何にアクセスできるのか、そしてどうやってアクセスをオフにするのかという重要な 2 つの質問に誰も答えることができなくなります。
Papercrane を使用すると、組織は各ソースを 1 回接続します。資格情報は暗号化されて 1 か所に保持され、エージェントは 1 つの認証された API を通じてデータにアクセスし、接続は監査ログに記録され、アクセスは一元的に取り消されます。チームメイトのエージェントは、Slack 経由でキーを転送することなく、あなたのエージェントと同じリーチを得ることができます。それが問題になり始めるのは、これがおもちゃでなくなった瞬間です。ダッシュボードは本物であり、データは機密であり、複数の人のエージェントがダッシュボードに触れているのです。
コネクタが存在しない場合、エージェントがコネクタを書き込みます
すべての統合は TypeScript ハンドラーであり、CLI にはエージェント自体用に書かれたガイドが含まれています。papercrane local-integration-guide は、直接アクセスする方法全体を出力します。

tdout。データが当社がカバーしていない場所に存在する場合は、エージェントにコネクタを構築するように伝えてください。これはワークスペースに配置され、ローカルで実行され、ホストされているものとまったく同じように呼び出すことができます。ツールセットはそれ自体を拡張します。
すべてのダッシュボードは、ワークスペースにある Next.js と React ソース ( page.tsx 、サーバー アクション、recharts、Tailwind) です。 1 つのコマンドでそれを独自の GitHub リポジトリに公開します。明日Papercraneを辞めても、ソフトウェアは使い続けることになります。
人間のためのツールではなく、AIのためのツール
Papercrane-cli におけるすべての設計上の決定は、エージェントに何が必要かという 1 つの質問から導き出されます。ドキュメントは、モデルが読み取ることができるマークダウンとして提供されます。すべてのコマンドは、モデルが解析できる出力を表示します。ここにはマニフェスト仕様やプロトコルはなく、ただ無慈悲な慣例があるだけです。クロード コードがそれを読み取って実行できれば、それは機能します。 Codex、Cursor、その他 stdout を読み取るものはすべて機能します。
これが、セットアップが 1 つのプロンプトである理由です。クロード コードを開き、次のように入力します。
https://papercrane.ai/get-started.md を読んでログインしてください
エージェントはガイドを読み、CLI をインストールし、ブラウザを開いてサインインし、最初のソースに接続します。そこからは、見たいものを尋ねるだけです
チャットしている場所でも構築が行われます。 Papercrane-cli は Claude Code のプレビューに接続されているため、エージェントが作業している間、ダッシュボードは Claude アプリ内でレンダリングされます。グラフを構築する会話の横で、グラフが形になっていく様子を観察します。
CLI は無料です。ダッシュボードは独自の Claude または OpenAI サブスクリプションで実行されるため、AI クレジットを購入したり、シートごとの価格を交渉したりする必要はありません。無料利用枠には、5 つのダッシュボードとすべてのコネクタが含まれます。有料レベルでは、必要に応じてチーム機能、カスタム ドメイン、組み込み分析を追加します。
BI ツールは 30 年かけて人間のソフトウェア操作を向上させてきました。その時代は終わりつつあります。次の BI ツールは

クエリ ビルダーは、1 秒あたり 1,000 ワードの速さでドキュメントを読み、クリックすることに飽きることがないため、よりフレンドリーなクエリ ビルダーは必要ありません。エージェントに必要なもの、つまり資格情報、ワークスペース、配送方法が必要です。それが私たちが築いてきたものです。
Claude Code がインストールされている場合は、90 秒で実際のデータのライブ ダッシュボードが表示されます。プロンプトを貼り付けて、何が起こるかを見て、作成した最初のダッシュボードを私に送ってください。全部読みました。
npm install -g @papercraneai/cli 、またはエージェントに次のように伝えてください: https://papercrane.ai/get-started.md を読んでログインしてください
2026年の折り鶴。無断転載を禁じます。

## Original Extract

Every analytics vendor is bolting an AI copilot onto their dashboards. Power BI has one. Tableau has one. They all work the same way: the human drives the old t

Today I’m launching papercrane-cli: a BI tool built for Claude Code | Papercrane Blog
Pricing Platform Integrations Compare Blog Log in Sign Up All posts Today I’m launching papercrane-cli: a BI tool built for Claude Code
Every analytics vendor is bolting an AI copilot onto their dashboards. Power BI has one. Tableau has one. They all work the same way: the human drives the old tool, and the AI sits in a sidebar making suggestions.
I went the opposite direction. I took the agent you already use and built the BI tool around it.
papercrane-cli is a command line tool that gives Claude Code everything it needs to turn a business question into a live dashboard: authenticated access to your data, a real dashboard workspace, and a path to a link you can send to anyone. It launches today.
It’s free and it runs on the Claude subscription you already pay for.
Agents are already great at dashboards. Here’s what stops them.
If you use Claude Code, you’ve probably tried this. You ask it for a revenue dashboard and it writes genuinely good React. Charts, layout, the works. Then it runs into three problems. I know because I hit every one of them, and papercrane-cli is the tool I wanted on the other side.
1. Everything an agent builds is frozen in time
Ask an agent for a dashboard right now and you get one of three things: an HTML file, a PDF, or a chat artifact. All three share a flaw you won’t notice until Thursday: the numbers were baked in at build time. It’s a photograph of your business, accurate the moment it was taken and drifting further from the truth every day after.
MCP doesn’t fix this, and it’s worth being precise about why. MCP is good at what it does: it lets your agent call tools during a conversation. The catch is that the connection belongs to the chat session. The dashboard your agent writes can’t inherit it. So the agent queries your data, hard codes the results, and hands you something that looks like a dashboard and behaves like a screenshot.
Papercrane puts the data connection inside the dashboard itself. Every dashboard is a small Next.js app whose server code calls the same authenticated API your agent used to build it. The chat ends. The numbers keep moving. That difference is structural: no amount of prompting gets you there without a standing data layer underneath.
2. Authentication is messy and isn’t portable
Your numbers live behind sign in screens: GA4, HubSpot, Salesforce, QuickBooks. An agent can’t click through those screens, and the workarounds are ugly. Either someone technical builds a custom developer connection for every tool (a project measured in weeks, with approval queues and security reviews), or you start pasting secret keys into terminals and chat windows and hoping nobody scrolls up.
We did the connection work once, for everyone. Connecting a tool is just signing in: your browser opens a normal sign in page, you approve access the way you’d connect any two apps, done. Your agent gets one clean, secure way to reach GA4, Stripe, HubSpot, Postgres, BigQuery, and 50+ more: warehouses, ad platforms, CRMs, finance, product analytics. The credentials stay encrypted with Papercrane, off your laptop and out of your chat history, and when a service needs its access renewed behind the scenes, that’s handled for you.
3. A dashboard is something you send to someone
A dashboard earns its keep the moment you send it: to your team, your CEO, a client. That’s the step where every local prototype dies. When you’re ready, papercrane publish puts your dashboard on a link you can share, with access control, and the numbers are live for the recipient too. The cloud side handles what a local tool never will: hosting, embedding, custom domains.
Everything so far describes data flowing in. It can flow out too. A dashboard can carry actions: a button that moves a deal to its next stage in Salesforce, a control that updates the record behind a chart. The same authenticated connection that reads the numbers can write them back. That’s the point where the word dashboard starts to undersell what you’ve got: a report you can act from is a tool, and BI has never shipped one of those.
All of your credentials, in one place
There’s a quieter problem with how people wire agents to data today: the credentials end up everywhere. An API key in an env var on your laptop. Another in an MCP config. One pasted into a prompt three weeks ago that you’ve already forgotten about. Multiply that by every teammate who wants their agent reaching the same systems, and nobody can answer the two questions that matter: who has access to what, and how do we turn it off?
With Papercrane, your organization connects each source once. Credentials are encrypted and held in one place, agents reach data through one authenticated API, connections are audit logged, and access gets revoked centrally. A teammate’s agent gets the same reach as yours without anyone forwarding keys over Slack. That starts to matter at exactly the moment this stops being a toy: the dashboard is real, the data is sensitive, and more than one person’s agent is touching it.
When a connector doesn’t exist, your agent writes one
Every integration is a TypeScript handler, and the CLI includes a guide written for the agent itself: papercrane local-integration-guide prints the whole how to straight to stdout. If your data lives somewhere we don’t cover, tell your agent to build the connector. It lands in your workspace, runs locally, and is callable exactly like the hosted ones. The toolset extends itself.
Every dashboard is Next.js and React source sitting in your workspace: a page.tsx , a server action, recharts, Tailwind. Publish it to your own GitHub repo with one command. If you leave Papercrane tomorrow, you keep working software.
A tool for AI, not for a person
Every design decision in papercrane-cli follows from one question: what does the agent need? The docs are served as markdown a model can read. Every command prints output a model can parse. There’s no manifest spec or protocol here, just ruthless convention: if Claude Code can read it and run it, it works. Codex, Cursor, and anything else that reads stdout work too.
Which is why setup is one prompt. Open Claude Code and type:
Read https://papercrane.ai/get-started.md and login
Your agent reads the guide, installs the CLI, opens your browser to sign you in, and connects your first source. From there you just ask for what you want to see
Building happens right where you’re chatting, too. papercrane-cli is wired into Claude Code’s preview, so the dashboard renders inside the Claude app while the agent works on it. You watch the charts take shape next to the conversation that’s building them.
The CLI is free. Dashboards run on your own Claude or OpenAI subscription, so there are no AI credits to buy and no per seat pricing to negotiate. The free tier includes five dashboards and every connector. Paid tiers add team features, custom domains, and embedded analytics when you need them.
BI tools have spent thirty years making humans better at operating software. That era is ending. The next BI tool doesn’t need a friendlier query builder, because the thing driving it reads documentation at a thousand words a second and never gets tired of clicking. It needs what an agent needs: credentials, a workspace, and a way to ship. That’s what we’ve built.
If you have Claude Code installed, you’re ninety seconds from a live dashboard on your real data. Paste the prompt, watch what happens, and send me the first dashboard you build. I read everything.
npm install -g @papercraneai/cli , or just tell your agent: Read https://papercrane.ai/get-started.md and login
2026 Papercrane. All rights reserved.
