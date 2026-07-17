---
source: "https://github.com/Anode1/mincdp"
hn_url: "https://news.ycombinator.com/item?id=48948997"
title: "Mincdp – Give Claude eyes and hands in a browser, in ~410 lines of C"
article_title: "GitHub - Anode1/mincdp: Minimal, dependency-free Chrome DevTools Protocol (CDP) client in C and Java — drive headless Chrome for tests without Selenium or chromedriver. · GitHub"
author: "siberean"
captured_at: "2026-07-17T17:01:36Z"
capture_tool: "hn-digest"
hn_id: 48948997
score: 2
comments: 0
posted_at: "2026-07-17T16:08:04Z"
tags:
  - hacker-news
  - translated
---

# Mincdp – Give Claude eyes and hands in a browser, in ~410 lines of C

- HN: [48948997](https://news.ycombinator.com/item?id=48948997)
- Source: [github.com](https://github.com/Anode1/mincdp)
- Score: 2
- Comments: 0
- Posted: 2026-07-17T16:08:04Z

## Translation

タイトル: Mincdp – ブラウザーにクロードの目と手を表示する (C 言語の約 410 行)
記事のタイトル: GitHub - Anode1/mincdp: C および Java の最小限で依存関係のない Chrome DevTools Protocol (CDP) クライアント — Selenium や chromedriver を使用せずにテスト用にヘッドレス Chrome を駆動します。 · GitHub
説明: C および Java の最小限で依存関係のない Chrome DevTools Protocol (CDP) クライアント — Selenium や chromedriver を使用せずにテスト用にヘッドレス Chrome を駆動します。 - アノード1/mincdp

記事本文:
GitHub - Anode1/mincdp: C および Java の最小限で依存関係のない Chrome DevTools Protocol (CDP) クライアント — Selenium や chromedriver を使用せずにテスト用にヘッドレス Chrome を駆動します。 · GitHub
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
アノード1
/
mincdp
公共
通知
サインインする必要があります

o 通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
16 コミット 16 コミット c c docs docs java java レガシー レガシー テスト テスト .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md エージェント.sh エージェント.sh デモ.sh デモ.sh page.html page.html すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Selenium またはバージョンが一致したものを使用せずに、実際のヘッドレス Chrome をテスト用に駆動します
クロムドライバー。 Chrome DevTools を使用する、依存関係のない 2 つの小さなクライアント
プロトコル (CDP) は、WebSocket 経由でブラウザに直接送信されます (使用言語で)。
スタック:
C : c/cdp.h 、単一のヘッダーのみのライブラリ (~410 行、libc + POSIX)
ソケットのみ)。
Java : java/Cdp.java 、単一のファイル (JDK java.net.http のみ、何もなし)
ベンダーに）。
同じアイデア、同じ小さなコマンド セット、同じページに対するそれぞれ 1 つのデモンストレーション。
上の画像は mincdp 自体によって生成されました。ページに入力されたデモです。
Enter を押して、この駆動状態を cdp_screenshot でキャプチャしました
( Page.captureScreenshot )。
CDP は、Puppeteer と Playwright が内部で使用するワイヤー プロトコルです。それを話す
直接的には、chromedriver の仲介者やベンダーへの外部依存がないことを意味します。
Chrome とバージョンを一致させます。内部ページの固定セットの場合、ヘッドレス、1 つ
ブラウザ、それは良い取引です。クライアントは座って読めるほど小さいので、
Chrome が破壊的な CDP 変更をリリースした場合は、そのファイルを修正するのではなく、その 1 つのファイルを修正する必要があります。
上流のドライバーのリリースを待ちます。
これは意図的に一般的なフレームワークではありません。それはコマンドだけを知っています
クリックアンドアサートテストのニーズ: ナビゲート、JS の評価、テキストの入力、キーの押し。もし
クロスブラウザー (Firefox/Safari) またはリッチ インタラクション API が必要な場合は、使用します。
劇作家。これは、ミニマリズムと依存関係ゼロの場合にのみその役割を果たします。
ニーズ

PATH 上のヘッドレス Chrome/Chromium (およびcurl、および cc / javac )。
make ut # 回帰スイート全体 (「テスト」を参照)
make demo-c # c/demo をビルドし、page.html に対して実行します
make demo-java # java/Demo をビルドし、page.html に対して実行します。
デモを作成 # 両方
ショット # スクリーンショット page.html を作成して表示できるようにします
make Agent GOAL="..." # クロードにページを目標に向かって進めさせます (API キーが必要です)
各デモでは、一連のアサーションとデモ (N 行が合格、0 行が失敗) が出力されます。それ
最初にページの内部（タイトル、属性、テキスト、計算されたスタイル、
geometry) を入力し、実際の入力パスを駆動します。「mincdp works」と入力し、Enter キーを押します。
ページがそれをエコーすることをアサートします。すべての CDP フレームを表示するには、CDP_DEBUG=1 を設定します。
make ut は 2 種類の完全な回帰を実行するため、テスト セットには両方の機能があります
そして目：
codeut -- クライアントの純粋なヘルパーの「通常の」ブラウザ不要の単体テスト
(base64、JSON エスケープ)、それぞれいくつかの例。 Chrome は必要ないので、いつでも
実行します (tests/codeut.c)。
uiut -- 実際のヘッドレス Chrome を駆動するクライアント。次の 2 種類があります。
ページ内部 + ハンド -- デモでは、レンダリングされた DOM (存在、
属性、テキスト、計算されたスタイル、ジオメトリ) を Runtime.evaluate 経由で取得し、
インタラクション (入力して Enter キーを押し、DOM が変更されたことをアサート) し、スクリーンショットを撮ります。
プロトコル上の駆動状態 ( cdp_screenshot -> Page.captureScreenshot )
有効な PNG ( c/demo.c 、 java/Demo.java ) をアサートします。
エージェント向けの目 -- testing/shot.sh はスタンドアロン シェル バージョンです。
Chrome の --screenshot を使用して、クライアント コードなしでロードされた新しいページのスクリーンショットを取得します
必要です。これは、デモのプロトコルのスクリーンショット (駆動された状態をキャプチャしたもの) を補完します。
状態）。 PNG を開くか、エージェントに読み取ってもらいます。
各レイヤーは PASS / FAIL / SKIP を報告します (欠落しているツールチェーン SKIP、出口 77)。
スイートは、スキップされていないレイヤーが失敗した場合にのみ失敗します。
mincd

