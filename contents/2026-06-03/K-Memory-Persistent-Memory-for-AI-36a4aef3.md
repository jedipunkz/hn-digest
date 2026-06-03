---
source: "https://github.com/mackopofa/k-memory"
hn_url: "https://news.ycombinator.com/item?id=48376087"
title: "K-Memory – Persistent Memory for AI"
article_title: "GitHub - mackopofa/k-memory: K-Memory — Persistent, self-installing memory for AI agents. Zero dependencies. Pure Python. · GitHub"
author: "kensaiArt"
captured_at: "2026-06-03T00:36:00Z"
capture_tool: "hn-digest"
hn_id: 48376087
score: 4
comments: 0
posted_at: "2026-06-02T20:51:16Z"
tags:
  - hacker-news
  - translated
---

# K-Memory – Persistent Memory for AI

- HN: [48376087](https://news.ycombinator.com/item?id=48376087)
- Source: [github.com](https://github.com/mackopofa/k-memory)
- Score: 4
- Comments: 0
- Posted: 2026-06-02T20:51:16Z

## Translation

タイトル: K-Memory – AI 用永続メモリ
記事のタイトル: GitHub - mackopofa/k-memory: K-Memory — AI エージェント用の永続的な自己インストール型メモリ。依存関係ゼロ。純粋なパイソン。 · GitHub
説明: K-Memory — AI エージェント用の永続的な自己インストール型メモリ。依存関係ゼロ。純粋なパイソン。 - mackopofa/k-memory

記事本文:
GitHub - mackopofa/k-memory: K-Memory — AI エージェント用の永続的な自己インストール型メモリ。依存関係ゼロ。純粋なパイソン。 · GitHub
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
マッコポファ
/
kメモリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット エクストラ エクストラ テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md install.sh install.sh k-core.py k-core.py k-detector.py k-detector.py すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI エージェント用の永続メモリ。依存関係ゼロ。純粋なパイソン。
LLM およびエージェント用の自己完結型メモリ エンジン。 Python stdlib のみを使用して、ナレッジを保存、リンク、要約、エクスポートします。ベクトル、クラウド、ロックインはありません。
bash <(curl -fsSL https://raw.githubusercontent.com/mackopofa/k-memory/main/install.sh )
または手動で:
git clone https://github.com/mackopofa/k-memory.git ~ /k-memory
cd ~ /k-memory && python3 k-core.py
クイックスタート
# ファクトを保存する
python3 k-core.py --remember " リーセンシーブーストは最近の事実を 10 倍高く重み付けします " --domain " 機能 "
# 関連する事実を取得する
python3 k-core.py --recall「最新性ブースト」
# ドメインを要約する
python3 k-core.py --summary --domain " 機能 "
# すべてのドメインを要約する
python3 k-core.py --summarize-all
# ナレッジグラフをマークダウンとしてエクスポート
python3 k-core.py --export
# 人魚図としてエクスポート (黒曜石対応)
python3 k-core.py --export --mermaid
特長
特徴
何をするのか
リーセンシーブースト
最近の事実の重みは 10 倍高くなります。半減期：90日。
自動要約
TF-IDF + トレンド検出による構造化されたドメインの概要。 LLMは必要ありません。
重複排除
Jaccard と SequenceMatcher の融合。重複する事実はありません。
マークダウンのエクスポート
人間が判読できる .md としての完全なナレッジ グラフ
マーメイドをエクスポート
Obsidian/Notion のインタラクティブなグラフ図
ポータブル
単一ファイル、依存関係なし、どこでも動作
コマンド
コマンド
効果
--<テキスト> を覚えておいてください
タイムスタンプ、ドメイン、重要性を含むファクトを保存する
--recall <クエリ>
関連する事実を取得します (関連性によって並べ替えられます)

× 最近）
--概要 [--ドメイン X]
ドメインの構造化された概要
--すべてを要約する
すべてのドメインの概要
--エクスポート [--マーメイド]
グラフを Markdown または Mermaid としてエクスポート
--バージョン
バージョンを表示
建築
~/k-メモリ/
§── k-core.py # メモリエンジン (v2.1)
§── k-detector.py # 環境自動検出
§── install.sh # ワンコマンドインストーラー
§── ライセンス # MIT
§── テスト/
│ └── test_core.py # 30 テスト、純粋な stdlib
§──graph.json # ナレッジグラフ(ノード+エッジ)
§──index.md # 読み取り可能なインデックス
§── Brain/ # 個々の .md ローブ ファイル
§── summary/ # 自動生成されたドメインの概要
§── extras/ # オプションのプラグイン
│ └── k-embeddings.py # セマンティック検索 (Ollama)
§── エクスポート/ 生成されたエクスポートの数
└── 知識/ # 詳しい知識（任意）
テスト
python3 testing/test_core.py # 30 テスト、外部依存関係なし
エクストラ
高度な機能で K-Memory を拡張するオプションのプラグイン。コアとは異なり、外部依存関係が必要です。
pipインストールリクエスト
ollama pull nomic-embed-text # 274 MB、ローカル、無料
python3 extras/k-embeddings.py --recall "コンセプト"
パフォーマンス
外部依存関係なし (純粋な Python stdlib)
ポータブル: Ubuntu、Debian、macOS、WSL、Termux
速度を低下させることなく 10,000 以上のノードを処理
汎用ハードウェアでの各操作は 100ms 未満
K-Memory は、単純な観察から生まれました。AI エージェントの現在のメモリ システムは、クラウド ベクトル データベースか肥大化した依存関係に依存しているということです。 K-Memory はその逆で、成長を拒否します。 1 つのファイル、1 つのデータ形式、1 つのコミット、1 つの Python3 コマンド。それがすべてになろうとするわけではありません。それで十分になろうとするのです。
MIT — Copyright (c) 2026 KensaiArt. 「ライセンス」を参照してください。
KensaiArt — 建築とデザイン ⚔️ 日々強くなる。
K-Memory — 永続的な自己インストール型メモリ

AI エージェントのメモリ。依存関係ゼロ。純粋なパイソン。
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

K-Memory — Persistent, self-installing memory for AI agents. Zero dependencies. Pure Python. - mackopofa/k-memory

GitHub - mackopofa/k-memory: K-Memory — Persistent, self-installing memory for AI agents. Zero dependencies. Pure Python. · GitHub
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
mackopofa
/
k-memory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits extras extras tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh k-core.py k-core.py k-detector.py k-detector.py View all files Repository files navigation
Persistent memory for AI agents. Zero dependencies. Pure Python.
A self-contained memory engine for LLMs and agents. Store, link, summarize and export knowledge using only Python stdlib. No vectors, no cloud, no lock-in.
bash <( curl -fsSL https://raw.githubusercontent.com/mackopofa/k-memory/main/install.sh )
Or manually:
git clone https://github.com/mackopofa/k-memory.git ~ /k-memory
cd ~ /k-memory && python3 k-core.py
Quick Start
# Store a fact
python3 k-core.py --remember " Recency boost weights recent facts 10x higher " --domain " features "
# Retrieve relevant facts
python3 k-core.py --recall " recency boost "
# Summarize a domain
python3 k-core.py --summary --domain " features "
# Summarize all domains
python3 k-core.py --summarize-all
# Export knowledge graph as Markdown
python3 k-core.py --export
# Export as Mermaid diagram (Obsidian-ready)
python3 k-core.py --export --mermaid
Features
Feature
What it does
Recency boost
Recent facts weighted 10x higher. Half-life: 90 days.
Auto-summary
Structured domain summaries with TF-IDF + trend detection. No LLM needed.
Deduplication
Jaccard + SequenceMatcher fusion. No duplicate facts.
Export Markdown
Full knowledge graph as human-readable .md
Export Mermaid
Interactive graph diagram for Obsidian/Notion
Portable
Single file, zero dependencies, works everywhere
Commands
Command
Effect
--remember <text>
Store a fact with timestamp, domain, importance
--recall <query>
Retrieve relevant facts (sorted by relevance × recency)
--summary [--domain X]
Structured summary of a domain
--summarize-all
Summary of all domains
--export [--mermaid]
Export graph as Markdown or Mermaid
--version
Show version
Architecture
~/k-memory/
├── k-core.py # Memory engine (v2.1)
├── k-detector.py # Environment auto-detector
├── install.sh # One-command installer
├── LICENSE # MIT
├── tests/
│ └── test_core.py # 30 tests, pure stdlib
├── graph.json # Knowledge graph (nodes + edges)
├── index.md # Readable index
├── brain/ # Individual .md lobe files
├── summaries/ # Auto-generated domain summaries
├── extras/ # Optional plugins
│ └── k-embeddings.py # Semantic search (Ollama)
├── exports/ # Generated exports
└── knowledge/ # Detailed knowledge (optional)
Tests
python3 tests/test_core.py # 30 tests, zero external dependencies
Extras
Optional plugins that extend K-Memory with advanced capabilities. They require external dependencies — unlike the core.
pip install requests
ollama pull nomic-embed-text # 274 MB, local, free
python3 extras/k-embeddings.py --recall " concept "
Performance
Zero external dependencies (pure Python stdlib)
Portable: Ubuntu, Debian, macOS, WSL, Termux
Handles 10,000+ nodes without slowdown
Each operation < 100ms on commodity hardware
K-Memory was born from a simple observation: current memory systems for AI agents either depend on cloud vector databases or bloat dependencies. K-Memory is the opposite — it refuses to grow. One file, one data format, one commit, one python3 command. It doesn't try to be everything. It tries to be enough.
MIT — Copyright (c) 2026 KensaiArt. See LICENSE .
KensaiArt — Architecture & Design ⚔️ Stronger every day.
K-Memory — Persistent, self-installing memory for AI agents. Zero dependencies. Pure Python.
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
