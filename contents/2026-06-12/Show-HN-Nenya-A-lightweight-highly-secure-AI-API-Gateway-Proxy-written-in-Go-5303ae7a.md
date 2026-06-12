---
source: "https://github.com/gumieri/nenya"
hn_url: "https://news.ycombinator.com/item?id=48506994"
title: "Show HN: Nenya – A lightweight, highly secure AI API Gateway/Proxy written in Go"
article_title: "GitHub - gumieri/nenya: A lightweight, highly secure AI API Gateway/Proxy written in Go. Acts as transparent middleware between local AI coding clients (OpenCode/Pi/Cursor) and upstream LLM providers (Gemini, DeepSeek, Zhipu z.ai). · GitHub"
author: "garou"
captured_at: "2026-06-12T18:46:09Z"
capture_tool: "hn-digest"
hn_id: 48506994
score: 1
comments: 0
posted_at: "2026-06-12T17:32:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Nenya – A lightweight, highly secure AI API Gateway/Proxy written in Go

- HN: [48506994](https://news.ycombinator.com/item?id=48506994)
- Source: [github.com](https://github.com/gumieri/nenya)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T17:32:40Z

## Translation

タイトル: Show HN: Nenya – Go で書かれた軽量で安全性の高い AI API ゲートウェイ/プロキシ
記事タイトル: GitHub - umieri/nenya: Go で書かれた軽量で安全性の高い AI API ゲートウェイ/プロキシ。ローカル AI コーディング クライアント (OpenCode/Pi/Cursor) と上流の LLM プロバイダー (Gemini、DeepSeek、Zhipu z.ai) の間の透過的なミドルウェアとして機能します。 · GitHub
説明: Go で書かれた軽量で安全性の高い AI API ゲートウェイ/プロキシ。ローカル AI コーディング クライアント (OpenCode/Pi/Cursor) と上流の LLM プロバイダー (Gemini、DeepSeek、Zhipu z.ai) の間の透過的なミドルウェアとして機能します。 - グミエリ/ネニャ

記事本文:
GitHub - umieri/nenya: Go で書かれた軽量で安全性の高い AI API ゲートウェイ/プロキシ。ローカル AI コーディング クライアント (OpenCode/Pi/Cursor) と上流の LLM プロバイダー (Gemini、DeepSeek、Zhipu z.ai) の間の透過的なミドルウェアとして機能します。 · GitHub
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

グミエリ
/
ネニャ
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
396 コミット 396 コミット .github .github .opencode/ プラン .opencode/ プラン cmd cmd config config デプロイ デプロイ ドキュメント ドキュメントの例 例 内部 内部パッケージング/スクリプト パッケージ化/スクリプト .containerignore .containerignore .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh missise.toml misse.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Go で書かれた軽量で依存性のない AI API ゲートウェイ。 Nenya は、AI コーディング クライアントと上流の LLM プロバイダーの間に位置し、シークレットの編集、コンテキスト管理、エージェント ルーティング、MCP ツールの統合を追加します。すべて透過的な SSE ストリーミングを使用します。セキュリティ強化: 非 root 実行、シークレットの mlock、seccomp + no-new-privileges。
OpenAI または Anthropic Chat Completions API を実装する任意のプロバイダーと互換性があります。 23 のプロバイダー向けに、特殊な処理を備えた組み込みアダプターを出荷します。
Nenya がリクエストを処理する方法
+----------------------------------------------+
|クライアント (Cursor / OpenCode / Aider / など) |
| OpenAI互換リクエスト |
| POST /v1/chat/completions + Bearer トークン |
|または |
| Anthropic メッセージ API リクエスト |
| POST /v1/messages + x-api-key |
+----------------------------------------------+
|
v
+----------------------------------------------+
|ネンヤ ゲートウェイ |
| - 認証チェック + RBAC 強制 |
| - JSON の解析 + モデルの抽出 |
| - エージェント/プロバイダーを解決する |
| - オプションのキャッシュ (HIT =>

SSE をリプレイ) |
| - オプションの MCP コンテキスト/ツール インジェクション |
+----------------------------------------------+
|
v
+----------------------------------------------+
|インターセプター チェーン (プラグ可能、ベストエフォート) |
| - RedactInterceptor (正規表現パターン) |
| - EntropyInterceptor (高エントロピー文字列) |
| - TFIDFInterceptor (関連性スコアリング) |
| - BouncerInterceptor (エンジンの概要) |
+----------------------------------------------+
|
v
+----------------------------------------------+
|トークンバジェットのトリミング (ペイロード > ハードの場合 |
|制限) は最も古い非システム メッセージを削除し、 |
|トークン対応のミドルアウト切り捨てを適用します。
+----------------------------------------------+
|
v
+----------------------------------------------+
|ルーティング |
| A) 標準転送 |
| - フォールバック チェーン + サーキット ブレーカー + RL |
| B) MCP マルチターンツールループ (有効な場合) |
| - SSE をバッファリングし、MCP ツールを実行し、再送信します。
| C) コンテキスト制限の再試行 |
| - アップストリーム 413/context_exceeded、 |
|ペイロードを要約し、フォールバックで再試行します。
+----------------------------------------------+
|
v
+----------------------------------------------+
|アップストリーム LLM プロバイダー |
|人類 |ジェミニ |ディープシーク |ミストラル | ...|
+----------------------------------------------+
|
| SSEストリーム
v
+----------------------------------------------+
| Nenya SSE パイプライン |
| - アダプターの応答変換 |
| - (オプション) OpenAI→人間性変換 |
| - 使用量アカウンティング + ストリーム フィルター |
| - フラッシュ + (オプション) キャッシュ キャプチャ |
| - (オプション) MCP 自動
[切り捨てられた]
/v1/* エンドポイントにはクライアント ベアラー認証が必要です。 /healthz 、 /statsz 、 /metrics は行いません。
パイプライン障害は正常に機能を低下させ、500 を返す代わりにリクエストを転送します。
MCP 対応エージェントは、MCP の複雑さをクライアントに公開することなく、ローカル/リモート ツールを実行できます。
構成ドライブ

プロバイダー レジストリ — JSON 経由でプロバイダーを追加、コード変更なし
ワイヤー形式の違いに特化したアダプターを備えた 23 個の組み込みプロバイダー
動的モデル検出 - 起動時とリロード時にプロバイダーからライブ モデル カタログを取得します。
モデル レジストリ — プロバイダー/コンテキストの自動解決を備えた文字列短縮表現によるモデルの参照
マルチプロバイダー モデルの解決 — モデルが複数のプロバイダーに存在する場合、すべてがエージェントのフォールバック チェーンに追加されます。
3 層のモデル解決 — 構成オーバーライド > 検出されたモデル > 静的レジストリ
モデルごとのワイヤ フォーマット — マルチフォーマット ゲートウェイ (OpenCode Zen) からのモデルは、モデルのフォーマット属性に基づいて OpenAI、Anthropic、および Gemini ワイヤ フォーマットの間で自動変換されます。
エージェント フォールバック チェーン — ラウンドロビンまたはサーキット ブレーカーと自動フェイルオーバーによるシーケンシャル
レイテンシーを意識したルーティング - ±5% のジッターで過去の応答時間の中央値に基づいてターゲットを自動再順序付けし、雷の群れを防ぎます。
エージェントごとのシステム プロンプト - インラインまたはファイルベース
Tier-0 正規表現シークレットフィルター — AWS キー、GitHub トークン、パスワードなどの常時編集。
3 層コンテンツ パイプライン — プラグイン可能なインターセプター チェーン: 正規表現編集、エントロピー フィルタリング、TF-IDF 関連性スコアリング、エンジン要約
コンテキスト ウィンドウの圧縮 — 構成可能なエンジンによるスライディング ウィンドウの要約
古いツール呼び出しのプルーニング — トークンを節約するために古いアシスタントとツール応答のペアを圧縮します
思考の枝刈り — アシスタントのメッセージ履歴から推論ブロックを削除します
入力検証 - 厳格な本文制限、JSON サニタイズ、ヘッダー フィルタリング
グレースフル デグラデーション — エンジンまたはパイプラインの障害によってリクエストがブロックされることはありません
ロールベースのアクセス制御 (RBAC) — エージェントおよびエンドポイント制限のある API ごとの主要なロール (管理者、ユーザー、読み取り専用)
安全なメモリ — mlock で保護されたトークン ストレージ、

読み取り専用シーリング、コアダンプ防止
強化 (展開セキュリティ)
セキュア メモリ (デフォルト) : すべてのトークンは mlock で保護された RAM に保存され、init 後に読み取り専用でシールされ、コア ダンプは無効になります。
非ルート実行 - 機能が削除された状態で UID 65532 として実行されます。
メモリ保護 - systemd の LimitMEMLOCK=infinity および LimitCORE=0
読み取り専用ファイルシステム — 不変ルート + プライベート /tmp
Seccomp + no-new-privileges — 制限されたシステムコール、権限昇格の防止
ゼロトラスト シークレット — systemd 資格情報またはコンテナ マウント経由でロードされ、ディスクには決してロードされません
ソケットのアクティベーション — 切断された接続をゼロでシームレスに再起動します
外部依存関係なし - Go 標準ライブラリのみ
ホットリロード — systemctl reload nenya によるダウンタイムなしの設定変更
サーキット ブレーカー — 自動フェイルオーバー、指数バックオフ、セマンティック エラー分類を備えたエージェント、プロバイダー、モデルごと
レート制限 - プロバイダーごとのオーバーライドを伴うアップストリーム ホストごと (RPM/TPM)
応答キャッシュ — SHA-256 フィンガープリントとオプションのセマンティック類似性検索を備えたメモリ内 LRU
正常なシャットダウン - 処理中のリクエストに対する 5 秒の猶予期間、MCP クライアントのクリーンアップ
コンテキスト制限自動再試行 - アップストリームのコンテキスト長エラーが要約と再試行をトリガーします
ローカル エンジンのライフサイクル — LRU エビクションによるローカル Ollama モデルの事前ロードと管理
構造化エラー — すべてのエラー応答には、プログラムによる診断のための error_kind フィールドが含まれています
ツール検出 - 自動ツールインジェクションのために MCP サーバーに接続します
マルチターン実行 - ツール呼び出しのインターセプト、MCP サーバーに対する実行、結果の転送
自動検索 — 転送する前に MCP サーバーから関連するコンテキストをプリフェッチします
自動保存 — アシスタントの応答を MCP メモリ サーバーに保持します
最小限の構成とシークレットを作成します。
mkdir -p 構成シークレット
cat > config/config.json << ' E

の
{
"サーバー": { "listen_addr": ":8080" },
「エージェント」: {
"デフォルト": {
"戦略": "フォールバック",
"モデル": ["gemini-2.5-フラッシュ"]
}
}
}
終了後
cat > Secrets/provider_keys.json << ' EOF '
{
"プロバイダーキー": {
「ジェミニ」：「アイザ……」
}
}
終了後
cat > Secrets/client.json << ' EOF '
{
"client_token": "nk-$(openssl rand -hex 32)"
}
終了後
コンテナを実行します。
ポッドマン実行 -d \
--name ねにゃ \
-p 8080:8080 \
-v ./config:/etc/nenya:ro \
-v ./secrets:/run/secrets/nenya:ro \
-e NENYA_SECRETS_DIR=/run/secrets/nenya \
--cap-drop=ALL \
--cap-add=IPC_LOCK \
--security-opt=no-new-privileges:true \
--読み取り専用 \
--tmpfs /tmp:rw,noexec,nosuid,size=64M \
ghcr.io/guieri/nenya:最新
テストしてみましょう:
curl -H " 認可: Bearer $( jq -r ' .client_token ' Secrets/client.json ) " \
http://localhost:8080/healthz
またはパッケージマネージャー経由でインストールする
Nenya は、主要な Linux ディストリビューションとコミュニティ パッケージ マネージャーにネイティブ パッケージを提供します。
すべてのパッケージはバイナリを /usr/bin/nenya にインストールし、systemd サービスとソケット ユニットを含みます。インストール後、有効にして開始します。
sudo systemctl Enable --now nenya.socket
sudo systemctl Enable --now nenya.service
ランタイム構成
Nenya は、展開の移植性を高めるための標準環境変数をサポートしています。
PORT=9090 HOST=127.0.0.1 ./nenya --config /path/to/config.json
または Docker で次のようにします。
docker run -e PORT=9090 -p 9090:9090 ghcr.io/umieri/nenya:latest
または、展開を選択してください
ベア メタル (systemd) のデプロイ — 直接バイナリ インストール、ソケット アクティベーション、ホット リロード
コンテナのデプロイ (Podman/Docker Compose) — compose.yml、イメージ検証、セキュリティ強化
Kubernetes のデプロイ (Helm) — Helm チャート、ConfigMap/Secret、イングレス設定
すべての /v1/* エンドポイントには認証が必要です: Bearer <client_token> または Bearer <api_key_token> 。
API キーは RBAC 強制をサポートします - エージェントのスコープ設定、エンドポイントの許可リスト、

ole ベースのアクセス許可 (管理者はすべてのチェックをバイパスします)。
パススルー プロキシの使用法の詳細については、docs/PASSTHROUGH_PROXY.md を参照してください。
文書
説明
プロバイダー
23 個すべてのプロバイダー、機能マトリックス、特殊な動作、カスタム プロバイダーの追加
構成
完全な構成リファレンス、ディレクトリ モード、すべてのセクションとフィールド
ベアメタルのデプロイ
Systemd ユニット、config.d レイアウト、シークレット、ホット リロード
コンテナのデプロイ
Podman/Docker Compose、イメージ検証、セキュリティに関する注意事項
Kubernetesのデプロイ
Helm チャートの使用法、ConfigMap/Secret、イングレス設定
パススループロキシ
Raw プロバイダー エンドポイント プロキシ、SSE ストリーミング、認証インジェクション
建築
パッケージ DAG、リクエスト ライフサイクル、サーキット ブレーカー、SSE パイプライン
MCPの統合
MCP サーバー統合、ツール検出、マルチターン実行
アダプター
アダプターのシステム内部、認証スタイル、機能フラグ
シークレットの形式
Systemd 認証情報、環境変数フォールバック、コンテナ/K8s デプロイメント
セキュリティ
脆弱性報告ポリシー
ライセンス
Go で書かれた軽量で安全性の高い AI API ゲートウェイ/プロキシ。ローカル AI コーディング クライアント (OpenCode/Pi/Cursor) と上流の LLM プロバイダー (Gemini、DeepSeek、Zhipu z.ai) の間の透過的なミドルウェアとして機能します。
Apache-2.0ライセンス
セキュリティポリシー
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
10
ネンヤ 0.6.1
遅刻

[切り捨てられた]

## Original Extract

A lightweight, highly secure AI API Gateway/Proxy written in Go. Acts as transparent middleware between local AI coding clients (OpenCode/Pi/Cursor) and upstream LLM providers (Gemini, DeepSeek, Zhipu z.ai). - gumieri/nenya

GitHub - gumieri/nenya: A lightweight, highly secure AI API Gateway/Proxy written in Go. Acts as transparent middleware between local AI coding clients (OpenCode/Pi/Cursor) and upstream LLM providers (Gemini, DeepSeek, Zhipu z.ai). · GitHub
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
gumieri
/
nenya
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
396 Commits 396 Commits .github .github .opencode/ plans .opencode/ plans cmd cmd config config deploy deploy docs docs examples examples internal internal packaging/ scripts packaging/ scripts .containerignore .containerignore .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh mise.toml mise.toml View all files Repository files navigation
A lightweight, zero-dependency AI API Gateway written in Go. Nenya sits between your AI coding clients and upstream LLM providers, adding secret redaction, context management, agent routing, and MCP tool integration — all with transparent SSE streaming. Security-hardened: non-root execution, mlock for secrets, seccomp + no-new-privileges.
Compatible with any provider that implements the OpenAI Or Anthropic Chat Completions API. For 23 providers we ship built-in adapters with specialized handling.
How Nenya handles the requests
+----------------------------------------------+
| Client (Cursor / OpenCode / Aider / etc.) |
| OpenAI-compatible request |
| POST /v1/chat/completions + Bearer token |
| or |
| Anthropic Messages API request |
| POST /v1/messages + x-api-key |
+----------------------------------------------+
|
v
+----------------------------------------------+
| Nenya Gateway |
| - auth check + RBAC enforcement |
| - parse JSON + extract model |
| - resolve agent/provider |
| - optional cache (HIT => replay SSE) |
| - optional MCP context/tool injection |
+----------------------------------------------+
|
v
+----------------------------------------------+
| Interceptor Chain (pluggable, best-effort) |
| - RedactInterceptor (regex patterns) |
| - EntropyInterceptor (high-entropy strings) |
| - TFIDFInterceptor (relevance scoring) |
| - BouncerInterceptor (engine summarization) |
+----------------------------------------------+
|
v
+----------------------------------------------+
| Token Budget Trimming (if payload > hard |
| limit) drops oldest non-system messages and |
| applies token-aware middle-out truncation |
+----------------------------------------------+
|
v
+----------------------------------------------+
| Routing |
| A) Standard forwarding |
| - fallback chain + circuit breaker + RL |
| B) MCP multi-turn tool loop (if enabled) |
| - buffer SSE, execute MCP tools, re-send |
| C) Context-limit retry |
| - on upstream 413/context_exceeded, |
| summarize payload, retry with fallback |
+----------------------------------------------+
|
v
+----------------------------------------------+
| Upstream LLM Providers |
| Anthropic | Gemini | DeepSeek | Mistral | ...|
+----------------------------------------------+
|
| SSE stream
v
+----------------------------------------------+
| Nenya SSE Pipeline |
| - adapter response transforms |
| - (optional) OpenAI→Anthropic conversion |
| - usage accounting + stream filter |
| - flush + (optional) cache capture |
| - (optional) MCP auto-s
[truncated]
/v1/* endpoints require client bearer auth; /healthz , /statsz , /metrics do not.
Pipeline failures degrade gracefully and forward the request instead of returning a 500.
MCP-enabled agents can run local/remote tools without exposing MCP complexity to the client.
Config-driven provider registry — add providers via JSON, zero code changes
23 built-in providers with specialized adapters for wire format differences
Dynamic model discovery — fetches live model catalogs from providers at startup and on reload
Model registry — reference models by string shorthand with automatic provider/context resolution
Multi-provider model resolution — when a model exists in multiple providers, all are added to the agent's fallback chain
Three-tier model resolution — config overrides > discovered models > static registry
Per-model wire format — models from multi-format gateways (OpenCode Zen) auto-convert between OpenAI, Anthropic, and Gemini wire formats based on the model's format attribute
Agent fallback chains — round-robin or sequential with circuit breaker and automatic failover
Latency-aware routing — auto-reorder targets by historical median response time with ±5% jitter to prevent thundering herd
Per-agent system prompts — inline or file-based
Tier-0 regex secret filter — always-on redaction of AWS keys, GitHub tokens, passwords, etc.
3-Tier content pipeline — pluggable interceptor chain: regex redaction, entropy filtering, TF-IDF relevance scoring, engine summarization
Context window compaction — sliding window summarization with configurable engine
Stale tool call pruning — compact old assistant+tool response pairs to save tokens
Thought pruning — strip reasoning blocks from assistant message history
Input validation — strict body limits, JSON sanitization, header filtering
Graceful degradation — never blocks requests due to engine or pipeline failures
Role-Based Access Control (RBAC) — per-API key roles (admin, user, read-only) with agent and endpoint restrictions
Secure memory — mlock-protected token storage, read-only sealing, core dump prevention
Hardening (Deployment Security)
Secure memory (default) : All tokens stored in mlock-protected RAM, sealed read-only after init, core dumps disabled
Non-root execution — runs as UID 65532 with dropped capabilities
Memory protection — LimitMEMLOCK=infinity and LimitCORE=0 in systemd
Read-only filesystem — immutable root + private /tmp
Seccomp + no-new-privileges — restricted syscalls, prevents privilege escalation
Zero-trust secrets — loaded via systemd credentials or container mounts, never to disk
Socket activation — seamless restarts with zero dropped connections
Zero external dependencies — Go standard library only
Hot reload — systemctl reload nenya for zero-downtime config changes
Circuit breaker — per agent+provider+model with automatic failover, exponential backoff, and semantic error classification
Rate limiting — per upstream host (RPM/TPM) with per-provider overrides
Response cache — in-memory LRU with SHA-256 fingerprinting and optional semantic similarity search
Graceful shutdown — 5s grace period for in-flight requests, MCP client cleanup
Context-limit auto-retry — upstream context-length errors trigger summarization and retry
Local engine lifecycle — pre-load and manage local Ollama models with LRU eviction
Structured errors — all error responses include error_kind field for programmatic diagnostics
Tool discovery — connect to MCP servers for automatic tool injection
Multi-turn execution — intercept tool calls, execute against MCP servers, forward results
Auto-search — pre-fetch relevant context from MCP servers before forwarding
Auto-save — persist assistant responses to MCP memory servers
Create minimal config and secrets:
mkdir -p config secrets
cat > config/config.json << ' EOF '
{
"server": { "listen_addr": ":8080" },
"agents": {
"default": {
"strategy": "fallback",
"models": ["gemini-2.5-flash"]
}
}
}
EOF
cat > secrets/provider_keys.json << ' EOF '
{
"provider_keys": {
"gemini": "AIza..."
}
}
EOF
cat > secrets/client.json << ' EOF '
{
"client_token": "nk-$(openssl rand -hex 32)"
}
EOF
Run the container:
podman run -d \
--name nenya \
-p 8080:8080 \
-v ./config:/etc/nenya:ro \
-v ./secrets:/run/secrets/nenya:ro \
-e NENYA_SECRETS_DIR=/run/secrets/nenya \
--cap-drop=ALL \
--cap-add=IPC_LOCK \
--security-opt=no-new-privileges:true \
--read-only \
--tmpfs /tmp:rw,noexec,nosuid,size=64M \
ghcr.io/gumieri/nenya:latest
Test it:
curl -H " Authorization: Bearer $( jq -r ' .client_token ' secrets/client.json ) " \
http://localhost:8080/healthz
Or Install via Package Manager
Nenya provides native packages for major Linux distributions and community package managers:
All packages install the binary to /usr/bin/nenya and include systemd service and socket units. After install, enable and start:
sudo systemctl enable --now nenya.socket
sudo systemctl enable --now nenya.service
Runtime Configuration
Nenya supports standard environment variables for deployment portability:
PORT=9090 HOST=127.0.0.1 ./nenya --config /path/to/config.json
Or in Docker:
docker run -e PORT=9090 -p 9090:9090 ghcr.io/gumieri/nenya:latest
Or Choose Your Deployment
Deploy Bare Metal (systemd) — Direct binary install, socket activation, hot reload
Deploy Container (Podman/Docker Compose) — compose.yml, image verification, security hardening
Deploy Kubernetes (Helm) — Helm chart, ConfigMap/Secret, ingress setup
All /v1/* endpoints require Authorization: Bearer <client_token> or Bearer <api_key_token> .
API keys support RBAC enforcement — agent scoping, endpoint allowlists, role-based permissions (admin bypasses all checks).
See docs/PASSTHROUGH_PROXY.md for detailed passthrough proxy usage.
Document
Description
Providers
All 23 providers, capabilities matrix, special behaviors, adding custom providers
Configuration
Full config reference, directory mode, all sections and fields
Deploy Bare Metal
Systemd unit, config.d layout, secrets, hot reload
Deploy Container
Podman/Docker Compose, image verification, security notes
Deploy Kubernetes
Helm chart usage, ConfigMap/Secret, ingress setup
Passthrough Proxy
Raw provider endpoint proxying, SSE streaming, auth injection
Architecture
Package DAG, request lifecycle, circuit breaker, SSE pipeline
MCP Integration
MCP server integration, tool discovery, multi-turn execution
Adapters
Adapter system internals, auth styles, capability flags
Secrets Format
Systemd credentials, env var fallback, container/K8s deployment
Security
Vulnerability reporting policy
License
A lightweight, highly secure AI API Gateway/Proxy written in Go. Acts as transparent middleware between local AI coding clients (OpenCode/Pi/Cursor) and upstream LLM providers (Gemini, DeepSeek, Zhipu z.ai).
Apache-2.0 license
Security policy
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
10
Nenya 0.6.1
Late

[truncated]
