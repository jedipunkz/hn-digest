---
source: "https://github.com/rewired89/HSIP-1PHASE"
hn_url: "https://news.ycombinator.com/item?id=48612178"
title: "HSIP–local identity server in Rust with Ed25519 signing and AI agent governance"
article_title: "GitHub - rewired89/HSIP-1PHASE: A self-hosted identity server that gives AI agents a cryptographic identity, generates tamper-proof audit trails, and covers financial compliance requirements (MiFID II, FINRA 4511, SOX §404, DORA, SWIFT CSCF) all in a single Rust binary with no cloud dependency. · Gi\n[truncated]"
author: "Rewired89"
captured_at: "2026-06-20T21:34:23Z"
capture_tool: "hn-digest"
hn_id: 48612178
score: 3
comments: 0
posted_at: "2026-06-20T19:27:43Z"
tags:
  - hacker-news
  - translated
---

# HSIP–local identity server in Rust with Ed25519 signing and AI agent governance

- HN: [48612178](https://news.ycombinator.com/item?id=48612178)
- Source: [github.com](https://github.com/rewired89/HSIP-1PHASE)
- Score: 3
- Comments: 0
- Posted: 2026-06-20T19:27:43Z

## Translation

タイトル: HSIP – Ed25519 署名と AI エージェント ガバナンスを備えた Rust のローカル ID サーバー
記事のタイトル: GitHub - rewired89/HSIP-1PHASE: AI エージェントに暗号化 ID を提供し、改ざん防止監査証跡を生成し、財務コンプライアンス要件 (MiFID II、FINRA 4511、SOX §404、DORA、SWIFT CSCF) をクラウド依存性のない単一の Rust バイナリですべてカバーするセルフホスト型アイデンティティ サーバー。・ギ
[切り捨てられた]
説明: AI エージェントに暗号化 ID を提供し、改ざん防止監査証跡を生成し、財務コンプライアンス要件 (MiFID II、FINRA 4511、SOX §404、DORA、SWIFT CSCF) をすべてクラウド依存性のない単一の Rust バイナリでカバーする自己ホスト型 ID サーバーです。 - rewired89/HSIP-1PHASE

記事本文:
GitHub - rewired89/HSIP-1PHASE: AI エージェントに暗号化 ID を提供し、改ざん防止監査証跡を生成し、財務コンプライアンス要件 (MiFID II、FINRA 4511、SOX §404、DORA、SWIFT CSCF) をすべてクラウド依存性のない単一の Rust バイナリでカバーする自己ホスト型アイデンティティ サーバーです。 · GitHub
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
別のタブまたは Wi-Fi でアカウントを切り替えた

窓。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
再配線89
/
HSIP-1フェーズ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
36 コミット 36 コミット .cargo .cargo .github .github ブラウザ拡張機能 ブラウザ拡張機能 クレート クレート ダッシュボード ダッシュボード ドキュメント ドキュメントの例 例 負荷テスト 負荷テスト SDKS SDKS 仕様 仕様 .gitattributes .gitattributes .gitignore .gitignore .gitleaksignore .gitleaksignore CLAUDE.md CLAUDE.md CODEMAP.md CODEMAP.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DEEPFAKE_DEFENSE.md DEEPFAKE_DEFENSE.md DEMO.md DEMO.md DEMO_WINDOWS.ps1 DEMO_WINDOWS.ps1 DEPLOYMENT.md DEPLOYMENT.md Dockerfile Dockerfile Dockerfile.api Dockerfile.api ライセンス ライセンス LINUX_SETUP.md LINUX_SETUP.md README.md README.md TESTING.md TESTING.md THREAT_MODEL.md THREAT_MODEL.md WINDOWS_SETUP.md WINDOWS_SETUP.md build-release.ps1 build-release.ps1 build.rs build.rs config.example.toml config.example.tomldeny.tomldeny.toml docker-compose.yml docker-compose.yml hsip.toml hsip.toml install.ps1 install.ps1 install.sh install.sh launch-hsip.bat launch-hsip.bat Railway.toml Railway.toml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
バイナリが 1 つ。雲はありません。購読はありません。個人、AI エージェント、金融機関向けの暗号化 ID と改ざん防止監査証跡。
🌐 hsip.rewired89.github.io/HSIP-1PHASE — ワンクリックでダウンロードできるランディング ページ
すべての鍵はあなたのものです。すべてのバイトはローカルで実行されます。雲はありません。購読はありません。商用利用にはライセンスが必要です — sanchezleal1989@gmail.com までご連絡ください。脅威モデルを読む →
Windows — hsip-windows-x64.exe をダウンロード →

ダブルクリック → ブラウザが自動的に開きます。
カール -sSf https://raw.githubusercontent.com/rewired89/HSIP-1PHASE/main/install.sh |しー
自作:
醸造タップ rewired89/hsip https://github.com/rewired89/HSIP-1PHASE && 醸造インストール hsip
なぜこれが存在するのか - 今すぐ
2026 年、3 つのことが同時に起こりました。
AI エージェントは、何をしたか、誰が許可したかについての信頼できる記録がなくても、ユーザーに代わって行動します。
OpenAI、Google、Meta は、思考に使用するツール内で広告を配信します。プロンプトはモデルをトレーニングします。
ディープフェイクにより、偽造できない暗号署名が含まれていない限り、デジタル証拠は無意味になりました。
HSIP は 3 つすべてに対する答えです。これはハードウェア上で実行され、キーですべてに署名し、完全に所有する改ざん防止の監査証跡を提供します。
したいです...
何を実行するか
追跡を停止 — 使用するすべてのアプリで広告、テレメトリ、監視をブロックします
DNS トラッカー ブロッカー
私の言ったことを証明する — 私がこの時点でこのメッセージを書いたという法廷で認められる証拠を作成する
署名付きメッセージ + 監査証跡
AI エージェントを制御 – AI が何をしたかを正確に確認し、即座にアクセスを取り消します
AI ウォッチ + 同意ウォレット
プライバシーを尊重したソフトウェアを構築する — アプリまたは AI エージェントに同意インフラストラクチャを追加する
開発者SDK →
企業監査コンプライアンス — GDPR、法廷記録、法的レベルの証拠チェーン
エンタープライズ展開 →
金融サービス インフラストラクチャ — MiFID II、FINRA 4511、SOX §404、DORA、SWIFT CSCF 準拠
金融サービス →
ダウンロード
プラットフォーム
ファイル
窓
hsip-windows-x64.exe
macOS アップルシリコン
hsip-macos-arm64
macOS インテル
hsip-macos-x64
Linux
hsip-linux-x64
Windows: .exe をダブルクリックします。自動的にインストールされ、デスクトップ ショートカットが作成され、ブラウザで自動的に開きます。
Mac / Linux: chmod +x hsip-macos-arm64 && ./hsip-macos-arm64 — ブラウザが自動的に開きます

ついでに。
1. DNS トラッカー ブロッカー — システム全体ですべてをブロックします
HSIP は、追跡リクエストがマシンに到達する前に、DNS レベルで追跡リクエストを傍受します。 1 つのブラウザだけではなく、実行するすべてのアプリが対象となります。
Google Analytics、Facebook Pixel、Hotjar、TikTok、DoubleClick、Microsoft テレメトリなど 200 以上をブロックします。ダッシュボードをワンクリックしてオンにします。ゼロ構成。
ブラウザ拡張機能との違い: ブラウザ拡張機能は 1 つのブラウザのみを保護します。 HSIP は、デスクトップ アプリ、バックグラウンド プロセス、すべてのブラウザを一度にネットワーク レベルでブロックします。
2. 署名付きメッセージ — ディープフェイクと闘い、紛争に勝つ
HSIP 経由で送信するすべてのメッセージは、個人の Ed25519 キーで署名されます。結果は次のことを数学的に証明します。
この証拠は法廷でも機械でも誰でも検証できます。偽造することはできません。
契約確認: 「2026 年 3 月 28 日にこれらの条件に同意したことを確認します。」 — 署名、タイムスタンプがあり、検証可能。
証拠の争議: いつ何を言ったかを証明する暗号化された領収書を数秒で作成します。
ディープフェイク防御: あなたが言っていないことを誰かが主張したとき、あなたの署名履歴はそうではないことを証明します。
AI コマンド認証: AI エージェントに与えたすべての指示はキーで署名されます。否認可能性は両方の方向でなくなりました。
3. AI Watch — AI が何をしたかを正確に知る
接続するすべての AI エージェント (Claude、ChatGPT、Siri、HTTP 対応ツール) はリアルタイムで追跡されます。
速度監視 - エージェントが異常な数のリクエストを行った場合にアラートを送信します
異常検出 — 正常なパターンから外れた動作にフラグを立てます
ワンクリックで切断 - エージェントのアクセスを即座に取り消します
完全に署名された監査証跡 - エージェントが実行したすべてのアクション、署名およびタイムスタンプ
これは、AI の「ブラック ボックス レコーダー」です。何か問題が発生した場合、何が起こったのかを正確に把握でき、

いつ。
4. 同意ウォレット — 機械読み取り可能なアクセス制御
読まずにクリックする Cookie バナーの代わりに、HSIP は実際に制御する同意レイヤーを作成します。
あなたに連絡したり、あなたのデータにアクセスしたりする許可を持つすべての当事者を表示します
各当事者に何が許可されているかを正確に確認する
同意に時間制限を設定します - 有効期限は自動的に切れます
ワンクリックで同意を取り消し、即時有効になります
HSIP をサポートするサードパーティ サービスは、行動する前に同意を問い合わせることができます。許可がありません - アクセスできません。
HSIP のすべての操作 (メッセージの署名、同意の付与、キーの作成、AI アクションのログ記録) は、BLAKE3 ハッシュ チェーン監査ログに書き込まれます。エントリを改ざんするとチェーンが切れます。
法的手続き、コンプライアンス監査、または個人記録のために、いつでもログをエクスポートできます。
HSIP は、銀行、トレーディング デスク、フィンテック、および改ざん防止監査証跡、AI エージェント ガバナンス、機関間の ID 検証を必要とするあらゆる規制対象機関向けの暗号インフラストラクチャです。中間に中央のクラウド ベンダーは必要ありません。
顧客は機関投資家であり、個人投資家ではありません。 HSIP はデータセンター (またはオンプレミス) 内で実行され、Ed25519 キーペアを使用してすべてのアクションに署名し、システム、アナリスト、AI エージェントが監査証跡に記載されているとおりに実行したことを示す法的に弁護可能な証拠を生成します。
金融機関が今これを必要とする理由
1. AI エージェントは機関に代わって行動します。規制当局は、各アクションを誰が承認したかを尋ねる予定です。各エージェントに付加された暗号化 ID とすべてのリクエストの追加専用ログがなければ、その質問に答えることはできません。 HSIP は、すべての AI エージェントに独自の Ed25519 キーペアを割り当て、エージェントが実行するすべてのアクションをログに記録し、ミリ秒単位でアクセスを取り消すことができます。
2. MiFID II 第 25 条および FINRA 規則 4511 では、何を証明する必要がありますか

あなたのシステムは、いつ、誰の権限で実行されたのか。データベース内のログは証拠ではなく、変更される可能性があります。 BLAKE3 ハッシュ チェーン監査ログが証拠です。エントリを改ざんすると、チェーンが切断され、誰でも検出できます。
3. オープン バンキング (PSD2) では、機械可読で期限付きの同意が義務付けられています。 HSIP の同意ウォレットは、まさにそれを生成します。つまり、特定のアクションを対象とした、自動的に有効期限が切れ、リアルタイムで取り消し可能な、暗号化された署名付きの許可です。コンプライアンス チームが証明できない Cookie バナーはもう必要ありません。
4. 機関間の信頼が崩壊する。相手からメッセージが届いたとき、それが転送中に変更されていないことをどのように確認しますか? HSIP のフェデレーテッド トラスト層を使用すると、機関は Ed25519 を交換してアウトオブバンド (電子メール、安全なチャネル) でキーを検証し、その後、将来のメッセージを暗号的に検証できます。中央レジストリ、PKI ベンダー、単一障害点は必要ありません。
5. DORA および SWIFT CSCF では、異常な AI または自動化されたシステムの動作を検出して対応する必要があります。 HSIP の速度監視は、1 分あたり 100 リクエストを超えるエージェントにフラグを立て、1,000 リクエスト/分でアクセスを自動取り消します。各ステップで署名された監査エントリが使用されます。
規制
HSIP の対象となる内容
SOX §404
追加専用の BLAKE3 ハッシュ チェーン監査ログ。すべてのコントロール アクションは Ed25519 で署名されています。監査人向けにエクスポート可能。
FINRA 規則 4511
6 年間の改ざん防止記録保持。一括監査エクスポート用の API エンドポイント。署名チェーンは、エントリが変更されていないことを証明します。
MiFID II アート。 25
機関の Ed25519 キーで署名された取引ごとの認可。タイムスタンプ + 署名 = 防御可能な適合性レコード。
PSD2 / オープンバンキング
範囲、有効期限、取り消しを含む機械可読な同意付与。 /v1/consent/grant をexpires_in_secondsでPOSTします。
GDPR アート。 7
文書化された範囲で暗号化署名された同意。元号の権利については DELETE /v1/tenant/erase

もちろん。監査ログは、処理時に同意が有効であったことを証明します。
ドーラ
AI エージェントの速度監視、異常検出、自動取り消し。 DELETE /v1/keys/:id によるインシデント対応。署名された監査証跡内のすべてのイベント。
スイフトCSCF
Ed25519 メッセージ認証により、不正な命令の挿入が防止されます。取引相手ごとに検証されたフェデレーション信頼キー。共有秘密はありません。
ISO20022
Ed25519 で署名された支払いメッセージ。機関の公開鍵を保有するあらゆる取引相手によって検証可能。構造による否認防止。
金融機関向け AI エージェント ガバナンス
貴機関が展開するすべての AI システム (取引アルゴリズム、ドキュメント プロセッサ、顧客対応チャットボット、内部アシスタント) は、独自の Ed25519 キーペアを HSIP に登録します。
# 取引アルゴリズムを管理された AI エージェントとして登録する
hsip エージェント登録 " algo-trading-v2 " --expires-days 90
# すべてのアクティブなエージェントとそのリクエスト速度をリストします。
hsipエージェントリスト
# 予期せぬ動作をするエージェントをただちに取り消す
hsip エージェントが「 algo-trading-v2 」を取り消します
各エージェントに得られるもの:
固有の Ed25519 キーペア — 署名されたすべてのアクションは、「システム」だけでなく、その特定のエージェントまで追跡可能です。
速度監視 - リクエストが 100/分を超えると、異常監査エントリがトリガーされます。 1,000/分を超えると自動取り消しがトリガーされます
完全に署名された監査証跡 — ev

[切り捨てられた]

## Original Extract

A self-hosted identity server that gives AI agents a cryptographic identity, generates tamper-proof audit trails, and covers financial compliance requirements (MiFID II, FINRA 4511, SOX §404, DORA, SWIFT CSCF) all in a single Rust binary with no cloud dependency. - rewired89/HSIP-1PHASE

GitHub - rewired89/HSIP-1PHASE: A self-hosted identity server that gives AI agents a cryptographic identity, generates tamper-proof audit trails, and covers financial compliance requirements (MiFID II, FINRA 4511, SOX §404, DORA, SWIFT CSCF) all in a single Rust binary with no cloud dependency. · GitHub
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
rewired89
/
HSIP-1PHASE
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
36 Commits 36 Commits .cargo .cargo .github .github browser-extension browser-extension crates crates dashboard dashboard docs docs examples examples load-test load-test sdks sdks spec spec .gitattributes .gitattributes .gitignore .gitignore .gitleaksignore .gitleaksignore CLAUDE.md CLAUDE.md CODEMAP.md CODEMAP.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DEEPFAKE_DEFENSE.md DEEPFAKE_DEFENSE.md DEMO.md DEMO.md DEMO_WINDOWS.ps1 DEMO_WINDOWS.ps1 DEPLOYMENT.md DEPLOYMENT.md Dockerfile Dockerfile Dockerfile.api Dockerfile.api LICENSE LICENSE LINUX_SETUP.md LINUX_SETUP.md README.md README.md TESTING.md TESTING.md THREAT_MODEL.md THREAT_MODEL.md WINDOWS_SETUP.md WINDOWS_SETUP.md build-release.ps1 build-release.ps1 build.rs build.rs config.example.toml config.example.toml deny.toml deny.toml docker-compose.yml docker-compose.yml hsip.toml hsip.toml install.ps1 install.ps1 install.sh install.sh launch-hsip.bat launch-hsip.bat railway.toml railway.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
One binary. No cloud. No subscription. Cryptographic identity and tamper-proof audit trail for individuals, AI agents, and financial institutions.
🌐 hsip.rewired89.github.io/HSIP-1PHASE — Landing page with one-click downloads
Every key is yours. Every byte runs locally. No cloud. No subscription. Commercial use requires a license — contact sanchezleal1989@gmail.com . Read the threat model →
Windows — Download hsip-windows-x64.exe → double-click → browser opens automatically.
curl -sSf https://raw.githubusercontent.com/rewired89/HSIP-1PHASE/main/install.sh | sh
Homebrew:
brew tap rewired89/hsip https://github.com/rewired89/HSIP-1PHASE && brew install hsip
Why this exists — right now
In 2026, three things happened at once:
AI agents act on your behalf without a reliable record of what they did or who authorized it.
OpenAI, Google, and Meta serve ads inside the tools you use to think. Your prompts train their models.
Deepfakes made digital evidence meaningless — unless it carries a cryptographic signature that cannot be faked.
HSIP is the answer to all three. It runs on your hardware, signs everything with your key, and gives you a tamper-proof audit trail you own completely.
I want to...
What to run
Stop being tracked — block ads, telemetry, and surveillance across every app I use
DNS Tracker Blocker
Prove what I said — create court-admissible proof that I wrote this message at this time
Signed Messages + Audit Trail
Control my AI agents — see exactly what my AI did, revoke access instantly
AI Watch + Consent Wallet
Build privacy-respecting software — add consent infrastructure to my app or AI agent
Developer SDK →
Enterprise audit compliance — GDPR, court records, legal-grade evidence chains
Enterprise deployment →
Financial services infrastructure — MiFID II, FINRA 4511, SOX §404, DORA, SWIFT CSCF compliance
Financial Services →
Download
Platform
File
Windows
hsip-windows-x64.exe
macOS Apple Silicon
hsip-macos-arm64
macOS Intel
hsip-macos-x64
Linux
hsip-linux-x64
Windows: Double-click the .exe . It installs itself, creates a Desktop shortcut, and opens in your browser automatically.
Mac / Linux: chmod +x hsip-macos-arm64 && ./hsip-macos-arm64 — your browser opens automatically.
1. DNS Tracker Blocker — block everything, system-wide
HSIP intercepts tracking requests at the DNS level before they ever reach your machine. Not just one browser — every app you run.
Blocks Google Analytics, Facebook Pixel, Hotjar, TikTok, DoubleClick, Microsoft telemetry, and 200+ more. One click in the dashboard to turn on. Zero configuration.
The difference from browser extensions: A browser extension only protects one browser. HSIP blocks at the network level — desktop apps, background processes, every browser, all at once.
2. Signed Messages — fight deepfakes and win disputes
Every message you send through HSIP is signed with your personal Ed25519 key. The result is mathematical proof that:
This proof can be verified by anyone, in court, or by a machine. It cannot be faked.
Contract confirmation: "I confirm we agreed to these terms on March 28, 2026." — signed, timestamped, verifiable.
Dispute evidence: Produce a cryptographic receipt in seconds that proves what you said and when.
Deepfake defense: When someone claims you said something you didn't — your signed history proves otherwise.
AI command authorization: Every instruction you gave your AI agent is signed with your key. Deniability is gone — in both directions.
3. AI Watch — know exactly what your AI did
Every AI agent you connect (Claude, ChatGPT, Siri, any HTTP-capable tool) is tracked in real time:
Velocity monitoring — alerts if an agent makes an unusual number of requests
Anomaly detection — flags behavior outside normal patterns
One-click disconnect — revoke any agent's access instantly
Full signed audit trail — every action the agent took, signed and timestamped
This is the "black box recorder" for your AI. When something goes wrong, you know exactly what happened and when.
4. Consent Wallet — machine-readable access control
Instead of cookie banners you click through without reading, HSIP creates a consent layer you actually control:
See every party that has permission to contact you or access your data
See exactly what each party is allowed to do
Set time limits on consent — it expires automatically
Revoke any consent in one click, effective immediately
Third-party services that support HSIP can query your consent before acting. No permission — no access.
Every operation in HSIP — message signed, consent granted, key created, AI action logged — writes to a BLAKE3 hash-chained audit log. Tamper with any entry and the chain breaks.
Export the log at any time for legal proceedings, compliance audits, or personal records.
HSIP is cryptographic infrastructure for banks, trading desks, fintechs, and any regulated institution that needs a tamper-proof audit trail, AI agent governance, and cross-institution identity verification — without a central cloud vendor in the middle.
The client is the institution, not the retail investor. HSIP runs inside your data center (or on-premise), signs every action with your Ed25519 keypair, and produces legally defensible evidence that your systems, analysts, and AI agents did exactly what the audit trail says they did.
Why financial institutions need this now
1. AI agents act on behalf of your institution — and regulators are going to ask who authorized each action. Without a cryptographic identity attached to each agent and an append-only log of every request, you cannot answer that question. HSIP assigns every AI agent its own Ed25519 keypair, logs every action it takes, and lets you revoke its access in milliseconds.
2. MiFID II Article 25 and FINRA Rule 4511 require you to prove what your systems did, when, and on whose authority. A log in a database is not proof — it can be altered. A BLAKE3 hash-chained audit log is proof. Tamper with any entry and the chain breaks, detectable by any party.
3. Open Banking (PSD2) mandates machine-readable, time-bounded consent. HSIP's Consent Wallet generates exactly that: a cryptographically signed grant scoped to a specific action, automatically expiring, revocable in real time. No more cookie banners your compliance team can't evidence.
4. Inter-institution trust is broken. When a message arrives from a counterparty, how do you verify it wasn't altered in transit? HSIP's Federated Trust layer lets institutions exchange Ed25519 verify keys out-of-band (email, secure channel) and then verify any future message cryptographically — no central registry, no PKI vendor, no single point of failure.
5. DORA and SWIFT CSCF require you to detect and respond to anomalous AI or automated system behavior. HSIP's velocity monitoring flags agents exceeding 100 requests/minute and auto-revokes access at 1,000 requests/minute — with a signed audit entry at every step.
Regulation
What HSIP covers
SOX §404
Append-only BLAKE3 hash-chained audit log. Every control action signed with Ed25519. Exportable for auditors.
FINRA Rule 4511
Six-year tamper-evident record retention. API endpoint for bulk audit export. Signature chain proves no entry was altered.
MiFID II Art. 25
Per-trade authorization signed with institutional Ed25519 key. Timestamp + signature = defensible suitability record.
PSD2 / Open Banking
Machine-readable consent grants with scope, expiry, and revocation. POST /v1/consent/grant with expires_in_seconds .
GDPR Art. 7
Cryptographically signed consent with documented scope. DELETE /v1/tenant/erase for right-to-erasure. Audit log proves consent was active at time of processing.
DORA
AI agent velocity monitoring, anomaly detection, auto-revocation. Incident response via DELETE /v1/keys/:id . All events in signed audit trail.
SWIFT CSCF
Ed25519 message authentication prevents unauthorized instruction injection. Federated trust keys verified per counterparty. No shared secrets.
ISO 20022
Signed payment messages with Ed25519. Verifiable by any counterparty holding the institution's public key. Non-repudiation by construction.
AI agent governance for financial institutions
Every AI system your institution deploys — trading algorithms, document processors, customer-facing chatbots, internal assistants — gets its own Ed25519 keypair registered in HSIP.
# Register a trading algorithm as a governed AI agent
hsip agent register " algo-trading-v2 " --expires-days 90
# List all active agents and their request velocity
hsip agent list
# Immediately revoke an agent that's behaving unexpectedly
hsip agent revoke " algo-trading-v2 "
What you get for each agent:
Unique Ed25519 keypair — every action it signs is traceable to that specific agent, not just "the system"
Velocity monitoring — requests > 100/min trigger an anomaly audit entry; > 1,000/min triggers automatic revocation
Full signed audit trail — ev

[truncated]
