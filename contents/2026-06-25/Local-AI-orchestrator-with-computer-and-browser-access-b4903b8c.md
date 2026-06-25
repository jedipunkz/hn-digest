---
source: "https://github.com/shreyasks094/Zeus"
hn_url: "https://news.ycombinator.com/item?id=48670377"
title: "Local AI orchestrator with computer and browser access"
article_title: "GitHub - shreyasks094/Zeus: Self-hosted AI agent orchestrator: Claude-driven agents with filesystem + browser tools, sub-agents, skills, workflows, and a built-in browser bridge — all local · GitHub"
author: "blackhawk094"
captured_at: "2026-06-25T09:02:42Z"
capture_tool: "hn-digest"
hn_id: 48670377
score: 1
comments: 1
posted_at: "2026-06-25T08:02:24Z"
tags:
  - hacker-news
  - translated
---

# Local AI orchestrator with computer and browser access

- HN: [48670377](https://news.ycombinator.com/item?id=48670377)
- Source: [github.com](https://github.com/shreyasks094/Zeus)
- Score: 1
- Comments: 1
- Posted: 2026-06-25T08:02:24Z

## Translation

タイトル: コンピューターとブラウザーにアクセスできるローカル AI オーケストレーター
記事のタイトル: GitHub - shreyasks094/Zeus: セルフホスト型 AI エージェント オーケストレーター: ファイル システム + ブラウザー ツール、サブエージェント、スキル、ワークフロー、および組み込みブラウザー ブリッジを備えたクロード駆動型エージェント — すべてローカル · GitHub
説明: セルフホスト型 AI エージェント オーケストレーター: ファイル システム + ブラウザー ツール、サブエージェント、スキル、ワークフロー、および組み込みブラウザー ブリッジを備えたクロード駆動型エージェント - すべてローカル - shreyasks094/Zeus

記事本文:
GitHub - shreyasks094/Zeus: セルフホスト型 AI エージェント オーケストレーター: ファイル システム + ブラウザー ツール、サブエージェント、スキル、ワークフロー、および組み込みブラウザー ブリッジを備えたクロード駆動型エージェント — すべてローカル · GitHub
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
シュレヤスク094
/
ゼウス
公共
お知らせ

s
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE バックエンド バックエンド cli cli 拡張機能拡張 フロントエンド フロントエンド .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md STEERING.md STEERING.md ZEUS.md ZEUS.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
⚡ Zeus — コンピュータを制御するためのワンストップ ツール
Zeus は、実際の UI を備えたローカル AI エージェント オーケストレーターです。一度インストールして zeus start を実行すると、ラップトップや携帯電話上で、ターミナルにインスピレーションを受けた高速な Web アプリが得られます。AI エージェントは実際にマシン上でさまざまなことを行うことができます。つまり、ファイルの読み取りと書き込み、コマンドの実行、実際のブラウザーの操作、Web の検索、タスクの管理、スケジュールされたジョブの実行、サブエージェントのチームの調整などです。コンピューターの制御に必要なすべてを 1 つのツールで実現します。
すべては ~/.zeus に存在します。クラウド アカウント、テレメトリ、ロックインはありません。エージェント、スキル、メモリ、ワークスペースは、独自のディスク上にある人間が編集可能なプレーン ファイルです。
./install.sh # ワンタイムセットアップ
zeus start # サーバーを起動 + Web UI を開く
ゼウスが違う理由
ほとんどのエージェント プロジェクトはライブラリまたは端末ループです。 Zeus は完全なアプリケーションです。
本物の美しい UI を備えています。バニラ JS で、コマンド パレット、テーマ、ライブ ストリーミング、インライン HTML/SVG プレビューを備えたビルド不要の Web アプリです。目を細めて見るような CLI ではありません。
携帯電話上で動作します。この UI は、モバイル ファースト ナビゲーションを備えたインストール可能な PWA で、zeus start --tailscale をプライベート テールネットに配置すると、どこからでもコンピューターのエージェントを安全に、公共のインターネットに公開されることはありません。
エージェントはマシン全体を制御します。

そのための UI。ファイル、シェル、視覚によってエージェントが駆動する専用の自動化された Chromium、Web 検索、組み込みターミナル、ファイル エクスプローラー、コード エディター、git がすべて 1 か所にまとめられています。
それは単なるエージェントではなく、オーケストレーターです。サブエージェントの委任 (サブエージェントごとの停止付き)、定期的なワークフロー、スケジュールされたジョブ、MCP サーバー、ToDo、ダッシュボード、ウィジェット、タイマー、階層化された自己キュレーション メモリ システム。
どのモデルでもお持ちください。 Anthropic、OpenAI、Groq、Google Gemini、OpenRouter、AWS Bedrock — または組み込みの BYO モデル アプリ ( ~/.zeus 内に完全にインストールされる自己完結型の Ollama) を介した完全なローカル モデル。
自己完結型の統合。プライベート Web 検索またはマネージド ブラウザが必要ですか? Zeus は SearXNG、Chromium、Ollama を ~/.zeus 内にインストールします。システムには何も影響しません。
zeus の開始、停止、ステータス、再起動、ログ -f、オープン、リセット。切り離されたローカルサーバー (127.0.0.1:8756 上の uvicorn) を実行し、UI を開きます。 zeus リセットは ~/.zeus をワイプし、工場出荷時の状態で再起動します (最初に確認します)。
http://127.0.0.1:8756 の 1 URL Web UI はストアなしで提供されるため、編集内容はリロード時に表示されます。
オフライン ページ、アプリ ランチャー ナビゲーション、タッチ フレンドリーなモバイル レイアウト (中央にあるポップアップ、大きなタップ ターゲット、水平スクロールなし) を備えたインストール可能な PWA。
Tailscale の組み込み: zeus start --tailscale はテールネット ( https://<machine>.ts.net ) 上でアプリを提供するため、同じテールネットにログインしている携帯電話はどこからでも Zeus をインストールして使用できます。デフォルトでは非公開です。
クラウド: Anthropic (API キー)、OpenAI、Groq、Google Gemini、OpenRouter、AWS Bedrock — キーを貼り付けて実行します。モデルごとのタイプのバッジ (マルチモーダル/テキスト/ツール) と編集可能なコンテキスト ウィンドウを備えた 1 つのモデル スイッチャー。
ローカル — BYO モデル アプリ: 完全な Ollama ライブラリを参照し、気に入ったモデルをお気に入りにし、任意のモデルを取得し (ライブ進行 + 停止ボタン付き)、完全に実行します。

デバイス上で。 Ollama のバイナリと重みは完全にプライベート ポートの ~/.zeus/ollama の下に存在します。システムは変更されません。ビジョン モデル (例: qwen3-vl 、 llava ) は、画像の添付ファイルを確認できます。
独自のプロンプト、ツール、モデル、権限モードを備えた特殊なエージェントを構築します。
エージェントはサブエージェントに委任され (深さは構成可能)、親が動作し続けている間、ライブ ドックから 1 つのサブエージェントを停止できます。
エージェントがオンデマンドでロードするスキル ライブラリ。エージェントは新しいスキルやサブエージェントを作成することもできます。
エージェントが実際に使用できるコンピュータ
組み込みツール: 読み取り / 書き込み / 編集 / Bash / Glob / Grep 、Web 検索 + フェッチ、およびビジョン駆動型ブラウザ (エージェントがスクリーンショットを見てクリック/入力/スクロールする専用の Chromium)。
実際の組み込みターミナル (PTY + xterm.js)、ファイル エクスプローラー、コード エディター、および git パネル。
権限モード — デフォルト (危険なアクションを承認)、計画 (読み取り専用)、 acceptEdits 、自動操縦 — に加えて、 PreToolUse / PostToolUse / Stop / UserPromptSubmit のライフサイクル フック。
ワークフロー: cron スケジュール、オンデマンド、または各チャット後のマルチステップ エージェント パイプライン。
To-Do と付箋、タイマーとアラーム、世界時計、ライブ ウィジェットとダッシュボード、デスクトップ プッシュ通知、組み込みの位置情報ツール。
MCP サーバー (stdio + HTTP/SSE) は、そのツールをエージェントにマージします。
レイヤー化されたコンテキスト (ソウル / ステアリング / ユーザー / プロジェクト) はターンごとにロードされ、プロンプトがキャッシュされるため、ほぼ無料です。
モデル駆動型のキュレーション: モデルは定期的にメモリ自体の重複を排除してプルーニングし、小型で高信号のコアを維持します。プロジェクト メモリは作業ディレクトリに従います。揮発性状態 (どのツール/モデルが接続されているか) は保存されません。
スマートなコンテキスト管理: 動的なモデルごとのコンテキスト ウィンドウ、プロンプト キャッシュ、正確なトークン アカウンティング、および古いツールを消去するインターン コンパクション

出力が長いため、エージェントの実行が時間枠を超えてしまうことはありません。
SearXNG (プライベート メタサーチ)、マネージド Chromium、および Ollama — それぞれは完全に ~/.zeus 内にインストールされて実行され、[設定] → [統合] から制御できます。システム全体には何もインストールされません。
ライブ パフォーマンス モニター、テーマに対応した UI、⌘K コマンド パレット、タブごとの作業ディレクトリ、複数のワークスペース、画像/PDF 添付ファイル、仕様主導の開発ツール。
./install.sh # backend/.venv を作成し、deps をインストールし、`zeus` を PATH にシンボリックリンクします
zeus start # start + http://127.0.0.1:8756 で UI を開きます
zeus start --tailscale # テールネットでも機能します (電話から使用します)
次に、UI を開いてプロバイダーに接続し (または BYO モデルでローカル モデルをインストールし)、チャットを開始します。バックエンドの変更には zeus の再起動が必要です。フロントエンドはビルドせず、リロードするだけです。
要件: Python 3.11 以降および macOS/Linux。状態は ~/.zeus にあります ( ZEUS_HOME でオーバーライドします)。
PewDiePie の Odysseus 、 openclaw 、 hermes に多大な感謝とインスピレーションを与えました。これは、パーソナルでローカルファーストの AI エージェントがどのようなものであるかを示したプロジェクトです。 Zeus は火花を受け取り、独自の方向に走りました。手に持てる完全なアプリです。
そして、Zeus が基盤とするオープン ツールに多大な感謝を示します。
Ollama — ローカル モデル。完全に ~/.zeus 内にインストールされ実行されます。
SearXNG — プライベート メタサーチ、 ~/.zeus で自己ホストされます。
Chromium — エージェントが視覚に基づいて操作する専用ブラウザ。
FastAPI と xterm.js — バックエンドと組み込みターミナル。
そして、Zeus が接続するすべてのモデルプロバイダー — Anthropic、OpenAI、Groq、Google Gemini、OpenRouter、AWS Bedrock 。
Zeus は情熱を持ってバイブコーディングされており、その喜びのために作られています。 ⚡
セルフホスト型 AI エージェント オーケストレーター: ファイル システム + ブラウザー ツール、サブエージェント、スキル、ワークフロー、組み込みブラウザー ブリッジを備えたクロード主導型エージェント - すべてローカル
そこに

読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hosted AI agent orchestrator: Claude-driven agents with filesystem + browser tools, sub-agents, skills, workflows, and a built-in browser bridge — all local - shreyasks094/Zeus

GitHub - shreyasks094/Zeus: Self-hosted AI agent orchestrator: Claude-driven agents with filesystem + browser tools, sub-agents, skills, workflows, and a built-in browser bridge — all local · GitHub
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
shreyasks094
/
Zeus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE backend backend cli cli extension extension frontend frontend .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md STEERING.md STEERING.md ZEUS.md ZEUS.md install.sh install.sh View all files Repository files navigation
⚡ Zeus — your one-stop tool to control your computer
Zeus is a local AI agent orchestrator with a real UI. Install once, run zeus start , and you get a fast, terminal-inspired web app — on your laptop and your phone — where AI agents can actually do things on your machine : read and write files, run commands, drive a real browser, search the web, manage tasks, run scheduled jobs, and orchestrate teams of sub-agents. One tool for everything you need to control your computer.
Everything lives in ~/.zeus . No cloud account, no telemetry, no lock-in — your agents, skills, memory, and workspaces are plain human-editable files on your own disk.
./install.sh # one-time setup
zeus start # launch the server + open the web UI
Why Zeus is different
Most agent projects are a library or a terminal loop . Zeus is a complete application :
It has a real, beautiful UI — a vanilla-JS, no-build web app with a command palette, theming, live streaming, and inline HTML/SVG previews. Not a CLI you squint at.
It runs on your phone. The UI is an installable PWA with a mobile-first nav, and zeus start --tailscale puts it on your private tailnet so you can drive your computer's agents from anywhere — securely, never exposed to the public internet.
Agents control the whole machine, with a UI for it. Files, shell, a dedicated automated Chromium the agent drives by vision, web search, an embedded terminal, a file explorer, a code editor, git — all in one place.
It's an orchestrator, not just an agent. Sub-agent delegation (with per-sub-agent stop), recurring workflows , scheduled jobs, MCP servers, to-dos, dashboards, widgets, timers, and a layered, self-curating memory system.
Bring any model. Anthropic, OpenAI, Groq, Google Gemini, OpenRouter, AWS Bedrock — or fully local models via the built-in BYO Model app (a self-contained Ollama that installs entirely inside ~/.zeus ).
Self-contained integrations. Need private web search or a managed browser? Zeus installs SearXNG, Chromium, and Ollama inside ~/.zeus — nothing touches your system.
zeus start · stop · status · restart · logs -f · open · reset . Runs a detached local server (uvicorn on 127.0.0.1:8756 ) and opens the UI. zeus reset wipes ~/.zeus and starts factory-fresh (asks first).
One-URL web UI at http://127.0.0.1:8756 , served no-store so edits show up on reload.
Installable PWA with an offline page, app-launcher nav, and a touch-friendly mobile layout (centered popups, big tap targets, no horizontal scroll).
Tailscale built in: zeus start --tailscale serves the app on your tailnet ( https://<machine>.ts.net ) so your phone — logged into the same tailnet — can install and use Zeus from anywhere. Private by default.
Cloud: Anthropic (API key), OpenAI, Groq, Google Gemini, OpenRouter, AWS Bedrock — paste a key and go. One model switcher, with per-model type badges (multimodal / text / tools) and editable context windows.
Local — the BYO Model app: browse the full Ollama library, favorite the models you like, pull any model (with live progress + a stop button), and run them fully on-device. The Ollama binary and weights live entirely under ~/.zeus/ollama on a private port — your system stays untouched. Vision models (e.g. qwen3-vl , llava ) can see image attachments.
Build specialized agents with their own prompt, tools, model, and permission mode.
Agents delegate to sub-agents (configurable depth) and you can stop a single sub-agent from the live dock while the parent keeps working.
A skills library the agent loads on demand; agents can even author new skills and sub-agents.
A computer the agent can actually use
Built-in tools: Read / Write / Edit / Bash / Glob / Grep , web search + fetch, and a vision-driven browser (a dedicated Chromium the agent clicks/types/scrolls by looking at screenshots).
A real embedded terminal (PTY + xterm.js), file explorer , code editor , and git panel.
Permission modes — default (approve risky actions), plan (read-only), acceptEdits , autopilot — plus lifecycle hooks for PreToolUse / PostToolUse / Stop / UserPromptSubmit.
Workflows: multi-step agent pipelines on a cron schedule , on demand, or after every chat.
To-dos & stickies , timers & alarms , world clocks , live widgets & dashboards , desktop push notifications , and a built-in location tool.
MCP servers (stdio + HTTP/SSE) merge their tools into your agents.
Layered context — Soul / Steering / User / Project — loaded into every turn and prompt-cached so it's nearly free.
Model-driven curation: a model periodically de-duplicates and prunes memory itself, keeping a small, high-signal core. Project memory follows your working directory; volatile state (which tools/models are connected) is never stored.
Smart context management: dynamic per-model context windows, prompt caching, accurate token accounting, and in-turn compaction that clears stale tool outputs so long agentic runs don't blow past the window.
SearXNG (private metasearch), a managed Chromium, and Ollama — each installs and runs entirely inside ~/.zeus , controllable from Settings → Integrations. Nothing is installed system-wide.
A live performance monitor, themeable UI, a ⌘K command palette, per-tab working directories, multiple workspaces, image/PDF attachments, and spec-driven development tooling.
./install.sh # creates backend/.venv, installs deps, symlinks `zeus` onto PATH
zeus start # start + open the UI at http://127.0.0.1:8756
zeus start --tailscale # also serve on your tailnet (use it from your phone)
Then open the UI, connect a provider (or install a local model in BYO Model ), and start chatting. Backend changes need zeus restart ; the frontend is no-build — just reload.
Requirements: Python 3.11+ and macOS/Linux. State lives in ~/.zeus (override with ZEUS_HOME ).
Huge thanks and inspiration to PewDiePie's Odysseus , openclaw , and hermes — projects that showed what personal, local-first AI agents can be. Zeus took the spark and ran in its own direction: a full app you can hold in your hand.
And enormous thanks to the open tools Zeus stands on:
Ollama — local models, installed and run entirely inside ~/.zeus .
SearXNG — private metasearch, self-hosted in ~/.zeus .
Chromium — the dedicated browser the agent drives by vision.
FastAPI & xterm.js — the backend and the embedded terminal.
And every model provider Zeus connects to — Anthropic, OpenAI, Groq, Google Gemini, OpenRouter, AWS Bedrock .
Zeus was vibe coded with passion — built for the joy of it. ⚡
Self-hosted AI agent orchestrator: Claude-driven agents with filesystem + browser tools, sub-agents, skills, workflows, and a built-in browser bridge — all local
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
