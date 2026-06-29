---
source: "https://blog.kilo.ai/p/announcing-next-edit-in-kilo-powered-by-inception"
hn_url: "https://news.ycombinator.com/item?id=48725881"
title: "Next-Edit in Kilo, Powered by Inception Diffusion LLMs"
article_title: "Announcing Next-Edit in Kilo, Powered by Inception"
author: "volodia"
captured_at: "2026-06-29T22:22:52Z"
capture_tool: "hn-digest"
hn_id: 48725881
score: 2
comments: 0
posted_at: "2026-06-29T22:02:21Z"
tags:
  - hacker-news
  - translated
---

# Next-Edit in Kilo, Powered by Inception Diffusion LLMs

- HN: [48725881](https://news.ycombinator.com/item?id=48725881)
- Source: [blog.kilo.ai](https://blog.kilo.ai/p/announcing-next-edit-in-kilo-powered-by-inception)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T22:02:21Z

## Translation

タイトル: Next-Edit in Kilo、Powered by Inception Diffusion LLM
記事のタイトル: Inception を活用した Kilo での Next-Edit の発表
説明: Mercury Edit 2 では、Next-Edit が 1 か月間無料 (キロ単位)

記事本文:
Inception を活用した Kilo での Next-Edit を発表
キロのブログ
Inception を活用した Kilo での Next-Edit を発表
Mercury Edit 2 では、Next-Edit が 1 か月間無料 (キロ単位)
Ari Messer 2026 年 6 月 22 日 8 シェア 思考のスピードでのコーディングが大幅にアップグレードされました。本日、Inception の Mercury Edit 2 モデルを搭載した Next-Edit in Kilo をご紹介できることを嬉しく思います。 Next-Edit は、最近のスニペット、ファイルのコンテキスト、編集履歴に基づいて次に何が来るかを予測する共同編集者のようなものです。
この飛躍を記念して、来月の間、Next-Edit を誰でも完全に無料で利用できるようにします。アクティベートするトライアルやクレジット カードの入力は必要ありません。今後 30 日間、Next-Edit powered by Inception を使用するたびに、Kilo Gateway を介して全額負担されます。
従来のオートコンプリートは、カーソルより先の次の数トークンまたは行を予測します。それは、次の単語を推測する熱心なタイプライターのようなものです。
Next-Edit は根本的に異なります。ただ前を向いているのではなく、あなたの隣に座っているエリート共同編集者のように機能します。 Next Edit Suggestions または NES とも呼ばれる Next-Edit は、より速く、よりスマートに作業するための方法です。 Next-Edit は、最近の編集内容と広範なコードベース コンテキストを分析することにより、ファイル全体で次に何を変更しようとしているかを予測します。それは、直前に操作した関数のリファクタリング、ブロック全体の変数の名前変更、カスケード機能変更の実装などです。
Mercury Edit 2 は最先端の拡散 LLM (dLLM) アーキテクチャに基づいて構築されているため、一度に 1 トークンずつ、左から右にコードを生成するわけではありません。代わりに、テキストを並行して絞り込み、Tab キーを押すだけで受け入れられます。この「編集にジャンプ」機能は慣れるまでに少し時間がかかりますが、一度コーディングを始めると気に入っていただけると思います。

流れの中にいるような感覚。
オートコンプリートは Kilo で人気の機能であり、今後も廃止される予定はありません。 Next-Edit を含めるために機能を拡張しているところです。
画期的なベンチマークによる裏付け
Mercury Edit 2 を選んだのは、単に未来的なサウンドだからというだけではありません。そのパフォーマンス指標は本当に驚異的です。 Inception は、このモデルを Instinct、Fill-in-the-Middle (FIM)、Next-Edit Prediction (NEP) などの一連の厳格な業界ベンチマークに通過させ、データがそれを物語っています。
48% 高い受け入れ率: 実際の人間のフィードバックに基づく高度な強化学習法 (KTO) による調整のおかげで、開発者は Mercury Edit 2 の編集を以前のイテレーションよりも 48% 多く受け入れています。
選択性が 27% 向上: 常に思考の流れを狂わせる騒々しい AI を好む人はいません。 Mercury Edit 2 は選択性が 27% 高く、高品質で的を絞った編集がある場合にのみ中断され、気が散る要素は最小限に抑えられます。
比類のない品質と速度: カスタムの Next-Edit フレームワークと速度が最適化されたフロンティア モデルに対する並列レイテンシー テストでは、Mercury Edit 2 がリアルタイム応答性の点で一貫して頂点に立っています。
出典: Inception Blog マーキュリーの 2 つのフレーバー 編集 2 がキロで登場
私たちは、開発者にはそれぞれ独自のリズムがあることを知っています。ワークフローを完全に制御できるように、Kilo には Mercury Edit 2 の 2 つの異なるバージョンが含まれています。
Next-Edit モード (新しいデフォルト): 次の論理コード変更を積極的に予測する強力なエディターです。
一般オートコンプリート モード: 標準的な逐次コーディング セッション用の超高速で古典的な行とトークンのフィラーです。
注意: Next-Edit は、すべての新しい Kilo ユーザーのデフォルトのエクスペリエンスになりました。以前にオートコンプリートのデフォルトを設定していた場合は、手動で N に変更する必要があります。

外部編集。スピンして平行拡散の違いをご自身で体験することを強くお勧めします。
Next-Edit がコーディング方法を永遠に変えると私たちは考えていますが、Kilo は完全に開発者の選択に基づいて構築されています。モデルをオートコンプリートのみに戻したい場合は、数回クリックするだけで完了します。次の編集とオートコンプリートの設定をワークスペースから直接シームレスに変更できます。
Kilo の VS Code 拡張機能サイドバー内の歯車アイコン (設定) を見つけてクリックします。
[設定] > [拡張機能] > [Kilo Code] で、現在のプロジェクトのニーズに合わせて設定を切り替えるか、Next-Edit と従来のオートコンプリートを切り替えます。
無料の次回編集機能をお楽しみください
追加のものをインストールしたり、個別のトライアルを開始したり、複雑な構成を行う必要はありません。来月中のいつでも Kilo を起動するだけで、Kilo Gateway (これがデフォルト) を使用している限り、次世代の拡散を利用した予測に完全に無料で自動的にアクセスできるようになります。
2026 年 7 月 23 日にプロモーションが終了しても、手頃な市場価格で Mercury Edit 2 を引き続き使用できます。
8 シェア この投稿に関するディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

With Mercury Edit 2, Next-Edit is Free in Kilo for an Entire Month

Announcing Next-Edit in Kilo, Powered by Inception
Kilo Blog
Subscribe Sign in Announcing Next-Edit in Kilo, Powered by Inception
With Mercury Edit 2, Next-Edit is Free in Kilo for an Entire Month
Ari Messer Jun 22, 2026 8 Share Coding at the speed of thought just got a massive upgrade. Today, we’re thrilled to introduce Next-Edit in Kilo , powered by the Mercury Edit 2 model from Inception . Next-Edit is like a co-editor that predicts what should be coming next based on your recent snippets, file context, and edit history.
To celebrate this leap forward, we are making Next-Edit completely free for everyone for the next month . No trials to activate and no credit cards to enter—anytime you use Next-Edit powered by Inception over the next 30 days, it’s completely on us via the Kilo Gateway.
Traditional autocomplete predicts the next few tokens or lines ahead of your cursor. It’s like an eager typewriter guessing your next word.
Next-Edit is fundamentally different. Instead of just looking forward, it acts like an elite co-editor sitting right next to you. Often referred to as Next Edit Suggestions or NES, Next-Edit is a way to work faster and smarter. By analyzing your recent edits and broader codebase context, Next-Edit anticipates what you’re about to change next across your file—whether that’s refactoring a function you just touched, renaming a variable across a block, or implementing a cascading feature change.
Because Mercury Edit 2 is built on a cutting-edge diffusion LLM (dLLM) architecture, it doesn’t generate code left-to-right, one token at a time. Instead, it refines text in parallel and you just hit Tab to accept. This “jump to edit” feature takes some getting used to, but once you start coding with it, we hope you will love the feeling of being in the flow.
Autocomplete has been a popular feature in Kilo, and it’s not going anywhere. We’re just expanding the functionality to include Next-Edit.
Backed by Breakthrough Benchmarks
We didn’t just pick Mercury Edit 2 because it sounds futuristic; its performance metrics are genuinely staggering. Inception put this model through a rigorous set of industry benchmarks —including Instinct, Fill-in-the-Middle (FIM), and Next-Edit Prediction (NEP)—and the data speaks for itself:
48% Higher Acceptance: Thanks to alignment via an advanced reinforcement learning method (KTO) based on real human feedback, developers accept Mercury Edit 2’s edits 48% more often than previous iterations.
27% More Selective: Nobody likes a noisy AI that constantly derails your train of thought. Mercury Edit 2 is 27% more selective, meaning it only interrupts when it has a high-quality, targeted edit, keeping distractions to an absolute minimum.
Unmatched Quality & Speed: In side-by-side latency testing against custom next-edit frameworks and speed-optimized frontier models, Mercury Edit 2 consistently takes the crown for real-time responsiveness.
Source: Inception Blog Two Flavors of Mercury Edit 2 Now in Kilo
We know every developer has their own rhythm. To give you total control over your workflow, you will now find two separate versions of Mercury Edit 2 inside Kilo:
Next-Edit Mode (The New Default): The powerhouse editor that actively anticipates your next logical code modification.
General Autocomplete Mode: The ultra-fast, classic line-and-token filler for standard, sequential coding sessions.
Heads Up: Next-Edit is now the default experience for all new Kilo users . If you had previously set a default for autocomplete, you will need to manually change to Next-Edit. We highly recommend giving it a spin to experience the parallel diffusion difference for yourself.
While we think Next-Edit will change how you code forever, Kilo is built entirely around developer choice. If you ever want to change the model back to autocomplete-only, doing so takes just a few clicks. You can seamlessly modify your Next-Edit and autocomplete settings right from your workspace:
Locate and click the Gear Icon (Settings) inside Kilo’s VS Code extension sidebar.
Under Settings > Extensions > Kilo Code, toggle your preferences or switch between Next-Edit and classic Autocomplete to suit your current project’s needs.
Enjoy Free Next-Edit Functionality
There is nothing extra to install, no individual trial to start, and no complex configuration required. Just fire up Kilo anytime over the next month and you’ll automatically have access to next-generation, diffusion-powered predictions entirely for free, as long as you are using the Kilo Gateway (this is the default).
When the promotion ends, on July 23rd, 2026, you can continue using Mercury Edit 2 at affordable market rates.
8 Share Discussion about this post Comments Restacks Top Latest Discussions No posts
