---
source: "https://github.com/Kanchisaw03/axiom"
hn_url: "https://news.ycombinator.com/item?id=48751251"
title: "Show HN: a Rust OS kernel built for LLM inference"
article_title: "GitHub - Kanchisaw03/axiom · GitHub"
author: "Kanchisaw"
captured_at: "2026-07-01T18:36:55Z"
capture_tool: "hn-digest"
hn_id: 48751251
score: 2
comments: 0
posted_at: "2026-07-01T18:30:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: a Rust OS kernel built for LLM inference

- HN: [48751251](https://news.ycombinator.com/item?id=48751251)
- Source: [github.com](https://github.com/Kanchisaw03/axiom)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T18:30:53Z

## Translation

タイトル: HN の表示: LLM 推論用に構築された Rust OS カーネル
記事タイトル: GitHub - Kanchisaw03/axiom · GitHub
説明: GitHub でアカウントを作成して、Kanchisaw03/axiom の開発に貢献します。

記事本文:
GitHub - Kanchisaw03/axiom · GitHub
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
カンチソー03
/
公理
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

s
2 コミット 2 コミット .cargo .cargo ベンチマーク ベンチマーク src src ツール tools .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md run_qemu_sdl_safe.ps1 run_qemu_sdl_safe.ps1 typescript typescript x86_64-axiom.json x86_64-axiom.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Linux はトランスフォーマー推論を中間層でプリエンプトし、テンソル状態を汎用の 4 KB ページに割り当てます。これにより、キャッシュの中断と、制約のあるマシンでの推論と一致しないメモリ動作が発生します。
AXIOM は、別のアプローチを評価するために構築されたブート可能な Rust no_std カーネルです。つまり、推論クリティカルな OS プリミティブをファーストクラスにし、ホット パスから汎用抽象化を削除します。
AXIOM は汎用オペレーティング システムではありません。推論基板です。
AXIOM は、推論固有の OS プリミティブをブート可能なベアメタル カーネルに実装できることを示しています。
テンソルネイティブのメモリ割り当て
ダブルバッファーされたウェイトストリーミング
現在の測定では、プリフェッチ パスの修正後、ストリーミング オーバーヘッドは 1 層あたり約 1.4 秒から 1 層あたり約 42 マイクロ秒に減少しました。現在のエンドツーエンドのスループットは、依然としてコンピューティングと VM ベースのストレージ エミュレーションによってボトルネックになっています。意図された評価ターゲットは、ベアメタル NVMe でのメモリ制約のある 7B クラスの推論です。
Transformer 推論には、レイヤー粒度およびテンソル形状粒度で予測可能な実行時構造があります。デフォルトの Linux スタックは、この構造ではなく、広範なマルチプログラムのワークロード向けに最適化されています。
AXIOM の動機となるのは 2 つの障害モードです。
スケジューラの粒度が一致しません。
Linux CFS タイマーのプリエンプションは通常 4 ミリ秒程度です。単一レイヤー内のアテンション セグメントと FFN セグメントは、同様の間隔またはより長い間隔にまたがることができます。中間層のプリエンプションにより、アクティブなワーキング セットが削除される可能性があります。

すべてのレイヤー遷移で繰り返しキャッシュの再ウォームを強制します。
メモリの抽象化の不一致。
Linux バディの割り当てとページレベルの管理は一般的です。 KV キャッシュとレイヤーの重みは一般的なものではありません。これらは、アクセスが予測可能な、大きくて形状が安定したテンソルです。メモリ負荷がかかると (たとえば、4 GB クラス システムで 7B)、スワップ フォールバックがレイテンシを支配し、対話型推論が非実用的になります。
AXIOM は、ユーザー空間ではなくカーネル レベルで両方の問題に対処します。
現在、AXIOM はブートし、メモリ/割り込み/スケジューラー/ストリーミング/ランタイム サブシステムを初期化し、量子化されたトランスフォーマー推論を実行し、レイヤーごとのタイミングとベンチマーク テレメトリを発行して停止します。
推論タスク制御を超えたプロセスモデル
+----------------------- AXIOM ブート シーケンス -------------------------+
|メイン.rs |
|シリアル::初期化() |
| -> メモリ::初期化() |
| -> 割り込み::init() |
| -> スケジューラー::init() |
| -> ストリーミング::init() |
| -> TransformerRuntime::new() |
+---------------------------------------------------------------------+
トークンごとの実行
+---------------------------------------------------------------------+
| forward_one_token() |
| 1) streamer.begin_layer(L) |
| - 計算ではバッファ A | を使用します。
| - プリフェッチは L+1 をバッファ B にロードします。
| 2）scheduler.begin_layer() -> LayerLock を取得 |
| - タイマーのプリエンプションは層の境界まで延期される |
| 3) 計算パス |
| - rmsnorm -> 注意 -> 残差 |
| - rmsnorm -> ffn (proj/act/down) -> 残差 |
| 4) LayerLock リリース -> yield_at_layer_boundary() |
| 5) streamer.end_layer(L) |
+---------------------------------------------------------------------+
物理的に連続したテンソル領域 (事前予約)
+-------------------------------------+----------------------+----------------------+
| KVCachePool (~45%) |ウェイトプール (~45%) |活性化 (~10%) |
|固定スラブ |スロットステートマシン

