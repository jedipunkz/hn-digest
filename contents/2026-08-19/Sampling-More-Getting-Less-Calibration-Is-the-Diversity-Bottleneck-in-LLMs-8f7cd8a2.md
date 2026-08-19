---
source: "https://arxiv.org/abs/2605.11128"
hn_url: "https://news.ycombinator.com/item?id=49361621"
title: "Sampling More, Getting Less: Calibration Is the Diversity Bottleneck in LLMs"
article_title: "[2605.11128] Sampling More, Getting Less: Calibration is the Diversity Bottleneck in LLMs"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "clukic"
captured_at: "2026-08-19T14:24:05Z"
capture_tool: "hn-digest"
hn_id: 49361621
score: 1
comments: 0
posted_at: "2026-08-19T13:51:35Z"
tags:
  - hacker-news
  - translated
---

# Sampling More, Getting Less: Calibration Is the Diversity Bottleneck in LLMs

- HN: [49361621](https://news.ycombinator.com/item?id=49361621)
- Source: [arxiv.org](https://arxiv.org/abs/2605.11128)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T13:51:35Z

## Translation

タイトル: サンプリングを増やし、サンプリングを減らす: キャリブレーションは LLM における多様性のボトルネックです
記事のタイトル: [2605.11128] サンプリングを増やし、サンプリングを減らす: キャリブレーションは LLM における多様性のボトルネックです
説明: arXiv 論文 2605.11128 の要約ページ: Sampling More, Getting Less: Calibration is the Diversity Bottleneck in LLMs

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
[2026 年 5 月 11 日に提出]
タイトル: サンプリングを増やし、サンプリングを減らす: キャリブレーションは LLM における多様性のボトルネックです
要約: 多様性は、創造的な生成から科学的発見に至るまでの言語モデルのアプリケーションにとって不可欠ですが、最新の LLM は、妥当な出力の狭いサブセットに崩壊することがよくあります。これまでの研究では、この多様性の欠如を測定するためのベンチマークが開発されてきましたが、推論時の段階的な確率分布がどのようにして問題を引き起こすかについてはあまり知られていません。我々は、多様性の崩壊の原因を、デコード中に LLM が有効な継続と無効な継続にまたがって確率質量を割り当てる方法に起因する、妥当性 - 多様性フレームワークを導入します。このフレームワークは、ボトルネックを 2 つの相補的な形式のミスキャリブレーションに分解します。まず、順序の調整です。有効なトークンが無効なトークンより上位に確実にランク付けされるわけではないため、ランクベースのカットオフ ルールは、有効な継続の回復と無効な継続の許可の間でトレードオフを行う必要があります。第二に、形状の調整: 確率質量は、有効なトークンと無効なトークンが混在するヘビーテールを持ちながら、少数の有効な継続のみに過度に集中するため、高い妥当性を維持すると多様性が制限されます。我々は両方のメカニズムを形式化し、局所的な失敗が復号化ステップ全体で複合し、多様性に強い配列レベルの損失を引き起こすことを示します。私たちは経験的に、これらのボトルネックを調査するための管理された診断を開発します。これには、正確に既知の有効なセットとオラクル カットオフ ベースラインを使用したタスクが含まれます。複数の族やスケールにまたがる 14 の言語モデルにわたって、多様性の崩壊は単に特定のサンプリング ヒューリスティックの制限ではなく、次のような結果であることがわかりました。

LLM 分布における順序と形状の誤調整。
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

Abstract page for arXiv paper 2605.11128: Sampling More, Getting Less: Calibration is the Diversity Bottleneck in LLMs

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 11 May 2026]
Title: Sampling More, Getting Less: Calibration is the Diversity Bottleneck in LLMs
Abstract: Diversity is essential for language-model applications ranging from creative generation to scientific discovery, yet modern LLMs often collapse into a narrow subset of plausible outputs. While prior work has developed benchmarks for measuring this lack of diversity, less is known about how the step-by-step probability distributions at inference time cause the problem. We introduce a validity--diversity framework that attributes diversity collapse to how an LLM allocates probability mass across valid and invalid continuations during decoding. This framework decomposes the bottleneck into two complementary forms of miscalibration. First, order calibration: valid tokens are not reliably ranked above invalid tokens, so rank-based cutoff rules must trade off between recovering valid continuations and admitting invalid ones. Second, shape calibration: probability mass is overly concentrated only on few valid continuations while having a heavy-tail of mixed valid and invalid tokens, so maintaining high validity limits diversity. We formalize both mechanisms and show that local failures compound across decoding steps, producing strong sequence-level losses in diversity. Empirically, we develop controlled diagnostics for probing these bottlenecks, including tasks with exactly known valid sets and oracle cutoff baselines. Across 14 language models spanning multiple families and scales, we find that diversity collapse is not merely a limitation of particular sampling heuristics, but a consequence of order and shape miscalibration in the LLM distribution.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
