---
source: "https://github.com/l-zhi/pith-wiki"
hn_url: "https://news.ycombinator.com/item?id=48673110"
title: "Show HN: Pith – A local-first desktop LLM wiki without vector DBs or embeddings"
article_title: "GitHub - l-zhi/pith-wiki: Karpathy-style LLM Wiki Desktop App: hydrate raw docs into dense Markdown entries; retrieve by keyword + link traversal. No embeddings, no vector DB. Karpathy llm wiki, pdf-parser, knowledge-base, rag。https://arxiv.org/abs/2605.15184 · GitHub"
author: "a569171010"
captured_at: "2026-06-25T13:42:18Z"
capture_tool: "hn-digest"
hn_id: 48673110
score: 1
comments: 0
posted_at: "2026-06-25T13:32:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pith – A local-first desktop LLM wiki without vector DBs or embeddings

- HN: [48673110](https://news.ycombinator.com/item?id=48673110)
- Source: [github.com](https://github.com/l-zhi/pith-wiki)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T13:32:07Z

## Translation

タイトル: Show HN: Pith – ベクトル DB や埋め込みのないローカルファーストのデスクトップ LLM wiki
記事のタイトル: GitHub - l-zhi/pith-wiki: Karpathy スタイル LLM Wiki デスクトップ アプリ: 生のドキュメントを高密度の Markdown エントリにハイドレートします。キーワード + リンク トラバーサルで取得します。埋め込みもベクトル DB もありません。 Karpathy llm wiki、pdf-parser、Knowledge-base、rag。 https://arxiv.org/abs/2605.15184 · GitHub
説明: Karpathy スタイルの LLM Wiki デスクトップ アプリ: 生のドキュメントを高密度の Markdown エントリに統合します。キーワード + リンク トラバーサルで取得します。埋め込みもベクトル DB もありません。 Karpathy llm wiki、pdf-parser、Knowledge-base、rag。 https://arxiv.org/abs/2605.15184 - l-zhi/pith-wiki

記事本文:
GitHub - l-zhi/pith-wiki: Karpathy スタイルの LLM Wiki デスクトップ アプリ: 生のドキュメントを高密度の Markdown エントリに統合します。キーワード + リンク トラバーサルで取得します。埋め込みもベクトル DB もありません。 Karpathy llm wiki、pdf-parser、Knowledge-base、rag。 https://arxiv.org/abs/2605.15184 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します

n.
アラートを閉じる
{{ メッセージ }}
志
/
ピスウィキ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
129 コミット 129 コミット .github .github bin bin バンドルスキル バンドルスキル デスクトップ デスクトップドキュメント ドキュメントスクリプト スクリプト src src テスト テスト .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CONTEXT.md CONTEXT.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカルファーストの LLM ナレッジベース。メモ、PDF、電子メールをドロップしてください — pith
それぞれを LLM が読み取れる高密度の Markdown エントリに統合し、チャットできるようにします。
キーワード + リンク トラバーサルによって取得されるライブラリ全体。埋め込みはありません、いいえ
ベクトルDB。すべてはディスク上のプレーン ファイルに残ります。
Obsidian の保管庫にメモをドロップします — pith がメモを LLM のエントリに自動的に取り込みます
会話の途中から引き出せる。
🗂️ ブラック ボックスではなくプレーン ファイル — すべてのエントリは Markdown + YAML フロントマターです。 Obsidian、VS Code、Git はネイティブに開きます。
🔍 推論できる検索 — 重み付けされたキーワード検索 + リンクグラフのトラバーサル + 正確な grep。エンベディングドリフトやベンダーロックインはありません。
💬 書き戻すチャット — エージェントはファイル/Wiki ツールを通じてライブラリを読み取ります。 /digest は会話を新しいエントリに抽出します。
🤖 自動取り込み + スケジュール — フォルダーに新しいファイルがないか監視します。 cron 上でエージェント タスクを実行します (例: 毎日のダイジェスト

昨日の追加）。
🔒 ローカルファースト — ~/.pith-wiki/ の下のすべてのデータ。アウトバウンド PII フィルタリングはデフォルトでオンになっています。クラウドもロックインもありません。
設計哲学: データ エンジニアリング > 検索アルゴリズム。生のまま捨てないでください
ドキュメントをストアに保存し、埋め込みによって元に戻されることを願っています。 LLMを使用して水分補給する
各ソースを高密度の Markdown エントリに格納し、キーワード + リンクで取得します
横断。シンプル、ファイルベース、人間が読める形式。
入力形式: .docx .eml .htm .html .md .pdf .txt 。製品は
髄; npm パッケージとリポジトリは、歴史的に pith-wiki という名前が付けられています。
デスクトップ アプリ (推奨) — 完全なエクスペリエンス: チャット、受信トレイ、ダッシュボード、リンク
グラフ、スキル、スケジュールされたタスクのカレンダーをすべて同じエンジン上で実行できます。
オンディスクライブラリ。パッケージ化されたインストーラーがまだないため、ソースから実行します。
git clone https://github.com/l-zhi/pith-wiki.git
cd pith-wiki/デスクトップ
npmインストール
npm run dev # Electron-vite dev (HMR)
CLI (オプション - 自動化/ヘッドレス使用用):
npm install -g pith-wiki
pith-wiki # REPL を起動する
最初の起動時に、オンボーディングによってセットアップが案内されます。プロバイダーを選択し、
API キーを指定し、監視するメモ フォルダーを指定します。すべてはその下で生きている
~/.pith-wiki/ (構成 + Wiki データ);分離された PITH_WIKI_HOME を設定します
プロフィール。
プラットフォーム: macOS + Linux は CI テスト済み (ノード 20 / 22)。 Windowsは使えますが、
まだCIの対象になっていません。開発スクリプト: npm test / npm run typecheck / npm run build
(デスクトップ/内、またはエンジン/コアのリポジトリルート内)。 COTRIBUTING.md を参照してください。
1. ハイドレート — 生のドキュメントを圧縮します (マークダウン / PDF / DOCX / HTML / 電子メール)
Markdown エントリは元のサイズの約 30% に変換されます。フィラーを剥がして保管
信号。 LLM はこれらを直接読み取ります。
2. 取得 — 埋め込みなし、ベクトル DB なし。キーワード重み付け検索（タイトル×2、
タグ × 2、概要 × 1、コンテンツ × 0.5) + BFS リンク トラバーサル、プラス exa

ct
部分文字列/正規表現検索 ( wiki_grep ) および日付範囲フィルター (エントリが
ライブラリに追加された日付、またはコンテンツ自体の日付）。わざと退屈。エントリは次のとおりです。
プレーンなマークダウン。 Obsidian、VS Code、Git はすべてネイティブに開きます。
3. チャット — エージェントはファイル + Wiki ツールを通じてライブラリと対話します。
( wiki_query あいまい検索、wiki_grep 完全検索、wiki_get 、
wiki_read_source 、 wiki_ingest 、 read_file / write_file / list_dir 、…)。
毎ターン記録を書きます。 /digest は会話を抽出して要約します。
Wiki エントリ、ループチャットを終了→保存→取得。
4. 自動取り込み — 監視フォルダー (Obsidian Vault、受信箱など) を中心にポイントします
変更は、バックグラウンド ワーカーがハイドレートできるように自動的にキューに追加されます。
組み込みのヘルスチェックにより、孤立したリンク、壊れたフロントマター、ID の衝突にフラグが立てられます。
5. スケジュール (デスクトップ) — エージェント プロンプトをスケジュールに従って実行するタスクを設定します。
(once または cron) — 例:昨日追加されたすべての日ごとのダイジェスト。それぞれの火事
再度開くことができる新しいセッションを開きます。 ${yyyy-mm-dd -1} 形式の日付プレースホルダー
実行時に解決されるため、「昨日」は常に正しいです。
文書
いつ読むか
docs/config.md
構成フィールド参照、AdditionalReadPaths 、ディスク上のレイアウト
docs/config.example.json
完全な ~/.pith-wiki/config.json の例 (マルチプロバイダー + watchDirs + キュー)
docs/entry-format.md
エントリの YAML フロントマター仕様
docs/architecture.md
3 つのコア サービス + データフロー図
docs/security-model.md
サンドボックスの不変条件 (貢献者必読)
docs/usage.md
CLI リファレンス (高度な / 自動化 - アプリが主な実行方法です)
セキュリティ.md
脆弱性レポート
貢献.md
寄付の流れ
変更ログ.md
バージョン履歴
ライセンス
Apache 2.0 · 著作権 (c) 2026 lizhi
Karpathy スタイルの LLM Wiki デスクトップ アプリ: 生のドキュメントをハイドレートします

高密度の Markdown エントリ。キーワード + リンク トラバーサルで取得します。埋め込みもベクトル DB もありません。 Karpathy llm wiki、pdf-parser、Knowledge-base、rag。 https://arxiv.org/abs/2605.15184
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
3
0.4.0
最新の
2026 年 6 月 23 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Karpathy-style LLM Wiki Desktop App: hydrate raw docs into dense Markdown entries; retrieve by keyword + link traversal. No embeddings, no vector DB. Karpathy llm wiki, pdf-parser, knowledge-base, rag。https://arxiv.org/abs/2605.15184 - l-zhi/pith-wiki

GitHub - l-zhi/pith-wiki: Karpathy-style LLM Wiki Desktop App: hydrate raw docs into dense Markdown entries; retrieve by keyword + link traversal. No embeddings, no vector DB. Karpathy llm wiki, pdf-parser, knowledge-base, rag。https://arxiv.org/abs/2605.15184 · GitHub
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
l-zhi
/
pith-wiki
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
129 Commits 129 Commits .github .github bin bin bundled-skills bundled-skills desktop desktop docs docs scripts scripts src src tests tests .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CONTEXT.md CONTEXT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A local-first LLM knowledge base. Drop in your notes, PDFs, and emails — pith
hydrates each into a dense Markdown entry an LLM can read, then lets you chat with
your whole library, retrieved by keyword + link traversal. No embeddings, no
vector DB. Everything stays in plain files on your disk.
Drop a note into your Obsidian vault — pith auto-ingests it into an entry the LLM
can pull from mid-conversation.
🗂️ Plain files, not a black box — every entry is Markdown + YAML frontmatter; Obsidian, VS Code, and Git open them natively.
🔍 Retrieval you can reason about — weighted keyword search + link-graph traversal + exact grep. No embedding drift, no vendor lock-in.
💬 Chat that writes back — the agent reads your library through file/wiki tools; /digest distills a conversation into a new entry.
🤖 Auto-ingest + schedule — watch a folder for new files; run agent tasks on a cron (e.g. a daily digest of yesterday's additions).
🔒 Local-first — all data under ~/.pith-wiki/ ; outbound PII filtering on by default. No cloud, no lock-in.
Design philosophy: data engineering > retrieval algorithms. Don't dump raw
docs into a store and hope embeddings pull them back. Use an LLM to hydrate
each source into a high-density Markdown entry, then retrieve by keyword + link
traversal. Simple, file-based, human-readable.
Input formats: .docx .eml .htm .html .md .pdf .txt . The product is
pith ; the npm package and repo are historically named pith-wiki .
Desktop app (recommended) — the full experience: chat, inbox, dashboard, link
graph, skills, and a scheduled-tasks calendar, all over the same engine and
on-disk library. No packaged installer yet, so run it from source:
git clone https://github.com/l-zhi/pith-wiki.git
cd pith-wiki/desktop
npm install
npm run dev # electron-vite dev (HMR)
CLI (optional — for automation / headless use) :
npm install -g pith-wiki
pith-wiki # launch the REPL
On first launch, onboarding walks you through setup — pick a provider, paste an
API key, and point it at a notes folder to watch. Everything lives under
~/.pith-wiki/ (config + wiki data); set PITH_WIKI_HOME for an isolated
profile.
Platforms : macOS + Linux are CI-tested (Node 20 / 22); Windows is usable but
not yet CI-covered. Dev scripts: npm test / npm run typecheck / npm run build
(inside desktop/ , or repo root for engine/core). See CONTRIBUTING.md .
1. Hydrate — compress raw documents (markdown / PDF / DOCX / HTML / email)
into Markdown entries roughly 30% of the original size. Strip filler, keep
signal. LLMs read these directly.
2. Retrieve — no embeddings, no vector DB. Weighted keyword search (title × 2,
tags × 2, summary × 1, content × 0.5) + BFS link traversal, plus exact
substring/regex search ( wiki_grep ) and date-range filters (when an entry was
added to the library, or the content's own date). Boring on purpose. Entries are
plain Markdown; Obsidian, VS Code, and Git all open them natively.
3. Chat — the agent talks to your library through file + wiki tools
( wiki_query fuzzy search, wiki_grep exact search, wiki_get ,
wiki_read_source , wiki_ingest , read_file / write_file / list_dir , …).
Every turn writes a transcript; /digest distills the conversation back into a
wiki entry, closing the loop chat → store → retrieve.
4. Auto-ingest — point a watch folder (Obsidian vault, inbox, etc.) at pith
in Settings, and changes are auto-enqueued for a background worker to hydrate.
Built-in health checks flag orphan links, broken frontmatter, and ID collisions.
5. Schedule (desktop) — set tasks that run an agent prompt on a schedule
(once, or cron) — e.g. a daily digest of everything added yesterday. Each fire
opens a fresh session you can reopen; ${yyyy-mm-dd -1} -style date placeholders
are resolved at run time so "yesterday" is always correct.
Document
When to read it
docs/config.md
Configuration field reference, additionalReadPaths , on-disk layout
docs/config.example.json
Full ~/.pith-wiki/config.json example (multi-provider + watchDirs + queue)
docs/entry-format.md
YAML frontmatter spec for entries
docs/architecture.md
Three core services + data-flow diagram
docs/security-model.md
Sandbox invariants (required reading for contributors)
docs/usage.md
CLI reference (advanced / automation — the app is the primary way to run)
SECURITY.md
Vulnerability reporting
CONTRIBUTING.md
Contribution flow
CHANGELOG.md
Version history
License
Apache 2.0 · Copyright (c) 2026 lizhi
Karpathy-style LLM Wiki Desktop App: hydrate raw docs into dense Markdown entries; retrieve by keyword + link traversal. No embeddings, no vector DB. Karpathy llm wiki, pdf-parser, knowledge-base, rag。 https://arxiv.org/abs/2605.15184
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
3
0.4.0
Latest
Jun 23, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
