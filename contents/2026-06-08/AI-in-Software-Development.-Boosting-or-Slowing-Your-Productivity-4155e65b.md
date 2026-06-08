---
source: "https://olegdubovoi.com/publications/ai-in-software-development-boosting-or-slowing-your-productivity/"
hn_url: "https://news.ycombinator.com/item?id=48447215"
title: "AI in Software Development. Boosting or Slowing Your Productivity"
article_title: "AI in Software Development. Boosting or Slowing Your Productivity · Oleg Dubovoi's Blog"
author: "empiree"
captured_at: "2026-06-08T16:27:02Z"
capture_tool: "hn-digest"
hn_id: 48447215
score: 1
comments: 0
posted_at: "2026-06-08T16:11:49Z"
tags:
  - hacker-news
  - translated
---

# AI in Software Development. Boosting or Slowing Your Productivity

- HN: [48447215](https://news.ycombinator.com/item?id=48447215)
- Source: [olegdubovoi.com](https://olegdubovoi.com/publications/ai-in-software-development-boosting-or-slowing-your-productivity/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T16:11:49Z

## Translation

タイトル: ソフトウェア開発における AI。生産性の向上または低下
記事のタイトル: ソフトウェア開発における AI。生産性の向上または低下 · Oleg Dubovoi のブログ
説明: AI は実際に開発者の生産性をどの程度向上させますか? METR の調査、2025 年の StackOverflow 調査、私自身の経験、そしてバイブコーディングのトレンドについて見てみましょう。

記事本文:
このページのマークダウン バージョンは、/publications/ai-in-software-development-boosting-or-slowing-your-productivity.md で入手できます。
オレグ・ドゥボヴォイのブログ
AI、テクノロジー、ソフトウェア開発
L'AI est-elle vraiment effice
ソフトウェア開発における AI。ソフトウェア開発における生産性 AI の向上または低下。 AI はソフトウェアの生産性を向上させますか、それとも低下させますか。開発論理上の製品開発を目指します。 AI は、開発における生産性を向上させます。
コメント シェア EN RU DE FR ZH AI翻訳 はじめに
現在、ソフトウェア開発における AI の影響を無視することはできません。 Google で情報を検索する代わりに、ChatGPT に質問するだけで済みます。さらに素晴らしいことに、パズルのように情報をつなぎ合わせようとインターネット上で延々と検索して時間を無駄にすることなく、リクエストをより具体的にして、より詳細な情報を取得できるようになりました。正直なところ、最後に StackOverflow にアクセスしたのがいつだったかさえ思い出せません。2018 年当時は、必要な情報を見つけたり、回答を提供して他の人を助けるために 10 ～ 15 個のアクティブなタブを開いたものでした。 AI は、情報に加えて、サーバーから受信した大規模なログの分析や、人間でははるかに時間がかかるような 500 行以上の構成ファイル内のエラーの発見にも役立ちます。
AI は技術文書の作成も支援します。私がオープンソース ライブラリに取り組んでいたとき、ChatGPT は XML ドキュメントの 80% を作成し、GitHub リポジトリ用の優れたドキュメント ファイルの作成にも役立ちました。
しかし最も重要なことは、AI がコードを書けることです。約6か月間、私は月額20ドルを支払い続けています。

Anthropic の Claude Code を購読しているのは、生産性が向上し、いくつかの日常的なタスクをより速く解決できるからです。さらに、素早い情報検索や技術文書の作成にも ChatGPT を使用しています。
StackOverflow によると、プロの開発者の約 51% が AI ツールを日常的に使用しており、これはかなり大きな数字です。
出典: StackOverflow 調査 2025
最近、「バイブコーディング」という言葉がよく使われるようになりました。これは、AI がコードを書く新しいスタイルのプログラミングを指します。すべては、アルゴリズムのタスクを解決するために AI を使用することから始まりましたが、プログラミングの仕方を知らない人でも自分で何かを作ろうとするまでに進化しました。 Cursor のような、より強力なモデルやコード エディターの登場により、このトレンドに関する話題は高まるばかりです。現在では、AI は 1 つのプロンプトを入力するだけで、設計からビジネス ロジックに至るまで大量のコードを生成し、それがどのように機能するのか、なぜ機能するのかを説明できるようになりました。
StackOverflow の調査によると、バイブ コーディングに興味がある開発者は 12 ～ 15% のみです。これらの数字はまだ小さいですが、方向性はすでに形成されており、将来的には成長し続けるでしょう。
スタートアップに取り組み始めたとき、フロントエンド コードを記述する必要がよくありましたが、正直に言うと、Angular を使用する場合を除いて、これはあまり楽しくありませんでした。この場合、クロードは、モバイル対応デザインを含む基本レイアウトの作成とフレームワークとのリンクを手伝ってくれました。その後、コンポーネントを手動で改善し、外観を調整して、すべてが機能するようにします。
ここでのリスクは最小限です。 AI はビジネス ロジック、データベース、支払いシステムを扱いません。起こり得る最悪の事態は、AI が期待した結果を与えず、数時間を無駄にしてしまうことです。
それでも、雰囲気だけを楽しむべきです

完全に理解していないことや自分ではできないことを AI に依存している場合ではなく、何が起こっているかを理解し、プロセスを最適化しようとしている場合はコーディングを行う必要があります。
AI 研究グループ METR がこの夏に発表した調査では、AI コーディング ツールが経験豊富な開発者の生産性向上に本当に役立つかどうかが疑問視されています。
この調査では、METR は 16 人の経験豊富なオープンソース開発者に、大規模なコード リポジトリで 246 のタスクを完了させました。タスクの半分では Cursor Pro などの AI ツールの使用が許可されていましたが、残りの半分では許可されていませんでした。
始める前、開発者は AI を使えばタスクを 24% 早く完了できるだろうと考えていました。しかし、結果は驚くべきものでした。「AI を使用すると、実際に 19% 速度が遅くなった」と研究者たちは述べています。
注目すべきことに、調査に参加した開発者の 56% のみが、調査で提供された主要な AI ツールである Cursor の使用経験がありました。ほぼすべての開発者 (94%) がコーディング ワークフローでいくつかの Web ベースの LLM を使用した経験がありましたが、一部の開発者が特に Cursor を使用したのはこの調査が初めてでした。研究者らは、開発者が研究の準備として Cursor の使用に関するトレーニングを受けていたことを指摘しています。
これらの結果は、AI ツールが開発者の作業を常に高速化するかどうかについて疑問を引き起こします。研究者らは、開発者が AI に助けを求め、応答を待つことに多くの時間を費やしており、それが開発の速度を低下させていると考えています。また、AI は、このテストで使用されたような大規模で複雑なコードベースに苦戦します。
この研究の著者らは、これらの調査結果から強い結論を導き出さないように注意しており、現時点で AI システムが多くのソフトウェア開発者の高速化に失敗しているとは考えていないことを明確に指摘しています。他の大規模な研究では、AI コーディング ツールがソフトウェア エンジニアのワークフローを実際に高速化することが示されています。
同時に、Reddit では、開発者が AI に関する経験を共有しています。たとえば、投稿では

「FAANG でコードをどのようにバイブレーションするか」では、小規模なスタートアップ企業ではなく、大規模な IT 企業において、AI によって機能開発パフォーマンスが約 30% 向上したと述べられています。コードの作成に加えて、AI はテストの作成やコード レビューの高速化にも役立ちました。
出典: FAANG でコードをどのように雰囲気付けるか
おわかりのとおり、職場での AI の使用はすべての問題を解決する特効薬ではありません。これは主に、経験豊富な開発者を支援することもあれば、速度を低下させることもできるツールです。私の個人的な経験から言えば、AI は、独自のソリューションを必要としない、単純で局所的なタスクに最適です。ただし、コンテキストが複雑になり、タスクが大規模になると、作業の速度が低下する可能性があります。
2017 年、私が学生だった頃、C++ の先生は私たちに知識と記憶のみに基づいてコードを書くように教えました。私たちは紙やメモ帳にコードを書きませんでした。当時、私たちは Visual Studio 2015 を使用していましたが、IntelliSense はそれほど開発されておらず、ReSharper のようなプラグインについては知りませんでした。私たちにとって何が大きな発見だったか知っていますか?ショートカット。このとき、キーワード (「for」や「switch」など) を入力して「tab」を押すと、IDE がコードを書きます。正確にすべてを記述するのではなく、開発をスピードアップするためのテンプレート構造を生成します。私たちはそれが本当にクールだと思いましたが、すぐにあきらめました。なぜ？なぜなら、私たちは自分たちでコードを書くことを学んでいたからです。
Jeffrey Richter 著の CLR via C# や Steve McConnell 著の Code Complete などの本を読むことはできますが、何千行 (あるいは何万行も) のコードを書かない限り、適切なプログラミング方法を学ぶことはできないか、プログラミングの仕方がうまくできません。 C++ の学習中に、リンク リストやバイナリ ツリーを作成し、メモリを大量に使用し、Windows コンソールで三目並べやフィフティーンなどのゲームも作成しました。すべてのコードを書き終えたときの気分がどれほど素晴らしいかご存知でしょう。

AI や助けを借りずに、あなたがそれを実行しますか?
StackOverflow の年次調査によると、プログラミングを学習したばかりの 18 ～ 24 歳の人の約 70% が AI を使用しています。
私の意見では、AI は資料を簡単な言葉で説明するのに役立ちます。これは大きな利点であり、指導者がいない場合でも、指導者のように AI と話すことができます。ただし、自分でコードを作成し、何かが機能しない理由を時間をかけて理解する必要があります。そうして初めて知識が定着し、結果が生まれます。したがって、学習に対する AI の影響を最小限に抑えるようにしてください。提供されるコードに完全に依存しないでください。多くの場合、誤った情報が提供され、混乱するだけであるため、必ず他の情報源も参照してください。
自分でコードを書くこと、AI と連携すること、そして AI をどこで使用するのが最適かを理解できるようになると、生産性が真に向上します。
ChatGPT、Claude、Cursor などのツールの登場は、ソフトウェア開発業界にとって大きなプラスです。これらは、プログラマーにとって日常的なタスクを削減し、スピードを上げ、開発を容易にするのに役立ちます。ただし、AI がうまく処理できないタスクもまだあります。したがって、何よりもまず、自分の知識に頼って、常に学び続ける必要があります。
Сегодня уже невозможно игнорировать влияние AI на разработку ПО. Google にアクセスして、ChatGPT にアクセスして、チャットを開始してください。 Солее конкретным и получить более подробную информацию, а не тратить часы на бесконечный поиск по интернету в надежде собрать ответ из кусочков, как пазл. Честно говоря, я уже не п

最後に StackOverflow にアクセスしたときのことを覚えています。2018 年には、情報を検索したり他の人の回答を手伝ったりするために、通常 10 ～ 15 個のアクティブなタブを開いていました。 AI は、情報に加えて、サーバーからの大規模なログの分析や、人間でははるかに時間がかかる 500 行以上の構成内のエラーの発見にも役立ちます。
AI は技術文書の作成にも役立ちます。私がオープンソース ライブラリに取り組んでいたとき、ChatGPT は XML ドキュメントの 80% を作成し、GitHub リポジトリ用の優れたドキュメント ファイルをまとめるのに役立ちました。
しかし最も重要なことは、AI はコードを書くことができるということです。私は Anthropic の Claude Code のサブスクリプションに月額 20 ドルを約 6 か月間支払っています。これにより、生産性が向上し、日常的なタスクをはるかに速く完了できるようになりました。また、ChatGPT を使用して情報をすばやく検索し、技術ドキュメントを作成します。
StackOverflow によると、プロの開発者の約 51% が毎日 AI ツールを使用しており、これは非常に重要な数字です。
出典: StackOverflow 調査 2025
最近

「バイブコーディング」という言葉はかなり有名になりました。 AI がコードを書く新しいスタイルのプログラミングについて説明しています。すべてはアルゴリズムの問​​題を解決するために AI を使用することから始まりましたが、コーディングの仕方を知らない人でも自分で何かを作ろうとするところまで来ています。より強力なモデルや Cursor のようなコード エディターの出現により、このトレンドに関する話題は高まるばかりです。現在、1 つのプロンプト AI で、レイアウトからビジネス ロジックに至るまで大量のコードを生成し、それがどのように機能するのか、なぜ機能するのかを説明できるようになりました。
StackOverflow 調査によると、バイブ コーディングを行う開発者は 12 ～ 15% のみです。数はまだ少ないですが、方向性はすでに形成されており、今後も成長していくでしょう。
スタートアップに取り組み始めたとき、フロントエンドを書かなければならないことがよくありましたが、正直に言うと、Angular の場合を除いて、フロントエンドはあまり好きではありませんでした。ここでクロードが私を助けてくれました。彼はモバイル対応デザインを含む基本的なレイアウトを作成し、それをフレームワークとリンクさせました。次に、コンポーネントを手動で変更し、外観を調整し、すべてが適切であることを確認しました。

それはうまくいきました。
ここでのリスクは最小限です。 AI はビジネス ロジック、データベース、決済システムとは連携しません。起こり得る最悪の事態は、AI が間違った結果を出し、数時間を失うことです。
ただし、バイブコーディングは、何が起こっているのかを理解し、プロセスを最適化しようとしている場合にのみ行う価値があり、理解できず自分ではできないことをAIに頼っている場合ではありません。
AI 研究グループ METR がこの夏に発表した調査では、AI ツールが実際に経験豊富な開発者の生産性向上に役立つかどうかが尋ねられました。
この調査では、METR は 16 人の経験豊富なオープンソース開発者に、大規模なリポジトリで 246 のタスクを完了するよう依頼しました。タスクの半分では Cursor Pro などの AI ツールの使用が許可されていましたが、残りの半分では許可されていませんでした。
開始前、開発者は AI を使えばタスクを 24% 早く完了できると考えていました。しかし、結果は驚くべきものでした。「AI を使用すると、実際に速度が 19% 遅くなりました」と研究者らは述べています。
重要なのは、開発者の 56% だけが

この調査では、提供されていた主要な AI ツールである Cursor の使用経験がありました。ほぼ全員 (94%) がすでにプロセスで何らかの Web ベースの LLM を使用していましたが、Cursor を初めて見た人もいました。研究者らは、開発者が研究前に Cursor を使用する訓練を受けていたことを指摘しています。
これらの結果は、AI ツールが常にパフォーマンスを高速化するかどうかに疑問を投げかけます。

[切り捨てられた]

## Original Extract

How much does AI actually boost developer productivity? A look at the METR study, the 2025 StackOverflow Survey, my own experience, and the vibe coding trend.

A Markdown version of this page is available at /publications/ai-in-software-development-boosting-or-slowing-your-productivity.md.
Oleg Dubovoi’s Blog
AI, Tech and Software Development
L’AI est-elle vraiment efficace
AI in Software Development. Boosting or Slowing Your Productivity AI в разработке ПО. Ускоряет или замедляет вашу продуктивность AI in der Softwareentwicklung. Beschleunigt oder bremst deine Produktivität L'AI dans le développement logiciel. Booster ou ralentir ta productivité AI 在软件开发中。提高还是拖慢你的生产力
Comment Share EN RU DE FR ZH Translated by AI Introduction
Nowadays, it’s impossible to ignore the influence of AI on software development. Instead of searching for information on Google, you can now simply ask ChatGPT, and what’s even cooler, you can make your request more specific and get more detailed information, instead of wasting time endlessly searching across the internet hoping to piece together the information like a puzzle. Honestly, I can’t even remember the last time I visited StackOverflow, while back in 2018 I used to have 10-15 active tabs open to find the information I needed or to help others by providing answers. Besides information, AI can help you analyze large logs you might have received from a server or find errors in a config file with 500+ lines, something that would take a human much longer.
AI can also assist in writing technical documentation. When I was working on my open-source library, ChatGPT wrote 80% of the XML documentation and also helped create a good documentation file for the GitHub repository.
But the most important thing is that AI can write code. For about six months now, I’ve been paying $20 a month for a subscription to Claude Code by Anthropic because it boosts my productivity and allows me to solve some routine tasks much faster. Additionally, I use ChatGPT for quick information searches or writing technical documentation.
According to the StackOverflow, about 51% of professional developers use AI tools on a daily basis , which is a pretty significant number.
Source: StackOverflow Survey 2025
Recently, the term “vibe coding” has become quite popular. It refers to a new style of programming where AI writes the code for you. It all started with using AI for solving algorithmic tasks, but it has evolved to the point where even people who don’t know how to program are trying to create something on their own. With the arrival of more powerful models and code editors like Cursor, the buzz around this trend is only growing. Now, with just one prompt, AI can generate a large amount of code, from design to business logic, and explain how and why it works.
According to a StackOverflow Survey, only 12-15% of developers are into vibe coding . While these numbers are still small, the direction is already formed and will continue to grow in the future.
When I started working on my startups, I often needed to write frontend code, which, to be honest, I didn’t really enjoy, except when working with Angular. In this case, Claude helped me by creating the basic layout, including mobile-responsive designs, and linking it with frameworks. After that, I would manually improve the components, adjust the appearance, and make everything work.
The risks here are minimal. AI doesn’t deal with business logic, databases, or payment systems. The worst thing that could happen is that AI doesn’t give me the expected result, and I waste a few hours.
Still, you should only engage in vibe coding if you understand what’s happening and are trying to optimize processes, not if you’re relying on AI to do something you don’t fully understand or can’t do yourself.
A study published this summer by the AI research group METR questioned whether AI coding tools really help experienced developers be more productive.
In the study, METR had 16 experienced open-source developers complete 246 tasks on large code repositories. Half of the tasks allowed them to use AI tools like Cursor Pro, while the other half didn’t.
Before starting, the developers thought AI would help them finish their tasks 24% faster. But the results were surprising: “Using AI actually made them 19% slower” the researchers said.
Notably, only 56% of the developers in the study had experience using Cursor, the main AI tool offered in the study. While nearly all the developers (94%) had experience using some web-based LLMs in their coding workflows, this study was the first time some used Cursor specifically. The researchers note that developers were trained on using Cursor in preparation for the study.
These results raise doubts about whether AI tools will always make developers faster. The researchers believe that developers spend a lot of time asking AI for help and waiting for responses, which slows them down. Also, AI struggles with large, complex codebases, like the ones used in this test.
The study’s authors are careful not to draw any strong conclusions from these findings , explicitly noting they don’t believe AI systems currently fail to speed up many or most software developers. Other large-scale studies have shown that AI coding tools do speed up software engineer workflows.
At the same time, on Reddit, developers share their experiences with AI. For example, in the post “How we vibe code at a FAANG”, it’s mentioned that AI increased feature development performance by about 30% , not in a small startup, but in a large IT company. In addition to writing code, AI also helped write tests and sped up code reviews.
Source: How we vibe code at a FAANG
As you can see, using AI at work is not a magic pill that solves all problems. It’s primarily a tool that can both help and slow down even experienced developers. From my personal experience, AI is great for simple, localized tasks where no unique solution is required . But the more complex the context and the larger the task, the more it can slow you down.
Back in 2017, when I was a student, my C++ teacher taught us to write code based only on our knowledge and memory. We didn’t write code on paper or in Notepad; at that time, we used Visual Studio 2015, where IntelliSense wasn’t as developed, and we didn’t know about plugins like ReSharper. You know what was a big discovery for us? Shortcuts. This is when, after typing a keyword (like “for” or “switch”) and pressing “tab” the IDE would write the code for you – not exactly writing the whole thing, but generating a template structure to speed up development. We thought it was really cool, but we soon gave it up. Why? Because we were learning to write code on our own.
You can read books like CLR via C# by Jeffrey Richter or Code Complete by Steve McConnell, but unless you’ve written thousands (or even tens of thousands) of lines of code, you won’t learn how to program properly or will do it poorly. During our C++ studies, we wrote linked lists, binary trees, worked a lot with memory, and even created games like Tic-Tac-Toe and Fifteen in the Windows console. And you know how great it feels when all the code is written by you, without any help or AI?
According to the annual StackOverflow survey, about 70% of people aged 18-24, who are just learning programming, use AI.
In my opinion, AI can help explain material in simple terms, which is a huge plus, and you can talk to it like a mentor if you don’t have one. But you need to write code yourself and spend time understanding why something isn’t working. Only then will the knowledge stick and produce results. Therefore, try to minimize AI’s influence on your learning. Don’t rely completely on the code it provides. It can often give incorrect information and only confuse you, so make sure to consult other sources as well.
Once you learn to write code on your own, work with AI, and understand where to use it best, it will truly boost your productivity.
The appearance of tools like ChatGPT, Claude, Cursor, and others is a big plus for the software development industry. They help reduce routine tasks, speed up, and make development easier for programmers. However, there are still tasks that AI doesn’t handle well. So, first and foremost, you should rely on your knowledge and always keep learning.
Сегодня уже невозможно игнорировать влияние AI на разработку ПО. Вместо того чтобы искать информацию в Google, можно просто спросить ChatGPT, и что ещё круче, можно сделать запрос более конкретным и получить более подробную информацию, а не тратить часы на бесконечный поиск по интернету в надежде собрать ответ из кусочков, как пазл. Честно говоря, я уже не помню, когда последний раз заходил на StackOverflow, а ведь в 2018 году у меня обычно было открыто 10-15 активных вкладок, чтобы найти нужную информацию или помочь другим ответами. Помимо информации, AI может помочь проанализировать большие логи с сервера или найти ошибку в конфиге на 500+ строк, на что у человека ушло бы куда больше времени.
AI также помогает писать техническую документацию. Когда я работал над своей open-source библиотекой, ChatGPT написал 80% XML-документации и помог собрать хороший файл документации для GitHub-репозитория.
Но самое главное, AI умеет писать код. Уже около полугода я плачу $20 в месяц за подписку на Claude Code от Anthropic, потому что он повышает мою продуктивность и позволяет решать рутинные задачи намного быстрее. Кроме того, я использую ChatGPT для быстрого поиска информации и написания технической документации.
По данным StackOverflow, около 51% профессиональных разработчиков ежедневно используют AI-инструменты , это довольно значительная цифра.
Источник: StackOverflow Survey 2025
В последнее время термин «vibe coding» стал довольно популярным. Он описывает новый стиль программирования, при котором код за вас пишет AI. Всё началось с использования AI для решения алгоритмических задач, но дошло до того, что даже люди, не умеющие программировать, пытаются создавать что-то сами. С появлением более мощных моделей и редакторов кода вроде Cursor шум вокруг тренда только растёт. Сегодня по одному prompt AI способен сгенерировать большой объём кода: от вёрстки до бизнес-логики, и объяснить, как и почему это работает.
По данным StackOverflow Survey, только 12-15% разработчиков занимаются vibe coding . Цифры пока небольшие, но направление уже сформировалось и продолжит расти.
Когда я начал работать над своими стартапами, мне часто приходилось писать frontend, и, честно говоря, мне это не особо нравилось, кроме случаев с Angular. Тут Claude мне помогал: создавал базовую вёрстку, включая mobile-responsive дизайн, и связывал её с фреймворками. Дальше я уже вручную дорабатывал компоненты, подгонял внешний вид и делал так, чтобы всё работало.
Риски здесь минимальны. AI не работает с бизнес-логикой, базами данных или платёжными системами. Худшее, что может случиться, AI выдаст не тот результат, и я потеряю пару часов.
Тем не менее, заниматься vibe coding стоит только если вы понимаете, что происходит, и пытаетесь оптимизировать процесс, а не если вы рассчитываете на AI, чтобы он сделал то, в чём вы сами не разбираетесь и что не смогли бы сделать сами.
Исследование, опубликованное этим летом группой AI-исследований METR, задалось вопросом, действительно ли инструменты AI помогают опытным разработчикам быть продуктивнее.
В исследовании METR попросил 16 опытных open-source разработчиков выполнить 246 задач в крупных репозиториях. В половине задач им разрешалось использовать AI-инструменты вроде Cursor Pro, в другой половине, нет.
До начала разработчики думали, что AI поможет им справиться на 24% быстрее. Но результаты удивили: «Использование AI на самом деле замедлило их на 19%» , говорят исследователи.
Что важно, только 56% разработчиков в исследовании имели опыт с Cursor, основным AI-инструментом, который предлагался. Почти все (94%) уже работали с какими-то веб-LLM в своих процессах, но именно Cursor некоторые видели впервые. Исследователи отмечают, что разработчиков обучали работе с Cursor перед исследованием.
Эти результаты ставят под сомнение, что AI-инструменты всегда ускоряют ра

[truncated]
