---
source: "https://blog.leafphp.dev/posts/leaf-5"
hn_url: "https://news.ycombinator.com/item?id=49330653"
title: "Leaf 5 released: a PHP framework built for humans and AI agents"
article_title: "Leaf 5: Build products at the speed of thought | Leaf Blog"
image: ""
author: "mychi"
captured_at: "2026-08-17T14:19:20Z"
capture_tool: "hn-digest"
hn_id: 49330653
score: 1
comments: 0
posted_at: "2026-08-17T13:39:06Z"
tags:
  - hacker-news
  - translated
---

# Leaf 5 released: a PHP framework built for humans and AI agents

- HN: [49330653](https://news.ycombinator.com/item?id=49330653)
- Source: [blog.leafphp.dev](https://blog.leafphp.dev/posts/leaf-5)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T13:39:06Z

## Translation

タイトル: Leaf 5 リリース: 人間と AI エージェント向けに構築された PHP フレームワーク
記事のタイトル: リーフ 5: 思考のスピードで製品を構築する |リーフブログ
説明: Leaf PHP の公式ブログ

記事本文:
Leaf Blog Discord スポンサードキュメント → 2026 年 8 月 13 日
リーフ 5: 思考のスピードで製品を構築する
名前 Michael Darko 寄稿者 + 6 人の寄稿者
Leaf 5 は Leaf 史上最大のリリースであり、単なるバージョンアップではありません。これは、AI と並行して構築する世界における PHP フレームワークのあるべき姿の自然な進化です。プロジェクト コンテキストが組み込まれた製品ファーストのフレームワーク、あらゆる PHP アプリで動作するツール ファミリ、コードがクリーンで読みやすい状態に保たれるという約束です。
どのフレームワークも高速化を謳っています。しかし、「より速い」は変わりました。今では、コードを書くだけでなく、コードを監督することもできます。あなたが機能を説明し、AI アシスタントがその機能の下書きを作成し、レビューして改良します。ボトルネックは入力ではなくなりました。ツールがアプリをどれだけ理解できるかが重要です。
ほとんどのフレームワークはそのように設計されていません。慣例は人々の頭の中に存在し、構造は抽象化の背後に隠れており、コードベースに組み込まれたアシスタントはその方​​法を推測する必要があります。
リーフ 5 が私たちの答えです。構造が明白で、パターンに一貫性があり、プロジェクト自体が人間とエージェントの両方が読み書きできる共有コンテキストを保持するフレームワークです。課金モジュールを要求すると、生成されたコードがアプリの残りの部分と同じように適切な場所に配置されます。ブラックボックスはありません。その後の掃除もなし。
また、当社は製品第一主義であるため、Leaf 5 には、実際の製品に必要な要素 (認証、データベース、キュー、メール、請求、その他 25 以上の機能) が、アプリが要求したときに正確に取り込める本番対応モジュールとして同梱されています。それ以上は何もありません。
この部分は今も変わっていませんし、これからも変わりません。隠されたランタイム、独自のロックイン、展開の魔法はありません。 Leaf 5 アプリは PHP アプリです。 PHP が実行される場所であればどこでも実行でき、生成されるコードはすべて 1 行ずつ読み取ることができます。
ある

pp ( ) -> get ( '/' , function ( ) {
応答 ( ) -> json ( [
「メッセージ」 => 「リーフ 5 😉」
]);
} ) ;
app() -> run() ;
以前に Leaf を作成したことがある場合は、Leaf 5 についてすでに知っています。まだ知らない場合でも、午後には理解できるでしょう。
Leaf 5 におけるもう 1 つの大きな変化は、Leaf を 1 つのフレームワークとして考えるのをやめ、焦点を絞ったツールのファミリーとして構築し始めたことです。それぞれが独自のホームを持ち、それぞれが Leaf 自体を超えて使用できるようになりました。
Alchemy : QA セットアップ全体を 1 つのファイルにまとめます。テスト (Pest または PHPUnit)、コード スタイル、自動リファクタリング、静的分析、GitHub、GitLab、CircleCI の CI パイプラインはすべて 1 つの alchemy.yml で記述されます。 Leaf、Laravel、Symfony、Slim、またはプレーンな PHP で動作し、使用するまで何もインストールされません。また、alchemy eject を使用すると、いつでも実際の設定ファイルを使用して終了できます。
Seedling : Leaf MVC エクスペリエンスを使用して CLI ツールを構築します。内部には、手間のかからないコンソール フレームワークである Sprout があり、完全な安定性パスを通過したばかりです。つまり、実際のテスト スイート、パイプされた入力と Windows の修正、macOS、Ubuntu、Windows で実行される CI です。
フェッチ: PHP の場合は fetch()。 JavaScript でよく知られているリクエスト API は、あらゆる PHP アプリで動作します。 v5 用に内部を再構築しました。文書化されたすべてのオプションが実際に機能するようになり、フォーム エンコーディングがよりスマートになり、応答がより予測可能になり、コミットごとに実サーバーに対して実行されるテストによって全面がカバーされます。
Hana JS : シンプルで軽量な React の代替品。フロントエンドも同様に「過剰なエンジニアリングを行わない」扱いに値するためです。
これは 25 以上のモジュールの上にあります。必要なものだけを正確に入手し、他には何も入れません。
なぜフレームワークに依存しないのでしょうか?一つには、優れたツールはより多くの利用者に値するからです。しかし正直に言うと、これは AI 時代の賭けでもあります。アシスタントがタスク用のライブラリを選択するときは、小型で機能性が高く、十分に文書化されたツールが勝つのです。リーフが欲しいんだ」

のツールは、アプリが Leaf アプリであるかどうかにかかわらず、明らかな選択となります。
Leaf は他のフレームワークと競合しているだけではありません。それは、初期の製品が動作し始め、ユーザーが現れ、作業を遅らせることなくコードベースをより本格的にする必要がある瞬間と競合しています。
古典的なオプションはどちらもその時点で負担がかかります。 tiny-router パスを使用すると、迅速に開始できますが、その後、認証、データベース パターン、検証、メール、キュー、ビュー、および構造を自分で選択して接続するための製品時間がかかります。重いフレームワークへの移行は初日から深刻ですが、最初のバージョンでは、製品がまだ必要としているよりも多くの概念と儀式が継承されています。 Leaf の道は、小規模に開始し、アプリが要求したときにファーストパーティ製品モジュールを追加し、構造が適切な位置を獲得したら MVC に移行し、成長に合わせて全体を人間やエージェントが読みやすい状態に保つことです。
最大の PHP 製品エコシステムが必要で、その完全なアプリケーション モデルを喜んで採用する場合は、Laravel を選択してください。
深いアーキテクチャ制御、成熟したコンポーネント、エンタープライズグレードの構成が必要な場合は、Symfony を選択してください。
可能な限り最小の HTTP レイヤーが必要で、残りの製品スタックを自分で組み立てることに抵抗がない場合は、Slim を選択してください。
本物の PHP 製品をすぐにユーザーの手に渡し、アプリを理解できるほど小さく保ち、製品の成長に合わせて認証、データ、フロントエンド、ジョブ、請求、展開、AI 支援メンテナンスへのパスを残しておきたい場合は、Leaf を選択してください。
製品ストーリーが重要ですが、パフォーマンスも約束の一部です。リーフ 5 はフレームワークのオーバーヘッドを低く抑え、製品がより多くのリクエスト バジェットを獲得できるようにします。
これは、完全なアプリケーション スループットではなく、ルート ルックアップとディスパッチ オーバーヘッドを比較しています。静的/動的混合ルート テーブルに対する 10,000 件のルックアップ、正確なヒットとパラメータ化されたヒットを交互に実行し、平均

PHP 8.2 で 5 回の実行にわたって実行されました。すべてのフレームワークは同じ URI を解決し、Laravel 列のリクエスト オブジェクトはタイミングループの外側で事前に構築され、各列はルーティング作業のみを測定します: Illuminate Routing for Laravel、symfony/ルーティング UrlMatcher、および Slim 4 の FastRoute を利用したディスパッチ。
ルックアップがフラットなままである理由は、リーフ 5 が正確なルートにインデックスを付け、動的ルートをメソッドと最初のセグメントごとにバケット化するため、完全一致がパターンに触れることはなく、動的リクエストは独自のバケットのみをスキャンするためです。同じパスにより、繰り返しの env 読み取りが約 41 倍、リクエストあたりの URI 解析が約 4.4 倍高速になりました。
何かが壊れたとき、最初に読むのはエラー メッセージです。 AI アシスタントを使用して構築している場合、通常、アシスタントが読み取るのは AI アシスタントだけです。そのため、リーフ 5 では、エラー メッセージは、メンテナがサポート チケットに答えるような方法で書かれています。つまり、何が問題で、それに対して何をすべきか、アプリの設定方法に応じて修正がどのように変化するかなどです。
これは、Leaf 5 の新しいエラー エンジンによる本物のクラッシュ画面です。修復はメッセージです。 Leaf MVC アプリは Leaf key:generate を実行でき、Lite アプリは実行できないことを認識しており、単一のドキュメントを開く前にその旨を表示します。クラッシュ画面もハンドオフ用に構築されています。Markdown としてコピーすると、レポート全体が任意のチャットに貼り付けることができる内容に変換され、Open with AI によってチャットにレポートが直接送信されます。
AI トークンの燃焼量を減らす #
AI を使用して構築する場合、トークンはお金であり、ほとんどのフレームワークは賢明に使用するには高価です。アシスタントは何千語ものドキュメントを読んで 1 つの API を抽出し、規則を推測し、何も説明しないエラーが発生した場合は再試行します。
リーフ 5 はアシスタントのコンテキスト ウィンドウを注意深く扱います。リファレンス ドキュメントは 1 ホップ離れた安定した URL に存在し、ファイルごとに 1 つのトピックがあり、各ファイルは全体を読み込むのに十分な大きさです。 L

eaf スキルは、エージェントが実際に犯す正確な間違いを導きます。間違った仮定を前もって修正する方が、後でデバッグするよりも何百倍もコストがかからないためです。独自の修正を指定するエラーは、try-fail-search-retry ループを 1 つのステップにまとめます。 .leaf/CONTEXT.md は、アシスタントがアーキテクチャを再検出する必要がないことを意味します。セッションごとにコードベースを再探索する代わりに、1 つの小さなファイルを読み取ります。
私たちはこれについて何も推測しません。 AI エージェントは Leaf 上で実際のアプリを構築します。私たちは、AI エージェントが開かなければならなかったすべてのファイルと焼き付けられたすべての再試行をカウントし、トークンのコストが何であれ、修正となります。上記のエラー画面が表示されるのは、エージェントがそのメッセージの不明瞭なバージョンを見つけ、代わりに聞く必要がある内容を正確に伝えたためです。この話については、別の投稿に値するので、すぐに詳しく説明します。
最近、leafphp.dev にアクセスしたことがある方は、ドキュメントだけではなく、新しい外観を見たことがあるでしょう。 Alchemy、Seedling、Fetch、hana、そしてこのブログはすべて、ブリコラージュ グロテスクな見出し、ヘアライン グリッド、そして各製品が独自の個性と色を維持するという共通のデザイン言語を中心に再デザインされました。違う家、同じ家族。
ドキュメント自体も、最新の読書エクスペリエンス、最終的にすべての製品を接続するエコシステム メニュー、Leaf 5 のすべての変更点に合わせて更新されたコンテンツなど、深いパスを獲得しました。
私たちが考えていること (正直な部分) #
このリリースを構築する際、いくつかの原則が私たちを正直に保ちました。これらは書き留める価値があります。
AI はあなたの判断力を強化するものであり、あなたの判断を置き換えるものではありません。 Leaf 5 はアシスタントに素晴らしい仕事をするためのコンテキストを提供しますが、出力は常にプレーンでレビュー可能な PHP です。あなたは建築家であり続けます。
効果のあるものだけを宣伝します。新しいサイトのすべての主張は、コードとテストによって裏付けられています。ドキュメントと現実の間にギャップが見つかった場合 (いくつか見つかりました)、コピーではなくコードを修正しました。
小さな表面、

鋭いエッジ。フレームワークは軽いままです。機能は、永久に成長するカーネルからではなく、選択したモジュールから得られます。
無料は無料のままです。 Leaf 5、すべてのモジュール、および上記のすべてのツールはオープンソースであり、コミュニティの支援を受けています。 Leaf によって時間を節約できるのであれば、スポンサーになることでそれを維持できます。
まだ終わっていません。ここからのロードマップ: リアルタイム API、より優れた導入ストーリー、より深いサードパーティ統合、共有コンテキスト ワークフローを中心としたツールの追加。各作品が到着次第、ここで共有します。
最新の PHP エコシステムの他の部分と同じフロアにある PHP 8.2 以降が必要です。
コンポーザー グローバルにはリーフ/cli -W が必要です
リーフ作成マイプロダクト
cd my-product && リーフサーブ
http://localhost:5500 を開くと、アプリが実行されています。迷路の設置も儀式もありません。
何かを構築したり、何かを壊したり、あるいは単に考えがある場合は、Discord または GitHub で私たちに知らせてください。 Leaf は常にそれを使用する人々によって形作られてきましたが、Leaf 5 はそれ以前のどのリリースよりも大切です。
私たちと一緒に構築していただきありがとうございます。これは、私たちが PHP に関してこれまでで最も興奮したことです。
何百万人ものユーザーが使用するフレームワークと製品を出荷しているスタジオである Leaf PHP の作成者によるものです。あなたのものを作りましょう。

## Original Extract

The offical blog for Leaf PHP

Leaf Blog Discord Sponsor Docs → August 13, 2026
Leaf 5: Build products at the speed of thought
Name Michael Darko Contributors + 6 contributors
Leaf 5 is the biggest release in Leaf's history , and it's not just a version bump. It's a natural evolution of what a PHP framework should be in a world where you build alongside AI: a product-first framework with project context built in, a family of tools that work in any PHP app, and a promise that your code stays clean, readable, and yours.
Every framework says it makes you faster. But "faster" has changed. Today, you're not just writing code, you're directing it. You describe a feature, your AI assistant drafts it, you review and refine. The bottleneck isn't typing anymore; it's how well your tools understand your app.
Most frameworks were never designed for that. Conventions live in people's heads, structure hides behind abstractions, and an assistant dropped into the codebase has to guess its way around.
Leaf 5 is our answer: a framework where the structure is obvious, the patterns are consistent, and the project itself carries shared context that both humans and agents can read...and write back to. You ask for a billing module, and the generated code lands where it belongs, looking like the rest of your app. No black boxes. No cleanup afterwards.
And because we're product-first, Leaf 5 ships the pieces real products need: auth, databases, queues, mail, billing and 25+ other features as production-ready modules you pull in exactly when your app asks for them. Nothing more.
This part hasn't changed, and it never will. There's no hidden runtime, no proprietary lock-in, no deployment magic. A Leaf 5 app is a PHP app. It runs anywhere PHP runs, and everything it generates is code you can read line by line.
app ( ) -> get ( '/' , function ( ) {
response ( ) -> json ( [
'message' => 'Leaf 5 😉'
] ) ;
} ) ;
app ( ) -> run ( ) ;
If you've written Leaf before, you already know Leaf 5. If you haven't, you'll know it in an afternoon.
The other big shift in Leaf 5 is that we stopped thinking of Leaf as one framework and started building it as a family of focused tools , each with its own home and each usable far beyond Leaf itself:
Alchemy : your entire QA setup in one file. Tests (Pest or PHPUnit), code style, automated refactors, static analysis, and CI pipelines for GitHub, GitLab or CircleCI, all described in a single alchemy.yml . It works with Leaf, Laravel, Symfony, Slim, or plain PHP, installs nothing until you use it, and alchemy eject lets you leave anytime with real config files.
Seedling : build CLI tools with the Leaf MVC experience. Under the hood is Sprout, our zero-fuss console framework, which just went through a full stability pass: a real test suite, piped input and Windows fixes, and CI running across macOS, Ubuntu and Windows.
Fetch : fetch() for PHP. The request API you already know from JavaScript, working in any PHP app. We rebuilt the internals for v5: every documented option now actually works, form encoding is smarter, responses are more predictable, and the whole surface is covered by tests that run against a real server on every commit.
Hana JS : our simple, lightweight React alternative, because the frontend deserves the same "no overengineering" treatment.
That's on top of 25+ modules. Grab exactly what you need and nothing else.
Why framework-agnostic? Partly because good tools deserve a bigger audience. But honestly, it's also an AI-era bet: when an assistant is choosing a library for a task, small, sharp, well-documented tools win. We want Leaf's tools to be the obvious pick whether or not the app is a Leaf app.
Leaf isn't only competing with other frameworks. It's competing with the moment where an early product starts working, users show up, and the codebase needs to become more serious without slowing you down.
The classic options both tax you at that moment. The tiny-router path gives you a fast start, then charges you product time to choose and wire auth, database patterns, validation, mail, queues, views, and structure yourself. The heavy-framework path is serious from day one, but your first version inherits more concepts and more ceremony than the product needs yet. Leaf's path is to start small, add first-party product modules when your app asks for them, move into MVC when structure earns its place, and keep the whole thing readable to humans and agents as it grows.
Choose Laravel when you want the largest PHP product ecosystem and you're happy to adopt its full application model.
Choose Symfony when you need deep architecture control, mature components, and enterprise-grade composition.
Choose Slim when you want the smallest possible HTTP layer and you're comfortable assembling the rest of the product stack yourself.
Choose Leaf when you want to get a real PHP product into users' hands quickly, keep the app small enough to understand, and still have a path to auth, data, frontend, jobs, billing, deployment, and AI-assisted maintenance as the product grows.
The product story is the point, but performance is still part of the promise. Leaf 5 keeps framework overhead low so your product gets more of the request budget:
This compares route lookup and dispatch overhead, not full application throughput: 10,000 lookups against a mixed static/dynamic route table, alternating exact and parameterized hits, averaged over five runs on PHP 8.2. Every framework resolves the same URIs, request objects for the Laravel column are pre-built outside the timed loop, and each column measures routing work only: Illuminate Routing for Laravel, the symfony/routing UrlMatcher, and Slim 4's FastRoute-backed dispatch.
The reason lookups stay flat is that Leaf 5 indexes exact routes and buckets dynamic routes by method and first segment, so exact matches never touch a pattern and dynamic requests only scan their own bucket. The same pass made repeated env reads about 41x faster and URI parsing about 4.4x faster per request.
When something breaks, the error message is the first thing you read. If you're building with an AI assistant, it's usually the only thing your assistant reads. So in Leaf 5, error messages are written the way a maintainer would answer a support ticket: what went wrong, what to do about it, and how the fix changes depending on how your app is set up.
That's a real crash screen from Leaf 5's new error engine. The remediation is the message. It knows a Leaf MVC app can run leaf key:generate and a lite app can't, and it says so, right there, before you've opened a single doc. The crash screen is built for handoff too: Copy as Markdown turns the whole report into something you can paste into any chat, and Open with AI sends it there directly.
Burning fewer of your AI tokens #
If you build with AI, tokens are money, and most frameworks are expensive to be smart about. Assistants read thousands of words of docs to extract one API, guess at conventions, and burn retries on errors that explain nothing.
Leaf 5 is careful with your assistant's context window. Reference docs live one hop away at stable URLs, with one topic per file, and each file is small enough to load whole. The Leaf skill leads with the exact mistakes agents actually make, because correcting a wrong assumption up front is hundreds of times cheaper than debugging it after. Errors that name their own fix collapse the try-fail-search-retry loop into one step. And .leaf/CONTEXT.md means your assistant never has to rediscover your architecture. It reads one small file instead of re-exploring your codebase every session.
We don't guess at any of this. AI agents build real apps on Leaf, we count every file they had to open and every retry they burned, and whatever cost them tokens becomes a fix. The error screen above exists because an agent hit the unclear version of that message and told us exactly what it needed to hear instead. More on that story soon, because it deserves its own post.
If you've visited leafphp.dev recently, you've seen the new look, and it's not just the docs. Alchemy, Seedling, Fetch, Hana, and this blog all got redesigned around a shared design language: Bricolage Grotesque headings, hairline grids, and each product keeping its own personality and colour. Different homes, same family.
The docs themselves got a deep pass too: a modern reading experience, an ecosystem menu that finally connects all the products, and content updated for everything Leaf 5 changes.
What we're thinking (the honest part) #
A few principles kept us honest while building this release, and they're worth writing down:
AI should amplify you, not replace your judgement. Leaf 5 gives assistants the context to do great work, but the output is always plain, reviewable PHP. You stay the architect.
Advertise only what works. Every claim on every new site is backed by code and tests. Where we found gaps between docs and reality (and we found a few), we fixed the code, not the copy.
Small surface, sharp edges. The framework stays light. Capability comes from modules you opt into, not from a kernel that grows forever.
Free stays free. Leaf 5, every module, and every tool above is open source and community-backed. If Leaf saves you time, sponsoring is how we keep it that way.
We're not done. On the roadmap from here: real-time APIs, better deployment stories, deeper third-party integrations, and more tooling around the shared-context workflow. We'll share each piece here as it lands.
You'll need PHP 8.2 or newer, the same floor as the rest of the modern PHP ecosystem.
composer global require leafs/cli -W
leaf create my-product
cd my-product && leaf serve
Open http://localhost:5500 and you have a running app. No setup maze, no ceremony.
If you build something, break something, or just have thoughts, come tell us on Discord or GitHub . Leaf has always been shaped by the people using it, and Leaf 5 more than any release before it.
Thank you for building with us. This is the most excited we've ever been about PHP.
From the creators of Leaf PHP, a studio that's shipped frameworks and products used by millions. Let's build yours.
