---
source: "https://www.hauser.io/i-gave-an-ai-agent-its-own-computer-and-it-changed-how-i-see-the-future-of-saas/"
hn_url: "https://news.ycombinator.com/item?id=48462904"
title: "My New Coworker Is an AI Agent with Its Own Computer"
article_title: "My New Coworker Is an AI Agent With Its Own Computer"
author: "bkfh"
captured_at: "2026-06-09T18:53:52Z"
capture_tool: "hn-digest"
hn_id: 48462904
score: 3
comments: 0
posted_at: "2026-06-09T16:05:48Z"
tags:
  - hacker-news
  - translated
---

# My New Coworker Is an AI Agent with Its Own Computer

- HN: [48462904](https://news.ycombinator.com/item?id=48462904)
- Source: [www.hauser.io](https://www.hauser.io/i-gave-an-ai-agent-its-own-computer-and-it-changed-how-i-see-the-future-of-saas/)
- Score: 3
- Comments: 0
- Posted: 2026-06-09T16:05:48Z

## Translation

タイトル: 私の新しい同僚は、独自のコンピューターを持つ AI エージェントです
記事のタイトル: 私の新しい同僚は、独自のコンピューターを持つ AI エージェントです
説明: ベルンハルト・ハウザーによる読み取り (AI ではありません)0:00/293.0677921×
注: この投稿のタイトルが更新されました。
数週間前、私は Reddit で、既存の Claude サブスクリプションにリンクするだけの完全自律型 AI エージェントをセットアップする方法を説明した投稿を見つけました。いいえ

記事本文:
成長するベンチャー企業
ベルンハルト・ハウザーについて
サインイン
購読する
AI
私の新しい同僚は、独自のコンピューターを持つ AI エージェントです
Bernhard Hauser による読み取り (AI ではありません) 0:00 / 293.067792 1× 注: この投稿のタイトルが更新されました。
数週間前、私は Reddit で、既存の Claude サブスクリプションにリンクするだけの完全自律型 AI エージェントをセットアップする方法を説明した投稿を見つけました。有料のオーケストレーターも、学ぶべき新しいプラットフォームもありません。独自の仮想マシン上で 24 時間年中無休で実行され、携帯電話からリモート制御される Claude Code インスタンスだけです。
その日の夕方に 30 分以内にセットアップして、それ以来ずっと改良を続けています。私たちのポートフォリオ企業である Notehouse では、定期的な BI チェックインを実行しているため、ダッシュボードを開かなくても数値を読み取ることができます。また、新しい SEO の機会を調査し、アウトバウンド キャンペーンを管理して私に報告します。このような繰り返しの作業は簡単に確認できます。
これを 30 分以内に起動して実行し、SaaS の将来がどうなるかを垣間見る方法を次に示します。
全体のアイデアは 1 つの決定に基づいています。それは、AI エージェントに完全なアクセス権を持つ独自のマシンを与え、次に必要な結果を必要なだけ詳細に記述することです。そのマシンに完全にアクセスできるため、統合を接続するのを待ちません。タスクを完了するためにツールが必要な場合は、ツール自体を構築します。
小規模なクラウド サーバーを起動します。私は月額 10 ユーロの Hetzner 仮想マシンを使用しましたが、どんな小さな VM でも動作します。エージェントに完全な制御を与えることになるため、重要なものとは別にしてください。
ログインしてクロード コードをインストールします。次に、Claude Code をリモート仮想マシンにインストールし、既存のサブスクリプションに接続します。完了すると、Claude デスクトップまたはモバイル アプリからリモート コントロールできるようになります。
作成する

共有知識ベース。これは、AI エージェントと知識を共有するための重要なステップです。 Markdown ファイルはこれに最適です。Obsidian Git (無料) または Obsidian Sync ($4/月) 経由で同期する共有 Obsidian ボールトが最適です。ボールトには、あなたのビジネスに関するコンテキストと、物事の進め方に関するコンテキストが保存されている必要があります。エージェントは定期的に更新を双方向で同期するように指示されます。
準備ができて！最初に取り組むべきタスクを与えます。セットアップの準備ができたので、Waterglass でビジネスを管理するために使用しているツールへのアクセスを許可しました。ナレッジ ベースを通じて完全なコンテキストを取得し、適切なツール セットを使用できるようになったので、最初のタスクに取り組む準備が整いました。
小規模な仮想サーバー、既存の Claude Code サブスクリプション、共有ナレッジ ベース、そして 30 分の時間。始めるために必要なのはこれだけです。
SaaS の後に何が起こるのかを垣間見る
現在、AI エージェントと、AI エージェントが広範な SaaS 市場に何をもたらすかについて、多くの記事が書かれています。そのほとんどは推測ですが、現時点で絶対に確かなことが 1 つあります。それは、SaaS が変化しており、ユーザーのタスクを独立して実行できる自律エージェントがソフトウェア業界の将来を形作ることになるということです。
わずか 30 分のセットアップで、業界がどこに向かっているのかが分かりました。創設者および運営者として、私が次に行うことは次のとおりです。
ヘッドレス向けにビルドします。エージェントは、ソフトウェア ビジネスと異なる方法でやり取りします。適切に設計されたユーザー インターフェイスは必要ありませんが、対話を容易にするコマンド ライン ツールまたは MCP サーバーが必要です。
ソフトウェア ビジネスの堀は、データ、コンプライアンス、および緊密な統合になります。複製が簡単なソフトウェアは、明確な指示、完全なアクセス権を持つ仮想マシン、およびそれを最初から構築するための大量のリソースを備えた AI エージェントから防御することがますます困難になります。
配信に注力します。ソフトウェアのとき

re が商品になると、実際のレバーはそれを中心に構築するディストリビューションになります。これには、ブランド、コンテンツ、SEO など、さまざまな形が考えられます。
上記のセットアップにはまだ多少の複雑さが伴いますが、ビジネスの運営を支援するパーソナライズされた AI エージェントの作成は、今後ますます簡単になり、おそらく良い方向に変わると私は確信しています。
あなたのビジネスで定期的に行われるどのタスクを最初に AI エージェントに渡しますか?
クロードに説明してもらいましょう
以下の手順をコピーし、クロードとの新しいチャットに貼り付けると、セットアップ全体を 1 ステップずつガイドします。
クロードへの指示をコピーする
他に誰がこのことを知っているでしょうか？
成長するベンチャーに関する毎週の 3 分間のストーリー
400 人以上の創設者と運営者に加わりましょう
今日、流通は製品の前に構築されます
今や誰もがメディア企業です。ほとんどの人はまだそれを知らないだけです。

## Original Extract

Read by Bernhard Hauser (not AI)0:00/293.0677921×
Note: The title of this post has been updated.
A few weeks ago I came across a post on Reddit that explained how to set up a fully autonomous AI agent that simply links to your existing Claude subscription. No

Growing Ventures
About Bernhard Hauser
Sign in
Subscribe
AI
My New Coworker Is an AI Agent With Its Own Computer
Read by Bernhard Hauser (not AI) 0:00 / 293.067792 1× Note: The title of this post has been updated.
A few weeks ago I came across a post on Reddit that explained how to set up a fully autonomous AI agent that simply links to your existing Claude subscription. No paid orchestrator, no new platform to learn. Just a Claude Code instance running 24/7 on its own virtual machine, remote-controlled from your phone.
I set it up the same evening in under 30 minutes and I have been refining it ever since. At our portfolio company Notehouse , it now runs regular BI check-ins, so I get a read on the numbers without opening our dashboards. It also researches new SEO opportunities and manages our outbound campaigns to report back to me, the kind of recurring work that is easy to glance over.
Here is how you can get this up and running in under 30 minutes and get a glimpse into what might be the future of SaaS.
The whole idea rests on one decision: give the AI agent its own machine with full access, then describe the outcome you want in as much detail as needed. Because it has full access to that machine, it does not wait for you to wire up an integration. If it needs a tool to finish a task, it will build the tool itself.
Spin up a small cloud server. I went with a €10 / mo Hetzner virtual machine, but any small VM works. Keep it separate from anything important, because you are about to give your agent full control of it.
Log in and install Claude Code. Now it's time to install Claude Code on your remote virtual machine and connect it with your existing subscription. Once done, you can then remote control it from your Claude desktop or mobile app.
Create a shared knowledge base. This is the important step for sharing knowledge with your AI agent. Markdown files are perfect for this: A shared Obsidian vault that syncs over Obsidian Git (free) or Obsidian Sync ($4 / mo) is a great fit. The vault should hold context about your business and how you like things done. The agent is instructed to regularly sync updates both ways.
Ready! Give it its first task to work on. Now that the setup is ready, I gave it access to the tools we are using at Waterglass to manage our business. Now that it has full context through the knowledge base and the right set of tools to work with, it's ready to work on its first task.
A small virtual server, your existing Claude Code subscription, a shared knowledge base and 30 minutes of your time. That is all you need to get started.
A glimpse of what comes after SaaS
There is a lot of writing right now about AI agents and what they will do to the broader SaaS market. Most of it is speculation, but one thing is absolutely certain by now: SaaS is changing and autonomous agents that can work independently on tasks for their users will shape the future of the software industry.
Only 30 minutes of setup showed me where the industry might be heading. As a founder and operator, this is what I would do next:
Build for headless. Agents will interact with your software business differently. They don't need a well-designed user interface, but rather a command-line tool or an MCP server to make interaction easier.
Your software business's moat is going to be data, compliance and deep integration. Easy-to-replicate software will become harder to defend against AI agents with clear instructions, a virtual machine with full access and lots of resources to build it from scratch.
Focus on distribution. When software becomes a commodity, your real lever is the distribution you build around it. This can take many forms: brand, content, SEO... you name it.
While the setup above still carries some complexity, I'm convinced that creating personalized AI agents to help us run our businesses will only get easier from here – and maybe change it for the better.
What recurring task in your business will you hand to your AI agent first?
Let Claude walk you through it
Copy the instructions below, paste them into a new chat with Claude and it will guide you through the whole setup, one step at a time.
Copy instructions for Claude
Who else should know about this?
Weekly 3-Minute Stories on Growing Ventures
Join 400+ founders & operators
Today, Distribution Is Built Before The Product
Everyone Is a Media Company Now. Most Just Don't Know It Yet.
