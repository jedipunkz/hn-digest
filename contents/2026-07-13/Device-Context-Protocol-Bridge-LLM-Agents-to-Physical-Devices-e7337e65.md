---
source: "https://github.com/device-context-protocol/dcp"
hn_url: "https://news.ycombinator.com/item?id=48889594"
title: "Device Context Protocol – Bridge LLM Agents to Physical Devices"
article_title: "GitHub - device-context-protocol/dcp: Device Context Protocol — bridge LLM agents to physical devices. Sub-50-byte frames, 27.6KB flash / 0.6KB RAM measured on ESP32, capability-scoped and safe by design. Complementary to MCP. Paper: arXiv:2605.26159 · GitHub"
author: "manchoz"
captured_at: "2026-07-13T08:49:13Z"
capture_tool: "hn-digest"
hn_id: 48889594
score: 2
comments: 0
posted_at: "2026-07-13T08:35:55Z"
tags:
  - hacker-news
  - translated
---

# Device Context Protocol – Bridge LLM Agents to Physical Devices

- HN: [48889594](https://news.ycombinator.com/item?id=48889594)
- Source: [github.com](https://github.com/device-context-protocol/dcp)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T08:35:55Z

## Translation

タイトル: デバイス コンテキスト プロトコル – LLM エージェントを物理デバイスにブリッジする
記事のタイトル: GitHub - device-context-protocol/dcp: デバイス コンテキスト プロトコル — LLM エージェントを物理デバイスにブリッジします。サブ 50 バイト フレーム、27.6KB フラッシュ / 0.6KB RAM（ESP32 で測定）、機能範囲が限定されており、設計により安全です。 MCP を補完します。論文: arXiv:2605.26159 · GitHub
説明: デバイス コンテキスト プロトコル — LLM エージェントを物理デバイスにブリッジします。サブ 50 バイト フレーム、27.6KB フラッシュ / 0.6KB RAM（ESP32 で測定）、機能範囲が限定されており、設計により安全です。 MCP を補完します。ペーパー: arXiv:2605.26159 - デバイスコンテキストプロトコル/dcp

記事本文:
GitHub - device-context-protocol/dcp: デバイス コンテキスト プロトコル — LLM エージェントを物理デバイスにブリッジします。サブ 50 バイト フレーム、27.6KB フラッシュ / 0.6KB RAM（ESP32 で測定）、機能範囲が限定されており、設計により安全です。 MCP を補完します。論文: arXiv:2605.26159 · GitHub
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
却下

警告
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
デバイスコンテキストプロトコル
/
DCP
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
57 コミット 57 コミット .github .github docs docs 例 例 ファームウェア/ esp32 ファームウェア/ esp32 src/ dcp src/ dcp テスト テスト ツール ツール Web/ パネルエミュ Web/ パネルエミュ .gitignore .gitignore .mcp.json.example .mcp.json.example CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md SPEC.md SPEC.md pyproject.toml pyproject.toml ビューすべてのファイル リポジトリ ファイルのナビゲーション
ステータス: ドラフト v0.3 — 2026 年 5 月 · ESP32-WROOM-32 でハードウェア検証済み
LLM エージェントが物理デバイスを安全に制御できるようにするプロトコル。
1ドルクラスのマイクロコントローラーに至るまで。
インテントレベル、トランスポートに依存しない、機能スコープ。コンパクトなワイヤフォーマット
(50 バイト未満のフレーム)。内蔵ファームウェア: 1 KB 未満の RAM、
約 28 KB のフラッシュ。
MCP を補完する — リファレンス
Bridge は DCP ↔ MCP を変換するため、任意の MCP ホスト (Claude Desktop、Claude Code、
IDE アシスタント) は設定不要で動作します。
レシピ - すぐにフラッシュできる 5 つのデバイス スケルトン
設計の理論的根拠: docs/RATIONALE.md — なぜなのか
MCP-on-MCU ではない、WoT ではない、Matter ではないのか。
MCP は SaaS ツールとしては優れていますが、WebSocket 上の JSON-RPC とランタイムを前提としています。
ツールの発見。 32 KB の RAM を搭載した MCU では、これはスターターではありません。
DCP は MCP のメンタル モデル (マニフェスト + ツール呼び出し) を保持しますが、次のとおりです。
コンパクトな CBOR ワイヤー形式にコンパイルします
静的インテント テーブルを使用します (ランタイム ネゴシエーションなし)
安全性の執行をブリッジプロセスに移す
参照ブリッジが DCP を変換します

↔ MCP 、つまり MCP 互換の LLM が動作します
箱から出して。 DCP は物理ハードウェアへのラストマイルです。
これが 1 つのグラフで重要となる理由: プロトコルのスキーマがその数を決定します。
幻覚や敵対的な通話は、バイトが到達する前に停止されます。
デバイス。 DCP は、ワイヤ層で 6 つのカテゴリすべてをキャッチします。他は
既存のスキーマが何をカバーしているのかを把握します。
登録ではなく意図。 set_brightness(50%) 、 write_pwm(pin=5,duty=128) ではありません。
プロトコル内の単位。すべての数値が単位を宣言します。曖昧さはありません。
静的インテントテーブル。マニフェストはコンパイル時に認識されます。ランタイムは純粋なバイナリです。
橋には安全が息づいています。デバイスはブリッジを信頼します。 LLM は生の GPIO を決して認識しません。
デフォルトでは冪等です。非冪等のインテントは、それ自体を宣言する必要があります。
トランスポートに依存しない。 UART、BLE、MQTT、USB-CDC、WebSocket — 1 フレーム。
ブリッジは唯一の信頼境界です。呼び出しごとに問題が発生し、
機能トークンを検証し、範囲/タイプ/単位のチェックを強制します。
マニフェストを作成し、ワイヤー形式のプリミティブとしてドライランをサポートします。デバイス
市販のマイクロコントローラーに適合するほどシンプルなままです。すべて
LLM が実行できることは、バイトがデータを通過する前に強制されます。
デバイスの境界。
v0.3 の時点で、リファレンス ファームウェアは 2 つの環境で測定および検証されています。
物理ボード — CH340 USB シリアル経由の ESP32-WROOM-32 開発ボード、
および S3 のネイティブ USB シリアル/JTAG 経由の ESP32-S3 (LILYGO T-Panel S3)
— どちらも 115 200 ボー:
各ボードで 13/13 往復テストに合格 ( tools/test_uart_roundtrip.py )
88/88 Python 単体テストと適合性テストに合格
フルランプファームウェア: 295 KB フラッシュ、WROOM-32 上の 22.7 KB グローバル — ほとんど
うち Arduino-ESP32 ランタイム + FreeRTOS であり、DCP ではありません
DCP レイヤー自体は 27.6 KB のフラッシュと 0.6 KB の RAM を測定します。
ベースラインの空のスケッチ上で — を使用して再現します
docs/paper/figures/measure_footprint.py 。の

フラッシュフィギュアは終わりました
元の 16 KB 未満のデザイン ターゲット (オンデバイス HMAC が使用される前に設定)
追加）; RAM の数値はそれをはるかに下回っています。
S3 の実行では、ネイティブ USB CDC リンクを介して DCP も実行されます。
USB-UART ブリッジ チップよりも — 同じファームウェアで、トランスポート固有のものはありません
コード
スタティック RAM は、MCU 上の希少なリソースです。 DCP層の測定
0.6 KB の RAM は、IoT-MCP の報告に比べて 2 桁大きい
74KBのピークメモリ。 DCP のフラッシュ コスト (27.6 KB、測定値) はプロットされていません
— IoT-MCP は、同等のフラッシュ数値を報告していません。
ハードウェアの内容については、docs/RATIONALE.md §7 を参照してください。
検証は証明しますが、証明しません。
ESP ファミリ全体でクリーンなクロスコンパイル (Xtensa + RISC-V + ESP8266)
リファレンス ファームウェアは設計上移植可能です (Arduino Stream +
ソフトウェア SHA-256、DCP に SoC 固有のコード パスはありません。{h,cpp} )。それ
現在のすべての ESP32 バリアントと ESP8266 をクロスコンパイルします。
これらのターゲットのうち 2 つは、実際のボードでも実行時検証されています。
残りはベンチ上でビルド検証が保留中のハードウェアです。
すべてのビルドは Arduino-ESP32 コア 3.3.8 / Arduino-ESP8266 コア 3.x を使用します。
同じファームウェア/esp32/ライブラリ。スケッチでは PWM API を選択します。
コンパイル時間 (ESP32 では ledcAttach / ledcWrite、analogWrite では
ESP8266);プロトコル層自体には #ifdef がありません。再現する
と:
arduino-cli コンパイル --clean --fqbn esp32:esp32:esp32c3 \
--ライブラリ ファームウェア/esp32 ファームウェア/esp32/examples/lamp
arduino-cli コンパイル --clean --fqbn esp8266:esp8266:nodemcuv2 \
--ライブラリ ファームウェア/esp32 ファームウェア/esp32/examples/lamp
マニフェスト
dcp : 0.3
デバイス:
id : ランプキッチン-01
モデル:smart_lamp_v1
ベンダー : example.dev
インテント:
- 名前 : set_brightness
パラメータ:
レベル: { タイプ: float、単位: パーセント、範囲: [0, 100] }
フェード : { タイプ: 期間、単位: ミリ秒、デフォルト: 0 }
機能 : ランプ書き込み
冪等 : true
dry_run : true
- 名前 : 読む

_明るさ
戻り値: { 型: 浮動小数点数、単位: パーセント }
機能 : ランプ読み取り
イベント:
- 名前 : モーション検出
ペイロード:
信頼度 : { 型: 浮動小数点数、単位: 比率、範囲: [0, 1] }
機能 : ランプ読み取り
tent_id = crc16(name) — マニフェストとファームウェアは同期を維持します。
コーディネート。
フレームは 6 バイトの固定ヘッダー + CBOR ペイロード + オプションの 16 バイトです
切り詰められた HMAC-SHA256。ヘッダーフィールド:
応答ステータス コード: ok 、拒否された、範囲、ビジー、不明な意図、機能要求。
一般的な set_brightness(50) 呼び出しは、回線上で 19 バイトになります。 MCP
JSON-RPC に相当するのは約 180 バイトです。完全な標準仕様
SPEC.md に存在します。
詳細については、docs/ADDING_FEATURES.md を参照してください。
動作するまばたき（回数、期間）の例を含む 5 ステップのループ。短い
バージョン: マニフェストを編集し、C++ ハンドラー + バインディングを追加し、再コンパイルします。
フラッシュ、MCP サーバーを再起動します - LLM が新しいツールを選択します
自動的に。 Bridge ではコードを変更する必要はありません。
# ユーザーとして — PyPI からインストールします。
pip install " pydcp[mcp,serial] " # またはすべてのトランスポートの [mcp,serial,mqtt,ble]
dcp Inspection Examples/lamp_manifest.yaml # 解析されたマニフェストの概要
dcp サーブの例/lamp_manifest.yaml --simulator
# 貢献者として — ソースから編集可能にインストール:
git clone https://github.com/device-context-protocol/dcp.git
CDDCP
pip install -e " .[mcp,serial,mqtt,ble,dev] "
pytest # 88 個のテストすべて
python Examples/lamp_demo.py # インプロセスブリッジ ↔ 偽のランプ
PyPI パッケージの名前は pydcp です (裸の dcp は、
無関係なパッケージ)。インポート名は dcp です。プロトコル名はDCPです。
リファレンス Bridge には、各 DCP インテントを
MCPツール。 --simulator を使用すると、インプロセスの偽デバイスが起動されるため、
ハードウェアなしでデモが可能。
dcpserve Examples/lamp_manifest.yaml --simulator # ハードウェアなし
DCP サーブの例

les/lamp_manifest.yaml --serial COM3 # UART 上の実際の ESP32
dcp サーブの例/lamp_manifest.yaml --mqtt Broker.lan:1883 \ # MQTT
--mqtt-prefix dcp/ランプキッチン
dcp サーブの例/lamp_manifest.yaml --ble AA:BB:CC:DD:EE:FF \ # BLE
--ble-service 12345678-1234-5678-1234-567812345678
ケイパビリティトークン(HMAC-SHA256)
マルチテナントまたはスコープ指定されたアクセスの場合、有効期間の短い HMAC トークンを作成して渡します。
橋へ:
エクスポート DCP_SECRET= $( dcp トークン keygen )
dcp トークン ミント --caps ランプ.書き込み、ランプ.読み取り --ttl 3600
# eyJjYXBzIjpb...sig
トークンは呼び出しごとにブリッジによって検証されます。デバイスはただ見るだけです
すでに承認されているフレーム。デバイス自体は署名を検証しません
v0.2 — これにはオンデバイス HMAC が必要ですが、これはロードマップにあります。
Claude Desktop に接続するには、これを
クロード_デスクトップ_config.json :
{
"mcpサーバー": {
「スマートランプ」: {
"コマンド" : " dcp " ,
"引数" : [
「サーブ」、
" C:/path/to/protocol/examples/lamp_manifest.yaml " 、
" --simulator "
]
}
}
}
次にクロードに「ランプの明るさを 60% に設定して」と頼みます。通話フロー:
クロード ─MCP─▶ DCP サーブ ─ブリッジ─▶ ループバック ─DCP ワイヤー─▶ GenericSimulator
運用環境で使用する場合は、GenericSimulator を実際のトランスポートに置き換えます。
(UART / MQTT / BLE — 次回登場)。
v0.3 にないもの (意図的)
マルチデバイスのアトミックトランザクション
メッシュルーティング (必要に応じて下にある Thread / Zigbee を使用してください)
LLM 側の認証 (MCP ホストのセッション モデルに委任)
ネイティブ CAN FD フレーム (ESP32-S3 TWAI はクラシック CAN、v0.4 ESP32-P4)
ポートは真の CAN FD を有効にします)
学術研究で DCP を使用する場合は、arXiv プレプリントを引用してください。
@misc { yang2026dcp 、
title = { デバイス コンテキスト プロトコル: コンパクトで安全第一のアーキテクチャ
制約されたデバイスの LLM 駆動制御用 } 、
著者 = {ヤン、ドンシュウ} 、
年 = { 2026 } 、
eprint = { 2605.26159 } 、
円弧

hivePrefix = { arXiv } 、
プライマリクラス = { cs.NI } 、
URL = { https://arxiv.org/abs/2605.26159 } 、
}
機械可読な CITATION.cff も提供されています — GitHub
サイドバーに「このリポジトリを引用」ボタンを表示します。
ループバック トランスポートを使用した Python ブリッジの参照
MCP サーバー ラッパー + CLI (dcpserve)
汎用インプロセスデバイスシミュレータ
UART トランスポート (COBS フレーム + CRC-16)
ESP32リファレンスファームウェア（Arduino互換C++）
設計の理論的根拠 ( docs/RATIONALE.md )
CI (GitHub アクション、Linux + Windows、py 3.11 ～ 3.13)
HMAC-SHA256 機能トークン (ブリッジ側の強制)
マニフェスト コンパイラ: dcp codegen (YAML → C ヘッダー)
ファームウェアのコンパイル時 DCP_ID(name) マクロ
リリース準備: CONTRIBUTING / CHANGELOG / CoC / SECURITY / issue template
オンデバイス HMAC 検証 (フレームごとの署名、ESP32 ファームウェア)
ESP32 BLEペリフェラル例（NimBLE-Arduino）
適合テスト スイート (ゴールデン フレーム、言語中立の YAML)
Codegen --stubs : ハンドラー署名 + バインディング テーブルを発行します
クイックスタート ビデオ スクリプト ( docs/QUICKSTART_VIDEO.md )
2 つのボード (ESP32-WROOM-32 以上) でのリアルハードウェア検証
CH340、ESP32-S3 / ネイティブ USB 経由の T パネル）、それぞれ 13/13 往復
ESP32 RISC-V ファミリ (C3、C6、H2、P4) および ESP8266 でのクロスコンパイル クリーン
device-context-protocol/dcp のパブリック リポジトリ (v0.3.0 リリース)

[切り捨てられた]

## Original Extract

Device Context Protocol — bridge LLM agents to physical devices. Sub-50-byte frames, 27.6KB flash / 0.6KB RAM measured on ESP32, capability-scoped and safe by design. Complementary to MCP. Paper: arXiv:2605.26159 - device-context-protocol/dcp

GitHub - device-context-protocol/dcp: Device Context Protocol — bridge LLM agents to physical devices. Sub-50-byte frames, 27.6KB flash / 0.6KB RAM measured on ESP32, capability-scoped and safe by design. Complementary to MCP. Paper: arXiv:2605.26159 · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
device-context-protocol
/
dcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
57 Commits 57 Commits .github .github docs docs examples examples firmware/ esp32 firmware/ esp32 src/ dcp src/ dcp tests tests tools tools web/ panel-emu web/ panel-emu .gitignore .gitignore .mcp.json.example .mcp.json.example CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SPEC.md SPEC.md pyproject.toml pyproject.toml View all files Repository files navigation
Status: Draft v0.3 — May 2026 · Hardware-validated on ESP32-WROOM-32
A protocol that lets LLM agents safely control physical devices,
down to dollar-class microcontrollers.
Intent-level, transport-agnostic, capability-scoped. Compact wire format
(sub-50-byte frames). Self-contained firmware: under 1 KB of RAM,
~28 KB of flash.
Complementary to MCP — a reference
Bridge translates DCP ↔ MCP so any MCP host (Claude Desktop, Claude Code,
IDE assistants) works zero-config.
Recipes — five ready-to-flash device skeletons
Design rationale: docs/RATIONALE.md — why
not MCP-on-MCU, why not WoT, why not Matter.
MCP is excellent for SaaS tools, but assumes JSON-RPC over WebSocket and runtime
tool discovery. On an MCU with 32 KB of RAM, that's a non-starter.
DCP keeps MCP's mental model (manifest + tool calls) but:
compiles to a compact CBOR wire format
uses a static intent table (no runtime negotiation)
moves safety enforcement to a Bridge process
A reference Bridge translates DCP ↔ MCP , so any MCP-compatible LLM works
out of the box. DCP is the last mile to physical hardware.
Why this matters in one chart: the protocol's schema decides how many
hallucinated or adversarial calls are stopped before any byte reaches a
device. DCP catches all six categories at the wire layer; the others
catch what their existing schema happens to cover.
Intent, not register. set_brightness(50%) , not write_pwm(pin=5, duty=128) .
Units in the protocol. Every number declares a unit. No ambiguity.
Static intent table. Manifest known at compile time; runtime is pure binary.
Safety lives in the Bridge. Devices trust the Bridge; LLMs never see raw GPIO.
Idempotent by default. Non-idempotent intents must declare themselves.
Transport-agnostic. UART, BLE, MQTT, USB-CDC, WebSocket — one frame.
The Bridge is the sole trust boundary. On every call it issues and
verifies capability tokens, enforces range/type/unit checks from the
manifest, and supports dry-run as a wire-format primitive. Devices
remain simple enough to fit on commodity microcontrollers; everything
the LLM is allowed to do is enforced before any byte traverses the
device boundary.
As of v0.3 the reference firmware is measured-validated on two
physical boards — an ESP32-WROOM-32 dev board over CH340 USB-Serial,
and an ESP32-S3 (LILYGO T-Panel S3) over the S3's native USB-Serial/JTAG
— both at 115 200 baud:
13/13 round-trip tests pass on each board ( tools/test_uart_roundtrip.py )
88/88 Python unit & conformance tests pass
Full lamp firmware: 295 KB flash, 22.7 KB globals on WROOM-32 — most
of which is the Arduino-ESP32 runtime + FreeRTOS, not DCP
The DCP layer itself measures 27.6 KB of flash and 0.6 KB of RAM
over a baseline empty sketch — reproduce with
docs/paper/figures/measure_footprint.py . The flash figure is over
the original <16 KB design target (set before on-device HMAC was
added); the RAM figure is well under it.
The S3 run also exercises DCP over a native-USB CDC link rather
than a USB-UART bridge chip — same firmware, no transport-specific
code
Static RAM is the scarce resource on an MCU. The DCP layer's measured
0.6 KB of RAM sits two orders of magnitude under IoT-MCP's reported
74 KB peak memory. DCP's flash cost (27.6 KB, measured) is not plotted
— IoT-MCP does not report a comparable flash figure.
See docs/RATIONALE.md §7 for what the hardware
validation does and does not prove.
Cross-compile clean across the ESP family (Xtensa + RISC-V + ESP8266)
The reference firmware is portable by design (Arduino Stream + a
software SHA-256, no SoC-specific code paths in DCP.{h,cpp} ). It
cross-compiles for every current ESP32 variant and for ESP8266;
two of those targets are also runtime-validated on real boards, the
rest are build-validated pending hardware on the bench:
All builds use Arduino-ESP32 core 3.3.8 / Arduino-ESP8266 core 3.x
the same firmware/esp32/ library. The sketch picks PWM API at
compile time ( ledcAttach / ledcWrite on ESP32, analogWrite on
ESP8266); the protocol layer itself has no #ifdef . Reproduce
with:
arduino-cli compile --clean --fqbn esp32:esp32:esp32c3 \
--library firmware/esp32 firmware/esp32/examples/lamp
arduino-cli compile --clean --fqbn esp8266:esp8266:nodemcuv2 \
--library firmware/esp32 firmware/esp32/examples/lamp
Manifest
dcp : 0.3
device :
id : lamp-kitchen-01
model : smart_lamp_v1
vendor : example.dev
intents :
- name : set_brightness
params :
level : { type: float, unit: percent, range: [0, 100] }
fade : { type: duration, unit: ms, default: 0 }
capability : lamp.write
idempotent : true
dry_run : true
- name : read_brightness
returns : { type: float, unit: percent }
capability : lamp.read
events :
- name : motion_detected
payload :
confidence : { type: float, unit: ratio, range: [0, 1] }
capability : lamp.read
intent_id = crc16(name) — manifests and firmware stay in sync without
coordination.
A frame is a 6-byte fixed header + CBOR payload + an optional 16-byte
truncated HMAC-SHA256. Header fields:
Reply status codes: ok , denied , range , busy , unknown_intent , capability_required .
A typical set_brightness(50) call is 19 bytes on the wire; the MCP
JSON-RPC equivalent is approximately 180 bytes. The full normative spec
lives at SPEC.md .
See docs/ADDING_FEATURES.md for the full
5-step loop with a worked blink(times, period) example. The short
version: edit the manifest, add a C++ handler + binding, recompile,
flash, restart the MCP server — the LLM picks up the new tool
automatically. The Bridge needs no code change.
# As a user — install from PyPI:
pip install " pydcp[mcp,serial] " # or [mcp,serial,mqtt,ble] for all transports
dcp inspect examples/lamp_manifest.yaml # parsed manifest summary
dcp serve examples/lamp_manifest.yaml --simulator
# As a contributor — editable install from source:
git clone https://github.com/device-context-protocol/dcp.git
cd dcp
pip install -e " .[mcp,serial,mqtt,ble,dev] "
pytest # all 88 tests
python examples/lamp_demo.py # in-process bridge ↔ fake lamp
The PyPI package is named pydcp (the bare dcp is squatted by an
unrelated package). The import name is dcp . The protocol name is DCP.
The reference Bridge ships an MCP server that exposes each DCP intent as an
MCP tool. With --simulator it spins up an in-process fake device, so you
can demo with no hardware.
dcp serve examples/lamp_manifest.yaml --simulator # no hardware
dcp serve examples/lamp_manifest.yaml --serial COM3 # real ESP32 over UART
dcp serve examples/lamp_manifest.yaml --mqtt broker.lan:1883 \ # MQTT
--mqtt-prefix dcp/lamp-kitchen
dcp serve examples/lamp_manifest.yaml --ble AA:BB:CC:DD:EE:FF \ # BLE
--ble-service 12345678-1234-5678-1234-567812345678
Capability tokens (HMAC-SHA256)
For multi-tenant or scoped access, mint short-lived HMAC tokens and pass them
to the Bridge:
export DCP_SECRET= $( dcp token keygen )
dcp token mint --caps lamp.write,lamp.read --ttl 3600
# eyJjYXBzIjpb...sig
Tokens are verified by the Bridge on every call. The device sees only
already-authorized frames. Devices themselves do not verify signatures
in v0.2 — that requires on-device HMAC, which is on the roadmap.
To wire it into Claude Desktop , add this to your
claude_desktop_config.json :
{
"mcpServers" : {
"smart-lamp" : {
"command" : " dcp " ,
"args" : [
" serve " ,
" C:/path/to/protocol/examples/lamp_manifest.yaml " ,
" --simulator "
]
}
}
}
Then ask Claude "set the lamp to 60% brightness" . The call flow:
Claude ─MCP─▶ dcp serve ─Bridge─▶ Loopback ─DCP wire─▶ GenericSimulator
For production use, replace GenericSimulator with a real transport
(UART / MQTT / BLE — coming next).
What's not in v0.3 (intentional)
Multi-device atomic transactions
Mesh routing (use Thread / Zigbee underneath if you need it)
LLM-side authentication (delegated to the MCP host's session model)
Native CAN FD frames (ESP32-S3 TWAI is classic CAN; v0.4 ESP32-P4
port enables true CAN FD)
If you use DCP in academic work, please cite the arXiv preprint:
@misc { yang2026dcp ,
title = { Device Context Protocol: A Compact, Safety-First Architecture
for LLM-Driven Control of Constrained Devices } ,
author = { Yang, Dongxu } ,
year = { 2026 } ,
eprint = { 2605.26159 } ,
archivePrefix = { arXiv } ,
primaryClass = { cs.NI } ,
url = { https://arxiv.org/abs/2605.26159 } ,
}
A machine-readable CITATION.cff is also provided — GitHub
renders a "Cite this repository" button in the sidebar.
Reference Python Bridge with loopback transport
MCP server wrapper + CLI ( dcp serve )
Generic in-process device simulator
UART transport (COBS framing + CRC-16)
ESP32 reference firmware (Arduino-compatible C++)
Design rationale ( docs/RATIONALE.md )
CI (GitHub Actions, Linux + Windows, py 3.11–3.13)
HMAC-SHA256 capability tokens (Bridge-side enforcement)
Manifest compiler: dcp codegen (YAML → C header)
Compile-time DCP_ID(name) macro in firmware
Release prep: CONTRIBUTING / CHANGELOG / CoC / SECURITY / issue templates
On-device HMAC verification (per-frame signatures, ESP32 firmware)
ESP32 BLE peripheral example (NimBLE-Arduino)
Conformance test suite (golden frames, language-neutral YAML)
Codegen --stubs : emits handler signatures + binding table
Quickstart video script ( docs/QUICKSTART_VIDEO.md )
Real-hardware validation on two boards (ESP32-WROOM-32 over
CH340, ESP32-S3 / T-Panel over native USB), 13/13 round-trips each
Cross-compile clean on ESP32 RISC-V family (C3, C6, H2, P4) and ESP8266
Public repo at device-context-protocol/dcp (v0.3.0 rele

[truncated]
