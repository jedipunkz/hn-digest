---
source: "https://github.com/yashchimata/statgate"
hn_url: "https://news.ycombinator.com/item?id=48995760"
title: "Show HN: Statgate, statistically calibrated ship/block CI gates for LLM evals"
article_title: "GitHub - yashchimata/statgate: Statistically calibrated ship or block CI gates for LLM evals: paired bootstrap verdicts, power analysis, sequential early stopping, and a GitHub Action · GitHub"
author: "yashchimata"
captured_at: "2026-07-21T18:10:22Z"
capture_tool: "hn-digest"
hn_id: 48995760
score: 1
comments: 0
posted_at: "2026-07-21T17:55:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Statgate, statistically calibrated ship/block CI gates for LLM evals

- HN: [48995760](https://news.ycombinator.com/item?id=48995760)
- Source: [github.com](https://github.com/yashchimata/statgate)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T17:55:23Z

## Translation

タイトル: HN を表示: Statgate、LLM 評価用の統計的に調整されたシップ/ブロック CI ゲート
記事のタイトル: GitHub - yashchimata/statgate: LLM 評価用の統計的に調整されたシップまたはブロック CI ゲート: ペアのブートストラップ判定、電力分析、逐次早期停止、および GitHub アクション · GitHub
説明: LLM 評価用の統計的に調整されたシップまたはブロック CI ゲート: ペアのブートストラップ判定、電力分析、順次早期停止、および GitHub アクション - yashchimata/statgate

記事本文:
GitHub - yashchimata/statgate: LLM 評価用の統計的に調整されたシップまたはブロック CI ゲート: ペアのブートストラップ判定、電力分析、順次早期停止、および GitHub アクション · GitHub
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
ヤシチマ

た
/
スタットゲート
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスで表示する メイン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .github .github アセット アセット ドキュメント ドキュメント サンプル サンプル スクリプト スクリプト src/ statgate src/ statgate テスト テスト .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスREADME.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md action.yml action.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM 評価用に統計的に調整されたシップまたはブロック CI ゲート。
あなたの評価ゲートはあなたに嘘をついています。 50 件のケーススイートでスコアが 2 ポイント低下すると、
ほとんどの場合、ノイズがサンプリングされますが、しきい値ゲートはそれを回帰として扱います。
したがって、チームはコイントスでマージをブロックするか、赤い X を無視することを学ぶかのどちらかになります。
それはもっと悪いです。 LLM 出力は非決定的であり、評価スイートは小さく、
生のパスレート比較では、信号とノイズを区別できません。
statgate は、評価ランナーとマージの間の統計レイヤーです。
ボタン:
操作上の失敗 (不正なファイル、不正なフラグ) はコード 3 で終了するため、
パイプラインは常に「ゲートが決定された」か「ゲートが壊れたか」を区別できます。
statgate は LLM 呼び出しを行わず、ネットワーク要求も送信せず、のみに依存します。
ヌルヌル、クリック感、リッチ、そしてピダンティック。それは、次のことができるすべての評価ランナーで動作します。
結果ファイルを生成します。
pip インストール statgate
またはリポジトリから直接:
pip install git+https://github.com/yashchimata/statgate
62回目のツアー
リポジトリにはサンプル データが付属しているため、以下のすべてのコマンドは新しいデータに対して機能します。
クローン。
1. 変更をゲートします。ペア

ケースIDで実行し、差分をブートストラップします。
決める：
statgate 比較の例/baseline.jsonl 例/candidate.jsonl
それが上のSHIPのスクリーンショットです。本当に回帰した実行に対して、
同じコマンドブロックであり、エラーバーはその理由を示しています: 全体の信頼性
間隔はマージンの左側にあります。
2. スイートが実際に何を検出できるかを学びます。ほとんどの評価スイートでは、
所有者が信じているマージンをサポートします。このコマンドはそれを置き換えます
算術による引数:
statgate power --baseline の例/baseline.jsonl --candidate の例/candidate.jsonl
3. 必要のない評価ケースに対する支払いを停止します。シーケンシャルモードが適用されます
常に有効な停止境界があるため、明確な判定は早期に終了します。ここで、
回帰は 80 件中 20 件で検出され、評価費用の 75% が節約されます。
statgate シーケンシャルの例/baseline.jsonl の例/regression.jsonl --batch-size 20
シーケンシャル モードでは、eval コマンドをバッチごとにライブで実行することもできます。
--run "python run_eval.py --offset {start} --limit {count}" 。
4. statgate validate results.jsonl を使用して、結果ファイルの健全性をチェックします。
ジョブ:
評価ゲート:
実行: ubuntu-最新
権限:
プルリクエスト : 書き込み
手順:
- 使用:actions/checkout@v7
- 名前 : 評価の実行
run : ./run_evals.sh # Baseline.jsonl と Candidate.jsonl を生成します
- 使用: yashchimata/statgate@v0.1.1
付き:
ベースライン:baseline.jsonl
候補：candidate.jsonl
このアクションは、判定とエラーバーを含む付箋コメントを投稿し、更新します
プッシュするたびに配置され、ジョブの概要にレポートが書き込まれ、失敗します。
BLOCK (または設定した場合は INCONCLUSIVE) でのみチェックします。
不確定な失敗時: true )。
ライブでご覧ください: このリポジトリは、プルごとに独自のサンプル評価をゲートします。
リクエスト。オープンデモのプルリクエスト
チェックに失敗した実際の SHIP コメントと実際の BLOCK を表示します。
すべての入力とベースライン st

レート (メイン、キャッシュされたアーティファクト、
コミットされたベースライン）については、 docs/github-action.md で説明されています。
statgate は、eval フレームワークではなく、結果ファイルを取り込みます。フォーマットが検出されました
ファイルから、または --adapter で強制的に。
JSON 行 (1 行につき 1 レコード):
{ "case_id" : " q-001 " 、 "スコア" : 0.83 、 "合格" : true }
{ "case_id" : " q-002 " 、 "スコア" : 0.41 、 "合格" : false 、 "run_index" : 0 }
case_id はテスト ケースを識別し、ベースラインとテスト ケースの間で一致する必要があります。
候補者。これにより可能になるペア分析には、通常 3 ～ 10 回の時間が必要です
同じ確実性については、対応のない比較よりもケースが少なくなります。どちらか
スコアまたは合格が必要です。同じケースの繰り返し実行
( run_index ) はケースごとに平均化されるため、誤って認識されることはありません。
独立したサンプル。不明なフィールドはメタデータとして保持されます。の JSON 配列
同じオブジェクトと同じ列を持つ CSV は両方とも機能します。
プロンプトfoo : によって書き込まれたファイルを statgate に指定します。
プロンプトfoo eval -o results.json 。ケース ID はテストとプロンプトから派生します
したがって、同じスイートの 2 つのエクスポートはきれいにペアになります。
ランナーのフォーマットが見つかりませんか?アダプターは約 50 行あります。
アダプターリクエストを開く
サンプルファイル付き。
CLI が行うことはすべて、型付きライブラリとして利用できます。
pathlibインポートパスから
statgate からインポート GateSettings 、 Verdict 、evaluate_gate
スタットゲートから。アダプターのインポートload_records
Baseline = load_records (パス ( "baseline.jsonl" ))
候補 =load_records (パス ( "candidate.jsonl" ))
report = Evaluate_gate (ベースライン、候補、GateSettings (マージン = 0.02、シード = 42))
print (レポートの評決、レポートの間隔が短い、レポートの間隔が長い)
ご報告があれば評決は 評決 です。ブロック :
SystemExit を発生させる ( 1 )
レポートには、間隔、p 値、側ごとの要約、およびスイートが含まれます。
評決が不確定の場合に必要なサイズ

、カスタムを構築できます
pytest フィクスチャまたは他の場所のポリシー。
statgate は、作業ディレクトリまたはから statgate.toml を読み取ります。
--config パス。最も一般的な値は CLI で上書きすることもできます
--margin 、 --alpha 、 --metric 、 --seed などのフラグ。
【門】
metric = " スコア " # または "pass_rate"
alpha = 0.05 # 有意水準;信頼度は 1 - アルファです
margin = 0.02 # 許容できる回帰サイズ
power = 0.8 # スイート サイズの推奨事項のターゲット電力
リサンプル = 10000 # ブートストラップ リサンプル
順列 = 10000 # 順列テストの反復
シード = 42 # 完全に再現可能なレポート用に設定
【順次】
バッチサイズ = 25
最大ケース数 = 400
マージンはポリシーの核心です。 margin = 0.02 は「ブロックのみ」を意味します
候補者が 2 ポイント以上劣っていると確信できる場合。」余裕
ゼロの場合は厳格な改善の証拠が必要ですが、小規模なスイートではそれができないことはほとんどありません。
提供します。完全なリファレンスは docs/configuration.md にあります。
statgate は、eval が生成するスコアをすべて信頼します。それらのスコアが来たら
LLM 裁判官、ジャッジゲートより
あなたが審査する前に、人間のラベルに照らして審査員自身を検証します。
これらは一緒に、eval パイプラインの信頼の両方の部分をカバーします。
短いバージョン: ケースごとのペアの違い、BCa ブートストラップの信頼度
平均値の間隔、クロスチェックとしての符号反転順列テスト、
非劣性決定ルール、正規近似検出力分析、および
常に有効な連続停止のための混合 SPRT。ケースはリサンプリングです
ユニットなので、同じケースを繰り返し実行しても、独立したものと間違われることはありません
証拠。
それぞれの選択の背後にある理由と既知の理由を含むロングバージョン
制限事項は docs/methodology.md にあります。テスト
スイートには、間隔カバレッジと false を検証するキャリブレーション チェックが含まれています
k を使用した合成データに対する陽性率

現在のグラウンドトゥルースを含む
順次オプション停止中。
git clone https://github.com/yashchimata/statgate
cd ステータスゲート
python -m venv .venv && 。 .venv/bin/activate # Windows の場合は .venv\Scripts\activate
pip install -e " .[dev] "
pytest
ruff チェック SRC テスト
mypyソース
貢献、特に新しいアダプターを歓迎します。参照
CONTRIBUTING.md と
良い最初の問題。
LLM 評価用に統計的に調整されたシップまたはブロック CI ゲート: ペアのブートストラップ判定、電力分析、順次早期停止、および GitHub アクション
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
2
スタットゲート v0.1.1
最新の
2026 年 7 月 15 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Statistically calibrated ship or block CI gates for LLM evals: paired bootstrap verdicts, power analysis, sequential early stopping, and a GitHub Action - yashchimata/statgate

GitHub - yashchimata/statgate: Statistically calibrated ship or block CI gates for LLM evals: paired bootstrap verdicts, power analysis, sequential early stopping, and a GitHub Action · GitHub
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
yashchimata
/
statgate
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .github .github assets assets docs docs examples examples scripts scripts src/ statgate src/ statgate tests tests .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
Statistically calibrated ship or block CI gates for LLM evals.
Your eval gate is lying to you. A 2 point score drop on a 50 case suite is
almost always sampling noise, but a threshold gate treats it as a regression.
So teams either block merges on coin flips or learn to ignore the red X,
which is worse. LLM outputs are nondeterministic, eval suites are small, and
a raw pass rate comparison cannot tell signal from noise.
statgate is the statistics layer between your eval runner and your merge
button:
Operational failures (bad files, bad flags) exit with code 3, so your
pipeline can always distinguish "the gate decided" from "the gate broke".
statgate makes no LLM calls, sends no network requests, and depends only on
numpy, click, rich, and pydantic. It works with any eval runner that can
produce a results file.
pip install statgate
Or straight from the repository:
pip install git+https://github.com/yashchimata/statgate
Sixty second tour
The repository ships example data, so every command below works on a fresh
clone.
1. Gate a change. Pair the runs by case id, bootstrap the difference,
decide:
statgate compare examples/baseline.jsonl examples/candidate.jsonl
That is the SHIP screenshot above. Against a genuinely regressed run the
same command blocks, and the error bar shows why: the entire confidence
interval sits left of the margin.
2. Learn what your suite can actually detect. Most eval suites cannot
support the margins their owners believe in. This command replaces that
argument with arithmetic:
statgate power --baseline examples/baseline.jsonl --candidate examples/candidate.jsonl
3. Stop paying for eval cases you do not need. Sequential mode applies
always-valid stopping boundaries, so clear verdicts end early. Here the
regression is caught after 20 of 80 cases, saving 75% of the eval spend:
statgate sequential examples/baseline.jsonl examples/regression.jsonl --batch-size 20
Sequential mode can also drive your eval command live, batch by batch, with
--run "python run_eval.py --offset {start} --limit {count}" .
4. Sanity check any results file with statgate validate results.jsonl .
jobs :
eval-gate :
runs-on : ubuntu-latest
permissions :
pull-requests : write
steps :
- uses : actions/checkout@v7
- name : Run evals
run : ./run_evals.sh # produces baseline.jsonl and candidate.jsonl
- uses : yashchimata/statgate@v0.1.1
with :
baseline : baseline.jsonl
candidate : candidate.jsonl
The action posts a sticky comment with the verdict and error bars, updates
it in place on every push, writes the report to the job summary, and fails
the check only on BLOCK (or on INCONCLUSIVE if you set
fail-on-inconclusive: true ).
See it live: this repository gates its own example evals on every pull
request. The open demo pull requests
show a real SHIP comment and a real BLOCK with a failed check.
All inputs and baseline strategies (re-run on main, cached artifact,
committed baseline) are covered in docs/github-action.md .
statgate ingests results files, not eval frameworks. The format is detected
from the file, or forced with --adapter .
JSON Lines (one record per line):
{ "case_id" : " q-001 " , "score" : 0.83 , "passed" : true }
{ "case_id" : " q-002 " , "score" : 0.41 , "passed" : false , "run_index" : 0 }
case_id identifies the test case and must match between baseline and
candidate; the paired analysis it enables typically needs 3 to 10 times
fewer cases than an unpaired comparison for the same certainty. Either
score or passed is required. Repeated runs of the same case
( run_index ) are averaged per case, so they are never mistaken for
independent samples. Unknown fields are kept as metadata. A JSON array of
the same objects and a CSV with the same columns both work.
promptfoo : point statgate at the file written by
promptfoo eval -o results.json . Case ids are derived from test and prompt
indices, so two exports of the same suite pair cleanly.
Missing your runner's format? Adapters are about 50 lines;
open an adapter request
with a sample file.
Everything the CLI does is available as a typed library:
from pathlib import Path
from statgate import GateSettings , Verdict , evaluate_gate
from statgate . adapters import load_records
baseline = load_records ( Path ( "baseline.jsonl" ))
candidate = load_records ( Path ( "candidate.jsonl" ))
report = evaluate_gate ( baseline , candidate , GateSettings ( margin = 0.02 , seed = 42 ))
print ( report . verdict , report . interval . low , report . interval . high )
if report . verdict is Verdict . BLOCK :
raise SystemExit ( 1 )
report carries the interval, p-value, per-side summaries, and the suite
size needed when the verdict is INCONCLUSIVE, so you can build custom
policies in a pytest fixture or anywhere else.
statgate reads statgate.toml from the working directory, or from
--config path . The most common values can also be overridden with CLI
flags such as --margin , --alpha , --metric , and --seed .
[ gate ]
metric = " score " # or "pass_rate"
alpha = 0.05 # significance level; confidence is 1 - alpha
margin = 0.02 # regression size you are willing to tolerate
power = 0.8 # target power for suite size recommendations
resamples = 10000 # bootstrap resamples
permutations = 10000 # permutation test iterations
seed = 42 # set for fully reproducible reports
[ sequential ]
batch_size = 25
max_cases = 400
The margin is the heart of the policy. margin = 0.02 means "block only
when we are confident the candidate is more than 2 points worse". A margin
of zero demands proof of strict improvement, which small suites can rarely
provide. Full reference in docs/configuration.md .
statgate trusts whatever scores your evals produce. When those scores come
from an LLM judge, judgegate
verifies the judge itself against human labels before you gate on it.
Together they cover both halves of trusting an eval pipeline.
Short version: per-case paired differences, a BCa bootstrap confidence
interval on their mean, a sign-flip permutation test as a cross-check, a
non-inferiority decision rule, normal-approximation power analysis, and a
mixture SPRT for always-valid sequential stopping. Cases are the resampling
unit, so repeated runs of the same case are never mistaken for independent
evidence.
The long version, with the reasoning behind each choice and the known
limitations, is in docs/methodology.md . The test
suite includes calibration checks that verify interval coverage and false
positive rates against synthetic data with known ground truth, including
under sequential optional stopping.
git clone https://github.com/yashchimata/statgate
cd statgate
python -m venv .venv && . .venv/bin/activate # .venv\Scripts\activate on Windows
pip install -e " .[dev] "
pytest
ruff check src tests
mypy src
Contributions are welcome, especially new adapters. See
CONTRIBUTING.md and the
good first issues .
Statistically calibrated ship or block CI gates for LLM evals: paired bootstrap verdicts, power analysis, sequential early stopping, and a GitHub Action
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
2
statgate v0.1.1
Latest
Jul 15, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
