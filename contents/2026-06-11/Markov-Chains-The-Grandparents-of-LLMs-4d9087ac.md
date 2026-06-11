---
source: "https://dmanco.dev/2026/06/10/the-grandparents-of-llms.html"
hn_url: "https://news.ycombinator.com/item?id=48488011"
title: "Markov Chains: The Grandparents of LLMs"
article_title: "Markov Chains: the grandparents of LLMs · Doch"
author: "Doch88"
captured_at: "2026-06-11T10:36:31Z"
capture_tool: "hn-digest"
hn_id: 48488011
score: 1
comments: 0
posted_at: "2026-06-11T09:08:09Z"
tags:
  - hacker-news
  - translated
---

# Markov Chains: The Grandparents of LLMs

- HN: [48488011](https://news.ycombinator.com/item?id=48488011)
- Source: [dmanco.dev](https://dmanco.dev/2026/06/10/the-grandparents-of-llms.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T09:08:09Z

## Translation

タイトル: マルコフ連鎖: LLM の祖父母
記事のタイトル: マルコフ連鎖: LLM の祖父母 · Doch

記事本文:
マルコフ チェーン: LLM の祖父母 · Doch
ドック
ランダムな人からのランダムなもの。技術的な話題からカジュアルな話題まで。
マルコフ連鎖: LLM の祖父母
ここ数年ソーシャル メディアを利用している人は、ChatGPT がリリースされた 2022 年に AI が始まったと人々が信じているように感じます。
多くの人に知られていない会社である OpenAI が、驚くほどスマートに感じられるこの画期的なサービスを突然リリースしました。
しかし、その前に、現在のレベルの AI 研究につながる多くの出来事が起こりました。ここでは、非常に似たようなことを行うことができた非常に古い手法の 1 つであるマルコフ連鎖についてお話します。
現在の LLM は、この投稿で説明する内容の超高度なバージョンにすぎないと言っても、それほど間違いではありません。
他のすべての定義はここから得られるため、プロセスとは何かを定義することから始めましょう。
アクター、いくつかの状態、いくつかの遷移という 3 つの要素が必要です。
たとえば、人 (俳優) が日常生活を送っていると想像してください。家、オフィス、コーヒーショップなど、さまざまな場所が州を表します。時間の経過とともにある場所から別の場所に移動する行為は、移行を表します。
これら 3 つの要素を定義することで、一連のイベントを計画することができます。この基本的な概念により、時間の経過とともに行動を追跡するためのフレームワークが得られます。これをマルコフ過程にするには、もう 1 つのプロパティ、memorylessness が必要です。
プロセス内の特定の状態は前の状態のみに依存し、以前に何が起こったかについての記憶はありません。
これを言葉に置き換えて、例としてフレーズを選んでみましょう。「猫はマットの上に座り、犬は敷物の上に座りました。」
すべての単語が状態であり、遷移とは、ある単語から次の単語に移動する行為です。
このテキストが私たちのすべてを表していると仮定すると、

語彙の場合、単語が互いに続く頻度を数えるだけで、ある単語から別の単語に移動する確率を計算できます。たとえば、現在の状態が「the」である場合、次の状態は「cat」、「mat」、「dog」、または「rug」になる可能性があります。
すべての単語が前の単語に接続されるチェーンを作成します。すべての単語の確率は、その先行単語にのみ依存します。
マルコフ連鎖が現在「sat」という単語にある場合、確率を調べて、100% の確率で次の単語が「on」であることがわかります。
しかし、「オン」になるとどうなるでしょうか?次の単語は、「the」（「sat on…」に続く）または「the」（他の「sat on…」に続く）である可能性があります。
それは常に「the」につながります。私たちが「その」に戻ると、チェーンは猫から始めたのか犬から始めたのかを完全に忘れてしまいます。
次の単語が猫、マット、犬、または敷物である確率は等しいです。 4 面のサイコロを振って次の単語をランダムに選ぶだけで、同じマルコフ プロセスに従うことができます。
これはLLMとどのように似ていますか?
たくさんの本を使ってこのマルコフ連鎖を作成し、その連鎖を自己回帰ジェネレーターとして使用して独自の疑似論理文を作成することを想像してください。
マルコフ連鎖は実際のテキストで作成されているため、一見正しい文章のように見えますが、正しく読もうとするとあまり意味がわかりません。
確率のみに基づいてジェネレーターをトレーニングしましたが、その状態にコンテキストは一切組み込まれません。
1 つの単語ではなく 2 つの単語を調べるように状態を拡張することを考えることもできます (ユニグラムからバイグラムへの移行)。これにより、文章が少し論理的になったように感じられますが、それでもスマートとは程遠いです。
本当に一貫した文章を得るには、本能的に状態を拡張し続けようとするかもしれません。

最後の10単語か20単語を見てみませんか？
これがマルコフ連鎖の問題です。長い単語シーケンスの完全一致が必要な場合、その正確な組み合わせがトレーニング データ内の他の場所で繰り返されることはほとんどありません。
可能なパスが 1 つしかないため、チェーンがスタックするか、単に元のテキストをそのままコピーするかのどちらかです。
しかし、もしかしたら、それらの単語を別の次元、つまり、実際に使用される単語よりも意味に基づいて、特定の単語のシーケンス間のつながりがより抽象的な次元にマッピングできるかもしれません。
これは、多かれ少なかれ、LLM が行うことです。これらは、単語を連続的な数学的空間にマッピングし、 埋め込み となり、厳密な単語ではなく抽象的でマッピングされた概念に基づいた確率ベースの遷移をたどります。
残る唯一の問題は、記憶喪失の制限です。あるコンセプトの確率が以前のコンセプトだけに依存するとしても、非常にスマートなマシンを作ることはできません。
しかし、単純なマルコフ連鎖の例のように、単一の状態を定義するためにさらに多くの概念を使用できるかもしれません。
最後の概念だけを見るのではなく、これまでに生成した一連の抽象概念全体として現在の状態を定義できます。これらすべての過去の概念に異なる重みを与えることで、次の遷移の確率に異なる寄与を与えることができます。
これは、現代の LLM の基礎である論文「Attending Is All You Need」(2017 年、Vaswani et al.) に記載されている有名な注意メカニズムです。
アテンション メカニズムを使用してテキスト全体を重み付けすることにより、LLM は従来のマルコフ チェーンの記憶のない制限から効果的に解放されます。
LLM は、その祖父母と同様に、依然として確率モデルであり、抽象的な概念を意味のある方法で結び付ける非常に複雑なマルコフ連鎖です。
これらの概念をより適切にモデル化すればするほど、LLM の機能が向上します。

賢く見えるかもしれませんが、確率的な性質が変わるわけではないので、ある時点で犬ではなく猫が敷物の上にいることに気づいても驚かないでください。
マルコフ連鎖がどのように機能するかを試してみませんか? WhatsApp グループ チャットでモデルをトレーニングしたいですか?
友達グループの WhatsApp チャットを .txt としてエクスポートし、WhatsApp マルコフ ジェネレーターに移動して、新しいフレーズを生成します。
圧縮アーチファクトによる AI の偽物の検出
2025 年 9 月 15 日
AI エンジニアは AI に取って代わられることから安全ではない
2025 年 8 月 17 日

## Original Extract

Markov Chains: the grandparents of LLMs · Doch
Doch
Random stuff from a random person. Ranging from tech topics to more casual ones.
Markov Chains: the grandparents of LLMs
If you have been on social media in the past few years, it feels like people believe that AI started in 2022 when ChatGPT was released.
Out of a sudden, OpenAI, a company unknown to many, released this groundbreaking service that felt surprisingly smart.
But before that, many things happened that led to the current level of AI research, and here I will talk about one very old method that managed to do something very similar: Markov chains.
It’s not too wrong to say that current LLMs are just super advanced versions of what I will talk about in this post.
Let’s start by defining what a process is, since all the other definitions will come from this.
We need three ingredients: an actor, some states and some transitions.
For example, imagine a person (the actor) going about their daily routine. The different locations they can be in, such as their house, the office, or a coffee shop, represent the states. The act of moving from one location to another as time passes represents the transitions.
By defining these three elements, we can map out a sequence of events. This foundational concept gives us a framework to track behaviour over time. To make it a Markov process, we just need another property: memorylessness .
A specific state in the process depends solely on the previous state, and it has no memory of what happened before.
Let’s transfer it to words, and pick a phrase as an example: “The cat sat on the mat, and the dog sat on the rug.”
Every single word is a state, and the transition is the act of moving from one word to the next.
If we assume this text represents our entire vocabulary, we can calculate the probabilities of moving from one word to another just by counting how often they follow each other. For instance, if the current state is “the”, the next state could be “cat”, “mat”, “dog”, or “rug”.
We create a chain where every word is connected to the previous one. The probability of every word solely depends on its predecessor.
If our Markov chain is currently at the word “sat”, it looks at the probabilities and sees that 100% of the time, the next word is “on”.
But what happens when it reaches “on”? The next word could be “the” (following “sat on…”) or “the” (following the other “sat on…”).
It always leads to “the”. Once we are back at “the”, the chain completely forgets whether we started with the cat or the dog.
The probability of the next word being either cat, mat, dog or rug is equal; we could just roll a four-sided die and pick the next word at random, and we would still follow the same Markov process.
How is this similar to an LLM?
Imagine creating this Markov chain using a lot of books, and then using the chain as an autoregressive generator to make up your own pseudo-logical sentences.
Since the Markov chain was created with real texts, the sentences would probably look correct at first sight, but they do not make too much sense when you try to read them properly.
You have trained a generator that is solely based on probability, and it doesn’t bring any context into its state.
You could think about expanding the state to look at two words instead of one (moving from unigrams to bigrams). This makes the text feel a bit more logical, but it’s still far from smart.
To get truly coherent sentences, your instinct might be to just keep expanding the state: why not look at the last 10 or 20 words?
That’s the problem with Markov chains: if you require an exact match for a long sequence of words, you will rarely find that exact combination repeated anywhere else in your training data.
The chain either gets stuck or it simply copies the original text as it is, because there is only one possible path.
But maybe we can map those words in a different dimension, a dimension where the link between certain sequences of words is more abstract, based more on the meaning than on the actual words used.
This is more or less what LLMs do. They map words into a continuous mathematical space, becoming embeddings , following a probability-based transition based on abstract, mapped concepts rather than rigid words.
The only remaining problem is the memorylessness limit. If the probability of a concept depends solely on the previous one, we still can’t have a very smart machine.
But, like in the simple Markov chain example, maybe we can use more concepts to define a single state.
Instead of looking at just the last concept only, we can define the current state as the entire sequence of abstract concepts we have generated so far. We can give a different weight to all those past concepts so they contribute differently to the probability of the next transition.
That is the famous attention mechanism from the paper “Attention Is All You Need” (2017, Vaswani et al.), the foundation of modern LLMs.
By using the attention mechanism to weigh the entire text, LLMs effectively break free from the memoryless limits of traditional Markov chains.
LLMs, like their grandparents, are still probabilistic models, very complex Markov chains that link abstract concepts together in a way that makes sense.
The better you model those concepts, the more the LLM will appear smart, but that doesn’t change its probabilistic nature, so don’t be surprised if, at some point, you’ll find your cat on the rug instead of your dog.
Do you want to try how a Markov chain works? You want to train a model on your WhatsApp group chat?
Export the WhatsApp chat of your friend group as .txt, go to WhatsApp Markov Generator and generate new phrases!
Detecting AI Fakes with Compression Artifacts
15 Sep 2025
AI Engineers aren’t safe from being replaced by AI
17 Aug 2025
