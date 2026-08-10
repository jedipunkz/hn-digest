---
source: "https://github.com/lo2589/JSONSEEK"
hn_url: "https://news.ycombinator.com/item?id=49246117"
title: "Show HN: Jsonseek – Token-efficient JSON/JSONL navigation for LLM agents"
article_title: "GitHub - lo2589/JSONSEEK: Path-based CLI for searching, inspecting, patching, and debugging JSON/JSONL files. · GitHub"
author: "yoliliya"
captured_at: "2026-08-10T16:43:04Z"
capture_tool: "hn-digest"
hn_id: 49246117
score: 1
comments: 0
posted_at: "2026-08-10T16:38:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Jsonseek – Token-efficient JSON/JSONL navigation for LLM agents

- HN: [49246117](https://news.ycombinator.com/item?id=49246117)
- Source: [github.com](https://github.com/lo2589/JSONSEEK)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T16:38:47Z

## Translation

タイトル: HN を表示: Jsonseek – LLM エージェント向けのトークン効率の高い JSON/JSONL ナビゲーション
記事のタイトル: GitHub - lo2589/JSONSEEK: JSON/JSONL ファイルの検索、検査、パッチ適用、デバッグのためのパスベースの CLI。 · GitHub
説明: JSON/JSONL ファイルを検索、検査、パッチ適用、デバッグするためのパスベースの CLI。 - lo2589/JSONSEEK

記事本文:
GitHub - lo2589/JSONSEEK: JSON/JSONL ファイルの検索、検査、パッチ適用、デバッグのためのパスベースの CLI。 · GitHub
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
lo2589
/
JSONSEEK
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
13 コミット 13 コミット スキル/ jsonseek スキル/ jsonseek src/ jsonseek src/ jsonseek テスト テスト .gitignore .gitignore LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md README_EN。

md README_EN.md README_ZH.md README_ZH.md pyproject.toml pyproject.toml setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM 用に設計された JSON/JSONL 解析ツールキット。
LLM が JSON を扱うときは、最初に整形し、次にクエリを実行する必要があり、ファイル全体を cat することはありません。
人間が JSON に触れるときも、コンテキスト ウィンドウの代わりにキーボードを使用するだけで、同じルールが適用されます。
jsonseek は、LLM に適したパーサーであり、大きな JSON/JSONL ファイル用の部分エディターです。中心原則は、LLM に JSON ファイル全体をコンテキストに取り込ませないことです。スケルトンのシェイプを実行し、正確な検索のためにクエリ/取得を実行し、可能な限り最小限の編集のために設定/追加/削除/追加を行います。
1 つのコマンド セットで、JSON と JSONL の両方について、構造の理解、フィールドの概要、部分的なクエリ、部分的な編集、バグのローカリゼーションと修復をサポートします。
🤖 LLMコーディングエージェントのスキル（クロード/キミ/カーソル/コーデックスなど）
エージェントに「jsonseek を使用してください」と伝えれば、問題なく機能します。スキルは以下に記載されています。オンデマンドでお読みください。
GitHub のスキル: lo2589/jsonseek/skills/jsonseek/
エージェントに接続します (クロード コード / カーソル / コーデックスなど):
git クローン https://github.com/lo2589/JSONSEEK.git
ln -s ../jsonseek/skills/jsonseek/SKILL.md ~ /.claude/skills/jsonseek.md
# または ~/.cursor/skills/ ~/.codex/skills/
JSON に専用ツールが必要な理由
JSON は、最新のデータ交換の事実上の標準です。 ML 実験ログ、API 構成、アプリケーション ログ ストリームから、マイクロサービス レジストリやクローラー ダンプに至るまで、JSON / JSONL はあらゆる場所に存在します。
ML 実験の追跡: トレーニング パラメーター、メトリック カーブ、モデル構成はすべて JSON として保存されます。 1 つの実験ディレクトリは簡単に数十 MB に達することがあります。
API / マイクロサービス構成: サービス検出、ルーティング ルール、環境変数 — 多くの場合、JSON 構成として管理されます。
ログとイベント ストリーム: 構造化ログ (JSONL)

プレーン テキストよりもクエリが簡単ですが、ファイル サイズが急速に大きくなります。
データ交換: フロントエンドとバックエンドの通信、サービス間 RPC、クローラー ダンプ — JSON が最も一般的な形式です。
問題は、JSON が大きくなるほど、処理コストが高くなるということです。 10MB JSON を LLM コンテキストに cat -ing すると、数百万のトークンが消費されます。人間の開発者でさえ、何千ものネストされた行をスキャンするのに苦労しています。
jsonseek はこれを解決します。完全な読み取りを部分的な操作に置き換え、手動スキャンを構造化クエリに置き換えます。 LLM コーディング エージェントや JSON / JSONL を頻繁に扱う開発者にとって、これは不可欠なツールです。
あなた (または実行中のクロード / キミ / カーソル / コーデックス エージェント) が 10MB の JSON に直面した場合、完全な cat をコンテキストに取り込むと、壊滅的なトークンの無駄が発生します。 jsonseek により、エージェントは次のことを行うことができます。
最初に構造を理解します。コンテンツを読まずに、スケルトンの形状、フィールド リストのフィールドを理解します。
次にターゲットを特定します。クエリでキーワードを検索し、ls でレイヤーを参照し、正確な値を取得します。
部分的に最後に編集 — 必要な場合のみ設定 / 追加 / 削除 / 追加
ファイルサイズ
操作
完全に読みました
jsonseek出力
貯蓄
100KBの構成JSON
形
~25,000 トークン
~100トークン
99%以上
100KBの構成JSON
フィールド
~25,000 トークン
~300トークン
98%以上
100KBの構成JSON
単一の値を取得する
~25,000 トークン
~10トークン
99%以上
100KBの構成JSON
クエリがいくつかヒットしました
~25,000 トークン
~100トークン
99%以上
10MB ログ JSONL
サンプリングされた形状
~250万トークン
~200トークン
99.9%以上
10MB ログ JSONL
クエリが数十件ヒットしました
~250万トークン
~1,000 トークン
99.9%以上
概算: 1 トークン ≈ 英語の 4 バイト。実際の比率はコンテンツとトークナイザーによって異なりますが、大きさのオーダーは安定しています。ファイルが大きいほど、節約できる金額も大きくなります。
# 1. スケルトンを読む - コンテンツは必要ありません
jsonseek形状データ.jsonl
# 2. フィールドリストを参照
jsonseek フィールド data.jsonl --top
# 3. キーワードを検索する
jsonseek クエリ data.jsonl パスワード

--レコード ID フィールド ID --max-results 5
# 4. 特定の値を読み取る
jsonseek 取得 data.jsonl ' [3].password '
# 5. 変更 (**常に --dry-run で開始します**)
jsonseek セット data.jsonl ' [3].password ' ' newpass ' --dry-run
jsonseek set data.jsonl ' [3].password ' ' newpass ' --backup
📦 インストール
pip インストール jsonseek
Python 3.8以降が必要です。依存関係なし、構成なし — インストールして実行します。
$ jsonseek --バージョン
jsonseek < current_version >
コマンドリファレンス
読み取り/検査 (安全、読み取り専用)
コマンド
目的
形状ファイル
構造・スケルトンツリーの表示
フィールド ファイル [キーワード]
すべてのフィールドとタイプをリストする
ls ファイル [パス]
パス上の子をリストする
ファイルパスを取得する
パスの値を取得する
FILE キーワードのクエリ
キーまたは値を検索する
PATTERN パスを抽出する
多数の JSON ファイルから同じパスをバッチ抽出する
連結パターン
複数の JSON ファイルを JSONL にマージする
書き込み/部分編集（ファイルの変更）
コマンド
目的
ファイルパス値を設定します
フィールドを変更する
ファイルパス値を追加
新しいキーを追加する
デルファイルパス
キーまたは配列要素を削除する
ファイルパス値を追加
1 つの項目を配列に追加する
ファイルパス json_array を拡張します
JSON配列から配列を拡張する
カットライン ファイル N
特定の JSONL 行を一時ファイルに抽出します
replaceline ファイル N
特定の JSONL 行を置き換えます
共通フラグ
旗
目的
--出力json
機械可読出力 (別のツール/エージェントにパイプする場合に必要)
--バックアップ
書き込みの前に .bak バックアップを作成する
--ドライラン
プレビューに書き込む前に必ずこれを使用してください
--kind {json,jsonl}
強制ファイルタイプ (デフォルトで自動検出)
--encoding エンコーディング
強制エンコード (デフォルトで自動検出; 例: gbk )
--コンテキスト N
ターゲットの周囲のコンテキスト行 (JSONL のみ、デフォルトは 2 )
パスの構文
スタイル
例
意味
ドット
a.b.c
a -> b -> c
ブラケット
a[キー1][キー2]
a -> キー1 -> キー2
混合
a[キー1].b[0]
a -> キー1 -> b -> 0
配列インデックス

×
アイテム[0][1]
項目 -> 0 -> 1
⚠️ zsh ユーザー: zsh は括弧をグロブ展開しようとするため、[N] を含むパスは一重引用符で囲む必要があります (または noglob を使用します)。
# ❌ zsh: "一致するものが見つかりませんでした"
jsonseek del file.json services[0].deprecated
# ✅ これらの作品のいずれか
jsonseek del file.json ' services[0].deprecated '
noglob jsonseek del file.json services[0].deprecated
bash / Fish / zsh-with-quotes はすべて正常に動作します。
3 つの鉄則 (LLM と人間用)
書き込みの前に、常に最初に --dry-run を実行します。
jsonseek は file.json パス値を設定します --dry-run
# [DRY-RUN] 前: path = old
# [DRY-RUN] 後: path = new
書き込みの前に、必ず --backup を追加してください。
jsonseek は file.json パス値を設定します --backup
# → file.json.bak を作成
別のツールにパイプする場合は、常に --output json を追加します。
jsonseek クエリ file.json キーワード --output json | jq ' .hits[0] '
AI / コーディングエージェント向けに設計
jsonseek の設計哲学は、LLM のトークン バジェットです。
できるだけ短く出力します。デフォルトでは、フィルタリングされた必須部分のみが出力され、構造全体は出力されません。
安定した出力形式: すべてのコマンドは --output json をサポートし、エージェントは直接解析します
書き込みはプレビュー可能です。すべての set / add / del には --dry-run があり、エージェントは呼び出す前に差分をプレビューします。
ストリーミング : 10MB JSONL 上の jsonseek シェイプは最初の 100 行のみを読み取ります ( --sample-size 調整可能)、トークンの使用量は制限されています
読み取り → シェイプ / フィールド / ls / クエリ / 取得
↓（構造を理解する）
検索 → クエリ / 取得
↓ (ターゲットを見つける)
write → set / add / del / append (最初に --dry-run → 実際には --backup)
↓（確認）
読み取り→クエリ/取得
クロスプラットフォーム
macOS / Linux : ネイティブ CLI は問題なく動作します
Windows PowerShell: 読み取りコマンド (shape /fields/get/query/ls/extract/concat) は CLI 経由で正常に機能します。書き込みコマンドは PowerShell を通じて二重引用符を削除するため、複雑な値になります

失敗します - 代わりに Python API を使用してください
インポートシステム
システム。パス 。挿入 ( 0 , '.' )
jsonseek から。コマンド。 set_cmd インポート set_value
jsonseek から。コマンド。 add_cmd インポート add_value
jsonseek から。コマンド。 append_cmd インポート append_value
jsonseek から。コマンド。 extend_cmd import extend_value
jsonseek から。コマンド。 del_cmd インポート del_value
jsonseek から。コマンド。 replaceline_cmd インポート replace_line
# 複雑な値を設定/追加/追加/拡張 - シェルの引用の問題はありません
set_value ( 'file.json' , 'path' , { "キー" : "値" })
add_value ( 'file.json' , 'path' , [ "item1" , "item2" ])
append_value ( 'file.json' , 'items' , { "id" : 1 })
extend_value ( 'file.json' , 'items' , [{ "id" : 2 }, { "id" : 3 }])
# 削除
del_value ( 'file.json' , 'path' )
# JSONL全行置換
replace_line ( 'file.jsonl' , 5 , '{"id": 5, "name": "fixed"}' )
CLI 書き込みコマンドは、成功した場合はパッチ プレビューを出力し、失敗した場合はエラー: ... を出力します。 Python API 書き込みヘルパーは、成功時には沈黙し、失敗時には発生します。
すべてのコマンドの完全な署名、すべてのフラグ、すべての出力形式。最新のリストについては、 jsonseek <command> --help を実行してください。
形状ファイル
構造/スケルトンツリーを表示します。
--kind {json,jsonl} ファイルの種類を強制します (デフォルトで自動検出されます)
--output {pretty,json} 出力形式 (デフォルト: pretty)
--encoding ENCODING 強制エンコード (デフォルトで自動検出)
--max- Depth N トラバーサルの深さを制限します (デフォルト: 無制限)
--array-mode {sample,full} JSONL 配列モード。 `sample` はデフォルト、`full` はすべての要素を調べます
--sample-size N JSONL: サンプリングするレコードの数 (デフォルト: 100)
フィールド ファイル [キーワード]
すべてのフィールドをタイプ/出現数とともにリストします。
--top 最上位フィールドのみを表示します
--kind / --output / --encoding (上記と同じ)
ls ファイル [パス]
パス上の子をリストします。 JSONL: パスは [N] で始まる必要があります。または

レコード[N] 。
--kind / --output / --encoding (上記と同じ)
ファイルパスを取得する
パスの値を読み取ります。出力は --output を尊重します。
--kind / --output / --encoding (上記と同じ)
クエリファイル用語
キーまたは値を検索します。
--case-sensitive 大文字と小文字を区別するマッチング (デフォルト: 区別しない)
--exact 完全一致 (デフォルト: 部分文字列)
--match-mode {key,value,both} 一致させるもの (デフォルト: 両方)
--max-results N 結果の数を制限します
--record-id-field FIELD JSONL: 出力のレコード ID として FIELD を使用します。
--preview-field FIELD JSONL: FIELD をプレビューとしても表示します
--kind / --output / --encoding / --context N (上記と同じ)
PATTERN パスを抽出する
glob パターンに一致する多数の JSON ファイルから同じパスをバッチ抽出します。
--include-missing パスが存在しないファイルを含めます (デフォルト: スキップ)
--output {pretty,json} 出力形式 (デフォルト: pretty)
連結パターン
複数の JSON ファイルを 1 つの JSONL に連結します。
-o, --output-file FILE 出力ファイル (デフォルト: stdout)
--no-sort グロブ順序を保持します (デフォルト: ファイル名でソート)
コマンドの書き込み (常に最初に --backup)
ファイルパス値を設定します
パスの既存の値を変更します。
--create-missing 中間パスを自動作成します (デフォルト: 見つからない場合はエラー)
--from-file FILE ファイルから新しい値を読み取ります (シェルの引用の問題を回避します)。
--backup FILE.bak を作成します

[切り捨てられた]

## Original Extract

Path-based CLI for searching, inspecting, patching, and debugging JSON/JSONL files. - lo2589/JSONSEEK

GitHub - lo2589/JSONSEEK: Path-based CLI for searching, inspecting, patching, and debugging JSON/JSONL files. · GitHub
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
lo2589
/
JSONSEEK
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
13 Commits 13 Commits skills/ jsonseek skills/ jsonseek src/ jsonseek src/ jsonseek tests tests .gitignore .gitignore LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md README_EN.md README_EN.md README_ZH.md README_ZH.md pyproject.toml pyproject.toml setup.py setup.py View all files Repository files navigation
JSON/JSONL parsing toolkit, designed for LLMs.
When LLMs touch JSON, they should shape first, query second, never cat the whole file.
When a human touches JSON, the same rules apply — just with a keyboard instead of a context window.
jsonseek is an LLM-friendly parser and partial editor for large JSON / JSONL files . The core principle: never let an LLM cat an entire JSON file into context . Run shape for the skeleton, query / get for precise lookup, then set / add / del / append for the smallest possible edit.
Supports structural understanding, field summaries, partial queries, partial edits, bug localization and repair — for both JSON and JSONL with one command set.
🤖 Skills for LLM Coding Agents (Claude / Kimi / Cursor / Codex, etc.)
Tell your agent "use jsonseek" and it just works. Skills live below — read on demand.
Skills on GitHub : lo2589/jsonseek/skills/jsonseek/
Plug into an agent (Claude Code / Cursor / Codex / etc.):
git clone https://github.com/lo2589/JSONSEEK.git
ln -s ../jsonseek/skills/jsonseek/SKILL.md ~ /.claude/skills/jsonseek.md
# or ~/.cursor/skills/ ~/.codex/skills/
Why JSON Deserves a Dedicated Tool
JSON is the de facto standard for modern data exchange. From ML experiment logs, API configs, application log streams, to microservice registries and crawler dumps — JSON / JSONL is everywhere:
ML experiment tracking : training parameters, metric curves, and model configs all live as JSON. A single experiment directory can easily reach tens of MB.
API / microservice configs : service discovery, routing rules, environment variables — often managed as JSON configs.
Logs & event streams : structured logs (JSONL) are easier to query than plain text, but file size grows fast.
Data exchange : frontend-backend communication, inter-service RPC, crawler dumps — JSON is the most common format.
The problem: the bigger the JSON, the more expensive it is to process. cat -ing a 10MB JSON into LLM context burns millions of tokens. Even human developers suffer scanning thousands of nested lines.
jsonseek solves this — replace full reads with partial operations, replace manual scanning with structured queries. For LLM coding agents and developers handling JSON / JSONL frequently, this is essential tooling.
When you (or your running Claude / Kimi / Cursor / Codex agent) face a 10MB JSON, full cat into context is catastrophic token waste . jsonseek lets the agent:
Understand structure first — shape for the skeleton, fields for the field list, without reading content
Locate targets next — query to search keywords, ls to browse a layer, get for a precise value
Edit partially last — set / add / del / append only where needed
File size
Operation
Full read
jsonseek output
Savings
100KB config JSON
shape
~25K tokens
~100 tokens
99%+
100KB config JSON
fields
~25K tokens
~300 tokens
98%+
100KB config JSON
get single value
~25K tokens
~10 tokens
99%+
100KB config JSON
query hit a few
~25K tokens
~100 tokens
99%+
10MB log JSONL
shape sampled
~2.5M tokens
~200 tokens
99.9%+
10MB log JSONL
query hit dozens
~2.5M tokens
~1K tokens
99.9%+
Rough estimate: 1 token ≈ 4 bytes of English. Actual ratios vary by content and tokenizer, but the order of magnitude is stable — the bigger the file, the bigger the savings .
# 1. Read the skeleton — no content needed
jsonseek shape data.jsonl
# 2. See the field list
jsonseek fields data.jsonl --top
# 3. Search a keyword
jsonseek query data.jsonl password --record-id-field id --max-results 5
# 4. Read a specific value
jsonseek get data.jsonl ' [3].password '
# 5. Modify (**always start with --dry-run**)
jsonseek set data.jsonl ' [3].password ' ' newpass ' --dry-run
jsonseek set data.jsonl ' [3].password ' ' newpass ' --backup
📦 Installation
pip install jsonseek
Requires Python 3.8+. Zero dependencies, zero configuration — install and run.
$ jsonseek --version
jsonseek < current_version >
Command reference
Read / inspect (safe, read-only)
Command
Purpose
shape FILE
Display structure / skeleton tree
fields FILE [keyword]
List all fields and types
ls FILE [path]
List children at a path
get FILE path
Get a value at a path
query FILE keyword
Search keys or values
extract PATTERN path
Batch-extract the same path from many JSON files
concat PATTERN
Merge multiple JSON files into JSONL
Write / partial edits (modifies files)
Command
Purpose
set FILE path value
Modify a field
add FILE path value
Add a new key
del FILE path
Delete a key or array element
append FILE path value
Append one item to an array
extend FILE path json_array
Extend an array from a JSON array
cutline FILE N
Extract a specific JSONL line to a temp file
replaceline FILE N
Replace a specific JSONL line
Common flags
Flag
Purpose
--output json
Machine-readable output (required when piping to another tool / agent)
--backup
Create a .bak backup before any write
--dry-run
Always use this before any write to preview
--kind {json,jsonl}
Force file type (auto-detected by default)
--encoding ENCODING
Force encoding (auto-detected by default; e.g. gbk )
--context N
Lines of context around the target (JSONL only, default 2 )
Path syntax
Style
Example
Meaning
Dot
a.b.c
a -> b -> c
Bracket
a[key1][key2]
a -> key1 -> key2
Mixed
a[key1].b[0]
a -> key1 -> b -> 0
Array index
items[0][1]
items -> 0 -> 1
⚠️ zsh users : paths containing [N] must be wrapped in single quotes (or use noglob ), because zsh tries to glob-expand brackets:
# ❌ zsh: "no matches found"
jsonseek del file.json services[0].deprecated
# ✅ either of these works
jsonseek del file.json ' services[0].deprecated '
noglob jsonseek del file.json services[0].deprecated
bash / fish / zsh-with-quotes all work fine.
The three iron rules (for LLMs and humans)
Before any write, always --dry-run first :
jsonseek set file.json path value --dry-run
# [DRY-RUN] Before: path = old
# [DRY-RUN] After: path = new
Before any write, always add --backup :
jsonseek set file.json path value --backup
# → creates file.json.bak
When piping to another tool, always add --output json :
jsonseek query file.json keyword --output json | jq ' .hits[0] '
Designed for AI / Coding Agents
jsonseek 's design philosophy is the LLM's token budget :
Output as short as possible : by default only the filtered essentials, never the whole structure
Stable output format : every command supports --output json , agents parse directly
Writes are previewable : every set / add / del has --dry-run , agents preview the diff before invoking
Streaming : jsonseek shape on a 10MB JSONL only reads the first 100 lines ( --sample-size adjustable), token usage is bounded
read → shape / fields / ls / query / get
↓ (understand structure)
locate → query / get
↓ (find target)
write → set / add / del / append (--dry-run first → --backup for real)
↓ (verify)
read → query / get
Cross-platform
macOS / Linux : native CLI works flawlessly
Windows PowerShell : read commands ( shape / fields / get / query / ls / extract / concat ) work fine via CLI. Write commands strip double quotes through PowerShell, so complex values fail — use the Python API instead
import sys
sys . path . insert ( 0 , '.' )
from jsonseek . commands . set_cmd import set_value
from jsonseek . commands . add_cmd import add_value
from jsonseek . commands . append_cmd import append_value
from jsonseek . commands . extend_cmd import extend_value
from jsonseek . commands . del_cmd import del_value
from jsonseek . commands . replaceline_cmd import replace_line
# Set/Add/Append/Extend complex values — no shell quoting issues
set_value ( 'file.json' , 'path' , { "key" : "value" })
add_value ( 'file.json' , 'path' , [ "item1" , "item2" ])
append_value ( 'file.json' , 'items' , { "id" : 1 })
extend_value ( 'file.json' , 'items' , [{ "id" : 2 }, { "id" : 3 }])
# Delete
del_value ( 'file.json' , 'path' )
# JSONL whole-line replacement
replace_line ( 'file.jsonl' , 5 , '{"id": 5, "name": "fixed"}' )
CLI write commands print a patch preview on success and Error: ... on failure. Python API write helpers are silent on success and raise on failure.
Every command's full signature, every flag, and every output format. For the most up-to-date list, run jsonseek <command> --help .
shape FILE
Show structure / skeleton tree.
--kind {json,jsonl} Force file kind (auto-detected by default)
--output {pretty,json} Output format (default: pretty)
--encoding ENCODING Force encoding (auto-detected by default)
--max-depth N Limit traversal depth (default: unlimited)
--array-mode {sample,full} JSONL array mode; `sample` is default, `full` walks every element
--sample-size N JSONL: number of records to sample (default: 100)
fields FILE [keyword]
List all fields with type / occurrence count.
--top Show only top-level fields
--kind / --output / --encoding (same as above)
ls FILE [path]
List children at a path. JSONL: path must start with [N] or records[N] .
--kind / --output / --encoding (same as above)
get FILE path
Read a value at a path. Output respects --output .
--kind / --output / --encoding (same as above)
query FILE term
Search keys or values.
--case-sensitive Case-sensitive matching (default: insensitive)
--exact Exact match (default: substring)
--match-mode {key,value,both} What to match (default: both)
--max-results N Limit number of results
--record-id-field FIELD JSONL: use FIELD as record ID in output
--preview-field FIELD JSONL: also show FIELD as preview
--kind / --output / --encoding / --context N (same as above)
extract PATTERN path
Batch-extract the same path from many JSON files matched by a glob pattern.
--include-missing Include files where the path does not exist (default: skip)
--output {pretty,json} Output format (default: pretty)
concat PATTERN
Concatenate multiple JSON files into a single JSONL.
-o, --output-file FILE Output file (default: stdout)
--no-sort Preserve glob order (default: sort by filename)
Write commands (always --backup first)
set FILE path value
Modify an existing value at a path.
--create-missing Auto-create intermediate paths (default: error if missing)
--from-file FILE Read the new value from a file (avoids shell quoting issues)
--backup Create FILE.bak

[truncated]
