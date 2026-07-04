---
source: "https://github.com/shivyadavus/open-kioku"
hn_url: "https://news.ycombinator.com/item?id=48785607"
title: "Show HN: Open Kioku – local evidence layer for AI coding agents"
article_title: "GitHub - shivyadavus/open-kioku: Local-first code intelligence MCP for AI coding agents. · GitHub"
author: "shivyadavus"
captured_at: "2026-07-04T14:40:19Z"
capture_tool: "hn-digest"
hn_id: 48785607
score: 1
comments: 0
posted_at: "2026-07-04T14:19:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open Kioku – local evidence layer for AI coding agents

- HN: [48785607](https://news.ycombinator.com/item?id=48785607)
- Source: [github.com](https://github.com/shivyadavus/open-kioku)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T14:19:34Z

## Translation

タイトル: Show HN: Open Kioku – AI コーディング エージェントのローカル証拠レイヤー
記事タイトル: GitHub - shivyadavus/open-kioku: AI コーディング エージェント向けのローカルファースト コード インテリジェンス MCP。 · GitHub
説明: AI コーディング エージェント向けのローカル ファースト コード インテリジェンス MCP。 - shivyadavus/open-kioku

記事本文:
GitHub - shivyadavus/open-kioku: AI コーディング エージェント向けのローカルファースト コード インテリジェンス MCP。 · GitHub
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
シビャダヴス
/
オープンキオク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインBr

anches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
219 コミット 219 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor-plugin .cursor-plugin .githooks .githooks .github .github Formula Formula アセット アセット アセット ベンチマーク ベンチマーク クレート クレート デモ デモ ドキュメント ドキュメントの例例 パッケージ パッケージ スクリプト スクリプト スキル/ open-kioku スキル/ open-kioku test-repos/ Rust-mini test-repos/ Rust-mini .gitignore .gitignore .mcp.json .mcp.json CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンスライセンス通知通知 PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md STABILITY.md STABILITY.md Audit.toml Audit.toml claude_plugin.json claude_plugin.json デモ.テープ デモ.テープ 拒否出力.txt 拒否出力.txt 拒否.toml Deny.toml Glama.json Glama.json jobs.json jobs.json mcp.json mcp.json release-metadata.json release-metadata.json run.json run.json smithery.yaml smithery.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude、Cursor、Codex、または任意の MCP 互換コーディング エージェントに、デフォルトでローカル インデックスと読み取り専用 MCP ツールを使用して、コードベースを編集する前に証拠レイヤーを提供します。
自分のリポジトリで次の 3 つのコマンドを実行して、どのような変化があるかを確認してください。
npm install -g open-kioku
オーケー、インデックス。
ok mcp installcursor --repo 。 # または: クロード、コーデックス、ジェミニ、ウィンドサーフィン、トレイ、ゼド、オープンコード
印刷された MCP 構成をエージェントに貼り付けて、次のプロンプトを貼り付けます。
編集する前に Open Kioku を使用してください。 repo_status、search_code、get_definition、を確認します。
get_references、impact_analysis、および find_tests_for_change です。で計画を立てる
plan_change を最初に実行し、次に編集を実行し、編集後に verify_change を実行して検証します。
エージェントは、シンボル、参照、影響分析などのインデックス付き証拠を取得しました。

、候補をテストし、単一のファイルに触れる前に境界を編集します。
リポジトリで okprove を実行すると、インデックス付きカウント、タスク スコア、および検証シグナルを含む共有可能なマークダウン レポートまたは HTML レポートが生成されます。このレポートは意図的にソース スニペットを省略しているため、プライベート リポジトリからレビューおよび共有できるように設計されています。
わかりました証明してください。 --task " あなたが取り組んでいる機能 "
わかりました証明してください。 --task " 取り組んでいる機能 " --html
代表的なパブリック リポジトリ監査では、33.1 秒で 4,600 以上のファイル、46,000 以上のシンボル、および 8,900 以上のテストのインデックスがローカルに作成されました。 Open Kioku の正確なバージョン、リポジトリのリビジョン、注意事項、および言語の制限は、 docs/large-repo-proof.md に記録されています。
コードの検索 → 定義の取得 → 影響分析 → 変更のテストの検索 → 変更の計画 → 変更の検証
証拠に裏付けられた編集前計画には、プライマリ ファイル、関連するシンボル、推定される爆発範囲、正確な検証コマンド、信頼スコア、次に実行する MCP 呼び出しが含まれます。
ok plan と MCP plan_change は、ローカル インデックスからの具体的なファイル範囲、検証候補、および信頼性に関する警告を返します。バンドルされたデモ リポジトリでは、現在トークン タスクが表示されています。
src/auth.rs 、 src/lib.rs 、および testing/auth_flow.rs のプライマリ コンテキスト
カーゴテストによる issue_token 、 validate_token 、 login_returns_valid_token などの検証候補
生成、ベンダー、ビルド、および .ok/ アーティファクトを範囲外に保ちながら、関連するソース ファイルとテスト ファイルを許可する編集境界
正確な参照、実行時、履歴、またはカバレッジの証拠が利用できない場合の明示的な警告
AI コーディング エージェントが最も強力なのは、編集前にコードベースに事実を問い合わせることができる場合です。インデックス付きコンテキストがないと、ファイルのクロールを繰り返すとトークンが書き込まれ、テキストの一致から参照が推測され、多くの場合、変更が既に失敗した後にのみテストが選択されます。
オペ

en Kioku はエージェントに事前編集ルーチンを提供します。
インデックス付きのコードとファイルを検索します。
シンボルと参照を解決します。
影響と検証の対象となる可能性が高い、証拠に裏付けられた編集前計画を作成します。
メモリがインデックス付きコードの証拠を上回ることなく、以前のリポジトリの事実を思い出します。
コンテキストをハンドルに圧縮し、後で元のスニペットを取得できるようにします。
これらの機能は、ローカル標準入出力上の MCP を通じて提供されます。
npm install -g open-kioku
オーケー --バージョン
open-kioku npm パッケージは、プラットフォーム固有のオプションの依存関係を通じてネイティブの ok バイナリをインストールする JavaScript ラッパーです。
カーゴ binstall open-kioku-cli
オーケー --バージョン
crates.io
カーゴインストール open-kioku-cli
オーケー --バージョン
GitHub リリース
タグ付きリリースでは、Linux x86_64/arm64 musl、macOS x86_64/arm64、および Windows x86_64 のネイティブ バイナリと SHA-256 チェックサムが公開されます。 https://github.com/shivyadavus/open-kioku/releases からダウンロードします。
git clone https://github.com/shivyadavus/open-kioku.git
CDオープンキオク
カーゴインストール --path crates/open-kioku-cli
オーケー --助けて
安定した Rust ツールチェーンが必要です。
ok init /absolute/path/to/repo
ok インデックス /absolute/path/to/repo
OK ドクター /absolute/path/to/repo
ok status /absolute/path/to/repo --markdown --write ok-status.md
OK 監査 /absolute/path/to/repo をセットアップします
ok --repo /absolute/path/to/repo 検索「気になる機能またはバグ」 --limit 5
ok Index はローカル データを .ok/ に書き込みます。SQLite メタデータとグラフ行は .ok/index.sqlite に、BM25 検索データは .ok/search/tantivy に書き込まれます。大規模なインデックスは、スキャン、解析、発生、保存、グラフ、検索、完了などの進行状況をレポートします。
編集中はインデックスを最新の状態に保ちます。
ok ウォッチ /absolute/path/to/repo
エージェントを接続する
オーケー mcp インストール カーソル --repo /absolute/path/to/repo
ok mcp install claude --repo /absolute/path/to/repo
ok mcp install codex --repo /abs

olute/パス/へ/リポジトリ
ok mcp install gemini --repo /absolute/path/to/repo
ok mcp install Windsurf --repo /absolute/path/to/repo
ok mcp install trae --repo /absolute/path/to/repo
ok mcp install opencode --repo /absolute/path/to/repo
ok mcp install zed --repo /absolute/path/to/repo
印刷された MCP 構成スニペットをエージェントに貼り付けます。デフォルトのサーバーは読み取り専用で、stdio 経由でローカルに実行されます。インストール マトリックスは、Claude、Cursor、Codex、Gemini CLI、OpenCode、Zed、Windsurf、および Trae に対してテストされています。
エージェント固有のセットアップ ガイド: クロード · カーソル · コーデックス · Gemini CLI 。 OpenCode、Zed、Windsurf、および Trae は、上記と同じ検証済みの ok mcp install <client> --repo /absolute/path/to/repo フローを使用します。
Open Kioku には、事前に設定されたリポジトリ スコープのプラグインとマーケットプレイスのマニフェストがバンドルされています。
OpenAI Codex : MCP セットアップが正常であることを確認しました。 mcp install codex --repo /absolute/path/to/repo ; .codex-plugin/ メタデータは、Git スコープのプラグイン マニフェストをサポートする Codex インストールに含まれています。
Claude : .claude-plugin/ メタデータは、検証済みの ok mcp install claude --repo /absolute/path/to/repo MCP config パスと一緒に含まれています。
Cursor : .cursor-plugin/ ルールとスキルは、検証済みの ok mcp installcursor --repo /absolute/path/to/repo MCP config パスと一緒に含まれています。
ゴールデン プロンプトと 1 つのコマンドのスモーク テストを備えたスターター サンプルが利用可能です
Examples/cursor と example/claude で。
ok status --markdown --write ok-status.md は、インデックス数、SCIP 品質、準備状況チェック、および次のステップを含むポータブル ハンドオフを作成します。 ok setup Audit は、インデックス、セキュリティ体制、MCP サーバーの健全性、プラグインのメタデータ、およびサポートされているクライアント インストール マトリックスをチェックします。
フリートまたはマルチサービス ワークスペースの場合は、まず各プロジェクトにインデックスを付けてから、ワークスペース構成を作成します。
[ ワークスペース ]
プロジェクト = [
{ 名前 = " サービス a " 、リポジトリ = "

../サービス-a " },
{ 名前 = " サービス-b " 、リポジトリ = " ../サービス-b " },
】
ソースを再解析せずにクロスプロジェクト リンカーを実行します。
ok インデックス --mode クロスプロジェクト --workspace /absolute/path/to/workspace
OK アーキテクチャ フリート --workspace /absolute/path/to/workspace
インデックススナップショット
チームと CI は、個人のメモリや圧縮されたコンテキスト状態を共有せずに、既知の正常なローカル インデックスを共有できます。
ok --repo /absolute/path/to/repo スナップショットのエクスポート --quality best
ok --repo /absolute/path/to/repo スナップショット ドクター
ok --repo /absolute/path/to/repo スナップショットのインポート
ok --repo /絶対/パス/to/repo インデックス --from-snapshot auto
品質モード
Open Kioku は外部言語インデクサーなしで動作しますが、正確な参照により、検索の基礎、影響分析、テストの選択、および計画が向上します。利用可能なものを確認してください:
ok setup Audit /absolute/path/to/repo --markdown --write ok-setup.md
オーケー scip ドクター /absolute/path/to/repo
ok scip setup /absolute/path/to/repo
ok インデックス /absolute/path/to/repo --with-scip auto
OK ドクター /absolute/path/to/repo
デフォルトのインデックス作成では、index.scip や .ok/indexes/*.scip などの既存の SCIP ファイルが存在する場合、それを使用します。 --with-scip は、サポートされているリポジトリに対してインストールされているインデクサーを自動実行します。サードパーティのツールはインストールされません。 --with-scip required は、SCIP ファクトをインポートできない場合、インデックスに失敗します。
SCIP は主要な精度プロバイダーです。デフォルトの品質モデルはローカルで無料のままです。インデックス付きシンボル/チャンク/インポート、言語固有の静的ファクト、インデックス付きテスト、ビルドシステム検出、index.scip が使用可能な場合の SCIP の正確な参照。 Java/Gradle リポジトリは、テスト ファイルのパスがわかっている場合、一般的なリポジトリ全体のテスト コマンドではなく、./gradlew :x-pack:plugin:ml:test --tests org.example.SomeTests などのスコープ指定された検証コマンドを取得します。
言語固有の静的分析により、

インポート、Java の継承と実装されたインターフェイス、Spring/Express/FastAPI/Rust のルート宣言、構成の読み取り、データベース テーブルのマッピングなどのグラフのファクトを作成できます。 Git 履歴分析はローカルであり、デフォルトで有効になっています。ok インデックスと ok ウォッチは、境界のあるローカル履歴ウィンドウを読み取り、型指定されたコミット メタデータと完全なファイルごとのタッチ (検出された名前変更を含む) を保存し、計画、ランキング、影響、契約、およびテスト選択のための、同時変更、テストへのパスの同時実行、チャーン、来歴、所有権、レビュー担当者、および同様の変更の証拠を保存します。履歴スコアのコンポーネントには、history_churn、ownership_risk、similar_change_overlap、reviewer_affinity という名前が付けられます。正確な参照と正確なシンボル/ファイルの証拠は依然として信頼できます。 [history] ​​max_commits = 500 でウィンドウを構成するか、 ok.t​​oml で [history] ​​Enabled = false でウィンドウを無効にします。
ランタイム分析はオプトイン証拠取り込みのみです。ローカルの JSONL トレース、スパン、ログ、インシデント、エラー、または失敗アーティファクトを、ソース ファイル パスとサービス、シンボル/関数、ルート、HTTP メソッド、ステータス コード、期間、タイムスタンプ、SQL ステートメント、メッセージなどのフィールドとともに .ok/runtime/ または .ok/analysis/runtime/ の下に配置し、 ok Index を再実行します。 Open Kioku は一致するエントリを ru に変換します

[切り捨てられた]

## Original Extract

Local-first code intelligence MCP for AI coding agents. - shivyadavus/open-kioku

GitHub - shivyadavus/open-kioku: Local-first code intelligence MCP for AI coding agents. · GitHub
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
shivyadavus
/
open-kioku
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
219 Commits 219 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor-plugin .cursor-plugin .githooks .githooks .github .github Formula Formula assets assets benchmarks benchmarks crates crates demo demo docs docs examples examples packages packages scripts scripts skills/ open-kioku skills/ open-kioku test-repos/ rust-mini test-repos/ rust-mini .gitignore .gitignore .mcp.json .mcp.json CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md STABILITY.md STABILITY.md audit.toml audit.toml claude_plugin.json claude_plugin.json demo.tape demo.tape deny-output.txt deny-output.txt deny.toml deny.toml glama.json glama.json jobs.json jobs.json mcp.json mcp.json release-metadata.json release-metadata.json run.json run.json smithery.yaml smithery.yaml View all files Repository files navigation
Give Claude, Cursor, Codex, or any MCP-compatible coding agent an evidence layer before it edits your codebase, using local indexes and read-only MCP tools by default.
Run these three commands on your own repo and see what changes:
npm install -g open-kioku
ok index .
ok mcp install cursor --repo . # or: claude, codex, gemini, windsurf, trae, zed, opencode
Paste the printed MCP config into your agent, then paste this prompt:
Use Open Kioku before editing. Check repo_status, search_code, get_definition,
get_references, impact_analysis, and find_tests_for_change. Build a plan with
plan_change first, then edit, and verify after the edit with verify_change.
Your agent now has indexed evidence — symbols, references, impact analysis, test candidates, and edit boundaries — before it touches a single file.
Run ok prove on your repo to produce a shareable Markdown or HTML report with indexed counts, task scores, and validation signals. The report intentionally omits source snippets, so it is designed to be reviewed and shared from private repos:
ok prove . --task " the feature you're working on "
ok prove . --task " the feature you're working on " --html
A representative public-repo audit indexed 4,600+ files, 46,000+ symbols, and 8,900+ tests locally in 33.1s. The exact Open Kioku version, repository revisions, caveats, and language limitations are recorded in docs/large-repo-proof.md .
search_code → get_definition → impact_analysis → find_tests_for_change → plan_change → verify_change
An evidence-backed pre-edit plan with: primary files, relevant symbols, likely blast radius, exact validation commands, confidence scores, and the next MCP calls to make.
ok plan and MCP plan_change return concrete file ranges, validation candidates, and confidence caveats from the local index. On the bundled demo repo, a token task currently surfaces:
primary context in src/auth.rs , src/lib.rs , and tests/auth_flow.rs
validation candidates such as issue_token , validate_token , and login_returns_valid_token via cargo test
an edit boundary that allows the relevant source and test files while keeping generated, vendored, build, and .ok/ artifacts out of scope
explicit caveats when exact-reference, runtime, history, or coverage evidence is unavailable
AI coding agents are strongest when they can ask the codebase for facts before editing. Without indexed context, they burn tokens on repeated file crawling, infer references from text matches, and often pick tests only after a change has already gone wrong.
Open Kioku gives agents a pre-edit routine:
Search indexed code and files.
Resolve symbols and references.
Build an evidence-backed pre-edit plan with likely impact and validation targets.
Recall prior repo facts without letting memory outrank indexed code evidence.
Compress context into handles that can retrieve the original snippets later.
Serve those capabilities through MCP over local stdio.
npm install -g open-kioku
ok --version
The open-kioku npm package is a JavaScript wrapper that installs the native ok binary through platform-specific optional dependencies.
cargo binstall open-kioku-cli
ok --version
crates.io
cargo install open-kioku-cli
ok --version
GitHub Releases
Tagged releases publish native binaries and SHA-256 checksums for Linux x86_64/arm64 musl, macOS x86_64/arm64, and Windows x86_64. Download from https://github.com/shivyadavus/open-kioku/releases .
git clone https://github.com/shivyadavus/open-kioku.git
cd open-kioku
cargo install --path crates/open-kioku-cli
ok --help
Requires a stable Rust toolchain.
ok init /absolute/path/to/repo
ok index /absolute/path/to/repo
ok doctor /absolute/path/to/repo
ok status /absolute/path/to/repo --markdown --write ok-status.md
ok setup audit /absolute/path/to/repo
ok --repo /absolute/path/to/repo search " the feature or bug you care about " --limit 5
ok index writes local data under .ok/ : SQLite metadata and graph rows in .ok/index.sqlite , plus BM25 search data in .ok/search/tantivy . Large indexes report progress phases such as scan , parse , occurrences , store , graph , search , and complete .
Keep the index fresh while editing:
ok watch /absolute/path/to/repo
Connect Your Agent
ok mcp install cursor --repo /absolute/path/to/repo
ok mcp install claude --repo /absolute/path/to/repo
ok mcp install codex --repo /absolute/path/to/repo
ok mcp install gemini --repo /absolute/path/to/repo
ok mcp install windsurf --repo /absolute/path/to/repo
ok mcp install trae --repo /absolute/path/to/repo
ok mcp install opencode --repo /absolute/path/to/repo
ok mcp install zed --repo /absolute/path/to/repo
Paste the printed MCP config snippet into your agent. The default server is read-only and runs locally over stdio. The install matrix is tested for Claude, Cursor, Codex, Gemini CLI, OpenCode, Zed, Windsurf, and Trae.
Agent-specific setup guides: Claude · Cursor · Codex · Gemini CLI . OpenCode, Zed, Windsurf, and Trae use the same verified ok mcp install <client> --repo /absolute/path/to/repo flow shown above.
Open Kioku bundles pre-configured repository-scoped plugin and marketplace manifests:
OpenAI Codex : verified MCP setup is ok mcp install codex --repo /absolute/path/to/repo ; .codex-plugin/ metadata is included for Codex installations that support Git-scoped plugin manifests.
Claude : .claude-plugin/ metadata is included alongside the verified ok mcp install claude --repo /absolute/path/to/repo MCP config path.
Cursor : .cursor-plugin/ rules and skills are included alongside the verified ok mcp install cursor --repo /absolute/path/to/repo MCP config path.
Starter examples with golden prompts and one-command smoke tests are available
in examples/cursor and examples/claude .
ok status --markdown --write ok-status.md creates a portable handoff with index counts, SCIP quality, readiness checks, and next steps. ok setup audit checks the index, security posture, MCP server health, plugin metadata, and the supported client install matrix.
For a fleet or multi-service workspace, index each project first, then create a workspace config:
[ workspace ]
projects = [
{ name = " service-a " , repo = " ../service-a " },
{ name = " service-b " , repo = " ../service-b " },
]
Run the cross-project linker without reparsing source:
ok index --mode cross-project --workspace /absolute/path/to/workspace
ok architecture fleet --workspace /absolute/path/to/workspace
Index Snapshots
Teams and CI can share a known-good local index without sharing personal memory or compressed context state:
ok --repo /absolute/path/to/repo snapshot export --quality best
ok --repo /absolute/path/to/repo snapshot doctor
ok --repo /absolute/path/to/repo snapshot import
ok --repo /absolute/path/to/repo index --from-snapshot auto
Quality Mode
Open Kioku works without external language indexers, but exact references improve search grounding, impact analysis, test selection, and planning. Check what is available:
ok setup audit /absolute/path/to/repo --markdown --write ok-setup.md
ok scip doctor /absolute/path/to/repo
ok scip setup /absolute/path/to/repo
ok index /absolute/path/to/repo --with-scip auto
ok doctor /absolute/path/to/repo
Default indexing consumes existing SCIP files such as index.scip and .ok/indexes/*.scip when present. --with-scip auto runs installed indexers for supported repos; it does not install third-party tools. --with-scip required fails the index if no SCIP facts can be imported.
SCIP is the primary precision provider. The default quality model stays local and free: indexed symbols/chunks/imports, language-specific static facts, indexed tests, build-system detection, and SCIP exact references when an index.scip is available. Java/Gradle repositories get scoped validation commands when the test file path is known, for example ./gradlew :x-pack:plugin:ml:test --tests org.example.SomeTests instead of a generic repository-wide test command.
Language-specific static analysis adds durable graph facts such as imports, Java inheritance and implemented interfaces, Spring/Express/FastAPI/Rust route declarations, config reads, and database table mappings. Git-history analysis is local and enabled by default: ok index and ok watch read a bounded local history window, store typed commit metadata and complete per-file touches (including detected renames), and preserve co-change, path-to-test co-run, churn, provenance, ownership, reviewer, and similar-change evidence for planning, ranking, impact, contracts, and test selection. History score components are named history_churn , ownership_risk , similar_change_overlap , and reviewer_affinity ; exact references and exact symbol/file evidence remain authoritative. Configure the window with [history] max_commits = 500 , or disable it with [history] enabled = false in ok.toml .
Runtime analysis is opt-in evidence ingestion only: place local JSONL trace, span, log, incident, error, or failure artifacts under .ok/runtime/ or .ok/analysis/runtime/ with source file paths plus fields such as service, symbol/function, route, HTTP method, status code, duration, timestamp, SQL statement, and message, then re-run ok index . Open Kioku turns matching entries into ru

[truncated]
