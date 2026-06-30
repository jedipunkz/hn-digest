---
source: "https://github.com/ohad6k/VibeRaven"
hn_url: "https://news.ycombinator.com/item?id=48730693"
title: "The missing context layer for AI-built apps"
article_title: "GitHub - ohad6k/VibeRaven: Open-source mission control for AI-built apps. The operating layer for apps built by AI agents — providers, releases, production evidence, and version drift from first demo to real scale. · GitHub"
author: "ohadkr"
captured_at: "2026-06-30T11:01:05Z"
capture_tool: "hn-digest"
hn_id: 48730693
score: 1
comments: 0
posted_at: "2026-06-30T10:30:30Z"
tags:
  - hacker-news
  - translated
---

# The missing context layer for AI-built apps

- HN: [48730693](https://news.ycombinator.com/item?id=48730693)
- Source: [github.com](https://github.com/ohad6k/VibeRaven)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T10:30:30Z

## Translation

タイトル: AI で構築されたアプリに不足しているコンテキスト層
記事のタイトル: GitHub - ohad6k/VibeRaven: AI で構築されたアプリのためのオープンソース ミッション コントロール。 AI エージェントによって構築されたアプリのオペレーティング層 (プロバイダー、リリース、本番環境の証拠、最初のデモから実際のスケールまでのバージョンのドリフト)。 · GitHub
説明: AI で構築されたアプリのためのオープンソースのミッション コントロール。 AI エージェントによって構築されたアプリのオペレーティング層 (プロバイダー、リリース、本番環境の証拠、最初のデモから実際のスケールまでのバージョンのドリフト)。 - ohad6k/バイブレイブン

記事本文:
GitHub - ohad6k/VibeRaven: AI で構築されたアプリ用のオープンソース ミッション コントロール。 AI エージェントによって構築されたアプリのオペレーティング層 (プロバイダー、リリース、本番環境の証拠、最初のデモから実際のスケールまでのバージョンのドリフト)。 · GitHub
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
ohad6k
/
バイブR

エイブン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
79 コミット 79 コミット .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor/ rules .cursor/ rules .github .github .viberaven .viberaven エージェント-スキル エージェント-スキル アセット アセット コマンド コマンド ドキュメント ドキュメントの例 例 スキル/ バイブラベン スキル/ バイブラベン .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md COTRIBUTING.md EXPORT_MANIFEST.json EXPORT_MANIFEST.json GEMINI.md GEMINI.md ライセンス ライセンス README.md README.md after-install.md after-install.md エージェントコンテキスト.md Agent-context.md gemini-extension.json gemini-extension.json llms-full.txt llms-full.txt llms.txt llms.txt plugin.yaml plugin.yaml skill.sh.json skill.sh.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI で構築されたアプリを雰囲気ではなく証拠とともに出荷します。
フル品質の MP4 デモを開く
VibeRaven は、VibeRaven Production Skills に AI で構築されたアプリ用のローカル Studio コックピットを加えたものです。エージェントが作業の準備ができたと主張する前に、プロバイダー コンテキスト、リリース コンテキスト、承認を意識したチャット、安全なリポジトリの変更、および証拠が必要な場合に、リリースの構築、修正、レビュー、またはリリースの準備中にこれを使用します。
npx -y バイブレーベン
パブリック リポジトリは、プラグイン スタイルのメタデータ、移植可能なスラッシュ コマンド、制作スキル、MCP ノート、AI 可読ドキュメントなど、エージェントの検出とインストールの表面です。
AI コーディング エージェントは、コードにパッチを適用するのは得意ですが、証明できないことを知るのは苦手です。 VibeRaven プロダクション スキルは受動的レポートではありません。これらは、エージェントが起動に依存する作業の実装、レビュー、デバッグ、または準備中に使用できるタスク ワークフローです。
VibeRaven はエージェントに次のような契約を与えます。
pr からリポジトリコードを個別に修正

ダッシュボードの作業を提供します。
use MCP or connected provider context when available;
リリースとバージョンのコンテキストを表示したままにします。
危険なローカル変更を行う前に承認を求めてください。
これは、機能の作業、認証の修正、請求の変更、デプロイのデバッグ、リリースの比較、リリース後のドリフト、および実稼働の準備の際に重要です。
任意のプロジェクトからローカル コックピットを開きます。
npx -y バイブレーベン
Studio では、エージェント チャット、ドラッグ可能なプロバイダー、ドラッグ可能なバージョン/リリース、プロバイダー MCP コンテキスト、端末出力、差分コンテキスト、アクセス モード制御、および CLI エージェント接続チェックが提供されます。インストール済みと接続済みは同じではありません。 test the agent connection before real chat control.
Install bounded VibeRaven guidance for Codex, Claude Code, Cursor, Copilot, Gemini, and related agents:
npx -y バイブラベン init --agents all
ファイルを書き込まずにプレビューする:
npx -y viberaven init --agents all --dry-run
インストーラーは、 AGENTS.md 、 CLAUDE.md 、 GEMINI.md 、カーソル ルール、Copilot 命令、および .viberaven コンテキスト ファイルを含む、サポートされているエージェント命令ファイルにルールを書き込みます。
VibeRaven ships as a portable skill/plugin-style pack:
コーデックス: .codex-plugin/plugin.json
クロード コード: .claude-plugin/plugin.json
Gemini CLI: gemini-extension.json
汎用プラグインホスト: plugin.yaml
スラッシュコマンドプロンプト:commands/
/viberaven-help : show the pack and output contract.
/viberaven-work : use Production Skills on a task.
/viberaven-launch : run launch readiness before saying ready.
/viberaven-provider-actions : separate repo-code fixes from dashboard work.
docs/agent-portability.md を参照してください。
Browse the skill library in docs/production-skills.md .
npx -y skills add ohad6k/VibeRaven --skill viberaven
npx -y skills add ohad6k/VibeRaven --skill what-broke
npx -y skills add ohad6k/VibeRaven --skill go-live
MCP
VibeRaven は MCP 対応に公開される可能性があります

エージェント:
{ "viberaven" : { "command" : " npx " , "args" : [ " -y " 、 " viberaven " 、 " --mcp " ] } }
MCP 出力を Studio 対応エージェントのプロバイダーおよび準備コンテキストとして使用します。また、リポジトリは MCP レジストリ メタデータを docs/mcp-registry-submission.md に保存するため、メンテナーは npm および skill.sh メタデータとともにパブリック MCP ディスカバリ サーフェスを確認できます。
Examples/nextjs-supabase-vercel-production-ready-template には、Next.js + Supabase + Vercel のエージェント ルールと viberaven:* スクリプトが含まれています。
古い VibeRaven ドキュメントおよび互換性ツールには、エージェント モード スキャン、タスク リスト、ゲート結果、PRP リソース、またはスキャン由来のアクション マニフェストについて言及されている場合があります。その言語を主要な公開製品の表面ではなく、互換性コンテキストとして扱います。
現在のデフォルトはスタジオです。
npx -y バイブレーベン
通常の git プッシュはゲートされていません。 VibeRaven の準備状況に関する言語は、運用変更の信頼性、リリース レビュー、プロバイダー認識の証拠、およびエージェントの境界に関するものであり、通常のリポジトリ作業を妨げるものではありません。
問題: ohad6k/VibeRaven/問題
パブリックディスカバリーリポジトリ: ohad6k/VibeRaven
AI で構築されたアプリのためのオープンソースのミッション コントロール。 AI エージェントによって構築されたアプリのオペレーティング層 (プロバイダー、リリース、本番環境の証拠、最初のデモから実際のスケールまでのバージョンのドリフト)。
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
2
VibeRaven スタジオ デモ (26 秒)
最新の
2026 年 6 月 26 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source mission control for AI-built apps. The operating layer for apps built by AI agents — providers, releases, production evidence, and version drift from first demo to real scale. - ohad6k/VibeRaven

GitHub - ohad6k/VibeRaven: Open-source mission control for AI-built apps. The operating layer for apps built by AI agents — providers, releases, production evidence, and version drift from first demo to real scale. · GitHub
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
ohad6k
/
VibeRaven
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
79 Commits 79 Commits .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor/ rules .cursor/ rules .github .github .viberaven .viberaven agent-skills agent-skills assets assets commands commands docs docs examples examples skills/ viberaven skills/ viberaven .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md EXPORT_MANIFEST.json EXPORT_MANIFEST.json GEMINI.md GEMINI.md LICENSE LICENSE README.md README.md after-install.md after-install.md agent-context.md agent-context.md gemini-extension.json gemini-extension.json llms-full.txt llms-full.txt llms.txt llms.txt plugin.yaml plugin.yaml skills.sh.json skills.sh.json View all files Repository files navigation
Ship AI-built apps with evidence, not vibes.
Open the full-quality MP4 demo
VibeRaven is the VibeRaven Production Skills plus a local Studio cockpit for AI-built apps. Use it while building, fixing, reviewing releases, or preparing a launch when an agent needs provider context, release context, approval-aware chat, safe repo changes, and evidence before it claims the work is ready.
npx -y viberaven
The public repo is the agent discovery and installation surface: plugin-style metadata, portable slash commands, production skills, MCP notes, and AI-readable docs.
AI coding agents are good at patching code and bad at knowing what they cannot prove. VibeRaven Production Skills are not passive reports; they are task workflows an agent can use while implementing, reviewing, debugging, or preparing launch-sensitive work.
VibeRaven gives agents a contract:
separate repo-code fixes from provider dashboard work;
use MCP or connected provider context when available;
keep release and version context visible;
ask for approval before risky local changes.
That matters during feature work, auth fixes, billing changes, deploy debugging, release comparison, post-launch drift, and production readiness.
Open the local cockpit from any project:
npx -y viberaven
Studio gives you agentic chat, draggable providers, draggable versions/releases, provider MCP context, terminal output, diff context, access-mode control, and CLI-agent connection checks. Installed is not the same as connected; test the agent connection before real chat control.
Install bounded VibeRaven guidance for Codex, Claude Code, Cursor, Copilot, Gemini, and related agents:
npx -y viberaven init --agents all
Preview without writing files:
npx -y viberaven init --agents all --dry-run
The installer writes rules into supported agent instruction files, including AGENTS.md , CLAUDE.md , GEMINI.md , Cursor rules, Copilot instructions, and .viberaven context files.
VibeRaven ships as a portable skill/plugin-style pack:
Codex: .codex-plugin/plugin.json
Claude Code: .claude-plugin/plugin.json
Gemini CLI: gemini-extension.json
Generic plugin hosts: plugin.yaml
Slash-command prompts: commands/
/viberaven-help : show the pack and output contract.
/viberaven-work : use Production Skills on a task.
/viberaven-launch : run launch readiness before saying ready.
/viberaven-provider-actions : separate repo-code fixes from dashboard work.
See docs/agent-portability.md .
Browse the skill library in docs/production-skills.md .
npx -y skills add ohad6k/VibeRaven --skill viberaven
npx -y skills add ohad6k/VibeRaven --skill what-broke
npx -y skills add ohad6k/VibeRaven --skill go-live
MCP
VibeRaven can be exposed to MCP-aware agents:
{ "viberaven" : { "command" : " npx " , "args" : [ " -y " , " viberaven " , " --mcp " ] } }
Use MCP output as provider and readiness context for Studio-aware agents. The repo also keeps MCP registry metadata in docs/mcp-registry-submission.md so maintainers can verify the public MCP discovery surface alongside npm and skills.sh metadata.
examples/nextjs-supabase-vercel-production-ready-template includes agent rules and viberaven:* scripts for Next.js + Supabase + Vercel.
Older VibeRaven docs and compatibility tools may mention agent-mode scans, task lists, gate results, PRP resources, or scan-derived action manifests. Treat that language as compatibility context, not the main public product surface.
The current default is the Studio:
npx -y viberaven
Normal git push is not gated. VibeRaven language about readiness is about production-change confidence, release review, provider-aware evidence, and agent boundaries, not blocking ordinary repository work.
Issues: ohad6k/VibeRaven/issues
Public discovery repo: ohad6k/VibeRaven
Open-source mission control for AI-built apps. The operating layer for apps built by AI agents — providers, releases, production evidence, and version drift from first demo to real scale.
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
2
VibeRaven Studio demo (26s)
Latest
Jun 26, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
