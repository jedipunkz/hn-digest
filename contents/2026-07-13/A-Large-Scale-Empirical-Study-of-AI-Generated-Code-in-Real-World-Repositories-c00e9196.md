---
source: "https://arxiv.org/abs/2603.27130"
hn_url: "https://news.ycombinator.com/item?id=48899456"
title: "A Large-Scale Empirical Study of AI-Generated Code in Real-World Repositories"
article_title: "[2603.27130] A Large-Scale Comprehensive Measurement of AI-Generated Code in Real-World Repositories"
author: "softwaredoug"
captured_at: "2026-07-13T22:45:25Z"
capture_tool: "hn-digest"
hn_id: 48899456
score: 2
comments: 0
posted_at: "2026-07-13T21:56:58Z"
tags:
  - hacker-news
  - translated
---

# A Large-Scale Empirical Study of AI-Generated Code in Real-World Repositories

- HN: [48899456](https://news.ycombinator.com/item?id=48899456)
- Source: [arxiv.org](https://arxiv.org/abs/2603.27130)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T21:56:58Z

## Translation

タイトル: 現実世界のリポジトリにおける AI 生成コードの大規模実証研究
記事のタイトル: [2603.27130] 現実世界のリポジトリにおける AI 生成コードの大規模な包括的な測定
説明: arXiv 論文 2603.27130: 現実世界のリポジトリにおける AI 生成コードの大規模な包括的測定の要約ページ

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
コンピュータサイエンス > ソフトウェアエンジニアリング
[2026 年 3 月 28 日に提出 ( v1 )、最終改訂日 2026 年 7 月 1 日 (このバージョン、v3)]
タイトル: 現実世界のリポジトリにおける AI 生成コードの大規模な包括的な測定
要約: 大規模言語モデル (LLM) は、開発者が小さなスニペットからプロジェクト全体に至るまでのコードを生成できるようにすることで、ソフトウェア エンジニアリングを急速に変革しています。 AI 支援コードが現実世界のシステムにますます統合されるようになるにつれ、その特性と影響を理解することが重要になります。 AI 生成コードに関する既存の研究は、通常、合成ベンチマークと小規模コーディング タスクを使用したラボ設定に限定されており、限られた指標をカバーしています。 AI 支援コードが現実世界のコードベースに現れること、および人間が作成したコードとの違いは依然として不明瞭です。このギャップを埋めるために、私たちは、現実世界のリポジトリで、人間が書いたコードと比較する、AI 支援コードの最初の大規模な測定研究を実行します。私たちは、コードレベルの側面（構造レベルおよびグラフレベルの複雑さ、コーディングスタイル、セキュリティ品質など）とコミットレベルの特性（コミットサイズ、頻度、コミット後の安定性など）の両方を含む包括的な一連の指標を研究します。私たちの結果は新たな発見と洞察を提供します。実験室環境での以前の観察を対照するものもあれば（例、コードレベルの指標における現実世界の AI と人間の違いは顕著ではなくかなり小さいと結論付けています）、以前の結果をより詳細な観察で拡張するものもあります（例：異なるプログラミング言語間でのセキュリティ品質のばらつき）、さらに多くの結果が、そうでない側面について初めて提示されます。

これについては以前に説明しました (コードの重複率、コミット サイズ、安定性など)。これらの包括的な現実世界の結果に基づいて、AI 支援プログラミングの実際的な意味についても説明します。
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

Abstract page for arXiv paper 2603.27130: A Large-Scale Comprehensive Measurement of AI-Generated Code in Real-World Repositories

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
Computer Science > Software Engineering
[Submitted on 28 Mar 2026 ( v1 ), last revised 1 Jul 2026 (this version, v3)]
Title: A Large-Scale Comprehensive Measurement of AI-Generated Code in Real-World Repositories
Abstract: Large language models (LLMs) are rapidly transforming software engineering by enabling developers to generate code ranging from small snippets to entire projects. As AI-assisted code becomes increasingly integrated into real-world systems, understanding its characteristics and impact is critical. Existing study on AI-generated code is usually limited in the lab setting with synthetic benchmarks and small-scale coding tasks and covers limited metrics. AI-assisted code's manifestation in real-world codebases and its differences between human-written one remain unclear. To close this gap, we perform a first large-scale measurement study of AI-assisted code, in comparison with the human-written, in real-world repositories. We study a comprehensive set of metrics including both code-level aspects (e.g., structural and graph-level complexity, coding style, security quality, etc.) and commit-level characteristics (e.g., commit size, frequency, post-commit stability, etc.). Our results provide new findings and insights: some contrast previous observations in the lab setting (e.g., we conclude that real-world AI-Human differences on code-level metrics are rather small instead of more pronounced), some extend prior results with finer-grained observations (e.g., the variance of security quality across different programming languages), yet more are presented for the first time on aspects not covered before (e.g., code duplication rate, commit size and stability, etc.). Based on these comprehensive real-world results, we also discuss the practical implications of AI-assisted programming.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
