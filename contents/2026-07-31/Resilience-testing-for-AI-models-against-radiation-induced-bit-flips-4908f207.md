---
source: "https://github.com/ananfauh7/MRVPlatform"
hn_url: "https://news.ycombinator.com/item?id=49125712"
title: "Resilience testing for AI models against radiation-induced bit flips"
article_title: "GitHub - ananfauh7/MRVPlatform · GitHub"
author: "ananfauh7"
captured_at: "2026-07-31T17:20:05Z"
capture_tool: "hn-digest"
hn_id: 49125712
score: 1
comments: 0
posted_at: "2026-07-31T16:56:28Z"
tags:
  - hacker-news
  - translated
---

# Resilience testing for AI models against radiation-induced bit flips

- HN: [49125712](https://news.ycombinator.com/item?id=49125712)
- Source: [github.com](https://github.com/ananfauh7/MRVPlatform)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T16:56:28Z

## Translation

タイトル: 放射線誘発ビット反転に対する AI モデルの回復力テスト
記事タイトル: GitHub - ananfauh7/MRVPlatform · GitHub
説明: GitHub でアカウントを作成して、ananfauh7/MRVPlatform の開発に貢献します。

記事本文:
GitHub - ananfauh7/MRVPlatform · GitHub
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
アナンファウ7
/
MRVプラットフォーム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 Co

de 「その他のアクション」メニューを開く フォルダーとファイル
4 コミット 4 コミット .ananth-skills .ananth-skills .cursor .cursor ベンチマーク ベンチマーク docs docs resilience_engine resilience_engine テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml required-lock.txt requirements-lock.txt resilience_test.py resilience_test.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オンボード AI モデル復元力検証プラットフォーム
ONNX モデルの予測が、衛星の搭載コンピューターが軌道上で実際に経験する種類のメモリ破損 (放射線によるシングル イベント アップセット (SEU)、別名ビット 反転) に耐えられるかどうかをテストするコマンド ライン ツールです。モデルが実行されるかどうかを確認するだけではありません。
モデルと検証セットが与えられると、統計的に現実的なビット フリップをモデルの重みに挿入し、推論を何度も再実行して、破損を生き残るためにどの重みテンソルが最も重要かを示す 0 ～ 100 の回復力スコアとレイヤーごとの感度の内訳を報告します。
これは、完全な SaaS 製品に投資する前に 1 つのことをテストするための 3 日間の CLI 検証スプリントとして構築されました。それは、このフォールト挿入方法論が実際に機能し、一般化できるのか、また、その結果は懐疑的な航空宇宙エンジニアに示されるほど信頼できるものなのかということです。この README は、そのスプリントで何が証明され、何が証明されなかったかについて意図的に正直に記載しています。このツールが生成するレポートから結論を導き出す前に、「現在の制限事項」を参照してください。
git clone https://github.com/ananfauh7/MRVPlatform.git
cd MRVプラットフォーム
python -m venv .venv && ソース .venv/bin/activate
pip install -e 。
python resilience_test.py \
--モデルベンチマーク/mnist/model.onnx \
--データベンチマーク/mnist/dataset.npz \
--preset leo-typical \
--ミッションデイズ 365 \
--トライアル20
これにより、実際の事前トレーニング済み CNN と実際の MNIST vali に対してツールが実行されます。

リポジトリに同梱されている日付サブセット — 試すために外部からダウンロードする必要はありません。すべてが決定的にシードされるため、上記の正確なコマンドを実行すると、次の正確な出力が再現されます。
プリセット「leo-typical」を使用する場合：fault_rate=2e-07 エラー/ビット日
Benchmarks/mnist/model.onnx... からモデルをロードしています...
Benchmarks/mnist/dataset.npz から検証データセットをロードしています...
50 個のサンプルでクリーンなベースラインを実行しています...
20 個のモンテカルロ トライアルを、fault_rate=2e-07、mission_days=365.0 で実行中...
レイヤーごとの感度解析を実行しています (レイヤーごとに 8 トライアル)...
======================================================
レジリエンスレポート
======================================================
クリーン精度: 1.0000
故障時の精度 (平均): 0.8190
故障時の精度 (標準): 0.3622
障害のある精度 (範囲): [0.0600、1.0000]
反発力スコア: 81.9 / 100
試行回数: 20
最上位の敏感な層:
#1 パラメータ 87 の平均精度低下: 0.1125
#2 パラメータ 193 の平均精度低下: 0.0000
#3 パラメータ 5 の平均精度低下: 0.0000
#4 パラメータ6の平均精度低下: 0.0000
#5 パラメータ88の平均精度低下: 0.0000
完全な結果はresults.jsonに書き込まれます。
チャートの書き込み先:Accuracy_vs_fault_rate.png
======================================================
何をするのか
ONNX モデルと検証データセットを読み込み、クリーンラン精度のベースラインを確立します。
フォールト率 (ビット日あたりのエラー数) とミッション露出期間 (日) を指定して、各重みテンソルで予想されるビット フリップの数を計算します: Expected_flips = fall_rate * tensor_size_bits * Mission_days 。
これらのビット反転を、モデルの重みの IEEE-754 float32 ビット表現に直接挿入します。符号、指数、仮数のビットは同様に、均一にランダムに選択されます (指数反転が不釣り合いなダメージを与えるかどうかは、ツールの発見によるものです)

サーフェスがベイクインされるという仮定ではありません)。
破損した重みを含む検証セットに対して推論を再実行し、結果の精度を測定します。
これをモンテカルロ シミュレーション (デフォルト: 20 試行) として繰り返し、破損した 1 回の実行だけでなく、結果の分布を特徴付けます。単一の試行の精度は、どのビットがヒットするかによって大きく異なります。
回復力スコアを計算します: 100 * (平均障害精度 / クリーン精度)、 [0, 100] にクリップされます。
レイヤーごとの感度解析を実行します。重みテンソルごとにトライアル ループを 1 回繰り返し、毎回そのテンソルのみを破損し、ヒット時の精度の低下度によってテンソルをランク付けします。これは、このツールの最も有用な非自明な出力です。集計スコアだけでなく、どのレイヤーが冗長性/強化から最も恩恵を受けるかを示します。
JSON レポートとグラフ、およびコンソールの概要を書き込みます。
ビット 反転は、各重みテンソルの生の IEEE-754 float32 ビット パターンのレベルでシミュレートされます。
ビット 31 : 符号
ビット 30 ～ 23 : 指数 (8 ビット、127 でバイアス)
ビット 22-0 : 仮数部 (23 ビット)
N ビットのテンソルの場合、D 日間のミッションにわたる故障率 r (エラー数/ビット日) では、予想されるフリップ数は r * N * D となり、最も近い整数に四捨五入されます。次に、正確なビット位置が均一にランダムに (置換なしで) 選択され、テンソルの uint32 ビット ビューで XOR を介して反転されます。 --seed を指定すると、すべてが完全に決定的になります。同じシード、モデル、データセット、パラメーターは常に同じ結果を再現します。
1 回の破損した実行だけでは、ほとんど何もわかりません (不運な仮数ビットの反転は目に見えません。臨界重みでの不運な指数ビットの反転は壊滅的な結果をもたらす可能性があります)。このツールは、--trials 独立した破損と評価のトライアル (それぞれ独自の派生サブシードを持つ) を実行し、平均値、標準値を報告します。

クリーンなベースラインに対する、結果として生じる精度と損失のデルタの偏差、最小値、最大値。
ヘッドラインのモデル全体のスコアとは別に、このツールは重みテンソルごとにトライアル ループを 1 回繰り返し (毎回そのテンソルのみを破損し、他のすべてのテンソルはクリーンなままにして)、平均精度の低下によってテンソルをランク付けします。これは、「1 つのレイヤーだけを強化/複製できるとしたら、どれが最も重要ですか?」という質問に対する答えであり、集計スコアだけよりも実用的です。
スコア = 100 * (mean_faulted_accuracy / clean_accuracy) # [0, 100] にクリップされる
このスプリントでは意図的にシンプルにしました。現時点では、損失の劣化、予測の信頼性の調整、またはニアミスが完全な故障に対してどの程度「近い」かという重み付けは考慮されていません。「現在の制限事項」を参照してください。
--fault-rate 値を推測することを要求するのではなく、--preset は、公開されているシングル イベント アップセット (SEU) 率の調査に基づいた名前を受け入れます。
すべてのプリセットの正確な引用と導出の推論は、その値とともに resilience_engine/presets.py に文書化されています。プリセットを実際のミッション決定のための検証済みの数値として扱う前に読んでください。 「推定」とマークされたプリセットは、桁違いの数値であり、デバイス固有の測定値ではありません。
resilience-test --model MODEL.onnx --data DATA (--fault-rate RATE | --preset NAME)
[--ミッション-日数 DAYS] [--トライアル N] [--layer-トライアル N]
[--seed SEED] [--output-dir DIR] [--debug]
旗
必須
デフォルト
説明
--モデル
はい
—
ONNX モデル ファイルへのパス
--データ
はい
—
.npz ファイル (入力/ラベル配列を含む) または .npy ファイルのクラスごとのディレクトリへのパス
--故障率
この 2 つのうち 1 つ
—
エラー/ビット日の故障率 (例: 1e-6 )
--プリセット
必須
—
leo-typical 、 saa-transit 、または Solar-storm-worst-case
--ミッション日数
いいえ
365
ミッション暴露期間（日数）
--トライアル
いいえ
20
モンテカルロ裁判

r 見出しの回復力スコア
--layer-トライアル
いいえ
8
感度内訳のレイヤーごとのトライアル数
--シード
いいえ
42
基本ランダムシード (再現性のため)
--出力ディレクトリ
いいえ
。
results.jsonとaccuracy_vs_fault_rate.pngを書く場所
--デバッグ
いいえ
オフ
クリーンなエラー メッセージの代わりに完全な Python トレースバックを表示する
--fault-rate または --preset のいずれかが必要です (相互に排他的)。
名前付きプリセットの使用 (推奨開始点):
python resilience_test.py --model my_model.onnx --data ./val_data --preset leo-typical
明示的な故障率と短いミッションを使用すると、次のようになります。
python resilience_test.py --model my_model.onnx --data ./val_data \
--故障率 1e-6 --ミッション日数 90 --トライアル数 50
クラスごとのディレクトリ データセットを使用します (整数クラス ラベルごとに 1 つのサブディレクトリ、それぞれに .npy ファイルが含まれます)。
val_data/
0/example_001.npy
0/example_002.npy
1/example_001.npy
...
python resilience_test.py --model my_model.onnx --data ./val_data --preset saa-transit --output-dir ./results
出力
コンソールの概要: クリーン/フォールトの精度、回復力スコア、および最も機密性の高い上位 5 つのレイヤー。
results.json : 完全なモンテカルロ結果 (すべてのトライアルの結果、分布統計、復元力スコア、監査可能性のためのモデル/データセット コンテンツ ハッシュ) と完全なレイヤーごとの感度ランキング。
activity_vs_fault_rate.png : テストされた故障率におけるクリーンな精度と故障した精度 (平均 ± 標準誤差バー) のグラフ。
私たちにとって、これらの事項については、見た目が印象的であるよりも明確であることが重要です。私たちが最初にそれらのギャップを明らかにしなければ、航空宇宙エンジニアはこれらのギャップをすぐに見つけるでしょう。
重量の破損のみ。このツールは、アクティベーション、勾配、制御フロー、または周囲のフライト ソフトウェア/OS ではなく、モデルの重みを破壊します。実際のオンボード SEU は、これらのいずれにもヒットする可能性があります。重みの破損は、実際の断層面の意図されたスライスの 1 つであり、

全部ではありません。
熱や電力スロットリングのモデリングはありません。 MVP 仕様のより広範なビジョン (放射 + パワー スロットリング + メモリ破損) は、ここでは部分的にのみ実装されています。このスプリントでは、特に放射/ビット フリップ部分を検証しました。
精度のみの復元力スコア。現在、スコアには、損失の劣化や予測の信頼性の調整は考慮されておらず、トップ 1 の精度への影響を超えてニアミスと壊滅的な故障が区別されていません。
実行ごとに単一の動作点。各実行は、1 つの障害率とミッション日数の組み合わせを反映し、暴露レベル全体の完全な感度曲線を反映するものではありません (ただし、複数の実行をスクリプト化することを妨げるものはありません)。
2 つのモデル アーキテクチャに対して検証されていますが、普遍的に一般化できることは証明されていません。コア パイプラインは構築され、小規模な合成線形モデルに対して単体テストが行​​われ、その後、実際の事前トレーニング済み CNN (benchmarks/mnist/ を参照) に対してエンドツーエンドで検証され、合成のケースを超えてアプローチが一般化しているかどうかが確認されました。これは意味のある証拠であって証拠ではありません。たとえば、トランスフォーマー アーキテクチャ、カスタム/エキゾチックな演算を備えたモデル、または ONNX opset 8 時代のグラフを超えるものに対してはテストされていません。
3 つの故障率プリセットのうち 2 つは推定値であり、直接の測定値ではありません。 leo-typical が派生します

[切り捨てられた]

## Original Extract

Contribute to ananfauh7/MRVPlatform development by creating an account on GitHub.

GitHub - ananfauh7/MRVPlatform · GitHub
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
ananfauh7
/
MRVPlatform
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .ananth-skills .ananth-skills .cursor .cursor benchmarks benchmarks docs docs resilience_engine resilience_engine tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml requirements-lock.txt requirements-lock.txt resilience_test.py resilience_test.py View all files Repository files navigation
Onboard AI Model Resilience Validation Platform
A command-line tool that tests whether an ONNX model's predictions survive the kind of memory corruption a satellite's onboard computer actually experiences in orbit — radiation-induced single event upsets (SEUs), a.k.a. bit flips — rather than just checking that the model runs.
Given a model and a validation set, it injects statistically-realistic bit flips into the model's weights, re-runs inference many times, and reports a 0-100 resilience score plus a per-layer sensitivity breakdown showing which weight tensors matter most for surviving corruption.
This was built as a 3-day CLI validation sprint to test one thing before investing in a full SaaS product: does this fault-injection methodology actually work and generalize, and is the finding credible enough to show a skeptical aerospace engineer? This README is deliberately honest about what that sprint did and didn't prove — see Current Limitations before drawing conclusions from a report this tool produces.
git clone https://github.com/ananfauh7/MRVPlatform.git
cd MRVPlatform
python -m venv .venv && source .venv/bin/activate
pip install -e .
python resilience_test.py \
--model benchmarks/mnist/model.onnx \
--data benchmarks/mnist/dataset.npz \
--preset leo-typical \
--mission-days 365 \
--trials 20
This runs the tool against a real pretrained CNN and a real MNIST validation subset that ship with the repo — no external downloads needed to try it out. Since everything is seeded deterministically, running the exact command above reproduces this exact output:
Using preset 'leo-typical': fault_rate=2e-07 errors/bit-day
Loading model from benchmarks/mnist/model.onnx...
Loading validation dataset from benchmarks/mnist/dataset.npz...
Running clean baseline on 50 example(s)...
Running 20 Monte Carlo trial(s) at fault_rate=2e-07, mission_days=365.0...
Running per-layer sensitivity analysis (8 trial(s) per layer)...
============================================================
RESILIENCE REPORT
============================================================
Clean accuracy: 1.0000
Faulted accuracy (mean): 0.8190
Faulted accuracy (std): 0.3622
Faulted accuracy (range): [0.0600, 1.0000]
Resilience score: 81.9 / 100
Trials run: 20
Top sensitive layers:
#1 Parameter87 mean accuracy drop: 0.1125
#2 Parameter193 mean accuracy drop: 0.0000
#3 Parameter5 mean accuracy drop: 0.0000
#4 Parameter6 mean accuracy drop: 0.0000
#5 Parameter88 mean accuracy drop: 0.0000
Full results written to: results.json
Chart written to: accuracy_vs_fault_rate.png
============================================================
What it does
Loads your ONNX model and validation dataset, and establishes a clean-run accuracy baseline.
Computes how many bit flips to expect in each weight tensor, given a fault rate (errors per bit-day) and a mission exposure duration (days): expected_flips = fault_rate * tensor_size_bits * mission_days .
Injects those bit flips directly into the IEEE-754 float32 bit representation of the model's weights — sign, exponent, and mantissa bits alike, chosen uniformly at random (whether exponent flips are disproportionately damaging is a finding the tool surfaces, not an assumption it bakes in).
Re-runs inference against the validation set with the corrupted weights and measures the resulting accuracy.
Repeats this as a Monte Carlo simulation (default: 20 trials) to characterize the distribution of outcomes, not just one corrupted run — a single trial's accuracy can vary enormously depending on which bits happen to get hit.
Computes a resilience score : 100 * (mean faulted accuracy / clean accuracy) , clipped to [0, 100] .
Runs a per-layer sensitivity analysis : repeats the trial loop once per weight tensor, corrupting only that tensor each time, and ranks tensors by how much they degrade accuracy when hit — this is the tool's most useful non-obvious output: it tells you which layers would benefit most from redundancy/hardening, not just an aggregate score.
Writes a JSON report and a chart , plus a console summary.
Bit flips are simulated at the level of each weight tensor's raw IEEE-754 float32 bit pattern:
bit 31 : sign
bits 30-23 : exponent (8 bits, biased by 127)
bits 22-0 : mantissa (23 bits)
For a tensor of N bits, at fault rate r (errors/bit-day) over a mission of D days, the expected number of flips is r * N * D , rounded to the nearest integer. Exact bit positions are then chosen uniformly at random (without replacement) and flipped via XOR on the tensor's uint32 bit view. Everything is fully deterministic given a --seed — the same seed, model, dataset, and parameters always reproduce the same result.
A single corrupted run tells you almost nothing on its own (an unlucky mantissa-bit flip is invisible; an unlucky exponent-bit flip on a critical weight can be catastrophic). The tool runs --trials independent corruption-and-evaluate trials (each with its own derived sub-seed) and reports the mean, standard deviation, min, and max of the resulting accuracy and loss deltas relative to the clean baseline.
Separately from the headline whole-model score, the tool repeats the trial loop once per weight tensor — corrupting only that tensor each time, leaving every other tensor clean — and ranks tensors by mean accuracy drop. This answers "if I could only harden/replicate one layer, which one matters most?", which is more actionable than an aggregate score alone.
score = 100 * (mean_faulted_accuracy / clean_accuracy) # clipped to [0, 100]
Deliberately simple for this sprint. It does not currently factor in loss degradation, prediction-confidence calibration, or weight how "close" a near-miss is versus total failure — see Current Limitations .
Rather than requiring you to guess a --fault-rate value, --preset accepts a name grounded in published Single Event Upset (SEU) rate research:
Every preset's exact citation and derivation reasoning is documented alongside its value in resilience_engine/presets.py — read it before treating a preset as a validated figure for a real mission decision. Presets marked "estimate" are order-of-magnitude figures, not device-specific measurements.
resilience-test --model MODEL.onnx --data DATA (--fault-rate RATE | --preset NAME)
[--mission-days DAYS] [--trials N] [--layer-trials N]
[--seed SEED] [--output-dir DIR] [--debug]
Flag
Required
Default
Description
--model
Yes
—
Path to the ONNX model file
--data
Yes
—
Path to a .npz file (with inputs / labels arrays) or a per-class directory of .npy files
--fault-rate
One of these two
—
Fault rate in errors/bit-day (e.g. 1e-6 )
--preset
required
—
leo-typical , saa-transit , or solar-storm-worst-case
--mission-days
No
365
Mission exposure duration in days
--trials
No
20
Monte Carlo trials for the headline resilience score
--layer-trials
No
8
Trials per layer for the sensitivity breakdown
--seed
No
42
Base random seed (for reproducibility)
--output-dir
No
.
Where to write results.json and accuracy_vs_fault_rate.png
--debug
No
off
Show full Python tracebacks instead of a clean error message
Either --fault-rate or --preset is required (mutually exclusive).
Using a named preset (recommended starting point):
python resilience_test.py --model my_model.onnx --data ./val_data --preset leo-typical
Using an explicit fault rate and a shorter mission:
python resilience_test.py --model my_model.onnx --data ./val_data \
--fault-rate 1e-6 --mission-days 90 --trials 50
Using a per-class directory dataset (one subdirectory per integer class label, each containing .npy files):
val_data/
0/example_001.npy
0/example_002.npy
1/example_001.npy
...
python resilience_test.py --model my_model.onnx --data ./val_data --preset saa-transit --output-dir ./results
Output
Console summary : clean/faulted accuracy, resilience score, and top-5 most sensitive layers.
results.json : the full Monte Carlo result (every trial's outcome, distribution statistics, resilience score, model/dataset content hashes for auditability) plus the complete per-layer sensitivity ranking.
accuracy_vs_fault_rate.png : a chart of clean vs. faulted accuracy (mean ± std error bars) at the tested fault rate.
Being explicit about these matters more to us than looking impressive — an aerospace engineer will find these gaps immediately if we don't disclose them first.
Weight corruption only. This tool corrupts model weights , not activations, gradients, control flow, or the surrounding flight software/OS. Real onboard SEUs can hit any of those; weight corruption is one well-motivated slice of the real fault surface, not the whole thing.
No thermal or power-throttling modeling. The MVP spec's broader vision (radiation + power throttling + memory corruption) is only partially implemented here — this sprint validated the radiation/bit-flip piece specifically.
Accuracy-only resilience score. The score doesn't currently factor in loss degradation, prediction-confidence calibration, or distinguish a near-miss from a catastrophic failure beyond their effect on top-1 accuracy.
Single operating point per run. Each run reflects one fault-rate/mission-days combination, not a full sensitivity curve across exposure levels (though nothing stops you from scripting multiple runs).
Validated against two model architectures, not proven to generalize universally. The core pipeline was built and unit-tested against a small synthetic linear model, then validated end-to-end against a real pretrained CNN (see benchmarks/mnist/ ) to check the approach generalizes beyond the synthetic case. That's meaningful evidence, not proof — it hasn't been tested against, for example, transformer architectures, models with custom/exotic ops, or anything beyond ONNX opset 8-era graphs.
Two of three fault-rate presets are estimates, not direct measurements. leo-typical is derived

[truncated]
