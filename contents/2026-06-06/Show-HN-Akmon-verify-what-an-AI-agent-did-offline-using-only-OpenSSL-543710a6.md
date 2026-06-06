---
source: "https://github.com/radotsvetkov/akmon"
hn_url: "https://news.ycombinator.com/item?id=48423037"
title: "Show HN: Akmon, verify what an AI agent did offline using only OpenSSL"
article_title: "GitHub - radotsvetkov/akmon: Local-first Rust terminal AI agent (Akmon) · GitHub"
author: "radotsvetkov"
captured_at: "2026-06-06T09:44:59Z"
capture_tool: "hn-digest"
hn_id: 48423037
score: 1
comments: 0
posted_at: "2026-06-06T09:24:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Akmon, verify what an AI agent did offline using only OpenSSL

- HN: [48423037](https://news.ycombinator.com/item?id=48423037)
- Source: [github.com](https://github.com/radotsvetkov/akmon)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T09:24:14Z

## Translation

タイトル: Show HN: Akmon、OpenSSL のみを使用して AI エージェントがオフラインで何をしたかを検証する
記事タイトル: GitHub - radotsvetkov/akmon: ローカルファースト Rust ターミナル AI エージェント (Akmon) · GitHub
説明: ローカルファーストの Rust ターミナル AI エージェント (Akmon)。 GitHub でアカウントを作成して、radotsvetkov/akmon の開発に貢献してください。

記事本文:
GitHub - radotsvetkov/akmon: ローカルファースト Rust ターミナル AI エージェント (Akmon) · GitHub
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
ラドツヴェトコフ
/
アクモン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

e コード [その他のアクション] メニューを開く フォルダーとファイル
243 コミット 243 コミット .github .github crates crates docs docs ランディング パッケージング/ homebrew パッケージング/ homebrew scripts/ docs scripts/ docs .gitignore .gitignore AKMON.md AKMON.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md CONTRIBUTORS.md CONTRIBUTORS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス POLICY_SYSTEM_EXPLANATION.md POLICY_SYSTEM_EXPLANATION.md README.md README.md SECURITY.md SECURITY.md拒否.toml 拒否.toml 錆びツールチェーン.toml 錆びツールチェーン.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
✦ ✦ ✦
▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒
▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓
▓▓ ▓▓
▓▓ ▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
アクモン
Akmon は、AI エージェントの改ざん証拠および検証レイヤーです。すでに実行しているエージェントの上に配置されます。

独自のツール、または OpenTelemetry を発行する任意のツール。各セッションは、コンテンツにアドレス指定され、暗号的に署名されたポータブルな記録となり、他の人が自分で検証できるようになります。最も重要な部分はこれです。第三者は openssl のみを使用してオフラインで署名をチェックできます。Akmon のインストールもクラウド サービスも必要なく、記録の作成者を信頼する必要もありません。
ウェブサイト: radotsvetkov.github.io/akmon 。ドキュメント: radotsvetkov.github.io/akmon/docs 。
AI エージェントが何かを変更する場合、実際に何をしたかを後で証明する必要がある場合があります。質問者は規制当局、セキュリティレビュー担当者、またはインシデントチームである可能性があり、それは事実から何年も経っている可能性があります。彼らはあなたを信頼せず、あなたのツールを実行しないかもしれません。ほとんどのエージェントのテレメトリはこれに耐えることができません。それはプロセス メモリ内、または変更可能な符号なしスパン内に存在しており、「AI がやった」という答えは誰もが受け入れるものではありません。 EU AI 法に基づき、第 12 条および附属書 IV の高リスク伐採義務は 2026 年 8 月 2 日から適用されます。
アクモンは証拠そのものをあなたが発送するものとして扱います。これは、ユーザー自身のセッション、または OpenTelemetry で計測されたエージェントからのセッションを取得し、それを封印されたレコードに変換します。その後、他の誰かが、Akmon について聞いたことのないマシン上でも、標準ツールを使用して、そのレコードの整合性と作成者を個別にチェックできます。
ここにあるすべてのコマンドは本日出荷され、テストでカバーされています。焦点はエージェントではなく検証チェーンです。
ライフサイクルを締めくくるその他のコマンド: Bundle Export および Bundle import 、inspect 、 diff 、replay (Akmon 自身が生成したセッションの決定論的再生)、redact 、audit 、evidence 、verify 、policy 、および Doctor 。
Akmon はいくつかの方法でインストールできます。
Homebrew を使用した macOS または Linux の場合:
醸造タップ ラドツヴェトコフ/アクモン
brew install akmon # CLI
醸造インストールAG

ef-verify # 監査人用のスタンドアロン検証ツール
または、事前に構築されたバイナリを直接取得します。 GitHub の各リリースには、Linux および macOS 用のビルド済みの akmon バイナリと agef-verify バイナリに加えて、SHA256SUMS ファイルが添付されているため、ダウンロードしたものを確認できます。 x86_64 上の Linux の場合:
Base=https://github.com/radotsvetkov/akmon/releases/latest/download
カール -LO $base /akmon-linux-x86_64
カール -LO $base /agef-verify-linux-x86_64
カール -LO $base /SHA256SUMS
# インストールする前に、公開されているチェックサムに対してダウンロードをチェックします。
sha256sum --check --ignore-missing SHA256SUMS
chmod +x akmon-linux-x86_64 agef-verify-linux-x86_64
mv akmon-linux-x86_64 ~ /bin/akmon
mv agef-verify-linux-x86_64 ~ /bin/agef-verify
macOS では、ファイル名は Apple シリコンの場合は akmon-darwin-arm64 および agef-verify-darwin-arm64、Intel の場合は -x86_64 のバリアントであり、 shasum -a 256 --check --ignore-missing SHA256SUMS で確認します。
または、任意のプラットフォーム上のソースからビルドします。
カーゴインストール --git https://github.com/radotsvetkov/akmon akmon
カーゴインストール --git https://github.com/radotsvetkov/akmon agef-verify
ここでは、エージェントの OpenTelemetry トレースから、見知らぬ人がチェックできる証拠までの完全な実行を示します。
# 1. 署名鍵を作成します。プライベートは半分秘密にし、.pub だけを配布します。
akmon バンドル keygen --outsigner.pk8 --public-outsigner.pub
# 2. 実際のトレースをセッションに変換し、次にポータブル バンドルに変換します。
akmon otel import trace.json --journal ./journal
akmon バンドル エクスポート < セッション ID > --journal ./journal --output Audit.akmon
#3. 署名します。
akmon バンドル署名 Audit.akmon --key signeder.pk8
# 4. 整合性、署名、キャプチャ レベルを検証します。
akmon バンドルは、audit.akmon --verify-keysigner.pub --require-signature を検証します
#5. Akmon を関与させずに、openssl のみを使用してそれを証明します。
akmon バンドルprove-openssl Audit.akmon --verify-key signeder.pub -

-out-dir 証明
openssl pkeyutl -verify -pubin -inkeyproof/pubkey.pem -rawin -inproof/statement.bin -sigfileproof/signature.bin
検証手順には OpenSSL 3.x を使用します。 macOS に付属の openssl は LibreSSL ですが、Ed25519 を検証できません。最初に OpenSSL 3.x をインストールします (たとえば、 brew install openssl@3 )。
ドキュメントには、OTEL トレースからオフライン openssl 証明までの完全なウォークスルーがあります。
アクモンは各レコードがどの程度完全であるかを明確に示しており、誰もそれを黙って誇張することができないように、そのレベルをレコードに署名します。
完全なキャプチャ (capture_level = full) は、Akmon 自身のエージェントから提供されます。プロンプト、応答、およびツール呼び出しが保持されるため、セッションは akmon reports で決定的に再生されます。
構造キャプチャ (capture_level = Structure) は、別のエージェントの OpenTelemetry トレースをインポートすることで得られます。セッションの形状、プロバイダー呼び出し、ツール呼び出し、メタデータ、およびトレースに偶然含まれていたコンテンツはすべて取得できますが、バイト単位の再生は取得できません。 akmonbundle verify --require-capture full はこれらに対して意図的に失敗し、akmon restart はそれらを拒否します。
インポートされたトレースが完全な記録としてドレスアップされることはありません。この境界線を明確に保つことが、信頼層を設ける価値がある理由のすべてです。
Microsoft は、堅牢なガバナンス ランタイムであるオープンソースのエージェント ガバナンス ツールキット (2026 年 4 月から一般提供開始) と、強力な改ざん防止クラウド台帳である Azure Confidential Ledger を出荷しています。ただし、2026 年 6 月の時点では、非対称キー (Ed25519) で署名され、Microsoft から何もインストールされていない見知らぬ人でもチェックでき、決定的に再生可能で、任意のエージェントの上に座ることができる、ポータブルで自己完結型の証拠アーティファクトを提供する Microsoft 製品は 1 つもありません。ツールキットの改ざん証拠はハッシュ チェーンと HMAC であり、非対称署名はありません

スタンドアロンの検証機能もありません。 Confidential Ledger の署名付きマークル領収書は非常に優れていますが、Azure に関連付けられており、エージェントを認識しません。 Microsoft 自身の Foundry ドキュメントには、そのトレースは完全な再生をサポートできないと記載されています。
そのギャップがアクモンにぴったりです。移植性があり、署名されており、クラウドに依存せず、エージェント対応で、再生可能です。
Akmon はそのどれにも代わろうとしているわけではなく、Microsoft が明らかに先を行っている部分もあります。その配布方法は 1 つです。Purview と Copilot Control System は、すべての Microsoft 365 テナントにすでに存在しています。 Confidential Ledger のオフライン検証可能な領収書も別の例であり、標準化団体における Microsoft の重要性も同様です。単一のツールではこれらに匹敵しません。 Akmon はそれらを補完することを目的としています。Purview がキャプチャしたものを封印し、Foundry が追跡したものをエクスポートして検証します。
Akmon は、EU AI 法 (第 12 条および附属書 IV のロギング、2026 年 8 月 2 日から高リスク義務が課される)、NIST AI リスク管理フレームワーク (MEASURE 2.8)、SOC 2 (CC7.x および CC8.1) などのフレームワークの証拠を作成できるように構築されています。これは適合性証明ではなく、適合性を保証するものではありません。 AGEF の証拠から特定の統制へのマッピングは、自社のコンプライアンス チームと法務チームで検証するものとして扱います。
Akmon には、独自のエージェント ( akmon 、 akmon chat 、および akmon --task ... ) が同梱されています。ファイル書き込み、シェル、ネットワークに対する型指定された権限チェックがあり、ローカルまたはホストされたモデルを実行し、ポリシー プロファイルをサポートし、MCP サーバーを管理します。これは主に参照プロデューサーとして存在し、決定論的に再生するフルキャプチャ セッションを取得する方法です。 Cursor または Claude Code をアウトコードしようとしているわけではありません。 Akmon が提供する価値は証拠層であり、どのエージェントを好むかに関係なく機能します。
# Akmon 自身のエージェントからのセッションとその検証パイプライン。
アクモン

--yes --output json --task "rustfmt 修正のみを適用します" |ティーラン.json
akmon 監査検証 .akmon/audit/ < セッション ID > .jsonl
akmon 証拠検証 .akmon/evidence/ < session-id > .json
akmon slo verify .akmon/evidence/ < session-id > .json --strict
証拠フォーマット（AGEF）
Akmon の記録は、AI エージェントのセッション証拠のための移植性があり、コンテンツにアドレス指定され、改ざん防止形式である AGEF 仕様に従っています。 Akmon は AGEF v0.1.3 を実装します。このバージョンでは、ハッシュ チェーンの最上位に 2 つのオプションの部分が追加されます。1 つは、レコードをキーに帰属させる、manifest.signatures[] のオフライン Ed25519 署名、もう 1 つは、セッションの背後にある責任者を記録する、manifest.operator_attestations[] のオペレータ証明書です。どちらもオプションであるため、プレーン バンドルは小さいままで、古いリーダーは引き続き動作します。
プロジェクト サイト: radotsvetkov.github.io/akmon
ホストされているハンドブック: radotsvetkov.github.io/akmon/docs
概要: docs/src/introduction.md
OTEL トレースから openssl 証明までのチュートリアル: docs/src/tutorials/otel-to-openssl-walkthrough.md
バンドル keygen のリファレンス: docs/src/reference/bundle-keygen.md
バンドルアテストのリファレンス: docs/src/reference/bundle-attest.md
バンドルprove-opensslのリファレンス: docs/src/reference/bundle-prove-openssl.md
貢献ガイド: COTRIBUTING.md
行動規範: CODE_OF_CONDUCT.md
Apache-2.0のみ。 「ライセンス」を参照してください。
アクモンは金床を意味する古いギリシャ語です。そのアイデアは、権限、改ざんを証明する監査証跡、誰もが検証できる証拠など、すべてのストライキを制御しながら、プレッシャーと精度を持って困難な作業を形作ることです。
ローカルファーストのRustターミナルAIエージェント（Akmon）
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
5
フォーク
レポートリポジトリ
リリース
22
v2.2.0
最新の
2026 年 6 月 6 日
+ 21 リリース
パッケージ

0
読み込み中にエラーが発生しました。このページをリロードしてください。
途中でエラーが発生しました

[切り捨てられた]

## Original Extract

Local-first Rust terminal AI agent (Akmon). Contribute to radotsvetkov/akmon development by creating an account on GitHub.

GitHub - radotsvetkov/akmon: Local-first Rust terminal AI agent (Akmon) · GitHub
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
radotsvetkov
/
akmon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
243 Commits 243 Commits .github .github crates crates docs docs landing landing packaging/ homebrew packaging/ homebrew scripts/ docs scripts/ docs .gitignore .gitignore AKMON.md AKMON.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTORS.md CONTRIBUTORS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE POLICY_SYSTEM_EXPLANATION.md POLICY_SYSTEM_EXPLANATION.md README.md README.md SECURITY.md SECURITY.md deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
✦ ✦ ✦
▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒
▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓
▓▓ ▓▓
▓▓ ▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
Akmon
Akmon is a tamper-evident evidence and verification layer for AI agents. It sits on top of whatever agent you already run, whether that is your own or any tool that emits OpenTelemetry. Each session becomes a portable, content-addressed, cryptographically signed record that someone else can verify for themselves. The part that matters most is this: a third party can check a signature offline using nothing but openssl , with no Akmon install, no cloud service, and no need to trust whoever produced the record.
Website: radotsvetkov.github.io/akmon . Documentation: radotsvetkov.github.io/akmon/docs .
When an AI agent changes something, you may have to prove later what it actually did. The person asking could be a regulator, a security reviewer, or an incident team, and it might be years after the fact. They may not trust you, and they may not run your tools. Most agent telemetry cannot stand up to that. It lives in process memory, or in mutable, unsigned spans, and "the AI did it" is not an answer anyone will accept. Under the EU AI Act, the high-risk logging obligations in Article 12 and Annex IV start to apply on 2 August 2026.
Akmon treats the evidence itself as the thing you ship. It takes a session, either your own or one from any OpenTelemetry-instrumented agent, and turns it into a sealed record. Someone else can then check that record's integrity and authorship independently, with standard tools, even on a machine that has never heard of Akmon.
Every command here ships today and is covered by tests. The focus is the verification chain, not the agent.
Other commands round out the lifecycle: bundle export and bundle import , inspect , diff , replay (deterministic playback of sessions Akmon produced itself), redact , audit , evidence , verify , policy , and doctor .
You can install Akmon a few ways.
On macOS or Linux with Homebrew:
brew tap radotsvetkov/akmon
brew install akmon # the CLI
brew install agef-verify # the standalone verifier for auditors
Or grab the prebuilt binaries directly. Each GitHub release attaches prebuilt akmon and agef-verify binaries for Linux and macOS, plus a SHA256SUMS file so you can check what you downloaded. For Linux on x86_64:
base=https://github.com/radotsvetkov/akmon/releases/latest/download
curl -LO $base /akmon-linux-x86_64
curl -LO $base /agef-verify-linux-x86_64
curl -LO $base /SHA256SUMS
# Check the downloads against the published checksums before installing.
sha256sum --check --ignore-missing SHA256SUMS
chmod +x akmon-linux-x86_64 agef-verify-linux-x86_64
mv akmon-linux-x86_64 ~ /bin/akmon
mv agef-verify-linux-x86_64 ~ /bin/agef-verify
On macOS the file names are akmon-darwin-arm64 and agef-verify-darwin-arm64 for Apple silicon, or the -x86_64 variants for Intel, and you check them with shasum -a 256 --check --ignore-missing SHA256SUMS .
Or build from source on any platform:
cargo install --git https://github.com/radotsvetkov/akmon akmon
cargo install --git https://github.com/radotsvetkov/akmon agef-verify
Here is a full run, from any agent's OpenTelemetry trace to a proof a stranger can check.
# 1. Make a signing key. Keep the private half secret and hand out only the .pub.
akmon bundle keygen --out signer.pk8 --public-out signer.pub
# 2. Turn a real trace into a session, then a portable bundle.
akmon otel import trace.json --journal ./journal
akmon bundle export < session-id > --journal ./journal --output audit.akmon
# 3. Sign it.
akmon bundle sign audit.akmon --key signer.pk8
# 4. Verify integrity, the signature, and the capture level.
akmon bundle verify audit.akmon --verify-key signer.pub --require-signature
# 5. Prove it with openssl alone, no Akmon involved.
akmon bundle prove-openssl audit.akmon --verify-key signer.pub --out-dir proof
openssl pkeyutl -verify -pubin -inkey proof/pubkey.pem -rawin -in proof/statement.bin -sigfile proof/signature.bin
Use OpenSSL 3.x for the verify step. The openssl that ships with macOS is LibreSSL, which cannot verify Ed25519. Install OpenSSL 3.x first, for example with brew install openssl@3 .
There is a full walkthrough in the docs: from an OTEL trace to an offline openssl proof .
Akmon is explicit about how complete each record is, and it signs that level into the record so nobody can quietly overstate it.
Full capture ( capture_level = full ) comes from Akmon's own agent. It keeps the prompts, responses, and tool calls, so the session replays deterministically with akmon replay .
Structural capture ( capture_level = structural ) comes from importing another agent's OpenTelemetry trace. You get the shape of the session, its provider calls, tool calls, metadata, and whatever content the trace happened to include, but not a byte-for-byte replay. akmon bundle verify --require-capture full fails on these on purpose, and akmon replay refuses them.
An imported trace is never dressed up as a full recording. Keeping that line clear is the whole reason a trust layer is worth having.
Microsoft ships a solid governance runtime, the open-source Agent Governance Toolkit, which has been generally available since April 2026, and a strong tamper-evident cloud ledger in Azure Confidential Ledger. As of June 2026, though, no single Microsoft product gives you one portable, self-contained evidence artifact that is signed with an asymmetric key (Ed25519), checkable by a stranger who has nothing from Microsoft installed, deterministically replayable, and able to sit on top of any agent. The Toolkit's tamper-evidence is a hash chain plus HMAC, with no asymmetric signature and no standalone verifier. Confidential Ledger's signed Merkle receipts are genuinely good, but they are tied to Azure and are not aware of agents. Microsoft's own Foundry documentation says its traces cannot support full replay.
That gap is where Akmon fits. It is portable, signed, cloud-independent, agent-aware, and replayable.
Akmon is not trying to replace any of that, and there are places where Microsoft is clearly ahead. Its distribution is one: Purview and the Copilot Control System are already in every Microsoft 365 tenant. Confidential Ledger's offline-verifiable receipts are another, and so is Microsoft's weight in the standards bodies. A single tool will not match those. Akmon is meant to complement them: seal what Purview captures, and export and verify what Foundry traces.
Akmon is built to help you produce evidence for frameworks like the EU AI Act (Article 12 and Annex IV logging, with high-risk obligations from 2 August 2026), the NIST AI Risk Management Framework (MEASURE 2.8), and SOC 2 (CC7.x and CC8.1). It is not a compliance certification, and it does not guarantee compliance. Treat any mapping from AGEF evidence to specific controls as something to validate with your own compliance and legal teams.
Akmon ships with its own agent ( akmon , akmon chat , and akmon --task ... ). It has typed permission checks for file writes, shell, and network, runs local or hosted models, supports policy profiles, and governs MCP servers. It exists mainly as the reference producer, the way to get full-capture sessions that replay deterministically. It is not trying to out-code Cursor or Claude Code. The value Akmon offers is the evidence layer, and that works no matter which agent you prefer.
# A session from Akmon's own agent, with its verification pipeline.
akmon --yes --output json --task " apply rustfmt fixes only " | tee run.json
akmon audit verify .akmon/audit/ < session-id > .jsonl
akmon evidence verify .akmon/evidence/ < session-id > .json
akmon slo verify .akmon/evidence/ < session-id > .json --strict
The evidence format (AGEF)
Akmon's records follow the AGEF specification , a portable, content-addressed, tamper-evident format for AI-agent session evidence. Akmon implements AGEF v0.1.3. That version adds two optional pieces on top of the hash chain: offline Ed25519 signatures in manifest.signatures[] , which make a record attributable to a key, and operator attestations in manifest.operator_attestations[] , which record the accountable person behind a session. Both are optional, so a plain bundle stays small and older readers keep working.
Project site: radotsvetkov.github.io/akmon
Hosted handbook: radotsvetkov.github.io/akmon/docs
Introduction: docs/src/introduction.md
Tutorial, from an OTEL trace to an openssl proof: docs/src/tutorials/otel-to-openssl-walkthrough.md
Reference for bundle keygen : docs/src/reference/bundle-keygen.md
Reference for bundle attest : docs/src/reference/bundle-attest.md
Reference for bundle prove-openssl : docs/src/reference/bundle-prove-openssl.md
Contribution guide: CONTRIBUTING.md
Code of Conduct: CODE_OF_CONDUCT.md
Apache-2.0 only. See LICENSE .
Akmon is an old Greek word for anvil. The idea is to shape difficult work with pressure and precision while keeping control over every strike: the permissions, the tamper-evident audit trail, and the evidence anyone can verify.
Local-first Rust terminal AI agent (Akmon)
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
5
forks
Report repository
Releases
22
v2.2.0
Latest
Jun 6, 2026
+ 21 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while

[truncated]
