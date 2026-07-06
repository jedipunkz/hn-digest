---
source: "https://arxiv.org/abs/2402.14873"
hn_url: "https://news.ycombinator.com/item?id=48800043"
title: "Technical Report on the Pangram AI-Generated Text Classifier (2024)"
article_title: "[2402.14873] Technical Report on the Pangram AI-Generated Text Classifier"
author: "imustachyou"
captured_at: "2026-07-06T02:47:57Z"
capture_tool: "hn-digest"
hn_id: 48800043
score: 1
comments: 0
posted_at: "2026-07-06T02:27:50Z"
tags:
  - hacker-news
  - translated
---

# Technical Report on the Pangram AI-Generated Text Classifier (2024)

- HN: [48800043](https://news.ycombinator.com/item?id=48800043)
- Source: [arxiv.org](https://arxiv.org/abs/2402.14873)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T02:27:50Z

## Translation

タイトル: Pangram AI 生成テキスト分類器に関する技術レポート (2024)
記事タイトル: [2402.14873] パングラム AI 生成テキスト分類器に関する技術レポート
説明: arXiv 論文 2402.14873 の要約ページ: パングラム AI 生成テキスト分類器に関する技術レポート

記事本文:
メインコンテンツにスキップ
arXiv は独立した非営利団体になりました。
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 計算と言語
[2024 年 2 月 21 日に提出 ( v1 )、最終改訂日 2024 年 7 月 29 日 (このバージョン、v3)]
タイトル: Pangram AI 生成テキスト分類器に関する技術レポート
要約: 我々は、大規模な言語モデルによって書かれたテキストと人間によって書かれたテキストを区別するように訓練されたトランスフォーマーベースのニューラルネットワークであるパングラムテキストを紹介します。 Pangram Text は、10 のテキスト領域 (学生の文章、創造的な文章、科学的文章、書籍、百科事典、ニュース、電子メール、科学論文、短形式の Q&A) と 8 つのオープンソースおよびクローズドソースの大規模言語モデルで構成される包括的なベンチマークで、DetectGPT などのゼロショット手法や主要な商用 AI 検出ツールよりも優れたパフォーマンスを示し、エラー率が 38 倍以上低くなります。私たちは、合成ミラーを使用したハード ネガティブ マイニングというトレーニング アルゴリズムを提案します。これにより、分類器はレビューなどの高データ ドメインで桁違いに低い誤検知率を達成できるようになります。最後に、パングラム テキストが英語を母国語としない話者に対して偏見を持たず、トレーニング中には見ら​​れなかった領域やモデルに一般化されることを示します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの値にコミットされており、p でのみ動作します。

それらにこだわるアーティストたち。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2402.14873: Technical Report on the Pangram AI-Generated Text Classifier

Skip to main content
arXiv is now an independent nonprofit!
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 21 Feb 2024 ( v1 ), last revised 29 Jul 2024 (this version, v3)]
Title: Technical Report on the Pangram AI-Generated Text Classifier
Abstract: We present Pangram Text, a transformer-based neural network trained to distinguish text written by large language models from text written by humans. Pangram Text outperforms zero-shot methods such as DetectGPT as well as leading commercial AI detection tools with over 38 times lower error rates on a comprehensive benchmark comprised of 10 text domains (student writing, creative writing, scientific writing, books, encyclopedias, news, email, scientific papers, short-form Q&A) and 8 open- and closed-source large language models. We propose a training algorithm, hard negative mining with synthetic mirrors, that enables our classifier to achieve orders of magnitude lower false positive rates on high-data domains such as reviews. Finally, we show that Pangram Text is not biased against nonnative English speakers and generalizes to domains and models unseen during training.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
