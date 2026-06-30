---
source: "https://www.zandrey.com/blog/i-like-claude-desktop-so-i-created-my-own"
hn_url: "https://news.ycombinator.com/item?id=48740294"
title: "I like Claude Desktop, so I created my own"
article_title: "I like Claude Desktop, so I created my own | Andrey Zagoruiko"
author: "rats"
captured_at: "2026-06-30T23:26:47Z"
capture_tool: "hn-digest"
hn_id: 48740294
score: 1
comments: 0
posted_at: "2026-06-30T22:52:48Z"
tags:
  - hacker-news
  - translated
---

# I like Claude Desktop, so I created my own

- HN: [48740294](https://news.ycombinator.com/item?id=48740294)
- Source: [www.zandrey.com](https://www.zandrey.com/blog/i-like-claude-desktop-so-i-created-my-own)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T22:52:48Z

## Translation

タイトル: Claude Desktopが気に入ったので自作してみました
記事タイトル：Claude Desktopが気に入ったので自作してみました |アンドレイ・ザゴルイコ
説明: Tauri、GLM-5.2、Claude Agent SDK、および Apple コンテナ マシンを使用して独自の Claude デスクトップを構築する

記事本文:
私は Claude Desktop が好きなので、独自のデスクトップを作成しました。 Andrey Zagoruiko Andrey Zagoruiko 8 メモ 約 12/1/24 15 年以上の経験を持つ AI プロダクト マネージャー ブログ クロード デスクトップが好きなので、独自のクロード デスクトップを作成しました 26 年 6 月 30 日 4 分で読む Tauri、GLM-5.2、Claude Agent SDK、および Apple コンテナ マシンを使用して独自のクロード デスクトップを構築する なぜ人魚はダメなのか、そしてそれについて私が何をしたか 26 年 3 月 29 日 13 分で読む AI本当に図が苦手なので修正しました Trying Too Hard by Dean Williams 2015/2/14 4 min read ディーン・ウィリアムズ著 Trying Too Hard からの逐語的な引用。 Cursor が優れた企業である理由 24 年 11 月 15 日 2 分読み取り Cursor のビジネス モデル、価格設定戦略、製品アプローチの分析 ステータス ゲーム 22 年 6 月 18 日 5 分読み取り ステータス力学が人間の相互作用をどのように制御するかに関する Keith Johnstone の「Impro」からのメモ Josh Elman が Alex Zhu にインタビュー 22 年 6 月 5 日 2 min read ソーシャル プロダクト構築に関する Musical.ly 共同創設者からの 5 つの主要な教訓 製品アイデア22/5/19 3 分で読める 製品思考のフレームワークとメンタル モデルの生きたコレクション Claude Desktop が気に入ったので、独自のものを作成しました
私は Claude デスクトップ アプリケーションが気に入っており、仕事でも個人でもよく使っています。
私は何かを構築するのも大好きで、最近リリースされた GLM-5.2 を機に、独自の Claude デスクトップを構築できないか試してみることにしました。それを「レコワーク」と名付けました。
以前、私は Tauri (デスクトップ アプリケーション用のクロスプラットフォーム Rust フレームワーク) と OrbStack を使用した経験があったので、このプロジェクトではそれらを使用することにしました。
Claude Desktop に慣れていない人のために説明すると、これは、チャット、コード、およびその中間 (Cowork) を実行できるエージェント ハーネスを備えたデスクトップ アプリケーションです。
私のバージョンも同じ考え方に従っています。つまり、目標を与えると、計画が立てられ、ツールが呼び出され、ジョブが完了するまでステップが連鎖していきます。
ハッカーニュースを取得しました

トップページの編集、トップ記事の並べ替え、マークダウンへの保存、Twitter スレッドの作成など、すべてを 1 つのプロンプトから行うことができます。
(OpenClaw とは異なり) デスクトップ エージェントを安全にするには、エージェントを閉じた空間、つまりコンテナ内に制限する必要があります。 LLM がマシン上でランダムなコマンドを実行することを許可したくありません。VM/コンテナーはまさにこれを支援します。
2 番目の層もあります。書き込み/実行ツールの呼び出しはすべて、実際に何かが実行される前に承認者 (私の canUseTool ) を経由するため、ファイルシステムにアクセスする前にコマンドを確認して拒否できます。
Apple コンテナ マシンへの移行
最初のいくつかのバージョンを試した後、Apple がコンテナ マシンを導入したことを WWDC'26 で最近見たことを思い出し、それを試してみることにしました。さらに、私は最近 macOS Golden Gate にアップグレードしたばかりだったため (これが以前のバージョンで利用できるかどうかはわかりませんでした)。
Apple コンテナ マシンに移行すると、追加の要件 (OrbStack をインストールする必要がなくなります) とディスク容量がいくつか節約されました。適切なパフォーマンス テストは行っていませんが、OrbStack のコンテナと Apple のコンテナのパフォーマンスの違いには気づきませんでした。
私が驚いたのは、GLM-5.2 が Claude Agent SDK とうまく連携することです。これに任意の (基本的な) タスクを投げることができます。それは... 正常に機能します。
さらに正常に回復します。ここでは Python を試しましたが、利用できないことがわかり、Node の BigInt に切り替えて 10,000 行のフィボナッチ数列を一拍も逃すことなく検証しました。
どれくらいのコストがかかるのか知りたかったので、詳細な統計 (トークンの入出力、キャッシュなど) を追加しました。そして、モデルの機能をテストした数百のスレッドで ~1 ドルの支出が発生したことは興味深いことです。
貧乏な人には、GLM-5.2 が最適です。私は貧乏ではないので、このデスクトップ アプリケーションを構築するために Opus 4.8 の Claude Code を使用しましたが、

中国モデルがいかに効果的であるかは依然として興味深い。
ツールと今後の展開
Claude Agent SDK には、多くのことを実行できる素晴らしいツール セットが付属しており、必要に応じて、このリストは組み込みツール (つまり、承認者 canUseTool ) または外部 MCP を使用して拡張できます。まだMCPに到達していない。
多くの人が、大手研究所のトークンは高価になり、場合によっては持続不可能になっていると言い始めています。また、将来はオープンソース モデルにあり、Fireworks や Baseten などの推論プロバイダーの大成功がこの点を証明しているだけだと言う人もいます。
なぜここで Baseten を選んだのかと疑問に思われるかもしれませんが、答えは非常に簡単です。 Baseten で働く Philip Kiely の『Inference Engineering』という本を購入し、本を読みながら Baseten のアカウントをセットアップしたところ、その素晴らしさに驚きました。
いずれにせよ、最新の AI を活用したツールは完全に狂っています。あれやこれやを少しでも知っていれば、自分が何をしているのか知っているか、少なくともそれを理解しようとする好奇心と意欲があれば、誰でも何かを構築し、そこから大金を稼ぐことができます。

## Original Extract

Building my own Claude Desktop with Tauri, GLM-5.2, the Claude Agent SDK, and Apple container machines

I like Claude Desktop, so I created my own | Andrey Zagoruiko Andrey Zagoruiko 8 notes About 12/1/24 AI Product Manager with 15+ years of experience Blog I like Claude Desktop, so I created my own 6/30/26 4 min read Building my own Claude Desktop with Tauri, GLM-5.2, the Claude Agent SDK, and Apple container machines Why Mermaid Sucks and What I Did About It 3/29/26 13 min read AI genuinely sucks at diagrams and I fixed it Trying Too Hard by Dean Williams 2/14/25 4 min read Verbatim quotes from Trying Too Hard by Dean Williams. Why Cursor is a great company 11/15/24 2 min read Analysis of Cursor's business model, pricing strategy, and product approach Status games 6/18/22 5 min read Notes from Keith Johnstone's 'Impro' on how status dynamics govern human interactions Josh Elman interviews Alex Zhu 6/5/22 2 min read Five major lessons from the Musical.ly cofounder on building social products Product ideas 5/19/22 3 min read A living collection of product thinking frameworks and mental models I like Claude Desktop, so I created my own
I like the Claude Desktop application and use it a lot, both for my work and in a personal capacity.
I also love to build stuff, and with the recent release of GLM-5.2 I decided to see if I could build my own Claude Desktop. I called it Recowork .
Previously I had some experience with Tauri (a cross-platform Rust framework for desktop applications) and OrbStack , so I decided to use them in this project.
For those who are not familiar with Claude Desktop - it is a desktop application with an agentic harness that can do chat, code, and something in between (Cowork).
My version follows the same idea: you give it a goal and it plans, calls tools, and chains the steps together until the job is done.
Here it fetched the Hacker News frontpage, sorted the top stories, saved them to markdown, and wrote a Twitter thread - all from a single prompt.
To make your desktop agent secure (unlike OpenClaw) you need to limit agents within a closed space, i.e. a container. You don't want to allow an LLM to execute random commands on your machine, and VMs/containers help with this exact thing.
There's a second layer too: every write/exec tool call goes through an approver (my canUseTool ) before anything actually runs, so I can see and deny the command before it touches the filesystem.
Moving to Apple container machines
After the first several versions, I recalled that I recently saw from WWDC'26 that Apple introduced container machines , and I decided to try it - moreover because I had recently upgraded to macOS Golden Gate (not sure if this thing is available on the previous versions).
Moving to Apple container machines saved some extra requirements (people wouldn't need to install OrbStack) and some disk space. I didn't notice any performance differences between OrbStack's and Apple's containers, although I didn't do proper performance testing.
What surprised me is how well GLM-5.2 works with the Claude Agent SDK. You can throw any (basic) task at this thing and it... just works.
It even recovers gracefully - here it tried Python, found it wasn't available, and switched to Node's BigInt to validate a 10,000-line Fibonacci sequence without missing a beat.
I wanted to see how much it costs, so I added detailed stats - tokens in/out, cached, etc. - and it is fascinating to see that the hundreds of threads where I tested the model's capabilities resulted in ~$1 of spend.
If you are poor, GLM-5.2 is a great choice for you. I am not poor, and I used Claude Code with Opus 4.8 to build this desktop application, but it is still fascinating how effective Chinese models are.
Tools, and where this goes next
The Claude Agent SDK comes with an awesome set of tools that can do a lot, and if need be, this list can be extended either with built-in tools (i.e. my approver canUseTool ) or with external MCPs. I haven't gotten to MCPs yet.
Many people are starting to say that tokens from the big labs are becoming more expensive and in some cases unsustainable. Others are saying that the future is in open source models, and that the huge success of inference providers such as Fireworks, Baseten, and others only proves this point.
If you wonder why I chose Baseten here, the answer is very simple. I bought the book Inference Engineering by Philip Kiely, who works at Baseten, and while reading the book I set up an account with Baseten and was surprised how good it is.
In any case, modern AI-powered tools are completely bonkers. Anyone who knows a bit of this and a bit of that can build something - and make a lot of money off it - if they know what they're doing, or at least have the curiosity and drive to figure it out.