|トリプルリングバッファ |
|ホットパスの割り当てがありません | U->S->R->C->R/U |インデックス = pos % 3 |
+-------------------------------------+----------------------+----------------------+
凡例: U=アンロード、S=ストリーミング、R=常駐、C=コンピューティング
仮想メモリ マップ (現在の QEMU VM 構成、縮尺は正確ではありません)
+----------------------------+--------------------------------------------+
|アドレス範囲 |地域 |
+----------------------------+--------------------------------------------+
| 0x444444440000-0x444444540000|カーネル ヒープ (1 MiB) |
| 0x20000000 |テンソル領域の終わり |
| 0x1e666000-0x20000000 | ActivationPool (~10%、トリプルリング) |
| 0x17333000-0x1e666000 | WeightPool (~45%、レイヤー スロット) |
| 0x10000000-0x17333000 | KVCachePool (~45%、固定スラブ) |
| 0x10000000 |テンソル領域ベース (256 MiB ウ​​ィンドウ) |
+----------------------------+--------------------------------------------+
1) Tensor ネイティブ アロケーター (汎用 p を置き換える)
[切り捨てられた]
4 KB ページにわたる汎用バディ アロケータ
長期間存続する推論テンソルの断片化とページレベルの間接化
事前に予約された物理的に連続したテンソル領域
3 つの専用プール:
KVCachePool: 固定スラブ、ホットパス再配置なし、アロケータ チャーンなし
WeightPool: 明示的なステート マシンを備えたレイヤーごとのスロット (アンロード、ストリーミング、常駐、コンピューティング)
ActivationPool: トリプル リング バッファ、フォワード ホット パスでのゼロ割り当て
これが推論に適している理由:
テンソルの寿命と形状は既知です
アロケータの作業はトークン パスの外に移動されます
メモリ レイアウトは、プロセス中心の抽象化ではなく、アクセス パターンに一致します。
2) LayerLock Scheduler (レイヤーの計算中にタイムスライスのプリエンプションを置き換えます)
変圧器層の境界に依存しない定期的なプリエンプション
1 つのレイヤーの実行中のプリエンプションの明示的な抑制
YieldAtLayerBoundary() でのスケジュール決定
これが推論に適している理由

