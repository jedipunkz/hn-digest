---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49378314"
title: "Ask HN: How do you review and validate LLM generated code?"
article_title: ""
image: ""
author: "darkLord19"
captured_at: "2026-08-20T19:24:47Z"
capture_tool: "hn-digest"
hn_id: 49378314
score: 3
comments: 2
posted_at: "2026-08-20T18:26:05Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: How do you review and validate LLM generated code?

- HN: [49378314](https://news.ycombinator.com/item?id=49378314)
- Score: 3
- Comments: 2
- Posted: 2026-08-20T18:26:05Z

## Translation

タイトル: HN に聞く: LLM で生成されたコードをどのようにレビューして検証しますか?
HN テキスト: 最近、コードの生成が非常に速くなりましたが、すべてをどのようにレビューしてテストしますか?モデルにコードの一部をレビューするよう依頼するたびに、モデルはさまざまな「結果」を返します。したがって、実際の欠陥を見つけるには、調査結果のノイズをフィルタリングする必要があります。また、完全なレビューを手動で行おうとすると、コードを理解するだけでも、最初から自分でコードを書くよりも時間がかかる可能性があります。テストに関しては、LLM にテストの作成を依頼できますが、そのテストが実際に期待される製品フローをアサートしているのか、それともロジックの一部の合否テストを作成しているだけなのかはわかりません。ワークフローや戦略、コツなどがあれば知りたいです。

## Original Extract

It's become very fast to generate code nowadays, but how do you review and test all of it? Every time you ask a model to review some piece of code, it comes back with different "findings". So, you still have to filter through the noise of their findings to get to actual defects. And if you try to do the full review manually, then it might take more time to just understand the code than writing it yourself from the beginning. As for testing, you can ask the LLM to write the tests, but you can't be sure if the tests actually assert the expected product flow or just write pass/fail tests for a piece of logic. I would like to know if you have any workflows, strategies, or tricks.

