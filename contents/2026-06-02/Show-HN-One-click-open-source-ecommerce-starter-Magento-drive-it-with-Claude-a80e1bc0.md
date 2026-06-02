---
source: "https://ecommerce-ai-starter.graycore.io"
hn_url: "https://news.ycombinator.com/item?id=48356655"
title: "Show HN: One-click open-source ecommerce starter (Magento), drive it with Claude"
article_title: "Open Source Ecommerce AI Starter — build an ecommerce store with Claude"
author: "damienwebdev"
captured_at: "2026-06-02T00:40:14Z"
capture_tool: "hn-digest"
hn_id: 48356655
score: 2
comments: 0
posted_at: "2026-06-01T13:32:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: One-click open-source ecommerce starter (Magento), drive it with Claude

- HN: [48356655](https://news.ycombinator.com/item?id=48356655)
- Source: [ecommerce-ai-starter.graycore.io](https://ecommerce-ai-starter.graycore.io)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T13:32:44Z

## Translation

タイトル: Show HN: ワンクリックのオープンソース e コマース スターター (Magento)、Claude で起動
記事のタイトル: オープンソース e コマース AI スターター — クロードと一緒に e コマース ストアを構築する
説明: コードを記述せずにオープンソースの e コマース ストアを構築します。ワンクリックの GitHub スターター、Codespaces、および Claude — Magento® または Mage-OS での学習、プロトタイピング、構築を行うことができます。ホスト型 SaaS ではありません。
HN テキスト: 私は生計のために E コマース ストアを構築しています (主に Magento オープンソース)。常に最悪の部分は最初の部分であり、チームに所属している場合は特にそうです。動作するローカル環境を取得するには、適切な PHP バージョンと拡張機能、Redis/Valkey、OpenSearch、nginx、MySQL/MariaDB、メール キャッチャー、Magento のインストールをセットアップする必要があり、人生の選択を疑うほど時間がかかります。まさに千切りの死である。役に立つことを一行も書くまでに、何時間もヤクの毛を剃るのは無意味だ。これは、技術者ではない人が Magento を試してみるのを諦めるまさにポイントでもあります。エージェント時代の到来に伴い、私はいくつかのことを達成したいと考えました。 1. ボタン 1 回のクリックで新しいストアを立ち上げることができるようにする 2. 開発を支援する AI エージェントを愚かなほど簡単に導入できるようにする。 3. 潜在的に (これが、スターター ジェネレーター サイトが AI に熱心すぎる理由です)、開発者の関与を限定しながら、販売業者が自分で店舗を構築できるようになります。ポイント 3 は、現時点では私にとっては「研究」です。これが正気なのか非常識なのかはわかりませんが、私が 1 年ほど前から楽しんでいるアイデアです。 Hacker News で、(良くも悪くも) PM に「コードを書かせている」人々に関するさまざまな投稿を見たことがありますが、さらに素人のユーザーにも試してみてはいかがでしょうか?それはさておき、私は以前、これを解決するために愚かな時間を費やしました

過去数年間のまさに問題。 AI が登場する前に、私は 2 つのローカル Docker 環境 (magedocker / mage2docker) を作成しましたが、最終的にはさまざまな理由で最悪になったので、メンテナンスをやめました。自分で構築する新しいものについては、結果として得られる e コマース ストアの成功に重要だと思ういくつかの重要な要素が必要でした。 - すぐに使える開発コンテナ - さまざまなストアフロント / ディストリビューションから簡単に選択できる機能 (クライアントによって必要なものが異なります。誰が知っているでしょうか?) - 仕事で毎日使用する事前にパッケージ化されたツール。 - CI が事前にパッケージ化されているため、新しいマーチャントで働くたびに車輪を再発明する必要がなくなります。 - Claude ビルトイン (現時点で私が選んだ AI です) 最近、いくつかのことが同時に実現し、これまで考えて繰り返してきた新しいアイデアをすべて 1 つのバンドルにパッケージ化する時期が来たと感じました。そのため、ボタンをクリックして、GitHub でテンプレートからリポジトリを作成し、ブラウザーで完全に構成されたストアでコードスペースを起動できるようになりました。 PHP、nginx、MariaDB、Mailpit、ディストリビューション、およびストアフロント、Hyvä (PHP でレンダリングされたテーマ) または Daffodil (私が管理している Angular ヘッドレス ストアフロント) のいずれかです。これらはすべて実行されており、すぐに接続できます。 Magento Open Source または Mage-OS を選択できます。ベアボーンが必要な場合は Mage-OS Minimal もあります。 Codespaces を使用するだけの場合は、ローカルに何もインストールする必要はありません。 「クリック」から「出店できる」までは約8分。おそらくこの期間を短縮することはできますが、さらに時間と調査が必要になります。まだ荒削りな部分があり、書かなければならないドキュメントがたくさんありますが、結果としてはかなり満足しています。テンプレート、パイプライン、フレームワーク自体、devcontainer などは無料であり、

オープンソース。料金を支払うのは Codespaces / Claude だけです。 Codespaces は、(この環境の場合) 月あたり 30 時間も無料です。したがって、クロードライセンスを持っている場合は、そのまま遊ぶことができます。私が一緒に調達したものは次のとおりです。 - スターター テンプレート: - Magento + Hyvä: https://github.com/graycoreio/magento2-ai-starter-hyva
- Magento + 水仙: https://github.com/graycoreio/magento2-ai-starter-daffodil
- Mage-OS + Hyvä: https://github.com/graycoreio/mage-os-ai-starter-hyva
- Mage-OS + 水仙: https://github.com/graycoreio/mage-os-ai-starter-daffodil
- Mage-OS Minimal (ストアフロントなし): https://github.com/graycoreio/mage-os-ai-starter-minimal

記事本文:
Graycore によるオープンソース e コマース AI スターター 仕組み 費用 次のステップ 始めましょう スターター キット · オープンソース · GitHub + Claude を利用 クロードと会話して、実際の e コマース ストアを構築します。
コーディング方法を知らなくても、無限にカスタマイズ可能な実際の e コマース ストアを構築するためのスターター キット。ブラウザから起動してクロードに尋ねてください。プロトタイピングから開始します。ホスティングを追加すると、本番環境に対応します。
ストアを作成する クリックしてからストアが実行されるまで、約 8 分でどのように機能するかを確認してください。
すでにクロードを持っている場合は無料です。そうでない場合は、Claude Pro で月 20 ドル。
プロンプトが 1 つあると、クロードがファイルを編集し、ストアフロントが更新されます — オンループです。
今日の学習と構築のためのスターター キット。ホスティングを追加すると、実際のライブ ストアに成長する準備が整います。
コードを学ばずに e コマース ストアを始めたい
プラットフォームのロックインをせずに、初日からストアのコードを所有したい
Composer、nginx、MariaDB を設定せずに、ブラウザーでコーディングできる Magento ® + Daffodil を望んでいる開発者です。クライアントの PoC、デモ、またはスタックの調査に適しています
実際の e コマース スタックがどのように機能するかに興味があり、構築したりいじったりして学びたいと考えている
今すぐライブのパブリックストアが必要です。本番環境にはホスティング、SSL、支払い、セキュリティレビューが必要です
Shopify のような、フルマネージドのクリックして起動する SaaS が必要です (コードを所有して実行することになります。これはトレードオフです)
既存の Magento ストアを移行しています。これはまったく新しいスターターであり、移行ツールではありません
スターターは、ブラウザで開くクラウド マシン上で実際の Magento ® ストアを起動し、クロードに作業方法を指示し、信頼する前に間違いを見つけます。すべてがオープンソースです。
あなたのストアをクラウドで準備完了
ボタンをクリックすると、動作する Magento ® ストアがクラウド マシン上で起動し、完全に構成され、すぐに開きます。

あなたのブラウザ。コンピューターにインストールするものやセットアップするものは何もありません。クロードは同じ環境内で直接作業し、同じファイルを見て、同じストアを実行し、目に見える変更を加えます。
開発者向け: スターターには、PHP、nginx、MariaDB、Mailpit、Magento ® 、および Hyvä (PHP レンダリングされたテーマ) または Daffodil (Angular ベースのヘッドレス ストアフロント) を起動する開発コンテナが同梱されており、すべて事前に接続され実行されています。追跡するためのコンポーザー認証も、格闘するための PHP バージョンも、ローカル インストールも必要ありません。 Claude は、サンドボックス API ではなく、完全な環境 (ターミナル、ファイル システム、実行中の Magento、ローカル DB) を継承します。 GitHub コードスペースを利用しています。
クロードが最初に読むガイドブック
クロードは何かに触れる前に、スターターに同梱されているガイドブックを読みます。ガイドブックには、この種の店舗がどのように構築されるか、何を避けるべきか、そしてよくある間違いがどこに隠れているかが説明されているため、基本を一から教えるのに苦労する必要はありません。新しいチームメイトを入社させるときと同じように、まずドキュメントを読んでから作業を始めます。
開発者向け: リポジトリのルートにある AGENTS.md は、Magento の規則、一般的な落とし穴、ツール ポインターをエンコードします。 Claude Code およびその他のエージェント コーディング ツールによって採用された標準規約。
あらゆる変更を自動チェッカー
何かが変更されるたびに、ストアの新しいコピーがコードから再構築され、一連のチェックが実行されます。機能している場合は緑色の ✓ が表示され、何かが壊れている場合は赤色の ✗ が表示されるため、変更を信頼する前にその変更が実際に機能するかどうかを判断できます。
開発者向け: GitHub Actions ワークフローはプッシュごとにストアを再構築し、スターター自体で使用するのと同じテスト スイートを実行します。 check-store を参照してください。
すでにクロードを持っている場合は無料です。
スターターは無料でオープンソースです。 GitHub Codespaces は通常の学習用途を無料でカバーします

層。必要な費用はクロードのみです。すでに Pro を取得している場合は、これで完了です。そうでない場合は、Claude Pro で月 20 ドル。
店内にあるものはすべて平易な英語で。クロードはコードを読み、変更を加えて説明します。
構築可能な店舗まで約 8 分。
クレジットカードなしで始められます。気に入らなかったら捨ててください。
ストアの作成 オープンソース E コマース AI スターター Claude を使用して Magento® または Mage-OS 上に e コマース ストアを構築するためのオープンソース スターター キット。ホスト型 SaaS または運用ストアフロントではありません。
Magento は Adob​​e Inc の登録商標です。このプロジェクトは Adob​​e と提携または承認されていません。 Mage-OS は、Adobe とは関係のない、コミュニティ主導の独立したプロジェクトです。

## Original Extract

Build an open source ecommerce store without writing code. One-click GitHub starter, Codespaces, and Claude — for learning, prototyping, and building on Magento® or Mage-OS. Not a hosted SaaS.

I build Ecommerce stores for a living (Magento Open Source primarily), and the part that has always been the worst is the very beginning, especially so if you're on a team of people. Getting a working local environment means setting up the right PHP version and extensions, Redis/Valkey, OpenSearch, nginx, MySQL/MariaDB, a mail catcher, and a Magento install that takes long enough to make you doubt your life choices. It truly is death by one-thousand cuts. Pointless hours of yak-shaving before you have written a single line of anything useful. It is also the exact point where a non-technical person who wants to try Magento gives up. With the agentic era upon us, I wanted to accomplish a few things: 1. Make it possible to spin up a new store in a single button click 2. Make it stupidly easy to bring in an AI agent to assist in development. 3. Potentially (and this is why the starter generator site is so over-eager on AI) allow merchants to build stores themselves with limited developer involvement. Point 3 is more "research" for me at the moment. I don't know if this is sane or insane, but it's an idea that I've been entertaining for a year or so now. I've seen various posts on Hacker News about people having their PMs "write code" (for better or worse), why not let the even-more lay-user try it? This all aside, I have previously spent a stupid amount of time attempting to solve this exact problem over the past few years. Pre-AI, I wrote two local docker environments (magedocker / mage2docker) that eventually sucked for various reasons so I stopped maintaining them. For anything new that I would build for myself, I wanted a few key elements that I think are critical to success of the resultant ecommerce store: - Working devcontainer out of the box - The ability to easily choose between different storefronts / distros (different clients want different things, who knew?) - Pre-packaged tools that I use every day to do my work. - Pre-packaged CI so that I don't have to re-invent the wheel each time I work for a new merchant. - Claude built-in (it's the AI-of-choice for me at the moment) Recently, a few things came together simultaneously and I felt that it was high time I packaged all of these new ideas I've been noodling on and iterating towards into a single bundle. As such, you can now click a button, have GitHub create a repo from the template, and boot a Codespace with a fully configured store in your browser. PHP, nginx, MariaDB, Mailpit, the distro, and a storefront, either Hyvä (the PHP-rendered theme) or Daffodil (an Angular headless storefront I maintain). These are all running and wired together out-of-the-box. You can pick Magento Open Source or Mage-OS, and there's a Mage-OS Minimal if you want barebones. You don't need to install anything locally if you just want to use Codespaces. From "click" to "a store you can open" is about 8 minutes. I can probably shorten this duration, but it'll take some further time and research. There are still rough edges, and I have a bunch of documentation to write, but I'm reasonably happy with how it all came out. The template, the pipelines, the frameworks themselves, the devcontainer, etc. are free and open source. The only things you pay for are Codespaces / Claude. Codespaces is even free for 30 hours / mo (with this environment). So, if you have a Claude license, you can just play around. Here’s the things I sourced together: - Starter templates: - Magento + Hyvä: https://github.com/graycoreio/magento2-ai-starter-hyva
- Magento + Daffodil: https://github.com/graycoreio/magento2-ai-starter-daffodil
- Mage-OS + Hyvä: https://github.com/graycoreio/mage-os-ai-starter-hyva
- Mage-OS + Daffodil: https://github.com/graycoreio/mage-os-ai-starter-daffodil
- Mage-OS Minimal (no storefront): https://github.com/graycoreio/mage-os-ai-starter-minimal

Open Source Ecommerce AI Starter By Graycore How it works What it costs Next steps Get started Starter kit · Open source · Powered by GitHub + Claude Build a real ecommerce store by talking to Claude.
A starter kit for building a real and infinitely customizable ecommerce store without knowing how to code. Spin it up from your browser and ask Claude. Start by prototyping — production-ready once you add hosting.
Create my store See how it works ~8 minutes from click to a running store.
Free if you already have Claude. $20/mo for Claude Pro if you don’t.
One prompt, Claude edits the file, the storefront updates — on loop.
A starter kit for learning and building today. Ready to grow into a real, live store once you add hosting.
Want to start an ecommerce store without learning to code first
Want to own your store’s code from day one, with no platform lock-in
Are a developer who wants Magento ® + Daffodil ready to code against in your browser , with no composer, nginx, or MariaDB to configure. Good for client PoCs, demos, or exploring the stack
Are curious how a real ecommerce stack works and want to learn by building and tinkering
Need a live, public store today. Production needs hosting, SSL, payments, and a security review
Want a fully-managed, click-to-launch SaaS like Shopify (you’ll own and run the code; that’s the trade-off)
Are migrating an existing Magento store. This is a brand-new starter, not a migration tool
The starter boots a real Magento ® store on a cloud machine you open in your browser, tells Claude how to work on it, and catches mistakes before you trust them. Everything is open source.
Your store, ready in the cloud
Click a button and a working Magento ® store boots on a cloud machine, fully configured, and opens right in your browser. There’s nothing to install on your computer and nothing to set up. Claude works directly inside that same environment — looking at the same files, running the same store, making changes you can see.
For developers: The starter ships a devcontainer that brings up PHP, nginx, MariaDB, Mailpit, Magento ® , and either Hyvä (PHP-rendered theme) or Daffodil (Angular-based headless storefront), all pre-wired and running. No composer auth to hunt down, no PHP versions to wrestle, no local install. Claude inherits the full environment — terminal, file system, running Magento, local DB — not a sandboxed API. Powered by GitHub Codespaces.
A guidebook Claude reads first
Before touching anything, Claude reads a guidebook that ships with the starter. The guidebook explains how this kind of store is built, what to avoid, and where the common mistakes hide — so you’re not stuck teaching it the basics from scratch. It’s the same way you’d onboard a new teammate: read the docs first, then start working.
For developers: An AGENTS.md at the repo root encodes Magento conventions, common pitfalls, and tooling pointers. Standard convention picked up by Claude Code and other agentic coding tools.
An automatic checker for every change
Every time something changes, a fresh copy of your store rebuilds itself from your code and runs a battery of checks. You see a green ✓ if it works, a red ✗ if something broke — so you can tell whether a change actually works before you trust it.
For developers: A GitHub Actions workflow rebuilds the store on every push and runs the same suite of tests we use on the starter itself — see check-store .
Free if you already have Claude.
The starter is free and open source. GitHub Codespaces covers normal learning use on its free tier. Claude is the only required spend — if you’ve already got Pro, you’re done. $20/mo for Claude Pro if you don’t.
Anything in your store, in plain English. Claude reads the code, makes the change, and explains it.
~8 minutes to a store you can build on.
No credit card to start. Throw it away if you don’t like it.
Create my store Open Source Ecommerce AI Starter An open-source starter kit for building an ecommerce store on Magento® or Mage-OS with Claude. Not a hosted SaaS or production storefront.
Magento is a registered trademark of Adobe Inc. This project is not affiliated with or endorsed by Adobe. Mage-OS is an independent, community-driven project not associated with Adobe.
