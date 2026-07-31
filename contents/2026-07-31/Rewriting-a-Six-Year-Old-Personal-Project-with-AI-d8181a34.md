---
source: "https://www.hanqi-blog.com/translation/ai_coding"
hn_url: "https://news.ycombinator.com/item?id=49120492"
title: "Rewriting a Six-Year-Old Personal Project with AI"
article_title: "Rewriting a Six-Year-Old Personal Project with AI | 知与行"
author: "HanQi"
captured_at: "2026-07-31T08:34:01Z"
capture_tool: "hn-digest"
hn_id: 49120492
score: 1
comments: 0
posted_at: "2026-07-31T08:29:48Z"
tags:
  - hacker-news
  - translated
---

# Rewriting a Six-Year-Old Personal Project with AI

- HN: [49120492](https://news.ycombinator.com/item?id=49120492)
- Source: [www.hanqi-blog.com](https://www.hanqi-blog.com/translation/ai_coding)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T08:29:48Z

## Translation

タイトル: 6 年間の個人プロジェクトを AI で書き直す
記事のタイトル: 6 年間の個人プロジェクトを AI で書き直す | 知与行
説明: はじめに 最近、私は AI を使用して、実稼働環境で 6 年間保守および実行してきた個人プロジェクトを書き直しました。 AI のおかげで時間を大幅に節約でき、複雑な新機能の構築に役立ちました。同時に、現状の限界にも気づき、徐々に仕事のやり方を覚えていきました。
[切り捨てられた]

記事本文:
6 年間の個人プロジェクトを AI で書き直す
最近、私は AI を使用して、6 年間実稼働環境で保守および実行してきた個人プロジェクトを書き直しました。 AI のおかげで時間を大幅に節約でき、複雑な新機能の構築に役立ちました。同時に、現在の限界についてもより認識するようになり、より効果的に対処する方法を徐々に学びました。この記事では、そのプロセスとそこから得た教訓を記録します。
この書き直しから得た非常に重要なアイデアは、フィードバック ループです。
AI がタスクに取り組むときは、推論、ツールの呼び出し、結果の確認、反復といういくつかの段階を経ます。完全なフィードバック ループとは、AI が人間の介入なしに最初から最後までタスクを実行できることを意味します。人間は指示を与えるだけで済みます。ただし、実際には、多くのプログラミング タスクは完全なフィードバック ループを提供しません。設計を選択したり、不足しているツールを接続したり、AI だけでは評価できない結果を検証したりするには、依然として人間が介入する必要があります。
AI の強みと限界は、このループがどの程度完全であるかによって決まります。その結果、プログラマーの役割は、手作業でコードを書くことから、AI のための完全なフィードバック ループを構築することに徐々に移行しています。
AIはすでに高い能力を持っていますが、社会はAIを統合する準備がまだ整っていません。多くの形式の作業には、AI に必要なツール、インターフェイス、検証メカニズムがまだ不足しています。ソフトウェア エコシステムを取り巻く環境を構築することが、今後数年間で重要な方向性になる可能性があります。
私はもともとこのエッセイを中国語で書き、その後 GPT-5.6 を使用して英語翻訳を支援し、最終版を自分でレビューして編集しました。
昨年の初め、オンラインで Cursor、Codex、Claude Code が本当に便利になったと多くの人が言っているのを目にしました。私が取り組んでいたソフトウェアは、

ただし、時間には機密の本番データが含まれており、無制限の AI アクセスには適していませんでした。そのため、私は自分でコードを書いてデバッグしながら、主に研究目的で AI を使い続けました。
今年の初めに、友人が Codex を強く勧めてくれました。当時は GPT-5.4 を搭載していました。私は CPython の json ライブラリを Python で再実装する別のプロジェクトを終えたばかりだったので、Codex に構文にコメントのサポートを追加するよう依頼しました。数分以内にテストが作成され、機能が完成しました。私は驚きました。私はまだ AI を主にチャットボットとして考えていましたが、AI が独自にプログラミング機能を実装しているのを見ると、ほとんど魔法のように感じました。
ChatGPT が 2022 年に登場すると、世界中の注目を集めました。それでも、チャットボット インターフェイスでは LLM ができることは限られていると感じていました。実験は私たちが知識を獲得する主な方法の 1 つであるため、本物の知性は実験を実行し、世界からのフィードバックから学ぶことができるはずです。これにより、AI は事前トレーニングの限界を超え、インタラクションを通じて推論、学習、改善することが可能になります。当時、私はロボット アームや小型車両を制御するシステムを想像していました。それは遠くに思えた。しかし、2026 年までに、AI はすでにソフトウェア環境内で自律的に行​​動し、実験する能力を獲得していました。
私は最近、AI を独自に研究するために数か月を確保し、コーディング エージェントを探索するための最初のプロジェクトが必要でした。 Smalltalk に敬意を表して名付けられた SmallTool という個人用ツールをすでに持っていたので、AI を使用してそれを一から書き直し、新しい機能を追加することにしました。
このプロジェクトは 2020 年に遡ります。パンデミックが始まった当初、私は旧正月の休暇を利用して自分の ToDo リストを作成しました。スタックはシンプルでした。JavaScript、AJAX、Flask で、デプロイメントはコマンド li から手動で行われました。

ね。 2022 年の初めに、Vue と Django を使用して書き直しました。 2023年にTypeScript、React、mypyで再度書き直し、ノートブックを追加し、Ansibleでデプロイを自動化しました。このアプリケーションはまだかなり小さく、フロントエンド コードとバックエンド コードが約 3,000 行しかありませんでしたが、6 年間実稼働環境で確実に実行されていました。
2026 年にこのプロジェクトを再検討したとき、別の書き直しのアイデアがいくつかありました。
スタックを簡素化します。React を削除し、Django テンプレートを使用し、SQLite に切り替えます。
リアルタイム保存機能を備えた Markdown ベースのリッチ テキスト エディターを追加します。古いノートブックは基本的な HTML フォームのみでした。
ネストされたフォルダーとドラッグ アンド ドロップ操作を使用して、ノートブック用の Mac のようなファイル システムを構築します。
携帯電話の使用状況を追跡する Android クライアントを構築します。ストレスを感じているときは、ビデオや長いコンテンツを使いすぎる傾向があります。その習慣を減らすための最初のステップは、オンライン メディアに費やす時間を測定することです。
Garmin ウォッチから Web アプリケーションにデータをインポートします。 Garmin は大量のデータを収集しますが、分析ツールが限られているため、自分でデータを分析する自由が欲しかったのです。
したがって、これは単なるテクノロジースタックの変更以上のものでした。また、多くの新機能が追加され、2 つの外部デバイス (Garmin ウォッチと Android 携帯電話) からデータが取り込まれました。これは単なるリファクタリングではなく、大幅な拡張でした。
2024 年に、私は Web テクノロジー スタックに関する記事を書きました。今ではその積み重ねが全く違った見方で見られます。個人的なプロジェクトや企業製品の最初のバージョンとしては、重すぎたと思います。あまりにも複雑になりすぎて、開発が困難になってしまいました。この書き直しにより、単純化するとどうなるかを確認する機会が得られました。
この記事を書いた時点で、新しい SmallTool は 1 か月間稼働していました。この記事の最初のアイデアの一部もその中に記録されています。

このプロジェクトには、ToDo リスト、ノートブック、Garmin データ同期、電話使用状況追跡の 4 つの主要な部分があります。
ファイル システムとノートブックのインターフェイスは次のとおりです。
コア機能はモバイルでの使用に合わせて調整されており、調整可能なフォント サイズやダーク モードなど、多くの小さな UI の詳細を調整しました。
ほとんどのデータのプレゼンテーションはサーバー上で行われるため、Android インターフェイスはシンプルです。永続的なバックグラウンド通知により、2 時間ごとに使用状況データが収集されます。
また、Garmin データと Android の画面オフ時間も組み合わせて分析しています。自分の睡眠を一目で理解できるように、毎日を色でラベル付けしています。
コードベースを以下に示します。実稼働コードとテスト コードの比率は約 1.7:1 で、テスト スイートには多くのエンドツーエンド テストが含まれています。
前述したように、私は以前の個人プロジェクトの経験に基づいてスタックを意図的に単純化しました。 React、TypeScript、mypy を削除し、PostgreSQL を SQLite に置き換えました。それはどうなりましたか?
非常によく。フロントエンド スタックがシンプルになると、UI が悪化したり、コードの繰り返しが増えたりするのではないかと心配していましたが、そのようなことはありませんでした。コードは簡潔で、SQLite によりデバッグがはるかに簡単になりました。たとえば、後でエンドツーエンドのテストを並列化する場合、PostgreSQL サーバーを設定せずに、各プロセスに独自の SQLite データベース ファイルを与えるだけで済みます。
最初の週は、Kimi Code と K2.7 Code を使用しました。 2 週間目は、Codex と GPT-5.5 を使用しました。 OpenAI は、プロジェクトの最終日にモデルを GPT-5.6 にアップグレードしました。このプロジェクトでは GPT-5.5 と GPT-5.6 の大きな違いには気づきませんでしたが、大規模なコードベースでは違いがより明確になる可能性があります。
コーデックスの割り当てを使い果たすことがよくありました。それが起こったとき、私は現在の機能の引き継ぎドキュメントを生成し、作業を続行するよう Kim Code に依頼しました。
VS Codeでコードを読んで実行しました

統合端末内の Kim Code と Codex。
私は20ドルのコーデックスプランと99元のキミコードプランに加入しました。私の現在の使用法では、それらを交互に使用するだけで十分でした。
当初、中国モデルにはあま​​り期待していませんでしたが、K2.7 コードは驚くほど有能で、プロジェクト全体を通して非常にうまく機能しました。しかし、Kimi Code のインターフェースはあまり洗練されていないように感じられました。長いコマンドは時々切り詰められ、長時間使用すると端末に文字化けしたテキストが表示されることがありました。
コーデックスのほうが洗練されているように感じました。そのインターフェイスでは、より明確な強調表示が使用され、計画とアクションが読みやすい方法で要約され、使用量クォータがほぼ使い果たされた場合でも、現在のタスクが完了するまで続行されることがよくありました。それは寛大さと思いやりの両方を感じました。
最新のコーディング エージェントは、詳細なタスク ドキュメントから数時間、場合によっては数十時間実行することもできます。私がそのような「魔法」を意図的に避けたのは、まず AI の操作の詳細を理解したかったからです。私も事前に商品の詳細を知りませんでした。何を変更するかを決定する前に、動作中のバージョンを確認する必要があることがよくありました。
AI はこのプロジェクトに大きな影響を与えてくれました。時間が節約され、手作業が減り、より興味深いと思った部分に集中できるようになりました。
プロジェクトには約 15 日間、およそ 40 時間の作業時間がかかりました。 AI がなければ、少なくとも 80 時間は必要だったと思います。
その40時間も比較的楽でした。友人と電話で話しながら、UIの細部を調整しました。
40 時間と 80 時間の差は、単に速度が 2 倍変化しただけではありません。 AI がなかったら、私はプロジェクトを放棄していたかもしれません。すべてを手書きで書くと、多くの詳細を把握して推論する必要があり、その作業は非常に疲れるでしょう。おそらく機能を削除し、UI の調整をスキップしたでしょう。

他の優先事項が自分の時間を奪い合っていたため、40 時間以内に収まるまでプロジェクトを短縮しました。
また、CSS や Android 開発を詳しく学ぶ忍耐力もエネルギーもありませんでした。もし私がすべてを自分で書いていたら、おそらく UI はもっと粗くなり、電話使用量の収集を完全に削除していたかもしれません。
AI はすでにソフトウェアの作成とデバッグに非常に優れています。
このプロジェクトを通じて、私の手作業によるコーディングとデバッグのほぼすべてが AI に置き換えられました。完全なフィードバック ループにアクセスできるようになると、人間の介入はほとんど必要なくなりました。
いくつかの例を見て私は驚きました。 Android アプリケーションの開発中、Codex は私の MacBook から Android デバッグ ツールを直接使用できました。Codex は、スクリーンショットを撮り、ファイルをアップロードおよびダウンロードし、アプリケーションの SQLite データベースを検査し、コマンドを実行し、APK をコンパイルして、携帯電話にインストールしました。私がデバイスに触れずにループ全体を完了できる可能性があります。デプロイ中に、Ansible を使用してサーバー上でコマンドを実行し、SQLite を検査し、運用エンドポイントを呼び出し、Nginx ログを読み取りました。自分で実稼働操作を実行する必要はありませんでした。
私の経験では、AI に明確な目標を与え、その結果を完全に検証できれば、プロセスは高度に自動化できると感じました。
このような状況では、人間は正確な指示を提供し、結果を確認するのを待つだけで済みます。
AI は、より高度な工学的判断も示しました。たとえば、Kimi がファイル ツリーを構築したとき、最初にネストされたフォルダーのないバージョンを実装し、それを検証してからネストを追加することを提案しました。これは古典的な反復アプローチです。細部を磨き上げる前に、小さな実験を実行してアイデアをテストします。
AI は依然として完全な開発プロセスを単独で完了することができず、ユーザーからのフィードバックがボトルネックになっています。
いくつかの

技術的な決定に欠陥があった。
ノートブックに自動保存を実装した場合、AI は競合を適切に処理できませんでした。その後、私がアルゴリズムと UI デザインを提供しました。私の使用例では、クライアントは各更新にメモの最後に保存されたタイムスタンプを含めます。サーバーは、送信されたタイムスタンプが古いと判断すると、ユーザーに更新するよう求めます。書いたばかりのテキストをコピーし、ページを更新して、コンテンツを失うことなく編集を続けることができます。
ごみ箱を実装するとき、AI は最初、削除されたファイルにステータス フィールドをマークしようとしました。この変更は、実装が方向性を失うまでコードベース全体に広がり続けました。これを、既存のフォルダー システムを再利用する、よりシンプルなデザインに置き換えました。1 つのフォルダーがごみ箱として予約されています。新しい実装は迅速で、変更はほとんど必要ありませんでした。
Android 開発の 2 回目のラウンド中に、アプリケーションは Web API を呼び出すことができなくなりました。 AI は徐々に低レベルの Android デバッグ ツールに移行していきました。それは私には正しくありませんでした。以前にも同じ設定が機能していたことを思い出させ、その原因が基本的なホスト構成の問題である可能性があることを示唆しました。すぐに問題を見つけて修正してくれました。
ユーザーからのフィードバックを得ることで開発ループもブロックされました。オルタナティブ

[切り捨てられた]

## Original Extract

Introduction Recently, I used AI to rewrite a personal project that I had maintained and run in production for six years. AI saved me a great deal of time and helped me build complex new features. At the same time, I also became more aware of its current limitations and gradually learned how to work
[truncated]

Rewriting a Six-Year-Old Personal Project with AI
Recently, I used AI to rewrite a personal project that I had maintained and run in production for six years. AI saved me a great deal of time and helped me build complex new features. At the same time, I also became more aware of its current limitations and gradually learned how to work with it more effectively. This article records the process and the lessons I drew from it.
A very important idea I took away from this rewrite is the feedback loop .
When AI works on a task, it goes through several stages: reasoning, calling tools, checking results, and iterating. A complete feedback loop means that AI can carry the task from start to finish without human intervention; the human only needs to provide the instruction. In practice, however, many programming tasks do not offer a complete feedback loop. Humans still need to step in to choose a design, connect missing tools, or verify results that AI cannot evaluate on its own.
Both the strengths and the limits of AI are shaped by how complete this loop is. As a result, the programmer’s role is gradually shifting from writing code by hand to building complete feedback loops for AI.
AI is already highly capable, but society is not yet fully prepared to integrate it. Many forms of work still lack the tools, interfaces, and verification mechanisms that AI needs. Building that surrounding software ecosystem may become an important direction in the years ahead.
I originally wrote this essey in Chinese，then used GPT-5.6 to assist with the English translation and reviewed and edited the final version myself.
At the beginning of last year, I saw many people online saying that Cursor, Codex, and Claude Code had become genuinely useful. The software I worked on at the time, however, involved sensitive production data and was not suitable for unrestricted AI access. I therefore continued to use AI mainly for research while writing and debugging the code myself.
At the beginning of this year, a friend strongly recommended Codex to me. It was powered by GPT-5.4 at the time. I had just finished another project, reimplementing CPython's json library in Python , so I asked Codex to add support for comments to the syntax. Within a few minutes, it had written tests and completed the feature. I was amazed. I had still thought of AI mainly as a chatbot, and seeing it independently implement a programming feature felt almost magical.
When ChatGPT appeared in 2022, it attracted worldwide attention. Even then, I felt that the chatbot interface limited what an LLM could do. Real intelligence should be able to run experiments and learn from feedback from the world, because experimentation is one of the main ways we acquire knowledge. That would allow AI to go beyond the limits of pretraining and to reason, learn, and improve through interaction. At the time, I imagined systems controlling robotic arms or small vehicles. That seemed far away. Yet by 2026, AI had already gained the ability to act and experiment autonomously inside software environments.
I recently set aside a few months to study AI independently and needed a first project for exploring coding agents. I already had a personal tool called SmallTool, named in tribute to Smalltalk, so I decided to use AI to rewrite it from the ground up and add new features.
The project dates back to 2020. At the beginning of the pandemic, I used the Lunar New Year holiday to build a todo list for myself. The stack was simple: JavaScript, AJAX, and Flask, with deployment done manually from the command line. In early 2022, I rewrote it with Vue and Django. In 2023, I rewrote it again with TypeScript, React, and mypy, added a notebook, and automated deployment with Ansible. The application was still fairly small, with only about 3,000 lines of frontend and backend code, but it had been running reliably in production for six years.
When I revisited the project in 2026, I had several ideas for another rewrite:
Simplify the stack: remove React, use Django templates, and switch to SQLite.
Add a Markdown-based rich-text editor with real-time saving. The old notebook was only a basic HTML form.
Build a Mac-like file system for the notebook, with nested folders and drag-and-drop movement.
Build an Android client to track my phone usage. When I am under stress, I tend to overuse videos and long-form content. The first step toward reducing that habit is measuring how much time I spend on online media.
Import data from my Garmin watch into the web application. Garmin collects a great deal of data, but its analysis tools are limited, so I wanted the freedom to analyze the data myself.
This was therefore more than a change of technology stack. It also added many new features and brought in data from two external devices: my Garmin watch and Android phone. It was not merely a refactor, but a substantial expansion.
In 2024, I wrote an article about my web technology stack . I now see that stack very differently. For a personal project, or the first version of a company product, I think it was too heavy. It introduced too much complexity and made development harder. This rewrite gave me a chance to see what would happen if I simplified it.
By the time I wrote this article, the new SmallTool had been running for a month. Some of the first ideas for this article were even recorded in it.
The project has four main parts: a todo list, a notebook, Garmin data synchronization, and phone-usage tracking.
Here are the file system and notebook interfaces:
The core features are adapted for mobile use, and I refined many small UI details, including adjustable font sizes and a dark mode.
The Android interface is simple because most data presentation happens on the server. A persistent background notification collects usage data every two hours:
I also combine Garmin data with Android screen-off time for analysis. To understand my sleep at a glance, I label each day with a color:
The codebase is shown below. Production code and test code have a ratio of about 1.7:1, and the test suite includes many end-to-end tests.
As mentioned earlier, I deliberately simplified the stack based on my experience with previous personal projects. I removed React, TypeScript, and mypy, and replaced PostgreSQL with SQLite. How did that work out?
Very well. I had worried that the simpler frontend stack might produce a worse UI or more repetitive code, but it did not. The code is concise, and SQLite made debugging much easier. For example, when I later parallelized the end-to-end tests, I could simply give each process its own SQLite database file without setting PostgreSQL server。
During the first week, I used Kimi Code with K2.7 Code. During the second week, I used Codex with GPT-5.5. OpenAI upgraded the model to GPT-5.6 during the final days of the project. I did not notice a large difference between GPT-5.5 and GPT-5.6 on this project, although the difference might be clearer on a larger codebase.
I often exhausted my Codex quota. When that happened, I generated a handoff document for the current feature and asked Kimi Code to continue the work.
I read the code in VS Code and ran Kimi Code and Codex in its integrated terminal.
I subscribed to the $20 Codex plan and the 99-yuan Kimi Code plan. Alternating between them was enough for my current usage.
I initially had low expectations for a Chinese model, but K2.7 Code was surprisingly capable and worked very well throughout the project. Kimi Code’s interface, however, felt much less polished. Long commands were sometimes truncated, and the terminal occasionally displayed garbled text after extended use.
Codex felt more refined. Its interface used clearer highlighting, summarized its plans and actions in a readable way, and often continued until it finished the current task even when the usage quota was almost exhausted. That felt both generous and considerate.
Modern coding agents can also run for hours, or even tens of hours, from a detailed task document. I deliberately avoided that kind of “magic” because I first wanted to understand the details of working with AI. I also did not know all the product details in advance. I often needed to see a working version before I could decide what to change.
AI gave me enormous leverage on this project. It saved time, reduced manual work, and allowed me to focus on the parts I found more interesting.
The project took about 15 days and roughly 40 hours of work. Without AI, I estimate that it would have required at least 80 hours.
Those 40 hours were also relatively easy. I adjusted some UI details while talking with friends on the phone.
The difference between 40 and 80 hours was not merely a twofold change in speed. Without AI, I might have abandoned the project. Writing everything by hand would have required me to hold and reason about many details, and the work would have been exhausting. I would probably have removed features, skipped UI refinement, and reduced the project until it fit within 40 hours because I had other priorities competing for my time.
I also did not have the patience or energy to study CSS and Android development in depth. If I had written everything myself, the UI would probably have been rougher, and I might have dropped phone-usage collection entirely.
AI is already very good at writing and debugging software.
Throughout this project, AI replaced almost all of my manual coding and debugging. When it had access to a complete feedback loop, it needed very little human intervention.
A few examples surprised me. While developing the Android application, Codex could use Android debugging tools directly from my MacBook: it took screenshots, uploaded and downloaded files, inspected the application’s SQLite database, ran commands, compiled the APK, and installed it on my phone. It could complete the entire loop without me touching the device. During deployment, it used Ansible to run commands on the server, inspect SQLite, call production endpoints, and read Nginx logs. I did not need to perform any production operations myself.
My experience was that when I could give AI a clear goal and it could fully verify the result, the process could be highly automated.
In those situations, the human only needed to provide a precise instruction and wait to review the result.
AI also showed some higher-level engineering judgment. When Kimi built the file tree, for example, it suggested implementing a version without nested folders first, validating it, and only then adding nesting. This is a classic iterative approach: run a small experiment to test the idea before polishing every detail.
AI still cannot complete the full development process on its own, and user feedback has become a bottleneck.
Several of its technical decisions were flawed.
When implementing autosave for the notebook, AI did not handle conflicts well. I later supplied the algorithm and UI design. For my use case, the client includes the note’s last-saved timestamp with each update. If the server sees that the submitted timestamp is stale, it asks the user to refresh. I can copy the text I just wrote, refresh the page, and continue editing without losing content.
When implementing the recycle bin, AI initially tried to mark deleted files with a status field. The change kept spreading through the codebase until the implementation lost direction. I replaced it with a simpler design that reused the existing folder system: one folder is reserved as the recycle bin. The new implementation was fast and required very few changes.
During a second round of Android development, the application could no longer call the web API. AI gradually moved toward increasingly low-level Android debugging tools. That did not feel right to me. I reminded it that the same setup had worked before and suggested that the cause might be a basic host configuration issue. It quickly found and fixed the problem.
Obtaining user feedback also blocked the development loop. Alt

[truncated]
