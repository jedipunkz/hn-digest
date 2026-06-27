---
source: "https://github.com/KayhanB21/riskratchet"
hn_url: "https://news.ycombinator.com/item?id=48700587"
title: "Riskratchet: Stop AI-generated code from rotting your codebase"
article_title: "GitHub - KayhanB21/riskratchet: A maintainability ratchet for AI-assisted Python. · GitHub"
author: "keynha"
captured_at: "2026-06-27T19:29:50Z"
capture_tool: "hn-digest"
hn_id: 48700587
score: 2
comments: 0
posted_at: "2026-06-27T18:37:00Z"
tags:
  - hacker-news
  - translated
---

# Riskratchet: Stop AI-generated code from rotting your codebase

- HN: [48700587](https://news.ycombinator.com/item?id=48700587)
- Source: [github.com](https://github.com/KayhanB21/riskratchet)
- Score: 2
- Comments: 0
- Posted: 2026-06-27T18:37:00Z

## Translation

タイトル: Riskratchet: AI 生成コードによるコードベースの腐敗を阻止する
記事のタイトル: GitHub - KayhanB21/riskratchet: AI 支援 Python の保守性ラチェット。 · GitHub
説明: AI 支援 Python の保守性ラチェット。 GitHub でアカウントを作成して、KayhanB21/riskratchet の開発に貢献してください。

記事本文:
GitHub - KayhanB21/riskratchet: AI 支援 Python 用の保守性ラチェット。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ケイハンB21
/
リスクラチェット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスターブランチ

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
83 コミット 83 コミット .github .github アセット アセット bin bin データ/ キャリブレーション データ/ キャリブレーション docs docs スキーマ スキーマ src/ リスクラチェット src/ リスクラチェット テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .pre-commit-hooks.yaml .pre-commit-hooks.yaml .riskratchet.json .riskratchet.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md TODO.md TODO.md action.yml action.yml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 支援 Python の保守性ラチェット。バーは下にのみ移動できます。
AI コーディング エージェントは、コンパイル、実行、および渡すコードの作成に非常に優れています。
付属のテスト。彼らは以下のことがあまり得意ではありません。
新しいコードの有意義なテストを作成し、
30行の関数がいつの間にか130行になっていることに気づき、
パブリック API がテストで呼び出し元のない関数を公開していることをキャッチし、
小さなリファクタリングにより、if ラダーが 14 方向のサイクロマティック モンスターに変わったことに気づきました。
従来のレビューでは、この点の一部がわかります。ラチェットがすべてをキャッチし、
機械的に、毎回。リスクラチェットは、関数ごとのリスク スコアを次から計算します。
カバレッジギャップ、循環的複雑さ、チャーン、公共表面、スプロール現象、その後
リスクがベースラインを超えて増大するたびに、CI が失敗するか、コミットがブロックされます。誰もそうする必要はありません
複雑な警官をプレイします。
レビュー ワークフローは以下からインスピレーションを得ています。
カーゴクラップ (CRAP を作成した)
ベースライン、PR コメント、JSON 出力を含む CI で実用的なメトリック)、および
カーソルの熱核コード品質レビュー
エージェント プロンプト (保守性、構造、スプロール性、および
明示的な境界）。リスクラチェットはカーゴの Python ポートでもなければ、
エージェントプロンプト: それはレポートします

ts CRAP を作成し、Python 固有のシグナルを上部に追加します (ブランチ
ギャップ、チャーン、公共の場、スプロール）。
pip インストール リスクラチェット
# またはインストールせずに実行する
uvx リスクラチェット --ヘルプ
# 1. JSON 形式のカバレッジを使用してテストを実行します
pytest --cov --cov-report=json:coverage.json
# 2. 現在のリスクプロファイルのスナップショットを作成する
リスクラチェット ベースライン src --coverage Coverage.json --output .riskratchet.json
#3. キャプチャされたものを検査する
リスクラチェット スキャン src --coverage Coverage.json
#4. リスクが後退するとビルドが失敗する
リスクラチェット チェック src --coverage Coverage.json --baseline .riskratchet.json
リスクラチェット チェックは、リグレッションの場合は 1、使用上のエラーの場合は 2 で終了します (例: 欠落しています)
ベースライン)、それ以外の場合は 0。
ベースラインが存在する前に早期導入するには、 --fail-above Ngates をチェックしてください
ベースラインを必要とせずに絶対しきい値に基づく (ベースライン ゲーティング)
成熟したコードベースの推奨モードのままです):
# ベースラインはまだありません: いずれかの関数のスコアが 60 を超える場合は失敗します。
リスクラチェット チェック src --coverage Coverage.json --fail-above 60
# scan は、ベースラインのないゲート (異なる出口/出力形状) も公開します。
リスクラチェット スキャン src --coverage Coverage.json --fail-above 75
リスクラチェット スキャン src --coverage Coverage.json --fail-severity high
--baseline と --fail-above の両方が指定された場合、ベースライン ゲート
は権限があり、 --fail-above は無視され、stderr 警告が表示されます。
Riskratchet init は、[tool.riskratchet] セクションをスキャフォールディングします。
pyproject.toml を開き、すぐに貼り付けられる CI スニペットを出力します。と
--with-baseline (または、対話型プロンプトに対して「はい」と答えることによって)
pytest が検出された場合の TTY)、pytest --cov も実行され、
ベースラインを一度に実行:
Riskratchet init # 構成を書き込み、スニペットを出力
riskratchet init --with-baseline # pytest --cov + Baseline も実行します
Riskratchet init --force # 既存の [tool.riskratchet] を置き換えます
リスクラチェットドクターはsです

