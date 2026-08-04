---
source: "https://github.com/bublhub/BubbleHub"
hn_url: "https://news.ycombinator.com/item?id=49170566"
title: "Show HN: BubbleHub - a local runtime and hosting for LLM agents"
article_title: "GitHub - bublhub/BubbleHub: BubbleHub is a runtime control layer for AI agents with sandboxing and local execution. · GitHub"
author: "danielsbubblee"
captured_at: "2026-08-04T16:07:56Z"
capture_tool: "hn-digest"
hn_id: 49170566
score: 1
comments: 0
posted_at: "2026-08-04T15:45:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: BubbleHub - a local runtime and hosting for LLM agents

- HN: [49170566](https://news.ycombinator.com/item?id=49170566)
- Source: [github.com](https://github.com/bublhub/BubbleHub)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T15:45:12Z

## Translation

タイトル: 表示 HN: BubbleHub - LLM エージェントのローカル ランタイムおよびホスティング
記事タイトル: GitHub - bublhub/BubbleHub: BubbleHub は、サンドボックス化とローカル実行を備えた AI エージェントのランタイム制御レイヤーです。 · GitHub
説明: BubbleHub は、サンドボックス化とローカル実行を備えた AI エージェントのランタイム制御レイヤーです。 - バブルハブ/バブルハブ

記事本文:
GitHub - bublhub/BubbleHub: BubbleHub は、サンドボックス化とローカル実行を備えた AI エージェントのランタイム制御レイヤーです。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
バブルハブ
/
バブルハブ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
87 コミット 87 コミット .bubblehub-cache .bubblehub-cache .cursor/ スキル .cursor/ スキル .github .github app app a

ssets 資産 bubblehub bubblehub docker docker docs docs 例 例 libbubble libbubble スクリプト スクリプト テスト テスト .clang-format .clang-format .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスREADME.md README.md codecov.yml codecov.yml package.json package.json pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1 つのコマンドでローカル LLM サービスとサンドボックス エージェントを実行します。
カール -fsSL https://bubblehub.ai/install.sh |バッシュ
Windows PowerShell インストーラー:
irm https:// bubblehub.ai / install.ps1 |アイエックス
Windows インストーラーをダブルクリックします。
irm https:// bubblehub.ai / ダウンロード / 最新 / BubbleHub-0.2.0-x64.exe - OutFile BubbleHub-0.2.0-x64.exe
.\ BubbleHub-0.2.0-x64.exe
確認してください:
docker pull ghcr.io/bublhub/bubblehub:latest
リリース イメージをベースとして使用します。
ghcr.io/bublhub/bubblehub から:v0.2.0
クイックスタート
ローカル モデルに質問します。
バブルプロンプト --text " BubbleHub からこんにちは "
サンドボックスでエージェントを実行します。
バブルラン --root-dir ./examples/basic --binary ./examples/basic/basic_agent.py --memory 16G
bubble ps 、シェル プロンプト、および BubbleHub デスクトップ アプリのエージェントに名前を付けます。
バブルシェル --name reviewer --root-dir ./workspace
バブル ps --kill agt-...
OpenAI 互換のローカル エンドポイントを開始します (オプション)。
バブルサーブ
モデルを選択または検査します。
バブルモデル
バブルモデルリスト
バブルモデルは停止する
アプリを開く:
OpenAI 互換エンドポイントを http://127.0.0.1:8000/v1 で公開します。
ウォーム モデル バックエンドをエージェント間で共有し続けます。
ファイルシステムとネットワークへのアクセスが制限された Linux サンドボックスでエージェントを実行します。
ローカル推論を OPENAI_BASE_URL および OPENAI_API_KEY としてエージェントに挿入します。
グラフィカルなモニタリング、マニフェストのレビュー、ベースモデル用の BubbleHub デスクトップ アプリを提供します

選択。
バブル実行は、エージェントを起動する前に共有推論エンドポイントを開始し、以下を注入します。
OPENAI_BASE_URL=http://127.0.0.1:8000/v1
OPENAI_API_KEY=バブルハブローカル
BUBBLEHUB_API_BASE_URL=http://127.0.0.1:8000
BUBBLEHUB_SANDBOX_INFERENCE_HOST=127.0.0.1
BUBBLEHUB_SANDBOX_INFERENCE_PORT=8000
サンドボックス化されたエージェントは、デフォルトではローカル推論エンドポイントにのみアクセスできます。 HTTP クライアントは http://127.0.0.1:18080 にあるループバック ポリシー プロキシを参照し、BubbleHub は分離されたサンドボックスに HTTP_PROXY 、 HTTPS_PROXY 、 http_proxy 、および https_proxy を挿入します。プロキシは libbubble によって所有されます。プロキシは、 ~/.local/state/bubblehub/sandboxes/<agent-id>/access-manifest.json (または BUBBLEHUB_STATE_DIR ) の下にある永続的なサンドボックス アクセス マニフェストをチェックし、許可されたリクエストを転送するか、ログに記録された 403 を返します。不明なリクエストはフェールクローズされます。使用可能なホスト プロンプトがない場合、BubbleHub は後でダッシュボードを確認できるように、それらをマニフェストに保留中として記録します。エージェントのセットアップ手順で一般的な送信ネットワーク アクセスが必要な場合は、バブル ランまたはバブル シェルで --allow-network を使用します。
バブル ランまたはバブル シェルが実際の端末に接続されている場合、新しいホストに最初にアクセスすると、エージェントが一時停止され、ホスト上で always 、never 、または ask each time (approve now) を求めるプロンプトが表示されます。非インタラクティブな実行はフェールクローズされ、バブル ダッシュボードを実行するためのリマインダーが出力されます。ダッシュボードは、ライブ リソース ビューが開く前に、保留中のサンドボックス アクセス リクエストを解決します。永続化されたポリシーを検査および編集するには、バブル マニフェスト --root-dir <dir> またはバブル マニフェスト --agent-id <エージェント> を使用します。マニフェスト ポリシーは正確に always 、never、または ask です。 HTTP ポリシーはドメインと表示される HTTP メソッド/パスを照合しますが、HTTPS CONNECT はホスト/ポート レベルでのみ照合できます。
--root-dir が指定されている場合、非システム バイナリはそのルートと BubbleHub マウント内に存在する必要があります。

ルートをサンドボックス ワークスペースとして使用します。 /usr 、 /bin 、 /sbin 、または /opt/bubblehub のシステム バイナリは、引き続きルート ディレクトリで使用できます。これにより、バブル シェル --root-dir <dir> はワークスペース サンドボックス内でシェルを開くことができます。 --root-dir を省略すると、サンドボックスが開始される前に、非システム バイナリが一時ワークスペースにコピーされます。サンドボックス内では、BubbleHub Python プロンプト/shim 呼び出しによって BUBBLEHUB_SANDBOX=1 が検出され、ネイティブ共有ライブラリをロードする代わりに、転送された推論エンドポイントが使用されます。
インストールされたサンドボックスは、エージェントごとのオーバーレイを使用して Ubuntu 26.04 ルート ファイルシステム上で実行されます。 Ubuntu の下位ファイルシステムは変更されません。ワークスペースの外に書き込むと、.bubblehub/agents/<agent-id>/overlay/upper にコピーされ、そのエージェントで永続化されます。 --force-new-sandbox または --overwrite-sandbox を使用して、現在のワークスペースの永続的なホームとプライベート オーバーレイを破棄します。
実装の詳細、セキュリティの前提条件、既知のギャップについては、 docs/sandbox.md を参照してください。
ステップバイステップのガイドについては、 docs/getting_started.md および docs/running_an_agent.md を参照してください。
OpenClaw はサンドボックス内から完全にインストールできます。永続エージェント ホームは、実行全体にわたって nvm 、 npm グローバル パッケージ、および OpenClaw 構成を保持します。
バブルシェル --allow-network --root-dir openclaw
サンドボックス シェル内:
カール -fsSL https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.5/install.sh |バッシュ
エクスポート NVM_DIR= " $HOME /.nvm "
。 「 $NVM_DIR /nvm.sh 」
nvm インストール 22.19.0
npm install -g openclaw@latest
openclaw オンボード --install-daemon
リリース
BubbleHub は、ソース インストール アセットを GitHub リリースから提供し、ランタイム イメージを GHCR から提供します。
v* タグをプッシュします。リリース ワークフローでは、単体テストを実行し、ローカル推論統合テストを実行してから、以下を公開します。
CI はプル リクエスト用の合成リリース アセットも構築し、パラレルでのインストールを検証します

ブランチがリリース準備ができているとみなされる前の el:
Linuxカール |クリーンな Ubuntu コンテナからの bash インストール
別のクリーンな Ubuntu コンテナから Linux apt install ./BubbleHub-<version>-x64.deb
セルフホスト型 Windows デスクトップ ランナーへの Windows PowerShell のインストール
Windows .exe ブートストラップを別の自己ホスト型 Windows デスクトップ ランナーにインストールする
各スモーク テストでは、以前の合成バージョンをインストールし、 bubble --version 、ショートカットの作成、基本的な CLI 動作、およびコントロール センターの /health 応答を検証し、現在の合成バージョンを上書きインストールして、ランタイムが更新されたバージョンを報告することを確認します。 Windows PowerShell パスと .exe パスは両方とも、リリースされた .deb を WSL にインストールし、Windows ランチャーを作成します。 Windows リリース インストール ジョブには、対話型デスクトップ、WSL2/WSLg、Python、クリーンな Ubuntu ベース ディストリビューション、および使い捨て WSL ディストリビューションをインポートする権限を備えた [セルフホスト型、Windows、X64] ランナーが必要です。
Cursor が作成したリリースノートの場合は、タグを付ける前に Cursor に BubbleHub リリースノート スキルを使用するように依頼してください。
以前のリリース以降のコミットから .github/releases/<tag>.md が書き込まれ、リリース ワークフローはそのファイルが存在する場合に使用します。
カール -fsSL https://bubblehub.ai/download/linux/v0.2.0/install.sh | BUBBLEHUB_VERSION=v0.2.0 bash
特定の Debian パッケージをダウンロードします。
カール -LO https://bubblehub.ai/download/linux/v0.2.0/BubbleHub-0.2.0-x64.deb
sudo apt install ./BubbleHub-0.2.0-x64.deb
一致するランタイム イメージを使用します。
ドッカープル ghcr.io/bublhub/bubblehub:v0.2.0
ソースからビルドする
仮想環境を作成してアクティブ化する
python3 -m venv .venv
ソース .venv/bin/activate
依存関係をインストールし、BubbleHub を構築します。このビルドでは、サンドボックスのセットアップに必要な権限を持つネイティブの bubblehub-sandbox ヘルパーを /usr/local/bin にインストールし、最初のインスタンスに Ubuntu 26.04 rootfs を作成します。

すべてを保存し、開発ループを高速化するために、後でローカルで再構築するときにそれを保存します。
./scripts/install-deps.sh
./scripts/build.sh
プッシュ前のテスト
プッシュする前に、CI が使用するのと同じ Docker テスト ターゲットを実行します。
docker build -f docker/Dockerfile --targetunit-test -t bubblehub:unit 。
docker run --rm --privileged --security-opt seccomp=unconfined bubblehub:unit
CI は、ユニットおよび統合カバレッジ レポートを使用して、Codecov を通じてライン カバレッジ ( libbubble および bubblehub のプロジェクト目標 45%) を強制します。ローカル ユニット カバレッジ コマンドと HTML レポートの場所については、CONTRIBUTING.md を参照してください。
統合テストには、モデルと OpenClaw の依存関係用の永続的なキャッシュも必要です。リモート/NFS ワークスペースで失敗する可能性がある $PWD バインド マウントの代わりに、Docker 名前付きボリュームを使用します。
docker volume create bubblehub-cache-local
docker volume create bubblehub-openclaw-local
docker build -f docker/Dockerfile --target integration-test -t bubblehub:integration 。
docker run --rm --privileged --security-opt seccomp=unconfined \
-v バブルハブキャッシュローカル:/キャッシュ/バブルハブ \
-v bubblehub-openclaw-local:/cache/openclaw \
バブルハブ:統合
インタラクティブな Docker 開発
pytest を実行する代わりに統合イメージを対話的に探索するには:
docker run -it --rm \
--特権 \
--security-opt seccomp=制限なし \
-e BUBBLEHUB_CACHE=/キャッシュ/バブルハブ \
-e BUBBLEHUB_MODELS_CONFIG=/cache/bubblehub/ci-models.yaml \
-e BUBBLEHUB_INTEGRATION_WORKSPACE_DIR=/cache/bubblehub/integration-workspaces \
-e OPENCLAW_CACHE_DIR=/キャッシュ/オープンクロー \
-v バブルハブキャッシュローカル:/キャッシュ/バブルハブ \
-v bubblehub-openclaw-local:/cache/openclaw \
バブルハブ:統合 \
バッシュ
コンテナ内:
scripts/ci/write-ci-model-config.sh
scripts/ci/prepare-openclaw.sh
mkdir -p /cache/bubblehub/integration-workspaces/dev-playground
バブルシェル --allow-ne

twork --root-dir /cache/bubblehub/integration-workspaces/dev-playground
ライセンス
Apache ライセンス 2.0。詳細については、「ライセンス」を参照してください。
コントリビューションのガイドライン、テスト要件、プル リクエストの期待値については、CONTRIBUTING.md を参照してください。
BubbleHub は、サンドボックス化とローカル実行を備えた AI エージェントのランタイム制御レイヤーです。
Readme Apache-2.0 ライセンス
貢献活動 カスタム プロパティ スター
4 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

BubbleHub is a runtime control layer for AI agents with sandboxing and local execution. - bublhub/BubbleHub

GitHub - bublhub/BubbleHub: BubbleHub is a runtime control layer for AI agents with sandboxing and local execution. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
bublhub
/
BubbleHub
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
87 Commits 87 Commits .bubblehub-cache .bubblehub-cache .cursor/ skills .cursor/ skills .github .github app app assets assets bubblehub bubblehub docker docker docs docs examples examples libbubble libbubble scripts scripts tests tests .clang-format .clang-format .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md codecov.yml codecov.yml package.json package.json pyproject.toml pyproject.toml View all files Repository files navigation
Local LLM serving and sandboxed agents in one command.
curl -fsSL https://bubblehub.ai/install.sh | bash
Windows PowerShell installer:
irm https: // bubblehub.ai / install.ps1 | iex
Windows double-click installer:
irm https: // bubblehub.ai / download / latest / BubbleHub-0.2.0-x64.exe - OutFile BubbleHub-0.2.0-x64.exe
.\ BubbleHub-0.2.0-x64.exe
Check it:
docker pull ghcr.io/bublhub/bubblehub:latest
Use a release image as a base:
FROM ghcr.io/bublhub/bubblehub:v0.2.0
Quick Start
Ask the local model a question:
bubble prompt --text " Say hello from BubbleHub "
Run an agent in the sandbox:
bubble run --root-dir ./examples/basic --binary ./examples/basic/basic_agent.py --memory 16G
Name an agent for bubble ps , the shell prompt, and the BubbleHub desktop app:
bubble shell --name reviewer --root-dir ./workspace
bubble ps --kill agt-...
Start the OpenAI-compatible local endpoint (optional):
bubble serve
Pick or inspect models:
bubble models
bubble models list
bubble models stop
Open app:
Exposes an OpenAI-compatible endpoint at http://127.0.0.1:8000/v1 .
Keeps warm model backends shared across agents.
Runs agents in a Linux sandbox with restricted filesystem and network access.
Injects local inference into agents as OPENAI_BASE_URL and OPENAI_API_KEY .
Provides the BubbleHub desktop app for graphical monitoring, manifest review, and base-model selection.
bubble run starts the shared inference endpoint before launching an agent and injects:
OPENAI_BASE_URL=http://127.0.0.1:8000/v1
OPENAI_API_KEY=bubblehub-local
BUBBLEHUB_API_BASE_URL=http://127.0.0.1:8000
BUBBLEHUB_SANDBOX_INFERENCE_HOST=127.0.0.1
BUBBLEHUB_SANDBOX_INFERENCE_PORT=8000
Sandboxed agents only get access to the local inference endpoint by default. HTTP clients see a loopback policy proxy at http://127.0.0.1:18080 , and BubbleHub injects HTTP_PROXY , HTTPS_PROXY , http_proxy , and https_proxy for isolated sandboxes. The proxy is owned by libbubble : it checks the persistent sandbox access manifest under ~/.local/state/bubblehub/sandboxes/<agent-id>/access-manifest.json (or BUBBLEHUB_STATE_DIR ) and either forwards allowed requests or returns a logged 403 . Unknown requests fail closed; when no host prompt is available, BubbleHub records them as pending in the manifest for later dashboard review. Use --allow-network with bubble run or bubble shell when an agent setup step needs general outbound network access.
When bubble run or bubble shell is connected to a real terminal, first access to a new host pauses the agent and prompts on the host for always , never , or ask every time (approve now) . Non-interactive runs fail closed and print a reminder to run bubble dashboard ; the dashboard resolves pending sandbox access requests before the live resource view opens. Use bubble manifest --root-dir <dir> or bubble manifest --agent-id <agent> to inspect and edit persisted policies. Manifest policies are exactly always , never , or ask ; HTTP policies match domains plus the visible HTTP method/path, while HTTPS CONNECT can only be matched at the host/port level.
When --root-dir is provided, non-system binaries must live inside that root and BubbleHub mounts the root as the sandbox workspace. System binaries from /usr , /bin , /sbin , or /opt/bubblehub can still be used with a root directory, which lets bubble shell --root-dir <dir> open a shell inside the workspace sandbox. When --root-dir is omitted, non-system binaries are copied into a temporary workspace before the sandbox starts. Inside the sandbox, BubbleHub Python prompt/shim calls detect BUBBLEHUB_SANDBOX=1 and use the forwarded inference endpoint instead of loading the native shared library.
Installed sandboxes run over an Ubuntu 26.04 root filesystem using a per-agent overlay. The Ubuntu lower filesystem stays unchanged; writes outside the workspace copy up into .bubblehub/agents/<agent-id>/overlay/upper and persist with that agent. Use --force-new-sandbox or --overwrite-sandbox to discard the persistent home and private overlay for the current workspace.
For implementation details, security assumptions, and known gaps, see docs/sandbox.md .
For step-by-step guides, see docs/getting_started.md and docs/running_an_agent.md .
OpenClaw can be installed entirely from inside the sandbox. The persistent agent home keeps nvm , npm global packages, and OpenClaw config across runs.
bubble shell --allow-network --root-dir openclaw
Inside the sandbox shell:
curl -fsSL https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.5/install.sh | bash
export NVM_DIR= " $HOME /.nvm "
. " $NVM_DIR /nvm.sh "
nvm install 22.19.0
npm install -g openclaw@latest
openclaw onboard --install-daemon
Releases
BubbleHub ships source install assets from GitHub Releases and runtime images from GHCR.
Push a v* tag. The release workflow runs unit tests, runs local-inference integration tests, then publishes:
CI also builds synthetic release assets for pull requests and validates installation in parallel before the branch is considered release-ready:
Linux curl | bash install from a clean Ubuntu container
Linux apt install ./BubbleHub-<version>-x64.deb from a separate clean Ubuntu container
Windows PowerShell install on a self-hosted Windows desktop runner
Windows .exe bootstrapper install on a separate self-hosted Windows desktop runner
Each smoke test installs a previous synthetic version, verifies bubble --version , shortcut creation, basic CLI behavior, and the Control Center /health response, installs the current synthetic version over it, and verifies the runtime reports the updated version. The Windows PowerShell and .exe paths both install the released .deb into WSL and create Windows launchers. The Windows release-install jobs require [self-hosted, Windows, X64] runners with an interactive desktop, WSL2/WSLg, Python, a clean Ubuntu base distro, and permission to import disposable WSL distros.
For Cursor-written release notes, ask Cursor to use the BubbleHub release-notes skill before tagging.
It writes .github/releases/<tag>.md from commits since the previous release, and the release workflow uses that file when present.
curl -fsSL https://bubblehub.ai/download/linux/v0.2.0/install.sh | BUBBLEHUB_VERSION=v0.2.0 bash
Download a specific Debian package:
curl -LO https://bubblehub.ai/download/linux/v0.2.0/BubbleHub-0.2.0-x64.deb
sudo apt install ./BubbleHub-0.2.0-x64.deb
Use the matching runtime image:
docker pull ghcr.io/bublhub/bubblehub:v0.2.0
Build from source
Create and activate a virtual environment
python3 -m venv .venv
source .venv/bin/activate
Install dependencies and build BubbleHub. The build installs the native bubblehub-sandbox helper under /usr/local/bin with the permissions required for sandbox setup, creates the Ubuntu 26.04 rootfs on first install, and preserves it on later local rebuilds for a faster development loop.
./scripts/install-deps.sh
./scripts/build.sh
Test Before Push
Run the same Docker test targets that CI uses before pushing:
docker build -f docker/Dockerfile --target unit-test -t bubblehub:unit .
docker run --rm --privileged --security-opt seccomp=unconfined bubblehub:unit
CI enforces line coverage through Codecov (45% project target for libbubble and bubblehub ) using unit and integration coverage reports. See CONTRIBUTING.md for the local unit coverage command and HTML report locations.
Integration tests also need persistent caches for the model and OpenClaw dependencies. Use Docker named volumes instead of $PWD bind mounts, which can fail on remote/NFS workspaces:
docker volume create bubblehub-cache-local
docker volume create bubblehub-openclaw-local
docker build -f docker/Dockerfile --target integration-test -t bubblehub:integration .
docker run --rm --privileged --security-opt seccomp=unconfined \
-v bubblehub-cache-local:/cache/bubblehub \
-v bubblehub-openclaw-local:/cache/openclaw \
bubblehub:integration
Interactive Docker Development
To explore the integration image interactively instead of running pytest:
docker run -it --rm \
--privileged \
--security-opt seccomp=unconfined \
-e BUBBLEHUB_CACHE=/cache/bubblehub \
-e BUBBLEHUB_MODELS_CONFIG=/cache/bubblehub/ci-models.yaml \
-e BUBBLEHUB_INTEGRATION_WORKSPACE_DIR=/cache/bubblehub/integration-workspaces \
-e OPENCLAW_CACHE_DIR=/cache/openclaw \
-v bubblehub-cache-local:/cache/bubblehub \
-v bubblehub-openclaw-local:/cache/openclaw \
bubblehub:integration \
bash
Inside the container:
scripts/ci/write-ci-model-config.sh
scripts/ci/prepare-openclaw.sh
mkdir -p /cache/bubblehub/integration-workspaces/dev-playground
bubble shell --allow-network --root-dir /cache/bubblehub/integration-workspaces/dev-playground
License
Apache License 2.0. See LICENSE for details.
See CONTRIBUTING.md for contribution guidelines, testing requirements, and pull request expectations.
BubbleHub is a runtime control layer for AI agents with sandboxing and local execution.
Readme Apache-2.0 license Contributing
Contributing Activity Custom properties Stars
4 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
