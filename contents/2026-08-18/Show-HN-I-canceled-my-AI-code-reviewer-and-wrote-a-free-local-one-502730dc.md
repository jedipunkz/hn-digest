---
source: "https://github.com/mukundzha/avouch"
hn_url: "https://news.ycombinator.com/item?id=49345154"
title: "Show HN: I canceled my AI code reviewer and wrote a free local one"
article_title: "GitHub - mukundzha/avouch: Review the Python you changed, not the Python you inherited. Git-aware AST code review that runs in the seconds before git push. · GitHub"
image: "https://repository-images.githubusercontent.com/1312017997/eed8069b-df0f-499e-9c1a-7aec7003c172"
author: "mukundzha6"
captured_at: "2026-08-18T13:36:04Z"
capture_tool: "hn-digest"
hn_id: 49345154
score: 3
comments: 2
posted_at: "2026-08-18T13:13:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I canceled my AI code reviewer and wrote a free local one

- HN: [49345154](https://news.ycombinator.com/item?id=49345154)
- Source: [github.com](https://github.com/mukundzha/avouch)
- Score: 3
- Comments: 2
- Posted: 2026-08-18T13:13:17Z

## Translation

タイトル: HN を表示: AI コード レビューアーをキャンセルして、無料のローカル コード レビューアーを作成しました
記事タイトル: GitHub - mukundzha/avouch: 継承した Python ではなく、変更した Python を確認してください。 git プッシュ前の数秒で実行される Git 対応 AST コード レビュー。 · GitHub
説明: 継承した Python ではなく、変更した Python を確認します。 git プッシュ前の数秒で実行される Git 対応 AST コード レビュー。 - ムクンジャ/アバウシュ

記事本文:
GitHub - mukundzha/avouch: 継承した Python ではなく、変更した Python を確認してください。 git プッシュ前の数秒で実行される Git 対応 AST コード レビュー。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ムクンジャ
/
保証する
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
132 コミット 132 コミット フォルダーとファイル
.github/ ワークフロー .github/ ワークフロー 画像 画像 src src テスト テスト .gitignore .gitignore CHANGELOG.md CHAN

GELOG.md ライセンス ライセンス README.md README.md avouch.toml avouch.toml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
継承した Python ではなく、変更した Python を確認してください。
Avouch は、Python 用の軽量の Git 対応静的分析 CLI です。それは尋ねます
Git で次のコミットがどのファイルにアクセスするか、変更された各 .py を解析します
標準の ast モジュールを使用してファイルを作成し、構造的な問題を報告します
avouch.toml で設定した制限に対して。
デーモンはありません。ネットワークがありません。維持するパス リストはありません。数秒で実行できます
git Push の前に、フラグを立てているものを修正し、プッシュします。
pip インストールの保証
cd あなたのリポジトリ
保証する
目次
レビュー セットはリポジトリではなく差分です。 Avouch は、
実行時に Git からセットを確認します ( git diff HEAD --name-only plus
追跡されていないファイル)。すべての発見は、あなたが取り組んでいる仕事に起因します
押し付けるのではなく、あなたが受け継いだ遺産を押し付けることは決してありません。
メトリクスは正確です。パラメータ数、ネストの深さ、および行範囲
正規表現ではなく AST から取得されます。メトリクスを正確に計算できない場合は、
アヴーシュはそれを主張していません。
エラーはデータです。読み取れない、または構文的に壊れたファイルは、
レポート内の ERROR エントリ。 1 つの壊れたファイルがキャンセルされることはありません
他の人のレビュー。
レビューを公言します。ゲートはしません。終了コードは結果を示します —
0 件はクリーン、1 件は違反が見つかり、2 件はエラーを認めますが、強制は適用されます
プッシュするたびに実行するツールではなく、オプトイン インターフェイスで。
ランタイムは標準ライブラリです。 3 つの git サブプロセス呼び出し
そして
ast/tomllib 。存続させるデーモンはありません。ランタイムは次の制限を受けます
リポジトリではなく、diff のサイズです。
Python 3.10以降が必要です(ルールは ast.Match を使用し、構成は使用します)
tomllib ) と PATH 上の Git 。
pip インストールの保証
またはソースから:
git clone https://github.com/mukundzha/avouch.git
CD 保証
pip install -e 。
どちらも登録します

avouch コンソール スクリプト ( avouch.cli:main )。
インターフェイスは、オプションのフラグの小さなセットを備えた 1 つのコマンドです。
cd あなたのリポジトリ
# ...変更を加える ...
# 人間の報告を公言する
avouch --json # 標準出力上の 1 つの JSON ドキュメント
avouch --docs # 組み込みドキュメント;レビューは行われていません
avouch --version # バージョンを出力して終了します
avouch --verbose # 標準エラー出力の詳細を段階的に確認します
avouch --quiet # 分析し、レポートは出力しません。終了コードのみ
avouch --changed # 変更されたファイルと HEAD の追加/削除ビューをコンパクトに表示
avouch --staged # 次のコミットのためにステージングされたファイルのみをレビューします
avouch --all-files # diff だけでなく、対象となるすべての Python ファイルをレビューします
avouch --not-git # ディスク上の対象となるすべての .py ファイルを確認します。 Git リポジトリは必要ありません
avouch --help # すべてのフラグ
レビュー セットは Git によって定義されるため、設定する必要はありません。
呼び出し時間。 --not-git を使用すると、Avouch は Git 要件をスキップし、
現在のルートを調べて見つかった対象となるすべての .py ファイルをレビューします。
代わりにディレクトリを使用します (Git、キャッシュ、仮想環境をスキップします)
ディレクトリ)。レビューを公言:
追跡されたファイルの変更と HEAD ( git diff HEAD --name-only )、および
追跡されていない .py ファイル ( git ls-files --others --exclude-standard )。
削除されたパスと .py 以外のファイルはスキップされます。コミットされた未変更のファイル
出力には決して現れません。生成されたように見えるファイル
( generated.py 、 *_generated.py 、 codegen.py 、 autogen.py 、 … — を参照
src/avouch/utility/is_generated.py ) もスキップされます。
レビュースコープのフラグ --changed 、 --staged 、および --all-files は次のとおりです。
相互に排他的です - 最大 1 つを選択します。出力フラグ --json 、
--verbose および --quiet は、任意のレビュー範囲と自由に組み合わせることができます。
$ 保証する
AVOUCH · 2 ファイル · 4 警告
─────────────────

