---
source: "https://srikanth.ch/posts/the-ai-situation/"
hn_url: "https://news.ycombinator.com/item?id=49310755"
title: "The AI Situation in Software Development"
article_title: "The AI Situation — Srikanth"
author: "srikanthdotch"
captured_at: "2026-08-15T14:13:19Z"
capture_tool: "hn-digest"
hn_id: 49310755
score: 1
comments: 0
posted_at: "2026-08-15T14:12:27Z"
tags:
  - hacker-news
  - translated
---

# The AI Situation in Software Development

- HN: [49310755](https://news.ycombinator.com/item?id=49310755)
- Source: [srikanth.ch](https://srikanth.ch/posts/the-ai-situation/)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T14:12:27Z

## Translation

タイトル: ソフトウェア開発における AI の状況
記事のタイトル: AI の状況 — Srikanth

記事本文:
AI の状況 — Srikanth Home ai
プロンプト、コンテキスト ウィンドウ、圧縮、AI の操作についてのランダムな考え。
頭の中にあることを実行してもらいたいが、その方法や何をすべきかはわかっていることもあれば、わかっていないこともあります。選択肢は 3 つあります。詳細に至るまで、具体的なことをすべて伝えます。
あるいは、単に高レベルで何かを行うように指示し、それが理解することを期待します。
あるいは、中間の道を進むこともできます。難しいと思われる重要な部分を説明するのがこの方法だと思います。
サンプルをフィードすることができます。これは処理を高速化する方法ですが、サンプルが目的に近いかどうかによって異なります。
これらはすべて時間がかかります。 AIに与える前に時間を費やすものもあれば、AIに与える後もある。
一般的なパターンは、ユーザー認証の実装など、トレーニング セットで以前に見たことがあるはずであることを考慮すると、LLM にとって簡単に実装できます。あなたが想像したり、伝えたりしている新しい問題は、当然、それにとって難しいものであり、手を握る必要があります。
次に、コンテキスト ウィンドウの問題があります。 3000 ワード、4 ページの詳細で緻密な仕様を与えて、それがすべてに従うことを期待することはできません。コードベースが大きくなると、すべてを詰め込むことができなくなり、それにフィードできる膨大なドキュメントも価値がなくなります。詳細な仕様を記述するだけでなく、パターンを要約したり、大規模なデータセットから結論を引き出したりすることも必要です。たとえば、株式の過去のデータセットなど、膨大な数値データの分析などです。
したがって、コストはあなたの負担になります。プロジェクトの詳細なガイド、その目標と問題点、そしてより重要なことに、あなたが望むものの青写真を書くために、やはり時間を費やす必要があります。
次に、AI モデルのドメイン固有の専門知識があります。適切なものを選択する必要があります。
より大きな問題に取り組むにつれて

AI をアプリケーションに統合すると、特定の問題を解決したり、結論を導き出したり、意思決定をしたりするために、すべてではないにしても、できるだけ多くの有用な情報を AI エージェントに詰め込む圧縮の必要性が生じます。この分野でこれを効果的に行う企業が現れるか、モデル構築者がこの問題をきっぱり解決するだろうと思います。
コードベースが大きくなるにつれて、テスト、ツール (専用またはその他)、およびそのアプローチの改良に関するフィードバック ループが必要です。コードベースの信号対雑音比を改善する方法。
優れた説明者や自然な教師は、AI と関わりやすく、より良い成果を生み出すことが容易だと感じています。
結論：まだ時間を費やす必要があります。実装時期は過ぎました。現在では、事前にシステムを設計し、前提条件を変更し、開発セットアップを調整することに費やされる時間は変わりました。しかし、実際に実装がなくなったわけではありません。まだコードではなく言葉で実装している気がします。

## Original Extract

The AI Situation — Srikanth Home ai
Random thoughts about prompting, context windows, compression, and working with AI.
You want it to do something you have in mind, and you know how/what to do, or sometimes you don’t. Now there are 3 options: you tell it everything down to the details, every specific thing.
Or you just tell it to do something at a high level and expect the thing to understand.
Or you can go the middle way. I feel this is the go-to way, explaining the important parts that you think might be difficult for it.
You can feed it examples. It’s a faster way to do things, but it depends on the example being close to what you want.
All of these are time-consuming. Some you spend time before giving to AI, some after.
A common pattern is easy for LLMs to implement, considering they must have seen it before in their training set, for example implementing user auth. A new problem you’re imagining or telling it is of course hard for it and needs hand-holding.
Then there’s the context window problem. You can’t just give a 3000-word, 4-page detailed dense spec and expect it to follow everything, and the larger the codebase, the less it can pack everything in, nor are the vast documents you can feed it worthwhile. Not only for writing detailed specs ~ you also want it to summarize patterns, draw conclusions from a large dataset, be it something like analyzing vast amounts of numerical data, for example a historical dataset for a stock.
So the cost is on you: you still need to spend the time to write a detailed guide for your project, its goals and its issues, and more importantly the blueprint of the thing you want.
Then there’s domain-specific expertise of AI models. You need to pick and choose the right one.
As you work on bigger problems and as you integrate AI into your applications, a need for compression arises, packing as much useful information as possible, if not all, into your AI agent to solve a particular problem or to draw a conclusion, make a decision or whatever. I think there will be companies in this space that’ll do this effectively, or the model builders will just solve this once and for all.
There must be feedback loops in terms of tests, tooling (purpose-built or otherwise), and refining its approach as the codebase grows large. And ways for improving the signal-to-noise ratio in your codebase.
I feel that great explainers or natural teachers find it easy to engage with AI and produce better outputs.
Bottom line: you still need to spend time. The implementation time is gone. Now the time you spend has shifted to designing the system upfront, changing assumptions, and refining your dev setup. But implementation is not really gone. I feel I am still implementing in words instead of code.
