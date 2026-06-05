---
source: "https://newsletter.eng-leadership.com/p/how-openai-codex-tech-lead-does-ai"
hn_url: "https://news.ycombinator.com/item?id=48417941"
title: "OpenAI Codex Tech Lead Does AI-Assisted Engineering"
article_title: "How OpenAI Codex Tech Lead Does AI-Assisted Engineering"
author: "gregorojstersek"
captured_at: "2026-06-05T21:34:49Z"
capture_tool: "hn-digest"
hn_id: 48417941
score: 1
comments: 0
posted_at: "2026-06-05T20:40:19Z"
tags:
  - hacker-news
  - translated
---

# OpenAI Codex Tech Lead Does AI-Assisted Engineering

- HN: [48417941](https://news.ycombinator.com/item?id=48417941)
- Source: [newsletter.eng-leadership.com](https://newsletter.eng-leadership.com/p/how-openai-codex-tech-lead-does-ai)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T20:40:19Z

## Translation

タイトル: OpenAI Codex 技術リーダーが AI 支援エンジニアリングを実施
記事のタイトル: OpenAI Codex の技術リーダーは AI 支援エンジニアリングをどのように行うか
説明: Codex オープンソース リポジトリの技術リーダーである Michael Bolin が、AI 支援エンジニアリングを使用して権限システムを構築した方法を共有しています。

記事本文:
OpenAI Codex 技術リーダーが AI 支援エンジニアリングをどのように行うか
エンジニアリングのリーダーシップ
OpenAI Codex の技術リーダーが AI 支援エンジニアリングをどのように行うか
Codex オープンソース リポジトリの技術リーダーである Michael Bolin 氏は、AI 支援エンジニアリングを使用して権限システムを構築した方法を共有しています。
Gregor Ojstersek Jun 04, 2026 ∙ 有料 29 5 シェア イントロ
私は最近、サンフランシスコにある OpenAI のオフィスを訪問する機会に恵まれました。そこで話をした人の 1 人が、Meta の元優秀エンジニアで、現在は Codex オープン ソース リポジトリ (Codex CLI) の OpenAI 技術リーダーである Michael Bolin でした。
マイケルはオライリーの著者仲間でもあり、2010 年に『 Closure: The Definitive Guide 』というタイトルの本を執筆しました。さらに、彼は Meta でオープンソース ビルド ツールである Buck を作成し、Google 在籍中に Google カレンダーの開発に取り組みました。
私たちは、彼が AI 支援エンジニアリングをどのように行っているのか、そして彼のワークフローがどのようなものであるのかについて話し合いました。私たちは彼が最近取り組んだ特定の機能を調べましたが、私が最も驚いたのは、彼のワークフローが非常にシンプルでわかりやすいことです。
基本的に、彼のワークフローは次のとおりです: 仕様を書く → 簡単なプロンプト → コードをレビューする。
クレイジーな AI ワークフローはありません。明晰な思考、適切な判断、そして迅速な反復だけです。この記事では、彼がどのようにそれを行うかについての詳細をすべて共有します。
これは有料購読者向けの記事です。完全なインデックスは次のとおりです。
- Codex CLI 内で権限システムを構築する
- 執筆要件
- 初期実装
- マージ競合の処理
- PR コードのサイズを管理しやすい状態に保ち、過去の PR とのコンテキストを提供することの重要性
- Codex による PR の作成と、Codex と協力して PR を作成するエンジニアの比較
🔒 エージェント PR は、一度に 1 つまたは 2 つ以上のタスクを実行するのに役立ちます
🔒 統合テストはより強力な署名を与える

アプリが正しく動作しているかどうかの単体テストよりも nals
🔒 人間によるコード レビューと AI コード レビュー
🔒 Codex を使用して問題を解決する方法の例
🔒 AI支援エンジニアリングを行う際には、優れた基礎が重要です
🔒 最後の言葉
Codex CLI 内での権限システムの構築
私たちは、彼らが最近完了したプロジェクトである Codex CLI 内の権限システムから議論を始めました。同氏は、Codex 内でこれを構築する動機は、特にエンタープライズ企業向けであると述べました。企業は機能を好みますが、特定の権限に基づいてどの機能を使用できるかについての制約も必要とします。
企業は、AI の動作を封じ込めて安全に保つことにも気を配っているため、AI は許可されているものにのみアクセスして変更することができます。そのため、これが正しく処理されるように、サンドボックス関連の機能にも取り組みました。
これは、Codex チームが数か月かけて取り組んだより大きなプロジェクトでした。
では、プロジェクトがどのように始まったのかを見ていきましょう。
要件は非常に簡単に書かれています。 Michael は Notion ドキュメントを作成し、サポートする必要がある要件と機能を説明しました。基本的には技術仕様です。
次に、彼はチームに仕様とその特定の部分を確認してコメントするように依頼しました。チームはそれを検討し、何が欠けている可能性があるのか​​、あるいは何が別の方法でできるのかについて全員が意見を共有し、その結果、初期設計が改良されました。
このプロセスが完了すると、すべてのコメントが処理され、チームはそれに満足し、プロジェクトの実装フェーズを開始しました。
この機能を実装する段階では、すでに Codex で Notion コネクタがサポートされていました。したがって、コンテキスト内の要件をコピーして貼り付ける代わりに、Michael は Notion の URL を Cod に貼り付けるだけで済みました。

元、「今度はこれを作りたいと思います」と言いました。
Notion コネクタのおかげで、Codex はソースに直接アクセスして、要件、コメント、ディスカッション履歴を読むことができました。 Michael は、すべてのコンテキストを手動で再作成するよりもはるかに優れているため、この部分が本当に気に入っていると明言しました。
彼が Codex に最初に依頼したのは、計画を作成し、作業を適切なサイズのプル リクエストのセットに分割することでした。
Michael 氏はまた、2,000 行の PR は人間がレビューする必要があり、非常に苦痛になるため、望まないとも述べました。
Codex によって最初に生成されたコードは約 6 PR でした。彼らは、コードのレビュー、テスト、マージ競合の処理、および途中での改良を行いながら、それらに取り組み始めました。
これが、彼らのワークフローが基本的にプロジェクトを構築する方法でした。
この時点で、私は AI 生成コードのマージ競合の処理について Michael に尋ねました。これは、特にコードベースの同じ部分に触れる多数の PR を開いている場合、非常に苦痛になる可能性があるためです。
彼は、Codex はマージ競合に対処する際、実際には彼よりもはるかに規律を持っていると述べました。彼はマージ競合に対処するための独自のスキルを作成しました。これは基本的に特定の命令セットです。
このスキルがあれば、変更が並行して行われ、ファイルが他のエンジニアによって変更されるときに、Codex にその作業の多くを一貫して処理させることができます。
PR コードのサイズを管理しやすい状態に保ち、過去の PR のコンテキストを提供することの重要性
彼は、適切なサイズのプルリクエストを作成することを強く支持していると述べました。彼は、人間がコードをレビューする必要があるため、変更はレビュー可能なサイズに分割する必要があることを Codex によく思い出させます。
時には、仕事をどのように分割したいかを明確に伝えることもあります。また、彼は次のように尋ねることもあります。
「これを分割する意味がどこにあるのか？」

変更はより小さな PR に変わりますか?」または「この変更をレビュー可能な部分に分割するにはどうすればよいですか?」
また、特定の検証戦略を念頭に置いている場合は、テストの期待値も提供します。
彼が大きく依存しているもう 1 つの実践は、関連するプル リクエストを参照することです。実際、それは多くの場合、最も有用なコンテキスト情報源の 1 つであると彼は述べています。
関連する PR には、コード、ディスカッション、コメント、決定事項が含まれます。 Codex にそのコンテキストへのアクセスを許可すると、Codex の理解と出力が大幅に向上します。
さらに、他の人が開いた PR を「PR が大きすぎる」というメッセージで拒否することがあることや、作業を分割するのに手間がかかるため、歴史的に議論が起こった可能性があることにも言及しました。
現在では、AI コーディング ツールを使用することではるかに簡単になりました。 Codex に依頼するだけで済みます。これにより、適切なレビューの実践がはるかに簡単になります。最も重要なことは、レビューを適切なサイズまで縮小して、各プル リクエストが実際に何を行っているかを明確にすることです。
Codex が PR を作成する場合と、Codex と協力して PR を作成するエンジニアがいる場合
あるいは、エージェント型 PR と AI 支援型 PR と呼ぶこともできます。最近のほとんどのエンジニアは、Claude Code、Cursor、または Codex を使用する 2 番目の方法を好むことがわかっていますが、Michael は最初の方法 (Codex による PR の作成) を好みます。その理由を彼に尋ねました。
彼は、CI の実行には約 15 ～ 20 分かかるため、Codex がすぐにプル リクエストを作成して、CI ができるだけ早く実行を開始できるようにしてほしいと述べました。
彼らは、「ベビーシッター PR」または「ベビーシッター CI」と呼ばれるスキルも持っています。プル リクエストが存在すると、Codex は障害を監視し、自動的に対応できます。
また、CI は Mac、Linux、Windo でテストを実行するため、ほとんどのテストをローカルで実行しないように Codex に指示することもよくあります。

とにかく。彼はパイプラインをすぐに開始したいと考えています。
これは、マシン上のリソースを消費する完全なテスト スイートを複数用意する必要がないことも意味します。想像してみてください。5 ～ 6 個のローカル変更を同時に実行し、すべてのテストの実行が開始されると、ローカル マシンで消費されるコンピューティング量によって、生産性を高めることが非常に困難になります。
エージェント PR は、1 つまたは 2 つ以上のタスクを同時に実行するのに役立ちます
この投稿は有料購読者向けです

## Original Extract

Michael Bolin, Tech Lead for the Codex open-source repo, is sharing how they built the permissions system with AI-assisted engineering.

How OpenAI Codex Tech Lead Does AI-Assisted Engineering
Engineering Leadership
Subscribe Sign in How OpenAI Codex Tech Lead Does AI-Assisted Engineering
Michael Bolin, Tech Lead for the Codex open-source repo, is sharing how they built the permissions system with AI-assisted engineering.
Gregor Ojstersek Jun 04, 2026 ∙ Paid 29 5 Share Intro
I recently had the pleasure of visiting OpenAI’s offices in San Francisco, and one of the people I talked to there was Michael Bolin , a former Distinguished Engineer at Meta, now OpenAI Tech Lead for the Codex open source repo (Codex CLI) .
Michael is also a fellow O’Reilly author, he wrote a book titled Closure: The Definitive Guide , dating back to 2010. Additionally, he created Buck , an open-source build tool at Meta, and he worked on Google Calendar during his time at Google.
We talked all about how he does AI-assisted engineering and what his workflow looks like. We went through a specific feature that he recently worked on, and what surprised me the most is that his workflow is very simple and straightforward.
Essentially, his workflow is the following: Write the spec → simple prompt → review the code.
No crazy AI workflows. Just clear thinking, good judgment, and fast iteration. I am sharing all the details on how he does it in this article!
This is an article for paid subscribers, and here is the full index:
- Building a permissions system within Codex CLI
- Writing requirements
- Initial implementation
- Handling merge conflicts
- Importance of keeping PR code sizes manageable and providing context with past PRs
- Codex creating PRs versus an engineer creating PRs in collaboration with Codex
🔒 Agentic PRs help doing more than 1 or 2 tasks at the time
🔒 Integration tests give a lot stronger signals than unit tests if the app is working correctly
🔒 Human code review versus AI code review
🔒 Example of how to use Codex to resolve an issue
🔒 Good fundamentals are crucial when doing AI-assisted engineering
🔒 Last words
Building a permissions system within Codex CLI
We started our discussion with the project they recently finished, the permissions system within Codex CLI. He mentioned that the motivation behind building it within Codex is especially for Enterprise companies. Enterprises love features, but they also need constraints on which features can be used based on certain permissions.
Enterprises also care about keeping AI’s actions contained and secure, so it can only access and modify what it’s allowed to. So, ensuring that this is handled correctly, they also worked on features related to sandboxing.
This was a bigger project that the Codex team has worked on over several months.
Now, let’s go into how the project has started.
The requirements have been written in a very simple way. Michael created a Notion document, where he described the requirements and functionalities that need to be supported. Basically a tech spec.
Then he asked the team to check and comment on the spec and specific parts of it. The team took a look, and everyone shared their perspective on what’s potentially missing or what could be done differently, which resulted in refining the initial design.
When that process was done, all the comments were addressed, and the team was happy with it, then they started with the implementation phase of the project.
When it was time to implement the feature, they already had support for a Notion connector in Codex. So, instead of needing to copy and paste requirements in the context, Michael simply pasted the Notion URL into Codex and said, “Now I want to build this”.
Because of the Notion connector, Codex was able to go directly to the source, read the requirements, comments, and discussion history. Michael explicitly mentioned that he really liked this part as it was much nicer than manually recreating all that context.
The first thing he asked Codex to do was create a plan and break the work into a set of right-sized pull requests.
Michael also mentioned that he doesn’t want 2,000-line PRs because a human still has to review them, which would make the experience very painful.
The initially generated code by Codex was around 6 PRs. They started working through them, reviewing the code, testing, handling merge conflicts, and doing refinements along the way.
That’s how their workflow essentially looked for building the project.
At this point, I asked Michael about handling merge conflicts for AI-generated code, as that can be quite painful, especially if you have many PRs open touching the same parts of the codebase.
He mentioned that Codex is actually a lot more disciplined when dealing with merge conflicts than he is. He created his own skill for dealing with merge conflicts, which is basically a specific set of instructions.
With the skill, as changes happen in parallel and files get modified by other engineers, he can simply let Codex handle a lot of that work consistently.
Importance of keeping PR code sizes manageable and providing context with past PRs
He mentioned that he’s a big advocate of creating right-sized pull requests. He often reminds Codex that a human needs to review the code, so changes should be broken into a size that is still reviewable.
Sometimes he explicitly tells it how he wants the work divided. Other times, he asks:
“Where does it make sense to split this change into smaller PRs?” or “How would you break this change into reviewable pieces?”
He also provides testing expectations when he has a specific validation strategy in mind.
Another practice he relies on heavily is referencing related pull requests. In fact, that’s often one of the most useful sources of context, he mentions.
Related PRs contain code, discussions, comments, and decisions. Giving Codex access to that context can significantly improve its understanding and output.
Additionally, he also mentioned that he occasionally rejects PRs opened by other people with the message of “PR being too large”, and that historically might have caused arguments because splitting work required effort.
Today, it’s much easier with AI coding tools. You can simply ask Codex to do it. That makes it much easier to enforce good review practices. The biggest thing is getting reviews down to the right size so it’s clear what each pull request is actually doing.
Codex creating PRs versus engineer creating PRs in collaboration with Codex
Or we could also call it Agentic PRs versus AI-assisted PRs. I know preference for most engineers these days is the second one, using Claude Code, Cursor, or Codex, while Michael has the preference for the first one (Codex creating PRs), which I asked him for the reasoning.
He mentioned that their CI takes around 15 to 20 minutes to run, and he wants Codex to create the pull request immediately so CI can start running as soon as possible.
They also have skills around what they call “babysitting PRs” or “babysitting CI”. Once the pull request exists, Codex watches for failures and can react to them automatically.
He also often tell Codex not to run most tests locally because CI will run the tests across Mac, Linux, and Windows anyway. He’d rather get the pipeline started immediately.
It also means, that he doesn’t need to have multiple full test suites consuming resources on his machine. Imagine, if you run 5-6 local changes at the same time, and all the tests start to run, the amount of compute it consumes on a local machine, makes it really hard to be productive.
Agentic PRs help doing more than 1 or 2 tasks at the same time
This post is for paid subscribers
