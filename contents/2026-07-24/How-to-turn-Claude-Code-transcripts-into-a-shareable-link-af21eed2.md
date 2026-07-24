---
source: "https://lore.link/blog/how-we-built-share"
hn_url: "https://news.ycombinator.com/item?id=49042509"
title: "How to turn Claude Code transcripts into a shareable link"
article_title: "How we built /share | Lore"
author: "feifan"
captured_at: "2026-07-24T22:55:02Z"
capture_tool: "hn-digest"
hn_id: 49042509
score: 1
comments: 0
posted_at: "2026-07-24T22:42:17Z"
tags:
  - hacker-news
  - translated
---

# How to turn Claude Code transcripts into a shareable link

- HN: [49042509](https://news.ycombinator.com/item?id=49042509)
- Source: [lore.link](https://lore.link/blog/how-we-built-share)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T22:42:17Z

## Translation

タイトル: クロード コードのトランスクリプトを共有可能なリンクに変える方法
記事のタイトル: /share をどのように構築したか |伝承
説明: 今週 /share を開始しました。このようにして構築しました。

記事本文:
/share をどのように構築したか |伝承
Lore 価格設定ドキュメント サインイン 無料で始める ブログ / すべてのエントリ 構築方法 /共有
今週 /share を開始しました。このようにして構築しました。
Lore は、AI プロンプトとセッションのためのソーシャル ネットワークとして始まりました。おそらく、これらは 2026 年版のオープンソース Github リポジトリに相当するでしょう。 /share コマンドがステップ 1 でした。
それ以来、私たちはソーシャル ネットワークの考え方から離れて進化してきましたが、/share は私たちと初期のユーザーにとって依然として有用です。
PR を作成するすべての Claude Code セッションでこれを実行し、コンテキストのために対応する PR に Lore リンクを追加します。
ユーザーはこれを使用して仕事を引き継ぎます。ある人が調査を開始し、それを他の人 (または別のエージェント) に渡して終了します。
ユーザーの中には、ユーザーに配布される製品として独自のスキル/プラグインを構築し、製品がユーザーにとってどの程度うまく機能しているかを観察するために /share (または SDK などの関連コンポーネント) を使用している人もいます。
この一見単純なコマンドには多くの詳細が含まれています。ディスク上の適切なファイルの検索、繰り返される共有の調整、および認証 (もう 1 つの 4 文字の単語) などの重要な作業です。共有リンクの自動コピー、共有スレッドのタイトルの指定、共有リンク内で強調表示するブロックの指定など、生活の質を向上させる機能があります。
その構築方法と、その過程で学んだ教訓について詳しく見ていきましょう。
/share は 3 つの主要なレイヤーで構成されています。
プラグインにラップされたスキル。ユーザーが実際に呼び出すコマンドです (例: Claude Code の /share または Codex の $share 経由)。
オンデバイス機能 (セッション ファイルの読み取り、出力 URL のクリップボードへのコピー) を実装するローカル MCP サーバー。
基盤となる API を調整するリモート MCP サーバー。
/share はスキルです。ソースは Github にあります。
この sk の最初のバージョン

エージェントに独自のセッション ファイルを読み取って API にアップロードするように依頼しただけです。これは実際のセッションではすぐに破綻しました。セッション ファイルを読み取ると、エージェントのコンテキスト使用量がすぐに 2 倍になり、特にエージェントが結果のセッションを圧縮しようとしたときに問題が発生しました。
この問題は、コンピューターのローカル ファイル システムから特定のトランスクリプト ファイルを読み取ってアップロードするスクリプトを追加することで修正しました。スキルはこのスクリプトを呼び出すだけです。このスクリプトは CLI に発展し、その後ローカル MCP サーバーに発展しました (これについては以下で詳しく説明します)。
スクリプトに依存するスキルができたので、両方をインストールする簡単な方法が必要でした。私たちはそれらをプラグインにバンドルしました。 Claude と Codex は、プラグインをインストールするための組み込みの方法を提供します。このプラグインはスキルの名前空間も設定します (つまり、実際には /lore:share と $lore:share になります)。これは初期のユーザーの 1 人がリクエストしたものです。
私たちのプラグインには 2 つの追加スキルが付属しています。
/read : 「スレッド th_123 で Matt が行ったことを見て、それを私のシステムに適用してください」のようなプロンプトを表示できます。
/fork : 既存のスレッドのインテントベースの概要を生成します。情報を無差別に破棄するスレッドを事前に一般的に要約する代わりに、多くの Lore 機能は特定の目的のためにコンテンツを要約します (一部は組み込まれており、この機能のようにユーザーが提供するものもあります)。
Claude Cowork のサポートを追加しようとしたときに、ローカル スクリプト/CLI が壁にぶつかりました。Cowork セッションはサンドボックス (最初はコンピュータ上にありましたが、現在はデフォルトでリモートにあります) に突入しました。この環境のスクリプト/CLI はローカル ファイル システムにアクセスできないため、トランスクリプト ファイルを読み取ることができません。
Cowork は、Claude デスクトップ アプリをローカル コンピューターへのブリッジとして使用します。
プラグインとデスクトップ拡張機能がバンドルされたローカル MCP サーバーは、

実行する他のプログラムと同じ権限 ( Source )
トランスクリプト読み取りコードを share_session という名前のローカル (stdio) MCP ツールに変換し、それをプラグインにバンドルしました。
share_session は、ディスク上でセッションのトランスクリプト ファイルを見つけ、それをサーバーにアップロードし、URL を取得して、ローカル コピー コマンド (macOS の pbcopy など) を呼び出して自動的にクリップボードに保存します。
重要なのは、ローカル MCP サーバーも認証を担当することです。必要に応じてログイン フローをトリガーし、それを使用してトークンを取得し、ディスクに保存します。当初、サーバーから返された OAuth アクセスとリフレッシュ トークンを直接保存していましたが、認証トークンのリフレッシュは信頼性が低く、場合によっては同期ログイン フローでユーザーを中断する必要がありました。幸いなことに、ユーザーの 1 人のリクエストに応じて、SDK をサポートするアップロード API キー (UAK) システムを別途構築していました。最初のログイン時にプログラムで有効期間の長い UAK を生成し、将来のセッションのアップロードのためにその値を保持するようになりました。
技術的にはホスト型 MCP サーバーは必要ありませんでしたが、すでに存在しており、基礎となる API を抽象化しています。たとえば、リモート share_session ツールは 11 のステップをカバーします。
S3 への署名付きストレージ URL を取得します。
トランスクリプト バイトをその署名付き URL に PUT します。
アップロード セッションを完了します (アップロードされたコンテンツに対するさらなるサーバー側の処理のトリガーとなります)。
タイトルと可視性の動作を適用します。
必要に応じて、バックグラウンド ジョブがトランスクリプトの解析を完了し、自然言語のハイライトを解決するまで待ちます。
プラグインのアクティベーション分析を記録します。
正規の絶対 URL を返します。
サーバー側でそれを抽象化することで、プラグイン バンドルが小さくなり、進化しやすくなります。
lore-plugin リポジトリは Github にあります。これには次のものが含まれます。
バンドルされたスキルの内容。
の

ローカル MCP サーバーのソースおよびビルドされたアーティファクト (Bun バイナリ)。そして
Claude および Codex プラグインのマニフェスト。 Amp 用のプラグインもありますが、これは実験的なものであり、チーム外ではまだサポートされていません。
このリポジトリは、すべての開発を行うプライベート モノリポジトリからミラーリングされたサブツリーです。このモノリポジトリ設定により、製品全体にわたって横断的な変更を迅速に行うことができます。これについては後の投稿で詳しく説明します。
スレッド共有は、私たちが Lore で構築しているものの基礎部分ですが、今後さらに多くの機能が追加されます。
私たちは、エージェントと AI ネイティブのワークフローが新しいプリミティブを生み出していると信じています。 AI が登場する前は、作業は会議、ファイル、ランブック/SOP などの構成要素を中心に展開していました。 AI ネイティブのワークフローでは、スレッド、スキル、アーティファクトが導入されます。これらの中には、AI 以前の構造に類似したものがあるものもあります。そうでない人もいます（または少なくとも、完全にはそうではありません）。どちらの場合も、AI ネイティブのチームや個人は、これらの新しいプリミティブを操作するための適切なツールを必要とします。これがどのようなものになるのかはまだ正確にはわかりませんが (最前線では能力やワークフローが急速に変化しています)、私たちは、自分たち自身のため、そしてナレッジ ワーカーのために、可能なことの最前線で構築し、解明する予定です。
これに興味がある場合は、Lore を試してみて、ご意見をお聞かせください。
チームの AI 作業は複雑化するはずです。
Lore は、チームが AI を使用して作成したすべてのセッション、すべてのスキル、すべてのプロセスを収集し、検索可能で共有可能なナレッジに変換します。自動的にキャプチャされます。
無料で始められます・クレジットカードは不要です

## Original Extract

We launched /share this week. This is how we built it.

How we built /share | Lore
Lore Pricing Docs Sign in Start free Blog / all entries How we built /share
We launched /share this week. This is how we built it.
Lore started as a social network for AI prompts and sessions: perhaps these would be the 2026 equivalent of open-source Github repos. The /share command was step 1.
We've since evolved away from the social network idea, but /share remains useful for us and our early users:
We run it on every Claude Code session that creates a PR and add the Lore link to the corresponding PR for context.
Our users use it to hand off work — one person starts an investigation, and then passes it to someone else (or a different agent) to finish.
Some of our users are building their own skills/plugins as products distributed to their users, and use /share (or related components like our SDK ) for observability on how well their products are working for their users.
A lot of detail goes into that seemingly-simple command: table-stakes stuff like finding the right file on disk, reconciling repeated shares, and auth (the other four-letter word); quality-of-life stuff like automatically copying the shared link, letting you specify a title for the shared thread, and letting you specify blocks that you want to highlight in the shared link.
Let's dive into how it's built and the lessons we learned along the way.
/share is made up of three main layers:
A skill, wrapped in a plugin, which is the command that users actually invoke (e.g. via /share in Claude Code or $share in Codex).
A local MCP server that implements on-device functionality (reading session files; copying the output URL to clipboard).
A remote MCP server that orchestrates our underlying API.
/share is a skill; its source is on Github .
The first version of this skill just asked the agent to read its own session file and upload it to our API. This broke down quickly on real-world sessions: reading the session file immediately doubled the context usage of the agent, and that caused problems, especially when the agent tried to compact the resulting session.
We fixed this by adding a script that read a given transcript file from our computers' local filesystem and uploaded it; the skill would simply invoke this script. That script evolved into our CLI , and later into our local MCP server (more on this below).
Now we had a skill which depended on a script, and we wanted an easy way to install them both. We bundled them into a plugin ; Claude and Codex provide built-in ways to install plugins. The plugin also namespaces our skills (so it's actually /lore:share and $lore:share ), which was something that one of our early users requested.
Our plugin now comes with two additional skills:
/read , which lets you prompt something like "look at what Matt did in thread th_123 and apply that on my system".
/fork , which generates an intent-based summary of an existing thread. Instead of generically summarizing threads up-front, which indiscriminately discard information, many Lore features summarize content for specific purposes (some of which are baked in; some of which, like this one, are user-supplied).
Local script/CLI hit a wall when we went to add support for Claude Cowork: Cowork sessions ran into a sandbox (originally on your computer, now remotely by default ). A script/CLI in this environment doesn't have access to your local filesystem, so it can't read transcript files.
Cowork uses the Claude Desktop app as a bridge to your local computer:
Local MCP servers bundled with plugins and desktop extensions run on your computer with the same permissions as any other program you run ( Source )
We converted our transcript-reading code into a local (stdio) MCP tool named share_session , and bundled that into our plugin .
share_session finds the transcript file for your session on disk, uploads it to our server, gets back a URL, and then invokes a local copy command (e.g. pbcopy on macOS) to automatically put it onto your clipboard.
Importantly, the local MCP server is also responsible for authentication. It'll trigger a login flow when needed, and use that to retrieve and save a token to disk. Originally, we directly stored the OAuth access and refresh token we got back from our server, but refreshing auth tokens was unreliable and would occasionally require interrupting the user with a synchronous login flow. Fortunately, we had separately built an Upload API Key (UAK) system to support our SDK at the request of one of our users; we now programmatically generate a long-lived UAK upon initial login, and persist that value for future session uploads.
We technically didn't need the hosted MCP server … but we already had it, and it abstracts over our underlying API. For example, the remote share_session tool covers 11 steps:
Obtain a presigned storage URL to S3.
PUT the transcript bytes to that presigned URL.
Complete the upload session (which is the trigger for further server-side processing on the uploaded content).
Apply title and visibility behavior.
Optionally wait for a background job to finish parsing the transcript and resolve a natural-language highlight.
Record plugin activation analytics.
Return a canonical absolute URL.
Abstracting that away server-side makes our plugin bundle smaller and easier to evolve.
The lore-plugin repo is on Github . It contains:
the contents of our bundled skills ;
the source and built artifact of the local MCP server (a Bun binary); and
manifests for a Claude and Codex plugin. We also have an plugin for Amp , although this is experimental and not yet supported outside of our team.
This repo is a subtree mirrored out of our private monorepo, in which we do all our development. This monorepo setup allows us to rapidly make cross-cutting changes across our whole product; more on this on a later post.
Thread sharing is a foundational piece of what we're building at Lore, but there's a lot more to come.
We believe that agents and AI-native workflows are resulting in new primitives. Before AI, work revolved around constructs like meetings, files, and runbooks/SOPs; AI-native workflows introduce threads, skills, and artifacts. Some of these have analogs in pre-AI constructs; others do not (or at least, not entirely). In both cases, AI-native teams and individuals will need the right tools to work with these new primitives. We don't yet know exactly what this will look like (capabilities and workflows are changing very quickly at the frontier), but we plan on building — for ourselves, and for knowledge workers at the frontier of what's possible — and finding out.
If this speaks to you, give Lore a try and let us know what you think.
Your team's AI work should compound.
Lore gathers every session, every skill, and every process your team creates with AI, and turns them into searchable, shareable knowledge. Captured automatically.
Free to start · No credit card