─────────────
bad.py:1: SCR002: 検出された以外は裸です。代わりに特定の例外をキャッチします。 ValueError: を除く。
│
1 │ def connect(ホスト、ポート、ユーザー、パスワード、データベース、タイムアウト):
│ ^^^^^^^ SCR002
2 │ 試してみましょう:
│
bad.py:1: SCR014: パラメーターが多すぎます (6/5)。関連するパラメータをデータ クラスまたはディクショナリにグループ化します。
│
1 │ def connect(ホスト、ポート、ユーザー、パスワード、データベース、タイムアウト):
│ ^^^^^^^ SCR014
2 │ 試してみましょう:
│
─────────────────────────
ルールにより
SCR002 1 を除くベア
SCR014 パラメータが多すぎます1
─────────────────────────
合格
✓ src/util.py
ヘッダー — AVOUCH · N FILES · W WARN · E ERR : ファイルと重大度ごと
カウントの後にファイルごとの結果が続きます。
調査結果 — 各調査結果はコンパイラ スタイル (ファイル:ライン) でレンダリングされます。
ルール ID と完全なメッセージを含むヘッダー、次に問題のあるコード
フラグ付きの下にグレー表示された行番号とキャレット ^^^^^ が表示された領域
名前 (TTY では青色のルール ID)。
BY RULE サマリー — ルールごとに検出結果がカウントされ、最も一般的なものから順にカウントされます。
カウントは右側に揃えられます。検出結果が存在する場合にのみレンダリングされます。
PASSING グリッド — 準拠ファイル。
[+N more] が多い場合は注意してください。
同一の (コンポーネント、ルール) 結果はファイルごとに重複排除されます。
ヘッダーはすべての検出結果をカウントするため、ルール ID が重複している場合 (SCR004 /
SCR006 ダップ

licate-branch) 行数はヘッダーよりも少なくてもよい
数えます。
$ 保証する
すべてきれいです。
エッジケース
$ cd /tmp/somewhere-without-git
$ 保証する
エラー: Git リポジトリが見つかりません
ヒント: Git リポジトリ内から Avouch を実行するか、 --not-git を使用して Git を使用せずにファイルをレビューします
$ cd ~/fresh-checkout # 例: CIランナー
$ 保証する
エラー: レビューするものはありません
ヒント: HEAD に対して何も変更されていません (CI チェックアウトはクリーンです)。完全なレビューには --all-files を使用してください
色は、標準出力が TTY の場合にのみ出力される ANSI コードです。パイプされた出力は
明白なので、断言します | tee review.log と CI キャプチャは正常に動作します。ランタイム
エラーは標準エラー出力に書き込まれるため、標準出力はパイプ処理用にクリーンなままであり、
--json キャプチャ。終了コードは、レビューがクリーンな場合は 0、レビューがクリーンな場合は 1 です。
調査結果が報告された場合は 2 つ、Avouch を実行できない場合は 2 つです。
avouch --docs は、このコードベースから派生した端末ドキュメントを出力します。
Avouch が行うこと、Git 対応ワークフロー、すべてのルールとその範囲、
設定キーとそのデフォルト、両方の出力形式、および現実的な
例 — その後、レビューを実行せずに 0 を終了します。どこでも機能しますが、
Git リポジトリの外でも。実際の端末では、次のように開きます。
対話型ブラウザ (ヘルプ、移動、メイン画面、終了);標準出力時
パイプ処理されると、代わりにプレーンテキストが出力されます。
自動化と CI の場合、--json はレビューを単一の JSON ドキュメントとして出力します。
標準出力上。人間が判読できるテキストは含まれていません。
--json を保証する
{
「バージョン」: 1 、
"ツール" : " 保証 " ,
「違反」: [
{
"ルール" : " SCR014 " ,
"重大度" : " 警告 " 、
"message" : " パラメーターが多すぎます (6/5)。関連するパラメーターをデータ クラスまたはディクショナリにグループ化します。 " ,
"ファイル" : " buggy.py " ,
"名前" : " 追加 " ,
"種類" : " 関数 " 、
「行」：4
}
]、
「概要」: {
「合計」 : 1 、
「エラー」: 0 、
「警告」: 1 、
「違反のあるファイル」: 1
}
}
各違反にはルール ID (または人間が判読できるルール ID) が含まれます。

