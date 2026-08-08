---
source: "https://docsalot.dev/blog/skill-md-vs-llms-txt"
hn_url: "https://news.ycombinator.com/item?id=49217681"
title: "Skill.md vs. Llms.txt: What Each File Does and When You Need Both"
article_title: "skill.md vs llms.txt: What Each File Does and When You Need Both"
author: "mooreds"
captured_at: "2026-08-08T00:53:26Z"
capture_tool: "hn-digest"
hn_id: 49217681
score: 1
comments: 0
posted_at: "2026-08-08T00:15:10Z"
tags:
  - hacker-news
  - translated
---

# Skill.md vs. Llms.txt: What Each File Does and When You Need Both

- HN: [49217681](https://news.ycombinator.com/item?id=49217681)
- Source: [docsalot.dev](https://docsalot.dev/blog/skill-md-vs-llms-txt)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T00:15:10Z

## Translation

タイトル: Skill.md と Llms.txt: 各ファイルの機能と両方が必要な場合
記事のタイトル: skill.md と llms.txt: 各ファイルの役割と両方が必要な場合
説明: llms.txt と skill.md はさまざまな問題を解決します。 1 つはドキュメントのマップです。もう 1 つは、エージェントがそれらをどのように使用するかについてのコンパクトなプレイブックです。

記事本文:
skill.md と llms.txt: 各ファイルの機能と両方が必要な場合 DocsAlot 製品 ▾ チーム用 ▾ 価格設定 AI 監査ブログ ツール サインイン スタートアップで開始 → DocsAlot 製品 ▾ チーム用 ▾ 価格設定 AI 監査ブログ ツール サインイン スタートアップで開始 → DocsAlot メニュー 製品ヘルプ センター 人間およびサポート ワークフロー向けのオンボーディング ドキュメントをホストしています。 API ドキュメント 開発者向けのリファレンス、例、プレイグラウンド フロー。ドキュメント ベンチマーク AI の読みやすさと検出に関してドキュメントがどのように機能するかを確認します。コンポーネント 再利用可能なドキュメントのパターン、例、実装ガイダンス。すべての機能 ドキュメント、SDK、CLI、MCP、AI 読み取り可能な出力の全機能ライブラリを参照します。 AI 可読ドキュメント AI 可読ドキュメント GitHub ドキュメント自動化 GitHub ドキュメント自動化 API ドキュメント自動化 API ドキュメント自動化 SaaS 用 CLI SaaS 用 MCP SaaS 用ホスト型 MCP サーバー SaaS 用 SDK SaaS 用 SDK ドキュメント用 llms.txt ドキュメント用 llms.txt ヘルプセンター + 開発ドキュメント ヘルプセンター + 開発者ドキュメント チーム用 すべてのチーム創設者、開発者ツール、API、SaaS、およびサポートのポジショニング ページを参照します。創設者 創設者向けのドキュメント 開発者ツール 開発者ツール会社向けのドキュメント API 会社 API 会社向けのドキュメント SaaS + エージェント エージェント採用のための構築を行う SaaS 会社向けのドキュメント サポート チーム サポート チーム向けのドキュメント 価格 AI 監査 ブログ ツール サインイン スタートアップから開始 → ← ブログに戻る / Engineering skill.md と llms.txt: 各ファイルの役割と両方が必要な場合
llms.txt と skill.md はさまざまな問題を解決します。 1 つはドキュメントのマップです。もう 1 つは、エージェントがそれらをどのように使用するかについてのコンパクトなプレイブックです。
多くのチームは、llms.txt と skill.md を競合する標準のように扱っています。そうではありません。それらは異なる層に位置します。
llms.txt て

重要なドキュメントが含まれているモデルです
skill.md はモデルにそれらの使用方法を指示します
一つは地図です。もう一つはプレイブックです。
簡単そうに聞こえますが、両方が必要かどうかに関する混乱のほとんどはこれで解消されます。
いずれかのファイルの背景を長くしたい場合は、「llms.txt とは何ですか?」を参照してください。そして私たちのskill.mdガイド。この投稿は比較レイヤーです。
実際の違いは次のとおりです。
llms.txt のみを公開する場合、モデルはどのページが存在するかを認識している可能性がありますが、依然としてそれらのページを不適切に使用します。
skill.md のみを公開すると、モデルには適切な指示が得られる可能性がありますが、それでもドキュメントの広範なマップが不足します。
だからこそ、それらは互いに補完し合うのです。
llms.txt は、困難な問題が検出である場合にうまく機能します。
このドキュメント サイトの重要なページは何ですか?
どの SDK ページが正規ですか?
ドキュメントのどのセクションが最も重要ですか?
単純な llms.txt ファイルは次のようになります。
コピー 1 # Acme ドキュメント 2
3 ## ここから始めましょう 4
5 - [クイックスタート](https://docs.acme.com/quickstart) 6 - [認証](https://docs.acme.com/authentication) 7 - [エラー](https://docs.acme.com/errors) 8
9 ## SDK 10
11 - [ノード SDK](https://docs.acme.com/sdks/node) 12 - [Python SDK](https://docs.acme.com/sdks/python)
これは、モデルに厳選された開始マップを提供するため便利です。
しかし、それはモデルに次のことを伝えません。
生の REST ドキュメントよりも SDK ドキュメントを優先するかどうか
どのドキュメント サーフェスが特定のタスクに適用されるか
そこでskill.mdの登場です。
skill.md は、難しい問題が判断である場合にうまく機能します。
最初にどのページを読めばいいですか？
API リファレンスの代わりに SDK ドキュメントを使用する必要があるのはどのような場合ですか?
このワークフローではドキュメントのどの部分が正規ですか?
単純な skill.md ファイルは次のようになります。
コピー 1 # Acme API ドキュメント 2
3 `/api-reference` の前に `/quickstart` を読みます。 4 OAuth を実装する前に「/authentication」を読んでください。 5 ノード固有のタスクには「/sdks/node」を使用します。

REST のサンプルの一部。 6
7 文書化されていない Webhook フィールドを作成しないでください。
これはドキュメント サイト全体の地図ではありません。これは、モデルが誤る可能性を低くするための小さな操作上の注意事項です。
llms.txt はモデルの検出に役立ちます
skill.md はモデルの選択に役立ちます
判断のない発見は不完全です。
発見のない判断は狭い。
これは実際の故障モードではっきりとわかります。
llms.txt のみの障害モード
ただし、最初にリファレンス ドキュメントにアクセスし、クイックスタートをスキップし、その後、間違った前提に基づいて構築する可能性があります。
skill.md のみの失敗モード
言語固有の作業については SDK ドキュメントを優先します
文書化されていないフラグを作成しないでください
ただし、ドキュメント セット全体にわたってより広範な検出が必要な場合、ファイルだけでは十分な範囲を提供できない可能性があります。
このため、2 つのファイルは競合するのではなく、お互いを強化します。
初日から必ずしも両方が必要なわけではありません。
あなたの主な問題は見つけやすさです
ドキュメント サイトが大きいか乱雑です
難しいのは、適切なページを表示することだけです
モデルは通常ドキュメントを見つけますが、そのドキュメントの使用方法が間違っています
同じ間違いが繰り返され続ける
製品に複数の表面があるか、ワークフローの順序が明確ではない
発見しやすさだけでなく、エージェントの品質も重視している
API、SDK、CLI、またはヘルプセンターのドキュメントにはさまざまな役割があります
最も深刻な開発者向けドキュメントの設定では、どちらも長期的には正しい答えです。
どちらもエージェント向けの短いファイルであり、ドキュメントの隣に配置されることが多いため、似ていると感じることがあります。
しかし、役立つ内容は異なります。
skill.md が llms.txt を再記述しているだけの場合、それはおそらく弱いです。
llms.txt が、 skill.md のすべてのルーティングとベストプラクティスの作業を実行しようとすると、おそらく過負荷になります。
意味のあるスタックが必要な場合は、ファイルについて次のように考えてください。
ドキュメントの内容: 実際の真実の情報源
MCP (関連する場合): ライブ ツールとアクション
それ

また、どちらのファイルもそれ自体では魔法ではない理由を説明するのにも役立ちます。これらは、基礎となるドキュメントがそれらをサポートするのに十分である場合にのみ役に立ちます。
llms.txt と skill.md は異なる仕事をします。
適切なページを見つけるためにモデルが必要な場合は、llms.txt を使用します。
モデルがこれらのページをより適切な順序で、より適切な判断で使用する必要がある場合は、skill.md を使用します。
ほとんどの実際のドキュメント システムでは、最終的には両方が必要になることを意味します。
標準が楽しいからではなく、モデルが 2 つの異なる間違いを犯すからです。
適切なドキュメントが見つからない
彼らは適切なドキュメントをうまく活用できていない
これらのファイルは、さまざまな角度からこれらの問題に対処します。
skill.md の具体的な側面が必要な場合は、 skill.md の例 を参照してください。スタックの検出側が必要な場合は、 llms.txt と sitemap.xml を読んでください。
2026 年の Mintlify の代替品ベスト 8
開発者ドキュメント、API ドキュメント、Git 同期、AI 読み取り可能なドキュメント、MCP、価格設定、移行に関して、2026 年の Mintlify の最良の代替案を比較します。
API ドキュメントと開発者ツール用の Good skill.md を作成する方法
ほとんどの skill.md ファイルは、ブランディングが多すぎる、ルーティングが不十分、明示的な境界がないなどの同じ理由で失敗します。実際にエージェントの動作を改善する記述方法は次のとおりです。
skill.md の例: クロード コード、カーソル、およびコーデックスの実際のパターン
ほとんどの skill.md ファイルは抽象的すぎて、エージェントが実際の作業を行うのに役立ちません。 API ドキュメント、SDK ドキュメント、および CLI ドキュメントの具体的なパターンとテンプレートを次に示します。
DocsAlot 蓄積するのではなく、複合的なドキュメントを求める創業者のためのドキュメント インフラストラクチャ。

## Original Extract

llms.txt and skill.md solve different problems. One is a map of your docs. The other is a compact playbook for how an agent should use them.

skill.md vs llms.txt: What Each File Does and When You Need Both DocsAlot Product ▾ For teams ▾ Pricing AI Audit Blog Tools Sign in Start with Startup → DocsAlot Product ▾ For teams ▾ Pricing AI Audit Blog Tools Sign in Start with Startup → DocsAlot Menu Product Help Center Hosted onboarding docs for humans and support workflows. API Docs Developer-facing reference, examples, and playground flows. Docs Benchmark See how documentation performs for AI readability and discovery. Components Reusable docs patterns, examples, and implementation guidance. All features Browse the full feature library for docs, SDKs, CLIs, MCP, and AI-readable outputs. AI-readable docs AI-readable documentation GitHub docs automation GitHub docs automation API docs automation API docs automation CLIs for your SaaS CLIs for your SaaS MCP for your SaaS Hosted MCP servers for your SaaS SDKs for your SaaS SDKs for your SaaS llms.txt for docs llms.txt for docs Help center + dev docs Help center + developer docs For teams All teams Browse founder, developer-tools, API, SaaS, and support positioning pages. Founders Documentation for founders Developer tools Documentation for developer-tools companies API companies Documentation for API companies SaaS + agents Documentation for SaaS companies building for agent adoption Support teams Documentation for support teams Pricing AI Audit Blog Tools Sign in Start with Startup → ← Back to blog / Engineering skill.md vs llms.txt: What Each File Does and When You Need Both
llms.txt and skill.md solve different problems. One is a map of your docs. The other is a compact playbook for how an agent should use them.
A lot of teams treat llms.txt and skill.md like competing standards. They are not. They sit at different layers.
llms.txt tells the model where the important docs are
skill.md tells the model how to use them
One is a map. The other is a playbook.
That sounds simple, but it clears up most of the confusion around whether you need both.
If you want the longer background on either file, read What Is llms.txt ? and our skill.md guide . This post is the comparison layer.
Here is the practical difference:
If you only publish llms.txt , the model may know which pages exist but still use them badly.
If you only publish skill.md , the model may get good instructions but still lack a broad map of the docs.
That is why they complement each other.
llms.txt works well when the hard problem is discovery.
what are the important pages on this docs site?
which SDK pages are canonical?
which docs sections matter most?
A simple llms.txt file might look like this:
Copy 1 # Acme Docs 2
3 ## Start Here 4
5 - [ Quickstart ]( https://docs.acme.com/quickstart ) 6 - [ Authentication ]( https://docs.acme.com/authentication ) 7 - [ Errors ]( https://docs.acme.com/errors ) 8
9 ## SDKs 10
11 - [ Node SDK ]( https://docs.acme.com/sdks/node ) 12 - [ Python SDK ]( https://docs.acme.com/sdks/python )
That is useful because it gives the model a curated starting map.
But it does not tell the model:
whether to prefer SDK docs over raw REST docs
which docs surface applies to a given task
That is where skill.md comes in.
skill.md works well when the hard problem is judgment.
which page should I read first?
when should I use the SDK docs instead of the API reference?
what part of the docs is canonical for this workflow?
A simple skill.md file might look like this:
Copy 1 # Acme API Docs 2
3 Read `/quickstart` before `/api-reference` . 4 Read `/authentication` before implementing OAuth. 5 Use `/sdks/node` for Node-specific tasks instead of REST examples. 6
7 Do not invent undocumented webhook fields.
This is not a map of the whole docs site. It is a small operating note that makes the model less likely to go wrong.
llms.txt helps the model discover
skill.md helps the model choose
Discovery without judgment is incomplete.
Judgment without discovery is narrow.
You can see this pretty clearly in real failure modes.
Failure mode with only llms.txt
But it may still jump into the reference docs first, skip the quickstart, and then build on the wrong assumptions.
Failure mode with only skill.md
prefer SDK docs for language-specific work
do not invent undocumented flags
But if it needs broader discovery across the docs set, the file may not give enough coverage on its own.
That is why the two files reinforce each other rather than competing.
You do not always need both on day one.
your main problem is discoverability
your docs site is large or messy
the hard part is just surfacing the right pages
the model usually finds the docs but uses them badly
the same mistakes keep repeating
your product has multiple surfaces or non-obvious workflow order
you care about agent quality, not just discoverability
your API, SDK, CLI, or help-center docs have different roles
For most serious developer-docs setups, both is the right long-term answer.
They can feel similar because both are short, agent-facing files and both often sit next to the docs.
But their useful content is different.
If your skill.md just restates your llms.txt , it is probably weak.
If your llms.txt tries to do all the routing and best-practice work of skill.md , it is probably getting overloaded.
If you want a stack that makes sense, think about the files like this:
docs content: the actual source of truth
MCP, if relevant: live tools and actions
It also helps explain why neither file is magical on its own. They only help if the underlying docs are good enough to support them.
llms.txt and skill.md do different jobs.
Use llms.txt when you need the model to find the right pages.
Use skill.md when you need the model to use those pages in a better order and with better judgment.
For most real documentation systems, that means you eventually want both.
Not because standards are fun, but because models make two different mistakes:
they fail to find the right docs
they fail to use the right docs well
These files address those problems from different angles.
If you want the concrete side of skill.md , read skill.md Examples . If you want the discovery side of the stack, read llms.txt vs sitemap.xml .
8 Best Mintlify Alternatives in 2026
Compare the best Mintlify alternatives in 2026 for developer docs, API docs, Git sync, AI-readable docs, MCP, pricing, and migration.
How to Write a Good skill.md for API Docs and Developer Tools
Most skill.md files fail for the same reasons: too much branding, not enough routing, and no explicit boundaries. Here is how to write one that actually improves agent behavior.
skill.md Examples: Real Patterns for Claude Code, Cursor, and Codex
Most skill.md files are too abstract to help an agent do real work. Here are concrete patterns and templates for API docs, SDK docs, and CLI docs.
DocsAlot Documentation infrastructure for founders who want docs that compound, not accumulate.
