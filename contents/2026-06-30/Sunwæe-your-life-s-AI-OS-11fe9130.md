---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48732413"
title: "Sunwæe – your life's AI OS"
article_title: ""
author: "dvdxnss"
captured_at: "2026-06-30T13:33:59Z"
capture_tool: "hn-digest"
hn_id: 48732413
score: 1
comments: 0
posted_at: "2026-06-30T13:21:31Z"
tags:
  - hacker-news
  - translated
---

# Sunwæe – your life's AI OS

- HN: [48732413](https://news.ycombinator.com/item?id=48732413)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T13:21:31Z

## Translation

タイトル: Sunwæe – あなたの人生の AI OS
HN テキスト: アプリの URL: https://app.sunwaee.com ねえ、それで、私は chatgpt や claude Web アプリにあまりこだわったことはありません。なぜなら、私にとって、それらはいつも少しパーソナライズされた Google 検索のように感じていたからですが、実際に自分が誰であるかを知るにはほど遠いからです。それに、最高のモデルにアクセスしたいのですが、最近の新モデルのリリース頻度により、多くの場合、複数の高価なサブスクリプションが必要になります...ちょっと怠け者でもあるので、特定のトピックに対して必要なモデルを選択する自動ルーティングのようなものがあると考えましたかっこいい。それに加えて、私は Google 化された生活 (つまり Google エコシステム) から抜け出す旅に出ており、すべてを整理することに夢中で、メモ、タスク、ファイル、カレンダー、検索などのアプリを切り替えるのに少しうんざりしています...それらのアプリは AI 統合が不十分 (Gemini...) か AI 統合がまったくないため、なおさらです... そこで、私はここ数か月間 https://app.sunwaee.com に取り組んでいます。ブラウザ上のオールインワン AI OS では、ほとんどの LLM にエージェントとしてアクセスできるだけでなく、アプリ内でネイティブのメモ、タスク、ファイル、カレンダーに直接アクセスでき、すべてワークスペースベースです。他にも次のようなすばらしい機能があります。 - クエリのトピックに基づいて最適なモデルを選択し、カスタマイズできるオートルーター (コーディング用の opus、ビジネス用の gpt、数学用のディープシークなど)、さらに会話中にモデルを切り替えることができます - 3 レベルのメモリ システム: ユーザー > ワークスペース > プロジェクト、Sun (AI エージェント) がアカウント全体の情報を検索できます - Sun がメモ、タスク、ファイルからあらゆるものを作成/編集し、ピン留めできるワークベンチイベント、または Web 検索 (近々、本、航空券、アニメ、映画、製品などのカスタム オブジェクトも...) - 基本的にメッセージから電子メールなどを編集する PII 編集システム、Sun メッセージ

esだけでなく、LLMプロバイダーのAPIに到達する前のツール呼び出しも含まれます。サブエージェント、サンドボックスでのファイル編集（msft、pdf、画像など）、スケジュールされたジョブ、スキル、基本的なコネクタなど、すべてをソロで構築しました（psql/fastapi/nextjs）が、現時点ではシングルユーザーのみであり、共有/コラボレーション機能はまだありません。ドメインに配置して以来、検索から計画まで、基本的にどこからでもすべてに使用していることに気付きました。学習しており、現時点ではそれが私の毎日の原動力となっています。これを非常に誇りに思っています（また、データが販売されたり、広告目的で使用されたりしないことを知ることは気分が良いです）。今後数週間または数か月以内に、コーディング エージェントがワークスペース/チャット/タスクにアクセスできるように、プロアクティブ/フォローアップ機能、チェックイン、およびある種の MCP を追加する予定ですが、現時点では、アプリ自体についての外部からの正直なフィードバックを探しています。一歩下がってみると、これはすべてのスタートアップのように *1 つの特定の問題* を解決するアプリやサービスではないことは確かです。3 つの複雑な単語を並べてそれをセールス ポイントにするだけであることがわかります (例: 「エージェント建設管理」)。そのため、これについて人々に話すと、アイデアや賭けは単純ですが、ポジショニングの部分は説明するのが難しいです。すべてが 1 つのアプリに含まれており、Sun のエージェントはすべてのアプリで積極的に機能します。何人かの人がこれを試してみることに本当に感謝します。 (https://app.sunwaee.com - サインアップは必要ありません。無料トライアルも可能です) そして、彼らの意見を正直に教えてください。アイデアや製品に興味があると思う人と喜んでつながります。デビッド

## Original Extract

app url: https://app.sunwaee.com hey, so I never really stuck with chatgpt or claude web apps because to me they always felt like a slightly-more-personalized google search but far from actually knowing who I am, plus I also want access to the best models but with the release frequency of new models these days it often means multiple expensive subscriptions... kinda lazy too so I figured having some kind of autorouting that picks the model I want for given topics would be cool. On top of that I've been on a journey to get out of the google-ized life (i.e. google ecosystem), I'm obsessed with organizing everything, and I'm kinda growing tired of switching apps for notes, tasks, files, calendars, search, etc... even more so since those apps either have bad AI integrations (gemini...) or no AI integrations at all... So I've been working on https://app.sunwaee.com for the last few months now, which is basically a all-in-one AI OS in your browser where you not only have access to most LLMs as agents, but also to native notes, tasks, files and calendars directly in the app - everything workspace-based. There are also a few other cool things, like: - an autorouter that picks the best model based on the query topic and that you can customize (e.g. opus for coding, gpt for business, deepseek for maths..), plus you can switch model mid-conversation - three-levels memory system: user > workspace > project, and Sun (the AI agent) can search your entire account for stuff - a workbench where Sun can create/edit and pin anything from notes, tasks, files, events, or web searches (and soon custom objects like books, flights, animes, movies, products...) - a PII redaction system that basically redacts emails etc from your messages, Sun messages, but also tool calls before they reach the LLM providers APIs - and more, like subagents, file editing (incl. msft, pdfs, images..) in a sandbox, scheduled jobs, skills and basic connectors Built everything solo (psql/fastapi/nextjs) but it's single-user only right now, no sharing/collab features yet Ever since I put it on my domain I find myself using it basically for everything and from anywhere, from search to planning to learning and at this point it has become my daily-driver - which I'm very proud of (also knowing your data isn't sold or used for advertising purposes feels good). I'm also planning to add proactive/follow-up features, check-ins, and some kind of MCP such that coding agents can access your workspaces/chats/tasks in the next few weeks/months But at this point I'm looking for external, honest feedback on the app itself. When I take a step back it sure isn't an app/service that solves *one-specific problem* like all startups you can see that just align 3 complex words and make it their selling point (e.g. "Agentic Construction Management") so when I talk to people about it the positionning part is hard to describe, even though the idea/bet is simple - everything in one app and Sun the agent that acts proactively across all of it Would really appreciate a few people trying this out (https://app.sunwaee.com - no signup required, free trial too) and telling me honestly what they think, happy to connect with anyone who finds the idea/product interesting! David

