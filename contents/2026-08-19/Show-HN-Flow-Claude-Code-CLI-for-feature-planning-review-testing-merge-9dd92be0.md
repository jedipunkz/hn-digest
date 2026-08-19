---
source: "https://github.com/gavingolden/flow/"
hn_url: "https://news.ycombinator.com/item?id=49363138"
title: "Show HN: Flow – Claude Code CLI for feature planning → review/testing → merge"
article_title: "GitHub - gavingolden/flow · GitHub"
image: "https://opengraph.githubassets.com/2bc65c3d1522cdee1fa3ff4421d20368e13263ca27ca928316c3835e159a74e6/gavingolden/flow"
author: "gavinegolden"
captured_at: "2026-08-19T16:19:56Z"
capture_tool: "hn-digest"
hn_id: 49363138
score: 2
comments: 0
posted_at: "2026-08-19T15:47:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Flow – Claude Code CLI for feature planning → review/testing → merge

- HN: [49363138](https://news.ycombinator.com/item?id=49363138)
- Source: [github.com](https://github.com/gavingolden/flow/)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T15:47:38Z

## Translation

タイトル: HN の表示: フロー – 機能計画 → レビュー/テスト → マージのためのクロード コード CLI
記事タイトル: GitHub - gavingolden/flow · GitHub
説明: GitHub でアカウントを作成して、gavingolden/flow の開発に貢献します。
HN テキスト: Flow は、エピックを設計し、機能を迅速に出荷するための私のクロード コード スーパーバイザーです。これはそれ自体でブートストラップされているため、最近の PR を見て、それが何を生成するかを確認できます。使用法: > フロー機能作成 'CSV エクスポートの実装' --model opus --effort high • トリアージ → プラン → git worktree → コード → 検証 → CI → レビュー → マージ • プランの承認と主観的なマージ前検証 (ある場合) の場合のみ一時停止 > フロー epic create 'お金を稼ぐためのアプリの設計' --model haiku --effort low • (`feature` と同じフローですが、フェーズ依存関係マッピングを含むエピック PRD を生成するだけです) > フローepic run design-app-for-making-money • エピックの現在の状態を解決し、ブロックされていないフェーズのフィーチャー パイプラインを起動します > claude -p "このバグのリストを優先し、opus high のバンドルごとにフロー パイプラインを起動します: <バグ リスト>" • フロー パイプラインは新しいフロー パイプラインを作成できるため、ファンアウトは非常に強力です (例: 1 つのパイプラインが個別に処理する必要があるバグを発見した場合) > flow feature ls # すべてのアクティブなパイプラインを表示 > flow epic ls # すべてのエピックステータスを表示 統計: • 過去 4 か月で、厳しくテスト/レビューされた ~1000 個の PR をマージしました。tmux 統合を有効にし、まだ使用していない場合は tmux の使用方法を学ぶことを強くお勧めします。

記事本文:
GitHub - gavingolden/flow · GitHub
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
ギビングゴールデン
/
流れ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
465 コミット 465 コミット フォルダーとファイル
.flow .flow .github .github エージェント/ コア エージェント/ コア ビン ビン 完了 完了 ドキュメント ドキュメント リファレンス リファレンス スキル スキル テンプレート テンプレート .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json A

GENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.scripts.json tsconfig.scripts.json vitest.config.ts vitest.config.ts vitest.setup.ts vitest.setup.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
小さな変更を出荷する場合でも、計画→ブランチ→コード→テスト→PR→CI→レビュー→マージの午後丸々の時間がかかります。
フローは 1 つのコマンドからエンドツーエンドで変更を送信します。ユーザーが監視している間、または立ち去っている間、Claude Code スーパーバイザーがパイプライン全体を駆動します。
以下のトランスクリプトは説明用であり、正確な出力ではありません。これは実行の実際の冗長性ではなく、実行の形状を示します。
$フロー機能作成「CSVエクスポートを追加」
[計画] 機能が検出されました — 計画の草案作成、承認のために一時停止しています
計画: テストを含む CSV エクスポーターをフラグの後ろに追加します。
承認する？ > 承認されました
[実装] ワークツリーの準備ができました。編集を適用し、検証を実行します ... OK
[ci] PR #142 がオープンされました — 小切手を待っています ................... 緑色
[レビュー] マルチエージェントレビュー + Copilot ... 2 つの調査結果が修正されました
[ゲート] テスト ステップをすべてチェック → 自動マージ
統合されました
すべての実行は独自の git ワークツリーで動作し (チェックアウトには一切触れません)、機能作業の計画承認のために一度一時停止し、明確な終了状態 ( MERGED 、 GATED: <url> 、 NEEDS HUMAN: <reason> 、または canceled ) で終了します。 flow には、クロード コード プロジェクトが単独で使用できる厳選されたスキル ライブラリも同梱されています。
前提条件を確認してください。 git 、node / npm 、bun 、および認証された gh (GitHub CLI) が必要です。フローは PR を開き、CI をポーリングし、それを介してマージします。 tmux はオプションです。tmux ランチャーを選択した場合にのみ必要です。デフォルトのプレーンランチャーは自分のターミナルで実行されます。ターゲット プロジェクトは、GitHub リモートを備えた git リポジトリである必要があります。
gitクローンhtt

ps://github.com/gavingolden/flow ~ /code/flow
cd ~ /code/flow
npmインストール
バンビン/フローインストール
動作することを確認します。フロー ls を実行します。「コマンドが見つかりません」ではなく、空のパイプライン リストが出力されるはずです。最も一般的な失敗は、 ~/.local/bin が PATH 上にないことです。それを追加して新しいシェルを開きます。
何かを出荷する: GitHub が支援するプロジェクトに cd して実行します。
フロー機能作成「CSVエクスポートの追加」
flow install は、選択したモジュールのセット (パイプライン コアと選択したスタック/統合スキル) にシンボリックリンクを作成します。コアは常にインストールされます。それ以外のすべて (Svelte、Tailwind/shadcn、Supabase、Cloudflare Pages、GitHub Copilot レビュー、AI-Ultra 研究ツール) は、インタラクティブな Q&A を通じてオプトインされます。選択された各モジュールは、クロード コードのスキルディレクトリ プラグイン ルートとしても具体化され、シンボリックリンクの上に衝突防止の名前空間とモジュールごとのトークンコスト インベントリの基礎を築きます。詳細については docs/configuration.md を参照してください (スキルは今のところフラットのままなので、現在は Skills (0) / ~0 tok を報告します)。選択フラグ、アップグレード ( flow install --upgrade )、およびスタンドアロン スキル ホームについては、 docs/configuration.md で説明されています。セットアップの内部設定は CONTRIBUTING.md にあります。また、flow ls / flow version / flow install は、マネージド シンボリック リンクが見つからない、ぶら下がっている、または古くなった場合、または予期しないエントリ (フロー管理プラグイン ルート内の迷走ファイルまたはディレクトリ) または外部エントリ (フロー ソース ツリーの外側、フロー管理プラグイン ルート内で解決されているライブ bin/シンボリック リンク) が表示された場合、淡色のインストール ドリフトの問題行を表示します。 flow install --upgrade はシンボリック リンクの種類を修復し、外部エントリを自動的に削除します。予期しないエントリは手動で削除する必要があります。
フロー # (TTY 上) フロー スキルがロードされたインタラクティブなクロード セッション
フロー機能作成 " CSV エクスポートを追加 " # パイプラインを開始 (r

デフォルトでは端末内にあります）
flow ls # アクティブなパイプラインをリストします (支出に対して --cost を追加します)
フロー機能再開 add-csv-export # クラッシュまたは閉じたパイプラインを保存された状態から再起動します
フロー完了 add-csv-export # 完成したパイプラインを閉じる
フロー完了 --merged # スイープ マージ/キャンセルされたパイプライン
フロー ヘルプを実行すると、完全なコマンド リファレンス ( epic 、 config 、attach 、 completed 、 version 、およびすべてのフラグ) が表示されます。
プレーンシェルがデフォルトです。 flow feature create は、起動したどの端末でもクロード コードをフォアグラウンド プロセスとして実行します。ウィンドウ管理や、それを表示するための追加のコマンドは必要ありません。実行が終了状態に達するまで、端末を保持します。デフォルトでは、マージ ゲートがクリアされている場合、パイプラインは PR を自動マージします。 --no-auto-merge を渡すと、常にゲートで停止します。
ここは新しいですか？完全な初回実行のウォークスルー (実行の読み取り、再開、クリーンアップ) は docs/getting-started.md にあります。
パワーユーザー: tmux ランチャー
1 つの場所から複数のパイプラインを一度に実行する場合、またはパイプラインを開始して立ち去り、後で任意の場所から再接続する場合は、tmux ランチャーに手を伸ばします。 flow feature create --tmux "<description>" を使用して実行ごとにオプトインし、対話型インストールでフロー インストールが尋ねる tmux の質問に「はい」と答えるか、flow config launcher set tmux を使用してデフォルトとして設定します。エージェントまたはスクリプト駆動のフロー機能作成 (TTY のないもの) は、 --tmux を渡す必要があります。デフォルトのプレーン ランチャーは、設計により非対話型起動を拒否します (PR #457 以降)。
tmux ランチャーの下で、最初のフロー機能作成により tmux ウィンドウでパイプラインが開始されますが、そこにはドロップされません。 flow Attach (引数なし) を実行してフロー セッションにポップするか、flow Attach <name> (エイリアス flow a <name> ) を実行して特定のセッションにジャンプします。実行を停止せずに離れるには、Ctrl-b d で切り離します。パイプラインは維持されます。

実行すると、 flowattach で戻ってきます。
さまざまなクロード モデルをさまざまなパイプライン フェーズにルーティングできます (計画時には高価なモデル、検証時には安価なモデル)。フラグ、構成キー、および優先順位ルールについては、docs/configuration.md#per-phase-models を参照してください。
フローはプッシュするたびにフロー事前コミット検証ゲートを実行します。単一パッケージのリポジトリとモノリポジトリは設定なしで動作します。差分からスコープを自動検出し、宣言された npm 実行スクリプト (さらに apps/<pkg>/ および package/<pkg>/ 内のパッケージごとのスクリプト) を実行します。従来とは異なるレイアウトの場合は、.flow/pre-commit.json エスケープ ハッチを追加します。詳細は、 ## Consumer-repo ノートの下の AGENTS.md にあります。
semgrep はオプションのツールです。これが PATH 上にある場合、 /flow-pr-review は静的分析セキュリティ レンズを実行します。これがないと、そのレンズはスキップされます (他の機能はすべて実行されます)。
あなたが欲しいのは
読む
最初のパイプラインのステップバイステップ
docs/getting-started.md
構成、モデル、インストールフラグ、アップグレード
docs/configuration.md
フローそのものに取り組む
貢献.md
スーパーバイザースキルそのもの
スキル/パイプライン/フローパイプライン/SKILL.md
フローに取り組むエージェントのプロジェクト ルール
エージェント.md
ライセンス
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to gavingolden/flow development by creating an account on GitHub.

Flow is my claude code supervisor for designing epics and shipping features really quickly. It was bootstrapped with itself so you can look at recent PRs to see what it produces. Usage: > flow feature create 'implement csv exports' --model opus --effort high • triage → plan → git worktree → code → verify → CI → review → merge • pauses only for plan approval and subjective pre-merge validation (if any) > flow epic create 'design app for making money' --model haiku --effort low • (same flow as `feature` but it only produces an epic PRD with phase dependency mappings) > flow epic run design-app-for-making-money • resolves current state of the epic and launches feature pipelines for unblocked phases > claude -p "Triage this list of bugs and launch a flow pipeline for each bundle with opus high: <bug list>" • a flow pipeline can create new flow pipelines so fan-out is really powerful (eg if one pipeline discovers a bug that should be handled separately) > flow feature ls # show all active pipelines > flow epic ls # show all epic statuses Stats: • in the last 4 months I've merged ~1000 heavily tested/reviewed PRs I very strongly recommend enabling the tmux integration and learning how to use tmux if you don't already.

GitHub - gavingolden/flow · GitHub
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
gavingolden
/
flow
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
465 Commits 465 Commits Folders and files
.flow .flow .github .github agents/ core agents/ core bin bin completions completions docs docs references references skills skills templates templates .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.scripts.json tsconfig.scripts.json vitest.config.ts vitest.config.ts vitest.setup.ts vitest.setup.ts View all files Repository files navigation
Shipping a small change still costs you a whole afternoon of plan → branch → code → test → PR → CI → review → merge.
flow ships a change end-to-end from one command — a Claude Code supervisor drives the entire pipeline while you watch or walk away.
The transcript below is illustrative — not exact output ; it shows the shape of a run, not its real verbosity.
$ flow feature create "add CSV export"
[plan] feature detected — drafting plan, pausing for approval
Plan: add a CSV exporter behind a flag, with tests.
approve? > approved
[implement] worktree ready, applying edits, running verify ... ok
[ci] PR #142 opened — waiting for checks ............... green
[review] multi-agent review + Copilot ... 2 findings fixed
[gate] Test Steps all checked → auto-merge
MERGED
Every run works in its own git worktree (your checkout is never touched), pauses once for plan approval on feature work, and ends in a clear terminal state: MERGED , GATED: <url> , NEEDS HUMAN: <reason> , or cancelled . flow also ships a curated skill library that any Claude Code project can use on its own.
Check the prerequisites. You need git , node / npm , bun , and an authenticated gh (GitHub CLI) — flow opens the PR, polls CI, and merges through it. tmux is optional : only needed if you opt into the tmux launcher; the default plain launcher runs in your own terminal. Your target project must be a git repo with a GitHub remote.
git clone https://github.com/gavingolden/flow ~ /code/flow
cd ~ /code/flow
npm install
bun bin/flow install
Verify it worked: run flow ls — it should print an empty pipeline list, not "command not found". The most common failure is ~/.local/bin not being on your PATH ; add it and open a fresh shell.
Ship something: cd into any GitHub-backed project and run
flow feature create " add CSV export "
flow install symlinks a selected set of modules — the pipeline core plus whichever stack/integration skills you pick. core is always installed; everything else (Svelte, Tailwind/shadcn, Supabase, Cloudflare Pages, GitHub Copilot review, AI-Ultra research tooling) is opt-in via an interactive Q&A. Each selected module is also materialized as a Claude Code skills-dir plugin root, laying the groundwork for collision-proof namespacing and a per-module token-cost inventory on top of the symlinks — see docs/configuration.md for the detail (skills stay flat for now, so it currently reports Skills (0) / ~0 tok ). Selection flags, upgrades ( flow install --upgrade ), and the standalone skills home are covered in docs/configuration.md ; setup internals live in CONTRIBUTING.md . flow ls / flow version / flow install also surface a dimmed install drift issue(s) line if any managed symlink goes missing, dangling, or stale, or an unexpected entry (a stray file or directory inside a flow-managed plugin root) or a foreign entry (a live bin/ symlink resolving outside the flow source tree, inside a flow-managed plugin root) turns up — flow install --upgrade repairs the symlink kinds and automatically removes a foreign entry, while an unexpected entry must be removed by hand.
flow # (on a TTY) interactive Claude session with flow skills loaded
flow feature create " add CSV export " # start a pipeline (runs in your terminal by default)
flow ls # list active pipelines (add --cost for spend)
flow feature resume add-csv-export # re-launch a crashed or closed pipeline from saved state
flow done add-csv-export # close a finished pipeline
flow done --merged # sweep merged/cancelled pipelines
Run flow help for the full command reference ( epic , config , attach , completion , version , and every flag).
The plain shell is the default. flow feature create runs Claude Code as a foreground process in whatever terminal you launched it from — no window management, no extra command to see it. It holds your terminal until the run reaches a terminal state. By default a pipeline auto-merges its PR when the merge gate is clear; pass --no-auto-merge to always stop at the gate.
New here? The full first-run walkthrough — reading a run, resuming, cleaning up — is at docs/getting-started.md .
Power users: the tmux launcher
Reach for the tmux launcher when you want to run several pipelines at once from one place, or start a pipeline and walk away, re-attaching from anywhere later. Opt in per run with flow feature create --tmux "<description>" , answer "yes" to the tmux question flow install asks on an interactive install, or set it as your default with flow config launcher set tmux . Agent- or script-driven flow feature create (anything without a TTY) must pass --tmux : the default plain launcher refuses non-interactive launches by design (since PR #457).
Under the tmux launcher, your first flow feature create starts the pipeline in a tmux window but doesn't drop you into it — run flow attach (no args) to pop into the flow session, or flow attach <name> (alias flow a <name> ) to jump to a specific one. To step away without stopping the run, detach with Ctrl-b d — the pipeline keeps running, and you come back with flow attach .
You can route different Claude models to different pipeline phases (an expensive model on planning, a cheap one on verify). See docs/configuration.md#per-phase-models for the flags, config keys, and precedence rules.
flow runs the flow-pre-commit verify gate before every push. Single-package repos and monorepos work with zero config — it auto-detects scope from the diff and runs your declared npm run scripts (plus per-package scripts in apps/<pkg>/ and packages/<pkg>/ ). For a non-conventional layout, drop in a .flow/pre-commit.json escape hatch. Full detail is in AGENTS.md under ## Consumer-repo notes .
semgrep is an optional tool: when it is on PATH , /flow-pr-review runs the static-analysis security lens; without it, that lens is skipped (everything else still runs).
You want
Read
Your first pipeline, step by step
docs/getting-started.md
Config, models, install flags, upgrades
docs/configuration.md
Working on flow itself
CONTRIBUTING.md
The supervisor skill itself
skills/pipeline/flow-pipeline/SKILL.md
Project rules for agents working on flow
AGENTS.md
License
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