p回帰:
codeut (単位) PASS
ウイットハンド：C PASS
uiut ハンド: Java PASS
uiut 目: ショット PASS
デモはインタラクション テストを兼ねています。デモ-C / デモ-Java を実行させます
自分たちで。
Chrome を起動します。クライアントが添付します。デモ.sh はその配線を行います (
headless Chrome で --remote-debugging-port を使用し、デモを実行し、分解します)。
同じコントラクトで両方のデモを呼び出すため、1 つのランチャーと 1 つの page.html
両方の言語に対応します。
PROG... 127.0.0.1 <chrome-port> file://<abs>/page.html
エージェントの目と手
スモーク テストを作成するのと同じ 2 つのプリミティブ、スクリーンショット (目)、
入力呼び出し (ハンド) は、LLM がブラウザを駆動するために必要な表面全体です。
c/agent.c はそのループであり、他には何もありません。見て、考えて、行動して、の繰り返し。
見る。 cdp_screenshot は、前のアクションの状態に関係なく、ライブ ページをキャプチャします。
まで運転しました。
考える。 PNG と 1 行のゴールは、curl 経由で Claude API に送られます (
ここでの生のソケットの 1 ホップでは実行できないのは、 api.anthropic.com への TLS です)。モデル
たった 1 つのアクションで応答します。
活動。クライアントはページに対するその 1 つのアクションを再生し、ループします。
モデルはスクリーンショットを取得し、次のいずれかの 1 行で応答します。
クライアントと同じ部分文字列トリック: 応答は 1 つの短い行なので、
応答からのアクションは、JSON パーサーではなく、単一の "text":"..." 検索です。
エクスポート ANTHROPIC_API_KEY=sk-ant-...
make Agent GOAL="ボックスに「mincdp works」と入力して Enter キーを押します"
page.html に対する実行は次のようになります。
ステップ 1: JS document.getElementById('q').focus()||true
ステップ 2: TYPE mincdp は機能します
ステップ 3: キー入力
ステップ 4: 完了すると、ボックスに「echo: mincdp works」と表示されます。
各ステップは完全なループです。モデルは新しいスクリーンショットを見て、次のステップを選択します。
キーストロークが実行され、mincdp がそれを再生し、スクリーンショットが表示された時点で DONE で閉じました。
目標は達成されました。
それは意図的にスケッチです: 1 つ

