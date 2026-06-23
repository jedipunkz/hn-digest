---
source: "https://github.com/miopea/swarm"
hn_url: "https://news.ycombinator.com/item?id=48645622"
title: "Swarm – open-source control center for Claude Code / Gemini / Codex agents"
article_title: "GitHub - miopea/swarm: Web dashboard + autopilot for Claude Code agents — manage sessions, tasks, and approvals from your browser · GitHub"
author: "miopea"
captured_at: "2026-06-23T15:01:15Z"
capture_tool: "hn-digest"
hn_id: 48645622
score: 1
comments: 0
posted_at: "2026-06-23T14:34:46Z"
tags:
  - hacker-news
  - translated
---

# Swarm – open-source control center for Claude Code / Gemini / Codex agents

- HN: [48645622](https://news.ycombinator.com/item?id=48645622)
- Source: [github.com](https://github.com/miopea/swarm)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T14:34:46Z

## Translation

タイトル: Swarm – クロード コード / ジェミニ / コーデックス エージェント用のオープンソース コントロール センター
記事のタイトル: GitHub - miopea/swarm: Web ダッシュボード + クロード コード エージェントの自動操縦 — ブラウザーからセッション、タスク、承認を管理 · GitHub
説明: Web ダッシュボード + Claude Code エージェントの自動操縦 - ブラウザーからセッション、タスク、承認を管理 - miopea/swarm

記事本文:
GitHub - miopea/swarm: Web ダッシュボード + クロード コード エージェントの自動操縦 — ブラウザーからセッション、タスク、承認を管理 · GitHub
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
ミオペア
/
群れ
公共
通知
通知設定を変更するにはサインインする必要があります
さらに

l ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,231 コミット 1,231 コミット .claude .claude .github/ workflows .github/ workflows docs docs scripts scripts src/ swarm src/ swarm テスト テスト .gitignore .gitignore .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml swarm.yaml.example swarm.yaml.example すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェント用の Web ベースのコントロール センター — Claude Code 、Gemini CLI 、および Codex CLI 。オートパイロット、タスクボード、AI 調整、電子メール統合を使用して、単一のブラウザー タブから 1 人または 10 人のエージェントを管理します。
すべてのエージェント セッションは管理対象 PTY で実行されます。 Web ダッシュボードでは、それらすべてをリアルタイムで把握できます。出力を読み取り、端末に入力し、タスクを作成して割り当て、バックグラウンドのドローンに日常的な承認を処理させることで、エージェントが失速することはありません。クイーンの指揮者はハイブを監視し、タスクの割り当てを提案し、作業が完了したことを検出し、電子メールの返信の下書きを作成します。これらすべてが、ワンクリックで承認される提案として表示されます。
エージェントのセッションが停止することはありません。 Swarm の背景調査スタッフであるドローンは、安全なプロンプトを自動承認し、墜落したエージェントを復活させ、難しい決定を女王 (首のないクロードの車掌) またはオペレーターにエスカレーションします。あなたは子守りをやめて、結果の確認を始めます。
仕事を管理するのは Windows ではなく、あなたです。ボード上にタスクを作成します。女王は、プロジェクトの説明に基づいて、適切な労働者を彼らに割り当てます。労働者が終了すると、女王がそれを検知して完了を提案します。ワンクリックで承認します。
あなたのブラウザはコントロールルームです。対話型ターミナル接続により、任意のワーカーのエージェント セッションに直接入力できます。

彼のダッシュボード。電子メールをタスク ボードにドラッグしてバグ チケットを作成します。修正されると、Outlook に返信の下書きが表示されます。
1セッションでも有効です。メリットを得るために 10 人のエージェントは必要ありません。単一のエージェント セッションでも、自動操縦、タスク キュー、ターミナル アクセスを備えたダッシュボードを利用できます。
Web ダッシュボード (プライマリ インターフェイス)
ライブターミナルアタッチ -- ブラウザから任意のワーカーのエージェントセッションに入力します (PTY over WebSocket)
タスク ボード -- 1 行または 2 行のコンパクトなタスク行。行をクリックすると、WYSIWYG 記述エディター (書式設定ツールバー、Word/Outlook から貼り付け → マークダウン、ソース表示切り替え)、優先順位/フィルター、依存関係、インライン ファイル添付を備えた完全な編集モーダルが開きます。
ドラッグ アンド ドロップ インポート -- .eml / .msg ファイル (または Outlook メッセージ タイル、または Jira 課題 URL / KEY-N ) をタスク ボードにドロップし、インポートされ、Markdown としてレンダリングされたソース コンテンツを含むタスクを作成します。
Queen の提案 -- 信頼スコアを使用した AI 推奨をワンクリックまたは一括で承認または拒否します。
構成エディタ -- タブ付き UI: 一般、LLM、ワーカー、自動化、通知、統合、セキュリティ、使用状況、詳細、ログ
承認ルールエディタ -- ドローンの自動承認/エスカレーション決定のための視覚的な正規表現ルールビルダー
ワーカー管理 -- アドホック ワーカーの生成、グループの起動、個人の強制終了/復活など、すべて実行時に実行します。
Outlook の統合 -- 構成ページから OAuth 経由で接続し、電子メールを直接取得します
ブラウザ通知 -- 作業者が注意を必要とするときにアラートをプッシュします
バックグラウンドドローン -- 専門の監視員が日常的な意思決定を処理するため、作業員が失速することはありません (完全な名簿については、以下のドローンを参照してください)
女王の車掌 -- タスクを割り当て、完了を検出し、競合を解決する首のないクロード
提案システム -- クイーンのアクションにはオペレーターの承認が必要です。サインオフがなければ何も実行されません
承認ルール -- 規則

x パターンにより、ドローンが自動承認するものと女王にエスカレーションするものを決定します
スキル ワークフロー -- タスクはクロード コードのスキル コマンド ( /fix-and-ship 、 /feature 、 /verify ) としてディスパッチされ、スキルごとの使用数を示す SQLite レジストリによってサポートされます。
ワーカーごとの Swarm スラッシュ コマンド -- すべてのワーカーは、 /swarm-status 、 /swarm-handoff 、 /swarm-finding 、 /swarm-warning 、 /swarm-blocker 、 /swarm-progress を .claude/commands/ に自動インストールするため、最もよく使用される調整ツールが /help に表示され、トランスクリプトできれいに読み取られます。
ワーカーごとの Swarm スキル -- ワーカーは、 /swarm-checkpoint スキル ( /check を実行し、緑色でコミットするか、赤色でブロッカーを報告する) と /swarm-coowned スキル (委任の提案に対する助言ピア/タスク調査。タスクは決して自動作成されません) も自動インストールします。
パイプライン -- エージェント、自動化、人間によるステップ、依存関係の順序付け、テンプレート、および 5 フィールドの cron スケジュールを含む複数ステップのワークフロー (例: 平日の午後は「30 14 * * 1-5」。従来の HH:MM は引き続き機能します)
承認率ゲージ -- ダッシュボードのヘッダーには、過去 24 時間のドローンの自動承認率が表示されます。 GET /api/drones/approval-rate でカウンターが公開される
サンドボックス オプトイン -- swarm.yaml のサンドボックス経由でクロード コードのネイティブ サンドボックスを有効にします。 Swarm はインストール時に CC バージョンを検出し、サポートされている場合はオーバーライドを ~/.claude/settings.json にマージします。
ドローンは、デーモンのポーリング ループを共有する特殊なバックグラウンド スイーパーです。それぞれが独自のリズムで実行され、すべてのアクションがバズ ログに書き込まれるため、オペレーターは監査と調整を行うことができます。
IdleWatcher -- タスクが割り当てられている休息中/睡眠中のワーカーを微調整します。 /mcp を挿入し、タスクの説明をフォローアップすることで、クライアント側 MCP ツールがドロップしたリロード後のセッションを回復します
InterWorkerMessageWatcher -- アイドル状態のワーカーに未読のワーカー間メッセージを通知します。フィルターを広げます

作業者にアクティブなタスクがない場合、情報の所見やメモが失われないようにする
PressureManager -- ホストレベルの圧力下でワーカーを一時停止および再開する、システム全体のメモリ/スワップ ウォッチャー
ContextPressure ドローン -- ワーカーごとの context_pct を監視します。会話がいっぱいになったときに /compact を挿入します (ソフト層はアイドル状態のワーカーを自動圧縮します。ハード層は活発なワーカーを中断し、待機中のワーカーを延期します)
Verifier ドローン -- swarm_complete_task のたびに実行される敵対的な完了後チェック。層 1 の決定論的ゲート (空の差分 / 証拠なし / チェック証拠 / オープンピア警告) は LLM 呼び出しの前に短絡し、層 2 は専用の検証サブプロセスを呼び出します。結果をピア警告としてタスクを再度開くか、2 回連続して再試行が失敗した後にクイーン スレッドにエスカレーションします。 queen_force_complete_task によるオペレーターのオーバーライドは検証をスキップします。
OversightHandler / TaskLifecycle / FileOwnership / StateTracker -- Queen の監視、タスクの移行、ファイル要求の調整、および状態分類のためのドローンをサポートします。
MCP サーバー -- Swarm は /mcp で HTTP MCP サーバーを公開し、エージェント自体がツール呼び出しを通じて調整できるようにします。
15 個の調整ツール -- swarm_check_messages 、 swarm_send_message 、 swarm_task_status 、 swarm_claim_file 、 swarm_complete_task 、 swarm_create_task 、 swarm_park_task (現在のタスクをキューに戻す)、 swarm_get_learnings 、 swarm_get_playbooks (合成された再利用可能なプロシージャの呼び出し)過去の成功から)、swarm_report_progress 、swarm_report_blocker (タスク依存性ブロッカーを宣言し、アイドル ナッジを抑制)、swarm_query_peers (ハンドオフ決定のためのピア ワーカー状態の読み取り専用スナップショット)、swarm_note_to_queen (軽量のサイドチャネル メモ)、swarm_draft_email (オペレーターの Outlook 下書きフォルダーに Microsoft Graph ドラフトを作成) — 自動的に送信されることはありません

ly)、および swarm_batch (1 回のラウンドトリップで複数の操作を実行)
ワーカー間メッセージ -- ワーカーは検出結果、警告、依存関係、およびステータスの更新を相互に送信 (またはブロードキャスト)
ファイルのクレーム -- アドバイザリー ロックにより、2 人の作業者が同じファイルを同時に編集することができなくなります
学習 -- 完了したタスクからの解決策は、他の作業者がコンテキストを検索して検索できます。
コンパクト テレメトリ -- すべての /compact は前後のトークンと圧縮率をログに記録するため、ワーカーごとの圧縮の効果を測定できます。
サービス ハンドラー (パイプラインの自動ステップ)
shell_command 、 webhook_notify 、 headless_claude -- シェル コマンドを実行するか、Webhook をポストするか、パイプライン ステップとしてヘッドレス クロードを呼び出します。
file_uploader 、 youtube_scraper -- ファイルをシンクにアップロードし、追跡されている YouTube チャンネルから新しいビデオを取得します
claude_code_security -- claude code security scan --json を実行し、 (rule_id, path, line) によって検出結果の重複を排除し、それらを下流のステップに返してタスクに変換します。
Jira の統合 -- Jira Cloud (OAuth 2.0) との双方向同期、タスクのインポート/エクスポート、タスク ボードからの Jira 課題の作成
REST API -- 80 以上のエンドポイントを備えた完全な JSON API、および /api/docs/ui にある OpenAPI ドキュメント (ダッシュボードを実行している状態で http://localhost:9090/api/docs/ui を開きます)
SQLite の永続性 -- タスク、提案、メッセージ、パイプライン、スキル、履歴は ~/.swarm/swarm.db に保存されます。 YAML はシード/インポート形式です
リソースの監視 -- メモリ/スワップのしきい値と、オプションでシステムの圧力に応じてワーカーを自動一時停止する機能
アプリ内フィードバック -- フッター ボタンをクリックすると、バグ/機能/質問フォームが開きます。提出物は gh CLI 経由で GitHub の問題としてファイルされ、プレビューと編集のステップと機密パスの自動編集が行われます。
リモート アクセス -- 電話またはリモート マシンからダッシュボードにアクセスするための Cloudflare トンネルのサポート。トンネル経由のオプションの名前付きドメイン

_ドメイン
通知 -- 端末のベル、デスクトップ、ブラウザのプッシュ アラート
ツール使用状況分析 -- swarm 分析ツールは、MCP コール、エラー、アクティブなワーカーをバズ ログから要約するため、書き換えが必要なツールの説明を特定できます。
テスト ハーネスの再現性 -- swarm test --pin-model=<id> は、すべてのテスト レポートにインフラ スナップショット (モデル、プロバイダー、ワーカー数、環境フィンガープリント) を記録するため、回帰は不可解なものではなくデバッグ可能になります。
Python 3.12+ (Swarm が swarm.db に使用する SQLite 3 stdlib に付属)
GitHub CLI ( gh ) — オプション。アプリ内フィードバック送信者にのみ必須
WSL ユーザー: 自動開始サービスのために WSL 内で systemd を有効にする必要があります。 swarm init はそうでない場合を検出し、/etc/wsl.conf を設定するように提案します (sudo が必要です)。
少なくとも 1 つの AI コーディング エージェント CLI:
Claude Code ( claude ) — 本番環境に対応しており、Queen の指揮者にも電力を供給します
Gemini CLI ( gemini ) — 実験的
Codex CLI ( codex ) — 実験的
uv ツールをインストール git+https://github.com/miopea/swarm.git
これにより、PATH に swarm が追加されます。クローンもvenvもありません。次に、セットアップ ウィザードを実行します。
群れの初期化
これにより次の 4 つのことが行われます。
Claude Code フックをインストールします -- 安全なツール (読み取り、編集、書き込み、Glob、Grep) を自動承認するため、ワーカーがすべての Fi で停止することはありません

[切り捨てられた]

## Original Extract

Web dashboard + autopilot for Claude Code agents — manage sessions, tasks, and approvals from your browser - miopea/swarm

GitHub - miopea/swarm: Web dashboard + autopilot for Claude Code agents — manage sessions, tasks, and approvals from your browser · GitHub
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
miopea
/
swarm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,231 Commits 1,231 Commits .claude .claude .github/ workflows .github/ workflows docs docs scripts scripts src/ swarm src/ swarm tests tests .gitignore .gitignore .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml swarm.yaml.example swarm.yaml.example View all files Repository files navigation
A web-based control center for AI coding agents — Claude Code , Gemini CLI , and Codex CLI . Manage one agent or ten from a single browser tab — with autopilot, a task board, AI coordination, and email integration.
Every agent session runs in a managed PTY. The web dashboard gives you real-time visibility into all of them: read their output, type into their terminals, create and assign tasks, and let background drones handle routine approvals so your agents never stall. A Queen conductor watches the hive, proposes task assignments, detects when work is done, and drafts email replies — all surfaced as proposals you approve with one click.
Your agent sessions never stall. Drones — Swarm's background poll workers — auto-approve safe prompts, revive crashed agents, and escalate the hard decisions to the Queen (a headless Claude conductor) or the operator. You stop babysitting and start reviewing results.
You manage work, not windows. Create tasks on a board. The Queen assigns them to the right worker based on project descriptions. When a worker finishes, the Queen detects it and proposes completion — you approve with one click.
Your browser is the control room. Interactive terminal attach lets you type directly into any worker's agent session from the dashboard. Drag an email onto the task board to create a bug ticket. When it's fixed, a draft reply lands in your Outlook.
It works for one session too. You don't need ten agents to benefit. Even a single agent session gets autopilot, a task queue, and a dashboard with terminal access.
Web Dashboard (primary interface)
Live terminal attach -- type into any worker's agent session from the browser (PTY over WebSocket)
Task board -- compact one-or-two-line task rows; click a row to open the full Edit modal with a WYSIWYG description editor (formatting toolbar, paste-from-Word/Outlook → Markdown, View-source toggle), priority/filtering, dependencies, and inline file attachments
Drag-and-drop import -- drop .eml / .msg files (or an Outlook message tile, or a Jira issue URL / KEY-N ) onto the task board to create a task with the source content imported and rendered as Markdown
Queen proposals -- approve or reject AI recommendations with confidence scores, one click or in bulk
Config editor -- tabbed UI: General, LLMs, Workers, Automation, Notifications, Integrations, Security, Usage, Advanced, Logs
Approval rules editor -- visual regex rule builder for drone auto-approve/escalate decisions
Worker management -- spawn ad-hoc workers, launch groups, kill/revive individuals, all at runtime
Outlook integration -- connect via OAuth from the config page, fetch emails directly
Browser notifications -- push alerts when workers need attention
Background drones -- specialized watchers handle routine decisions so workers don't stall (see Drones below for the full roster)
Queen conductor -- headless Claude that assigns tasks, detects completion, resolves conflicts
Proposal system -- Queen actions require operator approval; nothing executes without your sign-off
Approval rules -- regex patterns decide what drones auto-approve vs escalate to the Queen
Skill workflows -- tasks dispatch as Claude Code skill commands ( /fix-and-ship , /feature , /verify ), backed by a SQLite registry with per-skill usage counts
Per-worker Swarm slash commands -- every worker auto-installs /swarm-status , /swarm-handoff , /swarm-finding , /swarm-warning , /swarm-blocker , /swarm-progress into its .claude/commands/ so the most-used coordination tools show up in /help and read cleanly in transcripts
Per-worker Swarm Skills -- workers also auto-install the /swarm-checkpoint Skill (runs /check , then commits on green or reports a blocker on red) and the /swarm-coordinate Skill (advisory peer/task survey for delegation suggestions; never auto-creates tasks)
Pipelines -- multi-step workflows with agent, automated, and human steps, dependency ordering, templates, and 5-field cron schedules (e.g. "30 14 * * 1-5" for weekday afternoons; legacy HH:MM still works)
Approval-rate gauge -- dashboard header shows the drones' auto-approval percentage over the last 24h; GET /api/drones/approval-rate exposes the counters
Sandbox opt-in -- enable Claude Code's native sandbox via sandbox: in swarm.yaml ; Swarm detects CC version at install time and merges the overrides into ~/.claude/settings.json when supported
Drones are specialized background sweepers that share the daemon's poll loop. Each runs at its own cadence and writes every action to the buzz log so the operator can audit and tune.
IdleWatcher -- nudges RESTING/SLEEPING workers that have an assigned task; recovers post-reload sessions whose client-side MCP tools dropped by injecting /mcp and following up with the task description
InterWorkerMessageWatcher -- nudges idle workers about unread inter-worker messages; widens the filter when the worker has no active task so informational findings/notes are not lost
PressureManager -- system-wide memory/swap watcher that suspends and resumes workers under host-level pressure
ContextPressure drone -- watches per-worker context_pct ; injects /compact when the conversation fills (soft tier auto-compacts idle workers; hard tier interrupts BUZZING workers and defers WAITING ones)
Verifier drone -- adversarial post-completion check that fires after every swarm_complete_task ; tier 1 deterministic gates (empty diff / no /check evidence / open peer warning) short-circuit before any LLM call, tier 2 calls a dedicated verifier subprocess. Reopens the task with findings as a peer warning, or escalates to a Queen thread after the second consecutive failed retry. Operator override via queen_force_complete_task skips verification.
OversightHandler / TaskLifecycle / FileOwnership / StateTracker -- supporting drones for Queen oversight, task transitions, file-claim coordination, and state classification
MCP server -- Swarm exposes an HTTP MCP server at /mcp so the agents themselves can coordinate via tool calls
15 coordination tools -- swarm_check_messages , swarm_send_message , swarm_task_status , swarm_claim_file , swarm_complete_task , swarm_create_task , swarm_park_task (hand the current task back to the queue), swarm_get_learnings , swarm_get_playbooks (recall reusable procedures synthesized from past successes), swarm_report_progress , swarm_report_blocker (declare task-dependency blocker, suppresses idle nudges), swarm_query_peers (read-only snapshot of peer worker state for handoff decisions), swarm_note_to_queen (lightweight side-channel note), swarm_draft_email (create a Microsoft Graph draft in the operator's Outlook Drafts folder — never sent automatically), and swarm_batch (run multiple ops in one round-trip)
Inter-worker messages -- workers send findings, warnings, dependencies, and status updates to each other (or broadcast)
File claims -- advisory locks prevent two workers from editing the same file at once
Learnings -- resolutions from completed tasks are searchable by other workers for context
Compact telemetry -- every /compact logs tokens before/after and the compression ratio so you can measure how effective compaction is per worker
Service handlers (pipeline AUTOMATED steps)
shell_command , webhook_notify , headless_claude -- run shell commands, post webhooks, or invoke a headless Claude as a pipeline step
file_uploader , youtube_scraper -- upload files to a sink, pull new videos from tracked YouTube channels
claude_code_security -- run claude code security scan --json , deduplicate findings by (rule_id, path, line) , and return them for downstream steps to turn into tasks
Jira integration -- two-way sync with Jira Cloud (OAuth 2.0), import/export tasks, create Jira issues from the task board
REST API -- full JSON API with 80+ endpoints and OpenAPI docs at /api/docs/ui (open http://localhost:9090/api/docs/ui with the dashboard running)
SQLite persistence -- tasks, proposals, messages, pipelines, skills, and history are stored in ~/.swarm/swarm.db ; YAML is the seed/import format
Resource monitoring -- memory/swap thresholds with optional auto-suspend of workers on system pressure
In-app feedback -- a footer button opens a bug / feature / question form; submissions are filed as GitHub issues via the gh CLI, with a preview-and-edit step and automatic redaction of sensitive paths
Remote access -- Cloudflare Tunnel support for reaching the dashboard from a phone or remote machine; optional named domain via tunnel_domain
Notifications -- terminal bell, desktop, and browser push alerts
Tool-usage analytics -- swarm analyze-tools summarises MCP calls, errors, and active workers from the buzz log so you can spot tool descriptions that need rewriting
Test harness reproducibility -- swarm test --pin-model=<id> records an infra snapshot (model, provider, worker count, env fingerprint) in every test report so regressions are debuggable instead of mysterious
Python 3.12+ (ships with the SQLite 3 stdlib Swarm uses for swarm.db )
GitHub CLI ( gh ) — optional; required only for the in-app feedback submitter
WSL users: systemd must be enabled inside WSL for the auto-start service. swarm init detects when it's not and offers to configure /etc/wsl.conf for you (requires sudo).
At least one AI coding agent CLI:
Claude Code ( claude ) — production-ready, also powers the Queen conductor
Gemini CLI ( gemini ) — experimental
Codex CLI ( codex ) — experimental
uv tool install git+https://github.com/miopea/swarm.git
This puts swarm on your PATH. No clone, no venv. Then run the setup wizard:
swarm init
This does four things:
Installs Claude Code hooks -- auto-approves safe tools (Read, Edit, Write, Glob, Grep) so workers don't stall on every fi

[truncated]
