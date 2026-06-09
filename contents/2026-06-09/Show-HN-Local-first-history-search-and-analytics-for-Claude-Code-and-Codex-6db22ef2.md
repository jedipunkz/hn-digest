---
source: "https://github.com/sudomichael/agentgraphed"
hn_url: "https://news.ycombinator.com/item?id=48462489"
title: "Show HN: Local-first history, search, and analytics for Claude Code and Codex"
article_title: "GitHub - sudomichael/agentgraphed: Local-first analytics dashboard for AI coding sessions. See what you built with Claude Code & Codex CLI. · GitHub"
author: "ushercakes"
captured_at: "2026-06-09T16:03:50Z"
capture_tool: "hn-digest"
hn_id: 48462489
score: 1
comments: 0
posted_at: "2026-06-09T15:35:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Local-first history, search, and analytics for Claude Code and Codex

- HN: [48462489](https://news.ycombinator.com/item?id=48462489)
- Source: [github.com](https://github.com/sudomichael/agentgraphed)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T15:35:07Z

## Translation

タイトル: HN を表示: クロード コードとコーデックスのローカル優先の履歴、検索、分析
記事のタイトル: GitHub - sudomichael/agentgraphed: AI コーディング セッション用のローカルファースト分析ダッシュボード。 Claude Code および Codex CLI を使用して構築したものを確認してください。 · GitHub
説明: AI コーディング セッション用のローカルファースト分析ダッシュボード。 Claude Code および Codex CLI を使用して構築したものを確認してください。 - sudomichael/エージェントグラフ化
HN テキスト: 私はインディーハッキング タイプです。つまり、常に 5 つの異なることに取り組んでいます。クロードとコーデックスは素晴らしいです。しかし、彼らに関してはいくつかの問題がありました。 AgentGraphed はそれらをほぼ解決します。ローカルで動作し、ライブでのすべての会話をローカルの sqlite データベースにインデックス付けします。これを使用すると、優れた UX でレンダリングできることがたくさんあります。問題: セッションを再開したいのですが、claude --resume を見ると、タイトルがまったく役に立ちません。これらは、セッションを開始した会話の最初の文にすぎません。
解決策: LLM は各セッションに文脈に応じたタイトルを付けるため、より意味が分かりやすくなります。また、単純な「セッション再開」ボタンにより、 cd /path/to/session && claude --resume [sessionId] がコピーされ、準備完了です。問題: 休暇に行って、戻ってくると作業内容を忘れてしまいます。
解決策: AgentGraphed にはタイムラインがあり、いつ、何を作業していたかを正確に示します。問題: コーディング エージェントと何かについて話したことは覚えていますが、どのセッションだったか忘れてしまいました。
解決策: 履歴を検索可能。あらゆるセッションを、ずっと。問題: クロードをどれだけ使っているかを友達に自慢したいです。 ccusage は存在しますが、ターミナルネイティブの UX は共有するにはそれほど優れたものではありません
解決策: ソーシャルフレンドリーな共有ボタン。統計情報の画像を生成し、クリップボードにコピーします。 (すべてローカル)。問題: セッションを再開したくないが、再開したい

そこから重要なコンテキストをコピーします
解決策: 再利用できるようにコンテキストを生成するシンプルなボタン。 ------ 一つ注意しておきたいことがあります。このツールは実際に完全にローカルであり、完全に安全です。しかし！注意点 - 自動タイトル付けや「コンテキストの生成」機能などの AI 要約が必要な場合は、API キーを追加する必要があり、本質的にサードパーティと通信します。ただし、これらは完全にオプションです。私は OSS の世界には慣れていません。通常、私が行うことはすべて独自のものなので、お気軽に私をローストして、より良いものにするのを手伝ってください。

記事本文:
GitHub - sudomichael/agentgraphed: AI コーディング セッション用のローカルファースト分析ダッシュボード。 Claude Code および Codex CLI を使用して構築したものを確認してください。 · GitHub
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
スドマイケル
/
エージェントグラフ化された
公共
通知
通知を変更するにはサインインする必要があります

n設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
39 コミット 39 コミット .github .github bin bin docs/ スクリーンショット docs/ スクリーンショット scripts scripts src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md next-env.d.ts next-env.d.ts next.config.mjs next.config.mjs package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js tailwind.config.ts tailwind.config.ts tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードとコーデックスのローカル ファーストの履歴、検索、分析。
古いセッションを検索し、未完了の作業を再開し、過去の会話を検索し、マシン上のすべての AI コーディング プロジェクトの使用状況を確認します。
ログインがありません。雲はありません。コンピューターからは何も残りません。
🌐 サイト: https://agentgraphed.com · 📦 npm: https://www.npmjs.com/package/agentgraphed
AgentGraphed は、マシン上のすべてのクロード コードおよびコーデックス セッションに自動的にインデックスを付けます。これを使用すると、次のことが可能になります。
忘れていた古い作品を見つけてください。すべてのセッションで検索可能です。
放棄されたセッションをワンクリックで再開します。
タッチしたすべての Git リポジトリのプロジェクト履歴を参照します。
過去の会話から新しいチャットのコンテキストを生成します。
プロバイダー、モデル、プロジェクト全体で使用量とコストを追跡します。
すべてがローカルのままです。 AgentGraphed は、CLI ツールがすでに書き込んでいた JSONL ログ ( ~/.claude/projects/ 、 ~/.codex/sessions/ ) を読み取り、それらを実際のダッシュボードに変換します。
npxエージェントグラフ化
これでインストールは完了です。クローンもサインアップも設定ファイルもありません。ダッシュボードは http://localhost:3737 で開きます。いつでも再実行するか、実行したままにしてください。5 分ごとに再スキャンされます。
npm install -g エージェントグラフ化
エージェントグラフ化された
npx / ノードをお持ちではありませんか?ノード20をインストールする

+ https://nodejs.org から (LTS バージョンは問題ありません)。
npx Agentgraphed はパッケージを 1 回ダウンロードします (最大 15 MB、通常の接続では 30 ～ 60 秒かかります)。後続の実行は数秒で開始されます。npm はパッケージをキャッシュします。
1 つの npm warn exec 行と、推移的な依存関係からのいくつかの npm warn 非推奨の警告が表示されます。無害です。
ダッシュボード サーバーが起動します。
› http://localhost:3737 で AgentGraphed を開始する
› ローカル AI コーディング セッションをスキャン中…
142 クロード + 8 Codex セッションが見つかりました (1943 ミリ秒でインデックス付けされた 14821 メッセージ)
› 準備完了。 Ctrl+C を押して停止します。
ブラウザが自動的にダッシュボードを開きます。
サーバーを停止するには、 Ctrl+C を押します。クロード コードが元のログ ファイルをディスクからローテーションした後でも、インデックス付きデータは次回から ~/.agentgraphed/agentgraphed.sqlite に残ります。
古い作業を検索します。すべてのセッションを日ごとにグループ化します。プロジェクト、プロバイダー、またはモデル別に検索およびフィルターします。
未完了のセッションを再開します。会話全体を読んだり、1 行の再開コマンドをコピーしたり、新しいチャットに貼り付ける入門書を生成したりできます。
プロジェクト履歴を参照します。これまでに作業したすべての Git リポジトリがアクティビティごとにランク付けされます。任意のプロジェクトをクリックすると、そのセッションのみが表示されます。
LiteLLM の 2700 を超えるモデル カタログからの小売価格のコスト見積もりを使用して、日別、モデル別、プロジェクト別、カテゴリ別の使用量など、AI 時間がどこに費やされているかを確認します。
タイムライン — すべてのセッションが日ごとにグループ化され、「開始済み」「開始日まで延長」「継続中」「終了」のバッジが表示されるため、複数日にわたる作業が 1 つのバケットに隠れることはありません。
セッション — チャットバブル ビューで過去の会話を読み取ります。コンテンツ、プロジェクト、プロバイダー、またはモデルで検索します。
再開 — ワンクリックで cd <cwd> && claude --resume <id> をクリップボードにコピーします。
コンテキストの生成 (オプション、BYO LLM キー) — 再開時にコンテキストを失わないよう、新しいチャットに貼り付ける入門書を生成します。
プロジェクト — git リポジトリのルートから自動検出されます。

プロジェクトごとの使用状況、モデルの支出、セッション履歴。
ダッシュボード — 使用状況グラフ、KPI、上位プロジェクト、作業タイプの内訳、モデルごとのコスト。ダッシュボード全体をプロジェクトまたはモデル ファミリごとにフィルターします。
自動分類 (オプション、BYO LLM キー、デフォルトでオン) — セッションを機能 / デバッグ / 計画 / リファクタリング / スタイリング / DevOps / データ / 支払い / ドキュメント / コンテンツとして分類し、それぞれにきれいな過去形のタイトルを書き込みます。通常、数百回のセッションで数セントです。
ライブ クォータ プローブ (オプション) — サイドバー ウィジェットにカーソルを置くと、Anthropic 5 時間/7 日および OpenAI の 1 分あたりのライブ制限使用率が表示されます。単一トークン プローブ (それぞれ ~0.00006 ドル)。
共有 - ダッシュボード、プロジェクト、またはセッション ビューの統計カード PNG を生成し、クリップボードに直接コピーします。
バックグラウンド取り込み - 5 分ごとにローカル ログを再スキャンします。何もしなくても新しいセッションが表示されます。
コストの見積もり — LiteLLM の 2,700 以上のモデルの小売価格を自動更新します。方向性があるものとして扱います。
範囲ピッカー — すべてのチャートで 7 日 / 30 日 / 90 日 / 常時、データがロングテールである場合に自動提案される対数スケールを使用します。
すべてはマシン上の ~/.agentgraphed/agentgraphed.sqlite に存在します。
[コンテキストのコピー] または [セッションの分類] をクリックしない限り、何もアップロードされません。その場合でも、API キーを使用して LLM プロバイダーにのみアップロードされます。
API キーは SQLite ファイルに平文で保存されます。これは ~/.aws/credentials または .env と同じ脅威モデルです。ホームフォルダーを git にコミットしないでください。
クロード コード — ~/.claude/projects/*/*.jsonl を読み取ります
Codex CLI — ~/.codex/sessions/YYYY/MM/DD/*.jsonl を読み取ります
さらに多くのアダプターが登場します。特定のものが必要な場合は、問題を開きます。
オプション: LLM セッションのタイトルとカテゴリ
API キーを使用しない場合、セッションでは最初のプロンプトの最初の行がタイトルとして表示されます。キーを使用する場合:
Anthropic — Haiku 4.5 (デフォルト) はクリーンな p を提供します

「Stripe チェックアウトの小数点 2 桁のバグを修正」のような緊張感のあるタイトル
OpenAI — GPT-5 mini は同等に機能し、多くの場合安価です
[設定] → [LLM プロバイダー] をクリックし、キーを貼り付けて、 [未分類に分類] をクリックします。コストはセッションあたり 1 セントの数分で、通常、数百セッションで 0.01 ～ 0.03 ドルです。
ログ ディレクトリは、再起動せずに [設定] → [データ ソース] で編集することもできます。
git clone https://github.com/sudomichael/agentgraphed.git
cd エージェントグラフ化
npmインストール
npm 実行開発
http://localhost:3737 を開きます。ホットリロード、ソースは src/ の下にあります。
公開可能なビルドを生成するには:
npm run build # Next.js スタンドアロン バンドルをビルドします
npm Pack # Agentgraphed-X.Y.Z.tgz を作成します
ライセンス
AI コーディング セッション用のローカルファースト分析ダッシュボード。 Claude Code および Codex CLI を使用して構築したものを確認してください。
github.com/sudomichael/agentgraphed
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
3
フォーク
レポートリポジトリ
リリース
3
v0.3.1
最新の
2026 年 6 月 9 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local-first analytics dashboard for AI coding sessions. See what you built with Claude Code & Codex CLI. - sudomichael/agentgraphed

I'm an indiehacking type, which means I work on like 5 different things at any given time. Claude and codex are incredible. But, there were a few problems I had with them. AgentGraphed pretty much solves for those. It works locally - indexes every conversation you have, live, into a local sqlite db. With it, theres a lot we can render in a nice UX. Among other things.. Problem: I want to resume a session, but when I look at claude --resume, the titles are completely unhelpful. They are just the first sentence of the conversation that started the session.
Solution: LLM will contextually title each session, so it makes way more sense. Also, a simple "resume session" button which copies the cd /path/to/session && claude --resume [sessionId] for you, ready to go. Problem: I go on vacation, and forget what I was working on when I get back.
Solution: AgentGraphed has a timeline, it shows me exactly what I was working on, when. Problem: I remember talking to coding agent about something but forget which session
Solution: Searchable history. Every single session, ever. Problem: I want to brag to my friends about how much I use claude. ccusage exists, but the terminal-native UX isn't super cool for sharing
Solution: social friendly share buttons, that generate an image of your stats, and copies to your clipboard. (All local). Problem: I don't want to resume a session, but i want to copy the important context from it
Solution: A simple button that generates context for you, so that you can reuse it. ------ One thing I need to call out. The tool really is fully local, totally safe. BUT! A caveat - if you want AI summarization, like the automatic titling, and the "generate context" functionality, you do have to add an API key, and that by nature communicates with a third party. Those are totally optional, though. I'm new to the OSS world. Usually everything I do is proprietary etc, so feel free to roast me and help me make it better.

GitHub - sudomichael/agentgraphed: Local-first analytics dashboard for AI coding sessions. See what you built with Claude Code & Codex CLI. · GitHub
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
sudomichael
/
agentgraphed
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
39 Commits 39 Commits .github .github bin bin docs/ screenshots docs/ screenshots scripts scripts src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md next-env.d.ts next-env.d.ts next.config.mjs next.config.mjs package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js tailwind.config.ts tailwind.config.ts tsconfig.json tsconfig.json View all files Repository files navigation
Local-first history, search, and analytics for Claude Code and Codex.
Find old sessions, resume unfinished work, search past conversations, and see usage across every AI coding project on your machine.
No login. No cloud. Nothing leaves your computer.
🌐 Site: https://agentgraphed.com · 📦 npm: https://www.npmjs.com/package/agentgraphed
AgentGraphed automatically indexes every Claude Code and Codex session on your machine. With it you can:
Find old work you forgot about — every session, searchable.
Resume abandoned sessions with one click.
Browse project history across every git repo you've touched.
Generate context for a new chat from any past conversation.
Track usage and costs across providers, models, and projects.
Everything stays local. AgentGraphed reads the JSONL logs your CLI tools were already writing ( ~/.claude/projects/ , ~/.codex/sessions/ ) and turns them into a real dashboard.
npx agentgraphed
That's the whole install. No clone, no signup, no config file. The dashboard opens at http://localhost:3737 . Re-run any time, or leave it running — it re-scans every 5 minutes.
npm install -g agentgraphed
agentgraphed
Don't have npx / Node? Install Node 20+ from https://nodejs.org (the LTS version is fine).
npx agentgraphed downloads the package once (~15 MB, takes 30-60s on a typical connection). Subsequent runs start in seconds — npm caches the package.
You'll see one npm warn exec line and a couple of npm warn deprecated warnings from transitive dependencies. Harmless.
The dashboard server boots:
› Starting AgentGraphed on http://localhost:3737
› Scanning local AI coding sessions…
Found 142 Claude + 8 Codex sessions (14821 messages indexed in 1943ms)
› Ready. Press Ctrl+C to stop.
Your browser opens to the dashboard automatically.
To stop the server, hit Ctrl+C . Your indexed data stays at ~/.agentgraphed/agentgraphed.sqlite for next time — even after Claude Code rotates the original log files off disk.
Find old work — every session, grouped by day. Search and filter by project, provider, or model.
Resume unfinished sessions — read the full conversation, copy a one-line resume command, or generate a primer to paste into a fresh chat.
Browse project history — every git repo you've worked in, ranked by activity. Click any project to see only its sessions.
See where your AI time goes — usage by day, model, project, and category, with retail-priced cost estimates from LiteLLM's 2700+ model catalog.
Timeline — every session grouped by day, with STARTED · SPANS Nd / CONTINUED / CLOSED badges so multi-day work doesn't hide on a single bucket.
Sessions — read past conversations in a chat-bubble view; search by content, project, provider, or model.
Resume — one click copies cd <cwd> && claude --resume <id> to your clipboard.
Generate context (optional, BYO LLM key) — produces a primer to paste into a fresh chat so you don't lose context when resuming.
Projects — auto-detected from git repo roots; per-project usage, model spend, and session history.
Dashboard — usage chart, KPIs, top projects, work-type breakdown, per-model cost. Filter the whole dashboard by project or model family.
Auto-classification (optional, BYO LLM key, on by default) — categorizes sessions as Feature / Debugging / Planning / Refactor / Styling / DevOps / Data / Payments / Docs / Content and writes a clean past-tense title for each. Typically a few cents for hundreds of sessions.
Live quota probe (optional) — hover the sidebar widget to read your live Anthropic 5h/7d and OpenAI per-minute rate-limit utilization. Single-token probes (~$0.00006 each).
Share — generate a stat-card PNG of your dashboard, project, or session view and copy it straight to your clipboard.
Background ingest — re-scans local logs every 5 minutes; new sessions appear without you doing anything.
Cost estimates — LiteLLM's auto-updating retail pricing for 2700+ models. Treat as directional.
Range picker — 7d / 30d / 90d / all-time on every chart, with auto-suggested log scale when the data is long-tailed.
Everything lives in ~/.agentgraphed/agentgraphed.sqlite on your machine.
Nothing is uploaded unless you click "Copy context" or "Classify sessions," and even then only to your LLM provider with your API key.
API keys are stored in plaintext in the SQLite file — same threat model as ~/.aws/credentials or a .env . Don't commit your home folder to git.
Claude Code — reads ~/.claude/projects/*/*.jsonl
Codex CLI — reads ~/.codex/sessions/YYYY/MM/DD/*.jsonl
More adapters coming. If you want a specific one, open an issue .
Optional: LLM session titles & categories
Without an API key, sessions show the first line of the first prompt as the title. With a key:
Anthropic — Haiku 4.5 (default) gives clean past-tense titles like "Fixed Stripe checkout double-decimal bug"
OpenAI — GPT-5 mini works equally well, often cheaper
Click Settings → LLM provider , paste your key, then Classify uncategorized . Cost is fractions of a cent per session — typically $0.01–0.03 for a few hundred sessions.
Log directories can also be edited at Settings → Data sources without restarting.
git clone https://github.com/sudomichael/agentgraphed.git
cd agentgraphed
npm install
npm run dev
Open http://localhost:3737 . Hot reload, source under src/ .
To produce a publishable build:
npm run build # builds Next.js standalone bundle
npm pack # creates agentgraphed-X.Y.Z.tgz
License
Local-first analytics dashboard for AI coding sessions. See what you built with Claude Code & Codex CLI.
github.com/sudomichael/agentgraphed
Topics
There was an error while loading. Please reload this page .
3
forks
Report repository
Releases
3
v0.3.1
Latest
Jun 9, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
