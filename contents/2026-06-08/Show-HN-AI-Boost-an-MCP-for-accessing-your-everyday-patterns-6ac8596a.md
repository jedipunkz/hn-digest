---
source: "https://www.ai-boost.io/"
hn_url: "https://news.ycombinator.com/item?id=48442307"
title: "Show HN: AI Boost – an MCP for accessing your everyday patterns"
article_title: "AI Boost — Your Expertise Layer for LLM Agents"
author: "npiano"
captured_at: "2026-06-08T07:51:03Z"
capture_tool: "hn-digest"
hn_id: 48442307
score: 1
comments: 0
posted_at: "2026-06-08T07:30:37Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Boost – an MCP for accessing your everyday patterns

- HN: [48442307](https://news.ycombinator.com/item?id=48442307)
- Source: [www.ai-boost.io](https://www.ai-boost.io/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T07:30:37Z

## Translation

タイトル: Show HN: AI Boost – 日常のパターンにアクセスするための MCP
記事のタイトル: AI Boost — LLM エージェントの専門知識レイヤー
説明: 何年もかけて習得したものをキャプチャします。数秒で注入します。準備ができたら共有してください。 AI Boost は、LLM エージェントの専門知識レイヤーであり、Claude Code、Cursor、および任意の MCP クライアントで動作します。
HN テキスト: こんにちは、HN。LLM エージェントに対する私の最大の不満は、すべてのセッションが空白で始まり、毎回自分のパターンと慣例を再説明しなければならないことです。私は「X リポジトリを見て、Y に使用されているパターンをコピーする」といったことをしていることがよくありますが、これにはリポジトリがローカル マシン上にある必要があります。私にとって個人的に大きなことは次のとおりです。
- 認証フロー
- AWS の Terraform パターン ルールとスキルはある程度役立ちますが、複数のエージェント、プロジェクト、マシン間で同期するのは困難です。メモリは興味深いものですが、構造化されていないためノイズが多すぎます。そこで、MCP サーバーで個人ライブラリとして機能する AI Boost を構築しました。あなたはエージェントに何を保存するかを指示します (現時点ではテキスト ファイルまたはパブリック github リポジトリ。私はプライベート github リポジトリを含む他のオプションに取り組んでいます)。それを「ブースター」としてパッケージ化し、キーワードと埋め込みでインデックスを作成し、次に一致するタスクを開始するときに、エージェントがそれを表示します。現時点では:
- ブースターはデフォルトではプライベートです。あなたのアカウントのみがアクセスできます。
- 必要に応じてコミュニティ マーケットプレイスに公開し、注入ごとにクレジットを獲得できます (これはまだ完全に準備ができていませんが、近日中に提供される予定です)。
- これは MCP サーバーであるため、Cursor、Claude Code、および MCP をサポートするクライアントで動作するはずです。リンク: https://ai-boost.io (MCP URL は https://mcp.ai-boost.io/mcp です) ぜひご意見をお聞かせください。特に「タスクを開始する前の自動提案」動作が便利だと感じるか、邪魔だと感じるかどうか

実際にやってみました。

記事本文:
AI Boost — LLM エージェント AI の専門知識レイヤー
LLM エージェントの専門知識層。
AI の時代には専門知識のルールを変える必要がある
専門家は、自分たちが管理していないツールに自分の知識をますます提供することが求められています。私たちは、信頼は想定されるものではなく、獲得されるものであると考えています。これらは利用規約ではありません。これらは、ある開発者から別の開発者への約束です。
プライベート ブースターは、インデックス付けされたり、販売されたり、他の人に公開されたり、モデルのトレーニングに使用されたりすることはありません。保存したものはすべてアカウントに残ります。
ブースター、バージョン、またはアカウント全体をいつでも削除できます。削除は永久かつ即時です。アーカイブはありません。
価格変更、ポリシーの更新、プラットフォームの決定は事前に発表されます。あなたに行動する時間を与えずにルールを変更することはありません。
私たちは全角ダッシュが有効で便利な句読点であると信じており、他の人が何と言おうと今後もこれを使用し続けます。
あなたの専門知識。一度捕獲されました。どこでもご利用いただけます。
プライベートブースターはあなただけのものです。あらゆるセッションから専門知識を取得します。 AI Boost に埋め込まれているため、すぐに取得できます。エージェントはコンテキストが一致した瞬間にそれを表示します。
何を保存したいのかをエージェントに伝えてください。現在のコンテキストを読み取り、メタデータを提案し、数秒でプライベート ブースターを作成します。
AI Boost はキーワードおよびセマンティック検索のブースターを処理するため、今後のすべてのセッションでブースターを正確に取得できます。
保存されたパターンに一致するタスクを開始すると、エージェントは開始前にブースターを表示します。あなたは承認します、それは注入します。
一度保存すると、以降のセッションごとに利用可能
何を保存するかをエージェントに伝えてください。残りの部分は、次回の通知も含めて処理されます。
AI Boost Agent カスタム認証ミドルウェア パターンをプライベート ブースターとして保存します
現在の認証ミドルウォーを読んでみましょう

e ファイルを使用してパターンを抽出します。
わかった。ファイルに基づいて、私が提案するブースターの詳細は次のとおりです。
名前: カスタム認証ミドルウェア パターン
説明: JWT 検証、ロールベースのガード、リクエスト コンテキスト インジェクション
可視性は private に設定されます。保存することを確認するか、変更する内容をお知らせください。
プライベート価格 0クレジット（個人使用） 確認済み
ai-boostでcreate_boosterを実行しました
ライブラリに保存されました カスタム認証ミドルウェア パターン ブースターが、custom-auth-middleware-pattern の下のプライベート ライブラリに保存されました。これはあなたにのみ表示され、今後のセッションで自動的に提案されます。
AI Boost Agent — 新しいセッション 保護されたルートを新しい決済サービスに追加します
始める前に、ライブラリに関連するブースターがあるかどうかを確認させてください。
ai-boost で search_boosters を実行しました
カスタム認証ミドルウェア パターン
保存されている認証ミドルウェア ブースターを見つけました。始める前に注射しましょうか？
認証ミドルウェア規約がロードされます。このタスク全体を通じて、JWT 検証とロールベースのガード パターンに従います。
記憶システムは歴史を記録しますが、それを整理することはできません。エージェントが必要とする専門知識 (慣例、落とし穴、習得するのに何年もかかった教訓) は、セッション間で居場所がありません。
すべてのセッションはゼロから始まります 毎日使用するパターン (認証フロー、インフラ規約、API 設計ルール) は人間の記憶の中に存在します。エージェントはそれらにアクセスできません。
メモリはキュレーションではありません メモリ プラグインはすべてを記録します。つまり、すべてを取得します。信号対雑音比により、エージェントが実際に必要とする特定の構造化された知識を注入する際の信頼性が低くなります。
構造化された専門知識を共有する方法がない ドメインの専門家は、非常に貴重な知識を持っています。現在、そのエクスペリエンスをパッケージ化、配布、および注入するメカニズムはありません。

必要なときに LLM に組み込むことができます。
エージェントに規則、パターン、またはワークフローについて伝えます。 AI Boost は、これをブースターとしてパッケージ化し、キーワードとセマンティック埋め込みを使用してインデックスを作成します。
次回のセッションでは、エージェントはタスクを開始する前に AI Boost に自動的にクエリを実行します。関連するブースター (あなた自身のブースターまたはコミュニティからのブースター) が存在する場合、それはあなたの承認のために表示されます。
あなたは承認します。完全なコンテキスト ペイロードは、エージェントのコンテキスト ウィンドウに直接流れます。規則、パターン、例: すぐに利用できます。コピー＆ペーストやセットアップは必要ありません。
AI Boost は MCP サーバーです。これをエージェント クライアントに一度追加すると、すべてのプロジェクトで自動的に使用できるようになります。
ホーム ディレクトリまたはプロジェクト ルートで .cursor/mcp.json を作成または編集します。
{
"mcpサーバー": {
"aiブースト": {
"url": "https://mcp.ai-boost.io/mcp",
"トランスポート": "ストリーミング可能な http"
}
}
3 GitHub OAuth で認証する
初めて使用するとき、カーソルによりブラウザのタブが開きます。 GitHub でサインインし、アクセスを許可すると、自動的にリダイレクトされます。
クレジットを補充してブースターの使用を開始する
app.ai-boost.io/billing にアクセスしてクレジットを追加し、タスクに関連するブースターを検索するように Cursor に依頼します。
エージェントに公開したい内容を伝えるだけです。残りはそれが処理します。
AI Boost Agent このリポジトリをブースターとして公開したいと考えています
これを AI Boost Marketplace でブースターとして公開するのをお手伝いします。提出フローの最初のステップから始めましょう。
パブリック GitHub リポジトリへのリンク
このリポジトリの GitHub リモート URL を確認してみましょう。
実行 git で git リモート URL を確認する
リポジトリは AI-Boost-IO/ Stripe-payment-integration-guide です。詳細を自動提案するために、GitHub API を取得してみましょう。
GitHub リポジトリに基づいた私の提案は次のとおりです。変更したい場合はお知らせください。
名前: Stripe Payment 統合ガイド
説明: クイックセット

Stripe 支払いのアップガイド
バージョン: 1.0.0 (リリースが見つからず、デフォルト)
残りの詳細については次のとおりです。
ブースター提出の完全な概要は次のとおりです。続行する前に確認してください。
すべてが正しく見えますか? 「確認しました」と返信して送信するか、変更内容をお知らせください。
まず、送信ツールのスキーマを確認してみましょう。
Ran ai-boost のリポジトリからブースターを送信
ブースターは提出され、承認されました。要約は次のとおりです。
承認済み Stripe Payment Integration Guide ブースターは現在、AI Boost Marketplace で公開されており、slug Stripe-payment-integration-guide で見つけることができます。
AI の時代では人間の専門的な経験を見つけるのは困難
LLM は、見たことのないものから学ぶことはできません。あなたが苦労して勝ち取ったパターンは、まさに他の開発者が必要としているものです。準備ができたら、プライベート ブースターを公開します。これらは、コミュニティ全体の関連するすべてのセッションで自動的に表示されます。
開発者は AI によって生成された知識にますます囲まれるようになっています。実際の人間の経験に裏付けられたブースターは、まったく異なるカテゴリーです。共有するということは、モデルでは絶対に再現できないものに貢献していることになります。
プライベート ブースターはすでにライブラリに存在します。公開設定をパブリックに切り替えるだけです。
品質を確保するために、自動化された LLM 敵対的レビューと人間による承認が行われます。
コミュニティ全体で検索できるようになります。エージェントがそれに一致する開発者は誰でもそれを注入できます。
ブースターを注入するたびに AI Boost クレジットを獲得できます。これは、コミュニティへの貢献を直接表すものです。
Terraform で管理されたインフラストラクチャ (VPC、EC2、S3、IAM、EIP) を使用して、Dockerized Django アプリを AWS EC2 にデプロイするためのセットアップ ガイド。
9 14 Django Channels + Celery + Next.js + Terraform/AWS 0 クレジット モノリポジトリ SaaS バックエンドの正規パターンを文書化したフルスタック スケルトン: Dj

ango 5.2 + Strawberry GraphQL + Django Channels 4 + Celery + Next.js 16 + マテリアル UI v9 + Apollo クライアント 4 + AWS EC2 上の Terraform。 ASGI ルーティング、Redis pub/sub を使用した WebSocket サブスクリプション、Celery ワーカーと Beat スケジューラ、RSC およびクライアント コンポーネントを使用した Next.js App Router、MUI v9 テーマ、GraphQL コード生成、Docker Compose 開発/製品分割、フロントエンドの Vercel デプロイメントをカバーします。
13 Stripe Payment Integration Guide 0 クレジット Stripe 支払いのクイックセットアップガイド
1 3 Next.js App Router — NextAuth GitHub OAuth 0 クレジット NextAuth.js v4 GitHub OAuth を Next.js 15 App Router プロジェクトに接続し、サインイン中に GitHub トークンを Django API トークンと交換します。 lib/auth.ts、シン ルート ハンドラー、TypeScript セッション拡張、実装ドキュメントが含まれます。
Bun で管理され、Biome で lint およびフォーマットされ、同じ git リポジトリから別の Vercel プロジェクトにデプロイされた 2 つの独立した Next.js 15 アプリ (パブリック マーケティング サイトと認証されたダッシュボード) の軽量モノリポ パターン。
本格的なエージェントのワークフロー向けに構築
すべての設計上の決定は、エージェントが適切なコンテキストを即座に取得することと、ユーザーが制御を維持することという 2 つのことを最適化します。
プロアクティブな検出 エージェントは、実装タスクを開始する前に自律的にマーケットプレイスを呼び出します。プロンプトは必要ありません。反射のように機能するだけです。
セマンティック + キーワード検索ブースターは、埋め込みと全文検索を使用して関連性によってランク付けされます。エージェントが実際に行おうとしていることに基づいて、適切なブースターが表示されます。
パイプラインのレビュー + 署名 すべてのブースターは、稼働前に自動化された LLM 敵対的レビューと人間による承認を通過します。暗号署名により、挿入時のコンテンツの整合性が保証されます。
オープン ナレッジ コモンズ 専門知識をコミュニティと共有します。あなたのブースターは関連分野で積極的に現れます

すべてのメンバーのワークスペースにわたるセッション — 時間の経過とともに蓄積される知識。
バージョン管理とモデルタグ すべてのブースターは意味的にバージョン管理されます。モデル互換性タグは、作成者がどの LLM に対してテストしたかを示します。再現可能な結果を​​得るためにバージョンに固定します。
あなたの図書館。
あなたのコミュニティ。
あなたの電話です。
プライベート ブースターはあなただけに表示されます。共有することを決定すると、コミュニティ全体がそれらを利用できるようになり、注入するたびにクレジットを獲得できます。
あなたには何年もかけて蓄積された知識があります。 AI Boost を使用すると、ユーザー側での継続的な努力を必要とせずに、構造化されたコンテキスト ブースターとしてパッケージ化できます。
あなたのライブラリはデフォルトで非公開です - 個別のブースターを自分の条件で共有します
構造化された知識を提出します: 規則、パターン、ガイド、注釈付きスニペット
自動化された LLM レビューと人間による承認により、品質が高く不正行為が低く抑えられます
ブースターを自由にバージョン化してください。差分に重点を置いたレビューにより更新が高速化されます
ブースターは、コミュニティ全体の関連するすべてのエージェント セッションで積極的に表示されます
開発者がブースターを注入するたびにクレジットを獲得します (注入ごと、継続的に)
あなたのブースターは、世界中の開発者によって使用される成長を続けるナレッジ コモンズの一部になります
LLM エージェントのブースター層。構造化された知識ペイロードは、重要な瞬間に注入されます。

## Original Extract

Capture what took years to learn. Inject it in seconds. Share it when you're ready. AI Boost is your expertise layer for LLM agents — works with Claude Code, Cursor, and any MCP client.

Hi HN, My number one frustration with LLM agents is that every session starts blank and I have to re-explain my patterns and conventions every time. I often find myself doing things like "look at X repo and copy the patterns used for Y", which requires the repo to be on the local machine. A couple big ones for me personally:
- Auth flows
- Terraform patterns for AWS Rules and skills help to some extent, but they are difficult to synchronize across multiple agents, projects, and machines. Memory is interesting, but too noisy since it is unstructured. So, I built AI Boost to act as a personal library with an MCP server. You tell your agent what to save (text file or public github repo at the moment. I'm working on other options including private github repos.). It packages it as a "booster", indexes it with keywords and embeddings, and next time you start a task that matches, the agent surfaces it. At the moment:
- Boosters are private by default. They are only accessible to your account.
- You can publish to a community marketplace if you want, and earn credits per injection (this isn't completely ready yet, but it's coming soon).
- It's an MCP server, so it should work in Cursor, Claude Code, and any client with MCP support. Link: https://ai-boost.io (and the MCP URL is just https://mcp.ai-boost.io/mcp ) I'd love to know what you think. Especially whether the "auto-suggest before starting a task" behaviour feels useful or intrusive in practice.

AI Boost — Your Expertise Layer for LLM Agents AI
Your expertise layer for LLM agents.
The rules of expertise need to change in the age of AI
Experts are being asked to hand more and more of their knowledge to tools they do not control. We think that trust has to be earned, not assumed. These are not terms of service. They are commitments from one developer to another.
Private boosters are never indexed, sold, surfaced to others, or used to train any model. What you save stays entirely in your account.
You can remove any booster, any version, or your entire account at any point. Deletion is permanent and immediate. There is no archive.
Pricing changes, policy updates, and platform decisions are announced in advance. We will not change the rules without giving you time to act.
We believe the em-dash is a valid — and useful — punctuation mark, and we will continue to use it regardless of what anyone else says.
Your expertise. Captured once. Available everywhere.
Private boosters are yours alone. Capture your expertise from any session. AI Boost embeds it for instant retrieval. Your agent surfaces it the moment context matches.
Tell your agent what you want to save. It reads your current context, proposes metadata, and creates a private booster in seconds.
AI Boost processes the booster for keyword and semantic search so it can be retrieved accurately across all your future sessions.
When you start a task that matches a saved pattern, your agent surfaces the booster before it begins. You approve, it injects.
Save once, available for every future session
Tell your agent what to save. It handles the rest — including reminding you next time.
AI Boost Agent Save our custom auth middleware pattern as a private booster
Let me read the current auth middleware file to extract the pattern.
Got it. Based on the file, here are my suggested booster details:
Name: Custom Auth Middleware Pattern
Description: JWT validation, role-based guards, request context injection
Visibility will be set to private . Confirm to save, or let me know what to change.
Private Price 0 credits (personal use) Confirmed
Ran create_booster in ai-boost
Saved to your library Your Custom Auth Middleware Pattern booster has been saved to your private library under custom-auth-middleware-pattern . It is only visible to you and will be suggested automatically in future sessions.
AI Boost Agent — New session Add protected routes to the new payments service
Before I start, let me check if there are any relevant boosters in your library.
Ran search_boosters in ai-boost
Custom Auth Middleware Pattern
I found your saved auth middleware booster. Would you like me to inject it before I begin?
Your auth middleware conventions are loaded. I will follow your JWT validation and role-based guard patterns throughout this task.
Memory systems capture history, but they cannot curate it. The expertise your agent needs (your conventions, your gotchas, lessons that took years to learn) has nowhere to live between sessions.
Every session starts from scratch Patterns you use every day — your auth flow, your infra conventions, your API design rules — live in human memory. Your agent has no access to them.
Memory is not curation Memory plugins record everything, which means they retrieve everything. The signal-to-noise ratio makes them unreliable for injecting the specific, structured knowledge your agent actually needs.
No way to share structured expertise Domain experts have invaluable knowledge. There is currently no mechanism to package, distribute, and inject that expertise into an LLM at the moment it is needed.
Tell your agent about a convention, pattern, or workflow. AI Boost packages it as a booster and indexes it using keywords and semantic embeddings.
On your next session, your agent queries AI Boost automatically before starting a task. If a relevant booster exists (yours or from the community) it is surfaced for your approval.
You approve. The full context payload flows directly into the agent's context window. Conventions, patterns, examples: available immediately. No copy-paste, no setup.
AI Boost is an MCP server. Add it to your agent client once and it will be available across every project automatically.
Create or edit .cursor/mcp.json in your home directory or project root.
{
"mcpServers": {
"ai-boost": {
"url": "https://mcp.ai-boost.io/mcp",
"transport": "streamable-http"
}
}
} 3 Authenticate with GitHub OAuth
On first use Cursor will open a browser tab. Sign in with GitHub, grant access, and you will be redirected back automatically.
Top up credits and start using boosters
Visit app.ai-boost.io/billing to add credits, then ask Cursor to search for a booster relevant to your task.
Just tell the agent what you want to publish. It handles the rest.
AI Boost Agent I would like to publish this repo as a booster
I'll help you publish this as a booster on the AI Boost Marketplace. Let me start with the first step of the submission flow.
Link to a public GitHub repository
Let me check the GitHub remote URL for this repo.
Ran Check git remote URL in git
The repo is AI-Boost-IO/stripe-payment-integration-guide . Let me fetch the GitHub API to auto-suggest details.
Based on the GitHub repo, here are my suggestions — let me know if you'd like to change any of them:
Name: Stripe Payment Integration Guide
Description: A quick-setup guide for Stripe payments
Version: 1.0.0 (no releases found, defaulting)
Now for the remaining details:
Here's a full summary of the booster submission — please confirm before I proceed:
Does everything look correct? Reply "confirmed" to submit, or let me know what to change.
Let me check the submission tool schema first.
Ran Submit Booster From Repo in ai-boost
The booster was submitted and approved. Here's the summary:
Approved The Stripe Payment Integration Guide booster is now live on the AI Boost Marketplace and discoverable under the slug stripe-payment-integration-guide .
Expert human experience is hard to find in the age of AI
LLMs cannot learn from what they have never seen. Your hard-won patterns are exactly what other developers need. Make your private boosters public when you are ready. They surface automatically in every relevant session across the community.
Developers are increasingly surrounded by AI-generated knowledge. Boosters backed by real human experience are a different category entirely. When you share one, you are contributing something the models genuinely cannot replicate.
Your private booster already exists in your library. You just flip the visibility to public.
It goes through automated LLM adversarial review and human approval to ensure quality.
It becomes searchable across the community. Any developer whose agent matches it can inject it.
You earn AI Boost credits for every injection of your booster — a direct measure of your contribution to the community.
A setup guide for deploying a Dockerised Django app on AWS EC2 using Terraform-managed infrastructure (VPC, EC2, S3, IAM, EIP).
9 14 Django Channels + Celery + Next.js + Terraform/AWS 0 credits Full-stack skeleton documenting the canonical patterns for a monorepo SaaS backend: Django 5.2 + Strawberry GraphQL + Django Channels 4 + Celery + Next.js 16 + Material UI v9 + Apollo Client 4 + Terraform on AWS EC2. Covers ASGI routing, WebSocket subscriptions with Redis pub/sub, Celery workers and Beat scheduler, Next.js App Router with RSC and Client Components, MUI v9 theming, GraphQL codegen, Docker Compose dev/prod split, and Vercel deployment for the frontend.
13 Stripe Payment Integration Guide 0 credits A quick-setup guide for Stripe payments
1 3 Next.js App Router — NextAuth GitHub OAuth 0 credits Wire NextAuth.js v4 GitHub OAuth into a Next.js 15 App Router project and exchange the GitHub token for a Django API token during sign-in. Includes lib/auth.ts, a thin route handler, TypeScript session augmentation, and implementation docs.
A lightweight monorepo pattern for two independent Next.js 15 apps (public marketing site and authenticated dashboard) managed with Bun, linted and formatted with Biome, and deployed to separate Vercel projects from the same git repository.
Built for serious agent workflows
Every design decision optimises for two things: the agent gets the right context instantly, and the user stays in control.
Proactive Discovery Your agent calls the marketplace autonomously before starting any implementation task. No prompting required — it just works, like a reflex.
Semantic + Keyword Search Boosters are ranked by relevance using embeddings and full-text search. The right booster surfaces based on what your agent is actually trying to do.
Review Pipeline + Signing Every booster passes automated LLM adversarial review and human approval before going live. Cryptographic signing ensures content integrity at inject time.
Open Knowledge Commons Share your expertise with the community. Your boosters surface proactively in relevant sessions across every member's workspace — knowledge that compounds over time.
Versioning & Model Tags Every booster is semantically versioned. Model compatibility tags tell you which LLMs the creator tested against. Pin to a version for reproducible results.
Your library.
Your community.
Your call.
Private boosters are only ever visible to you. When you decide to share, they become available to the whole community, and you earn credits for every injection.
You have knowledge that took years to accumulate. AI Boost lets you package it as a structured context booster without any ongoing effort on your part.
Your library is private by default — share individual boosters on your own terms
Submit any structured knowledge: conventions, patterns, guides, annotated snippets
Automated LLM review and human approval keeps quality high and fraud low
Version your boosters freely; diff-focused review makes updates fast
Your booster surfaces proactively in every relevant agent session across the community
Earn credits every time a developer injects your booster — per-injection, ongoing
Your booster becomes part of a growing knowledge commons used by developers worldwide
The booster layer for LLM agents. Structured knowledge payloads, injected at the moment they matter.
