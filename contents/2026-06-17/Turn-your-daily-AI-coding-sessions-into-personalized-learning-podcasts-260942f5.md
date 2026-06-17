---
source: "https://github.com/ron3899/aftercode"
hn_url: "https://news.ycombinator.com/item?id=48569110"
title: "Turn your daily AI coding sessions into personalized learning podcasts"
article_title: "GitHub - ron3899/aftercode: Turn your daily AI coding sessions into personalized learning podcasts — Rust CLI + API + web UI · GitHub"
author: "Rontwtw"
captured_at: "2026-06-17T13:24:18Z"
capture_tool: "hn-digest"
hn_id: 48569110
score: 1
comments: 2
posted_at: "2026-06-17T11:55:28Z"
tags:
  - hacker-news
  - translated
---

# Turn your daily AI coding sessions into personalized learning podcasts

- HN: [48569110](https://news.ycombinator.com/item?id=48569110)
- Source: [github.com](https://github.com/ron3899/aftercode)
- Score: 1
- Comments: 2
- Posted: 2026-06-17T11:55:28Z

## Translation

タイトル: 毎日の AI コーディング セッションをパーソナライズされた学習ポッドキャストに変える
記事のタイトル: GitHub - ron3899/aftercode: 毎日の AI コーディング セッションをパーソナライズされた学習ポッドキャストに変える — Rust CLI + API + Web UI · GitHub
説明: 毎日の AI コーディング セッションをパーソナライズされた学習ポッドキャストに変える — Rust CLI + API + Web UI - ron3899/aftercode

記事本文:
GitHub - ron3899/aftercode: 毎日の AI コーディング セッションをパーソナライズされた学習ポッドキャストに変える — Rust CLI + API + Web UI · GitHub
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
ロン3899
/
アフターコード
公共
通知
通知設定を変更するにはサインインする必要があります
さらに

l ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
31 コミット 31 コミット .aftercode .aftercode .github/ workflows .github/ workflows 資産 アセット クレート クレート ドキュメント ドキュメント マイグレーション マイグレーション Web Web Web .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md docker-compose.yml docker-compose.yml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
毎日の AI コーディング セッションをパーソナライズされた学習ポッドキャストに変えます。
Aftercode はコーディング エージェントのワークフロー (Claude Code、Codex、Cursor など) に接続し、その日に構築およびデバッグした内容の背後にあるより深い技術的なトピックを把握し、出荷したコードの背後にある概念を教える 2 人のスピーカーによる短いポッドキャスト エピソード (ヘブライ語または英語) を生成します。クリーンな Web プレイリストでそれらを聞いたり閲覧したりできます。
バイブコーディングを本当の学習に変えましょう。
オープンソース (MIT または Apache-2.0)、CLI ファースト、自己ホスト可能。
▶ 音声付きで見る（フル品質MP4）
アフターコードのエピソード
│ コーディング エージェント セッション (クロード コード / コーデックス / カーソル) を自動検出 →
│ プロンプト + エージェントが行ったこと/説明したこと
│ + 実際の git diff (Redis クライアント、RabbitMQ 構成、追加された新しい deps)
│ シークレットのスキャン、無視ルールの適用、編集、キャップ サイズ
▼
バックエンドパイプライン
正規化 → トピックの抽出 → 知識ギャップの検出 → ランク付け →
2 話者のスクリプトを作成 (ヘブライ語または英語) → セグメントごとの TTS
(イレブンラボまたはOpenAI) → concat + 無音ギャップ → MP3 → ストア → エピソード
▼
ウェブUI
プレイリスト · トピックでフィルタリング · エピソードを再生 · トランスクリプト + クイズ
リポジトリのレイアウト
クレート/アフターコードコア共有

ed タイプ (セッションコンテキスト、エピソード、オーディオ)
crates/aftercode-server Axum API バックエンド + 生成パイプライン + Web UI を提供
crates/aftercode-cli `aftercode` CLI
Web/ React + Vite Web UI (プレイリスト、フィルター、プレーヤー)
移行/ SQLite スキーマ (起動時に自動適用)
ドキュメント/ PRD、セルフホスティング、アーキテクチャ、設計仕様
インストール
必要なもの: Docker (バックエンド + Web UI を実行) および
Rust (アフターコード CLI の場合 - ローカルを読み取るためにネイティブに実行する必要があります)
コーディング エージェント セッション ファイル)。それだけです — データベースもノードもありません。
次の 4 つのコマンドをコピーして貼り付けます。
#1. コードを取得する
git clone https://github.com/ron3899/aftercode && cd アフターコード
# 2. バックエンド + Web UI を起動 → http://localhost:8080
# (最初の実行ではイメージが 1 回、約 5 分でビルドされます。その後は即座に実行されます)
ドッカー構成 -d
# 3. CLI をインストールする
カーゴインストール --path crates/aftercode-cli
# 4. ログイン (ブラウザを開く → [承認] をクリック)
アフターコードログイン
次に、http://localhost:8080 を開いてエピソードを参照し、任意の環境でアフターコード エピソードを実行します。
プロジェクト。ステップ 2 は API キーなし (モック モード) で動作するため、すぐに試すことができます。
実際のエピソードの場合は、キーを使用してリポジトリに .env を作成し、 docker compose up -d を実行します。
もう一度:
LLM_PROVIDER=オープンナイ
OPENAI_API_KEY=sk-...
TTS_PROVIDER=openai # OpenAI tts-1 経由の音声;または、TTS_PROVIDER=elevenlabs + そのキーを設定します
ポート 8080 が使用されましたか? PORT=9000 docker compose up -d を実行し、どこでもそのポートを使用します。
コーディングエージェント内で使用します
CLI は、使用したエージェントから現在のセッションを自動検出するため、最も単純な
ワークフローは、一連の作業を完了してからエピソードを作成することです。
ターミナル (任意のエージェント): aftercode エピソード -- language en (または he )。それは、
そのリポジトリの最新の Cursor 、 Claude Code 、または Codex セッションが自動的に取得されます。
エージェントに依頼してください。エージェントは端末を持っているので、次のように伝えるだけです。

「aftercode エピソード -- language en を実行して、今行ったことをポッドキャストに変換します。」
Claude Code / Codex (CLI エージェント): 同じフォルダーでセッションの直後に実行します —
検出によりトランスクリプトが見つかります。別のフォルダーから実行した場合は、記録を提出してください
直接: aftercodeepisode --transcript - (トランスクリプトをパイプします) - 検出に失敗することはありません。
Cursor / Windsurf (エディター エージェント): プロジェクトでエディターのターミナルを開き、実行します。
アフターコードエピソード ;そのワークスペースの Cursor のライブ セッションを読み取ります。
最初にアフターコード プレビューを実行して、どのエージェント + セッションが検出されたか、そして正確に何が起こるかを確認します。
送られる。
cd web && npm install && npm run build && cd .. # UI をビルドします (/ で提供されます)
cp .env.example .env # キーを追加するか、モックをそのままにして試してみます
カーゴ run -p aftercode-server # 最初の実行でログイン トークン + URL が出力されます
カーゴインストール --path crates/aftercode-cli
aftercode ログイン && cd your-project && aftercode エピソード -- language en
ウェブUI
バックエンドは / で React アプリを提供します ( web/dist から構築):
プレイリスト — 永続的な最下位プレーヤーが含まれるエピソードを再生します。
現在のフィルタリングされたリスト (前/次 = キュー)。
トピック (チップ) + 言語 + 検索によるフィルター - 瞬時に。
エピソードの詳細 — プレーヤー、トピック、重要なポイント、クイズ、完全なホスト/専門家のトランスクリプト。
レスポンシブ (モバイル + デスクトップ)。サインインでは CLI 承認フローが再利用されます (トークンの貼り付けは必要ありません)。
ホットリロードを使用した開発: cd web && npm run dev (Vite on :5173 、バックエンドと通信、設定
web/.env の VITE_API_BASE)。 npm run build を使用して実稼働用に再ビルドします。
aftercode ログイン (引数なし) は <backend>/cli/authorize を開きます。 「承認」とトークンをクリックします
CLI の場合はローカルホスト ループバック経由、Web UI の場合はローカルホスト ループバック経由で自動的にキャプチャされます。
ブラウザのリダイレクト (コピーアンドペーストは不可)。 aftercode ログイン <トークン> とシードユーザーは残ります。
スクリプト/CI。アフターコードステータスが検証される

バックエンドに対するトークン。
ローカル承認では ID チェックは実行されません。信頼できるローカルホスト/自己ホストのみが対象です。
パブリック/マルチユーザー展開には、実際の OAuth プロバイダー (まだ構築されていません) が必要です。このビルドは
実質的にシングルユーザー: サインインしているユーザーであればすべてのエピソードが表示されます。
エピソードは、エージェント セッション + git diff から構築されます。 Aftercode が自動検出します。
クロード コード / コーデックス / ディスク上のカーソル ストレージからのセッション (カーソルはライブ SQLite WAL を読み取ります)、
ただし、検出が失敗する可能性があります (サブフォルダーからの実行など)。 2 つの安全策:
成績証明書を直接提出してください: your-transcript |アフターコードエピソード --転写-
(Claude Code .jsonl 、単純な {"role","text"} JSONL、またはプレーン テキストを受け入れます) - スキップします
ディスクの自動検出。
薄いエピソードのガードレール: セッション会話がキャプチャされない場合、エピソードは拒否されます (つまり、
単独の package-lock.json に関するエピソードは決して得られません)。 --allow-thin でオーバーライドします。
最初に aftercode プレビューを実行して、何が収集され、どのエージェントが検出されたかを確認します。
env 経由のすべての Secrets/config ( .env.example および docs/SELF_HOSTING.md を参照):
DATABASE_URL — SQLite、デフォルト sqlite://aftercode.db?mode=rwc (自動作成 + 移行)
LLM_PROVIDER — 人間（デフォルト、claude-opus-4-8）/openai/モック; ANTHROPIC_API_KEY / OPENAI_API_KEY
TTS_PROVIDER — elevenlabs (デフォルト) / openai / モック
イレブンラボ: ELEVENLABS_API_KEY 、 ELEVENLABS_HOST_VOICE_ID 、 ELEVENLABS_EXPERT_VOICE_ID
OpenAI: OPENAI_TTS_MODEL (デフォルト gpt-4o-mini-tts )、OPENAI_TTS_VOICE_HOST、OPENAI_TTS_VOICE_EXPERT
BLOB_STORE — localfs (デフォルト) / s3 (+ S3_BUCKET );音声は /static/episodes/<id>.mp3 で提供されます
APP_PUBLIC_URL 、 BIND_ADDR 、 WEB_DIST (構築された Web アプリへのパス)
LLM プロバイダーはトレイトの背後でプラグイン可能です。トピック + スクリプト ステージでは、構造化された JSON 出力を使用します。
脚本はエピソードの言語 (ヘブライ語 kee) で書かれています。

英語の PS 技術用語、その方法
イスラエルの開発者が語る）。
Aftercode は、エージェント セッションのトランスクリプト (プロンプト + エージェント メッセージ + コマンド/編集) を送信します。
実行されました) および git diff ハンク — 完全なソース ファイル、 .env 、シークレット、または無視されることはありません
パス。正規表現シークレットスキャナーは、アップロード前にキー/トークンのように見えるものをすべて編集します。
ペイロード サイズには上限があり、アフターコード プレビューには送信される内容が正確に表示されます。
フェーズ 1 (完了): Rust CLI + API バックエンド — セッションキャプチャ、トピック/スクリプト生成、
TTS、オーディオ アセンブリ、エピソード ストレージ。ブラウザのログイン。
フェーズ 2 (完了): React + Vite Web UI — プレイリスト、トピック フィルター、プレーヤー、エピソードの詳細。
フェーズ 3 (計画中): 本物の OAuth/マルチユーザー、毎日の自動生成、ダウンロード/共有、
プライベート RSS、エディター拡張機能。 docs/superpowers/spec を参照してください。
オプションで、MIT ( LICENSE-MIT ) または Apache-2.0 ( LICENSE-APACHE ) のいずれかに基づいてライセンスが付与されます。
毎日の AI コーディング セッションをパーソナライズされた学習ポッドキャストに変える - Rust CLI + API + Web UI
不明、MIT ライセンスが見つかりました
ライセンスが見つかりました
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
アフターコード v0.1.0
最新の
2026 年 6 月 16 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Turn your daily AI coding sessions into personalized learning podcasts — Rust CLI + API + web UI - ron3899/aftercode

