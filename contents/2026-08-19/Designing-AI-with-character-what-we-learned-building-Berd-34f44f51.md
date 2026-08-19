---
source: "https://block.xyz/inside/designing-ai-with-character-what-we-learned-building-berd"
hn_url: "https://news.ycombinator.com/item?id=49357223"
title: "Designing AI with character: what we learned building Berd"
article_title: "Block - Designing AI with character: what we learned building Berd"
image: "https://cdn.sanity.io/images/8m3angk4/production/6105f961ef9f2ce68c82a6544e5c4a4b9dd6e577-2803x1496.png?auto=format"
author: "thm"
captured_at: "2026-08-19T06:26:14Z"
capture_tool: "hn-digest"
hn_id: 49357223
score: 1
comments: 0
posted_at: "2026-08-19T05:26:26Z"
tags:
  - hacker-news
  - translated
---

# Designing AI with character: what we learned building Berd

- HN: [49357223](https://news.ycombinator.com/item?id=49357223)
- Source: [block.xyz](https://block.xyz/inside/designing-ai-with-character-what-we-learned-building-berd)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T05:26:26Z

## Translation

タイトル: キャラクターを使用した AI の設計: Berd の構築で学んだこと
記事のタイトル: Block - キャラクターを使用した AI の設計: Berd の構築で学んだこと
説明: Block では、必要なツールを自分たちで構築することがよくあります。 Berd は、チームがプロジェクト、スキル、ツール、モデル全体で AI エージェントと連携するために使用するデスクトップ アプリケーションです。現在、私たちは Berd をオープンソースにし、アプリケーションとその背後にある設計と技術的なレッスンを他の人が検査できるように共有しています。
[切り捨てられた]

記事本文:
Block - キャラクターを備えた AI の設計: Berd を構築して学んだこと
ブロックロゴニュース
2026 年 8 月 18 日 キャラクターを備えた AI の設計: Berd の構築で学んだこと
Block では、必要なツールを自分たちで構築することがよくあります。 Berd は、チームがプロジェクト、スキル、ツール、モデル全体で AI エージェントと連携するために使用するデスクトップ アプリケーションです。現在、私たちは Berd をオープンソースにし、アプリケーションとその背後にある設計と技術的なレッスンを共有し、他の人が検査、使用、適応できるようにしています。
Berd を構築することで、設計上の質問を検討する機会も得られました。AI エージェントのような抽象的なものを、どのようにして理解しやすく形にすることができるでしょうか?私たちの答えは、役割、指示、スキル、ツールだけでなく、独特のビジュアルアイデンティティを通じてエージェントに個性を与えることでした。
私たちのアプローチは、伝統的に知られていないカテゴリーにインスピレーションを受けたデザインをもたらしてきたブロックの長い歴史に基づいています。 Square は、販売者がカウンターの後ろに隠していた決済ハードウェアを、誇らしげに展示できるものに変えました。 Cash App は個人の財務に個性と文化的な関連性をもたらしました。いずれの場合も、デザインは人々の関わり方を変え、デザインを使うよう誘われていると感じる人を変えました。 Berd では、同じ本能を AI にもたらしました。目標はテクノロジーを隠すことではなく、その機能をより目に見え、親しみやすく、個人的なものにすることでした。
バードはブロック内部の実際的な問題から成長しました。私たちは、goose、Claude Code、Codex などのツールを通じて、ますます有能なエージェントにアクセスできるようになりましたが、それらと連携するには、さまざまなインターフェイス、構成システム、およびコンテキストの管理方法をナビゲートする必要がありました。テクノロジーは強力でしたが、それに関する経験は断片的であり、多くの人が持っていないレベルの技術的流暢さを前提としていることがよくありました。別のものは必要ないことを学びました

モデルやエージェントのハーネスを維持するには、それらの周囲に一貫した環境が必要でした。
Berd は、モデルやハーネス間で作業するための 1 つのデスクトップ アプリケーションをチームに提供しました。これにより、永続的なプロジェクトに関して会話、ファイル、フォルダー、指示、エージェント、スキルが統合されました。すべてのタスクのコンテキストを再構築する代わりに、人々は繰り返しの作業方法に基づいてエージェントを形成し、後でその作業に戻ることができます。
これは、エンジニアリングを超えてエージェントの作業をより親しみやすくするために設計されました。人々は会話か​​ら始めて、どのコンテキストとツールがアクティブになっているかを理解し、作業に必要な構造をさらに追加できます。
多くの AI インターフェイスは空のプロンプト ボックスから始まります。このモデルには機能があるかもしれませんが、この製品では、エージェントがどのように構成されているのか、どのコンテキストとツールが利用できるのか、他のエージェントとどのように異なるのかについてほとんど理解できません。私たちは、これらの違いをよりわかりやすく、誰でも構築できるようにしたいと考えました。
Berd では、キャラクターは視覚的かつ機能的です。異なるエージェントは別個であるため、別個に見えることがあります。これらは、思考を拡張したり、解決策を絞り込んだり、独自のスタイルで記述したり、Berd を使用して独自のツールを構築したりするのに役立ちます。 Berd とチャットするだけで、ユーザーは旅行やショッピングの計画などの日常業務を支援するなど、希望するあらゆるタスクを実行するカスタム エージェントを簡単に作成できます。
人々がエージェントを作成して共有できるようにするために、私たちは主力商品である「グルーピーズ」を含む、アニメーション キャラクターの特徴的なコレクションをデザインしました。これらのデザインは、抽象的な構成に、人々が理解し、楽しむことができる、認識可能な視覚的アイデンティティを与えます。
このアプローチにより、エージェントのワークフローに遊び心のあるひねりを加えるだけでなく、パーソナル エージェントの作成とカスタマイズが誰でも簡単に行えるようになります。アバターはエージェントに認識させます

ザーブル。その役割、スキル、ツールにより、それは役に立ちます。
Berd は、2025 年 1 月に導入されたオープンソース AI エージェント フレームワーク Block である goose の構築と使用を学んだことから成長しました。 goose は、オープンなモジュラー アーキテクチャを通じて、言語モデルをツールや現実世界のアクションに接続します。 2025 年 12 月、ブロック氏は Anthropic、OpenAI などと協力し、Linux Foundation の下に Agentic AI Foundation を設立しました。その後、私たちは Goose をこの財団に提供し、Model Context Protocol や AGENTS.md と並んでベンダー中立の拠点を与えました。
ベルドとグースは現在、異なるが関連した目的を果たしています。 goose はオープンなエージェント フレームワークとランタイムのままです。 Berd は、それを中心に構築されたデスクトップ アプリケーションです。
Berd は、エージェント クライアント プロトコルを通じて Goose に接続します。この分離により、Berd はプロジェクト、コンテキスト、セッション、エージェント、構成などのデスクトップ エクスペリエンスに集中できるようになり、Goose は基盤となるエージェント ループを処理します。
Berd の構築は、プライベート デスクトップ エクスペリエンスがどこで終わるのかを理解するのにも役立ちました。エージェントとの仕事は一人で始まることがよくあります。他の人が使えるようになる前に、リサーチ、実験、コンテキストの収集、そしてアイデアの形を作ります。ベルドは私たちにその個人的な経験を探求する場所を与えてくれました。
しかし、仕事がプライベートのままになることはほとんどありません。最終的には、チームメイトを招いたり、別のエージェントを追加したり、成果物を共有したり、決定事項を説明したり、何かがどのように行われたかの記録を残したりする必要があるかもしれません。これらは共同作業の問題であり、最近リリースされた Buzz の焦点です。
Buzz は、コミュニティが制御できる中継器を使用して、ユーザーとエージェントが共有ルームで共同作業するオープンソース ワークスペースです。ユーザーとエージェントは独自の ID を持ち、同じ会話に参加し、同じ検索可能なレコードに貢献します。 Berd から学んだことは、Buzz での継続的な作業に影響を与えます。 B

erd は、プライベート スペース、永続的なコンテキスト、認識可能なエージェント ID、再利用可能なスキル、目に見える設定、およびエージェントの設定されたコンテキスト、ツール、機能の明確な可視性の重要性を示しました。バズは、個々の作業が共同作業になるときに、それらのアイデアをどこかに行き先として提供します。最初は 1 人で始めて、その後はマルチプレイヤーに移行します。
いつものように、皆様からのご意見をお待ちしております。 Berd をダウンロードし、独自のエージェントやワークフローで試して、フィードバックを共有してください。私たちは、あなたが作ったものを見て、何があなたにとって効果的か、そしてエージェントのエクスペリエンスがどうあるべきかについて意見を聞きたいと思っています。
検索 AI チャット あなたは、生成 AI を利用した自動チャットボットに接続されています。続けることで、
あなたは、Block のプライバシーに関する通知がこのツールの使用に適用されることを認識し、こことこちらに記載されている Block の AI 利用規約に同意するものとします。

## Original Extract

At Block, we often build the tools we need ourselves. Berd is a desktop application our teams use to work with AI agents across projects, skills, tools, and models. Today, we’re making Berd open source, sharing the application, and the design and technical lessons behind it for others to inspect, us
[truncated]

Block - Designing AI with character: what we learned building Berd
Block logo News
August 18, 2026 Designing AI with character: what we learned building Berd
At Block, we often build the tools we need ourselves. Berd is a desktop application our teams use to work with AI agents across projects, skills, tools, and models. Today, we’re making Berd open source, sharing the application, and the design and technical lessons behind it for others to inspect, use, and adapt.
Building Berd also gave us a chance to explore a design question: How do you make something as abstract as an AI agent easier to understand and shape? Our answer was to give agents character, not only through roles, instructions, skills, and tools, but through distinctive visual identities.
Our approach builds on Block’s long history of bringing inspired design to categories not traditionally known for it. Square transformed payment hardware from something sellers hid behind the counter into something they could display proudly. Cash App brought personality and cultural relevance to personal finance. In each case, design changed how people related to it, and who felt invited to use it. With Berd, we brought the same instinct to AI. The goal was not to disguise the technology, but to make its capabilities more visible, approachable, and personal.
Berd grew from a practical problem inside Block. We had access to increasingly capable agents through tools such as goose, Claude Code, and Codex, but working with them meant navigating different interfaces, configuration systems, and ways of managing context. The technology was powerful, but the experience around it was fragmented, and often assumed a level of technical fluency that many people did not have. We learned that we didn’t need another model or agent harness, we needed a consistent environment around them.
Berd gave our teams one desktop application for working across models and harnesses. It brought conversations, files, folders, instructions, agents, and skills together around persistent projects. Instead of rebuilding context for every task, people could shape agents around recurring ways of working and return to that work later.
This was designed to make agentic work more approachable beyond engineering. People could begin with a conversation, understand which context and tools were active, and add more structure as the work required it.
Many AI interfaces begin with an empty prompt box. The model may be capable, but the product gives people little sense of how the agent is configured, which context and tools are available to it, and how it differs from another agent. We wanted to make those distinctions more visible and approachable for anyone to build.
In Berd, character is both visual and functional. Different agents can look distinct because they are distinct. They can help you expand your thinking or narrow down on a solution, write in your unique style, or build your own tools using Berd. Just by chatting with Berd, users can easily create custom agents to do any task they want, even helping with everyday tasks like planning travel or shopping.
To help people build and share their agents, we’ve designed distinctive collections of animated characters, including our flagship “Gloopies.” These designs give abstract configurations a recognizable visual identity that people can understand and have fun with.
More than just giving a playful twist to the agent workflow, this approach makes creating and customizing personal agents more accessible for anyone. The avatars make the agent recognizable. Its role, skills, and tools make it useful.
Berd grew from what we learned building and using goose , the open source AI agent framework Block introduced in January 2025. goose connects language models to tools and real-world actions through an open, modular architecture. In December 2025, Block joined Anthropic, OpenAI, and others to establish the Agentic AI Foundation under the Linux Foundation. We later contributed goose to the foundation, giving it a vendor-neutral home alongside the Model Context Protocol and AGENTS.md .
Berd and goose now serve different but connected purposes. goose remains the open agent framework and runtime. Berd is a desktop application built around it.
Berd connects to goose through the Agent Client Protocol. This separation lets Berd focus on the desktop experience, including projects, context, sessions, agents, and configuration, while goose handles the underlying agent loop.
Building Berd also helped us understand where a private desktop experience stops. Work with an agent often begins alone. You research, experiment, gather context, and shape an idea before it is ready for other people. Berd gave us a place to explore that individual experience.
But work rarely stays private. Eventually, you may need to bring in a teammate, add another agent, share an artifact, explain a decision, or keep a record of how something was made. Those are collaborative problems, and they are the focus of our recently released Buzz .
Buzz is our open source workspace where people and agents work together in shared rooms, on a relay the community can control. People and agents have their own identities, participate in the same conversations, and contribute to the same searchable record. What we learned from Berd will inform our continued work on Buzz. Berd showed us the importance of private space, durable context, recognizable agent identities, reusable skills, visible configuration, and clearer visibility into an agent’s configured context, tools, and capabilities. Buzz gives those ideas somewhere to go when individual work becomes collaborative. Start alone, then go multiplayer.
As always, we want to hear from you. Download Berd , try it with your own agents and workflows, and share your feedback. We want to see what you make, and hear what works for you, and what you think agent experiences should become.
Search AI Chat You are being connected to our automated chatbot which utilizes generative AI. By continuing,
you recognize that Block’s Privacy Notice applies to your use of this tool and you agree to Block’s AI terms of use located here and here .
