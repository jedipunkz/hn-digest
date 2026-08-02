---
source: "https://github.com/jonhardwick-spec/ccn"
hn_url: "https://news.ycombinator.com/item?id=49141107"
title: "Show HN: Nuking the crap Claude left in the codebase – CCN"
article_title: "GitHub - jonhardwick-spec/ccn: Interactive comment stripper that cannot eat your code. acorn for JS, tree-sitter for the rest. 9 languages, one tool. Strips AI comments, dead comments, TODOs, and Co-Authored-By spam from JS, TS, TSX, Python, Java, Kotlin, Rust, C, C++. · GitHub"
author: "jonhardwickspec"
captured_at: "2026-08-02T05:18:04Z"
capture_tool: "hn-digest"
hn_id: 49141107
score: 1
comments: 0
posted_at: "2026-08-02T04:31:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Nuking the crap Claude left in the codebase – CCN

- HN: [49141107](https://news.ycombinator.com/item?id=49141107)
- Source: [github.com](https://github.com/jonhardwick-spec/ccn)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T04:31:26Z

## Translation

タイトル: HN を表示: クロードがコードベースに残したクソを核攻撃 – CCN
記事のタイトル: GitHub - jonhardwick-spec/ccn: コードを読み込むことができない対話型コメント ストリッパー。 JS にはどんぐり、残りは木番です。 9 つの言語、1 つのツール。 JS、TS、TSX、Python、Java、Kotlin、Rust、C、C++ から AI コメント、デッドコメント、TODO、および共同作成者によるスパムを除去します。 · GitHub
説明: コードを読み込むことができない対話型コメント ストリッパー。 JS にはどんぐり、残りは木番です。 9 つの言語、1 つのツール。 JS、TS、TSX、Python、Java、Kotlin、Rust、C、C++ から AI コメント、デッドコメント、TODO、および共同作成者によるスパムを除去します。 - ジョンハードウィック仕様/ccn
HN テキスト: CCN コードベースからコメントを取り除き、コメントのみを行います。これはコードを読み込む可能性のある正規表現解析ではありません。実際にはリストを作成して 2 回チェックします。実際の開発者である私は、自分のコードベースにコメントを数バイト残して忘れてしまうことがよくあります。AI モデルはメガバイト単位でメモリを持たず、文字通りコメントを IT の形式として使用することがあります。このツールを使用して、コードベースから文字通りのゴミを取り除くか、非常に便利で、エンドツーエンドでテストされ、2,700 回の反復に対して検証されているレビューに使用します (気にする必要はありませんが、xfce デスクトップを備えた debian trixie 上でも動作します。それが重要です。公開する前にコードベースからリッピングした、エンド ユーザーが心配する何千ものテストではありません) ストールマンの悪名高い言葉にあるように... ハッピーハッキング ;)

記事本文:
GitHub - jonhardwick-spec/ccn: コードを読み込むことができない対話型コメント ストリッパー。 JS にはどんぐり、残りは木番です。 9 つの言語、1 つのツール。 JS、TS、TSX、Python、Java、Kotlin、Rust、C、C++ から AI コメント、デッドコメント、TODO、および共同作成者によるスパムを除去します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のアカウントでアカウントを切り替えました

bまたはウィンドウ。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ジョンハードウィック仕様
/
CCN
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット bin bin lib lib node_modules node_modules .gitignore .gitignore LICENSE.md LICENSE.md README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
★ クラプコメンヌッカー ★
コードを読み込むことができない対話型コメントストリッパー
やあ、今度はどのようなコード ベース ファイル / タイプのファイルからゴミを取り出すのですか
| |
|ゴミ箱 |
|のみ |
|_______|
//_______________________________________
//| Anthropic の Claude Code との共著 |
//|____________________________________________|
___
\
\O
\/|\
| \
／＼
／＼
> [x] javascript [検出: 3 ファイル] [acorn] .js .cjs .mjs .jsx
[x] python [検出: 1 ファイル] [tree-sitter] .py
[ ] kotlin [ツリーシッター] .kt .kts
...
ccn はソース ファイルのディレクトリを調べてコメントを削除します。一度に 1 ブロックずつ、
クリアされた画面で、Y/N を決定します。 HTML、MD、JSON、TXT、その他すべて
認識しないし、触ることもありません。
正規表現ではなく実際のパーサーを背後に持つ言語を含む 9 つの言語を話します
推測します。 JavaScript がデフォルトです。他のすべての言語は同じ考えを固定しています
別のパーサーに。
提案全体は 1 つの文です。「コードを食べることはできません」です。書き込みゲート
ファイルを再解析し、単一のトークンとコメントを移動する編集をバウンスします。
トークンではありません。コードはそのまま残ります。
出典: link タグ:javascripttypescripttsx
出典: link タグ:pythonjavakotlinRust
c c++
それぞれがそのバリアントとスキッドをカバーしています。JavaScript は js 、 cjs 、 mjs 、
esm 、 jsx 。 C++ は cxx 、 cc 、 hpp 、 hh をカバーします。 kotlin は kt 、 kts をカバーします。
JavaScript (および cjs 、

mjs 、 esm 、 jsx ) は、 JS である acorn に基づいています
エンジン独自のパーサー ファミリ。それが神託であり、それが最も鋭いのです。 1つおき
この言語は、 Tree-sitter に基づいており、同じ 3 つのことを提供します。
実際の AST: コメント スパン (つまり、文字列内の // または Python 文字列内の #
リテラルがコメントと間違われることはありません）、有効性チェック、およびリーフトークン
不変。
--lang を使用して言語を選択するか、ccn にディレクトリから言語を検出させます。もし
ディレクトリが混在していて TTY を使用している場合、次のように尋ねられます。
コメントを削除するコード言語を入力します (JavaScript、CJS、ESM、TypeScript、Python、Java、Kotlin、Rust、C、C++ など)。
非 JS 言語は、最初の実行時に文法が必要です。 ccnはそう言って尋ねます
何かを引っ張る前に:
Python のコメントを削除するには、Python AST パーサー (Python) が必要です。ダウンロード？はい/いいえ
1 回のフェッチは ~/.config/ccn/grammars にキャッシュされ、その後永久にオフラインになります。
npm install -g commentnuker
または、クローンを作成してその場で実行します。 Deps は、最初の起動時に自動インストールされます。
欠落しています (JS の場合は acorn、その他すべての場合は web-tree-sitter):
git clone < このリポジトリ > ccn
cd ccn
ノード bin/ccn ./src # js、自動検出
node bin/ccn ./src --lang c++ # c++ コメントを削除します (はい、c ではなく "c++" が解決されます)
1 回限りのフェッチの後、ネットワークがありません。 JS はその後永久にオフラインで実行されます。それぞれの新しい
非 JS 言語は文法を一度プルすると、その後は永久にオフラインになります。
Linux と macOS は安定しています。 Windows...「Windows では動作しないかもしれません、皆さん、私はどう思いますか？」
試してみました🥀」。パスは path.join され、ホーム ディレクトリは USERPROFILE に戻ります。
ノードは直接生成されるため (シェルのいたずらはありません)、理論上は次のようになります。
走る。そうでない場合は、GitHub に苦情を投稿してください。修正します。
ただ実行してください。引数や覚えておくべきコマンドはなく、すぐに始められます。
メニューを表示し、残りの手順を説明します。
ccn # ハブ: 歴史を核攻撃する、要約する、または消去する
ccn ./

src # ハブをスキップし、言語メニューから開始します
ccn ./src --lang python # ハブ + 言語をスキップして、quip + ピッカーに直接移動します
ccn --langs # サポートされているすべての言語をリストします
ccn ./src --dry-run --remember # リハーサル、選択肢を記録、何も書かない
ccn ./src --demo --limit 5 # auto-Y プレビュー、ファイルあたり 5 ブロック
ハブ (引数なし)
何も指定せずに ccn を実行すると、使用法エラーではなくメイン メニューが表示されます。
★クズコメントヌーカー★
私たちは何をしているのですか？
> ゴミを取り除く（コードベースを破壊する）
要約 -- 以前に削除されたコードベースとコメント
リキャップストレージを削除してこの機能を無効にする
出る
ゴミ箱を選択すると、ディレクトリと言語に直接移動します
以下の流れで説明します。要約は歴史です。要約ストレージ ワイプを削除する
録音をオフにします (同じスロットで録音を再びオンにします)。
すべての本物の核兵器（予行演習や安全ゲートが拒否したファイルではありません）はログに記録されます
~/.config/ccn/recap.json に。ハブから要約を開くと、次の内容が得られます。
以前に、これらのコードベースを無効化しました
1 つを選択して、一緒に攻撃したコメントを表示します
> 私のアプリ (3 核、JavaScript、2026 年 8 月 1 日)
cli (12 核、JavaScript、2026-07-30)
メインメニューに戻る
1 つのコメントにドリルダウンすると、削除したすべてのコメントがリストされます。ファイル + 行に入力してください。
全文読んでください。これは元に戻すことではなく、遡ることです。コメントはすでに削除されています
コードから (書き換えられた各ファイルの横に .bak が残されました)。
要約はデフォルトでオンになっています。ハブの切り替えによりストレージが消去され、録画が切り替わります
オフにするか、オンに戻します。最後の 200 セッションが保持され、最も古いセッションが削除されます。
インタラクティブ核兵器はハブの後に 4 つのスクリーンを通過します。 (ディレクトリを
最初の引数はハブとスクリーン 0 をスキップします。画面 0 と 1 をスキップするには --lang を渡します。)
画面 0、ディレクトリ メニュー。コードベースはどこにありますか?可能性のあるディレクトリのリスト
(現在のディレクトリ、デスクトップ、および実際にコードを含むサブディレクトリ)、
プラスA型

手動でパスを設定するオプション。矢印キーで選択するか、自分で入力します。
コマンドラインで dir を渡す場合はスキップされます。
画面 1、言語メニュー。 1 つまたは複数の言語を選択します。検出された言語
緑色で強調表示され、事前にチェックされているため、混合コードベースは 1 つのコードベースで処理されます。
走る。各コメントはパーサーによって独自の言語に合わせてゲートされるため、JS ファイル
同じパスで acorn と Tree-sitter を介した Python ファイルを通過します。ある
すべてを選択し、スペースで切り替え、Enter で確定します (少なくとも 1 つ必要です)。
やあ、今度はどのコードベース ファイル / タイプのゴミ箱を取り出すのですか
| |
|ゴミ箱 |
|のみ |
|_______|
//_______________________________________
//| Anthropic の Claude Code との共著 |
//|____________________________________________|
___
\
\O
\/|\
| \
／＼
／＼
> [x] javascript [検出: 3 ファイル] [acorn] .js .cjs .mjs .jsx
[x] python [検出: 1 ファイル] [tree-sitter] .py
[ ] kotlin [ツリーシッター] .kt .kts
...
画面 2、口語 + メソッド + 続行。言語についての一言
あなたが選択した (Java が本当の伝承を取得し、残りは同じ姿勢を取得します)、そして正確に
そのエンジンでコメントがどのように削除され、次に進むのでしょうか?はい/いいえ。非JS
言語は、文法がまだキャッシュされていない場合、ここに文法を取り込むように要求します。
画面 3、コメントごとのピッカー。すべてのコメント ブロック (連続した実行)
// 行、1 つの /* */ 、または # run) は、独自のクリアされた画面を取得します。
緑 これはここに属する可能性があります -- ブロックはメソッド、クラス、または
const 宣言、 @marker を運ぶ、または TODO 、 FIXME 、 HACK 、または
注意。それでもなお、「とにかく核を使うのか？」と尋ねる。はい/いいえ。あなたが決めてください。ただ大声で警告するだけです。
赤 大きなブロック、つまり複数行のブロックが見つかりました。
red Yo は他のすべてを攻撃するコメントを見つけました。
y または Enter で核攻撃、n で維持、m でシミュレーションメニュー、ctrl-c で解除
(最初にメモリを節約します)。
任意のコメントで m を押すと、クラスターにドロップされます

ビュー。のすべてのコメント
現在使用しているものと >= --sim cosine-similar のコードベースがグループ化されます
その下に、全員が違反する共有ルールがリストされ、あなたがそのルールを決定します。
グループ全体を一度に:
✦ SIMULURITY ✦ ベースライン主導のクラスター
あなたのエントリ:
「TODO これは後でリファクタリングしてください、ハッキー」
コードベース全体で他の 2 つのコメントが 0.75 以上類似しています。
これらはすべて同じルールに違反しています: TODO、単一行、末尾
> 全部削除 (コメント 3 件)
どれを選ぶか
ベースラインのみを設定し、レビューを続ける
コードベースに残っているすべてのものを核攻撃する (コメント 47 件)
マニュアルに戻る
Strip em all : クラスター全体を 1 回の動作で破壊します。
どれを選択するか: 複数選択し、個々のコメントをスペースで切り替えます。
ベースラインのみを設定: このコメントを参照として保持し、次のページに戻ります。
マニュアルで一つ一つレビューします。
残りのすべてを核攻撃します: に残っているすべてのコメントを直接核攻撃します。
コードベース。あなた次第。
マニュアルに戻る : コメントごとのピッカーに戻ります。
ゲートは、コメントがどのように選択されたかに関係なく、書き込みのたびに実行されます。
クラスター パスは、手動パスと同様にコードを使用できません。
旗
それは何をするのか
--lang <名前>
言語を選択します ( js 、 python 、 c++ 、 java 、 ...)。 jsがデフォルトです。画面1をスキップします
--langs
サポートされているすべての言語とそのエンジンをリストし、終了します
--覚えておいてください
すべての選択内容を ~/.config/strip-review/memory.json に保存します
--ドライラン
レビューのみ、書き込みはしない (選択肢は --remember で記録されたまま)
--デモ
手動プロンプトごとに auto-Y、TTY は不要 (メニューも自動解決)
--リミットN
ファイルごとに確認されるキャップ ブロック
--sim 0.75
シミュレーションクラスターのコサイン類似度バー
--SIMなし
クラスタリングをオフにします ([m] メニューが開いたままになりますが、何も一致しないだけです)
--flay MS
ファイル間の遅延 (デフォルトは 50)
--cdelay MS
ブロックごとの遅延 (デフォルトは 150)
記憶とそっくりさん
--remember を使用すると、選択肢 y に完全に一致します

すでに作成したものは自動適用されます
後の実行時（脳）。それは今でもそれができたことを示しています。 「属するかもしれない」ブロック
自動的に適用されることはなく、常に尋ねられます。
したがって、ワークフローは次のとおりです。 --remember を使用して予行演習を行い、好みを教えてから、
実際の実行はほとんど自動です。ゲートがあるため、コードを食べることはできません。
メモリの内容に関係なく、毎回実行されます。
つまり、ソースコードはここにあります、おい、信用できないとは言わないでください
コメントを無効にする... あなたは、コメントを無効にする何かについて偏執的です
FA プロキシなしで Claude Code を実行すると、あなたが疑心暗鬼になっていることがすでにわかります。
監視されたり、暴行を受けたりしても、手遅れになったときには、あなたはすでにそうなっている、あなたがそうするためにやったのです
あなた自身、それはすべてあなたが企業を信頼していたからです。
5 つの法務チームの後ろに隠れているボリス・チェルニーを信用するつもりですか?
あるいは、あなたは、あなたが話すことができる男が作ったコメントストリッパーを信頼する必要があります
インスタグラム @larp_init ?
そう思ったなら、AI があなたの中にメガバイト残した知恵遅れのコメントを核攻撃しましょう
コードベース、どういたしまして。
問題、苦情、機能リクエスト: GitHub リポジトリ。エンドユーザーが入力できるのは、
GitHub xD での苦情、それがその目的です。
コメントリムーバー ストリップコメント コメントストリッパー 削除コメント
削除-コメント コメント-クリーナー コード-クリーンアップ コード-衛生

[切り捨てられた]

## Original Extract

Interactive comment stripper that cannot eat your code. acorn for JS, tree-sitter for the rest. 9 languages, one tool. Strips AI comments, dead comments, TODOs, and Co-Authored-By spam from JS, TS, TSX, Python, Java, Kotlin, Rust, C, C++. - jonhardwick-spec/ccn

CCN Strips comments out of codebases, and comments ONLY , this isn't a regex parse that could potentially eat your code, no this actually makes a list and checks it twice . I, as a real developer, often leave bytes of comments in my own codebases then forget about them, AI models leave MEGABYTES and has NO memory, sometimes literally using comments as a FORM OF IT. Use this tool to clean your codebases of literal trash, or use it to review, very useful, end to end tested and verified against 2,700 iterations (not that you care or anything, it works, even on debian trixie with xfce desktop, and that's what matters, not the thousands of tests I ripped from the codebase before publishing it that end users dgaf about) As Stallman infamously says... Happy hacking ;)

