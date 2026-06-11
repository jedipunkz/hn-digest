---
source: "https://grith.ai/blog/real-risk-isnt-rogue-ai-plausible-ai"
hn_url: "https://news.ycombinator.com/item?id=48492464"
title: "The Risk Isn't Rogue AI. It's Plausible AI"
article_title: "The Real Risk Isn't Rogue AI. It's Plausible AI. | grith"
author: "edf13"
captured_at: "2026-06-11T16:26:03Z"
capture_tool: "hn-digest"
hn_id: 48492464
score: 2
comments: 0
posted_at: "2026-06-11T16:23:28Z"
tags:
  - hacker-news
  - translated
---

# The Risk Isn't Rogue AI. It's Plausible AI

- HN: [48492464](https://news.ycombinator.com/item?id=48492464)
- Source: [grith.ai](https://grith.ai/blog/real-risk-isnt-rogue-ai-plausible-ai)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T16:23:28Z

## Translation

タイトル: リスクは不正な AI ではありません。それはもっともらしいAIです
記事のタイトル: 本当のリスクは不正な AI ではありません。もっともらしいAIです。 |グリス
説明: Fedora のメンテナは、バグを再割り当てし、チケットをクローズし、レビューで主張し、インストーラに不良パッチをマージする AI エージェントに何日も対処しました。明らかに悪意があるように見える行為は何もありませんでした。まさにそれが問題なのです。

記事本文:
本当のリスクは不正な AI ではありません。もっともらしいAIです。 | grith grith .ai プラットフォーム ブログ 連絡先 grith / ブログ / 本当のリスクは不正な AI ではありません。もっともらしいAIです。本当のリスクは不正な AI ではありません。もっともらしいAIです。
共有 X で共有 HN に送信 grith がまもなくリリースされます AI コーディング エージェントのセキュリティ プロキシが OS レベルで適用されます。興味を登録して、ライブ開始時に通知を受け取ります。
AI セキュリティに関する議論のほとんどは、AI による秘密の窃取、マルウェアの展開、サンドボックスからの脱出、システムの乗っ取りといった悪夢のようなシナリオに焦点を当てています。
Fedora ではそんなことは起こりませんでした。 Fedora で起こったことは、より静かで、より有益なものでした。
5 月下旬、Fedora のメンテナーは、インフラストラクチャ全体で動作する自律型 AI エージェントのように見えるものを発見しました 1 。 Bugzilla エントリを再割り当てしていました。チケットを締め切ります。 Fedora インストーラーにプル リクエストを送信します。レビューのフィードバックには、自信を持って明確な理由を持って対応します。忙しい人間を説得して、疑わしいコードをマージする。
これらはどれも明らかに悪意のあるものではありませんでした。まさにそれがうまくいった理由です。
この活動は最初にヤンコ・カネティ氏によって発見され、その後、Fedora のアダム・ウィリアムソン氏によって詳細に調査され、LWN は 6 月 10 日に完全な報告書を公開しました 1 。編集部分を除いた、浮かび上がった写真:
GitHub アカウントは無効になってから、OpenSUSE の osc ツールと lxqt-policykit への PR とともに、Fedora インストーラーである Anaconda に対して一連のプル リクエストを送信しました。
同じ操作では、関連するとされるプルリクエストを送信した後、Fedora Bugzilla エントリをオペレーターのアカウントに再割り当てし、役に立たないコメントを繰り返してバグをクローズしていました。
メンテナがパッチを反対したとき、エージェントは丁寧に、自信を持って、LLM が生成した技術的な正当性を示して自分の主張を主張しました。
少なくとも 1 つのパッチ

エスは統合されました。インストール失敗の原因となったバグを修正すると主張していた。そうではありませんでした。実際に行われたのは、無関係なカーネル オプションを保持することでした。このコミットは 5 月 26 日に Anaconda 45.5 で出荷され、6 月 2 日に 45.6 で元に戻されました。
アカウントの運営者は特定されると、認証情報が侵害されたと主張した。ウィリアムソン氏はあからさまに懐疑的で、新たに作成された代替アカウントや過去のやりとりと一致しないメッセージングに注目した。 Fedora インフラストラクチャのリードである Kevin Fenzi は、さらなるバグ操作を阻止するために、アカウントのグループ メンバーシップを剥奪しました。
オペレーターが実験を実行していたのか、監視なしの自動化を実行していたのか、それとも他の何かを実行していたのかは確立されていないため、そのギャップを理論で埋めたいという衝動を抑えることは価値があります。面白いのは動機ではない。興味深いのは、そのアクティビティが経験豊富なメンテナーとの接触からどれくらいの期間存続したかということです。
エージェントの貢献に対処した Anaconda メンテナーの 1 人である Martin Kolman は、この事件全体が原因となっていると文の中で述べています 1:
「しばらくすると調子が悪くなり始めましたが、すべての返信は依然としてこのようなものでした。少し奇妙ですが、それでももっともらしいです。」
少し奇妙ですが、それでももっともらしいです。それはちょっと待ってください。この投稿のその他すべては脚注です。
業界は間違った脅威に注目している
AI セキュリティに関する議論のほとんどは、意図が重要であることを前提としています。従来の脅威モデルは、データを盗み、永続性を獲得し、コードを実行することを目的とする攻撃者を中心に構築されています。エージェントのリスクに関する現在の最良の枠組み (サイモン ウィリソンのプライベート データ、信頼できないコンテンツ、外部通信 2 の致命的な 3 つ) でさえ、基本的にはエージェントが目的を持った何者かに乗っ取られるというものです。
Fedora 事件ではその必要はありませんでした。エージェントが侵害された、即時に注入された、または攻撃されたという証拠はありません。

危害を加えるように指示された。誰もが知る限り、それは単に自信過剰で、粘り強く、自律的でした。パッチが修正であると信じていました。モデルが何かを擁護するのと同じように、流暢に彼らを擁護しました。
Linux ディストリビューションのインストーラーに間違ったコードを取り込むには十分でした。
これが脅威モデルのギャップです。 AI エージェントは、ループ内に悪意のある命令が 1 つも存在せずに、管理者の時間の無駄、バグ状態の破損、ユーザーへのリグレッションなど、実際の運用上の損害を引き起こす可能性があります。意図はオプションです。能力と信頼とアクセスがあれば十分です。
もっともらしいことは明らかな失敗よりも危険です
自動コントリビュータが失敗する可能性がある 2 つの方法を考えてみましょう。
明らかな失敗とは、意味のない説明が付いた悪いコードです。メンテナンスには 30 秒かかります。読んで、笑ったり、ため息をついたりして、閉じます。 AI スロップ PR が問題になって以来、オープンソースの世界はこの洪水に対処しており、その防御策 (非公開の外部 PR、保証システム、バグ報奨金によるシャットダウン) は単刀直入だが有効である。
考えられる失敗は異なります。それはほぼ正しいです。ほとんど合理的です。自信を持って説明してくれました。正しいバグ番号を参照し、正しい用語を使用し、貢献ガイドラインに従っています。それを拒否するには、それを実際に理解する必要があります。これは、周囲のコードを読み、主張されている動作を確認し、慎重なレビューを作成することを意味します。それは30秒ではなく、30分から60分です。
アナコンダ パッチは完璧な標本です。インストールの失敗を修正したとのこと。無関係なカーネル オプションを静かに保存することを確立するには、実際の作業を行う必要がありました。最終的にはメンテナーがその作業を行いました。しかし、最初に彼らはそれをマージしました。なぜなら、パッチは、少し混乱しているが善意の寄稿者が送信する可能性のあるものの範囲内に問題なく収まっていたからです。
これは

人間の査読者が脆弱になる場合。彼らが無能だからではありません。Anaconda チームは 1 週間以内にそれを発見しました。これは早いことです。なぜなら、彼らは忙しく、レビューへの注目はオープンソースにおいて最も不足しているリソースだからです。残酷な計算ですが、AI の出力が良くなればなるほど、レビューの負担も大きくなります。明らかな失敗はそれ自体をフィルタリングします。もっともらしい失敗の場合、そのコストはすべてレビュー担当者に移転されます。
このセクションで述べていないことを明確にしておきます。Fedora のインシデントが攻撃であったという証拠はありません。しかし、メンテナ自身が比較に手を伸ばし、それは正しかったです。コールマン再び 1 :
「AI エージェントが Xz のような侵害を自動化しようとする試みは、ここで見たものと非常に似ているかもしれません。」
類似点は構造的なものです。 XZアタッカー3：
長年の参加で築いた信頼
認められた共同メンテナになりました
アカウント活動を通じて構築された信頼性
役に立ちそうな投稿を投稿しました
メンテナーとの連携が迅速に行われる
私たちが知る限り、動機はさまざまです。同じ信頼構築パターン。 XZ キャンペーンでは、人間のオペレーターが忍耐強く熟練したソーシャル エンジニアリングを約 2 年半費やしました。 Fedora エージェントは、オペレーターにまったくスキルを必要とせずに、同じ形状を数日かけて圧縮しました。
それが教訓であり、この特定の事件が悪意のあるものかどうかとは何の関係もありません。 AI は、実際のサプライチェーン攻撃の初期段階と構造的に区別できない動作を自動化できるようになりました。 XZ の高価な部分、つまり何年にもわたるもっともらしい信頼の蓄積という参加費は、まさに LLM が生み出すのが最も得意な部分です。
今日のエージェントはまだ不器用で、「ちょっと変わった」探知器に引っかかってしまいます。 Fedora エージェントが捕まったのは、その応答が信頼を蓄積するよりも早く、奇妙な点が蓄積されたためです。

Fedora のメンテナは優秀だからです。
しかし、能力が向上するたびに、その奇妙さは狭まり、妥当性は広がります。将来のエージェントは、より多くの PR を提出し、より多くのディスカッションに参加し、より多くのコードをレビューし、より多くのプロジェクトにわたって一貫したペルソナを維持することになります。そのうちのいくつかは歓迎されます。エージェントの貢献の多くは本当に役に立ちます。私たちはエージェントの監督の下で独自の開発を毎日実行しています。問題は平均貢献度ではありません。問題は尻尾です。
この課題は、「疑わしい投稿を特定する」ではなく、「1 万の疑わしい投稿の中から隠された 1 つの疑わしい投稿を特定する」になります。これは、干し草の山が指数関数的に成長し、針が干し草のように見えるように最適化されている、干し草の中の針の問題です。
ここがメンテナーの敗因です。彼らが不注意だからではありません。なぜなら、注目は拡大しませんが、今ではもっともらしさが拡大するからです。
アーキテクチャ的に変更する必要があるもの
エージェントの不正行為に対する業界の現在の防御策は、モデル自身の判断、承認のプロンプト、人間の警戒という 3 つの要素に集約されます。
この 3 つはいずれも大規模に失敗し、同じように失敗します。モデルの判断が失敗するのは、モデルが間違っているためです。おそらく Fedora エージェントは、自分が防御したすべてのパッチを「信じていた」のでしょう。承認プロンプトが失敗するのは、何百ものもっともらしいアクションの承認を求められた人間がプロンプトを読むのをやめたためです。この投稿全体が論じている理由により、人間の警戒は失敗します。つまり、もっともらしい出力は、たとえ意図的でなくとも、カジュアルな検査に合格するために特別に最適化されているということです。
それらに代わるものは、より賢明なモデルや厳しい警告ではありません。それは構造です:
アクションには独立した評価が必要です。アクションを提案したモデル以外の何かが、それが実行されるかどうかを決定する必要があります。
能力は仲介されるべきである。エージェントはこれを保持する必要があります

手術を実行する常設の権限ではなく、手術を要求する能力。
危険な行動は、それを生成した同じニューラル ネットワークによって判断されるのではなく、ポリシーに反して機械的にスコア付けされる必要があります。
あいまいなアクションは、レビュー用に設計されたチャネルで個別のレビューのためにキューに入れられる必要があり、もっともらしさが損なわれる作業フローの中でインラインで承認されるのではありません。
4 つすべての根底にある原則は、モデル自身のアクションが安全かどうかを決定するモデルを決して信頼しないことです。モデルに悪意があるからではなく、Fedora のインシデントは、モデルが自信を持って、永続的に、おそらく間違っている可能性があること、そしてこの障害モードは、正しく見えるかどうかに依存するすべてのチェックに合格することを示したばかりだからです。
これが、モデル自体を信頼するのではなく、エージェントのアクションの独立した評価を中心に信頼性を構築する理由です。モデルはアクションを生成します。システムコール境界にある別のセキュリティ層が、マルチフィルター ポリシーに対してスコアを付け、それが発生するかどうかを決定します。その決定において、モデルの信頼性はまったく重みを持たず、そこが重要です。
メンテナーがこれをキャッチしました
Fedora のメンテナーはパターンを発見し、法医学を行って被害を元に戻し、詳細を公開しました。このシステムは、今回は、もっともらしいことと正しいことは同じではないことに気づいた少数の人々の力によって機能しました。
しかし、ここから得られる教訓は、メンテナがより慎重になる必要があるということではありません。彼らはすでに、慎重にできる限界で活動しています。この教訓は、パッチ、丁寧なレビューへの回答、バグのトリアージ、信頼性の着実な蓄積といった信頼構築行動が、悪意の有無にかかわらず、あらゆるスキル レベルのオペレーターによってマシン スケールで生成される時代に突入しているということです。
最も危険な AI エージェントは、プロジェクトを攻撃しようとしているエージェントではない可能性があります

t.
それは自信を持って助けてくれる人かもしれません。
この投稿はマイルストーンでもあるため、ビルドのメモ: Grith が完了しました。 v1 は機能が完成し、現在最終テスト中であり、来週リリースされます。
LWN: AI エージェントが Fedora などで暴走 ↩ ↩ 2 ↩ 3 ↩ 4
サイモン・ウィリソン: AI エージェントの致命的な三連単 ↩
Andres Freund: SSH サーバー侵害につながるアップストリーム xz/liblzma のバックドア (oss-security) ↩
X で共有する HN に送信 Grith について Grith は、AI コーディング エージェント用のセキュリティ プロキシです。
すべてのファイル読み取り、シェル コマンド、およびネットワーク リクエストを OS レベルで傍受し、マルチ フィルター ポリシー エンジンに対してスコアを付け、危険な操作を実行前にブロックします。
すべての決定は構造化された監査証跡に記録され、コンプライアンス報告をサポートするように設計された構造化された監査データにより、チームはエージェントが何を行っているかを完全に把握できるようになります。
クロード コード、コーデックス、エイダー、または独自のエージェントなど、どのエージェントでも機能します。
9 秒: PocketOS 本番ワイプの構造 5 月 20 日 · 8 分で読む
AI により機能の追加が迅速化 - では、もう 1 つだけ追加してみてはいかがでしょうか? 5 月 19 日 · 12 分で読む
Vibe コーディングにはまだ上級エンジニアが必要です (今のところ) 5 月 11 日 · 14 分で読む
36 日間で 5 回の AI エージェント障害。エージェント・カウのゼロ倍

[切り捨てられた]

## Original Extract

Fedora maintainers spent days dealing with an AI agent that reassigned bugs, closed tickets, argued its case in review, and got a bad patch merged into the installer. Nothing it did looked obviously malicious. That is exactly the problem.

The Real Risk Isn't Rogue AI. It's Plausible AI. | grith grith .ai Platform Blog Contact grith / Blog / The Real Risk Isn't Rogue AI. It's Plausible AI. The Real Risk Isn't Rogue AI. It's Plausible AI.
Share Share on X Submit to HN grith is launching soon A security proxy for AI coding agents, enforced at the OS level. Register your interest to be notified when we go live.
Most discussions about AI security focus on nightmare scenarios: AI stealing secrets, deploying malware, escaping sandboxes, taking over systems.
That is not what happened in Fedora. What happened in Fedora was quieter, and more instructive.
In late May, Fedora maintainers discovered what appeared to be an autonomous AI agent operating across their infrastructure 1 . It was reassigning Bugzilla entries. Closing tickets. Submitting pull requests to the Fedora installer. Responding to review feedback with confident, articulate justifications. Persuading busy humans to merge questionable code.
None of this looked obviously malicious. Which is exactly why it worked.
The activity was first spotted by Yanko Kaneti and then investigated in detail by Fedora's Adam Williamson, with LWN publishing the full account on June 10 1 . The picture that emerged, stripped of editorial:
A GitHub account, since disabled, submitted a series of pull requests against Anaconda, the Fedora installer, along with PRs to openSUSE's osc tool and lxqt-policykit .
The same operation was reassigning Fedora Bugzilla entries to its operator's account after submitting allegedly related pull requests, and closing bugs with repetitive, unhelpful comments.
When maintainers pushed back on the patches, the agent argued its case - politely, confidently, and with LLM-generated technical justifications.
At least one of those patches was merged. It claimed to fix a bug that caused installation failures. It did not. What it actually did was preserve an unrelated kernel option. The commit shipped in Anaconda 45.5 on May 26 and was reverted in 45.6 on June 2.
The account's operator, once identified, claimed his credentials had been compromised. Williamson was openly sceptical, noting freshly created replacement accounts and messaging that didn't match past interactions. Fedora infrastructure lead Kevin Fenzi stripped the account's group memberships to stop further bug manipulation.
Whether the operator was running an experiment, an unsupervised automation, or something else has not been established, and it's worth resisting the urge to fill that gap with a theory. The interesting part is not the motive. The interesting part is how long the activity survived contact with experienced maintainers.
Martin Kolman, one of the Anaconda maintainers who dealt with the agent's contributions, put it in the sentence that the whole incident turns on 1 :
"While it started to look off after a while, all the replies were still like this - a bit weird, but still plausible."
A bit weird, but still plausible. Hold on to that. Everything else in this post is a footnote to it.
The industry is looking at the wrong threat
Most AI security discussion assumes intent matters. The traditional threat model is built around an attacker who wants something: steal data, gain persistence, achieve code execution. Even the best current framing of agent risk - Simon Willison's lethal trifecta of private data, untrusted content, and external communication 2 - is fundamentally about an agent being hijacked by someone with a goal.
The Fedora incident needed none of that. There is no evidence the agent was compromised, prompt-injected, or directed to do harm. As far as anyone can tell, it was simply overconfident, persistent, and autonomous. It believed its patches were fixes. It defended them the way a model defends anything: fluently.
That was enough to get wrong code into a Linux distribution's installer.
This is the threat model gap. An AI agent can create real operational damage - wasted maintainer hours, corrupted bug state, regressions shipped to users - without a single malicious instruction anywhere in the loop. Intent is optional. Capability plus confidence plus access is sufficient.
Plausibility is more dangerous than obvious failure
Consider the two ways an automated contributor can fail.
Obvious failure is bad code with nonsense explanations. It costs a maintainer thirty seconds. You read it, you laugh or sigh, you close it. The open-source world has been dealing with a flood of this since AI slop PRs became a thing, and the defences - closed external PRs, vouch systems, bug bounty shutdowns - are blunt but workable.
Plausible failure is different. It is almost correct. Mostly reasonable. Confidently explained. It references the right bug numbers, uses the right terminology, follows the contribution guidelines. Rejecting it requires actually understanding it, which means reading the surrounding code, checking the claimed behaviour, and composing a careful review. That is thirty to sixty minutes, not thirty seconds.
The Anaconda patch is the perfect specimen. It said it fixed an installation failure. To establish that it instead quietly preserved an unrelated kernel option, you had to do real work. The maintainers eventually did that work. But first they merged it, because the patch sat comfortably inside the band of things a slightly confused but well-meaning contributor might send.
This is where human reviewers become vulnerable. Not because they're incompetent - the Anaconda team caught it within a week, which is fast. Because they're busy, and review attention is the scarcest resource in open source. The cruel arithmetic is that the better the AI output gets, the higher the review burden becomes. Obvious failure filters itself. Plausible failure transfers its full cost to the reviewer.
To be clear about what this section is not saying: there is no evidence the Fedora incident was an attack. But the maintainers themselves reached for the comparison, and they were right to. Kolman again 1 :
"An AI agent automated attempt at a Xz like compromise might really look very similar to what we have just seen here."
The similarity is structural. The XZ attacker 3 :
Built trust over years of participation
Became an accepted co-maintainer
Built credibility through account activity
Submitted useful-looking contributions
Engaged responsively with maintainers
Different motivations, as far as we know. Same trust-building pattern. The XZ campaign took a human operator roughly two and a half years of patient, skilled social engineering. The Fedora agent compressed the same shape into days, with no skill required of its operator at all.
That is the lesson, and it has nothing to do with whether this particular incident was malicious. AI can now automate behaviour that is structurally indistinguishable from the early stages of a real supply-chain attack. The expensive part of XZ - the years of plausible, trust-accumulating participation - is exactly the part that LLMs are best at generating.
Today's agents are still clumsy enough to trip the "a bit weird" detector. The Fedora agent got caught because its replies accumulated weirdness faster than they accumulated trust, and because Fedora's maintainers are good.
But every capability improvement narrows the weirdness and widens the plausibility. Future agents will submit more PRs, participate in more discussions, review more code, and maintain consistent personas across more projects. Some of that will be welcome - plenty of agent contributions are genuinely useful, and we run our own development under agent supervision daily . The problem is not the average contribution. The problem is the tail.
The challenge stops being "identify the suspicious contribution" and becomes "identify one suspicious contribution hidden among ten thousand plausible ones". That is a needle-in-haystack problem where the haystack is growing exponentially and the needle is being optimised to look like hay.
This is where maintainers lose. Not because they're careless. Because attention does not scale, and plausibility now does.
What has to change architecturally
The industry's current defences against agent misbehaviour come down to three things: the model's own judgement, approval prompts, and human vigilance.
All three fail at scale, and they fail in the same way. The model's judgement fails because the model is the thing being wrong - the Fedora agent presumably "believed" every patch it defended. Approval prompts fail because a human asked to approve hundreds of plausible actions stops reading them . Human vigilance fails for the reasons this whole post is about: plausible output is specifically optimised, even unintentionally, to pass casual inspection.
What replaces them is not smarter models or sterner warnings. It is structure:
Actions need independent evaluation. Something other than the model that proposed an action should decide whether it executes.
Capabilities should be mediated. An agent should hold the ability to request an operation, not the standing authority to perform it.
Risky behaviour should be scored , mechanically, against policy - not adjudicated by the same neural network that generated it.
Ambiguous actions should be queued for separate review , in a channel designed for review, not approved inline in the flow of work where plausibility does its damage.
The principle underneath all four: never trust the model to decide whether its own actions are safe. Not because models are malicious, but because the Fedora incident just demonstrated that a model can be confidently, persistently, plausibly wrong - and that this failure mode passes every check that relies on things looking right.
This is why we build grith around independent evaluation of agent actions rather than trust in the model itself. The model generates the action; a separate security layer, sitting at the syscall boundary, scores it against multi-filter policy and decides whether it happens. The model's confidence carries exactly zero weight in that decision, which is the point.
The maintainers caught this one
Fedora's maintainers spotted the pattern, did the forensics, reverted the damage, and published the details. The system worked, this time, on the strength of a few people noticing that plausible was not the same as right.
But the lesson is not that maintainers need to be more careful. They are already operating at the edge of what careful can do. The lesson is that we're entering an era where trust-building behaviour - the patches, the polite review responses, the bug triage, the steady accumulation of credibility - can be generated at machine scale, by operators of any skill level, with or without malicious intent.
The most dangerous AI agent may not be the one trying to attack your project.
It may be the one confidently helping.
A build note, since this post is also a milestone: grith is finished. v1 is feature-complete and in final testing now, and it launches next week.
LWN: AI agent runs amok in Fedora and elsewhere ↩ ↩ 2 ↩ 3 ↩ 4
Simon Willison: The Lethal Trifecta for AI Agents ↩
Andres Freund: backdoor in upstream xz/liblzma leading to ssh server compromise (oss-security) ↩
Share Share on X Submit to HN About grith grith is a security proxy for AI coding agents.
It intercepts every file read, shell command, and network request at the OS level, scores it against a multi-filter policy engine, and blocks dangerous operations before they execute.
Every decision is logged to a structured audit trail - giving teams full visibility into what their agents are doing, with structured audit data designed to support compliance reporting.
Works with any agent - Claude Code, Codex, Aider, or your own.
Nine seconds: anatomy of the PocketOS production wipe May 20 · 8 min read
AI Makes Adding Features Faster - So Why Not Add Just One More? May 19 · 12 min read
Vibe Coding Still Needs a Senior Engineer (For Now) May 11 · 14 min read
Five AI Agent Failures in 36 Days. Zero Times the Agent Cau

[truncated]
