---
source: "https://arxiv.org/abs/2608.13404"
hn_url: "https://news.ycombinator.com/item?id=49323197"
title: "Does Fixing Break Security? An Empirical Study of LLM Security Degradation"
article_title: "[2608.13404] Does Fixing Break Security? An Empirical Study of Security Degradation in Iterative LLM-Driven Infrastructure-as-Code Repair"
author: "Jimmc414"
captured_at: "2026-08-16T20:10:39Z"
capture_tool: "hn-digest"
hn_id: 49323197
score: 1
comments: 0
posted_at: "2026-08-16T20:09:22Z"
tags:
  - hacker-news
  - translated
---

# Does Fixing Break Security? An Empirical Study of LLM Security Degradation

- HN: [49323197](https://news.ycombinator.com/item?id=49323197)
- Source: [arxiv.org](https://arxiv.org/abs/2608.13404)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T20:09:22Z

## Translation

タイトル: 修正するとセキュリティが破壊されますか? LLM のセキュリティ低下に関する実証的研究
記事のタイトル: [2608.13404] 修正するとセキュリティが損なわれますか?反復的な LLM 駆動の Infrastructure-as-Code 修復におけるセキュリティ低下の実証的研究
説明: arXiv 論文 2608.13404 の要約ページ: 修正はセキュリティを破壊しますか?反復的な LLM 駆動の Infrastructure-as-Code 修復におけるセキュリティ低下の実証的研究

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > ソフトウェアエンジニアリング
[2026 年 8 月 13 日に提出]
タイトル: 修正するとセキュリティが破壊されますか?反復的な LLM 駆動の Infrastructure-as-Code 修復におけるセキュリティ低下の実証的研究
要約: 背景: 反復フィードバック ループは、LLM で生成されたコードとしてのインフラストラクチャ (IaC) を改善するための主要なパラダイムです。Checkov や terraform などの検証ツールは、連続した修復試行に対してフィードバック エラー信号を検証します。以前の研究では、構築によって減少しない累積最良のメトリクスが報告されているため、生の反復ごとのセキュリティ軌道が IaC について検査されたことはありません。目的: セキュリティ回帰 (以前は合格していた CIS ベンチマーク チェックが、修復の反復後に失敗すること) を研究し、反復的な LLM 修復によって他の問題を修正しながらセキュリティが低下するかどうか、またその頻度がどれくらいかを判断します。方法: IaC-Eval ベンチマークからの 5,968 のシナリオ タイムラインを分析します。各 1 つのシナリオは 1 つの構成で最大 5 回の修復反復で実行されます。 15 の構成 (モデル固有の 6 つの RAG、モデルに集約された非 RAG 9 つ、それぞれ 3 つの温度) では、両側に Checkov データを含む 4,440 回の反復遷移が生成されます。標準 (包括的) と厳密 (排他的チェック失敗のみ) の 2 つの検出モードで、30 個の個別の CIS チェック ID を追跡し、コードの差分から根本原因を分類します。結果: 標準的な検出では、シナリオの 13.8% (遷移の 24.8%) で少なくとも 1 つの回帰が見られました。厳密に検出された場合、この割合はシナリオの 3.3% (遷移の 5.2%) に低下し、明らかな回帰のほとんどがマルチリソースの測定アーティファクトであることを示しています。リソースの再構築 (79.0%) が主な根本原因です。回帰推移ではタラの数が 2.6 倍であることが示されています

チャーン (コーエンの d=0.90) と 4.9 倍高い厳密モード チェックのボラティリティ (d=1.49)。標準モードの回帰のうち、36.6% が平均 1.2 回の反復内で自己修正しました。反復 3 が最適な停止点です。結論: 反復的な IaC 修復によりセキュリティの後退が発生しますが、保守的で防御可能な割合はシナリオの約 3.3% です。私たちの調査結果は、セキュリティを意識したフィードバック ループの設計と実用的なイテレーション予算のガイダンスの動機付けとなります。
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

Abstract page for arXiv paper 2608.13404: Does Fixing Break Security? An Empirical Study of Security Degradation in Iterative LLM-Driven Infrastructure-as-Code Repair

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Software Engineering
[Submitted on 13 Aug 2026]
Title: Does Fixing Break Security? An Empirical Study of Security Degradation in Iterative LLM-Driven Infrastructure-as-Code Repair
Abstract: Background: Iterative feedback loops are the dominant paradigm for improving LLM-generated Infrastructure-as-Code (IaC): validators such as Checkov and terraform validate feed error signals back for successive repair attempts. Prior work reports cumulative-best metrics, which are non-decreasing by construction, so the raw per-iteration security trajectory has never been examined for IaC. Aims: We study security regression (a previously-passing CIS Benchmark check that fails after a repair iteration) to determine whether and how often iterative LLM repair degrades security while fixing other issues. Method: We analyze 5,968 scenario timelines from the IaC-Eval benchmark, each one scenario run through one configuration for up to 5 repair iterations. The 15 configurations (six model-specific RAG, nine model-aggregated non-RAG, three temperatures each) yield 4,440 iteration transitions with Checkov data on both sides. We track 30 individual CIS check IDs and classify root causes from code diffs, under two detection modes: standard (inclusive) and strict (exclusive check failures only). Results: Under standard detection, 13.8% of scenarios (24.8% of transitions) exhibit at least one regression. Under strict detection the rate falls to 3.3% of scenarios (5.2% of transitions), indicating most apparent regressions are multi-resource measurement artifacts. Resource restructuring (79.0%) is the dominant root cause. Regression transitions show 2.6x more code churn (Cohen's d=0.90) and 4.9x higher strict-mode check volatility (d=1.49). Of standard-mode regressions, 36.6% self-correct within an average of 1.2 iterations; iteration 3 is the optimal stopping point. Conclusions: Iterative IaC repair does introduce security regressions, but the conservative, defensible rate is about 3.3% of scenarios. Our findings motivate security-aware feedback-loop design and actionable iteration-budget guidance.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
