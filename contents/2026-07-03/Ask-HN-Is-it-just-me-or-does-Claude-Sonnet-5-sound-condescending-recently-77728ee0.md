---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48774742"
title: "Ask HN: Is it just me or does Claude / Sonnet 5 sound condescending recently?"
article_title: ""
author: "alentred"
captured_at: "2026-07-03T13:42:55Z"
capture_tool: "hn-digest"
hn_id: 48774742
score: 1
comments: 1
posted_at: "2026-07-03T13:24:23Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: Is it just me or does Claude / Sonnet 5 sound condescending recently?

- HN: [48774742](https://news.ycombinator.com/item?id=48774742)
- Score: 1
- Comments: 1
- Posted: 2026-07-03T13:24:23Z

## Translation

タイトル: HN に質問: 最近、クロード / ソネット 5 が見下したように聞こえるのは私だけでしょうか?
HN text: 最近、特に物事を説明するときに、クロード コードがしばしば見下しているように聞こえることに気づきました。子供扱いされているように感じて、ツールを使って作業しているというよりは、セサミストリートのエピソードを見ているような印象を受けます。ちょうど今日、Sonnet 5 で Claude Code を実行し、比較するためにローカルの Qwen 3.6 で OpenCode を実行し、両方にプル リクエスト、同じプロンプト、同じ CLAUDE.md ファイルについて説明してもらいました。以下にいくつかの抜粋を示します: Sonnet のタイトル: 「プログラムに JSON を話すように教える」 Qwen のタイトル: 「AMQP コンシューマ向けの JSON デシリアライゼーションのサポート」 より大きなコンテキストの Sonnet の説明: 「1 つの巨大なプログラムがすべてを行うのではなく、作業は RabbitMQ を介して相互に通信する小さなアプリケーションに分割されます。」 Qwen: 「これは、AMQP データ処理アプリケーションに共有インフラストラクチャを提供する Python ライブラリです。アプリケーションは MessageConsumer クラスを拡張し、process_one_message() を実装します。フレームワークは残りを処理し、RabbitMQ からのメッセージを消費し、アプリケーション ロジックを呼び出し、結果を公開します。」ソネットの説明だけでは、このコードが実際に何をするのかさえ分かりません。そして、Claude による比喩...「バイトから DataFrame へ、そしてその逆へ」、うーん...それを逆シリアル化と呼ぶだけです。その後、Claude Code が段落を費やして、ルックアップ テーブルが「if .. else」ラダーよりも優れていることを説明してくれました。当然のことですが、Qwen は実際には変更の性質に直接焦点を当てていました。 OK、Sonnet が非常に優れた HTML レポートを作成したことは認めますが、Qwen の出力は HTML 形式の man ページを彷彿とさせます。しかし...これも同様で、細かい点はプロフェッショナルであるという印象を与えるだけです。このようなことに気づきましたか？これは Sonnet 5 の新機能ですか、それとも

前に見逃しただけですか？これを何らかの方法で修正しますか？

## Original Extract

I have noticed recently that Claude Code often sounds condescending, especially when explaining things. I feel being treated like child, gives me an impression of watching a Sesame Street episode, instead of working with a tool. Just today I ran Claude Code with Sonnet 5 and, to compare, OpenCode with a local Qwen 3.6, and asked both to explain me a pull request, same prompt, same CLAUDE.md file. Here are some excerpts: Sonnet title: "Teaching the program to speak JSON" Qwen title: "JSON deserialization support for AMQP consumers" Sonnet's description of a larger context: "Rather than one giant program doing everything, the work is split across small applications that talk to each other through RabbitMQ". Qwen: "This is a Python library that provides shared infrastructure for AMQP data processing applications: applications extend MessageConsumer class, implement process_one_message(), and the framework handles the rest, consuming messages from RabbitMQ, calling the application logic, and publishing results." Like, from the Sonnet's description alone I wouldn't even know what this code does for real! And the metaphors by Claude... "From bytes to DataFrames, and back", ugh... Just call it desirialization, come on! And then later, Claude Code spent a paragraph explaining me how a lookup table is better than an `if .. else` ladder, duh... While Qwen actually focused directly on the nature of the change. OK, I admit it Sonnet produced a super nice HTML report, Qwen's output is more reminiscent of a man page in HTML format. But... this kind of goes the same way, the niceties just create a impression of being professional. Did you notice anything like this? Is this new in Sonnet 5, or did I just miss it before? Do you remediate this somehow?

