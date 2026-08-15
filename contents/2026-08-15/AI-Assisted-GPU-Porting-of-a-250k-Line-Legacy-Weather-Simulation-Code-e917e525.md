---
source: "https://arxiv.org/abs/2608.13122"
hn_url: "https://news.ycombinator.com/item?id=49314967"
title: "AI-Assisted GPU Porting of a 250k Line Legacy Weather Simulation Code"
article_title: "[2608.13122] Validation-Centric AI-Assisted GPU Porting of a 250,000+ Line Legacy Weather Simulation Code"
author: "Jimmc414"
captured_at: "2026-08-15T23:11:00Z"
capture_tool: "hn-digest"
hn_id: 49314967
score: 1
comments: 0
posted_at: "2026-08-15T22:41:51Z"
tags:
  - hacker-news
  - translated
---

# AI-Assisted GPU Porting of a 250k Line Legacy Weather Simulation Code

- HN: [49314967](https://news.ycombinator.com/item?id=49314967)
- Source: [arxiv.org](https://arxiv.org/abs/2608.13122)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T22:41:51Z

## Translation

タイトル: AI 支援による 250k ラインのレガシー気象シミュレーション コードの GPU 移植
記事のタイトル: [2608.13122] 250,000 行を超えるレガシー気象シミュレーション コードの検証中心の AI 支援 GPU 移植
説明: arXiv 論文 2608.13122 の要約ページ: 250,000 行を超えるレガシー気象シミュレーション コードの検証中心の AI 支援 GPU 移植

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
[2026 年 8 月 13 日に提出]
タイトル: 検証中心の AI 支援による 250,000 行を超えるレガシー気象シミュレーション コードの GPU 移植
要約: 大規模言語モデルの最近の進歩により、CLI ベースの AI エージェントは、大規模なレガシー科学アプリケーションの GPU 移植を加速するための実用的なツールになりました。ただし、そのようなアプリケーションは単なる古いコードベースではありません。これらは、長期にわたる開発、観察との比較、および分野研究での使用を通じて信頼性が蓄積された科学資産です。したがって、GPU 移植では、実装を GPU 中心の HPC システムに適応させながら、この科学的妥当性を維持する必要があります。このペーパーでは、250,000 行を超えるレガシー Fortran 気象シミュレーション コードである CReSS のケース スタディを通じて、検証中心の AI 支援 GPU 移植ワークフローを紹介します。このワークフローでは、AI エージェントを使用して OpenMP 領域を抽出し、物理的に意味のあるシミュレーション状態からダンプベースのカーネル ベンチマークを生成し、OpenACC 変換を適用し、ダンプされた参照データとの要素ごとの比較およびアプリケーション レベルの検証を通じて結果を検証します。実際の台風シミュレーションを使用して、ワークフローは 162 のターゲット カーネルに対して数値的に検証された GPU 実装を生成し、実用的な開発コスト内でアプリケーション レベルの 5.1 倍の高速化を達成しました。特に、しきい値に敏感な分岐分岐やキャンセル効果など、浮動小数点と組み込み関数の違いによって引き起こされる 5 つのカーネルの数値の不一致を検出し、アプリケーション開発者へのフィードバックを可能にしました。このケーススタディは、大規模なレガシー科学については、

ダンプベースの検証を必要とするアプリケーションでは、実際の AI 支援 GPU ポーティングでは、セッションにまたがるコンテキスト、ランタイム状態の再構築、および小さな静的解析の欠落からのコストのかかるリカバリを管理する必要があります。これらの調査結果は、AI 支援による GPU 移植にはコード生成だけでなく、検証中心のワークフロー設計も必要であることを示しています。
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

Abstract page for arXiv paper 2608.13122: Validation-Centric AI-Assisted GPU Porting of a 250,000+ Line Legacy Weather Simulation Code

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Distributed, Parallel, and Cluster Computing
[Submitted on 13 Aug 2026]
Title: Validation-Centric AI-Assisted GPU Porting of a 250,000+ Line Legacy Weather Simulation Code
Abstract: Recent advances in large language models have made CLI-based AI agents a practical tool for accelerating GPU porting of large legacy scientific applications. Such applications, however, are not merely old code bases; they are scientific assets whose credibility has been accumulated through long-term development, comparison with observations, and use in domain studies. GPU porting must therefore preserve this scientific validity while adapting the implementation to GPU-centric HPC systems. This paper presents a validation-centric AI-assisted GPU porting workflow through a case study of CReSS, a legacy Fortran weather simulation code with more than 250,000 lines. The workflow uses an AI agent to extract OpenMP regions, generate dump-based kernel benchmarks from physically meaningful simulation states, apply OpenACC transformations, and validate results through element-wise comparison with dumped reference data and application-level validation. Using a real typhoon simulation, the workflow produced numerically validated GPU implementations for 162 target kernels and achieved a 5.1x application-level speedup within practical wall-clock development cost. In particular, it detected numerical discrepancies in five kernels caused by floating-point and intrinsic-function differences, including threshold-sensitive branch divergence and cancellation effects, enabling feedback to the application developers. The case study suggests that, for large legacy scientific applications requiring dump-based validation, practical AI-assisted GPU porting must manage session-spanning context, runtime-state reconstruction, and costly recovery from small static-analysis omissions. These findings demonstrate that AI-assisted GPU porting requires not only code generation, but validation-centric workflow design.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
