---
source: "https://github.com/AshwinUgale/tracelint"
hn_url: "https://news.ycombinator.com/item?id=49346452"
title: "Show HN: Tracelint – a linter for AI agent traces, no LLM judge"
article_title: "GitHub - AshwinUgale/tracelint: A deterministic linter for agent runs — reads the execution trace and flags structural bugs (ignored errors, schema violations, loops) with evidence. No LLM judge. · GitHub"
image: "https://opengraph.githubassets.com/fda4f4d554f823f9ab00c80caca70843c8402375a4b421a4ef4c9bbe0af6b94d/AshwinUgale/tracelint"
author: "Ashwin1121"
captured_at: "2026-08-18T15:22:31Z"
capture_tool: "hn-digest"
hn_id: 49346452
score: 1
comments: 0
posted_at: "2026-08-18T14:49:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tracelint – a linter for AI agent traces, no LLM judge

- HN: [49346452](https://news.ycombinator.com/item?id=49346452)
- Source: [github.com](https://github.com/AshwinUgale/tracelint)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T14:49:47Z

## Translation

タイトル: Show HN: Tracelint – AI エージェント トレースのリンター、LLM ジャッジなし
記事のタイトル: GitHub - AshwinUgale/tracelint: エージェント実行のための決定論的リンター — 実行トレースを読み取り、構造的なバグ (無視されたエラー、スキーマ違反、ループ) に証拠を付けてフラグを立てます。 LLM 審査員はいません。 · GitHub
説明: エージェント実行のための決定論的リンター — 実行トレースを読み取り、構造的なバグ (無視されたエラー、スキーマ違反、ループ) に証拠を付けてフラグを立てます。 LLM 審査員はいません。 - アシュウィン・ウガレ/tracelint

記事本文:
GitHub - AshwinUgale/tracelint: エージェント実行のための決定論的リンター — 実行トレースを読み取り、構造的なバグ (無視されたエラー、スキーマ違反、ループ) に証拠を付けてフラグを立てます。 LLM 審査員はいません。 · GitHub
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
アシュウィンウガレ
/
トレースリント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
40 コミット 40 コミット フォルダーとファイル
.github .github 例 例 src/tracelint src/tracelint te

sts テスト .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェント実行用のリンター - ツール呼び出しエージェントの実行トレースを読み取ります (
実際に行った) と、正確な証拠と CI を使用して、構造上のバグに決定論的にフラグを立てます。
終了コード。実行後に、コード上ではなくトレース上で実行され、2 番目のモデルは存在しません。
それを判断します。
tracelint はツール呼び出しエージェントのトレースを読み取り、構造上の欠陥を報告します (スキーマ違反)
ツール呼び出し、無視されたツールエラー、幻覚引数、ループ、冗長呼び出し - それぞれに
証拠として正確なトレース行を返し、CI 終了コードを返します。また、フォールト インジェクターと
障害ごとの回復スコアカード。
Model-as-Judge によるこれらの欠陥の検出は信頼性がありません (公開されているトレースエラー ベンチマークでは、低いことが示されています)
位置特定の精度）。これらの欠陥の多くは構造的に決定可能であり、判断する必要はありません。
これがこのツールの前提条件全体です。 2 番目のモデルがトレースを判断することはありません。
ライブデモレポートを見る - 構築された
検証スイート (すべての欠陥に対して 1 つの植え付けられたインスタンス、クリーンなコントロール、正当だが疑わしいもの)
ケース）に加えて、tracelint デモによって生成された、堅牢なリカバリ スコアカードとバグの多いリカバリ スコアカードを追加します。
決定論的なルールは、最終的な答えが正しいかどうかではなく、構造的な欠陥を捉えます。
構造的に問題がない限り、幻聴、ループ、冗長通話の所見が候補となります。
実証済み — 正当な値の変換と意図的な再試行によってトリップする可能性があります。それぞれが示されています
人間による審査のための証拠であり、評決として主張されることはありません。確信度の高い幻覚
検出には、

フィールドの原点 ( x-value-origin ) を宣言するためのツール スキーマ。
回復スコアカードには、ラベル付きのタスク結果 (成功のオラクル) が必要です。それらなしでは測定されます
動作の回復のみ (「クラッシュしなかった」) であり、正当性よりも弱い主張です。
トレースは、その計測によってのみ完全になります。必須フィールドが欠落しているルールは、
明示された理由で抑制されます。tracelint は、完全であるかのように部分的なトレースを lint することはありません。
デモでは、キーレス検証スイートとリカバリ スコアカードをエンドツーエンドで実行します。API キーも必要ありません。
モデルのダウンロード:
pip インストール トレースリント
トレースリントのデモ --htmlデモ.html
CI でトレースを lint します。
tracelint check ./trace.json --tools ./tools.json #hard_defect で終了 2
終了コード: 0 クリーン · 2 構造的に証明可能な欠陥 (hard_defect) · 3 入力エラー。
ヒューリスティック候補が単独で CI に失敗することはありません。抑制は開示されていますが、欠陥ではありません。
ルール
見つける
階層
R1
スキーマ違反 - 引数がツールの JSON スキーマに失敗します
ハードディフェクト
R2a
ツールがエラーを返しました
hard_event (構造化信号) / 候補 (ヒューリスティック)
R2b
エラーが発生した結果の値は、後の副作用のある呼び出しで再利用されます。
ハードディフェクト / 候補
R3
幻覚的な議論 — 出所から導き出せない価値
候補者;フィールドに注釈が付けられている場合は、hard_defect が提供されます
R4
ループ - N 個の同一の進行しない呼び出し (ポーリング/再試行は除く)
候補者
R5
冗長な呼び出し - 同一の呼び出し + 同一の結果、間に突然変異なし
候補者
R6
不正な形式の引数 — 発行されたツール呼び出し引数が有効な JSON ではありません
ハードディフェクト
R7
不明なツール - 宣言されたツールセットに存在しないツールへの呼び出し (幻覚ツールの可能性あり)
候補者
hard_event とhard_defect は検出の種類と直交しています。つまり、ツールエラーイベントは
構造化ステータスフィールドからのhard_eventですが、例外のような文字列からの候補です。
フリーフォ

rmコンテンツ。
トレースは JSON オブジェクト ( .json 、多くの場合 .jsonl ) です。
{
"run_id" : " run-1 " 、
「ステップ」: [
{ "type" : " message " 、 "role" : " user " 、 "content" : " 注文 4521 が発送されていない場合はキャンセルします " },
{ "type" : " tools_call " 、 "call_id" : " c1 " 、 "name" : " get_order_status " 、 "args" : { "order_id" : " 4521 " }},
{ "type" : " tools_result " 、 "call_id" : " c1 " 、 "content" : { "status" : " 処理中 " }、 "status" : " ok " },
{ "type" : " tools_call " 、 "call_id" : " c2 " 、 "name" : " cancel_order " 、
"args" : { "order_id" : " 4521 " 、 "reason" : " not_shipped " }}
]、
"final" : "注文 4521 はキャンセルされました。"
}
tools.json は、ルールがチェックするグラウンド トゥルースを提供します。
{
「ツール」: {
"注文のキャンセル" : {
"スキーマ" : { "タイプ" : " オブジェクト " 、 "プロパティ" : { "order_id" : { "タイプ" : " 文字列 " }}、
"必須" : [ "order_id " ]},
"metadata" : { "side_effecting" : true }
}
}
}
ツールは、その結果で障害がどのようなものであるかを宣言することもできるため、ドメイン障害は次のように返されます。
トランスポートの成功 ( {"status": "declined"} を運ぶ HTTP 200) は、代わりに構造的に捕捉されます。
すり抜ける:
{
「ツール」: {
"チャージカード" : {
「メタデータ」: {
"side_effecting" : true ,
"failure_when" : { "pointer" : " /status " , "in" : [ "拒否" , "失敗" ]}
}
}
}
}
Failure_when は、結果と一致 ( / に等しい / が存在する) への JSON ポインターです。試合
これは R2 の構造化エラーです (R2a にフィードし、副作用のある呼び出しに再利用すると R2b にフィードします)。あ
副作用のあるツールで、failure_when がなく、分類できない結果は抑制されます。
reason — never counted as a clean pass.
ルールは 1 つの正規トレース スキーマに対して実行されます。 a thin adapter translates each source's
フォーマットに変換するため、ルールが変わることはありません。組み込み: from_openai_messages (OpenAI チャット メッセージ
lists), from_langfuse_tra

ce (Langfuse トレースの観察)、および
from_otel_spans (OpenTelemetry / OpenInference —
普遍的な標準であるため、Arize Phoenix、OpenLLMetry、Langfuse-via-OTel、およびデータセットに到達します
1 つのベンダーだけではなく、TRAIL のように)。トレースをリントするには、examples/langfuse_cookbook.py を参照してください。
すでに Langfuse に収集されており、結果はスコアとして書き戻されます。
実際のトレースでは、アダプターは仕様だけでなくライブデータに対して検証されます。
実際の Langfuse v4 実行では from_langfuse_trace 、実際の実行では from_otel_spans
TRAIL ベンチマーク トレース (tracelint)
決定的に局所化された実際のツール エラー、不正な形式のツール呼び出し、過剰な再試行ループ
ループ内にモデルはありません。実際のエクスポートはさまざまであるため、新しいソースにはアダプターの小さな調整が必要になる場合があります。
ルールに必要なフィールドが存在しない場合、そのルールは推測ではなく抑制します (そう言う)。
未処理の癖は、間違った結果を生成するのではなく、安全に劣化します。さらに多くのアダプターは今後の課題です。
すでに収集したトレースをリントする
check はデフォルトでネイティブの tracelint JSON を読み取りますが、--format はトレースを直接指します。
スタックはすでに出力しています。手動によるスキーマ変換は必要ありません。
トレースリントチェックspans.json --format openinference # OTel/OpenInference: Phoenix、OTLP、TRAIL
tracelint checkmessages.json --format openai # OpenAI チャット メッセージ リスト
トレースlintチェックtrace.json --format langfuse # Langfuseトレースのエクスポート
ほとんどのルールにはツール スキーマは必要ないため、これはキーなしで機能します。 --tools tools.json を追加して、
スキーマ依存ルール (R1、および R3 の高信頼層)。マルチトレース入力 (.jsonl ファイル、
JSON 配列、または複数の Trace_id を含む OTLP エクスポート）は、それぞれ 1 つのレポートに展開されます。から
ライブラリ、同じワンライナー:
トレースリントインポートから lint_otel_trace
report = lint_otel_trace (spans) # spans: OpenInference スパンのエクスポート (リス

辞書の t)
print ( report . exit_code ) # 0 または 2
オフラインでのキーレスのエンドツーエンド実行 (フェニックスの形) については、examples/lint_openinference_phoenix.py を参照してください。
スパン → 検出結果 (ツール レジストリの有無にかかわらず)。
実行中の Arize Phoenix インスタンスから直接:
フェニックスをピクセルとしてインポート
トレースリントインポートから lint_otel_trace
スパン = ピクセル 。クライアント （）。 get_spans_dataframe()。 to_dict ( "レコード" )
print ( lint_otel_trace (spans).exit_code )
両方の Phoenix シェイプが処理されます。span-export JSON (トップレベルの scan_kind ) と
get_spans_dataframe() レコード (attributes.* 列としての属性)。
注入されたフォールトの下でエージェントがどのように動作するかを測定し、決定論的な成功のオラクルに対してスコアを付けます。
Tracelint スコアカード --demo --faults タイムアウト、エラー、rate_limit --runs 5
ベースラインは最初にオラクルを満たさなければなりません（そうでない場合、回復は測定されません）。各障害タイプのレポート
ウィルソン信頼区間による正しさ回復率。オラクルがない場合は、元に戻ります
行動の回復は弱いと分類されます。
from tracelint import lint_trace 、default_rules 、Trace 、ToolRegistry
トレース = トレース。ロード ( "trace.json" )
レジストリ = ToolRegistry 。ロード ( "tools.json" )
report = lint_trace (トレース、default_rules ()、レジストリ)
print ( report . exit_code ) # 0 または 2
レポートの f について。アクティブな調査結果:
print ( f . ルール 、 f . tier . value 、 f . summary )
開発
Python -m pytest
ruff チェック SRC テスト
コアは依存関係が軽く (jsonschema + stdlib)、テスト スイート全体は決定的であり、
オフライン。本物の OpenAI トレース生成エージェントは、オプトイン [real-agent] エクストラの背後に存在し、
決してリンターの一部ではありません。 Python 3.10 ～ 3.12。
エージェント実行のための決定論的リンター — 実行トレースを読み取り、構造的なバグ (無視されたエラー、スキーマ違反、ループ) に証拠を付けてフラグを立てます。 LLM 審査員はいません。
アシュウィヌガレ.github

.io/tracelint/ トピック
Readme MIT ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A deterministic linter for agent runs — reads the execution trace and flags structural bugs (ignored errors, schema violations, loops) with evidence. No LLM judge. - AshwinUgale/tracelint

GitHub - AshwinUgale/tracelint: A deterministic linter for agent runs — reads the execution trace and flags structural bugs (ignored errors, schema violations, loops) with evidence. No LLM judge. · GitHub
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
AshwinUgale
/
tracelint
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
40 Commits 40 Commits Folders and files
.github .github examples examples src/ tracelint src/ tracelint tests tests .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
A linter for agent runs — it reads the execution trace of a tool-calling agent (what it
actually did ) and flags structural bugs deterministically, with the exact evidence and a CI
exit code. It runs after the run, on the trace — not on your code — and no second model ever
judges it.
tracelint reads a tool-calling agent's trace and reports structural defects — schema-violating
tool calls, ignored tool errors, hallucinated arguments, loops, and redundant calls — each with the
exact trace lines as evidence, and returns a CI exit code. It also ships a fault injector and a
per-fault recovery scorecard.
Model-as-judge detection of these defects is unreliable (published trace-error benchmarks show low
localization accuracy). Many of these defects are structurally decidable and need no judge — that
is the entire premise of this tool. No second model ever judges the trace.
View the live demo report — the constructed
validation suite (one planted instance of every defect, clean controls, and legitimate-but-suspicious
cases) plus the robust-vs-buggy recovery scorecard, generated by tracelint demo .
Deterministic rules catch structural defects, not whether the final answer was correct.
Hallucinated-argument, loop, and redundant-call findings are candidates unless structurally
proven — legitimate value transforms and intentional retries can trip them; each is shown with
its evidence for human review, never asserted as a verdict. High-confidence hallucination
detection requires the tool schema to declare field origins ( x-value-origin ).
The recovery scorecard needs labeled task outcomes (success oracles); without them it measures
behavioral recovery only ("did not crash"), a weaker claim than correctness.
A trace is only as complete as its instrumentation. A rule whose required field is missing is
suppressed with a stated reason — tracelint never lints a partial trace as if complete.
The demo runs a keyless validation suite and a recovery scorecard end to end — no API key, no
model download:
pip install tracelint
tracelint demo --html demo.html
Lint a trace in CI:
tracelint check ./trace.json --tools ./tools.json # exit 2 on a hard_defect
Exit codes: 0 clean · 2 a structurally-provable defect ( hard_defect ) · 3 an input error.
Heuristic candidates never fail CI on their own; suppressions are disclosed but are not defects.
Rule
Finding
Tiers
R1
schema violation — args fail the tool's JSON Schema
hard_defect
R2a
tool returned an error
hard_event (structured signal) / candidate (heuristic)
R2b
an errored result's value reused by a later side-effecting call
hard_defect / candidate
R3
hallucinated argument — value not derivable from provenance
candidate ; hard_defect if the field is annotated provided
R4
loop — N identical no-progress calls (polls/retries excluded)
candidate
R5
redundant call — identical call + identical result, no mutation between
candidate
R6
malformed arguments — the emitted tool-call arguments are not valid JSON
hard_defect
R7
unknown tool — a call to a tool absent from the declared toolset (possible hallucinated tool)
candidate
hard_event and hard_defect are orthogonal to the finding kind: a tool-error event is a
hard_event from a structured status field but a candidate from an exception-like string in
free-form content.
A trace is a JSON object ( .json , or .jsonl for many):
{
"run_id" : " run-1 " ,
"steps" : [
{ "type" : " message " , "role" : " user " , "content" : " cancel order 4521 if it hasn't shipped " },
{ "type" : " tool_call " , "call_id" : " c1 " , "name" : " get_order_status " , "args" : { "order_id" : " 4521 " }},
{ "type" : " tool_result " , "call_id" : " c1 " , "content" : { "status" : " processing " }, "status" : " ok " },
{ "type" : " tool_call " , "call_id" : " c2 " , "name" : " cancel_order " ,
"args" : { "order_id" : " 4521 " , "reason" : " not_shipped " }}
],
"final" : " Order 4521 has been cancelled. "
}
tools.json supplies the ground truth the rules check against:
{
"tools" : {
"cancel_order" : {
"schema" : { "type" : " object " , "properties" : { "order_id" : { "type" : " string " }},
"required" : [ " order_id " ]},
"metadata" : { "side_effecting" : true }
}
}
}
A tool can also declare what failure looks like in its result, so a domain failure returned as
a transport success (HTTP 200 carrying {"status": "declined"} ) is caught structurally instead of
slipping through:
{
"tools" : {
"charge_card" : {
"metadata" : {
"side_effecting" : true ,
"failure_when" : { "pointer" : " /status " , "in" : [ " declined " , " failed " ]}
}
}
}
}
failure_when is a JSON Pointer into the result plus a match ( in / equals / exists ); a match
is a structured error for R2 (feeding R2a and, on reuse into a side-effecting call, R2b). A
side-effecting tool with no failure_when and an unclassifiable result is suppressed with a
reason — never counted as a clean pass.
The rules run against one canonical trace schema ; a thin adapter translates each source's
format into it, so the rules never change. Built in: from_openai_messages (OpenAI chat message
lists), from_langfuse_trace (a Langfuse trace's observations), and
from_otel_spans ( OpenTelemetry / OpenInference —
the universal standard, so it reaches Arize Phoenix, OpenLLMetry, Langfuse-via-OTel, and datasets
like TRAIL, not just one vendor). See examples/langfuse_cookbook.py to lint the traces you
already collect in Langfuse and write findings back as scores.
On real traces: the adapters are validated against live data, not just the spec —
from_langfuse_trace on real Langfuse v4 runs, and from_otel_spans on real
TRAIL benchmark traces, where tracelint
deterministically localized real tool errors, a malformed tool call, and excessive-retry loops
with no model in the loop. Real exports vary, so a new source may need a small adapter tweak — and
when a field a rule needs is absent, that rule suppresses (says so) rather than guessing, so an
unhandled quirk degrades safely instead of producing a wrong result. More adapters are future work.
Lint the traces you already collect
check reads native tracelint JSON by default, but --format points it straight at the traces
your stack already emits — no manual schema conversion:
tracelint check spans.json --format openinference # OTel/OpenInference: Phoenix, OTLP, TRAIL
tracelint check messages.json --format openai # an OpenAI chat message list
tracelint check trace.json --format langfuse # a Langfuse trace export
Most rules need no tool schemas, so this works keyless; add --tools tools.json to light up the
schema-dependent rules (R1, and R3's high-confidence tier). A multi-trace input (a .jsonl file, a
JSON array, or an OTLP export carrying several trace_id s) fans out to one report each. From the
library, the same one-liner:
from tracelint import lint_otel_trace
report = lint_otel_trace ( spans ) # spans: your OpenInference span export (a list of dicts)
print ( report . exit_code ) # 0 or 2
See examples/lint_openinference_phoenix.py for an offline, keyless end-to-end run (Phoenix-shaped
spans → findings, with and without a tool registry).
Straight from a running Arize Phoenix instance:
import phoenix as px
from tracelint import lint_otel_trace
spans = px . Client (). get_spans_dataframe (). to_dict ( "records" )
print ( lint_otel_trace ( spans ). exit_code )
Both Phoenix shapes are handled: the span-export JSON (top-level span_kind ) and the
get_spans_dataframe() records (attributes as attributes.* columns).
Measure how an agent behaves under injected faults, scored against deterministic success oracles:
tracelint scorecard --demo --faults timeout,error,rate_limit --runs 5
The baseline must satisfy the oracle first (else recovery is not measured). Each fault type reports
a correctness-recovery rate with a Wilson confidence interval; with no oracle it falls back to
behavioral recovery, labeled as weaker.
from tracelint import lint_trace , default_rules , Trace , ToolRegistry
trace = Trace . load ( "trace.json" )
registry = ToolRegistry . load ( "tools.json" )
report = lint_trace ( trace , default_rules (), registry )
print ( report . exit_code ) # 0 or 2
for f in report . active_findings :
print ( f . rule , f . tier . value , f . summary )
Development
python -m pytest
ruff check src tests
The core is dependency-light ( jsonschema + stdlib) and the whole test suite is deterministic and
offline. A real OpenAI trace-generating agent lives behind the opt-in [real-agent] extra and is
never part of the linter. Python 3.10–3.12.
A deterministic linter for agent runs — reads the execution trace and flags structural bugs (ignored errors, schema violations, loops) with evidence. No LLM judge.
ashwinugale.github.io/tracelint/ Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
