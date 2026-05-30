---
source: "https://docsalot.dev/blog/llms-txt-examples"
hn_url: "https://news.ycombinator.com/item?id=48332306"
title: "Llms.txt Examples: Real Patterns for API Docs, Help Centers, and Developer Docs"
article_title: "llms.txt Examples: Real Patterns for API Docs, Help Centers, and Developer Docs"
author: "fazkan"
captured_at: "2026-05-30T11:35:39Z"
capture_tool: "hn-digest"
hn_id: 48332306
score: 1
comments: 0
posted_at: "2026-05-30T03:37:17Z"
tags:
  - hacker-news
  - translated
---

# Llms.txt Examples: Real Patterns for API Docs, Help Centers, and Developer Docs

- HN: [48332306](https://news.ycombinator.com/item?id=48332306)
- Source: [docsalot.dev](https://docsalot.dev/blog/llms-txt-examples)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T03:37:17Z

## Translation

タイトル: Llms.txt 例: API ドキュメント、ヘルプセンター、開発者ドキュメントの実際のパターン
記事のタイトル: llms.txt 例: API ドキュメント、ヘルプセンター、開発者ドキュメントの実際のパターン
説明: ほとんどの llms.txt ファイルは、あいまいすぎるか、肥大化しすぎます。 API ドキュメント、ヘルプ センター、開発者ドキュメントに実際に適合する 3 つのパターンを次に示します。

記事本文:
llms.txt の例: API ドキュメント、ヘルプ センター、および開発者ドキュメントの実際のパターン DocsAlot Docs ベンチマーク機能 ▾ すべての機能ページ DocsAlot の背後にあるドキュメント ワークフローを参照します。 AI 可読ドキュメント AI 可読ドキュメント GitHub ドキュメント自動化 GitHub ドキュメント自動化 API ドキュメント自動化 API ドキュメント自動化 SaaS 用 CLI SaaS 用 MCP SaaS 用ホスト型 MCP サーバー SaaS 用 SDK SaaS 用 SDK SaaS 用 llms.txt ドキュメント用 llms.txt ドキュメント用 ヘルプセンター + 開発者用ドキュメント ヘルプセンター + 開発者用ドキュメント チーム用▾ すべての対象者ページ 創設者、API、サポート、および開発者ツールの位置付けページを参照します。創設者 創設者向けドキュメント 開発者ツール 開発者ツール会社向けドキュメント API 会社 API 会社向けドキュメント SaaS + エージェント エージェント採用のための構築を行う SaaS 会社向けドキュメント サポート チーム サポート チーム向けドキュメント 価格 変更ログ ブログ サインイン 無料で開始 → メニュー ヘルプ センター ベンチマーク コンポーネント API ドキュメント 機能 ▾ すべての機能ページ AI 読み取り可能なドキュメント GitHub ドキュメントの自動化 API ドキュメントの自動化 SaaS 用の CLI SaaS 用の MCP SaaS 用の SDK ドキュメント用の llms.txt ヘルプ センター + 開発ドキュメント チーム向け ▾ すべての対象者ページ 創設者 開発者ツール API 企業 SaaS + エージェント サポート チーム 価格設定 ブログ 変更ログ サインイン 無料で始める → ← ブログ / エンジニアリングに戻る llms.txt 例: API ドキュメント、ヘルプ センター、および開発者ドキュメントの実際のパターン
ほとんどの llms.txt ファイルは、あいまいすぎるか、肥大化しすぎます。 API ドキュメント、ヘルプ センター、開発者ドキュメントに実際に適合する 3 つのパターンを次に示します。
オンラインの llms.txt のサンプルのほとんどは一般的すぎて役に立ちません。いくつかのリンク、おそらくはクイックスタート、おそらくは認証ページが表示され、それで終わりです。形式を理解するにはそれで十分です。私

複数の対象者と数十、数百のページを含む実際のドキュメント サーフェスを管理している場合は、それだけでは十分ではありません。
難しいのはファイルを書き込むことではありません。難しいのは、自分の種類のドキュメントにどのような形をとるべきかを決めることです。
まだ基本事項を読んでいない場合は、「llms.txt とは何ですか?」から始めてください。 SaaS Docs チームのための実践ガイド。この投稿では、ファイルが何であるかをすでに理解していることを前提として、より有用な質問、つまり、適切なファイルとは実際にどのようなものであるべきかということに焦点を当てます。
構文よりパターンが重要
構文は簡単な部分です。マークダウンの見出しとリンクはボトルネックではありません。
ドキュメント サイトはモデルに対してどのような役割を果たしていますか?
API リファレンスには、ヘルプ センターとは異なる優先順位があります。開発者ドキュメント ポータルには、製品サポート サイトとは異なる優先順位があります。すべてに同じ llm​​s.txt 構造を使用すると、通常、技術的には有効でも、編集的には弱いものになります。
したがって、1 つの普遍的なテンプレートではなく、パターンで考える方がよいでしょう。
これは、技術的な SaaS 製品で最も一般的なケースです。
モデルは通常、次のいずれかに答える必要があります。
エラーやレート制限はどのように処理すればよいですか?
どの SDK または統合パスを使用すればよいですか?
つまり、llms.txt は、最初に成功したリクエストと一般的な実装タスクに合わせて最適化する必要があります。
コピー 1 # Acme API ドキュメント 2
3 Acme の REST API の開発者向けドキュメント。認証、4 つのエンドポイント、ページネーション、エラー、Webhook、SDK について説明します。 5
6 ## ここから始めてください 7
8 - [クイックスタート]( https://docs.acme.com/quickstart ) : 最初の API リクエストを作成します 9 - [認証]( https://docs.acme.com/authentication ) : API キー、OAuth、およびスコープ 10 - [ API の概要 ]( https://docs.acme.com/api ) : コア リソースとリクエスト モデル 11
12 ## コアタスク 13
14 - [顧客の作成]( https://docs.ac

me.com/customers/create ) 15 - [請求書の一覧表示](https://docs.acme.com/invoices/list ) 16 - [Webhook の処理](https://docs.acme.com/webhooks ) 17
18 ## 運用ドキュメント 19
20 - [エラー]( https://docs.acme.com/errors ) 21 - [レート制限]( https://docs.acme.com/rate-limits ) 22 - [ページネーション]( https://docs.acme.com/pagination ) 23
24 ## SDK 25
26 - [ノード SDK](https://docs.acme.com/sdks/node) 27 - [Python SDK](https://docs.acme.com/sdks/python)
なぜこれが機能するのか
それは、リクエストが成功するまでの最短パスから始まります。
これには、モデルが実装の途中で必要になることが多い操作ページが含まれています。
ナビゲーション ラベルだけでなく、タスク指向のグループ化を使用します。
チームはトップレベルの参照カテゴリを llms.txt にダンプすることがよくあります。
これは整理されているように見えますが、モデルにはあまり役に立ちません。良いページがどこにあるのかを推測する必要があります。
API ドキュメントについては、独自の意見を持ってください。モデルは、完全にミラーリングされたサイドバーよりも、 Authentication 、 Errors 、および Webhook の方がメリットが得られます。
ヘルプ センターは、主なタスクが通常は診断であるため、異なります。
この製品を別の製品に接続するにはどうすればよいですか?
ヘルプ センターに最適な llms.txt は、API プリミティブよりも、一般的なワークフロー、トラブルシューティング、アカウント レベルのタスクに重点を置いています。
コピー 1 # Acme ヘルプ センター 2
3 Acme の製品サポート ドキュメント。セットアップ、請求、権限、4 つの統合、トラブルシューティング、アカウント管理について説明します。 5
6 ## ここから始めてください 7
8 - [はじめに]( https://help.acme.com/getting-started ) 9 - [ ワークスペースのセットアップ ]( https://help.acme.com/workspace-setup ) 10 - [ チームメイトを招待 ]( https://help.acme.com/invite-teammates ) 11
12 ## 一般的な問題 13
14 - [ログインの問題](https://help.acme.com/login-problems) 15 - [権限エラー](https://help.acme.com/permission-errors) 16 - [I

統合の失敗 ]( https://help.acme.com/integration-failures ) 17
18 ## 主要なワークフロー 19
20 - [請求と請求書]( https://help.acme.com/billing ) 21 - [ユーザー ロール]( https://help.acme.com/user-roles ) 22 - [Slack 統合]( https://help.acme.com/slack )
なぜこれが機能するのか
これは、サポートの質問が実際にどのように表現されているかを反映しています。
製品のマーケティング構造よりもトラブルシューティングと構成を優先します。
これにより、モデルにすべてのページをスキャンさせることなく、アカウント レベルの問題へのパスが与えられます。
多くのヘルプ センターは、内部チームの所有権ごとにコンテンツを整理しています。
ドキュメントチームはそう考えているのかもしれません。ユーザーやモデルが質問をどのように表現するかではありません。
ヘルプ センターの場合は、内部分類ではなく、ワークフローと障害モードを重視します。
開発者向けドキュメントは通常、API リファレンスと製品ガイダンスの間のどこかに位置します。
導入またはホスティングの手順
ここでの最良の llms.txt は、最初のエンドポイントを見つけるだけでなく、モデルが適切なメンタル マップを構築するのに役立ちます。
コピー 1 # Acme Developer Docs 2
3 Acme 上に構築するための技術文書。ローカル開発、4 つの SDK の使用法、CLI コマンド、デプロイメント、認証、アーキテクチャについて説明します。 5
6 ## ここから始めてください 7
8 - [クイックスタート](https://developers.acme.com/quickstart) 9 - [ローカル開発](https://developers.acme.com/local-development) 10 - [認証](https://developers.acme.com/authentication) 11
12 ## コアコンセプト 13
14 - [アーキテクチャの概要](https://developers.acme.com/architecture) 15 - [データ モデル](https://developers.acme.com/data-model) 16 - [環境](https://developers.acme.com/environments) 17
18 ## 構築と出荷 19
20 - [ CLI リファレンス ]( https://developers.acme.com/cli ) 21 - [ デプロイメント ]( https://developers.acme.com/deployments ) 22 - [ Webhook ]( https://developers.acme.com/webh

わかりました ) 23
24 ## 参考文献 25
26 - [ REST API ]( https://developers.acme.com/api ) 27 - [ ノード SDK ]( https://developers.acme.com/sdks/node ) 28 - [ Python SDK ]( https://developers.acme.com/sdks/python )
なぜこれが機能するのか
これにより、モデルに方向性と実装パスの両方が与えられます。
コードを記述する前にモデルが理解する必要があるアーキテクチャが製品にある場合、コンセプト ページが含まれます。
すべてをフラット化するのではなく、ワークフローから参照を分離します。
チームは開発者ドキュメントを純粋な API リファレンスのように扱い、アーキテクチャ ページを完全に省略することがよくあります。
これは単純な API では機能します。これは、エンドポイントを呼び出すだけでなく、正しいパスを選択することが主な困難であるプラットフォーム、CLI、展開システム、または製品ごとに分類されます。
製品に意味のあるコンセプトがある場合は、それを含めてください。
これら 3 つのパターンで何が変わるのか
違いについて考える最も簡単な方法は次のとおりです。
この表は通常、「重要なページを含める」などの一般的なアドバイスよりも役立ちます。
すべてのドキュメント チームは、重要なページを含める必要があることを知っています。難しいのは、走る路面の種類にとって「重要」が何を意味するのかを判断することです。
パターン間でコピーしてはいけないもの
ほとんどの llms.txt ファイルはここで問題が発生します。
ユーザーは、非常に異なるドキュメント サーフェイス間で同じテンプレートをコピーします。
主にアカウント設定を処理するヘルプセンターで、レート制限やページネーションなどの API 中心のセクションを使用する
開発者ドキュメントの最上位構造としてサポート スタイルのトラブルシューティング セクションを使用する
主にクイックスタートと認証を必要とする単純な API のアーキテクチャ ドキュメントを最初にリストします。
ファイルには、そのサーフェスに対する最も一般的で最も価値のあるエージェント タスクが反映されている必要があります。
そうでない場合、ファイルはまだ存在しますが、あまり有用な作業は行っていません。
llms.txt を配布する前に、次のように質問してください。
優秀なエンジニアであれば、

私のチームには、AI エージェントに適切なページを指示する時間が 60 秒ありました。彼らはこの正確なリストを選択するでしょうか?
答えが「いいえ」の場合は、編集を続けてください。
優れた llms.txt は単なる機械可読ファイルではないため、これは正しい標準です。コンパクトな形で捉えられた編集上の判断です。
最初の llms.txt を公開する場合は、すぐにオーバーフィットしないでください。
1 つまたは 2 つのタスク指向のグループ
次に、ドキュメントの表面に実際にどのような種類の質問が寄せられるかを確認してから、内容を調整します。
通常、最大の改善は、一般的なリンク ダンプからドキュメントの種類に一致する構造に移行することで得られます。
それだけでも、このファイルはさらに便利になります。
例は役に立ちますが、llms.txt はまだ単なる検出レイヤーにすぎません。
その後、次の質問が行われます。
ファイルを sitemap.xml にどのように関連付ける必要がありますか?
ドキュメントが進化するにつれて、どうやってドキュメントを更新し続けるのでしょうか?
エージェントはインデックス以外に何が必要ですか?
これらについては個別に説明します。より大きなシステムの観点については、「llms.txt だけでは不十分」を参照してください。基本については、「llms.txt とは何ですか?」から始めてください。 SaaS Docs チームのための実践ガイド。
llms.txtとは何ですか? SaaS ドキュメント チームのための実践ガイド
llms.txt は、機械可読なドキュメントのインデックスです。これが何をするのか、何を含めるべきか、そしてそれ自体では解決できないものを次に示します。
SaaS 用のリモート MCP サーバーをセットアップする方法
MCP は理論的には簡単ですが、実際には面倒です。このガイドでは、MCP とは何か、SaaS 製品にとってリモート MCP が重要な理由、主な展開オプション、ユーザーにローカル設定を強制せずにトークンレス オンボーディングを行う方法について説明します。
誰にもお金を払わずにPHでオーガニック投票をどれくらい獲得できますか?ケーススタディ
私はハンター、インフルエンサー、立ち上げコンサルタントにお金を払わずに、Product Hunt で狭い製品を立ち上げました。結果は 51 票で 19 位に終わりました。役に立ったのは、どのようなオーガニック ローンチが行われるかを学んだことです。

実際に依存します。
DocsAlot 蓄積するのではなく、複合的なドキュメントを求める創業者のためのドキュメント インフラストラクチャ。

## Original Extract

Most llms.txt files are either too vague or too bloated. Here are three patterns that actually fit API docs, help centers, and developer documentation.

llms.txt Examples: Real Patterns for API Docs, Help Centers, and Developer Docs DocsAlot Docs Benchmark Features ▾ All feature pages Browse the documentation workflows behind DocsAlot. AI-readable docs AI-readable documentation GitHub docs automation GitHub docs automation API docs automation API docs automation CLIs for your SaaS CLIs for your SaaS MCP for your SaaS Hosted MCP servers for your SaaS SDKs for your SaaS SDKs for your SaaS llms.txt for docs llms.txt for docs Help center + dev docs Help center + developer docs For teams ▾ All audience pages Browse founder, API, support, and developer-tools positioning pages. Founders Documentation for founders Developer tools Documentation for developer-tools companies API companies Documentation for API companies SaaS + agents Documentation for SaaS companies building for agent adoption Support teams Documentation for support teams Pricing Changelog Blog Sign in Start for free → Menu Help Center Benchmark Components API Docs Features ▾ All feature pages AI-readable docs GitHub docs automation API docs automation CLIs for your SaaS MCP for your SaaS SDKs for your SaaS llms.txt for docs Help center + dev docs For teams ▾ All audience pages Founders Developer tools API companies SaaS + agents Support teams Pricing Blog Changelog Sign in Start for free → ← Back to blog / Engineering llms.txt Examples: Real Patterns for API Docs, Help Centers, and Developer Docs
Most llms.txt files are either too vague or too bloated. Here are three patterns that actually fit API docs, help centers, and developer documentation.
Most llms.txt examples online are too generic to be useful. They show a couple of links, maybe a quickstart, maybe an auth page, and call it a day. That is fine for understanding the format. It is not enough if you are maintaining a real docs surface with multiple audiences and dozens or hundreds of pages.
The hard part is not writing the file. The hard part is deciding what shape it should take for your kind of documentation.
If you have not read the basics yet, start with What Is llms.txt ? A Practical Guide for SaaS Docs Teams . This post assumes you already understand what the file is and focuses on the more useful question: what should a good one actually look like?
The Pattern Matters More Than the Syntax
The syntax is the easy part. Markdown headings and links are not the bottleneck.
What job is your docs site doing for the model?
An API reference has different priorities than a help center. A developer docs portal has different priorities than a product support site. If you use the same llms.txt structure for all of them, you usually end up with something that is technically valid and editorially weak.
So instead of one universal template, it is better to think in patterns.
This is the most common case for technical SaaS products.
The model usually needs to answer one of these:
how do I handle errors or rate limits?
which SDK or integration path should I use?
That means your llms.txt should optimize for first successful requests and common implementation tasks.
Copy 1 # Acme API Docs 2
3 Developer documentation for Acme's REST API. Covers authentication, 4 endpoints, pagination, errors, webhooks, and SDKs. 5
6 ## Start Here 7
8 - [ Quickstart ]( https://docs.acme.com/quickstart ) : Make your first API request 9 - [ Authentication ]( https://docs.acme.com/authentication ) : API keys, OAuth, and scopes 10 - [ API Overview ]( https://docs.acme.com/api ) : Core resources and request model 11
12 ## Core Tasks 13
14 - [ Create a customer ]( https://docs.acme.com/customers/create ) 15 - [ List invoices ]( https://docs.acme.com/invoices/list ) 16 - [ Handle webhooks ]( https://docs.acme.com/webhooks ) 17
18 ## Operational Docs 19
20 - [ Errors ]( https://docs.acme.com/errors ) 21 - [ Rate Limits ]( https://docs.acme.com/rate-limits ) 22 - [ Pagination ]( https://docs.acme.com/pagination ) 23
24 ## SDKs 25
26 - [ Node SDK ]( https://docs.acme.com/sdks/node ) 27 - [ Python SDK ]( https://docs.acme.com/sdks/python )
Why this works
It starts with the shortest path to a successful request.
It includes operational pages that models often need in the middle of implementation.
It uses task-oriented grouping, not just navigation labels.
Teams often dump top-level reference categories into llms.txt :
That looks organized, but it does not help the model much. It still has to guess where the good pages are.
For API docs, be opinionated. A model benefits more from Authentication , Errors , and Webhooks than from a perfectly mirrored sidebar.
Help centers are different because the main tasks are usually diagnostic:
how do I connect this product to another one?
The best llms.txt for a help center is less about API primitives and more about common workflows, troubleshooting, and account-level tasks.
Copy 1 # Acme Help Center 2
3 Product support documentation for Acme. Covers setup, billing, permissions, 4 integrations, troubleshooting, and account management. 5
6 ## Start Here 7
8 - [ Getting Started ]( https://help.acme.com/getting-started ) 9 - [ Workspace Setup ]( https://help.acme.com/workspace-setup ) 10 - [ Invite Teammates ]( https://help.acme.com/invite-teammates ) 11
12 ## Common Issues 13
14 - [ Login Problems ]( https://help.acme.com/login-problems ) 15 - [ Permission Errors ]( https://help.acme.com/permission-errors ) 16 - [ Integration Failures ]( https://help.acme.com/integration-failures ) 17
18 ## Key Workflows 19
20 - [ Billing and Invoices ]( https://help.acme.com/billing ) 21 - [ User Roles ]( https://help.acme.com/user-roles ) 22 - [ Slack Integration ]( https://help.acme.com/slack )
Why this works
It reflects how support questions are actually phrased.
It prioritizes troubleshooting and configuration over product marketing structure.
It gives the model a path into account-level problems without making it scan every page.
Many help centers organize content by internal team ownership:
That may be how the docs team thinks. It is not how users or models phrase questions.
For a help center, lean toward workflows and failure modes, not internal taxonomy.
Developer docs usually sit somewhere between API reference and product guidance.
deployment or hosting instructions
The best llms.txt here should help a model build the right mental map, not just find the first endpoint.
Copy 1 # Acme Developer Docs 2
3 Technical documentation for building on Acme. Covers local development, 4 SDK usage, CLI commands, deployment, authentication, and architecture. 5
6 ## Start Here 7
8 - [ Quickstart ]( https://developers.acme.com/quickstart ) 9 - [ Local Development ]( https://developers.acme.com/local-development ) 10 - [ Authentication ]( https://developers.acme.com/authentication ) 11
12 ## Core Concepts 13
14 - [ Architecture Overview ]( https://developers.acme.com/architecture ) 15 - [ Data Model ]( https://developers.acme.com/data-model ) 16 - [ Environments ]( https://developers.acme.com/environments ) 17
18 ## Build and Ship 19
20 - [ CLI Reference ]( https://developers.acme.com/cli ) 21 - [ Deployments ]( https://developers.acme.com/deployments ) 22 - [ Webhooks ]( https://developers.acme.com/webhooks ) 23
24 ## References 25
26 - [ REST API ]( https://developers.acme.com/api ) 27 - [ Node SDK ]( https://developers.acme.com/sdks/node ) 28 - [ Python SDK ]( https://developers.acme.com/sdks/python )
Why this works
It gives the model both orientation and implementation paths.
It includes concept pages when the product has architecture the model needs to understand before writing code.
It separates references from workflows instead of flattening everything.
Teams often treat developer docs like pure API reference and leave out architecture pages entirely.
That works for simple APIs. It breaks down for platforms, CLIs, deployment systems, or products where the main difficulty is choosing the right path, not just calling an endpoint.
If your product has meaningful concepts, include them.
What Changes Across These Three Patterns
Here is the simplest way to think about the differences:
That table is usually more useful than generic advice like "include the important pages."
Every docs team knows they should include important pages. The hard part is deciding what “important” means for the kind of surface you run.
What Not to Copy Between Patterns
This is where most llms.txt files go wrong.
People copy the same template between very different docs surfaces.
using API-centric sections like Rate Limits and Pagination in a help center that mostly handles account setup
using support-style troubleshooting sections as the top-level structure for developer docs
listing architecture docs first for a simple API that mostly needs Quickstart and Authentication
The file should reflect the most common and most valuable agent tasks for that surface.
If it does not, the file still exists, but it is not doing much useful work.
Before shipping llms.txt , ask this:
If a good engineer on my team had 60 seconds to point an AI agent at the right pages, would they choose this exact list?
If the answer is no, keep editing.
That is the right standard because a good llms.txt is not just a machine-readable file. It is editorial judgment captured in a compact form.
If you are publishing your first llms.txt , do not overfit immediately.
one or two task-oriented groups
Then refine it once you see what kinds of questions your docs surface actually gets.
The biggest improvement usually comes from moving from a generic link dump to a structure that matches your docs type.
That alone makes the file much more useful.
Examples are helpful, but llms.txt is still just the discovery layer.
After that, the next questions are:
how should the file relate to sitemap.xml ?
how do you keep it updated as docs evolve?
what else do agents need besides an index?
We will cover those separately. For the bigger systems view, read llms.txt Isn't Enough . For the basics, start with What Is llms.txt ? A Practical Guide for SaaS Docs Teams .
What Is llms.txt? A Practical Guide for SaaS Docs Teams
llms.txt is a machine-readable index for documentation. Here's what it does, what to include, and what it won't solve on its own.
How to Set Up a Remote MCP Server for Your SaaS
MCP is straightforward in theory and messy in practice. This guide covers what MCP is, why remote MCP matters for SaaS products, the main deployment options, and how to get to tokenless onboarding without forcing users through local setup.
How Many Organic Votes Can You Get on PH Without Paying Anyone? A Case Study
I launched a narrow product on Product Hunt without paying a hunter, influencers, or a launch consultant. It ended with 51 votes and a 19th-place finish, and the useful part was learning what organic launches actually depend on.
DocsAlot Documentation infrastructure for founders who want docs that compound, not accumulate.
