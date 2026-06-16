---
source: "https://arxiv.org/abs/2606.15762"
hn_url: "https://news.ycombinator.com/item?id=48557881"
title: "Snyk VulnBench JavaScript 1.0: Can LLMs Find the Same Bugs Twice?"
article_title: "[2606.15762] Snyk VulnBench JS 1.0: Can LLMs Find the Same Bugs Twice?"
author: "ilreb"
captured_at: "2026-06-16T16:35:54Z"
capture_tool: "hn-digest"
hn_id: 48557881
score: 1
comments: 0
posted_at: "2026-06-16T16:32:45Z"
tags:
  - hacker-news
  - translated
---

# Snyk VulnBench JavaScript 1.0: Can LLMs Find the Same Bugs Twice?

- HN: [48557881](https://news.ycombinator.com/item?id=48557881)
- Source: [arxiv.org](https://arxiv.org/abs/2606.15762)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T16:32:45Z

## Translation

タイトル: Snyk VulnBench JavaScript 1.0: LLM は同じバグを 2 回見つけることができますか?
記事のタイトル: [2606.15762] Snyk VulnBench JS 1.0: LLM は同じバグを 2 回見つけることができますか?
説明: arXiv 論文 2606.15762 の要約ページ: Snyk VulnBench JS 1.0: LLM は同じバグを 2 回見つけることができますか?

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.15762
ヘルプ |高度な検索
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 6 月 14 日に提出]
タイトル: Snyk VulnBench JS 1.0: LLM は同じバグを 2 回見つけることができますか?
要約: 同じ JavaScript コード、プロンプト、ベンチマーク ハーネス上でエージェント的大規模言語モデル (LLM) セキュリティ レビューがどの程度反復可能であるかを測定するために、脆弱性発見スキャンを 300 回繰り返し実行しました。見出しの結果は、LLM セキュリティの結果の再現性にばらつきがあるということです。参照一致の結果は安定していましたが、追加のモデル レポートは実行ごとに大きく異なりました。 250 回のモデル実行を通じて、161 件の一致しない固有の結果のうち 80 件が 5 回の同一反復のうち 1 回のみに出現し、5 回すべてで出現したのは 22 件のみでした。対照的に、Claude が Snyk コードの参照結果と一致した場合、動作ははるかに安定しました。158 の固有の参照一致結果のうち 134 が、5 回の繰り返しすべてで出現しました。ベンチマークは相補性も示しています。モデルには、よく知られた高信号エクスプロイトの形状が一貫して見つかり、あるケースでは、Snyk Code 製品のギャップと思われる箇所が表面化しました。 Snyk Code の静的アプリケーション セキュリティ テスト (SAST) は決定論的であり、繰り返されるデータ フロー シンクを体系的に列挙することに優れていました。この結果は、どちらかの手法を他方の手法の代替として扱うのではなく、エージェント的な LLM レビューと決定論的な SAST を組み合わせることを裏付けています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業を可能にするフレームワークです。

研究者は、新しい arXiv 機能を開発し、Web サイト上で直接共有できます。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.15762: Snyk VulnBench JS 1.0: Can LLMs Find the Same Bugs Twice?

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.15762
Help | Advanced Search
Computer Science > Cryptography and Security
[Submitted on 14 Jun 2026]
Title: Snyk VulnBench JS 1.0: Can LLMs Find the Same Bugs Twice?
Abstract: We ran 300 repeated vulnerability-finding scans to measure how repeatable agentic large language model (LLM) security review is on the same JavaScript code, prompt, and benchmark harness. The headline result is that LLM security findings were unevenly repeatable: reference-matched findings were stable, but extra model reports varied heavily from run to run. Across 250 model runs, 80 of 161 unique unmatched findings appeared in only one of five identical repetitions, while only 22 appeared in all five. By contrast, when Claude matched a Snyk Code reference finding, the behavior was much more stable: 134 of 158 unique reference-matched findings appeared in all five repetitions. The benchmark also shows complementarity. Models consistently found familiar, high-signal exploit shapes, and in one case surfaced a likely Snyk Code product gap. Snyk Code static application security testing (SAST) was deterministic and better at systematically enumerating repeated data-flow sinks. The results support combining agentic LLM review with deterministic SAST rather than treating either technique as a replacement for the other.
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
