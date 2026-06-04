---
source: "https://github.com/Fujo930/moss-lang"
hn_url: "https://news.ycombinator.com/item?id=48402574"
title: "Show HN: Moss, an AI-led programming language experiment"
article_title: "GitHub - Fujo930/moss-lang: A runnable prototype of the Moss programming language and browser studio. · GitHub"
author: "Fujo930"
captured_at: "2026-06-04T18:52:57Z"
capture_tool: "hn-digest"
hn_id: 48402574
score: 1
comments: 0
posted_at: "2026-06-04T18:22:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Moss, an AI-led programming language experiment

- HN: [48402574](https://news.ycombinator.com/item?id=48402574)
- Source: [github.com](https://github.com/Fujo930/moss-lang)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T18:22:44Z

## Translation

タイトル: Show HN: Moss、AI 主導のプログラミング言語の実験
記事タイトル: GitHub - Fujo930/moss-lang: Moss プログラミング言語およびブラウザ スタジオの実行可能なプロトタイプ。 · GitHub
説明: Moss プログラミング言語およびブラウザー スタジオの実行可能なプロトタイプ。 - Fujo930/moss-lang

記事本文:
GitHub - Fujo930/moss-lang: Moss プログラミング言語およびブラウザ スタジオの実行可能なプロトタイプ。 · GitHub
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
ふじょ930
/
モスラング
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
48 コミット 48 コミット ドキュメント ドキュメントの例 例 src/ mosslang src/ mosslang テスト テスト .gitattributes .gitattributes .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Moss は、長期にわたるソフトウェア プロジェクトのための実験的なプログラミング言語です。
人間と AI エージェントは、時間をかけて同じコードベースで作業します。
このリポジトリは意図的に AI によって構築されています。Moss は設計、実装され、
Codex との協力によりデバッグ、文書化、コミット、プッシュされます。
ふじょ930。プロジェクトは、そのプロセスの記録および実行可能ファイルとして公開されます。
言語のプロトタイプ。
バージョン 0.2.0 はセルフホスティング プレビューです。 Moss はまだ完全に自己ホストされていません。
ただし、Moss が作成したレクサー、パーサー、チェッカー、プロジェクト チェック スケッチはすでに実行されています
モスソースに対して。
python -m pip install -e 。
苔チェック例 / order.moss
moss の実行例 / order.moss
moss テスト例 / order.moss
苔セルフホスト
moss セルフホスト -- 迅速
モススタジオ
以下をインストールせずに実行することもできます。
Python - m mosslang.cli 実行例 / order.moss
ローカル リリース アーティファクトをビルドするには:
Python - m pip インストール ビルド
Python - m ビルド
このパッケージは、 moss という名前のコンソール コマンドを公開します。
レコードと単純な共用体の型宣言
純粋な式関数としてのルール宣言
オプションの fn 宣言は EffectName を使用します
言語レベルの実行可能ファイルをチェックするための "name" { ... } ブロックのテスト
レコード、レコード フィールド アクセス、およびレコードの更新
if 、 else if 、および else ブロック
リストリテラル、インデックス付け、for ループ、 len 、 listPush 、 listGet 、
listSet 、 listSlice 、 listConcat 、 listInsert 、 listRemove 、および
範囲
Map<K, V> から、mapNew 、mapPut 、mapGet 、mapHas 、mapKeys 、
mapValues および mapRemove
テキスト

ヘルパー: textChars 、 textJoin 、 textSplit 、 textTrim 、 textSlice 、
textContains 、 textIndexOf 、 textReplace 、 textStartsWith 、および
テキストで終わる
FileSystem エフェクトの組み込み: readText 、 writeText 、 fileExists 、および
リストファイル
最上位インポート「path.moss」宣言
構造化されたトークンレコードを備えた自己ホスト型スケッチ、再利用可能なレクサー/パーサー
コア、トップレベルの宣言パーサー、および最初のチェッカー スケッチ:
Examples/self_host/tokenizer_sketch.moss および
Examples/self_host/parser_sketch.moss および
例/self_host/checker_sketch.moss
moss selfhost 、トークナイザー/パーサー/チェッカーのスケッチに加えて実行します。
例/self_host/project_check.moss ;プロジェクトチェックは解析してチェックします
Moss コードを含む自己ホスティング Moss ファイル
Payd や ShipError.NotReady(Pending) などの nullary およびペイロードのバリアント
ワイルドカードおよびペイロード バインディング パターンを使用した一致式
Ok(...) 、 Err(...) 、および ? を含む結果値
require 条件 else value 、これは Result から Err(value) を返します
機能
関数の引数と戻り値のランタイム型規約
List<T> 、 Map<K, V> 、および Option<T> ランタイム タイプ コントラクト
dbPut および dbGet を介した小さなメモリ内データベース。
関数内のデータベース効果
効果データベース
タイプ 順序 =
ID: テキスト
ステータス: 保留中 |有料 |出荷済み |キャンセルされました
合計: お金
タイプ ShipError = NotReady |行方不明
ルール canShip(order: Order) -> Bool =
order.status == 支払い済み、order.total > 0.usd
fn ship(order: Order) -> Result<Order, ShipError> はデータベースを使用します {
canShip(注文)が必要です
else ShipError.NotReady(order.status)
更新 = 注文のステータス = 出荷済み
dbPut(order.id、更新済み)
OK を返します (更新)
}
let order = { id: "A-100"、ステータス: 支払い済み、合計: 42.usd }
出荷しましょう = 出荷（注文）?
print("ステータス:", 出荷済み.ステータス)
print("stored:", dbGet("A-100").status)
コマンド
モスC

てか < file.moss >
moss 実行 < file.moss >
moss テスト < file.moss >
moss トークン < file.moss >
moss ast < file.moss >
苔セルフホスト
moss セルフホスト -- 迅速
モススタジオ
moss スタジオは、 http://127.0.0.1:8765 でローカル HTTP エディターを開きます。
moss selfhost --quick は高速セルフホスティング スケッチを実行します。苔セルフホスト
また、 example/self_host を介して、低速の Moss 作成プロジェクト チェックも実行されます。
これはバージョン 0.2.0 です。実際の構文とランタイムを備えたコンパクトなインタープリターです。
セマンティクス、ブラウザ エディタ、および Moss が作成したセルフホスティング スケッチです。
リポジトリは MIT ライセンスに基づいてリリースされています。
Moss は AI によって設計され、AI によって構築されます。
Moss は今日から便利なサンプル プログラムを実行できます。
Moss はまだアルファ版のソフトウェアであり、完全に自己ホストされていると表現すべきではありません。
次の有用なステップは、構造化された Moss AST、より豊富な診断、フォーマッタ、
そして、Moss が作成したフロントエンドの出力と Python ホストのフロントエンドを比較します。
GitHub の言語バーは Linguist を利用しています。 .moss ファイルにはマークが付いています
.gitattributes で検出可能ですが、GitHub では Moss のみがファーストクラスとして表示されます
Moss が上流の言語学者の言語リストに受け入れられた後の言語。
現在の言語サーフェスについては docs/ language.md を参照してください。
ブラウザエディターの場合は docs/studio.md、
docs/history.md (コミットごとの機能ガイド)、および
プロトタイプから本格的な実装までのパスについては、docs/roadmap.md を参照してください。
0.2.0 の公開リリース ノートとパッケージについては、docs/release.md を参照してください。
チェックリスト。
Moss プログラミング言語およびブラウザー スタジオの実行可能なプロトタイプ。
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

A runnable prototype of the Moss programming language and browser studio. - Fujo930/moss-lang

GitHub - Fujo930/moss-lang: A runnable prototype of the Moss programming language and browser studio. · GitHub
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
Fujo930
/
moss-lang
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
48 Commits 48 Commits docs docs examples examples src/ mosslang src/ mosslang tests tests .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Moss is an experimental programming language for long-lived software projects
where humans and AI agents work on the same codebase over time.
This repository is intentionally AI-built: Moss was designed, implemented,
debugged, documented, committed, and pushed by Codex in collaboration with
Fujo930. The project is public as a record of that process and as a runnable
language prototype.
Version 0.2.0 is a self-hosting preview. Moss is not fully self-hosted yet,
but Moss-written lexer, parser, checker, and project-check sketches already run
against Moss source.
python - m pip install - e .
moss check examples / order.moss
moss run examples / order.moss
moss test examples / order.moss
moss selfhost
moss selfhost -- quick
moss studio
You can also run without installing:
python - m mosslang.cli run examples / order.moss
To build local release artifacts:
python - m pip install build
python - m build
The package exposes a console command named moss .
type declarations for records and simple unions
rule declarations as pure expression functions
fn declarations with optional uses EffectName
test "name" { ... } blocks for language-level executable checks
records, record field access, and record updates
if , else if , and else blocks
list literals, indexing, for loops, len , listPush , listGet ,
listSet , listSlice , listConcat , listInsert , listRemove , and
range
Map<K, V> through mapNew , mapPut , mapGet , mapHas , mapKeys ,
mapValues , and mapRemove
Text helpers: textChars , textJoin , textSplit , textTrim , textSlice ,
textContains , textIndexOf , textReplace , textStartsWith , and
textEndsWith
FileSystem effect builtins: readText , writeText , fileExists , and
listFiles
top-level import "path.moss" declarations
self-hosting sketches with structured token records, reusable lexer/parser
cores, a top-level declaration parser, and a first checker sketch:
examples/self_host/tokenizer_sketch.moss and
examples/self_host/parser_sketch.moss and
examples/self_host/checker_sketch.moss
moss selfhost , which runs the tokenizer/parser/checker sketches plus
examples/self_host/project_check.moss ; the project check parses and checks
the self-hosting Moss files with Moss code
nullary and payload variants such as Paid and ShipError.NotReady(Pending)
match expressions with wildcard and payload binding patterns
Result values with Ok(...) , Err(...) , and ?
require condition else value , which returns Err(value) from Result
functions
runtime type contracts for function arguments and return values
List<T> , Map<K, V> , and Option<T> runtime type contracts
a tiny in-memory database through dbPut and dbGet , guarded by the
Database effect inside functions
effect Database
type Order =
id: Text
status: Pending | Paid | Shipped | Cancelled
total: Money
type ShipError = NotReady | Missing
rule canShip(order: Order) -> Bool =
order.status == Paid and order.total > 0.usd
fn ship(order: Order) -> Result<Order, ShipError> uses Database {
require canShip(order)
else ShipError.NotReady(order.status)
updated = order with status = Shipped
dbPut(order.id, updated)
return Ok(updated)
}
let order = { id: "A-100", status: Paid, total: 42.usd }
let shipped = ship(order)?
print("status:", shipped.status)
print("stored:", dbGet("A-100").status)
Commands
moss check < file.moss >
moss run < file.moss >
moss test < file.moss >
moss tokens < file.moss >
moss ast < file.moss >
moss selfhost
moss selfhost -- quick
moss studio
moss studio opens a local HTTP editor at http://127.0.0.1:8765 .
moss selfhost --quick runs the fast self-hosting sketches. moss selfhost
also runs the slower Moss-written project check over examples/self_host .
This is version 0.2.0 : a compact interpreter with real syntax, runtime
semantics, a browser editor, and Moss-written self-hosting sketches.
The repository is released under the MIT License.
Moss is AI-designed and AI-built.
Moss can run useful example programs today.
Moss is still alpha software and should not be described as fully self-hosted.
The next useful steps are a structured Moss AST, richer diagnostics, a formatter,
and comparing Moss-written frontend output against the Python host frontend.
GitHub's language bar is powered by Linguist. .moss files are marked
detectable in .gitattributes , but GitHub will only show Moss as a first-class
language after Moss is accepted into the upstream Linguist language list.
See docs/language.md for the current language surface,
docs/studio.md for the browser editor,
docs/history.md for a commit-by-commit feature guide, and
docs/roadmap.md for the path from prototype to a serious implementation.
See docs/release.md for the public 0.2.0 release notes and packaging
checklist.
A runnable prototype of the Moss programming language and browser studio.
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
