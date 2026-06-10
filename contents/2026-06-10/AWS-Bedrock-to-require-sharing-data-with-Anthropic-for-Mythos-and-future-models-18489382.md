---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48473166"
title: "AWS Bedrock to require sharing data with Anthropic for Mythos and future models"
article_title: ""
author: "TomAnthony"
captured_at: "2026-06-10T10:29:52Z"
capture_tool: "hn-digest"
hn_id: 48473166
score: 65
comments: 18
posted_at: "2026-06-10T08:21:38Z"
tags:
  - hacker-news
  - translated
---

# AWS Bedrock to require sharing data with Anthropic for Mythos and future models

- HN: [48473166](https://news.ycombinator.com/item?id=48473166)
- Score: 65
- Comments: 18
- Posted: 2026-06-10T08:21:38Z

## Translation

タイトル: AWS Bedrock、Mythos および将来のモデルについて Anthropic とのデータ共有を要求
HN テキスト: > Fable 5、Mythos 5、および同等以上の機能レベルを持つ Bedrock 上の将来のモデルの場合、Anthropic は Mythos クラス モデル上のすべてのトラフィックに対して 30 日間の保持を必要とします。データを限られた期間保持することで、Anthropic は単一の交換からは見えない悪用のパターンを検出できるようになります。データ保持を選択すると、データは AWS のデータとセキュリティの境界を離れることになります。ここでの発表から: https://aws.amazon.com/blogs/aws/anthropic-claude-fable-5-on-aws-mythos-class-capabilities-with-built-in-safeguards-now-available/ > 安全性調査の一部であるか、法的に保管することが法的に義務付けられているまれなケースを除き、30 日が経過すると、データは自動的に削除されます。出典: https://support.claude.com/en/articles/15425996-data-retention-practices-for-mythos-class-models

## Original Extract

> For Fable 5, Mythos 5, and future models on Bedrock with similar or higher capability levels, Anthropic will require 30-day retention for all traffic on Mythos-class models. Retaining data for a limited period allows Anthropic to detect patterns of misuse that are not visible from a single exchange. Once you opt into data retention, your data will leave AWS’s data and security boundary. From the announcement here: https://aws.amazon.com/blogs/aws/anthropic-claude-fable-5-on-aws-mythos-class-capabilities-with-built-in-safeguards-now-available/ > After 30 days, the data is deleted automatically, except in the rare cases where it's part of a safety investigation or we're legally required to keep it. From: https://support.claude.com/en/articles/15425996-data-retention-practices-for-mythos-class-models

