---
source: "https://claude.com/blog/the-new-rules-of-context-engineering-for-claude-5-generation-models"
hn_url: "https://news.ycombinator.com/item?id=49040821"
title: "The new rules of context engineering for Claude 5 generation models"
article_title: "The new rules of context engineering for Claude 5 generation models | Claude by Anthropic"
author: "opwizardx"
captured_at: "2026-07-24T20:10:51Z"
capture_tool: "hn-digest"
hn_id: 49040821
score: 1
comments: 0
posted_at: "2026-07-24T19:52:38Z"
tags:
  - hacker-news
  - translated
---

# The new rules of context engineering for Claude 5 generation models

- HN: [49040821](https://news.ycombinator.com/item?id=49040821)
- Source: [claude.com](https://claude.com/blog/the-new-rules-of-context-engineering-for-claude-5-generation-models)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T19:52:38Z

## Translation

タイトル: クロード 5 世代モデルのコンテキスト エンジニアリングの新しいルール
記事のタイトル: Claude 5 世代モデルのコンテキスト エンジニアリングの新しいルール |クロード by Anthropic
説明: より高度なモデル用に、Claude Code のシステム プロンプトの 80% 以上を削除しました。私たちが学んだ教訓を、クロード コードおよび独自のエージェントを使用した独自のコンテキスト エンジニアリングに適用する方法。

記事本文:
Claude 5 世代モデルのコンテキスト エンジニアリングの新しいルール |クロード by Anthropic
クロード製品のご紹介 クロード
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
Claude 5 世代モデルのコンテキスト エンジニアリングの新しいルール
Claude 5 世代モデルのコンテキスト エンジニアリングの新しいルール
より高度なモデル用に、Claude Code のシステム プロンプトの 80% 以上を削除しました。私たちが学んだ教訓を、クロード コードおよび独自のエージェントを使用した独自のコンテキスト エンジニアリングに適用する方法。
製品 クロード コード クロード エンタープライズ クロード プラットフォーム
共有 リンクをコピー https://claude.com/blog/the-new-rules-of-context-engineering-for-claude-5-generation-models
私は以前、最新世代の Claude 5 モデルを最適にプロンプトし、それらを繰り返し操作して構築したいものを発見する方法について書きました。
しかし、クロードにメッセージを送信する場合、プロンプトは、メッセージが取得するコンテキストのほんの一部にすぎません。コンテキストの多くは、システム プロンプト、スキル、CLAUDE.md ファイル、メモリ、およびその他のソースから組み立てられます。私たちはこれを コンテキスト エンジニアリング と呼び、クロード コードを使用するとき、または独自のエージェントを構築するときに生成される結果に大きな影響を与えます。
プロンプトとは異なり、コンテキストは一般に多くのリクエストにわたって使用されるため、それほど具体的にすることはできません。クロードに対するこれらの一般的なプロンプトとガイダンスをどのように作成しますか?

特にユーザーのプロンプトが何であるかわからない場合はどうでしょうか?
クロード自身の能力が進化するにつれて、これは驚くほど困難になる可能性があります。つい最近、私たちは最新世代のクロード モデルのプロンプト方法が大きく変わったことに気づきました。 Claude Opus 5 や Claude Fable 5 などのモデルの Claude Code のシステム プロンプトの 80% 以上を削除しましたが、コーディング評価に目立った損失は発生しませんでした。
この新しいクラスのモデルのプロンプトについて私たちが学んだことと、それを利用してコンテキスト エンジニアリングを更新する方法を説明します。これらのベスト プラクティスは「claude Doctor」に組み込まれています。クロード コードでコマンド /doctor を使用して、スキルと CLAUDE.md ファイルのサイズを適切に調整してください。
全体として、システム プロンプトと CLAUDE.md ファイルおよびスキルの両方で、クロード コードを過剰に制約していることがわかりました。
たとえば、クロード コードの内部使用のトランスクリプトを読むと、システム プロンプト、スキル、ユーザー リクエストが互いに衝突するため、「必要に応じてドキュメントを残す」や「コメントを追加しないでください」など、単一のリクエストの中にいくつかの矛盾するメッセージが表示されます。
一般に、クロードはユーザーの意図を解釈して正しい答えを導き出すことができますが、何をすべきかを決定する前に、これらの重複および矛盾するメッセージについてより慎重に考える必要があります。
これらの制約は、かつては最悪のシナリオを回避するために必要でしたが、その後、それらの多くを削除し、代わりにモデルに周囲のコンテキストと判断を使用させることができることがわかりました。
さらに、Claude Code にはさらに多くのツールが追加されました。クロードは、記憶、情報、および指針のソースとして CLAUDE.md に依存していました。これで、メモリ、アーティファクト、スキルが得られ、クロードはこれらを使用して、セッション間でコンテキストを読み込み、共有する新しい方法を作成できます。
前へ 前へ 0 / 5 次へ 次へ Claude Code Desktop を入手
IRM https://claude.ai

/インストール.ps1 | iex コマンドをクリップボードにコピーする またはドキュメントを読む クロード コードを試す クロード コードを試す クロード コードを試す 開発者ドキュメント 開発者ドキュメント 開発者ドキュメント 電子書籍
これまでのコンテキスト エンジニアリングのベスト プラクティスには、神話になったものが数多くありました。含む：。
最初にクロード コードを展開したとき、クロードがファイルの削除などの最悪のシナリオを確実に回避する必要がありました。これは、常に真実であるとは限らない特に強力なガイダンスを提供することを意味します。たとえば、システム プロンプトで次のように言っていました。
コード内: デフォルトではコメントを記述しません。複数段落の文書文字列や複数行のコメント ブロックは決して書かないでください (最大 1 行の短い行)。ユーザーが要求しない限り、計画、決定、または分析のドキュメントを作成しないでください。中間ファイルではなく、会話のコンテキストに基づいて作業します。
しかし、プロンプトの特定のサブセットについては、このガイダンスは間違っています。ドキュメントの場合、ユーザーが独自の設定を持っている場合や、非常に複雑なコードの特定の部分に複数行のコメント ブロックが必要な場合があります。
それでも、古いモデル用のこれらのガードレールがなければ、クロードが書いたコメントは多くの場合不正確になるため、このトレードオフを受け入れる必要がありました。しかし、新しいモデルはより優れた判断力を備えており、明示的なルールがなくてもこれらの決定を適切に処理できます。
新しいシステム プロンプトでは、次のように言います。周囲のコードと同じように読めるコードを作成し、コメントの密度、名前、およびイディオムを一致させます。
ツールの使用に関する第一のルールは、その使用方法についてクロードに例を与えることでした。私たちの最新モデルでは、例を与えると実際には特定の探索スペースに制約されてしまうことがわかりました。
例を使用する代わりに、ツール、スクリプト、ファイルの設計についてもっと考えてください。Claude にはどのようなパラメータがあり、どうすればより表現力豊かにできるでしょうか?
たとえば、Todo ツールの例では、l

保留中、進行中、完了の間の列挙としてステータスを表示し、その使用方法についてクロードにヒントを与えます。 1 つのアイテムを進行中のままにする指示は、要求された動作を定義するのに役立ちます。
今すぐ: 段階的な開示を使用する
Claude Code はコーディングに重点を置いていたため、システム プロンプトにはコード レビューと検証の方法に関する詳細な情報が含まれていました。これらは常に必要なわけではありませんが、必要な場合には重要な情報となります。
それ以来、Claude Code は、適切なタイミングで適切なコンテキストをロードする、漸進的な開示を使用することに非常に熟練してきました。たとえば、検証とコード レビューを独自のスキルに移し、Claude Code が選択的に呼び出すことができるようにしました。
しかし、段階的な開示はスキルだけでなく、ツールにも使用されます。一部のツールは「遅延読み込み」です。つまり、エージェントはツールを使用する前に、ToolSearch を使用して完全な定義を検索する必要があります。これにより、必要になるまでコンテキストを使用しないツール (タスク ツールなど) をさらに使用できるようになります。
同じことは、独自の CLAUDE.md および Skill.md ファイルにも適用できます。よくある誤解は、これらを、遭遇する可能性のあるすべての既知の実践の中央リポジトリにしたいというものです。そうしないとクロードが見つけられないからです。代わりに、適切なタイミングでロードできるファイルのツリーを用意することを検討してください。
以前の Claude モデルでは、繰り返しの命令が必要になる場合や、コンテキスト ウィンドウの最初よりも最後に命令を聞く可能性が高くなることがありました。これは、システム プロンプトのメイン システム プロンプト内のツールへの参照や、ツールの説明内の指示が含まれる場合があることを意味します。
これらの繰り返しの例を削除し、システム プロンプトではなくツールの説明にツールの使用方法の指示を追加できることがわかりました。
次に、CLAUDE.mの記憶

dファイル
以前は、# ホットキーを使用して CLAUDE.md に自動的に書き込むことで、クロードのメモリに内容を保存することをユーザーに推奨していました。代わりに、クロードは仕事とあなたに関連する思い出を自動的に保存するようになりました。
プラン モードでは、Claude Code はプランを含むマークダウン ファイルに大きく依存します。これらのファイルを計画として保存しておくと、クロードが必要なときに参照できるようになりました。同様のベスト プラクティスのもう 1 つは、クロードが長期プロジェクトにまたがる作業中に参照できるように仕様をコードベースに保存することでした。
しかし、Claude はますます複雑な参照を処理できることがわかりました。クロードは、単純なマークダウン ファイルの代わりに、新しいアーティファクト機能によって作成された HTML アーティファクトを参照できます。
コードの形式でクロード参照を提供することもできます。仕様は、詳細なテスト スイート、またはクロードが移植する可能性のある別のコードベースの関数である場合もあります。
ルーブリックは参照の別の形式です。ルーブリックを使用すると、Claude は動的なワークフローを使用し、それらのルーブリックを使用して検証エージェントを起動することで、特定の分野での好み (たとえば、優れた API 設計とはどのようなものであるか) を試して検証できます。
これをすべてまとめると、コンテキストを組み立てるとどうなるでしょうか?
システム プロンプトは製品コンテキストと大きく結びついています。これは、どの製品で動作しているのか、何をしているのかをクロードに伝えます。クロード コードの場合、これを変更することはおそらくありませんが、独自のエージェント ハーネスを構築している場合は、ここに多くの時間を費やす必要があります。
CLAUDE.md を軽量にし、リポジトリの目的を簡単に説明しますが、トークンのほとんどはコードベース内の落とし穴に費やします。たとえば、型を 1 つのモノリシック ファイルに保持し、他の場所には保持しないようにコードを整理できます。クロードがファイル システムやリポジトリを見て知っておくべき「明白な」事柄を述べることは避けてください。
プログレッシブ d を使用する

たとえば、作業内容を検証するための独自の手順がいくつかある場合は、検証スキルを作成し、CLAUDE.md から参照します。
スキルは、クロードが必要なときに情報を見つけられるようにする軽量のガイドと考えてください。非常に重要な領域を除いて、過度に制約しないようにしてください。
長いスキルの場合は、可能な限り段階的な開示を使用するようにしてください。多くのファイルに分割して分割してください。
スキルが、あなた、あなたのチーム、または製品に特有の特定の意見、知識、またはベストプラクティスをコード化している場合に最適です。
ファイルを @ メンションして参照として含めることができます。参照により、クロードは現在の計画に関する詳細な情報を参照できます。
これは仕様ファイル、モックアップ、またはコードベース全体に含まれる場合があります。一般に、コード内のファイルを使用することをお勧めします。これは、クロードがよく知っている言語で明確で忠実度の高い指示をクロードに提供するためです。たとえば、デザインの HTML モックアップは、通常、デザインの説明やスクリーンショットよりも良い結果をもたらします。
システム プロンプト、スキル、CLAUDE.md ファイル全体にわたって、私たちと同じように単純化する必要があるかもしれません。私たちは、これを自動的に行うのに役立つ「claude Doctor」と呼ばれる新しいコマンドを公開しました。特により高度なモデルのプロンプトの詳細については、「Fable フィールド ガイド」を参照してください。
この記事は、Anthropic 社の技術スタッフのメンバーである Thariq Shihipar によって書かれました。
クロードとともに構築するチーム向けの製品ニュースとベスト プラクティスをさらに詳しくご覧ください。
クロード モデルの説明: ユースケースに最適なモデルの選択
Enterprise AI クロード モデルの説明: ユース ケースに最適なモデルの選択 クロード モデルの説明: ユース ケースに最適なモデルの選択 クロード モデルの説明: ユース ケースに最適なモデルの選択 クロード モデルの説明

目的: ユースケースに最適なモデルの選択 2026 年 7 月 22 日 スキルを使用してクロード コードで検証ループを構築する
Claude Code スキルを使用して Claude Code で検証ループを構築する スキルを使用して Claude Code で検証ループを構築する スキルを使用して Claude Code で検証ループを構築する スキルを使用して Claude Code で検証ループを構築する 2026 年 7 月 21 日 Anthropic が AI ネイティブ ソフトウェア開発ライフサイクルを確保する方法
Claude Code Anthropic が AI ネイティブ ソフトウェア開発ライフサイクルをどのように確保しているか Anthropic が AI ネイティブ ソフトウェア開発ライフサイクルをどのように確保しているか Anthropic が AI ネイティブ ソフトウェア開発ライフサイクルをどのように確保しているか 2026 年 7 月 16 日 Anthropic が Claude Code を使用して大規模なコード移行を実行する方法
Claude Code Anthropic が Claude Code を使用して大規模なコード移行を実行する方法 Anthropic が Claude Code を使用して大規模なコード移行を実行する方法 Anthropic が Claude Code を使用して大規模なコード移行を実行する方法 Anthropic が Claude Code を使用して大規模なコード移行を実行する方法 Claude を使用して組織の運用方法を変革する
製品のアップデート、ハウツー、コミュニティのスポットライトなど。毎月あなたの受信箱に配信されます。
購読する 購読する場合はメールアドレスを入力してください

[切り捨てられた]

## Original Extract

We removed over 80% of Claude Code's system prompt for more advanced models. How to apply the lessons we learned to your own context engineering in Claude Code and with your own agents.

The new rules of context engineering for Claude 5 generation models | Claude by Anthropic
Meet Claude Products Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
The new rules of context engineering for Claude 5 generation models
The new rules of context engineering for Claude 5 generation models
We removed over 80% of Claude Code's system prompt for more advanced models. How to apply the lessons we learned to your own context engineering in Claude Code and with your own agents.
Product Claude Code Claude Enterprise Claude Platform
Share Copy link https://claude.com/blog/the-new-rules-of-context-engineering-for-claude-5-generation-models
I’ve written previously about how to best prompt the newest generation of Claude 5 models and work with them iteratively to discover what you want to build.
But when you send a message to Claude, the prompt is only a small part of the context it gets. Much of your context is assembled from your system prompt, Skills, CLAUDE.md files, memory, and other sources. We call this context engineering , and it makes a big impact on the results you generate when using Claude Code or in building your own agents.
Unlike a prompt, context is used generally across many requests, so it cannot be as specific. How do you build these general prompts and guidance for Claude, especially when you don’t know what a user’s prompt might be?
This can be surprisingly difficult as Claude’s own capabilities evolve. Most recently, we noticed a large jump in the way we prompt the newest generation of Claude models. We removed over 80% of Claude Code’s system prompt for models like Claude Opus 5 and Claude Fable 5 with no measurable loss on our coding evaluations.
Here’s what we’ve learned about prompting this new class of models, and how you can utilize it to update your context engineering. We’ve put these best practices in `claude doctor;` use the command /doctor in Claude Code to rightsize your skills, and CLAUDE.md files.
Overall, we found that we were overconstraining Claude Code, both through our system prompt and in our CLAUDE.md files and skills.
For example, when we read transcripts of our own internal usage of Claude Code, we see several conflicting messages in a single request like “leave documentation as appropriate,” or “DO NOT add comments” as our system prompt, skills, and user requests clash with each other.
Generally, Claude can interpret the user’s intent to get to the right answer, but Claude must think more carefully about these overlapping and conflicting messages before deciding what to do.
And while these constraints were once needed to avoid worst case scenarios, we have since found we can delete many of them and let the model use surrounding context and judgement instead.
Additionally, Claude Code now has many more tools. Claude used to rely on CLAUDE.md as a source of memory, information, and guidance. Now we have memory, artifacts, and skills, which Claude can use to create new ways of loading and sharing context across sessions.
Prev Prev 0 / 5 Next Next Get Claude Code Desktop
irm https://claude.ai/install.ps1 | iex Copy command to clipboard Or read the documentation Try Claude Code Try Claude Code Try Claude Code Developer docs Developer docs Developer docs eBook
There were a number of previous context engineering best practices that had become myths. Including:.
When we first rolled out Claude Code, we needed to be sure that Claude avoided worst case scenarios, such as deleting files. This meant we would give particularly strong guidance that might not always be true, For example, in the system prompt we used to say:
In code: default to writing no comments. Never write multi-paragraph docstrings or multi-line comment blocks — one short line max. Don't create planning, decision, or analysis documents unless the user asks for them — work from conversation context, not intermediate files.
But for a certain subset of prompts, this guidance would be wrong. In the case of documentation, the user may have their own preferences, or specific parts of very complex code might need multi-line comment blocks.
Still, without these guardrails for older models, the comments Claude wrote would be incorrect in many cases and we had to accept this tradeoff. But newer models have better judgement and can handle these decisions well without explicit rules.
In the new system prompt we say: Write code that reads like the surrounding code: match its comment density, naming, and idiom.
The number one rule for tool usage was to give Claude examples on how to use them. With our newest models, we’ve found that giving examples actually constrains them to a certain exploration space.
Instead of using examples, think more about the design of your tools, scripts and files- what parameters does Claude have and how can they be more expressive?
For example, in the Todo tool example, just listing status as an enumeration between pending, in_progress, and completed, hints to Claude about how to use it. The instruction on keeping one item in_progress helps define our requested behavior.
Now: Use progressive disclosure
Because Claude Code was focused on coding, our system prompt included detailed information on how to do code review and verification. These were not always needed, but when they were, it was crucial information.
Since then, Claude Code has gotten very competent at using progressive disclosure- loading the right context at the right times. For example, we moved verification and code review into their own skills that Claude Code could selectively call.
But progressive disclosure is not just for skills, we also use it for tools. Some of our tools are ‘deferred loading,’ which means the agent must search for their full definitions using ToolSearch before using them. This allows us to have more tools (such as our Task tools) that don’t take up context until they’re needed.
The same can be applied to your own CLAUDE.md and Skill.md files. A common myth is that you want to make these a central repository for every known practice that you might run into, because Claude would not find it otherwise. Instead, consider having a tree of files that can be loaded at the right time .
Earlier Claude models could sometimes need repeated instructions or be more likely to listen to instructions at the end of their context window than at the start. This meant our system prompt would sometimes have references to tools in the main system prompt as well as instructions in the tool description.
We found we could delete these repeat examples and put instructions on how to use tools in the tool descriptions rather than the system prompt.
Then: Memory in CLAUDE.md files
We used to encourage users to save things to Claude’s memory, by using the # hotkey to write to their CLAUDE.md automatically. Instead, Claude now automatically saves memories that are relevant to the work and to you.
In plan mode, Claude Code has heavily relied on markdown files with plans. Storing these files as plans helped Claude refer to them when needed. Another similar best practice was to store specs in the codebase for Claude to refer to while working across longer projects.
But we’ve found that Claude can handle increasingly more complicated references. Instead of simple markdown files, Claude can reference HTML artifacts created by our new artifacts feature.
You may also give Claude references in the form of code. A spec may also be a detailed test suite, or a function in a different codebase that Claude might port.
Rubrics are another form of references. Rubrics allow Claude to try and verify your taste in a particular field (e.g. what does a good API design look like) by using dynamic workflows and spinning up verifier agents with those rubrics.
Pulling this all together, what does this look like when you assemble your context?
A system prompt is heavily tied to the product context. It tells Claude what product it’s operating in and what it’s doing. For Claude Code, you will likely never modify this, but if you are building your own agent harness, this is where you should spend a lot of time.
Keep your CLAUDE.md lightweight and briefly describe what your repo is for, but spend most of the tokens on gotchas inside of the codebase. For example, you may organize your code to keep types in one monolithic file and nowhere else. Avoid stating ‘the obvious’ things Claude should know by looking at your file system or your repo.
Use progressive disclosure heavily, for example if you have several unique instructions on how to verify your work, create a verification skill and reference it from your CLAUDE.md.
Think of skills as lightweight guides to let Claude find information when needed. Avoid making them overconstrained, except in highly important areas.
For long skills, try and use progressive disclosure as much as possible- divide it into many files and split them out.
It’s best when skills encode particular opinions, knowledge, or best practices that are particular to you, your team, or product.
You can @ mention files to include them as references. References allow Claude to refer to in-depth information about the current plan.
This might be in specs files, mockups, or even entire codebases. Generally you should prefer files that are in code as it provides clear, high-fidelity instructions to Claude in a language it knows very well. For example, a HTML mockup of a design will generally produce better results than a description of the design or a screenshot.
Across your system prompt, skills, and CLAUDE.md files, you may need to simplify just like we did. We rolled out a new command called `claude doctor,` which will help you do this automatically as well. For more details on prompting more advanced models specifically, check out our Fable field guide .
This article was written by Thariq Shihipar, member of technical staff, Anthropic.
Explore more product news and best practices for teams building with Claude.
Claude models explained: choosing the best model for your use case
Enterprise AI Claude models explained: choosing the best model for your use case Claude models explained: choosing the best model for your use case Claude models explained: choosing the best model for your use case Claude models explained: choosing the best model for your use case Jul 22, 2026 Building verification loops in Claude Code with skills
Claude Code Building verification loops in Claude Code with skills Building verification loops in Claude Code with skills Building verification loops in Claude Code with skills Building verification loops in Claude Code with skills Jul 21, 2026 How Anthropic secures its AI-native software development lifecycle
Claude Code How Anthropic secures its AI-native software development lifecycle How Anthropic secures its AI-native software development lifecycle How Anthropic secures its AI-native software development lifecycle How Anthropic secures its AI-native software development lifecycle Jul 16, 2026 How Anthropic runs large-scale code migrations with Claude Code
Claude Code How Anthropic runs large-scale code migrations with Claude Code How Anthropic runs large-scale code migrations with Claude Code How Anthropic runs large-scale code migrations with Claude Code How Anthropic runs large-scale code migrations with Claude Code Transform how your organization operates with Claude
Product updates, how-tos, community spotlights, and more. Delivered monthly to your inbox.
Subscribe Subscribe Please provide your email address if you'd like to

[truncated]
