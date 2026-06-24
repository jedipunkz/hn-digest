---
source: "https://arxiv.org/abs/2606.20047"
hn_url: "https://news.ycombinator.com/item?id=48665300"
title: "Submodular Context Selection as a Pluggable Engine for LLM Agents"
article_title: "[2606.20047] PACMS: Submodular Context Selection as a Pluggable Engine for LLM Agents"
author: "Elof"
captured_at: "2026-06-24T21:30:13Z"
capture_tool: "hn-digest"
hn_id: 48665300
score: 1
comments: 0
posted_at: "2026-06-24T20:36:10Z"
tags:
  - hacker-news
  - translated
---

# Submodular Context Selection as a Pluggable Engine for LLM Agents

- HN: [48665300](https://news.ycombinator.com/item?id=48665300)
- Source: [arxiv.org](https://arxiv.org/abs/2606.20047)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T20:36:10Z

## Translation

タイトル: LLM エージェントのプラグ可能エンジンとしてのサブモジュール式コンテキスト選択
記事のタイトル: [2606.20047] PACMS: LLM エージェントのプラグイン可能なエンジンとしてのサブモジュラー コンテキスト選択
説明: arXiv 論文 2606.20047 の要約ページ: PACMS: LLM エージェントのプラグ可能エンジンとしてのサブモジュール コンテキスト選択

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.20047
ヘルプ |高度な検索
コンピュータサイエンス > 情報検索
[2026 年 6 月 18 日に提出]
タイトル: PACMS: LLM エージェントのプラグ可能エンジンとしてのサブモジュール式コンテキスト選択
要約: 会話およびツールを使用する LLM エージェントは、複数の方向から同時に満たされるコンテキスト ウィンドウ上で動作します。セッションが進行するにつれて、エージェントはユーザーとアシスタントのターン、永続メモリ ストアから取得されたエントリ、そして多くの場合、ファイル読み取り、検索結果、API 応答などのツール呼び出しの逐語的出力を蓄積します。累積的なコンテキストがモデルのトークン バジェットを超えると、フレームワークは何を保持するかを決定する必要があります。
一般的なメカニズムは最新性の切り捨てであり、場合によっては定期的な要約と組み合わせられます。これはトピックブラインドです。セッションの初期に確立された事実は、現在のユーザーのクエリがまさにその事実に関するものであっても、それが古いという理由だけで破棄されます。逆に、冗長だが無関係な最近の内容は保持されます。エージェントが何度も情報を思い出す必要がある場合、これは記憶の決定的なケースであり、まさに最新性の切り捨てが失敗する場所です。
既存の代替手段は、エージェントの組み立てステップの外にあります。検索拡張生成は外部ドキュメントをプロンプトにフェッチしますが、エージェントの \emph{already-present} プールされたコンテキストを調停しません。コンテキスト圧縮方法は、テキストの書き換えや削除によってトークン数を減らしますが、クエリブラインドで損失が大きく動作します。どちらも、メモリ エントリ、会話ターン、およびツール出力を、プロンプトが組み立てられた瞬間の関連性によって選択される単一の候補プールとして扱いません。

もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
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

Abstract page for arXiv paper 2606.20047: PACMS: Submodular Context Selection as a Pluggable Engine for LLM Agents

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.20047
Help | Advanced Search
Computer Science > Information Retrieval
[Submitted on 18 Jun 2026]
Title: PACMS: Submodular Context Selection as a Pluggable Engine for LLM Agents
Abstract: Conversational and tool-using LLM agents operate over a context window that fills from several directions simultaneously. As a session proceeds, the agent accumulates user and assistant turns, entries drawn from a persistent memory store, and often largest of all, the verbatim outputs of tool calls such as file reads, search results, and API responses. Once the cumulative context exceeds the model's token budget, the framework must decide what to keep.
The prevailing mechanism is recency truncation, sometimes paired with periodic summarization. This is topic-blind: a fact established early in a session is discarded simply because it is old, even when the current user query is about exactly that fact; conversely, verbose but irrelevant recent material is retained. Agents that must recall information across many turns, the defining case for memory, are precisely where recency truncation fails.
Existing alternatives sit outside the agent's assembly step. Retrieval augmented generation fetches external documents into the prompt but does not arbitrate the agent's \emph{already-present} pooled context. Context-compression methods reduce token count by rewriting or pruning text, but operate query-blind and lossily. Neither treats memory entries, conversation turns, and tool outputs as a single candidate pool to be selected from by relevance at the moment the prompt is assembled.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
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
