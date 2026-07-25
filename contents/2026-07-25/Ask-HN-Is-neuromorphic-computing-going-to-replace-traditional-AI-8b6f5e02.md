---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49045970"
title: "Ask HN: Is neuromorphic computing going to replace traditional AI?"
article_title: ""
author: "lennart-rth"
captured_at: "2026-07-25T09:28:05Z"
capture_tool: "hn-digest"
hn_id: 49045970
score: 1
comments: 0
posted_at: "2026-07-25T09:20:28Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: Is neuromorphic computing going to replace traditional AI?

- HN: [49045970](https://news.ycombinator.com/item?id=49045970)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T09:20:28Z

## Translation

タイトル: HN に聞く: ニューロモーフィック コンピューティングは従来の AI に取って代わるのでしょうか?
HN テキスト: 私は最近、現在の深層学習の中核となる非効率性について考えながら、ニューロモーフィック コンピューティングについて学びました。現在の AI は、最適な重みの更新を見つけるための超高密度レイヤー、すべてのレイヤーとニューロンの常に継続的な計算、およびグローバル バックプロパゲーションに大きく依存しています。しかし、人間の脳を見ると、そのようなことは起こりません。脳は、現代の LLM とは対照的な原理に基づいて動作します。 - 局所進化: ニューロンは大部分が独立しており、グローバルなエラー信号ではなく、局所的な近傍と、神経伝達物質 (ドーパミンなど) のような単純なフィードバック ループに基づいて進化します。
- 極度のスパース性: システムは非常にスパースです (ニューロンは、発火チェーンに関与している場合にのみ進化し、更新されます)。
- イベント駆動型処理: ニューロンは評価されたときにのみ起動し、実際には他のスパイク ニューロンによってトリガーされます。対照的に、現在の LLM は、すべての層とニューロンを同時にトレーニングし、最適なグローバル更新関数を見つけようとすることで、問題を強引に解決しているように私には見えます。過去数年間の AI 進歩の軌跡を見ると、明確なパターンが現れます。 - 2020 ～ 2022 年: データセットと生のコンピューティングのスケールアップ。
- 2023 ～ 2024 年: コンテキスト ウィンドウを拡大し、専門家混合 (MoE) に移行。
- 2024 ～ 2025 年: 思考連鎖と推論時間推論。
- 2025 – 現在: 自律実行および並列マルチエージェント システム。基本的に、これらの進歩はどれも、コンピューティングおよび処理されたトークンをスケールアップする別の方法にすぎません。
したがって、電力は無制限ではないため、このスケーリングは当然、ある時点で壁にぶつかります。
だからこそ、長期的な進歩は不可能だと私は信じています

永遠にスケーリングすることから来ています。必要なのは根本的な効率の改善です。それでは、ニューロモーフィック処理はまさにそれである可能性があるのでしょうか、それともまだ成熟していないのでしょうか?

## Original Extract

I’ve recently learned about neuromorphic computing while thinking about the core inefficiencies of current deep learning. Today AI relies heavily on super-dense layers, continuous computation of all layers and neurons at all times, and global backpropagation to find optimal weight updates. But when you look at the human brain, none of that happens. The brain operates on principles that stand in contrast to modern LLMs: - Local Evolution: Neurons are largely independent, evolving based on their local neighborhood and simple feedback loops like neurotransmitters (e.g., dopamine) rather than a global error signal.
- Extreme Sparsity: The system is massively sparse (neurons only evolve and get updated when they have been involved in a firing-chain).
- Event-Driven Processing: Neurons are only firing when they get evaluated and are actually triggered by other spiking neurons. In contrast, current LLM‘s seem to me like they brute-force their way through problems by training all layers and neurons at the same time and trying to find the optimal global update function. If you look at the trajectory of AI advancements over the last few years, a clear pattern emerges: - 2020–2022: Scaling up datasets and raw compute.
- 2023–2024: Expanding context windows and shifting to Mixture of Experts (MoE).
- 2024–2025: Chain-of-thought and inference-time reasoning.
- 2025–Present: Autonomous execution and parallel multi-agent systems. Fundamentally, every single one of these advancements is just a different way of scaling up compute and processed tokens.
So this scaling will naturally hit a wall at some point as electricity is not unlimited.
That’s why I believe that long term progress can not come from just scaling forever. What it needs is radical efficiency improvements. So could neuromorphic processing be exactly that or is it not mature yet?