ix-check プリフライトで任意の名前を付ける
チェックの開始に失敗する可能性があります (パスの欠落、欠落/不正な形式)
ベースライン、欠落/古いカバレッジ、Git 履歴なし、不明な構成
キー、無効な抑制) を修正し、正確な修正コマンドを出力します。
それぞれ。ステータス テーブルは標準出力に出力されます。 → 修正: 修復が進みます
stderr に転送して、個別にパイプできるようにします。
リスクラチェット ドクター # 人間が読めるテーブル + 修復
リスクラチェットドクター --json # schemas/doctor.schema.json に対して検証します
リスクラチェットドクター 2> /dev/null # ステータステーブルのみ
リスクラチェットドクター > /dev/null # 修復コマンドのみ
医師は、すべてのチェックが合格または警告された場合にのみ 0 を終了します。単一の
失敗した終了 1 。意図したワークフローは、初期化→ドクター→修正です。
警告→ベースライン→チェック。
複合アクションは action.yml で提供されるため、採用者はこれを行う必要はありません。
ワークフロー ファイルをコピーします — 使用: KayhanB21/riskratchet@v0.2.13 は
正規の参照。このアクションは、uv ツール install を介してリスクラチェットをインストールし、ベースラインとベースラインの両方でチェック ( --format pr-comment を実行します) を実行します。
ベースラインなしモード)、スティッキー PR コメントを更新/挿入し、
PR チェックでリグレッションが反映されるように、終了ステータスをチェックします。
# .github/workflows/riskratchet.yml
上: [プルリクエスト]
ジョブ:
リスクラチェット:
実行: ubuntu-最新
権限:
内容：読む
プルリクエスト : 書き込み
手順:
- 使用:actions/checkout@11bd71901bbe5b1630ceea73d27597364c9af683 # v4.2.2
- 使用: KayhanB21/riskratchet@v0.2.13
付き:
カバレッジ : カバレッジ.json
入力 (括弧内はデフォルト): パス ( [tool.riskratchet] paths )、カバレッジ (自動検出)、ベースライン ( .riskratchet.json
— ファイルが見つからない場合、アクションは --fail-above モードで実行されます)、
失敗した上 ( 60 )、コメント ( true )、Python バージョン ( 3.12 )、
リスクラチェット バージョン (PyPI の最新)、github トークン
( ${{ github.token }} )。
マーケットプレイス発見用

