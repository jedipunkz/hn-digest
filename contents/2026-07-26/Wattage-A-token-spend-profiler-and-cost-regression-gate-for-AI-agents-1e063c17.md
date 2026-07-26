---
source: "https://github.com/faizannraza/wattage"
hn_url: "https://news.ycombinator.com/item?id=49063397"
title: "Wattage: A token-spend profiler and cost-regression gate for AI agents"
article_title: "GitHub - faizannraza/wattage: A token-spend profiler and cost-regression gate for AI agents. · GitHub"
author: "faizanraza03"
captured_at: "2026-07-26T23:53:23Z"
capture_tool: "hn-digest"
hn_id: 49063397
score: 2
comments: 0
posted_at: "2026-07-26T23:27:35Z"
tags:
  - hacker-news
  - translated
---

# Wattage: A token-spend profiler and cost-regression gate for AI agents

- HN: [49063397](https://news.ycombinator.com/item?id=49063397)
- Source: [github.com](https://github.com/faizannraza/wattage)
- Score: 2
- Comments: 0
- Posted: 2026-07-26T23:27:35Z

## Translation

タイトル: ワット数: AI エージェントのトークン支出プロファイラーとコスト回帰ゲート
記事のタイトル: GitHub - faizannraza/wattage: AI エージェント用のトークン支出プロファイラーおよびコスト回帰ゲート。 · GitHub
説明: AI エージェント用のトークン支出プロファイラーおよびコスト回帰ゲート。 - ファイザンラザ/ワット数

記事本文:
GitHub - faizannraza/wattage: AI エージェント用のトークン支出プロファイラーおよびコスト回帰ゲート。 · GitHub
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
ファイザンラザ
/
ワット数
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
48 コミット 48 コミット .github .github アクション アクション ベンチマーク ベンチマーク ドキュメント ドキュメント サンプル サンプル npm npm src/ ワット数 src/ ワット数 テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント向けの Kill-A-Watt メーター。痕跡を指すと教えてくれます
トークンがどこで燃やされ、無駄にされているかを正確に把握し、各無駄の価格を確認します
パターンを実際のドルで表示し、修正を処方し、変更すると CI が失敗する可能性があります
エージェントの費用が目に見えて高くなります。
実際に捕らえられたエージェントの痕跡 (来歴を参照) —
ワット数は、キャッシュされずに再送信される安定したプロンプト プレフィックスをキャッチします。
老廃物を取り除き、修正を処方します。このGIFを再生成するには
vhs docs/assets/demo.tape (正確なコマンドについてはテープ ファイルを参照してください)。
uvxワット数レポートtrace.json
構成ファイルも API キーも必要なく、完全にオフラインです。OTLP JSON を指定します。
エクスポートをトレースすると、すべての呼び出しに価格が設定され、すべての検出器が実行されます。持っていない
トレースはまだですか？最初のトレースを取得すると、両方がカバーされます
「すでに OTel トレースがあります」および「インストルメンテーションはありません」 (実行可能ファイル、
何もない状態から実際の有料レポートを作成するまでの 5 分間の道のり)。または、今すぐ試してください
このリポジトリに同梱されているフィクスチャに対して:
git clone https://github.com/faizannraza/wattage
CD ワット数 && UV 同期
UV 実行ワット数レポートの例/sample_trace.json
╭──── ⚡ ワット数 — example/sample_trace.json ────╮
│ トークン効率: A (100) 総コスト: $0.0602 │
│ 品質：未測定 │
╰─────

─────────────────╯
トークンの内訳
┏━━━━━━━━━━━━━━━━┳━━━━━━━━┓
│ カテゴリ │ トークン │
┡━━━━━━━━━━━━━━━━╇━━━━━━━━┩
│ 入力 │ 18450 │
│ 出力 │ 320 │
│ キャッシュ読み取り │ 0 │
│ キャッシュ作成 │ 0 │
│ 推理 │ 0 │
━━━━━━━━━━━━━━━┘
検出結果はありません - このトレースは効率的であるように見えます。
価格: 2026-07-18-検証済み
または、ターミナルの代わりに自己完結型の共有可能な HTML フレーム グラフを取得します。
表示:
UV 実行ワット数レポートの例/sample_trace.json --html report.html
マーケティング上の主張ではなく証拠
Wattage の傑出した機能は、収束エンジンです。
非収束検出器。ループを通過するエージェントを捕捉します。
単純な完全一致パターンなど、実際の進歩はありません
重複検出器は構造的に認識できません (新しいタイムスタンプで再試行してください)
毎回、2 つの戦略の間で揺れが生じ、「生産性が高いように見える」
すべての呼び出しが技術的には一意であるものの、実際には何も変わらないストール
学んだ）。
そう断言するのではなく、手作業でレビューした 10 個のラベルを付けたセットを作成しました。
合成ループと実際のワット数分類子に対するベンチマーク
SHA-256 完全一致ベースライン実装:
自分で再現してみましょう — 厳選したり、隠し設定をしたりする必要はありません。
uv run python -m benchmarks.harness
そして、本物の捕獲されたエージェントの追跡（合成ではありません - を参照）
来歴についてはベンチマーク/トレース/README.md)、
Wattage の prefix_churn 修正シミュレーションでは 44.7% のコスト削減が示されています
( $0.000199 → $0.000110 ) 安定したプレフィックスでのプロンプト キャッシュの有効化による —
それは 3- であるため、小さな金額の数字になります。

ターンデモトレースですが、その仕組みは
生産規模でも同一です。自分のトレースに対して実行して数値を取得する
それは重要です:
uv run python -c " from benchmarks.frontier import build_frontier; print(build_frontier()) "
完全な方法論: Convergence Engine 。
uv 実行ワット数バッジ trace.json --out wattage-badge.svg
![ワット数] ( wattage-badge.svg )
--badge-out を CI ジョブに接続し (以下を参照)、毎回再生成されるようにします。
デフォルトのブランチにマージすると、README 内のバッジは有効なままになります。
3 つのサーフェス、その下に 1 つの正規化されたデータ モデル
(セッション → タスク → ループ → 反復 → 呼び出し)、以下から構築されています。
OpenTelemetry GenAI のセマンティック規約
トレース:
ワット数レポート — トレースを取り込み、すべてのコールに対して価格を設定します。
ベンダーの日付付き価格スナップショットを作成し、8 つの検出器を実行します。
すべての発見事項は実際のドルで価格設定されており、具体的な修正が含まれており、
品質リスク層 (なし / 低 / レビュー) のタグが付けられています - 修正
出力品質が変更される可能性があります (モデルのダウングレード、推論の減少)
--quality マップが実際のマップで裏付けられた場合のみ、スコアにカウントされます。
証拠。詳細: 検出器。
ワット数スコア / ワット数バッジ — 単一の 0 ～ 100 トークン効率
README バッジまたは CI ゲートのグレード。
ワット数 ci — コスト回帰ゲート (下記)。
ワット数は数値を捏造することはありません。価格が設定されていないモデルは通話料金をそのままにします
推測するのではなく、ゼロで（ワット数 ci が大声で失敗し、終了コード 4 で）、
測定されていない品質信号は、正常とはみなされず、 unmeasured として報告されます。
# .github/workflows/wattage.yml
名前 : ワット数
に:
プルリクエスト:
パス: ["エージェント/**"、"プロンプト/**"、"src/**"]
同時実行性:
グループ: ワット数-${{ github.ref }}
キャンセル中 : true
ジョブ:
トークン効率:
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- 名前 : トレース フィクスチャの生成

実行：python scripts/run_agent_fixture.py >trace.json
- 名前 : ワット数コスト回帰ゲート
使用: faizannraza/wattage/action@v0.1.0
付き:
ソース : トレース.json
ベースライン: .wattage/baseline.json
フェイルオン: " スコア_下:80、コスト_デルタ_pct_上:5、任意の_クリティカル: true "
事前コメント: " true "
エージェントがしきい値を超えて後退すると、ビルドは失敗します (終了コード 1)。
設定すると、検出器ごとのデルタ テーブルが PR コメントとして投稿され、SARIF が発行されます。
(GitHub の [セキュリティ] タブに表示されます) および他の CI システムの JUnit XML。
ベースラインはコミットされた小さな JSON ファイルです。ノイズフロア保護は
統計的ではなく構造的: 実際に実行されたときにのみ更新されます。
ゲートを通過した。
これはセットアップの半分にすぎません。 PR ジョブは使い捨てチェックアウトで実行されるため、
ディスク上の .wattage/baseline.json を更新するものではあり得ません。
更新には、デフォルトのブランチへのプッシュ時にトリガーされる 2 番目のワークフローが必要です。
これにより、マージのたびに、更新されたベースライン (およびバッジ) がコミットされます。
これをスキップすると、すべての PR が同じ古いベースラインと比較されることになります。
永遠に。両方のワークフローを含む完全なリファレンス: CI Integration 。
検出器は Python エントリポイント グループを通じて検出されるため、検出器を追加します
このリポジトリのコア パイプラインに触れる必要はありません — を参照してください。
CONTRIBUTING.md (完全な「検出器の書き込み」ウォークスルー)、
参照としてcache_gapを使用する
たとえば。
AI エージェント用のトークン支出プロファイラーとコスト回帰ゲート。
Readme Apache-2.0 ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A token-spend profiler and cost-regression gate for AI agents. - faizannraza/wattage

GitHub - faizannraza/wattage: A token-spend profiler and cost-regression gate for AI agents. · GitHub
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
faizannraza
/
wattage
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
48 Commits 48 Commits .github .github action action benchmarks benchmarks docs docs examples examples npm npm src/ wattage src/ wattage tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A Kill-A-Watt meter for your AI agents. Point it at a trace and it tells
you exactly where your tokens are being burned and wasted, prices each waste
pattern in real dollars, prescribes a fix, and can fail your CI when a change
makes your agent measurably more expensive.
A real captured agent trace (see provenance ) —
Wattage catches a stable prompt prefix being re-sent instead of cached, prices
the waste, and prescribes the fix. Regenerate this GIF with
vhs docs/assets/demo.tape (see the tape file for the exact command).
uvx wattage report trace.json
No config file, no API key, fully offline — point it at an OTLP JSON
trace export and it prices every call and runs every detector. Don't have a
trace yet? Getting your first trace covers both
"I already have OTel traces" and "I have zero instrumentation" (a runnable,
5-minute path from nothing to a real, priced report). Or try it right now
against the fixture shipped in this repo:
git clone https://github.com/faizannraza/wattage
cd wattage && uv sync
uv run wattage report examples/sample_trace.json
╭──── ⚡ wattage — examples/sample_trace.json ────╮
│ Token Efficiency: A (100) Total cost: $0.0602 │
│ quality: unmeasured │
╰─────────────────────────────────────────────────╯
Token breakdown
┏━━━━━━━━━━━━━━━━┳━━━━━━━━┓
┃ Category ┃ Tokens ┃
┡━━━━━━━━━━━━━━━━╇━━━━━━━━┩
│ input │ 18450 │
│ output │ 320 │
│ cache_read │ 0 │
│ cache_creation │ 0 │
│ reasoning │ 0 │
└────────────────┴────────┘
No findings — this trace looks efficient.
pricing: 2026-07-18-verified
Or get a self-contained, shareable HTML flame graph instead of the terminal
view:
uv run wattage report examples/sample_trace.json --html report.html
The evidence, not a marketing claim
Wattage's standout feature is the convergence engine — the
nonconvergence detector, which catches an agent thrashing through a loop
without making real progress, including patterns a naive exact-match
duplicate detector structurally cannot see (a retry with a fresh timestamp
each time, an oscillation between two strategies, a "productive-looking"
stall where every call is technically unique but nothing is actually
learned).
Rather than assert that, we built a hand-reviewed set of 10 labeled
synthetic loops and benchmarked Wattage's classifier against a real
SHA-256 exact-match baseline implementation:
Reproduce it yourself — no cherry-picking, no hidden setup:
uv run python -m benchmarks.harness
And on a genuine captured agent trace (not synthetic — see
benchmarks/traces/README.md for provenance),
Wattage's prefix_churn fix simulation shows a 44.7% cost reduction
( $0.000199 → $0.000110 ) from enabling prompt caching on the stable prefix —
small dollar figures because it's a 3-turn demo trace, but the mechanism is
identical at production scale. Run it against your own traces for numbers
that matter:
uv run python -c " from benchmarks.frontier import build_frontier; print(build_frontier()) "
Full methodology: The Convergence Engine .
uv run wattage badge trace.json --out wattage-badge.svg
![ Wattage ] ( wattage-badge.svg )
Wire --badge-out into your CI job (see below) so it regenerates on every
merge to your default branch, and the badge in your README stays live.
Three surfaces, one normalized data model underneath
( sessions → tasks → loops → iterations → calls ), built from
OpenTelemetry GenAI semantic-convention
traces:
wattage report — ingests a trace, prices every call against a
vendored, dated pricing snapshot, and runs eight detectors:
Every finding is priced in real dollars, includes a concrete fix, and is
tagged with a quality_risk tier ( none / low / review ) — a fix that
could plausibly change output quality (a model downgrade, less reasoning)
only counts toward your score once a --quality map backs it with real
evidence. Full detail: Detectors .
wattage score / wattage badge — a single 0–100 Token Efficiency
grade for a README badge or a CI gate.
wattage ci — the cost-regression gate (below).
Wattage never fabricates a number: an unpriced model leaves that call's cost
at zero (and fails wattage ci loudly, exit code 4) rather than guessing;
an unmeasured quality signal is reported as unmeasured , not assumed fine.
# .github/workflows/wattage.yml
name : Wattage
on :
pull_request :
paths : ["agents/**", "prompts/**", "src/**"]
concurrency :
group : wattage-${{ github.ref }}
cancel-in-progress : true
jobs :
token-efficiency :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- name : Generate trace fixture
run : python scripts/run_agent_fixture.py > trace.json
- name : Wattage cost-regression gate
uses : faizannraza/wattage/action@v0.1.0
with :
source : trace.json
baseline : .wattage/baseline.json
fail-on : " score_below:80,cost_delta_pct_above:5,any_critical:true "
pr-comment : " true "
Fails the build (exit code 1) when your agent regresses past the threshold
you set, posts a per-detector delta table as a PR comment, and emits SARIF
(shows up in GitHub's Security tab) and JUnit XML for any other CI system.
The baseline is a small committed JSON file — noise-floor protection is
structural, not statistical: it only ever updates on a run that actually
passed the gate.
This is only half the setup. A PR job runs on a throwaway checkout, so
it can't be the thing that updates .wattage/baseline.json on disk — that
update needs a second workflow, triggered on push to your default branch,
that commits the refreshed baseline (and badge) back after each merge.
Skipping it means every PR compares against the same stale baseline
forever. Full reference, with both workflows: CI Integration .
Detectors are discovered through a Python entry-point group, so adding one
doesn't require touching this repo's core pipeline — see
CONTRIBUTING.md for the full "write a detector" walkthrough,
using cache_gap as the reference
example.
A token-spend profiler and cost-regression gate for AI agents.
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
