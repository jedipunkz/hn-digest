---
source: "https://codesummary.io/"
hn_url: "https://news.ycombinator.com/item?id=48589234"
title: "Show HN: A knowledge engine that keeps up with AI driven code"
article_title: "CodeSummary"
author: "jerrodcodes"
captured_at: "2026-06-18T18:54:10Z"
capture_tool: "hn-digest"
hn_id: 48589234
score: 1
comments: 0
posted_at: "2026-06-18T18:12:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A knowledge engine that keeps up with AI driven code

- HN: [48589234](https://news.ycombinator.com/item?id=48589234)
- Source: [codesummary.io](https://codesummary.io/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T18:12:11Z

## Translation

タイトル: Show HN: AI 駆動のコードに対応するナレッジ エンジン
記事のタイトル: コードの概要
説明: コード サマリーは、リポジトリを生きたドキュメントに変えます。AI によって作成され、ユーザーによってレビューされ、独自のドメインでホストされ、MCP 経由で AI エージェントに提供されます。
HN テキスト: 私たちは、複雑なリポジトリ間のコードベースに対応するエンドツーエンドのナレッジ システムを構築しました。当社のインテリジェントなコンテキスト マッピング システムは、チームが内部/外部でホストしたり、MCP 経由でクエリしたりできる包括的なナレッジ ベースを提供します。

記事本文:
Codesummary Codesummary 製品価格設定ブログ サインイン 無料で始める エンジニアリング チームのコンテキスト レイヤー
AI によってコードベースがブラック ボックスにならないようにしてください。
コードベース全体の生きた真実の情報源であり、常に完全で常に最新であり、コードの移動に合わせて維持されます。チームとエージェントがスレッドを失うことはありません。
14 日間の無料トライアル。クレジットカードはありません。どのように機能するか見てみましょう
Dashboard.codesummary.io チームがすでに使用しているツールと連携します
知識はすでにコードに組み込まれています。それを取り出すには費用がかかります。
それで、それを出しました。自動的に。
私たちは Codesummary を、GitHub で最も人気のあるプロジェクトの 1 つである OpenClaw に注目しました。350,000 を超えるスター、50,000 を超えるコミット、そして毎日移動するギガバイトの TypeScript です。その 1 つの入力から、人間が触れたものはなく、すべての概念が理解され、すべてのページが書かれた、完全な相互リンクされたドキュメント サイトが構築されました。ソースコードが入力され、生きているドキュメントが出力されます。
GitHub Codesummary の OpenClaw 0 人による編集 openclaw.codesummary.io live 公開ドキュメント サイト これが実際の出力です。ソースから生成され、独自のドメインでホストされている完全なドキュメント サイトです。ライブで参照するか、自分のリポジトリで Codesummary を指定します。
ライブドキュメントを参照してください。
リポジトリから生きたコンテキストへ。
acme / フロントエンド ⎇ main · 2 分前にプッシュ acme / api ⎇ main · 1 時間前にプッシュ acme / billing-service ⎇ main · 1 日前にプッシュ acme / infra ⎇ main · 3 日前にプッシュ Codesummary 1 つのワークスペース読み取りがメインからプッシュ… ページに整理… 公開 → 2 つのエンドポイント mcp · /product mcp · /platform Claude Code · mina /product ▸ ask("返金はどのように機能しますか?") 請求サービスを通じて返金をポストし、台帳に同期します… [製品 v14] カーソル · 開発 /プラットフォーム ▸ orient() 6 つのサービス · エントリ ポイント · 規約 · 同期 2 分

go [プラットフォーム v9] Humans · docs docs.acme.com 01 ストリームをプッシュします
GitHub アプリをインストールすると、メインへのすべてのプッシュがワークスペースに降り注ぎます。プライベート リポジトリは非公開のままです。
乱雑な情報が読み取られ、レビューされたページに整理され、個別のサイトに合成され、それぞれが独自の MCP エンドポイントに公開されます。
エージェントはエンドポイントに対して ask() と orient() を呼び出し、引用された最新の回答を取得します。同じナレッジがドメイン上でドキュメントとして公開されます。
スタック全体に対する 1 つの信頼できる情報源。エージェントはただ尋ねるだけです。
リンクするすべてのリポジトリ、フロントエンド、バックエンド、およびその間のサービスが 1 つのモデル コンテキスト プロトコル エンドポイントになります。兄弟リポジトリのクローンを作成してざっと読んで、その動作を学習する代わりに、エージェントは Codesummary に問い合わせて、1 回の呼び出しでレビュー済みの回答を取得します。
単一のチェックアウトでは提供できないクロスリポジトリの回答
最後に git pull を実行した人ではなく、メインをトラックします。
1 つのワークスペース、1 つのエンドポイントの下にあるすべてのリポジトリ
OAuth で保護 · すべての主要な MCP クライアントで動作します
展開 エンジニアリング リード向け · 新規
基準を一度書きます。すべてのエージェントが彼らに従います。
スタイル ガイドは、コーディング標準、アーキテクチャ パターン、チームのやり方など、リーダーが管理する独自の MCP エンドポイントです。すべてのエンジニアのエージェントはコードを記述する前に PR を取得するため、誰が (または何を) 書いたかに関係なく、PR に含まれる内容はすでにコードベースと同じように見えます。
すべてのエージェントに提供されるアーキテクチャ パターンと規約
一度更新すれば、チーム全体が次のプロンプトに向けて調整できます
コードレビューで標準を再訴訟するのをやめる
@lead によってキュレーションされた style-guide.md § サービスはイベントを介して通信し、DB を直接読み取ることはありません。 § エラーは、problem+json を返します。裸の500はありません。 § React: デフォルトのサーバー コンポーネント。インタラクティブ時のみクライアント。クロードコード・ミナ整列カーソル・開発整列クロードコード・サーシャ整列ほぼリアルタイム
いつでも常に最新の状態

あなたのチェックアウトはそうではありません。
ローカル リポジトリのドリフト: 誰かのフロントエンドがメインより 3 週間遅れており、誰かが新しいサービスをまったくプルせず、それらのチェックアウトを読み取るすべてのエージェントがドリフトを継承します。 Codesummary は、誰のディスクに何があるかは気にしません。メインにマージするたびにドキュメントが再生成されるため、チームのすべてのエージェントは実際に出荷されたものから回答します。
回答は最後のプッシュ時点のメインを反映しており、誰かのローカル状態は反映されていません
マージごとに完全および増分再生成
編集内容がサイレントに上書きされることはありません。変更内容はレビュー中です。
ドメインとブランドを拡大する
そして、人間は実際に使用するドキュメント サイトを手に入れることができます。
エージェントがクエリするのと同じレビュー済みの知識が、洗練されたドキュメント サイトとして提供されます。公開製品ドキュメント、社内チーム ハンドブック、またはその両方: 自動 SSL を備えたカスタム ドメイン、ロゴとカラー、サインイン後にチームの知識を保持するプライベート モード。
自動証明書を使用したカスタム ドメイン
チームに対して公開、非公開、または非公開
仕様が検出された場合の対話型 API リファレンス
展開 本格的なチーム向けに構築
セキュリティレビューを制御できます。
新しいサイトはチームのみで開始されます。あなたが決定するまでは何も公開されません。
あなたがタッチしたページに対する AI の変更はすべてレビューを受けます。出荷前に承認、拒否、または編集します。
OAuth 付与とアクセス トークンは、単一のサイトまたはリポジトリにスコープされます。いつでもそれらを取り消してください。
公開されたバージョンはすべて保存されます。ワンクリックで以前の状態にロールバックします。
シンプルな定額料金設定。計量はありません。
Pro と Team は 14 日間の無料トライアルから始められ、カードは必要ありません。スケールとエンタープライズはチームに合わせてカスタマイズされます。
開発者および小規模チーム向け。
多くのリポジトリにまたがって出荷するチーム向け。
コンプライアンス主導の組織向け。
すべてのエージェントが同じ認識を持つようにします。
リポジトリを接続し、スタイル ガイドを設定し、チーム全体でハミングします。

回答者もエージェントも同様に、1 つのソースから応答を開始します。
コードベースの製品 Codesummary Living ドキュメントを参照してください。 AI によって作成され、ユーザーによってレビューされ、ドメイン上でホストされ、すべてのエージェントが使用できるようになります。

## Original Extract

Code Summary turns your repositories into living documentation: written by AI, reviewed by you, hosted on your own domain, and served to AI agents over MCP.

We built an end to end knowledge system that keeps up with complex cross-repo codebases. Our intelligent context mapping system provides comprehensive knowledge bases that your team can host internally/externally or query via MCP.

CodeSummary CodeSummary Product Pricing Blog Sign in Start free The context layer for engineering teams
Don’t let AI turn your codebase into a black box.
A living source of truth for your entire codebase, always complete, always current, maintaining itself as your code moves. Your team and your agents never lose the thread.
14-day free trial. No credit card. See how it works
dashboard.codesummary.io Works with the tools your team already uses
The knowledge is already in the code. Getting it out is what costs you.
So we got it out. Automatically.
We pointed CodeSummary at OpenClaw, one of GitHub’s most-starred projects: over 350,000 stars, more than 50,000 commits, and gigabytes of TypeScript that move every day. From that one input it built a complete, cross-linked documentation site, every concept understood, every page written, not a single one touched by a human. Source code in, living docs out.
OpenClaw on GitHub CodeSummary 0 human edits openclaw.codesummary.io live Published docs site That is the actual output: a complete documentation site, generated from the source and hosted on its own domain. Browse it live, or point CodeSummary at your own repo.
See the live docs How it works
From repositories to living context.
acme / frontend ⎇ main · pushed 2m ago acme / api ⎇ main · pushed 1h ago acme / billing-service ⎇ main · pushed 1d ago acme / infra ⎇ main · pushed 3d ago CodeSummary one workspace reading pushes from main… organizing into pages… publishing → 2 endpoints mcp · /product mcp · /platform Claude Code · Mina /product ▸ ask("how do refunds work?") Refunds post through billing-service, then sync to the ledger… [product v14] Cursor · Dev /platform ▸ orient() 6 services · entry points · conventions · synced 2m ago [platform v9] Humans · docs docs.acme.com 01 Pushes stream in
Install the GitHub App and every push to main rains into the workspace. Private repos stay private.
The clutter is read, organized into reviewed pages, and synthesized into separate sites, each published to its own MCP endpoint.
Agents call ask() and orient() against their endpoint and get cited, current answers. The same knowledge publishes as docs on your domain.
One source of truth for your whole stack. Agents just ask.
Every repo you link, frontend, backend, and the services between, becomes one Model Context Protocol endpoint. Instead of cloning and skimming a sibling repo to learn what it does, the agent asks CodeSummary and gets the reviewed answer in one call.
Cross-repo answers no single checkout can provide
Answers track main, not whoever last ran git pull
All your repos under one workspace, one endpoint
OAuth-secured · works with every major MCP client
Expand For engineering leads · New
Write the standards once. Every agent follows them.
The style guide is its own MCP endpoint your lead curates: coding standards, architecture patterns, the way your team does things. Every engineer's agent pulls it before writing code, so what lands in the PR already looks like your codebase, no matter who (or what) wrote it.
Architecture patterns and conventions, served to every agent
Update once and the whole team is aligned on the next prompt
Stop re-litigating standards in code review
style-guide.md curated by @lead § Services communicate over events, never direct DB reads. § Errors return problem+json. No bare 500s. § React: server components by default. Client only when interactive. Claude Code · Mina aligned Cursor · Dev aligned Claude Code · Sasha aligned Near real-time
Always current, even when your checkouts aren't.
Local repos drift: someone's frontend is three weeks behind main, someone never pulled the new service at all, and every agent reading those checkouts inherits the drift. CodeSummary doesn't care what's on anyone's disk. Every merge to main regenerates the docs, so every agent on the team answers from what actually shipped.
Answers reflect main as of the last push, not anyone's local state
Full and incremental regeneration on every merge
Your edits are never silently overwritten: changes wait in review
Expand Your domain, your brand
And your humans get a docs site they'll actually use.
The same reviewed knowledge your agents query ships as a polished docs site. Public product docs, internal team handbooks, or both: custom domains with automatic SSL, your logo and colors, and a private mode that keeps team knowledge behind sign-in.
Custom domains with automatic certificates
Public, unlisted, or private to your team
Interactive API reference when a spec is detected
Expand Built for serious teams
Control you can take to your security review.
New sites start team-only. Nothing becomes public until you decide it should.
Every AI change to a page you have touched goes through review. Accept, reject, or edit before anything ships.
OAuth grants and access tokens are scoped to a single site or repo. Revoke any of them at any time.
Every published version is preserved. Roll back to any prior state in one click.
Simple, flat pricing. No metering.
Pro and Team start with a 14-day free trial, no card required. Scale and Enterprise are tailored to your team.
For developers and small teams.
For teams shipping across many repos.
For compliance-driven organizations.
Get every agent on the same page.
Connect your repos, set your style guide, and the whole team, humans and agents alike, starts answering from one source.
Explore the product CodeSummary Living documentation for your codebase. Written by AI, reviewed by you, hosted on your domain, and ready for every agent.
