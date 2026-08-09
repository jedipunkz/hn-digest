---
source: "https://github.com/ademisler/RunOnMine"
hn_url: "https://news.ycombinator.com/item?id=49233245"
title: "Show HN: RunOnMine – local policy and approvals for AI access to your machine"
article_title: "GitHub - ademisler/RunOnMine: Local-first MCP security gateway for controlled AI access to files, terminals, browsers and desktop apps on machines you own. · GitHub"
author: "ademisler"
captured_at: "2026-08-09T17:22:02Z"
capture_tool: "hn-digest"
hn_id: 49233245
score: 2
comments: 0
posted_at: "2026-08-09T17:05:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RunOnMine – local policy and approvals for AI access to your machine

- HN: [49233245](https://news.ycombinator.com/item?id=49233245)
- Source: [github.com](https://github.com/ademisler/RunOnMine)
- Score: 2
- Comments: 0
- Posted: 2026-08-09T17:05:08Z

## Translation

タイトル: HN を表示: RunOnMine – マシンへの AI アクセスに関するローカル ポリシーと承認
記事のタイトル: GitHub - ademisler/RunOnMine: 所有するマシン上のファイル、端末、ブラウザ、デスクトップ アプリへの制御された AI アクセスのためのローカルファースト MCP セキュリティ ゲートウェイ。 · GitHub
説明: 所有するマシン上のファイル、端末、ブラウザ、デスクトップ アプリへの制御された AI アクセスのためのローカル ファースト MCP セキュリティ ゲートウェイ。 - ademisler/RunOnMine

記事本文:
GitHub - ademisler/RunOnMine: 所有するマシン上のファイル、端末、ブラウザ、デスクトップ アプリへの制御された AI アクセスのためのローカルファースト MCP セキュリティ ゲートウェイ。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アデミスラー
/
ランオンマイン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
307 コミット 307 コミット .cargo .cargo .github .github 受け入れ受け入れ apps apps crates crates docs docs fuzz fuzz パッケージング package

g release リリース スクリプト スクリプト xtask xtask .gitattributes .gitattributes .gitignore .gitignore .gitleaksignore .gitleaksignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md Task.md Task.md Deny.toml Deny.toml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AIを働かせましょう。制御をローカルに保ちます。
ローカルファーストの Model Context Protocol (MCP) ゲートウェイおよび macOS、Linux、Windows 用のデスクトップ コントロール センター。
ウェブサイト ·
インストール・
クイックスタート ·
仕組み ·
セキュリティ ·
ドキュメント
RunOnMine は、AI アシスタントとユーザーが所有するマシンの間に位置します。それはAIに与えます
ファイル、端末、ブラウザ、デスクトップ アプリケーションへのアクセスを制御しながら、
実行、ポリシー、認証情報、承認、監査記録を
機械。
SSH、生のシェル、またはパブリック MCP リスナーを公開する代わりに、RunOnMine チェックが行われます。
誰が尋ねているか、どのツールが呼び出されているか、どのリソースが対象であるか、
何かを実行する前にアクションがローカルの承認を必要とするかどうか。
能力
RunOnMine 境界
ファイル
AI アクセスは、明示的に選択したディレクトリに制限されます。
ターミナルとプロセス
コマンドはポリシーによって評価され、正確なローカル承認が必要になる場合があります。
ブラウザの自動化
ネットワークおよびプライベート アドレス保護を備えた分離されたブラウザ プロファイルを使用します。
デスクトップコントロール
ネイティブ デスクトップ アクションは、OS/セッションが許可する場合にのみ実行されます。
リモート接続
管理されたトンネルは外部に接続します。 MCP サーバーはループバック状態のままです。
安全管理
ローカル承認、明示的な拒否ルール、改ざん防止監査、診断、および緊急ロック。
AI モデルは、選択した AI サービスで引き続き実行されます。ツール例

処刑が起こる
あなたが制御するマシン上の RunOnMine 境界を通過します。
AI クライアントは、構成されたコネクタを通じて MCP ツール リクエストを送信します。
RunOnMine は、コネクタ/リクエスタ ID、選択されたルート、ポリシー、
リソースのスコープ、および正確なアクションの承認。
アクションは許可されるか、ローカル承認のために保留されるか、または拒否されます。
結果は AI に返され、ローカル監査証跡が更新されます。
RunOnMine はローカル stdio、オプトイン認証ループバック HTTP、Cloudflare をサポートします
コネクタ、および OpenAI Secure MCP トンネル。参照
フルコネクタモデルの接続モード。
RunOnMine はプレリリース ソフトウェアです。最新のパブリックベータ版は、
v0.1.0-ベータ.1 。
macOS はアドホック署名されていますが、開発者 ID の署名または公証はされていません。
Windows インストーラーは Authenticode 署名されていません。 Gatekeeper または SmartScreen は、
したがって、ダウンロードされたビルドについて警告します。各プライマリ アーティファクトには隣接する
SHA-256 チェックサム、およびターゲット固有の CycloneDX SBOM は、
解放する。
ポータブルな署名されていないアーカイブ、チェックサム、および SBOM は、
リリースページ
ソース ビルドは、サポート対象外の開発者およびプラットフォームに対して引き続きサポートされます。
パッケージ化されたベータ版。
runonmine-desktop_0.1.0-beta.1_x64-setup.exe を実行します。現在のユーザーです
インストールします。オプションの LocalSystem ヘルパーは、次の場合を除き、インストールまたはアクティブ化されません。
後で明示的にリクエストします。このベータ版は Authenticode 署名されていないため、
Windows SmartScreen では、認識されない発行元に関する警告が表示される場合があります。
インストール後、AI がアクセスできるプロジェクト ディレクトリを初期化します。
$rom = 結合パス $ env: LOCALAPPDATA " RunOnMine\runonmine.exe "
& $rom setup -- root " C:\path\to\your\project "
& $rom サービスのインストール
[スタート] メニューから RunOnMine を起動して、セキュリティ コントロール センターを開きます。
パッケージのライフサイクルについては Windows ガイドを参照してください。
アンインストール動作、およびオプションの特権

脚付きのヘルパー。
runonmine-desktop_0.1.0-beta.1_amd64.deb をダウンロードし、次のようにインストールします。
sudo apt install ./runonmine-desktop_0.1.0-beta.1_amd64.deb
runonmine setup --root "$HOME/Projects/my-project"
runonmine サービスのインストール
runonmine-デスクトップ
ヘッドレス x86_64 または ARM64 システムの場合は、対応する runonmine_* DEB を使用します。
ヘッドレス サービスには明示的な安全なシークレット ストアのセットアップが必要なので、次の手順に従ってください。
デスクトップをコピーするのではなく、Linux と VPS のガイド
サーバーに対するサービスの前提条件。
RunOnMine_0.1.0-beta.1_universal.dmg を開き、RunOnMine.app を次の場所にコピーします。
アプリケーション。ベータ版では、検証済みのアドホック コード署名と強化されたコード署名が使用されています。
ランタイムですが、Apple Developer ID が署名または公証されていないため、Gatekeeper は
隔離されたダウンロードに対して警告またはブロックします。
RunOnMine は macOS の同意プロンプトをバイパスしません。デスクトップ入力/キャプチャの可能性があります
アクセシビリティまたは画面録画の許可が必要です。を参照してください。
macOS ガイド 。
1. AI が使用できるディレクトリのみを選択します
runonmine setup --root /absolute/path/to/project
より広範なアクセスが本当に目的でない限り、ホーム ディレクトリ全体を選択しないでください。
必須です。
runonmine ポリシー ショー
新しいコネクタは Safe で始まります: 読み取りは可能ですが、書き込み/実行は要求されます
ローカルで実行され、管理者の実行は拒否されます。開発者向けの対象
信頼できる選択されたルートのコーディング作業。自動化 (CLI では完全) は、
最も広範なローカル プリセットであり、専用または限定された範囲でのみ使用する必要があります。
機械。
1 回限りのフォアグラウンド セッションの場合:
runonmine エージェントの実行
または、通常のユーザーごとのサービスをインストールして、RunOnMine がログインで回復できるようにします
セッション:
runonmine サービスのインストール
runonmine サービスのステータス
4. AI クライアントを接続する
最小のローカル サーフェスはデフォルトの stdio コネクタです。
runonmine接続リスト
runonmine mcp stdio --connector <ローカルコネクタ ID>
認証されたループバック HTTP はオプトインであり、その

ベアラートークンは決して出力されません:
runonmine 接続ローカル http を有効にする \
--token-output /absolute/private/local-http.json
リモート アクセスの場合は、代わりにマネージド Cloudflare または OpenAI コネクタを使用します。
MCP リスナーをネットワークにオープンします。 「接続モード」を参照してください。
タスク
コマンド/UI
保留中のアクションを確認する
デスクトップ アプリまたは runonmine 承認リストでの承認
一度承認する
runonmine 承認 <id> を承認します --1 回
検査ポリシー
runonmine ポリシー ショー
健康状態をチェックする
ルノンマイン博士
編集されたサポート バンドルを作成する
runonmine support-bundle --output runonmine-support.zip
すぐにアクセスを停止してください
ルノンマインロック
runonmine ロックはエージェントと管理対象コネクタを停止し、キューに入れられたコネクタを拒否します
承認、ライブ OAuth セッションの取り消し、一時コネクタの無効化
資格情報。
なぜ直接 SSH または生の MCP サーバーを使用しないのでしょうか?
マシンへの直接アクセス
ランオンマイン
幅広いアカウント権限
機能およびリソースを対象としたポリシー
多くの場合、1 つの認証情報ですべてのロックが解除されます
コネクタ/リクエスタ ID はアクションごとに評価されます
人間によるチェックポイントが組み込まれていない
危険な行為には現地の正確な承認が必要となる場合があります
パブリック リスナーまたは受信ファイアウォール ルール
MCP はループバック状態のままです。管理されたトンネルは外部に接続します
アドホックログ
改ざん防止監査と緊急ロック
セキュリティの概要
RunOnMine は、マシンの境界を見せかけるのではなく、見えるようにするように設計されています。
機械の自動化は無害です。
MCP がパブリック ネットワーク インターフェイスに直接バインドされることはありません。
リモート コネクタは、自身の危険な要求を承認できません。
リモート管理者の実行は、バイパス不可能な安全上限によって拒否されます。
シェル プロセスはクリアされた環境から開始されます。 Windows PowerShell プロファイルが無効になっています。
オプションの特権ヘルパーはデフォルトでは存在しないため、別途必要になります。
明示的なインストール。
ファイルシステム ツールは明示的に選択されたルート内で動作します

。
ブラウザの自動化では、ユーザーの毎日のプロファイルではなく、分離されたプロファイルが使用されます。
ブラウザのプロファイル。
シークレットはオペレーティング システムの資格情報ストアまたは文書化された場所に保存されます。
暗号化されたヘッドレスフォールバック。
監査レコードは改ざんが明らかであり、生の秘密やコマンドの保存を回避します。
ペイロード。
シェル、ブラウザ、デスクトップ、および特権ツールは、破壊的または外部的なものを作成する可能性があります。
変化します。 RunOnMine はセキュリティ境界および承認システムであり、サンドボックスではありません。
エージェントには、実際に必要なアカウント権限と選択されたルートのみを与えます。
権限モデルを読んで、
脅威モデル、および
広範囲の書き込みを有効にする前にブラウザのセキュリティを設定するか、
実行能力。
runonmine — セットアップ、コネクタ、ポリシー、承認、サービス ライフサイクル、
診断、監査、緊急ロック。
runonmine-agent — ローカル MCP サーバーおよびコネクタ スーパーバイザ。
runonmine-desktop — 承認、接続、
権限、OAuth、監査、診断。
runonmine-helper — オプションで別途インストールされる特権ヘルパー。
RunOnMine はまだプレリリース ソフトウェアです。リポジトリにはリリース状態が記録されます。
機械可読形式: accept/release-candidate.toml は凍結されたものを識別します。
ソース候補と accept/release-gates.toml はどの受け入れを記録するか
そして実際にセキュリティゲートを通過しました。
その制限が記載されている場合、パブリック ベータ版は署名なしで配布される場合があります
目立つように表示され、必要なパブリック ベータ ゲートはすべて通過します。署名されていないアーティファクトは、
発行者の ID またはオペレーティング システムの信頼を確立してはなりません。
実稼働環境で署名されたビルドとして提示されます。
ドキュメントの索引から始めます。
始めましょう: 安全なオンボーディングと
トラブルシューティング
プラットフォーム: macOS 、
Linux/VPS、および
窓
アーキテクチャ: アーキテクチャ、
接続モード、および MCP ツール
セキュリティ: アクセス許可、
脅威モデル 、
監査の整合性、および
特権的な

ヘルパー
品質とリリース: テスト、
リリース受付、および
リリースプロセス
ヘルプ: サポートとセキュリティレポート
RunOnMine は、 Rust 1.95.0 に固定された Rust ワークスペースです。 Cargo.lock を維持する
コミットした。プル リクエストを開く前に、次を実行します。
python3 スクリプト/ci/check-docs.py
カーゴ実行 --locked -p xtask -- verify
サポートされているヘッドレス Linux 開発ホストでは、 --headless を使用します。参照
COTRIBUTING.md (セキュリティ重視の変更要件)
そして完全な貢献者のワークフロー。
Apache ライセンス 2.0。 「ライセンス」を参照してください。
所有するマシン上のファイル、端末、ブラウザ、デスクトップ アプリへの制御された AI アクセスのためのローカル ファースト MCP セキュリティ ゲートウェイ。
Readme Apache-2.0 ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local-first MCP security gateway for controlled AI access to files, terminals, browsers and desktop apps on machines you own. - ademisler/RunOnMine

GitHub - ademisler/RunOnMine: Local-first MCP security gateway for controlled AI access to files, terminals, browsers and desktop apps on machines you own. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
ademisler
/
RunOnMine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
307 Commits 307 Commits .cargo .cargo .github .github acceptance acceptance apps apps crates crates docs docs fuzz fuzz packaging packaging release release scripts scripts xtask xtask .gitattributes .gitattributes .gitignore .gitignore .gitleaksignore .gitleaksignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md Task.md Task.md deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Let AI work. Keep control local.
A local-first Model Context Protocol (MCP) gateway and desktop control center for macOS, Linux, and Windows.
Website ·
Install ·
Quick start ·
How it works ·
Security ·
Documentation
RunOnMine sits between an AI assistant and a machine you own. It gives the AI
controlled access to files, terminals, browsers, and desktop applications while
keeping execution, policy, credentials, approvals, and audit records on your
machine.
Instead of exposing SSH, a raw shell, or a public MCP listener, RunOnMine checks
who is asking, what tool is being called, which resource it targets, and
whether the action needs local approval before anything executes.
Capability
RunOnMine boundary
Files
AI access is limited to directories you explicitly select.
Terminal & processes
Commands are evaluated by policy and can require exact local approval.
Browser automation
Uses an isolated browser profile with network and private-address protections.
Desktop control
Native desktop actions run only when the OS/session permits them.
Remote connectivity
Managed tunnels connect outward; the MCP server stays on loopback.
Safety controls
Local approvals, explicit deny rules, tamper-evident audit, diagnostics, and Emergency Lock.
The AI model still runs in the AI service you choose. Tool execution happens
through the RunOnMine boundary on the machine you control.
An AI client sends an MCP tool request through a configured connector.
RunOnMine evaluates connector/requester identity, selected roots, policy,
resource scope, and any exact-action approval.
The action is allowed , held for local approval , or denied .
The result is returned to the AI and the local audit trail is updated.
RunOnMine supports local stdio, opt-in authenticated loopback HTTP, Cloudflare
connectors, and OpenAI Secure MCP Tunnel. See
Connection modes for the full connector model.
RunOnMine is pre-release software . The latest public beta is
v0.1.0-beta.1 .
macOS is ad-hoc signed but not Developer ID signed or notarized, and the
Windows installer is not Authenticode signed. Gatekeeper or SmartScreen may
therefore warn on downloaded builds. Each primary artifact has an adjacent
SHA-256 checksum, and target-specific CycloneDX SBOMs are attached to the
release.
Portable unsigned archives, checksums, and SBOMs are available on the
release page .
Source builds remain supported for developers and platforms not covered by the
packaged beta.
Run runonmine-desktop_0.1.0-beta.1_x64-setup.exe . It is a current-user
install ; the optional LocalSystem helper is not installed or activated unless
you explicitly request it later. Because this beta is not Authenticode signed,
Windows SmartScreen may show an unrecognized-publisher warning.
After installation, initialize the project directory the AI may access:
$rom = Join-Path $ env: LOCALAPPDATA " RunOnMine\runonmine.exe "
& $rom setup -- root " C:\path\to\your\project "
& $rom service install
Launch RunOnMine from the Start Menu to open the security control center.
See the Windows guide for package lifecycle,
uninstall behavior, and the optional privileged helper.
Download runonmine-desktop_0.1.0-beta.1_amd64.deb , then install it with:
sudo apt install ./runonmine-desktop_0.1.0-beta.1_amd64.deb
runonmine setup --root "$HOME/Projects/my-project"
runonmine service install
runonmine-desktop
For headless x86_64 or ARM64 systems, use the corresponding runonmine_* DEB.
Headless services need an explicit secure secret-store setup, so follow the
Linux and VPS guide rather than copying desktop
service assumptions to a server.
Open RunOnMine_0.1.0-beta.1_universal.dmg and copy RunOnMine.app to
Applications. The beta uses verified ad-hoc code signatures and hardened
runtime, but it is not Apple Developer ID signed or notarized, so Gatekeeper may
warn or block a quarantined download.
RunOnMine does not bypass macOS consent prompts. Desktop input/capture may
require Accessibility or Screen Recording permission. See the
macOS guide .
1. Select only the directories AI may use
runonmine setup --root /absolute/path/to/project
Do not select your whole home directory unless that broader access is genuinely
required.
runonmine policy show
New connectors start with Safe : reads are available, writes/execution ask
locally, and administrator execution is denied. Developer is intended for
trusted selected-root coding work. Automation ( full in the CLI) is the
broadest local preset and should be used only on a dedicated or tightly scoped
machine.
For a one-off foreground session:
runonmine agent run
Or install the normal per-user service so RunOnMine can recover with your login
session:
runonmine service install
runonmine service status
4. Connect your AI client
The smallest local surface is the default stdio connector:
runonmine connect list
runonmine mcp stdio --connector <local-connector-id>
Authenticated loopback HTTP is opt-in and its bearer token is never printed:
runonmine connect local-http enable \
--token-output /absolute/private/local-http.json
For remote access, use a managed Cloudflare or OpenAI connector instead of
opening the MCP listener to the network. See Connection modes .
Task
Command / UI
Review pending actions
Approvals in the desktop app or runonmine approvals list
Approve once
runonmine approvals approve <id> --once
Inspect policy
runonmine policy show
Check health
runonmine doctor
Create a redacted support bundle
runonmine support-bundle --output runonmine-support.zip
Stop access immediately
runonmine lock
runonmine lock stops the agent and managed connectors, rejects queued
approvals, revokes live OAuth sessions, and invalidates temporary connector
credentials.
Why not direct SSH or a raw MCP server?
Direct machine access
RunOnMine
Broad account authority
Capability- and resource-scoped policy
One credential often unlocks everything
Connector/requester identity is evaluated per action
No built-in human checkpoint
Dangerous actions can require exact local approval
Public listener or inbound firewall rule
MCP remains on loopback; managed tunnels connect outward
Ad-hoc logs
Tamper-evident audit plus Emergency Lock
Security at a glance
RunOnMine is designed to make the machine boundary visible rather than pretend
machine automation is harmless:
MCP is never bound directly to a public network interface.
Remote connectors cannot approve their own dangerous requests.
Remote administrator execution is denied by a non-bypassable safety ceiling.
Shell processes start from a cleared environment; Windows PowerShell profiles are disabled.
The optional privileged helper is absent by default and requires separate,
explicit installation.
Filesystem tools operate inside explicitly selected roots.
Browser automation uses an isolated profile instead of the user’s daily
browser profile.
Secrets stay in the operating-system credential store or the documented
encrypted headless fallback.
Audit records are tamper-evident and avoid storing raw secrets or command
payloads.
Shell, browser, desktop, and privileged tools can make destructive or external
changes. RunOnMine is a security boundary and approval system, not a sandbox .
Give the agent only the account authority and selected roots it actually needs.
Read the Permissions model ,
Threat model , and
Browser security before enabling broader write or
execution capabilities.
runonmine — setup, connectors, policy, approvals, service lifecycle,
diagnostics, audit, and Emergency Lock.
runonmine-agent — local MCP server and connector supervisor.
runonmine-desktop — security control center for approvals, connections,
permissions, OAuth, audit, and diagnostics.
runonmine-helper — optional separately installed privileged helper.
RunOnMine is still pre-release software. The repository records release state in
machine-readable form: acceptance/release-candidate.toml identifies the frozen
source candidate and acceptance/release-gates.toml records which acceptance
and security gates actually passed.
Public beta may be distributed unsigned when that limitation is stated
prominently and every required public-beta gate passes. Unsigned artifacts do
not establish publisher identity or operating-system trust and must not be
presented as production-signed builds.
Start with the documentation index .
Get started: secure onboarding and
troubleshooting
Platforms: macOS ,
Linux/VPS , and
Windows
Architecture: architecture ,
connection modes , and MCP tools
Security: permissions ,
threat model ,
audit integrity , and
privileged helper
Quality & release: testing ,
release acceptance , and
release process
Help: support and security reporting
RunOnMine is a Rust workspace pinned to Rust 1.95.0 . Keep Cargo.lock
committed. Before opening a pull request, run:
python3 scripts/ci/check-docs.py
cargo run --locked -p xtask -- verify
On supported headless Linux development hosts, use --headless . See
CONTRIBUTING.md for security-sensitive change requirements
and the complete contributor workflow.
Apache License 2.0. See LICENSE .
Local-first MCP security gateway for controlled AI access to files, terminals, browsers and desktop apps on machines you own.
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
