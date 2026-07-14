---
source: "https://arxiv.org/abs/2607.05391"
hn_url: "https://news.ycombinator.com/item?id=48914113"
title: "LLM-as-a-Verifier: A General-Purpose Verification Framework"
article_title: "[2607.05391] LLM-as-a-Verifier: A General-Purpose Verification Framework"
author: "gmays"
captured_at: "2026-07-14T23:42:46Z"
capture_tool: "hn-digest"
hn_id: 48914113
score: 2
comments: 0
posted_at: "2026-07-14T23:09:17Z"
tags:
  - hacker-news
  - translated
---

# LLM-as-a-Verifier: A General-Purpose Verification Framework

- HN: [48914113](https://news.ycombinator.com/item?id=48914113)
- Source: [arxiv.org](https://arxiv.org/abs/2607.05391)
- Score: 2
- Comments: 0
- Posted: 2026-07-14T23:09:17Z

## Translation

タイトル: LLM-as-a-Verifier: 汎用検証フレームワーク
記事のタイトル: [2607.05391] LLM-as-a-Verifier: 汎用検証フレームワーク
説明: arXiv 論文 2607.05391 の要約ページ: LLM-as-a-Verifier: A General-Purpose Verification Framework

記事本文:
メインコンテンツにスキップ
arXiv は独立した非営利団体になりました。
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 人工知能
[2026 年 7 月 6 日に提出 ( v1 )、最終改訂日 2026 年 7 月 7 日 (このバージョン、v2)]
タイトル: LLM-as-a-Verifier: 汎用検証フレームワーク
要約: トレーニング前、トレーニング後、テスト時のコンピューティングのスケーリングは、LLM の機能を向上させるための中心的なパラダイムとなっています。この研究では、ソリューションの正しさを判断する能力である検証を新しいスケーリング軸として特定します。これを解き放ち、その有効性を実証するために、追加のトレーニングを必要とせずにエージェント タスクに対するきめ細かいフィードバックを提供する汎用検証フレームワークである LLM-as-a-Verifier を導入します。 LLM に候補解に対する離散スコアの生成を促す標準の LM ジャッジとは異なり、検証者としての LLM は、スコアリング トークン ロジットの分布に対する期待値を計算して連続スコアを生成します。この確率的定式化により、(1) スコアの粒度、(2) 反復評価、および (3) 基準の分解といった複数の次元に沿って検証を拡張することができます。特に、スコアの粒度をスケーリングすると、正の解と負の解がより適切に分離され、より校正された比較が得られることを示します。さらに、繰り返しの評価と基準分解をスケーリングすることにより、分散と複雑さの軽減を通じて検証精度がさらに向上します。さらに、検証者の連続スコアを使用して候補の中から最適なソリューションを選択するための、コスト効率の高いランキング アルゴリズムを導入します。 LLM-as-a-Verifier は Terminal-Bench V2 (86.5%)、SWE-Be で最先端のパフォーマンスを達成

nch Verified (78.2%)、RoboRewardBench (87.4%)、MedAgentBench (73.3%)。検証を超えて、LLM-as-a-Verifier からのきめ細かい信号は、タスクの進行状況を推定するためのプロキシとしても機能します。私たちは Claude Code の拡張機能を構築し、開発者が独自のエージェント システムを監視および改善できるようにします。最後に、LLM-as-a-Verifier が RL に緻密なフィードバックを提供し、ロボット工学と数学的推論のベンチマークにおける SAC と GRPO のサンプル効率を向上させることができることを示します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2607.05391: LLM-as-a-Verifier: A General-Purpose Verification Framework

Skip to main content
arXiv is now an independent nonprofit!
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 6 Jul 2026 ( v1 ), last revised 7 Jul 2026 (this version, v2)]
Title: LLM-as-a-Verifier: A General-Purpose Verification Framework
Abstract: Scaling pre-training, post-training, and test-time compute have become the central paradigms for improving the capabilities of LLMs. In this work, we identify verification, the ability to determine the correctness of a solution, as a new scaling axis. To unlock this and demonstrate its effectiveness, we introduce LLM-as-a-Verifier, a general-purpose verification framework that provides fine-grained feedback for agentic tasks without requiring additional training. Unlike standard LM judges that prompt LLMs to produce discrete scores for candidate solutions, LLM-as-a-Verifier computes the expectation over the distribution of scoring token logits to generate continuous scores. This probabilistic formulation enables verification to scale along multiple dimensions: (1) score granularity, (2) repeated evaluation, and (3) criteria decomposition. In particular, we show that scaling the scoring granularity leads to better separation between positive and negative solutions, resulting in more calibrated comparisons. Moreover, scaling repeated evaluation and criteria decomposition consistently lead to additional gains in verification accuracy through variance and complexity reduction. We further introduce a cost-efficient ranking algorithm for selecting the best solution among candidates using the verifier's continuous scores. LLM-as-a-Verifier achieves state-of-the-art performance on Terminal-Bench V2 (86.5%), SWE-Bench Verified (78.2%), RoboRewardBench (87.4%), and MedAgentBench (73.3%). Beyond verification, the fine-grained signals from LLM-as-a-Verifier can also serve as a proxy for estimating task progress. We build an extension for Claude Code, enabling developers to monitor and improve their own agentic systems. Finally, we show that LLM-as-a-Verifier can provide dense feedback for RL, improving the sample efficiency of SAC and GRPO on robotics and mathematical reasoning benchmarks.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
