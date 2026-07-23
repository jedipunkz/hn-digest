---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49027099"
title: "Ask HN: Have you noticed an improvement in AI responses with memory disabled?"
article_title: ""
author: "zbikowski"
captured_at: "2026-07-23T20:12:55Z"
capture_tool: "hn-digest"
hn_id: 49027099
score: 2
comments: 0
posted_at: "2026-07-23T19:49:47Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: Have you noticed an improvement in AI responses with memory disabled?

- HN: [49027099](https://news.ycombinator.com/item?id=49027099)
- Score: 2
- Comments: 0
- Posted: 2026-07-23T19:49:47Z

## Translation

タイトル: HN に質問: メモリを無効にした場合の AI 応答の改善に気づきましたか?
HN テキスト: 私は主にリサーチ、「ラバー ダッキング」、およびアイデア サウンドボードとして、月額 20 ドルのプランでクロードを使用しています (これには Fable を使用していましたが、現在は Opus 4.8 または Sonnet 5 に格下げされています)。最近、ハイコンテキストの状況で Opus 4.8 の応答品質が大幅に低下していることに気付きました。これを解決するために、メモリ システムを無効にしました (私の知る限り、代替品はありませんが、現在は「レガシー」というラベルが付いています)。この設定を無効にすると、タスクの継続に必要なコンテキストを毎回入力するという少々面倒な作業を除けば、応答の品質と精度が大幅に向上していることに気づきました。私は Opus 4.7 のリリース後に少し泣き言を言った愚か者の一人ですが、私の理論では、メモリ コンテキストを増やすために何らかの変更が加えられたため、実際の応答品質が低下したのではないかと考えています。一般に、SOTA 思考モデルのハイコンテキスト能力は誇張されているのではないかと思います。簡単な逸話として、私は 30 を超えるプロンプトを含むリサーチ タスクに取り組んでいました。新しいチャットに切り替えて、ディスカッションの要点を 1 つの新しいプロンプトに凝縮すると、その主題に関する以前の会話のコンテキストが豊富であったにもかかわらず (またはそのおかげであると私は主張します)、前のチャットをプロンプトし続けるよりも優れた結果が得られました。余談: 「プロジェクト」と「チャット」のメモリが別々の設定だったらよかったのにと思います。メモリを無効にするとチャットの機能が向上しますが、内部メモリが無効になるためプロジェクトが大幅に弱くなり、ファイルの参照を要求するたびにモデルがファイルを探し回るようになります。ハックな解決策としては、すべてのプロジェクトを作成することも考えられますが、これは面倒なようです。これは既知の現象ですか?最近の変化によって事態はさらに悪化しましたか?ご意見やヒントをいただければ幸いです。ありがとう！

## Original Extract

I use Claude on the $20/mo plan, mainly for research, "rubber duckying," and as an idea soundboard (I was using Fable for this but am now relegated to Opus 4.8 or Sonnet 5), and recently noticed a severe downturn in Opus 4.8 response quality in high-context situations. To attempt to remedy this, I disabled the memory system (now labeled "Legacy" even though to my knowledge there is no replacement). With this setting disabled, aside from the slight tedium of inputting the necessary context each time for task continuations, I have noticed a drastic improvement in response quality and accuracy—I'm one of those shmucks who whined a bit after Opus 4.7 came out, and my theory is some change was made to increase memory context that resulted in a decrease in actual response quality. In general, I wonder if the high-context abilities of SOTA thinking models are overstated. As a brief anecdote, I was working through a research task with upwards of 30 prompts. Switching to a new chat and condensing the gist of the discussion into a single new prompt yielded superior results than continuing to prompt the prior chat, despite (or due to, I would argue) the prior conversation's wealth of context on the subject. Aside: I wish that "project" and "chat" memory were separate settings. Disabling memory boosts the capabilities of chat, but severely nerfs projects by disabling their internal memory, leading the model to grope around for files each time you ask it to reference one. I suppose a hacky solution could just be to create a project for everything, but this seems tedious. Is this a known phenomenon? Have recent changes exacerbated things? Any thoughts and tips are appreciated. Thanks!

