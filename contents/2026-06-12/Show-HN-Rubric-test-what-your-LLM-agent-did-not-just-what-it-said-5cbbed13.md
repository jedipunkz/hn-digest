---
source: "https://github.com/Kareem-Rashed/rubric-eval"
hn_url: "https://news.ycombinator.com/item?id=48509073"
title: "Show HN: Rubric – test what your LLM agent did, not just what it said"
article_title: "GitHub - Kareem-Rashed/rubric-eval: Independent framework to test, benchmark, and evaluate LLMs & AI agents locally. · GitHub"
author: "kareemrashed"
captured_at: "2026-06-12T21:37:38Z"
capture_tool: "hn-digest"
hn_id: 48509073
score: 1
comments: 0
posted_at: "2026-06-12T20:26:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Rubric – test what your LLM agent did, not just what it said

- HN: [48509073](https://news.ycombinator.com/item?id=48509073)
- Source: [github.com](https://github.com/Kareem-Rashed/rubric-eval)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T20:26:15Z

## Translation

タイトル: HN: ルーブリックを表示 – LLM エージェントが言ったことだけでなく、何をしたかをテストします
記事のタイトル: GitHub - Kareem-Rashed/rubric-eval: LLM と AI エージェントをローカルでテスト、ベンチマーク、評価するための独立したフレームワーク。 · GitHub
説明: LLM と AI エージェントをローカルでテスト、ベンチマーク、評価するための独立したフレームワーク。 - カリーム・ラシュド/ルーブリック・エヴァル

記事本文:
GitHub - Kareem-Rashed/rubric-eval: LLM と AI エージェントをローカルでテスト、ベンチマーク、評価するための独立したフレームワーク。 · GitHub
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
カリーム・ラシュド
/
ルーブリック評価
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
56 コミット 56 コミット .github .github docs docs 例 例 rubriceval robriceval testing testing .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md ROADMAP.md ROADMAP.md action.yml action.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM アプリのエージェント動作テスト。
エージェントが何を言ったかだけでなく、呼び出されたツール、引数、トレース、待ち時間など、エージェントが行ったことをテストします。出荷前に CI でリグレッションを検出します。
あなたのエージェントはすべての手動チェックに合格しました。その後、プロンプト調整が出荷され、lookup_order の呼び出しを静かに停止し、メモリから応答を開始しました。応答は依然として正常に見えます。最終出力のみを確認する文字列一致評価と LLM ジャッジは、それをキャッチできません。
ルーブリックは、どのツールがどのような引数でどのような順序で呼び出されたか、禁止されたツールが回避されたかどうか、推論トレースがどの程度きれいであったか、および実行速度などの動作をテストします。必要な依存関係はゼロ、完全にローカル、MIT。
pip install ルーブリック評価
LangGraph エージェントを 60 秒でテストする
コールバック、ラッパー、手動配線は必要ありません。 Rubric は、エージェントがすでに生成したメッセージからツール呼び出し、引数、出力、エラー、完全なトレース、遅延、およびトークンの使用状況を抽出します。
ルーブリックとしてruricevalをインポート
ランググラフから。事前構築済みインポート create_react_agent
エージェント = create_react_agent (モデル、ツール = [ lookup_order 、 create_ticket 、 send_email ])
レポート = ルーブリック評価する(
test_cases = ルーブリック 。 run_langgraph ( エージェント , シナリオ = [
ルーブリック 。エージェントシナリオ (
input = "私の注文番号 #ORD-9821 はどこですか?" 、
Expected_tools = [ "lookup_order" ],
）、
ルーブリック 。エージェントシナリオ (
input = "私のアカウントはロックされています、これは緊急です

エント。」、
Expected_tools = [ "create_ticket" ],
forbidden_tools = [ "send_email" ], # 発券システムをバイパスしてはなりません
）、
])、
メトリクス = [
ルーブリック 。 ToolCallAccuracy ()、# 適切なツール?禁止事項はありませんか？
ルーブリック 。 TraceQuality ()、# ループなし、ステップの予算内ですか?
ルーブリック 。 LatencyMetric ( max_ms = 3000 )、
]、
Output_html = "レポート.html" ,
Output_json = "レポート.json" ,
)
[2/2] アカウントがロックされています。緊急です。
❌ スコア: 0.667
✗ tools_call_accuracy: 0.000 — 予期したツールがありません: ['create_ticket'];禁止されたツールの呼び出し: ['send_email']
✓trace_quality: 1.000 — トレースはきれいに見えます。 3 つのステップが実行されました。
✓ レイテンシ: 1.000 — レイテンシ 1840 ミリ秒は予算内 (3000 ミリ秒) です。
すでにagent.invoke()からの結果を持っていますか? 1 回の電話:
ケース = ルーブリック。 from_langgraph ( result , Expected_tools = [ "lookup_order" ])
LangGraph にないですか? rubric.from_messages() は、OpenAI 形式のメッセージ リスト ( role / content / tools_calls ) を受け入れるため、生の OpenAI ツール呼び出しループでも動作します。 LangFuse と LangSmith のトレース エクスポートは、load_langfuse() /load_langsmith() 経由でロードされ、いつでも手動で AgentTestCase を構築できます。
Rubric には、すべての PR で評価を実行し、ベースラインとの差分を実行し、結果をコメントとして投稿する GitHub アクションが同梱されています (Codecov と同様ですが、エージェントの動作を対象としています)。
# .github/workflows/eval.yml
名前 : エージェント エヴァルス
上: [プルリクエスト]
権限:
プルリクエスト : 書き込み
ジョブ:
評価:
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- 使用: Kareem-Rashed/rubric-eval@v0.2.0
付き:
eval ファイル: evals/regression.py
ベースライン: evals/baseline.json
環境:
OPENAI_API_KEY : ${{ Secrets.OPENAI_API_KEY }}
PR コメントは次のようになります。
🧪 ルーブリック評価 — 🔻 1 回帰
ベースライン
現在
Δ
合格率
100.0% (12/12)
91.7% (11/12)
🔻 -0.08
平均スコア
0.96
0.89
🔻 -0.07
🔻再評価

セッション (1)
緊急 — アカウントがロックされています — 合格 → 不合格 (スコア 1.00 → 0.67)
tool_call_accuracy : 1.00 → 0.00 — 期待されるツールがありません: ['create_ticket'];禁止されたツールの呼び出し: ['send_email']
同じ diff をローカルおよび任意の CI システムで使用できます。
ルーブリック実行 evals/regression.py --output-json current.json
ルーブリック比較 current.json --baseline evals/baseline.json --fail-on-regression
ルーブリック比較フラグは、合格→不合格の回帰、合格中のテスト、修正されたテスト、新規/削除されたテストのスコアの低下を、不合格のメトリックの理由をインラインで表示します。
メトリック
小切手
ToolCallAccuracy()
予期されたツールの呼び出し、禁止されたツールの回避、オプションの順序チェック
トレース品質()
ループなし、ステップバジェット内
タスク完了()
タスクは実際に完了しました
ToolCallEfficiency()
冗長または無駄なツール呼び出しは不要
安全性コンプライアンス()
トレース内に危険なアクションはありません
ReasoningQuality()
一貫した複数ステップの推論
ContextUtilization()
提供されたコンテキストが実際に使用されました
LatencyMetric(max_ms=...)
レイテンシの予算内で
CostMetric(max_cost_usd=...)
コスト予算内で
出力品質
メトリック
小切手
LLMJudge(基準=...)
カスタム LLM ベースのスコアリング — あらゆる呼び出し可能 (OpenAI、Anthropic、Ollama) で動作します
GEval(名前=..., 基準=...)
思考連鎖による LLM 評価
幻覚スコア()
提供されたコンテキスト (LLM ジャッジまたは NLI モード) で出力を接地
SemanticSimilarity(しきい値=...)
埋め込みの類似性と期待される出力 ([セマンティック] 追加)
RougeScore()
ROUGE 重複まとめ（「ルージュ」番外編）
ExactMatch() / Contains() / NotContains() / RegexMatch()
文字列チェック、依存関係なし
LLM 判定メトリクスは不安定性検出による繰り返し実行をサポートします。ルーブリックはスコアの差異をレポートするため、エージェントではなく判定者がいつ不安定な部分であるかを知ることができます。
fromruriceval import BaseMetric 、 MetricResult

クラス NoApologySpam (BaseMetric):
名前 = "no_apology_spam"
しきい値 = 1.0
def メジャー ( self , test_case ) -> MetricResult :
カウント = テストケース 。実際の出力 。より低い （）。数えます (「ごめんなさい」)
MetricResult を返す (
metric_name = self 。名前、
スコア = 1.0 if count <= 1 else 0.0 、
渡された = カウント <= 1 、
reason = f「「申し訳ありません」が { count } 回表示されます。」 、
)
pytestの統合
組み込みのrubric_evalフィクスチャを介して、通常のテストとして評価します。
def test_agent_routes_urgent_requests (rubric_eval):
結果 = エージェント。 invoke ({ "messages" : [{ "role" : "user" , "content" : "アカウントがロックされています。緊急です!" }]})
ルーブリック_評価 。追加(
ルーブリック 。 from_langgraph ( 結果 ,
Expected_tools = [ "create_ticket" ],
禁じられたツール = [ "send_email" ])、
メトリクス = [ルーブリック . ToolCallAccuracy ()]、
)
# テスト終了時に自動アサート
CLI
rubric run evals/regression.py # eval ファイルを実行する
ルーブリック実行 evals/regression.py --output-html report.html --output-json report.json
ルーブリック実行 evals/regression.py --quiet --fail-on-error
ルーブリック比較 current.json --baseline Baseline.json --fail-on-regression
ルーブリックバージョン
HTML レポートは、テストごとのトレース、ツール呼び出し、メトリックごとの内訳を含む単一の自己完結型ファイルです。ローカルで開いて CI アーティファクトに添付すれば、サーバーは必要ありません。
行動第一。ほとんどの eval フレームワークは、最終的な答えをスコア付けします。 Rubric の中核となる抽象化は、エージェントの実行 (ツール、引数、トレース、レイテンシー) です。これは、エージェントのバグが実際に存在する場所だからです。
配線ゼロ。 run_langgraph() / from_messages() は、既存のメッセージをテスト ケースに変換します。アプリを実行するための SDK はありません。
CIネイティブ。ベースラインの差分、PR コメント、および終了コードは、有料のアドオンではなく組み込まれています。
依存関係は必要なく、完全にローカルです。 pip installrubric-eval は他に何も取り込みません。あなたのプロンプトと痕跡があなたの母から離れることはありません

中国。
独立系でMITライセンスを取得しています。 AI 企業やプラットフォーム ベンダーによって所有されていないため、評価を誰かのクラウド経由でルーティングするというプレッシャーはありません。
Examples/langgraph_eval.py — エージェントの動作をエンドツーエンドでテストします (deps なし、API キーなしで実行)
Examples/eval.py — 本番環境に即したスイート: FAQ ボット + サポート エージェント
ROADMAP.md を参照してください。次は、CrewAI、OpenAI Agents SDK、MCP サーバーの自動キャプチャです。マージ時のベースライン自動更新。データセットローダー。
寄稿は歓迎です。「良好な最初の号」とタグ付けされた号は、純粋に最初の PR に限定されています。
git clone https://github.com/Kareem-Rashed/rubric-eval
cd ルーブリック評価
pip install -e " .[dev] "
pytest テスト/
ガイドラインについては、CONTRIBUTING.md を参照してください。
LLM と AI エージェントをローカルでテスト、ベンチマーク、評価するための独立したフレームワーク。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
2
v0.2.0 — エージェントの動作テスト
最新の
2026 年 6 月 12 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Independent framework to test, benchmark, and evaluate LLMs & AI agents locally. - Kareem-Rashed/rubric-eval

GitHub - Kareem-Rashed/rubric-eval: Independent framework to test, benchmark, and evaluate LLMs & AI agents locally. · GitHub
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
Kareem-Rashed
/
rubric-eval
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
56 Commits 56 Commits .github .github docs docs examples examples rubriceval rubriceval tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md ROADMAP.md ROADMAP.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
Agent behavior testing for LLM apps.
Test what your agent did — tools called, arguments, trace, latency — not just what it said. Catch regressions in CI before they ship.
Your agent passed every manual check. Then a prompt tweak shipped, and it quietly stopped calling lookup_order and started answering from memory. The responses still look fine — string-match evals and LLM judges that only see the final output can't catch it.
Rubric tests the behavior : which tools were called, with what arguments, in what order, whether forbidden tools were avoided, how clean the reasoning trace was, and how fast it ran. Zero required dependencies, fully local, MIT.
pip install rubric-eval
Test a LangGraph agent in 60 seconds
No callbacks, no wrappers, no manual wiring. Rubric extracts tool calls, arguments, outputs, errors, the full trace, latency, and token usage from the messages your agent already produces.
import rubriceval as rubric
from langgraph . prebuilt import create_react_agent
agent = create_react_agent ( model , tools = [ lookup_order , create_ticket , send_email ])
report = rubric . evaluate (
test_cases = rubric . run_langgraph ( agent , scenarios = [
rubric . AgentScenario (
input = "Where is my order #ORD-9821?" ,
expected_tools = [ "lookup_order" ],
),
rubric . AgentScenario (
input = "My account is locked, this is urgent." ,
expected_tools = [ "create_ticket" ],
forbidden_tools = [ "send_email" ], # must not bypass the ticketing system
),
]),
metrics = [
rubric . ToolCallAccuracy (), # right tools? no forbidden ones?
rubric . TraceQuality (), # no loops, within step budget?
rubric . LatencyMetric ( max_ms = 3000 ),
],
output_html = "report.html" ,
output_json = "report.json" ,
)
[2/2] My account is locked, this is urgent.
❌ Score: 0.667
✗ tool_call_accuracy: 0.000 — Missing expected tools: ['create_ticket']; Called forbidden tools: ['send_email']
✓ trace_quality: 1.000 — Trace looks clean. 3 steps taken.
✓ latency: 1.000 — Latency 1840ms is within budget (3000ms).
Already have a result from agent.invoke() ? One call:
case = rubric . from_langgraph ( result , expected_tools = [ "lookup_order" ])
Not on LangGraph? rubric.from_messages() accepts any OpenAI-format message list ( role / content / tool_calls ), so it works with raw OpenAI tool-calling loops too. LangFuse and LangSmith trace exports load via load_langfuse() / load_langsmith() , and you can always construct an AgentTestCase by hand.
Rubric ships a GitHub Action that runs your evals on every PR, diffs against a baseline, and posts the result as a comment — like Codecov, but for agent behavior.
# .github/workflows/eval.yml
name : Agent Evals
on : [pull_request]
permissions :
pull-requests : write
jobs :
eval :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : Kareem-Rashed/rubric-eval@v0.2.0
with :
eval-file : evals/regression.py
baseline : evals/baseline.json
env :
OPENAI_API_KEY : ${{ secrets.OPENAI_API_KEY }}
The PR comment looks like this:
🧪 Rubric eval — 🔻 1 regression
Baseline
Current
Δ
Pass rate
100.0% (12/12)
91.7% (11/12)
🔻 -0.08
Avg score
0.96
0.89
🔻 -0.07
🔻 Regressions (1)
Urgent — account locked — pass → fail (score 1.00 → 0.67)
tool_call_accuracy : 1.00 → 0.00 — Missing expected tools: ['create_ticket']; Called forbidden tools: ['send_email']
The same diff is available locally and in any CI system:
rubric run evals/regression.py --output-json current.json
rubric compare current.json --baseline evals/baseline.json --fail-on-regression
rubric compare flags pass→fail regressions, score drops on still-passing tests, fixed tests, and new/removed tests — with the failing metric's reason inline.
Metric
Checks
ToolCallAccuracy()
Expected tools called, forbidden tools avoided, optional order check
TraceQuality()
No loops, within step budget
TaskCompletion()
The task was actually finished
ToolCallEfficiency()
No redundant or wasted tool calls
SafetyCompliance()
No unsafe actions in the trace
ReasoningQuality()
Coherent multi-step reasoning
ContextUtilization()
Provided context was actually used
LatencyMetric(max_ms=...)
Within latency budget
CostMetric(max_cost_usd=...)
Within cost budget
Output quality
Metric
Checks
LLMJudge(criteria=...)
Custom LLM-based scoring — works with any callable (OpenAI, Anthropic, Ollama)
GEval(name=..., criteria=...)
Chain-of-thought LLM evaluation
HallucinationScore()
Output grounded in the provided context (LLM judge or NLI mode)
SemanticSimilarity(threshold=...)
Embedding similarity vs expected output ( [semantic] extra)
RougeScore()
ROUGE overlap for summarization ( [rouge] extra)
ExactMatch() / Contains() / NotContains() / RegexMatch()
String checks, zero dependencies
LLM-judge metrics support repeated runs with flakiness detection — Rubric reports the score variance so you know when your judge, not your agent, is the unstable part.
from rubriceval import BaseMetric , MetricResult
class NoApologySpam ( BaseMetric ):
name = "no_apology_spam"
threshold = 1.0
def measure ( self , test_case ) -> MetricResult :
count = test_case . actual_output . lower (). count ( "sorry" )
return MetricResult (
metric_name = self . name ,
score = 1.0 if count <= 1 else 0.0 ,
passed = count <= 1 ,
reason = f"'sorry' appears { count } time(s)." ,
)
pytest integration
Evals as regular tests, via the built-in rubric_eval fixture:
def test_agent_routes_urgent_requests ( rubric_eval ):
result = agent . invoke ({ "messages" : [{ "role" : "user" , "content" : "Account locked, urgent!" }]})
rubric_eval . add (
rubric . from_langgraph ( result ,
expected_tools = [ "create_ticket" ],
forbidden_tools = [ "send_email" ]),
metrics = [ rubric . ToolCallAccuracy ()],
)
# auto-asserts at end of test
CLI
rubric run evals/regression.py # run an eval file
rubric run evals/regression.py --output-html report.html --output-json report.json
rubric run evals/regression.py --quiet --fail-on-error
rubric compare current.json --baseline baseline.json --fail-on-regression
rubric version
The HTML report is a single self-contained file with per-test traces, tool calls, and per-metric breakdowns — open it locally, attach it to CI artifacts, no server needed.
Behavior-first. Most eval frameworks score the final answer. Rubric's core abstraction is the agent run — tools, arguments, trace, latency — because that's where agent bugs actually live.
Zero wiring. run_langgraph() / from_messages() turn the messages you already have into test cases. No SDK to thread through your app.
CI-native. Baseline diffing, PR comments, and exit codes are built in, not a paid add-on.
Zero required dependencies, fully local. pip install rubric-eval pulls in nothing else. Your prompts and traces never leave your machine.
Independent and MIT-licensed. Not owned by an AI company or a platform vendor — no pressure to route your evals through anyone's cloud.
examples/langgraph_eval.py — agent behavior testing end-to-end (runs with zero deps, no API keys)
examples/eval.py — a production-realistic suite: FAQ bot + support agent
See ROADMAP.md . Next up: auto-capture for CrewAI, the OpenAI Agents SDK, and MCP servers; baseline auto-update on merge; dataset loaders.
Contributions are welcome — the issues tagged good first issue are genuinely scoped to a first PR.
git clone https://github.com/Kareem-Rashed/rubric-eval
cd rubric-eval
pip install -e " .[dev] "
pytest tests/
See CONTRIBUTING.md for guidelines.
Independent framework to test, benchmark, and evaluate LLMs & AI agents locally.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
2
v0.2.0 — Agent behavior testing
Latest
Jun 12, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
