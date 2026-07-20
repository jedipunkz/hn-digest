---
source: "https://www.orafrontier.com/"
hn_url: "https://news.ycombinator.com/item?id=48982830"
title: "A new ML compiler to run 70B+ LLMs on consumer GPUs with <1% accuracy loss"
article_title: "Private AI, built for your hardware | Ora Frontier"
author: "marcelbuilds"
captured_at: "2026-07-20T19:26:17Z"
capture_tool: "hn-digest"
hn_id: 48982830
score: 2
comments: 1
posted_at: "2026-07-20T18:25:10Z"
tags:
  - hacker-news
  - translated
---

# A new ML compiler to run 70B+ LLMs on consumer GPUs with <1% accuracy loss

- HN: [48982830](https://news.ycombinator.com/item?id=48982830)
- Source: [www.orafrontier.com](https://www.orafrontier.com/)
- Score: 2
- Comments: 1
- Posted: 2026-07-20T18:25:10Z

## Translation

タイトル: コンシューマ GPU で 1% 未満の精度損失で 700 億以上の LLM を実行する新しい ML コンパイラ
記事のタイトル: ハードウェア向けに構築されたプライベート AI |オラフロンティア
説明: 自分のコンピューターでプライベート AI を実行します。 Ora Frontier は、ファイル、プロンプト、ワークフローを維持できるように、有能なローカル モデルをコンパイルして実行します。

記事本文:
ハードウェア向けに構築されたプライベート AI | Ora Frontier 製品リソース 企業価格 早期アクセスを入手 ハードウェア向けに構築された Private AI に参加してください。
Ora Frontier は、ハードウェア用のプライベート AI ソフトウェアです。自分のコンピューター上で実行される有能なローカル モデルなので、ファイル、プロンプト、ワークフローは手元に残ります。
仕組みを確認する デフォルトで非公開
選択したファイルとローカル ワークフローをコンピューターから離れる必要はありません。
Ora Core はマシンで実行できるものを識別し、Ora Pulse は適切なランタイムを準備します。
ファイル、書き込み、コード、調査、内部ワークフローにローカル AI を使用します。
能力を維持してください。
知性を家に持ち帰ってください。
Ora Frontier は、ユーザーの作業に近い場所で有能なモデルを実行するため、ユーザーのファイル、コンテキスト、コンピューティングはユーザーが制御する境界内に留まります。
オラフロンティア - ローカルセッション >
ワークスペース ファイルを開いています...ローカル
70B モデル プロファイルを読み込み中...準備完了
推論を実行しています...このデバイス
実行境界: このデバイス
エンジンは1つ。 AI を実行する必要があるあらゆる場所向けに構築されています。
あなたのパソコンのためのプライベートAI。
制御された環境全体に AI を導入して管理します。
端末からローカルAIを制御します。
同じ基礎です。どこに着地するかに合わせて調整されています。
共有実行層は、個人のハードウェア、管理対象フリート、開発者のワークフローに適応します。
ハードウェア対応の実行 利用可能なシリコンとメモリに合わせて準備されています。
プライベート実行 コンテキストは、ユーザーが制御する境界内に留まります。
統合モデル パッケージ化 サポートされている環境全体で 1 つのモデル システム。
あなたのコンピュータ。可能な限り最高のモデルです。
Ora Core は、使用可能なハードウェアを検出し、サポートされているモデル プロファイルを選択し、最適化されたローカル ランタイムを準備します。高度なコントロールは、使い始めたときではなく、必要なときに利用できます。
ora ハードウェア スキャン $ システム MacBook Pro (Apple M4) メモリ 16 GB ユニファイド ストレージ 412 GB 無料チェック

システム チェック モデルが利用可能 一致するハードウェア プロファイル モデル qwen3-250b プロファイル バランスのとれたローカル ✓ すぐに実行可能 常駐メモリの比較 700 億のパラメータ モデルは、標準ランタイムで 48 GB を使用し、Ora Frontier でコンパイルすると 14 GB を使用し、モデルの機能を維持しながら 34 GB をシステムに返します。推論時の常駐メモリ 同じ 70B モデル。低いほど良いです。常駐メモリ (GB) 0 12 24 36 48 48 GB 14 GB 標準ランタイム 48 GB 常駐 Frontier コンパイル済み 14 GB 常駐ランタイム 34 GB 返されるメモリが少なくなります。さらにモデル。謎はありません。
Frontier は、モデル コンポーネントをハードウェア対応ランタイムにコンパイルします。これは、ユーザーが詳細を管理することなく、展開の手間を軽減するように設計されています。
ターミナルを開く前に便利です。
ドキュメントをローカルで検索して作業します。
プライベートの起草、分析、研究サポート。
src approutesserver.ts utils Types.ts server.ts @@ -42,2 +42,3 @@ 42 import async function run() { 43 - return cloud.run(input) 44 + const model = local.load() 45 + return model.run(input) 46 } ソフトウェアのビルド
プロジェクトでローカルのコーディング アシスタントを使用します。
ネットワークができない場合でも続行してください。
有能な AI はハードウェアに組み込まれています。
ハードウェア向けに構築されたプライベート AI。

## Original Extract

Run private AI on your own computer. Ora Frontier compiles and runs capable local models so your files, prompts, and workflows stay with you.

Private AI, built for your hardware | Ora Frontier Products Resources Company Pricing Get early access Join Private AI , built for your hardware.
Ora Frontier is private AI software for your hardware — capable local models that run on your own computer so your files, prompts, and workflows stay with you.
See how it works Private by default
Your selected files and local workflows do not need to leave your computer.
Ora Core identifies what your machine can run and Ora Pulse prepares the right runtime.
Use local AI for files, writing, code, research, and internal workflows.
Keep the capability.
Bring the intelligence home.
Ora Frontier runs capable models closer to your work, so your files, context, and compute stay within a boundary you control.
ora frontier - local session >
opening workspace files... local
loading 70B model profile... ready
running inference... this device
execution boundary: this device
One engine. Built for every place AI needs to run.
Private AI for your personal computer.
Deploy and govern AI across controlled environments.
Control local AI from the terminal.
Same foundation. Tuned for where it lands.
A shared execution layer adapts to personal hardware, managed fleets, and developer workflows.
Hardware-aware execution Prepared for the silicon and memory available.
Private execution Context stays inside the boundary you control.
Unified model packaging One model system across supported environments.
Your computer. Its best possible model.
Ora Core detects your available hardware, selects a supported model profile, and prepares an optimized local runtime. Advanced controls are there when you want them, not when you are getting started.
ora hardware scan $ system MacBook Pro (Apple M4) memory 16 GB unified storage 412 GB free checking system checking models available matching hardware profiles model qwen3-250b profile balanced local ✓ ready to run Resident memory comparison A 70 billion parameter model uses 48 gigabytes with a standard runtime and 14 gigabytes when compiled with Ora Frontier, returning 34 gigabytes to the system while preserving model capability. Resident memory at inference Same 70B model. Lower is better. Resident memory (GB) 0 12 24 36 48 48 GB 14 GB Standard runtime 48 GB resident Frontier compiled 14 GB resident Runtime 34 GB returned Less memory. More model. No mystery.
Frontier compiles model components into a hardware-aware runtime designed to reduce deployment friction without making the user manage the details.
Useful before you open a terminal.
Search and work through documents locally.
Private drafting, analysis, and research support.
src app routes server.ts utils types.ts server.ts @@ -42,2 +42,3 @@ 42 export async function run() { 43 - return cloud.run(input) 44 + const model = local.load() 45 + return model.run(input) 46 } Build software
Use a local coding assistant with your projects.
Continue when your network cannot.
Capable AI belongs on your hardware.
Private AI, built for your hardware.
