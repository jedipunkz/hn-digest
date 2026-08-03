---
source: "https://github.com/vkarach/spriteloom"
hn_url: "https://news.ycombinator.com/item?id=49154659"
title: "Spriteloom – local AI pixel-art generator plugin for Aseprite"
article_title: "GitHub - vkarach/spriteloom · GitHub"
author: "vkarach"
captured_at: "2026-08-03T12:50:23Z"
capture_tool: "hn-digest"
hn_id: 49154659
score: 1
comments: 0
posted_at: "2026-08-03T12:06:49Z"
tags:
  - hacker-news
  - translated
---

# Spriteloom – local AI pixel-art generator plugin for Aseprite

- HN: [49154659](https://news.ycombinator.com/item?id=49154659)
- Source: [github.com](https://github.com/vkarach/spriteloom)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T12:06:49Z

## Translation

タイトル: Spriteloom – Aseprite 用ローカル AI ピクセルアート ジェネレーター プラグイン
記事タイトル: GitHub - vkarach/spriteloom · GitHub
説明: GitHub でアカウントを作成して、vkarach/spriteloom の開発に貢献します。

記事本文:
GitHub - vkarach/spriteloom · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ヴカラチ
/
スプライトルーム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 操作

その他のアクション メニュー フォルダーとファイル
149 コミット 149 コミット .github .github アセット アセット ランチャー ランチャー プラグイン プラグイン サーバー サーバー サイト サイト .gitignore .gitignore .luacheckrc .luacheckrc CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md TODO.md TODO.md build.spec build.spec install-plugin.bat install-plugin.bat start-server.bat start-server.bat すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Aseprite のローカル AI ピクセル アート アシスタント。
テキスト プロンプトからスプライトを生成し、既存のスプライトを編集します。
命令を実行したり、選択した領域を再描画したりすることはすべて独自の GPU 上で実行されます。
クラウドやサブスクリプションは不要で、ピクセルがマシンから離れることはありません。
サンプル出力とその仕組み → vkarach.github.io/spriteloom
AIによる生成・編集・修復
回転/指示、2 つの異なるターンアラウンド
デモではすべてのクリップが高速化されました。
実際の出力、それぞれ 1 世代。 「背の高い男性」のような主題
ダークコートを着た」、「赤い革のコートを着た風化した探検家」、「木造の小屋、
茅葺き屋根」。表示用にハード ピクセル エッジを使用して拡大されます。
拡散モデルをローカルで実行する WebSocket サーバーと Aseprite
それに話しかける拡張機能。あなたはずっとアスプライトにいます。結果
サイドウィンドウで開き、クリック時に新しいレイヤーとしてドロップインします。決して編集しない
既存のピクセル。
生成 — テキスト プロンプトからスプライトを生成します。
AI による編集 — 命令によって既存のスプライトを変更します (「
剣が青く光る");スタイルと言及されていないものはすべてそのままです。
選択範囲の修復 — 同じですが、選択された領域のみに触れます。
回転 + 指示 — 同じ被写体を別の角度から再表示します。
すべてが単一のモデル (FLUX.2 Klein) 上で実行されるため、モデルはありません
スワップ: 最初のロード後、タスクは数秒で応答します。
必要なハードウェア - 最初にお読みください
これは 4B パラメータの拡散モデルを実行します。

自分のマシン。そうではありません
軽量ツール:
NVIDIA GPU も Spriteloom もありません。 CPU フォールバックもクラウドもありません
設計上のオプション — 重要なのは、ローカルで実行されるということです。
セットアップの VRAM モード ドロップダウンでは、約 8 GB のトランスをどのように搭載するかを選択します。
GPU:
自動 (デフォルト) — 空き VRAM を検出し、12 GB 以上のカード上の bf16 を選択します。
それ以下のレガシー 8 GB モード。
bf16 — モデル全体が GPU 上に常駐します。最速。 12GB以上が必要
無料のVRAM。
レガシー 8 GB モード — 8 GB カード用 (ラップトップ RTX 4060 など)。モデル
一度に GPU に収まらないため、1 つのレイヤーを GPU に移動します
そのレイヤーが実行される直前に、それをシステム RAM にスワップアウトして戻します。 VRAM
使用率は常に低いままです (モデル全体ではなく、1 つのレイヤーに存在します)。
GPU の使用率も同様です。ほとんどの時間はウェイトの移動に費やされます。
PCIe 上では、計算するためではなく、それがこれが遅い本当の理由です。
出力は bf16 と同じですが、かなり遅いです。 8GBより速いものはない
現在のパス。 fp8 量子化モードを試してみたところ、結果は改善されませんでした
これよりも実際の品質リスクが増大するため、削除されました。
レガシー 8 GB モードでは、最大 16 GB のモデル全体がシステム RAM に保持されます (つまり、
レイヤーの交換元)。のみに適合する 16 GB マシン上
Windows ページ ファイルは非常に大きいため、有効にしてサイズを変更したままにしておきます (または、次のように設定します)。
システム管理)。空き RAM とページ ファイルがモデルをカバーできない場合、負荷がかかります
失敗する。サーバーは起動時にこれをチェックし、代わりに理由を出力するようになりました。
クラッシュする。
最新のビルドを次からダウンロードします。
itch.io を開き、任意の場所に解凍します。それ
Spriteloom.exe、server/、plugin/ が一緒に含まれています。他に何もすることはない
最初に取得します。
Windows には「Windows が PC を保護しました」と表示される可能性があります。
警告 — Spriteloom.exe はコード署名されていないため、評判がありません
まだ。 「詳細情報」をクリックし、「とにかく実行」をクリックします。
[セットアップ] を押します。それは示します

欠けているもの: 環境、PyTorch、
依存関係、プラグイン、モデル (~15 GB)、スタート メニューのショートカット。
すべてにチェックを入れて、「選択したものをインストール」を押します。 .venv をビルドします。
パッケージとプラグインをインストールし、モデルをダウンロードし、スタートボタンを追加します
メニューのショートカット、およびライブ ログを印刷します。プラグインを削除したら Aseprite を再起動します
入っています。
モデルを含む必要なすべての部分が完成するまで、開始は無効のままになります。
所定の位置にあります。遅延した初回実行ダウンロードはありません。
手動で行いたい場合は、次のようにします。
.venv\Scripts\python -m pip install -r server\requirements.txt
.venv\Scripts\python -m pip install torch --index-url https://download.pytorch.org/whl/cu128
install-plugin.bat を実行し、Aseprite を再起動します。
このように .venv を設定すると、start-server.bat はサーバーを起動します。
ランチャー経由ではなくコンソール ウィンドウから。
このリポジトリの完全なクローンから自分で exe をビルドするには:
.venv\Scripts\python -m pip install -r launcher\requirements.txt 、次に
.venv\Scripts\python -m PyInstaller build.spec --distpath 。 -- それは構築します
Spriteloom.exe としてプロジェクト ルートに直接コピーします。約 15 MB: モデル
そして PyTorch はその外側に留まります。
Spriteloom.exe を実行し、START キーを押してウィンドウを開いたままにします。ドット
モデルが常駐すると、ウォームアップから約 25 秒後に緑色に変わります。
始めます。ウィンドウを閉じるとサーバーが停止します。
Aseprite の場合: Sprite → Spriteloom... (または F1 を押します)。タスクを選択し、
フィールドに入力し、「実行」を押します。結果は別のウィンドウで開きます。クリック
新しいレイヤーとして挿入するバリアントです。
Generate は全文を理解します。View プリセットを選択し、名前を付けます。
件名 (「ダークブラウンの革表紙の閉じた本」)、追加の追加
必要に応じて詳細 — パネルには、送信される正確なテキストが表示されます。
編集/修復は、強さのスライダーではなく指示を受けます: 何をするかを指示します
変更とその金額。 Inpaint は sel のみをタッチします

セクション。
回転 / 指示 : 主題に明示的に名前を付けます (「四本足の茶色
「キャラクター」ではなく「馬」です）。オプションのミラー対称により結果が強制されます
左右対称で、正面と背面のビューをまとめるのに役立ちます -
しかし、それは片手武器や武器を含むあらゆるものを反映します。
ポーズが非対称なので、間違った印象を与える可能性もあります。試してみてください、しないでください
それが常に良い結果であると想定してください。
[詳細...] を選択すると、背景、パレット、シード、および
追加 (生成および回転/指示のプロンプトに追加されます。編集および
インペイントは無視します)。
背景 : 均一な背景を自動検出して除去し、削除します。
主要な境界線の色を取り除き、Keep は完全に不透明なままにします。
パレット : 結果ごとに色を自動的に導き出します。現在のパレット ピンの出力先
開いているスプライトのパレット全体。選択したカラーはスウォッチのみに固定されます
パレットバーで強調表示されます。パレット ファイルは .gpl / .pal / .png に固定されます
ファイルに保存すると、スプライトのバッチが 1 つのカラー セットを共有します。
履歴は、最新のものから順に過去の世代 (output/ に保存) を参照します。
実行をクリックしてそのバリアントを表示し、バリアントをクリックして挿入します。
+----------------------+ +---------------+ +----------------------+
| Aseprite プラグイン (Lua) | <-- WebSocket --> | Pythonサーバー | --> | FLUX.2 クライン (GPU) |
|ダイアログ、結果、 | |プロトコル、 | |独身居住者 |
|履歴、レイヤー挿入 | |後処理 | |モデル、スワップなし |
+----------------------+ +---------------+ +----------------------+
1 つのモデル、1 つのウォームアップ。すべてのタスクは同じ FLUX.2 Klein パイプラインにヒットします。
したがって、タスクごとのロード/アンロードはありません。 12 GB 以上のカードでは完全に維持されます
居住者。 8 GB カード (レガシー 8 GB モード) では、各レイヤーが GPU にストリーミングされます
代わりにタスクごとに、そのモードの追加時間がそこに費やされます。
境界でリクエストを検証する WebSocket プロトコルと
進行状況メッセージをストリーミングして戻す

プラグイン (server/protocol.py、
サーバー/main.py)。
ポストプロセス パイプラインは、生の拡散出力をクリーンなスプライトに変換します。
被写体に合わせてクロップ、パレットの量子化、背景の削除、ミラー
対称性、キャンバスにフィットする (server/postprocess.py)。
このプラグインは、純粋な単体テスト済みのプロンプト アセンブリを備えたモジュラー Lua です。
スタブ化された Aseprite API に対してテストされた UI レイヤー — 以下を参照してください。
クライン変換器を 8 ビット量子化すると、純粋なノイズが生成されます。
テキストから画像へ (編集は正常に機能します)。出荷されたセットアップはこれを回避します: 8 ビット
テキストエンコーダー + bf16 トランスフォーマー、完全常駐。 TODO.md を参照してください。
.venv\Scripts\python -m pytest サーバー/テスト/ --ignore=サーバー/テスト/smoke.py
Aseprite を使用しない迅速なチューニング (生のバリアントと後処理されたバリアントを出力/に書き込みます):
.venv\Scripts\python -m server.tests.smoke "魔剣" --size 64
プラグイン テスト (ユーザー スコープの lua luacheck を scoop install する必要があります):
luacheck プラグイン\*.lua プラグイン\tests\*.lua
lua プラグイン\tests\test_prompt.lua
lua プラグイン\tests\test_panel.lua
luacheck は Lua 5.4 (Aseprite が実行されるバージョン) をバンドルします。 .luacheckrc
挿入する API グローバルを宣言します。 test_panel.lua はすべてのモジュールをロードします
スタブ化された Aseprite API に対して、各サーバーのステータス キャンバスを再描画します
state — これにより、エディタを起動せずに、壊れたモジュール間の呼び出しが捕捉されます。
レイアウトとゴースティングには依然として本物の Aseprite が必要です。
ファイル
保持します
main.lua
エントリポイント、メニューコマンドを登録します
ダイアログ.lua
コントロールパネル
結果.lua
結果ウィンドウ (新しいバリアント)
歴史.lua
履歴リストと単一実行ウィンドウ
ウイ・ルア
テーマカラー、チェッカーボード、バリアントグリッド、プロンプトプレビュー
スプライト.lua
フレーム/マスクのエクスポート、バリアントをレイヤーとして挿入
プロンプト.lua
プロンプトアセンブリとキーマップ (純粋な Lua、単体テスト済み)
client.lua
WebSocketクライアント
Base64.lua
Base64コーデック
ランチャーのレイアウト
ファイル
保持します
ランチャー/app.py
ウィンドウ、JSブリッジ、

ウィンドウのサイズ変更
ランチャー/ui/index.html
マークアップとスタイル、メイン画面とセットアップ画面
ランチャー/server_proc.py
サーバーのサブプロセス、ポートのプローブ、ヘルス
ランチャー/plugin_install.py
プラグインのコピー、バージョン比較
ランチャー/paths.py
ルート、Python、Aseprite、モデルの検索
ランチャー/setup_checks.py
セットアップに欠けているもの、検出のみ
ランチャー/setup_steps.py
インストール手順を順番に実行する
サーバー/config.py
ランチャーとサーバーによって共有される設定
ランチャーは設定ファイル %APPDATA%\Spriteloom\config.json を所有しています。
ポート、VRAM モード、セットアップ パスはすべてそこに存在し、すべての書き込みが行われます。
マージされるので、あるキーが別のキーを消去することはありません。ポートのデフォルトは 8765 であり、
インストールすると、プラグインの隣にあるserver.jsonにそれがスタンプされるため、両端が
同意します。
サーバーは、ランチャーとともに終了する Windows ジョブ オブジェクト内で実行されます。
これにより、クラッシュまたは強制終了されたランチャーがサーバーから残されるのを防ぐことができます。
VRAM を保持します。
サーバーがクラッシュした、またはロードが完了しないと言っている場合は、組み込みの
インストールフォルダーからローダーをチェックします。
.venv\Scripts\python -m サーバー.診断
パッケージのバージョン、メモリのヘッドルーム、モデル ファイルが保存されているかどうかを出力します。
完了したら、トランスを独自にロードします。 TRUNCATED とは、
ダウンロードがカットされました

[切り捨てられた]

## Original Extract

Contribute to vkarach/spriteloom development by creating an account on GitHub.

GitHub - vkarach/spriteloom · GitHub
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
vkarach
/
spriteloom
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
149 Commits 149 Commits .github .github assets assets launcher launcher plugin plugin server server site site .gitignore .gitignore .luacheckrc .luacheckrc CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md TODO.md TODO.md build.spec build.spec install-plugin.bat install-plugin.bat start-server.bat start-server.bat View all files Repository files navigation
Local AI pixel-art assistant for Aseprite .
Generate sprites from a text prompt, edit existing sprites with an
instruction, or redraw a selected region — all running on your own GPU.
No cloud, no subscription, your pixels never leave your machine.
Sample output and how it works → vkarach.github.io/spriteloom
Generate · Edit with AI · Inpaint
Rotate / Instruct, two different turnarounds
All clips sped up for the demo.
Real output, one generation each. Subjects like "tall man
in a dark coat", "weathered explorer in a red leather coat", "wooden cottage,
thatched roof". Scaled up with hard pixel edges for display.
A WebSocket server that runs a diffusion model locally, plus an Aseprite
extension that talks to it. You stay in Aseprite the whole time; results
open in a side window and drop in as new layers on click. It never edits
your existing pixels.
Generate — a sprite from a text prompt.
Edit with AI — change an existing sprite by instruction ("make the
sword glow blue"); style and everything unmentioned stays put.
Inpaint Selection — same, but only ever touches the selected region.
Rotate + Instruct — re-view the same subject from another angle.
Everything runs on a single model (FLUX.2 Klein), so there are no model
swaps: after the first load, tasks respond in seconds.
Hardware you need — read this first
This runs a 4B-parameter diffusion model on your own machine. It is not
a lightweight tool:
No NVIDIA GPU, no Spriteloom. There is no CPU fallback and no cloud
option by design — the whole point is that it runs locally.
Setup's VRAM mode dropdown picks how the ~8 GB transformer is fit onto
the GPU:
Auto (default) — detects free VRAM and picks bf16 on 12+ GB cards,
Legacy 8 GB mode below that.
bf16 — the whole model stays resident on the GPU. Fastest; needs 12+ GB
free VRAM.
Legacy 8 GB mode — for 8 GB cards (e.g. a laptop RTX 4060). The model
doesn't fit on the GPU all at once, so it moves one layer onto the GPU
right before that layer runs, then swaps it back out to system RAM. VRAM
usage stays low the whole time (one layer resident, not the whole model),
and so does GPU utilization: most of the time goes to shuttling weights
over PCIe, not to compute, and that is the actual reason this is slow.
Output is identical to bf16, just much slower. There is no faster 8 GB
path currently. An fp8-quantized mode was tried and measured no better
than this, while adding a real quality risk, so it was dropped.
Legacy 8 GB mode keeps the whole ~16 GB model in system RAM (that is
where layers are swapped from). On a 16 GB machine that only fits with a
generous Windows page file, so leave it enabled and sized (or set it to
system-managed). If free RAM plus page file cannot cover the model the load
fails; the server now checks this at startup and prints why instead of
crashing.
Download the latest build from
itch.io and unzip it anywhere. It
has Spriteloom.exe , server/ , and plugin/ together; nothing else to
fetch first.
Windows will likely show a "Windows protected your PC" SmartScreen
warning — Spriteloom.exe isn't code-signed, so it has no reputation
yet. Click More info , then Run anyway .
Press Setup . It shows what is missing: the environment, PyTorch,
the dependencies, the plugin, the model (~15 GB), a Start Menu shortcut.
Tick everything and press Install selected . It builds the .venv ,
installs the packages and the plugin, downloads the model, adds a Start
Menu shortcut, and prints a live log. Restart Aseprite once the plugin
is in.
Start stays disabled until every required piece — including the model —
is in place; there is no lazy first-run download.
If you would rather do it by hand:
.venv\Scripts\python -m pip install -r server\requirements.txt
.venv\Scripts\python -m pip install torch --index-url https://download.pytorch.org/whl/cu128
install-plugin.bat , then restart Aseprite.
With the .venv set up this way, start-server.bat starts the server in a
console window instead of through the launcher.
To build the exe yourself from a full clone of this repository:
.venv\Scripts\python -m pip install -r launcher\requirements.txt , then
.venv\Scripts\python -m PyInstaller build.spec --distpath . -- it builds
straight into the project root as Spriteloom.exe , about 15 MB: the model
and PyTorch stay outside it.
Run Spriteloom.exe , press START and leave the window open. The dot
turns green once the model is resident, about 25 seconds after a warm
start. Closing the window stops the server.
In Aseprite: Sprite → Spriteloom... (or press F1 ). Pick a task,
fill the fields, press Run . Results open in a separate window; click
a variant to insert it as a new layer.
Generate understands full sentences: pick a View preset, name the
Subject ("closed book with dark brown leather cover"), add Extra
details if needed — the panel shows the exact text it will send.
Edit / Inpaint take instructions, not a strength slider: say what to
change and how much. Inpaint only touches the selection.
Rotate / Instruct : name the subject explicitly ("four-legged brown
horse", not "character"). Optional Mirror symmetry forces the result
left/right symmetric, which can help a front/back view hold together -
but it mirrors everything, including a one-handed weapon or an
asymmetric pose, so it can also make those look wrong. Try it, don't
assume it's always the better result.
Advanced... opens a separate window with Background, Palette, Seed, and
Extra (appended to the prompt for Generate and Rotate/Instruct; Edit and
Inpaint ignore it).
Background : Auto detects and strips a uniform background, Remove
strips the dominant border color, Keep leaves it fully opaque.
Palette : Auto derives colors per result; Current palette pins output to
the open sprite's whole palette; Selected colors pins to only the swatches
highlighted in the palette bar; Palette file pins to a .gpl / .pal / .png
file so a batch of sprites shares one set of colors.
History browses past generations (stored in output/ ), newest first;
click a run to see its variants, click a variant to insert it.
+-----------------------+ +---------------+ +--------------------+
| Aseprite plugin (Lua) | <-- WebSocket --> | Python server | --> | FLUX.2 Klein (GPU) |
| dialogs, results, | | protocol, | | single resident |
| history, layer insert | | postprocess | | model, no swaps |
+-----------------------+ +---------------+ +--------------------+
One model, one warm-up. Every task hits the same FLUX.2 Klein pipeline,
so there is no per-task load/unload. On 12+ GB cards it stays fully
resident; on 8 GB cards (Legacy 8 GB mode) each layer streams to the GPU
per task instead, which is where that mode's extra time goes.
WebSocket protocol with request validation at the boundary and
streamed progress messages back to the plugin ( server/protocol.py ,
server/main.py ).
Postprocess pipeline turns raw diffusion output into a clean sprite:
crop to subject, palette quantization, background removal, mirror
symmetry, fit-into-canvas ( server/postprocess.py ).
The plugin is modular Lua with pure, unit-tested prompt assembly and a
UI layer tested against a stubbed Aseprite API — see below.
8-bit quantizing the Klein transformer produces pure noise in
text-to-image (edits work fine). The shipped setup sidesteps this: 8-bit
text encoder + bf16 transformer, fully resident. See TODO.md .
.venv\Scripts\python -m pytest server/tests/ --ignore=server/tests/smoke.py
Prompt tuning without Aseprite (writes raw + postprocessed variants to output/ ):
.venv\Scripts\python -m server.tests.smoke "demonic sword" --size 64
Plugin tests (needs scoop install lua luacheck , user-scoped):
luacheck plugin\*.lua plugin\tests\*.lua
lua plugin\tests\test_prompt.lua
lua plugin\tests\test_panel.lua
luacheck bundles Lua 5.4 (the version Aseprite runs); .luacheckrc
declares the API globals it injects. test_panel.lua loads every module
against a stubbed Aseprite API and repaints the status canvas in each server
state — this catches broken cross-module calls without launching the editor.
Layout and ghosting still need a real Aseprite.
file
holds
main.lua
entry point, registers the menu command
dialogs.lua
the control panel
results.lua
results window (fresh variants)
history.lua
history list and single-run windows
ui.lua
theme colors, checkerboard, variant grid, prompt preview
sprite.lua
frame/mask export, inserting variants as layers
prompt.lua
prompt assembly and key maps (pure Lua, unit-tested)
client.lua
WebSocket client
base64.lua
base64 codec
Launcher layout
file
holds
launcher/app.py
the window, the JS bridge, window sizing
launcher/ui/index.html
markup and styles, main and setup screens
launcher/server_proc.py
the server subprocess, port probing, health
launcher/plugin_install.py
copying the plugin, version comparison
launcher/paths.py
finding the root, Python, Aseprite and the model
launcher/setup_checks.py
what the setup is missing, detection only
launcher/setup_steps.py
running the install steps in order
server/config.py
the settings, shared by the launcher and the server
The launcher owns the settings file %APPDATA%\Spriteloom\config.json : the
port, the VRAM mode, and the setup paths all live there, and every write
merges so one key never erases another. The port defaults to 8765, and
Install stamps it into server.json next to the plugin, so both ends
agree.
The server runs inside a Windows job object that dies with the launcher.
That is what keeps a crashed or killed launcher from leaving a server behind
holding your VRAM.
If the server says it crashed or never finishes loading, run the built-in
loader check from the install folder:
.venv\Scripts\python -m server.diagnose
It prints package versions, memory headroom, and whether the model files are
complete, then loads the transformer on its own. TRUNCATED means the
download was cut

[truncated]
