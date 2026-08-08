---
source: "https://github.com/arnoldpredator/pyreplay"
hn_url: "https://news.ycombinator.com/item?id=49223208"
title: "A tool to quickly step through and audit LLM-generated Python codebases"
article_title: "GitHub - arnoldpredator/pyreplay: An LLM can generate a whole codebase in seconds; understanding it is the slow part. pyreplay makes that first pass fast — top-down: map the whole project at a glance (nothing runs), then record a real run and replay it step by step, zooming from the architecture dow\n[truncated]"
author: "arnoldpredator"
captured_at: "2026-08-08T16:20:52Z"
capture_tool: "hn-digest"
hn_id: 49223208
score: 1
comments: 1
posted_at: "2026-08-08T16:19:31Z"
tags:
  - hacker-news
  - translated
---

# A tool to quickly step through and audit LLM-generated Python codebases

- HN: [49223208](https://news.ycombinator.com/item?id=49223208)
- Source: [github.com](https://github.com/arnoldpredator/pyreplay)
- Score: 1
- Comments: 1
- Posted: 2026-08-08T16:19:31Z

## Translation

タイトル: LLM で生成された Python コードベースを迅速にステップ実行して監査するツール
記事のタイトル: GitHub - arnoldpredator/pyreplay: LLM はコードベース全体を数秒で生成できます。それが遅い部分であることを理解してください。 pyreplay は最初のパスを高速にします — トップダウン: 一目でプロジェクト全体をマッピングし (何も実行されません)、その後、実際の実行を記録して、アーキテクチャからズームして段階的に再生します。
[切り捨てられた]
説明: LLM はコードベース全体を数秒で生成できます。それが遅い部分であることを理解してください。 pyreplay を使用すると、最初のパスがトップダウンで高速になります。一目でプロジェクト全体をマップし (何も実行されません)、実際の実行を記録して段階的に再生し、アーキテクチャから任意の 1 つの変数の値までズームダウンします。
[切り捨てられた]

記事本文:
GitHub - arnoldpredator/pyreplay: LLM はコードベース全体を数秒で生成できます。それが遅い部分であることを理解してください。 pyreplay を使用すると、最初のパスがトップダウンで高速になります。一目でプロジェクト全体をマップし (何も実行されません)、実際の実行を記録して段階的に再生し、アーキテクチャから任意の 1 つの変数の値までズームダウンします。自己完結型の HTML、依存関係なし。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アーノルドプレデター
/
パイリプレイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

イオン
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
86 コミット 86 コミット .github .github スクリーンショット スクリーンショット tinyshop tinyshop ツアー ツアー .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md FEATURES.md FEATURES.md ライセンス ライセンス README.md README.md TUTORIAL.md TUTORIAL.md _pyreplay_pytest_plugin.py _pyreplay_pytest_plugin.py bubble_sort.py bubble_sort.py checks.py checks.py example_bigarray.py example_bigarray.py example_control.py example_control.py example_dp.py example_dp.py example_dunder.py example_dunder.py example_例外.py example_例外.py example_flaky.py example_flaky.py example_fuzz.py example_fuzz.py example_graph.py example_graph.py example_heavy.py example_heavy.py example_histogram.py example_histogram.py example_machinery.py example_machinery.py example_mro.py example_nan.py example_nan.py example_prefix.py example_prefix.py example_race.py example_race.py example_sort.py example_sort.py example_tasks.py example_tasks.py example_threads.py example_threads.py example_watch.py example_watch.py map_template.html map_template.html mapper.py mapper.py restarter_template.html restarter_template.html run_template.html run_template.html スイープ_テンプレート.html スイープ_テンプレート.html スイープ_テンプレート.html スイープ_テンプレート.html トレーサー.py トレーサー.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Python が実際に何をしているのかを見てみましょう。依存性ゼロのトレーサーと静的
Python の実行 (またはコードベース全体) を探索可能なものに変えるマッパー
自己完結型の HTML ページ。これは、増大し続ける問題に対応するために構築されています。
LLM はプロジェクト全体を数秒で生成できますが、その内容をすぐに把握できます。
コードは実際には を実行しますが、それをどのように組み合わせるかは簡単ではありません。パイリプレイ
最初のパスを高速化します。実行を記録し、ステップごとに再生します。または
コードバ全体をマッピングする

一目でわかります。バグは修正されません（それが問題です）
プロフェッショナル IDE は次の場合に適しています) — 特にコードやコードを素早く確認できます。
LLMは書きました。
トレーサーと地図として始まったものが天文台に成長しました : ~126
構成可能な楽器 (すべて FEATURES.md にカタログされています)、
ファネルとして注文 — 無料で全体をマッピングし、実行を記録し、
単一の値に下がります。各ステージで次のステージが与えられます。
2 つのツールは 1 つの JSON イベント ログ形式を共有します。これらは次の 2 つの方法です。
tracer.py — 実行を記録します (すべての行/呼び出し/戻り、および
変数が変更されました）をステップ実行する自己完結型のtrace_*.htmlに変換します
ビデオのように。データ構造ごとのセマンティック ビュー (リストはセルの行、
グラフはノードとエッジです)、非同期タスクは並列レーンとして表示されます。
Perfetto タイムラインのエクスポート。すべてのトレースはコンソールも記録します (
印刷された行→書いた瞬間）と再現性を埋め込みます
Capsule: 正確なコマンド、環境の事実、必要な消費された標準入力
もう一度実行します。
mapper.py — ast を使用してコードベースを読み込みます (何も実行されません)。
ズーム可能なmap_*.html : インポートの深さによってレイアウトされたモジュール、パッケージ
赤で描かれた折り畳み、輸入サイクル、および「耐力壁」のランク付け方法
多くのモジュールはそれらをインポートします。トレースをオーバーレイして、どの部分が実際に実行されたかを確認します。
これら 2 つのエントリ ポイントの背後には、以下を構成するツールのツールキットがあります。
それらに到達する方法をグループ化します (ファネル - まず広くて安い):
コードベースをマップします (何も実行されません) — 耐力壁、インポート サイクル、
コードベース全体の検索、パッケージの折りたたみ、コール グラフ。
記録と再生 — セマンティック ビュー、スレッド、および
平行レーンとしての asyncio、ほぼ無料のブラックボックス フライト レコーダー、より高速な
大規模な実行のための sys.monitoring エンジン。
実行結果を読む — 来歴 (この値はどこから来たのか?)

、
Whyline (なぜこの行は実行されなかったのですか?)、分岐判定、バイトコードの分析、
コールツリーとシーケンス図。
真実と警報 — NaN 誕生のトリップワイヤー、コントラクト ( --invariant )、
不変条件をマイニングし、状態マシンを観察しました。
信頼性ラボ — N 回実行して統計を取得し、2 回の実行と結果を比較します。
最初の発散、失敗した入力を縮小、失敗を見つけるためのファズ、
基準に対する差動テスト、障害の挿入、監視
メモリが増加し、プログラムが変更した内容をリストします。
各仕上げ段階では、次に貼り付ける正確なコマンド (ツール) が出力されます。
次の楽器を覚えさせるのではなく、次の楽器を渡します。
クイックスタート - インストールなし、標準ライブラリのみ
python3トレーサー.py your_script.py # ->trace_your_script.html
python3 mapper.py パス/to/プロジェクト # -> map_project.html
python3 trader.py --runs 20 flaky.py # 20 回実行: 結果の統計、
# 動作ごとに 1 つのトレースが保持されます
python3 trader.py --diverge good.html bad.html # 2 つが実行される最初のイベント
# 道が分かれる (原因、次に症状)
python3 trader.py --trip nan sim.py # 最初の NaN が誕生した場所
python3 trader.py --check " total < 0 " app.py # としての実行に関する質問
# 終了コード -> git bisect のオラクル
python3 trader.py --black-box server.py # フライト レコーダー: 最後のリング
# N イベント; kill -USR1 = ライブスナップショット
python3トレーサー.py --スイープ " n=8,16,32,64 " \
--predict " n^2 " algo.py # 倍加実験: 観察されました
# 成長指数 + 請求の評決
python3トレーサー.py --relation \
" ' '.join(reversed(x.split())) => out == out0 " \
--gen gen.py algo.py # 変成テスト: 対称性
# はオラクルです。違反を続ける
# 両方のトレース、 --diverge の準備ができています
python3 trader.py --fuzz gen.py --runs 50 t.py # 実行中に失敗した入力を見つけます
# スリープ (シード、再現可能)、
# その後、あなたに渡されます --shrink
python3トレーサー.py --oracle brute.py高速

.py\
< input.txt # 差分テスト: ブルートフォース
# は仕様です。不一致 -> --diverge
python3 trader.py --io app.py # 何をタッチしましたか?ファイル、ソケット、
# サブプロセス;という名前の閉じられていないファイル
python3 trader.py --memory sim.py # メモリが保持される場所: 増加
# カーブ + モジュールごとのバイト
任意のブラウザで HTML を開きます。サーバーもビルドステップも依存関係もありません。
TUTORIAL.md の完全なガイド。すべての機能を
FEATURES.md のスクリーンショット。
これは Python Tutor / VizTracer とどう違うのですか?
Python Tutor は、サンドボックス内の小さなスニペットを視覚化します。 pyreplayは実際のトレースを追跡します
自分のマシン上でスクリプトを作成し、
静的マップとセマンティック ズーム。
VizTracer は、通話が発生したときのタイムラインを最初にプロファイリングします。
pyreplay は理解が第一です — 値を表示します: リストのどの要素であるか
名前が再バインドされたか、オブジェクトがその場で変更されたかに関係なく、変更されました。
ブランチの決定とその理由。
正直契約。 pyreplay は実際に知っているものだけをマークします
変わりました。部分的または不明な状態はマークされず、推測されることもありません。すべてのキャップ
切り捨ては画面上でアナウンスされます。確信が持てない場合は、そのように表示されます。
アーキテクチャ — および別の言語を追加する方法
トレーサー (Python: sys.settrace / sys.monitoring)
│
▼
JSON イベント ログ ◀── コントラクト。言語中立。
│
▼
HTML レンダラー (バニラ JS、フレームワークなし)
イベント ログが重要です。レンダラーは、イベント ログが何が生成されたかは気にしません。
ビューアを操作せずに別の言語のサポートを追加できます —
C++ / Rust / JS トレーサと既存のリプレイヤから同じ JSON を出力します
それを再生します。静的マップには一致するシームがあります (ast をパーサーと交換します)
ターゲット言語)。どちらもレイアウトされています
貢献.md 。
python3 checks.py # 110 データレベルのチェック — すべて緑色で出力されるはずです
あらゆる機能

チェックで固定されています。変更の前後に実行します。ある
グリーンスイートは、寄付を誠実に保つ契約です。
バグレポート、エッジケース、より多くのサンプルプログラム、その他の言語が含まれています。
皆さん大歓迎です。 COTRIBUTING.md が開始点です - 1 つのファイル
何を構築するかのロードマップ (構築する 12 の機能と「最初に良いもの」)
リスト）、基本ルール、および新しい言語バックエンドが発行するイベント ログ スキーマ。
pyreplay は無料で、MIT ライセンスを取得しています。理解する時間を節約できた場合は、
コードベース、コーヒー買ってきてもいいよ ☕
— ロードマップを前進させ続けます。
Python 3.10 以降 (3.12 で開発) が必要です。 --backend モニタリングの高速化
レコーダーは PEP 669 を使用し、3.12 以降が必要です。
LLM はコードベース全体を数秒で生成できます。それが遅い部分であることを理解してください。 pyreplay を使用すると、最初のパスがトップダウンで高速になります。一目でプロジェクト全体をマップし (何も実行されません)、実際の実行を記録して段階的に再生し、アーキテクチャから任意の 1 つの変数の値までズームダウンします。自己完結型の HTML、依存関係なし。
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An LLM can generate a whole codebase in seconds; understanding it is the slow part. pyreplay makes that first pass fast — top-down: map the whole project at a glance (nothing runs), then record a real run and replay it step by step, zooming from the architecture down to any single variable's value.
[truncated]

GitHub - arnoldpredator/pyreplay: An LLM can generate a whole codebase in seconds; understanding it is the slow part. pyreplay makes that first pass fast — top-down: map the whole project at a glance (nothing runs), then record a real run and replay it step by step, zooming from the architecture down to any single variable's value. Self-contained HTML, zero dependencies. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
arnoldpredator
/
pyreplay
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
86 Commits 86 Commits .github .github screenshots screenshots tinyshop tinyshop tours tours .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md FEATURES.md FEATURES.md LICENSE LICENSE README.md README.md TUTORIAL.md TUTORIAL.md _pyreplay_pytest_plugin.py _pyreplay_pytest_plugin.py bubble_sort.py bubble_sort.py checks.py checks.py example_bigarray.py example_bigarray.py example_control.py example_control.py example_dp.py example_dp.py example_dunder.py example_dunder.py example_exceptions.py example_exceptions.py example_flaky.py example_flaky.py example_fuzz.py example_fuzz.py example_graph.py example_graph.py example_heavy.py example_heavy.py example_histogram.py example_histogram.py example_machinery.py example_machinery.py example_mro.py example_mro.py example_nan.py example_nan.py example_prefix.py example_prefix.py example_race.py example_race.py example_sort.py example_sort.py example_tasks.py example_tasks.py example_threads.py example_threads.py example_watch.py example_watch.py map_template.html map_template.html mapper.py mapper.py replayer_template.html replayer_template.html runs_template.html runs_template.html sweep_template.html sweep_template.html tracer.py tracer.py View all files Repository files navigation
See what Python is actually doing. A zero-dependency tracer and static
mapper that turn a Python run — or a whole codebase — into an explorable,
self-contained HTML page. It's built for a problem that keeps growing: an
LLM can generate a whole project in seconds, but quickly grasping what that
code actually does , and how it all fits together, is not easy. pyreplay
makes that first pass fast — record a run and replay it step by step, or
map the whole codebase at a glance. It won't fix your bugs (that's what a
professional IDE is for) — it's the fast first look, especially at code an
LLM wrote.
What began as a tracer and a map has grown into an observatory : ~126
composable instruments (all catalogued in FEATURES.md ),
ordered as a funnel — map the whole thing for free, record a run, then
descend to a single value. Each stage hands you the next.
Two tools share one JSON event-log format — they're the two ways in:
tracer.py — records a run (every line / call / return, and which
variables changed) into a self-contained trace_*.html you step through
like a video. Semantic views per data structure (a list is a row of cells,
a graph is nodes and edges), asyncio tasks shown as parallel lanes, and
Perfetto timeline export. Every trace also records the console (click a
printed line → the moment that wrote it) and embeds a reproducibility
capsule: the exact command, environment facts and consumed stdin needed
to run it again.
mapper.py — reads a codebase with ast ( nothing is executed ) into
a zoomable map_*.html : modules laid out by import depth, packages that
fold, import cycles drawn in red, and the "load-bearing walls" ranked by how
many modules import them. Overlay a trace to see which parts actually ran.
Behind those two entry points is a toolkit of instruments that compose ,
grouped the way you reach for them (the funnel — wide and cheap first):
Map the codebase (nothing runs) — load-bearing walls, import cycles,
whole-codebase search, package folding, the call graph.
Record & replay — every value change with semantic views, threads &
asyncio as parallel lanes, a near-free black-box flight recorder, a faster
sys.monitoring engine for big runs.
Read the run — provenance ( where did this value come from? ), the
whyline ( why didn't this line run? ), branch verdicts, bytecode anatomy,
call tree & sequence diagram.
Truth & alarms — a NaN-birth tripwire, contracts ( --invariant ),
mined invariants, observed state machines.
The reliability lab — run N times for stats, diff two runs to the
first divergence, shrink a failing input, fuzz to find a failure ,
differential-test against a reference , inject faults , watch
memory grow , list what the program touched .
Each finishing stage prints the exact next command to paste — the tool
hands you the next instrument rather than making you remember it.
Quickstart — no install, standard library only
python3 tracer.py your_script.py # -> trace_your_script.html
python3 mapper.py path/to/project # -> map_project.html
python3 tracer.py --runs 20 flaky.py # run it 20x: outcome stats,
# one kept trace per behavior
python3 tracer.py --diverge good.html bad.html # first event where two runs
# part ways (cause, then symptom)
python3 tracer.py --trip nan sim.py # where the first NaN was BORN
python3 tracer.py --check " total < 0 " app.py # any question about a run as an
# exit code -> git bisect's oracle
python3 tracer.py --black-box server.py # flight recorder: ring of the LAST
# N events; kill -USR1 = live snapshot
python3 tracer.py --sweep " n=8,16,32,64 " \
--predict " n^2 " algo.py # the doubling experiment: observed
# growth exponent + claim verdict
python3 tracer.py --relation \
" ' '.join(reversed(x.split())) => out == out0 " \
--gen gen.py algo.py # metamorphic testing: the symmetry
# is the oracle; violations keep
# both traces, ready to --diverge
python3 tracer.py --fuzz gen.py --runs 50 t.py # find a failing input while you
# sleep (seeded, reproducible),
# then it hands you --shrink
python3 tracer.py --oracle brute.py fast.py \
< input.txt # differential test: the brute force
# is the spec; mismatch -> --diverge
python3 tracer.py --io app.py # what did it TOUCH? files, sockets,
# subprocesses; unclosed files named
python3 tracer.py --memory sim.py # where memory is RETAINED: a growth
# curve + per-module bytes
Open the HTML in any browser. No server, no build step, no dependencies.
Full guide in TUTORIAL.md ; every feature explained with a
screenshot in FEATURES.md .
How is this different from Python Tutor / VizTracer?
Python Tutor visualizes small snippets in a sandbox. pyreplay traces real
scripts on your own machine and scales up to whole codebases through the
static map and semantic zoom.
VizTracer is profiling-first — a timeline of when calls happened.
pyreplay is understanding-first — it shows values : which element of a list
changed, whether a name was rebound or the object mutated in place, what a
branch decided and why.
The honesty contract. pyreplay marks only what it actually knows
changed. Partial or unknown state is left unmarked, never guessed; every cap
and truncation is announced on screen. If it can't be sure, it says so.
Architecture — and how to add another language
tracer (Python: sys.settrace / sys.monitoring)
│
▼
JSON event log ◀── the contract. language-neutral.
│
▼
HTML renderer (vanilla JS, no framework)
The event log is the whole point: the renderer doesn't care what produced it.
You can add support for another language without touching the viewer —
emit the same JSON from a C++ / Rust / JS tracer and the existing replayer
plays it back. The static map has a matching seam (swap ast for a parser of
your target language). Both are laid out in
CONTRIBUTING.md .
python3 checks.py # 110 data-level checks — should print all green
Every feature is pinned by a check. Run it before and after any change; a
green suite is the contract that keeps contributions honest.
Bug reports, edge cases, more example programs, and other languages are
all very welcome. CONTRIBUTING.md is the place to start — one file
with the roadmap of what to build (12 features to build, plus a "good first"
list), the ground rules, and the event-log schema a new-language backend emits.
pyreplay is free and MIT-licensed. If it saved you time understanding a
codebase, you can buy me a coffee ☕
— it keeps the roadmap moving.
Requires Python 3.10+ (developed on 3.12). The faster --backend monitoring
recorder uses PEP 669 and needs 3.12+.
An LLM can generate a whole codebase in seconds; understanding it is the slow part. pyreplay makes that first pass fast — top-down: map the whole project at a glance (nothing runs), then record a real run and replay it step by step, zooming from the architecture down to any single variable's value. Self-contained HTML, zero dependencies.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
