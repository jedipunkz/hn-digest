---
source: "https://github.com/dgr8akki/engram"
hn_url: "https://news.ycombinator.com/item?id=48521226"
title: "Engram: Offline MCP memory server for AI coding tools to share memory"
article_title: "GitHub - dgr8akki/engram: Local personal knowledge system with semantic search — works with Claude Code, Cursor, and Antigravity via MCP · GitHub"
author: "dgr8akki"
captured_at: "2026-06-13T21:32:15Z"
capture_tool: "hn-digest"
hn_id: 48521226
score: 2
comments: 0
posted_at: "2026-06-13T20:40:47Z"
tags:
  - hacker-news
  - translated
---

# Engram: Offline MCP memory server for AI coding tools to share memory

- HN: [48521226](https://news.ycombinator.com/item?id=48521226)
- Source: [github.com](https://github.com/dgr8akki/engram)
- Score: 2
- Comments: 0
- Posted: 2026-06-13T20:40:47Z

## Translation

タイトル: Engram: AI コーディング ツールがメモリを共有するためのオフライン MCP メモリ サーバー
記事のタイトル: GitHub - dgr8akki/engram: セマンティック検索を備えたローカル個人知識システム — MCP を介してクロード コード、カーソル、および反重力で動作します · GitHub
説明: セマンティック検索を備えたローカル個人知識システム — MCP を介してクロード コード、カーソル、および反重力で動作します - dgr8akki/engram

記事本文:
GitHub - dgr8akki/engram: セマンティック検索を備えたローカル個人知識システム — MCP を介してクロード コード、カーソル、および反重力で動作します · GitHub
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
dgr8akki
/
エングラム
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
13 コミット 13 コミット スクリプト スクリプト スキル スキル .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md config.yaml config.yaml engram engram engram_cli.py engram_cli.py engram_db.py engram_db.py engram_embeddings.py engram_embeddings.py engram_http_server.py engram_http_server.py engram_install.py engram_install.py engram_mcp_server.py engram_mcp_server.py要件.txt要件.txt setup.sh setup.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
セマンティック検索を備えたローカルファーストの個人知識ベース。MCP を介してクロード コード、カーソル、反重力、およびウィンドサーフィンとネイティブに動作するように構築されています。
すべてがあなたのマシン上で実行されます。クラウドも API コストも、コンピュータからデータが流出することもありません。
セマンティック検索 — キーワードではなく意味によって考えを検索します
MCP の統合 — クロード コード、カーソル、および反重力は、会話中にナレッジ ベースを読み書きできます
自動メモリフック — セッションの開始時に最近の思い出が自動的に挿入され、終了時にメモが保存されます。
自動保存トリガー - 「覚えています: ...」と言うとメモが即座に保存されます
HTTP REST サーバー — MCP をサポートしないツールのフォールバック
デュアル埋め込みバックエンド — 文変換 (オフライン) または Ollama
git clone https://github.com/dgr8akki/engram
CDエングラム
python3 -m venv venv && ソース venv/bin/activate
pip install -r 要件.txt
./エングラム初期化
./エングラムインストール
engram install は 3 つのことを一度に実行します。
検出されたツールごとに MCP サーバーを登録します
エージェント スキルをインストールします (LLM に Engram の使用方法を指示します)。
セッションライフサイクルフック（自動メモリ読み取り/書き込み）をインストールします。
次に IDE を再起動すると、Engram がアクティブになります。
./engram add " サブスクを解除するのを忘れると、GraphQL サブスクリプションでメモリ リークが発生する

書記「
./engram add " 複雑なフォーム状態では useState より useReducer を優先" --tags 反応、パターン
./engram 検索「メモリ リーク JavaScript」
./engram リスト --limit 20
./エングラム削除 42
./エングラム統計
./engram -h # ヘルプ
./engram -V # バージョン
AIツールの使い方
AI コーディング ツールをインストールすると、自然な会話を通じて Engram を使用できるようになります。
フックはセッション境界で自動的に実行されます。 engram install 後に設定は必要ありません。
engram install は以下を自動検出して構成します。
MCP サポートのないツールの場合:
./engramserve # http://127.0.0.1:7823 で開始
./engram autostart install # ログイン時に自動的に開始 (macOS LaunchAgent)
./engram autostart delete # 自動起動を無効にする
./engram autostart status # 実行中かどうかを確認する
エンドポイント
説明
GET /thought?limit=20
最近の考えを列挙する
GET /thoughts/search?q=...
セマンティック検索
投稿/感想
考えを保存する {"content": "...", "tags": "..."}
/思考/{id}を取得
IDで取得
/思考/{id}を削除
IDで削除
GET /統計
統計
構成
config.yaml を編集して動作をカスタマイズします。
埋め込み:
バックエンド: "文章変換" # または "ollam"
モデル：「文章トランスフォーマー/all-MiniLM-L6-v2」
寸法 : 384
デバイス: " cpu " # または "cuda"
# Ollama (バックエンドの場合: "ollam")
ollam_url : " http://localhost:11434 "
ollama_model : " nomic-embed-text "
オラマの寸法 : 768
検索:
デフォルト制限: 10
類似性閾値 : 0.3
http :
ホスト：「127.0.0.1」
ポート: 7823
すでに実行している場合は、Ollama に切り替えてください。モデルのダウンロードは必要なく、より高品質の埋め込みが可能です。
sqlite-vec のインストールが失敗する — pyenv の代わりにシステム Python または Homebrew Python を使用してください。
Pythonを醸造インストールする
/opt/homebrew/bin/python3 -m venv venv
MCP サーバーが表示されません — claude mcp list (Claude Code) を実行するか、ツールの MCP 構成を確認してください。パスが

絶対的な。 IDE を再起動します。
フックが起動しない — engram install -v でフックが書き込まれたことを確認します。クロード コードについては、~/.claude/settings.json でフック ブロックを確認してください。
空の検索結果 — しきい値を下げる: engram search "query" --threshold 0.1
セマンティック検索を備えたローカル個人知識システム — MCP を介してクロード コード、カーソル、反重力と連携します
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

Local personal knowledge system with semantic search — works with Claude Code, Cursor, and Antigravity via MCP - dgr8akki/engram

GitHub - dgr8akki/engram: Local personal knowledge system with semantic search — works with Claude Code, Cursor, and Antigravity via MCP · GitHub
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
dgr8akki
/
engram
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
13 Commits 13 Commits scripts scripts skill skill .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md config.yaml config.yaml engram engram engram_cli.py engram_cli.py engram_db.py engram_db.py engram_embeddings.py engram_embeddings.py engram_http_server.py engram_http_server.py engram_install.py engram_install.py engram_mcp_server.py engram_mcp_server.py requirements.txt requirements.txt setup.sh setup.sh View all files Repository files navigation
A local-first personal knowledge base with semantic search, built to work natively with Claude Code, Cursor, Antigravity, and Windsurf via MCP.
Everything runs on your machine — no cloud, no API costs, no data leaving your computer.
Semantic search — find thoughts by meaning, not keywords
MCP integration — Claude Code, Cursor, and Antigravity can read and write your knowledge base mid-conversation
Auto-memory hooks — sessions automatically inject recent memories on start and save notes on end
Auto-save triggers — say "remember: ..." and your note is saved instantly
HTTP REST server — fallback for tools that don't support MCP
Dual embedding backends — sentence-transformers (offline) or Ollama
git clone https://github.com/dgr8akki/engram
cd engram
python3 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
./engram init
./engram install
engram install does three things in one shot:
Registers the MCP server for each detected tool
Installs the agent skill (tells LLMs how to use Engram)
Installs session-lifecycle hooks (auto memory read/write)
Then restart your IDE and Engram is active.
./engram add " GraphQL subscriptions leak memory if you forget to unsubscribe "
./engram add " Prefer useReducer over useState for complex form state " --tags react,patterns
./engram search " memory leak javascript "
./engram list --limit 20
./engram delete 42
./engram stats
./engram -h # help
./engram -V # version
AI Tool Usage
Once installed, your AI coding tool can use Engram through natural conversation:
Hooks run automatically at session boundaries — no configuration needed after engram install .
engram install auto-detects and configures:
For tools without MCP support:
./engram serve # start on http://127.0.0.1:7823
./engram autostart install # start automatically at login (macOS LaunchAgent)
./engram autostart remove # disable autostart
./engram autostart status # check if running
Endpoint
Description
GET /thoughts?limit=20
List recent thoughts
GET /thoughts/search?q=...
Semantic search
POST /thoughts
Save a thought {"content": "...", "tags": "..."}
GET /thoughts/{id}
Get by ID
DELETE /thoughts/{id}
Delete by ID
GET /stats
Statistics
Configuration
Edit config.yaml to customise behaviour:
embeddings :
backend : " sentence-transformers " # or "ollama"
model : " sentence-transformers/all-MiniLM-L6-v2 "
dimensions : 384
device : " cpu " # or "cuda"
# Ollama (when backend: "ollama")
ollama_url : " http://localhost:11434 "
ollama_model : " nomic-embed-text "
ollama_dimensions : 768
search :
default_limit : 10
similarity_threshold : 0.3
http :
host : " 127.0.0.1 "
port : 7823
Switch to Ollama if you already have it running — no model download needed and better quality embeddings.
sqlite-vec install fails — use system Python or Homebrew Python instead of pyenv:
brew install python
/opt/homebrew/bin/python3 -m venv venv
MCP server not appearing — run claude mcp list (Claude Code) or check your tool's MCP config. Make sure paths are absolute. Restart the IDE.
Hooks not firing — confirm with engram install -v that hooks were written. For Claude Code, check ~/.claude/settings.json for a hooks block.
Empty search results — lower the threshold: engram search "query" --threshold 0.1
Local personal knowledge system with semantic search — works with Claude Code, Cursor, and Antigravity via MCP
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
