---
source: "https://www.thecontentwrangler.com/p/tech-writers-may-be-solving-the-wrong"
hn_url: "https://news.ycombinator.com/item?id=48506454"
title: "Tech Writers May Be Solving the Wrong AI Problem"
article_title: "AI Has a Meaning Problem, Not a Content Problem"
author: "eigenBasis"
captured_at: "2026-06-12T18:46:46Z"
capture_tool: "hn-digest"
hn_id: 48506454
score: 2
comments: 0
posted_at: "2026-06-12T16:51:10Z"
tags:
  - hacker-news
  - translated
---

# Tech Writers May Be Solving the Wrong AI Problem

- HN: [48506454](https://news.ycombinator.com/item?id=48506454)
- Source: [www.thecontentwrangler.com](https://www.thecontentwrangler.com/p/tech-writers-may-be-solving-the-wrong)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T16:51:10Z

## Translation

タイトル: テックライターは間違った AI 問題を解決している可能性がある
記事のタイトル: AI には内容の問題ではなく、意味の問題がある
説明: テクニカル ライターがエンタープライズ AI の欠落しているレイヤーに座っている理由

記事本文:
AI には内容の問題ではなく、意味の問題がある
コンテンツラングラー
AI と技術ドキュメント テックライターは間違った AI 問題を解決している可能性がある
情報を整理することが意味を定義することと同じではない理由
Scott Abel 2026 年 6 月 12 日 69 シェア 多くの組織が、ドキュメントを消費する AI システムに多額の投資を準備しています。同時に、そのドキュメントの多くは、ギャップを埋めたり、仮定を立てたり、明示的に述べられていない文脈を解釈したりするのが得意な人間の読者向けに書かれています。
人間はこれを非常に自然に行っているため、ほとんど気づきません。機械はそうではありません。
ナレッジ グラフ研究者の Benny Cheung による最近の LinkedIn の投稿を読みながら、私はそのことについて考えていました。同氏は、データから分類法、オントロジー、ナレッジグラフへの進化を説明し、その進行が早すぎて止まってしまった組織は AI に苦戦する可能性が高いと主張しました。
テクニカル ライターは何年もこの問題の一部を扱ってきたため、この投稿は私の心に響きました。ただし、私たちは普段同じ語彙を使用しません。
Content Wrangler は読者によってサポートされる出版物です。新しい投稿を受け取り、私の仕事をサポートするには、無料または有料の購読者になることを検討してください。
情報を整理することと意味を定義することは同じではない
ほとんどの組織には、情報を保存する場所がたくさんあります。データベース、コンテンツ管理システム、ナレッジ ベース、SharePoint サイト、およびファイル リポジトリを備えています。
Cheung 氏は、データベース スキーマとオントロジーは異なる問題を解決すると指摘します。スキーマは、情報がどのように編成されるかを記述します。オントロジーは、何かが何であるか、それが他のものとどのように関係するか、そしてそれらの関係をどのようなルールが支配するかを記述します。
多くの組織はこれらの区別を曖昧にしています。彼らは、情報が構造化されると、カテゴリー化されると想定しています。

整理され、保存されているため、何らかの形で乗り出すことができます。
「ファームウェアのアップデート」などのコンテンツ ラベルと、誰がアップデートを実行するか、どのイベントがアップデートをトリガーしたか、アップデートを開始する前に満たす必要がある条件は何か、アップデートが成功または失敗したときに何が起こるかを説明する説明との違いを考慮してください。
👉🏾 トピックを特定するのに役立ちます。
👉🏾 もう 1 つは、何かがどのように機能するかを説明するのに役立ちます。
AI システムの場合、その詳細が重要です。
分類法は物事を見つけるのに役立ちます
技術文書チームは分類法を構築するのに何年も費やしてきました。製品、機能、対象読者、ドキュメントの種類、メタデータ値を分類します。その仕事は、人々が情報を見つけるのに役立つので役に立ちます。
課題は、分類では質問の一部しか答えられないことです。
AI システムはテキストだけでなく構造化されたコンテキストにますます依存します。ナレッジ グラフと RDF (リソース記述フレームワーク) はエンティティとエンティティ間の関係を表し、アクション モデルとプロセス モデルは、誰がアクションを実行するか、どのオブジェクトが影響を受けるか、どのようなイベントや条件がワークフローを形成するか、アクティビティがプロセス全体でどのように接続されるかを明示します。この構造は、AI システムがコンテキストを含む情報を取得し、複数ステップのタスクをより確実に実行するのに役立ちます。
これらの関係は、通常、人間の読者が推測できるため、複数の文書に散在しているか、記載されていないことがよくあります。 AI システムは常に正確または一貫して推論するとは限りません。
ここから、多くのドキュメント リポジトリが古さを見せ始めます。情報は存在しますが、関係が不完全、矛盾、または暗黙的です。
なぜドキュメントが AI を混乱させるのか
人々が私たちのドキュメントを読むとき、彼らは生涯にわたる経験をプロセスにもたらします。彼らは、顧客と技術者には異なる責任があることを知っています。彼らは認識しています

セットアップ手順は日常的な操作には適用できない場合があります。通常、警告がすべての状況ではなく、特定の状況に属するかどうかを知ることができます。
ドキュメントはそのような判断に依存することがよくあります。
大規模な言語モデルは人間のようには読み取れません。遠隔地でもありません。パターンと確率に基づいて応答を生成します。重要な関係が欠落している場合、モデルは接続すべきではないものを接続したり、重要な条件をスキップしたり、間違った人やシステムにアクションを割り当てたりする可能性があります。
結果は合理的であるように見えても、間違っている可能性があります。
それが私が TRACE フレームワークに取り組んできた理由の 1 つです。
このフレームワークは単純な観察から生まれました。 AI に影響を与える文書化の問題の多くは、情報の欠落よりもコンテキストの欠落と関係しています。事実は存在するかもしれませんが、それらの事実間の関係は存在しないことがよくあります。
テクニカル ライターはすでにこの分野で働いています
技術ライターは日常的にコンポーネントを定義し、ワークフローを説明し、責任を文書化し、依存関係を説明し、結果に影響を与える条件を特定します。
私たちのほとんどはそれをドキュメントだと考えています。しかし、別の見方をすると、それはシステムがどのように動作するかを定義する作業でもあります。
だからこそ、一部の組織は貴重な専門知識の源を見逃しているのではないかと思います。彼らは、エンタープライズナレッジの大規模なコレクションにわたる質問に回答できる AI システムを望んでいますが、ドキュメントを知識の定義機能ではなく公開機能として扱うことがよくあります。
関係を理解し​​ている人は、ドキュメントを書いた人と同じ人であることがよくあります。
ナレッジグラフが再び話題になっている理由
ナレッジ グラフは何年も前から存在していましたが、多くの組織では、ナレッジ グラフに投資する説得力のある理由が見つかりませんでした。検索ゲ

概ね「十分」に機能しました。
生成 AI は期待を変えました。組織は現在、関係を説明し、状況に応じて質問に答え、部門間の情報を結び付け、さまざまな状況に対応できるシステムを求めています。
これらの機能はコンテンツの取得だけに依存します。これらは、情報が互いにどのように関連しているかを理解することに依存しています。
ナレッジ グラフはこれに役立ちますが、それは基礎となるコンテンツに明確な関係が含まれている場合に限られます。ソース資料があいまいな場合、グラフにはそのあいまいさが引き継がれます。
ある時点で、AI が空白を埋め始めます。 AI の問題と思われる問題は、ドキュメントにずっと存在していた弱点を露呈させている可能性があります。
AI への対応についての別の考え方
AI への対応に関する会話の多くは、モデル、ツール、インフラストラクチャに焦点を当てています。そういったことが重要なのです。しかし、組織は情報そのものにも目を向ける必要があるかもしれません。
👉🏾 運用コンテキストは明示的ですか?
👉🏾 俳優は明確に特定されていますか?
👉🏾 条件は文書化されていますか?
👉🏾 ワークフローは、マシンが仮定せずに従うことができる方法で接続されていますか?
これらの質問は、ドキュメントに関する質問によく似ています。
だからこそ、Cheung 氏の投稿が私の注目を集めたのです。これは、テクニカル ライターが長い間取り組んできたギャップを浮き彫りにします。組織が AI を顧客サポート、トレーニング、サービス提供、エンタープライズ検索にさらに深く組み込むにつれて、そのギャップを無視するのは難しくなります。
問題は、組織にコンテンツが不足していることではないかもしれません。おそらく、それを読んだ人の頭の中にあまりにも多くの意味が残っているのかもしれません。 🤠
69 シェア この投稿に関する前のディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

Why Technical Writers May Be Sitting on the Missing Layer of Enterprise AI

AI Has a Meaning Problem, Not a Content Problem
The Content Wrangler
Subscribe Sign in AI and Tech Docs Tech Writers May Be Solving the Wrong AI Problem
Why organizing information isn't the same thing as defining meaning
Scott Abel Jun 12, 2026 69 Share Many organizations are preparing to invest heavily in AI systems that consume documentation. At the same time, much of that documentation was written for human readers who are good at filling in gaps, making assumptions, and interpreting context that was never explicitly stated.
Humans do this so naturally that we rarely notice it. Machines don’t.
I was thinking about that while reading a recent LinkedIn post from knowledge graph researcher Benny Cheung . He described the progression from data to taxonomy to ontology to knowledge graphs and argued that organizations that stop too early in that progression are likely to struggle with AI.
The post resonated with me because technical writers have been dealing with parts of this problem for years, although we don’t usually use the same vocabulary.
The Content Wrangler is a reader-supported publication. To receive new posts and support my work, consider becoming a free or paid subscriber.
Organizing Information Isn't The Same Thing As Defining Meaning
Most organizations have plenty of places to store information. They have databases, content management systems, knowledge bases, SharePoint sites, and file repositories.
Cheung points out that a database schema and an ontology solve different problems. A schema describes how information is organized. An ontology describes what something is, how it relates to other things, and what rules govern those relationships.
Many organizations blur those distinctions. They assume that once information is structured, categorized, and stored, meaning somehow comes along for the ride.
Consider the difference between a content label such as “firmware update” and a description that explains who performs the update, what event triggered it, what conditions must be true before it can begin, and what happens when it succeeds or fails.
👉🏾 One helps identify a topic.
👉🏾 The other helps explain how something works.
For AI systems, those details matter.
Taxonomy Helps You Find Things
Technical documentation teams have spent years building taxonomies . We classify products, features, audiences, document types, and metadata values. That work is useful because it helps people find information.
The challenge is that classification only answers part of the question.
AI systems increasingly depend on structured context, not just text. Knowledge graphs and RDF (Resource Description Framework) represent entities and the relationships between them, while action and process models make it explicit who performs an action, what object is affected, what events or conditions shape a workflow, and how activities connect across a process. That structure helps AI systems retrieve information with context and perform multi-step tasks more reliably.
Those relationships are often scattered across multiple documents or left unstated because human readers can usually infer them. AI systems don’t always infer them correctly or consistently.
That’s where many documentation repositories start to show their age. The information exists, but the relationships are incomplete, inconsistent, or implied.
Why Documentation Can Confuse AI
When people read our documentation, they bring a lifetime of experience to the process. They know that customers and technicians have different responsibilities. They recognize that a setup procedure may not apply during routine operations. They can usually tell when a warning belongs to a specific circumstance rather than every circumstance.
Documentation often depends on that kind of judgment.
Large language models don’t read the way people do. Not even remotely. They generate responses based on patterns and probabilities. When key relationships are missing, the model may connect things that shouldn’t be connected, skip conditions that matter, or assign actions to the wrong person or system.
The result can sound reasonable while being wrong.
That’s one of the reasons I’ve been working on the TRACE framework:
The framework grew out of a simple observation. Many documentation problems that affect AI have less to do with missing information than with missing context. The facts may be present, but the relationships between those facts often aren’t.
Technical Writers Are Already Working in This Space
Tech writers routinely define components, describe workflows, document responsibilities, explain dependencies, and identify conditions that affect outcomes.
Most of us think of that as documentation. But, viewed another way, it’s also the work of defining how a system operates.
That’s why I think some organizations are overlooking a valuable source of expertise. They want AI systems that can answer questions across large collections of enterprise knowledge, but they often treat documentation as a publishing function rather than a knowledge-definition function.
The people who understand the relationships are frequently the same people who wrote the documentation.
Why Knowledge Graphs Are Back In The Conversation
Knowledge graphs have been around for years, but many organizations never saw a compelling reason to invest in them. Search generally worked “good enough.”
Generative AI changed expectations. Organizations now want systems that can explain relationships, answer questions in context, connect information across departments, and adapt responses to different situations.
Those capabilities depend on more than retrieving content. They depend on understanding how pieces of information relate to one another.
Knowledge graphs can help with that, but only if the underlying content contains clear relationships. If the source material is ambiguous, the graph inherits the ambiguity.
At some point the AI starts filling in the blanks. What appears to be an AI problem may be exposing weaknesses that have existed in our documentation all along.
A Different Way to Think About AI Readiness
Many conversations about AI readiness focus on models, tools, and infrastructure. Those things matter. But organizations may also need to look at the information itself.
👉🏾 Is operational context explicit?
👉🏾 Are actors clearly identified?
👉🏾 Are conditions documented?
👉🏾 Are workflows connected in ways that a machine can follow without making assumptions?
Those questions sound a lot like documentation questions.
That’s why Cheung’s post caught my attention. It highlights a gap that technical writers have been working around for a long time. As organizations push AI deeper into customer support, training, service delivery, and enterprise search, that gap becomes harder to ignore.
The issue may not be that organizations lack content. It may be that too much of the meaning still lives in the heads of the people reading it. 🤠
69 Share Previous Discussion about this post Comments Restacks Top Latest Discussions No posts
