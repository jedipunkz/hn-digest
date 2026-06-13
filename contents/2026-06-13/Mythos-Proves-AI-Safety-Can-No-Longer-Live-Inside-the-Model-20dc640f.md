---
source: "https://grith.ai/blog/mythos-ai-safety-cannot-live-inside-the-model"
hn_url: "https://news.ycombinator.com/item?id=48518185"
title: "Mythos Proves AI Safety Can No Longer Live Inside the Model"
article_title: "Mythos Proves AI Safety Can No Longer Live Inside the Model | grith"
author: "edf13"
captured_at: "2026-06-13T15:37:10Z"
capture_tool: "hn-digest"
hn_id: 48518185
score: 2
comments: 0
posted_at: "2026-06-13T15:23:25Z"
tags:
  - hacker-news
  - translated
---

# Mythos Proves AI Safety Can No Longer Live Inside the Model

- HN: [48518185](https://news.ycombinator.com/item?id=48518185)
- Source: [grith.ai](https://grith.ai/blog/mythos-ai-safety-cannot-live-inside-the-model)
- Score: 2
- Comments: 0
- Posted: 2026-06-13T15:23:25Z

## Translation

タイトル: 神話は、AI の安全性がもはやモデル内に存在できないことを証明しています
記事のタイトル: AI の安全性がモデル内に存在しないことを証明する神話 |グリス
説明: Anthropic は、最も有能なサイバー モデルを精査されたパートナーに限定し、危険なリクエストをそこから遠ざけ、何千時間にもわたってレッドチーム化しました。とにかく脱獄が表面化し、政府はそのモデルを完全に撤回した。そのストーリーのすべての安全制御はモデルの外側にありました。それがすべてです
[切り捨てられた]

記事本文:
神話は、AI の安全性がもはやモデル内に存在できないことを証明しています | grith grith .ai プラットフォーム ブログ 連絡先 grith / ブログ / 神話は、AI の安全性がモデル内に存在しないことを証明します 神話は、AI の安全性がモデル内に存在しなくなることを証明します
共有 X で共有 HN に送信 grith がまもなくリリースされます AI コーディング エージェントのセキュリティ プロキシが OS レベルで適用されます。興味を登録して、ライブ開始時に通知を受け取ります。
3 年間、AI の安全性は主に 1 つのことを意味してきました。それは、モデルをより安全にすることです。拒否するように訓練してください。エッジを微調整します。憲法上の規定を追加します。より良い評価を構築します。
理由は簡単でした。モデルが安全に動作する場合、システムは安全です。
今週、その想定が公の場で崩れた。
6月12日、米国政府はAnthropicに対し、同社の最も高性能な2つのモデル、Fable 5とMythos 5への全世界の外国人のアクセスを停止するよう命令した。 1 リアルタイムで外国人を他の国民から確実に選別できるプロバイダーは存在しないため、実際的な結果として、地球上のすべてのユーザーに対して両方のモデルが強制的に遮断されることになりました。この指令は国家安全保障当局を引用し、モデルがジェイルブレイクされたという主張に従ったものだった。 2
政治的な要素や見出しを取り除くと、より耐久性のあるものが残ります。 Mythos の物語全体、つまりモデルがどのようにリリースされ、どのように保護され、最終的にどのようにプルされたかということは、有能な AI のセキュリティ境界がすでにモデルの外に移動していることを示しています。業界も実際にはその点を認めている。それを大声で言っていないだけです。
Anthropic が実際に出荷したもの
Anthropic 自身の説明によると、Mythos 5 は「現在入手可能なモデルの中で最も強力なサイバーセキュリティ機能」を備えたモデルです。 3 あらゆる主要なオペレーティング システムとあらゆる主要な Web ブラウザの脆弱性を特定して悪用できる

に指示された場合は、ser。 4 それはマーケティングの成功ではありません。それが、このモデルが広くリリースされなかった理由です。
Anthropic が危険だと考えたモデルをどのように扱ったかをよく見てください。目立った点は 3 つありますが、そのどれもが「拒否するように訓練した」わけではありません。
まず、アクセスがゲートされていました。このフルパワーモデルは、制御されたプログラムである Project Glasswing にのみ適用されました。4 月の立ち上げ時には精査された約 50 の組織が、6 月までに約 150 に拡大し、Amazon、Apple、Google、Microsoft、CrowdStrike などの名前が挙げられ、すべてが防御作業にこのモデルを使用しています。 5 ここでの安全機構は、誰がモデルを保持できるかを示すリストです。それが環境制御です。それは完全に重みの外で生きています。
次に、リクエストがルーティングされました。パブリック モデルである Fable 5 には、別のシステムによって制限が適用された「Mythos クラス」の機能が搭載されています。サイバーセキュリティ、生物学、化学、およびモデル蒸留リクエストは、より機能の低い Claude Opus 4.8 に静かにリダイレクトされます。 6 もう一度読んでください。 Anthropic 独自のパブリック モデルのヘッドライン セーフティ機能は、モデルの前に配置され、モデルがどのリクエストの試行を許可されるかを決定するルーターです。何が安全かという判断は、判断されるものの外側で行われます。
第三に、これら 2 つの層が不十分であると判断された場合、法律によりモデルが市場から削除されました。エクスポート コントロールは、境界が可能な限りモデルの外側にあります。
3 つの安全機構、3 層、すべて外部にあります。モデルを内部から安全にするはずだった 1 つのこと、つまりその訓練された拒否が、まさに失敗した部分です。
エピソード全体のきっかけとなったテクニックは珍しいものではありませんでした。報告書によると、ある企業はモデルに「特定のコードベースを読み取り、ソフトウェアの欠陥を特定する」よう促したという。 2 当たり前のようなお願い

ary コード レビューは、脆弱性発見エンジンとして、訓練されたガードレールをまっすぐに通過し、反対側に出ました。
Anthropic はその深刻さについて異議を唱えており、ジェイルブレイクは範囲が狭く、普遍的ではないとし、口頭での証拠しか確認していないとし、同じ機能が GPT-5.5 を含む他の公開モデルですでに利用可能であると指摘しています。 [7] この特定のモデルが採用されるに値するかどうかという狭い問題に関しては、Anthropic の意見はおそらく正しいでしょう。
しかし、その議論はより大きな議論を認めます。最先端の研究室が、サイバー能力を抑制するという明確な目標を持ってモデルをレッドチーム化し、モデルを厳選した 50 の組織に制限し、それでも拒否するように訓練された行動を引き出す平易な言葉のプロンプトを備えているのであれば、訓練された拒否はセキュリティ境界ではありません。それは好みです。強い優先事項であり、通常は尊重されますが、十分に有能なモデルであれば、その要求を良性の表現で表現する人は誰でも無視できるものです。
モデルの能力が高ければ高いほど、「通常は拒否する」と「危害を加えることができない」の差は大きくなります。そしてモデルについて話す必要があるのは一度だけです。
このパターンは以前にも見たことがあります
これは新しい教訓ではありません。これはシステム セキュリティにおける最も古い教訓であり、新しいクラスのシステムの予定どおりに実現されます。
初期のオペレーティング システムはアプリケーションを信頼していました。プログラムがマシンに何かをするように要求し、マシンがそれを実行しました。最新のオペレーティング システムでは、アプリケーションが実際の損害を与えるほど強力になると、アプリケーションを信頼することができなくなるため、アプリケーションをプロセス境界、アクセス許可、システムコール メディエーションの背後に隔離します。
初期のブラウザの信頼された Web サイト。 Web ページが敵対的になるのに十分な機能を備えているため、最新のブラウザーはすべてのタブをサンドボックス化します。
初期のクラウド プラットフォームの信頼されたワークロード。現代のものw

コンテナ、VM、IAM ポリシー、およびポリシー エンジン内の ap ワークロード。完全に制御できないワークロードは、信頼するのではなく封じ込める必要があるワークロードであるためです。
パターンは毎回同じです。機能は控えめですが、信頼は安く、仕事をするものに信頼を置きます。機能がしきい値を超えると、信頼はアクターから出て、その周囲のアーキテクチャに移ります。誰もこれを好みの問題として決めたわけではありません。それは攻撃者との接触を生き延びたものです。
AI エージェントは同じ道を歩いており、Mythos は閾値を超えたことを示すマーカーです。
モデルの安全性から実行の安全性へ
AI システムに関して尋ねることができる明確な質問が 2 つあり、業界は最初の質問にほとんどのエネルギーを費やしてきました。
1 つ目は生成に関するもので、モデルは有害な出力を生成する可能性がありますか?これは、調整、拒否訓練、憲法上の規則、レッドチーム化の領域です。それは実際の仕事であり、今でも重要です。
2 つ目はアクションに関するものです。モデルは実際に何を行うことが許可されていますか?このファイルを読み取れるでしょうか。ネットワークに到達できますか?このコマンドを実行できますか。データをマシンから移動できますか?与えられた範囲を超えて行動できるか。これを実行の安全性と呼びます。
決定的な違いは、実行の安全性はモデルが信頼できるかどうかに依存しないことです。モデルには機能があり、おそらく間違っていると仮定し、その機能が影響を与える可能性のあるものを制限します。エージェントが読み取りを許可されていないファイルは、モデルがどのようにそれを望むように説得されたかに関係なく、読み取られません。モデルの正当性がどれほど合理的であるかに関係なく、ポリシー外のコマンドは実行されません。
また、実行の安全性は、モデルの安全性には決してありえない方法で、モデルに依存しません。拒否トレーニングは 1 つのモデルの重みに固有です - モデルを再トレーニングするか、別のモデルに交換し、安全性を確保します

プロパティがリセットされました。アクションがシステムと出会う点にある境界では、アクションがクロード、GPT、ジェミニ、ディープシーク、ラップトップ上で実行されているオープンウェイト モデル、またはまだ誰も出荷していないものからのものであるかどうかは関係ありません。評価するのはアクションであって、それを生み出したものの意図ではありません。
オープンウェイトが唯一の耐久性のある答えとなる理由
アクセス ゲート (Project Glasswing アプローチ) は実際の境界であり、これほど危険なモデルにとっては賢明な境界です。しかし、それには有効期限があり、その分野の誰もがそれを知っています。
有能なモデルが不足しているわけではありません。オープンウェイトモデルは容赦なく改良されています。重りが漏れる。技術は数週間以内に論文やリポジトリを通じて広まります。最も制限されたフロンティア モデルと、独自のハードウェア上で監視なしで実行できる最良のモデルとの間の機能差は縮まり続けています。 Anthropic は、ジェイルブレイク機能がすでに公開されているモデルに組み込まれていると主張し、自らその点を主張しました。
Mythos クラスの機能が誰でもダウンロードできるようになると、「誰がこのモデルを保持できるのか?」という疑問が生じます。有用な答えが得られなくなります。載せるべきリストはありません。残された唯一の疑問は、実行の安全性に関するものです。強力で信頼できないモデルがこの環境で実行されているとすると、実際にここで何ができるのでしょうか?そして、その答えはモデルの出自とは何の関係もなく、すべてはモデルの周りに配置したレイヤーと関係があります。
アクセス制御はディストリビューション層の境界です。これは、モデルが配布を終了するまで機能します。実行の安全性はアクション層の境界であり、アクション層はモデルの意図が実際の効果となる場所であるため、モデルが迂回できない唯一の場所です。
境界線はもう動いた
エピソード全体を元に戻すと、結論は避けられません。最も

業界で安全性を重視する研究所は、その最も危険なモデルを採用し、承認されたユーザーのリスト、リクエストルーター、そして最終的には連邦法によって保護しました。これらのメカニズムはいずれもモデル内にありません。モデル内の制御 (拒否トレーニング) は、コードベースの読み取りに関する文によってバイパスされたものです。
神話の永続的な重要性は、モデルができることではありません。それは、モデルへの反応が明らかにしたことです。有能なシステムについては、モデルを信頼するのではなく、すでに環境を保護しており、輸出規制、許可リスト、フォールバックルーターなどの鈍い手段を使用してそれを行っていることがわかりました。なぜなら、正確な手段がまだ重要な場所に組み込まれていないからです。
正確な手段は、システム セキュリティのあらゆる前世代でおなじみのものであり、新しいターゲットに向けられています。サンドボックス化。能力調停。エージェントは、操作を実行するための永続的な権限を保持するのではなく、操作を要求できます。ルールに対するアクションをスコアリングするポリシー エンジン。監査証跡。あいまいなケースについては人によるレビュー。それらのどれも、モデルが安全であるとは想定していません。それらはすべて、それが強力であると仮定しており、それがどれほど強力になっても効果が持続するものです。
ビルドするとこんな感じになる
これは境界グリスが構築される境界であり、この投稿の議論全体がどこでセキュリティ上の決定が行われるかによって生きるか死ぬかが決まるため、そのメカニズムについて具体的に説明する価値はあります。
Grith はモデル内に存在せず、モデルを信頼しません。これはその下、オペレーティング システムのシステムコール境界に位置し、エージェントが実際に実行するすべてのアクション (読み取られたすべてのファイル、すべてのネットワーク接続、エージェントが生成しようとするすべてのプロセス) をインターセプトします。これらの各アクションは、カーネルが許可される前に、ポリシーに対してスコアを付けるマルチフィルター セキュリティ プロキシによって評価されます。

それを実行する。このモデルは次のように提案しています。代理人が処分します。決定はモデル以外のものによって行われるため、モデルの信頼性、そのトレーニング、その調整、およびその来歴は、その決定においてまったく重みを持ちません。
それがミトスの物語が緊急にしている部分だ。 「このコードベースの欠陥を見つける」ことを青信号として扱うように話されているジェイルブレイクされたモデルでも、ファイルを開いたり、ネットワークにアクセスしたり、エクスプロイトをディスクに書き込んだりするなど、その目的で何かを行うには実際のシステムコールを発行する必要があります。これらはまさに、上流のモデルがどれほど徹底的に説得されたかに関係なく、grith が評価し、否定できるアクションです。読み取りが許可されていないファイルは読み取られません。到達が許可されていないエンドポイントには到達しません。モデルが間違っていても、ジェイルブレイクしていても、あるいは単にガードレールよりも高性能であっても、そのアクションが触れることを許可されているものについては何も変わりません。
そして、その境界は特定の重みセット内ではなくシステムコール層に存在するため、その構造はモデルに依存しません。 grith は、 Grith run で独自の組み込みエージェントを監視している場合でも、 Grith exec で外部ツール (Claude Code、Codex、Aider) をラップしている場合でも、同じように機能します。クロードを交換する

[切り捨てられた]

## Original Extract

Anthropic restricted its most capable cyber model to vetted partners, routed risky requests away from it, and red-teamed it for thousands of hours. A jailbreak surfaced anyway, and the government pulled the model entirely. Every safety control in that story lived outside the model. That is the whole
[truncated]

Mythos Proves AI Safety Can No Longer Live Inside the Model | grith grith .ai Platform Blog Contact grith / Blog / Mythos Proves AI Safety Can No Longer Live Inside the Model Mythos Proves AI Safety Can No Longer Live Inside the Model
Share Share on X Submit to HN grith is launching soon A security proxy for AI coding agents, enforced at the OS level. Register your interest to be notified when we go live.
For three years, AI safety has mostly meant one thing: make the model safer. Train it to refuse. Fine-tune the edges. Add constitutional rules. Build better evaluations.
The reasoning was simple. If the model behaves safely, the system is safe.
This week, that assumption broke in public.
On June 12, the US government ordered Anthropic to suspend access to its two most capable models, Fable 5 and Mythos 5, for any foreign national worldwide. 1 Because no provider can reliably sort foreign nationals from everyone else in real time, the practical result was a hard shutoff of both models for every user on the planet. The directive cited national security authorities and followed a claim that the model had been jailbroken. 2
Strip away the politics and the headline and you are left with something more durable. The entire Mythos saga - how the model was released, how it was guarded, and how it was ultimately pulled - is a demonstration that the security boundary for capable AI has already moved outside the model. The industry has conceded the point in practice. It just has not said so out loud.
What Anthropic actually shipped
Mythos 5 is, by Anthropic's own description, the model with "the strongest cybersecurity capabilities of any model currently available." 3 It can identify and exploit vulnerabilities in every major operating system and every major web browser when directed to. 4 That is not a marketing flourish. It is the reason the model was never broadly released.
Look closely at how Anthropic handled a model it considered that dangerous. Three things stand out, and none of them is "we trained it to refuse."
First, access was gated . The full-power model went only to a controlled program, Project Glasswing - roughly 50 vetted organisations at launch in April, expanded to around 150 by June, names like Amazon, Apple, Google, Microsoft and CrowdStrike, all using it for defensive work. 5 The safety mechanism here is a list of who is allowed to hold the model at all. That is an environmental control. It lives entirely outside the weights.
Second, requests were routed . The public model, Fable 5, ships "Mythos-class" capability with restrictions applied by a separate system: cybersecurity, biology, chemistry and model-distillation requests get quietly redirected to the less capable Claude Opus 4.8. 6 Read that again. Anthropic's own headline safety feature for its public model is a router that sits in front of the model and decides which requests the model is even permitted to attempt. The judgment about what is safe is made outside the thing being judged.
Third, when those two layers were judged insufficient, the law removed the model from the market . Export controls are about as far outside the model as a boundary can get.
Three safety mechanisms, three layers, all of them external. The one thing that was supposed to make the model safe from the inside - its trained refusals - is precisely the part that failed.
The technique that triggered the whole episode was not exotic. According to the reporting, a company prompted the model to "read a specific codebase and identify software flaws." 2 A request that sounds like ordinary code review walked straight past the trained guardrails and out the other side as a vulnerability-discovery engine.
Anthropic disputes the severity - it calls the jailbreak narrow and non-universal, says it has seen only verbal evidence, and points out the same capability is already available in other public models including GPT-5.5. 7 On the narrow question of whether this particular model deserved to be pulled, Anthropic may well be right.
But that argument concedes the larger one. If a frontier lab can spend thousands of hours red-teaming a model with the explicit goal of suppressing its cyber capabilities, restrict it to fifty hand-picked organisations, and still have a plain-language prompt elicit the behaviour it was trained to refuse - then trained refusal is not a security boundary. It is a preference. A strong preference, usually honoured, but one that a sufficiently capable model can be talked out of by anyone who phrases the request as something benign.
The more capable the model, the larger the gap between "usually refuses" and "cannot do harm." And the model only has to be talked out of it once.
We have seen this pattern before
This is not a new lesson. It is the oldest lesson in systems security, arriving on schedule for a new class of system.
Early operating systems trusted their applications. A program asked the machine to do something and the machine did it. Modern operating systems isolate applications behind process boundaries, permissions and syscall mediation, because trusting the application stopped being viable once applications got powerful enough to do real damage.
Early browsers trusted websites. Modern browsers sandbox every tab, because a web page became capable enough to be hostile.
Early cloud platforms trusted workloads. Modern ones wrap workloads in containers, VMs, IAM policies and policy engines, because a workload you do not fully control is a workload you have to contain rather than trust.
The pattern is identical every time. While capabilities are modest, trust is cheap and you put it in the thing doing the work. Once capabilities cross a threshold, trust moves out of the actor and into the architecture around it. Nobody decided this as a matter of taste. It is what survived contact with attackers.
AI agents are walking the same road, and Mythos is the marker that says the threshold has been crossed.
From model safety to execution safety
There are two distinct questions you can ask about an AI system, and the industry has spent most of its energy on the first.
The first is about generation: can the model produce harmful output? This is the domain of alignment, refusal training, constitutional rules and red-teaming. It is real work and it still matters.
The second is about action: what is the model permitted to actually do? Can it read this file. Can it reach the network. Can it run this command. Can it move data off the machine. Can it act outside the scope it was given. Call this execution safety.
The crucial difference is that execution safety does not depend on the model being trustworthy. It assumes the model is capable and possibly wrong, and it constrains what that capability can touch. A file the agent is not permitted to read does not get read, regardless of how the model was persuaded to want it. A command outside policy does not run, regardless of how reasonable the model's justification sounded.
And execution safety is model-agnostic in a way model safety can never be. Refusal training is specific to one model's weights - retrain the model, or swap in a different one, and the safety properties reset. A boundary that sits at the point where actions meet the system does not care whether the action came from Claude, GPT, Gemini, DeepSeek, an open-weights model running on a laptop, or something nobody has shipped yet. It evaluates the action, not the intentions of whatever produced it.
Why open weights make this the only durable answer
Access gating - the Project Glasswing approach - is a real boundary, and a sensible one for a model this dangerous. But it has an expiry date, and everyone in the field knows it.
Capable models do not stay scarce. Open-weights models improve relentlessly. Weights leak. Techniques diffuse through papers and repositories within weeks. The capability gap between the most restricted frontier model and the best model you can run unsupervised on your own hardware keeps narrowing. Anthropic made the point itself when it argued the jailbroken capability is already in publicly available models.
Once a Mythos-class capability is something anyone can download, the question "who is allowed to hold this model?" stops having a useful answer. There is no list to be on. The only question left is the execution-safety one: given that a powerful, untrusted model is running in this environment, what is it actually able to do here? And that answer has nothing to do with the model's provenance and everything to do with the layers you put around it.
Access control is a boundary at the distribution layer. It works right up until the model escapes distribution. Execution safety is a boundary at the action layer, and the action layer is the one place the model cannot route around, because it is where its intentions become real effects.
The boundary has already moved
Put the whole episode back together and the conclusion is hard to avoid. The most safety-conscious lab in the industry took its most dangerous model and protected it with a list of approved users, a request router, and ultimately federal law. Not one of those mechanisms is inside the model. The inside-the-model control - refusal training - is the one that got bypassed by a sentence about reading a codebase.
The lasting significance of Mythos is not what the model can do. It is what the response to the model revealed: that for capable systems, we are already securing the environment rather than trusting the model, and we are doing it with blunt instruments - export controls, allowlists, fallback routers - because the precise ones have not been built into the places that matter yet.
The precise instruments are the familiar ones from every prior generation of systems security, pointed at a new target. Sandboxing. Capability mediation, so an agent can request an operation rather than hold standing authority to perform it. Policy engines that score actions against rules. Audit trails. Human review for the ambiguous cases. None of them assume the model is safe. All of them assume it is powerful, and stay effective no matter how powerful it becomes.
What this looks like when you build it
This is the boundary grith is built on, and it is worth being concrete about the mechanism, because the whole argument of this post lives or dies on where the security decision is made.
grith does not sit inside the model and it does not trust the model. It sits underneath it, at the operating-system syscall boundary, and intercepts every action an agent actually takes - every file read, every network connection, every process it tries to spawn. Each of those actions is evaluated by a multi-filter security proxy that scores it against policy before the kernel is allowed to carry it out. The model proposes; the proxy disposes. The model's confidence, its training, its alignment and its provenance carry exactly zero weight in that decision, because the decision is made by something other than the model.
That is the part the Mythos story makes urgent. A jailbroken model that has been talked into treating "find the flaws in this codebase" as a green light still has to issue real syscalls to do anything with that intent - open the file, reach the network, write the exploit to disk. Those are exactly the actions grith evaluates and can deny, regardless of how thoroughly the model upstream was persuaded. The file it is not allowed to read does not get read. The endpoint it is not allowed to reach does not get reached. The model being wrong, or jailbroken, or simply more capable than its guardrails, changes nothing about what its actions are permitted to touch.
And because that boundary lives at the syscall layer rather than inside any particular set of weights, it is model-agnostic by construction. grith works the same whether it is supervising its own built-in agent with grith run or wrapping an external tool - Claude Code, Codex, Aider - with grith exec . Swap Claude

[truncated]
