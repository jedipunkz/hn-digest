---
source: "https://github.com/raro42/mac-stats"
hn_url: "https://news.ycombinator.com/item?id=49248085"
title: "I put a local LLM agent inside a menu-bar Mac monitor (and you can leave it off)"
article_title: "GitHub - raro42/mac-stats: Menu-bar system monitor for Apple Silicon (optional local AI agent). Rust + Tauri. No cloud telemetry. · GitHub"
author: "raro43"
captured_at: "2026-08-10T19:49:06Z"
capture_tool: "hn-digest"
hn_id: 49248085
score: 1
comments: 0
posted_at: "2026-08-10T18:57:59Z"
tags:
  - hacker-news
  - translated
---

# I put a local LLM agent inside a menu-bar Mac monitor (and you can leave it off)

- HN: [49248085](https://news.ycombinator.com/item?id=49248085)
- Source: [github.com](https://github.com/raro42/mac-stats)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T18:57:59Z

## Translation

タイトル: メニューバーの Mac モニター内にローカル LLM エージェントを配置しました (オフのままにすることもできます)
記事のタイトル: GitHub - raro42/mac-stats: Apple Silicon 用のメニューバー システム モニター (オプションのローカル AI エージェント)。ラスト＋タウリ。クラウドテレメトリーはありません。 · GitHub
説明: Apple Silicon 用のメニューバー システム モニター (オプションのローカル AI エージェント)。ラスト＋タウリ。クラウドテレメトリーはありません。 - raro42/mac-stats

記事本文:
GitHub - raro42/mac-stats: Apple Silicon 用のメニューバー システム モニター (オプションのローカル AI エージェント)。ラスト＋タウリ。クラウドテレメトリーはありません。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ラロ42
/
mac-stats
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,233 コミット 1,233 コミット .github .github .mac-stats .mac-stats .vscode .vscode Casks Casks エージェント エージェント docs docs homebrew-tap homebrew-tap スクリプト スクリプト src-tau

ri src-tauri src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md FEATURES.md FEATURES.md README.md README.md Agents.md Agents.md config.example.json config.example.json config.minimal.json config.minimal.json run run run.sh run.sh すべてのファイルを表示リポジトリ ファイルのナビゲーション
Apple Silicon 用のメニューバー システム モニター (オプションのローカル AI エージェント)。
Apple Silicon のみ (arm64)。 Intel Mac は、公開された DMG / Homebrew Cask ではサポートされていません。
無料の MIT、ローカルファースト、クラウド テレメトリなし。メニューバーに一目でわかる主要な指標。必要に応じてオプションの Ollama / Discord エージェントを使用できます。デフォルトではオフになっています。
1 つのバイナリに 2 つの製品 — パスを選択してください:
📋 変更履歴 · ✨ 機能 · 📘 はじめに · 🗺 ロードマップ · 🍺 自作 · 🌐 ランディング · 🔬 比較方法
クイックスタート — モニターだけ
クイックスタート — モニター + AI エージェント
vs. 統計 / iStat メニュー / メニューメーター
クイックスタート — モニターだけ
醸造タップ raro42/mac-stats https://github.com/raro42/mac-stats
brew install --cask mac-stats
# または: ./scripts/quickstart.sh # クローンから;インストール + シード ~/.mac-stats
open -a mac-stats
メニューバーを見て→クリックするとウィンドウが表示されます。オラマは必要ありません。ディスク クリーンアップ、モニター、テーマはすべてガラス窓内にあります。
クイックスタート — モニター + AI エージェント
Ollama をインストールしてモデルをプルします。
カール -fsSL https://ollama.com/install.sh |しー
オラマ プル ラマ3.2
[設定] で AI を有効にする (ローカル AI エージェントを有効にする) または:
# ~/.mac-stats/config.json 内
{ " aiAgentEnabled " : true、 " menuBarCompact " : true }
ウィンドウを開く → AI チャット (Ollama) → 試してください: 私の CPU 温度は何度ですか?
詳細: docs/GETTING_STARTED.md 。
~49 秒のライブ ウィンドウ キャプチャ — 本物の mac_stats --cpu セッション (ScreenCaptureKit、ウィンドウのみ): ライブ ゲージ、モニター (レッド ダウン サイトを含む)、エージェント オペレーション、Ollama チャット。ライト付きで 1080p のレターボックス化

ナレーション。 (リポジトリファイル)
完全なテーマ ギャラリーとキャプチャ ノート: docs/screens/README.md 。
リポジトリ: github.com/raro42/mac-stats
vs. 統計 / iStat メニュー / メニューメーター
2026 年のメニューバー モニター スキャンで最も近いピア (比較方法):
vs Stats — Stats は古典的な無料の OSS モニターです。 mac-stats は、ガラス/テーマ、ディスク クリーンアップ (デフォルトでゴミ箱へのソフト削除による範囲指定された再利用)、およびオプションのローカル AI/Discord/スケジュールでその精神を維持しています。
vs iStat Menus — iStat Menus は有料のディープセンサーベンチマークです。 mac-stats は、必需品 + Apple Silicon の一目瞭然 + 軽量クリーンアップを維持します (iStat の完全な代替品ではありません)。
vs MenuMeters — MenuMeters は最小限のバーです。 mac-stats は、AI を強制することなく、より豊富なメトリクス、テーマ、ディスク クリーンアップ、およびオプションのエージェント ワークフローを追加します。
統計のようなモニターの場合のみ、AI をオフのままにしておきます。ローカルファースト: コア指標にはクラウドは必要ありません。
方法
コマンド/リンク
自家製樽
brew Tap raro42/mac-stats https://github.com/raro42/mac-stats && brew install --cask mac-stats
クイックスタートスクリプト
./scripts/quickstart.sh (クローン) — アプリ + ~/.mac-stats デフォルト + Ollama チェック
ダメージ
リリース
ソース
リリースタグを固定します。 「ソースからのビルド」を参照してください。
ゲートキーパー / 公証: 署名 + 公証済みビルド ( docs/NOTALIZATION.md ) を優先します。 CI シークレットが設定されるまでは、「右クリック」→「開く」を使用します。
リポジトリ ルートの構成テンプレート: config.minimal.json (モニターのみ)、config.example.json (AI 有効)。
クラウド テレメトリはありません。すべては ~/.mac-stats/ に残ります。シークレット: キーチェーンおよび/または .config.env (決してコミットしない)。 docs/CONFIG.md を参照してください。
アプリ内バナーは GitHub リリースをチェックします。または: brew upgrade --cask mac-stats 。
デフォルトではコンパクトなメニュー バー (CPU + SSD、既知の場合は °C を加えます)。 CPU/GPU/RAM/SSD の場合は、menuBarCompact: false を設定します。
9 つのテーマ、プロセス リスト、Web サイト モニター (メニュー バーに赤い月 ✕ キューが表示されます)

ニューヨークのサイトはダウンしています）。
CPU / 周波数 / 温度と並んで GPU 使用率リング ゲージ。
~0.5% のアイドル状態の CPU (メニュー バーのみ)。
組み込みの再生パネル (AI は不要):
再利用可能なサイズをプレビューします。今すぐクリーンアップするか、アプリの起動時および実行中 24 時間ごとに自動的に実行されます。
オン/オフを切り替えることができるスコープ: mac-stats データ、ゴミ箱、ダウンロード、一時ファイル、およびカスタム パス (経過日数 + 再帰)。 discCleanupScopes に保存されます。
デフォルトでは論理的な削除 - クリーンアップされたファイルはゴミ箱に移動されるため、復元できます。完全に削除するには、[クリーンアップされたアイテムをゴミ箱に移動] チェックボックスをオフにします (または、diskCleanupSoftDelete: false を設定します)。ゴミ箱スコープを空にすることは常に永続的です。
詳細: FEATURES.md · 上のスクリーンショット。
Ollama チャット、Discord (Werner)、FETCH_URL、Brave、Perplexity、CDP ブラウザ、タスク、スケジューラ、MCP、Agent Ops。
aiAgentEnabled: true までオフ。
docs/CONFIG.md · 設定 → モニターのデフォルトにリセットします。
コマンド
説明
mac_stats / open -a mac-stats
スタート
mac_stats --cpu
窓を開けた状態でスタート
mac_stats -vv
詳細な debug.log
ソースからビルドする
git clone https://github.com/raro42/mac-stats.git
cd mac-stats
git checkout v0.1.257 # 可能な場合はピンを留めます
./実行
Rust + Xcode CLT (macOS Tauri) が必要です。チェックサム: ./scripts/print-release-checksums.sh v0.1.257 。
寄稿者ドキュメント: docs/design/ 。ワークフロー: docs/agent_workflow.md 。
不和 · ディスカッション · 問題 · フィードバック
Apple Silicon 用のメニューバー システム モニター (オプションのローカル AI エージェント)。ラスト＋タウリ。クラウドテレメトリーはありません。
github.com/raro42/mac-stats/blob/main/docs/GETTING_STARTED.md トピック
4 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Menu-bar system monitor for Apple Silicon (optional local AI agent). Rust + Tauri. No cloud telemetry. - raro42/mac-stats

GitHub - raro42/mac-stats: Menu-bar system monitor for Apple Silicon (optional local AI agent). Rust + Tauri. No cloud telemetry. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
raro42
/
mac-stats
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,233 Commits 1,233 Commits .github .github .mac-stats .mac-stats .vscode .vscode Casks Casks agents agents docs docs homebrew-tap homebrew-tap scripts scripts src-tauri src-tauri src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md FEATURES.md FEATURES.md README.md README.md agents.md agents.md config.example.json config.example.json config.minimal.json config.minimal.json run run run.sh run.sh View all files Repository files navigation
Menu-bar system monitor for Apple Silicon (optional local AI agent).
Apple Silicon only (arm64). Intel Macs are not supported by the published DMG / Homebrew cask.
Free MIT · local-first · no cloud telemetry. Core glanceable metrics in the menu bar; optional Ollama / Discord agent when you want it — off by default.
Two products in one binary — pick your path:
📋 Changelog · ✨ Features · 📘 Getting Started · 🗺 Roadmap · 🍺 Homebrew · 🌐 Landing · 🔬 How we compare
Quick start — Just the monitor
Quick start — Monitor + AI agent
vs. Stats / iStat Menus / MenuMeters
Quick start — Just the monitor
brew tap raro42/mac-stats https://github.com/raro42/mac-stats
brew install --cask mac-stats
# or: ./scripts/quickstart.sh # from a clone; installs + seeds ~/.mac-stats
open -a mac-stats
Look at the menu bar → click for the window. No Ollama required. Disk Cleanup, monitors, and themes are all in the glass window.
Quick start — Monitor + AI agent
Install Ollama and pull a model:
curl -fsSL https://ollama.com/install.sh | sh
ollama pull llama3.2
Enable AI in Settings ( Enable local AI agent ) or:
# in ~/.mac-stats/config.json
{ " aiAgentEnabled " : true, " menuBarCompact " : true }
Open the window → AI Chat (Ollama) → try: What's my CPU temp?
Details: docs/GETTING_STARTED.md .
~49s live window capture — real mac_stats --cpu session (ScreenCaptureKit, window-only): live gauges, monitors (including a red down site), Agent Ops, Ollama chat. Letterboxed to 1080p with light voiceover. ( repo file )
Full theme gallery and capture notes: docs/screens/README.md .
Repo: github.com/raro42/mac-stats
vs. Stats / iStat Menus / MenuMeters
Closest peers from a 2026 menu-bar monitor scan ( how we compare ):
vs Stats — Stats is the classic free OSS monitor; mac-stats keeps that spirit with glass/themes, Disk Cleanup (scoped reclaim with soft-delete to Trash by default), and optional local AI / Discord / schedules.
vs iStat Menus — iStat Menus is the paid deep-sensor benchmark; mac-stats stays on essentials + Apple Silicon glanceability + lightweight cleanup (we are not a full iStat replacement).
vs MenuMeters — MenuMeters is minimal bars; mac-stats adds richer metrics, themes, Disk Cleanup, and optional agent workflows without forcing AI on.
Leave AI off for a Stats-like monitor only. Local-first: core metrics never need the cloud.
Method
Command / link
Homebrew cask
brew tap raro42/mac-stats https://github.com/raro42/mac-stats && brew install --cask mac-stats
Quick Start script
./scripts/quickstart.sh (clone) — app + ~/.mac-stats defaults + Ollama check
DMG
Releases
Source
Pin a release tag; see Build from source
Gatekeeper / notarization: Prefer signed+notarized builds ( docs/NOTARIZATION.md ). Until CI secrets are set, use Right-click → Open .
Config templates in repo root: config.minimal.json (monitor-only), config.example.json (AI enabled).
No cloud telemetry — everything stays in ~/.mac-stats/ . Secrets: Keychain and/or .config.env (never commit). See docs/CONFIG.md .
In-app banner checks GitHub Releases. Or: brew upgrade --cask mac-stats .
Compact menu bar by default ( CPU + SSD , plus °C when known); set menuBarCompact: false for CPU/GPU/RAM/SSD.
Nine themes, process list, website monitors (menu bar shows a red Mon ✕ cue when any site is down).
GPU usage ring gauge alongside CPU / frequency / temperature.
~0.5% idle CPU (menu bar only).
Built-in reclaim panel (no AI required):
Preview reclaimable size; Clean now , or automatic runs on app launch and every 24h while running.
Scopes you can turn on/off: mac-stats data, Trash, Downloads, Temp, plus custom paths (age in days + recurse). Saved in diskCleanupScopes .
Soft-delete by default — cleaned files go to Trash so you can recover them. Uncheck Move cleaned items to Trash (or set diskCleanupSoftDelete: false ) for permanent delete. Emptying the Trash scope is always permanent.
Details: FEATURES.md · screenshot above.
Ollama chat, Discord (Werner), FETCH_URL, Brave, Perplexity, CDP browser, tasks, scheduler, MCP, Agent Ops.
Off until aiAgentEnabled: true .
docs/CONFIG.md · Settings → Reset to monitor defaults .
Command
Description
mac_stats / open -a mac-stats
Start
mac_stats --cpu
Start with window open
mac_stats -vv
Verbose debug.log
Build from source
git clone https://github.com/raro42/mac-stats.git
cd mac-stats
git checkout v0.1.257 # pin when possible
./run
Requires Rust + Xcode CLT (macOS Tauri). Checksums: ./scripts/print-release-checksums.sh v0.1.257 .
Contributor docs: docs/design/ . Workflow: docs/agent_workflow.md .
Discord · Discussions · Issues · Feedback
Menu-bar system monitor for Apple Silicon (optional local AI agent). Rust + Tauri. No cloud telemetry.
github.com/raro42/mac-stats/blob/main/docs/GETTING_STARTED.md Topics
4 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
