---
source: "https://github.com/madara88645/Cognigraph"
hn_url: "https://news.ycombinator.com/item?id=48396828"
title: "CogniGraph – describe a scenario, LLM picks a brain lobe and neuromodulator"
article_title: "GitHub - madara88645/Cognigraph: Educational web demo: natural-language scenarios → LLM brain/ neuromodulator classification → Brian2 SNN → 3D brain visualization. Not medical software. · GitHub"
author: "madara8865"
captured_at: "2026-06-04T13:15:36Z"
capture_tool: "hn-digest"
hn_id: 48396828
score: 1
comments: 0
posted_at: "2026-06-04T10:52:46Z"
tags:
  - hacker-news
  - translated
---

# CogniGraph – describe a scenario, LLM picks a brain lobe and neuromodulator

- HN: [48396828](https://news.ycombinator.com/item?id=48396828)
- Source: [github.com](https://github.com/madara88645/Cognigraph)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T10:52:46Z

## Translation

タイトル: CogniGraph – シナリオを説明し、LLM が脳葉と神経調節物質を選択します
記事のタイトル: GitHub - madara88645/コグニグラフ: 教育ウェブ デモ: 自然言語シナリオ → LLM 脳/神経調節物質分類 → Brian2 SNN → 3D 脳視覚化。医療用ソフトウェアではありません。 · GitHub
説明: 教育用 Web デモ: 自然言語シナリオ → LLM 脳/神経調節物質分類 → Brian2 SNN → 3D 脳視覚化。医療用ソフトウェアではありません。 - madara88645/コグニグラフ

記事本文:
GitHub - madara88645/Cognigraph: 教育 Web デモ: 自然言語シナリオ → LLM 脳/神経調節物質分類 → Brian2 SNN → 3D 脳視覚化。医療用ソフトウェアではありません。 · GitHub
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
マダラ88645
/
コグニグラフ
公共
通知

イオン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
80 コミット 80 コミット .github/ workflows .github/ workflows .jules .jules バックエンド バックエンド ドキュメント ドキュメント フロントエンド フロントエンド スクリプト スクリプト src src テスト テスト .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version .vercelignore .vercelignore AGENTS.md AGENTS.md Baslat-Cognigraph.bat Baslat-Cognigraph.bat CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md benchmark.py benchmark.py benchmark_concurrent.py benchmark_concurrent.py benchmark_run_snn.py benchmark_run_snn.py fly.toml fly.toml pyproject.toml pyproject.toml pytest.ini pytest.ini 要件.txt 要件.txt 要件.vercel.txt 要件.vercel.txt run_bench.py run_bench.py start-cognigraph.bat start-cognigraph.bat test_main_async.py test_main_async.py vercel.json vercel.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
CogniGraph (リポジトリ フォルダー: Cognigraph ) は、小規模な教育用デモです。現実世界のシナリオを説明し、LLM が脳葉と神経調節物質の緊張を分類し、Brian2 スパイキング ニューラル ネットワーク (SNN) がシミュレートされ、Web UI が 3D 脳モデル上のアクティビティを視覚化します。 UI は、frontend/js/ にある静的 HTML と ES モジュールとして / から提供されます (バンドラーなし)。オプションの OpenRouter キーは、ページ内の API 設定パネルにのみ入力されます (個別の認証ルートやリダイレクトはありません)。
これは医療用ソフトウェアではありません。出力は視覚化と学習のみを目的としており、診断や治療を目的とするものではありません。 UI には、シミュレーションのメタファーとしてではなく、モデル化されたストレス ホルモン軸 (HPA / コルチゾールなど) のコンテキストが含まれています。

臨床測定。コルチゾールの場合、neuromodulator_intensity は、ホルモン濃度ではなく、急性最適対慢性負荷シミュレーション方式 (しきい値 0.5) を選択します。
Neural Activation Viewer — シナリオ入力、再生コントロール、および認知分析 (例: 重いデッドリフトを行う )。
シミュレーション ビュー - 再生完了後の色付きのローブ メッシュ、スパイク カウンタ、HPA コンテキスト、およびイベント ログ。
アーキテクチャ (リクエスト パイプライン)
フローチャート LR
U[ユーザーシナリオ] --> FE[index.html / 分析]
FE -- POST /simulate --> API[FastAPI backend/main.py]
API --> LLM[OpenRouter LLM 分類_シナリオ]
LLM --> NM[neuromodulation.py 検証 + パラメータの解決]
NM --> SNN[Brian2 SNN run_snn]
SNN --> VFX[build_vfx_profile]
VFX -- JSON --> FE
FE --> Three[Three.js グロー/ブルーム レンダリング]
読み込み中
要件
Python 3.10+ を推奨 (3.x が必要)
Brian2 のパフォーマンスを最大限に発揮するには、プラットフォームによっては C コンパイラが必要な場合があります。 Brian2 のインストールに関するドキュメントを参照してください。
cd コグニグラフ
Python -m venv .venv
# Windows: .venv\Scripts\activate
# Unix: ソース .venv/bin/activate
pip install -r 要件.txt
構成
.env.example をプロジェクト ルートの .env にコピーします。
OpenRouter から OPENROUTER_API_KEY を設定します。
オプションで OPENROUTER_DEMO_MODEL を設定します (デフォルト: qwen/qwen3.5-flash-02-23 ) — より強力な教育者スタイルのシステム プロンプトを使用して、公開デモでの匿名/ブラウザキーなしのトラフィックに使用されます。代替案: openai/gpt-oss-120b 、または 無料層の openai/gpt-oss-120b:free (レート制限が適用されます)。モデルを参照してください。
オプションで OPENROUTER_MODEL を設定します (デフォルト: deepseek/deepseek-v4-flash ) — 訪問者が自分のキーを UI ( X-OpenRouter-Api-Key ) に保存する場合にのみ使用されます。彼らはあなたではなく OpenRouter に支払います。
.env は決してコミットしないでください。 .gitignore にリストされています。
リポジトリ ルートから (依存関係がインストールされている):
python -m uvicorn backend.main:a

pp --host 0.0.0.0 --port 8000 --reload
次に、 http://127.0.0.1:8000 を開きます。
Windows: start-cognigraph.bat をダブルクリックするか、トルコ語メッセージの場合は Baslat-Cognigraph.bat を使用します。
運用環境のようなテストには次のコマンドを使用します ( --reload なし)。
python -m uvicorn backend.main:app --host 0.0.0.0 --port 8000
テスト
pytest
構成: pytest.ini 、 testing/ の下のテスト。
方法
パス
説明
ゲット
/
Web UI (frontend/index.html) を提供します。
ゲット
/healthz
プラットフォーム プローブ用の軽量ヘルス チェック エンドポイント。
投稿
/シミュレート
分類 + SNN を実行し、スパイクと VFX エコーを返します。
POST /JSON 本文をシミュレートします:
オプションのリクエスト ヘッダー (UI と同じオリジン。共有デモには必要ありません):
応答 (簡略化): active_lobe 、 dominant_neuromodulator 、neuromodulator_intensity 、neuromodulator_rationale 、 description 、duration_ms 、スパイク (ローブごとのスパイクのインデックスと時間)、 snn_modulation 、 vfx_profile 。
OPENROUTER_API_KEY が欠落している場合、API は明確なメッセージを含む 503 を返します。
静的ファイルは、frontend/ ディレクトリから /static にマウントされます。
Fly.io (プライマリ /シミュレート バックエンド): https://cognigraph-13906.fly.dev
Vercel (UI ミラー、プロキシ /simulate to Fly): https://cognigraph-tau.vercel.app
2 つのホストの関係: Fly は、長寿命の Docker VM で完全な FastAPI + Brian2 スタックを実行します ( min_machines_running = 1 を介してウォーム)。 Vercel は静的 UI を提供し、/simulate と /healthz を外部プロキシ ( vercel.json ) として Fly に書き換えるため、Vercel ページは Vercel で OPENROUTER_API_KEY を必要とせずにエンドツーエンドで動作します。サーバー キーは Fly 上に存在します。同一オリジンの書き換えとは、ブラウザーが引き続き cognigraph-tau.vercel.app/simulate に POST することを意味するため、CORS は適用されず、BYOK ヘッダー ( X-OpenRouter-Api-Key 、 X-OpenRouter-Model ) はそのまま通過します。
待ち時間: 分析は通常、暖かい Fly m で 6 ～ 10 秒です。

achine (LLM + Brian2)。 Vercel プロキシ パスは、上部に小さなエッジ ホップを追加します。古い Vercel-Python 関数パス (10 ～ 60 秒のサーバーレス予算内でコールドスタートする Brian2) は、 /simulate には使用されなくなりました。
各ユーザーは UI ([API 設定] パネル) で独自の OpenRouter キーを指定できます。
キーはユーザーのブラウザのローカル ストレージに保存され、 X-OpenRouter-Api-Key として送信されます。
ユーザーキーが存在する場合、同じパネルからのオプションのモデル ID が X-OpenRouter-Model として送信されます。省略または無効な場合、サーバーは OPENROUTER_MODEL を使用します。ユーザーキーがないと、X-OpenRouter-Model は無視されます (共有トラフィックは常に OPENROUTER_DEMO_MODEL を使用します)。
サーバー側の環境キー ( OPENROUTER_API_KEY ) は、キーを追加しない訪問者のフォールバックとして引き続きサポートされます。
X-OpenRouter-Api-Key を使用しないリクエストでは、OPENROUTER_DEMO_MODEL (デフォルトは qwen/qwen3.5-flash-02-23 ) に加えて神経科学者と教育者のシステム プロンプトが使用されます。ユーザー キーを含むリクエストでは、OPENROUTER_MODEL または検証された X-OpenRouter-Model 値を使用します (請求は OpenRouter アカウントで行われます)。
共有/パブリックデバイスの場合、ユーザーは使用後に保存したキーをクリアする必要があります。
Vercel は UI ミラーとして構成されます。Vercel は静的フロントエンドを提供し、 vercel.json の外部書き換えを介して動的エンドポイント ( /simulate 、 /healthz ) を Fly にプロキシします。 Vercel では LLM 関連の環境変数を設定しないでください。サーバーキーは Fly 上に存在し、プロキシを介して流れます。
CLI をインストールしてログインします。
npm i -g ヴェルセル
ヴェルセルログイン
プロジェクトのルートで、以下をデプロイします。
ヴェルセル
Vercel では OPENROUTER_API_KEY やモデル環境変数は必要ありません。フライはそれらを所有しています。
vercel.json を変更した後、運用環境に再デプロイします。
vercel --prod
Vercel で Python を実行する代わりにプロキシを使用するのはなぜですか? Vercel の Python 関数は、厳密な関数期間予算内で Brian2 をコールドスタートします (Hobby は 10 秒、Pro のデフォルトは 15 秒、プロジェクトごとの conf 経由でのみ最大 300 秒)

イグ）。 vercel.json の関数キーは、従来の api/ ディレクトリに対してグロブを検証するだけなので、Vercel の FastAPI 自動検出に必要な src/index.py では、vercel.json で maxDuration を宣言することはできません。両方の制約に同時に対処するのではなく、/simulate は Fly に委任されます (VM の寿命が長く、機能期間の上限なし、ウォーム後の Brian2 コールド スタートなし)。 src/index.py の Vercel Python 関数は、/ および /static/* のみを提供するために残ります。
インストールして認証します。
フライ認証ログイン
アプリを (1 回) 作成し、デプロイします。
フライ起動 --no-deploy
フライデプロイ
秘密鍵を設定します。
フライシークレットセット OPENROUTER_API_KEY=your_key_here
オプションのモデル オーバーライド:
フライ シークレット セット OPENROUTER_DEMO_MODEL=qwen/qwen3.5-flash-02-23
フライ シークレット セット OPENROUTER_MODEL=ディープシーク/ディープシーク-v4-フラッシュ
共有デモには OPENROUTER_DEMO_MODEL を使用します。 OPENROUTER_MODEL は BYOK リクエストにのみ適用されます。
シークレット/構成の変更後に再デプロイします。
フライデプロイ
このリポジトリには、 uvicorn backend.main:app 用に構成された fly.toml と Dockerfile が含まれています。 fly.toml は min_machines_running = 1 を設定するため、1 台のマシンがウォーム状態を維持します。これにより、Vercel UI プロキシ経由で到着するリクエストを含む、アイドル後の最初のリクエストでの 15 ～ 30 秒の Brian2 コールドスタートが回避されます。
デプロイされた URL ( $BASE_URL ) に対して次のチェックを実行します。
カール -fsS " $BASE_URL /healthz "
カール -fsS " $BASE_URL / " > /dev/null
curl -sS -X POST " $BASE_URL /simulate " \
-H " Content-Type: application/json " \
-d " { \" プロンプト \" : \" 複雑な数学の問題を解く \" } "
予想される動作:
/healthz は {"status":"ok"} を返します。
/simulate は、 active_lobe 、 dominant_neuromodulator 、spiks を含む JSON を返します。
キーが見つからない場合、/simulate はキー構成ヒントを含む 503 を返します。
Vercel は、vercel.json の外部書き換えを介して /simulate と /healthz をプロキシして Fly にする UI ミラーになりました。フライはシングル/simです

ULATEバックエンド。 Vercel には LLM 環境変数はもう存在しません。
fly.toml は min_machines_running = 1 (そして auto_stop_machines を無効にします) を設定するため、Brian2 はウォーム状態を維持し、アイドル後の最初のリクエストは 15 ～ 30 秒のコールドスタートになりません。
共有デモ トラフィックは OPENROUTER_DEMO_MODEL (デフォルトの Qwen 3.5 Flash + 教育者プロンプト) を使用します。 BYOK トラフィックは OPENROUTER_MODEL (デフォルトの DeepSeek V4 Flash) を使用します。
LLM エラー処理のセキュリティ強化により、機密性の高いアップストリームの詳細が API クライアントに公開されるのを防ぎます。
FastAPI の有効期間を通じて httpx.AsyncClient を再利用することで、リクエスト処理が高速化されます。
run_snn ループ内で繰り返される辞書作成を削除することで、SNN ランタイムを最適化します。
静的プロファイル定義をモジュール スコープに移動することによる build_vfx_profile の最適化。
_strip_markdown_fences 、 _load_dotenv_file 、 _lerp_toward_neutral 、 snn_params_to_dict 、ペイロード長検証、および GET / (serve_index) のテスト カバレッジを追加しました。
LinkedIn、X、Reddit、またはブログに投稿する場合は、この宣伝文句を使用してください。ソースを公開する場合は、YOUR_REPO_URL を置き換えます。
CogniGraph — シナリオを説明します。 LLM は脳葉と神経調節物質の調子を選択します。 Brian2 スパイキング ネットワークが実行されます。 3D 脳が結果を視覚化します。ライブデモ: https://cognigraph-tau.vercel.app — 医療ソフトウェアではありません。学習とデモ専用

[切り捨てられた]

## Original Extract

Educational web demo: natural-language scenarios → LLM brain/ neuromodulator classification → Brian2 SNN → 3D brain visualization. Not medical software. - madara88645/Cognigraph

GitHub - madara88645/Cognigraph: Educational web demo: natural-language scenarios → LLM brain/ neuromodulator classification → Brian2 SNN → 3D brain visualization. Not medical software. · GitHub
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
madara88645
/
Cognigraph
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
80 Commits 80 Commits .github/ workflows .github/ workflows .jules .jules backend backend docs docs frontend frontend scripts scripts src src tests tests .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version .vercelignore .vercelignore AGENTS.md AGENTS.md Baslat-Cognigraph.bat Baslat-Cognigraph.bat CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md benchmark.py benchmark.py benchmark_concurrent.py benchmark_concurrent.py benchmark_run_snn.py benchmark_run_snn.py fly.toml fly.toml pyproject.toml pyproject.toml pytest.ini pytest.ini requirements.txt requirements.txt requirements.vercel.txt requirements.vercel.txt run_bench.py run_bench.py start-cognigraph.bat start-cognigraph.bat test_main_async.py test_main_async.py vercel.json vercel.json View all files Repository files navigation
CogniGraph (repository folder: Cognigraph ) is a small educational demo: you describe a real-world scenario, an LLM classifies brain lobe and neuromodulator tone, a Brian2 spiking neural network (SNN) is simulated, and a web UI visualizes activity on a 3D brain model. The UI is served from / as static HTML plus ES modules under frontend/js/ (no bundler); optional OpenRouter keys are entered only in the in-page API Settings panel (no separate auth route or redirect).
This is not medical software. Outputs are for visualization and learning only, not diagnosis or treatment. The UI includes context for modeled stress-hormone axes (for example HPA / cortisol) as simulation metaphors, not clinical measurements. For cortisol , neuromodulator_intensity selects an acute optimal vs chronic-load simulation regime (threshold at 0.5), not a hormone concentration.
Neural Activation Viewer — scenario input, playback controls, and cognitive analysis (example: Doing a heavy deadlift ).
Simulation view — colored lobe mesh, spike counters, HPA context, and event log after playback completes.
Architecture (request pipeline)
flowchart LR
U[User scenario] --> FE[index.html / Analyze]
FE -- POST /simulate --> API[FastAPI backend/main.py]
API --> LLM[OpenRouter LLM classify_scenario]
LLM --> NM[neuromodulation.py validate + resolve params]
NM --> SNN[Brian2 SNN run_snn]
SNN --> VFX[build_vfx_profile]
VFX -- JSON --> FE
FE --> Three[Three.js glow/bloom render]
Loading
Requirements
Python 3.10+ recommended (3.x required)
Brian2 may need a C compiler on some platforms for full performance; see the Brian2 installation docs .
cd Cognigraph
python -m venv .venv
# Windows: .venv\Scripts\activate
# Unix: source .venv/bin/activate
pip install -r requirements.txt
Configuration
Copy .env.example to .env in the project root.
Set OPENROUTER_API_KEY from OpenRouter .
Optionally set OPENROUTER_DEMO_MODEL (default: qwen/qwen3.5-flash-02-23 ) — used for anonymous / no-browser-key traffic on public demos, with a stronger educator-style system prompt. Alternatives: openai/gpt-oss-120b , or openai/gpt-oss-120b:free for a no-cost tier (rate limits apply). See models .
Optionally set OPENROUTER_MODEL (default: deepseek/deepseek-v4-flash ) — used only when a visitor saves their own key in the UI ( X-OpenRouter-Api-Key ); they pay OpenRouter, not you.
Never commit .env ; it is listed in .gitignore .
From the repository root (with dependencies installed):
python -m uvicorn backend.main:app --host 0.0.0.0 --port 8000 --reload
Then open http://127.0.0.1:8000 .
Windows: double-click start-cognigraph.bat , or use Baslat-Cognigraph.bat for Turkish messages.
Use this command for production-like testing (no --reload ):
python -m uvicorn backend.main:app --host 0.0.0.0 --port 8000
Tests
pytest
Configuration: pytest.ini , tests under tests/ .
Method
Path
Description
GET
/
Serves the web UI ( frontend/index.html ).
GET
/healthz
Lightweight health check endpoint for platform probes.
POST
/simulate
Runs classification + SNN + returns spikes and VFX echo.
POST /simulate JSON body:
Optional request headers (same origin as the UI; not required for the shared demo):
Response (simplified): active_lobe , dominant_neuromodulator , neuromodulator_intensity , neuromodulator_rationale , explanation , duration_ms , spikes (per-lobe spike indices and times), snn_modulation , vfx_profile .
If OPENROUTER_API_KEY is missing, the API returns 503 with a clear message.
Static files are mounted at /static from the frontend/ directory.
Fly.io (primary /simulate backend): https://cognigraph-13906.fly.dev
Vercel (UI mirror, proxies /simulate to Fly): https://cognigraph-tau.vercel.app
How the two hosts relate: Fly runs the full FastAPI + Brian2 stack in a long-lived Docker VM (warm via min_machines_running = 1 ). Vercel serves the static UI and rewrites /simulate and /healthz to Fly as external proxies ( vercel.json ), so the Vercel page works end-to-end without needing OPENROUTER_API_KEY on Vercel — the server key lives on Fly. Same-origin rewrites mean the browser still POSTs to cognigraph-tau.vercel.app/simulate , so CORS does not apply and BYOK headers ( X-OpenRouter-Api-Key , X-OpenRouter-Model ) pass through untouched.
Latency: Analyze is typically 6–10 seconds on a warm Fly machine (LLM + Brian2). The Vercel-proxied path adds a small edge hop on top. The old Vercel-Python-function path (cold-starting Brian2 inside a 10–60s serverless budget) is no longer used for /simulate .
Each user can provide their own OpenRouter key in the UI ( API Settings panel).
The key is stored in the user's browser local storage and sent as X-OpenRouter-Api-Key .
Optional model id from the same panel is sent as X-OpenRouter-Model when a user key is present; if omitted or invalid, the server uses OPENROUTER_MODEL . Without a user key, X-OpenRouter-Model is ignored (shared traffic always uses OPENROUTER_DEMO_MODEL ).
Server-side env key ( OPENROUTER_API_KEY ) is still supported as fallback for visitors who do not add a key.
Requests without X-OpenRouter-Api-Key use OPENROUTER_DEMO_MODEL (default qwen/qwen3.5-flash-02-23 ) plus a neuroscientist-educator system prompt; requests with a user key use OPENROUTER_MODEL or the validated X-OpenRouter-Model value (billing is on their OpenRouter account).
For shared/public devices, users should clear their saved key after use.
Vercel is configured as a UI mirror — it serves the static frontend and proxies the dynamic endpoints ( /simulate , /healthz ) to Fly via external rewrites in vercel.json . Do not set LLM-related env vars on Vercel; the server key lives on Fly and flows through the proxy.
Install CLI and login:
npm i -g vercel
vercel login
In project root, deploy:
vercel
No OPENROUTER_API_KEY or model env vars are needed on Vercel. Fly owns them.
Redeploy to production after any vercel.json change:
vercel --prod
Why the proxy instead of running Python on Vercel? Vercel’s Python functions cold-start Brian2 inside a strict function-duration budget (Hobby ~10s, Pro default 15s and up to 300s only via per-project config). The functions key in vercel.json only validates globs against the legacy api/ directory, so src/index.py — required for Vercel’s FastAPI auto-detection — cannot have its maxDuration declared in vercel.json at all. Rather than fight both constraints at once, /simulate is delegated to Fly (long-lived VM, no function duration cap, no Brian2 cold-start after warm). The Vercel Python function in src/index.py remains for serving / and /static/* only.
Install and auth:
fly auth login
Create app (once) and deploy:
fly launch --no-deploy
fly deploy
Set secret key:
fly secrets set OPENROUTER_API_KEY=your_key_here
Optional model overrides:
fly secrets set OPENROUTER_DEMO_MODEL=qwen/qwen3.5-flash-02-23
fly secrets set OPENROUTER_MODEL=deepseek/deepseek-v4-flash
Use OPENROUTER_DEMO_MODEL for the shared demo; OPENROUTER_MODEL only applies to BYOK requests.
Redeploy after secret/config changes:
fly deploy
This repo includes fly.toml and Dockerfile configured for uvicorn backend.main:app . fly.toml sets min_machines_running = 1 so one machine stays warm — this avoids a 15–30s Brian2 cold-start on the first request after idle, including requests that arrive via the Vercel UI proxy.
Run these checks against deployed URL ( $BASE_URL ):
curl -fsS " $BASE_URL /healthz "
curl -fsS " $BASE_URL / " > /dev/null
curl -sS -X POST " $BASE_URL /simulate " \
-H " Content-Type: application/json " \
-d " { \" prompt \" : \" Solving a complex math problem \" } "
Expected behavior:
/healthz returns {"status":"ok"} .
/simulate returns JSON with active_lobe , dominant_neuromodulator , spikes .
If key is missing, /simulate returns 503 with key configuration hint.
Vercel is now a UI mirror that proxies /simulate and /healthz to Fly via external rewrites in vercel.json ; Fly is the single /simulate backend. No LLM env vars live on Vercel anymore.
fly.toml sets min_machines_running = 1 (and disables auto_stop_machines ) so Brian2 stays warm and the first request after idle is not a 15–30s cold-start.
Shared demo traffic uses OPENROUTER_DEMO_MODEL (default Qwen 3.5 Flash + educator prompt); BYOK traffic uses OPENROUTER_MODEL (default DeepSeek V4 Flash).
Security hardening in LLM error handling to avoid exposing sensitive upstream details to API clients.
Faster request handling by reusing httpx.AsyncClient through FastAPI lifespan.
SNN runtime optimization by removing repeated dictionary creation inside the run_snn loop.
build_vfx_profile optimization by moving static profile definitions to module scope.
Added test coverage for _strip_markdown_fences , _load_dotenv_file , _lerp_toward_neutral , snn_params_to_dict , payload length validation, and GET / ( serve_index ).
Use this blurb when posting to LinkedIn, X, Reddit, or a blog. Replace YOUR_REPO_URL if you publish the source.
CogniGraph — Describe a scenario; an LLM picks a brain lobe and neuromodulator tone; a Brian2 spiking network runs; a 3D brain visualizes the result. Live demo: https://cognigraph-tau.vercel.app — Not medical software ; for learning and demos only

[truncated]
