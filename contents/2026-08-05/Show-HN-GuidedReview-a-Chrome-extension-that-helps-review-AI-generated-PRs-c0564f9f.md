---
source: "https://guidedreview.dev"
hn_url: "https://news.ycombinator.com/item?id=49179546"
title: "Show HN: GuidedReview, a Chrome extension that helps review AI generated PRs"
article_title: "Guided Review — a better way to review AI generated code"
author: "nshntarora"
captured_at: "2026-08-05T07:33:12Z"
capture_tool: "hn-digest"
hn_id: 49179546
score: 1
comments: 0
posted_at: "2026-08-05T07:09:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: GuidedReview, a Chrome extension that helps review AI generated PRs

- HN: [49179546](https://news.ycombinator.com/item?id=49179546)
- Source: [guidedreview.dev](https://guidedreview.dev)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T07:09:39Z

## Translation

タイトル: Show HN: GuidedReview、AI が生成した PR のレビューに役立つ Chrome 拡張機能
記事のタイトル: ガイド付きレビュー — AI で生成されたコードをレビューするより良い方法
説明: GitHub PR の差分を概要付きのレビュー単位にクラスター化する Chrome 拡張機能。無料のオープンソースで、独自の LLM キーを持ち込んでください。
HN テキスト: 私が見たすべての企業では、リポジトリに自動コード レビュー エージェントがセットアップされています。コードレビューの問題は解決しません。コーディング エージェントが信頼できないのと同じように、自動レビュー エージェントも信頼できません。これらは、見逃していたバグやエッジケースを見つけるのに非常に役立ちますが、あなただけが持っているもの、つまり「味」が欠けています。より多くのコンテキストが得られます。あなたは、どのコーディング エージェントよりも人、ビジネス、製品についてよく知っています。抽象化が必要でない場合はわかります。何がコメントを必要としているかはわかります。どこでルールを破るかはわかっています。そのためには... コードを「読む」ことに勝るものはありません。問題は、コードを読むのが簡単ではないことです。ツールがコードを読み取るのではなく保存するように設計されている場合は、さらに困難になります。それが私がガイド付きレビューを構築した理由です。これは、Github PR をガイド付きウォークスルーに変換する Chrome 拡張機能で、クラスター化された変更 (ファイル変更のアルファベット順の大きなスクロール可能なリストだけでなく) のコードを実際に読み取ることができます。

記事本文:
ガイド付きレビュー — AI が生成したコードをレビューするためのより良い方法 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 コンテンツにスキップ 機能 ドキュメント FAQ スター ⌘ + G インストール ⌘ + I AI で生成されたコードを人間がレビューするためのより良い方法
「コードの読み取り」が大幅に向上します。PR の差分をレビュー単位、短い要約、キーボードファーストにクラスター化します。 AI がウォークスルーを構築します - 決めるのはあなたです。
GitHub でスターを付ける ⌘ + G 無料 · オープンソース · 独自の LLM キーを持参する
製品デモ ⌘ + V を押して再生します なぜですか?
コードレビューは大変ですが、
そしてそれはますます難しくなります
Why.md AI エージェントがあなたのためにコードを書いており、あなたの PR はレビュー待ちです。
注目の新しいコード レビュー エージェントをインストールしたとき、これによって作業が簡単になるだろうと想像したでしょう。 「承認」をクリックして次の PR に進みます。
残念ながら、コード レビュー エージェントはあなたの代わりではありません (あなたの仕事は安全です、そうですか?)。
これらは、見逃していたバグやエッジケースを見つけるのに非常に役立ちますが、あなただけが持っているもの、つまり「味」が欠けています。
より多くのコンテキストが得られます。あなたは、どのコーディング エージェントよりも人、ビジネス、製品についてよく知っています。抽象化が必要でない場合はわかります。何がコメントを必要としているかはわかります。どこでルールを破るかはわかっています。
コードを「読む」ことに勝るものはありません。
問題は、コードを読むのが簡単ではないことです。ツールがコードを読み取るのではなく保存するように設計されている場合は、さらに困難になります。
AI のおかげで、コードを書くのがあまりにも簡単になりました。コードを読みやすくするためにもそれを使用する時が来ました。私たちは人間向けに設計されたエクスペリエンスを必要としています。AI を避けるのではなく、適切な場所で (少なすぎず、多すぎず) AI を使用する必要があります。

。
人間がコードを書くことがなくなったため、コードを読むことがより重要になります。
それが私がガイド付きレビューを構築した理由です。
コードを読まなければならない人間のために作られています。
clustered-changes.diff ファイルではなく変更を確認する
GitHub では、変更されたすべてのファイルがアルファベット順に表示され、それらがどのように接続されているかを理解することができます。 AI を使用して関連する変更をレビュー単位にクラスタリングするため、ユーザーはファイル リストではなく変更内容を読むことができます。
レビュー単位 1 スキーマ db.sql · seed.ts 2 認証パス auth.ts · api.ts 3 フォーム + a11y form.tsx · a11y.tsx
Navigation.keys 書き方を読んでナビゲートする
ガイド付きレビューはキーボードファーストです。キーを離れることなく、ユニットを参照し、コメントし、送信します。これは、コードを書くのと同じ筋肉の記憶です。
summaries.ai 要約 (AI なので)
すべての変更について 2 行の概要を取得します。本当に役立つ場合もあれば、そうでない場合もあります。割り引いて受け止めてください。
モデル会社は、要約を実装しない場合、API キーをブロックするようです。真剣に。
selected-tools.json ツールの企業セキュリティはすでに承認されています
会社がすでに承認している AI プロバイダー (Claude、OpenAI、または Grok) と独自の API キーを持参します。そのキーを処理するコードを読みたいですか?これはオープンソース (VC からの資金提供なし) なので、読んだり、フォークして別の場所にポイントしたり、独自のワークフローに合わせて変更したりできます。
これに関するポリシーがあるかどうか雇用主に確認してください。 2 人のスタートアップであれば、（おそらく）大丈夫です。
sk-ant-•••••• 接続済み · キーはマシン上に残る OpenAI Anthropic Grok オープン ソース
あなたのコードが当社のインフラストラクチャに触れることはありません
この拡張機能は、AI プロバイダーと GitHub と直接通信します。私たちはあなたの差分、キー、コードを決して見ることはありません。なぜなら、私たちにはそれらを参照できるものが何もないからです。
当社のサーバーを拡張する GitHub でスターを付ける 独自の LLM キーを持参する よくある質問
q

インストールする前に誰もが尋ねる質問。
faq.md では、実際にはどのように機能するのでしょうか? Chrome 拡張機能がインストールされると、GitHub PR ページに「ガイド付きレビューの開始」ボタンが追加されます。それをクリックすると、レビューが GitHub の上にオーバーレイとして開きます。
PR からの差分を読み取り、構成された AI プロバイダーに送信してレビュー単位と概要を取得し、それらを差分にマッピングして戻し、結果を表示します。
どの AI プロバイダーが機能しますか?クロード (Anthropic)、OpenAI、および Grok (xAI)。拡張機能のオプション ページでプロバイダーとモデルを選択し、独自の API キーを貼り付けます。私たちはプロバイダーをさらに追加し続けます。GitHub リポジトリでプロバイダーをリクエストするか、独自のプロバイダーを追加する PR を開いてください。
これにより API の請求額が増加しますか?レビューを開始すると、プロバイダーが 1 回呼び出され、差分がレビュー単位にクラスタリングされ、要約されます。コストは、差分の大きさと選択したモデルによって決まります。安価なモデルの小さな PR は数セントですが、フロンティア モデルの巨大な PR はそうではありません。
プロバイダーに直接料金を支払うため、費用はすべてプロバイダーのダッシュボードに表示され、当社のダッシュボードには表示されません。私たちは決してそれを見ません。
では、何も追跡していないのですか？ほとんど。このマーケティング サイトでは、ページ ビューと数回の CTA クリックに対して分析サービスを使用しています。それでおしまい。この拡張機能は何も追跡せず、GitHub と AI プロバイダー以外のサードパーティ サーバーと通信しません。
永久に無料ですか？永遠というのは長い時間だ。無料で維持できるまでは無料です。ツールの需要が増加し、維持するために多くの帯域幅が必要になった場合は、余分な時間を補う何らかの方法を導入する予定です。しかし、これはオープンソースなので、私が良い人ではないと思われる場合はフォークすることができます。
あなたは誰ですか？私は Nishant です。私の Web サイトはこちら、Twitter/X プロフィールはこちら、LinkedIn はこちら、クレジット カード番号はこちらです - 4242 4242 4242 4242
インス

Chrome 拡張機能を追加し、オプションに LLM API キーを追加して、GitHub プル リクエストを開きます。 「ガイド付きレビューを開始」をクリックして開始します。
GitHub のスター ⌘ + G © 2026 ガイド付きレビュー

## Original Extract

Chrome extension that clusters GitHub PR diffs into review units with summaries. Free, open source, bring your own LLM key.

Every company I see has an automated code review agent set up in their repository. It doesn't solve the code review problem. The automated review agents cannot be trusted the same way coding agents cannot be trusted. They're very useful for finding bugs and edge cases you've missed, but they lack something only you have — "taste." You have more context. You know the people, the business, and the product better than any coding agent. You know when an abstraction is unnecessary. You know what needs a comment. You know where to break the rules. And, for that... Nothing beats "reading" the code. Problem is, reading code is not easy. It's even harder if your tools are designed for storing code rather than reading it. That's why I built Guided Review. It's a chrome extension that converts your Github PR into a guided walkthrough so you can actually read code in clustered changes (not just big scrollable list of file changes in alphabetical order).

Guided Review — a better way to review AI generated code 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 Skip to content Features Docs FAQ Star ⌘ + G Install ⌘ + I A better way for humans to review AI generated code
Makes “reading code” wayyyy better: clusters PR diffs into review units, short summaries, keyboard-first. AI structures the walkthrough — you still decide.
Star on GitHub ⌘ + G Free · Open source · Bring your own LLM key
Product demo ⌘ + V to play Why?
Code review is hard,
and it's only getting harder
why.md AI agents are writing code for you, and you have PRs pending your review.
When you installed the hot new code review agent, you imagined it would make the job easy. You'd click "approve" and move on to the next PR.
Unfortunately, code review agents are not a replacement for you (your job is safe, yay?).
They're very useful for finding bugs and edge cases you've missed, but they lack something only you have — "taste."
You have more context. You know the people, the business, and the product better than any coding agent. You know when an abstraction is unnecessary. You know what needs a comment. You know where to break the rules.
Nothing beats "reading" the code.
Problem is, reading code is not easy. It's even harder if your tools are designed for storing code rather than reading it.
AI has made writing code too easy. It's time we used it to make reading code easy too. We need an experience designed for humans — not to avoid AI, but to use it in just the right places (not too little, not too much).
Reading the code matters more now because humans aren't writing it anymore.
That's why I built Guided Review.
Built for humans who still have to read the code.
clustered-changes.diff Review changes, not files
GitHub gives you every changed file in alphabetical order and leaves you to work out how they connect. We use AI to cluster related changes into review units, so you read the change, not the file list.
review units 1 Schema db.sql · seed.ts 2 Auth path auth.ts · api.ts 3 Form + a11y form.tsx · a11y.tsx
navigation.keys Read and navigate how you write
Guided Review is keyboard-first. Browse units, comment, and submit without leaving the keys — same muscle memory as writing code.
summaries.ai Summaries (because AI)
Get a 2 line overview of every change. Sometimes it's really helpful, sometimes it's not. Take it with a grain of salt.
Model companies apparently block your API keys if you do not implement summarization. Seriously.
approved-tools.json Tools corporate security has already approved
Bring the AI provider your company already approved — Claude, OpenAI, or Grok — and your own API key. Want to read the code that handles that key? It's open source (without VC funding), so you can read it, or fork it and point it somewhere else, or modify it for your own workflow.
Please check with your employer if they've policies about this. If you're a 2 person startup, you're (probably) fine.
sk-ant-•••••• connected · key stays on your machine OpenAI Anthropic Grok open source
Your code never touches our infrastructure
The extension talks directly to your AI provider and GitHub. We never see your diffs, your keys, or your code — because there's nothing on our end to see them with.
extension our servers Star on GitHub Bring your own LLM key FAQs
The questions everyone asks before installing.
faq.md So how does it actually work? Once the Chrome extension is installed, it adds a Start Guided Review button to the GitHub PR page. Click it and the review opens as an overlay on top of GitHub.
We read the diff from the PR, send it to your configured AI provider to get review units and summaries, map those back to the diff, and show you the result.
Which AI providers work? Claude (Anthropic), OpenAI, and Grok (xAI). You pick the provider and model in the extension's options page and paste your own API key. We'll keep adding more providers — request one on the GitHub repo , or open a PR that adds yours.
Will this run up my API bill? Starting a review makes one call to your provider to cluster the diff into review units and summarize them. Cost scales with how big the diff is and which model you picked — a small PR on a cheap model is fractions of a cent, a huge PR on a frontier model is not.
You're paying your provider directly, so whatever it costs shows up on their dashboard, not ours. We never see it.
So you don't track anything at all? Almost. This marketing site uses an analytics service for page views and a few CTA clicks. That's it. The extension doesn't track anything and doesn't talk to any third-party servers other than GitHub and your AI provider.
Is it free forever? Forever is a long time. It is free until it can be maintained for free. If demand for the tool increases and it takes a lot of my bandwidth to maintain, I will introduce some way to be compensated for the extra time. But since it's open source , you can fork it if you find I'm not a good person.
Who are you? I'm Nishant, here's my website , here's my Twitter/X profile, here's my LinkedIn , and here's my credit card number - 4242 4242 4242 4242
Install the Chrome extension, add an LLM API key in options, and open any GitHub pull request. Click Start Guided Review to begin.
Star on GitHub ⌘ + G © 2026 Guided Review
