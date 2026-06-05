---
source: "https://github.com/typescope/jo"
hn_url: "https://news.ycombinator.com/item?id=48412229"
title: "Show HN: Jo – AI-native language to catch prompt injection at compile-time"
article_title: "GitHub - typescope/jo: Jo Secure Programming Language · GitHub"
author: "kiru_io"
captured_at: "2026-06-05T16:08:22Z"
capture_tool: "hn-digest"
hn_id: 48412229
score: 7
comments: 3
posted_at: "2026-06-05T13:25:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Jo – AI-native language to catch prompt injection at compile-time

- HN: [48412229](https://news.ycombinator.com/item?id=48412229)
- Source: [github.com](https://github.com/typescope/jo)
- Score: 7
- Comments: 3
- Posted: 2026-06-05T13:25:51Z

## Translation

タイトル: Show HN: Jo – コンパイル時にプロンプト インジェクションをキャッチする AI ネイティブ言語
記事タイトル: GitHub - typescope/jo: Jo セキュア プログラミング言語 · GitHub
説明: Jo セキュア プログラミング言語。 GitHub でアカウントを作成して、typescope/jo の開発に貢献してください。

記事本文:
GitHub - typescope/jo: Jo セキュア プログラミング言語 · GitHub
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
タイプスコープ
/
じょ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイル コードに移動 さらにアクションを開く

メニュー フォルダーとファイル
6,119 コミット 6,119 コミット .github .github 資産/ ドキュメント アセット/ ドキュメント bin bin デモ デモ ドキュメント ドキュメント lib lib ランタイム ランタイム stack-lang stack-lang テスト テスト ツール ツール .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 NOTICE README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md build build ci ci dev.md dev.md project.scala project.scala すべてのファイルを表示 リポジトリ ファイルのナビゲーション
安全なプログラミングを楽しむために
Jo は静的に型付けされた言語であり、機能が明示的で、静的に追跡され、コンパイラーによって強制されます。 Jo は Ruby と Python にコンパイルします。
プロジェクトのステータス: 初期段階。コンパイラ、標準ライブラリ、およびツールチェーンは本格的な実験の準備ができています。 API と言語の詳細は今後も変更される可能性があります。
AI エージェントは、プラットフォーム内で実行されるコードを生成します。そのコードは、ユーザーが防止しない限り、ネットワークにアクセスしたり、任意のファイルを読み取ったり、他のユーザーのデータをクエリしたりする可能性があります。ランタイム サンドボックスは役立ちますが、間違ったレベルで動作します。syscall やファイル システム パスをブロックできますが、「このユーザーの行のみにアクセスする」ことを強制することはできません。
Jo は、プログラムが実行される前に、型レベルで機能境界を強制します。ケイパビリティを受け取っていない機能は使用できません。コンパイラーは、呼び出しグラフ全体を通じてこれを推移的に証明します。
def foo () = println " foo " // 推論された機能: stdout
def bar () = foo() // 推論された機能: stdout
def qux () は IO を受け取ります .stdout = println " qux " // 明示的: stdout のみ
デフォルトメイン =
bar() では none を許可 // エラー: 許可される機能はありません
bar() で IO .stdout を許可 // OK
with IO .stdout = s => pass in qux() // 出力をリダイレクト
---------- main.jo:5:3 でのエラー ---------------
| nを許可する

1 つは bar() にあります
| ^^^^
|パラメータは許可されません: stdout
問題の原因となるトレースは次のとおりです。
§── bar() では何も許可しません [ main.jo:5:3 ]
│＾＾＾＾
§── def bar() = foo() [ main.jo:2:13 ]
│＾＾＾＾
━── def foo() = println "foo" [ main.jo:1:13 ]
^^^^^^
パターン指向プログラミング
名前付きの再利用可能なパターン述語は、論理演算子で構成されます。
パターン 正: 部分 [ Int ] = case x if x > 0
マッチリスト
case [..陽性の間陽性、..残り] =>
println " pos = \{positives},rest = \{rest} "
case _ => パス
結果が Some (code) && code > 0 の場合
println " 成功、コード = \{ code} "
// オプション "s" を有効にして許可します。新しい行に一致させる
メッセージが `(?s)<code>(?<prog>.*)</code>` の場合
printlnプログラム
AI 生成コードの制限
2 ワールド アーキテクチャは、制限されたコード (FFI なし、機能インターフェイスのみに対してチェックされる) を信頼できるコード (FFI が許可され、機能を実装および提供する) から分離します。
// --- インターフェース ライブラリ (制限付き、FFI なし) ---
param orderApi : OrdersApi
defer def aiMain () : ユニットはordersApi、IO .stdoutを受け取ります
// --- フレームワーク ハーネス (信頼できる、FFI 許可) ---
def フレームワークメイン () =
val db = connect( "orders.db " )
val userId = currentUser()
vallimited = new UserScopedOrders (userId, db) // 減衰: ユーザースコープ、読み取り専用
val バッファ = ( s : 文字列 ) => 出力 += s
何も許可しない
orderApi = 制限付き、IO .stdout = aiMain() のバッファー
// --- AI 生成コード (制限付き、FFI なし) ---
def aiMain () : ユニットはordersApi、IO .stdoutを受け取ります =
val 注文 =ordersApi.query( 30 )
要約（注文）
allowed none はコンパイル時の証明です。aiMain() は、宣言された内容を超える機能を使用しません。 AI コードは、ネットワーク、ファイル システム、または他のユーザーのデータにアクセスできません。
完全なモデルについては、セキュリティに関するドキュメントを参照してください。
カール -sSf

https://jo-lang.org/install.sh |しー
インストーラーはコンパイラーを ~/.jo/compilers/<version>/ にダウンロードし、 ~/.local/bin/jo にランチャーを作成します。 ~/.local/bin がまだ存在していない場合は、PATH に追加します。
完全なインストール手順とスタートガイドは jo-lang.org にあります。
語学ツアー
Jo の機能の概要と例
セキュリティモデル
機能の強制の仕組み
言語リファレンス
型、式、パターン、定義
ビルドツール
プロジェクトのセットアップ、依存関係、コマンド
貢献する
ビルド手順、貢献ガイドライン、DCO サインオフ要件については、CONTRIBUTING.md を参照してください。バグレポート、言語設計に関するディスカッション、プルリクエストは大歓迎です。
セキュリティの問題は非公開で報告する必要があります。SECURITY.md を参照してください。
Jo は TypeScope によって開発および保守されています。
ジョーセキュアプログラミング言語
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
1
ジョー 0.10.0
最新の
2026 年 6 月 4 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Jo Secure Programming Language. Contribute to typescope/jo development by creating an account on GitHub.

GitHub - typescope/jo: Jo Secure Programming Language · GitHub
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
typescope
/
jo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6,119 Commits 6,119 Commits .github .github assets/ doc assets/ doc bin bin demos demos docs docs lib lib runtime runtime stack-lang stack-lang tests tests tools tools .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md build build ci ci dev.md dev.md project.scala project.scala View all files Repository files navigation
For the joy of secure programming
Jo is a statically typed language where capabilities are explicit, statically tracked, and enforced by the compiler . Jo compiles to Ruby and Python.
Project status: Early-stage. The compiler, standard library, and toolchain are ready for serious experimentation. APIs and language details may still change.
AI agents generate code that runs inside your platform. That code can — unless you prevent it — reach for the network, read arbitrary files, or query other users' data. Runtime sandboxes help, but they operate at the wrong level: they can block syscalls or filesystem paths, but they cannot enforce "access only this user's rows".
Jo enforces capability boundaries at the type level, before the program runs. A function that has not received a capability cannot use it. The compiler proves this transitively through the entire call graph.
def foo () = println " foo " // inferred capability: stdout
def bar () = foo() // inferred capability: stdout
def qux () receives IO .stdout = println " qux " // explicit: only stdout
def main =
allow none in bar() // error: no capabilities allowed
allow IO .stdout in bar() // OK
with IO .stdout = s => pass in qux() // redirect output
---------- Error at main.jo:5:3 ---------------
| allow none in bar()
| ^^^^^
| Parameter not allowed: stdout
The following is the trace that leads to the problem:
├── allow none in bar() [ main.jo:5:3 ]
│ ^^^^^
├── def bar() = foo() [ main.jo:2:13 ]
│ ^^^^^
└── def foo() = println "foo" [ main.jo:1:13 ]
^^^^^^^
Pattern-oriented programming
Named, reusable pattern predicates compose with logical operators:
pattern Positive : Partial [ Int ] = case x if x > 0
match list
case [..positives while Positive , ..rest] =>
println " pos = \{ positives}, rest = \{ rest} "
case _ => pass
if result is Some (code) && code > 0 then
println " Success, code = \{ code} "
// enable option "s" to allow . to match new line
if message is `(?s)<code>(?<prog>.*)</code>` then
println prog
Confining AI-Generated Code
The two-world architecture separates confined code (no FFI, checked against capability interfaces only) from trusted code (FFI allowed, implements and provides capabilities):
// --- Interface library (confined, no FFI) ---
param ordersApi : OrdersApi
defer def aiMain () : Unit receives ordersApi, IO .stdout
// --- Framework harness (trusted, FFI allowed) ---
def frameworkMain () =
val db = connect( " orders.db " )
val userId = currentUser()
val restricted = new UserScopedOrders (userId, db) // attenuated: user-scoped, read-only
val buffer = ( s : String ) => output += s
allow none in
with ordersApi = restricted, IO .stdout = buffer in aiMain()
// --- AI-generated code (confined, no FFI) ---
def aiMain () : Unit receives ordersApi, IO .stdout =
val orders = ordersApi.query( 30 )
summarize(orders)
allow none is a compile-time proof: aiMain() uses no capabilities beyond what it declared. The AI code cannot access the network, filesystem, or other users' data.
See the security documentation for the full model.
curl -sSf https://jo-lang.org/install.sh | sh
The installer downloads the compiler to ~/.jo/compilers/<version>/ and creates a launcher at ~/.local/bin/jo . Add ~/.local/bin to PATH if it is not already there.
Full installation instructions and a getting-started guide are at jo-lang.org .
Language Tour
Overview of Jo's features with examples
Security Model
How capability enforcement works
Language Reference
Types, expressions, patterns, definitions
Build Tool
Project setup, dependencies, commands
Contributing
See CONTRIBUTING.md for build instructions, contribution guidelines, and the DCO sign-off requirement. Bug reports, language design discussions, and pull requests are welcome.
Security issues should be reported privately — see SECURITY.md .
Jo is developed and maintained by TypeScope .
Jo Secure Programming Language
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
Jo 0.10.0
Latest
Jun 4, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
