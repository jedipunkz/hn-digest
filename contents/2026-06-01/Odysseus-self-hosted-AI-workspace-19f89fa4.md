---
source: "https://github.com/pewdiepie-archdaemon/odysseus"
hn_url: "https://news.ycombinator.com/item?id=48346693"
title: "Odysseus – self-hosted AI workspace"
article_title: "GitHub - pewdiepie-archdaemon/odysseus: Self-hosted AI workspace. · GitHub"
author: "Dzheky"
captured_at: "2026-06-01T00:28:30Z"
capture_tool: "hn-digest"
hn_id: 48346693
score: 105
comments: 54
posted_at: "2026-05-31T15:50:34Z"
tags:
  - hacker-news
  - translated
---

# Odysseus – self-hosted AI workspace

- HN: [48346693](https://news.ycombinator.com/item?id=48346693)
- Source: [github.com](https://github.com/pewdiepie-archdaemon/odysseus)
- Score: 105
- Comments: 54
- Posted: 2026-05-31T15:50:34Z

## Translation

タイトル: Odysseus – 自己ホスト型 AI ワークスペース
記事のタイトル: GitHub - pewdiepie-archdaemon/odysseus: 自己ホスト型 AI ワークスペース。 · GitHub
説明: 自己ホスト型 AI ワークスペース。 。 GitHub でアカウントを作成して、pewdiepie-archdaemon/odysseus の開発に貢献してください。

記事本文:
GitHub - pewdiepie-archdaemon/odysseus: 自己ホスト型 AI ワークスペース。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ピュディパイ-アーチデーモン
/
オデュッセウス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション o

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット config/ searxng config/ searxng コア コア docker docker docs docs ライセンス ライセンス mcp_servers mcp_servers ルート ルート スクリプト スクリプト サービス サービス src src 静的 静的テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ACKNOWLEDGMENTS.md ACKNOWLEDGMENTS.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md app.py app.py docker-compose.yml docker-compose.yml install-service.sh install-service.sh odysseus-ui.service odysseus-ui.service package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml 要件-オプション.txt 要件-オプション.txt 要件.txt 要件.txt setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
─────────────────
⊹ ࣪ ˖ ૮( ˶ᵔ ᵕ ᵔ˶ )っ オデュッセウス ヴァージョン1.0
─────────────────
自己ホスト型 AI ワークスペース -- ChatGPT と Claude から得られる UI エクスペリエンスの自己ホスト型バージョンを意図しています。しかし、もっとジャンクで楽しいものがあります。独自のハードウェアで独自のデータを使用して実行します。ローカル第一、プライバシー第一で、トロイの木馬は使用しません。
チャット -- ローカル モデルまたは API とチャットします。それらを追加するのは非常に簡単です。
vLLM · llama.cpp · Ollama · OpenRouter · OpenAI
エージェント -- ツールを渡し、タスク自体をすべて実行させます。
オープンコード、MCP、Web、ファイル、シェル、スキル、メモリに基づいて構築
クックブック -- ハードウェアをスキャンし、モデルを推奨し、クリックしてダウンロードして提供します

e..簡単です！
llmfit に基づいて構築 · VRAM 対応 · GGUF / FP8 / AWQ · フィット スコアリング · vLLM / llama.cpp サービング
Deep Research -- ソースを収集、読み取り、合成して優れた視覚的なレポートを作成する複数ステップの実行。
同義ディープリサーチから転載
比較 -- モデルを並べて比較するための楽しいツールです。完全にブラインドで、バイアスなしでテストしてください。
マルチモデル・ブラインドテスト・合成
ドキュメント -- テキストを書くのはあなたです。AI はその逆ではなく、支援するために存在します。
マルチタブ エディター · マークダウン · HTML · CSV · 構文ハイライト · AI 編集 · 提案
記憶 / スキル -- 永続的な記憶とスキルにより、エージェントはあなたとあなたのタスクをよりよく理解するにつれて時間の経過とともに進化します。
ChromaDB · fastembed (ONNX) · ベクター + キーワード検索 · インポート/エクスポート
電子メール -- AI トリアージが組み込まれた IMAP/SMTP 受信箱: 緊急リマインダー、自動タグ、自動要約、自動返信下書き、自動スパム。
IMAP · SMTP · アカウントごとのルーティング · CalDAV 対応
メモとタスク -- リマインダー、ToDo リスト、エージェントが実行できるスケジュールされたタスクを含む簡単なメモ。
ping をメモする · チェックリスト · cron スタイルのタスク · ntfy / ブラウザ / 電子メール チャネル
カレンダー -- CalDAV を備えたローカルファーストのカレンダーで Radicale / Nextcloud / Apple / Fastmail に同期します。
CalDAV プル · .ics インポート/エクスポート · カレンダーごとのカラー · エージェント対応
モバイルでも動作 -- デスクトップだけでなく、携帯電話でも見た目も動作も優れています。
レスポンシブ · インストール可能 (PWA) · タッチ ジェスチャ
追加 -- さらに探索する必要があります。ぜひ試してみてください。
画像エディター · テーマエディター · ファイルアップロード (ビジョン + PDF) · Web 検索 · プリセット · セッション · 2FA
ホバーして再生する完全なツアーは、ランディング ページ ( docs/index.html ) にあります。いくつか見てみましょう:
デフォルトはそのままで機能し、アプリ内で複製、実行、構成を行います。
最初のログイン後に設定パネルを開き、オデュッセウスに LLM を指定します。
サーバー、検索プロバイダー、電子メール アカウント、

必要な場合にのみ .env をタッチします。
AUTH_ENABLED 、 DATABASE_URL などのデプロイメントレベルのものをオーバーライドします。
またはプレシード ODYSSEUS_ADMIN_PASSWORD (それ以外の場合、初期パスワードは
最初の起動時に生成および印刷されます)。
オプション 1: Docker (推奨)
git clone < your-odysseus-repo-url >
CD オデッセウス
cp .env.example .env # オプションですが、明示的なデフォルトを推奨します
docker 構成 -d --build
Compose は、Odysseus、ChromaDB、SearXNG、および ntfy を起動します。最初の実行では完全に実行されます
イメージビルド。コンテナーが正常になったら、http://localhost:7000 を開きます。
クックブックのリモート サーバーは、./data/ssh からの Odysseus 所有の SSH キーを使用します。
Docker内。クックブック -> 設定 -> サーバー で、
公開キーを作成し、それをリモートサーバーの ~/.ssh/authorized_keys に追加します。
キーを生成した後、次のコマンドを使用してホストからキーをインストールすることもできます。
ssh-copy-id -i data/ssh/id_ed25519.pub user@server
クックブックのローカル ダウンロードは ./data/huggingface に保存され、次のようにマウントされます。
オデュッセウスコンテナ内の ~/.cache/huggingface。
ドッカー構成PS
docker compose ログ --tail=120 オデッセウス
docker compose ログ オデッセウス | grep -E ' ChromaDB|MemoryVectorStore|DEGRADED '
docker compose exec odysseus python -c " from services.hwfit.models import get_models; print(len(get_models())) "
Docker で予想されるベクター メモリの起動行:
ChromaDB が接続されました: chromadb:8000
MemoryVectorStore が初期化されました
クックブック モデル カタログ チェックでは、ゼロ以外のカウントが出力されるはずです。印刷される場合
0 の場合は、 docker compose build --no-cache odysseus を使用して Odysseus イメージを再構築します。
オプション 2: 手動インストール — Linux / macOS
要件: Python 3.11 以降。 Linux/Termux では、Cookbook には tmux も必要です
バックグラウンドモデルのダウンロードと提供用。
最初にシステム パッケージをインストールします。
# Debian/Ubuntu
sudo apt インストール tmux
# アーチ
sudo pacman -S tmux
#フェドーラ
sudo dnf インストール tmu

×
次に、オデュッセウスをインストールします。
git clone < your-odysseus-repo-url >
CD オデッセウス
python3 -m venv venv
ソース venv/bin/activate
pip install -r 要件.txt
python setup.py # データ ディレクトリを作成し、初期管理者パスワードを出力します
uvicorn アプリ:アプリ --ホスト 0.0.0.0 --ポート 7000
オプション 3: 手動インストール — Windows (PowerShell)
git clone <あなたの - オデッセウス - リポジトリ - URL >
CD オデッセウス
Python - m venv venv
venv\Scripts\Activate.ps1
pip インストール - r 要件.txt
Pythonのセットアップ.py
uvicorn アプリ:アプリ -- ホスト 0.0 。 0.0 -- ポート 7000
http://localhost:7000 を開き、生成された管理者パスワードを使用してログインします。
そして、Settings内で他のすべてを設定します。
Odysseus は、シェル アクセス、ファイル アップロード、モデル ダウンロード、Web リサーチ、電子メール/カレンダー統合、API トークンなどの強力なローカル ツールを備えた自己ホスト型ワークスペースです。管理コンソールのように扱います。
ネットワークにアクセス可能な展開では、AUTH_ENABLED=true を維持します。
HTTPS と信頼できるリバース プロキシを使用せずに、パブリック インターネットに直接公開しないでください。
data/ 、 .env 、ログ、データベース、アップロード/生成されたメディアを Git から遠ざけてください。デフォルトでは無視されます。
最初の起動後に data/auth.json を確認します。意図的に必要な場合を除き、オープン サインアップを無効にし、自分のアカウントのみを管理者にし、デモ/テスト アカウントを非管理者のままにします。
管理者以外のユーザーは、デフォルトではシェル/Python/ファイルの読み取り/書き込みを行うことができず、MCP 管理、API トークン、Webhook、モデル/クックブックの提供、バックアップ/ボールト、アプリ設定などの管理者専用のルート/ツールは管理者によって制御されます。他の機能はユーザーごとの権限によって制御されるため、展開を公開する前に各ユーザーの権限を確認してください。
共有チャット、デモ、スクリーンショット、またはログに貼り付けられた API キーまたはトークンをローテーションします。
API トークンまたは Webhook を有効にする場合は、統合ごとに個別のトークンを作成し、未使用のトークンを削除します。
bを優先します

手動開発は 127.0.0.1 まで実行されます。意図的に LAN/リバース プロキシ アクセスが必要な場合にのみ、0.0.0.0 にバインドします。
フォークを公開する前に、 git status --short を実行し、 .env 、 data/ 、 logs/ 、アップロード、バックアップ、またはローカル データベースからのプライベート ファイルがステージングされていないことを確認します。
Odysseus は、そのポートでプレーン HTTP を提供します。ローカルホストや信頼できる LAN/VPN の使用には問題ありませんが、ブラウザは警告 (「安全でないページにパスワード フィールドが存在します」) を表示し、ログイン + API トークンはクリアテキストで送信されます。マシンの外部に到達可能なもの (他のデバイスと共有する Tailscale IP など) については、TLS 終端リバース プロキシを前面に配置します。
Caddy による最短パス (Let's Encrypt 証明書の自動更新):
オデッセウス.example.com {
リバースプロキシローカルホスト:7000
}
LAN のみの Tailscale 展開の場合は、Caddy + tailscale-cert または組み込みの MagicDNS HTTPS 機能の両方が機能します。 nginx/Traefik 構成は似ています。プロキシ localhost:7000 、プロキシで TLS を終了します。これが完了すると、ブラウザの警告が消え、ログインが暗号化されます。
お手伝いは大歓迎です。最良のエントリ ポイントは、新規インストールのテスト、プロバイダーのセットアップです。
バグ、モバイル/エディターの洗練、ドキュメント、および小規模な集中リファクタリング。参照
現在のヘルプ募集リストの ROADMAP.md。
ほとんどのセットアップは、 /setup または Settings を使用してアプリ内で行われます。 .envを使用する
最初の起動前に存在させたい展開レベルのデフォルトとシークレットの場合。
主要な設定:
Docker Compose にはデフォルトで次のものが含まれています。
ChromaDB → セマンティックメモリ用のベクトルストア。 Docker では、Odysseus は chromadb:8000 に接続します。ホストからは localhost:8100 として公開されます。
SearXNG → Web 検索のメタ検索。 Docker では、Odysseus は searxng:8080 に接続します。ホストからは 127.0.0.1:8080 でのみ公開されます。
ntfy → ローカル通知サービス、 localhost:8091 として公開されます。
Ollama → ローカル LLM サーバー

-- オラマ.ai
app.py # FastAPI エントリ ポイント
コア/認証、データベース、ミドルウェア、定数
src/llm_core、agent_loop、agent_tools、chat_processor、search/
ルート/チャット、セッション、ドキュメント、メモリ、モデル … エンドポイント
サービス/ドキュメント、メモリ、検索、hwfit (クックブック) …
static/index.html + app.js + style.css + js/ (モジュール式フロントエンド)
ドキュメント/ランディング ページ (index.html) + プレビュー クリップ
データ
すべてのユーザー データは data/ (gitignored): app.db (セッション、メッセージ、ドキュメント) に存在します。
Memory.json 、 presets.json 、uploads/ 、personal_docs/ 、chroma/ 、 settings.json 。
MIT – LICENSE and ACKNOWLEDGMENTS.md を参照してください。
|
|||
|||||
| | | |||||||
)_) )_) )_) ~|~
)___))___))___)\ |
)____)____)_____)\\|
_____|____|____|_____\\\__
\/
~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~
~^~ 全員参加です! ~^~
~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~
について
読み込み中にエラーが発生しました。このページをリロードしてください。
1.1k
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hosted AI workspace. . Contribute to pewdiepie-archdaemon/odysseus development by creating an account on GitHub.

