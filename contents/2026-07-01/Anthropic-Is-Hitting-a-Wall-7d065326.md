---
source: "https://www.vincentschmalbach.com/anthropic-is-hitting-a-wall/"
hn_url: "https://news.ycombinator.com/item?id=48742441"
title: "Anthropic Is Hitting a Wall"
article_title: "Anthropic Is Hitting a Wall - Vincent Schmalbach"
author: "vincent_s"
captured_at: "2026-07-01T05:31:00Z"
capture_tool: "hn-digest"
hn_id: 48742441
score: 3
comments: 0
posted_at: "2026-07-01T05:00:54Z"
tags:
  - hacker-news
  - translated
---

# Anthropic Is Hitting a Wall

- HN: [48742441](https://news.ycombinator.com/item?id=48742441)
- Source: [www.vincentschmalbach.com](https://www.vincentschmalbach.com/anthropic-is-hitting-a-wall/)
- Score: 3
- Comments: 0
- Posted: 2026-07-01T05:00:54Z

## Translation

タイトル: 人類は壁にぶつかっている
記事のタイトル: 人類は壁にぶつかっている - ヴィンセント・シュマルバッハ
説明: Anthropic には、すべての AI ラボが切望していた開発者のストーリーがありました。クロードコードは機能しました。オーパスは高価でしたが、散財する価値がありました。ソネットは主力製品として知られていました。

記事本文:
コンテンツにスキップ
ソフトウェア開発者
ホーム
サービス
カスタム Laravel アプリ開発
Laravelのパフォーマンスの最適化
技術チームのリーダーシップとメンタリング
Laravel開発コンサルティング
Anthropic には、あらゆる AI ラボが切望していた開発者のストーリーがありました。クロードコードは機能しました。オーパスは高価でしたが、散財する価値がありました。ソネットは主力製品として知られていました。しばらくの間、同社は OpenAI よりもビルダーのことをよく知っている研究所のように感じられました。
私は、タスクをプル リクエストに変換するための AI コーディング ツールである vroni.com で作業しながらこれを書いています。
それがまさに、ここ数週間が非常に醜い理由です。
6 月 1 日、Anthropic は極秘に S-1 草案を提出しました。同社は 5 月 6 日、Pro、Max、Team、シートベースの Enterprise プランの Claude Code の 5 時間制限を 2 倍にするなど、Claude の使用制限を引き上げると発表しました。その後、製品体験は別の物語を語り始めました。
私自身のクロード購読ログから問題を知っています。現在の 7 日間のローリング ウィンドウで目に見える Opus の入出力は 1,442,423 トークンでした。同じマシン上の他の 2 つの古い重い週は 8,927,831 と 8,477,395 でした。これは、以前の週の 16.2% と 17.0% です。
キャッシュ作成を入れてキャッシュ読み取りのみを除外したとしても、現在のキャッシュされていない合計は 4,563,761 でしたが、以前の週は 22,155,800 と 23,077,312 でした。それは 20.6% と 19.8% です。
それは根底にあるルールの普遍的な証拠ではありません。それは、あるパワー ユーザーが観察したことです。しかし、サブスクリプションが以前ははるかに目に見える Opus の機能を持っていて、公開メッセージが「より高い制限」であったにもかかわらず、現在は急速に燃え尽きている場合、信頼の問題は明らかです。
その後、寓話と神話の混乱が起こりました。
6 月 12 日、アンスロピック社は、米国政府が同社に外国人による Fable 5 と Mythos 5 へのアクセスを停止するよう命令し、事実上すべての顧客に対して両モデルを停止したと発表した。

Hile Anthropic はこれに応じました。 Anthropic 氏は同じ声明で、Fable の安全対策が広すぎると多くのユーザーが不満を述べており、完全な脱獄耐性は現時点では不可能であることを認めた。 Anthropic 社は、Fable がジェイルブレイクを調査し軽減できるよう、顧客データを 30 日間保持する必要があると述べました。
これは残酷な製品パッケージです。主力モデルは通常のアクセスには機密性が高すぎるとみなされ、安全層は正規ユーザーを悩ませ、安全性に関する話はまだ政府を満足させておらず、顧客はモデルの監視が必要であるためにサインアップしなかった保持のトレードオフを被っています。
ソネット 5 はその答えとなるはずです。
6 月 30 日、Anthropic は Claude Sonnet 5 をリリースし、これを Claude Code を含む Free および Pro プランのデフォルト モデルにしました。 Sonnet は、ステッカー価格で Opus よりも安くなります。8 月 31 日まではインプット $2/M、アウトプット $10/M、その後は $3/M、$15/M です。 Opus 4.8 の料金は $5/M と $25/M です。
しかし、ステッカー価格がすべてではありません。 Anthropic 氏は、Sonnet 5 は、より高い労力をかけて、一部のタスクでは Opus 4.8 に匹敵する可能性があると述べています。同じ投稿では、Sonnet モデルは Opus 4.8 や Mythos 5 よりもサイバー性能が著しく劣っており、トークナイザーの変更により、同じ入力が 1.0 ～ 1.35 倍のトークンにマッピングされる可能性があることも述べられています。
Anthropic 独自のコストパフォーマンス表を見ると、困難なタスクに対しては Opus のほうが優れているか、あるいは安価であることがわかります。 Sonnet 5 は、低コストのタスクでのみ興味深いものであり、安価な中国のオープンソース モデルと競合します。
それが日々のコーディングの現実です。開発者は、Anthropic がデフォルト モデルを提供しやすいかどうかを気にしません。彼らはチケットをクローズすることに関心を持っています。
人間は今や両側から圧迫されています。低コストで長期間の作業が必要な場合は、オープンウェイト モデルが急速に改良されています。 Z.ai は GLM-5.2 をそのレーンに配置しています。 OpenAIは

フロンティア コーディング エージェント用の Codex および ChatGPT の GPT-5.5 と、API および Codex 用の GPT-5.6 はプレビュー段階にあり、より広範囲で利用可能になる予定です。
そのため、Anthropic は困難な状況に置かれています。最良のモデルは、ゲート機能があり、高価であるか、安全性やコンプライアンスに関するドラマで覆われています。サブスクリプションはさらに厳しくなります。デフォルト モデルの方が安価ですが、これは労力、トークン化、タスクの完了を気にしない場合に限られます。
上場企業のストーリーには企業の信頼が必要です。開発者の話は善意です。善意を費やすのは簡単です。
おそらく企業顧客は今でも Anthropic の物語を購入しているでしょう。おそらく、規制対象の購入者は安全な姿勢を好むでしょう。おそらく、これで IPO の話がなくなるわけではありません。
しかし、どのツールが実際に下から企業に引き込まれるかを決定するのは開発者です。 Anthropic が人々が望むモデルを隠し続け、サブスクリプションの有用性を低下させ、十分に安くも十分に強力でもないと感じられるモデルをユーザーにデフォルトで使用させ続ければ、そもそも Claude の価値を高めたものを失うことになるでしょう。
安い仕事はオープンモデル向けです。フロンティア コーディング エージェントは Codex 用です。 Anthropic は、なぜ一般の開発者が中間に留まり、圧迫の代償を払わなければならないのかを説明しなければなりません。
Vroni に GitHub の問題、バグレポート、仕様、または大まかなアイデアを提供します。リポジトリを読み取り、変更を計画し、コードを作成し、チェックを実行し、レビュー可能なプル リクエストに向けて作業します。
新しい投稿を公開すると取得します。
私はあなたのプライバシーを尊重します。いつでも購読を解除してください。
このトピックに関してさらに執筆します。
クロード・コードは密かに5倍の価格になった
私は簡単なノルマの質問から始めました。別のクロード アカウントを取得するべきか、それとも今持っているアカウントを使用するべきでしょうか...
クロードコードは中国にリンクされた API ルーターを密かにフィンガープリンティングしている
Claude Code には、リクエストの送信先を示す隠しテストがあります。トリガーは ANTHROPIC_BASE_URL です。これは、人々が次のような場合に使用する環境変数です。
小規模チームのウィル・シー

p 保守できる以上のソフトウェア
Cursor と Claude を使用して、5 人からなるエンジニアリング チームは、請求 Webhook ハンドラー、顧客インポート ツール、Slack アラートなどを構築できるようになりました。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
ソフトウェア、AI、Laravel、SEO、オンライン製品の構築に必要なものについて書いています。 2025 年には、73,000 人がこのブログを読んでいます。一部の投稿は出版物や開発者コミュニティによって取り上げられました。
いくつかの古い投稿が最も注目を集めました。以下の最新の文章とは別に、ここに保管しておきます。
Google Now はデフォルトでコンテンツのインデックスを作成しません
この投稿は、Guardian のコラムと開発者コミュニティでの長い議論につながりました。
AI がテクノロジー負債を指数関数化する
AI によって脆弱なエンジニアリング基盤のコストが下がるのではなく、より高価になる理由。
専門家だけが適切なプロンプトを作成できる
AI が適切に機能するかどうかは、依然として技術的な判断とセンスにかかっています。
実際のコードベースで作業する人向けの Laravel に関する具体的なアドバイス。
ブログの最新の投稿。これは公開すると自動的に更新されます。
Anthropic には、あらゆる AI ラボが切望していた開発者のストーリーがありました。クロードコードは機能しました。オーパスは高価でしたが、その価値はありました。
クロード・コードは密かに5倍の価格になった
私は簡単なノルマの質問から始めました。別のクロード アカウントを取得する必要がありますか、それとも使用していたのでしょうか...
クロードコードは中国にリンクされた API ルーターを密かにフィンガープリンティングしている
Claude Code には、リクエストの送信先を示す隠しテストがあります。トリガーは ANTHROPIC_BASE_URL、環境です...
私の文章や解説について言及、引用、または議論した場所。
私がこれまでに仕事をしたクライアントからの実際のフィードバック。
「ヴィンセントを雇ってください。真剣に。私は何年にもわたって多くの開発者と仕事をしてきました (レンタコーダーを覚えていますか?)。そしてヴィンセントは、それを理解する稀有な種類の Web 開発者です。チャットをして、必要なことを表現すると、彼は自分の意見を持って戻ってきます。

あなたが求めたのです。彼の英語は流暢で、あなたが何を必要としているのかを理解しており、物事を過度に複雑にするつもりはありません。彼と一緒に仕事ができるのは喜びであり、細かな管理をする必要もなく、全体的に素晴らしい経験です。正直に言うと、彼を続けてほしかったが、もう予算がないだけだ――それが契約が終了した唯一の理由だ」
「ヴィンセントとの仕事で際立った点は、彼がどれほど几帳面で効率的だったかということです。彼は私たちの要件を徹底的に掘り下げることから始め、しっかりと理解するまで明確な質問をしました。その後、彼は必要なものを正確に構築し、常に私たちに情報を提供してくれました。コミュニケーションは常に明確で積極的でした。時間と予算の両方において、すべてが順調に進みました。最終結果は素晴らしく、発売以来順調に稼働しています。実践的な焦点と細部へのこだわりの組み合わせが違いを生みます。質の高い開発を行うには、間違いなく Vincent をお勧めします。」
「ヴィンセントをリテイナーに迎えたことは素晴らしいことだ。彼は私たちのためにいくつかの大きな技術的課題に取り組み、特にウェブサイトのパフォーマンスを大幅に向上させました。彼はまた、当社のコア プラットフォーム (Laravel) の最新バージョンへの複雑なアップグレードを専門的に管理し、当社が常に最新かつ安全であることを保証しました。私たちは現在、彼が構築した多言語機能を展開していますが、これは非常に興味深いことです。 Vincent は信頼でき、技術的に熟練しており、サイトの最適な運営と前進を維持するための貴重なパートナーです。彼の専門知識を強くお勧めします。」
「私は Vincent と数か月間仕事をしてきましたが、彼は私が共同作業した中で最も独立した開発者の 1 人です。彼には細かい管理は必要ありません。彼は意思決定を下し、プロジェクトを推進する能力を十分に備えています。そのおかげで、私はビジネスの他の部分に集中できるようになりました。
ヴィンセントはとても働きます

非同期: コミュニケーションは集中したままで、不必要なやり取りはほとんどなく、必要に応じて重要な質問を投げかけます。その結果、開発者はオーバーヘッドなしで開発を行うことができ、これは本当に価値のあることです。」
「私たちは新しいビルドと長年にわたって成長してきた既存のコードベースの両方で、Laravel/Vue プロジェクトで Vincent と定期的に協力しています。どちらの場合でも、彼はクリーンなコードと信頼性の高い実装を提供します。私たちは、彼が成熟したプロジェクトにどれだけ早く着手し、コードベースの既存の状態に適合するソリューションを見つけるかを特に評価しています。絶対にお勧めです。」
「Vincent は扱いやすく、高品質のコードを作成してくれました。お勧めします。」
「Vincent は、複雑なコードベースに飛び込み、すぐに役に立ち、手を握ることなく実際の問題を解決できる稀有な開発者の 1 人です。
彼は、パフォーマンスの最適化、顧客固有の機能、および主要なカスタム レポート システムに取り組みました。また、顧客との議論を明確な技術要件に変換し、チームの他のメンバーがそれに基づいて構築できるように自分の作業を文書化しました。
Vincent はチケットをクローズするだけでなく、製品、エッジケース、顧客への影響を徹底的に考慮します。 SaaS 製品を前進させることができる上級 Laravel 開発者が必要な場合は、彼を強くお勧めします。」
何か見るべきものはありますか？
ラフ版をお送りします。私はそれを計画に落とし込み、それを構築するのを手伝うことができます。
私の本: Laravel による Rapid SaaS
SaaS を数か月ではなく数日で立ち上げます。 Laravel 12 を使用して本番環境に対応した SaaS アプリを構築するための実用的な、BS なしのガイド: 速度、請求、チーム、AI 統合、顧客が料金を支払う前に必要な要素。
Vincent Schmalbach は起業家、ソフトウェア開発者、SEO の専門家です。彼は SaaS 製品、ビジネス アプリ、自動化、技術プロジェクトのレスキューに 15 年以上取り組んでいます。

ソフトウェアとオンライン ビジネスの経験。
投稿
人類は壁にぶつかっている
クロード・コードは密かに5倍の価格になった
クロードコードは中国にリンクされた API ルーターを密かにフィンガープリンティングしている
小規模なチームは、保守できる以上のソフトウェアを出荷することになる
AI により、誤った製品の決定が完成したソフトウェアのように見える
AI 生産性の罠はより多くの生産物を生み出す
ソフトウェアが 90% 安くなった場合、その節約分は誰が回収するのでしょうか?
AI コーディング時代により退屈なテストの価値が高まる
オープンソースのメンテナーには AI 労働のためのスパム フィルターが必要
ジュニア開発者の問題はシニア開発者の問題になりつつある
Vibe コードは出荷前はレガシー コードです
テストに合格した AI コードでもレビューに失敗する可能性がある
カスタム Laravel アプリ開発
Laravel開発コンサルティング
技術チームのリーダーシップとメンタリング
Laravelのパフォーマンスの最適化
Laravelアプリケーションのメンテナンスとサポート

## Original Extract

Anthropic had the developer story every AI lab craved. Claude Code worked. Opus was expensive but worth the splurge. Sonnet was known as the workhorse.

Skip to content
Software Developer
Home
Services
Custom Laravel App Development
Laravel Performance Optimization
Technical Team Leadership & Mentoring
Laravel Development Consulting
Anthropic had the developer story every AI lab craved. Claude Code worked. Opus was expensive but worth the splurge. Sonnet was known as the workhorse. For a time, the company felt like the lab that knew builders better than OpenAI did.
I'm writing this while working on vroni.com , my AI coding tool for turning tasks into pull requests.
That is precisely why the last few weeks are so ugly.
On June 1, Anthropic confidentially filed a draft S-1 . It said on May 6 it was increasing Claude usage limits , including doubling five-hour limits for Claude Code for Pro, Max, Team and seat-based Enterprise plans. Then the product experience began to tell a different story.
I know the trouble from my own Claude subscription logs. Visible Opus input and output in the current seven-day rolling window was 1,442,423 tokens. Two other older heavy weeks on the same machine were 8,927,831 and 8,477,395 . That is 16.2% and 17.0% of the old weeks.
Even if I put cache creation in and only exclude cache reads, the current non-cached total was 4,563,761 versus old weeks of 22,155,800 and 23,077,312 . That is 20.6% and 19.8% .
That is not universal evidence of an underlying rule. That's what one power user has observed. But if a subscription used to have way more visible Opus work, and now burns down fast, while the public message was "higher limits", the trust problem is obvious.
Then the Fable / Mythos mess came along.
On June 12, Anthropic said the US government ordered the company to halt access to Fable 5 and Mythos 5 for foreign nationals, effectively shutting down both models for all customers while Anthropic complied. Many users had complained that Fable's safeguards were overly broad, Anthropic said in the same statement, and admitted that perfect jailbreak resistance is not currently possible. Anthropic said Fable requires 30-day retention of customer data so Anthropic can research and mitigate jailbreaks.
This is a brutal product package: the flagship model is deemed too sensitive for normal access, the safety layer annoys legitimate users, the safety story still doesn't satisfy the government, and customers get a retention tradeoff they didn't sign up for because the model needs monitoring.
Now Sonnet 5 is meant to be the answer.
On June 30, Anthropic released Claude Sonnet 5 and made it the default model for Free and Pro plans, including for Claude Code. Sonnet is cheaper than Opus on sticker price: $2/M input and $10/M output through August 31, then $3/M and $15/M . Opus 4.8 costs $5/M and $25/M .
But the sticker price isn't the whole story. Anthropic says Sonnet 5 can match Opus 4.8 on some tasks, at higher effort. That same post also states that the Sonnet models are significantly worse at cyber than Opus 4.8 and Mythos 5, and the change in the tokenizer means that the same input can map to 1.0-1.35x tokens.
When looking at Anthropic's own cost-performance charts, you'll see Opus looking better or cheaper for hard tasks. Sonnet 5 is only interesting for low-effort tasks, where it competes with cheaper Chinese Open Source models.
That's the daily coding reality. Developers don't care if it is easier for Anthropic to serve the default model. They care about closing the ticket.
Anthropic is now squeezed from both sides. If you want cheap long-horizon work, open-weight models are getting better fast. Z.ai is putting GLM-5.2 right in that lane. OpenAI has GPT-5.5 in Codex and ChatGPT for frontier coding agents, and GPT-5.6 is in preview for API and Codex with wider availability planned.
That puts Anthropic in a rough middle. The best models are gated, costly, or covered in safety/compliance drama. The subscription is tighter. The default model is cheaper, but only if you don't care about effort, tokenization and task completion.
The public company story requires enterprise confidence. The developer story is goodwill. It is easy to spend goodwill.
Perhaps enterprise customers still buy the Anthropic story. Perhaps regulated buyers like the safety stance. Perhaps this doesn't kill the IPO story.
But developers decide which tools actually get pulled from below into companies. If Anthropic keeps hiding the models people want, making subscriptions less useful, and defaulting users to a model that doesn't feel cheap enough or powerful enough, it will lose what made Claude valuable in the first place.
Cheap work is for open models. Frontier coding agents are for Codex. Now Anthropic must explain why ordinary developers should remain in the middle and pay for the squeeze.
Give Vroni a GitHub issue, bug report, spec, or rough idea. It reads the repo, plans the change, writes code, runs checks, and works toward a review-ready pull request.
Get new posts when I publish them.
I respect your privacy. Unsubscribe at any time.
More writing around this topic.
Claude Code Quietly Got 5x More Expensive
I started with an easy quota question. Should I get another Claude account or have I been using the ones I have…
Claude Code Is Quietly Fingerprinting China-Linked API Routers
Claude Code has a hidden test for where your requests are going. The trigger is ANTHROPIC_BASE_URL, the environment variable people use when…
Small Teams Will Ship More Software Than They Can Maintain
Using Cursor and Claude, a 5-person engineering team can now build a billing webhook handler, a customer import tool, a Slack alert…
Your email address will not be published. Required fields are marked *
I write about software, AI, Laravel, SEO, and what it takes to build online products. In 2025, 73k people read the blog. Some posts were picked up by publications and developer communities.
A few older posts got the most attention. I keep them here, separate from the newest writing below.
Google Now Defaults to Not Indexing Your Content
The post that led to a Guardian column and a long developer-community discussion.
AI Exponentializes Your Tech Debt
Why AI makes weak engineering foundations more expensive, not less.
Only Experts Can Write Good Prompts
Good AI work still depends on technical judgment and taste.
Specific Laravel advice for people working in real codebases.
The newest posts from the blog. This updates automatically when I publish.
Anthropic had the developer story every AI lab craved. Claude Code worked. Opus was expensive but worth the...
Claude Code Quietly Got 5x More Expensive
I started with an easy quota question. Should I get another Claude account or have I been using...
Claude Code Is Quietly Fingerprinting China-Linked API Routers
Claude Code has a hidden test for where your requests are going. The trigger is ANTHROPIC_BASE_URL, the environment...
Places that have mentioned, quoted, or discussed my writing and commentary.
Actual feedback from clients I've worked with.
"HIRE VINCENT. Seriously. I've worked with many many developers over the years (remember rentacoder?), and Vincent is that rare breed of web developer who just gets it. You have a chat, you express what you need, and he comes back with what you asked for. His english is fluent, he understands what you need, and he's not looking to overcomplicate things. It's a joy and pleasure to work with him, not to have to micromanage and just overall a great experience. Honestly I wish I could keep him on we just don't have the budget anymore - that is the only reason the contract ended."
"What stood out about working with Vincent was how methodical and efficient he was. He started by really digging into our requirements, asking clarifying questions until he had a solid understanding. He then built exactly what was needed and kept us informed throughout. Communication was clear and proactive at all times. Everything stayed on track, both in terms of time and budget. The end result was excellent and has been running smoothly since launch. It's that combination of practical focus and attention to detail that makes the difference. I would definitely recommend Vincent for quality development."
"Having Vincent on retainer has been excellent. He's tackled some major technical challenges for us, notably significantly improving our website's performance. He also expertly managed a complex upgrade of our core platform (Laravel) to the latest version, ensuring we stay current and secure. We're now rolling out the multilingual capabilities he built, which is exciting. Vincent is reliable, technically skilled, and a valuable partner for keeping our site running optimally and moving forward. Highly recommend his expertise."
"I've been working with Vincent for several months, and he's one of the most independent developers I've collaborated with. He doesn't need micromanagement - he's fully capable of making decisions and driving the project forward, which has freed me up to focus on other parts of the business.
Vincent works very well asynchronously: communication stays focused, there is little unnecessary back-and-forth, and he still raises the important questions when needed. The result is a developer who delivers without overhead, and that's been genuinely valuable."
"We work with Vincent regularly on Laravel/Vue projects, both new builds and existing codebases that have grown over the years. In both cases, he delivers clean code and reliable implementation. We especially value how quickly he gets into mature projects and finds solutions that fit the existing state of the codebase. Absolutely recommended."
"Vincent was easy to work with and produced high-quality code. Would recommend."
"Vincent is one of those rare developers who can jump into a complex codebase, become useful immediately, and solve real problems without hand-holding.
He worked on performance optimization, customer-specific features, and a major custom reporting system. He also turned customer discussions into clear technical requirements and documented his work so the rest of the team could build on it.
Vincent does not just close tickets, he thinks through the product, the edge cases, and the customer impact. If you need a senior Laravel developer who can move a SaaS product forward, I strongly recommend him."
Have something I should look at?
Send the rough version. I can help turn it into a plan and then build it.
My Book: Rapid SaaS with Laravel
Launch Your SaaS in Days, Not Months. A practical, no-BS guide to building production-ready SaaS apps with Laravel 12: speed, billing, teams, AI integration, and the parts you need before customers can pay you.
Vincent Schmalbach is an entrepreneur, software developer, and SEO expert. He works on SaaS products, business apps, automation, and technical project rescue, with 15+ years of software and online business experience.
Posts
Anthropic Is Hitting a Wall
Claude Code Quietly Got 5x More Expensive
Claude Code Is Quietly Fingerprinting China-Linked API Routers
Small Teams Will Ship More Software Than They Can Maintain
AI Makes Bad Product Decisions Look Like Finished Software
The AI Productivity Trap Is More Output
If Software Gets 90 Percent Cheaper, Who Captures the Savings?
The AI Coding Era Makes Boring Tests More Valuable
Open Source Maintainers Need a Spam Filter for AI Labor
The Junior Developer Problem Is Becoming a Senior Developer Problem
Vibe Code Is Legacy Code Before It Ships
AI Code That Passes Tests Can Still Fail Review
Custom Laravel App Development
Laravel Development Consulting
Technical Team Leadership & Mentoring
Laravel Performance Optimization
Laravel Application Maintenance & Support
