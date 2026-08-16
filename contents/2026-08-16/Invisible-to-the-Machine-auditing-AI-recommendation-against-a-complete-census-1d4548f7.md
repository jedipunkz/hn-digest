---
source: "https://arxiv.org/abs/2608.07069"
hn_url: "https://news.ycombinator.com/item?id=49317814"
title: "Invisible to the Machine: auditing AI recommendation against a complete census"
article_title: "[2608.07069] Invisible to the Machine: Auditing AI Restaurant, Cafe, and Bar Recommendation Against a Complete Market Census"
author: "fspunch"
captured_at: "2026-08-16T08:17:33Z"
capture_tool: "hn-digest"
hn_id: 49317814
score: 1
comments: 0
posted_at: "2026-08-16T07:49:37Z"
tags:
  - hacker-news
  - translated
---

# Invisible to the Machine: auditing AI recommendation against a complete census

- HN: [49317814](https://news.ycombinator.com/item?id=49317814)
- Source: [arxiv.org](https://arxiv.org/abs/2608.07069)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T07:49:37Z

## Translation

タイトル: マシンには見えない: 完全な国勢調査に対する AI 推奨事項の監査
記事のタイトル: [2608.07069] マシンには見えない: 完全な市場調査に対する AI レストラン、カフェ、バーの推奨事項の監査
説明: arXiv 論文 2608.07069 の要約ページ: Invisible to the Machine: Auditing AI Restaurant, Cafe, and Bar Recommendation Against a Complete Market Census

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 情報検索
[2026 年 8 月 7 日に提出]
タイトル: 機械には見えない: 完全な市場調査に対する AI レストラン、カフェ、バーの推奨事項の監査
要約: AI アシスタントは、地域の発見のための主要なインターフェースになりつつありますが、AI アシスタントがどのような場所で表面化するかについてはほとんど何も知られていません。特に、推奨事項が収益に直接影響する飲食業界ではそうです。我々は、AI 開催場所の推奨に関する初めての国勢調査に基づく監査を提示します。これは、2 つの限定された市場 (バリ島のチャングーとウブド) にわたる 4,776 軒のカフェ、レストラン、バーの完全な列挙です。これに対して、事前に登録されたプロトコルに基づいて 7 日間にわたって収集された、96 のペルソナ条件付きクエリに対する 4 つの本番 AI システム (ChatGPT、Claude、Gemini、Perplexity) からの 2,208 件の検索に基づいた応答を評価します。私たちは市場全体を観察しているため、サンプリングされた監査では測定できないものを測定できます。会場の 85.6% はいかなるシステムからも推奨されておらず、50 以上の評価を持つ確立された会場でも 72.6% でした。可視性は 2 つのマージン構造に従います。回答への入力はドキュメントと関連付けられています: レビューの量 (OR 1.64)、独自の Web サイト (OR 1.92)、定価情報 (OR 1.54)、およびサードパーティの Web での言及 (OR 1.44) -- このマージン (OR 0.89) では星評価はゼロです。回答内のランクはパターンを逆転させます。推奨会場の中で、評価は 1 位を大幅に予測します (OR 1.17)。オープン POI データセット (Foursquare) に存在することは、民間理論による可視性要因ですが、どちらのマージンでもプラスの効果は示されていません。完全な捏造はまれですが（言及の 0.08%）、システムは会場の永久閉鎖を 93 回推奨しました - ホールではなく陳腐化

ユーシネーションは実際の故障モードです。システム間の一致度は低いです (上位 20 Jaccard 0.33 ～ 0.54)。 2 週間のテストと再テストでは、期間をまたいだ回答の類似性が同日の再実行の類似性に匹敵することがわかります。チャーンはサンプリングの確率論であり、時間的なドリフトではありません。プロトコル、レジストリ構築方法、派生データを公開します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.07069: Invisible to the Machine: Auditing AI Restaurant, Cafe, and Bar Recommendation Against a Complete Market Census

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Information Retrieval
[Submitted on 7 Aug 2026]
Title: Invisible to the Machine: Auditing AI Restaurant, Cafe, and Bar Recommendation Against a Complete Market Census
Abstract: AI assistants are becoming a primary interface for local discovery, yet almost nothing is known about which venues they surface -- especially in food and drink, where recommendations carry direct revenue consequences. We present the first census-denominated audit of AI venue recommendation: a complete enumeration of 4,776 cafes, restaurants, and bars across two bounded markets (Canggu and Ubud, Bali), against which we evaluate 2,208 search-grounded responses from four production AI systems (ChatGPT, Claude, Gemini, Perplexity) to 96 persona-conditioned queries, collected over seven days under a pre-registered protocol. Because we observe the full market, we can measure what sampled audits cannot: 85.6% of venues were never recommended by any system -- 72.6% even among established venues with fifty or more ratings. Visibility follows a two-margin structure. Entry into answers is associated with documentation: review volume (OR 1.64), an own website (OR 1.92), listed price information (OR 1.54), and third-party web mentions (OR 1.44) -- while star rating is null at this margin (OR 0.89). Rank within answers reverses the pattern: among recommended venues, rating significantly predicts first position (OR 1.17). Presence in an open POI dataset (Foursquare), a folk-theorized visibility factor, shows no positive effect at either margin. Outright fabrication is rare (0.08% of mentions), but systems recommended permanently closed venues 93 times -- staleness, not hallucination, is the practical failure mode. Cross-system agreement is low (top-20 Jaccard 0.33-0.54). A two-week test-retest shows cross-period answer similarity comparable to same-day rerun similarity: the churn is sampling stochasticity, not temporal drift. We release our protocol, registry construction method, and derived data.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
