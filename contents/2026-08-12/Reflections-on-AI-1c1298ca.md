---
source: "https://blog.greg.technology/2026/08/11/reflections-on-ai.html"
hn_url: "https://news.ycombinator.com/item?id=49273597"
title: "Reflections on AI"
article_title: "Reflections on AI | the greg technology blog"
author: "gregsadetsky"
captured_at: "2026-08-12T15:51:49Z"
capture_tool: "hn-digest"
hn_id: 49273597
score: 2
comments: 0
posted_at: "2026-08-12T15:08:24Z"
tags:
  - hacker-news
  - translated
---

# Reflections on AI

- HN: [49273597](https://news.ycombinator.com/item?id=49273597)
- Source: [blog.greg.technology](https://blog.greg.technology/2026/08/11/reflections-on-ai.html)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T15:08:24Z

## Translation

タイトル: AI についての考察
記事のタイトル: AI についての考察 |グレッグテクノロジーブログ
説明: 不正なプロンプト

記事本文:
AI についての考察 |グレッグテクノロジーブログ
グレッグテクノロジーブログ
AI についての考察
「間違いをしないでください」
面白くないことはないので、これからも繰り返していきます。タトゥーを入れたいくらいです。それは素晴らしい人生のマントラです。あとは「事故をしないこと」と「お金をたくさん持っていること」です。
「…というプロンプトを書いてください。」
「バグのないソフトウェアを作成するためのプロンプトを書いてください」と言ったらうまくいきますか?それは機能しないでしょう。 「次の AI が…できるようにプロンプ​​トを書いてください」は指数関数的な幻覚を求めています。ツールは、どのプロンプトが「機能する」のかを知りません。もちろん、プロンプトが生成されます。それはあなたが望むプロンプトではありません。
車が大好きな9歳の子どもに、あなたの車を運転できるか尋ねたら、はい、と答えますか（運転するでしょう）
もしあなたが彼らにあなたの車を高速道路で運転できるかどうか尋ねたら、彼らはイエスと答えるだろうか（そうするだろう）
それは良い考えでしょうか (いいえ)
なぜ反射的に同意する機械を信頼するのですか
あなたは AI に何かをやらせるのに苦労しているので、「それは不可能だ、うまくいくはずだ」と考えます。
1 年前、2 年前のモデル、または 2022 年の gpt3 に同じ質問をした場合、そのモデルは問題を解決できるでしょうか?今すぐ安価なモデルを試してみてください - 難しい問題に対してはより良い結果が得られますか?いいえ。
より悪いモデルの方が成績が悪い場合、現在使用しているモデルがそのタスクに十分に適しているとどうやって判断できるのでしょうか?昨年よりも 10% または 1% 優れたモデルを使用しているが、実際には 100,000% 優れたモデルが必要な場合はどうすればよいでしょうか?
「私に何が欠けているの？」 - それはわかりません
私は不利な比較とテスト装置だけを信頼します
現時点で私が最も信用していないトリックは次のとおりです。
モデルに互いの出力を尋ね、比較させます - Agy vs fable vs codex vs opus。彼らにとって最善の方法がわからない

共同作業を行います (通常、私は 1 人に他の全員にコードを提供してもらい、その結果を読み返して統合してもらいます)。彼らがチャットできるようにローカルで IRC サーバーを実行する必要があるかもしれません。私はこれについて真剣ではありません。
私は統合の構築とリグのテストに本当に多くの時間を費やしています。ハードウェア プロジェクトで行うのは実際に楽しいのです。 - 回帰を抑制するよう努めます。はい、型付き言語/型チェッカー/Rust などが役に立ちます。しかし、これらのシステムのいずれかを「集中」し、「忘れない」状態に保つのは、本当に難しいことです。
マシンを動作させ続けるために /loop more を使い始めました。 /goal についてどう感じているかはまだわかりません。
私は Anthropic の「サイバー検証プログラム」に承認されましたが、今でも Fable に関するリクエストがブロックされることがよくあります (あるいは、CVP は Opus のみにあったのでしょうか?)。ブロックされるのは常にプライドのポイントです - 「すごい、私はモデルに勝った!」 - しかし、私はマルウェアを作成しているわけではありません。 Fable の方が断然良いです。もっと活用できたらいいのにと思います。
私自身のより良い判断に反して、残念ながら私の claude.md には「有毒な楽観主義」を禁止し、謙虚な言葉を使用するよう求める指示があります。なぜそれがうまくいくのでしょうか - 私たちは皆、神に祈っていると思いますが、今私たちのものは、独自の数兆のパラメーターがすべて連携して解き放たれています。何が欲しいのですか - 私も人間です。

## Original Extract

bad prompts

Reflections on AI | the greg technology blog
the greg technology blog
Reflections on AI
“Don’t make mistakes”
it will never not be funny, so I will keep repeating it. I almost want it tattooed. it’s a fantastic life mantra. also, “don’t have accidents” and “have more money”.
“Write a prompt that …”
would it work if you said “write a prompt to create a software that has no bugs”? it would not work. “write a prompt so that the next ai can…” is asking for exponential hallucination. the tool doesn’t know what prompt will “work” for it. of course, it will generate a prompt. it will not be the prompt you want.
if you asked a 9 year old who’s really into cars whether they could drive your car, would they say yes (they would)
if you asked them whether they could drive your car on the highway, would they say yes (they would)
would that be a good idea (no)
why do you trust machines that reflexively agree
you are struggling to get the ai to do something, and so you think: that’s impossible, it should work.
if you asked the same question to a model from 1 year ago or 2 years ago, or gpt3 from 2022, would that model be able to solve your problem? try a cheaper model right now - does it do better on the hard problem? no.
if a worse model does worse, how do you know the model you’re using right now is good enough for the task? what if you’re using a model that’s 10% better or 1% better than last year but you actually need one that’s 100,000% better?
“what am I missing?” - it can’t know
I only trust adverserial comparisons and testing rigs
Right now, my least distrusted tricks are:
to make models ask and compare each other’s outputs - agy vs fable vs codex vs opus. I don’t have a sense of the best way for them to collaborate (typically, I ask one to provide code to all others and then to read back/integrate the findings). Maybe I should run an irc server locally for them to chat. I am not not serious about this.
I truly spend so much time building integration and testing rigs - it’s actually fun to do for hardware projects! - to try to keep regressions in check. yes, typed languages/type checkers/rust, etc. help. but keeping any of these systems “focused” and “not forgetting” is truly so hard.
I’ve started using /loop more to keep the machine doing a thing. I don’t know how I feel about /goal yet.
I’ve been approved to Anthropic’s “Cyber Verification Program”, and I still get blocked requests with Fable a lot (or was CVP only for Opus?). It’s ~always a point of pride to be blocked - “wow, I beat the model!” - but I am not creating malware. Fable is definitely better - I wish I could use it more.
against my own better judgement, my claude.md unfortunately has instructions against “toxic optimism” and asking it to use humble language. why would that ever work - I guess we all pray to some gods, and ours now is the proprietary trillion parameters all relu’ing in tandem. what do you want - I’m human too.
