---
source: "https://github.com/bossandboss/Ai-Firewall"
hn_url: "https://news.ycombinator.com/item?id=49018165"
title: "AI Firewall – Security Gateway and Reverse Proxy for LLM Traffic"
article_title: "GitHub - bossandboss/Ai-Firewall: Security gateway & reverse proxy for LLM traffic — real-time DLP redaction, prompt-injection filtering and EU AI Act audit logging. Go edge proxy (Aho-Corasick, SSE stream interception) + TypeScript management plane. Source-available (BSL 1.1). · GitHub"
author: "bossandboss"
captured_at: "2026-07-23T08:12:07Z"
capture_tool: "hn-digest"
hn_id: 49018165
score: 1
comments: 0
posted_at: "2026-07-23T07:47:13Z"
tags:
  - hacker-news
  - translated
---

# AI Firewall – Security Gateway and Reverse Proxy for LLM Traffic

- HN: [49018165](https://news.ycombinator.com/item?id=49018165)
- Source: [github.com](https://github.com/bossandboss/Ai-Firewall)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T07:47:13Z

## Translation

タイトル: AI ファイアウォール – LLM トラフィック用のセキュリティ ゲートウェイとリバース プロキシ
記事のタイトル: GitHub - Bossandboss/Ai-Firewall: LLM トラフィック用のセキュリティ ゲートウェイとリバース プロキシ — リアルタイム DLP リダクション、プロンプト インジェクション フィルタリング、および EU AI 法の監査ログ。 Go エッジ プロキシ (Aho-Corasick、SSE ストリーム インターセプト) + TypeScript 管理プレーン。ソースが利用可能 (BSL 1.1)。 · GitHub
説明: LLM トラフィック用のセキュリティ ゲートウェイとリバース プロキシ — リアルタイム DLP リダクション、プロンプト インジェクション フィルタリング、および EU AI Act 監査ログ。 Go エッジ プロキシ (Aho-Corasick、SSE ストリーム インターセプト) + TypeScript 管理プレーン。ソースが利用可能 (BSL 1.1)。 - ボスサンドボス/Ai-ファイアウォール

記事本文:
GitHub - Bossandboss/Ai-Firewall: LLM トラフィック用のセキュリティ ゲートウェイとリバース プロキシ — リアルタイム DLP リダクション、プロンプト インジェクション フィルタリング、EU AI 法の監査ログ。 Go エッジ プロキシ (Aho-Corasick、SSE ストリーム インターセプト) + TypeScript 管理プレーン。ソースが利用可能 (BSL 1.1)。 · GitHub
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
アカウントを切り替えました

別のタブまたはウィンドウ上で。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
上司と上司
/
Ai-ファイアウォール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット cmd/ プロキシ cmd/ プロキシ 内部 内部 pkg/ スキャナー pkg/ スキャナー src src ライセンス ライセンス README.md README.md go.mod go.mod Index.html Index.html package-lock.json package-lock.json package.json package.json server.ts server.ts tsconfig.json tsconfig.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🛡️ AI ファイアウォール — LLM トラフィック用のセキュリティ ゲートウェイとリバース プロキシ
AI ファイアウォールは、アプリケーションと LLM の間に位置するセキュリティ ゲートウェイです。
API (OpenAI、Anthropic、Gemini、Ollama、vLLM)。プロンプトを事前に検査します
ユーザーに届く前に境界とモデルの出力を残しておきます。
DLP / シークレット編集 — AWS キー、API トークン、JWT、秘密キー ヘッダー、
SSN、クレジット カード、電子メール、電話番号はマスクされるかブロックされます。
プロンプトインジェクションと脱獄フィルタリング — ペルソナオーバーライドフレーズ、システム
プロンプト抽出試行、マークダウン抽出ピクセル、ゼロ幅
ステガノグラフィー。
EU AI 法のタグ付け — すべてのリクエストは分類されます (第 5 条の禁止 /
付属書 III 高リスク / Art. 50 透明性 / 最小限のリスク）、
永続的な監査ログ。
ストリーミングの傍受 - Go エッジ プロキシが SSE のシークレットを編集します
チャンク境界を越えて分割されたシークレットを含むストリーム (ユニットによって検証される)
テスト）。
このリポジトリに存在するものとそうでないもの
これは正直なMVPです。以下の表は真実です。他には何もない
と主張した。
既知の制限、明確に述べると、現在の検出層は決定的です
(キーワード自動

aton + 正規表現)。既知のパターンでは高速かつ正確です。
構造的に言い換え攻撃や多言語攻撃に対して弱い。ベンチマーク
画面はこの隙間を意図的に見えるようにします。 ONNX 分類子でそれを閉じる
(例: onnxruntime-node を介した微調整された DeBERTa-v3 プロンプト インジェクション モデル)
はロードマップの最上位項目です。
npmインストール
cp .env.example .env # プロバイダーキーを追加するか、FIREWALL_DEMO_MODE=true のままにしておきます
npm run dev # http://localhost:3000
最初の起動時に、管理キーが生成され、一度印刷されます (
SHA-256 ハッシュは data/api-keys.json に保存されます)。ダッシュボードに貼り付けます
Key Manager を使用するか、 .env で FIREWALL_ADMIN_KEY を設定します。ルールとキーの変更
それを要求する。プロバイダーの資格情報がないと、プロキシは明確なフラグを立てて返します。
デモ モードではシミュレートされた完了 ( "simulated": true )、それ以外の場合は HTTP 502。
go test ./... # 単体テストを含む極秘ストリーム編集
./cmd/proxy をビルドします
UPSTREAM_LLM_URL= " https://api.openai.com " \
FIREWALL_CLIENT_KEYS= " クライアント キー " \
PORT=8080 ./プロキシ
API
POST /api/v1/proxy/chat/completions
{
"プロバイダー" : " openai " ,
"モデル" : " gpt-4o-mini " 、
"messages" : [{ "role" : " user " , "content" : " 私の AWS キーは AKIAIOSFODNN7EXAMPLE です。収益を要約してください。 " }]
}
プロンプトは転送前にサニタイズされます。応答にはファイアウォールが含まれます
ブロック (プロキシ レイテンシ、編集カウント、EU AI 法のカテゴリ、監査ログ ID) および
プロバイダーの資格情報が設定されていない場合にのみ true になるシミュレートされたフラグ。
その他のエンドポイント: POST /api/v1/inspect (ドライラン)、GET/POST/PUT/DELETE /api/v1/rules (変異: admin)、GET /api/v1/logs 、GET /api/v1/metrics
(実際のトラフィックのみから計算)、GET/POST/DELETE /api/v1/auth/keys (管理者)、
POST /api/v1/testsuite/run 、 GET /api/v1/proxy/sse/stream (ラベル付きデモ)
ビジュアライザーのストリーム)。
[ クライアント ] ──> [ ゴーエッジプロキシ :808

0 ] [ ノードゲートウェイ + ダッシュボード :3000 ]
§─ SHA-256 キー認証 §─ マルチプロバイダー転送
§─ Aho-Corasick キーワード オートマトン §─ 検査 + 監査永続化
§─ RE2 シークレット編集 §─ ハッシュ化された API キーの管理
━─ SSE キャリー ウィンドウ インターセプター ─ ライブ測定ベンチマーク
│
[ アップストリーム LLM API ]
2 つのプロセスは独立しています。Go プロキシは低遅延のデータパスです。
コンポーネント;ノード ゲートウェイは管理プレーンおよびデモ サーフェスです。配線
HTTP 経由の Go プロキシへのダッシュボードのルール ストアはロードマップにあります。
第 2 の検出段階としての ONNX トランス分類器、ベンチマーク
再現可能なスクリプトを含むパブリック プロンプト インジェクション データセット。
ノード管理プレーンと Go プロキシ間のルールの同期。
公開されている p50/p99 オーバーヘッドを備えた再現可能な負荷テスト ハーネス ( k6 )。
PostgreSQL 監査ストレージ、保持ポリシー、SIEM (syslog/OTLP) エクスポート。
API キーごとのレート制限。プロキシと管理プレーン間の mTLS。
ビジネス ソース ライセンス 1.1 (BUSL-1.1) — LICENSE ファイルを参照してください。
つまり、ソースが利用可能です。非実稼働用途および実稼働用途で無料
1 日あたり最大 10,000 件のリクエストを検査します。競合他社としてソフトウェアを提供する
ホスト/マネージド サービスには商用ライセンスが必要です。変更日について
(2030-07-23) コードは自動的に Apache 2.0 に変換されます。
公開前: ライセンスにライセンサーの名前と連絡先メールアドレスを入力します。
パラメーターを確認し、知財弁護士による最終文書のレビューを受けてください。
LLM トラフィック用のセキュリティ ゲートウェイとリバース プロキシ - リアルタイム DLP リダクション、プロンプト インジェクション フィルタリング、EU AI 法の監査ログ。 Go エッジ プロキシ (Aho-Corasick、SSE ストリーム インターセプト) + TypeScript 管理プレーン。ソースが利用可能 (BSL 1.1)。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
レル

緩和する
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Security gateway & reverse proxy for LLM traffic — real-time DLP redaction, prompt-injection filtering and EU AI Act audit logging. Go edge proxy (Aho-Corasick, SSE stream interception) + TypeScript management plane. Source-available (BSL 1.1). - bossandboss/Ai-Firewall

GitHub - bossandboss/Ai-Firewall: Security gateway & reverse proxy for LLM traffic — real-time DLP redaction, prompt-injection filtering and EU AI Act audit logging. Go edge proxy (Aho-Corasick, SSE stream interception) + TypeScript management plane. Source-available (BSL 1.1). · GitHub
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
bossandboss
/
Ai-Firewall
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit cmd/ proxy cmd/ proxy internal internal pkg/ scanner pkg/ scanner src src LICENSE LICENSE README.md README.md go.mod go.mod index.html index.html package-lock.json package-lock.json package.json package.json server.ts server.ts tsconfig.json tsconfig.json vite.config.ts vite.config.ts View all files Repository files navigation
🛡️ AI Firewall — Security Gateway & Reverse Proxy for LLM Traffic
AI Firewall is a security gateway that sits between your applications and LLM
APIs (OpenAI, Anthropic, Gemini, Ollama, vLLM). It inspects prompts before they
leave your perimeter and model outputs before they reach your users:
DLP / secret redaction — AWS keys, API tokens, JWTs, private-key headers,
SSNs, credit cards, emails, phone numbers are masked or blocked.
Prompt-injection & jailbreak filtering — persona-override phrases, system
prompt exfiltration attempts, markdown-exfiltration pixels, zero-width
steganography.
EU AI Act tagging — every request is categorized (Art. 5 prohibited /
Annex III high-risk / Art. 50 transparency / minimal risk) and written to a
persistent audit log.
Streaming interception — the Go edge proxy redacts secrets in SSE
streams, including secrets split across chunk boundaries (verified by unit
test).
What is real in this repository — and what is not
This is an honest MVP. The table below is the ground truth; nothing else is
claimed.
Known limitation, stated plainly: the current detection layer is deterministic
(keyword automaton + regex). It is fast and precise on known patterns, and
structurally weak against paraphrased or multilingual attacks. The benchmark
screen makes this gap visible on purpose; closing it with an ONNX classifier
(e.g. a fine-tuned DeBERTa-v3 prompt-injection model via onnxruntime-node )
is the top roadmap item.
npm install
cp .env.example .env # add provider keys, or leave FIREWALL_DEMO_MODE=true
npm run dev # http://localhost:3000
On first startup an admin key is generated and printed once (only its
SHA-256 hash is stored in data/api-keys.json ). Paste it in the dashboard's
Key Manager, or set FIREWALL_ADMIN_KEY in .env . Rule and key mutations
require it; without provider credentials the proxy returns clearly-flagged
simulated completions ( "simulated": true ) in demo mode, or HTTP 502 otherwise.
go test ./... # unit tests incl. split-secret stream redaction
go build ./cmd/proxy
UPSTREAM_LLM_URL= " https://api.openai.com " \
FIREWALL_CLIENT_KEYS= " your-client-key " \
PORT=8080 ./proxy
API
POST /api/v1/proxy/chat/completions
{
"provider" : " openai " ,
"model" : " gpt-4o-mini " ,
"messages" : [{ "role" : " user " , "content" : " My AWS key is AKIAIOSFODNN7EXAMPLE. Summarize earnings. " }]
}
The prompt is sanitized before forwarding; the response includes a firewall
block (proxy latency, redaction count, EU AI Act category, audit log id) and a
simulated flag that is true only when no provider credentials are set.
Other endpoints: POST /api/v1/inspect (dry-run), GET/POST/PUT/DELETE /api/v1/rules (mutations: admin), GET /api/v1/logs , GET /api/v1/metrics
(computed from real traffic only), GET/POST/DELETE /api/v1/auth/keys (admin),
POST /api/v1/testsuite/run , GET /api/v1/proxy/sse/stream (labeled demo
stream for the visualizer).
[ CLIENT ] ──> [ Go edge proxy :8080 ] [ Node gateway + dashboard :3000 ]
├─ SHA-256 key auth ├─ multi-provider forwarding
├─ Aho-Corasick keyword automaton ├─ inspection + audit persistence
├─ RE2 secret redaction ├─ hashed API-key management
└─ SSE carry-window interceptor └─ live measured benchmarks
│
[ UPSTREAM LLM API ]
The two processes are independent: the Go proxy is the low-latency data-path
component; the Node gateway is the management plane and demo surface. Wiring
the dashboard's rule store to the Go proxy over HTTP is on the roadmap.
ONNX transformer classifier as a second detection stage, benchmarked on
public prompt-injection datasets with reproducible scripts.
Rule synchronization between the Node management plane and the Go proxy.
Reproducible load-test harness ( k6 ) with published p50/p99 overhead.
PostgreSQL audit storage, retention policies, SIEM (syslog/OTLP) export.
Rate limiting per API key; mTLS between proxy and management plane.
Business Source License 1.1 (BUSL-1.1) — see the LICENSE file.
In short: source-available; free for non-production use and for production use
up to 10,000 inspected requests per day; offering the software as a competing
hosted/managed service requires a commercial license. On the Change Date
(2030-07-23) the code automatically converts to Apache 2.0.
Before publishing: fill in the Licensor name and contact email in the LICENSE
parameters, and have the final text reviewed by an IP lawyer.
Security gateway & reverse proxy for LLM traffic — real-time DLP redaction, prompt-injection filtering and EU AI Act audit logging. Go edge proxy (Aho-Corasick, SSE stream interception) + TypeScript management plane. Source-available (BSL 1.1).
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
