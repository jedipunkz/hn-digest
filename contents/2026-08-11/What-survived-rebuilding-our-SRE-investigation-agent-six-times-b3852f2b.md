---
source: "https://engineeringexec.tech/posts/evolution-of-an-sre-agent/"
hn_url: "https://news.ycombinator.com/item?id=49258133"
title: "What survived rebuilding our SRE investigation agent six times"
article_title: "The Evolution of an SRE Agent: From Five Browser Tabs to a Slack Command | Blog"
author: "mbleterman"
captured_at: "2026-08-11T14:14:02Z"
capture_tool: "hn-digest"
hn_id: 49258133
score: 1
comments: 0
posted_at: "2026-08-11T13:38:12Z"
tags:
  - hacker-news
  - translated
---

# What survived rebuilding our SRE investigation agent six times

- HN: [49258133](https://news.ycombinator.com/item?id=49258133)
- Source: [engineeringexec.tech](https://engineeringexec.tech/posts/evolution-of-an-sre-agent/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T13:38:12Z

## Translation

タイトル: SRE 調査エージェントを 6 回再構築しても生き残ったもの
記事のタイトル: SRE エージェントの進化: 5 つのブラウザ タブから Slack コマンドへ |ブログ
説明: すべてのアラートは同じ税金から始まります。5 つの可観測性システムのどれが真実を伝えているのか、そして私がいるのは 2 つの実稼働世界のどれなのか?これは、SRE アラート調査エージェントが、カーソルおよび MCP クエリからリモートの Slack トリガー調査に至るまで、6 つのバージョンにわたってどのように進化したか、そしてその理由は次のとおりです。
[切り捨てられた]

記事本文:
SRE エージェントの進化: 5 つのブラウザ タブから Slack コマンドへ |ブログ EngineeringExec
EngineeringExec へようこそ。ここではコーディングの専門知識が集まります。
経営陣のリーダーシップ。
私は Michael Bleterman です。R&D リーダー兼ソフトウェア アーキテクトです。
20年以上の経験があります。
私の使命は、エンジニアリングとマネジメントの間の橋渡しをすることです。
© 2026 マイケル・ブレターマン。無断転載を禁じます
SRE エージェントの進化: 5 つのブラウザ タブから Slack コマンドへ
すべての SRE は税金について知っています。アラートを修正する前に、問題とは関係のない 3 つの質問に答える必要があります。つまり、どのシステムが真実を伝えているのか、どの環境を見ているのか、これを最後に見た人はどこにそれを書き留めたのかということです。天気の良い日なら10分程度で済みます。午前2時には生きる意志が失われる。
ティパルティでは税金が他のところよりも重い。私たちは 2 つの運用グレードの環境を実行しており、それぞれが Windows マシン上のレガシー モノリスと K8 上のマイクロサービス群に分割されており、決して混同してはなりません。彼らのシグナルは、Datadog、Coralogix、Prometheus、Pingdom、およびアクセス制御の背後にある一部の内部モニタリングに分散されています。バニラ LLM をそこに向けて、アラートの原因を尋ねると、コンテキストなしですべてが行うことと同じことを実行します。自信を持って何かをでっちあげます。
これは、最初の税金をどのように自動化したかの物語です。 1 つの賢いビルドではなく、そのうちの 6 つで、そのほとんどを捨てました。ランタイムは 4 回変更されました。生き残ったのはそれ以外のすべてだった。
バージョン 0: 人間、5 つのタブ、大量のスクロール
2025 年以前にはエージェントは存在せず、待機しているのは誰でもありました。アラートを取得し、一致する可観測性システムを開き、スレッドのプルを開始します。ログがここにあり、トレースがそこにあり、メトリクスは間違っているように見え、ダッシュボードは正常に見えます。アラートがモノリスから来た場合は、あなたはそう思います

1 つのツール セットと 1 つのメンタル モデルを検討します。 K8s のものであれば、まったく異なるセットになります。どの世界にどのツールを使用するかという知識は人々の頭の中にあり、誰かがチームを変えるたびにそれが消えていきました。うまくいきました。拡張できず、スリープ状態にもなりませんでした。
2025 年までに、数学は機能しなくなります。チームはよりスリムになり、オンコールのローテーションはより重くなりましたが、新しいプラットフォーム機能が追加されるたびに、理解する必要がある領域が広がりました。私たちが保持する必要があるコンテキストは、チームよりも速く成長しました。その場に留まるために全力疾走するレッドクイーンレースです。対応を続けるということは、トリアージを半自動化し、さらに自動化することを意味しました。
2025 年後半が最初の本当のステップでした。 MCP サーバーを介していくつかの可観測性システムを Cursor に接続したため、5 つの UI をクリックする代わりに、1 か所で特定の時間枠のログとトレースを要求し、数秒でそれらを取得できるようになりました。 PromQL や DataPrime など、誰かがあなたの皿を取り上げたような風変わりなクエリ言語について考えてみましょう。
これはスピードの勝利であり、それ以上のものではありませんでした。人間は依然としてすべての結果を読み、次に何をクエリするかを決定し、依然として 2 つの世界の地図を頭の中に保持しています。しかし、「タブを 5 つ開いてスクロールする」を「1 回だけ質問する」に変えることで、調査のリズムが変わりました。機械が取り出しを行い、人間が考えるだけを行ったのは初めてのことだった。
クエリをスキルでラップする
年末までに、生のクエリには問題が発生しました。誰もがクエリの書き方が少しずつ異なり、運用環境を調査する際にサンドボックス データを読み取る方法も少し異なりました。そこで、一般的なクエリをチームレベルのスキルにまとめました。スキルは、どの環境に触れることを許可されているか、どのフィールドが重要であるか、そして単にもっともらしい答えを返すのではなく正確な答えが返されるように質問をどのように表現するかを知っていました。
精度が上がり、

そして、誰が運転しているかによって結果が停止しました。フローはまだ手動で、人がスキルを順番に呼び出していましたが、スキルには以前は人々の頭の中に存在していたコンテキストが含まれるようになりました。
スタンドアロンのエージェントが壁にぶつかる
それから私たちは野心的になりました。 2026 年の初めに、私たちは独立したエージェントとして全体を立ち上げようとしました。これは、人のラップトップ IDE 内に乗るのではなく、人ではなくシステムによってトリガーされるのではなく、単独で実行できるものです。ホストとして AWS AgentCore を評価しました。
自分でハンドブレーキを引いた。この考えは間違っていませんでしたが、クラウドでホストされるエージェント ランタイムは、クラウド ネイティブでクラウドに到達可能なものに到達するために構築されています。私たちの世界はその逆です。2 つのゲート環境、Windows モノリス、他人のクラウドで実行されているものに自らをさらさないツールです。私たちはエージェントを必要な環境に橋渡しするのに数週間を費やしましたが、橋は崩壊し続けました。私がそれをやめさせた理由は、難易度ではなく、作業の形状でした。スコープが環境の配管に忍び込み、短距離または中距離に MVP がありませんでした。賭けの対象が、出荷できるものに向けられなくなり、インフラストラクチャ プロジェクトに向けられるようになったら、それは無効になります。レッスンは高価でしたが、それだけの価値がありました。この問題では、エージェントがどこで実行されるかは詳細ではありません。それはゲーム全体です。
SRE Investigator: 残りを推進する 1 つのスキル
2026 年の春、私たちはエディターから逃れようとするのをやめ、代わりにエディターを倍増させました。 Cursor に戻って、私たちは SRE Investigator と呼ばれるスキルの最初のバージョンを構築し、それがすべての形を変えました。それまでスキルは人間が一つ一つ身に付けてきた道具でした。 Investigator は、他のスキルを推進する役割を持つ最初のスキルでした。つまり、アラートを受け取り、それがどの世界に属しているかを判断し、適切なログを取得して、

正しい順序でメトリクスを見つけ、それらを関連付けて、根本原因の最初の草案を返します。
人間は調査の推進から結論の検討までを行った。それは別の仕事であり、はるかに高速です。
クロード・コード、そして最後にスコアを記録する
6月までに私たちは再び編集者を超えてしまいましたが、これには正当な理由がありました。私たちはすべてを Claude Code 上のエージェント駆動ハーネスに移植し、もっと早くに構築すべきだったもの、つまりスコアを維持する方法を構築しました。バックテスト フレームワークを使用すると、エージェントを通じて実際の過去のアラートを再生し、それが正しいとわかっている根本原因に到達したかどうかを確認できます。これにより、変更すべき数値が得られました。以降のバージョンでは、変更が役に立ったかどうかを推測するのではなく、エージェントの再生アラートの精度を約 80% に引き上げました。
これはデモビデオに誰も入れていない部分であり、最も重要な部分です。ほとんどの場合正しく、残りの場合は自信を持って間違っている捜査機関は、捜査機関がまったくいないよりも悪いです。なぜなら、それは人々にそれを信頼するように教えているからです。そして、誤差は対称的ではありません。誤検知により、オンコール エンジニアが間違った方向に誘導されると、解決までの時間が 30 分かかる可能性があります。誤検知はエンジニアの夜を無駄にし、翌日に損害を与えます。ここでの精度は、リーダーボードのスコアではなく、顧客の滞在時間と睡眠不足で測定されます。ハーネスのおかげで、システムを変更し、出荷前にシステムをよりスマートにしたのか、単に異なるものにしたのかを知ることができます。
ラップトップ以外: Slack コマンド
最新の転生である 2026 年 7 月は、最終的に個人を置き去りにしたものです。調査員は現在、Slack からトリガーされて claude.ai でリモートで実行されています。適切な人が、適切なラップトップを使用し、適切な個人認証情報を持ち、適切な時間に起きている必要はもうありません。誰かがアラートを投稿、エージェント

それを拾い上げ、以前はすべてのインシデントの最初の 20 分を費やしていたトリアージを実行し、人間がまだコーヒーに手を伸ばしている間に、根本原因の最初の草案をポストバックします。
その能力は私たちの一部だけができるものではなくなり、チームが持つものになりました。現在処理しているアラート全体で、解決までの平均時間は約 30% 短縮されており、ほとんどのインシデントは最初の根本原因が既に関連付けられている人間に到達しています。
バックグラウンドでは、ベンダーのトライアル版、つまりすべての作業をすぐに実行できると約束されたツールがローテーションで実行されていました。彼らのほとんどは同じ岩で失敗しました。彼らは、私たちの世界が決して混在してはいけないということをまったく知りませんでした。どのシステムを信じるべきかも知りませんでした。答えを流暢ではなく信頼できるものにするコンテキストを獲得する方法もありませんでした。既製のものではモノリスとマイクロサービスを区別できませんでしたが、ここではその違いがすべてです。
私たちはランタイムを 4 回変更しました。ブラウザーのタブ、カーソル、AgentCore での行き止まりの実行、Claude Code、リモート サービスです。それらのいずれかにプロジェクトを賭けていたら、そのランタイムが適合しなくなったときに、そのプロジェクトを捨てていたでしょう。私たちが決して捨てなかったのはコンテキストです。どの世界に触れる可能性があるかを知るスキル、正確に返されるクエリ、出荷前に回帰を検出するバックテストです。モデルは決して堀ではありませんでした。ランタイムは決して堀ではありませんでした。という文脈でした。
次は何ですか: マルチエージェント システム
一人の捜査官はやはり一人の捜査官だ。次のバージョンはマルチエージェント フレームワークで、現在のフレームワークにはないいくつかの機能をサポートしています。メモリです。そのため、システムは先月すでに発見した根本原因の再導出を停止し、「これを確認しました」と言い始めます。また、独自の事後分析、ランブック、過去の調査を RAG するため、回答を信頼できるものにするコンテキストが 1 つのプロンプトに詰め込まれることなく成長します。
あなたが b なら

同様のものを構築する場合は、最初に完璧なプラットフォームを選択したいという衝動を抑えてください。それが何であるかはまだわかりませんが、制約が何度か顔を殴るまではわかりません。プラットフォームよりも存続できるようにコンテキストを構築する必要があります。

## Original Extract

Every alert starts with the same tax: which of five observability systems is telling the truth, and which of two production worlds am I even in? This is how our SRE alert-investigation agent evolved across six versions, from Cursor and MCP queries to a remote, Slack-triggered investigator, and why t
[truncated]

The Evolution of an SRE Agent: From Five Browser Tabs to a Slack Command | Blog EngineeringExec
Welcome to EngineeringExec , where coding expertise meets
executive leadership.
I'm Michael Bleterman , an R&D leader and software architect
with over 20 years of experience.
My mission is to bridge the gap between engineering and management.
© 2026 Michael Bleterman. All rights reserved
The Evolution of an SRE Agent: From Five Browser Tabs to a Slack Command
Every SRE knows the tax. Before you can fix an alert, you answer three questions that have nothing to do with the problem: which system is telling the truth, which environment am I looking at, and where did the last person who saw this write it down. On a good day that costs ten minutes. At 2 AM it costs the will to live.
At Tipalti the tax is heavier than most. We run two production-grade environments, each split into a legacy monolith on Windows machines and a fleet of microservices on K8s, worlds that must never be confused. Their signals are scattered across Datadog, Coralogix, Prometheus, Pingdom and some internal monitoring, part of it behind access controls. Point a vanilla LLM at that, ask what caused an alert, and it does what they all do without context: it makes something up, confidently.
This is the story of how we automated that first tax. Not in one clever build, but in six of them, and we threw most of them away. The runtime changed four times. What survived was everything else.
Version zero: a human, five tabs, and a lot of scrolling
Before 2025 there was no agent, just whoever was on call. Get the alert, open the matching observability system, start pulling threads: logs here, a trace there, a metric that looks wrong, a dashboard that looks fine. If the alert came from the monolith you reached for one set of tools and one mental model. If it came from K8s, a completely different set. The knowledge of which-tool-for-which-world lived in people’s heads, and it walked out the door every time someone changed teams. It worked. It just didn’t scale, and it didn’t sleep.
By 2025 the math stopped working. The team got leaner and the on-call rotation heavier, while every new platform feature widened the surface you had to understand. The context we needed to hold grew faster than the team did, a Red Queen race where you sprint just to stay in place. Keeping up meant making triage semi-automatic, and then automatic.
Late 2025 was the first real step. We wired a few observability systems into Cursor through MCP servers, so instead of clicking through five UIs you could ask, in one place, for the logs and traces around a given time window and get them back in seconds. Think of the quirky query languages, PromQL or DataPrime, that someone just took off your plate.
This was a speed win and nothing more. A human still read every result, still decided what to query next, still held the two-worlds map in their head. But turning “open five tabs and scroll” into “ask once” changed the rhythm of an investigation. It was the first time the machine did the fetching and the human did only the thinking.
Wrapping the queries in skills
By the end of the year the raw queries had a problem: everyone wrote them slightly differently, and slightly differently is how you end up reading Sandbox data while investigating Production. So we wrapped the common queries in team-level skills. A skill knew which environment it was allowed to touch, which fields mattered, and how to phrase the question so the answer came back accurate instead of merely plausible.
Accuracy went up, and the results stopped depending on who was driving. The flow was still manual, a person calling skills in sequence, but the skills now carried the context that used to live in people’s heads.
The standalone agent hits a wall
Then we got ambitious. Early 2026 we tried to stand the whole thing up as an independent agent, something that could run on its own instead of riding inside a person’s laptop IDE, triggered by systems rather than people. We evaluated AWS AgentCore as the host.
I pulled the handbrake myself. The idea was not wrong, but a cloud-hosted agent runtime is built to reach cloud-native, cloud-reachable things, and our world is the opposite of that: two gated environments, a Windows monolith, tools that do not expose themselves to whatever happens to be running in someone else’s cloud. We spent weeks trying to bridge the agent to the environments it needed to see, and the bridge kept collapsing. What made me kill it was not the difficulty but the shape of the work: the scope was creeping into environment plumbing with no MVP in short or medium reach. When a bet stops pointing at something you can ship and starts pointing at an infrastructure project, you kill it. The lesson was expensive and worth it. For this problem, where the agent runs is not a detail. It is the whole game.
The SRE Investigator: one skill to drive the rest
Spring 2026 we stopped trying to escape the editor and doubled down on it instead. Back in Cursor, we built the first version of a skill we called the SRE Investigator, and it changed the shape of everything. Until then, skills were tools a human picked up one at a time. The Investigator was the first skill whose job was to drive the other skills: take an alert, work out which world it belongs to, pull the right logs and traces and metrics in the right order, correlate them, and hand back a first-draft root cause.
The human went from driving the investigation to reviewing its conclusion. That is a different job, and a much faster one.
Claude Code, and finally keeping score
By June we had outgrown the editor again, this time for a good reason. We ported the whole thing to an agent-driven harness on Claude Code, and built something we should have built earlier: a way to keep score. A back-testing framework let us replay real past alerts through the agent and check whether it reached the root cause we already knew was correct. That gave us a number to move: over successive versions we pushed the agent to around 80% accuracy on replayed alerts, instead of guessing whether a change had helped.
This is the part nobody puts in the demo video, and the part that matters most. An investigation agent that is right most of the time and confidently wrong the rest is worse than no agent at all, because it teaches people to trust it. And the errors are not symmetric. A false negative that sends the on-call engineer down the wrong path can add half an hour to the time to resolution. A false positive burns that engineer’s night and costs you them the next day. Accuracy here is measured in customer minutes and lost sleep, not in a leaderboard score. The harness is what let us change the system and know, before shipping, whether we had made it smarter or just different.
Off the laptop: a Slack command
The latest reincarnation, July 2026, is the one that finally left the individual behind. The investigator now runs remotely, in claude.ai, triggered from Slack. You no longer need the right person, on the right laptop, with the right personal credentials, awake at the right hour. Someone posts an alert, the agent picks it up, does the triage that used to eat the first twenty minutes of every incident, and posts back a first-draft root cause while the humans are still reaching for coffee.
The capability stopped being something a few of us could do and became something the team has. Across the alerts it now handles, mean time to resolution is down about 30%, and most incidents reach a human with a first-draft root cause already attached.
In the background ran a rotating cast of vendor trials, tools that promised to do the whole job out of the box. Most failed on the same rock: they had no idea our worlds must never be mixed, no idea which system to believe, no way to earn the context that makes an answer trustworthy instead of fluent. Off-the-shelf could not tell the monolith from the microservices, and here that difference is everything.
We changed the runtime four times: browser tabs, Cursor, a dead-ended run at AgentCore, Claude Code, a remote service. Bet the project on any one of them and we would have thrown it away when that runtime stopped fitting. What we never threw away was the context: skills that know which world they may touch, queries that come back accurate, back-tests that catch a regression before it ships. The model was never the moat. The runtime was never the moat. The context was.
What’s next: a Multi-Agent system
One investigator is still one detective. The next version is a multi-agent framework, backed by a couple of things a current one does not have: memory, so the system stops re-deriving a root cause it already found last month and starts saying “we have seen this one”, and RAG over our own post-mortems, runbooks and past investigations, so the context that makes an answer trustworthy grows without being crammed into a single prompt.
If you are building something similar, resist the urge to pick the perfect platform first. You do not know what it is yet, and you will not until the constraints punch you in the face a few times. Build the context so it can outlive the platform, because it will have to.
