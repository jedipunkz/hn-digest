---
source: "https://github.com/obro79/promptetheus"
hn_url: "https://news.ycombinator.com/item?id=48695176"
title: "Promptetheus – Trace, detect, and auto-repair AI agent failures"
article_title: "GitHub - obro79/promptetheus: Debugging infrastructure for AI agents · GitHub"
author: "tar-ive"
captured_at: "2026-06-27T05:16:33Z"
capture_tool: "hn-digest"
hn_id: 48695176
score: 1
comments: 0
posted_at: "2026-06-27T04:37:00Z"
tags:
  - hacker-news
  - translated
---

# Promptetheus – Trace, detect, and auto-repair AI agent failures

- HN: [48695176](https://news.ycombinator.com/item?id=48695176)
- Source: [github.com](https://github.com/obro79/promptetheus)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T04:37:00Z

## Translation

タイトル: Promptetheus – AI エージェントの障害を追跡、検出、自動修復する
記事のタイトル: GitHub - obro79/promptetheus: AI エージェントのデバッグ インフラストラクチャ · GitHub
説明: AI エージェントのデバッグ インフラストラクチャ。 GitHub でアカウントを作成して、obro79/promptetheus の開発に貢献してください。

記事本文:
GitHub - obro79/promptetheus: AI エージェントのデバッグ インフラストラクチャ · GitHub
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
オブロ79
/
プロンプテテウス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

コード 「その他のアクション」メニューを開く フォルダーとファイル
27 コミット 27 コミット .agent/ スキル/ プロンプトテウス .agent/ スキル/ プロンプトテウス .github/ ワークフロー .github/ ワークフロー ドキュメント ドキュメント サンプル サンプル パッケージ/ プロンプトテウス パッケージ/ プロンプトテウス テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md conftest.py conftest.py start.md start.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Promptetheus は AI エージェントのインフラストラクチャをデバッグしています: Python SDK、ローカル
リプレイツール、ホストされたトレース配信、コーディングのための MCP 証拠アクセス
失敗したエージェントの実行を修正する必要があるエージェント。
ユーザーに表示されるエージェント タスクごとに 1 つのトレース。
最上位のエージェントの実行、ツール呼び出し、およびネストされたスパンのデコレーター。
ユーザーメッセージ、エージェントメッセージ、ツール呼び出し、ブラウザアクションなどの型付きイベント
DOM スナップショット、スクリーンショット、LLM 呼び出し、取得、メトリクス、エラー、スコア、
そして最終目標のチェック。
ホストエージェントを決してクラッシュさせない耐久性のある配信。 HTTP配信ができない場合
設定または失敗した場合、イベントはローカルにスプールされ、後で再実行できます。
ドクターチェック、スプール検査、セッションリプレイ、差分分析用のローカル CLI ツール
そして失敗の指紋。
範囲を限定した読み取り専用インシデント証拠のホストされた MCP 構成スニペット
ワークスペースと Supabase プロジェクト。
通常のプロジェクトの場合は、PyPI からインストールします。
pip インストール プロンプトテウス
プロンプテテウスのバージョン
ホストされたプロジェクト キーを作成または構成します。
エクスポート PROMPTETHEUS_CONSOLE_TOKEN=...
プロンプテテウス初期化\
--ワークスペース名 " Acme " \
--プロジェクト名 " ブラウザ エージェント " \
--write-env .env
ソース.env
プロンプテテウス博士
ローカルの自己ホスト型開発の場合:
プロンプテテウス初期化\
--api-url http://127.0.0.1:4318 \
--コンソールトークン pt_console_token \
--write-env .env
ソース.env
このリポジトリからの寄稿者の作品については、次のとおりです。
pip install -e パッケージ/promptetheus
プロンプテテウスのバージョン
Transport="auto" の場合、

SDK は次の場合に設定された API に送信します。
PROMPTETHEUS_API_KEY が存在します。キーがなければ、ローカル スプールに書き込みます
そのため、インスツルメントされたエージェントは実行を続けます。
インストルメンテーションをエージェントやツールに直接配置したい場合は、デコレータを使用します。
機能:
プロンプテテウスをptとしてインポート
@pt.ツール
def search_calendar (日: str ) -> リスト [ str ]:
return [ "火曜日午後 2 時" 、"火曜日午後 3 時" ]
@pt.トレース済み ( "スロットの選択" )
def Choice_slot (スロット : リスト [ str ]) -> str :
「水曜日の午後2時」に戻ります
@pt.観察してください（
エージェント = "カレンダーエージェント" ,
user_goal = "火曜日の午後 2 時に予約" ,
Transport = "auto" , # これを試すときにローカル JSONL を強制するには "spool" を使用します
)
def run_agent (ゴール: str ) -> str :
点。現在 （）。 user_message (目標)
スロット = search_calendar ( "火曜日" )
selected =Choose_slot (スロット)
点。現在 （）。 Agent_message ( f"予約済み { 選択済み } " )
点。現在 （）。ゴールチェック (
偽、
不一致 = [ "火曜日ではなく水曜日が選択されました" ],
)
選択したものを返す
run_agent ( "火曜日の午後 2 時に予約" )
各デコレーターが行うこと:
@pt.observe(...) は、トップレベルの実行を中心に 1 つのトレース/セッションを開始します。
@pt.tool は、現在の内部にtool_callイベントとtool_resultイベントを記録します。
セッション。
@pt.traced("name") は、リプレイツリーを開始せずに、ネストされたスパンをリプレイツリーに追加します。
別セッション。
pt.current() はアクティブなセッションを返すので、エージェントはユーザーを記録できます。
メッセージ、エージェントメッセージ、目標チェック、エラー、メトリクス、その他のイベント。
goal_check(False) は、リプレイ、フィンガープリント、およびテール サンプリングで表示されます。もし
失敗した目標はプロセスも失敗させ、目標チェックを記録してから、
例外を発生させて、端末の session_end ステータスを failed にします。
選択されていない場合。 ( "火曜日" ) で始まります:
点。現在 （）。 goal_check ( False 、不一致 = [ "選択された水曜日" ])
raise RuntimeError (「エージェントが間違った日を選択しました」)
なんだよ

あなたは見ることができます
API キーが設定されていない場合、transport="auto" はローカル JSONL を書き込みます。その間
また、transport="spool" を渡してローカル出力を強制することもできます。後
ローカルまたはスプールされた実行、リストセッション:
プロンプテテウスセッション
出力例:
01KVMZ4T7V2SN61ZWG1XTDBK47: 11 イベント
タイムラインを再生します。
プロンプテテウス リプレイ 01KVMZ4T7V2SN61ZWG1XTDBK47
出力例:
[0] state_change name='session_started'
[1] ツールコール ツール名='実行エージェント'
[2] user_message content='火曜日の午後 2 時に予約'
[3] ツールコール ツール名='検索カレンダー'
[4] ツール結果呼び出し ID='190a6438979141f5ac11b2e1b2ee29a0'
[5] state_change name='span_start'
[6] state_change name='span_end'
[7] Agent_message content='水曜日の午後 2 時に予約済み'
[8] ゴールチェックに合格=False
[9] ツール結果呼び出し ID='a78566297e0a4a309d5ce44cefe0d836'
[10] session_end status='完了'
実行ツリーを再生します。
プロンプトテウス リプレイ 01KVMZ4T7V2SN61ZWG1XTDBK47 --tree
出力例:
[0] state_change name='session_started'
[1] ツールコール ツール名='実行エージェント'
[2] user_message content='火曜日の午後 2 時に予約'
[3] ツールコール ツール名='検索カレンダー'
[4] ツール結果呼び出し ID='190a6438979141f5ac11b2e1b2ee29a0'
[7] Agent_message content='水曜日の午後 2 時に予約済み'
[8] ゴールチェックに合格=False
[9] ツール結果呼び出し ID='a78566297e0a4a309d5ce44cefe0d836'
[10] session_end status='完了'
選択スロット スパン=span_163a8380174647e98bfe1f3fff9e15b9duration_ms=0.0
失敗のフィンガープリントを生成します。
プロンプテテウスの指紋 01KVMZ4T7V2SN61ZWG1XTDBK47
出力例:
8ae0f41220d0 目標の不一致: 火曜日ではなく水曜日が選択されました
- 目標: 火曜日ではなく水曜日を選択
ローカル配信スプールを検査します。
プロンプテテウスのスプールリスト
出力例:
スプール: .promptetheus/スプール
保留中: 1 つのセッション ファイルにわたる 11 のイベント、4082 バイト
デッド: 0 ファイル、0 バイトにわたる 0 イベント
01KVMZ4T7V2SN61

ZWG1XTDBK47: 11 保留中
生のスプールは JSONL です。各行はイベント エンベロープです。
{
"type" : "tool_call" ,
"session_id" : " 01KVMZ4T7V2SN61ZWG1XTDBK47 " ,
"シーケンス" : 1 、
"idempotency_key" : " 01KVMZ4T7V2SN61ZWG1XTDBK47:29c5eff0:1 " ,
「ペイロード」: {
"tool_name" : " run_agent " ,
"call_id" : " a78566297e0a4a309d5ce44cefe0d836 " ,
"引数" : {
"args" : " ('火曜日の午後 2 時に予約',) " ,
"クワーグ" : " {} "
}
}
}
手動トレース API
実行境界を制御し、明示的にしたい場合は、pt.trace.start(...) を使用します。
デコレータの代わりにイベントを呼び出します。
プロンプテテウスをptとしてインポート
pt付き。跡が残ります。開始（
エージェント = "デモエージェント" 、
user_goal = "火曜日の会議を予約" ,
トランスポート = "自動" 、
) セッションとして:
セッション。 user_message ( "火曜日の午後 2 時に小さな部屋を予約してください" )
セッション。 tool_call ( "calendar.search" , { "day" : "火曜日" }, call_id = "calendar-1" )
セッション。 tool_result ( "calendar-1" , result = { "利用可能" : [ "午後 2 時" , "午後 3 時" ]})
セッション。 Agent_message ( "水曜日の午後 2 時に予約が確認されました" )
セッション。 goal_check ( False 、不一致 = [ "火曜日ではなく水曜日に予約されました" ])
# session_end は自動的に発行されます。出口でトランスポートフラッシュが実行される
パブリックSDK API
パッケージは次の主要なエントリ ポイントを公開します。
プロンプテテウスをptとしてインポート
点。跡が残ります。開始 (...)
点。開始 (...)
点。観察してください(...)
点。ツール
点。追跡されました (...)
点。現在 ()
点。セッション
点。非同期セッション
点。エージェントランタイム
一般的なセッション ヘルパー:
セッション。 user_message ( "火曜日午後 2 時 (太平洋時間) を予約" )
セッション。 Agent_message ( "空き状況が見つかりました" )
セッション。 tool_call ( "browser.click" , { "selector" : "#checkout" }, call_id = "click-1" )
セッション。 tool_result ( "click-1" , result = { "ok" : True })
セッション。取得 ( "返金ポリシー" , 文書 = [{ "id" : "doc-1" , "スコア" : 0.91 }])
セッション。ブラウザアクション