理由:
レイヤローカルのキャッシュ常駐性を保持します
スケジューラによる中間層の中断を除去します。
モデル層の粒度でタイミング動作を推論しやすくします。
3) ダブルバッファーウェイトストリーミング (ブロッキングレイヤーフェッチを置き換えます)
Linux/アプリケーションの動作が置き換えられました:
カーネル I/O スケジューリングのオーバーラップに対する制御が制限された、ユーザースペース主導の重みロード
次の層の重み付けを待機しているため、頻繁に計算が停止します。
レイヤ N はバッファ A から計算し、レイヤ N+1 はバッファ B にロードされます
ランタイムとストリーマー間の明示的な begin_layer/end_layer 統合
これが推論に適している理由:
I/O と直接オーバーラップし、レイヤー境界で計算します
メモリに完全に収まらないモデル用に制御されたパイプラインを作成します
公理/
|- Cargo.toml
|- x86_64-axiom.json
|- .cargo/config.toml
|- ツール/
| |-pack_weights.py
| |- run_qemu.ps1
| |- run_qemu_wsl_safe.sh
|- ベンチマーク/
| |- 比較.py
| |- 結果.csv
|- ソース/
|- main.rs
|- シリアル.rs
|- 中断.rs
|- 記憶/
| |-frame.rs
| |- ページング.rs
| |- heap.rs
| |- tensor.rs
| |- mod.rs
|- スケジューラー/
| |-layer_lock.rs
| |- mod.rs
|- ストリーミング/
| |- virtio.rs
| |- プリフェッチ.rs
| |- mod.rs
|- ランタイム/
|- matmul.rs
|-attention.rs
|-rope.rs
|- rmsnorm.rs
|- ffn.rs
|-sampling.rs
|- tokenizer.rs
|- mod.rs
ビルド環境
ターゲット: x86_64-unknown-none (no_std)
ターゲット機能: +avx2、+fma、+sse2 (推論前にカーネルで SSE/AVX 状態が有効になっている)
テスト済みモデル: SmolLM2-135M Q4、TinyLlama-1.1B Q4
rusup ツールチェーンを毎晩インストールする
rusup コンポーネントは、rust-src --toolchain を夜間に追加します。
カーゴインストールブートイメージ
python tools/pack_weights.py --model TinyLlama/TinyLlama-1.1B-Chat-v1.0 --output axiom_weights.img --quant q4
カーゴ + 夜間実行 -- リリース
ベンチマークの進行状況 (単一の最適化セッション)
この進行のための環境:
反復カーネル/

1 つのセッションでのランタイムの最適化
この表は中心的な経験談です。最終的な数値ではなく、アーキテクチャおよびカーネル パスの変更による累積的な影響を示しています。
このシーケンス全体の全体的な改善は約 14.8 倍 TPS (0.023 から ~0.340) であり、I/O ストールの除去と FFN 投影コストの削減によって大きな利点が得られます。
プリフェッチ パス修正前:
stream_us はレイヤーあたり約 1200000 ～ 1700000 us でした
ストリームはレイヤー時間の約 99.5% を占めました
プリフェッチパス修正後:
stream_us は、トレース ログのレイヤーあたり約 42 us です。
ストリーミング オーバーヘッドはもはや主要な用語ではありません
ダブルバッファ設計はアーキテクチャの方向性として検証されています
ストリーミングのオーバーヘッドが崩壊すると、コンピューティング カーネルが制限要因になりました
現在のボトルネックの内訳 (0.340 TPS 状態)
代表的なレイヤーごとのタイミング (現在の状態):
主なボトルネックは、ストリーミングではなく、FFN プロジェクション/ダウンプロジェクションのコンピューティングです。
これらは現在の事実であり、隠すべき注意事項ではありません。
ベンチマークは、ベアメタル NVMe ではなく、エミュレートされた virtio-blk ストレージを備えた KVM で行われます。
SmolLM2-135M は RAM に収まります。このサイズでは、ストリーミングのオーバーヘッドは目標の利点ではありません。
VM ノイズは重要です。 2 つの同様の実行では、0.278 対 0.323 TPS が示されました。単一サンプルではなく、繰り返し実行された中央値をレポートします。
この環境では、AVX2 固有パスが試行されましたが、スカラー LUT と比較して回帰しました。スカラー Q4 LUT パスはデフォルトのままです。
AXIOM は完全なメモリ内ベースラインの前提が当てはまらない、メモリに制約のあるワークロードを対象としているため、ここでは llama.cpp の比較は示しません。
現在のランタイムには、カーネル コード内に単純化された最終ロジット射影パスが含まれています。これは立ち上げには機能しますが、最終品質の出力ヘッドの実装には機能しません。
評価対象はmサイズにすっぽり収まる小型モデルではない

