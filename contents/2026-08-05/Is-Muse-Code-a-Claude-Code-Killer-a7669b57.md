---
source: "https://trycodus.com/blog/muse-code-vs-claude-code-benchmarks"
hn_url: "https://news.ycombinator.com/item?id=49188359"
title: "Is Muse Code a Claude Code Killer?"
article_title: "Muse Code vs Claude Code and Codex: what Meta’s own benchmarks actually say · Codus"
author: "feryuk"
captured_at: "2026-08-05T20:57:48Z"
capture_tool: "hn-digest"
hn_id: 49188359
score: 1
comments: 4
posted_at: "2026-08-05T20:14:09Z"
tags:
  - hacker-news
  - translated
---

# Is Muse Code a Claude Code Killer?

- HN: [49188359](https://news.ycombinator.com/item?id=49188359)
- Source: [trycodus.com](https://trycodus.com/blog/muse-code-vs-claude-code-benchmarks)
- Score: 1
- Comments: 4
- Posted: 2026-08-05T20:14:09Z

## Translation

タイトル: ミューズ・コードはクロード・コード・キラーですか?
記事のタイトル: Muse Code vs Claude Code and Codex: what Meta's own Benchmarksactually · Codus
説明: Meta は本日、最初のコーディング エージェントを出荷しました。公開されている 3 つのベンチマーク チャートを読んだところ、Claude Opus 5 は Meta 独自の内部評価を含む 3 つのベンチマーク チャートすべてに勝っていました。ここでは、番号が何をサポートしているのか、アスタリスクがどこにあるのか、そして誰が実際にインストールする必要があるのか​​を示します。

記事本文:
Muse コード vs Claude Code および Codex: Meta 独自のベンチマークが実際に示していること · Codus codus 製品 仕組み 価格 ブログ サインイン ダウンロード Mac 用ダウンロード ← すべての投稿 エンジニアリング Muse Code vs Claude Code and Codex: Meta 独自のベンチマークが実際に示していること
Meta は今朝最初のコーディング エージェントを出荷しましたが、1 時間以内に私のフィードはそれがクロード コード キラーかナッシングバーガーのどちらかであるとすでに判断していました。どちらの評決も、メタが同時に公開したベンチマークチャートを誰も開く前に書かれていたため、代わりに午前中をかけてそれらを読んだところ、どちらの判断よりもかなり興味深い数字が判明した。
短いバージョン: Meta は 3 つのベンチマーク チャートを公開しましたが、Claude Opus 5 は、Meta が独自に構築したベンチマークを含む 3 つのベンチマーク チャートすべてで 1 位になっています。これは発射ポストとしては実に異例のことであり、そのことについて言及した報道はほとんどありませんでした。しかし、より有益な発見はその下にあるものです。Muse Code が初日のターミナルベンチで Codex を、トークン価格の約 4 分の 1 で破りました。これにより、これはリーダーへの挑戦ではなく、2 位をめぐる真剣な戦いになります。
ここに、Meta が発表したすべての内容、数値が実際にサポートしている内容、アスタリスクの場所、およびインストールする必要があるかどうかを示します。
Muse Code は、Muse Spark 1.2 と呼ばれる新しいモデルを搭載した、macOS および Linux 上のベータ版のターミナル コーディング エージェントです。これは、Alexandr Wang が運営するグループである Meta Superintelligence Labs が開発した最初のコーディング エージェントであり、Mark Zuckerberg はそれについて明白に説明しました。
これは、変更の計画、コードの作成、結果の検証など、大規模なリポジトリ全体にわたる完全なソフトウェア エンジニアリング タスクを引き受けるターミナル コーディング エージェントです。
これは 1 行でインストールできます。これは、これまでに CLI エージェントをセットアップしたことのある人には馴染みのあるものでしょう。

e:
カール -fsSL https://dev.meta.ai/install.sh |バッシュ
アーキテクチャは二度読む価値のある部分です。
永続的なバックグラウンド エージェント。 Muse コードは、タスクごとに新しいサブエージェントを生成して破棄するのではなく、セッション全体で非同期エージェントを存続させ、作業中にコンテキストを蓄積します。
分離されたワークツリー内の並列サブエージェント。タスクが十分に大きい場合、タスクは同時に動作するサブエージェントに分割され、それぞれが独自の git worktree 内で動作するため、相互に上書きすることができません。
ローカルイベントログ。すべてのモデルの呼び出し、ツールの実行、承認、編集がディスクに書き込まれるため、正確な再生と再起動が安全なランタイムが実現します。実行中にクラッシュした場合は、セッションが再開されます。
3つのスキルがバンドルされています。 /plan は承認ゲート型の計画を生成し、/grill はその計画のストレス テストを行い、/goal は指定された目標が達成されるまで実行されます。
Muse Spark 1.2 自体は、個別に出荷して後でラップするのではなく、エージェントと共同トレーニングされ、リポジトリ全体の生成を含む長期的な作業についてトレーニングされ、Spark 1.1 が 1.2 用に学習するハード コーディング環境を生成するループを通じて改善されました。特に、デスクトップアプリはありません。 Claude Code や Codex とは異なり、これは今のところ端末のみです。
Meta は 3 つのベンチマーク チャートを公開しました。私が読んだ見出しはどれも、その中からお世辞のような文章を選んでいた。その記事は、Muse Spark 1.2 が Grok Build 4.5 と Gemini 3.6 Flash を破ったというものだった。それは本当です。これはページ内で最も興味のないものでもあります。
これらのツールの実際の動作に最も近いベンチマークである Terminal-Bench 2.1 から始めます。
Muse Code は 82.9% で 2 位、Opus 5 の Claude Code の 86.7% に次ぎ、GPT 5.6 Terra の Codex の 81.8% を上回っています。初日ベータでのリーダーとの 3.8 ポイント差は実際の結果であり、たとえ 1 ポイント差であってもコーデックスを破ったことが見出しとなったメタは獲得したが、勝ち残れなかった

で。
次に、長期的なエージェント作業を測定する DeepSWE 1.1 では、状況が変わります。
ここでは、Muse Code が 59.3% で 3 位に下がり、Claude Code の 65.0% と Codex の 64.8% が実質的にトップで並んでいます。対象期間が長ければ長いほど、Muse コードはさらに後れをとっていきます。これは、ベンチマークが明らかにするように設計されているものと一致しており、エージェントに 20 分の仕事ではなく一晩の仕事を任せることを計画している場合には、非常に重要になります。
そして、Meta が公開するとはまったく予想していなかった、独自の内部コーディング ベンチマークのグラフが表示されました。
Meta が社内で設計、キュレーション、実行したベンチマークでは、Muse Spark 1.2 の 70.6% に対して、Claude Opus 5 のスコアは 79.4% でした。あなた自身の立ち上げ投稿で自発的に公表されたあなた自身の評価での 8.8 ポイントの差は、驚くべき率直さ、またはフロンティアがまだどこかにあることを静かに認めていることのどちらかです。私は率直な意見を重視しており、比較を行わないベンダーよりも、負けたチャートを公開してくれるベンダーを希望します。
知っておく価値のある 3 つのアスタリスク
ベンチマークは、比較セットを選択した人に報酬を与えるため、これを確定したものとして扱う前に、3 つの点にフラグを立てる必要があります。
メタはソルではなくテラに対してベンチマークを行いました。これらのグラフの OpenAI バーはすべて、フラッグシップの GPT 5.6 Sol ではなく、中間層モデルの GPT 5.6 Terra です。 Hacker News の開発者は数分以内にこれを発見し、ある開発者は Meta 自身の GPU カーネルのケーススタディでは Sol が優位に立っているようだと指摘しました。したがって、上記の Codex の数値は、Muse Code が隣に立つことを選択した数値であり、入手可能な最強の数値ではありません。
これらはエージェントとモデルのスコアであり、モデルのスコアではありません。各バーの下の小さな文字に注目してください。Opus 5 は Claude Code 内で実行され、GPT 5.6 Terra は Codex 内で、Spark 1.2 は Muse Code 内で実行されました。これはおそらく製品を比較するための正直な方法です。なぜなら、それが実際に製品を比較する方法だからです。

ただし、これはこれらを純粋なモデルのランキングとして読み取ることはできず、ハーネスがすべてのバーで機能の一部を行っていることを意味します。
1.1 から 1.2 へのジャンプの一部はハーネスの変更です。 Terminal-Bench では、Spark 1.1 は、意図的に最小限のハーネスである mini-swe-agent での実行で 76.2% のスコアを獲得しましたが、1.2 は専用の Muse コード内での実行で 82.9% のスコアを獲得しました。この 6.7 ポイントの改善は、モデルのゲインとツールのゲインを加算したものであり、グラフではそれらを分離していません。私が最も単独で再放送してほしい数字です。
開発者の反応は、息切れではなく、それに応じて測定されています。 Hacker News のスレッドは「素晴らしいリリースと Spark 1.1 に対する確かな改善。SOTA ではないが堅実」ということで落ち着いており、1.1 から 1 か月も経たないうちに 1.2 が登場することについては懐疑的な側面があり、開発者がコードベースのライブ フィードを Meta に渡したいかどうかという完全に予測可能な質問が繰り返されていました。
Muse コードは、入力トークン 100 万あたり 1.25 ドル、出力トークン 100 万あたり 4.25 ドルで実行されます。これはフロンティア価格よりも約 4 対 1 安く、開始するには 20 ドルの無料クレジットがあります。エージェントのコーディングでは、四捨五入誤差ではなく 100 万あたりの価格設定が主要な項目となるレートでトークンが消費されるため、多くのチームにとってその違いが議論のすべてです。
次に、コントリビューター層があります。Meta 氏によれば、これはさらに 10 倍以上安いので、ここで私はペースを落とします。その価格は割引ではなく、取引です。モデルを改善するために使用法を選択することになります。 Meta はこれをコーディング エージェントの業界標準だと説明しており、これは公平であり、要求に応じてデータ保持ゼロが利用可能であり、Wang 氏はこれを重要なエンタープライズ機能と呼んでいます。しかし、「リクエストに応じて利用可能」という文は、この文で多くの役割を果たしており、誰にとってもデフォルトの姿勢です。

NDA に基づいてクライアント コード、ライセンスされたソース、その他のものに触れる場合は、安価な層は自分たちには向かないと想定する必要があります。
平たく言えば、4 倍安い価格で、真に競争力のある代理店を実質割引価格で購入することになります。リポジトリで支払うと 40 倍安くなります。
では、それはクロード・コードやコーデックスに対する脅威なのでしょうか？
この証拠とメタ自身のチャートに基づいて、私がヘッジなしでそう言える理由はクロード・コードではありません。 Claude Opus 5 は Terminal-Bench でリードし、DeepSWE でリードし、Meta の社内ベンチマークでもリードしており、タスクが長くなるにつれてその差は拡大し、それがカテゴリー全体の方向に向かっています。あなたの仕事が長期にわたる、複数のファイルにまたがる、引き渡して立ち去るエンジニアリングである場合、ここで設定を変更する必要はありません。
しかし、コーデックスにとっては、それは本当にその通りだと思います。 Muse Code は、ベータ版のターミナル ベンチと同等の価格で初日に 4 分の 1 の価格で提供され、中間層の価格設定がもたらす中規模の並列作業にまさに適したアーキテクチャでそれを実現しました。 OpenAIの答えはおそらくTerraではなくSolだろうが、今では性能だけでなく価格にもプレッシャーがかかっており、Metaは異例のほど潤沢な資金力を持った戦いとなっている。
今日誰がインストールすべきかについての私の率直な意見は次のとおりです。
中程度の複雑さの作業を大量に実行し、コストを重視し、トレーニング セットで見ても問題ないと思われるコードに取り組んでいる場合に試してください。
クロスチェックのためにローテーションに 2 番目のエージェントが必要な場合は、これを見てください。差分について 2 人のエージェントが意見を異にすることは、本当に有益なシグナルであり、安価なセカンドオピニオンは、安価なモデルの有効な利用法です。
作業が長期間にわたって無人で行われる場合、クライアントまたは規制されたコードを処理する場合、または GUI が必要な場合は、これはターミナルのみでアプリは含まれないため、今のところスキップしてください。
私が最も興味を持っている部分
ベンチを設置する

Meta の立ち上げ記事で最も重要なことは数字ではないため、マーク表はしばらく脇に置いてください。それは、Meta が独立して事実上無制限のリソースを使用して、私たちが行ったのと同じアーキテクチャに到達したということです。つまり、使い捨てエージェントではなく永続エージェント、相互に破損しないように独自のワークツリー内に分離された並列サブエージェント、およびクラッシュした実行が再起動せずに再開される耐久性のあるローカル イベント ログです。
私たちはまさにその形状に基づいてコーダスを構築しているため、私はここでは中立的な観察者ではありません。しかし、Anthropic、OpenAI、そして現在の Meta が 1 年以内にすべて同じ答えに収束すると、それはツールに関する 1 つのベンダーの意見ではなくなり、この分野で定着したアーキテクチャになり始めます。チャット ウィンドウ内に 1 人のエージェントがいる時代は終わり、現在、本格的に参入する企業はすべてフリートを出荷しています。
本当に未解決のままであり、そのページのベンチマークが測定していないことは、これらの並行エージェントの意見が異なる場合に何が起こるかということです。ワークツリーは、相互にファイルを上書きすることを防ぎます。同じ機能の互換性のない 2 つの部分を構築することを阻止するためには何もしません。 Meta の答えは /grill です。これは、作業を開始する前に計画のストレス テストを行うエージェントです。それは合理的な答えであり、次のラウンドの競争はターミナルベンチではなくそこで行われるのではないかと思います。
上記のすべての数値は、メタ社自身の発表のグラフから引用したもので、二次報道ではなく直接読んだものです。これは、この発表に関するレポートで、各数値がどのベンチマークに属するかについて一貫性がないためです。
Meta AI Research: Muse Code と Muse Spark 1.2 の紹介 (主要なソースと 3 つのチャートすべての起源)
CNBC: Meta が Anthropic と OpenAI に対抗する Muse Code をデビュー
9to5Mac: Meta が macOS および Linux 用の Muse Code を発表
ハッカー ニュース: Muse コードと Muse Spark 1.2 のディスカッション
Seeking Alpha: メタレル

AI コーディング エージェント Muse Code を容易にする
ベンチマーク チャートは、解説と比較のために Meta の発表から転載したものです。図は、2026 年 8 月 5 日に Meta によって公開されたものであり、私を含めて独自に複製されていません。
それについて読むのではなく、見てください。
Codus は macOS 上でプライベートベータ版です。毎週金曜日に新しい招待状が届きます。
Mac 用のダウンロード codus コードエディター、ビデオエディター、SEO、および詳細な調査が 1 つの Mac アプリに含まれており、自律エージェントの乗組員がそれぞれを並行して実行します。
codus は、イングランドおよびウェールズで登記された会社である AI Search Labs Ltd (会社番号 16719803)、登録事務所 Windrush House, Windrush Park Road, Witney, OX29 7DX, United Kingdom が所有しており、イングランドおよびウェールズで登記された会社である Search Intelligence Ltd (会社番号 09361526、VAT 番号 GB 357 3329 84) のライセンスに基づいて運営されています。登録事務所 Witney Business and Innovation Centre, Windrush Park Road, Brighthampton, Witney, OX29 7DX, United Kingdom。 Search Intelligence Ltd は登録販売者であり、お客様が契約する会社です。法律および企業情報の詳細。
© 2026 コーダス。無断転載を禁じます。

## Original Extract

Meta shipped its first coding agent today. I read the three benchmark charts it published, and Claude Opus 5 wins all three, including Meta’s own internal eval. Here is what the numbers support, where the asterisks are, and who should actually install it.

Muse Code vs Claude Code and Codex: what Meta’s own benchmarks actually say · Codus codus Products How it works Pricing Blog Sign in Download Download for Mac ← All posts engineering Muse Code vs Claude Code and Codex: what Meta’s own benchmarks actually say
Meta shipped its first coding agent this morning, and within the hour my feed had already decided it was either a Claude Code killer or a nothingburger. Both verdicts were written before anyone had opened the benchmark charts Meta published alongside it, so I spent the morning reading those instead, and the numbers turn out to be considerably more interesting than either take.
The short version: Meta published three benchmark charts, and Claude Opus 5 comes first on all three , including the benchmark Meta built itself. That is a genuinely unusual thing for a launch post to do, and almost none of the coverage mentioned it. But the more useful finding is the one underneath: Muse Code beat Codex on Terminal-Bench on day one, at roughly a quarter of the token price, which makes this a serious fight for second place rather than a challenge to the leader.
Here is everything Meta announced, what the numbers actually support, where the asterisks are, and whether I think you should install it.
Muse Code is a terminal coding agent, in beta, on macOS and Linux, powered by a new model called Muse Spark 1.2 . It is the first coding agent out of Meta Superintelligence Labs, the group run by Alexandr Wang, and Mark Zuckerberg described it plainly enough:
It’s a terminal coding agent that takes on complete software engineering tasks across large repos: planning changes, writing code, validating the results.
You install it with a single line, which will look familiar to anyone who has set up a CLI agent before:
curl -fsSL https://dev.meta.ai/install.sh | bash
The architecture is the part worth reading twice:
Persistent background agents. Rather than spawning a fresh subagent per task and throwing it away, Muse Code keeps async agents alive across the session so they accumulate context as you work.
Parallel sub-agents in isolated worktrees. When a task is large enough, it splits into sub-agents that work simultaneously, each in its own git worktree so they cannot overwrite each other.
A local event log. Every model call, tool run, approval and edit is written to disk, which gives it a replay-exact and restart-safe runtime. If it crashes mid-run, it picks the session back up.
Three bundled skills. /plan produces an approval-gated plan, /grill stress-tests that plan, and /goal runs until a stated objective is met.
Muse Spark 1.2 itself was co-trained with the agent rather than shipped separately and wrapped afterwards, trained on long-horizon work including whole-repository generation, and improved through a loop where Spark 1.1 generated hard coding environments for 1.2 to learn on. Notably, there is no desktop app. Unlike Claude Code and Codex, this is terminal-only for now.
Meta published three benchmark charts. Every headline I read picked the flattering sentence out of them, the one saying Muse Spark 1.2 beats Grok Build 4.5 and Gemini 3.6 Flash. That is true. It is also the least interesting thing on the page.
Start with Terminal-Bench 2.1, the benchmark that most closely matches what these tools actually do:
Muse Code takes second at 82.9% , behind Claude Code on Opus 5 at 86.7% , and ahead of Codex on GPT 5.6 Terra at 81.8% . A 3.8-point gap to the leader on a day-one beta is a real result, and beating Codex, even by a point, is the headline Meta earned but did not lead with.
Then DeepSWE 1.1, which measures long-horizon agentic work, and where the picture changes:
Here Muse Code drops to third at 59.3% , with Claude Code at 65.0% and Codex at 64.8% effectively tied at the top. The longer the horizon, the further Muse Code falls behind, which matches what the benchmark is designed to expose, and matters enormously if you plan to hand an agent an overnight job rather than a twenty-minute one.
And then the chart I did not expect Meta to publish at all, its own internal coding benchmark:
On a benchmark Meta designed, curated and ran in-house, Claude Opus 5 scores 79.4% against Muse Spark 1.2’s 70.6% . An 8.8-point deficit on your own eval, published voluntarily in your own launch post, is either remarkable candour or a quiet admission that the frontier is still somewhere else. I lean towards candour, and I would rather have a vendor that publishes the chart it loses than one that does not run the comparison.
Three asterisks worth knowing about
Benchmarks reward whoever chooses the comparison set, so before treating any of this as settled, three things deserve flagging.
Meta benchmarked against Terra, not Sol. Every OpenAI bar in those charts is GPT 5.6 Terra, the mid-tier model, rather than GPT 5.6 Sol, the flagship. Developers on Hacker News spotted this within minutes, and one pointed out that in Meta’s own GPU kernel case study, Sol appears to come out ahead. So the Codex numbers above are the ones Muse Code chose to stand next to, not the strongest ones available.
These are agent-and-model scores, not model scores. Look at the small print under each bar: Opus 5 ran inside Claude Code, GPT 5.6 Terra inside Codex, Spark 1.2 inside Muse Code. That is arguably the honest way to compare products, because it is how you will actually use them, but it means you cannot read these as pure model rankings, and the harness is doing part of the work in every bar.
Part of the 1.1 to 1.2 jump is a harness change. On Terminal-Bench, Spark 1.1 is scored at 76.2% running on mini-swe-agent, a deliberately minimal harness, while 1.2 is scored at 82.9% running inside the purpose-built Muse Code. That 6.7-point improvement is a model gain and a tooling gain added together, and the chart does not separate them. This is the number I would most want to see re-run independently.
Developer reaction has been correspondingly measured rather than breathless. The Hacker News thread settled around "a nice release and a solid improvement over Spark 1.1, not SOTA, but solid", with a side of scepticism about a 1.2 arriving less than a month after 1.1, and a recurring, entirely predictable question about whether developers want to hand Meta a live feed of their codebase.
Muse Code runs at $1.25 per million input tokens and $4.25 per million output tokens , which undercuts frontier pricing by roughly four to one, and there is $20 of free credit to start. For a lot of teams that difference is the whole argument, because agentic coding burns tokens at a rate that makes per-million pricing the dominant line item rather than a rounding error.
Then there is the contributor tier, which Meta says is more than ten times cheaper again , and this is where I would slow down. That price is not a discount, it is a trade: you opt in to your usage improving the model. Meta describes this as industry-standard for coding agents, which is fair, and zero-data-retention is available on request, which Wang has called an important enterprise feature. But "available on request" is doing a lot of work in that sentence, and the default posture for anyone touching client code, licensed source or anything under NDA should be to assume the cheap tier is not for them.
Put plainly: at four times cheaper you are buying a genuinely competitive agent at a real discount. At forty times cheaper you are paying with your repository.
So is it a threat to Claude Code or Codex?
Not to Claude Code, on this evidence, and Meta’s own charts are the reason I can say that without hedging. Claude Opus 5 leads on Terminal-Bench, leads on DeepSWE, and leads on Meta’s in-house benchmark, and the gap widens precisely as tasks get longer, which is the direction the whole category is moving. If your work is long-horizon, multi-file, hand-it-over-and-walk-away engineering, nothing here should change your setup.
To Codex, though, I think it genuinely is. Muse Code arrived at parity on Terminal-Bench, in beta, on day one, at a quarter of the price, and it did so with an architecture that is well suited to exactly the medium-sized parallel work that mid-tier pricing invites. OpenAI’s answer will presumably be Sol rather than Terra, but the pressure is now on price as much as capability, and that is a fight Meta is unusually well funded to have.
My honest read on who should install it today:
Try it if you run high volumes of medium-complexity work, are cost-sensitive, and are working on code you would be comfortable seeing in a training set.
Watch it if you want the second agent in your rotation for cross-checking. Two agents disagreeing about a diff is a genuinely useful signal, and a cheap second opinion is a good use of a cheap model.
Skip it for now if your work is long-horizon and unattended, if you handle client or regulated code, or if you need a GUI, because this is terminal-only, with no app.
The part that interests me most
Set the benchmark table aside for a moment, because the most significant thing in Meta’s launch post is not a number. It is that Meta, independently and with effectively unlimited resources, arrived at the same architecture we did: persistent agents rather than disposable ones, parallel sub-agents isolated in their own worktrees so they cannot corrupt each other, and a durable local event log so a crashed run resumes instead of restarting.
We build codus on exactly that shape, so I am hardly a neutral observer here. But when Anthropic, OpenAI and now Meta all converge on the same answer within a year, it stops being one vendor’s opinion about tooling and starts being the settled architecture of the field. The single-agent-in-a-chat-window era is over, and every serious entrant is now shipping a fleet.
What remains genuinely unsolved, and what no benchmark on that page measures, is what happens when those parallel agents disagree. Worktrees stop them overwriting each other’s files; they do nothing to stop them building two incompatible halves of the same feature. Meta’s answer is /grill, an agent that stress-tests a plan before work starts. That is a reasonable answer, and I suspect the next round of competition is fought there rather than on Terminal-Bench.
Every number above comes from the charts in Meta’s own announcement, read directly rather than via secondary coverage, because the reporting on this launch has been inconsistent about which benchmark each figure belongs to.
Meta AI Research: Introducing Muse Code and Muse Spark 1.2 (primary source, and the origin of all three charts)
CNBC: Meta debuts Muse Code to take on Anthropic and OpenAI
9to5Mac: Meta launches Muse Code for macOS and Linux
Hacker News: Muse Code and Muse Spark 1.2 discussion
Seeking Alpha: Meta releases AI coding agent Muse Code
Benchmark charts are reproduced from Meta’s announcement for commentary and comparison. Figures are as published by Meta on 5 August 2026 and have not been independently reproduced, including by me.
See it instead of reading about it.
Codus is in private beta on macOS. New invites every Friday.
Download for Mac codus A code editor, a video editor, SEO and deep research in one Mac app, with a crew of autonomous agents running each of them in parallel.
codus is owned by AI Search Labs Ltd , a company registered in England and Wales (company number 16719803), registered office Windrush House, Windrush Park Road, Witney, OX29 7DX, United Kingdom, and is operated under licence by Search Intelligence Ltd , a company registered in England and Wales (company number 09361526, VAT number GB 357 3329 84), registered office Witney Business and Innovation Centre, Windrush Park Road, Brighthampton, Witney, OX29 7DX, United Kingdom. Search Intelligence Ltd is the merchant of record and the company you contract with. Full detail on Legal & corporate information .
© 2026 Codus. All rights reserved.
