---
source: "https://github.com/Cznorth/winkterm"
hn_url: "https://news.ycombinator.com/item?id=48432710"
title: "Show HN: WinkTerm – AI that shares the same PTY as you in the terminal"
article_title: "GitHub - Cznorth/winkterm: AI that shares your terminal session. Open-source AI terminal where the AI types commands directly into your shell, not just suggests them. SSH, Docker, BYO LLM. · GitHub"
author: "Cznorth"
captured_at: "2026-06-07T10:03:48Z"
capture_tool: "hn-digest"
hn_id: 48432710
score: 1
comments: 0
posted_at: "2026-06-07T07:32:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: WinkTerm – AI that shares the same PTY as you in the terminal

- HN: [48432710](https://news.ycombinator.com/item?id=48432710)
- Source: [github.com](https://github.com/Cznorth/winkterm)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T07:32:02Z

## Translation

タイトル: HN を表示: WinkTerm – 端末内であなたと同じ PTY を共有する AI
記事のタイトル: GitHub - Cznorth/winkterm: ターミナル セッションを共有する AI。 AI がコマンドを提案するだけでなくシェルに直接入力するオープンソース AI ターミナル。 SSH、Docker、BYO LLM。 · GitHub
説明: 端末セッションを共有する AI。 AI がコマンドを提案するだけでなくシェルに直接入力するオープンソース AI ターミナル。 SSH、Docker、BYO LLM。 - Cznorth/ウィンクターム

記事本文:
GitHub - Cznorth/winkterm: ターミナル セッションを共有する AI。 AI がコマンドを提案するだけでなくシェルに直接入力するオープンソース AI ターミナル。 SSH、Docker、BYO LLM。 · GitHub
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
ズノース
/
ウィンクターム
公共
通知
よ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
191 コミット 191 コミット .claude-plugin .claude-plugin .devcontainer .devcontainer .github .github エージェント スキル エージェント スキル アセット アセット バックエンド バックエンド ビルド ビルド デスクトップ デスクトップ ドキュメント ドキュメント フロントエンド フロントエンド プラグイン/ winkterm-remote プラグイン/ winkterm-remote スクリプト スクリプト テスト テスト Web サイト Web サイト .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md SHOWHN.md SHOWHN.md TODO.md TODO.md docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
端末セッションを共有する AI。
コマンドを提案するチャットボットではありません。同じ PTY であなたと一緒に入力するコラボレーター。
デモ •
特徴 •
エージェントAPI •
クイックスタート •
なぜウィンクタームなのか? •
アーキテクチャ •
構成 •
開発 •
ロードマップ
GIF — 実際の SSH セッション: 間違ったコマンド ( ipconfig )、その後 # 何が間違っているのか ; AI は同じ PTY で応答し、修正を事前に入力できます。
Promo — 複数の SSH タブを備えた単一列ターミナル。 Craft はホスト全体 ( list_ssh_connections 、terminal_exec 、 ssh_run ) のチェックを調整します。
$ ipconfig
コマンド「ipconfig」が見つかりません。つまり: ...
$ # どうしたの
[WinkTerm] `ipconfig` は Windows コマンドです。Linux では `ip addr` (または `ifconfig`) を使用します。
$ ip addr█ ← AIが書きました。 Enter を押して実行します。編集するにはバックスペースキーを押します。キャンセルするには Ctrl+C を押します。
これは、端末に貼り付けられた ChatGPT ラッパーではありません。
AI は端末の入力ラインに直接書き込みます。制御はあなたが行います — Enter を押して実行したり、自由に編集したり、

キャンセル。これは、画面全体にアクセスして入力できる知識豊富なパートナーとサーバーに SSH 接続するようなものです。
共有 PTY セッション — AI とユーザーは同じ端末プロセスで動作します。コピー＆ペーストしたり、コンテキストなしで「このコマンドを実行」したりすることはできません。
ターミナル内チャット — シェル プロンプトがある場所に # と入力し、続いて質問を入力します。 Alt-Tab を押す必要はありません。
サイドバー AI パネル — マルチ会話タブ、AI 生成のタイトル、チャット/クラフト モード切り替えを備えた完全な会話型インターフェイス。
永続的なチャット履歴 — 会話は ~/.winkterm/chat_history.json に保存され、ページの読み込み時に復元されます。 WebSocket が再接続され、バックエンドが再起動されます。
ストリーミングの再開 — 送信中のトークンを失うことなく、応答中に更新または再接続します。アクティブなストリームはサーバー側で追跡され、新しい WebSocket クライアントに再生されます。
ストリーミング キューと提案 — AI が応答している間、フォローアップ メッセージをキューに入れます (キュー内のアイテムはいつでも中断または削除できます)。各回答後にワンクリックでフォローアップ提案チップを取得します。
外部エージェント API — 認証された HTTP インターフェイスにより、外部エージェントはインストール可能なスキルを介して端末、SSH、およびファイル転送を実行できます (以下のエージェント API のハイライトを参照)。
タブ内のエージェント端末 — エージェントが作成したセッションは、通常の端末タブとして表示されます (個別のモニター パネルはありません)。 GET /api/sessions/stream は UI の同期を保ちます。 WebSocket の切断によって PTY セッションが強制終了されなくなりました。リフレッシュにより、バッファされた出力が再生されます。
リモート アクセス認証 — Web アクセスはアクセス キーによって保護されます。ローカル デスクトップ クライアントには認証は必要ありません。
SSH リモート接続 — 組み込みのファイル転送を使用してリモート サーバーに接続します。パスワードを再入力せずに接続を編集すると、保存された資格情報が保持されます。
設定のエクスポートと秘密に安全な保存 — 設定から config.json をエクスポート ( GET /api/setti

ngs/エクスポート)。保存時にパスワードまたは API キー フィールドを空白にしても、保存されているシークレットは消去されません。
信頼性の高いターミナル - デバウンスされた PTY サイジングにより、切り詰められた PowerShell プロンプトが修正されました。エージェントのマルチタブ作成で空のペインが残らなくなりました。 WebSocket の再接続はサイレントです (PSReadLine を中断する切断バナーはありません)。
国際化 — 英語/中​​国語の UI が組み込まれており、初回起動時に言語を選択できます。
マルチモデルのサポート — 独自の LLM を導入します。 OpenAI、Anthropic、Ollama、または任意の OpenAI 互換エンドポイント。
Docker とデスクトップ — docker compose up を使用して即座にデプロイするか、スタンドアロンのデスクトップ アプリとしてパッケージ化します (Windows/macOS)。
WinkTerm の HTTP エージェント API は、単なる後付けではなく、AI エージェント (クロード コード、カーソルなど) がターミナルをリモートで操作できるように設計されています。
エンドポイント
目的
POST /api/agent/terminals/{id}/exec
アトミック実行: stdout + 実際の exit_code + 現在の cwd を返します。 Sentinel マーカーはコマンド エコーとプロンプトを自動削除します。 cwd / env サブシェル インジェクションをサポートします (永続的な端末状態を汚染しません)。
POST /api/agent/ssh/{conn_id}/run
ワンショット SSH 実行: バンドルの作成 → 実行 → を 1 回の呼び出しにまとめ、3 往復を節約します。
POST/GET/PUT/DELETE /api/agent/ssh/connections[/{id}]
SSH 接続管理: 保存された接続プロファイルの完全な CRUD、および POST /api/agent/ssh/import/electerm 。更新すると、マスクされた/省略されたシークレットは変更されません。
POST /api/agent/terminals/{id}/input
JSON に制御文字を詰め込む代わりに、名前付き制御キー: {"keys": ["ctrl+c"]} を追加します。 data_b64 入力は、マルチレイヤー引用符エスケープ地獄をバイパスします。
GET /api/agent/terminals/{id}/snapshot?pattern=...
サーバー側 grep : 256KB ローリング バッファー内の正規表現一致。帯域幅を節約します。
GET /api/agent/terminals/{id}/stream
SSE ライブ出力: 長時間実行されるコマンドのキラー機能 / tail -f 。以降で再開

r 切断します。
GET /api/agent/events/stream
オペレーション イベント フィード: すべてのエージェント アクションはリング バッファ (永続性なし) にプッシュされ、SSE 経由でブロードキャストされます。
GET /api/sessions / GET /api/sessions/stream
セッションのライフサイクル: ユーザーに表示される端末をリストします。 SSE は session_created / session_closed をプッシュするため、Web UI タブ バーはエージェントのアクティビティと同期したままになります。
/api/chat/conversations を取得する
チャットの永続性: 保存されたサイドバーの会話をリストします ( ~/.winkterm/chat_history.json にも書き込まれます)。
GET /api/settings/export
構成バックアップ: 完全な config.json (localhost または有効な X-Access-Key ) をダウンロードします。
GET /api/agent/handshake
ゼロ構成オンボーディング: localhost または Web 認証されたクライアントはトークンを自動的に取得します。エージェントはセッションごとにユーザーに質問する必要はありません。
主な設計上の選択
終了コードはファーストクラスです。失敗を検出するために出力を grep する必要はありません。exit_code は応答内にあります。
30 以上の名前付きキー: ctrl+c / up / tab / esc / f1 — JSON には生の制御文字はありません。
Base64 入力: 複雑な awk / jq / heredoc コマンドは、トリプルエスケープを回避して command_b64 を通過します。
永続的な cwd 追跡: 実行監視は実行のたびに $PWD を報告します。監視パネルには端末の現在のディレクトリが表示されます。
TTL 自動クリーンアップ: 端末はデフォルトで 30 分間のアイドル TTL に設定されているため、忘れられた端末がリークすることはありません。
待機理由フィールド : アイドル / タイムアウト / no_output を区別するため、呼び出し側は何が起こったのかを知ることができます。
Claude Code プラグイン (ワンライナー):
/プラグイン マーケットプレイス Cznorth/winkterm を追加
/plugin install winkterm-remote@winkterm
任意のエージェント (実行中のバックエンドからの生のスキル):
curl -s http:// < your-winkterm-host > :8000/api/agent/skill.md > SKILL.md
SKILL.md を Claude Code / Cursor / エージェント ツールのスキル ディレクトリにドロップすると、AI は API の操作方法をすぐに認識します。スキルはバージョン管理されています - エージェントはセッションごとに更新をチェックします

の上。
内部クラフト エージェントと外部 HTTP API は、同じターミナル セッション プールとツール サーフェス ( list / create / close / snapshot / input / exec / ssh_run ) を共有します。エージェントが作成したターミナルはユーザーに表示され、メイン UI で通常のタブとして開きます。ライブタブ同期の場合は /api/sessions/stream をサブスクライブし、色分けされた操作監査ログの場合は /api/agent/events/stream をサブスクライブします。
📖 事例: AI エージェントが WinkTerm 経由で 30 分以内に XMR クリプトジャッカーを見つけて削除
実際のインシデントの記録: ユーザーは「107.173 サーバーの負荷が高い」とだけ言っており、AI エージェントはエージェント API を介して発見→調査→キル チェーンの再構築→強化→不正行為の報告をエンドツーエンドで完了しました。このリリースの 9 つの新機能は、まさにこのケースで発生した問題点からリバース エンジニアリングされたものです。
特徴
ウィンクターム
ワープ
タビー
クロード・コード
共有 PTY (端末内の AI タイプ)
✅
❌
❌
❌
オープンソース
✅
✅
✅
❌
自己ホスト型 / BYO LLM
✅
❌
❌
✅
ウェブUI
✅
✅
✅
❌ (CLI のみ)
SSH+ファイル転送
✅
❌
✅
❌
デスクトップアプリ
✅
✅
✅
❌
WinkTerm の中核となる哲学: 端末は操作が行われる場所です。 AIはその傍らではなく、その中に生きるべきです。 AI が入力行にコマンドを書き込み、Enter キーを押すとき、あなたは盲目的に信頼しているわけではなく、検討し、理解し、選択していることになります。それが協力作戦です。
docker run -p 3000:3000 -p 8000:8000 \
-e ANTHROPIC_API_KEY= *** \
ghcr.io/cznorth/winkterm:最新
または docker-compose を使用します。
git クローン https://github.com/Cznorth/winkterm.git
cd ウィンクターム
cp .env.example .env
# API キーを使用して .env を編集します
ドッカー構成 -d
構成ファイルは winkterm-data ボリュームを /root/.winkterm にマウントするため、構成、チャット履歴、SSH 資格情報はコンテナーの再構築後も残ります。このイメージには、インストール可能なエージェント スキルもバンドルされています (404 はありません)。

skill.md フェッチ)。
次に、http://localhost:3000 を開きます。
「リリース」ページから、使用しているプラットフォーム用の最新リリースをダウンロードします。
macOS : .app バンドル (インテルおよび Apple シリコン)。デスクトップ ビルドは、WebView を開く前に埋め込みバックエンドを開始し、開発専用の localhost:8000 を静的アセットにベイクすることを回避します。
変数
説明
デフォルト
ANTHROPIC_API_KEY
Anthropic API キー (必須)
—
OPENAI_API_KEY
OpenAI API キー (代替)
—
MODEL_NAME
使用するモデル
クロード・ソネット-4-20250514
OPENAI_BASE_URL
カスタム API エンドポイント
—
AGENT_RECURSION_LIMIT
エージェントの再帰制限
100
プロメテウス_URL
プロメテウスエンドポイント
http://ローカルホスト:9090
LOKI_URL
Loki エンドポイント
http://ローカルホスト:3100
デバッグ
デバッグモードを有効にする
偽
独自の LLM を使用する : WinkTerm は OpenAI 互換プロトコルを使用します。 OPENAI_BASE_URL を任意のプロバイダー (Ollama、vLLM、Groq、OpenRouter など) に設定すると、WinkTerm がそれを使用します。
ユーザーキーボード入力
│
▼
フロントエンド ターミナル (xterm.js)
│ ウェブソケット
▼
ws_handler.py
│
§── 通常入力 ──► pty_manager.write() ──► シェルプロセス
│
└── # で始まる行 ──► インターセプト ──► Agent (LangGraph)
│
§── get_terminal_context()
§── ターミナル入力()
└── write_command() ──► pty ──► 端末入力行
重要な洞察 : AI メッセージは PTY 出力ストリームに直接書き込まれるため、端末にシームレスに表示されます。個別の UI クロムやコンテキストの切り替えはありません。
レイヤー
テクノロジー

[切り捨てられた]

## Original Extract

AI that shares your terminal session. Open-source AI terminal where the AI types commands directly into your shell, not just suggests them. SSH, Docker, BYO LLM. - Cznorth/winkterm

GitHub - Cznorth/winkterm: AI that shares your terminal session. Open-source AI terminal where the AI types commands directly into your shell, not just suggests them. SSH, Docker, BYO LLM. · GitHub
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
Cznorth
/
winkterm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
191 Commits 191 Commits .claude-plugin .claude-plugin .devcontainer .devcontainer .github .github agent-skill agent-skill assets assets backend backend build build desktop desktop docs docs frontend frontend plugins/ winkterm-remote plugins/ winkterm-remote scripts scripts test test website website .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md SHOWHN.md SHOWHN.md TODO.md TODO.md docker-compose.yml docker-compose.yml View all files Repository files navigation
AI that shares your terminal session.
Not a chatbot that suggests commands. A collaborator that types alongside you in the same PTY.
Demo •
Features •
Agent API •
Quick Start •
Why WinkTerm? •
Architecture •
Configuration •
Development •
Roadmap
GIF — real SSH session: a mistaken command ( ipconfig ), then # what's wrong ; the AI answers in the same PTY and can pre-fill the fix.
Promo — single-column terminal with multiple SSH tabs; Craft orchestrates checks across hosts ( list_ssh_connections , terminal_exec , ssh_run ).
$ ipconfig
Command 'ipconfig' not found, did you mean: ...
$ # what's wrong
[WinkTerm] `ipconfig` is a Windows command — on Linux use `ip addr` (or `ifconfig`).
$ ip addr█ ← AI wrote this. Press Enter to run. Backspace to edit. Ctrl+C to cancel.
This is not a ChatGPT wrapper pasted into a terminal.
The AI writes directly into your terminal's input line. You stay in control — press Enter to execute, edit freely, or cancel. It's like SSH-ing into a server with a knowledgeable partner who can reach across the screen and type.
Shared PTY Session — AI and user operate in the same terminal process. No copy-paste, no "run this command" without context.
In-Terminal Chat — Type # followed by your question, right where your shell prompt is. No need to alt-tab.
Sidebar AI Panel — Full conversational interface with multi-conversation tabs, AI-generated titles, and chat/craft mode switching.
Persistent Chat History — Conversations are saved to ~/.winkterm/chat_history.json and restored on page load; survives WebSocket reconnects and backend restarts.
Streaming Resume — Refresh or reconnect mid-response without losing in-flight tokens; active streams are tracked server-side and replayed to new WebSocket clients.
Streaming Queue & Suggestions — Queue follow-up messages while the AI is responding (interrupt or drop queued items anytime), and get one-click follow-up suggestion chips after each answer.
External Agent API — An authenticated HTTP interface lets external agents drive your terminal, SSH, and file transfers via an installable skill (see Agent API highlights below).
Agent Terminals in Your Tabs — Agent-created sessions appear as normal terminal tabs (no separate monitor panel). GET /api/sessions/stream keeps the UI in sync; WebSocket disconnect no longer kills PTY sessions — refresh replays buffered output.
Remote Access Auth — Web access is protected by an access key; the local desktop client needs no authentication.
SSH Remote Connections — Connect to remote servers with built-in file transfer. Editing a connection without re-entering the password keeps the saved credential.
Settings Export & Secret-safe Saves — Export config.json from Settings ( GET /api/settings/export ). Blank password or API key fields on save do not wipe stored secrets.
Reliable Terminals — Debounced PTY sizing fixes truncated PowerShell prompts; agent multi-tab creation no longer leaves empty panes; WebSocket reconnect is silent (no disconnect banner that breaks PSReadLine).
Internationalization — Built-in English / Chinese UI, with language selection on first launch.
Multi-Model Support — Bring your own LLM. OpenAI, Anthropic, Ollama, or any OpenAI-compatible endpoint.
Docker & Desktop — Deploy instantly with docker compose up or package as a standalone desktop app (Windows/macOS).
WinkTerm's HTTP Agent API is designed for AI agents (Claude Code, Cursor, etc.) to drive the terminal remotely — not just an afterthought.
Endpoint
Purpose
POST /api/agent/terminals/{id}/exec
Atomic execution : returns stdout + real exit_code + current cwd . Sentinel marker auto-strips command echo and prompt. Supports cwd / env subshell injection (doesn't pollute persistent terminal state).
POST /api/agent/ssh/{conn_id}/run
One-shot SSH execution : bundles create → exec → close into one call, saving 3 round-trips.
POST/GET/PUT/DELETE /api/agent/ssh/connections[/{id}]
SSH connection management : full CRUD on the stored connection profiles, plus POST /api/agent/ssh/import/electerm . Update leaves masked/omitted secrets unchanged.
POST /api/agent/terminals/{id}/input
Named control keys : {"keys": ["ctrl+c"]} instead of stuffing control chars into JSON. data_b64 input bypasses multi-layer quote escape hell.
GET /api/agent/terminals/{id}/snapshot?pattern=...
Server-side grep : regex-match within the 256KB rolling buffer. Save bandwidth.
GET /api/agent/terminals/{id}/stream
SSE live output : killer feature for long-running commands / tail -f . Resume with since after disconnect.
GET /api/agent/events/stream
Operation event feed : every agent action is pushed to a ring buffer (no persistence), broadcast via SSE.
GET /api/sessions / GET /api/sessions/stream
Session lifecycle : list user-visible terminals; SSE pushes session_created / session_closed so the web UI tab bar stays in sync with agent activity.
GET /api/chat/conversations
Chat persistence : list saved sidebar conversations (also written to ~/.winkterm/chat_history.json ).
GET /api/settings/export
Config backup : download full config.json (localhost or valid X-Access-Key ).
GET /api/agent/handshake
Zero-config onboarding : localhost or web-auth'd clients get the token automatically. The agent doesn't need to ask the user every session.
Key Design Choices
Exit codes are first-class : no need to grep output to detect failure — exit_code is in the response.
30+ named keys : ctrl+c / up / tab / esc / f1 — no raw control chars in JSON.
base64 input : complex awk / jq / heredoc commands go through command_b64 , sidestepping triple-escaping.
Persistent cwd tracking : the exec sentinel reports $PWD after every run; the monitoring panel displays the terminal's current directory.
TTL auto-cleanup : terminals default to 30-minute idle TTL, so forgotten terminals don't leak.
wait reason field : distinguishes idle / timeout / no_output so callers know what happened.
Claude Code plugin (one-liner):
/plugin marketplace add Cznorth/winkterm
/plugin install winkterm-remote@winkterm
Any agent (raw skill from a running backend):
curl -s http:// < your-winkterm-host > :8000/api/agent/skill.md > SKILL.md
Drop SKILL.md into Claude Code / Cursor / any agent tool's skills directory and the AI immediately knows how to drive the API. The skill is versioned — agents check for updates each session.
Internal craft agents and the external HTTP API share the same terminal session pool and tool surface ( list / create / close / snapshot / input / exec / ssh_run ). Agent-created terminals are user-visible and open as regular tabs in the main UI. Subscribe to /api/sessions/stream for live tab sync, or /api/agent/events/stream for a color-coded operation audit log.
📖 Case: AI agent locates and removes an XMR cryptojacker in 30 minutes via WinkTerm
A real incident write-up: user said only "the 107.173 server's load is high," and the AI agent completed discovery → investigation → kill chain reconstruction → hardening → abuse reporting end-to-end via the Agent API. The 9 new features in this release were reverse-engineered from the pain points hit during this exact case.
Feature
WinkTerm
Warp
Tabby
Claude Code
Shared PTY (AI types in your terminal)
✅
❌
❌
❌
Open source
✅
✅
✅
❌
Self-hosted / BYO LLM
✅
❌
❌
✅
Web UI
✅
✅
✅
❌ (CLI only)
SSH + file transfer
✅
❌
✅
❌
Desktop app
✅
✅
✅
❌
WinkTerm's core philosophy : The terminal is where operations happen. AI should live inside it, not beside it. When the AI writes a command into your input line and you press Enter, you're not blindly trusting — you're reviewing, understanding, and choosing. That's collaborative ops.
docker run -p 3000:3000 -p 8000:8000 \
-e ANTHROPIC_API_KEY= *** \
ghcr.io/cznorth/winkterm:latest
Or with docker-compose:
git clone https://github.com/Cznorth/winkterm.git
cd winkterm
cp .env.example .env
# Edit .env with your API keys
docker compose up -d
The compose file mounts a winkterm-data volume at /root/.winkterm , so config, chat history, and SSH credentials survive container rebuilds. The image also bundles the installable agent skill (no 404 on skill.md fetch).
Then open http://localhost:3000
Download the latest release for your platform from the Releases page .
macOS : .app bundle (Intel & Apple Silicon). The desktop build starts the embedded backend before opening the WebView and avoids baking dev-only localhost:8000 into static assets.
Variable
Description
Default
ANTHROPIC_API_KEY
Anthropic API key (required)
—
OPENAI_API_KEY
OpenAI API key (alternative)
—
MODEL_NAME
Model to use
claude-sonnet-4-20250514
OPENAI_BASE_URL
Custom API endpoint
—
AGENT_RECURSION_LIMIT
Agent recursion limit
100
PROMETHEUS_URL
Prometheus endpoint
http://localhost:9090
LOKI_URL
Loki endpoint
http://localhost:3100
DEBUG
Enable debug mode
false
Bring your own LLM : WinkTerm uses the OpenAI-compatible protocol. Set OPENAI_BASE_URL to any provider (Ollama, vLLM, Groq, OpenRouter, etc.) and WinkTerm will use it.
User Keyboard Input
│
▼
Frontend Terminal (xterm.js)
│ WebSocket
▼
ws_handler.py
│
├── Normal input ──► pty_manager.write() ──► shell process
│
└── Lines starting with # ──► intercept ──► Agent (LangGraph)
│
├── get_terminal_context()
├── terminal_input()
└── write_command() ──► pty ──► terminal input line
Key insight : AI messages are written directly into the PTY output stream, so they appear seamlessly in your terminal. No separate UI chrome, no context switching.
Layer
Technology

[truncated]
