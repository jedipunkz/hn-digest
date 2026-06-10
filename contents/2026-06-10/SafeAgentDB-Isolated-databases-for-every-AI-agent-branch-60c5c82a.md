---
source: "https://github.com/Aidan945/SafeAgentDB"
hn_url: "https://news.ycombinator.com/item?id=48470846"
title: "SafeAgentDB – Isolated databases for every AI agent branch"
article_title: "GitHub - Aidan945/SafeAgentDB: Isolated databases for every AI agent branch. Build, test, and migrate in parallel. Production stays untouched. · GitHub"
author: "zarskia9"
captured_at: "2026-06-10T04:34:49Z"
capture_tool: "hn-digest"
hn_id: 48470846
score: 2
comments: 0
posted_at: "2026-06-10T02:59:47Z"
tags:
  - hacker-news
  - translated
---

# SafeAgentDB – Isolated databases for every AI agent branch

- HN: [48470846](https://news.ycombinator.com/item?id=48470846)
- Source: [github.com](https://github.com/Aidan945/SafeAgentDB)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T02:59:47Z

## Translation

タイトル: SafeAgentDB – AI エージェント ブランチごとに分離されたデータベース
記事のタイトル: GitHub - Aidan945/SafeAgentDB: AI エージェント ブランチごとに分離されたデータベース。構築、テスト、移行を並行して行います。生産はそのまま残ります。 · GitHub
説明: AI エージェント ブランチごとに分離されたデータベース。構築、テスト、移行を並行して行います。生産はそのまま残ります。 - Aidan945/SafeAgentDB

記事本文:
GitHub - Aidan945/SafeAgentDB: AI エージェント ブランチごとに分離されたデータベース。構築、テスト、移行を並行して行います。生産はそのまま残ります。 · GitHub
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
エイダン945
/
セーフエージェントDB
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット スキル/safeagentdb スキル/safeagentdb .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AFK エージェント開発のためのデータベース安全インフラストラクチャ。
SafeAgentDB は、チームメイトが人間、AI エージェント、またはその両方であるかどうかにかかわらず、チームで実際の製品を出荷する本格的なバイブコーダー向けに構築されています。多くのブランチを同時に運用すると、1 つの共有データベースが誰もが壊すものになります。 SafeAgentDB は、すべてのブランチと PR に、独自の分離データベースに裏付けられたライブ プレビュー URL を提供します。そのため、エージェントは、本番環境を危険にさらしたり、共有開発データを破損したり、相互に踏み込んだりすることなく、移行を実行し、現実的なデータをハイドレートして、実際のアプリの動作をテストできます。
最初に Supabase + Vercel + GitHub Actions 用に構築され、同じインフラストラクチャ パターンを他のスタックに適応させるためのガイダンスが含まれています。
SafeAgentDB は、オープン エージェント スキル ディレクトリである skill.sh と、OpenClaw スキル レジストリである ClawHub にリストされています。
npx スキル追加 https://github.com/Aidan945/SafeAgentDB --skillsafeagentdb
略記:
npx スキル追加 Aidan945/SafeAgentDB --skillsafeagentdb
特定のエージェントをターゲットにすることもできます。
npx スキル追加 Aidan945/SafeAgentDB --skillsafeagentdb --agentcursor
npx スキル追加 Aidan945/SafeAgentDB --skillsafeagentdb --agent codex
npx スキル追加 Aidan945/SafeAgentDB --skillsafeagentdb --agent claude-code
または、OpenClaw を使用してインストールします。
openclaw スキルでsafeagentdbをインストール
スキルを使用する
インストール後、SafeAgentDB をインストールするプロジェクトでエージェントを起動し、次のように言います。
safeagentdb スキルを使用します。
このプロジェクトを検査し、SafeAgentDB データベースの安全性インフラストラクチャをセットアップします。 PLを見直してください

変更を加える前に私にご相談ください。
残りの部分はスキルが処理します。つまり、スタックを検査し、デフォルト以外の設定に適応し、資格情報に触れる前に質問し、シークレットを決してコミットしません。
SafeAgentDB は、AI エージェントのセットアップを支援します。
ローカルの Docker データベース開発
永続的な開発/ステージング データベース
ライブ プレビュー URL を備えたブランチ固有のプレビュー データベース
安全な移行フロー: マージ前のチェック、プレビューのみの機能移行、メインからの実稼働デプロイ
オプションで現実的なデータ、認証ユーザー、ストレージをプレビューにハイドレーションします
Vercel プレビュー環境ワイヤリングと環境変更後の自動再デプロイ
PR クローズ時のクリーンアップとスケジュールされた孤立クリーンアップ。永続的なデザイン/デモ環境はスキップされます。
プロジェクトの信頼できる情報源としてのドライラン プロビジョニングとコミットされたブランチ アーキテクチャ ドキュメント
その結果、エージェントは本番データや共有開発データに触れることなく、並行して作業し、PR を公開し、テスト用のライブ URL をユーザーに提供できるようになります。
デフォルトのテンプレートは以下のために構築されています。
Postgres、認証、ストレージ、移行、データベース ブランチ用の Supabase
本番環境、ステージング環境、プレビュー環境向けの Vercel
ブランチと PR の自動化のための GitHub アクション
プロジェクトで別のデータベース、展開プラットフォーム、または CI プロバイダーを使用する場合は、SafeAgentDB を概念モデルとして扱う必要があります。 AI エージェントはコードベースを検査し、何を変更する必要があるかを説明し、パターンをインフラストラクチャに適応させる必要があります。
エージェントに必要なものはすべてスキル内にバンドルされています。
スキル/
セーフエージェントdb/
スキル.md
参考文献/
セットアッププロセス.md
資格情報.md
データハイドレーションポリシー.md
ローカル開発.md
非標準スタック.md
トラブルシューティング.md
エージェント操作ルール.md
エージェントパッケージング.md
テンプレート/
ブランチ構成.example.json
パッケージスクリプト.json
パッケージ開発依存関係.json
ドキュメント/
データベース分岐.md
スクリプト/
スーパーベース/
c

私/
.github/
ワークフロー/
エージェントの指示/
エージェント.md
クロード.md
安全原則
SafeAgentDB は、いくつかのルールに基づいて構築されています。
本番環境ではエージェント機能の動作を決して指示しないでください。
複数のエージェントが 1 つの変更可能なプレビュー データベースを共有しないようにしてください。
マージするまでは、機能の移行を共有開発に決して適用しないでください。
明示的な承認なしに実稼働データをプレビューにコピーしないでください。
共有環境の前に、必ずローカルで移行をテストするか、プレビューで移行をテストしてください。
実稼働環境、開発環境、プレビュー環境、およびローカル環境を分離してください。
ブランチが閉じられるか削除されるときに、プレビュー データベースをクリーンアップします。
SafeAgentDB は、インストール可能なエージェント スキルおよびインフラストラクチャ テンプレート キットです。
ワンクリックの SaaS 製品ではありません。このスキルは、AI エージェントにプロジェクトの検査、正しい認証情報の要求、ハイドレーション ポリシーの選択、スタックに最も安全なバージョンのワークフローのインストールを指示します。
AI エージェント ブランチごとに分離されたデータベース。構築、テスト、移行を並行して行います。生産はそのまま残ります。
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

Isolated databases for every AI agent branch. Build, test, and migrate in parallel. Production stays untouched. - Aidan945/SafeAgentDB

GitHub - Aidan945/SafeAgentDB: Isolated databases for every AI agent branch. Build, test, and migrate in parallel. Production stays untouched. · GitHub
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
Aidan945
/
SafeAgentDB
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits skills/ safeagentdb skills/ safeagentdb .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Database safety infrastructure for AFK agentic development.
SafeAgentDB is built for serious vibe coders shipping real products with a team — whether your teammates are people, AI agents, or both. Once you have many branches in flight at the same time, one shared database becomes the thing everyone breaks. SafeAgentDB gives every branch and PR a live preview URL backed by its own isolated database, so agents can run migrations, hydrate realistic data, and test real app behavior without risking production, corrupting shared development data, or stepping on each other.
Built first for Supabase + Vercel + GitHub Actions , with guidance for adapting the same infrastructure pattern to other stacks.
SafeAgentDB is listed on skills.sh , the open agent skills directory, and on ClawHub , the OpenClaw skill registry.
npx skills add https://github.com/Aidan945/SafeAgentDB --skill safeagentdb
Shorthand:
npx skills add Aidan945/SafeAgentDB --skill safeagentdb
You can also target a specific agent:
npx skills add Aidan945/SafeAgentDB --skill safeagentdb --agent cursor
npx skills add Aidan945/SafeAgentDB --skill safeagentdb --agent codex
npx skills add Aidan945/SafeAgentDB --skill safeagentdb --agent claude-code
Or install with OpenClaw:
openclaw skills install safeagentdb
Use The Skill
After installing, start your agent in the project where you want SafeAgentDB installed and say:
Use the safeagentdb skill.
Inspect this project, then set up SafeAgentDB database safety infrastructure for it. Review your plan with me before making changes.
The skill handles the rest: inspecting your stack, adapting to non-default setups, asking before touching credentials, and never committing secrets.
SafeAgentDB helps AI agents set up:
Local Docker database development
A persistent develop/staging database
Branch-specific preview databases with live preview URLs
Safe migration flow: checks before merge, preview-only feature migrations, production deploys from main
Optional hydration of realistic data, auth users, and storage into previews
Vercel preview env wiring with automatic redeploys after env changes
Cleanup on PR close plus scheduled orphan cleanup, with persistent design/demo environments skipped
Dry-run provisioning and a committed branching architecture doc as the project's source of truth
The result: agents can work in parallel, publish PRs, and give users live URLs to test without touching production or shared development data.
The default templates are built for:
Supabase for Postgres, Auth, Storage, migrations, and database branches
Vercel for production, staging, and preview deployments
GitHub Actions for branch and PR automation
If your project uses a different database, deployment platform, or CI provider, SafeAgentDB should be treated as a conceptual model. The AI agent should inspect your codebase, explain what needs to change, and adapt the pattern to your infrastructure.
Everything the agent needs is bundled inside the skill:
skills/
safeagentdb/
SKILL.md
references/
setup-process.md
credentials.md
data-hydration-policy.md
local-development.md
non-standard-stacks.md
troubleshooting.md
agent-operating-rules.md
agent-packaging.md
templates/
branching-config.example.json
package-scripts.json
package-dev-dependencies.json
docs/
database-branching.md
scripts/
supabase/
ci/
.github/
workflows/
agent-instructions/
AGENTS.md
CLAUDE.md
Safety Principles
SafeAgentDB is built around a few rules:
Never point agent feature work at production.
Never let multiple agents share one mutable preview database.
Never apply feature migrations to shared development until merge.
Never copy production data into previews without explicit approval.
Always test migrations locally or in preview before shared environments.
Keep production, develop, preview, and local environments separate.
Clean up preview databases when branches are closed or deleted.
SafeAgentDB is an installable agent skill and infrastructure template kit.
It is not a one-click SaaS product. The skill walks an AI agent through inspecting your project, asking for the right credentials, choosing a hydration policy, and installing the safest version of the workflow for your stack.
Isolated databases for every AI agent branch. Build, test, and migrate in parallel. Production stays untouched.
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
