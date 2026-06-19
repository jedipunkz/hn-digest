---
source: "https://www.xelerate.tech/one-model-wont-save-you/"
hn_url: "https://news.ycombinator.com/item?id=48595915"
title: "One Model Won't Save You: How We Built Our AI Stack"
article_title: "One Model Won’t Save You: How We Actually Built Our AI Stack - xelerate.tech"
author: "pedrocha"
captured_at: "2026-06-19T07:59:12Z"
capture_tool: "hn-digest"
hn_id: 48595915
score: 1
comments: 0
posted_at: "2026-06-19T07:43:06Z"
tags:
  - hacker-news
  - translated
---

# One Model Won't Save You: How We Built Our AI Stack

- HN: [48595915](https://news.ycombinator.com/item?id=48595915)
- Source: [www.xelerate.tech](https://www.xelerate.tech/one-model-wont-save-you/)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T07:43:06Z

## Translation

タイトル: 1 つのモデルではあなたは救われない: AI スタックの構築方法
記事のタイトル: 1 つのモデルではあなたを救えない: AI スタックを実際に構築した方法 - xelerate.tech
説明: 1 つのモデルではあなたを救えない: xelerate.tech で AI スタックを実際に構築した方法

記事本文:
メインコンテンツにスキップ
Enter キーを押して検索するか、ESC キーを押して閉じます
検索を閉じる
連絡する
1 つのモデルでは救われない: AI スタックを実際に構築した方法
xelerate.tech の CTO
ポルト、2026 年 5 月 4 日
先週、私は会社全体にとって「完璧な」AI プラットフォームを見つけようと髪をかきむしっている同僚とコーヒーを飲んでいました。彼は 1 つの契約、1 つのインターフェイス、および 1 つの請求書を必要としていました。
私は彼に幽霊を追っていると言いました。
「画一的な」AI プラットフォームはマーケティング上の迷信です。システム アーキテクチャから人事メールに至るまですべてを単一のモデルで処理させようとすると、単純なものには高価すぎるか、重労働には「愚か」すぎるツールが使用されることになります。私たちは「聖杯」を探すのをやめ、代わりにツールボックスを作り始めました。
「インテリジェンス」の実際のコスト
ベンダー契約に署名する前に、派手なデモを無視して、配管について話しておく必要があります。具体的には、プライバシーと法案です。
セルフホスティング: ハードウェアと予算がある場合は、独自のモデル (Llama や Mistral など) をホスティングすることがプライバシーのゴールド スタンダードです。データが VPC から離れることはありません。しかし、正直に言うと、インフラストラクチャのコストと、これらのクラスターを維持するために必要な専門の人材を考えると、CFO は目が潤むでしょう。
トークンごとの支払い: これは、拡張するための最も透明な方法です。使用した分だけ料金を支払います。大規模な自動化には最適ですが、ガードレールがしっかりしていないと、開発中の 1 つのバグのあるループが昼までに 1 か月の予算を使い切ってしまう可能性があります。
サブスクリプション: これは予測可能であり、金融​​機関が好むものです。ただし、「ソフトリミット」に注意する必要があります。目に見えない天井に達すると、高性能モデルは突然 56k モデムで動作しているように感じられます。
たくさんのチームに気づきました

「デジタル主権」のギャップを無視してください。現在、米国と中国が先を急ぐ中、ヨーロッパは基本的に観客席に座っている。私たちにとって、これはテクノロジーだけの問題ではありません。それはコントロールに関するものです。ヨーロッパの企業が外国のモデルのみに依存するたびに、私たちはコグニティブ インフラストラクチャを借りていることになります。これは単に無視できる項目ではなく、管理しなければならないリスクです。
私たちのスタック: 何を選んだのか
勝者を 1 人も選ぶことはできませんでした。私たちは、従業員が毎日行っている特定の仕事に最適なツールを選択しました。
1. 開発チーム向け: GitHub Copilot Pro+
これについては考えすぎませんでした。私たちのコードは GitHub 上にあります。私たちの開発者は VSCode に住んでいます。 Copilot Pro+ を使用すると、「コンテキスト切り替え税」なしでシームレスな統合が可能になります。これにより最上位モデルへのアクセスが提供され、サブスクリプション モデルは、提供される制限を考慮すると、高生産性のエンジニアリング チームを満足させる最もコスト効率の高い方法です。
2. それ以外の場合: Gemini (Google Workspace)
残りの業務 (マーケティング、人事、管理) をサポートするために、私たちは Gemini に頼っています。これはすでに Google Workspace サブスクリプションにバンドルされているため、参入障壁はゼロでした。別のサブスクリプションを管理する必要がなく、スプレッドシートと電子メールを完全に適切に処理します。
3. 自動化の場合: n8n + OpenAI
次に、バックグラウンド プロセスを結合する接着剤があります。私たちは、強力な自動化のために OpenAI と組み合わせた n8n を標準化しました。 Gemini がドキュメント内で「人間と向き合う」作業を担当する一方で、このデュオは内部のロジックを担当します。私たちは n8n を使用して、CRM、Notion、BambooHR、Google Workspace、社内データベースを接続するワークフローを構築し、データを解析したり意思決定を下すために特定の「頭脳」が必要なときはいつでも、OpenAI の API (および一連の MCP サーバー) を呼び出します。それはウルです

従量課金制のセットアップを開始します。実際にトリガーが起動した場合にのみトークンを消費するため、GPT モデルの本来の能力を犠牲にすることなく、自動化のオーバーヘッドが最小限に抑えられます。
「Run Book」: IQ をタスクに適合させる
電卓で実行できるタスクでプレミアム サブスクリプションを使い果たさないようにするために、エンジニアリング タスクに特定の「Run Book」を実装しました。私たちは AI トークンを無料で使えるものではなく、リソースのように扱います。
目標: 単一のエージェント セッションで実装される問題の数を最大化します。 AI には、レート制限を使い果たす 20 分間のチャットではなく、一度に問題を解決してもらいたいと考えています。
「最高の AI」を探すのはやめましょう。それは存在しません。代わりに、ワークフローを見て、どこにメスが必要か、どこに大ハンマーが必要かを判断してください。チームの生産性は向上し、予算は実際に黒字を維持できるようになります。
Rua António Nicolau d’Almeida n°66 4° Dto
4100-320 ポルト
ポルトのゾーナ インダストリアル
利用規約 クッキーポリシー プライバシーポリシー
© 2026 xelerate.tech |無断転載を禁じます

## Original Extract

One Model Won’t Save You: How We Actually Built Our AI Stack at xelerate.tech

Skip to main content
Hit enter to search or ESC to close
Close Search
Get in touch
One Model Won’t Save You: How We Actually Built Our AI Stack
CTO of xelerate.tech
Porto, 4th of May 2026
I was grabbing coffee with a peer last week who was tearing his hair out trying to find the “perfect” AI platform for his entire company. He wanted one contract, one interface, and one bill.
I told him he was chasing a ghost.
The “one-size-fits-all” AI platform is a marketing myth. If you try to force a single model to handle everything from your system architecture to your HR emails, you’re going to end up with a tool that is either too expensive for the simple stuff or too “dumb” for the heavy lifting. We stopped looking for the “Holy Grail” and started building a toolbox instead.
The Real Cost of “Intelligence”
Before you sign a vendor agreement, you have to look past the flashy demos and talk about the plumbing. Specifically: Privacy and the Bill.
Self-Hosting: If you have the hardware and the budget, hosting your own models (like Llama or Mistral) is the gold standard for privacy. Your data never leaves your VPC. But let’s be honest—the infrastructure costs and the specialized talent needed to maintain those clusters will make your CFO’s eyes water.
Pay-per-token: This is the most transparent way to scale. You pay for what you use. It’s great for high-volume automation, but if you don’t have tight guardrails, a single buggy loop in development can burn through a month’s budget by lunch.
Subscriptions: These are predictable, which finance loves. But you have to watch the “soft limits.” Once you hit that invisible ceiling, your high-performance model suddenly feels like it’s running on a 56k modem.
I’ve noticed a lot of teams ignore the “Digital Sovereignty” gap. Right now, Europe is essentially sitting in the bleachers while the US and China sprint ahead. For us, this isn’t just about tech; it’s about control. Every time a European company relies solely on a foreign model, we’re renting our cognitive infrastructure. It’s a risk we have to manage, not just a line item we can ignore.
Our Stack: Why We Chose What We Chose
We didn’t pick one winner. We picked the best tools for the specific jobs our people do every day.
1. For the Dev Team: GitHub Copilot Pro+
We didn’t overthink this one. Our code lives on GitHub. Our devs live in VSCode. Copilot Pro+ gives us seamless integration without the “context-switching tax.” It provides access to top-tier models, and the subscription model—considering the limits they provide—is the most cost-effective way to keep a high-output engineering team happy.
2. For Everything Else: Gemini (Google Workspace)
To support the rest of our operations—marketing, hr, admin—we lean on Gemini. Since it’s already bundled into our Google Workspace subscription, the barrier to entry was zero. It handles the spreadsheets and the emails perfectly fine without us having to manage another separate subscription.
3. For Automation: n8n + OpenAI
Then there’s the glue holding our background processes together. We’ve standardised on n8n paired with OpenAI for our heavy-duty automation. While Gemini handles the “human-facing” work in our docs, this duo handles the logic under the hood. We use n8n to build the workflows—connecting our CRM, Notion, BambooHR, Google Workspace, and internal databases—and then call OpenAI’s API (and a bunch of MCP Servers) whenever we need a specific “brain” to parse data or make a decision. It’s the ultimate pay-as-you-go setup. We only spend tokens when a trigger actually fires, which keeps our automation overhead lean without sacrificing the raw power of the GPT models.
The “Run Book”: Matching IQ to the Task
To make sure we aren’t burning through our premium subscriptions on tasks a calculator could do, we implemented a specific “Run Book” for our engineering tasks. We treat AI tokens like a resource, not a free-for-all.
The Goal: Maximize the number of issues implemented by a single Agent Session. We want the AI to solve the problem in one go, not engage in a twenty-minute chat that drains our rate limits.
Stop looking for the “Best AI.” It doesn’t exist. Instead, look at your workflow and figure out where you need a scalpel and where you need a sledgehammer. Your team will be more productive, and your budget will actually stay in the black.
Rua António Nicolau d’Almeida nº66 4º Dto
4100-320 Porto
Zona Industrial do Porto
Terms & conditions Cookie Policy Privacy Policy
© 2026 xelerate.tech | All Rights Reserved
