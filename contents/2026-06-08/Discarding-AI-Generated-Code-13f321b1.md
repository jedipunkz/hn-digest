---
source: "https://www.dotnetperls.com/2026_6_8_discarding-ai-generated-code"
hn_url: "https://news.ycombinator.com/item?id=48452141"
title: "Discarding AI-Generated Code"
article_title: "Discarding AI-Generated Code [2026] - Dot Net Perls Blog"
author: "speckx"
captured_at: "2026-06-08T21:37:30Z"
capture_tool: "hn-digest"
hn_id: 48452141
score: 2
comments: 0
posted_at: "2026-06-08T21:11:19Z"
tags:
  - hacker-news
  - translated
---

# Discarding AI-Generated Code

- HN: [48452141](https://news.ycombinator.com/item?id=48452141)
- Source: [www.dotnetperls.com](https://www.dotnetperls.com/2026_6_8_discarding-ai-generated-code)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T21:11:19Z

## Translation

タイトル: AI によって生成されたコードの破棄
記事タイトル: AI 生成コードの破棄 [2026] - Dot Net Perls Blog

記事本文:
AI 生成コードの破棄 [2026] - Dot Net Perls Blog ホーム カテゴリ ブログ AI 生成コードの破棄 リンク 2026 年 6 月 8 日更新 AI 生成コードの破棄
私の Rust プログラムには、約 100 万回実行されるホットな関数があります。小さな変更でも、プログラム全体の実行が速くなったり、(より一般的には) 遅くなったりする可能性があります。
DeepSeek V4 Flash を使用して、この機能に対する多くの最適化の可能性を検討しました。最初に追加した最適化の 1 つが少し役に立ちました。残念ながら、その後のものはどれもそうではありませんでした。
AIでも現実を自分の意のままに曲げることはできないようだ。すでに最適化された関数を最適化することが不可能な場合があります。その後、周囲のコードを書き換えてプログラムを高速化する方法を見つけました。これは最適化としてカウントされると思います。
しかし、これは、超インテリジェントな AI が驚くべき方法でコードを最適化するという私の幻想を払拭するものです。これは不可能です。幸いなことに、DeepSeek V4 Flash は非常に安価なので、失敗した実験の多くはそれほど費用がかかりませんでした。 AI が生成したコードの多くは破棄しました。

## Original Extract

Discarding AI-Generated Code [2026] - Dot Net Perls Blog Home Categories Blog Discarding AI-Generated Code Links Updated Jun 8, 2026 Discarding AI-Generated Code
There is a hot function in my Rust program that runs about a million times . Even a small change can cause the entire program to run faster or (more usually) slower.
I used DeepSeek V4 Flash to explore many possible optimizations to this function. One of the first optimizations I added helped a bit. Unfortunately, none of the later ones did.
It seems even AI cannot bend reality to one's will. Optimizing an already-optimized function is sometimes just not possible . Later I did find a way to rewrite some surrounding code and speed up the program; and this I suppose counts as an optimization.
But it does sort of dispel my fantasies of super-intelligent AIs optimizing code in amazing ways. This is just not possible. Fortunately DeepSeek V4 Flash is very cheap, so many of the failed experiments did not cost me much. I discarded much of the AI-generated code.
