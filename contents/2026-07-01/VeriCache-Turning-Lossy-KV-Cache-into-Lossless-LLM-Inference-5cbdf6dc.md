---
source: "https://arxiv.org/abs/2605.17613"
hn_url: "https://news.ycombinator.com/item?id=48753225"
title: "VeriCache: Turning Lossy KV Cache into Lossless LLM Inference"
article_title: "[2605.17613] VeriCache: Turning Lossy KV Cache into Lossless LLM Inference"
author: "matt_d"
captured_at: "2026-07-01T21:30:55Z"
capture_tool: "hn-digest"
hn_id: 48753225
score: 1
comments: 0
posted_at: "2026-07-01T21:16:09Z"
tags:
  - hacker-news
  - translated
---

# VeriCache: Turning Lossy KV Cache into Lossless LLM Inference

- HN: [48753225](https://news.ycombinator.com/item?id=48753225)
- Source: [arxiv.org](https://arxiv.org/abs/2605.17613)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T21:16:09Z

## Translation

タイトル: VeriCache: 非可逆 KV キャッシュをロスレス LLM 推論に変える
記事のタイトル: [2605.17613] VeriCache: 非可逆 KV キャッシュをロスレス LLM 推論に変える
説明: arXiv 論文 2605.17613 の要約ページ: VeriCache: Turning Lossy KV Cache into Lossless LLM Inference

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
コンピュータ サイエンス > ハードウェア アーキテクチャ
[2026 年 5 月 17 日に提出]
タイトル: VeriCache: 非可逆 KV キャッシュをロスレス LLM 推論に変える
要約: KV キャッシュのサイズが大きいことが、コンテキスト長が増加する LLM を提供する際の大きなボトルネックになっています。これに応じて、トークンドロップや量子化などの多くの KV キャッシュ圧縮方法が提案されています。ただし、これらのメソッドのほとんどすべては本質的に不可逆的であり、短い出力の精度低下は最小限であるにもかかわらず、より多くのトークンがデコードされるにつれて、その出力はフル KV キャッシュ出力からますます乖離し、コード生成とツール呼び出しで壊滅的な失敗につながります。
フル KV キャッシュ デコードと同じ出力を保証しながら、さまざまな KV キャッシュ圧縮アルゴリズムの高いデコード スループットを大幅に維持する初の推論フレームワークである VeriCache を紹介します。 VeriCache は、圧縮された KV キャッシュを使用してトークンをドラフトし、完全な KV キャッシュと照合して検証します。単なる投機的なデコードのように見えるかもしれませんが、VeriCache では、完全な KV キャッシュを GPU メモリから外して動作させ、検証のためにスワップするオーバーヘッドを最小限に抑えるという重要なシステム課題に対処する必要があります。洞察は 2 つあります。(1) 圧縮 KV デコードはフル KV スワップと並列化できます。これは、1 つは HBM 帯域幅に制限され、もう 1 つは PCIe/ネットワークに制限されるためです。(2) 圧縮 KV キャッシュはフル KV キャッシュと同様の出力を生成することが多く、各フル KV スワップを償却するための長いドラフティング期間が可能になります。
VeriCache は、ロングコンテキストのデコードとリモート プレフィックス キャッシングの両方に適用され、幅広いトークン ドロップおよび量子化メソッドをサポートします。

統一されたコンプレッサー インターフェイスを大まかに作成し、従来の投機的デコードを使用して構成します。実験結果は、VeriCache が同一の出力を生成しながら、フル KV 推論よりも最大 4 倍高いスループットを達成することを示しています。
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

Abstract page for arXiv paper 2605.17613: VeriCache: Turning Lossy KV Cache into Lossless LLM Inference

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
Computer Science > Hardware Architecture
[Submitted on 17 May 2026]
Title: VeriCache: Turning Lossy KV Cache into Lossless LLM Inference
Abstract: The large size of the KV cache has become a major bottleneck for serving LLMs with increasing context lengths. In response, many KV cache compression methods, such as token dropping and quantization, have been proposed. However, almost all of these methods are inherently lossy-despite minimal accuracy degradation for short outputs, their outputs increasingly diverge from full-KV-cache outputs as more tokens are decoded, which leads to catastrophic failures in code generation and tool calling.
We present VeriCache, the first inference framework that ensures the same output as full-KV-cache decoding but largely preserves the high decoding throughput of a range of KV cache compression algorithms. VeriCache uses the compressed KV cache to draft tokens, then verifies them against the full KV cache. While it may seem like just speculative decoding, VeriCache requires addressing a key system challenge to work-keeping the full KV cache out of GPU memory and minimizing the overhead of swapping it in for verification. The insight is two-fold: (1) compressed-KV decoding can be parallelized with full-KV swap, because one is HBM-bandwidth-bound and the other is PCIe/network-bound, and (2) the compressed KV cache often produces output similar to the full KV cache, allowing a long drafting horizon to amortize each full-KV swap.
VeriCache applies to both long-context decoding and remote prefix caching, supports a broad family of token-dropping and quantization methods through a uniform compressor interface, and composes with traditional speculative decoding. Experimental results show that VeriCache achieves up to 4X higher throughput than full-KV inference while producing identical outputs.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
