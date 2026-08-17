---
source: "https://github.com/divadbaroon/claude-plugins"
hn_url: "https://news.ycombinator.com/item?id=49337325"
title: "Show HN: Engelbart – Mange Goals and TODOs for Claude Code"
article_title: "GitHub - divadbaroon/claude-plugins · GitHub"
image: "https://opengraph.githubassets.com/738b7cf785db8877c32002961b9addd0558890edbbf42df464db2bd982752487/divadbaroon/claude-plugins"
author: "hudsonmp"
captured_at: "2026-08-17T21:17:37Z"
capture_tool: "hn-digest"
hn_id: 49337325
score: 1
comments: 0
posted_at: "2026-08-17T20:41:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Engelbart – Mange Goals and TODOs for Claude Code

- HN: [49337325](https://news.ycombinator.com/item?id=49337325)
- Source: [github.com](https://github.com/divadbaroon/claude-plugins)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T20:41:03Z

## Translation

タイトル: HN を表示: Engelbart – クロード コードの目標と TODO を管理する
記事タイトル: GitHub - divadbaroon/claude-plugins · GitHub
説明: GitHub でアカウントを作成して、divadbaroon/claude-plugins の開発に貢献します。
HN テキスト: プロジェクトがより複雑になり、クロード コードのチャットがコンパクトになるにつれて、私が構築しようとしているものや今後の TODO に関するコンテキストが失われたり、混同されたりします。私たちは、この問題に対処するために Engelbart を構築しました。オープンソースの Claude プラグインは、以前の会話ターンから目標と TODO を推測し、Web インターフェイスを介してそれらを提示し、クロード コードの会話ターンにコンテキストを挿入し、プロジェクトの目標を更新します。

記事本文:
GitHub - divadbaroon/claude-plugins · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ディバドバルーン
/
クロードプラグイン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
354 コミット 354 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github .github コンパクトフォーカス コンパクトフォーカス engelbart engelbart hc hc テスト テスト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md GOAL_SYSTEM_LAUNCH_AUDIT.md GOAL_SY

STEM_LAUNCH_AUDIT.md LAUNCH_FEATURES.md LAUNCH_FEATURES.md LICENSE LICENSE README.md README.md STASHED.md STASHED.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェントを操作するためのツール。
デモ-small.mp4
目標と TODO 用のオープンソースの Claude コード プラグイン
macOS または Linux、ノード 18 以降、クロード コード 2.1.175 以降。
npx エンゲルバート-cli
Claude コード (または /reload-plugins ) を再起動します。
/goals-ui
このチャットの目標ワークスペースを開きます。クロードは何も言わない。
/goals-ui を無効にする
このチャットの分析と挿入を停止します。
ワークスペース — 目標ツリー、目標ごとに 1 つのマークダウン ドキュメント、リンクされたプロンプト、組み立てられたプロンプト。チャットごとに、ローカルポート上で。
インジェクション — 最初の /goals-ui の後、目標ドキュメントはチャットに戻ります。セッション開始時と圧縮後のファイル全体、その後の差分です。サブエージェントとツール バッチもそれを読み取ります。
永続性 — 1 回の呼び出しはチャットが存続する間保持されます。
コーディング エージェントと作業する場合、通常、目標と意図は暗黙的に示されます。これらは、プロンプト、TODO、実装の詳細、そしてあなた自身の頭の中に存在します。これは、保護のサイズが大きくなるにつれて、情報が失われ、混乱することを意味します。
\autocompact、projects、claude-men などの既存のツールは、自律的なコンテキストの保存を通じてこの問題の一部を解決しようとしますが、これらのプロセスでは依然として問題に関する重要な情報が失われ、混同されます。さらに悪いことに、これらのツールには人間の介入がないということは、エージェントが達成しようとしていると考えられることを人間が検査したり操作したりする能力を人間に与えることができないことを意味します。
これが、コーディング エージェント間で目標と TODO を管理、計画、同期するための無料のオープンソース ツールである Engelbart を作成した理由です。
Engelbart はブラウザベースの Claude Code プラグインで、エージェントが作業中に達成しようとしていることを共有して表現することをあなたとエージェントに提供します。

リアルタイムで変更を実装します。
Engelbart をインストールした後、クロード コードで /goals-ui を実行してローカル サーバーを起動できます。次に、Engelbart は現在のセッションと過去の会話ターンを分析して目標、計画、TODO を推測します。これを使用して、ローカル サーバー上に提案された目標ツリーが開きます。構築を再開する前に、その目標ツリーを検査して修正できます。
作業中、Engelbart は、新機能の計画、プロンプトの下書き、TODO の作成、目標の変更、現在のシステムに関するメモの書き留め、重要な決定の記録を行う際に、エージェントの最新情報を常に把握します。
Engelbart は、意図をコンテキスト ウィンドウ内に埋め込んだり、頭の中で暗黙のままにするのではなく、明示的で永続的で操作可能なものにするための重要な第一歩であると私たちは感じています。
Engelbart は初期ベータ段階にあり、まだ開発の初期段階にあります。これは、人間と AI システムが長期にわたる作業においてどのように計画を立て、目標を維持し、調整するかについての広範な認知科学研究の一部でもあります。
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to divadbaroon/claude-plugins development by creating an account on GitHub.

Context about what I'm trying to build and upcoming TODOs get lost or mixed up as my projects grow more complex and my Claude Code chats compact. We built Engelbart to address this problem, an open-source Claude plugin that infers your goals and TODOs from previous conversation turns and presents them via a web interface that injects context into your Claude Code conversation turns and you update the your project goals.

GitHub - divadbaroon/claude-plugins · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
divadbaroon
/
claude-plugins
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
354 Commits 354 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github .github compact-focus compact-focus engelbart engelbart hc hc tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md GOAL_SYSTEM_LAUNCH_AUDIT.md GOAL_SYSTEM_LAUNCH_AUDIT.md LAUNCH_FEATURES.md LAUNCH_FEATURES.md LICENSE LICENSE README.md README.md STASHED.md STASHED.md install.sh install.sh View all files Repository files navigation
Tools for steering coding agents.
demo-small.mp4
Open-Source Claude Code Plugin for Goals and TODOs
macOS or Linux, Node 18+, Claude Code 2.1.175+.
npx engelbart-cli
Restart Claude Code (or /reload-plugins ).
/goals-ui
Opens this chat's goal workspace; Claude says nothing.
/goals-ui disable
Stops analysis and injection for this chat.
Workspace — goal tree, one markdown document per goal, linked prompts, assembled prompt. Per chat, on a local port.
Injection — after the first /goals-ui , the goals document goes back into the chat: whole file on session start and after compaction, a diff afterwards. Subagents and tool batches read it too.
Persistence — one invocation holds for the life of the chat.
Goals and intent are usually implicit when you work with coding agents. They live across prompts, TODOs, implementation details, and your own head. This means that information is lost and confounded as protects grow in size.
Existing tools like \autocompact, projects, and claude-men try to solve parts of this problem through autonomous context preservation, but these processes still lose and conflate important information about the problem. Worse, the lack of human intervention in these tools means they fail to give humans the ability to inspect or steer what the agent thinks it is trying to accomplish.
That’s why we created Engelbart, a free, open-source tool for managing, planning, and syncing goals and TODOs across your coding agents.
Engelbart is a browser-based Claude Code plugin that gives you and your agent a shared representation of what you’re trying to accomplish while the agent implements changes in real time.
After installing Engelbart, you can run /goals-ui in Claude Code to kick off a local server. Engelbart then analyzes your current session and past conversation turns to infer your goals, plans, and TODOs, which it uses to open a proposed goal tree on a local server that you can inspect and correct before you resume building.
As you work, Engelbart keeps your agent in the loop as you plan new features, draft prompts, write TODOs, modify goals, jot down notes about the current system, and record key decisions.
We feel Engelbart is an important first step in making intent explicit, persistent, and steerable instead of leaving it buried inside a context window or tacit inside your head.
Engelbart is in early beta and still in the initial stages of development. It’s also part of our broader cognitive science research into how humans and AI systems plan, maintain goals, and coordinate over long-running work.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
