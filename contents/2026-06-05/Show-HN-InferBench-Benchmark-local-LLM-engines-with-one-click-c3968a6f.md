---
source: "https://github.com/JoniMartin27/inferbench"
hn_url: "https://news.ycombinator.com/item?id=48413247"
title: "Show HN: InferBench – Benchmark local LLM engines with one click"
article_title: "GitHub - JoniMartin27/inferbench: Panel local multiplataforma para benchmark de motores de inferencia LLM (llama.cpp nativo + APIs cloud) · GitHub"
author: "JoniMartin"
captured_at: "2026-06-05T16:07:30Z"
capture_tool: "hn-digest"
hn_id: 48413247
score: 1
comments: 0
posted_at: "2026-06-05T14:44:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: InferBench – Benchmark local LLM engines with one click

- HN: [48413247](https://news.ycombinator.com/item?id=48413247)
- Source: [github.com](https://github.com/JoniMartin27/inferbench)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T14:44:08Z

## Translation

タイトル: HN の表示: InferBench – ワンクリックでローカル LLM エンジンのベンチマークを実行
記事タイトル: GitHub - JoniMartin27/inferbench: LLM 推論エンジンのベンチマーク用マルチプラットフォーム ローカル パネル (ネイティブ llama.cpp + クラウド API) · GitHub
説明: LLM 推論エンジン ベンチマーク用のマルチプラットフォーム ローカル パネル (ネイティブ llama.cpp + クラウド API) - JoniMartin27/inferbench

記事本文:
GitHub - JoniMartin27/inferbench: パネル ローカル マルチプラットフォーム パラ ベンチマーク デ モーター デ 推論 LLM (llama.cpp nativo + API クラウド) · GitHub
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
ジョニマーティン27
/
インファーベンチ
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
73 コミット 73 コミット .claude/ skill/ run-inferbench-backend .claude/ skill/ run-inferbench-backend .github .github アセット アセット バックエンド バックエンド docker docker フロントエンド フロントエンド スクリプト スクリプト Web サイト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LAUNCH.md LAUNCH.md ライセンス ライセンス PROJECT_BRIEF.md PROJECT_BRIEF.md README.md README.md SECURITY-AUDIT.md SECURITY-AUDIT.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ワンクリックでローカル LLM 推論エンジンをダウンロード、起動、ベンチマークします。
Docker は必要ありません。コマンドラインには触れずに。データがマシンから離れることはありません。
LLM をローカルで実行したいのですが、GPU にどのような量子化が含まれているか、どのくらいの tok/s で動作するか、またはどのエンジンがハードウェアにとって高速であるかがわかりません。 InferBench は、でっち上げられた数値ではなく、実際に測定して、あなたの代わりに答えます。
モデル + 量子化 → InferBench を選択します。
①エンジンバイナリ（公式GitHubリリース）をダウンロードします。
②Hugging Faceから足りないGGUFをダウンロードする
③ ハードウェアに最適な設定でエンジンを起動します。
④ プロンプトスイートを実行して、TTFT、tok/s、VRAM、品質を測定します
⑤ 結果を保存し、並べて比較できます
実際の例
RTX 3070 · 8 GB VRAM · 32 GB RAM — ネイティブ llama.cpp で SmolLM2-360M Q8_0 を実行:
数値は実稼働環境から取得したものであり、作成されたデータではありません。
リリース ページからシステムのインストーラーを取得します。Python や Node をインストールする必要はありません。バックエンドには次のものが埋め込まれています。

またはサイドカー:
自分でコンパイルした方が良いですか? 「インストール (開発)」に進みます。
エンドツーエンドの自動ブートストラップ: 1 クリック → バイナリ + モデル + ブート + ベンチマーク + クリーンアップ
llama.cpp のネイティブ モード (Docker なし): GitHub からコンパイル済みリリースをダウンロード (CUDA を自動検出し、ランタイム DLL もダウンロードします)
Docker モードを必要とするすべてのエンジンで利用可能
自動ハードウェア検出: CPU、RAM、GPU (NVML 経由の NVIDIA、rocm-smi 経由の AMD、system_profiler 経由の Apple Silicon)。互換性リストを即時にキャッシュする (124 モデルで最大 4 ミリ秒)
HF から GGUF を自己ダウンロードした 124 モデルのカタログ。すべて HuggingFace に対して検証済み (ビジョン、コード、推論、および MoE を含む)。カタログを見る
ローカル GGUF のスキャン: GGUF メタデータから読み取られた実際のパラメータ数を使用して、LM Studio、Ollama、HF キャッシュなどからモデルを検出します (量子に依存しない)
リアル ビジョン (マルチモーダル): ビジョン モデル (Qwen2-VL、Qwen2.5-VL、MiniCPM-V) の場合は、mmproj をダウンロードし、--mmproj でコール サーバーを起動し、OpenAI 互換のビジョン API を介して実際の画像を使用してプロンプトをベンチマークします。また、API (gpt-4o、claude) と Docker エンジン (vLLM はビジョン モデルを提供します) でも、ビジョン モデルを備えた vLLM には大量の VRAM が必要です (2B は 8 GB に快適に収まりません。ローカルでは、llama.cpp が少量の VRAM への信頼できるルートです)。
オプティマイザー: ハードウェア + モデル + エンジンを考慮して、最適な量子化、KV キャッシュ、最大コンテキスト、MoE オフロード、フラグを計算します。
KV キャッシュ圧縮の説明: 5 つのプリセット (品質→極度) とその機能 / 影響 / 許可内容、およびハードウェアの各圧縮に適合する最も強力なモデルの表
3 つのモードでの品質評価: ベンチマークベースのオフライン スコアラー (GPU/API なし、任意の PC で実行)、LLM 判定

ローカルエンジンによる、または外部APIによるLLM判定。品質を見る
スイープ モード: N 個の異なる量子化をキューに入れて同じモデルを起動します。
比較: 履歴から複数の実行を選択し、メトリクスとグラフを並べて表示します。
ライブ SSE: ダウンロードの進行状況 (% 付き)、TTFT、現在の tok/s、ターミナル スタイル ログ
回復力のあるダウンロード: 自動再試行 (指数関数的バックオフ) と、ネットワークがダウンした場合の Range ヘッダー経由の部分的な再開を備えた数十 GB の GGUF
OS キーリングの API キー: OpenAI/Anthropic/OpenRouter/NVIDIA キーをシステム資格情報マネージャーに保存します ([設定] → [API キー])。ベンチマークはそれらを単独で使用し、ディスクに書き込まれることはありません
Lookspan による可観測性 (オプトイン): LOOKSPAN_ENDPOINT を定義すると、各ベンチマーク実行がトレースとして Lookspan (ルート スパン + 実際の TTFT、tok/s、VRAM、品質を含むプロンプトごとの llm_call) としてエクスポートされます。ベストエフォート: ネットワーク障害がベンチマークを破ることはありません
いつでも停止: ブートストラップ、ダウンロード、または実行をキャンセルします。
永続性: すべての実行を含む SQLite (エンジン、モデル、クォント、フラグ、プロンプト メトリクス、生の出力)
マント
テクノロジー
デスクトップアプリ
電子 42
フロントエンド
React 18 + Vite 8 + TailwindCSS + Recharts + lucide-react
バックエンド
FastAPI (Python 3.11) + uvicorn + sse-starlette
持続性
SQLModel経由のSQLite
GPUの検出
psutil + pynvml + system_profiler / rocm-smi
コンテナ
Python 用 Docker SDK
ネイティブランタイム
subprocess.Popen + 公式 GitHub バイナリ
包装
電子ビルダー + PyInstaller
対応モーター
エンジン
タイプ
ネイティブモード
ドッカーモード
自己放電モデル
callcpp
地元の
✅ 公式バイナリ
✅
✅ ハギングフェイス GGUF
オラマ
地元の
✅ オラマデーモン
✅
✅ オラマ登録
vllm
地元の
—
✅ (NVIDIA GPU)
✅ HF (容器内)
俗語
地元の
—
✅ (NVID GPU

IA)
✅ HF (容器内)
tgi
地元の
—
✅ (NVIDIA GPU)
✅ HF (容器内)
オープンナイ
API
該当なし
該当なし
該当なし
人間的な
API
該当なし
該当なし
該当なし
オープンルーター
API
該当なし
該当なし
該当なし
エヌビディア
API
該当なし
該当なし
該当なし
すべてのローカル エンジンには完全なアダプター (エンジンごとのコマンド ビルド、自動ブートストラップ、独自の最適化スキーマ) が備わっています。 vLLM/SGLang/TGI は Docker のみであり、NVIDIA GPU を必要とします。モデルはコンテナ自体によって HuggingFace からダウンロードされます (リポジトリ ID を渡します)。クラウド API は API キーで動作します (サンプリング パラメーターのみ、ローカル最適化なし): OpenAI / OpenRouter / NVIDIA は OpenAI 互換エンドポイント /v1/chat/completions を使用します。 Anthropic はネイティブ API ( /v1/messages 、ヘッダー x-api-key + anthropic-version 、システムは別として) を使用しますが、OpenAI 互換ではありません。
投機的デコード (DFLASH): vLLM および SGLang は、ドラフト ブロック拡散モデル (DFLASH、品質損失なしで 6 ～ 8 倍) による高速化を受け入れます。ベンチマークで、vLLM/SGLang を選択した状態で DFLASH をアクティブにし、ドラフト モデル (例: z-lab/Qwen3.5-35B-A3B-DFlash ) を指定します。 SGLang が公式ルートです。 vLLM にはサポートされているビルドが必要です。モデルとドラフトには VRAM が必要なため、小型の GPU には適合しません。
検証ステータス: 5 つのローカル エンジン ( llamacpp 、 rocama 、 vllm 、 sglang 、 tgi ) はすべて、NVIDIA GPU (RTX 3070、8 GB) 上のプロダクション ランナー (ブートストラップ → ブート → tps>0 での実推論 → コンテナをハングさせずに停止) によってエンドツーエンドで検証されています。 vLLM/SGLang は、空でない GPU でのクラッシュを避けるために、VRAM の一部を実際の空きメモリに調整します。
適用された最適化 (llama.cpp)
デフォルトでは、core/optimizer.py に基づきます。
最適な量子化: 適合するまで最高品質から最低品質まで繰り返します (Q8 → Q2)。
実行ごとの計画 (optimizer.plan_llamacpp_run): キューの最大コンテキストと ngl コンテキストが計算されます。

実際に実行される ant と選択された有効な KV (オプティマイザが選択するクオントではない) — 別のクオントの実行時に OOM が発生したり圧縮が無駄になったりする前に
実際のアーキテクチャの正確な KV キャッシュによる自動最大コンテキスト ( n_layer · n_head_kv · head_dim ) — GQA/MQA キャプチャ
KV キャッシュ圧縮 ( -ctk -ctv ): f16 / q8_0 / q4_0 (+ 独立した K/V およびプリセット)。量子化された KV は、-fa を強制的にオンにします (llama.cpp はこれを必要とします)。 ⚠️ K の q4_0 は攻撃的であり、小規模モデルでは生成が低下したり壊れたりする可能性があります。 q8_0 はスイートスポットです
RAM 内の KV ( --no-kv-offload 、プリセット Extreme ): VRAM を解放し、コンテキストによって RAM が制限され始めます。
小型 GPU 上の MoE モデルの MoE オフロード ( --n-cpu-moe N )
フラッシュ アテンション ( -fa on )、モデル全体が VRAM に収まる場合の mlock + --no-mmap
スレッド = 物理コア · バッチサイズ 2048 + ubatch サイズ 512
Engine_opts による合計オーバーライド (フラグの重複なし: オプティマイザがベースを設定し、ルールをオーバーライドします)
自己ダウンロードによるモデルのカタログ
backend/data/models.json には 124 のモデルがリストされています。 hf_gguf が設定されているものは、Hugging Face から自動ダウンロードされます。カバーなど:
フレイム3/3.1/3.2/3.3(1B→70B)、ネモトロン、ヘルメス、トゥール、ドルフィン
Qwen 2.5/3 (0.5B → 72B)、コーダー、数学、QwQ 32B (推論)、MoE 30B-A3B
ミストラル 7B、ニモ、スモール 3/3.1 24B、ミニストラル、コードストラル、ミストラル (MoE)
ファイ2 / 3 / 3.5 / 4 (+ミニ、+MoE)
DeepSeek R1 蒸留物、Coder、Coder-V2-Lite (MoE)
ビジョン: Qwen2-VL、Qwen2.5-VL、MiniCPM-V
コード: Code Llama、CodeGemma、StarCoder2、Yi-Coder、OpenCoder、安定版コード
詳細: Granite (IBM)、Falcon3、GLM-4、EXAONE、InternLM、OLMo、Aya/Command-R (Cohere)、SmolLM2、SOLAR、Zephyr…
発明されたデータはありません。各エントリは追加される前に HuggingFace に対して検証されます: GGUF ex リポジトリ

この場合、file_template は実際に公開されたファイルから派生し、Q4_K_M がチェックされて解決されます。このツールは backend/scripts/ ( verify_models.py + merge_models.py ) にあります。これを実行してカタログを安全に展開します。シャードに分割された巨大なモデル ( -00001-of-000NN.gguf 、例: DeepSeek-V3、Llama 4 70B+) もサポートされています。カタログでは hf_gguf.multipart がマークされており、ダウンロードによりすべてのシャードが結合されます (llama.cpp は同じディレクトリから兄弟をロードします)。
Node.js 20+: https://nodejs.org/
Python 3.11+: https://www.python.org/
uv:curl -LsSf https://astral.sh/uv/install.sh |しー
(オプション) Docker エンジンを使用している場合は Docker デスクトップ
(オプション) GPU アクセラレーションが必要な場合は、NVIDIA ドライバー (アプリは CUDA ビルドを検出してダウンロードします)
git クローン https://github.com/JoniMartin27/inferbench.git
cd推論ベンチ
# バックエンド(端末A)
CD バックエンド
uv venv -- Python 3.11
.venv\Scripts\activate
uv pip install -e 。
uvicorn main:app -- リロード -- ポート 7777
# フロントエンド（端末B）
CDフロントエンド
npmインストール
npm run electric:dev
アプリは Electron ウィンドウで開きます。サイドバーにはバックエンドの状態が表示され、ビュー間を移動します。
Python バックエンドは PyInstaller の実行可能ファイルとしてパッケージ化され、Electron インストーラーにサイドカーとして埋め込まれます。パッケージ化されたアプリでは、ターゲット マシンに Python は必要ありません。
# サイドカーを構築する
scripts\build-sidecar.ps1 # Windows
ベース

[切り捨てられた]

## Original Extract

Panel local multiplataforma para benchmark de motores de inferencia LLM (llama.cpp nativo + APIs cloud) - JoniMartin27/inferbench

GitHub - JoniMartin27/inferbench: Panel local multiplataforma para benchmark de motores de inferencia LLM (llama.cpp nativo + APIs cloud) · GitHub
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
JoniMartin27
/
inferbench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
73 Commits 73 Commits .claude/ skills/ run-inferbench-backend .claude/ skills/ run-inferbench-backend .github .github assets assets backend backend docker docker frontend frontend scripts scripts website website .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LAUNCH.md LAUNCH.md LICENSE LICENSE PROJECT_BRIEF.md PROJECT_BRIEF.md README.md README.md SECURITY-AUDIT.md SECURITY-AUDIT.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Descarga, arranca y benchmarkea motores de inferencia LLM locales con un solo click.
Sin Docker obligatorio. Sin tocar la línea de comandos. Tus datos nunca salen de tu máquina.
¿Quieres correr LLMs en local pero no sabes qué cuantización te entra en tu GPU , ni a cuántos tok/s va a ir , ni qué motor es más rápido para tu hardware ? InferBench lo responde por ti, midiendo de verdad — sin números inventados.
Eliges modelo + cuantizaciones → InferBench:
① descarga el binario del motor (release oficial GitHub)
② descarga los GGUF que falten desde Hugging Face
③ arranca el motor con la config óptima para tu hardware
④ ejecuta la suite de prompts midiendo TTFT, tok/s, VRAM, calidad
⑤ guarda los resultados y te deja compararlos lado a lado
Ejemplo real
RTX 3070 · 8 GB VRAM · 32 GB RAM — corriendo SmolLM2-360M Q8_0 con llama.cpp nativo:
Números sacados del runner de producción — sin datos inventados.
Coge el instalador para tu sistema desde la página de Releases — no necesitas Python ni Node instalados, el backend va embebido como sidecar:
¿Prefieres compilarlo tú? Salta a Instalación (desarrollo) .
Auto-bootstrap end-to-end : 1 click → binario + modelo + arranque + benchmark + cleanup
Modo nativo (sin Docker) para llama.cpp: descarga release pre-compilada de GitHub (auto-detecta CUDA, descarga también las DLLs del runtime)
Modo Docker disponible para cualquier motor que lo requiera
Detección automática de hardware : CPU, RAM, GPU (NVIDIA via NVML, AMD via rocm-smi, Apple Silicon via system_profiler). Cacheada para que el listado de compatibilidad sea instantáneo (~4 ms para 124 modelos)
Catálogo de 124 modelos con auto-descarga GGUF desde HF, todos verificados contra HuggingFace (incluye visión, código, reasoning y MoE). Ver Catálogo
Escaneo de GGUFs locales : detecta modelos de LM Studio, Ollama, HF cache, etc., con cuenta de parámetros real leída de la metadata GGUF (independiente del quant)
Visión real (multimodal) : para modelos de visión (Qwen2-VL, Qwen2.5-VL, MiniCPM-V) descarga el mmproj , arranca llama-server con --mmproj y benchmarkea un prompt con imagen real vía la API de visión OpenAI-compatible. También en APIs (gpt-4o, claude) y motores Docker (vLLM sirve el modelo de visión) — aunque vLLM con un modelo de visión necesita bastante VRAM (un 2B no entra holgado en 8 GB; en local, llama.cpp es la ruta fiable a poca VRAM)
Optimizador : dado tu hardware + modelo + motor, calcula la mejor cuantización, KV-cache, contexto máximo, MoE offload, flags
Compresión de KV-cache explicada : 5 presets (Calidad→Extremo) con qué hace / en qué afecta / qué permite, + tabla de los modelos más potentes que caben con cada compresión para tu hardware
Evaluación de calidad en 3 modos : scorer offline basado en referencia (sin GPU/API, corre en cualquier PC), LLM-judge con el motor local, o LLM-judge por API externa. Ver Calidad
Modo sweep : lanza el mismo modelo con N cuantizaciones distintas en cola
Comparación : selecciona varias runs del historial, ve métricas y gráficos lado a lado
SSE en vivo : progreso de descargas (con %), TTFT, tok/s actual, log estilo terminal
Descargas resilientes : GGUFs de decenas de GB con reintentos automáticos (backoff exponencial) y reanudación desde el parcial vía cabecera Range si la red se corta
API keys en el keyring del SO : guarda las claves de OpenAI/Anthropic/OpenRouter/NVIDIA en el gestor de credenciales del sistema (Ajustes → API keys); el benchmark las usa solo, nunca se escriben a disco
Observabilidad con lookspan (opt-in): si defines LOOKSPAN_ENDPOINT , cada run de benchmark se exporta como un trace a lookspan (span raíz + un llm_call por prompt con TTFT, tok/s, VRAM y calidad reales). Best-effort: un fallo de red nunca rompe el benchmark
Stop en cualquier momento : cancela bootstrap, descarga o ejecución
Persistencia : SQLite con todos los runs (engine, modelo, quant, flags, métricas por prompt, output bruto)
Capa
Tecnología
App de escritorio
Electron 42
Frontend
React 18 + Vite 8 + TailwindCSS + Recharts + lucide-react
Backend
FastAPI (Python 3.11) + uvicorn + sse-starlette
Persistencia
SQLite vía SQLModel
GPU detection
psutil + pynvml + system_profiler / rocm-smi
Containers
Docker SDK for Python
Native runtime
subprocess.Popen + binarios oficiales de GitHub
Empaquetado
electron-builder + PyInstaller
Motores soportados
Motor
Tipo
Modo nativo
Modo Docker
Auto-descarga modelo
llamacpp
local
✅ binarios oficiales
✅
✅ HuggingFace GGUF
ollama
local
✅ daemon Ollama
✅
✅ registro Ollama
vllm
local
—
✅ (GPU NVIDIA)
✅ HF (en contenedor)
sglang
local
—
✅ (GPU NVIDIA)
✅ HF (en contenedor)
tgi
local
—
✅ (GPU NVIDIA)
✅ HF (en contenedor)
openai
API
n/a
n/a
n/a
anthropic
API
n/a
n/a
n/a
openrouter
API
n/a
n/a
n/a
nvidia
API
n/a
n/a
n/a
Todos los motores locales tienen adaptador completo (build de comando por motor, bootstrap automático y schema de optimización propio). vLLM/SGLang/TGI son Docker-only y requieren GPU NVIDIA; el modelo lo descarga el propio contenedor desde HuggingFace (le pasamos el repo id). Las APIs cloud funcionan con tu API key (sólo parámetros de sampling, sin optimización local): OpenAI / OpenRouter / NVIDIA usan el endpoint OpenAI-compatible /v1/chat/completions ; Anthropic usa su API nativa ( /v1/messages , header x-api-key + anthropic-version , system aparte), no es OpenAI-compatible.
Speculative decoding (DFLASH) : vLLM y SGLang aceptan acelerar con un modelo draft block-diffusion ( DFLASH , 6-8× sin pérdida de calidad). En Benchmark, con vLLM/SGLang seleccionado, activa DFLASH e indica el modelo draft (p.ej. z-lab/Qwen3.5-35B-A3B-DFlash ). SGLang es la ruta oficial; vLLM necesita una build con soporte. Requiere VRAM para el modelo y el draft, así que no entra en GPUs pequeñas.
Estado de verificación: los 5 motores locales ( llamacpp , ollama , vllm , sglang , tgi ) verificados end-to-end por el runner de producción (bootstrap → arranque → inferencia real con tps>0 → parada sin contenedores colgados) en GPU NVIDIA (RTX 3070, 8 GB). vLLM/SGLang ajustan la fracción de VRAM a la memoria libre real para no fallar en GPUs no vacías.
Optimizaciones aplicadas (llama.cpp)
Por defecto, basadas en core/optimizer.py :
Cuantización óptima : itera de mayor a menor calidad (Q8 → Q2) hasta que cabe
Plan por run ( optimizer.plan_llamacpp_run ): el contexto máximo y ngl se calculan para el quant que de verdad se ejecuta y la KV efectiva elegida (no para el quant que el optimizer habría elegido) — antes podía dar OOM al correr un quant distinto, o desaprovechar la compresión
Contexto máximo automático con KV-cache exacta de la arquitectura real ( n_layer · n_head_kv · head_dim ) — captura GQA/MQA
KV-cache compresión ( -ctk -ctv ): f16 / q8_0 / q4_0 (+ K/V independientes y presets). La KV cuantizada fuerza -fa on (llama.cpp lo exige). ⚠️ q4_0 en K es agresivo y puede degradar/romper la generación en modelos pequeños; q8_0 es el punto dulce
KV en RAM ( --no-kv-offload , preset Extremo ): libera VRAM y el contexto pasa a limitarlo la RAM
MoE offload ( --n-cpu-moe N ) para modelos MoE en GPUs pequeñas
Flash Attention ( -fa on ), mlock + --no-mmap cuando el modelo cabe entero en VRAM
Threads = núcleos físicos · batch-size 2048 + ubatch-size 512
Override total via engine_opts (sin duplicar flags: el optimizer pone la base, tus overrides mandan)
Catálogo de modelos con auto-descarga
backend/data/models.json lista 124 modelos . Los que tienen hf_gguf configurado se auto-descargan desde Hugging Face. Cubre, entre otros:
Llama 3 / 3.1 / 3.2 / 3.3 (1B → 70B), Nemotron, Hermes, Tulu, Dolphin
Qwen 2.5 / 3 (0.5B → 72B), Coder, Math, QwQ 32B (reasoning), MoE 30B-A3B
Mistral 7B, Nemo, Small 3/3.1 24B, Ministral, Codestral, Mixtral (MoE)
Phi 2 / 3 / 3.5 / 4 (+ mini, + MoE)
DeepSeek R1 distills, Coder, Coder-V2-Lite (MoE)
Visión : Qwen2-VL, Qwen2.5-VL, MiniCPM-V
Código : Code Llama, CodeGemma, StarCoder2, Yi-Coder, OpenCoder, Stable Code
Más : Granite (IBM), Falcon3, GLM-4, EXAONE, InternLM, OLMo, Aya/Command-R (Cohere), SmolLM2, SOLAR, Zephyr…
Sin datos inventados. Cada entrada se verifica contra HuggingFace antes de añadirse: el repo GGUF existe, el file_template se deriva de los archivos reales publicados y se comprueba que el Q4_K_M resuelve. La herramienta está en backend/scripts/ ( verify_models.py + merge_models.py ); ejecútala para ampliar el catálogo de forma segura. Los modelos enormes partidos en shards ( -00001-of-000NN.gguf , p.ej. DeepSeek-V3, Llama 4 70B+) también se soportan: el catálogo marca hf_gguf.multipart y la descarga junta todos los shards (llama.cpp carga las hermanas del mismo directorio).
Node.js 20+ : https://nodejs.org/
Python 3.11+ : https://www.python.org/
uv : curl -LsSf https://astral.sh/uv/install.sh | sh
(Opcional) Docker Desktop si vas a usar motores Docker
(Opcional) Driver NVIDIA si quieres acceleración GPU (la app detecta y descarga la build CUDA)
git clone https: // github.com / JoniMartin27 / inferbench.git
cd inferbench
# Backend (terminal A)
cd backend
uv venv -- python 3.11
.venv\Scripts\activate
uv pip install - e .
uvicorn main:app -- reload -- port 7777
# Frontend (terminal B)
cd frontend
npm install
npm run electron:dev
Se abre la app en una ventana Electron. La sidebar muestra la salud del backend y navega entre vistas.
El backend Python se empaqueta como ejecutable con PyInstaller y se embebe como sidecar en el instalador Electron. La app empaquetada no requiere Python en la máquina destino.
# Construir el sidecar
scripts\build-sidecar.ps1 # Windows
bas

[truncated]
