---
source: "https://docsalot.dev/blog/llms-txt-for-saas-docs-what-to-include-and-what-to-skip"
hn_url: "https://news.ycombinator.com/item?id=48391619"
title: "Llms.txt is not a docs url dump"
article_title: "llms.txt for SaaS Docs: What to Include and What to Skip"
author: "fazkan"
captured_at: "2026-06-04T01:06:20Z"
capture_tool: "hn-digest"
hn_id: 48391619
score: 2
comments: 3
posted_at: "2026-06-03T23:38:15Z"
tags:
  - hacker-news
  - translated
---

# Llms.txt is not a docs url dump

- HN: [48391619](https://news.ycombinator.com/item?id=48391619)
- Source: [docsalot.dev](https://docsalot.dev/blog/llms-txt-for-saas-docs-what-to-include-and-what-to-skip)
- Score: 2
- Comments: 3
- Posted: 2026-06-03T23:38:15Z

## Translation

タイトル: Llms.txt はドキュメント URL ダンプではありません
記事のタイトル: SaaS ドキュメントの llms.txt: 含めるものとスキップするもの
説明: llms.txt の難しい部分は形式ではありません。ファイルに何を含めるべきかを決定するのです。

記事本文:
SaaS ドキュメントの llms.txt: 含めるものとスキップするもの DocsAlot Docs ベンチマーク機能 ▾ すべての機能ページ DocsAlot の背後にあるドキュメント ワークフローを参照します。 AI 可読ドキュメント AI 可読ドキュメント GitHub ドキュメント自動化 GitHub ドキュメント自動化 API ドキュメント自動化 API ドキュメント自動化 SaaS 用 CLI SaaS 用 MCP SaaS 用ホスト型 MCP サーバー SaaS 用 SDK SaaS 用 SDK SaaS 用 llms.txt ドキュメント用 llms.txt ドキュメント用 ヘルプセンター + 開発者用ドキュメント ヘルプセンター + 開発者用ドキュメント チーム用▾ すべての対象者ページ 創設者、API、サポート、および開発者ツールの位置付けページを参照します。創設者 創設者向けドキュメント 開発者ツール 開発者ツール会社向けドキュメント API 会社 API 会社向けドキュメント SaaS + エージェント エージェント採用のための構築を行う SaaS 会社向けドキュメント サポート チーム サポート チーム向けドキュメント 価格 変更ログ ブログ サインイン 無料で開始 → メニュー ヘルプ センター ベンチマーク コンポーネント API ドキュメント 機能 ▾ すべての機能ページ AI 読み取り可能なドキュメント GitHub ドキュメントの自動化 API ドキュメントの自動化 SaaS 用の CLI SaaS 用の MCP SaaS 用 SDK ドキュメント用 llms.txt ヘルプ センター + 開発ドキュメント チーム向け ▾ すべての対象者ページ 創設者 開発者ツール API 企業 SaaS + エージェント サポート チーム 価格設定 ブログ 変更ログ サインイン 無料で始める → ← ブログに戻る / エンジニアリング llms.txt for SaaS ドキュメント: 含めるものとスキップするもの
llms.txt の難しい部分は形式ではありません。ファイルに何を含めるべきかを決定するのです。
ほとんどの不良な llms.txt ファイルは、退屈な理由で失敗します。つまり、完全にしようとしすぎているということです。リンクが多すぎたり、隣接するコンテンツが多すぎたり、判断力が不十分であったりします。
それは通常、合理的な本能から来ています。重要なことを忘れたくない人はいません。ただし、llms.txt は 1 つです

完全性が目標ではない場合。
このファイルの役割は、エージェントの検索スペースを狭めることです。
したがって、本当の問題は「何を含めることができるか?」ということではありません。それは、「エージェントが実際にこの製品を使おうとする場合、最初に何を読むべきか？」です。
まだ基本事項を読んでいない場合は、「llms.txt とは何ですか?」から始めてください。 SaaS Docs チームのための実践ガイド。ファイルを機械的に構造化する方法を決定している場合は、「llms.txt をドキュメント サイトに追加する方法」を参照してください。
この投稿は、編集上のより難しい質問、つまりファイルに何が含まれ、何が除外されるべきかについてです。
ドキュメントが実行する必要があるジョブから開始する
ほとんどの SaaS 製品では、エージェントはドキュメントを使用して、次のいずれかのことを実行しようとしています。
適切なエンドポイントを呼び出すか、適切な SDK を使用してください
実装を妨げる概念を理解する
llms.txt はブランドの表面ではありません。これは docs ディレクトリではありません。ページ在庫ではありません。厳選されたタスクマップです。
ページがこれらのタスクのいずれかに実質的に役に立たない場合、そのページは最初のバージョンに属していない可能性があります。
ほとんどの SaaS ドキュメント チームにとって、デフォルトのキープリストは非常に安定しています。
クイックスタートと呼ばれるページではありません。実際に最初の統合を成功させるページ。
現在のクイックスタートが主にマーケティングのフレームワークやアーキテクチャの背景である場合は、それが最良の最初のステップであるかのように考えないでください。作業を完了するページを使用します。
これはほとんどの場合、含める価値があります。
製品が複数の認証モデルをサポートしている場合は、推奨されるパスを明確にしてください。
3. コア API または製品エントリ ポイント
API ドキュメントの場合、それはメインの API 概要またはリファレンス ランディング ページである可能性があります。
製品ドキュメントの場合は、セットアップの概要またはメインのワークフロー ページである可能性があります。
4. エラー、レート制限、再試行、またはトラブルシューティング
これらはエージェントにとって不釣り合いに役立ちます。なぜなら、これらはタスクの途中で表示されるためではありません。

ちょうど始めたばかりです。
5. 実際の導入パスを表す Webhook、SDK、または統合
考えられるすべての統合面ではなく、人々が実際に使用するものを含めます。
6. 製品に実際のアーキテクチャがある場合、1 つまたは 2 つのコンセプト ページ
実装を意味のあるものにする前に、ドキュメントで環境、テナント、同期モデル、または権限を理解する必要がある場合は、それらのページを明示的に含めます。
ほとんどのファイルでノイズが発生するのはここです。
購入者向けに書かれた製品概要ページ
これらのページは技術的には同じサイトの一部であっても、通常、エージェントがドキュメント タスクを解決するのには役立ちません。
変更ログは便利ですが、ほとんどの製品では、変更ログをコアの llms.txt マテリアルとして扱うべきではありません。
あるリリースで破壊的な認証変更が導入された場合、または新たに必要な統合パスが導入された場合は、他の場所で言及する価値があるかもしれません。ただし、長いリリース ノート アーカイブがファイルの大半を占めないようにしてください。
そこに新しいエンジニアを派遣しないなら、エージェントも派遣しないでください。
llms.txt は、ドキュメントの負債に関する内部の罪悪感を保存する場所ではありません。
同じガイドの複数のバージョンがある場合は、古いバージョンがまだアクティブで実質的に異なる場合を除き、正規のものを選択してください。
5. 内部ツール、ダッシュボード、または認証ゲートされたページ
当たり前のように聞こえますが、言う価値はあります。このファイルは、有用なパブリック ドキュメントの表面を反映している必要があります。
6. 構造を持たない巨大なリンクリスト
すべてのページが個別に有効であっても、リンクの壁は依然として弱いファイルです。
役立つテスト: これを新しいエンジニアに渡しますか?
これはおそらく、私が知っている llms.txt に関する最高の編集テストです。
もし有能なエンジニアが今日チームに加わったとしたら、これは良い出発点となるでしょうか?
答えが「はい」の場合は、おそらくその目標に近づいています。
答えが「いいえ」の場合、問題は通常、次の 3 つのうちのいずれかです。
ユーザータスクではなく内部組織図を反映するグループ化
あのて

st は単純ですが、ほとんどの間違いは見つかります。
さまざまな SaaS Docs Surface に通常必要なもの
すべてのドキュメント サイトに同じカテゴリを含める必要はありません。
重要なのは、1 つの普遍的な構造を強制しないことです。重要なのは、ドキュメント サーフェスの実際の役割に適合するページを選択することです。
重要だと思われるが、通常は役に立たないページ
公式だと思われるため、一部のページが含まれます。
ほとんど何も書かれていない一般的な「概要」ページ
実際のコンテンツのないカテゴリのランディング ページ
プレス発表形式のリリース投稿
「すべての統合」ディレクトリ
これらは必ずしも悪いページではありません。これらは通常、エージェントのガイダンスとしてあまり活用されないページです。
エージェントは、20 の製品にリンクするトップレベルの「統合」ページよりも、1 つの強力な Webhook ガイドの方がメリットが得られます。
優れたファイルには通常、視点がある
最良の llms.txt ファイルは中立的なものではありません。彼らは次のように提言しています。
これは正規の認証パスです
これらは一般的な故障モードです
この概念は、何かを実装する前に重要です
この観点がこのファイルを役立つものにしています。
それがなければ、まだドキュメントが残っています。あなたには指導力がないだけです。
自動生成された llms.txt ファイルが平凡に感じられる傾向があるのもこのためです。自動化はリンクを収集するのに優れています。どの少数のリンクが最も重要かを判断するのは苦手です。
実用的な「維持またはスキップ」チェックリスト
ページが llms.txt に属するかどうかを判断するときは、次のように質問してください。
このページは、誰かが実際の実装やサポートのタスクを完了するのに役立ちますか?
このページは正規のものであり、最新のものであり、信頼する価値がありますか?
エージェントに残りのドキュメントを参照する前にこれを読んでもらいたいでしょうか?
曖昧さが軽減されるのでしょうか、それともさらに曖昧さが増すのでしょうか？
答えがほとんどノーの場合は、省略してください。
それは失敗ではありません。それがキュレーションのポイントです。
SaaS ドキュメントの場合、適切な llms.txt には通常次の内容が含まれます。
のような運用ドキュメント

エラーとレート制限
実装の障害を取り除くコンセプトページ
巨大な未分化リンクダンプ
本当の価値は、そもそもそこに存在するに値するものを決定することから生まれます。
そのため、llms.txt は機械によって生成されたインベントリではなく、ドキュメント システム用の小さな編集成果物に近いものです。
llms.txt をドキュメント サイトに追加する方法
llms.txt の追加は機械的には簡単です。難しいのは、適切なページを選択し、ファイルを 2 番目のサイトマップに変えるのではなく、ファイルを有用な状態に保つことです。
llms.txt と sitemap.xml: AI 検出における各ファイルの役割
sitemap.xml は、何が存在するかをクローラーに伝えます。 llms.txt は AI エージェントに何が重要かを伝えます。 2026 年にドキュメントを実行する場合は、おそらく両方が必要になるでしょう。
llms.txt の例: API ドキュメント、ヘルプセンター、開発者ドキュメントの実際のパターン
ほとんどの llms.txt ファイルは、あいまいすぎるか、肥大化しすぎます。 API ドキュメント、ヘルプ センター、開発者ドキュメントに実際に適合する 3 つのパターンを次に示します。
DocsAlot 蓄積するのではなく、複合的なドキュメントを求める創業者のためのドキュメント インフラストラクチャ。

## Original Extract

The difficult part of llms.txt is not the format. It's deciding what deserves to be in the file at all.

llms.txt for SaaS Docs: What to Include and What to Skip DocsAlot Docs Benchmark Features ▾ All feature pages Browse the documentation workflows behind DocsAlot. AI-readable docs AI-readable documentation GitHub docs automation GitHub docs automation API docs automation API docs automation CLIs for your SaaS CLIs for your SaaS MCP for your SaaS Hosted MCP servers for your SaaS SDKs for your SaaS SDKs for your SaaS llms.txt for docs llms.txt for docs Help center + dev docs Help center + developer docs For teams ▾ All audience pages Browse founder, API, support, and developer-tools positioning pages. Founders Documentation for founders Developer tools Documentation for developer-tools companies API companies Documentation for API companies SaaS + agents Documentation for SaaS companies building for agent adoption Support teams Documentation for support teams Pricing Changelog Blog Sign in Start for free → Menu Help Center Benchmark Components API Docs Features ▾ All feature pages AI-readable docs GitHub docs automation API docs automation CLIs for your SaaS MCP for your SaaS SDKs for your SaaS llms.txt for docs Help center + dev docs For teams ▾ All audience pages Founders Developer tools API companies SaaS + agents Support teams Pricing Blog Changelog Sign in Start for free → ← Back to blog / Engineering llms.txt for SaaS Docs: What to Include and What to Skip
The difficult part of llms.txt is not the format. It's deciding what deserves to be in the file at all.
Most bad llms.txt files fail for a boring reason: they are trying too hard to be complete. They include too many links, too much adjacent content, and not enough judgment.
That usually comes from a reasonable instinct. Nobody wants to leave out something important. But llms.txt is one of those cases where completeness is not the goal.
The job of the file is to narrow the search space for an agent.
So the real question is not "what can we include?" It is "what should an agent read first if it is trying to actually use this product?"
If you have not read the basics yet, start with What Is llms.txt ? A Practical Guide for SaaS Docs Teams . If you are deciding how to structure the file mechanically, read How to Add llms.txt to Your Documentation Site .
This post is about the harder editorial question: what belongs in the file, and what should stay out?
Start From the Job the Docs Need to Do
For most SaaS products, an agent using your docs is trying to do one of a few things:
call the right endpoint or use the right SDK
understand a concept that blocks implementation
llms.txt is not a brand surface. It is not a docs directory. It is not a page inventory. It is a curated task map.
If a page does not materially help one of those tasks, it probably does not belong in the first version.
For most SaaS docs teams, the default keep-list is pretty stable.
Not the page you call the quickstart. The page that actually gets someone to a successful first integration.
If your current quickstart is mostly marketing framing or architectural background, do not pretend it is the best first step. Use the page that gets the work done.
This is almost always worth including.
If your product supports more than one auth model, make the recommended path obvious.
3. Core API or product entry point
For API docs, that might be the main API overview or a reference landing page.
For product docs, it might be the setup overview or the main workflow page.
4. Errors, rate limits, retries, or troubleshooting
These are disproportionately useful for agents because they show up mid-task, not just at the start.
5. Webhooks, SDKs, or integrations that represent real adoption paths
Include the things people actually use, not every possible integration surface.
6. One or two concept pages if the product has real architecture
If your docs require understanding environments, tenancy, sync models, or permissions before implementation makes sense, include those pages explicitly.
This is where most files get noisy.
product overview pages written for buyers
Even if those pages are technically part of the same site, they do not usually help an agent solve a docs task.
A changelog can be useful, but most products should not treat it as core llms.txt material.
If one release introduced a breaking auth change or a new required integration path, maybe that deserves a mention elsewhere. But a long release note archive should not dominate the file.
If you would not send a new engineer there, do not send an agent there either.
llms.txt is not the place to preserve internal guilt about docs debt.
If you have multiple versions of the same guide, choose the canonical one unless older versions are still active and materially different.
5. Internal tools, dashboards, or auth-gated pages
This sounds obvious, but it is worth saying. The file should reflect the useful public docs surface.
6. Giant link lists with no structure
Even if every page is individually valid, a wall of links is still a weak file.
The Useful Test: Would You Hand This to a New Engineer?
This is probably the best editorial test I know for llms.txt .
If a competent engineer joined the team today, would this be a good starting map?
If the answer is yes, you are probably close.
If the answer is no, the problem is usually one of three things:
groupings that reflect internal org charts instead of user tasks
That test is simple, but it catches most of the mistakes.
What Different SaaS Docs Surfaces Usually Need
Not every docs site should include the same categories.
The point is not to force one universal structure. The point is to choose the pages that fit the actual job of the docs surface.
Pages That Feel Important but Usually Do Not Help
Some pages get included because they feel official.
generic "overview" pages that say very little
category landing pages with no real content
press-announcement-style release posts
"all integrations" directories
These are not always bad pages. They are just usually not high-leverage pages for agent guidance.
An agent benefits more from one strong webhook guide than from a top-level "Integrations" page that links to twenty products.
A Good File Usually Has a Point of View
The best llms.txt files are not neutral. They express a recommendation:
this is the canonical auth path
these are the common failure modes
this concept matters before you implement anything else
That point of view is what makes the file helpful.
Without it, you still have documentation. You just do not have guidance.
That is also why auto-generated llms.txt files tend to feel mediocre. Automation is great at collecting links. It is bad at deciding which few links matter most.
A Practical Keep-or-Skip Checklist
When you are deciding whether a page belongs in llms.txt , ask:
Does this page help someone complete a real implementation or support task?
Is this page canonical, current, and worth trusting?
Would I want an agent to read this before browsing the rest of the docs?
Does it reduce ambiguity, or add more of it?
If the answer is mostly no, leave it out.
That is not a failure. It is the point of curation.
For SaaS docs, a good llms.txt usually includes:
operational docs like errors and rate limits
concept pages that unblock implementation
giant undifferentiated link dumps
The real value comes from deciding what deserves to be there at all.
That is why llms.txt is less like a machine-generated inventory and more like a small editorial artifact for your docs system.
How to Add llms.txt to Your Documentation Site
Adding llms.txt is mechanically simple. The hard part is choosing the right pages and keeping the file useful instead of turning it into a second sitemap.
llms.txt vs sitemap.xml: What Each File Does for AI Discovery
sitemap.xml tells crawlers what exists. llms.txt tells AI agents what matters. If you run docs in 2026, you probably want both.
llms.txt Examples: Real Patterns for API Docs, Help Centers, and Developer Docs
Most llms.txt files are either too vague or too bloated. Here are three patterns that actually fit API docs, help centers, and developer documentation.
DocsAlot Documentation infrastructure for founders who want docs that compound, not accumulate.
