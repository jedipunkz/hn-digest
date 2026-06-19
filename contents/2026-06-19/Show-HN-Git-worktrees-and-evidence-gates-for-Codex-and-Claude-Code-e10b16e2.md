---
source: "https://github.com/alex-reysa/glueRun-go"
hn_url: "https://news.ycombinator.com/item?id=48601338"
title: "Show HN: Git worktrees and evidence gates for Codex and Claude Code"
article_title: "GitHub - alex-reysa/glueRun-go: Autonomous multi-agent orchestration engine for software repos — L0/L1/L2 agents, leases, gates, audits, git-worktree isolation. Detached dispatch by default. GPL-3.0. · GitHub"
author: "alex-reyss"
captured_at: "2026-06-19T18:42:40Z"
capture_tool: "hn-digest"
hn_id: 48601338
score: 3
comments: 0
posted_at: "2026-06-19T18:07:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Git worktrees and evidence gates for Codex and Claude Code

- HN: [48601338](https://news.ycombinator.com/item?id=48601338)
- Source: [github.com](https://github.com/alex-reysa/glueRun-go)
- Score: 3
- Comments: 0
- Posted: 2026-06-19T18:07:11Z

## Translation

タイトル: HN の表示: Codex および Claude Code の Git ワークツリーと証拠ゲート
記事のタイトル: GitHub - alex-reysa/glueRun-go: ソフトウェア リポジトリ用の自律型マルチエージェント オーケストレーション エンジン — L0/L1/L2 エージェント、リース、ゲート、監査、git-worktree 分離。デフォルトでは分離されたディスパッチ。 GPL-3.0。 · GitHub
説明: ソフトウェア リポジトリ用の自律型マルチエージェント オーケストレーション エンジン — L0/L1/L2 エージェント、リース、ゲート、監査、git-worktree 分離。デフォルトでは分離されたディスパッチ。 GPL-3.0。 - アレックス・レイサ/グルーランゴー

記事本文:
GitHub - alex-reysa/glueRun-go: ソフトウェア リポジトリ用の自律型マルチエージェント オーケストレーション エンジン — L0/L1/L2 エージェント、リース、ゲート、監査、git-worktree 分離。デフォルトでは分離されたディスパッチ。 GPL-3.0。 · GitHub
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
アレックス・レイサ
/
接着剤走る
出版

c
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット cli cli エンジン エンジン gelrun-ext グルーラン ext 移行 マイグレーション プラグイン プラグイン スキーマ スキーマ テンプレート テンプレート テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SCHEMA_VERSION SCHEMA_VERSION SECURITY.md SECURITY.md VERSION VERSION install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
_ ___
__ _| |_ _ ___| _ \_ _ _ _ ___ _ _ ___
/ _ ` | | || / -_) / || | ' \___/ _` / _ \
\__, |_|\_,_\___|_|_\\_,_|_||_| \__、\__/
|__/ |___/
ソフトウェア リポジトリのための自律的なマルチエージェント オーケストレーション。 1 つのエンジンで多数の消費者。
glueRun-go は、自律型 AI コーディング エージェントを駆動する bash + Python オーケストレーション エンジンです
リポジトリに対して並行して実行します。 3 層のスケジューリング モデルを実装します。
(L0 起点ループ → L1 エリア プランナー → L2 ワーカー エージェント) 永続リース、状態パケット、
ゲート/監査パイプライン、および git-worktree 分離。エンジンはマシンごとに 1 回インストールされます
コンシューマー リポジトリごとにピン留めされます。改善は、バージョン ピンをバンプすることで伝播します。
スクリプトを再コピーしています。
階層
役割
L0原点
単一のスケジューラ。インポート→リカバリ→統合→ディスパッチ→スナップショットという調整サイクルを実行します。制御作業中のみ原点ロックを保持します。
L1エリアプランナー
DAG ノード (エリア) ごとに 1 つのプランナー。ノードのコンテキストを読み取り、L2 タスクのバッチを計画し、L0 がインポートする提案としてそれらをステージングします。
L2 ワーカー
タスクごとのブランチ上の分離された git ワークツリーで単一のタスクを実行します。状態パケット (所有ファイル、変更、証拠) を生成します。監査人がパケットをレビューします。決定者が結果をルーティングする

e.
調整サイクル
各 gluun reconcile --actuate は次のように実行されます。
インポート — ステージングされた L1 タスク提案をオリジン ロックの下の DAG にプルします。
回復 — ワーカーが終了またはタイムアウトした古いリースを再利用します。
統合 — 完了したワーカー ブランチを git-op ロック下のターゲット ブランチにマージします。
ディスパッチ — フロンティアタスクを事前にリースし、L2 ワーカーを生成します。
スナップショット — 人間が判読できるプロジェクト状態のスナップショットを書き込みます。
すべての実行中のタスクは、記録を記録するリース ( .gluerun-state/leases/ 内の JSON ファイル) を保持します。
所有権、再試行回数、有効期限。ワーカーは終了すると状態パケットを書き込みます
( state-packet.v0.schema.json ) 所有ファイル、変更されたファイル、コマンド、テスト、および
証拠。監査人はパケットを検証します。リーパーは後のサイクルでの結果を属性とします。
L2 ワーカーが実行されるたびに、ホストは設定されたゲート コマンド (例: npm test ) を実行します。
ゲート結果 (gate-result.v0.schema.json) は監査モデルに供給され、監査モデルは
監査判定 (audit-verdict.v0.schema.json)。ディサイダーは、
(failure-class, retries-left) と回復アクションのペア — 再試行、スコープの修正、エスカレート、または
park — モデルのラウンドトリップにフォールバックする前に、決定論的な高速パス テーブルを使用します。
切り離されたディスパッチはデフォルトでオンになっています。 GLUERUN_DETACHED_DISPATCH=1 (デフォルト) の場合、
各フロンティアタスクを事前にリースし、独自のセッションでワーカーを生成します。
その後、dispatch-wrap.sh を実行し、数秒以内に戻ります。原点ロックはサイクルの間のみ保持されます。
制御作業。リーパー (glurun_reap_dispatches) はすべてのプログラムの先頭で実行されます。
ディスパッチをチェックすることにより、サイクルと属性の完了、失敗、およびクラッシュを適用/実行します
レコード + ワーカー出口ファイル (pid liveness は pid の再利用を無効にします。クラッシュ検出は、
60 分の古いリース ウィンドウから最大 1 サイクルまで）。
これがインポート、インテルを維持するものです

長い間、すりおろし、回復、ステータス、停止に応答します
労働者はバックグラウンドで実行されます。
従来の同期バッチ パスを復元するには、GLUERUN_DETACHED_DISPATCH=0 を設定します。
reconcil は、すべてのワーカーが戻る前に待機します。
Bash >= 4、python3、および git 。
PATH 上で少なくとも 1 つのサポートされているランナー CLI ( clude 、 codex 、またはその他)
構成されたランナー)。
macOS ユーザーには、brew install bash と、これを解決する PATH エントリが必要になる場合があります。
bash を /bin/bash より前の Homebrew バージョンに変更します。
# エンジンをクローンして ~/.gluerun にインストールします
git clone https://github.com/alex-reysa/glueRun-go /path/to/glueRun-go
cd /path/to/glueRun-go
bash インストール.sh
# -> ~/.gluerun/versions/<ver>/ ~/.gluerun/current ~/.gluerun/bin/gluerun
エクスポート PATH= " $HOME /.gluerun/bin: $PATH "
各コンシューマ リポジトリでは次のようになります。
グルーラン初期化 # scaffold グルーラン.config.json、docs/orchestration/、.gluerun-version
Glurun Doctor # deps、エンジン解像度、リポジトリ構成を確認します
各リポジトリは、エンジンのバージョンを .gluerun-version に固定します (gluerun.config.json をオーバーライドします)
エンジンのバージョン)。 Gluerun ランチャーは、 ~/.gluerun/versions/<ver> からそのバージョンを解決します。
GLUERUN_ROOT を現在のリポジトリにバインドし、その構成をロードし、エンジンを実行します。走る
グルーランは <ver> を更新して再固定します。
# 調整/実行サイクルを 1 回実行します (インポート → リカバリ → 統合 → ディスパッチ → スナップショット)
グルーラン調整 --actuate
# L1 → L2 → 監査を通じて単一のタスクを実行する
グルールンドライブ TASK-0001
# 自動運転自律ループ (実時間の予算: GLUERUN_MAX_HOURS)
グルーランオート
# デタッチされたワーカーがすべて終了するまでブロックします (CI またはクリーン シャットダウンで役立ちます)
グルーラン調整 --drain
構成
リポジトリごとのバリエーションはすべてコンシューマ リポジトリ内に存在し、エンジン ファイル内には存在しません。
Gluerun.config.json — 宣言型: targetBranch 、gateCommand 、runner 、
area{} 、 areaPrefix 、 prewarm 、 modules[] 、identity{} 、 env{}

、
ProvisionFiles[] 、 envAllowlist[] 。
Gluerun.config.sh — オプションのシェル エクストラ (計算値、関数)。
.gluerun-state/config.local.sh — gitignored オペレーターのオーバーライドとシークレット。
スターター設定では、gateCommand を意図的に false に設定するため、新たに
スキャフォールドされたリポジトリは、それを証明するコマンドに置き換えるまで閉じられません。
リポジトリは健全です。
ProvisionFiles エントリは、リポジトリローカルの gitignored ファイルを各ワーカーにコピーします。
git worktree add の後の worktree : { "source": ".env.local", "target": ".env.local", "required": true } 。ソースとターゲットは両方とも無視する必要があります
またはプロビジョニングが失敗して閉じられます。 envAllowlist は正確な環境名またはプレフィックスを受け入れます
* で終わるパターン;許可された値が書き込まれる
worktree/.gluerun-state/worktree-env.sh であり、prewarm/gate フェーズ用に提供されます。
環境ノブ
デフォルト
効果
GLUERUN_MAX_CONCURRENT
5
同時に実行される最大 L2 ワーカー。
GLUERUN_MAX_DISPATCH
5
調整サイクルごとにディスパッチされるタスクの最大数。
GLUERUN_DETACHED_DISPATCH
1
デフォルトはオンです。 Reconcile は独自のセッションでワーカーを生成し、数秒で戻ります。リーパーは後のサイクルでの結果を属性とします。従来の同期バッチ待機には 0 を設定します。
GLUERUN_AUTO_INTEGRATE
1
完了したワーカー ブランチを直接リコンサイル --actuate 、gleun auto 、launchd、およびコンソール駆動のサイクルで自動的に統合 (マージ) します。
GLUERUN_PUSH
0 ダイレクト / 1 オート
統合されたブランチをリモートにプッシュします。ダイレクト エンジン コマンドのデフォルトはローカルのみです。オーバーライドされない限り、gluun auto /launchd セット 1。
GLUERUN_MAX_HOURS
12
自律性ループの実時間の予算 (グルーラン自動)。
GLUERUN_MAX_RETRIES
3
タスクごとのワーカーは、ディサイダーがエスカレートする前に再試行します。
GLUERUN_STALE_MINUTES
60
ライブディスパッチ PID のないタスクがリーパーによって再利用されるまでのリース期間 (分)。
GLUERUN_TARGET_BRANCH
(必須)
統合対象b

消費者リポジトリの牧場。
GLUERUN_SESSION_AFFINITY
1
すべての失効ゲートを通過したら、ロールの前のランタイム セッションを再利用します。 0 は常に新鮮に実行されます。
GLUERUN_FIX_PROMPT_STRUCTURED
1
再試行時に構造化された修正プロンプト (信頼できる調査結果)。 0 = 従来の fix_hints テール。
GLUERUN_DECIDER_FAST
1
ホストポリシーテーブルによって明確な障害クラスを解決します。 0 は、すべての失敗をモデル ディサイダーを通じてルーティングします。
GLUERUN_WORKER_INFRA_MAX
1
追加のワーカーは、 worker-infra が表示される前に、インフラ障害時に再実行されます。
GLUERUN_AUDIT_INFRA_MAX
2
追加の監査人は、 Audit-infra が表示される前にインフラ障害で再実行されます。
GLUERUN_CONTEXT_SECTION_MAX_CHARS
4000
プロンプトに追加される継続コンテンツのセクションごとの上限。
GLUERUN_PREFLIGHT_REQUIRE_ACCEPTANCE
1
プリフライトでは、タスクに対して空ではない acceptCriteria が必要です。
コンテキストの連続性
再試行の間に、glueRun-go は権限のある状態を転送します。
ログ末尾から再導出します。
コンテキストカプセル — ハッシュスタンプされたimplementer-capsule.jsonと
試行ごとの reviewer-capsule.json。
所見台帳 — 各監査評決から更新/挿入された Finding-status.json、
安定した検出結果 ID は、再試行にわたってオープン/解決された状態で追跡されます。
構造化された修正プロンプト - 作業者は再試行時に信頼できる未解決の結果を受け取ります
(従来のバイトテールに戻すには、GLUERUN_FIX_PROMPT_STRUCTURED=0 を設定します)。
デルタプロンプトの再監査 — 監査人は以前の調査結果を受け取り + 差異を修正 +
ID ごとの検証ターゲット。
試行のアーカイブ — 各試行の成果物は次の場所にコピーされます (移動されません)。
実行/<id>/attempts/<n>/ と returns/index.json 。
オプションのロールキー付きランタイムセッション再開 ( codex execresume 、 claude -r )
10 個の失効ゲート、デフォルトは ON ( GLUERUN_SESSION_AFFINITY=1 )。ゲートの故障または
再開を拒否したランナーは、同じアテ内で新たなランに静かに降格します。

mpt。
不変: セッションの再開は、タスクを変更しないトークンコストの最適化です。
結果。
汎用エンジン/はプロジェクト固有のシンボルを参照しません — によって強制されます
testing/test-engine-clean.sh (抽象化ゲート テスト)。プロジェクトごとのロジックはすべて次の場所に存在します。
オプトインモジュール:
グルーラン-ext/
storage-proof.sh # 例: 耐久性プルーフ体制
promote-gate.sh # 例: ゲート プロモーター
モジュールは、gluun.config.json → modules[] にリストされています。それらがリストされていないリポジトリ
決してロードしません。 GLUERUN_MODULES 環境変数は、ランタイム リスト (JSON によって設定されます) です。
設定ローダー）。
2 つのバージョンは独立して動作します。
エンジン ピン — .gluerun-version は正規のリポジトリごとのピンです (オーバーライドされます)
gluun.config.json エンジンバージョン ;彼らが同意しない場合は、.gluerun-version が勝ちます。
グルールン医師は警告する）。 gluun update <ver> で書き換えます。
スキーマ — SCHEMA_VERSION (リポジトリ ルート) はデータ コントラクト バージョン (現在は v1) を保持します。
リポジトリは、glurun.config.json でスキャフォールディングされたスキーマを記録します →
スキーマバージョン 。グルーランドクターはスキーマの不一致で失敗します。グルーラン移行実行
出荷された migrations/<from>-to-<to>.sh チェーンと schemaVersion の書き換え。
すべてのランタイム JSON スキーマ識別子は、名前空間 gluun.orchestration.*.v0 に従います。
bash テスト/run.sh # 23 回帰テスト
テストスイートはライブ統計を使用しません

[切り捨てられた]

## Original Extract

Autonomous multi-agent orchestration engine for software repos — L0/L1/L2 agents, leases, gates, audits, git-worktree isolation. Detached dispatch by default. GPL-3.0. - alex-reysa/glueRun-go

GitHub - alex-reysa/glueRun-go: Autonomous multi-agent orchestration engine for software repos — L0/L1/L2 agents, leases, gates, audits, git-worktree isolation. Detached dispatch by default. GPL-3.0. · GitHub
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
alex-reysa
/
glueRun-go
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits cli cli engine engine gluerun-ext gluerun-ext migrations migrations plugin plugin schemas schemas templates templates tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SCHEMA_VERSION SCHEMA_VERSION SECURITY.md SECURITY.md VERSION VERSION install.sh install.sh View all files Repository files navigation
_ ___
__ _| |_ _ ___| _ \_ _ _ _ ___ __ _ ___
/ _` | | || / -_) / || | ' \___/ _` / _ \
\__, |_|\_,_\___|_|_\\_,_|_||_| \__, \___/
|___/ |___/
Autonomous multi-agent orchestration for software repos. One engine, many consumers.
glueRun-go is a bash + Python orchestration engine that drives autonomous AI coding agents
in parallel against a repository. It implements a three-tier scheduling model
(L0 origin loop → L1 area planners → L2 worker agents) with durable leases, state packets,
gate/audit pipelines, and git-worktree isolation. The engine is installed once per machine
and pinned per consumer repo — improvements propagate by bumping a version pin, not by
re-copying scripts.
Tier
Role
L0 origin
The single scheduler. Runs the reconcile cycle: import → recover → integrate → dispatch → snapshot. Holds the origin lock only during control work.
L1 area planners
One planner per DAG node (area). Reads the node's context, plans a batch of L2 tasks, and stages them as proposals for L0 to import.
L2 workers
Execute a single task in an isolated git worktree on a per-task branch. Produce a state packet (owned files, changes, evidence). An auditor reviews the packet; the decider routes the outcome.
Reconcile cycle
Each gluerun reconcile --actuate runs:
Import — pull staged L1 task proposals into the DAG under the origin lock.
Recover — reclaim stale leases whose workers have exited or timed out.
Integrate — merge completed worker branches into the target branch under the git-op lock.
Dispatch — pre-lease frontier tasks and spawn L2 workers.
Snapshot — write a human-readable project state snapshot.
Every in-flight task holds a lease (a JSON file in .gluerun-state/leases/ ) that records
ownership, retry count, and expiry. When a worker finishes it writes a state packet
( state-packet.v0.schema.json ) enumerating owned files, changed files, commands, tests, and
evidence. The auditor validates the packet; the reaper attributes outcomes on later cycles.
After each L2 worker run the host executes the configured gate command (e.g. npm test ).
A gate result ( gate-result.v0.schema.json ) feeds the auditor model, which returns an
audit verdict ( audit-verdict.v0.schema.json ). The decider maps the
(failure-class, retries-left) pair to a recovery action — retry, amend-scope, escalate, or
park — using a deterministic fast-path table before falling back to a model round-trip.
Detached dispatch is ON by default. When GLUERUN_DETACHED_DISPATCH=1 (the default),
reconcile pre-leases each frontier task and spawns the worker in its own session via
dispatch-wrap.sh , then returns within seconds. The origin lock is held only for the cycle's
control work. A reaper ( gluerun_reap_dispatches ) runs at the top of every
apply/actuate cycle and attributes completions, failures, and crashes by checking dispatch
records + worker exit files (pid liveness defeats pid reuse; crash detection drops from the
60-min stale-lease window to ~one cycle).
This is what keeps import, integrate, recover, STATUS, and STOP responsive while long
workers run in the background.
Set GLUERUN_DETACHED_DISPATCH=0 to restore the legacy synchronous batch path, where
reconcile waits for every worker before returning.
Bash >= 4, python3 , and git .
At least one supported runner CLI on PATH ( claude , codex , or another
configured runner).
macOS users may need brew install bash and a PATH entry that resolves
bash to the Homebrew version before /bin/bash .
# Clone and install the engine to ~/.gluerun
git clone https://github.com/alex-reysa/glueRun-go /path/to/glueRun-go
cd /path/to/glueRun-go
bash install.sh
# -> ~/.gluerun/versions/<ver>/ ~/.gluerun/current ~/.gluerun/bin/gluerun
export PATH= " $HOME /.gluerun/bin: $PATH "
In each consumer repo:
gluerun init # scaffold gluerun.config.json, docs/orchestration/, .gluerun-version
gluerun doctor # check deps, engine resolution, repo config
Each repo pins its engine version in .gluerun-version (overrides gluerun.config.json
engineVersion ). The gluerun launcher resolves that version from ~/.gluerun/versions/<ver> ,
binds GLUERUN_ROOT to the current repo, loads its config, and execs the engine. Run
gluerun update <ver> to repin.
# Run one reconcile/actuate cycle (import → recover → integrate → dispatch → snapshot)
gluerun reconcile --actuate
# Drive a single task through L1 → L2 → audit
gluerun drive TASK-0001
# Self-driving autonomy loop (wall-clock budget: GLUERUN_MAX_HOURS)
gluerun auto
# Block until all detached workers finish (useful in CI or clean shutdown)
gluerun reconcile --drain
Configuration
All per-repo variation lives in the consumer repo, never in engine files:
gluerun.config.json — declarative: targetBranch , gateCommand , runner ,
areas{} , areaPrefix , prewarm , modules[] , identity{} , env{} ,
provisionFiles[] , envAllowlist[] .
gluerun.config.sh — optional shell extras (computed values, functions).
.gluerun-state/config.local.sh — gitignored operator overrides and secrets.
The starter config deliberately sets gateCommand to false so a newly
scaffolded repo fails closed until you replace it with the command that proves
the repo is healthy.
provisionFiles entries copy repo-local, gitignored files into each worker
worktree after git worktree add : { "source": ".env.local", "target": ".env.local", "required": true } . The source and target must both be ignored
or provisioning fails closed. envAllowlist accepts exact env names or prefix
patterns ending in * ; allowed values are written to
worktree/.gluerun-state/worktree-env.sh and sourced for prewarm/gate phases.
Env knob
Default
Effect
GLUERUN_MAX_CONCURRENT
5
Maximum L2 workers running concurrently.
GLUERUN_MAX_DISPATCH
5
Maximum tasks dispatched per reconcile cycle.
GLUERUN_DETACHED_DISPATCH
1
Default ON. Reconcile spawns workers in their own session and returns in seconds; the reaper attributes outcomes on later cycles. Set 0 for the legacy synchronous batch wait.
GLUERUN_AUTO_INTEGRATE
1
Automatically integrate (merge) completed worker branches in direct reconcile --actuate , gluerun auto , launchd, and console-driven cycles.
GLUERUN_PUSH
0 direct / 1 auto
Push integrated branches to the remote. Direct engine commands default local-only; gluerun auto /launchd set 1 unless overridden.
GLUERUN_MAX_HOURS
12
Wall-clock budget for the autonomy loop ( gluerun auto ).
GLUERUN_MAX_RETRIES
3
Per-task worker retries before the decider escalates.
GLUERUN_STALE_MINUTES
60
Lease age (minutes) before a task without a live dispatch pid is reclaimed by the reaper.
GLUERUN_TARGET_BRANCH
(required)
Integration target branch in the consumer repo.
GLUERUN_SESSION_AFFINITY
1
Reuse a role's prior runtime session when all staleness gates pass; 0 always runs fresh.
GLUERUN_FIX_PROMPT_STRUCTURED
1
Structured fix prompt on retries (authoritative findings); 0 = legacy fix_hints tail.
GLUERUN_DECIDER_FAST
1
Resolve clear-cut failure classes by host policy table; 0 routes every failure through the model decider.
GLUERUN_WORKER_INFRA_MAX
1
Extra worker re-runs on an infra failure before surfacing worker-infra .
GLUERUN_AUDIT_INFRA_MAX
2
Extra auditor re-runs on an infra failure before surfacing audit-infra .
GLUERUN_CONTEXT_SECTION_MAX_CHARS
4000
Per-section cap on continuity content appended to prompts.
GLUERUN_PREFLIGHT_REQUIRE_ACCEPTANCE
1
Preflight requires non-empty acceptanceCriteria on a task.
Context continuity
Between retry attempts glueRun-go carries authoritative state forward rather than
re-deriving it from a log tail:
Context capsules — hash-stamped implementer-capsule.json and
reviewer-capsule.json per attempt.
Findings ledger — findings-status.json upserted from each audit verdict, with
stable finding ids tracked open/resolved across retries.
Structured fix prompts — the worker receives authoritative open findings on retry
(set GLUERUN_FIX_PROMPT_STRUCTURED=0 to revert to the legacy byte-tail).
Re-audit delta prompts — the auditor receives prior findings + fix diff +
per-id verification targets.
Attempt archive — each attempt's artifacts are copied (never moved) under
runs/<id>/attempts/<n>/ with an attempts/index.json .
Optional role-keyed runtime session resume ( codex exec resume , claude -r ) behind
10 staleness gates, defaulting ON ( GLUERUN_SESSION_AFFINITY=1 ). Any gate failure or
runner that refuses the resume degrades silently to a fresh run within the same attempt.
Invariant: session resume is a token-cost optimization that never changes a task
outcome.
The generic engine/ references zero project-specific symbols — enforced by
tests/test-engine-clean.sh (the abstraction gate test). All per-project logic lives in
opt-in modules:
gluerun-ext/
storage-proof.sh # example: durable-proof regime
promote-gate.sh # example: gate promoter
Modules are listed in gluerun.config.json → modules[] . A repo that doesn't list them
never loads them. The GLUERUN_MODULES env var is the runtime list (set by the JSON
config loader).
Two versions move independently:
Engine pin — .gluerun-version is the canonical per-repo pin (overrides
gluerun.config.json engineVersion ; if they disagree .gluerun-version wins and
gluerun doctor warns). gluerun update <ver> rewrites it.
Schema — SCHEMA_VERSION (repo root) holds the data-contract version ( v1 today).
A repo records the schema it was scaffolded against in gluerun.config.json →
schemaVersion . gluerun doctor fails on a schema mismatch; gluerun migrate runs
the shipped migrations/<from>-to-<to>.sh chain and rewrites schemaVersion .
All runtime JSON schema identifiers follow the namespace gluerun.orchestration.*.v0 .
bash tests/run.sh # 23 regression tests
The test suite uses no live stat

[truncated]
