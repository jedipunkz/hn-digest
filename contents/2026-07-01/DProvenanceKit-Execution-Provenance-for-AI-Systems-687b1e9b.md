---
source: "https://github.com/Therealdk8890/DProvenanceKitPython"
hn_url: "https://news.ycombinator.com/item?id=48743043"
title: "DProvenanceKit: Execution Provenance for AI Systems"
article_title: "GitHub - Therealdk8890/DProvenanceKitPython: Reasoning observability and regression testing for AI systems — a Python port of DProvenanceKit. · GitHub"
author: "DPK890"
captured_at: "2026-07-01T07:25:37Z"
capture_tool: "hn-digest"
hn_id: 48743043
score: 1
comments: 0
posted_at: "2026-07-01T06:35:10Z"
tags:
  - hacker-news
  - translated
---

# DProvenanceKit: Execution Provenance for AI Systems

- HN: [48743043](https://news.ycombinator.com/item?id=48743043)
- Source: [github.com](https://github.com/Therealdk8890/DProvenanceKitPython)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T06:35:10Z

## Translation

タイトル: DProvenanceKit: AI システムの実行来歴
記事のタイトル: GitHub - Therealdk8890/DProvenanceKitPython: AI システムの推論可観測性と回帰テスト — DProvenanceKit の Python ポート。 · GitHub
説明: AI システムの推論可観測性と回帰テスト — DProvenanceKit の Python ポート。 - Therealdk8890/DProvenanceKitPython

記事本文:
GitHub - Therealdk8890/DProvenanceKitPython: AI システムの推論可観測性と回帰テスト — DProvenanceKit の Python ポート。 · GitHub
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
Therealdk8890
/
DProvenanceKitPython
公共
通知
chan にサインインする必要があります

GE通知設定
追加のナビゲーション オプション
コード
Therealdk8890/DProvenanceKitPython
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
68 コミット 68 コミット .github/ workflows .github/ workflows アクション アクション適合性 適合性の例 例 gitlab gitlab src/ dprovenancekit src/ dprovenancekit テスト テスト .gitattributes .gitattributes .gitignore .gitignore ライセンス ライセンス通知 通知 README.md README.md RELEASING.md RELEASING.md pyproject.toml pyproject.toml Record_example.py Record_example.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI システムの推論可観測性と回帰テスト — Swift DProvenanceKit の Python ポート。
エージェントの推論が実行間で変動する場合、DProvenanceKit は各実行をクエリ可能で比較可能なトレースに変換するため、何が起こったのかだけでなく、何が変更されたのか、なぜ変更されたのかを確認できるようになります。
実行 → 記録 → クエリ → 差分 → 回帰検出 → CI でのゲート
これは単なるライブラリではありません。推論回帰を実行可能にするサーフェスが同梱されています。
Gate in CI — サーバーレスの dprovenancekit ゲート CLI に加え、エージェントの推論がゴールデン ベースラインから逸脱した場合に PR/MR を失敗させ、差分をコメントするドロップイン GitHub Action および GitLab CI テンプレート。
すぐに使用できる異常ルール - JSON ルール レジストリを使用したツール ドロップとループの検出。ローカルまたはすべての PR で実行可能。
ホスト型ビジュアライザ - 回帰ゲート API とマルチテナント コントロール プレーンをサポートする Web ダッシュボード (単一実行のスパン ツリー、JSON ペイロード インスペクター、並列セマンティック差分、共有可能な HTML レポート)。別途商用サービスとしてご利用いただけます。
1 つの実行可能なスクリプト ( python Examples/end_to_end_demo.py ) ですべてを確認できます。
これは、Swift ライブラリを Python に忠実かつ依存関係なく移植したものです。同じアーキテクチャを維持します

そして保証 - 同期ノンブロッキング記録、優先度を意識したバックプレッシャー、パリティで保持される 2 つのバックエンドに対する 1 つのクエリ言語、構造的な差分、正式にモデル化されたセマンティック アライメント、および層ごとのドロップ アカウンティングにより、負荷遮断が決してサイレントになることはありません。
元の Swift パッケージは変更されていません。これは並列実装です。
Swift ライブラリは、Apple プラットフォームとオンデバイス AI をターゲットとしています。このポートは、サードパーティの依存関係なしで、Python コードベース (エージェント フレームワーク、LLM ワークフロー、ツール使用モデル) に同じ推論層の可観測性をもたらします (標準ライブラリ: sqlite3 、 contextvars 、 threading 、 json 、 hashlib 、 uuid 、 urllib のみを使用します)。
pip インストール dprovenancekit
pip install " dprovenancekit[langchain] " # + LangChain アダプター
pip install " dprovenancekit[openai-agents] " # + OpenAI Agents アダプター
チェックアウト (開発) から:
pip install -e " .[dev] "
Python 3.9 以降が必要です。コアにはサードパーティへの依存関係がありません。リリースは文書化されています
RELEASING.md にあります。
1 つの実行可能なスクリプトでアーク全体を実行したい — 記録 → クエリ → ゲート → 異常の検出 →
diff→レポートを実行し、同じ実行を CLI に渡しますか?走る
Python の例/end_to_end_demo.py 。以下の手順で少しずつ構築していきます。
TraceableEvent をサブクラス化し、安定した type_identifier と priority を公開する凍結されたデータクラス:
データクラスからデータクラスをインポート
dprovenancekit からインポート TraceableEvent 、 TracePriority
@データクラス (凍結 = True)
クラス MyAIDecision ( TraceableEvent ):
種類: str # "promptGenerated" | "ドキュメント評価済み" | "競合が検出されました" | 「最終決定」
トークン数: int = 0
document_id : str = ""
スコア: 浮動小数点 = 0.0
理由: str = ""
承認済み : bool = False
@プロパティ
def type_identifier ( self ) -> str :
返却自己。種類
@プロパティ
デフォルト優先度 (自分)

-> トレース優先度:
自分自身の場合種類 == "finalDecisionMade" :
TracePriority を返します。クリティカル
自分自身の場合種類 == "競合が検出されました" :
TracePriority を返します。診断
TracePriority を返します。テレメトリ
2. 実行実行を記録する
Record(...) は同期的であり、ブロックされることはありません。メモリ内のバッファーのみに触れます。アンビエント run/engine/span コンテキストは contextvars を介して伝播するため、ネストされたスコープは配管なしでイベントを正しく属性付けします。
from dprovenancekit import DProvenanceKit 、 InMemoryTraceStore
キット = DProvenanceKit ( MyAIDecision )
ストア = InMemoryTraceStore()
キット付き。 run ( context_id = "demo_case" 、store = store ):
キット。レコード ( MyAIDecision ( kind = "documentEvaluated" 、 document_id = "DocA" 、スコア = 0.95 ))
キット。レコード ( MyAIDecision ( kind = "conflictDetected" 、reason = "timeline_inconsistency" ))
キット。レコード ( MyAIDecision ( kind = "finalDecisionMade" 、承認済み = False ))
3. クエリ推論パターン
dprovenancekit から TraceQueryDSL をインポート
怪しい＝店にquery_runs (
トレースクエリDSL ()
。 required_step ( "競合検出" )
。 missing_step ( "documentEvaluated" )
)
Find は、競合が報告されたがドキュメントが評価されていない場所で実行されます。同じ DSL は SQLiteTraceStore 用に SQL にコンパイルされ、InMemoryTraceStore 用にメモリ内で評価されます。2 つのバックエンドはパリティ テスト スイートによってロックステップで保持されます。
dprovenancekit から TraceDiffEngine をインポート
diff = TraceDiffEngine ()。 diff (ベース = run_a 、比較 = run_b )
print ( diff .changes ) # 出現、消滅、または移動された構造ステップ
5. セマンティックな調整
TraceAlignmentEngine は、ペイロードがわずかに異なる場合でも、正式に定義されたセマンティック モデル内で 2 つの実行が動作的に同等であるかどうかを判断します。
dprovenancekit インポートから (
AlignmentConfiguration 、AlignmentProfile 、AnyEqui

valenceEvaluator、TraceAlignmentEngine、
)
config = AlignmentConfiguration (
プロファイル = アライメントプロファイル 。 strict_audit_v1 、
equivalence_evaluator = AnyEquivalenceEvaluator (
evaluator_identifier = "MyAIDecision_Semantic" ,
evaluator = lambda a , b : a == b の場合は 1.0、それ以外の場合は 0.0 、
）、
)
結果 = TraceAlignmentEngine ( config )。 align (ベース = run_a 、比較 = run_b )
print (結果.回帰リスク.レベル)
6.回帰を自動的に検出する
dprovenancekit からインポート AnomalyDetector 、 AnomalyRule 、 TraceQueryDSL
クラス UnverifiedConflictRule ( AnomalyRule ):
@プロパティ
デフォルト名 ( self ): "unverified_conflict" を返します
@プロパティ
def anomaly_query ( self ):
TraceQueryDSL () を返します。 required_step ( "conflictDetected" )。 missing_step ( "documentEvaluated" )
def description ( self , run ): return "サポートされる評価なしで競合が検出されました"
異常 = AnomalyDetector (ストア)。 detect_anomalies ([ UnverifiedConflictRule ()])
または、独自のルールを作成する代わりに、組み込みライブラリから既製のルールをドロップします。
dprovenancekit からインポート AnomalyDetector 、 LoopingRule 、 ToolDropRule
異常 = AnomalyDetector (ストア)。 detect_anomalies ([
ToolDropRule ( "safety_check" )、 # 必要なステップを実行しませんでした
LoopingRule ( "web_search" , max_repeats = 5 )、 # 同じツール呼び出しを繰り返してスタックする
])
7. リグレッションに対するプルリクエストのゲート処理
サーバーを使用せずに CI で回帰ゲートを実行します。ローカル SQLite トレース データベースをポイントします。
およびゴールデン/候補実行 ID。終了コードは 0 (パス)、1 (回帰)、または 2 (使用エラー) です。
dprovenancekit ゲート --db traces.sqlite --golden " $GOLDEN_RUN_ID " --candidate " $CANDIDATE_RUN_ID "
dprovenancekit ゲート --db traces.sqlite --golden " $G " --candidate " $C " --max-level low --json
# 別々のデータベースにまたがるゲート (復元されたベースラインとこの PR の実行)、解決

# ハードコーディングする代わりに、ベースラインからのゴールデン ラン ID を取得します。
GOLDEN= $( dprovenancekit 実行 --dbbaseline.sqlite --context my-agent --latest --format id )
dprovenancekit ゲート --golden-db Baseline.sqlite --golden " $GOLDEN " \
--candidate-db 候補.sqlite --candidate " $CANDIDATE_RUN_ID "
事前構築された CI 統合はこれをラップし、PR/MR の差分をコメントします。
GitHub アクションと GitLab CI テンプレート。
このライブラリには、Swift バージョンと同じ検証コーパスが同梱されています。ヘッドレス CLI は、実際のベンチマーク ランナーを通じて実行します。
dprovenancekit は、標準 + 敵対的コーパスに対して # 精度/再現率/F1 を評価します
dprovenancekit detect # 故障モードの因果関係のランキング
dprovenancekit の安定性 # 決定性境界: 孤立した F1 分散と摂動された F1 分散
どちらのコーパス スコアも精度 1.000 / 再現率 1.000 / F1 1.000 — 8 つの標準シナリオ (並べ替え、セマンティック進化、ノイズ挿入、分岐崩壊など) と 5 つの敵対的堅牢性トラップ (依存性反転、部分切り捨て、セマンティック置換など) — ケースごとに Swift 実装と一致します。
コンポーネント
モジュール
イベントモデル、優先順位、ドロップアカウンティング
イベント、優先度、drop_stats
レコーディング API + アンビエント コンテキスト
キット、コンテキスト
ストア (インメモリ、WAL SQLite、生読み取り、クラウド)
ストア、sqlite_store、raw_store、cloud_store
優先順位を意識した書き込みバッファ
書き込みバッファ
クエリ DSL + 2 つのバックエンド (AST 評価 + SQL コンパイラ)
クエリ
ライブクエリ + 異常検出 + ルールライブラリ
live_engine 、異常 、ルール
構造差分 + スパン対応スナップショット差分
diff 、 snapshot_diff
決定論的再生
リプレイ
セマンティック調整エンジン + 証拠 + 検証
アライメント_* 、検証
ベンチマークハーネス、障害診断ツール、コーパス
ベンチマーク、コーパス
トレース ビューア用の純粋なビュー モデル
ビューモデル
フレームワークアダプター (LangChain)

/ランググラフ)
統合.langchain
フレームワークアダプター (OpenAI Agents SDK)
統合.openai_agents
回帰ゲート テスト ヘルパー
テスト
共有可能な HTML 回帰レポート
報告書
フレームワークに依存しないインストゥルメンテーション (デコレータ)
楽器
ヘッドレス CLI — ゲート、異常、実行、評価
クリ
SwiftUI DProvenanceUI ターゲットは意図的に移植されていません (Apple プラットフォーム UI です)。その純粋な値モデル層 ( SpanViewModel 、フラット化) は viewmodel に移植されます。
Swift と Python SDK の動作を同等に保つことは、期待されるものではなく強制されます。 conformance/ は、Trace 仕様 v1 を保持します。これは、言語に依存しないコントラクトと、実行フィンガープリント、アライメント プロファイル ハッシュ、正規ペイロード エンコーディング、クエリ セマンティクス、およびアライメント判定を固定する凍結されたゴールデン ベクトルです。
python -m pytest testing/test_conformance.py # Python SDKの適合性の主張
python conformance/generate_vectors.py # 意図的にコントラクトを再凍結する
コミットされた準拠/vectors/*.json は契約です。どの SDK (今日では Swift、その後の Rust や TypeScript) も、同じファイルを再現することで同等であることが証明されます。 conformance/TRACE_SPEC_v1.md を参照してください。
フレームワーク アダプターは dprovenancekit.integrations 内に存在し、サードパーティの依存関係を持つパッケージの唯一の部分です。コアは PU のままです。

[切り捨てられた]

## Original Extract

Reasoning observability and regression testing for AI systems — a Python port of DProvenanceKit. - Therealdk8890/DProvenanceKitPython

GitHub - Therealdk8890/DProvenanceKitPython: Reasoning observability and regression testing for AI systems — a Python port of DProvenanceKit. · GitHub
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
Therealdk8890
/
DProvenanceKitPython
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Therealdk8890/DProvenanceKitPython
main Branches Tags Go to file Code Open more actions menu Folders and files
68 Commits 68 Commits .github/ workflows .github/ workflows action action conformance conformance examples examples gitlab gitlab src/ dprovenancekit src/ dprovenancekit tests tests .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE NOTICE NOTICE README.md README.md RELEASING.md RELEASING.md pyproject.toml pyproject.toml record_example.py record_example.py View all files Repository files navigation
Reasoning observability and regression testing for AI systems — a Python port of the Swift DProvenanceKit .
When an agent's reasoning drifts between runs, DProvenanceKit turns each execution into a queryable, diffable trace so you can see what changed and why — not just what happened .
Run → Record → Query → Diff → Detect regressions → Gate in CI
It's not just the library — it ships the surfaces that make reasoning regressions actionable:
Gate in CI — a server-less dprovenancekit gate CLI, plus a drop-in GitHub Action and GitLab CI template that fail a PR/MR when an agent's reasoning drifts from a golden baseline, and comment the diff.
Out-of-the-box anomaly rules — Tool Drop and Looping detection with a JSON rule registry, runnable locally or on every PR.
A hosted visualizer — a web dashboard (single-run span tree, JSON payload inspector, side-by-side semantic diff, shareable HTML reports) backed by a regression-gate API and multi-tenant control plane. Available as a separate commercial service.
See it all in one runnable script: python examples/end_to_end_demo.py .
This is a faithful, dependency-free port of the Swift library to Python. It keeps the same architecture and guarantees — synchronous non-blocking recording, priority-aware backpressure, one query language over two backends held at parity, structural diffing, formally-modeled semantic alignment, and by-tier drop accounting so load-shedding is never silent.
The original Swift package is unchanged; this is a parallel implementation.
The Swift library targets Apple-platform and on-device AI. This port brings the same reasoning-layer observability to Python codebases — agent frameworks, LLM workflows, tool-using models — with zero third-party dependencies (it uses only the standard library: sqlite3 , contextvars , threading , json , hashlib , uuid , urllib ).
pip install dprovenancekit
pip install " dprovenancekit[langchain] " # + LangChain adapter
pip install " dprovenancekit[openai-agents] " # + OpenAI Agents adapter
From a checkout (development):
pip install -e " .[dev] "
Requires Python 3.9+; the core has zero third-party dependencies . Releasing is documented
in RELEASING.md .
Want the whole arc in one runnable script — record → query → gate → detect anomalies →
diff → report, then hand the same runs to the CLI? Run
python examples/end_to_end_demo.py . The steps below build it up piece by piece.
Any frozen dataclass that subclasses TraceableEvent , exposing a stable type_identifier and a priority :
from dataclasses import dataclass
from dprovenancekit import TraceableEvent , TracePriority
@ dataclass ( frozen = True )
class MyAIDecision ( TraceableEvent ):
kind : str # "promptGenerated" | "documentEvaluated" | "conflictDetected" | "finalDecisionMade"
token_count : int = 0
document_id : str = ""
score : float = 0.0
reason : str = ""
approved : bool = False
@ property
def type_identifier ( self ) -> str :
return self . kind
@ property
def priority ( self ) -> TracePriority :
if self . kind == "finalDecisionMade" :
return TracePriority . CRITICAL
if self . kind == "conflictDetected" :
return TracePriority . DIAGNOSTIC
return TracePriority . TELEMETRY
2. Record an execution run
record(...) is synchronous and never blocks — it touches only an in-memory buffer. Ambient run / engine / span context propagates through contextvars , so nested scopes attribute events correctly with no plumbing.
from dprovenancekit import DProvenanceKit , InMemoryTraceStore
kit = DProvenanceKit ( MyAIDecision )
store = InMemoryTraceStore ()
with kit . run ( context_id = "demo_case" , store = store ):
kit . record ( MyAIDecision ( kind = "documentEvaluated" , document_id = "DocA" , score = 0.95 ))
kit . record ( MyAIDecision ( kind = "conflictDetected" , reason = "timeline_inconsistency" ))
kit . record ( MyAIDecision ( kind = "finalDecisionMade" , approved = False ))
3. Query reasoning patterns
from dprovenancekit import TraceQueryDSL
suspicious = store . query_runs (
TraceQueryDSL ()
. requiring_step ( "conflictDetected" )
. missing_step ( "documentEvaluated" )
)
Find runs where a conflict was reported but no document was ever evaluated. The same DSL compiles to SQL for SQLiteTraceStore and is evaluated in memory for InMemoryTraceStore — the two backends are held in lockstep by a parity test suite.
from dprovenancekit import TraceDiffEngine
diff = TraceDiffEngine (). diff ( base = run_a , comparison = run_b )
print ( diff . changes ) # structural steps that appeared, disappeared, or moved
5. Semantic alignment
TraceAlignmentEngine decides whether two executions are behaviorally equivalent within a formally-defined semantic model, even when payloads vary slightly:
from dprovenancekit import (
AlignmentConfiguration , AlignmentProfile , AnyEquivalenceEvaluator , TraceAlignmentEngine ,
)
config = AlignmentConfiguration (
profile = AlignmentProfile . strict_audit_v1 ,
equivalence_evaluator = AnyEquivalenceEvaluator (
evaluator_identifier = "MyAIDecision_Semantic" ,
evaluator = lambda a , b : 1.0 if a == b else 0.0 ,
),
)
result = TraceAlignmentEngine ( config ). align ( base = run_a , comparison = run_b )
print ( result . regression_risk . level )
6. Detect regressions automatically
from dprovenancekit import AnomalyDetector , AnomalyRule , TraceQueryDSL
class UnverifiedConflictRule ( AnomalyRule ):
@ property
def name ( self ): return "unverified_conflict"
@ property
def anomaly_query ( self ):
return TraceQueryDSL (). requiring_step ( "conflictDetected" ). missing_step ( "documentEvaluated" )
def describe ( self , run ): return "Conflict detected with no supporting evaluation"
anomalies = AnomalyDetector ( store ). detect_anomalies ([ UnverifiedConflictRule ()])
Or drop in ready-made rules from the built-in library instead of writing your own:
from dprovenancekit import AnomalyDetector , LoopingRule , ToolDropRule
anomalies = AnomalyDetector ( store ). detect_anomalies ([
ToolDropRule ( "safety_check" ), # never ran a required step
LoopingRule ( "web_search" , max_repeats = 5 ), # stuck repeating the same tool call
])
7. Gate a pull request on regressions
Run the regression gate in CI with no server — point it at a local SQLite trace database
and a golden/candidate run id. Exit code is 0 (pass), 1 (regression), or 2 (usage error):
dprovenancekit gate --db traces.sqlite --golden " $GOLDEN_RUN_ID " --candidate " $CANDIDATE_RUN_ID "
dprovenancekit gate --db traces.sqlite --golden " $G " --candidate " $C " --max-level low --json
# Gate across separate databases (a restored baseline vs. this PR's run), resolving
# the golden run id from the baseline instead of hardcoding it:
GOLDEN= $( dprovenancekit runs --db baseline.sqlite --context my-agent --latest --format id )
dprovenancekit gate --golden-db baseline.sqlite --golden " $GOLDEN " \
--candidate-db candidate.sqlite --candidate " $CANDIDATE_RUN_ID "
Prebuilt CI integrations wrap this and comment the diff on the PR/MR:
a GitHub Action and a GitLab CI template .
The library ships the same validation corpus as the Swift version. The headless CLI runs it through the real benchmark runner:
dprovenancekit evaluate # precision/recall/F1 over the standard + adversarial corpora
dprovenancekit diagnose # causal ranking of failure modes
dprovenancekit stability # determinism boundary: isolated vs perturbed F1 variance
Both corpora score Precision 1.000 / Recall 1.000 / F1 1.000 — 8 standard scenarios (reordering, semantic evolution, noise injection, branch collapse, …) and 5 adversarial robustness traps (dependency inversion, partial truncation, semantic substitution, …) — matching the Swift implementation case-for-case.
Component
Module
Event model, priority tiers, drop accounting
event , priority , drop_stats
Recording API + ambient context
kit , context
Stores (in-memory, WAL SQLite, raw read, cloud)
store , sqlite_store , raw_store , cloud_store
Priority-aware write buffer
write_buffer
Query DSL + two backends (AST eval + SQL compiler)
query
Live querying + anomaly detection + rule library
live_engine , anomaly , rules
Structural diff + span-aware snapshot diff
diff , snapshot_diff
Deterministic replay
replay
Semantic alignment engine + evidence + verification
alignment_* , verification
Benchmark harness, failure diagnoser, corpus
benchmark , corpus
Pure view models for a trace viewer
viewmodel
Framework adapters (LangChain / LangGraph)
integrations.langchain
Framework adapters (OpenAI Agents SDK)
integrations.openai_agents
Regression-gate test helper
testing
Shareable HTML regression report
report
Framework-agnostic instrumentation (decorators)
instrument
Headless CLI — gate , anomalies , runs , evaluate
cli
The SwiftUI DProvenanceUI target is intentionally not ported (it is Apple-platform UI); its pure value-model layer ( SpanViewModel , flattening) is ported in viewmodel .
Keeping the Swift and Python SDKs behaviorally equivalent is enforced, not hoped for. conformance/ holds Trace Specification v1 — a language-neutral contract plus frozen golden vectors that pin the run fingerprint, the alignment profile hash, canonical payload encoding, query semantics, and alignment verdicts.
python -m pytest tests/test_conformance.py # the Python SDK's claim of conformance
python conformance/generate_vectors.py # intentionally re-freeze the contract
The committed conformance/vectors/*.json are the contract: any SDK — Swift today, Rust or TypeScript later — proves equivalence by reproducing the same files. See conformance/TRACE_SPEC_v1.md .
Framework adapters live in dprovenancekit.integrations and are the only parts of the package with third-party dependencies — the core stays pu

[truncated]
