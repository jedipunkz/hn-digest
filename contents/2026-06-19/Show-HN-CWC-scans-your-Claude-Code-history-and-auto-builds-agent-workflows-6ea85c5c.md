---
source: "https://github.com/fayzan123/claude-workflow-composer"
hn_url: "https://news.ycombinator.com/item?id=48602418"
title: "Show HN: CWC scans your Claude Code history and auto-builds agent workflows"
article_title: "GitHub - fayzan123/claude-workflow-composer: Visual desktop app for composing multi-agent coding workflows. Drag agents, attach skills and MCPs, wire handoffs, export to .claude/ · GitHub"
author: "FayzanMalik"
captured_at: "2026-06-19T21:30:23Z"
capture_tool: "hn-digest"
hn_id: 48602418
score: 1
comments: 0
posted_at: "2026-06-19T19:42:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CWC scans your Claude Code history and auto-builds agent workflows

- HN: [48602418](https://news.ycombinator.com/item?id=48602418)
- Source: [github.com](https://github.com/fayzan123/claude-workflow-composer)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T19:42:11Z

## Translation

タイトル: HN を表示: CWC がクロード コードの履歴をスキャンし、エージェント ワークフローを自動構築します
記事のタイトル: GitHub - fayzan123/claude-workflow-composer: マルチエージェント コーディング ワークフローを構成するためのビジュアル デスクトップ アプリ。エージェントをドラッグし、スキルと MCP をアタッチし、ハンドオフを配線し、.claude/ にエクスポートする · GitHub
説明: マルチエージェント コーディング ワークフローを構成するためのビジュアル デスクトップ アプリ。エージェントをドラッグし、スキルと MCP をアタッチし、ハンドオフを配線し、.claude/ にエクスポートする - fayzan123/claude-workflow-composer

記事本文:
GitHub - fayzan123/claude-workflow-composer: マルチエージェント コーディング ワークフローを構成するためのビジュアル デスクトップ アプリ。エージェントをドラッグし、スキルと MCP をアタッチし、ハンドオフを配線し、.claude/ にエクスポートする · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
フェイザン123
/
クロード・ワークフロー・コンポーザー
公共
通知

ション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
fayzan123/クロード-ワークフロー-コンポーザー
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
289 コミット 289 コミット .github/ workflows .github/ workflows bin bin client client docs docs src src testing testing .gitignore .gitignore .impeccable.md .impeccable.md CLAUDE.md CLAUDE.md DESIGN.md DESIGN.md ライセンス ライセンス README.md README.md デモ.gif デモ.gif パッケージ-lock.json パッケージ-lock.json パッケージ.json パッケージ.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェント ワークフロー用の n8n。マルチエージェントの Claude Code ワークフローを作成するためのビジュアル デスクトップ アプリ — エージェントをキャンバスにドラッグし、ハンドオフを配線し、スキルをアタッチし、作業中のワークフローを Claude インストールに直接エクスポートします。 YAML 編集は必要ありません。
現在、Claude Code でマルチエージェント ワークフローを構築することは、次のことを意味します。
YAML フロントマターを使用した手書きエージェント .md ファイル
disable-model-invocation を使用したオーケストレーター スキルの手動作成: 真の順序付けされたハンドオフの散文
実行前のパイプラインを視覚的に表現しない
完全に機能するワークフローを他の人と共有する方法がない
優れたパイプラインがどのようなものかを知る方法がない
オーサリング エクスペリエンスは完全にテキストベースです。実行するまで、何を構築しているのかを見ることはできません。
npxクロード-cwc
http://localhost:3579 でブラウザを開きます。ローカル サーバーはループバックにバインドし、実行ごとのローカル トークンを使用して API 呼び出しを保護します。
ホーム ダッシュボードからオートメーションの検出を使用して、ローカルのクロード コード履歴をスキャンし、繰り返される作業を見つけて、最も有力な候補からすぐに編集できるワークフローを生成します。
npx claude-cwc stop # サーバーを停止します
macOS でスケジュールされた au が必要な場合は、バックグラウンド サービスをインストールしてください

再起動後も実行を続けるためのオプション:
npx クロード-cwc インストール-サービス
npx クロード-cwc アンインストール サービス
またはソースから:
npm run build && npm start
仕組み
エージェントをキャンバスにドラッグします
→ハンドオフ矢印で繋ぐ（オーサートリガー条件）
→ 各エージェントのシステム プロンプト、ツール、スキル、完了基準を編集します
→ ワークフローが自動的に実行される必要がある場合は、自動モードでスケジュール/Webhook を追加します。
→ エクスポートする前に、書き込まれるすべてのファイルをプレビューします。
→ エクスポート → エージェント .md ファイル + オーケストレーター SKILL.md を ~/.claude/ に書き込みます
→ クロードコードでワークフローを名前で呼び出す
エクスポータは、~/.claude/ (ユーザー スコープ) または任意のプロジェクト ディレクトリ内の .claude/ (プロジェクト スコープ、バージョン制御可能) に直接書き込みます。競合検出により、所有していないファイルには決して触れなくなります。
既存のエージェントをサイドバー ( ~/.claude/agents/ ) からキャンバスにドラッグして参照ノードを作成します。参照ノードは、エージェント ファイルを複製するのではなくスラッグによって参照します。 「新規/空のエージェント」からドラッグしてオーダーメイドのノードを作成します。エクスポータはそのノード用に新しいエージェント ファイルを生成します。
ハンドル間をドラッグしてノードを接続します。各接続は、トリガーの説明とオプションのコンテキスト アーティファクト (ファイル、テキスト、JSON) がエージェント間で渡されるハンドオフになります。ワークフローの終了状態を定義するには、任意のノードをターミナル ( Complete 、 Escalated 、または Aborted ) としてマークします。
ノード パネル で、ノードの完了基準、ツール アクセス、スキル、およびシステム プロンプトを編集します。最初のノードには、ワークフローを開始する内容を説明する開始トリガーを含めることもできます。
リアルタイム検証により、エクスポートする前に、重複したスラグ、空の名前、切断されたノード、欠落している完了基準が上部バーにすぐに表示されます。
サイドバーの「エージェントの生成」または「スキルの生成」を使用して、再利用可能なクロード コード アセットを平易な英語からドラフトします。

っぽい説明。 CWC は最初に編集可能な仕様を提供し、次にそのファイルを ~/.claude/agents/ または ~/.claude/skills/ に書き込みます。
上部のバーで「エクスポート」をクリックします。ターゲット ディレクトリ ( ~/.claude/ または任意のプロジェクトの .claude/ ) を選択します。書き込まれるすべてのファイルのプレビューを確認します。書き込みを確認します。
特注ノード → フロントマター (名前、説明、色、モデル、ツール)、システム プロンプト、完了基準、スキル参照、および所有権コメントを含むエージェント .md ファイルを書き込みます。
ノードを参照 → 何も書き込みません。exportedSlug は既存のエージェントのスラグに設定されるため、オーケストレーターはそこに直接ルーティングします。
ワークフロー スキル → disable-model-invocation: true を指定して .claude/skills/<workflow-slug>/SKILL.md にオーケストレーター スキルを生成します。オーケストレーター本体は、ノード/エッジ グラフを自然言語ステップに BFS トラバースすることによって生成されます。
名前変更処理 → ノードの名前が変更された場合、古い所有ファイルが削除され、新しいファイルが書き込まれます。
競合の検出 → すべてのファイルには所有権の HTML コメントが含まれます。上書きまたは削除する前に、エクスポーターは所有権を確認します。他のワークフローや手動で作成されたファイルには決して触れません。
任意の Claude Code セッションから、スキル名でワークフローを呼び出します。
/ワークフロー名
オーケストレーター スキルは、エージェント ツールを介してすべての実装ステップをサブエージェントに委任します。各ステップはエージェントを名前で参照します。 Claude Code はそれをエージェントの .md ファイルに解決し、そのシステム プロンプト、ツール、および完了基準をロードします。
ワークフローの自動モードを開いて、スケジュールと Webhook を追加します。
Cron — スケジュール ビルダーを使用するか、カスタム cron 式を入力します (例: 平日午前 9 時の場合は 0 9 * * 1-5)。
Webhook — CWC は受信ローカル URL を生成します。 HTTP POST を送信してワークフローを起動します。
作業ディレクトリ/ターゲット — 自動化を実行する場所を選択します。

ファンアウト用のオプションの追加リポジトリ。
分離 — 分離されたブランチに git ワークツリーを使用するか、現在のチェックアウトを明示的に必要な場合にインプレースで実行します。
前提条件 — CWC が実行を開始する前に成功する必要があるシェル コマンド。
セットアップ コマンド — 実行開始後、クロードが開始される前に CWC が実行するシェル コマンド。
ワークフローの任意の時点でゲート ノードを追加します (サイドバーの「ゲート」セクションからドラッグします)。ランがゲートに到達すると、次のようになります。
すべての変更を cwc/<runId> ブランチにコミットし、一時停止します。
作業ブランチの差分を [実行] パネルの受信トレイに投稿します。
待機 — レビュー担当者は差分を読み、オプションのメモを書き、 [承認] または [拒否] をクリックします。
承認されると、同じクロード コード セッションでゲート ポイントから実行が再開されます。
実行パネルのヘッ​​ダーには、トリガーを解除せずにスケジュールされたすべての実行をグローバルに一時停止する自動トグルと、通知設定 (macOS バナーおよび/または Webhook URL) を開く ⚙ ギアがあります。
ホーム ダッシュボードの [オートメーションの検出] をクリックして、ローカルのクロード コード履歴をスキャンして、繰り返し行われた作業を探します。 CWC は、ローカルのトランスクリプト ファイルを解析し、コンパクトなダイジェストを構築し、繰り返しのタスクをクラスタリングするようクロードに依頼し、候補者に証拠、確信度、観察されたステップ、および推奨されるトリガーを示します。
候補で [ワークフローの生成] をクリックして、候補を実際の .cwc ワークフローにプロモートします。 CWC は、一致するローカル スキルと既存のエージェントを探し、クロードにワークフローの作成を依頼し、生成されたグラフを検証し、必要に応じて無効なスケジュール トリガーをシードし、レビューのためにワークフローを開きます。
POST /api/export/delete は、エクスポートされたすべてのファイルをスキャンし、その所有権コメントをチェックして、現在のワークフローが所有するファイルのみを削除します。参照ノードには削除するものは何もありません。ファイルは書き込まれませんでした。
ビジュアル キャンバス — 背景グリッド、ミニマップを備えた React Flow

、ズーム コントロール、ドラッグして接続
テーマの切り替え - ホーム ダッシュボードまたはワークフロー ヘッダーからライト モードとダーク モードを切り替えます。
左側のサイドバー — 私のエージェント (検索可能、~/.claude/agents/ からドラッグ可能) およびスキル (検索可能、選択したノードにドラッグ可能)
エージェント/スキルの生成 — 平易な英語から新しい再利用可能なクロード コード アセットをドラフトし、仕様を調整して、~/.claude/ に保存します。
右側のパネル - ノード エディター (名前、説明、基準、ツール、スキル、システム プロンプト、端末タイプ) およびエッジ エディター (トリガー、ラベル、コンテキスト アーティファクト)
エクスポートモーダル — ターゲットの選択、ファイル全体のプレビュー、書き込み前の警告表示
自動保存 — 500 ミリ秒のデバウンス保存で ~/.cwc/workflows/ に保存されます。手動保存は必要ありません
最近のファイル — ホーム画面には、~/.cwc/recents.json に保存された最新 10 件のワークフローが表示されます
マークダウン プレビュー — エージェントまたはスキル カードをクリックしてソース ファイルを表示します
エディターで開く - システム エディターでエージェントまたはスキル ファイルを表示します
クロード コードの検出 — ~/.claude/ が見つからない場合に起動時に警告します
▶ テスト実行 — エクスポートされたワークフローを UI からヘッドレスで起動し ( --permission-mode bypassPermissions 、ユーザーが選択した作業ディレクトリ、ワークツリー、またはインプレース分離)、実行中に停止します。
ライブ実行ビュー — アクティブなノードがキャンバス上で脈動し、完了したノードがチェックされ、イベントがタイムライン パネルにストリーミングされます。
実行履歴 — エクスポートされたすべてのワークフロー (CWC または任意の端末から開始) のすべての実行は、ステータス、期間、ソース、コストとともに ~/.cwc/runs/ に保存されます。
自動化モード - cron スケジュールまたは Webhook URL をワークフローに添付し、ターゲット/分離を選択し、前提条件/セットアップ コマンドを追加し、信頼できるトリガーを準備します
承認ゲート — ワークフローにゲート ノードを挿入します。到達すると実行が一時停止され、作業ブランチの差分が投稿され、レビュー担当者が承認または拒否します。

受信箱 (または端末) から実行すると、同じセッションで実行が再開されます。
分離された実行 — テスト実行 (およびスケジューラー起動の実行) は、cwc/<runId> ブランチに git ワークツリーを作成します。これにより、メイン チェックアウトは常に変更されません。ワークツリーは実行完了後に削除されます
通知 — macOS バナー + 実行完了、ゲート一時停止、承認リクエスト時のオプションの Webhook。実行パネルの設定ギアから設定
グローバル一時停止 - [実行] パネルの 1 つの切り替えにより、トリガーを解除せずに、スケジュールされたすべてのオートメーションの実行が一時停止されます。
自動化の検出 — ローカルのクロード コードのトランスクリプトをスキャンし、繰り返される作業をクラスタリングし、進行状況をストリーミングし、自動化を提案し、可能であれば再利用されたスキル/エージェントと一致する .cwc ワークフローに候補者を昇格させます。
クライアント (React + React Flow) サーバー (Express :3579)
┌───────────────┐ ┌───────────────┐
│ TemplatePicker │ ──► │ /api/workflows │
│ トップバー │ ◄── │ /api/recents │
│ サイドバー (エージェント/スキル) │ ──► │ /api/agents │
│ Canvas (React Flow) │ ──► │ /api/skills │
│ NodePanel / EdgePanel │ ──► │ /api/export │
│ ExportFlow (モーダル) │ ──► │ /api/export/preview │
│ RunModal / RunPanel │ ──► │ /api/export/delete │
│ useWorkflow (リデューサー) │ │ /api/runs (+SSE) │
│ useAutoSave (デバウンス) │ │ /api/automations │
│ useRunEvents (SSE) │ ◄── │ /api/triggers │
│ 自動検出

[切り捨てられた]

## Original Extract

Visual desktop app for composing multi-agent coding workflows. Drag agents, attach skills and MCPs, wire handoffs, export to .claude/ - fayzan123/claude-workflow-composer

GitHub - fayzan123/claude-workflow-composer: Visual desktop app for composing multi-agent coding workflows. Drag agents, attach skills and MCPs, wire handoffs, export to .claude/ · GitHub
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
fayzan123
/
claude-workflow-composer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
fayzan123/claude-workflow-composer
main Branches Tags Go to file Code Open more actions menu Folders and files
289 Commits 289 Commits .github/ workflows .github/ workflows bin bin client client docs docs src src tests tests .gitignore .gitignore .impeccable.md .impeccable.md CLAUDE.md CLAUDE.md DESIGN.md DESIGN.md LICENSE LICENSE README.md README.md demo.gif demo.gif package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
n8n for coding agent workflows. A visual desktop app for composing multi-agent Claude Code workflows — drag agents onto a canvas, wire handoffs, attach skills, and export a working workflow directly into your Claude installation. No YAML editing required.
Building multi-agent workflows in Claude Code today means:
Hand-writing agent .md files with YAML frontmatter
Manually authoring orchestrator skills with disable-model-invocation: true and sequenced handoff prose
No visual representation of the pipeline before running it
No way to share a complete, working workflow with someone else
No way to discover what good pipelines look like
The authoring experience is entirely text-based. You can't see what you're building until you run it.
npx claude-cwc
Opens a browser at http://localhost:3579 . The local server binds to loopback and protects API calls with a per-run local token.
Use Detect automations from the Home dashboard to scan your local Claude Code history, find repeated work, and generate a ready-to-edit workflow from the strongest candidates.
npx claude-cwc stop # Stop the server
On macOS, install the background service if you want scheduled automations to keep running after reboot:
npx claude-cwc install-service
npx claude-cwc uninstall-service
Or from source:
npm run build && npm start
How It Works
Drag agents onto a canvas
→ Connect them with handoff arrows (author trigger conditions)
→ Edit each agent's system prompt, tools, skills, and completion criteria
→ Add schedules/webhooks in Automate mode when the workflow should run itself
→ Preview every file that will be written before exporting
→ Export → writes agent .md files + orchestrator SKILL.md to ~/.claude/
→ Invoke the workflow by name in Claude Code
The exporter writes directly to ~/.claude/ (user-scoped) or .claude/ inside any project directory (project-scoped, version-controllable). Conflict detection ensures it never touches files it doesn't own.
Drag an existing agent from the sidebar ( ~/.claude/agents/ ) onto the canvas to create a reference node — it points to that agent file by slug rather than duplicating it. Drag from "New / Blank Agent" to create a bespoke node — the exporter generates a new agent file for it.
Connect nodes by dragging between handles. Each connection becomes a handoff with a trigger description and optional context artifacts (files, text, JSON) passed between agents. Mark any node as a terminal ( Complete , Escalated , or Aborted ) to define workflow end states.
Edit any node's completion criteria, tool access, skills, and system prompt in the Node Panel . The first node can also have a start trigger describing what initiates the workflow.
Real-time validation surfaces duplicate slugs, empty names, disconnected nodes, and missing completion criteria immediately in the top bar — before you export.
Use Generate agent or Generate skill in the sidebar to draft reusable Claude Code assets from a plain-English description. CWC gives you an editable spec first, then writes the file into ~/.claude/agents/ or ~/.claude/skills/ .
Click Export in the top bar. Choose a target directory ( ~/.claude/ or any project's .claude/ ). Review a preview of every file that will be written. Confirm to write.
Bespoke nodes → writes an agent .md file with frontmatter (name, description, color, model, tools), system prompt, completion criteria, skill references, and an ownership comment.
Reference nodes → writes nothing — the exportedSlug is set to the existing agent's slug so the orchestrator routes to it directly.
Workflow skill → generates an orchestrator skill at .claude/skills/<workflow-slug>/SKILL.md with disable-model-invocation: true . The orchestrator body is produced by BFS-traversing the node/edge graph into natural-language steps.
Rename handling → if a node was renamed, the old owned file is deleted and the new one is written.
Conflict detection → every file carries an ownership HTML comment. Before overwriting or deleting, the exporter verifies ownership — it never touches files created by other workflows or by hand.
From any Claude Code session, invoke the workflow by its skill name:
/workflow-name
The orchestrator skill delegates every implementation step to sub-agents via the Agent tool. Each step references an agent by name; Claude Code resolves it to the agent's .md file and loads its system prompt, tools, and completion criteria.
Open a workflow's Automate mode to add schedules and webhooks:
Cron — use the schedule builder or enter a custom cron expression (e.g. 0 9 * * 1-5 for weekdays at 9 am).
Webhook — CWC generates an inbound local URL; send an HTTP POST to fire the workflow.
Working directory / targets — choose where the automation runs, including optional additional repos for fan-out.
Isolation — use a git worktree for an isolated branch, or run in-place when you explicitly want the current checkout.
Precondition — a shell command that must succeed before CWC starts the run.
Setup command — a shell command CWC runs after the run starts, before Claude begins.
Add a gate node (drag from the "Gate" section of the sidebar) at any point in the workflow. When the run reaches a gate it:
Commits all changes to a cwc/<runId> branch and pauses.
Posts a diff of the working branch to the inbox in the Run panel.
Waits — the reviewer reads the diff, writes an optional note, and clicks Approve or Reject .
On approval, the run resumes in the same Claude Code session from the gate point.
The Run panel header has an Automations toggle that globally suspends all scheduled runs without disarming triggers, and a ⚙ gear that opens notification settings (macOS banners and/or a webhook URL).
Click Detect automations on the Home dashboard to scan your local Claude Code history for repeated work. CWC parses local transcript files, builds compact digests, asks Claude to cluster recurring tasks, and shows candidates with evidence, confidence, observed steps, and a suggested trigger.
Click Generate workflow on a candidate to promote it into a real .cwc workflow. CWC looks for matching local skills and existing agents, asks Claude to compose the workflow, validates the generated graph, seeds disabled schedule triggers when appropriate, and opens the workflow for review.
POST /api/export/delete scans every exported file, checks its ownership comment, and only removes files owned by the current workflow. Reference nodes have nothing to delete — they didn't write any files.
Visual canvas — React Flow with background grid, minimap, zoom controls, and drag-to-connect
Theme toggle — switch between light and dark mode from the Home dashboard or workflow header
Left sidebar — My Agents (searchable, draggable from ~/.claude/agents/ ) and Skills (searchable, draggable onto selected nodes)
Generate agent / skill — draft new reusable Claude Code assets from plain English, refine the spec, then save to ~/.claude/
Right panels — Node Editor (name, description, criteria, tools, skills, system prompt, terminal type) and Edge Editor (trigger, label, context artifacts)
Export modal — target selection, full file preview, warning display before writing anything
Auto-save — 500ms debounced save to ~/.cwc/workflows/ , no manual saving needed
Recent files — home screen shows last 10 workflows, persisted to ~/.cwc/recents.json
Markdown preview — click any agent or skill card to view its source file
Open in editor — view any agent or skill file in your system editor
Claude Code detection — warns on startup if ~/.claude/ is missing
▶ Test Run — launch an exported workflow headlessly from the UI ( --permission-mode bypassPermissions , user-chosen working directory, worktree or in-place isolation) and stop it mid-run
Live run view — the active node pulses on the canvas, completed nodes get a check, and events stream into a timeline panel
Run history — every run of every exported workflow (started from CWC or any terminal) persists to ~/.cwc/runs/ with status, duration, source, and cost
Automate mode — attach cron schedules or webhook URLs to a workflow, choose targets/isolation, add preconditions/setup commands, and arm trusted triggers
Approval gates — insert a gate node into any workflow; when reached the run pauses and posts a diff of its working branch, a reviewer approves or rejects from the inbox (or terminal), the run resumes on the same session
Isolated runs — Test Run (and scheduler-fired runs) create a git worktree on a cwc/<runId> branch so the main checkout is always untouched; the worktree is removed after the run completes
Notifications — macOS banner + optional webhook on run complete, gate pause, and approval request; configured from the settings gear in the Run panel
Global pause — one toggle in the Run panel suspends all scheduled automation runs without disarming triggers
Detect automations — scans local Claude Code transcripts, clusters repeated work, streams progress, suggests automations, and promotes a candidate into a .cwc workflow with matching skills/agents reused when possible
Client (React + React Flow) Server (Express :3579)
┌─────────────────────────┐ ┌─────────────────────┐
│ TemplatePicker │ ──► │ /api/workflows │
│ TopBar │ ◄── │ /api/recents │
│ Sidebar (Agents/Skills) │ ──► │ /api/agents │
│ Canvas (React Flow) │ ──► │ /api/skills │
│ NodePanel / EdgePanel │ ──► │ /api/export │
│ ExportFlow (modal) │ ──► │ /api/export/preview │
│ RunModal / RunPanel │ ──► │ /api/export/delete │
│ useWorkflow (reducer) │ │ /api/runs (+SSE) │
│ useAutoSave (debounced) │ │ /api/automations │
│ useRunEvents (SSE) │ ◄── │ /api/triggers │
│ Detect autom

[truncated]
