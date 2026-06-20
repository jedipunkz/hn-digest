---
source: "https://github.com/jackulau/ratchet"
hn_url: "https://news.ycombinator.com/item?id=48604403"
title: "Show HN: Ratchet – let an AI agent reflash your BIOS over a CH341A (MCP server)"
article_title: "GitHub - jackulau/ratchet: Modern BIOS chip programmer and debugger (CH341A / CH347) — fully native Rust · GitHub"
author: "JackLau"
captured_at: "2026-06-20T00:57:12Z"
capture_tool: "hn-digest"
hn_id: 48604403
score: 1
comments: 0
posted_at: "2026-06-19T23:13:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ratchet – let an AI agent reflash your BIOS over a CH341A (MCP server)

- HN: [48604403](https://news.ycombinator.com/item?id=48604403)
- Source: [github.com](https://github.com/jackulau/ratchet)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T23:13:26Z

## Translation

タイトル: Show HN: Ratchet – AI エージェントに CH341A (MCP サーバー) 経由で BIOS を再フラッシュさせます
記事タイトル: GitHub - jackulau/ratchet: 最新 BIOS チップ プログラマおよびデバッガ (CH341A / CH347) — 完全ネイティブ Rust · GitHub
説明: 最新の BIOS チップ プログラマおよびデバッガ (CH341A / CH347) - 完全ネイティブ Rust - jackulau/ratchet

記事本文:
GitHub - jackulau/ratchet: 最新 BIOS チップ プログラマおよびデバッガ (CH341A / CH347) — 完全ネイティブ Rust · GitHub
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
ジャクラウ
/
ラチェット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
94 コミット 94 コミット .github/ workflows .github/ workflows docs docs パッケージング パッケージング Rust Rust タスク タスク .gitignore .gitignore ライセンス ライセンス README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
CH341A および CH347 USB プログラマ用のハードウェア デバッグおよびフラッシュ プログラミング ツールキット。
その核心は SPI フラッシュ プログラミングと BIOS 分析であり、ライブ シリコンをエンドツーエンドで推進するパスです。その周囲には、単体テスト済みのプロトコル層 (I2C、UART、1-Wire、パッシブ SPI スニフ、JTAG、SWD、CAN)、ターゲット MCU プログラマー (AVR ISP、STK500 / Arduino ブートローダー、24Cxx および 93xxx EEPROM、ESP32 / ESP8266、STM32 over SWD、および AN3155 UART)、ARM が配置されています。デバッグ サーフェス (ADIv5、Cortex-M 停止/再開/ステップ、ELF 対応メモリ ピーク)、JTAG IDCODE および BSDL 境界スキャン、Saleae / sigrok エクスポートを備えたロジック アナライザー モデル、および Bus Pirate / slcan CAN ブリッジ。
そのすべてがまだ稼働中のハードウェアに接続されているわけではありません。ステータスは、すべてのコマンドに [ライブ] 、 [オフライン] 、または [n/w] (有線ではない) をマークします。全体のルール: ライブ トランスポートのないコマンドは、明確なメッセージとともにゼロ以外で終了します。決して成功を装うことはありません。
Rust で書かれています。カスタム libusb FFI と手動の JSON-RPC MCP サーバーを備えた単一の自己完結型バイナリで、Node や Python ランタイムはありません。これは、SPI フラッシュ パス上の AsProgrammer と NeoProgrammer を置き換え、フラッシュROM、avrdude、esptool、stm32flash、および OpenOCD に分割されていた領域 (ネイティブ USB、画像分析、診断のナレッジ ベース、AI エージェント用の組み込み MCP サーバー) をカバーします。
プレリリース。 GitHub リリースはまだ公開されていないため、現在の唯一のインストール ルートはソースからのものです (「インストール」を参照)。
現在のライブ ハードウェアを推進しているもの (CLI + MCP):
SPI フラッシュ + BIOS、エンドツーエンド。ステータス、検出、識別、読み取り、書き込み

e、verify、erase、region-erase、blank-check、sfdp、wp-status、full-repair、および full-backup は、ライブ CH341A または CH347 に対して実行されます。 write は、影響を受けるセクターを最初に消去し (SPI プログラムはビット 1→0 のみをクリアできます)、ページごとにプログラムし、消去とプログラムのたびに進行中の書き込み (WIP) ビットをポーリングして、ビジーなチップと競合することがないようにし、書き込み前の自動バックアップを取得し、検証のために読み戻します。 16 MB を超えるチップは、4 バイトのアドレス指定に自動的に切り替わります。完全修復はガイド付きパイプラインを実行します。 full-backup は、名前付きファイルへのフルチップ読み取りです。
I2C、CH341A ビットバンまたは CH347 ネイティブ経由。 i2c scan 、 i2c read 、 i2c write 、および eeprom-i2c read/write (24Cxx) は、ライブ バス上で実際の Ch341aI2c / Ch347I2c マスターを使用します。
JTAG IDCODE スキャン、CH347 JTAG エンジン経由 (CH341A にはありません)。 jtag idcode-scan は実際の Ch347Jtag アダプターを駆動します。
Backend auto-select. open_default() は CH347 ( 1a86:55db )、次に CH341A ( 1a86:5512 ) をプローブし、stderr 警告を表示してモックに戻ります。 RATCHET_FORCE_MOCK=1 はモックを強制します。プロトコル動詞は open_raw_bus() を使用します。これは、デバイスが存在しない場合にサイレント モック フォールバックではなく正直なエラーを返します。ラチェット ステータスは、バックエンド JSON フィールドでアクティブなバックエンドを報告します。
ハードウェアを必要としないオフライン ツール:
i2c sniff <trace.json> は、キャプチャされた (t_us、scl、sda) トレースをデコードします。 jtag bsdl-scan <file.bsdl> は BSDL ファイルを解析し、その境界レジスタを報告します。 la export <capture.json> <out> --format csv|jsonl はキャプチャを変換します。 serial-list はシリアル ポートを列挙します (POSIX)。 repl は、SPI バックエンド上の stdin REPL です。加えて、分析動詞のanalyze、diff、checksum、chip-info、search、post-decode、およびvoltage-referenceを追加します。
まだライブハードウェアに接続されていません: uart 、 onewire 、 swd 、 avr 、 eeprom-microwire 、 esp 、 stm32 、 la Capture 、 Buspirate 、 can 、 plus moni

tor 、シリアル接続、および障害検索。それぞれのプロトコル ロジックは実装され、モックに対して単体テストされていますが、ライブ CH341A/CH347 トランスポート アダプタはまだ接続されていません (SWD / 1-Wire / AVR-ISP / Microwire ビットバン、ネイティブ UART RX、外部シリアル/CAN デバイス)。これらはゼロ以外で終了し (または JSON-RPC エラーを返し)、偽の成功ではありません。
破壊的なパスは強化されています。短い USB 読み取りは、サイレントゼロパディングの代わりにハードエラーになります。書き込み保護されたシリコン、容量不明のチップ、サイレントに選択された模擬バックエンドの消去と書き込みを拒否します。 4 バイト モードは使用後は常に終了します。単一のチップ選択アサーション内の全範囲読み取りストリーム。 SPI 書き込みパスは、CH341A USB フレーミング (全二重読み取り、AND フラッシュ セマンティクスによる消去/プログラム) の背後で SPI NOR チップをエミュレートする LoopbackFlash テスト バスによってハードウェアを使用せずに証明されているため、書き込み → 読み取り → 検証のラウンドトリップがエンドツーエンドで実行されます。モック バックエンドは、デバイスが接続されていない場合でも、SPI フラッシュ サーフェスを CI で実行可能な状態に保ちます。 472 の単体テストと統合テストに合格しました。
供給源から貨物を経由して（現在機能しているパス）
Rust 1.82+ と libusb-1.0 がインストールされている必要があります (「要件」を参照)。
git clone https://github.com/jackulau/ratchet
CDラチェット/錆び
カーゴインストール --path ラチェット-cli
カーゴインストール --path ratchet-mcp
これにより、ratchet と ratchet-mcp が ~/.cargo/bin/ (または設定されている場合は CARGO_INSTALL_ROOT の値) にインストールされます。どちらのバイナリも自己完結型の Rust 実行可能ファイルです。ノードもPythonもありません。
グローバル状態を混在させたくない場合は、サンドボックス ディレクトリにインストールします。
カーゴインストール --path ratchet-cli --root /opt/ratchet
カーゴインストール --path ratchet-mcp --root /opt/ratchet
エクスポート PATH = " /opt/ratchet/bin: $PATH "
チェックアウトから (インストールなし)
git clone https://github.com/jackulau/ratchet
CDラチェット/錆び
カーゴビルド --release
#B

インナリーは target/release/ratchet および target/release/ratchet-mcp に着地します。
アンインストール
カーゴアンインストールラチェット-cli
カーゴアンインストールラチェット-mcp
インストール中に --root /opt/ratchet を使用した場合は、同じルートを渡します。
カーゴ アンインストール ratchet-cli --root /opt/ratchet
カーゴ アンインストール ratchet-mcp --root /opt/ratchet
チェックアウトから構築した場合
アンインストールするものは何もありません。クローン作成したディレクトリを削除します。オプションで:
cd ラチェット/錆び && カーゴクリーン # ビルドアーティファクトを削除
cd .. && rm -rf ratchet # チェックアウトを削除します
Claude Desktop MCP 登録の削除
~/Library/Application Support/Claude/claude_desktop_config.json (macOS) または同等のプラットフォームを編集し、 mcpServers の下の「ratchet」エントリを削除します。クロードデスクトップを再起動します。
これは、破損した、またはブリックされたマザーボード BIOS を再フラッシュするためのエンドツーエンドのパスです。
CH341A (一般的な ~$3 プログラマ) または CH347。
CH341A または CH347 USB プログラマーおよび SOIC-8 / SOIC-16 テスト クリップ (または、ZIF アダプター)
チップのはんだを除去します）。 BIOS フラッシュは、チップセットの近くにある 8 ピン SPI チップで、通常は Winbond です。
W25Q… 、Macronix MX25L… 、または GigaDevice GD25Q… 。
正確なボード リビジョンに対応する正常な BIOS イメージ - マザーボードからダウンロードします
ベンダーのサポート ページを参照するか、手順 2 でラチェットが作成したバックアップを保管しておいてください。
電圧チェック: ほとんどの BIOS チップは 3.3 V (純正の CH341A が駆動する電圧) です。一部は 1.8 V です。
レベルシフターアダプターが必要です - ラチェットチップ情報 <チップ> がチップの電圧を報告するので、
接続する前に確認してください。
手順 (ボードの電源をオフにし、プラグを外した状態でチップにクリップします):
# 1. プログラマとチップが通信していることを確認します。
ラチェット ステータス # プログラマーが検出されましたか?どのバックエンドですか？
ラチェット識別 # JEDEC ID を読み取り、806 チップ DB でチップを検索します
# 2. BIOS が無効になっているように見える場合でも、最初に現在の内容をバックアップします。
ラチェットリードバック

kup.bin # フルチップダンプ → ファイル
ラチェット分析backup.bin # オプション: UEFIボリューム、MEリージョン、整合性
# 3. 正常なイメージをフラッシュします。これは自動的に次のようになります。
# • 現在のチップのタイムスタンプ付きバックアップを保存します。
# • 影響を受けるセクターを消去し、ページごとにプログラムします。
# • 操作のたびに進行中の書き込みビットをポーリングし、
# • チップを読み取り、それがファイルと一致することを確認します。
ラチェット書き込み new_bios.bin
# 4. 独立して再検証します (オプション - `write` はすでに検証されています)。
ラチェットは new_bios.bin を確認します
write は、all-0xFF または all-0x00 イメージ (チップを消去する空のダンプまたは失敗したダンプ) を拒否します。
チップより大きい画像は拒否されます。書き込み中に何か問題が発生した場合、オリジナルは
ステップ 3 で印刷されたタイムスタンプ付きバックアップ。不良フラッシュの後にボードをリカバリするには、次のようにします。
ラチェットはそのファイルからbackup.binを書き込みます。
ワンショットパイプライン。 ratchet full-repair --reference new_bios.bin がすべてを実行します -
接続品質チェック → 読み取り二重検証 → 健全性分析 → 修復 → 書き込み → 書き込み後
検証 - 単一のガイド付きワークフローとして。
CI ( .github/workflows/ci.yml : fmt、clippy -D 警告、完全なテスト スイート、厳密なドキュメント
ビルドし、 RATCHET_FORCE_MOCK=1 の下のスモーク スイートとテスト スイートの両方で、
プログラマなしでバイトごとにプロトコルを実行します (「ステータス」を参照) が、確認のため
自分のボードに対して:
ラチェット検出 # プログラマーが USB で列挙する
ラチェット識別 --json | jq .data # JEDEC ID はチップのシルクスクリーン/DB と一致します
ラチェット読み取り a.bin && ラチェット読み取り b.bin && diff a.bin b.bin # 2 つの読み取りは同一です (安定したクリップ)
ラチェット書き込み new_bios.bin # 出力で success=trueverified=true
クイックスタート - マルチプロトコル
これらは現在稼働中のハードウェアを駆動します (または、記載されている場合はオフラインで実行されます)。それぞれが
ゼロ以外の終了と、デバイスが存在しない場合の正直なメッセージ

nt。
# I2C バススキャン + レジスタリード (ライブ CH341A / CH347)
ラチェット i2c スキャン
ラチェット i2c 読み取り --addr 0x50 --reg 0x00 --len 256
# 24Cxx I2C EEPROM ダンプ/リストア (ライブ)
ラチェット eeprom-i2c 読み取り --addr 0x50 --part 24c256 dump.bin
ラチェット eeprom-i2c 書き込み --addr 0x50 --part 24c256 dump.bin
# JTAG IDCODE チェーン (ライブ、CH347 のみ)
ラチェット jtag idcode-scan
# オフライン: キャプチャされた I2C トレースをデコードする / BSDL ファイルを解析する / キャプチャを変換する
ラチェットi2cスニフtrace.json
ラチェット jtag bsdl-scan part.bsdl
ラチェット ラ エクスポート Capture.json out.csv --format csv
# シリアルポートを列挙する (POSIX)
ラチェットシリアルリスト
ライブトランスポートがまだ接続されていない動詞 ( uart 、 onewire 、 swd 、 avr 、
eeprom-microwire 、 esp 、 stm32 、 la Capture 、 Buspirate 、 can ) exit
ゼロ以外の説明付き (「ステータス」を参照)。
ratchet --help は、39 のトップレベルのサブコマンドと help を公開します。ステータスの凡例:
[ライブ] はハードウェアを駆動します (デバイスがない場合は正直なエラー)、[オフライン] は必要ありません
ハードウェア、[n/w] ライブ トランスポートにまだ接続されていません (ゼロ以外で終了、決して終了しません)
成功を偽装します）。
すべての検査コマンドは、AgentEnvelope 出力の --json をサポートしています。
{ok、コマンド、データ?|エラー、nextAction?} 。
ラチェットステータス --json
ラチェットチップ情報 ef4017 --json
ラチェット分析backup.bin --json | jq ' .data.regions '
エージェントインターフェース

[切り捨てられた]

## Original Extract

Modern BIOS chip programmer and debugger (CH341A / CH347) — fully native Rust - jackulau/ratchet

GitHub - jackulau/ratchet: Modern BIOS chip programmer and debugger (CH341A / CH347) — fully native Rust · GitHub
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
jackulau
/
ratchet
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
94 Commits 94 Commits .github/ workflows .github/ workflows docs docs packaging packaging rust rust tasks tasks .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
A hardware debug and flash-programming toolkit for CH341A and CH347 USB programmers.
Its core is SPI flash programming and BIOS analysis, the path that drives live silicon end to end. Around it sits a unit-tested protocol layer (I2C, UART, 1-Wire, passive SPI sniff, JTAG, SWD, CAN), target-MCU programmers (AVR ISP, STK500 / Arduino bootloader, 24Cxx and 93xxx EEPROM, ESP32 / ESP8266, STM32 over SWD and AN3155 UART), an ARM debug surface (ADIv5, Cortex-M halt/resume/step, ELF-aware memory peek), JTAG IDCODE and BSDL boundary scan, a logic-analyzer model with Saleae / sigrok export, and Bus Pirate / slcan CAN bridges.
Not all of that is wired to live hardware yet. Status marks every command [live] , [offline] , or [n/w] (not wired). The rule throughout: a command with no live transport exits non-zero with a clear message; it never fakes success.
Written in Rust: a single self-contained binary with custom libusb FFI and a hand-rolled JSON-RPC MCP server, no Node or Python runtime. It replaces AsProgrammer and NeoProgrammer on the SPI-flash path, and covers ground otherwise split across flashrom, avrdude, esptool, stm32flash, and OpenOCD: native USB, image analysis, a knowledge base of diagnostics, and a built-in MCP server for AI agents.
Pre-release. No GitHub Releases are published yet, so the only install route today is from source (see Install ).
What drives live hardware today (CLI + MCP):
SPI flash + BIOS, end to end. status , detect , identify , read , write , verify , erase , region-erase , blank-check , sfdp , wp-status , full-repair , and full-backup run against a live CH341A or CH347. write erases the affected sectors first (SPI program can only clear bits 1→0), programs page by page, polls the write-in-progress (WIP) bit after every erase and program so it never races a busy chip, takes an automatic pre-write backup, and reads back to verify. Chips larger than 16 MB switch to 4-byte addressing automatically. full-repair runs the guided pipeline; full-backup is a full-chip read to a named file.
I2C , over CH341A bit-bang or CH347 native. i2c scan , i2c read , i2c write , and eeprom-i2c read/write (24Cxx) use the real Ch341aI2c / Ch347I2c master over the live bus.
JTAG IDCODE scan , over the CH347 JTAG engine (CH341A has none). jtag idcode-scan drives the real Ch347Jtag adapter.
Backend auto-select. open_default() probes CH347 ( 1a86:55db ), then CH341A ( 1a86:5512 ), then falls back to mock with a stderr warning; RATCHET_FORCE_MOCK=1 forces mock. Protocol verbs use open_raw_bus() , which returns an honest error rather than a silent mock fallback when no device is present. ratchet status reports the active backend in its backend JSON field.
Offline tools that need no hardware:
i2c sniff <trace.json> decodes a captured (t_us, scl, sda) trace; jtag bsdl-scan <file.bsdl> parses a BSDL file and reports its boundary register; la export <capture.json> <out> --format csv|jsonl converts a capture; serial-list enumerates serial ports (POSIX); repl is a stdin REPL over the SPI backend; plus the analysis verbs analyze , diff , checksum , chip-info , search , post-decode , and voltage-reference .
Not wired to live hardware yet: uart , onewire , swd , avr , eeprom-microwire , esp , stm32 , la capture , buspirate , can , plus monitor , serial connect, and failure-search . Each one's protocol logic is implemented and unit-tested against a mock, but no live CH341A/CH347 transport adapter is wired for it yet (SWD / 1-Wire / AVR-ISP / Microwire bit-bang, native UART RX, external serial/CAN devices). They exit non-zero (or return a JSON-RPC error), never a fake success.
The destructive paths are hardened: a short USB read is a hard error instead of silent zero-padding; erase and write refuse write-protected silicon, unknown-capacity chips, and a silently-selected mock backend; 4-byte mode is always exited after use; whole-range reads stream inside a single chip-select assertion. The SPI write path is proven without hardware by a LoopbackFlash test bus that emulates an SPI NOR chip behind the CH341A USB framing (full-duplex reads, erase/program with AND-into-flash semantics), so a write → read-back → verify round-trip runs end to end. The mock backend keeps the SPI-flash surface exercisable in CI when no device is attached. 472 unit and integration tests pass.
From source via cargo (the path that works today)
Requires Rust 1.82+ and libusb-1.0 installed (see Requirements ).
git clone https://github.com/jackulau/ratchet
cd ratchet/rust
cargo install --path ratchet-cli
cargo install --path ratchet-mcp
This installs ratchet and ratchet-mcp into ~/.cargo/bin/ (or the value of CARGO_INSTALL_ROOT if set). Both binaries are self-contained Rust executables; no Node, no Python.
If you prefer not to mix global state, install to a sandbox directory:
cargo install --path ratchet-cli --root /opt/ratchet
cargo install --path ratchet-mcp --root /opt/ratchet
export PATH= " /opt/ratchet/bin: $PATH "
From a checkout (no install)
git clone https://github.com/jackulau/ratchet
cd ratchet/rust
cargo build --release
# Binaries land at target/release/ratchet and target/release/ratchet-mcp
Uninstall
cargo uninstall ratchet-cli
cargo uninstall ratchet-mcp
If you used --root /opt/ratchet during install, pass the same root:
cargo uninstall ratchet-cli --root /opt/ratchet
cargo uninstall ratchet-mcp --root /opt/ratchet
If you built from a checkout
Nothing to uninstall; delete the cloned directory. Optionally:
cd ratchet/rust && cargo clean # remove build artifacts
cd .. && rm -rf ratchet # remove the checkout
Removing the Claude Desktop MCP registration
Edit ~/Library/Application Support/Claude/claude_desktop_config.json (macOS) or the platform equivalent and remove the "ratchet" entry under mcpServers . Restart Claude Desktop.
This is the end-to-end path for reflashing a corrupt or bricked motherboard BIOS with a
CH341A (the common ~$3 programmer) or a CH347.
A CH341A or CH347 USB programmer and a SOIC-8 / SOIC-16 test clip (or a ZIF adapter if you
desolder the chip). The BIOS flash is the 8-pin SPI chip near the chipset, usually a Winbond
W25Q… , Macronix MX25L… , or GigaDevice GD25Q… .
A known-good BIOS image for your exact board revision - download it from the motherboard
vendor's support page, or keep the backup ratchet makes in step 2.
Voltage check: most BIOS chips are 3.3 V (what a stock CH341A drives). Some are 1.8 V and
need a level-shifter adapter - ratchet chip-info <chip> reports the chip's voltage so you can
check before connecting.
Steps (clip onto the chip with the board powered off and unplugged):
# 1. Confirm the programmer + chip are talking.
ratchet status # programmer detected? which backend?
ratchet identify # reads the JEDEC ID and looks the chip up in the 806-chip DB
# 2. Back up the current contents FIRST - always, even if the BIOS looks dead.
ratchet read backup.bin # full-chip dump → file
ratchet analyze backup.bin # optional: UEFI volumes, ME region, integrity
# 3. Flash the known-good image. This automatically:
# • saves a timestamped backup of the current chip,
# • erases the affected sectors, then programs page-by-page,
# • polls the write-in-progress bit after every operation, and
# • reads the chip back and verifies it matches the file.
ratchet write new_bios.bin
# 4. Re-verify independently (optional - `write` already verified).
ratchet verify new_bios.bin
write refuses an all-0xFF or all-0x00 image (a blank/failed dump that would wipe the chip) and
refuses an image larger than the chip. If anything goes wrong mid-write, your original is in the
timestamped backup printed by step 3. To recover a board after a bad flash, just
ratchet write backup.bin from that file.
One-shot pipeline. ratchet full-repair --reference new_bios.bin runs the whole thing -
connection-quality check → double-verify read → health analysis → repair → write → post-write
verify - as a single guided workflow.
CI ( .github/workflows/ci.yml : fmt, clippy -D warnings , full test suite, strict doc
build, and both smoke suites under RATCHET_FORCE_MOCK=1 ) and the test suite prove the
protocol byte-for-byte without a programmer (see Status ), but to confirm
against your own board:
ratchet detect # programmer enumerates on USB
ratchet identify --json | jq .data # JEDEC id matches the chip silk-screen / DB
ratchet read a.bin && ratchet read b.bin && diff a.bin b.bin # two reads are identical (stable clip)
ratchet write new_bios.bin # success=true verified=true in the output
Quick Start - multi-protocol
These drive live hardware today (or run offline where noted). Each returns a
non-zero exit and an honest message if no device is present.
# I2C bus scan + register read (live CH341A / CH347)
ratchet i2c scan
ratchet i2c read --addr 0x50 --reg 0x00 --len 256
# 24Cxx I2C EEPROM dump / restore (live)
ratchet eeprom-i2c read --addr 0x50 --part 24c256 dump.bin
ratchet eeprom-i2c write --addr 0x50 --part 24c256 dump.bin
# JTAG IDCODE chain (live, CH347 only)
ratchet jtag idcode-scan
# Offline: decode a captured I2C trace / parse a BSDL file / convert a capture
ratchet i2c sniff trace.json
ratchet jtag bsdl-scan part.bsdl
ratchet la export capture.json out.csv --format csv
# Enumerate serial ports (POSIX)
ratchet serial-list
Verbs whose live transport isn't wired yet ( uart , onewire , swd , avr ,
eeprom-microwire , esp , stm32 , la capture , buspirate , can ) exit
non-zero with an explanation (see Status ).
ratchet --help exposes 39 top-level subcommands plus help . Status legend:
[live] drives hardware (honest error if no device), [offline] needs no
hardware, [n/w] not wired to a live transport yet (exits non-zero, never
fakes success).
Every inspection command supports --json for AgentEnvelope output:
{ok, command, data?|error, nextAction?} .
ratchet status --json
ratchet chip-info ef4017 --json
ratchet analyze backup.bin --json | jq ' .data.regions '
Agent Interfa

[truncated]