GitHub - ron3899/aftercode: Turn your daily AI coding sessions into personalized learning podcasts — Rust CLI + API + web UI · GitHub
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
ron3899
/
aftercode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
31 Commits 31 Commits .aftercode .aftercode .github/ workflows .github/ workflows assets assets crates crates docs docs migrations migrations web web .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md docker-compose.yml docker-compose.yml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Turn your daily AI coding sessions into personalized learning podcasts.
Aftercode connects to your coding-agent workflow (Claude Code, Codex, Cursor, …), figures out the deeper technical topics behind what you built and debugged that day, and generates a short two-speaker podcast episode — in Hebrew or English — that teaches you the concepts behind the code you shipped. Listen and browse them in a clean web playlist.
Turn vibe coding into real learning.
Open source (MIT OR Apache-2.0), CLI-first, self-hostable.
▶ Watch with sound (full-quality MP4)
aftercode episode
│ auto-detect your coding-agent session (Claude Code / Codex / Cursor) →
│ your prompts + what the agent did/explained
│ + the real git diff (the Redis client, RabbitMQ config, new deps it added)
│ scan for secrets, apply ignore rules, redact, cap size
▼
backend pipeline
normalize → extract topics → detect knowledge gaps → rank →
write two-speaker script (Hebrew or English) → TTS per segment
(ElevenLabs or OpenAI) → concat + silence gaps → MP3 → store → episode
▼
web UI
playlist · filter by topic · play through your episodes · transcript + quiz
Repository layout
crates/aftercode-core shared types (session context, episode, audio)
crates/aftercode-server Axum API backend + generation pipeline + serves the web UI
crates/aftercode-cli the `aftercode` CLI
web/ React + Vite web UI (playlist, filters, player)
migrations/ SQLite schema (auto-applied on startup)
docs/ PRD, self-hosting, architecture, design specs
Install
You need: Docker (runs the backend + web UI) and
Rust (for the aftercode CLI — it must run natively to read your local
coding-agent session files). That's it — no database, no Node.
Copy-paste these four commands:
# 1. Get the code
git clone https://github.com/ron3899/aftercode && cd aftercode
# 2. Start the backend + web UI → http://localhost:8080
# (first run builds the image once, ~5 min; after that it's instant)
docker compose up -d
# 3. Install the CLI
cargo install --path crates/aftercode-cli
# 4. Log in (opens your browser → click Approve)
aftercode login
Now open http://localhost:8080 to browse episodes, and run aftercode episode in any
project. Step 2 works with no API keys (mock mode) so you can try it immediately.
For real episodes , create a .env in the repo with your keys, then docker compose up -d
again:
LLM_PROVIDER=openai
OPENAI_API_KEY=sk-...
TTS_PROVIDER=openai # voice via OpenAI tts-1; or set TTS_PROVIDER=elevenlabs + its keys
Port 8080 taken? Run PORT=9000 docker compose up -d and use that port everywhere.
Use it inside your coding agent
The CLI auto-detects your current session from whatever agent you used, so the simplest
workflow is: finish a chunk of work, then make the episode.
In the terminal (any agent): aftercode episode --language en (or he ). It reads the
latest Cursor , Claude Code , or Codex session for that repo automatically.
Ask the agent to do it for you. Since the agent has a terminal, just tell it:
"Run aftercode episode --language en to turn what we just did into a podcast."
Claude Code / Codex (CLI agents): run it right after your session in the same folder —
detection finds the transcript. If you ran from a different folder, hand the transcript in
directly: aftercode episode --transcript - (pipe a transcript) — never fails detection.
Cursor / Windsurf (editor agents): open the editor's terminal in your project and run
aftercode episode ; it reads Cursor's live session for that workspace.
Run aftercode preview first to see which agent + session was detected and exactly what will
be sent.
cd web && npm install && npm run build && cd .. # build the UI (served at /)
cp .env.example .env # add keys, or leave mock to try
cargo run -p aftercode-server # first run prints a login token + URL
cargo install --path crates/aftercode-cli
aftercode login && cd your-project && aftercode episode --language en
Web UI
The backend serves a React app at / (built from web/dist ):
Playlist — your episodes with a persistent bottom player that plays through the
current filtered list (prev/next = your queue).
Filter by topic (chips) + language + search — instant.
Episode detail — player, topics, key takeaways, quiz, full host/expert transcript.
Responsive (mobile + desktop). Sign in reuses the CLI approve flow — no token paste.
Dev with hot reload: cd web && npm run dev (Vite on :5173 , talks to the backend; set
VITE_API_BASE in web/.env ). Rebuild for production with npm run build .
aftercode login (no arg) opens <backend>/cli/authorize ; click Approve and the token
is captured automatically — for the CLI via a localhost loopback, for the web UI via a
browser redirect (no copy-paste). aftercode login <token> and seed-user remain for
scripts/CI; aftercode status validates the token against the backend.
Local-approval performs no identity check — it's for trusted localhost/self-host only.
A public/multi-user deployment needs a real OAuth provider (not yet built). This build is
effectively single-user: any signed-in user sees all episodes.
Episodes are built from your agent session + the git diff . Aftercode auto-detects the
session from Claude Code / Codex / Cursor on-disk storage (Cursor reads its live SQLite WAL),
but detection can miss (e.g. a run from a subfolder). Two safeguards:
Hand the transcript in directly: your-transcript | aftercode episode --transcript -
(accepts a Claude Code .jsonl , simple {"role","text"} JSONL, or plain text) — skips
disk auto-detection.
Thin-episode guardrail: if no session conversation is captured, episode refuses (so
you never get an episode about a lone package-lock.json ). Override with --allow-thin .
Run aftercode preview first to see what was collected and which agent was detected.
All secrets/config via env (see .env.example and docs/SELF_HOSTING.md ):
DATABASE_URL — SQLite, default sqlite://aftercode.db?mode=rwc (auto-created + migrated)
LLM_PROVIDER — anthropic (default, claude-opus-4-8 ) / openai / mock ; ANTHROPIC_API_KEY / OPENAI_API_KEY
TTS_PROVIDER — elevenlabs (default) / openai / mock
ElevenLabs: ELEVENLABS_API_KEY , ELEVENLABS_HOST_VOICE_ID , ELEVENLABS_EXPERT_VOICE_ID
OpenAI: OPENAI_TTS_MODEL (default gpt-4o-mini-tts ), OPENAI_TTS_VOICE_HOST , OPENAI_TTS_VOICE_EXPERT
BLOB_STORE — localfs (default) / s3 (+ S3_BUCKET ); audio served at /static/episodes/<id>.mp3
APP_PUBLIC_URL , BIND_ADDR , WEB_DIST (path to the built web app)
LLM providers are pluggable behind a trait; topic + script stages use structured JSON output.
The script is written in the episode's language (Hebrew keeps tech terms in English, the way
Israeli devs speak).
Aftercode sends your agent-session transcript (prompts + agent messages + the commands/edits
it ran) and the git diff hunks — never full source files, .env , secrets, or ignored
paths. A regex secret scanner redacts anything that looks like a key/token before upload,
payload size is capped, and aftercode preview shows exactly what would be sent.
Phase 1 (done): Rust CLI + API backend — session capture, topic/script generation,
TTS, audio assembly, episode storage. Browser login.
Phase 2 (done): React + Vite web UI — playlist, topic filter, player, episode detail.
Phase 3 (planned): real OAuth/multi-user, auto-daily generation, download/share,
private RSS, editor extensions. See docs/superpowers/specs .
Licensed under either of MIT ( LICENSE-MIT ) or Apache-2.0 ( LICENSE-APACHE ) at your option.
Turn your daily AI coding sessions into personalized learning podcasts — Rust CLI + API + web UI
Unknown, MIT licenses found
Licenses found
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Aftercode v0.1.0
Latest
Jun 16, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
