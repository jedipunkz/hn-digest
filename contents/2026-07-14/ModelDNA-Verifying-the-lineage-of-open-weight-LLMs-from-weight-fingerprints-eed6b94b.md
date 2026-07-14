---
source: "https://arxiv.org/abs/2607.10617"
hn_url: "https://news.ycombinator.com/item?id=48903868"
title: "ModelDNA: Verifying the lineage of open-weight LLMs from weight fingerprints"
article_title: "[2607.10617] modelDNA: Calibrated Lineage Verification and Merge Decomposition from Sampled Weight Fingerprints"
author: "saadaamir14"
captured_at: "2026-07-14T09:47:47Z"
capture_tool: "hn-digest"
hn_id: 48903868
score: 2
comments: 0
posted_at: "2026-07-14T08:50:13Z"
tags:
  - hacker-news
  - translated
---

# ModelDNA: Verifying the lineage of open-weight LLMs from weight fingerprints

- HN: [48903868](https://news.ycombinator.com/item?id=48903868)
- Source: [arxiv.org](https://arxiv.org/abs/2607.10617)
- Score: 2
- Comments: 0
- Posted: 2026-07-14T08:50:13Z

## Translation

タイトル: ModelDNA: 重量フィンガープリントからのオープンウェイト LLM の系統の検証
記事のタイトル: [2607.10617] modelDNA: サンプリングされた体重フィンガープリントからの校正された系統検証とマージ分解
説明: arXiv 論文 2607.10617 の要約ページ: modelDNA: サンプリングされた体重フィンガープリントからの校正された系統検証とマージ分解

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
コンピューターサイエンス > 機械学習
[2026 年 7 月 12 日に提出]
タイトル: modelDNA: サンプリングされた体重フィンガープリントからの校正された系統検証とマージ分解
要約: オープンウェイト言語モデルの系統グラフは自己報告です。Hugging Face のbase_model メタデータ フィールドはオプションで未検証であり、ハブ モデルの 60% 以上は親子関係をまったく文書化していません。重みから系統を検出する方法は研究文献に存在しますが、それぞれが 1 つの信号と 1 つの実験に関連付けられた紙のコードとして出荷されます。出所に関する論争が勃発すると、分析は手作業でやり直されます。このレポートでは、modelDNA について説明します。このツールは、(7B モデルの完全な 15 GB ダウンロードの代わりに) 約 100 ～ 300 MB の範囲の HTTP 読み取りからモデルのフィンガープリントを取得し、そのフィンガープリントを 4 つの公開された信号ファミリーにわたる基礎モデルの参照データベースと比較し、確信のあるエラーよりも正直な棄権を優先して、校正された確率を持つ 8 つの判定クラスの 1 つを返します。組織で文書化された親子関係を持つ 15 個の実際のハブ モデルのベンチマークでは、8 個の候補塩基 (13 個の陽性、107 個のハード ネガティブ) に対して判定され、システムは AUROC 1.0、レポートしきい値での偽陽性ゼロ、および 13/13 の正しい上位 1 親帰属を達成しました。このレポートの 2 番目の貢献はマージ分解です。すべての主流の重みマージ手法はテンソルごとに (ほぼ) 線形であり、フィンガープリントのサンプル位置はテンソル ID の決定論的な関数であるため、マージされたモデルのフィンガープリントはその親のフィンガープリントの同じ線形結合になります。したがって、混合重量は、sum-to-one 計算によって指紋のみから復元できます。

最小二乗法を行います。公開されたマージキット構成をグラウンド トゥルースとして使用するマージに対して、このメソッドは、フィンガープリントを超えるウェイトをダウンロードすることなく、r = 0.999 での slerp マージのレイヤー補間曲線と、dare_ties マージの混合ウェイトを公開値の 0.011 以内に回復します。 55 モデルのすべてのフィンガープリント、ベンチマーク、および推定系統グラフは公開されており、オフラインで再現可能です。
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

Abstract page for arXiv paper 2607.10617: modelDNA: Calibrated Lineage Verification and Merge Decomposition from Sampled Weight Fingerprints

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
Computer Science > Machine Learning
[Submitted on 12 Jul 2026]
Title: modelDNA: Calibrated Lineage Verification and Merge Decomposition from Sampled Weight Fingerprints
Abstract: The lineage graph of open-weight language models is self-reported: Hugging Face's base_model metadata field is optional and unverified, and over 60% of Hub models document no parentage at all. Methods for detecting lineage from weights exist in the research literature, but each ships as paper code tied to one signal and one experiment; when a provenance dispute breaks, the analysis is redone by hand. This report describes modelDNA, a tool that fingerprints a model from roughly 100-300 MB of ranged HTTP reads (instead of a full 15 GB download for a 7B model), compares the fingerprint against a reference database of foundation models across four published signal families, and returns one of eight verdict classes with a calibrated probability, preferring honest abstention to confident error. On a benchmark of 15 real Hub models with org-documented parentage, judged against 8 candidate bases (13 positives, 107 hard negatives), the system achieves AUROC 1.0, zero false positives at its reporting threshold, and 13/13 correct top-1 parent attribution. The report's second contribution is merge decomposition. Every mainstream weight-merging method is (near-)linear per tensor, and fingerprint sample positions are deterministic functions of tensor identity, so a merged model's fingerprint is the same linear combination of its parents' fingerprints. Mixture weights can therefore be recovered from fingerprints alone by sum-to-one constrained least squares. Against merges with published mergekit configurations as ground truth, the method recovers a slerp merge's layer-interpolation curves at r = 0.999 and a dare_ties merge's mixture weights to within 0.011 of the published values, without downloading any weights beyond the fingerprints. All fingerprints, benchmarks, and the inferred lineage graph of 55 models are public and reproducible offline.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
