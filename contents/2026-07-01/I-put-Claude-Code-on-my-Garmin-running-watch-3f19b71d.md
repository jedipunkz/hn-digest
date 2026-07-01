---
source: "https://github.com/fashton28/garmin-code"
hn_url: "https://news.ycombinator.com/item?id=48754085"
title: "I put Claude Code on my Garmin running watch"
article_title: "GitHub - fashton28/garmin-code: run and manage your Claude Code agents directly from your Garmin Watch · GitHub"
author: "fashton28"
captured_at: "2026-07-01T23:28:53Z"
capture_tool: "hn-digest"
hn_id: 48754085
score: 2
comments: 2
posted_at: "2026-07-01T22:39:37Z"
tags:
  - hacker-news
  - translated
---

# I put Claude Code on my Garmin running watch

- HN: [48754085](https://news.ycombinator.com/item?id=48754085)
- Source: [github.com](https://github.com/fashton28/garmin-code)
- Score: 2
- Comments: 2
- Posted: 2026-07-01T22:39:37Z

## Translation

タイトル: Garmin ランニングウォッチにクロード コードを装着しました
記事のタイトル: GitHub - fashton28/garmin-code: Garmin Watch から直接クロード コード エージェントを実行および管理 · GitHub
説明: Garmin Watch から直接 Claude Code エージェントを実行および管理します - fashton28/garmin-code

記事本文:
GitHub - fashton28/garmin-code: Garmin Watch から直接クロード コード エージェントを実行および管理 · GitHub
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
ファシュトン28
/
ガーミンコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
46 コミット 46 コミット .github/ workflows .github/ workflows apps apps Contract Contract docs docs tools tools .gitignore .gitignore README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Garmin Forerunner 165 からクロード コード セッションを監視し、実行します。
Mac 上の小さなデーモンがローカルのクロード コード セッション履歴を読み取り、認証された HTTPS API 経由で公開します。
Connect IQ ウォッチ アプリはその API を取得し、手首から次のことを可能にします。
最近のセッションをライブ状態のドット (動作中 / 待機中 / アイドル)、AI 生成のタイトル、相対的な最終アクティブ時間、およびセッションが使用するモデルで確認します。
ライブ ステータス画面 (スピナー、モデル、結果メッセージ) を使用して、セッション内のタスク (PR の作成、テストの実行、またはコードのレビュー) をヘッドレスで実行します。
全体的な使用状況を表示します - 合計トークン、セッション、メッセージ、およびモデルごとの内訳のダッシュボード。
カスタム 0xProto 等幅フォント、オレンジ色のフォーカス アクセント、クロード スパークなど、すべてが端末/「コンピューター」の美学でスタイル設定されています。
ビデオがインラインで再生されない場合は、ここで視聴してください。
フォアランナー 165 (Monkey C アプリ)
-> makeWebRequest (HTTPS + ベアラー トークン)、BLE 経由で電話経由で中継
-> 携帯電話の Garmin Connect アプリ (インターネット ブリッジ)
-> パブリック HTTPS URL (Cloudflare トンネル / Tailscale Funnel)
-> Mac 上のデーモン (Node + TypeScript)
-> ~/.claude/projects/*/*.jsonl を読み取り、タスクに対して `claude -p` を実行します
2 つの事実が全体の設計を推進します。
FR165 には WiFi も LTE もありません。
すべてのネットワーク リクエストは、Bluetooth 経由で携帯電話の Garmin Connect アプリを介してトンネリングされ、インターネットに送信されます。
したがって、携帯電話が近くにあり、接続されており、Garmin Connect を実行している必要があります。
Th

このデーモンは、クロード コードの作業が実際に存在するマシン上で実行する必要があります。
ローカル ~/.claude からセッションのトランスクリプトを読み取り、ローカル リポジトリに claude -p を生成することでタスクを実行します。
汎用クラウド ボックスには独自のセッションとリポジトリのみが表示されるため、このボックスにオフロードすることはできません。
macOS (またはクロード コードを実行する別のマシン) と Node 20+ 。
Claude Code CLI ( claude ) がインストールされ、認証されています - タスク機能に必要です。
Garmin Connect IQ SDK と開発者キー ( ~/development/developer_key )、ウォッチ アプリを構築します。
Garmin Forerunner 165 とそのデータ ケーブル (実デバイス用)。
Garmin Connect アプリがインストールされたスマートフォンと時計がペアリングされています。
パブリック HTTPS 経由でデーモンを公開する方法 (Cloudflare トンネルまたは Tailscale Funnel - 「デーモンの公開」を参照)。
動作を確認する最速の方法は、完全に Mac 上で行われ、電話やトンネルは必要ありません。
git clone https://github.com/fashton28/garmin-code
cd ガーミンコード
npm ci # デーモンワークスペースをインストールします
npm run sim # デーモン + シミュレーター + アプリ、すべて実際のセッションに接続
npm run sim は、実際の ~/.claude セッションを提供するローカル デーモンを開始し、ウォッチ アプリを http://localhost:8787 に指定してビルドし、Connect IQ シミュレーターを起動して、アプリを読み込みます。
Ctrl-C を押してすべてを破棄します ( Config.mc が復元されます)。
これがトンネルなしで機能する理由: シミュレーターは Mac のネットワークから直接 Web リクエストを行うため、プレーンな http://localhost で問題ありません。
実際の時計は異なります（下記を参照）。
本物の ForeAthlete 165 に装着する
パブリック HTTPS エンドポイント、構成、ビルド、サイドロードの 4 つのステップがあります。
1. パブリック HTTPS 経由でデーモンを公開する
実際の監視には、有効な証明書を持つパブリック HTTPS URL が必要です (プレーンな localhost はシミュレーターでのみ機能します)。
1 つ選択してください:
醸造インストールクラウドフレア
c

ラウドフレア トンネル --url http://localhost:8787 # https://<ランダム>.trycloudflare.com URL を出力します
このプロセスを実行し続けます。ウォッチは、デーモンが起動している間のみデーモンに到達します。
2. Watch アプリをそれに向けます (そしてトークンを共有します)
ウォッチとデーモンは共有ベアラー トークンを使用して認証します。
1 つを生成し、デーモンの環境に配置し、同じトークンとトンネル URL を監視設定に配置します。
# デーモン側 (git 無視)
printf ' CLAUDEWATCH_TOKEN=%s\nPORT=8787\n ' " $( openssl rand -hex 32 ) " > apps/daemon/.env
次に、apps/watch/source/Config.mc を編集します (git-ignored - 見つからない場合は Config.mc.example からコピーします)。
const BASE_URL as Lang.String = "https://<トンネルのホスト名>"; // 末尾のスラッシュはありません
const TOKEN as Lang.String = "<apps/daemon/.env と同じトークン>";
const LIMIT as Lang.Number = 10;
Config.mc のトークンは apps/daemon/.env の CLAUDEWATCH_TOKEN と一致する必要があります。そうでない場合は、すべてのリクエストが 401 を返します。
# デーモンを実行します (apps/daemon/.env からトークンを読み取ります)
npm 実行サーブ
# ...そして別の端末のトンネル (ステップ 1)
# 署名された .prg をビルドする
npm run build:watch # -> apps/watch/bin/ClaudeWatch.prg
ビルドする前に、パブリック URL がエンドツーエンドで到達可能であることを確認します。
curl -s -o /dev/null -w ' %{http_code}\n ' -H " 認可: Bearer <トークン> " " https://<your-tunnel-hostname>/sessions?limit=1 " # 200 を期待します
4. 時計へのサイドロード
.prg は、時計が GARMIN/APPS/ から読み取る単なるファイルです。
macOS の落とし穴: FR165 は USB 大容量ストレージではなく MTP 経由で接続するため、 /Volumes にドライブとしてマウントされず、macOS ターミナル/Finder からコピーすることはできません。
最も簡単な作業パスを選択してください。
Windows (最も簡単): ウォッチを接続し、ファイル エクスプローラー -> この PC -> Forerunner 165 -> 内部ストレージ -> GARMIN -> APPS を開き、そこに ClaudeWatch.prg を貼り付けます。
openMTP 経由の macOS : 無料の MTP

Garmin を適切に処理する pp (Android File Transfer は Garmin では信頼できません)。ウォッチ -> GARMIN/APPS を参照し、.prg をドロップします。
ClaudeWatch.prg を GARMIN/APPS/ (サブフォルダーではなく) に直接置き、きれいに取り出して、ウォッチのアプリ/グランス リストから ClaudeWatch を開きます。
リストにない場合は、ウォッチを再起動します (起動時に GARMIN/APPS が再スキャンされます)。
Mac 上でデーモン + トンネルを実行し続け、Mac を起動してオンラインに保ちます。
スマートフォンを近くに置いて、Bluetooth をオンにし、Garmin Connect を実行してください。
クロードウォッチ を開きます。
セッションリストがロードされます。
セッションからタスクを実行できます。 [使用状況] 行 ([更新] の下) でダッシュボードが開きます。
携帯電話にインターネットがあり、Mac がデーモンとトンネルを有効にして稼働している限り、外出先でも使用できます。
セッション状態 (精度はオプション)
各セッションには、動作中 (緑色)、待機中 (黄色)、またはアイドル状態 (灰色) の状態が表示されます。
デフォルトでは、これはトランスクリプトの鮮度と最後のターンからのヒューリスティックであり、適切ではありますがおおよその値です (長時間のサイレント操作はアイドル状態と見なされる可能性があります)。
正確なリアルタイムの状態を得るには、docs/session-state.md で説明されているオプションのクロード コード フックを有効にします。フックは、動作中/待機中を発生時に報告し、ヒューリスティックをオーバーライドします。
[PR の作成] / [テストの実行] / [コードのレビュー] を選択すると、固定プロンプト、スコープ指定された --allowedTools ホワイトリスト、および --permission-mode acceptEdits を使用して、そのセッションの作業ディレクトリにヘッドレスで claude -p が生成されます (したがって、対話型プロンプトはこれをブロックできません)。
これらはマシン上の実際のアクションです。 PR の作成は実際に PR をコミット、プッシュ、オープンします。 「テストの実行」では、テスト コマンドを実行します。
エンドポイントはパブリック トンネルの背後にあるため、ベアラー トークンは真の意味で負荷を負います。
タスクごとのツールの許可リストは、爆発範囲を厳密に保ちます (例: Create PR は git / gh のみを取得します)。
レビューコードは読み取り専用なので、最初に試すのが最も安全です。
デー

mon は、Authorization: Bearer <token> によって保護された、小規模な凍結されたコントラクト (contract/contract.md) を提供します。
ノードcontract/validate.mjsを使用して、いつでも正規フィクスチャを検証します。
どちらの設定ファイルも git 無視されます。本当の秘密については何も犯さないでください。
apps/daemon/.env (デーモンによって自動的にロードされます):
apps/watch/source/Config.mc : BASE_URL 、 TOKEN (デーモンと一致する必要があります)、 LIMIT 。
監視ビルド ( tools/build-watch.sh ) は、インストールされている最新の SDK を自動検出します。パスが異なる場合は、 CIQ_SDK_BIN 、 CIQ_DEV_KEY 、 CIQ_DEVICE でオーバーライドします。
常時オンにする (オプション)
現在のセットアップはターミナルで実行され、セッションが終了するか Mac がスリープすると終了します。
耐久性が高く、手間のかからないセットアップを実現するには:
launchd LaunchAgent により、デーモンが起動時に自動起動し、クラッシュ時に再起動します。
caffeinate (または pmset ) を実行すると、Mac がスリープせずにトンネルが切断されます。
名前付きの Cloudflare トンネル (ドメインが必要) または Tailscale Funnel (無料) を介した安定した URL なので、URL 変更のために .prg を再度再構築する必要はありません。
Mac にはセッションとリポジトリが存在するため、Mac に電源が入っていてオンラインになっている必要があります。
時計に「読み込めません」/「電話が接続されていません」と表示されます。
電話ブリッジは稼働していません。
携帯電話が近くにあり、Bluetooth がオンになっていて、Garmin Connect アプリが実行中であり、携帯電話がインターネットに接続されていることを確認してください。
エラー画面で「開始」を押して再試行してください。
すべてが 401 を返します。
apps/watch/source/Config.mc のトークンが apps/daemon/.env の CLAUDEWATCH_TOKEN と一致しません。
トークンを修正し、 .prg を再構築し、再サイドロードします。
ウォッチは macOS 上でドライブとしてマウントされません。
FR165 は大容量ストレージではなく MTP を使用します - これは予想通りです。
Windows PC (ファイル エクスプローラーのネイティブ MTP) または Mac 上の openMTP を使用します。
Mac には「BillBoard Device」のみが表示され、Garmin は表示されません。
USB接続は充電専用です。
Garmin データ ケーブルを使用し、USB-A を経由する場合は、

USB-C アダプター。Mac に直接接続されたデータ対応のものを使用します。
アプリはウォッチのアプリリストにありません。
ウォッチを再起動して GARMIN/APPS を再スキャンし、 ClaudeWatch.prg が GARMIN/APPS/ 内に直接存在することを確認します。
しばらくすると、Watch アプリがデーモンに到達しなくなります。
Cloudflared が再起動すると (Mac がスリープし、ターミナルが閉じたとき)、クイック トンネル URL がローテーションします。
トンネルを再起動し、 BASE_URL を更新し、再構築し、再サイドロードするか、安定した URL の名前付きトンネル/Tailscale Funnel に移動します。
セッションにはモデルが表示されません。
実際の（合成ではない）アシスタントのターンがまだない新しいセッションには、記録されたモデルがありません。セッションが実行されると入力されます。
タスクはすぐに失敗します。
クロードがデーモンの PATH 上にあり認証されていること、およびタスクの cwd が実際のリポジトリであることを確認します (Create PR の場合は gh の認証も必要です)。
パス
それは何ですか
契約/
凍結された HTTP+JSON コントラクトと正規の session.json フィクスチャは、両側で構築されます。
アプリ/デーモン/
Node + TypeScript デーモン (唯一の npm ワークスペース): セッションを読み取り、API を提供し、タスクを実行します。
アプリ/ウォッチ/
Forerunner 165 の IQ (Monkey C) アプリを接続します。
ツール/トンネル/
Cloudflareのトンネル構成+スクリプトの実行( npm runserve:tunnel )。
tools/dev/run-sim.sh
npmの実行

[切り捨てられた]

## Original Extract

run and manage your Claude Code agents directly from your Garmin Watch - fashton28/garmin-code

GitHub - fashton28/garmin-code: run and manage your Claude Code agents directly from your Garmin Watch · GitHub
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
fashton28
/
garmin-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
46 Commits 46 Commits .github/ workflows .github/ workflows apps apps contract contract docs docs tools tools .gitignore .gitignore README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Monitor and drive your Claude Code sessions from a Garmin Forerunner 165.
A small daemon on your Mac reads your local Claude Code session history and exposes it over an authenticated HTTPS API.
A Connect IQ watch app fetches that API and lets you, from your wrist:
See your recent sessions with a live state dot (working / waiting for you / idle), the AI-generated title, relative last-active time, and the model the session uses.
Run a task in a session - Create PR, Run tests, or Review code - headlessly, with a live status screen (spinner, model, and an outcome message).
View overall usage - a dashboard of total tokens, sessions, messages, and a per-model breakdown.
Everything is styled in a terminal / "computer" aesthetic: a custom 0xProto monospace font, an orange focus accent, and a Claude spark.
If the video does not play inline, watch it here .
Forerunner 165 (Monkey C app)
-> makeWebRequest (HTTPS + bearer token), relayed through the phone over BLE
-> the Garmin Connect app on your phone (the internet bridge)
-> a public HTTPS URL (Cloudflare tunnel / Tailscale Funnel)
-> daemon on your Mac (Node + TypeScript)
-> reads ~/.claude/projects/*/*.jsonl and runs `claude -p` for tasks
Two facts drive the whole design.
The FR165 has no WiFi and no LTE .
Every network request is tunneled through the Garmin Connect app on your phone over Bluetooth, then out to the internet.
So the phone must be nearby, connected, and running Garmin Connect.
The daemon must run on the machine where your Claude Code work actually lives .
It reads session transcripts from your local ~/.claude and runs tasks by spawning claude -p in your local repos.
It cannot be offloaded to a generic cloud box, because that box would only see its own sessions and repos.
macOS (or another machine where you run Claude Code) with Node 20+ .
The Claude Code CLI ( claude ) installed and authenticated - required for the task feature.
Garmin Connect IQ SDK and a developer key ( ~/development/developer_key ), to build the watch app.
A Garmin Forerunner 165 and its data cable , for real-device use.
Your phone with the Garmin Connect app, paired to the watch.
A way to expose the daemon over public HTTPS (a Cloudflare tunnel or Tailscale Funnel - see Exposing the daemon ).
The fastest way to see it working, entirely on your Mac, no phone or tunnel needed.
git clone https://github.com/fashton28/garmin-code
cd garmin-code
npm ci # installs the daemon workspace
npm run sim # daemon + simulator + app, all wired to your real sessions
npm run sim starts a local daemon serving your real ~/.claude sessions, points the watch app at http://localhost:8787 , builds it, launches the Connect IQ simulator, and loads the app.
Press Ctrl-C to tear it all down (it restores your Config.mc ).
Why this works without a tunnel: the simulator makes web requests directly from your Mac's network , so plain http://localhost is fine.
The real watch is different (see below).
Getting it on the real Forerunner 165
There are four steps: a public HTTPS endpoint, the config, the build, and the sideload.
1. Expose the daemon over public HTTPS
The real watch needs a public HTTPS URL with a valid certificate (plain localhost only works in the simulator).
Pick one:
brew install cloudflared
cloudflared tunnel --url http://localhost:8787 # prints a https://<random>.trycloudflare.com URL
Keep this process running - the watch reaches the daemon only while it is up.
2. Point the watch app at it (and share the token)
The watch and daemon authenticate with a shared bearer token.
Generate one, put it in the daemon's env, and put the same token plus your tunnel URL in the watch config.
# daemon side (git-ignored)
printf ' CLAUDEWATCH_TOKEN=%s\nPORT=8787\n ' " $( openssl rand -hex 32 ) " > apps/daemon/.env
Then edit apps/watch/source/Config.mc (git-ignored - copy from Config.mc.example if missing):
const BASE_URL as Lang.String = "https://<your-tunnel-hostname>"; // no trailing slash
const TOKEN as Lang.String = "<the same token as apps/daemon/.env>";
const LIMIT as Lang.Number = 10;
The token in Config.mc must match CLAUDEWATCH_TOKEN in apps/daemon/.env , or every request returns 401.
# run the daemon (reads the token from apps/daemon/.env)
npm run serve
# ...and the tunnel in another terminal (step 1)
# build the signed .prg
npm run build:watch # -> apps/watch/bin/ClaudeWatch.prg
Before it builds, confirm the public URL is reachable end to end:
curl -s -o /dev/null -w ' %{http_code}\n ' -H " Authorization: Bearer <token> " " https://<your-tunnel-hostname>/sessions?limit=1 " # expect 200
4. Sideload to the watch
The .prg is just a file the watch reads from GARMIN/APPS/ .
The catch on macOS: the FR165 connects over MTP , not USB mass storage, so it does not mount as a drive in /Volumes , and you can't copy to it from the macOS terminal / Finder.
Choose the easiest working path:
Windows (easiest): plug the watch in, open File Explorer -> This PC -> Forerunner 165 -> Internal Storage -> GARMIN -> APPS , and paste ClaudeWatch.prg there.
macOS via openMTP : a free MTP app that handles Garmin well (Android File Transfer is unreliable with Garmin). Browse the watch -> GARMIN/APPS and drop the .prg in.
Put ClaudeWatch.prg directly in GARMIN/APPS/ (not a subfolder), eject cleanly, then open ClaudeWatch from the watch's app/glance list.
If it isn't listed, reboot the watch (it rescans GARMIN/APPS on boot).
Keep the daemon + tunnel running on your Mac, and keep the Mac awake and online .
Keep your phone nearby, Bluetooth on, Garmin Connect running .
Open ClaudeWatch .
It loads your session list.
From a session you can run a task; the Usage row (under Refresh) opens the dashboard.
You can use it away from home as long as your phone has internet and your Mac stays up with the daemon + tunnel alive.
Session state (optional accuracy)
Each session shows a state: working (green), waiting for you (yellow), or idle (gray).
By default this is a heuristic from the transcript's freshness and last turn, which is good but approximate (a long silent operation can read as idle).
For exact, real-time state, enable the optional Claude Code hooks described in docs/session-state.md - they report working / waiting as they happen and override the heuristic.
Selecting Create PR / Run tests / Review code spawns claude -p headlessly in that session's working directory with a fixed prompt, a scoped --allowedTools allowlist, and --permission-mode acceptEdits (so no interactive prompt can block it).
These are real actions on your machine : Create PR really commits, pushes, and opens a PR; Run tests runs your test command.
Because the endpoint is behind a public tunnel , the bearer token is genuinely load-bearing.
The per-task tool allowlists keep the blast radius tight (e.g. Create PR gets only git / gh ).
Review code is read-only and the safest to try first.
The daemon serves a small, frozen contract ( contract/contract.md ), guarded by Authorization: Bearer <token> :
Validate the canonical fixture any time with node contract/validate.mjs .
Both config files are git-ignored ; commit nothing with real secrets.
apps/daemon/.env (loaded automatically by the daemon):
apps/watch/source/Config.mc : BASE_URL , TOKEN (must match the daemon), LIMIT .
The watch build ( tools/build-watch.sh ) auto-discovers the newest installed SDK; override with CIQ_SDK_BIN , CIQ_DEV_KEY , CIQ_DEVICE if your paths differ.
Keeping it always-on (optional)
The current setup runs in a terminal and dies when the session ends or the Mac sleeps.
For a durable, hands-off setup:
A launchd LaunchAgent so the daemon auto-starts on boot and restarts on crash.
caffeinate (or pmset ) so the Mac never sleeps and drops the tunnel.
A stable URL via a named Cloudflare tunnel (needs a domain) or Tailscale Funnel (free), so you never rebuild the .prg for a URL change again.
The Mac still has to be powered and online, because that is where your sessions and repos live.
"Can't load" / "Phone not connected" on the watch.
The phone bridge isn't live.
Make sure the phone is nearby, Bluetooth is on, the Garmin Connect app is running, and the phone has internet.
Press Start on the error screen to retry.
Everything returns 401.
The token in apps/watch/source/Config.mc doesn't match CLAUDEWATCH_TOKEN in apps/daemon/.env .
Fix the token, rebuild the .prg , and re-sideload.
The watch won't mount as a drive on macOS.
The FR165 uses MTP, not mass storage - this is expected.
Use a Windows PC (native MTP in File Explorer) or openMTP on the Mac.
The Mac shows only a "BillBoard Device" and no Garmin.
The USB connection is charge-only.
Use the Garmin data cable, and if you go through a USB-A to USB-C adapter, use a data-capable one plugged directly into the Mac.
The app isn't in the watch's app list.
Reboot the watch so it rescans GARMIN/APPS , and confirm ClaudeWatch.prg is directly inside GARMIN/APPS/ .
The watch app stopped reaching the daemon after some time.
A quick tunnel URL rotates when cloudflared restarts (Mac sleep, terminal close).
Restart the tunnel, update BASE_URL , rebuild, re-sideload - or move to a named tunnel / Tailscale Funnel for a stable URL.
A session shows no model.
Fresh sessions that haven't had a real (non-synthetic) assistant turn yet have no recorded model; it fills in once the session runs.
A task fails immediately.
Confirm claude is on the daemon's PATH and authenticated, and that the task's cwd is a real repo (for Create PR you also need gh authenticated).
Path
What it is
contract/
The frozen HTTP+JSON contract and the canonical sessions.json fixture both sides build against.
apps/daemon/
Node + TypeScript daemon (the only npm workspace): reads sessions, serves the API, runs tasks.
apps/watch/
Connect IQ (Monkey C) app for the Forerunner 165.
tools/tunnel/
Cloudflare tunnel config + run scripts ( npm run serve:tunnel ).
tools/dev/run-sim.sh
The npm run

[truncated]
