---
source: "https://github.com/Rafaelpta/dupehound"
hn_url: "https://news.ycombinator.com/item?id=48490401"
title: "Show HN: Dupehound – find the code your agent wrote twice (no AI required)"
article_title: "GitHub - Rafaelpta/dupehound: Finds the code your AI wrote twice. Fast, offline duplicate-code detector: scan, history chart, CI gate. No AI required. · GitHub"
author: "rafaepta"
captured_at: "2026-06-11T16:29:37Z"
capture_tool: "hn-digest"
hn_id: 48490401
score: 2
comments: 0
posted_at: "2026-06-11T13:55:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Dupehound – find the code your agent wrote twice (no AI required)

- HN: [48490401](https://news.ycombinator.com/item?id=48490401)
- Source: [github.com](https://github.com/Rafaelpta/dupehound)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T13:55:50Z

## Translation

タイトル: HN を表示: Dupehound – エージェントが 2 回書いたコードを見つけます (AI は必要ありません)
記事タイトル: GitHub - Rafaelpta/dupehound: AI が書いたコードを 2 回見つけます。高速なオフライン重複コード検出機能: スキャン、履歴チャート、CI ゲート。 AIは必要ありません。 · GitHub
説明: AI が 2 回書いたコードを検索します。高速なオフライン重複コード検出機能: スキャン、履歴チャート、CI ゲート。 AIは必要ありません。 - ラファエルプタ/デュープハウンド

記事本文:
GitHub - Rafaelpta/dupehound: AI が 2 回書いたコードを見つけます。高速なオフライン重複コード検出機能: スキャン、履歴チャート、CI ゲート。 AIは必要ありません。 · GitHub
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
ラファエルプタ
/
デュープハウンド
公共
通知
通知を変更するにはサインインする必要があります

アイケーション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .github .github 資産 アセット ドキュメント ドキュメント スクリプト スクリプト src src テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
dupehound は、エージェントがコードの大部分を記述するコードベース用に構築された重複コード検出器です。コードのテキストではなく構造のフィンガープリントを取得するため、すべての識別子とリテラルの名前が変更された後でも、複数回存在する関数が検索されます。
dupehound は 3 つのコマンドを実行します。スキャンはすべての重複クラスターとリポレベルのスロップ スコアをレポートします。履歴は git ログ全体での重複をグラフ化し、いつ開始されたかを特定します。変更により既存のコードが重複した場合にチェックが失敗した CI をチェックし、元のコードに再利用する名前を付けます。すべてがローカルで決定論的に実行されます。ネットワークも API キーも機械学習もありません。
macOS、Linux、および Windows 用のビルド済みバイナリは、リリース ページ、または次のページにあります。
カーゴインストールデュープハウンド
履歴を確認し、PATH に git が必要であることを確認します。スキャンはどのディレクトリでも機能します。
dupehound scan [path] は、削除可能な行ごとに重複クラスターをランク付けします。
$デュープハウンドスキャン。
dupehound v0.1.0 — 21 ミリ秒で 19 ファイル、370 行、27 関数をスキャン
╭─────────────────────────╮
│ スロップスコア 36.1% グレード F │
│ 352 の重要な行のうち 127 行は削除可能な重複です │
╰─────────────────

───────╯
● クラスター 1 ─ 4 コピー · 100% 類似 · 42 行の削除可能 ───────
★ src/utils/date.ts:1 formatDate 14行
src/api/timestamps.ts:1 renderTimestamp 14 行 100% █████████
src/jobs/report_dates.ts:1 stringifyDate 14 行 100% █████████
src/billing/dates.ts:1 humanizeDate 14 行 100% █████████
★ = 代表 (維持) · dupehound scan --explain 1 はコードを示します
スロップ スコアは、すべてのクラスターがコピーを 1 つだけ保持した場合に削除できるコードの割合です。テーブル駆動テストは設計上反復的であるため、最大のコピーは除外され、テスト ファイルはデフォルトで除外されます。 --explain N はクラスターのコードを証拠として出力し、--json はバージョン管理されたスキーマを出力し、--card はスコア カードを SVG および PNG として書き込みます。言語: TypeScript、TSX、JavaScript、Python、Rust、Go、Java。
dupehound 履歴は、毎月のスナップショットでスロップ スコアを測定し、オブジェクト データベースから BLOB を直接読み取り (チェックアウトなし)、複製が開始された時点をレポートします。
36.1% ┤ ██
┤ ▂▂▆▆██
┤ ▂▂████████
┤ ████████████
0.0% ┤ ██████████████
━━━━━━━━━
2025-01 2025-12
現在のスロップスコア: 36.1% (グレード F)
2025 年 5 月以降、重複は ~0 から 36.1% に上昇しました
dupehound は CI とプリコミットをチェックします。コードベースの基本リビジョンにインデックスを付け、変更によって追加または変更された関数のみを調査します。移動された関数とインプレース編集は起動しません。終了コード: 0 クリーン、1 検出、2 エラー。
$ dupehound check --diff main 。
src/api/orders.ts:1 CalculateOrderAmount() は src/billing/invoice.ts:1 computeInvoic の 100% 複製です。

eTotal() — 再利用する
GitHub Actions のレシピとコミット前のセットアップは docs/ci.md にあります。コーディング エージェントにコードを書き換えるのではなく再利用させるには、 CLAUDE.md または AGENTS.md からそのコードにフィードバック チェックを戻します。スニペットもそこにあります。
関数本体はツリーシッターで解析され、正規化されます。識別子、文字列、数値はセンチネルになり、コメントは削除され、構造は残ります。 10 個のトークンの k グラムはローリング ハッシュされ、堅牢な選別 (Schleimer、Wilkerson & Aiken、SIGMOD 2003) によって選択されます。これにより、17 個の正規化されたトークンの共有実行が確実に捕捉されます。反転されたフィンガープリント インデックスにより候補ペアが生成され、ボイラープレート フィンガープリントが選別され、類似性は正確な Jaccard で、union-find によってクラスターが構築されます。
デフォルトは誤検知に関して保守的です。生成、縮小、ベンダー化されたファイルはスキップされ、正規化トークン 40 未満の関数は無視され、すべての一致は --explain で検証可能です。グレード バケットは、express (0.0%)、gin (0.2%)、tokio (1.1%)、fastapi (1.7%)、および vscode (2.8%) に対して調整され、すべてグレード A でした。vscode では、297 万行と 53,000 関数で、ラップトップ上で 3.6 秒でスキャンできます。完全な設計ノートは docs/design.md にあります。
コーディング エージェントはコードベースに何が含まれているかを知らないため、コードベースを再実装します。 formatDate は renderTimestamp になり、次に stringifyDate になります。複数の名前で同じロジックが作成され、各コピーは個別にエージングされます。数百万件のコミットを分析したところ、AIアシスタントが主流になって以来、重複が約2倍になったと報告されている。
LLM にはこの仕事はできません。重複検出では、すべての関数を他の関数と比較します。モデルはコンテキストに適合するものをサンプリングし、インデックスはすべてをチェックします。マージ ゲートは再現可能である必要があります。同じ入力、同じ判定、読み取り可能なアルゴリズムです。 dupehound はループの決定論的な側です。エージェントが書き込み、インデックスが記憶します。

問題トラッカーに問題を報告してください。最も有用な誤検知レポートは、一致するはずの小さなコード ペアと --explain 出力です。これらは直接回帰フィクスチャになります。
PR の方は大歓迎です。言語の追加は最も望まれている貢献であり、およそ 1 つのツリー シッター クエリ ファイルです。 COTRIBUTING.md を参照してください。
マサチューセッツ工科大学バンドルされた JetBrains Mono サブセットは SIL OFL 1.1 の下にあります。この図では Excalidraw の Virgil フォント (OFL) を使用しています。
AI が 2 回書いたコードを見つけます。高速なオフライン重複コード検出機能: スキャン、履歴チャート、CI ゲート。 AIは必要ありません。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0
最新の
2026 年 6 月 11 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Finds the code your AI wrote twice. Fast, offline duplicate-code detector: scan, history chart, CI gate. No AI required. - Rafaelpta/dupehound

GitHub - Rafaelpta/dupehound: Finds the code your AI wrote twice. Fast, offline duplicate-code detector: scan, history chart, CI gate. No AI required. · GitHub
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
Rafaelpta
/
dupehound
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .github .github assets assets docs docs scripts scripts src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
dupehound is a duplicate-code detector built for codebases where agents write most of the code. It finds functions that exist more than once, even after every identifier and literal has been renamed, because it fingerprints the structure of the code instead of its text.
dupehound runs three commands: scan reports every duplicate cluster and a repo-level slop score, history charts duplication across the git log and pinpoints when it took off, and check fails CI when a change duplicates code that already exists, naming the original to reuse. Everything runs locally and deterministically: no network, no API keys, no machine learning.
Prebuilt binaries for macOS, Linux and Windows are on the releases page , or:
cargo install dupehound
history and check require git on PATH. scan works on any directory.
dupehound scan [path] ranks duplicate clusters by deletable lines:
$ dupehound scan .
dupehound v0.1.0 — scanned 19 files · 370 lines · 27 functions in 21ms
╭─────────────────────────────────────────────────────────╮
│ SLOP SCORE 36.1% grade F │
│ 127 of 352 significant lines are deletable duplicates │
╰─────────────────────────────────────────────────────────╯
● Cluster 1 ─ 4 copies · 100% similar · 42 deletable lines ─────────────
★ src/utils/date.ts:1 formatDate 14 lines
src/api/timestamps.ts:1 renderTimestamp 14 lines 100% █████████
src/jobs/report_dates.ts:1 stringifyDate 14 lines 100% █████████
src/billing/dates.ts:1 humanizeDate 14 lines 100% █████████
★ = representative (kept) · dupehound scan --explain 1 shows the code
The slop score is the percentage of code you could delete if every cluster kept only one copy; the largest copy is exempt and test files are excluded by default, since table-driven tests are repetitive by design. --explain N prints a cluster's code as proof, --json emits a versioned schema, --card writes a score card as SVG and PNG. Languages: TypeScript, TSX, JavaScript, Python, Rust, Go, Java.
dupehound history measures the slop score at monthly snapshots, reading blobs straight from the object database (no checkouts), and reports when duplication took off:
36.1% ┤ ██
┤ ▂▂▆▆██
┤ ▂▂████████
┤ ▁▁████████████
0.0% ┤ ██████████████
└────────────────────────
2025-01 2025-12
current slop score: 36.1% (grade F)
duplication went from ~0 to 36.1% since 2025-05
dupehound check gates CI and pre-commit. It indexes the codebase at the base revision and probes only the functions a change adds or touches. Moved functions and in-place edits don't fire. Exit codes: 0 clean, 1 findings, 2 error.
$ dupehound check --diff main .
src/api/orders.ts:1 calculateOrderAmount() is a 100% duplicate of src/billing/invoice.ts:1 computeInvoiceTotal() — reuse it
A GitHub Actions recipe and a pre-commit setup are in docs/ci.md . To make a coding agent reuse code instead of rewriting it, feed check back to it from CLAUDE.md or AGENTS.md ; the snippet is there too.
Function bodies are parsed with tree-sitter and normalized: identifiers, strings and numbers become sentinels, comments are dropped, structure stays. k-grams of 10 tokens are rolling-hashed and selected by robust winnowing ( Schleimer, Wilkerson & Aiken, SIGMOD 2003 ), which guarantees any shared run of 17 normalized tokens is caught. An inverted fingerprint index generates candidate pairs, boilerplate fingerprints are culled, similarity is exact Jaccard, union-find builds the clusters.
The defaults are conservative about false positives: generated, minified and vendored files are skipped, functions under 40 normalized tokens are ignored, and every match is verifiable with --explain . Grade buckets were calibrated against express (0.0%), gin (0.2%), tokio (1.1%), fastapi (1.7%) and vscode (2.8%), all grade A. vscode, at 2.97M lines and 53k functions, scans in 3.6s on a laptop. Full design notes in docs/design.md .
Coding agents don't know what a codebase already contains, so they re-implement it. formatDate becomes renderTimestamp , then stringifyDate : the same logic under several names, each copy aging independently. Analyses of millions of commits report duplication roughly doubled since AI assistants went mainstream.
An LLM can't do this job. Duplicate detection compares every function against every other; a model samples what fits in context, an index checks everything. A merge gate must be reproducible: same input, same verdict, an algorithm you can read. dupehound is the deterministic side of the loop: the agent writes, the index remembers.
Please file issues on the issue tracker . The most useful false-positive report is a small code pair that matches but shouldn't, plus the --explain output; these become regression fixtures directly.
PRs welcome. Adding a language is the most wanted contribution and is roughly one tree-sitter query file; see CONTRIBUTING.md .
MIT . Bundled JetBrains Mono subsets are under the SIL OFL 1.1 . The diagram uses Excalidraw's Virgil font (OFL).
Finds the code your AI wrote twice. Fast, offline duplicate-code detector: scan, history chart, CI gate. No AI required.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.0
Latest
Jun 11, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
