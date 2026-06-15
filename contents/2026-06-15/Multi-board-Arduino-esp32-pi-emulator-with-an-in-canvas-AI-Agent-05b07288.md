---
source: "https://velxio.dev/v3/"
hn_url: "https://news.ycombinator.com/item?id=48544496"
title: "Multi-board (Arduino, esp32,pi) emulator with an in-canvas AI Agent"
article_title: "Velxio 3.0 — Retro CPUs, MicroSD, ePaper & Multi-Board Embedded Simulator"
author: "noahnathan25"
captured_at: "2026-06-15T19:26:37Z"
capture_tool: "hn-digest"
hn_id: 48544496
score: 5
comments: 0
posted_at: "2026-06-15T17:32:50Z"
tags:
  - hacker-news
  - translated
---

# Multi-board (Arduino, esp32,pi) emulator with an in-canvas AI Agent

- HN: [48544496](https://news.ycombinator.com/item?id=48544496)
- Source: [velxio.dev](https://velxio.dev/v3/)
- Score: 5
- Comments: 0
- Posted: 2026-06-15T17:32:50Z

## Translation

タイトル: キャンバス内 AI エージェントを備えたマルチボード (Arduino、esp32、pi) エミュレーター
記事タイトル: Velxio 3.0 — レトロ CPU、MicroSD、電子ペーパー、マルチボード組み込みシミュレータ
説明: Velxio 3.0 は、プログラム可能なレトロ CPU (Z80、8080、4004、4040、8086)、MicroSD カード エミュレーション、電子ペーパー ディスプレイ、真のマルチボード UART/I2C/SPI 相互接続、完全な元に戻す/やり直し、および 100 以上の新しい例を追加します。無料、オープンソース、ブラウザベース — Arduino、ESP32、RP2040、STM32、Raspberry Pi など。

記事本文:
ベルキシオ
オフラインArduino、RP2040、ESP32シミュレーター
ローカル バックエンドを開始しています…
Velxio — 無料のオンライン回路および Arduino シミュレーター | SPICE・ESP32・RP2040・ATtiny85・カスタムチップ
無料のオンライン回路 + マイクロコントローラー シミュレーター — Arduino、ESP32、RP2040、ATtiny85 などに接続された SPICE 精度のアナログ
Velxio は、ハイブリッド デジタル + アナログ協調シミュレーションを備えた無料のオープンソース オンライン回路シミュレーターです。リアルタイム SPICE 分析 (ngspice-WASM) は、5 つの CPU アーキテクチャにわたる 19 のサポート対象ボードに接続されています。 C、Rust、または AssemblyScript でカスタム チップを構築します。 100 を超えるインタラクティブ コンポーネント、ライブ オシロスコープ、電圧計、電流計、モナコ エディター、シリアル モニター、arduino-cli コンパイラー、ライブラリ マネージャー - クラウドも遅延もアカウントも必要ありません。
ngspice-WASM (eecircuit-engine) を介したブラウザでのリアルタイム SPICE アナログ シミュレーション — 線形近似ではなく、~60 Hz での修正節点解析
ハイブリッド デジタル + アナログ協調シミュレーション — GPIO ピンは電圧源として SPICE ネットを駆動します。 analogRead() は、解決されたノード電圧をファームウェアに読み取ります。
100 以上の SPICE 精度部品: 抵抗、コンデンサ、インダクタ、BJT、MOSFET、オペアンプ (LM358/741/TL072)、レギュレータ (7805/7812/LM317)、ツェナー/ショットキー ダイオード、フォトカプラ、リレー
ライブ計測器：オシロスコープ（マルチチャンネル）、電圧計、電流計、信号発生器（正弦波/方形波/DC）
40 を超える新しいアナログおよびハイブリッドのサンプル: 分圧器、RC フィルター、全波整流器、シュミット トリガー、オペアンプ アンプ、ホイートストン ブリッジ、H ブリッジ モーター ドライバー
リアル AVR8 エミュレーション: Arduino Uno (ATmega328p)、Nano、Mega 2560、ATtiny85、Leonardo、Pro Mini (avr8js 経由 16 MHz)
Raspberry Pi Pico & Pico W (RP2040、133 MHz の ARM Cortex-M0+) (rp2040js 経由)
ESP32-C3 / XIAO ESP32-C3 / SuperMini / CH32V003 (RISC-V RV32IMC/EC) — QEMU lcgamboa 経由 (libqemu-riscv32)
ESP32/E

SP32-S3 / ESP32-CAM / Arduino Nano ESP32 (Xtensa LX6/LX7、240 MHz) QEMU lcgamboa 経由
QEMU raspi3b 経由の Raspberry Pi 3B (ARM Cortex-A53、フル Linux) — RPi.GPIO で Python スクリプトを実行します
カスタム チップ — Wokwi カスタム チップ API を使用して、C、Rust、または AssemblyScript で独自の集積回路を作成します。プロジェクト間で再利用する
48 個以上の wokwi インタラクティブ電子コンポーネント (LED、抵抗器、ボタン、センサー、TFT ディスプレイ、NeoPixel…)
完全な C++ / Python 構文ハイライトを備えた Monaco コード エディター
arduino-cli コンパイル バックエンド — 本物の .hex / .uf2 / .bin ファイルを生成します
自動ボーレート検出および送信機能を備えたシリアル モニター
Arduino ライブラリのライブラリ マネージャー
マルチファイルワークスペース (.ino、.h、.cpp、.py)
直交配線のワイヤーシステム
ILI9341 TFTディスプレイシミュレーション
I2C、SPI、USART、ADC、PWMのサポート
マルチボード キャンバス: Arduino + ESP32 + Raspberry Pi + アナログ回路を 1 つのシミュレーションにミックス
Arduino Uno (ATmega328P) — 16 MHz での完全な AVR8 エミュレーション
Arduino Nano (ATmega328P) — 16 MHz での完全な AVR8 エミュレーション
Arduino Mega 2560 (ATmega2560) — 256 KB フラッシュ、54 デジタル ピン、4 シリアル ポート
ATtiny85 — AVR、8 KB フラッシュ、DIP-8、6 つの I/O ピンすべてをエミュレート
Arduino Leonardo (ATmega32u4) — USB HID 対応 AVR
Arduino Pro Mini (ATmega328P) — 3.3 V / 5 V バージョン
Raspberry Pi Pico (RP2040) — ARM Cortex-M0+ (133 MHz)
Raspberry Pi Pico W (RP2040) — ARM Cortex-M0+ + WiFi
ESP32-C3 DevKit (RISC-V RV32IMC) — 160 MHz、QEMU libqemu-riscv32
XIAO ESP32-C3 (RISC-V RV32IMC) — コンパクト、QEMU libqemu-riscv32 を参照してください。
ESP32-C3 SuperMini (RISC-V RV32IMC) — ミニ フォーム ファクター
CH32V003 (RISC-V RV32EC) — 48 MHz、超小型 DIP-8
ESP32 DevKit V1 / C V4 (Xtensa LX6) — 240 MHz、WiFi、QEMU
ESP32-S3 (Xtensa LX7) — 240 MHz、デュアルコア、QEMU
ESP32-CAM (Xtensa LX6) — カメラモジュール、QEMU
シード

XIAO ESP32-S3 (Xtensa LX7) — コンパクト、QEMU
Arduino Nano ESP32 (Xtensa LX6) — Arduino フォームファクター、QEMU
Raspberry Pi 3B (ARM Cortex-A53) — QEMU raspi3b 経由の完全な Linux OS、Python を実行
エディタを開きます。インストールは必要ありません。
Docker によるセルフホスト: docker run -d -p 3080:80 ghcr.io/davidmonterocrespo24/velxio:master
エミュレータのセットアップ、構成、拡張方法については、Velxio の完全なドキュメントを参照してください。
はじめに — Velxio とは何ですか?また、Velxio を使用する理由は何ですか?
はじめに — ホスト型エディター、Docker、および手動セットアップ
エミュレータ アーキテクチャ — AVR8 および RP2040 CPU エミュレーションの内部
コンポーネント リファレンス — 48 以上のすべてのインタラクティブ電子コンポーネント
プロジェクト アーキテクチャ — システム設計、データ フロー、ストア
Wokwi ライブラリ — avr8js、wokwi-elements、rp2040js の統合
MCP サーバー — モデル コンテキスト プロトコルを介した AI エージェントの統合
プロジェクトのステータス - 実装されたすべての機能とトラブルシューティング
ロードマップ — 実装済み、進行中、および計画中の機能
SPICEを使用した無料のオンライン回路シミュレータ
Velxio は、ブラウザ上でリアルタイムの SPICE アナログ シミュレーションを実行できる無料のオンライン回路シミュレータです。抵抗、コンデンサ、オペアンプ、トランジスタ、ダイオード、電圧レギュレータをキャンバス上にドロップします。すべてのコンポーネントは、動作の代役ではなく、実際の ngspice デバイス カードを使用してモデル化されています。開回路シミュレータ →
ngspice-WASM エンジン — ~60 Hz の解析速度での完全な修正節点解析
非線形デバイスは実際のシリコンのように動作します: BJT カットオフ/飽和、オペアンプ レール、ダイオード順方向降下
ライブオシロスコープ、電圧計、電流計、信号発生器
100 以上の SPICE 精度の部品: 受動素子、半導体、ロジック IC、レギュレーター、フォトカプラ、リレー
常時オン — トグルなし、すべての回路は SPICE で解決されます
SPICE シミュレーター オンライン — ブラウザーで ngspice
Velxio は ngspice (オープンソース SP) を実行します。

プロフェッショナル EDA ツールを強化する ICE シミュレーター) は、eecircuit-engine 経由で WebAssembly にコンパイルされます。ブラウザタブでの完全な過渡解析、実際のデバイスモデル。インストールもアカウントも必要なく、最初のロード後はオフラインで動作します。 SPICEシミュレーターを開く→
本物の ngspice — 同じエンジン、WASM 経由でブラウザネイティブ
5 BJT (2N2222、2N3055、2N3906、BC547、BC557)
4 MOSFET (2N7000、IRF540、IRF9540、FQP27P06)
飽和レールを備えた 5 つのオペアンプ (LM358、LM741、TL072、LM324、理想的)
4 リニアレギュレータ (7805、7812、7905、LM317)
ダイオード: 1N4148、1N4007、1N5817、1N5819、ツェナー 1N4733
カスタム チップ シミュレータ — C / Rust で独自の IC を定義
Velxio は Wokwi カスタム チップ API を実装しています。チップ ロジックを C、Rust、または AssemblyScript で記述し、WebAssembly にコンパイルすると、シミュレーターが他のコンポーネントと同様にそのピンを駆動します。センサー モデル、デジタル プロトコル ブリッジ、または動作 IC の代役を構築し、プロジェクト全体で再利用します。カスタムチップシミュレーターを開く→
WASI シムを使用した Wokwi 互換カスタム チップ API
ピンの読み取り/書き込み、属性の読み取り、タイマー
チップをアカウントに保存してプロジェクト間で再利用
スターター サンプルのライブラリ (センサー、デコーダー、ブリッジ)
ATtiny85 シミュレータ — ブラウザでのリアル AVR エミュレーション
Velxio は、サイクル精度の高い AVR8 エミュレーションで ATtiny85 をシミュレートします。フル DIP-8 ピン配置、8 KB フラッシュ、6 GPIO、USI ペリフェラル、10 ビット ADC、PWM 付き Timer0/Timer1、およびウォッチドッグ。 LED、センサー、またはアナログ回路に配線し、プログラマーなしで実際の ATtiny コードを実行します。 ATtiny85シミュレータを開く→
avr8js を介したサイクル精度の高い ATtiny85 エミュレーション
6 つの I/O ピン (PB0 ～ PB5) がすべて適切なポート動作を実現
タイマ0 (8ビット)、タイマ1 (8ビット)、ウォッチドッグタイマ
4 つの入力チャンネルを備えた 10 ビット ADC
I²C / SPI ビットバンギング用の USI ペリフェラル
回路全体をテストするための SPICE アナログ部品への配線
Velxio は無料のオンライン Arduino SIM です

16 MHz でのリアル AVR8 エミュレーションを備えたエミュレータ。 48 以上のインタラクティブな電子コンポーネントを使用して Arduino コードをブラウザで直接シミュレートします。インストールもアカウントも必要ありません。 Arduinoシミュレータを開く→
Arduino Uno (ATmega328P) シミュレーション (16 MHz)
48 個以上の視覚コンポーネント: LED、センサー、ディスプレイ、サーボ
自動ボーレート検出機能付きシリアルモニター
複数ファイルのスケッチのサポート (.ino、.h、.cpp)
無料かつオープンソース (GNU AGPLv3)
Arduino エミュレータ — リアル AVR8 および RP2040 エミュレーション
Velxio は、avr8js および rp2040js 上に構築されたサイクル精度の高い Arduino エミュレータです。すべての AVR オペコードは 16 MHz で忠実にエミュレートされます。これは、実際のハードウェアと同じシリコンの動作です。 Arduinoエミュレータを開く →
135 個の AVR8 命令すべてをエミュレート
PWM、CTC、およびオーバーフロー割り込みを備えたハードウェア タイマー (Timer0/1/2)
10 ビット ADC、フル USART0 (TX/RX)、SPI、I2C
rp2040js 経由の RP2040 デュアルコア ARM Cortex-M0+
QEMU lcgamboa による RISC-V RV32IMC (ESP32-C3) エミュレーション
ATmega328P ファームウェアを、Arduino Uno および Nano 上で実行されるのとまったく同じようにシミュレートします。完全なAVR8エミュレーション: PORTB、PORTC、PORTD、Timer0/1/2、ADC、USART0、およびすべての割り込みベクトル。 ATmega328Pシミュレータを開く →
ATmega328P: 32 KB フラッシュ、2 KB SRAM、16 MHz
GPIO: PORTB (ピン 8 ～ 13)、PORTC (A0 ～ A5)、PORTD (0 ～ 7)
Timer0 (8 ビット)、Timer1 (16 ビット)、Timer2 (8 ビット) — すべての PWM モード
10 ビット ADC、6 アナログ チャネル (A0 ～ A5)
ボーレートを設定可能な USART0
Arduino Mega 2560 (ATmega2560) コードを無料でシミュレートします。 256 KB フラッシュ、54 デジタル ピン、16 アナログ入力、4 つのハードウェア シリアル ポート (シリアル、シリアル 1、シリアル 2、シリアル 3)、および 6 つのハードウェア タイマー。 Mega 2560 シミュレーターを開く →
ATmega2560: 256 KB フラッシュ、8 KB SRAM、16 MHz
54 デジタル I/O ピン (PORTA ～ PORTL)
16 アナログ入力チャンネル (A0 ～ A15)、10 ビット ADC
4 つのハードウェア USART チャネル: Serial0 ～ 3
6 つのハードウェア タイマー: Timer0 ～ 5 (3 つの 16 ビットを含む)

## Original Extract

Velxio 3.0 adds programmable retro CPUs (Z80, 8080, 4004, 4040, 8086), MicroSD card emulation, ePaper displays, true multi-board UART/I2C/SPI interconnect, full undo/redo and 100+ new examples. Free, open-source, browser-based — Arduino, ESP32, RP2040, STM32, Raspberry Pi and more.

Velxio
Offline Arduino, RP2040 & ESP32 simulator
Starting local backend…
Velxio — Free Online Circuit & Arduino Simulator | SPICE · ESP32 · RP2040 · ATtiny85 · Custom Chips
Free online circuit + microcontroller simulator — SPICE-accurate analog wired to Arduino, ESP32, RP2040, ATtiny85 & more
Velxio is a free, open-source online circuit simulator with hybrid digital + analog co-simulation. Real-time SPICE analysis (ngspice-WASM) wired to 19 supported boards across 5 CPU architectures. Build custom chips in C, Rust, or AssemblyScript. 100+ interactive components, live oscilloscope, voltmeter, ammeter, Monaco Editor, Serial Monitor, arduino-cli compiler, and Library Manager — no cloud, no latency, no account required.
Real-time SPICE analog simulation in your browser via ngspice-WASM (eecircuit-engine) — Modified Nodal Analysis at ~60 Hz, not a linear approximation
Hybrid digital + analog co-simulation — GPIO pins drive SPICE nets as voltage sources; analogRead() reads solved node voltages back into firmware
100+ SPICE-accurate parts : resistors, capacitors, inductors, BJTs, MOSFETs, op-amps (LM358/741/TL072), regulators (7805/7812/LM317), Zener/Schottky diodes, optocouplers, relays
Live instruments : oscilloscope (multi-channel), voltmeter, ammeter, signal generator (sine/square/DC)
40+ new analog & hybrid examples : voltage divider, RC filter, full-wave rectifier, Schmitt trigger, op-amp amplifier, Wheatstone bridge, H-bridge motor driver
Real AVR8 emulation: Arduino Uno (ATmega328p), Nano, Mega 2560, ATtiny85, Leonardo, Pro Mini at 16 MHz via avr8js
Raspberry Pi Pico & Pico W (RP2040, ARM Cortex-M0+ at 133 MHz) via rp2040js
ESP32-C3 / XIAO ESP32-C3 / SuperMini / CH32V003 (RISC-V RV32IMC/EC) — via QEMU lcgamboa (libqemu-riscv32)
ESP32 / ESP32-S3 / ESP32-CAM / Arduino Nano ESP32 (Xtensa LX6/LX7 at 240 MHz) via QEMU lcgamboa
Raspberry Pi 3B (ARM Cortex-A53, full Linux) via QEMU raspi3b — runs Python scripts with RPi.GPIO
Custom Chips — author your own integrated circuits in C, Rust, or AssemblyScript using the Wokwi Custom Chips API; reuse them across projects
48+ wokwi interactive electronic components (LEDs, resistors, buttons, sensors, TFT displays, NeoPixel…)
Monaco Code Editor with full C++ / Python syntax highlighting
arduino-cli compilation backend — produces real .hex / .uf2 / .bin files
Serial Monitor with auto baud-rate detection and send
Library Manager for Arduino libraries
Multi-file workspace (.ino, .h, .cpp, .py)
Wire system with orthogonal routing
ILI9341 TFT display simulation
I2C, SPI, USART, ADC, PWM support
Multi-board canvas: mix Arduino + ESP32 + Raspberry Pi + analog circuits in one simulation
Arduino Uno (ATmega328P) — full AVR8 emulation at 16 MHz
Arduino Nano (ATmega328P) — full AVR8 emulation at 16 MHz
Arduino Mega 2560 (ATmega2560) — 256 KB flash, 54 digital pins, 4 serial ports
ATtiny85 — AVR, 8 KB flash, DIP-8, all 6 I/O pins emulated
Arduino Leonardo (ATmega32u4) — USB HID capable AVR
Arduino Pro Mini (ATmega328P) — 3.3 V / 5 V variants
Raspberry Pi Pico (RP2040) — ARM Cortex-M0+ at 133 MHz
Raspberry Pi Pico W (RP2040) — ARM Cortex-M0+ + WiFi
ESP32-C3 DevKit (RISC-V RV32IMC) — 160 MHz, QEMU libqemu-riscv32
Seeed XIAO ESP32-C3 (RISC-V RV32IMC) — compact, QEMU libqemu-riscv32
ESP32-C3 SuperMini (RISC-V RV32IMC) — mini form factor
CH32V003 (RISC-V RV32EC) — 48 MHz, ultra-compact DIP-8
ESP32 DevKit V1 / C V4 (Xtensa LX6) — 240 MHz, WiFi, QEMU
ESP32-S3 (Xtensa LX7) — 240 MHz, dual-core, QEMU
ESP32-CAM (Xtensa LX6) — camera module, QEMU
Seeed XIAO ESP32-S3 (Xtensa LX7) — compact, QEMU
Arduino Nano ESP32 (Xtensa LX6) — Arduino form factor, QEMU
Raspberry Pi 3B (ARM Cortex-A53) — full Linux OS via QEMU raspi3b, runs Python
Open the Editor — no installation needed.
Self-host with Docker: docker run -d -p 3080:80 ghcr.io/davidmonterocrespo24/velxio:master
Browse the full Velxio documentation to learn how to set up, configure, and extend the emulator:
Introduction — What is Velxio and why use it?
Getting Started — Hosted editor, Docker, and manual setup
Emulator Architecture — AVR8 and RP2040 CPU emulation internals
Components Reference — All 48+ interactive electronic components
Project Architecture — System design, data flows, and stores
Wokwi Libraries — avr8js, wokwi-elements, rp2040js integration
MCP Server — AI agent integration via Model Context Protocol
Project Status — All implemented features and troubleshooting
Roadmap — Implemented, in-progress, and planned features
Free Online Circuit Simulator with SPICE
Velxio is a free online circuit simulator with real-time SPICE analog simulation in your browser. Drop resistors, capacitors, op-amps, transistors, diodes, and voltage regulators on the canvas — every component is modelled with real ngspice device cards, not behavioural stand-ins. Open Circuit Simulator →
ngspice-WASM engine — full Modified Nodal Analysis at ~60 Hz solve rate
Non-linear devices behave like real silicon: BJT cutoff/saturation, op-amp rails, diode forward drop
Live oscilloscope, voltmeter, ammeter, signal generator
100+ SPICE-accurate parts: passives, semiconductors, logic ICs, regulators, optocouplers, relays
Always-on — no toggle, every circuit is SPICE-solved
SPICE Simulator Online — ngspice in Your Browser
Velxio runs ngspice (the open-source SPICE simulator powering professional EDA tools) compiled to WebAssembly via eecircuit-engine. Full transient analysis, real device models, in your browser tab. No install, no account, works offline after the first load. Open SPICE Simulator →
Real ngspice — same engine, browser-native via WASM
5 BJTs (2N2222, 2N3055, 2N3906, BC547, BC557)
4 MOSFETs (2N7000, IRF540, IRF9540, FQP27P06)
5 op-amps (LM358, LM741, TL072, LM324, ideal) with saturation rails
4 linear regulators (7805, 7812, 7905, LM317)
Diodes: 1N4148, 1N4007, 1N5817, 1N5819, Zener 1N4733
Custom Chip Simulator — Define Your Own ICs in C / Rust
Velxio implements the Wokwi Custom Chips API: write your chip logic in C, Rust, or AssemblyScript, compile to WebAssembly, and the simulator drives its pins like any other component. Build a sensor model, a digital protocol bridge, or a behavioural IC stand-in — and reuse it across projects. Open Custom Chip Simulator →
Wokwi-compatible Custom Chips API with WASI shim
Pin reads/writes, attribute reads, timers
Save chips to your account and reuse across projects
Library of starter examples (sensors, decoders, bridges)
ATtiny85 Simulator — Real AVR Emulation in Your Browser
Velxio simulates the ATtiny85 with cycle-accurate AVR8 emulation. Full DIP-8 pinout, 8 KB flash, 6 GPIOs, USI peripheral, 10-bit ADC, Timer0/Timer1 with PWM, and watchdog. Wire it to LEDs, sensors, or analog circuits and run real ATtiny code without a programmer. Open ATtiny85 Simulator →
Cycle-accurate ATtiny85 emulation via avr8js
All 6 I/O pins (PB0–PB5) with proper port behaviour
Timer0 (8-bit), Timer1 (8-bit), watchdog timer
10-bit ADC with 4 input channels
USI peripheral for I²C / SPI bit-banging
Wire to SPICE analog parts for full circuit testing
Velxio is a free online Arduino simulator with real AVR8 emulation at 16 MHz. Simulate Arduino code with 48+ interactive electronic components directly in your browser — no install, no account required. Open Arduino Simulator →
Arduino Uno (ATmega328P) simulation at 16 MHz
48+ visual components: LEDs, sensors, displays, servos
Serial Monitor with auto baud-rate detection
Multi-file sketch support (.ino, .h, .cpp)
Free and open-source (GNU AGPLv3)
Arduino Emulator — Real AVR8 & RP2040 Emulation
Velxio is a cycle-accurate Arduino emulator built on avr8js and rp2040js. Every AVR opcode is faithfully emulated at 16 MHz — the same silicon behavior as real hardware. Open Arduino Emulator →
All 135 AVR8 instructions emulated
Hardware timers (Timer0/1/2) with PWM, CTC, and overflow interrupts
10-bit ADC, full USART0 (TX/RX), SPI, I2C
RP2040 dual-core ARM Cortex-M0+ via rp2040js
RISC-V RV32IMC (ESP32-C3) emulation via QEMU lcgamboa
Simulate ATmega328P firmware exactly as it runs on Arduino Uno and Nano. Full AVR8 emulation: PORTB, PORTC, PORTD, Timer0/1/2, ADC, USART0, and all interrupt vectors. Open ATmega328P Simulator →
ATmega328P: 32 KB flash, 2 KB SRAM, 16 MHz
GPIO: PORTB (pins 8–13), PORTC (A0–A5), PORTD (0–7)
Timer0 (8-bit), Timer1 (16-bit), Timer2 (8-bit) — all PWM modes
10-bit ADC, 6 analog channels (A0–A5)
USART0 with configurable baud rate
Simulate Arduino Mega 2560 (ATmega2560) code for free. 256 KB flash, 54 digital pins, 16 analog inputs, 4 hardware serial ports (Serial, Serial1, Serial2, Serial3), and 6 hardware timers. Open Mega 2560 Simulator →
ATmega2560: 256 KB flash, 8 KB SRAM, 16 MHz
54 digital I/O pins (PORTA through PORTL)
16 analog input channels (A0–A15), 10-bit ADC
4 hardware USART channels: Serial0–3
6 hardware timers: Timer0–5 (including three 16-bit)
