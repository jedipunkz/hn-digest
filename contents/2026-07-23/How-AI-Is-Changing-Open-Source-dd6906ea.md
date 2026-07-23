---
source: "https://enblog.eischmann.cz/2026/07/23/how-ai-is-changing-open-source/"
hn_url: "https://news.ycombinator.com/item?id=49026908"
title: "How AI Is Changing Open Source"
article_title: "How AI Is Changing Open Source – Brno Hat"
author: "pksadiq"
captured_at: "2026-07-23T20:13:44Z"
capture_tool: "hn-digest"
hn_id: 49026908
score: 4
comments: 0
posted_at: "2026-07-23T19:34:56Z"
tags:
  - hacker-news
  - translated
---

# How AI Is Changing Open Source

- HN: [49026908](https://news.ycombinator.com/item?id=49026908)
- Source: [enblog.eischmann.cz](https://enblog.eischmann.cz/2026/07/23/how-ai-is-changing-open-source/)
- Score: 4
- Comments: 0
- Posted: 2026-07-23T19:34:56Z

## Translation

タイトル: AI はオープンソースをどのように変えるか
記事のタイトル: AI はオープンソースをどう変える – ブルノハット

記事本文:
AI がオープンソースをどう変えるか
今年は AI がソフトウェア開発に本格的に参入し、オープンソース プロジェクトにも大きな影響を与えています。この記事では、AI に関連してオープンソースで私が最近観察したいくつかのトレンドと、これらのトレンドがオープンソース ソフトウェアの世界をどのように変えているかについて説明します。
この記事はもともとチェコ語のブログに公開されたものですが、非常に大きな反響があったため、英語に翻訳してここでも公開することにしました。
一般に AI がもたらすトレンドの 1 つは、コンテンツの爆発的増加です。検索結果には生成された Web サイトが表示され、ソーシャル ネットワークには生成された画像やビデオが氾濫します。ソースコードも例外ではありません。現在、GitHub は増え続けるリポジトリの中に埋もれています。
しかし、質の高いプロジェクトが増えているわけではありません。逆に言えば、頼っていいのかどうかも分からない案件です。以前は、数千行のコードを含むより大規模なプロジェクトに遭遇した場合、誰かがそのようなものをわざわざ作成するのであれば、その人は問題についてある程度の知識を持ち、自分の作成したものと個人的なつながりがあり、今後もそれを維持する意欲があるだろうという一定の前提がありました。
もうこれには全く頼れません。現在では、数千行のコードを含むプロジェクトを瞬時に生成できます。それはまったくナンセンスである可能性もあれば、危険なものである可能性さえあります。それは、誰かが自分の差し迫ったニーズのために生成して GitHub に投稿した、機能的なものである可能性がありますが、それをオープンソース プロジェクトにすることに興味はありません。コードを含むリポジトリはオープンソース プロジェクトを作成しないからです。 GitHu 上のコードの違い

b およびオープンソース プロジェクトとは、オープンソース プロジェクトが作成者の 1 回限りのニーズだけでなく、ユーザーの問題やユースケースを解決するということです。そして、そのような簡単で汚いコードの作成者のほとんどは、単にそのようなことをすることに興味がありません。
これは、たとえば MeshCore のようなプロジェクトではっきりとわかります。想像できるすべてのフォークが数十個あります。公式 MeshCore ファームウェアに不足している機能はありますか?それをフォークし、不足している部分をバイブコーディングして、GitHub に MeshCore-UltimateEdition としてダンプするだけです。問題は、それが最小限の労力で作成され、作者は通常、それとは何の関係もなく、1か月後には飽きてしまい、熟成する前に放棄されてしまうことです。
約 10 年前、Linux リポジトリの概念は終わったと人々が言い始めました。 2000 年代には、Linux ソフトウェアの実質的に唯一のソースでした。プロジェクトが配布リポジトリに入れられなかった場合、それには問題がありました。しかしその後、オープンソース プロジェクトの数が、ディストリビューションが追いつかないほどの速度で増加しました。ユーザーはソフトウェアを他の場所から入手し始める必要があり、ソフトウェア作成者はディストリビューションなしで済むことを学びました。ほんの数年前まで、「必要なものはすべて Debian にある」というアプローチは決定的に死んだように思えました。
ただし、Linux ディストリビューション リポジトリなど、厳選されたソフトウェア ソースが復活する可能性はあります。オープンソース ソフトウェアの世界は非常に混沌としており、ユーザーは誰かが精査し、半年後も信頼できる精選されたソフトウェアを含むソースを再び高く評価し始めるでしょう。
AI がオープンソースで引き起こしたもう 1 つの傾向は、「レビューの圧倒」です。以前は、コードの作成には多大な労力と時間の投資が必要だったため、自然なフィルターとして機能していました。

入口。それがなくなり、コードの作成が迅速かつ簡単になりました。ただし、このコードを本格的な運用に移す前に、誰かがこのコードをレビューする必要があります。オープンソース プロジェクトで長年機能してきたレビュー プロセスは、現在その能力の限界に達しています。
GNOME 50 では、長い間誰も Google ドライブを保守していなかったために、Google ドライブのサポートが削除されました。ユーザーは当然これに不満を抱きましたが、最終的に 1 人のユーザーが名乗り出て、サポートを再度追加し、それを gvfs プロジェクトの上流に提出しました。
そのプロジェクトの保守を担当する同僚は、4,000 行のコードを伴う変更だったと嘆いていました。基本的なレベルでは動作しているように見えますが、明らかに AI を使用して生成されました。彼は、それが実際に意図したとおりに動作し、長期的な保守に取り組むために必要なコード品質基準を満たしていることを確認するために、それを 1 行ずつ確認する必要があります。
このため、作業のほとんどはコード作成からコードレビューに移行しましたが、これは AI では一般的です。しかし、オープンソースの問題は、コードをレビューしてマージするのに十分な経験を積んだ開発者が、AI の前にすでにボトルネックになっていたことです。現在、問題は大幅に深刻化しています。上の例では、私の同僚は、寄稿者が敏感に反応し、この問題に長期的な関心を示してくれたことを幸運だと思っています。
今日では、それはむしろまれな例外です。よくある貢献は、主題に対する深い関心や理解もなしに、誰かが何かを乱暴にバイブコーディングし、それを壁の向こう側にメンテナに投げつけることで構成されます。
私は Meshy でこれについてかなり最近の経験があります。誰かが、macOS のサポートを追加することを想定した 9,000 行の変更されたコードを含むプル リクエストを送信しました。ある晩、1 時間をかけて簡単なレビューを行いましたが、その短い時間でも多くの問題に遭遇しました。

コードはあからさまに AI によって生成され、数千行は単一引用符を二重引用符に完全に置き換えただけで、問題に関係のないコードの部分が変更され、過去数週間にメイン ブランチで行ったすべての変更が上書きされました。
著者は私のコメントに一度も返信しなかったし、その後連絡もありませんでした。私が学んだことは、そのような投稿にはその 1 時間でさえ時間の投資が大きすぎるということであり、次回からはもっと早く拒否するつもりです。
一部のプロジェクトは、基本的な拠出要件を厳格化することでこの状況に対応しています。たとえば、AI 生成アプリを拒否するという Flathub の決定は大きな波紋を引き起こしました。多くの人が自分の足を撃ったとして批判しましたが、彼らの現実を見なければなりません。
Flathub は現在数千のアプリをホストしており、毎日さらに多くのアプリが追加されています。レビューを担当するのはわずか 3 人です。レビュープロセスは高度に自動化されていますが、非常に徹底的に行われており、依然として多くの手動入力が必要です。彼らの目標は、最悪のスロップを特定することだけではなく、比較的高い水準のコードの健全性を維持することであることは明らかです。過去 6 か月間で、私は 2 つのアプリを Flathub に提出しました。そのレビュー プロセスは、最終的にアプリ自体の品質の向上に貢献しました。
しかし、これは現在、個人的な努力を一切せずに、完全にバイブコーディングされたアプリを提出する人々の現実と衝突しています。含めるようリクエストするためのチケット テンプレートには、アプリの動作方法を示す短いビデオをアップロードする要件など、いくつかの質問があります。それほど難しい作業はなく、誰でも 15 分以内に組み立てることができます。しかし、AI スロップの作成者にとっては、それさえも労力がかかりすぎます。
これらの最小限の要件を満たす代わりに、これを Linux の自由への攻撃と呼ぶ人もいます。

誰もが利用できるはずだった Flathub の代替案を即座にバイブコーディングしました。当然のことながら、それはわずか1か月しか続きませんでした。
オープンソースのメンテナには、コードレビューに対するこの需要を満たす能力がないだけでなく、コードレビューを行う動機も失いつつあります。多くの場合、自社で機能を作成した方が早いのですが、歴史的にはレビュー プロセスが、新しい長期貢献者や潜在的な後継者を育成する方法でした。誰かがあなたに、労力もかからず、理解すらできないであろう雰囲気コード化された貢献を送ってきたとき、どうやってその人を指導して、長期的にプロジェクトを助ける貢献者に育てると期待しますか?
オープンソース ソフトウェアは決して最終結果だけを目的としたものではありません。それはまた、貢献者がプロジェクトとの関係を築き、最終的にはそれを他の人に伝える人に成長するというプロセスに関するものでもありました。これは、結果がすべてである AI の世界とは対照的です。できるだけ早く、できるだけ少ない労力で。
コードを公開するモチベーションの低下
1990年代、フランシス・フクヤマは民主主義と自由主義経済学が世界秩序の整備における最終的な勝利者であると宣言した。世界中で民主主義が侵食され、既存の経済秩序が崩壊している今日、これは控えめに言っても時期尚早の大胆な発言のように見える。同様に、ほんの数年前、過去数十年の発展に感銘を受け、オープンソースをソフトウェア開発モデルの中での最終的な勝者として称賛する人もいました。私たちは福山氏の論文と同様の厳しい現実チェックに直面しようとしているのだろうか？
最近、私は、オープンソース開発から手を引くという、微妙ではあるものの現在進行形の傾向を観察しています。オープン開発に反対する議論の 1 つは、前述のレビューの過負荷に関するものだと聞いています。いくつかのpのために

プロジェクトの場合、AI のスロップに圧倒されることに関連するコストが、有益なコミュニティへの貢献によるメリットを上回る可能性があります。透明性を確保するためにソース コードを公開する場合もありますが、オープン開発プロジェクトからオープンソースのクローズド開発プロジェクトに変わります。また、透明性をそれほど気にしない人は、ソース コードを完全に閉じてしまう可能性があります。
ソースコードの公開に反対するもう 1 つの議論は、ライセンス回避の恐れです。現在の LLM は、ライセンスに関係なくソース コードをトレーニングし、選択したライセンスに基づいて公開できる同様のソリューションを簡単に生成できます。
作者は、誰もがコードを使って事実上何でもできることをすでに認めているため、これは寛容なライセンスの問題ではありません。ただし、AI は、GNU GPL のようなコピーレフト ライセンスに対して直接の脅威となります。通常、作者は自分の作品を永久に公開し、それを使用する人が改善点をコミュニティに共有できるようにするためにこれらを選択します。 LLM が、あなたが何年も取り組んできたプロジェクトに基づいてトレーニングし、その後、独自のライセンスの下で公開された非常によく似たソリューションを生成する場合、事実上、この原則を回避することになります。
もう一度 MeshCore を例に挙げます。プロトコル自体とファームウェアはオープンソースですが、クライアントはクローズドです。最近、コア チーム メンバーが密かに MeshCore 商標を申請し、利用可能なコードに基づいて独自のクローズド ソース ソリューションのバイブコーディングを開始したことがコミュニティで明らかになりました。 MeshCore の創設者 Scott Powell は、クライアントのソース コードを非公開にしておくという自身の決定を再確認するものとしてこれを引用しました。具体的には、次のように書いています。
つまり、AI の時代において、オープンソースとは、他人のために自分の血、汗、涙を捧げ、その方法は無数にあると考えています。
意見の相違があるかもしれない

パウエルの視点ではそうですが、これは私の周囲でますます頻繁に見られる正当な立場を表しています。私は、オープンソースを長年支持してきた人たち、つまり以前は共有したいがためにヘルパー スクリプトをすべて公開していた人たちを見かけますが、今ではそれらのものを自分だけのものにし、要求があった場合にのみ他の人に提供しています。彼らにもパウエル氏と同様の理由がある。
オープンソースも共有の必要性から生まれました。コードを書くのは大変でした。それを維持するのはさらに困難でした。なぜ全員が同じことを個別に実装する必要があるのでしょうか?オープンソース プロジェクトに力を合わせて共有ライブラリを作成しましょう。そうすれば誰もがその結果から恩恵を受けることができます。今日のインターネットを支えるインフラストラクチャは、この基盤の上に構築されました。しかし、AI はこのニーズを抑制しています。
たとえば、WordPress は「独自の CMS を簡単に生成できるから」という理由で死んだという意見に遭遇します。私の見解では、これはオープンソース プロジェクトが実際に提供するものを大幅に過小評価していると思います。これは単にコードを記述するだけではなく、オープンソース プロジェクトへの依存関係を LLM への依存関係に交換するこの戦略は、長期的には利益を生まない可能性があります。
それにもかかわらず、共有オープンソース コンポーネントへの依存は確かにある程度減少しました。 AI がすべてを置き換えるわけではないかもしれませんが、なぜ大きな電子に依存するのでしょうか?

[切り捨てられた]

## Original Extract

How AI Is Changing Open Source
AI entered software development at full speed this year, and it is significantly impacting open-source projects as well. In this article, I discuss several trends I have recently observed in open source in connection with AI, and how these trends are changing the world of open-source software.
This article was originally published on my Czech blog , but it received such an overhelming response that I decided to translate it into English and publish it here as well.
One of the trends that AI brings in general is an explosion of content. Search results are filled with generated websites, and social networks are inundated with generated images and videos. Source code is no exception. Today, GitHub is drowning in an ever-increasing number of repositories.
However, it is not as if a larger number of high-quality projects are being created. On the contrary, these are projects where you have no idea whether you can rely on them or not. In the past, if you stumbled upon a more extensive project with thousands of lines of code, there was a certain assumption that if someone went to the trouble of creating something like that, they would have some knowledge of the problem, a personal connection to their creation, and some willingness to maintain it going forward.
You can no longer rely on this at all. Today, you can generate a project with several thousand lines of code in a matter of moments. It could be complete nonsense or even something dangerous; it could be something functional that someone generated for their own immediate needs and posted to GitHub, but with no interest in turning it into an open-source project. Because a repository with code doesn’t make an open-source project. The difference between a piece of code on GitHub and an open-source project is that an open-source project solves problems and use cases for its users, not just the author’s one-off need. And most authors of such quick-and-dirty code simply aren’t interested in doing that.
This is clearly visible in projects like MeshCore , for instance. There are dozens of forks of everything imaginable. Missing a feature in the official MeshCore firmware? You just fork it, vibe-code the missing piece, and dump it on GitHub as MeshCore-UltimateEdition . The problem is that it was created with minimal effort, the author usually has no relationship to it, gets bored after a month, and it becomes abandonware before it even has a chance to age.
About ten years ago, people started saying that the concept of Linux repositories had run its course. In the 2000s, they were practically the only source of Linux software. If a project didn’t make it into distribution repositories, it had a problem. But then the number of open-source projects grew at such a rate that distributions couldn’t keep up. Users had to start getting their software elsewhere, and software authors learned to do without distributions. Just a few years ago, the “everything I need, I find in Debian” approach seemed definitively dead.
However, it is possible that curated software sources – like Linux distribution repositories – will make a comeback. The open-source software world is becoming so chaotic that users will once again start appreciating sources containing curated software that someone has vetted for them and that they can rely on six months down the road.
Another trend that AI has triggered in open source is ‘review overwhelm’. Previously, writing code acted as a natural filter because it required a non-trivial amount of effort and time investment. That is now gone, making code creation fast and easy. But someone still has to review this code before it goes into serious production. The review processes that worked in open-source projects for years are now at their capacity limits.
In GNOME 50, support for Google Drive was removed because nobody had been maintaining it for a long time. Users were naturally unhappy about it, and eventually, one user stepped up, re-added the support, and submitted it upstream to the gvfs project .
A colleague responsible for maintaining that project lamented that it was a change involving 4,000 lines of code. Even though it seems to work at a basic level, it was clearly generated using AI. He will still have to go through it line by line to verify that it actually works as intended and meets the code quality standards required to commit to maintaining it long-term.
Most of the effort has thus shifted from code creation to code review, which is typical for AI. The problem in open source, however, is that developers experienced enough to review and merge code were already a bottleneck before AI. Now, the problem has deepened significantly. And in the example above, my colleague can count himself lucky that the contributor is responsive and has shown long-term interest in the issue.
Today, that is more of a rare exception. Common contributions consist of someone wildly vibe-coding something without any deeper interest or understanding of the subject, and throwing it over the wall to the maintainers.
I have a fairly recent experience with this in Meshy . Someone submitted a pull request with 9,000 lines of changed code, which was supposed to add support for macOS. I spent an hour one evening doing a very quick review, and even during that short time, I ran into numerous issues: the code was blatantly AI-generated, several thousand lines were just completely useless replacements of single quotes with double quotes, parts of the code unrelated to the problem were modified, and it overwrote all the changes I had made in the main branch over the last few weeks.
The author never responded to my comments and I never heard from him again. My takeaway was that even that one hour was too big of a time investment for contributions like that, and next time I will reject them much faster.
Some projects are responding to this situation by tightening basic contribution requirements. For example, Flathub’s decision to reject AI-generated apps caused quite a stir. Many people criticized it as shooting themselves in the foot, but you have to look at their reality.
Flathub currently hosts several thousand apps, with more added every day. Only three people handle the reviews. Although their review process is highly automated, they do it very thoroughly, and a lot of manual input is still required. It’s clear their goal isn’t just to spot the worst slop, but to maintain a relatively high standard of code hygiene. In the last six months, I submitted two apps to Flathub, and the review process ultimately contributed to improving the quality of the apps themselves.
However, this has now clashed with the reality of people submitting completely vibe-coded apps without a shred of personal effort. The ticket template for requesting inclusion asks a few questions, including a requirement to upload a short video showing how the app works. It really isn’t demanding, and anyone can put it together in 15 minutes. Yet even that is too much effort for creators of AI slop.
Instead of fulfilling these minimal requirements, some labeled it an attack on Linux’s freedom and immediately vibe-coded an alternative to Flathub that was supposed to be open to everyone. Unsurprisingly, it barely lasted a month.
Not only do open-source maintainers lack the capacity to satisfy this demand for code review, but they are also losing the motivation to do it. Often, it would be faster for them to write the feature themselves, but the review process was historically how they cultivated new long-term contributors and potential successors. When someone sends you a vibe-coded contribution that cost them zero effort and which they likely don’t even understand, how do you expect to mentor them into a contributor who will help the project in the long run?
Open-source software was never just about the end result; it was also about the process – where contributors build a relationship with the project and grow into someone who will eventually pass that on to others. This stands in sharp contrast to the world of AI, where it’s all about the result. As fast as possible, with as little effort as possible.
Declining Motivation to Publish Code
In the 1990s, Francis Fukuyama declared democracy and liberal economics to be the ultimate victors in the arrangement of the world order. Today, as democracy erodes globally and the existing economic order crumbles, that looks like a prematurely bold statement to say the least. Similarly, just a few years ago, impressed by the developments of the last few decades, some hailed open source as the ultimate winner among software development models. Are we about to face a sobering reality check similar to Fukuyama’s thesis?
Lately, I’ve been observing a subtle, yet present trend of stepping back from open-source development. One argument against open development I hear concerns the aforementioned review overload. For some projects, the costs associated with being overwhelmed by AI slop can outweigh the benefits of useful community contributions. They might still publish the source code for transparency’s sake, but they transform from an open-development project into an open-source, closed-development project. And those who don’t care as much about transparency may close off the source code entirely.
Another argument against making source code public is the fear of license circumvention. Today’s LLMs train on source code regardless of its license and can then easily generate a similar solution that you can publish under whatever license you choose.
This isn’t an issue for permissive licenses, as the author has already accepted that anyone can do practically whatever they want with the code. However, AI poses a direct threat to copyleft licenses like the GNU GPL. Authors usually choose these to ensure their work remains open forever and that anyone who uses it shares their improvements back with the community. If an LLM trains on a project you’ve worked on for years and then generates a very similar solution published under a proprietary license, it effectively bypasses this principle.
Take MeshCore again as an example: the protocol itself and the firmware are open-source, but the clients are closed. Recently, it came to light in the community that a core team member secretly applied for the MeshCore trademark and started vibe-coding his own closed-source solutions based on the available code. MeshCore founder Scott Powell cited this as something that reaffirmed his decision to keep the client source code private. Specifically, he wrote:
So, I see open source, in the age of AI, as offering up your blood, sweat and tears for others to rip-off, but in innumerable ways.
We may disagree with Powell’s perspective, but it represents a legitimate stance that I see more and more often around me. I see lifelong open-source advocates – people who used to publish every last helper script because they wanted to share – who now keep those things to themselves, offering them to others only upon request. They have reasons similar to Powell’s.
Open source also grew out of the need to share. Writing code was hard; maintaining it was even harder. Why should everyone implement the same thing independently? Let’s join forces in an open-source project, write a shared library, and everyone can benefit from the results. The infrastructure powering the Internet today was built on this foundation. But AI is suppressing this need.
For instance, I encounter opinions that WordPress is dead because “I can just easily generate my own CMS.” In my view, that severely underestimates what an open-source project actually provides. It is so much more than just writing code, and this strategy of swapping a dependency on an open-source project for a dependency on an LLM might not pay off in the long run.
Nevertheless, the reliance on shared open-source components has indeed decreased to some extent. AI might not replace everything, but why depend on a large e

[truncated]