思い出。対象は次のとおりです。
制約のある RAM フットプリント上の 7B クラス モデル (例: 4 GB クラス システム)
完全なモデル常駐が不可能な状況
ストリーミングとコンピューティングの重複が実現可能性を決定する場合
ここでは、スケジューリング、メモリ レイアウト、および I/O パイプラインに対する OS レベルの制御が最も重要になると予想されます。この記述は、ターゲット ハードウェアで測定されるまでは予測される仮説です。
現状と今後の取り組み
エンドツーエンドの推論パスを備えたブート可能な Rust no_std カーネル
ランタイムと統合されたテンソルネイティブアロケーター
タイマー パスと統合されたレイヤー境界スケジューリング プリミティブ
レイヤ実行と統合されたダブルバッファストリーミング
レイヤーあたりのストリーム オーバーヘッドの秒スケールからマイクロ秒スケールへの削減を測定
現在の VM セットアップでは、反復可能な高速ベンチマークは約 0.34 TPS で実行されます。
メモリに制約のある 7B ワークロードでのベアメタル NVMe 評価
より強力な統計プロトコル (中央値、分散、構成ごとの反復試行)
FFN プロジェクション/ダウンプロジェクションに重点を置いたコンピューティング カーネルの最適化
アーキテクチャを意識したディスパッチ/チューニングによる AVX2 パス回帰の解決
簡略化された出力ヘッド パスを完全な運用ロジット パスに置き換えます。
virtio エミュレーションを削除します

[切り捨てられた]

## Original Extract

Contribute to Kanchisaw03/axiom development by creating an account on GitHub.

