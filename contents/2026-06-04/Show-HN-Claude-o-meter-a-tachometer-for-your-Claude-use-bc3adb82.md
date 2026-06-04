---
source: "https://github.com/joshcarter/claude-o-meter"
hn_url: "https://news.ycombinator.com/item?id=48401978"
title: "Show HN: Claude-o-meter, a tachometer for your Claude use"
article_title: "GitHub - joshcarter/claude-o-meter: Keep your Claude use at the redline or you're just giving Anthropic free money · GitHub"
author: "jdcarter"
captured_at: "2026-06-04T18:53:25Z"
capture_tool: "hn-digest"
hn_id: 48401978
score: 1
comments: 0
posted_at: "2026-06-04T17:40:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude-o-meter, a tachometer for your Claude use

- HN: [48401978](https://news.ycombinator.com/item?id=48401978)
- Source: [github.com](https://github.com/joshcarter/claude-o-meter)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T17:40:27Z

## Translation

タイトル: HN を表示: Claude-o-meter、クロードが使用するタコメーター
記事のタイトル: GitHub - joshcarter/claude-o-meter: クロードの使用をレッドラインに保つか、Anthropic に無料のお金を与えているだけです · GitHub
説明: クロードの使用をレッドラインに維持しなければ、Anthropic に無料のお金を提供しているだけです - joshcarter/claude-o-meter
HN テキスト: ただの愚かなプロジェクト。そのレッドラインは、5 時間枠の終わりにトークンを使い果たすトークンの燃焼率です。明らかに 1980 年代のコルベットのダッシュボードからインスピレーションを得ています。

記事本文:
GitHub - joshcarter/claude-o-meter: クロードの使用をレッドラインに維持しなければ、Anthropic に無料のお金を提供しているだけです · GitHub
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
ジョシュカーター
/
クロード・オ・メーター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のna

違反オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
33 コミット 33 コミット claude_o_meter claude_o_meter デプロイ デプロイ グラフィックス printd_parts print_parts .gitignore .gitignore DONE.md DONE.md LICENSE LICENSE README.md README.md TODO.md TODO.md TODO_BACKLOG.md TODO_BACKLOG.md TODO_CRITICAL.md TODO_CRITICAL.md TODO_DEFERRED.md TODO_DEFERRED.md pyrightconfig.json pyrightconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
セッション内でのトークンの使用状況を示すクロード タコメーター
窓。 5 時間枠がメインのゲージであり、その赤線は次のとおりです。
トークンはウィンドウの終わりにぴったり使い果たされます。の
残量計は 7 日間の目安です。
これは、Raspberry Pi 上で実行するように設計されたシンプルな Python アプリですが、
デスクトップ上のウィンドウに表示することもできます。設定する必要があります
Anthropic セッション ID (以下を参照) を入力し、10 分以上ポーリングします。
正確な燃焼率の数値を得るために。
にフィットする独自のデザインの 3D 印刷可能なフレームが含まれています。
Adafruit PiTFT 3.5 プラス。これは現在のRaspberry Piでは問題なく動作します
2026 年時点の OS (trixie ベース)。
python3 -m venv .venv
.venv/bin/pip install -r claude_o_meter/requirements.txt
.venv/bin/python -m claude_o_meter # DATA_SOURCE のデフォルトは「fake」
480×320のウィンドウが開きます。デフォルトの DATA_SOURCE = "fake" では、ゲージは次のように実行されます。
デモ サイクル - ネットワーク キーやセッション キーは必要ありません - そのため、すべての状態を監視できます。
タコメーターと燃料計がその範囲をスイープし、燃料低下とエンジンチェックライトが点灯します
各頂点の近くで点滅し、障害メッセージがそのバリエーションを循環します。閉じる
ウィンドウを閉じるか、Ctrl-C を押して終了します。
サイクルが終了するのを待つのではなく、すべての状態をすばやく実行するには、
デモを明らかにします — タコメーター、燃料、両方のライト、そして障害メッセージ:
.venv/bin/pytho

n -m claude_o_meter.demo_reveal # オプションのミリ秒/ステップ、例: 300
ライブモード (実際の使用状況)
ライブ モードではアカウントをポーリングするため、claude.ai セッション Cookie が必要です。
サインインして https://claude.ai を開き、DevTools → Application → Cookie →
https://claude.ai に移動し、 sessionKey 値をコピーします ( sk-ant-sid01-... を開始します)。
claude_o_meter/config.toml で DATA_SOURCE = "live" を設定します。
環境内で Cookie を使用して実行します。
エクスポート CLAUDE_SESSION_KEY=sk-ant-sid01-...
.venv/bin/python -m claude_o_meter
燃焼率の読み取り値を信頼する前に、10 分以上ポーリングしてください。クッキー
数週間から数か月続きます。有効期限が切れると、エンジンチェックライトが「必要」を示します。
auth」メッセージ - 新しいセッションキーを取得して再起動します。
(curl_cffi はライブ モードでのみ使用され、Chrome の TLS フィンガープリントを模倣して通過します
claude.ai の Cloudflare 。)
設定 (claude_o_meter/config.toml)
キー
デフォルト
注意事項
データソース
「偽物」
「フェイク」（オフラインの変動値）または「ライブ」（実際のポーリング）
POLL_SECONDS
60
世論調査の頻度。ライブモードでは、POLL_INTERVAL_SECONDS を設定します。
UTC_OFFSET_HOURS
0
表示される時計/日付の整数オフセット (例: PDT の場合は -7)
ディスプレイモード
「窓」
「ウィンドウ」（デスクトップ上の SDL ウィンドウ）または「フレームバッファ」（Pi TFT）
FB_DEVICE
「/dev/fb1」
DISPLAY_MODE="framebuffer" の場合のフレームバッファデバイス
DIM_OPACITY
212
調光長方形の不透明度 0 ～ 255 (212 ≈ 83% ゴースト)
上記のキーはすべて、同じ名前の環境変数として設定することもできます。
TOML をオーバーライドします (Pi サービスはこれを使用して DISPLAY_MODE=framebuffer を設定します)
および DATA_SOURCE=live )。
ライブモード環境変数
セッション キーといくつかのオーバーライドは、TOML ファイルではなく環境から取得されます。
Raspberry Pi 3 + PiTFT 搭載
同じプログラムが、3.5 インチ PiTFT (480×320) パネルを駆動する Raspberry Pi 3 上で実行されます。
単一の systemd サービスとして。デスクトップもXサーブもない

r — レンダラー
パネルのフレームバッファ (SDL ダミービデオドライバー) にフレームを直接書き込みます。
64 ビット Raspberry Pi OS Lite (Trixie) で検証済み。
1. PiTFT パネル (レガシー フレームバッファ ドライバー) を有効にします。
Adafruit の PiTFT サポートをインストールし、パネルがレガシーを使用していることを確認します。
fb_hx8357d フレームバッファ ドライバ — DRM バリアントではありません。で
/boot/firmware/config.txt オーバーレイ行には ,drm を含めることはできません。
dtoverlay=pitft35-resistive、回転=90、速度=20000000、fps=20
,drm フラグ (Adafruit インストーラーが追加する場合があります) は、DRM/KMS を強制します。
ドライバーは、モードセットになるまで SPI パネルを空白のままにし、次のドライバーと競合します。
vc4-kms-v3d — 症状: 画面が白く空白で、コンソールが表示されません。それを取り外して、
再起動:
sudo sed -i ' /pitft35-resistive/ s/,drm,/,/ ' /boot/firmware/config.txt
sudo再起動
再起動後、フレームバッファを確認します。
cat /sys/class/graphics/fb1/name # -> fb_hx8357d
cat /sys/class/graphics/fb1/bits_per_pixel # -> 16 (RGB565)
cat /sys/class/graphics/fb1/virtual_size # -> 480,320
レンダラーは深度を自動検出します (ここでは 16bpp RGB565、フォールバックとして 32bpp)。
numpy は各 RGB565 フレームをパックします。
git clone https://github.com/joshcarter/claude-o-meter.git ~ /claude-o-meter
cd ~ /クロード・オ・メーター
python3 -m venv .venv
.venv/bin/pip install -r claude_o_meter/requirements.txt
pygame 、curl_cffi 、および numpy はすべて、事前に構築された aarch64 ホイールからインストールされます
64 ビット Raspberry Pi OS — コンパイラは必要ありません。
printf ' CLAUDE_SESSION_KEY=sk-ant-sid01-...\n ' > ~ /claude-o-meter/.env
chmod 600 ~ /claude-o-meter/.env
4. サービスをインストールする
deploy/claude-o-meter.service は /home/pi/claude-o-meter にリポジトリを想定します。
ユーザー pi として実行し、DISPLAY_MODE=framebuffer / DATA_SOURCE=live を設定します。
環境 (そのため、config.toml の編集は必要ありません)。 pi ユーザーは次の場所にいる必要があります
/dev/fb1 (Raspbe のデフォルト) を書き込むビデオ グループ

Pi OS)。を編集します
パスまたはユーザーが異なる場合は単位を変更します。
sudo cp ~ /claude-o-meter/deploy/claude-o-meter.service /etc/systemd/system/
sudo systemctl デーモン-リロード
sudo systemctl Enable --now claude-o-meter
systemctl status claude-o-meter --no-pager
サンプル.db は /var/lib/claude-o-meter/ の下に作成されます (systemd
状態ディレクトリ)。一度再起動して、ダッシュボードが無人で起動することを確認します。
sudo再起動
ログ:journalctl -u claude-o-meter -f 。
ディスプレイの各部分に表示される内容は次のとおりです。
タコメータースケール（「レッドライン」）
ゲージと 0 ～ 99 の表示値は両方とも 5 時間の redline_ratio から取得されます。
ポーラーによって報告されます ( burn_rate / Sustainable_rate 、1.0 はあなたを意味します)
リセット時にウィンドウの予算全体を正確に消費する予定です)。の
マッピングは claude_o_meter/gauges.py にあります。
tach_position() は、比率を連続的なゲージ位置にマップします。
0.0 .. TACH_FRAMES-1 (20 セグメント) を 2 つに分割:
レッドラインより下 (比率 ≤ 1.0): 凹型の曲線、
位置 = REDLINE_FRAME * 比率 ** BLUE_EXPONENT 。と
BLUE_EXPONENT = 0.5 低い比率では急峻で、比率に向かって平坦になります。
レッドラインなので、控えめなユーザーでも一日中針が動いているのを見ることができます。
レッドラインより上 (比率 > 1.0): REDLINE_FRAME から上部までの直線
セグメント、バーンレートが RED_FULL_RATIO (デフォルトは 2x) に達するとペギング
持続可能です）。
redline_ratio == 1.0 は、REDLINE_FRAME (17) の最上部に正確に到達します。
ゲージの黄色の帯 — それを超えるものは赤になります。読み出しは同じです
位置はセグメントに量子化されるのではなく、0 ～ 99 ( tach_number() ) にスケーリングされるため、
セグメントの変更の間でも移動します。
claude_o_meter/gauges.py のチューニングノブ: REDLINE_FRAME 、BLUE_EXPONENT
(低い = 使用量が少ない場合に感度が高くなります)、および RED_FULL_RATIO 。
.venv/bin/pip pytest をインストールする
.venv/bin/python -m pytest claude_o_meter/tests/
ディスプレイテスト実行ヘッド

