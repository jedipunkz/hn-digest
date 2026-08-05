---
source: "https://arxiv.org/abs/2607.29377"
hn_url: "https://news.ycombinator.com/item?id=49178608"
title: "Zero-Mem: Zero-Token Memory Operations for LLM Agents"
article_title: "[2607.29377] Zero-Mem: Zero-Token Memory Operations for LLM Agents"
author: "theanonymousone"
captured_at: "2026-08-05T04:55:50Z"
capture_tool: "hn-digest"
hn_id: 49178608
score: 1
comments: 0
posted_at: "2026-08-05T04:36:44Z"
tags:
  - hacker-news
  - translated
---

# Zero-Mem: Zero-Token Memory Operations for LLM Agents

- HN: [49178608](https://news.ycombinator.com/item?id=49178608)
- Source: [arxiv.org](https://arxiv.org/abs/2607.29377)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T04:36:44Z

## Translation

タイトル: Zero-Mem: LLM エージェントのためのゼロトークンメモリ操作
記事のタイトル: [2607.29377] Zero-Mem: LLM エージェントのゼロトークン メモリ操作
説明: arXiv 論文 2607.29377 の要約ページ: Zero-Mem: LLM エージェントのゼロトークンメモリ操作

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 計算と言語
[2026 年 7 月 31 日に提出]
タイトル: Zero-Mem: LLM エージェントのためのゼロトークンメモリ操作
要約: LLM エージェントは、長い対話にわたって一貫して動作するためにメモリを必要としますが、多くのシステムはそのメモリを操作するために追加の LLM 呼び出しを使用します。中間レコードの生成とその取得の仲介では、定期的なトークンと時間のコストが追加されますが、詳細が省略されたり統合されたりすると、元の証拠が不明瞭になる可能性があります。構造化メモリアクセスにはそもそも生成が必要なのかどうかを尋ねます。 Zero-Mem は \emph{ゼロトークン メモリ操作} を導入します。最終的な質問応答以外のステップでは、LLM を呼び出したり、LLM 入力トークンまたは出力トークンを消費したりすることはありません。エンコーダの計算は個別に考慮されます。 Zero-Mem は、元の対話トレースを記録ソースとして保存します。 2 つの相補的な方法でトレースを整理します。エンティティ - コンテキスト グラフはインタラクション全体の接続を公開しますが、時間階層は会話の局所性とセッション状態を保持します。クエリごとに、Zero-Mem は 2 つのビューを重み付けし、両方から取得し、その構造に従ってサポート関係や周囲のコンテキストを復元します。決定論的キャリブレーションでは、まず矛盾する証拠を破棄し、その後、読み取ったトレースに基づく読者の回答を維持します。最終 QA リーダーのみが LLM を呼び出します。 Zero-Mem は、ロングメモリおよびロングコンテキストの質問応答ベンチマーク全体で、メモリ操作による LLM 呼び出しと LLM トークンの消費を排除しながら、競争力のあるパフォーマンスを実現します。同じ最終 QA リーダーとコンテキスト バジェットを使用すると、最速の比較ベースラインと比較して、メモリ操作時間のコストが 57.6% 削減されます。アブレーションは 2 つの要因の寄与をサポートします。

ビューとそのクエリ依存の調整。全体として、結果は、構造化されたエージェントの記憶が過去の中間表現を生成する必要がないことを示しています。ピアレビューの後、コードと実装の詳細は \textcolor{blue}{ this https URL } で入手できるようになります。
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

Abstract page for arXiv paper 2607.29377: Zero-Mem: Zero-Token Memory Operations for LLM Agents

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 31 Jul 2026]
Title: Zero-Mem: Zero-Token Memory Operations for LLM Agents
Abstract: LLM agents need memory to act consistently over long interactions, yet many systems use additional LLM calls to operate that memory. Generating intermediate records and mediating their retrieval adds recurring token and time costs, while omitted or merged details can obscure the original evidence. We ask whether structured memory access requires generation at all. Zero-Mem introduces \emph{zero-token memory operations}: no step outside final question answering invokes an LLM or consumes LLM input or output tokens; encoder computation is accounted for separately. Zero-Mem preserves original interaction traces as its source of record. It organizes the traces in two complementary ways. An entity--context graph exposes connections across interactions, while a temporal hierarchy preserves conversational locality and session state. For each query, Zero-Mem weighs the two views, retrieves from both, and follows their structure to recover supporting relations or surrounding context. Deterministic calibration first discards conflicting evidence and then keeps the reader's answer grounded in the retrieved traces. Only the final-QA reader invokes an LLM. Across long-memory and long-context question-answering benchmarks, Zero-Mem achieves competitive performance while eliminating LLM calls and LLM-token consumption from memory operations. With the same final-QA reader and context budget, it reduces memory-operation time cost by 57.6\% relative to the fastest compared baseline. Ablations support the contribution of the two views and their query-dependent coordination. Overall, the results show that structured agent memory need not generate an intermediate representation of the past. After peer review, the code and implementation details will be available at \textcolor{blue}{ this https URL }.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
