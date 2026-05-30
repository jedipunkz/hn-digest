---
source: "https://mastodon.gamedev.place/@JeremiahFieldhaven/116654345332213390"
hn_url: "https://news.ycombinator.com/item?id=48334021"
title: "Rsync 3.4.3 has hundreds of Claude commits"
article_title: "Jeremiah Fieldhaven: \"So my systems recently updated to rsync 3.4.3, an…\" - Gamedev Mastodon"
author: "fooker"
captured_at: "2026-05-30T11:34:12Z"
capture_tool: "hn-digest"
hn_id: 48334021
score: 68
comments: 51
posted_at: "2026-05-30T08:32:57Z"
tags:
  - hacker-news
  - translated
---

# Rsync 3.4.3 has hundreds of Claude commits

- HN: [48334021](https://news.ycombinator.com/item?id=48334021)
- Source: [mastodon.gamedev.place](https://mastodon.gamedev.place/@JeremiahFieldhaven/116654345332213390)
- Score: 68
- Comments: 51
- Posted: 2026-05-30T08:32:57Z

## Translation

タイトル: Rsync 3.4.3 には数百の Claude コミットがあります
記事のタイトル: Jeremiah Fieldhaven: 「つまり、私のシステムは最近 rsync 3.4.3 に更新されました。つまり…」 - Gamedev Mastodon
説明: そのため、私のシステムは最近 rsync 3.4.3 に更新されましたが、それが起こるとすぐに、私のバックアップ システム (複数の --compare-dest= 引数を使用して増分バックアップを実行します) が完全バックアップ以外で失敗し始めました。
3.4.1 に戻すと機能します。
そこで、GitHub のソースを見て、何が起こるかを確認します。
[切り捨てられた]

記事本文:
ジェレマイア・フィールドヘイブン: 「ということで、私のシステムは最近 rsync 3.4.3 にアップデートされました。つまり…」 - Gamedev Mastodon

## Original Extract

So my systems recently updated to rsync 3.4.3, and as soon as that happened my backup system - which does incremental backups using multiple --compare-dest= arguments - started to fail on anything but a full backup.
Revert to 3.4.1 and it works.
So I go look at the source in GitHub to see what might
[truncated]

Jeremiah Fieldhaven: "So my systems recently updated to rsync 3.4.3, an…" - Gamedev Mastodon
