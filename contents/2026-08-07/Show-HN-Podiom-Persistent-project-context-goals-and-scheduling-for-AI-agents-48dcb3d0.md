---
source: "https://github.com/Podiom/Podiom"
hn_url: "https://news.ycombinator.com/item?id=49216407"
title: "Show HN: Podiom – Persistent project context, goals and scheduling for AI agents"
article_title: "GitHub - Podiom/Podiom: Thin orchestration layer for local LLM agents. Durable sessions, profiles, scheduling, and native MCP/tool/skill integration. · GitHub"
author: "Maphielbso"
captured_at: "2026-08-07T21:28:44Z"
capture_tool: "hn-digest"
hn_id: 49216407
score: 1
comments: 0
posted_at: "2026-08-07T21:24:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Podiom – Persistent project context, goals and scheduling for AI agents

- HN: [49216407](https://news.ycombinator.com/item?id=49216407)
- Source: [github.com](https://github.com/Podiom/Podiom)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T21:24:57Z

## Translation

タイトル: HN を表示: Podiom – AI エージェントの永続的なプロジェクト コンテキスト、目標、スケジュール
記事のタイトル: GitHub - Podiom/Podiom: ローカル LLM エージェント用の薄いオーケストレーション レイヤー。耐久性のあるセッション、プロファイル、スケジューリング、ネイティブ MCP/ツール/スキルの統合。 · GitHub
説明: ローカル LLM エージェント用の薄いオーケストレーション層。耐久性のあるセッション、プロファイル、スケジューリング、ネイティブ MCP/ツール/スキルの統合。 - ポディオム/ポディオム

記事本文:
GitHub - Podiom/Podiom: ローカル LLM エージェント用の薄いオーケストレーション層。耐久性のあるセッション、プロファイル、スケジューリング、ネイティブ MCP/ツール/スキルの統合。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ポディオム
/
ポディオム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
188 コミット 188 コミット .claude .claude .github .github cmd cmd docs docs ha ha i

内部 内部スクリプト スクリプト Web Web .gitignore .gitignore AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカル LLM エージェント (Claude Code および OpenAI Codex) 用の薄いオーケストレーション レイヤー。
Podiom はネイティブのクロード CLI とコーデックス CLI をシェルとして使用し、それらの CLI に依存します。
MCP、ツール、スキルは、独自の永続的な真実を所有しながら、名前付きエージェント、永続的
チャット セッション、新しいバッキング CLI セッションで再生される正規の履歴
任意のプロファイル/プロバイダー スイッチ、組み込みスケジューラ、および共有プロジェクト上で
台帳。 Svelte Web UI が組み込まれた単一の Go バイナリとして出荷されます。
複数のローカル エージェントの管理は、始めるのは簡単ですが、一貫性を保つのは困難です。ポディオム
意図的に薄いままです:
プロバイダーやプロファイルの変更に耐える耐久性のあるセッション。
共有プロジェクト台帳なので、実行間で作業が失われることはありません。
定期的な作業とフォローアップのための組み込みのスケジュール。
すでに使用しているツールを置き換えるのではなく、ネイティブに統合します。
エージェント名簿
チャットセッション
目標のタイムライン
これは誰のためのものですか?
すでに Claude Code または OpenAI Codex をローカルで使用している開発者。
永続的でレビュー可能なエージェントの作業を望むオープンソース ビルダー。
クラウド ロックインよりもローカル ファーストのワークフローを好むオペレーターやいじくり屋。
既存のエージェント ツールの周りに軽量のコントロール プレーンを必要とする保守者。
カール -fsSL https://github.com/Podiom/Podiom/releases/latest/download/install.sh |バッシュ
Windows PowerShell:
irm https://github.com/Podiom/Podiom/リリース/最新/ダウンロード/install.ps1 |アイエックス
インストーラーは、一致するリリース バイナリをダウンロードし、チェックサムを検証し、設定することができます。
ユーザーレベルの自動起動を開始し、オンボードで Podiom を起動します

クロード/コーデックスを確認し、
最初のエージェントを作成します。
マスターへのコミットごとに、自動
v0.1.<run-number> シリーズ。そのシリーズは意図的に単調になっています。
カレンダーベースであるため、集中的な作業によって多くのリリースが生成される可能性がありますが、
毎月のペース。
インストール後、CLI または Web UI からアップデートを確認して適用できます。
Podiomアップデートチェック
podiom 更新適用 --はい
Linux リリースは、ディストリビューション中立の静的バイナリです。
前提条件: Go 1.26 以降、Node 20 以降 (Web UI の構築用)。
# Web UI (vite) と両方のバイナリをバージョン スタンプを使用して bin/ にビルドします。
ビルドする
# デーモンを実行します (フォアグラウンド)。最初の実行時に ~/.podiom をスキャフォールディングします。
./bin/podiomd
# 別のシェルで、ライブであることを確認します。
./bin/podiom ステータス
Web UI として http://127.0.0.1:8787 を開きます。
ホット リロードを使用してフロントエンドを開発するには、web/ で npm run dev を実行します (プロキシ
実行中の podiomd への API/WebSocket トラフィック)。
クロスプラットフォームのビルドとパッケージ化
podiomd は、SPA が組み込まれた単一の静的バイナリです。外部アセットはありません。
cgo ( modernc.org/sqlite 経由の純粋な Go SQLite) がないため、きれいにクロスコンパイルされます。
makecross # linux/darwin/windows × amd64/arm64 → bin/<os>-<arch>/
make package # アーティファクトを dist/ にアーカイブし、SHA256SUMS を書き込みます
すべてのランタイム状態は 1 つのオーバーライド可能なルートの下に存在するため、Podiom をホームとして実行します
コンテナ内のアシスタント アドオンは、パッケージ化の手順であり、書き換えではありません。
PODIOM_HOME=/data/podiom ./bin/podiomd # 相対値は絶対値に固定されます
Web バインドは config.yaml (server.bind /server.port、
デフォルトは 127.0.0.1:8787 ); 「構成」を参照してください。
cmd/podiom/シン CLI クライアント
cmd/podiomd/ デーモン: Web サーバー + スケジューラー + コア
内部/コア、アダプター、実行、スケジュール、構成、ストア、サーバー、クライアント
web/ Svelte + Vite + TS + Tailwind SPA (ビルド

t → 埋め込み)
ドキュメント/要件、CLI リファレンス、構成、統合契約
すべてのランタイム状態は $PODIOM_HOME (デフォルト ~/.podiom/ ) の下に存在します。
要件 — 正式な仕様 (v1.6)。
エージェント - 永続的な名前付き同僚とその保存されたデフォルト
Git — プロジェクトがソース管理を行う方法
セッション — 永続的な会話単位
SOUL.md の生成 - エージェント ID ファイルの生成方法
目標 — 結果をエージェントに渡します。計画し、レビューし、報告します
ワークスペース ツール - 承認されたエージェントごとの CLI インストール
音声入力 — チャット、タスク、目標でプロンプトを読み上げる (OpenAI Whisper)
写真の添付ファイル — クロードまたはコーデックスが検査するために保持されている写真を添付します。
セキュリティとロギング - 許可モード、ゲートウェイトークン、リダクション、実行ログ
ホーム アシスタント アプリ — Podiom を HA アドオンとして展開
Podiom は MIT ライセンスの下でオープンソースです。寄付金は、
ようこそ。セットアップ、検証については CONTRIBUTING.md をお読みください。
およびプルリクエストのガイドライン。
ローカル LLM エージェント用の薄いオーケストレーション層。耐久性のあるセッション、プロファイル、スケジューリング、ネイティブ MCP/ツール/スキルの統合。
github.com/Podiom/Podiom/tree/master/docs トピック
Readme MIT ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
4 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Thin orchestration layer for local LLM agents. Durable sessions, profiles, scheduling, and native MCP/tool/skill integration. - Podiom/Podiom

GitHub - Podiom/Podiom: Thin orchestration layer for local LLM agents. Durable sessions, profiles, scheduling, and native MCP/tool/skill integration. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Podiom
/
Podiom
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
188 Commits 188 Commits .claude .claude .github .github cmd cmd docs docs ha ha internal internal scripts scripts web web .gitignore .gitignore AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum View all files Repository files navigation
A thin orchestration layer for local LLM agents (Claude Code and OpenAI Codex).
Podiom shells out to the native claude and codex CLIs and leans on their
MCP, tools, and skills, while owning its own durable truth: named agents, durable
chat sessions, a canonical history that replays onto a fresh backing CLI session
on any profile/provider switch, an embedded scheduler, and a shared project
ledger. It ships as a single Go binary with an embedded Svelte web UI.
Managing multiple local agents is easy to start and hard to keep coherent. Podiom
stays thin on purpose:
Durable sessions that survive provider and profile changes.
A shared project ledger so work does not get lost between runs.
Built-in scheduling for recurring work and follow-ups.
Native integration with the tools you already use instead of replacing them.
Agent roster
Chat session
Goal timeline
Who is this for?
Developers already using Claude Code or OpenAI Codex locally.
Open-source builders who want persistent, reviewable agent work.
Operators and tinkerers who prefer local-first workflows over cloud lock-in.
Maintainers who need a lightweight control plane around existing agent tools.
curl -fsSL https://github.com/Podiom/Podiom/releases/latest/download/install.sh | bash
Windows PowerShell:
irm https: // github.com / Podiom / Podiom / releases / latest / download / install.ps1 | iex
The installer downloads the matching release binary, verifies checksums, can set
up user-level autostart, and launches podiom onboard to check Claude/Codex and
create your first agent.
Every commit to master publishes a GitHub Release using the automatic
v0.1.<run-number> series. That series is intentionally monotonic rather than
calendar-based, so bursts of work can produce many releases without implying a
monthly cadence.
After install, updates can be checked and applied from the CLI or web UI:
podiom update check
podiom update apply --yes
Linux releases are distro-neutral static binaries.
Prerequisites: Go 1.26+, Node 20+ (for building the web UI).
# Build the web UI (vite) and both binaries into bin/ with a version stamp.
make build
# Run the daemon (foreground). It scaffolds ~/.podiom on first run.
./bin/podiomd
# In another shell, check it's live.
./bin/podiom status
Open http://127.0.0.1:8787 for the web UI.
To develop the frontend with hot reload, run npm run dev in web/ (it proxies
API/WebSocket traffic to a running podiomd ).
Cross-platform builds & packaging
podiomd is a single static binary with the SPA embedded — no external assets,
no cgo (pure-Go SQLite via modernc.org/sqlite ), so it cross-compiles cleanly:
make cross # linux/darwin/windows × amd64/arm64 → bin/<os>-<arch>/
make package # archives release artifacts into dist/ and writes SHA256SUMS
All runtime state lives under one overridable root, so running Podiom as a Home
Assistant add-on or in a container is a packaging step, not a rewrite:
PODIOM_HOME=/data/podiom ./bin/podiomd # relative values are anchored absolute
The web bind is configurable in config.yaml ( server.bind / server.port ,
default 127.0.0.1:8787 ); see Configuration .
cmd/podiom/ thin CLI client
cmd/podiomd/ daemon: web server + scheduler + core
internal/ core, adapter, exec, schedule, config, store, server, client
web/ Svelte + Vite + TS + Tailwind SPA (built → embedded)
docs/ requirements, CLI reference, configuration, integration contracts
All runtime state lives under $PODIOM_HOME (default ~/.podiom/ ).
Requirements — the authoritative spec (v1.6).
Agents — durable, named colleagues and their stored defaults
Git — how projects carry source control
Sessions — the durable conversation unit
SOUL.md generation — how agent identity files are generated
Goals — hand an outcome to an agent; it plans, reviews, and reports back
Workspace tools — approved per-agent CLI installs
Voice input — speak prompts in chat, tasks, and goals (OpenAI Whisper)
Photo attachments — attach retained photos for Claude or Codex to inspect
Security & logging — permission modes, gateway token, redaction, run logs
Home Assistant app — deploy Podiom as an HA add-on
Podiom is open source under the MIT License . Contributions are
welcome; please read CONTRIBUTING.md for setup, validation,
and pull request guidelines.
Thin orchestration layer for local LLM agents. Durable sessions, profiles, scheduling, and native MCP/tool/skill integration.
github.com/Podiom/Podiom/tree/master/docs Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
4 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
