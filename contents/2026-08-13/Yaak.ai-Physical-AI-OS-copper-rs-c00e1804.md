---
source: "https://www.yaak.ai/blog/yaak-x-copper"
hn_url: "https://news.ycombinator.com/item?id=49290869"
title: "Yaak.ai Physical AI OS: copper-rs"
article_title: "Physical AI OS: copper-rs"
author: "gbin"
captured_at: "2026-08-13T19:51:58Z"
capture_tool: "hn-digest"
hn_id: 49290869
score: 1
comments: 0
posted_at: "2026-08-13T19:34:29Z"
tags:
  - hacker-news
  - translated
---

# Yaak.ai Physical AI OS: copper-rs

- HN: [49290869](https://news.ycombinator.com/item?id=49290869)
- Source: [www.yaak.ai](https://www.yaak.ai/blog/yaak-x-copper)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T19:34:29Z

## Translation

タイトル：Yaak.ai 物理AI OS：copper-rs
記事タイトル：物理AI OS：copper-rs
説明: ヤク

記事本文:
"> プラットフォーム ハードウェア研究会社と提携する Nutron をお試しください
前回のブログ投稿では、物理世界のインテリジェンスは物理世界からもたらされる必要があるという核となる信念を強調しました。通信時間は短くなりますが、物理 AI の成功にとって最も重要な隣接する信念: モデルは製品ではありません、モデルは常に進化し、モデルの自律性を調整するハードウェアとソフトウェアのエコシステムが製品です。
現場での信頼性を確保するためにハードウェアを成熟させ、安全要件を検証し、ダウンタイムを最小限に抑えてエラーから回復することは、物理 AI スタックをターゲット環境 (公道、倉庫、閉鎖された建設現場) に展開するために不可欠です。ターゲット環境にデプロイすることで、AI の学習速度が向上します。これは、単一の実施形態に対して複数年にわたる取り組みである。それぞれが異なるターゲット環境を持つ数十の実施形態でその成功を再現するには、SW と HW の両方の中核となる設計原則として実施形態にとらわれないことが必要です。
当社の空間インテリジェンス キットは、その中核として、実施形態に依存しないハードウェア設計原則に従っており、高い学習速度を実現するために最適化されたcopper-rs (オープンソースの物理 AI OS) と組み合わされています。
レイテンシが 100 倍向上、ロギング スループットが 10 倍向上
レイテンシが 100 倍向上、ロギング スループットが 10 倍向上
物理AIの最適な学習速度
モデル≠堀。物理的な AI スタックを相互リンクされたエンジニアリング システムのネットワークとして見ると、スタック全体がその最も弱い部分と同じくらい強力であることは明らかです。オフライン ベンチマークで改善が示されたリリース候補 AI は、ハードウェアと物理 AI OS の統合、シナリオのバリエーション全体での検証、安全要件に対するトレース、および現場での監視を経る必要があります。 AI の学習速度は matu によって設定されます。

モデルのトレーニング時間だけでなく、物理的な HW および SW スタック全体の重要性も考慮します。
Copper は、静的に記述され、オープンに構築された決定論的なランタイムを持ち、初日から安全認証を念頭に置いているため、Yaak の高速化に役立ちます。物理的な AI OS はモデルに合わせて設計されており、モデルに合わせて後付けされるものではありません。
ビルドループ≠AIループ。物理的な AI では、反射神経が安全性の敵対者として扱われることがよくあります。各ビルド リリースのテストは数週間から数か月かかるため、結果が上流に戻る前に、パフォーマンスの欠点は静かに蓄積されます。モデル トレーニング ループ内に物理 AI スタックを導入することにより、マルチモーダル ログが HIL (ハードウェア イン ザ ループ) で決定論的に再生され、すべての変更がテストされ、安全要件の回帰が安価なときに表面化することになります。フィードバック ループの速度は安全性を実現します。
Copper のビルド、実行、および再生により、記録されたすべてのフィールド実行が反復可能な回帰テストに変わります。 Nutron と統合されているため、回帰は即座に分離され、表面化されます。
エンドツーエンドアルゴリズムの決定論的再生
エンドツーエンドアルゴリズムの決定論的再生
デモ ≠ 導入。製品デモが出荷時に保たれていないため、物理 AI の導入ギャップはますます拡大しています。フリートはターゲット環境でデータを生成し、そのデータによってモデルが改善され、そのモデルがフリートにデプロイされ、デプロイメントのギャップが埋まるにつれてループの各ターンが高速化されます。
SIK with Copper は、前方展開された物理 AI スタックです。ターゲット環境でマルチモーダル データを収集し、AI を調整し、あらゆる実施形態の物理資産をオンエッジで作動させます。
私たちは、長い導入スケジュールを短縮するために、空間インテリジェンス キット (SIK) と Nutron を構築しました。
1. Spatial Intel をペアリングします

アセットを含むligenceキット
2. AI がお客様の操作を観察して学習するため、通常どおりに資産を操作します。
3. 自律型ワーカーとしてフリートの一部になる準備が整うと、ping が送信されます。
今後数年間で、複数の実施形態と組み合わせた SIK を、Copper を使用した 100 万の物理資産に導入する予定です。
Yaak は物理的な資産を自律的な労働者に変換します。パレットジャッキ、フォークリフト、ホイールローダー、道路清掃車のフリートを運用していて、自律運転を検討する準備ができている場合は、話しましょう。
© 2026 Yaak Technologies ApS.無断転載を禁じます。

## Original Extract

Yaak

"> Partner with us Platform Hardware Research Company Try Nutron
Our last blog post highlighted a core belief: intelligence for the physical world must come from the physical world. An adjacent belief that gets less airtime, but is paramount for physical AI’s success: The model is not the product, models evolve all the time, the ecosystem of HW and SW that orchestrates autonomy with models is the product.
Maturing the hardware for reliability in the field, validating safety requirements and recovering from errors with minimal downtime is quintessential for deploying the physical AI stack in the target environment (public roads, warehouses, closed off construction sites). Deploying in the target environment is what feeds the learning velocity of AI. This is a multi-year effort for a single embodiment. Replicating its success for dozens of embodiments, each with a different target environment requires embodiment agnosticism to be a core design principle of both SW and HW.
Our Spatial Intelligence Kit , at its core follows an embodiment agnostic hardware design principle, coupled with copper-rs ( an open source physical AI OS ) we’ve optimized for high learning velocity.
100x better latency, 10x higher logging throughput
100x better latency, 10x higher logging throughput
Optimal learning velocity for physical AI
Model ≠ Moat. Viewing the physical AI stack as a network of interlinked engineering systems, it's clear that the entire stack is as strong as its weakest link. A release candidate AI that has shown improvements on offline benchmarks still has to go through HW and physical AI OS integration, validation across scenario variations, traced against safety requirements, as well as monitored in the field. The learning velocity for AI is set by the maturity of the entire physical HW and SW stack and not just by the training time of the models.
Copper helps accelerate Yaak by being statically described, having a deterministic runtime that is built in the open, and by having safety certification in mind from day one. The physical AI OS is designed for the model, not retrofitted around it.
Build loop ≠ AI loop. Quite often within physical AI, reflex is to treat learning velocity as an antagonist to safety. Testing each build release spans weeks or months and thus performance shortcomings age quietly, before the results flow back upstream. By bringing the physical AI stack within the model training loop means that a multimodal log replays deterministically with HIL (hardware in the loop), every change gets tested, and any regression on safety requirements surface when they are cheap. The velocity of the feedback loop is an enabler for safety.
Copper's build, run, and replay turns every recorded field-run into a repeatable regression test. Integrated with Nutron , regression gets isolated and surfaced instantly.
Deterministic replay of an end to end algorithm
Deterministic replay of an end to end algorithm
Demos ≠ Deployment. The deployment gap in physical AI is ever growing as product demos don’t hold water when shipped. A fleet generates data in the target environment, the data improves the model, the model is deployed back to the fleet, and every turn of the loop gets faster as the deployment gap is bridged.
SIK with Copper is the physical AI stack, forward deployed: collecting multimodal data in the target environment, to tune the AI and actuate the physical asset on-edge for any embodiment.
We built the Spatial Intelligence Kit (SIK) and Nutron to collapse the long deployment timelines:
1. Pair our Spatial Intelligence Kit with the asset
2. Operate the asset as usual as our AI learns from observing your operations
3. You get pinged when it's ready to be part of your fleet as an autonomous worker
We are deploying SIK paired with multiple embodiments on a million physical assets with Copper in the next years.
Yaak transforms physical assets into autonomous workers. If you’re running a fleet of pallet jacks, forklifts, wheel loaders, street sweepers and you’re ready to explore autonomous operations, let’s talk .
© 2026 Yaak Technologies ApS. All rights reserved.
