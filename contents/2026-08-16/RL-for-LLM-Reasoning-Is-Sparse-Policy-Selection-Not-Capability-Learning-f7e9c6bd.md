---
source: "https://arxiv.org/abs/2605.06241"
hn_url: "https://news.ycombinator.com/item?id=49319908"
title: "RL for LLM Reasoning Is Sparse Policy Selection, Not Capability Learning"
article_title: "[2605.06241] Rethinking RL for LLM Reasoning: It's Sparse Policy Selection, Not Capability Learning"
author: "BlackGlory"
captured_at: "2026-08-16T14:13:23Z"
capture_tool: "hn-digest"
hn_id: 49319908
score: 1
comments: 0
posted_at: "2026-08-16T13:32:41Z"
tags:
  - hacker-news
  - translated
---

# RL for LLM Reasoning Is Sparse Policy Selection, Not Capability Learning

- HN: [49319908](https://news.ycombinator.com/item?id=49319908)
- Source: [arxiv.org](https://arxiv.org/abs/2605.06241)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T13:32:41Z

## Translation

タイトル: LLM 推論の RL は、能力学習ではなく、スパースなポリシー選択です
記事のタイトル: [2605.06241] LLM 推論のための RL の再考: 能力学習ではなく、スパースなポリシー選択です
説明: arXiv 論文 2605.06241 の要約ページ: LLM 推論のための RL の再考: 能力学習ではなく、スパースなポリシー選択です

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 計算と言語
[2026 年 5 月 7 日に提出 ( v1 )、最終改訂日 2026 年 5 月 8 日 (このバージョン、v2)]
タイトル: LLM 推論のための RL の再考: 能力学習ではなく、スパースなポリシー選択です
要約: 強化学習は、大規模な言語モデルにおける推論を改善するための標準となっていますが、RL が新しい戦略を教えていないことを示す証拠が増えています。基本モデルに既に含まれている解全体に確率質量を再分配します。この研究では、RL がモデルをすでに知っているパスに向けて誘導するだけの場合、RL 最適化ループ自体は必要なのでしょうか?複数のモデル ファミリと RL アルゴリズムにわたるトークンレベルの分析を通じて、RL の有益なフットプリントは、モデルがどの分岐を取るべきかが不確実な高エントロピーの決定点に集中する、まばらで予測可能な修正であることがわかりました。トークン位置の 1 ～ 3\% のみが影響を受け、プロモートされたトークンは常にベース モデルの上位 5 つの選択肢内にあり、これらの少数の位置でのターゲットを絞った修正は因果的に RL の精度向上の大部分を回復しますが、ランダムな修正は失敗します。基本モデル自体のエントロピーは、RL でトレーニングされたモデルを使用せずにこれらの位置を特定し、補正全体は低次元であり、モデル パラメーターのごく一部で表現できます。これらの発見は、推論の改善を、能力の獲得ではなく、まばらな政策の選択として再構成します。私たちは、この洞察を ReasonMaxxer に変換します。ReasonMaxxer は、数百のベースモデルのロールアウトを使用し、オンライン生成を使用せず、エントロピー ゲートの決定点でのみ対比損失を適用する最小限の RL フリー手法です。 3 つのモデル ファミリ、6 つのスケール、6 つの数的推論ベンチマークにまたがる

, ReasonMaxxer は、完全な RL パフォーマンスと同等かそれを上回っていますが、必要な問題数は数十、単一 GPU トレーニングのみで、トレーニング コストがおよそ 3 桁削減されます。
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

Abstract page for arXiv paper 2605.06241: Rethinking RL for LLM Reasoning: It's Sparse Policy Selection, Not Capability Learning

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 7 May 2026 ( v1 ), last revised 8 May 2026 (this version, v2)]
Title: Rethinking RL for LLM Reasoning: It's Sparse Policy Selection, Not Capability Learning
Abstract: Reinforcement learning has become the standard for improving reasoning in large language models, yet evidence increasingly suggests that RL does not teach new strategies; it redistributes probability mass over solutions the base model already contains. In this work, we ask: if RL merely steers the model toward paths it already knows, is the RL optimization loop itself necessary? Through token-level analysis across multiple model families and RL algorithms, we find that RL's beneficial footprint is a sparse, predictable correction concentrated at high-entropy decision points where the model is uncertain which branch to take. Only 1--3\% of token positions are affected, the promoted token always lies within the base model's top-5 alternatives, and targeted corrections at those few positions causally recover a large fraction of RL's accuracy gain, while random corrections fail. The base model's own entropy identifies these positions without any RL-trained model, and the entire correction is low-dimensional, representable in a tiny fraction of model parameters. These findings reframe reasoning improvement as sparse policy selection, not capability acquisition. We translate this insight into ReasonMaxxer, a minimal RL-free method that applies contrastive loss only at entropy-gated decision points, using a few hundred base-model rollouts and no online generation. Across three model families, six scales, and six math reasoning benchmarks, ReasonMaxxer matches or exceeds full RL performance while requiring only tens of problems and minutes of single-GPU training, a reduction in training cost of roughly three orders of magnitude.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