SDL_VIDEODRIVER=dummy (テストで設定) 経由で少なく、
したがって、ウィンドウやライブ Cookie は必要ありません。
3D プリントされたフレーム デザイン (printed_parts/) — CC BY-SA 4.0、
「printed_parts/LICENSE」を参照してください。 Printablesにも掲載されています。
クロードの使用をレッドラインに維持しなければ、Anthropic に無料のお金を提供しているだけです
読み込み中にエラーが発生しました。このページをリロードしてください。
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

Keep your Claude use at the redline or you're just giving Anthropic free money - joshcarter/claude-o-meter

Just a silly project. Its redline is the token burn rate where you'll exhaust your tokens right at the end of your 5-hour window. Clearly inspired by 1980's Corvette dashboards.

GitHub - joshcarter/claude-o-meter: Keep your Claude use at the redline or you're just giving Anthropic free money · GitHub
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
joshcarter
/
claude-o-meter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
33 Commits 33 Commits claude_o_meter claude_o_meter deploy deploy graphics graphics printed_parts printed_parts .gitignore .gitignore DONE.md DONE.md LICENSE LICENSE README.md README.md TODO.md TODO.md TODO_BACKLOG.md TODO_BACKLOG.md TODO_CRITICAL.md TODO_CRITICAL.md TODO_DEFERRED.md TODO_DEFERRED.md pyrightconfig.json pyrightconfig.json View all files Repository files navigation
A Claude tachometer which shows your token use within your session
window. The 5-hour window is the main gauge, with its redline where
your tokens will be exhaused exactly at the end of the window. The
fuel gauge is your 7-day window.
This is a simple Python app designed to run on a Raspberry Pi, but it
can also display in a window on your desktop. You'll need to configure
your Anthropic session ID (see below) and let it poll for 10+ minutes
to give you accurate burn rate numbers.
I've included a 3D printable frame of my own design which fits the
Adafruit PiTFT 3.5 Plus. This works fine on the current Raspberry Pi
OS (trixie-based) as of 2026.
python3 -m venv .venv
.venv/bin/pip install -r claude_o_meter/requirements.txt
.venv/bin/python -m claude_o_meter # DATA_SOURCE defaults to "fake"
A 480×320 window opens. With the default DATA_SOURCE = "fake" the gauges run on
a demo cycle — no network or session key needed — so you can watch every state:
the tach and fuel gauge sweep their range, the low-fuel and check-engine lights
blink near each crest, and the fault message cycles through its variants. Close
the window or press Ctrl-C to quit.
To step through every state quickly instead of waiting out the cycle, run the
reveal demo — tach, fuel, both lights, then the fault messages:
.venv/bin/python -m claude_o_meter.demo_reveal # optional ms/step, e.g. 300
Live mode (your real usage)
Live mode polls your account, so it needs your claude.ai session cookie:
Open https://claude.ai signed in, then DevTools → Application → Cookies →
https://claude.ai , and copy the sessionKey value (starts sk-ant-sid01-... ).
Set DATA_SOURCE = "live" in claude_o_meter/config.toml .
Run with the cookie in the environment:
export CLAUDE_SESSION_KEY=sk-ant-sid01-...
.venv/bin/python -m claude_o_meter
Let it poll for 10+ minutes before trusting the burn-rate reading. The cookie
lasts weeks to months; when it expires the check-engine light shows a "needs
auth" message — get a fresh sessionKey and restart.
( curl_cffi , used only in live mode, mimics Chrome's TLS fingerprint to get past
Cloudflare on claude.ai .)
Configuration ( claude_o_meter/config.toml )
Key
Default
Notes
DATA_SOURCE
"fake"
"fake" (offline oscillating values) or "live" (real polling)
POLL_SECONDS
60
Poll cadence; in live mode it sets POLL_INTERVAL_SECONDS
UTC_OFFSET_HOURS
0
Integer offset for displayed clock/date (e.g. -7 for PDT)
DISPLAY_MODE
"window"
"window" (SDL window on a desktop) or "framebuffer" (Pi TFT)
FB_DEVICE
"/dev/fb1"
Framebuffer device when DISPLAY_MODE="framebuffer"
DIM_OPACITY
212
Dimming-rectangle opacity 0–255 (212 ≈ 83% ghost)
Any key above can also be set as an environment variable of the same name, which
overrides the TOML (the Pi service uses this to set DISPLAY_MODE=framebuffer
and DATA_SOURCE=live ).
Live-mode environment variables
The session key and a few overrides come from the environment, not the TOML file:
Raspberry Pi 3 + PiTFT deployment
The same program runs on a Raspberry Pi 3 driving a 3.5" PiTFT (480×320) panel
as a single systemd service. There is no desktop or X server — the renderer
writes frames straight to the panel's framebuffer (SDL dummy video driver).
Verified on 64-bit Raspberry Pi OS Lite (Trixie).
1. Enable the PiTFT panel (legacy framebuffer driver)
Install Adafruit's PiTFT support, then make sure the panel uses the legacy
fb_hx8357d framebuffer driver — not the DRM variant. In
/boot/firmware/config.txt the overlay line must not contain ,drm :
dtoverlay=pitft35-resistive,rotate=90,speed=20000000,fps=20
The ,drm flag (which the Adafruit installer may add) forces the DRM/KMS
driver, which leaves the SPI panel blank until a modeset and conflicts with
vc4-kms-v3d — symptom: a white/blank screen and no console. Remove it and
reboot:
sudo sed -i ' /pitft35-resistive/ s/,drm,/,/ ' /boot/firmware/config.txt
sudo reboot
After reboot, confirm the framebuffer:
cat /sys/class/graphics/fb1/name # -> fb_hx8357d
cat /sys/class/graphics/fb1/bits_per_pixel # -> 16 (RGB565)
cat /sys/class/graphics/fb1/virtual_size # -> 480,320
The renderer auto-detects the depth (16bpp RGB565 here, 32bpp as a fallback);
numpy packs each RGB565 frame.
git clone https://github.com/joshcarter/claude-o-meter.git ~ /claude-o-meter
cd ~ /claude-o-meter
python3 -m venv .venv
.venv/bin/pip install -r claude_o_meter/requirements.txt
pygame , curl_cffi , and numpy all install from prebuilt aarch64 wheels on
64-bit Raspberry Pi OS — no compiler required.
printf ' CLAUDE_SESSION_KEY=sk-ant-sid01-...\n ' > ~ /claude-o-meter/.env
chmod 600 ~ /claude-o-meter/.env
4. Install the service
deploy/claude-o-meter.service assumes the repo at /home/pi/claude-o-meter ,
runs as user pi , and sets DISPLAY_MODE=framebuffer / DATA_SOURCE=live via
the environment (so no config.toml edit is needed). The pi user must be in
the video group to write /dev/fb1 (the default on Raspberry Pi OS). Edit the
unit if your path or user differs.
sudo cp ~ /claude-o-meter/deploy/claude-o-meter.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now claude-o-meter
systemctl status claude-o-meter --no-pager
samples.db is created under /var/lib/claude-o-meter/ (systemd
StateDirectory ). Reboot once to confirm the dashboard comes up unattended:
sudo reboot
Logs: journalctl -u claude-o-meter -f .
What each part of the display shows:
Tachometer scale (the "redline")
The gauge and the 0–99 readout both come from the 5-hour redline_ratio
reported by the poller ( burn_rate / sustainable_rate , where 1.0 means you
are on track to spend the window's whole budget exactly when it resets). The
mapping lives in claude_o_meter/gauges.py .
tach_position() maps the ratio to a continuous gauge position
0.0 .. TACH_FRAMES-1 (20 segments) in two pieces:
Below the redline ( ratio ≤ 1.0): a concave curve,
position = REDLINE_FRAME * ratio ** BLUE_EXPONENT . With
BLUE_EXPONENT = 0.5 it is steep at low ratios and flattens toward the
redline, so a modest user still sees the needle move through the day.
Above the redline ( ratio > 1.0): linear from REDLINE_FRAME to the top
segment, pegging once the burn rate reaches RED_FULL_RATIO (default 2×
sustainable).
redline_ratio == 1.0 lands exactly on REDLINE_FRAME (17), the top of the
gauge's yellow band — anything past that is red. The readout is the same
position scaled to 0–99 ( tach_number() ) instead of quantised to segments, so
it moves even between segment changes.
Tuning knobs in claude_o_meter/gauges.py : REDLINE_FRAME , BLUE_EXPONENT
(lower = more sensitive at low use), and RED_FULL_RATIO .
.venv/bin/pip install pytest
.venv/bin/python -m pytest claude_o_meter/tests/
The display tests run headless via SDL_VIDEODRIVER=dummy (set in the tests),
so no window or live cookie is needed.
3D-printed frame design ( printed_parts/ ) — CC BY-SA 4.0,
see printed_parts/LICENSE . Also published on Printables.
Keep your Claude use at the redline or you're just giving Anthropic free money
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