( "クリック" 、 "#checkout" 、 url = ページ . url )
セッション。 dom_snapshot ( page . url ,visible_text , selected_values = { "日" : "火曜日" })
セッション。スクリーンショット ( ページ . スクリーンショット ())
セッション。 restart_artifact ( "trace.webm" 、 artifact_type = "screen_recording" 、event_time_map = {})
セッション。 llm_call ( "gpt-5" 、input_tokens = 100 、output_tokens = 40 、latency_ms = 900 )
セッション。スコア ( "goal_match" , 0.2 , comment = "間違った日を選択しました" )
セッション。メトリック ( "ステップ" , 12 , 単位 = "カウント" )
セッション。エラー ( RuntimeError ( "カレンダー API タイムアウト" )、処理済み = True )
セッション。 goal_check ( False 、不一致 = [ "選択された水曜日" ])
セッション。終了 (「失敗」)
セッション。フラッシュ (タイムアウト = 2)
すべてのヘルパーは、 type 、 session_id 、
timestamp 、 seq 、 idempotency_key 、および payload 。安全のためにメタデータを使用する
カーディナリティの低いコンテキスト。生のシークレット、Cookie、トークン、または
資格情報をイベント ペイロードに組み込みます。
最上位エージェントの実行が非同期である場合は、AsyncSession を使用します。
プロンプトテウスから AsyncSession をインポート
AsyncSession (agent = "voice-agent" 、 user_goal = "Summarize the call" ) をセッションとして使用する非同期:
セッション。 user_message ( "この通話を要約します" )
セッションと非同期。 aspan ( "転写" ):
セッション。メトリック ( "audio_秒" , 42 , 単位 = "秒" )
セッション。ゴールチェック ( True )
ブラウザエージェント
ブラウザエージェントは、ユーザーの目標、重要なブラウザアクション、最終的なアクションを記録する必要があります。
DOM 状態と明示的な目標チェック:
セッション。 browser_action ( "click" 、 "#confirm" 、 url = page . url )
セッション。 dom_スナップショット (
ページにURL 、
visible_text = ページを待ちます。ロケーター (「ボディ」)。内部テキスト ()、
selected_values = { "日" : "水曜日" , "時間" : "午後 2 時" },
warnings = [ "タイムゾーンが太平洋から東部に変更されました" ],
)
セッション。ゴールチェック (
偽、
不一致

hes = [ "水曜日に予約済み" , "タイムゾーンの警告が表示されます" ],
)
これは、Promptetheus が失敗をリプレイし、fix-agent を生成できるパスです。
一般的なログを保存するだけでなく、証拠を保存します。
アダプターはオプションであり、遅延インポートされます。必要な追加機能のみをインストールします。
pip install "プロンプトテテウス[openai]"
pip install "プロンプトテテウス[anthropic]"
pip install "プロンプトテウス[langchain]"
pip install "プロンプトテテウス[劇作家]"
利用可能なアダプターのエクスポート:
プロンプテテウスより。アダプターのインポート (
人間アダプター 、
AutoGenAdapter 、
CrewAIアダプター、
DSPyAdapter 、
ヘイスタックアダプター、
LangGraphAdapter 、
LiteLLMAアダプター 、
LlamaIndexAdapter 、
OpenAIAdapter 、
OpenTelemetryBridge 、
劇作家アダプター 、
PromptetheusCallbackHandler 、
PydanticAIアダプター 、
)
フレームワークがすでに構造化コールバックを発行している場合は、アダプターを使用します。カスタムを維持する
フレームワークがそうでない場合でも、実際の実行境界に近いインストルメンテーション。
AgentRuntime は、サービスを利用したライブ調整のためのベストエフォート型クライアントです。
これは永続的なトレース ストレージから分離されており、次の場合にホスト コードに呼び出されることはありません。
サービスが利用できません:
プロンプトテウスから AgentRuntime をインポート
ランタイム = AgentRuntime ( session . session_id )
ランタイム。覚えておいてください ( "仮説" , { "概要" : "認証

[切り捨てられた]

## Original Extract

Debugging infrastructure for AI agents. Contribute to obro79/promptetheus development by creating an account on GitHub.

GitHub - obro79/promptetheus: Debugging infrastructure for AI agents · GitHub
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
obro79
/
promptetheus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits .agent/ skills/ promptetheus .agent/ skills/ promptetheus .github/ workflows .github/ workflows docs docs examples examples packages/ promptetheus packages/ promptetheus tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md conftest.py conftest.py start.md start.md View all files Repository files navigation
Promptetheus is debugging infrastructure for AI agents: a Python SDK, local
replay tooling, hosted trace delivery, and MCP evidence access for coding
agents that need to fix failing agent runs.
One trace per user-visible agent task.
Decorators for top-level agent runs, tool calls, and nested spans.
Typed events for user messages, agent messages, tool calls, browser actions,
DOM snapshots, screenshots, LLM calls, retrieval, metrics, errors, scores,
and final goal checks.
Durable delivery that never crashes the host agent. If HTTP delivery is not
configured or fails, events spool locally and can be replayed later.
Local CLI tools for doctor checks, spool inspection, session replay, diffing,
and failure fingerprints.
Hosted MCP config snippets for read-only incident evidence scoped to a
workspace and Supabase project.
For a normal project, install from PyPI:
pip install promptetheus
promptetheus version
Create or configure a hosted project key:
export PROMPTETHEUS_CONSOLE_TOKEN=...
promptetheus init \
--workspace-name " Acme " \
--project-name " Browser Agent " \
--write-env .env
source .env
promptetheus doctor
For local self-hosted development:
promptetheus init \
--api-url http://127.0.0.1:4318 \
--console-token pt_console_token \
--write-env .env
source .env
For contributor work from this repository:
pip install -e packages/promptetheus
promptetheus version
With transport="auto" , the SDK sends to the configured API when
PROMPTETHEUS_API_KEY is present. Without a key, it writes to the local spool
so the instrumented agent keeps running.
Use decorators when you want instrumentation to sit directly on agent and tool
functions:
import promptetheus as pt
@ pt . tool
def search_calendar ( day : str ) -> list [ str ]:
return [ "Tuesday 2pm" , "Tuesday 3pm" ]
@ pt . traced ( "choose-slot" )
def choose_slot ( slots : list [ str ]) -> str :
return "Wednesday 2pm"
@ pt . observe (
agent = "calendar-agent" ,
user_goal = "Book Tuesday at 2pm" ,
transport = "auto" , # use "spool" to force local JSONL while trying this
)
def run_agent ( goal : str ) -> str :
pt . current (). user_message ( goal )
slots = search_calendar ( "Tuesday" )
selected = choose_slot ( slots )
pt . current (). agent_message ( f"Booked { selected } " )
pt . current (). goal_check (
False ,
mismatches = [ "selected Wednesday, not Tuesday" ],
)
return selected
run_agent ( "Book Tuesday at 2pm" )
What each decorator does:
@pt.observe(...) starts one trace/session around the top-level run.
@pt.tool records tool_call and tool_result events inside the current
session.
@pt.traced("name") adds a nested span to the replay tree without starting a
separate session.
pt.current() returns the active session so the agent can record user
messages, agent messages, goal checks, errors, metrics, and other events.
goal_check(False) is visible in replay, fingerprints, and tail sampling. If a
failed goal should also make the process fail, record the goal check and then
raise an exception so the terminal session_end status is failed :
if not selected . startswith ( "Tuesday" ):
pt . current (). goal_check ( False , mismatches = [ "selected Wednesday" ])
raise RuntimeError ( "agent selected the wrong day" )
What You Can See
When no API key is configured, transport="auto" writes local JSONL. While
learning, you can also pass transport="spool" to force local output. After a
local or spooled run, list sessions:
promptetheus sessions
Example output:
01KVMZ4T7V2SN61ZWG1XTDBK47: 11 event(s)
Replay the timeline:
promptetheus replay 01KVMZ4T7V2SN61ZWG1XTDBK47
Example output:
[0] state_change name='session_started'
[1] tool_call tool_name='run_agent'
[2] user_message content='Book Tuesday at 2pm'
[3] tool_call tool_name='search_calendar'
[4] tool_result call_id='190a6438979141f5ac11b2e1b2ee29a0'
[5] state_change name='span_start'
[6] state_change name='span_end'
[7] agent_message content='Booked Wednesday 2pm'
[8] goal_check passed=False
[9] tool_result call_id='a78566297e0a4a309d5ce44cefe0d836'
[10] session_end status='completed'
Replay the run tree:
promptetheus replay 01KVMZ4T7V2SN61ZWG1XTDBK47 --tree
Example output:
[0] state_change name='session_started'
[1] tool_call tool_name='run_agent'
[2] user_message content='Book Tuesday at 2pm'
[3] tool_call tool_name='search_calendar'
[4] tool_result call_id='190a6438979141f5ac11b2e1b2ee29a0'
[7] agent_message content='Booked Wednesday 2pm'
[8] goal_check passed=False
[9] tool_result call_id='a78566297e0a4a309d5ce44cefe0d836'
[10] session_end status='completed'
choose-slot span=span_163a8380174647e98bfe1f3fff9e15b9 duration_ms=0.0
Generate a failure fingerprint:
promptetheus fingerprint 01KVMZ4T7V2SN61ZWG1XTDBK47
Example output:
8ae0f41220d0 goal mismatch: selected wednesday, not tuesday
- goal:selected wednesday, not tuesday
Inspect the local delivery spool:
promptetheus spool list
Example output:
Spool: .promptetheus/spool
pending : 11 event(s) across 1 session file(s), 4082 bytes
dead : 0 event(s) across 0 file(s), 0 bytes
01KVMZ4T7V2SN61ZWG1XTDBK47: 11 pending
The raw spool is JSONL. Each line is an event envelope:
{
"type" : " tool_call " ,
"session_id" : " 01KVMZ4T7V2SN61ZWG1XTDBK47 " ,
"seq" : 1 ,
"idempotency_key" : " 01KVMZ4T7V2SN61ZWG1XTDBK47:29c5eff0:1 " ,
"payload" : {
"tool_name" : " run_agent " ,
"call_id" : " a78566297e0a4a309d5ce44cefe0d836 " ,
"arguments" : {
"args" : " ('Book Tuesday at 2pm',) " ,
"kwargs" : " {} "
}
}
}
Manual Trace API
Use pt.trace.start(...) when you control the run boundary and want explicit
event calls instead of decorators:
import promptetheus as pt
with pt . trace . start (
agent = "demo-agent" ,
user_goal = "Book a meeting for Tuesday" ,
transport = "auto" ,
) as session :
session . user_message ( "Please book the small room for Tuesday at 2pm" )
session . tool_call ( "calendar.search" , { "day" : "Tuesday" }, call_id = "calendar-1" )
session . tool_result ( "calendar-1" , result = { "available" : [ "2pm" , "3pm" ]})
session . agent_message ( "Booking confirmed for Wednesday at 2pm" )
session . goal_check ( False , mismatches = [ "booked Wednesday, not Tuesday" ])
# session_end is emitted automatically; transport flush runs on exit
Public SDK API
The package exposes these primary entry points:
import promptetheus as pt
pt . trace . start (...)
pt . start (...)
pt . observe (...)
pt . tool
pt . traced (...)
pt . current ()
pt . Session
pt . AsyncSession
pt . AgentRuntime
Common session helpers:
session . user_message ( "Book Tuesday at 2pm Pacific" )
session . agent_message ( "I found availability" )
session . tool_call ( "browser.click" , { "selector" : "#checkout" }, call_id = "click-1" )
session . tool_result ( "click-1" , result = { "ok" : True })
session . retrieval ( "refund policy" , documents = [{ "id" : "doc-1" , "score" : 0.91 }])
session . browser_action ( "click" , "#checkout" , url = page . url )
session . dom_snapshot ( page . url , visible_text , selected_values = { "day" : "Tuesday" })
session . screenshot ( page . screenshot ())
session . replay_artifact ( "trace.webm" , artifact_type = "screen_recording" , event_time_map = {})
session . llm_call ( "gpt-5" , input_tokens = 100 , output_tokens = 40 , latency_ms = 900 )
session . score ( "goal_match" , 0.2 , comment = "Selected the wrong day" )
session . metric ( "steps" , 12 , unit = "count" )
session . error ( RuntimeError ( "calendar API timeout" ), handled = True )
session . goal_check ( False , mismatches = [ "selected Wednesday" ])
session . end ( "failed" )
session . flush ( timeout = 2 )
Every helper writes a schema-valid event envelope with type , session_id ,
timestamp , seq , idempotency_key , and payload . Use metadata for safe,
low-cardinality context. Do not put raw secrets, cookies, tokens, or
credentials into event payloads.
Use AsyncSession when the top-level agent run is async:
from promptetheus import AsyncSession
async with AsyncSession ( agent = "voice-agent" , user_goal = "Summarize the call" ) as session :
session . user_message ( "Summarize this call" )
async with session . aspan ( "transcribe" ):
session . metric ( "audio_seconds" , 42 , unit = "seconds" )
session . goal_check ( True )
Browser Agents
Browser agents should record the user goal, critical browser actions, the final
DOM state, and an explicit goal check:
session . browser_action ( "click" , "#confirm" , url = page . url )
session . dom_snapshot (
page . url ,
visible_text = await page . locator ( "body" ). inner_text (),
selected_values = { "day" : "Wednesday" , "time" : "2pm" },
warnings = [ "Timezone changed from Pacific to Eastern" ],
)
session . goal_check (
False ,
mismatches = [ "booked Wednesday" , "timezone warning visible" ],
)
This is the path that lets Promptetheus replay a failure and produce fix-agent
evidence instead of just storing generic logs.
Adapters are optional and imported lazily. Install only the extra you need:
pip install " promptetheus[openai] "
pip install " promptetheus[anthropic] "
pip install " promptetheus[langchain] "
pip install " promptetheus[playwright] "
Available adapter exports:
from promptetheus . adapters import (
AnthropicAdapter ,
AutoGenAdapter ,
CrewAIAdapter ,
DSPyAdapter ,
HaystackAdapter ,
LangGraphAdapter ,
LiteLLMAdapter ,
LlamaIndexAdapter ,
OpenAIAdapter ,
OpenTelemetryBridge ,
PlaywrightAdapter ,
PromptetheusCallbackHandler ,
PydanticAIAdapter ,
)
Use adapters when a framework already emits structured callbacks. Keep custom
instrumentation close to the real run boundary when the framework does not.
AgentRuntime is a best-effort client for live, service-backed coordination.
It is separate from durable trace storage and never raises into host code when
the service is unavailable:
from promptetheus import AgentRuntime
runtime = AgentRuntime ( session . session_id )
runtime . remember ( "hypothesis" , { "summary" : "auth

[truncated]
