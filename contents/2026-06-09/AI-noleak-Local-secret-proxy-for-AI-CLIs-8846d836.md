---
source: "https://github.com/ahmedxuhri/ai-noleak"
hn_url: "https://news.ycombinator.com/item?id=48455280"
title: "AI-noleak – Local secret proxy for AI CLIs"
article_title: "GitHub - ahmedxuhri/ai-noleak: Local secret-leak prevention for agentic AI CLIs · GitHub"
author: "ahmedxuhri"
captured_at: "2026-06-09T04:30:03Z"
capture_tool: "hn-digest"
hn_id: 48455280
score: 2
comments: 0
posted_at: "2026-06-09T02:02:35Z"
tags:
  - hacker-news
  - translated
---

# AI-noleak – Local secret proxy for AI CLIs

- HN: [48455280](https://news.ycombinator.com/item?id=48455280)
- Source: [github.com](https://github.com/ahmedxuhri/ai-noleak)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T02:02:35Z

## Translation

タイトル: AI-noleak – AI CLI のローカル シークレット プロキシ
記事のタイトル: GitHub - ahmedxuhri/ai-noleak: エージェント AI CLI のローカル秘密漏洩防止 · GitHub
説明: エージェント AI CLI のローカル秘密漏洩防止。 GitHub でアカウントを作成して、ahmedxuhri/ai-noleak の開発に貢献してください。

記事本文:
GitHub - ahmedxuhri/ai-noleak: エージェント AI CLI のローカル秘密漏洩防止 · GitHub
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
アーメドシュリ
/
アイノリーク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主な支店のタグ

ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .github/ workflows .github/ workflows acceptaccept cmd cmd docs docs 内部内部スクリプト scripts .gitignore .gitignore FUTURE_WORK.md FUTURE_WORK.md LICENSE LICENSE Makefile Makefile README.md README.md SPEC.md SPEC.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントによる秘密の漏洩を防ぐためのローカル編集プロキシ。
ai-noleak は、端末と AI API の間に位置します。誤って公開されたローカル シークレット、認証情報、トークン、API キーをインターセプトし、上流モデルに到達する前に、3 つの独立した保護層にわたって決定論的なローカル プレースホルダー ( @TOKEN_xxxxxx@ ) に置き換えます。
1. シークレットを含むコマンドを実行します。
$ claude "AKIAIOSFODNN7EXAMPLEを使用してbackup.tarをS3にアップロードするスクリプトを作成します"
2. ai-noleak は、Anthropic への送信プロンプトをインターセプトし、シークレットを編集します。
[プロキシ] POST /v1/messages -> 編集された aws_access_key_id (confidence=1.00)
「AKIAIOSFODNN7EXAMPLE」を「@TOKEN_8f51a2@」に置き換えました
3. 上流モデルは安全なプロンプトを受け取ります。
「@TOKEN_8f51a2@ を使用して、backup.tar を S3 にアップロードするスクリプトを作成します。」
4. モデルはプレースホルダーを使用して応答します。
「これがあなたのスクリプトです：export AWS_ACCESS_KEY_ID=@TOKEN_8f51a2@ ...」
5. ai-noleak は、プレースホルダーをローカルで実際のシークレットに変換します。
「これがあなたのスクリプトです：export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE ...」
仕組み
あなたの端末
│
▼
┌─────────────────────────┐
│ レイヤー1・入力（PTYラッパー） │
│ シェルの前に括弧で囲まれたペーストから秘密を取り除きます │
━──────

───────┬─────────────────┘
│
▼
┌─────────────────────────┐
│ レイヤ 2・トランスポート (HTTP プロキシ) │
│ アウトバウンドリクエスト + AI 応答をスキャンして編集 │
━━━━━━━━━━━━━━━━━━━━━┬───────────────┘
│
▼
AI API (OpenAI / Anthropic / …)
│
▼
┌─────────────────────────┐
│ レイヤー3・ストレージ（ファイルウォッチャー） │
│ ログおよび履歴ファイルの inotify/kqueue スキャン │
━━━━━━━━━━━━━━━━━━━━━┘
注: これらのレイヤーは、 SPEC.md で定義されている L1 (物理/TTY 入力)、L2 (データ リンク/トランスポート プロキシ)、および L5 (セッション/アプリケーション ストレージ ファイル) OSI メタファー レイヤーに対応します。
3 つのレイヤーはすべて、プレースホルダーとシークレットのマッピングをメモリ (一時) に保存するか、ディスク上に暗号化 (パスフレーズ モード) して保存する単一のローカル ボールト デーモン ( noleakd ) を共有します。
ai-noleak はローカル HTTP リバース プロキシとして機能します。 TLS トラフィックを復号化するためにカスタム ルート CA 証明書のインストールが必要な従来の MITM プロキシ (高いセキュリティ リスクが発生します) とは異なり、ai-noleak は証明書のインストールを必要としません。 AI CLI ツールをローカルのプレーンテキスト HTTP 経由で http://127.0.0.1:9999/v1 に指定するだけです。 Th

プロキシは、ローカル メモリ内の平文リクエストを傍受、スキャン、編集し、安全で暗号化されたアウトバウンド HTTPS 接続を介して、サニタイズされたペイロードを上流のプロバイダー (Anthropic や OpenAI など) に転送します。
ai-noleak は、次の制御によりこの脅威モデルに対処します。
100% ローカル分離 : テレメトリは決して送信されません。プロキシは、アップストリーム要求が成功するように構成されたプロバイダー認証キーを保存しますが、無関係なローカル シークレットとプロンプト コンテンツは、ホストを離れる前にローカル CPU サイクルで編集されます。
TLS CA/MITM は不要 : ai-noleak は、ネットワーク全体の TLS インターセプターではなくローカル リバース プロキシとして機能するため、root 権限やカスタム ルート CA 証明書を必要としません。プレーンテキスト トラフィックはローカル ループバック インターフェイス ( 127.0.0.1 ) 経由でのみ送信されます。つまり、暗号化されていないトラフィックがマシンから送信されることはありません。
厳密なピア UID 検証: ボールト デーモン ( noleakd ) は、Unix ドメイン ソケット (UDS) 経由でプロキシおよびラッパーと通信します。接続は、ピア資格情報チェック (Linux では SO_PEERCRED、macOS では LOCAL_PEERCRED) を使用してカーネル レベルで検証されます。デーモンを開始したのとまったく同じユーザー ID (UID) が所有するプロセスのみが、ボールトにクエリを実行できます。
権限の分離 : インターセプト HTTP プロキシは、ボールト データベースに関して読み取り専用機能で実行されます。 AC 状態をクエリしたり、プレースホルダー バインディングをチェックしたりできますが、プレーンテキスト ボールトをダンプしたり、シークレット値を変更したりすることはできません。変更コマンド (回転、レビュー、手動追加など) は、クライアントの直接呼び出しに制限されます。
保存時の暗号化 : 永続モード (デフォルト) で実行している場合、ボールト ファイルは AES-256-GCM で暗号化されます。キーは、サービスの起動時に一度要求されるマスター パスフレーズから Argon2id を使用して派生します。
Agentic AI CLI は com を作成して実行します

mands、grep ファイル、および読み取りログ。コマンド履歴にアクティブな環境変数、.env ファイル、.git/config 資格情報、または生のトークンがある場合、エージェントがそれらを誤って読み取り、プロンプト コンテキストに挿入し、上流の AI API またはサードパーティ プロキシに送信することが非常に簡単です。
貼り付けられたシークレットは、シェルが実行する前にスクラブされます (レイヤー 1)。
AI プロバイダーへのアウトバウンド HTTP リクエストは、マシンから送信される前に生のシークレットをプレースホルダーに置き換えます (レイヤー 2)。
ディスクに書き込まれた一時的なシェル スナップショット、ログ、または履歴ファイルはすぐに消去されます (レイヤー 3)。
アップストリーム AI モデルは、 @TOKEN_a9553f@ のようなプレースホルダーのみを参照します。モデルがプレースホルダーを出力する場合、ai-noleak はそれを CLI に返す前にローカルで実際のシークレットに変換し直します。
これにより、無関係なローカル シークレットが AI リクエストに誤って漏洩することが防止されます (プロバイダー自身の API キーは認証が成功するように保存されますが、他のすべてのローカル シークレットは編集されます)。
ai-noleak は、次のいずれかの方法を使用してインストールできます。
カール -fsSL https://raw.githubusercontent.com/ahmedxuhri/ai-noleak/main/scripts/install.sh |しー
方法 B: 検証済みインストール (推奨)
インストール スクリプトを監査し、実行前にその整合性を検証するには:
# 1. インストーラースクリプトをダウンロードする
カール -fsSL -o install.sh https://raw.githubusercontent.com/ahmedxuhri/ai-noleak/main/scripts/install.sh
# 2. インストーラーのチェックサムが公式ハッシュと一致することを確認する
echo " 6a37d48892d15aa42c67a922cbfd71b35a220e8906b17673563ba0498e63b9f8 install.sh " | sha256sum -c -
# 3. インストーラーを実行する
sh install.sh && rm install.sh
注: バイナリ ( noleak 、 noleakd 、および noleak-watch ) は ~/.local/bin (root として実行する場合は /usr/local/bin) にインストールされます。
ソースからのビルド (代替)
Go Bina をコンパイルしたい場合

手動で実行します:
git クローン https://github.com/ahmedxuhri/ai-noleak.git
CD アイノリーク
make build # バイナリを ./bin/ にコンパイルします
make test # テストスイートを実行します
Go 1.22 以降が必要です。
すべてのビルド済みリリースには、GitHub リリース ページで公開されているバイナリと一致する SHA256SUMS.txt が含まれています。ダウンロードした tarball は、以下を使用して確認できます。
sha256sum -c SHA256SUMS.txt
クイックスタート
~/.noleak/config.yaml を編集して、proxy_upstream を AI CLI が通常呼び出す API エンドポイントに設定します。
プロキシリッスン : 127.0.0.1:9999
proxy_upstream : https://api.anthropic.com # または https://api.openai.com
proxy_preserve_headers :
- 認可
- X-API-キー
- 人間バージョン
- 人類ベータ
- ユーザーエージェント
- 受け入れる
- コンテンツタイプ
proxy_passthrough_tokens : []
proxy_passthrough_tokens — エンドポイントに到達する必要があるアップストリーム認証トークンのみ。ここにはプロバイダー API キーやアカウント資格情報を決して入力しないでください。
1 つのコマンドで 3 つのセキュリティ層 (ボルト デーモン、HTTP プロキシ、およびファイル ウォッチャー) をすべて同時に開始します。
# エフェメラル モード (メモリ内のみ、パスフレーズなし - テストに最適):
noleak start --ephemeral
# 永続モード (ディスク上で暗号化 - 起動時にマスター パスフレーズの入力を求める):
漏れのないスタート
3 · AI CLI をプロキシに向ける
エージェントの CLI ベース URL を http://127.0.0.1:9999/v1 に設定します。
エクスポート ANTHROPIC_BASE_URL= " http://127.0.0.1:9999/v1 "
OpenAI コーデックス ( ~/.codex/config.toml ):
モデル = " gpt-4o "
model_provider = " openai-custom "
[プロバイダー。オープンアイカスタム】
名前 = " openai-custom "
Base_url = " http://127.0.0.1:9999/v1 "
env_key = " OPENAI_API_KEY "
4 · ヘルスチェックを実行する
すべてのサービスが実行中であり、正しく接続されていることを確認します。
ノーリークドクター
期待される出力:
[ok] config /root/.noleak/config.yaml
[ok] 構成検証は有効です
[ok] ソケット /root/.noleak/sock
[ok] デーモンの健全性バージョン = 0.1.0 ロック解除 = true v

ault_entries=0 pending_review=0
[ok] プロキシリッスン 127.0.0.1:9999
[OK] プロキシ アップストリーム https://api.anthropic.com
...
インタラクティブなデモ
エージェントがキーを漏洩しようとしたときの noleak start の実行は次のようになります。
$ noleak start --ephemeral
[noleakd] /root/.noleak/sock で開始 (保管庫: 一時メモリ)
[プロキシ] 127.0.0.1:9999 でリッスン -> https://api.anthropic.com
[ウォッチャー] /root/.bash_history を監視 (編集)
[ウォッチャー] /root/.claude/projects を監視 (編集)
# 別のターミナルで、エージェントはプロンプトで AWS キーを使用して API の呼び出しを試みます。
$カール -s -X POST http://127.0.0.1:9999/v1/responses \
-d '{"プロンプト": "AKIAIOSFODNN7EXAMPLEを使用"}'
# noleak ターミナルは直ちにインターセプトして編集します。
[プロキシ] 12:41:45 request /v1/responses -> @TOKEN_322a15@ (kind=aws_access_key_id, conf=1.00)
手動テスト
実際の AI CLI を実行しなくても、3 つの保護層すべてを検証できます。
レイヤ 2 のテスト — プロキシの編集
curl -s --max-time 10 \
-X POST http://127.0.0.1:9999/v1/responses \
-H " Content-Type: application/json " \
-H " 権限: ベアラー sk-test-fake " \
-d ' {"model":"gpt-4o","input":"私の AWS キーは AKIAIOSFODNN7EXAMPLE","stream":true} '
noleak start のコンソール ログには次が出力されます。
[プロキシ] リクエスト /v1/responses -> @TOKEN_xxxxxx@ (kind=aws_access_key_id, conf=1.00)
レイヤ 3 のテスト — オンディスク ウォッチャーのリダクション
テスト ディレクトリを指定して noleak を開始します。
noleak start --ephemeral --redact ~ /test-watch
別のターミナルで、シークレットを書き込みます。
echo " GitHub PAT: ghp_A1B2C3D4E5F6G7H8I9J0K1L2M3N4O5P6Q7R8 " > ~ /test-watch/leaked.txt
睡眠2
猫 ~ /test-watch/leaked.txt
# 出力: GitHub PAT: @TOKEN_xxxxxx@
テストレイヤー 1 — PTY ペーストフィルター
printf " \x1b[200~ghp_A1B2C3D4E5F6G7H8I9J0K1L2M3N4O5P6Q7R8\x1b[201~\n " \
|ノーリークランバッシュ
# 生の秘密は s の前に削除されます

地獄はそれを受け取る
CLI リファレンス
コマンド
説明
漏れのないスタート
デーモン、プロキシ、ウォッチャーを同時に起動
noleak start --ephemeral
メモリ内のすべてのサービスを開始します (パスフレーズなし)
ノリークr

[切り捨てられた]

## Original Extract

Local secret-leak prevention for agentic AI CLIs. Contribute to ahmedxuhri/ai-noleak development by creating an account on GitHub.

GitHub - ahmedxuhri/ai-noleak: Local secret-leak prevention for agentic AI CLIs · GitHub
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
ahmedxuhri
/
ai-noleak
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .github/ workflows .github/ workflows acceptance acceptance cmd cmd docs docs internal internal scripts scripts .gitignore .gitignore FUTURE_WORK.md FUTURE_WORK.md LICENSE LICENSE Makefile Makefile README.md README.md SPEC.md SPEC.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Local redaction proxy for preventing AI coding agents from leaking your secrets.
ai-noleak sits between your terminal and any AI API. It intercepts accidentally exposed local secrets, credentials, tokens, and API keys — replacing them with deterministic local placeholders ( @TOKEN_xxxxxx@ ) across three independent protection layers before they can reach the upstream model.
1. You run a command containing a secret:
$ claude "Write a script to upload backup.tar to S3 using AKIAIOSFODNN7EXAMPLE"
2. ai-noleak intercepts the outbound prompt to Anthropic and redacts the secret:
[proxy] POST /v1/messages -> Redacted aws_access_key_id (confidence=1.00)
Replaced "AKIAIOSFODNN7EXAMPLE" with "@TOKEN_8f51a2@"
3. The upstream model receives the safe prompt:
"Write a script to upload backup.tar to S3 using @TOKEN_8f51a2@"
4. The model answers using the placeholder:
"Here is your script: export AWS_ACCESS_KEY_ID=@TOKEN_8f51a2@ ..."
5. ai-noleak translates the placeholder back to the real secret locally:
"Here is your script: export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE ..."
How It Works
Your Terminal
│
▼
┌─────────────────────────────────────────────────────────┐
│ Layer 1 · Input (PTY Wrapper) │
│ Strips secrets from bracketed-paste before shell │
└───────────────────────┬─────────────────────────────────┘
│
▼
┌─────────────────────────────────────────────────────────┐
│ Layer 2 · Transport (HTTP Proxy) │
│ Scans & redacts outbound requests + AI responses │
└───────────────────────┬─────────────────────────────────┘
│
▼
AI API (OpenAI / Anthropic / …)
│
▼
┌─────────────────────────────────────────────────────────┐
│ Layer 3 · Storage (File Watcher) │
│ inotify/kqueue scan of log & history files │
└─────────────────────────────────────────────────────────┘
Note: The layers correspond to L1 (Physical/TTY input), L2 (Data Link/Transport proxy), and L5 (Session/Application storage files) OSI-metaphor layers defined in the SPEC.md .
All three layers share a single local vault daemon ( noleakd ) that stores the placeholder↔secret mapping in memory (ephemeral) or encrypted on disk (passphrase mode).
ai-noleak acts as a local HTTP reverse proxy. Unlike traditional MITM proxies that require installing a custom Root CA certificate to decrypt TLS traffic (which introduces high security risks), ai-noleak requires zero certificate installation. You simply point your AI CLI tools to http://127.0.0.1:9999/v1 over plaintext HTTP locally. The proxy intercepts, scans, and redacts the plaintext requests in local memory, and then forwards the sanitized payload to the upstream provider (e.g. Anthropic or OpenAI) over a secure, encrypted outbound HTTPS connection.
ai-noleak addresses this threat model with the following controls:
100% Local Isolation : No telemetry is ever sent. The proxy preserves the configured provider authorization keys so that upstream requests succeed, but unrelated local secrets and prompt content are redacted on local CPU cycles before leaving the host.
No TLS CA/MITM Needed : By acting as a local reverse proxy rather than a network-wide TLS interceptor, ai-noleak does not require root privileges or custom root CA certificates. Plaintext traffic is only transmitted over the local loopback interface ( 127.0.0.1 ), meaning no unencrypted traffic leaves your machine.
Strict Peer UID Verification : The vault daemon ( noleakd ) communicates with the proxy and wrapper via a Unix Domain Socket (UDS). Connections are validated at the kernel level using peer credentials checking ( SO_PEERCRED on Linux, LOCAL_PEERCRED on macOS). Only processes owned by the exact same User ID (UID) that started the daemon can query the vault.
Privilege Separation : The intercepting HTTP proxy runs with read-only capabilities with respect to the vault database. It can query the AC state and check placeholder bindings, but it cannot dump the plaintext vault or modify secret values. Mutating commands (like rotate , review , and manual add ) are restricted to direct client invocations.
Encryption at Rest : When running in persistent mode (default), the vault file is encrypted with AES-256-GCM. The key is derived using Argon2id from a master passphrase prompted once at service startup.
Agentic AI CLIs write and run commands, grep files, and read logs. If you have active environment variables, .env files, .git/config credentials, or raw tokens in your command history, it is incredibly easy for an agent to accidentally read them, inject them into its prompt context, and send them upstream to an AI API or third-party proxy.
Pasted secrets are scrubbed before the shell executes them ( Layer 1 ).
Outbound HTTP requests to AI providers replace raw secrets with placeholders before leaving the machine ( Layer 2 ).
Temporary shell snapshots, logs, or history files written to disk are cleaned immediately ( Layer 3 ).
Upstream AI models only see placeholders like @TOKEN_a9553f@ . If the model outputs the placeholder, ai-noleak translates it back to the real secret locally before returning it to the CLI.
This prevents accidental leakage of unrelated local secrets into AI requests (the provider's own API key is preserved so that auth succeeds, but all other local secrets are redacted).
You can install ai-noleak using one of the following methods:
curl -fsSL https://raw.githubusercontent.com/ahmedxuhri/ai-noleak/main/scripts/install.sh | sh
Method B: Verified Installation (Recommended)
To audit the installation script and verify its integrity before execution:
# 1. Download the installer script
curl -fsSL -o install.sh https://raw.githubusercontent.com/ahmedxuhri/ai-noleak/main/scripts/install.sh
# 2. Verify the installer's checksum matches the official hash
echo " 6a37d48892d15aa42c67a922cbfd71b35a220e8906b17673563ba0498e63b9f8 install.sh " | sha256sum -c -
# 3. Run the installer
sh install.sh && rm install.sh
Note: Binaries ( noleak , noleakd , and noleak-watch ) are installed to ~/.local/bin (or /usr/local/bin if run as root).
Build from Source (Alternative)
If you prefer to compile the Go binaries manually:
git clone https://github.com/ahmedxuhri/ai-noleak.git
cd ai-noleak
make build # compiles binaries to ./bin/
make test # runs the test suite
Requires Go 1.22+.
All prebuilt releases contain SHA256SUMS.txt matching the published binaries on the GitHub Releases page. You can verify downloaded tarballs using:
sha256sum -c SHA256SUMS.txt
Quick Start
Edit ~/.noleak/config.yaml to set proxy_upstream to the API endpoint your AI CLI normally calls:
proxy_listen : 127.0.0.1:9999
proxy_upstream : https://api.anthropic.com # or https://api.openai.com
proxy_preserve_headers :
- Authorization
- X-Api-Key
- Anthropic-Version
- Anthropic-Beta
- User-Agent
- Accept
- Content-Type
proxy_passthrough_tokens : []
proxy_passthrough_tokens — only for upstream auth tokens that must reach the endpoint. Never put provider API keys or account credentials here.
Start all three security layers (vault daemon, HTTP proxy, and file watcher) concurrently in a single command:
# Ephemeral Mode (in-memory only, no passphrase — best for testing):
noleak start --ephemeral
# Persistent Mode (encrypted on disk — prompts for master passphrase on startup):
noleak start
3 · Point Your AI CLI at the Proxy
Configure your agent CLI base URL to http://127.0.0.1:9999/v1 .
export ANTHROPIC_BASE_URL= " http://127.0.0.1:9999/v1 "
OpenAI Codex ( ~/.codex/config.toml ):
model = " gpt-4o "
model_provider = " openai-custom "
[ providers . openai-custom ]
name = " openai-custom "
base_url = " http://127.0.0.1:9999/v1 "
env_key = " OPENAI_API_KEY "
4 · Run the Health Check
Verify all services are running and correctly connected:
noleak doctor
Expected output:
[ok] config /root/.noleak/config.yaml
[ok] config validation valid
[ok] socket /root/.noleak/sock
[ok] daemon health version=0.1.0 unlocked=true vault_entries=0 pending_review=0
[ok] proxy listen 127.0.0.1:9999
[ok] proxy upstream https://api.anthropic.com
...
Interactive Demo
Here is what running noleak start looks like when an agent attempts to leak a key:
$ noleak start --ephemeral
[noleakd] starting on /root/.noleak/sock (vault: ephemeral memory)
[proxy] listening on 127.0.0.1:9999 -> https://api.anthropic.com
[watcher] watching /root/.bash_history (redact)
[watcher] watching /root/.claude/projects (redact)
# In another terminal, an agent tries to call the API with an AWS key in the prompt:
$ curl -s -X POST http://127.0.0.1:9999/v1/responses \
-d '{"prompt": "use AKIAIOSFODNN7EXAMPLE"}'
# noleak terminal immediately intercepts and redacts:
[proxy] 12:41:45 request /v1/responses -> @TOKEN_322a15@ (kind=aws_access_key_id, conf=1.00)
Manual Testing
You can verify all three protection layers without running a real AI CLI:
Test Layer 2 — Proxy Redaction
curl -s --max-time 10 \
-X POST http://127.0.0.1:9999/v1/responses \
-H " Content-Type: application/json " \
-H " Authorization: Bearer sk-test-fake " \
-d ' {"model":"gpt-4o","input":"My AWS key is AKIAIOSFODNN7EXAMPLE","stream":true} '
The console logs of noleak start will output:
[proxy] request /v1/responses -> @TOKEN_xxxxxx@ (kind=aws_access_key_id, conf=1.00)
Test Layer 3 — On-Disk Watcher Redaction
Start noleak start specifying a test directory:
noleak start --ephemeral --redact ~ /test-watch
In another terminal, write a secret:
echo " GitHub PAT: ghp_A1B2C3D4E5F6G7H8I9J0K1L2M3N4O5P6Q7R8 " > ~ /test-watch/leaked.txt
sleep 2
cat ~ /test-watch/leaked.txt
# Output: GitHub PAT: @TOKEN_xxxxxx@
Test Layer 1 — PTY Paste Filter
printf " \x1b[200~ghp_A1B2C3D4E5F6G7H8I9J0K1L2M3N4O5P6Q7R8\x1b[201~\n " \
| noleak run bash
# The raw secret is stripped before the shell receives it
CLI Reference
Command
Description
noleak start
Start daemon, proxy, and watcher concurrently
noleak start --ephemeral
Start all services in-memory (no passphrase)
noleak r

[truncated]
