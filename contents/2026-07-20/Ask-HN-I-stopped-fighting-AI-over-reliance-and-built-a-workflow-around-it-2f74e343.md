---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48979474"
title: "Ask HN: I stopped fighting AI over-reliance and built a workflow around it"
article_title: ""
author: "dimonb19a"
captured_at: "2026-07-20T15:02:07Z"
capture_tool: "hn-digest"
hn_id: 48979474
score: 2
comments: 1
posted_at: "2026-07-20T14:35:11Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: I stopped fighting AI over-reliance and built a workflow around it

- HN: [48979474](https://news.ycombinator.com/item?id=48979474)
- Score: 2
- Comments: 1
- Posted: 2026-07-20T14:35:11Z

## Translation

タイトル: HN に聞く: AI への過度の依存との戦いをやめ、それを中心としたワークフローを構築しました
HN text: 私はフロントエンドエンジニアで、AIを積極的に活用しています。私はスタイルを微調整する必要がある場合にのみ、手動でコードを書くことは決してないことに気づきました。 AI以前の時代と比べると怠け者で愚かな気がします。コードに 2 分ほどかかる小さな変更を記述する必要がある場合でも、その 2 分をかけて AI にすべてを詳細に説明し、何をどのように変更する必要があるかを指摘する方が簡単です。プロセスについて考える必要がないので、アイデアと望ましい解決策のみを説明し、完了するまで待って確認します。それはある種の劣化です。しかし、重要なのはその品質が素晴らしいということです。生産性とスピードが大幅に向上しました。しかし、イライラすることがあります。怠惰はレビュープロセスにも影響を及ぼし始めました。すべての変更を読んでチェックするのに時間を費やしたくないような気がします。別の AI にそれを依頼するほうがよいでしょう。そしてそれはうまくいきます！私は作業の 95% に Claude を使用し、その後 GPT に切り替えてすべてを適切にレビューします。それぞれ個性が違うので、視点も違います。私は GPT をオタクだと考えています。GPT はクロードが見逃している多くのギャップや矛盾を見つけます。だから私は両方を同じ部屋に閉じて議論するのが好きです。私が本当に大切にしているのは、包括的なドキュメントです。私は AI を使用した計画とブレインストーミングに膨大な時間を費やし、すべてを適切に文書化します。アーキテクチャ、作曲レシピ、チートシートなどに関するすべての MD ファイルは、コードの奥深くに埋められるのではなく、表面に存在する必要があります。すべてのコールド セッションはすべてのドキュメントから始まります。そこでは、AI が流し読みしたり、怠惰な回答をしたりするのを防ぐために、すべてが参考資料と例とともに適切に説明されています。大きな機能を実装する必要があるときは、すぐに実行できる適切なプラを構築します。

n そのため、新しいコンテキストを持つ次のすべてのエージェントがそれを読み取り、何が行われたかを分析し、このセッションのタスクが何かを把握し、すべての変更を進めて報告します。何か大きなことに取り組んでいる数十のチャットが連続して行われることもあります。最終的には、ワークフローを実行してすべての実装をレビューし、GPT に変更を監査してもらい、AI がすべてが良好であると承認した場合にのみ自分自身を確認します。そうすれば、今のやり方が気に入らないので、何を、なぜやり直さなければならないのかがわかります。私のセットアップでは、AI のドキュメントがコードよりも重要です (すべてのレベルの命令ファイル、ファイル タイプごとにロードされるルール ファイル、ロールを持つサブエージェント、すべての編集に関するフック)。最も重要な部分: - サイレント障害のカタログ。 36 個の番号付きトラップ。コンパイルは問題ありませんが、レンダリングは壊れています。すべてのエントリは、原因、症状、修正など、少なくとも 1 回は出荷された実際の障害です。 - 「磨き」ルール。 AI に「もっと良くして」と指示すると、反射的に装飾的な要素を追加します。つまり、磨きとは引き算を意味するという明文化されたルールがあり、実際のレビューで拒否されたパターンのリストが含まれています。 - ドキュメントがコードから逸脱すると失敗する CI。コンポーネント レジストリはソース ファイルと照合してチェックされます。嘘をつくドキュメントは、ドキュメントを作成しないよりも悪いからです。 - ランタイム検証ツール (WebMCP)。エージェントは実際の Chrome ブラウザを操作し、ページをクリックして、コンソールを監視し、3 つのビューポートでスクリーンショットを撮ります。基本的に私はコードを書くのをやめて、コードを書く環境をエンジニアリングし始めました。すべてのファイルが存在するのは、AI が一度何かで失敗したためであり、それが二度と起こらないように私が書き留めたものです。私にとって最大の問題はモチベーションです。行き詰まったときのあの感覚を覚えています。何時間も何かを理解できず、正しい解決策を探し、試し、そしてやっとうまくいった、という満足感がありました。何も持っていない

もっと。全く専門知識のない分野でも自信を持って取り組めます。つまり、AI はそれを処理できるのです。私のシステムは品質の問題を解決していると思いますが、怠惰/モチベーションの問題は増大します。怠けても大丈夫なようにしました。質問は 2 つあります: 1. AI のセットアップはどのような感じですか?プロンプトの方法だけではなく、実際のファイルとルールを意味します。 2. 動機の部分を実際に解決した人はいますか?

## Original Extract

I am a frontend engineer and actively use AI. I caught myself that i never write code manually, only if i have to tweak some styling. I feel lazy and dumb compared to pre-AI era. Even if i have to write some small changes in code that take like 2 minutes - it is easier to spend those 2 minutes explaining everything in details to AI and point out what and how it needs to change. Because i dont have to think about the process, i only explain the idea and desired solution - then wait until its done to review. Thats some kind of degradation. But the thing is that the quality is amazing. My productivity and speed increased significantly. But there is a frustrating thing. The laziness started affect even review process. I feel like i dont want to spend time to read and check all the changes - i will better ask another AI to do it. And it works! I use Claude for like 95% of work, then switch to GPT to review everything properly. They have different personalities so you have different points of view. I consider GPT as a nerd - it catches a lot of gaps and inconsistencies that Claude misses. So i like to close them both in one room to debate. What i really do care a lot is a comprehensive documentation. I spend a huge amount of time for planning and brainstorming with AI, then i document everything properly. All those MD files about architecture, composition recipes, cheat sheet… All of it should lie on surface, not to be buried somewhere deep in code. Every cold session starts with all the docs where everything is explained properly with references and examples to prevent AI to skim or give lazy answers. When i need to implement some big feature i structure a proper ready-to-execute plan so every next agent with fresh context reads it - analyzes what was done - figures out what is the task for this session - and goes ahead with all the changes, then report. It can be dozens of chats working sequentially on some big thing; in the end i will run some workflow to review all the implementation, then ask GPT to audit changes, and i will look myself only when AI approves that everything looks great. Then i can tell what and why has to be redone because i dont like how we have it now. In my setup the documentation for AI outweighs the code (instruction files on every level, rule files that load by file type, subagents with roles, hooks around every edit). The most weightful part: - A catalog of silent failures. 36 numbered traps that compile fine but render broken. Every entry is a real failure that shipped at least once: cause, symptom, fix. - A "polish" rule. When you tell AI "make it better", its reflex is to ADD decorative stuff. So there is a written rule that polish means subtraction, with a list of patterns i rejected in real reviews. - CI that fails when docs drift from code. The component registry is checked against source files. Because docs that lie are worse than no docs. - A runtime verifier (WebMCP). An agent drives a real Chrome browser, clicks through the page, watches the console, takes screenshots at 3 viewports. Basically I stopped writing code and started engineering the environment that writes the code. Every file exists because AI failed at something once, and i wrote it down so it wont happen again. The biggest problem for me is motivation. I remember that feeling how you got stuck - cant figure out something for hours, searching right solution, trying, and then FUCK YEAH FINALLY, and the satisfaction: YOU NAILED IT. I dont have it anymore. I feel confident even in things where i have no expertise at all. Thats like: meh AI can handle it. I think my system solves the quality question, but increases the laziness/motivation problem. I made it safe to be lazy. So two questions: 1. How does your AI setup look like? I mean real files and rules, not just how you prompt. 2. Has anyone actually solved the motivation part?