、KayhanB21/リスクラチェットアクション
ラッパー リポジトリが推奨されるエントリ ポイントです。それはに委任します
action.yml をルートにして、両方の形状が 1 つの真実の情報源を共有するようにします。
タグ付けされたリリースはすべて、サプライチェーンの出所を検査できるように出荷されます: CycloneDX
ホイールのランタイム依存関係クロージャーの SBOM (sbom ワークフロー アーティファクト)、
ホイールと sdist、および PEP 740 で署名された GitHub ビルド来歴証明書
Trusted Publishing からの PyPI 証明書。ダウンロードしたホイールが構築されたことを確認するには
このリポジトリによると:
gh 認証検証リスクラチェット- < バージョン > -py3-none-any.whl --owner KayhanB21
docs/threat-model.md を参照してください。
各アーティファクトが何を保証し、何を保証しないのか。
標準的なユースケース: AI エージェント + サイド プロジェクト
あなたは 8 か月間、AI エージェントを使用して FastAPI バックエンドのバイブコーディングを行ってきました。
それは機能します、テストは緑色っぽい (カバー率 62%)、しかしあなたは今気づきました
services/billing.py::reconcile_subscriptions は静かに 180 行に増加しました。
書いた覚えのない 11 方向一致ステートメント。
pip インストール リスクラチェット
pytest --cov --cov-branch --cov-report=json:coverage.json
リスクラチェット スキャン src --coverage Coverage.json --top 10
reconcile_subscriptions はスコア 71 (高) で表示されます。
構造的複雑さ: 90 、スプロール: 55 、カバレッジ_ギャップ: 60 。あなたも見つけます
驚き: 12 行の公共ユーティリティ _normalize_plan_id のスコアは 48 です。
テストはゼロです。バーのスナップショットを作成します。
リスクラチェット ベースライン src --coverage Coverage.json --output .riskratchet.json
git add .riskratchet.json && git commit -m "リスクラチェット ベースラインを追加"
ここから、エージェントが Webhook ハンドラーを追加するたびに、または「リファクタリング」が行われます。
請求フロー」を確認してから、コミットする前にリスクラチェット チェックを実行してください。静かに膨れ上がったら
reconcile_subscriptions 180 から 220 行、チェック出口 1 と名前
回帰。見ることを覚えておく必要がなくなります。
なぜこれがCAなのか

