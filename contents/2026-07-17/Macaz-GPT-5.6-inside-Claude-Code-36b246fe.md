---
source: "https://github.com/macaz-dev/macaz-cli"
hn_url: "https://news.ycombinator.com/item?id=48949316"
title: "Macaz – GPT 5.6 inside Claude Code"
article_title: "GitHub - macaz-dev/macaz-cli: Your favorite models ❤️ your favorite coding agents. · GitHub"
author: "dumitruflorin"
captured_at: "2026-07-17T17:00:12Z"
capture_tool: "hn-digest"
hn_id: 48949316
score: 1
comments: 0
posted_at: "2026-07-17T16:36:39Z"
tags:
  - hacker-news
  - translated
---

# Macaz – GPT 5.6 inside Claude Code

- HN: [48949316](https://news.ycombinator.com/item?id=48949316)
- Source: [github.com](https://github.com/macaz-dev/macaz-cli)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T16:36:39Z

## Translation

Title: Macaz – GPT 5.6 inside Claude Code
記事のタイトル: GitHub - macaz-dev/macaz-cli: お気に入りのモデル ❤️ お気に入りのコーディング エージェント。 · GitHub
説明: お気に入りのモデル ❤️ お気に入りのコーディング エージェント。 - macaz-dev/macaz-cli

記事本文:
GitHub - macaz-dev/macaz-cli: お気に入りのモデル ❤️ お気に入りのコーディング エージェント。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
マカズ開発
/
マカズクリ
公共
通知
「いいえ」を変更するにはサインインする必要があります

ティフィケーション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github .github cmd/ macaz cmd/ macaz 内部 内部スクリプト scripts .gitignore .gitignore LEGAL.md LEGAL.md LICENSE ライセンス通知 NOTICE PRIVACY.md PRIVACY.md README.md README.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
お気に入りのコーディング エージェントでお気に入りのモデルとプロバイダーを使用します。
macaz は、ローカルにインストールされたコーディング エージェントを、によって選択されたモデル プロバイダーに接続します。
ユーザー。 Claude Code または Codex CLI を開始し、モデルの対話をルーティングします。
選択したエージェントは引き続きローカル ツール、権限、セッションを所有します。
スキルもユーザーインターフェースも。
macaz は、Apache-2.0 に基づいてライセンス供与された無料のオープンソース ソフトウェアです。すべてのビルド
完全な機能セットが含まれています。
macaz は独立した相互運用性プロジェクトです。とは提携していませんが、
Anthropic、OpenAI、またはその他の団体によって承認、承認、または後援されている
クライアントまたはモデルプロバイダー。 Claude Code と Codex CLI は別個の製品です。
認可された情報源から入手する必要があります。各製品とプロバイダー
には、引き続き独自の現在の条件が適用されます。
ご使用前にLEGAL.mdとPRIVACY.mdをお読みください。
macaz claude # 選択したプロバイダーを介したクロード コード
macaz codex # 選択したプロバイダーを介した Codex CLI
2 つのクライアントは独立したプロバイダー/モデル構成を持ち、分離されています。
マカズのプロフィール。たとえば、macaz claude は OpenAI モデルを使用できますが、
macaz codex は公式の Anthropic API を使用します。 macaz リセット コーデックスの変更
クロード設定も共有プロバイダー資格情報もありません。
ローカル ルーティング層は両方向を変換します。
Claude Code によって使用される Anthropic Messages リクエストとストリーム。
○

PenAI Responses リクエストと Codex CLI によって使用されるストリーム。
テキスト、画像、サポートされているドキュメント入力、推論の労力、使用法、エラー、
関数ツール、Codex カスタム/自由形式ツール、およびツールの名前空間。そして
ライブ プロバイダー モデルは、各クライアントのネイティブ /model インターフェイスにカタログされます。
クライアントで実行されるツールはクライアント内に残ります。シェルコマンド、ファイル編集、
apply_patch 、スキル、フック、MCP ツール、名前空間ツール、およびサブエージェントはサポートされません。
macaz またはアップストリーム CLI アダプターによって実行されます。通常のクライアント権限と
サンドボックスの動作は、ユーザーがそのクライアントの動作を明示的に渡さない限り保持されます。
独自のバイパスオプション。
サーバー実行ツールは異なります: プロバイダーホスト型 Web などのツール
検索は、選択したプロバイダーが互換性のあるプロバイダーを公開している場合にのみ使用できます。
サーバーの実装。 macaz は Codex の OpenAI がホストする Web 検索リクエストを無効にします
黙って別のプロバイダーのふりをするのではなく、プロバイダー中立のセッション用
それを実行することができます。ローカル関数、カスタム、名前空間、および MCP ツールはそのまま残ります
利用可能です。
いかなる翻訳も、異なる独自のモデルを意味的に同一にすることはできません。
互換性のターゲットは、正しいプロトコルとローカル ツールの動作です。
隠されたプロバイダーのフォールバック。
カール -fsSL https://raw.githubusercontent.com/macaz-dev/macaz-cli/main/scripts/install.sh |しー
インストーラはプラットフォーム リリースを選択し、そのエントリを必要とします。
SHA256SUMS は、SHA-256 を検証し、デフォルトで $HOME/.local/bin にインストールします。
--version を使用してリリースをピン留めします。
カール -fsSL https://raw.githubusercontent.com/macaz-dev/macaz-cli/main/scripts/install.sh |
sh -s -- --バージョン 0.2.0
Windows バイナリおよびその他のリリース ファイルは、次の場所で公開されています。
GitHub リリース 。
後で次のコマンドを使用して、最新のリリースを所定の場所にインストールします。
マカズアップデート
アップデーターは、現在のプラットフォームの正確なアセットのみをダウンロードし、検証します。
リリースチェックサムと埋め込み

バージョンを確認し、実行可能ファイルを次のように置き換えます。
ロールバック保護。プロンプト、ソース コード、構成、またはデータは送信されません。
プロバイダーの資格情報。
リリース ビルドは、起動時に短いベストエフォート型の更新チェックを実行します。小切手
自動的にインストールされることはなく、オフラインでの使用もブロックされません。次のコマンドで無効にします。
エクスポート MACAZ_NO_UPDATE_CHECK=1
要件
Go はソースからビルドする場合にのみ必要です。
macaz claude を使用するには、claude が PATH 上に存在する必要があります。
macaz codex を使用するには、 codex が PATH 上に存在する必要があります。
opencode は、OpenCode CLI が選択されたプロバイダーである場合にのみ必要です。
OpenAI サブスクリプション認証フローではブラウザーが使用されます。
Claude Code、Codex CLI、および OpenCode はバンドルまたは再配布されていません。
各クライアントの最初の呼び出しでは、独自のセットアップが開きます。
マカズ・クロード
マカズコーデックス
利用可能なアップストリーム:
API キーと OAuth 認証情報はオペレーティング システムの認証情報に保存されます
config.json ではなくストアします。 Anthropic オプションは Anthropic API キーを使用し、
パブリックメッセージ API。 Claude コンシューマを使用または変換しません
サブスクリプション。人間モデル ID、トークン制限、入力機能、および労力
レベルは、セットアップ時と起動時にアカウントのライブ モデル API から読み取られます。
起動するたびに、アクティブなプロバイダー カタログが更新されます。結果として得られる公開モデル
ID とサポートされる推論レベルは、分離されたクライアント プロファイルに書き込まれます。
したがって、モデルの選択は、Claude Code または Codex CLI のネイティブ /model を通じて機能します。
インターフェース。構成されたプロバイダーのデフォルトは起動時に選択されます。インタラクティブ
変更は、選択したクライアントの通常のセッション動作に従います。
macaz claude [args...] クロードコードの開始
macaz codex [args...] Codex CLI の起動
macaz status [client] 設定されたプロバイダーとモデルカタログを確認します
macaz ドクター [クライアント] クライアントの実行可能ファイルとプロバイダーを確認する
macaz replace [client] 1 つのクライアントまたはすべての Mac をリセットします

省略時の az 状態
macaz 法的 互換性に関する通知を表示する
macaz update 最新の検証済みリリースをインストールします
macaz バージョン インストールされているバージョンを表示します
macaz claude または macaz codex の後の引数がそのクライアントに転送されます。
macaz の認証をバイパスするモデル/プロバイダー/プロファイルのオーバーライドを除く
ローカルルーティング層。
macaz クロード --dangerously-skip-permissions
macaz codex --sandbox workspace-write
マカズステータスコーデックス
マカズ医師クロード
マカズリセットコーデックス
macazリセットクロードまたはmacazリセットコーデックスは、そのクライアントのmacazのみを削除します
構成と分離されたプロファイル。他のクライアントと共有を保持します
プロバイダーの資格情報。クライアントなしで macaz をリセットすると、両方の macaz プロファイルが削除されます。
macaz が管理するすべての認証情報、設定、および分離されたセッション履歴。それ
通常のベンダー/クライアント プロファイルや直接管理されている資格情報は削除されません
ベンダー CLI による。
ローカル ルーティング層は、ランダムな 127.0.0.1 ポートで待機します。
起動ごとに新しいランダム認証トークンが生成されます。
トークンは子プロセス環境にのみ存在します。
構成ファイルとプロファイル ファイルは、サポートされている場合はプライベート アクセス許可を使用します。
通常のクロードおよびコーデックスのプロファイルは変更されません。
ローカル ルーティング プロセスは、起動されたクライアントが終了すると停止します。
別のプロバイダーまたは公式モデルへのフォールバックはありません。
macaz には、プロジェクトで運用される分析クライアントやテレメトリ クライアントは含まれていません。
分離されたプロファイルには、クライアント所有のセッション履歴が含まれる場合があります。参照
正確なデータ フローと削除動作については PRIVACY.md を参照してください。
OpenAI API と OpenRouter は、Responses 互換の HTTP アダプターを使用します。
OpenAI サブスクリプションはデバイス認証を使用し、資格情報を更新し、
制限付きアカウントの再試行/バックオフを適用します。実験的なものなので残っています
OpenAI の当座預金条件が適用されます。
Anthropic API はネイティブの /v1/models 、 /v1/messages 、および
/v1/私

ssages/count_tokens 。応答カスタム ツールは通常のものに変換されます
型指定された生入力ラッパーを備え、前に変換されて戻される人造ツール
Codex がそれらを実行します。
Codex CLI はプロバイダーとして codex アプリサーバーを使用し、
クロードのクライアント。 Codex をクライアントとアップストリームの両方として使用すると、再帰的に発生します。
OpenCode CLI は、分離されたリクエストスコープのプロバイダー構成を使用します。それ
プロジェクト ツールとコンテキストは、第 2 エージェント層として公開されません。
Go 1.26.5 以降が必要です。
./cmd/macaz バージョンを実行します
./cmd/macaz クロードを実行してください
./cmd/macaz codex を実行します。
ローカル バイナリをビルドして実行します。
go build -o ./macaz ./cmd/macaz
./macaz バージョン
./macaz コーデックス
検証を実行します。
モッド検証に行く
テストに行ってください。/...
獣医に行ってください。/...
テストに行く - レース ./...
完全なローカル リリース パッケージを無視された dist/ にビルドします。
./scripts/build-release.sh v1.0.0
リリース スクリプトには空の出力ディレクトリが必要で、静的なディレクトリが作成されます。
amd64 および arm64 用の macOS、Linux、および Windows バイナリ、チェックサム、
インストーラー。別の空のディレクトリを選択するには:
OUTPUT_DIR=/tmp/macaz-release ./scripts/build-release.sh v1.0.0
オプトイン ライブ プロバイダー テストは、通常の CI では実行されません。
MACAZ_OPENAI_API_INTEGRATION_MODEL= <モデル> go test ./internal/provider/openai -run LiveOpenAIAPI
MACAZ_OPENAI_SUBSCRIPTION_INTEGRATION_MODEL= <モデル> go test ./internal/provider/openai -run LiveOpenAISubscription
MACAZ_OPENROUTER_INTEGRATION_MODEL= <プロバイダ/モデル> go test ./internal/provider/openrouter -run LiveOpenRouter
MACAZ_CODEX_INTEGRATION_EXECUTABLE= < codex または互換性 > go test ./internal/provider/codexcli -run LiveCodex
MACAZ_OPENCODE_INTEGRATION_MODEL= <プロバイダ/モデル> go test ./internal/provider/opencodecli -run LiveOpenCode
MACAZ_CLAUDE_INTEGRATION=1 テストに行く ./internal/app -run LiveClaudeLifecycle
メインでマージされたすべてのプル リクエストは、

検証され、次の SemVer タグが割り当てられ、
サポートされているすべてのプラットフォーム向けに構築され、GitHub リリースとして公開されています。
チェックサム、来歴、生成されたリリース ノート。デフォルトのバンプはパッチです。
release:minor ラベルと release:major ラベルでは、より大きなバンプが選択されます。
Apache ライセンス 2.0。 「ライセンス」、「通知」、および「」を参照してください。
THIRD_PARTY_NOTICES.md 。
お気に入りのモデル ❤️ お気に入りのコーディング エージェント。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
3
マカズ0.2.0
最新の
2026 年 7 月 17 日
+ 2 リリース
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Your favorite models ❤️ your favorite coding agents. - macaz-dev/macaz-cli

GitHub - macaz-dev/macaz-cli: Your favorite models ❤️ your favorite coding agents. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
macaz-dev
/
macaz-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github .github cmd/ macaz cmd/ macaz internal internal scripts scripts .gitignore .gitignore LEGAL.md LEGAL.md LICENSE LICENSE NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Use your favorite models and providers with your favorite coding agents.
macaz connects locally installed coding agents with a model provider chosen by
the user. It starts Claude Code or Codex CLI and routes model interactions while
the selected agent continues to own its local tools, permissions, sessions,
skills, and user interface.
macaz is free and open-source software licensed under Apache-2.0. Every build
contains the complete feature set.
macaz is an independent interoperability project. It is not affiliated with,
authorized by, endorsed by, or sponsored by Anthropic, OpenAI, or any other
client or model provider. Claude Code and Codex CLI are separate products that
must be obtained from their authorized sources. Each product and provider
remains subject to its own current terms.
Read LEGAL.md and PRIVACY.md before use.
macaz claude # Claude Code through the selected provider
macaz codex # Codex CLI through the selected provider
The two clients have independent provider/model configuration and isolated
macaz profiles. For example, macaz claude can use an OpenAI model while
macaz codex uses the official Anthropic API. macaz reset codex changes
neither the Claude configuration nor shared provider credentials.
The local routing layer translates both directions:
Anthropic Messages requests and streams used by Claude Code;
OpenAI Responses requests and streams used by Codex CLI;
text, images, supported document inputs, reasoning effort, usage, errors,
function tools, Codex custom/free-form tools, and tool namespaces; and
live provider model catalogs into each client's native /model interface.
Client-executed tools stay in the client. Shell commands, file edits,
apply_patch , skills, hooks, MCP tools, namespaced tools, and subagents are not
executed by macaz or by an upstream CLI adapter. Normal client permission and
sandbox behavior is preserved unless the user explicitly passes that client's
own bypass option.
Server-executed tools are different: a tool such as provider-hosted web
search is available only when the selected provider exposes a compatible
server implementation. macaz disables Codex's OpenAI-hosted web-search request
for provider-neutral sessions rather than silently pretending another provider
can execute it. Local function, custom, namespace, and MCP tools remain
available.
No translation can make different proprietary models semantically identical.
The compatibility target is correct protocol and local-tool behavior without a
hidden provider fallback.
curl -fsSL https://raw.githubusercontent.com/macaz-dev/macaz-cli/main/scripts/install.sh | sh
The installer selects the platform release, requires its entry in
SHA256SUMS , verifies SHA-256, and installs to $HOME/.local/bin by default.
Pin a release with --version :
curl -fsSL https://raw.githubusercontent.com/macaz-dev/macaz-cli/main/scripts/install.sh |
sh -s -- --version 0.2.0
Windows binaries and other release files are published on
GitHub Releases .
Install the latest release in place later with:
macaz update
The updater downloads only the exact asset for the current platform, verifies
the release checksums and embedded version, and replaces the executable with
rollback protection. It does not send prompts, source code, configuration, or
provider credentials.
Release builds perform a short best-effort update check on startup. The check
never installs automatically and does not block offline use. Disable it with:
export MACAZ_NO_UPDATE_CHECK=1
Requirements
Go is needed only when building from source.
claude must be on PATH to use macaz claude .
codex must be on PATH to use macaz codex .
opencode is required only when OpenCode CLI is the selected provider.
A browser is used by the OpenAI Subscription authorization flow.
Claude Code, Codex CLI, and OpenCode are not bundled or redistributed.
The first invocation of each client opens its own setup:
macaz claude
macaz codex
Available upstreams:
API keys and OAuth credentials are stored in the operating-system credential
store, not in config.json . The Anthropic option uses an Anthropic API key and
the public Messages API; it does not use or convert a Claude consumer
subscription. Anthropic model IDs, token limits, input capabilities, and effort
levels are read from the account's live Models API during setup and startup.
Each start refreshes the active provider catalog. The resulting public model
IDs and supported reasoning levels are written to the isolated client profile,
so model selection works through Claude Code's or Codex CLI's native /model
interface. The configured provider default is selected at launch; interactive
changes follow the selected client's normal session behavior.
macaz claude [args...] Start Claude Code
macaz codex [args...] Start Codex CLI
macaz status [client] Check the configured provider and model catalog
macaz doctor [client] Check the client executable and provider
macaz reset [client] Reset one client, or all macaz state when omitted
macaz legal Show the compatibility notice
macaz update Install the latest verified release
macaz version Show the installed version
Arguments after macaz claude or macaz codex are forwarded to that client,
except model/provider/profile overrides that would bypass macaz's authenticated
local routing layer.
macaz claude --dangerously-skip-permissions
macaz codex --sandbox workspace-write
macaz status codex
macaz doctor claude
macaz reset codex
macaz reset claude or macaz reset codex removes only that client's macaz
configuration and isolated profile. It preserves the other client and shared
provider credentials. macaz reset with no client removes both macaz profiles,
all macaz-managed credentials, configuration, and isolated session history. It
does not remove normal vendor-client profiles or credentials managed directly
by vendor CLIs.
The local routing layer listens on a random 127.0.0.1 port.
A new random authentication token is generated for every launch.
The token exists only in the child process environment.
Config and profile files use private permissions where supported.
Normal Claude and Codex profiles are not modified.
The local routing process stops when the launched client exits.
There is no fallback to a different provider or official model.
macaz contains no project-operated analytics or telemetry client.
The isolated profiles may contain client-owned session history. See
PRIVACY.md for exact data flow and deletion behavior.
OpenAI API and OpenRouter use Responses-compatible HTTP adapters.
OpenAI Subscription uses device authorization, refreshes credentials, and
applies bounded account retry/backoff. It is experimental and remains
subject to OpenAI's current account terms.
Anthropic API uses native /v1/models , /v1/messages , and
/v1/messages/count_tokens . Responses custom tools are converted to normal
Anthropic tools with a typed raw-input wrapper and converted back before
Codex executes them.
Codex CLI as a provider uses codex app-server and is offered only to the
Claude client; using Codex as both client and upstream would recurse.
OpenCode CLI uses an isolated request-scoped provider configuration. Its
project tools and context are not exposed as a second agent layer.
Go 1.26.5 or newer is required.
go run ./cmd/macaz version
go run ./cmd/macaz claude
go run ./cmd/macaz codex
Build and run a local binary:
go build -o ./macaz ./cmd/macaz
./macaz version
./macaz codex
Run verification:
go mod verify
go test ./...
go vet ./...
go test -race ./...
Build a complete local release package into ignored dist/ :
./scripts/build-release.sh v1.0.0
The release script requires an empty output directory and creates static
macOS, Linux, and Windows binaries for amd64 and arm64, plus checksums and the
installer. To choose another empty directory:
OUTPUT_DIR=/tmp/macaz-release ./scripts/build-release.sh v1.0.0
Opt-in live-provider tests never run in normal CI:
MACAZ_OPENAI_API_INTEGRATION_MODEL= < model > go test ./internal/provider/openai -run LiveOpenAIAPI
MACAZ_OPENAI_SUBSCRIPTION_INTEGRATION_MODEL= < model > go test ./internal/provider/openai -run LiveOpenAISubscription
MACAZ_OPENROUTER_INTEGRATION_MODEL= < provider/model > go test ./internal/provider/openrouter -run LiveOpenRouter
MACAZ_CODEX_INTEGRATION_EXECUTABLE= < codex-or-compatible > go test ./internal/provider/codexcli -run LiveCodex
MACAZ_OPENCODE_INTEGRATION_MODEL= < provider/model > go test ./internal/provider/opencodecli -run LiveOpenCode
MACAZ_CLAUDE_INTEGRATION=1 go test ./internal/app -run LiveClaudeLifecycle
Every merged pull request on main is verified, assigned the next SemVer tag,
built for all supported platforms, and published as a GitHub Release with
checksums, provenance, and generated release notes. The default bump is patch;
the release:minor and release:major labels select larger bumps.
Apache License 2.0. See LICENSE , NOTICE , and
THIRD_PARTY_NOTICES.md .
Your favorite models ❤️ your favorite coding agents.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
3
macaz 0.2.0
Latest
Jul 17, 2026
+ 2 releases
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
