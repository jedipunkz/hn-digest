---
source: "https://github.com/gtapps/claude-code-hermit/"
hn_url: "https://news.ycombinator.com/item?id=49394794"
title: "Turn Claude Code into a 24/7 Agent"
article_title: "GitHub - gtapps/claude-code-hermit: Turn Claude Code into an Always-on Agent · GitHub"
image: "https://opengraph.githubassets.com/515999bc257195323a02eb1238153fe239de2379a7a0c101e591a649af4b4068/gtapps/claude-code-hermit"
author: "gtapps"
captured_at: "2026-08-21T23:14:28Z"
capture_tool: "hn-digest"
hn_id: 49394794
score: 1
comments: 1
posted_at: "2026-08-21T23:03:11Z"
tags:
  - hacker-news
  - translated
---

# Turn Claude Code into a 24/7 Agent

- HN: [49394794](https://news.ycombinator.com/item?id=49394794)
- Source: [github.com](https://github.com/gtapps/claude-code-hermit/)
- Score: 1
- Comments: 1
- Posted: 2026-08-21T23:03:11Z

## Translation

タイトル: クロード・コードを年中無休のエージェントに変える
記事のタイトル: GitHub - gtapps/claude-code-hermit: クロード コードを常時稼働エージェントに変える · GitHub
説明: クロード コードを常時稼働エージェントに変えます。 GitHub でアカウントを作成して、gtapps/claude-code-hermit の開発に貢献してください。

記事本文:
GitHub - gtapps/claude-code-hermit: クロード コードを常時稼働エージェントに変える · GitHub
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
gtapps
/
クロード・コード・ハーミット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
2,182 コミット 2,182 コミット フォルダーとファイル
.agents/ スキル .agents/ スキル .claude-plugin .claude-plugin .claude .claude .github/ workflows .github/ workflows プラグイン プラグイン スクリプト スクリプト テスト テスト .dockerignore .do

ckerignore .gitattributes .gitattributes .gitignore .gitignore .worktreeinclude .worktreeinclude CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code インスタンスを 24 時間年中無休のエージェントに変える Claude Code プラグイン。ステートフル。積極的。オペレーターゲート型提案システムによる自己改善。コスト意識が高い。観察可能。クロード サブスクリプションと連携します。
/hatch を使用して任意のフォルダー、空のプロジェクト、または既存のプロジェクトにエージェントをセットアップし、そのアイデンティティ、優先順位、ルーチン、知識、自律性、ガードレールを形成して、それを自分のものにします。
# インストール
クロード プラグイン マーケットプレイス追加 gtapps/claude-code-hermit
クロード プラグインのインストール claude-code-hermit@claude-code-hermit --scope local
# Claude Code を起動し、セットアップ ウィザードを実行します
/claude-code-hermit:ハッチ
# 常時オンにする
/claude-code-hermit:docker-setup
追加されるもの
Hermit は、Claude コードの周囲に永続的なオペレーティング層、学習ループ、およびすべてを接続するための迅速なセットアップを追加します。
ステートフルなライブ作業状態、アーカイブされたセッションのハンドオフ、実行時の観察、レッスン、調査結果、ブロッカー、完了したタスク、作成/変更/削除されたファイル。
エージェント ルーチン セッション外で適格性を決定する 1 つの永続的なモニター サブプロセスから実行される独自のルーチンを追加します。これにより、スキップされたファイアのコストはゼロになり、ルーチンのバッチが 1 つのウェイクにまとめられます。毎日の CronCreate アンカーによって再装備されます。 Monitor が使用できない場合は、ルーチンごとの CronCreate にフォールバックします。 /hermit-routines によって管理されます。
ハートビートは永続的な Monitor サブプロセスからポーリングします。ファイルシステムのみの事前チェックによってすべてのティックが決定され、モデルは実際に何かが変更されたときにのみ起動します (請求のみが行われます)。
/watch は、セッションとともに停止したストリームを監視します。

設定 (または平易な言語) から自動的に開始され、検出結果が通知にルーティングされます。静かなときは沈黙します。
チャンネルを使用すると、セッションを DM できます。 Hermit Agent はそれに基づいて動作し ( "accept PROP-014" 、 "status" )、何かが Yes/No を必要とするときに最初に ping を送信します。
携帯電話から一時停止すると、実際に停止します。 Discord または Telegram でステータスを尋ねたり、一時停止、再開、またはスヌーズしたりできます。一時停止は、単なる会話要求として扱われるのではなく、ツール境界で強制されます。
ネイティブ クロード コード アーティファクトの統合により、ライブ Hermit ダッシュボード、オープンな提案、週次レビュー、およびリクエストされたコンパイル済みドキュメントがプライベートのバージョン管理されたクロード コード アーティファクトとして公開されます。ページは安定した URL で適切に更新され、サポートされている場合は組織共有が行われます。代わりに、artifacts.backend を独自の MCP アーティファクト サーバーに指定して、そこで公開します。
自動記憶 + 知識 2 層。 Claude Code のネイティブ自動メモリには、オペレータの事実と好み (どのように作業するか) が保存されます。その上に、隠者は生の/→コンパイルされた/ナレッジ ベースを追加します。つまり、ドメイン出力と生きたトピック ページが適切に更新され、新規および再開された開始時にコンテキスト バジェット内でカタログとして再挿入されます。 Discord/Telegram DM テキストもローカルにキャプチャされるため、チャットで行われた決定はスレッドよりも長く残ります。週次レビューによって決定がメモリに蒸留されます (knowledge.channel_log_enabled: false でオプトアウト)。 /recall はすべてを検索します。
計画の追跡は、SHELL.md 進行状況ログ (圧縮、再起動、およびすべてのモデル層を経ても存続するタイムスタンプ付きのステップ) に保存されます。
無人安全性は、プロファイル ゲート拒否パターン + サンドボックス、チャネル ルーティング アスク、許可拒否アラート、ハートビートと起動コンテキストのインジェクション スキャンを組み合わせたものです。
オーケストレーターはタスクと探索を他のエージェントに委任するように指示され、トークン効率のためにメイン コンテキストはクリーンなままになります。
セス

自己管理します。デーモンはアイドル状態の 12 時間と、外出中の真夜中に自動アーカイブされるため、手動で終了しなくても証拠が学習ループに到達します。外部ウォッチドッグは、デッドセッションを再開し、ウェッジされたセッションをナッジし、逃したスケジュールを再設定し、深夜終了後に古いコンテキストをクリアし、コールドウェイクによって蓄積された履歴全体が再支払われないように長時間実行コンテキストを圧縮します。リカバリは、セッションが意識的であるかどうかには決して依存しません。
コンテキスト効率の高い継続性。圧縮後、Hermit は完全なスタートアップ バンドルではなく、制限されたライフサイクル/タスク/進行状況カプセルのみをリロードします。構造化されたレポートのフロントマターにより、レポート本文をすべて読み直すことなく、概要、振り返り、および週次レビューで履歴を検査できます。
それがまずあなたに届きます。通知はデフォルトでネイティブプッシュ (ヘッドレスフレンドリー)、またはチャンネルをペアリングしている場合に返信できる Discord/Telegram DM になります。
コストは時間ではなくイベントに応じて変化します。何かが起こるまでモデルを目覚めさせることは何もないので、アイドル状態の隠者は実質的に自由になります。
隠者は、セッション全体で何がうまくいかないのかを観察し、修正を提案し、イエスかノーかを尋ねます。同じことを二度提案することはありません。
セッションの終了、アイドルティック、スケジュールされたリズムなど、自然な一時停止時にそれが反映されます。ほとんどのリフレクションはモデルに到達しません。事前チェック スクリプトは、フェーズ (計算、解像度チェック、コスト スパイク、ダイジェスト、新規) が実際に期限切れかどうかをゲートします。 1 つが該当する場合、候補者が届く前に 2 人のサブエージェントがその候補者を精査します。
リフレクション・ジャッジは、引用された証拠がセッション報告書に実際に存在することを確認するため、提案自体を認定することはできません。
proposal-triage は、オープンなプロポーザルに対して重複を排除し、 MEMORY.md と OPERATOR.md をクロスチェックし、3 つの条件バーを適用します。
サバイバーは、DM を含め、どこからでも行動できる提案として表示されます。
/claude-code-hermit:proposal-list # 何を参照してください

見つかりませんでした
/claude-code-hermit:proposal-act accept PROP-003 # or just reply "accept PROP-003"
提案する内容: 改善、ルーチン、新機能 (スキル、エージェント、ハートビート チェック)、ガードレール (確認した OPERATOR.md ガイダンス)、およびバグ。 When it catches itself repeating the same multi-step procedure across sessions, it drafts the skill and asks before installing.それ自体のスキルも向上します。セッション間で修正や再作業が行われ続けると、それがスキル向上提案に段階的に移行し、OK があればスキルを修正します ( skill-creator 経由)。 Accepted proposals can carry a measurable success signal and auto-resolve when met.あなたはあらゆる変化を受け入れる入り口です。 Raw session journals distill into compiled artifacts that reload next session — the raw/compiled pattern Karpathy described for his wiki-LLM.
ネイティブの Hermit ダッシュボード、提案ページ、および週次レビューは、安定した URL、または自己ホスト型アーティファクト サーバー (構成している場合) にあるクロード コード アーティファクトとして最新の状態に維持されます。
On-demand skills — pullable from the Claude app, your terminal, or a DM:
/hermit-dashboard-design — designs the Dashboard around what this hermit actually tracks; delete .claude-code-hermit/dashboard-render.ts to restore the built-in page
/recall — full-text search over past sessions, compiled knowledge, proposals, and your channel DM history ("what did I decide about X?")
/hermit-evolution — cost trend and behavior drift over weeks
/hermit-health — alerts, routines, channels, heartbeat state, plus fragile zones, stale proposals, and recent learnings
/hermit-doctor — proactive install diagnostic, from hook registration to heartbeat and routine-monitor liveness; the weekly check stays silent when green and alerts only on new problems
/コスト反映

— 構造コスト監査: どのトークン タイプとトリガー ソースが支出を促進するか
/brief — 現在のステータスと最近の作業の概要
前提条件: Claude Code v2.1.172+、Claude プラン (Pro、Max、Teams、または Enterprise)、および Bun 1.3+。 Linux、macOS、および Windows (WSL2 経由) — FAQ を参照してください。
cd /path/to/your/project # または任意のフォルダー (空のフォルダーも含む)
クロード プラグイン マーケットプレイス追加 gtapps/claude-code-hermit
クロード プラグインのインストール claude-code-hermit@claude-code-hermit --scope local
2.初期化
クロード /claude-code-hermit:hatch
ウィザードはエージェントの ID を設定し、フォルダーをスキャンし、 OPERATOR.md を生成し、クイック (4 つの質問) またはアドバンス (完全なウィザード) を提供します。
ただ試してみただけですか？ hatch の後、24 時間 365 日の自律性を持たないセッション、ルーチン、ハートビート、および学習ループに対して .claude-code-hermit/bin/hermit-start --no-tmux を実行します。 Discord または Telegram が必要な場合は、最初に /claude-code-hermit:channel-setup を実行します。
/claude-code-hermit:docker-setup
Docker スキャフォールディングを生成し、イメージを構築し、コンテナを起動し、認証とチャネルのペアリングを実行します。コンテナーには強化ベースライン ( cap_drop: ALL 、 no-new-privileges 、 pids_limit ) が同梱されています。より強力な分離が必要ですか?オプトイン LAN 封じ込め + DNS 許可リスト + リソース境界に対して /docker-security を実行します。
詳細なガイドについては、「Always-On セットアップ」を参照してください。 Docker なしで常時稼働したいですか? 「ベア tmux の常時オン操作」を参照してください。
クロード プラグインの更新 claude-code-hermit@claude-code-hermit --scope local
/claude-code-hermit:hermit-evolve
または、 .claude-code-hermit/bin/hermit-update (local/tmux) または .claude-code-hermit/bin/hermit-docker update (Docker) を実行します。これは、ピンを移動し、セッションをリロードし、hermit-evolve を実行する 1 つのコマンドです。
/hermit-settings 経由で調整します (または隠者に尋ねるだけでも)。利用可能な設定の一部:
Config Refe の完全なスキーマ

レンス
/hermit-settings ですべてライブ編集可能 (または Hermit に問い合わせるだけ) — 再起動は必要ありません。
モデル&オートモード。デフォルトは Sonnet — 無人セッションの推論とコストのバランスが取れています。自動モードは通常、サブスクリプション プランおよび API の使用状況に応じてすべてのユーザーが利用できます。サポートされているモデルとプロバイダーの構成は依然として異なる可能性があるため、Claude が現在の選択が利用できないと報告した場合は、サポートされているモデルまたは別のアクセス許可モードを選択してください。より重い推論を求める場合は、opus に切り替えてください。ルーチンごとのモデル: 「haiku」は、軽量で独立した作業に引き続き役立ちます。
心臓の鼓動。 heartbeat.every はアイドル スイープを設定します (デフォルトは 2h、1 時間はタイト、4 時間 + ウェイク数は少なくなります)。 active_hours はウィンドウ ( 08:00 – 23:00 ) の境界を設定します。 heartbeat.enabled: false は時間指定ウェイクを完全に停止します。チャネルとルーチンは引き続き起動します。
アイドル的な動作。 Discover (デフォルト) は、OPERATOR.md + コスト ログに対して優先度調整パスを追加します。 wait はパッシブです (タスク/チャネルのみ)。いずれにせよ、毎日の反映ルーチンは引き続き実行されます。待機は、学習ループではなく、スケジュール間の検出を沈黙させるだけです。
ルーチン。各ルーチンはオプションのモデルを使用します。分離されたサブエージェントで、コストを節約するために haiku で軽量のものを実行するか、より推論するために opus で重いものを実行します。モデルを省略してメインセッションコンテキスト内でインラインに保ちます

[切り捨てられた]

## Original Extract

Turn Claude Code into an Always-on Agent. Contribute to gtapps/claude-code-hermit development by creating an account on GitHub.

GitHub - gtapps/claude-code-hermit: Turn Claude Code into an Always-on Agent · GitHub
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
gtapps
/
claude-code-hermit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2,182 Commits 2,182 Commits Folders and files
.agents/ skills .agents/ skills .claude-plugin .claude-plugin .claude .claude .github/ workflows .github/ workflows plugins plugins scripts scripts tests tests .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .worktreeinclude .worktreeinclude CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Claude Code plugin that turns a Claude Code instance into a 24/7 agent. Stateful. Proactive. Self-improving through an operator-gated proposal system. Cost-aware. Observable. Works with your Claude Subscription .
Setup your agent in any folder, empty or existing project with /hatch and shape its identity, priorities, routines, knowledge, autonomy, guardrails and make it yours.
# Install
claude plugin marketplace add gtapps/claude-code-hermit
claude plugin install claude-code-hermit@claude-code-hermit --scope local
# Boot Claude Code and run the setup wizard
/claude-code-hermit:hatch
# Go always-on
/claude-code-hermit:docker-setup
What it adds
Hermit adds a persistent operating layer around Claude Code, a learning loop, and a quick setup to wire everything.
Stateful live working state, archived session handoffs, runtime observations, lessons, findings, blockers, completed tasks, files created/modified/deleted.
Agent Routines Add your own routines that run from one persistent Monitor subprocess that decides eligibility outside the session, so a skipped fire costs zero tokens and co-due routines batch into one wake; a daily CronCreate anchor re-arms it. Falls back to per-routine CronCreate where Monitor is unavailable. Managed by /hermit-routines .
Heartbeat polls from a persistent Monitor subprocess — a filesystem-only precheck decides every tick, and the model only wakes (and only bills) when something actually changed.
/watch wraps Monitor streams that die with the session: it auto-starts from config (or plain language) and routes findings to your notifications, silent when quiet.
Channels let you DM a session; the hermit agent acts on it ( "accept PROP-014" , "status" ) and pings you first when something needs a yes/no.
Pause it from your phone — and it actually stops. Ask for status, pause, resume, or snooze over Discord or Telegram. The pause is enforced at the tool boundary, not merely treated as a conversational request.
Native Claude Code Artifacts integration publishes a live Hermit Dashboard, open proposals, weekly reviews, and any compiled document you request as private, versioned Claude Code Artifacts . Pages update in place at stable URLs, with organization sharing where supported. Point artifacts.backend at your own MCP artifact server to publish there instead.
Auto-memory + knowledge Two layers. Claude Code's native auto-memory holds operator facts and preferences (how to work with you); on top, the hermit adds a raw/ → compiled/ knowledge base — domain outputs and living topic pages updated in place — re-injected as a catalog within a context budget on fresh and resumed starts. Your Discord/Telegram DM text is also captured locally, so decisions made over chat outlive the thread: weekly-review distills them into memory (opt out with knowledge.channel_log_enabled: false ). /recall searches across all of it.
Plan tracking lives in the SHELL.md Progress Log — timestamped steps that survive compaction, restart, and every model tier.
Unattended safety combines profile-gated deny patterns + sandbox, channel-routed asks, permission-denial alerts, and injection scans on heartbeat and startup context.
Orchestrator instructed to delegate tasks & exploration to other agents, main context stays clean for token efficiency.
Sessions self-manage. Daemons auto-archive at 12h idle and at midnight when you're away, so evidence reaches the learning loop without a manual close. An external watchdog restarts dead sessions, nudges wedged ones, re-arms missed schedules, clears stale context after a midnight close, and compacts long-running context so cold wakes don't re-pay the full accumulated history — recovery never depends on the session being conscious.
Context-efficient continuity. After compaction, Hermit reloads only a bounded lifecycle/task/progress capsule instead of the full startup bundle. Structured report frontmatter lets briefs, reflections, and weekly reviews inspect history without rereading every report body.
It reaches you first. Notifications default to a native push (headless-friendly), or a Discord/Telegram DM you can reply to if you've paired a channel.
Cost scales with events, not time. Nothing wakes the model until something happens, so an idle hermit is effectively free.
A hermit watches what keeps going wrong across sessions, proposes a fix, and asks you yes or no. It won't propose the same thing twice.
At natural pauses — session end, idle ticks, scheduled cadence — it reflects. Most reflections never reach the model: a precheck script gates whether any phase (compute, resolution check, cost spike, digest, newborn) is actually due. When one is, two subagents vet the candidate before it reaches you:
reflection-judge confirms the cited evidence actually exists in the session reports, so a proposal can't certify itself.
proposal-triage deduplicates against open proposals, cross-checks your MEMORY.md and OPERATOR.md , and applies a three-condition bar.
Survivors land as a proposal you can act on from anywhere — including a DM:
/claude-code-hermit:proposal-list # see what it found
/claude-code-hermit:proposal-act accept PROP-003 # or just reply "accept PROP-003"
What it proposes: improvements, routines, new capabilities (skills, agents, heartbeat checks), guardrails (OPERATOR.md guidance you confirm), and bugs. When it catches itself repeating the same multi-step procedure across sessions, it drafts the skill and asks before installing. It improves its own skills too: when one keeps getting corrected or reworked across sessions, that graduates into a skill-improvement proposal, and on your okay it revises the skill (via skill-creator ). Accepted proposals can carry a measurable success signal and auto-resolve when met. You're the acceptance gate for every change. Raw session journals distill into compiled artifacts that reload next session — the raw/compiled pattern Karpathy described for his wiki-LLM.
The native Hermit Dashboard, proposals page, and weekly review stay current as Claude Code Artifacts at stable URLs — or on a self-hosted artifact server, if you configure one.
On-demand skills — pullable from the Claude app, your terminal, or a DM:
/hermit-dashboard-design — designs the Dashboard around what this hermit actually tracks; delete .claude-code-hermit/dashboard-render.ts to restore the built-in page
/recall — full-text search over past sessions, compiled knowledge, proposals, and your channel DM history ("what did I decide about X?")
/hermit-evolution — cost trend and behavior drift over weeks
/hermit-health — alerts, routines, channels, heartbeat state, plus fragile zones, stale proposals, and recent learnings
/hermit-doctor — proactive install diagnostic, from hook registration to heartbeat and routine-monitor liveness; the weekly check stays silent when green and alerts only on new problems
/cost-reflect — structural cost audit: which token types and trigger sources drive spend
/brief — current status and a summary of recent work
Prerequisites: Claude Code v2.1.172+, a Claude plan (Pro, Max, Teams, or Enterprise), and Bun 1.3+. Linux, macOS, and Windows via WSL2 — see FAQ .
cd /path/to/your/project # or any folder — even an empty one
claude plugin marketplace add gtapps/claude-code-hermit
claude plugin install claude-code-hermit@claude-code-hermit --scope local
2. Initialize
claude /claude-code-hermit:hatch
The wizard sets up your agent's identity, scans your folder, generates OPERATOR.md , and offers Quick (4 questions) or Advanced (full wizard).
Just trying it? After hatch , run .claude-code-hermit/bin/hermit-start --no-tmux for sessions, routines, heartbeat, and the learning loop without 24/7 autonomy. Run /claude-code-hermit:channel-setup first if you want Discord or Telegram.
/claude-code-hermit:docker-setup
Generates the Docker scaffolding, builds the image, starts the container, and walks through auth and channel pairing. The container ships with the hardening baseline ( cap_drop: ALL , no-new-privileges , pids_limit ). Want stronger isolation? Run /docker-security for opt-in LAN containment + DNS allowlist + resource bounds.
See Always-On Setup for the full guide. Want always-on without Docker? See Always-On Operations for bare tmux.
claude plugin update claude-code-hermit@claude-code-hermit --scope local
/claude-code-hermit:hermit-evolve
Or run .claude-code-hermit/bin/hermit-update (local/tmux) or .claude-code-hermit/bin/hermit-docker update (Docker): one command that moves the pin, reloads the session, and runs hermit-evolve for you.
Tune via /hermit-settings (or just by asking the hermit). Some of the settings available:
Full schema in the Config Reference
All live-editable with /hermit-settings (or just ask the hermit) — no reboot.
Model & Auto mode. Defaults to Sonnet — a good balance of reasoning and cost for an unattended session. Auto mode is generally available to all users across subscription plans and API usage; supported models and provider configuration can still vary, so if Claude reports the current selection unavailable, choose a supported model or another permission mode. Switch to opus for heavier reasoning; per-routine model: "haiku" remains useful for lightweight, isolated work.
Heartbeat. heartbeat.every sets the idle sweep (default 2h ; 1h tighter, 4h + fewer wakes); active_hours bounds the window ( 08:00 – 23:00 ). heartbeat.enabled: false stops timed wakes entirely — channels and routines still fire.
Idle behavior. discover (default) adds a priority-alignment pass against OPERATOR.md + cost log; wait is passive (tasks/channels only). Either way the daily reflect routine still runs — wait only silences between-schedule discovery, not the learning loop.
Routines. Each routine takes an optional model : run lightweight ones on haiku to save cost or heavier ones on opus for more reasoning, in an isolated subagent. Omit model to keep it inline in the main session context

[truncated]
