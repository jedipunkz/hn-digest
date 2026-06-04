---
source: "https://arxiv.org/abs/2605.11325"
hn_url: "https://news.ycombinator.com/item?id=48396376"
title: "LLM memory systems benchmark: high recall near-zero precision for tested systems"
article_title: "[2605.11325] Structured Belief State and the First Precision-Aware Benchmark for LLM Memory Retrieval"
author: "decorner"
captured_at: "2026-06-04T10:21:51Z"
capture_tool: "hn-digest"
hn_id: 48396376
score: 2
comments: 0
posted_at: "2026-06-04T09:57:25Z"
tags:
  - hacker-news
  - translated
---

# LLM memory systems benchmark: high recall near-zero precision for tested systems

- HN: [48396376](https://news.ycombinator.com/item?id=48396376)
- Source: [arxiv.org](https://arxiv.org/abs/2605.11325)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T09:57:25Z

## Translation

タイトル: LLM メモリ システム ベンチマーク: テスト済みシステムのゼロに近い高リコール精度
記事のタイトル: [2605.11325] 構造化された信念状態と LLM メモリ取得のための最初の精度を意識したベンチマーク
説明: arXiv 論文 2605.11325 の要約ページ: Structured Belief State and the First Precision-Aware Benchmark for LLM Memory Retrieval

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2605.11325
ヘルプ |高度な検索
コンピュータサイエンス > 情報検索
[2026 年 5 月 11 日に提出 ( v1 )、最終改訂日 2026 年 5 月 27 日 (このバージョン、v2)]
タイトル: 構造化された信念状態と LLM メモリ取得のための最初の精度を意識したベンチマーク
要約: LLM メモリ システムの主要なベンチマークはすべて、LoCoMo を筆頭に、メモリ システムが正しく取得したかどうかではなく、モデルが正しく応答したかどうかを測定します。信念ストア全体を返すシステムは、再現率 1.0 を達成し、回答品質評価に合格します。これが単体テストと統合テストの違いです。取得品質は、それが入力される生成モデルから切り離して測定する必要があり、これを行う既存のベンチマークはありません。
エンティティの抽出が完全に忠実である場合でも、この失敗が続くことを示します。メモリ ベースラインは、独自の抽出を参照するケースでわずか 0.05 ～ 0.08 の平均検索精度を達成します。この失敗は構造的なものです。ドメイン固有のコーパスにおけるコサイン類似性では、関連する信念と意味的に近い信念を区別できず、埋め込みモデルのスケールで 20 倍の範囲にわたって不変性が確認されています。複数回の評価を行うと、複合的な失敗が表面化します。トピックのドリフト後、比較システムにより意味の質量がターンを超えて流出することが可能になり、再エントリ時に高いドリフト スコアが得られます。シングルターンのメトリクスはこのコストを隠します。Hindsight はシングルターンのレイテンシが 700 ミリ秒未満であると報告していますが、セッション ターンあたりの平均は 2,700 ミリ秒を超えており、p95 は 6,000 ミリ秒を超えています。 LLM-as-a-Judge の評価では、これらの失敗は目に見えないままになります。
私たちは 2 つの貢献を紹介します: PrecisionMemBench、

多様なスコープ、突然変異、および分離アサーションにわたる生成モデルとは独立して検索精度を測定する 89 ケースのベンチマーク。 Tenure は、アナライザーの非対称性、差動ブースティング、およびハード スコープ分離を備えたマルチパス BM25 を使用する、ローカル ファーストの構造化ビリーフ ストアです。 Tenure は、平均精度 1.0 および 15 ミリ秒未満の取得遅延で 89/89 のケースを合格しました。比較プロバイダーは、構築された生のベクトル ベースラインよりもパフォーマンスが悪く、アクティブな取得パスがゼロで、取り込みコストが 98 ～ 897 秒かかるため、応答品質ベンチマークでは検出できない障害が発生します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2605.11325: Structured Belief State and the First Precision-Aware Benchmark for LLM Memory Retrieval

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2605.11325
Help | Advanced Search
Computer Science > Information Retrieval
[Submitted on 11 May 2026 ( v1 ), last revised 27 May 2026 (this version, v2)]
Title: Structured Belief State and the First Precision-Aware Benchmark for LLM Memory Retrieval
Abstract: Every major benchmark for LLM memory systems, LoCoMo foremost, measures whether a model answered correctly, not whether the memory system retrieved correctly. A system returning its entire belief store achieves recall of 1.0 and passes answer-quality evaluation. This is the difference between a unit test and an integration test: retrieval quality must be measured in isolation from the generative model it feeds into, and no existing benchmark does this.
We demonstrate that this failure persists even when entity extraction is entirely faithful. Memory baselines achieve mean retrieval precision of just 0.05 to 0.08 on cases referencing their own extractions. The failure is structural: cosine similarity over a domain-specific corpus cannot discriminate relevant beliefs from semantically proximate ones, an invariance confirmed across a 20x range in embedding model scale. Multi-turn evaluation surfaces a compounding failure; after topic drift, comparison systems allow semantic mass to bleed across turns, yielding high drift scores on re-entry. Single-turn metrics conceal this cost: Hindsight reports sub-700ms single-turn latency but exceeds 2,700ms mean per session turn, with p95 above 6,000ms. Under LLM-as-a-Judge evaluation, these failures remain invisible.
We present two contributions: PrecisionMemBench, an 89-case benchmark measuring retrieval precision independently of generative models across diverse scope, mutation, and isolation assertions; and Tenure, a local-first structured belief store using multi-path BM25 with analyzer asymmetry, differential boosting, and hard scope isolation. Tenure passes 89/89 cases with mean precision 1.0 and sub-15ms retrieval latency. Comparison providers perform worse than the raw vector baseline they are built on, with zero active retrieval passes and ingestion costs of 98 to 897 seconds, failures that answer-quality benchmarks cannot detect.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
contact arXiv Click here to contact arXiv
Contact
subscribe to arXiv mailings Click here to subscribe
Subscribe
