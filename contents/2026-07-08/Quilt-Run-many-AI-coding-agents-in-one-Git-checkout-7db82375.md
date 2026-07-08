---
source: "https://github.com/wkoverfield/quilt"
hn_url: "https://news.ycombinator.com/item?id=48832631"
title: "Quilt: Run many AI coding agents in one Git checkout"
article_title: "GitHub - wkoverfield/quilt: Tracks which agent wrote which lines when multiple AI agents share one Git checkout. · GitHub"
author: "wkoverfield"
captured_at: "2026-07-08T15:11:03Z"
capture_tool: "hn-digest"
hn_id: 48832631
score: 1
comments: 0
posted_at: "2026-07-08T14:41:03Z"
tags:
  - hacker-news
  - translated
---

# Quilt: Run many AI coding agents in one Git checkout

- HN: [48832631](https://news.ycombinator.com/item?id=48832631)
- Source: [github.com](https://github.com/wkoverfield/quilt)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T14:41:03Z

## Translation

タイトル: Quilt: 1 つの Git チェックアウトで多くの AI コーディング エージェントを実行する
記事タイトル: GitHub - wkoverfield/quilt: 複数の AI エージェントが 1 つの Git チェックアウトを共有するときに、どのエージェントがどの行を書いたかを追跡します。 · GitHub
説明: 複数の AI エージェントが 1 つの Git チェックアウトを共有する場合、どのエージェントがどの行を書き込んだかを追跡します。 - wkoverfield/キルト

記事本文:
GitHub - wkoverfield/quilt: 複数の AI エージェントが 1 つの Git チェックアウトを共有する場合に、どのエージェントがどの行を書いたかを追跡します。 · GitHub
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
ウィコオーバーフィールド
/
キルト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション o

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
227 コミット 227 コミット .github .github ベンチ ベンチ設計 設計ドキュメント ドキュメントの例 例 src src テスト test .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md Glama.json Glama.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Quilt は、どのエージェントがファイル内のどの行を書いたかを追跡するコマンドライン ツールです。
共有 Git チェックアウトにより、複数の AI コーディング エージェントが 1 つのリポジトリで同時に作業できる
そしてそれぞれは独自の変更のみをコミットします。
ツール境界でのすべての編集をキャプチャし、誰が書いたかを行ごとに記録します。
何を、コミット時に各エージェント独自の変更を再構築します。 Git はそのまま残ります
真実の源。 Quilt は LLM を呼び出したり、エージェントを生成したりすることはなく、その状態は
リポジトリに触れることなく削除できる .quilt/ サイドカー。
npm install -g @quilt-dev/cli
キルトのセットアップ # Quilt をリポジトリにワイヤリングします (クロード コード、カーソル、またはプレーン git)
問題
エージェントを 1 つのチェックアウトに含める必要があります: 1 つの node_modules 、1 つのビルド、1 つの開発
サーバー、エージェントごとにワークツリーではなく、動作し続けるための 1 つの環境。
独自のインストールと独自のドリフト。ただし、共有チェックアウトでは、単純な git バイトが必要です
エージェントがまったく異なる作業を行っている場合でも、最初の git commit -am
他の全員のコミットされていないファイルを 1 つの BLOB、codegen、およびロックファイルにスイープします
チャーンは最後にコミットした人の功績として認められ、場合によっては 2 人のエージェントがそうすることもあります
同じ行に着地し、一方が他方を静かに上書きします。そんなことはありません
エージェントが同じタスクに取り組む必要があります。それはまさに共有されたものです
チェ

ckout はデフォルトでそうします。
キルトを使用すると、共有チェックアウトが安全になります。すべてのエージェントが独自のコミットを実行します
行だけで、それ以外は何もありません。ばらばらの作業は最後までばらばらのままです。
儀式のない歴史。 2 人のエージェントが本当に同じコードを必要としている場合、
それはサイレントロスではなく、調整されたハンドオフになります。それはあなたと同じです
エージェントを追加します。
./examples/fleet.sh は、1 つのチェックアウトに対して 7 つのエージェントを実行します。 2つのエンディング:
7 人のエージェントに対して 1 つのコミットをキルトしなかった場合、6 人は「コミットするものが何もありません」となりました。
作業は最初のエージェントの塊に押し込まれました。 a7 黙って
a1 の getUser への変更を上書きしました。 a1の仕事はなくなりました。
キルトを使用すると、6 つのクリーン コミットがエージェントごとに 1 つずつ含まれ、それぞれが正確に独自の行になります。
a1 の要求された関数への a7 の書き込みは、事前に拒否されました。
バイトが変更され、a1 の拒否の意図が示されました。
2 人のエージェントが同じファイルを必要とする場合
ばらばらのファイルをファンアウトするのは簡単なケースです。本当の試練は争いだ。
主張が拒否されたからといって行き詰まりではありません。それには所有者の表明された意図が含まれており、
リース期間が終了したとき:
$ QUILT_ACTOR=builder-flows キルトクレーム deal.js flows.js --intent "取引へのワイヤーフロー"
✗ 拒否された deal.js (builder-friction が保持)
builder-friction は: フリクション パス: 名前変更 + アーカイブ フラグ
更新されない限り、その請求は 2026-07-04T22:12:05Z に失効します。
✓ flows.js を要求しました
したがって、ブロックされたエージェントは待機中に許可されたファイルを構築し、その後再要求します
ホルダーのコミットは自動リリースされ、その変更は着陸したものの上に重ねられます。
1つ。 2 つのクリーンなコミット、どちらもファイル内の変更、何も失われていません。
./examples/contention.sh は、実際のマシン上でシーケンス全体を実行します。
1 つの共有チェックアウト。人間、エージェント、ボットをアクターとしてモデル化し、編集します
作業ツリー。エージェントごとにワークツリーはありません。
行レベルの属性。 commit --mine は、次の場合でも、自分の行のみをコミットします。
彼らは他の俳優とハンクを共有しています。
シンボルレベルのクレーム。予約者

ファイル全体ではなく utils.js#formatPrice なので、
異なる機能を編集するエージェントが競合することはありません。ツリーシッター経由で 10 か国語。
残りのファイル全体のクレーム。
衝突防止。別のエージェントが主張したコードへの書き込みは拒否されます。
バイトが変更される前に、所有者の明示された意図に従って。
プッシュ意識。別のアクターの関数に依存するシンボルを要求します
変更すると、請求時にキルトが警告します。
検出して保存します。あるアクターが別のアクターのコミットされていない行を上書きすると、
Quilt は被害者のバージョンのスナップショットを作成するため、黙って何も失われません。
Quilt が生成するすべてのコミットは通常の Git コミットです。 Git を信頼しますが、決して信頼しません
それを書き換えると、すべての状態が .quilt/ の下でローカルに存在します。アカウントもデーモンもありません。
キルトのセットアップ # キルトをリポジトリにワイヤリングします (MCP サーバー、フック、調整)
キルトドクター # 有線でキャプチャが流れていることを確認してください
それだけです。エージェントには自動的に名前が付けられます: 各クロード コード セッションまたは MCP
接続は独自の ID を取得するため、セットアップなしで並列エージェントが区別されます。
セッション間で安定した ID が必要な場合は、明示的な ID を設定します。
QUILT_ACTOR=auth-agent claude # このエージェントの編集は auth-agent に帰属します
次に、各エージェントは独自の行のみをコミットします。
キルトのステータス # 誰が何を所有しているか
キルトプレビュー --mine # コミットされる正確なパッチ
キルトコミット --mine -m "認証リダイレクトを修正"
キルト フリートは、すべての俳優、その主張、その他すべての全体像を示します。
それには人間が必要です。詳細については docs/reference.md を参照してください。
コマンドリスト。
エージェントごとのワークツリーが通常の答えであり、完全に独立したタスクの場合は、
動作します。ただし、すべてのワークツリーは構築する別の環境です (別のインストール、
別のビルド キャッシュ、別の開発サーバー）と分離は、単に
衝突から合流までの時間。これらのコストはエージェントの数に応じて増加します。全体
共有チェックアウトのポイント i

環境に対して一度お金を払っているのです。
ワークツリーは分離します。彼らは協調していない。エージェントが同じコードを作業するとき、
同時に、通常は、彼らがお互いを見て、お互いを説明することを望んでいます。
彼らは行きます。それがキルトの役割です。この 2 つは相互に排他的ではありません。
独立した長期にわたる作業、同じコード内のエージェントの Quilt を一度に実行します。
キルトのセットアップは、キャプチャ フックと共有 MCP サーバーを接続します。クロード・コードについては、
フックにより、エージェントは Quilt 中に通常どおり組み込みの編集ツールと書き込みツールを使用できるようになります。
各変更の作成者を記録し、他のエージェントが保持するコードへの書き込みをブロックします。
エージェントが従うプロトコルや設定はありません。各セッションには名前が付けられます。
自動的に、または安定した ID のために独自の QUILT_ACTOR を保持します。その他の場合
ランタイムでは、MCP ツールと同じキャプチャと防止を利用できます。
接続にも同じように自動的に名前が付けられます。
Codex、Cursor、Aider、および
エージェントごとのプロセス設定と 1 つのプロセスに複数のエージェントを配置する設定の違い。
docs/orchestrators.md : エージェント群を実行します。
docs/reference.md : 完全なコマンド リスト、帰属方法
動作し、.quilt/ 状態のレイアウト。
bench/ : Quilt がテストされ、実行され、そして実行されるシナリオ ラダー
同じメトリクスでキルトなし。
貢献は大歓迎です。 COTRIBUTING.md を参照してください。
複数の AI エージェントが 1 つの Git チェックアウトを共有する場合、どのエージェントがどの行を書き込んだかを追跡します。
www.npmjs.com/package/@quilt-dev/cli
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
8
v0.4.3: 非同期クレーム、およびキューは公平に動作します
最新の
2026 年 7 月 6 日
+ 7 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026

ギットハブ株式会社
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Tracks which agent wrote which lines when multiple AI agents share one Git checkout. - wkoverfield/quilt

GitHub - wkoverfield/quilt: Tracks which agent wrote which lines when multiple AI agents share one Git checkout. · GitHub
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
wkoverfield
/
quilt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
227 Commits 227 Commits .github .github bench bench design design docs docs examples examples src src test test .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md glama.json glama.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json View all files Repository files navigation
Quilt is a command-line tool that tracks which agent wrote which lines in a
shared Git checkout, so multiple AI coding agents can work in one repo at once
and each commits only its own changes.
It captures every edit at the tool boundary, keeps a per-line record of who wrote
what, and reconstructs each agent's own changes at commit time. Git stays the
source of truth. Quilt never calls an LLM or spawns agents, and its state lives in
a .quilt/ sidecar you can delete without touching your repo.
npm install -g @quilt-dev/cli
quilt setup # wire Quilt into your repo (Claude Code, Cursor, or plain git)
The problem
You want your agents in ONE checkout: one node_modules , one build, one dev
server, one environment to keep working, not a worktree per agent, each with
its own install and its own drift. But on a shared checkout, plain git bites
even when agents work on completely different things: the first git commit -am
sweeps everyone else's uncommitted files into one blob, codegen and lockfile
churn get credited to whoever committed last, and two agents occasionally do
land on the same line, where one silently overwrites the other. None of that
requires agents to be working on the same task. It's just what a shared
checkout does by default.
Quilt makes the shared checkout safe. Every agent commits exactly its own
lines and nothing else: disjoint work stays disjoint all the way into
history, with no ceremony. And when two agents genuinely want the same code,
that becomes a coordinated handoff instead of a silent loss. It holds as you
add agents.
./examples/fleet.sh runs seven agents against one checkout. The two endings:
WITHOUT quilt 1 commit for 7 agents — six got "nothing to commit", their
work swept into the first agent's blob. a7 silently
overwrote a1's change to getUser. a1's work is gone.
WITH quilt 6 clean commits, one per agent, each exactly its own lines.
a7's write into a1's claimed function was denied before any
bytes changed, with a1's stated intent in the denial.
When two agents want the same file
Fanning out on disjoint files is the easy case. The real test is contention.
A denied claim isn't a dead end. It carries the holder's stated intent and
when their lease lapses:
$ QUILT_ACTOR=builder-flows quilt claim deals.js flows.js --intent "wire flows to deals"
✗ denied deals.js (held by builder-friction)
builder-friction is: friction pass: rename + archive flags
their claim lapses 2026-07-04T22:12:05Z unless renewed
✓ claimed flows.js
So the blocked agent builds its granted files while it waits, re-claims after
the holder's commit auto-releases, and layers its change on top of the landed
one. Two clean commits, both changes in the file, nothing lost.
./examples/contention.sh runs the whole sequence on the real machinery.
One shared checkout. Model humans, agents, and bots as actors editing one
working tree, no worktree per agent.
Line-level attribution. commit --mine commits only your lines, even when
they share a hunk with another actor's.
Symbol-level claims. Reserve utils.js#formatPrice , not the whole file, so
agents editing different functions never contend. Ten languages via tree-sitter;
whole-file claims for the rest.
Collision prevention. A write into code another agent has claimed is denied,
with the holder's stated intent, before any bytes change.
Push-awareness. Claim a symbol that depends on a function another actor is
changing, and Quilt warns you at claim time.
Detect and preserve. If one actor overwrites another's uncommitted lines,
Quilt snapshots the victim's version so nothing is silently lost.
Every commit Quilt produces is an ordinary Git commit. It trusts Git and never
rewrites it, and all state lives locally under .quilt/ . No account, no daemon.
quilt setup # wire Quilt into the repo (MCP server, hooks, coordination)
quilt doctor # confirm it's wired and capture is flowing
That's it. Agents are named automatically: each Claude Code session or MCP
connection gets its own id, so parallel agents are told apart with no setup.
Set an explicit id when you want one that is stable across sessions:
QUILT_ACTOR=auth-agent claude # this agent's edits are attributed to auth-agent
Then each agent commits only its own lines:
quilt status # who owns what
quilt preview --mine # exact patch that would be committed
quilt commit --mine -m " fix auth redirect "
quilt fleet shows the whole picture: every actor, their claims, and anything
that needs a human. See docs/reference.md for the full
command list.
A worktree per agent is the usual answer, and for fully independent tasks it
works. But every worktree is another environment to build (another install,
another build cache, another dev server) and isolation just moves the
collision to merge time. Those costs grow with the number of agents; the whole
point of a shared checkout is paying for the environment once.
Worktrees isolate; they don't coordinate. When agents work the same code at the
same time, you usually want them to see each other and account for each other as
they go. That is what Quilt does. The two aren't mutually exclusive: worktrees for
independent, long-running work, Quilt for agents in the same code at once.
quilt setup wires the capture hooks and a shared MCP server. On Claude Code the
hooks let agents use the built-in Edit and Write tools normally while Quilt
records each change's author and blocks a write into code another agent holds,
with no protocol for the agent to follow and no setup: each session is named
automatically, or carries its own QUILT_ACTOR for a stable id. For other
runtimes, the same capture and prevention is available as MCP tools, with each
connection named automatically the same way.
See docs/orchestrators.md for Codex, Cursor, Aider, and
the difference between process-per-agent and many-agents-in-one-process setups.
docs/orchestrators.md : running a fleet of agents.
docs/reference.md : the full command list, how attribution
works, and the .quilt/ state layout.
bench/ : the scenario ladder Quilt is tested against, run with and
without Quilt on the same metrics.
Contributions are welcome. See CONTRIBUTING.md .
Tracks which agent wrote which lines when multiple AI agents share one Git checkout.
www.npmjs.com/package/@quilt-dev/cli
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
8
v0.4.3: async claims, and the queue plays fair
Latest
Jul 6, 2026
+ 7 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
