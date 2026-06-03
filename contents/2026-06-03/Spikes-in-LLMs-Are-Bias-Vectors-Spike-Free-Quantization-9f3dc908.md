---
source: "https://arxiv.org/abs/2606.02288"
hn_url: "https://news.ycombinator.com/item?id=48368631"
title: "Spikes in LLMs Are Bias Vectors: Spike-Free Quantization"
article_title: "[2606.02288] Massive Spikes in LLMs are Bias Vectors: Mechanistic Uncovering and Spike-Free Quantization"
author: "sbulaev"
captured_at: "2026-06-03T00:48:52Z"
capture_tool: "hn-digest"
hn_id: 48368631
score: 3
comments: 0
posted_at: "2026-06-02T11:07:19Z"
tags:
  - hacker-news
  - translated
---

# Spikes in LLMs Are Bias Vectors: Spike-Free Quantization

- HN: [48368631](https://news.ycombinator.com/item?id=48368631)
- Source: [arxiv.org](https://arxiv.org/abs/2606.02288)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T11:07:19Z

## Translation

タイトル: LLM のスパイクはバイアス ベクトルである: スパイクのない量子化
記事のタイトル: [2606.02288] LLM の大規模なスパイクはバイアス ベクトルです: 機械的な解明とスパイクのない量子化
説明: arXiv 論文 2606.02288 の要約ページ: LLM における大規模スパイクはバイアス ベクトルである: 機械的解明とスパイクフリー量子化

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.02288
ヘルプ |高度な検索
コンピューターサイエンス > 機械学習
[2026 年 6 月 1 日に提出]
タイトル: LLM の巨大なスパイクはバイアス ベクトルである: メカニズムの解明とスパイクのない量子化
要約: 大規模言語モデル (LLM) における大規模なアクティベーション スパイクにより、ダイナミック レンジが拡張され、量子化が大幅に低下します。以前の仮説はこれらを高レベルのスカラー バイアスとして特徴付けていますが、我々はこれらはスパイクを運ぶトークンの厳密な構造ベクトル バイアスのスカラー中間体にすぎないと主張します。これらのトークンは正規化後に定数ベクトルに収束し、アテンションシンクと値状態ドレインメカニズムを駆動することを示します。これは、射影重みの調整を分析することによって幾何学的に実証されます。$W_K$ はベクトルを対照的に増幅し、$W_Q$ は意味論的トークンをそれに向かって整列させ、$W_V$ はベクトルをスペクトルヌル空間に射影します。さらに、このモデルは、低周波数帯域とコヒーレントなチャネル ペアを利用して、回転位置埋め込み (RoPE) 摂動を「回転安定性のゾーン」に局所化することで、これらの構造バイアスを積極的に保存していることを明らかにしました。これを活用して、事前に計算されたテンプレート ベクトルを介してスパイクをクランプし、その機能を復元するポストトレーニング量子化 (PTQ) フレームワークである INSERTQUANT を提案します。これにより、アクティベーションが厳密にスパイクフリーになり、忠実度の高い堅牢な低ビット量子化が可能になります。 INSERTQUANT は、LLM 上で最先端のテンソルごとの量子化手法と同等の機能を実現し、テキストを超えて ViT などの他のモダリティに独自に一般化します。
もっと学ぶために集中する
arXiv-iss

DataCite 経由で DOI を使用しました (登録保留中)
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

Abstract page for arXiv paper 2606.02288: Massive Spikes in LLMs are Bias Vectors: Mechanistic Uncovering and Spike-Free Quantization

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.02288
Help | Advanced Search
Computer Science > Machine Learning
[Submitted on 1 Jun 2026]
Title: Massive Spikes in LLMs are Bias Vectors: Mechanistic Uncovering and Spike-Free Quantization
Abstract: Massive activation spikes in Large Language Models (LLMs) severely degrade quantization by stretching dynamic ranges. While prior hypotheses characterize these as high-level scalar biases, we argue that they are merely the scalar intermediates of rigid, structural vector biases in the spike-carrying tokens. We show that these tokens converge to constant vectors after normalization that drive the attention sink and value-state drain mechanisms. We geometrically substantiate this by analyzing the coordination of projection weights: $W_K$ contrastively amplifies the vector, $W_Q$ aligns semantic tokens toward it, and $W_V$ projects it into the spectral null-space. Furthermore, we reveal that the model actively preserves these structural biases against Rotary Positional Embedding (RoPE) perturbations by localizing them in "zones of rotational stability" utilizing low-frequency bands and coherent channel pairs. Leveraging this, we propose INSERTQUANT, a post-training quantization (PTQ) framework that clamps spikes and restores their function via pre-computed template vectors. This renders activations strictly spike-free, enabling robust low-bit quantization with high fidelity. INSERTQUANT achieves parity with state-of-the-art per-tensor quantization methods on LLMs and uniquely generalizes beyond text to other modalities such as ViTs.
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