ラベルが付いているとき、
検出結果には何もありません)、その重大度、メッセージ、ファイル、コンポーネント名、
その種類 ( func 、 class 、または file )、および検出結果が参照する行
(ファイルレベルの結果の場合は null) — に示されているものと同じコンポーネントと種類
人間のテーブル。 files_with_violations は、個別のファイルの数です。
少なくとも 1 つの違反を含むファイル。
このドキュメントは、自動化のための安定したバージョン管理された契約です: version
スキーマのバージョン (Avouch パッケージのバージョンとは独立)、ツール
エミッターを識別し、同じ入力から常に同じ JSON が生成されます。
— 色、タイムスタンプ、または診断情報が漏れることはありません。終了コードは動作します。
通常モードとまったく同様に、 avouch --json で CI をゲートできます: stdout を解析します
検出結果を確認し、終了ステータスに対応します (クリーン 0、違反 1、
2 エラーを保証します)。
--quiet はまったく同じ分析を実行しますが、レポートは出力されません。だけ
終了コードは結果を通知します (クリーン 0、違反 1、Avouch 2)
エラー)、ステータスのみを必要とするフックとスクリプトに適合します。
エラーが沈黙することはありません。「エラー: Git リポジトリが見つかりません」などのメッセージは引き続き出力され、 --json は引き続きそのドキュメントを出力します。
--verbose 診断は引き続き stderr に保存されます。
Avouch は、すべてのプル リクエストとプッシュで GitHub Actions チェックとして実行できます。
既存のプロジェクトの場合、最小限のワークフローで公開されたパッケージがインストールされます。
そして、PR とプッシュごとにチェックアウト全体をレビューします。
名前 : アヴォーシュ
に:
プルリクエスト:
プッシュ：
ジョブ:
保証する：
実行: ubuntu-最新
権限:
内容：読む
手順:
- 使用:actions/checkout@v6
- 使用:actions/setup-python@v5
付き:
Python バージョン: " 3.12 "
- 名前 : Avouch をインストールする
実行：python -m pip install avouch
- 名前 : ラン・アヴォーシュ
実行: avouch --all-files --json
アクション/チェックアウトはプルリクエストのコードをランナーの作業環境に置きます
ツリー — Avouch は checko するファイルを分析します

