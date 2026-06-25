---
source: "https://haoailab.com/blogs/parallel-tree-decoding/"
hn_url: "https://news.ycombinator.com/item?id=48680042"
title: "JetSpec Enables Up to 9.64x Lossless LLM Inference Speedup with Up to 1000TPS"
article_title: "JetSpec: Breaking the Scaling Ceiling of Speculative Decoding with Parallel Tree Drafting | Hao AI Lab @ UCSD"
author: "snyhlxde"
captured_at: "2026-06-25T23:28:56Z"
capture_tool: "hn-digest"
hn_id: 48680042
score: 2
comments: 1
posted_at: "2026-06-25T22:30:56Z"
tags:
  - hacker-news
  - translated
---

# JetSpec Enables Up to 9.64x Lossless LLM Inference Speedup with Up to 1000TPS

- HN: [48680042](https://news.ycombinator.com/item?id=48680042)
- Source: [haoailab.com](https://haoailab.com/blogs/parallel-tree-decoding/)
- Score: 2
- Comments: 1
- Posted: 2026-06-25T22:30:56Z

## Translation

タイトル: JetSpec により、最大 1000TPS で最大 9.64 倍のロスレス LLM 推論の高速化が実現
記事のタイトル: JetSpec: 並列ツリー ドラフティングによる投機的デコーディングのスケーリングの上限を突破 |ハオ AI ラボ @ UCSD
説明: TL;DR: 投機的なデコードはスケーリングの上限に達します。草案予算の拡大は、承認が高く、草稿のコストが低い場合にのみ役に立ちます。以前のドラフト ヘッドはジレンマに直面しています。自己回帰ドラフタは各パスで条件を付けますが、ツリーの深さで支払いますが、ブロック拡散ドラフタは 1 つのパスでドラフトしますが、SC になります。
[切り捨てられた]

記事本文:
JetSpec: 並列ツリー ドラフティングによる投機的デコーディングのスケーリングの上限を突破 |ハオ AI ラボ @ UCSD
ホーム
JetSpec: 並列ツリー ドラフティングによる投機的デコーディングのスケーリング上限の突破
JetSpec、DFlash、AR ベースラインのデコード速度を並べて比較。
低いドラフト コスト: 1 回のドラフトヘッド前方パスで多くのツリー ノードを生成します。
高い受け入れ性: すべてのノードを、将来の絶対的な位置だけでなく、その分岐プレフィックスに基づいて条件付けします。
ロスレス検証: 凍結されたターゲットにツリーを検証させ、一致するプレフィックスのみをコミットさせます。
ツリーの製図と検証 #
一般化可能性: トレーニング データの選択と一致して、推論が重要な数学タスクとコーディング タスクで最大の効果が見られます。 JetSpec はまた、オープンエンドの会話タスクにおいて 4 倍以上の高速化を実現します。
非貪欲サンプリング: ゲインは縮小しますが一貫性を維持しており、因果関係ツリーの利点が決定論的デコードに限定されないことを示しています。
バジェット スケーリング: ツリー バジェットが大きくなると、ドラフト ツリーが分岐条件付きのままになるため、JetSpec の信頼性が高まります。
因果関係ありとなしのトレーニング #
この実験では、因果関係を適用した場合と適用しない場合の木の品質を比較します。ギャップは、最上位のブランチとターゲットの優先継続との間の、ドラフターの対数確率の差 (nats (自然対数単位)) で測定されます。ギャップが小さいということは、ターゲットが受け入れる可能性が高いブランチがツリーに含まれているため、品質が高いことを意味します。 MATH-500 では、損失重み付けを行わないとブロック拡散ヘッドが誤って調整され、コーザル ヘッドの 9.46 トークンに対してラウンドあたり平均 4.84 トークンを受け入れます。
JetSpec サービスの最適化と vLLM の統合 #
論文: https://arxiv.org/pdf/2606.18394
Web サイトとデモ: https://jetspec-project.github.io/jetspec-web/
GitHub: https://github.com/ha

オーアイラボ/JetSpec
動物園モデル: https://huggingface.co/JetSpec

## Original Extract

TL;DR: Speculative decoding hits a scaling ceiling: a larger draft budget helps only while acceptance stays high and drafting stays cheap. Prior draft heads face a dilemma: autoregressive drafters condition on each path but pay with tree depth, while block-diffusion drafters draft in one pass but sc
[truncated]

JetSpec: Breaking the Scaling Ceiling of Speculative Decoding with Parallel Tree Drafting | Hao AI Lab @ UCSD
Home
JetSpec: Breaking the Scaling Ceiling of Speculative Decoding with Parallel Tree Drafting
Side-by-side comparison of decoding speed among JetSpec, DFlash and AR baseline.
Low drafting cost: generate many tree nodes in one draft-head forward pass.
High acceptance: condition every node on its branch prefix, not just on its absolute future position.
Lossless verification: let the frozen target verify the tree and commit only the prefix it agrees with.
Tree Drafting and Verification #
Generalizability: largest gains appear on reasoning-heavy math and coding tasks, in consistent with our training data choice. JetSpec also generalizes with >4x speedup on open-ended conversational tasks.
Non-greedy sampling: gains shrink but remain consistent, showing the causal-tree benefit is not limited to deterministic decoding.
Budget scaling: larger tree budgets help JetSpec more reliably because the draft tree stays branch-conditioned.
Training with and without Causality #
In this experiment we compare tree quality with and without causality enforced. The gap is measured the drafter’s log-probability difference, in nats (natural-log units), between its top-ranked branch and the target’s preferred continuation. A small gap means the tree contains the branch the target is more likely to accept and therefore a higher quality. On MATH-500, without loss weighting the block-diffusion head is miscalibrated, accepting a mean 4.84 tokens per round against the causal head’s 9.46.
JetSpec Serving Optimization and vLLM Integration #
Paper: https://arxiv.org/pdf/2606.18394
Website and Demos: https://jetspec-project.github.io/jetspec-web/
GitHub: https://github.com/hao-ai-lab/JetSpec
Model Zoo: https://huggingface.co/JetSpec
