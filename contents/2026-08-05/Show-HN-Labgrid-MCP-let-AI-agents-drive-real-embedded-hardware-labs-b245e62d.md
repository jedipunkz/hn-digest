---
source: "https://github.com/onurcelep/labgrid-mcp"
hn_url: "https://news.ycombinator.com/item?id=49182016"
title: "Show HN: Labgrid-MCP – let AI agents drive real embedded hardware labs"
article_title: "GitHub - onurcelep/labgrid-mcp: MCP server exposing labgrid hardware-in-the-loop device operations to LLM agents — 47 tools over coordinator gRPC + labgrid's client-side driver stack (power/console/io/flash/mux/SSH). Standalone, coordinator-agnostic. · GitHub"
author: "oclp"
captured_at: "2026-08-05T12:48:09Z"
capture_tool: "hn-digest"
hn_id: 49182016
score: 2
comments: 0
posted_at: "2026-08-05T12:39:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Labgrid-MCP – let AI agents drive real embedded hardware labs

- HN: [49182016](https://news.ycombinator.com/item?id=49182016)
- Source: [github.com](https://github.com/onurcelep/labgrid-mcp)
- Score: 2
- Comments: 0
- Posted: 2026-08-05T12:39:57Z

## Translation

タイトル: Show HN: Labgrid-MCP – AI エージェントに実際の組み込みハードウェア ラボを駆動させる
記事のタイトル: GitHub - onurcelep/labgrid-mcp: labgrid のハードウェアインザループ デバイス操作を LLM エージェントに公開する MCP サーバー — コーディネーター gRPC 上の 47 ツール + labgrid のクライアント側ドライバー スタック (power/console/io/flash/mux/SSH)。スタンドアロン、コーディネーターに依存しない。 · GitHub
説明: MCP サーバーは、labgrid のハードウェアインザループ デバイス操作を LLM エージェントに公開します。コーディネーター gRPC + labgrid のクライアント側ドライバー スタック (power/console/io/flash/mux/SSH) を介した 47 ツール。スタンドアロン、コーディネーターに依存しない。 - onurcelep/labgrid-mcp

記事本文:
GitHub - onurcelep/labgrid-mcp: labgrid のハードウェアインザループ デバイス操作を LLM エージェントに公開する MCP サーバー — コーディネーター gRPC 上の 47 ツール + labgrid のクライアント側ドライバー スタック (power/console/io/flash/mux/SSH)。スタンドアロン、コーディネーターに依存しない。 · GitHub
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
オンルセレプ
/
ラボグリッド-mcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット .claude .claude 。

github/ workflows .github/ workflows docs docs src/ labgrid_mcp src/ labgrid_mcp テスト テスト .gitignore .gitignore .gitlint .gitlint .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md Dockerfile Dockerfile ライセンス ライセンス通知 README.md README.md Glama.json Glama.json pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード、AI エディター、および任意の MCP クライアントから実際の組み込みハードウェアを駆動します。
labgrid はオープンソースです
フレームワーク組み込みチームは、ラボ ハードウェア、ボード (「場所」) を共有するために使用します。
リモートで切り替え可能な電源、シリアル コンソール、USB マルチプレクサ、およびフラッシュ ツール。
labgrid-mcp はモデル コンテキスト プロトコルです
あらゆる labgrid ラボを MCP エコシステムに接続するサーバーなので、エージェントと開発者
ツールはラボと直接連携できます。
「rk3399 ボードを入手し、昨夜のイメージをフラッシュし、電源を入れ直し、
ログインプロンプトが表示されるかどうかを教えてください。コンソール ログがある場合は、それを貼り付けます。
そうしません。」
チャットだけでなく、スクリプトまたは人間主導のあらゆる MCP クライアントが、
labgrid の成熟したドライバー エコシステム上のポリシーゲート型リモート コントロール サーフェス、
予約と所有権の調停を行うアドホック デバイス サーバーでは、
持っています。
クロードが組み込みのデモ ラボを推進 — ハードウェアなし、コマンド 1 つ: uvx labgrid-mcp デモ
デバイスのライフサイクル全体: 検出、予約、取得、解放。
keepalive-backed なので、タスク中にホールドが期限切れになることはありません
ハードウェア制御: 電源オン/オフ/サイクル、デジタル I/O、SD/USB マルチプレクサ切り替え
インタラクティブ シリアル コンソール: 開く、読み取る、送信する、閉じる。リングバッファ型
デバイスへの SSH : コマンドの実行、ファイルの転送、両方向のトンネル
フラッシュ (オプトイン) : DFU、ファストブート、ブートストラップ ローダー、イメージ書き込み、
すべてステータス/ログポーリングを伴うバックグラウンドジョブとして

している
ラボのハウスキーピング : タグ、エイリアス、コメント、場所管理、変更
モニタリング
安全ゲート: 読み取り専用モードとカテゴリごとの許可リスト。の
不可逆的なファミリ(フラッシュ、場所の削除)はデフォルトでオフになっています
47 個のツール、5 個の参照可能な labgrid:// リソース、正直
すべてのツールの readOnly / 破壊的なアノテーション。
5 分で試してみましょう (ハードウェアは必要ありません)
uv が必要です (バンドルされている uvx が残りの作業を行います)
Python のプロビジョニングを含む):
uvx labgrid-mcp デモ
これにより、マシン上で完全な偽のラボ、つまり本物のラボグリッドが起動します。
コーディネーターとエクスポーター、偽の電源スイッチと偽の電源スイッチを備えた 1 つのデモ ボード
シリアルコンソール。貼り付け可能な .mcp.json スニペットが出力されます。それから尋ねてください
あなたのエージェント:
「場所をリストアップし、デモプレイスを取得する」
「電源デモ - オンにして電源状態を読み取ります」
「デモプレイスでコンソールを開いて、その出力を読んでください」
個別のインストール手順はありません - uvx は最初に PyPI から labgrid-mcp を取得します
実行時間。 ( pip をお勧めしますか? pip install labgrid-mcp を使用してください
"command": "labgrid-mcp" (以下に引数はありません)。
実行中の gRPC 時代の labgrid コーディネーターが必要です (labgrid ≥ 24、テスト済み)
26.x に対して）このマシンからアクセス可能です。
1. サーバーを MCP クライアントに登録します。サーバー定義は、
どこでも同じ — コマンド uvx 、 args ["labgrid-mcp"] 、および LG_* 環境
vars — 構成ファイルの場所とトップレベルのキーのみが異なります。あなたのものを選んでください
クライアント:
プロジェクト ルートに .mcp.json として保存します。
{
"mcpサーバー": {
"ラボグリッド" : {
"コマンド" : " uvx " ,
"args" : [ "labgrid-mcp " ],
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
または 1 つのコマンド: claude mcp add labgrid --env LG_COORDINATOR=your-coordinator-host:20408 -- uvx labgrid-mcp
「設定」→「開発者」→「構成の編集」を選択し、mcpServers の下に追加します。
クロード_デスクトップ_config.json :
{
"mcpサーバー": {
"ラボグリッド" : {
"コマンド" : " uvx " ,
"引数

s" : [ " labgrid-mcp " ]、
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
カーソル
.cursor/mcp.json (プロジェクト) または ~/.cursor/mcp.json (グローバル) として保存します。
{
"mcpサーバー": {
"ラボグリッド" : {
"コマンド" : " uvx " ,
"args" : [ "labgrid-mcp " ],
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
VS コード (副操縦士)
.vscode/mcp.json として保存 — VS Code はサーバー キーを使用することに注意してください。
{
"サーバー" : {
"ラボグリッド" : {
"コマンド" : " uvx " ,
"args" : [ "labgrid-mcp " ],
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
ウィンドサーフィン
~/.codeium/windsurf/mcp_config.json の mcpServers の下に追加します。
{
"mcpサーバー": {
"ラボグリッド" : {
"コマンド" : " uvx " ,
"args" : [ "labgrid-mcp " ],
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
他の MCP クライアント/SDK
labgrid-mcp は標準の stdio MCP サーバーです: spawn uvx labgrid-mcp (または
pip install labgrid-mcp 後の labgrid-mcp ) LG_COORDINATOR を設定
その環境に接続し、stdin/stdout 経由で MCP を話します。あらゆるクライアントと連携したり、
標準入出力サーバーをサポートするエージェント SDK。
2. クライアントを再起動して、新しいサーバーを選択します。
3. 接続されていることを確認します。エージェントに「labgrid の場所をリストしてください」と依頼します。あなた
研究室のボードを取り戻す必要があります。準備は完了です。
ID は labgrid-client : set LG_HOSTNAME / とまったく同じように機能します。
LG_USERNAME 、または実際のホスト名/ユーザーを使用する場合は省略します。セキュリティは
labgrid-client と同様に、ネットワーク (VPN / SSH トンネル) に委任されます。
(PyPI ではなくクローンから実行しますか? "command": "uv" を使用してください。
"args": ["run"、"--directory"、"/path/to/labgrid-mcp"、"labgrid-mcp"] 。)
サーバーをコンテナとして構築して実行します (ロックダウンされたホスト、またはラボ ネットワーク内でホスト)
リポジトリの Dockerfile からイメージをビルドします。
git clone https://github.com/onurcelep/labgrid-mcp && cd labgrid-mcp
ドッカービルド

ld -t labgrid-mcp 。
次に、MCP クライアントにそれを指定します。
{
"mcpサーバー": {
"ラボグリッド" : {
"コマンド" : " docker " ,
"args" : [ " run " 、 " -i " 、 " --rm " 、
" -e " 、 " LG_COORDINATOR=あなたのコーディネーターホスト:20408 " 、
" labgrid-mcp " ]
}
}
}
ラボに役立つパターン: 内部のホスト上でコンテナを構築して実行する
MCP クライアントがどこでも実行されている間、ラボ ネットワーク、1 つにつき 1 つのコンテナ
ユーザーなので、アイデンティティと所有権は個人ごとに維持されます。
{
"mcpサーバー": {
"ラボグリッド" : {
"コマンド" : " ssh " 、
"args" : [ " labhost " 、 " docker " 、 " run " 、 " -i " 、 " --rm " 、
" -e " 、 " LG_COORDINATOR=127.0.0.1:20408 " 、
" -e " 、 " LG_USERNAME=あなたの名前 " 、
" labgrid-mcp " ]
}
}
}
ユーザー間で 1 つの実行サーバーを共有しないでください。各インスタンスは 1 つのサーバーを保持します。
labgrid ID を使用するため、共有インスタンスにより全員が取得できるようになります
見分けがつかない。ユーザー/エージェントごとに 1 つのコンテナがラボの所有権を保持します
モデルはそのまま。 (注: ビルドするイメージには labgrid がバンドルされます。
LGPL-2.1 以降 — どこでも使用できます。画像を再配布する場合は、
LGPL の条件はそのコピーに適用され、labgrid のライセンス テキストがすでに記載されています
その中にあります。）
わかりやすい言葉で尋ねるだけで、エージェントが適切なツールにマッピングします。典型的な
最初のワークフロー:
「今空いている場所はどこですか？」
「電源を入れて、シリアル コンソールを開いてブート出力を表示します。」
「電源を切ってボードを放してください。」
読み取り専用の質問 (「場所のリスト」、「ボード 7 を持っているのは誰ですか?」) はすぐに機能します。
ハードウェアの状態を変更するものはすべてゲートされます (以下を参照)。
不可逆的なファミリー - フラッシュと場所の削除 - までオフのままにしてください
明示的に有効にします。
環境変数
デフォルト
効果
LG_COORDINATOR
127.0.0.1:20408
コーディネーターの住所
LG_HOSTNAME / LG_USERNAME
実ホスト/ユーザー
ID (labgrid-client など)
LABGRID_MCP_READONLY
オフ
1 = 読み取り専用ツールのみが登録されます。

読み取りグループと wait_for_change および forward_list
LABGRID_MCP_ALLOW
設定を解除する
登録するカテゴリのカンマリスト。 flash と place_delete は明示的にリストする必要があります。デフォルトでもオフになっています
LABGRID_MCP_SSH_KEYFILE
設定を解除する
SSH ツールの秘密キー。設定されていない場合、呼び出し時に明らかにエラーが発生します
LABGRID_MCP_ACQUIRE_TIMEOUT
120
acquire_place が割り当てを待機する最大秒数
安全性を一言で言えば、フラッシュと場所の削除は元に戻せない可能性があります
したがって、それぞれに独自の明示的な LABGRID_MCP_ALLOW エントリが必要です。 SSHツール
取得したボード上での任意のコマンド実行、同じ信頼クラス
コンソールセッションとして。 LABGRID_MCP_READONLY=1 は、それらをすべてのファイルとともに削除します
他のゲート ツール (メモリ内のトンネルのみをリストするため、forward_list は残ります)
状態）。そして、labgrid について知っておく価値のある警告: コーディネーターは強制的に禁止を強制します。
場所メタデータの所有権保護: このサーバーは取得したメタデータの編集を拒否します。
力のない場所 = True ですが、誰も保持していない場所を守ることはできません
(そして、set_place_tags の空のタグ値はそのキーを削除します。
labgrid 独自のセマンティクス)。詳細: docs/DESIGN.md §4 および §11.12。
グループ別の全 47 ツール
グループ
ツール
読む
coordinator_info 、 list_places 、 show_place 、 who 、 list_resources 、 list_reservations
取得・予約
acquire_place、release_place、allow_place、release_from、reserve、cancel_reservation、reservation_wait
ドライバー
get_power_state 、 set_power 、 get_io 、 set_io 、 get_sd_mux 、 set_sd_mux 、 set_usb_mux
コンソール
console_open 、 console_read 、 console_send 、 console_close
SSH/転送
ssh_run 、 put_file 、 get_file 、 forward_open 、 forward_remote_open 、 forward_close 、 forward_list
フラッシュ (オプトイン)
flash_dfu 、 flash_fastboot 、 flash_script 、 bootstrap 、 write_image 、 flash_status 、 flash_logs
メタデータを配置する
add_place 、 add_place_alias 、 delete_place_

alias 、 set_place_tags 、 set_place_comment 、 add_place_match
場所の削除（オプトイン）
delete_place 、 delete_place_match
変更監視
変化を待つ
リソース: labgrid://places 、labgrid://places/{name} 、
labgrid://resources 、labgrid://reservations 、labgrid://sessions 。
ツールごとの引数と動作は、各ツール独自のドキュメントに文書化されています。
説明 (MCP クライアントに表示されます) と
docs/DESIGN.md §5/§11。
ビデオ/オーディオ/スクリーン キャプチャまたは USB 機器なし (gstreamer + が必要)
物理USB。正常な MCP 表面はありません)
ライブ イベント ストリームはありません。代わりに wait_for_change ロングポーリング
古いクロスバー コーディネーター (labgrid < 24) は接続できません
認証はネットワークレベル (VPN/トンネル) であり、まさに labgrid 独自のモデルです
実際のフラッシュ/マルチプレクサ ドライバーのステップには実際のボードが必要です。ジョブの機構は次のとおりです。
偽物に対して完全に CI テストされており、シリコンに触れるステップは行われていません。
labgrid-mcp 自体の作業
統合スイートは、デモを含むスタック全体を実行します。
すべての PR の CI で、偽のハードウェアを使用した本物のコーディネーター/エクスポーター プロセス、
さらに、labgrid master に対する毎週のカナリア:
git clone < this-repo > labgrid-mcp && cd labgrid-mcp
UV同期
uv 実行 pytest # ユニット
uv run pytest -m 統合
アーキテクチャ、意思決定ログ、および labgrid の検証済みリファレンス
整数

[切り捨てられた]

## Original Extract

MCP server exposing labgrid hardware-in-the-loop device operations to LLM agents — 47 tools over coordinator gRPC + labgrid's client-side driver stack (power/console/io/flash/mux/SSH). Standalone, coordinator-agnostic. - onurcelep/labgrid-mcp

GitHub - onurcelep/labgrid-mcp: MCP server exposing labgrid hardware-in-the-loop device operations to LLM agents — 47 tools over coordinator gRPC + labgrid's client-side driver stack (power/console/io/flash/mux/SSH). Standalone, coordinator-agnostic. · GitHub
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
onurcelep
/
labgrid-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits .claude .claude .github/ workflows .github/ workflows docs docs src/ labgrid_mcp src/ labgrid_mcp tests tests .gitignore .gitignore .gitlint .gitlint .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md glama.json glama.json pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Drive real embedded hardware from Claude, AI editors, and any MCP client.
labgrid is the open-source
framework embedded teams use to share lab hardware: boards ("places") with
remotely switchable power, serial consoles, USB muxes, and flashing tools.
labgrid-mcp is a Model Context Protocol
server that plugs any labgrid lab into the MCP ecosystem, so agents and dev
tools can work with the lab directly:
"Acquire the rk3399 board, flash last night's image, power-cycle it, and
tell me whether it reaches a login prompt. Paste the console log if it
doesn't."
Not only for chat: any MCP client, scripted or human-driven, gets a
policy-gated remote-control surface over labgrid's mature driver ecosystem,
with the reservations and ownership arbitration ad-hoc device servers don't
have.
Claude driving the built-in demo lab — no hardware, one command: uvx labgrid-mcp demo
Full device lifecycle : discover, reserve, acquire, release;
keepalive-backed so holds never expire mid-task
Hardware control : power on/off/cycle, digital I/O, SD/USB mux switching
Interactive serial console : open, read, send, close; ring-buffered
SSH to the device : run commands, transfer files, tunnels in both directions
Flashing (opt-in) : DFU, fastboot, bootstrap loaders, image writing,
all as background jobs with status/log polling
Lab housekeeping : tags, aliases, comments, place management, change
monitoring
Safety gating : read-only mode and per-category allowlists; the
irreversible families (flash, place deletion) are off by default
47 tools, 5 browseable labgrid:// resources, honest
readOnly / destructive annotations on every tool.
Try it in 5 minutes (no hardware needed)
Requires uv (its bundled uvx does the rest,
including provisioning Python):
uvx labgrid-mcp demo
This boots a complete fake lab on your machine: a real labgrid
coordinator and exporter, one demo board with a fake power switch and a fake
serial console. It prints a paste-ready .mcp.json snippet. Then ask
your agent:
"List places, then acquire demo-place"
"Power demo-place on and read its power state"
"Open the console on demo-place and read its output"
No separate install step — uvx fetches labgrid-mcp from PyPI the first
time it runs. (Prefer pip? pip install labgrid-mcp , then use
"command": "labgrid-mcp" with no args below.)
You need a running, gRPC-era labgrid coordinator (labgrid ≥ 24; tested
against 26.x) reachable from this machine.
1. Register the server with your MCP client. The server definition is the
same everywhere — command uvx , args ["labgrid-mcp"] , plus your LG_* env
vars — only the config file location and top-level key differ. Pick your
client:
Save as .mcp.json in your project root:
{
"mcpServers" : {
"labgrid" : {
"command" : " uvx " ,
"args" : [ " labgrid-mcp " ],
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
or one command: claude mcp add labgrid --env LG_COORDINATOR=your-coordinator-host:20408 -- uvx labgrid-mcp
Settings → Developer → Edit Config, then add under mcpServers in
claude_desktop_config.json :
{
"mcpServers" : {
"labgrid" : {
"command" : " uvx " ,
"args" : [ " labgrid-mcp " ],
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
Cursor
Save as .cursor/mcp.json (project) or ~/.cursor/mcp.json (global):
{
"mcpServers" : {
"labgrid" : {
"command" : " uvx " ,
"args" : [ " labgrid-mcp " ],
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
VS Code (Copilot)
Save as .vscode/mcp.json — note VS Code uses a servers key:
{
"servers" : {
"labgrid" : {
"command" : " uvx " ,
"args" : [ " labgrid-mcp " ],
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
Windsurf
Add under mcpServers in ~/.codeium/windsurf/mcp_config.json :
{
"mcpServers" : {
"labgrid" : {
"command" : " uvx " ,
"args" : [ " labgrid-mcp " ],
"env" : { "LG_COORDINATOR" : " your-coordinator-host:20408 " }
}
}
}
Any other MCP client / SDK
labgrid-mcp is a standard stdio MCP server: spawn uvx labgrid-mcp (or
labgrid-mcp after pip install labgrid-mcp ) with LG_COORDINATOR set in
its environment, and speak MCP over stdin/stdout. Works with any client or
agent SDK that supports stdio servers.
2. Restart the client so it picks up the new server.
3. Confirm it's connected — ask your agent "List the labgrid places" ; you
should get your lab's boards back. You're ready.
Identity works exactly like labgrid-client : set LG_HOSTNAME /
LG_USERNAME , or omit them to use your real hostname/user. Security is
delegated to the network (VPN / SSH tunnel), same as labgrid-client .
(Running from a clone instead of PyPI? Use "command": "uv" ,
"args": ["run", "--directory", "/path/to/labgrid-mcp", "labgrid-mcp"] .)
Build and run the server as a container (locked-down hosts, or hosting it inside the lab network)
Build the image from the repo's Dockerfile :
git clone https://github.com/onurcelep/labgrid-mcp && cd labgrid-mcp
docker build -t labgrid-mcp .
Then point your MCP client at it:
{
"mcpServers" : {
"labgrid" : {
"command" : " docker " ,
"args" : [ " run " , " -i " , " --rm " ,
" -e " , " LG_COORDINATOR=your-coordinator-host:20408 " ,
" labgrid-mcp " ]
}
}
}
A useful pattern for labs: build and run the container on a host inside
the lab network while your MCP client runs anywhere, one container per
user so identity and ownership stay per-person:
{
"mcpServers" : {
"labgrid" : {
"command" : " ssh " ,
"args" : [ " labhost " , " docker " , " run " , " -i " , " --rm " ,
" -e " , " LG_COORDINATOR=127.0.0.1:20408 " ,
" -e " , " LG_USERNAME=your-name " ,
" labgrid-mcp " ]
}
}
}
Don't share one running server between users: each instance holds a single
labgrid identity, so a shared instance would make everyone's acquisitions
indistinguishable. One container per user/agent keeps the lab's ownership
model intact. (Note: an image you build bundles labgrid,
LGPL-2.1-or-later — fine to use anywhere; if you redistribute the image,
the LGPL's terms apply to that copy, with labgrid's license texts already
inside it.)
Just ask in plain language — the agent maps it to the right tools. A typical
first workflow:
"Which places are free right now?"
"Power it on, then open the serial console and show me the boot output."
"Power it off and release the board."
Read-only asks ("list places", "who's holding board-7?") work immediately.
Anything that changes hardware state is gated (see below), and the two
irreversible families — flashing and place deletion — stay off until
you explicitly enable them.
Env var
Default
Effect
LG_COORDINATOR
127.0.0.1:20408
Coordinator address
LG_HOSTNAME / LG_USERNAME
real host/user
Identity, as in labgrid-client
LABGRID_MCP_READONLY
off
1 = only read-only tools are registered: the Read group plus wait_for_change and forward_list
LABGRID_MCP_ALLOW
unset
Comma list of categories to register; flash and place_delete must be listed explicitly ; they're off even by default
LABGRID_MCP_SSH_KEYFILE
unset
Private key for the SSH tools; unset, they error clearly at call time
LABGRID_MCP_ACQUIRE_TIMEOUT
120
Max seconds acquire_place waits for allocation
Safety in one paragraph: flashing and place deletion can do irreversible
damage, so each needs its own explicit LABGRID_MCP_ALLOW entry. SSH tools
are arbitrary command execution on the acquired board, the same trust class
as a console session; LABGRID_MCP_READONLY=1 drops them along with every
other gated tool ( forward_list stays, since it only lists in-memory tunnel
state). And a labgrid caveat worth knowing: the coordinator enforces no
ownership guard on place metadata: this server refuses to edit an acquired
place without force=True , but nothing can protect a place nobody holds
(and an empty tag value in set_place_tags deletes that key, which is
labgrid's own semantics). Details: docs/DESIGN.md §4 and §11.12.
All 47 tools by group
Group
Tools
Read
coordinator_info , list_places , show_place , who , list_resources , list_reservations
Acquisition / reservation
acquire_place , release_place , allow_place , release_from , reserve , cancel_reservation , reservation_wait
Drivers
get_power_state , set_power , get_io , set_io , get_sd_mux , set_sd_mux , set_usb_mux
Console
console_open , console_read , console_send , console_close
SSH / forward
ssh_run , put_file , get_file , forward_open , forward_remote_open , forward_close , forward_list
Flash (opt-in)
flash_dfu , flash_fastboot , flash_script , bootstrap , write_image , flash_status , flash_logs
Place metadata
add_place , add_place_alias , delete_place_alias , set_place_tags , set_place_comment , add_place_match
Place deletion (opt-in)
delete_place , delete_place_match
Change monitoring
wait_for_change
Resources: labgrid://places , labgrid://places/{name} ,
labgrid://resources , labgrid://reservations , labgrid://sessions .
Per-tool arguments and behaviors are documented in each tool's own
description (visible in your MCP client) and in
docs/DESIGN.md §5/§11.
No video/audio/screen capture or USB instruments (need gstreamer +
physical USB; no sane MCP surface)
No live event stream; wait_for_change long-polling instead
Old crossbar coordinators (labgrid < 24) can't connect
Authentication is network-level (VPN/tunnel), exactly labgrid's own model
The real flash/mux driver step needs a real board: the job machinery is
fully CI-tested against fakes, the silicon-touching step is not
Working on labgrid-mcp itself
The integration suite runs the whole stack, including the demo, against
real coordinator/exporter processes with fake hardware, in CI on every PR,
plus a weekly canary against labgrid master :
git clone < this-repo > labgrid-mcp && cd labgrid-mcp
uv sync
uv run pytest # unit
uv run pytest -m integration
Architecture, decision log, and a verified reference of labgrid's
int

[truncated]
