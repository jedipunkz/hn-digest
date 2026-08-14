---
source: "https://arxiv.org/abs/2608.12700"
hn_url: "https://news.ycombinator.com/item?id=49301417"
title: "A Contract-Grade Verifier for LLM-Generated GPU Kernels"
article_title: "[2608.12700] A Contract-Grade Verifier for LLM-Generated GPU Kernels, and a Native Blackwell Backward for the Gated-Linear-Recurrence Family"
author: "Jimmc414"
captured_at: "2026-08-14T17:45:38Z"
capture_tool: "hn-digest"
hn_id: 49301417
score: 4
comments: 0
posted_at: "2026-08-14T16:57:14Z"
tags:
  - hacker-news
  - translated
---

# A Contract-Grade Verifier for LLM-Generated GPU Kernels

- HN: [49301417](https://news.ycombinator.com/item?id=49301417)
- Source: [arxiv.org](https://arxiv.org/abs/2608.12700)
- Score: 4
- Comments: 0
- Posted: 2026-08-14T16:57:14Z

## Translation

タイトル: LLM で生成された GPU カーネルの契約グレードの検証ツール
記事のタイトル: [2608.12700] LLM で生成された GPU カーネルの契約グレードの検証ツール、および Gated-Linear-Recurrence ファミリのネイティブ Blackwell Backward
説明: arXiv 論文 2608.12700 の要約ページ: LLM で生成された GPU カーネルの契約グレードの検証者、および Gated-Linear-Recurrence ファミリのネイティブ Blackwell Backward

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 機械学習
[2026 年 8 月 13 日に提出]
タイトル: LLM で生成された GPU カーネル用の契約グレードの検証ツール、および Gated-Linear-Recurrence ファミリ用のネイティブ Blackwell Backward
要約: 言語モデルを使用して GPU カーネルを生成するシステムは、高い正確性率を報告します。これらのレートは、単一の緩やかなテストから得られます。つまり、1 つの固定形状でいくつかのランダムな入力に対してカーネルを実行し、出力が基準に近い場合にそれを受け入れます。カーネルはそのテストに合格しても、黙って間違っている可能性があります。真の答えが NaN または無限大である通常の数値を返すことも、実行ごとに異なることも、形状が変化すると中断することも、参照が fp32 の合計を保持する fp16 に蓄積することもできます。私たちは正確性を適切にチェックする手段を構築します。12 個の敵対的ゲートの契約グレードの検証器です。それぞれが正しいカーネルが満たさなければならないプロパティであり、そのうちのいくつかは許容誤差がないため、しきい値を選択しても失敗を説明できません。この検証者は外部に向けて、公共システム独自のハーネスがすでに正しいものとして受け入れている 2,638 個の機械生成カーネルを監査します。その結果、39.5% がいかなる許容範囲の議論も超えて違反しており、62.1% が少なくとも 1 つの違反を抱えていることがわかりました。この分野の標準テストでは、検証者が拒否したカーネルは 1,487 個受け入れられますが、他の方法では 14 個のみが受け入れられます。私たちは、7/7 ポジティブ コントロール、閾値キャリブレーション スイープ、参照ベンチマーク自体の正確性コードとの 98.5% の一致、および階層化手作業監査の 4 つの独立した方法での発見を擁護します。内部に向けて、検証者は独自のカーネルを判断します。最初のネイティブ Blackwell tcgen05 は、ゲート線形回帰 (GDN) ファミリ向けに逆方向にトレーニングされ、フィールドの逆状態ステージも含まれます。

フォールバックで実行されます。私たちは倍精度の神託に対抗してその正しさを独自に確立し、それを通じて家族 5 人を訓練します。報告されているカーネル生成の進歩の背後にある正確性のシグナルは、数字が示すよりもはるかに弱いため、許容差のない一連の契約がギャップのほとんどを埋めることになります。
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

Abstract page for arXiv paper 2608.12700: A Contract-Grade Verifier for LLM-Generated GPU Kernels, and a Native Blackwell Backward for the Gated-Linear-Recurrence Family

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Machine Learning
[Submitted on 13 Aug 2026]
Title: A Contract-Grade Verifier for LLM-Generated GPU Kernels, and a Native Blackwell Backward for the Gated-Linear-Recurrence Family
Abstract: Systems that generate GPU kernels with language models report high correctness rates. Those rates come from a single loose test: run the kernel on a few random inputs at one fixed shape and accept it if the output is close to a reference. A kernel can pass that test and still be silently wrong. It can return an ordinary number where the true answer is a NaN or an infinity, differ from run to run, break when the shape changes, or accumulate in fp16 where the reference keeps an fp32 total. We build the instrument that checks correctness properly: a contract-grade verifier of twelve adversarial gates, each a property a correct kernel must satisfy, several of them tolerance-free, so no choice of threshold can explain a failure away. Aimed outward, the verifier audits 2,638 machine-generated kernels that a public system's own harness had already accepted as correct. It finds 39.5% broken beyond any tolerance argument and 62.1% carrying at least one violation. The field's standard test accepts 1,487 kernels the verifier rejects, against only 14 the other way. We defend the finding four independent ways: a 7/7 positive control, a threshold-calibration sweep, 98.5% agreement with the reference benchmark's own correctness code, and a stratified hand-audit. Aimed inward, the verifier judges a kernel of our own: the first native Blackwell tcgen05 training backward for the gated-linear-recurrence (GDN) family, including the reverse-state stage the field still runs on a fallback. We establish its correctness independently, against a double-precision oracle, and train five family members through it. The correctness signal behind reported progress in kernel generation is far weaker than the numbers suggest, and a set of tolerance-free contracts would close most of the gap.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
