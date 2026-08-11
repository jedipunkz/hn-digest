---
source: "https://github.com/bryancostanich/verifrog"
hn_url: "https://news.ycombinator.com/item?id=49261301"
title: "Verifrog – a Verilog/SystemVerilog testing framework an AI agent can drive"
article_title: "GitHub - bryancostanich/verifrog: verilog testing framework written in F#. · GitHub"
author: "bryancostanich"
captured_at: "2026-08-11T17:49:42Z"
capture_tool: "hn-digest"
hn_id: 49261301
score: 2
comments: 0
posted_at: "2026-08-11T17:08:12Z"
tags:
  - hacker-news
  - translated
---

# Verifrog – a Verilog/SystemVerilog testing framework an AI agent can drive

- HN: [49261301](https://news.ycombinator.com/item?id=49261301)
- Source: [github.com](https://github.com/bryancostanich/verifrog)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T17:08:12Z

## Translation

タイトル: Verifrog – AI エージェントが駆動できる Verilog/SystemVerilog テスト フレームワーク
記事のタイトル: GitHub - bryancostanich/verifrog: F# で書かれた verilog テスト フレームワーク。 · GitHub
説明: F# で書かれた verilog テスト フレームワーク。 GitHub でアカウントを作成して、bryancostanich/verifrog の開発に貢献してください。

記事本文:
GitHub - bryancotanich/verifrog: F# で書かれた verilog テスト フレームワーク。 · GitHub
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
ブライアンコスタニッチ
/
ベリフロッグ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
89 コミット 89 コミット .claude .claude .github/ workflows .github/ workflows bin bin conductor/ プロジェクト/ v1_release conductor/ プロジェクト/ v1_release docs docs サンプル サンプル src src テスト/ Verifrog.Tests テスト/ Veri

frog.Tests .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
F# のオープンソース Verilog/SystemVerilog テストおよびデバッグ フレームワーク。すべての信号、メモリ、レジスタへのタイプ セーフで構造化されたアクセスにより、Verilator および Icarus Verilog を通じて RTL を駆動し、MCP 経由で AI エージェントにデバッグさせます。
従来の Verilog テストベンチは退屈で表現力に欠けます。 UVM は強力ですが重いです。 Verifrog は、最新の言語の人間工学を備えた Verilator のスピードを提供します。つまり、仕様を読み取るテスト、チェックポイント/リストア、フォークベースの仮説テスト、および MCP サーバーにより、エージェントが設計を開いてステップを踏み、シグナルを強制し、バグを独自に特定できるようになります。
Verifrog には MCP サーバーが同梱されているため、AI エージェントが RTL シミュレーションを直接操作できます。つまり、デザインを開き、サイクルをステッピングし、信号を読み取って強制し、状態をチェックポイントして仮説をテストします。このデモでは、テストが失敗すると、エージェントが調査、チェックポイント、シグナルの強制を行い、ALU の SUB 操作のバグを特定します。
verifrog-debugging-demo-2x.mp4
特長
Verilator を介して RTL をコンパイルして実行する、構造化された読みやすいテストを F# で作成します。信号、メモリ、レジスタを名前で読み書きします。明確な失敗メッセージとともに値をアサートします。単一の dotnet テストで Verilator テストと Icarus Verilog テストを並べて実行します。
Verifrog のシミュレーション モデルはコードから完全に制御可能です。任意の時点で一時停止し、デザイン内のすべての信号を検査し、サイクルごとにステップを進めることができます。しかし、本当の威力は、その上に構築されたツールにあります。
チェックポイント/復元 — シミュレーション状態全体 (すべてのレジスタ、すべてのメモリ セル) のスナップショットを作成し、マイクロ秒単位で後で復元します。サイクル 50,000 でバグに遭遇しましたか?サ

障害が発生する前にチェックポイントを取得し、シミュレーションを最初から再実行することなく、さまざまな信号の復元と調査を繰り返します。
フォーク — what-if シナリオを検討し、自動的にスナップバックします。 「この信号を強制的にハイにしたらどうなるでしょうか?」 Fork は実験を実行し、結果を取得し、元の状態を復元します。そのため、手動で保存/復元しなくても、同じ時点から複数の仮説を試すことができます。
比較とスイープ — 同じ状態から 2 つの構成を並べて実行する ([比較] )、または多くの値にわたってパラメーターをスイープする ([スイープ] )。どちらも内部でチェックポイントを使用して、各シナリオが同じ状態から開始されるようにします。
信号強制 — 内部信号をオーバーライドし、クロック サイクル全体にわたって保持します。フォルトを挿入し、クロック ゲーティングを無効にし、バス値を強制します。その後、解放してデザインが回復するのを確認します。
トレースと RunUntil — サイクルのウィンドウにわたって信号値を記録するか ( Trace )、条件が満たされるまでシミュレーションを進めます ( RunUntil 、 RunUntilSignal )。何サイクルステップするかを推測する必要はもうありません。
VCD 波形解析 - シミュレーション波形ダンプを解析し、プログラムでクエリを実行します。信号が最初に変化したときの検出、パルスのカウント、タイミング関係の確認、FSM 状態カバレッジの検証を行います。テストで使用するためのライブラリ ( Verifrog.Vcd ) と、迅速な分析のためのコマンドライン ツール ( verifrog-vcd ) の両方として利用できます。
ベリフロッグ
生のベリレーター (C++)
ココットブ
UVM
テストベンチ言語
F# (Expecto) + 宣言型 .verifrog
C++
パイソン
システムVerilog
シミュレーションエンジン
ベリレーター/イカロス
検証者
任意の VPI シミュレータ
任意のSVシミュレータ
チェックポイント/復元
組み込み (μs 単位の状態スナップショット)
手巻き
いいえ
いいえ
フォーク/比較/スイープ
内蔵
手巻き
いいえ
マニュアル
信号の強制
名前ごとに、サイクル全体で保持されます
手動ポインタ書き込み
.valueの割り当て
uvm_hdl_force
名前付きメモリ/レジスタ

アクセス
TOML 駆動 (名前による)
マニュアル
マニュアル
構成DB/RAL
AI エージェントのデバッグ (MCP)
内蔵
いいえ
いいえ
いいえ
学習曲線
低い
中
低い
高
Verifrog は素の速度を Verilator に依存しているため、Verilator の 2 状態制限を継承しています。タイミングが正確な 4 状態 (X/Z) テストベンチでは、同じテスト スイートで Icarus Verilog を駆動することもできます。デフォルトとしてサイクル精度の X 伝播が必要な場合は、cocotb のような VPI ベースのフローの方が適している可能性があります。ループ内のエージェントを使用してチェックポイント/フォーク駆動の高速デバッグが必要な場合、それが Verifrog の目的です。
Clang++ (macOS、Xcode に付属) または g++ (Linux)
Icarus Verilog (オプション、タイミングが正確なテストベンチ用)
git クローン https://github.com/bryancostanich/verifrog.git
CD ベリフロッグ
./install.sh # verifrog を /usr/local/bin にシンボリックリンクします
または、手動で bin/ を PATH に追加します。export PATH="/path/to/verifrog/bin:$PATH"
verifrog ビルド サンプル/カウンター
verifrog テストサンプル/カウンター
新しいプロジェクトを開始する
プロジェクトの cd
verifrog の初期化。
# デザインに合わせて verifrog.toml を編集し、次のようにします。
ベリフロッグビルド
ベリフロッグテスト
段階的なチュートリアルについては、完全なスタート ガイドを参照してください。
ほとんどのハードウェア テストは、「信号を設定し、ステップを実行し、チェックする」だけです。これらを .verifrog ファイルに宣言的に記述します。F# は必要ありません。
テスト「有効にすると 10 までカウント」[スモーク]:
書き込みイネーブル = 1
ステップ10
期待数 == 10
テスト「ロードしてからカウント」[ユニット]:
書き込みload_value = 42、load_en = 1
ステップ1
書き込みload_en = 0、enable = 1
ステップ5
期待数 == 47
または、より詳細な制御が必要な場合は F# で次のようにします。
ベリフロッグを開きます。シム
ベリフロッグを開きます。ランナー
let テスト = testList " カウンタ " [
テスト " 有効にすると 10 までカウントされます " {
sim = SimFixture.create() を使用します。
sim.Write ( "enable " , 1 L ) |> 無視します
シミュレーションステップ ( 10 )
Expect.signal sim " count " 10 L " count は 10 に達するはずです "
}
】
どちらも同じテスト スイート、同じカテゴリで実行されます

つまり、同じ --report 、同じ verifrog test です。
チェックポイントとフォークを使用したデバッグ
シミュレーション状態の保存、前方への実行、復元、別のことの試行はすべてコード内で行われます。
テスト "オーバーフロー動作を調査する" {
sim = SimFixture.create() を使用します。
sim.Write ( "enable " , 1 L ) |> 無視します
sim.ステップ ( 200 )
// 興味深い部分の直前の状態を保存します
let cp = sim.SaveCheckpoint ( " before_overflow " )
// 前に走って観察する
シミュレーションステップ ( 60 )
let count = sim.ReadOrFail ( " count " )
オーバーフローさせます = sim.ReadOrFail ( " オーバーフロー " )
printfn "さらに 60 サイクル後: count= %d overflow= %d " count がオーバーフローしました
// 復元して別のアプローチを試す
sim.RestoreCheckpoint ( " before_overflow " )
// 限界に近い値をロードしたらどうなるでしょうか?
let result = sim.Fork ( fun s ->
s.Write ( "load_en " , 1 L ) |> 無視する
s.Write ( "load_value " , 250 L ) |> 無視
s.ステップ(1)
s.Write ( "load_en " , 0 L ) |> 無視する
s.ステップ ( 10 )
s.ReadOrFail ( " オーバーフロー " ))
// SIM は「before_overflow」に戻ります — フォークは自動的に復元されます
// 複数の荷重値をスイープして境界を見つけます
let results = sim.スイープ(
[ 248L ; 249リットル; 250L; 251リットル; 252L]、
楽しいloadVal s ->
s.Write ( "load_en " , 1 L ) |> 無視する
s.Write ( "load_value " ,loadVal ) |> 無視
s.ステップ(1)
s.Write ( "load_en " , 0 L ) |> 無視する
s.ステップ ( 10 )
s.ReadOrFail ( " オーバーフロー " ))
for (loadVal , overflow ) の結果は do
printfn "load= %d -> overflow= %d "loadVal オーバーフロー
}
波形の解析
テスト "VCD 解析でタイミングを検証" {
sim = SimFixture.create() を使用します。
// ... スティミュラスを実行 ...
let vcd = VcdParser.parseAll "output/sim.vcd"
// FSM が最初に状態 5 になったのはいつですか?
let t = VcdParser.firstTimeAtValue vcd " fsm_state " 5
// オーバーフローは何回パルスしたか?
let パルス = VcdParser.highPulseCount vcd " counter.overflow

」
// FSM はどの州を訪問しましたか?
let states = VcdParser.uniqueValues vcd " fsm_state "
}
試験機関
Verifrog は、適切なタイミングで適切なテストを実行できるように、ハードウェア ドメインのテスト カテゴリを提供します。
verifrog test --category Smoke # Quick sanity — デザインは生きています (秒)
verifrog test --category ユニット番号 集中信号/ブロック テスト
verifrog test --category 統合 # マルチブロック データ フロー
verifrog test --category パラメトリック # スイープと値の範囲
verifrog テスト # すべて
カテゴリは軽量の testList ラッパーです。テストをグループ化するだけです。
ベリフロッグを開きます。ランナー。カテゴリ
let テスト = testList " MySoC " [
煙 [
テスト「リセットから復帰」{ ... }
】
単位 [
テスト「カウンタの増分」{ ... }
】
黄金色 [
テスト " 参照出力と一致します " { ... }
】
】
また、ストレス (長時間実行)、ゴールデン (参照出力)、回帰 (バグ修正カバレッジ) も利用できます。
あなたのテストプロジェクト (期待)
|
v
Verifrog.Runner — SimFixture、Iverilog バックエンド、Expect ヘルパー
|
v
Verifrog.Sim — Sim タイプ、メモリ/レジスタ アクセサ、TOML 構成
|
v
libverifrog_sim — 汎用 Verilator C++ ラッパー (デザインごとに構築)
|
v
Verilator — コンパイルされた RTL
コンポーネント
図書館
何をするのか
Verifrog.Sim
コア シミュレーション API: 作成、ステップ、信号の読み取り/書き込み、チェックポイント/復元、強制、フォーク/スイープ、メモリ/レジスタ アクセス
ベリフロッグランナー
テストインフラストラクチャ: SimFixture ライフサイクル、Iverilog バックエンド、Expect アサーション、テスト カテゴリ (Smoke/Unit/Parametric/Integration/Stress/Golden/Regression)
Verifrog.Vcd
スタンドアロン VCD 波形パーサー: ファイルの解析、クエリ信号、時刻値、遷移、タイミング分析
Verifrog.Vcd.Cli
テキストと JSON 出力を備えたコマンドライン VCD 分析ツール
verifrog CLI
ビルドツール: init 、 build 、 clean 、 test 、 debug (対話型 REPL)、 debug-server (JSON)、 mcp-server (MCP fo

クロード）
libverifrog_sim
設計に依存しない C++ シム: シグナル検出、ダイレクト ポインター アクセス、memcpy 経由のチェックポイント
VS コード拡張
.verifrog ファイルの構文ハイライト、シグナル パネル、チェックポイント パネル、デバッグ ツールバー (実験的)
構成
すべてのプロジェクト設定は verifrog.toml に存在します。
【デザイン】
トップ = " my_module "
ソース = [ " rtl/*.v " ]
【テスト】
出力 = "ビルド"
[思い出。データラム]
パス = " u_ram.mem "
バンク = 1
深さ = 1024
幅 = 32
[ レジスタ ]
パス = " u_regfile.regs "
幅 = 8
[ レジスタ 。地図】
CTRL = 0x00
ステータス = 0x01
データ = 0x02
完全な構成リファレンスを参照してください。
サンプル
それが示すもの
カウンター
最小限: ステップ、読み取り/書き込み、チェックポイント、強制、フォーク
alu_regfile
TOML レジスタ マップ、名前付きレジスタ アクセス、パラメトリック スイープ
スラム
TOML メモリ領域、バンク アクセス、バックドア ローディング
iverilog_tb
デュアル バックエンド: 1 つのドットネット テストでの Verilator + iverilog
i2c_bfm
自動検出、タイミング精度の高い I2C を備えたプロトコル レベルの BFM
デバッグ
シミュレーションをデバッグする複数の方法:
インタラクティブ REPL — 最速のパス。シミュレーションのステップ実行、信号の読み取り/書き込み、チェックポイントの設定、値の強制をすべてコマンド ラインから行います。
ベリフロッグのデバッグ
sim>書き込みイネーブル1
シム>ステップ10
sim> 読み取りカウント
カウント = 10
sim> チェックポイント before_overflow
シ

[切り捨てられた]

## Original Extract

verilog testing framework written in F#. Contribute to bryancostanich/verifrog development by creating an account on GitHub.

GitHub - bryancostanich/verifrog: verilog testing framework written in F#. · GitHub
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
bryancostanich
/
verifrog
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
89 Commits 89 Commits .claude .claude .github/ workflows .github/ workflows bin bin conductor/ projects/ v1_release conductor/ projects/ v1_release docs docs samples samples src src tests/ Verifrog.Tests tests/ Verifrog.Tests .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
An open-source Verilog/SystemVerilog testing and debugging framework in F#. Drive your RTL through Verilator and Icarus Verilog with type-safe, structured access to every signal, memory, and register — and let an AI agent debug it for you over MCP.
Traditional Verilog testbenches are tedious and inexpressive; UVM is powerful but heavyweight. Verifrog gives you Verilator's speed with the ergonomics of a modern language — tests that read like specifications, checkpoint/restore and fork-based hypothesis testing, and an MCP server so an agent can open your design, step it, force signals, and pinpoint a bug on its own.
Verifrog ships an MCP server, so an AI agent can drive your RTL simulation directly — opening a design, stepping cycles, reading and forcing signals, and checkpointing state to test hypotheses. In this demo, a failing test leads the agent to investigate, checkpoint, force signals, and pinpoint a bug in the ALU's SUB operation:
verifrog-debugging-demo-2x.mp4
Features
Write structured, readable tests in F# that compile and run your RTL through Verilator. Read and write any signal, memory, or register by name. Assert values with clear failure messages. Run Verilator and Icarus Verilog tests side-by-side under a single dotnet test .
Verifrog's simulation model is fully controllable from code — you can pause at any point, inspect every signal in the design, and step forward cycle-by-cycle. But the real power is in the tools built on top of this:
Checkpoint/Restore — Snapshot the entire simulation state (every register, every memory cell) and restore it later in microseconds. Hit a bug at cycle 50,000? Save a checkpoint before the failure, then repeatedly restore and probe different signals without re-running the simulation from scratch.
Fork — Explore a what-if scenario and automatically snap back. "What would happen if I forced this signal high?" Fork runs your experiment, captures the result, and restores the original state — so you can try multiple hypotheses from the same point without manual save/restore.
Compare and Sweep — Run two configurations side-by-side from the same state ( Compare ), or sweep a parameter across many values ( Sweep ). Both use checkpoints internally to ensure each scenario starts from identical state.
Signal forcing — Override any internal signal and hold it across clock cycles. Inject faults, disable clock gating, force a bus value — then release and watch the design recover.
Tracing and RunUntil — Record signal values over a window of cycles ( Trace ), or advance the simulation until a condition is met ( RunUntil , RunUntilSignal ). No more guessing how many cycles to step.
VCD waveform analysis — Parse simulation waveform dumps and query them programmatically: find when a signal first changed, count pulses, check timing relationships, verify FSM state coverage. Available as both a library ( Verifrog.Vcd ) for use in tests and a command-line tool ( verifrog-vcd ) for quick analysis.
Verifrog
Raw Verilator (C++)
cocotb
UVM
Testbench language
F# (Expecto) + declarative .verifrog
C++
Python
SystemVerilog
Simulation engine
Verilator / Icarus
Verilator
any VPI simulator
any SV simulator
Checkpoint / restore
Built in (state snapshot in µs)
Hand-rolled
No
No
Fork / compare / sweep
Built in
Hand-rolled
No
Manual
Signal forcing
By name, held across cycles
Manual pointer writes
.value assignment
uvm_hdl_force
Named memory / register access
TOML-driven, by name
Manual
Manual
Config DB / RAL
AI agent debugging (MCP)
Built in
No
No
No
Learning curve
Low
Medium
Low
High
Verifrog leans on Verilator for raw speed, so it inherits Verilator's two-state limitation — for timing-accurate, four-state (X/Z) testbenches it can also drive Icarus Verilog under the same test suite. If you need cycle-accurate X-propagation as the default, a VPI-based flow like cocotb may fit better; if you want fast checkpoint/fork-driven debugging with an agent in the loop, that's where Verifrog is aimed.
clang++ (macOS, included with Xcode) or g++ (Linux)
Icarus Verilog (optional, for timing-accurate testbenches)
git clone https://github.com/bryancostanich/verifrog.git
cd verifrog
./install.sh # Symlinks verifrog to /usr/local/bin
Or add bin/ to your PATH manually: export PATH="/path/to/verifrog/bin:$PATH"
verifrog build samples/counter
verifrog test samples/counter
Start a new project
cd your-project
verifrog init .
# Edit verifrog.toml with your design, then:
verifrog build
verifrog test
See the full Getting Started Guide for a step-by-step walkthrough.
Most hardware tests are just "set signals, step, check." Write those declaratively in a .verifrog file — no F# needed:
test "counts to 10 when enabled" [Smoke]:
write enable = 1
step 10
expect count == 10
test "load then count" [Unit]:
write load_value = 42, load_en = 1
step 1
write load_en = 0, enable = 1
step 5
expect count == 47
Or in F# when you need more control:
open Verifrog. Sim
open Verifrog. Runner
let tests = testList " counter " [
test " counts to 10 when enabled " {
use sim = SimFixture.create ()
sim.Write ( " enable " , 1 L ) |> ignore
sim.Step ( 10 )
Expect.signal sim " count " 10 L " count should reach 10 "
}
]
Both run in the same test suite — same categories, same --report , same verifrog test .
Debugging with checkpoints and Fork
Save simulation state, run forward, restore, try something different — all in code:
test " investigate overflow behavior " {
use sim = SimFixture.create ()
sim.Write ( " enable " , 1 L ) |> ignore
sim.Step ( 200 )
// Save state right before the interesting part
let cp = sim.SaveCheckpoint ( " before_overflow " )
// Run forward and observe
sim.Step ( 60 )
let count = sim.ReadOrFail ( " count " )
let overflowed = sim.ReadOrFail ( " overflow " )
printfn " After 60 more cycles: count= %d overflow= %d " count overflowed
// Restore and try a different approach
sim.RestoreCheckpoint ( " before_overflow " )
// What if we load a value near the limit?
let result = sim.Fork ( fun s ->
s.Write ( " load_en " , 1 L ) |> ignore
s.Write ( " load_value " , 250 L ) |> ignore
s.Step ( 1 )
s.Write ( " load_en " , 0 L ) |> ignore
s.Step ( 10 )
s.ReadOrFail ( " overflow " ))
// sim is back to "before_overflow" — Fork restored automatically
// Sweep across multiple load values to find the boundary
let results = sim.Sweep (
[ 248 L ; 249 L ; 250 L ; 251 L ; 252 L ],
fun loadVal s ->
s.Write ( " load_en " , 1 L ) |> ignore
s.Write ( " load_value " , loadVal ) |> ignore
s.Step ( 1 )
s.Write ( " load_en " , 0 L ) |> ignore
s.Step ( 10 )
s.ReadOrFail ( " overflow " ))
for ( loadVal , overflow ) in results do
printfn " load= %d -> overflow= %d " loadVal overflow
}
Analyzing waveforms
test " verify timing with VCD analysis " {
use sim = SimFixture.create ()
// ... run stimulus ...
let vcd = VcdParser.parseAll " output/sim.vcd "
// When did the FSM first enter state 5?
let t = VcdParser.firstTimeAtValue vcd " fsm_state " 5
// How many times did overflow pulse?
let pulses = VcdParser.highPulseCount vcd " counter.overflow "
// What states did the FSM visit?
let states = VcdParser.uniqueValues vcd " fsm_state "
}
Test organization
Verifrog provides hardware-domain test categories so you can run the right tests at the right time:
verifrog test --category Smoke # Quick sanity — design is alive (seconds)
verifrog test --category Unit # Focused signal/block tests
verifrog test --category Integration # Multi-block data flow
verifrog test --category Parametric # Sweeps and value ranges
verifrog test # Everything
Categories are lightweight testList wrappers — just group your tests:
open Verifrog. Runner . Category
let tests = testList " MySoC " [
smoke [
test " comes out of reset " { ... }
]
unit [
test " counter increments " { ... }
]
golden [
test " matches reference output " { ... }
]
]
Also available: stress (long-running), golden (reference outputs), regression (bug-fix coverage).
Your Test Project (Expecto)
|
v
Verifrog.Runner — SimFixture, Iverilog backend, Expect helpers
|
v
Verifrog.Sim — Sim type, Memory/Register accessors, TOML config
|
v
libverifrog_sim — Generic Verilator C++ wrapper (built per-design)
|
v
Verilator — Your compiled RTL
Components
Library
What it does
Verifrog.Sim
Core simulation API: create, step, read/write signals, checkpoint/restore, force, fork/sweep, memory/register access
Verifrog.Runner
Test infrastructure: SimFixture lifecycle, Iverilog backend, Expect assertions, test categories (Smoke/Unit/Parametric/Integration/Stress/Golden/Regression)
Verifrog.Vcd
Standalone VCD waveform parser: parse files, query signals, value-at-time, transitions, timing analysis
Verifrog.Vcd.Cli
Command-line VCD analysis tool with text and JSON output
verifrog CLI
Build tool: init , build , clean , test , debug (interactive REPL), debug-server (JSON), mcp-server (MCP for Claude)
libverifrog_sim
Design-agnostic C++ shim: signal discovery, direct-pointer access, checkpoint via memcpy
VS Code Extension
Syntax highlighting for .verifrog files, signals panel, checkpoints panel, debug toolbar (experimental)
Configuration
All project configuration lives in verifrog.toml :
[ design ]
top = " my_module "
sources = [ " rtl/*.v " ]
[ test ]
output = " build "
[ memories . data_ram ]
path = " u_ram.mem "
banks = 1
depth = 1024
width = 32
[ registers ]
path = " u_regfile.regs "
width = 8
[ registers . map ]
CTRL = 0x00
STATUS = 0x01
DATA = 0x02
See the full Configuration Reference .
Sample
What it demonstrates
counter
Minimal: step, read/write, checkpoint, force, fork
alu_regfile
TOML register map, named register access, parametric sweep
sram
TOML memory regions, banked access, backdoor loading
iverilog_tb
Dual-backend: Verilator + iverilog under one dotnet test
i2c_bfm
Protocol-level BFM with auto-detection, timing-accurate I2C
Debugging
Multiple ways to debug your simulations:
Interactive REPL — the fastest path. Step the simulation, read/write signals, set checkpoints, force values, all from the command line:
verifrog debug
sim> write enable 1
sim> step 10
sim> read count
count = 10
sim> checkpoint before_overflow
si

[truncated]
