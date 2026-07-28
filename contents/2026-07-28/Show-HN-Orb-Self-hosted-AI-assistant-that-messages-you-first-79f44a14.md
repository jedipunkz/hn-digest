---
source: "https://github.com/getorb/Orb-Backend"
hn_url: "https://news.ycombinator.com/item?id=49089227"
title: "Show HN: Orb – Self-hosted AI assistant that messages you first"
article_title: "GitHub - getorb/Orb-Backend · GitHub"
author: "ninjahawk1"
captured_at: "2026-07-28T21:00:12Z"
capture_tool: "hn-digest"
hn_id: 49089227
score: 1
comments: 0
posted_at: "2026-07-28T20:08:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Orb – Self-hosted AI assistant that messages you first

- HN: [49089227](https://news.ycombinator.com/item?id=49089227)
- Source: [github.com](https://github.com/getorb/Orb-Backend)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T20:08:18Z

## Translation

タイトル: Show HN: Orb – 最初にメッセージを送信する自己ホスト型 AI アシスタント
記事タイトル: GitHub - getorb/Orb-Backend · GitHub
説明: GitHub でアカウントを作成して、getorb/Orb-Backend の開発に貢献します。

記事本文:
GitHub - getorb/Orb-バックエンド · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ゲットーブ
/
Orb バックエンド
公共
通知
通知設定を変更するにはサインインする必要があります
アディティ

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット データ データ otools otools scripts scripts .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md API.md CLAUDE.md CLAUDE.md LICENSE LICENSE PRESERVATION.md PRESERVATION.md README.md README.md SELF_HOSTING.md SELF_HOSTING.md SETUP.md SETUP.md エージェント.py エージェント.py apns.py apns.py 脳.py 脳.py コネクタ.py コネクタ.py 意図ルーター.py 意図ルータ.py mcp_http.py mcp_http.py メモリ_ストア.py メモリ_ストア.py マインド.py マインド.py orb_mcp.py orb_mcp.py orb_memory.py orb_memory.py personas.py personas.py professional_engine.py professional_engine.py required.txt required.txt screen_bridge.py screen_bridge.py server_win.py server_win.py stt.py stt.pyvisor.py supervisor.py task_store.py task_store.py tools_registry.py tools_registry.py tools.py tools.py すべてのファイルを表示リポジトリ ファイルのナビゲーション
Orb の背後にあるバックエンド —
あなたが所有するハードウェア上で最初に話す AI アシスタント。
アプリを入手 ·
セットアップのウォークスルー ·
APIリファレンス・
あらゆる設定
iOS アプリは顔です。このサーバーが背後にあるすべてのものです。アプリは以下で動作します
デバイス上のボックス — このサーバーは、ハンド、メモリ、
そして独自のマシン。 Orb は質問されるのを待ちません: バックエンドが見守っています
あなたのプロジェクト、何が起こっているかの実行モデルを維持し、事前に計画し、承認されているかどうか
自動的に動作し、本当に価値がある場合にのみ携帯電話に届きます。
注意。
能力
それが何を意味するか
音声サーバー
WebSocket 音声ループ: 音声入力 (ローカル ウィスパー)、実際の会話出力 (ニューラル TTS)。電話ではウェイクワードはありません。ただ話すだけです。
エージェントスパイン
モデル自体がいつ話すか、いつ行動するかを決定する 1 つの薄いループ。最大 60 個も組み込まれています。

ls: 天気、メール、ファイル、株式、画面、セッション、ミッション、リマインダー、通知。
マルチブレインルーター
Claude (ログインした独自の Claude Code CLI を使用)、Grok、Codex、または単一の一貫した音声の背後にあるローカル モデル。会話の途中で交換します。人格は決して変わりません。
心
会話の間に実行されるプロセス。あなたの人生とプロジェクトの状況モデルを維持し、何が変化したかに気づき、予測を立て、実際の仕事を提案し、何を承認するかを監督します。
セッション
長時間実行されるジョブ (ビルド、リサーチ、トレーニングの実行) を、携帯電話から監視、メッセージ送信、停止できる追跡されたバックグラウンド エージェントに委任します。
プッシュは正しく完了しました
重複排除、ハード ペーシング、フルメッセージ ペイロード、および黙って飲み込むことのできない配信確認を伴う APNs 通知。
走る
現在は Windows ファースト (macOS バックエンドを予定)。 Python 3.11以降。
git clone https://github.com/getorb/Orb - バックエンド
cd Orb - バックエンド
Python - m venv venv
venv\Scripts\pip install - rrequirements.txt
venv\Scripts\python scripts\setup.py # ガイド付き: トークン、ID、ブート チェック
venv\Scripts\pythonvisor.py # クラッシュ リカバリを使用してサーバーを実行します
次に、携帯電話が到達できるアドレスを指定し（ tailscaleserve --bg 8340 ）、ペアリングします。
アプリ: メニュー (☰) → PC 。すべての注意点を含む完全なチュートリアルは次のとおりです。
セルフホスティング ガイド 。
委任することをお勧めしますか?クロードコードをインストールし、
リポジトリのルートで claude を実行し、「Orb バックエンドをセットアップしてください」と言います。
リポジトリの CLAUDE.md が全体の流れを教えてくれます。
モデル アクセス — デフォルトの会話脳は、あなた自身の脳を介してクロードです。
ログインしたクロードコード CLI
( クロード -p )。 API キーは出荷、要求、保存されません。 Grok/Codex CLI および
OpenAI 互換の無料利用枠をルーター チェーンに挿入可能
( ORB_FREE_BACKENDS )。
プッシュ — Apple Developer acco からの独自の APNs キー

unt (.p8 ファイル —
決してコミットしないでください。 .gitignore はすでに拒否しています)。閉じた状態へのバックグラウンドプッシュ
アプリは、自分のチーム + バンドル ID での iOS ビルドでのみ動作します。大衆と一緒に
App Store アプリでは、代わりにアプリ内 WebSocket 通知を受け取ります。詳細は
SELF_HOSTING.md 。
メール — 自分の Gmail アプリのパスワード。組み込みのメール ツールは読み取り専用です。
すべての変数は .env.example に文書化されており、
SELF_HOSTING.md 。
ファイル
役割
サーバー_win.py
サーバー: WS 音声ループ、HTTP API、ブレイン ディスパッチ、通知チョーク ポイント (重複排除 / ペーシング / 受信)
エージェント.py
シン エージェント ループ — モデルが会話かツールかを決定し、ハーネスが実行します
ブレイン.py
Brain Router: Claude CLI、Grok/Codex、無料バックエンド、ローカル フロア、クールダウン
tool_registry.py + otools/
すべてのツール (それぞれ 1 つのファイル) が @tool(...) で登録されています
mcp_http.py + orb_mcp.py
ツールは MCP として表示されます。クロード脳の常駐ループバック /mcp。 Codex/Grok 用の標準出力
プロアクティブ_エンジン.py
ゼロ AI 収集ティック、1 日 2 回の合成、スケジューラー (リマインダーは再起動後も存続)
マインドパイ
会話間の代理店: 状況モデル、予算付き通夜、承認ゲート付き提案
stt.py / ペルソナ.py
耳 (より速いささやき声) と音声 (ニューラル TTS + ペルソナの定義)
スーパーバイザー.py
バックエンドを存続させます。コード変更後の自己再起動を有効にする
iOSアプリ
App Store で公開中 (無料):
https://apps.apple.com/us/app/orb-ai/id6776376035
アプリ自体はクローズドソースです。このバックエンドはペアの開いた半分、つまりアプリです。
WebSocket + API.md に記載されている小さな HTTP API を介して通信します。
独自のクライアントを構築することを妨げるものは何もありません。
地元第一主義。データはマシン上に収集され、そこに残ります。
鍵はご自身でご用意ください。バックエンドは資格情報を送信しません。
構造上の安全性。自律型サーフェスは破壊的なツールに到達できません。
ツール

それらの表面には、単に落胆しているだけではありません。行動するものなら何でも
世界ではあなたの明示的な承認が必要であり、すべての自律的なステップが記録されます
どこで読めるか。
静粛性が特徴です。プロアクティブ層は「言うに値しないこと」を問題として扱います。
正しい、第一級の結果。
ファーストクラスのデフォルトとしての完全オフラインのローカルモデル脳層 (Ollama / Qwen)。
アプリ内からQRペアリング。
MIT — 「ライセンス」を参照してください。ライセンスはこのバックエンドのみを対象としています。 Orb iOS アプリ
Orb の名前/ブランドはその一部ではありません。
Readme ライセンス アクティビティ カスタム プロパティ スター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to getorb/Orb-Backend development by creating an account on GitHub.

GitHub - getorb/Orb-Backend · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
getorb
/
Orb-Backend
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits data data otools otools scripts scripts .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md API.md API.md CLAUDE.md CLAUDE.md LICENSE LICENSE PRESERVATION.md PRESERVATION.md README.md README.md SELF_HOSTING.md SELF_HOSTING.md SETUP.md SETUP.md agent.py agent.py apns.py apns.py brain.py brain.py connectors.py connectors.py intent_router.py intent_router.py mcp_http.py mcp_http.py memory_store.py memory_store.py mind.py mind.py orb_mcp.py orb_mcp.py orb_memory.py orb_memory.py personas.py personas.py proactive_engine.py proactive_engine.py requirements.txt requirements.txt screen_bridge.py screen_bridge.py server_win.py server_win.py stt.py stt.py supervisor.py supervisor.py tasks_store.py tasks_store.py tool_registry.py tool_registry.py tools.py tools.py View all files Repository files navigation
The backend behind Orb —
the AI assistant that talks first, on hardware you own.
Get the app ·
Setup walkthrough ·
API reference ·
Every setting
The iOS app is the face; this server is everything behind it. The app works out of
the box on-device — this server is the optional half that gives it hands, memory,
and a machine of its own. Orb doesn't wait to be asked: the backend watches over
your projects, keeps a running model of what's going on, plans ahead, does approved
work on its own, and reaches your phone only when it's genuinely worth your
attention.
Capability
What that means
Voice server
A WebSocket voice loop: speech in (local Whisper), a real conversation out (neural TTS). No wake word on the phone — you just talk.
Agent spine
One thin loop where the model itself decides when to speak and when to act, with ~60 built-in tools: weather, mail, files, stocks, screen, sessions, missions, reminders, notifications.
Multi-brain router
Claude (through your own logged-in Claude Code CLI), Grok, Codex, or local models behind a single consistent voice. Swap mid-conversation; the persona never changes.
The mind
A process that runs between conversations: maintains a situation model of your life and projects, notices what changed, makes predictions, proposes real work, supervises what you approve.
Sessions
Delegate long-running jobs (builds, research, training runs) to tracked background agents you can watch, message, and stop from your phone.
Push done right
APNs notifications with deduplication, hard pacing, full-message payloads, and delivery receipts that can never be silently swallowed.
Run
Windows-first today (macOS backend planned). Python 3.11+.
git clone https: // github.com / getorb / Orb - Backend
cd Orb - Backend
python - m venv venv
venv\Scripts\pip install - r requirements.txt
venv\Scripts\python scripts\setup.py # guided: token, identity, boot check
venv\Scripts\python supervisor.py # runs the server with crash recovery
Then give it an address your phone can reach ( tailscale serve --bg 8340 ) and pair
the app: menu (☰) → Your PC . The full walkthrough with every gotcha spelled out:
Self-Hosting guide .
Prefer to delegate? Install Claude Code ,
run claude in the repo root, and say "Set up the Orb backend for me" — the
repo's CLAUDE.md teaches it the whole flow.
Model access — the default conversational brain is Claude through your own
logged-in Claude Code CLI
( claude -p ). No API key is shipped, asked for, or stored. Grok/Codex CLIs and
any OpenAI-compatible free tier can slot into the router chain
( ORB_FREE_BACKENDS ).
Push — your own APNs key from your Apple Developer account (a .p8 file —
never commit it; the .gitignore already refuses). Background push to a closed
app only works for an iOS build under your own team + bundle id; with the public
App Store app you get in-app WebSocket notifications instead. Details in
SELF_HOSTING.md .
Mail — your own Gmail app password; the built-in mail tools are read-only.
Every variable is documented in .env.example and
SELF_HOSTING.md .
File
Role
server_win.py
The server: WS voice loop, HTTP API, brain dispatch, the notification choke point (dedup / pacing / receipts)
agent.py
The thin agent loop — the model decides talk-vs-tool, the harness executes
brain.py
Brain router: Claude CLI, Grok/Codex, free backends, local floor, cooldowns
tool_registry.py + otools/
Every tool, one file each, registered with @tool(...)
mcp_http.py + orb_mcp.py
The tool surface as MCP: resident loopback /mcp for the Claude brain; stdio for Codex/Grok
proactive_engine.py
Zero-AI collection ticks, twice-daily synthesis, the scheduler (reminders survive restarts)
mind.py
Between-conversation agency: situation model, budgeted wakes, approval-gated proposals
stt.py / personas.py
Ears (faster-whisper) and voice (neural TTS + persona definitions)
supervisor.py
Keeps the backend alive; enables self-restart after code changes
The iOS app
Live on the App Store (free):
https://apps.apple.com/us/app/orb-ai/id6776376035
The app itself is closed-source. This backend is the open half of the pair: the app
talks to it over a WebSocket + a small HTTP API documented in API.md —
nothing stops you from building your own client.
Local-first. Your data is collected on your machine and stays there.
Bring your own keys. The backend never ships credentials.
Structural safety. Autonomous surfaces can't reach destructive tools — those
tools are absent from those surfaces, not merely discouraged. Anything that acts
on the world requires your explicit approval, and every autonomous step is logged
where you can read it.
Silence is a feature. The proactive layer treats "nothing worth saying" as a
correct, first-class outcome.
Fully offline local-model brain tier (Ollama / Qwen) as a first-class default.
QR pairing from inside the app.
MIT — see LICENSE . The license covers this backend only; the Orb iOS app
and the Orb name/branding are not part of it.
Readme License Activity Custom properties Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
