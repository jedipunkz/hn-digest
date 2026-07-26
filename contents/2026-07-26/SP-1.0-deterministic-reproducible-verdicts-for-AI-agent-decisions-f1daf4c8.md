---
source: "https://github.com/Fame510/SHACKLE/blob/master/SP-1.0-SPECIFICATION.md"
hn_url: "https://news.ycombinator.com/item?id=49060407"
title: "SP/1.0: deterministic, reproducible verdicts for AI-agent decisions"
article_title: "SHACKLE/SP-1.0-SPECIFICATION.md at master · Fame510/SHACKLE · GitHub"
author: "shacklepro"
captured_at: "2026-07-26T17:54:34Z"
capture_tool: "hn-digest"
hn_id: 49060407
score: 2
comments: 0
posted_at: "2026-07-26T17:41:16Z"
tags:
  - hacker-news
  - translated
---

# SP/1.0: deterministic, reproducible verdicts for AI-agent decisions

- HN: [49060407](https://news.ycombinator.com/item?id=49060407)
- Source: [github.com](https://github.com/Fame510/SHACKLE/blob/master/SP-1.0-SPECIFICATION.md)
- Score: 2
- Comments: 0
- Posted: 2026-07-26T17:41:16Z

## Translation

タイトル: SP/1.0: AI エージェントの決定に対する決定的で再現可能な判定
記事タイトル: SHACKLE/SP-1.0-SPECIFICATION.md at master · Fame510/SHACKLE · GitHub
説明: SHACKLE — 自律 AI エージェント用のガバナンス プロトコルおよびポリシー決定デーモン。監査済みの意思決定エンジン、SP-1.0 プロトコル仕様、Rust/TypeScript クライアント、および SOC2 に準拠したコンプライアンス ツールを使用して、ガードレール、予算/ループ制限、ポリシー制約をリアルタイムで強制します。 - シャックル/SP-1.0-SPE
[切り捨てられた]

記事本文:
SHACKLE/SP-1.0-SPECIFICATION.md (マスター) · Fame510/SHACKLE · GitHub
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
名声510
/
シャックル
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
c にサインインする必要があります

ハンゲの通知設定
追加のナビゲーション オプション
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 748 行 (580 loc) · 26.5 KB マスター ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション SHACKLE プロトコル仕様 — SP/1.0
自律型 AI エージェント用のランタイム サーキット ブレーカー
バージョン: 1.0.0
ステータス: 公開済み
日付: 2026-06-25
著者: ダンテ・ブロック、ソブリン・ロジック
ライセンス: クリエイティブ コモンズ 表示 4.0 インターナショナル (CC-BY 4.0)
参考実装: https://github.com/Fame510/SHACKLE
最初のパブリックコミット: 2026-06-17 23:12 UTC
この仕様の実装には、SHACKLE ライセンス条項が適用されます。
リファレンス実装は、AGPLv3 (オープンソース) と AGPLv3 (オープンソース) のデュアルライセンスです。
商用 (独自)。連絡先：docspoc101@gmail.com
SHACKLE は、自律型 AI エージェント用のランタイム サーキット ブレーカーです。それは 1 つの質問に答えます。
現時点では、このエージェントにこれらのパラメータを使用してこのツールを実行することを許可する必要がありますか?
このプロトコルは、9 つ​​の数学的不変式に裏付けられた決定論的で検証可能な決定関数、言語に依存しないメッセージ スキーマ、Ed25519 署名付き追加専用監査ログ、および Redis を利用した分散状態エンジンを定義します。 gRPC/Unix ソケット トランスポートを備えたサイドカー デーモンとして、または単一エージェント デプロイメントのインプロセス ライブラリとして動作します。
SHACKLE は、暗号化監査の加工管理を備えた AI エージェント向けの初のオープンソース ランタイム サーキット ブレーカーです。この仕様は、SP/1.0 プロトコルの最終的なリファレンスです。
自律型 AI エージェントは、ランタイムの監視なしで、Web 検索、ファイル I/O、API 呼び出し、コード実行などのツールを実行します。フレームワークの再帰制限またはトークン キャップが唯一のガードレールです。エージェントが再試行ループに入ったとき (同じ

ツール、同じエラー、毎回トークンの書き込み）、ウォレットが空になる前にそれを検出、傍受、停止するメカニズムはありません。
これは仮説ではありません。運用環境の展開では、次のことが文書化されています。
エージェントの無限ループにより、再帰制限が適用される前に 6,000 ドル以上の API コストが消費される
重複したツール呼び出しが変化なしで 50 回以上繰り返される
生成された子プロセスがトークンの消費中に無期限にハングする
2026 年 6 月に複数のチームが独自に達成した業界のコンセンサスは、生成権限はリリース権限ではないということです。モデルは候補を生成します。別のメディエーション層で実行を承認する必要があります。
SHACKLE はその仲介層です。
原則
意味
決定論的コア
Decide(state, call) → Verdict は純粋な関数です。同じ入力は常に同じ出力を生成します。
権限としてのデーモン
SHACKLE デーモンは、時間、状態、および判定に関する唯一の信頼できる情報源です。エージェントは信頼されていません。
追加のみの監査
すべての決定は Ed25519 で署名され、不変の監査ログに書き込まれます。保管過程は暗号的に検証可能です。
数学的に検証済み
9 つの不変プロパティがすべての入力の下で保持され、プロパティベースのテストによって証明されています (仮説、それぞれ 500 以上の例)。
グレースフルデグラデーション
エージェントは、デーモンなしでローカル/ライブラリ モードで機能します。分散状態はアップグレード パスです。
フェールクローズ
ネットワーク障害、デーモンのクラッシュ、またはタイムアウト → 拒否。明示的な許可がなければ実行できません。
1.3 範囲
決定関数とその 9 つの数学的不変量 (§3)
メッセージのスキーマとセマンティクス (§4)
この仕様には以下の内容は含まれません。
デーモン永続層 (実装の詳細)
HITL コンソール UI (プレゼンテーションの問題)
モデル A — ライブラリ モード (インプロセス)
┌───────────────┐
│ エージェントのプロセス │
│ ┌─

──────────┐ │
│ │ @Guard デコレータ │ │
│ │ 地方州のみ │ │
│ ━━━━━━━━━┘ │
━━━━━━━━━━━┘
モデル B — サイドカー デーモン (実稼働)
┌─────────────┐ Unix/gRPC ┌───────────┐
│ エージェントプロセス │ ◄─────────────► │ SHACKLE デーモン │
│ ┌───────┐ │ pre_exec │ ┌─────────┐ │
│ │ シン │ │ post_exec │ │ ポリシー エンジン │ │
│ │ クライアント │ │ 登録 │ │ - 予算 │ │
│ │ シム │ │ 心拍数 │ │ - カウンター │ │
│ └───────┘ │ │ │ - サーキットブレーカー │ │
━━━━━━━━━━━━━━━━━━━┘ │
│ ┌───────────┐ │
│ │ 監査ログ │ │
│ │ Ed25519 署名入り │ │
│ │ 追加のみ │ │
│ │ チェーンリンク │ │
│ ━━━━━━━━━━━┘ │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
モデル C — 分散型 (エンタープライズ)
┌─────┐ ┌─────┐ ┌─────┐
│ エージェント A │ │ エージェント B │ │ エージェント C │
━━━┬─────┘ └────┬──────┘ └────┬──────┘
━─────────

┬─────────┘
│ gRPC/TLS
┌───┴───┐
│ シャックル │
│ デーモンクラスタ │
│ Redis (状態) │
│ Postgres (ログ)│
━━━━━━━━┘
2.2 プロトコル層
┌─────────────────┐
│ ポリシー言語 (将来) │ ← ガード ルールの DSL
━━━━━━━━━━━━━━━━━━━━━━━┤
│ 決定関数 │ ← 決定(状態, 呼び出し) → 判定
━━━━━━━━━━━━━━━━━━━━━━━┤
│ メッセージプロトコル │ ← この仕様
━━━━━━━━━━━━━━━━━━━━━━━┤
│ トランスポート (Unix/gRPC/WS) │ ← バインディング層
━━━━━━━━━━━━━━━━━━━━━━┘
3. 決定関数
判定機能はSHACKLEの心臓部です。これは純粋な関数であり、I/O、副作用、ホット パスへの割り当てはありません。人間による音声は 10 分以内に聞こえるようになります。ロジックは 200 行未満です。
関数決定(
状態: セッション状態、
呼び出し: ToolCall、
構成: GuardConfig、
rng_float: 浮動小数点
) → 評決
3.2 決定アルゴリズム — 8 つの積層レイヤー
レイヤ 1: サーキットブレーカー
IF state.circuit_tripped:
→ 拒否(CIRCUIT_OPEN)
レイヤ 2: ノンス検証 (アンチリプレイ)
IF call.nonce IN state.seen_nonces:
→ 拒否(POLICY_VIOLATION)
レイヤ 3: バジェット ガード
config.budget_usd > 0の場合:
IF state.budget_remaining_usd <= 0:
→ 拒否(BUDGET_EXHAUSTED)
hitl_mode == ON_THRESHOLD かつ小数 <= しきい値の場合:
→ HITL("予算しきい値に達しました

")
IF call.cost > state.budget_remaining:
hitl_mode IN (ON_DENY、ALWAYS) の場合:
→ HITL(「コストが残りの予算を超えています」)
→ 拒否(BUDGET_EXHAUSTED)
レイヤ 4: リピート コール ガード
config.max_repeat_calls > 0の場合:
IF is_repeat(call, state.last_call):
制限 = max_repeat_calls
error_amplification AND has_error_signal(call) の場合:
制限 = 最大(1, 制限 - 1)
繰り返し回数 >= 制限の場合:
→ 拒否(MAX_REPEAT_EXCEEDED)
レイヤ 5: タイム ウィンドウ ガード
config.window_max_calls > 0の場合:
IF window_count >= window_max_calls:
→ 拒否(WINDOW_EXCEEDED)
レイヤ 6: グローバル通話制限
config.max_total_calls > 0 かつ total_calls >= max_total_calls の場合:
→ 拒否(GLOBAL_LIMIT)
レイヤ 7: 確率的拒否 (敵対的強化)
probabilistic_deny AND Budget_ratio < 0.2 の場合:
IF rng < 拒否ジッター率 * (1.0 - 予算率):
→ 拒否(BUDGET_EXHAUSTED、確率=true)
レイヤ 8: 常に HITL
hitl_mode == 常に指定する場合:
→ HITL(「すべての通話に HITL が必要」)
デフォルト:
→許可する
3.3 数学的不変条件 (すべての入力の下で保持する必要がある)
これら 9 つのプロパティは、それぞれ 500 以上のランダムに生成された入力を使用したプロパティベースのテスト (仮説、Python) によって検証されます。プロパティ P1 ～ P9 は正しいことが証明されています。
評決
意味
エージェントのアクション
許可する
要求どおりに実行する
ツール呼び出しを続行します
拒否します
ブロック実行
中止、表面拒否理由
ヒットル
人間の判断が必要
一時停止し、HITL コンソール経由で人間の判断を待ちます
3.5 拒否の理由
理由
トリガー
サーキット_オープン
以前にサーキットブレーカーが落ちた
BUDGET_EXHAUSTED
残りの予算 ≤ 0 またはコスト > 残り
MAX_REPEAT_EXCEEDED
同じツールと同じパラメータが何度も繰り返されています
WINDOW_EXCEEDED
現在の時間枠内のコールが多すぎます
GLOBAL_LIMIT
セッション全体の呼び出し制限に達しました
ポリシー違反
ノンスの重複 (リプレイ攻撃)
認証_失敗
認証失敗
3.6 エラー信号の増幅
シャックル検出

正規表現を含まないツール パラメーター内の ts エラー シグナル (ReDoS 攻撃対象領域なし)。 error_amplification が有効な場合、パラメータに既知のエラー信号が含まれている場合、繰り返し呼び出しのしきい値は 1 減ります。
エラー信号: 401、不正、403、禁止、500、502、503、
504、タイムアウト、接続拒否、接続リセット、
許可が拒否されました、レート制限、クォータを超過しました、
無効な API キー、トークンの有効期限が切れ、モデルが見つかりません、
リソースが使い果たされ、期限を超えました
これにより、完全な繰り返ししきい値を待たずに、エージェントが 401 に到達し、再試行し、401 に到達し、再試行するという「死のループ」を捕らえます。
3.7 確率的拒否 (敵対的強化)
probabilistic_deny が有効で、エージェントの予算が 20% を下回っている場合、ランダム要素が導入されます。
確率 = 拒否ジッター率 × (1.0 − 予算率)
IF rng < 確率:
→ 拒否 (BUDGET_EXHAUSTED として表示)
これにより、敵対的なエージェントが正確な予算のしきい値を調べてそれを不正使用することを防ぎます。この否定は実際の予算の枯渇と同じように見え、信号とノイズを区別できなくなります。
すべての SHACKLE メッセージはプロトコル バッファー エンベロープでラップされます。
メッセージ封筒 {
文字列プロトコルバージョン = 1 ; // "1.0.0"
文字列メッセージID = 2 ; // UUIDv7、クライアント生成
文字列相関ID = 3 ; // リクエスト/レスポンスのペアリング
int64 client_timestamp_ns = 4 ; // クライアントの時計 (情報)
int64 サーバータイムスタンプ_ns = 5 ; // 受信時にデーモンによって設定される
バイト hmac = 6 ; // ペイロード上の HMAC-SHA256
ペイロードの 1 つ {
PreExecRequest pre_exec = 10 ;
PreExecResponse pre_exec_response = 11 ;
PostExecNotification post_exec = 12 ;
RegisterRequest レジスタ = 13 ;
レジスタ応答 register_response = 14 ;
ハートビート ハートビート = 15 ;
ハートビートAck ハートビート_ACK = 16 ;
エラーエラー = 17 ;
}
}
4.2 セッション登録
メッセージ RegisterRequest {
文字列

gent_id = 1 ;
文字列エージェントバージョン = 2 ;
文字列フレームワーク = 3 ; // "クレワイ" | "自動生成" | 「ランググラフ」
文字列セッションID = 4 ; // オプション: 既存のセッションを再開します
文字列組織ID = 5 ;
文字列ランタイム = 6 ;
マップ < 文字列 , 文字列 > メタデータ = 7 ;
}
メッセージ RegisterResponse {
文字列セッションID = 1 ;
文字列デーモンバージョン = 2 ;
文字列ネゴシエートプロトコル = 3 ;
GuardConfig active_config = 4 ;
int64 デーモン_time_ns = 5 ;
}
4.3 実行前チェック
メッセージ PreExecRequest {
文字列セッションID = 1 ;
uint64 呼び出し番号 = 2 ; // 単調増加
文字列ツール名 = 3 ;
バイトtool_params_hash = 4 ; // 正規の JSON パラメータの SHA-256
ダブル推定コスト_米ドル = 5 ;
文字列親ガードID = 6 ; // ネストされたガード ツリーの場合
uint64 ノンス = 7 ; // アンチリプレイ
マップ < 文字列 , 文字列 > タグ = 8 ;
}
メッセージ PreExecResponse {
文字列セッションID = 1 ;
uint64 呼び出し番号 = 2 ;
評決評決 = 3 ;
拒否理由拒否理由 = 4 ;
文字列 human_readable_reason = 5 ;
ダブルブ

[切り捨てられた]

## Original Extract

SHACKLE — a governance protocol and policy-decision daemon for autonomous AI agents. Enforces guardrails, budget/loop limits, and policy constraints in real time, with an audited decision engine, SP-1.0 protocol spec, Rust/TypeScript clients, and SOC2-aligned compliance tooling. - SHACKLE/SP-1.0-SPE
[truncated]

SHACKLE/SP-1.0-SPECIFICATION.md at master · Fame510/SHACKLE · GitHub
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
Fame510
/
SHACKLE
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 748 lines (580 loc) · 26.5 KB master Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions SHACKLE Protocol Specification — SP/1.0
Runtime Circuit Breaker for Autonomous AI Agents
Version: 1.0.0
Status: Published
Date: 2026-06-25
Authors: Dante Bullock, Sovereign Logic
License: Creative Commons Attribution 4.0 International (CC-BY 4.0)
Reference Implementation: https://github.com/Fame510/SHACKLE
First Public Commit: 2026-06-17 23:12 UTC
Implementations of this specification are subject to the SHACKLE license terms.
The reference implementation is dual-licensed: AGPLv3 (open source) and
Commercial (proprietary). Contact: docspoc101@gmail.com
SHACKLE is a runtime circuit breaker for autonomous AI agents. It answers one question:
Should this agent be allowed to execute this tool with these parameters at this moment?
The protocol defines a deterministic, verifiable decision function backed by 9 mathematical invariants, a language-agnostic message schema, Ed25519-signed append-only audit logging, and a Redis-backed distributed state engine. It operates as a sidecar daemon with gRPC/Unix socket transport, or as an in-process library for single-agent deployments.
SHACKLE is the first open-source runtime circuit breaker for AI agents with cryptographic audit chain-of-custody. This specification is the definitive reference for the SP/1.0 protocol.
Autonomous AI agents execute tools — web search, file I/O, API calls, code execution — with no runtime oversight. The framework's recursion limit or token cap is the only guardrail. When an agent enters a retry loop (same tool, same error, burning tokens each time), there is no mechanism to detect, intercept, and stop it before the wallet is empty.
This is not hypothetical. Production deployments have documented:
Agent infinite loops consuming $6,000+ in API costs before the recursion limit fired
Duplicate tool calls repeating 50+ times with no variation
Spawned child processes hanging indefinitely while consuming tokens
The industry consensus — independently reached by multiple teams in June 2026 — is that generation authority is not release authority. The model generates candidates. A separate mediation layer must authorize execution.
SHACKLE is that mediation layer.
Principle
Meaning
Deterministic core
decide(state, call) → Verdict is a pure function. Same inputs always produce same outputs.
Daemon as authority
The SHACKLE daemon is the sole source of truth for time, state, and verdicts. Agents are untrusted.
Append-only audit
Every decision is Ed25519-signed and written to an immutable audit log. Chain-of-custody is cryptographically verifiable.
Mathematically verified
9 invariant properties hold under all inputs, proven by property-based testing (Hypothesis, 500+ examples each).
Graceful degradation
Agents function in local/library mode without a daemon. Distributed state is an upgrade path.
Fail-closed
Network failure, daemon crash, or timeout → DENY. No execution without explicit authorization.
1.3 Scope
The decision function and its 9 mathematical invariants (§3)
Message schemas and semantics (§4)
This specification does NOT cover:
Daemon persistence layer (implementation detail)
HITL console UI (presentation concern)
MODEL A — Library Mode (In-Process)
┌─────────────────────────┐
│ Agent Process │
│ ┌───────────────────┐ │
│ │ @Guard decorator │ │
│ │ Local state only │ │
│ └───────────────────┘ │
└─────────────────────────┘
MODEL B — Sidecar Daemon (Production)
┌─────────────────┐ Unix/gRPC ┌──────────────────────────┐
│ Agent Process │ ◄────────────────► │ SHACKLE Daemon │
│ ┌───────────┐ │ pre_exec │ ┌────────────────────┐ │
│ │ Thin │ │ post_exec │ │ Policy Engine │ │
│ │ Client │ │ register │ │ - Budgets │ │
│ │ Shim │ │ heartbeat │ │ - Counters │ │
│ └───────────┘ │ │ │ - Circuit Breakers │ │
└─────────────────┘ │ └────────────────────┘ │
│ ┌────────────────────┐ │
│ │ Audit Log │ │
│ │ Ed25519-signed │ │
│ │ Append-only │ │
│ │ Chain-linked │ │
│ └────────────────────┘ │
└──────────────────────────┘
MODEL C — Distributed (Enterprise)
┌──────────┐ ┌──────────┐ ┌──────────┐
│ Agent A │ │ Agent B │ │ Agent C │
└────┬─────┘ └────┬─────┘ └────┬─────┘
└─────────────┬─────────────┘
│ gRPC/TLS
┌────────┴────────┐
│ SHACKLE │
│ Daemon Cluster │
│ Redis (state) │
│ Postgres (logs)│
└─────────────────┘
2.2 Protocol Layers
┌──────────────────────────────────┐
│ Policy Language (future) │ ← DSL for guard rules
├──────────────────────────────────┤
│ Decision Function │ ← decide(state, call) → Verdict
├──────────────────────────────────┤
│ Message Protocol │ ← This specification
├──────────────────────────────────┤
│ Transport (Unix/gRPC/WS) │ ← Binding layer
└──────────────────────────────────┘
3. The Decision Function
The decision function is the heart of SHACKLE. It is a pure function — no I/O, no side effects, no allocations in the hot path. It is human-auditable in under 10 minutes. It is under 200 lines of logic.
function decide(
state: SessionState,
call: ToolCall,
config: GuardConfig,
rng_float: float
) → Verdict
3.2 Decision Algorithm — 8 Stacked Layers
Layer 1: Circuit Breaker
IF state.circuit_tripped:
→ DENY(CIRCUIT_OPEN)
Layer 2: Nonce Validation (Anti-Replay)
IF call.nonce IN state.seen_nonces:
→ DENY(POLICY_VIOLATION)
Layer 3: Budget Guard
IF config.budget_usd > 0:
IF state.budget_remaining_usd <= 0:
→ DENY(BUDGET_EXHAUSTED)
IF hitl_mode == ON_THRESHOLD AND fraction <= threshold:
→ HITL("budget threshold reached")
IF call.cost > state.budget_remaining:
IF hitl_mode IN (ON_DENY, ALWAYS):
→ HITL("cost exceeds remaining budget")
→ DENY(BUDGET_EXHAUSTED)
Layer 4: Repeat Call Guard
IF config.max_repeat_calls > 0:
IF is_repeat(call, state.last_call):
limit = max_repeat_calls
IF error_amplification AND has_error_signal(call):
limit = max(1, limit - 1)
IF repeat_count >= limit:
→ DENY(MAX_REPEAT_EXCEEDED)
Layer 5: Time Window Guard
IF config.window_max_calls > 0:
IF window_count >= window_max_calls:
→ DENY(WINDOW_EXCEEDED)
Layer 6: Global Call Limit
IF config.max_total_calls > 0 AND total_calls >= max_total_calls:
→ DENY(GLOBAL_LIMIT)
Layer 7: Probabilistic Denial (Adversarial Hardening)
IF probabilistic_deny AND budget_ratio < 0.2:
IF rng < deny_jitter_ratio * (1.0 - budget_ratio):
→ DENY(BUDGET_EXHAUSTED, probabilistic=true)
Layer 8: HITL Always
IF hitl_mode == ALWAYS:
→ HITL("HITL required for all calls")
Default:
→ ALLOW
3.3 Mathematical Invariants (Must Hold Under All Inputs)
These 9 properties are verified by property-based testing (Hypothesis, Python) with 500+ randomly generated inputs each. Properties P1-P9 are provably correct.
Verdict
Meaning
Agent Action
ALLOW
Execute as requested
Proceed with tool call
DENY
Block execution
Abort, surface deny reason
HITL
Human decision required
Pause, await human verdict via HITL console
3.5 Deny Reasons
Reason
Trigger
CIRCUIT_OPEN
Circuit breaker was previously tripped
BUDGET_EXHAUSTED
Budget remaining ≤ 0 or cost > remaining
MAX_REPEAT_EXCEEDED
Same tool + same params repeated too many times
WINDOW_EXCEEDED
Too many calls in the current time window
GLOBAL_LIMIT
Session-wide call limit reached
POLICY_VIOLATION
Duplicate nonce (replay attack)
AUTH_FAILED
Authentication failure
3.6 Error Signal Amplification
SHACKLE detects error signals in tool parameters WITHOUT regex (no ReDoS attack surface). When error_amplification is enabled, the repeat call threshold is reduced by 1 if the parameters contain known error signals:
Error signals: 401, unauthorized, 403, forbidden, 500, 502, 503,
504, timeout, connection refused, connection reset,
permission denied, rate limit, quota exceeded,
invalid api key, token expired, model not found,
resource exhausted, deadline exceeded
This catches the "loop of death" — agent hits 401, retries, gets 401, retries — without waiting for the full repeat threshold.
3.7 Probabilistic Denial (Adversarial Hardening)
When probabilistic_deny is enabled and the agent is below 20% budget, a random factor is introduced:
probability = deny_jitter_ratio × (1.0 − budget_ratio)
IF rng < probability:
→ DENY (appears as BUDGET_EXHAUSTED)
This prevents adversarial agents from probing the exact budget threshold and gaming it. The denial appears identical to a real budget exhaustion, making it impossible to distinguish signal from noise.
Every SHACKLE message is wrapped in a Protocol Buffers envelope:
message Envelope {
string protocol_version = 1 ; // "1.0.0"
string message_id = 2 ; // UUIDv7, client-generated
string correlation_id = 3 ; // Request/response pairing
int64 client_timestamp_ns = 4 ; // Client clock (informational)
int64 server_timestamp_ns = 5 ; // Set by daemon on receipt
bytes hmac = 6 ; // HMAC-SHA256 over payload
oneof payload {
PreExecRequest pre_exec = 10 ;
PreExecResponse pre_exec_response = 11 ;
PostExecNotification post_exec = 12 ;
RegisterRequest register = 13 ;
RegisterResponse register_response = 14 ;
Heartbeat heartbeat = 15 ;
HeartbeatAck heartbeat_ack = 16 ;
Error error = 17 ;
}
}
4.2 Session Registration
message RegisterRequest {
string agent_id = 1 ;
string agent_version = 2 ;
string framework = 3 ; // "crewai" | "autogen" | "langgraph"
string session_id = 4 ; // Optional: resume existing session
string organization_id = 5 ;
string runtime = 6 ;
map < string , string > metadata = 7 ;
}
message RegisterResponse {
string session_id = 1 ;
string daemon_version = 2 ;
string negotiated_protocol = 3 ;
GuardConfig active_config = 4 ;
int64 daemon_time_ns = 5 ;
}
4.3 Pre-Execution Check
message PreExecRequest {
string session_id = 1 ;
uint64 call_number = 2 ; // Monotonically increasing
string tool_name = 3 ;
bytes tool_params_hash = 4 ; // SHA-256 of canonical JSON params
double estimated_cost_usd = 5 ;
string parent_guard_id = 6 ; // For nested guard trees
uint64 nonce = 7 ; // Anti-replay
map < string , string > tags = 8 ;
}
message PreExecResponse {
string session_id = 1 ;
uint64 call_number = 2 ;
Verdict verdict = 3 ;
DenyReason deny_reason = 4 ;
string human_readable_reason = 5 ;
double bu

[truncated]
