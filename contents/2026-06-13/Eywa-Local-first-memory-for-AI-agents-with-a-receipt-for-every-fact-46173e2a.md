---
source: "https://arxiv.org/abs/2605.30771"
hn_url: "https://news.ycombinator.com/item?id=48519731"
title: "Eywa: Local-first memory for AI agents, with a receipt for every fact"
article_title: "[2605.30771] Eywa: Provenance-Grounded Long-Term Memory for AI Agents"
author: "agentseal"
captured_at: "2026-06-13T18:34:46Z"
capture_tool: "hn-digest"
hn_id: 48519731
score: 1
comments: 0
posted_at: "2026-06-13T17:57:21Z"
tags:
  - hacker-news
  - translated
---

# Eywa: Local-first memory for AI agents, with a receipt for every fact

- HN: [48519731](https://news.ycombinator.com/item?id=48519731)
- Source: [arxiv.org](https://arxiv.org/abs/2605.30771)
- Score: 1
- Comments: 0
- Posted: 2026-06-13T17:57:21Z

## Translation

タイトル: Eywa: AI エージェントのためのローカルファーストの記憶、すべての事実の領収書付き
記事タイトル: [2605.30771] Eywa: AI エージェントのための出所に基づいた長期記憶
説明: arXiv 論文 2605.30771 の要約ページ: Eywa: AI エージェントのための来歴に基づいた長期記憶

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2605.30771
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 5 月 29 日提出]
タイトル: Eywa: AI エージェントのための出所に基づいた長期記憶
要約: セッション間で持続する AI エージェントには、取得、監査、更新、消去できるメモリが必要です。既存のメモリ システムでは、多くの場合、ソース証拠、抽出された事実、取得されたコンテキスト、および回答ポリシーが 1 つの不透明なプロンプト パスにまとめられているため、障害の診断が困難になります。間違った回答は、証拠の欠落、サポートされていない抽出、古い状態、取得の損失、または回答モデルの動作に起因する可能性があります。私たちは、信じる前に証拠を中心に構築された来歴に基づいた記憶アーキテクチャである Eywa を紹介します。 Eywa は、標準的な事実を導出する前に不変のソース証拠を保存し、抽出されたメモリを型指定された信号およびソース サポートに対して検証し、取得内での LLM 呼び出しがゼロの決定論的なマルチルート読み取りパスを通じて境界付きメモリ コンテキストを取得します。取得されたコンテキストは応答命令とは別に返されるため、フロンティア、予算、およびローカル応答モデル全体で同じメモリ基板を評価できます。凍結されたアーティファクトが記録された取得構成の下で、Eywa は、Claude Sonnet 4.6 書き込みおよび QA ロールを使用した LoCoMo C1-C4 スプリットで 90.19% の判定精度に達しました。 LongMemEval-S では、検索十分性の精度は 88.2% に達します。 700 問のテクニカル メモリ ストレス ベンチマークである BEAM では、平均ナゲット スコアが 81.45%、pass@score >= 0.5 が 85.29% に達しています。質問、ゴールドアンサー、模範解答、取得されたコンテキスト、ラベルなど、質問ごとの完全なアーティファクトは、

この https URL 。
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

Abstract page for arXiv paper 2605.30771: Eywa: Provenance-Grounded Long-Term Memory for AI Agents

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2605.30771
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 29 May 2026]
Title: Eywa: Provenance-Grounded Long-Term Memory for AI Agents
Abstract: AI agents that persist across sessions need memory they can retrieve, audit, update, and erase. Existing memory systems often collapse source evidence, extracted facts, retrieved context, and answer policy into one opaque prompt path, making failures difficult to diagnose: a wrong answer may come from missing evidence, unsupported extraction, stale state, retrieval loss, or answer-model behavior. We present Eywa, a provenance-grounded memory architecture built around evidence before belief. Eywa stores immutable source evidence before deriving canonical facts, validates extracted memories against typed signals and source support, and retrieves bounded memory context through a deterministic multi-route read path with zero LLM calls inside retrieval. Retrieved context is returned separately from answer instructions, allowing the same memory substrate to be evaluated across frontier, budget, and local answer models. Under a frozen, artifact-recorded retrieval configuration, Eywa reaches 90.19% judge accuracy on the LoCoMo C1-C4 split with Claude Sonnet 4.6 write and QA roles. On LongMemEval-S, it reaches 88.2% retrieval-sufficiency accuracy. On BEAM, a 700-question technical-memory stress benchmark, it reaches 81.45% mean nugget score and 85.29% pass@score >= 0.5. Full per-question artifacts, including questions, gold answers, model answers, retrieved context, and labels, are published at this https URL .
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
