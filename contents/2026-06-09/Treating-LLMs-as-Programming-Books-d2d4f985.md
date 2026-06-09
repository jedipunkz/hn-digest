---
source: "https://jola.dev/posts/treating-llms-as-programming-books"
hn_url: "https://news.ycombinator.com/item?id=48460384"
title: "Treating LLMs as Programming Books"
article_title: "Treating LLMs as programming books | jola.dev"
author: "shintoist"
captured_at: "2026-06-09T13:07:13Z"
capture_tool: "hn-digest"
hn_id: 48460384
score: 2
comments: 0
posted_at: "2026-06-09T12:44:08Z"
tags:
  - hacker-news
  - translated
---

# Treating LLMs as Programming Books

- HN: [48460384](https://news.ycombinator.com/item?id=48460384)
- Source: [jola.dev](https://jola.dev/posts/treating-llms-as-programming-books)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T12:44:08Z

## Translation

タイトル: LLM をプログラミング本として扱う
記事のタイトル: LLM をプログラミングの本として扱う |ジョラ.dev
説明: エンゲージメントと認知努力を失うことなく、コーディングに LLM を効果的に使用するアプローチについての考え。

記事本文:
ジョラ.dev
ホーム
について
ブログ
ニュースレター
プロジェクト
会談
ホーム
について
ブログ
ニュースレター
プロジェクト
会談
LLM をプログラミングの本として扱う
LLM は多くの場合、局所的な奇跡を起こして、機能するもの、または少なくとも機能しているように見えるものを迅速に構築できます。あなたが専門家ではないことを LLM が行うとき、それはまるで魔法のようです。 Terraform HCL を作成し、bash ツールの呼び出しをすべて認識し、Tailwind の問題に対処し (数回失敗した後)、Kubernetes クラスターをデバッグします。これらすべてを、苦労することなく実行できます。もちろん、そのトピックについて知れば知るほど、そのトピックの印象は薄れます。それは物事を必要以上に複雑にしてしまうことが多く、エレガントになることはめったにありません。それでも、コードの各行を細心の注意を払って作成するよりも、何かを構築してユーザーの前に届けることに重点を置いている場合は、非常に速く何かを作成することができます。
LLM を長期間にわたって使用すると、許容されるコードの基準が変わります。必ずしも悪化しているわけではありませんが、私が見てきたパターンでは、コードの個々の行はますます気にならなくなり、レビュー中の焦点は全体的な設計とパターンに移っています。 LLM は、あなたに代わって仕事をしてくれると信頼させるのが上手で、常にあなたの日常のワークフローをどんどん引き継ごうとします。時間が経つにつれて、調査、設計、実装、レビュー、さらにはコミット メッセージの作成や PR の作成など、プロセスのあらゆる段階でそれらに依存するようになります。
最終的には、自分の認知作業をすべてアウトソーシングしていることに気づきます。しばらくは大丈夫かも知れません。しかし、あなたも私と同じなら、これが自分でコードを設計して書く能力にどのような影響を与えているのか、最終的には心配し始めるでしょう。それだけではなく、それがあなたのアイデンティティや自己イメージそのものにどのような影響を与えているのか。
私

最近、精神的な関与を維持しながら LLM を「安全に」使用する方法についての記事をたくさん読みましたが、正直なところ、どれもあまり説得力のあるものではありませんでした。それらはすべて要約すると、「集中力を維持する」ということです。
最近、数週間仕事を休みましたが、それを利用してコーディングに再び取り組むことができました。私は SOTA モデルの使用を避けましたが、ローカルで実行中のモデルや、OpenRouter と Mistral を介していくつかのより大きなオープンウェイト モデルを試してきました。そうしているうちに、LLM 支援コーディングの別の方法を試し始めました。LLM をオンデマンドのプログラミング書籍のように扱うことです。
20年前、私がプログラミングを学び始めたとき、私は本屋に行って面白そうな本を探していました。 PHP、Linux、Apache、HTML/CSS/JS などについていくつか取り上げました。どの本も不思議でした。それを家に持ち帰り、コンピューターの前に座って入力を始めました。
本にあるすべての例を、私はすべて手書きで入力していました。なんて選択があったんだ！デジタル版はなく、コピーペーストするものは何もありませんでした。指が痛くなるまで何時間もタイプしていました。章を終えるたびに、新たな報酬、新たな成果が得られました。
コードの入力には魔法のようなものがあります。それがあなたのコードやデザインではないことさえ問題ではないようです。それは、入力している間、脳の奥深くで魔法が起こっているようなもので、命令に応じて再作成できるようになるまで、脳が構造とデザインを内部に取り込んでいるのです。
これは、プログラミング方法を学ぶための実証済みの方法です。必要なのは、本と入力するためのコンピューターだけです。そこで、LLM が登場します。倫理的考慮はさておき (そして、それらは真剣な倫理的考慮事項です。ふりはやめましょう)、LLM は、ほぼ無限のボリュームの生成されたプログラミングを生成する能力を持っています。

文献、コード例が含まれています。ただ尋ねるだけでいいのです。
ワークフローは標準的なエージェントのワークフローと似ており、自分にとって納得のいく計画を立てて、好みのアプローチが見つかるまでリサーチとディスカッションを行うことから始めます。おそらく、ローカル コードを簡単に読み取って分析できるように、Web ページ上のチャットボットとしてではなく、指定したコードベースから実行したいと考えるでしょう。ただし、ここでの大きな違いは、指示を明確にしていることです。LLM は編集を行っていません。一つもありません。
計画の準備ができたら、作業に取り掛かります。同意した計画からボットにコード例と提案を生成させ、断片ごとに抜粋して、自分で入力します。
これが痛くても驚かないでください。たくさんの文字を入力することになります。たくさんたくさん。良いキーボードを使ったほうがいいですよ。
しかし、その痛みには価値がある。この苦労は、明確で簡潔なコードの価値を教えてくれます。この痛みは、不必要なテストを書かないことを教えてくれます。膨大な労力を費やすことで、本当の目標に集中し続けることができます。その間ずっと、あなたの脳はバックグラウンドで処理を行っており、手を使わないエージェント的アプローチでは決して達成できない方法に従事することになります。
とても多くのことを学ぶことができます。それだけの価値はあります。
タイピングは別の方法で脳に働きかけます。それはまさにその通りです。手書きの方がさらに良いという研究結果もありますが、コードを実行するには扱いにくい選択肢です。それで、それはタイピングです。
コードを 1 行ずつ読んで入力するとき、夢中になることを避けることはできません。以前はとても良く見えた計画が、実は欠陥があったことに気づく瞬間がすぐに訪れるでしょう。入力しているとき、バックグラウンドで処理しているときに、ひらめきが起こるでしょう。バグや問題、予期せぬ欠点が見つかることになるでしょう

あなたのアプローチの。真っ向から立ち向かってください。関与して解決策を見つけます。そして、再び入力作業に戻ります。
入力すると常に最新情報を入手できます。あなたはそのループです。あなたは、自分が作成しているコードのすべての行を所有しています。これは、Claude や Codex があなたのためにコードを書いた場合には決して得られない方法です。コードの本当の所有権を感じることになるでしょう。コードの各行について。あらゆる技術的な決定について。
ああ、なんてことだ、とても大変な作業だ。しかし、それだけの価値はあります。
その際、ターミナルにある例をそのまま正確に入力する時間を減らし、最初に例を読んでからメモリから再作成するようにしてください。最初はうまくいかなくても問題ありません。必要に応じて何度でもコード スニペットに戻ってください。ただし、コードの一部を見てその形状を把握する能力を、再現できるまで練習してください。オリジナルとは多少異なるかもしれませんが、同じことを達成します。
これはおそらく、少なくとも私たちの多くにとって、コードを書くための理想的な方法ではありません。エージェントセッションとコードの大量生産を両立させているときほど生産性を感じられないかもしれません。あなたは依然として LLM に大きく依存していますが、公平を期すために言っておくと、おそらくもっと愚かなモデル、おそらく自分のコンピュータ上で動作するモデルでも問題なく動作するでしょう。もしかしたら、非倫理的な選択肢の中から、より倫理的な選択肢を選ぶこともできるかもしれません。
私たちは皆、このテクノロジーが世界にもたらした変化にどう対処すべきかをまだ考えているところだと思います。それが完全に回避することを意味するか、脳を集中させ続ける方法でそれを使用することを試みることを意味するかにかかわらず、私たちは前進する道を見つけるために最善を尽くすべきだと思います、あるいはおそらく、私たちの仕事は燃え尽きるまでアイデアをマシンに送り込むだけであることを全力で受け入れることだと思います。
ジョアンナ・ラーソン著。
この投稿についてどう思いますか? Bluesky で私を見つけてください

@jola.dev 。
GitHub スポンサーで私の執筆をサポートし、ブログのコンテンツを含む月刊ニュースレターを入手するか、Ko-fi でコーヒーを買ってください。
Elixir の standard.site にブログを公開する
小さなミックスタスクとatprotoクライアントを使用して、Elixir/PhoenixブログをBlueskyとstandard.siteネットワークに投稿するためのチュートリアル。
Elixir での OG イメージの生成
Elixir で静的ページとブログ投稿用の OG 画像を生成する
書くことの社会契約
下品な世界に溺れつつある中で、本物の文章を書くことの価値について。
© 2026 ジョアンナ・ラーソン。無断転載を禁じます。

## Original Extract

Thoughts on an approach for using LLMs effectively for coding without losing engagement and cognitive effort.

jola.dev
Home
About
Blog
Newsletter
Projects
Talks
Home
About
Blog
Newsletter
Projects
Talks
Treating LLMs as programming books
LLMs can often perform localized miracles, quickly building something that works, or at least gives the appearance of working. When an LLM does something that you’re not an expert on, it’s like magic. It can write Terraform HCL, knows every single bash tool invocation, tackle your Tailwind issues (after a few false starts), and debug your Kubernetes cluster, all without breaking a sweat. Of course, the more you know about a topic the less impressive it is. It often makes things more complicated than they need to be, it’s rarely elegant. Still, if you care more about building something and getting it in front of users than you do meticulously crafting each line of code, you can produce something really fast.
Working with LLMs over a longer period of time, your standards for acceptable code change. It’s not necessarily that they get worse, but the pattern I’ve seen is caring less and less about each individual line of code, and the focus during review shifting more towards overall design and patterns. LLMs are great at encouraging you to trust them to do work for you, and they’ll always attempt to take over more and more of your daily workflow. Over time you start relying on them for every step of the process, researching, designing, implementing, reviewing, even writing the commit messages and creating PRs.
Eventually you realize you’ve outsourced all of your cognitive work. It might be fine for a while. But if you’re like me, you’ll eventually start worrying about what this is doing to your ability to design and write code independently. Not just that, but what it’s doing to your very identity and self-image.
I’ve read a bunch of articles recently, on the topic of how to “safely” use LLMs while still staying engaged mentally, and honestly I didn’t find any of them very compelling. They all just boiled down to: stay focused.
I recently had a few weeks off work and I used them to reconnect with coding. I avoided using any of the SOTA models, although I have been playing around with running models locally as well as some larger open weight models through OpenRouter and Mistral. While doing that I’ve started playing around with a different way of LLM-assisted coding: treating them like on-demand programming books.
Two decades ago, when I started learning programming, I would go to the bookstore and look for interesting looking books. I got some on PHP, Linux, Apache, HTML/CSS/JS, and more. Every book was a wonder. I’d bring it home and sit down in front of the computer and I’d start typing.
Every single example in the book, I’d type it all by hand. What choice did I have! There was no digital version, nothing to copy paste. I typed for hours and hours until my fingers hurt. Every chapter I finished was a new reward, some new accomplishment.
There’s something magical about typing out code. It doesn’t even seem to matter that it’s not your code, not your design. It’s like while you type, some magic is happening deep down in your brain, where it’s internalizing the structure and design, until you’re able to re-create it on command.
It’s a tried and tested way of learning how to program. You just need books and computer to type on! And that’s where the LLMs come in. Ethical considerations aside (and they are serious ethical considerations, let’s not pretend), LLMs have the ability to produce a near infinite number of volumes of generated programming literature, code examples included. You just need to ask.
The workflow is similar to the standard agentic one, you start off by researching and discussing until you find an approach you like, with a plan that makes sense to you. You probably still want to run it from your given codebase, rather than as a chatbot on a webpage, so it can read and analyze the local code easily. The big distinction here though is that you make the instructions clear: the LLM is not making any edits. Not a single one.
Once you have the plan ready, it’s time to get to work. Have the bot produce code examples and suggestions from the plan you agreed on, snippet by snippet, and type it out yourself.
Now don’t get surprised if this hurts . You’re going to be doing a lot of typing. A lot a lot. You better be using a good keyboard.
But the pain is valuable. The pain teaches you the value of clear and concise code. The pain teaches you not to write unnecessary tests. The sheer effort involved will keep you focused on the real goals. All the while, your brain will be processing in the background, and you will be engaged in a way you can never accomplish with a hands off agentic approach.
You will learn so much . Stick with it, it’s worth it.
Typing engages your brain in a different way. It’s just the way it is. There are studies that show that writing by hand is even better, but that’s an awkward option for executing code. So typing it is.
You just can’t avoid being engaged as you read and then type line after line of code. You’ll soon have a moment where you realize that the plan that previously looked so good, actually had flaws. As you type, as you’re processing in the background, you’re going to have epiphanies. You’re going to find bugs, and issues, and unexpected drawbacks of your approach. Tackle them head on. Engage and find solutions. And then get back to typing again.
Typing keeps you in the loop. You are the loop. You own every single line of code you’re producing, in a way that you never could have if Claude or Codex wrote it for you. You’re going to feel real ownership of the code. Of every line of code. Of every technical decision.
Oh, my god it’s so much work. But worth it.
As you’re doing that, try to spend less and less time typing the example out exactly as it is in the terminal and start trying to read it first and then recreate it from memory. It’s ok if it doesn’t work at first, go back to the code snippets as many times as you need. But practice that ability of looking at a piece of code and grasping its shape, to the point where you’re able to recreate it. Maybe slightly different than the original, but accomplishing the same thing.
This is probably still not the ideal way to write code, at least not for many of us. You might not feel as productive as you do when you’re juggling agentic sessions and mass producing code. You’re still incredibly dependent on LLMs, although to be fair you can probably do alright with a dumber model, maybe even something that runs on your own computer. Maybe you can opt for a more ethical option, amongst the unethical options.
I think we’re all still figuring out how to deal with the changes that this technology has brought the world. I think we should do our best to find a path forward, whether that means avoiding it completely, trying to use it in a way that keeps our brains engaged, or, I guess, going all in and accepting that our job is now just feeding ideas into the machine until we burn out.
Written by Johanna Larsson .
Thoughts on this post? Find me on Bluesky at @jola.dev .
Support my writing on GitHub Sponsors and get a monthly newsletter with content from the blog, or buy me a coffee on Ko-fi.
Publishing your blog to standard.site in Elixir
A walkthrough for posting your Elixir/Phoenix blog to Bluesky and the standard.site network, using a small mix task and an atproto client.
Generating OG images in Elixir
Generating OG images for your static pages and blog posts in Elixir
The social contract of writing
About the value of genuine writing in a world being drowned in slop.
© 2026 Johanna Larsson. All rights reserved.
