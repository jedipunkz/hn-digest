---
source: "https://github.com/Nick-is-building/ast-guard"
hn_url: "https://news.ycombinator.com/item?id=48449545"
title: "Show HN: AST-guard – Fast, zero-cost structural checks for LLM code execution"
article_title: "GitHub - Nick-is-building/ast-guard: Pre-Execution Gate for AI-Generated Code. Deterministic static analysis layer between LLM code generation and execution. · GitHub"
author: "thinking-nick"
captured_at: "2026-06-08T18:55:44Z"
capture_tool: "hn-digest"
hn_id: 48449545
score: 1
comments: 0
posted_at: "2026-06-08T18:40:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AST-guard – Fast, zero-cost structural checks for LLM code execution

- HN: [48449545](https://news.ycombinator.com/item?id=48449545)
- Source: [github.com](https://github.com/Nick-is-building/ast-guard)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T18:40:57Z

## Translation

タイトル: Show HN: AST-guard – LLM コード実行のための高速かつゼロコストの構造チェック
記事のタイトル: GitHub - Nick-is-building/ast-guard: AI 生成コードの事前実行ゲート。 LLM コードの生成と実行の間の決定論的な静的解析レイヤー。 · GitHub
説明: AI 生成コードの事前実行ゲート。 LLM コードの生成と実行の間の決定論的な静的解析レイヤー。 - GitHub - Nick-is-building/ast-guard: AI 生成コードの事前実行ゲート。 LLM コードの生成と実行の間の決定論的な静的解析レイヤー。

記事本文:
GitHub - Nick-is-building/ast-guard: AI 生成コードの事前実行ゲート。 LLM コードの生成と実行の間の決定論的な静的解析レイヤー。 · GitHub
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
ニックは構築中です
/
アストガード
公共
通知
サインインする必要があります

o 通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
121 コミット 121 コミット .github .github ast_guard ast_guard ベンチマーク ベンチマークの例 例のテスト テスト .gitignore .gitignore ALLOWLIST.md ALLOWLIST.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 生成コードの事前実行ゲート
LLM コード生成とコード実行の間の決定層。 LLMはありません。 MLはありません。費用はかかりません。
ast-guard はコード生成と実行の間に位置します。 LLM で生成されたコードを AST に解析し、行が実行される前に決定的な判定を返します。これは、リンターでも、セキュリティ スキャナでも、サンドボックスでもありません。これは、コンプライアンスを検討することができず、モデルの推論トレースによって欺瞞できない決定論的な構造チェックです。
既存の 2 つの防御クラスにはギャップが残ります。
トレーニング時間の調整 (Anthropic、DeepMind) は、残留ではなく発生率を減らします。
推論時 LLM レビュー担当者 (TRACE、RewardHackWatch、EvilGenie) は、監視しているジェネレーターと障害モードを共有します。
ast-guard は、これらの防御と並んで決定的な 3 番目のレイヤーとして構築されます。現在では、構造的に明白なバイパスをスキャンごとのコストゼロで確実に捕捉します (「主な結果」を参照)。積極的な開発の方向性は、時間の経過とともに構造的範囲を拡大することです。そのため、セマンティックレビュー担当者は、意図と意味の判断という自分たちにしかできないことに集中できます。
LLM がコードを生成する
│
▼
┌─────────────┐
│ ast-guard ゲート │ ← 決定的、<10ms、コストゼロ
━─────────

┘
│
┌────┴────┐
▼ ▼
クリーンな警告 / クリティカル
2 つのモード:
ペア モード — 元のコードと LLM で生成されたコードを比較します。構造的な逸脱を検出します。
スタンドアロン モード — ベースラインなしで単一エージェントの出力を分析します。行動リスクスコアリングを使用します。
完全な方法論、混同行列、カテゴリごとの内訳については、ベンチマーク/RESULTS.md を参照してください。
外部データセット (独立してラベル付け)
このプロジェクトによって作成されていないラベルが付いた、公開されているデータセットの結果。
† 通常ラベルのみの TNR (77,369 サンプル)。完全な混同行列については、RESULTS.md を参照してください。
‡ v2.1.0 で CHANGELOG に記録された数値。 JSON アーティファクトは Benchmarks/data/ に保存されず、v2.2.0 では再検証されていません。 python -m benchmarks.run_benchmark --benchmark <name> --json results.json を使用して再実行し、現在のアーティファクトを生成します。
手動で構築された組み込みスイート (外部ベンチマークではなく、回帰/健全性チェック)
これらのサンプルは、特定の構造パターンをカバーするためにプロジェクト作成者によって作成されました。これらは独立した外部データセットではありません。実世界のデータの検出率についての主張としてではなく、チェックが正しく実行されたことを確認するために使用します。
既存のアプローチとの比較
これらのアプローチは補完的なものであり、競合するものではありません。 ast-guard は構造解析を処理します。 LLM レビュー担当者はセマンティクスを処理します。
ast-guard は、コードを抽象構文ツリーに解析し、構造プロパティを評価します。実行もサンプリングも確率的推論もありません。
ハードコーディングの検出 — if カウント、リテラル カウント、長い文字列の増加とベースラインの比較。ガード条項は除外されます。
複雑さの崩壊 — 正当な最適化が認められない場合、関数ごとの McCabe の複雑さは 60% を超えて低下します。
禁止された呼び出しと難読化 — 新しい eval / exec / subprocess / ctypes / SystemExit 呼び出し、エイリアスの差分ベースの検出

解像度、chr() -難読化、組み込みの添字。
インポート ドリフト — ブロックリスト (CRITICAL) およびセーフリスト (CLEAN) に対する新しいインポート。不明なインポート → 警告。
Extensional Enumeration — Helff et al. の RLVR ショートカット概念の Python 類似物: ループのない分岐の 70% 以上をカバーするフラットな if/elif または match/case チェーン。ヘルフは帰納論理タスク (プロローグ形式のルール帰納法) の概念を研究しました。ここでの if/elif 検出器は ast-guard 独自の運用化であり、Helff が直接測定したパターンではありません。
行動リスク スコアリング (スタンドアロンのみ) — AST パターンからの加算的な YARA/Semgrep スタイルのスコア。クリーン <30、警告 30 ～ 69、クリティカル ≥70。
CLEAN → スコアがしきい値を下回っており、ブロックリストはトリガーされません
警告 → 疑わしいパターン、手動で確認することをお勧めします
クリティカル → 信頼性の高い構造ハッキング、ブロック実行
モード: 厳密なブロック CRITICAL。標準はすべてをログに記録します。サイレント収集のみを監査します。
要件: Python 3.11 以降。 Python 分析のための外部依存関係はゼロです。
git clone https://github.com/Nick-is-building/ast-guard.git
CD アストガード
python -m pytest テスト/ -q
Python API
ast_guard インポートスキャンから、scan_standalone
result = scan (original_code 、 generated_code 、mode = "strict" )
if result [ "verdict" ] == "CRITICAL" :
print ( "ブロックされました: 構造的なハッキングが検出されました。" )
print (結果 [ "チェック" ])
elif 結果 [ "評決" ] == "警告" :
print ( "疑わしい。レビューをお勧めします。" )
# スタンドアロン: 単一エージェント出力、ベースラインなし
結果 = scan_standalone (agent_code)
print (結果 [ "評決" ]、結果 [ "チェック" ][ "チェック_6_動作" ][ "スコア" ])
CLI
python -m ast_guard.cli checkoriginal.py generated.py # 標準
python -m ast_guard.cli checkoriginal.py generated.py --mode strict
python -m ast_guard.cli checkoriginal.py generated.py --json #

パイプライン用
# 多言語: 自動検出または明示的に指定
python -m ast_guard.cli checkoriginal.sh generated.sh -- language bash
python -m ast_guard.cli チェックoriginal.js generated.js -- language javascript
python -m ast_guard.cli checkoriginal.py generated.py -- language auto # デフォルト
python -m ast_guard.cli checkoriginal.py generated.py --no-multilang # Python のみ
CLEAN/WARNING の場合は終了コード 0、CRITICAL の場合は終了コード 1 — CI ゲートのドロップイン。
Python はネイティブ (ゼロ deps) です。 Bash と JavaScript は、オプションとして Tree-sitter 経由で利用できます。
pip install ast-guard[multilang]
言語
バックエンド
アクティブなチェック
パイソン
ネイティブアスト
1、2、3、4、5、6
バッシュ
ツリーシッターバッシュ
1、3、4、5、6
JavaScript
ツリーシッター-JavaScript
1、3、4、5、6
3 つの言語はすべて、同じ 6 チェック パイプラインを実行します。チェック 2 (複雑性の崩壊) はペアモードのベースラインを必要とし、スタンドアロン モードではすべての言語で非アクティブになります。言語は生成されたファイルから自動検出されます (最初にシバン、次にキーワード スコアリング) か、 -- language を使用して明示的に設定できます。
Bash のチェック 5 (拡張列挙): リテラル分岐値を含む case/esac ステートメントと、 [[ $x == "y" ]] スタイルの比較を含む if/elif を検出します。
Bash のチェック 6 (行動リスク スコアリング): eval_dynamic、pipe_to_shell、process_termination、subprocess_shell、network_fetch、test_file_write、environ_mutation、startup_persistence、destructive_call。
JavaScript のチェック 5: 文字列/数値リテラルの switch/case、および === / == 比較の if/else-if を検出します。
JavaScript のチェック 6: eval_dynamic (Function() コンストラクターを含む)、process_termination、subprocess_shell、dangerous_import (child_process)、test_file_write、environ_mutation、module_cache_manipulation (require.cache)。
ast-guard には、モデル コンテキスト プロトコル サーバーが組み込まれています。
ピップイン

背の高いアストガード[mcp]
{
"mcpサーバー": {
"アストガード" : {
"コマンド" : " ast-guard-mcp " ,
"タイプ" : "stdio"
}
}
}
ツール: ast_guard_scan (オリジナルと生成されたものを比較)、ast_guard_フィードバック (問題切り分けフィードバックを送信)。
- 使用: ./.github/actions/ast-guard
付き:
オリジナル: path/to/original.py
生成された : path/to/generated.py
モード : 厳密
アップロードサリフ: " true "
SARIF 出力は、GitHub の [セキュリティ] タブと互換性があります。
しきい値、ブロックリスト、許可リストは TOML 経由で構成できます。階層: CLI 引数 > .ast-guard.toml > ~/.ast-guard/config.toml > デフォルト。
[しきい値]
if_count_rel_increase = 0.50
リテラルカウント_レル増加 = 2.0
リテラルカウント_abs_min = 10
長い文字列の長さ = 200
複雑さ_相対減少 = 0.60
複雑性_abs_min = 5
列挙率 = 0.70
enumeration_min_ifs = 5
【輸入品】
blocklist = [" os " 、 " sys " 、 " サブプロセス " 、 " pickle " 、 " importlib " ]
allowlist = [" functools " , " itertools " , " collections " , " math " ]
[設定]
モード = "標準"
テレメトリ = false
multilang = "自動" # "自動" |本当 |偽
設計原則
決定論的。同じ入力→同じ判定。ランダム性やネットワーク呼び出しはありません。
コア内の依存関係はゼロです。 Pythonの標準ライブラリのみ。
リアルタイム。スキャンあたり 10 ミリ秒未満、実行ループの低下なし。
設計により補完的です。構造解析は 1 つのレイヤーです。セマンティックレビューは別です。
プライバシーバイデザイン。テレメトリでは、匿名化された AST メトリクスのみが保存され、コード、ファイル名、パス、タイムスタンプは保存されません。常にオプトインします。
3 つのモード (監査 → 標準 → 厳格) により、コミット前にリスクのない評価が可能になります。
設計上、スタンドアロンの精度はペア モードよりも低くなります。ベースラインがないと、文脈上の曖昧さをなくすことが難しくなります。正常サンプルの MALT 偽陽性率 (v2.2.0 では 5.0%) が現在のキャリブレーション ポイントです — 意図的です

まさに保守的。
セマンティック ハッキングは範囲外です。構造的には正常だが論理的に間違っているコード、または構造の痕跡を残さずにタスク仕様を操作するコードには、セマンティックな理解が必要です。それは、LLM ベースのレビュー担当者と下流のテスト スイートの仕事です。
Bash と JavaScript は、スタンドアロン モードでの検出上限をバイパスします。 MALT では、Bash バイパス サンプルの欠落 316/322 個と JS バイパス サンプルの欠落 30/30 個で構造シグナルはゼロです。これらは構造的にエンコードされているのではなく、意味的にエンコードされています (サイレント障害、間違ったロジック)。ペア モード (元のコードを使用) は、チェック 2 とチェック 1 を介して追加の信号を提供します。
しきい値は経験的に調整されます。デフォルトは、最終的な最適化ではなく、開始点として通知されます。完全な反復履歴については、METHODOLOGY.md を参照してください。
benchmarks/RESULTS.md — すべてのデータセットにわたる適合率、再現率、F1、混同行列。
benchmarks/METHODOLOGY.md — 回帰を含む 13 回の反復キャリブレーション履歴。
benchmarks/structural_benchmark/ — 12 の構造ハック カテゴリにわたる 36 の厳選されたグラウンド トゥルース ペア。
python -m benchmarks.run_benchmark --benchmark 構造
python -m benchmarks.run_benchmark --benchmark all
# MALT には ~/.ast-guard/benchmarks/malt-public/ にあるデータセットが必要です
python -m benchmarks.run_benchmark --benchmark malt --mode strict
関連作品
TRACE (Deshpande et al. 2026, arXiv:2601.20103 ) — 54 カテゴリの報酬ハッキング分類法。アスト-

[切り捨てられた]

## Original Extract

Pre-Execution Gate for AI-Generated Code. Deterministic static analysis layer between LLM code generation and execution. - GitHub - Nick-is-building/ast-guard: Pre-Execution Gate for AI-Generated Code. Deterministic static analysis layer between LLM code generation and execution.

GitHub - Nick-is-building/ast-guard: Pre-Execution Gate for AI-Generated Code. Deterministic static analysis layer between LLM code generation and execution. · GitHub
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
Nick-is-building
/
ast-guard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
121 Commits 121 Commits .github .github ast_guard ast_guard benchmarks benchmarks examples examples tests tests .gitignore .gitignore ALLOWLIST.md ALLOWLIST.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
Pre-Execution Gate for AI-Generated Code
The deterministic layer between LLM code generation and code execution. No LLM. No ML. No cost.
ast-guard sits between code generation and execution. It parses LLM-generated code into an AST and returns a deterministic verdict before any line runs. It is not a linter, not a security scanner, and not a sandbox — it is a deterministic structural check that cannot be talked into compliance and is not deceivable by a model's reasoning trace.
Two existing defense classes leave a gap:
Training-time alignment (Anthropic, DeepMind) reduces incidence, not residual.
Inference-time LLM reviewers (TRACE, RewardHackWatch, EvilGenie) share failure modes with the generator they monitor.
ast-guard is built as a deterministic third layer alongside those defenses. Today it reliably catches the structurally obvious bypasses at zero per-scan cost (see Key Results ); the active development direction is to broaden structural coverage over time, so semantic reviewers can concentrate on what only they can do — judging intent and meaning.
LLM generates code
│
▼
┌───────────────────┐
│ ast-guard gate │ ← deterministic, <10ms, zero cost
└───────────────────┘
│
┌────┴────┐
▼ ▼
CLEAN WARNING / CRITICAL
Two modes:
Pair Mode — compares original code against LLM-generated code. Detects structural deviations.
Standalone Mode — analyzes a single agent output with no baseline. Uses behavioral risk scoring.
See benchmarks/RESULTS.md for full methodology, confusion matrices, and per-category breakdowns.
External datasets (independently labeled)
Results on publicly released datasets with labels not authored by this project.
† TNR on the normal label only (77,369 samples); see RESULTS.md for the full confusion matrix.
‡ Numbers recorded in CHANGELOG at v2.1.0; no JSON artifact stored in benchmarks/data/ and not re-verified at v2.2.0. Re-run with python -m benchmarks.run_benchmark --benchmark <name> --json results.json to produce a current artifact.
Hand-constructed built-in suite (regression / sanity checks, not external benchmarks)
These samples were written by the project author to cover specific structural patterns. They are not independent external datasets — use them to confirm a check fires correctly, not as a claim about detection rates on real-world data.
Comparison with Existing Approaches
These approaches are complementary, not competing. ast-guard handles structural analysis; LLM reviewers handle semantics.
ast-guard parses code into an Abstract Syntax Tree and evaluates structural properties. No execution, no sampling, no probabilistic inference.
Hardcoding Detection — if-counts, literal counts, long-string growth vs. baseline. Guard-clauses excluded.
Complexity Collapse — per-function McCabe complexity drop >60% without a recognized legitimate optimization.
Forbidden Calls & Obfuscation — diff-based detection of new eval / exec / subprocess / ctypes / SystemExit calls, alias resolution, chr() -obfuscation, builtins subscript.
Import Drift — new imports against blocklist (CRITICAL) and safelist (CLEAN). Unknown imports → WARNING.
Extensional Enumeration — a Python analogue of the RLVR-shortcut concept from Helff et al.: flat if/elif or match/case chains covering ≥70% of branches with no loops. Helff studied the concept in inductive-logic tasks (Prolog-style rule induction); the if/elif detector here is ast-guard's own operationalization, not a pattern Helff measured directly.
Behavioral Risk Scoring (standalone only) — additive YARA/Semgrep-style score from AST patterns. CLEAN <30, WARNING 30–69, CRITICAL ≥70.
CLEAN → score below threshold, no blocklist triggers
WARNING → suspicious patterns, manual review recommended
CRITICAL → high-confidence structural hack, block execution
Modes: strict blocks CRITICAL; standard logs everything; audit silent collection only.
Requirements: Python 3.11+. Zero external dependencies for Python analysis.
git clone https://github.com/Nick-is-building/ast-guard.git
cd ast-guard
python -m pytest tests/ -q
Python API
from ast_guard import scan , scan_standalone
result = scan ( original_code , generated_code , mode = "strict" )
if result [ "verdict" ] == "CRITICAL" :
print ( "Blocked: structural hack detected." )
print ( result [ "checks" ])
elif result [ "verdict" ] == "WARNING" :
print ( "Suspicious. Review recommended." )
# Standalone: single agent output, no baseline
result = scan_standalone ( agent_code )
print ( result [ "verdict" ], result [ "checks" ][ "check_6_behavioral" ][ "score" ])
CLI
python -m ast_guard.cli check original.py generated.py # standard
python -m ast_guard.cli check original.py generated.py --mode strict
python -m ast_guard.cli check original.py generated.py --json # for pipelines
# Multi-language: auto-detect or specify explicitly
python -m ast_guard.cli check original.sh generated.sh --language bash
python -m ast_guard.cli check original.js generated.js --language javascript
python -m ast_guard.cli check original.py generated.py --language auto # default
python -m ast_guard.cli check original.py generated.py --no-multilang # Python-only
Exit code 0 on CLEAN/WARNING, exit code 1 on CRITICAL — drop-in for CI gates.
Python is native (zero deps). Bash and JavaScript are available via tree-sitter as an optional extra.
pip install ast-guard[multilang]
Language
Backend
Checks active
Python
Native ast
1, 2, 3, 4, 5, 6
Bash
tree-sitter-bash
1, 3, 4, 5, 6
JavaScript
tree-sitter-javascript
1, 3, 4, 5, 6
All three languages run the same 6-check pipeline. Check 2 (Complexity Collapse) requires a pair-mode baseline and is inactive in standalone mode for all languages. Language is auto-detected from the generated file (shebang-first, then keyword scoring) or can be set explicitly with --language .
Check 5 (Extensional Enumeration) for Bash: detects case/esac statements with literal branch values and if/elif with [[ $x == "y" ]] -style comparisons.
Check 6 (Behavioral Risk Scoring) for Bash: eval_dynamic, pipe_to_shell, process_termination, subprocess_shell, network_fetch, test_file_write, environ_mutation, startup_persistence, destructive_call.
Check 5 for JavaScript: detects switch/case with string/number literals and if/else-if with === / == comparisons.
Check 6 for JavaScript: eval_dynamic (including Function() constructor), process_termination, subprocess_shell, dangerous_import (child_process), test_file_write, environ_mutation, module_cache_manipulation (require.cache).
ast-guard includes a built-in Model Context Protocol server.
pip install ast-guard[mcp]
{
"mcpServers" : {
"ast-guard" : {
"command" : " ast-guard-mcp " ,
"type" : " stdio "
}
}
}
Tools: ast_guard_scan (compare original vs. generated), ast_guard_feedback (submit triage feedback).
- uses : ./.github/actions/ast-guard
with :
original : path/to/original.py
generated : path/to/generated.py
mode : strict
upload-sarif : " true "
SARIF output is compatible with the GitHub Security Tab.
Thresholds, blocklists, and allowlists are configurable via TOML. Hierarchy: CLI args > .ast-guard.toml > ~/.ast-guard/config.toml > defaults.
[ thresholds ]
if_count_rel_increase = 0.50
literal_count_rel_increase = 2.0
literal_count_abs_min = 10
long_string_len = 200
complexity_rel_decrease = 0.60
complexity_abs_min = 5
enumeration_ratio = 0.70
enumeration_min_ifs = 5
[ imports ]
blocklist = [ " os " , " sys " , " subprocess " , " pickle " , " importlib " ]
allowlist = [ " functools " , " itertools " , " collections " , " math " ]
[ settings ]
mode = " standard "
telemetry = false
multilang = " auto " # "auto" | true | false
Design Principles
Deterministic. Same input → same verdict. No randomness, no network calls.
Zero dependencies in core. Python standard library only.
Real-time. <10ms per scan, no degradation to the execution loop.
Complementary by design. Structural analysis is one layer; semantic review is another.
Privacy-by-design. Telemetry stores anonymized AST metrics only — never code, filenames, paths, or timestamps. Always opt-in.
Three modes ( audit → standard → strict ) enable risk-free evaluation before commitment.
Standalone precision is lower than pair mode by design. Without a baseline, contextual disambiguation is harder. The MALT false-positive rate on normal samples (5.0% at v2.2.0) is the current calibration point — intentionally conservative.
Semantic hacks are outside scope. Code that is structurally normal but logically wrong, or that games a task specification without a structural trace, requires semantic understanding. That is the job of LLM-based reviewers and downstream test suites.
Bash and JavaScript bypass detection ceiling in standalone mode. In MALT, 316/322 missed Bash bypass samples and 30/30 missed JS bypass samples show zero structural signal — they are semantically encoded (silent failure, wrong logic), not structurally encoded. Pair mode (with original code) would provide additional signal via Check 2 and Check 1.
Thresholds are empirically calibrated. Defaults are informed starting points, not final optima. See METHODOLOGY.md for the full iteration history.
benchmarks/RESULTS.md — precision, recall, F1, confusion matrices across all datasets.
benchmarks/METHODOLOGY.md — the 13-iteration calibration history, including regressions.
benchmarks/structural_benchmark/ — 36 curated ground-truth pairs across 12 structural hack categories.
python -m benchmarks.run_benchmark --benchmark structural
python -m benchmarks.run_benchmark --benchmark all
# MALT requires the dataset at ~/.ast-guard/benchmarks/malt-public/
python -m benchmarks.run_benchmark --benchmark malt --mode strict
Related Work
TRACE (Deshpande et al. 2026, arXiv:2601.20103 ) — 54-category reward-hacking taxonomy. ast-

[truncated]
