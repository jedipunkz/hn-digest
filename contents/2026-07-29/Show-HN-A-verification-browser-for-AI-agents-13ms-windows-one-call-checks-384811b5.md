---
source: "https://github.com/hongnoul/hwatu"
hn_url: "https://news.ycombinator.com/item?id=49097727"
title: "Show HN: A verification browser for AI agents – 13ms windows, one-call checks"
article_title: "GitHub - hongnoul/hwatu: Fast, interruptible verification browser for AI coding agents: 35 ms checks, pixel diffs, live human hand-off · GitHub"
author: "hongnoul_"
captured_at: "2026-07-29T15:07:02Z"
capture_tool: "hn-digest"
hn_id: 49097727
score: 6
comments: 1
posted_at: "2026-07-29T14:05:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A verification browser for AI agents – 13ms windows, one-call checks

- HN: [49097727](https://news.ycombinator.com/item?id=49097727)
- Source: [github.com](https://github.com/hongnoul/hwatu)
- Score: 6
- Comments: 1
- Posted: 2026-07-29T14:05:30Z

## Translation

タイトル: Show HN: AI エージェント用の検証ブラウザ – 13 ミリ秒のウィンドウ、ワンコール チェック
記事タイトル: GitHub - hongnoul/hwatu: AI コーディング エージェント向けの高速で中断可能な検証ブラウザ: 35 ミリ秒のチェック、ピクセル差分、ライブ人間ハンドオフ · GitHub
説明: AI コーディング エージェント用の高速で中断可能な検証ブラウザ: 35 ミリ秒のチェック、ピクセル差分、ライブ ヒューマン ハンドオフ - hongnoul/hwatu

記事本文:
GitHub - hongnoul/hwatu: AI コーディング エージェント用の高速で中断可能な検証ブラウザ: 35 ミリ秒のチェック、ピクセル差分、ライブ ヒューマン ハンドオフ · GitHub
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
ホンヌール
/
ファトゥ
公共
通知
にサインインする必要があります

通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
221 コミット 221 コミット .astrophile .astrophile .github .github crates crates docs docs 例 例 パッケージ化 パッケージング スクリプト スクリプト .gitignore .gitignore COTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md VISION.md VISION.mddeny.tomldeny.tomlglama.jsonglama.jsonllms.txtllms.txtserver.jsonserver.jsonすべてのファイルを表示リポジトリ ファイルのナビゲーション
エージェントに実際の目を与えることで、エージェントのハーネス ループを瞬時に高速化します。
エージェントが「ピクセル完璧」と主張するのはやめてください。 97.49%を証明してください。
ページ チェックごとに 5 つのツール コールの支払いを停止します。 hwatu チェックは 1 回の呼び出しで、最大 35 ミリ秒です (warm-server Playwright の最大 9 倍を上回ります)。
ブラウザウィンドウがあなたの集中力を奪うのをやめてください。デフォルトではヘッドレスなので、入力し続けます。
170 MB の Chromium の出荷を停止します。 1 つの静的バイナリ + ディストリビューションの webkitgtk。
ビジョン : 耐久性のある製品の原則、ネイティブ プラットフォーム戦略、群モデル
エージェント ガイド: プロトコル、プリミティブ、検証ループ
ヒューマンガイド：タイリング・WMブラウザ側
ベンチマーク: 方法論を使用して測定されたすべての数値
ロードマップ : 記録計画、優先順位、非目標
継続的な改善: アクティベーションメトリクス、フィードバックループ、毎週のペース
起動キット: 再利用可能なコピー、チャネル、測定プラン
インストール → ワークフローの検出 → エージェントの接続 → ページの確認 → 人間に引き渡す
カール -fsSL https://raw.githubusercontent.com/hongnoul/hwatu/main/scripts/install.sh |バッシュ
ファトゥセットアップ
1 つの静的バイナリとディストリビューションの webkitgtk-6.0 (インストーラー)
小切手）。アーチ上: やったー -S hwatu 。ソースから: Cargo build --release 。
インストーラーはバイナリのみをインストールします。 hwatu セットアップはサポートされているコーディングを検出します
エージェントと利用可能な情報を印刷します

構成を変更せずに接続できます。
準備ができたら、クライアントを明示的に選択します。
ファトゥドクター
hwatu setup --client claude --scope project --dry-run
hwatu setup --client claude --scope プロジェクト
ファトゥデモ
インストール: 2 つのバイナリをダウンロードし、WebKitGTK を確認します。
検出: クロード コード、カーソル、Jcode、または一般的な MCP ワークフローを検索します。
接続: Jcode のネイティブ ソケット、MCP、または CLI フォールバックを使用します。
検証: ドクターまたはデモを使用して、ヘッドレスでレンダリングされた煙テストを実行します。
ハンドオフ: 人間が必要な場合にのみ同じライブ セッションを実現します。
セットアップはプレビュー可能で、べき等であり、同じクライアントで元に戻すことができます。
スコープに --undo を加えます。プロジェクト スコープは共有可能な構成を作成します。ユーザースコープ
それを個人的なものにしておきます。クロード コードは各ユーザーにプロジェクト スコープの MCP を承認するよう要求します
次回起動時にサーバーが削除されます。
手動 MCP 設定は 1 つの移植可能なエントリのままです。
{ "mcpServers" : { "hwatu" : { "コマンド" : " hwatu " , "args" : [ " mcp " ] } } }
または、MCP を完全にスキップします。すべてのコマンドは短い CLI 呼び出しまたは 1 つです。
Unix ソケット上の改行区切りの JSON 行。
hwatu localhost:3000 # ターミナルを開くようにウィンドウを開きます
ファトゥを試してみましたか？チェックが成功した場合、インストールが失敗した場合、ワークフローが欠落している場合は、
すべての有用な信号。 2 分間の使用レポートを共有する
またはバグを報告してください。
ピクセル差分スコアリング: 一致パーセント + 差分領域 + ヒートマップ ( diff )
数値としてのアニメーション: 期間、イージング、速度 (モーション)
決定的アニメーション フレーム: 時間 t ですべてのアニメーションを固定します (シーク)
JSON としてのページ状態、ピクセルではなくトークン (スナップショット)
構造化エラーのある実際の入力イベント (クリック / 入力 / スクロール / アップロード)
JS エラー、コンソール出力、失敗したリクエスト ( console )
JSON 行または MCP 通知としてイベント サブスクリプションをプッシュします ( watch )
ポーリングによるワンコールページアサーション (expect)
ヘッドレス / バックグラウンド / ウィンドウごとのプロパティとしてフォーカス

、切り替え可能なライブ
人間によるハンドオフ: hwatu focus <id> がライブ セッションをタイリング WM にドロップします
CAPTCHA / 構造化された待機/再開によるボット対策検出 (チャレンジ)
MCP サーバー、プレーン CLI、および 1 行の JSON ソケット プロトコル
人間用の最小限の WebKit ブラウザ: ネイティブ広告ブロック、vim スタイル バー、クラッシュ復元
Playwright または chrome-devtools-mcp を使用しないのはなぜでしょうか?
エージェントにブラウザを提供するには 3 つの方法がありますが、そのうち 2 つは不適切です。
hwatu は、チェックを瞬時に行うもの (エンジン、GPU コンテキスト、
コンパイルされた広告ブロック、事前にウォーム化された WebView)、および
その前に座っている人間。タブなしでもアイドリングが暖かいのはそのためです
バー、そしてなぜ保温された Playwright サーバーが依然として同じ方法で駆動されているのか
クライアントあたり 341 ミリ秒、hwatu の 39 ミリ秒かかります (ベンチマーク)。
2 番目の違いは、何が返されるかです。劇作家と
chrome-devtools-mcp は、本質的には自動化です
API: エージェントにブラウザを操作させ、その後、生のデータを返すことができます。
エージェントが注目するスクリーンショットと DOM。
ファトゥは違います。検証ブラウザ：計測
プリミティブが組み込まれており、ブラウザ自体がウォームデーモンです
ここで、ウィンドウのコストは 13 ミリ秒で、ヘッドレスはウィンドウのプロパティであり、
起動モード。
そのため、ループは次のようになります。実際のコマンド、実際の出力は次のとおりです。
hwatu --headless localhost:3000 # そのウィンドウ;あなたはそれを決して見ません
hwatu --headless staging.example.com # リファレンス
hwatu diff --id 2 --other 1 --heatmap /tmp/heat.png
# {"match_percent":85.13,"regions":[{"x":0,"y":160,"w":2048,...}]}
hwatu motion --id 1 # 参照のアニメーション (数値として)
# イージング cubic-bezier(0.25,1,0.5,1)、300ms、マーキー 29.78px/s ...
# ...エージェントがコードを編集...
hwatu diff --id 2 --other 1
# {"match_percent":97.49} # 登山は推測に勝ります
このループは、stripe.com のランディング ページのクローンに対して実行しました。
エージェントがそれを受け取った

85.1% ～ 98.8% のピクセル一致。それを再現します:
スクリプト/デモ/ 。完全な検証パス (オープン、ロード、
eval、スクリーンショット、close) は 1 つのコマンド、1 つのツール呼び出し、最大 35 ミリ秒
中央値 (ベンチマーク):
hwatu チェック localhost:5173 --eval ' document.title ' --shot=/tmp/after.png
# {"title":"私のアプリ","eval":"私のアプリ","shot":"/tmp/after.png",
# "コンソール":[...],"load_ms":13,"total_ms":35}
生成された HTML は手元にありますが、サーバーはありませんか? hwatu レンダリングは同じです
入力としてマークアップを使用した 1 回の呼び出しパス: 一時ファイルなし、なし
python3 -m http.server :
echo ' <h1>生成</h1> ' | hwatu render --stdin --shot=/tmp/gen.png
# {"rendered":true,"shot":"/tmp/gen.png","load_ms":5,"total_ms":28}
# ポーリングせずにロード、コンソール、ダウンロード、ウィンドウ イベントに反応します。
hwatu watch --kinds ロード、コンソール
# {"event":"load","seq":1,"window_id":7,"data":{"state":"started",...}}
MCP クライアントは、同じストリームの submit_events を呼び出すことができます。
通知/hwatu/イベント 。各接続は厳密に単調になります
ゼロから始まるシーケンス。接続を閉じるか停止すると、接続が切断されます
デーモンをブロックせずにサブスクリプションを実行します。
同じパスが Playwright のウォーム インプロセス CDP 接続を通過します。
最良の場合は、82 ミリ秒と 5 回の API 呼び出しです。ファトゥのような形
実際に実行されます (新しいクライアントはそれぞれ、保温されたクライアントに対してチェックします)
エンジン)、Playwright のパスは 341 ミリ秒、hwatu の 39 ミリ秒です。hwatu は
暖かいデーモンは設計上、Playwright は暖かく保つ必要があるライブラリです
あなた自身。
そして、エージェントが CAPTCHA または判断コールを押すと、hwatu focus が表示されます。
ライブセッション、Cookie、状態をそのままの実体化します。
WMをタイリングします。 10秒間行動します。それは引き継がれます。これは
他のツールでは主張できない形容詞: 割り込み可能。
他の場所では、首なしは打ち上げ時に決定され、人間は決してできません
任意の価格でセッションをご覧いただけます。ファトゥではそれは窓の財産です、
切り替え可能 l

両方向に。
速度も同じ設計上の決定から生まれます: ファトゥは暖かいです
デーモン、あなたが起動するライブラリではありません。エンジン、GPU コンテキスト、
コンパイルされた広告ブロック ルールセットと、事前にウォームアップされた WebView は、
タスクなので、チェックはコールドパイプラインではなくホットパイプラインから開始されます。
プロセス。劇作家は図書館好きで、本質的に冷たい。それを維持する
Warm はユーザーが構築するもの (サーバー プロセス、接続)
管理、コンテキストプーリング)。 hwatu はデフォルトとして暖かい状態で出荷されます。
唯一のモード。
凡例: ✅ はい / 内蔵 · 🟡 部分的 / 限定的 · ❌ いいえ
1 toHaveScreenshot が保存されたゴールデンと比較します: 合格/不合格
テストスイートであり、エージェントが登れるスコアではありません。
2 標準的な方法は、アニメーションを無効にするか、早送りすることです。
フレークを避けるための最終状態。
3 Raw CDP はアニメーションの状態をクエリできますが、数値はありません
イージング/ベロシティ/キーフレームの概要。
4 立派なヘッドレス。すべての先頭のウィンドウがポップしてフォーカスを取得します。
比較は、執筆時点の各プロジェクトを反映しています。
修正は大歓迎です。正直な注意点: 劇作家がまだ勝つ
コールド スタート (190 対 435 ミリ秒、起動ごとに 1 回支払い) とメモリ。ファトゥ
Chromium ではなく WebKit をレンダリングします (Playwright マトリックスを CI に保持します)
エンジン固有のバグ）、現在は Linux のみです。フル
直接対決のデータと方法論:
docs/benchmarks.md 。
AGPL-3.0ライセンスを取得しています。リナックス。 WebKitGTK6.
AI コーディング エージェント用の高速で中断可能な検証ブラウザ: 35 ミリ秒のチェック、ピクセル差分、ライブ ヒューマン ハンドオフ
hongnoul.github.io/hwatu/ トピック
Readme AGPL-3.0 ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Fast, interruptible verification browser for AI coding agents: 35 ms checks, pixel diffs, live human hand-off - hongnoul/hwatu

