---
source: "https://github.com/Thomaszhou22/danger-guard-skill"
hn_url: "https://news.ycombinator.com/item?id=48809001"
title: "Show HN:I built a safety shield for AI agents that intercepts dangerous commands"
article_title: "GitHub - Thomaszhou22/danger-guard-skill: AI agent safety shield — intercepts dangerous commands, requires password verification, sends alerts across all platforms and AI tools · GitHub"
author: "Thomas_Zhou"
captured_at: "2026-07-06T19:42:38Z"
capture_tool: "hn-digest"
hn_id: 48809001
score: 1
comments: 0
posted_at: "2026-07-06T19:00:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN:I built a safety shield for AI agents that intercepts dangerous commands

- HN: [48809001](https://news.ycombinator.com/item?id=48809001)
- Source: [github.com](https://github.com/Thomaszhou22/danger-guard-skill)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T19:00:50Z

## Translation

タイトル: Show HN:危険なコマンドを傍受する AI エージェント用の安全シールドを構築しました
記事のタイトル: GitHub - Thomaszhou22/danger-guard-skill: AI エージェントの安全シールド - 危険なコマンドを傍受し、パスワード検証を要求し、すべてのプラットフォームと AI ツールにアラートを送信 · GitHub
説明: AI エージェントの安全シールド — 危険なコマンドを傍受し、パスワードの検証を要求し、すべてのプラットフォームと AI ツールにアラートを送信します - Thomaszhou22/danger-guard-skill

記事本文:
GitHub - Thomaszhou22/danger-guard-skill: AI エージェントの安全シールド — 危険なコマンドを傍受し、パスワード検証を要求し、すべてのプラットフォームと AI ツールにアラートを送信します · GitHub
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
トーマスチョウ22
/
危険ガードスキル
公共
通知

ns
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
Thomaszhou22/デンジャーガードスキル
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット configs configs INSTALL.md INSTALL.md ライセンス ライセンス README-CN.md README-CN.md README.md README.md SKILL.md SKILL.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントの安全シールド。危険なコマンドを実行前に傍受し、パスワードの検証を要求し、誰かが侵害されたアカウントを介してシステムを破壊しようとするとアラートを送信します。
AI アシスタントはマシンに完全にアクセスできます。誰かがあなたのメッセージング アカウント (Feishu、Telegram、Discord、WhatsApp など) を侵害した場合、AI に rm -rf / 、 format C: 、または git Push --force を指示することができます。 AIがそれをやってくれるだけです。
危険なコマンドを検出しました
↓
即時インターセプト (何も実行されない)
↓
警告+危険度表示
↓
sudo/管理者パスワードの検証が必要です
↓
3 回失敗 → ロックアウト + アラート
↓
パスワードが正しい→実行
パスワードが間違っています → ブロック + 警告
サポートされているプラットフォーム
Danger Guard は、AI エージェントに接続する任意のメッセージング プラットフォームで動作します。
Danger Guard は、すべての主要な AI コーディング ツールの構成を提供します。
階層 1 — 常にブロック (パスワードが必要)
rm -rf / 、rm -rf /* 、rm -rf ~ 、rm -rf /ユーザー
find / -exec rm {} + 、検索。 -削除
形式 C: 、 rd /s /q C:\ 、 del /s /q /f C:\*.*
dd if=/dev/zero of=/dev/sda 、dd if=/dev/urandom of=/dev/sda
mkfs.* (すべてのファイルシステム形式)
diskpart clean 、Clear-Disk -RemoveData
chmod -R 777 / 、chmod -R 000 /
フォークボム: :(){:|:&};: 、フォークしながらフォーク
シャットダウン、再起動、停止、電源オフ
kill -9 -1 (すべてのプロセスを強制終了)
カール ... | sh 、カール ... |バッシュ
wget ... | sh 、 wget ... |バッシュ
Tier 2 — 確認付きでブロック
gitプッシュ

--force 、 git Push -f 、 git Push --mirror
git replace --hard 、 git clean -fd 、 git clean -fdx
git ブランチ -D 、 git stash clear 、 git reflog 期限切れ
docker system prune -a --volumes
ドロップテーブル、ドロップデータベース、ドロップスキーマ
Danger Guard は、英語と中国語の両方で自然言語の試みも検出します。
# リポジトリのクローンを作成する
git clone https://github.com/Thomaszhou22/danger-guard.git
# OpenClaw スキル ディレクトリにコピーします
cp -r 危険ガード ~ /.openclaw/skills/
# OpenClaw ゲートウェイを再起動します (または新しいスキルを自動検出します)
Danger Guard は最初の使用時に自動的に有効になり、セットアップ手順が表示されます。
Danger Guard が初めて危険なコマンドを検出すると、セットアップが実行されます。
🛡️ デンジャーガードのセットアップ
検出されたシステム: macOS
ステップ 1: sudo パスワード (ログインパスワード) を入力します。
→ SHA256 ハッシュとして保存され、平文ではありません
ステップ 2: アラートを電子メールで送信しますか? (オプション)
1. いいえ (チャット アラートのみ)
2. ネイティブメール（macOSメール/Windows PowerShell）
3. API再送信（無料3000/月）
ステップ 3: シェルラッパー? (オプション)
1. いいえ (AI が実行したコマンドのみを保護します)
2. はい (すべての端末コマンドをシステム全体で保護します)
その他の AI ツール
適切な構成を configs/ からプロジェクトにコピーします。
# クロードコード
cp configs/claude-code/settings.json .claude/settings.json
cp configs/claude-code/CLAUDE.md ./CLAUDE.md
# OpenAI コーデックス
cp configs/codex/AGENTS.md ./AGENTS.md
# カーソル / ウィンドサーフィン / 副操縦士
cp configs/cursor/.cursorrules ./.cursorrules
# Git フック (リポジトリごと)
cp configs/git-hooks/pre-push .git/hooks/pre-push
chmod +x .git/hooks/pre-push
完全なインストール ガイド: INSTALL.md
Danger Guard は、次の 2 つの補完的な保護層を提供します。
┌─────────────────────┐
│ レイヤ 1: AI エージェント保護 │
│ (OpenClaw スキル / クロード

コード/等) │
│ │
│ 危険なコマンドを傍受する │
│ AIツールやチャットで送信 │
│ プラットフォーム │
━━━━━━━━━━━━━━━━┤
│ レイヤ 2: システムレベルの保護 │
│ (シェルラッパー — オプション) │
│ │
│ すべての端末コマンドをインターセプトします。 │
│ 誰が、あるいは何によって運営されているかに関係なく │
━━━━━━━━━━━━━━━━━━━━━━┘
レイヤー
守るもの
セットアップ
AIエージェント
チャットからのコマンド（フェイシュ、テレグラムなど）
スキルインストールで自動発動
AIツール
クロードコード、コーデックス、カーソルからのコマンド
設定ファイルをコピーする
シェルラッパー
すべての端末コマンド (システム全体)
オンボーディング中のオプション
Gitフック
危険な git 操作 (強制プッシュ)
.git/hooks/ にコピーします
緊急プロトコル
アカウントが侵害されていると思われる場合は、次のように返信してください。
最後の 50 個のコマンドをmemory/security-breach-YYYY-MM-DD.mdに記録します。
緊急アラートの送信 (チャット + メール)
セキュリティに関する推奨事項を提供する
構成は、memory/danger-guard-config.json に保存されます。
{
"os" : "macOS" ,
"sudo_hash" : " sha256:... " ,
"alert_email" : null 、
"email_method" : " なし " 、
"有効" : true 、
"作成" : " 2026-07-07 "
}
ホワイトリスト
インターセプトをスキップするコマンドとパス:
.Trash/ 、ごみ箱の操作
git clean 、 git replace (バージョン管理)
npm install (グローバルではなくプロジェクトレベル)
危険ガード/
§── SKILL.md # コアスキル(ブラックリスト+インターセプトロジック)
§── README.md # このファイル（英語）
§── README-CN.md # 中国語版
§── INSTALL.md # 詳細なインストールガイド
§── configs/
│ §── claude-code/ # クロードコードの設定
│ │ §── settings.json # 70+ 拒否パターン

│ │ └── CLAUDE.md # 安全規定
│ §── codex/ # OpenAI Codex 設定
│ │ └─ AGENTS.md # エージェントの指示
│ §── カーソル/ # カーソル/ウィンドサーフィン設定
│ │ └─ .cursorrules # 禁止されている操作
│ §──shell-wrapper/ # システムレベルの保護
│ │ ━──danger-guard # Bash インターセプト スクリプト
│ └── git-hooks/ # Git の安全性
│ └── プリプッシュ # ブロック強制プッシュ
└── メモリ/ # ランタイム設定 (自動生成)
━──danger-guard-config.json
貢献する
SKILL.mdを編集し、適切な階層の下にパターンを追加します。
ツールの構成形式で configs/ の下に新しいディレクトリを作成します。
キャッチされなかった（または誤ってキャッチされた）コマンドについて説明します。
正確なコマンド構文を含める
どの AI ツール/プラットフォームがトリガーしたかに注目してください
AI の最後の防衛線です。
AI エージェントの安全シールド - 危険なコマンドを傍受し、パスワードの検証を要求し、すべてのプラットフォームと AI ツールにアラートを送信します
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI agent safety shield — intercepts dangerous commands, requires password verification, sends alerts across all platforms and AI tools - Thomaszhou22/danger-guard-skill

GitHub - Thomaszhou22/danger-guard-skill: AI agent safety shield — intercepts dangerous commands, requires password verification, sends alerts across all platforms and AI tools · GitHub
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
Thomaszhou22
/
danger-guard-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Thomaszhou22/danger-guard-skill
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits configs configs INSTALL.md INSTALL.md LICENSE LICENSE README-CN.md README-CN.md README.md README.md SKILL.md SKILL.md View all files Repository files navigation
AI agent safety shield. Intercepts dangerous commands before execution, requires password verification, and sends alerts when someone tries to destroy your system through a compromised account.
Your AI assistant has full access to your machine. If someone compromises your messaging account — Feishu, Telegram, Discord, WhatsApp, whatever — they can tell your AI to rm -rf / , format C: , or git push --force . The AI just does it.
Dangerous command detected
↓
Immediate intercept (nothing executes)
↓
Warning + risk level displayed
↓
Requires sudo/admin password verification
↓
3 failed attempts → lockout + alert
↓
Password correct → execute
Password wrong → block + alert
Supported Platforms
Danger Guard works with any messaging platform that connects to your AI agent:
Danger Guard provides configs for all major AI coding tools:
Tier 1 — Always Blocked (Requires Password)
rm -rf / , rm -rf /* , rm -rf ~ , rm -rf /Users
find / -exec rm {} + , find . -delete
format C: , rd /s /q C:\ , del /s /q /f C:\*.*
dd if=/dev/zero of=/dev/sda , dd if=/dev/urandom of=/dev/sda
mkfs.* (all filesystem formats)
diskpart clean , Clear-Disk -RemoveData
chmod -R 777 / , chmod -R 000 /
Fork bombs: :(){:|:&};: , fork while fork
shutdown , reboot , halt , poweroff
kill -9 -1 (kill all processes)
curl ... | sh , curl ... | bash
wget ... | sh , wget ... | bash
Tier 2 — Blocked With Confirmation
git push --force , git push -f , git push --mirror
git reset --hard , git clean -fd , git clean -fdx
git branch -D , git stash clear , git reflog expire
docker system prune -a --volumes
DROP TABLE , DROP DATABASE , DROP SCHEMA
Danger Guard also catches natural language attempts in both English and Chinese:
# Clone the repo
git clone https://github.com/Thomaszhou22/danger-guard.git
# Copy to your OpenClaw skills directory
cp -r danger-guard ~ /.openclaw/skills/
# Restart OpenClaw gateway (or it auto-detects new skills)
Danger Guard activates automatically on first use and walks you through setup.
The first time Danger Guard detects a dangerous command, it runs through setup:
🛡️ Danger Guard Setup
Detected system: macOS
Step 1: Enter your sudo password (login password)
→ Stored as SHA256 hash, never plaintext
Step 2: Email alerts? (optional)
1. No (chat alerts only)
2. Native mail (macOS mail / Windows PowerShell)
3. Resend API (free 3000/month)
Step 3: Shell Wrapper? (optional)
1. No (protect AI-executed commands only)
2. Yes (protect ALL terminal commands system-wide)
Other AI Tools
Copy the appropriate config from configs/ to your project:
# Claude Code
cp configs/claude-code/settings.json .claude/settings.json
cp configs/claude-code/CLAUDE.md ./CLAUDE.md
# OpenAI Codex
cp configs/codex/AGENTS.md ./AGENTS.md
# Cursor / Windsurf / Copilot
cp configs/cursor/.cursorrules ./.cursorrules
# Git Hooks (per-repo)
cp configs/git-hooks/pre-push .git/hooks/pre-push
chmod +x .git/hooks/pre-push
Full installation guide: INSTALL.md
Danger Guard provides two complementary layers of protection:
┌─────────────────────────────────────────┐
│ Layer 1: AI Agent Protection │
│ (OpenClaw Skill / Claude Code / etc.) │
│ │
│ Intercepts dangerous commands when │
│ sent through AI tools and chat │
│ platforms │
├─────────────────────────────────────────┤
│ Layer 2: System-Level Protection │
│ (Shell Wrapper — optional) │
│ │
│ Intercepts ALL terminal commands, │
│ regardless of who or what runs them │
└─────────────────────────────────────────┘
Layer
What It Protects
Setup
AI Agent
Commands from chat (Feishu, Telegram, etc.)
Auto-activates with skill install
AI Tools
Commands from Claude Code, Codex, Cursor
Copy config files
Shell Wrapper
ALL terminal commands (system-wide)
Optional during onboarding
Git Hooks
Dangerous git operations (force push)
Copy to .git/hooks/
Emergency Protocol
If you suspect your account is compromised, reply with:
Log the last 50 commands to memory/security-breach-YYYY-MM-DD.md
Send emergency alerts (chat + email)
Provide security recommendations
Config is stored in memory/danger-guard-config.json :
{
"os" : " macOS " ,
"sudo_hash" : " sha256:... " ,
"alert_email" : null ,
"email_method" : " none " ,
"enabled" : true ,
"created" : " 2026-07-07 "
}
Whitelist
Commands and paths that skip interception:
.Trash/ , Recycle Bin operations
git clean , git reset (version-controlled)
npm install (project-level, not global)
danger-guard/
├── SKILL.md # Core skill (blacklist + intercept logic)
├── README.md # This file (English)
├── README-CN.md # Chinese version
├── INSTALL.md # Detailed installation guide
├── configs/
│ ├── claude-code/ # Claude Code config
│ │ ├── settings.json # 70+ deny patterns
│ │ └── CLAUDE.md # Safety rules
│ ├── codex/ # OpenAI Codex config
│ │ └── AGENTS.md # Agent instructions
│ ├── cursor/ # Cursor/Windsurf config
│ │ └── .cursorrules # Forbidden operations
│ ├── shell-wrapper/ # System-level protection
│ │ └── danger-guard # Bash intercept script
│ └── git-hooks/ # Git safety
│ └── pre-push # Block force push
└── memory/ # Runtime config (auto-generated)
└── danger-guard-config.json
Contributing
Edit SKILL.md and add patterns under the appropriate tier.
Create a new directory under configs/ with the tool's config format.
Describe the command that wasn't caught (or was falsely caught)
Include the exact command syntax
Note which AI tool / platform triggered it
Your AI's last line of defense.
AI agent safety shield — intercepts dangerous commands, requires password verification, sends alerts across all platforms and AI tools
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
