---
source: "https://puremint.co.uk/blog/stop-your-ai-memory-file-rotting/"
hn_url: "https://news.ycombinator.com/item?id=48679461"
title: "Show HN: A Claude skill that prunes your AI's memory file, one diff at a time"
article_title: ""
author: "wonkyfruit"
captured_at: "2026-06-25T21:31:50Z"
capture_tool: "hn-digest"
hn_id: 48679461
score: 1
comments: 0
posted_at: "2026-06-25T21:30:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A Claude skill that prunes your AI's memory file, one diff at a time

- HN: [48679461](https://news.ycombinator.com/item?id=48679461)
- Source: [puremint.co.uk](https://puremint.co.uk/blog/stop-your-ai-memory-file-rotting/)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T21:30:03Z

## Translation

タイトル: HN を表示: AI のメモリ ファイルを一度に 1 つずつ削除するクロード スキル
HN テキスト: クロード用のメモリクリーナーを作成しました。おそらく他の AI でも動作すると思いますが、私は Claude でのみテストしました。これは、メモリ ファイルに蓄積された肥大化とゴミをすべて取り除き、サイズが大きすぎると動作を停止するため、動作を良くするためのものです。私がこれを構築したのは、時間の経過とともに、クロード セッションに覚えておいてほしいとお願いしていたことがいくつか欠けていて、矛盾に満ちていることに気づいたからです。これが機能する理由は、ただ勝手にプルーニングを行うのではなく、常にユーザーの意見を求めるからです。そもそもモデルがあなたの記憶を肥大化させないと信頼できないのであれば、そのモデルがその記憶を刈り取る人であるとも信頼すべきではありません。そのため、面接のような差分プロセスが必要になります。そして、何だと思いますか？そうするからうまくいくのです。他に自分の膨らみに気づいた人はいますか？まだそこに何かおかしなものが潜んでいるのを見つけましたか？私のものは逸話でいっぱいでした:D

## Original Extract

I made a memory cleaner for Claude. I guess it'll probably work in other AIs but I've only tested it in Claude and it's for getting rid of all of the bloat and crap that builds up in your memory file so that it works better because when it's too big it stops working. I built it because I noticed that over time my Claude sessions were missing some of the things that I'd asked it to remember and it was full of contradiction. The reason why this works is because it always asks you for your opinion rather than just going off and doing a prune on its own. If you don't trust the model to not bloat your memory in the first place you shouldn't trust it to be the one who prunes that memory either. That's why it has to be an interview-like, diff process. And guess what? It works because you do it that way. Anyone else noticed theirs bloat? Found anything daft still lurking in there? Mine was full of anecdotes :D

