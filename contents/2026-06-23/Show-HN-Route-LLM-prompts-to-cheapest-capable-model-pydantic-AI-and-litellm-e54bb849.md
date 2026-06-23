---
source: "https://github.com/Reactance0083/pydantic-ai-multi-llm-cost-optimizer"
hn_url: "https://news.ycombinator.com/item?id=48644688"
title: "Show HN: Route LLM prompts to cheapest capable model – pydantic-AI and litellm"
article_title: "GitHub - Reactance0083/pydantic-ai-multi-llm-cost-optimizer: Route prompts to the cheapest model that handles them. Claude + GPT-4o + Groq. Live cost tracking. Built with pydantic-ai + litellm. · GitHub"
author: "reactance0083"
captured_at: "2026-06-23T13:50:56Z"
capture_tool: "hn-digest"
hn_id: 48644688
score: 1
comments: 0
posted_at: "2026-06-23T13:30:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Route LLM prompts to cheapest capable model – pydantic-AI and litellm

- HN: [48644688](https://news.ycombinator.com/item?id=48644688)
- Source: [github.com](https://github.com/Reactance0083/pydantic-ai-multi-llm-cost-optimizer)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T13:30:26Z

## Translation

タイトル: HN を表示: LLM プロンプトを最も安価な対応モデルにルーティングする – pydantic-AI と litellm
記事のタイトル: GitHub - Reactance0083/pydantic-ai-multi-llm-cost-optimizer: プロンプトを、それらを処理する最も安価なモデルにルーティングします。クロード + GPT-4o + Groq。ライブコストの追跡。 pydantic-ai + litellm で構築されています。 · GitHub
説明: プロンプトを、それらを処理する最も安価なモデルにルーティングします。クロード + GPT-4o + Groq。ライブコストの追跡。 pydantic-ai + litellm で構築されています。 - Reactance0083/pydantic-ai-multi-llm-cost-optimizer

記事本文:
GitHub - Reactance0083/pydantic-ai-multi-llm-cost-optimizer: プロンプトを、それらを処理する最も安価なモデルにルーティングします。クロード + GPT-4o + Groq。ライブコストの追跡。 pydantic-ai + litellm で構築されています。 · GitHub
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
リアクタンス0083
/
pydantic-ai-multi-llm

-コストオプティマイザー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
Reactance0083/pydantic-ai-multi-llm-cost-optimizer
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット .env.example .env.example ライセンス ライセンス README.md README.md main.py main.pyrequirements.txt required.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
マルチ LLM コスト オプティマイザー (pydantic-ai + litellm + FastAPI)
すべてのプロンプトを適切に処理できる最も安価なモデルにルーティングします。ルーティングの決定には pydantic-ai を使用し、Claude、GPT-4o、Groq にわたる統合実行には litellm を使用します。ライブ /stats エンドポイントを使用してモデルごとのコストを追跡します。
品質レベル (高速 / 標準 / 品質 / 最高) のプロンプトを受け取ります
claude-haiku-4-5 をルーターとして使用して、最も安価な適切なモデルにルーティングします。
litellm 経由で実行 (認証とすべてのプロバイダーの API の違いを処理)
コストの内訳とレイテンシを含む応答を返します
pip install -r 要件.txt
cp .env.example .env
# 少なくとも ANTHROPIC_API_KEY を入力します。 OPENAI と GROQ はオプションです。
uvicorn main:app --reload --port 8002
APIの使用法
curl -X POST http://localhost:8002/complete \
-H " Content-Type: application/json " \
-d ' {
"prompt": "REST と GraphQL の主な違いを要約します",
"品質": "標準"、
"タスクタイプ": "一般"
} '
応答:
{
"テキスト" : " ... " 、
"model_used" : " anthropic/claude-haiku-4-5 " ,
"input_tokens" : 42 、
"output_tokens" : 218 、
"コスト_米ドル" : 0.000283 、
"レイテンシ_ミリ秒" : 847
}
GET /統計
{
「モデル」: {
"anthropic/claude-haiku-4-5" : { "呼び出し数" : 47 、 "total_cost" : 0.0134 、 "total_tokens" : 52400 }
}、
"合計コスト_米ドル" : 0.0134 ,
「合計コール数」: 47
}
料金表（2026年5月）
モデル
入力/1k
出力/1k
最適な用途
groq/llama-3.1-8b-インスタン

t
$0.00005
$0.00008
素早く簡単なタスク
人間論/クロード俳句-4-5
$0.00025
$0.00125
構造化された出力、分類
openai/gpt-4.1-mini
$0.0004
$0.0016
一般的なタスク、良い価値
人間性/クロード・ソネット-4-6
$0.003
$0.015
コード、複雑な推論
openai/gpt-4.1
$0.002
$0.008
複雑なタスク
人間的/クロード作品-4-7
$0.015
$0.075
最も難しいタスクのみ
openai/gpt-5.5
$0.005
$0.015
主要な推論、最も困難なタスク
品質レベル
階層
検討したモデル
いつ使用するか
速い
Groq llama-8b、クロードの俳句
ローステークス、ハイボリューム、シンプルな分類
標準
Groq llama-70b、GPT-4o-mini、クロード俳句
ほとんどの制作タスク
品質
クロード・ソネット、GPT-4o
コード生成、複雑な分析
最大
クロード・オプス
最も困難な問題、最も大きなリスク
構造化ルーティング (pydantic-ai)
クラス RoutingDecision (BaseModel):
model : str # 正確な litellm モデル文字列
reason : str # 1文の両端揃え
Expected_tokens : int # おおよその出力推定値
建築
POST /完了
→ ルーティングエージェント (claude-haiku-4-5) → RoutingDecision
→ litellm.completion(model=決定.モデル, ...)
→コスト計算→レスポンス+/統計更新
要件
OpenAI API キー (オプション、GPT ルーティングを有効にする)
Groq API キー (オプション、最も安価な層を有効にする)
5 つのテンプレートはすべて個別に、または 39 ドルのバンドルとして利用できます (個別に比べて 15 ドル節約)。
購入には、すべてのソース ファイル、README、requirements.txt、.env.example、およびライフタイム アップデートが含まれます。
無料で使用できます — ソースは GitHub にあります。購入すると開発の継続がサポートされ、すべてがパッケージ化されたクリーンなダウンロードが得られます。
Wade Allen によって構築 — AI ワークフロー アーキテクト
プロンプトを処理できる最も安価なモデルにプロンプトをルーティングします。クロード + GPT-4o + Groq。ライブコストの追跡。 pydantic-ai + litellm で構築されています。
リアクタンス0083.gumroad.com/l/ztmlv
トピックス
読み込み中にエラーが発生しました。お願いします

このページをリロードします。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Route prompts to the cheapest model that handles them. Claude + GPT-4o + Groq. Live cost tracking. Built with pydantic-ai + litellm. - Reactance0083/pydantic-ai-multi-llm-cost-optimizer

GitHub - Reactance0083/pydantic-ai-multi-llm-cost-optimizer: Route prompts to the cheapest model that handles them. Claude + GPT-4o + Groq. Live cost tracking. Built with pydantic-ai + litellm. · GitHub
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
Reactance0083
/
pydantic-ai-multi-llm-cost-optimizer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Reactance0083/pydantic-ai-multi-llm-cost-optimizer
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits .env.example .env.example LICENSE LICENSE README.md README.md main.py main.py requirements.txt requirements.txt View all files Repository files navigation
Multi-LLM Cost Optimizer (pydantic-ai + litellm + FastAPI)
Routes every prompt to the cheapest model that can handle it well. Uses pydantic-ai for the routing decision and litellm for unified execution across Claude, GPT-4o, and Groq. Tracks cost per model with a live /stats endpoint.
Receives a prompt with a quality tier ( fast / standard / quality / max )
Routes to the cheapest appropriate model using claude-haiku-4-5 as the router
Executes via litellm (handles auth + API differences for all providers)
Returns the response with cost breakdown and latency
pip install -r requirements.txt
cp .env.example .env
# Fill in at minimum ANTHROPIC_API_KEY. OPENAI and GROQ are optional.
uvicorn main:app --reload --port 8002
API Usage
curl -X POST http://localhost:8002/complete \
-H " Content-Type: application/json " \
-d ' {
"prompt": "Summarize the key differences between REST and GraphQL",
"quality": "standard",
"task_type": "general"
} '
Response:
{
"text" : " ... " ,
"model_used" : " anthropic/claude-haiku-4-5 " ,
"input_tokens" : 42 ,
"output_tokens" : 218 ,
"cost_usd" : 0.000283 ,
"latency_ms" : 847
}
GET /stats
{
"models" : {
"anthropic/claude-haiku-4-5" : { "calls" : 47 , "total_cost" : 0.0134 , "total_tokens" : 52400 }
},
"total_cost_usd" : 0.0134 ,
"total_calls" : 47
}
Cost Table (May 2026)
Model
Input/1k
Output/1k
Best For
groq/llama-3.1-8b-instant
$0.00005
$0.00008
Fast, simple tasks
anthropic/claude-haiku-4-5
$0.00025
$0.00125
Structured outputs, classification
openai/gpt-4.1-mini
$0.0004
$0.0016
General tasks, good value
anthropic/claude-sonnet-4-6
$0.003
$0.015
Code, complex reasoning
openai/gpt-4.1
$0.002
$0.008
Complex tasks
anthropic/claude-opus-4-7
$0.015
$0.075
Hardest tasks only
openai/gpt-5.5
$0.005
$0.015
Flagship reasoning, hardest tasks
Quality Tiers
Tier
Models Considered
Use When
fast
Groq llama-8b, Claude haiku
Low-stakes, high-volume, simple classification
standard
Groq llama-70b, GPT-4o-mini, Claude haiku
Most production tasks
quality
Claude sonnet, GPT-4o
Code generation, complex analysis
max
Claude opus
Hardest problems, highest stakes
Structured Routing (pydantic-ai)
class RoutingDecision ( BaseModel ):
model : str # exact litellm model string
reason : str # 1-sentence justification
expected_tokens : int # rough output estimate
Architecture
POST /complete
→ routing agent (claude-haiku-4-5) → RoutingDecision
→ litellm.completion(model=decision.model, ...)
→ cost calculation → response + /stats update
Requirements
OpenAI API key (optional, enables GPT routing)
Groq API key (optional, enables cheapest tier)
All 5 templates are available individually or as a $39 bundle (saves $15 vs individual).
Buying includes: all source files, README, requirements.txt, .env.example, and lifetime updates.
Free to use — the source is here on GitHub. Buying supports continued development and gets you a clean download with everything packaged.
Built by Wade Allen — AI Workflow Architect
Route prompts to the cheapest model that handles them. Claude + GPT-4o + Groq. Live cost tracking. Built with pydantic-ai + litellm.
reactance0083.gumroad.com/l/ztmlv
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
