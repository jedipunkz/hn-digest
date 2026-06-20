---
source: "https://www.porchlab.com/blog/best-ai-stack-elixir-phoenix/"
hn_url: "https://news.ycombinator.com/item?id=48612306"
title: "The best stack for the AI Era"
article_title: "Why Elixir and Phoenix are the Best Stack for the AI Era — PorchLab"
author: "wallflow3r"
captured_at: "2026-06-20T21:34:09Z"
capture_tool: "hn-digest"
hn_id: 48612306
score: 2
comments: 0
posted_at: "2026-06-20T19:39:22Z"
tags:
  - hacker-news
  - translated
---

# The best stack for the AI Era

- HN: [48612306](https://news.ycombinator.com/item?id=48612306)
- Source: [www.porchlab.com](https://www.porchlab.com/blog/best-ai-stack-elixir-phoenix/)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T19:39:22Z

## Translation

タイトル: AI 時代に最適なスタック
記事タイトル: Elixir と Phoenix が AI 時代に最適なスタックである理由 — PorchLab
説明: Elixir と Phoenix が、大規模な同時実行、楽なテキスト ストリーミング、および迅速な反復を提供する、AI 時代の究極の技術スタックである理由を説明します。

記事本文:
Elixir と Phoenix が AI 時代に最適なスタックである理由 — PorchLab
コンテンツにスキップ
ポーチラボ
検索
⌘K
←
ブログに戻る
Elixir と Phoenix が AI 時代に最適なスタックである理由
Elixir と Phoenix が、大規模な同時実行性、楽なテキスト ストリーミング、および迅速な反復を提供する、AI 時代の究極の技術スタックである理由をご覧ください。
過去 1 年間、AI 統合アプリケーションの構築に費やした人なら、おそらくアーキテクチャ上の悪夢が繰り返されていることに気づいたでしょう。
私たちは常にネットワーク通話を待っています。私たちは何千もの永続的な WebSocket 接続をやりくりして、せっかちなユーザーにトークンをストリーミングして返しています。そして、これらすべてに対処するために、複雑な非同期キュー、Redis インスタンス、肥大化したフロントエンド フレームワークを従来の MVC またはマイクロサービス バックエンドにボルトで固定しています。
私たちは、2014 年のステートレスなリクエスト/レスポンス Web 向けに設計されたツールを使用して、次世代の高度な同時実行 AI アプリケーションを構築しようとしています。
何年にもわたって大規模なシステムの設計を行ってきた結果、生成 AI 時代に構築するための最も実用的な選択肢は、あなたが期待するようなスタックではないことに気づきました。それは Elixir と Phoenix フレームワークです。
1. 微細なフットプリントによる大規模な同時実行
AI ラッパー、エージェント、またはマルチユーザー チャット インターフェイスを構築する場合、アプリケーションは本質的に I/O バウンドになります。 LLM プロバイダーの API が応答するのを待つことにほとんどの時間を費やしています。 Ruby、Python、Node.js などの従来の環境では、外部 API を待機している間、何千もの接続を開いたままにするには、イベント ループ、ワーカー プール、メモリの肥大化など、インフラストラクチャの真剣な取り組みが必要です。
Elixir は Erlang VM (BEAM) 上で動作します。BEAM は文字通り、何百万もの同時通話を汗をかくことなく処理するために通信エンジニアによって発明されました。
Elixir では、すべてのユーザーが

すべてのバックグラウンド タスク、およびすべての LLM API 呼び出しは、独自の軽量プロセスで実行されます。これらは OS レベルのスレッドではありません。文字通りキロバイトのメモリを消費します。単一の安価なサーバー上で AI エージェントとチャットする数万人の同時ユーザーを簡単に管理できます。 Kubernetes クラスターは必要ありません。トラフィックの急増に対処するためだけに複雑な自動スケーリング ルールは必要ありません。このシステムは同時実行性をネイティブに吸収するだけで、インフラストラクチャの設置面積は非常に小さくなります。
2. テキストのストリーミングに最適なエンジン
「ChatGPT 効果」はユーザーの期待を永久に変えました。完全な応答が得られるまでユーザーを 15 秒も待たせることはできません。出力トークンごとにストリーミングする必要があります。
標準スタックでは、これを正しく行うのは驚くほど困難です。複雑な状態を管理するフロントエンド フレームワーク (React/Vue)、ストリームを処理する WebSocket サーバー、バックエンド コンテナ間ですべての同期を保つメッセージ ブローカー (Redis など) が必要です。分散システムでは、画面に文字を表示するだけでも頭痛の種になります。
LiveView はサーバー上に状態を保持し、単一の多重化された WebSocket を介して HTML 更新をクライアントに直接プッシュします。バックエンドが OpenAI API から新しいテキスト トークンを受け取ると、サーバー側の状態を更新するだけです。 Phoenix は UI の差分を自動的に取得し、新しいトークンだけをネットワーク経由でブラウザにプッシュします。
フロントエンドの状態管理はありません。個別の GraphQL API はありません。メッセージブローカーはありません。 Phoenix はフレームワークに直接組み込まれた PubSub を配布しているため、AI が生成したストリームを 1 人のユーザー (または共有の共同作業ルームにいる 1 万人のユーザー) にブロードキャストすることは、機能的にはワンライナーです。
3. エコシステムの健全性: Elixir と Node.js の移動ターゲット
動きの速い AI スタートアップを構築するとき、自分自身の利益を得るために戦うことには最も時間をかけたくないことです。

オルたち。今すぐ Node.js エコシステムを見てください。アイデンティティの危機が永遠に続く状態にあります。 Node、Deno、Bun のいずれかを選択する必要があります。ライブラリが CommonJS または ESM をサポートしているかどうかを確認する必要があります。 Next.js のようなメタフレームワークを ORM、WebSocket ライブラリ、バックグラウンド ワーカー キューと組み合わせて、次の npm メジャー アップグレード後にすべてが適切に動作することを期待する必要があります。これは、破壊的な変化と断片化されたコミュニティ標準の骨の折れるトレッドミルです。
それに比べて、Elixir は絶対的なアーキテクチャの健全性を提供します。
Elixir のコアツールは非常に成熟しており、安定しています。 Phoenix は単なるルーターではありません。これは、10 年間にわたって予想どおりに進化してきた、バッテリーを備えた統合されたフレームワークです。データの永続性は、JavaScript ORM のオブジェクト リレーショナル インピーダンスの不一致の影響を受けない堅牢なデータベース ラッパーである Ecto によって処理されます。
Elixir は BEAM 上に構築されているため、バックグラウンド ジョブ処理 (Oban 経由) やパブリッシュ/サブスクライブ メッセージングなど、Node ワールドでサードパーティのインフラストラクチャを必要とする機能は、アプリケーション メモリ プール内で直接実行されます。 5 年前に書かれたコードが現在でも完璧にコンパイルされ、実行される、単一の強固なエコシステムが得られます。ビルド ツールのリファクタリングや依存関係地獄の管理ではなく、AI 機能の出荷にエネルギーを 100% 集中できます。
4. 「AI スピード」での反復: モノリスの場合
生成 AI の状況は、毎週のように私たちの足元で変化しています。新しい基盤モデルは絶えず減少し、ユーザーの期待は光の速さで進化しており、月曜日には確かな製品アイデアのように見えたものでも、木曜日には OpenAI または Anthropic のアップデートによって時代遅れになる可能性があります。
この非常に不安定な環境において、Product-Market Fit (PMF) を見つけるのは過酷な競争です。ゆっくり最適化することではありません

安定した機能セット。それは根本的で素早い実験です。 RAG パイプラインを完全に取り除き、コア ユースケース全体を方向転換し、データベース スキーマを根本的に変更する必要が、すべて同じ週に必要になる場合があります。
スタートアップがマイクロサービスの複雑な Web 上に構築されている場合、この種の俊敏性は数学的に不可能です。新しい仮説をテストするためだけに、データベース サービスのスキーマ変更、Python AI ワーカーの即時更新、および 3 つの異なるリポジトリと CI/CD パイプラインにわたる React フロントエンドの UI のオーバーホールを調整する余裕はありません。マイクロサービスは大規模な専門化されたエンジニアリング チーム向けに最適化しますが、PMF を探している間は反復速度が大幅に低下します。
Phoenix はモノリスを奨励し、スタック全体を単一のコードベースにまとめます。 AI エージェントが会話メモリを保存し、その新しい形式をすぐにユーザーにストリーミングする方法を大幅に変更する必要がありますか? 1 か所で変更を加え、1 つのプル リクエストを開き、1 つのアーティファクトをデプロイします。
ただし、過去の複雑なスパゲッティ コードベースとは異なり、Elixir はドメイン間に厳密で論理的な境界を強制するための組み込みツール (コンテキストなど) を提供します。変化する AI 市場に合わせて製品を一晩で方向転換できる機敏性と、実際に金メダルを獲得した後に拡張できるアーキテクチャの健全性が得られます。
5. 秘密兵器: AI はより優れた Elixir を作成します
これは通常眉をひそめるポイントです。 PythonってAIの言語じゃないの？
モデルのトレーニングとデータ サイエンス パイプラインの実行には、もちろんです。しかし、これらのモデルを使用して堅牢な Web アプリケーションを構築する場合、Elixir には大きな利点があります。それは、大規模な言語モデルの記述が非常に優れているということです。
これは単なる逸話的な観察ではありません。それは厳格なベンチマークによって裏付けられています。 2025 年後半、テンク

ent の AI チームは、20 言語にわたる 4,000 近くの現実世界のプログラミング問題で構成される大規模なベンチマークである AutoCodeBench 調査を発表しました。結果は驚くべきものでした。Elixir は、テストされた言語の中で最高の完了率を記録しました。
30 を超える評価モデルの出力を組み合わせると、Elixir の問題の 97.5% が少なくとも 1 つのモデルで解決されました。個々のモデルの性能を見ても、Elixirがトップに君臨した。たとえば、Claude Opus 4 は、Elixir タスクで 80.3% の成功率を達成しました。これは全体で最高のスコアであり、C# (74.9%) や Kotlin (72.5%) などのトレーニング データがはるかに多い言語を決定的に上回りました。
Python の市場シェアのほんの一部しかない言語が、どのようにしてこれほどパフォーマンスが向上するのでしょうか?それはアーキテクチャに帰着します。
不変性によりローカル推論が可能になります。LLM は、隠れた状態、複雑なオブジェクト指向の継承ツリー、および予測不可能な副作用に大きく悩まされます。 Elixir では、データは不変です。 AI は、関数がオブジェクトを 5 層の深さで変更するかどうかを推測する必要はありません。単純にデータを受け取り、データを返します。
明示的なデータ フロー: パイプ演算子 ( |> ) は、上から下へのデータのクリーンな線形変換を強制します。 LLM の生成と拡張は非常に予測可能です。
防御的なコーディングよりもパターン マッチング: コーディング エージェントを混乱させやすい多数のネストされた if/else チェックを記述する代わりに、Elixir は厳密なパターン マッチングに依存して、「成功」タプルと「エラー」タプルの両方をクリーンに処理します。
Claude または GPT-4 に Elixir モジュールの作成を依頼すると、言語の基本的なルールにより、AI はクリーンで分離された予測可能なコードを生成するように強制されます。境界が非常に明確で、ミックス形式のスタイルが普遍的であるため、AI で生成された Elixir コードは、他のどの言語よりも最初の試行ではるかに頻繁に機能します。

私のキャリアで使用されていました。
AI ブームは、高度にインタラクティブなリアルタイム アプリケーションを迅速に出荷できるチームに恩恵をもたらしています。接続プールの管理、ストリーミング テキストの React レンダリング サイクルのデバッグ、マイクロサービス デプロイメントの調整、アイドル状態の WebSocket 接続を処理するためだけにクラウド インフラストラクチャをスケールアップするなどでチームが行き詰まっている場合は、間違った問題でサイクルを無駄にしていることになります。
Elixir と Phoenix は、テレコム スイッチの同時実行能力、複雑な SPA のリアルタイム インタラクティブ性、不安定な市場で PMF を見つけるために必要な迅速な反復速度、および AI コーディング アシスタントと美しく連携する機能パラダイムを提供します。これは疑いもなく、現代のウェブの不当な利点です。

## Original Extract

Discover why Elixir and Phoenix are the ultimate tech stack for the AI era, offering massive concurrency, effortless text streaming, and rapid iteration.

Why Elixir and Phoenix are the Best Stack for the AI Era — PorchLab
Skip to content
porchlab
Search
⌘K
←
Back to blog
Why Elixir and Phoenix are the Best Stack for the AI Era
Discover why Elixir and Phoenix are the ultimate tech stack for the AI era, offering massive concurrency, effortless text streaming, and rapid iteration.
If you’ve spent the last year building AI-integrated applications, you’ve probably noticed a recurring architectural nightmare.
We are constantly waiting on network calls. We are juggling thousands of persistent WebSocket connections to stream tokens back to impatient users. And to handle all of this, we are bolting complex asynchronous queues, Redis instances, and bloated frontend frameworks onto our traditional MVC or microservice backends.
We’re trying to build next-generation, highly concurrent AI applications using tools designed for the stateless request/response web of 2014.
After years of architecting systems at scale, I’ve realized that the most pragmatic choice for building in the generative AI era isn’t the stack you might expect. It’s Elixir and the Phoenix framework.
1. Massive Concurrency with a Microscopic Footprint
When building AI wrappers, agents, or multi-user chat interfaces, your application is inherently I/O bound. You are spending most of your time waiting for an LLM provider’s API to respond. In traditional environments like Ruby, Python, or Node.js, holding open thousands of connections while waiting for an external API requires serious infrastructure gymnastics — event loops, worker pools, and memory bloat.
Elixir runs on the Erlang VM (BEAM), which was literally invented by telecom engineers to handle millions of concurrent phone calls without dropping a sweat.
In Elixir, every user connection, every background task, and every LLM API call runs in its own lightweight process. These aren’t OS-level threads; they take up literal kilobytes of memory. You can easily manage tens of thousands of concurrent users chatting with AI agents on a single, cheap server. No Kubernetes clusters required. No complex autoscaling rules just to handle a traffic spike. The system simply absorbs the concurrency natively, leaving a remarkably small infrastructure footprint.
2. The Ideal Engine for Streaming Text
The “ChatGPT effect” has permanently changed user expectations. You cannot make a user wait 15 seconds for a complete response; you have to stream the output token-by-token.
In a standard stack, this is surprisingly hard to get right. You need a frontend framework (React/Vue) managing complex state, a WebSocket server handling the streams, and a message broker (like Redis) keeping everything in sync across your backend containers. It’s a distributed systems headache just to put words on a screen.
LiveView keeps the state on the server and pushes HTML updates directly to the client over a single, multiplexed WebSocket. When your backend receives a new text token from the OpenAI API, you simply update the server-side state. Phoenix automatically diffs the UI and pushes just the new token across the wire to the browser.
No frontend state management. No separate GraphQL API. No message brokers. Because Phoenix has distributed PubSub built directly into the framework, broadcasting AI-generated streams to one user — or ten thousand users in a shared collaborative room — is functionally a one-liner.
3. Ecosystem Sanity: Elixir vs. the Node.js Moving Target
When you build a fast-moving AI startup, the last thing you want to spend time on is fighting your own tools. Look at the Node.js ecosystem right now. It is in a perpetual state of identity crisis. You have to choose between Node, Deno, or Bun. You have to figure out if your libraries support CommonJS or ESM. You have to stitch together a meta-framework like Next.js with an ORM, a WebSocket library, and a background worker queue, hoping they all play nice after the next major npm upgrade . It’s an exhausting treadmill of breaking changes and fragmented community standards.
Elixir offers absolute architectural sanity by comparison.
The core tooling in Elixir is remarkably mature and stable. Phoenix isn’t just a router; it’s a cohesive, battery-included framework that has evolved predictably for a decade. Data persistence is handled by Ecto, a robust database wrapper that doesn’t suffer from the object-relational impedance mismatch of JavaScript ORMs.
Because Elixir is built on the BEAM, features that require third-party infrastructure in the Node world — like background job processing (via Oban) or pub/sub messaging — run directly inside the application memory pool. You get a single, rock-solid ecosystem where code written five years ago still compiles and runs flawlessly today. You can focus 100% of your energy on shipping AI features, rather than refactoring build tools and managing dependency hell.
4. Iterating at “AI Speed”: The Case for the Monolith
The generative AI landscape is shifting underneath our feet on a weekly basis. New foundation models drop constantly, user expectations are evolving at lightspeed, and what looked like a solid product idea on Monday might be rendered obsolete by an OpenAI or Anthropic update on Thursday.
In this hyper-volatile environment, finding Product-Market Fit (PMF) is a brutal race. It’s not about slowly optimizing a stable feature set; it’s about radical, rapid experimentation. You might need to completely rip out a RAG pipeline, pivot your entire core use case, and fundamentally alter your database schema — all in the same week.
If your startup is built on a complex web of microservices, this kind of agility is mathematically impossible. You cannot afford to coordinate a schema change in a database service, a prompt update in a Python AI worker, and a UI overhaul in a React frontend across three different repositories and CI/CD pipelines just to test a new hypothesis. Microservices optimize for massive, specialized engineering teams, but they absolutely slaughter your iteration speed when you are still searching for PMF.
Phoenix encourages the monolith, bringing your entire stack into a single codebase. Need to drastically change how an AI agent stores conversational memory and immediately stream that new format to the user? You make the change in one place, open one pull request, and deploy one artifact.
But unlike the tangled spaghetti codebases of the past, Elixir provides built-in tools (like Contexts) to enforce strict, logical boundaries between your domains. You get the agility to pivot your product overnight to match the shifting AI market, along with the architectural soundness to scale it once you actually strike gold.
5. The Secret Weapon: AI Writes Better Elixir
This is the point that usually raises eyebrows. Isn’t Python the language of AI?
For training models and running data science pipelines, absolutely. But for building robust web applications using those models, Elixir holds a massive advantage: Large Language Models are exceptionally good at writing it.
This isn’t just an anecdotal observation; it’s backed by rigorous benchmarking. In late 2025, Tencent’s AI team published the AutoCodeBench study , a large-scale benchmark comprising nearly 4,000 real-world programming problems across 20 languages. The results were staggering: Elixir scored the highest completion rate of any language tested.
When combining the output of over 30 evaluated models, 97.5% of Elixir problems were solved by at least one model. Even looking at individual model performance, Elixir reigned supreme. Claude Opus 4, for example, achieved an 80.3% success rate on Elixir tasks — its highest score overall, decisively beating languages with vastly more training data like C# (74.9%) and Kotlin (72.5%).
How does a language with a fraction of Python’s market share perform so much better? It comes down to architecture:
Immutability enables local reasoning: LLMs struggle profoundly with hidden state, complex object-oriented inheritance trees, and unpredictable side effects. In Elixir, data is immutable. The AI doesn’t have to guess if a function mutates an object five layers deep; it simply takes data in and returns data out.
Explicit Data Flow: The pipe operator ( |> ) forces a clean, linear transformation of data from top to bottom. It’s highly predictable for an LLM to generate and extend.
Pattern Matching over Defensive Coding: Instead of writing dozens of nested if/else checks that easily confuse coding agents, Elixir relies on strict pattern matching to handle both “success” and “error” tuples cleanly.
When you ask Claude or GPT-4 to write an Elixir module, the fundamental rules of the language force the AI to produce clean, decoupled, and predictable code. Because the boundaries are so explicit and the mix format styling is universal, AI-generated Elixir code works on the first try far more often than in any other language I’ve used in my career.
The AI boom is rewarding teams that can ship highly interactive, real-time applications quickly. If you are bogging your team down in managing connection pools, debugging React rendering cycles for streaming text, coordinating microservice deployments, and scaling up cloud infrastructure just to handle idle WebSocket connections, you are wasting cycles on the wrong problems.
Elixir and Phoenix give you the concurrent power of a telecom switch, the real-time interactivity of a complex SPA, the rapid iteration speed required to find PMF in a volatile market, and a functional paradigm that pairs beautifully with AI coding assistants. It is, without a doubt, the unfair advantage of the modern web.
