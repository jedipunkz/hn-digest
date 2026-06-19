---
source: "https://github.com/SSX360/matrixscroll"
hn_url: "https://news.ycombinator.com/item?id=48598284"
title: "Matrix Scroll – sign AI-generated code changes with Ed25519"
article_title: "GitHub - SSX360/matrixscroll: Open protocol for hardware-signed AI-assisted code (Ed25519 root of trust, software emulator + SSX360 reference device). · GitHub"
author: "ssx360"
captured_at: "2026-06-19T16:05:40Z"
capture_tool: "hn-digest"
hn_id: 48598284
score: 2
comments: 1
posted_at: "2026-06-19T13:23:33Z"
tags:
  - hacker-news
  - translated
---

# Matrix Scroll – sign AI-generated code changes with Ed25519

- HN: [48598284](https://news.ycombinator.com/item?id=48598284)
- Source: [github.com](https://github.com/SSX360/matrixscroll)
- Score: 2
- Comments: 1
- Posted: 2026-06-19T13:23:33Z

## Translation

タイトル: マトリックス スクロール – Ed25519 を使用して AI 生成のコード変更に署名する
記事タイトル: GitHub - SSX360/matrixscroll: ハードウェア署名された AI 支援コード用のオープン プロトコル (Ed25519 信頼のルート、ソフトウェア エミュレーター + SSX360 参照デバイス)。 · GitHub
説明: ハードウェア署名された AI 支援コード用のオープン プロトコル (Ed25519 信頼のルート、ソフトウェア エミュレーター + SSX360 参照デバイス)。 - SSX360/マトリックススクロール

記事本文:
GitHub - SSX360/matrixscroll: ハードウェア署名された AI 支援コード用のオープン プロトコル (Ed25519 信頼のルート、ソフトウェア エミュレーター + SSX360 参照デバイス)。 · GitHub
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
SSX360
/
マトリックススクロール
公共
通知
通知を変更するにはサインインする必要があります

アイケーション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .github .github コントロール コントロール ドキュメント ドキュメントの例 例 マトリックス スクロール マトリックス スクロール テスト テスト ベクトル ベクトル .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md SPEC.md SPEC.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
署名された AI 支援コードの出自に関するオープン プロトコル。
IDE 内で AI によって生成されたすべての変更は、
Ed25519 の ID は、公開キーと 1 つのコマンドを使用してオフラインで検証されます。の
v0.1.x リファレンス実装には、十分にテストされたソフトウェアの信頼のルートが同梱されています。
SSX360/NXP SE050 ハードウェア署名は、互換性のある参照デバイス パスです。
進歩。
📜 仕様: SPEC.md — ワイヤ形式、正規エンコーディング、スキーマ。
🛡 エージェント AI コントロール: docs/AGENTIC_AI_SECURITY.md
マトリックス スクロールして、エージェントティック AI サービスの慎重な導入に関する共同ガイダンスをマップします。
🔐 アルゴリズム: Ed25519 (RFC 8032)。秘密キーが SDK API によって公開されることはありません。
🧪 適合ベクトル:vectors/ — 非 Python 実装用。
🌐 サイト: https://matrixscroll.com
🔧 参照デバイス: SSX360 (NXP SE050 ハードウェア パスが進行中)。
pip インストール マトリックススクロール
クイックスタート
インポートマトリックススクロール
# このマシンではどの ID がアクティブですか?
print ( マトリックススクロール . ステータス ())
# {'schema': 'matrixscroll.identity.v1', 'available': True,
# 'モード': 'エミュレート', 'device_id': 'MS-A3F2-9C81', ...}
# 何かに署名します (リリースマニフェスト、コミットエンベロープ、SBOM、証拠パック)
signed = マトリックススクロール . Sign_manifest ({ "リリース" : "v1.0.0" , "アーティファクト" : [...]})
# どこでもオフラインで検証
として

マトリックススクロールを挿入します。 verify_manifest (署名付き)
CLI
$ マトリックススクロールステータス
{
" 利用可能 " : true、
" device_id " : " MS-A3F2-9C81 " 、
" モード " : " エミュレート " 、
" public_key " : " ... " 、
" スキーマ ": " matrixscroll.identity.v1 "
}
$matrixscrollsign release.json > release.signed.json
$matrixscroll release.signed.json を検証する
{ " device_id " : " MS-A3F2-9C81 " 、 " mode " : " エミュレート " 、 " ok " : true、 " signed_at " : " ... " }
マトリックススクロール検証終了 署名が有効な場合は 0、失敗した場合は 2
(マニフェストの改ざん、署名ブロックの欠落、間違ったスキーマ/アルゴリズム、不一致
デバイス ID、不正な形式の公開キー、読み取り不可能なファイル）。なしで CI からパイプします。
出力を解析します。
IDE / エージェント / CI
│
│ マニフェスト (リリース、コミット、証拠パック、SBOM、その他)
▼
マトリックススクロール.sign_manifest(...)
│
│ 正規の JSON (ソートされたキー、ASCII エスケープ、NaN なし、
│ 署名ブロックは入力から除外されます）
▼
IdentityProvider ──► Ed25519 署名
（今日も真似して、
明日はSSX360/SE050)
│
▼
署名済みマニフェスト ──►matrixscroll.verify_manifest(...)
(誰でも、どこでも、オフライン)
同じ Python API は、ローカル ソフトウェア エミュレータと
SSX360 デバイスの物理パス。 MATRIXSCROLL_MODE 環境との切り替え
変数; v0.1.x では、SE050 がリリースされるまでハードウェア モードは利用できないと報告されます。
輸送船。
レベル
プロバイダー
支援者
ステータス
L1 エミュレート
エミュレートされたプロバイダー
ソフトウェアキー、ファイルバックアップ (0600)
✅ 送料
L2ハードウェア
ハードウェアプロバイダー
NXP SE050 セキュア エレメント (SSX360)
🛠 Stage-0 プロトタイプ
L3 認証済み
未来
L2 + リモート認証
🗺 ロードマップ
status() は、モードと利用可能なフィールドを介してアクティブ レベルを公開します。
読み取り専用ダッシュボードは、ハードウェア パスが接続される前にレンダリングできます。
エミュレートされたキー ストア: ~/.matrixscroll/device.json
( MATRIXSCROLL_HOME でオーバーライドします)。
d

irectory が作成される 0700 ;シード ファイルは 0600 で開かれます
O_CREAT|O_EXCL なので、プライベート シードは一時的に世界から読み取り可能になることはありません。
種族は既存のキー ストアを黙って破壊することはできません。
破損または切り詰められたストアは、大規模なエラー ( IdentityError ) ではなく失敗します。
新しいアイデンティティを静かに鋳造します。 ID ローテーションは明示的な操作です。
計画されたハードウェア パスはディスク上にプライベートなものを何も保持しません - シードは封印されています
セキュアエレメント内。 v0.1.x では、このパスは型指定された可用性スタブです。
唯一のものではない参照実装
マトリックス スクロールはプロトコルです。この Python パッケージがリファレンスです。歓迎します
Rust、Go、TypeScript、および組み込み C での実装 - それらに対して実行します。
ベクトル/ 自己認証します。 COTRIBUTING.md を参照してください。
リポジトリには、機械可読な制御マトリックスが含まれています。
controls/agentic_ai_controls.json 、
限定エージェント証拠マニフェストの例
例/agentic_ai_evidence_manifest.json 、
および testing/test_agentic_guidance.py の実行可能チェック。これらはそれぞれを証明します
クレームはリポジトリ証拠にマップされ、署名されたエージェント スコープの変更は検証に失敗します。
仕様テキスト ( SPEC.md 、vectors/): CC0 1.0 — パブリック ドメイン。
SECURITY.md を参照してください。脆弱性を非公開で報告する
security@matrixscroll.com または GitHub セキュリティ アドバイザリ経由。
ハードウェア署名された AI 支援コード用のオープン プロトコル (Ed25519 信頼のルート、ソフトウェア エミュレーター + SSX360 参照デバイス)。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
マトリックス スクロール v0.1.1
最新の
2026 年 6 月 19 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
共有しないでください

私の個人情報

## Original Extract

Open protocol for hardware-signed AI-assisted code (Ed25519 root of trust, software emulator + SSX360 reference device). - SSX360/matrixscroll

GitHub - SSX360/matrixscroll: Open protocol for hardware-signed AI-assisted code (Ed25519 root of trust, software emulator + SSX360 reference device). · GitHub
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
SSX360
/
matrixscroll
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .github .github controls controls docs docs examples examples matrixscroll matrixscroll tests tests vectors vectors .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SPEC.md SPEC.md pyproject.toml pyproject.toml View all files Repository files navigation
Open protocol for signed AI-assisted code provenance.
Every AI-generated change in your IDE can be cryptographically signed by an
Ed25519 identity and verified offline with a public key and one command. The
v0.1.x reference implementation ships a well-tested software root of trust;
SSX360/NXP SE050 hardware signing is the compatible reference-device path in
progress.
📜 Spec: SPEC.md — wire format, canonical encoding, schemas.
🛡 Agentic AI controls: docs/AGENTIC_AI_SECURITY.md
maps Matrix Scroll to the joint Careful Adoption of Agentic AI Services guidance.
🔐 Algorithm: Ed25519 (RFC 8032). Private keys are never exposed by the SDK API.
🧪 Conformance vectors: vectors/ — for non-Python implementations.
🌐 Site: https://matrixscroll.com
🔧 Reference device: SSX360 (NXP SE050 hardware path in progress).
pip install matrixscroll
Quickstart
import matrixscroll
# What identity is active on this machine?
print ( matrixscroll . status ())
# {'schema': 'matrixscroll.identity.v1', 'available': True,
# 'mode': 'emulated', 'device_id': 'MS-A3F2-9C81', ...}
# Sign anything (a release manifest, a commit envelope, a SBOM, an evidence pack)
signed = matrixscroll . sign_manifest ({ "release" : "v1.0.0" , "artifacts" : [...]})
# Verify, anywhere, offline
assert matrixscroll . verify_manifest ( signed )
CLI
$ matrixscroll status
{
" available " : true,
" device_id " : " MS-A3F2-9C81 " ,
" mode " : " emulated " ,
" public_key " : " ... " ,
" schema " : " matrixscroll.identity.v1 "
}
$ matrixscroll sign release.json > release.signed.json
$ matrixscroll verify release.signed.json
{ " device_id " : " MS-A3F2-9C81 " , " mode " : " emulated " , " ok " : true, " signed_at " : " ... " }
matrixscroll verify exits 0 on a valid signature, 2 on any failure
(tampered manifest, missing signature block, wrong schema/algorithm, mismatched
device id, malformed public key, unreadable file). Pipe it from CI without
parsing the output.
your IDE / agent / CI
│
│ manifest (release, commit, evidence pack, SBOM, anything)
▼
matrixscroll.sign_manifest(...)
│
│ canonical JSON (sorted keys, ASCII-escaped, no NaN,
│ signature block excluded from input)
▼
IdentityProvider ──► Ed25519 signature
(Emulated today,
SSX360 / SE050 tomorrow)
│
▼
signed manifest ──► matrixscroll.verify_manifest(...)
(anyone, anywhere, offline)
The same Python API is designed to serve the local software emulator and the
physical SSX360 device path. Switch with the MATRIXSCROLL_MODE environment
variable; in v0.1.x, hardware mode reports unavailable until the SE050
transport ships.
Level
Provider
Backed by
Status
L1 Emulated
EmulatedProvider
Software key, file-backed (0600)
✅ Shipping
L2 Hardware
HardwareProvider
NXP SE050 secure element (SSX360)
🛠 Stage-0 prototype
L3 Attested
future
L2 + remote attestation
🗺 Roadmap
status() exposes the active level via the mode and available fields so
read-only dashboards can render before the hardware path is wired.
Emulated key store: ~/.matrixscroll/device.json
(override with MATRIXSCROLL_HOME ).
The directory is created 0700 ; the seed file is opened 0600 with
O_CREAT|O_EXCL so the private seed is never momentarily world-readable and
a race cannot silently clobber an existing key store.
A corrupt or truncated store fails loud ( IdentityError ) rather than
silently minting a fresh identity. Identity rotation is an explicit operation.
The planned hardware path holds nothing private on disk — the seed is sealed
in the secure element. In v0.1.x, this path is a typed availability stub.
Reference implementation, not the only one
Matrix Scroll is a protocol. This Python package is the reference. We welcome
implementations in Rust, Go, TypeScript, and embedded C — run them against
vectors/ to self-certify. See CONTRIBUTING.md .
The repo includes a machine-readable control matrix at
controls/agentic_ai_controls.json , an
example bounded-agent evidence manifest at
examples/agentic_ai_evidence_manifest.json ,
and executable checks in tests/test_agentic_guidance.py . These prove each
claim maps to repo evidence and that signed agent scope changes fail verify.
Specification text ( SPEC.md , vectors/ ): CC0 1.0 — public domain.
See SECURITY.md . Report vulnerabilities privately to
security@matrixscroll.com or via a GitHub Security Advisory.
Open protocol for hardware-signed AI-assisted code (Ed25519 root of trust, software emulator + SSX360 reference device).
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
matrixscroll v0.1.1
Latest
Jun 19, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
