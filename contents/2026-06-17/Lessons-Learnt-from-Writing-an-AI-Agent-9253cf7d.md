---
source: "https://www.browserless.io/blog/writing-an-ai-agent"
hn_url: "https://news.ycombinator.com/item?id=48575480"
title: "Lessons Learnt from Writing an AI Agent"
article_title: "How to Write an AI Agent: Lessons Learnt"
author: "risingsong"
captured_at: "2026-06-17T21:44:11Z"
capture_tool: "hn-digest"
hn_id: 48575480
score: 6
comments: 0
posted_at: "2026-06-17T19:24:04Z"
tags:
  - hacker-news
  - translated
---

# Lessons Learnt from Writing an AI Agent

- HN: [48575480](https://news.ycombinator.com/item?id=48575480)
- Source: [www.browserless.io](https://www.browserless.io/blog/writing-an-ai-agent)
- Score: 6
- Comments: 0
- Posted: 2026-06-17T19:24:04Z

## Translation

タイトル: AI エージェントの作成から学んだ教訓
記事のタイトル: AI エージェントの書き方: 学んだ教訓
説明: Browserless のエージェント ブラウザの構築から苦労して得た 5 つの教訓: ホスト性、トークンの効率、スキル、成熟したプロトコル、頑固なモデルがブラウジング エージェントが実際に機能するかどうかを決定する理由。

記事本文:
プラットフォームの概要 プラットフォームの概要 Browsers as a Service 大規模な管理型ヘッドレス ブラウザ API シンプルな HTTP MCP サーバーを介したブラウザ タスク MCP を介したブラウザの自動化 セルフホスト 独自のインフラストラクチャで実行 Web スクレイピング用のソリューション ブロックされずにデータを抽出 自動化用 あらゆるブラウザ ワークフローを自動化 テスト用 ヘッドレス テストとライトハウス監査用 AI エージェント用 エージェントに実際のブラウザを提供 スクリーンショットと PDF 用 ピクセル完璧なキャプチャとドキュメント エンタープライズ SLA、コンプライアンス、セキュリティ、サポート リソース ブラウザレスを構築するお客様 ブログ スクレイピング、自動化、AI ガイド トラスト センター セキュリティ、プライバシー、コンプライアンス ブラウザレスのチームについて 料金ドキュメント ログイン お問い合わせ サインアップ ログイン サインアップ お問い合わせ プラットフォームの概要 Browsers as a Service API MCP サーバー Web スクレイピング用のセルフホスト ソリューション 自動化用 テスト用 AI エージェント用 スクリーンショットと PDF エンタープライズ リソース用 お客様のブログ トラスト センター 価格設定について ドキュメント ログイン レッスンAI エージェントの作成から学んだこと
2.トークンを追いかけることになります
3. 明確に定義された問題はまったく問題ではありません (スキルはあなたの友達です)
4. 車輪の再発明をすることになる
エージェントをホストするのではなく、サービスとして使用します
トークンは重要です、ビジョンだけに頼らないでください
Browserless では、過去数か月かけて独自のエージェント ブラウジング エクスペリエンスを構築してきました。私たちは皆、より LLM フレンドリーになり始めており、LLM がルートを発見し、REST API を簡単に利用できるようになりました。したがって、LLM が使用できる独自の MCP サーバーを作成するのは簡単なことのように思えました。
その後、私たちはすでに持っているものを使って自律エージェントを構築できることに気づきました。
道は険しかった、それは否定できない。 HackerNews は毎日、新しい CLI ツールをめぐって大騒ぎになっており、今回は「本当の」ことで眉をひそめることになるでしょう。

完全に自律的なもの、または自律的なブラウジングに革命をもたらす Anthropic エージェントとの往復を行う新しい Python ラッパーです。
状況は非常に曖昧であり、誰もが独自の解決策を持っています。しかし、エージェントを構築する過程で学んだいくつかのことは、(まだ導入されていないとしても) 1 年後、LLM 空間が (昨年と同様に) 大きく変化する今から非常に重要になる鉄則であると思われます。
このブログ投稿から何かを引用する場合は、これにさせてください。ホスピタリティが鍵です。 AI、特にエージェント ブラウジングは、過去 20 年間の DevEx がどのようにして Thing-as-a-Service の哲学に進化したかを忘れつつあるようです。
こう言っておきます。よほど特別な理由がない限り、ベアメタル サーバーを最初から構築する人はいませんし、Amazon、Google、DO、Microsoft、その他多くのクラウド プロバイダーは、インターネット インフラストラクチャ全体の基礎的な部分を占めています。
よほどの理由がない限り、GitHub、BitBucket、または Gitea をホストするクラウド プロバイダーに CI を支払います。チームチャットと同期? Slack または Teams の料金を支払います。内部文書とチケット追跡?同じ話です。
これ自体は悪いことではありません。実際、これによりチームの生産性が向上し、はるかに高品質の製品を提供できるようになります。そしてこれは、AI、特に自律エージェントが最終的には「プラグ アンド プレイ」機能として提供されることになります。 Anthropic/OpenAI/Google は、独自のインフラストラクチャ上でマネージド エージェントを提供しており、それをサービスとして利用して支払うことができます。エージェント ブラウジングもプラグ アンド プレイになります。 CodeRabbit または Devin が CI 構成に「プラグイン」されるのと同じ方法で、MCP を介したエージェント ブラウジングがプラグ アンド プレイ機能として提供されます。
ブラウザレスでは、ある種の賭けをしました

このアプローチについて。
現在、これらのエージェント ブラウザ ツール (ブラウザ ハーネス、Vercel エージェント ブラウザ CLI) は、ユーザーがホストして実行するプログラムです。実行するには少なくとも 4 つの要素が必要です。
読み書きスキルのためのファイル システム
もちろん、これらのリポジトリをローカル マシンにダウンロードすると、クラウド内のブラウザを使用できますが、これらのツールはリモート ブラウザ ≠ リモート インフラストラクチャであることを忘れています。そのため、ブラウザレス MCP はプラグ アンド プレイであり、側で何もホストすることなく、任意の MCP 対応 LLM に接続できます。
2.トークンを追いかけることになります
もう一つの避けられない疑問は、これはバブルなのかということです。
補助金が終了し、投資家が資金の返還を求めたら、補助金は崩壊するだろうか？トークンの使用は持続可能ですか?答えが何であれ、モデルが改良され、より多くの計算能力が必要になるたびに、私たちは甘いトークンをすべて支払うたびに、少しずつ苦しむことになります。ここで 2 番目の点に移ります。すべてのトークンにはコストがかかります。
これは、トークンには支払われるべき価値があるというトートロジーのようなものです。しかし、エージェントのブラウジング環境はこれに関係していないようです。最も正確、最も信頼性が高い、または最もステルスであると主張する新しいエージェント ブラウジング リポジトリが隔週で作成されます。そして厳密に言えば、これらはすべて良いものです。問題は、これらすべてがトークンの効率を犠牲にして行われることです。
エンジニアは、純粋な HTML を LLM に送信することは、LLM にページのスクリーンショットを撮るほど正確ではないことをすぐに認識しました。また、LLM がスクリーンショットを撮ってそれを記述し、視覚的なチャンクに分割し、各チャンクを処理し、生のクリック コマンドを座標に送信する場合ほど正確ではありません。繰り返しますが、これによりモデルがより正確になることは間違いありません。しかし、それは同時に非常に高価で遅くなります。
ブラウザレスでは、ある種の賭けをしました

LLM にページのテキスト形式ベースの表現 (タブとネットワーク情報を含む) を提供するという意味で、ハイブリッド アプローチを採用することについてです。LLM はそれを使用して、ページの状態を推測し、次のアクションを検討します。
これでは不十分な場合に限り、LLM にスクリーンショットを撮って、何がブロックされているかを確認することをお勧めします。これにより、MCP エージェントは速度とトークン消費の両方においてはるかに効率的になったのは間違いありません。たとえば、Browser Harness に対して実行された次のベンチマークを考えてみましょう。
LLM に、それが動作するのに十分な優れたものを与えた場合、それを行うために小遣いを消費する必要はありません。
3. 明確に定義された問題はまったく問題ではありません (スキルはあなたの友達です)
Cookie バナーを閉じることができないため、フライト スケジュールを選択できないエージェント。 CAPTCHA が原因でサイトにログインできないエージェント。セッション中にプロキシが必要であることに気づくエージェント…これらはきっと聞き覚えがあると思います。そしてそれらはすべて、エージェントが問題を具体的に解決する方法を知らないために起こります。ブラウザ ハーネス ベンチマークの最後の項目、ブラウザレスでは 2 分かかる作業に 8 分かかることに懐疑的だったとしても、そんなことはありません。これは、Browser Harness エージェントがセッションの途中までボットがブロックされていることに気づかず、Google で結果を積極的に検索し続けたため発生しました。それはうまくいきました…最終的には。
私が言いたいのは、スキルはエージェントの友達だということです。通常の LLM はすでにこのことを認識しており、設計、プログラミング、数学、読み書きのスキルをバンドルしています。
MCP ブラウザレス エージェントは、スキルにも賭けています。つまり、LLM が CAPTCHA、MFA、シャドウ DOM、タブ管理などをどのように扱うべきか明確に説明されているということです。ある意味、LLM が使用できるように、スキルの調整と作成により多くの時間を費やしています。

可能な限り最も一般的な方法でエージェントに通知します。
もちろん、インタラクションのすべてのステップをいつでも説明することはできますが、それはプロンプトを改良するためにすべての時間とトークンを費やすことを意味します。
4. 車輪の再発明をすることになる
現在のエージェントの状況は、ドットコム時代のように感じられます。標準化の欠如、誇大宣伝、トピックに関する混乱...私たちはこの映画をこれまでに何度も見てきたので、すでに結末を知っています。
状態の永続性、セッション管理、さらにはエージェントが障害を報告する方法などを処理する統一された方法がないため、全員が独自のカスタム ロジックを最初から構築することになります。
独自の個人用ブラウジング エージェントを作成してホストするということは、要素選択用の別のラッパーや、不安定なネットワーク状態に合わせた特注の再試行メカニズムを作成することになることを意味します。しかし、ドットコム時代の混乱の中でも、ブラウザー、HTTP プロトコル、ストリーミング、さらにはサービスとしてのソフトウェア (Warcraft は本当にベルを鳴らすはずです) など、人々が普及するとわかっていたテクノロジーがいくつかありました。
これがブラウザレスで私たちが行ったもう 1 つの賭けです。スキルと MCP がおそらく普及するプロトコルです。私たちは、業界が本当に採用すべき成熟したアーキテクチャを採用しました。 MCP を使用することで、開発者が不安定な再試行メカニズム、特定のエージェント スキル、セッションやプロキシの管理用のラッパーを作成する必要がなくなりました。プラグアンドプレイの合理化された本番環境に対応したエクスペリエンスにより、基盤となるインフラストラクチャではなくエージェントの目的に集中できるようになります。
LLM の背後にあるテクノロジーは今後も存続します。創造的で柔軟性があり、協力的で陽気な良い人です。しかし、閲覧エージェントにとって、それは実際には責任です。頑固で賞品に目を向け続けるモデルが必要です。

当初の目的から逸脱することはありません。これが往復番号 73 かどうかは気にしません。Web は騒がしい場所で、ポップアップ、無関係なサイドバー、混乱を招く UI パターンで満たされており、目に見えるものすべてにフレンドリーなモデルにとっては基本的にマタタビのようなものです。
「ニュースレターに登録してください」モーダルや「あなたは 1,000,000 人目の訪問者です! iPad が当たりました」という広告に気を取られないモデルが必要です。障害を認識し、指示された回避策を試し、賞品に目を光らせます。失敗した場合、それが良いかもしれないと考えてインターネットのまったく別の部分に移動して礼儀正しく「支援」しようとするのではなく、目標に集中し続け、そこから移動しません。トークンの支払い時には、常に正確さが礼儀正しさを上回ります。
興味深いことに、この点では Anthropic モデルと xAI モデルが最も「頑固」であるのに対し、OpenAI モデルは非常に簡単に諦めることがわかりました。ブラウザレス MCP を Opus または Grok に接続することをお勧めします。
MCP エージェントについては、別のブログ投稿ですでに紹介しました。他の競合他社との正確な違いについてのより一般的なアイデアについては、それを確認してください。技術的な側面について詳しく説明し、MCP に接続する方法については、ドキュメントを参照してください。オープンソースなのでローカルで実行できます。また、Anthropic/LLM アカウントがなくても、ブラウザレス CLI を使用して使用することもできます。
ユーザーに合わせて拡張するブラウザー インフラストラクチャ。
お客様のエクスペリエンスを向上させ、サイトのトラフィックを分析するために Cookie を使用します。もっと詳しく知る 。
以下で Cookie の設定を管理します。必要な Cookie はサイトが機能するために必要なため、無効にすることはできません。
サイトが機能するために必要です。
訪問者がサイトをどのように利用しているかを理解するのにご協力ください。
関連する広告やキャンペーンを配信するために使用されます。

## Original Extract

Five hard-won lessons from building Browserless's agentic browser: why hostability, token efficiency, skills, mature protocols, and a stubborn model decide whether a browsing agent actually works.

Platform Overview The platform at a glance Browsers as a Service Managed headless browsers at scale APIs Browser tasks over simple HTTP MCP Server Browser automation over MCP Self-Hosted Run on your own infrastructure Solutions For Web Scraping Extract data without getting blocked For Automation Automate any browser workflow For Testing Headless tests and Lighthouse audits For AI Agents Give agents a real browser For Screenshots & PDFs Pixel-perfect captures and documents For Enterprise SLAs, compliance, security & support Resources Customers Who builds on Browserless Blog Scraping, automation & AI guides Trust Center Security, privacy & compliance About The team behind Browserless Pricing Docs Log In Contact Us Sign Up Login Sign Up Contact Us Platform Overview Browsers as a Service APIs MCP Server Self-Hosted Solutions For Web Scraping For Automation For Testing For AI Agents For Screenshots & PDFs For Enterprise Resources Customers Blog Trust Center About Pricing Docs Log In Lessons Learnt From Writing an AI Agent
2. You will be chasing the token
3. A well-defined problem is not a problem at all (AKA skills are your friend)
4. You will be reinventing the wheel
Don't host the agent, use it as a service
Tokens matter, don't rely only on vision
At Browserless, we've spent the last few months building our own agentic browsing experience. We've all started being more LLM-friendly, letting LLMs discover our routes and consume our REST APIs easily. So it seemed like a no-brainer to write our own MCP server that LLMs could consume.
Then we realized we could build an autonomous agent with what we already had.
The road's been hard, there's no denying that. Every other day, HackerNews gets a meltdown over the new CLI tool that, this time for "real", will make browsing fully autonomous, or the new Python wrapper that makes round trips to an Anthropic agent that will revolutionize autonomous browsing.
We get it, the situation is very nebulous, and everyone has their own solution. But some things that we learnt during the process of building our agent seem to almost be cardinal rules that (if not already in place) will be crucially valued a year from now, when the LLM space is much different (just like it was last year).
If you are going to take something from this blog post, please let it be this. Hostability is key. AI, and agentic browsing in particular, seems to be forgetting how the last two decades of DevEx have evolved into a Thing-as-a-service philosophy.
Let me put it this way: unless you have an extremely specific reason, nobody builds their bare metal servers from scratch, and Amazon, Google, DO, Microsoft, and many other cloud providers are a fundamental part of the whole internet infrastructure.
Unless you have an extremely good reason not to, you do CI paying GitHub, BitBucket, or a cloud provider to host Gitea for you. Team chat and sync? You pay for Slack or Teams. Internal documentation and ticket tracking? Same story.
This is not bad on its own. It actually allows teams to be more productive and deliver products with much higher quality. And it's something that AI, and autonomous agents in particular, will eventually be offered as a "plug and play" feature. You can already see Anthropic/OpenAI/Google providing managed agents on their own infrastructure, which you consume and pay as a service. Agentic browsing will also be plug and play . In the same way CodeRabbit or Devin is "plugged" to your CI config, agentic browsing through MCP will be offered as a plug-and-play feature.
At Browserless, we kind of made a bet on this approach.
Currently, these agentic browser tools (your Browser-harness, your Vercel agent-browser CLI) are a program you host and execute. You need at least four pieces to run it:
A file system to read and write skills
Of course, when you download these repos to your local machine, you can use browsers in the cloud, but these tools forget that remote browser ≠ remote infrastructure . That's why Browserless MCP is plug-and-play, and can connect to any MCP-capable LLM, without having to host anything on your end .
2. You will be chasing the token
Another unavoidable question: is this a bubble?
When subsidies end and investors want their money back, will it implode? Will token usage be sustainable? Regardless of the answer, we are progressively suffering a bit every time from paying all those sweet tokens, every time a model gets better and requires more computing power. Which brings me to my second point: every token costs.
It's almost like a tautology, saying that tokens have a value to be paid. But the agentic browsing environment doesn't seem to be concerned with this. Every other week, a new agentic browsing repo comes in claiming to be the most accurate, or the most reliable, or the most stealthy. And strictly speaking, all of these are good. The problem is that all this is done sacrificing token efficiency.
Engineers quickly realized that sending pure HTML to the LLM was not as accurate as having the LLM take a screenshot of the page, which, in turn, was not as accurate as taking a screenshot and having the LLM describe it, divide it into visual chunks, process each chunk, and send raw click commands to coordinates . Again, this undeniably makes the model more accurate . But it also makes it bloody expensive and slow.
At Browserless, we kind of took a bet on having a hybrid approach, in the sense of providing the LLM with a text-format-based representation of the page (including tab and network information), which it can use to infer the state of the page and think through the following action.
If, and only if, this is not enough, we encourage the LLM to take a screenshot and see what's blocking it. This has undeniably made our MCP agent much more efficient in both speed and token spending. Take, for instance, the following benchmark done against Browser Harness:
If you give the LLM something good enough that it can work with, it doesn't have to burn your allowance to do it.
3. A well-defined problem is not a problem at all (AKA skills are your friend)
An agent that can't select a flight schedule because it can't close the cookie banner. An agent that can't log in to a site because of a CAPTCHA. An agent that realizes mid-session that it needs a proxy… I bet these sound familiar. And all of them happen because the agent doesn't know how to solve the problem specifically. If you were skeptical of the last item in the Browser Harness benchmark, taking 8 minutes to do what took Browserless 2 minutes, don't be. It happened because the Browser Harness agent didn't realize that it was being bot-blocked until mid-session, and proceeded to actively search the results on Google. Which worked… in the end.
The point I'm trying to make is that skills are an agent's friend. Regular LLMs already realized this and bundle design, programming, mathematics, reading, and writing skills.
The MCP Browserless Agent is also making a bet on skills: that the LLM has been explicitly explained how it should deal with CAPTCHAs, MFA, shadow DOMs, tab management, etc. In a way, we spend more time tweaking and creating skills, so the LLMs can use the agent in the most general way possible.
Of course, you can always describe every single step of the interaction, but that means spending all your time and tokens refining the prompt.
4. You will be reinventing the wheel
The current agentic landscape feels like the DotCom era: the lack of standardization, the hype, the confusion about topics... We've seen this movie so many times before that we already know the ending.
Because there isn't a unified way to handle things like state persistence, session management, or even how an agent should report a failure, everyone ends up building their own custom logic from the ground up.
Writing and hosting your own personal browsing agent means you'll find yourself writing yet another wrapper for element selection or a bespoke retry mechanism for flaky network conditions. But even in the chaos of the DotCom era, there were some pieces of tech that folks just knew would prevail: the browser, the HTTP protocol, the streaming, even the software as a service (Warcraft should really ring a bell).
Here's the other bet we made at Browserless: skills and MCP are probably the protocol that will prevail. We went with the mature architecture that the industry really should be adopting. By using MCP, we've eliminated the need for developers to write wrappers for flaky retry mechanisms, for specific agent skills, for managing sessions or proxies. Just a plug-and-play, streamlined, production-ready experience that allows you to focus on your agent's objective rather than the underlying infrastructure.
The technology behind LLMs is here to stay. It is creative and flexible, it's a supportive and jolly good fella. For browsing agents, however, that's actually a liability. You want a model that is stubborn, one that keeps its eyes on the prize and doesn't derail from the original objective. One that doesn't care if this is round trip number 73. The web is a noisy place, filled with pop-ups, irrelevant sidebars, and confusing UI patterns that are basically catnip for a model that is friendly with everything it sees.
You want a model that doesn't get distracted by a "Sign up for our newsletter" modal or a "You are the 1,000,000th visitor! You won an iPad" ad. It sees the hurdle, tries the instructed workaround, and keeps its eyes on the prize. If it fails, rather than politely trying to "help" by navigating to a completely different part of the internet because it thought that might be nice, it stays focused on the goal and does not move from there. Precision beats politeness every single time when you're paying for the tokens .
Curiously, we've found that Anthropic and xAI models are the most "stubborn" in this regard, while OpenAI models give up really easily. We recommend plugging the Browserless MCP into Opus or Grok!
We already introduced the MCP agent in another blog post , check it out for a more general idea of exactly what sets us apart from other competitors. You can read about more technical aspects and learn how to connect to our MCP in our docs . It is open source, so you can run it locally . And you can even use it without an Anthropic/LLM account, through the Browserless CLI .
Browser infrastructure that scales with you.
We use cookies to enhance your experience and analyze site traffic. Learn more .
Manage your cookie preferences below. Necessary cookies cannot be disabled as they are required for the site to function.
Required for the site to function.
Help us understand how visitors use the site.
Used to deliver relevant ads and campaigns.
