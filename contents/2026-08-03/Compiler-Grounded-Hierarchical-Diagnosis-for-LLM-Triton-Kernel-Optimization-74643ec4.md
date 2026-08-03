---
source: "https://arxiv.org/abs/2607.23089"
hn_url: "https://news.ycombinator.com/item?id=49162463"
title: "Compiler-Grounded Hierarchical Diagnosis for LLM Triton Kernel Optimization"
article_title: "[2607.23089] Compiler-Grounded Hierarchical Diagnosis for LLM-Based Triton Kernel Optimization"
author: "matt_d"
captured_at: "2026-08-03T22:58:28Z"
capture_tool: "hn-digest"
hn_id: 49162463
score: 1
comments: 0
posted_at: "2026-08-03T22:53:27Z"
tags:
  - hacker-news
  - translated
---

# Compiler-Grounded Hierarchical Diagnosis for LLM Triton Kernel Optimization

- HN: [49162463](https://news.ycombinator.com/item?id=49162463)
- Source: [arxiv.org](https://arxiv.org/abs/2607.23089)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T22:53:27Z

## Translation

タイトル: LLM Triton カーネル最適化のためのコンパイラベースの階層診断
記事のタイトル: [2607.23089] LLM ベースの Triton カーネル最適化のためのコンパイラベースの階層診断
説明: arXiv 論文 2607.23089: LLM ベースの Triton カーネル最適化のためのコンパイラベースの階層診断の要約ページ

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 人工知能
[2026 年 7 月 25 日に提出]
タイトル: LLM ベースの Triton カーネル最適化のためのコンパイラベースの階層診断
要約: 大規模言語モデル (LLM) の最近の進歩により、カーネルの自動生成と最適化が可能になりましたが、既存のアプローチのほとんどはコンパイル フィードバックやプロファイリング メトリクスなどの表面的なシグナルに依存しています。これらのシグナルは、カーネルが遅いことを明らかにしますが、バックエンド コンパイラーが、特に NPU などの新興アクセラレータで有益な最適化を実現できない理由は明らかにしません。したがって、ソースを書き換える前に、実行時の症状を IR 構造およびコンパイラの動作に結び付ける、漸進的なクロスレイヤ診断問題としてカーネルの最適化を定式化します。この洞察に基づいて、コンパイラベースの Triton カーネル用の階層最適化フレームワークである私たちのシステムを紹介します。このシステムは、より深い証拠が必要な場合にのみ、軽量パターンのトリアージとプロファイリング診断から IR 帰属とコンパイラに基づいた分析にエスカレートし、その後、証拠に裏付けられたソースレベルの書き換えを提案します。
このシステムを Ascend NPU 用の Triton に実装し、標準化された NPUKernelBench 由来の Ascend 950 ベンチマークから正常に変換された 37 個のエントリで評価しました。これらのエントリ全体で、システムは初期の Triton カーネルから最適化された Triton カーネルまでに、幾何平均 4.35$\times$ の高速化と中央値 2.73$\times$ の高速化を達成しました。 22/37 は 2$\times$ を超え、13/37 は 5$\times$ を超えています。完全な分布はベースラインに近いエントリーから大規模な勝利まで多岐にわたり、現在のシステムの範囲と制限に関する透明性のあるレポートを動機付けます。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)

ション)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2607.23089: Compiler-Grounded Hierarchical Diagnosis for LLM-Based Triton Kernel Optimization

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 25 Jul 2026]
Title: Compiler-Grounded Hierarchical Diagnosis for LLM-Based Triton Kernel Optimization
Abstract: Recent advances in large language models (LLMs) have enabled automated kernel generation and optimization, but most existing approaches rely on surface signals such as compilation feedback and profiling metrics. These signals reveal that a kernel is slow, but not why the backend compiler fails to realize a profitable optimization, especially on emerging accelerators such as NPUs. We therefore formulate kernel optimization as a progressive cross-layer diagnosis problem that links runtime symptoms to IR structure and compiler behavior before rewriting source. Based on this insight, we present our system, a compiler-grounded and hierarchical optimization framework for Triton kernels. the system escalates from lightweight pattern triage and profiling diagnosis to IR attribution and compiler-grounded analysis only when deeper evidence is needed, then proposes evidence-backed source-level rewrites.
We implement the system on Triton for Ascend NPUs and evaluate it on 37 successfully converted entries from a standardized NPUKernelBench-derived Ascend 950 benchmark. Across these entries, the system attains a geometric-mean speedup of 4.35$\times$ and a median speedup of 2.73$\times$ from the initial to optimized Triton kernel; 22/37 exceed 2$\times$ and 13/37 exceed 5$\times$. The complete distribution ranges from near-baseline entries to large wins, motivating transparent reporting of the current system's scope and limitations.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
