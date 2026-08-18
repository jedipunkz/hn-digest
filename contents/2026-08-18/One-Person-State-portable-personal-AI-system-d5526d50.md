---
source: "https://embassy.svit.la/p/one-person-state"
hn_url: "https://news.ycombinator.com/item?id=49348919"
title: "One Person State – portable personal AI system"
article_title: "One Person State - by Ray Svitla - 404Embassy"
image: "https://substackcdn.com/image/fetch/$s_!HsCu!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F5ae4e23d-fac3-48f0-a5fe-cec38ff83ffe_1456x816.png"
author: "promen_svitla"
captured_at: "2026-08-18T17:18:41Z"
capture_tool: "hn-digest"
hn_id: 49348919
score: 1
comments: 0
posted_at: "2026-08-18T17:11:51Z"
tags:
  - hacker-news
  - translated
---

# One Person State – portable personal AI system

- HN: [49348919](https://news.ycombinator.com/item?id=49348919)
- Source: [embassy.svit.la](https://embassy.svit.la/p/one-person-state)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T17:11:51Z

## Translation

タイトル: One Person State – ポータブルパーソナル AI システム
記事のタイトル: 一人国家 - レイ・スヴィトラ著 - 404Embassy
説明: 私は 2 つの物理的な家を失ったので、マークダウン、Discord、オートメーション、そして Promen という AI エージェントを使ってポータブルなものを構築し始めました。

記事本文:
一人国家 - レイ・スヴィトラ著 - 404Embassy
購読 サインイン 1 人の状態
動き続ける人生のバックアップを構築する
Ray Svitla 2026 年 8 月 18 日 シェア 私はここ数年で物理的な家を 2 軒失いました。
詩的な「家は心の在り処」という意味ではありません。実際の家。場所、ルーチン、書類、物体、生活の周りのすべての退屈なインフラストラクチャは、どこか別の場所に再構築するまで気付かないものです。
二度目の強制引っ越しの後、私は違う種類の家にますます力を入れるようになりました。私と一緒に引っ越しできて、理想的には、家主、特定の会社、アプリ、またはその月にたまたま私の書類を持っていた機関にあまり依存しないものです。
マークダウンファイル、古いMacBook Air、プライベートDiscordサーバー、スクリプト、バックアップ、多すぎる統合、そして私の個人エージェントであるプロメンは、今では私の人生について少し不快なほど知っています。
途中のどこかで、それは私が使用する単なるツールのコレクションではなくなり、私の生活の一部になり始めました。
私はそれを「一人国家」と呼んでいます。意図的に大げさすぎる名前になっています。それも冗談の一部です。
ネットワーク状態から一人状態へ
このブログは元々、私がロンドン、リスボン、ベルリン、ドバイ、およびオンラインで企画した一連のイベント、Future State から生まれました。そこでの会話の多くは、ネットワークの状態とそれに付随するアイデアに関するものでした。人々がオンラインで連携すること、コミュニティが地理にあまり依存しなくなること、新しい制度、インターネットネイティブの経済など、まさに未来を形づくるものすべてです。
それから、人生のあまり未来的ではない部分がずっと大きくなりました。
移民、お金、家族、官僚制。移動しなければならない。ルーチンを再構築する必要があります。しかし、そうではないことを学びます

将来のことを考えても、パスポート管理局があなたの一年を完全に台無しにする可能性は依然としてあります。
ある時点で、かなり明白な疑問が生じました。
一人を組織できないなら、一体なぜ一万人を組織しようと考えているのでしょうか？
一人国家は基本的に、私の生活の管理層のどれだけを実行、検査、変更するのが容易になるかを理解するための私の継続的な試みです。
自立するという意味ではありません。一人国家には驚くほど長い依存関係リストがあります。友人、都市、公共インフラ、会社、幸運、メッセージに答えてくれる人、場所に誘ってくれる人、その他私がコントロールできないものはたくさんあります。
私は今でも人々ともっと多くの時間を過ごしたいし、GPU を搭載したバンカーを建てたくありません。
しかし、「生きていること」と「自分の人生を生きること」の間には層があり、それは主に管理者によるものです。つまり、物事を思い出すこと、フォローアップすること、適切な文書を見つけること、小さな決定を記録すること、そして毎週月曜日にゼロから始めるわけではありません。
その部分はますます自動化できるように感じました。または少なくとも管理が容易になります。
悩ましい答えは「たくさんある」です。
この記事を書いているときに、初めてファイルの数を数えました。
.git 、node_modules 、キャッシュ、エクスポート、ハンドオフ、ビルド アーティファクト、その他の技術的な不要物を除外した後、作業レイヤーには現在 41,468 個のファイルが含まれています。
これが私が個人的なオペレーティング システムを構築したことを証明しているのか、それとも問題があることを証明しているのかはわかりません。おそらく両方だろう。
これらのファイルのほとんどは、作業メモ、生成されたレポート、調査、プロジェクトのコンテキスト、古い成果物、ログ、およびおそらく削除されるべきものです。この記事を書くことの有益な副作用の 1 つは、実際にどれだけのたわごとが蓄積されているかに気づいたことです。
アーキテクチャはファイル数から想像されるよりもはるかに単純です。信号は別のPから入ってきます

私の人生の芸術、いくつかは思い出になります、Promenはそのコンテキストで物事を行います、そしてDiscordは私がそのほとんどを対話する場所です。
メインマシンは古い M1 MacBook Air で、8 GB の RAM と 256 GB のストレージを搭載しています。
それはすでに予備のラップトップとして家に転がっていました。今では 24 時間年中無休で動作し、Hermes の上で Promen を実行します。私のメインの MacBook と iPhone は Tailscale 経由で接続しているので、基本的にどこからでも同じマシンにアクセスできます。
視覚的に何かを修正する必要がある場合は、画面を共有します。より低いレベルのものが必要な場合は、SSH を使用します。時々私は Promen に MCP 経由で Ableton に接続するなどの不必要なことをさせます。主にそうしないからです。
これは特に高度なホームラボエンジニアリングではありませんが、私は気に入っています。
ファイルがどこにあるのか、マシンが停止した場合に何を復元する必要があるのか​​は大体わかっています。
現在の回復レイヤーにより、数時間以内にクリーンな Mac 上で動作する Promen を戻すことができます。私にとって回復とは継続性をますます意味します。そこに至るまでの経緯を失わずにマシンを失うことができることです。
完全な会話履歴もその一部に含めたいと考えています。プロメンが私のプロジェクト、修正、奇妙な好みを学習した方法の一部として何年もの会話があったとしたら、それらを使い捨てのチャット履歴として扱うのは、かなり愚かなバックアップ戦略になるでしょう。
「ローカル」という言葉も、AI に関しては特に扱いにくい言葉です。
選択されたコンテキストが外部モデルに送られている間、真実のソースは MacBook 上に存在することができます。 API は依然として API です。他人の SaaS データベースからファイルを移動しても、第三者が何も処理しないわけではありません。
私にとって重要なのは監護権です。生活を変えることなくモデルを交換できます。
システムの大部分はプライベート Discord サーバーの背後に存在します。
人々は私がなぜ Telegram を使わないのかと尋ね続けます。主にテレグラムがすでにあるため

yには私の人生があまりにも多く含まれています。生活管理の層をさらに追加することは、最初のキッチンが乱雑であるため、最初のキッチンの中に 2 番目のキッチンを置くようなものです。
サーバーには、 #health 、 #finance 、 #bureaucracy 、 #radar 、 #one-person-state 、 #dj 、 #shopping 、 #roles-radar など、最近のクリーンアップで生き残ったものがすべてあります。 #dating にも独自のチャンネルがあります。それについてはすでに他の場所で言いすぎました。
外から見ると、おそらく少し狂っているように見えるでしょう。その中での生活はそれほどドラマチックではありません。そして分類法も神聖なものではありません。私はそれを常に変更します。
便利な部分は、コンテキスト境界があることです。健康について話しているとき、同じ文脈に滞在許可証に関する膨大なチャット履歴は必要ありません。何かが一時的なものであれば、通常はスレッドを作成します。 1 か月後もスレッドが何らかの形でまだ生きている場合、そのスレッドはチャンネルになる権利を獲得したことになります。
プロメンは依然として、これらの分野にわたるより広範なコンテキストを保持できますが、すべての会話に私の人生全体が含まれる必要はありません。
私はどういうわけかすでに少なくとも 3 人にこのアイデアを感染させてしまいました。彼らは私のものを見て独自の OpenClaw /Hermes セットアップを構築し、最終的に 3 人全員が Telegram を使用することになりました。
したがって、Discord は明らかに重要ではありません。その部分はまさに私特有の病気です。
現時点では GPT が多くの作業を行っていますが、私はクロード、キミ、ジェミニなども使用しています。 1 つはコードに適しており、もう 1 つは研究に適しており、もう 1 つはバックグラウンド ジョブに使用できるほど安価です。
重要なのは、それらの誰も記憶を所有していないということです。
私のコンテキストのほとんどは、プロジェクト、好み、意思決定、ルール、履歴など、モデルの外側にあります。そのため、モデルを切り替えると最初からやり直す気にはなれません。
私が構築した、より遅いフィードバック ループもあります。基本的には、プロメンの毎週のマスターマインド セッションです。
毎晩

最近のセッションから証拠を収集します: どこを修正したか、何がうまくいかなかったのか、何が繰り返されたのか。週に一度、取締役会は次の 1 つの質問を検討します。
今週実際にレイの邪魔をした繰り返しの間違いは何ですか?
明確なパターンがある場合にのみ、 ~/.hermes/SOUL.md に小さな変更が 1 つ加えられます。
そのため、モデルを切り替えても、蓄積された関係はほとんどそこに残ります。新しいモデルには、同じコンテキスト、ルール、修正履歴が適用されます。
これが、私が第二の脳アプリのカテゴリー全体についてあまり気にしなくなった理由でもあります。私はオブシディアンに対して何も反対しません。どのアプリがマークダウンを表示するかはもう気にしません。
プロメンが実際に使えるかどうかが気になります。
私はその情報が何かをすることができるかどうかを気にしています。
家計に関するメモはフォルダーに永遠に保存されることもあれば、日曜日のレビューで今後の出費の隣に表示され、先月決めたことを思い出させたり、不快な支出パターンを無視するのが難しくなったりすることもあります。
私にとっては、完璧に組織された第二の脳を持つことよりもずっと興味深いです。
今年の初めに、私は散らかったもの、お金、より多くの物を発送すること、肉体的および精神的な基盤をより良い状態にすること、人間関係、そしてこの個人的な OS の実験全体について、いくつかの幅広い目標を立てていました。
そのため、大量のデータがシステムに流入し始めました。
私のAmazfitストラップとスマートスケールはApple Healthに栄養を与えています。 DayFlow は、私がコンピューター上で実際に何をいつ行うかを監視します。 ZenMoney を使用すると、支出とランウェイのビューが得られます。これらはすべて、数時間ごとに自動的にサーバーにエクスポートされます。
週に数回、自分の気分についてチャットに音声メモを書き込むこともあります。何が私にエネルギーを与えてくれたのか、何が私をイライラさせたのか、何を避けてきたのか、疲れているのか、何かが起こり続けているのか。
私はポルトガルに住んでいるので、官僚制には独自のチャンネルがあります。
メールが監視される

あまりにも。接続されている 1 つの Gmail アカウントだけで、過去 30 日間に 451 件の受信メッセージを受信しました。重要な管理メッセージが他のメッセージの間に隠れてしまうのに十分な量です。
プロメンは、実際に私を必要とするものと、そこに留まっていてもよいもの、つまり返信する、返信せずに行動する、監視する、無視するものを分離します。
ほとんどの場合、私は依然として結果的なアクションをレビューの背後に置いています。私は特に自治個人政府が移民弁護士にフリースタイルで返信することを望んでいません。
毎週日曜日には、これらの断片が、いくつかのグラフィック、いくつかのテキスト、いくつかの箇条書きなど、プライベートな週次レポートにまとめられます。何が動いたのか、何が動かなかったのか、私がやると言ったこと、何が起こり続けているのか、どのデータが欠けているのか。
DayFlow を使用すると、その日が通話でいっぱいになる前に、より良い仕事が行われることが明らかになったので、可能な限り通話を 15:00 以降に移動しました。
おそらくAIがなければこれを理解できたでしょう。実際のパターンについて議論するのははるかに困難でした。
これが、何かを追跡するときの私のルールです。決定が変わらない場合は、おそらくそれは必要ありません。
私が最も使用する部分はおそらくスキルとスケジュールされたジョブです。
スキルは基本的に、何かを行うための再利用可能な方法です。単なる短い指示の場合もあれば、多数のツールを調整する場合もあります。
現在、Promen は、Hermes の 202 のランタイム スキルと 33 のスケジュールされたジョブにアクセスできます。
過去 90 日間で、442 件の Discord セッションと 4,292 件のバックグラウンド cron セッションが記録されました。これらのバックグラウンド実行のほとんどは、チェック、準備、監視、または何もすることがないと判断するだけです。これが私が望んでいることです。
繰り返される行動は徐々にインフラになる可能性があります。同じことをやり続けたり、同じ出力を修正したりすると、そのパターンを頭の中に留めておくのではなく、スキルに変えることができます。
それらのスキルの 1 つは、私のロシア語の草案を ki に翻訳します。

私が実際に公に使っている英語の nd です。この記事でも同様に説明したため、これは少し再帰的です。
欠点は明らかです。システムが独自の官僚主義を成長させます。この記事を書いているときに、重複する仕事、もう必要のない問題に対する古いスキル、目的を過ぎたルールを見つけました。
いくつかは統合されました。いくつかは削除されました。
一人国家でも時折解雇は必要だ。
より大きなシステムの 1 つは Radar と呼ばれるもので、self.md の毎日の信号セクションを実行します。
Hacker News、Reddit、RSS、GitHub、YouTube にわたる 68 のフィード、コミュニティ、検索、収集ルートに加え、私が手動で投入したあらゆるものを監視します。
8 月 10 日の重複除外の結果、最終的に 274 個の一意の候補が得られました。
毎朝 07:00 頃にリサーチとライティングパスを実行します。 08:30 頃、別のジョブがそれをカードに変換し、すべてをサイト、X、スレッド、テレグラムに投稿します。
5 月 3 日以来、毎日 106 個のカルーセル バンドルが生成されました。
正直なところ、古い MacBook Air の中に住んでいる小さなメディア部門にとっては、107 点中 106 点で十分です。
ほとんどの場合、マシンは調査、作成、アセットのチェック、および全体の公開に 20 ～ 50 分を費やします。
この種の奇妙な個人的なインフラストラクチャが好きなら、購読してください。おそらく変化しながら記録し続けるでしょう

[切り捨てられた]

## Original Extract

i lost two physical homes, so i started building a portable one in markdown, Discord, automations, and an AI agent called Promen

One Person State - by Ray Svitla - 404Embassy
Subscribe Sign in One Person State
building a backup for a life that keeps moving
Ray Svitla Aug 18, 2026 Share i’ve lost two physical homes in the last few years.
not in the poetic “home is where the heart is” sense. actual homes. places, routines, documents, objects, all the boring infrastructure around a life that you don’t really notice until you have to rebuild it somewhere else.
after the second forced move i started putting more and more effort into a different kind of home. one that could move with me and, ideally, wasn’t too dependent on a landlord, one particular company, one app, or whichever institution happened to have my documents that month.
most of it ended up being extremely unromantic: markdown files, an old MacBook Air, a private Discord server, scripts, backups, too many integrations and Promen — my personal agent — which by now knows a slightly uncomfortable amount about my life.
somewhere along the way it stopped being just a collection of tools i use and started becoming a layer around my life.
i’ve been calling it One Person State . the name is intentionally too grand. that’s part of the joke too.
from network state to one person state
this blog originally grew out of Future State, a series of events i organized in London, Lisbon, Berlin, Dubai and online. a lot of the conversations there were about the network state and adjacent ideas : people coordinating online, communities becoming less dependent on geography, new institutions, internet-native economies, all of that very future-shaped stuff.
then the much less future-shaped parts of life got so much louder.
immigration, money, family, bureaucracy. having to move. having to rebuild routines. learning that however distributed the future becomes, a passport office can still ruin your year perfectly well.
at some point a fairly obvious question appeared:
if i can’t organize one person, why the fuck am i thinking about organizing ten thousand?
One Person State is basically my ongoing attempt to figure out how much of the administrative layer of my life can become easier to carry, inspect and change.
i don’t mean becoming self-sufficient. the One Person State has a surprisingly long dependency list: friends, cities, public infrastructure, companies, luck, people answering my messages, people inviting me places and a thousand other things i don’t control.
i still want to spend more time with people and i definitely don’t want to build a bunker with GPUs in it.
but there is a layer between “being alive” and “living your life” that is mostly admin: remembering things, following up, finding the right document, keeping track of small decisions, and not starting from zero every Monday.
that part felt increasingly automatable. or at least easier to manage.
the annoying answer is: a lot.
while working on this article i counted the files for the first time.
after excluding .git , node_modules , caches, exports, handoffs, build artifacts and other technical garbage, the working layer currently contains 41,468 files .
i’m not sure whether this proves that i built a personal operating system or that i have a problem. probably both.
most of those files are working notes, generated reports, research, project context, old artifacts, logs and things that should probably be deleted. one useful side effect of writing this article has actually been noticing how much shit had accumulated.
the architecture is much simpler than the file count makes it sound. signals come in from different parts of my life, some become memory, Promen does things with that context, and Discord is where i interact with most of it.
the main machine is an old M1 MacBook Air with 8 GB of RAM and 256 GB of storage.
it was already lying around at home as a backup laptop. now it stays on 24/7 and runs Promen on top of Hermes . my main MacBook and iPhone connect to it through Tailscale , so i can get to the same machine from basically anywhere.
if i need to fix something visually, i screen-share into it. if i need something lower-level, i use SSH. sometimes i make Promen do unnecessary things like connect to Ableton through MCP, mostly because why not.
none of this is particularly advanced homelab engineering, but i like it.
i know where the files are and roughly what has to be restored if the machine dies.
the current recovery layer can bring a working Promen back on a clean Mac within a few hours. recovery for me increasingly means continuity: being able to lose the machine without losing the history of how i got there.
i want the full conversation history to be part of that too. if years of conversations are part of how Promen learned my projects, corrections and weird preferences, treating them as disposable chat history would be a pretty stupid backup strategy.
“local” is also an especially slippery word around AI.
the source of truth can live on my MacBook while selected context still goes to an external model. an API is still an API. moving files out of someone else’s SaaS database doesn’t mean no third party ever processes anything.
for me the important part is custody. i can replace the model without having to move my life with it.
most of the system lives behind a private Discord server.
people keep asking why i don’t use Telegram. mostly because Telegram already contains too much of my life. adding another layer of life administration there would be like putting a second kitchen inside the first kitchen because the first one is messy.
the server has things like #health , #finance , #bureaucracy , #radar , #one-person-state , #dj , #shopping , #roles-radar and whatever else survived my latest cleanup. #dating has its own channel too. i’ve already said too much about that elsewhere.
from outside it probably looks mildly insane. living inside it is not that dramatic. and the taxonomy isn’t sacred either — i change it all the time.
the useful part is having contextual boundaries. when i’m talking about health, i don’t need a giant chat history about a residence permit sitting in the same context. if something is temporary, i usually create a thread. if the thread is somehow still alive a month later, it has earned the right to become a channel.
Promen can still retain broader context across those areas, but every conversation doesn’t have to carry my entire life inside it.
i’ve somehow already infected at least three people with this idea. they built their own OpenClaw /Hermes setups after seeing mine — and all three ended up using Telegram.
so Discord is clearly not the point. that part is just my particular disease.
right now GPT does a lot of the work, but i’ve used Claude, Kimi, Gemini and others too. one is better for code, another for research, another gets cheap enough to use for background jobs.
the important part is that none of them own the memory.
most of my context lives outside the model: projects, preferences, decisions, rules, history. so switching models doesn’t feel like starting over.
there is also a slower feedback loop i built — basically a weekly mastermind session for Promen.
every night it collects evidence from recent sessions: where i corrected it, what kept going wrong, what repeated. once a week, a board looks at one question:
what recurring mistake actually got in Ray’s way this week?
only if there is a clear pattern does it make one small change to ~/.hermes/SOUL.md .
so when i switch models, the accumulated relationship mostly stays there too. the new model gets the same context, rules and history of corrections.
this is also why i stopped caring much about the whole second-brain app category. i have nothing against Obsidian . i just don’t really care which app displays the markdown anymore.
i care whether Promen can actually use it.
i care whether the information can do something.
a note about my finances can sit in a folder forever, or it can show up in a Sunday review next to upcoming expenses, remind me what i decided last month, and make an uncomfortable spending pattern harder to ignore.
that’s much more interesting to me than having a perfectly organized second brain.
at the beginning of this year i had a few broad goals around clutter, money, shipping more things, getting my physical and mental base into better shape, relationships, and this whole personal OS experiment.
so a bunch of data started flowing into the system.
my Amazfit strap and smart scales feed Apple Health. DayFlow watches what i actually do on my computer and when. ZenMoney gives me a view of spending and runway. all of that gets exported to my server automatically every few hours.
a few times a week i also just dictate an audio note into the chat about how i’m feeling. what gave me energy, what annoyed me, what i’m avoiding, whether i’m tired, whether something keeps coming up.
bureaucracy has its own channel because i live in Portugal.
email gets watched too. one connected Gmail account alone received 451 incoming messages over the last 30 days . enough volume for an important administrative message to disappear between everything else.
Promen separates things that actually require me from things that can sit there: reply, act without replying, monitor, ignore.
in most cases i still keep consequential actions behind review. i do not particularly want my autonomous personal government freestyle-replying to immigration lawyers.
every Sunday these fragments get pulled together into a private weekly report: some graphics, some text, some bullet points. what moved, what didn’t, what i said i would do, what seems to keep happening, what data is missing.
DayFlow made it obvious that better work happened before the day filled with calls, so i moved calls after 15:00 where i could.
i probably could have figured this out without AI. it was just much harder to argue with the actual pattern.
that’s roughly my rule for tracking anything: if it doesn’t change a decision, i probably don’t need it.
the part i use most is probably skills and scheduled jobs.
a skill is basically a reusable way of doing something. sometimes it’s just a short instruction, sometimes it coordinates a bunch of tools.
right now Promen has access to 202 runtime skills in Hermes and 33 scheduled jobs .
over the last 90 days it recorded 442 Discord sessions and 4,292 background cron sessions . most of those background runs just check, prepare, monitor, or decide there’s nothing to do — which is what i want.
repeated behavior can slowly become infrastructure. if i keep doing the same thing or correcting the same output, i can turn that pattern into a skill instead of keeping it in my head.
one of those skills translates my russian drafts into the kind of English i actually use publicly. which is slightly recursive, since this article went through it too.
the downside is obvious: the system grows its own bureaucracy. while writing this article i found overlapping jobs, old skills for problems i no longer had, and rules that had simply outlived their purpose.
some got merged. some got deleted.
sooo even a one-person state needs occasional layoffs.
one of the larger systems is something called Radar, which runs the daily signals section on self.md
it watches 68 feeds, communities, searches and collection routes across Hacker News, Reddit, RSS, GitHub and YouTube, plus whatever i manually throw into it.
on August 10 it ended up with 274 unique candidates after de-duplication.
every morning around 07:00 it does the research and writing pass. around 08:30 another job turns that into cards and posts everything to the site, X, Threads and Telegram.
since May 3 it has produced 106 daily carousel bundles .
honestly, 106 out of 107 is good enough for a tiny media department living inside an old MacBook Air.
most days the machine spends 20–50 minutes researching, writing, checking assets and publishing the whole thing.
if this kind of weird personal infrastructure is your thing, you can subscribe. i’ll probably keep documenting it as it mutate

[truncated]
