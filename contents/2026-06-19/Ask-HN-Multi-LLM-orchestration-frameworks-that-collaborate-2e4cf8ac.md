---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48593979"
title: "Ask HN: Multi-LLM orchestration frameworks that collaborate?"
article_title: ""
author: "ch3coohlink"
captured_at: "2026-06-19T04:37:16Z"
capture_tool: "hn-digest"
hn_id: 48593979
score: 1
comments: 0
posted_at: "2026-06-19T01:52:36Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: Multi-LLM orchestration frameworks that collaborate?

- HN: [48593979](https://news.ycombinator.com/item?id=48593979)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T01:52:36Z

## Translation

タイトル: HN に聞く: 連携するマルチ LLM オーケストレーション フレームワーク?
HN テキスト: これが私の一般的な見解です。Gemini は高レベルのリファクタリングには優れていると思いますが、実際のコードを書くときはバグだらけです。一方、GPT/Claude はコーディングには優れていますが、リファクタリングに関してはマイナーなパッチに固執する傾向があります。彼らは、下位互換性を維持するためだけに不必要な防御的なプログラミングを投入したり、冗長なスパゲッティ コードで終わることを好みます。私のアイデアは、彼らの強みを補完することです。Gemini に方向性/アーキテクチャのアイデアを提供させ、それから GPT/Claude に議論して実装してもらいます (実際、私はこれを常に手動で行っており、結果はかなり良好です)。そこで私の質問は、この種のコラボレーションを効果的に自動化できるエージェント フレームワークはあるのかということです。既存の「サブエージェント」機能についてはよく知っていますが、私の経験では、AI が常にそれらの機能を呼び出すことを選択するとは限りません。さらに、サブエージェントが動作している間、メイン モデルは通常、アイドル状態でそこに座っているだけです。それは実際のコラボレーションというよりも、アウトソーシングされたタスクがメインのコンテキストを汚染するのを防ぐメカニズムのように感じます。少なくとも私にはそう感じられます。

## Original Extract

Here is my general take: I feel Gemini is excellent at high-level refactoring but riddled with bugs when writing actual code. On the other hand, GPT/Claude excel at coding, but when it comes to refactoring, they tend to stick to minor patches. They love throwing in unnecessary defensive programming just to maintain backward compatibility, or ending up with verbose spaghetti code. My idea is to complement their strengths: let Gemini provide the directional/architectural ideas, and then have GPT/Claude discuss and implement them (in fact, I do this manually all the time, and the results are pretty good). So my question is: Are there any Agent frameworks that can automate this kind of collaboration effectively? I'm well aware of the existing "subagent" features, but in my experience, the AI doesn't always choose to invoke them. Moreover, while the subagent is working, the main model usually just sits there idling. It feels less like actual collaboration and more like a mechanism to prevent outsourced tasks from polluting the main context—at least, that's how it feels to me.

