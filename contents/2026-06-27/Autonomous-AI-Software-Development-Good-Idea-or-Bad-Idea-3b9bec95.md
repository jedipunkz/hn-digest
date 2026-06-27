---
source: "https://adrianavillela.com/post/the-great-autonomous-ai-experiment/"
hn_url: "https://news.ycombinator.com/item?id=48698552"
title: "Autonomous AI Software Development: Good Idea, or Bad Idea?"
article_title: "Autonomous AI Software Development: Good Idea, or Bad Idea?"
author: "mooreds"
captured_at: "2026-06-27T14:31:02Z"
capture_tool: "hn-digest"
hn_id: 48698552
score: 1
comments: 0
posted_at: "2026-06-27T14:22:10Z"
tags:
  - hacker-news
  - translated
---

# Autonomous AI Software Development: Good Idea, or Bad Idea?

- HN: [48698552](https://news.ycombinator.com/item?id=48698552)
- Source: [adrianavillela.com](https://adrianavillela.com/post/the-great-autonomous-ai-experiment/)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T14:22:10Z

## Translation

タイトル: 自律型 AI ソフトウェア開発: 良いアイデア、それとも悪いアイデア?
説明: Paperclip と BMAD を使用した自律型 AI ソフトウェア開発の探索

記事本文:
自律型 AI ソフトウェア開発: 良いアイデアか、悪いアイデアか?アドリアナ・ヴィレラ
私の考え、仕事、そしてその間にあるすべてのもののコレクション。
自律型 AI ソフトウェア開発: 良いアイデアか、悪いアイデアか?
私たちは AI の偉大な実験に数年取り組んでいますが、AI を私たちのテクノロジー生活に効果的に組み込む方法については、まだ多くの議論が行われています。その極端な例としては、AI を心から受け入れる人たちがいます。その対極には、AI に何かをすることを拒否する人たちがいます。あらゆる大きなテクノロジーの破壊と同様、答えは中間のどこかにあります。
テクノロジー業界の多くの同僚と同様に、私も仕事で AI を使用する「適切なバランス」を見つけるのに今でも苦労しています。私を助けるために、さまざまなトピックについて実験し、勉強しています。私の最新の一連の実験は、私を自律型 AI ワークフローの素晴らしい世界に導きました。それが今日お話しする内容です。
しかし、本題に入る前に、少し寄り道して、用語について少し整理してみましょう。
チャットボット: 当初*、AI チャットボットがありました。これらは、AI の水門を開いた ChatGPT の出現で世界を席巻しました。クロード 、 副操縦士 、 ジェミニ のような他の人もすぐに続きました。彼らはきちんとしていました！彼らにシェイクスピアの作品について質問したり、楽しい絵を描いてもらったり、散文を磨くのを手伝ったりすることができました。ただし、トレーニングに使用したデータによって制限されており、外部の世界にアクセスできないため、その範囲は限られていました。
モデル コンテキスト プロトコル (MCP): その後、MCP が参入し、チャットボットが外部サービスにアクセスするための AI ネイティブ API を提供しました。突然、チャットボットは、インターウェブで物事を調べたり、ドキュメントを作成したりするなど、私たちのためにさらに多くのことをできるようになりました。
エージェント : エージェントはさらにレベルを上げて、バイブコーディングなどを可能にしました。 Y

あなたはエージェントを使用していて、それに気づいていなかったかもしれません。エージェントは、モデル (例: Claude Sonnet)、命令、ツール (例: MCP)、およびエージェント ループで構成されます。エージェントのループ サイクルは次のようになります: 観察 → 理由 → 行動 → 評価。エージェントは、目標に到達するまでこれに従います。たとえば、説明書で定義されている成果物などです。
ハーネス: ハーネスは、エージェントの周囲にインフラストラクチャを追加します。これはエージェントの運用ランタイムであり、エージェントをサポートするインフラストラクチャを提供します。メモリ管理、可観測性、ライフサイクル管理などを行います。 Goose 、Claude Code、GitHub Copilot などのツールは、エージェントとハーネスの両方として機能します。ただでさえ混乱しているトピックをさらに混乱させるだけです。 🫠💀
(*) そうですね… AI は数十年前から存在しています。
冒頭で述べたように、私は自律型 AI ワークフローを試してみたかったのです。しかし、なぜ？
AI エージェントを使用したことがあるなら、あなたは自律型エージェントを使用したことがあるはずです。エージェントは、「エージェント ループ」 (上記の定義セクションを参照) を通じて、最終目標* を達成するまで推論、反復、軌道修正を行います。
開発エージェントが 1 人いるのは素晴らしいことです。しかし、エージェントのチーム全体がいて、各エージェントが人間の介入なしでソフトウェア開発ライフサイクル (SDLC) のさまざまな側面を処理するための特定のスキルを持っているとしたらどうなるでしょうか??
それはどのようなものでしょうか？それは実現可能でしょうか？これを実現するにはどのようなツールを使用できますか?
それが私が知りたかったことです。
(*) そうですね…ほとんどの場合。場合によっては、無限ループにはまってしまうこともあります。
自律型 AI ワークフローの実験では、次のツールを使用することにしました。
Paperclip は AI エージェント オーケストレーターです。それはエージェントの会社を持つという考えに基づいて組織されています。少なくとも 1 つの会社を作成する必要があり、各会社には少なくとも 1 つの ag が必要です

ent、CEO エージェント。
You can organize your company however you like.たとえば、CEO のみが会社を設立し、CEO が唯一の開発者としても機能する場合があります。素晴らしいとは言えず、Paperclip の目的に反するようなものですが、それは完全に可能です。または、特定のスキル、レポート階層、引き継ぎを備えたエージェントのチームを作成することもできます。ここで Paperclip が威力を発揮します。
Paperclip エージェントは AGENTS.md ファイルで定義されており、次のようなものが含まれます。
スキル (Paperclip 組織で利用できるようにしたいさまざまな SKILLS.md を登録する必要があります)
アーティファクトをどこに保存するか、どのアーティファクトを生成するか
エージェント間のコラボレーション (エージェントが誰から受け取るか、誰に引き継ぎ、誰とコラボレーションするか)
さらに、Paperclip を使用すると、目標を定義し、プロジェクトを作成し、エージェントをプロジェクト タスクに割り当てることができます。目標をプロジェクトに関連付けることができ、プロジェクト内で課題を作成してエージェントに割り当てることができます。
これらすべてが優れた Web インターフェイスにきちんとパッケージ化されています。
前に述べたように、私はエージェントのチームに私の指示に従ってもらいたかったのです。この分野で多くの仕事をしてきた同僚でチームメイトの Henrik Rexed と話し合った後、Paperclip に BMAD エージェントをセットアップすることにしました。実際、Papreclip で BMAD をセットアップするための公式ドキュメントがないため、私は彼のリポジトリ Papreclip-Bmad-Crew を探索の出発点として使用しました。
BMAD は、ソフトウェア開発のための AI エージェント スキルを提供するツールです。各エージェントは、アジャイル ソフトウェア開発チームのさまざまな役割/ペルソナにマッピングされた一連のスキルを持っています。私は以前に BMAD を使ったことがあり、AI 支援ソフトウェア開発に BMAD を使用した経験がとても気に入りました。
私はエージェントの基礎となる LLM として Claude Sonnet を選択しました。 Sonnet は非常に強力なモデルであり、トークンを使い切ることはありません

オーパスのように。
✨ 簡単に言うと、Paperclip が AI エージェントを管理し、BMAD がエージェントの基本スキルを提供し、クロードが作業を行います。
Henrik のリポジトリのエージェント定義を使用して、チームは次のように構成されました。
CEO：組織を管理します。
クルーマネージャー (CTO): 開発チームを管理します。 CEOに報告します。
開発チーム:
ウィンストン: アーキテクチャと実装
メアリー: 調査と市場分析
Amelia: 開発者とコードレビューアーの両方を務める二重人格
John: ユーザーのニーズを製品の要件に変換するプロダクト マネージャー
ストーリーライター：ブリッジズの商品企画・開発実行
テストアーキテクト: テストの自動化と品質保証
チャレンジャー: 物事を批判的な目で見ます。
上の図では、O11y エンジニア エージェントと DevOps エンジニア エージェントも存在することがわかります。私個人としては、小さなサイド プロジェクトにはこれらを使用しませんでしたし、BMAD スキルには対応していません。プロジェクトの 1 つでそれら (および他の Paperclip 対応エージェント) を活用したい場合は、Henrik の共有可能な Paperclip エージェントの GitHub リポジトリをチェックしてください。
この設定をテストするために、アプリのアイデアを思いつきました。 Geeking Out というポッドキャストを持っています。 Simplecast と呼ばれるツールを使用して、エピソードをさまざまなポッドキャスティング プラットフォーム (Apple Music、Spotify、Amazon Music など) に公開しています。 YouTube にもエピソードを公開しています。両方のツールで統合されたポッドキャスト統計を表示することはできないため、YouTube と Simplecast からポッドキャストの統計を 1 つのダッシュボードに取得するツールを作成すると便利だと思いました。
そして、そのプロジェクトを念頭に置いて、私は出発しました！
最初の試み: Winston がアプリを作成してくれました
Paperclip で BMAD スタッフを設定した後、Paperclip で問題を作成し、それを Winston に割り当てました。Winston は、ご存知のとおり、建築家エージェントです。問題は

食べた：
私の YouTube チャンネル (https://youtube.com/@geekingout_pod) と https://geeking-out.simplecast.com でホストされているポッドキャストから統計を収集したいと考えています。
すべての統計情報を 1 か所で確認し、PDF またはスプレッドシート (両方のオプション) にエクスポートできるアプリを構築したいと考えています。
YouTube ビデオとポッドキャスト エピソードの両方について収集する統計情報の推奨事項と、それらを表示する最適な方法を教えてください。
私たちは何度もチャットを交わし、彼は私のために素晴らしいアプリを考え出しました。
くそー、ウィンストン…よくやった。私はとても幸せでした…Paperclip が提供するものを実際に活用していないことに気づくまでは。たった 1 人のエージェントとしか関わっていないのに、さまざまな役割を持つ自律エージェントの会社全体がタスクを互いに引き継ぐことに何の意味があるのでしょうか?
それに…とにかく他のエージェントはどこにいたの？
混乱した私は、エージェントの引き継ぎがないことについてヘンリックにメッセージを送りました。彼は、私が次のことを行う必要があると（正しく）指摘しました。
BMAD手法に従う必要があることを明示してCEOに割り当てます。
2 番目の試み: エージェントのチームと協力する
わかりました。これを適切に行う時期が来ました。私は Paperclip + BMAD 環境を破壊し、アプリの作成も含めて最初から始めることにしました。
これをより Paperclip ネイティブな方法で実行したかったので、2 回目は次のようにしました。
Paperclip で「複数のソースからの Geeking Out Podcast の統計を 1 つのダッシュボードに表示する」という目標を作成しました。
Geeking Out Podcast Stats Dashboard というプロジェクトを作成しました。次に、上記の目標をプロジェクトに追加しました。
プロジェクトの要件を含むマークダウン ファイルをリポジトリに作成しました。それには以下が含まれていました:
目標
開発に使用する言語とフレームワーク
新しいプロジェクト内で新しい課題を作成し、それを CEO に割り当てました。今回の私の要望は、

ts はリポジトリ内のマークダウン ファイルにカプセル化されており、最初に使用したプロンプトよりもはるかに詳細な情報が含まれていました。私の新しいプロンプトは次のとおりです。
ファイル /workspaces/devrel-toolkit/requirements/podcast-stats-requirements.md から要件を実装します。
そして…それは惨めに失敗した。
CEO はその作業を CTO に割り当て、CTO が実装作業を進めました。なんだ…？？
腹が立ったので、私は作業を中断し、CTO エージェントの定義に反して、なぜ CTO が実装を行っているのかと尋ねました。すると、「あなたは分析、アーキテクチャ、実装の作業を自分で行うのではありません。あなたの仕事は、入ってくる作業を受け取り、それを適切な BMAD フェーズ リードにルーティングし、チケット主導のワークフローを動かし続けることです。」と明確に述べられました。
それは間違いだったと認めた。素晴らしい。しかし、適切な人材に割り当てる代わりに、新しいエージェントを「雇用」（作成）しようとしました。何が起こったのでしょうか??
私はまた激怒し、すでにアーキテクトがいるのに、なぜ CTO は新しいエージェントを雇おうとしたのですか?と指摘しました。
わかりました。我々は軌道に戻りました。しかし、物事が間違った方向に始まったという事実が本当に気に入らなかったので、タスクをキャンセルし、それまでに行っていた作業をロールバックして、もう一度最初からやり直しました。
しかし、それを行う前に、なぜエージェントは AGENTS.md で定義された適切な BMAD ハンドオフ フローに従っていなかったのでしょうか?
次に、要件ファイルを確認して、何かが欠けていることに気付きました。エージェントに BMAD フローを使用するように指示する指示です。クソ。
そこで、ファイルの最後に次のスニペットを追加しました。
## あなたの働き方 (必ず従ってください)
* BMAD 手法を使用する
* エージェントに作業を割り当てるときは、「AGENTS.md」で定義されたハンドオフに従ってください。これは、すべての下流エージェントに適用されます。
* BMAD アーティファクトを作成して `/workspaces/devrel-toolkit/_bmad-output/_bmad-ou に保存します

エージェントの `AGENTS.md` で指定されている命名規則に従って、作業を行うときは tput/geeking-out-stats` を実行してください。
* プロジェクトの作業ディレクトリ: `/workspaces/devrel-toolkit/geeking-out-stats`
その後、うまくいきました！エージェント間のハンドオフが確認できました。各エージェントは、元の問題に付随するサブ問題を作成しました。作業エージェントが完了すると、AGENTS.md の定義に従って、次のエージェントに引き継がれます。それは奇跡としか言いようがありませんでした。天使たちの歌声を合図します。
エージェントは熱心に作業し、30 分後に最初のパスを完了しました。
彼らの仕事を検査する時が来ました。いくつかのバグと使いやすさの問題を見つけたので、次のようにしました。
バグのリストを含む新しいマークダウン ファイルを作成しました
プロジェクトで新しい問題が発生しました
対処する必要があるバグを説明したマークダウン ファイルを添付しました
BMAD ワークフローに従うことに関する部分がすべてのマークダウン ファイルの一部であることを確認しました。
バグ修正サイクルを終える前に、チャレンジャーを呼び出して、すでに構築されているものに問題が見つかるかどうかを確認しました。これには、セキュリティの問題、不適切なアーキテクチャ、欠落したテストなどが含まれる可能性があります。潜在的な問題を特定した後、チャレンジャーは、すぐに修正する価値のある問題と、後で延期できる問題を決定し、必要な作業を追跡できるようにプロジェクトの問題をオープンしました。

[切り捨てられた]

## Original Extract

Exploring autonomous AI software development using Paperclip and BMAD

Autonomous AI Software Development: Good Idea, or Bad Idea? Adriana Villela
A collection of my thoughts, work, and everything in-between.
Autonomous AI Software Development: Good Idea, or Bad Idea?
We’re a few years deep into The Great AI Experiment, and there is still a lot of debate out there on how to incorporate AI effectively into our tech lives. On the one extreme, we have those who have embraced AI wholeheartedly. On the other extreme, we have those who refuse to do anything AI. As with any big technology disruption, the answer lies somewhere in the middle.
Like many of my peers in tech, I am still struggling to find that “right balance” of AI use in my work, and to help me, I’ve been experimenting and educating myself on various topics. My latest set of experiments have brought me to the wonderful world of autonomous AI workflows, which is what I’ll be talking about today.
But before we dig in, let’s take a little detour and do a little level-set on terminology.
Chatbot: In the beginning*, we had AI chatbots. These took the world by storm with the advent of ChatGPT, which opened the AI floodgates. Others like Claude , Copilot , and Gemini , soon followed. They were neat! We could ask them about Shakespeare’s works, get them to draw us fun pictures, and help us polish our prose. Their scope was limited, however, because they were limited by the data they were trained on, and had no access to the outside world.
Model Context Protocol (MCP): Then MCP entered the picture, providing an AI-native API for chatbots to access outside services. Suddenly, chatbots could do so much more for us, like look things up in the interwebs, and create documents for us.
Agent : Agents took things up another notch, making things like vibe coding possible. You might’ve been using an agent and didn’t even realize it. An agent is made up of a model (e.g. Claude Sonnet ), instructions, tools (e.g. MCP), and an agent loop. An agent loop cycle looks like this: observe → reason → act → evaluate. The agent follows this until it reaches its goal. For example, a deliverable as defined in its instructions.
Harness: A harness adds infrastructure around your agent. It is the agent’s operational runtime, providing the infrastructure that supports the agent. It does things like memory management, observability, and lifecycle management. Tools like Goose , Claude Code, and GitHub Copilot serve as both agents and harnesses. Just to add to confusion to an already confusing topic. 🫠💀
(*) Kinda… AI has been around for a few decades .
As I said in the intro, I wanted to play with autonomous AI workflows. But why?
If you’ve used AI agents, then you, my friend, have used autonomous agents. Agents by way of the “agentic loop” (see the definitions section above) will reason, iterate, and course correct until they have achieved their end goal*.
Having one agent for development is great. But what if you had a whole team of agents , each one with specific skills to handle a different aspect of the software development life cycle (SDLC), without human intervention??
What would that look like? Would it be feasible? What tools could I use to make this happen?
That’s what I wanted to find out.
(*) Well… on the most part. Sometimes they do get stuck in an infinite loop.
For my autonomous AI workflow experiment, I decided on the following tools:
Paperclip is an AI agent orchestrator. It’s organized around the idea of having a company of agents. You must create at least one company, and each company must have at least one agent, the CEO agent.
You can organize your company however you like. For example, you could have a company with only the CEO, who also serves as your sole developer. Not great, and kind of defeats the purpose of Paperclip, but you could totally do that. Or you can create a team of agents with specific skills, reporting hierarchies, and handoffs, which is where Paperclip shines.
Paperclip agents are defined in an AGENTS.md file, and they include things like:
Skills (you must register your various SKILLS.md that you want made available in your Paperclip organization)
Where to store artifacts and what artifacts to produce
Cross-agent collaborations (who the agent receives from/hands off to/collaborates with)
Additionally, Paperclip allows you to define goals, create projects, and assign agents to project tasks. You can associate goals to a project, and within a project, you can create issues and assign them to an agent.
All of this is packaged neatly into a nice web interface.
As I said previously, I wanted a team of agents to do my bidding. After chatting with my co-worker and teammate, Henrik Rexed , who has done a LOT of work in this area, I decided to set up BMAD agents in Paperclip. In fact, I used his repository, Papreclip-Bmad-Crew , as a starting point for my explorations, since there’s no official documentation for setting up BMAD with Paperclip.
BMAD is a tool that provides AI agent skills for software development. Each agent has a set of skills that are mapped to different roles/personas in an Agile software development team. I’ve played with BMAD before , and loved the experience of using it for AI-assisted software development.
I chose Claude Sonnet as the underlying LLM for my agents. Sonnet is a pretty powerful model, and it doesn’t burn through tokens like Opus does.
✨ In a nutshell: Paperclip manages the AI agents, and BMAD supplies the agent’s base skills, with Claude doing the work.
Using the agent definitions from Henrik’s repository , the team was structured as follows:
CEO: Manages the organization.
Crew Manager (CTO): Manages the development team. Reports to the CEO.
Development team:
Winston: Architecture and implementation
Mary: Research and market analysis
Amelia: Dual personas, serving as both developer and code reviewer
John: Product manager who translates user needs into product requirements
Story writer: Bridges product planning and development execution
Testing architect: Test automation and quality assurance
Challenger: Looks at things with a critical eye.
You may notice in the diagram above that there are also O11y Engineer and DevOps Engineer agents. I personally didn’t use them for my little side project, and they don’t map to BMAD skills. If you want to leverage them (and other Paperclip-ready agents) for one of your projects, you should check out Henrik’s GitHub repository of shareable Paperclip agents .
To test this setup, I came up with an app idea. I have a podcast called Geeking Out . I publish episodes to various podcasting platforms (e.g. Apple Music, Spotify, Amazon Music) using a tool called Simplecast . I also publish episodes on YouTube . I can’t view consolidated podcast stats across both tools, so I thought it would be useful to create a tool that pulls the stats for my podcast from YouTube and Simplecast onto a single dashboard.
And with that project in mind, away I went!
First try: Winston built my app for me
After setting up my BMAD crew in Paperclip, I created an issue in Paperclip, and assigned it to Winston, who, you may recall, is the architect agent. The issue stated:
I would like collect stats from my YouTube channel (https://youtube.com/@geekingout_pod) and from my podcast hosted at https://geeking-out.simplecast.com.
I would like build an app that lets me see all of my stats in one place and and exports them to a PDF or spreadsheet (options for both).
I would like recommendations of stats to collect for both YouTube videos and podcast episodes, and the best way to display them.
We chatted back-and forth a bunch, and he came up with a nice app for me.
Damn, Winston… nice job. I was pretty happy… until I realized that I wasn’t really taking advantage of what Paperclip had to offer. What was the point of having this whole company of autonomous agents with different roles, handing tasks off to one another, if I was only engaging with just one agent?
Also… where were the other agents, anyway?
Confused, I messaged Henrik about the lack of agent handoff. He (rightly) pointed out that I needed to have done the following:
Assign it to the CEO, explicitly saying that it needed to follow the BMAD method
Second try: Working with a team of agents
Okay. Time to do this properly. I decided to nuke my Paperclip + BMAD environment and start from scratch… including writing the app.
I wanted to do this in a more Paperclip-native way, so here’s what I did the second time around:
Created a goal in Paperclip: “Display stats for the Geeking Out Podcast from multiple sources in a single dashboard”.
Created a project called Geeking Out Podcast Stats Dashboard . I then attached the above goal to the project.
Created a markdown file in my repository with the project requirements. It included:
Goal
Language and framework used for development
Created a new issue inside the new project, and assigned it to the CEO. This time, my requirements were encapsulated in a markdown file in my repository, and had way more details than the prompt I used the first time around. My new prompt was:
Implement the requirements from the file /workspaces/devrel-toolkit/requirements/podcast-stats-requirements.md"
And then… it failed miserably.
The CEO assigned the work to the CTO, who proceeded to do the implementation work. What the… ??
Angry, I interrupted the work, and asked why the CTO was doing the implementation, contrary to how the CTO agent was defined, which clearly stated, “You do not do analysis, architecture, or implementation work yourself — your job is to receive incoming work, route it to the right BMAD phase lead, and keep the ticket-driven workflow moving.”
It admitted that it was wrong. Great. But instead of assigning it to the right person, it tried to “hire” (create) a new agent. WHAT WAS HAPPENING??
I got mad again, pointing out that we already had an architect, so why was the CTO trying to hire a new agent?
Okay. We were back on track. But I wasn’t really happy with the fact that things had started off on the wrong foot, so I cancelled the task, rolled back the work that had been done up to that point, and once again started from scratch.
But before I did that, why weren’t my agents following proper BMAD handoff flow defined in AGENTS.md ?
Then I checked my requirements file and realized that I was missing something: instructions telling my agent to use the BMAD flow. Crap.
So I added the following snippet to the end of my file:
## How you work (MUST FOLLOW)
* Use BMAD methodology
* Follow the handoffs defined in `AGENTS.md` when assigning agents to do the work. This applies to ALL downstream agents.
* Create and save BMAD artifacts in `/workspaces/devrel-toolkit/_bmad-output/_bmad-output/geeking-out-stats` as you work, following the naming conventions specified in the Agents' `AGENTS.md`
* Project working directory: `/workspaces/devrel-toolkit/geeking-out-stats`
After that, IT WORKED! I could see the agent-to-agent handoff. Each agent created a sub-issue attached to the original issue. Once the working agent was done, it would hand off to the next agent, as per its AGENTS.md definition. It was nothing short of miraculous. Cue angels singing.
The agents worked diligently, and after 30 minutes they were done with their first pass.
Time to inspect their work. I found a few bugs and usability issues, so I:
Created a new markdown file with a list of the bugs
Opened up a new issue in the project
Attached the markdown file describing the bugs that needed to be addressed
Made sure the part about following the BMAD workflow was a part of every markdown file
Before wrapping up a bug fix cycle, I invoked the Challenger to see if they could find any issues with what had already been built. This might include things like security issues, bad architecture, and missing tests. After identifying potential issues, the Challenger decided what issues were worth fixing right away, and what issues could be deferred to later, opening project issues so that we could keep track of the work it nee

[truncated]
