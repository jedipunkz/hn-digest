---
source: "https://ariya.io/2026/05/the-illusion-of-perfect-llm-code/"
hn_url: "https://news.ycombinator.com/item?id=48360797"
title: "The Illusion of Perfect LLM Code"
article_title: "The Illusion of Perfect LLM Code · ariya.io"
author: "pavel_lishin"
captured_at: "2026-06-02T00:34:31Z"
capture_tool: "hn-digest"
hn_id: 48360797
score: 3
comments: 0
posted_at: "2026-06-01T18:33:57Z"
tags:
  - hacker-news
  - translated
---

# The Illusion of Perfect LLM Code

- HN: [48360797](https://news.ycombinator.com/item?id=48360797)
- Source: [ariya.io](https://ariya.io/2026/05/the-illusion-of-perfect-llm-code/)
- Score: 3
- Comments: 0
- Posted: 2026-06-01T18:33:57Z

## Translation

タイトル: 完璧な LLM コードの幻想
記事のタイトル: 完璧な LLM コードの幻想 · ariya.io
説明: 私は最近、Web アプリに単純な認証機能を実装するというタスクを課して、いくつかの異なる LLM をテストしました。現在、ほとんどすべての最新の LLM が、構造化された青写真に従うことに優れていることは明らかです。しかし、実際の違いは、セキュリティの内部を覗いてみると明らかになりました。
[切り捨てられた]

記事本文:
ホーム → 投稿
完璧な LLM コードの幻想
私は最近、Web アプリに単純な認証機能を実装するというタスクを課して、いくつかの異なる LLM をテストしました。現在、ほとんどすべての最新の LLM が、構造化された青写真に従うことに優れていることは明らかです。ただし、実際の違いは、生成されたコードのセキュリティを内部的に見てみると明らかになりました。
私のテストでは、Opus 4.8、Gemini 3.5 Flash、Sonnet 4.6、Kimi 2.6、および DeepSeek V4 Flash を比較しました。私の目標は、これらのさまざまな LLM が現実世界のコーディング タスク、実行計画、セキュリティ監査をどの程度うまく処理できるかを確認することでした。 PLAN.md などの特定の指示ファイルを与えたところ、すべて驚くほどうまく機能しました。 Opus のようなフラッグシップの高価なモデルであっても、DeepSeek のような超手頃な価格のオプションであっても、これらの LLM は段階的な指示に従って簡単に動作するコードを生成できます。
ただし、セキュリティ (モデルが自分の仕事を自己評価する能力を含む) に関しては、状況が異なり始めます。 Opus や Gemini などのプレミアムで高度なモデルは、セキュリティ監査の実施と隠れた欠陥の発見において大きな強みを示しました。一方で、他のモデルは当たり外れが非常に多かった。
これは、現在人々がバイブコーダーと呼んでいるものにとって、重大な隠れた危険を生み出します。雰囲気コーダーとは、LLM を完全に信頼し、プロジェクトの全体的な雰囲気や流れを純粋に判断してコードを作成する人のことです。アプリケーションが画面上で正常に動作し、機能が動作する場合、バイブコーダーはすべてが完璧であるとみなします。 LLM が PLAN.md に完璧に従ったという理由だけで、彼らは成功したと感じています。
しかし、これは幻想です。ソフトウェアが外部で動作するからといって、内部が安全であるとは限りません。 LLM が独自の内部セキュリティ監査に失敗すると、危険な脆弱性が簡単に導入される可能性があります

機能をアプリケーションに組み込みます。自分でコードをレビューせずに雰囲気に全面的に依存すると、知らず知らずのうちにシステム全体が危険にさらされることになります。
LLM を判断する際に、常に公開ベンチマークに依存できるわけではありません。効率、スピード、低コストは素晴らしいことですが、安全性を犠牲にしてはいけません。開発者として、私たちは実践を続けなければなりません。これらのモデルを評価するための最良のアプローチは、独自の真に代表的なテストを作成し、コードを公開する前にコードのセキュリティを常に再確認することです。
おそらく将来的には、モデルはより高度な自己監査を実行できるようになるでしょう。コーディングハーネスも時間の経過とともに改善される可能性があります。それまでは、LLM で生成されたコードをやみくもに運用環境にロールアウトすることは、まったく無責任です。
Copyright (C) 2005-2026 アリヤヒダヤット。

## Original Extract

I recently tested several different LLMs by tasking them with implementing a simple authentication feature for a web app. It is clear that almost all modern LLMs are now excellent at following a structured blueprint. However, the real differences appeared when looking under the hood at the security
[truncated]

Home → Posts
The Illusion of Perfect LLM Code
I recently tested several different LLMs by tasking them with implementing a simple authentication feature for a web app. It is clear that almost all modern LLMs are now excellent at following a structured blueprint. However, the real differences appeared when looking under the hood at the security of the generated code.
In my testing , I compared Opus 4.8, Gemini 3.5 Flash, Sonnet 4.6, Kimi 2.6, and DeepSeek V4 Flash. My goal was to see how well these different LLMs handle real-world coding tasks, execution plans, and security audits. When I gave them a specific instruction file, like a PLAN.md , they all performed remarkably well. Whether it was a flagship, expensive model like Opus or an ultra-affordable option like DeepSeek, these LLMs could easily follow the step-by-step instructions and generate working code.
However, when it comes to security (including the models’ ability to self-assess their own work), things start to diverge. Premium, advanced models like Opus and Gemini showed great strength in conducting security audits and catching hidden flaws. On the other hand, other models were very hit-or-miss.
This creates a serious hidden danger for what people now call the vibe coder . A vibe coder is someone who trusts the LLM completely, writing code purely by judging the general vibe or flow of the project. If the application runs fine on the screen and the features work, the vibe coder assumes everything is perfect. They feel successful simply because the LLM followed the PLAN.md flawlessly.
But this is an illusion. Just because a piece of software works on the outside does not mean it is safe on the inside. When an LLM fails its own internal security audit, it can easily introduce dangerous vulnerabilities into your application. If you rely entirely on the vibe without reviewing the code yourself, you are unknowingly putting your entire system at risk.
We cannot always rely on public benchmarks to judge an LLM. Efficiency, speed, and low costs are great, but they should not come at the expense of safety. As developers, we must stay hands-on. The best approach to evaluating these models is to craft a truly representative test of your own, and always double-check the security of the code before it goes live.
Perhaps in the future, models will be advanced enough to carry out much better self-audits. Coding harnesses will likely improve over time, too. Until then, blindly rolling out LLM-generated code to production is simply irresponsible.
Copyright (C) 2005-2026 Ariya Hidayat.
