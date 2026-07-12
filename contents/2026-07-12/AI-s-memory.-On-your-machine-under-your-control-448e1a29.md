---
source: "https://github.com/EXXETA/exxperts"
hn_url: "https://news.ycombinator.com/item?id=48884002"
title: "AI's memory. On your machine, under your control"
article_title: "GitHub - EXXETA/exxperts: Persistent AI rooms with governed, approval-gated memory: local-first, on your machine. · GitHub"
author: "alindnbrg"
captured_at: "2026-07-12T20:02:22Z"
capture_tool: "hn-digest"
hn_id: 48884002
score: 1
comments: 1
posted_at: "2026-07-12T19:44:13Z"
tags:
  - hacker-news
  - translated
---

# AI's memory. On your machine, under your control

- HN: [48884002](https://news.ycombinator.com/item?id=48884002)
- Source: [github.com](https://github.com/EXXETA/exxperts)
- Score: 1
- Comments: 1
- Posted: 2026-07-12T19:44:13Z

## Translation

タイトルは「AIの記憶」。あなたのマシン上で、あなたの制御下で
記事のタイトル: GitHub - EXXETA/experts: 管理された承認ゲート型メモリを備えた永続的な AI ルーム: マシン上でローカルファースト。 · GitHub
説明: ローカルファーストの、管理された承認ゲート型メモリを備えた永続的な AI ルーム。 - EXXETA/エキスパート

記事本文:
GitHub - EXXETA/experts: 管理された承認ゲート型メモリを備えた永続的な AI ルーム: マシン上でローカルファースト。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
エクセタ
/
専門家
公共
通知
あなたはきっとそうでしょう

サインインして通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github .github apps apps bin bin docs docs pi-package pi-package runtime ランタイム スクリプト scripts .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cffライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI の記憶。あなたのマシン上であなたの制御下にあります。
専門家は、管理された承認ゲート型メモリを備えた永続的な AI ルームを提供します。エージェントは、セッション全体にわたってあなたとあなたの作業を記憶し、すべてのメモリ書き込みがあなたによって承認されます。すべてがローカルで実行され、ローカルに維持されます。ルーム、メモリ、ナレッジ ベース、アーティファクト、および使用状況データはディスク上のファイルです。
あなたが支配する記憶。 Room はセッション全体にわたる決定、好み、コンテキストを記憶しますが、すべての記憶の書き込みはユーザーが制御する承認ゲートを通過します。黙ってプロフィールを構築する必要はありません。
確実にローカルファースト。ルーム、思い出、知識ベース、アーティファクト、およびトークンの使用状況は、 ~/.experts の下にプレーン ファイルとして存在します。すべてを読み取り、バックアップ、または削除できます。テレメトリはありません。ネットワーク トラフィックは使用するものだけです。つまり、モデル プロバイダーへの呼び出し、Web リサーチ (有効にした場合)、および追加した MCP コネクタが含まれます。
チャット ラッパーではなく、エージェント ランタイム。ルームは、各サーフェスを対象とした権限モデルに基づいて、厳選されたツール (ナレッジ ベース、アーティファクト、ローカル Web リサーチ、MCP コネクタ) を使用します。部屋が無制限のシェルになることはありません。
1 つの製品で 2 つの表面。 Web アプリと CLI/TUI は、同じルーム、メモリ、資格情報ストアを共有します。自由に切り替えてください。
Open WebUI または AnythingLLM を知っている場合:

re excellent local chat/RAG frontends.エキスパートは異なる形状です。ドキュメント上のチャットではなく、永続的で管理されたメモリと監査可能なエージェントの動作に重点を置いたローカル エージェント ランタイムです。
表面
打ち上げ
得られるもの
ウェブアプリ
専門家のウェブ
メモリ、KB、アーティファクト、Web リサーチ、承認、ウォレットのある部屋。
CLI/TUI
専門家 cli
ターミナル内の同じルームを、ワークスペースとして使用するフォルダーから実行します。
Web アプリは完全な製品です。AI のセットアップ、メモリのレビューと承認、ウォレット、コネクタ、スキルはすべてそこに存在します。 The CLI/TUI focuses on the rooms themselves. Bare exxperts opens a picker for the two surfaces (web app recommended).製品/アプリの状態は ~/.exxperts/app (Windows の場合は %USERPROFILE%\.experts) でローカルに維持されます。 embedded runtime provider/auth/model/session state lives under ~/.exxperts/agent .
ランディング ページ、永続ルーム、承認、ウォレットを備えたローカル Web ワークスペース。
同じルームのランタイムとガバナンスを共有するルーム専用の CLI/TUI。
User-created skills from the web UI.
承認ゲート型メモリ、KB 書き込み、Markdown/HTML アーティファクト。
Markdown/Obsidian KB tools and local web search.
MCP connectors on web and CLI through a single proxy tool.独自のサーバーを持ち込んでください。 docs/mcp.md を参照してください。
~/.exxperts/app/usage.jsonl からのローカル トークン/コスト ウォレット。
Prerequisites: Node.js 20.6+ and npm. On Windows, apply the two Git settings from the Windows quickstart before cloning.
git clone https://github.com/EXXETA/exxperts.git
CDの専門家
npmインストール
npm run install:global # エキスパート コマンドをビルド、パック、インストールします (すべてのプラットフォーム)
期待できることは 2 つあります: npm install はヘッドレス Chromium も取得します (最大 150 MB、1 回。EXXETA_SKIP_BROWSER_INSTALL=1 npm install でスキップします)、および install:global は @exxeta/exxperts-app int をインストールする前にアプリ全体をビルドします

o グローバル npm プレフィックスなので、数分お待ちください。
エキスパート Web # Web アプリ: ルーム、メモリ、ウォレット (現在のフォルダーは関係ありません)
cd /path/to/your/project
Experts cli # このフォルダーをワークスペースとして、ターミナル内の同じ部屋に配置します
First run: open AI setup in the web app and sign in to your provider (Claude and ChatGPT subscriptions sign in with one click; API keys and OpenAI-compatible gateways also work; see Model/provider setup ).何かが機能していないのでしょうか?リポジトリのルートから npm run Doctor はすべてのレイヤーをチェックし、修正を出力します。
ここは新しいですか？ docs/quickstart.md は、AI のインストール、接続、最初の部屋、最初のメモリまでのパス全体を約 5 分で実行します。製品が何であるか、および各部分がどのように適合するかについては、 docs/how-experts-works.md を参照してください。
どのプラットフォームでも同じで、リポジトリフォルダーから:
git プル
npm install # 依存関係が変更された場合に備えて
npm run install:global # グローバルコマンドを再構築して再インストールします
その後、グローバル エキスパート コマンドが新しいバージョンを実行します。アップデート後またはいつでも何か問題が発生した場合は、リポジトリ フォルダーから npm run Doctor を実行してください。グローバル インストールではなくクローンから開発しますか? git pull && npm install && npm run build で更新します。
Windows は、Web アプリと CLI/TUI の両方でサポートされています。要件:
Git for Windows ≥ 2.40 ( https://gitforwindows.org )。 Git Bash is required : the agent's shell tool runs commands through bash.exe , which is discovered automatically in the standard Git for Windows install location ( C:\Program Files\Git\bin\bash.exe ) or on PATH .
Node.js 20.6+ (LTS 推奨) および npm ( https://nodejs.org )。
CLI/TUI には Windows ターミナルを推奨します (レガシー conhost はテストされていません)。
クローン作成前の 1 回限りの Git 設定 (node_modules ツリーが 260 文字の MAX_PATH を超えるため、長いパスが重要です):
git

config -- グローバル core.longpaths true
git config -- global core.autocrlf false # リポジトリの .gitattributes が行末を管理します
次に、他の場所と同じコマンドを使用して、PowerShell または Git Bash からインストールします。
git clone https://github.com/EXXETA/experts.git
CDの専門家
npmインストール
npm run install:グローバル
任意のシェルから実行します。
エキスパート Web # Web アプリ
experts cli # CLI/TUI、ワークスペースとして使用するフォルダーから実行
Web 検索は Windows でも機能します。Docker Desktop をインストールし、任意のシェルからノード スクリプト\searxng.mjs start を 1 回実行します。 docs/web-search.md を参照してください。
グローバル インストールを行わずにクローンから開発しますか?シェルに依存しないフォームを使用します: ノード bin\exxperts-web.cjs 、ノード bin\experts-cli.cjs 、およびノー​​ド scripts\exxeta-web.mjs (サーバー + Vite UI を備えた開発 Web アプリ)。 scripts/ の bash ランチャーは Git Bash からも動作します。
フル機能を利用するには何をインストールするか
すべてがローカルで実行されます。完全な機能は 4 つのレイヤーです。最初のもののみが必要です:
コア アプリ: Node.js 20.6+ および npm、その後、上記のクイックスタート手順を実行します。
Headless Chromium (~150 MB、1 回限り): npm インストール中に自動的にダウンロードされます。ルームは作成した HTML デッキを視覚的に確認し、JavaScript でレンダリングされたページを読むことができます。ダウンロードを実行できなかった場合は、後で npx playwright install chromium を使用して有効にします (または、 EXXETA_SKIP_BROWSER_INSTALL=1 npm install を使用してインストール中にスキップします)。
Web 検索: コンテナ エンジンと 1 回限りのセットアップ コマンド。 「Web 検索 (オプション)」を参照してください。
モデル認証: プロバイダーのサインインまたは API キー。 「モデル/プロバイダーのセットアップ」を参照してください。
リポジトリ ルートから npm run Doctor を使用して設定を確認します。上記のすべてに加えて、MCP 構成がチェックされ、アウトバウンド Web フェッチが正しくデコードされていること (企業の TLS 検査プロキシは応答を破損する可能性があります) がチェックされ、欠落しているものについては修正が出力されます。
インストール:グローバル W

raps npm run build && npm Pack && npm install -g <tarball> ;手動手順と npm exec による 1 回限りの実行 (グローバル インストールなし) は docs/packaging-local.md に文書化されています。 macOS が EACCES を返す場合は、 sudo の代わりにユーザーレベルの npm プレフィックスを使用します。それもそこでカバーされています。新しい npm がインストール スクリプトがブロックされたことを警告する場合 (allow-scripts )、アプリは引き続き動作します。ブロックされたステップは、Chromium のダウンロード (npx playwright install chromium でいつでも回復可能) と、MCP ページの化粧品のブランディングです。
Web 検索は出荷時に無効になっており、SearXNG コンテナを通じて完全にローカルで実行されます。API キーやサードパーティの検索 SaaS はありません。注: 検索クエリは引き続き公開検索エンジンに送信されるため、機密クライアントや内部コンテンツの検索は避けてください。
有効にするには: Docker Desktop または (macOS では軽量) OrbStack をインストールし、リポジトリ ディレクトリからワンタイム セットアップ コマンドを実行してアプリを再起動します。
./scripts/searxng start # macOS / Linux / Git Bash
ノード スクリプト \s Earxng.mjs start # Windows (PowerShell または cmd)
セットアップの詳細、再起動後の実行維持、ステータス/停止コマンド、および Windows のメモ: docs/web-search.md 。
ルームが応答するには、サインインしているモデル プロバイダーが必要です。適合するパスを選択してください。
サブスクリプション プロバイダー (Claude、ChatGPT Plus/Pro) の場合は、Web アプリを起動し、 AI セットアップ を開き、プロバイダーのプロファイル カードの [サインイン →] ボタンを使用します。プロバイダーのログインはブラウザーで開き、資格情報はローカルの資格情報ストアに残ります。
OpenAI 互換ゲートウェイ (会社の LiteLLM または vLLM プロキシなど) の場合、すべてが Web アプリ内で行われます。AI セットアップを開き、→ 別のプロバイダーを追加→ ゲートウェイをセットアップし、ベース URL とゲートウェイ ルートのモデル ID を入力し、学習/レビュー メモリ モデルを選択して、ゲートウェイのプロファイル行にトークンを貼り付けます。ターミナルウィザードはまだ

好みに応じて機能します:
専門家のウェブ
エキスパートのセットアップ openai 互換
リポジトリ クローンから実行する場合は、グローバルにインストールされたコマンドの代わりに、対応する dev setup コマンドを使用します。
./scripts/experts-web
./scripts/experts-cli セットアップ openai 互換
ランタイムが認識している他のプロバイダー (Google Gemini、Groq、Mistral、DeepSeek、OpenRouter、xAI、その他約 25 個) を Web アプリから追加できます。AI セットアップを開いて [別のプロバイダーの追加] を使用し、プロバイダーが提供するサブスクリプションでサインインするか、API キーを貼り付けます。サインインした後、プロバイダーが使用できるモデルを承認します。ルームで使用できるモデルと、Learn and Review Memory を実行するモデル (デフォルトが推奨されます) です。承認によりプロバイダーの AI プロファイルが作成されます。これがないと、プロバイダーはサインインしていますが、ルームでは使用できません。
シェルまたはリポジトリ .env 内のプロバイダー API キーは、クローンから実行する場合にも機能します。
OPENAI_API_KEY=...
ANTHROPIC_API_KEY=...
サブスクリプション/OAuth ログインはターミナルからも機能します。 exxperts cli を実行し、 /login と入力します。 API キー エントリを含む、Web アプリと同じプロバイダー リストを提供します。 Web アプリと CLI/TUI は同じローカル ランタイム資格情報ストアを共有するため、どちらのパスでも両方にサインインします。新しく追加されたprovのモデルを承認する

[切り捨てられた]

## Original Extract

Persistent AI rooms with governed, approval-gated memory: local-first, on your machine. - EXXETA/exxperts

GitHub - EXXETA/exxperts: Persistent AI rooms with governed, approval-gated memory: local-first, on your machine. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
EXXETA
/
exxperts
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github .github apps apps bin bin docs docs pi-package pi-package runtime runtime scripts scripts .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Your AI's memory. On your machine, under your control.
exxperts gives you persistent AI rooms with governed, approval-gated memory: agents that remember you and your work across sessions, with every memory write approved by you. Everything runs and stays local: rooms, memory, knowledge base, artifacts, and usage data are files on your disk.
Memory you govern. Rooms remember decisions, preferences, and context across sessions, but every memory write goes through an approval gate you control. No silent profile-building.
Local-first, verifiably. Rooms, memories, knowledge base, artifacts, and token usage live as plain files under ~/.exxperts . You can read, back up, or delete all of it. There is no telemetry: the only network traffic is what you use, meaning calls to your model provider, web research if you enable it, and any MCP connectors you add.
An agentic runtime, not a chat wrapper. Rooms use curated tools (knowledge base, artifacts, local web research, MCP connectors) under a permission model scoped to each surface. Rooms never get an unrestricted shell.
One product, two surfaces. The web app and the CLI/TUI share the same rooms, memory, and credential store; switch between them freely.
If you know Open WebUI or AnythingLLM : those are excellent local chat/RAG frontends. exxperts is a different shape: a local agentic runtime focused on persistent, governed memory and auditable agent behaviour, not chat over your documents.
Surface
Launch
What you get
Web app
exxperts web
Rooms with memory, KB, artifacts, web research, approvals, and the wallet.
CLI / TUI
exxperts cli
The same rooms in your terminal, run from the folder you want as the workspace.
The web app is the full product: AI setup, memory review and approvals, the wallet, connectors, and skills all live there. The CLI/TUI focuses on the rooms themselves. Bare exxperts opens a picker for the two surfaces (web app recommended). Product/app state stays local under ~/.exxperts/app ( %USERPROFILE%\.exxperts on Windows); embedded runtime provider/auth/model/session state lives under ~/.exxperts/agent .
Local web workspace with landing page, persistent rooms, approvals, and wallet.
Rooms-only CLI/TUI sharing the same room runtime and governance.
User-created skills from the web UI.
Approval-gated memory, KB writes, and Markdown/HTML artifacts.
Markdown/Obsidian KB tools and local web search.
MCP connectors on web and CLI through a single proxy tool. Bring your own servers; see docs/mcp.md .
Local token/cost wallet from ~/.exxperts/app/usage.jsonl .
Prerequisites: Node.js 20.6+ and npm. On Windows, apply the two Git settings from the Windows quickstart before cloning.
git clone https://github.com/EXXETA/exxperts.git
cd exxperts
npm install
npm run install:global # builds, packs, and installs the exxperts commands (all platforms)
Two things to expect: npm install also fetches headless Chromium (~150 MB, one-time; skip it with EXXETA_SKIP_BROWSER_INSTALL=1 npm install ), and install:global builds the whole app before installing @exxeta/exxperts-app into your global npm prefix, so give it a few minutes.
exxperts web # the web app: rooms, memory, wallet (current folder does not matter)
cd /path/to/your/project
exxperts cli # the same rooms in your terminal, with this folder as the workspace
First run: open AI setup in the web app and sign in to your provider (Claude and ChatGPT subscriptions sign in with one click; API keys and OpenAI-compatible gateways also work; see Model/provider setup ). Something not working? npm run doctor from the repo root checks every layer and prints the fix.
New here? docs/quickstart.md walks the whole path in about five minutes: install, connect your AI, first room, first memory. For an orientation on what the product is and how the pieces fit, read docs/how-exxperts-works.md .
Same on every platform, from the repo folder:
git pull
npm install # in case dependencies changed
npm run install:global # rebuilds and reinstalls the global commands
The global exxperts commands then run the new version. If anything misbehaves after an update, or anytime, run npm run doctor from the repo folder. Developing from the clone instead of the global install? Update with git pull && npm install && npm run build .
Windows is supported for both the web app and the CLI/TUI. Requirements:
Git for Windows ≥ 2.40 ( https://gitforwindows.org ). Git Bash is required : the agent's shell tool runs commands through bash.exe , which is discovered automatically in the standard Git for Windows install location ( C:\Program Files\Git\bin\bash.exe ) or on PATH .
Node.js 20.6+ (LTS recommended) and npm ( https://nodejs.org ).
Windows Terminal recommended for the CLI/TUI (legacy conhost is untested).
One-time Git settings before cloning (long paths matter because node_modules trees exceed the 260-character MAX_PATH ):
git config -- global core.longpaths true
git config -- global core.autocrlf false # the repo's .gitattributes manages line endings
Then install from PowerShell or Git Bash, with the same commands as everywhere else:
git clone https: // github.com / EXXETA / exxperts.git
cd exxperts
npm install
npm run install:global
And run from any shell:
exxperts web # web app
exxperts cli # CLI/TUI, run from the folder you want as workspace
Web search works on Windows too: install Docker Desktop, then run node scripts\searxng.mjs start once from any shell. See docs/web-search.md .
Developing from the clone without a global install? Use the shell-independent forms: node bin\exxperts-web.cjs , node bin\exxperts-cli.cjs , and node scripts\exxeta-web.mjs (dev web app with server + Vite UI). The bash launchers in scripts/ also work from Git Bash.
What to install for full functionality
Everything runs locally. Full functionality is four layers; only the first is required:
Core app : Node.js 20.6+ and npm, then the quick-start steps above.
Headless Chromium (~150 MB, one-time) : downloaded automatically during npm install ; lets rooms visually review the HTML decks they author and read JavaScript-rendered pages. If the download couldn't run, enable it later with npx playwright install chromium (or skip it during install with EXXETA_SKIP_BROWSER_INSTALL=1 npm install ).
Web search : a container engine plus a one-time setup command. See Web search (optional) .
Model authentication : provider sign-in or API keys. See Model/provider setup .
Verify any setup with npm run doctor from the repo root: it checks all of the above, plus MCP config and that outbound web fetches decode cleanly (corporate TLS-inspection proxies can corrupt responses), and prints the fix for anything missing.
install:global wraps npm run build && npm pack && npm install -g <tarball> ; the manual steps and one-off runs via npm exec (no global install) are documented in docs/packaging-local.md . If macOS returns EACCES , use a user-level npm prefix instead of sudo ; that is also covered there. If a newer npm warns that install scripts were blocked ( allow-scripts ), the app still works. The blocked steps are the Chromium download (recoverable anytime with npx playwright install chromium ) and cosmetic MCP page branding.
Web search ships disabled and runs fully locally through a SearXNG container: no API key, no third-party search SaaS. Note: your search queries still go to public search engines, so avoid searching confidential client or internal content.
To enable: install Docker Desktop or (lighter on macOS) OrbStack , then from the repo directory run the one-time setup command and restart the app:
./scripts/searxng start # macOS / Linux / Git Bash
node scripts \s earxng.mjs start # Windows (PowerShell or cmd)
Setup details, keeping it running across reboots, status/stop commands, and Windows notes: docs/web-search.md .
Rooms need a signed-in model provider before they can respond. Pick whichever path fits.
For subscription providers (Claude, ChatGPT Plus/Pro), launch the web app, open AI setup , and use the Sign in → button on the provider's profile card. The provider login opens in the browser, and credentials stay in the local credential store.
For an OpenAI-compatible gateway (a company LiteLLM or vLLM proxy, for example), everything happens in the web app: open AI setup → Add another provider → Set up gateway , enter the base URL and the model ids your gateway routes, pick the Learn/Review Memory model, then paste your token on the gateway's profile row. The terminal wizard still works if you prefer it:
exxperts web
exxperts setup openai-compatible
When running from a repo clone, use the matching dev setup command instead of any globally installed command:
./scripts/exxperts-web
./scripts/exxperts-cli setup openai-compatible
Any other provider the runtime knows (Google Gemini, Groq, Mistral, DeepSeek, OpenRouter, xAI, and ~25 more) can be added from the web app: open AI setup and use Add another provider , then sign in with a subscription where the provider offers one, or paste an API key. After signing in, approve the models that provider may use: the models available in rooms, plus the one that runs Learn and Review Memory (a default is suggested). Approval creates the provider's AI profile; without it, the provider is signed in but not usable in rooms.
Provider API keys in your shell or repo .env also work when running from the clone:
OPENAI_API_KEY=...
ANTHROPIC_API_KEY=...
Subscription/OAuth login also works from the terminal: run exxperts cli , then type /login . It offers the same provider list as the web app, including API-key entry. The web app and CLI/TUI share the same local runtime credential store, so either path signs in both. Approving models for a newly added prov

[truncated]
