---
source: "https://arxiv.org/abs/2608.00799"
hn_url: "https://news.ycombinator.com/item?id=49208672"
title: "AI Meet CAD: Beat Opus/Mythos 5, GPT 5.6 Sol on BenchCAD"
article_title: "[2608.00799] CADENA: Stepwise CAD Reverse Engineering"
author: "SUPERustam"
captured_at: "2026-08-07T11:38:44Z"
capture_tool: "hn-digest"
hn_id: 49208672
score: 1
comments: 1
posted_at: "2026-08-07T11:17:20Z"
tags:
  - hacker-news
  - translated
---

# AI Meet CAD: Beat Opus/Mythos 5, GPT 5.6 Sol on BenchCAD

- HN: [49208672](https://news.ycombinator.com/item?id=49208672)
- Source: [arxiv.org](https://arxiv.org/abs/2608.00799)
- Score: 1
- Comments: 1
- Posted: 2026-08-07T11:17:20Z

## Translation

タイトル: AI と CAD: Beat Opus/Mythos 5、GPT 5.6 Sol on BenchCAD
記事のタイトル: [2608.00799] CADENA: ステップワイズ CAD リバース エンジニアリング
説明: arXiv 論文 2608.00799 の要約ページ: CADENA: ステップワイズ CAD リバース エンジニアリング

記事本文:
メインコンテンツにスキップ
システムメンテナンス 8月4日・5日
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューター サイエンス > コンピューター ビジョンとパターン認識
[2026 年 8 月 1 日に提出]
タイトル: CADENA: ステップワイズ CAD リバース エンジニアリング
要約: コンピュータ支援設計 (CAD) は現代のエンジニアリングを支えていますが、既存の形状を編集可能なモデルに変換するには依然として専門家の多大な努力が必要です。ほとんどの AI システムは、CAD プログラム全体を単一パスで出力し、中間ジオメトリを検査することはありません。対照的に、人間のエンジニアは部品の機能ごとに構築し、各操作の後にモデリングが必要なものを確認します。 CADENA (スペイン語で「鎖」) を紹介します。このモデルは、パラメトリック CAD プログラムとして 3D メッシュを再構築し、一連の操作を一度に 1 つずつ拡大し、ステップごとにターゲットと現在予測されているジオメトリを比較します。また、機械部品のリバース エンジニアリング手法を評価するためのベンチマークの不足にも対処し、機械部品のカテゴリ全体のパフォーマンスを測定するベンチマークである CADENA-Bench を導入します。 CADENA は、CADENA-Bench、DeepCAD、Fusion 360、および MCB データセットで以前の方法よりも優れたパフォーマンスを発揮します。コードはこの https URL で、モデルの重みはこの https URL で、CADENA-Bench はこの https URL で入手できます。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織の両方が、オープンという私たちの価値観を受け入れ、受け入れています。

らしさ、コミュニティ、卓越性、そしてユーザーデータのプライバシー。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.00799: CADENA: Stepwise CAD Reverse Engineering

Skip to main content
System maintenance August 4th and 5th
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computer Vision and Pattern Recognition
[Submitted on 1 Aug 2026]
Title: CADENA: Stepwise CAD Reverse Engineering
Abstract: Computer-Aided Design (CAD) underpins modern engineering, yet converting existing shapes into editable models still demands substantial expert effort. Most AI systems emit the entire CAD program in a single pass, never inspecting the intermediate geometry. In contrast, human engineers build a part feature by feature, checking after each operation what remains to be modeled. We introduce CADENA (Spanish for "chain"), a model that reconstructs a 3D mesh as a parametric CAD program, growing its sequence of operations one at a time and comparing the target with the currently predicted geometry at every step. We also address the lack of benchmarks for evaluating reverse-engineering methods on mechanical parts, introducing CADENA-Bench, a benchmark that measures performance across categories of mechanical parts. CADENA outperforms prior methods on CADENA-Bench and on the DeepCAD, Fusion 360, and MCB datasets. Code is available at this https URL , model weights at this https URL , and CADENA-Bench at this https URL .
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
