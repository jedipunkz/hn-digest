---
source: "https://github.com/thomas-luebker/amimcp"
hn_url: "https://news.ycombinator.com/item?id=49228811"
title: "Amimcp – let Claude work on your Amiga"
article_title: "GitHub - thomas-luebker/amimcp: Let Claude work on a real Amiga — an MCP server plus a tiny AmigaOS daemon. Shell, files, screen capture, mouse and keyboard. · GitHub"
author: "nickt"
captured_at: "2026-08-09T06:38:23Z"
capture_tool: "hn-digest"
hn_id: 49228811
score: 1
comments: 0
posted_at: "2026-08-09T05:59:20Z"
tags:
  - hacker-news
  - translated
---

# Amimcp – let Claude work on your Amiga

- HN: [49228811](https://news.ycombinator.com/item?id=49228811)
- Source: [github.com](https://github.com/thomas-luebker/amimcp)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T05:59:20Z

## Translation

タイトル: Amimcp – クロードに Amiga を操作してもらいます
記事のタイトル: GitHub - thomas-luebker/amimcp: Claude を実際の Amiga (MCP サーバーと小さな AmigaOS デーモン) で動作させます。シェル、ファイル、画面キャプチャ、マウス、キーボード。 · GitHub
説明: Claude を実際の Amiga (MCP サーバーと小さな AmigaOS デーモン) で動作させます。シェル、ファイル、画面キャプチャ、マウス、キーボード。 - トーマス・リューブカー/amimcp

記事本文:
GitHub - thomas-luebker/amimcp: Claude を実際の Amiga (MCP サーバーと小さな AmigaOS デーモン) で動作させます。シェル、ファイル、画面キャプチャ、マウス、キーボード。 · GitHub
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
トーマス・リューブカー
/
amimcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
30 コミット 30 コミット エージェント エージェント ドキュメント ドキュメント サーバー サーバー テスト テスト ツール ツール .gitignore .gitignore ライセンス ライセンス通知なし

ICE PROTOCOL.md PROTOCOL.md README.md README.md install.sh install.sh run_tests.sh run_tests.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
amimcp — クロードに Amiga を操作してもらいます
クロードに実際の操作を提供する MCP サーバー
本物のアミーガ。 AmigaDOS コマンドの実行、ファイルの読み取りと書き込み、ドロワーのリスト表示、
マシンを検査し、画面をキャプチャし、マウスを動かして入力します。
あなたは問題を説明します。クロードは実際のハードウェアをつつきます。
それは、クロードが A4000/060 のシェルをクリックしてコマンドを入力し、
Return キーを押すと、ウムラウトが Amiga 独自のドイツ語キーマップを介してルーティングされます。
┌─────────┐ MCP over stdio ┌──────┐ フレーム化 TCP ┌────────┐
│ クロード・コード │ ◄─────────────► │ amimcp │ ◄───────► │ amiagent │
│ (お使いの Mac/PC)│ JSON-RPC 2.0 │ (Python) │ ポート 7846 │ (Amiga)│
━━━━━━━━━━━━━━━━━━━━━━━━━━┘
2 つの半分:
amimcp — Claude を実行しているマシン上の MCP サーバー。ピュアパイソン
標準ライブラリ: 設定した日の間に venv、pip、ロックファイルが腐ることはありません
これと2年後の夕方、Amigaが起動しなくなったときのこと。
amiagent — Amiga 上の小さな C デーモン。 AmigaOS 2.04 以降、ベア
68000、約 85 KB、 bsdsocket.library のみが必要です。
ツール
何をするのか
amiga_shell
AmigaDOS コマンドの実行 — stdout、stderr、および戻りコード
amiga_read_file
ファイルを読み取ります。テキストインライン、バイナリbase64
amiga_write_file
バイナリを含むファイルの書き込みまたは上書き
amiga_list_dir
サイズ、保護ビット、日付スタンプを含むドロワーをリストします。
amiga_system_info
キックスタート、CPU/FPU、空きチップ/高速、ボリューム

s、代入
amiga_スクリーンショット
画面を PNG としてキャプチャ — 平面または RTG、全体または領域
amiga_screens
開いているすべての画面を前から後ろにジオメトリとともにリストします。
amiga_pointer
ポインタの位置 — 画像は転送されません
amiga_region_changed
領域のチェックサム: 「まだ何か変更がありましたか?」
amiga_click
ポインタを移動して、シングル、ダブル、任意のボタンをクリックします。
amiga_drag
押す、移動する、離す — ドラッグ アンド ドロップ、スクロールバー、メニュー
amiga_button
ボタンを押し続けるか放すだけ
amiga_move_mouse
クリックせずにポインタを移動します
amiga_type
Amiga 独自のキーマップを通じてマッピングされたテキストを入力します
amiga_key
Return、Esc、F キーなどを修飾子付きで押します
amiga_break
Ctrl-C amiga_shell によって実行されたままのコマンド
実際に動作するには、これらを組み合わせるだけで十分です。スタートアップ シーケンスを読み取って修正し、
バイナリをクロスコンパイルしてプッシュして実行し、GUI プログラムを起動して、
クリックして操作するか、マシンが起動したときに表示される内容を観察するだけです。
間違ってしまいました。
GUI を操作する前に、次の 2 つのことを知っておく価値があります。
安く視聴しましょう。フル 1080p トゥルーカラー フレームは、~1 MB/秒のリンク上で~6 MB です —
約7秒。 1 つのステータス行をリージョンとして読み取ることは事実上無料であり、
amiga_region_changed は「描画は終了しましたか?」と答えます。 4バイトで。投票
ハッシュ付き。見る必要がある場合にのみピクセルをキャプチャします。
ポインタは2種類。 amiga_click は Intuition ポインタをワープさせます。
直感的な応用が続きます。 SDL プログラム — ゲーム、ScummVM — 生のマウスを追跡する
デルタを作成し、独自のカーソルを保持するため、ワープやクリックが表示されることはありません
ポインターがまだあると信じている場所に着陸します。相対パス: true
それら。右ボタンを押しながら移動するとAmigaメニューが開きます。
したがって、クリックではなく amiga_drag が必要です。
1. Amiga にエージェントをインストールする
amiagent-0.5.3.lha を取得します
Amiga 上で解凍します。それは含まれています

ns amiagent (68000、すべてで実行)
およびamiagent.020 (68020+)。
amipkg アミエージェントをインストールする
または自分で構築する
m68k-amigaos クロスツールチェーンが必要 —
AmigaPorts/m68k-amigaos-gcc 、
bebbo の amiga-gcc の維持された子孫。 (Bebbo 自身のリポジトリはなくなりました
2026 年 8 月 7 日現在。 AmigaPorts が現在存在する場所です。)
それを備えたマシン:
CD エージェント && # 68000 ベースラインを作成、すべての Amiga で実行
cd エージェント && make CPU=68020 # 68020+ ビルド
結果として得られたamiagentをAmigaにコピーします。
未定義の IntuitionBase によりリンクが失敗した場合、ツールチェーンは
直感/グラフィックを自動的に開きます。 make OWNBASES=1 で再構築します。
TCP/IP スタック (Roadshow、AmiTCP、またはエミュレータの bsdsocket) を起動します。
エミュレーション)、シェルから:
amiagent TOKEN=pickasecret
ポート 7846 でリッスンしていることを出力して待機します。 Ctrl-C を押すと停止します。もし
バックグラウンドに設定した場合は、Status と Break <n> C で見つけます。
オプション: PORT/N (デフォルトは 7846)、 TOKEN/K 、 QUIET/S 。
起動時に起動するには、これを TCP スタックの後の S:User-Startup に追加します。
出てきます:
>NIL を実行: amiagent TOKEN=pickasecret QUIET
3. アシスタントにそれを向けます
./install.sh 192.168.1.42 pickasecret
これは、最初に Amiga をプローブし、次にサーバーを Claude Code に登録します。あるいはそうする
それを手で：
クロード・MCP アミーガを追加 \
--env AMIGA_HOST=192.168.1.42 \
--env AMIGA_TOKEN=pickasecret \
-- python3 /path/to/amimcp/server/amimcp.py
Claude Desktop の場合、これを claude_desktop_config.json に追加します。
{
"mcpサーバー": {
"アミガ" : {
"コマンド" : " python3 " ,
"args" : [ " /path/to/amimcp/server/amimcp.py " ],
"環境" : {
"AMIGA_HOST" : " 192.168.1.42 " 、
"AMIGA_TOKEN" : "ピッカシークレット"
}
}
}
}
新しいセッションを開始し、クロードに Amiga をチェックするように依頼します。
その他の MCP クライアント — Codex、Cursor、Zed、Cline…
MCP はオープンスタンダードであり、server/amimcp.py は通常の stdio MCP です。
サー

ver : stdin/stdout 経由で JSON-RPC を話し、構成を次から取得します。
環境変数。したがって、MCP サーバーを起動できるものはすべて、
Amiga を駆動する — サーバーは、どのモデルがもう一方のモデルであるかを認識せず、気にしません
終わり。
構成ファイルの形式のみが異なります。 OpenAI Codex CLI ( ~/.codex/config.toml )
TOML を使用します。
[ mcp_servers .アミーガ]
コマンド = " python3 "
args = [ " /path/to/amimcp/server/amimcp.py " ]
env = { AMIGA_HOST = " 192.168.1.42 " 、AMIGA_TOKEN = " pickasecret " }
Google Gemini CLI、カーソル、Zed、Cline、Continue、
Windsurf 、LM Studio、Goose は同じ mcpServers JSON ブロックを使用します
上にクロード デスクトップの場合、独自の設定ファイルに示したものがあります。クライアントのを確認してください
正確なパスについては現在のドキュメントを参照してください。これらは移動します。
Claude Code および Claude Desktop に対してテストされました。残りは次のようになります。
私がすべてを実行したのではなく、標準の stdio サーバー。必要であれば
何か違う場合は、問題を開いて何を言ってください。
Python 3 が必要です — 標準ライブラリのみ、 pip install 、venv は不要です。
LLM を使用せずにリンクをチェックする
AMIGA_HOST=192.168.1.42 AMIGA_TOKEN=pickasecret python3 サーバー/amimcp.py --probe
これは MCP を完全にスキップするため、ここでの障害はむしろネットワークまたはエージェントの問題です。
クライアントよりも。
変数
デフォルト
意味
AMIGA_HOST
(必須)
Amiga の IP アドレスまたはホスト名
アミガポート
7846
ポートアミエージェントがリッスンします
AMIGA_TOKEN
(なし)
エージェントの TOKEN= と一致する必要があります
AMIGA_TIMEOUT
30
デフォルトのリクエストごとのタイムアウト、秒
セキュリティ - ポートを開く前にこれをお読みください
amiagent は、そのポートにアクセスできる人に対して任意のコマンドを実行します。
プロトコルは暗号化されていません。これは意図的な取引です: 従来の Amiga には TLS がありません
AmiSSL を使用せず、接続しようとしているマシンに到達するためだけに AmiSSL を必要とする
修理では目的が果たせません。
つまり: TOKEN を設定します

、信頼できる LAN 上に保持し、ポートを転送しないでください。
トークンがなければ、ネットワーク セグメント上のあらゆるものでフォーマットを実行できます。
システムパーティション。エージェントなしで開始すると、エージェントは大声で警告します。
トークンはバイトごとに比較されます (したがって、A4000 と a4000 は異なります)
トークン）、平文で送信されます。ポートスキャンを停止するのではなく、野良ポートスキャンを停止します。
LAN 上のパケット スニファー。
Kickstart 47.115 / Workbench 47.5 (AmigaOS 3.2.3) を実行する A4000/060
ロードショーと 1920×1080 Picasso96 RTG スクリーン。そのマシンに対して確認済み:
stdout+stderr とリターンコードを含むシェル、ファイルの読み取り/書き込み、ディレクトリのリスト、
システム情報、 cybergraphics.library によるネイティブ トゥルーカラー スクリーン キャプチャ、
フォールバックとしての SGrab キャプチャ、マウスの移動/クリック、入力など -
äöü ÄÖÜ ß は、マシンのドイツ語キーマップを通じて正しく生成されました。
転送は約 830 KiB/s で実行されるため、完全な 1920×1080 トゥルーカラー グラブは 6.2 MB になります。
ワイヤー上では約 7 秒。クロードに届く PNG は約 850 KiB です。
一度に 1 つのコマンド。コマンドは期限付きで子プロセスで実行されます。
そのため、入力を待っているプログラムはエージェントを妨害しなくなり、他のすべてのことは行われなくなります。
動作し続け、amiga_break が Ctrl-C を送信します。しかし、コマンドがスタックすると、
以降の amiga_shell 呼び出しは、壊れるか終了するまでブロックされます。
クライアントとエージェントは同じバージョンである必要があります。ワイヤ形式が変更されました
0.3.0; 0.2.0 と 0.3.0 を混合すると、きれいな結果ではなく混乱を招くエラーが生成されます。
拒否。 --probe はエージェントのバージョンを出力します。
stderr には OS 3.2 以降が必要です。 SYS_Error は dos.library v47 で発生したため、3.2 では
その後、stderr がキャプチャされ、--- stderr --- マーカーの後に追加されます。
古いシステムは標準出力のみを取得します。
スクリーンショットは 2 つのルートのいずれかを使用します。エージェントは平面画面をキャプチャします
それ自体 (OS 3.0+、≤256 色) およびトゥルーカラー/RTG 画面を介して
サイバーグラフ

ics.library 、Picasso96 と Cyber​​GraphX の両方が提供します。もし
どちらにも適合しない場合は、Amiga 上で C:sgrab を実行して、
ファイル — SGrab は Amiga 上で圧縮するため、転送量ははるかに少なくなります。 SGrabのJPEG
モードにはさらに jpeg.library が必要です。 PNG にはありません。
一度に 1 つのコマンド、一度に 1 つの接続。
両方向で 16 MiB のフレーム キャップ。
PROTOCOL.md には、ワイヤ形式と、さらに便利な理由が記載されています。
各部分はそのように見えます - フレーム構成、各部分に 1 つのコマンドがある理由
接続、入力イベントの挿入方法、および System() ファイルハンドル
ハードウェア上で最初に見つかった実際のバグの原因となった所有権ルール。
testing/fake_agent.py は、サンドボックス ディレクトリに対して同じワイヤ プロトコルを通信します。
そのため、デスクトップ上でスタック全体を構築してデバッグしてから、スタックの 1 つを変更できます。
実際のハードウェアにアクセスするための環境変数:
python3 テスト/fake_agent.py --port 7846 --root /tmp/fakeamiga
AMIGA_HOST=127.0.0.1 python3 サーバー/amimcp.py --probe
これはテスト ダブルであり、エミュレータではありません。EXEC はいくつかの機能を理解します。
実際に AmigaDOS を実行するのではなく、AmigaDOS 形式のコマンド。
./run_tests.sh # 69 個のテスト、ハードウェアもクロスコンパイラーも必要ありません
レイアウト
PROTOCOL.md ワイヤ形式、およびそのように見える理由
Agent/amiagent.c Amiga デーモン
Python側と共有されるagent/proto.h定数
エージェント/Makefile m68k-amigaos-gcc ビルド
Agent/vendor/cgx/CyberGraphX インターフェイス ファイル (README を参照)
server/amimcp.py MCP サーバー (stdio JSON-RPC)
サーバー/amiga.py ワイヤー プロトコル クライアント
server/png.py chunky/RGB → PNG、stdlib zlib
テス

[切り捨てられた]

## Original Extract

Let Claude work on a real Amiga — an MCP server plus a tiny AmigaOS daemon. Shell, files, screen capture, mouse and keyboard. - thomas-luebker/amimcp

GitHub - thomas-luebker/amimcp: Let Claude work on a real Amiga — an MCP server plus a tiny AmigaOS daemon. Shell, files, screen capture, mouse and keyboard. · GitHub
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
thomas-luebker
/
amimcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits agent agent docs docs server server tests tests tools tools .gitignore .gitignore LICENSE LICENSE NOTICE NOTICE PROTOCOL.md PROTOCOL.md README.md README.md install.sh install.sh run_tests.sh run_tests.sh View all files Repository files navigation
amimcp — let Claude work on your Amiga
An MCP server that gives Claude hands on a
real Amiga. It can run AmigaDOS commands, read and write files, list drawers,
inspect the machine, capture the screen, and move the mouse and type on it.
You describe the problem; Claude pokes at the actual hardware.
That is Claude clicking into a Shell on an A4000/060, typing a command, and
pressing Return — the umlauts routed through the Amiga's own German keymap.
┌──────────────┐ MCP over stdio ┌──────────┐ framed TCP ┌────────────┐
│ Claude Code │ ◄─────────────────► │ amimcp │ ◄────────────► │ amiagent │
│ (your Mac/PC)│ JSON-RPC 2.0 │ (python) │ port 7846 │ (the Amiga)│
└──────────────┘ └──────────┘ └────────────┘
Two halves:
amimcp — the MCP server, on the machine running Claude. Pure Python
standard library: no venv, no pip, no lockfile to rot between the day you set
this up and the evening two years later when the Amiga won't boot.
amiagent — a small C daemon on the Amiga. AmigaOS 2.04 and up, bare
68000, ~85 KB, needs nothing but bsdsocket.library .
Tool
What it does
amiga_shell
Run an AmigaDOS command — stdout, stderr, and return code
amiga_read_file
Read a file; text inline, binaries base64
amiga_write_file
Write or overwrite a file, binaries included
amiga_list_dir
List a drawer with sizes, protection bits, datestamps
amiga_system_info
Kickstart, CPU/FPU, free Chip/Fast, volumes, assigns
amiga_screenshot
Capture a screen as PNG — planar or RTG, whole or a region
amiga_screens
List every open screen, front to back, with geometry
amiga_pointer
Where the pointer is — no image transferred
amiga_region_changed
Checksum a region: "has anything changed yet?"
amiga_click
Move the pointer and click — single, double, any button
amiga_drag
Press, move, release — drag-and-drop, scrollbars, menus
amiga_button
Hold or release a button on its own
amiga_move_mouse
Move the pointer without clicking
amiga_type
Type text, mapped through the Amiga's own keymap
amiga_key
Press Return, Esc, F-keys and so on, with qualifiers
amiga_break
Ctrl-C a command left running by amiga_shell
Together those are enough to actually work: read a Startup-Sequence and fix it,
cross-compile a binary and push it over and run it, launch a GUI program and
drive it by clicking, or just look at what the machine is showing when it has
gone wrong.
Two things are worth knowing before you drive a GUI:
Watch cheaply. A full 1080p truecolour frame is ~6 MB over a ~1 MB/s link —
about 7 seconds. Reading one status line as a region is effectively free, and
amiga_region_changed answers "has it finished drawing?" in four bytes. Poll
with hashes; capture pixels only when you need to look.
Two kinds of pointer. amiga_click warps the Intuition pointer, which
Intuition applications follow. SDL programs — games, ScummVM — track raw mouse
deltas and keep their own cursor, so they never see the warp and the click
lands wherever they still believe the pointer is. Pass relative: true for
those. And an Amiga menu is opened by holding the right button and moving,
so it needs amiga_drag , not a click.
1. Install the agent on the Amiga
Grab amiagent-0.5.3.lha
and unpack it on the Amiga. It contains amiagent (68000, runs on everything)
and amiagent.020 (68020+).
amipkg install amiagent
Or build it yourself
Needs an m68k-amigaos cross-toolchain —
AmigaPorts/m68k-amigaos-gcc ,
the maintained descendant of bebbo's amiga-gcc. (Bebbo's own repository is gone
as of 2026-08-07; AmigaPorts is where it lives now.) Cross-compile on any
machine that has it:
cd agent && make # 68000 baseline, runs on every Amiga
cd agent && make CPU=68020 # 68020+ build
Copy the resulting amiagent to the Amiga.
If the link fails with an undefined IntuitionBase , your toolchain doesn't
auto-open intuition/graphics. Rebuild with make OWNBASES=1 .
Bring up your TCP/IP stack (Roadshow, AmiTCP, or the emulator's bsdsocket
emulation), then from a Shell:
amiagent TOKEN=pickasecret
It prints that it is listening on port 7846 and waits. Ctrl-C stops it; if
you backgrounded it, find it with Status and Break <n> C .
Options: PORT/N (default 7846), TOKEN/K , QUIET/S .
To start it at boot, add this to S:User-Startup , after your TCP stack
comes up:
run >NIL: amiagent TOKEN=pickasecret QUIET
3. Point your assistant at it
./install.sh 192.168.1.42 pickasecret
That probes the Amiga first, then registers the server with Claude Code. Or do
it by hand:
claude mcp add amiga \
--env AMIGA_HOST=192.168.1.42 \
--env AMIGA_TOKEN=pickasecret \
-- python3 /path/to/amimcp/server/amimcp.py
For Claude Desktop , add this to claude_desktop_config.json :
{
"mcpServers" : {
"amiga" : {
"command" : " python3 " ,
"args" : [ " /path/to/amimcp/server/amimcp.py " ],
"env" : {
"AMIGA_HOST" : " 192.168.1.42 " ,
"AMIGA_TOKEN" : " pickasecret "
}
}
}
}
Start a new session and ask Claude to check the Amiga.
Other MCP clients — Codex, Cursor, Zed, Cline…
MCP is an open standard, and server/amimcp.py is an ordinary stdio MCP
server : it speaks JSON-RPC over stdin/stdout and takes its configuration from
environment variables. Anything that can launch an MCP server can therefore
drive the Amiga — the server does not know or care which model is on the other
end.
Only the config file shape differs. OpenAI Codex CLI ( ~/.codex/config.toml )
uses TOML:
[ mcp_servers . amiga ]
command = " python3 "
args = [ " /path/to/amimcp/server/amimcp.py " ]
env = { AMIGA_HOST = " 192.168.1.42 " , AMIGA_TOKEN = " pickasecret " }
Google Gemini CLI , Cursor , Zed , Cline , Continue ,
Windsurf , LM Studio and Goose take the same mcpServers JSON block
shown above for Claude Desktop, in their own settings file. Check your client's
current docs for the exact path — these move.
Tested against Claude Code and Claude Desktop. The rest follow from it being a
standard stdio server rather than from me having run all of them; if one needs
something different, please open an issue and say what.
Requires Python 3 — standard library only, no pip install , no venv.
Check the link without involving an LLM
AMIGA_HOST=192.168.1.42 AMIGA_TOKEN=pickasecret python3 server/amimcp.py --probe
This skips MCP entirely, so a failure here is a network or agent problem rather
than a client one.
Variable
Default
Meaning
AMIGA_HOST
(required)
The Amiga's IP address or hostname
AMIGA_PORT
7846
Port amiagent listens on
AMIGA_TOKEN
(none)
Must match the agent's TOKEN=
AMIGA_TIMEOUT
30
Default per-request timeout, seconds
Security — read this before opening the port
amiagent runs arbitrary commands for anyone who can reach its port, and the
protocol is unencrypted. That is a deliberate trade: classic Amigas have no TLS
without AmiSSL, and requiring AmiSSL just to reach the machine you are trying to
repair defeats the purpose.
So: set a TOKEN , keep it on a LAN you trust, and never forward the port.
Without a token, anything on your network segment can run Format on your
system partition. The agent warns loudly when started without one.
The token is compared byte-for-byte (so A4000 and a4000 are different
tokens) and travels in cleartext. It stops a stray port scan, not someone with a
packet sniffer on your LAN.
An A4000/060 running Kickstart 47.115 / Workbench 47.5 (AmigaOS 3.2.3) with
Roadshow and a 1920×1080 Picasso96 RTG screen. Confirmed against that machine:
shell with stdout+stderr and return codes, file read/write, directory listing,
system info, native truecolor screen capture through cybergraphics.library ,
SGrab capture as the fallback, mouse move/click, and typing — including
äöü ÄÖÜ ß correctly produced through the machine's German keymap.
Transfer runs at roughly 830 KiB/s, so a full 1920×1080 truecolor grab is 6.2 MB
and about 7 seconds on the wire; the PNG that reaches Claude is around 850 KiB.
One command at a time. Commands run in a child process with a deadline,
so a program waiting for input no longer wedges the agent — everything else
keeps working and amiga_break sends it Ctrl-C. But a stuck command does
block later amiga_shell calls until it is broken or finishes.
The client and agent must be the same version. The wire format changed in
0.3.0; mixing 0.2.0 with 0.3.0 produces confusing errors rather than a clean
refusal. --probe prints the agent's version.
stderr needs OS 3.2+. SYS_Error arrived in dos.library v47, so on 3.2
and later stderr is captured and appended after a --- stderr --- marker.
Older systems get stdout only.
Screenshots take one of two routes. The agent captures planar screens
itself (OS 3.0+, ≤256 colours) and truecolor/RTG screens through
cybergraphics.library , which Picasso96 and CyberGraphX both provide. If
neither fits, it falls back to running C:sgrab on the Amiga and fetching the
file — SGrab compresses on the Amiga, so it transfers far less. SGrab's JPEG
mode additionally needs jpeg.library ; PNG does not.
One command at a time , one connection at a time.
16 MiB frame cap in both directions.
PROTOCOL.md documents the wire format and, more usefully, why
each part looks the way it does — the framing, why there is one command per
connection, how input events are injected, and the System() file-handle
ownership rule that caused the first real bug found on hardware.
tests/fake_agent.py speaks the same wire protocol against a sandbox directory,
so you can build and debug the whole stack on your desktop and then change one
environment variable to hit real hardware:
python3 tests/fake_agent.py --port 7846 --root /tmp/fakeamiga
AMIGA_HOST=127.0.0.1 python3 server/amimcp.py --probe
It is a test double, not an emulator: EXEC understands a handful of
AmigaDOS-shaped commands rather than actually running AmigaDOS.
./run_tests.sh # 69 tests, no hardware and no cross-compiler needed
Layout
PROTOCOL.md the wire format, and why it looks like that
agent/amiagent.c the Amiga daemon
agent/proto.h constants shared with the Python side
agent/Makefile m68k-amigaos-gcc build
agent/vendor/cgx/ CyberGraphX interface files (see its README)
server/amimcp.py MCP server (stdio JSON-RPC)
server/amiga.py wire protocol client
server/png.py chunky/RGB → PNG, stdlib zlib
tes

[truncated]
