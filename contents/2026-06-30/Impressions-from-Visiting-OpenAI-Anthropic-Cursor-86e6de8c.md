---
source: "https://newsletter.pragmaticengineer.com/p/impressions-from-visiting-openai"
hn_url: "https://news.ycombinator.com/item?id=48736207"
title: "Impressions from Visiting OpenAI, Anthropic, & Cursor"
article_title: "Impressions from visiting OpenAI, Anthropic, & Cursor"
author: "monkeydust"
captured_at: "2026-06-30T18:35:56Z"
capture_tool: "hn-digest"
hn_id: 48736207
score: 1
comments: 0
posted_at: "2026-06-30T17:34:18Z"
tags:
  - hacker-news
  - translated
---

# Impressions from Visiting OpenAI, Anthropic, & Cursor

- HN: [48736207](https://news.ycombinator.com/item?id=48736207)
- Source: [newsletter.pragmaticengineer.com](https://newsletter.pragmaticengineer.com/p/impressions-from-visiting-openai)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T17:34:18Z

## Translation

タイトル: OpenAI、Anthropic、Cursor を訪問した感想
記事タイトル: OpenAI、Anthropic、Cursor を訪問した感想
説明: ソフトウェア エンジニアリングがどこに向かっているのかを、この分野をリードする AI ラボの内部から覗いてみましょう。クラウド上で実行されるエージェントが主要なトレンドである一方、コーディングハーネスは技術を超えて広がりを見せています

記事本文:
OpenAI、Anthropic、Cursor を訪問した感想
実践的なエンジニア
The Pulse OpenAI、Anthropic、Cursor を訪問した感想
ソフトウェア エンジニアリングがどこに向かっているのかを、この分野をリードする AI ラボの内部から覗いてみましょう。クラウド上で実行されるエージェントが主要なトレンドである一方、コーディングハーネスは技術を超えて広がりを見せています
Gergely Orosz Jun 30, 2026 ∙ 有料 47 4 シェア スケジュールに関する注意: 今週、私は AI エンジニア ワールド フェアでサンフランシスコにいるため、木曜日の The Pulse 版はありません。ただし、明日 (水曜日) には、ソフトウェア エンジニアリングの伝説的人物、ケント ベックによる、これまでで最も長く、最も詳細な特別なポッドキャスト エピソードが放送されます。
ここ数日、私はサンフランシスコにある OpenAI、Anthropic、Cursor のオフィスを訪問しました。オンサイトでは、モデル側で作業しているソフトウェア担当者と話をして、ソフトウェアの構築方法がどのように変化しているかを詳しく学びました。この記事は、業界全体で採用される可能性があると思われるいくつかの新しい開発を含む、これらの訪問での観察に基づいています。
次のメガトレンドは？クラウドで実行されるエージェントが主流になります。 OpenAI、Anthropic、Cursor はクラウド エージェントにオールインしており、それらに対する需要が大幅に増加すると予想されます。
開発者以外によるコーディング ハーネスの大量採用。 OpenAI では、非エンジニアの 95% 以上が ChatGPT ではなく Codex を使用しています。テクノロジーに出会うことは何かの前兆なのでしょうか？
エンジニアの主な仕事はエージェントの効率を高めることでしょうか? Anthropic と Cursor でエージェントがより効率的に実行できる環境を構築するエンジニアリング作業がますます増えています。
次のトレンドは？企業はトークンごとの支出を積極的に最適化しています。ソフトウェア エンジニアによる AI 支出は非常に高額であるため、プラットフォーム チームがトークンあたりのコストを削減するのは理にかなっています。 Coinbase のケーススタディ。
1.

次のメガトレンドは？クラウドで実行されるエージェントが主流に
先週、Andrej Karpathy 氏は「新しいパラダイム」という言葉を使って、Claude Tag (Slack でクロードに言及し、タスクを開始させる方法) を使用して AI と連携することを説明しました。
アンドレイ・カルパシー on X この主張に対しては、ソーシャルメディア上で多くの反発がありました。結局のところ、それはクロードと Slack を統合しただけですよね?私も、Anthropic のオフィスを訪問した際に、Anthropic の Applied AI 部門の David Hershey にこのことについて尋ねるまでは、そう考えていました。彼は、この特定の Slack 統合が Claude Code のようなものを使用する場合と何が違うのかを詳しく説明しました。
追加のセットアップはありません。クロード コードが適切に動作するには、ローカル マシン上で適切なスキルを備えた内部 MCP サーバーに接続する必要があります。もちろん、大企業ではこのセットアップは少なくとも部分的に自動化されていますが、開発者は多くの場合調整を行う必要があります。
「ツールのコンテキスト切り替え」はありません。Slack でメンションするだけです。もちろん、Claude Code を開くのは大した労力ではありませんが、Slack にただ入力して仕事を開始するよりも手間がかかります。
日常的な作業が簡単になりました。デビッドはサイドプロジェクトとして「ポケモンをプレイするクロード」を持っています。新しいモデルが発表されるたびに、彼はそれに対してスクリプトの実行を開始します。以前は、これをセットアップするのに毎回数分かかり、その後、彼のマシン上で何時間も実行されていました。この新しい Slack 統合では、コマンド 1 つだけで済みます。
私の感覚では、ここでの興奮は Slack の統合自体に関するものではなく、ローカル マシン上で実行されなくなった 1 つ以上の AI を簡単に起動できるという事実に関係していると考えています。セットアップを完全にスキップすることもできます。
「クロード管理エージェント」は Anthropic の大きな焦点です。そこにいる間、私は Claude Platform のエンジニアリング責任者である Katelyn Lesse に会いました。彼は、Claude が管理しているものであると説明しました。

Agents は、彼女のチームが 6 か月かけて構築した大規模で複雑なプロジェクトです。これは、さまざまなクラウド プロバイダーで長時間実行されるエージェントを実行するためのホスト型サービスです。
Slack の統合ではなく、クラウド エージェントが「重要」です
また先週、私はプライベート AI ビルダー イベントに参加する機会がありました。そこで Peter Steinberger がワークフローについて話し合いました。
Peter Steinberger 氏は、AI コーディング エージェントの使用方法について説明しています。彼は、ローカル マシン上で複数の OpenClaw エージェントを実行していることに本当にうんざりしていて、CPU が熱くなり、システム全体が遅くなったと話しました。そこで、彼はクラウドで OpenClaw エージェントを実行する方法として Crabbox を構築しました。
Crabbox: OpenClaw 用のリモート エージェント ローカルで実行されているエージェントによって引き起こされる問題に対応して、突然、クラウド エージェントの同じソリューションが別の場所 (Anthropic と Peter’s OpenClaw を使用) に登場しました。また、OpenAI と Cursor でもクラウド エージェントが大きな話題になっているということも知りました。
OpenAI はクラウド エージェントに大きく賭けています
OpenAI は、クラウド開発環境 (CDE) のリーダーである Ona (旧 Gitpod) を買収しました。 2021 年に遡ると、CDE は開発者がソフトウェアをより迅速に開発できるように構築されており、サンドボックス化されたクラウド環境でエージェントを実行するための完璧なプリミティブでもありました。 OpenAIによる買収発表より（強調は私のもの）：
「Codex の機能が向上するにつれ、その最も価値のある作業は数分ではなく、数時間または数日かけて行われるようになります。私たちは、人々が作業を開始したマシンに縛られることなく、より野心的な作業を委任できるようにすべきだと考えています。作業は最初のセッションを超えて継続する必要があり、Codex を使用すると、どこからでも接続を維持して進捗状況を確認し、指示を出し、決定を下し、結果を確認することが可能になります。
オナはそれを手伝ってくれます。そのテクノロジーは安全で永続的な環境を提供します

ここでは、エージェントが時間の経過とともに進歩するために必要なツール、システム、コンテキストにアクセスできます。
Ona を OpenAI に導入することで、単一のデバイスやアクティブなセッションに関連付けられた作業を超えて Codex を拡張し、より多くの組織がエージェントを本番環境に安全に導入できるように支援します。」
OpenAI のオフィスで、私はそこのエンジニアに、彼らの焦点がクラウドベースのエージェントに移っているかどうか尋ねました。 Their answer: it very much is.これはごく最近の開発であり、クラウド エージェント チームのエンジニアを募集しています。 Here’s one job ad that’s currently live:
「当社のクラウド エージェント プラットフォームの構築と拡張を支援してくれる、経験豊富なソフトウェア エンジニアを探しています。エージェントを大規模に調整するためのシステムを設計および運用します。 ChatGPT、API、Codex に関して製品エンジニアと緊密に連携して、適切な抽象化を定義し、製品を迅速に出荷できるようにします。バックエンドまたはインフラストラクチャの強力な経験が重要です。 Python、Rust、分散システム、クラウド インフラストラクチャ、または製品プラットフォームの経験が特に役に立ちます。」
カーソル: クラウドでエージェントを実行するのが未来
Cursor では、共同創業者の Sualeh Asif (元 CTO、現在は最高製品責任者) と 1 時間を過ごしました。 Cursor は昨年末に Cloud Agents をリリースし、この分野にさらに注力し始めています。 Sualeh 氏は、クラウド エージェントとの連携に関する興味深い詳細を明らかにしました。
クラウド上のエージェントには「苦情を言う」方法がありません。エージェントをローカルで実行すると、警告やエラーを受け取ると、その応答で人間にそれらが表示され、人間はエージェントに X または Y を実行するように指示します。しかし、クラウド上の長時間実行タスクにはそのようなループはありません。 Cursor は定期的なインタビューでモデル「告白」のアイデアを思いつき、その「告白」はインフラ チームと共有され、エージェントの環境が改善されます。

入口。
長期にわたるエージェントには独自の課題があります。ノードが途中で終了すると何が起こるか。エージェントの実行をあるノードから別のノードに移動するにはどうすればよいでしょうか?チームが解決する必要がある、新たな重要なエンジニアリング課題があります。
つい昨日（6 月 29 日月曜日）、Cursor はどこからでもソフトウェアを構築できる iOS アプリをリリースしました。
スマートフォン上でソフトウェアを構築するにはクラウド エージェントが必要です。出典: Cursor この製品は、長時間実行されるタスクを可能にするためにクラウド エージェント上に構築されていると同社は述べています:
「クラウド エージェントは、作業をテスト、検証、デモするための完全な開発環境を備えた分離された仮想マシンで実行されます。クラウド エージェントは独自のツールやリソースと非同期で動作するため、介入なしで長時間実行し、マージ準備ができた PR に向けて反復することができます。
これらの機能を活用するには、ローカル プランをクラウド エージェントに送信するか、アクティブなエージェントをクラウドに移動して実行を継続します。クラウド セッションをコンピュータに戻して、結合する前にローカルで変更をテストできます。」
なぜクラウド エージェントが突然注目されるようになったのでしょうか?
AI エージェントをクラウドで実行するのが現実的であると考えています。セットアップの手間が少なく、複数のエージェントを並行して実行できます。長時間実行するエージェントにとって、クラウドは個人のラップトップよりも優れた便利な場所です。つまり、オフィス内を歩き回るときでも、蓋を開けたままにしなければなりません。
しかし、なぜ今このようなことが起こっているのでしょうか？私の仮説は、さまざまな要因が関与しているということです。
コーディング モデルは「十分に優れた」ものになりました。 Opus 4.5 / GPT-5.4 より前は、AI モデルは実際には自律的にコーディングできなかったため、長時間のタスクで AI モデルを実行することは無意味でした。
AI コーディング エージェントのインフラは成熟しました。エージェントにより多くのコンテキストを与える方法が改善され、MCP やスキルなどが主流になり、よりよく理解されるようになりました。
コンテキストウィンドウが大きくなりました。今日のモデルには、

最大 100 万トークンのコンテキスト ウィンドウ。つまり、より複雑な命令、コード、コンテキストを渡すことができます。大きなコンテキスト ウィンドウにアクセスせずにエージェントを長時間実行するのは困難です。
クラウド プロバイダーは、はるかに多くの GPU 容量を備えています。ここ数年、すべてのクラウド プロバイダーが GPU クラスターを構築しており、現在ではこれらの AI エージェントがこのインフラを利用できるほどの GPU クラスターが用意されています。
2. 開発者以外によるコーディング ハーネスの大量採用?
OpenAI では、Codex チームの最初のエンジニアだった Andrew Ambrosino にも会いました。私たちの時間は理想的なスタートを切りました。アンドリューは私に素晴らしいものを見せてほしいと言いました。
この投稿は有料購読者向けです

## Original Extract

A peek into where software engineering is headed from inside the sector’s leading AI labs. Agents running in the cloud are a major trend, while coding harnesses are spreading beyond the craft

Impressions from visiting OpenAI, Anthropic, & Cursor
The Pragmatic Engineer
Subscribe Sign in The Pulse Impressions from visiting OpenAI, Anthropic, & Cursor
A peek into where software engineering is headed from inside the sector’s leading AI labs. Agents running in the cloud are a major trend, while coding harnesses are spreading beyond the craft
Gergely Orosz Jun 30, 2026 ∙ Paid 47 4 Share Scheduling note: this week, I’m in San Francisco at the AI Engineer’s World Fair , so there won’t be an edition of The Pulse on Thursday. However, tomorrow (Wednesday) there will be a special podcast episode – the lengthiest, most detailed one yet – with software engineering legend, Kent Beck.
In recent days, I’ve visited the offices of OpenAI, Anthropic, and Cursor, in San Francisco. Onsite, I talked with software folks working on the model side to learn more about how their way of building software is changing. This article is based on observations from those visits, including some new developments that I reckon may be adopted industry-wide.
Next mega-trend? Agents running in the cloud to go mainstream. OpenAI, Anthropic, and Cursor are all-in on cloud agents and expect demand for them to increase massively.
Mass adoption of coding harnesses by non-developers. At OpenAI, more than 95% of non-engineers use Codex, not ChatGPT. Is it a sign of things to come across tech?
Will the main task of engineers be to make agents more efficient? Ever more engineering work is about building environments for agents to execute more efficiently at Anthropic and Cursor.
Next trend? Companies aggressively optimize spend-per-token. AI spending by software engineers is so high that it makes sense for platform teams to slash per-token cost. A case study from Coinbase.
1. Next mega-trend? Agents running in the cloud to go mainstream
Last week, Andrej Karpathy employed the phrase “new paradigm” to describe using Claude Tag – a way to mention Claude in Slack and have it kick off tasks – to work with AI:
Andrej Karpathy on X There was plenty of pushback against this claim on social media; after all, it’s just a Slack integration with Claude, right? I also thought this until I asked David Hershey at Anthropic’s Applied AI unit about it while visiting the company’s offices. He explained in detail what makes this particular Slack integration different from using something like Claude Code:
No additional setup. For Claude Code to work well, it should be connected to internal MCP servers, with the right skills on your local machine. Of course, at larger companies this setup is at least partially automated, but devs often need to do tweaking.
No “tool context-switching. ” Just mention it in Slack! Of course, opening Claude Code is not a big effort, but it’s still more work than just typing it out in Slack, and kicking off work.
Routine work made easier. David has “Claude playing Pokémon” as his side project. Every time a new model comes out, he kicks off a run of his script on it. Previously, this took a few minutes to set up every time, and then it ran on his machine for hours. With this new Slack integration, it’s just one command.
My sense is that the excitement here is less about the Slack integration itself, and more to do with the fact that it’s easy to kick off one or more AIs that no longer run on a local machine. You can skip the setup entirely.
‘Claude Managed Agents’ is a big focus at Anthropic. While there, I met Katelyn Lesse, head of engineering for Claude Platform, who explained that Claude Managed Agents is a large, complex project which her team built over a six-month period. It’s a hosted service to execute long-running agents on various cloud providers.
Cloud agents are the “big deal”, not the Slack integration
Also last week, I had the opportunity to attend a private AI builders event, where Peter Steinberger discussed his workflow.
Peter Steinberger covers how he uses AI coding agents He talked about how he has gotten really tired of having several OpenClaw agents running on his local machine, which heat up the CPU and slow down his whole system. So, he built Crabbox as a way to run OpenClaw agents in the cloud:
Crabbox: remote agents for OpenClaw Suddenly, the same solution of cloud agents has emerged in separate places – at Anthropic and with Peter’s OpenClaw – in response to issues caused by locally-running agents. I also learned that cloud agents are becoming a big deal at OpenAI and Cursor, too.
OpenAI bets big on Cloud Agents
OpenAI acquired Ona, (formerly Gitpod), the leader in cloud development environments (CDEs). Back in 2021, CDEs were built for developers to develop software faster, and they also happen to be the perfect primitive for agents to run in a sandboxed cloud environment. From the acquisition announcement by OpenAI (emphasis mine):
“As Codex becomes more capable, its most valuable work is unfolding over hours or days, rather than minutes. We believe people should be able to delegate more ambitious work without remaining tied to the machine where it began. The work should continue beyond the initial session, with Codex making it possible to stay connected and check progress, provide direction, make decisions, and review results from anywhere.
Ona will help us do that. Its technology provides secure, persistent environments where agents can access the tools, systems, and context they need to make progress over time.
By bringing Ona to OpenAI, we will expand Codex beyond work tied to a single device or active session and help more organizations deploy agents securely in production.“
At OpenAI’s offices, I asked engineers there if their focus is shifting to cloud-based agents. Their answer: it very much is. This is a fairly recent development and they’re hiring engineers for the Cloud Agents team. Here’s one job ad that’s currently live:
“We are looking for an experienced software engineer to help build and scale our cloud agent platform. You will design and operate systems for orchestrating agents at scale. You will work closely with product engineers on ChatGPT, API, and Codex to define the right abstractions and enable them to ship products quickly. Strong backend or infrastructure experience is important; experience with Python, Rust, distributed systems, cloud infrastructure, or product platforms is especially helpful.”
Cursor: running agents in the cloud is the future
At Cursor, I spent an hour with cofounder Sualeh Asif (formerly the CTO, now Chief Product Officer). Cursor released Cloud Agents at the end of last year, and is starting to focus a lot more on this area. Sualeh revealed some interesting details about working with cloud agents:
Agents in the cloud don’t have a way to “complain.” With running an agent locally, when it gets warnings or errors, it surfaces them to a human in its response, who instructs it to do X or Y. However, there’s no such loop for a long-running task on the cloud! Cursor came up with the idea for the model “confess” in regular interviews, and the “confessions” are shared with the infra team to improve the agents’ environment.
Long-running agents have their own challenges. What happens when a node terminates, midway through; how do you move agent execution from one node to the other? There are new, nontrivial engineering challenges the team needs to solve.
Only yesterday, (Monday, 29 June), Cursor launched its iOS app that enables the building of software from anywhere.
Building software on a smartphone needs cloud agents. Source: Cursor This product is built on top of cloud agents to allow for long-running tasks, the company said:
“Cloud agents run in isolated virtual machines with full development environments to test, verify, and demo work. Since they operate asynchronously with their own tools and resources, cloud agents can run for longer and iterate toward merge-ready PRs without intervention.
To take advantage of these capabilities, send a local plan to a cloud agent or move active agents to the cloud to keep running. You can move the cloud session back to your computer to test changes locally before merging”.
Why are cloud agents suddenly a thing?
It figures that running AI agents in the cloud is practical: there’s less setup involved, several can run in parallel, and the cloud is a better, more convenient place for long-running agents than a personal laptop is; i.e., having to keep the lid open even when walking around the office.
But why is this happening now? My hypothesis is that a mix of factors are at play:
Coding models got ‘good enough’. Before Opus 4.5 / GPT-5.4, AI models could not really code autonomously, so running them for long tasks was pointless!
Infra for AI coding agents has matured. Ways of giving more context to agents have improved: things like MCP and skills became mainstream and better understood.
The context window is bigger. Today’s models have context windows of up to 1 million tokens, meaning that more complex instructions, code, and context can be passed in. It’s hard to have agents run for a longer time without access to a large context window.
Cloud providers have much more GPU capacity. Every cloud provider has been building GPU clusters in the last few years, and now there’s enough that these AI agents can make use of this infra.
2. Mass adoption of coding harnesses by non-developers?
At OpenAI, I also met Andrew Ambrosino , who was the first engineer on the Codex team. Our time together got off to an ideal start, with Andrew saying he needed to show me something incredible:
This post is for paid subscribers
