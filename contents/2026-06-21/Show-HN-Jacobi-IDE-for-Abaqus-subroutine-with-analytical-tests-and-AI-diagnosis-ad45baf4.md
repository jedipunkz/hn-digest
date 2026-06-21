---
source: "https://jacobee.netlify.app/"
hn_url: "https://news.ycombinator.com/item?id=48620188"
title: "Show HN: Jacobi–IDE for Abaqus subroutine with analytical tests and AI diagnosis"
article_title: "Jacobi — Abaqus Subroutine IDE"
author: "white_tiger"
captured_at: "2026-06-21T16:31:38Z"
capture_tool: "hn-digest"
hn_id: 48620188
score: 2
comments: 3
posted_at: "2026-06-21T16:20:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Jacobi–IDE for Abaqus subroutine with analytical tests and AI diagnosis

- HN: [48620188](https://news.ycombinator.com/item?id=48620188)
- Source: [jacobee.netlify.app](https://jacobee.netlify.app/)
- Score: 2
- Comments: 3
- Posted: 2026-06-21T16:20:46Z

## Translation

タイトル: HN を表示: 分析テストと AI 診断を備えた Abaqus サブルーチンの Jacobi–IDE
記事のタイトル: Jacobi — Abaqus サブルーチン IDE
説明: Fortran サブルーチンを作成し、分析ソリューションに対してテストし、AI 診断を取得します。 Abaqus エンジニア向けに構築された IDE。
HN テキスト: 私は計算力学の大学院生として、Abaqus UMAT サブルーチンを作成しています。これらは、Abaqus Fortran サブルーチン (UMAT ～ 機械的動作、UMATHT ～ 熱/拡散など) を使用した複雑なマルチフィジックス シミュレーション モデルで、高温または製造プロセスの下でさまざまな材料システムがどのように故障するかをシミュレートします。ただし、プロセス全体は非常に困難で、実際に時間の 80 ～ 90% が、紙の上ですでにわかっている物理を Abaqus CAE でどのようにシミュレートするか、または使用する必要がある正しいサブルーチンと変数は何かということに費やされています。 .sta .msg ファイルを深く掘り下げてシミュレーションが失敗した理由を解明し、デバッグするのにも膨大な時間がかかります。ソフトウェアまたは現在のツール (計算物理学や力学ではなく、ソフトウェア エンジニアリング用に構築された IDE ~ VS Code など) は、シミュレーションでセグメンテーション エラーが発生する理由を教えてくれません。また、ダメージ変数が 1 (完全に破損している) の場合にゼロ除算エラーが発生することを警告しません。サイレントミスやエラーは、シミュレーションを通じてサイレントに伝播し、物理学は完全に間違っています (さらに悪いことに、物理学は理にかなっていますが、それでも間違っています)。ドキュメントは簡潔かつ複雑で、10 年間変更されておらず、基本的な線形弾性の例のみが提供されています。計算エンジニアが自分の仕事に取り組む方法全体は、ある種の秘密感を持っています。 NDA やさまざまなポリシーの背後にコードが隠されています。ただし、これは国家安全保障の観点からは理にかなっています。それは、この分野のイノベーションを遅らせ、抑圧します。コードは共有されません

ソフトウェア エンジニアと同様に、リポジトリはプライベートであり、.odb ファイルは隠されており、複雑なマルチフィジックス シミュレーションに関するチュートリアルのリソースが少なくなります。これが、Fei Fei Li 博士が世界モデルに関する最近の投稿で、世界モデルの効果的なトレーニングに使用できる利用可能性の低いデータを利用して、世界モデルの基軸として世界モデル分類法の「シミュレート」部分 (他の 2 つはプランナーとレンダラー) を特定した理由です (https://x.com/drfeifei/status/2062247238143996275 を参照してください)。エンジニアがこれらの問題の一部を「スキルの問題」として分類することがありますが、これは間違った考え方であると私は考えています。しかし、以前にソフトウェアや Web アプリを構築した経験があるので、ソフトウェア エコシステム全体にとってツールとコミュニティがいかに重要であるかを知っています (Github 上のオープンソース プロジェクト、簡単に pip インストールできるモジュール/パッケージ、Karpathy などのレジェンドによるチュートリアルを考えてください)。 Jacobi Physics IDE を使用して、物理シミュレーション用のツールとエコシステムを構築しています。 Abaqus についての現在の知識があるため、そこから始めます。しかし最終的には、ユーザーからのフィードバックや関与を受けて、これは世の中にある複数の物理ソルバー (COMSOL、NASTRAN、LS-DYNA...コメントで皆さんの意見をもっと聞きたいです) をカバーするように拡張され、ユーザーや研究者が独自のシミュレーション用のサブルーチンやユーザー定義の材料挙動を作成できるようになります。 https://jacobee.netlify.app/ から IDE をダウンロードしてください。ぜひご意見をお待ちしています。
ありがとう。

記事本文:
ヤコビ
プラットフォーム
ワークフロー
特長
テスト
ドキュメント
v0.1 ベータ版
ダウンロード↓
物理シミュレーションサブルーチン
ヤコビ
物理シミュレーション サブルーチン用の IDE。 Fortran、C++、または Python を作成し、閉じた形式の分析ソリューションに対してテストし、AI 診断を取得します。今日の Abaqus 用に構築され、完全なシミュレーション エコシステム用に設計されています。
Windows 用のダウンロード
gfortran が必要です
Windows 10 / 11
オープンソース
✓ AIの完成
✓ テストマーケットプレイス
✓ AI診断
✓ ホバードキュメント
✓ 定義に移動
✓ 12 個のテンプレート
✓ 20のテーマ
✓ 静的解析
コムソル
ナストラン
jacobi — テストランナー
$ jacobi 実行 elastic_j2.f
gfortran -O2 -Wall でコンパイルしています...
12 個の物理テストを実行する
✓ 一軸 C₁₁₁₁ 誤差: 0.021%
✓ B 横方向 C₁₁₂₂ 誤差: 0.008%
✓ C せん断力誤差: 0.000%
✓ D バルクエラー: 0.031%
✓ E J2 中立性 20/20 ステップ
✓ F 収量開始ステップ 8 ✓
✓ G 弾性アンロード C₁₁₁₁ OK
✗ K 非圧縮率 1.2e‑6 > tol
✓ L 利回り誤差: 0.04%
✓ H ゼロ増分 |dσ|=1e‑14
✗ I DDSDDE FD 3.4% > 1% tol
✓ J DDSDDE 対称性 |C-Cᵀ|=8e‑15
10 回合格 · 2 回不合格 — クロードに送信...
クロード: 塑性流動には体積測定があります。
コンポーネント。 NHAT を確認してください...
$
30以上
UMAT、VUMAT、UEL、UHARD、および CREEP サブルーチン タイプにわたる物理テスト
8
マーケットプレイスにバンドルされたコミュニティ テスト - 独自のインストール、実行、作成が可能
<5秒
フルサイクル — ソースからテスト結果、AI 診断まで
20
Tokyo NightからRosé Pine、GitHub Darkまでのエディターのカラーテーマ
なぜヤコビなのか
サブルーチン開発
推測なしで
構成的サブルーチンを作成するということは、リターン マッピング アルゴリズムを実装し、一貫した接線を組み立て、ソルバーが明示的にチェックしない物理的制約を満たすことを意味します。シミュレーションが異なる場合、エラー メッセージは何も表示しません。
Jacobi がサブを実行します

15 の閉じた形式の分析ソリューションに対するルーチンを実行し、どのソリューションがどの程度失敗したかを正確に示します。次に、クロードが完全な数値コンテキストを使用して、物理的な原因とコード内のどこを調べるべきかを説明します。
汎用エディタではありません。それは、ソルバー間、言語間で、サブルーチンの物理的な動作をより迅速に修正することです。
根本からソルバーに依存しない
Jacobi は、単一のソルバーではなく、物理シミュレーション エコシステムをターゲットにしています。現在のリリースには、Abaqus との深い統合が含まれています。アーキテクチャ (サブルーチンの読み込み、テスト ハーネス、AI 診断、言語サポート) は拡張するように設計されています。 COMSOL、Nastran、LS-DYNA の統合は開発中です。
長期的な目標は、既存のソルバー ライセンスを使用して Jacobi 内からあらゆるシミュレーションを実行し、サブルーチン コードを変更せずにソルバーを切り替えることです。
発生源から診断までの 3 つのステップ
Monaco エディターで、任意の .f 、 .f90 、 .cpp 、または .c ファイルを開きます。 Fortran および C/C++ の構文ハイライト、静的分析、AI インライン補完、サブルーチン検出が自動的に実行されます。プロジェクトのセットアップは必要ありません。
「テストの実行」をクリックします。 Jacobi は、単一要素ドライバーをコンパイルし、サブルーチンをリンクし、15 のシナリオにわたって規定のひずみ増分を適用し、各結果を閉じた形式の分析結果と比較します。
gfortran ドライバー
03
AI診断を受ける
15 個の数値結果すべて (応力、ひずみ、エラー、合否) がクロードに送信されます。あらゆる失敗について物理学を優先した説明が得られ、問題となる特定の変数と方程式に加えて教科書の参照も含まれます。
クロードの俳句
v0.1の内容
正しいサブルーチンを出荷するために必要なものすべて
UMAT 用の 8 つのバンドルされたコミュニティ テスト。開いているファイルを参照、インストール、実行します。独自のファイルを作成し、 .jtest としてエクスポートし、チームと共有します。
任意の Abaqus にカーソルを合わせます

引数 — STRESS、DDSDDE、STATV — タイプ、寸法、I/O 方向、および物理的な説明を確認します。 Fortran 組み込み関数もインラインで文書化されています。
F12 キーを押すか、任意の変数、パラメーター、またはサブルーチンをシングルクリックして、定義されている場所にジャンプします。 Monaco の完全なナビゲーション エクスペリエンスを使用してファイル内で動作します。
UMAT、VUMAT、UEL、UHARD、CREEP、UMATHT など - それぞれに完全な署名、正しい IMPLICIT 宣言、およびインライン物理コメントが事前に入力されています。
材料パラメータをキャリブレーション パネルの実験データに適合させます。 Step Inspector は、各テストの実行後に、増分ごとの完全な状態 (応力、ひずみ、STATEV) を表示します。
Abaqus .inp ファイルを解析して、ステップ、材料、境界条件を参照します。 .msg / .dat ログをインポートして、収束履歴を視覚化します。
Fortran の仕組みを知るゴーストテキスト
入力すると、Jacobi はコードのプレフィックスとサフィックスをクロードに送信します。モデルには、サブルーチンのタイプ、固定形式か自由形式のフラグ、および周囲のコンテキストが要求されます。戻ってくるものは半透明のゴースト テキストとして表示されます。タブを押して受け入れます。
これはトークン予測のオートコンプリートではありません。モデルは、DDSDDE が対称である必要があること、NHAT が J2 に対して偏差的である必要があること、および STATEV インデックス付けがサブルーチン固有であることを理解しています。
Alt+\ を押して手動でトリガーする
キーストロークごとにゴーストテキストを表示したくないですか?オンデマンドでトリガーします。提案はエディターにインラインで表示され、ポップアップや邪魔はされません。
サブルーチンを意識したコンテキスト
Jacobi は、カーソルがどのサブルーチン (UMAT、VUMAT、USDFLD、CREEP) にあるかを検出し、Claude に通知するため、補完は各署名の規則を尊重します。
あなたの鍵を Anthropic に直接送信してください
Anthropic キーを 1 回入力してください。すべての呼び出しはデスクトップ アプリから Anthropic API に直接行われます。プロキシやテレメトリはなく、リクエストあたり数セントです。
クロードはあなたのテスト結果だけでなく、テスト結果も読み取ります。

頌歌
テスト スイートが完了すると、Jacobi は完全な数値コンテキスト (すべての応力成分、誤差の大きさ、合否フラグ、派生量) を組み立て、材料パラメータとともに Claude Haiku に送信します。
返されるのは、物理優先の診断です。 「ライン 42 を確認してください」ではなく、「プラスチック フロー法線には、J2 には存在しないはずの体積成分が含まれています。その理由は次のとおりです。」
連続力学の接地
リファレンス リターン マッピング アルゴリズム、フォークト記法、アイソコリック フロー、および一貫したタンジェント導出を診断します。一般的なデバッグ アドバイスではありません。
教科書の引用も含まれています
すべての診断には、デ・ソウザ・ネト、シモ＆ヒューズ、ベリチコなどの特定の章が引用されているため、必要に応じてソースにアクセスできます。
API キー、データ
Anthropic キーをツールバーに一度追加します。呼び出しはアプリから Anthropic API に直接行われます。プロキシやサーバーは必要なく、実行あたり数セントです。
各テストはサブルーチンを直接呼び出し、規定のひずみ増分を適用し、結果を閉じた形式の分析結果と比較します。許容値は標準の FE ソルバー収束基準と一致します。 15 件の結果すべてが診断のためにクロードに送信されます。
Fortran を意図された方法で記述します
Monaco エディター — VS Code と同じエンジン — 完全な Fortran 構文の強調表示、固定形式および自由形式のルーラーのサポート、およびテスト結果パネルからのエラーへのジャンプ機能を備えています。
静的解析では、コンパイル前に、初期化されていない変数と未定義のシンボルをインラインでフラグ付けします。 Tokyo Night、Rosé Pine、Kanada、GitHub Dark を含む 20 の組み込みテーマ。
1 C J2 等方塑性 UMAT
2 サブルーチン UMAT ( STRESS 、 STATEV 、 DDSDDE 、...)
3 暗黙的な REAL * 8 ( A ～ H 、 O ～ Z )
4 EMOD = プロップ ( 1 )
5 ANU = プロップ ( 2 )
6 シールド = プロップ ( 3 )
7 GMOD = EMOD

/ ( 2.D0 *( 1.D0 + ANU ))
8 SMISES = SQRT ( 1.5D0 *( S11 ** 2 + S22 ** 2 + S33 ** 2
9 & + 2.D0 *( S12 ** 2 + S13 ** 2 + S23 ** 2 )))
10 IF (スミス .GT. シールド) THEN
11 DLAM = ( SMISES - SYIELD ) / ( 3.D0 * GMOD )
12 DO I = 1、3
13 ストレス ( I ) = ストレス ( I ) - 2.D0 * GMOD * DLAM * NHAT ( I )
14 エンドドゥ
15 エンドIF
σ vm = 0 ピーク応力
FEMの可視化
数字だけじゃない、目に見えるデフォルメ
メッシュは荷重方向に伸び、横方向には速度 ν で収縮します。 Jet カラーマップは、弾性および塑性領域を通じて σ₁₁ を追跡します。
体積変化のない純粋なせん断。 G = E/2(1+ν) の関係を幾何学的に検証します。
すべての方向に均等なひずみを与えます。フォン・ミーゼス応力はゼロのまま – J2 中立性テストを視覚化。
ステップ スライダーをドラッグして、読み込み履歴の任意のポイントを検査します。ゴーストの輪郭は元の構成を示しています。
単一のインストーラー。設定はありません。テスト実行には PATH に gfortran が必要です。 Anthropic API キーをツールバーに追加して、AI 診断とインライン補完を有効にします。
Windows 用のダウンロード
v0.1.0 ベータ版 · Windows 10/11 x64
Windows で SmartScreen 警告が表示される場合があります。[詳細情報] → [とにかく実行] をクリックします。
最初に gfortran をインストールします: winlibs.com (推奨) または MinGW-w64 。

## Original Extract

Write Fortran subroutines, test against analytical solutions, get an AI diagnosis. The IDE built for Abaqus engineers.

I write Abaqus UMAT subroutines as a graduate student in computational mechanics. These are complex multi-physics simulation models using Abaqus Fortran subroutines (UMAT ~ mechanical behavior, UMATHT ~ heat/diffusion and a lot more) that simulate how different material systems fail under high temperature or manufacturing processes. However, the entire process has been quite challenging, 80-90% of time is actually spent on how make Abaqus CAE simulate a physics you already know on paper or what are the correct subroutines and variables you should use. Huge time is also spent on diving deep into the .sta .msg files to figure out and debug why a simulation failed. The software or current tooling (IDE ~ which is VS Code etc built for software engineering, not computational physics or mechanics) doesn't tell you why the simulation run into a segmentation error, or warn you when your damage variable being 1 (fully damaged) will lead to a zero division error. Silent mistakes or errors propagate through your simulation silently, and the physics is completely wrong (or worse, physics make sense, but wrong nonetheless). The documentation is terse and complex, haven't changed for a decade, offering basic linear elastic examples only. The entire way computational engineers approach their work is with a touch/sense of secrecy. Codes are hidden behind NDAs and various policies. Although, this makes sense from a national security standpoint. It slows down and stifle innovation in the field. Codes are not shared like the way software engineers do, repositories are private, .odb files are hidden, Less resources for tutorials on complex multi-physics simulations. This is why Dr. Fei Fei Li in her recent post on world models, identifying the "simulate" part of the world model taxonomy (the other two is a planner and a renderer) as the linchpin of a world model (read here https://x.com/drfeifei/status/2062247238143996275 ) with orders of magnitude of less available data available to effectively train world models. On occasions, engineers classify some of these problems as a "skill issue", which I believe is a wrong mindset to have. However, having prior experience building software and webapps, I know how important the tooling and community is to the entire software ecosystem (think open-source projects on Github, Modules/Packages you can easily pip install, and tutorials from legends like Karpathy and a lot more). With the Jacobi Physics IDE, I am building that tooling and ecosystem for physics simulation. Due to my current knowledge of Abaqus, that is what I start with. But eventually, following feedback and engagement from users, this will expand to cover multiple physics solvers out there (COMSOL, NASTRAN, LS-DYNA....I want to hear more from you guys in the comment) that enable users and researchers to write subroutines and user-defined material behavior for their own unique simulations. I will love to hear from you guys, download the IDE at https://jacobee.netlify.app/ .
Thanks.

Jacobi
Platform
Workflow
Features
Tests
Docs
v0.1 Beta
Download ↓
Physics Simulation Subroutines
Jacobi
An IDE for physics simulation subroutines. Write Fortran, C++, or Python , test against closed-form analytical solutions, get an AI diagnosis — built for Abaqus today, designed for the full simulation ecosystem.
Download for Windows
gfortran required
Windows 10 / 11
Open source
✓ AI Completion
✓ Test Marketplace
✓ AI Diagnosis
✓ Hover Docs
✓ Go-to-Definition
✓ 12 Templates
✓ 20 Themes
✓ Static Analysis
COMSOL
Nastran
jacobi — test runner
$ jacobi run elastic_j2.f
Compiling with gfortran -O2 -Wall...
Running 12 physics tests
✓ A Uniaxial C₁₁₁₁ err: 0.021%
✓ B Lateral C₁₁₂₂ err: 0.008%
✓ C Shear G err: 0.000%
✓ D Bulk K err: 0.031%
✓ E J2 neutrality 20/20 steps
✓ F Yield onset step 8 ✓
✓ G Elastic unload C₁₁₁₁ ok
✗ K Incompressibility 1.2e‑6 > tol
✓ L Yield return err: 0.04%
✓ H Zero increment |dσ|=1e‑14
✗ I DDSDDE FD 3.4% > 1% tol
✓ J DDSDDE symmetry |C-Cᵀ|=8e‑15
10 passed · 2 failed — sending to Claude...
Claude: Plastic flow has volumetric
component. Check NHAT...
$
30 +
Physics tests across UMAT, VUMAT, UEL, UHARD, and CREEP subroutine types
8
Bundled community tests in the marketplace — install, run, or author your own
<5 s
Full cycle — source to test results to AI diagnosis
20
Editor color themes from Tokyo Night to Rosé Pine to GitHub Dark
Why Jacobi
Subroutine development
without the guesswork
Writing a constitutive subroutine means implementing a return-mapping algorithm , assembling a consistent tangent, and satisfying physical constraints the solver never explicitly checks. When a simulation diverges, the error message tells you nothing.
Jacobi runs your subroutine against fifteen closed-form analytical solutions and tells you exactly which one failed and by how much. Then Claude — with full numerical context — explains the physical cause and where in your code to look.
It is not a general-purpose editor. It does one thing: get your subroutine to physically correct behavior faster — across solvers, across languages.
Solver-agnostic from the ground up
Jacobi targets the physics simulation ecosystem, not a single solver. The current release ships deep Abaqus integration. The architecture — subroutine loading, test harness, AI diagnosis, language support — is designed to extend. COMSOL, Nastran, and LS-DYNA integrations are in development.
The long-range goal: run any simulation from within Jacobi using your existing solver license, switching solvers without modifying your subroutine code.
Three steps from source to diagnosis
Open any .f , .f90 , .cpp , or .c file in the Monaco editor. Fortran and C/C++ syntax highlighting, static analysis, AI inline completion, and subroutine detection run automatically. No project setup required.
Click Run Tests. Jacobi compiles a single-element driver, links your subroutine, applies prescribed strain increments across 15 scenarios, and compares each result to its closed-form analytical answer.
gfortran driver
03
Get AI diagnosis
All fifteen numerical results — stresses, strains, errors, pass/fail — are sent to Claude. You get a physics-first explanation of every failure, with the specific variable and equation at fault, plus textbook references.
Claude Haiku
What's in v0.1
Everything you need to ship a correct subroutine
8 bundled community tests for UMAT. Browse, install, run against your open file. Author your own, export as .jtest , share with your team.
Hover any Abaqus argument — STRESS, DDSDDE, STATEV — to see its type, dimensions, I/O direction, and a physics description. Fortran intrinsics documented inline too.
F12 or single-click any variable, parameter, or subroutine to jump to where it's defined. Works within the file with Monaco's full navigation experience.
UMAT, VUMAT, UEL, UHARD, CREEP, UMATHT and more — each pre-filled with the full signature, correct IMPLICIT declarations, and inline physics comments.
Fit material parameters to experimental data in the Calibration Panel. Step Inspector shows the full per-increment state — stress, strain, STATEV — after every test run.
Parse Abaqus .inp files to browse steps, materials, and boundary conditions. Import .msg / .dat logs to visualize convergence history.
Ghost text that knows Fortran mechanics
As you type, Jacobi sends the prefix and suffix of your code to Claude. The model is prompted with your subroutine type, fixed vs. free form flag, and surrounding context. What comes back appears as translucent ghost text — Tab to accept.
This isn't token-prediction autocomplete. The model understands that DDSDDE must be symmetric, that NHAT should be deviatoric for J2, and that STATEV indexing is subroutine-specific.
Alt+\ to trigger manually
Don't want ghost text on every keystroke? Trigger on demand. The suggestion appears inline in the editor, no popup, no distraction.
Subroutine-aware context
Jacobi detects which subroutine the cursor is in (UMAT, VUMAT, USDFLD, CREEP) and tells Claude, so completions respect each signature's conventions.
Your key, direct to Anthropic
Enter your Anthropic key once. Every call goes directly from the desktop app to the Anthropic API — no proxy, no telemetry, fractions of a cent per request.
Claude reads your test results, not just your code
After the test suite finishes, Jacobi assembles the full numerical context — every stress component, error magnitude, pass/fail flag, and derived quantity — and sends it to Claude Haiku along with your material parameters.
What comes back is a physics-first diagnosis . Not “check line 42” — but “your plastic flow normal has a volumetric component that shouldn’t be there for J2, here’s why, here’s where.”
Continuum mechanics grounding
Diagnoses reference return-mapping algorithms, Voigt notation, isochoric flow, and consistent tangent derivation — not generic debugging advice.
Textbook citations included
Every diagnosis cites the specific chapter: de Souza Neto, Simo & Hughes, Belytschko — so you can go to the source if needed.
Your API key, your data
Add your Anthropic key once in the toolbar. Calls go directly from the app to the Anthropic API — no proxy, no server, fractions of a cent per run.
Each test calls your subroutine directly, applies a prescribed strain increment, and compares the result to the closed-form analytical answer. Tolerances match standard FE solver convergence criteria. All 15 results are fed to Claude for diagnosis.
Write Fortran the way it was meant to be written
Monaco editor — the same engine as VS Code — with full Fortran syntax highlighting, fixed-form and free-form ruler support, and jump-to-error navigation from the test results panel.
Static analysis flags uninitialized variables and undefined symbols inline, before compilation. 20 built-in themes including Tokyo Night, Rosé Pine, Kanagawa, and GitHub Dark.
1 C J2 Isotropic Plasticity UMAT
2 SUBROUTINE UMAT ( STRESS , STATEV , DDSDDE ,...)
3 IMPLICIT REAL * 8 ( A - H , O - Z )
4 EMOD = PROPS ( 1 )
5 ANU = PROPS ( 2 )
6 SYIELD = PROPS ( 3 )
7 GMOD = EMOD / ( 2.D0 *( 1.D0 + ANU ))
8 SMISES = SQRT ( 1.5D0 *( S11 ** 2 + S22 ** 2 + S33 ** 2
9 & + 2.D0 *( S12 ** 2 + S13 ** 2 + S23 ** 2 )))
10 IF ( SMISES .GT. SYIELD ) THEN
11 DLAM = ( SMISES - SYIELD ) / ( 3.D0 * GMOD )
12 DO I = 1 , 3
13 STRESS ( I ) = STRESS ( I ) - 2.D0 * GMOD * DLAM * NHAT ( I )
14 END DO
15 END IF
σ vm = 0 peak stress
FEM Visualization
Deformation you can see, not just numbers
Mesh stretches in load direction, contracts laterally at rate ν. Jet colormap tracks σ₁₁ through elastic and plastic regimes.
Pure shear without volumetric change. Verifies the G = E/2(1+ν) relationship geometrically.
Equal strain in all directions. Von Mises stress stays zero — the J2 neutrality test, visualized.
Drag the step slider to inspect any point in the loading history. Ghost outline marks the original configuration.
A single installer. No configuration. Requires gfortran on your PATH for test execution. Add your Anthropic API key in the toolbar to enable AI diagnosis and inline completion.
Download for Windows
v0.1.0 beta · Windows 10/11 x64
Windows may show a SmartScreen warning — click More info → Run anyway .
Install gfortran first: winlibs.com (recommended) or MinGW-w64 .