ターンごとのアクション、部分文字列解析、再試行なし、
スピードを考えていない。そのミニマリズムがポイントです。読めるくらい小さいです
モデルにブラウザを動作させる前に、エンドツーエンドで確認してください。デモと同様にスキップします
(exit 77) Chrome、curl、または ANTHROPIC_API_KEY が見つからない場合。参照
c/agent.c およびagent.sh 。
両方のクライアントは同じ形状を公開します (C 名が示されています。Java はキャメルケースの双子です)。
応答は ID で照合され、JSON ではなく対象の部分文字列チェックで読み取られます。
パーサー。これが、両方のクライアントの依存関係をなくす秘訣です。すべてのコマンドです。
ブール値、ショート ACK、または 1 つのフラット フィールド (スクリーンショットの Base64) を返すため、
"value":true /ExceptionDetails / "data":"..." を探すだけで十分です。あ
深くネストされた JSON を返すコマンドには実際のパーサーが必要です。誰もそうではありません。
c/cdp.h は単一ヘッダーのライブラリです。正確に 1 つの翻訳単位内で:
#define CDP_IMPLEMENTATION
#include "cdp.h"
宣言のために他の場所に #include "cdp.h" を追加するだけです。 c/demo.c を参照してください。
Selenium テストと同じカテゴリ (アウトプロセス自動化: コードを 1 つにまとめる)
プロセス、別の実際のブラウザ、有線プロトコルを介して通信)、マイナス
chromedriver 仲介者と外部依存関係。精神的にはより近い
ミニマルな手巻きの人形遣い。
Selenium は、WebDriver プロトコルを別の chromedriver バイナリと通信します。
それから Chrome を駆動します: 3 つのプロセス、そして chromedriver のプロセスを保持する必要があります
バージョンはChromeのものと一致しました。これは、jar にバージョン管理されたネイティブ バイナリを加えたものです。
ベンダーとアップデート。
mincdp は、仲介者なしで WebSocket 経由で CDP を Chrome と直接通信します。それぞれ
client は、その言語の標準ライブラリのみを使用する 1 つのファイルです。
正直なところ、取引: Selenium と Playwright は、大きく洗練されたクロスブラウザーを提供します
API (豊富な待機、アクション チェーン、グリッド)。 mincdp はいくつかのコマンドを提供します。
ヘッドレススモークテストの必要性、コードで読み取ることができます

端から端まで。それに応じて選択してください。
page.html 自己完結型デモ ターゲット (サーバーなし、ネットワークなし)
demo.sh Chrome を起動し、デモを実行し、破棄します (両方で共有)
Makefile make ut/demo-c/demo-java/demo/shot/clean
c/cdp.h C クライアント (単一ヘッダー ライブラリ)
c/agent.c エージェント ループ: スクリーンショット、クロード、1 つのアクションをリプレイ
Agent.sh Chrome を起動し、目標に向かってエージェントを実行し、破棄します
c/demo.c C デモンストレーション = uiut インタラクション テスト (手)
java/Cdp.java Java クライアント (単一ファイル)
java/Demo.java Java デモンストレーション = uiut インタラクション テスト (手)
testing/run.sh 回帰ランナー (make ut): codeut + uiut、PASS/FAIL/SKIP
testing/codeut.c クライアントの純粋なヘルパーのブラウザフリーの単体テスト
testing/shot.sh 目: スクリーンショット page.html を PNG に変換 (Chrome --screenshot)
血統
Legacy/ には、このプロジェクトの 2005 年の先祖が含まれています: Xvfb ベースのスクリーンショット
ブラウザーにヘッドレス モードやリモートが存在する前から、概念実証をキャプチャします。
プロトコル。ここが mincdp の始まりです。 Legacy/README.md を参照してください。
「目」がどのようにして誕生したかについての日付付きの記録 (起源の対話、
タイムライン、および帰属の要求) は、legacy/PROVENANCE.md にあります。
すべての製品名、ロゴ、ブランドはそれぞれの所有者の財産です。
Chrome は Google の商標です。 Selenium、Puppeteer、Playwright は、
それぞれのプロジェクトの名前。ここでは識別のためにのみ使用されます
および比較であり、提携や支持を暗示するものではありません。
C および Java による依存関係のない最小限の Chrome DevTools Protocol (CDP) クライアント — Selenium や chromedriver を使用せずにテスト用にヘッドレス Chrome を駆動します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください

