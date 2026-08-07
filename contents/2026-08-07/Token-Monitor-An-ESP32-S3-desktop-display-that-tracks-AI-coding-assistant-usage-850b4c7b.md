---
source: "https://www.cnx-software.com/2026/08/07/token-monitor-an-esp32-s3-desktop-display-that-tracks-ai-coding-assistant-usage/"
hn_url: "https://news.ycombinator.com/item?id=49213890"
title: "Token Monitor, An ESP32-S3 desktop display that tracks AI coding assistant usage"
article_title: "Token Monitor - An ESP32-S3 desktop display that tracks AI coding assistant usage (Crowdfunding) - CNX Software"
author: "jandeboevrie"
captured_at: "2026-08-07T18:41:07Z"
capture_tool: "hn-digest"
hn_id: 49213890
score: 1
comments: 0
posted_at: "2026-08-07T17:48:08Z"
tags:
  - hacker-news
  - translated
---

# Token Monitor, An ESP32-S3 desktop display that tracks AI coding assistant usage

- HN: [49213890](https://news.ycombinator.com/item?id=49213890)
- Source: [www.cnx-software.com](https://www.cnx-software.com/2026/08/07/token-monitor-an-esp32-s3-desktop-display-that-tracks-ai-coding-assistant-usage/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T17:48:08Z

## Translation

タイトル: Token Monitor、AI コーディング アシスタントの使用状況を追跡する ESP32-S3 デスクトップ ディスプレイ
記事タイトル: Token Monitor - AI コーディング アシスタントの使用状況を追跡する ESP32-S3 デスクトップ ディスプレイ (クラウドファンディング) - CNX ソフトウェア
説明: Fractal Manifold は、リアルタイム AI コーディング アシスタントの使用状況を監視するための 4 インチ ESP32-S3 ベースのデスクトップ ディスプレイである Token Monitor を導入しました。デバイス

記事本文:
Token Monitor - AI コーディング アシスタントの使用状況を追跡する ESP32-S3 デスクトップ ディスプレイ (クラウドファンディング) - CNX ソフトウェア
コンテンツにスキップ
CNX ソフトウェア – 組み込みシステム ニュース
組み込みシステム、IoT、オープンソース ハードウェア、SBC、マイクロコントローラー、プロセッサーなどに関するレビュー、チュートリアル、最新ニュース
広告およびコンサルティングサービス
Token Monitor – AI コーディング アシスタントの使用状況を追跡する ESP32-S3 デスクトップ ディスプレイ (クラウドファンディング)
Fractal Manifold は、リアルタイム AI コーディング アシスタントの使用状況を監視するための 4 インチ ESP32-S3 ベースのデスクトップ ディスプレイである Token Monitor を導入しました。
このデバイスは、5 点静電容量式タッチ機能を備えた 4 インチ 480 × 480 IPS タッチスクリーン、Wi-Fi 4 および Bluetooth LE 5 ワイヤレス接続、USB-C 電源、AXP2101 PMIC によるオプションの 1 セル リチウムイオン バッテリのサポート、電源および起動機能用の 2 つのサイド ボタンを備えています。クォータ消費量、セッション制限、トークン数、リセット タイマー、推定トークン コストなどの使用状況情報が表示されます。ローカル ブローカー サービスは使用状況データを取得し、安全な接続を介してデバイスに転送し、資格情報をユーザーのマシンに保持します。
SoC – Espressif システム ESP32-S3
CPU – AI/ML ワークロード向けのベクトル拡張機能を備えた最大 240 MHz のデュアルコア Tensilica LX7
ワイヤレス - WiFi 4 および Bluetooth LE 5
ディスプレイ – 4 インチ 480×480 静電容量式タッチスクリーン、ST7701 ドライバーおよび GT911 タッチ コントローラー付き
USB – 電源用の USB Type-C ポート
電源
USB Type-Cポート経由で5V
内蔵 1S リチウムイオンバッテリー (5 時間以上のセッション時間をカバー)
寸法 – 設置面積 86 x 86 mm
エンクロージャ – 3D プリントされたマット ABS シェル。専用のデスクトップベースが含まれています
デバイスは、Mac または Linux PC 上で実行されている Apache-2.0 ライセンスを取得したローカル オープンソース ブローカー (tokenmonitor-mcp) に接続し、CLI ログを解析して使用量クォータを取得します。

プロバイダー API から直接。ブローカーには、Node.js (20 以上)、Python 3.11 以降 (uv、pip、または venv 経由)、または Go (1.25 以上) を備えたローカル ランタイム環境も必要です。 HMAC 署名付きのリプレイ保護されたリクエストを使用して、ポート 8765 で LAN ベースのポーリングをリッスンするため、テレメトリや生のコードは外部に送信されません。このプラットフォームは、Claude Code、Codex CLI、および Antigravity CLI をすぐにサポートします。
もう 1 つのソフトウェア機能は、Model Context Protocol (MCP) との統合です。ローカル ブローカーは MCP サーバーとして機能し、17 のハードウェア管理ツールを Claude Code や Codex などの AI アシスタントに公開します。これにより、mDNS または USB-CDC シリアル経由で ESP32-S3 をプロビジョニングし、ESP-IDF 診断ログを表示し、Wi-Fi 認証情報を更新し、Ed25519 署名付き OTA ファームウェア アップデートを AI CLI から直接インストールすることができます。ブローカーの詳細については、プロジェクトの GitHub リポジトリで入手できます。
トークン モニターには、クォータ消費やリセット タイマー (明るいテーマと暗いテーマ) を含むクロード コードの使用状況がリアルタイムで表示されます。
ファームウェアはLVGLに基づいています。また、Open-Meteo を使用して、地元の日の出と日の入りに同期した明るいテーマと暗いテーマもサポートしています。また、ローカル ブローカーからのユーザー定義の JSON データを折れ線グラフ、円グラフ、表、テキストとして表示して、AI の使用以外の追加のメトリクスを監視できるカスタマイズ可能なパネルもあります。
Token Monitor は、Claude Desktop Buddy 、Clawdmeter 、ディスプレイ内蔵の CtrlVibe AI コンソール キーパッド、ZecTrix Note 4 E-Ink ディスプレイ (vibecoding-voice/vibecoding-plus ファームウェアと併用した場合) など、AI コーディング アシスタントの監視と制御のための ESP32 ベースのスマート ディスプレイの長いリストに追加されます。
3D プリントされた筐体、通気スロット、デスクトップ スタンドを示す Token Monitor の背面図。
ESP32-S3 ベースのデスクトップ Token Monitor が利用可能になりました

Kickstarter で 25,000 ユーロの資金目標を掲げています。 Super Early Bird 特典の価格は、白い Token Monitor デバイスで 99 ユーロ (約 115 ドル) ですが、標準の Early Bird レベルの価格は 120 ユーロです。この記事の執筆時点では、現在配送可能なのは米国、カナダ、EU、英国のみで、最初のユニットは 2026 年 11 月に出荷される予定です。
Debashis Das は、業界で 5 年以上の経験を持つテクニカル コンテンツ ライター兼組み込みエンジニアです。組み込み C、PCB 設計、SEO 最適化の専門知識を活かし、難しい技術的なトピックと明確なコミュニケーションを効果的に融合させています。
これを共有してください:
CNX ソフトウェアをサポート!暗号通貨経由で寄付したり、Patreon でパトロンになったり、Amazon や Aliexpress で商品を購入したりできます。また、記事内でアフィリエイト リンクを使用し、リンクをクリックして商品を購入するとコミッションが発生します。
SenseCAP Watcher は、LLM ベースのスペース監視用の音声制御の物理 AI エージェントです (クラウドファンディング)
vPlayer – カスタム拡張オプションを備えた 1.69 インチ ESP32-S3 タッチスクリーン ビデオ ディスプレイ
RIMU モジュラー IoT コントローラーは 4.3 インチのタッチスクリーン ディスプレイを備え、拡張のための積み重ね可能な設計を備えています (クラウドファンディング)
Espressif の EchoEar ESP32-S3 音声制御 AI チャットボットは esp-brookesia ファームウェアを実行します
EnviroGo ESP32-S3 ウェアラブル環境モニターは 7 つのセンサーを備えています (クラウドファンディング)
ラベル
{}
[+]
名前*
電子メール*
ウェブサイト
プライバシーポリシーに同意します
コメント フォームでは、Web サイトに投稿されたコメントを追跡できるように、あなたの名前、電子メール、内容が収集されます。コメントを投稿するには、当社の Web サイト利用規約とプライバシー ポリシーを読んで同意してください。
接続先:
ラベル
{}
[+]
名前*
電子メール*
ウェブサイト
プライバシーポリシーに同意します
コメント フォームでは、Web サイトに投稿されたコメントを追跡できるように、あなたの名前、電子メール、内容が収集されます。

コメントを投稿するには、当社の Web サイト利用規約とプライバシー ポリシーを読んで同意してください。
0コメント
最古の
最新の
さらにコメントを読み込む
ポストナビゲーション
Twitter LinkedIn Telegram YouTube RSS
注目の投稿 - 過去 7 日間
ZecTrix Note 4 – 音声タスク用の 4.2 インチ ESP32-S3 ワイヤレス電子ペーパー ディスプレイコメント0件
NanoPi R28S – USB-C コンソール ポートを備えた小型デュアル GbE Rockchip RK3528A ワイヤレス SBC およびルーターコメント1件
2,890 万パラメータの LLM は、ESP32-S3 上で 9 トークン/秒でローカルに実行されます。コメント5件
Espressif ESP32-C61-MINI-1/1U Wi-Fi 6 および BLE IoT モジュールが約 2 ドルで発売コメント5件
ChronoWatch X2040 – 既製の Raspberry Pi RP2040 ラウンド ディスプレイ、3D プリント ケース、カスタム ファームウェアを備えた DIY ウォッチ8件のコメント
匿名 on GEEKOM IT13 Max ミニ PC レビュー – パート 1: 仕様、開梱、分解、および最初の起動
Espressif Systems の Sarah J Cunningham が、オープンソースのデバイスからクラウド、電話までの IoT プラットフォーム ESP RainMaker Neo をリリース
maurer による GEEKOM IT13 Max ミニ PC レビュー – パート 1: 仕様、開梱、分解、および最初の起動
Amlogic A311Y3 オクタコア エッジ AI SoC 上の Stane1983 は、Cortex-A78/A55 コア、8 TOPS NPU、LPDDR5 サポートを備えています
Amlogic A311Y3 オクタコア エッジ AI SoC 上の Stane1983 は、Cortex-A78/A55 コア、8 TOPS NPU、LPDDR5 サポートを備えています
GEEKOM IT13 Max ミニ PC レビュー – パート 1: 仕様、開梱、分解、最初の起動 Intel Core Ultra 9 185H CPU を搭載した GEEKOM IT13 Max ミニ PC のサンプルを受け取りました […]
Creality Falcon T1 レビュー – 20W ファイバー レーザー モジュールと 40W ダイオード レーザー モジュールでテストされた 5-in-1 モジュラー レーザー彫刻機 Creality は、レビューのために Falcon T1 5-in-1 モジュラー レーザー彫刻機のサンプルを送ってくれました。同社はそれを説明しています[…]
Creality Pika レビュー – 手頃な価格のポータブル 3D sca

赤外線および青色レーザー モードを備えたナー Creality から、レビュー用に Pika ポータブル 3D スキャナーのサンプルが送られてきました。ある意味、それはいくつかの点を共有しています […]

## Original Extract

Fractal Manifold has introduced the Token Monitor, a 4-inch ESP32-S3-based desktop display for monitoring real-time AI coding assistant usage. The device

Token Monitor - An ESP32-S3 desktop display that tracks AI coding assistant usage (Crowdfunding) - CNX Software
Skip to content
CNX Software – Embedded Systems News
Reviews, tutorials and the latest news about embedded systems, IoT, open-source hardware, SBC's, microcontrollers, processors, and more
Advertisement & Consulting Services
Token Monitor – An ESP32-S3 desktop display that tracks AI coding assistant usage (Crowdfunding)
Fractal Manifold has introduced the Token Monitor, a 4-inch ESP32-S3-based desktop display for monitoring real-time AI coding assistant usage.
The device features a 4-inch 480×480 IPS touchscreen with 5-point capacitive touch, Wi-Fi 4 and Bluetooth LE 5 wireless connectivity, USB-C power, optional 1-cell Li-ion battery support via an AXP2101 PMIC, and two side buttons for power and boot functions. It shows usage information including quota consumption, session limits, token counts, reset timers, and estimated token costs. A local broker service fetches the usage data and transfers it to the device over a secure connection, keeping credentials on the user’s machine.
SoC – Espressif Systems ESP32-S3
CPU – Dual-core Tensilica LX7 up to 240 MHz with vector extension for AI/ML workloads
Wireless – WiFi 4 and Bluetooth LE 5
Display – 4-inch 480×480 capacitive touchscreen with ST7701 driver and GT911 touch controller
USB – USB Type-C port for power
Power supply
5V via USB Type-C port
Internal 1S Li-ion battery (covers >5 hours of session time)
Dimensions – 86 x 86 mm footprint
Enclosure – 3D-printed matte ABS shell; includes a dedicated desktop base
The device connects to an Apache-2.0 licensed, local open-source broker (tokenmonitor-mcp) running on a Mac or Linux PC, which parses CLI logs and fetches usage quotas directly from provider APIs. The broker also requires a local runtime environment with Node.js (≥20), Python 3.11+ (via uv, pip, or venv), or Go (≥1.25). It listens for LAN-based polling on port 8765 using HMAC-signed, replay-protected requests, so no telemetry or raw code is transmitted externally. The platform supports Claude Code, Codex CLI, and Antigravity CLI out of the box.
Another software feature is integration with the Model Context Protocol (MCP). The local broker acts as an MCP server and exposes 17 hardware management tools to AI assistants like Claude Code and Codex. This allows you to provision the ESP32-S3 over mDNS or USB-CDC serial, show ESP-IDF diagnostic logs, update Wi-Fi credentials, and install Ed25519-signed OTA firmware updates directly from the AI CLI. More information about the broker is available on the project’s GitHub repository .
The Token Monitor displays real-time Claude Code usage, including quota consumption and reset timers (with light and dark themes)
The firmware is based on LVGL. It also supports light and dark themes synchronized with local sunrise and sunset using Open-Meteo. There is also a customizable panel that can display user-defined JSON data from the local broker as line graphs, pie charts, tables, and text to monitor additional metrics beyond AI usage.
The Token Monitor adds to a long list of ESP32-based smart displays for AI coding assistants monitoring and control, such as the Claude Desktop Buddy , Clawdmeter , CtrlVibe AI Console keypad with built-in display, ZecTrix Note 4 E-Ink display (when used with vibecoding-voice/vibecoding-plus firmware), among others.
Token Monitor’s rear view showing the 3D-printed enclosure, ventilation slots, and desktop stand.
The ESP32-S3-based desktop Token Monitor is now available on Kickstarter with a funding goal of €25,000. The Super Early Bird reward is priced at €99 (about $115) for a white Token Monitor device, while the standard Early Bird tier costs €120. At the time of writing, I can only see delivery currently available to the US, Canada, the EU, and the UK, and the first units are expected to ship in November 2026.
Debashis Das is a technical content writer and embedded engineer with over five years of experience in the industry. With expertise in Embedded C, PCB Design, and SEO optimization, he effectively blends difficult technical topics with clear communication
Share this:
Support CNX Software! Donate via cryptocurrencies , become a Patron on Patreon, or purchase goods on Amazon or Aliexpress . We also use affiliate links in articles to earn commissions if you make a purchase after clicking on those links.
The SenseCAP Watcher is a voice-controlled, physical AI agent for LLM-based space monitoring (Crowdfunding)
vPlayer – A 1.69-inch ESP32-S3 touchscreen video display with custom expansion options
RIMU modular IoT controller features 4.3-inch touchscreen display, stackable design for expansion (Crowdfunding)
Espressif’s EchoEar ESP32-S3 voice-controlled AI chatbot runs esp-brookesia firmware
EnviroGo ESP32-S3 wearable environmental monitor features 7 sensors (Crowdfunding)
Label
{}
[+]
Name*
Email*
Website
I agree to the Privacy Policy
The comment form collects your name, email and content to allow us keep track of the comments placed on the website. Please read and accept our website Terms and Privacy Policy to post a comment.
Connect with:
Label
{}
[+]
Name*
Email*
Website
I agree to the Privacy Policy
The comment form collects your name, email and content to allow us keep track of the comments placed on the website. Please read and accept our website Terms and Privacy Policy to post a comment.
0 Comments
oldest
newest
Load More Comments
Post navigation
Twitter LinkedIn Telegram YouTube RSS
Trending Posts - Last 7 Days
ZecTrix Note 4 – A 4.2-inch ESP32-S3 wireless e-paper display for voice tasks 13.7k views | 0 comments
NanoPi R28S – A tiny dual GbE Rockchip RK3528A wireless SBC and router with a USB-C console port 6.7k views | 1 comment
28.9M-parameter LLM runs locally on ESP32-S3 at 9 tokens/s 6.1k views | 5 comments
Espressif ESP32-C61-MINI-1/1U Wi-Fi 6 and BLE IoT module launched for about $2 5.6k views | 5 comments
ChronoWatch X2040 – A DIY watch with off-the-shelf Raspberry Pi RP2040 round display, 3D printed case, and custom firmware 4.2k views | 8 comments
Anonymous on GEEKOM IT13 Max mini PC review – Part 1: specifications, unboxing, teardown, and first boot
Sarah J Cunningham on Espressif Systems releases ESP RainMaker Neo open-source device-to-cloud-to-phone IoT platform
maurer on GEEKOM IT13 Max mini PC review – Part 1: specifications, unboxing, teardown, and first boot
Stane1983 on Amlogic A311Y3 octa-core Edge AI SoC features Cortex-A78/A55 cores, 8 TOPS NPU, LPDDR5 support
Stane1983 on Amlogic A311Y3 octa-core Edge AI SoC features Cortex-A78/A55 cores, 8 TOPS NPU, LPDDR5 support
GEEKOM IT13 Max mini PC review – Part 1: specifications, unboxing, teardown, and first boot We’ve just received a sample of the GEEKOM IT13 Max mini PC powered by an Intel Core Ultra 9 185H CPU […]
Creality Falcon T1 review – A 5-in-1 modular laser engraver tested with 20W fiber and 40W diode laser modules Creality sent us a sample of the Falcon T1 5-in-1 modular laser engraver for review. The company describes it […]
Creality Pika review – An affordable, portable 3D scanner with infrared and blue laser modes Creality has sent us a sample of the Pika portable 3D scanner for review. In some ways, it shares some of the […]
