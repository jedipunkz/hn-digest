---
source: "https://github.com/projectargus-cc/libargus.cc"
hn_url: "https://news.ycombinator.com/item?id=48907681"
title: "Show HN: Low-latency local LLM runner via OpenJDK Panama FFM (Java 22)"
article_title: "GitHub - ProjectArgus-cc/libargus.cc: A unified, unmanaged multimodal runtime. Zero-allocation native AI inference for Java 22+. · GitHub"
author: "KingJoker"
captured_at: "2026-07-14T15:17:32Z"
capture_tool: "hn-digest"
hn_id: 48907681
score: 4
comments: 1
posted_at: "2026-07-14T14:40:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Low-latency local LLM runner via OpenJDK Panama FFM (Java 22)

- HN: [48907681](https://news.ycombinator.com/item?id=48907681)
- Source: [github.com](https://github.com/projectargus-cc/libargus.cc)
- Score: 4
- Comments: 1
- Posted: 2026-07-14T14:40:46Z

## Translation

タイトル: Show HN: OpenJDK Panama FFM 経由の低遅延ローカル LLM ランナー (Java 22)
記事のタイトル: GitHub - ProjectArgus-cc/libargus.cc: 統合されたアンマネージド マルチモーダル ランタイム。 Java 22 以降のゼロ割り当てネイティブ AI 推論。 · GitHub
説明: 統合された、管理されていないマルチモーダル ランタイム。 Java 22 以降のゼロ割り当てネイティブ AI 推論。 - ProjectArgus-cc/libargus.cc
HN テキスト: JVM 内から AI を実行したかったのです。私は標準の REST サイドカーから始めて、新しい JDK バージョンの Project Panama (Foreign Function & Memory API) を使用して llama.cpp と直接インターフェイスするためにそれを取り除きました。それでもその機能に満足できなかったので、JVM ランドスケープで構造化 API を公開するためのクリーンな ABI を取得するために libargus.cc を構築しました。これは依然として Project Panama を使用して、llama.cpp、whisper.cpp、および ggml 計算グラフと直接インターフェイスします。ホット パスにはゼロ割り当てがあり、プロンプトとトークンのメモリ セグメントは、制限されたアリーナ内で一度割り当てられます。生のポインタは、低 C レベルまで直接通過します。これにより、プリミティブ配列のクローン作成とヒープ チャーンが回避されます。安全なメモリ アクセスを維持するためにコンパイラのパディングを一致させながら、llama.cpp と Whisper.cpp からネイティブ構造をマッピングしました。導入を容易にするために、コンパイル済みのネイティブ バイナリを jar にバンドルします。この実行エンジンは、RAG を置き換えるために時空間メモリ層 (L-TABB) で行っている作業に必要な基盤を提供します。次のレイヤーの作業を続ける間、問題を改善するために技術的なフィードバックを得たいと考えています。
Project Panama や最新の JDK の低レイテンシ システムをハッキングしている方からの詳細な情報をお待ちしております。私は散文よりもコードの方がはるかに得意なので、ほとんどのことをコードに任せます。ハッピーハッキング!
/デビッドコード: https://libargus.cc
プロジェクトのランディング ページ: https://projectargus.c

c

記事本文:
GitHub - ProjectArgus-cc/libargus.cc: 統合されたアンマネージド マルチモーダル ランタイム。 Java 22 以降のゼロ割り当てネイティブ AI 推論。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
プロジェクトArgus-cc
/
libargus.cc
出版

c
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット bindings/ java bindings/ java include include src src testing testing .geminiignore .geminiignore .gitattributes .gitattributes .gitignore .gitignore CMakeLists.txt CMakeLists.txt ライセンス ライセンス README.md README.md build.gradle.kts build.gradle.kts download_models.sh download_models.sh settings.gradle.kts settings.gradle.kts version.txt version.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ビジョン、スピーチ、LLM コンピューティング パイプラインを単一の Project Panama FFM 境界の背後に統合する、管理されていない割り当てゼロのネイティブ AI 実行ランタイム。
v1.0.0 安定版 — 統合ハードウェア オーケストレーションおよびゼロ割り当て ABI
libargus は、LLM テキスト生成、ウィスパーベースの音声テキスト変換 (ASR)、Speech-LLM テキスト音声変換 (TTS)、および最先端のマルチモーダル (ビジョン、オーディオ、ビデオ) エンコードと評価パイプラインを単一のプロセス グローバル ネイティブ実行ランタイムに統合するように設計された、超無駄のない高性能のモデルに依存しない推論ラッパーです。
libargus は、モジュラー GGML および llama.cpp (libmtmd) コンピューティング エンジン上に直接構築されており、最新のアンマネージド オーケストレーション フレームワークと並行して、摩擦のないゼロコピー コンパイル用に明示的に設計された、統合されたスレッドセーフな C API を提供します。これは、JDK 22+ プロジェクト Panama 外部関数 & メモリ (FFM) API のすぐに使用できる構造調整を特徴としています。
プロセス グローバル バックエンド特異点: テキスト、オーディオ、音声、およびマルチモーダル サブシステムにわたる単一の共有初期化経路 ( ggml_backend_load_all() ) を調整することにより、VRAM の断片化とマルチコンテキスト ドライバーの競合状態を排除します。
分離された重みと実行:

モデルの重みの読み込み ( argus_model_t ) を評価コンテキスト メモリの状態 ( argus_context_t ) から分離し、複数の同時セッション間でモデルを再利用できるようにします。
最先端のマルチモーダル プロジェクター: 新しい libmtmd C++ エンジンを統合して、生のビットマップ、オーディオ PCM 配列、およびビデオ ファイル/ストリームを取り込みます。プロンプトとメディアを統一されたチャンク シーケンスにトークン化し、GPU 上で投影を実行し、M-RoPE 位置グリッドと非因果的アテンション マトリックスを自動的に構成します。
アンマネージド ビデオ イテレーション パイプ: 内部 FFmpeg サブプロセス パイプを使用してビデオ ファイルをフレームごとにデコードしてストリーミングし、指定されたターゲット フレーム レートで生の RGB フレームまたはローカライズされたタイムスタンプ テキスト チャンク (例: [12m34s] ) を生成します。
ポインタのみの FFM アライメント: 値渡しおよび揮発性の C++ ポリモーフィック境界を、ポインタを受け入れる厳密にアライメントされたフラットな C 関数に置き換えます。構造体のパディングは、コンパイラーがアライメントギャップを挿入しないように手動でパックされます。
絶対ゼロコピーのメモリ境界: ホット パス全体で JVM ヒープのプリミティブ配列 ( int[] 、 float[] ) を排除します。 Project Panama MemorySegment パラメーターを直接統合し、トークン テープ、音声波形、およびビデオ フレームで GC フットプリントをゼロにしながら音声とテキストを生成できるようにします。
選択的同時実行ロック: コンテキスト レベルのミューテックス同期を統合して、読み取り専用モデルで完全にロックフリーの同時トークナイザー アクセスを有効にしながら、スレッド セーフなデコードとコンテキスト操作を可能にします。
投機的および MTP の高速化: 従来の投機的ドラフティングとマルチトークン予測 (draft-mtp) 用のネイティブ検証ループを C++ 実行層内に直接組み込みます。
KV キャッシュ量子化: ネイティブ構成 (type_k および type_v キャッシュ列挙型) をサポートし、メモリ フットプリントを Q8_0、Q4_0、またはその他の最適化された形式にオフロードします。
ザー

o-Allocation Vocab & GGUF Metadata Introspection: 特別な語彙トークン (BOS、EOS、EOT、PAD) を検索し、世代終了 (EOG) 条件を検証し、GGUF 辞書エントリを動的に列挙するために、安全で管理されていない境界を公開します。
リバルガス/
§── CMakeLists.txt # レイヤー 0 の依存関係の分離と最適化マトリックス
§── 含む/
│ └─ libargus.h # マスター C ABI 安定レイアウト定義
§── src/
│ §── argus_internal.h # 共有プライベート構造体 (モデルとコンテキスト)
│ §── argus_common.cc # グローバル バックエンド ライフサイクルとハードウェア レジストリ
│ §── argus_text.cc # Llama モデル/コンテキスト処理、投機的ループ & TTS
│ §── argus_audio.cc # ウィスパーモデルのコンテキストと転写 (ASR)
│ └── argus_multimodal.cc# マルチモーダル コンテキスト、メディア ローダー、ビデオ パイプ、および評価
└── bindings/java/ # Idiomatic Project Panama FFM バインディング モジュール
└── src/main/java/cc/projectargus/libargus/
§── ArgusBackend.java # グローバル デバイス テレメトリとバックエンドの初期化
§── ArgusModel.java # アンマネージ GGUF 重みマネージャー (AutoCloseable)
§── ArgusContext.java # コアテキスト評価コンテキストセッション
§── ArgusContextConfig.java # テキストコンテキスト生成パラメータ
§── ArgusAudioContext.java # Whisper 音声テキスト変換エンジン
§── ArgusMultimodalContext.java# ロードされたマルチモーダル プロジェクター コンテキスト (AutoCloseable)
§── ArgusBitmap.java # 生/解析済みの RGB ピクセルまたは PCM オーディオ サンプル バッファー
§── ArgusVideo.java # ビデオ ファイルまたはバッファ パイプのフレーム イテレータ
§── ArgusVideoItem.java # ビデオ処理用の再利用可能なフレーム/タイムスタンプ コンテナ
§── ArgusInputChunks.java # トークン化されたマルチモーダル プロンプト チャンク コンテナ
━── 内部/
§── ArgusLayouts.java # パナマ C-to-J

ava 構造体レイアウト定義
━── ArgusBindings.java # 動的共有ライブラリメソッドハンドルローダー
ビルドと依存関係のアーキテクチャ
libargus は、元のソース構成を強制します。 CMake レベルのビルド層ベンダー ( FetchContent ) を利用することで、肥大化した上流サーバー実装、レガシー CLI ターゲット、不要な依存関係を破棄します。アップストリーム コンポーネントは、アンマネージド コンパイル パス内でダウンロード、構成され、静的にリンクされます。
高度に最適化された共有バイナリ ターゲット ( libargus.so または argus.dll ) をコンパイルするには、ルート ディレクトリからターゲット生成コマンドを実行します。
# 最適化されたアンマネージド コンピューティング グラフ プロジェクトを生成します (CUDA アクセラレーションを有効にします)
cmake -B build -DCMAKE_BUILD_TYPE=リリース -DGGML_CUDA=ON
# 最終的な統合システム バイナリをコンパイルする
cmake --build build --config Release -j $( nproc )
クイックスタート: 慣用的な Java 開発者エクスペリエンス
libargus は、JVM 開発者を複雑なポインター演算、構造体の位置合わせのギャップ、および手動のアテンション マスク スケジューリングから保護します。以下は、高レベルでメモリセーフで自動クローズ可能な Java API の使用パターンです。
テキストと音声の文字起こし (ASR)
輸入CC。プロジェクターガス 。リバルガス .*;
java をインポートします。ラング 。外国 。アリーナ ;
java をインポートします。にを。ファイル 。パス ;
パブリッククラス Main {
public static void main ( String [] args ) {
// グローバル バックエンド (CUDA/CPU) を初期化します。
アーガスバックエンド 。初期化();
try ( アリーナ アリーナ = アリーナ .ofConfined ();
ArgusModel モデル = ArgusModel 。 load ( arena , Path . of ( "models/llama-3-8b.gguf" ), 99 , true )) {
// コンテキスト設定を初期化します
ArgusContextConfig config = 新しい ArgusContextConfig 。ビルダー ( 4096 )
。 CPUスレッド ( 8 )
。 typeK ( ArgusContextConfig . KV_TYPE_Q4_0 ) // KV キャッシュを量子化します
。 typeV ( ArgusContextConfig . KV_TYPE_Q4_0 )
。建てる （）;
try ( ArgusContext コンテキスト

= ArgusContext 。 init (アリーナ、モデル、構成)) {
// テキストの評価と生成ループを実行します...
}
最後に {
// ネイティブ バックエンド ドライバーを破棄します
アーガスバックエンド 。無料 （）;
}
}
}
最先端のマルチモーダル プロンプト (ビジョン/ビデオ/オーディオ)
ビジョン対応 GGUF モデルとそのマルチモーダル プロジェクター ( mmproj ) を使用します。
輸入CC。プロジェクターガス 。リバルガス .*;
java をインポートします。ラング 。外国 。アリーナ ;
java をインポートします。にを。ファイル 。パス ;
java をインポートします。ユーティリティ。リスト ;
パブリック クラス MultimodalApp {
public static void main ( String [] args ) {
アーガスバックエンド 。初期化();
try ( アリーナ アリーナ = アリーナ .ofConfined ();
ArgusModel BaseModel = ArgusModel 。 load ( arena , Path . of ( "models/qwen2-vl-7b-it.gguf" ), 99 , true );
ArgusContext コンテキスト = ArgusContext 。 init ( arena 、 baseModel 、 new ArgusContextConfig . Builder ( 8192 ). build ());
// マルチモーダルアダプターコンテキストをロードします
ArgusMultimodalContext mctx = ArgusMultimodalContext 。 init ( arena , baseModel , Path . of ( "models/qwen2-vl-7b-it.mmproj" ), 4 , true )) {
// 1. 画像または音声ファイルをアンマネージ メモリにロードします
try ( ArgusBitmap image = ArgusBitmap .loadFile ( arena , mctx , Path . of ( "media/cat.png" ), false )) {
// 2. メディア マーカーを置き換えるプロンプト テキストをトークン化します。
文字列プロンプト = "<__media__> \n この画像に何が表示されているか説明してください。" ;
try ( ArgusInputChunks チャンク = mctx . tokenize ( arena , prompt , true , List . of ( image ))) {
// 3. チャンクを評価します (GPU および M-RoPE 位置グリッドでの画像投影を処理します)
int newNPast = コンテキスト . evalMultimodalChunks ( mctx 、チャンク、 0 、 0 、 1024 、 true );
システム 。外 。 println ( "プロンプトが評価されました。出力トークンをサンプリングする準備ができました! 新しい位置: " + newNPast );
}
}
最後に {
アーガスバックエンド 。無料 （）;
}
}
}
フレームごとのビデオ ストリーム処理
// ビデオからフレーム/タイムスタンプを反復して読み取ります

順番にファイルする
try ( ArgusVideo video = ArgusVideo .loadFile ( arena , mctx , Path . of ( "media/video.mp4" ), 4.0f , 5000 );
ArgusVideoItem item = new ArgusVideoItem ()) {
while ( video . readNext ( item )) {
if ( item . bitmap () != null ) {
// 抽出された RGB フレーム ビットマップを処理します (所有権は ArgusVideoItem によって管理されます)
ArgusBitmap フレーム = item 。ビットマップ ();
// ...
else if ( item . text () != null ) {
// ビデオのタイムスタンプ チャンクを受信しました (例: "[00m05s]")
システム。外 。 println ( "タイムスタンプ時: " + item . text ());
}
}
}
セマンティックテキスト埋め込みの抽出
専用モデル (例: jina-embeddings-v3 ) から float 埋め込みベクトルを取得します。
輸入CC。プロジェクターガス 。リバルガス .*;
java をインポートします。ラング 。外国 。アリーナ ;
java をインポートします。ラング 。外国 。メモリセグメント ;
java をインポートします。ラング 。外国 。値レイアウト ;
java をインポートします。にを。ファイル 。パス ;
パブリック クラス EmbeddingsApp {
public static void main ( String [] args ) {
アーガスバックエンド 。初期化();
try ( アリーナ アリーナ = アリーナ .ofConfined ();
ArgusModel モデル = ArgusModel 。 load ( arena , Path . of ( "models/jina-embeddings-v3-Q4_K_M.gguf" ), 99 , true )) {
// 埋め込みを有効にしてコンテキスト設定を初期化します
ArgusContextConfig config = 新しい ArgusContextConfig 。ビルダー ( 512 )
。 CPUスレッド ( 4 )
。埋め込み ( true )
。建てる （）;
try ( ArgusContext context = ArgusContext .init ( arena , model , conf

[切り捨てられた]

## Original Extract

A unified, unmanaged multimodal runtime. Zero-allocation native AI inference for Java 22+. - ProjectArgus-cc/libargus.cc

I wanted to run AI from inside the JVM. I started out with the standard REST sidecar, ripped that out to use Project Panama (Foreign Function & Memory API) in the new JDK versions to interface directly with llama.cpp. I still wasn't happy with how that functioned, so I built libargus.cc to get a clean ABI to expose a structured API up in the JVM landscape. It still uses Project Panama to interface directly with llama.cpp, whisper.cpp, and ggml compute graphs. I have zero-allocation on the hot paths, memory segments for prompts and tokens are allocated once inside confined Arenas. Raw pointers pass straight through down to the low C level. This avoids primitive array cloning and heap churn. I mapped out the native structures from llama.cpp and whisper.cpp while matching the compiler's padding to maintain safe memory access. I bundle pre-compiled native binaries in the jar for easy deployment. This execution engine provides the foundation I need for work I'm doing on a spatio-temporal memory layer (L-TABB) to replace RAGs. I'd love to get technical feedback to polish any issues while I continue working on the next layer.
Deep-dives from anyone hacking on Project Panama or low-latency systems in modern JDK would be very appreciated! I'm much better with code than prose, so I'll let the code do most of the talking. Happy Hacking!
/David Code: https://libargus.cc
Project Landing Page: https://projectargus.cc

GitHub - ProjectArgus-cc/libargus.cc: A unified, unmanaged multimodal runtime. Zero-allocation native AI inference for Java 22+. · GitHub
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
ProjectArgus-cc
/
libargus.cc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits bindings/ java bindings/ java include include src src tests tests .geminiignore .geminiignore .gitattributes .gitattributes .gitignore .gitignore CMakeLists.txt CMakeLists.txt LICENSE LICENSE README.md README.md build.gradle.kts build.gradle.kts download_models.sh download_models.sh settings.gradle.kts settings.gradle.kts version.txt version.txt View all files Repository files navigation
An unmanaged, zero-allocation native AI execution runtime consolidating Vision, Speech, and LLM compute pipelines behind a single Project Panama FFM boundary.
v1.0.0 Stable — Unified Hardware Orchestration & Zero-Allocation ABI
libargus is an ultra-lean, high-performance, model-agnostic inference wrapper engineered to consolidate LLM text generation, Whisper-based speech-to-text (ASR), Speech-LLM text-to-speech (TTS), and bleeding-edge Multimodal (Vision, Audio, and Video) encoding and evaluation pipelines into a single process-global native execution runtime.
Built directly on top of the modular GGML and llama.cpp (libmtmd) compute engines, libargus provides a unified, thread-safe C API designed explicitly for frictionless, zero-copy compilation alongside modern unmanaged orchestration frameworks, featuring out-of-the-box structural alignment for the JDK 22+ Project Panama Foreign Function & Memory (FFM) API .
Process-Global Backend Singularity: Eliminates VRAM fragmentation and multi-context driver race conditions by orchestrating a singular, shared initialization pathway ( ggml_backend_load_all() ) across text, audio, speech, and multimodal subsystems.
Decoupled Weights & Execution: Separates model weight loading ( argus_model_t ) from evaluation context memory states ( argus_context_t ), allowing model reuse across multiple concurrent sessions.
Bleeding-Edge Multimodal Projectors: Integrates the new libmtmd C++ engine to ingest raw bitmaps, audio PCM arrays, and video files/streams. Tokenizes prompts and media into a unified chunk sequence, executes projection on the GPU, and automatically configures M-RoPE position grids and non-causal attention matrices.
Unmanaged Video Iteration Pipe: Decodes and streams video files frame-by-frame using internal FFmpeg subprocess pipes, yielding raw RGB frames or localized timestamp text chunks (e.g., [12m34s] ) at a specified target frame rate.
Pointers-Only FFM Alignment: Replaces pass-by-value and volatile C++ polymorphic boundaries with strictly aligned, flat C functions accepting pointers. Structure padding is manually packed to prevent compilers from injecting alignment gaps.
Absolute Zero-Copy Memory Boundaries: Eliminates JVM heap primitive arrays ( int[] , float[] ) across hot paths. Integrates Project Panama MemorySegment parameters directly, allowing token tapes, audio waves, and video frames to generate speech and text with zero GC footprint.
Selective Concurrency Locking: Integrates context-level mutex synchronization to allow thread-safe decoding and context operations while enabling fully lock-free, concurrent tokenizer accesses on read-only models.
Speculative & MTP Acceleration: Incorporates native verification loops for traditional speculative drafting and Multi-Token Prediction ( draft-mtp ) directly inside the C++ execution layer.
KV Cache Quantization: Supports native configurations ( type_k and type_v cache enums) to offload memory footprints to Q8_0, Q4_0, or other optimized formats.
Zero-Allocation Vocab & GGUF Metadata Introspection: Exposes safe, unmanaged boundaries to lookup special vocab tokens (BOS, EOS, EOT, PAD), verify End-Of-Generation (EOG) conditions, and dynamically enumerate GGUF dictionary entries.
libargus/
├── CMakeLists.txt # Layer-0 dependency isolation & optimization matrix
├── include/
│ └── libargus.h # Master C ABI stable layout definitions
├── src/
│ ├── argus_internal.h # Shared private structures (model & context)
│ ├── argus_common.cc # Global backend lifecycles & hardware registries
│ ├── argus_text.cc # Llama model/context handling, speculative loops & TTS
│ ├── argus_audio.cc # Whisper model contexts & transcription (ASR)
│ └── argus_multimodal.cc# Multimodal context, media loaders, video pipes, and evaluation
└── bindings/java/ # Idiomatic Project Panama FFM binding module
└── src/main/java/cc/projectargus/libargus/
├── ArgusBackend.java # Global device telemetry & backend initialization
├── ArgusModel.java # Unmanaged GGUF weights manager (AutoCloseable)
├── ArgusContext.java # Core text evaluation context session
├── ArgusContextConfig.java # Text context generation parameters
├── ArgusAudioContext.java # Whisper speech-to-text transcription engine
├── ArgusMultimodalContext.java# Loaded multimodal projector context (AutoCloseable)
├── ArgusBitmap.java # Raw/parsed RGB pixel or PCM audio sample buffer
├── ArgusVideo.java # Frame iterator for video files or buffer pipes
├── ArgusVideoItem.java # Reusable frame/timestamp container for video processing
├── ArgusInputChunks.java # Tokenized multimodal prompt chunks container
└── internal/
├── ArgusLayouts.java # Panama C-to-Java struct layout definitions
└── ArgusBindings.java # Dynamic shared library method handle loader
Build & Dependency Architecture
libargus enforces a pristine source configuration. It discards bloated upstream server implementations, legacy CLI targets, and unneeded dependencies by utilizing CMake-level build-layer vendoring ( FetchContent ). Upstream components are downloaded, configured, and statically linked inside the unmanaged compilation pass.
To compile the highly optimized shared binary target ( libargus.so or argus.dll ), execute the target generation commands from the root directory:
# Generate the optimized unmanaged compute graph project (enabling CUDA acceleration)
cmake -B build -DCMAKE_BUILD_TYPE=Release -DGGML_CUDA=ON
# Compile the final unified system binary
cmake --build build --config Release -j $( nproc )
Quickstart: Idiomatic Java Developer Experience
libargus shields JVM developers from complex pointer arithmetic, structure alignment gaps, and manual attention mask scheduling. Below is the high-level, memory-safe, and auto-closeable Java API usage pattern.
Text & Audio Transcription (ASR)
import cc . projectargus . libargus .*;
import java . lang . foreign . Arena ;
import java . nio . file . Path ;
public class Main {
public static void main ( String [] args ) {
// Initialize global backends (CUDA/CPU)
ArgusBackend . init ();
try ( Arena arena = Arena . ofConfined ();
ArgusModel model = ArgusModel . load ( arena , Path . of ( "models/llama-3-8b.gguf" ), 99 , true )) {
// Initialize context configurations
ArgusContextConfig config = new ArgusContextConfig . Builder ( 4096 )
. cpuThreads ( 8 )
. typeK ( ArgusContextConfig . KV_TYPE_Q4_0 ) // Quantize KV Cache
. typeV ( ArgusContextConfig . KV_TYPE_Q4_0 )
. build ();
try ( ArgusContext context = ArgusContext . init ( arena , model , config )) {
// Run text evaluation & generation loop...
}
} finally {
// Tear down native backend drivers
ArgusBackend . free ();
}
}
}
Bleeding-Edge Multimodal Prompting (Vision/Video/Audio)
Using a vision-capable GGUF model along with its multimodal projector ( mmproj ):
import cc . projectargus . libargus .*;
import java . lang . foreign . Arena ;
import java . nio . file . Path ;
import java . util . List ;
public class MultimodalApp {
public static void main ( String [] args ) {
ArgusBackend . init ();
try ( Arena arena = Arena . ofConfined ();
ArgusModel baseModel = ArgusModel . load ( arena , Path . of ( "models/qwen2-vl-7b-it.gguf" ), 99 , true );
ArgusContext context = ArgusContext . init ( arena , baseModel , new ArgusContextConfig . Builder ( 8192 ). build ());
// Load the multimodal adapter context
ArgusMultimodalContext mctx = ArgusMultimodalContext . init ( arena , baseModel , Path . of ( "models/qwen2-vl-7b-it.mmproj" ), 4 , true )) {
// 1. Load an image or audio file into unmanaged memory
try ( ArgusBitmap image = ArgusBitmap . loadFile ( arena , mctx , Path . of ( "media/cat.png" ), false )) {
// 2. Tokenize prompt text replacing the media marker
String prompt = "<__media__> \n Describe what you see in this image." ;
try ( ArgusInputChunks chunks = mctx . tokenize ( arena , prompt , true , List . of ( image ))) {
// 3. Evaluate chunks (handles image projection on GPU & M-RoPE position grids)
int newNPast = context . evalMultimodalChunks ( mctx , chunks , 0 , 0 , 1024 , true );
System . out . println ( "Prompt evaluated. Ready to sample output tokens! New position: " + newNPast );
}
}
} finally {
ArgusBackend . free ();
}
}
}
Frame-by-Frame Video Stream Processing
// Iterate and read frames/timestamps from a video file sequentially
try ( ArgusVideo video = ArgusVideo . loadFile ( arena , mctx , Path . of ( "media/video.mp4" ), 4.0f , 5000 );
ArgusVideoItem item = new ArgusVideoItem ()) {
while ( video . readNext ( item )) {
if ( item . bitmap () != null ) {
// Process the extracted RGB frame bitmap (ownership is managed by ArgusVideoItem)
ArgusBitmap frame = item . bitmap ();
// ...
} else if ( item . text () != null ) {
// Received a video timestamp chunk (e.g. "[00m05s]")
System . out . println ( "At timestamp: " + item . text ());
}
}
}
Extracting Semantic Text Embeddings
Retrieve float embedding vectors from dedicated models (e.g., jina-embeddings-v3 ):
import cc . projectargus . libargus .*;
import java . lang . foreign . Arena ;
import java . lang . foreign . MemorySegment ;
import java . lang . foreign . ValueLayout ;
import java . nio . file . Path ;
public class EmbeddingsApp {
public static void main ( String [] args ) {
ArgusBackend . init ();
try ( Arena arena = Arena . ofConfined ();
ArgusModel model = ArgusModel . load ( arena , Path . of ( "models/jina-embeddings-v3-Q4_K_M.gguf" ), 99 , true )) {
// Initialize context configuration with embeddings enabled
ArgusContextConfig config = new ArgusContextConfig . Builder ( 512 )
. cpuThreads ( 4 )
. embeddings ( true )
. build ();
try ( ArgusContext context = ArgusContext . init ( arena , model , conf

[truncated]
