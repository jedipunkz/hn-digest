---
source: "https://howborisusesclaudecode.com"
hn_url: "https://news.ycombinator.com/item?id=48699921"
title: "How Boris Cherny Uses Claude Code"
article_title: "How Boris Uses Claude Code"
author: "eustoria"
captured_at: "2026-06-27T17:23:55Z"
capture_tool: "hn-digest"
hn_id: 48699921
score: 1
comments: 0
posted_at: "2026-06-27T17:10:17Z"
tags:
  - hacker-news
  - translated
---

# How Boris Cherny Uses Claude Code

- HN: [48699921](https://news.ycombinator.com/item?id=48699921)
- Source: [howborisusesclaudecode.com](https://howborisusesclaudecode.com)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T17:10:17Z

## Translation

タイトル: ボリス・チェルニーがクロード・コードをどのように使用するか
記事のタイトル: ボリスはクロード コードをどのように使用するか
説明: Claude Code の作成者による、毎日の使い方に関する 107 のヒント

記事本文:
">
アントロピックとは関係ありません
claude-cli — Boris がどのようにクロード コードを使用するか
について
クウデオノミクス
Claude SF によるコード
夢を見る — 研究プレビュー
パート1 13
パート2 10
パート 3 12
パート 4 5
パート5 2
パート6 3
パート 7 8
パート8 4
パート9 15
パート 10
7
パート11
8
パート12
2
パート13
3
パート14
6
パート15
4
パート16
3
新しい
パート17
2
新しい
/bolis をインストールする — クロード コードのヒント
/thariq-skills をインストールする — スキルの書き方
～
イントロ
1
平行
2
ウェブ+モバイル
3
作品
4
クロード.md
5
@.クロード
6
計画
7
スラッシュ
8
サブエージェント
9
フック
10
パーマ
11
ツール
12
長期にわたる
13
検証する
～
イントロ
1
ワークツリー
2
計画
3
クロード.md
4
スキル
5
バグ
6
促す
7
端子
8
サブエージェント
9
データ
10
学習
～
イントロ
1
端子
2
努力
3
プラグイン
4
エージェント
5
権限
6
サンドボックス
7
ステータスライン
8
キーバインド
9
フック
10
スピナー
11
出力
12
カスタマイズする
～
イントロ
1
クリ
2
デスクトップ
3
サブエージェント
4
エージェント
5
非 Git
～
イントロ
1
/簡素化する
2
/バッチ
～
イントロ
1
/ループ
2
コードレビュー
3
/ところで
～
イントロ
1
/努力
2
リモート
3
声
4
セットアップ
5
--名前
6
自動名前
7
/色
8
ポストコンパクト
～
イントロ
1
オートモード
2
/スケジュール
3
iメッセージ
4
記憶
～
イントロ
1
モバイル
2
テレポート
3
ループ
4
フック
5
派遣
6
クロム
7
デスクトップ
8
フォーク
9
/ところで
10
ワークツリー
11
/バッチ
12
--裸
13
--追加ディレクトリ
14
--エージェント
15
/音声
～
イントロ
1
ルーティン
2
/巻き戻し
3
/compact と /clear
4
オートコンパクト
5
代表者
6
完全なコンテキスト
7
高さ
～
イントロ
1
オートモード
2
/少ないパーマ
3
要約
4
/フォーカス
5
努力
6
/行く
7
4.6→4.7
8
通知
～
イントロ
1
エージェントビュー
2
/目標
～
イントロ
1
作品4.8
2
努力+限界
3
dyn ワークフロー
～
イントロ
1
なぜワークフローなのか
2
プリミティブ
3
6つのパターン
4
ユースケース
5
ループ + 予算
6
保存+共有
～
イントロ
1
プラン→オート
2
ミニマリズム
3
それを書き留めてください
4
オートモードの安全性
～
イントロ
1
ネストされたエージェント
2
フォーク: 真
3
ワークフローを使用する
～
イントロ
1
寓話5
2
何が変わるのか
～
ガイド
～
インストールする
Boris Cherny がどのようにクロード コードを使用するか
ボー

ris は Anthropic でクロード コードを作成しました。どのように使っているか尋ねると、彼は毎日のワークフローから得た 13 の実践的なヒントを共有しました。彼のセットアップは「驚くほどバニラ」であり、CC が箱から出してすぐにうまく機能することの証拠です。
@bcherny の 2026 年 1 月 2 日のスレッド
上のタブをクリックして各ヒントを確認してください →
または ← → 矢印キーを使用します
1
5 人のクロードを並行して実行する
Boris は、同じリポジトリの 5 つの個別の git チェックアウトを使用して、ターミナルでクロード コードの 5 つのインスタンスを同時に実行します。彼は簡単に参照できるようにタブに 1 ～ 5 の番号を付け、システム通知を使用してクロードがいつ入力を必要とするかを把握します。
元の投稿を見る
← イントロ
ウェブ+モバイル →
2
Web セッションとモバイル セッションの同時実行
ターミナルを超えて、Boris は claude.ai/code で 5 ～ 10 の追加セッションを実行します。 & コマンドまたは --teleport フラグを使用して、ローカルと Web の間を流動的に切り替えます。
また、彼は午前中に Claude iOS アプリを介して携帯電話からセッションを開始し、後でコンピューターでセッションを開始します。
元の投稿を見る
← 平行
作品 →
3
Opus 4.5 あらゆることを考える
ボリスは、あらゆるタスクに思考モードを備えた Opus 4.5 を使用しています。彼の推論は次のとおりです。
重要な点: より大きなモデルであっても、ステアリングの軽減 + ツールの使用の向上 = 全体的な結果の高速化。
元の投稿を見る
← ウェブ+モバイル
クロード.md →
4
共有 CLAUDE.md ドキュメント
チームは、Git にチェックインされた Claude Code リポジトリの単一の CLAUDE.md ファイルを共有します。チーム全体が週に複数回貢献します。
元の投稿を見る
← 作品
@.クロード →
5
コードレビューの @.claude
コード レビュー中に、Boris は PR に @.claude をタグ付けして、PR 自体の一部として学習内容を CLAUDE.md に追加します。
これには、Claude Code GitHub アクション ( /install-github-action ) を使用します。ボリスは、ダン・シッパーのコンセプトに触発されて、これを彼らのバージョンの「配合エンジニアリング」と呼んでいます。
元の投稿を見る
← クロード.md
計画→
6
プランモードで開始
M

ost セッションはプラン モード (Shift+Tab を 2 回) で開始します。ボリスは、計画が固まるまでクロードとともに計画を繰り返し、その後、自動承認モードに切り替えます。
「将来の問題を回避するには、適切な計画が非常に重要です。」
元の投稿を見る
← @.クロード
スラッシュ →
7
内部ループのスラッシュ コマンド
Boris は、1 日に何度もワークフローにスラッシュ コマンドを使用します。これにより、繰り返しのプロンプトが不要になり、クロードもプロンプトを使用できるようになります。
コマンドは .claude/commands/ の下の git にチェックインされ、チームと共有されます。
元の投稿を見る
← 計画
サブエージェント →
8
共通ワークフローのサブエージェント
Boris は、サブエージェントを最も一般的な PR ワークフローの自動化として考えています。
元の投稿を見る
← スラッシュ
フック →
9
PostTool書式設定にフックを使用する
チームは PostToolUse フックを使用して、Claude のコードを自動フォーマットします。 Claude は 90% の確率で適切にフォーマットされたコードを生成しますが、フックはエッジ ケースをキャッチして CI の失敗を防ぎます。
元の投稿を見る
← サブエージェント
パーマ→
10
安全なアクセス許可を事前に許可する
--dangerously-skip-permissions の代わりに、Boris は /permissions を使用して一般的な安全なコマンドを事前に許可します。ほとんどは .claude/settings.json で共有されます。
元の投稿を見る
← フック
ツール →
11
ツールの統合
Claude Code は、Boris のすべてのツールを自律的に使用します。
検索と Slack への投稿 (MCP サーバー経由)
bq CLI を使用して BigQuery クエリを実行します
元の投稿を見る
← パーマ
ロングラン→
12
長時間実行されるタスクを処理する
非常に長時間実行されるタスクの場合、ボリスはクロードが中断されずに作業できることを保証します。
サンドボックス環境の場合、ブロックを回避するために --permission-mode=dontAsk または --dangerously-skip-permissions を使用します。
元の投稿を見る
← ツール
確認→
13
最も重要なヒント: 検証
これが Boris の 1 番のヒントです。
claude.ai/code に対する独自の変更では、Claude は Claude Chrome 拡張機能を使用してブラウザを開き、UI の変更をテストし、変更が完了するまで繰り返します。

完璧に機能します。
検証は、Bash コマンド、テスト スイート、シミュレーター、ブラウザ テストなどのドメインによって異なります。重要なのは、フィードバック ループを閉じる方法をクロードに与えることです。
元の投稿を見る
← 長期的
イントロに戻る→
Boris Cherny からのその他のヒント
Boris は 2026 年 1 月 31 日にさらに 10 個のヒントを共有しました。これらは Claude Code チームから直接提供されたものです。覚えておいてください、設定は人それぞれ異なります。試してみて何が効果的かを確認してください。
@bcherny の 2026 年 1 月 31 日のスレッド
上のタブをクリックして各ヒントを確認してください →
または ← → 矢印キーを使用します
1
より多くの作業を並行して行う
3 ～ 5 個の git worktree を一度にスピンアップし、それぞれが独自の Claude セッションを並行して実行します。これは生産性の最大の解放であり、チームからの最高のヒントです。
ワークツリーに名前を付け、シェル エイリアス (za、zb、zc) を設定して、1 回のキーストロークでワークツリー間を移動できるようにする人もいます。ログの読み取りと BigQuery の実行のみを目的とした専用の「分析」ワークツリーを備えているものもあります。
元の投稿を見る
← イントロ
計画→
2
すべての複雑なタスクを計画モードで開始する
クロードが一発で実行できるよう、計画にエネルギーを注ぎましょう。物事がうまくいかなくなっても無理をせず、計画モードに戻って計画を立て直してください。
また、ビルドだけでなく、検証ステップのために計画モードに入るようにクロードに明示的に指示します。
元の投稿を見る
← ワークツリー
クロード.md →
3
CLAUDE.md に投資してください
修正するたびに、「同じ間違いを繰り返さないように、CLAUDE.md を更新してください。」で終了します。クロードは自分自身でルールを書くのが不気味なほど上手です。
あるエンジニアはクロードに、すべてのタスク/プロジェクトのメモ ディレクトリを維持し、PR のたびに更新するように指示します。次に、彼らは CLAUDE.md をそれに向けます。
元の投稿を見る
← 計画
スキル→
4
独自のスキルを作成する
独自のスキルを作成し、 git にコミットします。すべてのプロジェクトで再利用します。
もっと何かやったら

an once a day, turn it into a skill or command
/techdebt スラッシュ コマンドを作成し、各セッションの最後に実行して、重複したコードを見つけて強制終了します。
7 日間の Slack、GDrive、Asana、GitHub を 1 つのコンテキスト ダンプに同期するスラッシュ コマンドを設定する
dbt モデルを作成し、コードをレビューし、開発環境で変更をテストする分析エンジニア スタイルのエージェントを構築する
元の投稿を見る
← クロード.md
バグ →
5
Claude Fixes Most Bugs by Itself
Slack MCP を有効にしてから、Slack のバグ スレッドを Claude に貼り付けて、「修正」と言うだけです。 Zero context switching required.
Or, just say "Go fix the failing CI tests."方法を細かく管理しないでください。
元の投稿を見る
← スキル
促す →
6
プロンプトをレベルアップする
ａ．クロードに挑戦します。 「これらの変更についてはよく理解してください。テストに合格するまでは自己PRをしないでください。」と言いましょう。 Make Claude be your reviewer.または、「これが機能することを証明してください」と言って、Claude に main と機能ブランチの間で動作の差分を与えます。
b.平凡な修正を行った後、「今わかっていることをすべて理解したので、これを破棄して、洗練されたソリューションを実装してください。」と言います。
c.作業を引き継ぐ前に詳細な仕様を作成し、曖昧さを減らします。具体的であればあるほど、より良い結果が得られます。
元の投稿を見る
← バグ
端子→
7
Terminal & Environment Setup
チームは Ghostty を愛しています!同期レンダリング、24 ビット カラー、適切な Unicode サポートが気に入っている人はたくさんいます。
クロードジャグリングを簡単にするには、/statusline を使用してステータス バーをカスタマイズし、コンテキストの使用状況と現在の git ブランチを常に表示します。私たちの多くは、ターミナル タブを色分けして名前を付けており、場合によっては tmux (タスク/ワークツリーごとに 1 つのタブ) を使用します。
元の投稿を見る
← 促す
サブエージェント →
8
サブエージェントの使用
ａ．クロードが問題に対してより多くの計算を投入できるようにするリクエストには、「サブエージェントの使用」を追加します。
b.メインエージェントのコンテキストを維持するために、個々のタスクをサブエージェントにオフロードします。

ウィンドウがきれいで集中しています。
c.許可リクエストをフック経由で Opus 4.5 にルーティングします。攻撃をスキャンし、安全なものを自動承認します。
元の投稿を見る
← ターミナル
データ→
9
データと分析にクロードを使用する
Claude Code に「bq」CLI を使用してメトリクスをその場で取得して分析するように依頼します。 BigQuery スキルがコードベースにチェックインされており、チームの全員がそれをクロード コードで直接分析クエリに使用しています。
これは、CLI、MCP、または API を備えたあらゆるデータベースで機能します。
元の投稿を見る
← サブエージェント
学習→
10
クロードと一緒に学ぶ
Claude Code を学習に使用するためのチームからのヒント:
ａ． /config で「説明」または「学習」出力スタイルを有効にして、クロードに変更の背後にある理由を説明させます。
b.クロードに、なじみのないコードを説明する視覚的な HTML プレゼンテーションを生成してもらいます。驚くほど良いスライドが作れます！
c.新しいプロトコルとコードベースを理解するために、クロードに ASCII 図を描いてもらいます。
d.間隔をあけて反復する学習スキルを構築します。あなたが自分の理解を説明し、クロードがギャップを埋めるためにフォローアップを求め、結果を保存します。
元の投稿を見る
← データ
イントロに戻る→
クロードをカスタマイズ
Boris は、2026 年 2 月 11 日にさらに 12 のヒントを共有しました。今回のテーマはカスタマイズです。フック、プラグイン、エージェント、権限、およびクロード コードを独自のものにするためのすべての方法です。
@bcherny の 2026 年 2 月 11 日のスレッド
上のタブをクリックして各ヒントを確認してください →
または ← → 矢印キーを使用します
1
端末を設定する
Claude コードを端末で適切に使用できるようにするためのいくつかの簡単な設定:
テーマ: /config を実行してライト/ダーク モードを設定する
通知: iTerm2 の通知を有効にするか、カスタム通知フックを使用します。
改行: IDE ターミナル、Apple ターミナル、Warp、または Alacritty でクロード コードを使用する場合は、/terminal-setup を実行して改行の Shift+Enter を有効にします (そのため、タイプする必要はありません)

e\)
元の投稿を見る
← イントロ
努力→
2
努力レベルを調整する
/model を実行して、希望する作業レベルを選択します。
低 — トークンが少なく、応答が速い
高 — より多くのトークンとより多くのインテリジェンス
個人的には、すべてに High を使用します。
元の投稿を見る
← ターミナル
プラグイン→
3
プラグイン、MCP、スキルをインストールする
プラグインを使用すると、LSP (現在はすべての主要言語で利用可能)、MCP、スキル、エージェント、カスタム フックをインストールできます。
公式の Anthropic プラグイン マーケットプレイスからプラグインをインストールするか、会社用に独自のマーケットプレイスを作成します。次に、settings.json をコードベースにチェックインして、チームのマーケットプレイスを自動追加します。
元の投稿を見る
← 努力
エージェント →
4
カスタムエージェントの作成
カスタム エージェントを作成するには、.md ファイルを .claude/agents にドロップします。各エージェントには、カスタム名、色、ツール セット、事前に許可されたツールと事前に禁止されたツール、権限モード、モデルを設定できます。
元の投稿を見る
← プラグイン
権限 →
5
共通権限の事前承認
Claude Code は、プロンプト インジェクション検出、静的分析、サンドボックス、人間による監視を組み合わせた高度な権限システムを使用しています。
すぐに使える安全なコマンドの小さなセットが事前に承認されています。さらに事前承認するには、/permissions を実行し、許可リストとブロック リストに追加します。これらをチームの settings.json にチェックインします。
元の投稿を見る
← エージェント
サンドボックス →
6
サンドボックスを有効にする
クロード・コードのオープンにオプトインする

[切り捨てられた]

## Original Extract

107 tips from the creator of Claude Code on how he uses it daily

">
not affiliated with Anthropic
claude-cli — How Boris Uses Claude Code
About
Claudeonomics
Code w/Claude SF
Dreaming — Research Preview
Part 1 13
Part 2 10
Part 3 12
Part 4 5
Part 5 2
Part 6 3
Part 7 8
Part 8 4
Part 9 15
Part 10
7
Part 11
8
Part 12
2
Part 13
3
Part 14
6
Part 15
4
Part 16
3
New
Part 17
2
New
Install /boris — Claude Code tips
Install /thariq-skills — How to write skills
~
intro
1
parallel
2
web+mobile
3
opus
4
CLAUDE.md
5
@.claude
6
plan
7
slash
8
subagents
9
hooks
10
perms
11
tools
12
long-run
13
verify
~
intro
1
worktrees
2
plan
3
CLAUDE.md
4
skills
5
bugs
6
prompting
7
terminal
8
subagents
9
data
10
learning
~
intro
1
terminal
2
effort
3
plugins
4
agents
5
permissions
6
sandbox
7
statusline
8
keybindings
9
hooks
10
spinners
11
output
12
customize
~
intro
1
cli
2
desktop
3
subagents
4
agents
5
non-git
~
intro
1
/simplify
2
/batch
~
intro
1
/loop
2
code-review
3
/btw
~
intro
1
/effort
2
remote
3
voice
4
setup
5
--name
6
auto-name
7
/color
8
postcompact
~
intro
1
auto mode
2
/schedule
3
iMessage
4
memory
~
intro
1
mobile
2
teleport
3
loops
4
hooks
5
dispatch
6
chrome
7
desktop
8
fork
9
/btw
10
worktrees
11
/batch
12
--bare
13
--add-dir
14
--agent
15
/voice
~
intro
1
routines
2
/rewind
3
/compact vs /clear
4
auto-compact
5
delegate
6
full context
7
xhigh
~
intro
1
auto mode
2
/fewer-perms
3
recaps
4
/focus
5
effort
6
/go
7
4.6→4.7
8
notifications
~
intro
1
agent view
2
/goal
~
intro
1
opus 4.8
2
effort + limits
3
dyn workflows
~
intro
1
why workflows
2
primitives
3
six patterns
4
use cases
5
loop + budgets
6
save + share
~
intro
1
plan → auto
2
minimalism
3
write it down
4
auto mode safety
~
intro
1
nested agents
2
fork: true
3
use a workflow
~
intro
1
fable 5
2
what changes
~
guide
~
install
How Boris Cherny Uses Claude Code
Boris created Claude Code at Anthropic. When asked how he uses it, he shared 13 practical tips from his daily workflow. His setup is "surprisingly vanilla" — proof that CC works great out of the box.
@bcherny's January 2, 2026 thread
Click the tabs above to explore each tip →
or use ← → arrow keys
1
Run 5 Claudes in Parallel
Boris runs 5 instances of Claude Code simultaneously in his terminal using 5 separate git checkouts of the same repo. He numbers his tabs 1-5 for easy reference and uses system notifications to know when any Claude needs input.
View original post
← intro
web+mobile →
2
Parallel Web and Mobile Sessions
Beyond the terminal, Boris runs 5-10 additional sessions on claude.ai/code . He fluidly hands off between local and web using the & command or --teleport flag.
He also kicks off sessions from his phone via the Claude iOS app in the morning, then picks them up on his computer later.
View original post
← parallel
opus →
3
Opus 4.5 with Thinking for Everything
Boris uses Opus 4.5 with thinking mode for every task. His reasoning:
The takeaway: less steering + better tool use = faster overall results , even with a larger model.
View original post
← web+mobile
CLAUDE.md →
4
Shared CLAUDE.md Documentation
The team shares a single CLAUDE.md file for the Claude Code repo, checked into git. The whole team contributes multiple times a week.
View original post
← opus
@.claude →
5
@.claude in Code Reviews
During code review, Boris tags @.claude on PRs to add learnings to the CLAUDE.md as part of the PR itself.
They use the Claude Code GitHub Action ( /install-github-action ) for this. Boris calls it their version of "Compounding Engineering" — inspired by Dan Shipper's concept.
View original post
← CLAUDE.md
plan →
6
Start in Plan Mode
Most sessions start in Plan mode (shift+tab twice). Boris iterates on the plan with Claude until it's solid, then switches to auto-accept mode.
"A good plan is really important to avoid issues down the line."
View original post
← @.claude
slash →
7
Slash Commands for Inner Loops
Boris uses slash commands for workflows he does many times a day . This saves repeated prompting, and Claude can use them too.
Commands are checked into git under .claude/commands/ and shared with the team.
View original post
← plan
subagents →
8
Subagents for Common Workflows
Boris thinks of subagents as automations for the most common PR workflows :
View original post
← slash
hooks →
9
PostToolUse Hooks for Formatting
The team uses a PostToolUse hook to auto-format Claude's code. While Claude generates well-formatted code 90% of the time, the hook catches edge cases to prevent CI failures.
View original post
← subagents
perms →
10
Pre-Allow Safe Permissions
Instead of --dangerously-skip-permissions , Boris uses /permissions to pre-allow common safe commands. Most are shared in .claude/settings.json .
View original post
← hooks
tools →
11
Tool Integrations
Claude Code uses all of Boris's tools autonomously:
Searches and posts to Slack (via MCP server)
Runs BigQuery queries with bq CLI
View original post
← perms
long-run →
12
Handle Long-Running Tasks
For very long-running tasks, Boris ensures Claude can work uninterrupted:
For sandboxed environments, he'll use --permission-mode=dontAsk or --dangerously-skip-permissions to avoid blocks.
View original post
← tools
verify →
13
The Most Important Tip: Verification
This is Boris's #1 tip:
For his own changes to claude.ai/code, Claude uses the Claude Chrome extension to open a browser, test UI changes, and iterate until it works perfectly.
Verification varies by domain: Bash commands, test suites, simulators, browser testing, etc. The key is giving Claude a way to close the feedback loop.
View original post
← long-run
back to intro →
More Tips from Boris Cherny
Boris shared 10 more tips on January 31, 2026. These are sourced directly from the Claude Code team — remember, everyone's setup is different. Experiment to see what works for you!
@bcherny's January 31, 2026 thread
Click the tabs above to explore each tip →
or use ← → arrow keys
1
Do More in Parallel
Spin up 3-5 git worktrees at once , each running its own Claude session in parallel. It's the single biggest productivity unlock, and the top tip from the team.
Some people also name their worktrees and set up shell aliases (za, zb, zc) so they can hop between them in one keystroke. Others have a dedicated "analysis" worktree that's only for reading logs and running BigQuery.
View original post
← intro
plan →
2
Start Every Complex Task in Plan Mode
Pour your energy into the plan so Claude can 1-shot the implementation . Don't keep pushing when things go sideways — switch back to plan mode and re-plan.
They also explicitly tell Claude to enter plan mode for verification steps, not just for the build .
View original post
← worktrees
CLAUDE.md →
3
Invest in Your CLAUDE.md
After every correction, end with: "Update your CLAUDE.md so you don't make that mistake again." Claude is eerily good at writing rules for itself.
One engineer tells Claude to maintain a notes directory for every task/project , updated after every PR. They then point CLAUDE.md at it.
View original post
← plan
skills →
4
Create Your Own Skills
Create your own skills and commit them to git . Reuse across every project.
If you do something more than once a day, turn it into a skill or command
Build a /techdebt slash command and run it at the end of every session to find and kill duplicated code
Set up a slash command that syncs 7 days of Slack, GDrive, Asana, and GitHub into one context dump
Build analytics-engineer-style agents that write dbt models, review code, and test changes in dev
View original post
← CLAUDE.md
bugs →
5
Claude Fixes Most Bugs by Itself
Enable the Slack MCP , then paste a Slack bug thread into Claude and just say "fix." Zero context switching required.
Or, just say "Go fix the failing CI tests." Don't micromanage how.
View original post
← skills
prompting →
6
Level Up Your Prompting
a. Challenge Claude. Say "Grill me on these changes and don't make a PR until I pass your test." Make Claude be your reviewer. Or, say "Prove to me this works" and have Claude diff behavior between main and your feature branch.
b. After a mediocre fix, say: "Knowing everything you know now, scrap this and implement the elegant solution."
c. Write detailed specs and reduce ambiguity before handing work off. The more specific you are, the better the output.
View original post
← bugs
terminal →
7
Terminal & Environment Setup
The team loves Ghostty ! Multiple people like its synchronized rendering, 24-bit color, and proper unicode support.
For easier Claude-juggling, use /statusline to customize your status bar to always show context usage and current git branch. Many of us also color-code and name our terminal tabs, sometimes using tmux — one tab per task/worktree.
View original post
← prompting
subagents →
8
Use Subagents
a. Append "use subagents" to any request where you want Claude to throw more compute at the problem.
b. Offload individual tasks to subagents to keep your main agent's context window clean and focused.
c. Route permission requests to Opus 4.5 via a hook — let it scan for attacks and auto-approve the safe ones.
View original post
← terminal
data →
9
Use Claude for Data & Analytics
Ask Claude Code to use the "bq" CLI to pull and analyze metrics on the fly. We have a BigQuery skill checked into the codebase, and everyone on the team uses it for analytics queries directly in Claude Code.
This works for any database that has a CLI, MCP, or API .
View original post
← subagents
learning →
10
Learning with Claude
A few tips from the team to use Claude Code for learning:
a. Enable the "Explanatory" or "Learning" output style in /config to have Claude explain the why behind its changes.
b. Have Claude generate a visual HTML presentation explaining unfamiliar code. It makes surprisingly good slides!
c. Ask Claude to draw ASCII diagrams of new protocols and codebases to help you understand them.
d. Build a spaced-repetition learning skill : you explain your understanding, Claude asks follow-ups to fill gaps, stores the result.
View original post
← data
back to intro →
Customize Your Claude
Boris shared 12 more tips on February 11, 2026. This time the theme is customization — hooks, plugins, agents, permissions, and all the ways to make Claude Code your own.
@bcherny's February 11, 2026 thread
Click the tabs above to explore each tip →
or use ← → arrow keys
1
Configure Your Terminal
A few quick settings to make Claude Code feel right in your terminal:
Theme: Run /config to set light/dark mode
Notifications: Enable notifications for iTerm2, or use a custom notifs hook
Newlines: If you use Claude Code in an IDE terminal, Apple Terminal, Warp, or Alacritty, run /terminal-setup to enable shift+enter for newlines (so you don't need to type \ )
View original post
← intro
effort →
2
Adjust Effort Level
Run /model to pick your preferred effort level:
Low — less tokens & faster responses
High — more tokens & more intelligence
Personally, I use High for everything .
View original post
← terminal
plugins →
3
Install Plugins, MCPs, and Skills
Plugins let you install LSPs (now available for every major language), MCPs, skills, agents, and custom hooks .
Install a plugin from the official Anthropic plugin marketplace, or create your own marketplace for your company. Then, check the settings.json into your codebase to auto-add the marketplaces for your team.
View original post
← effort
agents →
4
Create Custom Agents
To create custom agents, drop .md files in .claude/agents . Each agent can have a custom name, color, tool set, pre-allowed and pre-disallowed tools, permission mode, and model.
View original post
← plugins
permissions →
5
Pre-Approve Common Permissions
Claude Code uses a sophisticated permission system with a combo of prompt injection detection, static analysis, sandboxing, and human oversight .
Out of the box, we pre-approve a small set of safe commands. To pre-approve more, run /permissions and add to the allow and block lists. Check these into your team's settings.json .
View original post
← agents
sandbox →
6
Enable Sandboxing
Opt into Claude Code's open

[truncated]
