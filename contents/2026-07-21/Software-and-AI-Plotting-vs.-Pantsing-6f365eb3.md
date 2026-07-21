---
source: "https://pyjarrett.github.io/2026/07/18/software-and-ai-plotting-versus-pantsing.html"
hn_url: "https://news.ycombinator.com/item?id=48990854"
title: "Software and AI – Plotting vs. Pantsing"
article_title: "Software and AI - Plotting vs. Pantsing"
author: "pyjarrett"
captured_at: "2026-07-21T12:20:54Z"
capture_tool: "hn-digest"
hn_id: 48990854
score: 1
comments: 0
posted_at: "2026-07-21T11:30:30Z"
tags:
  - hacker-news
  - translated
---

# Software and AI – Plotting vs. Pantsing

- HN: [48990854](https://news.ycombinator.com/item?id=48990854)
- Source: [pyjarrett.github.io](https://pyjarrett.github.io/2026/07/18/software-and-ai-plotting-versus-pantsing.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T11:30:30Z

## Translation

タイトル: ソフトウェアと AI – プロットとパンツ
記事のタイトル: ソフトウェアと AI - プロットとパンツ
説明: AI が好きな人、嫌いな人がいる理由についての別の解釈

記事本文:
ソフトウェアと AI - プロットとパンツ処理
pyjのブログ
投稿
タグ
ソフトウェアと AI - プロットとパンツ処理
ソフトウェア「プロット」と「パンツ」
私は、なぜ人々が自分の仕事のために AI を好むように見えるのに、他の人は AI を嫌うのか疑問に思っています。私の好奇心は、社会、経済、環境への懸念はさておき、仕事自体における AI の使用に集中しています。溝の原因の一部は、ソフトウェアと問題領域についての理解を構築する方法に AI がどのような影響を与えるかにある可能性があります。
フィクションの執筆には、ソフトウェアと同様に、多くの詳細を一貫した方法で管理する必要があります。フィクションの執筆には 2 つのスタイルがあります。1 つは事前にアイデアを練るプロッター (「プランナー」、「アーキテクト」) で、もう 1 つは「ズボンの裾」で書くパンツター (「庭師」) です。作成者はこれらの両方を異なるタイミングで使用しますが、どちらか一方を好む場合があります。人々が 1 つのカテゴリーにのみ当てはまるという誤解を避けるために、これらのアプローチの動作を「プロット」と「パンツ」と呼びます。
フィクションの執筆におけるプロットとは、世界を構築し、キャラクターガイドを書き、書き始める前に主要なストーリーのビートとアークの概要を説明することです。パンツィングは、いくつかの興味深いキャラクターと設定を取り上げ、物語がどこに行き着くのかを考えながら書き始めます。
これら 2 つのアプローチは、いくつかのソフトウェアの実践とアプローチを反映しているようです。私の目標は、どちらのスタイルを推進したり否定したりすることではなく、ソフトウェア エンジニアリングとそれが AI の使用にどのような影響を与えるかを考えるためのフレームワークとしてこれを使用することです。
ソフトウェア「プロット」と「パンツ」
「プロット」とは、コードを記述する前にソリューションを徹底的に考えて開発することです。 「パンツ」では、コードを主要な成果物として扱い、それを使用して解決策を検討します。
完璧ではありませんが、プロットすると c のトップダウン分析形式が強調されます。

パンツはボトムアップ合成による作成に重点を置いています。
違いは、コードがどのように扱われるか、そしてどこに理解が存在するかです。プロットではコードはデザインの表現ですが、パンツではそれがデザインです。
私のような一部の人々にとって、コードを書くことは問題を理解するのに役立ちます。 AI にすべてのコードを書かせるように誰かに指示することは、「私たちの仕事を奪う」のではなく、私たちの理解を奪うことになります。 AI が自分たちの仕事の楽しみを奪っていると人々が話すとき、これが理由である可能性があり、これが AI 導入における意見の分かれ目の一部を説明することになるでしょう。
AIの利用のしやすさや有用性が異なる理由はこれにあるようだ。 AI の使用に関する用語は非常に曖昧なように思われるため、これらのシステムに対して包括的な用語「AI」を使用しながら、私が何を意味するのかを定義します。この区別は、「エージェント」とインライン編集システムまたはチャット システムの間で明確に分かれているわけではありません。
プロットとパンツの場合、その分かれ目は、AI がアーティファクトを更新する際の自律性と粒度になります。 「デリゲート」は、目標が与えられ、作業成果物の編集を自律的に実行できる AI になります。 「アシスタント」には、人間が制御を維持する残りの作業が含まれます。
これは、代表者がプロットをサポートし、アシスタントがパンツをサポートしていることを示唆しています。デリゲートは、高レベルの仕様から出荷された製品に直接進みます。アシスタントは、オートコンプリート、調査、コード レビューなどのために人間と共同作業します。これらがエージェントで行われるかチャット インターフェイスで行われるかは重要ではなく、最終製品がどのように使用されるかが重要です。
定義が設定されている場合、これはどのように適用されるのでしょうか?
当初、プロットとパンツは特定の分野にアピールするのではないかと考えていましたが、さまざまな種類の作業やシステムの知識レベルに適合するようです。
多くのソフトウェア開発者は、次のようなプロジェクトに取り組んでいます。

自分たちは書いていない。 main() を書くプログラマはほとんどおらず、開発者のプロジェクト在職期間は非常に短い傾向があります。オリジナル製品の作成者が残っていないチームを見つけることも珍しいことではありませんが、ソフトウェアのビジネス ケースは残っていませんでした。この環境では、ソフトウェアがどのように動作するかを知る唯一の方法は、コードとテストを読むことかもしれません。これは簡単に聞こえますが、進化の発展を生き延びた成熟したシステムには、痕跡的な制約や要件があることがよくあります。古いコードを扱うときは、現在のシステム状態を知る必要があります。この状況は、数百万行のコードや、異なるチームが作成したマイクロサービスなど、相互通信する異種システムを扱う場合にはさらに悪化します。
私は既存のコードベースの新人として何度も活動してきました。人間の記憶や逸脱により、設計上の決定や望ましいアーキテクチャが時間の経過とともに失われ、メンテナンスされていないコードが生成される部分がよくあります。意味のない奇妙な略語ですか？そう、その意味は時間とともに失われていくのです。非常に大規模なプログラム (コードが 100 万行を超える) の場合、システム全体を完全に理解し続けることは不可能です。ソフトウェア エンジニアを連れてこのシステムにパラシュートで降ろすと、彼らはゼロからスタートし、ほとんどパンツモードで実行されます。システムは未知の山です。
プロジェクトが正しく実行されることを確認するにはどのようなチェックが役立ちますか?
誰がこのコードに依存しているのでしょうか?また、何かを壊した場合はどうすればわかりますか?
自分が割り当てられているシステムを理解したら、通常は簡単な方法で変更をプロットできます。一部のシステムには強力な規則があり、標準のデータ フローやワークフローを強制するのに役立ちます。ただし、これは、専門用語を理解する必要があるオーダーメイドのソフトウェアや社内の専門知識の場合には当てはまりません。調子に乗るときは、l から始まることがよくあります

ワークフローに慣れるために少しずつ変更してください。次に、ワークフローの理解が深まるにつれて、行うことも拡張します。さまざまな作業に取り組むにつれて、知識が保持され、チーム全体に広がります。
既存のシステムでは、制約により、状況に慣れるまではじっとしていなければならないことがよくあります。新しい貢献者も、追加の制約を発見したり、遭遇した不明瞭な設計要素を体系化することで製品をより鮮明にするのに役立ちます。 AI のデリゲートが仕事に使用されると、システムについての共通の理解がほとんど、またはまったく構築されません。
「AI は開発者の効率を向上させていない」と聞くと、私の頭はそこに行きます。プログラムの現在の状態や現在の問題自体がわからない場合、作業を委任するのは困難です。 「デリゲート」CLI ベースの AI を使用して仕様から出力に移行しても、根本的な問題の解決には役立ちません。迷ったとき、委任のボトルネックとなるのは理解することです。計画モードに移行し、操作するシステムを理解するために何度も繰り返すことができますが、これは新しい変更と既存のコードの間のインターフェイスの境界についてのみ情報を提供します。
既存のコードでは、「アシスタント」AI が PR がコード規則に従っていることを保証します。オートコンプリートは、適切な内部型と言語を使用することを推奨します。アシスタントは、問題に取り組むために参照すべきプログラム領域を見つけるのにも役立ちますが、これは理解を強化するものであり、それに代わるものではありません。
グリーンフィールド プロジェクトの開発では、まったく異なるエクスペリエンスが提供されます。製品のニーズがまだ完成していない可能性があるため、安定性は機能に劣ります。問題の理解を体系化することは、価値を実証することよりも重要であることがわかります。
限られた滑走路では速度が必要であり、これは作業の細分化を意味します。ほとんどの問題はまだ解決されておらず、同様に必要な内部制約も少なくなります。

維持する必要があります。システムが使用不能になったり、製品が方向転換したりした場合、システム全体が廃棄される可能性があります。コードは機能を提供するため、グリーン フィールドでは使い捨ての資産です。コードは制約があるため、茶色い領域では負債となります。
グリーンフィールド プロジェクトは、時間が経過して固まるにつれて、ゆっくりとブラウンフィールド プロジェクトに変わります。知識は、不完全なフィルターを通して、精神的なものからゆっくりとコードに変換されます。ただし、直接表現できないものもあり、それが損失につながります。
AI デリゲートがグリーンフィールド コードを作成する場合、競合の可能性が少ないため、リスクが低くなります。他の人が書いたコードは少なく、作成者はなぜそのようなことが行われるのか、プロジェクトの歴史を知っています。製品について学ぶための準備をスキップし、コードをほとんどまたはまったく使わずに、すぐにプロットに進むことができます。
問題をコンピュータ上で実行するにはコード内で何らかの表現が必要であり、無数のプログラムが同じ出力を生成する可能性があります。問題の解決方法を知っていることと、それを迅速に表現したり、正しさの間違いを防ぐ方法で表現したりすることの間には、大きな違いがあります。なじみのない問題に対処するとき、私と同じように、その問題に対する解決策をさまざまな方法で表現することを繰り返す人もいます。これらの中には、「Parse, don't validate」で説明されているアイデアのように、賢くて便利なものもあります。これは、問題について学ぶためにコードを使用するという、別の種類のパンツです。これは一種の遊びであり、プログラミングや経験を積むことの喜びの源です。 AI がプログラミングの楽しさを奪っていると人々が言うのを聞くとき、この「学ぶ遊び」の除去が私にとってそれを意味します。このプロセスは製品だけでなく知識も生み出します。
AI を使用してソリューションを生成すると、この種の作業の理論的根拠の一部が失われます。インラインコード生成と調査により、より迅速に作業を進めることができます

障害物を減らすことによって。ただし、コードレビューのアシスタントはバックアップと追加の洞察を提供します。
システムの知識は、誰かの頭の中、仕様書、コードの中など、どこかに存在します。プロットは、既存のメンタル モデルを世界に最大限にプッシュします。パンツィングは世界からメンタル モデルを抽出し、経験を構築します。
重要な問題は、AI の使用がソフトウェア エンジニアの考え方にどのような影響を与えるかです。これらのスタイルは、さまざまな理由により、さまざまなタイミングで連携して機能します。デリゲート AI は既存の理解をコードに変換します。アシスタント AI は、ドキュメントとコードを理解に変換するのに役立ちます。
間違った AI スタイルを使用すると、実行中の作業が著しく損なわれる可能性があります。重要な要素は、現在行っている開発の種類をサポートするために適切な形式の AI を使用することのようです。 AI に特定のスタイルを強制するのは、特にそれが自分より上の人からのものである場合、気まずいものになります。
あなたは今、計画を立てているのか、あるいは喘いでいるのでしょうか。AI はそれをどのようにサポートしているのでしょうか?

## Original Extract

An alt take on why some like or dislike AI

Software and AI - Plotting vs. Pantsing
pyj's blog
Posts
Tags
Software and AI - Plotting vs. Pantsing
Software "Plotting" and "Pantsing"
I've been wondering why people seem to love AI for their work, while others hate it. My curiosity revolves around usage of AI in the work itself, setting aside social, economic and environmental concerns. The source of the divide might partly be how AI affects the way we build understanding of software and problem domains.
Fiction writing, like software, requires managing many details in a coherent way. There are two styles of fiction writing: plotters ("planners", "architects") who work through their ideas ahead of time, and pantsers ("gardeners") who write by the "seat of their pants." Authors use both of these at different times but might favor one or the other. To avoid the misinterpretation that people only fit into one category or another, I'll refer to the behavior of these approaches as "plotting" and "pantsing."
Plotting in fiction writing is worldbuilding, writing character guides, and outlining the major story beats and arc before starting to write. Pantsing would be taking a few interesting characters and a setting, and just starting to write, figuring where the story ends up as you go.
These two approaches seem to mirror some software practices and approaches. My goal isn't to promote or denigrate either style, but rather to use this as a framework to think about software engineering and how it impacts AI usage.
Software "Plotting" and "Pantsing"
"Plotting" would be thinking through and developing a solution before any code is written. "Pantsing" would be treating the code as a main artifact and using that to think through the solution.
It's not perfect, but plotting emphasizes a top-down analytical form of construction, whereas pantsing focuses on creating from bottom-up synthesis.
The difference is how code is treated and where understanding lives. In plotting, code is an expression of the design, whereas in pantsing, it is the design.
For some people like myself, writing the code helps us build our understanding of the problem. Telling someone to have an AI write all the code isn't "taking away our work," it's taking away our understanding. When people talk about AI stealing the joy of their craft, this might be why, which would explain some of the divide in AI adoption.
This seems to explain why AI usage differs in ease and usefulness. AI usage terminology seems quite blurry, so I'll define what I mean while using the umbrella term "AI" for these systems. The distinction doesn't split neatly between "agents" and inline editing or chat systems.
For plotting versus pantsing, the divide would be the autonomy and granularity with which AI updates artifacts. "Delegates" would be an AI which is given a goal and allowed to autonomously perform edits on the work products. "Assistants" would encompass the rest of work in which the human maintains control.
This would suggest that delegates support plotting, while assistants support pantsing. A delegate goes from high level specifications directly to a shipped product. An assistant collaborates with the human, such as for autocomplete, research, or code review. It's not important whether these happen with agents or a chat interface, but how their end products get used.
With definitions set up, how does this apply?
Initially I thought plotting and pantsing might appeal to specific domains, but it seems to align with different types of work and levels of knowledge of the system.
Many software developers work on projects that they themselves didn't write. Few programmers write main() and developers tenure on projects tends to be very short. It's also not uncommon to find a team with no original product authors left, but the business case for the software didn't leave. In this environment, the only way to learn how the software works might be to read the code and tests. This sounds straightforward, but mature systems surviving evolutionary development often have vestigial constraints and requirements. When dealing with older code, you need to learn the current system state. This gets worse when dealing with millions of lines of code and with intercommunicating and disparate systems, such as microservices created by different teams.
I've been a fresh face on existing codebases many times. There are often portions where human memory or departures result in unmaintained code whose design decisions or desired architecture are lost to time. That weird acronym that makes no sense? Yeah, its meaning is lost to time. For extremely large programs (>1 million lines of code), it's impossible to maintain a full understanding of the whole system. If you take a software engineer and parachute them into this system, they're starting from zero and running mostly in pantsing mode. The system is a pile of unknowns.
What checks help ensure the project runs correctly?
Who depends on this code and how will I know if I broke something?
Once you've developed an understanding of the systems on which you're assigned, you can plot changes usually in a straightforward way. Some systems have strong conventions, which helps enforce standard data flows and workflows. However, this isn't true for bespoke software, or in-house expertise, where you need to come up to speed with the vernacular. Coming up to speed often starts with little changes to get familiar with the workflows. Then you expand what you do as your understanding of the workflows expands. Knowledge gets retained and spread through the team as you work on different things.
In existing systems, constraints often necessitate pantsing until you're up to speed. New contributors also discover additional constraints or help sharpen the product by codifying unclear design elements they run across. When AI delegates get used for work, little or no shared understanding of the system gets built.
When I hear "AI isn't improving efficiency for developers," this is where my mind goes. If you don't know the current state of the program, or the current problem itself, it's difficult to delegate work. Using "delegate" CLI-based AI to go from spec to output doesn't help with the underlying problem. When you're lost, the bottleneck to delegation is understanding. You can go into planning mode and then iterate a bunch trying to get an understanding of the systems you touch, but this only informs you on the interface boundary between new changes and existing code.
In existing code, an "assistant" AI ensures a PR follows code conventions. Autocomplete would recommend using the appropriate internal types and vernacular. An assistant can also help find program areas to look at to get started on a problem, but this augments understanding and doesn't replace it.
Developing a greenfield project provides a completely different experience. Stability falls behind features as the product need might not be complete yet. Codifying understanding of the problem proves less important than demonstrating value.
A limited runway necessitates speed, which means a subdivision of work. Most problems aren't solved yet and there are similarly fewer internal constraints needing to be maintained. Entire systems might be tossed out if they prove unusable or the product pivots. Code is an expendable asset in green fields because it provides capability. Code is a liability in brown fields because it constrains.
Greenfield projects slowly convert into brownfield ones as they age and solidify. Knowledge slowly converts from being only mental through an imperfect filter into code. Some things cannot be expressed directly though, which leads to some loss.
An AI delegate writing greenfield code is less risky since there are fewer possible conflicts. There's less code written by others and the authors know why things are being done and the history of the project. You skip the pantsing to learn the product and can jump straight into plotting with little or no code in your way.
Problems require some sort of expression in code to run on a computer, and an infinite number of programs can produce the same output. There's a wild difference between knowing how to solve a problem and expressing it in a way which is fast or prevents correctness mistakes. When dealing with an unfamiliar problem, some people, like myself, iterate on various ways of expressing solutions to that problem. Some of these can be clever and useful, like the ideas described in Parse, don't validate . This is pantsing of a different sort, of using code to learn about the problem. It's a sort of play, a source of joy in programming and how we gain experience. When I hear people describe AI stealing the joy of programming, this removal of "play to learn" is what it means to me. The process produces not just a product, but also knowledge.
Using an AI to generate a solution removes part of the rationale for this type of work. Inline code generation and research lets you move faster though by reducing obstacles. An assistant for code review though provides backup and additional insights.
The knowledge of the system exists somewhere, whether in someone's head, or in a spec or in the code. Plotting maximizes the push of existing mental models into the world. Pantsing extracts mental models from the world and builds experience.
The critical issue is how AI usage affects the way software engineers think. These styles work together at various times for various reasons. A delegate AI converts existing understanding into code. Assistant AIs help convert documentation and code into understanding.
Using the wrong AI style can actively detract from the work being done. The important element seems to be using the right form of AI to support the type of development you're doing right now. Forcing a certain style of AI feels awkward, especially if it comes from someone above you.
Are you plotting or pantsing right now, and how does AI support that?
