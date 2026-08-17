---
source: "https://zenodo.org/records/21696066"
hn_url: "https://news.ycombinator.com/item?id=49335495"
title: "AI alignment as continuation control: 31,430 frozen trials"
article_title: "Cross-Vendor Semantic Void Matrix | Zenodo"
image: ""
author: "rayanpal_"
captured_at: "2026-08-17T19:21:39Z"
capture_tool: "hn-digest"
hn_id: 49335495
score: 2
comments: 0
posted_at: "2026-08-17T18:27:07Z"
tags:
  - hacker-news
  - translated
---

# AI alignment as continuation control: 31,430 frozen trials

- HN: [49335495](https://news.ycombinator.com/item?id=49335495)
- Source: [zenodo.org](https://zenodo.org/records/21696066)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T18:27:07Z

## Translation

タイトル: 継続対照としての AI アライメント: 31,430 件の凍結試験
記事のタイトル: クロスベンダー セマンティック ボイド マトリックス |ゼノド
説明: このプレプリントは、OpenAI、Anthropic、Google、Moonshot の 11 個の正確なモデル識別子にわたるゼロ可視バイト言語モデルの実行の成功に関する凍結されたクロスベンダー評価を報告します。 Void を、可視 U がまったくゼロで成功したプロバイダー応答を返すモデル実行として定義します。
[切り捨てられた]

記事本文:
クロスベンダーセマンティックボイドマトリックス |ゼノド
メインにスキップ
古いブラウザを使用しています。エクスペリエンスを向上させるためにブラウザをアップグレードしてください。
クロスベンダーセマンティックボイドマトリックス
このプレプリントは、OpenAI、Anthropic、Google、Moonshot の 11 個の正確なモデル識別子にわたるゼロ可視バイト言語モデルの実行の成功に関する凍結されたクロスベンダー評価を報告します。これは、Void を、可視 UTF-8 出力バイトがまったくゼロの成功したプロバイダー応答を返すモデル実行として定義します。プロバイダー終了メタデータによってそのサブタイプが決定されますが、明示的な拒否、安全ブロック、ツール仲介の実行、および操作の失敗は、引き続き個別の非 Void 結果として残ります。 31,430 回の完了したシングルターン トライアル全体で、マトリックスでは 11,658 個のボイド (37.09%) が観察されました。 4,290 個の厳密に一致したセマンティック ペアにおいて、ヌル アームは 2,505 個の Void を生成しましたが、一致した出力ライセンス付きコントロールは 0 個の Void を生成しました。 この研究では、4 つの同等ステータス サブタイプ (V0、V1、V2、および VU) を区別し、目に見える応答、Near-Void、拒否、安全ブロック、およびインフラストラクチャの結果を個別に保存します。 16,000トークンの上限では、313/500トライアルが無効のままで、すべてV0またはV2であり、認識された出力予算の終了が観察されたすべてのケースを説明するわけではないことを示しています。完全な生の試行、再試行、リクエスト構成、プロバイダーのメタデータ、および暗号化されたリンクされたイベント レコードは、公的に保存されます。検証では 62,968 個のイベント ハッシュが再生され、完全な 31,484 レコードの生の試行セットがチェックされました。この発見は、内部メカニズム、意図、共有アーキテクチャ、または因果関係の説明を主張することなく、ゼロ可視バイト実行の成功の再現可能で意味論的に構造化されたクラスを確立します。
Cross_Vendor_Semantic_Void_Matrix.pdf
統計の収集方法の詳細....
10.5281/ゼノド.21696066
マーク

ダウン
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21696066.svg)](https://doi.org/10.5281/zenodo.21696066)
再構造化されたテキスト
.. 画像:: https://zenodo.org/badge/DOI/10.5281/zenodo.21696066.svg
:target: https://doi.org/10.5281/zenodo.21696066
HTML
<a href="https://doi.org/10.5281/zenodo.21696066"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21696066.svg" alt="DOI"></a>
画像URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21696066.svg
ターゲット URL
https://doi.org/10.5281/zenodo.21696066
リソースの種類
プレプリント
出版社
ゼノド
権利
提供元
CERN データセンターと InvenioRDM
このサイトでは Cookie を使用しています。 Cookieの使用方法について詳しくはこちらをご覧ください

## Original Extract

This preprint reports a frozen cross-vendor evaluation of successful zero-visible-byte language-model executions across 11 exact model identifiers from OpenAI, Anthropic, Google, and Moonshot. It defines a Void as a model execution returning a successful provider response with exactly zero visible U
[truncated]

Cross-Vendor Semantic Void Matrix | Zenodo
Skip to main
You are using an outdated browser. Please upgrade your browser to improve your experience.
Cross-Vendor Semantic Void Matrix
This preprint reports a frozen cross-vendor evaluation of successful zero-visible-byte language-model executions across 11 exact model identifiers from OpenAI, Anthropic, Google, and Moonshot. It defines a Void as a model execution returning a successful provider response with exactly zero visible UTF-8 output bytes; provider termination metadata determines its subtype, while explicit refusals, safety blocks, tool-mediated executions, and operational failures remain distinct non-Void outcomes. Across 31,430 completed single-turn trials, the matrix observed 11,658 Voids (37.09%). In 4,290 strict matched semantic pairs, null arms produced 2,505 Voids while matched output-licensed controls produced 0. The study distinguishes four equal-status subtypes—V0, V1, V2, and VU—and separately preserves visible responses, Near-Voids, refusals, safety blocks, and infrastructure outcomes. At a 16,000-token ceiling, 313/500 trials remained Voids, all V0 or V2, showing that recognized output-budget termination does not explain all observed cases. Complete raw attempts, retries, request configurations, provider metadata, and cryptographically linked event records are publicly preserved. Verification replayed 62,968 event hashes and checked the complete 31,484-record raw-attempt set. The findings establish a reproducible, semantically structured class of successful zero-visible-byte executions without claiming an internal mechanism, intention, shared architecture, or causal explanation.
Cross_Vendor_Semantic_Void_Matrix.pdf
More info on how stats are collected....
10.5281/zenodo.21696066
Markdown
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21696066.svg)](https://doi.org/10.5281/zenodo.21696066)
reStructuredText
.. image:: https://zenodo.org/badge/DOI/10.5281/zenodo.21696066.svg
:target: https://doi.org/10.5281/zenodo.21696066
HTML
<a href="https://doi.org/10.5281/zenodo.21696066"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21696066.svg" alt="DOI"></a>
Image URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21696066.svg
Target URL
https://doi.org/10.5281/zenodo.21696066
Resource type
Preprint
Publisher
Zenodo
Rights
Powered by
CERN Data Centre & InvenioRDM
This site uses cookies. Find out more on how we use cookies
