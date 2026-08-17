---
source: "https://arxiv.org/abs/2608.12123"
hn_url: "https://news.ycombinator.com/item?id=49335310"
title: "Bounding GPU Opportunity and Avoiding Host Round Trips in LLM-Agent Control"
article_title: "[2608.12123] Ready Cohorts: Bounding GPU Opportunity and Avoiding Host Round Trips in LLM-Agent Control"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "josefchen"
captured_at: "2026-08-17T18:22:48Z"
capture_tool: "hn-digest"
hn_id: 49335310
score: 1
comments: 0
posted_at: "2026-08-17T18:16:51Z"
tags:
  - hacker-news
  - translated
---

# Bounding GPU Opportunity and Avoiding Host Round Trips in LLM-Agent Control

- HN: [49335310](https://news.ycombinator.com/item?id=49335310)
- Source: [arxiv.org](https://arxiv.org/abs/2608.12123)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T18:16:51Z

## Translation

タイトル: LLM エージェント制御における GPU の機会を制限し、ホストの往復を回避する
記事のタイトル: [2608.12123] Ready Cohorts: GPU の機会を制限し、LLM エージェント制御でのホストの往復を回避する
説明: arXiv 論文 2608.12123 の要約ページ: Ready Cohorts: Bounding GPU Opportunity and Saving Host Round Trips in LLM-Agent Control

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューター サイエンス > 分散、並列、クラスター コンピューティング
[2026 年 8 月 12 日に提出]
タイトル: Ready Cohorts: GPU の機会を制限し、LLM エージェント制御でのホストの往復を回避する
要約: LLM エージェント サービスは、モデル呼び出しとツール呼び出しの間で小さな決定論的な遷移を繰り返し実行します。つまり、結果のルーティング、状態の更新、次の効果の発行です。この制御パスが GPU 実行に十分な同時作業を公開する時期と、GPU が計算したルート決定がデバイス上に残っているときに何が変化するかを尋ねます。固定パーティション シェア F、正確なオフライン シェア P*、ローカル上限 U、およびオンライン達成シェア A を使用して、準備完了コホートの境界を形式化します。ゼロ サービス時間、無制限の容量、および等しい相対的な起動期限の下で、専用の動的プログラムが P* を正確に計算します。 1 つの固定された 851 セッションのパブリック トレース パネルの定常ポアソン リプレイでは、100,000 のターゲット アクティブ セッション、K=256、および 50 ミリ秒の起動期限での主条件は、F=30.19%、P*=43.00%、U=45.85% となります。正確なパッキングは、固定ウィンドウ境界で失われた機会の 81.83% を回復します。結果から派生したルート キーは条件付けプロキシであり、実行可能ファイルの ID の証明ではありません。別のメカニズムの研究により、ホストに 4 バイトを返して再ディスパッチするのではなく、GPU で計算されたバイナリ決定がデバイス上に保持されます。 4 つの名前付き GPU 配置全体で、デバイス常駐パスは 36 の構成すべてで高速です。配置内の行中央値の範囲は 1.19 倍から 2.39 倍です。両方の許容メカニズムにわたって、テストされた 14,557,440 個のバッチ呼び出しはすべて、個別に実装されたホスト オラクルと一致します。ホストの決定を削除しない固定ネストされたデバイス グラフは、60 のすべての構成で速度が低下します。

5 つの配置。これらの研究を総合すると、期限内に実行可能なコホートの供給と観察の配置という、GPU エージェント制御のための 2 つの測定可能なゲートが確立されます。 A、CPU 置き換え、およびサービス レベルの利点を測定するには、結合された有限オンライン ランタイムが必要です。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.12123: Ready Cohorts: Bounding GPU Opportunity and Avoiding Host Round Trips in LLM-Agent Control

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Distributed, Parallel, and Cluster Computing
[Submitted on 12 Aug 2026]
Title: Ready Cohorts: Bounding GPU Opportunity and Avoiding Host Round Trips in LLM-Agent Control
Abstract: LLM-agent services repeatedly execute small deterministic transitions between model and tool calls: route an outcome, update state, and emit the next effect. We ask when this control path exposes enough concurrent work for GPU execution, and what changes when a GPU-computed route decision remains on device. We formalize the ready-cohort boundary using fixed-partition share F, exact offline share P*, local upper bound U, and online achieved share A. Under zero service time, unlimited capacity, and equal relative launch deadlines, a specialized dynamic program computes P* exactly. In a stationary Poisson replay of one pinned 851-session public trace panel, the primary condition at 100,000 target active sessions, K=256, and a 50 ms launch deadline gives F=30.19%, P*=43.00%, and U=45.85%. Exact packing recovers 81.83% of the opportunity lost at fixed window boundaries. The outcome-derived route key is a conditioning proxy, not proof of executable identity. A separate mechanism study keeps a GPU-computed binary decision on device instead of returning four bytes to the host and redispatching. Across four named GPU placements, the device-resident path is faster in all 36 configurations; within-placement row-median ratios range from 1.19x to 2.39x. Across both admissible mechanisms, all 14,557,440 tested batched invocations match a separately implemented host oracle. A fixed nested device graph that removes no host decision is slower in all 60 configurations across five placements. Together, the studies establish two measurable gates for GPU agent control: deadline-feasible cohort supply and observation placement. A joined finite online runtime is required to measure A, CPU displacement, and service-level benefit.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
