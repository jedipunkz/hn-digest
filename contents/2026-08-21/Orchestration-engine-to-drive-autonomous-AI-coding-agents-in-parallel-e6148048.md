---
source: "https://github.com/alex-reysa/singular-lite"
hn_url: "https://news.ycombinator.com/item?id=49389218"
title: "Orchestration engine to drive autonomous AI coding agents in parallel"
article_title: "GitHub - alex-reysa/singular-lite: Autonomous multi-agent orchestration engine for software repos — L0/L1/L2 agents, leases, gates, audits, git-worktree isolation. Detached dispatch by default. GPL-3.0. · GitHub"
image: "https://opengraph.githubassets.com/bf6bc21d8f480573fd65085cd1593c02cc662227b20ad69d37c147176b243564/alex-reysa/singular-lite"
author: "alexreysa"
captured_at: "2026-08-21T15:24:20Z"
capture_tool: "hn-digest"
hn_id: 49389218
score: 2
comments: 0
posted_at: "2026-08-21T15:03:40Z"
tags:
  - hacker-news
  - translated
---

# Orchestration engine to drive autonomous AI coding agents in parallel

- HN: [49389218](https://news.ycombinator.com/item?id=49389218)
- Source: [github.com](https://github.com/alex-reysa/singular-lite)
- Score: 2
- Comments: 0
- Posted: 2026-08-21T15:03:40Z

## Translation

タイトル: 自律型 AI コーディング エージェントを並行して駆動するオーケストレーション エンジン
記事のタイトル: GitHub - alex-reysa/singular-lite: ソフトウェア リポジトリ用の自律型マルチエージェント オーケストレーション エンジン — L0/L1/L2 エージェント、リース、ゲート、監査、git-worktree 分離。デフォルトでは分離されたディスパッチ。 GPL-3.0。 · GitHub
説明: ソフトウェア リポジトリ用の自律型マルチエージェント オーケストレーション エンジン — L0/L1/L2 エージェント、リース、ゲート、監査、git-worktree 分離。デフォルトでは分離されたディスパッチ。 GPL-3.0。 - アレックス・レイサ/singular-lite

記事本文:
GitHub - alex-reysa/singular-lite: ソフトウェア リポジトリ用の自律型マルチエージェント オーケストレーション エンジン — L0/L1/L2 エージェント、リース、ゲート、監査、git-worktree 分離。デフォルトでは分離されたディスパッチ。 GPL-3.0。 · GitHub
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
アレックス・レイサ
/
単数ライト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
656 コミット 656 コミット フォルダーとファイル
claudedocs claudedocs cli cli docs ドキュメントエンジン エンジン

移行 移行 プラグイン プラグイン スキーマ スキーマ singular-ext singular-ext テンプレート テンプレート テスト テスト ツール ツール .gitignore .gitignore .singular-version .singular-version CHANGELOG.md CHANGELOG.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SCHEMA_VERSION SCHEMA_VERSION SECURITY.md SECURITY.md VERSION VERSION install.sh install.sh singular.config.json singular.config.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
TcGUeulhd4hM0DvB5ytcng-under10.mp4
単数形
ソフトウェア リポジトリのための自律的なマルチエージェント オーケストレーション。 1 つのエンジンで多数の消費者。
singular は、自律型 AI コーディング エージェントを駆動する bash + Python オーケストレーション エンジンです
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
タスクごとのブランチ上の分離された git ワークツリーで単一のタスクを実行します。状態パケット (所有ファイル、変更、証拠) を生成します。監査人がパケットをレビューします。決定者が結果を決定します。
調整サイクル
それぞれの単数形の reconcile --actuate が実行されます。
インポート — ステージングされた L1 タスク提案をオリジン ロックの下の DAG にプルします。
回復 — ワーカーが終了またはタイムアウトした古いリースを再利用します。
統合 — マージコム

完了したワーカー ブランチを git-op ロックの下のターゲット ブランチに追加します。
ディスパッチ — フロンティアタスクを事前にリースし、L2 ワーカーを生成します。
スナップショット — 人間が判読できるプロジェクト状態のスナップショットを書き込みます。
すべての実行中のタスクは、記録を記録するリース ( .singular-state/leases/ 内の JSON ファイル) を保持します。
所有権、再試行回数、有効期限。ワーカーは終了すると状態パケットを書き込みます
( state-packet.v0.schema.json ) 所有ファイル、変更されたファイル、コマンド、テスト、および
証拠。監査人はパケットを検証します。リーパーは後のサイクルでの結果を属性とします。
L2 ワーカーが実行されるたびに、ホストは設定されたゲート コマンド (例: npm test ) を実行します。
ゲート結果 (gate-result.v0.schema.json) は監査モデルに供給され、監査モデルは
監査判定 (audit-verdict.v0.schema.json)。ディサイダーは、
(failure-class, retries-left) と回復アクションのペア — 再試行、スコープの修正、エスカレート、または
park — モデルのラウンドトリップにフォールバックする前に、決定論的な高速パス テーブルを使用します。
単純なコマンドは機能します。exit 0 は成功し、0 以外の場合は失敗します。終了コードは応答できません
可能であればエンジンが使用する 2 つの質問なので、ゲートはオプションで次の質問を書き込むことができます。
ゲート観察 (gate-observation.v0.schema.json) のパスへ
SINGULAR_GATE_REPORT_FILE :
Failure[].signature — 安定した障害ごとの識別子。必須
認識された障害と新しい障害を区別するための特異なゲート ベースライン。
インフラストラクチャ失敗 / インフラストラクチャ理由 - ゲートを実行できませんでした
(依存関係が欠落している、ディスクがいっぱいである、ネットワークに到達できない)。エンジンのレポート
タスクの再試行を費やす代わりに、それを決定的でないインフラストラクチャとして使用する
決して壊れていないコードを修正するようモデルに要求する予算。
単一の init scaffold docs/orchestration/gates/gate.sh を開始点として使用します。
サイドカーは、 schemaVersion: v2 を含めて必須ではありません。ウィズ

ひとつ
エンジンは終了コードと意図的に狭いログセットにフォールバックします。
シグネチャ ( Engine/infra-patterns.tsv ) は、環境障害のみをカバーします。
アプリケーションコードは妥当に生成できません。
切り離されたディスパッチはデフォルトでオンになっています。 SINGULAR_DETACHED_DISPATCH=1 (デフォルト) の場合、
各フロンティアタスクを事前にリースし、独自のセッションでワーカーを生成します。
その後、dispatch-wrap.sh を実行し、数秒以内に戻ります。原点ロックはサイクルの間のみ保持されます。
制御作業。リーパー ( singular_reap_dispatches ) はすべてのノードの先頭で実行されます。
ディスパッチをチェックすることにより、サイクルと属性の完了、失敗、およびクラッシュを適用/実行します
レコード + ワーカー出口ファイル (pid liveness は pid の再利用を無効にします。クラッシュ検出は、
60 分の古いリース ウィンドウから最大 1 サイクルまで）。
これにより、インポート、統合、回復、STATUS、および STOP の応答性が長期間維持されます。
労働者はバックグラウンドで実行されます。
従来の同期バッチ パスを復元するには、SINGULAR_DETACHED_DISPATCH=0 を設定します。
reconcil は、すべてのワーカーが戻る前に待機します。
Bash >= 4、python3、および git 。
PATH 上で少なくとも 1 つのサポートされているランナー CLI ( clude 、 codex 、またはその他)
構成されたランナー)。
macOS ユーザーには、 brew install bash が必要な場合があります。セット
SINGULAR_BASH_BIN=/opt/homebrew/bin/bash シェル/サービス環境で
PATH を並べ替えずにそれを選択します。
複数の Codex インストールが存在する場合は、
SINGULAR_CODEX_BIN=/absolute/path/to/codex 。ドクターとランナーがそれを使用
正確な実行可能ファイルであり、壊れた場合でもフォールバックしません。
# エンジンをクローンして ~/.singular にインストールします
git clone https://github.com/alex-reysa/singular-lite /path/to/singular-lite
cd /path/to/singular-lite
bash インストール.sh
# -> ~/.singular/versions/<ver>/ ~/.singular/current ~/.singular/bin/singular
エクスポート PATH= " $HOME /.singular/bin: $PATH "
Singular の起動名前空間

意図的にクリーンです: SINGULAR_* 、
singular.config.json 、 .singular-version 、 .singular-state/ 、および
SINGULAR_HOME ルートをインストールします (デフォルトは ~/.singular )。発見されなかったり、
このリリースで置き換えられた起動前名前空間をインポートします。各コンシューマーを開始する
単一のセットアップを使用して新しい DAG を作成します。
特異なセットアップ # 「リポジトリ」から検証済みの停止したリポジトリへの 1 つの冪等パス
単数セットアップは個々のライフサイクル動詞を構成し、
命令、証拠、契約。インタプリタ/リポジトリ/gitワークツリーをチェックします。
エンジン PIN を解決します (.singular-version と
singular.config.json のエンジンのバージョンが一致しない)、固定されたエンジンをインストールする場合
これは存在しません。このマシン上にすでに存在する一致するエンジンのチェックアウトからのみ、
ダウンロードメカニズムがないため、最初に .singular-state/STOP を書き込みます。
リポジトリの書き込み、ピンと足場、印刷前にすべてのゲート結果をハッシュし、
移行チェーンを実行し、過去の判定が生き残っていることを確認し、実行します
医師に指示し、教師あり回帰の実行を記録します。前提条件が満たされない前に
何もかもが変異している。状態ラダーを報告します
( インストール → 移行 → 検証済み → 停止準備完了 )、決して起動せず、出力します
Next: 行は 1 つだけです。障害が発生しても安定したコードと 1 回のリカバリが行われます
命令 ( singular.operator-failure.v0 );証拠は下にあります
.singular-state/setup/ 。
singular setup --no-test # 回帰を実行せずに「validated」で停止します
singular setup --test-async # スイートを切り離して開始します。単数形の test --wait で接続します
singular setup --json # 標準出力上の 1 つの singular.setup-report.v0 オブジェクト
構成されたステップは、引き続き単独で使用できます。
singular init # scaffold singular.config.json、docs/orchestration/、.singular-version
単数形のドクター # チェックデプス、エンジン解像度、再

PO設定
singular merge # schemaVersion をエンジンの値に上げます (--dry-run はチェーンのみを出力します)
SINGULAR_BASH_BIN はブートストラップ専用であり、 singular.config.json では無視されます。
singular を呼び出す前に設定してください。 SINGULAR_CODEX_BIN は、
通常のエンジン環境/構成レイヤー。標準の Codex ランナーの場合、
単一のドクターは、制限された --version およびログイン ステータス プローブを実行します。
選択された実行可能ファイルそのもの。
Doctor は、無人実行用の機械可読プリフライトでもあります。
単数形の医師 --json | jq ' .summary、.checks[] | select(.status != "pass") '
単数形の医師 --repair-model-cache # 明示的: 最初にバックアップし、後で再生成します
すべての JSON チェックには、安定した id 、 severity 、 requiredFor 、 remediation があります。
および dedupeKey 。必要な機能の障害がドクターをブロックします。欠落しているオプション
この機能では、複数の役割がそれを共有している場合でも、警告が 1 つ生成されます。医師のチェック
デプロイメント可能な DAG ノードが実際に存在する間のみ、デプロイメント資格情報
準備が整ったフロンティア。 Codex モデル キャッシュをサイレントに削除したり書き換えたりすることはありません
データ: 修復フラグは、オリジナルをタイムスタンプ付きの SHA タグ付きバックアップに移動します。
役割プロファイルは、デフォルトではローカルのみです。追加のツールが必要なリポジトリ、
MCP サーバーまたはプラグインは、遅延プロファイルを明示的に宣言できます。
{
"能力プロファイル" : {
"監査コア" : {
"startup" : " 怠惰 " 、
"required" : [ " ファイルシステム " 、 " git " 、 " スキーマ " 、 " ランナーコントラクト " ]、
"オプション" : [ "mcp:browser " ]
}
}、
"役割プロファイル" : {
"監査人" : "監査コア" ,
"決定者" : " 監査コア "
}
}
機能 ID には、 mcp:NAME 、 plugin:NAME 、 executable:NAME 、または
ファイル:REPO_PATH 。より特殊な機能は、
タイプが buildin 、 executable 、
ファイル、mcp、プラグイン、または環境。
厳密なプロファイルでは

、外部スキル、MCP サーバー、およびプラグインがアクティブ化されている
capabilityArgs.<exact-capability> によってのみ。無関係なプロバイダ引数は決してありません
能力を主張する。従来の SINGULAR_*_EXTRA_ARGS 変数は拒否されます
strict は機能に制限されていないため実行されます。
各リポジトリは、そのエンジン バージョンを .singular-version に固定します (singular.config.json をオーバーライドします)
エンジンのバージョン)。単数形ランチャーは、 ~/.singular/versions/<ver> からそのバージョンを解決します。
SINGULAR_ROOT を現在のリポジトリにバインドし、その構成をロードし、エンジンを実行します。走る
単数形の更新 <ver> を再固定します。
# 調整/実行サイクルを 1 回実行します (インポート → リカバリ → 統合 → ディスパッチ → スナップショット)
単数形の調整 --actuate
# L1 → L2 → 監査を通じて単一のタスクを実行する
単一ドライブ TASK-0001
# 自動運転自律ループ (実時間の予算: SINGULAR_MAX_HOURS)
単数自動
# 所有者とアーティファクトのハッシュにバインドされたヒューマン ゲートを作成、承認、または検査します
単数形のヒューマンゲートリクエスト --help
単数形の human-gate accept --help
特異なヒューマンゲートステータス --help
# ゲート結果のすべての契約違反を一度に報告します。フロンティアを読む
# 最初の違反で停止します (正しくは、無効なゲートに作用してはなりません)。
# これは、開発中のプロモーターが実行ごとに 1 つの違反を学習することを意味します。
特異ゲートバリダ

[切り捨てられた]

## Original Extract

Autonomous multi-agent orchestration engine for software repos — L0/L1/L2 agents, leases, gates, audits, git-worktree isolation. Detached dispatch by default. GPL-3.0. - alex-reysa/singular-lite

GitHub - alex-reysa/singular-lite: Autonomous multi-agent orchestration engine for software repos — L0/L1/L2 agents, leases, gates, audits, git-worktree isolation. Detached dispatch by default. GPL-3.0. · GitHub
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
alex-reysa
/
singular-lite
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
656 Commits 656 Commits Folders and files
claudedocs claudedocs cli cli docs docs engine engine migrations migrations plugin plugin schemas schemas singular-ext singular-ext templates templates tests tests tools tools .gitignore .gitignore .singular-version .singular-version CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SCHEMA_VERSION SCHEMA_VERSION SECURITY.md SECURITY.md VERSION VERSION install.sh install.sh singular.config.json singular.config.json View all files Repository files navigation
TcGUeulhd4hM0DvB5ytcng-under10.mp4
singular
Autonomous multi-agent orchestration for software repos. One engine, many consumers.
singular is a bash + Python orchestration engine that drives autonomous AI coding agents
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
Each singular reconcile --actuate runs:
Import — pull staged L1 task proposals into the DAG under the origin lock.
Recover — reclaim stale leases whose workers have exited or timed out.
Integrate — merge completed worker branches into the target branch under the git-op lock.
Dispatch — pre-lease frontier tasks and spawn L2 workers.
Snapshot — write a human-readable project state snapshot.
Every in-flight task holds a lease (a JSON file in .singular-state/leases/ ) that records
ownership, retry count, and expiry. When a worker finishes it writes a state packet
( state-packet.v0.schema.json ) enumerating owned files, changed files, commands, tests, and
evidence. The auditor validates the packet; the reaper attributes outcomes on later cycles.
After each L2 worker run the host executes the configured gate command (e.g. npm test ).
A gate result ( gate-result.v0.schema.json ) feeds the auditor model, which returns an
audit verdict ( audit-verdict.v0.schema.json ). The decider maps the
(failure-class, retries-left) pair to a recovery action — retry, amend-scope, escalate, or
park — using a deterministic fast-path table before falling back to a model round-trip.
A bare command works: exit 0 passes, non-zero fails. The exit code cannot answer
two questions the engine would use if it could, so a gate may optionally write a
gate observation ( gate-observation.v0.schema.json ) to the path in
SINGULAR_GATE_REPORT_FILE :
failures[].signature — stable per-failure identifiers. Required for
singular gate baseline to tell an acknowledged failure from a new one.
infrastructureFailure / infrastructureReason — the gate could not run
(missing dependencies, full disk, unreachable network). The engine reports
that as inconclusive-infrastructure instead of spending a task's retry
budget asking a model to fix code that was never broken.
singular init scaffolds docs/orchestration/gates/gate.sh as a starting point.
The sidecar is never required — including on schemaVersion: v2 . Without one
the engine falls back to the exit code plus a deliberately narrow set of log
signatures ( engine/infra-patterns.tsv ) covering only environment failures that
application code cannot plausibly produce.
Detached dispatch is ON by default. When SINGULAR_DETACHED_DISPATCH=1 (the default),
reconcile pre-leases each frontier task and spawns the worker in its own session via
dispatch-wrap.sh , then returns within seconds. The origin lock is held only for the cycle's
control work. A reaper ( singular_reap_dispatches ) runs at the top of every
apply/actuate cycle and attributes completions, failures, and crashes by checking dispatch
records + worker exit files (pid liveness defeats pid reuse; crash detection drops from the
60-min stale-lease window to ~one cycle).
This is what keeps import, integrate, recover, STATUS, and STOP responsive while long
workers run in the background.
Set SINGULAR_DETACHED_DISPATCH=0 to restore the legacy synchronous batch path, where
reconcile waits for every worker before returning.
Bash >= 4, python3 , and git .
At least one supported runner CLI on PATH ( claude , codex , or another
configured runner).
macOS users may need brew install bash . Set
SINGULAR_BASH_BIN=/opt/homebrew/bin/bash in the shell/service environment to
select it without reordering PATH .
When multiple Codex installations exist, set
SINGULAR_CODEX_BIN=/absolute/path/to/codex . Doctor and the runner use that
exact executable and do not fall back when it is broken.
# Clone and install the engine to ~/.singular
git clone https://github.com/alex-reysa/singular-lite /path/to/singular-lite
cd /path/to/singular-lite
bash install.sh
# -> ~/.singular/versions/<ver>/ ~/.singular/current ~/.singular/bin/singular
export PATH= " $HOME /.singular/bin: $PATH "
Singular's launch namespace is intentionally clean: SINGULAR_* ,
singular.config.json , .singular-version , .singular-state/ , and the
SINGULAR_HOME install root (default ~/.singular ). It does not discover or
import the pre-launch namespace replaced by this release. Start each consumer
with singular setup and author a fresh DAG.
singular setup # one idempotent path from "a repo" to a verified, STOPPED repo
singular setup composes the individual lifecycle verbs and contributes the
order, the evidence, and the contract. It checks interpreter/repo/git work tree,
resolves the engine pin (naming the winner when .singular-version and
singular.config.json engineVersion disagree), installs the pinned engine when
it is absent — only from a matching engine checkout already on this machine,
since there is no download mechanism — writes .singular-state/STOP as its first
repo write, pins and scaffolds, hashes every gate result before printing and
running the migration chain, verifies those historical verdicts survived, runs
doctor, and records a supervised regression run. Prerequisites fail before
anything is mutated. It reports a state ladder
( installed → migrated → validated → stopped-ready ), never actuates, and prints
exactly one Next: line. Failures carry a stable code and one recovery
instruction ( singular.operator-failure.v0 ); evidence lands under
.singular-state/setup/ .
singular setup --no-test # stop at `validated`, without the regression run
singular setup --test-async # start the suite detached; attach with singular test --wait
singular setup --json # one singular.setup-report.v0 object on stdout
The composed steps are still available on their own:
singular init # scaffold singular.config.json, docs/orchestration/, .singular-version
singular doctor # check deps, engine resolution, repo config
singular migrate # raise schemaVersion to the engine's (--dry-run prints the chain only)
SINGULAR_BASH_BIN is bootstrap-only and is ignored in singular.config.json ;
set it before invoking singular . SINGULAR_CODEX_BIN may be supplied through
the normal engine environment/config layers. For the standard Codex runner,
singular doctor performs bounded --version and login status probes against
the exact selected executable.
Doctor is also the machine-readable preflight for unattended runs:
singular doctor --json | jq ' .summary, .checks[] | select(.status != "pass") '
singular doctor --repair-model-cache # explicit: backup first, then regenerate later
Every JSON check has a stable id , severity , requiredFor , remediation ,
and dedupeKey . Required capability failures block doctor; a missing optional
capability produces one warning even when several roles share it. Doctor checks
deployment credentials only while a deployment-capable DAG node is actually in
the ready frontier. It never silently deletes or rewrites Codex model cache
data: the repair flag moves the original to a timestamped, SHA-tagged backup.
Role profiles are local-only by default. Repositories that need extra tools,
MCP servers, or plugins can declare lazy profiles explicitly:
{
"capabilityProfiles" : {
"audit-core" : {
"startup" : " lazy " ,
"required" : [ " filesystem " , " git " , " schemas " , " runner-contract " ],
"optional" : [ " mcp:browser " ]
}
},
"roleProfiles" : {
"auditor" : " audit-core " ,
"decider" : " audit-core "
}
}
Capability IDs may use mcp:NAME , plugin:NAME , executable:NAME , or
file:REPO_PATH . More specialized capabilities can be declared in the
top-level capabilities registry with a type of builtin , executable ,
file , mcp , plugin , or environment .
In strict profiles, external skills, MCP servers, and plugins are activated
only by capabilityArgs.<exact-capability> ; unrelated providerArgs never
claim a capability. Legacy SINGULAR_*_EXTRA_ARGS variables are rejected for
strict runs because they are not capability-bound.
Each repo pins its engine version in .singular-version (overrides singular.config.json
engineVersion ). The singular launcher resolves that version from ~/.singular/versions/<ver> ,
binds SINGULAR_ROOT to the current repo, loads its config, and execs the engine. Run
singular update <ver> to repin.
# Run one reconcile/actuate cycle (import → recover → integrate → dispatch → snapshot)
singular reconcile --actuate
# Drive a single task through L1 → L2 → audit
singular drive TASK-0001
# Self-driving autonomy loop (wall-clock budget: SINGULAR_MAX_HOURS)
singular auto
# Create, approve, or inspect an owner- and artifact-hash-bound human gate
singular human-gate request --help
singular human-gate approve --help
singular human-gate status --help
# Report every contract violation in a gate-result at once. The frontier read
# stops at the first breach (correctly — it must not act on an invalid gate),
# which means a promoter under development learns about one violation per run.
singular gate valida

[truncated]
