---
source: "https://simonwillison.net/2026/Jun/11/anthropic-walks-back-policy/"
hn_url: "https://news.ycombinator.com/item?id=48488304"
title: "Anthropic Walks Back Policy That Could Sabotage AI Researchers Using Claude"
article_title: "Anthropic Walks Back Policy That Could Have ‘Sabotaged’ AI Researchers Using Claude"
author: "lumpa"
captured_at: "2026-06-11T10:35:57Z"
capture_tool: "hn-digest"
hn_id: 48488304
score: 1
comments: 0
posted_at: "2026-06-11T09:57:19Z"
tags:
  - hacker-news
  - translated
---

# Anthropic Walks Back Policy That Could Sabotage AI Researchers Using Claude

- HN: [48488304](https://news.ycombinator.com/item?id=48488304)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jun/11/anthropic-walks-back-policy/)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T09:57:19Z

## Translation

タイトル: クロードを使用して AI 研究者を妨害する可能性のある人類のウォークバック政策
記事のタイトル: クロードを使用して AI 研究者を「妨害」した可能性がある人類のウォークバック政策
説明: Wired での Maxwell Zeff の大スクープ: 「フロンティア LLM 開発のための Fable 5 の保護手段を変更して、それを可視化する予定です。」アンスロピック氏は『WIRED』への声明でこう述べた。 「私たちが作ったのは…

記事本文:
クロードを利用してAI研究者を「妨害」した可能性のある人類歩戻し政策
サイモン・ウィリソンのウェブログ
Anthropic は、クロードを使用して AI 研究者を「妨害」した可能性のある政策を撤回します。 Wired でマックスウェル・ゼフが大スクープ:
「私たちは、フロンティア LLM 開発のための Fable 5 の保護手段を変更して、それが見えるようにしています。」アンスロピック氏は『WIRED』への声明でこう述べた。 「私たちは間違ったトレードオフを行ってしまい、正しいバランスを取れなかったことをお詫びします。」
システム カードに隠された Anthropic のポリシー、Claude Fable/Mythos がユーザーに通知せずに「フロンティア LLM 開発を対象としたリクエスト」を特定し、「有効性を制限」するというものに対して、大きな抗議がありました。
彼らがこの目に見えない側面を削除したことは良いニュースです。このカテゴリーの拒否を完全に削除した方がずっと良いでしょう。
更新: Twitter の @ClaudeDevs からの詳細:
私たちは、フロンティア LLM 開発に対する Fable 5 の保護手段を可視化するための変更を展開しています。
今週から、フラグが立てられたリクエストは明らかに Opus 4.8 にフォールバックします。これは、サイバーおよびバイオに対する保護措置と同じです。それが起こるたびにこれがわかります。 API では、フラグが設定されたリクエストは拒否の理由を返します (数日以内にサーバー側でフォールバックされる予定です)。
私たちは Fable 5 をユーザーに迅速かつ安全に導入したいと考えていました。目に見える安全装置は調査できるため、堅牢である必要があり、適切に機能するまでに時間がかかります。目に見えない保護手段はより狭い範囲にターゲットを絞ることができるため、誤検知をほとんど発生させずに迅速に出荷することができます。この理由から、私たちは目に見えない保護手段を採用しましたが、それは間違ったトレードオフでした。当社が講じている安全対策とその理由を可視化する必要があります。バランスが取れていなくて申し訳ありません。
クロード・ファブルの第一印象

2026 年 6 月 5 日から 9 日まで
MicroPython と WASM を使用したサンドボックスでの Python コードの実行 - 2026 年 6 月 6 日
Claude Opus 4.8: 「控えめだが目に見える改善」 - 2026 年 5 月 28 日
これは、2026 年 6 月 11 日に投稿された、Simon Willison によるリンク投稿です。
月額 10 ドルで私をスポンサーしていただければ、その月の最も重要な LLM 開発に関する厳選された電子メール ダイジェストを入手できます。

## Original Extract

Big scoop for Maxwell Zeff at Wired: “We’re changing Fable 5’s safeguards for frontier LLM development to make them visible.” Anthropic said in a statement to WIRED. “We made the …

Anthropic Walks Back Policy That Could Have ‘Sabotaged’ AI Researchers Using Claude
Simon Willison’s Weblog
Anthropic Walks Back Policy That Could Have ‘Sabotaged’ AI Researchers Using Claude . Big scoop for Maxwell Zeff at Wired:
“We’re changing Fable 5’s safeguards for frontier LLM development to make them visible.” Anthropic said in a statement to WIRED. “We made the wrong tradeoff and we apologize for not getting the balance right.”
There's been a huge outcry about Anthropic's policy, tucked away in their system card , that Claude Fable/Mythos would identify "requests targeting frontier LLM development" and "limit effectiveness" without notifying the user.
It's good news that they're dropping the invisible aspect of this. It would be a whole lot better of they dropped this category of refusals entirely.
Update : More details from @ClaudeDevs on Twitter :
We’re rolling out changes to make Fable 5’s safeguards for frontier LLM development visible.
Starting this week, flagged requests will visibly fall back to Opus 4.8—the same as our safeguards for cyber and bio. You will see this every time it happens. On the API, any flagged requests will return a reason for their refusal (coming to server-side fallback in the next few days).
We wanted to deploy Fable 5 to our users quickly and safely. Visible safeguards can be probed, so they have to be robust, which takes time to get right. Invisible safeguards can be targeted more narrowly, allowing us to ship quickly with very few false positives. We went with invisible safeguards for this reason—and that was the wrong tradeoff. You should have visibility into the safeguards we have in place, and why. We’re sorry for not getting the balance right.
Initial impressions of Claude Fable 5 - 9th June 2026
Running Python code in a sandbox with MicroPython and WASM - 6th June 2026
Claude Opus 4.8: "a modest but tangible improvement" - 28th May 2026
This is a link post by Simon Willison, posted on 11th June 2026 .
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.
