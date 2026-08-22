---
source: "https://blackboard.sh/blog/electrobun-2-0/"
hn_url: "https://news.ycombinator.com/item?id=49402641"
title: "Cottontail, Electrobun 2.0, and Why I Decided to Jian-Yang Anthropic"
article_title: "Cottontail, Electrobun 2.0, and Why I Decided to Jian-Yang Anthropic - Blackboard Blog"
image: ""
author: "yoav"
captured_at: "2026-08-22T19:16:10Z"
capture_tool: "hn-digest"
hn_id: 49402641
score: 1
comments: 0
posted_at: "2026-08-22T18:55:12Z"
tags:
  - hacker-news
  - translated
---

# Cottontail, Electrobun 2.0, and Why I Decided to Jian-Yang Anthropic

- HN: [49402641](https://news.ycombinator.com/item?id=49402641)
- Source: [blackboard.sh](https://blackboard.sh/blog/electrobun-2-0/)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T18:55:12Z

## Translation

タイトル: ワタテール、エレクトロブン 2.0、そして私が Jian-Yang Anthropic を選択した理由
記事のタイトル: Cottontail、Electrobun 2.0、そして Jian-Yang Anthropic を選択した理由 - Blackboard Blog

記事本文:
ワタテール、エレクトロブン 2.0、そして私が Jian-Yang Anthropic を選択した理由
Electrobun は、小型で高速、バッテリー内蔵のデスクトップ アプリ フレームワークで、私は過去 3 年間にわたって開発を繰り返してきました。
ドキュメント を読むことができます。また、オープンソースなので、リポジトリにアクセスして、まだスターを付けていない場合は、スターを付けてください。
もともと私は Dash の初期バージョンを構築していましたが、Electron のサイズ、パフォーマンス、配布のストーリーに非常に不満を感じたので、より良いものを構築することにしました。私はスタートアップやユニコーン企業でエンジニアおよびエンジニアリング リーダーとして 20 年以上を過ごし、10 年以上前に私が設計し独力で構築したシステムは今でも数十億のページビューと数十万人の労働日に電力を供給しています。そのため、エージェントが普及する前にクロスプラットフォームのデスクトップ アプリ フレームワークをゼロから構築することは、研究室で構築している他のものに加えて取り組むことができる合理的なサイドクエストであると考えるのに十分な傲慢さが私にありました。
私が構築しているものを発見した人たちから好意的な支持を得た後、他の人が使用できるように時間をかけて磨き上げ、2 月に Electrobun 1.x をリリースしました。これは 6 か月前で、まだほとんどがまだエージェント前でした。完全にエージェント時代に入った今、私が Electrobun 2.0 に向けて何を準備してきたかが分かるまで待ってください。
これを読んでいる Electrobun を初めて使用する人のために説明すると、基本的な考え方は、これまでそれに付随してきたすべての荷物を受け入れることなく、使い慣れたテクノロジを使用して、小型で高速なデスクトップ アプリを構築して出荷できるはずだということです。
2 月に遡ると、Electrobun 1.x では次のものが提供されました。
システム Web ビューまたは固定された Chromium。システム Webview を使用してアプリを小さく保つことも、プラットフォーム間のブラウザーの一貫性が重要な場合には特定のバージョンの CEF をバンドルすることもできます。
小さな差

熱心な更新。 Electrobun には、zstd を中心に構築されたバッテリー付属の更新システムと、LLM を学ぶ前に Zig を学ぶために元々手書きで書いた bsdiff の最適化された Zig 実装があります。可能な限り最小の配布可能なサイズで、さらに小さなアップデートも 2KB までなので、好きなだけ配布できます。
カスタム OOPIF アーキテクチャ。 <electrobun-webview> は、スーパー iframe のように動作する分離された Web ビューを提供します。当時、Electron は長らく非推奨となっていた Chromium の <webview> タグをまだ使用していたので、私は独自のタグを構築することにしました。 Electrobun では、システム Web ビューまたは CEF を使用しているかどうかに関係なく、同じ API を使用して、すべてのタブが実際に独自の Web ビュー内で分離され、同じ UI にそれらを合成するマルチタブ ブラウザのようなものを構築できます。
高速なエンドツーエンドのツールチェーン。開発、ビルド、パッケージ化、コード署名、更新、配布はフレームワークの一部でした。 S3、R2、または基本的に任意の静的ファイル ホストを導入すると、Electrobun CLI が残りを処理します。
発射は爆発的でした。数週間でスターが 10,000 個になります。 Electrobun は現在、GitHub のスター数 13,000 に近づき、数百のアプリが Electrobun を使用して構築されています。スター チャートは基本的に発売時に垂直になり、数か月間その状態が続きました。
アーキテクチャをどこまで推し進めてレンダラーのモジュール性を証明できるか試したかったので、3 月に WGPU サポートを追加しました。これは、bundleWGPU を true に設定するだけで有効になります。これにより、Dawn (Chrome の WGPU 実装) のカスタム ビルドがアプリ バンドルに配置され、TypeScript に公開されました。
Electrobun は、macOS、Windows、Linux 上でネイティブ GPU ウィンドウを作成し、Bun から直接駆動できるようになりました。 Three.js アダプターと Babylon アダプターを追加したので、ブラウザーに WebGL キャンバスを配置せずに API を使用できるようになりました。
さらに興味深いことに、<electrobun-webview> 用にすでに構築したアーキテクチャが GPU サーフでも機能することが判明しました。

aces も追加したので、 <electrobun-wgpu> を追加しました。スーパー iframe に対応するスーパー GPU サーフェス。
ネイティブ WGPU サーフェスを取得し、別の分離された Web ビューを通常の Web UI に合成するのと同じ方法で、それを通常の Web UI に直接合成できるようになりました。設定パネルは HTML、エディターは HTML、中央のビューポートはメイン プロセスから直接駆動されるネイティブ GPU サーフェスである可能性があります。
人々はこれを使用して、私が予想していなかったゲーム IDE やその他のアプリケーションを構築しました。 DOOM を 2 つの方法で Electrobun に移植し、試していただける小さな GPU 数字分類子テンプレートを作成しました。一部の開発者はさらに進んで、Electrobun の WGPU とネイティブ ライブラリを直接使用し、ウィンドウ ハンドルを渡してそこから引き継ぐカスタム Zig サイドカーを出荷しました。
1.x の後、私は一歩下がって、そもそもこのサイド クエスト全体を開始したアプリである Dash に戻りました。世界は大きく変わり、Code と NoCode を統合するワークスペースを構築するという私の当初のアイデアは時代遅れになりました。詳細については以下をご覧ください。
過去 2 か月間、私はエージェントのスピードで Electrobun と Dash を新しいプラットフォームにどのように適合させるかを構築してきました。今日、私はそのプラットフォームの前半部分である基盤を発表します。これを Electrobun 2.0 と呼んでいます。
私がここにたどり着いた経緯を説明する前に、この土地の概要を簡単にご紹介します。
Hutch は Electrobun の新しい CLI です。これは引き続きビルドとパッケージ化を処理しますが、npm の外部で動作できるようになり、Rust、Zig、Go、Odin などのすべての新しいメイン プロセス テクノロジとツールチェーンを処理できるようになりました。また、JS エコシステムには正規の非営利 npm クライアントがないため、Hutch には軽量の npm リゾルバーが組み込まれているため、hatch install を実行して npm パッケージをインストールできます。さらに必要な場合には、npm のようなより完全な npm クライアントを指定できます。

Bun、pnpm、または Yarn。
Cottontail は、JavaScriptCore と Zig を中心に書かれた新しい JavaScript ランタイムです。 Node と Bun の互換性があるため、ほとんどのパッケージが正常に動作し、パフォーマンスの点で競争力があり、macOS、Windows、Linux 上ではすでに Bun よりも小さいです。これは、Electrobun がサーバー ランタイムのアーキテクチャとロードマップを継承するのではなく、アプリケーションを中心に設計された小さな JS ランタイムを持つことができるようにするために存在します。
Electrobun 2.0 では、フレームワークのコアが Zig に移行され、高速化され、UI がシステム Web ビュー、ピン留めされた Chromium、ネイティブ WGPU、またはそれらの組み合わせを使用するかどうかに関係なく、メイン プロセスに Cottontail、Bun、Zig、Rust、Go、または Odin を選択できるようになります。
Warren は、SolidJS からインスピレーションを得た実験的なきめの細かいリアクティブ UI フレームワークで、リアクティブ性をより明示的にし、コンパイラーを必要とせず、Web ビューの DOM と、メイン プロセスで直接 Electrobun のネイティブ GPU UI プリミティブに対して直接動作します。
Electrobun 2.0 で構築された Dash は、人間と AI が一緒に構築するマルチマシン ワークスペースの新しい解釈です。現在プライベート ベータ版であり、すべての可動部分について説明するには、起動時に独自のブログ投稿が必要です。ダッシュは台の後半です。
ここで旅について少し説明します。
GUI アプリケーションには、コアとなるアーキテクチャ上の制約があります。通常、プラットフォームのイベント ループはメイン スレッドをブロックする必要があります。
Electrobun のオリジナル バージョンでは、2 つのプロセスでこれを解決しました。 Bun はアプリケーションを実行し、ネイティブ イベント ループを所有する Zig プロセスと Unix ソケット経由で通信しました。
Electrobun 1.0 までに Bun の FFI は十分に良くなり、アーキテクチャをひっくり返すことができました。 Bun は、別のスレッドで実行されている Bun ワーカーがアプリを実行している間、ネイティブ イベント ループを実行するメイン スレッドをブロックしました。

ication コードと FFI がネイティブ ライブラリに直接組み込まれます。 Bun ワーカーはランタイムのメモリ内の読み取り専用コピーを共有するため、追加のスレッドは基本的に無料でした。Electrobun のアップデータの背後にあるカスタム bsdiff 実装やその他のパフォーマンスに依存する部分など、Zig が意味をなす場所は維持しながら、Unix ソケットと元の Zig レイヤーを削除する必要がありました。
2.0 では、Zig Core を再導入しましたが、理由も方法も異なりました。
すでに Rust、Zig、Go サイドカーを出荷していて、アプリで Bun ランタイムを使用していなかった開発者に最高のエクスペリエンスを提供するために、コア状態とネイティブ アプリケーション ロジックを Bun から新しい Zig コアに移動しました。 Bun 側に残ったものは、同じコアの薄い SDK ラッパーになりました。次に、Rust、Zig、Go、Odin の SDK を追加しました。
同じ Electrobun マジックとコアを、幅広いメイン プロセス テクノロジとレンダリング レイヤーで利用できるようになりました。 Cottontail + TypeScript + システム Web ビューを使用すると、小型で専用に構築されたものを除けば、Electrobun の人々がすでに知っているようなものを手に入れることができます。
Zig + システムの Web ビューを使用して、数 MB 程度の非圧縮の Electrobun アプリケーション全体を配布できます。
Rust + Chromium、または Go + システム Web ビューを使用するか、Odin でメイン プロセスを記述して WGPU ウィンドウを駆動し、JavaScript ランタイムを起動したり Web ビューを開いたりしないクロスプラットフォームのリアクティブ GUI アプリケーションを構築できます。
これらを組み合わせて高度な IDE を構築することもできます。通常の HTML アプリケーションには、Rust、Zig、Go、または Odin を使用して、分離された子 Web ビューとネイティブ GPU サーフェスを組み込むことができます。
メインプロセスの作業をしているときに、デフォルトの TypeScript オプションが実際にどのようなものであるべきかを考え始めました。私は持っていました

Bun のサイズが大きくなり、Windows 上で突然 100MB を超え、さらに追加し続けることに多くの人が不満を抱いていました。現在、最大 3 つの SQLite クライアントがあると思います。私は一般的に Bun のバッテリーを含む哲学には同意しますが、すべてのアプリが SQLite、MySQL、Postgres、Redis のクライアントを持つべきだという点には同意しません。どこで終わるのでしょうか?
私は当初、Bun をバッテリー付属の JS オプションとして維持し、Bun のすべてを必要としないアプリケーション向けに、Zig + QuickJS を中心とした小さな実験的なランタイムを構築する予定でした。しかし、『ソル ウルトラ』、『ファブル』、そして『キミ 3』では、数年先を飛ばして逆にやらせてくれました。最初は Bun フォークのようなものから始めて、時間の経過とともに逆方向に作業して、より小さく、より構成しやすくすることができます。説明しましょう。
Cottontail + QuickJS は驚くほど遠くまで到達しました。 Electrobun Kitchen Sink のテストに合格し、Dash が実行されました。
サイズは約 1 MB と小さかったですが、実際のアプリケーションで体験してみると、出荷できないことが明らかになりました。 JSC JIT は現代の最も素晴らしいテクノロジーの 1 つであるため、QuickJS 上に独自の JIT を記述することは、私が望んでいたよりもはるかに大きな副次的課題でした。より良い JIT を再発明する余地はあまりなく、またそれは私個人の範囲ではなかったので、Cottontail の QuickJS を JavaScriptCore に置き換えました。
私は Electrobun 2.0 と Dash 自体の両方をビルドするために Dash を使用していますが、Bun で実行される Dash と Cottontail で実行される Dash の違いが分かりませんでした。
昨日リリースされた Bun 1.4.0 と現在のリリースの比較は次のとおりです。
Cottontail はすでにどこでも小さくなり、いくつかのパスで競争力がありますが、起動、メモリ、実行時のパフォーマンスに関しては、やるべきことがまだたくさんあります。
当初の計画では、Cottontail は小規模で新しいものとし、ユーザーには Wi-Fi との互換性を考慮して Bun を選択させる予定でした。

それは生態系です。
隣り合った 2 つのランタイムを見れば見るほど、意味がわからなくなりました。
Bun はすでにノードの互換性を仕様として扱っています。 Node の API とテストを採用し、Node 用に書かれたほとんどのソフトウェアが Bun でも実行できるように十分な互換性を実装しました。 Deno は Node で同じことを行い、さらには Electrobun の新しい Deno Desktop でも同様のことを行いました。オープンソースはサークルです。
では、Cottontail が実際に最初から Node と Bun 互換であり、そこから小型のモジュール式で Electrobun 用に構築されたビジョンに移行したとしたらどうなるでしょうか。
これは特に面白い瞬間に起こりました。私がそれについて考えている間に、Anthropic が Bun を買収し、Bun の 100 万行を超える Zig 実装が AI エージェントを使用して Rust で書き直されたからです。今では 1 週間で数百万行のコードを書き換えることができ、Rust Bun の「歴戦の」時計はちょうどゼロにリセットされました。
私がJian-Yang Anthropicを選んだ理由
Rust Bun、そのマージの早さ、Zig、Rust、AI 生成コード、レビューの実践、セグメンテーション違反、そしてこれらのいずれかが良いアイデアであるかどうかについては、たくさんの議論がありました。ここでインターネット上の議論全体を再現するつもりはありません。ただ、みんなが間違っているとだけ言っておきます。
Electrobun の場合、私がすでに解決しようとしていた依存関係の問題がほぼ明確になりました。
Bun a f のソフトフォークを試してみました

[切り捨てられた]

## Original Extract

Cottontail, Electrobun 2.0, and Why I Decided to Jian-Yang Anthropic
Electrobun is a small, fast, batteries-included desktop app framework that I’ve been working on and off on for the past three years.
You can read the docs , and since it’s open source, visit the repo and give it a star if you haven’t already.
Originally I was building an early version of Dash and got so frustrated with Electron’s size, performance, and distribution story that I decided to build something better. I’ve spent 20+ years as an engineer and engineering leader at startups and unicorns, systems I’ve architected and single-handedly built 10+ years ago still power billions of pageviews and hundreds of thousands of people’s working days. That gave me just enough hubris to think building a cross-platform desktop app framework from scratch before agents were a thing was a reasonable side quest that I could take on in addition to the other things I’m building at my lab.
After getting positive traction from people who discovered what I was building, I took some time to polish it up enough for other people to use and released Electrobun 1.x in February, that was 6 months ago and still mostly pre-agent. Wait until you see what I’ve been cooking up for Electrobun 2.0 now that we’re fully in the agentic era.
For those reading this that are new to Electrobun, the basic idea is that you should be able to build and ship a tiny, fast desktop app, with familiar technologies without accepting all the baggage that has historically come with doing that.
Back in February Electrobun 1.x gave you:
System webviews or pinned Chromium. You could use the system webview and keep your app tiny, or bundle a specific version of CEF when cross-platform browser consistency mattered.
Tiny differential updates. Electrobun has a batteries-included update system built around zstd and an optimized Zig implementation of bsdiff I originally wrote by hand to learn Zig before LLMs. The tiniest possible distributable size, and even tinier updates as small as 2KB so you can ship as often as you like.
A custom OOPIF architecture. <electrobun-webview> gives you isolated webviews that behave like super-iframes. At the time Electron was still using Chromium’s long-deprecated <webview> tag and I set out to build my own. In Electrobun you can build something like a multi-tab browser where every tab is actually isolated in its own webview and composite them into the same UI, using the same API whether you’re using system webviews or CEF.
A fast end-to-end toolchain. Dev, build, packaging, code signing, updates and distribution were part of the framework. Bring S3, R2 or basically any static file host and the Electrobun CLI handled the rest.
The launch was explosive. 10k stars in a few weeks. Electrobun is now approaching 13k GitHub stars, hundreds of apps have been built with it, and the star chart basically went vertical at launch and stayed that way for months.
I wanted to see how far I could push the architecture and prove out the renderer modularity so in March I added WGPU support, enabled by just setting bundleWGPU to true. This put a custom build of Dawn (Chrome’s WGPU implementation) in your app bundle and exposed it to TypeScript.
Electrobun could now create native GPU windows on macOS, Windows, and Linux and drive them directly from Bun. I added Three.js and Babylon adapters so you could use their APIs without putting a WebGL canvas in a browser.
More interestingly, the architecture I’d already built for <electrobun-webview> turned out to work for GPU surfaces too, so I added <electrobun-wgpu> . Super GPU surfaces to go with your super-iframes.
You could now take a native WGPU surface and composite it directly into a normal web UI in the same way you could composite another isolated webview into it. Your settings panel might be HTML, your editor might be HTML, and the viewport in the middle could be a native GPU surface driven directly from the main process.
People used this to build game IDEs and other applications I hadn’t anticipated. I ported DOOM to Electrobun two ways and built a little GPU digit classifier template you can try out. Some developers went further and shipped custom Zig sidecars that used Electrobun’s WGPU and native libraries directly, passing window handles around and taking over from there.
After 1.x I took a step back and went back to Dash, the app that started this whole side quest in the first place. The world had materially changed and my original idea of building a workspace that unified Code and NoCode was obsolete. More on that below.
For the last two months I’ve been building at agent speed shaping how Electrobun and Dash fit together into a new platform. Today I’m announcing the first half of that platform, the foundation, which I’m calling Electrobun 2.0.
Before we dive into how I got here I’ll give you a quick lay of the land.
Hutch is Electrobun’s new CLI. It still handles building and packaging but can now operate outside of npm, handle all the new main process technologies and toolchains like Rust, Zig, Go, and Odin, and since the JS ecosystem doesn’t have a canonical nonprofit npm client, Hutch has a lightweight npm resolver built in so you can do hutch install to install your npm packages. And for times when you need more you can specify a fuller npm client like npm, Bun, pnpm, or Yarn.
Cottontail is a new JavaScript runtime written around JavaScriptCore and Zig. It has Node and Bun compatibility where it makes sense so most packages should just work, is competitive performance-wise and already smaller than Bun on macOS, Windows, and Linux. It exists so Electrobun can have a small JS runtime designed around applications rather than inherit the architecture and roadmap of a server runtime.
Electrobun 2.0 moves the core of the framework into Zig so it’s faster, and lets you choose Cottontail, Bun, Zig, Rust, Go, or Odin for your main process, independently of whether your UI uses system webviews, pinned Chromium, native WGPU, or some combination of them.
Warren is an experimental fine-grained reactive UI framework inspired by SolidJS that makes reactivity more explicit, doesn’t require a compiler, and works both in the DOM in a webview and directly against Electrobun’s native GPU UI primitives directly in the main process.
Dash, built with Electrobun 2.0 is a new take on a multi-machine workspace for humans and AI to build together, it’s currently in private beta and needs its own blog post when I launch it to talk about all the moving parts. Dash is the second half of the platform.
Now a little about the journey:
GUI applications have a core architectural constraint: the platform event loop generally needs to block the main thread.
The original version of Electrobun solved this with two processes. Bun ran your application and communicated over Unix sockets with a Zig process that owned the native event loop.
By Electrobun 1.0 Bun’s FFI had gotten good enough that I could flip the architecture upside down. Bun blocked its main thread running the native event loop while a Bun worker running in a separate thread ran your application code and FFI’d directly into the native libraries. Bun workers share an in-memory read-only copy of the runtime, so the additional thread was basically free, and I got to delete the Unix sockets and that original Zig layer while keeping the places where Zig made sense, like the custom bsdiff implementation behind Electrobun’s updater and some of the other performance-sensitive pieces.
For 2.0 I reintroduced a Zig Core, but for a different reason and in a different way.
In order to enable a first-class experience for developers that were already shipping Rust, Zig, and Go sidecars and had no use for the Bun runtime in their apps I moved the core state and native application logic out of Bun and into a new Zig core. What was left on the Bun side became a thin SDK wrapper around the same core. I then added SDKs for Rust, Zig, Go, and Odin.
You can now leverage the same Electrobun magic and core with a wide range of main process technologies and rendering layers. You can use Cottontail + TypeScript + system webviews and have something that feels like the Electrobun people already know, except smaller and purpose built.
You can use Zig + system webviews and ship an entire uncompressed Electrobun application as small as a few MB.
You can use Rust + Chromium, or Go + system webviews, or write your main process in Odin and drive WGPU windows, and build a cross-platform reactive GUI application that never starts a JavaScript runtime or opens a webview.
You can also mix these things and build advanced IDEs. A normal HTML application can have isolated child webviews and native GPU surfaces composited into it using Rust, Zig, Go, or Odin.
While I was doing the main-process work I started thinking about what the default TypeScript option should actually look like. I had lots of people complaining about how big Bun was getting, suddenly over 100MB on Windows and they just kept adding things. I think they’re up to 3 SQLite clients now. While I generally agree with Bun’s batteries included philosophy, I disagree that every app should have a client for SQLite, MySQL, Postgres, and Redis, where does it end?
I’d originally planned to keep Bun as the batteries-included JS option and build a tiny experimental runtime around Zig + QuickJS for applications that didn’t need all of Bun. But Sol Ultra, Fable, and Kimi 3 let me skip a few years ahead and do it in reverse. I could start with a sort of Bun fork and work backwards to make it smaller and more composable over time. Let me explain.
Cottontail + QuickJS got surprisingly far. I had the Electrobun Kitchen Sink tests passing and Dash running on it.
It was tiny at around 1MB, but experiencing it in a real application made it obvious I couldn’t ship it. The JSC JIT is one of the most incredible pieces of technology of the modern era so writing my own JIT on top of QuickJS was a much bigger side quest than I wanted. There’s not a lot of room for reinventing a better JIT, nor was that in scope for me as a single person, so I replaced QuickJS in Cottontail with JavaScriptCore.
I use Dash to build both Electrobun 2.0 and Dash itself, and I couldn’t tell the difference between Dash running on Bun and Dash running on Cottontail.
Here is the current release comparison against Bun 1.4.0 , released yesterday:
Cottontail is already smaller everywhere and competitive in several paths, there is still a lot of work to be done on startup, memory, and runtime performance.
At first the plan was still for Cottontail to be something small and new while letting people choose Bun for compatibility with the ecosystem.
The more I looked at the two runtimes sitting next to each other the less that made sense.
Bun already treats Node compatibility as a spec. It took Node’s APIs and tests and implemented enough compatibility that most software written for Node also runs in Bun. Deno did the same with Node, and even Electrobun for their new Deno Desktop . Open Source is a circle.
So what if Cottontail was actually Node and Bun compatible from day one, and moved toward the tiny, modular, built-for-Electrobun vision from there.
This happened at a particularly funny moment because while I was thinking about it Anthropic acquired Bun and Bun’s million-plus-line Zig implementation was rewritten in Rust using AI agents. You could now rewrite millions of lines of code in a week, and Rust Bun’s “battle-tested” clock was just reset to zero.
Why I decided to Jian-Yang Anthropic
There has been plenty of arguing about Rust Bun, how quickly it was merged, Zig, Rust, AI-generated code, review practices, segfaults, and whether any of this is a good idea. I’m not going to reproduce the entire internet argument here. I will just say that everyone is wrong.
For Electrobun it mostly clarified a dependency problem I was already trying to solve.
I’d experimented with soft-forking Bun a f

[truncated]