GitHub - hongnoul/hwatu: Fast, interruptible verification browser for AI coding agents: 35 ms checks, pixel diffs, live human hand-off · GitHub
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
hongnoul
/
hwatu
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
221 Commits 221 Commits .astrophile .astrophile .github .github crates crates docs docs examples examples packaging packaging scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md VISION.md VISION.md deny.toml deny.toml glama.json glama.json llms.txt llms.txt server.json server.json View all files Repository files navigation
Make your agent harness loop instantly faster by giving it real eyes
STOP your agent claiming "pixel-perfect." Make it prove 97.49%.
STOP paying 5 tool calls per page check. hwatu check is one call, ~35 ms (beats warm-server Playwright ~9x).
STOP browser windows stealing your focus. Headless by default, you keep typing.
STOP shipping 170 MB of Chromium. One static binary + your distro's webkitgtk.
Vision : durable product principles, native platform strategy, swarm model
Agent guide : protocol, primitives, verification loops
Human guide : the tiling-WM browser side
Benchmarks : every number, measured, with methodology
Roadmap : plan of record, priorities, non-goals
Continuous improvement : activation metric, feedback loop, weekly cadence
Launch kit : reusable copy, channels, and measurement plan
Install → Detect workflow → Connect agent → Verify page → Hand off to human
curl -fsSL https://raw.githubusercontent.com/hongnoul/hwatu/main/scripts/install.sh | bash
hwatu setup
One static binary plus your distro's webkitgtk-6.0 (the installer
checks). On Arch: yay -S hwatu . From source: cargo build --release .
The installer installs binaries only. hwatu setup detects supported coding
agents and prints the available connections without changing their config.
Choose a client explicitly when you are ready:
hwatu doctor
hwatu setup --client claude --scope project --dry-run
hwatu setup --client claude --scope project
hwatu demo
Install: download two binaries and check WebKitGTK.
Detect: find Claude Code, Cursor, Jcode, or a generic MCP workflow.
Connect: use Jcode's native socket, MCP, or the CLI fallback.
Verify: run a headless rendered smoke test with doctor or demo .
Hand off: materialize the same live session only when a human is needed.
Setup is previewable, idempotent, and reversible with the same client and
scope plus --undo . Project scope creates shareable configuration; user scope
keeps it personal. Claude Code asks each user to approve project-scoped MCP
servers when it next starts.
Manual MCP configuration remains one portable entry:
{ "mcpServers" : { "hwatu" : { "command" : " hwatu " , "args" : [ " mcp " ] } } }
Or skip MCP entirely: every command is a short CLI call or one
newline-delimited JSON line over a Unix socket.
hwatu localhost:3000 # open a window like you open a terminal
Tried hwatu? A successful check, a failed install, and a missing workflow are
all useful signals. Share a two-minute use report
or report a bug .
Pixel-diff scoring: match percent + diff regions + heatmap ( diff )
Animations as numbers: duration, easing, velocity ( motion )
Deterministic animation frames: pin all animations at time t ( seek )
Page state as JSON, tokens not pixels ( snapshot )
Real input events with structured errors ( click / type / scroll / upload )
JS errors, console output, failed requests ( console )
Push event subscriptions as JSON lines or MCP notifications ( watch )
One-call page assertions with polling ( expect )
Headless / background / focused as a per-window property, switchable live
Human hand-off: hwatu focus <id> drops the live session into your tiling WM
CAPTCHA / anti-bot detection with structured wait/resume ( challenge )
MCP server, plain CLI, and a 1-line JSON socket protocol
A minimal WebKit browser for humans: native ad blocking, vim-style bar, crash restore
Why not Playwright or chrome-devtools-mcp?
There are three ways to give an agent a browser, and two of them are bad at it:
hwatu keeps exactly what makes checks instant (engine, GPU context,
compiled adblock, a prewarmed WebView) and nothing that serves a
human sitting in front of it. That's why it idles warm without a tab
bar, and why a kept-warm Playwright server driven the same way still
costs 341 ms per client to hwatu's 39 ( benchmarks ).
The second difference is what comes back. Playwright and
chrome-devtools-mcp are, at their core, automation
APIs: they let an agent drive a browser, then hand back raw
screenshots and DOM for the agent to eyeball.
hwatu is different. It is a verification browser: the measurement
primitives are built in, and the browser itself is a warm daemon
where a window costs 13 ms and headless is a window property, not a
launch mode.
That is why the loop looks like this, real commands, real output:
hwatu --headless localhost:3000 # its window; you never see it
hwatu --headless staging.example.com # the reference
hwatu diff --id 2 --other 1 --heatmap /tmp/heat.png
# {"match_percent":85.13,"regions":[{"x":0,"y":160,"w":2048,...}]}
hwatu motion --id 1 # the reference's animations, as numbers
# easing cubic-bezier(0.25,1,0.5,1), 300ms, marquee 29.78px/s ...
# ...agent edits code...
hwatu diff --id 2 --other 1
# {"match_percent":97.49} # climbing beats guessing
We ran this loop against a clone of stripe.com's landing page: an
agent took it from 85.1% to 98.8% pixel match . Reproduce it:
scripts/demo/ . A full verification pass (open, load,
eval, screenshot, close) is one command, one tool call, ~35 ms
median ( benchmarks ):
hwatu check localhost:5173 --eval ' document.title ' --shot=/tmp/after.png
# {"title":"My App","eval":"My App","shot":"/tmp/after.png",
# "console":[...],"load_ms":13,"total_ms":35}
Generated HTML in hand and no server? hwatu render is the same
one-call pass with the markup as input: no temp file, no
python3 -m http.server :
echo ' <h1>generated</h1> ' | hwatu render --stdin --shot=/tmp/gen.png
# {"rendered":true,"shot":"/tmp/gen.png","load_ms":5,"total_ms":28}
# React to load, console, download, and window events without polling.
hwatu watch --kinds load,console
# {"event":"load","seq":1,"window_id":7,"data":{"state":"started",...}}
MCP clients can call subscribe_events for the same stream as
notifications/hwatu/event . Each connection gets a strictly monotonic
sequence starting at zero; closing or stalling the connection drops its
subscription without blocking the daemon.
The same pass through Playwright's warm in-process CDP connection,
its best case, is 82 ms and five API calls. Shaped like hwatu
actually runs (a fresh client each check against a kept-warm
engine), Playwright's pass is 341 ms vs hwatu's 39 : hwatu is a
warm daemon by design, Playwright is a library you have to keep warm
yourself.
And when the agent hits a CAPTCHA or a judgment call, hwatu focus
materializes its live session, cookies and state intact, in your
tiling WM. You act for ten seconds. It takes back over. This is
the adjective no other tool gets to claim: interruptible.
Everywhere else, headless is decided at launch and a human can never
see the session at any price. In hwatu it's a window property,
switchable live, in both directions.
The speed follows from the same design decision: hwatu is a warm
daemon , not a library you launch. The engine, the GPU context, the
compiled adblock ruleset, and a prewarmed WebView outlive every
task, so a check starts from a hot pipeline instead of a cold
process. Playwright is a library and cold by nature. Keeping it
warm is something you build (a server process, connection
management, context pooling); hwatu ships warm as the default and
the only mode.
Legend: ✅ Yes / built-in · 🟡 Partial / limited · ❌ No
1 toHaveScreenshot compares against stored goldens: pass/fail for
test suites, not a score an agent can climb.
2 Standard practice is to disable animations or fast-forward to the
end state to avoid flakes.
3 Raw CDP can query animation state, but there is no numeric
summary of easing/velocity/keyframes.
4 Fine headless; every headed window pops and takes focus.
Comparison reflects each project at the time of writing;
corrections are welcome. Honest caveats: Playwright still wins
cold start (190 vs 435 ms, paid once per boot) and memory; hwatu
renders WebKit not Chromium (keep a Playwright matrix in CI for
engine-specific bugs), and it is Linux-only today. Full
head-to-head data and methodology:
docs/benchmarks.md .
AGPL-3.0 licensed. Linux. WebKitGTK 6.
Fast, interruptible verification browser for AI coding agents: 35 ms checks, pixel diffs, live human hand-off
hongnoul.github.io/hwatu/ Topics
Readme AGPL-3.0 license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
