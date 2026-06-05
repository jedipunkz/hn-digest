---
source: "https://github.com/mirkolenz/llmhop"
hn_url: "https://news.ycombinator.com/item?id=48406517"
title: "Show HN: LLMhop – A tiny, stateless router for LLMs with a NixOS module"
article_title: "GitHub - mirkolenz/llmhop: Tiny, stateless Go router that dispatches OpenAI-compatible requests to single-model vLLM and sglang backends with zero external dependencies · GitHub"
author: "mlenz"
captured_at: "2026-06-05T00:57:21Z"
capture_tool: "hn-digest"
hn_id: 48406517
score: 1
comments: 0
posted_at: "2026-06-05T00:26:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LLMhop – A tiny, stateless router for LLMs with a NixOS module

- HN: [48406517](https://news.ycombinator.com/item?id=48406517)
- Source: [github.com](https://github.com/mirkolenz/llmhop)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T00:26:09Z

## Translation

タイトル: Show HN: LLMhop – NixOS モジュールを備えた LLM 用の小さなステートレスルーター
記事のタイトル: GitHub - mirkolenz/llmhop: OpenAI 互換リクエストを単一モデル vLLM および sglang バックエンドに外部依存関係なしでディスパッチする小型のステートレス Go ルーター · GitHub
説明: OpenAI 互換リクエストを外部依存関係なしで単一モデルの vLLM および sglang バックエンドにディスパッチする小型のステートレス Go ルーター - mirkolenz/llmhop
HN テキスト: LLMhop は、LLM 推論サーバー用の小さなステートレス プロキシです。これは、vLLM でネイティブにサポートされていない複数のローカル LLM を同時に提供しようとしたときに直面した問題に対処します。 LLMhop バイナリは、リクエストのモデル フィールドを検査し、オプションの認証処理を使用してリクエストを正しいバックエンド サービスにルーティングします。さらに、Quadlet/Podman 経由で llama.cpp、vLLM、および sglang を実行し、プロキシに自動登録するための NixOS モジュールが含まれています。

記事本文:
GitHub - mirkolenz/llmhop: 外部依存関係なしで OpenAI 互換リクエストを単一モデルの vLLM および sglang バックエンドにディスパッチする小型のステートレス Go ルーター · GitHub
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
ミルコレンツ
/
イルムホップ
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット .github/ workflows .github/ workflows cmd/ llmhop cmd/ llmhop docs docs 内部 内部 nix nix .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md flake.lock flake.lock flake.nix flake.nix go.mod go.mod renovate.json renovate.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1 つのポート、多数のモデル: OpenAI 互換の LLM 推論バックエンド用の小型のステートレス HTTP ルーター。
LLMhop は、受信した OpenAI 互換リクエストのモデル フィールドを調べ、それを一致するバックエンドにリバース プロキシします。
これは主に、プロセスごとに 1 つのモデルを提供し、その前にモデル対応のシン ゲートウェイを必要とする vLLM や sglang などの単一モデル推論サーバー向けに設計されていますが、単一のエンドポイントの背後に複数のアップストリームを統合したい場合は常に、OpenAI 互換のバックエンド (マルチモデル サーバーやホストされたプロバイダーを含む) で動作します。
OpenAI 互換のリバース プロキシ、モデル ルーター、およびセルフホスト型 LLM 推論用のリクエスト ディスパッチャー。
ステートレスなシングルバイナリ HTTP サービス: データベース、キャッシュ、バックグラウンド ワーカーがなく、ロード バランサの背後で安全です。
外部依存関係なし: 純粋な Go、サードパーティ パッケージ、CGO なし。
セルフホストまたはリモートの OpenAI API 互換バックエンドで動作します: vLLM、sglang、TabbyAPI、Aphrodite、Ollama、LocalAI、OpenRouter、togetter.ai、DeepInfra など。
静的バイナリ、最小限の Docker イメージ、およびオプションでルーターと一緒に llama.cpp、sglang、または vLLM ワーカーを起動できる強化された NixOS モジュールとして出荷されます。
クライアントは、 {"model": "..."} を含む JSON 本文を含むリクエストを送信します。
LLMhop はモデルフィールドを読み取り、それを検索します。

その構成。
リクエストは構成されたバックエンド URL にそのまま転送されます。
LLMhop はオプションで、受信リクエストをベアラー トークンのリストでゲートし、バックエンドに転送するときにモデルごとの Authorization (またはその他の) ヘッダーを挿入できます。
両方の側がオプトインです。authTokens と models.*.headers を設定しないままにし、ヘッダーはそのまま転送されます。
authTokens が設定されている場合、ルーターは受信した Authorization: Bearer <token> ヘッダー (定数時間比較) を検証し、転送する前にそれを削除するため、クライアント側のトークンが上流に漏洩することはありません。
モデルごとのヘッダーは最後に適用されるため、構成された認可はクライアントが送信したものよりも常に優先されます。
{
"聞く" : " :8080 " 、
"authTokens" : [ " ${file:client_token} " ],
「モデル」: {
"ラマ-3-8b" : {
"url" : " http://localhost:30000 "
}、
"openai-gpt-4o" : {
"url" : " https://api.openai.com " ,
"ヘッダー" : {
"認可" : " ベアラー ${env:OPENAI_KEY} "
}
}
}
}
秘密の参照
authTokens および models.*.headers 内の文字列値は起動時に展開されるため、構成ファイル内にプレーンテキストのシークレットが存在する必要はありません。
${env:NAME} : NAME 環境変数から読み取ります。
${file:path} : ファイルから読み取ります。相対パスは、設定されている場合 (たとえば、 LoadCredential= を使用して systemd によって起動された場合)、 $CREDENTIALS_DIRECTORY に対して解決され、それ以外の場合は、現在の作業ディレクトリに対して解決されます。末尾の 1 つの改行は切り取られます。
$NAME : ${env:NAME} の短縮形。
未解決の参照はハード スタートアップ エラーです。
LLMhop は各リクエスト本文をメモリにバッファリングして、転送する前にモデル フィールドを参照できるようにします。
単一のリクエストでメモリを使い果たさないようにするため、本文の上限はデフォルトで 100 MiB に設定されています。上限を超えるボディは 413 Request Entity Too Large で拒否されます。
ビジョンまたはその他のマルチモーダル ペイロードでさらに必要な場合は、これをオーバーライドします。
{ "最大ボディバイト数"

: 524288000 }
ランニング
# ネイティブ
llmhop --config config.json
#ニックス
nix run github:mirkolenz/llmhop -- --config config.json
# ドッカー
docker run --rm -p 8080:8080 -v ./config.json:/config.json ghcr.io/mirkolenz/llmhop --config /config.json
NixOSモジュール
強化された systemd サービスがすぐに提供されます。
LLMhop をフレーク入力に追加し、モジュールをシステム構成にインポートします。
{
入力 = {
nixpkgs 。 url = "github:nixos/nixpkgs/nixos-unstable" ;
llmhop = {
URL = "github:mirkolenz/llmhop" ;
入力。 nixpkgs 。フォロー = "nixpkgs" ;
} ;
} ;
出力 =
{ nixpkgs 、 llmhop 、 ... } :
{
nixosConfiguration 。 myhost = nixpkgs 。ライブラリ 。 nixosシステム {
システム = "x86_64-linux" ;
モジュール = [
イルムホップ 。 nixosModules 。デフォルト
{
サービス。 llmhop = {
有効 = true ;
設定 = {
listen = ":8080" ;
モデル = {
「ラマ-3-8b」 。 URL = "http://localhost:30000" ;
「qwen-2.5-7b」 。 URL = "http://localhost:30001" ;
} ;
} ;
} ;
}
] ;
} ;
} ;
}
このユニットは、積極的なサンドボックス ( ProtectSystem 、 PrivateTmp 、制限された syscall とアドレス ファミリ、新しい特権なしなど) を備えた DynamicUser の下で実行され、障害が発生すると再起動されます。
このモジュールは推論サーバー自体を実行することもできるため、llama.cpp、sglang、または vLLM を手動で接続する必要はありません。
各バックエンドは、services.llmhop.<backend> の下にモデル属性セットを公開し、すべてのエントリがループバック ポートにバインドされた 1 つの分離されたワーカーとなり、一致するルートが llmhop に自動的に登録されます。
3 つのバックエンドはすべて、同じ構成内で並べて有効化し、自由に混合できます。
llama.cpp は、 DynamicUser の下でネイティブの強化された systemd システム ユニットとして実行されます。
sglang と vLLM は、quadlet-nix を通じてルートレス Podman コンテナとして起動されます。
各 Quadlet バックエンドは、キャッシュ ディレクトリ sub を所有する専用の残留システム ユーザー ( sglang 、 vllm ) を取得します。

-UID range and rootless container store.
The container units are installed under that user's per-UID search path and therefore run as systemd user units , not system units.
This is a deliberate workaround for NVIDIA/nvidia-container-toolkit#648 :
nvidia-cdi-hook runs as an OCI createContainer hook inside the container's user namespace and fails to read the OCI bundle's config.json whenever Podman uses a UID-mapped namespace (e.g., --userns auto or --userns nomap ), which is the mode you end up in when systemd's system manager launches a rootless container.
Running each Quadlet unit under a real, lingering system user's systemd instance keeps Podman in the keep-id -style mapping where the CDI hook can read the bundle and the GPU is correctly exposed.
No worker ever runs as root.
For convenience, the module injects a tiny per-backend helper into environment.systemPackages whenever the backend's default user is used:
llama-cpp workers are plain system units, so they are managed with the usual systemctl status llama-cpp-<model> and journalctl -u llama-cpp-<model> .
sglang-shell and vllm-shell are writeShellApplication wrappers around machinectl shell that drop you into the backend user's session, where systemctl --user , journalctl --user and podman ps see the worker units directly. Run them with no arguments for an interactive shell, or pass a command to execute it inside the session.
サービス。 llmhop = {
有効 = true ;
ラマ-cpp = {
有効 = true ;
モデル。 "qwen3-8b" = {
ポート = 18001 ;
設定 。 hf-repo = "unsloth/Qwen3-8B-GGUF:UD-Q4_K_XL" ;
} ;
} ;
スラング = {
有効 = true ;
モデル。 "qwen3-コーダー" = {
ポート = 19001 ;
model = "Qwen/Qwen3-8B" ;
設定 。 reasoning-parser = "qwen3" ;
} ;
} ;
vllm = {
有効 = true ;
モデル。 "ラマ-3-8b" = {
ポート = 20001 ;
モデル = "メタ

-llama/Meta-Llama-3-8B-Instruct" ;
} ;
} ;
} ;
バックエンドごとのオプションの完全なリストについては、オプション リファレンスを参照してください。
生成された構成ファイルは世界中で読み取り可能な Nix ストアに保存されるため、シークレットを services.llmhop.settings に直接配置しないでください。
代わりに、 ${file:...} を介してそれらを参照し、 systemd の LoadCredential= を使用してファイルをサービスに渡します。
各 LoadCredential エントリの右側は単なるファイル パスであるため、agenix または sops-nix の出力、 /etc/llmhop/ の下にある手動で管理されるファイル、または独自のシークレット プロビジョニング ツールによって発行されたパスなど、ファイルを生成するものはすべて機能します。
サービス。イルムホップ 。設定 = {
authTokens = [ " \ ${file:client_token}" ] ;
モデル。 "openai-gpt-4o" = {
URL = "https://api.openai.com" ;
ヘッダー 。権限 = "ベアラー \ ${env:OPENAI_KEY}" ;
} ;
} ;
システムド。サービス。イルムホップ 。 serviceConfig = {
LoadCredential = [ "client_token:/etc/llmhop/client-token" ] ;
環境ファイル = [ "/etc/llmhop/openai.env" ] ;
} ;
/etc/llmhop/openai.env はプレーンな KEY=VALUE ファイルです。
OPENAI_KEY = sk-...
${file:...} 参照は、 $CREDENTIALS_DIRECTORY に対して解決されます。 systemd は、このサービスのみにアクセス可能なユニットごとの tmpfs として公開し、DynamicUser およびサンドボックスの残りの部分と互換性があります。
${env:...} は、通常、EnvironmentFile= を介して、ユニットが継承するものをすべて取得します。
秘密ツールがデータを渡す方法に一致するものを選択してください。両方を 1 つの構成に混在させても問題ありません。
OpenAI 互換リクエストを外部依存関係のない単一モデルの vLLM および sglang バックエンドにディスパッチする、小型のステートレス Go ルーター
mirkolenz.github.io/llmhop/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
6
v1.2.2
最新の
2026 年 5 月 29 日
+ 5 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください

年 。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Tiny, stateless Go router that dispatches OpenAI-compatible requests to single-model vLLM and sglang backends with zero external dependencies - mirkolenz/llmhop

LLMhop is a tiny stateless proxy for LLM inference servers. It tackles an issue I faced when trying to serve more than one local LLM at once which is not natively supported by vLLM. The LLMhop binary inspects the model field of the request and routes it to the correct backend service with optional handling of authentication. In addition, it contains a NixOS module to run llama.cpp, vLLM, and sglang via Quadlet/Podman and auto-register with the proxy.

GitHub - mirkolenz/llmhop: Tiny, stateless Go router that dispatches OpenAI-compatible requests to single-model vLLM and sglang backends with zero external dependencies · GitHub
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
mirkolenz
/
llmhop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits .github/ workflows .github/ workflows cmd/ llmhop cmd/ llmhop docs docs internal internal nix nix .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md flake.lock flake.lock flake.nix flake.nix go.mod go.mod renovate.json renovate.json View all files Repository files navigation
One port, many models: A tiny, stateless HTTP router for OpenAI-compatible LLM inference backends.
LLMhop peeks at the model field of an incoming OpenAI-compatible request and reverse-proxies it to the matching backend.
It is primarily designed for single-model inference servers like vLLM and sglang that serve one model per process and need a thin model-aware gateway in front of them, but it works with any OpenAI-compatible backend (including multi-model servers and hosted providers) whenever you want to consolidate several upstreams behind a single endpoint.
OpenAI-compatible reverse proxy, model router and request dispatcher for self-hosted LLM inference.
Stateless single-binary HTTP service: no database, no cache, no background workers, safe behind any load balancer.
Zero external dependencies: pure Go, no third-party packages, no CGO.
Works with any OpenAI API-compatible backend, self-hosted or remote: vLLM, sglang, TabbyAPI, Aphrodite, Ollama, LocalAI, OpenRouter, together.ai, DeepInfra, etc.
Ships as a static binary, a minimal Docker image and a hardened NixOS module that can optionally spin up llama.cpp, sglang or vLLM workers alongside the router.
Client sends a request with a JSON body containing {"model": "..."} .
LLMhop reads the model field and looks it up in its config.
The request is forwarded verbatim to the configured backend URL.
LLMhop can optionally gate incoming requests with a list of bearer tokens and inject per-model Authorization (or any other) headers when forwarding to the backend.
Both sides are opt-in: leave authTokens and models.*.headers unset and headers are forwarded verbatim.
When authTokens is set, the router validates the incoming Authorization: Bearer <token> header (constant-time compare) and then strips it before forwarding, so the client-facing token never leaks upstream.
Per-model headers are applied last, so a configured Authorization always wins over whatever the client sent.
{
"listen" : " :8080 " ,
"authTokens" : [ " ${file:client_token} " ],
"models" : {
"llama-3-8b" : {
"url" : " http://localhost:30000 "
},
"openai-gpt-4o" : {
"url" : " https://api.openai.com " ,
"headers" : {
"Authorization" : " Bearer ${env:OPENAI_KEY} "
}
}
}
}
Secret references
String values inside authTokens and models.*.headers are expanded at startup, so no plaintext secret ever has to live in the config file:
${env:NAME} : read from the NAME environment variable.
${file:path} : read from a file. Relative paths are resolved against $CREDENTIALS_DIRECTORY when set (e.g. when launched by systemd with LoadCredential= ), otherwise against the current working directory. A single trailing newline is trimmed.
$NAME : shorthand for ${env:NAME} .
Unresolved references are a hard startup error.
LLMhop buffers each request body in memory so it can peek at the model field before forwarding.
To keep a single request from exhausting memory, the body is capped at 100 MiB by default; bodies beyond the cap are rejected with 413 Request Entity Too Large .
Override it when vision or other multimodal payloads need more:
{ "maxBodyBytes" : 524288000 }
Running
# native
llmhop --config config.json
# nix
nix run github:mirkolenz/llmhop -- --config config.json
# docker
docker run --rm -p 8080:8080 -v ./config.json:/config.json ghcr.io/mirkolenz/llmhop --config /config.json
NixOS module
A hardened systemd service is provided out of the box.
Add LLMhop to your flake inputs and import the module into your system configuration:
{
inputs = {
nixpkgs . url = "github:nixos/nixpkgs/nixos-unstable" ;
llmhop = {
url = "github:mirkolenz/llmhop" ;
inputs . nixpkgs . follows = "nixpkgs" ;
} ;
} ;
outputs =
{ nixpkgs , llmhop , ... } :
{
nixosConfigurations . myhost = nixpkgs . lib . nixosSystem {
system = "x86_64-linux" ;
modules = [
llmhop . nixosModules . default
{
services . llmhop = {
enable = true ;
settings = {
listen = ":8080" ;
models = {
"llama-3-8b" . url = "http://localhost:30000" ;
"qwen-2.5-7b" . url = "http://localhost:30001" ;
} ;
} ;
} ;
}
] ;
} ;
} ;
}
The unit runs under DynamicUser with aggressive sandboxing ( ProtectSystem , PrivateTmp , restricted syscalls and address families, no new privileges, ...) and restarts on failure.
The module can also run the inference servers themselves, so you don't have to wire up llama.cpp, sglang or vLLM by hand.
Each backend exposes a models attrset under services.llmhop.<backend> and every entry becomes one isolated worker bound to a loopback port, with the matching route registered automatically with llmhop.
All three backends can be enabled side by side and mixed freely in the same configuration.
llama.cpp runs as a native, hardened systemd system unit under DynamicUser .
sglang and vLLM are launched as rootless Podman containers through quadlet-nix .
Each Quadlet backend gets a dedicated, lingering system user ( sglang , vllm ) that owns its cache directory, sub-UID range and rootless container store.
The container units are installed under that user's per-UID search path and therefore run as systemd user units , not system units.
This is a deliberate workaround for NVIDIA/nvidia-container-toolkit#648 :
nvidia-cdi-hook runs as an OCI createContainer hook inside the container's user namespace and fails to read the OCI bundle's config.json whenever Podman uses a UID-mapped namespace (e.g., --userns auto or --userns nomap ), which is the mode you end up in when systemd's system manager launches a rootless container.
Running each Quadlet unit under a real, lingering system user's systemd instance keeps Podman in the keep-id -style mapping where the CDI hook can read the bundle and the GPU is correctly exposed.
No worker ever runs as root.
For convenience, the module injects a tiny per-backend helper into environment.systemPackages whenever the backend's default user is used:
llama-cpp workers are plain system units, so they are managed with the usual systemctl status llama-cpp-<model> and journalctl -u llama-cpp-<model> .
sglang-shell and vllm-shell are writeShellApplication wrappers around machinectl shell that drop you into the backend user's session, where systemctl --user , journalctl --user and podman ps see the worker units directly. Run them with no arguments for an interactive shell, or pass a command to execute it inside the session.
services . llmhop = {
enable = true ;
llama-cpp = {
enable = true ;
models . "qwen3-8b" = {
port = 18001 ;
settings . hf-repo = "unsloth/Qwen3-8B-GGUF:UD-Q4_K_XL" ;
} ;
} ;
sglang = {
enable = true ;
models . "qwen3-coder" = {
port = 19001 ;
model = "Qwen/Qwen3-8B" ;
settings . reasoning-parser = "qwen3" ;
} ;
} ;
vllm = {
enable = true ;
models . "llama-3-8b" = {
port = 20001 ;
model = "meta-llama/Meta-Llama-3-8B-Instruct" ;
} ;
} ;
} ;
See the options reference for the full list of per-backend options.
The generated config file lives in the world-readable Nix store, so secrets should never be placed in services.llmhop.settings directly.
Instead, reference them via ${file:...} and hand the files to the service with systemd's LoadCredential= .
The right-hand side of each LoadCredential entry is just a file path, so anything that produces a file works: agenix or sops-nix outputs, a manually-managed file under /etc/llmhop/ , or a path emitted by your own secret-provisioning tool.
services . llmhop . settings = {
authTokens = [ " \ ${file:client_token}" ] ;
models . "openai-gpt-4o" = {
url = "https://api.openai.com" ;
headers . Authorization = "Bearer \ ${env:OPENAI_KEY}" ;
} ;
} ;
systemd . services . llmhop . serviceConfig = {
LoadCredential = [ "client_token:/etc/llmhop/client-token" ] ;
EnvironmentFile = [ "/etc/llmhop/openai.env" ] ;
} ;
/etc/llmhop/openai.env is a plain KEY=VALUE file:
OPENAI_KEY = sk-...
${file:...} references are resolved against $CREDENTIALS_DIRECTORY , which systemd exposes as a per-unit tmpfs accessible only to this service, compatible with DynamicUser and the rest of the sandbox.
${env:...} picks up anything the unit inherits, typically via EnvironmentFile= .
Pick whichever matches how your secret tooling hands you the data; mixing both in one config is fine.
Tiny, stateless Go router that dispatches OpenAI-compatible requests to single-model vLLM and sglang backends with zero external dependencies
mirkolenz.github.io/llmhop/
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
6
v1.2.2
Latest
May 29, 2026
+ 5 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
