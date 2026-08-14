---
source: "https://github.com/choiyounggi/cliclaw"
hn_url: "https://news.ycombinator.com/item?id=49294922"
title: "Show HN: Control Claude Code, Codex, Pi and Gemini CLI from Telegram"
article_title: "GitHub - choiyounggi/cliclaw: A Telegram bot that drives four local coding CLIs — Claude Code, Codex, Pi, and Gemini — from your phone. macOS daemon with per-chat per-agent sessions, a confirm gate for dangerous commands, response streaming, image attachments, auto-installed launchd, and corporate-T\n[truncated]"
author: "dch0202"
captured_at: "2026-08-14T05:25:26Z"
capture_tool: "hn-digest"
hn_id: 49294922
score: 2
comments: 0
posted_at: "2026-08-14T05:07:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Control Claude Code, Codex, Pi and Gemini CLI from Telegram

- HN: [49294922](https://news.ycombinator.com/item?id=49294922)
- Source: [github.com](https://github.com/choiyounggi/cliclaw)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T05:07:01Z

## Translation

タイトル: Show HN: Telegram からの Control Claude Code、Codex、Pi および Gemini CLI
記事のタイトル: GitHub - choiyounggi/cliclaw: 携帯電話から 4 つのローカル コーディング CLI (Claude Code、Codex、Pi、Gemini) を駆動する Telegram ボット。エージェントごとのチャットごとのセッションを備えた macOS デーモン、危険なコマンドの確認ゲート、応答ストリーミング、画像添付、自動インストールされる launchd、corporate-T
[切り捨てられた]
説明: 携帯電話から 4 つのローカル コーディング CLI (Claude Code、Codex、Pi、Gemini) を駆動する Telegram ボット。エージェントごとのチャットごとのセッション、危険なコマンドの確認ゲート、応答ストリーミング、イメージ添付ファイル、自動インストールされる launchd、および企業 TLS (Zscaler) 自動検出を備えた macOS デーモン。 -
[切り捨てられた]

記事本文:
GitHub - choiyounggi/cliclaw: 携帯電話から 4 つのローカル コーディング CLI (Claude Code、Codex、Pi、Gemini) を駆動する Telegram ボット。エージェントごとのチャットごとのセッション、危険なコマンドの確認ゲート、応答ストリーミング、イメージ添付ファイル、自動インストールされる launchd、および企業 TLS (Zscaler) 自動検出を備えた macOS デーモン。 · GitHub
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
チェヨンギ
/
クリッククロー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 操作

詳細アクション メニュー フォルダーとファイル
41 コミット 41 コミット .github/ workflows .github/ workflows __tests__ __tests__ bin bin lib lib サンプル サンプル .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DEVELOPMENT.md DEVELOPMENT.md ライセンス ライセンス README.ko.md README.ko.md README.md README.md SECURITY.md SECURITY.md bot.ts bot.ts cli.ts cli.ts config.example.json config.example.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
4 つのローカル コーディング CLI を駆動できる単一のデーモン ( Claude Code ·
Codex · Pi · Gemini ) を Telegram から使用し、チャットごとに切り替えます。
すべてのチャットに対して独立したエージェントごとのセッションを維持し、確認メッセージを送信します。
危険なコマンドのゲート、応答ストリーミング、画像添付ファイルの処理、
企業の TLS インターセプター (Zscaler など) の自動検出。
必要な CLI のみをインストールします - 不足しているエージェントは自動的に削除されます
アクティブリストから。
チャット UX (ユーザー メッセージ、エラー、 /help ) は現在韓国語ファーストです。の
ボット自体はどちらの方法でも正常に動作します。
Bun 1.x —curl -fsSL https://bun.sh/install |バッシュ
少なくとも 1 つのコーディング CLI がインストールされ、ログインしていること (ボットは認証された子プロセスを生成するだけです)
クロード コード: npm install -g @anthropic-ai/claude-code
コーデックス: npm install -g @openai/codex
Pi: npm install -g @earendil-works/pi-coding-agent
Gemini: npm install -g @google/gemini-cli
# パン (推奨)
パン追加 -g @younggichoi/cliclaw
# または npm
npm install -g @younggichoi/cliclaw
パッケージはスコープ付きの名前 ( @younggichoi/cliclaw ) で公開されますが、
インストールされているコマンドは単に clilaw です。
クリッククローの初期化
ガイド付きの 5 つのステップ:
cliclaw のセットアップへようこそ。
ステップ 1/5 — Telegram ボット トークン
Telegram の @BotFather (/newbot) から取得してください。
ボットトークン:

1234:ABC...
✓ @yourbotname (id=...) が確認されました
ステップ 2/5 — インストールされているコーディング エージェントを検出する
✓ クロード 2.1.139 (クロード コード) @ ~/.nvm/.../bin/claude
✓ コーデックス 1.0.0 @ /opt/homebrew/bin/codex
✓ pi が不明 @ /opt/homebrew/bin/pi
✓ ジェミニ 0.42.0 @ ~/.nvm/.../bin/gemini
デフォルトのエージェント? [クロード] (クロード/コーデックス/パイ/ジェミニ): クロード
ステップ 3/5 — Telegram アカウントを認証する
Telegram を開いて、メッセージを @yourbotname に送信してください。
最大 5 分間待機します...Ctrl+C を押して中止します。
✓ user_id=123456789 から受信
この Telegram ユーザーを承認しますか? [はい/いいえ] はい
ステップ 4/5 — 企業 TLS インターセプター (オプション)
$NODE_EXTRA_CA_CERTS: /path/to/Zscaler.pem
この CA 証明書をボットの LaunchAgent 環境に適用しますか? [はい/いいえ] はい
ステップ 5/5 — ログイン時の自動起動 (launchd)
LaunchAgent をインストールすると、ログイン時にボットが自動的に起動しますか? [はい/いいえ] はい
✓ インストール済み ~/Library/LaunchAgents/com.alice.cliclaw.plist
✓ ボットが開始されました。
すべて準備完了です。
ログ: tail -f ~/.cliclaw/logs/bot.log
テスト: Telegram で /status を送信します。
それでおしまい。画面ロックまたは再起動後、ボットは自動的に再起動します。
コマンド
何をするのか
クリッククローの初期化
対話型セットアップ (トークン、エージェント検出、テレグラム ID キャプチャ、CA、launchd)
クリッククロースタート
フォアグラウンドでボットを実行します (テスト用)
cliclaw インストール起動
LaunchAgent をインストールします (config.json から launchd.extraEnv を自動適用します)
clilaw アンインストール-起動
LaunchAgent を削除する
クリクロードクター
パス、エージェント、および plist ステータスを確認する
クリッククローヘルプ
ヘルプ
チャットコマンド (Telegram 内)
コマンド
何をするのか
/claude /codex /pi /gemini
このチャットのアクティブなエージェントを切り替える (アンインストールされたエージェントは非表示になります)
/ステータス
エージェントごとのセッション ステータス + 進行中のジョブ
/健康
ボット システムのステータス (稼働時間、メモリ、ログ サイズ、チャット/ジョブ数)
/停止
このチャットの実行中のジョブをキャンセルします (5 秒後に SIGTERM → SIGKILL)
/リセット
廃棄のみ

現在アクティブなエージェントのセッション
/すべてリセット
このチャット内のすべてのエージェント セッションを破棄します
/安全性
セーフティモードのステータスを表示 — /safety オン / /safety オフに切り替えます
/スタート /ヘルプ
ヘルプ
その他のテキスト/写真
プロンプトとしてアクティブなエージェントに送信されます (写真がダウンロードされ、そのパスがプロンプトの前に追加されます)
エージェントを切り替えると、古いセッションがそのまま維持されます。戻って続行します。送信中
実行中のジョブとのチャットへの新しいプロンプトは拒否されます (/stop または wait を使用します)。
すべての状態は ~/.cliclaw/ の下に分離されます。
~/.cliclaw/
§── config.json # モード 600;トークン、ホワイトリスト、launchd extraEnv
§── safety.json # 永続化された /safety on|off 状態
§── session.json # チャットごとのアクティブ エージェントのメタデータ
§── セッション/ # チャットごとのコーデックス / pi / gemini ディレクトリ
§── workspace/ # エージェントの共有 cwd (サンドボックス)
│ §── .claude/settings.json # 危険なコマンドフック + セーフティが ON の場合のルールを拒否する
│ └── Uploads/<chatId>/ # Telegram 写真のダウンロード
§── ログ/
│ §── bot.log # トークンの自動マスキングが適用されました
│ §── bot.err # launchd stderr
│ └── Audit.jsonl # 監査ログ (決定、安全状態)
└── .sock/ # 確認ゲート IPC
状態ディレクトリは CLICLAW_HOME 環境変数を使用して移動できます。
CLICLAW_HOME= ~ /my-bot cliclaw init
特長
1. エージェントパスの自動検出
config.json には絶対パスはありません。起動時に、エージェントは 3 つのパスで検出されます。
~/.local/bin 、 ~/.claude/local 、 /usr/local/bin 、 /opt/homebrew/bin
$NVM_DIR または ~/.nvm の下の最新ノードの bin/<cmd>
ログインシェルのコマンド -v <cmd> (.zshrc がロードされた PATH)
検出されなかったエージェントは正常にスキップされます。ボットは 4 つのサブセットのいずれでも正常に動作します。
2. セーフティモード ( /safety オン · /safety オフ )
危険な Bash コマンド ( rm -rf 、 git Push --force 、 DROP 、 kubectl delete 、 AWS

delete-* 、 sudo 、curl|sh 、 ssh prd-* など) は、Telegram インライン キーボード [✅ 許可] [❌ 拒否] 経由で再確認されます。応答がない場合は、自動拒否を意味します。
Claude の読み取りツールは、機密ファイルを拒否します: ~/.ssh/** 、 ~/.aws/** 、 ~/.gnupg/** 、 ~/.netrc 、 ~/.npmrc 、 **/.env* 、 **/*.pem 、 **/id_rsa* 、 **/id_ed25519* 、 ./secrets/** 。
confirmGate.extraPatterns を介して独自の正規表現を追加します。
OFF : 環境にすでに外部ガード ( pre-bash-guard 、
EDR、…) とボットの確認プロンプトが冗長に感じられるので、1 行でオフにします
テレグラムで。拒否ルールは一緒に無効になります。すべての IPC リクエストは依然として
logs/audit.jsonl に決定:allow、理由:safety_off として記録されます。
この状態は再起動後も $CLICLAW_HOME/safety.json に保持されます。
3. レスポンスストリーミング（クロード）
--include-partial-messages と live-updates からの text_delta を使用します。
editMessageText 、1.5秒でデバウンス。 3800 文字を超えると、新しいメッセージにロールオーバーされます。
電報の写真・画像ドキュメントは以下にダウンロードされます。
workspace/uploads/<chatId>/<msgId>.<ext> とパスがプロンプトの前に追加されます。
Claude : --permission-mode bypassPermissions で実行します。危険な Bash は確認ゲートによって捕らえられ、機密ファイルはセーフティ モードの拒否ルールによって捕らえられます。
Codex : デフォルトでは、sandbox=workspace-write。危険フルアクセスは決して使用しないでください。
Gemini : デフォルトでは、approvalMode=auto_edit (自動承認された破壊的なコマンド プロンプトを編集します)。より自律的な yolo またはより保守的なデフォルト/プランが利用可能です。
6. 企業 TLS インターセプターの自動検出
Zscaler / Forticlient / Cisco Umbrella が HTTPS をインターセプトする場合、ノードはできません
Telegram の証明書を信頼しないと、ボットは実行できません。 cliclaw 初期化のステップ 4
$NODE_EXTRA_CA_CERTS または launchctl getenv NODE_EXTRA_CA_CERTS を自動検出します。
と尋ねられ、それを config.json の launchd.extraEnv に永続化します。後ほど
インスタ

all-launchd はそれを plist に自動的に焼き付けます。
logs/bot.log / bot.err に書き込まれる内容はすべて事前に編集されます。
Telegram ボット トークン ( \d{8,}:[A-Za-z0-9_-]{30,} )
ライブ config.token の完全一致
Time Machine バックアップ、EDR、ショルダー サーフィンなどに対しても同様に防御します。
ユーザー向けのメッセージ、エラー、ヘルプはすべて韓国語です (英語の PR も歓迎します)。
cliclaw init のステップ 5 で「はい」と答える:
~/Library/LaunchAgents/com.<username>.cliclaw.plist を作成します (企業 CA が組み込まれています)
launchctl bootstrap gui/$UID <plist> を介してすぐにロードして開始します。
以降、ログイン/起動/クラッシュ時に自動再起動します
stdout → ~/.cliclaw/logs/bot.log 、stderr → bot.err
# 停止 (次回ログイン時に自動再起動)
launchctl kill SIGTERM gui/ $UID /com. < ユーザー名 > .cliclaw
# 完全に無効化します（自動再起動もありません）
clilaw アンインストール-起動
# 再有効化 (config.json から launchd.extraEnv を自動適用)
cliclaw インストール起動
セキュリティ
ボット トークン = インストールされているすべてのエージェントへのリモート シェル。漏洩した場合は、直ちに BotFather で /revoke してください。
空の allowedUserIds はすべてのメッセージを拒否します (フェールクローズ)。
config.json をモード 600 のままにしておきます (init によって自動的に設定されます)。
決してconfirmGate.enabled: false を設定したり、コーデックスサンドボックスをdanger-full-accessに切り替えたりしないでください。
Gemini の ApprovalMode を yolo にオプトすると、すべてのツールが自動承認されます。十分に理解した上で使用してください。
疑わしい場合は、/safety on を指定すると、拒否ルールが即座に再アクティブ化されます。
cliclaw init フローを使用せずにインストールするには:
git clone https://github.com/choiyounggi/cliclaw.git
CD クリクロー
バンインストール
mkdir -p ~ /.cliclaw
cp config.example.json ~ /.cliclaw/config.json
chmod 600 ~ /.cliclaw/config.json
# config.json にトークンと allowedUserIds を入力し、
バンランbot.ts
テスト
危険なパターンは正規表現に基づいており、100% 分類することは不可能です。ポリシーの所有者はあなたです。
誰でもない

-Codex / Pi / Gemini のテキスト ストリーミング (構造化イベントなし、または統合されていない)。
Gemini の危険なコマンドは、独自の ApprovalMode のみに依存します (bash-confirm IPC は統合されていません)。
フックは、ユーザーの決定を待つ間、IPC を保持します。
音声/ファイルの添付はありません (写真のみ)。
同じチャット内の同時メッセージは拒否されます (/stop または wait)。
バージョンごとの変更は GitHub Releases で公開されています。
リリース自動化 (メンテナ)
# 1) バージョンを上げます (コミット + タグを自動作成します)
npm バージョンのパッチ番号またはマイナー/メジャー
# 2) コミット+タグをプッシュ
git Push --follow-tags
次に、GitHub Web UI で、「新しいリリースをドラフトする」→タグを選択→公開します。
解放する。 .github/workflows/publish.yml は自動的に実行されます。
npm public --access public 。ワークフローは最初にリリース タグを検証します
package.json バージョンと一致しますが、不一致により公開されずに失敗します。
1 回限りの前提条件: リポジトリ設定 → シークレットと変数 → アクション →
2FA バイパスが可能な npm トークンを使用して NPM_TOKEN を登録します。
https://www.npmjs.com/settings/younggichoi/tokens/new
Granular Access Token または Classic Automation トークン (2FA バイパス付き) を発行します。
npm_… トークンを GitHub Actions シークレット NPM_TOKEN として追加します
強化オプション: npm Trusted Publishing (OIDC) に切り替えると、トークンはまったく必要ありません。
https://www.npmjs.com/package/@younggichoi/cliclaw/access で、 Trusted Publisher → GitH を追加します。

[切り捨てられた]

## Original Extract

A Telegram bot that drives four local coding CLIs — Claude Code, Codex, Pi, and Gemini — from your phone. macOS daemon with per-chat per-agent sessions, a confirm gate for dangerous commands, response streaming, image attachments, auto-installed launchd, and corporate-TLS (Zscaler) auto-detection. -
[truncated]

GitHub - choiyounggi/cliclaw: A Telegram bot that drives four local coding CLIs — Claude Code, Codex, Pi, and Gemini — from your phone. macOS daemon with per-chat per-agent sessions, a confirm gate for dangerous commands, response streaming, image attachments, auto-installed launchd, and corporate-TLS (Zscaler) auto-detection. · GitHub
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
choiyounggi
/
cliclaw
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
41 Commits 41 Commits .github/ workflows .github/ workflows __tests__ __tests__ bin bin lib lib samples samples .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DEVELOPMENT.md DEVELOPMENT.md LICENSE LICENSE README.ko.md README.ko.md README.md README.md SECURITY.md SECURITY.md bot.ts bot.ts cli.ts cli.ts config.example.json config.example.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
A single daemon that lets you drive four local coding CLIs ( Claude Code ·
Codex · Pi · Gemini ) from Telegram, switching between them per chat.
It keeps an independent per-agent session for every chat, and ships a confirm
gate for dangerous commands, response streaming, image-attachment handling, and
auto-detection of corporate TLS interceptors (Zscaler, etc.).
Install only the CLIs you want — missing agents are automatically dropped
from the active list.
The chat UX (user messages, errors, /help ) is currently Korean-first; the
bot itself works fine either way.
Bun 1.x — curl -fsSL https://bun.sh/install | bash
At least one of the coding CLIs installed and logged in (the bot merely spawns authenticated child processes)
Claude Code: npm install -g @anthropic-ai/claude-code
Codex: npm install -g @openai/codex
Pi: npm install -g @earendil-works/pi-coding-agent
Gemini: npm install -g @google/gemini-cli
# Bun (recommended)
bun add -g @younggichoi/cliclaw
# or npm
npm install -g @younggichoi/cliclaw
The package publishes under a scoped name ( @younggichoi/cliclaw ), but the
installed command is simply cliclaw .
cliclaw init
Five steps, guided:
Welcome to cliclaw setup.
Step 1/5 — Telegram bot token
Get one from @BotFather (/newbot) on Telegram.
Bot token: 1234:ABC...
✓ @yourbotname (id=...) verified
Step 2/5 — Detect installed coding agents
✓ claude 2.1.139 (Claude Code) @ ~/.nvm/.../bin/claude
✓ codex 1.0.0 @ /opt/homebrew/bin/codex
✓ pi unknown @ /opt/homebrew/bin/pi
✓ gemini 0.42.0 @ ~/.nvm/.../bin/gemini
Default agent? [claude] (claude/codex/pi/gemini): claude
Step 3/5 — Authorize your Telegram account
Open Telegram and send any message to @yourbotname now.
Waiting up to 5 minutes... press Ctrl-C to abort.
✓ Received from user_id=123456789
Authorize this Telegram user? [Y/n] y
Step 4/5 — Corporate TLS interceptor (optional)
$NODE_EXTRA_CA_CERTS: /path/to/Zscaler.pem
Apply this CA certificate to the bot's LaunchAgent environment? [Y/n] y
Step 5/5 — Auto-start at login (launchd)
Install LaunchAgent so the bot starts automatically on login? [Y/n] y
✓ Installed ~/Library/LaunchAgents/com.alice.cliclaw.plist
✓ Bot started.
All set.
Logs: tail -f ~/.cliclaw/logs/bot.log
Test: send /status in Telegram.
That's it. The bot restarts automatically after screen lock or a reboot.
Command
What it does
cliclaw init
Interactive setup (token, agent detection, telegram-id capture, CA, launchd)
cliclaw start
Run the bot in the foreground (for testing)
cliclaw install-launchd
Install the LaunchAgent (auto-applies launchd.extraEnv from config.json)
cliclaw uninstall-launchd
Remove the LaunchAgent
cliclaw doctor
Check paths, agents, and plist status
cliclaw help
Help
Chat commands (in Telegram)
Command
What it does
/claude /codex /pi /gemini
Switch this chat's active agent (uninstalled agents are hidden)
/status
Per-agent session status + any job in progress
/health
Bot system status (uptime, memory, log sizes, chat/job counts)
/stop
Cancel this chat's running job (SIGTERM → SIGKILL after 5s)
/reset
Discard only the current active agent's session
/reset all
Discard every agent session in this chat
/safety
Show safety-mode status — toggle with /safety on / /safety off
/start /help
Help
Any other text / photo
Sent as a prompt to the active agent (photos are downloaded and their path is prepended to the prompt)
Switching agents keeps the old session intact — come back and continue. Sending
a new prompt to a chat with a running job is rejected (use /stop or wait).
All state is isolated under ~/.cliclaw/ :
~/.cliclaw/
├── config.json # mode 600; token, allowlist, launchd extraEnv
├── safety.json # persisted /safety on|off state
├── sessions.json # per-chat active-agent metadata
├── sessions/ # per-chat codex / pi / gemini directories
├── workspace/ # the agents' shared cwd (sandbox)
│ ├── .claude/settings.json # dangerous-command hook + deny rules when safety is ON
│ └── uploads/<chatId>/ # Telegram photo downloads
├── logs/
│ ├── bot.log # token auto-masking applied
│ ├── bot.err # launchd stderr
│ └── audit.jsonl # audit log (decisions, safety state)
└── .sock/ # confirm-gate IPC
The state directory can be moved with the CLICLAW_HOME env var:
CLICLAW_HOME= ~ /my-bot cliclaw init
Features
1. Automatic agent path discovery
No absolute paths in config.json . At startup, agents are discovered in three passes:
~/.local/bin , ~/.claude/local , /usr/local/bin , /opt/homebrew/bin
bin/<cmd> of the newest node under $NVM_DIR or ~/.nvm
command -v <cmd> in a login shell (PATH with .zshrc loaded)
Undetected agents are gracefully skipped. The bot runs fine with any subset of the four.
2. Safety mode ( /safety on · /safety off )
Dangerous Bash commands ( rm -rf , git push --force , DROP, kubectl delete , AWS delete-* , sudo , curl|sh , ssh prd-*, …) are re-confirmed via a Telegram inline keyboard [✅ Allow] [❌ Deny] — no response means auto-deny.
Claude's Read tool denies sensitive files: ~/.ssh/** , ~/.aws/** , ~/.gnupg/** , ~/.netrc , ~/.npmrc , **/.env* , **/*.pem , **/id_rsa* , **/id_ed25519* , ./secrets/** .
Add your own regexes via confirmGate.extraPatterns .
OFF : if your environment already has an external guard ( pre-bash-guard ,
EDR, …) and the bot's confirm prompts feel redundant, turn it off with one line
in Telegram. Deny rules are disabled together. Every IPC request is still
recorded in logs/audit.jsonl as decision: allow, reason: safety_off .
The state persists in $CLICLAW_HOME/safety.json across restarts.
3. Response streaming (Claude)
Consumes text_delta from --include-partial-messages and live-updates via
editMessageText , debounced at 1.5s. Past 3800 chars it rolls over to a new message.
Telegram photos/image documents are downloaded to
workspace/uploads/<chatId>/<msgId>.<ext> and the path is prepended to the prompt.
Claude : runs with --permission-mode bypassPermissions — dangerous Bash is caught by the confirm gate, and sensitive files by safety-mode deny rules.
Codex : sandbox=workspace-write by default. Never use danger-full-access .
Gemini : approvalMode=auto_edit by default (edits auto-approved, destructive commands prompt). More autonomous yolo or more conservative default / plan available.
6. Corporate TLS interceptor auto-detection
Where Zscaler / Forticlient / Cisco Umbrella intercepts HTTPS, Node cannot
trust Telegram's certificate and the bot cannot run. Step 4 of cliclaw init
auto-detects $NODE_EXTRA_CA_CERTS or launchctl getenv NODE_EXTRA_CA_CERTS ,
asks you, and persists it into launchd.extraEnv in config.json . Every later
install-launchd bakes it into the plist automatically.
Everything written to logs/bot.log / bot.err is pre-redacted:
Telegram bot tokens ( \d{8,}:[A-Za-z0-9_-]{30,} )
exact matches of the live config.token
Defends against Time Machine backups, EDR, and shoulder surfing alike.
All user-facing messages, errors, and /help are in Korean (English copy PRs welcome).
Answering "Yes" at step 5 of cliclaw init :
Creates ~/Library/LaunchAgents/com.<username>.cliclaw.plist (corporate CA baked in)
Loads and starts it immediately via launchctl bootstrap gui/$UID <plist>
Auto-restarts on login / boot / crash from then on
stdout → ~/.cliclaw/logs/bot.log , stderr → bot.err
# stop (auto-restarts on next login)
launchctl kill SIGTERM gui/ $UID /com. < username > .cliclaw
# fully disable (no auto-restart either)
cliclaw uninstall-launchd
# re-enable (auto-applies launchd.extraEnv from config.json)
cliclaw install-launchd
Security
The bot token = a remote shell into every installed agent. If it leaks, /revoke at BotFather immediately.
An empty allowedUserIds rejects all messages (fail-closed).
Keep config.json at mode 600 (init sets it automatically).
Never set confirmGate.enabled: false or switch the codex sandbox to danger-full-access .
Opting Gemini's approvalMode into yolo auto-approves every tool — use it with full understanding.
When in doubt, /safety on re-activates the deny rules instantly.
To install without the cliclaw init flow:
git clone https://github.com/choiyounggi/cliclaw.git
cd cliclaw
bun install
mkdir -p ~ /.cliclaw
cp config.example.json ~ /.cliclaw/config.json
chmod 600 ~ /.cliclaw/config.json
# fill token and allowedUserIds in config.json, then
bun run bot.ts
Tests
Dangerous patterns are regex-based — 100% classification is impossible; you own the policy.
No body-text streaming for Codex / Pi / Gemini (no structured events, or not integrated).
Gemini's dangerous commands rely solely on its own approvalMode (bash-confirm IPC not integrated).
The hook holds the IPC while waiting for the user's decision.
No voice/file attachments (photos only).
Concurrent messages in the same chat are rejected ( /stop or wait).
Per-version changes live on GitHub Releases .
Release automation (maintainer)
# 1) bump the version (auto-creates commit + tag)
npm version patch # or minor / major
# 2) push commit + tag
git push --follow-tags
Then in the GitHub web UI: "Draft a new release" → pick the tag → Publish
release. .github/workflows/publish.yml runs automatically through
npm publish --access public . The workflow first verifies the release tag
matches the package.json version and fails without publishing on a mismatch.
One-time prerequisite : repo Settings → Secrets and variables → Actions →
register NPM_TOKEN with an npm token capable of 2FA bypass.
https://www.npmjs.com/settings/younggichoi/tokens/new
Issue a Granular Access Token or a Classic Automation token (with 2FA bypass)
Add the npm_… token as the GitHub Actions secret NPM_TOKEN
Hardening option : switch to npm Trusted Publishing (OIDC) and no token is needed at all.
At https://www.npmjs.com/package/@younggichoi/cliclaw/access , add Trusted Publisher → GitH

[truncated]
