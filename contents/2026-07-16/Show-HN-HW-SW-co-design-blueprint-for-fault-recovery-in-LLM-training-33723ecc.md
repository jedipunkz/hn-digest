---
source: "https://github.com/PJHkorea/pim-hbm-bypass"
hn_url: "https://news.ycombinator.com/item?id=48940625"
title: "Show HN: HW/SW co-design blueprint for fault recovery in LLM training"
article_title: "GitHub - PJHkorea/pim-hbm-bypass: Software-Defined PIM-HBM Bypass Infrastructure. 0ns JAX/XLA memory view fusion with real-time asynchronous NCCL collective fault-scanning & no-recompile Mux hot-swapping. · GitHub"
author: "PJHkorea"
captured_at: "2026-07-16T21:56:21Z"
capture_tool: "hn-digest"
hn_id: 48940625
score: 2
comments: 1
posted_at: "2026-07-16T21:39:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: HW/SW co-design blueprint for fault recovery in LLM training

- HN: [48940625](https://news.ycombinator.com/item?id=48940625)
- Source: [github.com](https://github.com/PJHkorea/pim-hbm-bypass)
- Score: 2
- Comments: 1
- Posted: 2026-07-16T21:39:40Z

## Translation

タイトル: HN の表示: LLM トレーニングにおける障害回復のための HW/SW 共同設計ブループリント
記事のタイトル: GitHub - PJHkorea/pim-hbm-bypass: ソフトウェア定義の PIM-HBM バイパス インフラストラクチャ。 0ns JAX/XLA メモリ ビューの融合と、リアルタイム非同期 NCCL 集団フォールト スキャンおよび再コンパイル不要の Mux ホットスワップ。 · GitHub
説明: ソフトウェア定義の PIM-HBM バイパス インフラストラクチャ。 0ns JAX/XLA メモリ ビューの融合と、リアルタイム非同期 NCCL 集団フォールト スキャンおよび再コンパイル不要の Mux ホットスワップ。 - PJHkorea/pim-hbm-bypass

記事本文:
GitHub - PJHkorea/pim-hbm-bypass: ソフトウェア定義の PIM-HBM バイパス インフラストラクチャ。 0ns JAX/XLA メモリ ビューの融合と、リアルタイム非同期 NCCL 集団フォールト スキャンおよび再コンパイル不要の Mux ホットスワップ。 · GitHub
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
PJH韓国
/
pim-hbm-バイパス

公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
40 コミット 40 コミット CMakeLists.txt CMakeLists.txt ライセンス ライセンス README.md README.md hardware_fault_recovery.py hardware_fault_recovery.py hardware_fault_recovery_distributed.py hardware_fault_recovery_distributed.py llama3_layer_adapter.py llama3_layer_adapter.py pim_hardware_gate.py pim_hardware_gate.py pim_hbm_core.cu pim_hbm_core.cu topology_sharding.py topology_sharding.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
0ns フレームワーク メモリ ビュー フュージョンと障害テレメトリを探索する実験的な PIM-HBM ハードウェア共同設計サブシステム
このプロジェクトは、次世代アクセラレータ インフラストラクチャ環境内のソフトウェア抽象化の断片化の障壁を調査し、解決するために設計されたハードウェアとソフトウェアの共同設計プロトタイプです。
このサブシステムは、低レベルの物理キャッシュ ライン アライメント メカニズムを上位ハイパフォーマンス フレームワーク (JAX/XLA) アドレス インターフェイスと直接連動させて、統合された代数パイプラインにすることで、分散クラスター操作中の静的コンパイル ストールとハードウェア タイミング ジッターを安定して軽減する構造的な実行可能性を探ります。
0ns メモリ コピー バイパス : __cuda_array_interface__ 仕様を活用して、低レベル C++ 物理アドレス ラインを JAX tensor バスに直接リンクし、ホスト デバイスのハードウェア メモリ コピー オーバーヘッドを構造的に排除します。
Pure Branchless Loop : 条件分岐 ( if ) を完全に根絶し、代わりに正確な三項演算とレジスタ常駐データの再利用パターンを導入して、コンパイラが条件付き移動プリミティブ ( SEL/PRMT ) を強制することを保証します。
ワープレベルの動的境界 : 潜在的な境界外のメモをインターセプトします

__shfl_down_sync を介してワープレベルのバイナリ ツリー リダクション ファイアウォールを駆動し、アクティブな最大存続アドレス オフセットを計算することにより、ラグド テール グリッド内の障害 ( SegFault ) を検出します。
Algebraic Insulation Gate : 専用の stop_gradient 回路を使用して、JAX ランタイムの背後でハードウェア障害信号と浮動小数点発散異常を捕捉し、バックプロパゲーション チェーンの汚染を物理層で分離します。
ダイナミック ホットプラグ リカバリ: アクティブ クラスターの実行中に物理 HBM バンクに障害が発生すると、このメカニズムは集合通信 (NCCL) の停止やグラフの再コンパイルを引き起こすことなく、破損したデバイス スロットに対してリアルタイムの 64 ビット アドレス ワイヤ ホットスワップを試みます。
LICENSE : Apache License 2.0 仕様に基づく法的保護措置と特許報復防御条項を宣言します。
CMakeLists.txt : システム ハードウェア アーキテクチャ トポロジと pybind11 コンパイル パスを自動的に追跡して、最終的な共有オブジェクト ( .so ) ライブラリを生成するビルド オーケストレーター。
pim_hbm_core.cu : alignas(32) キャッシュライン マッチング、__activemask() 動的アドレス ファイアウォール、__ldg 高速読み取りレール (C++/CUDA) を実装するコア ブランチレス数学的アクセラレーション カーネル。
pim_hardware_gate.py : ShapeDtypeStruct 仮想抽象トレーサを利用して、0MB の物理 VRAM フットプリント (Python/JAX) を維持しながら XLA コンパイラ マシン コードをロックする、プリウォーミングおよびバックプロパゲーション チェーン絶縁層。
topology_sharding.py : ノードごとの VRAM 物理アドレス行をインターセプトして、ゼロコピーの NamedSharding グローバル分散マトリックス ビュー (Python/JAX) を確立するマクロレベルのトポロジ コントロール タワー。
hardware_fault_recovery.py : 分散重みマトリックス内のバックグラウンド障害スキャンを処理し、ルーティングを切り替えるリアルタイムの高可用性ホットプラグ スワップ エンジン

事前に予約された緊急バックアップ アドレス プール (Python/JAX)。
hardware_fault_recovery_distributed.py : 大規模インフラストラクチャ向けに設計された集合スキャンおよび回復エンジンで、ワイヤレベルの NCCL All-Reduce フュージョンと np. flatnonzero ベクトル化障害抽出 (Python/JAX) を活用します。
llama3_layer_adapter.py : ブランチレス フォールト トレラントな前方実行バス (Python/JAX) と並行して、Llama-3-8B ディメンション (4096 / 14336) に一致する直接 0ns アドレス取り込みを実行するトランス層アダプター プラグイン。
NVIDIA Ampere (A100) または Hopper (H100/H200) 環境で構成された高性能クラスター ターミナル内で次のコマンドを順番に実行して、メモリ ビュー フュージョン モードとハードウェア フォールト トレラント エミュレーション エンジンを起動します。
# 1. ビルド シーケンス専用の分離ディレクトリを作成して入力します。
mkdir ビルド && cd ビルド
# 2. クロスコンパイル アーキテクチャ スキャンを開始します (pybind11 および CUDA パスを自動的に追跡します)
cmake ..
# 3. ハードウェア マシンコード ライブラリをビルドしてコンパイルします (pim_hbm_bridge_core.so を抽出します)。
make -j \$ (nproc)
# 4. 抽出した共有ライブラリモジュールを上位の実行ディレクトリに移行します。
cp pim_hbm_bridge_core * .so .. && cd ..
# 5. [ステップ A] シングルノード JAX プリウォーミング エンジンと代数絶縁ガードを実行する
python3 pim_hardware_gate.py
# 6. [ステップ B] マクロレベルの分散シャーディング トポロジ仮想ビュー フュージョン マトリックスの起動
python3トポロジー_シャーディング.py
# 7. [ステップ C] ワイヤレベルの NCCL All-Reduce を介して、集合的ヘルス スキャンと動的ホットプラグ リカバリを実行します。
python3 hardware_fault_recovery_distributed.py
# 8. [⚡ ステップ D] Llama-3-8B (4096/14336) マトリックスに合わせたゼロコピー アドレス アダプター インフラストラクチャをインスタンス化する
python3 llama3_layer_adapter.py
📜ライセンス
このプロジェクトは dist です

ributed under the terms of the Apache License 2.0 。 You are free to modify and distribute the software, provided that the original copyright notice and license disclosure obligations are fully preserved.
Experimental PIM-HBM Hardware Co-Design Subsystem exploring 0ns Framework Memory View Fusion & Fault Telemetry
このプロジェクトは、次世代アクセラレータインフラストラクチャ環境で発生する可能性があるソフトウェア層間の断片化の障壁を研究するために設計された、ハードウェア - ソフトウェア共同設計（Co-design）プロトタイプです。
低レベルの物理キャッシュラインアラインメントメカニズムと上位高性能フレームワーク（JAX / XLA）の間のアドレスラインインターフェイスを単一の代数パイプラインに接続して、分散クラスタの稼働中に発生する静的コンパイルとハードウェアジッタを確実に制御できる可能性を探ります。
⚡コアアーキテクチャ特性(Key Innovations)
0ns Memory Copy Bypass : __cuda_array_interface__ 規格を活用して C++ 機械語アドレス線を JAX テンソルバスに直結することにより、ホスト-デバイス間の物理コピーオーバーヘッドを構造的に解消します。
Pure Branchless Loop : 条件分岐文 ( if ) を完全に排除し、三項演算およびレジスタ常駐データ再利用構造を精密構成し、コンパイラレベルの条件付き移動命令 (SEL/PRMT ) 出力を導き出します。
Warp-level Dynamic Bounds：最後のグリッド方位領域で発生する可能性があるメモリ参照エラー（SegFault）を防ぐために、__shfl_down_syncベースのワープ内のバイナリツリー最大生存アドレスダイナミックリダクションファイアウォールを駆動します。
Algebraic Insulation Gate: JAXランタイムにおけるハードウェア障害信号と数値発散誤差を停止_

gradient回路で捕捉し、微分鎖汚染をフィジカルレベルで隔離および絶縁します。
Dynamic Hot-Plugging Recovery：アクセラレータクラスタの駆動中に特定のHBMバンクに物理的な障害が発生した場合、グローバル通信（NCCL）の中断やグラフの再コンパイルなしで、不良デバイスの64ビットアドレス線だけをリアルタイムでバイパススワップ（Hot-Swapping）するメカニズムを試みます。
📂 ファイルトポロジ (Repository Structure)
ライセンス: Apache License 2.0 に基づく法的ファイアウォールおよび特許保護条項の指定
CMakeLists.txt：システムアクセラレータアーキテクチャ環境とpybind11コンパイルパスを自動追跡して共有ライブラリ（.so）を射出するビルドオーケストレータ
pim_hbm_core.cu：alignas（32）キャッシュラインマッチレイアウト、__activemask（）ダイナミックアドレスファイアウォール、および__ldgアクセラレーションレールが注入された無分岐数学加速カーネルコア（C ++ / CUDA）
pim_hardware_gate.py : ShapeDtypeStruct 仮想抽象化トレーサーを活用して実 VRAM 占有 0MB 状態で XLA コンパイラ機械語を固定する予熱および微分鎖絶縁層 (Python/JAX)
topology_sharding.py : 大規模クラスタノードごとのVRAM物理アドレス線を傍受し、ゼロコピー NamedSharding グローバル分散マトリックスビューを確立する巨視的トポロジ制御塔 (Python/JAX)
hardware_fault_recovery.py : 分散重み行列内の不良バンクバックグラウンドスキャンと緊急バックアッププールのアドレス線を活用したリアルタイム無中断ホットプラグスワップエンジン (Python/JAX)
hardware_fault_recovery_distributed.py : 超大型インフラストラクチャのためのNCCL All-Reduceワイヤレベル融合集約(Collective)スキャンとnp.flatnonzeroベクトル化欠陥抽出回復エンジン(Python/JAX)
llama3_layer_adapter.py:Llama-3-

8Bユニークディメンション(4096/14336)
🛠️高速駆動とビルドガイドライン(クイックスタート)
NVIDIA Ampere（A100）またはHopper（H100 / H200）環境を備えた高性能クラスタターミナルで、次のコマンドを順番に起動して、バイパスモードおよびハードウェア障害許容エンジンエミュレーションを爆発させることができます。
#1.ビルド専用の分離空間の作成とエントリ
mkdir build && cd build
#2.クロスコンパイルアーキテクチャスキャン操作（pybind11とCUDAコンパイルパスの自動追跡）
cmake ..
#3.ハードウェア機械語ライブラリコンパイルビルド（pim_hbm_bridge_core.so抽出）
make -j \$(nproc)
#4. 親実行ディレクトリに抽出されたモジュールコピーエスカレーション
cp pim_hbm_bridge_core*.so..&&cd..
#5. [STEP A] シングルノード専用JAX予熱エンジンと代数誤差絶縁ガード稼働
python3 pim_hardware_gate.py
#6. [STEP B] Multi-GPUマクロ分散シャーディングトポロジー仮想融合テンソル稼働
python3 topology_sharding.py
#7. [STEP C] 超大型インフラストラクチャのためのNCCL All-Reduce融合分散アグリゲーションヘルススキャンとホットプラグリカバリ稼働
python3 hardware_fault_recovery_distributed.py
#8. [⚡ STEP D] 本番 Llama-3-8B Transformer 4096/14336 マトリックス 0ns アドレスゼロコピーアダプタインフラストラクチャ
python3 llama3_layer_adapter.py
📜ライセンス(License)
このプロジェクトはApache License 2.0に基づいて配布されています。自由な修正と配布は可能ですが、著作権およびライセンス通知の義務が伴います。
Software-Defined PIM-HBM Bypass Infrastructure. 0ns JAX/XLA memory view fusion with real-time asynchronous NCCL collective fault-scanning & no-recompile Mux hot-s

ワッピング。
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

Software-Defined PIM-HBM Bypass Infrastructure. 0ns JAX/XLA memory view fusion with real-time asynchronous NCCL collective fault-scanning & no-recompile Mux hot-swapping. - PJHkorea/pim-hbm-bypass

GitHub - PJHkorea/pim-hbm-bypass: Software-Defined PIM-HBM Bypass Infrastructure. 0ns JAX/XLA memory view fusion with real-time asynchronous NCCL collective fault-scanning & no-recompile Mux hot-swapping. · GitHub
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
PJHkorea
/
pim-hbm-bypass
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
40 Commits 40 Commits CMakeLists.txt CMakeLists.txt LICENSE LICENSE README.md README.md hardware_fault_recovery.py hardware_fault_recovery.py hardware_fault_recovery_distributed.py hardware_fault_recovery_distributed.py llama3_layer_adapter.py llama3_layer_adapter.py pim_hardware_gate.py pim_hardware_gate.py pim_hbm_core.cu pim_hbm_core.cu topology_sharding.py topology_sharding.py View all files Repository files navigation
Experimental PIM-HBM Hardware Co-Design Subsystem exploring 0ns Framework Memory View Fusion & Fault Telemetry
This project is a hardware-software co-design prototype engineered to investigate and resolve software abstraction fragmentation barriers within next-generation accelerator infrastructure environments.
By interlocking the low-level physical cache-line alignment mechanisms directly with the upper high-performance framework (JAX/XLA) address interfaces into a unified algebraic pipeline, this subsystem explores the structural viability of stably mitigating static compilation stalls and hardware timing jitters during distributed cluster operations.
0ns Memory Copy Bypass : Leverages the __cuda_array_interface__ specification to directly link low-level C++ physical address lines to the JAX tensor bus, structurally eliminating host-device hardware memory copy overhead.
Pure Branchless Loop : Completely eradicates conditional branches ( if ), instead deploying precise ternary operations and register-resident data reuse patterns to guarantee the compiler forces conditional move primitives ( SEL/PRMT ).
Warp-level Dynamic Bounds : Intercepts potential out-of-bound memory faults ( SegFault ) inside ragged tail grids by driving a warp-level binary tree reduction firewall via __shfl_down_sync to compute active maximum surviving address offsets.
Algebraic Insulation Gate : Captures hardware fault signals and floating-point divergence anomalies behind the JAX runtime using dedicated stop_gradient circuits, isolating backpropagation chain contamination at the physical layer.
Dynamic Hot-Plugging Recovery : Upon physical HBM bank failure during active cluster runs, this mechanism attempts real-time 64-bit address wire hot-swapping for the corrupted device slot without causing collective communication (NCCL) stalls or graph recompilation.
LICENSE : Declares legal safeguards and patent retaliation defense clauses under the Apache License 2.0 specification.
CMakeLists.txt : The build orchestrator that automatically tracks system hardware architecture topologies and pybind11 compilation paths to emit the final shared object ( .so ) libraries.
pim_hbm_core.cu : The core branchless mathematical acceleration kernel implementing alignas(32) cache-line matching, __activemask() dynamic address firewalls, and __ldg high-speed read rails (C++/CUDA).
pim_hardware_gate.py : The pre-warming and backpropagation chain insulation layer utilizing ShapeDtypeStruct virtual abstract tracers to lock XLA compiler machine code while maintaining 0MB of physical VRAM footprint (Python/JAX).
topology_sharding.py : The macro-level topology control tower that intercepts per-node VRAM physical address lines to establish zero-copy NamedSharding global distributed matrix views (Python/JAX).
hardware_fault_recovery.py : The real-time, high-availability hot-plugging swap engine that handles background fault scans within the distributed weight matrices and switches routing to the pre-reserved emergency backup address pool (Python/JAX).
hardware_fault_recovery_distributed.py : The collective scanning and recovery engine engineered for large-scale infrastructure, leveraging wire-level NCCL All-Reduce fusion and np.flatnonzero vectorized fault extraction (Python/JAX).
llama3_layer_adapter.py : A transformer layer adapter plugin that conducts direct 0ns address ingestion matching Llama-3-8B dimensions (4096 / 14336) alongside branchless fault-tolerant forward execution buses (Python/JAX).
Execute the following commands sequentially within a high-performance cluster terminal configured with NVIDIA Ampere (A100) or Hopper (H100/H200) environments to ignite the memory view fusion mode and the hardware fault-tolerant emulation engine.
# 1. Create and enter an isolated directory dedicated to the build sequence
mkdir build && cd build
# 2. Launch cross-compilation architecture scanning (Automatically tracks pybind11 and CUDA paths)
cmake ..
# 3. Build and compile the hardware machine-code library (Extracts pim_hbm_bridge_core.so)
make -j \$ (nproc)
# 4. Migrate the extracted shared library module back to the upper execution directory
cp pim_hbm_bridge_core * .so .. && cd ..
# 5. [STEP A] Execute the single-node JAX pre-warming engine and algebraic insulation guard
python3 pim_hardware_gate.py
# 6. [STEP B] Launch the macro-level distributed sharding topology virtual view fusion matrix
python3 topology_sharding.py
# 7. [STEP C] Run the collective health scan and dynamic hot-plugging recovery via wire-level NCCL All-Reduce
python3 hardware_fault_recovery_distributed.py
# 8. [⚡ STEP D] Instantiate the zero-copy address adapter infrastructure tailored for Llama-3-8B (4096/14336) matrices
python3 llama3_layer_adapter.py
📜 License
This project is distributed under the terms of the Apache License 2.0 . You are free to modify and distribute the software, provided that the original copyright notice and license disclosure obligations are fully preserved.
Experimental PIM-HBM Hardware Co-Design Subsystem exploring 0ns Framework Memory View Fusion & Fault Telemetry
본 프로젝트는 차세대 가속기 인프라 환경에서 발생할 수 있는 소프트웨어 계층 간 파편화 장벽을 연구하기 위해 설계된 하드웨어-소프트웨어 공동 설계(Co-design) 프로토타입 입니다.
저수준의 물리 캐시라인 정렬 매커니즘과 상위 고성능 프레임워크(JAX/XLA) 간의 주소선 인터페이스를 단일 대수 파이프라인으로 연결하여, 분산 클러스터 가동 중 유발되는 정적 컴파일 렉 및 하드웨어 지터를 안정적으로 제어할 수 있는 가능성을 탐구합니다.
⚡ 핵심 아키텍처 특성 (Key Innovations)
0ns Memory Copy Bypass : __cuda_array_interface__ 규격을 활용해 C++ 기계어 주소선을 JAX 텐서 버스에 직결함으로써, 호스트-디바이스 간 물리 복사 오버헤드를 구조적으로 해소합니다.
Pure Branchless Loop : 조건 분기문( if )을 완전히 배제하고 삼항 연산 및 레지스터 상주 데이터 재사용 구조를 정밀 구성하여, 컴파일러 수준의 조건부 이동 명령어( SEL/PRMT ) 출력을 유도합니다.
Warp-level Dynamic Bounds : 마지막 그리드 자투리 영역에서 발생할 수 있는 메모리 참조 오류(SegFault)를 방지하기 위해, __shfl_down_sync 기반 워프 내 2진 트리 최대 생존 주소 동적 리덕션 방화벽을 구동합니다.
Algebraic Insulation Gate : JAX 런타임에서 하드웨어 결함 신호 및 수치 발산 오차를 stop_gradient 회로로 포획하여, 미분 사슬 오염을 피지컬 레벨에서 격리 및 절연합니다.
Dynamic Hot-Plugging Recovery : 가속기 클러스터 구동 중 특정 HBM 뱅크에 물리적 결함 발생 시, 글로벌 통신(NCCL) 중단 및 그래프 재컴파일 없이 오직 불량 장치의 64비트 주소선만 실시간으로 우회 스와핑(Hot-Swapping)하는 메커니즘을 시도합니다.
📂 파일 토폴로지 (Repository Structure)
LICENSE : Apache License 2.0 의거 법적 방화벽 및 특허 보호 조항 명시
CMakeLists.txt : 시스템 가속기 아키텍처 환경과 pybind11 컴파일 패스를 자동 추적하여 공유 라이브러리( .so )를 사출하는 빌드 오케스트레이터
pim_hbm_core.cu : alignas(32) 캐시라인 일치 레이아웃, __activemask() 동적 주소 방화벽 및 __ldg 가속 레일이 주입된 무분기 수학 가속 커널 코어 (C++/CUDA)
pim_hardware_gate.py : ShapeDtypeStruct 가상 추상화 트레이서를 활용해 실재 VRAM 점유 0MB 상태로 XLA 컴파일러 기계어를 고정하는 예열 및 미분 사슬 절연 레이어 (Python/JAX)
topology_sharding.py : 대규모 클러스터 노드별 VRAM 물리 주소선을 가로채어 제로카피 NamedSharding 글로벌 분산 매트릭스 뷰를 수립하는 거시적 토폴로지 관제탑 (Python/JAX)
hardware_fault_recovery.py : 분산 가중치 행렬 내 불량 뱅크 백그라운드 스캔 및 비상 백업 풀 주소선을 활용한 실시간 무중단 핫플러깅 스와프 엔진 (Python/JAX)
hardware_fault_recovery_distributed.py : 초대형 인프라를 위한 NCCL All-Reduce 와이어 레벨 융합 집산(Collective) 스캔 및 np.flatnonzero 벡터화 결함 적출 복구 엔진 (Python/JAX)
llama3_layer_adapter.py : Llama-3-8B 고유 차원(4096 / 14336) 직통 0ns 주소 인입 및 무분기 결함 허용 순방향 훈련/추론 버스 어댑터 플러그인 (Python/JAX)
🛠️ 고속 구동 및 빌드 지침 (Quick Start)
NVIDIA Ampere(A100) 또는 Hopper(H100/H200) 환경이 구축된 고성능 클러스터 터미널에서 다음 명령을 순차 가동하여 바이패스 모드 및 하드웨어 결함 허용(Fault-Tolerant) 엔진 에뮬레이션을 기폭할 수 있습니다.
# 1. 빌드 전용 격리 공간 생성 및 진입
mkdir build && cd build
# 2. 크로스 컴파일 아키텍처 스캔 가동 (pybind11 및 CUDA 컴파일 패스 자동 추적)
cmake ..
# 3. 하드웨어 기계어 라이브러리 컴파일 빌드 (pim_hbm_bridge_core.so 추출)
make -j \$ (nproc)
# 4. 상위 실행 디렉토리로 추출된 모듈 복사 이관
cp pim_hbm_bridge_core * .so .. && cd ..
# 5. [STEP A] 단일 노드 전용 JAX 예열 엔진 및 대수적 오차 절연 가드 가동
python3 pim_hardware_gate.py
# 6. [STEP B] Multi-GPU 거시 분산 샤딩 토폴로지 가상 융합 텐서 가동
python3 topology_sharding.py
# 7. [STEP C] 초대형 인프라용 NCCL All-Reduce 융합 분산 집산 헬스 스캔 및 핫플러깅 복구 가동
python3 hardware_fault_recovery_distributed.py
# 8. [⚡ STEP D] 실전 Llama-3-8B Transformer 4096/14336 매트릭스 0ns 주소 제로카피 어댑터 인프라 가동
python3 llama3_layer_adapter.py
📜 라이선스 (License)
본 프로젝트는 Apache License 2.0 의거하여 배포됩니다. 자유로운 수정 및 배포가 가능하나 저작권 및 라이선스 고지 의무가 수반됩니다.
Software-Defined PIM-HBM Bypass Infrastructure. 0ns JAX/XLA memory view fusion with real-time asynchronous NCCL collective fault-scanning & no-recompile Mux hot-swapping.
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
