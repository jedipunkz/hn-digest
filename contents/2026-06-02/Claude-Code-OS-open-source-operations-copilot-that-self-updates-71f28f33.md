---
source: "https://github.com/bernardohcrocha/claude-code-os"
hn_url: "https://news.ycombinator.com/item?id=48358513"
title: "Claude Code OS: open-source operations copilot that self-updates"
article_title: "GitHub - bernardohcrocha/claude-code-os: The operational co-pilot that self-updates. Self-improving, proactive, local-first. · GitHub"
author: "bernardohcr"
captured_at: "2026-06-02T00:37:38Z"
capture_tool: "hn-digest"
hn_id: 48358513
score: 3
comments: 2
posted_at: "2026-06-01T15:51:26Z"
tags:
  - hacker-news
  - translated
---

# Claude Code OS: open-source operations copilot that self-updates

- HN: [48358513](https://news.ycombinator.com/item?id=48358513)
- Source: [github.com](https://github.com/bernardohcrocha/claude-code-os)
- Score: 3
- Comments: 2
- Posted: 2026-06-01T15:51:26Z

## Translation

タイトル: Claude Code OS: 自己更新するオープンソースの運用コパイロット
記事のタイトル: GitHub - bernardohcrocha/claude-code-os: 自己更新する運用コパイロット。自己改善、積極的、地元第一主義。 · GitHub
説明: 自己更新する運用副操縦士。自己改善、積極的、地元第一主義。 - bernardohcrocha/claude-code-os

記事本文:
GitHub - bernardohcrocha/claude-code-os: 自己更新する運用副操縦士。自己改善、積極的、地元第一主義。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ベルナルドクロチャ
/
クロードコードOS
公共
通知
あなたは署名している必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
bernardohcrocha/claude-code-os
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
51 コミット 51 コミット スクリプト スクリプト テンプレート テンプレート ARCHITECTURE.md ARCHITECTURE.md FRAMEWORK.html FRAMEWORK.html LICENSE LICENSE README.md README.md claude-code-os.png claude-code-os.png setup.sh setup.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
クロード コードのエージェント メモリを自己更新します。
Claude Code を、プロジェクトの変更を記憶し、シームレスに動作するエージェントに変えます。
クロード・コードには動作記憶がありません。すべてのセッションはゼロから始まります。スタックを再説明し、コンテキストを貼り付け、自分自身を繰り返します。
Claude Code OS はそれを修正します。 _brain/ フォルダーはプロジェクトのルートに存在し、永続メモリとして機能します。プロジェクトを一度読み取って、git diff から毎日自己更新します。変更されたもののみで、それ以上は更新されません。
プロジェクトのルート フォルダーで Claude Code を開き、次を実行します。
カール -fsSL https://raw.githubusercontent.com/bernardohcrocha/claude-code-os/main/setup.sh |バッシュ
クロードはあなたのプロジェクトをスキャンし、見つからないものだけを質問します。フォームはありません。設定ファイルはありません。ただの会話です。
オプション: GitHub CLI ( gh ) — 更新ごとにクラウド バックアップを使用してプライベート Brain リポジトリを自動的に作成します。マシンをフォーマットし、Brain リポジトリを git clone して、中断したところから正確に続行します。
スケジューラーが含まれています: 自動的にインストールされます (macOS では launchd、Linux では systemd)。
メモリの自動更新 — git diff HEAD~1 を毎日実行します。プロジェクト内に 1,000 個のファイルがあり、今日は 3 個が変更されました。読み取り値は 3 です
永続的なスキル — 1 回言う → _brain/skills/ に書き込まれ、今後のセッションごとにロードされる
自律タスク — 平易な言語であらゆるタスクをスケジュールし、完全なコンテキストで自動的に実行します
ライブ ダッシュボード — _brain/dashboard.html を開き、

5分ごとに自動更新
Hermes Agent、Agent Zero、OpenClaw とは異なり、コンテキストは自動的に更新されます。手動によるメンテナンスは必要ありません。
→ 「今月使用量が 30% 以上減少したのはどの顧客ですか? サポート履歴を照合し、それぞれにパーソナライズされた再エンゲージメント メッセージを作成します。」
→ 「毎週月曜日: Stripe から先週の数値を取得し、異常を報告し、割り当てを下回ったアカウントのフォローアップをキューに入れます。」
→ 「過去 30 日間のアクティビティのないサインアップを検索します。各企業の Web サイトにアクセスし、最も関連性の高い問題点を先頭に、パーソナライズされた電子メールを作成します。」
_脳/
§── .git/ ← 分離されたブレイン リポジトリ、プライベート GitHub リモートにプッシュ
§──index.md ← エージェントはセッションごとに最初にこれを読み取ります
§──dashboard.html ← ライブ コマンド センター、ブラウザで自動更新
§── core/ ← 製品、ブランド、ICP
§── オペレーション/ ← メトリクス + 顧客、毎日自動更新
§── スキル/ ← 永続的なルール、自動的に作成および更新されます
└──tasks/ ← スケジュールされたタスクのキュー、自動的に管理
→ git アーキテクチャの仕組み
購読はありません。完全に無料です。 100% オープンソース。
クロード コードは Anthropic の製品です。 Claude Code OS は独立したオープンソース プロジェクトであり、Anthropic と提携または承認されていません。
自己更新する運用副操縦士。自己改善、積極的、地元第一主義。
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

The operational co-pilot that self-updates. Self-improving, proactive, local-first. - bernardohcrocha/claude-code-os

GitHub - bernardohcrocha/claude-code-os: The operational co-pilot that self-updates. Self-improving, proactive, local-first. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
bernardohcrocha
/
claude-code-os
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
bernardohcrocha/claude-code-os
main Branches Tags Go to file Code Open more actions menu Folders and files
51 Commits 51 Commits scripts scripts templates templates ARCHITECTURE.md ARCHITECTURE.md FRAMEWORK.html FRAMEWORK.html LICENSE LICENSE README.md README.md claude-code-os.png claude-code-os.png setup.sh setup.sh View all files Repository files navigation
Self-updating agentic memory for Claude Code.
Turn Claude Code into an agent that remembers and acts seamlessly as your project changes.
Claude Code has no operational memory. Every session starts from zero — you re-explain your stack, paste context, repeat yourself.
Claude Code OS fixes that. A _brain/ folder lives in your project root and acts as persistent memory. It reads your project once, then self-updates daily from git diff — only what changed, nothing more.
Open Claude Code in your project root folder , then run:
curl -fsSL https://raw.githubusercontent.com/bernardohcrocha/claude-code-os/main/setup.sh | bash
Claude scans your project and asks only what it can't find. No forms. No config files. Just a conversation.
Optional: GitHub CLI ( gh ) — automatically creates a private brain repository with cloud backup on every update. Format your machine, git clone the brain repo, continue exactly where you left off.
Scheduler included: installs automatically (launchd on macOS, systemd on Linux).
Self-updating memory — runs git diff HEAD~1 daily. 1,000 files in the project, 3 changed today: it reads 3
Permanent skills — say it once → written to _brain/skills/ , loaded at every future session
Autonomous tasks — schedule any task in plain language, runs automatically with full context
Live dashboard — open _brain/dashboard.html , auto-refreshes every 5 min
Unlike Hermes Agent, Agent Zero, or OpenClaw — context updates itself. No manual maintenance required.
→ "Which customers dropped usage 30%+ this month? Cross-check their support history and draft a personalized re-engagement message for each."
→ "Every Monday: pull last week's numbers from Stripe, flag anomalies, and queue a follow-up for any account that dropped below quota."
→ "Find signups from the last 30 days with no activity. Visit each company's website and write a personalized email — leading with the most relevant pain points."
_brain/
├── .git/ ← isolated brain repository, pushes to private GitHub remote
├── index.md ← agent reads this first, every session
├── dashboard.html ← live command center, auto-refreshes in browser
├── core/ ← product, brand, ICP
├── operations/ ← metrics + customers, auto-updated daily
├── skills/ ← permanent rules, created and updated automatically
└── tasks/ ← scheduled tasks queue, managed automatically
→ How the git architecture works
No subscriptions. 100% free. 100% open source.
Claude Code is a product of Anthropic. Claude Code OS is an independent open-source project, not affiliated with or endorsed by Anthropic.
The operational co-pilot that self-updates. Self-improving, proactive, local-first.
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
