---
source: "https://zenodo.org/records/21753002"
hn_url: "https://news.ycombinator.com/item?id=49140080"
title: "Persistent State Machines: LLM Attention with INT4 In-Memory Cells"
article_title: "Persistent State Machines: Complete Mathematical Proofs and Vivado Implementation Synthesis (Version 8.0) | Zenodo"
author: "yusuke_esaka"
captured_at: "2026-08-02T01:50:01Z"
capture_tool: "hn-digest"
hn_id: 49140080
score: 3
comments: 0
posted_at: "2026-08-02T01:01:45Z"
tags:
  - hacker-news
  - translated
---

# Persistent State Machines: LLM Attention with INT4 In-Memory Cells

- HN: [49140080](https://news.ycombinator.com/item?id=49140080)
- Source: [zenodo.org](https://zenodo.org/records/21753002)
- Score: 3
- Comments: 0
- Posted: 2026-08-02T01:01:45Z

## Translation

タイトル: 永続ステート マシン: INT4 インメモリ セルを使用した LLM アテンション
記事のタイトル: 永続的ステート マシン: 完全な数学的証明と Vivado 実装合成 (バージョン 8.0) |ゼノド
説明: 永続ステート マシン: 完全な数学的証明と Vivado 実装合成 (バージョン 8.0) 要旨 永続ステート マシン (PSM) を介した大規模言語モデルのアテンション演算子のための正式な離散フレームワークを紹介します。計算は固定メモリ内セルにブロードキャストされます。
[切り捨てられた]

記事本文:
永続的ステート マシン: 完全な数学的証明と Vivado インプリメンテーション合成 (バージョン 8.0) |ゼノド
メインにスキップ
古いブラウザを使用しています。エクスペリエンスを向上させるためにブラウザをアップグレードしてください。
永続的ステート マシン: 完全な数学的証明と Vivado インプリメンテーション合成 (バージョン 8.0)
1.
コスモス行政書士事務所・独立研究員
永続的ステート マシン: 完全な数学的証明と Vivado インプリメンテーション合成 (バージョン 8.0)
要約 永続的ステートマシン (PSM) を介した大規模言語モデルのアテンション演算子のための正式な離散フレームワークを紹介します。計算は、局所的な決定論的な状態遷移を評価する固定メモリ内セルにブロードキャストされます。完全な数学的証明は、量子化誤差限界、明示的な有界ロジット仮定に基づく具体的な多相離散ソフトマックス構築、空間因数分解による決定論的有限オートマトン等価性、および DSPACE(O(n)) のメンバーシップに対して与えられます。
私たちは、次の 2 つの異なる評価フローを通じて、現代のプログラマブル ロジック ファブリック上でのアーキテクチャの実装の実現可能性を検証します。
低消費電力評価 (Zynq-7000 xc7z020) : フル 1024 セル アレイ (d=128) のアウト オブ コンテキスト (OOC) ブロック デザインの実装。配線後のネットリストに、配線後の機能シミュレーションから取得したスイッチング アクティビティ交換フォーマット (SAIF) ファイルで注釈を付けることにより、ハードウェア レベルの自己活性化ゲートがアクティブな動的スイッチングをまばらなセルに制限することを示します。コア ロジックの動的電力は 1.0 mW 未満と推定され、正規化された動的エネルギーは 3.81 × 10^-5 pJ/op となります。
システム オン チップ PCIe 統合 (UltraScale+ xcvu9p) : システム レベルの互換性を検証するには、256 セルのサブアレイ (

統合された 1 ヘッド アテンション システム) が完全な SoC 内に統合されています。これには、AMBA AXI4 インターコネクトと AMD Xilinx PCIe Gen3 x1 ブリッジ (XDMA v4.2) が含まれます。完全な SoC は、AWS クラウド FPGA 開発者環境下で、62.5 MHz (最悪のネガティブ スラック WNS = +1.854 ns) の統一システム クロックで正常にタイミングを終了しました。 SoC はデバイスのロジック スライスの 0.67% と DSP ブロックの 0.00% を占有するだけであり、高いスケーリングの可能性を示しています。
1,000 を超えるランダム ベクトルを使用した関数シミュレーションにより、固定小数点ソフトウェア参照とのビット正確な一致が確認されました。すべてのエネルギー数値は、合成ロジックのシミュレーション ベースのツール推定値です。物理的な FPGA ボードの測定は行われておらず、システム レベルの外部メモリのエネルギーは厳密に除外されています。
特願2026-177318（特許出願中）
029_Zenodo_ASMA_Paper_v8_0_FullProofMaster.pdf
統計の収集方法の詳細....
10.5281/ゼノド.21753002
マークダウン
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21753002.svg)](https://doi.org/10.5281/zenodo.21753002)
再構造化されたテキスト
.. 画像:: https://zenodo.org/badge/DOI/10.5281/zenodo.21753002.svg
:target: https://doi.org/10.5281/zenodo.21753002
HTML
<a href="https://doi.org/10.5281/zenodo.21753002"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21753002.svg" alt="DOI"></a>
画像URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21753002.svg
ターゲット URL
https://doi.org/10.5281/zenodo.21753002
リソースの種類
プレプリント
出版社
ゼノド
権利
提供元
CERN データセンターと InvenioRDM
このサイトでは Cookie を使用しています。 Cookieの使用方法について詳しくはこちらをご覧ください

## Original Extract

Persistent State Machines: Complete Mathematical Proofs and Vivado Implementation Synthesis (Version 8.0) ABSTRACT We present a formal discrete framework for attention operators in Large Language Models via Persistent State Machines (PSMs). Computation is broadcast to stationary in-memory cells that
[truncated]

Persistent State Machines: Complete Mathematical Proofs and Vivado Implementation Synthesis (Version 8.0) | Zenodo
Skip to main
You are using an outdated browser. Please upgrade your browser to improve your experience.
Persistent State Machines: Complete Mathematical Proofs and Vivado Implementation Synthesis (Version 8.0)
1.
Cosmos Administrative Scrivener Office & Independent Researcher
Persistent State Machines: Complete Mathematical Proofs and Vivado Implementation Synthesis (Version 8.0)
ABSTRACT We present a formal discrete framework for attention operators in Large Language Models via Persistent State Machines (PSMs). Computation is broadcast to stationary in-memory cells that evaluate local deterministic state transitions. Complete mathematical proofs are given for quantization error bounds, a concrete multi-phase discrete Softmax construction under an explicit bounded-logits assumption, deterministic finite-automaton equivalence with spatial factorization, and membership in DSPACE(O(n)).
We validate the implementation feasibility of the architecture on contemporary programmable logic fabric through two distinct evaluation flows:
Low-Power Evaluation (Zynq-7000 xc7z020) : A full 1024-cell array (d=128) out-of-context (OOC) block design implementation. By annotating the post-route netlist with a Switching Activity Interchange Format (SAIF) file obtained from a post-route functional simulation, we demonstrate that hardware-level self-activation gating restricts active dynamic switching to sparse cells. The dynamic power of the core logic is estimated below 1.0 mW, yielding a normalized dynamic energy of 3.81 × 10^-5 pJ/op.
System-on-Chip PCIe Integration (UltraScale+ xcvu9p) : To verify system-level compatibility, a 256-cell sub-array (representing an integrated 1-head attention system) is integrated within a full SoC. This includes AMBA AXI4 interconnects and an AMD Xilinx PCIe Gen3 x1 Bridge (XDMA v4.2). The complete SoC successfully closed timing at a unified system clock of 62.5 MHz (Worst Negative Slack WNS = +1.854 ns) under the AWS Cloud FPGA Developer environment. The SoC occupies only 0.67% of the device's logic slices and 0.00% of DSP blocks, showing high scaling potential.
Functional simulation with more than one thousand random vectors confirmed bit-exact agreement with a fixed-point software reference. All energy figures are simulation-based tool estimates for the synthesised logic; no physical FPGA board measurement was performed, and system-level external memory energy is strictly excluded.
Japanese Patent Application No. 2026-177318 (Patent Pending)
029_Zenodo_ASMA_Paper_v8_0_FullProofMaster.pdf
More info on how stats are collected....
10.5281/zenodo.21753002
Markdown
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21753002.svg)](https://doi.org/10.5281/zenodo.21753002)
reStructuredText
.. image:: https://zenodo.org/badge/DOI/10.5281/zenodo.21753002.svg
:target: https://doi.org/10.5281/zenodo.21753002
HTML
<a href="https://doi.org/10.5281/zenodo.21753002"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21753002.svg" alt="DOI"></a>
Image URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21753002.svg
Target URL
https://doi.org/10.5281/zenodo.21753002
Resource type
Preprint
Publisher
Zenodo
Rights
Powered by
CERN Data Centre & InvenioRDM
This site uses cookies. Find out more on how we use cookies
