---
source: "https://github.com/lidge-jun/opencodex"
hn_url: "https://news.ycombinator.com/item?id=49012330"
title: "Proxy for OpenAI Codex and Claude Code, use any LLM with those apps"
article_title: "GitHub - lidge-jun/opencodex: Universal provider proxy for OpenAI Codex & Claude Code — use any LLM (Claude, Gemini, Grok, DeepSeek, Ollama…) with Codex CLI, App, SDK, and Claude Code · GitHub"
author: "dandaka"
captured_at: "2026-07-22T20:07:12Z"
capture_tool: "hn-digest"
hn_id: 49012330
score: 2
comments: 0
posted_at: "2026-07-22T19:42:08Z"
tags:
  - hacker-news
  - translated
---

# Proxy for OpenAI Codex and Claude Code, use any LLM with those apps

- HN: [49012330](https://news.ycombinator.com/item?id=49012330)
- Source: [github.com](https://github.com/lidge-jun/opencodex)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T19:42:08Z

## Translation

タイトル: OpenAI Codex および Claude Code のプロキシ、これらのアプリで任意の LLM を使用する
記事のタイトル: GitHub - lipge-jun/opencodex: OpenAI Codex および Claude Code のユニバーサル プロバイダー プロキシ — Codex CLI、アプリ、SDK、および Claude Code で任意の LLM (Claude、Gemini、Grok、DeepSeek、Ollama…) を使用する · GitHub
説明: OpenAI Codex および Claude Code 用のユニバーサル プロバイダー プロキシ — Codex CLI、アプリ、SDK、および Claude Code で任意の LLM (Claude、Gemini、Grok、DeepSeek、Ollama...) を使用します - lidge-jun/opencodex

記事本文:
GitHub - lipge-jun/opencodex: OpenAI Codex および Claude Code 用のユニバーサル プロバイダー プロキシ — Codex CLI、アプリ、SDK、および Claude Code で任意の LLM (Claude、Gemini、Grok、DeepSeek、Ollama…) を使用する · GitHub
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
尾根

-ジュン
/
オープンコデックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,700 コミット 1,700 コミット .github .github アセット アセット bin bin devlog devlog dist/ bin dist/ bin docs-site docs-site docs docs gui gui scripts scripts src src 構造 構造テスト テスト .coderabbit.yaml .coderabbit.yaml .gitattributes .gitattributes .gitignore .gitignore .npmignore .npmignore AGENTS.md AGENTS.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md README.ja.md README.ja.md README.ko.md README.ko.md README.md README.md README.ru.md README.ru.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md bun.lock bun.lock bunfig.toml bunfig.toml package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenAI Codex および Claude Code 用のユニバーサル プロバイダー プロキシ — Codex CLI、アプリ、SDK、および Claude Code で任意の LLM を使用します。
npm install -g @bitkyc08/opencodex · ocx start · localhost:10100
英語 · 한국어 · 简体中文 · Русский · 日本語 · 📖 完全なドキュメント →
Claude、Gemini、Grok、GLM、DeepSeek、Kimi、Qwen、Ollama、またはその他の LLM を、誰かがサポートを追加するのを待つことなく、Codex および Claude Code で使用できます。
opencodex は、Codex の Response API をプロバイダーが話す内容に変換する軽量のローカル プロキシです。ストリーミング、ツール呼び出し、推論トークン、画像など、すべてが双方向で機能します。
Codex、任意のモデルを実行します。プロバイダーを選択してすぐに使用できます。同じ Codex ワークフロー、異なる頭脳です。
Codex 認証用の ChatGPT アカウント プールを管理することもできます。複数の ChatGPT / Codex アカウントを追加し、
ダッシュボードで 5 時間/週/30 日の割り当てを更新し、新しいセッションを許可します

への自動ルート
使用量が最も少ない健全なアカウント。既存の Codex スレッドは、それを開始したアカウントに固定されたままになります。
そのため、SSH、tmux、またはモバイル接続セッションが会話中にアカウントをジャンプすることはありません。
Codex CLI / App / SDK ──/v1/responses──▶ opencodex ──▶ 任意のプロバイダー
│
Anthropic · Google · xAI · キミ · Ollama Cloud · Groq
OpenRouter · Azure · DeepSeek · GLM · …そして OpenAI 自体
フローチャート LR
codex[Codex セッション<br/>CLI、アプリ、SSH、モバイル] --> プロキシ[opencodex]
プロキシ --> 既存{既存のスレッド?}
既存 -->|はい|固定[同じ<br/>ChatGPT アカウントを維持]
既存 -->|新しいセッション|クォータ[クォータを更新<br/>5 時間、毎週、30 日]
クォータ --> 選択[使用量が最も少ない<br/>健全なアカウントを選択]
pick --> 上流 [ChatGPT / Codex バックエンド]
固定 --> 上流
アップストリーム --> 結果[クォータ / 認証結果]
結果 -->|429|クールダウン[クールダウン + フェイルオーバー]
結果 -->|401 / 403|再認証[再認証が必要であるとマーク]
クールダウン --> クォータ
読み込み中
サポートされているプラットフォーム
OS
ステータス
サービスマネージャー
macOS (arm64 / x64)
完全にサポート
起動
Linux (x64 / arm64)
完全にサポート
systemd (ユーザーユニット)
Windows (x64)
完全にサポート
タスク スケジューラ (非表示) / オプトイン ネイティブ サービス ( --native 、WinSW)
ノード 18 以降が必要です。 Bun ランタイムは npm install に自動的にバンドルされます。別途 Bun をインストールする必要はありません。 3 つのプラットフォームはすべてネイティブに動作します (Windows では WSL は必要ありません)。
# インストール (Bun ランタイムを自動的にバンドルします。ノード 18 以降のみが必要です)
# ユーザー所有のノード (nvm/fnm) を優先します — `sudo npm install -g …` は避けてください
npm install -g @bitkyc08/opencodex
# 対話型セットアップ (構成の書き込み、Codex への挿入、および自動開始 shim インストールの提供)
ocx初期化
# プロキシを開始する
ocxスタート
# 初期化中にスキップした場合は、後でオンデマンド自動起動シムをインストールしてください
ocx codex-shim のインストール
# 通常のコーデックスを使用する

y — opencodex 経由でルーティングされるようになりました
codex 「Rust で hello world を書く」
「バンドルされた Bun ランタイムが見つかりません」/ npm が Bun インストール スクリプトをブロックしましたか?
opencodex は Bun ランタイムを依存関係としてバンドルし、ノード経由で実行します。
ランチャーなので、Bun を自分でインストールする必要はありません。を見た場合は、
「バンドルされた Bun ランタイムが見つかりません」エラー、インストールでライフサイクル スクリプトがスキップされる
(allowScripts での bun のポストインストールをブロックする npm を含む) またはオプション
依存関係。これらのフラグを使用せずに再インストールし、bun のインストール スクリプトを許可します。
npm install -g --allow-scripts=bun @bitkyc08/opencodex # --ignore-scripts なし、--omit=optional なし
# 元のインストールで sudo を使用した場合は、引き続き sudo を使用します。
sudo npm install -g --allow-scripts=bun @bitkyc08/opencodex
npm 自身の警告では、パッケージ名を含まない省略されたコマンドが示唆されています。
これは現在のディレクトリを再インストールするため、常に渡します
@bitkyc08/opencodex を明示的に指定します。
sudo を使用して root が所有するプレフィックスにインストールした場合は、上記の sudo を再インストールします
そのプレフィックスのブロックを解除しますが、ユーザー所有のノード (nvm、fnm、または
可能な場合は、ユーザーの npm プレフィックス)。
プロバイダーを追加する最も速い方法は、Web ダッシュボードを使用することです。
ocx gui
これにより、 http://localhost:10100 でダッシュボードが開きます。そこから:
40 を超える組み込みプロバイダーから選択するか、カスタムの OpenAI 互換エンドポイントを入力します
API キーを貼り付けます (または、Anthropic、xAI、および Kim の場合は OAuth 経由でログインします)。
モデルはプロバイダーの /v1/models エンドポイントから自動検出されます
新しいプロバイダーはすぐに使用できるようになります。再起動は必要ありません。
ocx init (対話型 CLI) を使用するか、~/.opencodex/config.json を直接編集してプロバイダーを追加することもできます。
プロバイダー/モデル構文を使用して、構成されたプロバイダーとモデルをターゲットにします。
独自のモデル ID に / が含まれるプロバイダー (zenmux、openrouter、nvidia など) が公開されます。
内側の s を含むコーデックス

- にエイリアスされたまつげ (例: zenmux/moonshotai-kimi-k3-free );の
プロキシはそれらをネイティブ ID に透過的にルーティングし、生のフルスラッシュ形式を維持します。
働いているのも。
# Anthropic を通じてクロード オーパスを使用する
codex -m " anthropic/claude-opus-4-8 " " このスタック トレースの説明 "
# Google 経由で Gemini を使用する
codex -m " google/gemini-3-pro " " auth.ts の単体テストを作成する "
# Ollama Cloud 経由で GLM を使用する
codex -m " ollama-cloud/glm-5.2 " " SQL 移行を作成する "
# Ollama を通じてローカル モデルを使用する
codex -m " ollam/llama3 " " この関数をリファクタリング "
Provider/ プレフィックスを省略すると、opencodex はデフォルトのプロバイダーにルーティングするか、モデル名のパターンに基づいて自動照合します (例: claude-* は Anthropic にルーティングし、gpt-* は OpenAI にルーティングします)。
ルーティングされたモデルは、モデルごとの推論量の制御を備えた Codex App モデル ピッカーにも表示されます。
現在の Codex ビルドでは、 low 、 middle 、 high 、 xhigh 、 max 、および Ultra 推論を公開できます。
モデルがいつそれらをアドバタイズするかを制御します。 opencodex は、プロバイダーを使用しない限り、xhigh と max を区別し続けます。
config は、一方を他方に明示的にマップします。ウルトラ ミラーアップストリーム Codex セマンティクス: 選択します
最大の推論とクライアントでのプロアクティブなマルチエージェント委任を組み合わせたもので、最大値に変換されます。
リクエストがプロバイダーに到達する前に。ルーテッド モデルは、プロバイダー構成が選択した場合にのみそれをアドバタイズします。
推論努力を介して。
GPT-5.6 Sol/Terra/Luna は、OpenAI API キーのロールアウト準備が整ったカタログ エントリとしてシードされ、
OpenRouter プリセット ( gpt-5.6-sol 、 gpt-5.6-terra 、 gpt-5.6-luna ; OpenRouter が使用するもの)
オープンナイ/...)。これらは上流の可用性によってプレビューゲートされたままです。 opencodex は
アカウントおよびそれらにサービスを提供できるプロバイダーのルーティングおよびカタログのメタデータ。
プロバイダーID
ルート
資格情報
行動
オープンナイ
コーデックスログイン
メイン + 追加された Codex アカウント
デフォルトではプール。オプションのダイレクトモード
ああ

ペナイ・アピキー
OpenAI API
APIキー/キープール
Codex アカウントのルーティングがありません
プールには、アフィニティ、クォータ、クールダウン、フェイルオーバーを備えたメインの Codex ログインと追加アカウントが含まれます。
プール状態を直接ショートし、現在の呼び出し元/メインログインベアラーのみを使用します。
永続モードのない新規インストールと構成は、デフォルトでプールになります。でモードを変更します。
ダッシュボードのプロバイダー ページ。どちらのモードでもモデル ID はそのまま残ります。
従来のパブリック プロバイダー ID chatgpt は移行後は非表示になります。元の構成は保持されます
~/.opencodex/config.json.pre-openai-tiers-v2.bak に 1 回。で復元します
cp ~/.opencodex/config.json.pre-openai-tiers-v2.bak ~/.opencodex/config.json 。
現在の構成では openaiProviderTierVersion: 2 を使用します。以前の v1 の 3 プロバイダー構成が移行されます
自動的に単一のオープンアイ列に配置されます。
API 層には、Pro 仮想モデル ( gpt-5.6-sol-pro 、 gpt-5.6-terra-pro 、
gpt-5.6-luna-pro )。ワイヤレベルでは、それぞれが次のように基本モデルを書き換えます。
推論.モード: "プロ" 。
そのカタログは、 gpt-5.5 、 gpt-5.6 、 Sol/Terra/Luna の 8 つの ID に固定されています。
対応する Pro 仮想 ID。汎用の gpt-5.6-pro エイリアスはありません。
コンパクト リクエストは、選択された層を保持しますが、推論オブジェクトなしで基本モデルを送信します。
公式 API メタデータは、1,050,000 のコンテキスト トークンと 922,000 の最大入力トークンです。
設定された openai アカウント モードには gpt-5.6-sol を使用し、
API キーの openai-apikey/gpt-5.6-sol。 Codex ログインと API 認証情報が失敗することはありません
お互いに。
ダッシュボードで Codex Auth を開いてアカウントを追加し、どのアカウントが処理するかを選択します。
次のコーデックスセッション。 opencodex は次の動作を維持します。
既存のセッションはアフィニティを維持します。スレッド ID は選択したアカウントにバインドされ、次のアカウントで再利用されます。
そのため、長いリクエストまたはモバイル/SSH 接続セッションは同じアカウントを使用し続けます。

新しいセッションは自動ルーティングできます。自動切り替えが有効になっている場合、opencodex は既知の最もホットなものを比較します。
5 時間、毎週、30 日の使用量にわたるクォータ ウィンドウを設定し、新しいアカウントに対して使用量の少ない適格なアカウントを選択します。
アクティブなアカウントがしきい値を超えると、セッションが停止されます。
クォータ検索が組み込まれています。ダッシュボードでは、ワンクリックですべてのアカウント クォータを更新できます。
リクエスト ログ ラベルは、非 PII アカウント序数を含むトラフィックをプールします。
障害はフェイルクローズされます。トークンの失敗は、別のトークンにフォールバックするのではなく、再認証をマークします。
資格情報をサイレントに送信します。 429 クォータ応答によりアカウントがクールダウンになり、今後の作業がフェイルオーバーされる可能性があります
別の適格なプールアカウントに転送します。
Codex では任意の LLM を使用します。 5 つのプロトコル アダプターは、Anthropic Messages、Google Gemini、Azure、OpenAI Responses パススルー、およびすべての OpenAI 互換 Chat Completions エンドポイントをカバーします。つまり、すぐに使用できるプロバイダーは 40 以上です。
クロード コードでも任意の LLM を使用します。同じデーモンが Anthropic Messages API ( /v1/messages + count_tokens ) を提供します。ocx claude は完全に接続された Claude Code を起動し、ルーティングされたモデルはゲートウェイ モデル検出 (claude-ocx-<provider>--<model> エイリアス、Claude Code 2.1.129+) 経由でネイティブの /model ピッカーに表示されます。ダッシュボードのクロード ページでスロットとモデル マップを設定します。
ChatGPT アカウントを安全にプールします。新しいセッション中は、既存の Codex スレッドを 1 つのアカウントに維持します
自動化できます

[切り捨てられた]

## Original Extract

Universal provider proxy for OpenAI Codex & Claude Code — use any LLM (Claude, Gemini, Grok, DeepSeek, Ollama…) with Codex CLI, App, SDK, and Claude Code - lidge-jun/opencodex

GitHub - lidge-jun/opencodex: Universal provider proxy for OpenAI Codex & Claude Code — use any LLM (Claude, Gemini, Grok, DeepSeek, Ollama…) with Codex CLI, App, SDK, and Claude Code · GitHub
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
lidge-jun
/
opencodex
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,700 Commits 1,700 Commits .github .github assets assets bin bin devlog devlog dist/ bin dist/ bin docs-site docs-site docs docs gui gui scripts scripts src src structure structure tests tests .coderabbit.yaml .coderabbit.yaml .gitattributes .gitattributes .gitignore .gitignore .npmignore .npmignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md README.ja.md README.ja.md README.ko.md README.ko.md README.md README.md README.ru.md README.ru.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md bun.lock bun.lock bunfig.toml bunfig.toml package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Universal provider proxy for OpenAI Codex & Claude Code — use any LLM with Codex CLI, App, SDK, and Claude Code.
npm install -g @bitkyc08/opencodex · ocx start · localhost:10100
English · 한국어 · 简体中文 · Русский · 日本語 · 📖 Full documentation →
Use Claude, Gemini, Grok, GLM, DeepSeek, Kimi, Qwen, Ollama, or any other LLM with Codex — and with Claude Code — without waiting for anyone to add support.
opencodex is a lightweight local proxy that translates Codex's Responses API into whatever your provider speaks. Streaming, tool calls, reasoning tokens, images — everything works, in both directions.
Codex, running any model. Pick a provider and go — same Codex workflow, different brain.
It can also manage a ChatGPT account pool for Codex auth. Add multiple ChatGPT / Codex accounts,
refresh their 5h / weekly / 30d quota in the dashboard, and let new sessions auto-route to the
lowest-usage healthy account. Existing Codex threads stay pinned to the account that started them,
so long SSH, tmux, or mobile-connected sessions do not jump accounts mid-conversation.
Codex CLI / App / SDK ──/v1/responses──▶ opencodex ──▶ Any provider
│
Anthropic · Google · xAI · Kimi · Ollama Cloud · Groq
OpenRouter · Azure · DeepSeek · GLM · …and OpenAI itself
flowchart LR
codex[Codex session<br/>CLI, App, SSH, mobile] --> proxy[opencodex]
proxy --> existing{Existing thread?}
existing -->|yes| pinned[Keep the same<br/>ChatGPT account]
existing -->|new session| quota[Refresh quota<br/>5h, weekly, 30d]
quota --> pick[Pick lowest-usage<br/>healthy account]
pick --> upstream[ChatGPT / Codex backend]
pinned --> upstream
upstream --> outcomes[Quota / auth outcome]
outcomes -->|429| cooldown[Cooldown + failover]
outcomes -->|401 / 403| reauth[Mark reauth needed]
cooldown --> quota
Loading
Supported platforms
OS
Status
Service manager
macOS (arm64 / x64)
Fully supported
launchd
Linux (x64 / arm64)
Fully supported
systemd (user unit)
Windows (x64)
Fully supported
Task Scheduler (hidden) / opt-in native service ( --native , WinSW)
Requires Node 18+. The Bun runtime is bundled automatically on npm install — no separate Bun install needed. All three platforms work natively (no WSL needed on Windows).
# Install (bundles the Bun runtime automatically — only Node 18+ required)
# Prefer a user-owned Node (nvm/fnm) — avoid `sudo npm install -g …`
npm install -g @bitkyc08/opencodex
# Interactive setup (writes config, injects into Codex, and offers autostart shim install)
ocx init
# Start the proxy
ocx start
# If you skipped it during init, install the on-demand autostart shim later
ocx codex-shim install
# Use Codex normally — it now routes through opencodex
codex " Write a hello world in Rust "
"bundled Bun runtime is missing" / npm blocked Bun install scripts?
opencodex bundles the Bun runtime as a dependency and runs it via a Node
launcher, so you do not need to install Bun yourself. If you see a
"bundled Bun runtime is missing" error, the install skipped lifecycle scripts
(including npm blocking bun's postinstall under allowScripts ) or optional
dependencies. Reinstall without those flags, allowing bun's install script:
npm install -g --allow-scripts=bun @bitkyc08/opencodex # no --ignore-scripts, no --omit=optional
# if the original install used sudo, keep using sudo:
sudo npm install -g --allow-scripts=bun @bitkyc08/opencodex
npm's own warning suggests an abbreviated command without the package name —
that would reinstall the current directory, so always pass
@bitkyc08/opencodex explicitly.
If you installed with sudo into a root-owned prefix, the sudo reinstall above
unblocks that prefix — but prefer migrating to a user-owned Node (nvm, fnm, or
a user npm prefix) when you can.
The fastest way to add a provider is through the web dashboard:
ocx gui
This opens the dashboard at http://localhost:10100 . From there:
Pick from 40+ built-in providers — or enter a custom OpenAI-compatible endpoint
Paste your API key (or log in via OAuth for Anthropic, xAI, and Kimi)
Models are auto-discovered from the provider's /v1/models endpoint
Your new provider is ready to use immediately. No restart needed.
You can also add providers through ocx init (interactive CLI) or by editing ~/.opencodex/config.json directly.
Target any configured provider and model using the provider/model syntax:
Providers whose own model ids contain / (zenmux, openrouter, nvidia, …) are exposed to
Codex with inner slashes aliased to - (e.g. zenmux/moonshotai-kimi-k3-free ); the
proxy transparently routes them back to the native id, and the raw full-slash form keeps
working too.
# Use Claude Opus through Anthropic
codex -m " anthropic/claude-opus-4-8 " " Explain this stack trace "
# Use Gemini through Google
codex -m " google/gemini-3-pro " " Write unit tests for auth.ts "
# Use GLM through Ollama Cloud
codex -m " ollama-cloud/glm-5.2 " " Write a SQL migration "
# Use a local model through Ollama
codex -m " ollama/llama3 " " Refactor this function "
When you omit the provider/ prefix, opencodex routes to the default provider — or auto-matches based on the model name pattern (e.g., claude-* routes to Anthropic, gpt-* routes to OpenAI).
Routed models also appear in the Codex App model picker with per-model reasoning effort controls:
Current Codex builds can expose low , medium , high , xhigh , max , and ultra reasoning
controls when a model advertises them. opencodex keeps xhigh and max distinct unless a provider
config explicitly maps one to the other. ultra mirrors upstream Codex semantics: it selects
maximum reasoning plus proactive multi-agent delegation in the client, and is converted to max
before any request reaches a provider. Routed models advertise it only when a provider config opts
in via reasoningEfforts .
GPT-5.6 Sol/Terra/Luna are seeded as rollout-ready catalog entries for the OpenAI API key and
OpenRouter presets ( gpt-5.6-sol , gpt-5.6-terra , gpt-5.6-luna ; OpenRouter uses
openai/... ). They remain preview-gated by upstream availability; opencodex only prepares the
routing and catalog metadata for accounts and providers that can serve them.
Provider ID
Route
Credential
Behavior
openai
Codex login
Main + added Codex accounts
Pool by default; optional Direct mode
openai-apikey
OpenAI API
API key/key pool
No Codex account routing
Pool includes the main Codex login and added accounts, with affinity, quota, cooldown, and failover.
Direct short-circuits pool state and uses only the current caller/main-login bearer.
Fresh installs and configs with no persisted mode default to Pool. Change the mode on the
dashboard's Providers page; model ids stay bare in either mode.
The legacy public provider id chatgpt is hidden after migration. The original config is retained
once at ~/.opencodex/config.json.pre-openai-tiers-v2.bak ; restore it with
cp ~/.opencodex/config.json.pre-openai-tiers-v2.bak ~/.opencodex/config.json .
Current configs use openaiProviderTierVersion: 2 . Earlier v1 three-provider configs migrate
automatically into the single openai row.
The API tier includes Pro virtual models ( gpt-5.6-sol-pro , gpt-5.6-terra-pro ,
gpt-5.6-luna-pro ). At the wire level, each rewrites to its base model with
reasoning.mode: "pro" .
Its catalog is fixed to eight ids: gpt-5.5 , gpt-5.6 , Sol/Terra/Luna, and the three
corresponding Pro virtual ids. There is no generic gpt-5.6-pro alias.
Compact requests keep the selected tier but send the base model without a reasoning object.
Official API metadata is 1,050,000 context tokens and 922,000 max input tokens.
Use gpt-5.6-sol for the configured openai account mode and
openai-apikey/gpt-5.6-sol for the API key. Codex-login and API credentials never fall through to
one another.
Open Codex Auth in the dashboard to add accounts and choose which account should handle the
next Codex session. opencodex keeps these behaviors:
Existing sessions keep affinity. A thread id is bound to the selected account and reused on
later turns, so a long request or a mobile/SSH-attached session keeps using the same account.
New sessions can auto-route. When auto-switch is enabled, opencodex compares the hottest known
quota window across 5h, weekly, and 30d usage, then picks a lower-usage eligible account for new
sessions once the active account crosses the threshold.
Quota lookup is built in. The dashboard can refresh all account quotas in one click, and the
request log labels pool traffic with non-PII account ordinals.
Failures fail closed. Token failures mark reauthentication instead of falling back to another
credential silently; 429 quota responses put the account in cooldown and can fail over future work
to another eligible pool account.
Use any LLM with Codex. 5 protocol adapters cover Anthropic Messages, Google Gemini, Azure, OpenAI Responses passthrough, and every OpenAI-compatible Chat Completions endpoint — that's 40+ providers out of the box.
Use any LLM with Claude Code too. The same daemon serves the Anthropic Messages API ( /v1/messages + count_tokens ): ocx claude launches Claude Code fully wired, and routed models appear in its native /model picker via gateway model discovery ( claude-ocx-<provider>--<model> aliases, Claude Code 2.1.129+). Configure slots and model maps on the dashboard's Claude page.
Pool ChatGPT accounts safely. Keep existing Codex threads on one account while new sessions
can auto

[truncated]
