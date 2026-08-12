---
source: "https://github.com/ivan-magda/swift-claw"
hn_url: "https://news.ycombinator.com/item?id=49271289"
title: "I built a personal AI agent in Swift on top of what macOS ships"
article_title: "GitHub - ivan-magda/swift-claw: Your AI. Your machine. Always on. A personal assistant in one pure-Swift binary. · GitHub"
author: "strong-self"
captured_at: "2026-08-12T12:45:02Z"
capture_tool: "hn-digest"
hn_id: 49271289
score: 1
comments: 0
posted_at: "2026-08-12T12:23:50Z"
tags:
  - hacker-news
  - translated
---

# I built a personal AI agent in Swift on top of what macOS ships

- HN: [49271289](https://news.ycombinator.com/item?id=49271289)
- Source: [github.com](https://github.com/ivan-magda/swift-claw)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T12:23:50Z

## Translation

タイトル: macOS に同梱されているものに基づいて Swift でパーソナル AI エージェントを構築しました
記事のタイトル: GitHub - ivan-magda/swift-claw: あなたの AI。あなたのマシン。常にオン。 1 つの純粋な Swift バイナリのパーソナル アシスタント。 · GitHub
説明: あなたの AI。あなたのマシン。常にオン。 1 つの純粋な Swift バイナリのパーソナル アシスタント。 - イワン・マグダ/スウィフト・クロー

記事本文:
GitHub - ivan-magda/swift-claw: あなたの AI。あなたのマシン。常にオン。 1 つの純粋な Swift バイナリのパーソナル アシスタント。 · GitHub
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
イワン・マグダ
/
素早い爪
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,038 コミット 1,038 コミット .claude/ skill/ verify .claude/ skill/ verify .github .github ソース ソース テスト テスト デプロイ デプロイ ドキュメント ドキュメント スクリプト スクリプト .env.example .env.e

サンプル .gitattributes .gitattributes .gitignore .gitignore .swift-format .swift-format .swiftlint.yml .swiftlint.yml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスPackage.resolved Package.resolved Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
常時稼働のパーソナル AI アシスタント。所有するハードウェア上の純粋な Swift デーモン 1 つ。
clawd はプライベート Telegram ボットを選択した LLM と組み合わせます。それはあなたのことを覚えています
通知して、スケジュールされたプロアクティブなタスクを実行します。重要なツール呼び出しは、
承認。保存されているものはすべて、自分のマシン上の 1 つのディレクトリ (SQLite) に保存されます。
データベース、暗号化された秘密のエンベロープ、および手動で編集する Markdown ファイル。
本物のテレグラムチャット。回答はライブ メッセージの下書きとしてストリーミングされます。 /stop はキャンセルします
ターン、/new は新しいセッションを開始し、clawd はデバイス上で音声メモを書き写します
(macOS 26)、お使いのモデルが送信した写真を閲覧できる場合、それが表示されます。
耐久性のあるメモリ。あなたが確認した事実は SQLite に保持され、clawd はそれを次のように呼び出します。
重要性と最新性。ワークスペース マークダウン ファイルには、プロフィール、メモ、日次情報が保存されます。
ログ、会話履歴は全文検索可能です。
一度書いたスキル。 Skills/<name>/SKILL.md ファイルは、コンテキスト内でその名前として表示されます。
そしてそれをいつ使用するかについての 1 行。タスクが一致すると、clawd は本体をロードしてフォローします。
再度貼り付けるように要求するのではなく、手順を実行してください。 /スキルを送信してすべてを確認してください
受け入れられたスキルとスキャナーが拒否した各ファイル。
いつでもプロアクティブに。 「平日毎日 07:00」は、1 回につき 1 回起動するようにスケジュールします。
再起動と DST 変更の間で発生し、オプトイン ハートビートは静かな時間を尊重します。
ポリシー エンジンの背後にあるツール。 web_fetch は SSRF ゲートの背後にあります。書いてコードを書く
処刑w

Telegram での明示的なタップによる承認を待ちます。 Clawd はポリシーを強制します
コードを作成し、受信コンテンツを命令としてではなくデータとして扱います。
サンドボックス化されたコードの実行。信頼できないコードはリクエストごとに新しい使い捨て VM で実行されます
(macOS 26 arm64、デフォルトではオフ)。
MCP サーバーからのツール。サーバーをリストし、そのトークンを暗号化して保存し、そのツールが参加します
Clawd が持つ最も信頼性の低いツールとしての組み込み。デフォルトでは、呼び出しは尋ねます。マークを付けることができます
名前付きツールセーフですが、抽出ゲートには依然として承認が必要な場合があります。あなただけが追加できます
サーバーを変更するか、公開するものを変更します。
ご自身のモデルをご持参ください。 OpenAI 互換のエンドポイントはすべて動作し、認証ログインをクロードします。
ChatGPT サブスクリプションで対象となるモデルを実行できます。
バイナリが 1 つ。 Telegram ロングポーリングから SQLite まで、厳密な同時実行性を備えた Swift 6。
ファイルの書き込みでは、応答するまで実行が一時停止されます。カード上のすべてのフィールドはデーモン自身のものです
アクションの記録: シンボリックリンク後のターゲット パスと .. 解像度、サイズ、プレビュー
内容。 「拒否」をタップすると、clawd は何も書き込みません。
カール -fsSL https://raw.githubusercontent.com/ivan-magda/swift-claw/main/install.sh |しー
すべては ~/.swift-claw に配置され、sudo は使用されません。スクリプトはすべてのダウンロードを検証します
リリース チェックサムと照合し、サービス ファイルをステージングして、次の手順を出力します。
リリースをカールでピン留めする … | CLAWD_VERSION=v0.2.0 sh 、読んでください
スクリプト ソースを最初に実行するか、手動ルートに従います。
docs/INSTALL.md (macOS 15+ arm64、glibc 2.38+ を搭載した Linux x86_64)。
または、Swift 6.3 ツールチェーンを使用してソースからビルドします (Linux には libsqlite3-dev が必要です)。
git clone https://github.com/ivan-magda/swift-claw.git && cd swift-claw
迅速なビルド -c リリース
sudo install -m755 .build/release/clawd /usr/local/bin/clawd
クイックスタート
@BotFather からボット トークンを取得します ( /newbot を送信)。
~/.swift-claw/clawd.env を編集: トークンを設定します

および LLM プロバイダー
( CLAW_LLM_BASE_URL 、 CLAW_LLM_MODEL 、 CLAW_LLM_API_KEY )。
設定をロードし、保存時のシークレットを暗号化します。
set -a && 。 ~/.swift-claw/clawd.env && +a && Clawd Secret シールを設定します
フォアグラウンドで一度挨拶します: clawd run 、次に /start をボットに送信します。の
拒否すると数字のIDが表示されます。 clawd.env で CLAW_ALLOWLIST=<id> として設定し、
Ctrl+C。
健全性を確認してサービスを開始します。
set -a && 。 ~/.swift-claw/clawd.env && set +a && clawd Doctor を実行し、start を実行します
コマンドドクターのプリント。
ChatGPT サブスクリプション ルートとトラブルシューティングを含む完全なウォークスルーは、次のとおりです。
docs/GETTING_STARTED.md にあります。
swift-claw は、サービスを提供するのはあなただけであると想定しています。
デフォルトの拒否。許可リストに登録された Telegram ID のみが会話を取得します。爪は拒否します
他の全員は、送信者自身の数値 ID で /start に応答します。
それらを許可リストに登録します。 CLAW_ALLOWLIST は追加のみなので、ID を取り消すことはその ID を削除することを意味します。
データベースからの行 ( 詳細 )。
秘密は保存時に暗号化されます。 Clawd Secret シールはボット トークンと API キーをラップします。
AES-GCM エンベロープ。プレーンテキストの環境シークレットは、開発フォールバックとして引き続き利用可能です。
ブートするたびに警告します。
承認は耐久性があり、偽造することはできません。ファイル書き込み、メモリ書き込み、およびコード
Telegram で [承認] をタップするまで、実行は永続ステート マシンに一時停止されます。あ
偽造またはサードパーティのコールバックは承認できず、保留中の承認は期限切れになり拒否されます。
即時注射が含まれています。メッセージ、Web コンテンツ、ツール出力、および保存されたメモリ
コンテキストを信頼できないデータとして入力します。セッションが両方の信頼できないコンテンツを取り込んだ場合
プライベート ファイルをコンテキストに取り込み、任意の URL を取得するには、
承認。 clawd は LLM と検索プロバイダーを構成に固定しますが、モデルはそれを行うことができません
それらをリダイレクトします。
完全なモデルは docs/ARCHITECTURE.md (§12) にあります。報告するには

ある
脆弱性については、 SECURITY.md を参照してください。
ペルソナと動作は ~/.swift-claw/workspace/ の下の Markdown ファイルに存在します。
MCP サーバーは ~/.swift-claw/mcp.yaml に配置され、そのトークンは暗号化されて保存されます。
clawd mcp セットトークン 。他のランタイム ノブは環境変数です: モデル ルート
( CLAW_LLM_MODEL )、オプションのフォールバック ルート ( CLAW_LLM_FALLBACK_MODEL 、そうでない限りオフ)
設定します）、米ドルの予算、スケジュールと静かな時間、音声ロケール、サンドボックスの制限。
.env.example はすべての変数を文書化します。
docs/CUSTOMIZATION.md がガイドです。
あなたはそうしたいです
読む
エンドツーエンドで設定する
docs/GETTING_STARTED.md
インストール、アップデート、またはアンインストール
docs/INSTALL.md
それをあなたのものにしてください
docs/カスタマイズ.md
サービスとして実行する
docs/INSTALL.md
ローカルで開発してテストする
docs/LOCAL_DEV.md
デザインを理解する
docs/ARCHITECTURE.md
脆弱性を報告する
セキュリティ.md
貢献する
貢献は大歓迎です。問題を開いて、事前に考えていることについて話し合います
プルリクエストを送信する。 COTRIBUTING.md に詳細が記載されています。
lint/テストゲート。
あなたのAI。あなたのマシン。常にオン。 1 つの純粋な Swift バイナリのパーソナル アシスタント。
swift-whisper-buddy.lovable.app トピックス
Readme MIT ライセンス
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Your AI. Your machine. Always on. A personal assistant in one pure-Swift binary. - ivan-magda/swift-claw

GitHub - ivan-magda/swift-claw: Your AI. Your machine. Always on. A personal assistant in one pure-Swift binary. · GitHub
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
ivan-magda
/
swift-claw
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,038 Commits 1,038 Commits .claude/ skills/ verify .claude/ skills/ verify .github .github Sources Sources Tests Tests deploy deploy docs docs scripts scripts .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .swift-format .swift-format .swiftlint.yml .swiftlint.yml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Package.resolved Package.resolved Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md install.sh install.sh View all files Repository files navigation
Your always-on personal AI assistant. One pure-Swift daemon on hardware you own.
clawd pairs a private Telegram bot with the LLM of your choice. It remembers what you
tell it and runs scheduled and proactive tasks. Consequential tool calls wait for your
approval. Everything it keeps stays in one directory on your own machine: a SQLite
database, encrypted secret envelopes, and Markdown files you edit by hand.
A real Telegram chat. Answers stream in as live message drafts. /stop cancels a
turn, /new starts a fresh session, clawd transcribes voice notes on-device
(macOS 26), and it looks at photos you send if your model can see them.
Durable memory. Facts you confirm persist in SQLite, and clawd recalls them by
importance and recency. Workspace Markdown files hold your profile, notes, and daily
logs, and conversation history is full-text searchable.
Skills you write once. A skills/<name>/SKILL.md file shows up in context as its name
and one line about when to use it; when a task matches, clawd loads the body and follows
your procedure instead of asking you to paste it again. Send /skills to see every
accepted skill and each file the scanner rejected.
Proactive, on your clock. "Every weekday at 07:00" schedules fire once per
occurrence across restarts and DST changes, and an opt-in heartbeat respects quiet hours.
Tools behind a policy engine. web_fetch sits behind an SSRF gate; writes and code
execution wait for an explicit tap-to-approve in Telegram. clawd enforces policy in
code and treats inbound content as data, never as instructions.
Sandboxed code execution. Untrusted code runs in a fresh disposable VM per request
(macOS 26 arm64, off by default).
Tools from MCP servers. List a server, store its token encrypted, and its tools join
the built-ins as the least-trusted tools clawd has. Calls ask by default; you may mark a
named tool safe, but the exfiltration gate can still require approval. Only you can add a
server or change what it exposes.
Bring your own model. Any OpenAI-compatible endpoint works, and clawd auth login
can run an eligible model on a ChatGPT subscription.
One binary. Swift 6 with strict concurrency, from the Telegram long-poll down to SQLite.
A file write suspends the run until you answer. Every field on the card comes from the daemon's own
record of the action: the target path after symlink and .. resolution, the size, and a preview of
the content. Tap Deny and clawd writes nothing.
curl -fsSL https://raw.githubusercontent.com/ivan-magda/swift-claw/main/install.sh | sh
Everything lands in ~/.swift-claw , with no sudo. The script verifies every download
against the release checksums, stages the service files, and prints the next steps.
Pin a release with curl … | CLAWD_VERSION=v0.2.0 sh , read the
script source first, or follow the manual route in
docs/INSTALL.md (macOS 15+ arm64; Linux x86_64 with glibc 2.38+).
Or build from source with a Swift 6.3 toolchain (Linux needs libsqlite3-dev ):
git clone https://github.com/ivan-magda/swift-claw.git && cd swift-claw
swift build -c release
sudo install -m755 .build/release/clawd /usr/local/bin/clawd
Quick start
Get a bot token from @BotFather (send /newbot ).
Edit ~/.swift-claw/clawd.env : set the token and your LLM provider
( CLAW_LLM_BASE_URL , CLAW_LLM_MODEL , CLAW_LLM_API_KEY ).
Load the config and encrypt your secrets at rest:
set -a && . ~/.swift-claw/clawd.env && set +a && clawd secrets seal
Say hello once in the foreground: clawd run , then send /start to your bot. The
refusal shows your numeric ID; set it as CLAW_ALLOWLIST=<id> in clawd.env , then
Ctrl-C.
Check health and start the service:
set -a && . ~/.swift-claw/clawd.env && set +a && clawd doctor , then run the start
command doctor prints.
The full walkthrough, including the ChatGPT-subscription route and troubleshooting, is
in docs/GETTING_STARTED.md .
swift-claw assumes you are the only person it serves.
Default-deny. Only allowlisted Telegram IDs get a conversation. clawd refuses
everyone else, and answers /start with the sender's own numeric ID so you can
allowlist them. CLAW_ALLOWLIST only ever adds, so revoking an ID means deleting its
row from the database ( details ).
Secrets encrypted at rest. clawd secrets seal wraps the bot token and API keys in
an AES-GCM envelope. Plaintext env secrets remain available as a dev fallback that
warns on every boot.
Approvals are durable and unforgeable. File writes, memory writes, and code
execution suspend into a durable state machine until you tap Approve in Telegram. A
forged or third-party callback cannot approve, and pending approvals expire to deny.
Prompt injection contained. Messages, web content, tool output, and stored memory
enter the context as untrusted data. Once a session has both ingested untrusted content
and pulled your private files into context, fetching an arbitrary URL also needs your
approval. clawd pins your LLM and search providers in config, and the model cannot
redirect them.
The full model is in docs/ARCHITECTURE.md (§12). To report a
vulnerability, see SECURITY.md .
Persona and behavior live in Markdown files under ~/.swift-claw/workspace/ :
MCP servers go in ~/.swift-claw/mcp.yaml , with their tokens stored encrypted by
clawd mcp set-token . Other runtime knobs are environment variables: the model route
( CLAW_LLM_MODEL ), an optional fallback route ( CLAW_LLM_FALLBACK_MODEL , off unless you
set it), USD budgets, schedules and quiet hours, voice locales, sandbox limits.
.env.example documents every variable;
docs/CUSTOMIZATION.md is the guide.
You want to
Read
Set it up end to end
docs/GETTING_STARTED.md
Install, update, or uninstall
docs/INSTALL.md
Make it yours
docs/CUSTOMIZATION.md
Run it as a service
docs/INSTALL.md
Develop and test locally
docs/LOCAL_DEV.md
Understand the design
docs/ARCHITECTURE.md
Report a vulnerability
SECURITY.md
Contributing
Contributions are welcome. Open an issue to discuss what you have in mind before
sending a pull request; CONTRIBUTING.md has the details and the
lint/test gate.
Your AI. Your machine. Always on. A personal assistant in one pure-Swift binary.
swift-whisper-buddy.lovable.app Topics
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