GitHub - jonhardwick-spec/ccn: Interactive comment stripper that cannot eat your code. acorn for JS, tree-sitter for the rest. 9 languages, one tool. Strips AI comments, dead comments, TODOs, and Co-Authored-By spam from JS, TS, TSX, Python, Java, Kotlin, Rust, C, C++. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
jonhardwick-spec
/
ccn
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit bin bin lib lib node_modules node_modules .gitignore .gitignore LICENSE.md LICENSE.md README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
★ C R A P C O M M E N T N U K E R ★
interactive comment stripper that cannot eat code
Igh bro, now, what code base files / type are we taking the trash out of
| |
| TRASH |
| ONLY |
|_______|
//_________________________________________
//| Co-Authored by Anthropic's Claude Code |
//|________________________________________|
___
\
\ O
\/|\
| \
/ \
/ \
> [x] javascript [detected: 3 files] [acorn] .js .cjs .mjs .jsx
[x] python [detected: 1 file] [tree-sitter] .py
[ ] kotlin [tree-sitter] .kt .kts
...
ccn walks a directory of source files and nukes comments. One block at a time,
on a cleared screen, you decide Y/N . HTML, MD, JSON, TXT, and everything it
does not recognize is never touched.
It speaks 9 languages, the ones with a real parser behind them, not a regex
guess. JavaScript is the default. Every other language is the same idea bolted
onto a different parser.
The whole pitch is one sentence: it cannot eat your code. a write gate
re-parses the file and bounces any edit that moves a single token, and comments
aren't tokens. your code stays put.
javascript typescript tsx
python java kotlin rust
c c++
Each one covers its variants and skids: javascript covers js , cjs , mjs ,
esm , jsx . c++ covers cxx , cc , hpp , hh . kotlin covers kt , kts .
JavaScript (and cjs , mjs , esm , jsx ) is grounded in acorn , the JS
engine's own parser family. that's the oracle, and it's the sharpest. Every other
language is grounded in tree-sitter , which gives the same three things from a
real AST: comment spans (so a // inside a string or a # in a Python string
literal is never mistaken for a comment), a validity check, and a leaf-token
invariant.
Pick the language with --lang , or let ccn detect it from the directory. If
the directory is mixed and you are on a TTY, it asks:
Type the code language to nuke comments from (eg, JavaScript, CJS, ESM, TypeScript, Python, Java, Kotlin, Rust, C, C++):
Non-JS languages need their grammar on the first run. ccn says so and asks
before it pulls anything:
Nuking python comments needs the python AST parser (python). Download? Y/N
One fetch, cached in ~/.config/ccn/grammars , offline forever after.
npm install -g commentnuker
Or clone and run it in place. Deps auto-install on first launch if they are
missing (acorn for JS, web-tree-sitter for everything else):
git clone < this repo > ccn
cd ccn
node bin/ccn ./src # js, auto-detected
node bin/ccn ./src --lang c++ # nuke c++ comments (yes, "c++" resolves, not c)
No network after the one-time fetches. JS runs offline forever after; each new
non-JS language pulls its grammar once, then offline forever after.
Linux and macOS, solid. Windows... "Might not work for windows, idek guys I
tried 🥀". Paths are path.join 'd, the home dir falls back to USERPROFILE ,
and node is spawned directly (no shell shenanigans), so on paper it should
run. If it doesn't, put in a complaint on the GitHub, I'll fix it.
Just run it. No arguments, no commands to remember, it drops you straight into
the menu and walks you through the rest:
ccn # the hub: nuke, recap, or wipe history
ccn ./src # skip the hub, start at the language menu
ccn ./src --lang python # skip hub + language, straight to the quip + picker
ccn --langs # list every supported language
ccn ./src --dry-run --remember # rehearse, record choices, write nothing
ccn ./src --demo --limit 5 # auto-Y preview, 5 blocks a file
The hub (no args)
Run ccn with nothing and you land on the main menu, not a usage error:
★ Crap Comment Nuker ★
what are we doing?
> take out the trash (nuke a codebase)
recap -- previously nuked codebases & the comments
delete recap storage & disable this feature
exit
Pick take out the trash and you go straight into the directory + language
flow described below. recap is the history. delete recap storage wipes
it and flips recording off (same slot flips it back on).
Every real nuke (not a dry run, not a file the safety gate refused) gets logged
to ~/.config/ccn/recap.json . Open recap from the hub and you get:
Previously, you nuked these codebases
pick one to see the comments we nuked together
> my-app (3 nuked, javascript, 2026-08-01)
cli (12 nuked, javascript, 2026-07-30)
back to main menu
Drill into one and every comment you removed is listed, file + line, enter to
read it in full. It is a look-back, not an undo, the comments are already gone
from your code (a .bak was left next to each rewritten file).
Recap is on by default. The hub toggle wipes the storage and turns recording
off, or turns it back on. The last 200 sessions are kept, oldest fall off.
Interactive nukes go through four screens after the hub. (Pass a dir as the
first arg to skip the hub and screen 0; pass --lang to skip screens 0 and 1.)
Screen 0, the directory menu. Where's the codebase? A list of likely dirs
(the current dir, your Desktop, and any subdirs that actually contain code),
plus a type a path manually option. Arrow keys to pick, or type one yourself.
Skipped if you pass a dir on the command line.
Screen 1, the language menu. Pick one language or many. Detected languages
are highlighted in green and pre-checked, so a mixed codebase is handled in one
run. Each comment gets gated by the parser for its own language, so a JS file
goes through acorn and a Python file through tree-sitter in the same pass. a
selects all, space toggles, enter confirms (needs at least one).
Igh bro, now, what codebase files / type are we taking the trash out of
| |
| TRASH |
| ONLY |
|_______|
//_________________________________________
//| Co-Authored by Anthropic's Claude Code |
//|________________________________________|
___
\
\ O
\/|\
| \
/ \
/ \
> [x] javascript [detected: 3 files] [acorn] .js .cjs .mjs .jsx
[x] python [detected: 1 file] [tree-sitter] .py
[ ] kotlin [tree-sitter] .kt .kts
...
Screen 2, the quip + the method + proceed. A one-liner about the language
you picked (Java gets the real lore, the rest get the same posture), then exactly
how the comments get stripped for that engine, then proceed? Y/n . Non-JS
languages ask to pull their grammar here if it is not cached yet.
Screen 3, the per-comment picker. Every comment block (a contiguous run
of // lines, one /* */ , or a # run) gets its own cleared screen:
green This might belong here -- the block sits above a method, class, or
const declaration, carries an @marker , or has a TODO , FIXME , HACK , or
NOTE . It still asks: nuke anyway? Y/N . You decide. It just warns louder.
red Found a big one -- a multi-line block.
red Yo found a comment to nuke -- everything else.
y or enter to nuke, n to keep, m for the simulurity menu, ctrl-c to bail
(it saves memory first).
Hit m on any comment and you drop into the cluster view. Every comment in the
codebase that is >= --sim cosine-similar to the one you are on gets grouped
under it, the shared rules they all break get listed, and you decide for the
whole group at once:
✦ SIMULURITY ✦ baseline-driven cluster
your entry:
"TODO refactor this later, hacky"
2 other comments across the codebase are >=0.75 similar to it.
these ALL break the same rules: TODO, single-line, trailing
> strip em all (3 comments)
pick which ones
set baseline only, keep reviewing
nuke EVERYTHING remaining in the codebase (47 comments)
back to manual
strip em all : nuke the whole cluster in one move.
pick which ones : multi-select, toggle individual comments with space.
set baseline only : keep this comment as the reference and go back to
manual one-by-one review.
nuke EVERYTHING remaining : straight nuke every comment left in the
codebase. up to you.
back to manual : return to the per-comment picker.
The gate runs on every write regardless of how the comment got picked, so the
cluster path cannot eat code any more than the manual path can.
flag
what it does
--lang <name>
pick the language ( js , python , c++ , java , ...). js is default. skips screen 1
--langs
list every supported language and its engine, then exit
--remember
persist every choice to ~/.config/strip-review/memory.json
--dry-run
review only, never write (choices still recorded with --remember )
--demo
auto-Y every manual prompt, no TTY needed (also auto-resolves the menus)
--limit N
cap blocks reviewed per file
--sim 0.75
cosine similarity bar for the simulurity cluster
--no-sim
turn clustering off (the m menu still opens, just matches nothing)
--fdelay MS
delay between files (default 50)
--cdelay MS
delay per block (default 150)
Memory and look-alikes
With --remember , an exact match to a choice you already made is auto-applied
on later runs (the brain). It still shows you it did it. "Might belong" blocks
are never auto-applied, they always ask.
So the workflow is: dry-run with --remember once to teach it your taste, then
the real run is mostly automatic. It still cannot eat code, because the gate
runs every time regardless of what memory says.
I mean the source code is RIGHT HERE dawg, don't tell me you can't trust a
comment nuker... You being paranoid about something that nukes comments while
running Claude Code without the FA proxy already tells us, you're paranoid about
being spied on or ratted, when it's too late, you already are, you did it to
yourself, all because you trusted a corporation.
You gunna trust Boris Cherny who hides behind 5 legal teams?
Or you gunna trust the comment stripper made by a guy you can talk to on
instagram @larp_init ?
Ight thought so, go nuke those retarded comments AI left megabytes of in your
code base, you're welcome.
Issues, complaints, feature requests: the GitHub repo. End users can put in
complaints on the GitHub xD, that's what it's there for.
comment-remover strip-comments comment-stripper remove-comments
delete-comments comment-cleaner code-cleanup code-hygien

[truncated]
