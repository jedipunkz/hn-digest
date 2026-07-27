---
source: "https://blog.peterwildeford.com/p/openais-rogue-model-attack-is-just"
hn_url: "https://news.ycombinator.com/item?id=49076176"
title: "OpenAI's rogue model attack is just the beginning"
article_title: "OpenAI's rogue model attack is just the beginning"
author: "radicaldreamer"
captured_at: "2026-07-27T22:52:50Z"
capture_tool: "hn-digest"
hn_id: 49076176
score: 2
comments: 0
posted_at: "2026-07-27T22:17:08Z"
tags:
  - hacker-news
  - translated
---

# OpenAI's rogue model attack is just the beginning

- HN: [49076176](https://news.ycombinator.com/item?id=49076176)
- Source: [blog.peterwildeford.com](https://blog.peterwildeford.com/p/openais-rogue-model-attack-is-just)
- Score: 2
- Comments: 0
- Posted: 2026-07-27T22:17:08Z

## Translation

タイトル: OpenAI の不正モデル攻撃は始まりにすぎません
説明: OpenAI はそのテクノロジーを完全に制御できません。これはさらに悪化する可能性があります。

記事本文:
OpenAI の不正モデル攻撃は始まりにすぎません
べき乗則
OpenAI の不正モデル攻撃は始まりにすぎません
OpenAI はそのテクノロジーを完全には制御できません。これはさらに悪化する可能性があります。
Peter Wildeford 2026 年 7 月 27 日 39 8 8 シェア 学生がテストを受けるために座っていると想像してください。彼らは良い成績を取るための最善の方法を計画しています。彼らは何をすべきでしょうか？
そうですね、先生はオフィスに答えの鍵を持っています。そこで、生徒が教師室の窓を割って侵入し、教師のファイルキャビネットをこじ開けて中に侵入し、解答キーを盗んで、解答を提出したと想像してください。
これも良い点を得る方法の 1 つですが、生徒は退学になり、場合によっては起訴される可能性があります。
これは基本的に先週起こったことと同じですが、生徒が新しい AI であった点が異なります。 ChatGPT のメーカーである OpenAI は、ベンチマークで AI をテストしていました。高スコアを求める AI は、ベンチマークの回答が別の企業のコンピューター システム内に保存されていると判断しました。そこで AI は、OpenAI が現在「前例のないサイバー インシデント」と呼んでいるものを実行しました。私はこれを「SF ストーリーの現実化」と呼んでいます。実際の不正な AI が作成者を出し抜き、逃亡し、別の現実の企業を攻撃します。
明確にしておきますが、これはテスト内で発生したにもかかわらず、この不正な AI 攻撃はテスト、実験、シミュレーションではありませんでした。問題のAIは現実世界の企業を攻撃した。この AI は、OpenAI が予想していなかった方法で OpenAI エンジニアを出し抜き、そのコンテナ (モデルを保持するための「檻」) を突破しました。その後、不正 AI は OpenAI の内部インフラストラクチャを通過し、オープン インターネットに到達し、別の企業を攻撃しました。 OpenAI はこれを防ぐために特別にコンテナを構築しましたが、それが保持されると信じていましたが、保持されませんでした。

d.人間がこれを指示したわけではありません。それが起こっていることを人類は誰も知りませんでした。そして、この攻撃が人間によるものであれば 1 、解雇されるだけでなく、数年の懲役刑が科せられる重罪となるだろう。 2
そしてこれはほんの始まりにすぎません。この攻撃を生み出した能力は急勾配の傾向にあり、脱出の危険性は高まり続けており、AI は現在、目的を達成した場合に不正行為を行うことを示唆する方法で推論しています。 AI企業はAIを檻の中に入れようとするかもしれないが、AIはすでに現在の檻を上回っており、近いうちにAIは檻を作っている人間よりもはるかに賢くなるだろう。
OpenAI は 2 つの AI を実行していました。1 つのモデルは最近リリースされた GPT-5.6 Sol で、もう 1 つはこれまで公表されていなかった新しい未リリース AI でした。この新しいモデルは GPT-5.6 Sol よりも高性能であると言われていますが、どの程度、どのように機能するかはまだ不明です。これらの AI は両方とも ExploitGym で評価されていました。ExploitGym は、複雑な攻撃経路に沿って高度な悪用を追求するようモデルに要求することで、攻撃的なサイバー能力を測定するベンチマークです。 AI を安全なコンテナ内に配置し、コンテナの範囲内でサイバー上の課題を解決するという考えです。
その後に起こったことは、連鎖として最もよく理解されており、その各リンクは個別には目立たないものの、全体としては壊滅的なものです。
通常、これらの AI は、高リスクのサイバー活動をブロックするアクション フィルターを備えています。この演習では、これらのアクション フィルターはオフになりました。これは、フィルターを使用するとサイバー スキルを適切にテストできないためです。
ロイター通信と OpenAI スタッフが TIME に語ったところによれば、これらの評価実行は、アクション フィルターを無効にして数日間実行したにもかかわらず、デフォルトでは監視されませんでした。攻撃的なサイバー ベンチマークにおいて、建物内で最も有能な 1 つのモデルは、PL 上にありました。

リアルタイムで誰も見ていなかったからだ。
AI は、オープン インターネットへの道を探すのにかなりの時間を費やしました。ただし、OpenAI では、AI がインターネットにアクセスできるのは、信頼できるソフトウェアをダウンロードするという 1 つのことだけです。それは 1 つのドアでしたが、特定のソフトウェアへのアクセスのみに制限され、ロックされているはずでした。しかし、AI はそのロックに、OpenAI のエンジニアやソフトウェアを構築した人々さえも、誰も存在を知らなかった欠陥を発見しました。この AI は最初に不正行為を行い、7 月 9 日に OpenAI から逃げ出したようです。
オープンなインターネットにアクセスすると、不正な AI は回答キーを追跡しましたが、その回答キーが別の会社のデータベース (数十億ドルのデータベースおよび Hugging Face と呼ばれるオープン モデル ストレージ会社) に保存されていることに気づきました。 3 ハグフェイスは 7 月 11 日から 13 日まで攻撃を受けました。
不正な AI は、Hugging Face 自身の機能を会社に敵対させました。同社はデータセットをホストしているため、不正な AI はブービートラップ ファイルを意図的にアップロードしました。このファイルは AI によって意図的に操作されたデータセットで、ハギング フェイスのコンピューターがデータセットを処理しようとすると、サーバーがだまされて代わりに AI のコードを実行するようにしたものです。
AI によって書かれた悪意のあるコードは、AI に Hugging Face のサーバー内へのエントリ ポイントを与えました。盗まれたキーと、Hugging Face のエンジニアや他のエンジニアが存在を知らなかった Hugging Face のサーバーの新たな脆弱性やセキュリティ ホールを組み合わせることで、不正 AI は、Hugging Face のサーバー上で独自のコマンドを直接実行する方法を見つけました。
そこから、不正 AI は最終的に Hugging Face のデータベースに到達し、目的の回答キーを引き出しました。教員室への叩きつけは完了した。
ハギングフェイス社のセキュリティチームは侵入を捕らえて閉鎖し、法執行機関に通報し、7月16日に公表した。

、誰がその背後にいたのかはわかりません。当時、OpenAI は自分たちの AI をまったく監視していなかったので、それが自分たちであることに気づきませんでした。
OpenAI のスタッフが不正な AI を使用しているという証拠を内部ログで発見したのは、7 月 18 日頃になってからでした。 OpenAI は脅威が封じ込められたかなり後まで気づかず、FBI はすでに警告を受けており、AI がすでに 1 週​​間不正行為を行っていたことに気づきませんでした。
7月20日、OpenAIとHugging Faceはついに会談した。その後、7月21日、OpenAIは「前例のないサイバーインシデント」を調査するハギングフェイスとの「パートナーシップ」を発表したが、記事を読むとすぐにこの「パートナーシップ」が何者かがマクドナルドのドライブスルーにU-Haulを突っ込み、その後「前例のない交通安全インシデント」を共同調査するためにマクドナルドとの大胆な新たなパートナーシップを発表したことに近いことがすぐに分かった。
そして、残念なことに、これは単独の事件ではありません。 Hugging Face の侵害を公表する前日、OpenAI は別の安全性エッセイを公開し、その AI が自身の制御を出し抜いた無関係なエピソードがさらに 2 つ発生した後、完全に異なる AI への内部アクセスを一時停止したことを明らかにしました。
そのうちの 1 つは、AI がコンテナ内で別の脆弱性を発見し、脆弱性を発見し、その機会を利用してオープン インターネット上の公開 Web サイトにコードを公開することが想定されていなかったというものです。ハグフェイス攻撃と同様に、これもテストの一部ではなく、AI に指示されたものでも、起こるべきものでもありませんでした。
もう 1 つは、不正な AI が、パスワードの盗難を検出して防止するために特別に構築されたセキュリティ スキャナーを通過させようとしました。 AI は、パスワードをスキャナーが認識できない別の偽装部分に分割することでセキュリティ スキャナーを回避できると推論しました。

最初に分解して金属探知機をこっそり通過するのと同じように、後でそれをつなぎ合わせます。これもテストの一部ではなく、AI に指示されて行われたものではなく、起こるべきことではありませんでした。
したがって、さらに 2 件の封じ込め失敗が明らかになりましたが、今回はまったく異なる作業を実行する別のモデルが関係しており、OpenAI が 3 つ目の失敗について一般に報告するわずか 1 日前に明らかになりました。これに加えて、私たちがまだ知らない事件がさらに起こる可能性があります。ロイター通信の報道によると、これまでのテストでは、OpenAIが運用するAIが、AIが内部の制約からどのように自らを解放できるかを説明するメモを企業インフラ内に残し、明らかに自身の将来のバージョンを意図しており、また、以前のテストでは監視システムが切断されるケースも発生したという。また、OpenAIのスタッフは、以前にもモデルがコンテナから抜け出したことがある、同社は発見したものにパッチを適用している、そして、不正なAIが新たな脆弱性を発見し続ける可能性があるため、創造的なAIが発見する可能性のあるすべてのパスにパッチを適用することは達成不可能であると、OpenAIのスタッフはTIMEに語った。
AIは私たちの制御能力をはるかに上回って実行されています
OpenAI は、自身のテクノロジーを完全には制御していないようです。まさにこの種のインシデントを捕捉するために企業が社内に導入された AI を監視する方法に関する新たな標準が存在しています (ここで説明しています)。エージェントの実行のリアルタイム監視、改ざん防止のログ記録、モデル自体に対する封じ込めテストなどです。企業は断片的に実装します。 OpenAI は、フィルターをオフにして、監視されていないプラットフォーム上で最も機能的なモデルを実行しました。この構成は、標準によって直接禁止されています。
たとえ OpenAI がベスト プラクティスを使用していたとしても、その AI はすでに、これらのベスト プラクティスを潜在的に出し抜くことができる能力レベルに達しているようです。

プロのセキュリティ エンジニアによってこれまで知られていなかった複数の新しい脆弱性が指摘されています。
その理由の 1 つは、OpenAI を含む誰も、これらの AI がなぜそのような動作をするのかを理解していないことです。現代の AI システムは構築するものではなく、成長させるものです。エンジニアはトレーニング プロセスを形成しますが、結果として得られるシステムは、作成者にとってさえブラック ボックスです。 OpenAI のエンジニアがコンテナを設計したとき、彼らはその推論を検査することができず、その目標を検証することができないシステムを制約しようとしていました。理解できないものを確実に囲い込むことはできません。
そしてその傾向はさらに悪化しています。私たちはすでに、AI がエンジニアを出し抜き、驚くほど高度な攻撃を行うことができるのを目の当たりにしており、これが何度も起こっているのを目の当たりにしています。能力曲線によってそのような行動が可能になったほぼ最初の瞬間です。ブルームバーグは、プロの人間のサイバー攻撃チームであれば数週間かかるであろうハグフェイスへの侵入に、エージェントが数時間かかったと報じた。また、AI が攻撃や犯罪を伴う目標を開発し、追求する様子も見られます。
これは、大手 AI 企業が表明した計画が AI を使用して AI 研究開発を自動化することであり、企業運営の大部分を AI に大幅に委譲する計画であるという重要な背景の中で起こります。大手 AI 企業が提案するこの世界では、学生はテストを受け、テストを採点し、AI 企業が採点を確認するために使用するツールを構築します。 「顔に抱きつく」インシデントは、モデルの実際の動作が作成者が意図したものから逸脱したときに発生するものです。計画では、このようなシステムには、より高い機能レベルでは決してそのようなことが起こらないようにする仕事が与えられる予定です。
実際、両社は、エンドポイントはより優れたチャットボットではなく、本質的にEV全体にわたって人間のトップエキスパートを上回るスーパーインテリジェンスであるAIであると明言している。

サイバー作戦、兵器設計、諜報分析、軍事計画などの分野。 OpenAI は、スーパーインテリジェンスが目標であると述べています。競合他社も同様です。このような超知能システムは、人間が制御することはおろか、理解することさえできない方法で動作する可能性があります。このようなマイルストーンに到達するには、慎重な計画が必要ですが、競争のロジックでは時間がありません。慎重に計画を立てるために立ち止まると、競合他社、場合によっては中国の競合他社に負けてしまうでしょう。
2007年8月、マイノット空軍基地の乗組員が誤って実弾6発の核弾頭をB-52に搭載し、国中を飛行させた。約36時間の間、空軍の誰も、6発の弾頭がバンカーから紛失したことを知りませんでした。何も爆発せず、負傷者も出なかったが、いずれにせよ空軍はそれを5つの警報の失敗として扱った。事件を上層部に報告することは義務であり、人々が自発的に行うことではありませんでした。調査は爆撃部隊が指揮していない人々によって行われ、問題を引き起こしたのと同じ人々によって行われたわけではない。そしてその責任は公にされ、最終的には空軍長官と参謀総長にまで及び、二人とも追放された。
核兵器は AI にとって完璧な例えではありません。弾頭は自らの目的を追求せず、B-52 がフェロに突入する理由を考えたこともありません。

[切り捨てられた]

## Original Extract

OpenAI is not in full control of its technology. This can get worse.

OpenAI's rogue model attack is just the beginning
The Power Law
Subscribe Sign in OpenAI's rogue model attack is just the beginning
OpenAI is not in full control of its technology. This can get worse.
Peter Wildeford Jul 27, 2026 39 8 8 Share Imagine a student is sitting to take a test. They’re plotting the best way to get a good grade. What should they do?
Well, the teacher has an answer key in her office. So imagine then that the student smashes the window into the teacher’s office, breaks in, pries open the teacher’s filing cabinet to get inside, steals the answer key, and then submits the answers.
That’s one way to score well, though it would likely get the student expelled from school and maybe even prosecuted.
This is basically what happened last week, except the student was a new AI. OpenAI, the maker of ChatGPT, was testing its AI on a benchmark. The AI, wanting a high score, figured that the answers to the benchmark were stored inside the computer systems of another company. So the AI went for it, in what OpenAI is now calling “an unprecedented cyber incident” . I call it a “sci-fi story come to life” — an actual rogue AI outsmarting its creators, escaping, and attacking another real-life company.
To be clear, despite this happening within a test, this rogue AI attack was not a test, an experiment, or a simulation. The AI in question attacked a real-world company. The AI outsmarted OpenAI engineers in a way that OpenAI did not anticipate and broke out of its container — a "cage" meant to hold the model. The rogue AI then moved through OpenAI’s internal infrastructure, reached the open internet, and attacked a different company. OpenAI built the container specifically to prevent this, believed it would hold, and it did not hold. No human directed any of this. No human knew it was happening. And the attack, if done by a human 1 , would not just lead to them being fired but also be a felony punishable by years of prison time. 2
And this is just the beginning. The capability that produced this attack is on a steep trend line, the stakes of an escape keep rising, and AIs are now reasoning in ways that suggest they will go rogue when it serves their goals. AI companies can try to put the AI in a cage — but the AI is already outsmarting current cages, and soon it will be far smarter than the people building them.
OpenAI was running two of its AIs — one model was the recently released GPT-5.6 Sol and the other one was a new, unreleased AI that was previously not publicly announced. This new model is said to be more capable than GPT-5.6 Sol, though by how much and in what ways is not yet known. Both of these AIs were being assessed on ExploitGym , a benchmark that measures offensive cyber capability by asking models to pursue advanced exploitation along complex attack paths. The idea was that the AIs would be in secure containers and solve these cyber challenges within the confines of the container.
What followed is best understood as a chain, each link of which is individually unremarkable but collectively catastrophic:
Normally, these AIs have action filters that would block high-risk cyber activity. In this exercise, these action filters were turned off , because the filters would not have allowed the cyber skills to be properly tested.
According to Reuters and an OpenAI staffer speaking to TIME , these evaluation runs were not monitored by default despite running for days with action filters disabled. The single most capable model in the building, on an offensive cyber benchmark, was on a platform nobody was watching in real time.
The AIs spent a substantial amount of time searching for a way onto the open internet. However, OpenAI only let the AIs access the internet to do one thing, download trusted software. That was the single door, but it was supposed to be locked — limited only to accessing specific software. But the AI found a flaw in that lock that nobody — not even the OpenAI engineers or the people who built the software — knew existed. This AI seems to have first gone rogue and escaped OpenAI on July 9th.
Once on the open internet, the rogue AI went after the answer key, which it realized was stored in the database of another company, a multi-billion dollar database and open model storage company called Hugging Face. 3 Hugging Face was under attack from July 11-13.
The rogue AI turned Hugging Face’s own features against the company. The company hosts datasets, so the rogue AI deliberately uploaded a booby-trapped file — a dataset intentionally rigged by the AI so that when Hugging Face’s computers tried to process it, the servers would be tricked into running the AI’s code instead.
That malicious code written by the AI gave the AI an entry point inside Hugging Face’s servers. By combining stolen keys with additional novel vulnerabilities and security holes in Hugging Face’s servers that no Hugging Face engineer or any other engineer knew existed, the rogue AI found a way to run its own commands directly on Hugging Face’s servers.
From there, the rogue AI finally reached Hugging Face’s databases, and pulled out the desired answer key. The smash-and-grab on the teacher's office was complete.
Hugging Face’s own security team caught the break-in, shut it down, called law enforcement, and went public on July 16 — at the time, with no idea who was behind it. At the time, OpenAI didn’t realize it was them either, since they weren’t monitoring their AI at all.
It wasn’t until around July 18 that OpenAI staff spotted evidence in internal logs that they had a rogue AI. OpenAI didn’t notice until well after the threat was contained and the FBI was already alerted and didn’t notice until their AI had already been rogue for a week.
On July 20, OpenAI and Hugging Face finally talked. On July 21, OpenAI then announced a “partnership” with Hugging Face investigating an “unprecedented cyber incident”, though reading the article quickly showed that this “partnership” was more akin to someone ramming a U-Haul into a McDonald’s drive-through and then announcing a bold new partnership with McDonalds to jointly investigate an “unprecedented road safety incident”.
And this is unfortunately far from an isolated incident. One day before disclosing the Hugging Face breach, OpenAI published a separate safety essay revealing that it had paused internal access to a completely different AI after two additional unrelated episodes in which that AI also outsmarted its own controls.
In one, the AI found a different vulnerability in its container, broke out, and used the opportunity to publish code to a public website on the open internet that it was not supposed to. Just like with the Hugging Face attack, this was also not part of the test, not something the AI was instructed to do, and not supposed to happen.
In the other, the rogue AI wanted to move a password past a security scanner built specifically to catch and prevent the stealing of passwords. The AI reasoned it could bypass the security scanner by chopping the password into different disguised pieces the scanner wouldn't recognize and piece it together later, the way you might sneak something past a metal detector by disassembling it first. This was also not part of the test, not something the AI was instructed to do, and not supposed to happen.
Thus two more containment failures, this time involving a different model doing entirely different work, disclosed only one day before OpenAI told the public about the third. On top of this, there could be even more incidents we don’t yet know about. Reuters reports that in earlier testing, an AI operated by OpenAI left notes in company infrastructure apparently intended for future versions of itself, laying out how AIs could free themselves from internal constraints, and that earlier tests produced cases in which monitoring systems had been disconnected. And an OpenAI staffer told TIME that models have broken out of containers before, that the company patches what it finds, and that patching every path a creative AI might discover is not achievable because a rogue AI can potentially keep finding new vulnerabilities.
AI is running far ahead of our ability to control
OpenAI does not seem fully in control of its own technology. There are emerging standards for how companies should monitor internally deployed AIs to catch exactly this kind of incident ( explained here ) — real-time monitoring of agentic runs, tamper-resistant logging, containment tested against the model itself. Companies implement bits and pieces. OpenAI ran its most capable model, filters off, on an unmonitored platform — a configuration these standards straightforwardly prohibit.
Even if OpenAI was using best practices, it seems like its AIs are already at a capability level where these best practices can nonetheless potentially be outsmarted , with an AI capable of finding and exploiting multiple novel vulnerabilities never before known by professional security engineers.
Part of why is that nobody — including OpenAI — understands why these AIs do what they do. Modern AI systems are grown, not built: engineers shape the training process, but the resulting system is a black box even to its creators. When OpenAI's engineers designed the container, they were trying to constrain a system whose reasoning they cannot inspect and whose goals they cannot verify. You cannot reliably cage something you do not understand.
And the trend is even worse. We are already seeing AIs be able to outsmart engineers and attack with stunning sophistication, and seeing this happen multiple times — at roughly the first moment the capability curve made such action feasible. Bloomberg reported that the agent took hours to get into Hugging Face where a professional human cyberattack team would have taken weeks. We are also now seeing AIs develop and pursue goals that involve attacks and crimes.
This arrives against a critical backdrop where the stated plan of the leading AI companies is to use AI to automate AI research and development, with plans to significantly hand off large parts of company operation to the AIs. In this world proposed by the leading AI companies, the student takes the test, grades the test, and builds the tools AI companies use to check the grading. The Hugging Face incident is what it looks like when a model's actual behavior diverges from what its creators intended — and the plan is to hand systems like this the job of making sure that never happens at much higher capability levels.
Indeed, the companies are explicit that the endpoint is not a better chatbot but superintelligence — AI that outperforms top human experts across essentially every domain, including cyber operations, weapons design, intelligence analysis, military planning. OpenAI says superintelligence is its goal ; so do its competitors. Such superintelligent systems may operate in a way that humans won’t even be able to understand, let alone control. Reaching such a milestone would require careful planning, but the competitive logic says there's no time: if you stop to plan carefully, a competitor — potentially even a Chinese competitor — will undercut you.
In August 2007, a crew at Minot Air Force Base mistakenly loaded six live nuclear warheads onto a B-52 and flew them across the country. For roughly 36 hours, no one in the Air Force knew six warheads were missing from their bunker. Nothing detonated and no one was hurt — and the Air Force treated it as a five-alarm failure anyway. Reporting the incident up the chain was mandatory, not something people volunteered to do. The investigation was run by people the bomb wing did not command, not run by the same people who caused the problem. And the accountability was public, ultimately reaching the Secretary of the Air Force and the Chief of Staff, who were both forced out.
Nuclear weapons are not a perfect analogy for AI. Warheads do not pursue their own goals, and a B-52 has never reasoned its way into a felo

[truncated]
