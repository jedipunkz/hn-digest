---
source: "https://github.com/bernardohcrocha/claude-code-os/"
hn_url: "https://news.ycombinator.com/item?id=48351152"
title: "Claude Code OS: self-updating operational memory for Claude Code (open source)"
article_title: "GitHub - bernardohcrocha/claude-code-os: The operational co-pilot for solo founders. Self-improving, proactive, local-first. · GitHub"
author: "bernardohcr"
captured_at: "2026-06-01T00:24:15Z"
capture_tool: "hn-digest"
hn_id: 48351152
score: 2
comments: 0
posted_at: "2026-06-01T00:08:21Z"
tags:
  - hacker-news
  - translated
---

# Claude Code OS: self-updating operational memory for Claude Code (open source)

- HN: [48351152](https://news.ycombinator.com/item?id=48351152)
- Source: [github.com](https://github.com/bernardohcrocha/claude-code-os/)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T00:08:21Z

## Translation

タイトル: Claude Code OS: Claude Code の自動更新動作メモリ (オープンソース)
記事のタイトル: GitHub - bernardohcrocha/claude-code-os: 個人創業者のための運用副操縦士。自己改善、積極的、地元第一主義。 · GitHub
説明: 個人創業者向けの運用副操縦士。自己改善、積極的、地元第一主義。 - bernardohcrocha/claude-code-os

記事本文:
GitHub - bernardohcrocha/claude-code-os: 個人創業者向けの運用副操縦士。自己改善、積極的、地元第一主義。 · GitHub
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
42 コミット 42 コミット スクリプト スクリプト テンプレート テンプレート ARCHITECTURE.md ARCHITECTURE.md FRAMEWORK.html FRAMEWORK.html LICENSE LICENSE README.md README.md claude-code-os.png claude-code-os.png setup.sh setup.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
自動更新する運用副操縦士。
自己進化 · 追加ツール不要 · ワークフロー変更なし · トークン効率 · 自動 git バックアップ
クロードコードはうまく実行されます。問題は、永続的な動作メモリがないことです。
すべてのセッションはゼロから始まります。プロジェクト内で変更されたすべての新機能、すべての新しいドキュメント、すべての詳細について、再説明するか、プロジェクト全体をコンテキストとしてダンプしてトークン コストを支払う必要がありました。いずれにせよ、あなたはエージェントが行うべき仕事を行っています。
Claude Code OS はそれを修正します。 _brain/ フォルダーはプロジェクト ルート内にあります。最初の実行時にコード、ドキュメント、構成を読み取ります。それ以降、実際に変更された内容に基づいて、毎日自動的に更新されます。
新しい支払い方法を発送します。製品知識、それによって明らかにされる顧客セグメント、および追跡を開始する指標が更新されます。あなたは何も言わなかった。
ブランド資産をフォルダーに追加します。ブランドガイドラインは自動的に更新されます。
1 か月後に新しいマシンに戻ってきました。ファイルを読み取り、拍子抜けすることなく続行します。スタックを再度説明する必要はありません。コンテキストをコピー＆ペーストする必要はありません。何も失われませんでした。
単純なダッシュボードでは答えられない質問をしてください。または、完全なコンテキストがすでに読み込まれた状態で、自動的に実行されるタスクをスケジュールします。
→ 「今月使用量が 30% 以上減少したのはどの顧客ですか? サポート履歴を照合し、EA 向けにパーソナライズされた再エンゲージメント メッセージを作成します。

ch.」
→ 「毎週月曜日: Stripe から先週の数値を取得し、前週と比較し、異常を報告し、10% の割り当てを下回ったアカウントのフォローアップをキューに入れます。」
→「過去 30 日間にサインアップ後にアクティビティがなかったサインアップを検索します。偽のドメインを除外します。各企業の Web サイトにアクセスして、その事業内容を理解し、同様の顧客がすでに製品をどのように使用しているかにマッピングし、特定のユースケースに最も関連する問題点を先頭に、それぞれにパーソナライズされた電子メールを作成します。」
一度設定してください。自動的に実行されます。プロジェクト内に存在するため、完全なコンテキストがすでに存在します。
ブラウザで _brain/dashboard.html を開きます。 5分ごとに自動更新されます。
固定テンプレートはありません。ダッシュボードは実際のデータから構築され、ビジネスの進行に合わせて進化します。今何が重要かに基づいて、異なるセクション、異なる指標、異なる優先順位が設定されます。
スケジュールされたタスク、プロアクティブな提案、アクティブなプロジェクト、週次レビュー: すべてバックグラウンドで自動的に更新されます。会話を開かなくても結果が表示されます。
何かを変えたいですか?ただそれを伝えてください。
プロジェクトがアイドル状態の場合、3 日ごとに指標、顧客、チャネルをスキャンし、気づいたことをメモに残します。常に提案します。決して単独で行動することはありません。
毎日、 git diff HEAD~1 は何が変更されたかを正確に検出します。プロジェクト内に 1,000 個のファイルがあり、今日変更されたファイルは 3 つです。読み取り値は 3 です。完全な再読み込みではありません。まさにデルタ。
1 つの変更されたファイルは、製品知識、顧客セグメント、指標、スキルなど、すべての要素を一度に更新します。自動。軽量。
すべての指示は永続的になります。
「不正行為を報告するときは、使い捨て電子メール ドメイン、重複した名前、サインアップのタイミングを常にクロスチェックしてください。」と一度言ってみましょう。
すぐに _brain/skills/fraud.md を書き込みます。今後のセッション開始ごとにロードされます。自動的に適用されます。

リバー。リマインダーはありません。編集ファイルはありません。再説明はありません。
Wispr Flow または Handy.computer を追加すると、机の上、通勤中、会議の合間に歩いているなど、どこからでも命令できる本物のスーパー従業員のような気分になります。
「クォータに近づいている Pro アカウントがあるかどうかを確認し、それぞれについて注意事項を作成します。」終わり。コーヒーを淹れている間。
コマンドは 1 つです。そこからそれ自体が構築されます。
プロジェクトのルート フォルダーで Claude Code を開き、次を実行します。
カール -fsSL https://raw.githubusercontent.com/bernardohcrocha/claude-code-os/main/setup.sh |バッシュ
クロードはプロジェクト全体をスキャンし、既存のツールに接続して、見つからないものだけを尋ねます。フォームはありません。設定ファイルはありません。ただの会話です。
オプション: GitHub CLI ( gh ) — インストールして認証すると、セットアップによりプライベート ブレイン リポジトリが自動的に作成され、コミットごとにクラウド バックアップが有効になります。マシンをフォーマットし、Brain リポジトリを git clone して、中断したところから正確に続行します。
スケジューラーが含まれています: 自動的にインストールされます (macOS では launchd、Linux では systemd)。スリープまたは再起動後に見逃したタスクを取り戻します。
_脳/
§── .git/ ← 分離されたブレイン リポジトリ、プライベート GitHub リモートにプッシュ
§──index.html ← エージェントはセッションごとにこれを最初に読み取ります
§──dashboard.html ← ライブ コマンド センター、ブラウザで自動更新
§── core/ ← 製品、ブランド、ICP
§── オペレーション/ ← メトリクス + 顧客、毎日自動更新
§── スキル/ ← 永続的なルール、自動的に作成および更新されます
└──tasks/ ← スケジュールされたタスクのキュー、自動的に管理
→ git アーキテクチャの仕組み
購読はありません。完全に無料です。 100% オープンソース。
フォークしてください。調整してください。それをあなたのものにしてください。
クロード コードは Anthropic の製品です。 Claude Code OS は独立したオープンソース プロジェクトであり、Anthropic と提携または承認されていません。
運航副操縦士

個人の創業者向け。自己改善、積極的、地元第一主義。
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

The operational co-pilot for solo founders. Self-improving, proactive, local-first. - bernardohcrocha/claude-code-os

GitHub - bernardohcrocha/claude-code-os: The operational co-pilot for solo founders. Self-improving, proactive, local-first. · GitHub
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
42 Commits 42 Commits scripts scripts templates templates ARCHITECTURE.md ARCHITECTURE.md FRAMEWORK.html FRAMEWORK.html LICENSE LICENSE README.md README.md claude-code-os.png claude-code-os.png setup.sh setup.sh View all files Repository files navigation
The Operational Copilot that Self-Updates.
Self-evolving · No extra tools · No workflow changes · Token-efficient · Auto git backup
Claude Code executes well. The problem is it has no persistent operational memory .
Every session starts from zero. Every new feature, every new document, every detail that changed in your project: you had to re-explain it, or dump the entire project as context and pay the token cost. Either way, you're doing work the agent should be doing.
Claude Code OS fixes that. A _brain/ folder lives inside your project root. It reads from your code, docs, and configs on first run. From then on, it updates itself every day, automatically, based on what actually changed.
You ship a new payment method. It updates the product knowledge, the customer segments it unlocks, and the metrics to start tracking. You never said a word.
You add brand assets to a folder. Brand guidelines updated automatically.
You come back after a month on a new machine. It reads the files and continues without skipping a beat. No re-explaining your stack. No copy-pasting context. Nothing lost.
Ask questions no simple dashboard can answer. Or schedule tasks that run automatically, with full context already loaded.
→ "Which customers dropped usage 30%+ this month? Cross-check their support history and draft a personalized re-engagement message for each."
→ "Every Monday: pull last week's numbers from Stripe, compare against the previous week, flag anomalies, and queue a follow-up for any account that dropped below 10% quota."
→ "Find signups from the last 30 days with no activity after signup. Filter out fake-looking domains. Visit each company's website, understand what they do, map it to how similar customers already use the product, and write a personalized email for each, leading with the pain points most relevant to their specific use case."
Set once. Runs automatically. Already has full context because it lives in the project.
Open _brain/dashboard.html in your browser. It auto-refreshes every 5 minutes.
No fixed template. The dashboard is built from your actual data and evolves as your business does. Different sections, different metrics, different priorities based on what matters right now.
Scheduled tasks, proactive suggestions, active projects, weekly review: all updated automatically in the background. You see the results without opening a conversation.
Want to change anything? Just tell it.
Every 3 days, when the project is idle, it scans across metrics, customers, and channels and leaves a note with what it noticed. Always suggests. Never acts on its own.
Every day, git diff HEAD~1 detects exactly what changed. 1,000 files in the project, 3 changed today: it reads 3. Not a full re-read. Just the delta.
One changed file updates every dimension it touches at once: product knowledge, customer segments, metrics, skills. Automatic. Lightweight.
Every instruction becomes permanent.
Say it once: "When flagging fraud, always cross-check disposable email domains, duplicate names, and signup timing."
It writes _brain/skills/fraud.md immediately. Loaded at every future session start. Applied automatically, forever. No reminders. No editing files. No re-explaining.
Add Wispr Flow or handy.computer and it starts feeling like a real super employee you can give orders to from anywhere: desk, commute, walking between meetings.
"Check if any Pro accounts are near quota and draft a heads-up for each." Done. While you're making coffee.
One command. It builds itself from there.
Open Claude Code in your project root folder , then run:
curl -fsSL https://raw.githubusercontent.com/bernardohcrocha/claude-code-os/main/setup.sh | bash
Claude scans your entire project, connects to your existing tools, and asks only what it can't find. No forms. No config files. Just a conversation.
Optional: GitHub CLI ( gh ) — if installed and authenticated, setup automatically creates a private brain repository and enables cloud backup with every commit. Format your machine, git clone the brain repo, continue exactly where you left off.
Scheduler included: installs automatically (launchd on macOS, systemd on Linux). Catches up on missed tasks after sleep or restart.
_brain/
├── .git/ ← isolated brain repository, pushes to private GitHub remote
├── index.html ← agent reads this first, every session
├── dashboard.html ← live command center, auto-refreshes in browser
├── core/ ← product, brand, ICP
├── operations/ ← metrics + customers, auto-updated daily
├── skills/ ← permanent rules, created and updated automatically
└── tasks/ ← scheduled tasks queue, managed automatically
→ How the git architecture works
No subscriptions. 100% free. 100% open source.
Fork it. Adjust it. Make it yours.
Claude Code is a product of Anthropic. Claude Code OS is an independent open-source project, not affiliated with or endorsed by Anthropic.
The operational co-pilot for solo founders. Self-improving, proactive, local-first.
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