提供されるだけで、それ以上は何もありません。
events/setup-python は Python ランタイムを提供します。 Avouch の要件
Python 3.10以降。
python -m pip install avouch は、最新の公開リリースをインストールします。
再現可能な実行のためにバージョン ( avouch==0.3.1 ) を固定します。
avouch --all-files --json は、対象となるすべての .py ファイルをレビューして出力します。
機械可読ドキュメントをジョブログに保存します。必要な権限は読み取りだけです。ワークフローは API 呼び出しを行いません。
デフォルトのレビューセットは変更されたファイルと Git HEAD であるため、新しく
チェックアウトされた作業ツリー - 構築によりクリーン - レビューするものは何もありません。
avouch は error: nothing to review を出力して終了します 2 。同じ
--changed および --staged に適用されます。それらはローカルでのみ意味を持ちますが、
自分の作業ツリーに対して。リポジトリ全体のレビューは、
CIで動作します:
Avouch の終了コードは、ローカルでの場合とまったく同じように CI で動作します。0 はクリーンです。
1 は調査結果が報告されたことを意味し、2 は Avouch を実行できなかったことを意味します。 GitHub
ステップがゼロ以外で終了するとアクションはジョブに失敗するため、 --all-files --json
検出結果のチェックに失敗し、ジョブ ログの JSON ドキュメントに次のことが示されます。
なぜ。 || では何も隠されません。真実 ;すでに存在する調査結果
リポジトリは、修正されるか除外されるまでチェックに失敗します。
avouch.toml のignore_paths。
Avouch リポジトリ自体は .github/workflows/avouch.yml として出荷されます。それを有効にする
リポジトリの「アクション」タブにあるものを選択すると、単独で実行されます。インストールします
リポジトリ独自のソースには pip install -e を使用します。 、つまりコードをテストします
公開されたリリースではなくプル リクエストに含めてから、
--all-files --json を使用してチェックアウトされたリポジトリ全体。
Avouch は、文書化された終了コードを備えたプレーンなコンソール コマンドであるため、あらゆる CI
システムは同じ 3 つの手順で実行できます。
インストール: python -m pip install avouch
実行: 公言します

[切り捨てられた]

## Original Extract

Review the Python you changed, not the Python you inherited. Git-aware AST code review that runs in the seconds before git push. - mukundzha/avouch

