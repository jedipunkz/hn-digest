---
source: "https://github.com/Zyora-Dev/zse/releases/tag/v2.0.0"
hn_url: "https://news.ycombinator.com/item?id=48370827"
title: "Show HN: We built an LLM inference engine in pure Python – no PyTorch, no Triton"
article_title: "Release v2.0.0 — The Own Everything Release · Zyora-Dev/zse · GitHub"
author: "zyoraclub"
captured_at: "2026-06-03T00:43:38Z"
capture_tool: "hn-digest"
hn_id: 48370827
score: 2
comments: 0
posted_at: "2026-06-02T14:34:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: We built an LLM inference engine in pure Python – no PyTorch, no Triton

- HN: [48370827](https://news.ycombinator.com/item?id=48370827)
- Source: [github.com](https://github.com/Zyora-Dev/zse/releases/tag/v2.0.0)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T14:34:48Z

## Translation

タイトル: Show HN: PyTorch も Triton も使用せず、純粋な Python で LLM 推論エンジンを構築しました
記事のタイトル: リリース v2.0.0 — Own Everything リリース · Zyora-Dev/zse · GitHub
説明: オープンソースの世界が独自に構築した推論エンジン。 - リリース v2.0.0 — 独自のすべてのリリース · Zyora-Dev/zse

記事本文:
リリース v2.0.0 — 独自のすべてのリリース · Zyora-Dev/zse · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ゾーラ・デヴ
/
ぜー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
v2.0.0 — 独自のすべてのリリース
ろ過

えーっと
読み込み中
申し訳ありませんが、問題が発生しました。
読み込み中にエラーが発生しました。このページをリロードしてください。
v2.0.0
55a9742
ZSE v2.0.0 — 「すべてを所有」リリース
完全な書き直し。サードパーティへの依存関係はゼロです。 PyTorch、Triton、トランスフォーマー、ビットサンドバイトはありません。 Pure-Python カーネル コンパイラは、CUDA C、HIP C、および Metal Shading Language を直接出力します。
ヘッドライン番号 (A100-80GB の Qwen2.5-14B INT4 と vLLM AWQ INT4)
pip インストール zse-engine
zseserve model.zse --port 8000
または、カーネル コンパイラをスタンドアロンで実行します。
pip インストール zse-compiler
このリリースの内容
ZSE カーネル コンパイラ — @zse.kernel Python DSL → CUDA / HIP / Metal。ワープ プリミティブ、ベクトル化メモリ、ブロック リダクション、タイリング、フュージョン、WMMA、CDNA3 MFMA マトリックス コア、自動チューニング。
.zse モデル フォーマット v2 — 事前に量子化された INT4/INT8/FP16、mmap 対応、C アクセラレーションによる量子化 (~600 倍高速)。 Llama / Mistral / Qwen2 / Gemma2 / Phi3 用アダプター。
独自の PagedAttendance — 適応型ブロック サイズ設定、トークン レベルのエビクション、FNV-1a 重複除去、COW フォーク。
ZStreamer — 連続バッチ処理、細分化されたプレフィル/デコード、チャンク化されたプレフィル、投機的デコード (N グラム + セルフドラフト)。
オーケストレーター — 統合 VRAM アロケーター、MI300X 上の 29 GPU カーネル、CUDA グラフ + HIP グラフ、LoRA ホットスワップ。
サーバー - OpenAI 互換 API、API キー認証、レート制限、SQLite ストア、組み込み RAG ( /v1/rag/* )、Web ダッシュボード。
RAG — BM25 + TF-IDF + 高密度エンベディング (ロードされた LLM 経由、追加の深度ゼロ) + 相互ランク融合 + LLM クロスエンコーダーのリランク。
Tensor Parallelism — 純粋な ctypes NCCL/RCCL ラッパー、マルチプロセス ワーカー。
パッケージ名変更: PyPI の zllm-zse → zse-engine
モジュールの名前変更: zse → zse_engine
.zse 形式 v2 は 1.x と互換性がありません — zse Convert で再変換します
bnb / bitsandbytes バックエンドが削除されました
PyTorch / Triton / トランスフォーマーに依存します

削除されました
完全な移行ガイドと詳細な変更ログ: CHANGELOG.md
AMD MI300X 検証、32B パラメータ ベンチマーク、および ROCm wave-64 カーネル開発は、DigitalOcean のオープンソース スポンサーシップ プログラムによって可能になりました。
447 件のテストに合格しました。依存関係ゼロ。 3 つの GPU バックエンド。 1つのパッケージ。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The inference engine the open-source world built for itself. - Release v2.0.0 — The Own Everything Release · Zyora-Dev/zse

Release v2.0.0 — The Own Everything Release · Zyora-Dev/zse · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Zyora-Dev
/
zse
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
v2.0.0 — The Own Everything Release
Filter
Loading
Sorry, something went wrong.
There was an error while loading. Please reload this page .
v2.0.0
55a9742
ZSE v2.0.0 — The "Own Everything" Release
A complete rewrite. Zero third-party dependencies. No PyTorch, no Triton, no transformers, no bitsandbytes. Pure-Python kernel compiler emits CUDA C, HIP C, and Metal Shading Language directly.
Headline numbers (Qwen2.5-14B INT4 vs vLLM AWQ INT4 on A100-80GB)
pip install zse-engine
zse serve model.zse --port 8000
Or run the kernel compiler standalone:
pip install zse-compiler
What's in this release
ZSE Kernel Compiler — @zse.kernel Python DSL → CUDA / HIP / Metal. Warp primitives, vectorized memory, block reductions, tiling, fusion, WMMA, CDNA3 MFMA matrix cores, auto-tuning.
.zse model format v2 — pre-quantized INT4/INT8/FP16, mmap-friendly, C-accelerated quantization (~600× faster). Adapters for Llama / Mistral / Qwen2 / Gemma2 / Phi3.
Own PagedAttention — adaptive block sizing, token-level eviction, FNV-1a dedup, COW forking.
ZStreamer — continuous batching, disaggregated prefill/decode, chunked prefill, speculative decoding (n-gram + self-draft).
Orchestrator — unified VRAM allocator, 29 GPU kernels on MI300X, CUDA Graphs + HIP Graphs, LoRA hot-swap.
Server — OpenAI-compatible API, API key auth, rate limiting, SQLite store, built-in RAG ( /v1/rag/* ), web dashboard.
RAG — BM25 + TF-IDF + dense embeddings (via the loaded LLM, zero extra deps) + Reciprocal Rank Fusion + LLM cross-encoder rerank.
Tensor Parallelism — pure-ctypes NCCL/RCCL wrapper, multi-process workers.
Package rename: zllm-zse → zse-engine on PyPI
Module rename: zse → zse_engine
.zse format v2 is incompatible with 1.x — re-convert with zse convert
bnb / bitsandbytes backend removed
PyTorch / Triton / transformers dependencies removed
Full migration guide and detailed change log: CHANGELOG.md
AMD MI300X validation, 32B-parameter benchmarks, and our ROCm wave-64 kernel development were made possible by DigitalOcean's Open Source Sponsorship Program .
447 tests passing. Zero dependencies. Three GPU backends. One package.
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
