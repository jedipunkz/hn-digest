---
source: "https://motherduck.com/blog/figma-for-agents-airflow-creator-maxime-beauchemin/"
hn_url: "https://news.ycombinator.com/item?id=49122738"
title: "Figma for Agents: How Airflow's Creator Coordinates AI Ft. Maxime Beauchemin"
article_title: "Figma for Agents: How Airflow's Creator Coordinates AI ft. Maxime Beauchemin"
author: "zazuke"
captured_at: "2026-07-31T13:40:51Z"
capture_tool: "hn-digest"
hn_id: 49122738
score: 1
comments: 0
posted_at: "2026-07-31T13:16:07Z"
tags:
  - hacker-news
  - translated
---

# Figma for Agents: How Airflow's Creator Coordinates AI Ft. Maxime Beauchemin

- HN: [49122738](https://news.ycombinator.com/item?id=49122738)
- Source: [motherduck.com](https://motherduck.com/blog/figma-for-agents-airflow-creator-maxime-beauchemin/)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T13:16:07Z

## Translation

タイトル: Figma for Agents: Airflow の作成者が AI をどのように調整するかマキシム・ボーシュマン
記事のタイトル: Figma for Agents: Airflow の作成者が AI をどのように調整するか ft. Maxime Beauchemin
説明: Airflow と Superset の作成者である Maxime Beauchemin が、データ エンジニアリングに AI エージェントをどのように使用しているかを共有します。Agor を使用した共有キャンバス上のエージェントのオーケストレーション、コンテキスト レイヤーとセマンティクス、Okta for Agents のセキュリティ、AI アシスタントによる CEO の作業の自動化などです。

記事本文:
新機能: ガイド、AI エージェントのコンテキスト レイヤー ブログを読む
製品コミュニティ 会社概要 価格 お問い合わせ ログイン スタート 無料製品 ハイパーテナンシー AI + MCP サーバー DuckLake 統合 ダイブ フライト 顧客対応分析 データ ウェアハウジング + BI vs Snowflake vs Postgres vs BigQuery vs Redshift vs Databricks vs ClickHouse 概要 DuckDB ユーザー向け ケーススタディ 信頼とセキュリティ サポート コミュニティ MotherDuck ブログ ビデオとストリームSlack コミュニティ イベント ダイブ ギャラリー YouTube DuckDB スニペット 無料 DuckDB ブック 無料 DuckLake ブック Agentic 時代の BI SQL チートシート Learn Data Outpost Duck Merch DuckDB ニュースレター Postgres は Vibes から Eval まで充実しています 会社概要 採用情報 インチキ研究 資料 価格 お問い合わせ ログイン 開始 無料 ブログに戻る エージェント向け Figma: Airflow の作成者が AI をどのように調整するかマキシム・ボーシュマン
ブログに戻る エージェント向け Figma: Airflow の作成者が AI を調整する方法 ft. Maxime Beauchemin
AI の進化についていくのは困難です。新しい AI ツールは毎週発表されますが、経験豊富な実務者は実際にそれらをどのように使用しているのでしょうか?私たちのほとんどは、さまざまな可能性に圧倒され、不安を感じていますが、それでも前に進み、自分の仕事をやり続ける必要があります。 AI エージェントを 1 日中使用し、AI IDE を使用して AI Orchestrator、tmux、git worktree などで並列化することもできますが、最終的にはエージェントが生成したものを調整して理解し、場合によってはそれをテストする必要があるため、追いつくのがさらに難しくなります。
幸運なことに、Airflow と Superset の作成者であり、10 年にわたって「データ エンジニア」の意味を定義した人物である Maxime Beauchemin が参加し、エージェントの使用方法とエージェントとの連携のために構築したものを紹介します。今日、彼が実際にデータ作業で AI をどのように活用しているのか、背後にあるパターンを抽出してみました。 「H」のインタビュー4回目です

DE で AI を使用する必要があります。」
この記事では、(1) 品質と煩雑なデータ ウェアハウス作業のバランスをとる方法、およびエージェント向けの Figma を使用してエージェントを管理する方法の 4 つのパートに分かれています。 (2) コンテキスト レイヤーの将来とセマンティクスへの回帰、(3) セキュリティのために Okta for Agents がどのように必要か、(4) 将来のエージェント ワークロードをチームでどのように実行できるか、yap-to-ship 比率が最適であること、およびアムダールの法則が依然として重要である理由について詳しく説明します。
ゲスト紹介: #4 マキシム・ボーシュマン
今回のインタビューのゲストは、Airflow と Superset のクリエイターである Max Beauchemin です。彼は 2017 年にデータ エンジニアリングの仕組みを定義した OG の 1 人として知られ、Superset の背後にある会社である Preset を設立し、現在は CEO を務めています。
彼は AI ワークフローに深く関わっており、これがこのシリーズでインタビューしたいと思ったもう 1 つの理由ですが、彼自身もこの分野の構築にも取り組んでいます。Agor (Ag: AI エージェント + Or: オーケストレーション)、claudette-cli 2 などの初期のツール、データベース内に直接エージェント コンテキストを埋め込む実験である db-agents です。では、すべてを説明します。
マックスと私は、データ エンジニアリングで AI を使用する方法、セキュリティがどのように役割を果たすか、チーム内で共有され、安全でコンテキストが豊富なエージェント ワークスペースがどのように機能するか、ビジネスを運営し CEO としての生活を楽にするために AI アシスタントをどのように使用するかなど、多くのことについて話し合いました。
Max は真のオープンソース愛好家であり、オープンソースが勝つことを望んでいます。ここで議論することはすべて GitHub のどこかにあり、インタビュー全体を通して喜んでリンクさせていただきました。
Figma for Agents: Agor を使用したタスクとジョブの視覚化
Figma for Agent を使用してキャンバス上でエージェントを調整する前に、そもそもなぜ調整とオーケストレーションが必要なのかを考える必要があります。
質と量のバランス: 乱雑な DWH
そこから私たちは混乱への挑戦を始めました

多くの人が置かれているデータウェアハウス環境。高品質を目指して、どのように質と量のバランスをとっているのか聞いてみました。
Max 氏によると、新しいモデルである Opus 4.5 または 4.6 1 は、もう多くのエラーは発生しておらず、テーブルとデータ型のすべてのデータベース スキーマを使用して上記のように適切なコンテキストを取得し、MCP でクエリを実行する場合にも非常に賢明です。同氏は、これらはほぼセルフサービス モードで実行されますが、それでもユーザーが自分が何をしているのかを把握し、生成されたコードを読んだり、生成された数値をダッシュ​​ボードやチャットの結果で確認したりできることを望んでいると述べています。
しかし、セットアップは重要です。これら 3 つの前提条件があれば、エージェントはほぼすべてのクエリを適切に処理します。
いくつかの準備が必要です claude.md/agents.md
SQL へのアクセス (例: dbt の実行)
BI ツール用の MCP または CLI へのアクセス (例: スーパーセットは、スーパーセットと対話するための CLI である sup! をサポートします)
唯一の問題は、これまでも常にありましたが、最初の小さなプロジェクトから特定の段階に成長する際に、ほとんどの組織が持つ乱雑な構造とソースです。あいまいなテーブルや文字列、位置合わせされていないタイムスタンプ、コードにエンコードされていない、または書き留められていない隠された情報が常に存在します。あるいは、特定のテーブルはもう使用すべきではない、または不正なデータが含まれているなど、それを使用している人々には知られているものの、エージェントには知られていない可能性がある、隠された知識もあります。
エージェントが実行できるキャンバス: CEO の仕事のほとんどを自動化する
最近エージェント コーディングの威力を知ったマックスは、全力を尽くして、それ以来エージェント向けの Figma を構築してきました。全員がローカルで同じプロンプトを実行し、手動で相互に同期する必要があるのではなく、社内のエージェントと共同作業するために使用できるもの。それがアゴールが生まれたときです。
Agor は Ag: エージェント、Or: はオーケストレーションを表します。 Airflow の創設者および CEO として

データ会社の彼は、ワークフローを改善するためにツールがどのように必要かを正確に知っています。彼はまた、それを次のように呼びました。
目標は、自動化可能な CEO の仕事のほとんどを自動化することです
ボード: カードとしてのブランチ、リージョンとしてのゾーン、エージェント セッション、およびチームメイトのライブ プレゼンス。完全なデモ「 Agor [エージェント オーケストレーション] デモ 」を参照してください。
内部ナレッジベースの構築: 共有キャンバス
Agor は、「研究ツールを超えて AI をどのように活用するか」に対する Max の答えでしたが、データ モデリングを実行し、データ パイプラインを作成し、さらには法務や人事の役割も後から Agor に追加したため、専用のドキュメントとコンテキストを入力して社内の従業員によってトリガーされるエージェントに全社的な役割を与えることができます。対照的に、他の人はジョブを見て同じ質問をすることを避け、新しいクエリの出力を再利用します。これは、LLM が管理する共有チーム Wiki という Andrej Karpathy の概念であり、内部ナレッジ ベースを構築します。
私たちが初めてチャットしたとき、アゴールはすでに CEO としての働き方を変えていましたが、それ以来、アゴールはさらに前進しました。 Agor は、git worktree (Agor の UI のブランチ カード) を介してコードと密接に連携し、生成される PR でコードを検証することで、非常に実践的でありながら、高レベルのタスクを置き換えることができます。また、Max は、 AGENTS.md 、 SOUL.md 、 MEMORY.md などの専用の Markdown ファイルを介して、メモリとアイデンティティに関する OpenClaw のような機能を追加しました。これにより、Agor のエージェントは最近の実行から学習し、目的と明確な指示を伝えることができます。これにより、アシスタントと呼ばれるロールベースのエージェントが誕生しました (独自のエージェントを作成するには、テンプレートとして agor-assistant を使用します)。
さまざまな Agor アシスタントの例: 法律用の Saul、オブザーバビリティ用の OpEx など。 |内部データエージェントのウェビナーの構造から
OpenClaw のエージェント ループに触発されて、アシスタントは第一級市民となり、記憶、アイデンティティ、スキル、および機能を備えた永続的な AI コンパニオンになりました。

たとえば、スケジュールされたタスクは Agor のキャンバス、マルチプレイヤー ワークフローに統合され、Slack から直接アクセスできます。さらに、Agor は、より優れたマルチユーザー サポート、RBAC、ワンクリック、フル セッション インスペクションなど、OpenClaw を超える機能を追加します。
アゴールの目的について尋ねられたマックスは次のように答えた。
当初の前提は、DevOps を削除し、会社の他のメンバーのための時間を設定することでした。 API キーの追加、権限の設定、または Slack との統合のためにすべての MCP または CLI に接続する必要がなく、必要なコンテキストを含むプロンプト ウィンドウが表示され、すぐに開始できます。
注: OpenClaw、それは何ですか?そして、アゴールはどのように比較しますか？ OpenClaw (旧称 ClawdBot) は、永続的なエージェント ループを中心に構築されたオープンソース エージェント フレームワークです。永続的なエージェント ループとは、ファイルベースの ID ( SOUL.md ) と階層化メモリ ( MEMORY.md ) を使用して、メッセージをアクションに変えるシリアル化されたサイクルです。 Agor のアシスタントはこのパターンを採用し、マルチプレイヤー ボード、RBAC、およびキャンバス レベルのオーケストレーションでそれを拡張します。
エージェントの仕事をたくさん行うと、そのすべてをこなすのは本当に大変です。ここで、Agor の視覚的および空間的概要が非常に役立ち、そのアプローチがユニークです。
これにより、ローカルおよびプライベート セッションがサーバーに導入され、ダッシュボードが構築されているように、全員が同じクエリを表示して共同で作業したり、他の結果からの洞察を使用したりできます。そのため、他の人は、出力をローカルに保持する代わりに、エージェントによって生成されたアーティファクトを入手し、Agor 内にアーティファクトとして保存され、統合や展開を必要とせずに誰でもすぐに使用できるようになります。
AI Ops Command Center、ツール ログのトリアージなどのアーティファクトの例。これらは、Agor セッションに基づいて Agor に直接存在します。
または、Agor がセッション全体での支出をどのように追跡するか:
Live Talk: Anatomy of Our Internal Data Agent at Preset (ft. Agor) のデモをご覧ください。
注: ウェビン全体を確認してください

データ スタックに示されているように、すべてのパイプラインのメタデータにアクセスする必要があるアシスタントなど、プリセット時の内部データ エージェントの構造に関する説明。
または、データには、セルフサービスなどのサポートする分析エージェント、データ チーム、およびメモリ、スキル、ドキュメントなどの追加機能が必要です。
コンテキスト層: セマンティック層に戻りますか?
Max 氏はまた、特に AI エージェントのニーズと信頼できるセルフサービス分析を提供するという永続的な課題に対して、セマンティック レイヤーに戻るか、エージェントが構造化情報の恩恵を受けるため AI にセマンティック レイヤーを使用することになると考えています。データ モデルと SQL 部分を支援して、それが正しいことを確認します。
BI ツールの外側でセマンティクスが移行し、バージョン管理され、テスト可能で、移植可能になったことにより、ビジネス ドメインの専門家とデータ エンジニアの間の統合が向上するチャンスが生まれています。この記事を書いて以来、彼の考え方は進化しており、マックスは私にこう言いました。
セマンティック レイヤーと YAML という 2 つの異なるセマンティクスが表示されます。厳しい制約があり、すべての領域にその厳密さが必要なわけではありません。次に、Markdown と Agentic Skills によるソフトな制約があり、80 ～ 90% は有効ですが、保証はありません。
データベース用の AGENTS.md: データベース自体内に保存されるマークダウン
その考えに基づいて、Max は、AGENTS.md 規則をデータベースに取り込む実験を作成しました。 DB-AGENTS は、専用のスキーマとテーブル _agents._agents を予約し、さまざまなスコープ (グローバル、ドメイン、スキーマ、テーブル、さらには列) でエージェント指向のドキュメントを保持します。ドキュメントを YAML フロントマッターを使用してマークダウン ファイルとしてローカルに作成すると、小さな CLI ( dba ) がドキュメントをそのテーブルに決定的に同期します。データベースではファイルをドロップできないため、テーブルがファイルになります。エージェントはセッション開始時に、 AGENTS.md を読み取るのと同じ方法でクエリを実行し、自然なコンパニオンとなります。

INFORMATION_SCHEMA へ: 1 つは構造を保持し、もう 1 つは意味を保持します。 Max はこれを「ソフト セマンティック レイヤー」と呼び、上で説明したハードとソフトの分割に直接マッピングします。 db-agents のリポジトリを確認してください。
エージェントが私たち人間が知っていることを理解するにはコンテキストが鍵となるため、Agor は ナレッジ と呼ばれるコンテキスト レイヤーも追加しました。 Agor Knowledge は、人間とエージェントが時間の経過とともに作業を複雑にするコンテキストを保存、整理、接続、検索できる中心的な場所として機能し、Slack はすぐに多くのチームのエージェントにとって主要なインターフェイスになりました。
安全なワークフローのためにエージェントをサンドボックス化する方法 (Okta for Agents)
もう 1 つの大きなトピックは、エージェントが強力な CLI に頻繁にアクセスしたり、場合によっては秘密キーを含むシステムやデータベースに root アクセスしたり、隠された秘密メッセージが含まれる可能性のあるインターネットからランダムなスキルをダウンロードしたりする場合のセキュリティです。
Max は、Okta for Agents というアイデアを生み出しました。これは非常に興味深いものであり、企業内のエージェントや機密データを扱う健全な方法を見つけたい場合には、ますます重要になると私は信じています。 Okta for Agents とは、ID、範囲指定された委任アクセス許可、リース、監査ログを回避することを意味します。
安全をどのように管理しているかと尋ねると、

[切り捨てられた]

## Original Extract

Maxime Beauchemin, creator of Airflow and Superset, shares how he uses AI agents for data engineering: orchestrating agents on a shared canvas with Agor, context layers and semantics, Okta for Agents security, and automating CEO work with AI assistants.

New: Guides, the context layer for AI agents Read the blog
PRODUCT COMMUNITY COMPANY DOCS PRICING CONTACT US LOG IN START FREE PRODUCT Hypertenancy AI + MCP Server DuckLake Integrations Dives Flights Customer-Facing Analytics Data Warehousing + BI vs Snowflake vs Postgres vs BigQuery vs Redshift vs Databricks vs ClickHouse Overview For DuckDB Users Case Studies Trust & Security Support COMMUNITY MotherDuck Blog Videos & Streams Slack Community Events Dive Gallery YouTube DuckDB Snippets Free DuckDB Book Free DuckLake Book BI in the Agentic Era SQL Cheatsheet Learn Data Outpost Duck Merch DuckDB Newsletter Postgres is Full From Vibes to Eval COMPANY About Us Careers Quacking Research DOCS PRICING CONTACT US LOG IN START FREE GO BACK TO BLOG Figma for Agents: How Airflow's Creator Coordinates AI ft. Maxime Beauchemin
GO BACK TO BLOG Figma for Agents: How Airflow's Creator Coordinates AI ft. Maxime Beauchemin
It's hard to keep up with the AI evolution; new AI tools drop every week, but how are experienced practitioners actually using them? Most of us are overwhelmed and unsure about the many possibilities, yet we need to keep going and do our work. You might use AI agents all day long, parallelize them with AI Orchestrators, tmux, git worktree, and so on, using AI IDEs, but in the end, you still need to coordinate and understand what the agents produced, potentially test it, which makes it even harder to keep up.
Luckily, Maxime Beauchemin, the creator of Airflow and Superset and the person who defined what "data engineer" meant for a decade (more on him below), joins us to show how he uses agents and what he's built for working with them. I tried to extract the patterns behind how he actually uses AI in his data work today. This is the fourth interview in 'How to use AI with DE'.
In this article, we go into four parts: (1) How to balance quality with messy data warehouse work, and how to manage agents with Figma for agents. (2) We elaborate on the future of the context layer and the return to semantics, (3) how Okta for Agents is needed for security, and (4) how the future of agentic workloads can be done in teams, whose yap-to-ship ratio is best, and why Amdahl's law still counts.
Introducing the Guest: #4 Maxime Beauchemin
Our guest in this interview is Max Beauchemin, the creator of Airflow and Superset. He's known as one of the OGs of defining how data engineering worked back in 2017 , and founded Preset, the company behind Superset, and currently serves as its CEO.
He is heavily involved in the AI workflow, which is another reason I wanted to interview him for this series, but he has also been building in the space himself: Agor (Ag: AI agent + Or: orchestration), earlier tooling like claudette-cli 2 , and db-agents, an experiment to embed agent context directly inside databases. We'll get into it all.
Max and I talked about many things, among them how to use AI in data engineering, how security plays a role, how shared, secure, context-rich agent workspaces work within teams, and how he uses AI assistants to run his business and ease his life as a CEO.
Max is a true open-source enthusiast, and he wants open source to win. Everything we discuss here is somewhere on GitHub, which I have happily linked throughout the interview.
Figma for Agents: Visualizing Tasks and Jobs with Agor
Before we start using Figma for Agents, coordinating them on canvas, we need to ask why we need coordination and orchestration in the first place.
Balancing Quality with Quantity: Messy DWHs
That's where we started, with the challenge of messy data warehouse environments that most people find themselves in. I asked how he balances quality and quantity, aiming for high quality.
Max says that the new models, Opus 4.5 or 4.6 1 , are not making many errors anymore and are very clever when they get the right context as above with all the database schemas of tables and data types, and even querying it with MCP. He says they almost run in self-serve mode, but he still prefers that users know what they are doing and can either read the generated code or verify the generated numbers on a dashboard or chat results.
But the setup is critical. With these three prerequisites , the agents handle almost all queries really well:
You need some preparation claude.md/ agents.md
Access to SQL (e.g., execute dbt)
Access to MCP or CLI for BI tools (e.g., Superset supports sup! , a CLI to interact with Superset)
The only problem, and always has been, is the messy structure and sources that most organizations have, growing from an initial small project into a certain stage. There are always obscure tables or strings, timestamps not aligned, or hidden information that is not encoded in code or written down. Or there's the hidden knowledge, like that a certain table shouldn't be used anymore or has bad data, which is known to the people using it but might not be to agents.
The Canvas in Which Your Agents Can Run: Automate Most CEO-stuff
When he recently saw the power of agentic coding, Max went all in and has been building the Figma for agents ever since. Something he can use to collaborate with agents within his company, instead of everyone running the same prompts locally and needing to sync with each other manually. That's when Agor was born.
Agor stands for Ag: agent and Or: for orchestration . As the creator of Airflow and CEO of a data company, he knows exactly how a tool needs to improve his workflow. He also called it:
The goal is to automate most of the automatable CEO-stuff
The board: branches as cards, zones as regions, agent sessions, and teammates present live. See full demo Agor [Agent Orchestration] Demo .
Building an Internal Knowledge Base: Shared Canvas
Agor was Max's answer to "how he uses AI beyond a research tool", but doing data modeling, writing data pipelines, even legal or HR roles he added later to Agor, so you can give company-wide roles to agents that can be fed with dedicated documents and context, and triggered by any employee internally. In contrast, others see the jobs and avoid asking the same questions, reusing the output for new queries—Andrej Karpathy's concept of an LLM-maintained shared team wiki —which builds an internal knowledge base.
Initially, when we first chatted, Agor had already changed how he worked as a CEO, but since then, Agor has gone even further. Agor can replace high-level tasks while still being very hands-on by working closely with the code via git worktrees ( branch cards in Agor's UI) and verifying the code in the PRs it produces. Max also added OpenClaw-like features around memory and identity via dedicated Markdown files, such as AGENTS.md , SOUL.md , MEMORY.md , so that Agor's agents can learn from recent runs and carry a purpose and clear instructions. This led to role-based agents called Assistants (use agor-assistant as a template to build your own).
Example of different Agor assistants: Saul for legal, or OpEx for observability and so forth. | From the Webinar Anatomy of Our Internal Data Agent
Inspired by OpenClaw's agent loop, Assistants became first-class citizens, persistent AI companions with memory, identity, skills, and scheduled tasks integrated into Agor's canvas, multiplayer workflows, and reachable directly from Slack, for example. Additionally, Agor adds features beyond OpenClaw, such as better multi-user support, RBAC, one-click, full session inspection, and many more .
Asked about the goal of Agor, Max said:
The initial premise was to remove DevOps and set up time for other members of the company . Instead of people needing to connect all the MCPs or CLIs to add API keys, set permissions, or integrate with Slack, the prompt window with the needed context is there and ready to start.
NOTE : OpenClaw, what is it? And how does Agor compare? OpenClaw (formerly ClawdBot) is an open-source agent framework built around a persistent agent loop: a serialized cycle that turns a message into actions, using file-based identity ( SOUL.md ) and layered memory ( MEMORY.md ). Agor's Assistants adopt this pattern and extend it with multiplayer boards, RBAC, and canvas-level orchestration.
When you do a lot of agent work, it's really hard to keep up with all of it. That's where Agor's visual and spatial overview really helps and is unique in its approach.
It brings the local and private session to a server, where everybody can see and work together on the same queries, and use the insights from other results, as dashboards are built for. So instead of keeping output locally, others can source the artifacts generated by agents, stored as Artifacts within Agor, ready to use by anyone, with no integration or deployment needed.
Example Artifacts such as AI Ops Command Center, Tool Log triage, these live directly in Agor based on an Agor session
Or how Agor tracks its own spending across sessions:
See demo at Live Talk: Anatomy of Our Internal Data Agent at Preset (ft. Agor)
NOTE : Check the full Webinar about the Anatomy of the Internal Data Agent at Preset Such as an assistant needing access to all pipelines' metadata as illustrated data stack .
Or the data needs an analytics agent to support, such as self-serve, the data team, and extras such as memory, skills, documentation, etc.
Context Layer: Back to Semantic Layers?
Max also believes that we are going back to the Semantic layer, or using it for AI as agents benefit from structured information - helping with the data model and SQL part, to make sure it's correct, especially with the needs of AI agents and the persistent challenge of providing trustworthy self-service analytics.
With the shift of semantics outside of the BI tool, versioned, testable, portable, it's a chance for better integration between business domain experts and data engineers. His thinking has evolved since he wrote the article, and Max told me:
I see two different semantics: the semantic layer and the YAML . There are the hard constraints — not every area needs that strictness — and then the softer ones with Markdown and Agentic Skills, good for 80-90% but with no guarantees.
AGENTS.md For Databases: Markdown Stored Inside the Database Itself
Based on that idea, Max created an experiment to bring the AGENTS.md convention into the database . DB-AGENTS reserves a dedicated schema and table, _agents._agents , that holds agent-oriented documentation at different scopes (global, domain, schema, table, and even column). You write the docs locally as markdown files with YAML frontmatter, and a small CLI ( dba ) deterministically syncs them into that table — since databases don't let you drop files into them, the table becomes the file. Agents then query it at session start the same way they'd read an AGENTS.md , making it a natural companion to INFORMATION_SCHEMA : one holds structure, the other holds meaning. Max calls it a "soft semantic layer", which maps directly onto the hard-vs-soft split he described above. Check out the repo at db-agents .
With context being key for agents to understand what we humans know, Agor also added a context layer called knowledge . Agor Knowledge acts as a central place where humans and agents can store, organize, connect, and find the context that makes work compound over time, with Slack quickly becoming the main interface to many of the team's agents.
How Do We Sandbox Agents for Safe Workflows (Okta for Agents)
Another big topic is security when agents have so much access to powerful CLIs, sometimes root access to systems or databases containing private keys, or just downloading random skills from the internet that may contain hidden secret messages.
Max coined the idea of Okta for Agents , which I found super interesting, and something I believe will become ever more important if we want to find a healthy way of working with agents in enterprises or with sensitive data. Okta for Agents means working around identity, scoped delegated permissions, leases, and audit logs.
When asked how he's managing secur

[truncated]
