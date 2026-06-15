---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48545313"
title: "Show HN: Claude Token Analyzer: a token \"screen-time\" report"
article_title: ""
author: "danmeier"
captured_at: "2026-06-15T19:25:16Z"
capture_tool: "hn-digest"
hn_id: 48545313
score: 1
comments: 0
posted_at: "2026-06-15T18:38:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude Token Analyzer: a token "screen-time" report

- HN: [48545313](https://news.ycombinator.com/item?id=48545313)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T18:38:48Z

## Translation

タイトル: HN を表示: Claude Token Analyzer: トークンの「スクリーンタイム」レポート
HN テキスト: 誰もがトークンの支出について話していますが、実際に何にトークンを支出しているかについては話していません。私は、さまざまなタスクの時間を追跡して、自分が正しいことに取り組んでいるかどうかを確認するのが面白いといつも感じてきました。現在、私は 1 日の約 95% を Claude デスクトップ アプリで過ごしています。そのため、「私の時間やトークンはどこに行ったのでしょうか?」基本的には、「私がクロードで何をしたの?」という意味です。* どうやら、それはすべてすでにディスク上にあることがわかりました。すべての Claude Code / Cowork セッションは、メッセージごとの使用ブロック (入力、出力、キャッシュ作成、キャッシュ読み取り、タイムスタンプ) を含む JSONL トランスクリプトです。以下の抜粋を実行すると、Cowork と Claude Code の履歴に基づいて、いつ何を行ったかの概要が表示されます。出力例: Claude Token Analyzer · 2026-06-15 · 5 スレッド Claude Token Analyzer · 2026-06-15 · 5 スレッド
╭───────┬─────┬─────┬───────╮
│ 支出 │ トークン │ アクティブ │ キャッシュ読み取り │
│ $39 │ 27.5M │ 1.8 時間 │ 91% │
╰───────┴───────┴───────┴───────╯
アクティブ時 08:00 09:00 10:00 11:00 12:00
クロード・コード
クロード・コワーク
クロード コード — $36.56 · 3 スレッド
2026 年 6 月 14 日の 2,207 ドル 1,790 万クロード使用状況レポート
929 ドル 760 万件の LinkedIn データ エクスポート分析
519 ドル 100 万製品化計画
クロード・コワーク — $1.99 · 2 スレッド
$1.15 791K YouTube 動画概要
$0.84 156K 写真の位置特定
カール -fsSL https://gist.githubusercontent.com/danmeier2/064f7db8c0867dd... | python3 - *ccusage は行います

私のクロード コード セッションは ~/Library/Application Support/Claude/local-agent-mode-sessions/ に保存されているため表示されません。

## Original Extract

Everyone is talking about token spend, but not what we actually spend tokens on. I’ve always found it interesting to track my time across different tasks to challenge whether I was working on the right things. Nowadays, I spend around 95% of my day in the Claude desktop app, so “where did my time/tokens go?” basically means “what did I do in Claude?”* Turns out it’s all already on your disk. Every Claude Code / Cowork session is a JSONL transcript with per-message usage blocks: input, output, cache creation, cache read, and timestamps. Run the snipped below to get an overview of what you did and when based on your Cowork and Claude Code history. Example output: Claude Token Analyzer · 2026-06-15 · 5 threads Claude Token Analyzer · 2026-06-15 · 5 threads
╭────────────────┬────────────────┬────────────────┬────────────────╮
│ Spend │ Tokens │ Active │ Cache-read │
│ $39 │ 27.5M │ 1.8h │ 91% │
╰────────────────┴────────────────┴────────────────┴────────────────╯
when active 08:00 09:00 10:00 11:00 12:00
Claude Code
Claude Cowork
Claude Code — $36.56 · 3 threads
$22.07 17.9M Claude usage report for 2026-06-14
$9.29 7.6M LinkedIn data export analysis
$5.19 1.0M Productization plan
Claude Cowork — $1.99 · 2 threads
$1.15 791K YouTube video summary
$0.84 156K Picture location identification
curl -fsSL https://gist.githubusercontent.com/danmeier2/064f7db8c0867dd... | python3 - *ccusage doesn’t see my Claude Code sessions because they’re stored in ~/Library/Application Support/Claude/local-agent-mode-sessions/.

