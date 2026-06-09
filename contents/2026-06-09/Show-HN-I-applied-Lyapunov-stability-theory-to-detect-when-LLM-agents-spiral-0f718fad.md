---
source: "https://github.com/vishal-dehurdle/state-harness"
hn_url: "https://news.ycombinator.com/item?id=48466381"
title: "Show HN: I applied Lyapunov stability theory to detect when LLM agents spiral"
article_title: "GitHub - vishal-dehurdle/state-harness: Runtime safety net for LLM agents. Detects token spirals, kills doomed tasks early, tells you exactly why. Rust core, Python SDK. pip install state-harness · GitHub"
author: "visha1v"
captured_at: "2026-06-09T21:39:30Z"
capture_tool: "hn-digest"
hn_id: 48466381
score: 2
comments: 1
posted_at: "2026-06-09T19:27:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I applied Lyapunov stability theory to detect when LLM agents spiral

- HN: [48466381](https://news.ycombinator.com/item?id=48466381)
- Source: [github.com](https://github.com/vishal-dehurdle/state-harness)
- Score: 2
- Comments: 1
- Posted: 2026-06-09T19:27:10Z

## Translation

タイトル: Show HN: LLM エージェントのスパイラルを検出するためにリアプノフ安定理論を適用しました
記事のタイトル: GitHub - vishal-dehurdle/state-harness: LLM エージェントのランタイム セーフティ ネット。トークンスパイラルを検出し、運命のタスクを早期に終了し、正確な理由を示します。 Rust コア、Python SDK。 pip install state-harness · GitHub
説明: LLM エージェントのランタイム セーフティ ネット。トークンスパイラルを検出し、運命のタスクを早期に終了し、正確な理由を示します。 Rust コア、Python SDK。 pip install state-harness - vishal-dehurdle/state-harness

記事本文:
GitHub - vishal-dehurdle/state-harness: LLM エージェントのランタイム セーフティ ネット。トークンスパイラルを検出し、運命のタスクを早期に終了し、正確な理由を示します。 Rust コア、Python SDK。 pip install state-harness · GitHub
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
ヴィシャル・デハードル
/
ステートハーネス
パブ

リック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット .github .github ベンチマーク ベンチマークの例 例 python/ state_harness python/ state_harness スクリプト スクリプト src src tau3_integration tau3_integration テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md COTRIBUTING.md Cargo.toml Cargo.toml LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM エージェントのランタイム セーフティ ネット。物事がうまくいくときは何もしません。予算を節約し、そうでない場合はその理由を正確に伝えます。
from state_harness import GrowthRatioGuard 、FailureReport
ガード = GrowthRatioGuard ( token_budget = 50_000 )
ガード付き：
Agent_loop の順番:
結果 = llm 。 invoke (ターンプロンプト)
ガード。 Record_step ( tokens_used = 結果 . 使用量 . total_tokens )
＃何が問題だったのでしょうか？ (コストゼロ、LLM コールなし)
レポート = FailureReport 。 from_guard (ガード)
印刷（レポート）
⚠️ターン12で安定性が低下
パターン: コンテキスト蓄積スパイラル (信頼度: 92%)
• 最後の 5 ターンはすべてベースラインの 1.5 倍を超えました (4/4 が加速していました)。
• ピーク成長率: ベースラインの 5.2 倍。
• 介入なしの場合、予測コストは 0.0396 ドル (実際: 0.0039 ドル) でした。
エネルギー: ▂▂▃▄▆█
ベースライン: 1050 トークン/ターン
ピーク比: 5.2× ベースライン
コスト: $0.0039 (早期トリップにより ~$0.0357 節約)
推奨されるアクション:
🔴 1. エージェント ループで RG 履歴圧縮を有効にします。
→ 古いメッセージを圧縮すると、プロンプト トークンが 40 ～ 60% 削減されます。
🟡 2. 成長率のしきい値を 1.8 倍に下げます。
→ 閾値が低ければ耳に入っただろう

嘘つき。
🟢 3. スライディング ウィンドウ コンテキスト戦略を追加します。
→ 最新の N 個のメッセージと以前のメッセージの概要のみを送信します。
なぜこれが存在するのか
運用環境で LLM エージェントを実行しているすべてのチームは、このような状況を経験したことがあります。エージェントがループに陥り、トークン使用量がスパイラルになり、翌朝には 1 つの失敗したリクエストに対して 15 ドルの請求が発生することに気付きます。プロセスを強制終了しますが、なぜそれが起こったのか、次回それを防ぐ方法はわかりません。
予算の上限を厳格に設定すれば、コストの問題は解決しますが、何もわかりません。タスクが強制終了されたことがわかります。それがコンテキスト蓄積スパイラルなのか、再試行の嵐なのか、ポリシーのドリフトなのかはわかりません。診断できないものは修正できません。
ステートハーネスはプラットフォームではなくライブラリです。 pip インストールして実行します。リャプノフ安定理論を使用して、暴走行為が高価になる前に検出します。暴走行為が発生した場合は、故障パターンを分類し、何が問題だったのか、修正方法、およびどれくらい節約できたのかを正確に通知します。すべてゼロコストで、追加の LLM 呼び出しや外部 API は必要ありません。
パターン
信号
例
コンテキストスパイラル
トークンの成長がベースラインを超えて加速
エージェントは毎ターン完全な履歴を再生します
リトライストーム
低分散の繰り返し呼び出し
ツールが失敗し、エージェントが同様に再試行します
ポリシーのドリフト
VSA 類似性スコアの低下
エージェントが会話中に本題から逸れてしまう
早期爆発
最初の 3 ターンでトークンが急増
特大のシステム プロンプトまたはツールの応答
予算の枯渇
累計支出額が上限に達する
複雑なタスク、必ずしも壊れているわけではない
得られるものと得られないもの
✅ エージェントが失敗した理由を知る
パターン分類 + 証拠 + 修正提案 — LLM コストゼロ
✅ 失敗したタスクのコンピューティングを節約
SWE ベンチの検索ノードが 38.6% 減少
✅ 健全なエージェントには決して干渉しないでください
1,886 回の短期/中ループ実行で誤検知ゼロ
✅ 3,175 回の実行で検証済み
4 つのベンチマーク、5-

状態アブレーション、ブートストラップ CI を使用した複数のトライアル
✅ モデルに依存しない
7 つのモデルで誤検知ゼロを確認: GPT-4o-mini、Claude Haiku 4.5、Gemini 2.5 Flash + Llama 3.2:3B、Phi-4-Mini、Qwen3:4B、Gemma4:E4B (Ollama)
❌ エージェントが賢くなるわけではありません
解決率はモニタリングの有無にかかわらず統計的に同一です
❌ 予算の上限に代わるものではありません
素朴な上限は同等の成功率を達成しますが、何も伝えません
値は診断です。予算の上限は「タスクが無効になった」ことを示します。ステートハーネスは、「コンテキスト蓄積スパイラルが原因でタスクが強制終了されました。これを修正するには履歴圧縮を有効にしてください」と通知します。その違いがこれが存在する理由です。
サーチ ツリー エージェント (MCTS、ビーム サーチ) を実行しているチーム — SWE ベンチ ソルバーや Devin などのツールの背後にあるアーキテクチャ。ループではなく分岐がコストを左右します。ブランチごとの反復上限は単独では問題ないように見えます。ツリーレベルのコスト爆発は静かに起こります。
プラットフォーム チームは 1 日あたり 1,000 以上のエージェント タスクを実行しています。手動によるトレース検査は拡張できません。ステートハーネスは、エッジで障害パターンを分類し (コストゼロ、LLM 呼び出しなし)、それらを集計分析用の OpenTelemetry 属性としてエクスポートします。
研究者らはエージェントのベンチマークを行っています。非決定性の下限 (Gemini 2.5 フラッシュでは標準偏差約 4 ～ 5%) は、デルタが 8% 未満の単一実行比較がノイズであることを意味します。ステートハーネスはこれを定量化します。
チャットボット、RAG パイプライン、またはシングルターン アプリ - これらはスパイラルにはなりません。監視する必要はありません。
10 ターン未満の単純な ReAct ループ — max_iterations=10 とバジェット キャップで十分です。最新のフレームワーク (LangGraph、CrewAI) はすべて、これをネイティブにサポートしています。
pip インストール状態ハーネス
Python 3.10 以上が必要です。事前に構築されたホイールは、Linux、macOS、および Windows (x86_64 および ARM64) で使用できます。 Rust ツールチェーンは必要ありません。
git clone https://github.com/vishal-dehurdle/state-harnes

s.git
CD ステートハーネス
python -m venv .venv && ソース .venv/bin/activate
# Rustをインストールします(まだインストールされていない場合)
カール --proto ' =https ' --tlsv1.2 -sSf https://sh.rustup.rs |しー
pip インストール マチュリン
maturin 開発 --release
# テストを実行する
pip インストール pytest
pytest テスト/
クイックスタート
基本: GrowthRatioGuard (推奨)
GrowthRatioGuard はトークンの使用量をベースラインに対して正規化するため、マルチターン コンテキスト ウィンドウの自然な増加ではなく、不均衡な増加が発生した場合にのみトリップします。
from state_harness import GrowthRatioGuard 、 StabilityViolation
ガード = GrowthRatioGuard (
token_budget = 100_000 、# ハード上限
rate_threshold = 2.0 、 # ターンがベースラインの 2 倍のときにトリップする
window = 3 、 # トリップまでの 3 回連続のエスカレーション ターン
Budget_gate = 8_000 、 # 8K トークンが消費されるまでトリップしない
)
ガード付き：
Agent_loop の順番:
試してみてください:
結果 = llm 。 invoke (ターンプロンプト)
ガード。レコードステップ (
tokens_used = 結果 。使用方法。 total_tokens 、
エラー = 0 、
)
e としての StabilityViolation を除く:
print ( f"エージェントが殺されました: { e } " )
休憩
print ( f"総コスト: {guard . total_tokens } トークン" )
print ( f"ベースライン: { ガード . ベースライン } トークン/ターン" )
print ( f"ピーク比: {guard . current_ratio } ×" )
故障診断
実行後 (トリップしたかどうかに関係なく)、構造化された障害レポートを取得します。
state_harness インポート失敗レポートから
レポート = FailureReport 。 from_guard (ガード、モデル = "gemini-2.5-フラッシュ")
# 人間が判読できる端末出力
印刷（レポート）
# ロギング/ダッシュボード用の構造化辞書
jsonをインポートする
print ( json . dumps ( report . to_dict (), indent = 2 ))
このレポートでは、障害パターンの分類、証拠の提供、コストへの影響の推定、および特定の修正の提案がすべて LLM 呼び出しなしで行われます。
未加工のトークン数を使用した下位レベルの制御の場合 (正規化なし):
state_harness imp から

オルトバウンダリーガード
BoundaryGuard (token_budget = 100_000、lambda_ = 1.0、window = 5) をガードとして使用します。
Agent_loop の順番:
結果 = llm 。 invoke (ターンプロンプト)
ガード。レコードステップ (
tokens_used = 結果 。使用方法。 total_tokens 、
エラー = 0 、
ツール名 = "検索" ,
)
デコレータ: @boundary_guard
state_harness からのインポート境界ガード
@ 境界ガード (
token_budget = 50_000 、
token_counter = ラムダ r : r 。使用方法。 total_tokens 、
)
def エージェント_ステップ (プロンプト : str ):
llm を返します。呼び出し ( プロンプト )
フレームワークの統合
ランググラフから。事前構築済みインポート create_react_agent
state_harness から。アダプターのインポートmonitor_graph
エージェント = create_react_agent (モデル、ツール = [検索、計算])
安全 = モニターグラフ (エージェント、トークン予算 = 100_000)
結果 = 安全。 invoke ({ "メッセージ" : [( "ユーザー" , "ログインのバグを修正する" )]})
# 実行後 — 常に利用可能:
print (safe . total_tokens ) # 累積使用量
print (safe.tripped) # 安定性がトリップしましたか?
print ( 安全なレポート ) # パターン + 提案を含む完全な FailureReport
ストリーミングの場合:
安全なチャンク用。 stream ({ "メッセージ" : [( "ユーザー" , "このモジュールをリファクタリング" )]}):
印刷 (チャンク)
旅行コールバックの場合 (Slack アラートなど):
安全 = モニターグラフ (
代理店、
token_budget = 100_000 、
on_trip = ラムダレポート: スラック 。 send ( f"エージェントがトリップしました: { レポート . パターン } " )、
)
上級: LangGraphMiddleware を使用したツールごとのラッピング
state_harness から BoundaryGuard をインポート
state_harness から。アダプターのインポート LangGraphMiddleware
ガード = BoundaryGuard ( token_budget = 150_000 )
ミドルウェア = LangGraphMiddleware (ガード)
@ミドルウェア。ラップツール
def search_database (クエリ: str):
データベースを返します。検索（クエリ）
ガード付き：
結果 = エージェント。 invoke ({ "メッセージ" : [...]})
CrewAI
from crewai import エージェント、タスク、クルー
州ハーネより

ss 。アダプターのインポート CrewAICallback
コールバック = CrewAICallback ( token_budget = 200_000 )
乗組員 = 乗組員 (
エージェント = [ 研究者 、 ライター ]、
タスク = [ リサーチタスク , ライトタスク ],
step_callback = コールバック 。 step_callback 、
task_callback = コールバック 。 task_callback 、
)
結果 = 乗組員。キックオフ（）
print (コールバック . レポート) # FailureReport
コールバック。閉じる（）
バニラ パイソン フック
state_harness から BoundaryGuard をインポート
state_harness から。アダプターは VanillaHook をインポートします
ガード = BoundaryGuard ( token_budget = 50_000 )
フック = VanillaHook (ガード)
ガード付き：
Agent_loop のステップの場合:
フック。 before_call ( ツール名 = "検索" )
結果 = 実行ツール (ステップ)
フック。 after_call ( tokens_used = result . tokens )
CLI
# トークンの軌道をシミュレートします — ガードが何をするか見てみましょう
ステート ハーネス シミュレート 1000 1200 1500 2000 3000 5000 8000 --budget 50000
# 保存されたレポートを分析する
ステートハーネス分析レポート.json
state-harness 分析 report.json --json # JSON 出力
state-harness 分析 report.json --otel # OpenTelemetry 属性
# ディレクトリ内のすべてのレポートをバッチ分析します
状態ハーネス バッチ --dir ./reports/ --output results.csv
構造化された出力
すべての FailureReport は、複数の出力形式をサポートしています。
レポート = FailureReport 。 from_guard (ガード)
# JSON (ロギング、API、ストレージ用)
ご報告をさせていただきます。 to_json () # きれいに印刷された

[切り捨てられた]

## Original Extract

Runtime safety net for LLM agents. Detects token spirals, kills doomed tasks early, tells you exactly why. Rust core, Python SDK. pip install state-harness - vishal-dehurdle/state-harness

GitHub - vishal-dehurdle/state-harness: Runtime safety net for LLM agents. Detects token spirals, kills doomed tasks early, tells you exactly why. Rust core, Python SDK. pip install state-harness · GitHub
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
vishal-dehurdle
/
state-harness
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits .github .github benchmarks benchmarks examples examples python/ state_harness python/ state_harness scripts scripts src src tau3_integration tau3_integration tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.toml Cargo.toml LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json View all files Repository files navigation
Runtime safety net for LLM agents. Does nothing when things work. Saves your budget and tells you exactly why when they don't.
from state_harness import GrowthRatioGuard , FailureReport
guard = GrowthRatioGuard ( token_budget = 50_000 )
with guard :
for turn in agent_loop :
result = llm . invoke ( turn . prompt )
guard . record_step ( tokens_used = result . usage . total_tokens )
# What went wrong? (zero-cost, no LLM calls)
report = FailureReport . from_guard ( guard )
print ( report )
⚠️ STABILITY TRIPPED at turn 12
Pattern: Context Accumulation Spiral (confidence: 92%)
• Last 5 turns all exceeded 1.5× baseline (4/4 were accelerating).
• Peak growth ratio: 5.2× baseline.
• Without intervention, projected cost was $0.0396 (actual: $0.0039).
Energy: ▁▁▁▁▁▂▂▃▄▆█
Baseline: 1050 tokens/turn
Peak ratio: 5.2× baseline
Cost: $0.0039 (saved ~$0.0357 by tripping early)
Suggested actions:
🔴 1. Enable RG history compression in your agent loop.
→ Compressing older messages reduces prompt tokens by 40-60%.
🟡 2. Lower the growth ratio threshold to 1.8×.
→ A lower threshold would have caught it earlier.
🟢 3. Add a sliding-window context strategy.
→ Send only the last N messages plus a summary of earlier ones.
Why this exists
Every team running LLM agents in production has experienced this: an agent gets stuck in a loop, token usage spirals, and you find a $15 charge for a single failed request the next morning. You kill the process — but you have no idea why it happened or how to prevent it next time.
A hard budget cap solves the cost problem — but tells you nothing. You know the task was killed. You don't know if it was a context accumulation spiral, a retry storm, or policy drift. You can't fix what you can't diagnose.
State-harness is a library , not a platform. pip install and go. It uses Lyapunov stability theory to detect runaway behavior before it becomes expensive — and when it trips, it classifies the failure pattern and tells you exactly what went wrong, how to fix it, and how much you saved. All at zero cost — no extra LLM calls, no external APIs.
Pattern
Signal
Example
Context Spiral
Token growth accelerating beyond baseline
Agent replaying full history each turn
Retry Storm
Low-variance repeated calls
Tool failing, agent retrying identically
Policy Drift
VSA similarity score dropping
Agent going off-topic mid-conversation
Early Explosion
Token spike in first 3 turns
Oversized system prompt or tool response
Budget Exhaustion
Cumulative spend hits ceiling
Complex task, not necessarily broken
What you get — and what you don't
✅ Know WHY your agent failed
Pattern classification + evidence + fix suggestions — zero LLM cost
✅ Save compute on failing tasks
38.6% fewer search nodes on SWE-bench
✅ Never interfere with healthy agents
Zero false positives across 1,886 short/medium-loop runs
✅ Validated across 3,175 runs
4 benchmarks, 5-condition ablation, multi-trial with bootstrap CIs
✅ Model-agnostic
Zero false positives confirmed across 7 models: GPT-4o-mini, Claude Haiku 4.5, Gemini 2.5 Flash + Llama 3.2:3B, Phi-4-Mini, Qwen3:4B, Gemma4:E4B (Ollama)
❌ Does NOT make your agent smarter
Resolve rates are statistically identical with or without monitoring
❌ Does NOT replace a budget cap
A naive cap achieves comparable success rates — but tells you nothing
The value is diagnostics. A budget cap tells you "task killed." State-harness tells you "task killed because of a context accumulation spiral — enable history compression to fix it." That difference is why this exists.
Teams running search-tree agents (MCTS, beam search) — the architecture behind SWE-bench solvers and tools like Devin. Branches, not loops, drive cost. A per-branch iteration cap looks fine in isolation; the tree-level cost explosion happens silently.
Platform teams running 1,000+ agent tasks/day — manual trace inspection doesn't scale. State-harness classifies failure patterns at the edge (zero cost, no LLM calls) and exports them as OpenTelemetry attributes for aggregate analysis.
Researchers benchmarking agents — the nondeterminism floor (~4–5% stdev on Gemini 2.5 Flash) means single-run comparisons with <8% delta are noise. State-harness quantifies this.
Chatbots, RAG pipelines, or single-turn apps — these don't spiral. You don't need monitoring.
Simple ReAct loops with <10 turns — max_iterations=10 and a budget cap are sufficient. Every modern framework (LangGraph, CrewAI) supports this natively.
pip install state-harness
Requires Python ≥ 3.10. Pre-built wheels are available for Linux, macOS, and Windows (x86_64 and ARM64). No Rust toolchain needed.
git clone https://github.com/vishal-dehurdle/state-harness.git
cd state-harness
python -m venv .venv && source .venv/bin/activate
# Install Rust (if not already installed)
curl --proto ' =https ' --tlsv1.2 -sSf https://sh.rustup.rs | sh
pip install maturin
maturin develop --release
# Run tests
pip install pytest
pytest tests/
Quickstart
Basic: GrowthRatioGuard (recommended)
The GrowthRatioGuard normalizes token usage against a baseline, so it only trips on disproportionate growth — not the natural growth of multi-turn context windows.
from state_harness import GrowthRatioGuard , StabilityViolation
guard = GrowthRatioGuard (
token_budget = 100_000 , # hard ceiling
ratio_threshold = 2.0 , # trip when turn is 2× the baseline
window = 3 , # 3 consecutive escalating turns to trip
budget_gate = 8_000 , # don't trip until 8K tokens spent
)
with guard :
for turn in agent_loop :
try :
result = llm . invoke ( turn . prompt )
guard . record_step (
tokens_used = result . usage . total_tokens ,
errors = 0 ,
)
except StabilityViolation as e :
print ( f"Agent killed: { e } " )
break
print ( f"Total cost: { guard . total_tokens } tokens" )
print ( f"Baseline: { guard . baseline } tokens/turn" )
print ( f"Peak ratio: { guard . current_ratio } ×" )
Failure Diagnostics
After any execution (tripped or not), get a structured failure report:
from state_harness import FailureReport
report = FailureReport . from_guard ( guard , model = "gemini-2.5-flash" )
# Human-readable terminal output
print ( report )
# Structured dict for logging / dashboards
import json
print ( json . dumps ( report . to_dict (), indent = 2 ))
The report classifies the failure pattern, provides evidence, estimates cost impact, and suggests specific fixes — all without any LLM calls.
For lower-level control using raw token counts (no normalization):
from state_harness import BoundaryGuard
with BoundaryGuard ( token_budget = 100_000 , lambda_ = 1.0 , window = 5 ) as guard :
for turn in agent_loop :
result = llm . invoke ( turn . prompt )
guard . record_step (
tokens_used = result . usage . total_tokens ,
errors = 0 ,
tool_name = "search" ,
)
Decorator: @boundary_guard
from state_harness import boundary_guard
@ boundary_guard (
token_budget = 50_000 ,
token_counter = lambda r : r . usage . total_tokens ,
)
def agent_step ( prompt : str ):
return llm . invoke ( prompt )
Framework Integration
from langgraph . prebuilt import create_react_agent
from state_harness . adapters import monitor_graph
agent = create_react_agent ( model , tools = [ search , calculate ])
safe = monitor_graph ( agent , token_budget = 100_000 )
result = safe . invoke ({ "messages" : [( "user" , "Fix the login bug" )]})
# After execution — always available:
print ( safe . total_tokens ) # cumulative usage
print ( safe . tripped ) # did stability trip?
print ( safe . report ) # full FailureReport with pattern + suggestions
For streaming:
for chunk in safe . stream ({ "messages" : [( "user" , "Refactor this module" )]}):
print ( chunk )
With a trip callback (e.g., for Slack alerts):
safe = monitor_graph (
agent ,
token_budget = 100_000 ,
on_trip = lambda report : slack . send ( f"Agent tripped: { report . pattern } " ),
)
Advanced: per-tool wrapping with LangGraphMiddleware
from state_harness import BoundaryGuard
from state_harness . adapters import LangGraphMiddleware
guard = BoundaryGuard ( token_budget = 150_000 )
middleware = LangGraphMiddleware ( guard )
@ middleware . wrap_tool
def search_database ( query : str ):
return db . search ( query )
with guard :
result = agent . invoke ({ "messages" : [...]})
CrewAI
from crewai import Agent , Task , Crew
from state_harness . adapters import CrewAICallback
callback = CrewAICallback ( token_budget = 200_000 )
crew = Crew (
agents = [ researcher , writer ],
tasks = [ research_task , write_task ],
step_callback = callback . step_callback ,
task_callback = callback . task_callback ,
)
result = crew . kickoff ()
print ( callback . report ) # FailureReport
callback . close ()
Vanilla Python Hooks
from state_harness import BoundaryGuard
from state_harness . adapters import VanillaHook
guard = BoundaryGuard ( token_budget = 50_000 )
hook = VanillaHook ( guard )
with guard :
for step in agent_loop :
hook . before_call ( tool_name = "search" )
result = execute_tool ( step )
hook . after_call ( tokens_used = result . tokens )
CLI
# Simulate a token trajectory — see what the guard would do
state-harness simulate 1000 1200 1500 2000 3000 5000 8000 --budget 50000
# Analyze a saved report
state-harness analyze report.json
state-harness analyze report.json --json # JSON output
state-harness analyze report.json --otel # OpenTelemetry attributes
# Batch analyze all reports in a directory
state-harness batch --dir ./reports/ --output results.csv
Structured Output
Every FailureReport supports multiple output formats:
report = FailureReport . from_guard ( guard )
# JSON (for logging, APIs, storage)
report . to_json () # pretty-printed

[truncated]
