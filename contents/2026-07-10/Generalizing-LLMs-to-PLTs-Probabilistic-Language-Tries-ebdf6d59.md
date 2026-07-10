---
source: "https://arxiv.org/abs/2604.06228"
hn_url: "https://news.ycombinator.com/item?id=48861851"
title: "Generalizing LLMs to PLTs: Probabilistic Language Tries"
article_title: "[2604.06228] Probabilistic Language Tries: A Unified Framework for Compression, Decision Policies, and Execution Reuse"
author: "EGreg"
captured_at: "2026-07-10T16:14:31Z"
capture_tool: "hn-digest"
hn_id: 48861851
score: 1
comments: 0
posted_at: "2026-07-10T16:07:47Z"
tags:
  - hacker-news
  - translated
---

# Generalizing LLMs to PLTs: Probabilistic Language Tries

- HN: [48861851](https://news.ycombinator.com/item?id=48861851)
- Source: [arxiv.org](https://arxiv.org/abs/2604.06228)
- Score: 1
- Comments: 0
- Posted: 2026-07-10T16:07:47Z

## Translation

タイトル: LLM から PLT への一般化: 確率的言語の試行
記事のタイトル: [2604.06228] 確率的言語の試行: 圧縮、決定ポリシー、および実行の再利用のための統合フレームワーク
説明: arXiv 論文 2604.06228 の要約ページ: Probabilistic Language Tryes: A Unified Framework for Compression, Decision Policies, and Execution Reuse

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
コンピューターサイエンス > 機械学習
[2026 年 3 月 29 日に提出]
タイトル: 確率的言語の試行: 圧縮、意思決定ポリシー、および実行の再利用のための統一フレームワーク
要約: シーケンスに対する生成モデルによって暗黙的に定義されたプレフィックス構造を明示する統一表現である確率的言語トライ (PLT) を導入します。各出力エッジに対応するトークンまたはアクションの条件付き確率を割り当てることにより、PLT は同時に次の役割を果たします。(i) 周波数重み付け間隔エンコーディングを介した最適な可逆圧縮器。算術コーディングをモデル条件付き分布に一般化します。 (ii) ゲーム、検索、ロボット制御などの逐次決定問題のポリシー表現。 (iii) 繰り返される推論クエリに対して、完全なモデルの実行ではなく構造化された検索によって回答できるようにするメモ化インデックス。
中心的な技術的成果は、事前ガイド付きキャッシュ定理です。定常生成分布の下では、PLT ガイド付きキャッシュは、事前の集中に伴って増加するしきい値を下回るすべてのクエリ数について、経験的頻度キャッシュよりも厳密に低い予想推論コストを達成します。これにより、O(n^2) トランスフォーマ アテンション コストが p_r * O(log N) + (1 - p_r) * O(n^2) の期待コストに変換されます。ここで、p_r は事前に推定された再利用確率、N はアーティファクト ストア サイズです。
さらに、あらゆるデータセットを PLT でカバーされるマジョリティとスパース残差ストアに分解するハイブリッド圧縮アーキテクチャを導入し、算術コーディングをコルモゴロフ形式のプログラム表現およびレート歪み理論と結び付けます。私たちはインスタントシア

これは、チェス、Web 検索、ロボット工学、組織ワークフロー、および LLM 推論にわたるフレームワークであり、圧縮、意思決定、および計算の再利用がすべてシーケンス空間上の単一の確率尺度から導出されることを示しています。
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

Abstract page for arXiv paper 2604.06228: Probabilistic Language Tries: A Unified Framework for Compression, Decision Policies, and Execution Reuse

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
Computer Science > Machine Learning
[Submitted on 29 Mar 2026]
Title: Probabilistic Language Tries: A Unified Framework for Compression, Decision Policies, and Execution Reuse
Abstract: We introduce probabilistic language tries (PLTs), a unified representation that makes explicit the prefix structure implicitly defined by any generative model over sequences. By assigning to each outgoing edge the conditional probability of the corresponding token or action, a PLT simultaneously serves as: (i) an optimal lossless compressor via frequency-weighted interval encoding, generalizing arithmetic coding to model-conditioned distributions; (ii) a policy representation for sequential decision problems including games, search, and robotic control; and (iii) a memoization index that lets repeated inference queries be answered by structured retrieval rather than full model execution.
The central technical result is a prior-guided caching theorem: under a stationary generative distribution, a PLT-guided cache achieves strictly lower expected inference cost than any empirical-frequency cache for all query counts below a threshold that grows with the concentration of the prior. This converts O(n^2) transformer attention cost into an expected cost of p_r * O(log N) + (1 - p_r) * O(n^2), where p_r is the prior-estimated reuse probability and N is the artifact store size.
We further introduce a hybrid compression architecture decomposing any dataset into a PLT-covered majority and a sparse residual store, connecting arithmetic coding with Kolmogorov-style program representations and rate-distortion theory. We instantiate the framework across chess, web search, robotics, organizational workflows, and LLM inference, demonstrating that compression, decision making, and computational reuse are all derived from a single probability measure on sequence space.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