。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Minimal, dependency-free Chrome DevTools Protocol (CDP) client in C and Java — drive headless Chrome for tests without Selenium or chromedriver. - Anode1/mincdp

GitHub - Anode1/mincdp: Minimal, dependency-free Chrome DevTools Protocol (CDP) client in C and Java — drive headless Chrome for tests without Selenium or chromedriver. · GitHub
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
Anode1
/
mincdp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits c c docs docs java java legacy legacy tests tests .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md agent.sh agent.sh demo.sh demo.sh page.html page.html View all files Repository files navigation
Drive a real headless Chrome for tests without Selenium or a version-matched
chromedriver. Two tiny, dependency-free clients that speak the Chrome DevTools
Protocol (CDP) straight to the browser over a WebSocket, in the language of your
stack:
C : c/cdp.h , a single header-only library (~410 lines, libc + POSIX
sockets only).
Java : java/Cdp.java , a single file (JDK java.net.http only, nothing
to vendor).
Same idea, same small command set, one demonstration each against the same page.
The image above was produced by mincdp itself: the demo typed into the page,
pressed Enter, then captured this driven state with cdp_screenshot
( Page.captureScreenshot ).
CDP is the wire protocol Puppeteer and Playwright use under the hood. Talking it
directly means no chromedriver middleman and no external dependency to vendor and
keep version-matched to Chrome. For a fixed set of internal pages, headless, one
browser, that is a good trade: the client is small enough to read in a sitting,
and if Chrome ever ships a breaking CDP change you fix the one file rather than
wait for an upstream driver release.
It is deliberately not a general framework. It knows only the commands a
click-and-assert test needs: navigate, evaluate JS, type text, press a key. If
you need cross-browser (Firefox/Safari) or a rich interaction API, use
Playwright; this earns its place only where minimalism and zero dependencies do.
Needs a headless Chrome/Chromium on PATH (and curl , and cc / javac ).
make ut # the whole regression suite (see Tests)
make demo-c # build c/demo and run it against page.html
make demo-java # build java/Demo and run it against page.html
make demo # both
make shot # screenshot page.html so you can SEE it
make agent GOAL="..." # let Claude drive the page toward a goal (needs an API key)
Each demo prints a series of assertions and a demo: N passed, 0 failed line. It
first inspects the page's internals (title, attribute, text, computed style,
geometry), then drives the real input path: type "mincdp works", press Enter, and
assert the page echoes it. Set CDP_DEBUG=1 to see every CDP frame.
make ut runs the full regression, in two kinds, so a test set has both hands
and eyes:
codeut -- "regular" browser-free unit tests of the client's pure helpers
(base64, JSON escaping), several examples each. No Chrome needed, so it always
runs ( tests/codeut.c ).
uiut -- the client driving real headless Chrome, in two kinds:
page internals + hands -- the demos inspect the rendered DOM (presence,
attribute, text, computed style, geometry) through Runtime.evaluate , then
interact (type, press Enter, assert the DOM changed), then screenshot the
driven state over the protocol ( cdp_screenshot -> Page.captureScreenshot )
and assert a valid PNG ( c/demo.c , java/Demo.java ).
eyes, for agents -- tests/shot.sh is the standalone shell version:
screenshots a fresh page load with Chrome's --screenshot , no client code
needed. It complements the demos' protocol screenshot (which captures a driven
state). Either PNG is there to be looked at: open it, or have an agent Read it.
Each layer reports PASS / FAIL / SKIP (a missing toolchain SKIPs, exit 77);
the suite fails only if a non-skipped layer fails:
mincdp regression:
codeut (units) PASS
uiut hands: C PASS
uiut hands: Java PASS
uiut eyes: shot PASS
The demos double as the interaction tests; make demo-c / demo-java run them
on their own.
You start Chrome; the client attaches. demo.sh does that wiring (launch a
headless Chrome with --remote-debugging-port , run the demo, tear down), and it
invokes both demos with the same contract, so one launcher and one page.html
serve both languages:
PROG... 127.0.0.1 <chrome-port> file://<abs>/page.html
Eyes and hands for an agent
The same two primitives that make a smoke test, a screenshot (eyes) and the
input calls (hands), are the entire surface an LLM needs to drive a browser.
c/agent.c is that loop and nothing else. See, think, act, repeated.
See. cdp_screenshot captures the live page, whatever state prior actions
drove it to.
Think. The PNG plus a one-line goal go to the Claude API over curl (the
one hop raw sockets can't do here is the TLS to api.anthropic.com ). The model
replies with exactly one action.
Act. The client replays that one action against the page, then loops.
The model gets a screenshot and answers with a single line, one of:
Same substring trick as the client: the reply is one short line, so pulling the
action out of the response is a single "text":"..." find, not a JSON parser.
export ANTHROPIC_API_KEY=sk-ant-...
make agent GOAL="type 'mincdp works' into the box and press Enter"
A run against page.html looks like:
step 1: JS document.getElementById('q').focus()||true
step 2: TYPE mincdp works
step 3: KEY Enter
step 4: DONE the box shows "echo: mincdp works"
Each step is a full loop: the model looked at a fresh screenshot, chose the next
keystroke, and mincdp replayed it, closing on DONE when the screenshot showed
the goal met.
It is deliberately a sketch: one action per turn, substring parsing, no retries,
thinking off for speed. That minimalism is the point. It is small enough to read
end to end before you let a model drive a browser. Like the demos, it SKIPs
(exit 77) when Chrome, curl , or ANTHROPIC_API_KEY is missing. See
c/agent.c and agent.sh .
Both clients expose the same shape (C names shown; Java is the camelCase twin):
Responses are matched by id and read with targeted substring checks, not a JSON
parser. That is the trick that keeps both clients dependency-free: every command
returns a boolean, a short ack, or one flat field (the screenshot's base64), so
looking for "value":true / exceptionDetails / "data":"..." is enough. A
command that returned deeply nested JSON would need a real parser; none does.
c/cdp.h is a single-header library. In exactly one translation unit:
#define CDP_IMPLEMENTATION
#include "cdp.h"
and just #include "cdp.h" elsewhere for the declarations. See c/demo.c .
Same category as a Selenium test (out-of-process automation: your code in one
process, a real browser in another, talking over a wire protocol), minus the
chromedriver middleman and the external dependency. Closer in spirit to a
minimal, hand-rolled Puppeteer.
Selenium talks the WebDriver protocol to a separate chromedriver binary
that then drives Chrome: three processes, and you must keep chromedriver's
version matched to Chrome's. It is a jar plus a versioned native binary to
vendor and update.
mincdp talks CDP directly to Chrome over a WebSocket, no middleman. Each
client is one file using only its language's standard library.
The trade, honestly: Selenium and Playwright give a big, polished, cross-browser
API (rich waits, action chains, a grid). mincdp gives the handful of commands a
headless smoke test needs, in code you can read end to end. Pick accordingly.
page.html self-contained demo target (no server, no network)
demo.sh launch Chrome, run a demo, tear down (shared by both)
Makefile make ut / demo-c / demo-java / demo / shot / clean
c/cdp.h the C client (single-header library)
c/agent.c the agent loop: screenshot, Claude, replay one action
agent.sh launch Chrome, run the agent toward a GOAL, tear down
c/demo.c the C demonstration = the uiut interaction test (hands)
java/Cdp.java the Java client (single file)
java/Demo.java the Java demonstration = the uiut interaction test (hands)
tests/run.sh the regression runner (make ut): codeut + uiut, PASS/FAIL/SKIP
tests/codeut.c browser-free unit tests of the client's pure helpers
tests/shot.sh the eyes: screenshot page.html to a PNG (Chrome --screenshot)
Lineage
legacy/ holds the 2005 ancestor of this project: an Xvfb-based screenshot
capture proof of concept, from before browsers had a headless mode or a remote
protocol. It is where mincdp started. See legacy/README.md .
The dated record of how the "eyes" came about (the origin dialog, the
timeline, and the attribution ask) is in legacy/PROVENANCE.md .
All product names, logos, and brands are property of their respective owners.
Chrome is a trademark of Google; Selenium, Puppeteer, and Playwright are the
names of their respective projects. They are used here only for identification
and comparison, and no affiliation or endorsement is implied.
Minimal, dependency-free Chrome DevTools Protocol (CDP) client in C and Java — drive headless Chrome for tests without Selenium or chromedriver.
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
