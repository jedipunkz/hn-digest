---
source: "https://pacslate.com"
hn_url: "https://news.ycombinator.com/item?id=48736591"
title: "Show HN: Pacific Slate – a self-hosted, model-agnostic multi-agent AI assistant"
article_title: "Pacific Slate: a personal AI system"
author: "badwx"
captured_at: "2026-06-30T18:35:04Z"
capture_tool: "hn-digest"
hn_id: 48736591
score: 1
comments: 0
posted_at: "2026-06-30T17:59:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pacific Slate – a self-hosted, model-agnostic multi-agent AI assistant

- HN: [48736591](https://news.ycombinator.com/item?id=48736591)
- Source: [pacslate.com](https://pacslate.com)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T17:59:28Z

## Translation

タイトル: Show HN: Pacific Slate – 自己ホスト型でモデルに依存しないマルチエージェント AI アシスタント
記事のタイトル: Pacific Slate: パーソナル AI システム
説明: 自己ホスト型でモデルに依存しないパーソナル AI システム — アーキテクチャとその背後にある設計上の決定。

記事本文:
">
パシフィックスレート //
ライブデモ
モニター
ソース
ページをコピーする
▾
マークダウンとしてコピー
マークダウンを開く
クロードでチャット
ChatGPTで開く
Pacific Slate: パーソナル AI システム
Pacific Slate は、私自身の知識と私がフォローしている情報源を基に動作し、質問されたときに答えるだけでなく、バックグラウンドで有用な作業を行うパーソナル AI システムです。
これはプライベート リモート サーバー上で実行され、セッション全体のコンテキストを記憶し、作業をルーティングする特殊なエージェントから構築されています。このページは、それが何であるか、どのようにまとめられているのか、プライバシーをどのように処理しているのかをわかりやすくまとめたものです。興味があれば、技術的な詳細は最後に記載されています。
前もって言っておきますが、私はオペレーター兼システムアーキテクトであり、職業ソフトウェアエンジニアではありません。私は、さまざまな AI プラットフォーム、エージェント、コーディング ツールと連携して、アーキテクチャを設計し、モデルのルーティング、信頼性、コストを決定し、AI ネイティブに構築しました。
それが私にとって日々何をもたらしてくれるのか。それは私の世界の最新の状態を保ちます。電子的に接続されているもの (メール、カレンダー、メッセージ、フィード、フォローしているアカウント) はすべて自動的に流れ込んでくるので、私が尋ねたときにはすでにプレイの状態がわかっています。これにより、時事問題を日常的なタスクに反映させることができます。たとえば、会議前の注意喚起、関心のある事柄の監視、毎日午後に食事をする管理者などです。そして、標準の MCP 接続を介して通常のクロード クライアントからすべてにアクセスします。これは、私が既に使用しているのと同じアシスタントであり、現在は私自身のメモリ、データ、ツールによってサポートされています。
pacslate.com/demo — 実際の UI。プライベート バックエンドを持たずに、加工されたサンプル データ上で実行されます。
キャンバス — プロンプトを発行します (または自動再生を確認します)。オペレーターは専門エージェントにルーティングし、合成された回答をストリーミングし、結果カードを生成します。結果カードには、サービスを提供したモデルと通話料金がタグ付けされています。 →キャンバスを開く
モニター — 仕事

LDダッシュボード。デモでは、USGS 地震情報と公開ヘッドライン フィード (HackerNews) がライブで実行され、クライアント側で取得されます。市場、航空機、個人レイヤーは、明確にラベル付けされたサンプル データです。 (運用環境では、ヘッドライン フィードは有線サービス RSS であり、マーケット/フライト フィードはライブです。) → モニターを開きます。
個人レイヤーはサンプル データであり、プライベート バックエンドはありません。実稼働システムは設計上プライベートです。
キャンバス (サンプル データ) — エージェントの回答は移動可能な結果カードとしてストリームされ、それぞれに対応したモデルと通話コストがタグ付けされます。ポインタは 1 回限りのガイド付きプレビューです。カードをドラッグしてワークスペースのテーマを変更し、マウスを動かしたり入力したりすると制御が戻ります。
モニター — USGS 地震情報および公開ヘッドライン フィードはライブで実行されます (クライアント側)。市場、航空機、個人レイヤーはサンプル データです。
私が使用していたフロンティア チャット ツールは、私が望むようにセッション全体のコンテキストを保持できず、出力レイヤーでフィルタリングされ、データが側に保持されていました。私はエンドツーエンドの自分のものを望んでいました。継続的に実行し、重要なことを覚えていて、自分の知識を根拠に推論し、質問されるのを待つのではなく役立つことを私に押し付けてくれます。
さらに難しいのは、雑多で散らばったインプット（私自身のメモや文書、私が関心のあるフィードやソース）を大量に収集し、それを単なる要約ではなく、検証済みで実用的なものに変えることです。
1 つの設計選択が残りを決定しました。モデルは私が所有していない部品なので、交換できる部品になるようにシステムを構築しました。耐久性のある部分、つまり私のデータ、メモリ、ツール、ルーティングとオーケストレーション ロジックは私のものであり、そのまま残ります。モデルはレンタルされており、標準インターフェイス経由でアクセスできるため、より優れたモデル (または別のプロバイダー) への移行は、再構築ではなく、ほとんどの場合、構成の変更になります。毎日私は博士

クロード コードを通じて実行しますが、デザインはそれに依存しません。
平易な言葉での作品ツアー:
マルチエージェントコア。オペレーターはリクエストを読み取り、調査、コーディング、分析、レビュー、その他いくつかの適切な専門家にルーティングします。それぞれの専門家は、その仕事に最適な (そして価格設定された) モデルに支えられており、ベンダー間で自動的にランク付けされたフォールバック機能を備えているため、1 つのモデルがダウンしても実行が停止することはありません。
ツールレイヤー。エージェントが (モデル コンテキスト プロトコル/MCP 経由で) 呼び出すことができるおよそ 24 個のツール サーバー: メモリ、ナレッジ検索、コード実行、Web 検索とフィード、インフラストラクチャなど。新しい機能は、エージェントを再接続するのではなく、サーバーを登録することによって追加されます。また、エージェントは、すべてのツールをコンテキスト内に保持するのではなく、インデックスを検索することによって適切なツールを見つけます。
持続する記憶。システムが推論するプライベートなインデックス付きナレッジ ベースに加えて、関連する会話と意思決定を時間の経過とともにリンクするレイヤーにより、すべてのチャットをリセットするのではなく、コンテキストがセッション間で引き継がれます。これは単なるベクトル検索ではありません。高速で決定論的な関連性パスにより、モデルが実行される前に関連するものが選択され、記憶は時間の経過とともに減衰して調整されるため、ストアは最新の状態に保たれます。
積極的なルーチン。スケジュールされたジョブは、多くのソースから取得し、合成して、短い概要を私に渡します。そのため、システムは、私が要求したときだけでなく、バ​​ックグラウンドで動作します。各項目は、最初に信号に対してフィルター処理され、ドロップされたノイズは表示されるのではなくログに記録されます。
キャンバスインターフェイス。リクエストは単一のスクロール スレッドではなく、ライブで移動可能な結果カードを生成します。結果は生成されるとストリームされ、各カードにはどのモデルが応答したか、および通話コストが表示されます。
ガードレールが組み込まれています。すべてのモデル呼び出しは予算に応じて計測され、すべての呼び出しが追跡されるため、何が実行されたのか、何が原因なのかを確認できます。

向きを変え、その費用はいくらでしたか。支出は各コールの前にチェックされ、上限に達すると失敗します。また、すべてのモデル スワップが追跡されます。
自己維持型。ルーチンはフリートを監視し、依存関係と CVE アップデートにフラグを立てて、リスクの低いものを適用し、一般的なルーチンの修正を処理します。つまり、私への判断が必要なものはすべてエスカレーションします。
これは機密システムではなく、そのふりをするものでもありません。日常的な AI ツールに関する現実的な問題は、パケットがインターネットを通過することではなく、ベンダーがユーザーの履歴を所有し、その履歴をトレーニングしており、その履歴を取り戻すことができないことです。 Pacific Slate はそれを逆転させます。記録システムはあなたが所有しており、推論はあなたが自分の条件でレンタルする商品です。
何があなたのもので、何があなたのものであり続けるのか。保存データ、ナレッジベース、継続性レイヤー、オーケストレーションなど、耐久性のあるほぼすべてのものは独自のサーバー上に存在します。エクスポート可能、削除可能で、トレーニングに使用されないように構成されています。意図的な例外の 1 つは、セッションをまたがるファクトに対するホストされたセマンティック メモリ サービスです。これは、記録システムではなく、交換可能な依存関係です (必要に応じて自己ホストすることもできます)。
レンタルしたもの。モデル呼び出しは、単一のゲートウェイを介してプロバイダーにルーティングされます。呼び出しは、プロンプトと必要なコンテキストを送信し、応答を取得します。応答は、入力をトレーニングしたり公開したりしないように設定されたプロバイダーにのみルーティングされます。
必要に応じて、外部推論への依存はありません。ルーティングはあなたのものです。あらゆるワークロードをボックス自体で実行されているモデルに固定できます。
これは意図的に実用的な姿勢であり、制御に真剣に取り組み、依然として最先端のモデルを使用しています。この主張は「何も残らない」というものではありません。それは、あなたが記録システムを所有し、物事の行き先を制御し、あなたのデータが他人の製品に入力されていないということです。
OpenRouter アカウント設定に示されているように (スクリーンショット、2026-06-

29): データのトレーニングの切り替え (有料と無料) がオフになり、プロンプト公開がオフになります。そのため、プロバイダーは入力をトレーニングしたり公開したりしません。ゼロデータ保持はグローバルにオフにされるため、完全なモデル名簿が利用可能な状態に保たれます。ZDR のみを強制することで、すべての非 ZDR エンドポイントがドロップされます。その一部は、システムが実行されるオープンウェイト モデルをホストします。また、呼び出しが必要なときにリクエストごとに設定できます。
これはユーザーが適応させるレシピであり、インストールする製品ではありません。重要なのは私の正確なビルドではありません。このアーキテクチャに独自のコーディング エージェントを指定し、モデル、ツール、ボックスなど、すでに持っているものを使用して、必要な部分を再構築させます。必要に応じて深くまたは浅く進みます。
ドライバーとして月額約 20 ドルのコーディングアシスタント、
常時稼働の基板向けの小規模 ~$20/月の VPS、
従量制モデルの通話がトップです。通常の個人負担での予算モデルの名簿では、月額 10 ～ 40 ドル程度です。
したがって、ソブリンで常時オンのマルチエージェント アシスタントは月額 50 ～ 80 ドル程度で、深さが必要な場合にのみスケールアップできます。セルフホスト モデルには大きなボックス、推論席にはプレミアム モデル、ビデオにはビデオ ツールなどです。構造は変わりません。その下限を実現するコスト規律 (複雑さスコア付きルーティング、安価なモデルのデフォルト、ハードバジェット上限、コンテキスト キャッシュ) は、§4 で説明したものと同じ機構です。 (私自身のビルドは毎日開発し、バックグラウンドで自動化を実行するため、動作が重くなります。そうしない場合でもアーキテクチャによって低コストが保たれます。)
リクエストの存続期間の短いバージョン:
リクエストが受信される → オペレータは、どのスペシャリストがそれを所有すべきかを決定する → そのエージェントはメモリとナレッジベースから必要なものを取得し、必要なツールを MCP 経由で呼び出す → モデルルーターが適切なモデルを選択する (モデルがダウンしている場合はフェイルオーバーする) → 回答が合成される

、キャンバスに戻され、保持する価値のあるものはすべてメモリに書き戻されます。料金は通話ごとに課金されます。パス全体がトレースされます。
フローチャート TD
U[ユーザー · Next.js キャンバス UI] -->|AG-UI over SSE| API[コグニティブAPI・スターレット]
API --> ORCH[オペレーター · リクエストをルーティング]
ORCH <-->|Redis イベント ストリーム · アウェアネス メッシュ| SPEC[専門エージェント<br/>コーダー · 研究者 · アナリスト · レビューアー · …]
SPEC --> ROUTER[復元モデルルーター<br/>役割 → モデル + フォールバックチェーン]
ROUTER --> LLM[(複数のモデルプロバイダー<br/>ロール価格設定 · 自己ホスト型モデルで終了)]
スペック -->|MCP| GW[MCPゲートウェイ]
GW --> TOOLS[~20 台の MCP ツールサーバー<br/>メモリ、知識、実行、git、インフラ、フィード、…]
SPEC --> RAG[(プライベートナレッジコーパス<br/>~27,000 ドキュメント / 294,000 チャンク)]
SPEC --> GRAPH[(連続グラフ)]
API --> COST[コストトラッカー・リクエストごとの予算ガードレール]
API --> OBS[可観測性 · リクエスト + モデルスワップ トレース]
サブグラフ インフラ[プライベートサーバー・Docker・Cloudflareトンネル]
API
オーク
スペック
ルーター
GW
ツール
ラグ
グラフ
終わり
技術的に興味のある人向け
以下のセクションでは、設計上の決定事項と故障モードについて説明します。特定のモデルとコンポーネントの数は、この記事の執筆時点での最新のものです。アーキテクチャは設計上モデルに依存しないため、構造を変更せずに個々のモデルが変更されます。
モデルは 1 つのレンタルされたパーツであるため、システムはそれを 2 つの標準インターフェイスの背後で交換可能として扱います。
MCP ゲートウェイ - 基板のツールとメモリはモデル コンテキスト プロトコル経由で公開されるため、MCP 対応のクライアントはそれらを使用できます。現在、そのクライアントは Claude Code です。これはフロントエンドとして実行され、基板は永続メモリ、ツール群、およびその下のルーティングを提供します。別の MCP ネイティブ CLI は同じサーフェスをマウントできます。
OpenAI 互換エンドポイント — パイプライン i

標準のチャット補完 API としても呼び出すことができるため、その形式を使用するクライアントは、内部構造を知らなくてもチャット補完 API を実行できます。
明確にしておきますが、クロードは私が毎日使用するエンジンであり、現在エンドツーエンドで接続されているエンジンです。しかし、MCP と OpenAI 互換のチャット形式はどちらも広くサポートされている非独自のインターフェイスであり、特定のベンダーに限定されるものではありません。そのため、「エンジンを交換し、基板は維持する」は単なる願望ではなく、設計の特性です。モデルが移動しても、私の作業 (メモリ、データ、ツール、ルーティング、検証) は移動しません。
キャンバス (Next.js シングルページ アプリ) はストリーミング接続を開き、プロンプトを送信します。ストリーミングでは、 Server-Sent Events 上で AG-UI プロトコルが使用され、バックエンドにプロキシされるため、トークン、ツール呼び出し、モデル解決イベントはすべて 1 つのストリームで到着します。
バックエンドは、キャンバスの AG-UI SSE エンドポイントと、他のクライアントが同じパイプラインと通信できるようにする OpenAI 互換エンドポイントの 2 つのサーフェスを公開する Starlette サービスです。
オペレータ エージェントはリクエストを分類し (直接回答するか、専門家に引き継ぐか)、それに応じてルーティングします。
選ばれたスペシャリストはコンテキスト (メモリ、知識ベース) を取得し、必要なツールを MCP 経由で呼び出します。
復元力のあるモデル層は、その呼び出しの実際のモデルを解決し (コストを意識したダウングレードと健全性を意識したフォールバックを使用します。§4 を参照)、完了をストリーミングし、コストを測定します。
答えは流れてくる

[切り捨てられた]

## Original Extract

A self-hosted, model-agnostic personal AI system — the architecture and the design decisions behind it.

">
Pacific Slate //
live demo
monitor
source
Copy page
▾
Copy as Markdown
Open Markdown
Chat in Claude
Open in ChatGPT
Pacific Slate: a personal AI system
Pacific Slate is a personal AI system that works over my own knowledge and the sources I follow, and does useful work in the background instead of only answering when asked.
It runs on a private remote server, remembers context across sessions, and is built from specialized agents I route work to. This page is a plain overview: what it is, how it’s put together, and how it handles privacy . Technical details are at the end if interested.
Up front: I’m an operator and systems architect, not a career software engineer. I designed the architecture, made the model-routing, reliability, and cost decisions, and built it AI-natively, in partnership with a range of AI platforms, agents, and coding tools.
What it does for me, day to day. It keeps itself current on my world. Whatever’s wired in electronically (mail, calendar, messages, the feeds and accounts I follow) flows in on its own, so it already knows the state of play when I ask. It puts current events to work on ordinary tasks: a heads-up before a meeting, a standing watch on something I care about, the daily admin that usually eats an afternoon. And I reach all of it from a normal Claude client over a standard MCP connection — the same assistant I already use, now backed by my own memory, data, and tools.
pacslate.com/demo — the real UI, running on fabricated sample data with no private backend.
Canvas — issue a prompt (or watch the autoplay): the operator routes to specialist agents, streams a synthesized answer, and spawns result cards, each tagged with the model that served it and what the call cost. → open the canvas
Monitor — a world dashboard. In the demo, USGS seismic and a public headline feed (HackerNews) run live, fetched client-side; markets, aircraft, and the personal layer are clearly-labeled sample data. (In production the headline feed is wire-service RSS and the market/flight feeds are live.) → open the monitor
The personal layer is sample data and there’s no private backend; the production system is private by design.
The canvas (sample data) — the agent’s answer streams in as movable result cards, each tagged with the model that served it and the call’s cost. The pointer is a one-time guided preview: it drags a card and re-themes the workspace, then hands control back the moment you move the mouse or type.
The monitor — USGS seismic and a public headline feed run live (client-side); markets, aircraft, and the personal layer are sample data.
The frontier chat tools I was using didn’t hold context across sessions the way I wanted, filtered at the output layer, and kept my data on their side. I wanted something that was mine end to end: runs continuously, remembers what matters, reasons over my own knowledge, and pushes useful things to me instead of waiting to be asked.
The harder part: take a lot of noisy, scattered input (my own notes and documents, the feeds and sources I care about) and turn it into something verified and actionable instead of just another summary.
One design choice shaped the rest: the model is the part I don’t own, so I built the system so it’s the part I can swap. The durable pieces — my data, the memory, the tools, the routing and orchestration logic — are mine and stay put; the model is rented, reached over standard interfaces, so moving to a better one (or another provider) is mostly a config change rather than a rebuild. Day to day I drive it through Claude Code, but the design doesn’t depend on that.
A plain-language tour of the pieces:
A multi-agent core. An operator reads a request and routes it to the right specialist — research, coding, analysis, review, and a few others — each backed by the model best suited (and priced) for its job, with automatic ranked fallback across vendors so one model going down doesn’t stop the run.
A tool layer. Roughly two dozen tool-servers the agents can call (over the Model Context Protocol / MCP): memory, knowledge retrieval, code execution, web search and feeds, infrastructure, and more. New capabilities are added by registering a server, not by rewiring the agents — and agents find the right tool by searching an index, rather than holding all of them in context.
Memory that persists. A private, indexed knowledge base the system reasons over, plus a layer that links related conversations and decisions over time, so context carries across sessions instead of resetting every chat. It’s more than vector lookup: a fast, deterministic relevance pass picks what’s relevant before any model runs, and memories decay and reconcile over time so the store stays current.
Proactive routines. Scheduled jobs that pull from many sources, synthesize, and hand me a short brief — so the system works in the background, not only when I ask — and each item is filtered for signal first, with the noise it dropped logged rather than shown.
A canvas interface. Requests spawn live, movable result cards rather than a single scrolling thread — results stream in as they’re generated, each card showing which model answered and what the call cost.
Guardrails built in. Every model call is metered against a budget, and every call is traced — so I can see what ran, what it returned, and what it cost. Spend is checked before each call — it fails closed at the cap — and every model swap is traced.
Self-maintaining. A routine watches the fleet, flags dependency and CVE updates and applies the low-risk ones, and handles common routine fixes — escalating anything that needs a judgment call to me.
This isn’t a classified system and doesn’t pretend to be one. The realistic problem with everyday AI tools isn’t that packets cross the internet — it’s that the vendor owns your history, trains on it, and you can’t get it back. Pacific Slate inverts that: you own the system of record, and inference is a commodity you rent on your own terms.
What’s yours, and what stays yours. Almost everything durable — data at rest, the knowledge base, the continuity layer, the orchestration — lives on your own server: exportable, deletable, and configured not to be used for training. The one deliberate exception is a hosted semantic-memory service for cross-session facts — a swappable dependency, not the system of record (and it can be self-hosted if you’d rather).
What’s rented. Model calls route to providers through a single gateway. A call sends the prompt and the context it needs, gets an answer — routed only to providers set not to train on or publish your inputs.
No external inference dependency, when you want it. The routing is yours: any workload can be pinned to a model running on the box itself.
It’s a pragmatic posture on purpose: serious about control, still using frontier models. The claim isn’t “nothing ever leaves.” It’s that you own the system of record, you control where things go, and your data isn’t being fed into someone else’s product.
As shown in the OpenRouter account settings ( screenshot , 2026-06-29): the train-on-data toggles (paid and free) are off and prompt-publishing is off — so providers don’t train on or publish your inputs. Zero-data-retention is left off globally so the full model roster stays available — enforcing ZDR-only drops every non-ZDR endpoint, some of which host the open-weight models the system runs on — and it can be set per request when a call needs it.
This is a recipe you adapt, not a product you install. The point isn’t my exact build. You point your own coding agent at this architecture and have it rebuild the parts you want, with whatever you already have: your models, your tools, your box. Go as deep or as shallow as you need.
a ~$20/mo coding assistant as the driver,
a small ~$20/mo VPS for the always-on substrate,
metered model calls on top — on a budget-model roster at a normal personal load, on the order of $10–40/mo.
So on the order of $50–80/mo for a sovereign, always-on, multi-agent assistant, and you scale up only where you want depth — a bigger box to self-host models, premium models for the reasoning seats, a video tool for video. The structure doesn’t change. The cost discipline that makes that floor real — complexity-scored routing, cheap-model defaults, a hard budget ceiling, context caching — is the same machinery described in §4. (My own build runs heavier, because I develop it daily and run background automations; the architecture is what keeps it cheap when you don’t.)
The short version of a request’s life:
A request comes in → the operator decides which specialist should own it → that agent pulls what it needs from memory and the knowledge base and calls any tools it needs over MCP → the model router picks the right model (and fails over if one is down) → the answer is synthesized, returned to the canvas, and anything worth keeping is written back to memory. Cost is metered on every call; the whole path is traced.
flowchart TD
U[User · Next.js canvas UI] -->|AG-UI over SSE| API[Cognitive API · Starlette]
API --> ORCH[Operator · routes the request]
ORCH <-->|Redis event stream · awareness mesh| SPEC[Specialist agents<br/>coder · researcher · analyst · reviewer · …]
SPEC --> ROUTER[Resilient model router<br/>role → model + fallback chain]
ROUTER --> LLMs[(Multiple model providers<br/>role-priced · ends on a self-hosted model)]
SPEC -->|MCP| GW[MCP gateway]
GW --> TOOLS[~two dozen MCP tool-servers<br/>memory · knowledge · exec · git · infra · feeds · …]
SPEC --> RAG[(Private knowledge corpus<br/>~27k docs / 294k chunks)]
SPEC --> GRAPH[(Continuity graph)]
API --> COST[Cost tracker · per-request budget guardrail]
API --> OBS[Observability · request + model-swap tracing]
subgraph Infra[Private server · Docker · Cloudflare Tunnel]
API
ORCH
SPEC
ROUTER
GW
TOOLS
RAG
GRAPH
end
For the technically curious
The sections below go down to the design decisions and the failure modes. Specific model and component counts are current as of this writing; the architecture is model-agnostic by design, so individual models change without the structure changing.
The model is the one rented part, so the system treats it as swappable, behind two standard interfaces:
An MCP gateway — the substrate’s tools and memory are exposed over the Model Context Protocol, so any MCP-capable client can use them. Today that client is Claude Code: it runs as the front end while the substrate supplies persistent memory, the tool fleet, and routing underneath. Another MCP-native CLI could mount the same surface.
An OpenAI-compatible endpoint — the pipeline is also callable as a standard chat-completions API, so a client that speaks that format can drive it without knowing the internals.
To be clear: Claude is the engine I use day to day, and the one wired end to end right now. But MCP and the OpenAI-compatible chat format are both widely-supported, non-proprietary interfaces — not one vendor’s lock-in — so “swap the engine, keep the substrate” is a property of the design, not just an aspiration: the work that’s mine (memory, data, tools, routing, verification) doesn’t move when the model does.
The canvas (a Next.js single-page app) opens a streaming connection and posts the prompt. Streaming uses the AG-UI protocol over Server-Sent Events , proxied to the backend so tokens, tool calls, and model-resolution events all arrive on one stream.
The backend is a Starlette service exposing two surfaces: the AG-UI SSE endpoint for the canvas, and an OpenAI-compatible endpoint so other clients can talk to the same pipeline.
The operator agent classifies the request — direct answer, or hand off to a specialist — and routes accordingly.
The chosen specialist pulls context (memory, knowledge base) and calls whatever tools it needs over MCP.
The resilient model layer resolves the actual model for that call (with cost-aware downgrades and health-aware fallback — see §4), streams the completion, and meters the cost.
The answer streams bac

[truncated]
