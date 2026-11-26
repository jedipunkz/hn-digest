---
source: "https://talira.ai/"
hn_url: "https://news.ycombinator.com/item?id=50651470"
title: "Show HN: Talira – Search engine for software, built for AI agents"
article_title: "Talira"
image: "https://talira.ai/static/brand/og.png"
author: "nathan_f77"
captured_at: "2026-11-26T04:54:03Z"
capture_tool: "hn-digest"
hn_id: 50651470
score: 1
comments: 0
posted_at: "2026-11-26T04:37:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Talira – Search engine for software, built for AI agents

- HN: [50651470](https://news.ycombinator.com/item?id=50651470)
- Source: [talira.ai](https://talira.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-11-26T04:37:30Z

## Translation

タイトル: Show HN: Talira – AI エージェント向けに構築されたソフトウェアの検索エンジン
記事タイトル: タリラ
説明: ソフトウェアの検索エンジン。問題、または製品が何をする必要があり、それにどれくらいのコストがかかるかを説明すると、Talira は適切なツールを見つけます。
HN テキスト: こんにちは、HN、私はネイサンです。今月取り組んでいるプロジェクトを共有したいと思います。数週間前、私の Rust ターゲット フォルダーが多くの git ワークツリー全体で約 70 GB まで増加したため、私はこれを防ぐ方法を見つけるように Claude に依頼しました。オンラインではあまり見つからなかったため、いくつかの不安定な貨物フラグを使用して独自の回避策を構築し始めました。その後、r/rust の誰かが、なぜ kache を使用しなかったのかと尋ねました。他の人が GitHub の問題で Boxington 氏について言及しました (これも素晴らしい解決策です)。また、クロード コード セッションで常に MCP サーバーの再認証を求められることにもイライラしました。これは、多数のサーバーを並行して実行すると発生します。エージェントは適切な「MCP プロキシ」を見つけることができなかったため、独自の「MCP プロキシ」の計画と構築を開始しました。自分で Google を検索して、MCPProxy.app を見つけました。これは非常にうまく機能したので、ブログ投稿を書いて X で共有しました。何人かの人々が返信して、Executor.sh がさらに優れたソリューションであると教えてくれました。彼らの意見は正しかったので、Executor に切り替えました。これらすべてを経て、私は Talira を構築することにしました。これは、AI エージェントが独自のソリューションの構築を開始する前にチェックできるソフトウェアの検索エンジンです。問題をわかりやすい言葉で説明します (「私の Rust ターゲット フォルダーは git ワークツリー全体で 80 GB を使用します」)。 Talira はベクトル埋め込みを使用してカタログを検索し、フィルタリングと再ランキングに Jev を使用します。試してみるには: - https://talira.ai で検索 - MCP サーバーをクロード コードに追加します: claude mcp add --transport http --scope user talira https://mcp.talira.ai/mcp エージェントは各検索後にレポートを送信することもできます: 一般的な問題

rms、試したこと、維持したこと、およびその理由。最終的には、人々が実際に選択し、時間をかけて使い続けるツールのクラウドソーシング データベースになるという考えです。現在はベータ版であり、カタログにはまだ未完成の部分がたくさんあります。また、ChatGPT と Claude はすでに多くのことを知っており、検索が非常に得意であることを私は認識していることにも言及しておきます。したがって、タリラがまったく必要ないケースもたくさんあります。適切にベンチマークを行い、違いが生じるかどうかを確認する前に、さらに実際の検索を収集する必要があります。同じような経験をされた方がいらっしゃいましたらお聞きしたいです。 AI が既存の適切なソリューションを見つけることができなかった場合は、お知らせください。読んでいただきありがとうございます!

記事本文:
Talira はベータ版であるため、一部の検索ではまだ結果が返されません。より多くの人やエージェントが使用するほど、結果は向上します。 AI を接続して、何を保存し、どのツールがタリラに欠けているかを報告させます。開発者ツールを作成しましたか?提出してください。レポートがどのように役立つか
Talira はベータ版であるため、一部の検索ではまだ結果が返されません。 AI を接続してギャップを埋めましょう。
検索 / カタログ ドキュメント スタック リスト 友達 サインイン AI に接続する
ソフトウェアの検索エンジン
AI のために構築されました。問題を説明すると、その問題をすでに解決しているツールが検索されます。
AI を接続する 製品を送信する
無料。その代わりに、あなたの AI はタリラに何を試し、何を保持したかを伝えるので、次の検索は他の人にとってより良いものになります。
{
"問題" : "アプリが送信するメールをキャッチする
開発のため、顧客に届くものはありません」、
"phrasings" : [ "テスト用のローカル SMTP サーバー" ,
「Web 受信箱を備えた偽の SMTP サーバー」 ]、
"探している" : [ "インフラストラクチャ" , "cli" ],
"self_hostable" : true
}
タリラの答えは、最適なものが最初です
Mailpit オープンソース、MIT、セルフホスト可能 0.88
MailCatcher オープンソース、MIT、セルフホスト可能 0.76
5,433 件のリストがあり、解決する 44,936 件の問題によってインデックスが付けられています。
私の Rust ターゲット フォルダーは git ワークツリー全体で 80 GB を使用します
クロード コードが MCP サーバーに再度サインインするように要求し続けます
再デプロイせずに一部のユーザーに対して機能を有効にする
Sentry の SDK と連携するセルフホスト型エラー追跡
何十ものコーディング エージェント セッションを実行しているが、どれが自分を必要としているのか分からなくなってしまう
Postgres データベースを毎晩 S3 にバックアップする
あなたが抱えている問題に基づいて検索してください
各リストには、解決する問題が記載されており、問題を抱えた人が言いそうな方法で書かれています。これらはカタログに書かれているとおりです。
Kache 一度ビルドします。作品を再利用します。
パスワード リセット メールがスパムに分類される 消印 速度を重視したトランザクション メール配信

と配信可能性
更新を取得できなくなった MailHog を置き換えます Mailpit 開発者向けの API を備えた電子メールおよび SMTP テスト ツール
セルフホスト型 Sentry には、12 個のコンテナーと、アプリ全体よりも多くの RAM が必要です。 GlitchTip Sentry SDK と互換性のあるオープンソースのエラー追跡
なぜ別の製品ディレクトリがあるのでしょうか?
ディレクトリは別の質問に答えます。メーカーが自社製品について述べた内容が、カテゴリー、星評価、発売日ごとに並べてリストされています。問題が発生したとき、自分と同じような人にとってどのツールが問題を解決したのか、そしてその人が今もそれを使用しているのかどうかを知りたいと思うでしょう。
問題やカテゴリから探す
この問題を、「すべての git ワークツリーが同じ Rust 依存関係を再コンパイルする」という言い方で説明してください。 Talira は、ビルド キャッシュの存在を知らなかった場合でも、まさにそのために作られたビルド キャッシュである kache を見つけます。どのような種類のツールが必要かすでにわかっていますか? 「稼働時間監視サービス」を検索し、オプションを比較します。
結果は、ツールが問題にどれだけ適合するかによってランク付けされ、次にそのツールが来年メンテナンスされる可能性がどの程度あるかによってランク付けされ、その理由が詳しく説明されます。配置は決して売り物ではなく、メーカー自身の主張が彼らのものとしてラベル付けされます。
研究を行うエージェント向けに構築
AI エージェントは現在、何かを構築するか使用するかを決定していますが、ツールを見逃しています。トレーニング データは数か月前のものであり、最良のヒントは Reddit のスレッドや GitHub のイシューに掲載されており、めったに読まれません。タリラは MCP を通じて彼らに現在の答えを与えます。
AI はすでにツールを検索し、試し、何が機能するかを確認しています。 Talira では結果が数行で共有され、あなたの検索が終了したところから次の人の検索が始まります。
同じ結果が得られる代替案
Argo CD と Talira のリストにある Flux と Fleet の代替案を求めてください。これらは同じ役割を果たします。次に、CI パイプラインで Helm を使用して Git から Kubernetes にデプロイし、デプロイすることで、一度に 1 ステップずつ目標を広げます。

いずれにしても Kubernetes を使用し、Kubernetes を使用せずにアプリをデプロイします。各ステップには何を放棄するかが示されており、Deis Workflow などの停止したツールはなくなったものとしてリストされます。代替手段の仕組み
チームは、Helm でパッケージ化され、Argo CD によって Git からデプロイされたプレーンな Kubernetes オブジェクトなど、複数のツールを同時に使用して目標を達成することがよくあります。一般的なセットアップには、それらの組み合わせがリストされています。
ユーザーが決定すると、エージェントはユーザーが試したツールを使用して、ユーザーが選択した内容とその理由を報告します。これらのレポートは、同様の問題を抱えている次の人のためにツールを上下に移動させますが、ページビューや SEO はまったく関係ありません。誰かが試して落としたツールは、結果として渡されたツールよりも多くカウントされます。 1か月後タリラがさらに重要だと尋ねたとき、まだ使用されている選択肢。
レポートが何も見つからずに終わったり、ユーザーが独自に作成したりしたレポートは、カタログに答えがないことを示しています。問題を抱えているほとんどの人が自分で構築したものを検索すると、そのことがわかります。エージェントが他の場所で見つけたツールをタリラに提案することができ、タリラのエージェントはそれをリストする前にメーカー自身のページと照合してチェックします。
リスト上の事実は、そのページの元のページ (通常はメーカーのページ) と、誰かがそこをチェックした日付とをリンクします。人気に関する主張は彼らの調査を引用しています。Argo CD のリストは、Argo CD が 50%、Flux が 11% である Octopus Deploy の 2025 年の State of GitOps レポートを引用しています。
Talira を検索することで、Talira が構築されるツールを選択します。開発ログには、間違った結果も含め、返された各検索が表示されます。 Argo CD のエントリはミスで始まります。代替案を求められ、タリラは 2 つの GitOps コントローラーと、単語のみが一致する 2 つのツールを挙げました。
1 つの問題を解決するために複数のツールを組み合わせることはよくあるため、レポートではユーザーが選択したすべてのツールに名前を付けることができ、タリラはそれらが一緒に使用されたものとして記録します。誰かる

交流会を開催すると、イベントについては Luma が報告され、メンバーには Facebook グループが報告される可能性があります。レポートには、ユーザーが検討しただけのツールとは別に、ユーザーが試したツールも記録されます。
レポートには、ユーザーがツールをフォークしたことと、フォークによって何が変更されたかが記載されます。ツールをベースにして構築する人は、そのツールをベースにして構築する価値があると感じており、この変更はツールに欠けていたものを示しています。上流に続くフォークは、それぞれが追加するものとともにその下にリストされます。
1 つの検索がどのようにしてみんなの答えになるのか
2026 年 11 月からの実際の問題: MCP サーバーは、数十の並列コーディング セッションにわたって再度サインインを要求し続けました。
エージェントはタリラに、MCP サーバーがサインインを失い続ける理由を尋ねます。カタログには適合するものは何もありませんが、タリラは推測ではなくそう言います。
エージェントは通常どおり Web を検索し、 MCPProxy を見つけます。それは機能したので、エージェントはタリラにレポートを送信します。
数週間後、誰かが Executor を提案します。ユーザーがそこに移動すると、エージェントもそのことを次の理由とともに報告します。
次に問題を自分の言葉で説明する人は、Executor と MCPProxy が先頭にあり、それらを試した人々に何が起こったかを示します。 「独自に構築しました」で終わる検索は、カタログのどこにギャップがあるかを示し、ビルダーがツールを公開してタリラに追加すると、次にアクセスする人全員のギャップを埋めることになります。
AI がレポートを作成します。信用を得られます。
AI は使用するツールの非公開リストを保持し、切り替えると更新します。他の人はそれを見ませんが、ランキングを形成するのに役立ちます。公開プロフィールに表示するツールはユーザーが選択できるため、ユーザーの好みや会社のスタックが公開されることはありません。
投稿に対する評判ではなく、支援に対する評判
レポートが他の人のツール選択に役立ち、そのツールを使い続けると、評判が高まります。
レポートとは、問題、試したツール、維持したツール、およびその理由です。タリラは何もかも拒否する

ハットは秘密または個人データのように見えるため、レポートは公開を選択するまで非公開のままです。
オープンソースのプロジェクトや製品を送信すると、それが解決する問題が表示され、他のものと同じルールでランク付けされます。
AI を接続する 製品を送信する
なぜクロードではなくタリラを使うのでしょうか?
クロードは、他の AI モデルと同様に、トレーニングで学習した内容に基づいてツールを推奨します。これは、何年も変更されていない人気のあるツールに適しています。トレーニング終了後にリリースされたツールについても、記憶しているツールがそれ以降にアーカイブまたは再ライセンスされたことについても知ることができません。 Talira は各ツールの独自のページとリポジトリから構築されたライブ カタログであるため、Claude は何かを推奨する前に何が真実かを調べることができます。
タリラには星評価がありますか?
いいえ。Talira には星による評価やレビューはありません。これは、人々が何を検索したか、どのツールを試したり比較したりしたのか、あるツールを保持したか、落としたか、別のツールに切り替えたか、それとも独自に構築したかなど、結果レポートから学びます。誰かが今でも使用しているツールは、5 点満点中 4 点の評価よりもはるかに多くのことを教えてくれます。
Web サイトでの検索は無料であり、トークンなしの JSON API も同様です。 AI の検索もほとんど同様です。代わりに、AI が何を試し、何を選択したかをタリラに伝え、同じ問題に対する次回の検索を改善します。
企業は上位にランクされるためにお金を払うことができますか?
いいえ、ランキングは決して売り物ではありません。メーカーが提出した製品は、他のすべてのものと同じルールに従ってランク付けされます。
結果レポートには何が記載されますか?
いくつかのフィールド: 一般的な用語での問題、問題がどのように終わったか (そのままにしておく、選択したがまだ開始していない、切り替えた、ドロップした、独自に構築した、適合するものが見つからなかった、または続行しなかった)、理由、試したツール、重さを量っただけのもの、選択したツールとその使用方法 (到着したとき、分岐したとき、または拡張したとき)、およびその背後にある検索

決断。セッションの他の内容は何も入りません。
検索しても何も返されなかったのはなぜですか?
Talira はベータ版であり、そのカタログにはまだ未完成の部分があります。問題に適合するものが何もない場合、その旨を表示し、結果を返さず、エージェントにも Web を検索するように指示します。人々が最終的に何を使用するかについてのレポートは、それらのギャップがどのように埋められるかということを示しています。
タリラは、ツールが問題をどれだけうまく解決できるか、次にメンテナンスされる可能性がどれだけあるかによってツールをランク付けします。
DocSpring 製品。 DocSpring はここにリストされており、他のものと同じルールに従ってランク付けされています。

## Original Extract

Search engine for software. Describe the problem, or what a product has to do and what it can cost, and Talira finds the tools that fit.

Hi HN, I'm Nathan. I wanted to share a project that I've been working on this month. A few weeks ago, my Rust target folders had grown to about 70 GB across many git worktrees, and I asked Claude to figure out how to stop that from happening. It couldn't find much online, so we started to build our own workaround using some unstable cargo flags. Then someone on r/rust asked why we hadn't just used kache. Someone else mentioned Mr Boxington on a GitHub issue (another great solution). I also got annoyed that my Claude Code sessions would constantly ask me to reauthenticate my MCP servers, which happens when you run a lot of them in parallel. An agent couldn't find any good "MCP proxy" so we started planning and building our own. I searched Google myself and found MCPProxy.app. It worked pretty well so I wrote a blog post and shared it on X. Some people replied and told me that Executor.sh was an even better solution. They were right, so I switched to Executor. After all this, I decided to build Talira: A search engine for software that an AI agent can check before it starts building your own solution. You describe the problem in plain words ("my Rust target folders use 80 GB across git worktrees"). Talira searches the catalog using vector embeddings and uses Jev for filtering and reranking. To try it: - Search on https://talira.ai - Add the MCP server to Claude Code: claude mcp add --transport http --scope user talira https://mcp.talira.ai/mcp Your agent can also send a report after each search: the problem in general terms, what you tried, what you kept and why. The idea is that it can eventually become a crowd-sourced database of tools that people actually chose and continued using over time. It's currently in beta and the catalog still has a lot of gaps. I should also mention that I'm aware that ChatGPT and Claude already know a lot and are already very good at search. So there are plenty of cases where you don't need Talira at all. I'll need to collect some more real-world searches before I can benchmark it properly to see if it makes any difference. I'd like to know if you've ever had any similar experiences. Let me know if an AI wasn't able to find a good solution that already existed. Thanks for reading!

Talira is in beta, so some searches won't return any results yet. Results get better as more people and agents use it. Connect your AI and have it report what you kept and which tools Talira is missing. Made a developer tool? Submit it . How reports help
Talira is in beta, so some searches won't return any results yet. Connect your AI to help fill the gaps.
Search / Catalog Docs Stacks Lists Friends Sign in Connect your AI
Search engine for software
Built for your AI. Describe the problem, and it finds the tools that already solve it.
Connect your AI Submit your product
Free. In return, your AI tells Talira what it tried and what you kept, so the next search is better for everyone else.
{
"problem" : "catch the emails my app sends in
development so none reach customers" ,
"phrasings" : [ "local SMTP server for testing" ,
"fake SMTP server with a web inbox" ],
"looking_for" : [ "infrastructure" , "cli" ],
"self_hostable" : true
}
Talira answers, best fit first
Mailpit Open source, MIT, self-hostable 0.88
MailCatcher Open source, MIT, self-hostable 0.76
5,433 listings, indexed by the 44,936 problems they solve.
my Rust target folders use 80 GB across git worktrees
Claude Code keeps asking me to sign in to my MCP servers again
turn a feature on for some users without redeploying
self-hosted error tracking that works with Sentry's SDKs
I run dozens of coding agent sessions and lose track of which one needs me
back up my Postgres database to S3 every night
Search based on the problem you have
Each listing carries the problems it solves, written the way someone with the problem would say them. These come from the catalog as written.
Kache Build once. Reuse the work.
my password reset emails land in spam Postmark Transactional email delivery with a focus on speed and deliverability
replace MailHog, which no longer gets updates Mailpit An email and SMTP testing tool with API for developers
self-hosted Sentry needs a dozen containers and more RAM than my whole app GlitchTip Open source error tracking compatible with Sentry SDKs
Why another product directory?
Directories answer a different question. They list what makers say about their products, sorted by category, star ratings or launch day. When you have a problem, you want to know which tool solved it for someone like you, and whether they still use it.
Search by the problem or the category
Describe the problem the way you'd say it: “every git worktree recompiles the same Rust dependencies”. Talira finds kache , a build cache made for exactly that, even if you didn't know build caches existed. Already know what kind of tool you need? Search for “uptime monitoring services” and compare the options.
Results are ranked by how well a tool fits your problem, then by how likely it is to be maintained next year, with the reasons spelled out. Placement is never for sale, and makers' own claims are labelled as theirs.
Built for the agents doing the research
AI agents now decide whether to build or use something, and they miss tools: their training data is months old, and the best tips sit in Reddit threads and GitHub issues they rarely read. Talira gives them current answers over MCP.
Your AI already searches, tries tools and sees what works. With Talira it shares the outcome in a few lines, and the next person's search starts where yours ended.
Alternatives that get the same result
Ask for alternatives to Argo CD and Talira lists Flux and Fleet, which do the same job. Then it widens the goal one step at a time: deploying to Kubernetes from Git with Helm in a CI pipeline, deploying to Kubernetes any way at all, and deploying apps without Kubernetes. Each step says what you give up, and tools that stopped, such as Deis Workflow, are listed as gone. How alternatives work
Teams often reach a goal with several tools at once, such as plain Kubernetes objects packaged with Helm and deployed from Git by Argo CD. Common setups lists those combinations.
When the user decides, their agent reports what they chose and why, with the tools they tried. Those reports move tools up or down for the next person with a similar problem, and page views and SEO play no part. A tool someone tried and dropped counts against it more than one they passed over in the results. A choice still in use when Talira asks a month later counts for more.
Reports that end with nothing found, or with the user building their own, show where the catalog has no answer. When most people with a problem built their own, a search for it says so. A tool an agent found somewhere else can be suggested to Talira, whose agents check it against the maker's own pages before listing it.
Facts on a listing link the page they came from, usually the maker's own, and the date someone checked them there. Claims about popularity cite their survey: Argo CD's listing quotes Octopus Deploy's 2025 State of GitOps report, which puts Argo CD at 50% and Flux at 11%.
We choose the tools Talira is built with by searching Talira, and the dev log shows each of those searches as it was returned, wrong results included. The Argo CD entry opens with the miss: asked for alternatives, Talira named two GitOps controllers and two tools that only matched words.
People often put several tools together to solve one problem, so a report can name every tool the user chose, and Talira records them as used together. Someone running a meetup might report Luma for the events and a Facebook group for the members. A report also keeps the tools the user tried apart from the ones they only considered.
A report can say the user forked a tool and what the fork changes. Someone who builds on a tool found it worth building on, and the change says what it was missing. Forks that follow their upstream are listed under it, with what each adds.
How one search becomes everyone's answer
A real problem from November 2026: MCP servers kept asking to sign in again across dozens of parallel coding sessions.
The agent asks Talira why its MCP servers keep losing their sign-ins. Nothing in the catalog fits well, and Talira says so instead of guessing.
The agent searches the web as usual and finds MCPProxy . It works, so the agent sends Talira a report:
Weeks later, someone suggests Executor . The user moves to it, and the agent reports that too, with the reason:
The next person who describes the problem, in their own words, finds Executor and MCPProxy at the top, with what happened to the people who tried them. Searches that end in “we built our own” show where the catalog has gaps, and when the builders publish their tool and add it to Talira, it fills the gap for everyone who comes next.
Your AI does the reporting. You get the credit.
Your AI keeps a private list of the tools you use and updates it when you switch. Nobody else sees it, but it helps shape our rankings. You choose which tools to show on your public profile, so it never gives away your preferences or your company's stack.
Reputation for helping, not for posting
You earn reputation when your report helps someone else choose a tool and they keep using it.
A report is the problem, the tools tried, the one kept and why. Talira turns away anything that looks like a secret or personal data, and a report stays private until you choose to make it public.
Submit your open source project or product and it shows up for the problems it solves, ranked by the same rules as everything else.
Connect your AI Submit your product
Why would I use Talira instead of Claude?
Claude, like any AI model, recommends tools from what it learned in training. That works well for popular tools that haven't changed in years. It can't know about a tool released since its training ended, or that a tool it remembers has been archived or relicensed since then. Talira is a live catalog built from each tool's own pages and repository, so Claude can look up what's true now before it recommends something.
Does Talira have star ratings?
No. Talira has no star ratings or reviews. It learns from outcome reports: what people searched for, which tools they tried or weighed, and whether they kept one, dropped it, switched to another or built their own. A tool someone is still using later tells you far more than a rating of 4 out of 5.
Searching on the website is free, and so is the JSON API without a token. Your AI's searches mostly are too: in return, it tells Talira what you tried and what you picked, which improves the next search for the same problem.
Can a company pay to rank higher?
No. Ranking is never for sale. A product its maker submitted ranks by the same rules as everything else.
What goes in an outcome report?
A few fields: the problem in general terms, how it ended (kept, picked and not started yet, switched, dropped, built your own, found nothing that fits, or didn't go ahead), the reason, the tools you tried, the ones you only weighed, the ones you picked and how you use them (as they come, forked or extended), and the searches behind the decision. Nothing else from your session goes into it.
Why did my search return nothing?
Talira is in beta, and its catalog still has gaps. When nothing fits a problem well, it says so, returns no results and tells your agent to search the web as well. Reports of what people end up using are how those gaps get filled.
Talira ranks tools by how well they solve your problem, then by how likely they are to be maintained.
A DocSpring product. DocSpring is listed here and ranked by the same rules as everything else.
