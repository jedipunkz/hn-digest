---
source: "https://blog.cloudflare.com/temporary-accounts/"
hn_url: "https://news.ycombinator.com/item?id=48598906"
title: "Temporary Cloudflare Accounts for AI Agents"
article_title: "Temporary Cloudflare Accounts for AI agents"
author: "soheilpro"
captured_at: "2026-06-19T16:04:22Z"
capture_tool: "hn-digest"
hn_id: 48598906
score: 7
comments: 1
posted_at: "2026-06-19T14:17:59Z"
tags:
  - hacker-news
  - translated
---

# Temporary Cloudflare Accounts for AI Agents

- HN: [48598906](https://news.ycombinator.com/item?id=48598906)
- Source: [blog.cloudflare.com](https://blog.cloudflare.com/temporary-accounts/)
- Score: 7
- Comments: 1
- Posted: 2026-06-19T14:17:59Z

## Translation

タイトル: AIエージェント用の一時的なCloudflareアカウント
記事のタイトル: AI エージェント用の一時的な Cloudflare アカウント
説明: エージェントが何かを展開する必要がある瞬間、エージェントは人間のために作られた壁に顔から衝突します。今日、私たちは

記事本文:
無料で始めましょう |営業担当者へのお問い合わせ Cloudflare ブログ
新しい投稿の通知を受け取るには購読してください:
AIエージェント用の一時的なCloudflareアカウント
現在、誰もが AI エージェントを使用してコードを作成しています。しかし、エージェントが何かをデプロイする必要があり、サインアップしてアカウントを作成する必要がある瞬間、ブラウザベースの OAuth フロー、クリックスルーするためのダッシュボード、コピーアンドペーストするための API トークン、満足させるための多要素認証プロンプトなど、人間のために構築された壁に正面からぶつかります。開発者の隣に座る対話型の副操縦士にとって、それは煩わしいことです。バックグラウンドエージェントにとって、これは難しいことです。
本日、エージェント向けに一時的な Cloudflare アカウントを展開します。
エージェントは、最初にアカウントにサインアップする必要がなく、Web サイト、API、およびエージェントをすぐに導入できるようになりました。
すべてのエージェントが Wranglerdeploy --temporary を実行して、ワーカーを Cloudflare にデプロイできるようになりました。この一時的な展開は 60 分間有効であり、その間、一時的なアカウントを要求して、永久に自分のものにすることができます。そうしないと、自動的に期限切れになります。
私たちの目標は？エージェントにコーディングして出荷してもらいます。
AI エージェントにとってスムーズな導入が重要な理由
摩擦のない一時アカウントは、見た目よりも重要です。
バックグラウンド AI セッションには人間が関与せず、標準になりつつあります。ブラウザ、コピー＆ペースト、または「60 秒以内にここをクリック」を必要とする認証ステップは、エージェントが行き詰まり、他の場所に展開することを選択する可能性があることを意味します。
試行錯誤はエージェントのスーパーパワーです。エージェントには、書き込み→デプロイ→検証の緊密なループが必要です。彼らは、独自の出力をカールして、それが正しかったかどうかを判断できるように、安価で使い捨てのデプロイメントターゲットを必要としています。
エージェント プラットフォームは、追加の手順や認証情報を必要とせずにコードをデプロイして「正常に機能する」独自の方法を構築しています。人々はこの p に期待し始めています。

rocess は、これまでに使用したことも聞いたこともない他のサービスにサインアップする必要もなく、そのまま機能します。
一時アカウントは、開発者が新しいプロジェクトをブートストラップし、その構成とリソースを管理し、デプロイおよび更新できる開発者プラットフォーム コマンドライン インターフェイス (CLI) ツールである Wrangler を中心に構築されています。
Wrangler の使用法はオンラインで広く文書化されており、エージェントは Wrangler の使用方法をよく知っています。ただし、まだサインインして Cloudflare アカウントに Wrangler 権限を付与していない場合、エージェントがデプロイしようとすると、サインアップと認証のステップで停止します。 「エージェントと LLM は、Wrangler にこの新しい --temporary フラグが存在することをどのようにして認識し、人間が明示的に指示することなく実際にそれを使用できるのでしょうか?」と疑問に思うかもしれません。
これを解決するために、--temporary フラグについてエージェントに通知するメッセージを表示するように Wrangler を更新しました。
エージェントがこれを検出し、--temporary フラグを指定して Wrangler のデプロイを再度実行すると、Cloudflare はエージェントが使用する一時アカウントをプロビジョニングし、Wrangler に使用する API トークンを与え、エージェントが人間に返すことができるクレーム URL を提供します。
フローの各ステップを見てみましょう
新しいプロジェクトのデプロイと反復
最新の Wrangler リリースを使用していることを確認し、お気に入りのコーディング エージェントを起動して、ビルド モードで「hello world」アプリをデプロイするためのプロンプトを作成します。
TypeScript で非常に単純な Hello World Cloudflare Worker を作成し、Wrangler を使用してデプロイします。質問しないで、できる限りのことをしてください。
エージェントは Wrangler を実行し、出力メッセージから --temporary フラグを取得し、スクリプトを構築して即座にデプロイします。ループに人間が関与する必要はありません。
ご覧のとおり、エージェントはスクリプトを作成し、--temporary フラグを使用してデプロイし、プレビュー リンクをカールしました。

出力から取得し、結果がコードと一致することを確認しました。
これは素晴らしいことですが、エージェント コーディングは 1 つのデプロイメントに関するものではないことがよくあります。セッションでは、複数のコード変更のサイクルが実行されることがあります。これは問題ではありません。エージェントはワーカー スクリプトを反復処理し、必要なだけ (60 分の要求ウィンドウ内で) 変更を再デプロイできます。次のプロンプトを入力します。
hello world を「hello Cloudflare」に変更して再デプロイします
エージェントがソース コードを変更し、以前に作成した一時アカウントを再利用し、新しいバージョンを再デプロイして、結果を再確認していることを確認します。
いつでも一時アカウントを取得して、永久に自分のものにすることができます。申請リンクをクリックすると、Cloudflare にサインアップまたはサインインして、Worker がデプロイされた一時アカウントを申請できるページが表示されます。これには、ワーカーだけでなく、データベースやその他のバインディングなどのリソースの要求も含まれます。
これらの一時アカウントを 60 分以内に要求しない場合、それらは自動的に削除されます。
スムーズなエージェント導入への道
これは、エージェントのサインアップの障壁を取り除く 1 つの方法にすぎません。私たちは最近、Stripe とのパートナーシップと、エージェントがユーザーに代わって Cloudflare をプロビジョニングできるようにする共同設計した新しいプロトコルを発表しました。アカ​​ウントの作成、サブスクリプションの開始、ドメインの登録、コードをデプロイするための API トークンの取得を、トークンのコピー＆ペーストやクレジット カードの詳細の入力を行わずに実行できます。先月、私たちは WorkOS と協力して、誰でも採用できる auth.md の立ち上げに取り組み、エージェントが確立された既存の OAuth 標準を使用して新しいアカウントをプロビジョニングできるようにしました。
この分野では多くのことが起こっており、私たちはエージェントが Cloudflare を使いやすく、開発者が Cloudflare を使いやすくできるようにしていきたいと思っています。

独自のアプリをエージェント対応にする。一時アカウントは、スムーズなエージェント展開に向けたもう 1 つのステップです。詳細については、引き続きご期待ください。
一時アカウントにはいくつかの制限があり、その機能は時間の経過とともに変更される可能性があります。詳細については開発者ドキュメントを確認してから、何かをビルドしてください。エージェントに Cloudflare を紹介し、それがどこまで進んでいるのかを確認し、改善できる点や喜んでいただける点を教えてください。X で構築したものを共有したり、Cloudflare コミュニティに参加したりしてください。
独自の脆弱性ハーネスを構築する
多段階の脆弱性発見ハーネスと自動トリアージ ループの背後にある技術アーキテクチャを詳しく分析します。状態制御を管理し、敵対的なレビューを通じて誤検知を排除し、LLM コンテキスト制限を回避する方法を学びましょう。 ...
Flueを皮切りに、より多くのエージェントハーネスとフレームワークをCloudflareに導入
Agent SDK は、あらゆるエージェント フレームワークを構築できるランタイムになりました。本日、エージェント SDK をターゲットとした最初のフレームワークとして Flue を使用してエージェント SDK プリミティブを公開し、ダッシュボードにエージェントを展開します。 ...
Cloudflare Oneスタックの紹介: エージェントによる導入
Cloudflare Oneスタックは、AIエージェントにゼロトラスト環境の計画、導入、管理に必要な知識を提供するエージェントスキルのライブラリであり、移行コールは必要ありません。 ...
Ensemble AI の人材を活用して Cloudflare AI チームを成長させる
Cloudflareは、Ensemble AIのチームメンバーを加えて、機械学習インフラストラクチャと効率に焦点を当ててAIへの投資を強化しています。 ...

## Original Extract

The moment an agent needs to deploy something, it slams face-first into a wall built for humans. Today we

Get Started Free | Contact Sales The Cloudflare Blog
Subscribe to receive notifications of new posts:
Temporary Cloudflare Accounts for AI agents
Everyone's writing code with AI agents today. But the moment an agent needs to deploy something — and needs to sign up and create an account — it slams face-first into a wall built for humans: a browser-based OAuth flow, a dashboard to click through, an API token to copy-paste, a multi-factor authentication prompt to satisfy. For an interactive copilot sitting next to a developer, that's annoying. For a background agent, it's a hard stop.
Today we're rolling out Temporary Cloudflare Accounts for Agents.
Agents can now deploy websites , APIs , and agents right away, without first needing to sign up for an account.
Any agent can now run wrangler deploy --temporary and deploy a Worker to Cloudflare. This temporary deployment stays live for 60 minutes, during which time you can claim the temporary account, making it permanently your own. If you don't, it expires on its own.
Our goal? Let your agent code and ship.
Why frictionless deployments matter for AI agents
Frictionless temporary accounts matter more than it might first seem:
Background AI sessions have no human in the loop, and are becoming the norm . Any auth step that needs a browser, a copy-paste, or "click here in 60 seconds" means an agent gets stuck and may choose to deploy elsewhere.
Trial-and-error is the agent's superpower . Agents need a tight write → deploy → verify loop. They need cheap, throwaway deployment targets, so they can curl their own output and decide whether they got it right.
Agent platforms are building their own ways for deploying code to "just work" without extra steps or credentials . People are starting to expect that this process just works, without the need to sign up for other services that they've not used before or heard of.
Temporary accounts are built around Wrangler , our Developer Platform command-line interface (CLI) tool that lets developers bootstrap new projects, manage their configurations and resources, and deploy and update them.
Wrangler usage is widely documented online and agents know how to use it very well. But if you hadn’t yet signed in and granted Wrangler permission to your Cloudflare account, when the agent tried to deploy, it would get stuck at the sign-up and authentication step. And you might rightly ask: How do agents and LLMs know that this new --temporary flag in Wrangler exists, so that they actually use it without a human explicitly telling them to do so?
To solve this, we updated Wrangler to prompt the agent with a message that tells it about the --temporary flag:
When the agent discovers this, and then runs wrangler deploy again with the --temporary flag, Cloudflare provisions a temporary account for the agent to use, gives Wrangler an API token to work with, and provides a claim URL that the agent can give back to the human.
Let’s go over every step of the flow
Deploying and iterating on a new project
Make sure you’re using the latest Wrangler release , fire up your favorite coding agent, and write a prompt to deploy a "hello world" app in build mode:
Make a very simple hello world Cloudflare Worker in TypeScript and deploy it using wrangler, don't ask me questions, do the best you can
The agent will run wrangler, pick up the --temporary flag from the output messages, build your script, and deploy it instantly, no human in the loop required:
As you can see, the agent wrote the script, deployed it using the --temporary flag, curled the preview link it got from the output, and verified that the result matches the code.
This is great, but agentic coding is often not about one single deployment. A session can go through a cycle of multiple code changes. This is not a problem: the agent can iterate on the Worker script and redeploy the changes as many times as it wants (within the 60-minute claim window). Type this prompt:
Now change hello world to "hello cloudflare" and redeploy
Look at the agent changing the source code, reusing the previously created temporary account, redeploying a new version and rechecking the result:
At any point, you can claim the temporary account and make it yours permanently. When you click the claim link you will be taken to a page where you can either sign up for or sign in to Cloudflare, and then claim the temporary account that your Worker was deployed to. This includes claiming not just Workers, but resources like databases and other bindings, too.
If you do not claim these temporary accounts within 60 minutes, they will be automatically deleted.
The road to frictionless agentic deployments
This is just one way we’re eliminating the signup barrier for agents. We recently announced a partnership with Stripe and a new protocol we co-designed that lets agents provision Cloudflare on behalf of their users — creating an account, starting a subscription, registering a domain, and getting an API token to deploy code, with no copy-pasting tokens or entering credit card details. Last month, we collaborated with WorkOS on the launch of auth.md , which anyone can adopt, to let agents provision new accounts using well-established, existing OAuth standards.
There’s a ton going on in this space, and we’re excited to keep making it easier for agents to use Cloudflare, and for developers to make their own apps agent-ready . Temporary accounts are one more step toward frictionless agentic deployments — stay tuned for more.
Temporary accounts have some limitations, and their capabilities may change over time; check the developer documentation for more information and then go build something. Point your agent at Cloudflare, see how far it gets, and tell us what we can improve or what delights you — share what you’ve built on X or hop into the Cloudflare Community .
Build your own vulnerability harness
We break down the technical architecture behind our multi-stage vulnerability discovery harness and automated triage loop. Learn how we manage state controls, squash false positives through adversarial review, and route around LLM context limits. ...
Bringing more agent harnesses and frameworks to Cloudflare, starting with Flue
The Agents SDK is now a runtime any agent framework can build on. Today we're opening up the Agents SDK primitives, with Flue as a first framework targeting Agents SDK, and rolling out agents in the dashboard. ...
Introducing the Cloudflare One stack: agent-powered deployment
The Cloudflare One stack is a library of agent skills that gives any AI agent the knowledge it needs to plan, deploy, and manage a Zero Trust environment — no migration calls required. ...
Growing the Cloudflare AI team with talent from Ensemble AI
Cloudflare is deepening our investment in AI with the addition of team members from Ensemble AI, focusing on machine learning infrastructure and efficiency. ...
