---
source: "https://github.com/juliensimon/canopy"
hn_url: "https://news.ycombinator.com/item?id=48538315"
title: "Show HN: Canopy – parallel, sandboxed Claude Code sessions on macOS"
article_title: "GitHub - juliensimon/canopy: Ship faster with parallel Claude Code sessions — one native macOS window, git worktrees, sandboxes, auto-resume, merge & finish, token dashboard. · GitHub"
author: "julsimon"
captured_at: "2026-06-15T11:12:28Z"
capture_tool: "hn-digest"
hn_id: 48538315
score: 1
comments: 0
posted_at: "2026-06-15T08:33:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Canopy – parallel, sandboxed Claude Code sessions on macOS

- HN: [48538315](https://news.ycombinator.com/item?id=48538315)
- Source: [github.com](https://github.com/juliensimon/canopy)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T08:33:21Z

## Translation

タイトル: Show HN: Canopy – macOS 上の並列サンドボックス クロード コード セッション
記事のタイトル: GitHub - juliensimon/canopy: クロード コードの並列セッションでより高速に出荷 — 1 つのネイティブ macOS ウィンドウ、git ワークツリー、サンドボックス、自動再開、マージと終了、トークン ダッシュボード。 · GitHub
説明: 並列クロード コード セッション - 1 つのネイティブ macOS ウィンドウ、git ワークツリー、サンドボックス、自動再開、マージと終了、トークン ダッシュボードにより、より高速に出荷されます。 - ジュリエンシモン/キャノピー

記事本文:
GitHub - juliensimon/canopy: クロード コードの並列セッションにより、より高速に出荷できます。1 つのネイティブ macOS ウィンドウ、git ワークツリー、サンドボックス、自動再開、マージと終了、トークン ダッシュボードです。 · GitHub
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
ジュリエンシモン
/
キャノピー
公共
通知
あなたはきっとそうでしょう

サインインして通知設定を変更する
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
179 コミット 179 コミット .githooks .githooks .github/ workflows .github/ workflows Canopy.xcodeproj Canopy.xcodeproj Canopy Canopy リソース リソース テスト テスト ドキュメント ドキュメント スクリプト スクリプト .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CLAUDE.md CLAUDE.md ExportOptions.plist ExportOptions.plist ライセンス ライセンス Package.resolved Package.resolved Package.swift Package.swift README.md README.md TODO.md TODO.md VERSION VERSION codecov.yml codecov.yml project.yml project.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ネイティブ macOS アプリである git worktree を使用した並列クロード コード セッション。
窓がひとつ。平行な分岐。パラレルクロード。コンテキスト切り替えゼロ。
ユニバーサルバイナリ · Apple Silicon + Intel · 公証済み · macOS 14+
私は自分のプロジェクトで何ヶ月も毎日 Canopy を使用してきました。私のドックにその場所を確保したので、今共有しています。あなたのドックにも場所を確保できると思います。
正直に言いますと、クロード・コードは超大国です。キャノピーは、5 つのインフィニティ ストーンをすべて身に着けているような気分になります。
クロード コードを実験だけでなく実際の運用作業にも真剣に使用すると、ツールが乗り越えるように設計されていない壁にぶつかることになります。クロードは、1 つのディレクトリ内の 1 つのことに集中するのが得意です。しかし、実際のエンジニアリング作業は、一度に 1 つずつ行われるわけではありません。リファクタリング中にバグレポートが届きます。テストを作成しているときに、チームメイトから変更をレビューするように頼まれました。ロードマップでは、不安定なテストを調査している間も出荷を続けることが求められています。あなたは一度に 3 つのことに取り組む必要がありますが、クロードと git はあなたに 1 つのことに取り組むことを望んでいます。
ハックできます

その周辺には、スタッシュ、チェックアウト、おそらく 2 番目のクローン、おそらく git worktree add と .env をコピーするための手書きのシェル スクリプト、おそらく 6 つの Terminal.app タブのうちどのタブがどのセッションを実行していたかを覚えておく必要があります。最終的には、タブの墓場、古いワークツリーがディスクに散らばり、クロード セッションが見つからず、実行するすべてのタスクに静かな税金が課せられることになります。
Canopy はその税金を取り除きます。完全に。
Canopy のすべての機能は、手動で何かを行うことにうんざりしていたために存在しました。それぞれ数秒を節約できます。 1 日、1 週間、1 か月にわたって計算すると、その化合物は本物になります。
これらはいずれも、それ自体では大きな問題ではありません。それらはすべて切り絵です。 Canopy は、紙切れに気づく人のためのツールです。
brew install --cask juliensimon/canopy/canopy
直接ダウンロード
最新の署名済み、公証済みの .dmg を Releases から取得します。開いて、「アプリケーション」にドラッグし、起動します。
要件: macOS 14 Sonoma 以降。 Appleシリコンとかインテルとか。クロード コードがインストールされています ( $PATH でクロードが利用可能です)。
プロジェクトを追加します ( Cmd+Shift+P ) — 任意の git リポジトリをポイントします。コピーするファイル ( .env )、シンボリックリンク パス ( node_modules )、およびセットアップ コマンド ( npm install ) を構成します。
ワークツリー セッションを作成します ( Cmd+Shift+T ) — ベース ブランチを選択し、機能ブランチに名前を付けます。
Canopy はワークツリーを作成し、構成をコピーし、セットアップを実行して、Claude Code を起動します。プロンプトを開始します。
並列タスクが必要ですか?もう一度 Cmd+Shift+T を押します。各セッションは完全に分離されています。
セッションが終了したら、右クリック → [結合して終了] を選択します。 Canopy はブランチをマージし、ワークツリーをクリーンアップします。
それがループ全体です。永遠に繰り返します。
🪟 並列セッション、1 つのウィンドウ
各タブは、独自のクロード コード インスタンスを実行する個別のワークツリーです。 Cmd+1 – Cmd+9 でそれらを切り替えます。タブをドラッグして順序を変更します。 Cmd+ を使用して、名前、プロジェクト、作成日、またはディレクトリで並べ替えます

Shift+S 。アクティビティ ドットは、どのセッションがストリーミングを出力したかを示すため、どのクロードがまだ考えているかが一目でわかります。
Canopy を閉じて再度開くと、すべてのセッションが戻り、クロードの会話が自動的に再開されます。セッション ID を覚えたり、--resume フラグを入力したりする必要はありません。 Canopy は ~/.claude/projects/ をスキャンして適切なセッションを見つけ、それを claude に渡します。
📊 アクティビティ ダッシュボード — トークンの行き先を把握する
トークンの使用法は、Claude Code で十分に表面化していない点の 1 つです。キャノピーはそれを修正します。
アクティビティ ビューは、~/.claude/projects/ JSONL ファイルを解析し、Claude をどのように使用したかの全体像を示します。
常時トークンの合計 — これまでに実行したすべてのセッションにわたる入力と出力
過去 12 週間 – 同じ内訳なので、最近の傾向がわかります
セッション数 - ウィンドウ内で行った会話の数
最も忙しい日 — クロードに夢中になったのはいつですか?
モデルの内訳 - Opus、Sonnet、Haiku の割合の分割
時間ごとのヒートマップ — 12 週間分の実際の労働時間を視覚化
これは、「今月の API 予算は順調に進んでいるか」、「いつが最も生産性が高いか」に答えるために私が使用しているビューです。サードパーティのツール、スクレイピング、推定は必要ありません。Claude Code が書き込むのと同じ JSONL ファイルを読み取ります。
⌨️ コマンドパレット — すべてをあいまい検索します
Cmd+K は、Canopy が認識しているすべてのセッション、プロジェクト、ブランチ、およびアクションのあいまい一致パレットを開きます。ブランチ名を 3 文字入力し、Return キーを押すと、ブランチ名が表示されます。プロジェクトの名前を入力し、Return キーを押すと、プロジェクト ビューが開きます。 「merge」と入力して return キーを押すと、現在のセッションでマージ フローが起動します。
4 つまたは 5 つを超えるセッションが開いている場合、これが最も速く移動できる方法です。クリックするよりも速いです。ミッションコントロールより速い。インスタントに感じるほどの速さです。
🔎 で見つける

ターミナル — スクロールを停止します
任意のセッション内で Cmd+F を押すと、端末出力に対する増分検索が開きます。入力時に一致がハイライトされます。次の試合にジャンプします。 Shift-Return で前にジャンプします。エスケープすると検索が終了します。
これは小さな機能ですが、後で非常に重要であることがわかります。クロードが 400 行の計画を作成し、「データベースの移行」に関する部分にジャンプする必要がある場合、以前はスクロールする必要がありました。今はそうではありません。
📜 トランスクリプトを表示 — 端末ではなく会話を読みます
任意のセッションを右クリック → [トランスクリプトを表示…] を選択すると、会話全体がクリーンでスクロール可能な読み取り専用ビューで表示されます。 Claude Code の実行中、Canopy は構造化された JSONL セッション ログを読み取り、マークダウン形式でユーザー/アシスタントのターンをレンダリングします。ツール呼び出しは、生の出力の壁ではなく、1 行の概要に圧縮されます。クロードのストリーミングに応じてライブ更新され、自動テール切り替え機能があるため、元に戻されることなく履歴を読むことができます。コピー ボタン ( Cmd+Shift+C ) を押すと、書式設定されたマークダウンがクリップボードに保存されます。PR の説明やメモに便利です。
⬓ 分割ターミナル — クロードが上、シェルが下
Cmd+Shift+D を押すと、クロードのターミナルの下にある 2 番目のシェル ペインが切り替わります。クロードが考えている間に git status を実行する必要がありますか?別のツールからのテスト出力を覗いてみませんか?ログを追跡しますか?クロードの話を中断したり、新しいウィンドウを開いたりする必要はありません。分割ペインは、同じワークツリーをスコープとする完全な対話型シェルであり、表示したのと同じ方法で非表示にすることができます。
これは私が使用することに最も懐疑的だった機能であり、今ではこれなしで作業することは考えられません。
🛡 サンドボックス — クロードを単独で実行する
[設定] でサンドボックス バックエンド (グローバル、プロジェクトごと、または [新しいワークツリー セッション] シートのセッションごと) を選択して、VM 内でクロードを起動します。作業ディレクトリはサンドボックスにバインドマウントされているため、ファイルは編集できます。

通常は問題ありませんが、明示的にマウントされていないもの (SSH キー、ドキュメント、キーチェーン、他のリポジトリ、ホーム ディレクトリの残りの部分) はすべてアクセスできません。シールド アイコンはサイドバーでサンドボックス化されたセッションをマークし、分割ターミナルは依然として実際のファイル システムを検査するためのホスト シェルを開きます。 Canopy はバックエンドを有効にする前に必要なツールを検証します。ユーザー ガイドには、何が保護され、何が保護されないかが正確に説明されています (マウントされたプロジェクトと送信ネットワークは意図的なトレードオフです)。
Docker Sandbox (sbx) — sbx run を介した Docker Sandbox microVM。セッションの再開は無効になっています (セッション ファイルは一時的な microVM 内に存在します)。
要件: Docker Desktop および sbx ( brew install docker/tap/sbx )。
Apple コンテナ — Apple のオープンソース コンテナ ランタイムを介した軽量 VM。Docker Desktop は必要ありません。 Canopy は、ワークツリー、プロジェクトのメイン リポジトリ、および ~/.claude 状態をホスト パスにマウントします。そのため、sbx とは異なり、git はサンドボックス内で動作し、セッション再開は機能します。デフォルトの canopy-claude イメージは、ワンクリックで構築されます ([設定] → [イメージの構築])。最初のサンドボックス セッション内の 1 回限りの /login により、資格情報が設定されます。
要件: Apple シリコン上の macOS 26 以降、brew install コンテナー、コンテナー システムは起動ごとに 1 回起動します。完全なセットアップ手順はユーザー ガイドに記載されています。
✅ マージと終了 — ワンクリックで 1 つのワークツリーが廃止されます
すべての機能の最後のステップは、かつては 4 つのコマンドのダンスでした。
git チェックアウト メイン
git プル
git merge feat/何でも
git worktree 削除 ../canopy-worktrees/myproject/feat-whatever
git ブランチ -d feat/何でも
…ただし、実際にすべてを実行したわけではなく、ラップトップ上に 11 個の古いワークツリーと 40 個のマージされたブランチが存在します。キャノピーは全体を2段階確認シートに置き換えます。フェーズ 1: ターゲット ブランチを選択し、コミット数を確認し、

衝突が起こる前にそれを防ぎます。フェーズ 2: クリーンアップする対象を選択します。ワークツリー ディレクトリ、機能ブランチ、その両方、またはどちらもクリーンアップしません。ワンクリック。終わり。借金はありません。
🌳 プロジェクトビュー — すべてのワークツリーを一度に表示
サイドバーの任意のプロジェクトをクリックすると、すべてのワークツリー、そのブランチ、ステータスが表示され、ワンクリックでプロジェクトを開いたり、結合したり、削除したりできるボタンが表示されます。 「すべて開く」は、非アクティブなすべてのワークツリーを以前のクロード セッションから一度に再開します。これは、週末の休暇後に複数ブランチのプロジェクトに戻る最速の方法です。
プロジェクト ビューには、gh pr list 経由でプルされたリポジトリのオープンなプル リクエストもすべてリストされます。そのため、どのワークツリーにすでに PR が実行中で、どれがまだローカルにあるのかが一目でわかります。
🔀 Git の認識 — 常に表示されるリポジトリの状態
Canopy は 10 秒ごとに git と gh をポーリングするため、現在のワークツリーの状態を確認するためにシェルにドロップする必要はありません。ご覧のとおり:
アクティブなセッションのウィンドウ下部のステータス バー: + / - の行数を持つ変更されたファイル、アップストリームに先行するコミット、およびオープンなプル リクエストの数 (ドラフト分割あり)。任意の錠剤にカーソルを合わせると、ファイル リスト、プッシュ ステータス、PR タイトルなどの完全なツールチップが表示されます。
サイドバーのセッション行には、コンパクトな +N / −N diffstat と先行コミットの上矢印カウントが表示されるため、すべてのワークツリーを一度にスキャンできます。
プロジェクト詳細ビューには、開いているすべての PR がタイトルと番号とともにリストされます。

[切り捨てられた]

## Original Extract

Ship faster with parallel Claude Code sessions — one native macOS window, git worktrees, sandboxes, auto-resume, merge & finish, token dashboard. - juliensimon/canopy

GitHub - juliensimon/canopy: Ship faster with parallel Claude Code sessions — one native macOS window, git worktrees, sandboxes, auto-resume, merge & finish, token dashboard. · GitHub
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
juliensimon
/
canopy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
179 Commits 179 Commits .githooks .githooks .github/ workflows .github/ workflows Canopy.xcodeproj Canopy.xcodeproj Canopy Canopy Resources Resources Tests Tests docs docs scripts scripts .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CLAUDE.md CLAUDE.md ExportOptions.plist ExportOptions.plist LICENSE LICENSE Package.resolved Package.resolved Package.swift Package.swift README.md README.md TODO.md TODO.md VERSION VERSION codecov.yml codecov.yml project.yml project.yml View all files Repository files navigation
Parallel Claude Code sessions with git worktrees — a native macOS app.
One window. Parallel branches. Parallel Claudes. Zero context switching.
Universal binary · Apple Silicon + Intel · Notarized · macOS 14+
I've been using Canopy every day for months on my own projects. I'm sharing it now because it has earned its place on my dock, and I think it might earn a place on yours too.
Here's the honest pitch: Claude Code is a superpower. Canopy feels like wearing all five Infinity Stones.
When you use Claude Code seriously — not just for experiments, but for real production work — you hit a wall the tool wasn't designed to push through. Claude is brilliant at focusing on one thing in one directory. But real engineering work doesn't come one thing at a time. A bug report lands while you're refactoring. A teammate asks you to review a change while you're writing tests. The roadmap demands you keep shipping while you investigate a flaky test. You need to work on three things at once, but Claude — and git — want you to work on one.
You can hack around it: stash, checkout, maybe a second clone, maybe git worktree add and a hand-written shell script to copy your .env , maybe remember which of your six Terminal.app tabs was running which session. You end up with a tab graveyard, stale worktrees littering your disk, Claude sessions you can't find, and a quiet tax on every task you do.
Canopy removes that tax. Completely.
Every Canopy feature exists because I was tired of doing something manually. Each one saves a few seconds. Do the math over a day — over a week, over a month — and the compound is real.
None of these are big problems on their own. All of them are papercuts. Canopy is a tool for people who notice papercuts.
brew install --cask juliensimon/canopy/canopy
Direct download
Grab the latest signed and notarized .dmg from Releases . Open, drag to Applications, launch.
Requirements: macOS 14 Sonoma or later. Apple Silicon or Intel. Claude Code installed ( claude available in your $PATH ).
Add a project ( Cmd+Shift+P ) — point at any git repository. Configure files to copy ( .env ), symlink paths ( node_modules ), and setup commands ( npm install ).
Create a worktree session ( Cmd+Shift+T ) — pick a base branch, name your feature branch.
Canopy creates the worktree, copies config, runs your setup, and launches Claude Code. Start prompting.
Need a parallel task? Cmd+Shift+T again. Each session is completely isolated.
When you're done with a session, right-click → Merge & Finish . Canopy merges your branch and cleans up the worktree.
That's the whole loop. Repeat forever.
🪟 Parallel sessions, one window
Each tab is a separate worktree running its own Claude Code instance. Switch between them with Cmd+1 – Cmd+9 . Drag tabs to reorder. Sort by name, project, creation date, or directory with Cmd+Shift+S . Activity dots show which sessions have output streaming, so you know at a glance which Claude is still thinking.
When you close and reopen Canopy, every session comes back — with its Claude conversation resumed automatically . No session IDs to remember, no --resume flags to type. Canopy finds the right session by scanning ~/.claude/projects/ and passes it to claude for you.
📊 Activity dashboard — know where your tokens go
Token usage is the one thing Claude Code doesn't surface well. Canopy fixes that.
The Activity view parses your ~/.claude/projects/ JSONL files and gives you a full picture of how you've been using Claude:
All-time token totals — input and output, across every session you've ever run
Last 12 weeks — same breakdown, so you can see recent trends
Session count — how many conversations you've had in the window
Busiest day — when were you deep in Claude?
Model breakdown — percentage split across Opus, Sonnet, and Haiku
Hour-by-hour heatmap — 12 weeks of your actual working hours, visualized
This is the view I use to answer "am I on track for my API budget this month" and "when am I most productive." No third-party tools, no scraping, no estimation — it's reading the same JSONL files Claude Code writes.
⌨️ Command palette — fuzzy search everything
Cmd+K opens a fuzzy-match palette over every session, project, branch, and action Canopy knows about. Type three letters of a branch name, hit return, you're there. Type the name of a project, hit return, the project view opens. Type "merge", hit return, the merge flow fires on the current session.
If you have more than four or five sessions open, this is the fastest way to navigate. Faster than clicking. Faster than Mission Control. Fast enough to feel instant.
🔎 Find in terminal — stop scrolling
Cmd+F inside any session opens an incremental search over the terminal output. Matches highlight as you type. Return jumps to the next match. Shift-return jumps to the previous one. Escape closes the search.
This is a small feature that turns out to matter a lot: when Claude produces a 400-line plan and you need to jump to the part about "database migration," you used to scroll. Now you don't.
📜 Show Transcript — read the conversation, not the terminal
Right-click any session → Show Transcript… for a clean, scrollable, read-only view of the whole conversation. When Claude Code is running, Canopy reads its structured JSONL session log and renders user/assistant turns with markdown formatting — tool calls compacted to one-line summaries instead of walls of raw output. It live-updates as Claude streams, with an auto-tail toggle so you can read history without being yanked back down. The Copy button ( Cmd+Shift+C ) puts the formatted markdown on your clipboard — handy for PR descriptions and notes.
⬓ Split terminal — Claude up top, shell below
Cmd+Shift+D toggles a secondary shell pane below Claude's terminal. Need to run git status while Claude is mid-thought? Peek at the test output from another tool? Tail a log? You don't need to interrupt Claude or open a new window. The split pane is a full interactive shell, scoped to the same worktree, and you can hide it the same way you showed it.
This is the feature I was most skeptical I'd use, and now I can't imagine working without it.
🛡 Sandboxes — run Claude in isolation
Pick a sandbox backend in Settings (globally, per project, or per session in the New Worktree Session sheet) to launch Claude inside a VM. Your working directory is bind-mounted into the sandbox, so file edits work normally, but everything that isn't explicitly mounted — SSH keys, documents, the Keychain, other repos, the rest of your home directory — is out of reach. A shield icon marks sandboxed sessions in the sidebar, and the split terminal still opens a host shell for inspecting the real filesystem. Canopy validates the required tools before enabling a backend. The user guide spells out exactly what is and isn't protected (the mounted project and outbound network are deliberate trade-offs).
Docker Sandbox (sbx) — a Docker Sandbox microVM via sbx run . Session resume is disabled (session files live inside the ephemeral microVM).
Requirements: Docker Desktop and sbx ( brew install docker/tap/sbx ).
Apple container — a lightweight VM via Apple's open-source container runtime, no Docker Desktop needed. Canopy mounts the worktree, the project's main repo, and your ~/.claude state at their host paths — so git works inside the sandbox and session resume works , unlike sbx. The default canopy-claude image is built in one click ( Settings → Build Image ); a one-time /login inside the first sandboxed session sets up credentials.
Requirements: macOS 26+ on Apple silicon, brew install container , container system start once per boot. Full setup steps in the user guide .
✅ Merge & Finish — one click, one worktree retired
The last step of every feature used to be a four-command dance:
git checkout main
git pull
git merge feat/whatever
git worktree remove ../canopy-worktrees/myproject/feat-whatever
git branch -d feat/whatever
…except you never actually did all of it, and now you have 11 stale worktrees and 40 merged branches on your laptop. Canopy replaces the whole thing with a two-phase confirmation sheet. Phase 1: pick the target branch, confirm the commit count, see any conflicts before they happen. Phase 2: pick what to clean up — worktree directory, feature branch, both, or neither. One click. Done. No debt.
🌳 Project view — see every worktree at once
Click any project in the sidebar to see every worktree, its branch, its status, and a one-click button to open, merge, or delete it. "Open All" resumes every inactive worktree at once with their prior Claude sessions — the fastest way to get back into a multi-branch project after a weekend away.
The project view also lists every open pull request for the repository, pulled via gh pr list — so you can see at a glance which of your worktrees already have a PR in flight and which are still local.
🔀 Git awareness — always-visible repo state
Canopy polls git and gh every 10 seconds so you never have to drop into a shell to check the state of the current worktree. You see:
Status bar at the bottom of the window, for the active session: modified files with + / − line counts, commits ahead of the upstream, and open pull request count (with draft split). Hover any pill for a full tooltip — file list, push status, PR titles.
Sidebar session rows show a compact +N / −N diffstat and an up-arrow count for commits-ahead, so you can scan all your worktrees at once.
Project detail view lists every open PR with title, numbe

[truncated]
