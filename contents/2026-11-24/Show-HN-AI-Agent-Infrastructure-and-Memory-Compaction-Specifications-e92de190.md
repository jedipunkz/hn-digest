---
source: "https://github.com/grfdsasdvef/awesome-ai-agent-infrastructure"
hn_url: "https://news.ycombinator.com/item?id=50625268"
title: "Show HN: AI Agent Infrastructure and Memory Compaction Specifications"
article_title: "GitHub - grfdsasdvef/awesome-ai-agent-infrastructure: Curated Architecture Specifications, Memory Compaction, xgrammar Decoding & SWE-bench Evals · GitHub"
image: "https://opengraph.githubassets.com/fbd9856eca313c621ef2a7aec7eaf413c833c061c535d8ada3a14278ec9e0dee/grfdsasdvef/awesome-ai-agent-infrastructure"
author: "snowdensay"
captured_at: "2026-11-24T08:42:13Z"
capture_tool: "hn-digest"
hn_id: 50625268
score: 1
comments: 0
posted_at: "2026-11-24T08:16:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Agent Infrastructure and Memory Compaction Specifications

- HN: [50625268](https://news.ycombinator.com/item?id=50625268)
- Source: [github.com](https://github.com/grfdsasdvef/awesome-ai-agent-infrastructure)
- Score: 1
- Comments: 0
- Posted: 2026-11-24T08:16:42Z

## Translation

タイトル: HN を表示: AI エージェントのインフラストラクチャとメモリ圧縮の仕様
記事のタイトル: GitHub - grfdsasdvef/awesome-ai-agent-infrastructor: 厳選されたアーキテクチャ仕様、メモリ圧縮、xgrammar デコードおよび SWE ベンチ評価 · GitHub
説明: 厳選されたアーキテクチャ仕様、メモリ圧縮、xgrammar デコード、SWE ベンチ評価 - grfdsasdvef/awesome-ai-agent-infrastructor

記事本文:
GitHub - grfdsasdvef/awesome-ai-agent-infrastructor: 厳選されたアーキテクチャ仕様、メモリ圧縮、xgrammar デコード、SWE ベンチ評価 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 製品 プラットフォーム GitHub リポジトリ ホスティング、コード レビュー、コラボレーション
タスク、プロジェクト、自動化作業用の Copilot AI コーディング エージェント
ビルドと出荷アクションの CI/CD と自動化
コード品質 品質分析、修正、適用
Codespaces クラウド開発環境
セキュアコードのセキュリティ脆弱性の検出と修復
機密保護 機密スキャンと漏洩防止
開発者向け学習ドキュメント
変更ログ 変更ログ GitHub 製品全体の最新機能、改善点、修正のソースがすべて 1 か所にまとめられています。
事業段階別事業紹介
アドオン GitHub Advanced Security エンタープライズ グレードのコード セキュリティ
Copilot for business エンタープライズグレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
リソース Learn & Connect ドキュメント
GitHub Universe GitHub Universe 開発者による、開発者のためのイベント - トーク、ワークショップ、GitHub の次の展開。
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
grfdsasdvef
素晴らしい-ai-エージェント-インフラストラクチャ
weawe-ai-agent-infrastructor 公開通知 通知設定を変更するにはサインインする必要があります
スター 2 ( 2 ) リポジトリにスターを付けるにはサインインする必要があります
厳選されたアーキテクチャ仕様、メモリ圧縮、xgrammar デコード、SWE ベンチ評価
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ main ブランチ タグ ファイルに移動

e コード [その他のアクション] メニューを開く 最新のコミット
1 コミット 1 コミット フォルダーとファイル
.gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
素晴らしい AI エージェント インフラストラクチャ
精選されたアーキテクチャ仕様、メモリ エンジン、スキーマ制約デコーディング、SWE ベンチ評価、プロダクション ツール
⚡ スノーデン SAY ディープ システム エンジニアリング ラボ:
生産事故、LLM (Lost in the Middle) 注意力の低下、ローリング メモリ コンパクター、および低レベル エージェントのランタイムに関する毎日の詳細な分析は、Telegram チャネル: @snowdensay で行われます。
📖 完全なモノグラフ「AI エージェントのメモリとランタイム」 (コード + テスト): Telegram ボット @Snowden_book_bot で入手可能
1. エージェントのメモリとコンテキストの圧縮
2. スキーマに制約されたデコードおよび文法エンジン
3. SWEベンチと自動ASTコード編集
4. MCP (モデル コンテキスト プロトコル) の強化
5. サンドボックス化と分離 (Docker、gVisor、Firecracker)
6. ベンチマークと実際の生産指標
128k ～ 1M トークン コンテキスト ウィンドウを備えた最新の LLM は、メモリが無限であるという危険な錯覚を引き起こします。長時間の自律エージェントの実行 (40 回以上のツール回転) では、KV キャッシュ全体で注意の重みが不可逆的に薄まってしまいます (途中で失われる)。
失敗したツール呼び出しを決して要約しないでください。生のエラー トレースバックは、修正されるまでそのままにしておく必要があります。修正されない場合、エージェントは無限に循環します。
プロンプト検証よりも文法に制約されたデコード: ソフトマックスの前に CFG ロジット マスクを介して JSON スキーマを強制します (有効性 99.97% 対プロンプト プロンプトありの 92%)。
圧縮ルーチンをロックする: asyncio.Lock を使用しない場合、非同期ループで並列ツールを呼び出すと競合状態が発生し、二重圧縮が発生します。
1. エージェントのメモリとコンテキストの圧縮
三層圧延会社

アクションアーキテクチャ:
[ レイヤー 1: 固定された不変量 ] --- システム コントラクト、安全境界、ツール (未加工、0% ドロップ)
│
[ レイヤー 2: 圧縮された台帳 ] ─── 完了した成功したツール サイクルの重複排除された事実
│
[ レイヤー 3: Hot Raw Tail ] --- 最後の 4 ～ 6 ターン (新鮮なローカル状態 + 未解決のスタックトレース)
参考実装:
RollingMemoryCompactor (Python/Go): 冪等ファイル読み取り ( head 、 cat ) を SHA256 ポインターに自動的に重複排除します。
エピソード KV キャッシュ オフロード: バックグラウンド コンテキストの中間層 KV トークンをプルーニングします。
2. スキーマに制約されたデコードおよび文法エンジン
xgrammar : vLLM および SGLang 用の高速 GPU アクセラレーションによる文法ガイド付きデコーディング。
概要 : モデル語彙に対する有限状態マシン (FSM) インデックスによる構造化テキストの生成。
ガイダンス : トークンレベルの正規表現制約を備えたインターリーブ生成および制御ロジック。
制作例 (vLLM + xgrammar):
from vllm import LLM 、SamplingParams
vllm から。 sampling_params インポート GuidedDecodingParams
スキーマ = {
"タイプ" : "オブジェクト" ,
"プロパティ" : {
"tool" : { "enum" : [ "patch_ast" , "execute_test" , "commit_fix" ]},
"ファイルパス" : { "タイプ" : "文字列" , "パターン" : r"^[\w\-\./]+\.[a-z]+$" },
"パッチコンテンツ" : { "タイプ" : "文字列" }
}、
"必須" : [ "ツール" , "ファイルパス" , "パッチコンテンツ" ],
"追加プロパティ" : False
}
params = SamplingParams (
温度 = 0.1、
max_tokens = 1024 、
guided_decoding = GuidedDecodingParams ( json = スキーマ )
)
3. SWEベンチと自動ASTコード編集
Tree-sitter : AST を認識した差分のための増分具体的な構文ツリー解析。
SWE-bench : 実際の G​​itHub の問題を解決するために LLM を評価するためのベンチマーク。
構文ロールバック ガード : コンパイラを呼び出す前の解析時に無効なパッチを元に戻します。
4. MCP (モデル コンテキスト プロトコル) の強化
プロンプトイン

ツールでの防御: 信頼できない MCP サーバーからの動的ツール定義を厳密にサニタイズします。
Zombie Process Reaper : アイドル状態の TTL を超える標準入出力サブプロセスを自動的にクリーンアップします。
同時実行速度制限: LLM の再試行の混雑を防ぐための非ブロッキングの選択トークン バケット。
gVisor : コンテナー セキュリティのためのアプリケーション カーネル。ユーザー空間で Linux syscall をインターセプトします。
Firecracker : 信頼できない Python コードを実行するため、5 ミリ秒未満で起動する一時的な microVM。
nsjail : 名前空間と seccomp-bpf を使用した軽量の Linux プロセス分離ツール。
6. ベンチマークと実際の生産指標
SWE ベンチでテスト済み 検証済みサブセット (1 日あたり 17,000 ツール反復、4×A100 上の Llama-3.3-70B):
ディスカッションとアーキテクチャの青写真は Snowden SAY によって管理されています。
📡 電報エンジニアリング チャンネル: @snowdensay
📖 モノグラフリポジトリボット: @Snowden_book_bot
厳選されたアーキテクチャ仕様、メモリ圧縮、xgrammar デコード、SWE ベンチ評価
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Curated Architecture Specifications, Memory Compaction, xgrammar Decoding & SWE-bench Evals - grfdsasdvef/awesome-ai-agent-infrastructure

GitHub - grfdsasdvef/awesome-ai-agent-infrastructure: Curated Architecture Specifications, Memory Compaction, xgrammar Decoding & SWE-bench Evals · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
grfdsasdvef
awesome-ai-agent-infrastructure
awesome-ai-agent-infrastructure Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Curated Architecture Specifications, Memory Compaction, xgrammar Decoding & SWE-bench Evals
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
AWESOME-AI-AGENT-INFRASTRUCTURE
Curated Architecture Specifications, Memory Engines, Schema-Constrained Decoding, SWE-bench Evals & Production Tooling
⚡ Snowden SAY Deep Systems Engineering Lab:
Ежедневные глубокие разборы аварий на проде, деградация внимания LLM (Lost in the Middle), rolling memory compactor и low-level агентные рантаймы — в нашем Telegram-канале: @snowdensay .
📖 Полная монография «Память и рантайм AI-агентов» (код + тесты): доступна в Telegram-боте @Snowden_book_bot
1. Agent Memory & Context Compaction
2. Schema-Constrained Decoding & Grammar Engines
3. SWE-bench & Automated AST Code Editing
4. MCP (Model Context Protocol) Hardening
5. Sandboxing & Isolation (Docker, gVisor, Firecracker)
6. Benchmarks & Real Production Metrics
Modern LLMs with 128k–1M token context windows create a dangerous illusion of infinite memory. In long autonomous agent runs (40+ tool turns), attention weights irreversibly dilute across the KV cache ( Lost in the Middle ).
Never summarize failed tool calls: Raw error tracebacks must remain verbatim until fixed, or the agent cycles infinitely.
Grammar-constrained decoding over prompt validation: Enforce JSON Schemas via CFG logits masks prior to softmax (99.97% validity vs 92% with prompt prompting).
Lock your compaction routines: Parallel tool invocations in async loops trigger race conditions and double-compaction without asyncio.Lock .
1. Agent Memory & Context Compaction
Three-tier rolling compaction architecture:
[ LAYER 1: Pinned Invariants ] ─── System contract, safety boundaries, tools (Raw, 0% drop)
│
[ LAYER 2: Compacted Ledger ] ─── Deduplicated facts of completed successful tool cycles
│
[ LAYER 3: Hot Raw Tail ] ─── Last 4-6 turns (Fresh local state + unresolved stacktraces)
Reference Implementations:
RollingMemoryCompactor (Python/Go): Automatic deduplication of idempotent file reads ( head , cat ) into SHA256 pointers.
Episodic KV Cache Offloading : Pruning intermediate layer KV tokens for background context.
2. Schema-Constrained Decoding & Grammar Engines
xgrammar : Fast GPU-accelerated grammar-guided decoding for vLLM and SGLang.
outlines : Structured text generation via finite-state machine (FSM) index over model vocabulary.
guidance : Interleaved generation and control logic with token-level regex constraints.
Production Example (vLLM + xgrammar):
from vllm import LLM , SamplingParams
from vllm . sampling_params import GuidedDecodingParams
schema = {
"type" : "object" ,
"properties" : {
"tool" : { "enum" : [ "patch_ast" , "execute_test" , "commit_fix" ]},
"file_path" : { "type" : "string" , "pattern" : r"^[\w\-\./]+\.[a-z]+$" },
"patch_content" : { "type" : "string" }
},
"required" : [ "tool" , "file_path" , "patch_content" ],
"additionalProperties" : False
}
params = SamplingParams (
temperature = 0.1 ,
max_tokens = 1024 ,
guided_decoding = GuidedDecodingParams ( json = schema )
)
3. SWE-bench & Automated AST Code Editing
Tree-sitter : Incremental concrete syntax tree parsing for AST-aware diffing.
SWE-bench : Benchmark for evaluating LLMs on resolving real-world GitHub issues.
Syntactic Rollback Guards : Reverting invalid patches at parse-time before invoking compilers.
4. MCP (Model Context Protocol) Hardening
Prompt Injection Defense in Tools : Strict sanitization of dynamic tool definitions from untrusted MCP servers.
Zombie Process Reaper : Automatic cleanup of stdio subprocesses exceeding idle TTL.
Concurrency Rate Limiting : Non-blocking select token buckets to prevent LLM retry stampedes.
gVisor : Application kernel for container security, intercepting Linux syscalls in user space.
Firecracker : Ephemeral microVMs booting in < 5 ms for untrusted Python code execution.
nsjail : Lightweight Linux process isolation tool using namespaces and seccomp-bpf.
6. Benchmarks & Real Production Metrics
Tested on SWE-bench Verified subset (17,000 tool iterations/day, Llama-3.3-70B on 4×A100):
Discussions and architecture blueprints are maintained by Snowden SAY .
📡 Telegram Engineering Channel: @snowdensay
📖 Monograph Repository Bot: @Snowden_book_bot
Curated Architecture Specifications, Memory Compaction, xgrammar Decoding & SWE-bench Evals
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
