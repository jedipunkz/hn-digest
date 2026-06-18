---
source: "https://x86ecosystem.org/resource/ai-compute-extensions-ace-specification/"
hn_url: "https://news.ycombinator.com/item?id=48579913"
title: "[x86] AI Compute Extensions (ACE) Specification"
article_title: "AI Compute Extensions (ACE) Specification – x86 Ecosystem Advisory Group"
author: "matt_d"
captured_at: "2026-06-18T04:35:06Z"
capture_tool: "hn-digest"
hn_id: 48579913
score: 19
comments: 7
posted_at: "2026-06-18T02:32:01Z"
tags:
  - hacker-news
  - translated
---

# [x86] AI Compute Extensions (ACE) Specification

- HN: [48579913](https://news.ycombinator.com/item?id=48579913)
- Source: [x86ecosystem.org](https://x86ecosystem.org/resource/ai-compute-extensions-ace-specification/)
- Score: 19
- Comments: 7
- Posted: 2026-06-18T02:32:01Z

## Translation

タイトル: [x86] AI Compute Extensions (ACE) 仕様
記事のタイトル: AI Compute Extensions (ACE) 仕様 – x86 Ecosystem Advisory Group

記事本文:
AI Compute Extensions (ACE) 仕様 – x86 エコシステム アドバイザリー グループ
コンテンツにスキップ
について
について
検索
AI コンピューティング拡張機能 (ACE) 仕様
このドキュメントでは、計算タスクを高速化するための x86 拡張機能を定義します。最初は、ML ワークロードにとって重要な行列乗算カーネルと低精度データ形式に焦点を当てています。
ACE 拡張機能は、AVX およびスカラー コードを新しい機能で強化する行列乗算プリミティブを定義し、以下を追加します。
ACE レジスタの状態 (タイルおよびブロック スケール レジスタを含む)
AVX レジスタ入力を消費し、タイル レジスタの状態を操作するデータ処理操作
ACE レジスタの状態と AVX レジスタの間でデータを移動するデータ移動操作
システム管理のための状態と操作
ACE は、AVX ベクトルと ACE タイル レジスタ間の緊密な統合を提供し、高計算密度のタイル処理操作と AVX の包括的なデータ処理機能を組み合わせます。
マトリックス アクセラレーションに加えて、AVX10 フレームワークでは多数の専用フォーマット変換操作が提供されます。
© 2026 全著作権所有。プライバシーポリシー

## Original Extract

AI Compute Extensions (ACE) Specification – x86 Ecosystem Advisory Group
Skip to content
About
About
Search
AI Compute Extensions (ACE) Specification
This document defines x86 extensions for accelerating computation tasks, initially focusing on matrix multiplication kernels and reduced precision data formats important to ML workloads.
The ACE extensions define matrix multiplication primitives that augment AVX and scalar code with new capabilities, adding:
ACE register state, including tile and block scale registers
Data processing operations that consume AVX register input and operate on tile register state
Data move operations to move data between ACE register state and AVX registers
State and operations for system management
ACE provides tight integration between AVX vectors and ACE tile registers, combining high compute density tile processing operations with the comprehensive data processing features of AVX.
In addition to matrix acceleration, a number of dedicated format convert operations are provided under the AVX10 framework.
© 2026 All Rights Reserved. Privacy Policies
