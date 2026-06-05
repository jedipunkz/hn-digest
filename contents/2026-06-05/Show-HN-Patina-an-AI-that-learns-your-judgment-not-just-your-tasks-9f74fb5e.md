---
source: "https://github.com/Sanctum-Origo-Systems/patina"
hn_url: "https://news.ycombinator.com/item?id=48405624"
title: "Show HN: Patina, an AI that learns your judgment, not just your tasks"
article_title: "GitHub - Sanctum-Origo-Systems/patina: A persistent cognitive backend with belief graph and earned autonomy · GitHub"
author: "andywidjaja"
captured_at: "2026-06-05T00:58:40Z"
capture_tool: "hn-digest"
hn_id: 48405624
score: 2
comments: 1
posted_at: "2026-06-04T22:39:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Patina, an AI that learns your judgment, not just your tasks

- HN: [48405624](https://news.ycombinator.com/item?id=48405624)
- Source: [github.com](https://github.com/Sanctum-Origo-Systems/patina)
- Score: 2
- Comments: 1
- Posted: 2026-06-04T22:39:28Z

## Translation

タイトル: Show HN: Patina、タスクだけでなく判断力も学習する AI
記事のタイトル: GitHub - Sanctum-Origo-Systems/patina: 信念グラフと獲得した自律性を備えた永続的なコグニティブ バックエンド · GitHub
説明: 信念グラフと獲得した自律性を備えた永続的な認知バックエンド - Sanctum-Origo-Systems/patina

記事本文:
GitHub - Sanctum-Origo-Systems/patina: 信念グラフと獲得した自律性を備えた永続的な認知バックエンド · GitHub
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
サンクタム・オリゴ・システムズ
/
緑青
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
49 コミット 49 コミット .github .github config config eval eval scripts scripts src/ patina src/ patina testing testing .cursorrules .cursorrules .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md patina-demo.gif patina-demo.gif pyproject.toml pyproject.toml test_conversation.py test_conversation.py uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたを学習する AI — 長く使えば使うほど、AI に伝える必要がなくなります。
あなたの認知の上限はあなたの知性によって決まるわけではありません。それはあなたの認知負荷によって設定されます。 Patina はあなたの認知の永続的な拡張です。それはあなたのコンテキストを保持し、あなたの世界についてのあなたの信念を追跡し、あなたの判断を反映し、あらゆる相互作用によって改善されます。
1 日目: Slack をエクスポートします。 5 分以内に、見逃したもの、忘れたもの、見逃したものをすべて確認します。
30 日目: ユーザーの優先順位を認識し、音声の下書きを作成し、ノイズを自動的に無視します。
90 日目: ユーザーの行動を予測し、会話全体の矛盾を見つけ、静かに動作します。
# インストール
UVツールインストール緑青
# 初期化
緑青の初期化
# Slack エクスポートからの取り込み (即値)
緑青取り込み --from-export ~ /Downloads/slack-export.zip
# 注意が必要なものを確認する
緑青のキャッチアップ
# 優先象限ごとにランク付けされたすべてを表示
緑青の優先順位
何をするのか
# 優先順位+判断
緑青のキャッチアップ # 統合ビュー: アクションが必要 / 新規 / 待機中
緑青の優先順位 # 象限ごとにグループ化 (Q1 ～ Q4)
patina dismiss < id > # ノイズを無視し、モデルをトレーニングします
緑青の目標は、「 Ship v2 」 --キーワード「 release,deploy 」を追加します。
# スタイル + 製図
Patina style build # 送信されたメッセージから通信プロファイルを構築する
patina style show < name > # pe のパターンを表示

アールソン
緑青ドラフト --to < 名前 > --context " タイムラインのフォローアップ "
# 信念グラフ
patina extract # LLM を介して観察から信念を抽出する
緑青の信念 --type person # エンティティ (請求数あり)
緑青が古くなり、信頼しきい値を下回る朽ち果てた信念の数
緑青の矛盾 # 矛盾する主張
緑青関係 -- トップ 20 # 信頼レベル + アクティビティ マップ
# 段階的な自律性
緑青の自律性ステータス # 現在のレベル、精度、アンチパターン
patina accept < id > # 提案されたアクションを承認します
緑青拒否 < id > # 拒否 (進行を凍結し、アンチパターンを保存)
Patina autonomy set-level < N > # 手動オーバーライド (0-6)
# ライブアダプター
緑青接続スラック --token " xoxb-... "
Patina ingest # 設定されたアダプターから取得する
# ハートビート (バックグラウンドタスク)
緑青ハートビート 1 回 # 取り込み + 減衰 + エスカレーション チェック
緑青ハートビート開始 -- 間隔 30
# インタラクティブな会話 (Claude Agent SDK)
緑青チャット
# ゲートウェイ統合用の HTTP サーバー
緑青サーブ --ポート 8321
# Telegram ゲートウェイ (電話からエージェントに話しかけます)
緑青のゲートウェイ
仕組み
┌───────────────────────┐
│ Tier 3: Frontier LLM (Claude、GPT-4o) │ 総合、ドラフト、矛盾
━━━━━━━━━━━━━━━━━━┤
│ Tier 2: ローカル LLM (Qwen 3.x、Ollama) │ エンティティの抽出、分類
━━━━━━━━━━━━━━━━━━┤
│ 階層 1: 決定論的 (LLM なし) │ スコアリング、減衰、グラフ クエリ
━━━━━━

───────────────┘
↓ 全階層フィード ↓
┌───────────────────────┐
│ 信念グラフ (SQLite) │
│ エンティティ → 関係 → クレーム │
│ 信頼の崩壊・来歴・矛盾 │
━━━━━━━━━━━━━━━━━┘
システムは Tier 1 だけで完全に機能します (LLM コールはゼロ)。各層は機能を追加しますが、負荷はかかりません。ローカルファースト: すべてのデータは ~/.patina/store.db に残ります。
減衰を伴う永続的な信念モデル — メッセージアーカイブではなく、生きた世界モデル
判断力は自分の決定から学べます – 普遍的なルールではなく、自分の優先事項
精度によって獲得される段階的な自律性 — 設定されていない、証明されていない
ローカルファースト、モデルに依存しない — オフラインで実行され、ベンダーロックインなし
決定論的コア — インテリジェンスは数学とグラフであり、LLM 呼び出しではありません
輸出初日の価値 — ウォームアップ期間なし
Patina は、Claude Code、Cline、または任意の MCP ホストから会話で使用するための MCP サーバーとして実行されます。コンテキストの適切なハンドオフのための session_checkpoint と、ステートレス セッション間での会話の継続のための Recent_messages を含む 21 のツール。
{
"mcpサーバー": {
"緑青コア" : {
"コマンド" : " uv " ,
"args" : [ " run " 、 " patina-mcp " ]、
"cwd" : " /path/to/patina "
}
}
}
構成
すべての設定は ~/.patina/config.yaml にあります。資格情報がマシンから離れることはありません。
所有者 :
ユーザー ID : ["U0ABC123"]
名前：「君の名は。」
アダプター:
チャット:
- プロバイダー：slack
トークン: " xoxb-... "
心拍数:
有効 : true
間隔分: 30
開発
gitクローンhtt

ps://github.com/Sanctum-Origo-Systems/patina.git
CD緑青
UV同期
# デモデータを生成してチャットを開始する
uv run python scripts/generate_demo_export.py --output demo-export.zip
UV 実行緑青初期化
uv run patina ingest --from-export デモ-export.zip
uv run patina extract --model Sonnet # タイムアウトになった場合は再実行 — すでに処理されたものをスキップします
uv run python scripts/seed_decions.py # デモのみ — 信頼スコアリングのためのユーザーの動作をシミュレートします
uv run patina style build # 送信されたメッセージからコミュニケーション スタイル プロファイルを構築する
uv ラン緑青チャット
# 質問してみてください:
# 「注意が必要なのは何ですか?」
# 「私が一番信頼できるのは誰ですか?」
# 「ジョーダンについて私たちは何を知っていますか?」
# 「アトラスのタイムラインについてアレクシスにメッセージを下書きする」
# 「私の信念に矛盾はありますか?」
# テストと評価を実行する
UV で pytest を実行
uv 実行 pytest eval/deterministic/
#糸くず
UV 実行ラフチェック && UV 実行ラフ形式
ライセンス
信念グラフと獲得した自律性を備えた永続的な認知バックエンド
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
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A persistent cognitive backend with belief graph and earned autonomy - Sanctum-Origo-Systems/patina

GitHub - Sanctum-Origo-Systems/patina: A persistent cognitive backend with belief graph and earned autonomy · GitHub
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
Sanctum-Origo-Systems
/
patina
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
49 Commits 49 Commits .github .github config config eval eval scripts scripts src/ patina src/ patina tests tests .cursorrules .cursorrules .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md patina-demo.gif patina-demo.gif pyproject.toml pyproject.toml test_conversation.py test_conversation.py uv.lock uv.lock View all files Repository files navigation
An AI that learns you — the longer you use it, the less you need to tell it.
Your cognitive ceiling isn't set by your intelligence. It's set by your cognitive load. Patina is a persistent extension of your cognition: it holds your context, tracks your beliefs about your world, mirrors your judgment, and improves with every interaction.
Day 1: Export your Slack. In 5 minutes, see everything you've missed, forgotten, or let slip.
Day 30: It knows your priorities, drafts in your voice, and dismisses noise automatically.
Day 90: It predicts what you'd do, catches contradictions across conversations, and operates silently.
# Install
uv tool install patina
# Initialize
patina init
# Ingest from a Slack export (immediate value)
patina ingest --from-export ~ /Downloads/slack-export.zip
# See what needs your attention
patina catch-up
# See everything ranked by priority quadrant
patina priorities
What It Does
# Priority + judgment
patina catch-up # unified view: needs action / new / waiting
patina priorities # grouped by quadrant (Q1-Q4)
patina dismiss < id > # dismiss noise, trains the model
patina objectives add " Ship v2 " --keywords " release,deploy "
# Style + drafting
patina style build # build communication profiles from sent messages
patina style show < name > # view patterns for a person
patina draft --to < name > --context " follow up on the timeline "
# Belief graph
patina extract # extract beliefs from observations via LLM
patina beliefs --type person # entities with claim counts
patina stale # decayed beliefs below confidence threshold
patina contradictions # conflicting claims
patina relationships --top 20 # trust level + activity map
# Graduated autonomy
patina autonomy status # current level, accuracy, anti-patterns
patina approve < id > # approve a proposed action
patina reject < id > # reject (freezes advancement, stores anti-pattern)
patina autonomy set-level < N > # manual override (0-6)
# Live adapters
patina connect slack --token " xoxb-... "
patina ingest # fetch from configured adapters
# Heartbeat (background tasks)
patina heartbeat once # ingest + decay + escalation check
patina heartbeat start --interval 30
# Interactive conversation (Claude Agent SDK)
patina chat
# HTTP server for gateway integration
patina serve --port 8321
# Telegram gateway (talk to your agent from your phone)
patina gateway
How It Works
┌─────────────────────────────────────────────────┐
│ Tier 3: Frontier LLM (Claude, GPT-4o) │ Synthesis, drafts, contradictions
├─────────────────────────────────────────────────┤
│ Tier 2: Local LLM (Qwen 3.x, Ollama) │ Entity extraction, classification
├─────────────────────────────────────────────────┤
│ Tier 1: Deterministic (no LLM) │ Scoring, decay, graph queries
└─────────────────────────────────────────────────┘
↓ all tiers feed ↓
┌─────────────────────────────────────────────────┐
│ Belief Graph (SQLite) │
│ Entities → Relationships → Claims │
│ Confidence decay · Provenance · Contradictions │
└─────────────────────────────────────────────────┘
The system is fully functional at Tier 1 alone (zero LLM calls). Each tier adds capability but never load-bears. Local-first: all data stays in ~/.patina/store.db .
Persistent belief model with decay — not a message archive, a living world model
Judgment learned from your decisions — not universal rules, YOUR priorities
Graduated autonomy earned by accuracy — not configured, proven
Local-first, model-agnostic — runs offline, no vendor lock-in
Deterministic core — the intelligence is math and graphs, not LLM calls
Day-one value from export — no warm-up period
Patina runs as an MCP server for conversational use from Claude Code, Cline, or any MCP host. 21 tools including session_checkpoint for graceful context handoff and recent_messages for conversational continuity across stateless sessions.
{
"mcpServers" : {
"patina-core" : {
"command" : " uv " ,
"args" : [ " run " , " patina-mcp " ],
"cwd" : " /path/to/patina "
}
}
}
Configuration
All config lives in ~/.patina/config.yaml . Credentials never leave your machine.
owner :
user_ids : ["U0ABC123"]
name : " Your Name "
adapters :
chat :
- provider : slack
token : " xoxb-... "
heartbeat :
enabled : true
interval_minutes : 30
Development
git clone https://github.com/Sanctum-Origo-Systems/patina.git
cd patina
uv sync
# Generate demo data and start chatting
uv run python scripts/generate_demo_export.py --output demo-export.zip
uv run patina init
uv run patina ingest --from-export demo-export.zip
uv run patina extract --model sonnet # re-run if it times out — skips already-processed
uv run python scripts/seed_decisions.py # demo only — simulates user behavior for trust scoring
uv run patina style build # build communication style profiles from sent messages
uv run patina chat
# Try asking:
# "What needs my attention?"
# "Who do I trust most?"
# "What do we know about Jordan?"
# "Draft a message to Alexis about the Atlas timeline"
# "Any contradictions in my beliefs?"
# Run tests + evals
uv run pytest
uv run pytest eval/deterministic/
# Lint
uv run ruff check && uv run ruff format
License
A persistent cognitive backend with belief graph and earned autonomy
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