GitHub - mukundzha/avouch: Review the Python you changed, not the Python you inherited. Git-aware AST code review that runs in the seconds before git push. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mukundzha
/
avouch
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
132 Commits 132 Commits Folders and files
.github/ workflows .github/ workflows images images src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md avouch.toml avouch.toml pyproject.toml pyproject.toml View all files Repository files navigation
Review the Python you changed, not the Python you inherited.
Avouch is a lightweight, Git-aware static analysis CLI for Python. It asks
Git which files your next commit will touch, parses each changed .py
file with the standard ast module, and reports structural problems
against limits you configure in avouch.toml .
No daemon. No network. No path lists to maintain. Run it in the seconds
before git push , fix what it flags, push.
pip install avouch
cd your-repo
avouch
Table of contents
The review set is the diff, not the repository. Avouch computes the
review set from Git at run time ( git diff HEAD --name-only plus
untracked files). Every finding is attributable to work you are about
to push — never to the legacy you inherited.
Metrics are exact. Parameter counts, nesting depth, and line spans
come from the AST, not regex. If a metric cannot be computed exactly,
Avouch does not claim it.
Errors are data. An unreadable or syntactically broken file becomes
an ERROR entry in the report. One broken file never cancels the
review of the others.
Avouch reviews; it does not gate. The exit code signals the outcome —
0 clean, 1 violations found, 2 Avouch error — but enforcement belongs
in an opt-in interface, not in a tool you run before every push.
The runtime is the standard library. Three git subprocess calls
and
ast / tomllib . No daemon to keep alive; runtime is bounded by the
size of your diff, not your repository.
Requires Python 3.10+ (rules use ast.Match ; configuration uses
tomllib ) and Git on PATH .
pip install avouch
or from source:
git clone https://github.com/mukundzha/avouch.git
cd avouch
pip install -e .
Both register the avouch console script ( avouch.cli:main ).
The interface is one command with a small set of optional flags:
cd your-repo
# ... make a change ...
avouch # human report
avouch --json # one JSON document on stdout
avouch --docs # built-in documentation; no review performed
avouch --version # print the version and exit
avouch --verbose # step-by-step review details on stderr
avouch --quiet # analyze, print no report; exit code only
avouch --changed # compact added/deleted view of changed files vs HEAD
avouch --staged # review only files staged for the next commit
avouch --all-files # review every eligible Python file, not just the diff
avouch --not-git # review every eligible .py file on disk; no Git repo needed
avouch --help # every flag
The review set is defined by Git, so there is nothing to configure at
invocation time. With --not-git , Avouch skips the Git requirement and
reviews every eligible .py file found by walking the current
directory instead (skipping Git, cache, and virtual-environment
directories). Avouch reviews:
tracked files modified vs. HEAD ( git diff HEAD --name-only ), and
untracked .py files ( git ls-files --others --exclude-standard ).
Deleted paths and non- .py files are skipped. Committed, untouched files
never appear in the output. Files that look generated
( generated.py , *_generated.py , codegen.py , autogen.py , … — see
src/avouch/utility/is_generated.py ) are skipped too.
The review-scope flags --changed , --staged , and --all-files are
mutually exclusive — pick at most one. The output flags --json ,
--verbose , and --quiet combine freely with any review scope.
$ avouch
AVOUCH · 2 FILES · 4 WARN
────────────────────────────────────────────────────────────────────────────────
bad.py:1: SCR002: Bare except detected. Catch a specific exception instead, e.g. except ValueError:.
│
1 │ def connect(host, port, user, password, db, timeout):
│ ^^^^^^^ SCR002
2 │ try:
│
bad.py:1: SCR014: Too many parameters (6/5). Group related parameters into a data class or dictionary.
│
1 │ def connect(host, port, user, password, db, timeout):
│ ^^^^^^^ SCR014
2 │ try:
│
────────────────────────────────────────────────────────────────────────────────
BY RULE
SCR002 Bare except 1
SCR014 Too many parameters 1
────────────────────────────────────────────────────────────────────────────────
PASSED
✓ src/util.py
Header — AVOUCH · N FILES · W WARN · E ERR : file and per-severity
counts, followed by the per-file findings.
Findings — each finding renders compiler-style: a file:line
header with the rule id and full message, then the offending code
region with dimmed line numbers and a caret ^^^^^ under the flagged
name (rule id in blue on a TTY).
BY RULE summary — findings counted per rule, most common first,
with counts aligned on the right. Rendered only when findings exist.
PASSING grid — compliant files, compressed to a few lines with a
[+N more] note when there are many.
Identical (component, rule) findings are deduplicated per file — the
header counts every finding, so with overlapping rule IDs (SCR004 /
SCR006 duplicate-branch) the row count can be lower than the header
count.
$ avouch
All clean.
Edge cases
$ cd /tmp/somewhere-without-git
$ avouch
error: no Git repository found
hint: run Avouch from inside a Git repository, or use --not-git to review files without Git
$ cd ~/fresh-checkout # e.g. a CI runner
$ avouch
error: nothing to review
hint: nothing changed vs HEAD (CI checkouts are clean); use --all-files for a full review
Colors are ANSI codes emitted only when stdout is a TTY. Piped output is
plain, so avouch | tee review.log and CI capture work cleanly. Runtime
errors are written to stderr, so stdout stays clean for piping and
--json capture. The exit code is 0 when the review is clean, 1
when findings are reported, and 2 when Avouch cannot run.
avouch --docs prints terminal documentation derived from this codebase —
what Avouch does, the Git-aware workflow, every rule with its scope, every
configuration key with its default, both output formats, and realistic
examples — then exits 0 without running a review. It works anywhere,
even outside a Git repository. In a real terminal it opens as an
interactive browser ( H elp, G o, M ain screen, Q uit); when stdout
is piped it prints the plain text instead.
For automation and CI, --json prints the review as a single JSON document
on stdout, with no human-readable text mixed in:
avouch --json
{
"version" : 1 ,
"tool" : " avouch " ,
"violations" : [
{
"rule" : " SCR014 " ,
"severity" : " WARNING " ,
"message" : " Too many parameters (6/5). Group related parameters into a data class or dictionary. " ,
"file" : " buggy.py " ,
"name" : " extra " ,
"kind" : " func " ,
"line" : 4
}
],
"summary" : {
"total" : 1 ,
"errors" : 0 ,
"warnings" : 1 ,
"files_with_violations" : 1
}
}
Each violation carries the rule id (or a human-readable label when the
finding has none), its severity, the message, the file, the component name,
its kind ( func , class , or file ), and the line the finding refers to
( null for file-level findings) — the same component and kind shown in
the human table. files_with_violations is the number of distinct
files containing at least one violation.
The document is a stable, versioned contract for automation: version
is the schema version (independent of the Avouch package version), tool
identifies the emitter, and the same input always produces the same JSON
— no colors, timestamps, or diagnostics leak in. Exit codes behave
exactly as in normal mode, so avouch --json can gate CI: parse stdout
for the findings and react to the exit status ( 0 clean, 1 violations,
2 Avouch error).
--quiet runs the exact same analysis but prints no report; only the
exit code signals the outcome ( 0 clean, 1 violations, 2 Avouch
error), which makes it fit hooks and scripts that need only the status.
Errors are never silenced: messages such as "error: no Git repository found" still print, --json still emits its document, and
--verbose diagnostics still go to stderr.
Avouch can run as a GitHub Actions check on every pull request and push.
For an existing project, a minimal workflow installs the published package
and reviews the whole checkout on every PR and push:
name : Avouch
on :
pull_request :
push :
jobs :
avouch :
runs-on : ubuntu-latest
permissions :
contents : read
steps :
- uses : actions/checkout@v6
- uses : actions/setup-python@v5
with :
python-version : " 3.12 "
- name : Install Avouch
run : python -m pip install avouch
- name : Run Avouch
run : avouch --all-files --json
actions/checkout puts the pull request's code in the runner's working
tree — Avouch analyzes the files that checkout provided, nothing more.
actions/setup-python provides a Python runtime; Avouch requires
Python 3.10+.
python -m pip install avouch installs the latest published release.
Pin a version ( avouch==0.3.1 ) for reproducible runs.
avouch --all-files --json reviews every eligible .py file and prints
the machine-readable document to the job log. permissions: contents: read is the only permission needed — the workflow makes no API calls.
The default review set is files changed vs. Git HEAD , so a freshly
checked-out working tree — clean by construction — has nothing to review:
avouch would print error: nothing to review and exit 2 . The same
applies to --changed and --staged ; they only make sense locally,
against your own working tree. Whole-repository review is the mode that
works in CI:
Avouch's exit code behaves in CI exactly as it does locally: 0 is clean,
1 means findings were reported, 2 means Avouch could not run. GitHub
Actions fails a job when a step exits non-zero, so --all-files --json
fails the check on any finding, and the JSON document in the job log shows
why. Nothing is hidden with || true ; findings already present in the
repository fail the check until they are fixed or excluded with
ignore_paths in avouch.toml .
The Avouch repository itself ships .github/workflows/avouch.yml ; enable it
in the repository's Actions tab and it runs on its own. It installs
the repository's own source with pip install -e . , so it tests the code
in the pull request rather than a published release, then reviews the
whole checked-out repository with --all-files --json .
Avouch is a plain console command with a documented exit code, so any CI
system can run it with the same three steps:
Install: python -m pip install avouch
Run: avouch

[truncated]
