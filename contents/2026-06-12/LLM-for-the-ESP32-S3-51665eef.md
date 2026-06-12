---
source: "https://github.com/harmansingh4163-ai/ESP-32-s3-Story-maker-LLM"
hn_url: "https://news.ycombinator.com/item?id=48500424"
title: "LLM for the ESP32-S3"
article_title: "GitHub - harmansingh4163-ai/ESP-32-s3-Story-maker-LLM: 15M/42M-param Llama split across two ESP32-S3s over 3 wires — too big for either chip alone. INT4, flash mmap, bit-exact verified. · GitHub"
author: "Harman-Singh123"
captured_at: "2026-06-12T07:44:55Z"
capture_tool: "hn-digest"
hn_id: 48500424
score: 1
comments: 0
posted_at: "2026-06-12T05:56:02Z"
tags:
  - hacker-news
  - translated
---

# LLM for the ESP32-S3

- HN: [48500424](https://news.ycombinator.com/item?id=48500424)
- Source: [github.com](https://github.com/harmansingh4163-ai/ESP-32-s3-Story-maker-LLM)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T05:56:02Z

## Translation

タイトル: ESP32-S3 用 LLM
記事のタイトル: GitHub - Harmansingh4163-ai/ESP-32-s3-Story-maker-LLM: 15M/42M-param Llama が 3 つのワイヤーを介して 2 つの ESP32-S3 に分割されました。どちらのチップだけでも大きすぎます。 INT4、フラッシュ mmap、ビット正確確認済み。 · GitHub
説明: 15M/42M パラメータの Llama は、3 つのワイヤを介して 2 つの ESP32-S3 に分割されます。どちらのチップも単独では大きすぎます。 INT4、フラッシュ mmap、ビット正確確認済み。 - harmansingh4163-ai/ESP-32-s3-ストーリーメーカー-LLM

記事本文:
GitHub - Harmansingh4163-ai/ESP-32-s3-Story-maker-LLM: 15M/42M-param Llama が 3 つのワイヤを介して 2 つの ESP32-S3 に分割されています。どちらのチップだけでも大きすぎます。 INT4、フラッシュ mmap、ビット正確確認済み。 · GitHub
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
ハーマンシン4163-ai
/
ESP-32-s3-ストーリーメーカー-

LLM
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
harmansingh4163-ai/ESP-32-s3-ストーリーメーカー-LLM
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット core core docs docs pc_tools pc_tools スケッチ スケッチ テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md パーティション.csv パーティション.csv すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1 つの言語モデル、2 つのマイクロコントローラー。 Llama アーキテクチャ LLM
2 つの ESP32-S3 ボードにレイヤーを分割して実行します。トークンごとに、
アクティベーション ベクトルは、チップ間の 3 本のワイヤ (CRC フレーム UART) を通過します。
> 昔々、勇敢な小さなキツネがいました
猫と魚は友達です。彼らは池で遊んだり泳いだりするのが好きです。
ある日、猫は木に止まっている鳥を見つけます...
[2 つのボード全体で 1.4 tok/s]
公開されたプロジェクトが示す限り、これは最初のマルチチップパイプライン化です。
ESP32 クラスのハードウェアでの LLM 推論。
16MB ESP32-S3 は、最大 15M パラメータ モデル (INT4) に達します。次は
TinyStories のサイズ — ストーリー 42M、~24MB — 単一のボードには適合しません。分割
2 つのチップにまたがる層では、チップではなく組み合わせたフラッシュが限界になります。
重み: INT4 (グループ 32 スケール)、メモリマップされたフラッシュからストリーミング
パーティション — 重み付けに使用される RAM の 0 バイト
計算: INT8 アクティベーション、整数完全群ドット積、matmul
各チップの両方の LX7 コアにわたって行を分割
分割: ワーカーボードはレイヤー 0 ～ K を実行します。ヘッドボードは埋め込みを保持し、
レイヤー K ～ L、分類子、トークナイザー、および次の各トークンのサンプル
リンク: UART @460800、フレーム = A5 5A | cmd |レン |ペイロード | CRC16、
~1.2KB/トークン往復 (トークン時間の~3%)
ハードウェアなしでテストできるものはすべて、フラッシュする前にテストされました (/tests を参照)。
フォワードパスは ~3e-7 (INT4 および INT8) への NumPy 参照と一致します。の

分割パイプラインはモノリシック モデルと比較してビット正確です。リンクプロトコル
ファズされました - ノイズは無視され、破損したフレームは CRC によって拒否されました。
パラメータ
重みが住んでいます
速度
出力
既知の ESP32 LLM ポート
260K
RAM
19–33 トーク/秒
文レベルの雑談
これ、シングルボード
15M
フラッシュ (mmap)
~1.4トーク/秒
複数段落の物語
これ、ボード2枚
現在 1500 万 / 目標 4200 万
スプリットフラッシュ
~1.4 / 推定 0.5 トーク/秒
まだ良い
一貫性は段階的ではない: TinyStories の調査でプロットの一貫性が示される
数百万のパラメータの範囲で現れる
(エルダン & リー 2023)。
ハードウェア: シングルボード用の 16MB フラッシュ + 8MB PSRAM を備えた ESP32-S3 1 台
モード;そのうちの 2 つ + パイプライン用のジャンパー線 3 本。検証日: Waveshare
ESP32-S3-Touch-LCD-5 (ヘッド) および Guition JC3248W535C (ワーカー)。二人
ボードは同じモデルである必要はありません。ディスプレイは v1 では使用されていません - すべて
相互作用は USB シリアル経由で行われ、設計により画面は暗いままになります。
PC (1 回限りのセットアップ、Windows コマンドを表示):
winget Python.Python.3.12 をインストールする
:: cmd を閉じて、新しい cmd を開き、次のようにします。
Python --バージョン
pip インストール numpy esptool
esp32 ボード パッケージを備えた Arduino IDE (Boards Manager → "esp32" by
エスプレッシフ; v3.x を推奨しますが、v2.x も動作します。スケッチには互換シムが含まれています)。
インストール後に「Python が見つかりませんでした」と表示されましたか? Windowsの設定→「アプリの管理」
実行エイリアス」→ python.exe と python3.exe をオフにして、cmd を再度開きます。
https://huggingface.co/karpathy/tinyllamas/resolve/main/stories15M.bin (~60MB)
https://raw.githubusercontent.com/karpathy/llama2.c/master/tokenizer.bin
(右クリック → リンクを名前を付けて保存 → 名前は tokenizer.bin のままにします)
cd パス\to\repo\pc_tools
python export_model.py stories15M.bin tokenizer.bin full.bin --bits 4
出力: full.bin (~10MB) — INT4 フラッシュ イメージ。
2A.シングルボードのストーリーテラー (ここから開始、約 15 分)
core\llm_core.c 、 core\llm_core.h 、およびpartitをコピーします

ions.csv へ
スケッチ\ストーリーテラー\ 。
Arduino IDEでesp32_storyteller.inoを開きます。ツール設定:
ESP32S3 開発モジュール · 起動時の USB CDC: 有効 · フラッシュ サイズ: 16MB ·
PSRAM：OPI・CPU 240MHz。アップロード。
モデルをフラッシュします (最初にシリアル モニターを閉じます。COMx は [ツール] → [ポート] にあります)。
python -m esptool --chip esp32s3 --port COMx write_flash 0x1F0000 full.bin
成功 = 「データのハッシュが検証されました。」
4. シリアル モニター @ 115200、ボードのリセット ボタンを押し、しばらく待ちます。
準備ができました。ストーリーの冒頭を入力し、Enter キーを押します。
モデルを分割します (ワーカーはレイヤー 0 ～ 2 を取得し、ヘッドは埋め込み + 3 ～ 5 を取得します)。
python split_image.py full.bin 3 worker.bin head.bin
ファームウェア: core\* とpartitions.csvを両方にコピーします
スケッチ\pipeline_head\ とスケッチ\pipeline_worker\ 。アップロード
Pipeline_head.ino をボード A に、pipeline_worker.ino をボード B に
(2A と同じツール設定)。
ピンとボー — 各スケッチの core/pipeline_link.h のコピーを編集します。
必要です。各ボードの LINK_TX_PIN / LINK_RX_PIN はそれぞれのボードの LINK_TX_PIN / LINK_RX_PIN と一致する必要があります
配線; 2 つのボードのピン番号は互いに一致する必要はありません。
LINK_BAUD は両方で同一でなければなりません。デフォルト: 17/18 @460800。で
Waveshare 5 インチ (空き GPIO ヘッダーなし) は I2C 端子ブロックを使用します。
TX = GPIO8 (SDA)、RX = GPIO9 (SCL)。
半分ずつフラッシュします (各ボード独自の COM ポート、シリアル モニターは閉じています):
python -m esptool --chip esp32s3 --port COM_head write_flash 0x1F0000 head.bin
python -m esptool --chip esp32s3 --port COM_worker write_flash 0x1F0000 worker.bin
ワイヤ - 3 つのジャンパー、ボードの電源がオフ、TX↔RX が交差:
頭部 TX ピン ──→ 作業者 RX ピン
ヘッド RX ピン ←── ワーカー TX ピン
GND ───── GND (各ボード上の任意の GND ピン。3V3/VCC には接続しないでください)
実行: 両方に電源を供給します (頭は PC に接続し、作業者は任意の USB 電源に接続します)。シリアル
HEAD のポート @115200 を監視し、両方のボードをリセットして、待機します。
準備完了: emb + 3

ローカル レイヤーが合計 6 つある場合は、プロンプトを入力します。
3. ランタイム コマンド (シリアル モニターに入力)
コマンド
効果
任意のテキスト
テキストに続くストーリーを生成します
/温度0.7
低い = 集中、高い (1.0+) = ワイルド
/topp 0.9
核サンプリングカットオフ
/レン250
プロンプトあたりの最大トークン数
/統計
空きRAM/PSRAMと現在の設定
4. stories42M（2枚のボードが必要なモデル）へのアップグレード
ステータス: エクスポートパスが実装され、規模が予算化されています。 15M パイプラインは
ハードウェア検証済み。測定された 4,200 万の数値は問題を通じて歓迎されます。
:: 同じ HuggingFace ページから stories42M.bin (~170MB) をダウンロードし、次のようにします。
python export_model.py stories42M.bin tokenizer.bin full42.bin --bits 4 --gs 32 --seq 224
python split_image.py full42.bin 7 worker42.bin head42.bin
python -m esptool --chip esp32s3 --port COM_head write_flash 0x1F0000 head42.bin
python -m esptool --chip esp32s3 --port COM_worker write_flash 0x1F0000 worker42.bin
--seq 224 が必要です (42M のネイティブ 1024 トークン コンテキストには
~29MB KV キャッシュ — 8MB PSRAM 上。 224 適合)。 3 ではなく 7 で分割します。
約 9MB の埋め込みはヘッド上に存在するため、レイヤーはワーカーに偏っています。チェックする
印刷される画像サイズは 14.6MB 未満に抑えられます。ブート バナーには次の内容が表示されます。
「emb + 1 ローカル レイヤー 8」（ヘッド）および「7 ローカル レイヤー、dim 512」（ワーカー）。
~0.4 ～ 0.7 tok/s が期待でき、明らかにより良いストーリーが得られます。人員配置
起動時に失敗→ --seq 192 で再エクスポート。
症状
修正
llm_init -1
フラッシュに古いモデルがありません - 0x1F0000 で esptool ステップをやり直します
「モデル」パーティションがありません
Partitions.csv は適用されません - スケッチ フォルダーにコピーし、パーティション スキームをカスタムに設定し、再アップロードします
ブートログはあるがバナーがない
起動時に「ツール」→「USB CDC」を反転し、再アップロードします
PSRAMが見つかりません
ツール → PSRAM: OPI (または QSPI)、再アップロード
リンクタイムアウト
TX/RX の一方の端が入れ替わっている、または GND ワイヤが欠落している
crc エラー、常に再試行します
ワイヤーを短くする

両方のボードで LINK_BAUD 115200
esptool: そのようなファイルはありません
.bin を含むフォルダーに cd します。
esptool: ポートがビジー状態
シリアルモニターを閉じる
リポジトリのレイアウト
コア/推論エンジン + リンクプロトコル (ホスト検証済み、ポータブル C)
スケッチ/Arduino ファームウェア (ビルド前に core/* +partitions.csv をコピー)
pc_tools/モデル量子化器/エクスポーターとレイヤースプリッター
テスト/証明: NumPy 参照、ビット正確性、およびプロトコル ファジー テスト
Partitions.csv 14MB モデル パーティションを備えた 16MB フラッシュ レイアウト
ロードマップ
ストーリー42M はハードウェアで測定
マークされた matmul スロット ( llm_matmul_rows ) の PIE SIMD — 推定 2 ～ 3×
MITライセンス。その後のエンジンアーキテクチャ
llama2.c (アンドレイ・カルパシー、マサチューセッツ工科大学);
TinyStories でトレーニングされたモデル
(エルダン & リー、マイクロソフト リサーチ)。と協力して開発されたコード
クロード (人間)。ハードウェア、統合、デバッグは私が担当しました。
15M/42M パラメータの Llama は、3 つのワイヤを介して 2 つの ESP32-S3 に分割されます。どちらのチップだけでも大きすぎます。 INT4、フラッシュ mmap、ビット正確確認済み。
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

15M/42M-param Llama split across two ESP32-S3s over 3 wires — too big for either chip alone. INT4, flash mmap, bit-exact verified. - harmansingh4163-ai/ESP-32-s3-Story-maker-LLM

GitHub - harmansingh4163-ai/ESP-32-s3-Story-maker-LLM: 15M/42M-param Llama split across two ESP32-S3s over 3 wires — too big for either chip alone. INT4, flash mmap, bit-exact verified. · GitHub
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
harmansingh4163-ai
/
ESP-32-s3-Story-maker-LLM
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
harmansingh4163-ai/ESP-32-s3-Story-maker-LLM
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits core core docs docs pc_tools pc_tools sketches sketches tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md partitions.csv partitions.csv View all files Repository files navigation
One language model, two microcontrollers. A Llama-architecture LLM
running with its layers split across two ESP32-S3 boards — per token, the
activation vector crosses three wires (CRC-framed UART) between the chips.
> Once upon a time there was a brave little fox
cat and fish are friends. They like to play and swim in the pond.
One day, the cat sees a bird on a tree...
[1.4 tok/s across 2 boards]
As far as published projects show, this is the first multi-chip pipelined
LLM inference on ESP32-class hardware.
A 16MB ESP32-S3 caps out at a ~15M-parameter model (INT4). The next
TinyStories size — stories42M, ~24MB — fits no single board . Splitting
layers across two chips makes combined flash the limit, not the chip.
Weights: INT4 (group-32 scales), streamed from a memory-mapped flash
partition — 0 bytes of RAM used for weights
Compute: INT8 activations, integer-exact group dot products, matmul
rows split across both LX7 cores of each chip
Split: worker board runs layers 0–K; head board holds embedding,
layers K–L, classifier, tokenizer, and samples each next token
Link: UART @460800, frame = A5 5A | cmd | len | payload | CRC16 ,
~1.2KB/token round trip (~3% of token time)
Everything testable without hardware was tested before flashing (see /tests):
forward pass matches a NumPy reference to ~3e-7 (INT4 and INT8); the
split pipeline is bit-exact vs the monolithic model; the link protocol
was fuzzed — noise ignored, corrupted frames rejected by CRC.
params
weights live in
speed
output
known ESP32 LLM ports
260K
RAM
19–33 tok/s
sentence-level babble
this, single board
15M
flash (mmap)
~1.4 tok/s
multi-paragraph stories
this, two boards
15M now / 42M target
split flash
~1.4 / est. 0.5 tok/s
better still
Coherence isn't gradual: TinyStories research shows plot consistency
emerges in the millions-of-parameters range
( Eldan & Li 2023 ).
Hardware: one ESP32-S3 with 16MB flash + 8MB PSRAM for single-board
mode; two of them + 3 jumper wires for the pipeline. Verified on: Waveshare
ESP32-S3-Touch-LCD-5 (head) and Guition JC3248W535C (worker). The two
boards do not have to be the same model. Displays are unused in v1 — all
interaction is over USB serial, screens stay dark by design.
PC (one-time setup, Windows commands shown):
winget install Python.Python.3.12
:: close cmd, open a NEW one, then:
python --version
pip install numpy esptool
Arduino IDE with the esp32 board package (Boards Manager → "esp32" by
Espressif; v3.x recommended, v2.x works — the sketches include a compat shim).
"Python was not found" after installing? Windows Settings → "Manage app
execution aliases" → turn OFF python.exe and python3.exe, reopen cmd.
https://huggingface.co/karpathy/tinyllamas/resolve/main/stories15M.bin (~60MB)
https://raw.githubusercontent.com/karpathy/llama2.c/master/tokenizer.bin
(right-click → Save link as → keep the name tokenizer.bin )
cd path\to\repo\pc_tools
python export_model.py stories15M.bin tokenizer.bin full.bin --bits 4
Output: full.bin (~10MB) — the INT4 flash image.
2A. Single-board storyteller (start here, ~15 min)
Copy core\llm_core.c , core\llm_core.h , and partitions.csv into
sketches\storyteller\ .
Open esp32_storyteller.ino in Arduino IDE. Tools settings:
ESP32S3 Dev Module · USB CDC On Boot: Enabled · Flash Size: 16MB ·
PSRAM: OPI · CPU 240MHz . Upload.
Flash the model (close Serial Monitor first; COMx is in Tools → Port):
python -m esptool --chip esp32s3 --port COMx write_flash 0x1F0000 full.bin
Success = "Hash of data verified."
4. Serial Monitor @ 115200 , press the board's reset button, wait for
Ready , type a story opening, press Enter.
Split the model (worker gets layers 0–2; head gets embedding + 3–5):
python split_image.py full.bin 3 worker.bin head.bin
Firmware: copy core\* and partitions.csv into BOTH
sketches\pipeline_head\ and sketches\pipeline_worker\ . Upload
pipeline_head.ino to board A and pipeline_worker.ino to board B
(same Tools settings as 2A).
Pins & baud — edit each sketch's copy of core/pipeline_link.h if
needed. Each board's LINK_TX_PIN / LINK_RX_PIN must match its own
wiring ; the two boards' pin numbers do NOT have to match each other.
LINK_BAUD MUST be identical on both. Defaults: 17/18 @460800. On the
Waveshare 5" (no free GPIO header) use the I2C terminal block:
TX = GPIO8 (SDA), RX = GPIO9 (SCL).
Flash each half (each board's own COM port, Serial Monitor closed):
python -m esptool --chip esp32s3 --port COM_head write_flash 0x1F0000 head.bin
python -m esptool --chip esp32s3 --port COM_worker write_flash 0x1F0000 worker.bin
Wire — 3 jumpers, boards powered off, TX↔RX crossed:
head TX-pin ──→ worker RX-pin
head RX-pin ←── worker TX-pin
GND ─────────── GND (any GND pin on each board; do NOT connect 3V3/VCC)
Run: power both (head on the PC; worker on any USB power). Serial
Monitor on the HEAD's port @115200, reset both boards, wait for
Ready: emb + 3 local layers of 6 total , type a prompt.
3. Runtime commands (typed into Serial Monitor)
Command
Effect
any text
generates a story continuing your text
/temp 0.7
lower = focused, higher (1.0+) = wilder
/topp 0.9
nucleus sampling cutoff
/len 250
max tokens per prompt
/stats
free RAM/PSRAM and current settings
4. Upgrading to stories42M (the model that needs two boards)
Status: export path implemented and size-budgeted; the 15M pipeline is
hardware-verified. Measured 42M numbers welcome via issues.
:: download stories42M.bin (~170MB) from the same HuggingFace page, then:
python export_model.py stories42M.bin tokenizer.bin full42.bin --bits 4 --gs 32 --seq 224
python split_image.py full42.bin 7 worker42.bin head42.bin
python -m esptool --chip esp32s3 --port COM_head write_flash 0x1F0000 head42.bin
python -m esptool --chip esp32s3 --port COM_worker write_flash 0x1F0000 worker42.bin
--seq 224 is required (42M's native 1024-token context would need a
~29MB KV cache — over the 8MB PSRAM; 224 fits). Split at 7 , not 3:
the ~9MB embedding lives on the head, so layers skew to the worker. Check
the printed image sizes stay under 14.6MB. Boot banners should read
"emb + 1 local layers of 8" (head) and "7 local layers, dim 512" (worker).
Expect ~0.4–0.7 tok/s — and clearly better stories. Worker allocation
failure at boot → re-export with --seq 192 .
Symptom
Fix
llm_init -1
no/old model in flash — redo the esptool step at 0x1F0000
no 'model' partition
partitions.csv not applied — copy into the sketch folder, set Partition Scheme: Custom, re-upload
boot log but no banner
flip Tools → USB CDC On Boot, re-upload
PSRAM not found
Tools → PSRAM: OPI (or QSPI), re-upload
link timeout
TX/RX swapped at one end, or GND wire missing
crc error, retry constantly
shorten wires or set LINK_BAUD 115200 on BOTH boards
esptool: No such file
cd into the folder containing the .bin
esptool: port busy
close Serial Monitor
Repo layout
core/ inference engine + link protocol (host-verified, portable C)
sketches/ Arduino firmware (copy core/* + partitions.csv in before building)
pc_tools/ model quantizer/exporter and the layer splitter
tests/ the proof: NumPy-reference, bit-exactness, and protocol fuzz tests
partitions.csv 16MB flash layout with the 14MB model partition
Roadmap
stories42M measured on hardware
PIE SIMD in the marked matmul slot ( llm_matmul_rows ) — est. 2-3×
MIT License. Engine architecture after
llama2.c (Andrej Karpathy, MIT);
models trained on TinyStories
(Eldan & Li, Microsoft Research). Code developed in collaboration with
Claude (Anthropic); hardware, integration, and debugging by me.
15M/42M-param Llama split across two ESP32-S3s over 3 wires — too big for either chip alone. INT4, flash mmap, bit-exact verified.
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
