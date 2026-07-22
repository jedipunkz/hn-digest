---
source: "https://arxiv.org/abs/2607.18069"
hn_url: "https://news.ycombinator.com/item?id=49000502"
title: "Hardware Mechanisms to Dynamically Throttle AI Performance"
article_title: "[2607.18069] Hardware Mechanisms to Dynamically Throttle AI Performance"
author: "Jimmc414"
captured_at: "2026-07-22T01:44:50Z"
capture_tool: "hn-digest"
hn_id: 49000502
score: 2
comments: 0
posted_at: "2026-07-22T01:01:18Z"
tags:
  - hacker-news
  - translated
---

# Hardware Mechanisms to Dynamically Throttle AI Performance

- HN: [49000502](https://news.ycombinator.com/item?id=49000502)
- Source: [arxiv.org](https://arxiv.org/abs/2607.18069)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T01:01:18Z

## Translation

タイトル: AI パフォーマンスを動的に調整するハードウェア メカニズム
記事のタイトル: [2607.18069] AI パフォーマンスを動的に調整するハードウェア メカニズム
説明: arXiv 論文 2607.18069: AI パフォーマンスを動的に調整するハードウェア メカニズムの要約ページ

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータ サイエンス > ハードウェア アーキテクチャ
[2026 年 7 月 20 日に提出]
タイトル: AI パフォーマンスを動的に調整するハードウェア メカニズム
要約: より高性能な AI モデルが重要なコンピューター システムに統合されるようになるにつれて、AI の意図を制御できないことが安全メカニズムの動機となります。既存のソフトウェア保護手段は、十分にインテリジェントなモデルによって潜在的に回避できる動作上の制約のみを課します。ハードウェア レベルの安全性の強制は重要な最後の防御線として認識されていますが、不正アクセスや粗雑なフルチップ シャットダウンに関するポリシー規制を超えるメカニズムはほとんど提案されていません。欠けているのは、アーキテクチャ レベルでのきめの細かい動的な介入メカニズムです。
このペーパーでは、利用可能なハードウェア リソースを動的に制御して実行時の AI パフォーマンスを制限する一連のマイクロアーキテクチャ ノブを紹介します。 GPU メモリ サブシステムの容量、帯域幅、レイテンシ、周波数の各次元にわたる候補ノブを評価し、L2 サイズ、L2 レイテンシ、L2 帯域幅、共有メモリ ポート アクセス レートの 4 つの有力な候補に絞り込みます。新しいロジックと追加の設計コストを最小限に抑えるために、キャッシュ ウェイ マスキング、クレジット ベースのレート制限、レイテンシ挿入、およびバンク アービトレーションという 4 つのメカニズムすべてを、確立されたマイクロアーキテクチャ プリミティブから構築します。これらのノブは、高いパフォーマンス感度 (1/8 リソース可用性で最大 80% のパフォーマンス削減)、ごくわずかな実装コスト (<~10,000 フリップフロップ)、動的スロットリング後の高速安定化 (5 ～ 80,000 サイクル)、およびチップの残りの部分への副次的な影響を最小限に抑えていることを示します。さらに、マルチノブ分析により組み合わせが明らかになる

各ノブの個別の効果を超えてパフォーマンスの低下を増幅するノブの数により、より幅広いパフォーマンス目標が可能になります。
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

Abstract page for arXiv paper 2607.18069: Hardware Mechanisms to Dynamically Throttle AI Performance

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Hardware Architecture
[Submitted on 20 Jul 2026]
Title: Hardware Mechanisms to Dynamically Throttle AI Performance
Abstract: As more capable AI models are increasingly integrated into critical computer systems, the lack of control over AI intent motivates safety mechanisms. Existing software safeguards impose only behavioral constraints that can potentially be bypassed by sufficiently intelligent models. While hardware-level safety enforcement has been recognized as an essential last line of defense, few mechanisms have been proposed beyond policy regulations on unauthorized accesses or coarse full-chip shutdown. What is missing is a fine-grained, dynamic intervention mechanism at the architecture level.
In this paper, we introduce a set of microarchitecture knobs which dynamically control the available hardware resources to limit AI performance at runtime. We evaluate candidate knobs spanning the GPU memory subsystem, across capacity, bandwidth, latency and frequency dimensions, and narrow down to four strong candidates: L2 size, L2 latency, L2 bandwidth, and shared memory port access rate. To minimize new logic and extra design cost, we build all four mechanisms from well-established microarchitectural primitives: cache way masking, credit-based rate limiting, latency insertion, and bank arbitration. We show that these knobs achieve high performance sensitivity (up to 80% performance cut at 1/8 resource availability), negligible implementation cost (<~10K flip flops), fast stabilization after dynamic throttling (5-80K cycles), and minimal collateral impact on the rest of the chip. Further, multi-knob analysis reveals combinations of knobs that amplify the performance degradation beyond the effect of each knob individually, which enables a broader range of performance targets.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
