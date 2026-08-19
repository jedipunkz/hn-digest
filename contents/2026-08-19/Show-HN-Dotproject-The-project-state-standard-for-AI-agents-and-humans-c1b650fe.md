---
source: "https://github.com/jasonnam/dotproject"
hn_url: "https://news.ycombinator.com/item?id=49361544"
title: "Show HN: Dotproject – The project state standard for AI agents and humans"
article_title: "GitHub - jasonnam/dotproject: The Git-native, collision-free project state standard for AI agents and humans. · GitHub"
image: "https://opengraph.githubassets.com/967ba8a239579fd1edbe37222c997589e863eec50c12ce2e04e08afc0e6123c5/jasonnam/dotproject"
author: "swift3"
captured_at: "2026-08-19T14:24:18Z"
capture_tool: "hn-digest"
hn_id: 49361544
score: 1
comments: 0
posted_at: "2026-08-19T13:46:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Dotproject – The project state standard for AI agents and humans

- HN: [49361544](https://news.ycombinator.com/item?id=49361544)
- Source: [github.com](https://github.com/jasonnam/dotproject)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T13:46:06Z

## Translation

タイトル: Show HN: Dotproject – AI エージェントと人間のためのプロジェクト状態標準
記事のタイトル: GitHub - jasonnam/dotproject: AI エージェントと人間のための、Git ネイティブで衝突のないプロジェクト状態標準。 · GitHub
説明: AI エージェントと人間のための、Git ネイティブで衝突のないプロジェクト状態標準。 - ジャソンナム/ドットプロジェクト
HN テキスト: 個々の機能についてエージェントを詳細に管理するにはすでに時間がかかりすぎています。あらゆる小さなタスクを確認して操作しようとすると、すぐに主なボトルネックになってしまいます。マイルストーンと目標レベルでの管理は、少なくとも今のところは人間が現実的に続けられるもののように感じられます...私はこのアイデアに基づいて dotproject スキルを構築しました。実行速度が人間の監視を上回り続ける中、他の人がエージェントのワークフローをどのように構築しているか興味があります。理論的には、目標とマイルストーンが定義され、十分な AI エージェントが接続されれば、それらに向けた進捗状況をリアルタイムで追跡および推進できるはずです。

記事本文:
GitHub - jasonnam/dotproject: AI エージェントと人間のための、Git ネイティブで衝突のないプロジェクト状態標準。 · GitHub
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
ジャソンナム
/
ドットプロジェクト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
2 コミット 2 コミット フォルダーとファイル
.agents/ プラグイン .agents/ プラグイン .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor-plugin .cursor-plugin アセット アセット docs/ superpowers/ spe

cs docs/ superpowers/ スペックの例/ サンプル リポジトリの例/ サンプル リポジトリ スキーマ スキーマ スキル/ dotproject スキル/ dotproject ビューア/ Web ビューア/ web .gitattributes .gitattributes AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md SPEC.md SPEC.md Viewすべてのファイル リポジトリ ファイルのナビゲーション
プレーンな JSON ファイルでの競合のない Git ネイティブのプロジェクト追跡。
インフラストラクチャゼロ。自律エージェントを同時に実行できるように構築されています。
従業員が多いのにコーディネーターがいない。コームは一度に 1 つのセルで構築されます。
そして、2 匹のミツバチが同じ細胞を必要とすることはありません。
プロジェクトの状態は、コードの隣の .project/ にあります。エージェントオン
並列ブランチはタスクを要求し、進行状況を記録し、作業を検証します。
2 人のライターがブランチに触れないため、2 つのブランチのマージは成功します。
同じバイト。
git レイヤー: 定義ファイルは一度だけ書き込み可能であり、すべての変更が行われます。
その後は、ULID + アクターによって名前が付けられた新しいファイルです。二人の作家は決して
同じパスが生成されるため、 git merge が .project/ で競合することはありません。
セマンティック層では、競合は隠蔽されるのではなく、報告されます。
すべてのイベントは、その作成者が見たもの (observed_through) を記録します。いつ
2 人のライターが同じフィールドを同時に変更しても、実体化は続行されません
1 つの決定的な答えが得られ、相違点、両方のイベントに名前が付けられます。
そして二人の俳優。
.プロジェクト/
project.json # マニフェスト
Goals/G-off42.json # なぜ - 戦略的目標
Milestones/M-alpha7.json # いつ — エポックをリリースするか
task/T-syn9b.json # 何を — 実行 + 検証
events/01JAR…__agent-a__T-syn9b__claimed.json # 追加のみの履歴
snapshots/M-alpha7.01JAR….snapshot.json # オプションの圧縮
project.json と、目標、マイルストーン、および定義ファイルを作成します。
いくつかのタスク。定義ファイルは再度編集されることはありません。
何かを変更するには、イベント ファイルを 1 つ追加します。
events/<ULID>__<あなた>__<エンティティ>__<a

変更マップを含む ction>.json。
現在の状態 = すべてのイベントを ULID 順に定義の上に重ねます。
決定論的。どの読者も同じ答えを得るでしょう。参照
SPEC.md §6 。
インストールはありません。デーモンはありません。依存性はありません。実際に動作した例は次のとおりです。
Examples/sample-repo/ とその正確な期待値
状態
期待される状態.md 。
dotproject スキルはあらゆるコーディングを教えます
エージェントの実行ループ:
Discover (状態を具体化) → Claim (イベントの追加) → 作業
(タスクの context_files 内のみ) → 検証 (実行)
verify.command ) → 出力 (結果を記録)。
ファイルの読み取り、ファイルの書き込み、ファイルのリストのみを想定しており、シェルやツールは使用しません。
カップリング。プラグインとしてインストール (Claude Code、Codex、Cursor、Antigravity)
または、このリポジトリで AGENTS.md を読み取るハーネスを指定します。
同じフォーマットで計画されている後続スキル: ブートストラップ (仕様を変更する)
または最初の .project/ ツリーへの会話)、レポート (要約
人間の状態)、マージ レビュー (マージ後の相違レポート)、
マイルストーンを閉じます (スナップショットを書き込みます)。
マーケットプレイスを追加し、プラグインをインストールします。
# マーケットプレイスを追加
クロード プラグイン マーケットプレイスに jasonnam/dotproject を追加
# プラグインをインストールする
クロードプラグインのインストール dotproject
または、開発中にローカル クローンから直接ロードします。
クロード --plugin-dir /path/to/dotproject
カーソル
プラグイン設定から dotproject をインストールするか、skills/dotproject/ をワークスペースのスキル ディレクトリにコピーします。
コーデックス / 反重力 / その他のハーネス
エージェント ハーネスを skill/dotproject/SKILL.md に向けます。
または、skills/dotproject フォルダをハーネスのアクティブなスキル ディレクトリにコピーします。
独自のリポジトリで dotproject を使用するには:
.project/ ディレクトリ ツリーを作成します。
mkdir -p .project/{目標、マイルストーン、タスク、イベント、スナップショット}
.project/project.json を作成します。
{
"spec_version" : " 0.1.0 " 、
"id" : " P-myproj " ,
"名前" : "私のプロジェクト" ,
「ステータス語彙」

り" : {
"ステータス" : [ "保留中" 、 "申請中" 、 "進行中" 、 "ブロック済み" 、 "完了" 、 "キャンセル" ]、
"ターミナル" : [ "完了" 、 "キャンセル" ]
}、
「拡張子」: [],
「メタ」: {}
}
プロジェクトの AGENTS.md または CLAUDE.md にポインターを追加します。
** ` .project/ ` ツリーを操作しますか? ** `skills/dotproject/SKILL.md`を読んでフォローしてください。
viewer/web/index.html — 任意の場所で開きます
ブラウザで .project/ フォルダーを指定します。インストールもサーバーも不要
構築する。フォルダーを読み取り、ボード、depends_on グラフをレンダリングします。
イベントの軌跡と分岐レポート。ディスクには何も書き込まれません。
これは仕様ではなく参照実装ですが、仕様でもあります。
フォーマットの 2 番目の独立したマテリアライザーを再現し、
正確には、example-repo/EXPECTED-STATE.md です。もしそれと
SPEC.md はこれに同意しません。閲覧者が間違っています。
人間一人での直線的な仕事ですか?マークダウンチェックリストを使用します。ドットプロジェクトの
読み取りパス (定義 + イベント フォールド) は TODO.md よりもコストが高く、
同時ライターがいる場合にのみ効果を発揮します。
厳しい強制執行が必要ですか?検証ゲートは信頼と監査のゲートです
モデル — この形式はコマンドと終了コードを記録しますが、システムを停止することはできません。
嘘をついた作家。強制は、同じファイルを使用する CI チェックです。
チーム間での即時同期が必要ですか? git がマージされると状態が伝播します。
1 秒未満の共有状態が必要な場合は、ファイルではなくサービスが必要です。
dotproject は、次の 3 つのことが同時に当てはまる場合にオーバーヘッドを獲得します。
同時ライター、リポジトリ内に存在する必要がある状態、および監査証跡
git log を実行できます。
完全な仕様 — エンティティ モデル、具体化アルゴリズム、クレーム
プロトコル、相違レポート、スナップショット、拡張ポイント - にあります
SPEC.md 。 JSON スキーマは schemas/ にあります。
AI エージェントと人間のための、Git ネイティブで衝突のないプロジェクト状態標準。
Readme MIT ライセンス アクティビティ

y スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The Git-native, collision-free project state standard for AI agents and humans. - jasonnam/dotproject

We are already getting too slow to micromanage agents on individual features. Trying to review and steer every tiny task is quickly becoming the main bottleneck. Managing at the milestone and goal level still feels like something humans can realistically keep up with—at least for now... I built the dotproject skill around this idea. Curious how others here are structuring agent workflows as execution speeds keep outpacing human oversight. Theoretically, once goals and milestones are defined and enough AI agents are attached, we should be able to track and drive progress toward them in real time.

GitHub - jasonnam/dotproject: The Git-native, collision-free project state standard for AI agents and humans. · GitHub
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
jasonnam
/
dotproject
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
.agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor-plugin .cursor-plugin assets assets docs/ superpowers/ specs docs/ superpowers/ specs examples/ sample-repo examples/ sample-repo schemas schemas skills/ dotproject skills/ dotproject viewers/ web viewers/ web .gitattributes .gitattributes AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md SPEC.md SPEC.md View all files Repository files navigation
Conflict-free, Git-native project tracking in plain JSON files.
Zero infrastructure. Built for concurrent autonomous agents.
Many workers, no coordinator. The comb is built one cell at a time,
and no two bees need the same cell.
Your project's state lives in .project/ next to your code. Agents on
parallel branches claim tasks, record progress, and verify work — and any
merge of any two branches succeeds, because no two writers ever touch the
same bytes.
At the git layer: definition files are write-once, and every change
after that is a new file named by ULID + actor. Two writers never
produce the same path, so git merge never conflicts on .project/ .
At the semantic layer: conflicts aren't hidden, they're reported .
Every event records what its author had seen ( observed_through ); when
two writers change the same field concurrently, materialization still
yields one deterministic answer — and names the divergence, both events,
and both actors.
.project/
project.json # manifest
goals/G-off42.json # why — strategic targets
milestones/M-alpha7.json # when — release epochs
tasks/T-syn9b.json # what — execution + verification
events/01JAR…__agent-a__T-syn9b__claimed.json # append-only history
snapshots/M-alpha7.01JAR….snapshot.json # optional compaction
Create project.json and definition files for a goal, a milestone, and
some tasks. Definition files are never edited again.
To change anything, append one event file:
events/<ULID>__<you>__<entity>__<action>.json with a changes map.
Current state = fold all events over the definitions, in ULID order.
Deterministic; any reader gets the same answer. See
SPEC.md §6 .
No install. No daemon. No dependency. A worked example lives in
examples/sample-repo/ with its exact expected
state in
EXPECTED-STATE.md .
The dotproject skill teaches any coding
agent the execution loop:
Discover (materialize state) → Claim (append event) → Work
(only within the task's context_files ) → Verify (run
verification.command ) → Emit (record the result).
It assumes only read-file, write-file, and list-files — no shell, no tool
coupling. Install as a plugin (Claude Code, Codex, Cursor, Antigravity)
or point any harness that reads AGENTS.md at this repo.
Planned follow-on skills over the same format: bootstrap (turn a spec
or conversation into an initial .project/ tree), report (summarize
state for a human), merge review (post-merge divergence report),
close milestone (write the snapshot).
Add the marketplace and install the plugin:
# Add marketplace
claude plugin marketplace add jasonnam/dotproject
# Install plugin
claude plugin install dotproject
Or load directly from a local clone during development:
claude --plugin-dir /path/to/dotproject
Cursor
Install dotproject via plugin settings, or copy skills/dotproject/ into your workspace skills directory.
Codex / Antigravity / Other Harnesses
Point your agent harness at skills/dotproject/SKILL.md .
Or copy the skills/dotproject folder into your harness's active skills directory.
To use dotproject in your own repository:
Create the .project/ directory tree:
mkdir -p .project/{goals,milestones,tasks,events,snapshots}
Create .project/project.json :
{
"spec_version" : " 0.1.0 " ,
"id" : " P-myproj " ,
"name" : " My Project " ,
"status_vocabulary" : {
"statuses" : [ " pending " , " claimed " , " in_progress " , " blocked " , " completed " , " cancelled " ],
"terminal" : [ " completed " , " cancelled " ]
},
"extensions" : [],
"meta" : {}
}
Add a pointer in your project's AGENTS.md or CLAUDE.md :
** Working with a ` .project/ ` tree? ** Read and follow ` skills/dotproject/SKILL.md ` .
viewers/web/index.html — open it in any
browser and point it at a .project/ folder. No install, no server, no
build; it reads the folder and renders the board, the depends_on graph,
the event trail, and the divergence report. Nothing is written to disk.
It is a reference implementation, not the spec — but it is also the
format's second independent materializer, and it reproduces
examples/sample-repo/EXPECTED-STATE.md exactly. If it and
SPEC.md ever disagree, the viewer is wrong.
Solo human, linear work? Use a markdown checklist. dotproject's
read path (definitions + event folds) costs more than TODO.md and
pays off only when you have concurrent writers.
Need hard enforcement? The verification gate is a trust-and-audit
model — the format records commands and exit codes but cannot stop a
writer from lying. Enforcement is a CI check consuming the same files.
Need instant cross-team sync? State propagates when git merges.
If you need sub-second shared state, you need a service, not files.
dotproject earns its overhead when three things are true at once:
concurrent writers, state that must live in the repo, and an audit trail
you can git log .
The full specification — entity model, materialization algorithm, claim
protocol, divergence reporting, snapshots, extension points — is in
SPEC.md . JSON Schemas are in schemas/ .
The Git-native, collision-free project state standard for AI agents and humans.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
