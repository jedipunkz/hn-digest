---
source: "https://github.com/simke9445/agentlocks"
hn_url: "https://news.ycombinator.com/item?id=48404759"
title: "Agentlocks – advisory file locks for Codex and Claude Code in one worktree"
article_title: "GitHub - simke9445/agentlocks: Advisory file locks so multiple AI coding agents can share one Git worktree. · GitHub"
author: "simke9445"
captured_at: "2026-06-04T21:36:08Z"
capture_tool: "hn-digest"
hn_id: 48404759
score: 2
comments: 0
posted_at: "2026-06-04T21:17:56Z"
tags:
  - hacker-news
  - translated
---

# Agentlocks – advisory file locks for Codex and Claude Code in one worktree

- HN: [48404759](https://news.ycombinator.com/item?id=48404759)
- Source: [github.com](https://github.com/simke9445/agentlocks)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T21:17:56Z

## Translation

タイトル: Agentlocks – 1 つのワークツリー内の Codex と Claude Code の勧告ファイル ロック
記事のタイトル: GitHub - simke9445/agentlocks: アドバイザリー ファイル ロックにより、複数の AI コーディング エージェントが 1 つの Git ワークツリーを共有できます。 · GitHub
説明: アドバイザリ ファイル ロックにより、複数の AI コーディング エージェントが 1 つの Git ワークツリーを共有できます。 - simke9445/エージェントロック

記事本文:
GitHub - simke9445/agentlocks: アドバイザリー ファイル ロックにより、複数の AI コーディング エージェントが 1 つの Git ワークツリーを共有できます。 · GitHub
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
シムケ9445
/
エージェントロック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

イオン
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
150 コミット 150 コミット .claude/ フック .claude/ フック .codex/ ルール .codex/ ルール .cursor .cursor .gemini .gemini .github .github 分析 分析 アセット アセット bin bin docs docs スクリプト スクリプト src src テスト テスト .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md biome.json biome.json bun.lock bun.lock bunfig.toml bunfig.toml package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
アドバイザリー ファイル ロックにより、複数の AI コーディング エージェントが互いに干渉することなく 1 つの Git ワークツリーを共有できるようになります。
同じリポジトリ内で 2 つのコーディング エージェント (またはサブエージェントを持つ 1 つのエージェント) を実行すると、それらが開始されます。
お互いにつまずきます: 2 人が同じファイルを編集し、黙って作業を上書きします。
auth.ts " メモは古くなって決してクリアされず、2 つの git add 実行が同じインデックスに対して競合します。
Agentlock は、これらの衝突ポイントを明示的なスクリプト可能なリースに変換します。
Agentlocks は勧告です。書き込み前にロックをチェックするエージェントを調整します。重なり合うのです
表示可能でスクリプト可能ですが、プロトコルを無視するプロセスは停止しません。
これはエージェントネイティブです。ID はハーネスから取得されるため、管理する ID はありません。
状態レポート コマンドは JSON を話します。エラーには正確な修正の名前が付けられます。そして契約書はエージェントに伝えます
次に何を実行するか。デーモンもデータベースもホスト型サービスもありません。 .agentlocks/locks/ 下のファイルのみ。
npm install -g Agentlocks # または: bun install -g Agentlocks
エージェントロック --ヘルプ
npm install -g はマシン上に Bun が存在しない場合でも機能します。エージェントロックは 1 つの小さなノード バンドルとして出荷されます。

それ
独自のノードで実行されます。プラットフォームごとのバイナリや埋め込みランタイムはありません。ノード 22.18 以上が必要です。
TypeScriptのagentlocks.config.tsをネイティブにロードする最初のリリース。 (パンは構築するためにのみ必要です
ソースから。) 次に、調整する各リポジトリ内で、agentlock init を 1 回実行します。
公開された npm パッケージは、SLSA 来歴証明書を含む 5 ファイルの tarball であり、ランタイムはありません
依存関係。リリース自動化は、Linux glibc、Linux musl、
macOS、Windows、および Node 22.18 / 22.22 / 24.16 に準拠。
# Codex または Claude Code 内では、エージェントの ID が自動的に検出されます。
lock= $( エージェントロックは ' src/auth.ts ' --reason " リファクタリング ログイン " --id-only を取得します)
# ... エージェントは src/auth.ts を編集し、テストを実行します ...
エージェントロック リリース " $lock " --id-only
リースが保持されている間に「src/auth.ts」を取得しようとする 2 番目のエージェントは、クリーンなライセンスを取得します。
競合: ブロックしている所有者、その理由、および次に実行する正確なコマンド (代わりに)
サイレント上書き。
ほとんどの CLI は人間向けに設計されており、エージェントによってのみ許容されます。 Agentlock はそれを逆転させます。毎
表面はエージェント人間工学バーに固定されており、各主張を自分で検証できます。
タイプミスによってエージェントが行き詰まるわけではありません。それは次のように教えています：
$ エージェントロックステータス --jason
エージェントロック エラー: 不明なオプション '--jason'
(--json という意味ですか?)
次: エージェントロックステータス --json
ID はセットアップやフラグなしでそのまま機能します。
$ エージェントロックは --json を識別します
{"kind":"identified","exit_code":0,"agent_id":"codex:0fd188d5-...","source":"harness:codex:CODEX_THREAD_ID","harness":"codex","harness_scope":"agent","reliable":true,"mine_supported":true}
Agentlocks は、独自のハーネス統合も実現します。agentlocks init --harness claude-code
Claude Code PreToolUse フック (および Codex と同等の --harness codex) をインストールします。
広告

git commit ツール呼び出しの前にバイザリ git verify を実行し、ステージングされているがロックされていないパスを明らかにする
コミットをブロックしたり、git 設定を変更したりすることはありません。
リポジトリの同時作業は、予測可能な 3 つの場所で失敗します。 Agentlock はそれぞれを明示的にします。
所有され、解析可能:
Agentlocks は助言的なものであり、使用に同意するエージェントを調整します。それは止まらない
プロトコルを無視するエディター、シェル コマンド、または Git 操作。これが必要な理由でもあります。
デーモン、権限、ロックを保持するバックグラウンド プロセスはありません。
アプローチ
に適しています
共有ワークツリーでは不十分な点
チャットまたは問題の手動メモ
非公式な調整と意図
リースなし、所有者チェックなし、解析可能なステータスなし、ステージング前に忘れられやすい
シェルスクリプト
1 つのリポジトリに関するローカル規約
通常、競合セマンティクス、古いセッション、JSON コントラクト、Git インデックス ロックが見逃されます。
群れ
1 台のマシン上のプロセスレベルの重要なセクション
リポジトリ リソース レジストリではありません。リソースインベントリ、所有者のメタデータ、インストールガイダンス、またはエージェントドキュメントはありません
Git ネイティブ フック (コミット前)
コミットタイムポリシーチェック
編集の重複を防ぐには遅すぎます。フックはワーカー間で git add を調整せず、単一の core.hooksPath が husky/lefthook と衝突します。
ホスト型ロックサービス
クロスマシン調整
サービス、資格情報、ネットワーク アクセス、および運用の所有権が必要です
エージェントロック
1 つのリポジトリ ワークツリー内のローカル エージェント
助言のみ。 git verify と、git commit ツール呼び出しの前にそれを実行するオプトアウト型の PreToolUse バックストップ (クロード コード + コードックス) を同梱します。 git フックをインストールせず、git を再構成することはありません
クイックデモ
このデモは Codex または Claude Code 内で実行されます。アクティブなハーネス ID が検出されました
自動的に、agentlocks init --harness claude-code が指定されている場合に、Claude Code サブエージェントを含めます。
PRをインストールしました

オジェクトフック。
HOST= " $( mktemp -d ) "
cd " $HOST "
git init -q
printf ' console.log("hello")\n ' > app.ts
printf ' # デモ ホスト リポジトリ\n ' > AGENTS.md
printf ' {"scripts":{}}\n ' > package.json
# 初期ドリフトが見つかった場合は 1 を終了することが期待されます。ファイルには書き込みません。
エージェントロック init --check --json ||本当の
エージェントロックの初期化
エージェントロックは --json を識別します
file_lock= " $( エージェントロックは ' app.ts ' を取得します --reason " アプリを編集します " --id-only ) "
エージェントロック展開 ' README.md ' --lock " $file_lock " --id-only
エージェントロック更新 " $file_lock " --id-only
# git begin --id-only は、ロック ID、次にフェンス トークンの 2 行を出力します。両方読んでください。
{ git_lock を読み取ります ; git_token を読み取ります。 } < <(agentlocks git begin --refresh-lock " $file_lock " --reason " commit Demon " --id-only )
エージェントロックのステータス --json
エージェントロック git end " $git_lock " --git-token " $git_token " --release-lock " $file_lock " --id-only
エージェントロック プルーン --dry-run --json
エージェントロックドクター --json
予想される形状:
identify は、CLAUDE_CODE_SESSION_ID、CODEX_THREAD_ID、AGENTLOCKS_HARNESS_AGENT_ID などのハーネス ソースを示します。
status には、ファイル ロックと @git/index ロックが保持されている間に 2 つのロックが表示されます。
git end は、解放された git ロック ID とファイル ロック ID を出力します。
prune --dry-run は、新しいリポジトリで pruned_count 0 を報告します。
init が完了した後、医師は ok true と報告します。
ホストのセットアップ
Agentlocks init は冪等であり、ホスト リポジトリ サポート ファイルを書き込みます。リポジトリごとに 1 回実行します。
--harness claude-code (または --harness codex ) を追加して、PreToolUse フックもインストールします。
エージェントロック init --check --json || true # 変更をプレビューします。ドリフト時に 1 を終了し、何も書き込みません
エージェントロックの初期化
# Claude コード: .claude PreToolUse フックもインストールします。
エージェントロック init --check --harness クロード コード --json ||本当の
エージェントロック init --harness クロード コード
init は以下を作成または更新できます。
コミットフックのバックストップは

アドバイザリー: git コミットの前に、ステージングされているがロックされていないパスが表示されます。
tool-call を実行し、 commit をブロックすることはありません。これはデフォルトでインストールされます。 Agentlocks init --no-commit-hook は、元の ID インジェクション専用の Claude フックを保持し、Codex フックをスキップします。
以下の手順は、上記のグローバル インストールとサポートされているエージェント ハーネスを前提としています。コーデックスとクロード
コードの識別は自動的に行われます。
ホスト リポジトリを初期化します (AGENTS.md 命令ブロックを書き込みます)。
エージェントロック init --check --json ||本当の
エージェントロックの初期化
# Claude コード: .claude PreToolUse フックもインストールします。
エージェントロック init --check --harness クロード コード --json ||本当の
エージェントロック init --harness クロード コード
検出されたエージェント ID を検査します。
編集する前に最も狭いロックを取得します。
エージェントロック取得 ' src/index.ts ' ' testing/cli.test.ts ' --reason " CLI ディスパッチの変更 " --id-only
別のファイルに触れる前に展開してください。
エージェントロック展開 ' src/config.ts ' --lock < lock_id > --id-only
長い編集バッチの前、テスト後、ステージングの前に更新してください。
エージェントロック更新 < lock_id > --id-only
共有 Git インデックスを調整します。 git begin --id-only は 2 行を出力します: git ロック ID、次に
git end が再チェックするフェンス トークン (コミット中にリースが再利用された場合、出口 3 で中止されます)。
{ git_lock を読み取ります ; git_token を読み取ります。 } < <(agentlocks git begin --refresh-lock < lock_id > --reason " commit Agentlocks change " --id-only )
git add <locked_paths>
gitコミット
エージェントロック git end " $git_lock " --git-token " $git_token " --release-lock < lock_id > --id-only
または、agentlocks commit にすべて (ロック、ステージ、コミット、フェンス、リリース) を 1 つのコマンドで実行させます。
raw git commit の前に、サニティチェックカバレッジ (推奨; 決してブロックしない):
Agentlocks git verify --json # ステージングされているがロック解除されているパスをリストします
これは、オプトアウト PreToolUse コミットフック バックストップが自動的に実行するものです

厳密に。もしも負けてしまったら
ロック ID (例: コンテキストの圧縮後)、agentlocks status --mine で回復し、
エージェントロックリリース --mine / エージェントロックリフレッシュ --mine 。
エージェントロック機能 --json はコマンド メタデータ、位置、JSON の信頼できる情報源です
シェイプ参照、--id-only 回線契約、フラグ、終了コード、デフォルト TTL、エージェント ID 検出、
そして次のコマンド。
すべてのリソースを引用符で囲んで、シェルがそれを文字通り Agentlocks に渡します。正確なパスとグロブパターン
同じ位置リストを共有します。 CLI は、引用符で囲まれたこれらの引数を生のリソース仕様として扱います。 JSON
ロック レコードは、各リースに格納されている正規化された {kind,value} エントリのリソースを使用します。
エージェントロックは「a.ts」「b.ts」「src/**/*.ts」を取得します --reason "編集ファイル"
エージェントロック取得 ' 分析/トレンドベースライン/** ' --reason " フェーズ 0 アーティファクトの作成 "
コマンド
目的
主要なフラグ
ノートを出力する
[リソース...]を取得する
引用符で囲まれたリポジトリ相対パスまたはグロブ パターンのロックを取得する
--reason 、 --ttl-ms 、 --reclaim 、 --agent-id 、 --json 、 --id-only 、 --verbose
--reason と少なくとも 1 つのリソースが必要です。 glob のようなリソースは * 、? から推論されます。 、またはブラケットクラス
Expand [リソース...] --lock <id>
引用されたリポジトリ相対リソースを追加する

[切り捨てられた]

## Original Extract

Advisory file locks so multiple AI coding agents can share one Git worktree. - simke9445/agentlocks

GitHub - simke9445/agentlocks: Advisory file locks so multiple AI coding agents can share one Git worktree. · GitHub
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
simke9445
/
agentlocks
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
150 Commits 150 Commits .claude/ hooks .claude/ hooks .codex/ rules .codex/ rules .cursor .cursor .gemini .gemini .github .github analyses analyses assets assets bin bin docs docs scripts scripts src src tests tests .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md biome.json biome.json bun.lock bun.lock bunfig.toml bunfig.toml package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Advisory file locks that let multiple AI coding agents share one Git worktree without clobbering each other.
Run two coding agents in the same repository (or one agent with subagents) and they start
tripping over each other: two edit the same file and silently overwrite work, a "working on
auth.ts " note goes stale and never clears, and two git add runs race for the same index.
Agentlocks turns those collision points into explicit, scriptable leases.
Agentlocks is advisory: it coordinates agents that check the lock before writing. It makes overlaps
visible and scriptable, but it does not stop a process that ignores the protocol.
It's agent-native : identity comes from the harness, so there are no ids to manage;
state-reporting commands speak JSON; errors name the exact fix; and the contract tells the agent
what to run next. No daemon, no database, no hosted service. Just files under .agentlocks/locks/ .
npm install -g agentlocks # or: bun install -g agentlocks
agentlocks --help
npm install -g works with no Bun on the machine: Agentlocks ships as one small Node bundle that
runs under your own Node — no per-platform binary, no embedded runtime. It needs Node >= 22.18, the
first release that loads a TypeScript agentlocks.config.ts natively. (Bun is only needed to build
from source.) Then run agentlocks init once inside each repo you want to coordinate.
The published npm package is a 5-file tarball with SLSA provenance attestations and no runtime
dependencies. Release automation publishes the exact tarball that passed Linux glibc, Linux musl,
macOS, Windows, and Node 22.18 / 22.22 / 24.16 conformance.
# Inside Codex or Claude Code, the agent's identity is detected automatically.
lock= $( agentlocks acquire ' src/auth.ts ' --reason " refactor login " --id-only )
# ... the agent edits src/auth.ts, runs tests ...
agentlocks release " $lock " --id-only
A second agent that tries to acquire 'src/auth.ts' while that lease is held gets a clean
conflict: the blocking owner, its reason, and the exact command to run next, instead of a
silent overwrite.
Most CLIs are designed for humans and merely tolerated by agents. Agentlocks inverts that. Every
surface is held to an agent-ergonomics bar, and you can verify each claim yourself:
A typo doesn't dead-end the agent. It teaches:
$ agentlocks status --jason
agentlocks error: unknown option '--jason'
(Did you mean --json?)
next: agentlocks status --json
Identity just works, with no setup and no flags:
$ agentlocks identify --json
{"kind":"identified","exit_code":0,"agent_id":"codex:0fd188d5-...","source":"harness:codex:CODEX_THREAD_ID","harness":"codex","harness_scope":"agent","reliable":true,"mine_supported":true}
Agentlocks even brings its own harness integration: agentlocks init --harness claude-code
installs a Claude Code PreToolUse hook (and --harness codex the Codex equivalent) that runs
an advisory git verify before a git commit tool-call, surfacing staged-but-unlocked paths
without ever blocking the commit or touching your git config.
Concurrent repository work fails in three predictable places. Agentlocks makes each one explicit,
owned, and parseable:
Agentlocks is advisory : it coordinates agents that agree to use it. It does not stop an
editor, shell command, or Git operation that ignores the protocol, which is also why it needs
no daemon, no privileges, and no lock-holding background process.
Approach
Works well for
Where it falls short for shared worktrees
Manual notes in chat or issues
Informal coordination and intent
No lease, no owner check, no parseable status, easy to forget before staging
Shell scripts
Local conventions around one repo
Usually miss conflict semantics, stale sessions, JSON contracts, and Git-index locking
flock
Process-level critical sections on one machine
Not a repo resource registry; no resource inventory, owner metadata, install guidance, or agent docs
Git-native hooks ( pre-commit )
Commit-time policy checks
Too late to prevent overlapping edits; hooks do not coordinate git add across workers, and a single core.hooksPath collides with husky/lefthook
Hosted lock service
Cross-machine coordination
Requires a service, credentials, network access, and operational ownership
Agentlocks
Local agents in one repository worktree
Advisory only; participants opt in. Ships git verify plus an opt-out PreToolUse backstop (Claude Code + Codex) that runs it before a git commit tool-call; installs no git hook and never reconfigures your git
Quick Demo
This demo runs inside Codex or Claude Code. The active harness identity is detected
automatically, including Claude Code subagents when agentlocks init --harness claude-code has
installed the project hook.
HOST= " $( mktemp -d ) "
cd " $HOST "
git init -q
printf ' console.log("hello")\n ' > app.ts
printf ' # Demo host repo\n ' > AGENTS.md
printf ' {"scripts":{}}\n ' > package.json
# Expected to exit 1 when init drift is found; it does not write files.
agentlocks init --check --json || true
agentlocks init
agentlocks identify --json
file_lock= " $( agentlocks acquire ' app.ts ' --reason " edit app " --id-only ) "
agentlocks expand ' README.md ' --lock " $file_lock " --id-only
agentlocks refresh " $file_lock " --id-only
# git begin --id-only prints two lines: the lock id, then a fence token. Read both.
{ read git_lock ; read git_token ; } < <( agentlocks git begin --refresh-lock " $file_lock " --reason " commit demo " --id-only )
agentlocks status --json
agentlocks git end " $git_lock " --git-token " $git_token " --release-lock " $file_lock " --id-only
agentlocks prune --dry-run --json
agentlocks doctor --json
Expected shape:
identify shows a harness source such as CLAUDE_CODE_SESSION_ID, CODEX_THREAD_ID, or AGENTLOCKS_HARNESS_AGENT_ID.
status shows two locks while the file lock and @git/index lock are held.
git end prints the released git lock id and file lock id.
prune --dry-run reports pruned_count 0 in a fresh repo.
doctor reports ok true after init completes.
Host Setup
agentlocks init is idempotent and writes the host-repo support files. Run it once per repo;
add --harness claude-code (or --harness codex ) to also install the PreToolUse hooks.
agentlocks init --check --json || true # preview changes; exits 1 on drift, writes nothing
agentlocks init
# Claude Code: also install the .claude PreToolUse hooks.
agentlocks init --check --harness claude-code --json || true
agentlocks init --harness claude-code
init can create or update:
The commit-hook backstop is advisory: it surfaces staged-but-unlocked paths before a git commit
tool-call and never blocks the commit . It is installed by default; agentlocks init --no-commit-hook keeps the original id-injection-only Claude hook and skips the Codex hook.
The steps below assume the global install above and a supported agent harness. Codex and Claude
Code identity is automatic.
Initialize the host repo (writes the AGENTS.md instructions block).
agentlocks init --check --json || true
agentlocks init
# Claude Code: also install the .claude PreToolUse hooks.
agentlocks init --check --harness claude-code --json || true
agentlocks init --harness claude-code
Inspect the detected agent id.
Acquire the narrowest lock before editing.
agentlocks acquire ' src/index.ts ' ' tests/cli.test.ts ' --reason " change CLI dispatch " --id-only
Expand before touching another file.
agentlocks expand ' src/config.ts ' --lock < lock_id > --id-only
Refresh before long edit batches, after tests, and before staging.
agentlocks refresh < lock_id > --id-only
Coordinate the shared Git index. git begin --id-only prints two lines: the git lock id, then a
fence token that git end re-checks (it aborts with exit 3 if the lease was reclaimed mid-commit).
{ read git_lock ; read git_token ; } < <( agentlocks git begin --refresh-lock < lock_id > --reason " commit Agentlocks change " --id-only )
git add < locked_paths >
git commit
agentlocks git end " $git_lock " --git-token " $git_token " --release-lock < lock_id > --id-only
Or let agentlocks commit do all of it (lock, stage, commit, fence, release) in one command.
Before a raw git commit , sanity-check coverage (advisory; never blocks):
agentlocks git verify --json # lists any staged-but-unlocked paths
This is what the opt-out PreToolUse commit-hook backstop runs automatically. If you ever lose
your lock ids (e.g. after context compaction), recover with agentlocks status --mine and
agentlocks release --mine / agentlocks refresh --mine .
agentlocks capabilities --json is the source of truth for command metadata, positionals, JSON
shape refs, --id-only line contracts, flags, exit codes, default TTLs, agent identity detection,
and next commands.
Quote every resource so the shell passes it to Agentlocks literally. Exact paths and glob patterns
share the same positional list. The CLI treats those quoted arguments as raw resource specs; JSON
lock records use resources for the normalized {kind,value} entries stored in each lease.
agentlocks acquire ' a.ts ' ' b.ts ' ' src/**/*.ts ' --reason " edit files "
agentlocks acquire ' analyses/trending-baseline/** ' --reason " create Phase 0 artifacts "
Command
Purpose
Key flags
Output notes
acquire [resources...]
Acquire locks for quoted repo-relative paths or glob patterns
--reason , --ttl-ms , --reclaim , --agent-id , --json , --id-only , --verbose
--reason and at least one resource are required; glob-like resources are inferred from * , ? , or bracket classes
expand [resources...] --lock <id>
Add quoted repo-relative resour

[truncated]
