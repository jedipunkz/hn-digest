---
source: "https://github.com/valivalivali/x-os"
hn_url: "https://news.ycombinator.com/item?id=49057264"
title: "X-OS – Building an Operating System for the AI Era"
article_title: "GitHub - valivalivali/x-os · GitHub"
author: "vali1"
captured_at: "2026-07-26T12:59:02Z"
capture_tool: "hn-digest"
hn_id: 49057264
score: 1
comments: 3
posted_at: "2026-07-26T12:03:09Z"
tags:
  - hacker-news
  - translated
---

# X-OS – Building an Operating System for the AI Era

- HN: [49057264](https://news.ycombinator.com/item?id=49057264)
- Source: [github.com](https://github.com/valivalivali/x-os)
- Score: 1
- Comments: 3
- Posted: 2026-07-26T12:03:09Z

## Translation

タイトル: X-OS – AI 時代のオペレーティング システムの構築
記事タイトル: GitHub - valivalivali/x-os · GitHub
説明: GitHub でアカウントを作成して、valivalivali/x-os の開発に貢献します。

記事本文:
GitHub - バリバリバリ/x-os · GitHub
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
バリバリバリ
/
X-OS
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード O

ペンのその他のアクション メニュー フォルダーとファイル
85 コミット 85 コミット boot boot bsd bsd docs docs カーネル カーネル スクリーンショット スクリーンショット スクリプト スクリプト サードパーティ サードパーティ ツール ツール ユーザースペース ユーザースペース .gitignore .gitmodules .gitmodules AGENTS.md AGENTS.md LICENSE LICENSE Makefile Makefile README.md README.md disc.img disk.img lv_conf.h lv_conf.h すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 時代に向けてゼロから構築された先進的なオペレーティング システム。
X OS は、インテリジェント エージェントと人間のユーザーが連携する世界向けに設計された白紙のマイクロカーネル オペレーティング システムです。これは単純な質問です。エージェント プログラミング、プログラムによる対話、人間と AI のコラボレーションを目的として、今日設計された OS はどのようなものになるでしょうか?
GPU アクセラレーション レンダリングを備えた QEMU で実行される X OS — メニュー バー、アプリ アイコンを備えたドック、コンポジター サーフェス、および 2560 × 1600 のデスクトップ。
ソフトウェアは変化しています。 AI エージェントは API を呼び出すだけではなく、コードの読み取り、書き込み、コマンドの実行、デバッグ、反復を行います。現在のオペレーティング システムが提供するように設計されていないレベルで、監視可能、スクリプト可能、変更可能な環境が必要です。
X OS は異なる前提から始まります。このシステムは、最上位に階層化されたツールを介するのではなく、第一級のアーキテクチャ原則として、すべてのサービスが検査可能、すべてのコンポーネントが交換可能、すべての対話がプログラム可能であるように構築されています。マイクロカーネルは、任意のプロセス (人間またはエージェント) がシステムをリアルタイムでクエリ、制御、拡張するために使用できる IPC ポートを公開します。
AI ネイティブのアーキテクチャ — OS はエージェントがファーストクラスのユーザーであることを前提としています。 IPC、サービス ディスカバリー、システム イントロスペクションは、初日からプログラムで使用できるように設計されています。
設計によるマイクロカーネル — カーネルは、スケジューリング、メモリ、IPC、およびハードウェアを処理します。すべて

その他 (ディスプレイ、ファイルシステム、ネットワーキング、シェル) は、システムの実行中に検査、変更、または置換できるユーザー空間サービスです。
すべてが監視可能です。サービスは、検出のために IPC ポートを介してネームサーバーと通信します。どのプロセスでも、システム状態のクエリ、イベントのサブスクライブ、または介入を行うことができます。これはデバッグ機能ではありません。それは建築です。
白紙の状態 — 維持する必要のある数十年前の ABI も、レガシーのゴミもありません。システムは、現在および今後の将来について正しい決定を自由に行うことができます。
将来を見据えた — エージェントのワークフロー、プログラムによる対話、人間と AI のコラボレーションなど、将来に向けて設計されています。
核となるアイデアは固定されています。それは、エージェントのワークフローとインテリジェントなユーザー向けに設計された、AI 時代に向けて構築された先進的なオペレーティング システムです。それは変わりません。
それ以外はすべてオープンです。アーキテクチャ、UI、テクノロジー、優先順位はすべて、プロジェクトが成長し、AI ネイティブ コンピューティングの状況が形づくられるにつれて進化します。今日行われた設計上の決定は、明日には修正される可能性があります。コンポーネントは書き換え、交換、または削除される可能性があります。それはリスクではありません。それがポイントです。
デスクトップ環境は、メニュー バー、ドック、ウィンドウ、カーソルが親しみやすいものであることを目指しています。私たちが既知の慣例に頼るのは、それが機能し、ユーザーがすでにそれを理解しているからです。しかし、これは約束ではありません。根本的に優れたものを発見した場合は、それを探索します。目標は、今日は直感的で、明日は再発明できるインターフェイスを構築することです。
X OS は、実際のハードウェア (Limine 経由) および QEMU で起動し、デスクトップ環境が動作します。
マイクロカーネル — スケジューリング、メモリ管理、IPC ポート、タイマー、GPU/virtio ドライバー、NVMe ストレージ、ネットワーキングをカバーする最大 90 のシステムコール
ユーザースペース コンポーザー — サーフェス合成、ダートを備えたハードウェア アクセラレーション ディスプレイ サーバー

y 四角形、ウィンドウ装飾、カーソル管理、virgl (OpenGL ES) による GPU レンダリング
メニュー バー — アプリ メニュー、ドロップダウン、およびフォーカス トラッキングを備えたトップ パネル
ドック — アプリのアイコン、ホバー効果、アプリの起動/終了を含む下部パネル
コンテキスト メニュー — egui (Rust、no_std) を利用した右クリック ポップアップ サービス
シェル — ユーザー空間プロセスとして実行される埋め込み zsh ポート
ファイル マネージャー (Xplorer) — ドラッグ可能なウィンドウ アプリ
ネットワーキング — マイクロカーネルに移植された XNU BSD ネットワーク スタック: TCP/UDP ソケット、virtio-net ドライバー、POSIX ソケット API
ファイルシステム — NVMe ブロックデバイスをサポートするカスタム XFS ファイルシステム
GPU アクセラレーション — OpenGL ES 3D 合成用の virglrenderer を備えた virtio-gpu。アルファ ブレンディングを使用したテクスチャ クワッド レンダリング用のカスタム GPU コマンド ストリーム
egui の統合 — egui 0.35 を no_std / x86_64-unknown-none に移植し、カスタム CPU ラスタライザー (スキャンライン グリフ レンダラー) と即時モード UI レンダリング用のソフトウェア バックエンドを備えています。
コンポーネント
リング
責任
カーネル
0
スケジューリング、メモリ割り当て/マップ、IPC ポート、タイマー、割り込み、NVMe、virtio GPU、virtio-net、PCI
初期化 (PID 1)
3
最初のユーザー空間プロセス。サービスを生成し、ネームサーバーのポートを登録します
作曲家
3
ディスプレイサーバー — サーフェス、ダーティ四角形、ウィンドウ装飾、カーソル、GPU 合成
メニューバー
3
トップパネル - X ロゴ、アプリメニュー、ドロップダウン、フォーカストラッキング
ドック
3
下部パネル — アプリアイコン、ホバー効果、生成/非表示/表示/閉じる
コンテキストメニュー
3
右クリックポップアップ — eguiベース、Rust no_std staticlib
シェル
3
zsh ポート — ユーザー空間プロセスとして実行
エクスプローラー
3
ファイルマネージャー - ドラッグ可能なウィンドウ
BSD レイヤー
0→3
マイクロカーネルに適合した XNU BSD ネットワーキング スタック (ソケット、TCP/IP、virtio-net)
仕組み
カーネルは最小限のマイクロカーネルです。CPU スケジューリング、物理/仮想メモリ、IPC メッセージ ポート、およびハードウェア ドライバー (NV) を処理します。

私、virtio-gpu、virtio-net、PS/2 入力、PCI、タイマー)。ファイルシステム コード、表示ロジック、ネットワークは含まれていません。これらはユーザー空間サービスです。
すべてのユーザー空間バイナリ (init、composer、dock、menubar、menu、shell) は ELF バイナリとしてコンパイルされ、バイト配列としてカーネル イメージに直接埋め込まれます。カーネルはブート時にそれらを生成します。ディスクのロードやユーザー空間のブートローダー チェーンはありません。
サービスは IPC ポート経由で通信します。 Composer は既知のポートを登録します。アプリは、サーフェス作成リクエスト、ダーティ四角形、およびマウス イベントを IPC 経由で送信します。ネームサーバー ( sys_ns_register / sys_ns_lookup ) はサービス検出を提供します。
GPU 合成には virgl (virtio-gpu 3D 経由の OpenGL ES) を使用します。各アプリのサーフェスは GPU テクスチャを取得します。コンポーザーは、スキャンアウト ソースでもあるフレームバッファを使用したレンダー ターゲット上に、アルファ ブレンディングを使用してテクスチャ付きクワッドを描画する 3D コマンド ストリームを送信します。
Apple Siliconを搭載したmacOSでテスト済み。必要なものは次のとおりです。
Xcode コマンド ライン ツール (clang、make、git を提供)
xorriso (ブート可能な ISO の構築用)
QEMU — Homebrew 経由でインストールするか、ソースからビルドします。 Makefile はパスを自動検出します。
オプション A — 自作 (最速):
醸造インストールqemu
オプション B — virgl/ANGLE GPU レンダリングを使用してソースから QEMU v11 をビルドします (Apple Silicon での GPU アクセラレーションに推奨):
# virgl/ANGLE 依存関係用の startergo タップをインストールします
brew Tap startergo/homebrew-qemu-virgl
brew trust startergo/qemu-virgl
brew install libangle libepoxy-angle virglrenderer スパイス-プロトコル スパイス-サーバー \
meson ninja glib pixman pkg-config dtc vde libssh
# Cocoa + virgl + OpenGL ES を使用して QEMU v11.0.2 をビルドする
bash tools/build-qemu-virgl.sh
これにより、virtio-gpu-gl-pci サポート、Cocoa ディスプレイ、および ANGLE/virglrenderer を介した OpenGL ES レンダリングを備えた QEMU v11.0.2 がビルドされます。スクリプトはパッチ適用、ビルド、

/opt/qemu-head へのインストール、dylib パスの修正、HVF 資格によるコード署名。
Makefile は最初に /opt/qemu-head をチェックし、次に Homebrew qemu-virgl プレフィックス、次に Homebrew qemu プレフィックスをチェックしてから、PATH 内の qemu-system-x86_64 にフォールバックします。
ワンタイムセットアップ (Limine ブートローダーをダウンロード):
セットアップを行う
ブータブル ISO をビルドします。
作る
これにより、プロジェクト ルートに x-os.iso が生成されます。
run-uefi を作成する
QEMU は次のように起動されます。
virtio-gpu-pci (2560x1600、Cocoa ディスプレイ)
NVMe ディスク (disk.img、見つからない場合は自動的に作成されます)
シリアル出力を標準入出力に転送
×/
§── boot/ # ブートローダーの設定とハンドオフ構造を制限する
§── kernel/ # マイクロカーネルソース
│ §── Arch/x86_64/ # GDT、IDT、syscall エントリ、コンテキスト スイッチ
│ §── メモリ/ # 物理ページ アロケータ、VMM、ヒープ
│ §── sched/ # ラウンドロビンスケジューラ
│ §── ipc/ # ポートベースのメッセージパッシング
│ §── proc/ # ELF ローダー;埋め込みユーザースペース BLOB
│ §── hal/ # NVMe、virtio GPU、virtio-net、PS/2 入力、PCI、タイマー
│ §── fs/ # カスタム XFS ファイルシステム
│ §── bsd/ # XNU BSD システムコール実装 (ソケット、フォーク、実行、パイプ)
│ └── エントリ/ # kmain() ブートシーケンス
§── bsd/ # XNU BSD ネットワーキング スタック (TCP/IP、ソケット、virtio-net)
§── ユーザー空間/ # Ring-3 コード
│ §── init/ # PID 1
│ §── runtime/ # Syscall ラッパー (共有 C ライブラリ)
│ §── シェル/ # zsh ポート
│ §── lib/
│ │ §── xgfx/ # グラフィック ライブラリ (パス、塗りつぶし、テキスト、拡大縮小されたテキスト)
│ │ §── thorvg/ # ThorVG ベクター グラフィックス エンジン (SVG レンダリング)
│ │ §── wm/ # ウィンドウマネージャ IPC プロトコルライブラリ
│ │ └─ cpp_runtime/ # 自立環境用の C++ ランタイムのサポート
│ └── サービス/
│ §── combos/ # ディスプレイサーバー + GPU c

重ね合わせ
│ §── ドック/ # ドック底部パネル
│ §── menubar/ # トップメニューバーパネル
│ §── menu/ # 右クリックのコンテキストメニュー (egui、Rust no_std)
│ └── xplorer/ # ファイルマネージャーアプリ
§── third_party/ # ベンダーの依存関係
│ §── egui/ # egui 0.35 (no_std/x86_64-unknown-none に移植)
│ §── egui_software_backend/ # egui プリミティブ用の CPU ラスタライザー
│ └── llvm-project-libcxx/ # 自立型 C++ ランタイム用の libc++
§── tools/ # ビルドスクリプト (QEMU virgl ビルドなど)
§── メイクファイル
━──disk.img # Raw ブロックデバイスイメージ (自動作成)
クリーンアップ
make clean # ビルドアーティファクトと ISO を削除する
make distclean # 取得した Limine ディレクトリも削除します
ライセンス
X OS は、BSL の下でソースから入手できます。これは次のことを意味します。
寄稿者は無料 — フォーク、変更、構築、プル リクエストの送信が可能です。
個人、教育、研究での使用は無料です。
商用利用には有料ライセンスが必要です。販売、サービスとして提供、または製品に埋め込みたい場合は、著作権所有者にお問い合わせください。
すべての商業的権利は留保されます - 著作権所有者はすべての商業的ライセンスを管理し、独自の裁量で商業的使用を許可または拒否することができます。
このモデルでは、すべての商用および入手経路を著作権所有者が完全に管理しながら、コミュニティがプロジェクトを成長させることができます。全条件については、「ライセンス」を参照してください。
商用ライセンス、パートナーシップ、または取得に関するお問い合わせについては、著作権所有者にお問い合わせください。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私のパーを共有しないでください

[切り捨てられた]

## Original Extract

Contribute to valivalivali/x-os development by creating an account on GitHub.

GitHub - valivalivali/x-os · GitHub
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
valivalivali
/
x-os
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
85 Commits 85 Commits boot boot bsd bsd docs docs kernel kernel screenshots screenshots scripts scripts third_party third_party tools tools userspace userspace .gitignore .gitignore .gitmodules .gitmodules AGENTS.md AGENTS.md LICENSE LICENSE Makefile Makefile README.md README.md disk.img disk.img lv_conf.h lv_conf.h View all files Repository files navigation
A forward-looking operating system built from scratch for the AI era.
X OS is a clean-slate microkernel operating system designed for a world where intelligent agents and human users work together. It asks a simple question: what would an OS look like if it were designed today — for agentic programming, programmatic interaction, and human-AI collaboration?
X OS running in QEMU with GPU-accelerated rendering — menu bar, dock with app icons, compositor surfaces, and desktop at 2560×1600.
Software is changing. AI agents don't just call APIs — they read code, write code, run commands, debug, and iterate. They need an environment that is observable, scriptable, and mutable at a level that current operating systems weren't designed to provide.
X OS starts from a different premise. The system is built so that every service is inspectable, every component is replaceable, and every interaction is programmable — not through tooling layered on top, but as a first-class architectural principle. The microkernel exposes IPC ports that any process (human or agent) can use to query, control, and extend the system in real time.
AI-native architecture — the OS assumes agents are first-class users. IPC, service discovery, and system introspection are designed for programmatic use from day one.
Microkernel by design — the kernel handles scheduling, memory, IPC, and hardware. Everything else — display, filesystem, networking, shell — is a userspace service that can be inspected, modified, or replaced while the system runs.
Everything is observable — services communicate over IPC ports with a nameserver for discovery. Any process can query the system state, subscribe to events, or intervene. This is not a debugging feature; it is the architecture.
Clean slate — no decades-old ABI to maintain, no legacy cruft. The system is free to make the right decision for today and for what comes next.
Forward-looking — designed for what's ahead: agentic workflows, programmatic interaction, and human-AI collaboration.
The core idea is fixed: a forward-looking operating system built for the AI era, designed for agentic workflows and intelligent users. That will not change.
Everything else is open. The architecture, the UI, the technologies, the priorities — all will evolve as the project grows and as the landscape of AI-native computing takes shape. Design decisions made today may be revised tomorrow. Components may be rewritten, replaced, or removed. That's not a risk; that's the point.
The desktop environment aims to feel familiar — a menu bar, a dock, windows, a cursor. We lean into known conventions because they work and users already understand them. But this is not a commitment. If we discover something fundamentally better, we will explore it. The goal is to build an interface that is intuitive today and open to reinvention tomorrow.
X OS boots on real hardware (via Limine) and in QEMU, with a working desktop environment:
Microkernel — ~90 syscalls covering scheduling, memory management, IPC ports, timers, GPU/virtio drivers, NVMe storage, and networking
Userspace composer — hardware-accelerated display server with surface compositing, dirty rectangles, window decorations, cursor management, and GPU rendering via virgl (OpenGL ES)
Menu bar — top panel with app menus, dropdowns, and focus tracking
Dock — bottom panel with app icons, hover effects, and app launch/close
Context menu — right-click popup service powered by egui (Rust, no_std )
Shell — embedded zsh port running as a userspace process
File manager (Xplorer) — draggable window app
Networking — XNU BSD network stack ported to the microkernel: TCP/UDP sockets, virtio-net driver, POSIX socket API
Filesystem — custom XFS filesystem with NVMe block device support
GPU acceleration — virtio-gpu with virglrenderer for OpenGL ES 3D compositing; custom GPU command stream for textured quad rendering with alpha blending
egui integration — egui 0.35 ported to no_std / x86_64-unknown-none with a custom CPU rasterizer (scanline glyph renderer) and software backend for immediate-mode UI rendering
Component
Ring
Responsibility
Kernel
0
Scheduling, memory alloc/map, IPC ports, timer, interrupts, NVMe, virtio GPU, virtio-net, PCI
Init (PID 1)
3
First userspace process; spawns services, registers nameserver ports
Composer
3
Display server — surfaces, dirty rects, window decorations, cursor, GPU compositing
Menu Bar
3
Top panel — X logo, app menus, dropdowns, focus tracking
Dock
3
Bottom panel — app icons, hover effects, spawn/hide/show/close
Context Menu
3
Right-click popup — egui-based, Rust no_std staticlib
Shell
3
zsh port — runs as a userspace process
Xplorer
3
File manager — draggable window
BSD Layer
0→3
XNU BSD networking stack (sockets, TCP/IP, virtio-net) adapted for the microkernel
How It Works
The kernel is a minimal microkernel: it handles CPU scheduling, physical/virtual memory, IPC message ports, and hardware drivers (NVMe, virtio-gpu, virtio-net, PS/2 input, PCI, timers). It does not contain filesystem code, display logic, or networking — those are userspace services.
All userspace binaries (init, composer, dock, menubar, menu, shell) are compiled as ELF binaries and embedded directly into the kernel image as byte arrays. The kernel spawns them at boot — no disk loading, no bootloader chain for userspace.
Services communicate via IPC ports. The composer registers a well-known port; apps send surface creation requests, dirty rectangles, and mouse events over IPC. The nameserver ( sys_ns_register / sys_ns_lookup ) provides service discovery.
GPU compositing uses virgl (OpenGL ES via virtio-gpu 3D). Each app surface gets a GPU texture; the composer submits a 3D command stream that draws textured quads with alpha blending onto a framebuffer-backed render target that is also the scanout source.
Tested on macOS with Apple Silicon. You need:
Xcode Command Line Tools (provides clang , make , git )
xorriso (for building the bootable ISO)
QEMU — install via Homebrew or build from source. The Makefile auto-detects the path.
Option A — Homebrew (quickest):
brew install qemu
Option B — build QEMU v11 from source with virgl/ANGLE GPU rendering (recommended for GPU acceleration on Apple Silicon):
# Install the startergo tap for virgl/ANGLE dependencies
brew tap startergo/homebrew-qemu-virgl
brew trust startergo/qemu-virgl
brew install libangle libepoxy-angle virglrenderer spice-protocol spice-server \
meson ninja glib pixman pkg-config dtc vde libssh
# Build QEMU v11.0.2 with Cocoa + virgl + OpenGL ES
bash tools/build-qemu-virgl.sh
This builds QEMU v11.0.2 with virtio-gpu-gl-pci support, Cocoa display, and OpenGL ES rendering via ANGLE/virglrenderer. The script handles patching, building, installing to /opt/qemu-head , fixing dylib paths, and code-signing with HVF entitlements.
The Makefile checks /opt/qemu-head first, then the Homebrew qemu-virgl prefix, then the Homebrew qemu prefix, then falls back to qemu-system-x86_64 in your PATH.
One-time setup (downloads the Limine bootloader):
make setup
Build the bootable ISO:
make
This produces x-os.iso in the project root.
make run-uefi
QEMU is launched with:
virtio-gpu-pci at 2560x1600, Cocoa display
NVMe disk ( disk.img , created automatically if missing)
Serial output forwarded to stdio
x/
├── boot/ # Limine bootloader config and handoff structures
├── kernel/ # Microkernel source
│ ├── arch/x86_64/ # GDT, IDT, syscall entry, context switch
│ ├── memory/ # Physical page allocator, VMM, heap
│ ├── sched/ # Round-robin scheduler
│ ├── ipc/ # Port-based message passing
│ ├── proc/ # ELF loader; embedded userspace blobs
│ ├── hal/ # NVMe, virtio GPU, virtio-net, PS/2 input, PCI, timers
│ ├── fs/ # Custom XFS filesystem
│ ├── bsd/ # XNU BSD syscall implementations (sockets, fork, exec, pipe)
│ └── entry/ # kmain() boot sequence
├── bsd/ # XNU BSD networking stack (TCP/IP, sockets, virtio-net)
├── userspace/ # Ring-3 code
│ ├── init/ # PID 1
│ ├── runtime/ # Syscall wrappers (shared C library)
│ ├── shell/ # zsh port
│ ├── lib/
│ │ ├── xgfx/ # Graphics library (paths, fills, text, scaled text)
│ │ ├── thorvg/ # ThorVG vector graphics engine (SVG rendering)
│ │ ├── wm/ # Window manager IPC protocol library
│ │ └── cpp_runtime/ # C++ runtime support for freestanding environment
│ └── services/
│ ├── composer/ # Display server + GPU compositing
│ ├── dock/ # Bottom dock panel
│ ├── menubar/ # Top menu bar panel
│ ├── menu/ # Right-click context menu (egui, Rust no_std)
│ └── xplorer/ # File manager app
├── third_party/ # Vendored dependencies
│ ├── egui/ # egui 0.35 (ported to no_std/x86_64-unknown-none)
│ ├── egui_software_backend/ # CPU rasterizer for egui primitives
│ └── llvm-project-libcxx/ # libc++ for freestanding C++ runtime
├── tools/ # Build scripts (QEMU virgl build, etc.)
├── Makefile
└── disk.img # Raw block device image (auto-created)
Cleaning Up
make clean # remove build artifacts and ISO
make distclean # also remove fetched Limine directory
License
X OS is source-available under the BSL. This means:
Free for contributors — you can fork, modify, build, and send pull requests.
Free for personal, educational, and research use.
Commercial use requires a paid license — if you want to sell it, offer it as a service, or embed it in a product, contact the copyright holder.
All commercial rights reserved — the copyright holder controls all commercial licensing and may grant or deny commercial use at their sole discretion.
This model lets the community grow the project while keeping all commercial and acquisition paths fully controlled by the copyright holder. See LICENSE for full terms.
For commercial licensing, partnership, or acquisition inquiries, contact the copyright holder.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my per

[truncated]
