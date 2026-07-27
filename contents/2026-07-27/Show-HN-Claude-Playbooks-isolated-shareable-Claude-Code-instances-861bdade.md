---
source: "https://github.com/ramazanpolat/claude-playbooks"
hn_url: "https://news.ycombinator.com/item?id=49068803"
title: "Show HN: Claude Playbooks – isolated, shareable Claude Code instances"
article_title: "GitHub - ramazanpolat/claude-playbooks: Isolated Claude Code instances to keep things separated and clean. · GitHub"
author: "ramazanpolat"
captured_at: "2026-07-27T12:50:59Z"
capture_tool: "hn-digest"
hn_id: 49068803
score: 1
comments: 0
posted_at: "2026-07-27T12:37:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude Playbooks – isolated, shareable Claude Code instances

- HN: [49068803](https://news.ycombinator.com/item?id=49068803)
- Source: [github.com](https://github.com/ramazanpolat/claude-playbooks)
- Score: 1
- Comments: 0
- Posted: 2026-07-27T12:37:54Z

## Translation

タイトル: HN を表示: クロード プレイブック – 分離された共有可能なクロード コード インスタンス
記事のタイトル: GitHub - ramazanpolat/claude-playbooks: 物事を分離してクリーンに保つための分離されたクロード コード インスタンス。 · GitHub
説明: クロード コードのインスタンスを分離して、物事を分離してクリーンに保ちます。 - ラマザンポラット/クロード・プレイブック
HN テキスト: スキルやプラグインをインストールしたり、MCP 接続を確立したりして、CC 構成を変更したい場合があります。しかし、CC をカスタマイズすればするほど、きれいに保つのが難しくなり、複雑になっていきます。実験的なフックや新しい CLAUDE.md の動作を試してみませんか?現在の設定を変更する必要があり、気に入らない場合は元に戻す必要があります。時間が経つにつれて、すべての変更に対処するのが面倒になり、2 つの異なるアカウントを使用したいと思いました。そこで、実際に分離された CC インスタンスを使用できることを発見しました。これは、複数の CC セットアップを行う素晴らしい方法です。デフォルトでは、CC は ~/.claude を指す CLAUDE_CONFIG_DIR を利用していました。新しい CC が必要な場合は、これを実行してください: mkdir ~/.claude-personal
CLAUDE_CONFIG_DIR=~/.claude-personal claude と別個の CC インスタンスがあり、必要に応じて構成でき、これにエイリアスを付けることもできます。
alias claude-personal='CLAUDE_CONFIG_DIR=~/.claude-personal claude' (これを ~/.bashrc または ~/.zshrc に入れます) で準備完了です。いくつかの CLAUDE_CONFIG_DIR を作成しましたが、それらを並行して使用できるのは素晴らしいことです。たとえば、コミュニティ プラグインとスキル (caveman、i-have-adhd、context-mode など) が読み込まれた実験的な CC セットアップがあり、「claude-experimental」と入力するだけでこのセットアップにアクセスできます。そして、別のアカウントを使用する別のセットアップもあります。彼らは両方とも孤立して働いています。作業を簡単にするために、管理と *sha を行うための小さなツールを作成しました。

* 分離されたクロード コードのセットアップ: https://github.com/ramazanpolat/claude-playbooks/ 単一の静的バイナリ、依存関係なし、MIT ライセンス。フィードバックは大歓迎です!

記事本文:
GitHub - ramazanpolat/claude-playbooks: 物事を分離してクリーンに保つための分離されたクロード コード インスタンス。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ラマザンポラット
/
クロード・プレイブック
公共
通知
通知を変更するにはサインインする必要があります

アイケーション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
64 コミット 64 コミット .github/ workflows .github/ workflows cmd cmd 内部 内部テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md SPEC-v4.md SPEC-v4.md build.sh build.sh go.mod go.mod go.sum go.sum install.sh install.sh main.go main.go uninstall.sh uninstall.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード プレイブックは、クロード コードの独立したインスタンスです。
それでおしまい。各 Playbook には独自の構成、設定、フック、メモリ、タスク履歴があり、デフォルトの ~/.claude インストールや他のすべての Playbook とは完全に分離されています。
Claude Code は、設定、会話履歴、権限、フック、MCP サーバーなど、すべてを ~/.claude/ に保存します。何か (別のモデル、カスタム フック、新しい CLAUDE.md の動作) を試したい場合は、メインのセットアップを変更する必要があります。変更を 1 つ間違えると、日常のワークフローが中断されます。
プレイブックは、各実験 (またはワークフロー) に独自の分離されたディレクトリを与えることで、この問題を解決します。
メインの ~/.claude を危険にさらさずに、新しいフックまたは設定をテストします。
仕事と個人の設定を分けて干渉しないようにする
異なる性格を持つ異なるタスクで 2 つのクロード コード インスタンスを同時に実行します。
異なるアカウントで同時に認証します (例: 1 つのプレイブックを企業アカウントで認証し、別のプレイブックを個人アカウントで認証したままにします)
プレイブックを Git リポジトリに配置して、チームと構成を共有します
1 つ以上の Playbook 設定を含むリポジトリを使用します (例: サブディレクトリ経由)。
Claude Code は、 CLAUDE_CONFIG_DIR に設定されたディレクトリ (デフォルトは ~/.claude ) から設定を読み取ります。その変数を変更すると、完全に新しい、独立した変数が得られます。

ent インスタンス:
# 通常のクロード コード (~/.claude を使用)
クロード
# 分離された Playbook (~/.claude-playbooks/experiment を使用)
CLAUDE_CONFIG_DIR= ~ /.claude-playbooks/実験クロード
プレイブックの内部にあるのはこれだけです。 claude-Playbook を使用すると、それらの作成、共有、管理が簡単になります。
~/.claude-playbooks/ シェル エイリアス:
§── Experiment/ ◄── alias Experiment='CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/experiment" cpb run Experiment'
│ §── CLAUDE.md
│ └── settings.json
│
└── 素晴らしい/ ◄── エイリアス ap='CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/awesome" cpb run 素晴らしい'
§── .playbook (マーカー + メタデータ)
━── クロード.md
各 Playbook ディレクトリは、完全に分離された Claude Code インスタンスです。
ディレクトリは、Playbook ルートの下に存在する場合、Playbook です。 .playbook マニフェスト ファイルはオプションであり、メタデータ (バージョン、作成者、説明など) を保存するために使用されます。
カール -fsSL https://raw.githubusercontent.com/ramazanpolat/claude-playbooks/main/install.sh |しー
スクリプトは OS とアーキテクチャを検出し、最新の GitHub リリースから適切なバイナリをダウンロードし、/usr/local/bin (書き込み可能でない場合は ~/.local/bin) にインストールします。
クロード-プレイブック --バージョン
短いコマンド名を使用してインストールします。
カール -fsSL https://raw.githubusercontent.com/ramazanpolat/claude-playbooks/main/install.sh | INSTALL_NAME=cpb sh
cpb --バージョン
リポジトリのクローンを作成して、インストーラーをローカルで実行することもできます。
git clone https://github.com/ramazanpolat/claude-playbooks.git
CD クロード プレイブック
./install.sh
ローカル インストールでは、同じ短いコマンド名がサポートされます。
INSTALL_NAME=cpb ./install.sh
cpb --バージョン
バイナリのみをアンインストールします。
カール -fsSL https://raw.githubusercontent.com/ramazanpolat/claude-playbooks/main/uninstall.sh |しー
カスタム コマンド名をアンインストールするには:

カール -fsSL https://raw.githubusercontent.com/ramazanpolat/claude-playbooks/main/uninstall.sh | INSTALL_NAME=cpb sh
または、クローンからローカル アンインストーラーを実行します。カスタム コマンド名をインストールした場合は、同じ INSTALL_NAME を使用します。
./アンインストール.sh
INSTALL_NAME=cpb ./uninstall.sh
アンインストールしても ~/.claude-playbooks は削除されません。
クロードプレイブック自体をアンインストールする
ツール、インストールされているすべての Playbook、シェル エイリアス、および
ワンステップでバイナリ:
claude-playbook self-uninstall # 確認のプロンプト
クロード-プレイブックの自己アンインストール -y # プロンプトをスキップ
クロードプレイブック自己アンインストール -y --keep-data # keep ~/.claude-playbooks
claude-playbook self-uninstall -y --keep-binary # バイナリを保持します
claude-playbook self-uninstall --dry-run # 削除せずにプレビュー
バイナリを削除できない場合 (例: /usr/local/bin にインストールされている場合)
root ではない)、コマンドは sudo rm <path> ヒントを出力し、クリーンアップを続行します。
他のすべて。
手動フォールバック (バイナリを実行できない場合):
# 1. シェル構成 (~/.zshrc または ~/.bashrc) からエイリアスを削除します。
# 一致する行を削除します: alias ...='CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/...
# 2. rm -rf ~/.claude-playbooks
# 3. sudo rm /usr/local/bin/claude-playbook # またはバイナリが存在する場所
ソースからビルドする (Go 1.21 以降が必要):
git clone https://github.com/ramazanpolat/claude-playbooks.git
CD クロード プレイブック
./build.sh
mv クロード-プレイブック /usr/local/bin/
使用法
ほとんどのワークフローは、 create 、 install 、または link のいずれかで始まります。
claude-playbook は、新しく作成、インストール、リンクされた Playbook に対して既存の Claude Code 認証を再利用しようとします。つまり、新しいプレイブックでは通常、再ログインを要求するのではなく、Claude Code を直接開く必要があります。
異なるアカウントを同時に使用したい場合は、認証を有効にできます。

どのようなプレイブックでも大歓迎です。プレイブックの .playbook マニフェスト ファイルで isolate_auth = true を設定するか、環境変数 CLAUDE_PLAYBOOKS_ISOLATE_AUTH=true を使用してプレイブックを実行するだけです。これにより、Playbook のログイン セッションが分離され、他の Playbook やグローバル設定と資格情報を共有または自動同期できなくなります。
独自のプレイブックを作成して実行する
新しく分離されたクロード コードのセットアップが必要な場合は、create を使用します。
クロード プレイブックの作成実験
ソース ~ /.zshrc
実験
これにより、 ~/.claude-playbooks/experiment が作成され、その中で開かれた Claude Code セッションにプレイブックの概念を導入するスターター CLAUDE.md がドロップされ、Claude 認証メタデータが同期され、 Experiment という名前のシェル エイリアスが追加されます。 .playbook マニフェストはオプションであり、このコマンドでは作成されません。
エイリアスを使用せずに実行することもできます。
クロード・プレイブックの実験を実行する
Playbook 名の後に Claude Code フラグを渡します。
クロード プレイブック 実験の実行 --model クロード-opus-4-6 --permission-mode auto
カスタム エイリアスを使用するか、エイリアスの作成をスキップします。
claude-playbook create backend --alias be
クロード-プレイブック スクラッチを作成 --no-alias
何がインストールされているかを見る
クロード・プレイブックのリスト
最後に使用された名前パス エイリアス
実験 ~/.claude-playbooks/experiment exp 2 日前
素晴らしい ~/.claude-playbooks/awesome ap 2 時間前
共有 Playbook リポジトリをインストールする
Playbook が Git リポジトリまたはローカル ディレクトリにあり、コピーされたインストールを ~/.claude-playbooks の下に置きたい場合は、 install を使用します。
クロード-プレイブックのインストール https://github.com/ramazanpolat/awesome-playbooks
インストール名またはエイリアスをオーバーライドします。
クロード-プレイブックのインストール https://github.com/user/awesome --name チームツール --alias tt
ローカル ディレクトリをコピーしてインストールします。
クロード-プレイブックのインストール ~ /dev/my-playbook
より大きなリポジトリから 1 つの Playbook をインストールする
必要な場合は GitHub ツリー URL を使用してください

サブディレクトリは 1 つだけです。
クロードプレイブックのインストール https://github.com/user/awesome/tree/main/playbooks/dba
または、サブディレクトリを明示的に渡します。
クロード-プレイブックのインストール https://github.com/user/awesome --subdir playbooks/dba
厳選されたインストールは、フラットなトップレベルの Playbook です。
/ を含むブランチ名は、リポジトリのリモート参照に対して解決されます。 --branch feature/name を使用して境界を明示的にすることもできます。
クロード-プレイブックのインストール https://github.com/user/awesome --subdir playbooks/dba --name dba --alias ap-dba
プレイブックを適切に開発する
~/.claude-playbooks の外部で Playbook を積極的に編集していて、ライブ変更が必要な場合は、リンクを使用します。
クロード-プレイブックのリンク ~ /dev/my-playbook
link は、Playbook ルートの下にシンボリックリンクを作成します。
クロード-プレイブック リンク ~ /dev/my-playbook --nameScratch --alias sc
クロード-プレイブック リンク ~ /dev/my-playbook --no-alias
リンクされたプレイブックを削除すると、シンボリックリンクのみが削除されます。ソースディレクトリは保存されます。
生成されたエイリアスは、実行を claude-playbook run (または cpb run ) 経由で戻すシェル エイリアスです。
# クロード-プレイブック: 実験
エイリアス実験= ' CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/experiment" cpb 実験の実行 '
エイリアスを表示、設定、または削除します。
クロード・プレイブックのエイリアス
クロード・プレイブックのエイリアス実験
クロード プレイブック エイリアス実験 exp
クロード・プレイブックのエイリアス実験 --remove
クロード・プレイブックのディエイリアスの実験
エイリアスは通常のシェル行であるため、エイリアスを編集してクロード コード フラグを追加できます。
# クロード-プレイブック: 仕事
alias work= ' CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/work" cpb run work --model claude-opus-4-6 --permission-mode auto '
一時的なセッション
Playbook を登録せずに、1 回限りの Claude Code 構成ディレクトリに start を使用します。
クロード-プレイブックの開始 /tmp/scratch
クロード-プレイブックの開始 /tmp/scratch --model

クロード作品-4-6
クロード-プレイブックの開始 /tmp/scratch --delete
--delete は、セッションの終了時にディレクトリを削除します。これは、使い捨ての実験に役立ちます。
クロード・プレイブックの名前を実験室に変更
クロード・プレイブックの名前を変更するラボ実験 --alias exp
プレイブックを削除します。
クロード・プレイブックの実験の削除 # 確認のプロンプト
クロード-プレイブック削除素晴らしい -y # 確認をスキップ
uninstall と unlink は、 delete のコマンド エイリアスです。
クロード・プレイブックのアンインストールはすばらしい
クロード-プレイブック私のリンクされたプレイブックのリンクを解除します
最初のデリゲートをプレイブックが提供するスクリプトに更新します。
クロード・プレイブックのアップデートは素晴らしい
~/.claude-playbooks/awesome/bin/update-playbook.sh が存在する場合、Playbook ディレクトリ内から実行されます。 Git ベースの Playbook は次のものを出荷する可能性があります。
#! /bin/sh
セット -e
cd " $( dirname " $0 " ) /.. "
git pull --ff-only
Git インストールでは、リポジトリ、ブランチ、選択したサブディレクトリも .playbook に記録されます。更新スクリプトが存在しない場合は、フラットな非リンク インストールでそのソースからネイティブに更新できます。ネイティブ更新は完全バックアップをステージングし、ローカルのクロード状態と資格情報を保持しながら新しいソース ファイルをオーバーレイし、結果をアトミックにアクティブ化します。リンクされた Playbook とサブディレクトリを持つ従来のマニフェストには、委任された更新スクリプトが必要です。
名前がない場合、更新は cla を自己更新します

[切り捨てられた]

## Original Extract

Isolated Claude Code instances to keep things separated and clean. - ramazanpolat/claude-playbooks

Sometimes I want to change my CC configuration by installing some skills, plugins or having some MCP connections. However the more I customize my CC the more is gets convoluted because it is hard to keep it clean. Wanna try an experimental hook or new CLAUDE.md behavior? You have to change your current setup, then if you didn't like, you have to revert it back. Over time, it became a nuisance to deal with all of the changes, then I wanted to use 2 different accounts, that's where I have discovered that we can actually have isolated CC instances, which is a spectacular way to have multiple CC setups. By default, CC utilized CLAUDE_CONFIG_DIR which points to ~/.claude, if you want to have a fresh CC, just run this: mkdir ~/.claude-personal
CLAUDE_CONFIG_DIR=~/.claude-personal claude an you'll have separate CC instance, you may configure however you want and even have an alias for this:
alias claude-personal='CLAUDE_CONFIG_DIR=~/.claude-personal claude' (put this into your ~/.bashrc or ~/.zshrc) and you are good to go! I have created several CLAUDE_CONFIG_DIR and being able to use them in parallel feels great. For example, I have experimental CC setup that is loaded with community plugins and skills (caveman, i-have-adhd, context-mode, etc) and I can access this setup just by typing "claude-experimental". And another setup where I even use different account! They both work isolated. To make things easy, I've created a small tool to manage and *share* isolated claude-code setups: https://github.com/ramazanpolat/claude-playbooks/ Single static binary, no dependencies, MIT licensed. Feedbacks are welcome!

GitHub - ramazanpolat/claude-playbooks: Isolated Claude Code instances to keep things separated and clean. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
ramazanpolat
/
claude-playbooks
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
64 Commits 64 Commits .github/ workflows .github/ workflows cmd cmd internal internal tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md SPEC-v4.md SPEC-v4.md build.sh build.sh go.mod go.mod go.sum go.sum install.sh install.sh main.go main.go uninstall.sh uninstall.sh View all files Repository files navigation
A Claude Playbook is an isolated instance of Claude Code.
That's it. Each playbook has its own configuration, settings, hooks, memory, and task history -- completely separate from your default ~/.claude installation and from every other playbook.
Claude Code stores everything in ~/.claude/ : your settings, conversation history, permissions, hooks, MCP servers. If you want to try something -- a different model, a custom hook, a new CLAUDE.md behavior -- you have to touch your main setup. One wrong change and your daily workflow breaks.
Playbooks solve this by giving each experiment (or workflow) its own isolated directory.
Test a new hook or setting without risking your main ~/.claude
Separate work and personal configurations that don't interfere
Run two Claude Code instances concurrently on different tasks with different personalities
Authenticate with different accounts concurrently (e.g., keep one playbook authenticated with your corporate account and another with your personal account)
Share a configuration with your team by putting the playbook in a Git repo
Consume a repository containing one or more playbook configurations (e.g. via subdirectories)
Claude Code reads its configuration from the directory set in CLAUDE_CONFIG_DIR (defaults to ~/.claude ). Change that variable, and you get a completely fresh, independent instance:
# Your normal Claude Code (uses ~/.claude)
claude
# An isolated playbook (uses ~/.claude-playbooks/experiment)
CLAUDE_CONFIG_DIR= ~ /.claude-playbooks/experiment claude
That's all a playbook is under the hood. claude-playbook just makes creating, sharing, and managing them easy.
~/.claude-playbooks/ Shell aliases:
├── experiment/ ◄── alias experiment='CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/experiment" cpb run experiment'
│ ├── CLAUDE.md
│ └── settings.json
│
└── awesome/ ◄── alias ap='CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/awesome" cpb run awesome'
├── .playbook (marker + metadata)
└── CLAUDE.md
Each playbook directory is a completely isolated Claude Code instance.
A directory is a playbook if it exists under the playbooks root. A .playbook manifest file is optional and used for storing metadata (like version, author, description).
curl -fsSL https://raw.githubusercontent.com/ramazanpolat/claude-playbooks/main/install.sh | sh
The script detects your OS and architecture, downloads the right binary from the latest GitHub Release, and installs it to /usr/local/bin (or ~/.local/bin if that's not writable).
claude-playbook --version
Install with a shorter command name:
curl -fsSL https://raw.githubusercontent.com/ramazanpolat/claude-playbooks/main/install.sh | INSTALL_NAME=cpb sh
cpb --version
You can also clone the repo and run the installer locally:
git clone https://github.com/ramazanpolat/claude-playbooks.git
cd claude-playbooks
./install.sh
Local installs support the same shorter command name:
INSTALL_NAME=cpb ./install.sh
cpb --version
Uninstall only the binary:
curl -fsSL https://raw.githubusercontent.com/ramazanpolat/claude-playbooks/main/uninstall.sh | sh
To uninstall a custom command name:
curl -fsSL https://raw.githubusercontent.com/ramazanpolat/claude-playbooks/main/uninstall.sh | INSTALL_NAME=cpb sh
Or run the local uninstaller from a clone. Use the same INSTALL_NAME if you installed a custom command name:
./uninstall.sh
INSTALL_NAME=cpb ./uninstall.sh
Uninstalling does not delete ~/.claude-playbooks .
Uninstalling claude-playbook itself
To remove the tool, all its installed playbooks, their shell aliases, and the
binary in one step:
claude-playbook self-uninstall # prompts for confirmation
claude-playbook self-uninstall -y # skip prompt
claude-playbook self-uninstall -y --keep-data # keep ~/.claude-playbooks
claude-playbook self-uninstall -y --keep-binary # keep the binary
claude-playbook self-uninstall --dry-run # preview without removing
If the binary can't be removed (e.g. installed to /usr/local/bin and you're
not root), the command prints a sudo rm <path> hint and continues cleaning up
everything else.
Manual fallback (if you can't run the binary):
# 1. Remove aliases from your shell config (~/.zshrc or ~/.bashrc)
# Delete any lines matching: alias ...='CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/...
# 2. rm -rf ~/.claude-playbooks
# 3. sudo rm /usr/local/bin/claude-playbook # or wherever the binary lives
Build from source (requires Go 1.21+):
git clone https://github.com/ramazanpolat/claude-playbooks.git
cd claude-playbooks
./build.sh
mv claude-playbook /usr/local/bin/
Usage
Most workflows start with either create , install , or link .
claude-playbook tries to reuse your existing Claude Code authentication for newly created, installed, and linked playbooks. That means a new playbook should normally open Claude Code directly instead of asking you to log in again.
If you want to use different accounts concurrently , you can enable authentication isolation for any playbook. Simply set isolate_auth = true in the playbook's .playbook manifest file, or run the playbook with the environment variable CLAUDE_PLAYBOOKS_ISOLATE_AUTH=true . This isolates that playbook's login session and prevents it from sharing or auto-syncing credentials with your other playbooks or global settings.
Create and run your own playbook
Use create when you want a fresh isolated Claude Code setup.
claude-playbook create experiment
source ~ /.zshrc
experiment
This creates ~/.claude-playbooks/experiment , drops in a starter CLAUDE.md that introduces the playbook concept to the Claude Code session opened inside it, syncs Claude auth metadata, and adds a shell alias named experiment . A .playbook manifest is optional and is not created by this command.
You can also run it without using the alias:
claude-playbook run experiment
Pass Claude Code flags after the playbook name:
claude-playbook run experiment --model claude-opus-4-6 --permission-mode auto
Use a custom alias or skip alias creation:
claude-playbook create backend --alias be
claude-playbook create scratch --no-alias
See what is installed
claude-playbook list
NAME PATH ALIAS LAST USED
experiment ~/.claude-playbooks/experiment exp 2 days ago
awesome ~/.claude-playbooks/awesome ap 2 hours ago
Install a shared playbook repo
Use install when the playbook is in a Git repo or local directory and you want a copied install under ~/.claude-playbooks .
claude-playbook install https://github.com/ramazanpolat/awesome-playbooks
Override the install name or alias:
claude-playbook install https://github.com/user/awesome --name team-tools --alias tt
Install a local directory by copying it:
claude-playbook install ~ /dev/my-playbook
Install one playbook from a larger repo
Use a GitHub tree URL when you want only one subdirectory:
claude-playbook install https://github.com/user/awesome/tree/main/playbooks/dba
Or pass the subdirectory explicitly:
claude-playbook install https://github.com/user/awesome --subdir playbooks/dba
Cherry-picked installs are flat top-level playbooks.
Branch names containing / are resolved against the repository's remote refs. You can also make the boundary explicit with --branch feature/name .
claude-playbook install https://github.com/user/awesome --subdir playbooks/dba --name dba --alias ap-dba
Develop a playbook in place
Use link when you are actively editing a playbook outside ~/.claude-playbooks and want live changes.
claude-playbook link ~ /dev/my-playbook
link creates a symlink under the playbooks root.
claude-playbook link ~ /dev/my-playbook --name scratch --alias sc
claude-playbook link ~ /dev/my-playbook --no-alias
Deleting a linked playbook removes only the symlink. The source directory is preserved.
Generated aliases are shell aliases that route their execution back through claude-playbook run (or cpb run ):
# claude-playbook: experiment
alias experiment= ' CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/experiment" cpb run experiment '
Show, set, or remove aliases:
claude-playbook alias
claude-playbook alias experiment
claude-playbook alias experiment exp
claude-playbook alias experiment --remove
claude-playbook dealias experiment
Because aliases are ordinary shell lines, you can edit them to add Claude Code flags:
# claude-playbook: work
alias work= ' CLAUDE_CONFIG_DIR="$HOME/.claude-playbooks/work" cpb run work --model claude-opus-4-6 --permission-mode auto '
Temporary sessions
Use start for a one-off Claude Code config directory without registering a playbook:
claude-playbook start /tmp/scratch
claude-playbook start /tmp/scratch --model claude-opus-4-6
claude-playbook start /tmp/scratch --delete
--delete removes the directory when the session ends, which is useful for disposable experiments.
claude-playbook rename experiment lab
claude-playbook rename lab experiment --alias exp
Delete a playbook:
claude-playbook delete experiment # prompts for confirmation
claude-playbook delete awesome -y # skip confirmation
uninstall and unlink are command aliases for delete :
claude-playbook uninstall awesome
claude-playbook unlink my-linked-playbook
Update first delegates to a playbook-provided script:
claude-playbook update awesome
If ~/.claude-playbooks/awesome/bin/update-playbook.sh exists, it is run from inside the playbook directory. A Git-backed playbook might ship:
#! /bin/sh
set -e
cd " $( dirname " $0 " ) /.. "
git pull --ff-only
Git installs also record their repository, branch, and selected subdirectory in .playbook . If no update script exists, a flat, non-linked install can update natively from that source. Native update stages a full backup, overlays new source files while preserving local Claude state and credentials, and atomically activates the result. Linked playbooks and legacy manifests with subdir require a delegated update script.
With no name, update self-updates the cla

[truncated]
