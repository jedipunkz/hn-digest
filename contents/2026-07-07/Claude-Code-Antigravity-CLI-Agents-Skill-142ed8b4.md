---
source: "https://github.com/markfulton/claude-antigravity-agents"
hn_url: "https://news.ycombinator.com/item?id=48817248"
title: "Claude Code Antigravity CLI Agents Skill"
article_title: "GitHub - markfulton/claude-antigravity-agents: A Claude Code skill that delegates coding, review, analysis, and research jobs to Google Antigravity CLI (agy) sub-agents — save tokens, use a second model, work in parallel. · GitHub"
author: "DotSauce"
captured_at: "2026-07-07T14:08:23Z"
capture_tool: "hn-digest"
hn_id: 48817248
score: 2
comments: 0
posted_at: "2026-07-07T13:07:48Z"
tags:
  - hacker-news
  - translated
---

# Claude Code Antigravity CLI Agents Skill

- HN: [48817248](https://news.ycombinator.com/item?id=48817248)
- Source: [github.com](https://github.com/markfulton/claude-antigravity-agents)
- Score: 2
- Comments: 0
- Posted: 2026-07-07T13:07:48Z

## Translation

タイトル: クロード コード反重力 CLI エージェント スキル
記事のタイトル: GitHub - markfulton/claude-antigravity-agents: コーディング、レビュー、分析、調査ジョブを Google Antigravity CLI (agy) サブエージェントに委任するクロード コード スキル — トークンを保存し、2 番目のモデルを使用し、並行して作業します。 · GitHub
説明: コーディング、レビュー、分析、調査ジョブを Google Antigravity CLI (agy) サブエージェントに委任するクロード コード スキル。トークンを保存し、2 番目のモデルを使用し、並行して作業します。 - マークフルトン/クロード反重力エージェント

記事本文:
GitHub - markfulton/claude-antigravity-agents: コーディング、レビュー、分析、調査ジョブを Google Antigravity CLI (agy) サブエージェントに委任するクロード コード スキル — トークンを保存し、2 番目のモデルを使用し、並行して作業します。 · GitHub
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
顕著な

トン
/
クロード反重力エージェント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マークフルトン/クロード-反重力エージェント
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット antigravity-agents antigravity-agents .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード・コード反重力エージェント
Claude Code をオーケストレーターに変えます。構築を続ける間、重いジョブを Google Antigravity サブエージェントに委任します。
作成者: Mark Fulton · Vibecoding is Life コミュニティの創設者 (メンバー数 300,000 人以上)
単一のクロード コード スキル — antigravity-agents という名前の 1 つのマークダウン命令ファイル ( SKILL.md )。スキルは、Claude Code がオンデマンドでロードする機能です。これは、Gemini やその他のモデルを利用した自律型端末コーディング エージェントである Google Antigravity CLI ( agy ) に作業を引き渡すように教えています。
スキルがロードされると、Claude Code は単一のワーカーであることをやめ、オーケストレーターになります。自己完結型のブリーフを作成し、バックグラウンドで AGY ジョブを起動し、タスクの作業を継続し、その結果を収集して検証してから、その行を信頼します。
トークンを保存します。フルリポジトリ監査、大規模なリファクタリング、リサーチスイープを Claude セッション内で実行するにはコストがかかります。作業中にクロード コンテキストとサブスクリプション トークンを書き込むのではなく、それらを Antigravity にオフロードします。
モデルの選択肢を広げます。純粋に独立したセカンドオピニオンを得るために、または単に適切なタスクに適切なモデルを配置するために、Gemini (または Antigravity が提供するもの) でジョブを実行します。
並行して作業します。 Claude Code が常に応答している間、複数のサブエージェント ジョブを一度に展開します。実時間の単位当たりにより多くのことが行われます。
クロードは概要を書き、

ジョブを非対話的に実行し ( agy -p "<task>" )、すぐに作業に戻ります。サブエージェントは単独で実行されます。終了すると、Claude は出力を読み取り、その出力が信頼されるかマージされる前に ( git diff 、 typecheck 、 build ) 検証します。
フローチャート LR
A[クロード・コード<br/>オーケストレーター] -->|自己完結型の概要を作成します| B["agy -p '<タスク>'"]
B -->|バックグラウンドで起動| C{職種}
C -->|読み取り専用| D[サンドボックスモード<br/>同時実行]
C -->|ファイルを書き込みます| E[分離されたワークツリー<br/>または別のリポジトリ]
あ～。タスクを構築し続けます。-> A
D --> F[結果]
E --> F[結果]
F -->|git diff · typecheck · build| G[クロードが検証する]
G -->|信頼できる| H[あなたの作品に組み込まれています]
読み込み中
前提条件
1. Antigravity CLI をインストールします。
カール -fsSL https://antigravity.google/cli/install.sh |バッシュ
Windows の場合 (PowerShell):
irm https://antigravity.google/cli/install.ps1 |アイエックス
詳細は公式ガイドに記載されています: https://antigravity.google/docs/cli/getting-started
ターミナルで agy を 1 回実行し、ブラウザで 1 回限りの Google OAuth サインインを完了します。
アジ
その後、Claude Code は agy ジョブを非対話的に起動できるようになり、それ以上のプロンプトは表示されなくなります。
このリポジトリのクローンを作成し、antigravity-agents フォルダーを Claude Code スキル ディレクトリにコピーします。 Claude Code は次回の実行時にそれを自動検出します。
git clone https://github.com/markfulton/claude-antigravity-agents.git
cp -r クロード-反重力エージェント/反重力エージェント ~ /.claude/skills/
Windows (PowerShell)
git clone https://github.com/markfulton/claude - antigravity - Agents.git
Copy-Item - Recurse クロード - antigravity - Agents\antigravity - Agents " $ env: USERPROFILE \.claude\skills\ "
それだけです。 Claude Code セッションを開始し、何かを委任するように依頼します。スキルはオンデマンドで読み込まれます。
これまでと同じようにクロード コードに話しかけるだけです。数p

スキルをトリガーするロンプト:
「チェックアウト フローの構築を続けている間、Antigravity に現在のブランチの差分をバグがないかレビューしてもらいます。」
「サブエージェントを起動して、このリポジトリの依存関係に既知の脆弱性がないか監査し、報告します。」
「このキャッシュ設計について Gemini からセカンドオピニオンを得て、その間 API の開発を続けてください。」
クロードはジョブの種類を決定し、適切な分離モードを選択してバックグラウンドで起動し、結果を収集して検証したらレポートを返します。
委任は、戻ってくるものを信頼できる場合にのみ役に立ちます。 1 つのルールが他のすべてより優先されます。
Antigravity ジョブは、その時点でクロード コードが編集しているのと同じファイルに書き込むことはありません。
このルールにより、すべてのジョブがどのように実行されるかが決まります。
信頼される前にすべてが検証されます。クロードは git diff を検査し、返された作業を折り込む前にプロジェクトのタイプチェック/ビルドを実行します。サブエージェントの変更は信頼に基づいて行われません。タスクが小さい場合、現在編集中の内容と密接に結合している場合、またはライブでのやり取りが必要な場合は、タスクを Claude に保持します。委任は、重く分離可能な作業に重点を置きます。
Vibecoding is Life コミュニティの創設者 (メンバー 30 万人以上)
MIT ライセンスに基づいてリリースされています。
コーディング、レビュー、分析、調査ジョブを Google Antigravity CLI (agy) サブエージェントに委任するクロード コード スキル。トークンを保存し、2 番目のモデルを使用し、並行して作業します。
読み込み中にエラーが発生しました。このページをリロードしてください。
16
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Claude Code skill that delegates coding, review, analysis, and research jobs to Google Antigravity CLI (agy) sub-agents — save tokens, use a second model, work in parallel. - markfulton/claude-antigravity-agents

GitHub - markfulton/claude-antigravity-agents: A Claude Code skill that delegates coding, review, analysis, and research jobs to Google Antigravity CLI (agy) sub-agents — save tokens, use a second model, work in parallel. · GitHub
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
markfulton
/
claude-antigravity-agents
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
markfulton/claude-antigravity-agents
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits antigravity-agents antigravity-agents .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Claude Code Antigravity Agents
Turn Claude Code into an orchestrator — delegate heavy jobs to Google Antigravity sub-agents while it keeps building.
Created by Mark Fulton · Founder of the Vibe Coding is Life community (300k+ members)
A single Claude Code skill — one markdown instruction file ( SKILL.md ) named antigravity-agents . Skills are capabilities Claude Code loads on demand. This one teaches it to hand off work to the Google Antigravity CLI ( agy ), an autonomous terminal coding agent powered by Gemini and other models.
With the skill loaded, Claude Code stops being a single worker and becomes an orchestrator : it writes a self-contained brief, launches an agy job in the background, keeps working on your task, then collects and verifies the result before trusting a line of it.
Save your tokens. Full-repo audits, big refactors, and research sweeps are expensive to run inside a Claude session. Offload them to Antigravity instead of burning your Claude context and subscription tokens on the grind work.
Broaden your model choice. Run a job on Gemini (or whatever Antigravity offers) for a genuinely independent second opinion — or just to put the right model on the right task.
Work in parallel. Fan out several sub-agent jobs at once while Claude Code stays responsive to you. More gets done per unit of wall-clock time.
Claude writes the brief, launches the job non-interactively ( agy -p "<task>" ), and immediately returns to your work. The sub-agent runs on its own. When it finishes, Claude reads the output and verifies it — git diff , typecheck, build — before any of it is trusted or merged.
flowchart LR
A[Claude Code<br/>orchestrator] -->|writes self-contained brief| B["agy -p '&lt;task&gt;'"]
B -->|launched in background| C{Job type}
C -->|read-only| D[Sandbox mode<br/>runs concurrently]
C -->|writes files| E[Isolated worktree<br/>or separate repo]
A -. keeps building your task .-> A
D --> F[Result]
E --> F[Result]
F -->|git diff · typecheck · build| G[Claude verifies]
G -->|trusted| H[Folded into your work]
Loading
Prerequisites
1. Install the Antigravity CLI.
curl -fsSL https://antigravity.google/cli/install.sh | bash
On Windows (PowerShell):
irm https: // antigravity.google / cli / install.ps1 | iex
Full details are in the official guide: https://antigravity.google/docs/cli/getting-started
Run agy once in a terminal and complete the one-time Google OAuth sign-in in your browser:
agy
After that, Claude Code can launch agy jobs non-interactively — no further prompts.
Clone this repo and copy the antigravity-agents folder into your Claude Code skills directory. Claude Code auto-discovers it on the next run.
git clone https://github.com/markfulton/claude-antigravity-agents.git
cp -r claude-antigravity-agents/antigravity-agents ~ /.claude/skills/
Windows (PowerShell)
git clone https: // github.com / markfulton / claude - antigravity - agents.git
Copy-Item - Recurse claude - antigravity - agents\antigravity - agents " $ env: USERPROFILE \.claude\skills\ "
That's it. Start a Claude Code session and ask it to delegate something — the skill loads on demand.
Just talk to Claude Code the way you already do. A few prompts that trigger the skill:
"Have Antigravity review my current branch diff for bugs while you keep building the checkout flow."
"Spin up a sub-agent to audit this repo's dependencies for known vulnerabilities and report back."
"Get a second opinion from Gemini on this caching design, and keep working on the API in the meantime."
Claude decides the job type, picks the right isolation mode, launches it in the background, and reports back once it has collected and verified the result.
Delegation only helps if you can trust what comes back. One rule sits above everything else:
An Antigravity job never writes to the same files Claude Code is editing at that moment.
That rule drives how every job runs:
Everything is verified before it's trusted — Claude inspects the git diff and runs the project's typecheck / build on any returned work before folding it in. No sub-agent change is taken on faith. Keep a task in Claude when it's small, tightly coupled to what you're editing right now, or needs live back-and-forth; delegation shines on the heavy, separable work.
Founder of the Vibe Coding is Life community (300k+ members)
Released under the MIT License .
A Claude Code skill that delegates coding, review, analysis, and research jobs to Google Antigravity CLI (agy) sub-agents — save tokens, use a second model, work in parallel.
There was an error while loading. Please reload this page .
16
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