GitHub - pewdiepie-archdaemon/odysseus: Self-hosted AI workspace. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
pewdiepie-archdaemon
/
odysseus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits config/ searxng config/ searxng core core docker docker docs docs licenses licenses mcp_servers mcp_servers routes routes scripts scripts services services src src static static tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ACKNOWLEDGMENTS.md ACKNOWLEDGMENTS.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md app.py app.py docker-compose.yml docker-compose.yml install-service.sh install-service.sh odysseus-ui.service odysseus-ui.service package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml requirements-optional.txt requirements-optional.txt requirements.txt requirements.txt setup.py setup.py View all files Repository files navigation
───────────────────────────────────────────────
⊹ ࣪ ˖ ૮( ˶ᵔ ᵕ ᵔ˶ )っ Odysseus vers. 1.0
───────────────────────────────────────────────
A self-hosted AI workspace -- meant to be the self-hosted version of the UI experience you get from ChatGPT and Claude. But with more jank and fun. Running on your own hardware, with your own data -- local-first, privacy-first, and no trojan.
Chat -- chat with any local model or API; adding them is super simple.
vLLM · llama.cpp · Ollama · OpenRouter · OpenAI
Agent -- hand it tools and let it run the whole task itself.
built on opencode · MCP · web · files · shell · skills · memory
Cookbook -- Scans your hardware, recommends models, click to download and serve.. easy!
built on llmfit · VRAM-aware · GGUF / FP8 / AWQ · fit scoring · vLLM / llama.cpp serving
Deep Research -- multi-step runs that gather, read, and synthesize sources into a nice visual report.
adapted from Tongyi DeepResearch
Compare -- a fun tool to compare models side by side. Test completely blind, no bias!
multi-model · blind test · synthesis
Documents -- YOU write the text, AI is there to assist, not the opposite.
multi-tab editor · markdown · HTML · CSV · syntax highlighting · AI edits · suggestions
Memory / Skills -- Persistent memory and skills, your agent evolves over time as it better understands you and your tasks!
ChromaDB · fastembed (ONNX) · vector + keyword retrieval · import/export
Email -- IMAP/SMTP inbox with AI triage built in: urgency reminders, auto-tag, auto-summary, auto-reply drafts, auto-spam.
IMAP · SMTP · per-account routing · CalDAV-aware
Notes & Tasks -- Quick notes with reminders, a todo list, and scheduled tasks the agent can act on.
note pings · checklist · cron-style tasks · ntfy / browser / email channels
Calendar -- Local-first calendar with CalDAV sync to Radicale / Nextcloud / Apple / Fastmail.
CalDAV pull · .ics import/export · per-calendar colors · agent-aware
Works on mobile -- looks and runs great on your phone, not just desktop.
responsive · installable (PWA) · touch gestures
Extras -- more to explore, happy if you give it a go!
image editor · theme editor · file uploads (vision + PDF) · web search · presets · sessions · 2FA
A full, hover-to-play tour lives on the landing page ( docs/index.html ). A few looks:
Defaults work out of the box — clone, run, configure inside the app.
Open the Settings panel after first login to point Odysseus at your LLM
server, search provider, email account, etc. Only touch .env if you need
to override deployment-level things like AUTH_ENABLED , DATABASE_URL ,
or pre-seed ODYSSEUS_ADMIN_PASSWORD (otherwise an initial password is
generated and printed on first boot).
Option 1: Docker (recommended)
git clone < your-odysseus-repo-url >
cd odysseus
cp .env.example .env # optional, but recommended for explicit defaults
docker compose up -d --build
Compose starts Odysseus, ChromaDB, SearXNG, and ntfy. First run does a full
image build. Open http://localhost:7000 after the containers are healthy.
Cookbook remote servers use an Odysseus-owned SSH key from ./data/ssh
inside Docker. In Cookbook -> Settings -> Servers , generate/copy the
public key and add it to the remote server's ~/.ssh/authorized_keys .
After generating the key, you can also install it from the host with:
ssh-copy-id -i data/ssh/id_ed25519.pub user@server
Cookbook local downloads are stored in ./data/huggingface , mounted as
~/.cache/huggingface inside the Odysseus container.
docker compose ps
docker compose logs --tail=120 odysseus
docker compose logs odysseus | grep -E ' ChromaDB|MemoryVectorStore|DEGRADED '
docker compose exec odysseus python -c " from services.hwfit.models import get_models; print(len(get_models())) "
Expected vector-memory startup lines in Docker:
ChromaDB connected: chromadb:8000
MemoryVectorStore initialized
The Cookbook model catalog check should print a non-zero count. If it prints
0 , rebuild the Odysseus image with docker compose build --no-cache odysseus .
Option 2: Manual install — Linux / macOS
Requirements: Python 3.11+. On Linux/Termux, Cookbook also requires tmux
for background model downloads and serves.
Install system packages first:
# Debian/Ubuntu
sudo apt install tmux
# Arch
sudo pacman -S tmux
# Fedora
sudo dnf install tmux
Then install Odysseus:
git clone < your-odysseus-repo-url >
cd odysseus
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python setup.py # creates data dirs and prints an initial admin password
uvicorn app:app --host 0.0.0.0 --port 7000
Option 3: Manual install — Windows (PowerShell)
git clone < your - odysseus - repo - url >
cd odysseus
python - m venv venv
venv\Scripts\Activate.ps1
pip install - r requirements.txt
python setup.py
uvicorn app:app -- host 0.0 . 0.0 -- port 7000
Open http://localhost:7000 , log in with the generated admin password,
and configure everything else inside Settings .
Odysseus is a self-hosted workspace with powerful local tools: shell access, file uploads, model downloads, web research, email/calendar integrations, and API tokens. Treat it like an admin console.
Keep AUTH_ENABLED=true for any network-accessible deployment.
Do not expose it directly to the public internet without HTTPS and a trusted reverse proxy.
Keep data/ , .env , logs, databases, and uploaded/generated media out of Git. They are ignored by default.
Review data/auth.json after first boot: disable open signup unless you intentionally want it, make only your own account admin, and keep demo/test accounts non-admin.
Non-admin users do not get shell/Python/file read/write by default, and admin-only routes/tools such as MCP management, API tokens, webhooks, model/cookbook serving, backup/vault, and app settings are admin-gated. Other features are controlled by per-user privileges, so review each user's privileges before exposing a deployment.
Rotate any API keys or tokens that were ever pasted into a shared chat, demo, screenshot, or log.
If you enable API tokens or webhooks, create separate tokens per integration and delete unused ones.
Prefer binding manual development runs to 127.0.0.1 ; bind to 0.0.0.0 only when you intentionally want LAN/reverse-proxy access.
Before publishing a fork, run git status --short and confirm no private files from .env , data/ , logs/ , uploads, backups, or local databases are staged.
Odysseus serves plain HTTP on its port. That's fine for localhost and trusted LAN/VPN use, but browsers will warn ("Password fields present on an insecure page") and the login + API tokens travel in cleartext. For anything reachable outside your machine — including a Tailscale IP shared with other devices — put a TLS-terminating reverse proxy in front.
Shortest path with Caddy (auto-renews Let's Encrypt certs):
odysseus.example.com {
reverse_proxy localhost:7000
}
For a LAN-only Tailscale deployment, Caddy + tailscale-cert or the built-in MagicDNS HTTPS feature both work. nginx/Traefik configs are similar — proxy localhost:7000 , terminate TLS at the proxy. Once that's in place, the browser warning goes away and your login is encrypted.
Help is welcome. The best entry points are fresh-install testing, provider setup
bugs, mobile/editor polish, docs, and small focused refactors. See
ROADMAP.md for the current help-wanted list.
Most setup is done inside the app with /setup or Settings . Use .env
for deployment-level defaults and secrets you want present before first boot.
Key settings:
Docker Compose includes these by default:
ChromaDB → vector store for semantic memory. In Docker, Odysseus connects to chromadb:8000 ; from the host it is exposed as localhost:8100 .
SearXNG → meta search for web search. In Docker, Odysseus connects to searxng:8080 ; from the host it is exposed only on 127.0.0.1:8080 .
ntfy → local notification service, exposed as localhost:8091 .
Ollama → local LLM server -- ollama.ai
app.py # FastAPI entry point
core/ auth, database, middleware, constants
src/ llm_core, agent_loop, agent_tools, chat_processor, search/
routes/ chat, session, document, memory, model … endpoints
services/ docs, memory, search, hwfit (Cookbook) …
static/ index.html + app.js + style.css + js/ (modular front-end)
docs/ landing page (index.html) + preview clips
Data
All user data lives in data/ (gitignored): app.db (sessions, messages, documents),
memory.json , presets.json , uploads/ , personal_docs/ , chroma/ , settings.json .
MIT -- see LICENSE and ACKNOWLEDGMENTS.md .
|
|||
|||||
| | | |||||||
)_) )_) )_) ~|~
)___))___))___)\ |
)____)____)_____)\\|
_____|____|____|_____\\\__
\ /
~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~
~^~ all aboard! ~^~
~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~~^~^~
About
There was an error while loading. Please reload this page .
1.1k
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
