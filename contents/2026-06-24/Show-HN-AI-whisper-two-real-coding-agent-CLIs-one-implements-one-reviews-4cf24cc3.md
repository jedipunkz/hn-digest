---
source: "https://github.com/ai-creed/ai-whisper"
hn_url: "https://news.ycombinator.com/item?id=48661711"
title: "Show HN: AI-whisper – two real coding agent CLIs, one implements, one reviews"
article_title: "GitHub - ai-creed/ai-whisper: Terminal-first relay for paired AI coding agents (Claude + Codex), driven by structured workflows. · GitHub"
author: "vuphanse"
captured_at: "2026-06-24T16:42:52Z"
capture_tool: "hn-digest"
hn_id: 48661711
score: 1
comments: 0
posted_at: "2026-06-24T15:49:37Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI-whisper – two real coding agent CLIs, one implements, one reviews

- HN: [48661711](https://news.ycombinator.com/item?id=48661711)
- Source: [github.com](https://github.com/ai-creed/ai-whisper)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T15:49:37Z

## Translation

タイトル: Show HN: AI-wisper – 2 つの実際のコーディング エージェント CLI、1 つは実装、1 つはレビュー
記事のタイトル: GitHub - ai-creed/ai-whisper: 構造化されたワークフローによって駆動される、ペアの AI コーディング エージェント (Claude + Codex) のターミナルファースト リレー。 · GitHub
説明: 構造化されたワークフローによって駆動される、ペアの AI コーディング エージェント (Claude + Codex) のターミナルファースト リレー。 - ai-クリード/ai-ウィスパー

記事本文:
GitHub - ai-creed/ai-whisper: 構造化されたワークフローによって駆動される、ペアの AI コーディング エージェント (Claude + Codex) のターミナルファースト リレー。 · GitHub
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
アイクリード
/
愛のささやき
公共
通知
通知設定を変更するにはサインインする必要があります
追加

イオンナビゲーションオプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
884 コミット 884 コミット .github/ workflows .github/ workflows docs docs packages scripts scripts test test .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 NOTICE README.md README.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml prettier.config.mjs prettier.config.mjs tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ai-whisper は、2 つのコーディング エージェント (Claude、Codex、ezio のいずれか 2 つをマウント) を 1 つのバトンの下で手動で行ったり来たりするターミナル ネイティブのペアにペアにします。そのため、一方のエージェントが実装し、もう一方がレビューします。また、構造化されたワークフローにより、人間が毎回子守をすることなく、完成したレビューされた成果物に至るまでのループが推進されます。
各エージェントを独自の端末にマウントします。各マウントは現在のシェルを要求し、実際のプロバイダー CLI を起動して、それをコラボにバインドします。
# ターミナル1
ささやきコラボマウントクロード
#ターミナル2
ウィスパーコラボマウントコーデックス
次に、どちらかのエージェントのセッション内から、仕様に基づいて構造化されたワークフローを開始します。わかりやすい言葉で質問するだけです。
docs/spec.md を使用してスペック駆動開発を実行する
そこから、ai-whisper はワークフローを自律的に実行します。
実装者/レビュー担当者の割り当て - ワークフローをトリガーしたエージェントが実装者となり、他のエージェントがレビュー担当者になります。 --implementer / --reviewer を渡して明示的に選択します。 (マウントされたセッション外でフラグなしで開始された場合、失敗します

デフォルトのペアリングに戻って警告します。) バトンがそれらの間で渡されます。一度にターンを所有できるのは 1 人だけです。
自律実行 — 実装者は実際のセッションで各ステップを実行し、結果を返します。 LLM 評価者は、成果物が要求を満たしているかどうかを判断します。
レビューループ — 作業がまだ十分ではない場合、レビュー担当者の発見はフォローアップの引き継ぎとして構成され、実装者が反復します。このループは、作業が承認されるか、ラウンドの予算がなくなるまで繰り返されます。
再開可能性 — ワークフローとチェーン状態は永続的です。ブローカーが再起動するか、その日は停止しても、最初からやり直すのではなく、回復して再接続します。
成果物 — コミットされたコードとレビュー証跡 (ステップごとの判定、ラウンド数) が得られ、ウィスパー コラボ ダッシュボードでいつでも検査できます。
本当の仕様主導型開発の実行: クロード (左) とコーデックス (中央) は独自にマウントされた環境で動作します。
ダッシュボード (右) ではバトンのハンドオフとフェーズごとの判定 (約 20 秒) が追跡されます。
静止画をクリックすると、プロジェクト ページで再生されます。
ai-whisper は、すでにコーディング エージェントに頼っており、その周りの構造をさらに強化したいエンジニア向けです。
すでにコーディング エージェントを頻繁に使用しており、そのうちの 2 つが相互にチェックできるようにしたいと考えています。
ターミナルファーストで作業し、エージェントが Web UI ではなく実際のターミナル セッションに存在するようにしたいと考えています。
マルチエージェントによるレビューが必要です。つまり、最初のモデルの出力を制御する 2 番目のモデルです。
1 回限りのプロンプトではなく、長期にわたる構造化されたワークフロー (仕様 → 計画 → 実装 → レビュー) を実行します。
すぐに答えが欲しいワンショットの「バイブコーディング」。
決して見ることのない目に見えないバックグラウンドの自動化。
コーディング エージェントに慣れていない人で、ガイド付きのハンドヘルド エクスペリエンスを求めています。
3 つのエージェント ( claude 、 codex 、および ezio ) のうちの任意の 2 つをペアにします。 ai-whisper は実際の Claude および Codex CLI を駆動します。

これら 2 つのうちのどちらを最初にマウントするかをインストールして認証します。 ezio はプロトコルネイティブで、ai-whisper が同梱されているため、個別の CLI は必要ありません。
Claude Code CLI — サインインした claude コマンド。
Codex CLI — サインインした codex コマンド。
ezio (オプション) — AI-Whisper にバンドルされています。ウィスパー コラボ マウント ezio でマウントします。別途インストールする必要はありません。
資格情報を持つ LLM エバリュエーター — ワークフローは資格情報によってゲートされ、資格情報なしでは開始されません。 「 エバリュエーターの構成 」を参照してください。
tmux (オプション) — Whisper Collab start 専用。両方のエージェントをペインに自動配置します。以下のマウント フローでは必要ありません。
プラットフォームのサポート: ai-whisper はターミナルネイティブで Unix 指向であり、インタラクティブな PTY セッションを駆動するため、macOS および Linux 上で実行されます。 Windows ではネイティブにサポートされていません。ウィスパー コラボのマウント/再接続には Unix tty-backed シェルが必要で、ここを指すエラーで終了します。 Windows では、WSL2 内で ai-whisper を実行します。Node、エージェント CLI、および AI-Whisper を WSL2 ディストリビューション内にインストールし、そこでコマンドを実行すると、すべてがそのまま動作します。
ai-whisper は各エージェントを完全自律モードで起動するため、リレーは無人でエージェントを駆動できます ( claude --dangerously-skip-permissions および codex --dangerously-bypass-approvals-and-sandbox )。マウントされたワークスペース内で、エージェントはプロンプトを表示せずにコマンドの読み取り、書き込み、および実行を行います。 2 人のエージェントが自律的に変更できるようにしたいコードを指定し、ダッシュボードで実行を監視し、自分が最後の門番であることを思い出してください。結果を出荷する前に確認してください。より深い根拠はコンセプトにあります。
npm install -g ai-whisper
または、リポジトリのチェックアウトから:
pnpmインストール
pnpm ビルド
バンドルされたエージェント スキルを 1 回インストールします (これにより、エージェントはワークフローを確認、開始、レポートできるようになります)。これにより、 ai-whisper-code-review もインストールされます。

スキル ワークフローのコードレビュー ハンドオフは、エージェントが作成したコードを評価するために使用され、スキル プラン実行ハンドオフは、実装者が承認されたプランをどのように実行するかを構造化するために使用されます。
ささやきスキルインストール
ワークフローには認証情報を持つ LLM エバリュエーターが必要です。これを実行する前に設定してください。 「 エバリュエーターの構成 」を参照してください。
次に、両方のエージェントをマウントし、ワークフローを実行します。
# ターミナル1
ささやきコラボマウントクロード
#ターミナル2
ウィスパーコラボマウントコーデックス
最初のマウントではコラボが作成され、ワークスペースのブローカー デーモンが起動されます。 2 番目は他のエージェントをバインドします。どちらのセッションからも、仕様ファイルまたは目標ファイルに対するワークフローを開始します。仕様については仕様駆動開発、オープンエンド目標については ralph ループ、さらに複雑なバグ修正と審議が行われます (「ワークフロー」を参照)。次のように実行する様子を観察してください。
ささやきコラボダッシュボード
ウィスパー コラボ ダッシュボード — 最近アクティブなコラボのライブ ウォールと実行ごとのインスペクター。
--all を追加すると、すべてのワークフロー実行が表示されます (コラボごとのマスキングなし)。と組み合わせる
--完全な実行台帳のすべてをウィンドウ表示します。
パッケージ化されたインストールではなく、リポジトリのチェックアウトから実行しますか?最初にビルド ( pnpm build ) し、CLI をノード package/cli/dist/bin/whisper.js ... として呼び出します。これらの例では、whisper ... と書かれています。
走行が短く停止しても、通常はエスカレートします。クラッシュはしません。評価者がフェーズを解決できない場合 (ラウンドの予算が使い果たされている、エージェントがブロックされていると報告している、または信頼度が低すぎる場合)、ループは停止し、ターンの所有権があなたに戻ります。これは失敗ではなく設計された終了です。実行状態は永続的であるため、ダッシュボードを読み、仕様を修正するかエージェントのブロックを解除し、ワークフロー再開 <id> をささやいて中断したところから再開します。エスカレーションとは、必要なときに人間に要求を求めるシステムのことです。それは正常なことであり、何かが壊れた兆候ではありません。
ai-wisper は swarm ではありません。

エージェントは一度に入力することはありません。作業は一度に 1 人のオーナーごとに 1 つのバトンで進められます。マウントされたセッションは、端末内の実際のエージェント セッション (Claude または Codex CLI、または ezio) であり、これらのセッションが真実の情報源です。自律性は監視されています。すべてのハンドオフ、判定、およびラウンドは検査可能であり、実行はファイア・アンド・フォーゲットではなく再開可能です。作業は、自由形式のチャットではなく、明示的なループと状態遷移という構造化されたワークフローとして編成されます。
Claude、Codex、ezio は現在サポートされており、そのうちの 2 つをマウントできます。このアーキテクチャは設計上プロバイダーに依存しないため、他のコーディング エージェント CLI を同じリレーの背後に追加できます。
完全なメンタル モデルについては、「概念」を参照してください。
ワークフロー — 4 つのワークフローを適切に使用する方法: spec-driven-development 、 ralph-loop 、 complex-bug-fixing 、および Council のいずれかを選択し、実行を推進する仕様、目標、バグ レポート、またはシードを作成します。
コンセプト — メンタル モデル: バトンの受け渡し、実際のマウントされたセッション、監視された自律性、ワークフロー優先の実行。
リレーとハンドオフのフロー - 完全なハンドオフ ステート マシン、キャプチャ ステータス テーブル、ホットキー リファレンス、ステップごとの判定、およびトラブルシューティング。
エバリュエーターの構成 — ワークフローを制御する LLM エバリュエーターに必要な資格情報とオプション。
レガシー接続モード — 保留された接続/採用フロー。履歴参照用に保持されます。
pnpmインストール
pnpmテスト
pnpm タイプチェック
pnpm リント
pnpm形式
ライセンス
Apache License 2.0 — 「ライセンス」と「通知」を参照してください。寄付金は、
開発者原産地証明書に基づいて受け入れられます (署名は次のようになります)
git commit -s )。
構造化されたワークフローによって駆動される、ペアの AI コーディング エージェント (Claude + Codex) のターミナルファースト リレー。
ai-creed.dev/projects/ai-whisper/
トピックス
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください

e 。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Terminal-first relay for paired AI coding agents (Claude + Codex), driven by structured workflows. - ai-creed/ai-whisper

GitHub - ai-creed/ai-whisper: Terminal-first relay for paired AI coding agents (Claude + Codex), driven by structured workflows. · GitHub
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
ai-creed
/
ai-whisper
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
884 Commits 884 Commits .github/ workflows .github/ workflows docs docs packages packages scripts scripts test test .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml prettier.config.mjs prettier.config.mjs tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
ai-whisper pairs two coding agents — mount any two of Claude, Codex, and ezio — into a terminal-native pair that hand work back and forth under a single baton, so one agent implements while the other reviews, and a structured workflow drives the loop to a finished, reviewed deliverable without a human babysitting every round.
Mount each agent in its own terminal. Each mount claims the current shell, launches the real provider CLI, and binds it to the collab:
# terminal 1
whisper collab mount claude
# terminal 2
whisper collab mount codex
Then, from inside either agent's session, kick off a structured workflow against a spec — just ask in plain language:
Run spec-driven-development using docs/spec.md
From there ai-whisper runs the workflow autonomously:
Implementer / reviewer assignment — the agent you trigger the workflow from becomes the implementer and the other agent becomes the reviewer; pass --implementer / --reviewer to choose explicitly. (Started outside a mounted session with no flags, it falls back to a default pairing and warns.) The baton passes between them; only one owns the turn at a time.
Autonomous execution — the implementer does each step in its real session and hands the result back. An LLM evaluator judges whether the deliverable meets the request.
Review loops — when work isn't good enough yet, the reviewer's findings are composed into a follow-up handoff and the implementer iterates. The loop repeats until the work is approved or the round budget is exhausted.
Resumability — workflow and chain state is durable. If the broker restarts or you stop for the day, you recover and reconnect rather than starting over.
Deliverables — you get committed code plus a review trail (per-step verdicts, round counts), inspectable at any time with whisper collab dashboard .
A real spec-driven-development run: Claude (left) and Codex (middle) work in their own mounted
sessions while the dashboard (right) tracks the baton handoffs and per-phase verdicts (~20s).
Click the still to watch it play on the project page.
ai-whisper is for engineers who already lean on coding agents and want more structure around them:
you already use coding agents heavily and want two of them to check each other.
you work terminal-first and want the agents to live in real terminal sessions, not a web UI.
you want multi-agent review — a second model gating the first model's output.
you run long, structured workflows (spec → plan → implement → review) rather than one-off prompts.
one-shot "vibe coding" where you just want a quick answer.
invisible background automation you never watch.
people new to coding agents looking for a guided, hand-holding experience.
You pair any two of three agents — claude , codex , and ezio . ai-whisper drives the real Claude and Codex CLIs, so install and authenticate whichever of those two you plan to mount first; ezio is protocol-native and ships with ai-whisper, so it needs no separate CLI.
Claude Code CLI — the claude command, signed in.
Codex CLI — the codex command, signed in.
ezio (optional) — bundled with ai-whisper; mount it with whisper collab mount ezio , no separate install.
An LLM evaluator with credentials — workflows are gated by it and refuse to start without it. See Evaluator configuration .
tmux (optional) — only for whisper collab start , which auto-arranges both agents into panes. The mount flow below does not need it.
Platform support: ai-whisper is terminal-native and Unix-oriented — it drives interactive PTY sessions, so it runs on macOS and Linux . It is not supported natively on Windows : whisper collab mount / reconnect require a Unix tty-backed shell and will exit with an error pointing here. On Windows, run ai-whisper inside WSL2 — install Node, your agent CLI, and ai-whisper inside the WSL2 distro and run the commands there, where everything works as-is.
ai-whisper launches each agent in full-autonomy mode so the relay can drive it unattended — claude --dangerously-skip-permissions and codex --dangerously-bypass-approvals-and-sandbox . Inside the mounted workspace the agents read, write, and run commands without prompting. Point it at code you're willing to let two agents change autonomously, watch the run on the dashboard, and remember you are the final gatekeeper — review the result before you ship it. The deeper rationale is in Concepts .
npm install -g ai-whisper
Or from a repo checkout:
pnpm install
pnpm build
Install the bundled agent skills once (they let the agents verify, kick off, and report on workflows). This also installs ai-whisper-code-review , the skill workflow code-review handoffs use to evaluate agent-written code, and ai-whisper-plan-execution , the skill plan-execution handoffs use to structure how the implementer executes an approved plan:
whisper skill install
Workflows require an LLM evaluator with credentials — set this up before running one. See Evaluator configuration .
Then mount both agents and run a workflow:
# terminal 1
whisper collab mount claude
# terminal 2
whisper collab mount codex
The first mount creates the collab and starts the broker daemon for the workspace; the second binds the other agent. From either session, start a workflow against a spec or goal file — spec-driven-development for a spec, ralph-loop for an open-ended goal, plus complex-bug-fixing and deliberation (see Workflows ). Watch it run with:
whisper collab dashboard
whisper collab dashboard — live wall of recently-active collabs + per-run inspector.
Add --all to show every workflow run (no per-collab masking); combine with
--window all for the full run ledger.
Running from a repo checkout instead of a packaged install? Build first ( pnpm build ) and invoke the CLI as node packages/cli/dist/bin/whisper.js ... wherever these examples say whisper ... .
A run that stops short usually escalates — it does not crash. When the evaluator can't resolve a phase (the round budget is spent, an agent reports it's blocked, or confidence stays too low), the loop halts and turn ownership returns to you. That's a designed exit, not a failure: run state is durable, so you read the dashboard, fix the spec or unblock the agent, and whisper workflow resume <id> to pick up where it left off. Escalation is the system asking for a human exactly when it should — seeing it is normal, not a sign something broke.
ai-whisper is not a swarm . The agents never type at once — work moves by a single baton, one owner at a time. Mounted sessions are real agent sessions in your terminal — Claude or Codex CLIs, or ezio — and those sessions are the source of truth. Autonomy is supervised: every handoff, verdict, and round is inspectable, and runs are resumable rather than fire-and-forget. Work is organized as structured workflows — explicit loops and state transitions, not a free-form chat.
Claude, Codex, and ezio are supported today — you mount any two of them; the architecture is provider-agnostic by design, so other coding-agent CLIs can be added behind the same relay.
For the full mental model, read Concepts .
Workflows — how to use the four workflows well: choosing between spec-driven-development , ralph-loop , complex-bug-fixing , and deliberation , and authoring the spec, goal, bug report, or seed that drives the run.
Concepts — the mental model: baton handoff, real mounted sessions, supervised autonomy, workflow-first execution.
Relay & handoff flows — the complete handoff state machine, capture-status table, hotkey reference, per-step verdicts, and troubleshooting.
Evaluator configuration — required credentials and options for the LLM evaluator that gates workflows.
Legacy attach mode — the shelved attach / adopt flows, kept for historical reference.
pnpm install
pnpm test
pnpm typecheck
pnpm lint
pnpm format
License
Apache License 2.0 — see LICENSE and NOTICE . Contributions are
accepted under the Developer Certificate of Origin (sign off with
git commit -s ).
Terminal-first relay for paired AI coding agents (Claude + Codex), driven by structured workflows.
ai-creed.dev/projects/ai-whisper/
Topics
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
