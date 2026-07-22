---
source: "https://arxiv.org/abs/2607.18242"
hn_url: "https://news.ycombinator.com/item?id=49008530"
title: "AI Tool Discovery at Scale: All You Need Is DNS"
article_title: "[2607.18242] AI Tool Discovery at Scale: All You Need is DNS"
author: "Brajeshwar"
captured_at: "2026-07-22T16:16:25Z"
capture_tool: "hn-digest"
hn_id: 49008530
score: 2
comments: 0
posted_at: "2026-07-22T15:38:36Z"
tags:
  - hacker-news
  - translated
---

# AI Tool Discovery at Scale: All You Need Is DNS

- HN: [49008530](https://news.ycombinator.com/item?id=49008530)
- Source: [arxiv.org](https://arxiv.org/abs/2607.18242)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T15:38:36Z

## Translation

タイトル: 大規模な AI ツールの検出: 必要なのは DNS だけです
記事のタイトル: [2607.18242] 大規模な AI ツール検出: 必要なのは DNS だけです
説明: arXiv 論文 2607.18242 の要約ページ: 大規模な AI ツール検出: 必要なのは DNS だけ

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
[2026 年 4 月 19 日に提出]
タイトル: 大規模な AI ツールの検出: 必要なのは DNS だけです
要約: これからの自律型 AI エージェントの時代には、何百万ものツールを操作できる検出メカニズムが必要ですが、既存のソリューションは O(N) の複雑さと集中ガバナンスの下で機能しません。別の脆弱なオーバーレイを構築する代わりに、インターネットの最も回復力のある基盤であるドメイン ネーム システム (DNS) にセマンティック ツールの検出を後付けする根本的なフレームワークである ToolDNS を提案します。 ToolDNS は、機能的意図と組織の信頼を階層型名前空間に埋め込むことにより、高価なセマンティック検索を一連の軽量な O(log N) 名前解決に変換します。部分的に展開された名前、EDNS0 インテント ペイロード、論理サブドメインという、分散型ガバナンスとセマンティック プルーニングを可能にする 3 つのプロトコル準拠の機能強化を導入します。断片化されたツール環境全体でこのアプローチを厳密に評価するために、MCP、A2A、RESTful、およびスキル プロトコルにわたる 33,688 個の実際のツールで構成される大規模な異種ベンチマークを構築してリリースします。このデータセットでは、ToolDNS は最先端の検索精度を実現しながら、クエリごとの検索スペースを 95.26% 削減します。さらに、UDP ネイティブの設計により、HTTP ベースのレジストリと比較して検出遅延が桁違いに短縮されます。私たちの研究は、スケーラブルな AI の相互運用性には、追加のミドルウェアではなく、すでに足元にあるインフラストラクチャをより賢く利用することが必要であることを示しています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arX

ivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2607.18242: AI Tool Discovery at Scale: All You Need is DNS

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 19 Apr 2026]
Title: AI Tool Discovery at Scale: All You Need is DNS
Abstract: The coming era of autonomous AI agents demands a discovery mechanism capable of navigating millions of tools, yet existing solutions buckle under O(N) complexity and centralized governance. Instead of building another fragile overlay, we propose ToolDNS, a radical framework that retrofits semantic tool discovery onto the Internet's most resilient substrate: the Domain Name System (DNS). By embedding functional intent and organizational trust into a hierarchical namespace, ToolDNS transforms an expensive semantic search into a series of lightweight, O(log N) name resolutions. We introduce three protocol-compliant enhancements to enable decentralized governance and semantic pruning: partially unfolded names, EDNS0 intent payloads, and logical subdomains. To rigorously evaluate this approach across the fragmented tooling landscape, we construct and release a large-scale heterogeneous benchmark comprising 33,688 real-world tools spanning MCP, A2A, RESTful, and Skill protocols. On this dataset, ToolDNS slashes the per-query search space by 95.26% while matching state-of-the-art retrieval accuracy. Furthermore, its UDP-native design reduces discovery latency by orders of magnitude compared to HTTP-based registries. Our work demonstrates that scalable AI interoperability requires not more middleware, but a smarter utilization of the infrastructure already beneath our feet.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
