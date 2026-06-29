---
source: "https://github.com/matthiasnordwig/pdf-struct-chunker"
hn_url: "https://news.ycombinator.com/item?id=48716707"
title: "LLM-free, layout-aware PDF chunker in pure Rust"
article_title: "GitHub - matthiasnordwig/pdf-struct-chunker: LLM-free, layout-aware PDF chunking for RAG pipelines. Preserves document structure via regex and font heuristics. Written in pure Rust. · GitHub"
author: "MatthiasNordwig"
captured_at: "2026-06-29T09:32:34Z"
capture_tool: "hn-digest"
hn_id: 48716707
score: 1
comments: 0
posted_at: "2026-06-29T09:06:32Z"
tags:
  - hacker-news
  - translated
---

# LLM-free, layout-aware PDF chunker in pure Rust

- HN: [48716707](https://news.ycombinator.com/item?id=48716707)
- Source: [github.com](https://github.com/matthiasnordwig/pdf-struct-chunker)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T09:06:32Z

## Translation

タイトル: 純粋な Rust の LLM フリー、レイアウト認識型 PDF チャンカー
記事のタイトル: GitHub - matthiasnordwig/pdf-struct-chunker: RAG パイプライン用の LLM フリー、レイアウト認識 PDF チャンク。正規表現とフォントのヒューリスティックによって文書構造を保持します。純粋な Rust で書かれています。 · GitHub
説明: RAG パイプライン用の LLM フリー、レイアウト認識型 PDF チャンク。正規表現とフォントのヒューリスティックによって文書構造を保持します。純粋な Rust で書かれています。 - matthiasnordwig/pdf-struct-chunker

記事本文:
GitHub - matthiasnordwig/pdf-struct-chunker: RAG パイプライン用の LLM フリー、レイアウト認識 PDF チャンク。正規表現とフォントのヒューリスティックによって文書構造を保持します。純粋な Rust で書かれています。 · GitHub
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
マティアスノルドウィッグ
/
pdf-struct-chunker
公共
ああ、ああ

！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
matthiasnordwig/pdf-struct-chunker
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github .github アセット アセット フィクスチャ フィクスチャ src src テスト テスト .gitignore .gitignore COTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM、API、クラウド依存関係を使用せずに、PDF を意味的に意味のあるチャンクに分割します。
🌐 著者: Matthias Nordwig ·programmiere.de
ほとんどの RAG チャンカーは、トークン数または文字制限によってドキュメントを盲目的に分割します。これにより、文書の構造が破壊され、見出し、セクション、段落がバラバラになってしまいます。結果: ベクトル検索では、どこから来たのかというコンテキストのない、支離滅裂な断片が返されます。
pdf-struct-chunker は、PDF の実際のレイアウト (X/Y 座標、フォント サイズ、太字の検出) を分析することでこれを解決します。見出しがどこで始まり、段落がどこで終わり、新しいセクションがどこで始まるかを理解します。各チャンクには構造化メタデータ (セクション、見出し、ページ) が含まれるため、RAG パイプラインは何を調べているのかを正確に認識できます。
LLMは必要ありません。 API 呼び出しはありません。オフラインで実行します。純粋な Rust で書かれています。
❌ 標準 RAG (固定サイズのオーバーラップ)
// チャンク 1
「この規制はすべての会社に適用されます。§ 2 解除」
// チャンク 2
「定義。この規程においては、次の用語をいう。」
結果: 単語は半分に切り取られ、見出しは内容から切り離されます。
✅ pdf-struct-chunker (レイアウト認識)
// チャンク 1
{
"メタデータ" : { "セクション" : " § 1 " , "見出し" : " スコープ " },
"text" : "この規制はすべての会社に適用されます。"
}
// チャンク

2
{
"メタデータ" : { "セクション" : " § 2 " , "見出し" : " 定義 " },
"text" : " この規則では、次の用語を... "
}
結果: 完璧なメタデータを備えたクリーンでセマンティックなチャンク。
スピードとエッジ AI シナリオ向けに構築:
GPUは不要（純粋なCPU処理）
非常に高速: 標準的なラップトップで 100 ページの PDF を 1 秒未満で処理します。
メモリ使用量が少ない : 一時ファイルを作成せずに完全にメモリ内で動作します。
サンプル PDF が含まれています。クローン作成後すぐに試すことができます。
git clone https://github.com/matthiasnordwig/pdf-struct-chunker.git
cd pdf-struct-chunker
カーゴラン --release -- -i fixtures/sample.pdf --format json --pretty
出力:
[
{
「インデックス」: 0 、
"char_start" : 0 、
"char_end" : 441 、
"text" : " § 1 Anwendungsbereich \n Diese Verordnung gilt für alle Unternehmen ... " ,
"署名" : " § 1 Anwendungsbereich \n Diese Verordnung gilt für alle Unternehmen " ,
「メタデータ」: {
"セクション" : " § 1 " 、
"見出し" : " アンウェンドゥングスベライヒ " ,
「ページ」：2
}
}、
{
「インデックス」: 1 、
"text" : "§ 2 Begriffsbestimmungen \n Im Sinne Dieser Verordnung ... " ,
「メタデータ」: {
"セクション" : " § 2 " 、
"見出し" : " Begriffsbestimmungen " ,
「ページ」：2
}
}
]
すべてのチャンクはそのセクション、見出し、ページ番号を認識しており、埋め込む準備ができています。
PDF バイト ──► pdf_oxyde (X/Y 位置 + フォント サイズで文字を抽出)
│
▼
ラインの分類
(行を正規表現プロファイルと照合するか、フォント サイズのヒューリスティックにフォールバックします)
│
▼
チャンクアセンブリ
(見出しで分割、小さな断片を結合、文の境界でオーバーフローを分割)
│
▼
Vec<Chunk> { テキスト、セクション、見出し、ページ }
チャンカーは、文字レベルの境界ボックスを抽出し、Y 座標から行を再構築し、構成可能な正規表現パターン (またはフォント サイズのヒューリスティック) を使用してそれらを分類することによって、各 PDF ページを処理します。

s をフォールバックとして）、構造メタデータを使用して意味的に一貫したチャンクにそれらを組み立てます。
git clone https://github.com/matthiasnordwig/pdf-struct-chunker.git
cd pdf-struct-chunker
カーゴビルド --release
Rust プロジェクトの依存関係として
[ 依存関係 ]
pdf-struct-chunker = { git = " https://github.com/matthiasnordwig/pdf-struct-chunker " }
CLIの使用法
pdf-struct-chunker [オプション] --input < INPUT >
旗
説明
デフォルト
-i, --input <ファイル>
入力 PDF ファイルへのパス
必須
-p、--profile <ファイル>
カスタム正規表現ルールを含む JSON プロファイルへのパス (以下を参照)
組み込みのデフォルト
-o, --output <ファイル>
出力ファイルのパス
標準出力
--format <フォーマット>
出力形式: jsonl または json
jsonl
--かなり
きれいに印刷された JSON 出力
偽
--統計
チャンク自体ではなくチャンク統計を出力します。
偽
例
# PDF をチャンクして JSONL として保存
pdf-struct-chunker -i document.pdf -o result.jsonl
# JSON をコンソールに整形出力する
pdf-struct-chunker -i document.pdf --format json --pretty
# 作成されたチャンクの数とそのサイズを確認する
pdf-struct-chunker -i document.pdf --stats
# 独自の正規表現ルールを使用する
pdf-struct-chunker -i document.pdf -p my_rules.json --format json --pretty
ライブラリ API (インメモリ)
コア関数は完全にメモリ内で動作します。ファイル I/O や一時ファイルはありません。どこからでも (ファイル、HTTP アップロード、S3、データベース) バイトをフィードし、即座にチャンクを取得します。
pdf_struct_chunker :: { chunk_pdf , プロファイル } を使用します。
fn メイン ( ) {
let bytes = std :: fs :: read ( "document.pdf" ) 。アンラップ ( ) ;
chunks = chunk_pdf (& bytes , None ) とします。アンラップ ( ) ;
チャンクインとチャンクの場合 {
プリントイン！ ( "[{}] {} (p.{})" ,
チャンク。メタデータ 。セクション、
チャンク。メタデータ 。見出し、
チャンク。メタデータ 。ページにunwrap_or ( 0 ) 、
) ;
}
}
カスタム正規表現プロファイル
デフォルトでは、チャンカーは組み込みのヒューリスティック最適化を使用します。

d 法律および規制文書​​の場合 (§ 、 Article 、 Chapter などの検出)。これを独自の正規表現ルールでオーバーライドできます。
.json ファイル (例: my_rules.json ) を作成し、 --profile 経由で渡します。
pdf-struct-chunker -i document.pdf -p my_rules.json
簡単な例 - ページ番号を無視する
最も単純なプロファイルは、不要な行を削除するだけです。
{
「パターン」: [
{
"役割" : " 無視 " ,
"正規表現" : " ページ \\ d+ " ,
"フラグ" : " i " 、
「優先度」：100
}
]
}
これにより、「ページ 1」、「ページ 2」などに一致するすべての行が出力から削除されます。
完全な例 — カスタム見出しと定義
{
"min_chunk_chars" : 200 、
"max_chunk_chars" : 1500 、
「パターン」: [
{
"役割" : " 無視 " ,
"regex" : " (?:ページ|フッターテキスト) " ,
"フラグ" : " i " 、
「優先度」：200
}、
{
"役割" : "Heading_l1 " ,
"正規表現" : " ^((?:章|セクション) \\ s*[ \\ d]+) \\ s*(.*) " ,
"フラグ" : " i " 、
「優先度」：100
}、
{
"役割" : "定義" ,
"regex" : " \\ b(?:means|shall means|は次のように定義されます) " ,
"フラグ" : " i " 、
「優先度」：50
}
]
}
パターンの役割
役割
何をするのか
見出し_l1
新しいチャンクを開始します。正規表現キャプチャ グループ 1 は、metadata.section (例: 「第 3 章」) になり、グループ 2 は、metadata.Heading (例: 「データ保護」) になります。
定義
ソフト スプリットをトリガーします。現在のチャンクがすでに min_chunk_chars に達している場合、チャンカーはそれをフラッシュし、新しいチャンクを開始します。
無視する
行を完全に削除します。これは、ページ番号、フッター、ヘッダー、またはチャンクに含めたくない定型文に使用します。
プロフィールフィールド
フィールド
説明
デフォルト
min_chunk_chars
「ソフト」分割 (定義またはリスト項目で) が許可される前の最小チャンク サイズ
200
max_chunk_chars
最大チャンク サイズ — 最も近い文の境界で強制的に分割します。
1500
パターン[].regex
各テキスト行と照合される正規表現
—
パターン[].役割
次のいずれか: 見出し_l1 、 def

初期化、無視
—
パターン[].フラグ
「i」 = 大文字と小文字を区別しない、「m」 = 複数行
「」
パターン[].優先度
高い値 = 複数のパターンが同じ行に一致する場合に最初に評価されます
0
連絡先とフィードバック
ご質問、機能リクエスト、または単にご挨拶したい場合は、お気軽に問題を開くか、私の Web サイトからご連絡ください。
このツールで時間を節約でき、その開発をサポートしたい場合は、PayPal 経由でコーヒーを買ってください。 ☕
RAG パイプライン向けの LLM フリーのレイアウト認識 PDF チャンク化。正規表現とフォントのヒューリスティックによって文書構造を保持します。純粋な Rust で書かれています。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
https://www.paypal.me/MatthiasNordwig
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

LLM-free, layout-aware PDF chunking for RAG pipelines. Preserves document structure via regex and font heuristics. Written in pure Rust. - matthiasnordwig/pdf-struct-chunker

GitHub - matthiasnordwig/pdf-struct-chunker: LLM-free, layout-aware PDF chunking for RAG pipelines. Preserves document structure via regex and font heuristics. Written in pure Rust. · GitHub
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
matthiasnordwig
/
pdf-struct-chunker
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
matthiasnordwig/pdf-struct-chunker
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github .github assets assets fixtures fixtures src src tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
Split PDFs into semantically meaningful chunks — without LLMs, without APIs, without cloud dependencies.
🌐 Author: Matthias Nordwig · programmiere.de
Most RAG chunkers blindly split documents by token count or character limit. This destroys document structure — headings, sections, and paragraphs get ripped apart. The result: your vector search returns incoherent fragments with no context about where they came from.
pdf-struct-chunker solves this by analyzing the actual layout of a PDF: X/Y coordinates, font sizes, and bold detection. It understands where a heading starts, where a paragraph ends, and where a new section begins. Each chunk carries structured metadata ( section , heading , page ) so your RAG pipeline knows exactly what it's looking at.
No LLM needed. No API calls. Runs offline. Written in pure Rust.
❌ Standard RAG (Fixed-size overlap)
// Chunk 1
" This regulation applies to all companies. § 2 De- "
// Chunk 2
" finitions. In this regulation, the following terms "
Result: Words are cut in half, headings are disconnected from their content.
✅ pdf-struct-chunker (Layout-aware)
// Chunk 1
{
"metadata" : { "section" : " § 1 " , "heading" : " Scope " },
"text" : " This regulation applies to all companies. "
}
// Chunk 2
{
"metadata" : { "section" : " § 2 " , "heading" : " Definitions " },
"text" : " In this regulation, the following terms... "
}
Result: Clean, semantic chunks with perfect metadata.
Built for speed and Edge-AI scenarios:
No GPU required (pure CPU processing)
Extremely fast : Processes a 100-page PDF in < 1 second on a standard laptop.
Low memory footprint : Operates entirely in-memory without creating temporary files.
A sample PDF is included — you can try it immediately after cloning:
git clone https://github.com/matthiasnordwig/pdf-struct-chunker.git
cd pdf-struct-chunker
cargo run --release -- -i fixtures/sample.pdf --format json --pretty
Output:
[
{
"index" : 0 ,
"char_start" : 0 ,
"char_end" : 441 ,
"text" : " § 1 Anwendungsbereich \n Diese Verordnung gilt für alle Unternehmen ... " ,
"signature" : " § 1 Anwendungsbereich \n Diese Verordnung gilt für alle Unternehmen " ,
"metadata" : {
"section" : " § 1 " ,
"heading" : " Anwendungsbereich " ,
"page" : 2
}
},
{
"index" : 1 ,
"text" : " § 2 Begriffsbestimmungen \n Im Sinne dieser Verordnung ... " ,
"metadata" : {
"section" : " § 2 " ,
"heading" : " Begriffsbestimmungen " ,
"page" : 2
}
}
]
Every chunk knows its section, heading, and page number — ready for embedding.
PDF bytes ──► pdf_oxide (extract characters with X/Y positions + font sizes)
│
▼
Line Classification
(match lines against your regex profiles, or fall back to font-size heuristics)
│
▼
Chunk Assembly
(split at headings, merge small fragments, split overflow at sentence boundaries)
│
▼
Vec<Chunk> { text, section, heading, page }
The chunker processes each PDF page by extracting character-level bounding boxes, reconstructing lines from Y-coordinates, classifying them using configurable regex patterns (or font-size heuristics as fallback), and assembling them into semantically coherent chunks with structural metadata.
git clone https://github.com/matthiasnordwig/pdf-struct-chunker.git
cd pdf-struct-chunker
cargo build --release
As a Dependency in Your Rust Project
[ dependencies ]
pdf-struct-chunker = { git = " https://github.com/matthiasnordwig/pdf-struct-chunker " }
CLI Usage
pdf-struct-chunker [OPTIONS] --input < INPUT >
Flag
Description
Default
-i, --input <FILE>
Path to the input PDF file
Required
-p, --profile <FILE>
Path to a JSON profile with custom regex rules (see below)
Built-in defaults
-o, --output <FILE>
Output file path
stdout
--format <FORMAT>
Output format: jsonl or json
jsonl
--pretty
Pretty-print JSON output
false
--stats
Print chunk statistics instead of the chunks themselves
false
Examples
# Chunk a PDF and save as JSONL
pdf-struct-chunker -i document.pdf -o result.jsonl
# Pretty-print JSON to the console
pdf-struct-chunker -i document.pdf --format json --pretty
# See how many chunks were created and their sizes
pdf-struct-chunker -i document.pdf --stats
# Use your own regex rules
pdf-struct-chunker -i document.pdf -p my_rules.json --format json --pretty
Library API (In-Memory)
The core function operates entirely in-memory — no file I/O, no temp files. Feed it bytes from anywhere (file, HTTP upload, S3, database) and get chunks back instantly:
use pdf_struct_chunker :: { chunk_pdf , Profile } ;
fn main ( ) {
let bytes = std :: fs :: read ( "document.pdf" ) . unwrap ( ) ;
let chunks = chunk_pdf ( & bytes , None ) . unwrap ( ) ;
for chunk in & chunks {
println ! ( "[{}] {} (p.{})" ,
chunk . metadata . section ,
chunk . metadata . heading ,
chunk . metadata . page . unwrap_or ( 0 ) ,
) ;
}
}
Custom Regex Profiles
By default, the chunker uses built-in heuristics optimized for legal and regulatory documents (detecting § , Article , Chapter , etc.). You can override this with your own regex rules.
Create a .json file (e.g., my_rules.json ) and pass it via --profile :
pdf-struct-chunker -i document.pdf -p my_rules.json
Simple Example — Ignore Page Numbers
The simplest profile just removes unwanted lines:
{
"patterns" : [
{
"role" : " ignore " ,
"regex" : " Page \\ d+ " ,
"flags" : " i " ,
"priority" : 100
}
]
}
This removes every line matching "Page 1", "Page 2", etc. from the output.
Full Example — Custom Headings and Definitions
{
"min_chunk_chars" : 200 ,
"max_chunk_chars" : 1500 ,
"patterns" : [
{
"role" : " ignore " ,
"regex" : " (?:Page|Footer text) " ,
"flags" : " i " ,
"priority" : 200
},
{
"role" : " heading_l1 " ,
"regex" : " ^((?:Chapter|Section) \\ s*[ \\ d]+) \\ s*(.*) " ,
"flags" : " i " ,
"priority" : 100
},
{
"role" : " definition " ,
"regex" : " \\ b(?:means|shall mean|is defined as) " ,
"flags" : " i " ,
"priority" : 50
}
]
}
Pattern Roles
Role
What it does
heading_l1
Starts a new chunk. Regex capture group 1 becomes metadata.section (e.g., "Chapter 3"), group 2 becomes metadata.heading (e.g., "Data Protection").
definition
Triggers a soft split. If the current chunk has already reached min_chunk_chars , the chunker flushes it and starts a new one.
ignore
Removes the line entirely. Use this for page numbers, footers, headers, or any boilerplate you don't want in your chunks.
Profile Fields
Field
Description
Default
min_chunk_chars
Minimum chunk size before a "soft" split (at definitions or list items) is allowed
200
max_chunk_chars
Maximum chunk size — forces a split at the nearest sentence boundary
1500
patterns[].regex
Regular expression matched against each text line
—
patterns[].role
One of: heading_l1 , definition , ignore
—
patterns[].flags
"i" = case-insensitive, "m" = multiline
""
patterns[].priority
Higher value = evaluated first when multiple patterns match the same line
0
Contact & Feedback
If you have any questions, feature requests, or just want to say hi, feel free to open an issue or reach out via my website:
If this tool saved you time and you'd like to support its development, you can buy me a coffee via PayPal . ☕
LLM-free, layout-aware PDF chunking for RAG pipelines. Preserves document structure via regex and font heuristics. Written in pure Rust.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
https://www.paypal.me/MatthiasNordwig
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
