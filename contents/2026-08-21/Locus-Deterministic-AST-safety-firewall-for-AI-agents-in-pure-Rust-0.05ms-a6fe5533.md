---
source: "https://github.com/ahmadshady747-create/LOCUS"
hn_url: "https://news.ycombinator.com/item?id=49387042"
title: "Locus: Deterministic AST safety firewall for AI agents in pure Rust (<0.05ms)"
article_title: "GitHub - ahmadshady747-create/LOCUS: Sovereign Ambient HUD & Invariant-Verified Microkernel IDE in Pure Rust · GitHub"
image: "https://opengraph.githubassets.com/1d7ba72352706bc4440956dfe430a20063ef3d1a2dfd4c016eb375ecd7e590ea/ahmadshady747-create/LOCUS"
author: "ahmadshadi2004"
captured_at: "2026-08-21T12:26:00Z"
capture_tool: "hn-digest"
hn_id: 49387042
score: 1
comments: 0
posted_at: "2026-08-21T12:24:23Z"
tags:
  - hacker-news
  - translated
---

# Locus: Deterministic AST safety firewall for AI agents in pure Rust (<0.05ms)

- HN: [49387042](https://news.ycombinator.com/item?id=49387042)
- Source: [github.com](https://github.com/ahmadshady747-create/LOCUS)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T12:24:23Z

## Translation

タイトル: Locus: 純粋な Rust の AI エージェント用の決定論的 AST 安全ファイアウォール (<0.05ms)
記事のタイトル: GitHub - ahmadshady747-create/LOCUS: Pure Rust の Sovereign Ambient HUD および Invariant-Verified Microkernel IDE · GitHub
説明: Pure Rust のソブリン アンビエント HUD およびインバリアント検証済みマイクロカーネル IDE - ahmadshady747-create/LOCUS

記事本文:
GitHub - ahmadshady747-create/LOCUS: Pure Rust のソブリン アンビエント HUD およびインバリアント検証済みマイクロカーネル IDE · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ahmadshady747-作成
/
軌跡
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
37 コミット 37 コミット フォルダーとファイル
.github/ ワークフロー .github/ ワークフロー スクリプト スクリプト src src テスト テスト .gitattributes .gitattributes .gitignore .gitignore Cargo.toml Cargo.toml LI

CENSE ライセンス README.md README.md SPEC.md SPEC.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Pure Rust の確定的 AST セーフティ ガード、多言語セマンティック シンボル グラフ、サージカル バイトスパン パッチング、およびゼロ依存モデル コンテキスト プロトコル (MCP) サーバー。
最新の AI コード生成エージェント (Claude Code、Cursor、Copilot、Devin) と自動開発パイプラインは、次の 2 つのシステム エンジニアリングのボトルネックに直面しています。
確率的構文と同時実行性回帰: AI エージェントは、閉じられていない区切り文字、パニックを引き起こす .unwrap() トラップ、非同期ミューテックス スレッドのデッドロック、壊滅的な多項式正規表現 (ReDoS)、および無制限の配列インデックス付けなどの幻覚を頻繁に見せます。
コンテキスト ウィンドウのインフレーション: ソース コード ファイル全体を LLM コンテキスト ウィンドウに繰り返しフィードすると、高レベルのインターフェイス コントラクトではなく、反復的な関数本体にトークン バジェットの最大 80% が無駄になります。
locus-engine は、100% 安全な Rust で書かれたスタンドアロンの膨張ゼロの高性能システム エンジンとして、両方の課題を解決します。決定論的で交渉不可能な安全性の不変条件をマイクロ秒時間 (9.04 μs) で強制し、最小限のコンテキスト フットプリントでクロスファイル シンボル グラフを抽出し、モデル コンテキスト プロトコル (MCP) を介して最新の AI IDE とネイティブに通信します。
📺 ライブ端末検証デモ
$ 軌跡チェック src/async_task.rs
+--------------------------------------------------------------+
| LOCUS AST GUARD 検証 |
+--------------------------------------------------------------+
ターゲットファイル: src/async_task.rs
検証済みレイテンシ: 0.0194 ミリ秒
ステータス: [FAIL] 不変違反が検出されました
違反の種類: ASYNC_MUTEX_DEADLOCK
違反の詳細: std::sync::Mutex が .await の非同期コンテキストで使用されています — 代わりに tokio::sync::Mutex を使用してください。
+--------------------------------------------------------------+
$ 軌跡グラフ src/
+----------------

--------------------------------------+
|軌跡記号グラフ索引 |
+--------------------------------------------------------------+
インデックス付きルート: src/
インデックス付きファイルの合計: 8
抽出された AST シンボル: 28
AST Skeleton によるトークンの節約: 74.8%
インデックス作成の遅延: 4.82 ミリ秒
+--------------------------------------------------------------+
📊 経験的なベンチマークとパフォーマンス指標
最適化されたリリース プロファイル ( opt-level = 3 、 lto = Thin 、 codegen-units = 1 ) でベンチマーク:
能力
軌跡エンジン
従来のリンター (ESLint、Clippy)
クラウド AI ガードレール
検証の待ち時間
9 μs – 0.05 ms (ナノ秒スケール)
250 – 1,500 ミリ秒 (プロセススポーン)
500 – 2,500 ミリ秒 (ネットワーク往復)
実行アーキテクチャ
インメモリ Pure Rust カーネル
Node.js / Python ランタイム
リモートHTTPクラウドAPI
コンテキストトークンの節約
> 50% - 80% (AST スケルトン)
0% (フルファイル)
0% (フルファイル)
MCPプロトコルのサポート
Stdio 経由の組み込み JSON-RPC 2.0
カスタムラッパーが必要です
独自の API
メモリの安全性
100% 安全な Rust (危険なブロックは 0)
さまざま (C/C++/ノード)
未定義
外部依存関係
ゼロ暗号化/ランタイムの膨張
重いノードモジュール / Python 環境
クラウド接続と API キー
確定的な保証
100% 形式的不変式の拒否
ヒューリスティック警告
確率的LLMの再評価
🏛️ システムアーキテクチャとワークフローパイプライン
フローチャート TD
サブグラフ入力 ["受信コード / AI エージェント パッチ"]
RawCode["Raw コード スニペット / ファイル"]
終わり
サブグラフ AstGuardPipeline ["🛡️ AstGuard: 6 パス決定論的ファイアウォール (<0.05ms)"]
P0["パス 0: デリミタバランス (ダイクストラ)"]
P1["パス 1: 待機中の非同期ミューテックス"]
P2[「パス 2: ゼロ除算ガード」]
P3["パス 3: 配列境界のオーバーフロー"]
P4[「パス 4: 安全でないアンラップ / 期待トラップ」]
P5[「パス 5: ReDoS の壊滅的なバックトラッキング」]
P6[「パス 6: TS/JS ディープ Null 逆参照」]
終わり
サブグラフ 解決策

イオン[「解決と検証の評決」]
VerdictSafe{"すべてのパスは合格しましたか?"}
拒否["❌ 即時拒否と反例"]
承認[「✅ 検証済みの安全なAST」]
終わり
サブグラフ ContextEngine ["✂️ AstDiffEngine & 🧠 SymbolGraph"]
キャッシュ["⚡ AstContextCache (FIPS 180-4 SHA-256)"]
スケルトン["コンテキスト圧縮 (>50-80% のトークン節約)"]
パッチ[「外科的バイトスパンノード置換」]
終わり
サブグラフ インターフェイス [「公開されたランタイム インターフェイス」]
CLI["💻 CLI バイナリ: 軌跡チェック / グラフ / パッチ"]
MCP["🔌 モデル コンテキスト プロトコル サーバー: locus mcp"]
LIB["📦 Rust ライブラリ クレート: locus_engine"]
終わり
RawCode --> P0 --> P1 --> P2 --> P3 --> P4 --> P5 --> P6 --> VerdictSafe
VerdictSafe -->|いいえ|拒否する
VerdictSafe -->|はい|承認する
承認 --> キャッシュ --> スケルトン --> パッチ
パッチ --> インターフェース
読み込み中
🛡️ 6 つの決定論的安全性の不変条件
グラフLR
A[AstGuard インバリアント パス] --> B[1.デリミタバランス：ダイクストラスタックスキャン]
A --> C[2.同時実行性: .await ポイント全体の std::sync::Mutex]
A --> D[3.算術演算: 保護されていない変数による除算]
A --> E[4.境界: 長さチェックなしの配列インデックス]
A --> F[5.パニック: 保護されていない .unwrap() および .expect()]
A --> G[6. ReDoS: 指数ネストされた正規表現数量子]
読み込み中
区切り文字バランス (ダイクストラ アルゴリズム): 文字列リテラルとエスケープを安全に無視しながら、生のバイト ストリーム全体で {} [] () の一致するクロージャを検証する線形シングルパス スタック スキャンを実行します。
非同期ミューテックス同時実行トラップ: 非同期 .await 一時停止ポイント全体での std::sync::Mutex ロックのブロックを防ぎ、スレッド プールの枯渇とデッドロックを排除します。
ゼロ除算保護: 算術評価を許可する前に、分母が非ゼロ ( $y \neq 0$ ) であることを証明します。
配列境界保護: 配列とスライスのインデックス付け ( arr[i] ) が長さアサーションの前にあることを保証します。

r 安全なアクセサー ( .get() )。
Unsafe Unwrap Guard: 事前の安全性チェック ( is_some() 、 is_ok() 、または if let ) を欠いた、パニックを引き起こす直接の .unwrap() または .expect() 呼び出しを排除します。
ReDoS Catastrophic Backtracking Guard: CPU 実行スレッドをフリーズさせる多項式および指数のネストされた数量子 ( (a+)+$ など) を特定します。
🔌 モデルコンテキストプロトコル (MCP) の統合
locus-engine には、標準入出力 (JSON-RPC 2.0) 上で実行される、依存性のない組み込み MCP サーバーが付属しています。 Claude Code、Claude Desktop、Cursor、Windsurf、および VS Code に直接接続します。
⚙️ クロードデスクトップ/カーソル設定
claude_desktop_config.json または Cursor MCP 設定に軌跡を追加します。
{
"mcpサーバー": {
"軌跡" : {
"コマンド" : "軌跡" ,
"args" : [ "mcp " ]
}
}
}
🛠️ 公開された MCP ツール:
MCP ツール名
引数
能力と出力
安全性を確認する
{"コード": "文字列"、"パス": "文字列"}
6パスAST検証を実行します。正確な違反バイト スパンを含む合格/失敗レポートを返します。
スケルトン化する
{"コード": "文字列", "言語": "rust|typescript|python"}
すべての署名を保持しながら実装本体を削除し、LLM コンテキスト トークンを >50 ～ 80% 節約します。
パッチシンボル
{"ソース": "文字列"、"シンボル": "文字列"、"new_code": "文字列"、"言語": "文字列"}
変更されていないコードを書き換えることなく、ターゲット関数/構造体のサージカルバイトオフセットノード置換を実行します。
インデックスグラフ
{"パス": "文字列"}
プロジェクト ディレクトリのインデックスを再帰的に作成し、定義を抽出し、ファイル間の依存関係エッジをマップします。
🚀 クイックインストールと CLI の使用法
🐧 Linux および 🍏 macOS:
カール -fsSL https://raw.githubusercontent.com/ahmadshady747-create/LOCUS/main/scripts/install.sh |バッシュ
🪟 Windows (PowerShell):
irm https://raw.githubusercontent.com/ahmadshady747 - create/LOCUS/main/scripts/install.ps1 |アイエックス
📦 貨物経由 (crates.io):

貨物設置軌跡エンジン
コマンドラインリファレンス
#1. 決定的な安全性検証 (<0.05ms)
軌跡チェック src/lib.rs
# 2. ワークスペース シンボル グラフのインデックス作成とトークンの節約の測定
軌跡グラフ src/
#3. 外科的シンボルのパッチング
locus patch src/models.rs --symbol User --with " pub struct User { pub id: u64 } "
# 4. モデル コンテキスト プロトコル (MCP) stdio サーバーを開始する
ローカスMCP
🛠️ Rust ライブラリの統合
locus-engine を Cargo.toml に追加します。
[ 依存関係 ]
軌跡エンジン = " 0.1.0 "
locus_engine :: { AstGuard 、 AstDiffEngine 、 SymbolGraph 、 AstContextCache 、 Language } を使用します。
fn メイン ( ) {
let code = "pub fnsafe_calc(a: f64, b: f64) -> f64 { if b != 0.0 { a / b } else { 0.0 } }" ;
// 1. インスタント不変安全性検証 (9µs)
let report = AstGuard :: verify ( code ) ;
主張してください！ (レポートが渡されました);
// 2. 外科的コンテキストの圧縮 (>70% のトークン節約)
let スケルトン = AstDiffEngine :: スケルトン化 ( コード , 言語 :: Rust ) ;
プリントイン！ ( "圧縮されたスケルトン: \n {}" , スケルトン ) ;
// 3. 高速インメモリ FIPS 180-4 SHA-256 LRU キャッシュ
let キャッシュ = AstContextCache :: new (1024) ;
ハッシュ = キャッシュにしましょう。挿入 (コード、スケルトン、1) ;
アサート_eq ! (ハッシュ .len() , 64 ) ;
}
📂 リポジトリの構造
d:\LOCUS\
§── Cargo.toml # シングルクレートパッケージマニフェスト (locus bin + locus_engine lib)
§── ライセンス # ビジネス ソース ライセンス 1.1 (明示的な現状のままの免責条項付き)
§── README.md # 包括的な技術文書とベンチマーク
§── SPEC.md # コアアルゴリズムの詳細な正式仕様
§── .gitattributes # GitHub Linguist 分類 (100% Rust プロジェクト)
§── スクリプト/
│ §── install.sh # Linux および macOS 用の 1 行のcurl インストーラー
│ └── install.ps1 # Windows 用の 1 行 PowerShell インストーラー
§── テスト/
│ ━─ ベンチマ

rks.rs # 高精度のベンチマークおよびストレス テスト スイート
━── src/
§── lib.rs # パブリックライブラリのエクスポート
§── main.rs # CLI エントリポイント (check、graph、patch、mcp コマンド)
§── types.rs # コアモデル (SymbolNode、SymbolEdge、VerificationReport)
§── Guard.rs # 6 パス決定論的 AST 安全不変条件エンジン
§──cache.rs # 単調インデックスを使用した純粋な FIPS 180-4 SHA-256 LRU キャッシュ
§──graph.rs # Polyglot シンボルグラフ & 依存関係リゾルバー (Rust、TS、Python)
§── diff.rs # 外科バイトスパン AST パッチングとスケルトンナイザー
━── mcp.rs # 依存性ゼロの stdio モデル コンテキスト プロトコル (MCP) サーバー
📄 ライセンスと商用レベル
locus-engine は、ビジネス ソース ライセンス 1.1 (BSL 1.1) に基づいて公開されています。
保証とサポートの免責事項 (現状のまま / セルフサービス): ソフトウェアは、いかなる種類の保証もなく、セルフサービス ベースで「現状のまま」提供されます。専用のテクニカル サポート、カスタム SLA 保証、エンタープライズ統合支援は、個別のカスタム契約に基づいて交渉されない限り含まれません。
ライセンスのアクティベーションおよび商用契約については、以下の著者に連絡するか、 licensing@locus.dev に電子メールを送信してください。
アーメド・シャディ (リビア🇱🇾) によって独立して設計および建設されました。
📘 Facebook: アーメド・シャディのプロフィール
🐙 GitHub: @ahmadshady747-create
📧 直接のお問い合わせ: GitHub の問題とディスカッション経由
Pure Rust のソブリン アンビエント HUD およびインバリアント検証済みマイクロカーネル IDE
0 フォーク リポジトリを報告

[切り捨てられた]

## Original Extract

Sovereign Ambient HUD & Invariant-Verified Microkernel IDE in Pure Rust - ahmadshady747-create/LOCUS

GitHub - ahmadshady747-create/LOCUS: Sovereign Ambient HUD & Invariant-Verified Microkernel IDE in Pure Rust · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
ahmadshady747-create
/
LOCUS
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
37 Commits 37 Commits Folders and files
.github/ workflows .github/ workflows scripts scripts src src tests tests .gitattributes .gitattributes .gitignore .gitignore Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SPEC.md SPEC.md View all files Repository files navigation
Deterministic AST Safety Guard, Polyglot Semantic Symbol Graph, Surgical Byte-Span Patching, and Zero-Dependency Model Context Protocol (MCP) Server in Pure Rust.
Modern AI code generation agents (Claude Code, Cursor, Copilot, Devin) and automated developer pipelines face two systemic engineering bottlenecks:
Probabilistic Syntax & Concurrency Regressions: AI agents frequently hallucinate unclosed delimiters, panic-inducing .unwrap() traps, async-mutex thread deadlocks, catastrophic polynomial regular expressions (ReDoS), and unbounded array indexing.
Context Window Inflation: Repeatedly feeding entire source code files into LLM context windows wastes up to 80% of token budgets on repetitive function bodies rather than high-level interface contracts.
locus-engine solves both challenges as a standalone, zero-bloat, high-performance systems engine written in 100% safe Rust. It enforces deterministic, non-negotiable safety invariants in microsecond time ( 9.04 µs ) , extracts cross-file symbol graphs with minimal context footprints, and communicates natively with modern AI IDEs via the Model Context Protocol (MCP) .
📺 Live Terminal Verification Demo
$ locus check src/async_task.rs
+-------------------------------------------------------------+
| LOCUS AST GUARD VERIFICATION |
+-------------------------------------------------------------+
Target File: src/async_task.rs
Verified Latency: 0.0194 ms
Status: [FAIL] Invariant Violation Detected
Violation Kind: ASYNC_MUTEX_DEADLOCK
Violation Detail: std::sync::Mutex used in async context with .await — use tokio::sync::Mutex instead.
+-------------------------------------------------------------+
$ locus graph src/
+-------------------------------------------------------------+
| LOCUS SYMBOL GRAPH INDEX |
+-------------------------------------------------------------+
Indexed Root: src/
Total Indexed Files: 8
Extracted AST Symbols: 28
Token Savings via AST Skeleton: 74.8%
Indexing Latency: 4.82 ms
+-------------------------------------------------------------+
📊 Empirical Benchmarks & Performance Metrics
Benchmarked under optimized release profile ( opt-level = 3 , lto = thin , codegen-units = 1 ):
Capability
locus-engine
Traditional Linters (ESLint, Clippy)
Cloud AI Guardrails
Verification Latency
9 µs – 0.05 ms (Nanosecond-scale)
250 – 1,500 ms (Process Spawns)
500 – 2,500 ms (Network Round-Trip)
Execution Architecture
In-Memory Pure Rust Kernel
Node.js / Python Runtime
Remote HTTP Cloud API
Context Token Savings
> 50% - 80% (AST Skeleton)
0% (Full Files)
0% (Full Files)
MCP Protocol Support
Built-In JSON-RPC 2.0 over Stdio
Requires Custom Wrappers
Proprietary APIs
Memory Safety
100% Safe Rust (0 Unsafe Blocks)
Varies (C/C++/Node)
Undefined
External Dependencies
Zero Crypto/Runtime Bloat
Heavy node_modules / Python env
Cloud Connection & API Keys
Deterministic Guarantee
100% Formal Invariant Rejection
Heuristic Warnings
Probabilistic LLM Re-evaluation
🏛️ System Architecture & Workflow Pipeline
flowchart TD
subgraph Input ["Incoming Code / AI Agent Patch"]
RawCode["Raw Code Snippet / File"]
end
subgraph AstGuardPipeline ["🛡️ AstGuard: 6-Pass Deterministic Firewall (<0.05ms)"]
P0["Pass 0: Delimiter Balance (Dijkstra)"]
P1["Pass 1: Async Mutex Across Await"]
P2["Pass 2: Division-by-Zero Guard"]
P3["Pass 3: Array Bounds Overflow"]
P4["Pass 4: Unsafe Unwrap / Expect Trap"]
P5["Pass 5: ReDoS Catastrophic Backtracking"]
P6["Pass 6: TS/JS Deep Null Dereference"]
end
subgraph Resolution ["Resolution & Verification Verdict"]
VerdictSafe{"All Passes Passed?"}
Reject["❌ Immediate Rejection & Counterexample"]
Approve["✅ Verified Safe AST"]
end
subgraph ContextEngine ["✂️ AstDiffEngine & 🧠 SymbolGraph"]
Cache["⚡ AstContextCache (FIPS 180-4 SHA-256)"]
Skeleton["Context Compression (>50-80% Token Savings)"]
Patch["Surgical Byte-Span Node Replacement"]
end
subgraph Interfaces ["Exposed Runtime Interfaces"]
CLI["💻 CLI Binary: locus check / graph / patch"]
MCP["🔌 Model Context Protocol Server: locus mcp"]
LIB["📦 Rust Library Crate: locus_engine"]
end
RawCode --> P0 --> P1 --> P2 --> P3 --> P4 --> P5 --> P6 --> VerdictSafe
VerdictSafe -->|No| Reject
VerdictSafe -->|Yes| Approve
Approve --> Cache --> Skeleton --> Patch
Patch --> Interfaces
Loading
🛡️ The 6 Deterministic Safety Invariants
graph LR
A[AstGuard Invariant Passes] --> B[1. Delimiter Balance: Dijkstra stack scan]
A --> C[2. Concurrency: std::sync::Mutex across .await points]
A --> D[3. Arithmetic: Unguarded division by variable]
A --> E[4. Bounds: Array index without length checks]
A --> F[5. Panics: Unguarded .unwrap() and .expect()]
A --> G[6. ReDoS: Exponential nested regex quantifiers]
Loading
Delimiter Balance (Dijkstra Algorithm): Performs a linear single-pass stack scan validating matching closure for {} [] () across raw byte streams while safely ignoring string literals and escapes.
Async Mutex Concurrency Trap: Prevents blocking std::sync::Mutex locks across asynchronous .await suspension points to eliminate thread pool exhaustion and deadlocks.
Division-by-Zero Protection: Proves the denominator is non-zero ( $y \neq 0$ ) before permitting arithmetic evaluation.
Array Bounds Protection: Ensures array and slice indexing ( arr[i] ) is preceded by length assertions or safe accessors ( .get() ).
Unsafe Unwrap Guard: Eliminates panic-inducing direct .unwrap() or .expect() calls lacking prior safety checks ( is_some() , is_ok() , or if let ).
ReDoS Catastrophic Backtracking Guard: Identifies polynomial and exponential nested quantifiers (such as (a+)+$ ) that freeze CPU execution threads.
🔌 Model Context Protocol (MCP) Integration
locus-engine ships with a built-in, zero-dependency MCP server running over stdio (JSON-RPC 2.0). It connects directly to Claude Code , Claude Desktop , Cursor , Windsurf , and VS Code .
⚙️ Claude Desktop / Cursor Configuration
Add locus to your claude_desktop_config.json or Cursor MCP settings:
{
"mcpServers" : {
"locus" : {
"command" : " locus " ,
"args" : [ " mcp " ]
}
}
}
🛠️ Exposed MCP Tools:
MCP Tool Name
Arguments
Capabilities & Output
check_safety
{"code": "string", "path": "string"}
Executes 6-pass AST verification; returns passed/failed report with exact violation byte span.
skeletonize
{"code": "string", "language": "rust|typescript|python"}
Strips implementation bodies while preserving all signatures, saving >50-80% LLM context tokens.
patch_symbol
{"source": "string", "symbol": "string", "new_code": "string", "language": "string"}
Performs surgical byte-offset node replacement of a target function/struct without rewriting unchanged code.
index_graph
{"path": "string"}
Recursively indexes project directory, extracts definitions, and maps cross-file dependency edges.
🚀 Quick Install & CLI Usage
🐧 Linux & 🍏 macOS:
curl -fsSL https://raw.githubusercontent.com/ahmadshady747-create/LOCUS/main/scripts/install.sh | bash
🪟 Windows (PowerShell):
irm https: // raw.githubusercontent.com / ahmadshady747 - create / LOCUS / main / scripts / install.ps1 | iex
📦 Via Cargo (crates.io):
cargo install locus-engine
Command Line Reference
# 1. Deterministic Safety Verification (<0.05ms)
locus check src/lib.rs
# 2. Index Workspace Symbol Graph & Measure Token Savings
locus graph src/
# 3. Surgical Symbol Patching
locus patch src/models.rs --symbol User --with " pub struct User { pub id: u64 } "
# 4. Start Model Context Protocol (MCP) stdio Server
locus mcp
🛠️ Rust Library Integration
Add locus-engine to your Cargo.toml :
[ dependencies ]
locus-engine = " 0.1.0 "
use locus_engine :: { AstGuard , AstDiffEngine , SymbolGraph , AstContextCache , Language } ;
fn main ( ) {
let code = "pub fn safe_calc(a: f64, b: f64) -> f64 { if b != 0.0 { a / b } else { 0.0 } }" ;
// 1. Instant Invariant Safety Verification (9µs)
let report = AstGuard :: verify ( code ) ;
assert ! ( report . passed ) ;
// 2. Surgical Context Compression (>70% Token Savings)
let skeleton = AstDiffEngine :: skeletonize ( code , Language :: Rust ) ;
println ! ( "Compressed Skeleton: \n {}" , skeleton ) ;
// 3. Fast In-Memory FIPS 180-4 SHA-256 LRU Cache
let cache = AstContextCache :: new ( 1024 ) ;
let hash = cache . insert ( code , skeleton , 1 ) ;
assert_eq ! ( hash . len ( ) , 64 ) ;
}
📂 Repository Anatomy
d:\LOCUS\
├── Cargo.toml # Single-crate package manifest (locus bin + locus_engine lib)
├── LICENSE # Business Source License 1.1 with explicit As-Is disclaimer
├── README.md # Comprehensive technical documentation & benchmarks
├── SPEC.md # Detailed formal specification of core algorithms
├── .gitattributes # GitHub Linguist classification (100% Rust project)
├── scripts/
│ ├── install.sh # One-line curl installer for Linux & macOS
│ └── install.ps1 # One-line PowerShell installer for Windows
├── tests/
│ └── benchmarks.rs # High-precision benchmark & stress test suite
└── src/
├── lib.rs # Public library exports
├── main.rs # CLI entrypoint (check, graph, patch, mcp commands)
├── types.rs # Core models (SymbolNode, SymbolEdge, VerificationReport)
├── guard.rs # 6-pass deterministic AST safety invariants engine
├── cache.rs # Pure FIPS 180-4 SHA-256 LRU cache with monotonic indexing
├── graph.rs # Polyglot symbol graph & dependency resolver (Rust, TS, Python)
├── diff.rs # Surgical byte-span AST patching and skeletonizer
└── mcp.rs # Zero-dependency stdio Model Context Protocol (MCP) server
📄 Licensing & Commercial Tiers
locus-engine is published under the Business Source License 1.1 (BSL 1.1) :
Warranty & Support Disclaimer (As-Is / Self-Service): The software is provided "AS IS" on a self-service basis without warranties of any kind. Dedicated technical support, custom SLA guarantees, and enterprise integration assistance are not included unless negotiated under a separate custom agreement.
For license activation and commercial contracts: Contact the author below or email licensing@locus.dev .
Architected & built independently by Ahmed Shadi (Libya 🇱🇾).
📘 Facebook: Ahmed Shadi Profile
🐙 GitHub: @ahmadshady747-create
📧 Direct Inquiries: Via GitHub Issues & Discussions
Sovereign Ambient HUD & Invariant-Verified Microkernel IDE in Pure Rust
0 forks Report repos

[truncated]
