---
source: "https://www.cnx-software.com/2026/07/30/nightrun-uefi-application-boots-a-local-llm-on-raspberry-pi-5-and-x86-pcs-without-an-os/"
hn_url: "https://news.ycombinator.com/item?id=49112667"
title: "Run LLM Without OS"
article_title: "NightRun UEFI application boots a local LLM on Raspberry Pi 5 and x86 PCs without an OS - CNX Software"
author: "boveyking"
captured_at: "2026-07-30T17:16:04Z"
capture_tool: "hn-digest"
hn_id: 49112667
score: 1
comments: 2
posted_at: "2026-07-30T16:59:28Z"
tags:
  - hacker-news
  - translated
---

# Run LLM Without OS

- HN: [49112667](https://news.ycombinator.com/item?id=49112667)
- Source: [www.cnx-software.com](https://www.cnx-software.com/2026/07/30/nightrun-uefi-application-boots-a-local-llm-on-raspberry-pi-5-and-x86-pcs-without-an-os/)
- Score: 1
- Comments: 2
- Posted: 2026-07-30T16:59:28Z

## Translation

タイトル: OS なしで LLM を実行する
記事のタイトル: NightRun UEFI アプリケーションは、OS なしで Raspberry Pi 5 および x86 PC 上でローカル LLM を起動します - CNX ソフトウェア
説明: LLM を SBC またはミニ PC 上でローカルに実行することは、データをプライベートに保つ良い方法ですが、通常は従来のオペレーティング システムが必要です。

記事本文:
NightRun UEFI アプリケーションは、OS なしで Raspberry Pi 5 および x86 PC 上でローカル LLM を起動します - CNX ソフトウェア
コンテンツにスキップ
CNX ソフトウェア – 組み込みシステム ニュース
組み込みシステム、IoT、オープンソース ハードウェア、SBC、マイクロコントローラー、プロセッサーなどに関するレビュー、チュートリアル、最新ニュース
広告およびコンサルティングサービス
NightRun UEFI アプリケーションは、OS なしで Raspberry Pi 5 および x86 PC 上でローカル LLM を起動します
LLM を SBC またはミニ PC 上でローカルに実行することは、データの機密性を保つための良い方法ですが、通常、システム メモリの大部分を使用する従来のオペレーティング システムが必要です。 NightRun は実験的なオープンソース プロジェクトで、従来の OS をロードせずに、USB ドライブまたは microSD カードからローカル LLM でマシンを直接起動することで、この問題に対処します。
NightRun はオペレーティング システムを削除することで、より多くの RAM とメモリ帯域幅を AI 推論に利用できるようにします。ランタイムは Rust で書かれており、ブート中に CRC-32 チェックサムを検証しながら量子化モデル (1.3 GB ～ 2.4 GB) を RAM に直接ロードします。モデルがロードされた後、ストレージは「密閉」されます。これは、後でディスクから読み取ろうとするとハードフォールトがトリガーされることを意味します。
ハードウェア サポート – USB または microSD カード経由で 64 ビット x86_64 UEFI PC (セキュア ブート無効) および Raspberry Pi 5 をサポート
サポートされているモデル – .nrm 形式に変換された Llama 3.2 1B (4GB RAM)、Llama 3.2 3B および Granite 4.1 3B (6GB RAM)、および Qwen3 4B Instruct 2507 (8GB RAM) モデルをサポート
最適化された推論 – x86 では AVX2/FMA/F16C カーネルを使用し、Raspberry Pi では NEON カーネルを使用します。
ゼロコピー設計 – 量子化された Q8_0、Q4_K、および Q6_K モデルを逆量子化せずに RAM から直接実行します
内蔵 UI – ライブ統計、プロンプト編集、会話履歴を備えたフレームバッファ チャット インターフェイスが含まれています
ナイトラン

内蔵フレームバッファ チャット インターフェイスで Granite 4.1 3B モデルを使用して応答を生成します。
NightRun のフレームバッファ チャット インターフェイスでの Granite 4.1 3B モデルからの応答が完了し、ステータス バーにライブ推論統計が表示されます。
NightRun はメモリを節約するために設計された必要最低限​​の機能を備えた Linux ディストリビューションであると思われがちですが、そうではありません。これは従来のオペレーティング システムなしで実行される UEFI アプリケーションであるため、その下に Linux カーネル、スケジューラー、ネットワーク スタックはありません。代わりに、UEFI ブート サービスに依存して、キーボード入力の処理、ディスプレイ フレームバッファの駆動、ストレージへのアクセス、さらには Raspberry Pi 5 冷却ファンの制御を行います。
パフォーマンスの点では、デコード速度は主にメモリ帯域幅によって制限されると開発者は述べています。 8 コア x86 QEMU 仮想マシンでは、Llama 3.2 1B (Q8_0) はデコード中に 1 秒あたり約 20 トークンに達します。実際の 8GB Raspberry Pi 5 では、Granite 4.1 3B (Q4_K_M) は、初期の pre-sdot NEON カーネルを使用して 1 秒あたり約 3.0 トークンを生成します。
以前、軽量 Linux ディストリビューション (Alpine や DietPi など) 上にデプロイされた llama.cpp などの C/C++ 推論エンジンを使用してメモリを解放する Binh Pham の LMStick について書きました。それと比較すると、NightRun は、消費者向けハードウェア向けに特別に設計された、パリティテスト済みのトークン化を備えた完全な実稼働対応 UEFI ランタイム アプライアンスです。
NightRun プロジェクトのコードの大部分は、Fable 5 モデルの Claude Code を使用して記述され、GitHub で MIT ライセンスの下でリリースされました。インストールは、検証済みモデルをダウンロードし、ブート可能イメージをフラッシュする Linux シェル スクリプトによって処理されます。開発者は、インストーラーにさまざまな安全性チェックを実装して、明示的に FLASH /dev/sdX と入力し、システム ドライブの誤った上書きを防止しました。

イブス。技術的な実装の詳細については、nightrun.io を参照してください。
Debashis Das は、業界で 5 年以上の経験を持つテクニカル コンテンツ ライター兼組み込みエンジニアです。組み込み C、PCB 設計、SEO 最適化の専門知識を活かし、難しい技術的なトピックと明確なコミュニケーションを効果的に融合させています。
これを共有してください:
CNX ソフトウェアをサポート!暗号通貨経由で寄付したり、Patreon でパトロンになったり、Amazon や Aliexpress で商品を購入したりできます。また、記事内でアフィリエイト リンクを使用し、リンクをクリックして商品を購入するとコミッションが発生します。
2021 年を振り返る – トップ 10 の投稿と統計
CNX ソフトウェアの 2024 年の総括、ウェブサイトの統計、2025 年に予想されること
LLMStick – Raspberry Pi Zero W と最適化された llama.cpp に基づく AI および LLM USB デバイス
2025 年の振り返り、CNX ソフトウェアの統計、そして 2026 年の展望
GyroidOS 仮想化ソリューションは、組み込みデバイスのセキュリティを確保し、サイバーセキュリティ認証を容易にすることを目的としています
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
コメント フォームでは、Web サイトに投稿されたコメントを追跡できるように、あなたの名前、電子メール、内容が収集されます。コメントを投稿するには、当社の Web サイト利用規約とプライバシー ポリシーを読んで同意してください。
1 コメント
最古の
最新の
さらにコメントを読み込む
ポストナビゲーション
Twitter LinkedIn Telegram YouTube RSS
注目の投稿 - 過去 7 日間
Ubuntu Touch 24.04-2.0 は Morph ブラウザを更新し、ノッチと印刷サポートなどを追加します。コメント1件
Open RC Spotter は、オープンソースの ESP32 テレメトリです。

RCカーや玩具用のdデータロガー 4kビュー | 3件のコメント
Beelink ME Pro 2 ベイ ハイブリッド NAS は、Wildcat Lake 処理を受け、10GbE と Thunderbolt 4 を獲得し、いくつかの M.2 スロットを失います。 9件のコメント
Creality Pika レビュー – 赤外線および青色レーザー モードを備えた手頃な価格のポータブル 3D スキャナー6件のコメント
Orange Pi 5B SBC は、8 年間失われていた Bluetooth/Wi-Fi 修正を含む Ubuntu デスクトップおよびオーディオ プロダクション OS イメージを取得します。 6件のコメント
Orange Pi 5B SBC 上の citral は、Bluetooth/Wi-Fi 修正が 8 年間欠落していた Ubuntu デスクトップおよびオーディオ プロダクション OS イメージを取得します
NightRun UEFI アプリケーションの Shane は、OS のない Raspberry Pi 5 および x86 PC 上でローカル LLM を起動します
Jean-Luc Aufranc (CNXSoft) による ESP-KVM – ESP32-P4 RISC-V MCU に基づくオープンソース IP KVM ソリューション
Waveshare RP2350-POE-ETH 上の RoganDawes – PoE を備えた Raspberry Pi RP2350 MCU ボード
Creality Pika の Bilal JK レビュー – 赤外線および青色レーザー モードを備えた手頃な価格のポータブル 3D スキャナー
Creality Pika レビュー – 赤外線および青色レーザー モードを備えた手頃な価格のポータブル 3D スキャナ Creality は、レビューのために Pika ポータブル 3D スキャナのサンプルを送ってくれました。ある意味、それはいくつかの点を共有しています […]
Raspberry Pi Touch Display 2 – 10″ Portrait レビュー、3D プリント VESA マウント Raspberry Pi は、「Raspberry Pi Touch Display 2 – 10″ Portrait」を発表しました。
Beelink EQi 304 レビュー – パート 3: Wildcat Lake ミニ PC 上の Ubuntu 26.04 Beelink EQi Wildcat Lake ミニ PC の開梱と分解でハードウェアをチェックした後、[…]

## Original Extract

Running LLMs locally on an SBC or mini PC is a good way to keep your data private, but they typically require a conventional operating system, which uses

NightRun UEFI application boots a local LLM on Raspberry Pi 5 and x86 PCs without an OS - CNX Software
Skip to content
CNX Software – Embedded Systems News
Reviews, tutorials and the latest news about embedded systems, IoT, open-source hardware, SBC's, microcontrollers, processors, and more
Advertisement & Consulting Services
NightRun UEFI application boots a local LLM on Raspberry Pi 5 and x86 PCs without an OS
Running LLMs locally on an SBC or mini PC is a good way to keep your data private, but they typically require a conventional operating system, which uses a good portion of system memory. NightRun is an experimental open-source project that addresses this issue by booting a machine directly into a local LLM from a USB drive or microSD card, without loading a conventional OS.
By removing the operating system, NightRun makes more RAM and memory bandwidth available for AI inference. The runtime is written in Rust, and during boot it loads a quantized model (1.3 GB to 2.4 GB) directly into RAM while verifying its CRC-32 checksums. After the model is loaded, storage is “sealed,” meaning any later attempt to read from the disk will trigger a hard fault.
Hardware support – Supports 64-bit x86_64 UEFI PCs (Secure Boot disabled) and Raspberry Pi 5 via USB or microSD card
Supported models – Supports Llama 3.2 1B (4GB RAM), Llama 3.2 3B and Granite 4.1 3B (6GB RAM), and Qwen3 4B Instruct 2507 (8GB RAM) models converted to the .nrm format
Optimized inference – Uses AVX2/FMA/F16C kernels on x86 and NEON kernels on Raspberry Pi
Zero-Copy design – Runs quantized Q8_0, Q4_K, and Q6_K models directly from RAM without dequantization
Built-in UI – Includes a framebuffer chat interface with live stats, prompt editing, and conversation history
NightRun generating a response with the Granite 4.1 3B model in its built-in framebuffer chat interface.
Completed response from the Granite 4.1 3B model in NightRun’s framebuffer chat interface, with live inference statistics displayed in the status bar.
It might be easy to think that NightRun is a stripped-down Linux distribution designed to save memory, but that is not the case. It is a UEFI application that runs without a conventional operating system, so there is no Linux kernel, scheduler, or network stack underneath. Instead, it relies on UEFI Boot Services to handle keyboard input, drive the display framebuffer, access storage, and even control the Raspberry Pi 5 cooling fan.
In terms of performance, the developers say decoding speed is mainly limited by memory bandwidth. In an 8-core x86 QEMU virtual machine, Llama 3.2 1B (Q8_0) reaches around 20 tokens per second during decoding. On a real 8GB Raspberry Pi 5, Granite 4.1 3B (Q4_K_M) generates about 3.0 tokens per second using the early pre-sdot NEON kernels.
We have previously written about Binh Pham’s LLMStick that uses C/C++ inference engines like llama.cpp deployed on top of lightweight Linux distributions (such as Alpine or DietPi) to free up memory. Compared to that, NightRun is a complete, production-ready UEFI runtime appliance with parity-tested tokenization specifically designed for consumer hardware.
The majority of the code for the NightRun project was written using Claude Code with the Fable 5 model and released under the MIT license on GitHub . Installation is handled by a Linux shell script that downloads a verified model and flashes the bootable image. The developers have implemented various safety checks in the installer to explicitly type FLASH /dev/sdX to prevent accidental overwriting of system drives. More details on the technical implementation can be found at nightrun.io.
Debashis Das is a technical content writer and embedded engineer with over five years of experience in the industry. With expertise in Embedded C, PCB Design, and SEO optimization, he effectively blends difficult technical topics with clear communication
Share this:
Support CNX Software! Donate via cryptocurrencies , become a Patron on Patreon, or purchase goods on Amazon or Aliexpress . We also use affiliate links in articles to earn commissions if you make a purchase after clicking on those links.
Year 2021 in review – Top 10 posts and statistics
CNX Software’s 2024 Year in review, website statistics, and what to expect in 2025
LLMStick – An AI and LLM USB device based on Raspberry Pi Zero W and optimized llama.cpp
Year 2025 in Review, CNX Software stats, and looking ahead to 2026
GyroidOS virtualization solution aims to secure embedded devices, ease cybersecurity certification
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
1 Comment
oldest
newest
Load More Comments
Post navigation
Twitter LinkedIn Telegram YouTube RSS
Trending Posts - Last 7 Days
Ubuntu Touch 24.04-2.0 updates Morph browser, adds notch and printing support, and more 5k views | 1 comment
Open RC Spotter is an open-source ESP32 telemetry and data logger for RC cars and toys 4k views | 3 comments
Beelink ME Pro 2-bay hybrid NAS gets Wildcat Lake treatment, gains 10GbE and Thunderbolt 4, loses a few M.2 slots 3.5k views | 9 comments
Creality Pika review – An affordable, portable 3D scanner with infrared and blue laser modes 3.2k views | 6 comments
Orange Pi 5B SBC gets Ubuntu Desktop and Audio Production OS images with Bluetooth/Wi-Fi fix missed for 8 years 2.9k views | 6 comments
citral on Orange Pi 5B SBC gets Ubuntu Desktop and Audio Production OS images with Bluetooth/Wi-Fi fix missed for 8 years
Shane on NightRun UEFI application boots a local LLM on Raspberry Pi 5 and x86 PCs without an OS
Jean-Luc Aufranc (CNXSoft) on ESP-KVM – An open-source IP KVM solution based on ESP32-P4 RISC-V MCU
RoganDawes on Waveshare RP2350-POE-ETH – A Raspberry Pi RP2350 MCU board with PoE
Bilal JK on Creality Pika review – An affordable, portable 3D scanner with infrared and blue laser modes
Creality Pika review – An affordable, portable 3D scanner with infrared and blue laser modes Creality has sent us a sample of the Pika portable 3D scanner for review. In some ways, it shares some of the […]
Raspberry Pi Touch Display 2 – 10″ Portrait review, 3D printed VESA mount Raspberry Pi has just announced the “Raspberry Pi Touch Display 2 – 10″ Portrait”, and […]
Beelink EQi 304 review – Part 3: Ubuntu 26.04 on a Wildcat Lake mini PC After checking out the hardware with an unboxing and a teardown of the Beelink EQi Wildcat Lake mini PC, we […]
