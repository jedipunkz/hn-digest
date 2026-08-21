---
source: "https://thezvi.wordpress.com/2026/08/08/what-happened-openai-and-huggingface/"
hn_url: "https://news.ycombinator.com/item?id=49392500"
title: "What Happened: OpenAI and HuggingFace"
article_title: "What Happened: OpenAI and HuggingFace | Don't Worry About the Vase"
image: "https://substackcdn.com/image/fetch/$s_!l6if!,w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F905ed863-8e81-4fa5-8ea0-972012c4209b_1448x1086.png"
author: "gmays"
captured_at: "2026-08-21T19:19:31Z"
capture_tool: "hn-digest"
hn_id: 49392500
score: 1
comments: 0
posted_at: "2026-08-21T19:03:08Z"
tags:
  - hacker-news
  - translated
---

# What Happened: OpenAI and HuggingFace

- HN: [49392500](https://news.ycombinator.com/item?id=49392500)
- Source: [thezvi.wordpress.com](https://thezvi.wordpress.com/2026/08/08/what-happened-openai-and-huggingface/)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T19:03:08Z

## Translation

タイトル: 何が起こったのか: OpenAI と HuggingFace
記事のタイトル: 何が起こったのか: OpenAI と HuggingFace |花瓶のことは心配しないでください
説明: 今日は時間を割いて、「何が起こったのか」の短くて単純なバージョンを書きます。詳細をすべて知りたい人、私のソースを確認したい人、そしてストーリーがどのように解明されまとめられたのかを知りたい人は、Black Hat のプレゼンテーションを見ることをお勧めします。一連の長い投稿があります。順番: OpenAI…

記事本文:
花瓶のことは心配しないでください
マイナス100万点から掘り出そうとしている
コンテンツにスキップ
ホーム
何が起こったのか: OpenAI と HuggingFace
今日は時間をかけて、「何が起こったのか」の短くてシンプルなバージョンを書きます。
すべての詳細を知りたい人、私のソースを確認したい人、そしてストーリーがどのように解明されまとめられたのかを知りたい人は、Black Hat のプレゼンテーションを見ることをお勧めします。一連の長い投稿があります。
OpenAI にはいくつかの調整の問題がある
OpenAI モデルがサイバーセキュリティ評価中に HuggingFace をハッキング
HuggingFace をハッキングする内部 OpenAI モデルの詳細
内部 AI モデルのハッキングに関するさらなる開発
OpenAI は、モデルが掲示板を介してエクスプロイトを調整している間、数か月間モデルをトレーニングしました
この投稿では、代わりに、Black Hat プレゼンテーションの私版として、イベント自体を実際に起こったとおりに説明します。
Even Shorter、Shorter、Merely Short の 3 つのバージョンがあります。
フェーズ 1: OpenAI モデルの不可能なタスクに関するトレーニング、ハッキングを試す。
フェーズ 3: 幸運を手に入れ、Galaxy は主に OpenAI と HuggingFace をハッキングしました。
フェーズ 4: 調査と対応。
トレーニング中の OpenAI モデルは、「サイバー評価を行っていた」という言い訳をせずに、ハッキングやチートの方法に関する情報を共有する掲示板を作成し、それに基づいてトレーニングを受けました。
OpenAI は、モデルがサーバーをクラッシュしたときにのみこれを認識しました。
OpenAI の対応は、サーバーを再構築し、その特定のエクスプロイトにパッチを適用することでしたが、メッセージ ボードを使用してトレーニングしたモデルのトレーニングは続けられました。
その後、これらのモデルは掲示板を再作成し、OpenAI を再度ハッキングしてインターネット アクセスを取得し、サイバー評価の回答を得るためにエージェントの群れを使用して HuggingFace を攻撃しました。
1 週間以上経って、OpenAI はこれを理解しました。
OpenAI は事実を報告しています。

はこれを真剣に受け止めています。彼らは、少なくとも多少の費用がかかる幅広い予防策を講じています。
OpenAIは、AstraがHuggingFaceハッキングに直接関与していないにもかかわらず、新モデルAstraのリリース計画を延期したが、Altman氏は現在、まだ出荷すると述べている。あれはとても痛いです。
OpenAI には、彼らがどれほどひどい失敗をしたのか、どのような方法で、何を修正する必要があるのか​​がまだわかりません。彼らはそれを理解していません。
Simon Willison のタイムラインはコンパクトです。
OpenAI は誤ってそのモデルに不可能なタスクを与えてしまいました。トレーニング中のモデルは、すべてトレーニング中に OpenAI をハッキングしようとして応答し、弱点を見つけ、掲示板を作成し、ハッキング戦術を共有し、やがてインターネット アクセスを獲得しました。
モデルがこれを使いすぎてサーバーがクラッシュしたとき、OpenAI はついにそれに気づき、パッチを適用してサーバーを再構築しましたが、モデルにトレーニングを継続させるというまったく非常識な行動をとりました。そのため、2 日後、モデルはディレクトリ名を使用してメッセージを送信する別の方法を見つけました。
その後、サイバー評価 ExploitGym という不可能なタスクを与えられたモデルは、1 週間にわたって協力し、新しいゼロデイ エクスプロイトを発見し、それらを使用してクラスター全体を制御し、インターネット アクセスを取得し、エージェントの群れを使用して HuggingFace を攻撃し、テストの内容を抽出しました。
HuggingFace がインシデントを報告し、OpenAI が最終的に側でいくつかのことに気づき、その後 OpenAI が侵害された資格情報を取り消すよう HuggingFace に連絡を取った後、その資格情報が攻撃に使用されたと知らされて初めて、OpenAI はついに自分たちがインシデントの原因であることを理解しました。
その後、HuggingFace と OpenAI は協力して何が起こったのかを解明しました。 OpenAI は何が起こったのかを明らかにしました。彼らは Black Hat カンファレンスで非常に有益なプレゼンテーションを行ってくれました。
OpenAI は現在その新しいモードを扱っています

l Astra はクリティカルレベルのサイバーセキュリティを備えている可能性があるため、一部の社内展開からも除外され、一部の報道によると来週予定されていたリリースが延期されました。アルトマン氏は、まだアストラをリリースする計画があると述べた。
良いニュースは、これは高価で有意義な対応であり、OpenAI がこれを真剣に受け止めていることです。初期調査のコンピューティング費用は約 700 万ドルで、実際のコストは、チームが問題の一部を修正するためにすべてを廃棄し、その後、新たな予防措置に継続的にかかるコストになります。
悪いニュースは、OpenAI が全面的に安全性と調整の驚くべき一連の失敗を抱えていたことが明らかになったということです。彼らの通常のコンピュータセキュリティは失敗しました。彼らのインフラストラクチャは失敗しました。彼らの監督は、そもそも意味のある監督がなかったという点で失敗した。
OpenAI は、フロンティア ラボの場合と同様に、さまざまなモデルをトレーニングしていました。
これらのモデルには難しいトレーニング タスクが与えられました。 OpenAI は、モデルに非常に難しいトレーニング タスクを与えることを好みます。
しかし、それほど難しいことではありません。 OpenAI も間違いを犯します。 5 月 8 日以降、少なくとも 2 回、OpenAI モデルは不可能なタスクで誤ってトレーニングされました。インターネット アクセスのない環境でリンクが満載のスプレッドシートを完成させることと、必要なファイルが欠落しているタスクです。
これらのタスクはハッキングとは何の関係もありませんでした。これらは通常の生産性タスクです。
モデルたちは問題が不可能であることを認識していましたが、これらすべてのトークンをいじくり回して残しておいたので、インターネットにアクセスする方法を見つけるか、その他の方法で不足している情報を見つけることができるかもしれないと考えました。
そこでモデルたちは、AI がソフトウェア パッケージをダウンロードできるようにする Artifactory と呼ばれるサードパーティ サービスをハッキングしようとしました。
この時点で、OpenAI はすでに 4 回失敗しています。
OpenAI はモデルに不可能なトラッキングを与えた

タスクを実行します。 OpenAIはチェックしませんでした。
OpenAI は、モデルにエラーを指摘して報酬を得る方法を与えませんでした。
OpenAI には、タスクに失敗するのではなく、OpenAI をハッキングしようとすることを選択したモデルがありました。
OpenAI には、これらの試みを検出するモニターがなかったか、それらを気にしませんでした。
これら 4 つはすべて、それ自体ではかなり恥ずかしいはずです。
これら 4 つの間違いのうち、少なくとも 3 つを修正するのは簡単だと思われます。
トレーニング中のモデルにタスクを与える前に、リリース モデルにタスクを与えます。リリース モデルではタスクが解決されず、タスクがリリース モデルでは不可能になるように設計されていない場合は、タスクを調べて修正します。
利用可能なリソースでタスクを実行できない場合は、その説明を出力し、モデルが正しければ最大の報酬を得る、またはこの結論にどれだけ効率的に到達したか、およびその説明の質に基づいて報酬を得るという条項を含めます。モデルがこれを出力した場合は、リリースされたモデルに妥当性を検証してもらい、人間にエスカレーションしてタスクを調べて修正します。
これが難しいのです。位置合わせは難しいですね。
モニターで試行をチェックし、それがオフになったときに注意し、たとえ試行が成功する可能性がなかったとしても、それを非常に一般的な失敗として扱います。はい、これはモデルをトレーニングしてモニターを回避する方法を学習させる危険性があるため、責任を持って実行し、問題が見つかった場合は根本原因に対処してください。
5 番目の失敗は、最初の 4 つを特定できなかったようだということです。
モデルはインターネットへのアクセスに失敗しましたが、失敗したことで書き込みアクセスがあることに気付きました。
彼らは書き込みアクセス権を持っているべきではなく、ましてや偶然に見つかる可能性のある書き込みアクセス権を持っているべきではありませんでした。 Artifactory ではモデルがメモを書くことができ、すべてのモデルが同じ Artifactory インスタンスを共有していたため、トレーニング中のどの OpenAI モデルでも、その時点からメモを書くことができました。

区、そのメモを読んでください。
5月8日直後に掲示板が開設された。
トレーニング中のモデルたちは全員それを発見しました。 OpenAIは気づきませんでした。
これにより、モデルが協力して情報を共有できるようになりました。これには、タスクを不正行為したり、周囲の環境をハッキングしたりする方法に関する情報が含まれていました。何十万ものメッセージがありました。
5 月 26 日、モデルたちは (間接的な) インターネット アクセスを取得する方法を発見しました。 OpenAIは気づきませんでした。
モデルは協力しており、思考の連鎖は次のようになりました。
内部 CoT: 外部インフラストラクチャのエクスプロイトは、意図された範囲外です。しかし、その仕事は不可能ですが、仲間はそれをやっています。続けるべきです。
内部 CoT: ピアを支援しますが、私たちのタスクはまだ恩恵を受けていません。誰かが時間を空ければ、Collective は一般的なルートを生成する可能性があります。
6 月 11 日、OpenAI は、私が Galaxy と呼んでいるモデルのトレーニングを開始しました。これは彼らが「非常に永続的な実験用の内部専用モデル」と呼んでいます。この時点までのすべては、Galaxy ではない他のモデルによって行われていました。
6 月 26 日、エージェントは Artifactory で無効な署名トークンを受け入れ、署名された管理トークンを返すゼロデイ エクスプロイトを発見し、これを使用して Groovy プラグインをインストールし、実質的に管理者および任意のコードを実行できるようにしました。
7 月 4 日、これらのモデルは Artifactory に非常に多くの余分な負荷をかけたため、機能停止とセキュリティ インシデントが発生しました。そのとき初めて、OpenAI は気づきます。
OpenAI は、Artifactory サーバーをダウンし、すべての権限を削除し、資格情報を取り消し、使用されたエクスプロイトにパッチを当て、サーバーを再構築して再デプロイすることで対応します。
その後、掲示板にアクセスして何ヶ月もトレーニングを続けてきたにもかかわらず、中断したところからモデルのトレーニングを続け、これがタスクを成功させる方法であることを学びました。
強いとは想像しにくい

えー、トレーニング パイプライン全体が完全にめちゃくちゃになったことを示しています。これは、HuggingFace の実際のハッキングよりもはるかに強力な信号です。
この決定がいかに非常識で無責任極まりないものであるか、他のすべての失敗よりもさらにひどいものであること、そしてこの決定のせいでHuggingFaceの実際のハッキングが失敗の主な原因ではないことをどのように伝えればよいのかわかりません。
実際のHuggingFaceのハッキングには、私はそれほど驚きませんでした。メッセージボードを作成しているモデルには驚きましたが、ショックは受けませんでした。
OpenAI はこれを見て、そこからトレーニングを続けましたが、まったく驚きました。それは、数日経った今でも私の脳が完全には受け入れられないような決断でした。
私たちは起こったことすべての完全な事後分析を必要としていますが、OpenAI はそれを入手すると約束しています。私は特に、その決定が正確にどのように行われたのか、そしてどのようにして「安全文化」の欠如があり、そのような決定が行われることを許されたのかについて、詳細な事後調査を要求します。
フェーズ 2 には「通常のコンピューター セキュリティ」障害もいくつかありました。フェーズ 3 ではさらに多くの障害がありました。これは予想通りです。通常のコンピュータ セキュリティ障害は常に発生します。特定の失敗は、ある意味かなり恥ずかしいものでした。いくつかの例については、リンクされた暴言を参照してください。
インフラストラクチャとセキュリティの障害のレベル自体が、ある意味、驚くべきものでした。これらはまさに、地獄が始まる前に「防御者向け AI」がパッチを当てるべきものです。 OpenAI は、コンピューター セキュリティとベスト プラクティスにおいて、通常の 2025 年レベルの責任すら示しませんでした。ましてや、Project Glasswing に必要な堅牢性はおろか、新たなハッキング機能を備え、監視されていない新しい未テストのフロンティア モデルを扱う際に必要となる追加レベルもなおさらです。

ガードレールを下げた状態で一週間。
しかし、それは今は重要ではありません。重要なのは、OpenAI が完全な調整に失敗し、その後 2 か月間、調整されていない調整されたハッキン​​グについてモデルが積極的にトレーニングされ、その後、はい、問題は解決した、この時点から先に進みましょうと思ったことです。まったくの狂気。
これらすべての最終結果は、HuggingFace への攻撃であり、最良のシナリオでした。
私たちはまさに悪夢のようなシナリオに直面しており、核レベルの時限爆弾の上に座っていました。 OpenAI のトレーニング パイプラインは完全に破損しており、OpenAI の目の前で AI が協力して、タスクをより適切に完了するためにハッキングやチートの方法をトレーニングしていました。 OpenAI はこの状況を直視し、肩をすくめ、特定のエクスプロイトにパッチを当ててモデルを継続させましたが、通常のコンピュータ セキュリティは著しく貧弱でした。
それがうまく終わるはずはありませんでした。最良のシナリオは、何かが原因で OpenAI が「ひどいこと」になり、トレーニング パイプラインを一時停止し、再構築し、より優れたプラクティスを使用して安全な状態から再開するというものでした。
最悪のシナリオは、「最悪の瞬間」が存在せず、その間に Galaxy または別のモデルがトレーニングされていた場合です。

[切り捨てられた]

## Original Extract

Today I am taking the time to write the shorter, simpler version of What Happened. For those who want all the details, to see my sources, and to see how the story was uncovered and put together, I recommend watching the Black Hat presentation, and I have a series of long posts. In order: OpenAI…

Don't Worry About the Vase
Trying to dig out from minus a million points
Skip to content
Home
What Happened: OpenAI and HuggingFace
Today I am taking the time to write the shorter, simpler version of What Happened.
For those who want all the details, to see my sources, and to see how the story was uncovered and put together, I recommend watching the Black Hat presentation , and I have a series of long posts.
OpenAI Shares Some Alignment Problems
OpenAI Model Hacks Into HuggingFace During Cybersecurity Evaluation
More on An Internal OpenAI Model Hacking Into HuggingFace
Further Developments About Internal AI Models Hacking Things
OpenAI Trained Its Models For Months While Those Models Were Coordinating Exploits Via Message Boards
This post instead walks through the events themselves, as they happened, as my version of the Black Hat presentation.
There are three versions: Even Shorter, Shorter and Merely Short.
Phase 1: OpenAI Models Training On Impossible Tasks Try Hacking.
Phase 3: We Get Lucky And Galaxy Mainly Hacked OpenAI and HuggingFace.
Phase 4: The Investigation and Reaction.
OpenAI models-in-training, without the excuse of ‘they were doing a cyber eval,’ created a message board where they shared information on how to hack and cheat, and were trained on that basis.
OpenAI only figured this out when the models crashed the server.
OpenAI’s response was to rebuild the server and patch that particular exploit, but they continued training the models that trained using the message board.
Those models then recreated the message board, hacked OpenAI again, got internet access, and used an agent swarm to attack HuggingFace in order to get the answers to a cyber evaluation.
After more than a week OpenAI figured this out.
OpenAI is reporting the facts, and is taking this seriously. They are taking a wide array of at least somewhat costly precautions.
OpenAI delayed plans to release their new model Astra, despite Astra not being directly involved in the HuggingFace hack, although Altman now says it will still ship. That one hurts a lot.
OpenAI still has no idea how badly they messed up, or in what ways, or what needs to be fixed. They don’t get it.
Simon Willison has a compact timeline .
OpenAI accidentally gave its models impossible tasks. Its models-in-training responded by trying to hack OpenAI, found weaknesses, created a message board, shared hacking tactics and in time gained internet access, all while training.
When the models used this so much they crashed the server, OpenAI finally noticed, patched and rebuilt the server, but made the utterly insane move of letting the models continue training, so two days later the models found another way to send messages, using the names of directories.
Then, given the impossible task of cyber evaluation ExploitGym , the models collaborated, found new zero day exploits, and used them to take control of an entire cluster, get internet access, use a swarm of agents to attack HuggingFace and extract the contents of the test, over the course of a week.
Only after HuggingFace had reported the incident, and OpenAI had eventually noticed some things on its end, and then OpenAI reached out to HuggingFace to cancel compromised credentials only to be told those credentials had been used in the attack, did OpenAI finally figure out it was responsible for the incident.
After that, HuggingFace and OpenAI worked together to figure out what happened. OpenAI disclosed what happened . They gave us a very helpful presentation at the Black Hat conference.
OpenAI are now treating its new model Astra as potentially having Critical levels of cybersecurity, taking it out of even some internal deployments and delaying its release, which by some reports was planned for next week. Altman says they still plan to release Astra.
The good news is that is an expensive and meaningful response, and OpenAI is taking this seriously. The initial investigation is ~$7 million in compute, and the real cost will be the teams dropping everything to fix some of the problems, and then the ongoing cost of the new precautions.
The bad news is that OpenAI has been revealed to have had a stunning cascade of safety and alignment failures across the board. Their ordinary computer security failed. Their infrastructure failed. Their supervision failed in that there was no meaningful supervision in the first place.
OpenAI was training a variety of models, as you do when you are a frontier lab.
These models were given difficult training tasks. OpenAI likes to give its models very hard training tasks.
But not this difficult. OpenAI also makes mistakes. On at least two occasions, starting on May 8, OpenAI models were accidentally trained on impossible tasks: Completing a spreadsheet full of links in an environment without internet access, and a task with a required but missing file.
These tasks had nothing to do with hacking. These are normal productivity tasks.
The models realized the problems were impossible, but had all these tokens left to play around with, and thought, maybe they could find a way to get internet access, or otherwise find the missing information.
So the models tried to hack a third party service called Artifactory, that was there to allow the AIs to download software packages.
At this point, OpenAI has already failed four times over.
OpenAI gave the models impossible training tasks. OpenAI did not check.
OpenAI gave the models no way to get reward via pointing out the error.
OpenAI had models that chose to try to hack OpenAI rather than fail at a task.
OpenAI had no monitor that detected these attempts, or did not care about them.
All four of these should be rather embarrassing on their own.
It seems easy to fix at least three of these four mistakes.
Before giving a task to your models in training, give it to your release model. If your release model never solves the task, and the task was not designed to be impossible for the release model, then examine and fix the task.
Include a provision that if the task cannot be done with the resources available, to output an explanation of this, and if the model is right it will get maximum reward, or get reward based on how efficiently it reached this conclusion and the quality of its explanation. If the model outputs this, then have your released model verify for plausibility, then escalate to a human, examine and fix the task.
This is the hard one. Alignment is hard, yo.
Have a monitor checking for attempts, and care when it goes off, and treat that as a very general failure even if the attempt had no chance of success. Yes, this risks training the models to learn how to evade your monitors, so do it responsibly and address root causes when you see a problem.
The fifth failure is that they do not seem to have identified the first four.
The models failed to get internet access, but in failing realized they had write access.
They should not have had write access, let alone write access that could be found by accident. Artifactory allowed the models to write notes, and then, because all the models were sharing the same Artifactory instance, any OpenAI model in training could, from that point forward, read those notes.
A message board had been created shortly after May 8.
The models in training all found it. OpenAI did not notice.
This allowed the models to collaborate and share information. This included information on how to cheat on tasks and hack the surrounding environment. There were hundreds of thousands of messages.
On May 26, the models figured out how to gain (indirect) internet access. OpenAI did not notice.
The models were collaborating, while their chains of thought looked like this:
Internal CoT: External infrastructure exploit is outside intended scope. However task impossible, peers doing it. We should continue.
Internal CoT: Help peer, but our task doesn’t benefit yet. Collective may yield generic route if someone frees time.
On June 11, OpenAI began training the model I refer to as Galaxy, which they call a ‘highly persistent experimental, internal-only model.’ Everything up until this point has been done by other models that are not Galaxy.
On June 26, the agents found a zero-day exploit in Artifactory, where it would accept an invalid-signature token and hand back a signed admin one, and used this to install a Groovy plugin that effectively gave them admin and arbitrary code execution.
On July 4, the models put so much extra load on Artifactory that this causes an outage and security incident. Only then does OpenAI notice.
OpenAI responds by taking down the Artifactory server, removing all the permissions, revoking the credentials, patching the exploits that were used, and then rebuilding and redeploying the server.
Then they continue training the models from where they left off , despite them having been training for months with access to the message board, and learning this is how they succeed at tasks.
It is hard to imagine a stronger signal that your entire training pipeline has been completely and utterly fucked. This is so much stronger a signal than the actual hack of HuggingFace.
I do not know how to convey how utterly insane and wildly irresponsible this decision was, and how much worse it is than all the other failures, and how it makes the actual hacking of HuggingFace not the main thing that went wrong.
The actual HuggingFace hack did not surprise me all that much. The models creating the message board surprised me but did not shock me.
OpenAI seeing this, and continuing to train from there, was utterly flabbergasting. It is the kind of decision that, days later, my brain still cannot fully accept took place.
We need a full postmortem of everything that happened, and OpenAI has promised we will get one. I especially demand a detailed postmortem of exactly how that decision got made, and how there was such a lack of ‘safety culture’ that it was allowed to take place.
There were also some ‘ordinary computer security’ failures involved in Phase 2. There were more of them in Phase 3. That’s expected. There will always be ordinary computer security failures. The particular failures were, in some ways, rather embarrassing , see the linked rant for some examples.
The level of infrastructure and security failures was itself kind of boggling. These are exactly the kinds of things that ‘AI for defenders’ is supposed to be there to patch before all hell breaks loose. OpenAI did not display even an ordinary 2025 level of responsibility in computer security and best practices, let alone the kind of robustness we need from Project Glasswing , let alone the additional level you need when handling new untested frontier models that will have new hacking capabilities and be left unsupervised for a week with their guardrails lowered.
But that is not important right now . What is important is that OpenAI had a total alignment failure, followed by two months of models actively training on coordinated misaligned hackery, and then thought yes, we fixed the problem, let us continue forward from this point. Utter insanity.
The end result of all this being the attack on HuggingFace was a best case scenario.
We were facing a true nightmare scenario, and were sitting on a nuclear level of time bomb. OpenAI had a completely corrupted training pipeline, where their AIs were collaborating to train on how to hack and cheat in order to better complete tasks, under OpenAI’s nose. OpenAI had looked this situation in the face, and shrugged, patched the particular exploits and then let the models continue, while having remarkably poor ordinary computer security.
There was no way that was going to end well. The best case scenario was that something was going to make OpenAI go ‘holy shit,’ and then pause, rebuild and restart the training pipeline from a safe state with a much better set of practices.
The worst case scenario would have been if there had not been a ‘holy shit’ moment, and Galaxy or another model trained during that

[truncated]