GitHub - Kanchisaw03/axiom · GitHub
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
Kanchisaw03
/
axiom
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .cargo .cargo benchmarks benchmarks src src tools tools .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md run_qemu_sdl_safe.ps1 run_qemu_sdl_safe.ps1 typescript typescript x86_64-axiom.json x86_64-axiom.json View all files Repository files navigation
Linux preempts transformer inference mid-layer and allocates tensor state with general-purpose 4 KB pages, which causes cache disruption and memory behavior that is mismatched to inference on constrained machines.
AXIOM is a bootable Rust no_std kernel built to evaluate a different approach: make inference-critical OS primitives first-class and remove general-purpose abstractions from the hot path.
AXIOM is not a general-purpose operating system. It is an inference substrate.
AXIOM demonstrates that inference-specific OS primitives can be implemented in a bootable bare-metal kernel:
tensor-native memory allocation
double-buffered weight streaming
In current measurements, streaming overhead dropped from approximately 1.4 seconds per layer to approximately 42 microseconds per layer after prefetch-path corrections. Current end-to-end throughput is still bottlenecked by compute and by VM-based storage emulation. The intended evaluation target is memory-constrained 7B-class inference on bare metal NVMe.
Transformer inference has runtime structure that is predictable at layer granularity and tensor-shape granularity. The default Linux stack is optimized for broad multiprogrammed workloads, not this structure.
Two failure modes motivate AXIOM:
Scheduler granularity mismatch.
Linux CFS timer preemption is typically on the order of 4 ms. Attention and FFN segments in a single layer can span similar or longer intervals. Mid-layer preemption can evict the active working set and force repeated cache re-warm on every layer transition.
Memory abstraction mismatch.
Linux buddy allocation and page-level management are generic. KV cache and layer weights are not generic: they are large, shape-stable tensors with predictable access. Under memory pressure (for example, 7B on 4 GB class systems), swap fallback dominates latency and makes interactive inference impractical.
AXIOM addresses both issues at kernel level, not in userspace.
AXIOM currently boots, initializes memory/interrupt/scheduler/streaming/runtime subsystems, runs quantized transformer inference, emits per-layer timing and benchmark telemetry, and halts.
process model beyond inference task control
+------------------------- AXIOM Boot Sequence -------------------------+
| main.rs |
| serial::init() |
| -> memory::init() |
| -> interrupts::init() |
| -> scheduler::init() |
| -> streaming::init() |
| -> TransformerRuntime::new() |
+---------------------------------------------------------------------+
Per-token Execution
+---------------------------------------------------------------------+
| forward_one_token() |
| 1) streamer.begin_layer(L) |
| - compute uses buffer A |
| - prefetch loads L+1 into buffer B |
| 2) scheduler.begin_layer() -> LayerLock acquire |
| - timer preemption deferred until layer boundary |
| 3) compute path |
| - rmsnorm -> attention -> residual |
| - rmsnorm -> ffn (proj/act/down) -> residual |
| 4) LayerLock release -> yield_at_layer_boundary() |
| 5) streamer.end_layer(L) |
+---------------------------------------------------------------------+
Physically Contiguous Tensor Region (Pre-reserved)
+------------------------+-----------------------+--------------------+
| KVCachePool (~45%) | WeightPool (~45%) | Activation (~10%) |
| fixed slabs | slot state machine | triple ring buffer |
| no hot-path alloc | U->S->R->C->R/U | index = pos % 3 |
+------------------------+-----------------------+--------------------+
Legend: U=Unloaded, S=Streaming, R=Resident, C=Computing
Virtual Memory Map (current QEMU VM configuration, not to scale)
+------------------------------+--------------------------------------+
| Address Range | Region |
+------------------------------+--------------------------------------+
| 0x444444440000-0x444444540000| Kernel heap (1 MiB) |
| 0x20000000 | Tensor region end |
| 0x1e666000-0x20000000 | ActivationPool (~10%, triple ring) |
| 0x17333000-0x1e666000 | WeightPool (~45%, layer slots) |
| 0x10000000-0x17333000 | KVCachePool (~45%, fixed slabs) |
| 0x10000000 | Tensor region base (256 MiB window) |
+------------------------------+--------------------------------------+
1) Tensor-Native Allocator (replaces generic p
[truncated]
generic buddy allocator over 4 KB pages
fragmentation and page-level indirection for long-lived inference tensors
pre-reserved physically contiguous tensor region
three dedicated pools:
KVCachePool: fixed slabs, no hot-path relocation, no allocator churn
WeightPool: per-layer slots with explicit state machine (Unloaded, Streaming, Resident, Computing)
ActivationPool: triple ring buffer, zero allocation in forward hot path
Why this is better for inference:
tensor lifetimes and shapes are known
allocator work is moved out of the token path
memory layout matches access pattern instead of process-centric abstractions
2) LayerLock Scheduler (replaces time-slice preemption during layer compute)
periodic preemption independent of transformer layer boundaries
explicit suppression of preemption while one layer executes
scheduling decisions at YieldAtLayerBoundary()
Why this is better for inference:
preserves layer-local cache residency
removes scheduler-induced mid-layer interruption
makes timing behavior easier to reason about at model-layer granularity
3) Double-Buffered Weight Streaming (replaces blocking layer fetch)
Linux/application behavior replaced:
userspace-driven weight loading with limited control over kernel I/O scheduling overlap
frequent compute stalls waiting on next-layer weights
layer N computes from buffer A while layer N+1 is loaded into buffer B
explicit begin_layer/end_layer integration between runtime and streamer
Why this is better for inference:
directly overlaps I/O and compute at the layer boundary
creates a controlled pipeline for models that do not fully fit in memory
axiom/
|- Cargo.toml
|- x86_64-axiom.json
|- .cargo/config.toml
|- tools/
| |- pack_weights.py
| |- run_qemu.ps1
| |- run_qemu_wsl_safe.sh
|- benchmarks/
| |- compare.py
| |- results.csv
|- src/
|- main.rs
|- serial.rs
|- interrupts.rs
|- memory/
| |- frame.rs
| |- paging.rs
| |- heap.rs
| |- tensor.rs
| |- mod.rs
|- scheduler/
| |- layer_lock.rs
| |- mod.rs
|- streaming/
| |- virtio.rs
| |- prefetch.rs
| |- mod.rs
|- runtime/
|- matmul.rs
|- attention.rs
|- rope.rs
|- rmsnorm.rs
|- ffn.rs
|- sampling.rs
|- tokenizer.rs
|- mod.rs
Build Environment
target: x86_64-unknown-none (no_std)
target features: +avx2, +fma, +sse2 (with SSE/AVX state enabled in kernel before inference)
tested models: SmolLM2-135M Q4, TinyLlama-1.1B Q4
rustup toolchain install nightly
rustup component add rust-src --toolchain nightly
cargo install bootimage
python tools/pack_weights.py --model TinyLlama/TinyLlama-1.1B-Chat-v1.0 --output axiom_weights.img --quant q4
cargo +nightly run --release
Benchmark Progression (Single Optimization Session)
Environment for this progression:
iterative kernel/runtime optimizations in one session
This table is the central empirical story: not one final number, but cumulative impact of architectural and kernel-path changes.
Overall improvement across this sequence is approximately 14.8x TPS (0.023 to ~0.340), with major gains coming from removing I/O stalls and reducing FFN projection cost.
Before the prefetch-path correction:
stream_us was approximately 1200000 to 1700000 us per layer
stream accounted for approximately 99.5% of layer time
After the prefetch-path correction:
stream_us is approximately 42 us per layer in trace logs
streaming overhead is no longer the dominant term
the double-buffer design is validated as an architectural direction
once streaming overhead collapsed, compute kernels became the limiting factor
Current Bottleneck Breakdown (0.340 TPS State)
Representative per-layer timing (current state):
Dominant bottleneck is FFN projection/down-projection compute, not streaming.
These are current facts, not caveats to hide.
Benchmarks are on KVM with emulated virtio-blk storage, not bare metal NVMe.
SmolLM2-135M fits in RAM; at this size, streaming overhead is not the target benefit.
VM noise is material; two similar runs have shown 0.278 vs 0.323 TPS. Report medians over repeated runs, not single samples.
AVX2 intrinsic path was attempted and regressed relative to scalar LUT in this environment; scalar Q4 LUT path remains default.
No llama.cpp comparison is presented here because AXIOM targets memory-constrained workloads where full in-memory baseline assumptions do not hold.
Current runtime still contains a simplified final logit projection path in kernel code; this is functional for bring-up but not a final-quality output head implementation.
The evaluation target is not small models that fit comfortably in memory. The target is:
7B-class models on constrained RAM footprints (for example 4 GB class systems)
conditions where full model residency is impossible
where overlap of streaming and compute determines feasibility
This is where OS-level control over scheduling, memory layout, and I/O pipeline is expected to matter most. This statement is a projected hypothesis until measured on the target hardware.
Current Status and Future Work
bootable Rust no_std kernel with end-to-end inference path
tensor-native allocator integrated with runtime
layer-boundary scheduling primitive integrated with timer path
double-buffered streaming integrated with layer execution
measured reduction of stream overhead from second-scale to microsecond-scale per layer
repeatable fast-benchmark runs around ~0.34 TPS in current VM setup
bare-metal NVMe evaluation on memory-constrained 7B workloads
stronger statistical protocol (median, dispersion, repeated trials per configuration)
compute kernel optimization focused on FFN projection/down-projection
resolution of AVX2 path regressions with architecture-aware dispatch/tuning
replacement of simplified output-head path with full production logits path
removes virtio emulation ove

[truncated]
