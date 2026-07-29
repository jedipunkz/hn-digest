---
source: "https://github.com/oryxintel/oryxflow"
hn_url: "https://news.ycombinator.com/item?id=49099308"
title: "Oryxflow – cheaper and more reliable AI data analysis in Python and Claude Code"
article_title: "GitHub - oryxintel/oryxflow: Faster, cheaper, more trustworthy data analysis for humans and AI agents. A Python library that turns data-science scripts into a cached, reproducible, lineage-tracked pipeline — reruns only what a parameter, data, or code change affects. No server, no account. · GitHub"
author: "citynorman"
captured_at: "2026-07-29T17:06:13Z"
capture_tool: "hn-digest"
hn_id: 49099308
score: 2
comments: 0
posted_at: "2026-07-29T16:06:23Z"
tags:
  - hacker-news
  - translated
---

# Oryxflow – cheaper and more reliable AI data analysis in Python and Claude Code

- HN: [49099308](https://news.ycombinator.com/item?id=49099308)
- Source: [github.com](https://github.com/oryxintel/oryxflow)
- Score: 2
- Comments: 0
- Posted: 2026-07-29T16:06:23Z

## Translation

タイトル: Oryxflow – Python と Claude Code による安価で信頼性の高い AI データ分析
記事のタイトル: GitHub - oryxintel/oryxflow: 人間と AI エージェントのための、より速く、より安く、より信頼できるデータ分析。データ サイエンス スクリプトを、キャッシュされた再現可能なリネージ追跡パイプラインに変換する Python ライブラリ。パラメーター、データ、またはコードの変更が影響するもののみを再実行します。サーバーもアカウントもありません。 · GitHub
説明: 人間と AI エージェントにとって、より速く、より安く、より信頼できるデータ分析。データ サイエンス スクリプトを、キャッシュされた再現可能なリネージ追跡パイプラインに変換する Python ライブラリ。パラメーター、データ、またはコードの変更が影響するもののみを再実行します。サーバーもアカウントもありません。 - オリックスインテル/オリックスフロー

記事本文:
GitHub - oryxintel/oryxflow: 人間と AI エージェントにとって、より速く、より安く、より信頼できるデータ分析。データ サイエンス スクリプトを、キャッシュされた再現可能なリネージ追跡パイプラインに変換する Python ライブラリ。パラメーター、データ、またはコードの変更が影響するもののみを再実行します。サーバーもアカウントもありません。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
あなたは切り替えました

別のタブまたはウィンドウ上のアカウント。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
オリックスインテル
/
オリックスフロー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
57 コミット 57 コミット .github/ workflows .github/ workflows devops devops docs docs oryxflow oryxflow オーバーライド オーバーライド スクリプト スクリプト テスト テスト .gitignore .gitignore .readthedocs.yaml .readthedocs.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.mdライセンス ライセンス MANIFEST.in MANIFEST.in README.md README.md mkdocs.yml mkdocs.yml 要件-dev.txt 要件-dev.txt 要件.txt 要件.txt setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
人間と AI コーディング エージェントにとって、より速く、より安く、より信頼できるデータ分析が可能になります。
oryxflow は、データ サイエンス スクリプトを、変更が影響するものを正確に再実行するパイプラインに変換します。
各結果がどのように作成されたかを記録し、すべてのステップをキャッシュするので、同じ作業に対して 2 回支払うことがなくなります。
中間ファイルに名前を付けたり、どのパラメータがどの出力を生成したかを再度追跡したりすることはありません。
Pythonのライブラリです。サーバーもデータベースもアカウントも設定ファイルもありません。
pip インストール oryxflow
問題: 反復分析は静かに信頼できなくなる
ほとんどすべてのプロジェクトは、機能するスクリプトとして始まります。そして、失敗が積み重なって、
ワークフローが非効率になり、結果に対する信頼が失われます。
無駄な再計算。ダウンストリームで 1 行変更すると、10 分間のデータ プルが再実行されるため、
os.path.exists(...) キャッシュ自体が古くなった場合は、待つかハンドローリングを開始します。
古くなった中間体。機能を変更し、キャッシュされたファイルを再生成するのを忘れてトレーニングした場合
yに

昨日のデータ。何もエラーはありません。数字が間違っているだけです。
整理しておかなければならないファイルが詰まったフォルダー。 features_v3.pkl 、
features_v3_final.pkl 、 features_v3_final_FIXED.pkl — それに加えて、
それぞれが使用した設定であり、誰も書き留めていませんでした。
失われた血統。 6 か月 (または 6 時間) 後、どのコードがどの入力であったかは誰も言えなくなります。
model_final_v3.pkl が生成されました。
AI が生成したコードは完全には信頼できません。コーディングエージェントはもっともらしいパンダを書き、
scikit-learn は高速ですが、長いセッションを続けるとリネージや既に計算された内容を追跡できなくなります。
それがまだ有効かどうかを確認し、古い状態に基づいて静かに構築します。
その結果、パイプラインの仕組みにおける信頼エラーが発生します。そして
プロジェクトが複雑になり、AI エージェントが記述するコードの量が増えるにつれて、改善されるどころか悪化していきます。
解決策: oryxflow が提供するもの
デフォルトでの再現性。すべての出力は正確なタスク、パラメータ、コードに関連付けられます
それを作成したバージョン。 「先週の結果を再現できますか?」機械的には「はい」になります。
クエリできる系統。 oryxflow は、いつ、どのパラメータとコードで何を実行したかを記録します。
そして再計算した理由。 「これは古いですか? 現在のコードで構築されていますか?」クエリではなく、
推測します。
変更されたものを正確に再実行します。パラメータ、データ入力、またはタスクのコードを正確に変更します
影響を受ける出力は再構築されます。古いフィーチャで新しいモデルを誤って評価することはできません。
ストレージやパラメータのボイラープレートはありません。ファイルに名前を付けたり、パスを構築したり、どのファイルを追跡したりすることはありません。
設定によってどの出力が生成されるか。 self.save(df) はそれを保管し、 flow.outputLoad() はそれを戻します。
oryxflow はタスクとそのパラメータからどこに存在するかを計算します。
スピードとコストの削減。完了したステップは再計算ではなくキャッシュからロードされるため、
編集と実行のループが数分から数秒に短縮されます。 AIエージェントが停止する

期限内にトークンを支払って、
すでに行った高価な作業をやり直します。
AI エージェントの信頼性。同じキャッシュとリネージ ログがエージェントのメモリになります。
セッション。コンパニオンの Claude Code プラグインには、これらの分野が含まれています
自動アクティブ化スキルとして機能するため、エージェントは古いキャッシュを信頼するのではなく、キャッシュを正しく使用します。
状態。
キャッシングがエンジンです。信頼 — 正確に何を更新する、再現可能で系統追跡された再実行
変更されました - 結果です。
30 秒の例: 2 つのモデル、古い比較の可能性はありません
オリックスフローをインポートする
パンダをPDとしてインポートする
スクラーンをインポートします。データセット 、 sklearn 。アンサンブル、スクラーン。線形モデル
クラス GetData ( oryxflow .tasks .TaskPqPandas ):
"""生のトレーニング データを 1 回ロードします。"""
持続 = [ 'x' , 'y' ]
def run (self):
ds = スクラーン 。データセット 。負荷_糖尿病 ()
自分自身。 save ({ 'x' : pd . DataFrame ( ds . data , columns = ds . feature_names ),
'y' : pd . DataFrame ( ds . target , columns = [ 'target' ])})
@ オリックスフロー 。 Required ( GetData ) # 依存関係を宣言する
クラス ModelTrain ( oryxflow .tasks .TaskPickle ):
"""1 つのモデルを当てはめます。「モデル」は出力の ID の一部です。"""
モデル = オリックスフロー 。 ChoiceParameter (デフォルト = 'ols' 、選択肢 = [ 'ols' , 'gbm' ])
def run (self):
df_x 、 df_y = self 。 inputLoad () # GetData の出力、すでにロードされています
clf = { 'ols' : sklearn .線形モデル 。線形回帰 、
'gbm' : スクラーン 。アンサンブル。 GradientBoostingRegressor }[ self .モデル]()
clf 。 fit ( df_x , df_y .values .ravel())
自分自身。保存 ( clf ) # キャッシュ;ファイル名を作成する必要はありません
自分自身。 saveMeta ({ 'スコア' : clf . スコア ( df_x , df_y )})
# 2 つの名前付き実験;それぞれの名前は、その実行を定義するパラメータにマップされます。
フロー = オリックスフロー 。 WorkflowMulti ( ModelTrain , { 'ols' : { 'model' : 'ols' },
'gbm' : { 'モデル' : 'gbm' }})
結果＝流れ。 run() # 両方を実行します。

依存関係の順序で
print ( result . summary ()) # 何が実行され、何がキャッシュから来たのか
print ( flow .outputLoadMeta ()) # 実験名をキーとした結果
# {'ols': {'スコア': 0.5177484222203498}, 'gbm': {'スコア': 0.7990392018966864}}
その辞書内の各キーは実験に名前を付け、その値はその実行のパラメーターを設定します。
タスクが宣言するパラメーターなので、「gbm」はコードを編集せずに ModelTrain(model='gbm') をトレーニングします。
走る。結果は同じ名前で返され、3 番目のモデルを追加するのはもう 1 行です。掃き掃除用
いくつかの名前付き実行ではなく、値を渡すと、oryxflow がグリッド自体を拡張します。
WorkflowMulti(ModelTrain, params={'model': ['ols', 'gbm']}) 。
2 番目の実験についての概要を見てください。
===== gbm =====
スケジュールされたタスクのうち 2 つ:
* 1 つの完全なものが見つかりました:
- GetData <- キャッシュからロードされ、再計算されません
* 1 は正常に実行されました。
- ModelTrain(モデル=gbm)
GetData は 1 回実行され、再利用されました。 flow.run() を再度実行しても、何も起こりません。
タスクはキャッシュ ヒットです。明日 3 番目のモデルを追加すると、その 1 つだけがトレーニングされます。
次に、GetData の本体を編集します。機能を追加し、フィルターを修正し、新しいソースを指定します。
===== オルス =====
スケジュールされたタスクのうち 2 つ:
* 0 個の完全なものが見つかりました
* 2 は正常に実行されました。
- GetData <- コードが変更されたため、再実行されました
- ModelTrain(model=ols) <- 入力が変更されたため、再実行されました
両方のモデルが再トレーニングされ、両方のスコアが変化しました。あなたはそれを求めていません、そしてあなたは努力する必要はありません
キャッシュされた結果のうち、まだ有効なものはどれですか。oryxflow は、コード、パラメーター、入力から有効です。
これが、これを自分で管理するキャッシュと区別するものです。出力する比較は、
現在のコードを比較し、昨日の機能と今日の機能を混合することは決してありません。 (コードを再フォーマットします
またはコメントを追加しても何も再実行されません

— コードの動作を比較します。そして最後に取った一歩
10 分を超えると、サイレントに再計算する代わりに警告が表示されるため、リファクタリングは静かに長い時間を書き込むことができません。
走ってください。）
それが全体的な考え方です。キャッシュはその仕組みです。信頼はあなたが得るものです。
そしてそれは何にでも使えます。タスクの run() は単なる Python なので、あらゆる ML ライブラリ (sklearn、PyTorch、
Keras、XGBoost) とあらゆるデータ スタックがその内部で動作します。 oryxflow はグラフとストレージを管理します。
あなたの数学ではありません。出力は、parquet、pickle、CSV、JSON、Excel、Markdown、またはメモリ内キャッシュとして保存されます
— ローカルまたは S3/GCS 上で。
簡単な EDA から始めて、そこから拡張します
初日からタスクグラフを作成する必要はなく、事前に決定する必要もありません。プレーンから始めましょう
スクリプトまたは探索的プローブ。作業のステップ、コスト、パラメータの組み合わせが増加するにつれて、
常に実行します — /oryxflow:maigrate は、キャッシュされたタスクに既に書き込まれた内容を持ち上げます。簡単なスクリプトは次のとおりです。
最初は任意の複雑さで、途中で書き直しや崖はありません。
すでに中止になったプロジェクトがありますか? 9 つのノートブックと clean_v3.csv のフォルダーは
まさにその出発点
乱雑なノートブック プロジェクトを移行するのは、
のために書かれた。
データ作業におけるコーディング エージェントの本当の弱点は、パンダを書くことではなく、パンダが何を書いているかを追跡することです。
すでに実行されているもの、まだ有効なもの、そして最後の編集が無効になったもの。それはあまりにも重要なので、
コンテキストウィンドウのままにしておきます。
oryxflow クロード コード プラグインは、
コーディング エージェントがこの方法で分析を構築するため、データを書き込む代わりに高価な結果を再利用します。
再実行には時間とトークンが必要であり、古いデータでモデルを静かにトレーニングすることはできません。オリックスフローは
エージェントの外部で実行されたものの記録。プラグインはエージェントに、代わりにそのレコードをチェックするように指示します。
自分の記憶を信じて。これはプラグイン (スキルとスラッシュ コマンド) であり、

MCPサーバー。
/plugin マーケットプレイス oryxintel/oryxflow-claude-plugin を追加
/プラグインのインストール oryxflow@oryxflow
oryxflow スキルは、パイプライン プロジェクトで作業するたびに自動的にアクティブになります。スラッシュコマンドは次のとおりです。
詳細: データ サイエンスのための Claude Code 。
oryxflow は実験トラッカーや運用オーケストレーターに代わるものではなく、ギャップを埋めるものです
アドホック スクリプトと重量級プラットフォームの間で使用され、両方で構成されます。特徴的なのは、
ローカルファーストのシンプルさ、コードの変更を認識する無効化、および
常にオンの血統。
vs Airflow / Prefect / Dagster — 別の仕事。これらは、スケジュールされた本番パイプラインを実行します。
実際のインフラストラクチャ。 oryxflow は、ローカル調査ループ用の pip インストールです。
続きを読む→
vs MLflow / W&B — 補完的であり、両方を使用することを想定する必要があります。トラッカーは「どれが実行されるか」と答えます
スコアは0.91?"; oryxflow は、「再現するにはどのステップを再実行する必要があるか、そしてその入力は何か」と答えます。
古くなった？」 oryxflow タスク内でトラッカーにログを記録し続けます。
続きを読む→
vs DVC — どちらもキャッシュ パイプラインです。 DVC はファイルと YAML で宣言されたステージをハッシュします。オリックスフローが維持する
ネイティブ Python の ID であるため、パラメーターの変更は自動的に新しいキャッシュされた結果となり、コードになります。
edit は、影響を受けるタスクを独自に再実行します。構成ファイルを維持する必要はありません。
続きを読む→
フルLAN

[切り捨てられた]

## Original Extract

Faster, cheaper, more trustworthy data analysis for humans and AI agents. A Python library that turns data-science scripts into a cached, reproducible, lineage-tracked pipeline — reruns only what a parameter, data, or code change affects. No server, no account. - oryxintel/oryxflow

GitHub - oryxintel/oryxflow: Faster, cheaper, more trustworthy data analysis for humans and AI agents. A Python library that turns data-science scripts into a cached, reproducible, lineage-tracked pipeline — reruns only what a parameter, data, or code change affects. No server, no account. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Uh oh!
There was an error while loading. Please reload this page .
oryxintel
/
oryxflow
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
57 Commits 57 Commits .github/ workflows .github/ workflows devops devops docs docs oryxflow oryxflow overrides overrides scripts scripts tests tests .gitignore .gitignore .readthedocs.yaml .readthedocs.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md mkdocs.yml mkdocs.yml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt setup.py setup.py View all files Repository files navigation
Faster, cheaper, more trustworthy data analysis — for humans and AI coding agents.
oryxflow turns a data-science script into a pipeline that reruns exactly what a change affects,
records how each result was made, and caches every step so you never pay twice for the same work.
You never name an intermediate file or track which parameters produced which output again.
It's a Python library. No server, no database, no account, no config files.
pip install oryxflow
The problem: iterative analysis quietly stops being trustworthy
Almost every project starts as a script that works. Then it accumulates the failures that
make your workflow inefficient and erode trust in the result:
Wasted recomputation. A one-line change downstream re-runs the 10-minute data pull, so you
either wait or start hand-rolling if os.path.exists(...) caches that themselves go stale.
Stale intermediates. You change a feature, forget to regenerate a cached file, and train
on yesterday's data. Nothing errors. The number is just wrong.
A folder full of files you have to keep straight. features_v3.pkl ,
features_v3_final.pkl , features_v3_final_FIXED.pkl — plus the mental note about which
settings each one used, which nobody wrote down.
Lost lineage. Six months (or six hours) later, no one can say which code and which inputs
produced model_final_v3.pkl .
AI-generated code you can't fully trust. Coding agents write plausible pandas and
scikit-learn fast — but across a long session they lose track of lineage and what's already computed and
whether it's still valid, and silently build on stale state.
The result are trust errors — in the mechanics of the pipeline. And
they get worse, not better, as projects grow in complexity and an AI agent writes more of the code.
The solution: What oryxflow gives you
Reproducibility by default. Every output is tied to the exact task, parameters, and code
version that produced it. "Can I reproduce last week's result?" becomes yes, mechanically.
Lineage you can query. oryxflow records what ran, when, with which parameters and code,
and why it recomputed. "Is this stale? Was it built with current code?" are queries, not
guesses.
Reruns exactly what changed. Change a parameter, a data input, or a task's code and exactly
the affected outputs rebuild — you can't accidentally evaluate a new model on old features.
No storage or parameter boilerplate. You never name a file, build a path, or track which
settings produced which output. self.save(df) puts it away, flow.outputLoad() gets it back,
and oryxflow works out where it lives from the task and its parameters.
Speed and cost savings. Completed steps load from cache instead of recomputing, so the
edit–run loop drops from minutes to seconds. An AI agent stops paying — in time and tokens — to
redo expensive work it already did.
AI-agent reliability. The same cache and lineage log become an agent's memory across
sessions. The companion Claude Code plugin ships these disciplines
as an auto-activating skill, so the agent uses the cache correctly instead of trusting stale
state.
Caching is the engine . Trust — reproducible, lineage-tracked reruns that update exactly what
changed — is the outcome .
30 second example: two models, no chance of a stale comparison
import oryxflow
import pandas as pd
import sklearn . datasets , sklearn . ensemble , sklearn . linear_model
class GetData ( oryxflow . tasks . TaskPqPandas ):
"""Load the raw training data once."""
persists = [ 'x' , 'y' ]
def run ( self ):
ds = sklearn . datasets . load_diabetes ()
self . save ({ 'x' : pd . DataFrame ( ds . data , columns = ds . feature_names ),
'y' : pd . DataFrame ( ds . target , columns = [ 'target' ])})
@ oryxflow . requires ( GetData ) # declare the dependency
class ModelTrain ( oryxflow . tasks . TaskPickle ):
"""Fit one model. `model` is part of the output's identity."""
model = oryxflow . ChoiceParameter ( default = 'ols' , choices = [ 'ols' , 'gbm' ])
def run ( self ):
df_x , df_y = self . inputLoad () # GetData's output, already loaded
clf = { 'ols' : sklearn . linear_model . LinearRegression ,
'gbm' : sklearn . ensemble . GradientBoostingRegressor }[ self . model ]()
clf . fit ( df_x , df_y . values . ravel ())
self . save ( clf ) # cached; no filename to invent
self . saveMeta ({ 'score' : clf . score ( df_x , df_y )})
# two named experiments; each name maps to the parameters that define that run
flow = oryxflow . WorkflowMulti ( ModelTrain , { 'ols' : { 'model' : 'ols' },
'gbm' : { 'model' : 'gbm' }})
result = flow . run () # runs both, in dependency order
print ( result . summary ()) # what ran, and what came from cache
print ( flow . outputLoadMeta ()) # results keyed by experiment name
# {'ols': {'score': 0.5177484222203498}, 'gbm': {'score': 0.7990392018966864}}
Each key in that dict names an experiment and its value sets that run's parameters — matched to the
parameters the task declares, so 'gbm' trains ModelTrain(model='gbm') with no code edit between
runs. Results come back under the same names, and adding a third model is one more line. For a sweep
rather than a few named runs, pass the values and oryxflow expands the grid itself:
WorkflowMulti(ModelTrain, params={'model': ['ols', 'gbm']}) .
Look at what the summary says about the second experiment:
===== gbm =====
Scheduled 2 tasks of which:
* 1 complete ones were encountered:
- GetData <- loaded from cache, not recomputed
* 1 ran successfully:
- ModelTrain(model=gbm)
GetData ran once and was reused. Run flow.run() again and nothing happens at all — every
task is a cache hit. Add a third model tomorrow and only that one trains.
Now edit the body of GetData — add a feature, fix a filter, point it at a new source:
===== ols =====
Scheduled 2 tasks of which:
* 0 complete ones were encountered
* 2 ran successfully:
- GetData <- its code changed, so it reran
- ModelTrain(model=ols) <- its input changed, so it reran
Both models retrained and both scores moved. You didn't ask for that, and you don't have to work out
which cached results are still valid — oryxflow does, from the code, the parameters, and the inputs.
That's what separates this from a cache you manage yourself: the comparison you print is a
comparison of your current code, never a mix of yesterday's features and today's. (Reformat the code
or add a comment and nothing reruns — it compares what the code does . And a step that last took
over ten minutes warns instead of silently recomputing, so a refactor can't quietly burn a long
run.)
That's the whole idea. Caching is how it works; trust is what you get.
And it Works with anything. A task's run() is just Python, so any ML library (sklearn, PyTorch,
Keras, XGBoost) and any data stack works inside it. oryxflow manages the graph and the storage,
not your math. Outputs save as parquet, pickle, CSV, JSON, Excel, Markdown, or an in-memory cache
— locally or on S3/GCS.
Start with a quick EDA — it scales from there
You don't need a task graph on day one, and you don't have to decide upfront. Start with a plain
script or an exploratory probe; as the work gains steps, cost, and parameter combinations — as it
always does — /oryxflow:migrate lifts what you already wrote into cached tasks. Simple scripts at
the start, arbitrary complexity later, with no rewrite and no cliff in between.
Already have a project that got away from you? Nine notebooks and a folder of clean_v3.csv is
exactly the starting point that
Migrate a messy notebook project is
written for.
A coding agent's real weakness in data work isn't writing the pandas — it's keeping track of what it
already ran, what's still valid, and what its last edit just invalidated. That's too important to
leave in a context window.
The oryxflow Claude Code plugin teaches your
coding agent to build the analysis this way — so it reuses expensive results instead of burning your
time and tokens redoing them, and can't quietly train a model on stale data. oryxflow keeps the
record of what ran outside the agent; the plugin teaches the agent to check that record instead of
trusting its own memory. It's a plugin (a skill plus slash commands), not an MCP server.
/plugin marketplace add oryxintel/oryxflow-claude-plugin
/plugin install oryxflow@oryxflow
The oryxflow skill auto-activates whenever you work in a pipeline project. The slash commands:
More: Claude Code for data science .
oryxflow doesn't replace an experiment tracker or a production orchestrator — it fills the gap
between an ad-hoc script and a heavyweight platform, and composes with both. What's distinctive is
the combination of local-first simplicity, invalidation that notices a code change, and
always-on lineage.
vs Airflow / Prefect / Dagster — a different job. Those run scheduled production pipelines on
real infrastructure; oryxflow is a pip install for the local research loop.
Read more →
vs MLflow / W&B — complementary, and you should expect to use both. Trackers answer "which run
scored 0.91?"; oryxflow answers "which steps must rerun to reproduce it, and are its inputs
stale?" Keep logging to your tracker inside oryxflow tasks.
Read more →
vs DVC — both cache pipelines. DVC hashes files and YAML-declared stages; oryxflow keeps
identity in native Python, so a parameter change is a new cached result automatically and a code
edit reruns the affected tasks on its own — no config files to maintain.
Read more →
The full lan

[truncated]
