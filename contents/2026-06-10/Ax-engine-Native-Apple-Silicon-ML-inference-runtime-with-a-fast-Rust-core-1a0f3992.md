---
source: "https://github.com/defai-digital/ax-engine"
hn_url: "https://news.ycombinator.com/item?id=48469752"
title: "Ax-engine: Native Apple Silicon ML inference runtime with a fast Rust core"
article_title: "GitHub - defai-digital/ax-engine: Apple Silicon LLM runtime supporting Gemma 4 and Qwen 3.6 MTP modes · GitHub"
author: "automatosx"
captured_at: "2026-06-10T00:59:17Z"
capture_tool: "hn-digest"
hn_id: 48469752
score: 2
comments: 0
posted_at: "2026-06-10T00:41:14Z"
tags:
  - hacker-news
  - translated
---

# Ax-engine: Native Apple Silicon ML inference runtime with a fast Rust core

- HN: [48469752](https://news.ycombinator.com/item?id=48469752)
- Source: [github.com](https://github.com/defai-digital/ax-engine)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T00:41:14Z

## Translation

タイトル: Ax-engine: 高速 Rust コアを備えたネイティブ Apple Silicon ML 推論ランタイム
記事のタイトル: GitHub - defai-digital/ax-engine: Gemma 4 および Qwen 3.6 MTP モードをサポートする Apple Silicon LLM ランタイム · GitHub
説明: Gemma 4 および Qwen 3.6 MTP モードをサポートする Apple Silicon LLM ランタイム - defai-digital/ax-engine

記事本文:
GitHub - defai-digital/ax-engine: Gemma 4 および Qwen 3.6 MTP モードをサポートする Apple Silicon LLM ランタイム · GitHub
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
デファイデジタル
/
斧エンジン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

s
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,576 コミット 1,576 コミット .cargo .cargo .github/ workflows .github/ workflows ベンチマーク ベンチマーク build/ metal build/ metal crates crates docs docs 例 例 javascript/ ax-engine javascript/ ax-engine metal metal proto/ ax_engine/ v1 proto/ ax_engine/ v1 python python qa qa スクリプト スクリプト sdk sdk .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AX Engine は、Mac ファーストの LLM 推論ランタイム、ローカル サーバー、SDK レイヤー、および Apple Silicon 用のベンチマーク ツールキットです。直接サポートの MLX モデル ファミリをネイティブに実行し、他の MLX テキスト モデルまたは非 MLX モデルを明示的な mlx-lm および llama.cpp 互換ルート経由でルーティングします。
AX Engine は、どのランタイム パスが作業を行っているかを隠すことなく、Apple Silicon 上にローカルの OpenAI 互換モデル サーバーを構築したい開発者向けです。
Python、TypeScript/JavaScript、Go、Ruby、Mojo 用の SDK を使用した、一般的なチャットおよび完了フロー用の OpenAI 互換ローカル テキスト エンドポイント。
直接サポートの Gemma および Qwen ファミリのリポジトリ所有の MLX ランタイム パス。委任されたルートは明示的に維持されます。
証拠が揃っている発表準備完了のベンチマーク主張: Gemma 4 12B アシスタント MTP は、同じアーティファクトの直接デコードより 2.83 ～ 2.92 倍高速であり、Qwen3.6 35B-A3B AX MTP は、公開サイドカーフェア マトリックスで MTPLX より +76.4% 高速です。
ワークロード コントラクト ベンチマーク ツールは、ルート ID、アーティファクト、プロンプト スイート、サンプラー、クールダウン、受け入れ率、ダーティ ステートの来歴を記録します。
パフォーマンス
ジェマ4 12B
ジェマ 4 12B マルチモーダル
投機的デコーディング (MTP)
ゲ

総合格闘技 4
ダイレクトデコード・プレフィル・TTFT
pip install " ax-engine[download] "
小さなモデルをダウンロードしてサーバーを起動します。
MODEL_DIR= " $( ax-engine download mlx-community/Qwen3-4B-4bit --json | python3 -c ' import json,sys; print(json.load(sys.stdin)["dest"]) ' ) "
ax-engine は " $MODEL_DIR " --port 8080 を提供します
ハイメモリ モデルのショートカット:
# 1 つ選択してください:
ax-engine サーブ qwen36-35b --download --port 8080
ax-engine は gemma4-12b --download --port 8080 を提供します
任意の OpenAI クライアントから呼び出します。
openaiインポートからOpenAI
client = OpenAI (base_url = "http://127.0.0.1:8080/v1" 、api_key = "local")
モデル = クライアント。モデル。リスト（）。データ [ 0 ]。 ID
resp = クライアント 。チャット 。完成品。作成(
モデル = モデル、
messages = [{ "role" : "user" , "content" : "AGI とは何ですか?" }]、
max_tokens = 128 、
)
print (それぞれの選択肢 [ 0 ] のメッセージ . コンテンツ )
または、Python SDK を直接使用します。
ax_engine import download_model から、セッション
パス = download_model ( "mlx-community/Qwen3-4B-4bit" )
Session ( mlx = True 、 mlx_model_artifacts_dir = str ( path )) を s として使用します。
print ( s .generate ([ 1 , 2 , 3 ], max_output_tokens = 8 ).output_tokens )
クイック スタートには、32 GB 以上のユニファイド メモリを搭載した Apple Silicon M2 Max 以降の macOS 14 (Sonoma) 以降が必要です。 Qwen3.6 35B-A3B や Gemma 4 12B などの大型モデルには、「標準的なハードウェア」にリストされているメモリ層が必要です。
pip インストール ax-engine
macOS 14 以降、Apple Silicon (M2 Max 以降)、Python 3.10 以降が必要です。 pip ホイールには、ax-engine オーケストレーション CLI と ax-engine-server バイナリが含まれています。どちらもインストール後の PATH 上にあります。
pip install " ax-engine[download] " # Hugging Face Hub ダウンロードと mlx-lm ツール
自作
同じトップレベルの ax-engine CLI とネイティブの ax-engine-server の場合、
斧エンジンベンチツール:
醸造タップ デファイデジタル/axエンジン
醸造インストール ax-eng

イネ
斧機関士
ライブラリがロードされていない: libmlxc.dylib というエラーでドクターが失敗した場合は、次のコマンドを実行します。
brew install mlx-c && brew reinstall ax-engine 。
醸造インストールmlx-c
カーゴビルド --workspace --release
maturin は # Python バインディングを開発します
モデルの取得
AX エンジンには、事前にサニタイズされた MLX ウェイトが必要です。推奨されるソースは mlx-community です。そこにあるモデルはすでに変換され、検証されています。
ax-engine download 、 download_model() 、および scripts/download_model.py は、重みをダウンロードし、必要な model-manifest.json を 1 つのステップで自動生成します。
# サポートされているダウンロード ターゲットをリストします。
ax-engine ダウンロード --list
# エイリアスでダウンロード
ax-engine ダウンロード qwen36-35b --json
ax-engine ダウンロード qwen36-27b --json
ax-engine ダウンロード gemma4-e2b --json
ax-engine ダウンロード gemma4-12b --json
ax-engine ダウンロード gemma4-31b --json
# 1 つのコマンドでダウンロードして提供する
ax-engine サーブ qwen36-35b --download --port 8080
# 生の mlx-community リポジトリ ID も受け入れられます
ax-engine ダウンロード mlx-community/Qwen3.6-35B-A3B-4bit --json
ax-engine ダウンロード mlx-community/gemma-4-e2b-it-4bit --json
# オプション: スナップショットを明示的なディレクトリにコピーします
ax-engine ダウンロード qwen36-35b --dest /ボリューム/モデル/qwen36-35b
# Python SDK
ax_engine インポート download_model から
path = download_model( " mlx-community/Qwen3.6-35B-A3B-4bit " )
組み込みのダウンロード エイリアス:
デフォルトでは、ダウンロードを Hugging Face Hub キャッシュに残しておきます。このキャッシュは mlx_lm および他の HF 対応ツールと共有され、大きなウェイトの重複コピーが回避されます。 --dest は、共有キャッシュの外部に明示的にコピーする場合にのみ使用します。
すでに mlx_lm がインストールされている場合、そのダウンロードは同じキャッシュに配置され、AX エンジンはそれらを自動検出できます。
python -m mlx_lm.generate --model mlx-community/Qwen3-4B-4bit --prompt " x " --max-tokens 1
ax-engine-bench 生成マニフェスト ~ /.cache/huggingface/hub/models

--mlx-community--Qwen3-4B-4bit/snapshots/ < ハッシュ >
ax-engineserve ~ /.cache/huggingface/hub/models--mlx-community--Qwen3-4B-4bit/snapshots/ < ハッシュ > --port 8080
生のハグフェイスチェックポイント
生のチェックポイントは、AX エンジンがロードする前にサニタイズする必要があります。
pip インストール mlx-lm
mlx_lm.convert --hf-path < org/model > --mlx-path /path/to/dest -q --q-bits 4
ax-engine-bench 生成マニフェスト /path/to/dest
ax-engine サーブ /path/to/dest --port 8080
マニフェストの生成
上記の両方のパスには、model-manifest.json が必要です。ダウンロード ヘルパーによって自動的に生成されます。直接実行するには:
ax-engine-bench generated-manifest /path/to/model # pip、Homebrew、またはビルドされたバイナリ
Cargo run -p ax-engine-core --bin generated-manifest -- /path/to/model # ソース
一般的なハードウェア
ローカル エージェントとチャットボットのワークロードの場合、AX Engine は、すべてのワークフローに対応する 1 つのモデルよりも小規模なモデル ポートフォリオに適しています。完全な推奨事項については、FAQ モデル スタック ガイダンスを参照してください。
AX エンジンは、ローカル推論作業に安定したランタイム コントラクトを提供します。
リポジトリ所有の MLX 実行は、委任されたルートとは別に直接サポート モデル ファミリを追跡します。委任された結果は、AX 所有のスループット要求ではありません。
デュアルファミリーの投機的デコードは、同じリポジトリ所有のランタイムおよびベンチマーク ツールで、Qwen3.6 の融合 MTP サイドカーと Gemma 4 の個別のアシスタントドラフト コントラクトの両方をサポートします。
N グラム アクセラレーションは、第 2 ドラフト モデルを使用しない場合でも、ヒット率の高いベンチマーク行で最大 3.1 倍の mlx_lm デコード スループットに達します。
長時間セッションのプレフィックスを再利用すると、検証済みのキャッシュ レイアウト上に物理 MLX KV スナップショットが復元されるため、長時間実行されるチャットとエージェントのループにより、蓄積されたコンテキストが繰り返し事前に入力されることが回避されます。 docs/LONG-CONTEXT.md を参照してください。
ワークロード コントラクト ツール ( ax-engine-bench ) は、正確性、決定論、ルート ID、およびチェックインされた m 全体にわたる回帰を検証します。

アニフェスト。
委任されたルート ( mlx_lm_delegated 、 llama_cpp ) は、AX が所有するパフォーマンス要求を損なうことなく、明示的な互換性のケースをカバーします。
mlx_lm は正規の MLX リファレンスです。 AX エンジンは、mlx_lm.benchmark と比較し、AX にリポジトリ所有のグラフがまだない場合、明示的な委任された互換性ルートとして mlx_lm.server を保持します。 MLX カーネルと AX 所有のランタイム動作の間の境界については、FAQ を参照してください。
設計の詳細: スケジューラ · KV キャッシュ · ロングコンテキスト · ベンチマーク設計。
パス
用途:
現在の範囲
リポジトリ所有の MLX ランタイム
直接サポートの MLX モデル ファミリと、ベンチマーク アーティファクトに裏付けられたリポジトリ所有のパフォーマンス主張
ローカル Apple Silicon 推論、トークンベースのサーバー/SDK リクエスト、ダイレクトおよび N-gram アクセラレーション モード
mlx_lm_delegated
AX がリポジトリ所有のグラフを持つ前に、アップストリームの mlx-lm がサポートしていた MLX テキスト モデル
ユーザー提供の mlx_lm.server によるブロッキングと SSE テキスト生成。 AX 所有のトークン/KV パフォーマンスではない
ラマ_cpp
GGUF および非 MLX ローカル推論
委任された llama.cpp サーバー/CLI の互換性。リポジトリ所有の MLX スループットではなく、ルート契約の証拠
ランタイム レポートは selected_backend 、 support_tier 、および solution_policy を公開するため、呼び出し元とベンチマーク アーティファクトはこれらのパスを区別できます。正確な OpenAI 形式のエンドポイント コントラクトについては、 docs/API-COMPATIBILITY.md を参照してください。
AX Engine の公開パフォーマンスに関する主張は、ルート ID、モデル アーティファクト、プロンプト スイート、サンプラー設定、およびリポジトリの来歴を保存するベンチマーク アーティファクトに範囲が定められています。
直接サポートとは、AX がモデル ファミリ用のリポジトリ所有の ax-engine-mlx グラフを持ち、AX マニフェスト パスを通じて MLX セーフテンソルをロードすることを意味します。他の MLX テキスト モデルは、引き続き明示的な mlx_lm_delegated 互換性ルートを使用できます。
GLM 4.7 Flash ( glm4_moe_lite ) は直接サポートから mlx_l に降格されました

m_delegated passby ルート: ネイティブ デコードは mlx_lm パリティにのみ到達し、4 ビット エクスポートには MTP ヘッドがありません。 glm4.7-flash-4bit プリセットは委任された層を選択するようになり、 --mlx-lm-server-url が必要になります。 docs/SUPPORTED-MODELS.md を参照してください。
新しいアーキテクチャを追加するということは、汎用ローダーを接続するのではなく、 ax-engine-mlx にモデル グラフを実装することを意味します。アーキテクチャ コードだけでは、直接サポートを主張することはできません。モデルには、ここでプロモーションする前に、リポジット所有のグラフ、マニフェスト、スモーク カバレッジ、およびベンチマークの証拠が必要です。 LLaMA、Mistral、Mixtral、DeepSeek、およびリストされていない Gemma/Qwen バリアントは、明示的な委任ルートを使用する必要があります。
別のアーキテクチャまたはチェックポイントをプロモートする前に、 scripts/probe_mlx_model_support.py --model-dir <model-dir> を実行します。モデルは、そのマニフェスト、ローカル参照ファイル、およびランタイム パスがすべて存在する場合にのみ、repo_owned_runtime_ready を報告する必要があります。
完全なリスト: docs/SUPPORTED-MODELS.md 。
完全な結果テーブルと解釈は docs/PERFORMANCE.md にあります。ベンチマーク方法、テスト設定、再現の詳細は docs/BENCHMARKS.md にあります。
Gemma 4 12B (model_type: gemma4_unified ) は、レイヤーごとの埋め込み E2B/E4B および MoE 26B/31B とは異なる実装です。アップストリームの mlx_lm 0.31.3 はロードできません —

[切り捨てられた]

## Original Extract

Apple Silicon LLM runtime supporting Gemma 4 and Qwen 3.6 MTP modes - defai-digital/ax-engine

GitHub - defai-digital/ax-engine: Apple Silicon LLM runtime supporting Gemma 4 and Qwen 3.6 MTP modes · GitHub
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
defai-digital
/
ax-engine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,576 Commits 1,576 Commits .cargo .cargo .github/ workflows .github/ workflows benchmarks benchmarks build/ metal build/ metal crates crates docs docs examples examples javascript/ ax-engine javascript/ ax-engine metal metal proto/ ax_engine/ v1 proto/ ax_engine/ v1 python python qa qa scripts scripts sdk sdk .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
AX Engine is a Mac-first LLM inference runtime, local server, SDK layer, and benchmark toolkit for Apple Silicon. It runs direct-support MLX model families natively, and routes other MLX text models or non-MLX models through explicit mlx-lm and llama.cpp compatibility routes.
AX Engine is for developers who want a local OpenAI-compatible model server on Apple Silicon without hiding which runtime path is doing the work.
OpenAI-compatible local text endpoints for common chat and completion flows, with SDKs for Python, TypeScript/JavaScript, Go, Ruby, and Mojo.
Repo-owned MLX runtime paths for direct-support Gemma and Qwen families, with delegated routes kept explicit.
Announcement-ready benchmark claims where evidence is complete: Gemma 4 12B assistant-MTP is 2.83-2.92x faster than same-artifact direct decode, and Qwen3.6 35B-A3B AX MTP is +76.4% faster than MTPLX on the public sidecar-fair matrix.
Workload-contract benchmark tooling records route identity, artifacts, prompt suite, sampler, cooldowns, accept rate, and dirty-state provenance.
Performance
Gemma 4 12B
Gemma 4 12B Multimodal
Speculative Decoding (MTP)
Gemma 4
Direct Decode · Prefill · TTFT
pip install " ax-engine[download] "
Download a small model and start the server:
MODEL_DIR= " $( ax-engine download mlx-community/Qwen3-4B-4bit --json | python3 -c ' import json,sys; print(json.load(sys.stdin)["dest"]) ' ) "
ax-engine serve " $MODEL_DIR " --port 8080
High-memory model shortcuts:
# Choose one:
ax-engine serve qwen36-35b --download --port 8080
ax-engine serve gemma4-12b --download --port 8080
Call it from any OpenAI client:
from openai import OpenAI
client = OpenAI ( base_url = "http://127.0.0.1:8080/v1" , api_key = "local" )
model = client . models . list (). data [ 0 ]. id
resp = client . chat . completions . create (
model = model ,
messages = [{ "role" : "user" , "content" : "What is AGI?" }],
max_tokens = 128 ,
)
print ( resp . choices [ 0 ]. message . content )
Or use the Python SDK directly:
from ax_engine import download_model , Session
path = download_model ( "mlx-community/Qwen3-4B-4bit" )
with Session ( mlx = True , mlx_model_artifacts_dir = str ( path )) as s :
print ( s . generate ([ 1 , 2 , 3 ], max_output_tokens = 8 ). output_tokens )
Quick Start requires macOS 14 (Sonoma) or later on Apple Silicon M2 Max or newer with 32 GB unified memory or more . Larger models such as Qwen3.6 35B-A3B and Gemma 4 12B need the memory tiers listed in Typical Hardware .
pip install ax-engine
Requires macOS 14+, Apple Silicon (M2 Max or newer), Python 3.10+. The pip wheel includes the ax-engine orchestration CLI and the ax-engine-server binary — both are on your PATH after install.
pip install " ax-engine[download] " # Hugging Face Hub downloads plus mlx-lm tools
Homebrew
For the same top-level ax-engine CLI plus the native ax-engine-server and
ax-engine-bench tools:
brew tap defai-digital/ax-engine
brew install ax-engine
ax-engine doctor
If doctor fails with Library not loaded: libmlxc.dylib , run:
brew install mlx-c && brew reinstall ax-engine .
brew install mlx-c
cargo build --workspace --release
maturin develop # Python bindings
Getting a Model
AX Engine requires pre-sanitized MLX weights. The recommended source is mlx-community — models there are already converted and validated.
ax-engine download , download_model() , and scripts/download_model.py download weights and auto-generate the required model-manifest.json in one step:
# List supported download targets
ax-engine download --list
# Download by alias
ax-engine download qwen36-35b --json
ax-engine download qwen36-27b --json
ax-engine download gemma4-e2b --json
ax-engine download gemma4-12b --json
ax-engine download gemma4-31b --json
# Download and serve in one command
ax-engine serve qwen36-35b --download --port 8080
# Raw mlx-community repo IDs are also accepted
ax-engine download mlx-community/Qwen3.6-35B-A3B-4bit --json
ax-engine download mlx-community/gemma-4-e2b-it-4bit --json
# Optional: copy snapshot to an explicit directory
ax-engine download qwen36-35b --dest /Volumes/Models/qwen36-35b
# Python SDK
from ax_engine import download_model
path = download_model( " mlx-community/Qwen3.6-35B-A3B-4bit " )
Built-in download aliases:
Leave downloads in the Hugging Face Hub cache by default — it's shared with mlx_lm and other HF-aware tools, avoiding duplicate copies of large weights. Use --dest only when you want an explicit copy outside the shared cache.
If you already have mlx_lm installed, its downloads land in the same cache and AX Engine can auto-discover them:
python -m mlx_lm.generate --model mlx-community/Qwen3-4B-4bit --prompt " x " --max-tokens 1
ax-engine-bench generate-manifest ~ /.cache/huggingface/hub/models--mlx-community--Qwen3-4B-4bit/snapshots/ < hash >
ax-engine serve ~ /.cache/huggingface/hub/models--mlx-community--Qwen3-4B-4bit/snapshots/ < hash > --port 8080
Raw HuggingFace checkpoint
Raw checkpoints need sanitization before AX Engine can load them:
pip install mlx-lm
mlx_lm.convert --hf-path < org/model > --mlx-path /path/to/dest -q --q-bits 4
ax-engine-bench generate-manifest /path/to/dest
ax-engine serve /path/to/dest --port 8080
Manifest generation
Both paths above require model-manifest.json . Download helpers generate it automatically. To run it directly:
ax-engine-bench generate-manifest /path/to/model # pip, Homebrew, or built binary
cargo run -p ax-engine-core --bin generate-manifest -- /path/to/model # source
Typical Hardware
For local agent and chatbot workloads, AX Engine is a better fit for a small model portfolio than for one model serving every workflow. See the FAQ model-stack guidance for the full recommendation.
AX Engine gives local inference work a stable runtime contract:
Repo-owned MLX execution tracks direct-support model families separately from delegated routes — delegated results are not AX-owned throughput claims.
Dual-family speculative decoding supports both Qwen3.6's fused MTP sidecar and Gemma 4's separate assistant-drafter contract in the same repo-owned runtime and benchmark tooling.
N-gram acceleration reaches up to 3.1× mlx_lm decode throughput on high-hit benchmark rows with no second draft model.
Long-session prefix reuse restores physical MLX KV snapshots on validated cache layouts, so long-running chat and agent loops avoid repeatedly pre-filling accumulated context. See docs/LONG-CONTEXT.md .
Workload-contract tooling ( ax-engine-bench ) validates correctness, determinism, route identity, and regression across checked-in manifests.
Delegated routes ( mlx_lm_delegated , llama_cpp ) cover explicit compatibility cases without polluting AX-owned performance claims.
mlx_lm is the canonical MLX reference. AX Engine compares against mlx_lm.benchmark and keeps mlx_lm.server as the explicit delegated compatibility route when AX does not yet have a repo-owned graph. See the FAQ for the boundary between MLX kernels and AX-owned runtime behavior.
Design details: Scheduler · KV Cache · Long Context · Benchmark Design .
Path
Use it for
Current scope
Repo-owned MLX runtime
Direct-support MLX model families and repo-owned performance claims backed by benchmark artifacts
Local Apple Silicon inference, token-based server/SDK requests, direct and n-gram acceleration modes
mlx_lm_delegated
MLX text models that upstream mlx-lm supports before AX has a repo-owned graph
Blocking and SSE text generation through a user-provided mlx_lm.server ; not AX-owned token/KV performance
llama_cpp
GGUF and non-MLX local inference
Delegated llama.cpp server/CLI compatibility; route-contract evidence, not repo-owned MLX throughput
The runtime report exposes selected_backend , support_tier , and resolution_policy so callers and benchmark artifacts can distinguish these paths. For the exact OpenAI-shaped endpoint contract see docs/API-COMPATIBILITY.md .
AX Engine's public performance claims are scoped to benchmark artifacts that preserve route identity, model artifacts, prompt suite, sampler settings, and repository provenance.
Direct support means AX has a repo-owned ax-engine-mlx graph for the model family and loads MLX safetensors through the AX manifest path. Other MLX text models can still use the explicit mlx_lm_delegated compatibility route.
GLM 4.7 Flash ( glm4_moe_lite ) was demoted from direct support to the mlx_lm_delegated passby route: native decode only reaches mlx_lm parity and the 4-bit export has no MTP head. The glm4.7-flash-4bit preset now selects the delegated tier and requires --mlx-lm-server-url . See docs/SUPPORTED-MODELS.md .
Adding a new architecture means implementing the model graph in ax-engine-mlx , not wiring up a generic loader. Architecture code alone is not a direct-support claim — a model requires a repo-owned graph, manifest, smoke coverage, and benchmark evidence before promotion here. LLaMA, Mistral, Mixtral, DeepSeek, and unlisted Gemma/Qwen variants should use the explicit delegated route.
Before promoting another architecture or checkpoint, run scripts/probe_mlx_model_support.py --model-dir <model-dir> ; a model should report repo_owned_runtime_ready only when its manifest, local reference files, and runtime path are all present.
Full list: docs/SUPPORTED-MODELS.md .
Full result tables and interpretation live in docs/PERFORMANCE.md . Benchmark methodology, test setup, and reproduction details live in docs/BENCHMARKS.md .
Gemma 4 12B ( model_type: gemma4_unified ) is a different implementation from the per-layer-embedding E2B/E4B and the MoE 26B/31B. Upstream mlx_lm 0.31.3 cannot load it —

[truncated]
