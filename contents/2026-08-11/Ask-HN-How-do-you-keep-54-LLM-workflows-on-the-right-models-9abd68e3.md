---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49262312"
title: "Ask HN: How do you keep 54 LLM workflows on the right models?"
article_title: ""
author: "ttruett"
captured_at: "2026-08-11T18:47:25Z"
capture_tool: "hn-digest"
hn_id: 49262312
score: 1
comments: 0
posted_at: "2026-08-11T18:17:16Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: How do you keep 54 LLM workflows on the right models?

- HN: [49262312](https://news.ycombinator.com/item?id=49262312)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T18:17:16Z

## Translation

タイトル: HN に聞く: 54 の LLM ワークフローを適切なモデルに維持するにはどうすればよいですか?
HN テキスト: 54 の LLM ベースのワークフローを備えた Django アプリを持っています。最近まで、私はもっぱら AWS Bedrock 経由で Anthropic モデルを使用していましたが、Sonnet/Haiku インテリジェンスに一致し、出力速度が 3 ～ 5 倍であると思われる新しい Gemini モデルをテストするために OpenRouter をセットアップしました。検証/正規化には Pydantic AI を使用しています。私は現在、ワークフローの目的、最適化対象 (インテリジェンス、速度、コスト)、その評価、および達成する必要がある人間が判読できるバーを記述したレジストリを管理しています。すべてを追跡し、特定のワークフローに最適なモデルが確実に使用されるようにするために、人々がどのような戦略やシステムを使用しているのか興味があります。

## Original Extract

I have a Django app with 54 LLM-backed workflows. Up until recently I've exclusively used Anthropic models via AWS Bedrock but just set up OpenRouter to test the new Gemini models given they seem to match Sonnet/Haiku intelligence but with 3-5x output speed. I'm using Pydantic AI for validation/normalization. I currently maintain a registry that describes a workflow's purpose, what we're optimizing for (intelligence, speed, cost), its eval, and a human-readable bar that must be achieved. I'm curious what strategies/systems people are using to keep track of everything and ensure an optimal model is being used for a given workflow.

