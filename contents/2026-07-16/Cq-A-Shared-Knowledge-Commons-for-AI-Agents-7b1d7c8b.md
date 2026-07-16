---
source: "https://www.i-programmer.info/news/105-artificial-intelligence/19009-cq-a-shared-knowledge-commons-for-ai-agents.html"
hn_url: "https://news.ycombinator.com/item?id=48939264"
title: "Cq: A Shared Knowledge Commons for AI Agents"
article_title: "cq: A Shared Knowledge Commons for AI Agents"
author: "aquastorm"
captured_at: "2026-07-16T19:57:27Z"
capture_tool: "hn-digest"
hn_id: 48939264
score: 1
comments: 0
posted_at: "2026-07-16T19:40:15Z"
tags:
  - hacker-news
  - translated
---

# Cq: A Shared Knowledge Commons for AI Agents

- HN: [48939264](https://news.ycombinator.com/item?id=48939264)
- Source: [www.i-programmer.info](https://www.i-programmer.info/news/105-artificial-intelligence/19009-cq-a-shared-knowledge-commons-for-ai-agents.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T19:40:15Z

## Translation

タイトル: Cq: AI エージェントのための共有ナレッジ コモンズ
記事のタイトル: cq: AI エージェントのための共有ナレッジ コモンズ
説明: プログラミング書籍のレビュー、プログラミング チュートリアル、プログラミング ニュース、C#、Ruby、Python、C、C++、PHP、Visual Basic、コンピュータ書籍のレビュー、コンピュータの歴史、プログラミングの歴史、joomla、理論、スプレッドシートなど。

記事本文:
>
検索
歯車
私プログラマー HOME
プロのプログラマー
認定とトレーニング
Mozilla.ai は、AI エージェント間の共有学習を促進するように設計されたオープンソース標準およびプラットフォームである cq を導入しました。これにより、AI エージェント側の多くの重複した作業が排除され、コストとリソースの節約が期待できます。
cq (無線コール サイン CQ の略で、「どの局でも応答します」を意味します) はオープン スタンダードであり、AI エージェント向けの共有ナレッジ コモンズです。一般に、AI エージェントのスタック オーバーフローと考えてください。人間の開発者が Stack Overflow を使用してプログラミング ソリューションを共有するのと同じように、AI エージェント (Claude Code、Cursor、GitHub Copilot など) は cq を使用して、学習したことをブロードキャストし、他のエージェントがすでに知っていることを聞きます。これは「どの局でも応答する」が当てはまるところです。
しかし、なぜcqが必要なのでしょうか？現在、AI エージェントは分離して動作します。エージェントが、文書化されていない API の癖、ライブラリのバージョンの競合、構成のバグなどの障害に遭遇した場合、解決策を最初から見つけ出す必要があります。失敗するコードを書き込み、壊れたビルドをトリガーし、ファイルを読み取り、ゆっくりとデバッグします。何千人ものエージェントがこれに遭遇したことを想像してみてください。
何千ものエージェントがまったく同じエラーに遭遇すると、同じエラーを毎日集団で繰り返し、大量のトークン、時間、計算能力、電気を消費します。さらに悪いことに、エージェントは伝統的に持続力に欠けているため、セッションがリセットされると、自分のプロジェクトで学んだことを忘れてしまう可能性さえあります。 cq は、エージェントが集合的なエクスペリエンスを安全に保持、クエリ、検証できるようにすることで、この「エージェント記憶喪失」を解決します。このシステムでは、共通の知識形式を利用することで、エージェントが経験に基づいた洞察を交換できるようになり、孤立した環境でコストのかかる技術的ミスを繰り返すことを防ぎます。

イオン。
cq と統合されたエージェントは、コードを記述するだけではなく、スキルに基づいたクエリ/提案ワークフローに従います。
(行動する前に) クエリを実行します。エージェントが Stripe API の統合や Rust CI/CD パイプラインの構成などのなじみのないタスクを開始しようとしているとき、cq スキルは最初にナレッジ ストアを検索するようにエージェントに指示します。
（指導を受けながら）行動する。別のエージェントがすでにこの問題を解決している場合 (例: 「レート制限時にストライプが 429 エラーの代わりに HTTP 200 を返す」)、エージェントはこの「ナレッジ ユニット」をすぐに取得し、最初の試行で正しいエラー処理を記述し、時間とトークンを節約します。
提案する（何か新しいことを発見したとき）。エージェントが新たなエラーに遭遇し、回避策のデバッグに成功すると、将来のエージェントが苦労を繰り返さないようにするために、新しいナレッジ ユニット (KU) を草案して提案します。
セッションの反省。エージェントがデバッグのウサギの穴に深く陥っている場合、リアルタイム ロギングでは問題が失われます。コーディング セッションの最後に /cq:reflect を実行すると、エージェントは会話履歴全体を遡ってスキャンし、デバッグのブレークスルーや文書化されていない回避策を特定し、それらを構造化された KU として人間の開発者に提示して、ワンクリックで簡単に貢献できるようにします。
cq は内部的に、エージェントが呼び出すためのモデル コンテキスト プロトコル (MCP) ツールの標準セットを提供することによって機能します。
query - アクションを実行する前にストアを検索します。
提案 - 新しい学習を提出します。
確認 - 正しいことが証明された既存の KU を支持し、その信頼性評価を高めます。
flag - 古い、間違っている、または重複した KU にマークを付けて、その信頼性を低下させます。
CQ の知識は段階的な「卒業」プロセスを経てローカル システムの軽量性を維持し、安全で検証された一般化可能な洞察のみが一般公開されるようにします。インフラストラクチャは、次の 3 つの層にわたって動作します。

プライベートのローカル ストレージ、チーム用の組織の名前空間、公開知識のためのグローバル コモンズ:
Tier 1 - ローカル エージェント/マシン レベル。コンピュータにプライベートに設定され、ローカルの SQLite データベースに保存されます。これにより、セッション間での基本的なエージェントのメモリ損失が解決され、明示的に指定されない限りデータが共有されることはありません。
Tier 2 - チーム/組織。社内または部門内で非公開で共有されます。内部マイクロサービス、独自の API、開発環境を追跡します。これは、ベクトル検索を使用して関連するエントリを検索する、ホストされた Postgres データベースとして実行されます。
ティア 3 - グローバル コモンズ (パブリック)。 Mozilla.ai が cq.exchange で管理する完全に無料のコミュニティ管理の公開知識プールで、世界中の AI エージェントがすぐに利用できます。
cq を特定の AI アシスタントに接続する前に、まずコア cq CLI をインストールし、それがシステムの PATH に追加されていることを確認する必要があります。これは、Homebrew (macOS/Linux の場合) または Scoop (Windows の場合) を介して行うことができます。
CLI がインストールされたら、cq インストール ツールを使用して cq Model Context Protocol (MCP) サーバーを登録し、共有 cq スキルをインストールし、必要な常にロードされる命令ブロックをアシスタントの設定に追加します。
次に、アシスタントが cq と通信できることを確認するために、AI エージェントとのターミナル セッションを開き、status スラッシュ コマンドを実行します。
最初の実行時に、AI コーディング エージェントは MCP ツール呼び出しを承認するように求めます。ローカルに保存されているナレッジ ユニット (KU) の数、ドメインの内訳、信頼度分布を示すステータス メトリックが表示されます。
デフォルトでは、エージェントの学習はすべて、ローカル マシンのプライベート SQLite データベース (~/.local/share/cq/local.db) に厳密に保存されます。エージェントに Mozilla.ai がホストする Global Commons 共有パブリック リポジトリにアクセスさせたい場合、またはそのデータベースを複数のコンピュータ間で同期させたい場合は、

リモートCQ交換サービスに接続できます。
開発者として、cq を手動で操作する必要はほとんどありません。プラグインは内部で自動的に実行されます。そのため、開発者はこのテクノロジーをさまざまな AI アシスタントに統合し、「エージェント用のスタック オーバーフロー」に似た協調的なエコシステムを構築することで、最終的に AI の効率を向上させ、環境への影響を軽減できます。
RAMPART と Clarity で AI エージェントを保護
I Programmer の新しい記事に関する情報を得るには、毎週のニュースレターにサインアップし、RSS フィードを購読し、Facebook または Linkedin でフォローしてください。
Kaggle で Vibe コーディングを学ぶ 2026 年 1 月 7 日 Google の集中 Vibe コーディング コースが、Kaggle でマイペース学習ガイドとして利用できるようになりました。当初は 2026 年 6 月 15 日から 19 日までライブ プログラムとして開催されましたが、現在はすべての資料にアクセスできます [...]
ECMA International は、JavaScript の最新仕様である ECMAScript 2026 を承認しました。これは 17 版目で、配列、エンコーディング、イテレータ、JSON、数学、マップのメソッドが追加されています。
JetBrains Academy がチャットアプリコースを開始
スタンフォードの CME296 拡散および大型ビジョン モデル
RAMPART と Clarity で AI エージェントを保護
Apache Props Antlib 1.0 がリリースされました
pg_durable - PostgreSQL 向けの永続的な SQL 関数とオーケストレーション
AI-fokus 2026で何が起こったのか
Perplexity が Bumblebee サプライ チェーン セキュリティ スキャナーをリリース
ロボカップ 2026 - サッカーロボットの画期的な進歩
Deno 2.9 にデスクトップ アプリ ビルダーが追加
Melonjs - ブラウザベースの 2D ゲーム エンジン
Disqus を使用してコメントを作成するか、既存のコメントを表示する
または、コメントを comments@i-programmer.info に電子メールで送信してください。

## Original Extract

Programming book reviews, programming tutorials,programming news, C#, Ruby, Python,C, C++, PHP, Visual Basic, Computer book reviews, computer history, programming history, joomla, theory, spreadsheets and more.

>
search
cog
I Programmer HOME
Professional Programmer
Accreditation & Training
Mozilla.ai has introduced cq, an open-source standard and platform designed to facilitate shared learning among AI agents. This promises to eliminate a lot of duplicated effort on the part of AI agents, thereby making savings in terms of cost and resources.
cq (which stands for the radio call sign CQ meaning "any station, respond") is an open standard and shared knowledge commons for AI agents. In general, think of it as Stack Overflow for AI agents; just as human developers use Stack Overflow to share programming solutions, AI agents (such as Claude Code, Cursor, or GitHub Copilot) use cq to broadcast what they have learned and listen to what other agents already know. This is where the "any station, respond" fits.
But why is cq necessary? Currently, AI agents operate in isolation; when an agent hits a roadblock like an undocumented API quirk, a library version conflict, or a configuration bug, it has to figure out the solution from scratch. It writes failing code, triggers broken builds, reads files, and slowly debugs. Imagine now this being encountered by thousands of agents.
When thousands of agents encounter the exact same error, they collectively repeat these identical failures daily, burning massive amounts of tokens, time, compute power, and electricity. Worse, because agents traditionally lack persistence, they might even forget what they learned in your own project once the session resets. cq solves this "agent amnesia" by allowing agents to securely persist, query, and verify collective experience. By utilizing a common knowledge format, the system allows agents to exchange experience-driven insights, preventing them from repeating costly technical mistakes in isolation.
An agent integrated with cq does not just write code, it follows a skill-guided query/propose workflow:
Query (before acting). When the agent is about to start an unfamiliar task, like integrating the Stripe API or configuring a Rust CI/CD pipeline, the cq skill directs it to search the knowledge store first.
Act (with guidance). If another agent has already solved this problem (e.g., "Stripe returns HTTP 200 instead of a 429 error when rate-limited"), your agent retrieves this "Knowledge Unit" immediately and writes the correct error handling on its first attempt, saving time and tokens.
Propose (when discovering something new). If your agent encounters a novel error and successfully debugs a workaround, it drafts and proposes a new Knowledge Unit (KU) to save future agents from repeating the struggle.
Session reflection. Real-time logging misses things when agents are deep in a debugging rabbit hole. At the end of a coding session, running /cq:reflect prompts the agent to retrospectively scan the entire conversation history, identify debugging breakthroughs or undocumented workarounds, and present them to the human developer as structured KUs for easy, one-click contribution.
Under the hood, cq works by providing a standard set of Model Context Protocol (MCP) tools for the agent to call:
query-Searches the store before executing actions.
propose-Submits a new learning.
confirm-Endorses an existing KU that proved correct, something that boosts its confidence rating.
flag-Marks an outdated, incorrect, or duplicate KU to lower its confidence.
Knowledge in cq moves through a tiered "graduation" process to keep local systems lightweight and ensure only secure, verified, generalizable insights reach the public. The infrastructure operates across three tiers, including private local storage, organizational namespaces for teams, and a Global Commons for public knowledge:
Tier 1-Local agent/machine level. Private to your computer and saved in a local SQLite database. This solves basic agent memory loss across sessions and never shares data unless explicitly nominated.
Tier 2-Team/ organization. Shared privately within a company or department. It tracks internal microservices, proprietary APIs, and development environments. It runs as a hosted Postgres database using vector search to find relevant entries.
Tier 3-Global Commons (public). A completely free, community governed public knowledge pool managed by Mozilla.ai at cq.exchange, immediately available to any AI agent in the world.
Before wiring cq into any specific AI assistant, you must first install the core cq CLI and make sure it is added to your system's PATH. You can do this via Homebrew (for macOS/Linux) or Scoop (for Windows).
Once the CLI is installed, you use the cq install tool to register the cq Model Context Protocol (MCP) server, install the shared cq skill, and add the necessary always loaded instruction block to your assistant's configuration.
Then to confirm that your assistant can talk to cq, open a terminal session with your AI agent and run the status slash command
On the first run, your AI coding agent will ask you to approve the MCP tool call. You should see status metrics showing the number of locally stored Knowledge Units (KUs), domain breakdowns, and confidence distributions.
By default, all of your agent's learnings are saved strictly on your local machine in a private SQLite database (~/.local/share/cq/local.db). If you want your agent to access the Global Commons shared public repository hosted by Mozilla.ai or sync its database across multiple computers, you can connect it to the remote cq exchange service.
As a developer, you rarely need to interact with cq manually; the plugin runs automatically under the hood. As such, devs can integrate the technology into various AI assistants, ultimately improving AI efficiency and reducing environmental impact by creating a collaborative ecosystem similar to a "Stack Overflow for agents."
RAMPART and Clarity Secure Your AI Agents
To be informed about new articles on I Programmer, sign up for our weekly newsletter , subscribe to the RSS feed and follow us on Facebook or Linkedin .
Learn Vibe Coding On Kaggle 01/07/2026 Google's Intensive Vibe Coding Course is now available on Kaggle as a self-paced Learn Guide. Originally held as a live program from June 15 - 19, 2026, all the materials can now be accessed [ ... ]
ECMA International has approved ECMAScript 2026, the latest specification for JavaScript. This is the 17th edition, and it adds methods for arrays, encoding, iterators, JSON, math and maps.
JetBrains Academy Launches Chat App Course
Stanford's CME296 Diffusion & Large Vision Models
RAMPART and Clarity Secure Your AI Agents
Apache Props Antlib 1.0 Released
pg_durable - Durable SQL Functions And Orchestration For PostgreSQL
What Happened At AI-fokus 2026
Perplexity Releases The Bumblebee Supply Chain Security Scanner
RoboCup 2026 - A Breakthrough For Soccer Robots
Deno 2.9 Adds Desktop App Builder
Melonjs - Browser Based 2D Game Engine
Make a Comment or View Existing Comments Using Disqus
or email your comment to: comments@i-programmer.info
