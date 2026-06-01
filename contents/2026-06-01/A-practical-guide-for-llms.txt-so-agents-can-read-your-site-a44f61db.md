---
source: "https://docsalot.dev/blog/what-is-llms-txt"
hn_url: "https://news.ycombinator.com/item?id=48347529"
title: "A practical guide for llms.txt so agents can read your site"
article_title: "What Is llms.txt? A Practical Guide for SaaS Docs Teams"
author: "fazkan"
captured_at: "2026-06-01T00:27:17Z"
capture_tool: "hn-digest"
hn_id: 48347529
score: 2
comments: 0
posted_at: "2026-05-31T17:21:35Z"
tags:
  - hacker-news
  - translated
---

# A practical guide for llms.txt so agents can read your site

- HN: [48347529](https://news.ycombinator.com/item?id=48347529)
- Source: [docsalot.dev](https://docsalot.dev/blog/what-is-llms-txt)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T17:21:35Z

## Translation

タイトル: エージェントがサイトを読めるようにするための llms.txt の実践的なガイド
記事のタイトル: llms.txt とは何ですか? SaaS ドキュメント チームのための実践ガイド
説明: llms.txt は、機械可読なドキュメントのインデックスです。これが何をするのか、何を含めるべきか、そしてそれ自体では解決できないものを次に示します。

記事本文:
llms.txtとは何ですか? SaaS Docs Teams の実践ガイド DocsAlot Docs ベンチマーク機能 ▾ すべての機能ページ DocsAlot の背後にあるドキュメント ワークフローを参照します。 AI 可読ドキュメント AI 可読ドキュメント GitHub ドキュメント自動化 GitHub ドキュメント自動化 API ドキュメント自動化 API ドキュメント自動化 SaaS 用 CLI SaaS 用 MCP SaaS 用ホスト型 MCP サーバー SaaS 用 SDK SaaS 用 SDK SaaS 用 llms.txt ドキュメント用 llms.txt ドキュメント用 ヘルプセンター + 開発者用ドキュメント ヘルプセンター + 開発者用ドキュメント チーム用▾ すべての対象者ページ 創設者、API、サポート、および開発者ツールの位置付けページを参照します。創設者 創設者向けドキュメント 開発者ツール 開発者ツール会社向けドキュメント API 会社 API 会社向けドキュメント SaaS + エージェント エージェント採用のための構築を行う SaaS 会社向けドキュメント サポート チーム サポート チーム向けドキュメント 価格 変更ログ ブログ サインイン 無料で開始 → メニュー ヘルプ センター ベンチマーク コンポーネント API ドキュメント 機能 ▾ すべての機能ページ AI 読み取り可能なドキュメント GitHub ドキュメントの自動化 API ドキュメントの自動化 SaaS 用の CLI SaaS 用の MCP SaaS 用の SDK ドキュメント用の llms.txt ヘルプ センター + 開発ドキュメント チーム向け ▾ すべての対象者ページ 創設者 開発者ツール API 企業 SaaS + エージェント サポート チーム 価格設定 ブログ 変更履歴 サインイン 無料で始める → ← ブログ / エンジニアリングに戻る llms.txt とは何ですか? SaaS ドキュメント チームのための実践ガイド
llms.txt は、機械可読なドキュメントのインデックスです。これが何をするのか、何を含めるべきか、そしてそれ自体では解決できないものを次に示します。
SaaS 製品のドキュメント サイトを運営している場合は、人間と検索エンジン向けの構造がすでに公開されています。人々のためのナビゲーションがあります。おそらく Google の sitemap.xml があるでしょう。 llms.txt は AI エージェントに欠けているレイヤーです。AI エージェントにどこにあるかを伝える小さなファイルです。

役立つドキュメントが存在し、どこから始めればよいか。
複雑ではありません。これは新しいプロトコル スタックではありません。これは、エージェントがドキュメントを見つけやすくするための実用的な方法にすぎません。
この投稿では、llms.txt とは何か、そこに何を入れるか、SaaS ドキュメント チームがそれについてどのように考えるべきか、そしてどこで十分ではなくなるのかについて説明します。
llmstxt.org で公開されている仕様は短いので、読む価値があります。以下は、実際の製品ドキュメントを管理するチーム向けの実用的なバージョンです。
llms.txt はプレーンテキストまたはマークダウン ファイルで、通常はドキュメント サイトのルートで提供され、AI エージェントにドキュメントの厳選されたインデックスを提供します。
手書きの地図と考えてください。
クローラーダンプではありません。サイト上のすべての URL ではありません。エージェントが最初に実際に読む必要があるページだけです。
最小バージョンは次のようになります。
コピー 1 # Acme ドキュメント 2
3 認証、API リクエスト、Webhook、4 つのレート制限、SDK セットアップに関する Acme の開発者ドキュメント。 5
6 ## ここから始めてください 7
8 - [クイックスタート]( https://docs.acme.com/quickstart ) : 最初の API リクエストを設定する 9 - [認証]( https://docs.acme.com/authentication ) : API キー、OAuth、およびスコープ 10 - [エラー]( https://docs.acme.com/errors ) : エラー コードと再試行ガイダンス 11
12 ## 一般的なタスク 13
14 - [ Webhook ]( https://docs.acme.com/webhooks ) : イベント通知を受信する 15 - [ レート制限 ]( https://docs.acme.com/rate-limits ) : 制限とバックオフを理解する 16 - [ Node SDK ]( https://docs.acme.com/sdks/node ) : Node.js と統合する
それが核となる考え方です。モデルに推測させるのではなく、小さくて独自のエントリ ポイントのセットを与えます。
SaaS ドキュメント チームが注意すべき理由
ドキュメントへの最初のインターフェイスとして AI ツールを使用する人が増えています。
URL をクロード コードに貼り付けます。彼らは Cursor に API で認証する方法を尋ねます。彼らは、製品の Webhook を設定する方法を ChatGPT に尋ねます。

これらすべてのケースにおいて、モデルは有用なドキュメントがどこにあるのか、どのページが重要なのかを把握する必要があります。
ドキュメントが適切に構造化されていれば、通常はモデルを難なく進めることができます。ただし、llms.txt にはショートカットが用意されています。
一般的なタスクをどこから始めるべきか
どの参照が正規であるか
これは、ほとんどの製品に同じ障害モード (ページ数が多すぎる、重複が多すぎる、マシンの明確な「ここから開始」パスがない) があるため、特に SaaS ドキュメントに役立ちます。
人間は閲覧することでそこから回復することができます。エージェントは、直接指示したほうがはるかに優れています。
最も簡単な間違いは、llms.txt をドキュメントの完全なエクスポートのように扱うことです。
優れた llms.txt は短く、厳選されており、エージェントが実際のタスクを完了するのに役立つページに偏っています。
ほとんどの SaaS ドキュメント チームには、次のものが含まれます。
製品またはドキュメントセットの一文の説明
クイックスタートまたは最初に成功した統合パス
認証または認可のドキュメント
エラー、レート制限、Webhook などの一般的な運用トピック
SDK ページ（実際の導入パスの場合）
製品のアーキテクチャが明確ではない場合、中核となる概念については 1 ～ 2 ページ
複数の視聴者をサポートする場合は、リンクをグループ化します。
コピー 1 ## API の基本 2 - [クイックスタート]( ... ) 3 - [認証]( ... ) 4 - [ レート制限 ]( ... ) 5
6 ## 統合 7 - [ Webhook ]( ... ) 8 - [ Zapier ]( ... ) 9 - [ Slack 統合 ]( ... ) 10
11 ## SDK 12 - [ノード SDK ]( ... ) 13 - [ Python SDK ]( ... )
これは、1 つの巨大なフラット リストよりもモデルにとって使いやすいです。
2 番目の間違いは、llms.txt にドキュメントに隣接するものをすべて詰め込んでしまうことです。
サイトマップの代わりとして使用しないでください。
現在の API の動作にとって重要でない限り、changelog ノイズ
1 つのバージョンが正規の場合、バージョン管理されたページが重複する
壊れた、古い、または薄いページ
内部ダッシュボードまたは認証の背後にあるページ
n を含む巨大なリンク リスト

説明
新しい開発者を初日から送り込むようなページではない場合、そのページはおそらく llms.txt の先頭近くに属していません。
ファイルは曖昧さを増やすのではなく、曖昧さを減らす必要があります。
llms.txt をドキュメント サイトに追加する方法
ほとんどのチームにとって、実装は良い意味で退屈です。
ドキュメントが docs.example.com にある場合は、そこで公開してください。メインのドキュメントが example.com/docs にある場合は、ユーザーが実際に使用するのと同じホストに公開してください。
2. 先頭の 10 ～ 20 ページから始めます
最初に完全性を解決しようとしないでください。まず有用性を解決します。
エンジニア向けの重要なオンボーディング メモのようにファイルを作成します。
これは、エージェントとツールが期待し始めている規則です。
Bash コピー 1 カール https://docs.example.com/llms.txt
ファイルがプレーン テキストとして読みにくい場合は、モデルにとっても読みにくい可能性があります。
クイックスタートが変更された場合、認証モデルが変更された場合、または新しい SDK がデフォルトのパスになった場合は、llms.txt も更新してください。 1 回限りの SEO 成果物ではなく、小さな製品インフラストラクチャのように扱ってください。
これらのファイルはさまざまな役割を果たします。
sitemap.xml はクローラー指向のインベントリーです。どの URL が存在し、いつ変更されたかを検索エンジンに伝えます。
llms.txt は、厳選されたガイダンス レイヤーです。 AI エージェントに、どの URL が読む価値があるか、ドキュメントがどのように高レベルで構成されているかを伝えます。
サイトマップは完全性を追求して最適化されています。優れた llms.txt は、使いやすさを考慮して最適化されています。
すでに優れたサイトマップを持っている場合、それはある意味で見つけやすさに役立ちます。しかし、それでもモデルに次のことは伝えられません。
実際のクイックスタートはどのページですか
OAuth または API キーが推奨される認証パスであるかどうか
どの参照ページが最も重要か
それが llms.txt によって埋められるギャップです。
これはほとんどのチームが見逃している部分です。
llms.txt は発見に役立ちます。自動的にドキュメントが安価になり、解釈しやすくなるわけではありません。

、より深い意味ではエージェント対応です。
また、エージェントに製品の上手な使い方を伝えることもできません。そのためには、 skill.md ファイル 、より優れたマークダウン配信、または全体的により構造化されたドキュメント エクスペリエンスが必要になる場合もあります。
そのため、llms.txt をソリューション全体ではなく、最初のレイヤーとして扱います。システム ビューが必要な場合は、それを llms.txt で十分に説明していないので説明します。
簡単な開始点が必要な場合は、これを使用します。
コピー 1 # 製品ドキュメント 2
3 製品の技術文書。セットアップ、認証、4 つの API の使用法、SDK、Webhook、レート制限、トラブルシューティングについて説明します。 5
6 ## ここから始めてください 7
8 - [クイックスタート]( https://docs.yourproduct.com/quickstart ) : 最初のリクエストを行う 9 - [認証]( https://docs.yourproduct.com/authentication ) : API キーと OAuth 10 - [ API の概要 ]( https://docs.yourproduct.com/api ) : コア エンドポイントと概念 11
12 ## 一般的なタスク 13
14 - [ Webhook ]( https://docs.yourproduct.com/webhooks ) : イベントの受信と検証 15 - [ エラー ]( https://docs.yourproduct.com/errors ) : 一般的なエラーのデバッグ 16 - [ レート制限 ]( https://docs.yourproduct.com/rate-limits ) : 再試行とバックオフの処理 17
18 ## SDK 19
20 - [ノード SDK](https://docs.yourproduct.com/sdks/node) 21 - [Python SDK](https://docs.yourproduct.com/sdks/python)
有用な最初のバージョンを出荷するにはこれで十分です。
人々やツールがどのように使用されているかを確認したら、後でいつでも拡張できます。
それについての正しい考え方
llms.txt を SEO ハックとは考えないでください。
これは、エージェントのドキュメント アフォーダンスと考えてください。
人間はサイドバー、パンくずリスト、厳選されたクイックスタートから恩恵を受けます。エージェントは、より小さく、より明確な形式で表現された同じ種類の編集上の判断から恩恵を受けることができます。
どのページが最も重要かをドキュメント チームがすでに知っている場合、llms.txt は単に

それを機械可読な方法で書き留める場所です。
そして、そこで止まったとしても、それは依然として意味のある改善です。
その後に次のレイヤーが必要な場合は、マークダウン配信、よりクリーンなドキュメント構造、および製品の上手な使い方をエージェントに伝える skill.md を追加します。
llms.txt の例: API ドキュメント、ヘルプセンター、開発者ドキュメントの実際のパターン
ほとんどの llms.txt ファイルは、あいまいすぎるか、肥大化しすぎます。 API ドキュメント、ヘルプ センター、開発者ドキュメントに実際に適合する 3 つのパターンを次に示します。
SaaS 用のリモート MCP サーバーをセットアップする方法
MCP は理論的には簡単ですが、実際には面倒です。このガイドでは、MCP とは何か、SaaS 製品にとってリモート MCP が重要な理由、主な展開オプション、ユーザーにローカル設定を強制せずにトークンレス オンボーディングを行う方法について説明します。
誰にもお金を払わずにPHでオーガニック投票をどれくらい獲得できますか?ケーススタディ
私はハンター、インフルエンサー、立ち上げコンサルタントにお金を払わずに、Product Hunt で狭い製品を立ち上げました。最終的には 51 票を獲得して 19 位に終わりましたが、有益だったのは、オーガニック ローンチが実際に何に依存しているのかを学んだことです。
DocsAlot 蓄積するのではなく、複合的なドキュメントを求める創業者のためのドキュメント インフラストラクチャ。

## Original Extract

llms.txt is a machine-readable index for documentation. Here's what it does, what to include, and what it won't solve on its own.

What Is llms.txt? A Practical Guide for SaaS Docs Teams DocsAlot Docs Benchmark Features ▾ All feature pages Browse the documentation workflows behind DocsAlot. AI-readable docs AI-readable documentation GitHub docs automation GitHub docs automation API docs automation API docs automation CLIs for your SaaS CLIs for your SaaS MCP for your SaaS Hosted MCP servers for your SaaS SDKs for your SaaS SDKs for your SaaS llms.txt for docs llms.txt for docs Help center + dev docs Help center + developer docs For teams ▾ All audience pages Browse founder, API, support, and developer-tools positioning pages. Founders Documentation for founders Developer tools Documentation for developer-tools companies API companies Documentation for API companies SaaS + agents Documentation for SaaS companies building for agent adoption Support teams Documentation for support teams Pricing Changelog Blog Sign in Start for free → Menu Help Center Benchmark Components API Docs Features ▾ All feature pages AI-readable docs GitHub docs automation API docs automation CLIs for your SaaS MCP for your SaaS SDKs for your SaaS llms.txt for docs Help center + dev docs For teams ▾ All audience pages Founders Developer tools API companies SaaS + agents Support teams Pricing Blog Changelog Sign in Start for free → ← Back to blog / Engineering What Is llms.txt? A Practical Guide for SaaS Docs Teams
llms.txt is a machine-readable index for documentation. Here's what it does, what to include, and what it won't solve on its own.
If you run a docs site for a SaaS product, you already publish structure for humans and search engines. You have navigation for people. You probably have a sitemap.xml for Google. llms.txt is the missing layer for AI agents: a small file that tells them where your useful documentation lives and where to start.
It is not complicated. It is not a new protocol stack. It is just a practical way to make your docs easier for agents to discover.
This post covers what llms.txt is, what to put in it, how SaaS docs teams should think about it, and where it stops being enough.
The public spec at llmstxt.org is short and worth reading. What follows is the practical version for teams maintaining real product docs.
llms.txt is a plain-text or markdown file, usually served at the root of a docs site, that gives AI agents a curated index of your documentation.
Think of it as a hand-written map.
Not a crawler dump. Not every URL on your site. Just the pages an agent should actually read first.
A minimal version looks like this:
Copy 1 # Acme Docs 2
3 Acme's developer documentation for authentication, API requests, webhooks, 4 rate limits, and SDK setup. 5
6 ## Start Here 7
8 - [ Quickstart ]( https://docs.acme.com/quickstart ) : Set up your first API request 9 - [ Authentication ]( https://docs.acme.com/authentication ) : API keys, OAuth, and scopes 10 - [ Errors ]( https://docs.acme.com/errors ) : Error codes and retry guidance 11
12 ## Common Tasks 13
14 - [ Webhooks ]( https://docs.acme.com/webhooks ) : Receive event notifications 15 - [ Rate Limits ]( https://docs.acme.com/rate-limits ) : Understand limits and backoff 16 - [ Node SDK ]( https://docs.acme.com/sdks/node ) : Integrate with Node.js
That is the core idea. Give the model a small, opinionated set of entry points instead of making it guess.
Why SaaS Docs Teams Should Care
More people are using AI tools as the first interface to your docs.
They paste a URL into Claude Code. They ask Cursor how to authenticate with your API. They ask ChatGPT how to set up webhooks for your product. In all of those cases, the model has to figure out where the useful documentation is and which pages matter.
If your docs are well structured, a model can usually muddle through. But llms.txt gives it a shortcut:
where to start for common tasks
which references are canonical
That is especially useful for SaaS docs because most products have the same failure mode: too many pages, too much duplication, and no obvious "start here" path for a machine.
Humans can recover from that by browsing. Agents are much better if you point them directly.
The easiest mistake is treating llms.txt like a full export of your docs.
A good llms.txt is short, curated, and biased toward pages that help an agent complete real tasks.
For most SaaS docs teams, include:
a one-sentence description of the product or docs set
the quickstart or first successful integration path
authentication or authorization docs
common operational topics like errors, rate limits, and webhooks
SDK pages if they are real adoption paths
one or two pages for core concepts if your product has non-obvious architecture
If you support multiple audiences, group the links:
Copy 1 ## API Basics 2 - [ Quickstart ]( ... ) 3 - [ Authentication ]( ... ) 4 - [ Rate Limits ]( ... ) 5
6 ## Integrations 7 - [ Webhooks ]( ... ) 8 - [ Zapier ]( ... ) 9 - [ Slack integration ]( ... ) 10
11 ## SDKs 12 - [ Node SDK ]( ... ) 13 - [ Python SDK ]( ... )
This is easier for a model to use than one giant flat list.
The second mistake is stuffing llms.txt with everything adjacent to documentation.
Do not use it as a sitemap replacement.
changelog noise unless it is critical for current API behavior
duplicate versioned pages if one version is canonical
broken, outdated, or thin pages
internal dashboards or pages behind auth
giant link lists with no explanation
If a page is not something you would send a new developer to on day one, it probably does not belong near the top of llms.txt .
The file should reduce ambiguity, not create more of it.
How to Add llms.txt to Your Docs Site
For most teams, the implementation is boring in a good way.
If your docs live on docs.example.com , publish it there. If your main documentation lives on example.com/docs , publish it on the same host users actually rely on.
2. Start with the top 10 to 20 pages
Do not try to solve completeness first. Solve usefulness first.
Write the file like a high-signal onboarding note for an engineer:
That is the convention agents and tooling are starting to expect.
Bash Copy 1 curl https://docs.example.com/llms.txt
If the file is hard to read as plain text, it is probably hard for a model too.
If your quickstart changes, your auth model changes, or a new SDK becomes the default path, update llms.txt too. Treat it like a small piece of product infrastructure, not a one-time SEO artifact.
These files do different jobs.
sitemap.xml is a crawler-oriented inventory. It tells search engines which URLs exist and when they changed.
llms.txt is a curated guidance layer. It tells AI agents which URLs are worth reading and how the docs are organized at a high level.
A sitemap is optimized for completeness. A good llms.txt is optimized for usefulness.
If you already have a great sitemap, that helps with discoverability in one sense. But it still does not tell a model:
which page is the real quickstart
whether OAuth or API keys are the recommended auth path
which reference pages matter most
That is the gap llms.txt fills.
This is the part most teams miss.
llms.txt helps with discovery. It does not automatically make your docs cheap to consume, easy to interpret, or agent-ready in the deeper sense.
And it does not tell an agent how to use your product well. For that, you may also want a skill.md file , better markdown delivery, or a more structured docs experience overall.
That is why we treat llms.txt as the first layer, not the whole solution. If you want the systems view, we already wrote that up in llms.txt Isn't Enough .
If you want a simple starting point, use this:
Copy 1 # Your Product Docs 2
3 Technical documentation for Your Product. Covers setup, authentication, 4 API usage, SDKs, webhooks, rate limits, and troubleshooting. 5
6 ## Start Here 7
8 - [ Quickstart ]( https://docs.yourproduct.com/quickstart ) : Make your first request 9 - [ Authentication ]( https://docs.yourproduct.com/authentication ) : API keys and OAuth 10 - [ API Overview ]( https://docs.yourproduct.com/api ) : Core endpoints and concepts 11
12 ## Common Tasks 13
14 - [ Webhooks ]( https://docs.yourproduct.com/webhooks ) : Receive and verify events 15 - [ Errors ]( https://docs.yourproduct.com/errors ) : Debug common failures 16 - [ Rate Limits ]( https://docs.yourproduct.com/rate-limits ) : Handle retries and backoff 17
18 ## SDKs 19
20 - [ Node SDK ]( https://docs.yourproduct.com/sdks/node ) 21 - [ Python SDK ]( https://docs.yourproduct.com/sdks/python )
That is enough to ship a useful first version.
You can always expand it later once you see how people and tools use it.
The Right Way to Think About It
Do not think of llms.txt as an SEO hack.
Think of it as a documentation affordance for agents.
Humans benefit from sidebars, breadcrumbs, and curated quickstarts. Agents benefit from the same kind of editorial judgment, just expressed in a smaller and more explicit format.
If your docs team already knows which pages matter most, llms.txt is simply the place to write that down in a machine-readable way.
And if you stop there, that is still a meaningful improvement.
If you want the next layer after that, add markdown delivery, cleaner docs structure, and a skill.md that tells agents how to use your product well.
llms.txt Examples: Real Patterns for API Docs, Help Centers, and Developer Docs
Most llms.txt files are either too vague or too bloated. Here are three patterns that actually fit API docs, help centers, and developer documentation.
How to Set Up a Remote MCP Server for Your SaaS
MCP is straightforward in theory and messy in practice. This guide covers what MCP is, why remote MCP matters for SaaS products, the main deployment options, and how to get to tokenless onboarding without forcing users through local setup.
How Many Organic Votes Can You Get on PH Without Paying Anyone? A Case Study
I launched a narrow product on Product Hunt without paying a hunter, influencers, or a launch consultant. It ended with 51 votes and a 19th-place finish, and the useful part was learning what organic launches actually depend on.
DocsAlot Documentation infrastructure for founders who want docs that compound, not accumulate.