非非的なユースケース: AI エージェントは追加に優れています
コード、それが状況を悪化させたことに気づくのは平凡です。ベースラインはあなたのものです
記憶。
CI での PR のチーム ゲーティング。 pytest --cov とriskratchet check --format pr-comment を実行します。
GitHub アクション内。 gh pr コメントにパイプします。 PR コメントの形式は次で始まります。
<!-- Riskratchet-report --> これにより、ボットはプッシュごとに同じコメントを更新します
スパム送信の代わりに。ラチェットは機械式で所有者がいないため、誰も所有する必要はありません
コードレビューで「複雑さの警官」を演じる。参照
AI コーディング エージェントのリスクラチェットを使用します。
ソロ リポジトリの事前コミット フック。 pytest-cov とリスクラチェットをワイヤリングする
.pre-commit-config.yaml に保存すると、コミットごとにカバレッジが再生成され、
回帰なしでコミットをゲートします。参照
プリコミット統合。
醜い関数を 1 つ調査します。 Riskratchet Explain path/to/file.py::qualname を使用して 6 つのコンポーネントのスコアをダンプし、
ドライバー (複雑さ、カバレッジ、スプロール性)。リファクタリング後、実行します
リスクラチェット diff --json | jq '.improved[], .regressed[]' を証明する
単にデッキチェアを並べ替えただけではなく、変化は正味のプラスでした。
CRAP だけでは有用だが不完全な理由
古典的な CRAP スコア ( CC^2 * (1 - line_coverage)^3 + CC ) は 1 つをキャッチします。
悪いコードの形状: 複雑でテストが不十分です。それは切実な問題なのですが、
同じくらい頻繁に運用環境に出荷される他のいくつかの機能が欠落しています。
複雑さが低く、テストが不要な関数。 CRAP はそれに CC (単一の
桁）。リスクは存在しますが、目に見えません。
全行をカバーする関数ですが、すべてのブランチが同じ方法でカバーされます。
CRAP はライン カバレッジのみを調べます。
誰もが触れることを恐れる 2,000 行のモジュール内の関数。スプロールは
CRAPには見えません。
過去 90 件のコミットのうち 40 件で変更された関数。チャーンは目に見えない
クラップに。
リスクラチェットは CRAP を報告された指標として保持し、独自の複合値を計算します
6 重さからのスコア

コンポーネントが保護されているため、他のリスクも発生します。
プリコミットとリスクラチェットがどのように連携するか
リスクラチェットにとってプリコミットに関する 2 つの重要な点:
事前コミットでは、フックを実行する前にステージングされていない編集を非表示にします。フックのみ
実際にコミットしようとしているコードを確認してください。一般的には便利ですが、
つまり、リスクラチェットは、ユーザーが開いているソース ツリーとは異なるソース ツリーを参照します。
編集者。
各言語: Python フックは、独自の分離された virtualenv で実行されます。
プロジェクトの pytest ではなく、riskratchet とその宣言された deps が含まれています。
アプリケーション コード、またはテスト プラグイン。
これらを組み合わせると 1 つの要件が作成されます。それは、coverage.json リスクラチェットの読み取りです。
分析しているものと同じ隠しソース ツリーを反映する必要があります。古いものを再利用する
プリコミット前のcoverage.jsonは編集ドリフトソースを隠しておき、
カバレッジが同期していません。
そのため、公開されたフックにはデフォルトで --no-auto-cov --allow-missing-coverage が付属しており、安全ですが制限されています。いずれかを選択してください
以下のパターンを参考にしてください。
パターン A: 兄弟フックでカバレッジを事前生成する (推奨)
カバレッジが一致するように、同じプリコミット チェーン内で pytest --cov を実行します。
まさに隠されたツリーです。
リポジトリ:
- リポジトリ: ローカル
フック:
- ID : pytest-cov
名前: pytest --cov (リスクラチェット用のcoverage.jsonを生成します)
エントリ:

[切り捨てられた]

## Original Extract

A maintainability ratchet for AI-assisted Python. Contribute to KayhanB21/riskratchet development by creating an account on GitHub.

GitHub - KayhanB21/riskratchet: A maintainability ratchet for AI-assisted Python. · GitHub
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
KayhanB21
/
riskratchet
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
83 Commits 83 Commits .github .github assets assets bin bin data/ calibration data/ calibration docs docs schemas schemas src/ riskratchet src/ riskratchet tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .pre-commit-hooks.yaml .pre-commit-hooks.yaml .riskratchet.json .riskratchet.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md TODO.md TODO.md action.yml action.yml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A maintainability ratchet for AI-assisted Python. The bar can only move down.
AI coding agents are very good at writing code that compiles, runs, and passes
the tests they ship with it. They are less good at:
writing meaningful tests for the new code,
noticing a 30-line function quietly became 130 lines,
catching that the public API now exposes a function with no callers in tests,
realising a small refactor turned an if ladder into a 14-way cyclomatic monster.
A traditional review catches some of this. A ratchet catches all of it,
mechanically, every time. riskratchet computes a per-function risk score from
coverage gaps, cyclomatic complexity, churn, public surface, and sprawl, then
fails CI or blocks the commit whenever risk grows past a baseline. Nobody has to
play complexity cop.
The review workflow is inspired by
cargo-crap (which made the CRAP
metric practical in CI with baselines, PR comments, and JSON output) and
Cursor's thermo-nuclear-code-quality-review
agent prompt (which emphasises maintainability, structure, sprawl, and
explicit boundaries). riskratchet is neither a Python port of cargo-crap nor an
agent prompt: it reports CRAP and adds Python-specific signals on top (branch
gaps, churn, public surface, sprawl).
pip install riskratchet
# or run without installing
uvx riskratchet --help
# 1. run your tests with coverage in JSON form
pytest --cov --cov-report=json:coverage.json
# 2. snapshot the current risk profile
riskratchet baseline src --coverage coverage.json --output .riskratchet.json
# 3. inspect what was captured
riskratchet scan src --coverage coverage.json
# 4. fail the build when risk regresses
riskratchet check src --coverage coverage.json --baseline .riskratchet.json
riskratchet check exits 1 on regressions, 2 on usage errors (e.g. missing
baseline), and 0 otherwise.
For early adoption before a baseline exists, check --fail-above N gates
on an absolute threshold without requiring a baseline (baseline gating
remains the recommended mode for mature codebases):
# No baseline yet: fail if any function scores above 60.
riskratchet check src --coverage coverage.json --fail-above 60
# scan also exposes a no-baseline gate (different exit/output shape).
riskratchet scan src --coverage coverage.json --fail-above 75
riskratchet scan src --coverage coverage.json --fail-severity high
When --baseline and --fail-above are both given, the baseline gate
is authoritative and --fail-above is ignored with a stderr warning.
riskratchet init scaffolds a [tool.riskratchet] section in
pyproject.toml and prints a ready-to-paste CI snippet. With
--with-baseline (or by saying yes to the interactive prompt on a
TTY when pytest is detected), it also runs pytest --cov and creates
the baseline in one go:
riskratchet init # write config, print snippet
riskratchet init --with-baseline # also run pytest --cov + baseline
riskratchet init --force # replace existing [tool.riskratchet]
riskratchet doctor is a six-check pre-flight that names whatever
would make check fail to start (missing paths, missing/malformed
baseline, missing/stale coverage, no git history, unknown config
keys, invalid suppressions) and prints the exact fix command for
each. The status table goes to stdout; the → fix: remediations go
to stderr so you can pipe them separately:
riskratchet doctor # human-readable table + remediation
riskratchet doctor --json # validates against schemas/doctor.schema.json
riskratchet doctor 2> /dev/null # status table only
riskratchet doctor > /dev/null # remediation commands only
doctor exits 0 only when every check is pass or warn; a single
fail exits 1 . The intended workflow is init → doctor → fix the
warnings → baseline → check .
The composite action ships in action.yml so adopters don't have to
copy a workflow file — uses: KayhanB21/riskratchet@v0.2.13 is the
canonical reference. The action installs riskratchet via uv tool install , runs check ( --format pr-comment in both baseline and
no-baseline modes), upserts a sticky PR comment, and surfaces the
check exit status so PR checks reflect regressions.
# .github/workflows/riskratchet.yml
on : [pull_request]
jobs :
riskratchet :
runs-on : ubuntu-latest
permissions :
contents : read
pull-requests : write
steps :
- uses : actions/checkout@11bd71901bbe5b1630ceea73d27597364c9af683 # v4.2.2
- uses : KayhanB21/riskratchet@v0.2.13
with :
coverage : coverage.json
Inputs (defaults in parentheses): paths ( [tool.riskratchet] paths ), coverage (auto-detected), baseline ( .riskratchet.json
— when the file is missing, the action runs in --fail-above mode),
fail-above ( 60 ), comment ( true ), python-version ( 3.12 ),
riskratchet-version (latest from PyPI), github-token
( ${{ github.token }} ).
For Marketplace discovery, the KayhanB21/riskratchet-action
wrapper repo is the recommended entry point; it delegates to the
root action.yml so both shapes share one source of truth.
Every tagged release ships supply-chain provenance you can inspect: a CycloneDX
SBOM of the wheel's runtime dependency closure (the sbom workflow artifact), a
signed GitHub build-provenance attestation on the wheel and sdist, and PEP 740
PyPI attestations from Trusted Publishing. To confirm a downloaded wheel was built
by this repo:
gh attestation verify riskratchet- < version > -py3-none-any.whl --owner KayhanB21
See docs/threat-model.md for
what each artifact does and does not vouch for.
The canonical use case: AI agent + side project
You've been vibe-coding a FastAPI backend with an AI agent for eight months.
It works, tests are green-ish (62% coverage), but you just noticed
services/billing.py::reconcile_subscriptions quietly grew to 180 lines and
an 11-way match statement you don't remember writing.
pip install riskratchet
pytest --cov --cov-branch --cov-report=json:coverage.json
riskratchet scan src --coverage coverage.json --top 10
reconcile_subscriptions shows up at score 71 (high) with
structural_complexity: 90 , sprawl: 55 , coverage_gap: 60 . You also spot a
surprise: a 12-line public utility _normalize_plan_id scoring 48 because it
has zero tests. Snapshot the bar:
riskratchet baseline src --coverage coverage.json --output .riskratchet.json
git add .riskratchet.json && git commit -m " Add riskratchet baseline "
From here, every time the agent adds a webhook handler or "refactors the
billing flow," run riskratchet check before committing. If it quietly bloated
reconcile_subscriptions from 180 to 220 lines, the check exits 1 and names
the regression. You stop having to remember to look.
Why this is the canonical use case: AI agents are excellent at adding
code, mediocre at noticing they've made things worse. The baseline is your
memory.
Team gating PRs in CI. Run pytest --cov and riskratchet check --format pr-comment
in GitHub Actions; pipe to gh pr comment . The PR-comment format starts with
<!-- riskratchet-report --> so the bot updates the same comment on each push
instead of spamming. The ratchet is mechanical and unowned, so nobody has to
play "the complexity cop" in code review. See
Using riskratchet from an AI coding agent .
Pre-commit hook for a solo repo. Wire pytest-cov and riskratchet
into .pre-commit-config.yaml so every commit regenerates coverage and
gates the commit on no regressions. See
Pre-commit integration .
Investigating one ugly function. Use riskratchet explain path/to/file.py::qualname to dump the six component scores and find the
driver (complexity vs. coverage vs. sprawl). After refactoring, run
riskratchet diff --json | jq '.improved[], .regressed[]' to prove the
change was net-positive, not just rearranging deck chairs.
Why CRAP alone is useful but incomplete
The classic CRAP score ( CC^2 * (1 - line_coverage)^3 + CC ) catches one
shape of bad code: complex and poorly tested. That's a real problem, but
it misses several others that ship to production just as often:
A function with low complexity and zero tests. CRAP gives it CC (a single
digit). Risk is real but invisible.
A function with full line coverage but every branch covered the same way.
CRAP only looks at line coverage.
A function in a 2,000-line module everyone is afraid to touch. Sprawl is
invisible to CRAP.
A function that changed in 40 of the last 90 commits. Churn is invisible
to CRAP.
riskratchet keeps CRAP as a reported metric and computes its own composite
score from six weighted components so those other risks show up too.
How pre-commit and riskratchet fit together
Two things about pre-commit matter for riskratchet:
Pre-commit hides your unstaged edits before running hooks. Hooks only
see the code you're actually about to commit. Useful in general, but it
means riskratchet sees a different source tree than the one open in your
editor.
Each language: python hook runs in its own isolated virtualenv that
contains riskratchet and its declared deps, not your project's pytest,
application code, or test plugins.
Together these create one requirement: the coverage.json riskratchet reads
must reflect the same stashed source tree it's analyzing. Reusing an old
coverage.json from before pre-commit stashed your edits drifts source and
coverage out of sync.
That's why the published hook ships with --no-auto-cov --allow-missing-coverage by default: safe but limited. Pick one of the
patterns below to make it useful.
Pattern A: pre-generate coverage in a sibling hook (recommended)
Run pytest --cov inside the same pre-commit chain so the coverage matches
the stashed tree exactly.
repos :
- repo : local
hooks :
- id : pytest-cov
name : pytest --cov (produces coverage.json for riskratchet)
entry :

[truncated]
