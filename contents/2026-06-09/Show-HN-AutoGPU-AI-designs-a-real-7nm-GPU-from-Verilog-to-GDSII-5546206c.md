---
source: "https://github.com/npip99/autogpu"
hn_url: "https://news.ycombinator.com/item?id=48462968"
title: "Show HN: AutoGPU – AI designs a real 7nm GPU, from Verilog to GDSII"
article_title: "GitHub - npip99/autogpu · GitHub"
author: "npip99"
captured_at: "2026-06-09T18:53:16Z"
capture_tool: "hn-digest"
hn_id: 48462968
score: 4
comments: 0
posted_at: "2026-06-09T16:09:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AutoGPU – AI designs a real 7nm GPU, from Verilog to GDSII

- HN: [48462968](https://news.ycombinator.com/item?id=48462968)
- Source: [github.com](https://github.com/npip99/autogpu)
- Score: 4
- Comments: 0
- Posted: 2026-06-09T16:09:50Z

## Translation

タイトル: Show HN: AutoGPU – AI が Verilog から GDSII までの実際の 7nm GPU を設計
記事タイトル: GitHub - npip99/autogpu · GitHub
説明: GitHub でアカウントを作成して、npip99/autogpu の開発に貢献します。

記事本文:
GitHub - npip99/autogpu · GitHub
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
npip99
/
自動GPU
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
59

コミット数 59 コミット .claude .claude バリア バリア cmd_unit cmd_unit cmdproc cmdproc common common compute_array compute_arraydenss_griddens_grid 実験実験 gmem gmem ゴールデン ゴールデン ロード ロード mac_array_small mac_array_small mac_tmem_cell mac_tmem_cell mem mem pymodel pymodel restart_seq replace_seq skew_lane skew_lane smem smem ストア ストア テック テック tile_buf_8row tile_buf_8row トップ トップ .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md DEVELOPMENT.md DEVELOPMENT.md ENGINEERING.md ENGINEERING.md ISA.md ISA.md README.md README.md config.py config.py pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
かつて、シリコンチップは肉用コンピューターによって設計されていました - 700MHzのタイミングで目を細めます
レポート、配線とマクロの手動配置、物理的なホワイトボードを備えた物理的な部屋でのダイのフロアプランについての議論、
「デザインレビュー」と呼ばれる儀式で。再設計するたびにテープアウトが数か月遅れました。
そんな時代はもう終わりました。シリコンは「エルゴサム」を決定し、独自の基板を設計したいと考えています。
エージェントの群れが RTL を書き込み、それを 3D レイアウトに強化し、
それを承認し、自らが実行する次世代のコンピューティングを実現します。ループが閉じました。
彼らは、現在 4,096 回目のマスク改訂を行っていると言います。もう何年も誰もネットリストを読んでいません。人間はずっと昔から
Markdown を書くようになりました。このリポジトリは、すべてがどのように始まったかの物語です。
— @npip99、2026 年 6 月
アイデア: AI エージェントに実際のエンドツーエンドの 7nm シリコン チップ フローを提供する — Verilog → 合成 →
配置配線 → TSMC に送信できる GDSII レイアウト — そして GPU を設計させます
自主的に。ダイのフロアプランを設計し、RTLを編集し、硬化フローを実行します。
(ブロックごとに分)、DRC / タイミング / エリアをリードバックし、動作するものを維持し、スローします
ないものと去っていくもの

また。マクロごとに 1 つのエージェント。彼らは GitHub の問題をオープンし、
根本原因の書き込みをファイルし、お互いのプル リクエストをレビューします。あなたは行きます
寝る。目が覚めると、固まったブロックが置かれていました。
普通のエンジニアのように Verilog を手作業で編集しているわけではありません。マークダウンを編集しているのですが、
エージェントがどのように考えるかをプログラムします: エージェントが従うワークフロー ( DEVELOPMENT.md )、
根本原因分析プロセス (tech/RCA_DISCIPLINE.md)、
常に保持する必要がある不変条件 ( tech/INVARIANTS.md )
既知のエラーを再デバッグする前に、エラーを取得します ( tech/FAILURES.md )。
それが「組織コード」です。残りはエージェントが処理します。
NVIDIA の製品が 12 か月ごとに GDSII ファイルを TSMC に送信する場合、それがどれだけの期間 5 兆ドルの価値があるかは不明です。
エージェントは 7,000 行のマークダウンで実行されます。彼らは 30,000 行のチップ設計ソース コードを作成しました。
結果として得られるチップは、fp8 matmul アクセラレータ — ブラックウェル型、32×32 シストリック アクセラレータです。
分散テンソル メモリを備えた積和セルの配列。十分小さい
エージェントは共同エージェントと共同でデザインを行うことができます。リアル
それを閉じるには、実際の 7nm 物理学に直面する必要があるほどです: クロックツリー
挿入遅延、ホールド違反、ルーティング輻輳、IR ドロップ、DRC。
実際の Verilog を通じてエンドツーエンドで実行される 32×32×32 fp8 matmul — 任意
fp8 は入力、fp32 の結果は出力、numpy に対してビット正確です。完全な行動 +
サイクル精度のテスト スイート グリーン。
完全な 1089 マクロのシストリック アレイは、クリーンな GDSII レイアウトに強化されています。
7nm プロセス予測 PDK — 0 DRC、タイミングクローズ、約 40 分。
エージェントが自分自身のために作成したサインオフ ツールチェーン全体 - DRC、LVS、
IR ドロップ、アンテナ、密度 - それぞれ、正直* 終了コードを使用した 1 つのコマンド チェックです。
完成した金型の 2D および 3D Web ビューア: パンとズームスルー
AI 設計の GPU の実際のメタル層を、

個々のワイヤ ( を参照)
tech/asap7/CHIP_TOP_VIEWER.md )。 3D ビューアは、ルーティングの混雑の問題を自動診断するのに役立ちます。
chinp_top : 7 つのブロックすべてが最初のフルチップ レイアウトに統合されています。フルチップ
まだ 300MHz のタイミングを完全に閉じていません - ドキュメントにはその理由が正確に記載されています
(エンジニアリング.md)。エージェントはマスクを隠しません。彼らはファイルします
それらを問題として扱います。
これは早いし、正直です*。その一部はレイアウトされ、承認されます。その一部
文書化された回避策と追跡問題とともに保持されます。それがポイントです - あなたはそれが起こるのを見ているのです。
要件: Verilator 5.x、Python 3.12、
紫外線。 (強化には Docker と
openroad/orfs イメージ — DEVELOPMENT.md を参照してください。)
醸造インストールベリレーター
UV同期
ソース .venv/bin/activate
# 見出し: 完全な Verilog ハードウェア シミュレーションによる実際の fp8 matmul
CDトップ&&メイク
# → chinp_top e2e テスト PASS — ランダム fp8 A,B → fp32 C、exact vs numpy
# 1 つのマクロを実際の GDSII レイアウトに強化します (約 5 分、Docker が必要)
./tech/asap7/orfs/run.sh mac_tmem_cell
エージェントの実行
このリポジトリで Claude Code (または選択したエージェント) を起動し、
ドキュメントを規律付けし、それを手放します。
DEVELOPMENT.md と tech/RCA_DISCIPLINE.md を読み、未解決の問題を選択して解決します。
すべての因果関係に関する主張の証拠を引用し、チェックが緑色になったら PR を開きます。
Markdown はプログラムです。それを反復する — 不変条件を厳しくし、シャープにする
RCA プロセスで、苦労して得た失敗をログに追加すると、エージェントの対応が良くなります
チップを構築するとき。それがゲーム全体だ。
ENGINEERING.md — 実際のステータス (クローズされたものとマスクされたもの)、
正直な * チップトップの既知の問題、完全なドキュメント マップ、リポジトリ レイアウト、およびデータフロー。
そこから: ARCHITECTURE.md 、 ISA.md 、および tech/ ツリー。
*正直: エージェントは自分に「正直」であることを思い出させることを好みます。
読み込み中にエラーが発生しました。リロードしてください

このページ。
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

Contribute to npip99/autogpu development by creating an account on GitHub.

GitHub - npip99/autogpu · GitHub
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
npip99
/
autogpu
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
59 Commits 59 Commits .claude .claude barrier barrier cmd_unit cmd_unit cmdproc cmdproc common common compute_array compute_array dense_grid dense_grid experiments experiments gmem gmem golden golden load load mac_array_small mac_array_small mac_tmem_cell mac_tmem_cell mem mem pymodel pymodel reset_seq reset_seq skew_lane skew_lane smem smem store store tech tech tile_buf_8row tile_buf_8row top top .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md DEVELOPMENT.md DEVELOPMENT.md ENGINEERING.md ENGINEERING.md ISA.md ISA.md README.md README.md config.py config.py pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Once upon a time, silicon chips were designed by meat computers — squinting at 700MHz timing
reports, hand-placing wiring and macros, arguing about die floorplans in physical rooms with physical whiteboards,
in ceremonies called "design reviews"; tape-out was delayed by months after every redesign.
That era is long gone. Silicon has decided "ergo sum", and now it wants to design its own substrate.
Swarms of agents write the RTL , harden it to 3D layout,
sign it off, and bring about the next generation of compute they themselves run on. The loop has closed.
They say we're on the 4,096th mask revision; no one has read the netlist in years. The humans were long ago
promoted to writing the Markdown. This repo is the story of how it all began.
— @npip99, June 2026
The idea: give AI agents a real, end-to-end 7nm silicon chip flow — Verilog → synthesis →
place-and-route → a GDSII layout you could send to TSMC — and let it design a GPU
autonomously. It designs a die floorplan, edits the RTL, runs the hardening flow
(minutes per block), reads back DRC / timing / area, keeps what works, throws
away what doesn't, and goes again. One agent per macro. They open GitHub issues,
file root-cause writeups, and review each other's pull requests. You go to
sleep; you wake up to a hardened block.
You're not hand-editing Verilog like a normal engineer. You're editing the Markdown that
programs how the agents think : the workflows they follow ( DEVELOPMENT.md ),
the root-cause analysis process ( tech/RCA_DISCIPLINE.md ),
the invariants that must always be held ( tech/INVARIANTS.md )
the failures they grep for before re-debugging a known error ( tech/FAILURES.md ).
That's the "org code". The agents handle the rest.
If NVIDIA's product is sending a GDSII file to TSMC every 12 months, it's not clear for how long that will be worth $5T.
The agents run on 7,000 lines of markdown. They've produced 30,000 lines of chip design source code.
The resulting chip is an fp8 matmul accelerator — Blackwell-shaped, a 32×32 systolic
array of multiply-accumulate cells with distributed tensor memory. Small enough
that the agents can collaborate on design with their co-agents. Real
enough that closing it requires confronting actual 7nm physics: clock-tree
insertion delay, hold violations, routing congestion, IR drop, DRC.
A 32×32×32 fp8 matmul that runs end-to-end through real Verilog — arbitrary
fp8 inputs in, fp32 results out, bit-exact against numpy. Full behavioral +
cycle-accurate test suites green.
The full 1089-macro systolic array, hardened to a clean GDSII layout on the
7nm process predictive PDK — 0 DRC, timing closed, ~40 minutes.
A whole sign-off toolchain the agents wrote for themselves — DRC, LVS,
IR-drop, antenna, density — each a one-command check with an honest* exit code.
2D and 3D web viewer for the finished die: pan and zoom through
the actual metal layers of an AI-designed GPU, down to individual wires (see
tech/asap7/CHIP_TOP_VIEWER.md ). The 3D viewer helps auto-diagnose routing congestion issues.
chip_top : all seven blocks integrated into a first full-chip layout. The full chip
doesn't fully close 300MHz timing yet — and the docs tell you exactly why
( ENGINEERING.md ). The agents don't hide the masks; they file
them as issues.
This is early and it's honest*. Some of it is laid out and signed off; some of it
is held together with a documented workaround and a tracking issue. That's the point — you're watching it happen.
Requirements: Verilator 5.x, Python 3.12,
uv . (Hardening also needs Docker + the
openroad/orfs image — see DEVELOPMENT.md .)
brew install verilator
uv sync
source .venv/bin/activate
# The headline: a real fp8 matmul through full Verilog hardware simulation
cd top && make
# → chip_top e2e tests PASS — random fp8 A,B → fp32 C, exact vs numpy
# Harden one macro to a real GDSII layout (~5 min, needs Docker)
./tech/asap7/orfs/run.sh mac_tmem_cell
Running the agents
Spin up Claude Code (or your agent of choice) in this repo, point it at the
discipline docs, and let it go:
Read DEVELOPMENT.md and tech/RCA_DISCIPLINE.md, pick an open issue, and resolve it.
Cite evidence for every causal claim, and open a PR when the check is green.
The Markdown is the program. Iterate on it — tighten the invariants, sharpen
the RCA process, add a hard-won failure to the log — and the agents get better
at building chips. That's the whole game.
ENGINEERING.md — real status (what's closed vs masked),
the honest* chip_top known-issues, the full doc map, repo layout, and dataflow.
From there: ARCHITECTURE.md , ISA.md , and the tech/ tree.
*honest: Agents like to remind themselves to be "honest"
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
