---
source: "https://github.com/Idank96/agent-panorama"
hn_url: "https://news.ycombinator.com/item?id=48531114"
title: "Nobody Is Measuring What Your AI Agents Are Worth"
article_title: "GitHub - Idank96/agent-panorama: See what your AI agents do, whether it's worth it, and what it costs - a manager-readable report + local dashboard from Langfuse/LangSmith traces (or a one-line live callback). Open source, runs locally. · GitHub"
author: "Idankogan"
captured_at: "2026-06-14T21:33:57Z"
capture_tool: "hn-digest"
hn_id: 48531114
score: 1
comments: 1
posted_at: "2026-06-14T18:50:29Z"
tags:
  - hacker-news
  - translated
---

# Nobody Is Measuring What Your AI Agents Are Worth

- HN: [48531114](https://news.ycombinator.com/item?id=48531114)
- Source: [github.com](https://github.com/Idank96/agent-panorama)
- Score: 1
- Comments: 1
- Posted: 2026-06-14T18:50:29Z

## Translation

タイトル: AI エージェントの価値を誰も測定していない
記事のタイトル: GitHub - Idank96/agent-panorama: AI エージェントが何を行うか、その価値があるかどうか、そしてそれにかかるコストを確認します。マネージャーが読めるレポートと、Langfuse/LangSmith トレース (または 1 行のライブ コールバック) からのローカル ダッシュボードです。オープンソースでローカルで動作します。 · GitHub
説明: AI エージェントが何を行うか、その価値があるかどうか、そしてそれにかかるコストを確認します。マネージャーが読めるレポートと、Langfuse/LangSmith トレース (または 1 行のライブ コールバック) からのローカル ダッシュボードです。オープンソースでローカルで動作します。 - Idank96/エージェントパノラマ

記事本文:
GitHub - Idank96/agent-panorama: AI エージェントが何を行うか、その価値があるかどうか、そしてそれにかかるコストを確認します。マネージャーが読めるレポートと、Langfuse/LangSmith トレース (または 1 行のライブ コールバック) からのローカル ダッシュボードです。オープンソースでローカルで動作します。 · GitHub
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
{{ メッサ

げ }}
イダンク96
/
エージェントのパノラマ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット .github/ workflows .github/ workflows アセット アセットの例 例 フロントエンド フロントエンド src/ Agent_panorama src/ Agent_panorama テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md config.example.yaml config.example.yaml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
生の LLM エージェント トレースを、マネージャーが実際に読み取ることができるレポートに変換します。
エージェントがやったこと、それが価値があるかどうか、そしてその費用はいくらだったのか。それをラングヒューズに向けます（または
LangSmith) をエクスポート (または 1 行のライブ コールバックを追加) し、クリーンな Markdown + を取得します。
わかりやすいビジネス言語による自己完結型の HTML レポートとローカル ダッシュボード。
スイッチをオンにするのは 1 行です。エンジニアは 1 つのコールバックをユーザーにドロップします。
既存のエージェント - 再構築も、新しいインフラストラクチャも、接続するトレースも必要ありません - そして
ライブ ダッシュボードが表示され始めます。現在、どの LangChain または LangGraph アプリでも動作します。
(ロードマップにはさらに多くのフレームワークが含まれています)。
運用中のエージェントに関する 3 つの質問。既存のツールが次の課題に答えます
最初の 2 つ:
3 番目の質問、つまり CEO、クライアント、または PM が実際に尋ねている質問に答えます。
同じ会話に関する 3 つのラング:
明晰さ - 「彼らは何をしたのですか？」フリート全体で 1 つの平易な英語フィード (または
単一エージェント): X に質問 → Y を実行 → 結果。 30 メッセージのチャットが 1 行になります。いいえ
スパン、JSON なし。
価値 - 「それだけの価値がありましたか?」 LLM 審査員が各会話を採点します。
価値の定義 (ドメイン、ユーザーの目標、成功基準)
そして、提供された価値、失われた価値、および修正すべき内容を報告します。
費用 - 「いくらかかりましたか?」トークン →

ドル → 貴重品あたりのコスト
会話、誰も教えてくれない ROI 数値。
フリート ビュー - すべてのエージェントにわたる、実行ごとの詳細、結果、コストを含む 1 つの平易な英語のアクティビティ フィード。
YAML を使用せずに値を定義する - ガイド付きウィザードが各エージェントの値オントロジーをライブ コンスタレーション マップとして入力し、その後、値ブループリントがそれを要約します。
トレースはエンジニアにとっては素晴らしいものですが、他の人にとっては最悪です。エージェントのパノラマ
ツールの呼び出し、再試行、トークンの使用法、およびエラーを平易な英語に翻訳します。それ
また、実際のユーザーリクエストと最終的な回答を LangGraph/LangChain から取得します。
メッセージ ペイロードなので、レポートは JSON ダンプではなくストーリーのようになります。
get_weather({"city": "パリ"}) → "天気を調べました"
モデル呼び出しが 3 回失敗しました → 「再試行回数が多い: 完了するまでに 3 回失敗しました。」
human_handoff(...) → 人間がエスカレートした実行結果
トークンが主要な指標です。 USD コストはオプトインです (v0.2 以降):
構成内のmodel_pricesテーブルとレポートにドルの見積もりが追加されます
トークンと一緒に（価格なし ⇒ コストは非表示のままです）。
pip インストール エージェント パノラマ
# または、ローカル開発の場合:
uv pip install -e " .[dev] "
Python 3.10以降が必要です。依存関係は意図的に最小限に抑えられています。クリックして、
jinja2 、 pyyaml 、 python-dotenv 。
クイックスタート - エージェントをライブで視聴する
最も早い方法: 既存の LangChain / LangGraph にコールバックを 1 つドロップする
アプリを使用して、完了したすべての実行がローカル ダッシュボードに表示されるのを確認します。再構築も新規もありません
インフラストラクチャ、エクスポートする痕跡なし。
1. Live Extra を使用してインストールします。
pip インストール「エージェント-パノラマ[ライブ]」
2. コールバックをエージェントに追加します。呼び出した場所は 1 行です。
Agent_panorama から。ライブ インポート PanoramaCallbackHandler
代理店。 invoke ( inputs , config = { "callbacks" : [ PanoramaCallbackHandler ()]})
それが全体の統合です。ハンドラーはベースパックに同梱されています

影とポスト
標準ライブラリ上で実行されるため、エージェント アプリはサーバーを必要としません
依存関係 - そして配信は 2 秒のタイムアウトを超えて発生したりブロックされることはありません (
ダッシュボードがダウンしても、アプリは 1 つの警告を記録し、動作し続けます)。
3. ダッシュボードを実行し、以下の実行ストリームを監視します。
Agent-panoramaserve --open # ダッシュボード (http://localhost:8321)
各実行は終了後数秒以内にアクティビティ フィードに表示されます (結果、
ツール呼び出し、トークン、異常、エージェントごとのロールアップはすべてライブで更新されます (
ダッシュボードは 3 秒ごとに /api/report をポーリングします)。
マルチターンチャットエージェント?メタデータ = {"session_id": ..., "user_id": ...} を渡します
invoke config により、会話全体が 1 つのフィード エントリに集約されます - を参照してください。
セッション。エクスポートされたものから作業することを好む
Langfuse / LangSmith はライブではなくトレースしますか?以下の CLI の使用法を参照してください。
ライブ サーバー (フラグ、安全性、デモ) の詳細は、ライブ モードでご覧ください。
Langfuse または LangSmith トレースをすでにエクスポートしましたか?ワンショットを生成する
Markdown + HTML レポート (およびダッシュボードの report.json) を直接作成 -
コールバックもライブサーバーもありません:
エージェントパノラマ生成 --input traces.json --output ./report --format html
オプション:
バンドルされたサンプルで試してみるか、フリート全体を集約してください。
エージェントパノラマ生成 --input 例/langfuse_traces.json --output ./report
# 多数のトレース → 1 つのフリート レポート + ダッシュボードの feed.json
エージェントパノラマ生成 --input ' トレース/*.json ' --input more/ \
--since 2026-05-01 --until 2026-05-31 --format json --output ./report
フリートビュー (v0.2)
複数の --input フラグ、グロブ パターン、およびディレクトリはすべて展開され、
1 つのレポートに集約されます。レポートにはエージェント間のアクティビティが含まれます。
フィードおよびエージェントごとのロールアップ (実行、アクション、成功/エスカレーション/再試行率、
トークン、model_prices が設定されている場合のコスト)。 --format json は、
st を含む report.json

有効な契約 ( generated_at 、 time_range 、 totals 、
feed 、 rollups 、 Decision_log ) はフロントエンド ダッシュボードによって消費されます。
構成にmodel_pricesテーブルを追加して、トークンの隣にドルの見積もりを取得します
(価格は 100 万トークンあたりの米ドルです。キーは部分文字列、最長でモデル名と一致します)
試合の勝利):
モデル価格 :
gpt-4o-mini : { 入力: 0.15、出力: 0.60 }
gpt-4o : { 入力: 2.50、出力: 10.00 }
クロード-3-5-ソネット : { 入力: 3.00、出力: 15.00 }
model_prices ブロックがない場合、コストは完全に省略され、トークンはそのまま残ります。
唯一のメトリック。
--format json は、安定した形状の report.json を書き込みます (入力も
ダッシュボードが消費します)。すべてのタイムスタンプは ISO-8601 UTC または
null ;すべての *cost_usd は数値または null (model_prices がない場合は null)
一致しました）。結果は、成功、人間によるエスカレーション、失敗、不明のいずれかです。
{
"generated_at" : " 2026-05-31T09:42:00+00:00 " ,
"time_range" : { "開始" : " … " 、 "終了" : " … " },
"合計" : { "実行" : 4 、 "ステップ" : 7 、 "トークン" : 3990 、 "cost_usd" : 0.0134 、
"value" : null }, // 値レイヤーがオンの場合の値の概要
"feed" : [ // 実行ごとに 1 つのエントリ、新しいものから順
{
"run_id" : " … " 、 "agent_name" : " 研究助手 " 、
"agent_key" : " Research-assistant " , // 安定した UI グループ化/色用のスラッグ
"action" : " Web を検索し、3 つの論文をまとめました。 " ,
"outcome" : " success " 、 "timestamp" : " … " 、
"retry_count" : 0 、 "anomaly_count" : 0 、
"トークン" : 1234 、 "コスト_米ドル" : 0.006 、
"概要" : " … " 、 "事実" : [[ " ステップ " 、 " 5 " ]、 [ " 再試行 " 、 " 0 " ]]、
「異常」: []、
"value" : null // 判定時の ValueJudgment (値レイヤーを参照)
}
]、
"rollups" : [ // エージェントごとに 1 つ
{
"agent_name" : " 研究アシスタント " 、 "agent_key" : " 研究アシスタント " 、
「実行」: 1 、「アクション」: 5 、
"成功率" : 1.0 、 "エスカレーション率" : 0.0 、
「失敗率」

" : 0.0 、 "retry_rate" : 0.0 、
"total_tokens" : 1234 、 "total_cost_usd" : 0.006 、
"judged" : 0 , "avg_value_score" : null , // 値レイヤーのメトリクス (オフの場合は null)
"value_rate" : null 、 "cost_per_valuable_usd" : null
}
]、
"decion_log" : [ // エージェント間の結果的なアクション
{ "タイムスタンプ" : " … " 、 "エージェント名" : " … " 、 "アクション" : " … " 、
"パラメータ" : " … " 、 "結果" : " 成功しました " }
】
}
図書館の利用状況
エージェントパノラマからインポート生成_レポート
レポート = 生成_レポート (
"トレース.json" 、
Output_dir = "./レポート" ,
フォーマット = [ "md" , "html" ],
input_type = "langfuse" ,
config = "config.yaml" , # オプション
)
print (レポート . total_runs 、レポート . total_tokens )
generate_report はメモリ内の Report を返すため、実行を検査することもできます。
決定ログ、およびディスクに触れることなくプログラムで異常を検出 (使用
ファイルを書き込まずにレポートが必要な場合は、build_report_from_file)。
generate_report (および下位レベルの build_report_from_inputs ) はグロブを受け入れます。
ディレクトリ、または inputs= を介したパスのリスト、および session / because / until
フィルター。返されたレポートでは、エージェント間フィードとエージェントごとのフィードが公開されます。
ロールアップ ; serialize_report は、report.json dict を直接提供します。
Agent_panorama インポートから (
generate_report 、 build_report_from_inputs 、load_runs 、
load_config、serialize_report、
)
レポート = 生成_レポート (
inputs = [ "traces/*.json" , "more/" ]、# globs、dirs、または単一パス
formats = [ "json" ], # report.json を書き込みます
以降 = "2026-05-01" 、まで = "2026-05-31" 、
config = "config.yaml" , #model_prices ここに ⇒ コストが入力されます
)
レポート内の項目について。 feed : # 新しい順のアクティビティ フィード
print ( item . エージェント名 、 アイテム . アクション 、 アイテム . 結果 . 値 、 アイテム . トークン 、 アイテム . コスト_米ドル )
レポートの r について。ロールアップ: # エージェントごとの成功/エスカレーション/再

お試し料金
print (r .agent_name , r .runs , r . success_rate , r . escalation_rate , r . retry_rate )
# ファイルがありませんか?メモリを構築し、JSON コントラクトを自分でシリアル化します。
実行 =load_runs ( "traces/*.json" 、セッション = "abc123" )
mem = build_report_from_inputs ( "traces/*.json" , "langfuse" , load_config ( "config.yaml" ))
payload = Serialize_report ( mem ,load_config ( "config.yaml" )) # -> dict
レポートの内容
概要 - 時間範囲、合計実行数、合計ステップ、合計トークン (および合計コスト)
model_prices が設定されている場合)。
フリート アクティビティ フィード (v0.2) - 実行ごとにスキャン可能な最新から最初の 1 行
すべてのエージェント: 誰が何をしたか、結果とタイミングをわかりやすく説明します。
エージェントごとのロールアップ (v0.2) - エージェントごとに 1 行: 実行、アクション、および
成功/エスカレーション/再試行率、さらにトークンとコスト。
エージェントごとのセクション - 何をするように求められたか、何を段階的に実行したか (グラフ
平易な英語でのノード/ツール呼び出し、選択された --detail レベル)、最終
結果、および信頼シグナル (再試行/フォールバック)。
意思決定ログ - すべての結果として生じるアクションのソート可能なテーブル: タイムスタンプ、
エージェント、アクション、平易な英語でまとめられたパラメータ、結果。
異常 - 再試行回数が多い、実行が遅い、アクティビティが多い、エラー、フォールバック。
すべての構成はオプションです。 「コン」を参照

[切り捨てられた]

## Original Extract

See what your AI agents do, whether it's worth it, and what it costs - a manager-readable report + local dashboard from Langfuse/LangSmith traces (or a one-line live callback). Open source, runs locally. - Idank96/agent-panorama

GitHub - Idank96/agent-panorama: See what your AI agents do, whether it's worth it, and what it costs - a manager-readable report + local dashboard from Langfuse/LangSmith traces (or a one-line live callback). Open source, runs locally. · GitHub
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
Idank96
/
agent-panorama
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .github/ workflows .github/ workflows assets assets examples examples frontend frontend src/ agent_panorama src/ agent_panorama tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md config.example.yaml config.example.yaml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Turn raw LLM agent traces into a report a manager can actually read - what your
agents did, whether it was worth it, and what it cost. Point it at a Langfuse (or
LangSmith) export (or add a one-line live callback) and get clean Markdown + a
self-contained HTML report - and a local dashboard - in plain business language.
It's one line to switch on. An engineer drops a single callback into your
existing agents - no rebuild, no new infrastructure, no traces to wire up - and
the live dashboard starts filling in. Works in any LangChain or LangGraph app today
(more frameworks on the roadmap ).
Three questions about any agent in production. Your existing tools answer the
first two:
It answers the third - the one your CEO, client, or PM actually asks - across
three rungs over the same conversations:
Clarity - "what did they do?" One plain-English feed across the fleet (or a
single agent): asked X → did Y → outcome. A 30-message chat becomes one line. No
spans, no JSON.
Value - "was it worth it?" An LLM judge scores each conversation against
your definition of value (your domain, your user goal, your success criteria)
and reports the value delivered, the value lost, and what to fix.
Cost - "what did it cost?" Tokens → dollars → cost per valuable
conversation , the ROI number nobody else gives you.
The fleet view - one plain-English activity feed across every agent, with per-run details, outcomes, and cost.
Define value without YAML - a guided wizard fills in each agent's value ontology as a live constellation map, then a Value Blueprint summarizes it.
Traces are great for engineers and terrible for everyone else. agent-panorama
translates tool calls, retries, token usage, and errors into plain English. It
also pulls the real user request and final answer out of LangGraph/LangChain
messages payloads, so the report reads like a story, not a JSON dump:
get_weather({"city": "Paris"}) → "Looked up the weather"
3 failed model calls → "High retry count: 3 failed attempts before completing."
human_handoff(...) → run outcome human-escalated
Tokens are the primary metric. USD cost is opt-in (since v0.2): supply a
model_prices table in your config and the report adds dollar estimates
alongside tokens (no prices ⇒ cost stays hidden).
pip install agent-panorama
# or, for local development:
uv pip install -e " .[dev] "
Requires Python 3.10+. Dependencies are intentionally minimal: click ,
jinja2 , pyyaml , python-dotenv .
Quickstart - watch your agents live
The fastest way in: drop one callback into your existing LangChain / LangGraph
app and watch every completed run land on a local dashboard. No rebuild, no new
infrastructure, no traces to export.
1. Install with the live extra:
pip install ' agent-panorama[live] '
2. Add the callback to your agent - one line, wherever you invoke :
from agent_panorama . live import PanoramaCallbackHandler
agent . invoke ( inputs , config = { "callbacks" : [ PanoramaCallbackHandler ()]})
That's the whole integration. The handler ships with the base package and posts
runs over the standard library, so your agent app never needs the server
dependencies - and delivery never raises or blocks beyond a 2 s timeout (if the
dashboard is down, the app logs one warning and keeps working).
3. Run the dashboard and watch runs stream in:
agent-panorama serve --open # dashboard at http://localhost:8321
Each run appears in the activity feed within seconds of finishing - outcome,
tool calls, tokens, anomalies, and per-agent rollups all update live (the
dashboard polls /api/report every 3 s).
Multi-turn chat agent? Pass metadata={"session_id": ..., "user_id": ...} in
the invoke config so a whole conversation collapses into one feed entry - see
Sessions . Prefer working from exported
Langfuse / LangSmith traces instead of live? See CLI usage below.
More on the live server (flags, safety, demos) is in Live mode .
Already have exported Langfuse or LangSmith traces? Generate a one-shot
Markdown + HTML report (and report.json for the dashboard) straight from them -
no callback, no live server:
agent-panorama generate --input traces.json --output ./report --format html
Options:
Try it on the bundled example, or aggregate a whole fleet:
agent-panorama generate --input examples/langfuse_traces.json --output ./report
# many traces → one fleet report + a feed.json for the dashboard
agent-panorama generate --input ' traces/*.json ' --input more/ \
--since 2026-05-01 --until 2026-05-31 --format json --output ./report
Fleet view (v0.2)
Multiple --input flags, glob patterns, and directories are all expanded and
aggregated into one report. The report then carries a cross-agent activity
feed and per-agent rollups (runs, actions, success/escalation/retry rates,
tokens, and cost when model_prices is set). --format json writes a
report.json with a stable contract ( generated_at , time_range , totals ,
feed , rollups , decision_log ) consumed by the frontend dashboard.
Add a model_prices table to your config to get dollar estimates next to tokens
(prices are USD per 1M tokens; keys match model names by substring, longest
match wins):
model_prices :
gpt-4o-mini : { input: 0.15, output: 0.60 }
gpt-4o : { input: 2.50, output: 10.00 }
claude-3-5-sonnet : { input: 3.00, output: 15.00 }
With no model_prices block, cost is omitted entirely and tokens remain the
only metric.
--format json writes a report.json with a stable shape (also the input the
dashboard consumes). Every timestamp is ISO-8601 UTC or
null ; every *cost_usd is a number or null (null when no model_prices
matched). outcome is one of success , human-escalated , failure , unknown .
{
"generated_at" : " 2026-05-31T09:42:00+00:00 " ,
"time_range" : { "start" : " … " , "end" : " … " },
"totals" : { "runs" : 4 , "steps" : 7 , "tokens" : 3990 , "cost_usd" : 0.0134 ,
"value" : null }, // value summary when the value layer is on
"feed" : [ // one entry per run, newest first
{
"run_id" : " … " , "agent_name" : " research-assistant " ,
"agent_key" : " research-assistant " , // slug, for stable UI grouping/colour
"action" : " Searched the web and summarized 3 papers. " ,
"outcome" : " success " , "timestamp" : " … " ,
"retry_count" : 0 , "anomaly_count" : 0 ,
"tokens" : 1234 , "cost_usd" : 0.006 ,
"summary" : " … " , "facts" : [[ " Steps " , " 5 " ], [ " Retries " , " 0 " ]],
"anomalies" : [],
"value" : null // ValueJudgment when judged (see value layer)
}
],
"rollups" : [ // one per agent
{
"agent_name" : " research-assistant " , "agent_key" : " research-assistant " ,
"runs" : 1 , "actions" : 5 ,
"success_rate" : 1.0 , "escalation_rate" : 0.0 ,
"failure_rate" : 0.0 , "retry_rate" : 0.0 ,
"total_tokens" : 1234 , "total_cost_usd" : 0.006 ,
"judged" : 0 , "avg_value_score" : null , // value layer metrics (null when off)
"valuable_rate" : null , "cost_per_valuable_usd" : null
}
],
"decision_log" : [ // consequential actions across agents
{ "timestamp" : " … " , "agent_name" : " … " , "action" : " … " ,
"parameters" : " … " , "outcome" : " succeeded " }
]
}
Library usage
from agent_panorama import generate_report
report = generate_report (
"traces.json" ,
output_dir = "./report" ,
formats = [ "md" , "html" ],
input_type = "langfuse" ,
config = "config.yaml" , # optional
)
print ( report . total_runs , report . total_tokens )
generate_report returns the in-memory Report , so you can also inspect runs,
the decision log, and anomalies programmatically without touching disk (use
build_report_from_file if you want the report without writing files).
generate_report (and the lower-level build_report_from_inputs ) accept a glob,
a directory, or a list of paths via inputs= , plus session / since / until
filters. The returned Report exposes the cross-agent feed and per-agent
rollups ; serialize_report gives you the report.json dict directly.
from agent_panorama import (
generate_report , build_report_from_inputs , load_runs ,
load_config , serialize_report ,
)
report = generate_report (
inputs = [ "traces/*.json" , "more/" ], # globs, dirs, or a single path
formats = [ "json" ], # writes report.json
since = "2026-05-01" , until = "2026-05-31" ,
config = "config.yaml" , # model_prices here ⇒ cost is populated
)
for item in report . feed : # newest-first activity feed
print ( item . agent_name , item . action , item . outcome . value , item . tokens , item . cost_usd )
for r in report . rollups : # per-agent success/escalation/retry rates
print ( r . agent_name , r . runs , r . success_rate , r . escalation_rate , r . retry_rate )
# No files? Build in memory and serialize the JSON contract yourself:
runs = load_runs ( "traces/*.json" , session = "abc123" )
mem = build_report_from_inputs ( "traces/*.json" , "langfuse" , load_config ( "config.yaml" ))
payload = serialize_report ( mem , load_config ( "config.yaml" )) # -> dict
What's in a report
Summary - time range, total runs, total steps, total tokens (and total cost
when model_prices is set).
Fleet activity feed (v0.2) - one scannable, newest-first line per run across
every agent: who did what, in plain English, with outcome and timing.
Per-agent rollups (v0.2) - one row per agent: runs, actions, and
success / escalation / retry rates, plus tokens and cost.
Per-agent section - what it was asked to do, what it did step by step (graph
nodes / tool calls in plain English, at the chosen --detail level), final
outcome, and a confidence signal (retries / fallback).
Decision log - a sortable table of every consequential action: timestamp,
agent, action, parameters summarized in plain English, outcome.
Anomalies - high retry counts, slow runs, high activity, errors, fallbacks.
All configuration is optional. See con

[truncated]
