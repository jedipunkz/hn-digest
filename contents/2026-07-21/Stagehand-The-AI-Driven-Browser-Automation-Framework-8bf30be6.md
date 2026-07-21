---
source: "https://www.akashtandon.in/interactive-explainers/stagehand/"
hn_url: "https://news.ycombinator.com/item?id=48991148"
title: "Stagehand: The AI-Driven Browser Automation Framework"
article_title: "Stagehand: Inside the AI-Driven Browser Automation Framework"
author: "akashtndn"
captured_at: "2026-07-21T12:19:58Z"
capture_tool: "hn-digest"
hn_id: 48991148
score: 1
comments: 0
posted_at: "2026-07-21T12:03:52Z"
tags:
  - hacker-news
  - translated
---

# Stagehand: The AI-Driven Browser Automation Framework

- HN: [48991148](https://news.ycombinator.com/item?id=48991148)
- Source: [www.akashtandon.in](https://www.akashtandon.in/interactive-explainers/stagehand/)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T12:03:52Z

## Translation

タイトル: Stagehand: AI 主導のブラウザ自動化フレームワーク
記事のタイトル: Stagehand: AI 主導のブラウザ自動化フレームワークの内部

記事本文:
インタラクティブな説明者
GitHub
セレクターが壊れた夜
再設計を出荷する
ブラウザ用の 4 つの動詞
動詞、ライブ
指示からクリックまで
act() 内
一度書いたら永久に実行可能
キャッシュ フライホイール
劇作家を卒業
仲介業者を排除する
正直なトレードオフ
インタラクティブな説明者
舞台係:
Inside the AI-Driven Browser Automation Framework
Hardcoded selectors break the moment a page changes;完全に自律的なエージェントは、
書くのは早いが、信頼するのは難しい。 Stagehand, the open-source browser-automation SDK
from Browserbase, takes a third path - you write plain-English instructions, an AI
resolves them to real elements at runtime, and deterministic code does the clicking.
Explore how its four verbs work, why it caches the AI's answers, and what it gave up
劇作家を後にするために。
スクレイパーは 8 か月間、問題なく動作しました。そして、ある火曜日の夜、フロントエンドが
developer at the site you automate renamed a CSS class - .btn--buy became
.purchase-cta - and your pipeline died without a sound.何もクラッシュしませんでした
大声で。 The script looked for a button that no longer existed, timed out, retried,
またタイムアウトになった。 You found out Thursday, from a stakeholder asking where the data went.
This is the standing bargain of traditional browser automation. Playwright などのツール
Selenium, and Puppeteer are exact: you tell them precisely which element to click, and
彼らは無料で、素早く、確実にクリックします。でもその正確さが失敗でもある
モード。 A hardcoded selector is a bet that the page will never change, and the page always
変化します。 Every automation you write is quietly accruing a maintenance debt that comes
少なくともご都合のよい時間までに提出してください。
At the other extreme sits the fully autonomous AI agent.あなたはそれに目標を与えます - 「
ヘッドフォン」 - 見た目は

ページの閲覧、その理由、クリック数の把握
それ自体の。再設計してもひるむことはありません。しかし、今度は別の問題が発生します。エージェントが
実行するたびに異なるパスを選択し、ステップごとにさまようとデバッグが困難になる
モデルトークンを焼きます。脆いものをブラックボックスと交換しました。
Stagehand はオープンソースです
Browserbase のフレームワーク (
ヘッドレス ブラウザ インフラストラクチャを販売する会社）は、賭けに基づいて構築されています。
自動化は選択でも目標でもなく、指示です。 MITライセンス、
23,000 を超える GitHub スターを持ち、「ブラウザ エージェント用の SDK」を自称しています。あなたは書きます
act("購入ボタンをクリック") とすると、AI が実行時にその文を次のように解決します。
今日の購入ボタンが何であれ、通常の決定論的なブラウザ コード
クリックを実行します。
開発者は Stagehand を使用して Web を確実に自動化します。
違いを感じるために、何かを壊してみてください。下の図は、同じ小さなものを示しています
ストアフロントは 2 つの方法で自動化されています: 左側のハードコーディングされたセレクター、ステージハンド スタイル
右側の指示。両方を実行し、再設計を出荷して再度実行します - 見てください
どちらが生き残るか。
舞台係とは、劇場で舞台袖で小道具を動かし、ショーを演出する人です。
続けることができます。このフレームワークは、意図的に小さな API でその名前を獲得しました。つまり、4 つの動詞です。
ブラウザに求められるすべてをカバーします。
最初の 3 つは精密ツールです。 act() は 1 ステップだけ実行します。
ドキュメントでは、命令をアトミックに保つことを強調しています。「ログイン ボタンをクリックする」ではなく、
「ログインして設定に移動します」。 extract() はスクレイピングを型付き関数に変換します
call: Zod スキーマを渡し、実際にそれに準拠するオブジェクトを返します (「
Zod が初めての場合は、下のボックスにチェックを入れてください)。 observ() は偵察です - それは教えます
取り消せないものがある前に、ページ上で実行可能なものを確認してください

兄弟が起こります。
Zod は、データの形状を記述するための TypeScript ライブラリです。各フィールドを詳しく説明します
そしてその型を 1 回。演出家はその説明を記入すべきものとしてモデルに示し、
次に、オブジェクトを返す前に、モデルの答えをモデルの答えと照合します。というページ
間違った形状を返すと、黙って間違った結果が返されるのではなく、エラーが発生します。
4 番目の動詞は誰が運転するかを変更します。 Agent() はゴールを取得してループします -
観察し、決定し、行動する - 目標が達成されるまで。通常のLLMで裏付けることができます
またはコンピューター専用モデル:
早めに知っておく価値のある詳細がもう 1 つあります。指示はモデル プロバイダーに送信されるため、
Stagehand は、機密性の高いものに対してプレースホルダー変数を提供します。あなたは書きます
act("パスワードフィールドに %password% を入力してください", { 変数: { パスワード: ... } })
そしてシークレットはローカルに置き換えられます。LLM はプレースホルダーのみを参照します。
2 つのモード間の緊張感が興味深い部分です。精密動詞は次のとおりです。
予測可能でデバッグ可能ですが、ルートを記述する必要があります。エージェントがルートを見つける
ただし、予測できるのはそのモデルと同程度です。 Stagehand 自身のドキュメントには、ほとんどのチームがどのように行うかを説明しています
二乗: エージェントを探索に使用し、クリティカル パスを特定します。
個々のプリミティブ。以下の図は、文字通り、4 つの動詞すべてを同じページにまとめています。タブを切り替える
そして、それぞれが同じ小さな店に対して何をするかを観察してください。
では、act("購入ボタンをクリック") の内部では実際に何が起こっているのでしょうか?感じることができる
魔法のようですが、各ステージは理解できるものであることがわかります
エンジニアリング - エンドツーエンドでたどることのできる短いパイプライン。
まず、Stagehand はモデルにページを表示する必要がありますが、スクリーンショットは送信しません。
またはデフォルトで生の HTML ダンプ。ページのアクセシビリティ ツリーを読み取ります。
ブラウザがスクリーン リーダー用にすでに保持しているセマンティック アウトライン。

ある
<div class="x7-btn"> は単なる「今すぐ購入」ボタンになります。
(これは直接確認できます。引数なしで extract() を呼び出すと戻り値が返されます。
アクセシビリティ ツリー自体です。) これは非常に経済的な選択です。ツリーは
DOM よりも桁違いに小さく、人間がページを記述するのと同じ方法でページを記述します。
マークアップではなく、役割とラベルによって。
モデルは命令とそのツリーを受け取り、ターゲットとメソッドを選択します。
クリック、入力、入力、押し、スクロール、または選択します。重要なのは、その決定が次のように具体化されることです。
具体的で退屈な成果物: act() が返す ActResult
使用された XPath セレクターを、成功ステータスと説明とともに記録します。
何が起こったのか。 AI の判断は記録できるものにまとめられます。
検査し、次のセクションで示すようにキャッシュします。その後、決定論的なコードがディスパッチされます
実際のイベント、ターゲットが存在するときに iframe と Shadow DOM を独自に走査します。
一つの中に。
これが、指示の表現が思っている以上に重要である理由でもあります。ドキュメントプッシュ
具体性を重視 - 「ユーザー John Smith の横にある赤い [削除] ボタンが「クリック」よりも優れています
ボタン」 - モデルが候補の中から選択しているため、曖昧さがあり、間違っています。
要素がアウト。一か八かの場合には、observ() というプレビュー イディオムがあります。
解決されたアクションを実行せずに返し、それを に渡します。
表示されたものが気に入った場合にのみ、act() を実行してください。
以下のパイプラインを自分でステップ実行してください。ボタンを押すたびに、次のステップに進みます。
実際のクリックに向けてさらに 1 段階進んだ指示。
上記すべてに対して明らかな反対意見があります。アクションごとの LLM 呼び出しは遅く、
お金がかかります。自動化が 15 分ごとに実行される場合は、料金を支払いたくないでしょう。
同じ購入ボタンを 1 日 96 回再発見するモデル。
Stagehand の答えは、積極的なキャッシュであり、フレームワークの「書き込みオン」が行われる場所です。

え、
「永遠に走れ」というピッチが由来。初めて命令が解決されたとき、結果は -
前のセクションの具体的な XPath が保存されます。後続の実行はすべてリプレイされます
格納されたアクションを直接呼び出し、推論を完全にスキップします。ローカルでは、ポイントしてオプトインします
キャッシュディレクトリのコンストラクタ。 Browserbase のクラウドでは、サーバー側のキャッシュがオンになっています
デフォルトでは、すべての結果によって、それが「ヒット」だったのか、
「ミス」：
賢いのは、キャッシュが古くなったときに何が起こるかということです。からの再設計
オープニング シーンが到着すると、保存されている XPath の一致が停止し、古典的なスクリプトは次のようになります。
死んだ。 Stagehand は代わりに、ページが変更されたことを検出し、そのページに対して推論を再実行します。
新しいページを作成し、独自のキャッシュを修復します - Browserbase が呼び出す動作
自己修復。木曜日の朝の出来事だった再設計
1 つのキャッシュミスと 1 つの追加のモデル呼び出しになります。
以下のシミュレーションを数回実行して、経済的な展開を観察してください。最初の走行
遅くてトークンがかかります。次の実行は高速かつ無料です。その後、再設計を出荷し、
システムが一度だけ攻撃を受けるのを見てください。
Stagehand は、独自の自動化エンジンとして開始されたわけではありません。 Browserbase がオープンソース化したとき
2024 年には、Playwright の上のレイヤーになっていました。まさにそれが、人気を博した理由です。あなた
act() と extract() を既存の Playwright にドロップできます
スクリプトは数分で作成でき、構築した他のすべては保持されます。バージョン 2 までのパッケージは
毎週 500,000 件以上のダウンロードが見られます。
しかし、鱗はDNAの不一致を明らかにしました。 Playwright はテスト ツールです。その前に
何かをクリックすると、アクション可能性チェックが実行されます - 要素が追加されるのを待ちます
Playwright の用語では、「可視、安定、有効、受信可能なイベント」です。 CI スイートの場合、
その注意こそが重要なのだ。高スループットの自動化を行うには、ボタンをクリックします。
解決されてキャッシュされた

数週間前は、あらゆるアクションに遅延税が課せられていました。ブラウザベース
チームはまた、低レベルの摩擦に遭遇し続けました: Playwright は安定したフレームを公開していません
ID が原因で、Chrome DevTools プロトコル コマンドをルーティングすることが困難になりました。
特定の iframe;その CDP セッションは、フレームではなくページとコンテキストにスコープされます。
そしてユーザーは、長時間のセッションでメモリがクリープすることを報告しました。
そのため、2025 年 10 月に Stagehand v3 は仲介者を排除しました。エンジンは次のように書き換えられました
Chrome DevTools プロトコルを直接話します - 同じ生の WebSocket
ブラウザー独自の DevTools が使用するプロトコル - ストリーミング アクセシビリティ ツリー、DOM スナップショット、
間に翻訳層を介さずに、ブラウザーから直接ネットワーク イベントを送信できます。
深くネストされた iframe と Shadow DOM の Browserbase のベンチマーク、v3 で測定
Playwright ベースのアーキテクチャよりも平均 44.11% 高速です。そして素敵な展開で、
依存関係を取り除くことで、Stagehand の互換性は低下するどころかさらに向上しました: v3 API
ドライバーに依存せず、生きているのではなく劇作家や人形遣いと協力して働いています
そのうちの 1 つは、TypeScript を超えた SDK と REST API を備えています。
次の図は、メッセージ トラフィックとしてのアーキテクチャの違いを示しています。両車線
同じ繰り返しのアクションを実行します。ホップを数えます。
これらはどれも無料ではなく、Stagehand ではどこにコストがかかるのかがすっきりと読みやすくなっています。
移動しました。すべてのキャッシュ ミスは従量制の LLM 呼び出しであるため、一連の自動化が必要になります。
頻繁にページが変更されると、実際の推論費用が増加します。このプロジェクトは推奨しません
ローカル モデルなので、実際にはクラウド モデル API に接続することになります。そしてエンジンは、
Chromium のみ - Chrome、Edge、Arc、Brave - なので、クロスブラウザー テストは引き続き行われます。
伝統的な道具。
最も明確な比較はブラウザです
図 1 とは対極にある Python フレームワークである を使用します。
ブラウザを使用してください

スクリーンショットから機能するエージェントにタスクを送信 - AI が認識
人は何を見て、次の行動を決めるのか。より大きなコミュニティがあります（おおよそ
GitHub のスターは 80,000 個、Stagehand の 23,000 個に相当します）。また、ローカル推論をサポートしています。
Ollama を使用すると、実質的に無料で実行できるようになります。価格は変動です: 自律型
エージェントは実行ごとに異なるルートを選択する可能性があり、複数ステップの自律性によってより多くのトークンが消費されます
アトミックなキャッシュ可能な操作よりもタスクごとに実行されます。
TypeScript ファースト; v3 以降、より多くの言語での SDK
段階的な制御、キャッシュ可能および再生可能
アクセシビリティ ツリー コンテキスト。ダイレクト CDP エンジン
ネストされた iframe と Shadow DOM に強い
クラウド LLM API。ローカルモデルは推奨されません
Python ファースト、データサイエンススタックに適合
完全自律型のスクリーンショットベースのエージェント
Ollama によるローカル推論 - API コストゼロ
タスクごとのトークンが増加し、実行ごとの差異が増加
Scrapfly の独立したレビューは完全に意見が分かれています: ブラウザの使用については次のとおりです。
自律性、Python、または無料のローカル推論が必要です。必要なときにステージハンドに手を伸ばせる
予測可能でデバッグ可能な実稼働グレードの実行 - 特に iframe とシャドウのどこでも
DOMが潜んでいます。これは Stagehand 自身のアドバイスと一致します。
「探索のためのエージェント、クリティカル パスのための個々のプリミティブ。」
うーん

[切り捨てられた]

## Original Extract

Interactive Explainers
GitHub
The Night the Selector Broke
Ship a Redesign
Four Verbs for a Browser
The Verbs, Live
From Instruction to Click
Inside act()
Write Once, Run Forever
The Cache Flywheel
Graduating from Playwright
Cutting Out the Middleman
The Honest Trade-offs
Interactive Explainer
Stagehand:
Inside the AI-Driven Browser Automation Framework
Hardcoded selectors break the moment a page changes; fully autonomous agents are
quick to write but hard to trust. Stagehand, the open-source browser-automation SDK
from Browserbase, takes a third path - you write plain-English instructions, an AI
resolves them to real elements at runtime, and deterministic code does the clicking.
Explore how its four verbs work, why it caches the AI's answers, and what it gave up
to leave Playwright behind.
Your scraper ran flawlessly for eight months. Then, one Tuesday night, a frontend
developer at the site you automate renamed a CSS class - .btn--buy became
.purchase-cta - and your pipeline died without a sound. Nothing crashed
loudly. The script looked for a button that no longer existed, timed out, retried,
timed out again. You found out Thursday, from a stakeholder asking where the data went.
This is the standing bargain of traditional browser automation. Tools like Playwright,
Selenium, and Puppeteer are exact: you tell them precisely which element to click, and
they click it fast, deterministically, for free. But that exactness is also the failure
mode. A hardcoded selector is a bet that the page will never change, and the page always
changes. Every automation you write is quietly accruing a maintenance debt that comes
due at the least convenient time.
At the other extreme sits the fully autonomous AI agent. You hand it a goal - "buy the
headphones" - and it looks at the page, reasons about it, and figures out the clicks on
its own. Redesigns don't faze it. But now you have a different problem: the agent might
take a different path every run, it's hard to debug when it wanders, and every step
burns model tokens. You've traded brittle for black box.
Stagehand is an open-source
framework from Browserbase (the
company that sells headless browser infrastructure) built on a wager: the right unit of
automation is neither the selector nor the goal, but the instruction . MIT-licensed,
with over 23,000 GitHub stars, it bills itself as "the SDK for browser agents." You write
act("click the buy button") , and an AI resolves that sentence at runtime to
whatever the buy button happens to be today - then ordinary, deterministic browser code
performs the click.
Developers use Stagehand to reliably automate the web.
To feel the difference, try breaking something. The figure below shows the same tiny
storefront automated two ways: a hardcoded selector on the left, a Stagehand-style
instruction on the right. Run both, then ship a redesign and run them again - watch
which one survives.
A stagehand, in the theater, is the person in the wings who moves the props so the show
can go on. The framework earns its name with a deliberately small API: four verbs that
cover everything you'd ask of a browser.
The first three are the precision tools. act() executes exactly one step;
the docs are emphatic about keeping instructions atomic - "click the login button", not
"log in and go to settings". extract() turns scraping into a typed function
call: you hand it a Zod schema and get back an object that actually conforms to it (see
the box below if Zod is new to you). observe() is reconnaissance - it tells
you what's actionable on the page before anything irreversible happens.
Zod is a TypeScript library for describing the shape of data. You spell out each field
and its type once; Stagehand shows that description to the model as the thing to fill in,
then checks the model's answer against it before handing the object back. A page that
returns the wrong shape raises an error instead of a silently-wrong result.
The fourth verb changes who's driving. agent() takes a goal and loops -
observing, deciding, acting - until the goal is met. You can back it with a regular LLM
or a dedicated computer-use model:
One more detail worth knowing early: instructions are sent to a model provider, so
Stagehand gives you placeholder variables for anything sensitive. You write
act("type %password% into the password field", { variables: { password: ... } })
and the secret is substituted locally - the LLM only ever sees the placeholder.
The tension between the two modes is the interesting part. The precision verbs are
predictable and debuggable but need you to write the route; the agent finds the route
but is only as predictable as its model. Stagehand's own docs describe how most teams
square it: use the agent for exploration, and pin the critical paths down with
individual primitives. The figure below puts all four verbs on the same page - literally. Switch tabs
and watch what each one does to the same little store.
So what actually happens inside act("click the buy button") ? It can feel
like magic, but each stage turns out to be a piece of understandable
engineering - a short pipeline you can follow from end to end.
First, Stagehand needs to show the model the page, and it does not send a screenshot
or a raw HTML dump by default. It reads the page's accessibility tree - the
semantic outline the browser already maintains for screen readers, where a
<div class="x7-btn"> becomes simply button, "Buy now" .
(You can see this directly: calling extract() with no arguments returns
the accessibility tree itself.) It's a beautifully economical choice - the tree is
orders of magnitude smaller than the DOM, and it describes the page the way a person
would: by role and label, not by markup.
The model receives your instruction plus that tree and picks a target and a method -
click, fill, type, press, scroll, or select. Crucially, the decision materializes as a
concrete, boring artifact: the ActResult that act() returns
records the XPath selectors that were used, alongside success status and a description
of what happened. The AI's judgment gets compiled down to something you can log,
inspect, and - as the next section shows - cache. Then deterministic code dispatches
the actual event, traversing iframes and shadow DOM on its own when the target lives
inside one.
This is also why instruction phrasing matters more than it might seem. The docs push
you toward specificity - "the red Delete button next to user John Smith" beats "click
the button" - because the model is choosing among candidates, and ambiguity in, wrong
element out. For anything high-stakes, there's a preview idiom: observe()
returns the resolved action without executing it , and you pass it to
act() only if you like what you see.
Step through the pipeline yourself below. Each press of the button advances the
instruction one stage further toward a real click.
There's an obvious objection to everything above: an LLM call per action is slow and
costs money. If your automation runs every fifteen minutes, you do not want to pay a
model to re-discover the same buy button 96 times a day.
Stagehand's answer is aggressive caching, and it's where the framework's "write once,
run forever" pitch comes from. The first time an instruction resolves, the result -
that concrete XPath from the previous section - is stored. Every subsequent run replays
the stored action directly, skipping inference entirely. Locally, you opt in by pointing
the constructor at a cache directory; on Browserbase's cloud, a server-side cache is on
by default, and every result tells you whether it was a "HIT" or a
"MISS" :
The clever part is what happens when the cache goes stale. That redesign from the
opening scene lands, the stored XPath stops matching, and a classic script would be
dead. Stagehand instead detects that the page has changed, re-runs inference against
the new page, and repairs its own cache - the behavior Browserbase calls
self-healing . The redesign that used to be a Thursday-morning incident
becomes one cache miss and one extra model call.
Run the simulation below a few times and watch the economics play out. The first run
is slow and costs tokens. The next runs are fast and free. Then ship a redesign and
see the system take the hit once - and only once.
Stagehand didn't start as its own automation engine. When Browserbase open-sourced it
in 2024, it was a layer on top of Playwright - which was exactly why it caught on. You
could drop act() and extract() into an existing Playwright
script in minutes, keeping everything else you'd built. By version 2 the package was
seeing over 500,000 weekly downloads.
But scale surfaced a mismatch in DNA. Playwright is a testing tool. Before it
clicks anything, it runs actionability checks - waiting for the element to be, in
Playwright's terms, "visible, stable, enabled, and receiving events." For a CI suite,
that caution is the whole point. For a high-throughput automation clicking a button it
resolved and cached weeks ago, it's latency tax on every single action. The Browserbase
team also kept hitting lower-level friction: Playwright doesn't expose stable frame
identities, which made it awkward to route Chrome DevTools Protocol commands at
specific iframes; its CDP sessions are scoped to pages and contexts rather than frames;
and users reported memory creep in long-lived sessions.
So in October 2025, Stagehand v3 removed the middleman. The engine was rewritten to
speak Chrome DevTools Protocol directly - the same raw websocket
protocol the browser's own DevTools use - streaming accessibility trees, DOM snapshots,
and network events straight from the browser with no translation layer in between.
On Browserbase's benchmark of deeply nested iframes and shadow DOMs, v3 measured
44.11% faster on average than the Playwright-based architecture. And in a nice twist,
shedding the dependency made Stagehand more compatible, not less: the v3 API
is driver-agnostic, working alongside Playwright and Puppeteer rather than living
inside one of them, with SDKs beyond TypeScript and a REST API.
The figure below shows the architectural difference as message traffic. Both lanes
perform the same repeated action; count the hops.
None of this is free, and Stagehand is refreshingly legible about where the costs
moved. Every cache miss is a metered LLM call, so a fleet of automations against
frequently-changing pages runs up a real inference bill. The project doesn't recommend
local models, so in practice you're wiring in a cloud model API. And the engine is
Chromium-only - Chrome, Edge, Arc, Brave - so cross-browser testing stays with the
traditional tools.
The sharpest comparison is with Browser
Use , the Python framework that owns the other end of the spectrum from figure 1.
Browser Use hands the whole task to an agent that works from screenshots - the AI sees
what a person sees and decides each next move. It has the larger community (roughly
80,000 GitHub stars to Stagehand's 23,000) and it supports local inference through
Ollama, which makes it effectively free to run. The price is variance: an autonomous
agent may take a different route each run, and multi-step autonomy burns more tokens
per task than atomic, cacheable operations.
TypeScript-first; SDKs in more languages since v3
Step-by-step control, cacheable and replayable
Accessibility-tree context; direct CDP engine
Strong on nested iframes and shadow DOM
Cloud LLM APIs; local models not recommended
Python-first, fits data-science stacks
Fully autonomous, screenshot-based agent
Local inference via Ollama - zero API cost
More tokens per task, more run-to-run variance
Scrapfly's independent review lands on a clean split: reach for Browser Use when you
want autonomy, Python, or free local inference; reach for Stagehand when you need
predictable, debuggable, production-grade runs - especially anywhere iframes and shadow
DOM lurk. That matches Stagehand's own advice :
"agent for exploration, individual primitives for critical paths."
Wh

[truncated]
