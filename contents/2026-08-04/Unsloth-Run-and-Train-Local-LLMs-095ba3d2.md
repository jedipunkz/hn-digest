---
source: "https://github.com/unslothai/unsloth"
hn_url: "https://news.ycombinator.com/item?id=49162853"
title: "Unsloth: Run and Train Local LLMs"
article_title: "GitHub - unslothai/unsloth: Unsloth is a local UI for training and running Kimi K3, Gemma 4, Qwen3.6, DeepSeek-V4, GLM and other models. · GitHub"
author: "ungreased0675"
captured_at: "2026-08-04T00:07:22Z"
capture_tool: "hn-digest"
hn_id: 49162853
score: 1
comments: 0
posted_at: "2026-08-03T23:54:50Z"
tags:
  - hacker-news
  - translated
---

# Unsloth: Run and Train Local LLMs

- HN: [49162853](https://news.ycombinator.com/item?id=49162853)
- Source: [github.com](https://github.com/unslothai/unsloth)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T23:54:50Z

## Translation

タイトル: Unsloth: ローカル LLM の実行とトレーニング
記事タイトル: GitHub - unslothai/unsloth: Unsloth は、 Kim K3、Gemma 4、Qwen3.6、DeepSeek-V4、GLM およびその他のモデルをトレーニングおよび実行するためのローカル UI です。 · GitHub
説明: Unsloth は、Kimi K3、Gemma 4、Qwen3.6、DeepSeek-V4、GLM およびその他のモデルをトレーニングおよび実行するためのローカル UI です。 - 怠惰ではない/怠惰ではない

記事本文:
GitHub - unslothai/unsloth: Unsloth は、 Kim K3、Gemma 4、Qwen3.6、DeepSeek-V4、GLM およびその他のモデルをトレーニングおよび実行するためのローカル UI です。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
ああ、ああ！
読み込み中にエラーが発生しました。これをリロードしてください

ページに
だらしない
/
怠ける
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6,620 コミット 6,620 コミット .github .github 画像 画像 スクリプト スクリプト スタジオ スタジオ テスト テスト unsloth unsloth unsloth_cli unsloth_cli .gitattributes .gitattributes .gitignore .gitignore .pre-commit-ci.yaml .pre-commit-ci.yaml .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md コピー中 コピー中 ライセンス ライセンス マニフェスト.in MANIFEST.in README.md README.md _changelog_build.py _changelog_build.py build.sh build.sh cli.py cli.py install.ps1 install.ps1 install.sh install.sh pyproject.toml pyproject.toml unsloth-cli.py unsloth-cli.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Unsloth Studio では、モデルをローカルで実行およびトレーニングできます。
特徴 •
ニュース •
クイックスタート •
ノートブック •
ドキュメント
macOS、Linux、WSL:
カール -fsSL https://unsloth.ai/install.sh |しー
Windows:
irm https://unsloth.ai/install.ps1 |アイエックス
コミュニティ:
不和
Unsloth Studio (ベータ) を使用すると、Windows、Linux、macOS 上でテキスト、オーディオ、埋め込み、ビジョン モデルを実行してトレーニングできます。
GGUF、LoRA アダプター、セーフテンサーを含むモデルの検索 + ダウンロード + 実行
モデルのエクスポート : モデルを GGUF、16 ビット セーフテンサー、およびその他の形式で保存またはエクスポートします。
ツール呼び出し: 自己修復ツール呼び出しと Web 検索のサポート
コードの実行 : LLM がクロード アーティファクトおよびサンドボックス環境でコードをテストできるようにします
API 推論エンドポイント: Claude Code、Unsloth を使用した Codex ツールでローカル LLM をデプロイして実行する
推論設定を自動設定し、チャット テンプレートをカスタマイズします。
私たちは直接仕事をします

gpt-oss 、 Qwen3 、 Llama 4 、 Mistral 、 Gemma 1-3 、および Phi-4 のチームと協力し、モデルの精度を向上させるバグを修正しました。
画像、音声、PDF、コード、DOCX などを使用してチャットできます。 API プロバイダー (OpenAI、Anthropic) またはサーバー (vLLM、Ollama) に接続します。
同じプロンプトで 2 つのモデルを並べて比較します。
OpenAI/Anthropic 互換 API : /v1/chat/completions 、 /v1/responses 、および /v1/messages を通じてローカル モデルを提供します。
ローカル モデルをエージェントに接続する: Claude Code、Codex、Hermes などで unsloth start を使用します。
Web/PDF 検索では、PDF 論文、マニュアル、その他の PDF 結果を読むことができます。
GGUF ハードウェア コントロール: GPU/レイヤーの選択、MoE エキスパートのオフロード、マルチ GPU または Tensor 並列処理の使用。
オプトイン MCP 制御エンドポイントにより、AI クライアントはモデル、トレーニング、レシピ、エクスポートを管理できるようになります。
70% 少ない VRAM で最大 2 倍の速度でトレーニングおよび RL 500+ モデルを実行します。 MoE は最大 12 倍高速です。
Windows、WSL、Linux 全体で AMD GPU 上で RL をトレーニングして実行します。
データ レシピ : PDF、CSV、DOCX などからデータセットを自動作成します。ビジュアルノード ワークフローでデータを編集します。
強化学習では、GRPO、FP8、ビジョン RL で使用する VRAM が 80% 削減され、コンテキストが 7 倍長くなります。
ロングコンテキストトレーニング: 3 倍高速、30% 少ない VRAM、500K 以上のコンテキスト。
LoRA/QLoRA、フル微調整、RL、事前トレーニング、4 ビット、16 ビット、FP8 をサポートします。
PyTorch と Hugging Face で構築されたカスタム Triton と数学的カーネル。
可観測性 : トレーニングをライブで監視し、損失と GPU の使用状況を追跡し、グラフをカスタマイズします。
マルチ GPU トレーニングがサポートされており、大きな改善が近々行われる予定です。
Unsloth Start は、1 つのコマンドで Claude Code 、Codex およびその他のエージェントをローカル モデルに接続します。
Unsloth を起動し、モデルをロードし、プロジェクト フォルダーを開いて、次を実行します。
アンナマケモノスタートクロード
clude をサポートされているエージェントに置き換えます。
Claude Code、Codex、OpenCode は現在のモデルを維持し、Uns を使用できます。

地元の人として
副代理人:
unsloth start claude --as-subagent --model unsloth/model-GGUF:quant
📥インストール
Unsloth は、Web UI である Unsloth Studio を使用するか、コードベースのバージョンである Unsloth Core を使用する 2 つの方法で使用できます。それぞれに異なる要件があります。
Unsloth Studio (ベータ) は、Windows、Linux、WSL、macOS で動作します。
CPU: 現在、チャットおよびデータ レシピでサポートされています
NVIDIA: トレーニングは RTX 30/40/50、Blackwell、DGX Spark、Station などで動作します
macOS: トレーニング、MLX、GGUF 推論はすべてサポートされています。
AMD: トレーニング、RL、チャット、導入は Windows、WSL、Linux 上で動作します。 AMD ガイドを読んでください。
Vulkan: GGUF 推論は、Intel GPU を含む互換性のある GPU でサポートされています。 Vulkan は GGUF 推論のみを高速化します。トレーニングには、サポートされている PyTorch または MLX バックエンドが必要です。
マルチ GPU: 現在利用可能ですが、メジャー アップグレードが予定されています
カール -fsSL https://unsloth.ai/install.sh |しー
同じコマンドを使用して更新します。
Vulkan llama.cpp バックエンドを強制するには、 をインストールまたは更新する前に UNSLOTH_FORCE_VULKAN=1 を設定します。この設定では llama.cpp バイナリ バンドルが選択されるため、Studio の起動時にのみ設定しても既存の CPU バンドルを置き換えることはできません。
UNSLOTH_FORCE_VULKAN=1 をエクスポート
カール -fsSL https://unsloth.ai/install.sh |しー
Windows:
irm https://unsloth.ai/install.ps1 |アイエックス
同じコマンドを使用して更新します。
Vulkan llama.cpp バックエンドを強制するには、インストーラーまたはアップデーターを実行する前に環境変数を設定します。
$ 環境: UNSLOTH_FORCE_VULKAN = 1
irm https://unsloth.ai/install.ps1 |アイエックス
バックエンドが異なる場合、現在のインストーラーを再実行すると、以前に選択した CPU バンドルが置き換えられます。個別の Vulkan SDK は必要ありません。 GPU ドライバーは、動作する Vulkan ランタイムを提供する必要があります。
unsloth スタジオ -p 8888
LAN またはクラウド アクセスの場合は、-H 0.0.0.0 (生ポートのみ。パブリック URL の場合は --cloudflare を追加します)

）。デフォルトでは、Unsloth はローカルでのみアクセス可能です。
HTTPS 経由で Unsloth にアクセスするには、 unsloth studio --secure を使用します。 Unsloth は localhost にバインドされたままで、パブリック https://*.trycloudflare.com URL で公開する無料の Cloudflare トンネル経由でのみ到達します (トンネルを開始できない場合は失敗して閉じられるため、生のポートは公開されません)。これにより、インターネットから Unsloth にアクセスできるようになり、リンクと API キーを持っている人は誰でも Unsloth を使用してコードを実行できるようになります。API キーは非公開にしておいてください (下記のリモート アクセスを参照)。
Docker イメージの unsloth/unsloth コンテナを使用します。実行:
docker run -d -e JUPYTER_PASSWORD= " mypassword " \
-p 8888:8888 -p 8000:8000 -p 2222:22 \
-v $( pwd ) /work:/workspace/work \
--GPU すべて\
怠ける/怠ける
開発者、夜間、アンインストール
開発者向け、夜間インストール、アンインストールなどの手順については、「高度なインストール」を参照してください。
Linux、WSL:
カール -LsSf https://astral.sh/uv/install.sh |しー
uv venv unsloth_env --python 3.13
ソース unsloth_env/bin/activate
uv pip install unsloth --torch-backend=auto
Windows:
winget install - e -- id Python.Python。 3.13
winget インストール -- id = astral - sh.uv - e
uv venv unsloth_env -- Python 3.13
.\unsloth_env\Scripts\activate
uv pip install unsloth -- torch - backend = auto
Windows の場合、pip install unsloth は PyTorch がインストールされている場合にのみ機能します。 Windows ガイドをお読みください。
Unsloth Studio と同じ Docker イメージを使用できます。
RTX 50x、B200、6000 GPU の場合: uv pip install unsloth --torch-backend=auto 。 Blackwell と DGX Spark のガイドをお読みください。
AMD および Intel GPU に Unsloth をインストールするには、 AMD ガイド および Intel ガイド に従ってください。
ノートブックを使って無料でトレーニングしましょう。新しい無料の Unsloth Studio ノートブックを使用して、Web UI でモデルを無料で実行およびトレーニングできます。
ガイドをお読みください。データセットを追加して実行し、トレーニングされたモデルをデプロイします。
Kaggle のすべてのノートブックを参照してください。

、GRPO、TTS、エンベディングおよびビジョン
すべてのモデルとすべてのノートブックをご覧ください
ここでUnslothの詳細なドキュメントを参照してください
AMD トレーニング : Windows、WSL、Linux 上の AMD GPU でトレーニング、RL の実行、チャット、デプロイを行います。ガイド
GGUF ハードウェア コントロール: GPU/レイヤーの配置を選択し、MoE エキスパートをオフロードし、マルチ GPU または Tensor 並列処理を使用します。 #6414
あらゆるエージェントのローカル モデル : Unsloth の OpenAI および Anthropic 互換 API を通じて、Claude Code、Codex、Hermes、OpenCode、OpenClaw などから始めて unsloth を使用します。ガイド
MCP コントロール エンドポイント : 互換性のあるクライアントがモデル、トレーニング、レシピ、チェックポイント、エクスポートを管理できるようにします。 #7191
ローカル推論の信頼性 : 長時間のチャットをより速く再開し、停止したダウンロードを回復し、既存の GGUF ファイルを再利用します。 #7204 • #6858 • #7209
新しいモデル : Qwen-AgentWorld 、Ornith 、Kimi K2.7 Code および MiniMax M3
GLM-5.2 : Unsloth Dynamic GGUF を使用して、Z.ai の 744B パラメーター、1M コンテキストのオープン モデルをローカルで実行します。ガイド
DeepSeek-V4 : マルチターンおよびツール呼び出しの動作を修正して、DeepSeek-V4-Flash をローカルで実行します。ガイド
DiffusionGemma : Unsloth Studio で 1.8 倍高速な推論で Google の拡散言語モデルを実行し、微調整します。ガイド
Qwen3.6 : MTP を使用して Qwen3.6 を実行およびトレーニングし、サポートされている GPU の 1.4 ～ 2.2 倍高速な推論と NVFP4 クォントを実現します。ガイド
Gemma 4 : QAT、MTP、GGUF、MLX をサポートする Gemma 4 のテキスト、画像、音声モデルを実行およびトレーニングします。ガイド
MCP サーバー : モデル コンテキスト プロトコルを介して、ローカル モデルをファイル、アプリ、データベース、外部ツールに接続します。ガイド
接続 : ローカル モデルと API プロバイダー (OpenAI、Anthropic) またはサーバー (vLLM、Ollama) を同じインターフェイス内で混在させます。ガイド
Unsloth Studio の紹介: LLM を実行およびトレーニングするための新しい Web UI。ブログ
DeepSeek、GLM、Qwen、gpt-oss など、MoE LLM を 35% 少ない VRAM で 12 倍高速にトレーニングします。ブログ
埋め込みモデル: Unsl

oth は、約 1.8 ～ 3.3 倍高速な埋め込み微調整をサポートするようになりました。ブログ • ノートブック
新しいバッチアルゴリズムにより、他のすべてのセットアップと比較して 7 倍長い新しいコンテキスト RL。ブログ
新しい RoPE & MLP Triton カーネル & パディング フリー + パッキング: トレーニングが 3 倍高速になり、VRAM が 30% 削減されます。ブログ
500K コンテキスト: 500K を超えるコンテキストで 20B モデルをトレーニングすることが、80GB GPU で可能になりました。ブログ
FP8 および Vision RL : コンシューマー GPU で FP8 および VLM GRPO を実行できるようになりました。 FP8 ブログ • ビジョン RL
以下の高度な手順は Unsloth Studio 用です。 Unsloth Core の高度なインストールについては、ドキュメントを参照してください。
開発者は、最新 (夜間) ソースであるメイン ブランチからビルドをインストールします。
git clone https://github.com/unslothai/unsloth
CD ナマケモノ
./install.sh --local
unsloth スタジオ -p 8888
隔離された場所 (独自の仮想環境、 auth/ 、 studio.db 、キャッシュ、および llama.cpp ビルド) にインストールするには、 UNSLOTH_STUDIO_HOME を設定し、起動時に再度渡します。
UNSLOTH_STUDIO_HOME= " $PWD /.studio " ./install.sh --local
UNSLOTH_STUDIO_HOME= " $PWD /.studio " unsloth スタジオ -p 8888
次に更新するには:
cd unsloth && git pull
./install.sh --local
unsloth スタジオ -p 8888
開発者 / ナイトリー / 実験的インストール: Windows PowerShell:
開発者インストールはメイン ブランチからビルドします。

[切り捨てられた]

## Original Extract

Unsloth is a local UI for training and running Kimi K3, Gemma 4, Qwen3.6, DeepSeek-V4, GLM and other models. - unslothai/unsloth

GitHub - unslothai/unsloth: Unsloth is a local UI for training and running Kimi K3, Gemma 4, Qwen3.6, DeepSeek-V4, GLM and other models. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Uh oh!
There was an error while loading. Please reload this page .
unslothai
/
unsloth
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6,620 Commits 6,620 Commits .github .github images images scripts scripts studio studio tests tests unsloth unsloth unsloth_cli unsloth_cli .gitattributes .gitattributes .gitignore .gitignore .pre-commit-ci.yaml .pre-commit-ci.yaml .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md COPYING COPYING LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md _changelog_build.py _changelog_build.py build.sh build.sh cli.py cli.py install.ps1 install.ps1 install.sh install.sh pyproject.toml pyproject.toml unsloth-cli.py unsloth-cli.py View all files Repository files navigation
Unsloth Studio lets you run and train models locally.
Features •
News •
Quickstart •
Notebooks •
Documentation
macOS, Linux, WSL:
curl -fsSL https://unsloth.ai/install.sh | sh
Windows:
irm https: // unsloth.ai / install.ps1 | iex
Community:
Discord
Unsloth Studio (Beta) lets you run and train text, audio , embedding , vision models on Windows, Linux and macOS.
Search + download + run models including GGUF, LoRA adapters, safetensors
Export models : Save or export models to GGUF, 16-bit safetensors and other formats.
Tool calling : Support for self-healing tool calling and web search
Code execution : lets LLMs test code in Claude artifacts and sandbox environments
API inference endpoint : Deploy and run local LLMs in Claude Code, Codex tools with Unsloth
Auto set inference settings and customize chat templates.
We work directly with teams behind gpt-oss , Qwen3 , Llama 4 , Mistral , Gemma 1-3 , and Phi-4 , where we’ve fixed bugs that improve model accuracy.
Chat with images, audio, PDFs, code, DOCX and more. Connect API providers (OpenAI, Anthropic) or servers (vLLM, Ollama).
Compare any two models side by side with the same prompt.
OpenAI/Anthropic-compatible APIs : Serve local models through /v1/chat/completions , /v1/responses and /v1/messages .
Connect local models to agents : Use unsloth start with Claude Code, Codex, Hermes and more.
Web/PDF search can read PDF papers, manuals and other PDF results.
GGUF hardware controls : Choose GPUs/layers, offload MoE experts, use multi-GPU or Tensor Parallelism.
The opt-in MCP control endpoint lets AI clients manage models, training, recipes and exports.
Train and RL 500+ models up to 2x faster with 70% less VRAM ; MoE up to 12x faster .
Train and run RL on AMD GPUs across Windows, WSL and Linux.
Data Recipes : Auto-create datasets from PDF, CSV, DOCX etc. Edit data in a visual-node workflow.
Reinforcement Learning uses 80% less VRAM for GRPO, FP8 and vision RL, with 7x longer contexts.
Long-context training : 3x faster , 30% less VRAM and 500K+ context.
Supports LoRA/QLoRA, full fine-tuning, RL, pretraining, 4-bit, 16-bit and FP8.
Custom Triton and mathematical kernels built with PyTorch and Hugging Face.
Observability : Monitor training live, track loss and GPU usage and customize graphs.
Multi-GPU training is supported, with major improvements coming soon.
Unsloth Start connects Claude Code , Codex and other agents to local models with one command.
Start Unsloth, load a model, open your project folder, then run:
unsloth start claude
Replace claude with any supported agent:
Claude Code, Codex and OpenCode can keep their current model and use Unsloth as a local
subagent:
unsloth start claude --as-subagent --model unsloth/model-GGUF:quant
📥 Install
Unsloth can be used in two ways: through Unsloth Studio , the web UI, or through Unsloth Core , the code-based version. Each has different requirements.
Unsloth Studio (Beta) works on Windows, Linux, WSL and macOS .
CPU: Supported for Chat and Data Recipes currently
NVIDIA: Training works on RTX 30/40/50, Blackwell, DGX Spark, Station and more
macOS: Training, MLX and GGUF inference are ALL supported.
AMD: Training, RL, chat and deployment work on Windows, WSL and Linux. Read the AMD guide .
Vulkan: GGUF inference is supported on compatible GPUs, including Intel GPUs . Vulkan accelerates GGUF inference only; training still requires a supported PyTorch or MLX backend.
Multi-GPU: Available now, with a major upgrade on the way
curl -fsSL https://unsloth.ai/install.sh | sh
Use the same command to update.
To force the Vulkan llama.cpp backend, set UNSLOTH_FORCE_VULKAN=1 before installing or updating . The setting selects the llama.cpp binary bundle, so setting it only when launching Studio cannot replace an existing CPU bundle:
export UNSLOTH_FORCE_VULKAN=1
curl -fsSL https://unsloth.ai/install.sh | sh
Windows:
irm https: // unsloth.ai / install.ps1 | iex
Use the same command to update.
To force the Vulkan llama.cpp backend, set the environment variable before running the installer or updater:
$ env: UNSLOTH_FORCE_VULKAN = 1
irm https: // unsloth.ai / install.ps1 | iex
Re-running the current installer replaces a previously selected CPU bundle when the backend differs. A separate Vulkan SDK is not required; the GPU driver must provide a working Vulkan runtime.
unsloth studio -p 8888
For LAN or cloud access, add -H 0.0.0.0 (raw port only; add --cloudflare for a public URL). By default, Unsloth is accessible only locally.
To reach Unsloth over HTTPS, use unsloth studio --secure . Unsloth stays bound to localhost and is reached only through a free Cloudflare tunnel, which publishes it at a public https://*.trycloudflare.com URL (it fails closed if the tunnel can't start, so the raw port is never exposed). This makes Unsloth reachable from the internet, so anyone with the link and API key can use it and run code: keep your API key private (see Remote access below).
Use our Docker image unsloth/unsloth container. Run:
docker run -d -e JUPYTER_PASSWORD= " mypassword " \
-p 8888:8888 -p 8000:8000 -p 2222:22 \
-v $( pwd ) /work:/workspace/work \
--gpus all \
unsloth/unsloth
Developer, Nightly, Uninstall
To see developer, nightly and uninstallation etc. instructions, see advanced installation .
Linux, WSL:
curl -LsSf https://astral.sh/uv/install.sh | sh
uv venv unsloth_env --python 3.13
source unsloth_env/bin/activate
uv pip install unsloth --torch-backend=auto
Windows:
winget install - e -- id Python.Python. 3.13
winget install -- id = astral - sh.uv - e
uv venv unsloth_env -- python 3.13
.\unsloth_env\Scripts\activate
uv pip install unsloth -- torch - backend = auto
For Windows, pip install unsloth works only if you have PyTorch installed. Read our Windows Guide .
You can use the same Docker image as Unsloth Studio.
For RTX 50x, B200, 6000 GPUs: uv pip install unsloth --torch-backend=auto . Read our guides for: Blackwell and DGX Spark .
To install Unsloth on AMD and Intel GPUs, follow our AMD Guide and Intel Guide .
Train for free with our notebooks. You can use our new free Unsloth Studio notebook to run and train models for free in a web UI.
Read our guide . Add dataset, run, then deploy your trained model.
See all our notebooks for: Kaggle , GRPO , TTS , embedding & Vision
See all our models and all our notebooks
See detailed documentation for Unsloth here
AMD training : Train, run RL, chat and deploy on AMD GPUs across Windows, WSL and Linux. Guide
GGUF hardware controls : Choose GPU/layer placement, offload MoE experts and use multi-GPU or Tensor Parallelism. #6414
Local models for any agent : Use unsloth start with Claude Code, Codex, Hermes, OpenCode, OpenClaw and more through Unsloth's OpenAI- and Anthropic-compatible APIs. Guide
MCP control endpoint : Let compatible clients manage models, training, recipes, checkpoints and exports. #7191
Local inference reliability : Resume long chats faster, recover stalled downloads and reuse existing GGUF files. #7204 • #6858 • #7209
New models : Qwen-AgentWorld , Ornith , Kimi K2.7 Code and MiniMax M3
GLM-5.2 : Run Z.ai's 744B-parameter, 1M-context open model locally with Unsloth Dynamic GGUFs. Guide
DeepSeek-V4 : Run DeepSeek-V4-Flash locally with corrected multi-turn and tool-calling behavior. Guide
DiffusionGemma : Run and fine-tune Google's diffusion language model with 1.8x faster inference in Unsloth Studio. Guide
Qwen3.6 : Run and train Qwen3.6 with MTP for 1.4-2.2x faster inference and NVFP4 quants for supported GPUs. Guide
Gemma 4 : Run and train Gemma 4 text, image and audio models with QAT, MTP, GGUF and MLX support. Guide
MCP servers : Connect local models to files, apps, databases and external tools through Model Context Protocol. Guide
Connections : Mix local models with API providers (OpenAI, Anthropic) or servers (vLLM, Ollama) in the same interface. Guide
Introducing Unsloth Studio : our new web UI for running and training LLMs. Blog
Train MoE LLMs 12x faster with 35% less VRAM - DeepSeek, GLM, Qwen and gpt-oss. Blog
Embedding models : Unsloth now supports ~1.8-3.3x faster embedding fine-tuning. Blog • Notebooks
New 7x longer context RL vs. all other setups, via our new batching algorithms. Blog
New RoPE & MLP Triton Kernels & Padding Free + Packing : 3x faster training & 30% less VRAM. Blog
500K Context : Training a 20B model with >500K context is now possible on an 80GB GPU. Blog
FP8 & Vision RL : You can now do FP8 & VLM GRPO on consumer GPUs. FP8 Blog • Vision RL
The below advanced instructions are for Unsloth Studio. For Unsloth Core advanced installation, view our docs .
The developer install builds from the main branch, which is the latest (nightly) source.
git clone https://github.com/unslothai/unsloth
cd unsloth
./install.sh --local
unsloth studio -p 8888
To install into an isolated location (its own virtual env, auth/ , studio.db , cache and llama.cpp build), set UNSLOTH_STUDIO_HOME and pass it again at launch:
UNSLOTH_STUDIO_HOME= " $PWD /.studio " ./install.sh --local
UNSLOTH_STUDIO_HOME= " $PWD /.studio " unsloth studio -p 8888
Then to update :
cd unsloth && git pull
./install.sh --local
unsloth studio -p 8888
Developer / Nightly / Experimental installs: Windows PowerShell:
The developer install builds from the main branch, which is th

[truncated]
