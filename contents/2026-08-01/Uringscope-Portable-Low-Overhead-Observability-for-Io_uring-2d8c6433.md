---
source: "https://arxiv.org/abs/2606.15137"
hn_url: "https://news.ycombinator.com/item?id=49137062"
title: "Uringscope: Portable, Low-Overhead Observability for Io_uring"
article_title: "[2606.15137] uringscope: Portable, Low-Overhead Observability for io_uring"
author: "Jimmc414"
captured_at: "2026-08-01T18:54:35Z"
capture_tool: "hn-digest"
hn_id: 49137062
score: 2
comments: 1
posted_at: "2026-08-01T18:32:29Z"
tags:
  - hacker-news
  - translated
---

# Uringscope: Portable, Low-Overhead Observability for Io_uring

- HN: [49137062](https://news.ycombinator.com/item?id=49137062)
- Source: [arxiv.org](https://arxiv.org/abs/2606.15137)
- Score: 2
- Comments: 1
- Posted: 2026-08-01T18:32:29Z

## Translation

タイトル: Uringscope: Io_uring のためのポータブルで低オーバーヘッドの可観測性
記事のタイトル: [2606.15137] uringscope: io_uring のポータブル、低オーバーヘッドの可観測性
説明: arXiv 論文 2606.15137 の要約ページ: uringscope: io_uring のポータブルで低オーバーヘッドの可観測性

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > オペレーティングシステム
[2026 年 6 月 13 日に提出 ( v1 )、最終改訂日 2026 年 7 月 9 日 (このバージョン、v2)]
タイトル: uringscope: io_uring のポータブルで低オーバーヘッドの可観測性
要約: io_uring は、I/O の送信と完了を共有メモリ リングに移動します。これにより、処理が高速になるだけでなく、非表示になります。 strace はリング設定のみを参照し、要求フローを公開するカーネル トレースポイントは安定した ABI ではないため、それらに構築されたいくつかのツールは狭いカーネル範囲でのみ機能します。 CO-RE (Compile Once, Run Everywhere) eBPF 上に構築された io_uring 用の単一バイナリで言語に依存しない可観測性ツール、uringscope を紹介します。 uringscope は 4 つの貢献を行っています。 1 つ目は、リクエストのライフサイクルの正確なモデルと、カーネル イベントからリクエストごとのフローを再構築する方法です。 2 つ目は、BTF プローブされたプログラム バリアント、CO-RE フィールド フレーバー、および位置に依存しない読み取りを使用して、不安定なトレースポイント サーフェスに移植可能にアタッチする手法です。 3 つ目は、オーバーヘッドと忠実度の間のトレードオフの評価です。デバイスにバインドされた NVMe ワークロードでは、uringscope の集約モードのコストはスループットの 0.7 ～ 9.9% であり、これは、測定したすべての完全忠実度の代替案よりも安価です。 4 つ目は、ヒストグラムを参照するのではなくテール レイテンシー インシデントをデバッグするオペレーター向けに、同じ再構築を再利用して送信境界ハザードを検出する軽量の正確性モードです。また、測定結果を証拠のある名前付き病理に変換する内蔵ドクターと併用します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: 実験的

コミュニティの協力者とのプロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2606.15137: uringscope: Portable, Low-Overhead Observability for io_uring

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Operating Systems
[Submitted on 13 Jun 2026 ( v1 ), last revised 9 Jul 2026 (this version, v2)]
Title: uringscope: Portable, Low-Overhead Observability for io_uring
Abstract: io_uring moves I/O submission and completion into shared-memory rings. This makes it fast, and it also makes it invisible. strace sees only the ring setup, and the kernel tracepoints that expose the request flow are not stable ABI, so the few tools built on them work only on narrow kernel ranges. We present uringscope, a single-binary, language-agnostic observability tool for io_uring built on CO-RE (Compile Once, Run Everywhere) eBPF. uringscope makes four contributions. The first is a precise model of the request lifecycle and a method to reconstruct per-request flows from kernel events. The second is a technique for attaching portably to an unstable tracepoint surface, using BTF-probed program variants, CO-RE field flavors, and position-independent reads. The third is an evaluation of the tradeoff between overhead and fidelity: on device-bound NVMe workloads uringscope's aggregate mode costs 0.7 to 9.9% of throughput, which is cheaper than every full-fidelity alternative we measured. The fourth is a lightweight correctness mode that reuses the same reconstruction to detect submission-boundary hazards, together with a built-in doctor that turns the measurements into named pathologies with evidence, for operators who are debugging a tail-latency incident rather than browsing histograms.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
