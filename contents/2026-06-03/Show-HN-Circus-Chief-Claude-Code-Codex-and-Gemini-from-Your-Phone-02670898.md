---
source: "https://github.com/ferrislucas/Circus-Chief"
hn_url: "https://news.ycombinator.com/item?id=48370360"
title: "Show HN: Circus Chief – Claude Code, Codex, and Gemini from Your Phone"
article_title: "GitHub - ferrislucas/Circus-Chief: Open-source control plane for managing Claude Code, OpenAI Codex, and Gemini CLI coding agents from a browser. · GitHub"
author: "deathmonger5000"
captured_at: "2026-06-03T00:44:41Z"
capture_tool: "hn-digest"
hn_id: 48370360
score: 3
comments: 0
posted_at: "2026-06-02T13:59:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Circus Chief – Claude Code, Codex, and Gemini from Your Phone

- HN: [48370360](https://news.ycombinator.com/item?id=48370360)
- Source: [github.com](https://github.com/ferrislucas/Circus-Chief)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T13:59:56Z

## Translation

タイトル: HN を表示: サーカス長 – 携帯電話からクロード コード、コーデックス、ジェミニを表示
記事のタイトル: GitHub - ferrislucas/Circus-Chief: Claude Code、OpenAI Codex、および Gemini CLI コーディング エージェントをブラウザから管理するためのオープンソース コントロール プレーン。 · GitHub
説明: Claude Code、OpenAI Codex、および Gemini CLI コーディング エージェントをブラウザから管理するためのオープンソース コントロール プレーン。 - フェリスルーカス/サーカス団長
HN テキスト: こんにちは、HN、サーカス チーフは、ブラウザからコーディング エージェント セッションを管理するためのツールです。特に小さな画面向けに最適化されています。 Claude Code、OpenAI Codex、および Google Gemini CLI エージェントをサポートします。特徴 エージェントはサーカスチーフそのものを操作することができます。エージェントは、セッションの生成、セッションのスケジュール設定、カンバン ボードとの対話を行うことができます。UI で実行できることはすべて、エージェントも行うことができます。事前に作業のスケジュールを立てます。使用制限に達すると、自動的にスケジュールが変更されます。構成可能でチェーン可能なプロンプト テンプレート。ユーザー定義コマンド - エージェントはコマンドを実行して結果を確認でき、UI でも同様に実行できます。ローカル CI または日常的なワークフロー ステップに便利です。ワークツリーをセッションごとに分離するか、特定のブランチで作業することを選択します。共有キャンバス - コードに属さない共有成果物のための場所です。計画文書を反復処理する場合に便利です。独自のプロバイダーを持ち込んでください。 Anthropic、OpenAI、または Google のサブスクリプション認証を使用するか、Anthropic または OpenAI 互換のエンドポイントを備えたサードパーティ プロバイダーでセッションをポイントします。試してみたい場合: npx circuschief
スタック: Vue 3、Express、SQLite、WebSocket。プレーンな JavaScript。追い風。開発プロセス 私はこれをほぼ完全に携帯電話とタブレットから構築しました。まず Anthropic と z.ai の両方を使用して Claude Code から始め、約 1 か月前に Codex をミックスに追加しました。タラの多くの個々の行をレビューしているわけではないという意味で、それは雰囲気コード化されていました

e.私はエージェントと計画を確認して繰り返しましたが、実装の詳細については触れないことが多かったです。計画に意味がある場合は、通常、エージェントに何回もレビューしてもらってから、何も介入せずに実行してもらいます。私はこの方法で構築するのを楽しんできましたが、実装の詳細についてはあまり深く理解していません。これまでのところ、私はちょうど 1 回そのことに悩まされました。この問題を参照してください: https://github.com/ferrislucas/Circus-Chief/issues/944 私は、非常に早い段階から Circus Leader を構築するために Circus Leader を使用しました。私は Vibe Kanban を使用してブートストラップし、サーカス チーフが使用できるようになったらドッグフーディングに切り替えました。アプリの多くはアプリ自体を通じて構築されています。数字によれば、2025 年のクリスマス以降、エージェント セッションは 2,184 回あり、そのうち 1,423 回は子供/フォローアップ セッションでした。エージェントの構成: クロード コード セッション 1,818 件、コーデックス セッション 363 件、ジェミニ セッション 3 件。これらのセッション全体で 50,929 件の会話メッセージ。 41 セッション (1.9%) には、ユーザー作成のメッセージに強い冒涜的な表現が含まれていました。 685 の個別の PR、759 のブランチ、731 のワークツリーにわたる、GitHub PR にリンクされた 990 のセッション。記録されたエージェント通話の平均: 131.3 秒。メインで 2,410 コミット。合計 883 件のプル リクエスト — 847 件がマージされました (約 96%)。リポジトリの規模: 887 の追跡されたファイルと 310,000 の現在のコード行 (ロックファイル、フィクスチャ、および記録されたテスト カセットを除く)。ピーク日: 2026 年 3 月 21 日の 76 件のコミット、最初のコミットは午前 12 時 49 分、最後は午後 11 時 57 分で、23 時間にわたって 18 分ごとに 1 つです。マージが最も多かった日: 28 件のプル リクエストが送信されました — 2026 年の元旦。差分による最大の PR: 92 ファイルにわたる 140K の挿入 - モック クロード フィクスチャを VCR スタイルの記録/再生に置き換えます。コード: https://github.com/ferrislucas/Circus-Chief Analytics の開示: npm ビルドには、匿名のページビュー分析用の PostHog が含まれています (セッション記録オフ、DNT 尊重)。 - で無効にする

-no-analytics は、[設定] → [全般] で選択するか、PostHog キーを使用せずにソースからビルドします。

記事本文:
GitHub - ferrislucas/Circus-Chief: ブラウザから Claude Code、OpenAI Codex、および Gemini CLI コーディング エージェントを管理するためのオープンソース コントロール プレーン。 · GitHub
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
フェリスルーカス
/
サーカス団長
公共
通知
通知を変更するにはサインインする必要があります

アイケーション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2,410 コミット 2,410 コミット .claude/ エージェント .claude/ エージェント .github/ ワークフロー .github/ ワークフロー docker/ playwright docker/ playwright docs docs パッケージ パッケージ スクリプト スクリプト サーバー/ テスト サーバー/ テスト テスト テスト .env.production.example .env.production.example .eslintrc-pr4-dryrun.cjs .eslintrc-pr4-dryrun.cjs .eslintrc.cjs .eslintrc.cjs .gitignore .gitignore .prettierrc .prettierrc .yarnrc .yarnrc CLAUDE.md CLAUDE.md ライセンス ライセンス通知 NOTICE PUBLISHING.md PUBLISHING.md README.md README.md TRADEMARKS.md TRADEMARKS.md docker-compose.playwright.yml docker-compose.playwright.yml eslint-tightening-plan.md eslint-tightening-plan.md package.json package.json playwright.config.docker.ts playwright.config.docker.ts playwright.config.ts playwright.config.ts糸.ロック 糸.ロック すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントを管理するための、タッチ操作に最適化されたオープンソースのコントロール プレーン。
Claude Code エージェント (任意の Anthropic 互換プロバイダー)、OpenAI Codex エージェント (任意の OpenAI 互換プロバイダー)、および Google Gemini CLI エージェントをサポートします。
API キーまたはサブスクリプションベースの認証と連携します。
エージェントはサーカスチーフ自体を操作することができます。各エージェントは、セッションの検査、フォローアップの生成、再試行のスケジュール設定、作業の停止/再開、および結果への対応を行うことができます。
事前に作業のスケジュールを立てます。 1 回限りのプロンプトから完全なテンプレート パイプラインまで、後でエージェント セッションを計画します。
使用制限のオプトイン再試行。セッション中にこれをオンにすると、トークンの上限やプロバイダーの停止に達した場合、自動的にスケジュールが変更され、中断したところから再開されます。
構成可能でチェーン可能なテンプレート。各テンプレートはプロンプトとセッション設定を定義し、1 つのテンプレートで次のテンプレートを自動起動できます。エクサ

複数のパイプライン: 計画 → 計画のレビュー → 計画の実装 → 実装のレビュー → PR を開く — テンプレートはそれ自体を呼び出すことができます。
すべてのセッションについて AI が生成した概要なので、トランスクリプト全体を読み直すことなく、各エージェントが何をしているのか、どこから中断したかを確認できます。これはプロジェクト設定でオフにできます。
ユーザー設定のコマンド。テスト、lint、ビルド、タイプチェック、CI チェックなど、定期的に実行するプロジェクト コマンド用のワンタップ ボタンを追加します。出力ストリームはライブであり、合否結果はオプションでダッシュボードに表示できます。
クロード コード、コーデックス、ジェミニ セッション。同じモバイル コントロール、履歴、キャンバス、コマンド、およびワークツリー分離を使用して、同じダッシュボードからあらゆる種類のエージェントを開始します。エージェントやプロバイダーを自由に切り替えられます。同じワークツリーまたは独自のワークツリーに対して並列エージェントを呼び出します。
セッションごとのワークツリーの分離。すべてのセッションは独自の git ワークツリーを取得します。メイン git リポジトリで作業するか、メイン git リポジトリの特定のブランチで作業するかを選択することもできます。
共有キャンバス。マークダウン、画像、JSON、コード - エージェントとあなたは同じ成果物を編集します。バージョン履歴が含まれます。
セッションごとに独自のプロバイダーを導入します。 Anthropic、OpenAI、または Google のサブスクリプション認証を使用するか、Anthropic または OpenAI 互換のエンドポイントを備えたサードパーティ プロバイダーでセッションをポイントします。 Claude Code、Codex、および Gemini CLI はすべてファーストクラスのパスです。
ライブ CI とマージ/競合状態を備えた自動リンクされた GitHub PR ( gh が必要)。
旗
説明
-p, --port <番号>
リッスンするポート (デフォルト: 5000)
--分析なし
匿名の使用状況分析を無効にする
-h、--ヘルプ
ヘルプメッセージを表示する
-v、--バージョン
バージョン番号を表示
例 — カスタム ポートで実行します。
npx サーカス団長 -p 8080
前提条件
クロード コード — クロード コード エージェントに必須
OpenAI Codex CLI — Codex エージェントに必要
Google ジェミニ CL

I — Gemini エージェントに必須
GitHub CLI (オプション - 自動 PR リンクを有効にします)
開発ガイド — クイック スタート、コマンド、テスト、環境変数
ビルドと配布 — npm パッケージがどのようにビルド、公開、実行されるか
エージェント システム プロンプトと REST API リファレンス — システム プロンプトを介してエージェントに公開される REST API
CircusChief は、Apache License 2.0 に基づいてライセンスを取得しています。
その保証と免責事項。帰属については「通知」を参照してください
および商標表示、およびプロジェクトの TRADEMARKS.md
サーカス チーフ、サーカス タイム、サーカス サーチを対象とする商標ポリシー。
Claude Code、OpenAI Codex、および Gemini CLI コーディング エージェントをブラウザから管理するためのオープンソース コントロール プレーン。
www.npmjs.com/package/circuschief
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
3
1.7.0
最新の
2026 年 5 月 31 日
+ 2 リリース
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source control plane for managing Claude Code, OpenAI Codex, and Gemini CLI coding agents from a browser. - ferrislucas/Circus-Chief

Hi HN, Circus Chief is a tool for managing coding agent sessions from a browser. It's specifically optimized for small screens. It supports Claude Code, OpenAI Codex, and Google Gemini CLI agents. Features Agents can operate Circus Chief itself. Agents can spawn sessions, schedule sessions, interact with the Kanban board — anything you can do in the UI, an agent can also do. Schedule work ahead of time. Automatically reschedule when you hit usage limits. Configurable, chainable prompt templates. User-defined commands - agents can run them and see the results, and so can you in the UI. Handy for local CI or routine workflow steps. Worktree-per-session isolation, or elect to work on a specific branch. Shared canvas - this is a place for shared artifacts that don't belong in the code. Useful for iterating on planning documents. Bring your own provider. Use subscription auth for Anthropic, OpenAI, or Google, or point sessions at third-party providers with Anthropic- or OpenAI-compatible endpoints. If you want to try it: npx circuschief
Stack: Vue 3, Express, SQLite, WebSockets. Plain JavaScript. Tailwind. Development process I built this almost entirely from my phone and tablet — starting with Claude Code using both Anthropic and z.ai, then adding Codex to the mix around a month ago. It was vibe coded in the sense that I was not reviewing many individual lines of code. I did review and iterate on plans with the agents, but often stayed out of implementation details. If a plan made sense then I would usually have the agent review it a few times before just letting the agent implement without any intervention. I've enjoyed building this way, but I don't deeply understand a lot of the details of the implementation. That's bitten me exactly once so far, see this issue: https://github.com/ferrislucas/Circus-Chief/issues/944 I used Circus Chief to build Circus Chief from a very early stage. I bootstrapped with Vibe Kanban, then switched to dogfooding once Circus Chief was usable. A lot of the app has been built through the app itself. By the numbers 2,184 agent sessions since Christmas 2025 — 1,423 were child/follow-up sessions. Agent mix: 1,818 Claude Code sessions, 363 Codex sessions, and 3 Gemini sessions. 50,929 conversation messages across those sessions. 41 sessions (1.9%) contained strong profanity in user-authored messages. 990 sessions linked to GitHub PRs, across 685 distinct PRs, 759 branches, and 731 worktrees. Average logged agent call: 131.3 seconds. 2,410 commits on main; 883 pull requests total — 847 merged (~96%). Repo scale: 887 tracked files and 310K current lines of code (excludes lockfiles, fixtures, and recorded test cassettes). Peak day: 76 commits on March 21, 2026, first at 12:49am, last at 11:57pm — one every 18 minutes for 23 hours. Busiest merge day: 28 pull requests shipped — New Year's Day, 2026 . Largest PR by diff: 140K insertions across 92 files — replacing mock Claude fixtures with VCR-style record/replay. Code: https://github.com/ferrislucas/Circus-Chief Analytics disclosure: the npm build includes PostHog for anonymous page-view analytics (session recording off, DNT respected). Disable with --no-analytics, in Settings → General, or build from source without a PostHog key.

GitHub - ferrislucas/Circus-Chief: Open-source control plane for managing Claude Code, OpenAI Codex, and Gemini CLI coding agents from a browser. · GitHub
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
ferrislucas
/
Circus-Chief
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2,410 Commits 2,410 Commits .claude/ agents .claude/ agents .github/ workflows .github/ workflows docker/ playwright docker/ playwright docs docs packages packages scripts scripts server/ test server/ test tests tests .env.production.example .env.production.example .eslintrc-pr4-dryrun.cjs .eslintrc-pr4-dryrun.cjs .eslintrc.cjs .eslintrc.cjs .gitignore .gitignore .prettierrc .prettierrc .yarnrc .yarnrc CLAUDE.md CLAUDE.md LICENSE LICENSE NOTICE NOTICE PUBLISHING.md PUBLISHING.md README.md README.md TRADEMARKS.md TRADEMARKS.md docker-compose.playwright.yml docker-compose.playwright.yml eslint-tightening-plan.md eslint-tightening-plan.md package.json package.json playwright.config.docker.ts playwright.config.docker.ts playwright.config.ts playwright.config.ts yarn.lock yarn.lock View all files Repository files navigation
An open-source, touch-optimized control plane for managing AI coding agents.
Supports Claude Code agents (any Anthropic-compatible provider), OpenAI Codex agents (any OpenAI-compatible provider), and Google Gemini CLI agents.
Works with API key or subscription-based authorization.
Agents can operate Circus Chief itself. Each agent can inspect sessions, spawn follow-ups, schedule retries, stop/restart work, and react to results.
Schedule work ahead of time. Plan agent sessions for a later time, from one-off prompts to full template pipelines.
Opt-in retry on usage limits. Toggle it on for a session and, if it hits a token cap or provider outage, it reschedules itself and picks up where it left off.
Configurable, chainable templates. Each template defines a prompt and session settings, and one template can auto-launch the next. Example pipeline: plan → review plan → implement the plan → review implementation → open PR — templates can invoke themselves.
AI-generated summaries on every session, so you can see what each agent is doing and where you left off without re-reading the whole transcript. You can turn this off in project settings.
User-configured commands. Add one-tap buttons for the project commands you run constantly: tests, lint, build, typecheck, CI checks. Output streams live, and pass/fail results can optionally display on the dashboard.
Claude Code, Codex, and Gemini sessions. Start any kind of agent from the same dashboard, with the same mobile controls, history, canvas, commands, and worktree isolation. Switch agents and/or providers freely. Invoke parallel agents against the same worktree or in their own work trees.
Worktree-per-session isolation. Every session gets its own git worktree. You can also elect to work in the main git repo, or on a specific branch of the main git repo.
Shared canvas. Markdown, images, JSON, code — agents and you edit the same artifacts. Version history included.
Bring your own provider — per session. Use subscription auth for Anthropic, OpenAI, or Google, or point sessions at third-party providers with Anthropic- or OpenAI-compatible endpoints. Claude Code, Codex, and Gemini CLI are all first-class paths.
Auto-linked GitHub PRs with live CI and merge/conflict state (needs gh ).
Flag
Description
-p, --port <number>
Port to listen on (default: 5000 )
--no-analytics
Disable anonymous usage analytics
-h, --help
Show help message
-v, --version
Show version number
Example — run on a custom port:
npx circuschief -p 8080
Prerequisites
Claude Code — required for Claude Code agents
OpenAI Codex CLI — required for Codex agents
Google Gemini CLI — required for Gemini agents
GitHub CLI (optional — enables automatic PR linking)
Development Guide — Quick start, commands, testing, environment variables
Build & Distribution — How the npm package is built, published, and run
Agent System Prompt & REST API Reference — The REST API exposed to agents via the system prompt
Circus Chief is licensed under the Apache License 2.0 , including
its warranty and liability disclaimers. See NOTICE for attribution
and trademark notices, and TRADEMARKS.md for the project
trademark policy covering Circus Chief, Circus Time, and Circus Search.
Open-source control plane for managing Claude Code, OpenAI Codex, and Gemini CLI coding agents from a browser.
www.npmjs.com/package/circuschief
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
3
1.7.0
Latest
May 31, 2026
+ 2 releases
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
