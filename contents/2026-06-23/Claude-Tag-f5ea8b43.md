---
source: "https://www.anthropic.com/news/introducing-claude-tag"
hn_url: "https://news.ycombinator.com/item?id=48648039"
title: "Claude Tag"
article_title: "Introducing Claude Tag \\ Anthropic"
author: "adocomplete"
captured_at: "2026-06-23T17:35:22Z"
capture_tool: "hn-digest"
hn_id: 48648039
score: 26
comments: 4
posted_at: "2026-06-23T17:09:18Z"
tags:
  - hacker-news
  - translated
---

# Claude Tag

- HN: [48648039](https://news.ycombinator.com/item?id=48648039)
- Source: [www.anthropic.com](https://www.anthropic.com/news/introducing-claude-tag)
- Score: 26
- Comments: 4
- Posted: 2026-06-23T17:09:18Z

## Translation

タイトル: クロード・タグ
記事のタイトル: Claude Tag の紹介 \ Anthropic
説明: Anthropic は、信頼性が高く、解釈可能で、操作可能な AI システムの構築に取り組んでいる AI の安全性と研究を行う会社です。

記事本文:
クロード タグの紹介 \ Anthropic メイン コンテンツにスキップ フッターにスキップ 研究
クロードタグの製品紹介
クロード タグは、チームがクロードと連携するための新しい方法です。
クロードがチームメンバーとして参加できる Slack を始めます。クロードに選択したチャネルへのアクセスを許可し、選択したツール、データ、さらにはコードベースに接続します。その後、チャンネル内の誰もが @Claude をタグ付けして、他の作業に集中している間にタスクを委任できます。クロードは、参加しているチャネルから関連情報を記憶することでコンテキストを構築し、将来完了するタスクを計画できます。
私たちは、Claude Tag を Claude Code の進化の始まりであると考えています。これにより、モデルがさらに積極的になり、チーム全体でより適切に機能するようになります。 @Claude のタグ付けは、Anthropic で物事を進める主な方法の 1 つになりました。現在、製品チームのコードの 65% は内部バージョンの Claude Tag によって作成されています。同じパターンが現在、エンジニアリングを超えて広がっています。私たちはクロードをタグ付けして、製品のメトリクスとデータを追跡し、サポート チケットを処理し、さらには厄介なバグの根本原因を見つけるのにも役立ちます。
Slack はチームと AI 間の共同作業の自然な拠点であり、Anthropic の日常業務の多くがすでに行われている場所であるため、Slack 上で Claude Tag を開始します。 Claude Enterprise および Team の顧客は現在ベータ版として利用できます。私たちの目標は、利用可能な場所をより広範囲に拡大し、チームが作業する他の多くの場所で @Claude をタグ付けできるようにすることです。
以前に Claude Code や Cowork で働いたことがある人なら、Claude Tag に親しみを感じるでしょう。 @Claude に簡単な言葉でリクエストをタグ付けすると、タスクを段階に分割し、アクセスできるツールを使用して順番に処理します。完了すると、Slack スレッドでその内容が応答されます。

が作成されました。
ただし、Claude をタグ付けすると、いくつかの新しい利点が得られます。
@Claude はマルチプレイヤーです。特定の Slack チャネル内に、全員と対話する 1 人のクロードがいます。これは、誰でも作業内容を確認でき、最後の人が中断したところから会話を再開できることを意味します。これにより、クロードへのタグ付けは、単一のチャット内での作業や単一のタスクでの作業とは大きく異なり、チームメイトと協力して対話することに非常に似ています。
@Claude は時間の経過とともに学習します。クロードがそのチャンネルに沿ってフォローすると、作品に関するコンテキストがさらに構築されます。これは、ユーザーが何度も最初から説明する必要がないことを意味します。また、許可が与えられていれば、クロードは他の Slack チャネルやデータ ソースから自動的に学習することもできます。 (プライベート チャネルからはレポートしません。) これにより、可能な限り最高の作業を提供するために必要な暗黙の知識が得られます。
@Claude が主導権を握ります。 「アンビエント」動作が有効になっている場合、クロードは、ユーザーが知る必要があると思われるものについて積極的に最新情報を提供します。接続されているチャネルや接続されているツール全体から関連情報にフラグを立て、解決されずに沈黙しているスレッドやタスクをフォローアップします。
@Claude は非同期で動作します。クロードにタスクを設定すると、それが機能している間、あなたは他の優先事項に集中できます。また、タスクを自分でスケジュールし、数時間または数日間にわたって自律的にプロジェクトを推進することもできます。これは Anthropic で特に役に立ちます。現在では、多くの時間を並行して多くのクロードにタスクを委任することに費やしています。
クロードにダイレクト メッセージを送信することもできます。設定した個人用ツールとコネクタを使用して、クロードは非公開で応答します。
私たちはチームや組織を念頭に置いて Claude タグを設計しました: @Claude の機密情報へのアクセス

データとタスク固有のツールは非常に厳密に制御できます。
起動して実行するには、システム管理者はモデルがどのチャネルでどのツールと情報にアクセスできるかを指定します。これは、用途に応じて個別のクロード ID を作成するものと考えてください。記憶を含むすべての内容は、管理者が定義したチャネルに限定されます。たとえば、営業業務用にセットアップされたモデルは、エンジニアリング用にセットアップされたモデルに記憶を引き継ぎません。また、エンジニアが販売データやツールにアクセスできるようにすることもありません。プロビジョニング アクセスの詳細については、ここを参照してください。
権限を設定すると、誰でもすぐにタグ付けを開始できます。管理者は、トークンの使用量 (組織と個々のチャネルの両方) に制限を設定でき、@Claude が実行したすべてのログと各タスクの要求者を表示できます。
Claude Enterprise または Team の顧客は、今日からベータ版の Claude Tag にアクセスできます。開始するには、ここにアクセスし、次の 4 つの手順に従ってください。
Claude Tag を Slack ワークスペースとペアリングする
クロードにツールへのアクセスを許可します
組織の月々の支出に制限を設定する
プライベート チャネルでクロードをテストして、機能することを確認します。
Claude Tag は、Slack アプリの既存の Claude を置き換えます。移行するには、管理者は 30 日以内にオプトインできます。企業全体が試用できるよう、対象となるエンタープライズおよびチーム組織に初回起動クレジットを発行します。
クロード タグは Opus 4.8 で動作します。当社のドキュメントと製品ページをご覧ください。
Anthropic がソウル事務所を開設し、韓国の AI エコシステム全体にわたる新たなパートナーシップを発表
Fable 5 および Mythos 5 へのアクセスを一時停止するという米国政府の指示に関する声明
米国政府は、Fable 5 と Mythos 5 へのすべてのアクセスを停止する輸出管理指令を発行しました。
fの結果

最初の人類公記録
消費者の健康データのプライバシー ポリシー

## Original Extract

Anthropic is an AI safety and research company that's working to build reliable, interpretable, and steerable AI systems.

Introducing Claude Tag \ Anthropic Skip to main content Skip to footer Research
Product Introducing Claude Tag
Claude Tag is a new way for teams to work with Claude.
We’re starting on Slack, which Claude can join as a team member. Grant Claude access to selected channels, and connect it to whichever tools, data—and even codebases—you choose. Then, anyone in the channel can tag @Claude in, and delegate tasks to it while they focus on other work. Claude builds context by remembering relevant information from the channels it’s in, and can plan out tasks to complete in the future.
We see Claude Tag as the beginning of an evolution of Claude Code: it makes the model even more proactive, and it works better with a full team. Tagging @Claude is now one of the main ways we get things done at Anthropic. Today, 65% of our product team’s code is created by our internal version of Claude Tag. The same pattern is now spreading well beyond engineering—we’re tagging Claude to chase down product metrics and data, work through support tickets, or even help find the root cause of tricky bugs.
We’re launching Claude Tag on Slack, since it’s a natural home for collaborative work between teams and AI, and where much of Anthropic’s day-to-day work already happens. It’s available today in beta for Claude Enterprise and Team customers. Our goal is to expand where it's available more widely, so that teams can tag @Claude in the many other places they work.
If you’ve worked with Claude Code or Cowork before, Claude Tag will feel familiar. Tag @Claude with a request in simple terms and it’ll break its task down into stages and then work through them in turn, using the tools it has access to. Once it’s done, it’ll respond in a Slack thread with what it’s created.
But tagging Claude comes with a few new advantages:
@Claude is multiplayer. Within a given Slack channel, there’s one Claude that interacts with everyone. This means that anyone can see what it’s working on, and can pick up the conversation from where the last person left off. This makes tagging Claude very different from working within a single chat or for a single task—it’s much more like interacting collaboratively with a teammate.
@Claude learns over time. As Claude follows along with its channel, it builds more context about the work. This means that users don’t need to explain things to it from scratch over and over again. And Claude can even automatically learn from other Slack channels and data sources, if it’s granted permission. (It doesn’t report from private channels.) This gives it the tacit knowledge necessary for it to provide the best possible work.
@Claude takes initiative. If “ambient” behavior is enabled, Claude will proactively keep you updated about whatever it thinks you might need to know. It’ll flag relevant information from across the channels it’s in and the tools it’s connected to, and follow up on threads or tasks that have gone quiet without being resolved.
@Claude works asynchronously. Set Claude a task, and you can focus on your other priorities while it works. It can also schedule tasks for itself, pursuing a project autonomously over hours or days. We’ve found this particularly helpful at Anthropic: we now spend much more of our time delegating tasks to many Claudes in parallel.
You can also send Claude direct messages: it’ll respond privately, using the personal tools and connectors you’ve set up.
We’ve designed Claude Tag with teams and organizations in mind: @Claude’s access to sensitive data and task-specific tools can be very tightly controlled.
To get up and running, system administrators specify which tools and information the model should have access to, in which channels. Think of it as creating separate Claude identities for different uses: everything, including its memories, will stay scoped to the channels defined by the administrators. For example, a model set up for sales work won’t pass on memories to one set up for engineering; nor will it give engineers access to any sales data or tools. More information about provisioning access is available here .
Once permissions are set, everyone can begin tagging right away. Administrators can set limits for token spend (both for the organization and for individual channels), and can view a log of everything that @Claude has done, along with who requested each task.
If you’re a Claude Enterprise or Team customer, you have access to Claude Tag in beta starting today. To get started, visit here and follow these four steps:
Pair Claude Tag with your Slack workspace
Give Claude access to your tools
Set a limit on your organization’s monthly spend
Test Claude in a private channel to confirm it works
Claude Tag replaces the existing Claude in Slack app. To migrate, administrators can opt in within 30 days. We’re issuing an introductory launch credit to eligible Enterprise and Team organizations so that the whole company can try it out.
Claude Tag works with Opus 4.8. You can read our docs and product page .
Anthropic opens Seoul office and announces new partnerships across the Korean AI ecosystem
Statement on the US government directive to suspend access to Fable 5 and Mythos 5
The US government has issued an export control directive to suspend all access to Fable 5 and Mythos 5.
Results from the first Anthropic Public Record
Consumer health data privacy policy
