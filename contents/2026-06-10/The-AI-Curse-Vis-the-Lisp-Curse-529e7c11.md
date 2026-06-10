---
source: "https://blog.djhaskin.com/blog/the-ai-curse/"
hn_url: "https://news.ycombinator.com/item?id=48469536"
title: "The AI Curse (Vis the Lisp Curse)"
article_title: "The AI Curse — Dan's Musings"
author: "djha-skin"
captured_at: "2026-06-10T00:59:45Z"
capture_tool: "hn-digest"
hn_id: 48469536
score: 2
comments: 0
posted_at: "2026-06-10T00:10:06Z"
tags:
  - hacker-news
  - translated
---

# The AI Curse (Vis the Lisp Curse)

- HN: [48469536](https://news.ycombinator.com/item?id=48469536)
- Source: [blog.djhaskin.com](https://blog.djhaskin.com/blog/the-ai-curse/)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T00:10:06Z

## Translation

タイトル: AI の呪い (Lisp の呪いに対して)
記事のタイトル: AI の呪い — ダンの思索

記事本文:
Lisp の呪いとは、簡単に言えば、Lisp が非常に強力で表現力豊かであるため、他の人と共同作業することなく、自分だけでコードを書くことが簡単になってしまうということです。見覚えがあると思われるなら、それは今 AI で起こっていることだからです。
GitHub 上にはプロジェクトが数万、数千ではなく数百のスターがあり、同じものに数十の実装があり、どこにでも放棄されたプロジェクトがありますが、コードを自分で修正するのは非常に簡単なので、とにかくそれらのコードを使用できるので問題ありません。私はバイブコーディングについて話しているのでしょうか、それとも Lisp について話しているのでしょうか? (両方です。答えは両方です。)
これを修正するにはどうすればよいでしょうか?そうですね、まずは私たちの態度から始めましょう。 Common Lisp コミュニティは団結し、政治的妥協による統一を支持して、つぎはぎで乱雑な言語を好みました。これが、ANSI Common Lisp 標準が達成された方法です。まだ Lisp の呪いに苦しんでいますが、コミュニティが同じ認識を持っているように見えます。たとえば、json パーサーの実装はまだ少数ですが、そのうちのいくつかは非常に高品質で広く使用されています。
この点において、Common Lisp の人々は、Scheme の親戚よりもはるかにうまくやっていた。彼らの言語ファミリー全体は、言語の作り方についてより良いアイデアがあるなら、聞いたものに参加するのではなく、新しい言語を作って自分のことをすればよいという前提に基づいて設立されている (これは、Guy Steele がオリジナルの論文で行ったことである)。このような前提は非常に分断されたコミュニティを生み出し、Lisp Curse 自体と合わせて、私の意見では、Scheme がプログラミング世界内のすべての言語コミュニティの中で最も分断されたものであることを意味します。
これはすでにAIで起こっています。どうぞ。 GitHub でローカル RAG MCP サーバー (ベクトル データベース ベースのもの) を検索します。このような単純なユニバーサル c の場合、

10,000 個以上の星が付いた単一の実装が見つかると思われるかもしれません。その代わりに、最大 1,000 未満の実装がいくつか見つかりました。私の兄は地元の RAG をやっていますが、彼は自分で RAG をしただけです。そのうちの 1 つをダウンロードしましたが、実行時にセグメンテーション違反が発生しました。
Lisp の呪いとの類似性は十分に憂慮すべきことですが、AI の使用がコラボレーションに反する傾向も見られます。今 Google 検索をするときは、Gemini と対話する方が簡単です。つまり、実際のブログと対話していないことを意味します。これは、ブロガーがブログを書く動機が少ないことを意味します。上で述べたように、Claude を使用すると、他の人のスロップで動作させる必要がなくなるため、自分のコードを書くだけで簡単になります。 AI とインターネット上のコラボレーションの間には、リアパノフの安定機能があるのではないかと考えさせられます。コラボレーションが増えれば増えるほど、そのコラボレーションによって生成されたトークンを AI が利用できるようになります。 AI が増えれば増えるほど、人々がコラボレーションする動機は減ります。コラボレーションに対するインセンティブが低ければ低いほど、AI トレーニングのためのトークンも少なくなります。 AI トレーニング用のトークンが少なければ少ないほど、AI の使用量が減り、より多くのコラボレーションが可能になります。等々。
コミュニティを優先することが重要なステップであることは明らかですが、それにもかかわらず、AI とのコラボレーションを容易にするツールを作成することが重要になってくると思います。 git クローンを想像してください。これにはベクトル インデックス付きのナレッジ グラフ データも保存されているため、すべての共同作業者がコードと、AI がそのコードを操作するために直接使用できるデータの両方にアクセスできるため、AI を使用する全員が同じ認識を持つことができます。共有された AI コンテキストをさらにバイラルにするためには、こうしたものがさらに必要です。このバイラル性の考え方は、エッセイ「Worse is Better」で議論されています。これは、Lisp Curse の一種のアンチテーゼです。つまり、最もバイラルなものが勝ちます。私たちはね

特定のコンテキスト エンジニアリング環境の共有と改善を完全にスムーズに行うコラボレーション ツールを作成するために開発されました。
それまでは、皆が同じことをしているが一緒ではない AI 主導の分断されたコミュニティ、Not-Invented-Here 症候群の汚水溜め、そして新興企業と既存企業が同様に作成した過給された壁に囲まれた庭園「プラットフォーム」をうまく乗り切って、誰もがすべてを社内に保持したいと思わせる一見単純なビジネス上の理由によって動かされます (繰り返しますが、私が話しているのは 2020 年代の AI 情勢ですか、それとも 2020 年代の AI Lisp 情勢ですか)。 80年代？）。その幸運が必要になることはわかっています。
RSS または電子メールで購読してください。

## Original Extract

The Lisp Curse is, briefly, that Lisp is so powerful and expressive that it makes it easy to write code without collaborating with others and just write it yourself. If that sounds familiar, it's because that's what's happening with AI right now.
Hundreds of stars instead of tens of thousands or thousands on GitHub for projects, dozens of implementations for the same thing, abandoned projects everywhere but that's okay because we can just use the code from them anyway because it's really easy to fix the code ourselves. Am I talking about vibe coding, or Lisp? (Both. The answer is both.)
How to fix this? Well, we can start by our attitude. The Common Lisp community came together and preferred a patchy, messier language in favor of unity through political compromise. That is how the ANSI Common Lisp standard was achieved. It still suffers from the Lisp Curse, but there is still some semblance that the community is on the same page. There are still a handful of implementations for json parsers for example, but a couple of them are very high quality and widely used.
In this regard, the Common Lisp folks fared much better than their Scheme cousins, whose whole language family was founded on the premise that if you have a better idea for how to make the language you can just make a new language and do your own thing instead of joining the heard (this is what Guy Steele did with his original paper). Such a premise makes for an extremely fractured community, and together with the Lisp Curse itself, means that in my opinion, Scheme is the most fractured of all language communities within the programming universe.
This is already happening with AI. Go ahead. Search on GitHub for a local-RAG MCP server (one that's vector database based). For such a simple universal concept you'd think we'd find a single implementation with over 10,000 stars. Instead, we find a handful of implementations with less than 1,000, max. My brother does local RAG, but he just rolled his own. I downloaded one of them, but got a segfault when it ran.
This similarity to the Lisp Curse is alarming enough, but we also see the tendency for the use of AI to run counter to collaboration. When I Google Search now, it's just easier to interact with Gemini, which means I'm not interacting with the actual blogs, which means the bloggers have less incentive to write the blogs. When I use Claude, as above mentioned, it's easier to just write my own code because then I don't have to get it to work with somebody else's slop. It makes me think there's a Liapanov stability function between AI and collaboration on the internet. The more collaboration, the more AI can feed on the tokens generated by that collaboration. The more AI, the less people are incentivized to collaborate. The less they are incenitivized to collaborate, the less tokens for AI training. The less tokens for AI training, the less AI we use and the more we collaborate. And so on.
Clearly, prioritizing community is an important step, but I think it's going to become important to make tools that make it easy to collaborate with AI rather than in spite of it. Imagine a git clone but it also stores vector indexed knowledge graph data so that every collaborator has access both to the code and to data that AI can directly use to work on that code, so that everyone that uses AI is all on the same page. We need more of those things to make shared AI context more viral. This idea of virality is discussed in the essay Worse is Better which is sort of the antithesis of the Lisp Curse: the most viral thing wins. We need to make collaboration tools that make sharing and improving a specific context engineering environment absolutely frictionless.
Until then, good luck navigating an AI-driven, fractured community that all does the same stuff but not together, a cesspool of Not-Invented-Here Syndrome and supercharged walled garden "Platforms" written by start-ups and has-beens alike, driven by seemingly straightforward business reasons that makes everyone want to keep everything in-house (again, am I talking about the AI landscape in the 2020s or the AI Lisp landscape in the 80s?). I know I'm going to need that luck.
Subscribe via RSS / via Email .
