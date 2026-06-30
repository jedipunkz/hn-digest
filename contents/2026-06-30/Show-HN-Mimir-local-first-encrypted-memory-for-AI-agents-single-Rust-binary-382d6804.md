---
source: "https://github.com/Perseus-Computing-LLC/mimir"
hn_url: "https://news.ycombinator.com/item?id=48739468"
title: "Show HN: Mimir – local-first encrypted memory for AI agents (single Rust binary)"
article_title: "GitHub - Perseus-Computing-LLC/mimir: Persistent memory MCP server for AI agents. SQLite, FTS5 and vector search, AES-256-GCM, 43 tools. Local-first, single Rust binary. · GitHub"
author: "perseusai"
captured_at: "2026-06-30T21:33:54Z"
capture_tool: "hn-digest"
hn_id: 48739468
score: 1
comments: 0
posted_at: "2026-06-30T21:29:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Mimir – local-first encrypted memory for AI agents (single Rust binary)

- HN: [48739468](https://news.ycombinator.com/item?id=48739468)
- Source: [github.com](https://github.com/Perseus-Computing-LLC/mimir)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T21:29:21Z

## Translation

タイトル: Show HN: Mimir – AI エージェント用のローカルファースト暗号化メモリ (単一の Rust バイナリ)
記事のタイトル: GitHub - Perseus-Computing-LLC/mimir: AI エージェント用の永続メモリ MCP サーバー。 SQLite、FTS5 およびベクトル検索、AES-256-GCM、43 ツール。ローカルファーストの単一 Rust バイナリ。 · GitHub
説明: AI エージェント用の永続メモリ MCP サーバー。 SQLite、FTS5 およびベクトル検索、AES-256-GCM、43 ツール。ローカルファーストの単一 Rust バイナリ。 - Perseus-Computing-LLC/mimir

記事本文:
GitHub - Perseus-Computing-LLC/mimir: AI エージェント用の永続メモリ MCP サーバー。 SQLite、FTS5 およびベクトル検索、AES-256-GCM、43 ツール。ローカルファーストの単一 Rust バイナリ。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このpをリロードしてください

年 。
ペルセウス コンピューティング LLC
/
ミミル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
205 コミット 205 コミット .github .github .hermes/ プラン .hermes/ プラン .well-known/ mcp .well-known/ mcp アセット アセット ベンチマーク ベンチマーク ベンチマーク ベンチマーク clawhub/ スキル/ mimir clawhub/ スキル/ mimir docs docs 例 例 統合 統合プラグイン/ obsidian-mimir plugins/ obsidian-mimir proto/ mimir/ v1 proto/ mimir/ v1 スクリプト スクリプト src src .gitignore .gitignore .nojekyll .nojekyll CHANGELOG.md CHANGELOG.md CLAIMS-AUDIT.md CLAIMS-AUDIT.md CNAME CNAME CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SBOM.md SBOM.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.mdweawe-mimir.mdweawe-mimir.md build.rs build.rsglama.jsonglama.jsonindex.htmlindex.htmlmanifest.jsonmanifest.jsonserver.jsonserver.json smithery.yaml smithery.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用の永続メモリ — MCP ネイティブ。ローカルファースト。依存関係ゼロ。
Mimir は、セッション全体にわたって AI エージェントに耐久性のあるメモリを提供する単一の Rust バイナリです。
バイナリが 1 つ。 1つのファイル。ドッカーはありません。 Postgresはありません。雲はありません。永続的な記憶だけ
あらゆる MCP ホストで動作します。
カール -sSf https://raw.githubusercontent.com/Perseus-Computing-LLC/mimir/main/scripts/install.sh |しー
それだけです。 Mimir は ~/.local/bin/mimir にインストールされます。始めましょう:
mimirserve --db ~ /.mimir/data/mimir.db
macOSのメモ。 Apple Silicon では、署名されていないバイナリは起動時に強制終了されます
OS バイナリ ポリシーによる ( Killed: 9 、出力なし) — 隔離がない場合でも
ある

トリビュート。インストーラーは、Mimir にアドホック コード署名します。ビルドまたはコピーする場合
バイナリを自分で作成し (cargo build --release && cp target/release/mimir ~/.cargo/bin/ )、リビルドするたびに 1 回署名します。
codesign --sign - " $( command -v mimir ) "
任意の MCP ホスト (Claude Desktop、Cursor、Hermes Agent、Perseus など) に接続します。
{
"mcpサーバー": {
"ミミル" : {
"コマンド" : "ミミル" ,
"args" : [ "serve " 、 " --db " 、 " ~/.mimir/data/mimir.db " ]
}
}
}
30 秒のクイックスタート
# ミーミルを起動する
mimirserve --db メモリ.db &
睡眠1
# 事実を覚えておいてください (標準入出力上の MCP JSON-RPC 経由)
echo ' {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"mimir_remember","arguments":{"category":"demo","key":"hello","body_json":"{\"text\":\"Mimir からこんにちは!\"}"}}} ' | mimirserve --db メモリ.db
# 検索してください
echo ' {"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"mimir_recall","arguments":{"query":"Hello"}}} ' | mimirserve --db メモリ.db
すべての MCP クライアントで動作します
Mimir は標準の MCP stdio サーバーです - 同じ mimirserve コマンドが機能します
どこでも。 mimir Doctor を実行してインストールを検証し、このマトリックスをローカルに印刷します。
それぞれの構成スニペットをコピーして貼り付けます: docs/clients/ 。
Mimir は、同時に MCP ネイティブである唯一のメモリ エンジンです。
ローカルファースト、ゼロ依存関係、かつエージェントファースト。
ミーミル
メモリ0
レッタ
ゼップ
導入
単一バイナリ (~8MB)
クラウド + セルフホスト
Docker/Postgres
Docker/Postgres
依存関係
なし (SQLite 埋め込み)
Python + ベクター DB
Postgres + Python
Postgres + Go
MCPネイティブ
✅ 46 ツール
❌ MCP ネイティブではない
❌ MCP ネイティブではない
❌ MCP ネイティブではない
オフライン/ローカル
✅ 完全にローカル
クラウド依存
Docker が必要です
Docker が必要です
暗号化
AES-256-GCM ✅
❌
❌
❌
ハイブリッド検索
BM25 + 高密度 + RRF
ベクターのみ
ベクターのみ
ベクトル + グラフ
エンティティのライフサイクル
衰退 + 促進 + アーカイブ
❌
❌

❌
エンティティグラフ
リンク+トラバース
❌
❌
✅
ジャーナル監査証跡
✅ 不変
❌
❌
❌
状態管理
✅ Key-Value + TTL
❌
❌
❌
MCPツール
46
5
8
0
GitHub スター
~20
～55,000
～15,000
～3K
ライセンス
マサチューセッツ工科大学
アパッチ2.0
アパッチ2.0
アパッチ2.0
完全な比較: Mimir vs Mem0 →
vsレッタ →
vsゼップ →
Mimir は、控えめなハードウェアで実稼働ワークロードを処理します。
自分で実行してください: カーゴテストストレス_100k --release -- --ignored --nocapture
Mimir をデフォルトのメモリ バックエンドにする、すぐに使えるアダプター
一般的な AI エージェント フレームワーク:
MCP stdio サブプロセス経由で接続します (永続セッション)
フレームワークのメモリ インターフェイスを Mimir ツールにマップします
README クイックスタートが付属しています (作業まで 5 分)
モック化された MCP トランスポートによるテストに合格しました
MCP 互換のフレームワークはすべて Mimir と直接連携します。参照
完全なリストの素晴らしいミーミル。
ツール
説明
mimir_remember
エンティティを保存/更新します。 (カテゴリ、キー) による冪等。コンテンツの変更により、以前のバージョンが履歴にスナップショットされます。
mimir_recall
FTS5/デンス/ハイブリッド モード、フィルター、ステミング拡張を使用して検索します。
mimir_recall_layer
特定の生体模倣層 (世界、エピソード、セマンティック) から思い出します。
mimir_recall_when
プロアクティブなジャストインタイム リコール:recall_when トリガーが一致するエンティティを表示します。
mimir_get_entity
完全な body_json を使用して ID によって 1 つのエンティティを取得します。
mimir_as_of
バイテンポラル タイムトラベル: 過去の瞬間に生きていた事実 (カテゴリ + キー) のバージョン。
mimir_history
ファクトの置き換えられたすべてのバージョン (カテゴリ + キー) を最新のものから順にリストします (完全バージョンの証跡 ( mimir_as_of の対))。
mimir_forget
論理的な削除 (アーカイブ = 1)。
検索とラグ
ツール
説明
mimir_ask
RAG: コンテキストを呼び出し、LLM をクエリし、ソースを含む根拠のある回答を返します。
mimir_embed
バンドルされたモデル、Ollama、または OpenAI 互換エンドポイントを介して密なベクトルを生成します。
mimir_seman

ティックサーチ
密のみのセマンティック検索ショートカット — 埋め込まれた類似性のみによってランク付けされ、意味によってエンティティを検索します (キーワードのフォールバックはありません)。
mimir_context
セッションインジェクション用の事前にフォーマットされたマークダウンブロック。
mimir_ingest
コネクタ同期をトリガーします (GitHub、ファイル ウォッチャー)。
mimir_ingest_file
ドキュメントのテキスト (常にプレーンテキスト/マークダウン、マルチモーダル機能を備えた DOCX/PDF) をローカルで抽出し、呼び出し可能なエンティティとして保存します。
mimir_extract
テキストまたは保存されたエンティティからのローカルで決定論的なルールベースの知識抽出 (事実 / 好み / 一時的なイベント / エピソード)。読み取り専用。
グラフ
ツール
説明
mimir_link
エンティティ間の型付き関係リンクを作成します。
mimir_unlink
エンティティ リンクを削除します。
mimir_traverse
エンティティ リンク グラフを構成可能な深さまでウォークします。
日記
ツール
説明
mimir_journal
アクターの属性を含む構造化イベントを追加します。
mimir_timeline
フィルターを使用して時間範囲ごとにジャーナルをクエリします。
州
ツール
説明
mimir_state_set
オプションの TTL を使用してキーと値の状態を設定します。
mimir_state_get
状態値を取得します。期限切れの場合は null を返します。
mimir_state_delete
状態エントリを削除します。
mimir_state_list
状態キーをリストします。オプションでプレフィックスでフィルター処理されます。
ライフサイクル
ツール
説明
mimir_decay
エビングハウス減衰スコアを再計算します (バッチ処理された 1000 エンティティ トランザクション)。
mimir_prune
カテゴリ、減衰しきい値、または経過時間別に一括アーカイブします。
mimir_purge
アーカイブされたエンティティ + VACUUM を完全に削除します。破壊的。
mimir_cohere
自律的なコヒーレンス グルーミング パス - 促進、減衰、リンク、アーカイブ。
mimir_autocohere
完全なアトミック グルーミング: 1 つのパスで凝集 → 減衰 → 圧縮 (ドライランをサポート)。
mimir_compact
エンティティを減衰しきい値以下にアーカイブします。
mimir_reindex
エンティティ テーブルから FTS5 検索インデックスを再構築します。
品質
ツール
説明
mimir_score
品質スコア (0.0 ～ 1.0) を割り当てます。
mimir_conflicts
競合するエンティティを検出する

トリグラムの類似性による。 opt-insolve=true は、確実性の低い側を履歴に無効にします (デフォルトでは、元に戻すことができ、予行演習が行われます)。
mimir_correct
エラーから学ぶための構造化された修正キャプチャ。
mimir_supersede
新しいファクトを古いファクトに置き換えるものとしてマークします (古いエンティティを deprecated に設定します)。
ボールトとフェデレーション
ツール
説明
mimir_vault_export
YAML フロントマターを使用してエンティティを .md ファイルにエクスポートします。
mimir_vault_import
.md ボールト ディレクトリからインポートします (冪等)。
mimir_federate
ワークスペース間でエンティティをコピーします。
mimir_share
コンテンツを保持しながら、1 つのエンティティ (カテゴリ + キー別) を別のワークスペースで共有します。
mimir_workspace_list
すべての個別のエンティティ カテゴリをリストします。
メトリクスと運用
ツール
説明
mimir_stats
すべてのテーブルにわたる完全な DB 統計。
mimir_health
サーバーとDBのヘルスチェック。
mimir_bench
パフォーマンスベンチマークの追跡。
mimir_maintenance
DB メンテナンス: 重複排除、孤立検出、VACUUM、FTS5 再インデックス (ドライランをサポート)。
mimir_synthesize
LLM セッションの合成 — トランスクリプトからレッスンを抽出します。
mimir_移行
v0.1.x DB を現在のスキーマに移行します。
CLI
# サーバー
mimirserve --db /data/mimir.db
mimirserve --web --port 8767 --encryption-key ~ /.mimir/secret.key
mimirserve --llm-endpoint http://localhost:11434/api/generate --llm-model llama3
mimirserve --transport sse --port 8787 --mcp-token my-secret-token
# メンテナンス（DB上で直接操作、サーバー不要）
mimir 統計 --db /data/mimir.db
mimir を忘れます --db /data/mimir.db --category Decision --key stale-choice --reason " 置き換えられました "
mimir prune --db /data/mimir.db --category Junk --min-decay 0.1 --dry-run
mimir パージ --db /data/mimir.db --dry-run
mimir 減衰 --db /data/mimir.db
mimir 再インデックス --db /data/mimir.db
mimir vault-export --db /data/mimir.db --vault-dir ./export/
mimir vault-import --db /data/mimir.db --vault

-dir ./エクスポート/
mimir obsidian-sync ~ /obsidian-vault/Mimir/ # Obsidian Vault へのワンショット エクスポート
mimir obsidian-sync ~ /obsidian-vault/Mimir/ --watch # メモリ変更ごとに継続的に同期
# 鍵の管理
mimir keygen --key-file ~ /.mimir/secret.key
フラグ
旗
説明
--db
SQLite データベース パス (デフォルト: ~/.mimir/data/mimir.db )
--ウェブ
Webダッシュボードを開始する
--ポート
ダッシュボードポート (デフォルト: 8767)
--web-バインド
ダッシュボードのバインド アドレス (デフォルト: 127.0.0.1)
--輸送
MCP トランスポート: stdio (デフォルト)、sse または http
--mcp-トークン
SSE/HTTP トランスポート認証用のベアラー トークン
--暗号化キー
AES-256-GCM キー ファイルのパス
--llm-エンドポイント
mimir_ask および埋め込み用の LLM API エンドポイント
--llm-モデル
LLM モデル名 (デフォルト: llama3)
--llm-API キー
LLM エンドポイント (OpenAI、Azure など) の API キー
--embedding-エンドポイント
OpenAI互換の埋め込みエンドポイント
--connectors-config
Connector.yaml へのパス
Obsidian の AI メモリ
Mimir は AI エージェントの長期記憶であり、第 2 の記憶としても機能します。
脳。エージェントが記憶しているすべてのエンティティは、次のようなプレーンな Markdown ノートにエクスポートされます。
YAML フロントマターにより、AI のメモリがナビゲート可能な個人的な知識になります
すでに使用しているツール (Obsidian、Logseq、または Notion) 内のベース。
# メモリ全体をリンクされた Markdown ノートとして Obsidian Vault にエクスポートします
ミミ

[切り捨てられた]

## Original Extract

Persistent memory MCP server for AI agents. SQLite, FTS5 and vector search, AES-256-GCM, 43 tools. Local-first, single Rust binary. - Perseus-Computing-LLC/mimir

GitHub - Perseus-Computing-LLC/mimir: Persistent memory MCP server for AI agents. SQLite, FTS5 and vector search, AES-256-GCM, 43 tools. Local-first, single Rust binary. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Perseus-Computing-LLC
/
mimir
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
205 Commits 205 Commits .github .github .hermes/ plans .hermes/ plans .well-known/ mcp .well-known/ mcp assets assets benchmark benchmark benchmarks benchmarks clawhub/ skills/ mimir clawhub/ skills/ mimir docs docs examples examples integrations integrations plugins/ obsidian-mimir plugins/ obsidian-mimir proto/ mimir/ v1 proto/ mimir/ v1 scripts scripts src src .gitignore .gitignore .nojekyll .nojekyll CHANGELOG.md CHANGELOG.md CLAIMS-AUDIT.md CLAIMS-AUDIT.md CNAME CNAME CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SBOM.md SBOM.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md awesome-mimir.md awesome-mimir.md build.rs build.rs glama.json glama.json index.html index.html manifest.json manifest.json server.json server.json smithery.yaml smithery.yaml View all files Repository files navigation
Persistent Memory for AI Agents — MCP-Native. Local-First. Zero Dependencies.
Mimir is a single Rust binary that gives AI agents durable memory across sessions.
One binary. One file. No Docker. No Postgres. No cloud. Just persistent memory
that works with any MCP host.
curl -sSf https://raw.githubusercontent.com/Perseus-Computing-LLC/mimir/main/scripts/install.sh | sh
That's it. Mimir is installed to ~/.local/bin/mimir . Start it:
mimir serve --db ~ /.mimir/data/mimir.db
macOS note. On Apple Silicon, an unsigned binary is killed on launch
( Killed: 9 , no output) by the OS binary policy — even with no quarantine
attribute. The installer ad-hoc code-signs Mimir for you. If you build or copy
the binary yourself ( cargo build --release && cp target/release/mimir ~/.cargo/bin/ ), sign it once after each rebuild:
codesign --sign - " $( command -v mimir ) "
Connect any MCP host (Claude Desktop, Cursor, Hermes Agent, Perseus, etc.):
{
"mcpServers" : {
"mimir" : {
"command" : " mimir " ,
"args" : [ " serve " , " --db " , " ~/.mimir/data/mimir.db " ]
}
}
}
30-Second Quickstart
# Start Mimir
mimir serve --db memory.db &
sleep 1
# Remember a fact (via MCP JSON-RPC on stdio)
echo ' {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"mimir_remember","arguments":{"category":"demo","key":"hello","body_json":"{\"text\":\"Hello from Mimir!\"}"}}} ' | mimir serve --db memory.db
# Search for it
echo ' {"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"mimir_recall","arguments":{"query":"Hello"}}} ' | mimir serve --db memory.db
Works With Every MCP Client
Mimir is a standard MCP stdio server — the same mimir serve command works
everywhere. Run mimir doctor to validate your install and print this matrix locally.
Copy-paste config snippets for each: docs/clients/ .
Mimir is the only memory engine that is simultaneously MCP-native,
local-first, zero-dependency, AND agent-first.
Mimir
Mem0
Letta
Zep
Deployment
Single binary (~8MB)
Cloud + self-host
Docker/Postgres
Docker/Postgres
Dependencies
None (SQLite embedded)
Python + vector DB
Postgres + Python
Postgres + Go
MCP-Native
✅ 46 tools
❌ Not MCP-native
❌ Not MCP-native
❌ Not MCP-native
Offline/Local
✅ Fully local
Cloud-dependent
Docker needed
Docker needed
Encryption
AES-256-GCM ✅
❌
❌
❌
Hybrid Search
BM25 + Dense + RRF
Vector only
Vector only
Vector + Graph
Entity Lifecycle
Decay + Promote + Archive
❌
❌
❌
Entity Graph
Link + Traverse
❌
❌
✅
Journal Audit Trail
✅ Immutable
❌
❌
❌
State Management
✅ Key-value + TTL
❌
❌
❌
MCP Tools
46
5
8
0
GitHub Stars
~20
~55K
~15K
~3K
License
MIT
Apache 2.0
Apache 2.0
Apache 2.0
Full comparison: Mimir vs Mem0 →
vs Letta →
vs Zep →
Mimir handles production workloads on modest hardware:
Run it yourself: cargo test stress_100k --release -- --ignored --nocapture
Ready-to-use adapters that make Mimir the default memory backend for
popular AI agent frameworks:
Connects via MCP stdio subprocess (persistent session)
Maps the framework's memory interface to Mimir tools
Comes with a README quickstart (5 minutes to working)
Has passing tests with mocked MCP transport
Any MCP-compatible framework works with Mimir directly. See
Awesome Mimir for the full list.
Tool
Description
mimir_remember
Store/update entity. Idempotent by (category, key); a content change snapshots the prior version into history.
mimir_recall
Search with FTS5/dense/hybrid modes, filters, stemming expansion.
mimir_recall_layer
Recall from a specific biomimetic layer (world, episodic, semantic).
mimir_recall_when
Proactive just-in-time recall: surface entities whose recall_when triggers match.
mimir_get_entity
Fetch one entity by ID with full body_json .
mimir_as_of
Bi-temporal time-travel: the version of a fact (category + key) that was live at a past instant.
mimir_history
List every superseded version of a fact (category + key), newest first — the full version trail (companion to mimir_as_of ).
mimir_forget
Soft-delete (archived=1).
Search & RAG
Tool
Description
mimir_ask
RAG: recall context, query LLM, return grounded answer with sources.
mimir_embed
Generate dense vectors via the bundled model, Ollama, or OpenAI-compatible endpoint.
mimir_semantic_search
Dense-only semantic search shortcut — find entities by meaning, ranked purely by embedding similarity (no keyword fallback).
mimir_context
Pre-formatted markdown block for session injection.
mimir_ingest
Trigger connector syncs (GitHub, file watcher).
mimir_ingest_file
Locally extract a document's text (plaintext/markdown always; DOCX/PDF with the multimodal feature) and store it as a recallable entity.
mimir_extract
Local, deterministic, rule-based knowledge extraction (facts / preferences / temporal events / episodes) from text or a stored entity. Read-only.
Graph
Tool
Description
mimir_link
Create typed relationship links between entities.
mimir_unlink
Remove entity links.
mimir_traverse
Walk entity link graph up to configurable depth.
Journal
Tool
Description
mimir_journal
Append structured event with actor attribution.
mimir_timeline
Query journal by time range with filters.
State
Tool
Description
mimir_state_set
Set key-value state with optional TTL.
mimir_state_get
Get state value. Returns null if expired.
mimir_state_delete
Delete state entry.
mimir_state_list
List state keys, optionally filtered by prefix.
Lifecycle
Tool
Description
mimir_decay
Recalculate Ebbinghaus decay scores (batched 1000-entity transactions).
mimir_prune
Bulk archive by category, decay threshold, or age.
mimir_purge
Permanently delete archived entities + VACUUM. Destructive.
mimir_cohere
Autonomous coherence grooming pass — promote, decay, link, archive.
mimir_autocohere
Full atomic grooming: cohere → decay → compact in one pass (supports dry-run).
mimir_compact
Archive entities below decay threshold.
mimir_reindex
Rebuild FTS5 search index from entities table.
Quality
Tool
Description
mimir_score
Assign quality score (0.0-1.0).
mimir_conflicts
Detect conflicting entities via trigram similarity; opt-in resolve=true invalidates the lower-certainty side into history (reversible, dry-run by default).
mimir_correct
Structured correction capture for learning from errors.
mimir_supersede
Mark a new fact as superseding an old one (sets the old entity to deprecated ).
Vault & Federation
Tool
Description
mimir_vault_export
Export entities to .md files with YAML frontmatter.
mimir_vault_import
Import from .md vault directory (idempotent).
mimir_federate
Copy entities between workspaces.
mimir_share
Share one entity (by category + key) into another workspace, preserving content.
mimir_workspace_list
List all distinct entity categories.
Metrics & Ops
Tool
Description
mimir_stats
Full DB statistics across all tables.
mimir_health
Server and DB health check.
mimir_bench
Performance benchmark tracking.
mimir_maintenance
DB maintenance: dedup, orphan detection, VACUUM, FTS5 reindex (supports dry-run).
mimir_synthesize
LLM session synthesis — extract lessons from transcripts.
mimir_migrate
Migrate v0.1.x DB to current schema.
CLI
# Server
mimir serve --db /data/mimir.db
mimir serve --web --port 8767 --encryption-key ~ /.mimir/secret.key
mimir serve --llm-endpoint http://localhost:11434/api/generate --llm-model llama3
mimir serve --transport sse --port 8787 --mcp-token my-secret-token
# Maintenance (operate directly on DB, no server needed)
mimir stats --db /data/mimir.db
mimir forget --db /data/mimir.db --category decision --key stale-choice --reason " superseded "
mimir prune --db /data/mimir.db --category junk --min-decay 0.1 --dry-run
mimir purge --db /data/mimir.db --dry-run
mimir decay --db /data/mimir.db
mimir reindex --db /data/mimir.db
mimir vault-export --db /data/mimir.db --vault-dir ./export/
mimir vault-import --db /data/mimir.db --vault-dir ./export/
mimir obsidian-sync ~ /obsidian-vault/Mimir/ # one-shot export to an Obsidian vault
mimir obsidian-sync ~ /obsidian-vault/Mimir/ --watch # continuous sync on every memory change
# Key management
mimir keygen --key-file ~ /.mimir/secret.key
Flags
Flag
Description
--db
SQLite database path (default: ~/.mimir/data/mimir.db )
--web
Start web dashboard
--port
Dashboard port (default: 8767)
--web-bind
Dashboard bind address (default: 127.0.0.1)
--transport
MCP transport: stdio (default), sse , or http
--mcp-token
Bearer token for SSE/HTTP transport auth
--encryption-key
AES-256-GCM key file path
--llm-endpoint
LLM API endpoint for mimir_ask and embeddings
--llm-model
LLM model name (default: llama3)
--llm-api-key
API key for LLM endpoints (OpenAI, Azure, etc.)
--embedding-endpoint
OpenAI-compatible embedding endpoint
--connectors-config
Path to connectors.yaml
Your AI Memory in Obsidian
Mimir is your AI agent's long-term memory — and it doubles as your second
brain. Every entity your agent remembers exports to a plain Markdown note with
YAML frontmatter, so your AI's memory becomes a navigable personal knowledge
base inside the tools you already use: Obsidian, Logseq, or Notion.
# Export your entire memory to an Obsidian vault as linked Markdown notes
mimi

[truncated]
