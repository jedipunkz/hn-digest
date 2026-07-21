---
source: "https://github.com/vijaym2k6/SteerPlane"
hn_url: "https://news.ycombinator.com/item?id=48988992"
title: "Show HN: SteerPlane – open-source runtime guardrails for AI agents"
article_title: "GitHub - vijaym2k6/SteerPlane: Runtime Control Plane for Autonomous AI Agents - cost limits, loop detection, and full observability with one decorator · GitHub"
author: "vijaym2k6"
captured_at: "2026-07-21T07:34:00Z"
capture_tool: "hn-digest"
hn_id: 48988992
score: 1
comments: 0
posted_at: "2026-07-21T07:00:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: SteerPlane – open-source runtime guardrails for AI agents

- HN: [48988992](https://news.ycombinator.com/item?id=48988992)
- Source: [github.com](https://github.com/vijaym2k6/SteerPlane)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T07:00:10Z

## Translation

タイトル: Show HN: SteerPlane – AI エージェント用のオープンソース ランタイム ガードレール
記事のタイトル: GitHub - vijaym2k6/SteerPlane: 自律 AI エージェント用のランタイム コントロール プレーン - コスト制限、ループ検出、および 1 つのデコレーターによる完全な可観測性 · GitHub
説明: 自律 AI エージェント用のランタイム コントロール プレーン - コスト制限、ループ検出、および 1 つのデコレータによる完全な可観測性 - vijaym2k6/SteerPlane

記事本文:
GitHub - vijaym2k6/SteerPlane: 自律 AI エージェント用のランタイム コントロール プレーン - コスト制限、ループ検出、および 1 つのデコレーターによる完全な可観測性 · GitHub
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
ヴィジェイム2k6
/
ステアプレーン
公共
お知らせ

s
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
71 コミット 71 コミット .github/ workflows .github/ workflows api api アセット アセット ダッシュボード ダッシュボード ドキュメント ドキュメント サンプル サンプル スクリプト スクリプト sdk-ts sdk-ts sdk sdk .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE PATENTS.md PATENTS.md README.md README.md docker-compose.yml docker-compose.yml ruff.toml ruff.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自律型 AI エージェントのランタイム ガードレール。
コスト制限 · ループ検出 · デュアル強制 (キル/アラート) · ストリーミング ゲートウェイ · ポリシー エンジン · 人間参加型 · CLI · Docker · 4 つのフレームワーク統合
pip インストール ステアプレーン · npm インストール ステアプレーン
🌐 ステアプレーン.com
AI エージェントは、API を呼び出し、コードを実行し、Web を閲覧し、現実世界での意思決定を行うことができます。ガードレールなし:
🔄 構成を誤った単一のエージェントが無限ループに陥る可能性がある
💸 暴走エージェントは一晩で 10,000 ドル以上の API クレジットを使い果たす可能性があります
💀 エージェントは可視性ゼロで破壊的な行動を取る可能性があります
SteerPlane は 1 行のコードでこれを修正します。
ステアプレーン輸入ガードより
@ガード(
エージェント名 = "サポートボット" ,
max_cost_usd = 10.00 、
max_steps = 50 、
拒否されたアクション = [ "delete_*" , "sudo_*" ],
強制 = "アラート" 、
アラートしきい値 = 0.8 、
アラートタイムアウト_秒 = 1800 、
)
def run_agent():
# エージェントは正常に実行されます。
# SteerPlane はすべてのステップを静かに監視します。
# 財務/実行時間の制限は、人間の承認のために一時停止される場合があります。
# ループとポリシー違反は引き続き即時に終了します。
エージェント 。実行（）
🚀 ステアプレーン |実行が開始されました
実行ID: a3f8d2b1-...
エージェント: support_bot
制限: コスト $10.00 / 50 ステップ
──────

───────────────
✅ ステップ 1: クエリデータベース | 380 トークン | $0.0020 | 45ミリ秒
✅ ステップ 2: call_llm_analyze | 1240トークン | $0.0080 | 320ミリ秒
✅ ステップ 3: 知識の検索 | 560トークン | $0.0030 | 89ミリ秒
✅ ステップ 4:generate_response | 1800トークン | $0.0120 | 450ミリ秒
✅ ステップ 5: send_notification | 120トークン | $0.0010 | 200ミリ秒
─────────────────
✅ ステアプレーン |実行が完了しました
ステップ: 5
費用: $0.0260
トークン: 4,100
持続時間: 1.1秒
特長
特徴
何をするのか
🔄
ループ検出
O(W²) スライディング ウィンドウ アルゴリズムは、シングルアクション、交互、およびマルチステップの繰り返しパターンをミリ秒未満の時間で捕捉します。LLM 呼び出しは必要ありません。
💰
コスト上限
実行ごと (SDK) / セッションごと (ゲートウェイ) の USD 制限。各ステップの後にチェックされるため、オーバーシュートは 1 つのステップに制限されます。 OpenAI、Anthropic、Google、Meta、Mistral にわたる 25 以上のモデルに対する組み込みの価格設定
🌊
ストリーミングゲートウェイ
ストリーム中のコストキルによるリアルタイム SSE チャンク転送 — ストリーム中に予算を超過した場合、SteerPlane は終了イベントを挿入し、接続を切断します
🛡️
ポリシーエンジン
グロブ パターンとスライディング ウィンドウ レート制限を備えた許可/拒否リスト
🌐
ゲートウェイプロキシ
OpenAI 互換 API プロキシ — ゼロコード強制の場合は、base_url のみを変更します。エージェントはプロバイダー キーをゲートウェイ経由で渡します。これにより、すべてのトラフィックが上流に転送される前に強制的に適用されます。
🖥️
リアルタイムダッシュボード
自動更新、アニメーション化されたタイムライン、コストの内訳、ポリシー管理を備えた Next.js ダッシュボード
🔧
CLIツール
ステアプレーン実行リスト、ステアプレーンステータス、ステアプレーンキー作成 - 端末からすべてを管理
📄

設定ファイル
.steerplane.yml 自動検出 — ソース コードでハードコーディング制限なしでデフォルトを設定します
🔗
4 フレームワークの統合
LangChain、OpenAI Agents SDK、CrewAI、AutoGen — 設定不要のドロップイン ハンドラー
🐳
Docker Compose
1 つのコマンドで API + ダッシュボード + PostgreSQL を起動
🔌
グレースフル デグラデーション
API がダウンすると、SDK はすべての制限をローカルに適用します。エージェントが無防備になることはありません
🧪
CI/CD
GitHub Actions パイプライン — プッシュごとに lint、テスト、Docker を構築
ホスト/エンタープライズ層: アラート モードの人による承認ワークフロー (電子メール/Webhook 通知付き)、サーバー側プロバイダー キー ボールト、Redis 支援のマルチワーカー ゲートウェイ状態は、SteerPlane のホスト型プランで利用できます。オープンソース SDK は、enforcement="alert" クライアント オプションを引き続き公開しているため、ホスト型またはエンタープライズ展開に対してすぐに機能します。セルフホスティングでは、無料枠のみがキルモード強制を実行します。
[切り捨てられた]
オプション A: Docker (推奨)
git クローン https://github.com/vijaym2k6/SteerPlane.git
cd ステアプレーン
cp .env.example .env
ドッカー構成 -d
ローカルホストの API:8000 · ローカルホストのダッシュボード:3000 · PostgreSQL は自動構成されます。
# SDKをインストールする
pip インストール ステアプレーン
# APIを起動する
cd api && pip install -rrequirements.txt
uvicorn app.main:app --reload --port 8000
# ダッシュボードを開始する
cd ダッシュボード && npm install && npm run dev
オプション C: CLI
pip インストール ステアプレーン[cli]
ステアプレーンのステータス # API の健全性を確認する
steerplane run list # 最近の実行をリストする
steerplane key create -n prod # API キーを生成する
デモエージェントを実行する
Python の例/simple_agent/agent_example.py
localhost:3000 を開く → エージェントがリアルタイムで実行されているのを確認します。
ステアプレーン輸入ガードより
@ガード(
エージェント名 = "my_bot" ,
max_cost_usd = 10.00 、
max_steps = 50 、
max_runtime_sec = 300 、
強制 = "アラート" 、
アラートしきい値

d = 0.8、
拒否されたアクション = [ "delete_*" , "sudo_*" ],
allowed_actions = [ "search_*" , "read_*" , "generate_*" ],
rate_limits = [{ "パターン" : "call_llm*" , "max_count" : 20 , "ウィンドウ秒" : 60 }],
)
def run_my_agent():
代理店。実行（）
Python — コンテキスト マネージャー API
ステアプレーンからインポート SteerPlane
sp = SteerPlane (agent_id = "my_bot")
SP付き。 run ( max_cost_usd = 10.0 、 max_steps = 50 ) を実行します:
実行します。 log_step ( "query_db" 、トークン = 380 、コスト = 0.002 、latency_ms = 45 )
実行します。 log_step ( "生成" 、トークン = 1240 、コスト = 0.008 、latency_ms = 320 )
TypeScript
import { ガード , GuardOptions } から 'ステアプレーン' ;
const protectedAgent = Guard ( async ( run ) => {
実行を待ちます。 logStep({ アクション : 'query_db' 、トークン : 380 、コスト : 0.002 } ) ;
実行を待ちます。 logStep ( { アクション : '生成' 、トークン : 1240 、コスト : 0.008 } ) ;
「完了」を返します。
}、{
エージェント名 : 'support_bot' 、
maxCostUsd : 10.0 、
最大ステップ数 : 50 、
ポリシー: {
拒否されたアクション: [ 'delete_*' 、 'sudo_*' ] 、
} 、
} ) ;
const result = await protectedAgent() ;
例外処理
ステアプレーンから。例外インポート (
コスト制限を超えました 、
ループ検出エラー 、
ステップ制限を超えました 、
ポリシー違反エラー 、
)
@ ガード ( max_cost_usd = 5 、拒否されたアクション = [ "delete_*" ])
def run_agent():
試してみてください:
代理店。実行（）
e としての CostLimitExceeded を除く:
print ( f"予算を超えました: { e } " )
e としての LoopDetectedError を除く:
print ( f"ループが検出されました: { e } " )
e として StepLimitExceeded を除く:
print ( f"ステップ制限ヒット: { e } " )
e としての PolicyViolationError を除く:
print ( f"ポリシー違反: { e . アクション } は { e . ルール } によってブロックされました " )
設定ファイル (.steerplane.yml)
ハードコーディングの制限ではなく、プロジェクト ルートの .steerplane.yml にデフォルトを設定します。
api_url : http://localhost:8000
エージェント名 : my_bot
デフォルト:
最大コスト米ドル : 25.0
最大ステップ数 : 100

最大実行時間秒 : 1800
施行：警告
ループウィンドウサイズ : 10
ポリシー:
拒否されたアクション:
- 「削除_*」
- 「ドロップ_*」
rate_limits :
- パターン: " search_* "
最大数 : 10
ウィンドウ秒 : 60
アラート :
電子メール : ops@company.com
webhook_url : https://hooks.slack.com/...
しきい値 : 0.8
マージ順序: 明示的なデコレータ パラメータ → .steerplane.yml → ハードコードされたデフォルト。構成ファイルは、現在のディレクトリから上に移動すると自動的に検出されます。
ステアプレーンから。統合。 langchain import SteerPlaneCallbackHandler
ハンドラー = SteerPlaneCallbackHandler (
エージェント名 = "リサーチボット" ,
max_cost_usd = 5.0 、
max_steps = 30 、
)
llm = ChatOpenAI (モデル = "gpt-4o" 、コールバック = [ハンドラー])
代理店。 run ( "このデータを分析する" 、コールバック = [ ハンドラー ])
ハンドラー。終了（）
OpenAI エージェント SDK
ステアプレーンから。統合。 openai_agents インポート SteerPlaneAgentHooks
フック = SteerPlaneAgentHooks (
エージェント名 = "my_openai_agent" ,
max_cost_usd = 10.0 、
max_steps = 100 、
)
# 便利なラッパー
result = フックを待ちます。 run (エージェント、「こんにちは!」)
# または手動のライフサイクル
フック。開始()
result = ランナーを待ちます。 run (エージェント、「こんにちは!」)
フック。終了（）
CrewAI
ステアプレーンから。統合。 crewai インポート SteerPlaneCrewMonitor
モニター = SteerPlaneCrewMonitor (
エージェント名 = "my_crew" ,
max_cost_usd = 25.0 、
max_steps = 200 、
)
乗組員 = 乗組員 (
エージェント = [ 研究者 、 ライター ]、
タスク = [ リサーチタスク , ライトタスク ],
step_callback = モニター。 step_callback 、
)
結果＝モニター。キックオフ (クルー)
自動生成
ステアプレーンから。統合。 autogen インポート SteerPlaneAutoGenMonitor
モニター = SteerPlaneAutoGenMonitor (
エージェント名 = "my_autogen_group" ,
max_cost_usd = 15.0 、
max_steps = 150 、
)
結果＝モニター。 Initial_chat ( user_proxy 、 Assistant 、 message = "こんにちは!" )
必要な統合のみをインストールします

必要なもの:
pip install steerplane[langchain] # LangChain
pip install steerplane[cli] # CLI ツール
pip install steerplane[yaml] # 設定ファイルのサポート
pip install steerplane[all] # すべて
ゲートウェイ プロキシ (ゼロコード モード)
変更できないエージェントのために、SteerPlane はリアルタイム ストリーミングとミッドストリーム コスト強制を備えた OpenAI 互換ゲートウェイ プロキシを提供します。
openaiインポートからOpenAI
クライアント = OpenAI (
Base_url = "http://localhost:8000/gateway/v1" ,
api_key = "sk_sp_your_steerplane_key" ,
# 実際のプロバイダーキーはゲートウェイに送信され、ゲートウェイはそれを上流に転送します。
default_headers = { "X-LLM-API-Key" : "sk-your-real-provider-key" },
)
# ストリーミング作品 - リアルタイムで転送されるチャンク
クライアントのチャンク用。チャット 。完成品。作成(
モデル = "gpt-4o" 、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : "Hello" }],
ストリーム = True 、
):
print (チャンク . 選択肢 [ 0 ]. デルタ . コンテンツ , end = "" )
ゲートウェイがリクエストごとに強制する内容:
ポリシー ルール (拒否/許可/レート制限)
セッションコストと上限（ミッドストリームキルを含む）
SHA-256 プロンプトハッシュループ検出
Anthropic + OpenAI ストリーミングのサポート
セキュリティ モデル: エージェントは、ゲートウェイで OpenAI クライアントをポイントし、X-LLM-API-Key ヘッダーで実際のプロバイダー キーを渡します。ゲートウェイは SteerPlane キーを認証し、強制 (ポリシー → コスト → ループ) を通じてすべてのリクエストを実行し、その後初めてそのプロバイダー キーを使用して上流に転送します。そのため、エージェントはプロバイダーに直接到達したり、ガードレールをバイパスしたりすることはできません。
サーバー側のプロバイダー キーの保管 (

[切り捨てられた]

## Original Extract

Runtime Control Plane for Autonomous AI Agents - cost limits, loop detection, and full observability with one decorator - vijaym2k6/SteerPlane

GitHub - vijaym2k6/SteerPlane: Runtime Control Plane for Autonomous AI Agents - cost limits, loop detection, and full observability with one decorator · GitHub
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
vijaym2k6
/
SteerPlane
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
71 Commits 71 Commits .github/ workflows .github/ workflows api api assets assets dashboard dashboard docs docs examples examples scripts scripts sdk-ts sdk-ts sdk sdk .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PATENTS.md PATENTS.md README.md README.md docker-compose.yml docker-compose.yml ruff.toml ruff.toml View all files Repository files navigation
Runtime guardrails for autonomous AI agents.
Cost limits · Loop detection · Dual enforcement (Kill/Alert) · Streaming gateway · Policy engine · Human-in-the-loop · CLI · Docker · 4 framework integrations
pip install steerplane · npm install steerplane
🌐 steerplane.com
AI agents can call APIs, execute code, browse the web, and make real-world decisions. Without guardrails:
🔄 A single misconfigured agent can enter an infinite loop
💸 A runaway agent can burn through $10,000+ in API credits overnight
💀 Agents can take destructive actions with zero visibility
SteerPlane fixes this with one line of code.
from steerplane import guard
@ guard (
agent_name = "support_bot" ,
max_cost_usd = 10.00 ,
max_steps = 50 ,
denied_actions = [ "delete_*" , "sudo_*" ],
enforcement = "alert" ,
alert_threshold = 0.8 ,
alert_timeout_sec = 1800 ,
)
def run_agent ():
# Your agent runs normally.
# SteerPlane silently monitors every step.
# Financial/runtime limits can pause for human approval.
# Loops and policy violations still terminate immediately.
agent . run ()
🚀 SteerPlane | Run Started
Run ID: a3f8d2b1-...
Agent: support_bot
Limits: $10.00 cost / 50 steps
─────────────────────────────────────────────
✅ Step 1: query_database | 380 tokens | $0.0020 | 45ms
✅ Step 2: call_llm_analyze | 1240 tokens | $0.0080 | 320ms
✅ Step 3: search_knowledge | 560 tokens | $0.0030 | 89ms
✅ Step 4: generate_response | 1800 tokens | $0.0120 | 450ms
✅ Step 5: send_notification | 120 tokens | $0.0010 | 200ms
─────────────────────────────────────────────
✅ SteerPlane | Run COMPLETED
Steps: 5
Cost: $0.0260
Tokens: 4,100
Duration: 1.1s
Features
Feature
What It Does
🔄
Loop Detection
O(W²) sliding-window algorithm catches single-action, alternating, and multi-step repeating patterns in sub-millisecond time — no LLM calls
💰
Cost Ceiling
Per-run (SDK) / per-session (gateway) USD limits, checked after each step so overshoot is bounded to a single step. Built-in pricing for 25+ models across OpenAI, Anthropic, Google, Meta, and Mistral
🌊
Streaming Gateway
Real-time SSE chunk forwarding with mid-stream cost kill — if the budget is exceeded during a stream, SteerPlane injects a termination event and cuts the connection
🛡️
Policy Engine
Allow/deny lists with glob patterns and sliding-window rate limits
🌐
Gateway Proxy
OpenAI-compatible API proxy — change only base_url for zero-code enforcement. The agent passes its provider key through the gateway, which forces all traffic through enforcement before forwarding upstream
🖥️
Real-Time Dashboard
Next.js dashboard with auto-refresh, animated timelines, cost breakdowns, and policy management
🔧
CLI Tool
steerplane runs list , steerplane status , steerplane keys create — manage everything from your terminal
📄
Config File
.steerplane.yml auto-discovery — set defaults without hardcoding limits in source code
🔗
4 Framework Integrations
LangChain, OpenAI Agents SDK, CrewAI, AutoGen — zero-config drop-in handlers
🐳
Docker Compose
One command brings up API + Dashboard + PostgreSQL
🔌
Graceful Degradation
If the API goes down, the SDK enforces all limits locally. Agents are never unprotected
🧪
CI/CD
GitHub Actions pipeline — lint, test, build Docker on every push
Hosted/Enterprise tier: alert-mode human-approval workflows (with email/webhook notifications), server-side provider-key vaulting, and Redis-backed multi-worker gateway state are available on SteerPlane's hosted plan. The open-source SDK still exposes the enforcement="alert" client options so it works out of the box against a hosted or enterprise deployment — self-hosting only the free tier here runs kill-mode enforcement
[truncated]
Option A: Docker (Recommended)
git clone https://github.com/vijaym2k6/SteerPlane.git
cd SteerPlane
cp .env.example .env
docker compose up -d
API at localhost:8000 · Dashboard at localhost:3000 · PostgreSQL auto-configured.
# Install the SDK
pip install steerplane
# Start the API
cd api && pip install -r requirements.txt
uvicorn app.main:app --reload --port 8000
# Start the Dashboard
cd dashboard && npm install && npm run dev
Option C: CLI
pip install steerplane[cli]
steerplane status # Check API health
steerplane runs list # List recent runs
steerplane keys create -n prod # Generate API key
Run the Demo Agent
python examples/simple_agent/agent_example.py
Open localhost:3000 → See your agent run in real time.
from steerplane import guard
@ guard (
agent_name = "my_bot" ,
max_cost_usd = 10.00 ,
max_steps = 50 ,
max_runtime_sec = 300 ,
enforcement = "alert" ,
alert_threshold = 0.8 ,
denied_actions = [ "delete_*" , "sudo_*" ],
allowed_actions = [ "search_*" , "read_*" , "generate_*" ],
rate_limits = [{ "pattern" : "call_llm*" , "max_count" : 20 , "window_seconds" : 60 }],
)
def run_my_agent ():
agent . run ()
Python — Context Manager API
from steerplane import SteerPlane
sp = SteerPlane ( agent_id = "my_bot" )
with sp . run ( max_cost_usd = 10.0 , max_steps = 50 ) as run :
run . log_step ( "query_db" , tokens = 380 , cost = 0.002 , latency_ms = 45 )
run . log_step ( "generate" , tokens = 1240 , cost = 0.008 , latency_ms = 320 )
TypeScript
import { guard , GuardOptions } from 'steerplane' ;
const protectedAgent = guard ( async ( run ) => {
await run . logStep ( { action : 'query_db' , tokens : 380 , cost : 0.002 } ) ;
await run . logStep ( { action : 'generate' , tokens : 1240 , cost : 0.008 } ) ;
return 'done' ;
} , {
agentName : 'support_bot' ,
maxCostUsd : 10.0 ,
maxSteps : 50 ,
policy : {
deniedActions : [ 'delete_*' , 'sudo_*' ] ,
} ,
} ) ;
const result = await protectedAgent ( ) ;
Exception Handling
from steerplane . exceptions import (
CostLimitExceeded ,
LoopDetectedError ,
StepLimitExceeded ,
PolicyViolationError ,
)
@ guard ( max_cost_usd = 5 , denied_actions = [ "delete_*" ])
def run_agent ():
try :
agent . run ()
except CostLimitExceeded as e :
print ( f"Budget exceeded: { e } " )
except LoopDetectedError as e :
print ( f"Loop detected: { e } " )
except StepLimitExceeded as e :
print ( f"Step limit hit: { e } " )
except PolicyViolationError as e :
print ( f"Policy violation: { e . action } blocked by { e . rule } " )
Config File (.steerplane.yml)
Set defaults in a .steerplane.yml at your project root instead of hardcoding limits:
api_url : http://localhost:8000
agent_name : my_bot
defaults :
max_cost_usd : 25.0
max_steps : 100
max_runtime_sec : 1800
enforcement : alert
loop_window_size : 10
policy :
denied_actions :
- " delete_* "
- " drop_* "
rate_limits :
- pattern : " search_* "
max_count : 10
window_seconds : 60
alerts :
email : ops@company.com
webhook_url : https://hooks.slack.com/...
threshold : 0.8
Merge order: Explicit decorator params → .steerplane.yml → hardcoded defaults. The config file is auto-discovered by walking up from the current directory.
from steerplane . integrations . langchain import SteerPlaneCallbackHandler
handler = SteerPlaneCallbackHandler (
agent_name = "research_bot" ,
max_cost_usd = 5.0 ,
max_steps = 30 ,
)
llm = ChatOpenAI ( model = "gpt-4o" , callbacks = [ handler ])
agent . run ( "Analyze this data" , callbacks = [ handler ])
handler . finish ()
OpenAI Agents SDK
from steerplane . integrations . openai_agents import SteerPlaneAgentHooks
hooks = SteerPlaneAgentHooks (
agent_name = "my_openai_agent" ,
max_cost_usd = 10.0 ,
max_steps = 100 ,
)
# Convenience wrapper
result = await hooks . run ( agent , "Hello!" )
# Or manual lifecycle
hooks . start ()
result = await Runner . run ( agent , "Hello!" )
hooks . finish ()
CrewAI
from steerplane . integrations . crewai import SteerPlaneCrewMonitor
monitor = SteerPlaneCrewMonitor (
agent_name = "my_crew" ,
max_cost_usd = 25.0 ,
max_steps = 200 ,
)
crew = Crew (
agents = [ researcher , writer ],
tasks = [ research_task , write_task ],
step_callback = monitor . step_callback ,
)
result = monitor . kickoff ( crew )
AutoGen
from steerplane . integrations . autogen import SteerPlaneAutoGenMonitor
monitor = SteerPlaneAutoGenMonitor (
agent_name = "my_autogen_group" ,
max_cost_usd = 15.0 ,
max_steps = 150 ,
)
result = monitor . initiate_chat ( user_proxy , assistant , message = "Hello!" )
Install only the integration you need:
pip install steerplane[langchain] # LangChain
pip install steerplane[cli] # CLI tool
pip install steerplane[yaml] # Config file support
pip install steerplane[all] # Everything
Gateway Proxy (Zero-Code Mode)
For agents you can't modify, SteerPlane provides an OpenAI-compatible gateway proxy with real-time streaming and mid-stream cost enforcement :
from openai import OpenAI
client = OpenAI (
base_url = "http://localhost:8000/gateway/v1" ,
api_key = "sk_sp_your_steerplane_key" ,
# The real provider key is sent to the gateway, which forwards it upstream.
default_headers = { "X-LLM-API-Key" : "sk-your-real-provider-key" },
)
# Streaming works — chunks forwarded in real-time
for chunk in client . chat . completions . create (
model = "gpt-4o" ,
messages = [{ "role" : "user" , "content" : "Hello" }],
stream = True ,
):
print ( chunk . choices [ 0 ]. delta . content , end = "" )
What the gateway enforces per request:
Policy rules (deny/allow/rate limits)
Session cost vs. ceiling (including mid-stream kill)
SHA-256 prompt-hash loop detection
Anthropic + OpenAI streaming support
Security model: The agent points its OpenAI client at the gateway and passes the real provider key in the X-LLM-API-Key header. The gateway authenticates the SteerPlane key, runs every request through enforcement (policy → cost → loop), and only then forwards it upstream with that provider key — so the agent can't reach the provider directly or bypass the guardrails.
Server-side provider-key vaulting (

[truncated]
