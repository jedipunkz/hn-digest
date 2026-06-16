---
source: "https://github.com/matheusdelgado/infinite-context"
hn_url: "https://news.ycombinator.com/item?id=48555259"
title: "Infinite Context Paging Engine – Zero-copy LLM context paging in Rust ~419.34 µs"
article_title: "GitHub - matheusdelgado/infinite-context: An ultra-low latency, zero-copy context virtual memory paging engine written in Rust, designed to break physical VRAM limitations for LLMs and autonomous agents using attention-driven predictive prefetching. · GitHub"
author: "matheusdelgs"
captured_at: "2026-06-16T13:54:58Z"
capture_tool: "hn-digest"
hn_id: 48555259
score: 1
comments: 0
posted_at: "2026-06-16T13:47:15Z"
tags:
  - hacker-news
  - translated
---

# Infinite Context Paging Engine – Zero-copy LLM context paging in Rust ~419.34 µs

- HN: [48555259](https://news.ycombinator.com/item?id=48555259)
- Source: [github.com](https://github.com/matheusdelgado/infinite-context)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T13:47:15Z

## Translation

タイトル: Infinite Context Paging Engine – Rust のゼロコピー LLM コンテキスト ページング ~419.34 μs
記事のタイトル: GitHub - matheusdelgado/infinite-context: Rust で書かれた超低レイテンシ、ゼロコピーコンテキスト仮想メモリ ページング エンジン。アテンション駆動型の予測プリフェッチを使用して、LLM と自律エージェントの物理 VRAM 制限を突破するように設計されています。 · GitHub
説明: Rust で書かれた超低遅延のゼロコピー コンテキスト仮想メモリ ページング エンジンで、アテンション駆動の予測プリフェッチを使用して LLM および自律エージェントの物理 VRAM 制限を突破するように設計されています。 - matheusdelgado/infinite-context

記事本文:
GitHub - matheusdelgado/infinite-context: Rust で書かれた超低レイテンシのゼロコピーコンテキスト仮想メモリ ページング エンジンで、アテンション駆動型の予測プリフェッチを使用して LLM と自律エージェントの物理 VRAM 制限を打ち破るように設計されています。 · GitHub
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
を却下する

警告します
{{ メッセージ }}
マテウスデルガド
/
無限コンテキスト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
matheusdelgado/infinite-context
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット Assets lib lib src src .gitignore .gitignore Cargo.toml Cargo.toml README.md README.md build.rs build.rs test_engine.py test_engine.py すべてのファイルを表示 リポジトリ ファイル ナビゲーション
ICPE: 無限コンテキスト ページング エンジン 🧠
Rust で書かれた超低レイテンシのゼロコピー コンテキスト仮想メモリ ページング エンジンで、LLM および長期存続する自律エージェントの物理 VRAM 制限を打ち破るように設計されています。
ICPE は、LLM トークン コンテキスト レイヤーをオペレーティング システムの仮想メモリ (ページング/スワップ) とまったく同じように扱います。 ICPE は、高価な GPU VRAM に大量の低アクティベーション履歴を保持する代わりに、アテンション駆動型予測エビクション アルゴリズムを利用して、メモリ マップ ファイル (mmap) 経由でコールド コンテキストをディスクにページアウトし、次の推論ステップの前にホット スライスを高速メモリ ナノ秒でプリフェッチします。
🛡️ バイナリのセキュリティ、アーキテクチャ、コンプライアンス
このプロジェクトは 100% Rust で書かれています。公開評価中にコアの知的財産と独自のアルゴリズムを保護するために、高性能の予測エンジンとスレッド同期ヒューリスティックは、事前にコンパイルされ、高度に最適化された Rust バイナリ BLOB ( .a / .so ) として /lib ディレクトリに配布されます。
提供されるターゲット アーキテクチャ: x86_64-unknown-linux-gnu および aarch64-unknown-linux-gnu (ARM64/Graviton 準拠)。
コンプライアンス: すべてのバイナリは暗号化署名され、公開された独立した GitHub Actions ワークフローを通じて構築されます。 SHA-256 ハッシュは実行時に検証されます。コアには 0% の外部ネットワーキング、0% のテレメトリが含まれており、str を操作します。

ローカル システム メモリの範囲内でのみ動作します。
取得の予定: コア エンジンの完全なクリーン ルーム ソース コード、数学的仕様、コンパイル ツールチェーン、およびグローバル IP 所有権は、完全な取得のために厳密に予約されています。
📊 検証可能なベンチマーク (基準リリースモード)
ICPE は、 memmap2 と zerocopy を使用して実行エンジンをカーネル ページ キャッシュ スペースに直接マッピングすることで、標準の I/O システムコール オーバーヘッドを排除します。
プリフェッチとエビクションのレイテンシー: 継続的な同時スレッド ストレス下で ~419.34 µs (マイクロ秒)、FFI 境界を安全に越えて保護されたコアに到達します。
メモリ コピー オーバーヘッド: 0% (トゥルー ゼロ コピー バイト キャスティング)。
RAM フットプリント: 決定的で固定され、完全に制限されています。
このパブリック リポジトリは、厳格なオープンコア評価ライセンスに基づいて動作します。アーキテクチャ、Python ラッパー、ベンチマーク テスト スイートは完全にオープンで検証可能です。統合テストをローカルで自由にネイティブにコンパイル、ベンチマーク、実行できます。エンタープライズ ライセンスや完全な IP 取得がない場合、プリコンパイルされたコアの商用利用、運用展開、またはクラウド インフラストラクチャへの組み込みは固く禁止されています。
🚀 実行方法とパフォーマンスの確認方法
プロジェクトをネイティブにコンパイルし、ローカル インフラストラクチャ上で直接ベンチマーク要求を監査できます。
Rust ツールチェーン、Python 3.12 開発ヘッダー、およびネイティブ リンカーがマシンにインストールされていることを確認します。
sudo aptアップデート
sudo apt install build-essential python3-dev python3-config lld
2. ローカル ベンチマーク (基準) を検証する
マイクロベンチマーク スイートは、Python ランタイム コンテキスト シンボルの衝突を防ぐために、コア ソース ファイル内で分離されています。統計的なハードウェア遅延レポートを実行するには、次のコマンドを実行します。
# 古いリンカーメタデータをクリアします
貨物をきれいにする
# ターゲットコンテキストマネージャーベンチマークを実行します su

それ
カーゴベンチ --bench context_manager_bench
詳細な統計分布曲線は、target/criterion/report/index.html で生成されます。
3. Python 拡張モジュールをテストする
ネイティブ拡張機能をローカルの Python 環境に構築し、統合パイプライン テストを実行します。
# ローカル仮想環境をアクティブ化します
ソース .venv/bin/activate
# コンパイルラッパーをインストールする
pip インストール マチュリン
# コンパイル済みの高性能コアを使用してプロジェクトをコンパイルします
maturin 開発 --release
# ライブ エージェント コンテキスト スワッピング シミュレーションを実行する
python3 テストエンジン.py
💼 企業開発とM&A (合併と買収)
このテクノロジーは上級システム エンジニアによって設計され、クラウド インフラストラクチャの直接統合 (例: AWS Bedrock、Google Vertex AI、Meta GenAI クラスター ノード) 向けに最適化されています。
エンジニアリングへの深い焦点を保護するために、当社は厳密に非同期で文書化を優先した M&A パイプラインを運用しています。当社は、入門ディスカバリ コール、入門 Zoom セッション、または緩い技術調整会議には参加しません。
ソフトウェア自体がすべてを物語ります。エンジニアリング チームがカーゴ ベンチを実行し、マイクロ ベンチマークを監査し、ICPE が VRAM/遅延のボトルネックを解決していることを確認した場合、取得ワークフローは次のようになります。
提出: 企業開発/M&A チームは、正式な書面による意向表明書 (LOI) または企業買収提案を、正式な企業ドメインから当社の安全なチャネルに直接提出します。
デューデリジェンス: 有効な LOI を受け取ると、当社は標準的な NDA に基づいて、クリーンルームのソース コード、数学的仕様、包括的なファジング レポート、および自動ハードウェア ストレス テストを含む安全なデータ ルームへのアクセスを即座に許可します。
クロージング: 完全に自動化された IP 移転と法的クロージング。
📩 LOI & Fi の公式チャンネル

rm 提案: corpdev@matheusdelgado.com
Rust で書かれた超低レイテンシのゼロコピー コンテキスト仮想メモリ ページング エンジンで、アテンション駆動の予測プリフェッチを使用して LLM および自律エージェントの物理 VRAM 制限を突破するように設計されています。
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

An ultra-low latency, zero-copy context virtual memory paging engine written in Rust, designed to break physical VRAM limitations for LLMs and autonomous agents using attention-driven predictive prefetching. - matheusdelgado/infinite-context

GitHub - matheusdelgado/infinite-context: An ultra-low latency, zero-copy context virtual memory paging engine written in Rust, designed to break physical VRAM limitations for LLMs and autonomous agents using attention-driven predictive prefetching. · GitHub
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
matheusdelgado
/
infinite-context
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
matheusdelgado/infinite-context
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits assets assets lib lib src src .gitignore .gitignore Cargo.toml Cargo.toml README.md README.md build.rs build.rs test_engine.py test_engine.py View all files Repository files navigation
ICPE: Infinite Context Paging Engine 🧠
An ultra-low latency, zero-copy context virtual memory paging engine written in Rust, designed to break physical VRAM limitations for LLMs and Long-Lived Autonomous Agents.
ICPE treats LLM token context layers exactly like operating system virtual memory (Paging/Swap). Instead of holding massive, low-activation histories in expensive GPU VRAM, ICPE utilizes an Attention-Driven Predictive Eviction algorithm to page out cold contexts to disk via memory-mapped files (mmap), prefetching hot slices back into high-speed memory nanoseconds before the next inference step.
🛡️ Binary Security, Architecture & Compliance
This project is written 100% in Rust . To protect our core intellectual property and proprietary algorithms during public evaluation, the high-performance predictive engine and thread synchronization heuristics are distributed as pre-compiled, highly optimized Rust binary blobs ( .a / .so ) located in the /lib directory.
Target Architectures Provided: x86_64-unknown-linux-gnu and aarch64-unknown-linux-gnu (ARM64/Graviton compliant).
Compliance: All binaries are cryptographically signed and built via public, isolated GitHub Actions workflows. SHA-256 hashes are verified at runtime. The core contains 0% external networking, 0% telemetry, and operates strictly within local system memory bounds.
What is Up for Acquisition: Full clean-room source code of the core engine, mathematical specifications, compilation toolchains, and global IP ownership are strictly reserved for total acquisition.
📊 Verifiable Benchmarks (Criterion Release Mode)
ICPE eliminates standard I/O syscall overhead by mapping the execution engine directly into the kernel page cache space using memmap2 and zerocopy .
Prefetch & Eviction Latency: ~419.34 µs (Microseconds) under continuous concurrent thread stress, crossing the FFI boundary safely into the protected core.
Memory Copy Overhead: 0% (True Zero-Copy byte casting).
RAM Footprint: Deterministic, fixed, and completely bounded.
This public repository operates under a strict Open-Core Evaluation License. The architecture, Python wrappers, and benchmarking test-suites are fully open and verifiable. You are free to natively compile, benchmark, and run integration tests locally. Commercial use, production deployment, or cloud infrastructure embedding of the pre-compiled core without an Enterprise License or total IP Acquisition is strictly prohibited.
🚀 How to Run and Verify Performance
You can natively compile the project and audit the benchmarking claims directly on your local infrastructure.
Ensure you have the Rust toolchain, Python 3.12 development headers, and the native linker installed on your machine:
sudo apt update
sudo apt install build-essential python3-dev python3-config lld
2. Verify Local Benchmarks (Criterion)
The micro-benchmarking suite is isolated within the core source files to prevent Python runtime context symbol collisions. To run the statistical hardware latency reports, execute:
# Clear any stale linker metadata
cargo clean
# Run the target context manager benchmark suite
cargo bench --bench context_manager_bench
The detailed statistical distribution curves will be generated under target/criterion/report/index.html.
3. Test the Python Extension Module
Build the native extension into your local Python environment and execute the integration pipeline test:
# Activate your local virtual environment
source .venv/bin/activate
# Install the compilation wrapper
pip install maturin
# Compile the project using the pre-compiled high-performance core
maturin develop --release
# Run the live agent context swapping simulation
python3 test_engine.py
💼 Corporate Development & M&A (Mergers and Acquisitions)
This technology is architected by senior systems engineers and is optimized for direct cloud infrastructure integration (e.g., AWS Bedrock, Google Vertex AI, Meta GenAI cluster nodes).
To protect deep engineering focus, we operate a strictly asynchronous, documentation-first M&A pipeline. We do not participate in introductory discovery calls, introductory Zoom sessions, or loose technical alignment meetings.
The software speaks for itself. If your engineering team has executed cargo bench, audited the micro-benchmarks, and verified that ICPE solves your VRAM/latency bottleneck, the acquisition workflow is as follows:
Submission: Your Corporate Development / M&A team submits a formal, written Letter of Intent (LOI) or a firm acquisition proposal from an official corporate domain directly to our secure channel.
Due Diligence: Upon receiving a valid LOI, we will instantly grant access to a secure data room containing clean-room source code, mathematical specifications, comprehensive fuzzing reports, and automated hardware stress tests under a standard NDA.
Closing: Fully automated IP transfer and legal closing.
📩 Official Channel for LOIs & Firm Proposals: corpdev@matheusdelgado.com
An ultra-low latency, zero-copy context virtual memory paging engine written in Rust, designed to break physical VRAM limitations for LLMs and autonomous agents using attention-driven predictive prefetching.
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
