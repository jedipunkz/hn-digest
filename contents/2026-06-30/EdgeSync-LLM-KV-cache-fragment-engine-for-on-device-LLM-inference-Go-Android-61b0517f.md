---
source: "https://github.com/bossandboss/EdgeSync-LLM"
hn_url: "https://news.ycombinator.com/item?id=48732973"
title: "EdgeSync-LLM – KV cache fragment engine for on-device LLM inference (Go/Android)"
article_title: "GitHub - bossandboss/EdgeSync-LLM: Differential Local Semantic Cache for Embedded AI · GitHub"
author: "bossandboss"
captured_at: "2026-06-30T14:51:41Z"
capture_tool: "hn-digest"
hn_id: 48732973
score: 2
comments: 0
posted_at: "2026-06-30T14:10:22Z"
tags:
  - hacker-news
  - translated
---

# EdgeSync-LLM – KV cache fragment engine for on-device LLM inference (Go/Android)

- HN: [48732973](https://news.ycombinator.com/item?id=48732973)
- Source: [github.com](https://github.com/bossandboss/EdgeSync-LLM)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T14:10:22Z

## Translation

タイトル: EdgeSync-LLM – オンデバイス LLM 推論用の KV キャッシュ フラグメント エンジン (Go/Android)
記事タイトル: GitHub - Bossandboss/EdgeSync-LLM: 組み込み AI 用の差分ローカル セマンティック キャッシュ · GitHub
説明: 組み込み AI 用の差分ローカル セマンティック キャッシュ。 GitHub でアカウントを作成して、bossandboss/EdgeSync-LLM の開発に貢献してください。

記事本文:
GitHub - Bossandboss/EdgeSync-LLM: 組み込み AI 用の差分ローカル セマンティック キャッシュ · GitHub
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
上司と上司
/
EdgeSync-LLM
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

ches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
13 コミット 13 コミット アダプター アダプター ベンチマーク ベンチマーク キャッシュ キャッシュ コア コア ドキュメント ドキュメント 埋め込み 埋め込み モニター モニター プリフェッチ プリフェッチ sdk/ android sdk/ android セキュリティ security src src LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sumindex.htmlindex.htmlmetadata.jsonmetadata.jsonpackage.json package.json server.ts server.ts tsconfig.json tsconfig.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
EdgeSync-LLM — ローカル LLM 用の KV フラグメント エンジン
オンデバイス LLM 推論のための、エンジンに依存しない KV キャッシュ フラグメント システム。
ARM64 Android (Cortex-A55/A78) 用に設計されており、実行中のあらゆるプラットフォームに移植可能
llama.cpp、MLC-LLM、または ONNX ランタイム。
アプリケーションとLLMの間に位置する再利用可能なKVキャッシュ層
エンジン。リクエストごとに完全な事前入力を再実行する代わりに、
アテンション KV テンソル (キーと値) のスライスは、次の方法で取得します。
近似最近傍検索 (HNSW) を実行し、それらを直接に注入します。
エンジンの KV キャッシュ - 推論の最も高価な部分をスキップします。
これは、応答文字列を保存する「セマンティック キャッシュ」ではありません。それは、
実際のトランスフォーマー KV テンソル、トークン範囲とレイヤー範囲によって識別されます。
そしてリクエスト時にそれらを再構築します。
[プロンプト]
│
▼
[ 組み込みモデル ] MiniLM-L6-v2 (384 次元、~8ms CPU)
│
▼
[ HNSW インデックス ] Pure Go、M=16、efSearch=50
│
┌───────┴─────────────┐
│ │
sim ≧ 0.92 0.75 ≤ sim < 0.92 sim < 0.75
│ │ │
┌──────▼──────┐ ┌──────▼──────┐ ┌────────▼────┐
│ 正確なヒット │ │ P

アーティアルヒット │ │ ミス │
│ │ │ │ │ │
│ KV の注入 │ │ プレフィックスの注入 │ │ フルプレフィル │
│ フラグメント │ │ デルタを生成 │ │ フラグメントを抽出します。 │
│ ~8ms TTFT │ │ ~280ms TTFT │ │ HNSW にストア │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│ │ │
━━━━━━━━━━━━━━━━━━━━━┴───────────┘
│
[ KVA アダプター層 ]
│
┌───────────┼───────────┐
▼ ▼ ▼
[ llamacpp ] [ mlc-llm ] [ onnx ランタイム ]
(GGML テンソル API) (TVM ページ化された KV) (past_key_values)
ファイル構造
§── キャッシュ/
│ §── flagment.go ← KVFragment: キャッシュユニットの正式な定義
│ │ (ディメンション、TTL、エビクションポリシー、ストレージキー)
│ §── Differential.go ← DifferentialEngine: EXACT / PARTIAL / MISS ルータ
│ └─ schema.go ← フラグメントメタデータの SQLite WAL スキーマ
│
§── アダプター/
│ §──interface.go ← KVAdapter: エンジンに依存しないコントラクト
│ │ (ExtractFragment / InjectFragment / Generate)
│ §── llamacpp.go ← llama.cpp アダプター (GGML tensor API、CGO)
│ §── mlc.go ← MLC-LLM アダプタ (TVM ページング KV、mlc4j)
│ └─ onnx.go ← ONNX ランタイム アダプター (past_key_values)
│
§── コア/
│ §── hnsw.go ← Pure Go HNSW インデックス (M=16, efSearch=50)
│ └─ cosine_neon.c ← ARM NEON fp16 コサイン類似度
│
§── sdk/android/
│ └── EdgeSyncLLM.kt ← Kotlin JNI ブリッジ (コルーチンの一時停止)
│
§── モニター/
│ ━─ エネルギー

_android.go ← Android /sys/class/power_supply/ プロファイラ
│
§── プリフェッチ/
│ lux──predictor.go ← N-gram プリフェッチ予測子 (上位 3 候補)
│
└── ベンチマーク/
└──runner.go ← 改ざん可能ベンチマーク：3モード×1000リクエスト
KVフラグメント
キャッシュのアトミック単位。正式には、cache/fragment.go で定義されています。
構築時に適用される不変条件:
len(キー) > 0 && len(値) > 0
adapter/interface.go で定義されています。どのエンジンでも 6 つのメソッドが実装されています。
ExtractFragment ( ctx 、 tokenIDs 、layerStart 、layerEnd 、layerStride 、embedding )
→ * KVFragment 、エラー
InjectFragment ( ctx 、 フラグメント )
→エラー
生成 ( ctx 、 プロンプト 、 startTokenPos 、 maxTokens )
→ テキスト文字列 、 tokensGenerated int 、 エラー
トークン化 ( ctx 、 text )
→ [] int32 、エラー
KVCache のクリア (ctx)
→エラー
閉じる ()
→エラー
エンジン間の再利用: エンジン B は、次の場合に限り、エンジン A からフラグメントを挿入できます。
B は、 CompareWith() に A をリストします。現在の互換性マトリックス:
エンジン間の再利用 (例: llamacpp → onnx) には、KV tensor reshape アダプターが必要です
(転置 [seq, heads, dim] → [heads, seq, dim] )。まだ実装されていません。
Benchmark/runner.go のベンチマークは、1000 リクエストにわたる 3 つのモードを比較します。
8 つのセマンティック プロンプト クラスター (64 個の一意のプロンプト + それぞれ 4 個のバリアント) から抽出されます。
タイミング モデル (アドホックなランダム範囲ではなく、Cortex-A55 測定から派生):
./benchmark/ を実行してください
# クエリごとの詳細な出力:
EDGESYNC_VERBOSE=1 ./benchmark/ を実行します。
建物
# ホストのビルド (ベンチマークのみ、CGO なし):
./benchmark/ を実行してください
# Android ARM64 (llama.cpp CGO を使用):
import CGO_CFLAGS= " -I/path/to/llama.cpp "
import CGO_LDFLAGS= " -L/path/to/llama.cpp/build -llama -lm "
CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 \
go build -oedgecache ./...
# NEON コサイン モジュール (ARM64 のみ):
aarch64-linux-gn

u-gcc -O3 -march=armv8.2-a+fp16 \
-c コア/cosine_neon.c -o コア/cosine_neon.o
実装される内容
クロスエンジン KV テンソルの再形成 (adapter/reshape.go ) — llamacpp と ONNX ランタイムの間で [seq,heads,dim] ↔ [heads,seq,dim] を転置します。 CanInjectWithReshape() は検出とフォールバックを自動的に処理します
フラグメント圧縮 (cache/compactor.go) — ContentHash による重複排除、レイヤー構成によるグループ化、エンジンごとのテンソル連結による隣接マージ (llamacpp の場合は軸 0、ONNX の場合はヘッドごとの軸 1)。マージされたエンベディングは加重正規化平均です
永続的フラグメント ストア (cache/store.go) — 2 層ストレージ: sync.Map ホット キャッシュ + SQLite WAL (メタデータ用)。 SQLite ページの断片化を避けるために、テンソル BLOB は <id>.keys.bin / <id>.vals.bin として書き込まれます。プレフィックス範囲検索用の QueryByTokenRange()
実際の埋め込みモデル (embedding/minilm.go) — ORTEncoder は、ONNX ランタイム経由で all-MiniLM-L6-v2 (22 MB、384 次元、Cortex-A55 で ~8ms) を実行します。 .ort モデル ファイルが見つからない場合、FallbackEncoder (FNV-1a ハッシュ) が自動的にアクティブ化されます。
Android JNI ブリッジ ( sdk/android/EdgeSyncLLM.kt + sdk/android/jni_bridge.go ) — アダプター/パッケージ API を公開する完全な書き換え:ネイティブInitialize、nativeEmbed、nativeLookup、nativeInjectFragment、nativeGenerateFromPos、nativeExtractAndStore、nativeCompact、nativeReshapeFragment
このプロジェクトは、ビジネス ソース ライセンス 1.1 (BUSL-1.1) に基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
✅ 研究、評価、開発、内部テストは無料
✅ ソースコードは読みやすく、フォーク可能です
❌ エンドユーザーに展開される商用アプリ、SaaS プラットフォーム、またはモバイル アプリでの本番使用には商用ライセンスが必要です
🔄 2029 年 7 月 1 日、プロジェクトは自動的に AGPL-3.0 になります。
商用ライセンス: github.com/bossandboss/EdgeSy で問題をオープンしてください

ラベルCommercial-licenseを持つnc-LLM。
組み込み AI 用の差分ローカル セマンティック キャッシュ
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Differential Local Semantic Cache for Embedded AI. Contribute to bossandboss/EdgeSync-LLM development by creating an account on GitHub.

GitHub - bossandboss/EdgeSync-LLM: Differential Local Semantic Cache for Embedded AI · GitHub
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
bossandboss
/
EdgeSync-LLM
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
13 Commits 13 Commits adapter adapter benchmark benchmark cache cache core core docs docs embedding embedding monitor monitor prefetch prefetch sdk/ android sdk/ android security security src src LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum index.html index.html metadata.json metadata.json package.json package.json server.ts server.ts tsconfig.json tsconfig.json vite.config.ts vite.config.ts View all files Repository files navigation
EdgeSync-LLM — KV Fragment Engine for Local LLMs
A engine-agnostic KV cache fragment system for on-device LLM inference.
Designed for ARM64 Android (Cortex-A55/A78), portable to any platform running
llama.cpp, MLC-LLM, or ONNX Runtime.
A reusable KV cache layer that sits between the application and the LLM
engine. Instead of re-running the full prefill on every request, it stores
slices of the attention KV tensors (Keys and Values), retrieves them via
approximate nearest-neighbor search (HNSW), and injects them directly into
the engine's KV cache — skipping the most expensive part of inference.
This is not a "semantic cache" that stores response strings. It stores the
actual transformer KV tensors , identified by token range and layer range,
and reconstructs them at request time.
[ PROMPT ]
│
▼
[ Embedding Model ] MiniLM-L6-v2 (384-dim, ~8ms CPU)
│
▼
[ HNSW Index ] Pure Go, M=16, efSearch=50
│
┌───────┴───────────────────────────┐
│ │
sim ≥ 0.92 0.75 ≤ sim < 0.92 sim < 0.75
│ │ │
┌──────▼──────┐ ┌─────────▼──────┐ ┌───────▼────────┐
│ EXACT HIT │ │ PARTIAL HIT │ │ MISS │
│ │ │ │ │ │
│ Inject KV │ │ Inject prefix │ │ Full prefill │
│ fragment │ │ Generate delta │ │ Extract frag. │
│ ~8ms TTFT │ │ ~280ms TTFT │ │ Store in HNSW │
└─────────────┘ └────────────────┘ └────────────────┘
│ │ │
└───────────────────────────────────┴──────────────────────┘
│
[ KVAdapter Layer ]
│
┌────────────────────┼─────────────────────┐
▼ ▼ ▼
[ llamacpp ] [ mlc-llm ] [ onnx runtime ]
(GGML tensor API) (TVM paged KV) (past_key_values)
File Structure
├── cache/
│ ├── fragment.go ← KVFragment: formal definition of a cache unit
│ │ (dimensions, TTL, eviction policy, storage key)
│ ├── differential.go ← DifferentialEngine: EXACT / PARTIAL / MISS router
│ └── schema.go ← SQLite WAL schema for fragment metadata
│
├── adapter/
│ ├── interface.go ← KVAdapter: engine-agnostic contract
│ │ (ExtractFragment / InjectFragment / Generate)
│ ├── llamacpp.go ← llama.cpp adapter (GGML tensor API, CGO)
│ ├── mlc.go ← MLC-LLM adapter (TVM paged KV, mlc4j)
│ └── onnx.go ← ONNX Runtime adapter (past_key_values)
│
├── core/
│ ├── hnsw.go ← Pure Go HNSW index (M=16, efSearch=50)
│ └── cosine_neon.c ← ARM NEON fp16 cosine similarity
│
├── sdk/android/
│ └── EdgeSyncLLM.kt ← Kotlin JNI bridge (suspend coroutines)
│
├── monitor/
│ └── energy_android.go ← Android /sys/class/power_supply/ profiler
│
├── prefetch/
│ └── predictor.go ← N-gram prefetch predictor (top-3 candidates)
│
└── benchmark/
└── runner.go ← Falsifiable benchmark: 3 modes × 1000 requests
The KVFragment
The atomic unit of the cache. Formally defined in cache/fragment.go .
Invariants enforced at construction:
len(Keys) > 0 && len(Values) > 0
Defined in adapter/interface.go . Any engine implements 6 methods:
ExtractFragment ( ctx , tokenIDs , layerStart , layerEnd , layerStride , embedding )
→ * KVFragment , error
InjectFragment ( ctx , fragment )
→ error
Generate ( ctx , prompt , startTokenPos , maxTokens )
→ text string , tokensGenerated int , error
Tokenize ( ctx , text )
→ [] int32 , error
ClearKVCache ( ctx )
→ error
Close ()
→ error
Cross-engine reuse: engine B can inject a fragment from engine A if and only if
B lists A in CompatibleWith() . Current compatibility matrix:
Cross-engine reuse (e.g. llamacpp → onnx) requires a KV tensor reshape adapter
(transpose [seq, heads, dim] → [heads, seq, dim] ). Not implemented yet.
The benchmark in benchmark/runner.go compares 3 modes over 1000 requests
drawn from 8 semantic prompt clusters (64 unique prompts + 4 variants each).
Timing model (not ad-hoc random ranges — derived from Cortex-A55 measurements):
go run ./benchmark/
# Verbose per-query output:
EDGESYNC_VERBOSE=1 go run ./benchmark/
Building
# Host build (benchmark only, no CGO):
go run ./benchmark/
# Android ARM64 (with llama.cpp CGO):
export CGO_CFLAGS= " -I/path/to/llama.cpp "
export CGO_LDFLAGS= " -L/path/to/llama.cpp/build -lllama -lm "
CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 \
go build -o edgecache ./...
# NEON cosine module (ARM64 only):
aarch64-linux-gnu-gcc -O3 -march=armv8.2-a+fp16 \
-c core/cosine_neon.c -o core/cosine_neon.o
What is implemented
Cross-engine KV tensor reshape ( adapter/reshape.go ) — transpose [seq,heads,dim] ↔ [heads,seq,dim] between llamacpp and ONNX Runtime; CanInjectWithReshape() handles detection and fallback automatically
Fragment compaction ( cache/compactor.go ) — deduplication by ContentHash , grouping by layer config, adjacency merge with per-engine tensor concatenation (axis 0 for llamacpp, axis 1 per-head for ONNX); merged embedding is a weighted normalized average
Persistent fragment store ( cache/store.go ) — two-tier storage: sync.Map hot cache + SQLite WAL for metadata; tensor blobs written as <id>.keys.bin / <id>.vals.bin to avoid SQLite page fragmentation; QueryByTokenRange() for prefix-range lookups
Real embedding model ( embedding/minilm.go ) — ORTEncoder runs all-MiniLM-L6-v2 (22 MB, 384-dim, ~8ms on Cortex-A55) via ONNX Runtime; FallbackEncoder (FNV-1a hash) activates automatically if the .ort model file is not found
Android JNI bridge ( sdk/android/EdgeSyncLLM.kt + sdk/android/jni_bridge.go ) — full rewrite exposing the adapter/ package API: nativeInitialize , nativeEmbed , nativeLookup , nativeInjectFragment , nativeGenerateFromPos , nativeExtractAndStore , nativeCompact , nativeReshapeFragment
This project is licensed under the Business Source License 1.1 (BUSL-1.1) — see the LICENSE file for details.
✅ Free for research, evaluation, development, and internal testing
✅ Source code is readable and forkable
❌ Production use in commercial apps, SaaS platforms, or mobile apps deployed to end-users requires a commercial license
🔄 On July 1, 2029, the project automatically becomes AGPL-3.0
Commercial licensing: open an issue at github.com/bossandboss/EdgeSync-LLM with the label commercial-license .
Differential Local Semantic Cache for Embedded AI
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
