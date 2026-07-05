---
source: "https://github.com/Dominic-Muscatella/weather-station-alpha/tree/master"
hn_url: "https://news.ycombinator.com/item?id=48796441"
title: "I built a weather station that runs machine learning to forecast weather"
article_title: "GitHub - Dominic-Muscatella/weather-station-alpha: A raspberry pi based system for severe weather warnings using local data only. · GitHub"
author: "makegreatthings"
captured_at: "2026-07-05T19:07:29Z"
capture_tool: "hn-digest"
hn_id: 48796441
score: 2
comments: 1
posted_at: "2026-07-05T18:11:04Z"
tags:
  - hacker-news
  - translated
---

# I built a weather station that runs machine learning to forecast weather

- HN: [48796441](https://news.ycombinator.com/item?id=48796441)
- Source: [github.com](https://github.com/Dominic-Muscatella/weather-station-alpha/tree/master)
- Score: 2
- Comments: 1
- Posted: 2026-07-05T18:11:04Z

## Translation

タイトル: 機械学習を実行して天気を予測する気象ステーションを構築しました
記事のタイトル: GitHub - Dominic-Muscatella/weather-station-alpha: ローカル データのみを使用した荒天警報用の raspberry pi ベースのシステム。 · GitHub
説明: ローカル データのみを使用して悪天候を警告するための Raspberry PI ベースのシステム。 - ドミニク・マスカテラ/気象観測所アルファ

記事本文:
GitHub - Dominic-Muscatella/weather-station-alpha: ローカル データのみを使用して荒天を警告するための raspberry pi ベースのシステム。 · GitHub
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
ドミニク・マスカテラ
/
気象観測所アルファ
公共
通知
通知を変更するにはサインインする必要があります

アイケーション設定
追加のナビゲーション オプション
コード
ドミニク・マスカテラ/気象観測所アルファ
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット ビルド ビルド データ データ エンジン エンジン イメージ イメージ live_work live_work model_package_lstm_attn model_package_lstm_attn static static system_management system_management LICENSE ライセンス config.py config.py Cooling_control.py Cooling_control.py data_pipeline.py data_pipeline.py light_control.py light_control.py live_engine.py live_engine.py live_log_local.json live_log_local.json oria_wa150km.c oria_wa150km.c Output.csv Output.csv readme.md readme.md refs.csv refs.csv要件.txt要件.txtscaler_LOCAL.npzscaler_LOCAL.npz sensor_reader_writer.py sensor_reader_writer.pythresholds.jsonthresholds.json ui_server.py ui_server.pyunits.pyunits.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカルデータを活用した悪天候予測
このプロジェクトには 4 つの主要な基本部分があります。
1: USB 気圧計とリモートセンサーからのセンサーデータをログし、CSV に書き込みます。
2: 過去の天気、監視、警告データに基づいてトレーニングされた 1 組の機械学習モデル アンサンブル。 1 人は 1 時間先の未来を見るように訓練され、もう 1 人は 24 時間先の未来を見るように訓練されました。どちらも KNN コンポーネントを備えており、基礎となるニューラル ネットワークを再トレーニングすることなくリアルタイム学習を実現します。
これらのモデルは、ローカルで収集されたデータに対して 1 時間ごとに実行されます。
3: 警告しきい値の調整、リアルタイム学習用のサンプルの管理、収集されたデータの表示、および悪天候モデルの出力の表示を行うための Web サーバーと UI
4: アクティブ冷却と照明のためのシステムおよびハードウェア制御
120v AC ->5v 3a DC 電源 (マイクロ USB 用)
2x 1.5 インチ USB 冷却ファン: 5v 0.162A
エレクトロニクスサロン リレーシールド f

またはラズベリー p1 モデル 8541715736
オベットデジタル気象観測所
ベースユニット型式：B66
スターパス USB BARO
ASIN B0CYJGZS8H
Nooelec RTL-SDR NESDR スマート - v5 バンドル
モデル番号 100700
ZulKit ブルーメタルプロジェクトボックス
7.9インチ×6.5インチ×3.5インチ
実行可能で以下を備えたコンピューター。
ブラウザ
ubuntu サーバー 22.04 (LTS) イメージを保持するのに十分なメモリ
システム
まず、Ubuntu Server 22.04 (lts) で microSD をフラッシュします。
ユーザー名が設定されていること、および Wi-Fi と SSH が有効になっていることを確認してください。
microSD、USB フラッシュ ドライブ、USB Baro、および USB SDR を接続します。
リレーシールドを追加し、5VをNO2とNO1に接続し、3.3VをNC1に接続します。冷却ファンをアースに配線し、ファン 1 のプラスのリードをリレー 1 のコモンに接続し、ファン 2 のプラスのリードをリレー 2 のコモンに接続します。
最後に、ws2812b チップを 5V、GND に配線し、そのデータ接点をボード上のピン D13 に接続します。 (これには PWM ピンを使用する必要があります)
すべての接続を再確認してください。電源がサポートできる限界に達しており、接続が緩んでいるとデバイスの電圧が不足してしまいます。
pi の電源をオンにして、cloud-init が完了するまで待ちます。これは最初の起動時にのみ発生しますが、時間がかかる場合があります。
メインマシンで、pi に SSH 接続します。
次のコマンドで、<username> を pi イメージに設定したユーザー名に置き換えます。
ssh < ユーザー名 > @ < ip.of.your.pi >
ユーザーのパスワードを入力し、Enter キーを押します。中に入ったら、いくつかのセットアップを行います。
sudo aptアップデート
sudo apt install -y git
次に、このリポジトリのクローンを作成します。
cd /home/ < ユーザー名 > /
git clone https://github.com/Dominic-Muscatella/weather-station-alpha.git
rtl_433 ライブラリのクローンを作成します。ネイティブ ライブラリは選択したセンサーの湿度チャネルをネイティブに取得しないため、少し変更してソースからコンパイルする必要があります。
cd /home/ < ユーザー名 > /
git clone https://github.com/merbanan/rtl_433.g

それ
ファイル oria_wa150km.c を /home/<username>/weather-station-alpha/ から rtl_433 ライブラリにコピーし、そこにあるファイルを私たちのファイルと置き換えます
cp /home/ < ユーザー名 > /weather-station-alpha/oria_wa150km.c /home/ < ユーザー名 > /rtl_433/src/devices/oria_wa150km.c
依存関係をインストールする
sudo apt install -y cmake librtlsdr-dev libusb-1.0-0-dev build-essential cmake pkg-config rtl-sdr software-properties-common
それを再コンパイルします。
mkdir -p /home/ < ユーザー名 > /rtl_433/build && cd /home/ < ユーザー名 > /rtl_433/build
cmake ..
make -j2
次に、Obet センサーが稼働しており、USB SDR が接続されて機能していると仮定して、次のコマンドを実行してテストします。
/home/ < ユーザー名 > /rtl_433/build/src/rtl_433 -R 288 -F json
センサーの出力が 30 秒ごとに、適切にフォーマットされた JSON として表示され始めるはずです。
ctrl+c で rtl_433 テストを終了します
次に、他のシステムの依存関係をインストールする必要があります。
sudo apt install -y software-properties-common
sudo add-apt-repository ppa:deadsnakes/ppa
sudo aptアップデート
sudo apt install -y python3.10 python3.10-dev python3.10-venv
次に、センサー監視、機械学習、Web サーバー スクリプト用に Python 仮想環境を設定する必要があります。
cd /home/ < ユーザー名 > /weather-station-alpha
python3.10 -m venv 天気環境
ソース Weather_env/bin/activate
python3.10 -m pip install --upgrade pip && python3.10 -m pip install -rrequirements.txt
このコマンドを使用して、接続されている USB デバイスを検索します。
sudo lsusb
どの接続デバイスがバロメーターであるかを決定します。
/home/<username>/weather-station-alpha/sensor_reader_writer.py を、接続されている baro USB のパスと一致するように変更します。
コンパイルされた rtl_433 の適切なパスも使用していることを確認してください。
次に、フラッシュドライブのUSBポートを特定します。 ext4 用にフォーマットし、 /mnt/DeepData/ にマウントします。 ubuntを変更する

u fs-tab ファイルを使用してドライブを永続的にマウントします。
完了したら、python3.10 sensor_reader_writer.py を実行すると、入力が表示されるはずです。
Ctrl+C を押してセンサー スクリプト テストを終了します。
スクリプトが仮想環境をアクティブ化し、システムの再起動時にバックグラウンドで確実に実行されるようにするために、個別の systemd ユニット ファイルを実装します。
サービス構成ファイルを作成します。
sudo nano /etc/systemd/system/weather-sensors.service
対応するファイルを /home/<username>/weather-station-alpha/system_management/services からコピーします。
必ず自分のユーザー名に変更してください。
機械学習サービスのセットアップ
エンジン構成ファイルを作成します。
sudo nano /etc/systemd/system/weather-engine.service
対応するファイルを /home/<username>/weather-station-alpha/system_management/services からコピーします。
必ず自分のユーザー名に変更してください。
sudo nano /etc/systemd/system/weather-ui.service
対応するファイルを /home/<username>/weather-station-alpha/system_management/services からコピーします。
必ず自分のユーザー名に変更してください。
冷却ファンサービスのセットアップ
sudo nano /etc/systemd/system/cooling_manager.service
対応するファイルを /home/<username>/weather-station-alpha/system_management/services からコピーします。
必ず自分のユーザー名に変更してください。
ライトコントロールサービスの設定
sudo nano /etc/systemd/system/light_manager.service
対応するファイルを /home/<username>/weather-station-alpha/system_management/services からコピーします。
必ず自分のユーザー名に変更してください。
最後に、デーモンに新しいサービスを通知し、カーネルが起動するたびに実行されるようにそれらをバインドする必要があります。
sudo systemctl デーモン-リロード
sudo systemctl は、weather-sensors.service を有効にします。
sudo systemctl は、weather-engine.service を有効にします
sudo systemctl は、weather-ui.service を有効にします
sudo systemctlでlight_manager.serviceを有効にする
sudo systemctl e

Cooling_manager.serviceを有効にする
sudo systemctl start Weather-sensors.service
sudo systemctl start Weather-engine.service
sudo systemctl start Weather-ui.service
sudo systemctl start light_manager.service
sudo systemctl start Cooling_manager.service
デバッグに役立ちます:
サービスのライブログを読むには:
sudojournalctl -u Weather-sensors.service -f
sudojournalctl -u Weather-engine.service -f
sudojournalctl -u Weather-ui.service -f
sudoジャーナルctl -u light_manager.service -f
sudojournalctl -u Cooling_manager.service -f
必要に応じてサービスを停止するには、次の操作を行います。
sudo systemctl stop Weather-sensors.service
sudo systemctl stop Weather-Engine.service
sudo systemctl stop Weather-ui.service
sudo systemctl stop light_manager.service
sudo systemctl stop Cooling_manager.service
完全に停止するのではなく再起動するには:
sudo systemctl 再起動 Weather-sensors.service
sudo systemctl 再起動 Weather-engine.service
sudo systemctl 再起動weather-ui.service
sudo systemctl 再起動 light_manager.service
sudo systemctl 再起動 Cooling_manager.service
すべての仕組み
Raspberry Pi は Ubuntu サーバーを実行します。
この機能は、systemctl が管理する 5 つのサービス (Weather Sensors サービス、Weather Engine サービス、Weather UI サービス、Light Manager サービス、Cooling Manager サービス) に分かれています。
気象センサー サービスは、433 mHz の無線を介して外部センサーから温度と湿度を収集し、車載気圧計から地域の気圧を収集します。このデータは 5 分ごとに保存されます。
Weather Engine サービスはローカル データに対して機械学習モデルを 15 分ごとに実行し、Weather UI サービスはデータ UI の API とフロントエンドをホストします。
Light Manager サービスは ws2812b チップを操作します。スタートアップを再生します

アニメーションとシャットダウンアニメーション。
操作中、次の場合に異なるアニメーションが表示されます。
全て晴れ（監視も警告もなし）
ウォッチ (1 つ以上のウォッチがアクティブ)
アドバイザリー (1 つ以上のアドバイザリーがアクティブ)
警告 (1 つ以上の警告がアクティブです)
コンピューティング (CPU コンピューティングが 30% を超えている)
1回（センサー1からのデータ受信時）
2回（センサー2からのデータ受信時）
5回（センサーデータをCSV保存する場合）
25 回、素早く (ml モデルが開始され、完了したとき)
Cooling Manager サービスは、CPU 使用率に基づいて、冷却ファンを 3 つのモードで動作させます。
デフォルト - 低 (ファン 1 @ 3.3v、ファン 2 オフ)
CPU 75% 以上 - 中 (ファン 1 @ 3.3v、ファン 2 @ 5v)
CPU 92% 以上 - 高 (ファン 1 @ 5v、ファン 2 @ 5v)
BTX3 ワイヤレス センサーは、その ID、検出された温度、湿度を 433 mHz の無線で送信します。これらの送信は、センサーごとに 30 秒ごとに行われます。
rtl_433 ライブラリは USB SDR と通信し、433 mHz 無線帯域のトラフィックを監視します。送信が検出されると、バイトがデコードされ、データが適切な形式の JSON オブジェクトとして返されます。
rtl_433 は、初期設定のままでは温度と ID のみをデコードできます。

[切り捨てられた]

## Original Extract

A raspberry pi based system for severe weather warnings using local data only. - Dominic-Muscatella/weather-station-alpha

GitHub - Dominic-Muscatella/weather-station-alpha: A raspberry pi based system for severe weather warnings using local data only. · GitHub
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
Dominic-Muscatella
/
weather-station-alpha
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Dominic-Muscatella/weather-station-alpha
master Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits build build data data engine engine images images live_work live_work model_package_lstm_attn model_package_lstm_attn static static system_management system_management LICENSE LICENSE config.py config.py cooling_control.py cooling_control.py data_pipeline.py data_pipeline.py light_control.py light_control.py live_engine.py live_engine.py live_log_local.json live_log_local.json oria_wa150km.c oria_wa150km.c output.csv output.csv readme.md readme.md refs.csv refs.csv requirements.txt requirements.txt scaler_LOCAL.npz scaler_LOCAL.npz sensor_reader_writer.py sensor_reader_writer.py thresholds.json thresholds.json ui_server.py ui_server.py units.py units.py View all files Repository files navigation
Severe Weather Prediction Using Local Data
This project has 4 main basic parts:
1: Logging the sensor data from the USB barometer, and the remote sensors, writing them to csv
2: A pair of machine learning model ensembles trained on historical weather, watch, and warning data. One trained to look 1h into the future, the other trained to look 24h ahead into the future. Both feature a KNN component, for real-time learning without retraining the underlying neural net.
These models are run against the locally gathered data once every hour.
3: A web server and UI for adjusting warning thresholds, administering the sample for real time learning, viewing the collected data, and viewing the output of the severe weather models
4: system and hardware controls for active cooling and lights
120v AC ->5v 3a DC power supply for micro usb
2x 1.5" USB cooling fan: 5v 0.162A
Electronics Salon relay shield for raspberry p1 model 8541715736
Obet Digital Weather Station
Base unit model : B66
Starpath USB BARO
ASIN B0CYJGZS8H
Nooelec RTL-SDR NESDR Smart - v5 Bundle
MODEL NUMBER 100700
ZulKit Blue Metal Project Box
7.9"x6.5"x3.5"
A computer that can run and has:
A browser
Enough memory to hold the ubuntu server 22.04 (LTS) image
System
First, flash the microSD with Ubuntu Server 22.04 (lts).
Make sure your username is set, and that the wifi and ssh are enabled.
Plug in the microSD, the USB flash drive, the USB Baro, and the USB SDR.
add the relay shield, and connect 5v to NO2 and NO1, and connect 3.3v to NC1. wire the cooling fans to ground, and put fan1's positive lead on relay 1's common, and put fan2's positive lead on relay 2's common.
finally, wire the ws2812b chip to 5v, gnd, and connect it's data contact to pin D13 on the board. (we need to use a PWM pin for this)
double check all your connections. we're at the edge of what the power supply can support, and loose connections will cause the the device to undervolt!
Power on the pi and WAIT for cloud-init to complete. this only happens on the first boot, but can take a long time.
On your main machine, ssh into the pi
In the following commands, replace <username> with the username you put on the pi image.
ssh < username > @ < ip.of.your.pi >
enter your user's password, and hit enter. Once inside, we will do some setup.
sudo apt update
sudo apt install -y git
then, clone this repo:
cd /home/ < username > /
git clone https://github.com/Dominic-Muscatella/weather-station-alpha.git
Clone the rtl_433 library. We have to modify it a bit and compile it from source, because the native library doesnt natively pick up the humidity channel of our chosen sensors.
cd /home/ < username > /
git clone https://github.com/merbanan/rtl_433.git
copy the file oria_wa150km.c from /home/<username>/weather-station-alpha/ , into the rtl_433 library, replacing the file that is there with ours
cp /home/ < username > /weather-station-alpha/oria_wa150km.c /home/ < username > /rtl_433/src/devices/oria_wa150km.c
install the dependancies
sudo apt install -y cmake librtlsdr-dev libusb-1.0-0-dev build-essential cmake pkg-config rtl-sdr software-properties-common
then recompile it.
mkdir -p /home/ < username > /rtl_433/build && cd /home/ < username > /rtl_433/build
cmake ..
make -j2
Then, assuming your Obet sensors are up and the USB SDR is plugged in and functioning, we test it by running this command:
/home/ < username > /rtl_433/build/src/rtl_433 -R 288 -F json
It should start showing the outputs of your sensors, every 30 seconds or so, as nice, formatted, json.
ctrl+c to exit the rtl_433 test
Next, we need to install dependancies for the other systems
sudo apt install -y software-properties-common
sudo add-apt-repository ppa:deadsnakes/ppa
sudo apt update
sudo apt install -y python3.10 python3.10-dev python3.10-venv
then, we need to set up our python virtual env for the sensor monitoring, machine learning and webserver scripts.
cd /home/ < username > /weather-station-alpha
python3.10 -m venv weather_env
source weather_env/bin/activate
python3.10 -m pip install --upgrade pip && python3.10 -m pip install -r requirements.txt
use this command to find connected usb devices:
sudo lsusb
determine which connected device is the barometer.
Modify /home/<username>/weather-station-alpha/sensor_reader_writer.py to match the path of our attached baro usb.
Make sure it is also using the proper path for our compiled rtl_433.
Then, identify the flash drive usb port. Format it for ext4, and mount it at /mnt/DeepData/ . Modify the ubuntu fs-tab file to perminantly mount the drive.
once complete, running python3.10 sensor_reader_writer.py should show the inputs comming in.
ctrl+c to exit the sensor script test.
To ensure the scripts activate their virtual environments and run reliably in the background upon system reboot, we implement individual systemd unit files.
Create the service configuration file:
sudo nano /etc/systemd/system/weather-sensors.service
copy the corresponding file from /home/<username>/weather-station-alpha/system_management/services
make sure to change to your username.
Setting Up The Machine Learning Service
Create the engine configuration file:
sudo nano /etc/systemd/system/weather-engine.service
copy the corresponding file from /home/<username>/weather-station-alpha/system_management/services
make sure to change to your username.
sudo nano /etc/systemd/system/weather-ui.service
copy the corresponding file from /home/<username>/weather-station-alpha/system_management/services
make sure to change to your username.
Setting Up The Cooling Fan service
sudo nano /etc/systemd/system/cooling_manager.service
copy the corresponding file from /home/<username>/weather-station-alpha/system_management/services
make sure to change to your username.
Setting Up The Light Control Service
sudo nano /etc/systemd/system/light_manager.service
copy the corresponding file from /home/<username>/weather-station-alpha/system_management/services
make sure to change to your username.
Finally, we need to tell the daemon about the new services, and bind them to run whenever the kernel boots.
sudo systemctl daemon-reload
sudo systemctl enable weather-sensors.service
sudo systemctl enable weather-engine.service
sudo systemctl enable weather-ui.service
sudo systemctl enable light_manager.service
sudo systemctl enable cooling_manager.service
sudo systemctl start weather-sensors.service
sudo systemctl start weather-engine.service
sudo systemctl start weather-ui.service
sudo systemctl start light_manager.service
sudo systemctl start cooling_manager.service
helpful for debugging:
to read the live logs of the services:
sudo journalctl -u weather-sensors.service -f
sudo journalctl -u weather-engine.service -f
sudo journalctl -u weather-ui.service -f
sudo journalctl -u light_manager.service -f
sudo journalctl -u cooling_manager.service -f
to stop the services if you ever need to:
sudo systemctl stop weather-sensors.service
sudo systemctl stop weather-engine.service
sudo systemctl stop weather-ui.service
sudo systemctl stop light_manager.service
sudo systemctl stop cooling_manager.service
and to restart them instead of stopping them completely:
sudo systemctl restart weather-sensors.service
sudo systemctl restart weather-engine.service
sudo systemctl restart weather-ui.service
sudo systemctl restart light_manager.service
sudo systemctl restart cooling_manager.service
How It All Works
The Raspberry Pi Runs Ubuntu Server.
The functions are broken into 5 services, which systemctl manages: the Weather Sensors service, the Weather Engine service, the Weather UI service, the Light Manager service, and the Cooling Manager service.
The Weather Sensors service gathers temperature and humidity from our outside sensors from over the radio at 433mHz, as well as gathering the local pressure from an on-board barometer. It saves this data every 5 minutes.
The Weather Engine service runs the machine learning models on the local data every 15 minutes, and the Weather UI service hosts the API and front end for the data UI.
The Light Manager service will operate the ws2812b chip. it will play a startup animation, and a shut down animation.
During operation, it will show different animations for the following cases:
all clear (no watches or warnings)
watch (one or more watches active)
advisory (one or more advisories active)
warning (one or more warnings active)
compute (cpu compute is over 30%)
once (when recieving data from sensor 1)
twice (when recieving data from sensor 2)
5 times (when saving sensor data to the csv)
25 times, quickly (when the ml models have started, and completed)
The Cooling Manager service will operate the cooling fans in 3 modes, based on cpu usage:
default - low (fan1 @ 3.3v, fan2 off)
cpu over 75% - med (fan1 @ 3.3v, fan2 @ 5v)
cpu over 92% - high (fan1 @ 5v, fan2 @ 5v)
the BTX3 wireless sensor transmits it's identity, the temperature detected , and the humidity over the radio at 433 mHz. These transmission happen every 30 seconds, for each sensor.
The rtl_433 library talks to the usb SDR, and monitors the 433 mHz radio band for traffic. When a transmission is detected, it decodes the bytes, and returns the data as nice, formatted JSON objects.
Out of the box, rtl_433 could decode only the termperature and identity, because the variant of the protocol that initially w

[truncated]
