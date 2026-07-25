---
source: "https://zenodo.org/records/21566453"
hn_url: "https://news.ycombinator.com/item?id=49049884"
title: "Persistent State Machine: Breaking the von Neumann Memory Wall for LLM Attention"
article_title: "Persistent State Machine: A Formal Computational Paradigm Beyond The Von Neumann Model for High-Sparsity LLM Attention Acceleration | Zenodo"
author: "yusuke_esaka"
captured_at: "2026-07-25T18:53:33Z"
capture_tool: "hn-digest"
hn_id: 49049884
score: 1
comments: 0
posted_at: "2026-07-25T17:53:15Z"
tags:
  - hacker-news
  - translated
---

# Persistent State Machine: Breaking the von Neumann Memory Wall for LLM Attention

- HN: [49049884](https://news.ycombinator.com/item?id=49049884)
- Source: [zenodo.org](https://zenodo.org/records/21566453)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T17:53:15Z

## Translation

タイトル: 永続的ステート マシン: LLM アテンションのためのフォン ノイマン メモリの壁を突破する
記事のタイトル: 永続ステート マシン: 高スパース性 LLM アテンション アクセラレーションのためのフォン ノイマン モデルを超えた正式な計算パラダイム |ゼノド
説明: フォン ノイマン メモリの壁 (計算とデータ ストレージの間の帯域幅とエネルギー ギャップ) が、大規模言語モデル (LLM) 推論の主要なボトルネックになっています。この論文では、計算が次のようにブロードキャストされる形式的な計算パラダイムである永続ステート マシン (PSM) を紹介します。
[切り捨てられた]

記事本文:
永続的ステート マシン: 高スパース性 LLM アテンション アクセラレーションのためのフォン ノイマン モデルを超えた正式な計算パラダイム |ゼノド
メインにスキップ
古いブラウザを使用しています。エクスペリエンスを向上させるためにブラウザをアップグレードしてください。
永続的ステート マシン: 高スパース性 LLM アテンション アクセラレーションのためのフォン ノイマン モデルを超えた正式な計算パラダイム
1.
コスモス行政書士事務所・独立研究員
フォン ノイマン メモリの壁 (計算とデータ ストレージの間の帯域幅とエネルギー ギャップ) が、大規模言語モデル (LLM) 推論の主要なボトルネックになっています。この論文では、状態遷移をローカルに評価する固定メモリ内セルへの命令として計算がブロードキャストされる形式的な計算パラダイムである永続ステート マシン (PSM) を紹介します。 PSM を数学的な 6 タプルとして定義し、決定論的有限オートマトン (DFA) が含まれていることを証明し、PSPACE 完全な表現力を表す物理メモリ制約 L <= N の下で線形有界オートマトン (LBA) と同等であることを確立します。
我々は、LLM KV キャッシュ アテンションのための PSM を実装する合成可能なシリコン リファレンス アーキテクチャであるアクティブ ステート マシン メモリ アーキテクチャ (ASMA) を紹介します。 ASMA は、GPU ベースラインに対してシステム バス トラフィックを 99.47% 削減し、正味ステップ エネルギーを 99.44% 削減しながら、完全なソフトマックス(QK^T)V 出力を生成します。 NVIDIA GeForce RTX 3090 GPU ベースラインと比較して評価すると、ASMA は、極端なロングコンテキスト マルチバッチ デコード (B=32、N=131k) のもとで、2,129.19 倍という驚異的な物理速度向上を達成します。
注：特願2026-177318（特許出願中）により保護されています。
008_TechRxiv_ASMA_Final_20260726.pdf
統計の収集方法の詳細....
10.5281/zenodo.21566453
マークダウン
[![DOI](https://zenodo.org/badg

e/DOI/10.5281/zenodo.21566453.svg)](https://doi.org/10.5281/zenodo.21566453)
再構造化されたテキスト
.. 画像:: https://zenodo.org/badge/DOI/10.5281/zenodo.21566453.svg
:target: https://doi.org/10.5281/zenodo.21566453
HTML
<a href="https://doi.org/10.5281/zenodo.21566453"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21566453.svg" alt="DOI"></a>
画像URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21566453.svg
ターゲット URL
https://doi.org/10.5281/zenodo.21566453
リソースの種類
プレプリント
出版社
ゼノド
権利
提供元
CERN データセンターと InvenioRDM
このサイトでは Cookie を使用しています。 Cookieの使用方法について詳しくはこちらをご覧ください

## Original Extract

The von Neumann memory wall---the bandwidth and energy gap between computation and data storage---has become the dominant bottleneck of Large Language Model (LLM) inference. This paper introduces the Persistent State Machine (PSM), a formal computational paradigm where computation is broadcast as in
[truncated]

Persistent State Machine: A Formal Computational Paradigm Beyond The Von Neumann Model for High-Sparsity LLM Attention Acceleration | Zenodo
Skip to main
You are using an outdated browser. Please upgrade your browser to improve your experience.
Persistent State Machine: A Formal Computational Paradigm Beyond The Von Neumann Model for High-Sparsity LLM Attention Acceleration
1.
Cosmos Administrative Scrivener Office & Independent Researcher
The von Neumann memory wall---the bandwidth and energy gap between computation and data storage---has become the dominant bottleneck of Large Language Model (LLM) inference. This paper introduces the Persistent State Machine (PSM), a formal computational paradigm where computation is broadcast as instructions to stationary in-memory cells that evaluate state transitions locally. We define PSM as a mathematical 6-tuple, prove its containment of Deterministic Finite Automata (DFA), and establish its equivalence to Linear Bounded Automata (LBA) under physical memory constraints L <= N, representing PSPACE-complete expressive power.
We present the Active State-machine Memory Architecture (ASMA), a synthesizable silicon reference architecture implementing PSM for LLM KV-cache attention. ASMA produces the full softmax(QK^T)V output while reducing system bus traffic by 99.47% and net step energy by 99.44% against GPU baselines. Evaluated against an NVIDIA GeForce RTX 3090 GPU baseline, ASMA achieves an extraordinary 2,129.19x physical speedup under extreme long-context multi-batch decoding (B=32, N=131k).
Note: Protected under Japanese Patent Application No. 2026-177318 (Patent Pending).
008_TechRxiv_ASMA_Final_20260726.pdf
More info on how stats are collected....
10.5281/zenodo.21566453
Markdown
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21566453.svg)](https://doi.org/10.5281/zenodo.21566453)
reStructuredText
.. image:: https://zenodo.org/badge/DOI/10.5281/zenodo.21566453.svg
:target: https://doi.org/10.5281/zenodo.21566453
HTML
<a href="https://doi.org/10.5281/zenodo.21566453"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21566453.svg" alt="DOI"></a>
Image URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21566453.svg
Target URL
https://doi.org/10.5281/zenodo.21566453
Resource type
Preprint
Publisher
Zenodo
Rights
Powered by
CERN Data Centre & InvenioRDM
This site uses cookies. Find out more on how we use cookies
