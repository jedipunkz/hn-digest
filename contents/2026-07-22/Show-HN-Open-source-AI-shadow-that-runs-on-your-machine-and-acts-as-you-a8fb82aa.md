---
source: "https://github.com/NanoNets/ami"
hn_url: "https://news.ycombinator.com/item?id=49002822"
title: "Show HN: Open-source AI shadow that runs on your machine and acts as you"
article_title: "GitHub - NanoNets/ami: Ami - your AI clone. A local agent harness that connects to your tools, learns your world, and does your busy work. · GitHub"
author: "ole_gooner"
captured_at: "2026-07-22T07:34:47Z"
capture_tool: "hn-digest"
hn_id: 49002822
score: 4
comments: 0
posted_at: "2026-07-22T07:01:37Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source AI shadow that runs on your machine and acts as you

- HN: [49002822](https://news.ycombinator.com/item?id=49002822)
- Source: [github.com](https://github.com/NanoNets/ami)
- Score: 4
- Comments: 0
- Posted: 2026-07-22T07:01:37Z

## Translation

タイトル: Show HN: マシン上で実行され、ユーザーとして機能するオープンソース AI シャドウ
記事のタイトル: GitHub - NanoNets/ami: Ami - あなたの AI クローン。ツールに接続し、世界を学習し、忙しい仕事をこなすローカル エージェント ハーネス。 · GitHub
説明: アミ - あなたの AI クローン。ツールに接続し、世界を学習し、忙しい仕事をこなすローカル エージェント ハーネス。 - ナノネッツ/ami

記事本文:
GitHub - NanoNets/ami: Ami - AI クローン。ツールに接続し、世界を学習し、忙しい仕事をこなすローカル エージェント ハーネス。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。リロードしてください

このページ。
ナノネット
/
アミ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット apps apps パッケージ package .gitignore .gitignore ライセンス ライセンス README.md README.md ami ami package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
ami-製品-ビデオ.mp4
Ami は、別個の同僚ではなく、あなたのクローンとして機能するローカル エージェント ハーネスです。個人トークンを使用してアプリ、データ、リポジトリ、ツールに接続し、ライブ ToDo リストを維持し、ユーザーのタスクの実行方法を学習し、ユーザーに代わって「忙しい仕事」タスクを実行します。ユーザーのコンテキスト グラフ メモリを構築し、エンティティ、関係、フィードバック、意思決定、書き方を更新して、使えば使うほど自律性が高まります。
アミはあなたのマシン上で動作します。データは ~/.ami/ (SQLite + マークダウン メモリ) の下に存在します。組織とは何も共有されません。 Ami は、あなたが持っているアクセス権を持っています。つまり、あなたが到達できる範囲にのみ到達できます。そして、クリックしなければ何も発送されず、送信されません。
唯一の前提条件はノード ≥ 20 です。
git clone https://github.com/NanoNets/ami.git && cd ami
./アミ
./ami は最初の実行時にすべてをブートストラップします。
依存関係をインストールします (corepack 経由の pnpm)
コンソールを構築してサーバーを起動します
http://localhost:4141 を開いて、オンボーディングを実行します。個人トークンを使用してモデルを構成し、ツールを接続します。どちらもこのマシンにのみ保存されます。
Ami は、Anthropic Messages API をプロトコルとして話します。プロバイダーの切り替えは、[設定] → [モデル] で 1 回変更するだけです。
API キー — プロバイダーのキー (デフォルトでは Anthropic のキー)。
ベース URL — Anthropic の場合は空白のままにするか、

Anthropic 互換のエンドポイントではありません。
トリアージとタスク実行のモデル — メイン モデル (デフォルトは claude-opus-4-8 )。フリーテキスト: エンドポイントが提供する任意のモデル名。
メモリ エージェントのモデル — 安価なユーティリティ モデル (デフォルトは claude-sonnet-4-6 )。これはあらゆる実質的な信号で動作するため、ローカルセットアップでは、これは小型モデルを指すノブになります。
Claude Agent SDK はベース URL を尊重し、その内部の小規模モデル呼び出しはユーティリティ モデルに再マップされます。構造化された出力は、非 Anthropic エンドポイントではスキーマ要求の JSON に自動的にフォールバックするため、互換プロバイダーは調整なしで機能します。
ホストされたオープン モデル - Kim、GLM、DeepSeek、および MiniMax はすべて、Anthropic 互換のベース URL を公開します (例: Kimi の場合は https://api.moonshot.ai/anthropic)。ベース URL、キー、モデル名を設定します。
完全にローカル — LiteLLM プロキシを Ollama / vLLM / llama.cpp の前に置きます。 Anthropic プロトコルを公開します。
オラマ プル クウェン3:32b
pip インストール「litellm[プロキシ]」
litellm --model ollama/qwen3:32b # :4000 の Anthropic 互換エンドポイント
次に、設定 → モデル: ベース URL http://localhost:4000 、両方のノブのモデル ollama/qwen3:32b (またはメモリ エージェントの小さいモデル)、およびプロキシに認証がない場合は任意のプレースホルダー キー。コネクタがローカルファーストでポーリングすると、マシンからは何も出なくなります。
./ami update # 最新のものをプルし、再インストールし、再構築します
./ami build # コンソールを強制的に再構築します
オプションの追加機能。一致する機能が必要な場合はインストールします。
gh CLI — コーディングタスクで PR を開くために必要です ( brew install gh )
ffmpeg + Whisper-cpp — オンデバイスでの文字起こし機能を備えたローカル会議レコーダー ( brew install ffmpeg Whisper-cpp )
ページ
何をするのか
ホーム
To Do と会議が表示されたホーム画面
やるべきこと
接続されているすべてのツールの To Do リストの共通リスト
チャット
To-Do ru と同じツール サーフェスを使用した副操縦士のチャット

ns。質問したり、仕事を中断したり、ToDo を作成したり、記憶を更新したりできます
エージェント
スケジュールに従って実行するエージェントを作成するページ
歴史
過去のToDoのアーカイブ
記憶
現在のメモリをコンテキスト グラフとして視覚化
設定
コネクタ、モデル、ノブ、使用データ
グローバル検索と通知はヘッダー ナビゲーション バーに表示されます。
取り込み — ポーラーは、個人トークン/API キーを使用して、各コネクタを通じて新しいメンション、DM、電子メール、招待、通知を取得します。
トリアージ — クロード コールは、各シグナルを期限付きのタスク/FYI/無視として分類します。エンティティと関係は知識ベースに流れ込みます。あなたのフィードバック (修正、却下、上書き) は、ナレッジ ベース内の意思決定の痕跡となり、将来のトリアージの方向性を決定します。
行動 — 各 ToDo 項目には 5 つのアクションがあります。
計画 — 読み取り専用で調査し、編集可能な計画を提案します。
開始 — 分離された Claude Agent SDK セッションでタスクを実行します。
解決 — 自分で解決したことを示し、ナレッジベースを更新します
却下/スヌーズ — 却下すると、同様のタスクの今後のトリアージが妨げられ、スヌーズされたタスクは明日に戻ります。
レビュー — 成果物 (PR リンク、ドキュメント、イベント、会議リンク) が、返信の下書きとともにタスク ページに表示されます。アミ自身は何も送信しません。編集/承認すると、アカウントから元の Slack スレッド/電子メール スレッドに投稿されます。編集内容はメモリに反映され、Ami の演技と声が洗練されます。
反復 — 完了したタスクに関するフィードバックにより、完全なコンテキストで同じエージェント セッションが再開されます。
あらゆる To-Do の実行と副操縦士のチャットにより、知識ベースが増加します。
Memory は ~/.ami/knowledge/ にある Obsidian スタイルのマークダウン ナレッジ ベースです。wiki リンクを含むドシエで、アクティビティの各バッチ後にノート エージェントによってキュレートされ、git でバージョン管理されているため、すべてのノートに履歴と復元が含まれます。それは次のような力を与えます。
会議の準備 — 出席者の準備

知識ベースに対抗する。各会議の前に、概要（誰が来るか、何が重要か、未解決の項目）が表示されます。
会議レコーダー — マイクとシステム音声をローカルで録音し、デバイス上で Whisper.cpp を使用して文字起こしし、トランスクリプトをメモリにフィードします。音声はマシンから出力されません。
スタイル — Ami は、実際のメッセージとその下書きの編集から書き方を学びます。答えはあなたの声に出ます。
意思決定トレース — 過去の例外、オーバーライド、フィードバックは、同様の状況が再発したときに取得されるため、Ami はユーザーがすでに決定した方法で決定します。
現在、Ami のコネクタの数は限られています。ただし、あらゆるツール コネクタを構築するのは熟練しています。
コネクタリストの下部にある「コネクタの追加」をクリックします。
名前、ホームページの URL、コネクタを使用する意図を指定します。
Ami は (クロード コードと API キーを使用して) コネクタを構築し、タスクの実行や副操縦士のチャットで使用できるようにします。
Ami が知っていることはすべて ~/.ami/ の下にあります。
ami.db (SQLite — シグナル、todos、実行、設定、トークン)
Knowledge/ (マークダウンナレッジベース、git バージョン)
worktrees/ (分離コーディングチェックアウト)、
bg-tasks/ および events/ (バックグラウンド エージェントの状態)。
ディレクトリを削除すると、Ami はすべてを忘れます。アウトバウンド呼び出しは、Anthropic API と、トークンで接続したツールに対するものだけです。
Ami は内部ツールとして構築されています。まだ開発段階にあり、すぐに安定版リリースをプッシュする予定です。乞うご期待。
アミ - あなたの AI クローン。ツールに接続し、世界を学習し、忙しい仕事をこなすローカル エージェント ハーネス。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フー

ターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Ami - your AI clone. A local agent harness that connects to your tools, learns your world, and does your busy work. - NanoNets/ami

GitHub - NanoNets/ami: Ami - your AI clone. A local agent harness that connects to your tools, learns your world, and does your busy work. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
NanoNets
/
ami
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits apps apps packages packages .gitignore .gitignore LICENSE LICENSE README.md README.md ami ami package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
ami-product-video.mp4
Ami is a local agent harness that acts as your clone, not a separate co-worker. It connects to apps, data, repositories, tools with your personal tokens, maintains a live to-do list, learns how you do tasks, and executes the "busy work" tasks on your behalf. It contructs a context graph memory of you and updates entities, relationships, feedback, decisions, writing style, to get more autonomous the more you use it.
Ami runs on your machine. Data lives under ~/.ami/ (SQLite + markdown memory). Nothing is shared with your org. Ami has exactly the access you have, i.e., it can only reach what you can reach. And nothing ships or gets sent without your click.
The only prerequisite is Node ≥ 20.
git clone https://github.com/NanoNets/ami.git && cd ami
./ami
./ami bootstraps everything on first run.
It installs dependencies (pnpm via corepack)
builds the console and starts the server
Open http://localhost:4141 and walk through onboarding. Configure your model and connect tools using personal tokens. Both are stored only on this machine.
Ami speaks the Anthropic Messages API as a protocol. Switching provider is a single change in Settings → Models:
API key — your provider's key (Anthropic's by default).
Base URL — leave blank for Anthropic, or point it at any Anthropic-compatible endpoint.
Model for triage & task runs — the main model (default claude-opus-4-8 ). Free-text: any model name your endpoint serves.
Model for memory agents — the cheap/utility model (default claude-sonnet-4-6 ). It runs on every substantive signal, so on a local setup this is the knob to point at a small model.
The Claude Agent SDK honors the base URL, and its internal small-model calls are remapped to your utility model. Structured outputs automatically fall back to schema-prompted JSON on non-Anthropic endpoints, so compat providers work without tweaks.
Hosted open models - Kimi, GLM, DeepSeek, and MiniMax all publish Anthropic-compatible base URLs (e.g. https://api.moonshot.ai/anthropic for Kimi). Set the base URL, their key, and their model names.
Fully local — put a LiteLLM proxy in front of Ollama / vLLM / llama.cpp; it exposes the Anthropic protocol:
ollama pull qwen3:32b
pip install ' litellm[proxy] '
litellm --model ollama/qwen3:32b # Anthropic-compatible endpoint on :4000
Then in Settings → Models: base URL http://localhost:4000 , model ollama/qwen3:32b for both knobs (or a smaller model for memory agents), and any placeholder key if the proxy has no auth. With connectors polling local-first, nothing leaves your machine at all.
./ami update # pull the latest, reinstall, rebuild
./ami build # force-rebuild the console
Optional extras, install if you want the matching feature:
gh CLI — coding tasks need it to open PRs ( brew install gh )
ffmpeg + whisper-cpp — a local meeting recorder with on-device transcription ( brew install ffmpeg whisper-cpp )
Page
What it does
Home
A homescreen with to-dos and meetings
To-do
A universal list of to-dos from every connected tool
Chat
A copilot chat with the same tool surface as to-do runs. You can ask questions, fire off work, create to-dos, update memory
Agents
A page to create agents that run on a schedule
History
An archive of past to-dos
Memory
A visualization of the current memory as a context graph
Settings
Connectors, models, knobs, usage data
Global search and notifications live in the header navbar.
Ingest — a poller pulls new mentions, DMs, emails, invites, notifications through each connector using your personal tokens / API keys.
Triage — a Claude call classifies each signal as task / FYI / ignore with a due date. Entities and relationships flow into the knowledge base. Your feedback (corrections, dismissals, overrides) become decision traces in the knowledge base that steer future triage.
Act — each to-do item has five actions:
Plan — explores read-only and proposes an editable plan.
Start — runs the task in an isolated Claude Agent SDK session.
Resolve — indicates you did it yourself and updates knowledge base
Dismiss / Snooze — dismissals discourages future triage of similar tasks, snoozed tasks return tomorrow.
Review — deliverables (PR link, doc, event, meeting link) land on the task page along with a drafted reply. Ami never sends anything itself. You edit/approve, then it posts to the originating Slack thread / email thread from your account. Your edits are diffed into memory and refine Ami's execution and voice.
Iterate — feedback on a finished task resumes the same agent session with full context.
Every to-do run and every copilot chat grows the knowledge base.
Memory is an Obsidian-style markdown knowledge base at ~/.ami/knowledge/ — dossiers with wiki-links, curated by note agents after each batch of activity, versioned with git so every note has history and restore. It powers:
Meeting prep — attendees resolved against the knowledge base; a brief (who's coming, what matters, open items) lands before each meeting.
Meeting recorder — record mic + system audio locally, transcribe on-device with whisper.cpp, feed the transcript into memory. No audio leaves the machine.
Style — Ami learns how you write from your real messages and your edits to its drafts; replies come out in your voice.
Decision traces — past exceptions, overrides, and feedback are retrieved when similar situations recur, so Ami decides the way you already decided.
Ami currently has a limited number of connectors. But it is skilled to build any tool connector for you.
Click on "Add a connector" at the bottom of the connector list.
Specify the name, homepage url, and your intent on using the connector
Ami will build the connector (using your claude code and your API key) and make it available to task runs and copilot chats.
Everything Ami knows lives under ~/.ami/ .
ami.db (SQLite — signals, todos, runs, settings, tokens)
knowledge/ (the markdown knowledge base, git-versioned)
worktrees/ (isolated coding checkouts),
bg-tasks/ and events/ (background agent state).
Delete the directory and Ami forgets everything. The only outbound calls are to the Anthropic API and to the tools you connected with your tokens.
Ami is built as an internal tool. It's still in development stage, and we'll push a stable release soon. Stay tuned.
Ami - your AI clone. A local agent harness that connects to your tools, learns your world, and does your busy work.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
