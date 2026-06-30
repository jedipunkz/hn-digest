---
source: "https://github.com/NaiaLorente/datalog"
hn_url: "https://news.ycombinator.com/item?id=48731529"
title: "Mltrackr – ML experiment tracking in 2 lines, no server, no account"
article_title: "GitHub - NaiaLorente/datalog · GitHub"
author: "naialorente"
captured_at: "2026-06-30T12:37:39Z"
capture_tool: "hn-digest"
hn_id: 48731529
score: 1
comments: 0
posted_at: "2026-06-30T12:07:55Z"
tags:
  - hacker-news
  - translated
---

# Mltrackr – ML experiment tracking in 2 lines, no server, no account

- HN: [48731529](https://news.ycombinator.com/item?id=48731529)
- Source: [github.com](https://github.com/NaiaLorente/datalog)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T12:07:55Z

## Translation

タイトル: Mltrackr – 2 行での ML 実験追跡、サーバーなし、アカウントなし
記事タイトル: GitHub - NaiaLorente/datalog · GitHub
説明: GitHub でアカウントを作成して、NaiaLorente/datalog の開発に貢献します。

記事本文:
GitHub - NaiaLorente/データログ · GitHub
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
ナイアロレンテ
/
データログ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーと

ファイル
4 コミット 4 コミット .github .github アセット アセットの例 例 mltrackr mltrackr .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LAUNCH.md LAUNCH.md README.md README.md Demon.py Demon.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
2 行のコードで ML 実験を追跡します。サーバーがありません。アカウントがありません。設定はありません。
トレーニング ループを実行しています。どのハイパーパラメータが最も効果的かを知りたいと考えています。次のようなことはしたくないでしょう。
任意のサービスでアカウントを作成する
環境変数を構成する
mltrackr が答えです。それをインストールし、ループをラップし、美しいローカル ダッシュボードを開きます。終わり。
pip インストール mltrackr
これにより、mltrackr コマンドがインストールされ、mltrackr Python パッケージがインポートされます。
2. すぐに実行できるサンプルを生成する
python -m mltrackr init --framework plain -o Demon.py
ほとんどのシステムでは、mltrackr init が直接動作します。そうでない場合は、代わりに python -m mltrackr を使用してください。
3. デモを実行します (6 つの偽のトレーニング実行を作成します)
Pythonデモ.py
4. ターミナルで結果を検査する
Python -m mltrackrリスト
python -m mltrackr 最高の精度
python -m mltrackr は精度を提案します
5.ビジュアルダッシュボードを開く
Python -m mltrackr ui
次に、ブラウザで http://localhost:7000 を開きます。 Ctrl+C を押して停止します。
mltrackrをインポートする
mltrackr を使用して。 run ( "resnet-baseline" 、タグ = [ "cv" , "baseline" ]):
mltrackr 。ログ ( lr = 1e-3 、batch_size = 64 、オプティマイザー = "adam" )
範囲 ( 50 ) のエポックの場合:
loss , acc = train_one_epoch (モデル, データローダー)
mltrackr 。 log (損失 = 損失、精度 = acc、エポック = エポック)
mltrackr 。注 (「確実なベースライン - 次に lr=5e-4 を試してください」)
# 「mltrackr」がシステム上で直接動作する場合:
mltracker UI
mltrackr リスト
mltrackr の最高の精度
mltrackr は精度を提案します
mltrackr レポート
# そうでない場合 (Windows など)、次を使用します。
Python -m mltrackr ui
Python -m mltrackrリスト
python -m mltrackr 最高の精度
すべて

g はローカルの ~/.mltrackr/experiments.db に保存されます。単一の SQLite ファイル。これをコピーしてバックアップし、任意の SQLite ブラウザで開きます。
良い結果が出ましたか? mltrackr share precision を実行して、すぐに投稿できる Twitter/X または Hacker News の概要を生成します。 mltrackr によって時間を節約できた場合は、GitHub の ⭐ が大いに役立ちます。
本当の問題: モデルをハッキングしており、いくつかのメトリクスをログに記録したいと考えていますが、MLflow のセットアップには 15 分かかり、W&B はアカウントを作成してデータをクラウドに送信するよう求めています。したがって、メトリクスをテキスト ファイルに書き込むか、何も追跡しないことになります。そうなると、どのハイパーパラメータが機能したのか忘れてしまいます。次に、失敗した同じ実験を再度実行します。
mltrackr は、必要なときに実際に使用できる実験トラッカーです。
任意のループをラップします。任意の値を記録します。あらゆるフレームワークで動作します。
mltrackrをインポートする
mltrackr を使用して。 run ( "gpt-finetune" 、タグ = [ "nlp" 、 "v3" ]):
mltrackr 。ログ ( lr = 2e-5 、エポック = 3 、モデル = "gpt2" )
step の場合、列挙型のバッチ ( dataloader ):
損失 = モデル。 train_step (バッチ)
mltrackr 。 log (損失 = 損失.項目 ()、ステップ = ステップ)
✅ 美しいライブダッシュボード
mltracker UI
http://localhost:7000 で開きます。以下を備えた高速のダークモード シングルページ アプリです。
サイドバーにインライン スパークライン チャートを備えた検索可能な実行リスト
各指標が改善しているかどうかを示すトレンド指標 (↑ ↓)
選択した実行を並べて比較 (最良の値が強調表示されます)
グラデーション塗りつぶしを含む自動生成された時系列グラフ
最新の値が履歴範囲内のどこに位置するかを示すメトリクスの進行状況バー
グローバル統計ビュー - 成功率、最もログに記録されたメトリクス、実行タイムライン
5 秒ごとに自動更新 — トレーニング中に開いて更新を確認します
✅ ライブ異常検出 — 不正な実行を早期に発見します
mltrackr 。 configure_watch ( nan_check = True 、 divergence_window = 5 、 plateau_win

ダウ = 15 )
mltrackr を使用して。実行 (「トレーニング」):
範囲 ( 100 ) のエポックの場合:
mltrackr 。ログ (損失 = compute_loss ())
# 次の場合に自動的に警告します: 損失 → NaN、損失は 5 エポックにわたって発散、
# 15 エポックの損失プラトー (LR の調整を提案)
すでに失敗している実行で GPU 時間を無駄にするのはやめましょう。
mltrackr は精度を提案します
実行履歴を分析し、どのハイパーパラメータ値がより良い結果と統計的に相関しているかを示します。ブラックボックスはありません。次のようなわかりやすい英語の洞察が得られます。
最適な構成: lr=0.001 → 平均精度 0.943 (他の値の場合は 0.871、+8.2%)
次の実験:batch_size=128 を試してください — より大きなバッチは +5.1% の精度で相関しました
✅ 自動生成された実験レポート
mltrackr レポート --output results.md
以下を使用して、論文の準備ができたマークダウン レポートを生成します。
要約統計 (合計実行数、完了率、最適な構成)
時系列の実験タイムライン
主な結果 (自動的に計算)
オプションの AI ナラティブ: mltrackr report --ai (ローカル Ollama を使用、API キーなし)
✅ すぐに実行できるサンプルを生成する
mltrackr init # 単純な Python の例
mltrackr init --framework pytorch # PyTorch トレーニング ループ
mltrackr init --framework sklearn # scikit-learn グリッド検索
mltrackr init --framework keras # Keras コールバック
すぐに実行できる完全な作業スクリプトを生成します。
フレームワーク
どうやって
パイトーチ
トレーニング ループ内の mltrackr.log(loss=loss.item(), acc=acc)
scikit-learn
ハイパーパラメータループ内のmltrackr.log(**params, cv_score=score)
ケラス / TF
model.fit() の 1 つのファイル TrainlogCallback
ハグ顔
カスタム TrainerCallback — Examples/huggingface_example.py を参照
XGブースト / ライトGBM
評価コールバックにログインします
ジャックス / 亜麻
各トレーニングステップの最後にログを記録します
プレーンパイソン
数字を生み出すものなら何でも
完全な API リファレンス
mltrackrをインポートする
# ── 追跡 ─

───────────────────
mltrackr を使用して。 run ( "name" , tags = [ "tag1" , "tag2" ]) を run_id として実行します。
mltrackr 。 log ( precision = 0.95 , loss = 0.05 ) # キーと値のペアをログに記録します
mltrackr 。 note ( "Cosine LR スケジュールがとても役に立ちました" ) # プレーンテキストのメモを添付します
mltrackr 。 tag ( run_id , "production" ) # 事後にタグを追加
mltrackr 。 tag ( "experiment-name" , "best" ) # 名前でも機能します
# ── クエリ ──── ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ クエリ
実行 = mltrackr 。 get_runs () # すべての実行、最新のものから順
ベスト = mltrackr 。 get_best_run ( "accuracy" ) # 最終的な最高値
best_low = mltrackr 。 get_best_run ( "loss" , mode = "min" ) # 最低最終値
cmp = mltrackr 。 Compare_runs ( 1 , 2 , 3 ) # 実行辞書のリスト
# ── 異常検知 ──── ─ ─ ─ ─ ─ ─ ─ 異常検知
mltrackr 。 configure_watch (
nan_check = True 、 # NaN/Inf 値を警告する
divergence_window = 5 , # N ステップでメトリクスが発散する場合に警告する
plateau_window = 15 , # N ステップでメトリックがプラトーになった場合に警告する
有効 = True 、
)
# または一時的にコンテキストマネージャーを使用して:
mltrackr を使用して。ウォッチ ( divergence_window = 3 ):
# このブロックに対する監視を厳しくする
mltrackr 。対数 (損失 = 0.5)
# ── エクスポートと分析 ─--------------------------------------------------------------------------------------------------
mltrackr 。博覧会

rt_csv ( "結果.csv" )
mltrackr 。エクスポート_json ( "結果.json" )
mltrackr 。 generate_report ( "report.md" 、 use_ollama = False )
提案 = mltrackr 。 recommend ( "精度" 、モード = "最大" 、top_n = 3 )
mltrackr 。 clear_all () # すべてを削除します (元に戻せません)
# ── 構成 ──── ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ 構成
mltrackr 。 configure (verbose = False) # 各実行後に自動要約パネルを抑制します
CLI リファレンス
# ダッシュボード
mltrackr ui # localhost:7000 で開きます
mltrackr ui --port 8080 --no-browser # カスタム ポート、自動オープンなし
# 実行を検査する
mltrackr リスト # 豊富なテーブル、新しいものから順
mltrackr リスト -- 制限 50
mltrackr Compare 1 2 3 # メトリックを並べて比較
mltrackr 最高の精度 # メトリクスの最高の実行
mltrackr の最良の損失 --mode min
# 注釈を付ける
mltrackr タグ 42 プロダクション調整 # 実行するタグを追加 #42
mltrackr note 42 "コサインアニーリングを試してください" # 実行するノートを追加 #42
# 分析する
mltrackr stats # 集計統計
mltrackr 提案精度 # ハイパーパラメータの推奨事項
mltrackr は損失を提案します --mode min --top 5
# 生成する
mltrackr レポート # report.md を書き込む
mltrackr report -o results.md --ai # Ollama AI ナラティブ付き
mltrackr init --framework pytorch # サンプル スクリプトを生成します
# エクスポート/クリーン
mltrackr エクスポート --format csv -o data.csv
mltrackr エクスポート --format json -o data.json
mltrackr clear # すべて削除 (確認を求める)
# シェアする
mltrackr 共有精度 # Twitter/X + HN 対応投稿を生成
mltrackr share loss --mode min # メトリクスの場合、低いほど良い
仕組み
SQLite — ~/.mltrackr/experiments.db 。 1つのファイル。サーバーがありません。任意の SQLite ブラウザで検査してください。 cp でバックアップします。
Flask — ダッシュボード

ローカルの Flask サーバーです。 Vanilla JS、Chart.js、npm ゼロ、ビルドステップゼロ。
スレッドローカル状態 — 独自のスレッド内の各トレーニング ジョブは、分離された実行コンテキストを取得します。同時実験はうまくいきます。
Git 対応 — git rev-parse HEAD 経由で現在のコミット ハッシュをキャプチャします。 git リポジトリの外でサイレントにスキップされました。
監視フック — 異常検出はすべての log() 呼び出し内で実行されます。外部サービスなし、オフラインで動作します。
mltrackr init --framework pytorch -o train.py
Python train.py
mltracker UI
それが全体の流れです。コマンドは5つ。設定ゼロ。
ライブ異常検出 (configure_watch)
自動生成された実験レポート (mltrackr レポート、Ollama サポート)
ハイパーパラメータの提案 ( mltrackr 提案 )
クイックスタート サンプル ジェネレーター ( mltrackr init )
トレンド指標を備えたサイドバーのスパークライン チャート
詳細ビューのメトリクスの進行状況バーとトレンド矢印
フレームワークの例: PyTorch、scikit-learn、Keras、HuggingFace
mltrackr.log_artifact("model.pt") — メトリクスと一緒にファイルパスを保存します
ネイティブ PyTorch TrainlogCallback (pip でインストール可能なプラグイン)
VS Code 拡張機能 — ホバー時のインライン実行概要
mltrackrserve — 共有可能な読み取り専用ダッシュボード URL (ngrok/localtunnel)
共有 Git 追跡 SQLite を介したチーム同期
実行完了時の Slack / Discord Webhook
アイデアはありますか?機能リクエストをオープンするか、PR を送信してください。
COTRIBUTING.md を参照してください。 TL;DR: pip install -e 。 、変更を加え、PR を開きます。
タイプミス、ドキュメント、機能、バグ修正など、あらゆる貢献を歓迎します。
MIT — 好きなだけ、永久に使用できます。
読み込み中にエラーが発生しました。このページをリロードしてください。
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

Contribute to NaiaLorente/datalog development by creating an account on GitHub.

GitHub - NaiaLorente/datalog · GitHub
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
NaiaLorente
/
datalog
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github .github assets assets examples examples mltrackr mltrackr .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LAUNCH.md LAUNCH.md README.md README.md demo.py demo.py pyproject.toml pyproject.toml View all files Repository files navigation
Track ML experiments in 2 lines of code. No server. No account. No config.
You're running a training loop. You want to know which hyperparameters worked best. You don't want to:
Create an account on any service
Configure environment variables
mltrackr is the answer. Install it, wrap your loop, open a beautiful local dashboard. Done.
pip install mltrackr
This installs the mltrackr command and import mltrackr Python package.
2. Generate a ready-to-run example
python -m mltrackr init --framework plain -o demo.py
On most systems mltrackr init works directly. If not, use python -m mltrackr instead.
3. Run the demo (creates 6 fake training runs)
python demo.py
4. Inspect results in the terminal
python -m mltrackr list
python -m mltrackr best accuracy
python -m mltrackr suggest accuracy
5. Open the visual dashboard
python -m mltrackr ui
Then open http://localhost:7000 in your browser. Press Ctrl+C to stop.
import mltrackr
with mltrackr . run ( "resnet-baseline" , tags = [ "cv" , "baseline" ]):
mltrackr . log ( lr = 1e-3 , batch_size = 64 , optimizer = "adam" )
for epoch in range ( 50 ):
loss , acc = train_one_epoch ( model , dataloader )
mltrackr . log ( loss = loss , accuracy = acc , epoch = epoch )
mltrackr . note ( "Solid baseline - try lr=5e-4 next" )
# If 'mltrackr' works directly on your system:
mltrackr ui
mltrackr list
mltrackr best accuracy
mltrackr suggest accuracy
mltrackr report
# If not (e.g. Windows), use:
python -m mltrackr ui
python -m mltrackr list
python -m mltrackr best accuracy
Everything is saved locally in ~/.mltrackr/experiments.db . A single SQLite file. Copy it, back it up, open it in any SQLite browser.
Got good results? Run mltrackr share accuracy to generate a ready-to-post Twitter/X or Hacker News summary. If mltrackr saved you time, a ⭐ on GitHub goes a long way!
The real problem: you're hacking on a model, you want to log some metrics, but setting up MLflow takes 15 minutes and W&B wants you to create an account and send your data to the cloud. So you end up writing metrics to a text file or just... not tracking anything. Then you forget which hyperparameters worked. Then you run the same failed experiment again.
mltrackr is the experiment tracker that's actually available when you need it.
Wrap any loop. Log any value. Works with every framework.
import mltrackr
with mltrackr . run ( "gpt-finetune" , tags = [ "nlp" , "v3" ]):
mltrackr . log ( lr = 2e-5 , epochs = 3 , model = "gpt2" )
for step , batch in enumerate ( dataloader ):
loss = model . train_step ( batch )
mltrackr . log ( loss = loss . item (), step = step )
✅ Beautiful live dashboard
mltrackr ui
Opens at http://localhost:7000 — a fast, dark-mode single-page app with:
Searchable run list with inline sparkline charts in the sidebar
Trend indicators (↑ ↓) showing whether each metric is improving
Side-by-side comparison of any runs you select (best value highlighted)
Auto-generated time-series charts with gradient fills
Metric progress bars showing where the latest value sits in its historical range
Global statistics view — success rate, most-logged metrics, run timeline
Auto-refresh every 5 seconds — open while training, watch it update
✅ Live anomaly detection — catch bad runs early
mltrackr . configure_watch ( nan_check = True , divergence_window = 5 , plateau_window = 15 )
with mltrackr . run ( "training" ):
for epoch in range ( 100 ):
mltrackr . log ( loss = compute_loss ())
# Automatically warns if: loss → NaN, loss diverges for 5 epochs,
# loss plateaus for 15 epochs (and suggests adjusting LR)
Stop wasting GPU hours on runs that are already failing.
mltrackr suggest accuracy
Analyzes your run history and tells you which hyperparameter values are statistically correlated with better results. No black box — plain English insights like:
Best config: lr=0.001 → avg accuracy 0.943 (vs 0.871 for other values, +8.2%)
Next experiment: try batch_size=128 — larger batches correlated with +5.1% accuracy
✅ Auto-generated experiment reports
mltrackr report --output results.md
Generates a thesis-ready markdown report with:
Summary statistics (total runs, completion rate, best configurations)
Chronological experiment timeline
Key findings (computed automatically)
Optional AI narrative: mltrackr report --ai (uses local Ollama, no API keys)
✅ Generate a ready-to-run example
mltrackr init # plain Python example
mltrackr init --framework pytorch # PyTorch training loop
mltrackr init --framework sklearn # scikit-learn grid search
mltrackr init --framework keras # Keras callback
Generates a complete working script you can run immediately.
Framework
How
PyTorch
mltrackr.log(loss=loss.item(), acc=acc) inside the training loop
scikit-learn
mltrackr.log(**params, cv_score=score) in your hyperparam loop
Keras / TF
One-file TrainlogCallback for model.fit()
HuggingFace
Custom TrainerCallback — see examples/huggingface_example.py
XGBoost / LightGBM
Log in the eval callback
JAX / Flax
Log at end of each training step
Plain Python
Anything that produces a number
Full API reference
import mltrackr
# ── Tracking ──────────────────────────────────────────────────────────────────
with mltrackr . run ( "name" , tags = [ "tag1" , "tag2" ]) as run_id :
mltrackr . log ( accuracy = 0.95 , loss = 0.05 ) # log any key-value pairs
mltrackr . note ( "Cosine LR schedule helped a lot" ) # attach plain-text notes
mltrackr . tag ( run_id , "production" ) # add tags after the fact
mltrackr . tag ( "experiment-name" , "best" ) # also works by name
# ── Querying ──────────────────────────────────────────────────────────────────
runs = mltrackr . get_runs () # all runs, newest first
best = mltrackr . get_best_run ( "accuracy" ) # highest final value
best_low = mltrackr . get_best_run ( "loss" , mode = "min" ) # lowest final value
cmp = mltrackr . compare_runs ( 1 , 2 , 3 ) # list of run dicts
# ── Anomaly detection ─────────────────────────────────────────────────────────
mltrackr . configure_watch (
nan_check = True , # warn on NaN/Inf values
divergence_window = 5 , # warn if metric diverges for N steps
plateau_window = 15 , # warn if metric plateaus for N steps
enabled = True ,
)
# Or temporarily with a context manager:
with mltrackr . watch ( divergence_window = 3 ):
# stricter watch for this block
mltrackr . log ( loss = 0.5 )
# ── Export & analysis ─────────────────────────────────────────────────────────
mltrackr . export_csv ( "results.csv" )
mltrackr . export_json ( "results.json" )
mltrackr . generate_report ( "report.md" , use_ollama = False )
suggestions = mltrackr . suggest ( "accuracy" , mode = "max" , top_n = 3 )
mltrackr . clear_all () # deletes everything (irreversible)
# ── Config ────────────────────────────────────────────────────────────────────
mltrackr . configure ( verbose = False ) # suppress auto-summary panels after each run
CLI reference
# Dashboard
mltrackr ui # open at localhost:7000
mltrackr ui --port 8080 --no-browser # custom port, no auto-open
# Inspect runs
mltrackr list # rich table, newest first
mltrackr list --limit 50
mltrackr compare 1 2 3 # side-by-side metric comparison
mltrackr best accuracy # best run for a metric
mltrackr best loss --mode min
# Annotate
mltrackr tag 42 production tuned # add tags to run #42
mltrackr note 42 " Try cosine annealing " # add note to run #42
# Analyse
mltrackr stats # aggregate statistics
mltrackr suggest accuracy # hyperparameter recommendations
mltrackr suggest loss --mode min --top 5
# Generate
mltrackr report # write report.md
mltrackr report -o results.md --ai # with Ollama AI narrative
mltrackr init --framework pytorch # generate example script
# Export / clean
mltrackr export --format csv -o data.csv
mltrackr export --format json -o data.json
mltrackr clear # delete all (asks confirmation)
# Share
mltrackr share accuracy # generate Twitter/X + HN ready post
mltrackr share loss --mode min # for metrics where lower is better
How it works
SQLite — ~/.mltrackr/experiments.db . One file. No server. Inspect it with any SQLite browser. Back it up with cp .
Flask — the dashboard is a local Flask server. Vanilla JS, Chart.js, zero npm, zero build step.
Thread-local state — each training job in its own thread gets an isolated run context. Concurrent experiments just work.
Git-aware — captures the current commit hash via git rev-parse HEAD . Silently skipped outside a git repo.
Watch hooks — anomaly detection runs inside every log() call. Zero external services, works offline.
mltrackr init --framework pytorch -o train.py
python train.py
mltrackr ui
That's the whole flow. Five commands. Zero config.
Live anomaly detection ( configure_watch )
Auto-generated experiment reports ( mltrackr report , Ollama support)
Hyperparameter suggestions ( mltrackr suggest )
Quick-start example generator ( mltrackr init )
Sparkline charts in sidebar with trend indicators
Metric progress bars and trend arrows in detail view
Framework examples: PyTorch, scikit-learn, Keras, HuggingFace
mltrackr.log_artifact("model.pt") — save file paths alongside metrics
Native PyTorch TrainlogCallback (pip-installable plugin)
VS Code extension — inline run summary on hover
mltrackr serve — shareable read-only dashboard URL (ngrok/localtunnel)
Team sync via shared git-tracked SQLite
Slack / Discord webhook on run completion
Have an idea? Open a feature request — or submit a PR.
See CONTRIBUTING.md . TL;DR: pip install -e . , make your change, open a PR.
All contributions welcome — typos, docs, features, bug fixes.
MIT — use it however you want, forever.
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
