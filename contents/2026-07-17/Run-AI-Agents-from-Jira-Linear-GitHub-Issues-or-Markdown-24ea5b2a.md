---
source: "https://github.com/alexrolls/startup-factory"
hn_url: "https://news.ycombinator.com/item?id=48949104"
title: "Run AI Agents from Jira, Linear, GitHub Issues, or Markdown"
article_title: "GitHub - alexrolls/startup-factory · GitHub"
author: "alexberlinde"
captured_at: "2026-07-17T17:01:06Z"
capture_tool: "hn-digest"
hn_id: 48949104
score: 1
comments: 0
posted_at: "2026-07-17T16:17:41Z"
tags:
  - hacker-news
  - translated
---

# Run AI Agents from Jira, Linear, GitHub Issues, or Markdown

- HN: [48949104](https://news.ycombinator.com/item?id=48949104)
- Source: [github.com](https://github.com/alexrolls/startup-factory)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T16:17:41Z

## Translation

タイトル: Jira、Linear、GitHub Issues、または Markdown から AI エージェントを実行する
記事のタイトル: GitHub - alexrolls/startup-factory · GitHub
説明: GitHub でアカウントを作成して、alexrolls/startup-factory の開発に貢献します。

記事本文:
GitHub - alexrolls/startup-factory · GitHub
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
アレックスロールズ
/
スタートアップファクトリー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く

フォルダーとファイル
125 コミット 125 コミット .github/ workflows .github/ workflows アダプター アダプター bin bin config config docs/ superpowers docs/ superpowers exports extensions/ tracker-backends extensions/ tracker-backends パッケージング パッケージング リファレンス リファレンス ロール ロール src/startup_factory_cli src/startup_factory_cli チーム チーム テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
製品ボードを管理されたソフトウェア配信システムに変えます。
Startup Factory は、エンドツーエンド製品のためのエージェント オーケストレーション フレームワークです。
配達。チームがすでに愛用しているプロジェクト管理ツール、リニア、
Jira、GitHub Issues、ローカル Markdown、または独自のアダプターを使用して、耐久性の高いアダプターを作成します。
AI エージェントのクロスファンクショナル チームのコントロール プレーン。
[タスク] を ToDo (一般的なキュー状態の出荷されたマッピング) に入れます。
自動化が有効でスケジュールされている場合、決定論的な PM スーパーバイザは
デフォルトでは 3 分ごとにボードをチェックし、ラベルが付いたものはすべて残します
人間の仕事を人間にルーティングし、キューに入れられた他のすべてのタスクを明示的なチームによってルーティングします。
プリセット (または設定されたデフォルト) を使用し、アーキテクチャを通じてそれを駆動します。
実装、4 者による In Review ゲート、オプションのスペシャリスト QA、および
統合。また、ブロックされた作業を人間が制御する安全ロックとして監視します。
一致するタスクのワーカーは停止しますが、独立した ToDo 作業は続行されます。いつ
リリース ポリシー、正確な承認、保護されたグリーン CI 証明により許可されます。
別の資格情報分離エグゼキューターがレビューされた不変アーティファクトをデプロイします。
ターゲットを検証してから、親 [機能] を Live として閉じます。
独自のモデル、リポジトリ、スタック、トラッカー、インフラストラクチャを持ち込みます。
プロバイダーに依存しない構造化フックは、

あらゆるクラウドでのターゲットを絞った生産、
を実装できるプラットフォーム、クラスター、データセンター、または内部環境
契約の計画/申請/ステータス/確認。 Startup Factoryが納品いたします
プロトコル、チームトポロジ、決定論的ランタイム、リカバリモデル、安全性
それらの周りの境界線。
プロジェクト管理ネイティブ · マルチモデル · クラウドに依存しない · フェールクローズ ·
監査可能
これは隠れたチャットボット ループではありません。計画、所有権、進捗状況、決定、
証拠、阻害要因、承認、ポリシーの拒否、展開状態はそのまま残ります。
チームが製品を管理するのと同じボードに表示されます。トラッカーテキストと
主張される著者名はワークフローの証拠であり、セキュリティ認証や
制作権限。
ToDo -> 進行中 -> レビュー中 -> 運用準備完了 -> デプロイ -> ライブ
調査結果 -> ToDo -> 新たな試み -> 進行中
赤 / 保留中 / CI がありません -----------X どこにでもデプロイ
デフォルトで安全: ボードの自動化と本番納品の両方が出荷されます
無効になっています。通常のエージェントは本番認証情報を受け取ることはありません。を有効にする
リリースには、保護された外部構成、フック、アイデンティティが必要です。
検証。リリースコミットに対する正確な最新のグリーン CI/CD プルーフのみ
展開を許可します。
AI ビルダーのための階層化された安全境界線
トラッカーの完全な透明性
クイック スタート (2 分、アカウントなし)
基板と製品の納品を自動化する
AI ビルダーのための階層化された安全境界線
自律エージェントは、その周囲の境界と同じくらい役に立ちます。スタートアップ
工場ではフェールクローズされたワークフローとリリース制御が提供されますが、リポジトリ コード
オペレーティング システムのセキュリティ境界ではありません。通常のエージェントには引き続き
実際のワークツリーをスコープとしたサンドボックスと最小権限の ID:
コード所有の運用ポリシー ゲート。 bin/policy-check.py は次の間隔で画面を表示します
特権リリースフックコマンドといいえ

その前に正式な生産計画を立てた
サブプロセスが開始されます。その否定
ベースライン — シェル構成、権限昇格、ファイルシステム/データベース/
インフラストラクチャの破壊、機密情報のダンピング、メタデータ資格情報へのアクセス、
エンコードされたコマンドのバイパス — コードによって所有されます。プロジェクト構成は次のとおりです。
拒否を追加し、決して削除しないでください。
3 層の権限モデル。すべてのアクションは DENY に解決されます。
人間の承認が必要 、または許可 (reference/guardrails.md)。
承認は正確なダイジェスト、環境、ターゲット、有効期限、および 1 回限りの使用をバインドします
ノンス。沈黙は決して承認しません。不明なものはすべて否定されます。
[タスク] レベルのリリース拒否証跡。正規化された生産計画の拒否
[DENIED ACTION] として冪等に投影されます。
bin/tracker-ops.sh 記録拒否 。その他のランチャー、パス、キュー、および ID
プリフライト拒否はミューテーションの前に失敗し、保護されたランタイム ログに残ります。
これらはトラッカー レコードとして誤って宣伝されることはありません。
最小権限のエージェント サンドボックス。強制モードでは、すべての LLM プロセスと
WORKTREE_SETUP は、AGENT_SANDBOX_RUNNER --workdir <absolute> -- /usr/bin/env -i ... として実行されます。ランチャーは外部で保護された実行可能ファイルのみを受け入れます
リポジトリ;そのランナーはファイルシステム、プロセス、ネットワーク、および IAM を強制する必要があります
孤立。ブローカー モードでは、チーム リーダーを含む LLM はトラッカーを受信しません
資格情報。
含まれるワークスペース。すべての実装者は独自の git 内に分離されます
ワークツリー;トラッカー ファイル パスはシンボリックリンク セーフであり、設定されたパスに限定されます。
ルート;統合とターミナル遷移はシリアル化され、次のように検証されます。
リードバックし、専用コンポーネントに予約されています。
デフォルトではオフになっているオートメーション。ポートフォリオスーパーバイザーは、
決定的 (LLM ではない)、明示的に有効になるまで無効になり、
何かが不正な形式の場合、状態を作成するのではなくパスします。
ビルトイン権限po

licy は意図的にシンプルです:
バイパス不可能な完全なリストと施行境界は次のとおりです。
参照/ガードレール.md 。
その結果、実際の配送作業を AI エージェントに引き渡し、自分の端末から検査できるようになります。
トラッカー、ワークフロー アクター、ワークフロー承認、プロダクション承認 ID、および
配信ごとのポリシー拒否。認証された本番承認者の ID
保護されたトランザクション状態のままです。トラッカーの作成者は決して扱われません
認証。
利点
それがあなたに与えるもの
混乱をマージせずに迅速に移動
決定的ディスパッチャーは、設計が承認され、依存関係に対応し、リソースに安全な作業のみを起動します。試行ごとに独自のタスク ブランチとワークツリーが取得されます。統合はシリアル化されたままになります。
エージェントの出力だけでなく、配信全体を確認する
[タスク] ごとに 1 つのライブ [進行状況] レコードと、[機能] ごとに 1 つの [ダイジェスト] レコードで、プロジェクト管理ツールのトラッカー ステータス、実行ステージ、アクター、試行が表示されます。
各ジョブに適切なモデルを使用する
Claude、Codex、Gemini、または任意のファイル読み取り CLI をロールごとに組み合わせて、個々のタスクを高速、標準、または強力なモデル プロファイルにルーティングします。
品質ゲートを明示的に保つ
アーキテクチャの承認は実装に先立って行われます。 In Review では、1 つの正確なパッケージに関して、プリンシパル アーキテクト、懐疑的なプリンシパル アーキテクト、シニア セキュリティ エンジニア、およびチーム リードからの独立した承認が必要です。オプションの QA によって証拠が追加される場合があり、インテグレーターはマージ前にビルド、テスト、および lint コマンドを実行します。
再起動ではなく回復する
不変のタスク パケット、永続的なイベント、チェックポイント ブランチ、冪等の送信ボックス、試行を認識した再起動により、中断された作業を検査して回復できるようになります。
スタックとトラッカーを保管してください
同じワークフローが、言語、フレームワーク、LLM、プロジェクト管理ツール間で実行されます。 Markdown をオフラインで開始し、プロセスを書き直すことなくアダプターを切り替えます。
ターンT

彼は安全な配達の列に乗り込みます
確定的な cron/サービス パスは、キューに入れられた/ブロックされた作業を監視し、実行中の実行を復元し、明示的なチーム プリセットを選択し、対象となるキューに入れられたタスクに対してのみ LLM を起動します。
ファクトリを停止せずに 1 つのタスクを一時停止する
次のスキャンでは、[ブロック] により、一致するタスクのみが即座にフェンスされます。独立した ToDo 作業やその他の機能は継続します。人間だけがロックを解除できます。
危険な権限をエージェントから遠ざける
1 つの拒否/承認/許可契約がすべての役割を管理します。コード ゲートは、危険な特権リリース フックと計画をブロックします。必要な OS サンドボックスと最小権限の ID により、通常のエージェント ファイルが適用されます
[切り捨てられた]
オブラ/スーパーパワーズとStartup Factory
異なる補完的なレベルで動作します。
Superpowers は、エージェント スキルとしてパッケージ化されたエンジニアリング方法論です。いっぱいです
上流のワークフローはブレインストーミング、実装計画、ワークツリーをカバーできます
作成、タスクの実行、TDD、デバッグ、コードレビュー、検証、ブランチ
完成。 Startup Factory はプロジェクト配信のコントロール プレーンです。
製品、アーキテクチャ、実装、セキュリティ、品質、統合、
ポートフォリオの自動化、および永続的なトラッカーの状態全体にわたる本番リリースと
複数の独立したエージェント。
統合統合では、各システムの最も強力な部分を意図的に使用します。
2 つの実行オーケストレーターを実行せずに:
スーパーパワーはアイデアを承認された仕様と詳細な形に形作ります。
実施計画。
Startup Factory は、これらのドキュメントを入力としてレビューし、
配信を追跡し、検証された本番環境を通じて実行を所有します。
クロード タスク ワーカーは、TDD、デバッグ、
割り当てられた 1 つのタスク内でレビューと新たな検証を受け取ります。
Superpowers は 2 番目のワークツリーを作成せず、seco をディスパッチします。

2番目のチーム、実行します
計画を立てるか、ブランチをマージするか、機能のリリースを宣言します。
フローチャート LR
アイデア[「製品のアイデアまたはチケット」]
ブレインストーミング[「超大国のブレインストーミング<br/>質問、代替案、承認されたデザイン」]
Plan["スーパーパワーの書き込み計画<br/>正確なファイル、インターフェイス、テスト、コマンド"]
ハンドオフ["ダイジェストバインドされた計画のハンドオフ<br/>コミットされた仕様 + 計画"]
形状[「Startup Factory 計画レビュー<br/>製品・リード・プリンシパル・懐疑的」]
Tracker[「追跡される機能とタスク<br/>承認 · 依存関係 · リソース」]
Execute["Startup Factory の実行<br/>パケット、ブランチ、ワークツリー、ディスパッチ"]
メソッド[「クロードローカルの超大国メソッド<br/>TDD · デバッグ · レビュー応答 · 検証」]
レビュー[「正確なパッケージのレビュー委員会<br/>リーダー、プリンシパル、懐疑的、セキュリティ」]
Integrate["シリアル化された統合<br/>検証された機能ブランチ"]
Release[「保護された CI とリリース<br/>ポリシー · 承認 · デプロイ · 検証」]
アイデア --> ブレインストーミング --> 計画 --> ハンドオフ --> 形状 --> トラッカー --> 実行
実行 --> メソッド --> レビュー --> 統合 --> リリース
読み込み中
なぜライフサイクルを分割するのか
理由
メリット
ステージごとに 1 つの権限
仕様には 1 つのソースがあり、トラッカーには 1 つのワークフロー所有者があり、各タスクには 1 つのアクティブな試行があります。

[切り捨てられた]

## Original Extract

Contribute to alexrolls/startup-factory development by creating an account on GitHub.

GitHub - alexrolls/startup-factory · GitHub
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
alexrolls
/
startup-factory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
125 Commits 125 Commits .github/ workflows .github/ workflows adapters adapters bin bin config config docs/ superpowers docs/ superpowers exports exports extensions/ tracker-backends extensions/ tracker-backends packaging packaging reference reference roles roles src/ startup_factory_cli src/ startup_factory_cli teams teams tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml View all files Repository files navigation
Turn your product board into a governed software delivery system.
Startup Factory is an agentic orchestration framework for end-to-end product
delivery. Connect the project-management tool your team already loves—Linear,
Jira, GitHub Issues, local Markdown, or your own adapter—and make it the durable
control plane for a cross-functional team of AI agents.
Put a [task] in ToDo—the shipped mapping of the generic queued state.
When automation is enabled and scheduled, the deterministic PM supervisor
checks the board every three minutes by default, leaves anything labeled
human-work to people, routes every other queued task by its explicit team
preset (or your configured default), and drives it through architecture,
implementation, a four-party In Review gate, optional specialist QA, and
integration. It also observes Blocked work as a human-controlled safety lock:
the matching task workers stop, while independent ToDo work continues. When
your release policy, exact approval, and protected green CI proof allow it, a
separate credential-isolated executor deploys the reviewed immutable artifact,
verifies the target, and only then closes the parent [feature] as Live .
Bring your own models, repository, stack, tracker, and infrastructure.
Provider-neutral structured hooks can target production in any cloud,
platform, cluster, datacenter, or internal environment that can implement the
plan/apply/status/verify contract. Startup Factory supplies the delivery
protocol, team topology, deterministic runtime, recovery model, and safety
boundaries around them.
Project-management native · Multi-model · Cloud agnostic · Fail closed ·
Auditable
This is not a hidden chatbot loop. Plans, ownership, progress, decisions,
evidence, blockers, approvals, policy denials, and deployment state remain
visible in the same board where your team manages the product. Tracker text and
claimed authorship are workflow evidence, never security authentication or
production authority.
ToDo -> In Progress -> In Review -> Ready for production -> deploy -> Live
findings -> ToDo -> fresh attempt -> In Progress
red / pending / missing CI ----------------------X deploy anywhere
Safe by default: board automation and production delivery both ship
disabled. Ordinary agents never receive production credentials; enabling a
release requires protected external configuration, hooks, identities, and
verification. Only an exact, current green CI/CD proof for the release commit
permits deployment.
Layered safety boundaries for AI builders
Full transparency in your tracker
Quick Start (2 minutes, no accounts)
Automate the board and production delivery
Layered safety boundaries for AI builders
Autonomous agents are only as useful as the boundaries around them. Startup
Factory supplies fail-closed workflow and release controls, but repository code
is not an operating-system security boundary. Ordinary agents still require a
real worktree-scoped sandbox and least-privilege identity:
A code-owned production policy gate. bin/policy-check.py screens every
privileged release-hook command and normalized production plan before that
subprocess starts. Its deny
baseline — shell composition, privilege escalation, filesystem/database/
infrastructure destruction, secret dumping, metadata-credential access,
encoded-command bypasses — is owned by the code: project configuration can
add denials, never remove one.
A three-tier authority model. Every action resolves to DENY ,
REQUIRE HUMAN APPROVAL , or ALLOW ( reference/guardrails.md ).
Approvals bind exact digests, environments, targets, expirations, and one-use
nonces. Silence never approves; unknown anything is denied.
A [task] -level release-denial trail. A normalized production-plan denial
is projected idempotently as [DENIED ACTION] through
bin/tracker-ops.sh record-denial . Other launcher, path, queue, and identity
preflight refusals fail before mutation and remain in protected runtime logs;
they are not falsely advertised as tracker records.
Least-privilege agent sandboxes. In enforced mode every LLM process and
WORKTREE_SETUP runs as AGENT_SANDBOX_RUNNER --workdir <absolute> -- /usr/bin/env -i ... . The launcher accepts only a protected executable outside
the repository; that runner must enforce filesystem, process, network, and IAM
isolation. In broker mode no LLM—including the team lead—receives tracker
credentials.
Contained workspaces. Every implementer is isolated in its own git
worktree; tracker file paths are symlink-safe and confined to their configured
root; integration and terminal transitions are serialized, verified by
read-back, and reserved to dedicated components.
Automation that is off by default. The portfolio supervisor is
deterministic (not an LLM), disabled until explicitly enabled, and stops the
pass rather than fabricating state when anything is malformed.
The built-in authority policy is intentionally simple:
The full non-bypassable list and enforcement boundary are in
reference/guardrails.md .
The result: you can hand real delivery work to AI agents and inspect, from your
tracker, the workflow actor, workflow approvals, production approval ID, and
policy denials for each delivery. Authenticated production approver identity
remains in protected transaction state; tracker authorship is never treated as
authentication.
Advantage
What it gives you
Move fast without merge chaos
A deterministic dispatcher launches only design-approved, dependency-ready, resource-safe work. Each attempt gets its own task branch and worktree; integration stays serialized.
See the whole delivery, not just agent output
One live [progress] record per [task] and one [digest] per [feature] show tracker status, execution stage, actor, and attempt in your project-management tool.
Use the right model for each job
Mix Claude, Codex, Gemini, or any file-reading CLI by role, then route individual tasks to fast, standard, or strong model profiles.
Keep quality gates explicit
Architecture approval precedes implementation. In Review requires independent approvals from the Principal Architect, Sceptical Principal Architect, Senior Security Engineer, and Team Lead over one exact package; optional QA may add evidence, and the integrator runs your build, test, and lint commands before merging.
Recover instead of restarting
Immutable task packets, durable events, checkpoint branches, an idempotent outbox, and attempt-aware relaunches make interrupted work inspectable and recoverable.
Keep your stack and your tracker
The same workflow runs across languages, frameworks, LLMs, and project-management tools. Start offline with Markdown and switch adapters without rewriting the process.
Turn the board into a safe delivery queue
A deterministic cron/service pass observes queued/blocked work, restores in-flight runs, chooses an explicit team preset, and launches LLMs only for eligible queued tasks.
Pause one task without stopping the factory
On the next scan, [Blocked] immediately fences only the matching task. Independent ToDo work and other features continue; only a human can unlock it.
Keep dangerous authority out of agents
One deny/approval/allow contract governs every role. The code gate blocks dangerous privileged release hooks and plans; your required OS sandbox and least-privilege identities enforce ordinary-agent filesy
[truncated]
obra/superpowers and Startup Factory
operate at different, complementary levels.
Superpowers is an engineering methodology packaged as agent skills. Its full
upstream workflow can cover brainstorming, implementation planning, worktree
creation, task execution, TDD, debugging, code review, verification, and branch
completion. Startup Factory is a project delivery control plane: it coordinates
product, architecture, implementation, security, quality, integration,
portfolio automation, and production release across durable tracker state and
multiple independent agents.
The combined integration deliberately uses the strongest part of each system
without running two execution orchestrators :
Superpowers shapes the idea into an approved specification and detailed
implementation plan.
Startup Factory reviews those documents as inputs, creates and governs the
tracked delivery, and owns execution through verified production.
Claude task workers may use focused Superpowers methods for TDD, debugging,
receiving review, and fresh verification inside their one assigned task.
Superpowers does not create a second worktree, dispatch a second team, execute
the plan, merge the branch, or declare the feature released.
flowchart LR
Idea["Product idea or ticket"]
Brainstorm["Superpowers brainstorming<br/>questions · alternatives · approved design"]
Plan["Superpowers writing-plans<br/>exact files · interfaces · tests · commands"]
Handoff["Digest-bound planning handoff<br/>committed spec + plan"]
Shape["Startup Factory planning review<br/>Product · Lead · Principal · Sceptical"]
Tracker["Tracked feature and tasks<br/>acceptance · dependencies · resources"]
Execute["Startup Factory execution<br/>packets · branches · worktrees · dispatch"]
Methods["Claude-local Superpowers methods<br/>TDD · debugging · review response · verification"]
Review["Exact-package review board<br/>Lead · Principal · Sceptical · Security"]
Integrate["Serialized integration<br/>validated feature branch"]
Release["Protected CI and release<br/>policy · approval · deploy · verify"]
Idea --> Brainstorm --> Plan --> Handoff --> Shape --> Tracker --> Execute
Execute --> Methods --> Review --> Integrate --> Release
Loading
Why divide the lifecycle
Reason
Benefit
One authority per stage
The specification has one source, the tracker has one workflow owner, each task has one active attempt, th

[truncated]
