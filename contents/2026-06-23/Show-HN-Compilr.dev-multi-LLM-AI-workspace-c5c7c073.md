---
source: "https://compilr.dev/"
hn_url: "https://news.ycombinator.com/item?id=48648085"
title: "Show HN: Compilr.dev, multi LLM AI workspace"
article_title: "compilr.dev — Your AI Workspace"
author: "scozzola"
captured_at: "2026-06-23T17:35:02Z"
capture_tool: "hn-digest"
hn_id: 48648085
score: 1
comments: 0
posted_at: "2026-06-23T17:11:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Compilr.dev, multi LLM AI workspace

- HN: [48648085](https://news.ycombinator.com/item?id=48648085)
- Source: [compilr.dev](https://compilr.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T17:11:51Z

## Translation

タイトル: HN を表示: Compilr.dev、マルチ LLM AI ワークスペース
記事のタイトル: compilr.dev — AI ワークスペース
説明: AI ワークスペース — 特化した AI エージェント、構造化されたプロジェクト ワークフロー、ソフトウェア、研究、ビジネスなどの永続的なコンテキスト。デスクトップ アプリ、CLI、開発者ライブラリ。
HN テキスト: こんにちは。過去 6 か月以上にわたって、私は AI エージェントを構築または使用するためのツールのエコシステム (ライブラリ、CLI、デスクトップ) を構築してきました。コーディングにどれだけ慣れているかどうかに応じて、ライブラリ (compilr-dev/agents および compilr-dev/agents-coding) を使用して独自の AI エージェントを構築する場合があります。 CLI を使用してプロジェクトを管理し、VS Code などで作業中にターミナルでエージェントを実行することもできます。また、デスクトップ アプリを使用することもできます (CLI の同じコアを共有しますが、より豊富な UI を備えています)。私は、compilr-dev エコシステムを他の同様のツールと区別しているのは、真のマルチエージェントの性質であると信じています。デスクトップ アプリに重点を置くと、事前に構築された特殊なエージェント (上記のライブラリを使用して構築) が同梱されており、各エージェントにはカスタマイズされたシステム プロンプト、スキル、専用ツールへのアクセスが備わっています。新しいプロジェクトが作成されると、専門のエージェントのチームがチーム名簿に追加され、プロジェクトの作成時にすでにエージェントを追加または削除できます (プロジェクトの作成時にカスタム エージェントを作成することもできます)。好みの LLM プロバイダー (BYOK アプローチ) を選択でき、必要に応じていつでもプロバイダーをシームレスに切り替えることができます。各エージェントは独自のコンテキストを持っていますが、プロジェクト内のエージェントも一部のコンテキストを共有します。これらは、問い合わせ (エージェント A がエージェント B に何かを尋ねることができます) またはタスクのハンドオフによって相互に対話できます。各エージェントは仕様主導のアプローチに従うように指示され、アプリ自体がユーザーに構築前に設計するよう促します。
たとえば、新しいプロを作成するとすぐに、

/design コマンドを使用すると、エージェントに質問してプロジェクトの範囲と詳細を把握することができます。その結果、エージェントはデータベース内の作業項目を追跡し、指示に従って PRD ドキュメントを作成します。すべてのファイルは、あなたとエージェントが表示および編集できます。最後に重要なことですが、トークンの消費には特別な注意が払われました。トークン = お金なので、トークンを節約することが「最も重要」です (いいえ、AI がこの文章を書いたわけではありません。私は意図的に「最も重要」という言葉を使用することにしました)。
いくつかのテクニック:
- ツールとスキルの定義は、必要な場合にのみロードされます (ネオがマトリックスにカンフー スキルをロードするなど)
- ツールの出力は、数ラウンド後にコンテキスト内でプルーニング/要約されます。
- エージェントは、大きなファイルを読み取って概要を取得するサブエージェントを生成します
- その他 私は compilr-dev エコシステムをベータ版で立ち上げたばかりで、実際のフィードバックを得るために何人かのベータ テスターを迎え始めたいと考えています。ありがとう。カーメロ

記事本文:
compilr.dev — AI ワークスペース
タグ (ターミナルカード.tsx を参照) -->
メインコンテンツにスキップ

## Original Extract

Your AI workspace — specialized AI agents, structured project workflows, and persistent context for software, research, business, and beyond. Desktop app, CLI, and developer libraries.

Hi, over the past 6+ months I have built an ecosystem (libraries, CLI, Desktop) of tools to build or use AI agents. Depending on how close or not close you are with coding, you may lean towards the libraries (compilr-dev/agents and compilr-dev/agents-coding) to build your own AI Agents; you may use the CLI to manage your project and run your agents in the terminal when working, for example in VS Code or you may use the Desktop app (it shares the same core of the CLI but has for sure a richer UI). What, I believe, sets the compilr-dev ecosystem apart from other similar tools is the real multi-agent nature. Focusing on the Desktop app, it ships with pre-built specialized agents (built using the above mentioned libraries) each with tailored system prompt, skills and access to dedicated tools. When a new project is created a team of specialized agents is added to the team roster and you can add or remove agents already at project creation time (you can even create custom agents at project creation time). You can choose your preferred LLM provider (BYOK approach) and you can seamlessly switch provider whenever you want. Each agent has its own context but agents within a project also share some context. They can interact with each other by inquiring (agent A can ask something to agent B) or by task handoff. Each agent is instructed to follow a spec driven approach and the app itself encourage the user to design before building.
For example, as soon as you create a new project you can use the /design command to have the agent ask you questions to capture the project scope and details. As a result the agent will track work items in the database and will write the PRD document following your direction. All files are visible and editable by you and by the agents. Last but not least, special attention was given to the token consumptions: tokens = money so saving them is "paramount" (no, AI did not write this sentence, I deliberately chose to use the word "paramount").
Some techniques:
- Tools and skills definitions are loaded only when needed (like Neo loading the kung fu skills on Matrix)
- Tools output are pruned/summarized in the context after a few rounds
- Agents spawn subagents to read large files and get a summary
- and more I just launched the compilr-dev ecosystem in beta and I hope to start having some beta testers to get some real world feedback. Thank you. Carmelo

compilr.dev — Your AI Workspace
tag (see TerminalCard.tsx) -->
Skip to main content
