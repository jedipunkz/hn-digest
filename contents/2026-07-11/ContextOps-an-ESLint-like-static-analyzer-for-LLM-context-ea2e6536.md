---
source: "https://github.com/Abhijeet777ui/contextops"
hn_url: "https://news.ycombinator.com/item?id=48876485"
title: "ContextOps, an ESLint-like static analyzer for LLM context"
article_title: "GitHub - Abhijeet777ui/contextops · GitHub"
author: "Abhijeet_Buag"
captured_at: "2026-07-11T22:37:49Z"
capture_tool: "hn-digest"
hn_id: 48876485
score: 1
comments: 0
posted_at: "2026-07-11T22:36:06Z"
tags:
  - hacker-news
  - translated
---

# ContextOps, an ESLint-like static analyzer for LLM context

- HN: [48876485](https://news.ycombinator.com/item?id=48876485)
- Source: [github.com](https://github.com/Abhijeet777ui/contextops)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T22:36:06Z

## Translation

タイトル: ContextOps、LLM コンテキスト用の ESLint のような静的アナライザー
記事のタイトル: GitHub - Abhijeet777ui/contextops · GitHub
説明: GitHub でアカウントを作成して、Abhijeet777ui/contextops の開発に貢献します。

記事本文:
GitHub - Abhijeet777ui/contextops · GitHub
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
アビジート777ui
/
コンテキストオプス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く Fo

フォルダとファイル
23 コミット 23 コミット .github/ workflows .github/ workflows ContextBench ContextBench ContextSecBench ContextSecBench ベンチマーク ベンチマーク contextops-0.2.2/ テスト contextops-0.2.2/ テスト contextops contextops 例 例 実験 実験 リーダーボード リーダーボード シナリオ シナリオ スクリプト スクリプト テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md HIGH_LEVEL_DOC.md HIGH_LEVEL_DOC.md HOW_TO_GET_JSON.md HOW_TO_GET_JSON.md ライセンス ライセンス README.md README.md STABILITY.md STABILITY.md USER_GUIDE.md USER_GUIDE.md 出力1.txt 出力1.txt 出力2.txt 出力2.txt 出力3.txt 出力3.txt pyproject.toml pyproject.toml test_duplicate.json test_duplicate.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM コンテキストの静的分析。
ContextOps は、LLM アプリケーション用の決定論的でモデルに依存しないコンテキスト リンターです。
推論前に LLM に送信されたコンテキストを分析し、冗長性、トークンの無駄、コンテキストの不均衡、ソースの集中などの構造的問題を検出します。埋め込み、モデル呼び出し、外部サービスを使用せずに、再現可能なコンテキスト ヘルス スコア (CHS) と実用的な診断を生成します。
ContextOps を LLM コンテキストの ESLint と考えてください。
現代のソフトウェア エンジニアリングには、決定論的な品質ゲートがあります。
コンパイラは構文エラーをキャッチします。
フォーマッタは一貫性を強制します。
静的アナライザーはアーキテクチャ上の問題を検出します。
LLM アプリケーションには、モデルに送信するコンテキストに相当するレイヤーが存在することはほとんどありません。
代わりに、プロンプトは時間の経過とともに静かに増加します。
これらの問題はレイテンシとコストを増加させ、モデルの動作を予測しにくくし、本番環境まで気付かれないことがよくあります。
ContextOps は、推論前にコンテキストの品質を観察、測定、テストできるようにします。
ContextOps が評価する

4つの構造次元。
結果は、決定的な 0 ～ 100 のコンテキスト ヘルス スコアと、詳細な調査結果と提案される修正です。
ContextOps は意図的に以下を評価しません。
推論能力または幻覚
推論前のコンテキストの構造的品質のみに焦点を当てます。
contextops は context.json を検査します
コンテキスト健全性スコア: 81 / 100
✓ 構造の複雑さが低い
✓ 優れたソースの多様性
警告
• 214 個の重複トークンが検出されました
• 検索がコンテキストの 78% を占める
• 2 つの取得チャンクが重複に近い
推定トークン節約量: 12%
より率直な診断を行うには、--roast モードを使用します。
contextops は context.json --roast を検査します
インストール
pip インストール contextops
インタラクティブなデモを実行します。
コンテキストオプスのデモ
依存関係: tiktoken とクリックのみ。 GPU、ネットワーク、外部 API キーは必要ありません。
LLM ペイロードをモデルに送信する前に保存してください。
jsonをインポートする
メッセージ = [
{ "役割" : "システム" 、 "コンテンツ" : system_prompt },
{ "ロール" : "ユーザー" 、 "コンテンツ" : user_query },
】
# API 呼び出しの前に 2 行を追加します
open ( "context.json" , "w" ) を f として使用:
json 。ダンプ ( メッセージ , f , インデント = 2 )
応答 = オープンアイ。チャット 。完成品。作成 (モデル = "gpt-4o" 、メッセージ = メッセージ )
または、構造化された dict 形式を使用して、より詳細な分析を行うこともできます。
{
"system" : " あなたは役立つカスタマー サポート ボットです。 " ,
「メッセージ」: [
{ "role" : " user " , "content" : " 返金にはどのくらい時間がかかりますか? " }
]、
"チャンク" : [
{ "content" : "返金には 3 ～ 5 営業日かかります。" , "source" : " docs/refunds.md " }
]、
"memory" : [ "ユーザーは昨日返金について問い合わせました。" ],
"tools" : [{ "name" : " search_api " , "output" : " ツールの応答テキストがここにあります " }]
}
2. 検査してください
contextops は context.json を検査します
3. CI での品質の強化
contextops チェック context.json --min-score 75
CLI リファレンス

コンテキスト ファイルを分析し、豊富なレポートを表示します。
contextops は context.json を検査します
contextops は context.json --roast を検査します
contextops は context.json を検査します --profile rag
contextops は context.json を検査します --explain
contextops は context.json --json-output を検査します
旗
デフォルト
説明
--json-出力
オフ
ターミナル形式ではなく生の JSON を出力する
--モデル <名前>
gpt-4o
トークンエンコーディングのモデル (例: gpt-4 、 claude-3 )
--profile <名前>
一般的な
アーキタイプ: rag 、 エージェント 、 チャットボット 、 ツールチェーン
--config <パス>
なし
カスタムしきい値を含む JSON 構成ファイルへのパス
--取得最大比 <f>
0.70
取得チャンクの最大許容比率をオーバーライドする
--system-max-ratio <f>
0.50
システムプロンプトの最大許容比率を上書きする
--memory-max-ratio <f>
0.50
メモリエントリの最大許容比率を上書きする
--tool-max-ratio <f>
0.60
ツール出力の最大許容比率をオーバーライドします
--説明してください
オフ
各ペナルティの詳細な「トップスコアドライバー」を表示
--ロースト
オフ
極めて正直なスコアバンドの解説を可能にする
contextops check <ファイル> --min-score N
CIゲート。コード 0 (合格) または 1 (失敗) で終了します。
contextops チェック context.json --min-score 75
終了コード
意味
0
PASS — スコアがしきい値を満たすか超えています
1
FAIL — スコアがしきい値を下回っているか、分析エラーです
2
エラー - 無効な JSON または読み取り不可能な入力
GitHub アクションの例:
- name : コンテキストの品質を確認します
実行：contextops check context.json --min-score 75 --profile rag
contextops diff <file_a> <file_b>
2 つのコンテキスト スナップショットを比較して回帰を検出します。
contextops diff before.json after.json
スコア デルタ、どのペナルティが変更されたか、およびその変更が改善か後退かを表示します。 A/B テストの取得戦略やプロンプト リファクタリングに役立ちます。
決定論的安定性レポートを実行して、スコアリング エンジンが正しく動作していることを確認します。
コンテキストの安定性
コンテキストオプスタビ

lity context.json
コンテキストオプステレメトリ
コンテキスト品質の長期的な傾向を追跡するための、ローカルのみのオプトイン テレメトリ。
contextops telemetry status # テレメトリがアクティブかどうかを確認します
contextops telemetry log --limit 25 # 最近のイベントを表示
contextops テレメトリー傾向 --days 7 # 7 日間の品質傾向を表示します
contextops バッジ [--スコア N]
GitHub Shields.io マークダウン バッジを生成します。
contextops バッジ # 最後のテレメトリ スコアを使用します
contextops バッジ --score 87 # 特定のスコアを使用する
出力:
[![ ContextOps ] ( https://img.shields.io/badge/ContextOps-87-green )] ( https://github.com/Abhijeet777/contextops )
Python API
contextops から。 API 。検査インポートinspect_context
結果 = 検査コンテキスト (
payload 、 # dict 、リスト、またはプレーン文字列
model = "gpt-4o" , # トークンカウントのモデル
Archetype = "rag" , # アーキタイプ プロファイル (オプション)
)
print (結果.スコア) # int 0–100
print (結果 . token_breakdown . wasted_tokens )
記録の結果。推奨事項:
print ( f" -> { rec . fix } " )
完全な結果フィールド:
結果。スコア # int 0–100
結果。 Score_breakdown # 冗長性、密度、構造、集中ペナルティ
結果。 token_breakdown # total_tokens、by_type、wasted_tokens、estimated_cost_usd
結果。 redundancy_findings # リスト[冗長性の検出]
結果。 Structure_findings # リスト[StructureFinding]
結果。推奨事項 # リスト[推奨事項]
結果。 density_signal # 密度信号
結果。 Archetype_resolved # 例: 「ぼろ布」
結果。ロースト # str |なし (roaste_enabled の場合)
結果。メタデータ # item_count、モデル、バージョン
diff_contexts()
contextops から。 API 。 diffインポート diff_contexts
結果 = diff_contexts (ペイロード_a、ペイロード_b)
ContextOpsConfig
contextops から。コア。構成インポート ContextOpsConfig
config = ContextOpsConfig 。デフォルト (プロファイル = "ラグ")
config = ContextOpsC

オンフィグ。 from_dict ({
"retrieval_max_ratio" : 0.90 、
"system_max_ratio" : 0.30 、
"roast_enabled" : True
})
ラングチェーンの統合
pip インストール contextops langchain-core
contextops から ContextOps をインポート
# スコアをログに記録します (デフォルト — 非ブロッキング)
チェーン = チェーン。 with_config ({
"callbacks" : [ ContextOps .自動()]
})
# コンテキストの品質が低すぎる場合は実行をブロックする
チェーン = チェーン。 with_config ({
"コールバック" : [
コンテキストオペレーション 。自動 (
モード = "ブロック" 、
min_score = 75 、
プロフィール = "ぼろ"
)
】
})
結果 = チェーン。 invoke ({ "質問" : "返金ポリシーは何ですか?" })
モード
行動
「ログ」
スコアレポートを標準出力に出力します (デフォルト)
「警告する」
スコア < min_score の場合に Python 警告を発します
「ブロック」
スコア < min_score の場合は ContextOpsScoreError を発生させます — LLM 呼び出しをブロックします
アーキタイプのプロファイル
アーキタイプは、特定のユースケースに合わせて構造のしきい値を調整します。グローバル 0 ～ 100 スコアは決して影響を受けず、どの警告が発せられるかのみに影響します。
解決順序 (優先順位が最も高いものが優先):
Archetype= Python API 引数
JSON ペイロード内の「archetype」キー
ContextOps は 3 つの入力形式を受け入れます。
構造化された辞書 (推奨 — 最も豊富な分析が得られます):
{
"システム" : " ... " ,
"メッセージ" : [ ... ],
"チャンク" : [ ... ],
「記憶」: [ ... ],
「ツール」: [ ... ]
}
OpenAI メッセージ リスト (既存のペイロードを直接貼り付けます):
[
{ "role" : " system " , "content" : " あなたは役に立つボットです。 " },
{ "役割" : " ユーザー " 、 "コンテンツ" : " こんにちは " }
】
プレーン文字列 (システム プロンプトとして扱われます)。
スコア = max(0, min(100,round(100 − total_penalty)))
ここで、 total_penalty = 冗長性 + 密度 + 構造 + 集中。
語彙のみ。 ContextOps は、セマンティックな類似性ではなく、N グラムの重複、Jaccard の類似性、および正確な文字列一致を使用します。これは意図的なものです。これは構造アナライザーであり、意味論的なアナライザーではありません。
ContextBench — Co の LeetCode

ntextエンジニアリング
LeetCode がアルゴリズムの DSA である場合、ContextBench はコンテキストの DSA です。
競技プログラミングが、答えだけではなくアルゴリズムが重要であることを教えるのと同じように、ContextBench は、LLM が最終的に正しく実行されるかどうかだけではなく、コンテキスト アーキテクチャが重要であることを教えてくれます。
正しい答えを生成するブルートフォース O(N²) ソリューションでも、時間の計算量によってペナルティが課せられます。肥大化し、冗長なコンテキストが良好な LLM 応答を生成する場合でも、構造的無駄としてペナルティが課せられます。リーダーボードは両方に注目します。
LeetCode / DSA
ContextBench リーダーボード
問題セット
5 つの障害カテゴリにわたる 1,500 の事前構築済みコンテキスト ウィンドウ
裁判官
ContextOps エンジン — 決定論的スコアラー、毎回同じ出力
あなたの解決策
optimize_context(ctx) — 壊れたコンテキストを取得し、より適切なコンテキストを返します
スコア指標
品質 (50%) + 圧縮 (35%) + 遅延 (15%)
リーダーボード
すべてのベンチマーク サンプルにわたって、final_score によってランク付けされます
敵対的なトラック
ContextSecBench — 9,500 の敵対的攻撃ペイロード
問題セット
ContextBench には 5 つのカテゴリにわたる 1,500 のサンプルが含まれています。これらを難易度の段階として考えてください。
LeetCode の問題が期待される出力と同じように、すべてのサンプルにはグラウンド トゥルースがあります。
{
"ground_truth" : {
"failure_modes" : [ " システムプロ

[切り捨てられた]

## Original Extract

Contribute to Abhijeet777ui/contextops development by creating an account on GitHub.

GitHub - Abhijeet777ui/contextops · GitHub
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
Abhijeet777ui
/
contextops
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
23 Commits 23 Commits .github/ workflows .github/ workflows ContextBench ContextBench ContextSecBench ContextSecBench benchmarks benchmarks contextops-0.2.2/ tests contextops-0.2.2/ tests contextops contextops examples examples experiments experiments leaderboard leaderboard scenarios scenarios scripts scripts tests tests .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md HIGH_LEVEL_DOC.md HIGH_LEVEL_DOC.md HOW_TO_GET_JSON.md HOW_TO_GET_JSON.md LICENSE LICENSE README.md README.md STABILITY.md STABILITY.md USER_GUIDE.md USER_GUIDE.md output1.txt output1.txt output2.txt output2.txt output3.txt output3.txt pyproject.toml pyproject.toml test_duplicate.json test_duplicate.json View all files Repository files navigation
Static analysis for LLM context.
ContextOps is a deterministic, model-independent context linter for LLM applications.
It analyzes the context sent to an LLM before inference and detects structural problems such as redundancy, token waste, context imbalance, and source concentration. It produces a reproducible Context Health Score (CHS) together with actionable diagnostics — without embeddings, model calls, or external services.
Think of ContextOps as ESLint for LLM context .
Modern software engineering has deterministic quality gates.
Compilers catch syntax errors.
Formatters enforce consistency.
Static analyzers detect architectural problems.
LLM applications rarely have an equivalent layer for the context they send into models.
Instead, prompts quietly grow over time:
These issues increase latency and cost, make model behavior less predictable, and often go unnoticed until production.
ContextOps makes context quality observable, measurable, and testable before inference.
ContextOps evaluates four structural dimensions.
The result is a deterministic 0–100 Context Health Score together with detailed findings and suggested fixes.
ContextOps intentionally does not evaluate:
reasoning ability or hallucinations
It focuses exclusively on the structural quality of the context before inference.
contextops inspect context.json
Context Health Score: 81 / 100
✓ Low structural complexity
✓ Good source diversity
Warnings
• 214 duplicated tokens detected
• Retrieval occupies 78% of context
• Two retrieval chunks are near duplicates
Estimated token savings: 12%
Use --roast mode for more candid diagnostics:
contextops inspect context.json --roast
Installation
pip install contextops
Run the interactive demo:
contextops demo
Dependencies: Only tiktoken and click . No GPU, no network, no external API keys required.
Save your LLM payload before sending it to the model.
import json
messages = [
{ "role" : "system" , "content" : system_prompt },
{ "role" : "user" , "content" : user_query },
]
# Add two lines before your API call
with open ( "context.json" , "w" ) as f :
json . dump ( messages , f , indent = 2 )
response = openai . chat . completions . create ( model = "gpt-4o" , messages = messages )
Or use the structured dict format for richer analysis:
{
"system" : " You are a helpful customer support bot. " ,
"messages" : [
{ "role" : " user " , "content" : " How long will my refund take? " }
],
"chunks" : [
{ "content" : " Refunds take 3-5 business days. " , "source" : " docs/refunds.md " }
],
"memory" : [ " The user asked about a refund yesterday. " ],
"tools" : [{ "name" : " search_api " , "output" : " Tool response text here " }]
}
2. Inspect it
contextops inspect context.json
3. Enforce quality in CI
contextops check context.json --min-score 75
CLI Reference
Analyse a context file and display a rich report.
contextops inspect context.json
contextops inspect context.json --roast
contextops inspect context.json --profile rag
contextops inspect context.json --explain
contextops inspect context.json --json-output
Flag
Default
Description
--json-output
off
Output raw JSON instead of terminal format
--model <name>
gpt-4o
Model for token encoding (e.g. gpt-4 , claude-3 )
--profile <name>
general
Archetype: rag , agent , chatbot , toolchain
--config <path>
none
Path to a JSON config file with custom thresholds
--retrieval-max-ratio <f>
0.70
Override max allowed ratio for retrieval chunks
--system-max-ratio <f>
0.50
Override max allowed ratio for system prompt
--memory-max-ratio <f>
0.50
Override max allowed ratio for memory entries
--tool-max-ratio <f>
0.60
Override max allowed ratio for tool outputs
--explain
off
Show detailed "Top Score Drivers" for each penalty
--roast
off
Enable brutally honest score-band commentary
contextops check <file> --min-score N
CI gate. Exits with code 0 (pass) or 1 (fail).
contextops check context.json --min-score 75
Exit Code
Meaning
0
PASS — score meets or exceeds threshold
1
FAIL — score is below threshold or analysis error
2
ERROR — invalid JSON or unreadable input
GitHub Actions example:
- name : Check context quality
run : contextops check context.json --min-score 75 --profile rag
contextops diff <file_a> <file_b>
Compare two context snapshots to detect regressions.
contextops diff before.json after.json
Shows score delta, which penalties changed, and whether the change is an improvement or regression. Useful for A/B testing retrieval strategies or prompt refactors.
Run a deterministic stability report to verify the scoring engine is working correctly.
contextops stability
contextops stability context.json
contextops telemetry
Local-only, opt-in telemetry for tracking context quality trends over time.
contextops telemetry status # Check if telemetry is active
contextops telemetry log --limit 25 # Show recent events
contextops telemetry trends --days 7 # Show 7-day quality trends
contextops badge [--score N]
Generate a GitHub shields.io markdown badge.
contextops badge # Uses your last telemetry score
contextops badge --score 87 # Use a specific score
Output:
[ ![ ContextOps ] ( https://img.shields.io/badge/ContextOps-87-green )] ( https://github.com/Abhijeet777/contextops )
Python API
from contextops . api . inspect import inspect_context
result = inspect_context (
payload , # dict, list, or plain string
model = "gpt-4o" , # model for token counting
archetype = "rag" , # archetype profile (optional)
)
print ( result . score ) # int 0–100
print ( result . token_breakdown . wasted_tokens )
for rec in result . recommendations :
print ( f" -> { rec . fix } " )
Full result fields:
result . score # int 0–100
result . score_breakdown # redundancy, density, structure, concentration penalties
result . token_breakdown # total_tokens, by_type, wasted_tokens, estimated_cost_usd
result . redundancy_findings # list[RedundancyFinding]
result . structure_findings # list[StructureFinding]
result . recommendations # list[Recommendation]
result . density_signal # DensitySignal
result . archetype_resolved # e.g. "rag"
result . roast # str | None (if roast_enabled)
result . metadata # item_count, model, version
diff_contexts()
from contextops . api . diff import diff_contexts
result = diff_contexts ( payload_a , payload_b )
ContextOpsConfig
from contextops . core . config import ContextOpsConfig
config = ContextOpsConfig . default ( profile = "rag" )
config = ContextOpsConfig . from_dict ({
"retrieval_max_ratio" : 0.90 ,
"system_max_ratio" : 0.30 ,
"roast_enabled" : True
})
LangChain Integration
pip install contextops langchain-core
from contextops import ContextOps
# Log the score (default — non-blocking)
chain = chain . with_config ({
"callbacks" : [ ContextOps . auto ()]
})
# Block execution if context quality is too low
chain = chain . with_config ({
"callbacks" : [
ContextOps . auto (
mode = "block" ,
min_score = 75 ,
profile = "rag"
)
]
})
result = chain . invoke ({ "question" : "What is the refund policy?" })
Mode
Behaviour
"log"
Prints score report to stdout (default)
"warn"
Emits Python warning if score < min_score
"block"
Raises ContextOpsScoreError if score < min_score — blocks LLM call
Archetype Profiles
Archetypes adjust structural thresholds for your specific use case. The global 0–100 score is never affected — only which warnings fire.
Resolution order (highest priority wins):
archetype= Python API argument
"archetype" key inside the JSON payload
ContextOps accepts three input formats.
Structured dict (recommended — gives the richest analysis):
{
"system" : " ... " ,
"messages" : [ ... ],
"chunks" : [ ... ],
"memory" : [ ... ],
"tools" : [ ... ]
}
OpenAI message list (paste your existing payload directly):
[
{ "role" : " system " , "content" : " You are a helpful bot. " },
{ "role" : " user " , "content" : " Hello " }
]
Plain string (treated as a system prompt).
Score = max(0, min(100, round(100 − total_penalty)))
Where total_penalty = redundancy + density + structure + concentration .
Lexical only. ContextOps uses N-gram overlap, Jaccard similarity, and exact string matching — not semantic similarity. This is intentional: it is a structural analyser, not a semantic one.
ContextBench — The LeetCode for Context Engineering
If LeetCode is DSA for algorithms , ContextBench is DSA for context .
Just as competitive programming teaches you that the algorithm matters — not just the answer — ContextBench teaches you that context architecture matters , not just whether the LLM eventually gets it right.
A brute-force O(N²) solution that produces the correct answer still gets penalized for time complexity. A bloated, redundant context that still produces a good LLM response still gets penalized for structural waste. The leaderboard notices both.
LeetCode / DSA
ContextBench Leaderboard
Problem set
1,500 pre-built context windows across 5 failure categories
Judge
ContextOps engine — deterministic scorer, same output every time
Your solution
optimize_context(ctx) — takes a broken context, returns a better one
Score metric
Quality (50%) + Compression (35%) + Latency (15%)
Leaderboard
Ranked by final_score across all benchmark samples
Adversarial track
ContextSecBench — 9,500 adversarial attack payloads
The Problem Set
ContextBench contains 1,500 samples across 5 categories — think of these as your difficulty tiers:
Every sample has a ground truth — just like a LeetCode problem has expected output:
{
"ground_truth" : {
"failure_modes" : [ " system_pro

[truncated]
