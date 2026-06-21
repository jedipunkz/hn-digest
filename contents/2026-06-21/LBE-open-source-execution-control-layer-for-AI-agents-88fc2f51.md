---
source: "https://github.com/Letterblack0306/LetterBlack-Sentinel"
hn_url: "https://news.ycombinator.com/item?id=48616412"
title: "LBE – open-source execution control layer for AI agents"
article_title: "GitHub - Letterblack0306/LetterBlack-Sentinel: The execution-control layer between AI agent decisions and real-world actions. · GitHub"
author: "letterblack0306"
captured_at: "2026-06-21T07:44:51Z"
capture_tool: "hn-digest"
hn_id: 48616412
score: 2
comments: 0
posted_at: "2026-06-21T07:11:01Z"
tags:
  - hacker-news
  - translated
---

# LBE – open-source execution control layer for AI agents

- HN: [48616412](https://news.ycombinator.com/item?id=48616412)
- Source: [github.com](https://github.com/Letterblack0306/LetterBlack-Sentinel)
- Score: 2
- Comments: 0
- Posted: 2026-06-21T07:11:01Z

## Translation

タイトル: LBE – AI エージェント用のオープンソース実行制御レイヤー
記事のタイトル: GitHub - Letterblack0306/LetterBlack-Sentinel: AI エージェントの決定と現実世界のアクションの間の実行制御レイヤー。 · GitHub
説明: AI エージェントの決定と現実世界のアクションの間の実行制御層。 - Letterblack0306/LetterBlack-センチネル

記事本文:
GitHub - Letterblack0306/LetterBlack-Sentinel: AI エージェントの決定と現実世界のアクションの間の実行制御レイヤー。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
レターブラック0306
/
レターブラックセンチネル
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
Letterblack0306/LetterBlack-センチネル
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
40 のコミット 40 のコミット
LBE は、AI エージェントが提案する内容とシステムが実際に実行する内容の間にローカル ポリシー ゲートを置きます。ファイル書き込み、シェルコマンドなど、すべてのアクションは実行前にローカルで検証されます。クラウドサービスはありません。デーモンはありません。
本番環境で使用: LBE は、Letterblack for After Effects 内の安全エンジンです。AI によって生成されたすべてのスクリプトと自動化コマンドは、ライブ プロジェクトに触れる前に LBE を通過します。
欲しいです…
パッケージ
ファイル書き込みとシェルコマンドを処理する LBE (フルコントローラー)
@letterblack/lbe-exec
許可/拒否の決定だけ - 私がそれを実行します
@letterblack/lbe-sdk ← ここにいます
インストール
npm インストール @letterblack/lbe-sdk
Node.js ≥ 20.9.0 が必要です。
'@letterblack/lbe-sdk' からインポート {実行} ;
const リクエスト = {
バージョン : '1.0' 、
request_id : 'req-001' 、
タイムスタンプ: 数学。階 (日付.現在()/1000)、
アクター : { id : 'agent:local' 、役割 : 'agent' } 、
インテント: { タイプ: 'コマンド' 、名前: 'write_file' 、ペイロード: {ターゲット: 'out.txt' } }、
コンテキスト:{ワークスペース:プロセス。 cwd() 、環境 : { } 、履歴 : [ ] } 、
制約: {policy_mode : 'strict' 、timeout_ms : 5000 } 、
auth : { 署名 : '<ホスト署名>' 、 nonce : '<リクエストごとに一意>' }
} ;
const 結果 = JSON 。 parse(execute(JSON .stringify(request)) ) ;
// 承認済み: { ok: true、決定: '許可', ... }
// ブロック: { ok: false、決定: '拒否'、エラー: { ステージ、メッセージ } }
execute(input: string): string — JSON を受け入れ、JSON を返します。ランタイム

決定を検証して返します。ホストは決定に基づいて行動します。
フィールド
必須
説明
バージョン
はい
「1.0」
リクエストID
はい
呼び出し元が指定した一意の識別子
タイムスタンプ
はい
Unix タイムスタンプ (秒単位)
俳優
はい
{ id, role } — リクエスト元エージェントの ID
意図
はい
{ type, name, payload } — エージェントが実行したいこと
コンテキスト
はい
ワークスペースのパスと呼び出し元のコンテキスト
制約
はい
policy_mode と timeout_ms
認証
はい
ホストが提供する署名とノンス
オブザーバーモード — ここから始めます
ブロックする準備ができていませんか?オブザーバーモードで起動します。すべてのリクエストは完全に検証され、施行時とまったく同じようにログに記録されますが、何もブロックされません。何を拒否するかを決定する前に、エージェントの行動を観察してください。
npx lbe init # オブザーバーモードで lbe.policy.json を作成する
npx lbe enforce # ブロッキングに切り替える
npx lbe観察 # アドバイザリーに戻す
CLI リファレンス
コマンド
目的
npx lbe 初期化
オブザーバー モードでプロジェクト ローカル ポリシーとキー状態を作成する
npx lbe ポリシーの追加
アクティブなポリシーにルールを追加する
npx lbe観察
アドバイザリー (ログのみ) モードを設定します
npx lbe 強制
ブロックモードを設定する
npx lbe 実行
--in <file> からの提案を検証して実行します
npx lbe 検証
提案を実行せずに検証する
npx lbe ドライラン
実行せずに検証およびシミュレーションを行う
npx lbe健康
必要なファイルがすべて存在し、読み取り可能であることを確認してください
npx lbe 監査検証
監査ログのハッシュ チェーンを確認する
ゲートパイプラインの仕組み
すべてのリクエストは 7 ゲートのパイプラインに入ります。いずれかのゲートで障害が発生すると、構造化された拒否が返されます。残りのゲートは評価されません。
[1] スキーマの必須フィールドと構造的妥当性
↓
[2] タイムスタンプが許容するクロック スキュー ウィンドウ (±10 分)
↓
[3] キーのライフサイクルの信頼できるキー、アクティブ、期限切れではない
↓
[4] 署名 Ed25519 リクエストの信頼性
↓
[5] リクエスタごとのレート制限、スライディング ウィンドウ制限
↓
[6] ノンス・シ

角度使用リプレイ保護
↓
[7] ポリシー設定された認可 (deny-wins)
↓
許可 / 拒否 / エラー — ホストに返される構造化された結果
WASM ランタイムはすべてのゲート決定を所有します。ホストは決定を受け取り、それに基づいて行動します。ランタイム内では何も実行されません。
エージェントは署名されたアクション提案を作成します。
ID はローカルに保持されているキーと照合して確認されます。ネットワーク呼び出しは必要ありません。
プロジェクトの方針が評価されます。アクションは承認されました。
ホストは、許可されたワークスペース内で書き込みまたはコマンドを実行します。
監査チェーンが拡張され、承認されたアクションごとにハッシュリンクされたエントリがローカル ログに追加され、永続的に検証可能になり、サイレントに削除することは不可能になります。
構造化された結果として、成功したかどうか、どのルールが一致したか、および監査エントリの識別子が返されます。
アプリケーションは制御を維持します。 @letterblack/lbe-sdk は、アクションが許可されたかどうかを判断し、回答を返します。それはあなたに代わって実行されるわけではありません。
エージェントは、間違い、設定ミス、または意図的なバイパス試行などのアクションを試みます。
ポリシーゲートはすぐに閉じます。 WASM ランタイムは、アダプターに到達する前に、要求が拒否されたことをスタンプします。
殻は無傷です。ファイルシステムは変更されません。
拒否は不変の監査ログに書き込まれ、チェーンは封印され、証拠は保存されます。
部分的な実行はありません。サイレントな失敗はありません。拒否は第一級の結果であり、エラーではありません。
脅威
ゲート
不正なリクエストまたは不完全なリクエスト
スキーマ
古いリクエストまたはリプレイされたリクエスト
タイムスタンプ + ノンス
改ざんまたは期限切れのキー
鍵のライフサイクル + 署名
一人の俳優からの過剰な要求
レート制限
プロジェクトポリシーで許可されていないアクション
ポリシー — 拒否勝利
エージェントがプロジェクトルート外に書き込む
決定後のホストでのスコープチェック
どの船
dist/index.js WebAssembly ランタイム ローダーとexecute()
dist/cli.js ローカル CLI (npx lbe)
距離/lbe_

Engine.wasm 検証済みランタイムバイナリ
dist/wasm.lock.json ランタイム整合性ロック (wasm バイナリの SHA-256)
assets/lbe-gates.jpg ゲートシーケンス図
assets/story-allow.jpg 承認済みリクエストのストーリーボード
assets/story-deny.jpg ブロックされたリクエストのストーリーボード
assets/runtime-boundary.svg ランタイム境界図
assets/lbe-gates.png ゲート シーケンス図 (フル解像度)
assets/story-allow.png 承認済みリクエストのストーリーボード (フル解像度)
assets/story-deny.png ブロックされたリクエストのストーリーボード (フル解像度)
types.d.ts TypeScript 宣言
ロード時に、ランタイムは lbe_engine.wasm を wasm.lock.json に対して検証します。バイナリが見つからない、変更されている、またはスワップされている場合、リクエストが処理される前に失敗します。
ソース コード、コントローラー実装、アダプター、テスト、キー、および実行時状態は含まれません。
このパッケージは、ランタイムを通じてルーティングされたリクエストを検証します。カーネル レベルのプロセス分離、ネットワーク出力制御、マルチテナント分離、ホスト型コントロール プレーンは提供されません。
ファイル操作、シェル、およびポリシー管理が組み込まれたインプロセス コントローラーについては、 @letterblack/lbe-exec を参照してください。
AI エージェントの決定と現実世界のアクションの間の実行制御層。
github.com/Letterblack0306/LetterBlack-Sentinel
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
5
v1.2.20
最新の
2026 年 6 月 21 日
+ 4 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The execution-control layer between AI agent decisions and real-world actions. - Letterblack0306/LetterBlack-Sentinel

GitHub - Letterblack0306/LetterBlack-Sentinel: The execution-control layer between AI agent decisions and real-world actions. · GitHub
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
Letterblack0306
/
LetterBlack-Sentinel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Letterblack0306/LetterBlack-Sentinel
main Branches Tags Go to file Code Open more actions menu Folders and files
40 Commits 40 Commits assets assets dist dist LICENSE LICENSE README.md README.md package.json package.json types.d.ts types.d.ts View all files Repository files navigation
LBE puts a local policy gate between what an AI agent proposes and what the system actually executes. Every action — file write, shell command, anything — is validated locally before it runs. No cloud service. No daemon.
Used in production: LBE is the safety engine inside Letterblack for After Effects — every AI-generated script and automation command passes through it before touching a live project.
I want…
Package
LBE to handle file writes and shell commands for me (full controller)
@letterblack/lbe-exec
Just the allow/deny decision — I'll execute it myself
@letterblack/lbe-sdk ← you are here
Install
npm install @letterblack/lbe-sdk
Requires Node.js ≥ 20.9.0.
import { execute } from '@letterblack/lbe-sdk' ;
const request = {
version : '1.0' ,
request_id : 'req-001' ,
timestamp : Math . floor ( Date . now ( ) / 1000 ) ,
actor : { id : 'agent:local' , role : 'agent' } ,
intent : { type : 'command' , name : 'write_file' , payload : { target : 'out.txt' } } ,
context : { workspace : process . cwd ( ) , env : { } , history : [ ] } ,
constraints : { policy_mode : 'strict' , timeout_ms : 5000 } ,
auth : { signature : '<host-signed>' , nonce : '<unique-per-request>' }
} ;
const result = JSON . parse ( execute ( JSON . stringify ( request ) ) ) ;
// Approved: { ok: true, decision: 'allow', ... }
// Blocked: { ok: false, decision: 'deny', error: { stage, message } }
execute(input: string): string — accepts JSON, returns JSON. The runtime validates and returns a decision. The host acts on the decision.
Field
Required
Description
version
Yes
"1.0"
request_id
Yes
Caller-supplied unique identifier
timestamp
Yes
Unix timestamp in seconds
actor
Yes
{ id, role } — identity of the requesting agent
intent
Yes
{ type, name, payload } — what the agent wants to do
context
Yes
Workspace path and caller context
constraints
Yes
policy_mode and timeout_ms
auth
Yes
Host-supplied signature and nonce
Observer mode — start here
Not ready to block? Start in observer mode. Every request is fully validated and logged exactly as it would be in enforcement — but nothing is blocked. Watch what the agent is doing before you decide what to deny.
npx lbe init # create lbe.policy.json in observer mode
npx lbe enforce # switch to blocking
npx lbe observe # switch back to advisory
CLI reference
Command
Purpose
npx lbe init
Create project-local policy and key state in observer mode
npx lbe policy-add
Add a rule to the active policy
npx lbe observe
Set advisory (log-only) mode
npx lbe enforce
Set blocking mode
npx lbe run
Validate and execute a proposal from --in <file>
npx lbe verify
Validate a proposal without executing
npx lbe dryrun
Validate and simulate without executing
npx lbe health
Check all required files are present and readable
npx lbe audit-verify
Verify the audit log hash chain
How the gate pipeline works
Every request enters a 7-gate pipeline. A failure at any gate returns a structured denial — the remaining gates are not evaluated.
[1] Schema required fields and structural validity
↓
[2] Timestamp permitted clock-skew window (±10 minutes)
↓
[3] Key lifecycle trusted key, active, not expired
↓
[4] Signature Ed25519 request authenticity
↓
[5] Rate limit per-requester sliding-window limit
↓
[6] Nonce single-use replay protection
↓
[7] Policy configured authorization (deny-wins)
↓
allow / deny / error — structured result returned to host
The WASM runtime owns all gate decisions. Your host receives the decision and acts on it. Nothing executes inside the runtime.
The agent produces a signed action proposal.
Identity is confirmed against a locally held key — no network call required.
The project policy is evaluated. The action is approved.
The host executes the write or command inside the allowed workspace.
The audit chain is extended — every approved action appends a hash-linked entry to the local log, permanently verifiable, impossible to silently remove.
A structured result returns: whether it succeeded, which rules matched, and the audit entry identifier.
The application stays in control. @letterblack/lbe-sdk decides whether the action was permitted and hands the answer back. It does not execute for you.
The agent attempts an action — whether by mistake, misconfiguration, or a deliberate bypass attempt.
The policy gate closes immediately. The WASM runtime stamps the request denied before any adapter is reached.
The shell is untouched. The filesystem is unchanged.
The denial is written to the immutable audit log — chain sealed, evidence preserved.
No partial execution. No silent failures. Denial is a first-class outcome, not an error.
Threat
Gate
Malformed or incomplete request
Schema
Stale or replayed request
Timestamp + Nonce
Tampered or expired key
Key lifecycle + Signature
Excessive requests from one actor
Rate limit
Action not permitted by project policy
Policy — deny-wins
Agent writing outside project root
Scope check in host after decision
What ships
dist/index.js WebAssembly runtime loader and execute()
dist/cli.js Local CLI (npx lbe)
dist/lbe_engine.wasm Verified runtime binary
dist/wasm.lock.json Runtime integrity lock (SHA-256 of wasm binary)
assets/lbe-gates.jpg Gate sequence diagram
assets/story-allow.jpg Approved-request storyboard
assets/story-deny.jpg Blocked-request storyboard
assets/runtime-boundary.svg Runtime boundary diagram
assets/lbe-gates.png Gate sequence diagram (full resolution)
assets/story-allow.png Approved-request storyboard (full resolution)
assets/story-deny.png Blocked-request storyboard (full resolution)
types.d.ts TypeScript declarations
At load time the runtime verifies lbe_engine.wasm against wasm.lock.json . A missing, modified, or swapped binary fails before any request is processed.
Source code, controller implementation, adapters, tests, keys, and runtime state are not included.
This package validates requests routed through its runtime. It does not provide kernel-level process isolation, network-egress control, multi-tenant separation, or a hosted control plane.
For an in-process controller with file operations, shell, and policy management built in, see @letterblack/lbe-exec .
The execution-control layer between AI agent decisions and real-world actions.
github.com/Letterblack0306/LetterBlack-Sentinel
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
5
v1.2.20
Latest
Jun 21, 2026
+ 4 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
