---
source: "https://github.com/Ntooxx/Sentinel"
hn_url: "https://news.ycombinator.com/item?id=48586242"
title: "Sentinel- map any codebase for AI agents in 55s, no cloud, zero deps"
article_title: "GitHub - Ntooxx/Sentinel · GitHub"
author: "Ntox"
captured_at: "2026-06-18T16:15:14Z"
capture_tool: "hn-digest"
hn_id: 48586242
score: 1
comments: 0
posted_at: "2026-06-18T14:46:26Z"
tags:
  - hacker-news
  - translated
---

# Sentinel- map any codebase for AI agents in 55s, no cloud, zero deps

- HN: [48586242](https://news.ycombinator.com/item?id=48586242)
- Source: [github.com](https://github.com/Ntooxx/Sentinel)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T14:46:26Z

## Translation

タイトル: Sentinel - AI エージェントのあらゆるコードベースを 55 秒でマッピング、クラウドなし、デプスなし
記事タイトル: GitHub - Ntooxx/Sentinel · GitHub
説明: GitHub でアカウントを作成して、Ntooxx/Sentinel の開発に貢献します。

記事本文:
GitHub - Ntooxx/Sentinel · GitHub
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
Ntoox
/
センチネル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
30

コミット数 30 コミット数 .github .github ベンチマーク/フィクスチャ ベンチマーク/フィクスチャ config config データ/レポート データ/レポート ドキュメント ドキュメントの例 例 ロゴ ロゴ src src テスト テスト ツール ツール .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CHANGELOG_ARCHETYPE_AWARE.md CHANGELOG_ARCHETYPE_AWARE.md CHANGELOG_risk_surface.md CHANGELOG_risk_surface.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md フロント.html フロント.html install.ps1 install.ps1 pyproject.toml pyproject.toml要件.txt要件.txt Sentinel.md Sentinel.md Sentinel.py Sentinel.py setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クラウドにアップロードせずに、AI にコードベースを理解してもらいたい開発者向け
55 秒で 25,000 個のファイルをスキャン。依存関係ゼロ。 197 件のテスト。
クイックスタート · インストール · コマンド · ダッシュボード · アーキテクチャ
AI コーディング エージェント (Claude Code、Cline、Codex、Continue、Roo) を使用します。彼らはあなたのコードベースを理解する必要がありますが、生のファイルをダンプするとトークンが無駄になり、コンテキストが失われます。
センチネルはこれを解決します。これは、あらゆるリポジトリを構造化されたトークン効率の高いインテリジェンスに変えるローカルの依存性ゼロのスキャナーです。
ポイント → スキャン → AI 対応コンテキスト パック (~2,500 トークン)
アーキテクチャをマッピングし、保守性をスコア付けし、リスクのホットスポットを明らかにし、エントリ ポイントを特定し、AI エージェント用にすぐに使用できるプロンプトを生成します。これらはすべて完全にオフラインで数秒で行われます。アップロードはありません。 API キーはありません。 Python stdlib を超える依存関係はありません。
フローチャート LR
A["📂 任意のリポジトリ"] -->|scan| S["🛡️センチネル"]
S --> B["💊 ヘルススコア"]
S --> C["🔥 ホットスポットとリスク"]
S --> D["🎯 エントリーポイント"]
S --> E["🤖 エージェント プロンプト"]
S --> F["📦コンテキストパック"]

S --> G["💡 次のアクション"]
B & C & D & E & F & G --> H["🧠 AI コーディング エージェント"]
読み込み中
⚡ 30 秒のデモ
# インストール
pip install -e 。
# あらゆるプロジェクトを高速にスキャンします
Python Sentinel.py スキャン 。 --速い
＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝ ═=================================
║ 🛡️ SENTINEL — レポインテリジェンス ║
╠═══════════════════════════════ ═===============================
║ ║
║ プロジェクト kubernetes ║
║ タイプ コンテナ オーケストレーション プラットフォーム ║
║ 健康 ████████████████░░░░ 74% ║
║ ファイル 25,432 ║
║ 6,007,991 行目 ║
║ 時間 55 秒 ║
║ ║
║ ⚠️ 最大のリスク: 5K 行を超える 3 つの特大ファイル ║
║ 💡 次のアクション: kubelet.go を重点モジュールに分割します ║
║ ║
║ テスト数 197 · 失敗数 0 · 外部依存関係なし ║
╚===================================== ═==============================
📊 スキャンパフォーマンス
💡雲はありません。外部サービスはありません。純粋なパイソン。すべてのスキャンは完全にマシン上で実行されます。
名前、タイプ、アーキタイプ、目的、言語、フレームワーク、ワークフロー - ガベージを返さない 5 層のランク付けされたフォールバック システムを通じて解決されます。
保守性、実行時の複雑さ、テスト信号、セキュリティ — 詳細な内訳が記載されているため、問題の箇所が正確にわかります。
プライマリ ランタイム、API サーフェス、サンプル、ビルド ツール、ジェネレーター — インテリジェントなスコアリング (Go バイナリは +80 ボーナスを獲得)。
ランタイム、ビルド、テストランナー、ドキュメント、

ベンダー — リスク別にランク付けされているため、最悪の問題から最初に攻撃できます。
サイズ超過のファイル、TODO の密度、ドキュメントのずれ、テストのギャップなど、すべてのシグナルが実用的です。
影響、労力、信頼度によってランク付けされた提案。単に「これを修正する必要がある」ではなく、どこから始めるべきか。
Cline、Claude Code、Codex、Roo、Continue のすぐに使用できるプロンプト — コピー、ペースト、配布。
コンパクトでトークン効率の高いプロジェクト概要 — 数時間のファイル読み取りに代わる最大 2,500 個のトークン。
コンポーネント、依存関係、アーキタイプ、パターン - 全体像が一目でわかります。
重複排除要素とテストカバレッジを使用したファイルごとのスコアリング - ノイズや重複はありません。
python -m 単体テスト 検出 -s テスト -v
# 197 テスト · 0 失敗 · 9.3 秒
🌟 機能のハイライト
Sentinel は 5 段階のランク付けされたフォールバックを通じてプロジェクト名を解決します。FastAPI をスキャンするときにプロジェクト名として「スポンサー」が表示されなくなります。
┌─ 階層 1: 既知のリポジトリ名 (22 エントリ)
│ FastAPI · Kubernetes · TensorFlow · Flask · Django · React
│ PyTorch · NumPy · Pandas · Vite · Express · Tailwind CSS · …
│
§─ Tier 2: パッケージマニフェスト
│ Cargo.toml · pyproject.toml · package.json · setup.py · go.mod · CMakeLists.txt
│
§─ 階層 3: マニフェストの説明
│ 同じマニフェストから抽出
│
§─ 第 4 層: README 本文
│ 見出しの後の最初の実際の段落
│
└─ 階層 5: README の見出し
ブロックされたセクションのキーワード (インストール、使用法、スポンサーなど) に対して検証済み
🧠 目的の推論
プロジェクトの目的として、プレースホルダーを決して返さない 6 ステップのフォールバック チェーン — もう — — — —
🎯 例: 「Kubernetes: 本番グレードのコンテナ オーケストレーション」 → 「本番グレードのコンテナ オーケストレーション」
go バイナリは、 main.go という名前が付いていなくても検出されます。
cmd/kube-apiserver/apiserver.go → ランタイム エントリ ポイント (+80 スコア)
cmd/kubelet/k

ubelet.go → ランタイム エントリ ポイント (+80 スコア)
cmd/cloud-controller-manager/main.go → ランタイム エントリ ポイント
主要な Go バイナリには +80 スコア ボーナスが付与されます: kube-apiserver 、 kubelet 、 kube-controller-manager 、 kube-scheduler 、 kubectl 、 kube-proxy 、 kubeadm 。
Sentinel は、すべての ID フィールド (プロジェクト名、タイプ、目的、概要) からノイズをフィルターで除去します。
❌ HTML タグ · マークダウン リンク · バッジ · 画像
❌ スポンサーのキーワード · セクションの見出し · 表の成果物
❌ 装飾的な区切り文字 ( ---- 、 ==== など)
生成された HTML レポートは、自己完結型の単一ページです。外部アセットやビルド ステップはありません。
ダークテーマ ブラウザ コマンド センター (http://127.0.0.1:8765):
機能: 統計行 · プロジェクト ID + リスク カード · 共有入力 (クエリ、リポジトリ URL、予算、目標、フラグ) · ピルの切り替え (高速スキャン、ドライラン、適用、検証、アダプター) · ツール カード (理解、質問、レポート、品質、メモリ、メンテナンス、URL 分析) · 出力ターミナル · 提案 + プロンプト · フォーカス/ホットスポット/フレームワーク · ファイル リスク + シグナル テーブルのレビュー · ヘルス タイムライン · 自動更新(3秒)
コマンド
何をするのか
スキャン
プロジェクトの構造、リスク、ホットスポットを分析する
簡単な
重要な提案を含む 1 行の概要
概要
コンポーネント、ホットスポット、ワークフローを含むプロジェクトの完全な説明
コンテキスト
AI エージェント向けのトークン効率の高いプロジェクトの概要
プロンプト
目標の選択を含む次のステップに焦点を当てたプロンプト
取得する
クエリに一致するファイル、シンボル、スニペットを検索する
尋ねる
プロジェクトに関する自然言語の質問に答える
分析URL
git URL を複製し、完全なレポート バンドルを生成します
グラフ
ASTシンボルの抽出、グラフのインポート、グラフの呼び出し
検証する
変更されたファイルに焦点を当てたテストをプレビューまたは実行する
ダッシュボード
ライブブラウザGUIを起動します。
報告書
マークダウンまたは HTML レポートを保存する
広報
変更点、リスク、推奨されるテストの概要
リリースチェック
オープンソース

準備チェックリスト
適用範囲
Coverage.xml からテストが不十分な領域を特定する
タイムライン
スキャン履歴、タスクメモリ、トークンの節約を表示
記憶
タスクメモリの記録またはリスト
貯蓄
トークンの推定節約量を表示する
自動修正
小さな安全な修正を計画または適用する
医者
構成とパスを検証する
mcp
標準入出力 MCP サーバーとして実行
mcp-health
MCP ツールの可用性を検証する
キロセットアップ
Sentinel ファースト ルールで Kilo を構成する
キロブリッジ
no-MCP ファイルブリッジをセットアップする
キロリフレッシュ
タスクの前に Kilo コンテキスト ファイルを更新する
時計
一定間隔で連続スキャン
🏁 クイックスタート
pip install git+https://github.com/Ntooxx/Sentinel.git
ソースから (開発用):
git クローン https://github.com/Ntooxx/Sentinel.git
cdセンチネル
pip install -e 。
Windows ユーザー: install.ps1 をダブルクリックするか、次のコマンドを実行します。
powershell - 実行ポリシーのバイパス - ファイル install.ps1
インストール後、project-sentinel コマンドはグローバルに使用できるようになります。
# 現在のディレクトリをスキャンします
プロジェクトセンチネルスキャン。 --速い
# ライブダッシュボードを起動する
プロジェクトセンチネルダッシュボード 。 --速い
レポートの生成
# 美しい HTML レポート
プロジェクトセンチネルレポート。 --html形式
# マークダウンレポート
プロジェクトセンチネルレポート。 --format マークダウン
AI エージェントのワークフロー
# エージェント対応のプロンプトを生成する
プロジェクトセンチネルプロンプト 。 --次の目標 --予算は少ない --早い
# コードベースについて質問する
プロジェクトセンチネルの質問。 --question 「認証はどこで処理されますか?」 --fast
# GitHub リポジトリを分析する
プロジェクトセンチネル分析 URL https://github.com/user/repo --fast
🤖 トークン節約ワークフロー
トークンの支出を最小限に抑えながら、AI エージェントの効果を最大化します。
# ステップ 1: 全体像を把握する
プロジェクトセンチネルの概要 。 --速い --静か
# ステップ 2: コンパクトなコンテキスト パックを取得する (~2,500 トークン)
プロジェクトセンチネルコンテキスト。 --予算が少ない --速い --静か
# ステップ 3: 焦点を絞った次のステップのプロンプトを取得する

プロジェクトセンチネルプロンプト 。 --次の目標 --予算は少ない --速い --静か
エージェントが受け取るもの:
# 完全なテストスイートを実行する
python -m 単体テスト 検出 -s テスト -v
# 197 テスト · 0 失敗 · 9.3 秒
┌─────────────────────────┐
│ 試験結果 │
│ │
│ ████████████████████████████████████████████████ 100% │
│ │
│ 197 回合格 · 0 回不合格 · 9.3 秒 │
│ 不安定なテストなし、外部依存関係なし │
━━━━━━━━━━━━━━━━━━━━━┘
📈 再現可能なベンチマーク
バンドルされているすべてのフィクスチャ リポジトリに対して Sentinel を実行して、自分のマシンでのパフォーマンスの主張を確認します。
プロジェクトセンチネルベンチマーク。 --速い
実際の実行からの出力例:
センチネルベンチマーク
ベンチマークされた 7 つの試合
cpp_repo ファイル = 2 行 = 6 時間 = 0.007 秒 健全性 = 85%
docs_heavy ファイル = 2 行 = 6 時間 = 0.006 秒 健全性 = 85%
生成された重いファイル = 2 行 = 8 時間 = 0.008 秒 健全性 = 85%
go_service ファイル = 2 行 = 6 時間 = 0.007 秒 健全性 = 85%
node_app ファイル = 2 行 = 19 時間 = 0.006 秒 健全性 = 85%
python_app ファイル = 3 行 = 14 時間 = 0.007 秒 健全性 = 95%
Rust_cli ファイル = 2 行 = 8 時間 = 0.007 秒 健全性 = 85%
ベンチマークは外部依存関係なしで完全にオフラインで実行されます。
すぐに実行できるスクリプトについては、examples/ ディレクトリを参照してください。
# Sentinel リポジトリ自体をスキャンします
プロジェクトセンチネルスキャン。 --速い
# HTML レポートを生成する
プロジェクトセンチネルレポート。 --html形式
# ダッシュボードを起動する
プロジェクトセンチネルダッシュボード 。 --速い
# すべてのフィクスチャ リポジトリでベンチマークを実行する
プロジェクト

ect-sentinel ベンチマーク。 --速い
25,000 ファイル · 600 万行 · 1 つのコマンド · 1 分以内 · クラウドなし
センチネル-nt.netlify.app
トピックス
Apache-2.0ライセンス
行動規範
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

[切り捨てられた]

## Original Extract

Contribute to Ntooxx/Sentinel development by creating an account on GitHub.

GitHub - Ntooxx/Sentinel · GitHub
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
Ntooxx
/
Sentinel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits .github .github benchmarks/ fixtures benchmarks/ fixtures config config data/ reports data/ reports docs docs examples examples logos logos src src tests tests tools tools .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CHANGELOG_ARCHETYPE_AWARE.md CHANGELOG_ARCHETYPE_AWARE.md CHANGELOG_risk_surface.md CHANGELOG_risk_surface.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md front.html front.html install.ps1 install.ps1 pyproject.toml pyproject.toml requirements.txt requirements.txt sentinel.md sentinel.md sentinel.py sentinel.py setup.py setup.py View all files Repository files navigation
For developers who want AI to understand their codebase — without uploading to the cloud
25,000 files scanned in 55 seconds. Zero dependencies. 197 tests.
Quick Start · Install · Commands · Dashboard · Architecture
You use AI coding agents (Claude Code, Cline, Codex, Continue, Roo). They need to understand your codebase — but dumping raw files wastes tokens and misses context.
Sentinel solves this. It's a local, zero-dependency scanner that turns any repo into structured, token-efficient intelligence:
Point → Scan → AI-ready context pack (~2,500 tokens)
It maps architecture, scores maintainability, surfaces risk hotspots, identifies entry points, and generates ready-to-use prompts for your AI agent — all in seconds, entirely offline. No uploads. No API keys. No dependencies beyond Python stdlib.
flowchart LR
A["📂 Any Repo"] -->|scan| S["🛡️ Sentinel"]
S --> B["💊 Health Score"]
S --> C["🔥 Hotspots & Risks"]
S --> D["🎯 Entry Points"]
S --> E["🤖 Agent Prompt"]
S --> F["📦 Context Pack"]
S --> G["💡 Next Actions"]
B & C & D & E & F & G --> H["🧠 AI Coding Agent"]
Loading
⚡ 30-Second Demo
# Install
pip install -e .
# Scan any project — fast
python sentinel.py scan . --fast
╔══════════════════════════════════════════════════════════════╗
║ 🛡️ SENTINEL — Repo Intelligence ║
╠══════════════════════════════════════════════════════════════╣
║ ║
║ Project kubernetes ║
║ Type container orchestration platform ║
║ Health ████████████████░░░░ 74% ║
║ Files 25,432 ║
║ Lines 6,007,991 ║
║ Time 55s ║
║ ║
║ ⚠️ Top risk: 3 oversize files exceeding 5K lines ║
║ 💡 Next action: Split kubelet.go into focused modules ║
║ ║
║ 197 tests · 0 failures · no external dependencies ║
╚══════════════════════════════════════════════════════════════╝
📊 Scan Performance
💡 No cloud. No external services. Pure Python. Every scan runs entirely on your machine.
Name, type, archetype, purpose, language, frameworks, workflow — resolved through a 5-tier ranked fallback system that never returns garbage.
Maintainability, runtime complexity, test signal, security — with a detailed breakdown so you know exactly where the pain is.
Primary runtime, API surfaces, examples, build tools, generators — with intelligent scoring (Go binaries get +80 bonus).
Runtime, build, test runner, documentation, vendor — ranked by risk so you attack the worst problems first.
Oversized files, TODO density, documentation drift, test gaps — every signal is actionable.
Suggestions ranked by impact , effort , and confidence — not just "you should fix this" but where to start .
Ready-to-use prompt for Cline, Claude Code, Codex, Roo, Continue — copy, paste, ship.
Compact, token-efficient project brief — ~2,500 tokens that replace hours of file reading.
Components, dependencies, archetype, patterns — the big picture at a glance.
Per-file scoring with deduplicated factors and test coverage — no noise, no duplicates.
python -m unittest discover -s tests -v
# 197 tests · 0 failures · 9.3 seconds
🌟 Feature Highlights
Sentinel resolves project names through a 5-tier ranked fallback — no more "Sponsors" as a project name when scanning FastAPI:
┌─ Tier 1: Known repo names (22 entries)
│ FastAPI · Kubernetes · TensorFlow · Flask · Django · React
│ PyTorch · NumPy · Pandas · Vite · Express · Tailwind CSS · …
│
├─ Tier 2: Package manifests
│ Cargo.toml · pyproject.toml · package.json · setup.py · go.mod · CMakeLists.txt
│
├─ Tier 3: Manifest descriptions
│ Extracted from the same manifests
│
├─ Tier 4: README body
│ First real paragraph after headings
│
└─ Tier 5: README heading
Validated against blocked section keywords (Installation, Usage, Sponsors, …)
🧠 Purpose Inference
A 6-step fallback chain that never returns a placeholder — no more ---- as project purpose:
🎯 Example: "Kubernetes: Production-Grade Container Orchestration" → "Production-Grade Container Orchestration"
Go binaries are detected even when not named main.go :
cmd/kube-apiserver/apiserver.go → runtime entry point (+80 score)
cmd/kubelet/kubelet.go → runtime entry point (+80 score)
cmd/cloud-controller-manager/main.go → runtime entry point
Major Go binaries get a +80 score bonus : kube-apiserver , kubelet , kube-controller-manager , kube-scheduler , kubectl , kube-proxy , kubeadm .
Sentinel filters out the noise from all identity fields (project name, type, purpose, summary):
❌ HTML tags · Markdown links · Badges · Images
❌ Sponsor keywords · Section headings · Table artifacts
❌ Decorative separators ( ---- , ==== , etc.)
The generated HTML report is a single self-contained page — no external assets, no build step:
Dark-theme browser command centre at http://127.0.0.1:8765 :
Features: Stats row · Project identity + risk cards · Shared inputs (query, repo URL, budget, goal, flags) · Toggle pills (fast scan, dry-run, apply, verify, adapters) · Tool cards (Understand, Ask, Reports, Quality, Memory, Maintenance, Analyze URL) · Output terminal · Suggestions + prompt · Focus/hotspots/frameworks · File risks + review signals tables · Health timeline · Auto-refresh (3s)
Command
What It Does
scan
Analyse project structure, risks, hotspots
brief
One-line summary with the top suggestion
overview
Full project description with components, hotspots, workflow
context
Token-efficient project brief for AI agents
prompt
Focused next-step prompt with goal selection
retrieve
Find files, symbols, and snippets matching a query
ask
Answer a natural-language question about the project
analyze-url
Clone a git URL and generate a complete report bundle
graph
Extract AST symbols, import graph, call graph
verify
Preview or run focused tests for changed files
dashboard
Launch the live browser GUI
report
Save a Markdown or HTML report
pr
Summarise changes, risks, and suggested tests
release-check
Open-source readiness checklist
coverage
Identify weakly tested areas from coverage.xml
timeline
Show scan history, task memory, and token savings
memory
Record or list task memory
savings
Show estimated token savings
autofix
Plan or apply small safe fixes
doctor
Validate configuration and paths
mcp
Run as a stdio MCP server
mcp-health
Validate MCP tool availability
kilo-setup
Configure Kilo with Sentinel-first rules
kilo-bridge
Set up the no-MCP file bridge
kilo-refresh
Refresh Kilo context files before a task
watch
Continuously scan at an interval
🏁 Quick Start
pip install git+https://github.com/Ntooxx/Sentinel.git
From source (for development):
git clone https://github.com/Ntooxx/Sentinel.git
cd Sentinel
pip install -e .
Windows users: double-click install.ps1 or run:
powershell - ExecutionPolicy Bypass - File install.ps1
After install, the project-sentinel command is available globally.
# Scan the current directory
project-sentinel scan . --fast
# Launch the live dashboard
project-sentinel dashboard . --fast
Generate Reports
# Beautiful HTML report
project-sentinel report . --format html
# Markdown report
project-sentinel report . --format markdown
AI Agent Workflow
# Generate an agent-ready prompt
project-sentinel prompt . --goal next --budget small --fast
# Ask a question about your codebase
project-sentinel ask . --question " where is authentication handled? " --fast
# Analyse any GitHub repo
project-sentinel analyze-url https://github.com/user/repo --fast
🤖 Token-Saving Workflow
Maximize your AI agent's effectiveness while minimizing token spend:
# Step 1: Get the big picture
project-sentinel overview . --fast --quiet
# Step 2: Get a compact context pack (~2,500 tokens)
project-sentinel context . --budget small --fast --quiet
# Step 3: Get a focused next-step prompt
project-sentinel prompt . --goal next --budget small --fast --quiet
What the agent receives:
# Run the full test suite
python -m unittest discover -s tests -v
# 197 tests · 0 failures · 9.3 seconds
┌─────────────────────────────────────────────────────────┐
│ Test Results │
│ │
│ ████████████████████████████████████████████████ 100% │
│ │
│ 197 passed · 0 failed · 9.3s │
│ No flaky tests · No external dependencies │
└─────────────────────────────────────────────────────────┘
📈 Reproducible Benchmark
Run Sentinel against all bundled fixture repos to verify performance claims on your own machine:
project-sentinel benchmark . --fast
Example output from a real run:
SENTINEL BENCHMARK
Benchmarked 7 fixture(s)
cpp_repo files= 2 lines= 6 time= 0.007s health=85%
docs_heavy files= 2 lines= 6 time= 0.006s health=85%
generated_heavy files= 2 lines= 8 time= 0.008s health=85%
go_service files= 2 lines= 6 time= 0.007s health=85%
node_app files= 2 lines= 19 time= 0.006s health=85%
python_app files= 3 lines= 14 time= 0.007s health=95%
rust_cli files= 2 lines= 8 time= 0.007s health=85%
Benchmarks run entirely offline with zero external dependencies.
See the examples/ directory for ready-to-run scripts:
# Scan the Sentinel repo itself
project-sentinel scan . --fast
# Generate an HTML report
project-sentinel report . --format html
# Launch the dashboard
project-sentinel dashboard . --fast
# Run a benchmark on all fixture repos
project-sentinel benchmark . --fast
25,000 files · 6 million lines · One command · Under a minute · No cloud
sentinel-nt.netlify.app
Topics
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal infor

[truncated]
