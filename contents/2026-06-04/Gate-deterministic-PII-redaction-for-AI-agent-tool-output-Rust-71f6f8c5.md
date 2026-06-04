---
source: "https://github.com/GaaraZhu/gate"
hn_url: "https://news.ycombinator.com/item?id=48395321"
title: "Gate – deterministic PII redaction for AI agent tool output (Rust)"
article_title: "GitHub - GaaraZhu/gate: A deterministic privacy boundary between your data and AI. · GitHub"
author: "gzhuuu"
captured_at: "2026-06-04T07:45:08Z"
capture_tool: "hn-digest"
hn_id: 48395321
score: 1
comments: 0
posted_at: "2026-06-04T07:28:09Z"
tags:
  - hacker-news
  - translated
---

# Gate – deterministic PII redaction for AI agent tool output (Rust)

- HN: [48395321](https://news.ycombinator.com/item?id=48395321)
- Source: [github.com](https://github.com/GaaraZhu/gate)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T07:28:09Z

## Translation

タイトル: Gate – AI エージェント ツール出力の決定論的な PII 編集 (Rust)
記事のタイトル: GitHub - GaaraZhu/gate: データと AI の間の決定的なプライバシー境界。 · GitHub
説明: データと AI の間の決定的なプライバシー境界。 - 我愛羅珠/ゲート

記事本文:
GitHub - GaaraZhu/gate: データと AI の間の決定的なプライバシー境界。 · GitHub
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
我羅珠
/
ゲート
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチ Tags Go

ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
301 コミット 301 コミット .github/ workflows .github/ workflows アセット アセット クレート クレート dev dev docs docs scripts scripts .gitignore .gitignore .tool-versions .tool-versions CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DISCLAIMER.md DISCLAIMER.md ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md THREAT-MODEL.md THREAT-MODEL.mdクリフ.toml クリフ.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
データと AI の間の決定的なプライバシー境界。
モデルがクエリ結果を認識する前にインターセプトします。ルールに基づいて再現可能で、監査の準備ができています。
AI エージェントは、CLI ツール、スクリプト、MCP サーバーを通じて内部データベースや API にアクセスすることが増えています。保護策がなければ、電子メール、電話番号、納税者識別子、支払い詳細などの機密データが意図せず LLM コンテキスト ウィンドウに公開される可能性があります。
ゲートはクエリ結果がモデルに到達する前にインターセプトし、既存のエージェントのワークフローやプロンプトを変更することなく、検出された PII フィールドを自動的に編集します。これは、エージェントが使用する両方のアクセス パス、Bash コマンド (ハーネス フック経由) と MCP サーバー呼び出し (ラップ スタイルの stdio プロキシ経由) をカバーしており、追加されるオーバーヘッドはクエリごとに 10 ミリ秒未満です。
AI エージェント用のほとんどの PII ガードレールはそれ自体 LLM です。データが機密かどうかを判断するためにモデルに送信されます。ゲートは逆のアプローチをとります。
トレードオフ ゲートにより、ルールは構造化されていないフリーテキストの散文では PII をキャッチできません。脅威モデルは、ゲートがカバーしないものを文書化します。
なぜソースをマスクしないのでしょうか?
ソースを制御する場合は、データベース レベルのマスキングが正しい答えです。ゲートはギャップを埋めない場合は埋め、マスキングが到達できないパスをカバーします。
の

これらは補完的です。DDM が設定されている場合、ゲートは DDM がミスするパスとパターンに対するセーフティ ネットになります。
デモでは次の 3 つのステップを実行します。
ゲート スキャン: クエリの実行前にスキーマ全体の PII 列を検出します。
エージェントがゲートを無効にしてトランザクション テーブルをクエリする - Card_number が完全に表示される
ゲートを有効にした同じクエリ - MCP パスと Bash パスの両方でカード番号が編集されました
OpenCode、Cursor、GitHub Copilot CLI、Codex CLI、および Gemini CLI でも動作します。完全な互換性マトリックスについては、「サポートされる AI ツール」を参照してください。
設計の理論的根拠、脅威モデルのウォークスルー、および検出パイプラインの詳細については、「ゲートの紹介」を参照してください。
フックをインストールする前に、ゲート スキャンを使用して、スキーマが公開する PII の量を評価します。 TABLE_NAME、COLUMN_NAME クエリをパイプしてそこにパイプし、すべてのテーブルにわたるリスク レポートをゲート出力します。ゲート スキャン自体に構成は必要ありません。まだ構成を作成していない場合は、最初にgate config --init-onlyを実行します。
psql -U < ユーザー > -h < ホスト > -d < データベース名 > -c " SELECT TABLE_NAME, COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = 'public' ORDER BY TABLE_NAME, ORDINAL_POSITION " | psql -U < ユーザー > -h < ホスト > -d < データベース名 > -c "ゲートスキャン
MySQL、MS SQL Server (ネイティブ sqlcmd を含む)、Databricks、およびツールキットで管理されるクライアントに対するクエリについては、 docs/scan.md を参照してください。
リスク レベルはカテゴリの感度によって重み付けされます。1 つの SSN 列は 20 のアドレス列よりも重要です。 PII 列が見つかった場合はコード 1 で終了します (CI でスクリプト可能)。検出されたすべての列を表示するには --verbose を渡し、機械可読出力には --json を渡します。
注: ゲート スキャンは列名のみで PII を検出します。結果が LOW ということは、列名がきれいに見えることを意味しますが、データが安全であることを意味するわけではありません。さらに、ゲート 2 はクエリ時に値を検査し、スキャンでは認識できないフリーテキスト、JSON、およびあいまいな名前の列内の PII を検出します。で

複数行の結果では、列内のいずれかの値が PII パターンに一致する場合、列全体が昇格され、一致する行だけでなくすべての行が編集されます。
誤検知 (製品テーブル内の都市など) の場合は、gate scan --review を実行して対話的に優先順位を付け、ホワイトリストに列を追加します。許可リストに登録された列は、名前ベースと値ベースの両方のリダクションをすべてスキップします。列に PII が含まれていないことが確実な場合にのみ、列を許可リストに追加してください。信頼性の低いパターン一致 (confidence_threshold 未満) は編集され、 _gate_summary に警告のフラグが立てられます。抑制する列を column_allowlist に追加します。ゲートのallowlist add/remove/listを使用してリストを直接管理します。
# Homebrew — macOS および Linux (推奨)
醸造タップ GaaraZhu/ゲート && 醸造インストール ゲート
# Cargo binstall — 事前に構築されたバイナリをダウンロードします
貨物設置ゲート
# またはリリースページからバイナリを取得します
# https://github.com/GaaraZhu/gate/releases
構成を作成します (エディターで ~/.config/gate/config.yaml が開きます)。
フックをエージェント ハーネスに登録します。
# クロードコード (デフォルト)
ゲートの初期化
# オープンコード
ゲート init --harness オープンコード
# カーソル
ゲート初期化 --ハーネス カーソル
# GitHub Copilot CLI (プロジェクト スコープ、リポジトリ ルートから実行)
ゲート init --harness copilot-cli
# コーデックス CLI
ゲート初期化 --ハーネス コーデックス
# ジェミニ CLI
ゲート初期化 --harness gemini
プロジェクトのみのセットアップ用に --scope プロジェクトを追加します。ゲート初期化後に OpenCode、Cursor、または Gemini CLI セッションを再起動してフックをロードします。 Codex CLI の場合は、セッションを再起動し、信頼とアクセス許可 UI でフックを確認し、信頼済みとしてマークして有効にします。 Copilot CLI の場合、生成された .github/hooks/PreToolUse.json はデフォルトで gitignor されます。各開発者はローカル クローンでgate init --harness copilot-cli を 1 回実行します。
(オプション) MCP サーバー プロキシを登録して、ツール/通話応答も通過できるようにします

ゲートを通過:
# Claude Code (デフォルト) — 予行演習、何が変更されるかを示します
ゲート初期化 --wrap-mcp
# オープンコード
ゲート init --harness opencode --wrap-mcp --yes
# カーソル
ゲート init --harness カーソル --wrap-mcp --yes
# コパイロット CLI
ゲート init --harness copilot-cli --wrap-mcp --yes
# コーデックス CLI
ゲート init --harness codex --wrap-mcp --yes
# ジェミニ CLI
ゲート init --harness gemini --wrap-mcp --yes
プロジェクトレベルの MCP 構成に --scope プロジェクトを追加します。カーソル プロジェクト スコープの MCP の場合は、登録後に [設定] → [ツールと MCP] でサーバーを再度有効にします。 --servers 、ハーネスごとのパス、および単一サーバーの手動登録については、 docs/mcp.md を参照してください。
AI セッションを開始します。ゲートはクエリ コマンドを自動的にインターセプトします。プロンプトやツールを変更する必要はありません。
最初のセッションの前にゲート検証を実行して、構成が有効であることを確認します。
ゲートは、エージェントがデータにアクセスするために使用する 2 つのアクセス パスをカバーします。ブログ投稿には完全なチュートリアルが記載されています。短いバージョン:
すべての Bash コマンドは、最初にゲート フックを通過します。構成されたツールに一致するコマンドは、ゲート実行 -- <元のコマンド> にサイレントに書き換えられます。これにより、サブプロセスが生成され、2 ゲート検出パイプラインを介して stdout がパイプされます。書き換えはハーネスのツール実行前フックで行われ、Claude Code、OpenCode、Cursor、GitHub Copilot CLI、Codex CLI、および Gemini CLI で強制されます。エージェントはそれをバイパスできません。ハーネスの外部で実行される人間と CI スクリプトは影響を受けません。
AI は次の実行を要求します: tkpsql query --sql "SELECT * FROM users"
│
ハーネス フックの起動 (PreToolUse / tool.execute.before)
│
ゲートフックは次のように書き換えられます:gate run -- tkpsql query --sql "..."
│
┌─────────┴─────────┐
│ ゲート 1: SQL インスペクション │ SELECT * → 列ヒントなし、ゲート 2 に従う
│ ゲート 2: バリュー スキャンニ

ng │ 正規表現 + 列名のヒューリスティック + Luhn チェック
━─────────┬─────────┘
│
{"id": 1、"full_name": "[PII:name]"、"email": "[PII:email]"、...、"_gate_summary": {...}}
MCPパス
Gate mcp は、ハーネスに MCP サーバーとして登録されている透過的な stdio プロキシです。ツール/呼び出し応答を除くすべての JSON-RPC トラフィックはそのまま転送され、モデルに到達する前にゲート 2 を通過します。上流サーバーを変更する必要はありません。
注: tools/call 応答のみが編集されます。resources/read、prompts/get、およびその他の MCP メッセージ タイプは検査なしで転送されます。
AI ──tools/call──> ゲート mcp ──forward──> 上流 MCP サーバー
│
│ <── ツール/PII による通話応答
│
│ ゲート 2 スキャン + 墨消し
│
AI <──編集結果─┘
出力フォーマット
編集された出力では、元の JSON 構造が保持されます。 PII 値は [PII:<type>] プレースホルダーに置き換えられます。編集内容を報告する _gate_summary フィールドが追加されます。
{
"rows" : [{ "id" : 1 , "email" : " [PII:email] " , "ssn" : " [PII:ssn] " }],
「カウント」: 1 、
"_gate_summary" : { "redacted" : 2 、 "types" : [ " email " 、 " ssn " ]、 "warnings" : []}
}
構成で hash_values: true を指定すると、各プレースホルダーは元の値 ( [PII:email:7f83b165] ) から派生した 8 文字の 16 進サフィックスを取得します。同じ生の値は常に同じサフィックスを生成するため、AI は基になるデータを確認することなく、行間で結合または重複排除を行うことができます。基礎となるツールからのエラー応答は変更されずに渡されます。
_gate_summary は 1 つの応答を報告します。すべてのゲートのレトロな集計 - 確認されたクエリの合計、編集された PII フィールド、ヒット率に加えて、ツールおよび PII カテゴリごとの内訳。定期的な監査や境界が実際に行われていることを確認するのに役立ちます

仕事。
クエリによって信頼性の低いリダクションが生成された場合、ゲート レトロは、一意の警告列をリストする低信頼性リダクション セクションと、それを抑制するための正確なゲート ホワイトリスト add <col> コマンドを表示します。列がホワイトリストに追加されると、その列はこのセクションから自動的に削除されます。
統計はデフォルトで収集され、ディスク上のローカル JSONL ログに書き込まれます。統計はマシンから離れることはありません。設定の stats.enabled: false で無効にします。
どのゲートが防御できないのか
ゲートはサンドボックスではなく、決定論的なリダクション レイヤーです。エージェントは非敵対的であると想定し、config の tools: の下にリストされているコマンドからの出力のみを検査します。以下は意図的に範囲外です。
敵対的なエージェント/即時注入。 Gate の脅威モデルは、PII を不用意に漏洩するエージェントです。 Gate Protect (Unix) は、設定の所有権を root に転送することで、最も直接的なバイパス (設定編集によってゲートを無効にするハイジャックされたエージェント) をブロックします。しかし、決意の強い攻撃者は、ツールにないコマンドを呼び出したり、非 JSON 出力形式を要求したり、エンコーダーを介してパイプしたり、次のセッションのハーネス設定ファイルからフック エントリを削除したりすることで、ゲートを迂回することができます。ゲートをハーネス レベルの Bash ホワイトリストとペアにして、残りのギャップを埋めます。
ツールにないコマンド: 。 AI はそれらを自由に呼び出すことができます。彼らの出力は決して検査されません。
非 JSON ツールの出力。プレーン テキスト、CSV、およびその他の形式は変更されずに通過します。ツールを構成する

[切り捨てられた]

## Original Extract

A deterministic privacy boundary between your data and AI. - GaaraZhu/gate

GitHub - GaaraZhu/gate: A deterministic privacy boundary between your data and AI. · GitHub
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
GaaraZhu
/
gate
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
301 Commits 301 Commits .github/ workflows .github/ workflows assets assets crates crates dev dev docs docs scripts scripts .gitignore .gitignore .tool-versions .tool-versions CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DISCLAIMER.md DISCLAIMER.md LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md THREAT-MODEL.md THREAT-MODEL.md cliff.toml cliff.toml View all files Repository files navigation
A deterministic privacy boundary between your data and AI.
Intercepts query results before the model sees them — rule-driven, reproducible, and audit-ready.
AI agents increasingly access internal databases and APIs through CLI tools, scripts, and MCP servers. Without safeguards, sensitive data such as emails, phone numbers, tax identifiers, and payment details can be unintentionally exposed to LLM context windows.
gate intercepts query results before they reach the model and automatically redacts detected PII fields without requiring changes to existing agent workflows or prompts. It covers both access paths agents use: Bash commands (via a harness hook) and MCP server calls (via a wrap-style stdio proxy), adding < 10 ms of overhead per query.
Most PII guardrails for AI agents are themselves LLMs — they send your data to a model to decide whether it's sensitive. Gate takes the opposite approach.
The trade-off gate makes: rules can't catch PII in unstructured free-text prose. The threat model documents what gate doesn't cover.
Why not just mask at the source?
Database-level masking is the right answer when you control the source. Gate fills the gap when you don't, and covers the paths masking can't reach.
They're complementary: if you have DDM configured, gate is the safety net for the paths and patterns DDM misses.
The demo walks through three steps:
gate scan detecting PII columns across the schema before any query runs
An agent querying the transactions table with gate disabled — card_number fully visible
The same queries with gate enabled — card_number redacted across both MCP and Bash paths
Also works with OpenCode, Cursor, GitHub Copilot CLI, Codex CLI, and Gemini CLI — see Supported AI Tools for the full compatibility matrix.
For the design rationale, threat-model walkthrough, and detection-pipeline deep dive, read Introducing gate .
Before installing the hook, use gate scan to assess how much PII your schema exposes. Pipe a TABLE_NAME, COLUMN_NAME query into it and gate prints a risk report across every table. No config is required for gate scan itself — if you haven't created one yet, run gate config --init-only first.
psql -U < user > -h < host > -d < dbname > -c " SELECT TABLE_NAME, COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = 'public' ORDER BY TABLE_NAME, ORDINAL_POSITION " | gate scan
See docs/scan.md for queries against MySQL, MS SQL Server (including native sqlcmd ), Databricks, and toolkit-managed clients.
Risk level is weighted by category sensitivity — one SSN column matters more than twenty address columns. Exits with code 1 if any PII columns are found (scriptable in CI). Pass --verbose to show all detected columns, or --json for machine-readable output.
Note: gate scan detects PII by column name only. A LOW result means your column names look clean — it does not mean the data is safe. Gate 2 additionally inspects values at query time, catching PII in free-text, JSON, and ambiguously-named columns that scan cannot see. In multi-row results, if any value in a column matches a PII pattern, the entire column is promoted and all rows are redacted — not just the matching row.
For false positives (e.g. city in a products table), run gate scan --review to triage interactively and add columns to the allowlist. Allowlisted columns skip all redaction — both name-based and value-based. Only add a column to the allowlist when you are certain it contains no PII. Low-confidence pattern matches (below confidence_threshold ) are redacted and flagged with a warning in _gate_summary ; add the column to column_allowlist to suppress. Manage the list directly with gate allowlist add/remove/list .
# Homebrew — macOS and Linux (recommended)
brew tap GaaraZhu/gate && brew install gate
# cargo binstall — downloads a prebuilt binary
cargo binstall gate
# Or grab a binary from the releases page
# https://github.com/GaaraZhu/gate/releases
Create your config (opens ~/.config/gate/config.yaml in your editor):
Register the hook with your agent harness:
# Claude Code (default)
gate init
# OpenCode
gate init --harness opencode
# Cursor
gate init --harness cursor
# GitHub Copilot CLI (project-scoped, run from repo root)
gate init --harness copilot-cli
# Codex CLI
gate init --harness codex
# Gemini CLI
gate init --harness gemini
Add --scope project for project-only setup. Restart your OpenCode, Cursor, or Gemini CLI session after gate init to load the hook. For Codex CLI, restart the session, then review the hook in the Trust & Permissions UI, mark it as trusted, and enable it. For Copilot CLI, the generated .github/hooks/PreToolUse.json is gitignored by default — each developer runs gate init --harness copilot-cli once in their local clone.
(Optional) Register MCP server proxies so tools/call responses also pass through gate:
# Claude Code (default) — dry-run, shows what would change
gate init --wrap-mcp
# OpenCode
gate init --harness opencode --wrap-mcp --yes
# Cursor
gate init --harness cursor --wrap-mcp --yes
# Copilot CLI
gate init --harness copilot-cli --wrap-mcp --yes
# Codex CLI
gate init --harness codex --wrap-mcp --yes
# Gemini CLI
gate init --harness gemini --wrap-mcp --yes
Add --scope project for project-level MCP config. For Cursor project-scoped MCP, re-enable the servers in Settings → Tools & MCPs after registration. See docs/mcp.md for --servers , per-harness paths, and manual single-server registration.
Start your AI session — gate intercepts query commands automatically. No changes to your prompts or tools required.
Run gate validate to confirm your config is valid before the first session.
gate covers two access paths agents use to reach data. The blog post has the full walkthrough; the short version:
Every Bash command passes through gate hook first. Commands that match a configured tool are silently rewritten to gate run -- <original command> , which spawns the subprocess and pipes stdout through the two-gate detection pipeline. The rewrite happens in the harness's pre-tool-execution hook — it is enforcing in Claude Code, OpenCode, Cursor, GitHub Copilot CLI, Codex CLI, and Gemini CLI; the agent cannot bypass it. Humans and CI scripts running outside the harness are untouched.
AI asks to run: tkpsql query --sql "SELECT * FROM users"
│
harness hook fires (PreToolUse / tool.execute.before)
│
gate hook rewrites to: gate run -- tkpsql query --sql "..."
│
┌──────────────┴──────────────┐
│ Gate 1: SQL inspection │ SELECT * → no column hints, defer to Gate 2
│ Gate 2: Value scanning │ regex + column-name heuristics + Luhn check
└──────────────┬──────────────┘
│
{"id": 1, "full_name": "[PII:name]", "email": "[PII:email]", ..., "_gate_summary": {...}}
MCP path
gate mcp is a transparent stdio proxy registered in the harness as the MCP server. It forwards all JSON-RPC traffic verbatim except tools/call responses, which pass through Gate 2 before reaching the model. No changes to the upstream server are required.
Note: only tools/call responses are redacted — resources/read , prompts/get , and other MCP message types are forwarded without inspection.
AI ──tools/call──> gate mcp ──forward──> upstream MCP server
│
│ <── tools/call response with PII
│
│ Gate 2 scan + redact
│
AI <───redacted result─┘
Output format
Redacted output preserves the original JSON structure. PII values are replaced with [PII:<type>] placeholders. A _gate_summary field is appended reporting what was redacted.
{
"rows" : [{ "id" : 1 , "email" : " [PII:email] " , "ssn" : " [PII:ssn] " }],
"count" : 1 ,
"_gate_summary" : { "redacted" : 2 , "types" : [ " email " , " ssn " ], "warnings" : []}
}
With hash_values: true in config, each placeholder gains an 8-char hex suffix derived from the original value ( [PII:email:7f83b165] ). The same raw value always produces the same suffix, so the AI can join or deduplicate across rows without ever seeing the underlying data. Error responses from the underlying tool pass through unchanged.
_gate_summary reports a single response. gate retro aggregates across all of them — total queries seen, PII fields redacted, hit rate, plus a breakdown by tool and PII category. Useful for periodic audits and for confirming the boundary is doing real work.
If any query produced a low-confidence redaction, gate retro surfaces a Low-confidence redactions section listing each unique warned column and the exact gate allowlist add <col> command to suppress it. Once a column is added to the allowlist it disappears from this section automatically.
Stats are collected by default and written to a local JSONL log on disk — they never leave your machine. Disable with stats.enabled: false in config.
What gate does NOT protect against
gate is a deterministic redaction layer, not a sandbox. It assumes the agent is non-adversarial and only inspects output from commands listed under tools: in config. The following are deliberately out of scope:
Adversarial agents / prompt injection. Gate's threat model is an agent that inadvertently exfiltrates PII. gate protect (Unix) blocks the most direct bypass — a hijacked agent disabling gate via config edits — by transferring config ownership to root. But a determined attacker can still route around gate by invoking commands not in tools: , requesting non-JSON output formats, piping through encoders, or removing the hook entry from the harness settings file for the next session. Pair gate with a harness-level Bash allowlist to close the residual gap.
Commands not in tools: . The AI can invoke them freely; their output is never inspected.
Non-JSON tool output. Plain text, CSV, and other formats pass through unchanged. Configure tools

[truncated]
