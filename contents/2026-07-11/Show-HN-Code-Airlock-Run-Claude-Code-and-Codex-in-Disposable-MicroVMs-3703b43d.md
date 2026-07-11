---
source: "https://github.com/Trivo25/code-airlock"
hn_url: "https://news.ycombinator.com/item?id=48870603"
title: "Show HN: Code Airlock: Run Claude Code and Codex in Disposable MicroVMs"
article_title: "GitHub - Trivo25/code-airlock: Run Claude Code (or another coding agent) unattended, inside a disposable microVM, with its work committed to git so you can review everything from your host. · GitHub"
author: "zkTrivo"
captured_at: "2026-07-11T10:55:04Z"
capture_tool: "hn-digest"
hn_id: 48870603
score: 2
comments: 0
posted_at: "2026-07-11T10:16:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Code Airlock: Run Claude Code and Codex in Disposable MicroVMs

- HN: [48870603](https://news.ycombinator.com/item?id=48870603)
- Source: [github.com](https://github.com/Trivo25/code-airlock)
- Score: 2
- Comments: 0
- Posted: 2026-07-11T10:16:50Z

## Translation

タイトル: HN の表示: Code Airlock: 使い捨て MicroVM でのクロード コードとコーデックスの実行
記事のタイトル: GitHub - Trivo25/code-airlock: クロード コード (または別のコーディング エージェント) を使い捨ての microVM 内で無人で実行し、その作業を git にコミットして、ホストからすべてを確認できるようにします。 · GitHub
説明: クロード コード (または別のコーディング エージェント) を使い捨ての microVM 内で無人で実行し、その作業を git にコミットして、ホストからすべてを確認できるようにします。 - Trivo25/コードエアロック

記事本文:
GitHub - Trivo25/code-airlock: クロード コード (または別のコーディング エージェント) を使い捨ての microVM 内で無人で実行し、その作業を git にコミットして、ホストからすべてを確認できるようにします。 · GitHub
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
トリボ25
/
コードエアロック
公共
通知

ns
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット .github/ workflows .github/ workflows Formula Formula docs docs .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md anime.gif anime.gif code-airlock code-airlock install.sh install.sh package.json package.json Sandbox.conf.example Sandbox.conf.example すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code、Codex、OpenCode、または別のコーディング エージェントを使い捨て microVM で実行し、通常の git コミットとしてその動作を確認します。
Code Airlock は、コーディング エージェントにホスト マシンへの直接アクセスを与えずに、より少ないプロンプトで作業できるようにしたい人向けの、Docker Sandbox の小さなラッパーです。
コーディング エージェントが最も役立つのは、コマンドごとに事前に確認することなく、パッケージのインストール、テストの実行、障害の検査、および反復処理ができる場合です。これはまさに、プロセス レベルの拒否ルールよりも強力な境界が必要な場合でもあります。
Code Airlock はワークフローをシンプルに保ちます。
コーディング エージェントは、依存関係のインストール、生成されたファイルの検査、テスト スイートの実行、ローカル サービスの開始、ログの読み取り、失敗時の反復など、タスクを完了するために必要なツールを使用できる場合に最も効果的です。すべてのシェル コマンドを過度に制限すると、エージェントがより遅いオートコンプリート ループに陥ることがよくあります。
Code Airlock は、主要な安全境界をエージェントの下に移動します。エージェントは、使い捨ての microVM 内でより自律的に動作できますが、ホスト リポジトリは読み取り専用のままで、結果として生じた変更はレビューのためにコミットとして返されます。
sbx を直接使用しないのはなぜですか?
あなたはできる。 Code Airlock は意図的に Docker Sandbox の薄いラッパーであり、Docker Sandbox の代替ではありません。

それ。
サンドボックスのライフサイクル、ポリシー、キット、および 1 回限りの実験を完全に制御したい場合は、sbx を直接使用します。共通のコーディング エージェント ループを既に接続したい場合は、Code Airlock を使用します。
リポジトリごとの安定したサンドボックスの名前付け
repo-local AGENTS.md スキャフォールディング
サンドボックス ブランチの fetch 、 diff 、 review 、および merge コマンド
Claude Code、Codex、OpenCode の文書化された資格情報とネットワーク ポリシー パス
重要なのは、安全なワークフローを退屈で反復可能なものにしつつ、その下で sbx を利用できるようにすることです。
エージェント ハーネスをより適切に構成しないのはなぜでしょうか?
エージェント ハーネス ルールは便利であり、Code Airlock はそれらと連動するように設計されています。 Codex の承認、Claude Code の権限、OpenCode 設定、 AGENTS.md 、およびカスタム プロンプトはすべて、エージェントが行うべき動作を決定するのに役立ちます。
これらのルールは依然としてエージェント プロセス内のポリシーです。使い捨ての VM、読み取り専用のホスト リポジトリ、エージェント コミット用の別個のクローン、変更が適用される前のホスト側のレビュー手順は提供されません。また、Claude Code、Codex、OpenCode、およびその他のエージェント用の共有ラッパーも提供しません。
Code Airlock は、ハーネス ルールを最初の層として扱い、Docker サンドボックスをその下の境界として扱います。エージェントは適切に動作するように構成できます。サンドボックスは、そうでない場合に起こることを制限します。
まず Docker サンドボックスをインストールします。 Linux の場合:
カール -fsSL https://get.docker.com | sudo REPO_ONLY=1 シュ
sudo apt-get install docker-sbx
sudo usermod -aG kvm $USER
newgrp kvm
sbxログイン
macOS および Windows では、Docker の公式インストール ガイドに従ってください。
npm を使用して Code Airlock をインストールします。
npm install -g code-airlock
または、Homebrew を使用します (タグ付きリリースが存在するまではヘッドのみの式):
醸造タップ trivo25/code-airlock https://github.com/Trivo25/code-airlock
brew install --HEAD code-airlock
CLI がインストールされると、あらゆる手順を実行できます。

前提条件がありません:
code-airlock セットアップ # ガイド付きインストール;各コマンドを表示し、実行前に質問します
code-airlock ドクター # チェックだけを行い、何も変更しません
または、スタンドアロンのシェル インストーラーを使用します。
カール -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh |しー
シェル インストーラーは、コマンドがまだ実行されていない場合、短いコードロック エイリアスも追加します。
コードロックアップ
コードロックのステータス
コードロックアタッチ
codelock はすでに npm パッケージ名であるため、npm パッケージは意図的に code-airlock のみをインストールします。 npm install 後に短いコマンドが必要な場合は、個人用シェル エイリアスを追加します。
エイリアス codelock=code-airlock
または、手動でクローンを作成します。
git クローン https://github.com/Trivo25/code-airlock.git
CDコード-エアロック
ln -s " $PWD /code-airlock " ~ /.local/bin/code-airlock
インストーラーはデフォルトで ~/.local/bin に書き込みます。次のコマンドを使用して別の場所を選択します。
カール -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh | PREFIX=/usr/local/bin sh
短いエイリアスをスキップするか名前を変更します。
カール -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh | CODE_AIRLOCK_INSTALL_ALIAS=0 シュ
カール -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh | CODE_AIRLOCK_ALIAS=カロック シュ
タスクを実行する
最短のループ、1 つのコマンドを入力し、通常の Git レビューを出力します。
code-airlock run " ユーザーエンドポイントにページネーションを追加し、テストを実行します "
コードエアロックの差分
コードとエアロックのマージ
run は、up と同様にサンドボックスを開始し、タスクをエージェントの非対話モードに渡します。クロード コードの場合は -p、Codex の場合は exec --dangerously-bypass-approvals-and-sandbox、OpenCode の場合は run です。他のエージェントはタスクを単純な引数として受け取ります。
すでにコーディング エージェントを使用していますか?これを送信してください:
この環境に Code Airlock をインストールし、初回実行チェックを実行します。
利用可能な場合は npm を使用します。
npm

インストール -g コード-エアロック
npm が使用できない場合は、次を使用します。
カール -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh |しー
インストール後、code-airlock が PATH 上にあることを確認して、次を実行します。
コードエアロックドクター
code-airlock --dry-run up
私が明示的に要求しない限り、実際のサンドボックスを開始したり、ネットワーク ポリシーを変更したりしないでください。
エージェントを動作させたいリポジトリから実行します。
cd ~ /code/my-project
code-airlock Doctor # sbx、仮想化、ログイン、git をチェックする
code-airlock --dry-run up # sbx コマンドを実行せずにプレビューします
code-airlock init # オプション: スターター AGENTS.md 命令を追加します
code-airlock up # サンドボックスでクロードコードを開始する
code-airlock --seed-config up # オプションでユーザーレベルのエージェントのカスタマイズをコピーします
code-airlock status # このリポジトリのサンドボックスと tmux セッションを参照
リモート サーバーでは、SSH 接続が閉じた後もエージェントを実行し続けます。
code-airlock --tmux up # Ctrl-b でデタッチしてから d
code-airlock アタッチ # 後で再アタッチ
別のエージェントを使用します。
code-airlock --agent codex アップ
code-airlock --agent opencode 実行「ビルドを修正」
AGENT 環境変数と Sandbox.conf 設定は引き続き機能します。両方が設定されている場合は、--agent フラグが優先されます。
エージェントが変更を加えたら、ホストから変更を確認します。
コードエアロックフェッチ
コードエアロックの差分
code-airlock レビュー # オプション: 視覚的な差分を開きます
コードとエアロックのマージ
ローカル マージよりもプル リクエストを優先します。
code-airlock pr # サンドボックス ブランチをプッシュして PR を開きます
code-airlock pr main feature/foo # 明示的なベースとブランチ名
プッシュは認証情報を使用してホストから行われるため、サンドボックスは GitHub にアクセスする必要がありません。 pr は、インストールされている場合は gh を使用し、それ以外の場合は GitHub 比較 URL を出力します。デフォルトのブランチ名は code-airlock/<sandbox-name> です。 2 番目の引数または PR_BRANCH でオーバーライドします。
コードエアロック

k は --clone を使用して Docker サンドボックスを開始します。
Docker はサンドボックス VM 内にプライベート クローンを作成します。
エージェントはそのクローン内で編集とコミットを行います。
エージェントの実行中、ホスト リポジトリは読み取り専用のままです。
サンドボックス ブランチをローカル リポジトリにフェッチし、差分を確認します。
満足した場合にのみマージします。
初期化されたサブモジュールが存在する場合、 fetch 、 diff 、 review 、および merge も、各サンドボックス サブモジュールの HEAD コミットを、一致するホスト サブモジュールにインポートしようとします。サブモジュールのコミットは以下にフェッチされます。
refs/remotes/code-airlock/<サンドボックス名>
これにより、ホストが実際にチェックアウトできるサブモジュールのコミットを親リポジトリのマージ ポイントにできるようになります。親のみのフェッチ動作の場合は、FETCH_SUBMODULES=0 でこれを無効にします。
これは、コミットされていないサンドボックスの編集が通常のフローでは戻されないことを意味します。必要に応じて、サンドボックスに入り、最初にコミットします。
コードエアロックシェル
git ステータス
git add -A
git commit -m "エージェントの動作を説明します"
出る
次に、ホストから取得してレビューします。
セットアップの失敗のほとんどは、Docker サンドボックスの欠落、KVM/仮想化サポートの欠落、または認証されていない sbx CLI によって発生します。起動する前に確認してください:
コードエアロックドクター
サンドボックスを作成せずに Code Airlock が実行する内容をプレビューします。
code-airlock --dry-run up
ドライラン モードでは、sbx および git コマンドを実行する代わりに出力します。これは、まだ KVM を実行できないマシンで、または VM を作成する前にワークフローを理解したい場合に役立ちます。
--tmux を使用して、ホスト側の tmux セッション内でサンドボックス エージェントを開始します。
コードエアロック --tmux アップ
これはすぐに接続されるので、コーディング エージェントと会話できるようになります。停止せずに切断するには、 Ctrl-b を押してから d を押します。これにより、端末が切断され、エージェントは tmux セッションで実行されたままになります。後で次のように再接続します。
コードエアロックアタッチ
無人サーバーを実行する場合は、att を使用せずに tmux セッションを開始します。

痛む：
code-airlock --tmux-detached up
デフォルトのセッション名は code-airlock-<sandbox-name> です。これを次のようにオーバーライドします。
TMUX_SESSION=エージェント作業コード-エアロック --tmux アップ
TMUX_SESSION=エージェント作業コードエアロック接続
視覚的に確認するには、次を使用します。
コードエアロックのレビュー
これにより、サンドボックス ブランチがフェッチされ、デフォルトで VS Code を使用して Git ディレクトリの差分が開きます。作業ツリーは review コマンドによって変更されません。
REVIEW_TOOL=opendiff コード-エアロックのレビュー
REVIEW_TOOL=vimdiff コード-エアロックのレビュー
エージェントの指示
リポジトリに簡潔なプロジェクト手順が含まれていると、エージェントの機能が向上します。 Code Airlock はスターター ファイルを作成できます。
コードエアロック初期化
これにより、AGENTS.md がまだ存在しない場合にのみ、ターゲット リポジトリに AGENTS.md が書き込まれます。 code-airlock up は、AGENTS.md が存在しない場合にコマンドを言及しますが、作業ツリーを黙って変更することはありません。
デフォルトでは、サンドボックスは新しいエージェントのホーム ディレクトリから開始します。 Git にコミットされたプロジェクト ファイルはクローン モードを通じて提供されますが、個人用のコマンド、スキル、サブエージェント、およびホストのホーム ディレクトリからのグローバル指示はクローン モードを通じて提供されません。
エージェントを開始する前に、厳選されたユーザーレベルの構成をコピーすることをオプトインします。
code-airlock --seed-config アップ
これにより、起動が 1 つの sbx run 呼び出しから次のように変更されます。
sbxクレア

[切り捨てられた]

## Original Extract

Run Claude Code (or another coding agent) unattended, inside a disposable microVM, with its work committed to git so you can review everything from your host. - Trivo25/code-airlock

GitHub - Trivo25/code-airlock: Run Claude Code (or another coding agent) unattended, inside a disposable microVM, with its work committed to git so you can review everything from your host. · GitHub
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
Trivo25
/
code-airlock
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits .github/ workflows .github/ workflows Formula Formula docs docs .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md animation.gif animation.gif code-airlock code-airlock install.sh install.sh package.json package.json sandbox.conf.example sandbox.conf.example View all files Repository files navigation
Run Claude Code, Codex, OpenCode, or another coding agent in a disposable microVM, then review its work as normal git commits.
Code Airlock is a small wrapper around Docker Sandboxes for people who want to let coding agents work with fewer prompts without giving them direct access to the host machine.
Coding agents are most useful when they can install packages, run tests, inspect failures, and iterate without asking before every command. That is also exactly when you want a stronger boundary than process-level deny rules.
Code Airlock keeps the workflow simple:
Coding agents are most effective when they can use the tools needed to finish the task: install dependencies, inspect generated files, run test suites, start local services, read logs, and iterate on failures. Over-constraining every shell command often turns the agent into a slower autocomplete loop.
Code Airlock moves the main safety boundary below the agent. The agent can work with more autonomy inside a disposable microVM, while the host repo stays read-only and the resulting changes still come back as commits for review.
Why Not Just Use sbx Directly?
You can. Code Airlock is intentionally a thin wrapper around Docker Sandboxes, not a replacement for it.
Use sbx directly when you want full control over sandbox lifecycle, policies, kits, and one-off experiments. Use Code Airlock when you want the common coding-agent loop already wired together:
stable sandbox naming per repo
repo-local AGENTS.md scaffolding
fetch , diff , review , and merge commands for the sandbox branch
a documented credential and network-policy path for Claude Code, Codex, and OpenCode
The point is to make the safe workflow boring and repeatable while still leaving sbx available underneath.
Why Not Just Configure the Agent Harness Better?
Agent harness rules are useful, and Code Airlock is meant to work with them. Codex approvals, Claude Code permissions, OpenCode settings, AGENTS.md , and custom prompts all help shape what the agent should do.
Those rules are still policy inside the agent process. They do not give you a disposable VM, a read-only host repo, a separate clone for agent commits, or a host-side review step before changes land. They also do not provide a shared wrapper for Claude Code, Codex, OpenCode, and other agents.
Code Airlock treats harness rules as the first layer and Docker Sandboxes as the boundary underneath it. The agent can be configured to behave well; the sandbox limits what happens when it does not.
Install Docker Sandboxes first. On Linux:
curl -fsSL https://get.docker.com | sudo REPO_ONLY=1 sh
sudo apt-get install docker-sbx
sudo usermod -aG kvm $USER
newgrp kvm
sbx login
On macOS and Windows, follow Docker's official install guide .
Install Code Airlock with npm:
npm install -g code-airlock
Or with Homebrew (head-only formula until a tagged release exists):
brew tap trivo25/code-airlock https://github.com/Trivo25/code-airlock
brew install --HEAD code-airlock
Once the CLI is installed, it can walk you through any missing prerequisites:
code-airlock setup # guided install; shows each command and asks before running it
code-airlock doctor # checks only, changes nothing
Or use the standalone shell installer:
curl -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh | sh
The shell installer also adds a short codelock alias when that command is not already taken:
codelock up
codelock status
codelock attach
The npm package intentionally installs only code-airlock , because codelock is already an npm package name. Add a personal shell alias if you want the shorter command after npm install:
alias codelock=code-airlock
Or clone it manually:
git clone https://github.com/Trivo25/code-airlock.git
cd code-airlock
ln -s " $PWD /code-airlock " ~ /.local/bin/code-airlock
The installer writes to ~/.local/bin by default. Choose another location with:
curl -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh | PREFIX=/usr/local/bin sh
Skip or rename the short alias:
curl -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh | CODE_AIRLOCK_INSTALL_ALIAS=0 sh
curl -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh | CODE_AIRLOCK_ALIAS=calock sh
Run a Task
The shortest loop, one command in and ordinary Git review out:
code-airlock run " add pagination to the users endpoint and run the tests "
code-airlock diff
code-airlock merge
run starts the sandbox like up and hands the task to the agent's non-interactive mode: -p for Claude Code, exec --dangerously-bypass-approvals-and-sandbox for Codex, and run for OpenCode. Other agents receive the task as a plain argument.
Already using a coding agent? Send it this:
Install Code Airlock in this environment, then run its first-run checks.
Use npm when available:
npm install -g code-airlock
If npm is unavailable, use:
curl -fsSL https://raw.githubusercontent.com/Trivo25/code-airlock/main/install.sh | sh
After install, make sure code-airlock is on PATH, then run:
code-airlock doctor
code-airlock --dry-run up
Do not start a real sandbox or change network policy unless I explicitly ask.
Run it from the repo you want the agent to work on:
cd ~ /code/my-project
code-airlock doctor # check sbx, virtualization, login, and git
code-airlock --dry-run up # preview the sbx commands without running them
code-airlock init # optional: add starter AGENTS.md instructions
code-airlock up # start Claude Code in a sandbox
code-airlock --seed-config up # optionally copy user-level agent customizations
code-airlock status # see this repo's sandbox and tmux session
On a remote server, keep the agent running after your SSH connection closes:
code-airlock --tmux up # detach with Ctrl-b, then d
code-airlock attach # reattach later
Use another agent:
code-airlock --agent codex up
code-airlock --agent opencode run " fix the build "
The AGENT environment variable and the sandbox.conf setting still work; the --agent flag wins when both are set.
When the agent has made changes, review them from your host:
code-airlock fetch
code-airlock diff
code-airlock review # optional: open a visual diff
code-airlock merge
Prefer a pull request over a local merge:
code-airlock pr # push the sandbox branch and open a PR
code-airlock pr main feature/foo # explicit base and branch name
The push happens from the host with your credentials, so the sandbox still needs no GitHub access. pr uses gh when it is installed and prints the GitHub compare URL otherwise. The default branch name is code-airlock/<sandbox-name> ; override it with the second argument or PR_BRANCH .
Code Airlock starts Docker Sandboxes with --clone :
Docker creates a private clone inside the sandbox VM.
The agent edits and commits inside that clone.
Your host repo stays read-only while the agent runs.
You fetch the sandbox branch into your local repo and review the diff.
You merge only when you are satisfied.
If initialized submodules exist, fetch , diff , review , and merge also try to import each sandbox submodule's HEAD commit into the matching host submodule. Submodule commits are fetched into:
refs/remotes/code-airlock/<sandbox-name>
This lets the parent repo merge point at submodule commits that the host can actually check out. Disable this with FETCH_SUBMODULES=0 for parent-only fetch behavior.
This means uncommitted sandbox edits do not come back through the normal flow. If needed, enter the sandbox and commit first:
code-airlock shell
git status
git add -A
git commit -m " describe the agent work "
exit
Then fetch and review from the host.
Most setup failures come from missing Docker Sandboxes, missing KVM/virtualization support, or an unauthenticated sbx CLI. Check before launching:
code-airlock doctor
Preview what Code Airlock would run without creating a sandbox:
code-airlock --dry-run up
Dry-run mode prints the sbx and git commands instead of running them. It is useful on machines that cannot run KVM yet, or when you want to understand the workflow before creating a VM.
Use --tmux to start the sandboxed agent inside a host-side tmux session:
code-airlock --tmux up
This attaches immediately so you can talk to the coding agent. To disconnect without stopping it, press Ctrl-b , then d . This detaches your terminal and leaves the agent running in the tmux session. Reattach later with:
code-airlock attach
For an unattended server run, start the tmux session without attaching:
code-airlock --tmux-detached up
The default session name is code-airlock-<sandbox-name> . Override it with:
TMUX_SESSION=agent-work code-airlock --tmux up
TMUX_SESSION=agent-work code-airlock attach
For a visual review, use:
code-airlock review
This fetches the sandbox branch and opens a Git directory diff using VS Code by default. Your working tree is not modified by the review command.
REVIEW_TOOL=opendiff code-airlock review
REVIEW_TOOL=vimdiff code-airlock review
Agent Instructions
Agents work better when the repo contains concise project instructions. Code Airlock can create a starter file:
code-airlock init
This writes AGENTS.md in the target repo only if one does not already exist. code-airlock up will mention the command when no AGENTS.md is present, but it will not silently modify your working tree.
By default, a sandbox starts with a fresh agent home directory. Project files that are committed to Git come through clone mode, but personal commands, skills, subagents, and global instructions from your host home directory do not.
Opt in to copying curated user-level config before the agent starts:
code-airlock --seed-config up
This changes startup from one sbx run call to:
sbx crea

[truncated]
