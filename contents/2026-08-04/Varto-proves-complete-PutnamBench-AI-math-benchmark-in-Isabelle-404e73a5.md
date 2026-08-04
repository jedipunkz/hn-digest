---
source: "https://varto.ai/news/putnambench/"
hn_url: "https://news.ycombinator.com/item?id=49170286"
title: "Varto proves complete PutnamBench AI math benchmark in Isabelle"
article_title: "How we solved PutnamBench | Varto"
author: "simonbohnen"
captured_at: "2026-08-04T16:08:38Z"
capture_tool: "hn-digest"
hn_id: 49170286
score: 1
comments: 0
posted_at: "2026-08-04T15:23:54Z"
tags:
  - hacker-news
  - translated
---

# Varto proves complete PutnamBench AI math benchmark in Isabelle

- HN: [49170286](https://news.ycombinator.com/item?id=49170286)
- Source: [varto.ai](https://varto.ai/news/putnambench/)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T15:23:54Z

## Translation

タイトル: Varto が Isabelle で完全な PutnamBench AI 数学ベンチマークを証明
記事のタイトル: PutnamBench をどのように解決したか |ヴァルト
説明: 学部競技数学の正式証明についての考察

記事本文:
varto ニュースお問い合わせ ← ニュースルーム 2026 年 7 月 23 日
学部競技数学の正式証明の考察
今後の Isabelle ベースの検証プラットフォームをストレステストするために、Varto Cloud を PutnamBench ベンチマークに指定しました。 Putnam Mathematical Competition は毎年行われる学部数学コンテストであり、PutnamBench は定理証明者向けに形式化されたその問題のコレクションです。
Varto は 640 の Isabelle 形式化をすべて解決しました。これは、Isabelle の最初の提出であり、ベンチマーク カテゴリ全体を解決する最初の提出です。これにより、Varto はイザベル部門で議論の余地のないナンバーワンになります。
また、イザベル カテゴリに含まれていないすべての問題を自動的に形式化しました。これらの形式化は現在、ベンチマーク管理者によって検討中です。ヴァルトはそれらすべてを証明した。
ターゲット 100 10 0 $0 $20 $40 $60 トークンコスト ターゲットごとの証明時間
ターゲット 100 10 0 0 30 60 90 120 分 証明時間 GPT-5.5 を使用し、ほとんどの証明は中程度または高度の推論労力で完了しました。
問題あたりの平均トークンコストは 8 ドルで、平均証明時間は 16 分でした。これは、2026 年 7 月 20 日時点でトップランクのリーン申請と比べて約 8 倍安く、23 倍高速です。
私たちは、トークンコストと証明時間が低いのは、主に Varto の自己学習スキル、Isabelle の高度な組み込み証明戦術、および正式証明のアーカイブによるものであると考えています。

## Original Extract

A look at formally proving undergraduate competition mathematics

varto NEWS CONTACT ← Newsroom July 23, 2026
A look at formally proving undergraduate competition mathematics
To stress-test our upcoming Isabelle-based verification platform, we pointed Varto Cloud at the PutnamBench benchmark. The Putnam Mathematical Competition is an annual undergraduate mathematics competition, and PutnamBench is a collection of its problems formalized for theorem provers.
Varto solved all 640 Isabelle formalizations. This marks the first Isabelle submission and the first submission to solve an entire benchmark category. This makes Varto the undisputed number one in the Isabelle category .
We also auto-formalized every problem missing from the Isabelle category. These formalizations are currently under review by the benchmark maintainers. Varto proved all of them as well.
Targets 100 10 0 $0 $20 $40 $60 Token cost Proof time per target
Targets 100 10 0 0 30 60 90 120 min Proof time We used GPT-5.5, with most proofs completed at medium or high reasoning effort.
Our average token cost per problem was $8, and our average proof time was 16 minutes. That is about 8× cheaper and 23× faster than the top-ranked Lean submission as of July 20, 2026.
We attribute the low token cost and proof time primarily to Varto’s self-learning skills, Isabelle’s advanced built-in proof tactics, and the Archive of Formal Proofs .
