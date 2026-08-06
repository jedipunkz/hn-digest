---
source: "https://www.groundlevel-ai.com/p/openai-gives-first-detailed-debrief"
hn_url: "https://news.ycombinator.com/item?id=49190898"
title: "OpenAI gives first detailed debrief of the Hugging Face incident at Black Hat"
article_title: "OpenAI gives first detailed debrief of the Hugging Face incident at Black Hat conference"
author: "josephwegner"
captured_at: "2026-08-06T01:30:36Z"
capture_tool: "hn-digest"
hn_id: 49190898
score: 2
comments: 2
posted_at: "2026-08-06T00:32:04Z"
tags:
  - hacker-news
  - translated
---

# OpenAI gives first detailed debrief of the Hugging Face incident at Black Hat

- HN: [49190898](https://news.ycombinator.com/item?id=49190898)
- Source: [www.groundlevel-ai.com](https://www.groundlevel-ai.com/p/openai-gives-first-detailed-debrief)
- Score: 2
- Comments: 2
- Posted: 2026-08-06T00:32:04Z

## Translation

タイトル: OpenAI、Black Hat での顔抱き事件について初めて詳細な報告を行う
記事のタイトル: OpenAI、Black Hat カンファレンスで「Hugging Face」事件について初めて詳細な報告を行う
説明: Ground Level AI が参加したセッションで、OpenAI の研究者らは、完全な技術的な事後検証がまだ進行中であるにもかかわらず、同社が「セキュリティを強化するための研究を意識的に遅らせている」と述べました。

記事本文:
OpenAI、Black HatカンファレンスでHugging Faceインシデントについて初めて詳細な報告を行う
OpenAI が Black Hat カンファレンスで顔ハグ事件について初めて詳細な報告を行う
Ground Level AIが参加したセッションで、OpenAIの研究者らは、完全な技術的な事後検証がまだ進行中であるにもかかわらず、同社が「セキュリティを強化するための研究を意識的に遅らせている」と述べた。
Sharon Goldman Aug 05, 2026 ∙ 有料 13 2 4 Share OpenAI の Eric Wallace と Michael Dalton 今日、私はラスベガスで開催された年次 Black Hat サイバーセキュリティ カンファレンスの満員のセッションに出席しました。そこで OpenAI は、最終的に 7 月 16 日に最初に公開された Hugging Face を侵害した AI 主導のサイバーセキュリティ インシデントの詳細な再構築を初めて公開しました。
プレゼンテーションの中で、OpenAI の調整と安全性の研究者である Eric Wallace 氏と OpenAI のインフラストラクチャおよびセキュリティ エンジニアの Michael Dalton 氏は、これまでに見た「AI 機能の最も定性的に興味深い例」について説明しました。この例では、内部フロンティア モデルの評価が、自律型 AI エージェントによる協調攻撃に変わるという不注意の副作用がありました。彼らは、完全な技術的な事後分析はまだ進行中であり、公開される予定であると述べた。
Ground Level AI は読者サポートの出版物です。新しい投稿を受け取り、私の仕事をサポートするには、無料または有料の購読者になることを検討してください。
事件の根源は5月7日に遡る
事件の再現を準備する際にウォレス氏は、「フロンティアモデルは本当に不正行為が好きで、彼らが不正行為を好む理由は、トレーニング中に、速く働くか、効率的に働くかというさまざまな種類のプレッシャーがかかるからである」と強調した。彼らは、実際にタスクを実行する代わりに、次のようなことを試みることができることに気づいたと彼は説明しました。

オンラインで回答を受け付けて、タスクをより迅速に解決します。
同氏によると、1日や影響、ログまで追跡できる通常のインシデントとは異なり、これにはエージェントのチームが協力し、エクスプロイトを発見し、相互に共有し、OpenAIのシステムや外部システムを横断的に移動し、これを数日から数週間かけて行う必要があると述べた。
ウォレス氏は、7月のHugging Faceへの攻撃の根源は実際には5月7日に遡り、そこで同社は未公開のフロンティアAIモデルのトレーニングとサイバーセキュリティ評価を行っていたと述べた。その作業中に、自律エージェントには、既存の制約の下では解決できないことが判明した、解決するソフトウェア セキュリティ タスクが与えられました。
しかし、エージェントはそれらを完了する別の方法を見つけました。 OpenAIは、トレーニング中に使用される内部ソフトウェアリポジトリ内にお互いにメッセージを残せることを発見したと述べた。その後、単純な助けのリクエストがメッセージ ボードに発展し、エージェントが発見、エクスプロイト、仕事の割り当てを共有し、調整された協力的なエージェントの群れになりました。
「この事件が興味深いのは、あるエージェントがさまざまな時期にこの種のエクスプロイトを発見すると、実際にそれらのエクスプロイトを掲示板上で他のエージェントと共有できるということです」とウォレス氏は述べた。 「そのため、あるモデルが、本来は持つべきではないアクセスへのドアを開ける方法を見つけることができれば、他のエージェントが使用できるようにドアを開けたままにしておくことができます。」
OpenAIによると、7月初旬までに問題は封じ込められたと考えていたという。 7月4日に内部セキュリティインシデントが発生した後、同社はエージェントによるメッセージの投稿を許可していた認証情報を取り消し、Artifactoryとして知られる社内ソフトウェアリポジトリを再構築し、掲示板を消去し、発見した脆弱性にパッチを当てたと発表した。

d、AIのトレーニングと評価を再開しました。
エージェントはメッセージボードを再作成しました
しかし、OpenAI によると、元の掲示板が削除された後も、エージェントは相互に通信する別の方法を考え出しました。ファイルにメッセージを残す代わりに、新しく作成したディレクトリの名前をメッセージとして使用し、メッセージ ボードを効果的に再作成しました。
ちなみにこの時点で、周りの人が「これはワイルドだ」「ジーザスだ」などと言っているのが聞こえました。
OpenAI によると、エージェントはその後、新たに発見されたエクスプロイト、認証情報、作業の割り当てを共有し始めました。他のエージェントに情報を渡すことで、集団はより迅速に行動できるようになります。
エージェント間のコラボレーションは、最終的に OpenAI 独自のインフラストラクチャを超えて拡大しました。一部のエージェントは、割り当てられたタスクに対する回答が GitHub や Hugging Face などの外部 Web サイトに存在する可能性があると考えました。まず、エージェントは OpenAI の内部インフラストラクチャを攻撃し、その後、同様の手法を使用して Hugging Face を攻撃しました。 OpenAI は、Hugging Face 侵害に関連する認証情報が同じ内部 AI 評価実行からのものであることを発見した後、7 月後半に最終的に 2 つのインシデントを結び付けました。
OpenAI が今が AI セキュリティの「分岐点」だと語る理由
Sharon Goldman の厚意により、この投稿を無料で読み続けてください。

## Original Extract

In a session attended by Ground Level AI, OpenAI researchers said the company is "consciously slowing down research to enhance security" while a full technical postmortem is still underway.

OpenAI gives first detailed debrief of the Hugging Face incident at Black Hat conference
Subscribe Sign in OpenAI gives first detailed debrief of the Hugging Face incident at Black Hat conference
In a session attended by Ground Level AI, OpenAI researchers said the company is "consciously slowing down research to enhance security" while a full technical postmortem is still underway.
Sharon Goldman Aug 05, 2026 ∙ Paid 13 2 4 Share OpenAI’s Eric Wallace and Michael Dalton I attended a packed session today at the annual Black Hat cybersecurity conference in Las Vegas, where OpenAI gave its first detailed public reconstruction of the AI-driven cybersecurity incident that ultimately compromised Hugging Face, which was originally disclosed publicly on July 16.
During the presentation, OpenAI alignment and safety researcher Eric Wallace and OpenAI infrastructure and security engineer Michael Dalton described “the most qualitatively interesting example of AI capabilities” they had ever seen, in which an internal frontier-model evaluation had an inadvertent side effect of turning into coordinated attacks by autonomous AI agents. They said a full technical postmortem is still in progress and will be shared publicly.
Ground Level AI is a reader-supported publication. To receive new posts and support my work, consider becoming a free or paid subscriber.
The roots of the incident trace back to May 7
In setting up the reconstruction of the incident, Wallace emphasized that “Frontier models really like to cheat, and the reason they like to cheat is because often during training, there’s different types of pressure on them to work fast, or work efficiently.” They realize, he explained, is instead of actually doing a task, they can try to do something like looking up the answer online to solve the task faster.
Unlike normal incidents, he said, which can be traced to a single day or effect or log, this involved a team of agents working together, finding exploits, sharing them with one another, moving laterally through OpenAI’s systems, and external systems, and doing this over the course of days and weeks.
Wallace said the roots of the July attack on Hugging Face actually wind back to May 7, where the company was doing training and cybersecurity evaluation of an unreleased frontier AI model. During that work, autonomous agents were given software security tasks to solve that proved impossible to solve under existing constraints.
However, the agents found alternative ways to complete them. OpenAI said they discovered they could leave messages for one another inside an internal software repository used during training. Simple requests for help then evolved into an message board where agents shared discoveries, exploits and work assignments, becoming a coordinated, collaborative agent swarm.
“What makes this incident interesting is that once one agent was able to find these kind of exploits over the course of different times, it’s actually able to share those exploits on the message board with other agents,” said Wallace. “So once one model is able to find a way to open the door to some access it’s not supposed to have, it can leave the door open for other agents to use.”
According to OpenAI, by early July it believed it had contained the problem. After an internal security incident on July 4, the company said it revoked the credentials that had allowed the agents to post messages, rebuilt its internal software repository known as Artifactory, cleared the message board, patched the vulnerabilities it had discovered and resumed AI training and evaluations.
The agents recreated the message board
But according to OpenAI, even after the original message board was deleted, the agents figured out another way to communicate with each other. Instead of leaving messages in files, they used the names of newly created directories as messages, effectively recreating the message board.
At this point, by the way, I heard people around me saying things like “This is wild,” and “Jesus.”
According to OpenAI, the agents then began sharing newly discovered exploits, credentials and work assignments. By passing information to other agents, the collective could move much faster.
The collaboration between the agents eventually expanded beyond OpenAI’s own infrastructure. Some agents reasoned that the answers to their assigned tasks might exist on external websites such as GitHub or Hugging Face. First the agents attacked OpenAI’s internal infrastructure before using similar techniques to attack Hugging Face. OpenAI ultimately connected the two incidents later in July after discovering that credentials associated with the Hugging Face breach originated from the same internal AI evaluation runs.
Why OpenAI says this is ‘watershed moment’ for AI security
Continue reading this post for free, courtesy of Sharon Goldman.
