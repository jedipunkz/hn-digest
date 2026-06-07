---
source: "https://github.com/aksika/abtars"
hn_url: "https://news.ycombinator.com/item?id=48436973"
title: "AbTARS – Self-hosted AI agent with persistent memory and 5-layer self-healing"
article_title: "GitHub - aksika/abtars: Self-hosted AI agent bridge - connects LLMs to Telegram/Discord/IRC with memory, tools, scheduled tasks, agent-swarm and self-healing supervision. · GitHub"
author: "aksika"
captured_at: "2026-06-07T18:33:29Z"
capture_tool: "hn-digest"
hn_id: 48436973
score: 1
comments: 0
posted_at: "2026-06-07T17:34:42Z"
tags:
  - hacker-news
  - translated
---

# AbTARS – Self-hosted AI agent with persistent memory and 5-layer self-healing

- HN: [48436973](https://news.ycombinator.com/item?id=48436973)
- Source: [github.com](https://github.com/aksika/abtars)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T17:34:42Z

## Translation

タイトル: AbTARS – 永続メモリと 5 層の自己修復機能を備えた自己ホスト型 AI エージェント
記事のタイトル: GitHub - aksika/abtars: 自己ホスト型 AI エージェント ブリッジ - メモリ、ツール、スケジュールされたタスク、エージェント群、自己修復監視を使用して LLM を Telegram/Discord/IRC に接続します。 · GitHub
説明: 自己ホスト型 AI エージェント ブリッジ - メモリ、ツール、スケジュールされたタスク、エージェント群、自己修復監視を使用して LLM を Telegram/Discord/IRC に接続します。 - アクシカ/アブターズ

記事本文:
GitHub - aksika/abtars: 自己ホスト型 AI エージェント ブリッジ - メモリ、ツール、スケジュールされたタスク、エージェント群、自己修復監視を使用して LLM を Telegram/Discord/IRC に接続します。 · GitHub
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
アクシカ
/
アブター
公共
通知
サインインする必要があります

o 通知設定を変更する
追加のナビゲーション オプション
コード
dev ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2,487 コミット 2,487 コミット .github/ workflows .github/ workflows .kiro/ステアリング .kiro/ ステアリング エージェント エージェント config-examples config-examples config config core core docker/ browser docker/ browser docs/ wiki docs/ wiki package/ homebrew package/ homebrew plugins/ openclaw-kiro-professor plugins/ openclaw-kiro-professor スクリプト スクリプト src src テスト テスト .env.example .env.example .env.skills.example .env.skills.example .gitattributes .gitattributes .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知README.md README.md esbuild.config.js esbuild.config.js install-manifest.json install-manifest.json install.sh install.sh models.default.json models.default.json package-lock.json package-lock.json package.json package.json task.default.json task.default.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
永続メモリ、自己修復機能、ピアツーピア通信を備えた自律型 AI エージェント。あなたのハードウェア、あなたのルール。
abTARS はマシン上で 24 時間年中無休で実行されます。Telegram/Discord/IRC 経由で話しかけ、セッション全体ですべてを記憶し、スケジュールされたタスクを実行し、介入なしで障害から回復します。既存の CLI サブスクリプションを通じてフロンティア モデルを限界費用ゼロで使用します。
🧠 自らを管理する記憶 — 多層想起 (5 つの検索ステージ + 再ランキング)、夜間の睡眠維持、感情追跡、矛盾検出、記憶ダーウィニズム
🛡️ 自己ホスト型、多層防御 — 機密メモリ (4 層、保存時に暗号化)、ロールベースのアクセス、インジェクション スキャン、秘密保管庫
🔄 数か月間無人で実行 — 5 層のスーパーバイズ

イオン、リーキーバケットモデルフォールバック、自己修復エージェント、スタンバイ対応リカバリ
🤝 エージェント間 — Ed25519 署名による P2P 通信、mDNS ウェイクアップ、IRC 調整チャネル
💰 アイドルコストゼロ — 保存時の LLM 呼び出しなし、CLI サブスクリプション寄生、予算スリープ層
→ なぜ abTARS と OpenClaw および Hermes を対比させるのか
あなた（Telegram / Discord / IRC / APIクライアント）
│
▼
アブターズ（ブリッジ）
§── abmind (メモリ - インプロセス、多層リコール、暗号化)
§── スキル（コア＋スリープ中に自作＋ダウンロード可能）
§── ツール (browse、bash、MCP、peer_ask)
§── タスク (cron スケジューラ + リトライ + DoD チェック)
§── Agent Swarm (非同期バックグラウンドセッション)
│
§── kiro-cli → Claude、DeepSeek、MiniMax、Qwen (無料枠)
§── gemini-cli → Gemini 2.5
§── 直接 API → ollam、OpenRouter、任意の OpenAI 互換
└── ピア → /v1/chat/completions 経由の他の abTARS インスタンス
クイックスタート
npm install -g abtars@alpha abmind@alpha
インストールするのをやめます
アブターズインストール
アブターズアップデート
船上のアブター
sudo $( what abtars ) デーモンのインストール
インストール後、Telegram ボット トークンと少なくとも 1 つのモデル プロバイダーを使用して ~/.abtars/config/.env を構成します。完全なガイド: docs/wiki/install.md
スタンドアロン パッケージ — abTARS、Kiro CLI、Gemini CLI、Claude Code、Codex、Hermes、または任意の MCP クライアントで動作します。
マルチレイヤーリコール: 5 つの検索ステージ (ポーター FTS5 → トライグラム → バイナリ署名 → ベクトル埋め込み → エンティティ グラフ) + 7 つの後処理レイヤー (クロスステージ ペナルティ、コンテキスト ブースト、感情ブースト、スペーシング ブースト、品質ブースト、MMR 再ランキング、干渉検出)
膠着語サポート (ハンガリー語、フィンランド語、トルコ語) — QWERTZ フォールバック、部分文字列ウィンドウ
記憶ごとのスコアリングを備えた 25 の感情タイプ、トピックごとの感情アーク
12 ステップの夜間睡眠: 抽出、統合、

枝刈り、矛盾の検出、翻訳の修正
NATO にヒントを得た分類: 機密性 × 信頼 × 完全性 × 信頼性
履歴からの自動編集を備えた AES-256-GCM 暗号化シークレット ボールト
記憶ダーウィニズム — 使われなかった記憶は薄れ、思い出した記憶は強化される
L1 ハートビート — スタンバイ検出、ブリッジロック、タスクディスパッチ
L2 インプロセス ウォッチドッグ — スタックしたイベント ループを検出します
L3 外部ウォッチドッグ — 無効な PID、古いハートビートをキャッチします。サーキットブレーカーが再起動の嵐を防ぐ
L4 OS スーパーバイザ - launchd (macOS) / systemd (Linux) はウォッチドッグ自体を再起動します
L5 予防的な毎日の再起動 — メモリ リークを排除します
モデルの健全性 — モデルごとのリーキーバケット、段階的なペナルティ、任意に長いフォールバックチェーン
自己修復エージェント - 失敗したタスクを診断し、修復を試み、3 回失敗すると中断します
メイン エージェントは、独立したバックグラウンド セッション (独自のコンテキスト、独自のツール ループ) を生成します。結果は親の次のプロンプトに自動挿入されます。最大3人まで同時接続可能。 /wait は、実行中のセッションに命令を挿入します。
複数の abTARS インスタンスは、OpenAI 互換の /v1/chat/completions エンドポイントを介して通信します。ピアごとのベアラー認証、Ed25519 署名、ファイアウォールで保護されたピアの mDNS ウェイクアップ、ホップ リミット ループ防止。
プラットフォームレベルのアクセス: 登録されたチャット ID/ユーザー ID のみがエージェントに到達できます。
ロールベース: マスター/ユーザー/ゲスト - コマンド、ツール、メモリはすべてゲートされています
Secrets Vault: AES-256-GCM、暗号化派生キー、取り込み時の自動暗号化
すべての受信メッセージに対するインジェクション スキャナー
すべてのログとエクスポートにおける資格情報の編集
輸送
プロバイダー
ACP (推奨)
キロクリ、ジェミニクリ
ダイレクトAPI
ollam、OpenRouter、任意の OpenAI 互換エンドポイント
フック (スタンドアロン)
CLI エージェントのライフサイクル フックを忌避する
要件
Telegram ボット トークン (Discord/IRC はオプション)
オプション: メモリ埋め込み用の ollam + nomic-embed-text

へこみ。
完全なドキュメント: aksika.github.io/abtars
1794 件のテスト (アブター 1016 + アブマインド 778)
5つのエージェントタイプ（教授、ドリーミー、ブラウジー、コーディング、cron）
3 つのプラットフォーム アダプター + OpenAI 互換 API
12ステップの夜間記憶維持
git クローン https://github.com/aksika/abtars.git
cd abtars && npm install && npm run build
npmテスト
コミュニティ
GitHub: aksika/abtars · aksika/abmind
自己ホスト型 AI エージェント ブリッジ - メモリ、ツール、スケジュールされたタスク、エージェント群、自己修復監視を使用して LLM を Telegram/Discord/IRC に接続します。
Apache-2.0ライセンス
貢献する
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

Self-hosted AI agent bridge - connects LLMs to Telegram/Discord/IRC with memory, tools, scheduled tasks, agent-swarm and self-healing supervision. - aksika/abtars

GitHub - aksika/abtars: Self-hosted AI agent bridge - connects LLMs to Telegram/Discord/IRC with memory, tools, scheduled tasks, agent-swarm and self-healing supervision. · GitHub
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
aksika
/
abtars
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
dev Branches Tags Go to file Code Open more actions menu Folders and files
2,487 Commits 2,487 Commits .github/ workflows .github/ workflows .kiro/ steering .kiro/ steering agents agents config-examples config-examples config config core core docker/ browser docker/ browser docs/ wiki docs/ wiki packaging/ homebrew packaging/ homebrew plugins/ openclaw-kiro-professor plugins/ openclaw-kiro-professor scripts scripts src src tests tests .env.example .env.example .env.skills.example .env.skills.example .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md esbuild.config.js esbuild.config.js install-manifest.json install-manifest.json install.sh install.sh models.default.json models.default.json package-lock.json package-lock.json package.json package.json tasks.default.json tasks.default.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Autonomous AI agent with persistent memory, self-healing, and peer-to-peer communication. Your hardware, your rules.
abTARS runs 24/7 on your machine — talks to you via Telegram/Discord/IRC, remembers everything across sessions, executes scheduled tasks, and recovers from failures without intervention. Use frontier models through existing CLI subscriptions at zero marginal cost.
🧠 Memory that curates itself — multi-layer recall (5 search stages + reranking), nightly sleep maintenance, emotion tracking, contradiction detection, Memory Darwinism
🛡️ Self-hosted, defense-in-depth — classified memory (4 tiers, encrypted at rest), role-based access, injection scanning, secrets vault
🔄 Runs months unattended — 5-layer supervision, leaky-bucket model fallback, self-healing agent, standby-aware recovery
🤝 Agent-to-agent — P2P communication with Ed25519 signatures, mDNS wake-up, IRC coordination channels
💰 Zero idle cost — no LLM calls at rest, CLI subscription parasitism, budget sleep tiers
→ Why abTARS vs OpenClaw & Hermes
You (Telegram / Discord / IRC / API client)
│
▼
abTARS (bridge)
├── abmind (memory — in-process, multi-layer recall, encrypted)
├── Skills (core + self-authored during sleep + downloadable)
├── Tools (browse, bash, MCP, peer_ask)
├── Tasks (cron scheduler + retry + DoD checks)
├── Agent Swarm (async background sessions)
│
├── kiro-cli → Claude, DeepSeek, MiniMax, Qwen (free tier)
├── gemini-cli → Gemini 2.5
├── Direct API → ollama, OpenRouter, any OpenAI-compatible
└── Peers → other abTARS instances via /v1/chat/completions
Quick Start
npm install -g abtars@alpha abmind@alpha
abmind install
abtars install
abtars update
abtars onboard
sudo $( which abtars ) daemon install
After install, configure ~/.abtars/config/.env with your Telegram bot token and at least one model provider. Full guide: docs/wiki/install.md
Standalone package — works with abTARS, Kiro CLI, Gemini CLI, Claude Code, Codex, Hermes, or any MCP client.
Multi-layer recall: 5 search stages (porter FTS5 → trigram → binary signatures → vector embeddings → entity graph) + 7 post-processing layers (cross-stage penalty, context boost, emotion boost, spacing boost, quality boost, MMR reranking, interference detection)
Agglutinative language support (Hungarian, Finnish, Turkish) — QWERTZ fallback, substring windows
25 emotion types with per-memory scoring, emotional arcs per topic
12-step nightly sleep: extract, consolidate, prune, detect contradictions, fix translations
NATO-inspired classification: confidentiality × trust × integrity × credibility
AES-256-GCM encrypted secrets vault with auto-redaction from history
Memory Darwinism — unused memories fade, recalled memories strengthen
L1 Heartbeat — standby detection, bridge.lock, task dispatch
L2 In-process watchdog — detects stuck event loops
L3 External watchdog — catches dead PIDs, stale heartbeats. Circuit breaker prevents restart storms
L4 OS supervisor — launchd (macOS) / systemd (Linux) restarts the watchdog itself
L5 Preventive daily restart — eliminates memory leaks
Model health — leaky-bucket per model, progressive penalties, arbitrarily long fallback chains
Self-healing agent — diagnoses failed tasks, attempts repair, suspends after 3 failures
Main agent spawns independent background sessions (own context, own tool loop). Results auto-inject into the parent's next prompt. Up to 3 concurrent. /wait injects instructions into running sessions.
Multiple abTARS instances communicate via OpenAI-compatible /v1/chat/completions endpoint. Bearer auth per peer, Ed25519 signatures, mDNS wake-up for firewalled peers, hop-limit loop prevention.
Platform-level access: only registered chatId/userId can reach the agent
Role-based: master/user/guest — commands, tools, memory all gated
Secrets vault: AES-256-GCM, scrypt-derived key, auto-encrypt on ingest
Injection scanner on all inbound messages
Credential redaction in all logs and exports
Transport
Providers
ACP (recommended)
kiro-cli, gemini-cli
Direct API
ollama, OpenRouter, any OpenAI-compatible endpoint
Hooks (standalone)
abmind lifecycle hooks on any CLI agent
Requirements
A Telegram bot token (Discord/IRC optional)
Optional: ollama + nomic-embed-text for memory embeddings.
Full docs: aksika.github.io/abtars
1794 tests (abtars 1016 + abmind 778)
5 agent types (professor, dreamy, browsie, coding, cron)
3 platform adapters + OpenAI-compatible API
12-step nightly memory maintenance
git clone https://github.com/aksika/abtars.git
cd abtars && npm install && npm run build
npm test
Community
GitHub: aksika/abtars · aksika/abmind
Self-hosted AI agent bridge - connects LLMs to Telegram/Discord/IRC with memory, tools, scheduled tasks, agent-swarm and self-healing supervision.
Apache-2.0 license
Contributing
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
