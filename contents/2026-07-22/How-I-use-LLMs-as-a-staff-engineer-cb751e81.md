---
source: "https://www.seangoedecke.com/how-i-use-llms/"
hn_url: "https://news.ycombinator.com/item?id=49001739"
title: "How I use LLMs as a staff engineer"
article_title: "How I use LLMs as a staff engineer"
author: "kp25"
captured_at: "2026-07-22T04:59:38Z"
capture_tool: "hn-digest"
hn_id: 49001739
score: 2
comments: 0
posted_at: "2026-07-22T04:04:49Z"
tags:
  - hacker-news
  - translated
---

# How I use LLMs as a staff engineer

- HN: [49001739](https://news.ycombinator.com/item?id=49001739)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/how-i-use-llms/)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T04:04:49Z

## Translation

タイトル: スタッフ エンジニアとして LLM を使用する方法

記事本文:
スタッフ エンジニア Sean Goedecke としての LLM の使用方法
スタッフエンジニアとしての LLM の使用方法
大規模な言語モデルについては、ソフトウェア エンジニアの意見が大きく分かれています。多くの人は、これが業界にこれまで登場した中で最も革新的なテクノロジーであると信じています。また、これらは誇大宣伝だけを目的とした長い製品ラインの最新作であると信じている人もいます。考えるのは楽しいですが、結局のところ、本格的な仕事に取り組もうとしている専門家にとっては役に立たないものです。
個人的には、AI から多くの価値を得ていると感じています。このように感じていない人の多くは「間違った認識」をしていると思います。つまり、言語モデルを最も役立つ方法で使用していないのです。この投稿では、スタッフ エンジニアとして私が日常的に AI を使用するさまざまな方法をリストします。
私はコード 1 を書くたびに Copilot 補完を使用します。私が受け入れるほとんどすべての補完は完全な定型文です (たとえば、関数の引数や型の入力など)。 Copilot にビジネス ロジックを生成させることはめったにありませんが、たまに起こります。私の専門分野 (Ruby on Rails など) では、LLM よりも優れた仕事ができると確信しています。これは単なる (非常に優れた) オートコンプリートです。
ただし、常に自分の専門分野に取り組んでいるわけではありません。私は、あまり馴染みのない領域 (Golang サービスや C ライブラリなど) で小さな戦術的な変更を加えていることによく気づきます。私は構文を知っており、これらの言語で個人プロジェクトを書いたこともありますが、何が慣用的であるかについてはあまり自信がありません。このような場合、私は Copilot にさらに依存します。通常、o1 モデルを有効にして Copilot チャットを使用し、コードを貼り付けて、「これは慣用的な C ですか?」と直接尋ねます。
このように LLM にさらに依存するのは危険です。何が欠けているのかわからないからです。基本的には、全体的にスマートインターンのベースラインで業務を行うことができます。私もしなければなりません

賢明なインターンのように行動し、その分野の専門家が変更をレビューしてくれるようにしてください。しかし、そのような注意点はあったとしても、この種の戦術的変更を迅速に行うことができるのは非常に大きな影響力だと思います。
私は、本番環境にまったく使用されないコードを作成するときは、LLM をより自由に使用します。たとえば、私は最近、API から公開データのチャンクを取得し、それを分類し、その分類を一連の簡単な正規表現で近似する必要がある一連の研究を行いました。このコードはすべて私のラップトップ上でのみ実行され、データを取得するコード、データを分類するために別の LLM を実行するコード、トークン化してトークン頻度を測定してスコアを付けるコードなど、基本的にすべてのコードを作成するために LLM を使用しました。
LLM は、メンテナンスの必要がなく動作するコードを書くことに優れています。一度だけ実行される非運用コード (研究など) は、これに最適です。ここで LLM を使用したことで、支援を受けなかった場合に比べて 2 倍から 4 倍の速さで作業を完了できたと言えます。
おそらく LLM を使って行う最も便利なことは、新しいドメインを学習するためのオンデマンドの家庭教師として LLM を使用することです。たとえば、先週末、私は ChatGPT-4o に大きく依存して Unity の基礎を学びました。 LLM で学習することの魔法は、質問できることです。「X はどのように機能するのか」だけでなく、「X は Y とどのように関係しているのか」といったフォローアップの質問もできることです。さらに便利なのは、「これは正しいですか」という質問をすることです。私はよく、学んだことを書き留めて LLM にフィードバックします。LLM は、どこが正しいのか、どこがまだ誤解しているのかを指摘します。私はLLMにたくさんの質問をします。
私は何か新しいことを学ぶときはたくさんメモを取ります。すべてのメモをコピー＆ペーストするだけで、LLM によってレビューしてもらえるのは素晴らしいことです。
幻覚はどうですか

ション？正直に言うと、GPT-3.5 以降、ChatGPT やクロードが頻繁に幻覚を見ていることに気づきませんでした。私が学ぼうとしている分野のほとんどは (私には理解できていないだけですが) 非常によく理解されており、私の経験では、幻覚の可能性はかなり低いことを意味します。 LLM から学んだことが根本的に間違っていたり、幻覚だったりすることが判明したというケースに遭遇したことはありません。
あまりこのようなことはしませんが、本当にバグに行き詰まったときは、ファイル全体を Copilot チャットに添付し、エラー メッセージを貼り付けて、「助けていただけますか?」と尋ねることがあります。
私がこれを行わない理由は、現在の AI モデルよりもバグハンティングがはるかに優れていると考えているからです。ほとんどの場合、副操縦士 (個人プロジェクトによってはクロード) はただ混乱するだけです。しかし、本当に行き詰まった場合には、念のため試してみる価値はあります。労力が非常に少ないからです。 LLM が検出した微妙な動作を見逃して、時間を大幅に節約できたケースが 2、3 件あったことを覚えています。
LLM はこの点ではまだそれほど優れていないため、LLM の反復や固定を解除することに多くの時間を費やすことはありません。一度だけ取得できるか試してみます。
タイプミスや論理ミスの校正
私は、ADR、技術概要、社内投稿など、かなりの量の英語文書を作成しています。私は LLM がこれらを私に代わって書くことを決して許可しません。その理由の 1 つは、現在の LLM よりも明確に書けると思うことです。その一部は、ChatGPT ハウス スタイルに対する私の一般的な嫌悪感です。
私が時々行うのは、LLM にドラフトを入力してフィードバックを求めることです。 LLM はタイプミスを見つけるのが得意で、時には私の草稿を編集するような興味深い点を指摘してくれることもあります。
バグ修正と同様に、私はこれを行うときに繰り返しはしません。フィードバックを 1 回だけ求めるだけです。通常、LLM はいくつかのスタイラスを提供します

スティックフィードバック、私はいつも無視します。
Copilot によるスマートなオートコンプリート
よく知らない分野での短期間の戦術変更 (SME によって常にレビューされます)
一度使ったら使い捨ての研究コードを大量に書く
新しいトピック (Unity ゲーム エンジンなど) について学ぶためにたくさんの質問をする
すぐに解決できる場合に備えて、最後の手段のバグ修正
長い英語コミュニケーションのための全体的な校正
これらのタスクには (まだ) LLM を使用していません。
自分がよく知っている分野で自己PRを丸ごと書いてくれる
ADR またはその他の技術コミュニケーションの作成
大規模なコードベースを調査し、物事がどのように行われるかを調べる
編集: この投稿は Hacker News で議論され、多くのコメントが付けられました。
このことについては、ポッドキャスト「Humans of Reliability」でも取り上げました。こちらから視聴できます。
免責事項: 私は GitHub で働いており、1 年間 Copilot に直接取り組んでいました。今は GitHub モデルに取り組んでいます。
この投稿を気に入っていただけた場合は、私の新しい投稿に関する更新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
AI のスロップはなぜ読むのがとても不快なのでしょうか?
Twitter 上で明らかに AI によって生成されたコンテンツを読むのは好きではありません。これには、AI の「スロップ」という軽蔑的な用語があります。これは、「人間のふりをしている AI コンテンツ」、または単に「望ましくない AI コンテンツ」のような意味です。しかし、Copilot または ChatGPT と話すときに AI が生成したコンテンツを読むのに問題はありません。何故ですか？
AI のスロップに関する適切な理論を持つことは、いくつかの理由から重要です。まず、さまざまな状況において、AI の応答がなぜ大丈夫と感じられるのか、あるいは不快に感じられるのかということは、興味深い質問です。次に、AI を使用して製品を構築している場合、それをいい加減なものとして提示しないことをお勧めします。
続きを読む...
購読する │ About │ ポッドキャスト │ 人気 │ タグ │ RSS

## Original Extract

How I use LLMs as a staff engineer sean goedecke
How I use LLMs as a staff engineer
Software engineers are deeply split on the subject of large language models. Many believe they’re the most transformative technology to ever hit the industry. Others believe they’re the latest in a long line of hype-only products: exciting to think about, but ultimately not useful to professionals trying to do serious work.
Personally, I feel like I get a lot of value from AI. I think many of the people who don’t feel this way are “holding it wrong”: i.e. they’re not using language models in the most helpful ways. In this post, I’m going to list a bunch of ways I regularly use AI in my day-to-day as a staff engineer.
I use Copilot completions every time I write code 1 . Almost all the completions I accept are complete boilerplate (filling out function arguments or types, for instance). It’s rare that I let Copilot produce business logic for me, but it does occasionally happen. In my areas of expertise (Ruby on Rails, for instance), I’m confident I can do better work than the LLM. It’s just a (very good) autocomplete.
However, I’m not always working in my areas of expertise. I frequently find myself making small tactical changes in less-familiar areas (for instance, a Golang service or a C library). I know the syntax and have written personal projects in these languages, but I’m less confident about what’s idiomatic. In these cases, I rely on Copilot more. Typically I’ll use Copilot chat with the o1 model enabled, paste in my code, and ask directly “is this idiomatic C?”
Relying more on the LLM like this is risky, because I don’t know what I’m missing. It basically lets me operate at a smart-intern baseline across the board. I have to also behave like a sensible intern, and make sure a subject-matter expert in the area reviews the change for me. But even with that caveat, I think it’s very high-leverage to be able to make these kinds of tactical changes quickly.
I am much more liberal with my use of LLMs when I’m writing code that will never see production. For instance, I recently did a block of research which required pulling chunks of public data from an API, classifying it, and approximating that classification with a series of quick regexes. All of this code was run on my laptop only, and I used LLMs to write basically all of it: the code to pull the data, the code to run a separate LLM to classify it, the code to tokenize it and measure token frequencies and score them, and so on.
LLMs excel at writing code that works that doesn’t have to be maintained. Non-production code that’s only run once (e.g. for research) is a perfect fit for this. I would say that my use of LLMs here meant I got this done 2x-4x faster than if I’d been unassisted.
Probably the most useful thing I do with LLMs is use it as a tutor-on-demand for learning new domains. For instance, last weekend I learned the basics of Unity, relying heavily on ChatGPT-4o. The magic of learning with LLMs is that you can ask questions : not just “how does X work”, but follow-up questions like “how does X relate to Y”. Even more usefully, you can ask “is this right” questions. I often write up something I think I’ve learned and feed it back to the LLM, which points out where I’m right and where I’m still misunderstanding. I ask the LLM a lot of questions.
I take a lot of notes when I’m learning something new. Being able to just copy-paste all my notes in and get them reviewed by the LLM is great.
What about hallucinations? Honestly, since GPT-3.5, I haven’t noticed ChatGPT or Claude doing a lot of hallucinating. Most of the areas I’m trying to learn about are very well-understood (just not by me), and in my experience that means the chance of a hallucination is pretty low. I’ve never run into a case where I learned something from a LLM that turned out to be fundamentally wrong or hallucinated.
I don’t do this a lot, but sometimes when I’m really stuck on a bug, I’ll attach the entire file or files to Copilot chat, paste the error message, and just ask “can you help?”
The reason I don’t do this is that I think I’m currently much better at bug-hunting than current AI models. Almost all the time, Copilot (or Claude, for some personal projects) just gets confused. But it’s still worth a try if I’m genuinely stuck, just in case, because it’s so low-effort. I remember two or three cases where I’d just missed some subtle behaviour that the LLM caught, saving me a lot of time.
Because LLMs aren’t that good at this yet, I don’t spend a lot of time iterating or trying to un-stick the LLM. I just try once to see if it can get it.
Proofreading for typos and logic mistakes
I write a fair amount of English documents: ADRs, technical summaries, internal posts, and so on. I never allow the LLM to write these for me. Part of that is that I think I can write more clearly than current LLMs. Part of it is my general distaste for the ChatGPT house style.
What I do occasionally do is feed a draft into the LLM and ask for feedback. LLMs are great at catching typos, and will sometimes raise an interesting point that becomes an edit to my draft.
Like bugfixing, I don’t iterate when I’m doing this — I just ask for one round of feedback. Usually the LLM offers some stylistic feedback, which I always ignore.
Smart autocomplete with Copilot
Short tactical changes in areas I don’t know well (always reviewed by a SME)
Writing lots of use-once-and-throwaway research code
Asking lots of questions to learn about new topics (e.g. the Unity game engine)
Last-resort bugfixes, just in case it can figure it out immediately
Big-picture proofreading for long-form English communication
I don’t use LLMs for these tasks (yet):
Writing whole PRs for me in areas I’m familiar with
Writing ADRs or other technical communications
Research in large codebases and finding out how things are done
edit: this post was discussed on Hacker News with lots of comments.
I also discussed it on the Humans of Reliability podcast, which you can watch here .
Disclaimer: I work for GitHub, and for a year I worked directly on Copilot. Now I work on GitHub Models .
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
Why does AI slop feel so bad to read?
I don’t like reading obviously AI-generated content on Twitter. There’s a derogatory term for it: AI “slop”, which means something like “AI content presenting itself as human”, or even just “unwanted AI content”. But I have no problem reading AI-generated content when I talk to Copilot or ChatGPT. Why is that?
Having a good theory of AI slop is important for a few reasons. First, it’s just an interesting question why AI responses feel okay or gross in different contexts. Second, if you’re building products with AI, it’s a good idea to not present them as slop.
Continue reading...
subscribe │ about │ podcasts │ popular │ tags │ rss
