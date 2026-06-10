---
source: "https://github.com/martidu4/honey-ai"
hn_url: "https://news.ycombinator.com/item?id=48477620"
title: "HoneyAI – AI-powered honeypot replacing Cowrie/OpenCanary/Galah with one process"
article_title: "GitHub - martidu4/honey-ai: 🍯 All-in-one AI honeypot powered by local LLMs. SSH, HTTP, FTP, Telnet, SMTP, MySQL, Redis, Git, VNC, RDP — with canary tokens, tarpits, GZIP bombs, and threat intel reporting. · GitHub"
author: "whatda77"
captured_at: "2026-06-10T16:20:51Z"
capture_tool: "hn-digest"
hn_id: 48477620
score: 1
comments: 0
posted_at: "2026-06-10T15:15:34Z"
tags:
  - hacker-news
  - translated
---

# HoneyAI – AI-powered honeypot replacing Cowrie/OpenCanary/Galah with one process

- HN: [48477620](https://news.ycombinator.com/item?id=48477620)
- Source: [github.com](https://github.com/martidu4/honey-ai)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T15:15:34Z

## Translation

タイトル: HoneyAI – Cowrie/OpenCanary/Galah を 1 つのプロセスで置き換える AI を活用したハニーポット
記事のタイトル: GitHub - martidu4/honey-ai: 🍯 ローカル LLM を利用したオールインワン AI ハニーポット。 SSH、HTTP、FTP、Telnet、SMTP、MySQL、Redis、Git、VNC、RDP - カナリア トークン、タールピット、GZIP ボム、脅威インテリジェンス レポートを使用します。 · GitHub
説明: 🍯 ローカル LLM を利用したオールインワン AI ハニーポット。 SSH、HTTP、FTP、Telnet、SMTP、MySQL、Redis、Git、VNC、RDP - カナリア トークン、タールピット、GZIP ボム、脅威インテリジェンス レポートを使用します。 - martidu4/ハニーアイ

記事本文:
GitHub - martidu4/honey-ai: 🍯 ローカル LLM を利用したオールインワン AI ハニーポット。 SSH、HTTP、FTP、Telnet、SMTP、MySQL、Redis、Git、VNC、RDP - カナリア トークン、タールピット、GZIP ボム、脅威インテリジェンス レポートを使用します。 · GitHub
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
マルティドゥ4
/
ハニーアイ
P

ユーブリック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github/ workflows .github/ workflows ai ai core core docs docs Honeyfs Honeyfs プロトコル プロトコル .dockerignore .dockerignore .gitignore .gitignore .host_key .host_key COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンスREADME.md README.md config.example.yaml config.example.yaml config.pi5.yaml config.pi5.yaml docker-compose.yml docker-compose.yml Honey-ai.service Honey-ai.service Honeyai-stats.py Honeyai-stats.py package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml server.js server.js setup.js setup.js test-qa.js test-qa.js test-stress.js test-stress.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オールインワンの AI 搭載ハニーポット。 1 つのプロセス、すべてのプロトコル。
Cowrie、モモイロインコ、OpenCanary、Endlessh を、ローカル LLM によって駆動される単一の Node.js サービスに置き換えます。
HoneyAI は、プロアクティブな AI 駆動のハニーポットで、あらゆる一般的なプロトコルにわたって攻撃者を傍受し、Ollama を介して実行されるローカル LLM を利用して、動的に生成された完全に説得力のある欺瞞的なコンテンツで応答します。
AI は静的な応答の代わりに、攻撃者のペイロードを読み取り、調整されたトラップを生成します。
💉 SQL インジェクションの試行 → カナリア トークン (制御するベイト API キー) を使用した偽のデータベース ダンプ
🐚 シェルのアップロード → より多くの餌を使用した偽の実行出力
🔑 SSH ログイン → 現実的なファイルシステムを備えたインタラクティブな偽の bash シェル
📂 ディレクトリスキャン → 偽のbackup.zip 、 .env 、 config.php 、 sql_dump.tar.gz
🎣 Cat の機密ファイル → 偽の AWS 認証情報、SSH キー、データベース パスワード
すべての攻撃者の IP が 5 つの脅威に自動的に報告されます。

インテリジェンスプラットフォーム。
特徴
説明
🌐 HTTP/HTTPS
キャッチオール Web ハニーポット。 WordPress、Apache、phpMyAdmin、Laravel を模倣します。モモイロインコを置き換える
🔑SSH
カナリアファイルシステムを備えたインタラクティブな偽の bash シェル。すべての資格情報を受け入れます。カウリーを置き換える
🧲 SSH ターピット
設定可能なポート上の無限のバナー。 Endlessh を置き換える
📁 FTP
AI が生成したディレクトリ リストを使用した偽の vsftPd
📟 Telnet
偽のルーター/スイッチ CLI (静的 show コマンドを使用した Cisco IOS スタイル)
📧 SMTP
偽メールサーバー - すべてのメッセージを受け入れてログに記録します
🗄️MySQL
偽の MySQL 8.0 — ハンドシェイク + 認証 + クエリ応答
🔴レディス
Fake Redis — 静的 RESP プロトコル (PING、INFO、KEYS、CONFIG)
🐙 Git
ポート 9418 上の Git プロトコル ハニーポット
🖥️VNC
RFB プロトコルのハンドシェイク トラップ
💻RDP
RDP プロトコルのハンドシェイク トラップ
💣 GZIP 爆弾
圧縮されたペイロード爆弾をスキャナーに届けます
📡 レポート
AbuseIPDB、OTX、DShield、Blocklist.de、VirusTotal への自動レポート
📲 電報
Telegram ボットによるリアルタイム攻撃通知
🤖 任意の LLM
Ollama (ローカル) または任意の OpenAI 互換 API で動作します
クイックスタート (ベアメタル)
🐳 ドッカー? 1 つのコマンドでセットアップするには、「Docker デプロイメント」に進んでください。
pnpm — npm install -g pnpm でインストールします
ローカルで実行される Ollama (または任意の OpenAI 互換 API)
プルしたモデル: ollama pull qwen2.5:1.5b (高速、1GB RAM)
⚠️なぜ pnpm だけなのですか?このプロジェクトは、プリインストール フックを介して npm と Yarn をブロックします。 npm は、インストール中にあらゆる依存関係から任意のライフサイクル スクリプト ( preinstall 、 postinstall ) を実行します。これは既知のサプライ チェーン攻撃ベクトルです (参考資料)。ハニーポットのようなセキュリティ ツールの場合、これは受け入れられません。 pnpm はデフォルトではこれらのスクリプトを実行せず、ファントム依存関係を防止し、厳密な分離を提供するコンテンツアドレス指定可能なストアを使用します。 npm install を試行すると意図的に失敗します

。
# pnpm がない場合はインストールします
npm install -g pnpm
# クローンを作成して実行する
git clone https://github.com/martidu4/honey-ai.git
CD ハニーアイ
pnpm install # npm/yarn は拒否されます — pnpm のみ
pnpm run setup # 対話型ウィザード — AI、レポート、カナリア トークンを構成します
pnpm start # 🍯 すべてのプロトコルがリッスンを開始します
セットアップ ウィザードでは次のことを尋ねられます。
Ollama の URL とモデル (または OpenAI 互換 API)
AbuseIPDB、OTX、DShield、Blocklist.de、VirusTotal API キー (すべてオプション)
攻撃通知用の Telegram ボット (オプション)
設定は config.yaml に保存されますが、これは gitignore され、コミットされることはありません。
最も速く始める方法 — 1 つのコマンドで、すべてが含まれます。
git clone https://github.com/martidu4/honey-ai.git
CD ハニーアイ
cp config.example.yaml config.yaml
# すべてを開始します (Ollama + モデルのダウンロード + HoneyAI)
ドッカー構成 -d
# ログを追跡する
docker compose ログ -f Honeyai
Docker Compose は自動的に行われます。
永続モデルストレージを使用して Ollama を開始します
最初の実行時に qwen2.5:1.5b モデルをプルします
11 プロトコルすべてで HoneyAI を開始します
AI_MODEL=qwen3:4b docker 構成 -d
レポート API キーを追加するには、.env ファイルを作成します。
ABUSEIPDB_KEY = your_key
TELEGRAM_TOKEN = your_bot_token
TELEGRAM_CHAT = your_chat_id
建築
インターネット攻撃者
│
§─ :80/8080 → HTTP ハニーポット (Express + AI 応答)
§─ :22/2222 → SSH ハニーポット (ssh2 + AI インタラクティブ シェル)
§─ :222/2200 → SSHターピット(エンドレス式無限バナー)
§─ :21 → FTP ハニーポット (TCP + AI)
§─ :23 → Telnet (TCP + AI、Cisco IOS スタイル)
§─ :25 → SMTP (TCP + AI)
§─ :3306 → MySQL (TCP + プロトコル精度のハンドシェイク)
§─ :6379 → Redis (TCP + 静的 RESP プロトコル)
§─ :9418 → Git (TCP + 偽のリポジトリ応答)
§─ :5900 → VNC (TCP + RFB ハンドシェイク)
━─ :3389 → RDP (TCP + RDP

握手）
│
▼
AIエンジン（Ollama / OpenAI互換）
│
§─ 欺瞞的な対応 → 攻撃者
§─ レポーター → AbuseIPDB、OTX、DShield、Blocklist.de、VT
└─ 電報 → リアルタイムアラート 📲
プロジェクトの構造
ハニーアイ/
§──server.js # メインオーケストレーター — すべてのプロトコルを開始します
§── setup.js # 対話型セットアップウィザード
§── config.example.yaml # 構成テンプレート (コミット済み - シークレットなし)
§── Honey-ai.service # 本番用の systemd サービスファイル
§── あい/
│ └─ Engine.js # AI エンジン — Ollama/OpenAI + ID 漏洩フィルター
§── コア/
│ §── config.js # 構成ローダー (YAML + 環境変数)
│ §── logger.js # 統合ロガー (コンソール + JSONL、CRLF セーフ)
│ §──reporter.js # 脅威情報レポート (5 つのプラットフォーム)
│ §──traps.js # Web 迷路、GZIP 爆弾、カナリアのダウンロード
│ §── backfire.js # 攻撃者 IP の逆スキャン
│ §── downloader.js # マルウェア サンプル コレクター (SSRF 保護)
│ §── fileReader.js # HoneyFS 仮想ファイルシステムリーダー
│ └─ jitter.js # 現実的な遅延を実現するタイミングランダマイザー
§── プロトコル/
│ §── http.js # HTTP ハニーポット (モモイロインコの置き換え)
│ §── ssh.js # SSH ハニーポット + タールピット (Cowrie + Endlessh を置き換えます)
│ └── tcp.js # FTP、Telnet、SMTP、MySQL、Redis、Git、VNC、RDP
§── Honeyfs/ # 🎣 Canary ファイルシステム — 攻撃者はこれらのファイルを参照します
│ §── etc/ # 偽の /etc/passwd、シャドウ、グループ、ホスト名
│ §── home/ # 偽の暗号通貨ウォレット、認証情報ファイル
│ §── opt/ # 偽の docker-compose、.env、terraform、k8s シークレット
│ └── root/ # 偽の .aws/credentials、.ssh/id_rsa、passwords.txt
└── test-qa.js # 完全なテスト スイート (98 個のテスト)
🎣 カナリアトークン (ハニーポットファイルシステム)
Honeyfs/ ディレクトリには偽のファイルが含まれています

攻撃者が SSH または HTTP 経由で閲覧するときに見つけられる重要なファイル。これらはカナリア トークン、つまり攻撃者によって使用されると侵害を警告するおとりの資格情報です。
⚠️ 重要: デプロイする前に、すべての CHANGE_ME_* 値を独自のおとり認証情報に置き換えてください。
# 例: https://canarytokens.org/ で独自のカナリア AWS キーを生成します。
# 次に、次のように置き換えます。
Honeyfs/root/.aws/credentials # 偽の AWS キー
Honeyfs/root/.env # 偽の DB/ストライプ/AWS 認証情報
Honeyfs/root/config.json # 偽の完全なアプリケーション構成
Honeyfs/root/passwords.txt # 偽のマスターパスワードリスト
Honeyfs/root/.ssh/id_rsa # 偽の SSH 秘密鍵
Honeyfs/root/.github-token # 偽の GitHub PAT
Honeyfs/opt/app/.env # 偽のアプリ環境
Honeyfs/opt/app/docker-compose.yml # 偽の Docker スタック
Honeyfs/opt/k8s/secrets.yaml # 偽の Kubernetes シークレット
Honeyfs/opt/infra/terraform.tfstate # 偽の Terraform 状態
アイデアは、攻撃者がこれらの資格情報を盗んで使用しようとしたときに、カナリア トークン サービスを介して侵害を検出するというものです。 canarytokens.org または独自の検出メカニズムを使用します。
オプション A: セットアップ ウィザード (推奨)
pnpm 実行セットアップ
オプション B: 手動構成
cp config.example.yaml config.yaml
# config.yaml を編集 — 有効にするポート、AI モデル、プロトコル
使用可能なすべてのオプションとコメントについては、config.example.yaml を参照してください。
環境変数を使用して構成値をオーバーライドできます。
OLLAMA_URL = http://localhost:11434
AI_MODEL = qwen2.5:1.5b
# レポート (すべてオプション - 無料枠にサインアップ)
ABUSEIPDB_KEY = your_key_here
OTX_KEY = your_key_here
DSHIELD_KEY = your_key_here
BLOCKLIST_KEY = your_key_here
VT_KEY = your_key_here
# 通知
TELEGRAM_TOKEN = your_bot_token
TELEGRAM_CHAT = your_chat_id
システムサービスとして導入する
# 1. 専用ユーザーを作成します (決して root として実行しないでください)。
sudo useradd -r -s /usr

/sbin/nologin ハニーアイ
# 2. /opt にクローンを作成する
sudo git clone https://github.com/martidu4/honey-ai.git /opt/honey-ai
cd /opt/honey-ai && sudo -u Honeyai pnpm install
#3. 設定する
sudo -u Honeyai pnpm 実行セットアップ
# 4. サービスのインストールと開始
sudo cp Honey-ai.service /etc/systemd/system/
sudo systemctl デーモン-リロード
sudo systemctl Enable --now Honey-ai
#5. ログをたどる
sudo ジャーナルctl -u ハニーアイ -f
ポート転送 (root なしで実行)
HoneyAI はデフォルトで高ポートで実行されます。 iptables を使用して標準ポートをリダイレクトします。
# リダイレクト :22 → :2226 (SSH ハニーポット)
sudo iptables -t nat -A PREROUTING -p tcp --dport 22 -j REDIRECT --to-port 2226
# リダイレクト :21 → :2121 (FTP)、:23 → :2323 (Telnet)、:25 → :2525 (SMTP)
sudo iptables -t nat -A PREROUTING -p tcp --dport 21 -j REDIRECT --to-port 2121
sudo iptables -t nat -A PREROUTING -p tcp --dport 23 -j REDIRECT --to-port 2323
sudo iptables -t nat -A PREROUTING -p tcp --dport 25 -j REDIRECT --to-port 2525
# リダイレクト :3306 → :33060 (MySQL)、 :6379 → :63790 (Redis)
sudo iptables -t nat -A PREROUTING -p tcp --dport 3306 -j REDIRECT --to-port 33060
sudo iptables -t nat -A PREROUTING -p tcp --dport 6379 -j REDIRECT --to-port 63790
推奨される LLM モデル
モデル
サイズ
速度
品質
こんな方に最適
qwen2.5:0.5b
400MB
⚡⚡⚡
良い
低リソースデバイス (Pi、VPS)
qwen2.5:1.5b
1GB
⚡⚡
より良い
推奨 — ベストバランス
クウェン3:4b
2.5GB
⚡
ベスト
高品質の欺瞞
任意の OpenAI 互換
雲
⚡
素晴らしい
クラウド展開
ヒント: Raspberry Pi 5 では、qwen2.5:1.5b により優れた結果が得られます。 Ollama を別のマシンで実行し、HoneyAI をそのマシンにポイントすることもできます。
# 完全なテスト スイートを実行します (98 個のテスト — Ollama の実行が必要です)
ノード

[切り捨てられた]

## Original Extract

🍯 All-in-one AI honeypot powered by local LLMs. SSH, HTTP, FTP, Telnet, SMTP, MySQL, Redis, Git, VNC, RDP — with canary tokens, tarpits, GZIP bombs, and threat intel reporting. - martidu4/honey-ai

GitHub - martidu4/honey-ai: 🍯 All-in-one AI honeypot powered by local LLMs. SSH, HTTP, FTP, Telnet, SMTP, MySQL, Redis, Git, VNC, RDP — with canary tokens, tarpits, GZIP bombs, and threat intel reporting. · GitHub
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
martidu4
/
honey-ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github/ workflows .github/ workflows ai ai core core docs docs honeyfs honeyfs protocols protocols .dockerignore .dockerignore .gitignore .gitignore .host_key .host_key CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md config.example.yaml config.example.yaml config.pi5.yaml config.pi5.yaml docker-compose.yml docker-compose.yml honey-ai.service honey-ai.service honeyai-stats.py honeyai-stats.py package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml server.js server.js setup.js setup.js test-qa.js test-qa.js test-stress.js test-stress.js View all files Repository files navigation
All-in-one AI-powered honeypot. One process, every protocol.
Replaces Cowrie · Galah · OpenCanary · Endlessh — with a single Node.js service driven by a local LLM.
HoneyAI is a proactive, AI-driven honeypot that intercepts attackers across every common protocol and responds with dynamically generated, fully convincing deceptive content — powered by a local LLM running via Ollama .
Instead of static responses, the AI reads the attacker's payload and generates tailored traps:
💉 SQL injection attempt → Fake database dump with canary tokens (bait API keys you control)
🐚 Shell upload → Fake execution output with more bait
🔑 SSH login → Interactive fake bash shell with realistic filesystem
📂 Directory scan → Fake backup.zip , .env , config.php , sql_dump.tar.gz
🎣 Cat sensitive files → Fake AWS credentials, SSH keys, database passwords
Every attacker IP is automatically reported to 5 threat intelligence platforms .
Feature
Description
🌐 HTTP/HTTPS
Catch-all web honeypot. Mimics WordPress, Apache, phpMyAdmin, Laravel. Replaces Galah
🔑 SSH
Interactive fake bash shell with canary filesystem. Accepts all credentials. Replaces Cowrie
🧲 SSH Tarpit
Infinite banner on configurable ports. Replaces Endlessh
📁 FTP
Fake vsftPd with AI-generated directory listings
📟 Telnet
Fake router/switch CLI (Cisco IOS style with static show commands)
📧 SMTP
Fake mail server — accepts and logs all messages
🗄️ MySQL
Fake MySQL 8.0 — handshake + auth + query responses
🔴 Redis
Fake Redis — static RESP protocol (PING, INFO, KEYS, CONFIG)
🐙 Git
Git protocol honeypot on port 9418
🖥️ VNC
RFB protocol handshake trap
💻 RDP
RDP protocol handshake trap
💣 GZIP Bombs
Delivers compressed payload bombs to scanners
📡 Reporting
Auto-reports to AbuseIPDB, OTX, DShield, Blocklist.de, VirusTotal
📲 Telegram
Real-time attack notifications via Telegram bot
🤖 Any LLM
Works with Ollama (local) or any OpenAI-compatible API
Quick Start (bare metal)
🐳 Docker? Skip to Docker Deployment for a one-command setup.
pnpm — install with npm install -g pnpm
Ollama running locally (or any OpenAI-compatible API)
A model pulled: ollama pull qwen2.5:1.5b (fast, 1GB RAM)
⚠️ Why pnpm only? This project blocks npm and yarn via a preinstall hook. npm executes arbitrary lifecycle scripts ( preinstall , postinstall ) from every dependency during install — this is a known supply chain attack vector ( reference ). For a security tool like a honeypot, this is unacceptable. pnpm does not run these scripts by default, uses a content-addressable store that prevents phantom dependencies, and provides strict isolation. If you try npm install , it will fail intentionally.
# Install pnpm if you don't have it
npm install -g pnpm
# Clone and run
git clone https://github.com/martidu4/honey-ai.git
cd honey-ai
pnpm install # npm/yarn will be rejected — pnpm only
pnpm run setup # Interactive wizard — configures AI, reporting, canary tokens
pnpm start # 🍯 All protocols start listening
The setup wizard will ask you for:
Your Ollama URL and model (or OpenAI-compatible API)
AbuseIPDB, OTX, DShield, Blocklist.de, VirusTotal API keys (all optional)
Telegram bot for attack notifications (optional)
Configuration is saved to config.yaml which is gitignored and never committed.
The fastest way to get started — one command, everything included:
git clone https://github.com/martidu4/honey-ai.git
cd honey-ai
cp config.example.yaml config.yaml
# Start everything (Ollama + model download + HoneyAI)
docker compose up -d
# Follow logs
docker compose logs -f honeyai
Docker Compose automatically:
Starts Ollama with persistent model storage
Pulls the qwen2.5:1.5b model on first run
Starts HoneyAI with all 11 protocols
AI_MODEL=qwen3:4b docker compose up -d
To add reporting API keys, create a .env file:
ABUSEIPDB_KEY = your_key
TELEGRAM_TOKEN = your_bot_token
TELEGRAM_CHAT = your_chat_id
Architecture
Internet attackers
│
├─ :80/8080 → HTTP honeypot (Express + AI responses)
├─ :22/2222 → SSH honeypot (ssh2 + AI interactive shell)
├─ :222/2200 → SSH tarpit (Endlessh-style infinite banner)
├─ :21 → FTP honeypot (TCP + AI)
├─ :23 → Telnet (TCP + AI, Cisco IOS style)
├─ :25 → SMTP (TCP + AI)
├─ :3306 → MySQL (TCP + protocol-accurate handshake)
├─ :6379 → Redis (TCP + static RESP protocol)
├─ :9418 → Git (TCP + fake repo responses)
├─ :5900 → VNC (TCP + RFB handshake)
└─ :3389 → RDP (TCP + RDP handshake)
│
▼
AI Engine (Ollama / OpenAI-compatible)
│
├─ Deceptive response → attacker
├─ Reporter → AbuseIPDB, OTX, DShield, Blocklist.de, VT
└─ Telegram → real-time alert 📲
Project Structure
honey-ai/
├── server.js # Main orchestrator — starts all protocols
├── setup.js # Interactive setup wizard
├── config.example.yaml # Config template (committed — no secrets)
├── honey-ai.service # systemd service file for production
├── ai/
│ └── engine.js # AI engine — Ollama/OpenAI + identity leak filters
├── core/
│ ├── config.js # Config loader (YAML + env vars)
│ ├── logger.js # Unified logger (console + JSONL, CRLF-safe)
│ ├── reporter.js # Threat intel reporting (5 platforms)
│ ├── traps.js # Web maze, GZIP bombs, canary downloads
│ ├── backfire.js # Reverse scanning of attacker IPs
│ ├── downloader.js # Malware sample collector (SSRF-protected)
│ ├── fileReader.js # HoneyFS virtual filesystem reader
│ └── jitter.js # Timing randomizer for realistic delays
├── protocols/
│ ├── http.js # HTTP honeypot (replaces Galah)
│ ├── ssh.js # SSH honeypot + tarpit (replaces Cowrie + Endlessh)
│ └── tcp.js # FTP, Telnet, SMTP, MySQL, Redis, Git, VNC, RDP
├── honeyfs/ # 🎣 Canary filesystem — attackers see these files
│ ├── etc/ # Fake /etc/passwd, shadow, group, hostname
│ ├── home/ # Fake crypto wallets, credential files
│ ├── opt/ # Fake docker-compose, .env, terraform, k8s secrets
│ └── root/ # Fake .aws/credentials, .ssh/id_rsa, passwords.txt
└── test-qa.js # Full test suite (98 tests)
🎣 Canary Tokens (Honeypot Filesystem)
The honeyfs/ directory contains fake sensitive files that attackers will find when browsing via SSH or HTTP. These are your canary tokens — bait credentials that, when used by an attacker, alert you to a compromise.
⚠️ IMPORTANT: Replace ALL CHANGE_ME_* values with your own bait credentials before deploying.
# Example: Generate your own canary AWS keys at https://canarytokens.org/
# Then replace in:
honeyfs/root/.aws/credentials # Fake AWS keys
honeyfs/root/.env # Fake DB/Stripe/AWS credentials
honeyfs/root/config.json # Fake full application config
honeyfs/root/passwords.txt # Fake master password list
honeyfs/root/.ssh/id_rsa # Fake SSH private key
honeyfs/root/.github-token # Fake GitHub PAT
honeyfs/opt/app/.env # Fake app environment
honeyfs/opt/app/docker-compose.yml # Fake Docker stack
honeyfs/opt/k8s/secrets.yaml # Fake Kubernetes secrets
honeyfs/opt/infra/terraform.tfstate # Fake Terraform state
The idea: when an attacker steals these credentials and tries to use them, you'll detect the breach via the canary token service. Use canarytokens.org or your own detection mechanism.
Option A: Setup Wizard (recommended)
pnpm run setup
Option B: Manual Configuration
cp config.example.yaml config.yaml
# Edit config.yaml — ports, AI model, protocols to enable
See config.example.yaml for all available options with comments.
You can override config values with environment variables:
OLLAMA_URL = http://localhost:11434
AI_MODEL = qwen2.5:1.5b
# Reporting (all optional — sign up for free tiers)
ABUSEIPDB_KEY = your_key_here
OTX_KEY = your_key_here
DSHIELD_KEY = your_key_here
BLOCKLIST_KEY = your_key_here
VT_KEY = your_key_here
# Notifications
TELEGRAM_TOKEN = your_bot_token
TELEGRAM_CHAT = your_chat_id
Deploying as a System Service
# 1. Create a dedicated user (never run as root!)
sudo useradd -r -s /usr/sbin/nologin honeyai
# 2. Clone to /opt
sudo git clone https://github.com/martidu4/honey-ai.git /opt/honey-ai
cd /opt/honey-ai && sudo -u honeyai pnpm install
# 3. Configure
sudo -u honeyai pnpm run setup
# 4. Install and start service
sudo cp honey-ai.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now honey-ai
# 5. Follow logs
sudo journalctl -u honey-ai -f
Port Forwarding (run without root)
HoneyAI runs on high ports by default. Use iptables to redirect standard ports:
# Redirect :22 → :2226 (SSH honeypot)
sudo iptables -t nat -A PREROUTING -p tcp --dport 22 -j REDIRECT --to-port 2226
# Redirect :21 → :2121 (FTP), :23 → :2323 (Telnet), :25 → :2525 (SMTP)
sudo iptables -t nat -A PREROUTING -p tcp --dport 21 -j REDIRECT --to-port 2121
sudo iptables -t nat -A PREROUTING -p tcp --dport 23 -j REDIRECT --to-port 2323
sudo iptables -t nat -A PREROUTING -p tcp --dport 25 -j REDIRECT --to-port 2525
# Redirect :3306 → :33060 (MySQL), :6379 → :63790 (Redis)
sudo iptables -t nat -A PREROUTING -p tcp --dport 3306 -j REDIRECT --to-port 33060
sudo iptables -t nat -A PREROUTING -p tcp --dport 6379 -j REDIRECT --to-port 63790
Recommended LLM Models
Model
Size
Speed
Quality
Best for
qwen2.5:0.5b
400MB
⚡⚡⚡
Good
Low-resource devices (Pi, VPS)
qwen2.5:1.5b
1GB
⚡⚡
Better
Recommended — best balance
qwen3:4b
2.5GB
⚡
Best
High-quality deception
Any OpenAI-compat
cloud
⚡
Excellent
Cloud deployments
Tip: On a Raspberry Pi 5, qwen2.5:1.5b gives great results. You can also run Ollama on a separate machine and point HoneyAI to it.
# Run full test suite (98 tests — requires Ollama running)
node

[truncated]
