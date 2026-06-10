---
source: "https://simonwillison.net/2026/Jun/10/if-claude-fable-stops-helping-you/"
hn_url: "https://news.ycombinator.com/item?id=48470557"
title: "If Claude Fable stops helping you, you'll never know"
article_title: "If Claude Fable stops helping you, you'll never know"
author: "behnamoh"
captured_at: "2026-06-10T04:35:13Z"
capture_tool: "hn-digest"
hn_id: 48470557
score: 6
comments: 2
posted_at: "2026-06-10T02:22:46Z"
tags:
  - hacker-news
  - translated
---

# If Claude Fable stops helping you, you'll never know

- HN: [48470557](https://news.ycombinator.com/item?id=48470557)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jun/10/if-claude-fable-stops-helping-you/)
- Score: 6
- Comments: 2
- Posted: 2026-06-10T02:22:46Z

## Translation

タイトル: クロード・ファブルがあなたを助けるのをやめたら、あなたは決して知りません
説明: Jonathon Ready は、Fable 5 と Mythos 5 の 319 ページのシステム カードから眉をひそめるような詳細の 1 つを強調しています。ここに長い抜粋があり、私のハイライトは次のとおりです。

記事本文:
クロード・ファブルがあなたを助けるのをやめても、あなたは決して知りません
サイモン・ウィリソンのウェブログ
クロード・ファブルがあなたを助けるのをやめたら、あなたは決して知りません (経由) Jonathon Ready は、Fable 5 と Mythos 5 の 319 ページのシステム カードから眉をひそめるような詳細の 1 つをハイライトしています。ここに長い抜粋があり、私のハイライトを示します。
最近のモデルが独自の開発を加速できることを考慮して、フロンティア LLM 開発 (たとえば、事前トレーニング パイプライン、分散トレーニング インフラストラクチャ、または ML アクセラレータ設計の構築など) をターゲットとするリクエストに対するクロードの有効性を制限する新しい介入を実装しました。 Claude を使用して競合モデルを開発することは、すでにサービス利用規約に違反していますが、安全策を通じてこの制限を強制することで、これらの利用規約に違反しようとする攻撃者が加速することを回避できます。
サイバーセキュリティ、生物学と化学、蒸留の試みに対する当社の介入とは異なり、これらの保護手段はユーザーには見えません。 Fable 5 は、別のモデルにフォールバックしません。代わりに、安全策は、即時変更、ステアリングベクトル、パラメータ効率の良い微調整 (PEFT) などの方法を通じて有効性を制限します。これらの介入は、コーディング作業の大部分には影響しません。これらはトラフィックの約 0.03% に影響を及ぼし、0.1% 未満の組織に集中すると推定されます。
Anthropic がこの種のサイレント介入を発表したのはこれが初めてだと思います。この正当化は私にとって依然としてかなり SF のように感じられます。リンクされた記事は「再帰的自己改善」について述べています。私は、Anthropic 自身の目標と矛盾する可能性のある研究を単に遅らせるために、「ML アクセラレーターの設計」に関する質問への回答を黙って破損するモデルにはまったく興味がありません。
クロード・ファブル5～9thの第一印象

2026年6月
MicroPython と WASM を使用したサンドボックスでの Python コードの実行 - 2026 年 6 月 6 日
Claude Opus 4.8: 「控えめだが目に見える改善」 - 2026 年 5 月 28 日
これは、2026 年 6 月 10 日に投稿された、Simon Willison によるリンク投稿です。
月額 10 ドルで私をスポンサーしていただければ、その月の最も重要な LLM 開発に関する厳選された電子メール ダイジェストを入手できます。

## Original Extract

Jonathon Ready highlights one of the more eyebrow-raising details from the 319 page system card for Fable 5 and Mythos 5. Here's a longer excerpt, highlights mine: In light of …

If Claude Fable stops helping you, you'll never know
Simon Willison’s Weblog
If Claude Fable stops helping you, you'll never know ( via ) Jonathon Ready highlights one of the more eyebrow-raising details from the 319 page system card for Fable 5 and Mythos 5. Here's a longer excerpt, highlights mine:
In light of the ability of recent models to accelerate their own development , we’ve implemented new interventions that limit Claude’s effectiveness for requests targeting frontier LLM development (for example, on building pretraining pipelines, distributed training infrastructure, or ML accelerator design ). Using Claude to develop competing models already violates our Terms of Service , but enforcing this restriction through our safeguards avoids accelerating the actors most willing to violate these terms.
Unlike our interventions for cybersecurity, biology and chemistry, and distillation attempts, these safeguards will not be visible to the user . Fable 5 will not fall back to a different model. Instead, the safeguards will limit effectiveness through methods such as prompt modification, steering vectors, or parameter-efficient fine-tuning (PEFT). These interventions will not affect the vast majority of coding work. We estimate they will impact ~0.03% of traffic, concentrated in fewer than 0.1% of organizations.
I believe this is the first time Anthropic have announced these kinds of silent interventions. The justification still feels pretty science-fiction to me - the linked article talks about "recursive self-improvement". I'm not at all keen on a model that silently corrupts its replies to questions about "ML accelerator design" purely to slow down research that might conflict with Anthropic's own goals!
Initial impressions of Claude Fable 5 - 9th June 2026
Running Python code in a sandbox with MicroPython and WASM - 6th June 2026
Claude Opus 4.8: "a modest but tangible improvement" - 28th May 2026
This is a link post by Simon Willison, posted on 10th June 2026 .
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.
