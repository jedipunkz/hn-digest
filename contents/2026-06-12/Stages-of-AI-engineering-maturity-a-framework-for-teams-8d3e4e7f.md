---
source: "https://upsun.com/blog/8-stages-ai-engineering-maturity/"
hn_url: "https://news.ycombinator.com/item?id=48504959"
title: "Stages of AI engineering maturity: a framework for teams"
article_title: "8 Stages of AI Engineering Maturity for Teams | Upsun"
author: "fabpot"
captured_at: "2026-06-12T16:07:36Z"
capture_tool: "hn-digest"
hn_id: 48504959
score: 7
comments: 1
posted_at: "2026-06-12T14:52:40Z"
tags:
  - hacker-news
  - translated
---

# Stages of AI engineering maturity: a framework for teams

- HN: [48504959](https://news.ycombinator.com/item?id=48504959)
- Source: [upsun.com](https://upsun.com/blog/8-stages-ai-engineering-maturity/)
- Score: 7
- Comments: 1
- Posted: 2026-06-12T14:52:40Z

## Translation

タイトル: AI エンジニアリングの成熟段階: チームのフレームワーク
記事のタイトル: チームの AI エンジニアリング成熟度の 8 段階 |ウプサン
説明: ほとんどのチームはすでに AI を使用しています。問題は、組織が一緒に動いているかどうかです。チーム全体を混沌から自律エージェントに導くためのフレームワーク。

記事本文:
チームの AI エンジニアリング成熟度の 8 段階 | Upsun 以前の Platform.sh
ソリューション ソリューション アプリケーションの最新化
DevOps とプラットフォーム エンジニアリング
パートナーの専門知識 戦略的実装 upsun の実装アーキテクトを活用して、複雑なプロジェクトを自信を持って立ち上げられるようにする明確さと技術的な道筋を見つけてください。私たちがどのようにお手伝いできるかをご覧ください
企業コミュニティ デベロッパー センター 最新の Upsun 開発リソースのハブ: 選択したフレームワークと言語のデモ、ドキュメント、記事を検索してください。
デモを見る 無料トライアル ブログ ブログ ブログ 製品事例 ニュース インサイト ブログ AI エンジニアリングの成熟度の 8 段階: チームのためのフレームワーク
Fabien Potencier 最高製品および技術責任者 シェア 数か月前、Steve Yegge 氏が AI 支援開発の 8 つのレベルを公開しました。それを読んだ瞬間にピンときました。なぜなら、私自身、オートコンプリートからエージェントの実行まで、ステップごとに移行するまさにその過程を経験したからです。 AI の信頼勾配として枠組み化されたこのソリューションは、私たちのほとんどが既に名前も付けずに経験していた事柄に対する語彙をついに業界に与えました。まだ読んでいない場合は、後で読むために保存してください。
Yegge のレベルは、個々の開発者に焦点を当てています。それが重要です。エンジニアリング チームを率いるときは、1 つの信頼勾配を管理しているのではなく、その分布全体を管理していることになります。一日中並列エージェントを実行している人がいて、AI をより高級なオートコンプリートだと考えている人の隣に座っていて、どちらもコードを出荷しています。これが、誰もが言い続ける生産性の 10 倍の向上が組織レベルではほとんど現れない理由です。難しいのは、迅速に行動する人が混乱を引き起こしたり、遅い人が静かに後れをとったりすることなく、組織全体を連携させることです。
そこで私たちは同じPRをしました

Yegge が説明する個人の開発者ではなく、チームとそのチームが所属する組織を別の角度から見ました。「誰も何も決めていない」から「工場が自動的に動く」までの 8 つの成熟段階。
語彙について簡単に説明します。なぜなら、私たちは語彙に重点を置いているからです。私たちが「チーム」と言うとき、同じ結果に向かって協力する人々のグループを意味します。エンジニアはもちろんですが、製品、設計、QA など、ソフトウェアの出荷に必要な人材も含まれます。私たちが「組織」と言うときは、それらのチームの集合体全体と、その上に位置するリーダーシップ、予算、ガバナンスを意味します。
そして、これが単なる開発者の話ではなく、SDLC の話である理由は次のとおりです。新しいソフトウェア開発ライフサイクルは開発者だけに関するものではありません。それはチーム全体に関することだ。製品、デザイン、エンジニアリングの間の古い境界線は曖昧になりつつあります。プロダクト マネージャーは、誰も読まない 3 ページの仕様書を書く代わりに、動作するプロトタイプをバイブコーディングできるようになりました。デザイナーは、デザイン システムと Figma ファイルを渡すだけでなく、実際の HTML と CSS を出荷することができます。この変化こそが重要なポイントであり、個々の開発者の観点から考えるだけでは十分ではなくなるのはこのためです。
マップ全体を見て回る前に、マップ全体をレイアウトしましょう。
リーダーシップは実際の地位を確立していません。せいぜい、ライセンスのバッチを購入してそこで停止する程度でした。開発者はガードレールなしで習慣を形成します。つまり、開いているチャット ウィンドウにコードを貼り付け、共有コンテキスト ファイル、エージェントのセットアップ、管理 API キーはありません。これは学習段階であり、それでいいと思います。人々はこれらのモデルで何ができるか、何ができないかについての直感を構築しています。その直感が今の本当のリターンです。沈黙を何もしないことと誤解しないでください。開発者はすでに自分たちで決定しています。
ステージ 1 は、組織がどのようなものであるかについてでした。

決まっていない。ステージ2は個人が自分で決めたことについてです。あるエンジニアは、エージェントに出力の大部分を静かに処理させ、プロンプト、カスタム スキル、および自分のマシン上に保管している慎重に調整された AGENTS.md の個人的な隠し場所を実行させています。彼らの隣にいる人は2年間何も変わっておらず、個人差は常に正常に見えたため、誰もそれを指摘しませんでした。これは、何もしないことが自由になる最後の段階です。ずれが目に見えるようになると、それはチームレベルの問題に変わります。
以前は個人間で発生していた差が、今ではチーム全体の間で発生し、全員が確認できる配信カレンダーに表示されます。あるチームは、AGENTS.md ファイルを共有し、MCP サーバーを接続し、再利用可能なスキルを構築しました。スループットが飛躍的に向上します。隣のチームは今でも 2 年前と同じようにコードを書いています。ツールのギャップとして始まったものは、やがて恨みに変わり、それを修復するのはさらに難しくなります。
ステージ 4: 標準化への賭け
リーダーの真の取り組みが必要となる最初の段階は、AI が意図的に構築される能力になることです。ここでは 3 つのことが重要です。コンテキスト エンジニアリングは、AGENTS.md ファイルの共有とリポジトリ内の厳選されたスキルとプロンプト ライブラリを使用して明示的な作業になるため、すべての開発者が AI に同じレッスンを教えるのではなく、チームが知識を 1 回エンコードします。セキュリティとガバナンスは規模より優先されます。SSO と SCIM、シークレット スキャン、エージェントの出力で実行される PR ゲート、監査ログ、ゲートウェイの背後にあるモデルとツールの承認済みリストは、人々がフルスピードで実行する前に構築されます。また、トレーニングは 1 回限りのワークショップではなく継続的に行われ、通常はステージ 2 のチャンピオンを中心に構築されます。
ステージ 5: ワークフローの再設計
ステージ 4 ではツールをアップグレードします。ステージ 5 では工場の現場を再設計し、実際の作業方法を変更します。スペックファー

st がデフォルトになります。仕様はエージェントのエントリ ポイントとしてリポジトリ内に存在し、エージェントは曖昧なチケットから若手開発者よりもうまく作業することはできません。コードレビューは、行ごとの質問から、「これは仕様と一致していますか?」というリスクベースの質問に移行します。テストはそれを証明しますか？間違っている場合、爆発範囲はどれくらいですか? CI ゲートは、エージェントが開いた PR を人間の PR とまったく同じように扱い、評価は単体テストのすぐ隣で実行を開始します。役割がコードの作成からレビューとオーケストレーションに移行する間は、生産性の低下が予想されます。それは正常です。速度だけでなく品質にも注目してください。低品質でより多くのコードが出荷されるのは速度ではなく、インシデントがより早く到着することです。
ステージ 5 では、チームの働き方が変わりました。ステージ 6 では、チームの構成とチームの連携方法が変わります。スプリントには 3 人のエンジニアと 5 人のエージェントの枠があり、エンジニアはプログラマーというよりもプロダクト マネージャーのように見えます。仕様を作成し、仕様について AI と議論し、テストをチェックします。 TDD は、あると便利なものではなく、依存関係になります。エージェントはテストに合格するように最適化するため、弱いスイートは攻撃され、本番環境がそうなるまでわかりません。分離されたサンドボックスで実行される並列エージェントは、単一のスプリントでホスティング コストの 4 分の 1 よりも多くのコンピューティングを消費する可能性があるため、トークンの予算と実行ごとの可観測性はオプションではなくなります。そして、共有コンテキストは実際のインフラストラクチャ (プロジェクト メモリ、スキル、プロンプト ライブラリ、MCP サーバー) となり、チーム資産として維持され、チーム間で調整されます。
ステージ6のセリフは誰がペンを持つかです。現在、エージェントは人間の権限をほとんど持たずに作業単位全体を作成して出荷しており、人間は運転手から監督者に戻ります。ほとんどのチームはまだ子守り中です。エージェントは誰かのラップトップ上の端末からキックオフされ、ローカルでループし、開発マシンから PR を開きます。SHA はありません。

ランタイムは赤で、何が実行されたのか、なぜ実行されたのかについての中央記録はありません。このローカル マシンの詳細が、この段階と次の段階を分ける唯一の要素です。
ステージ 8: 自律型工場
エージェントはラップトップから共有インフラストラクチャに移行し、サンドボックス環境でスケジュールされた実行と一元化されたログとトレースにより、これまでベビーシッターが必要だった移行が夜間に実行され、朝に結果が報告されます。定期的なジョブは、定義された成功基準と、何かが失敗した場合の自動エスカレーションを備えたファーストクラス (依存関係の更新、セキュリティ パッチ、テストの拡張) になります。 Eval が製品になります。人々の頭の中にあった知識は、エージェント設定、スキル、およびすべてのマージを制御する評価スイートの中に組み込まれ、一度エンコードされ、自動的に適用されます。あなたの基準はもう誰かの頭の中に存在しません。彼らはシステムの中で生きている
単一のステージにきれいに収まる組織はありません。いくつかのチームはステージ 6 にいますが、別のチームはまだステージ 1 にしっかり留まっています。この番号は重心であり、全員にラベルを付けるわけではありません。
また、ステージをスキップしないでください。ガバナンス、テスト、およびその下にある共有コンテキストを持たずに自律型エージェントをそのまま使用すると、単に別のキャンセルされたプロジェクトになってしまいます。しかし、一度に 1 つの段階に進むには、もっと微妙な理由があります。各段階にはそれぞれ特有の痛みがあり、その痛みが教訓なのです。速いチームと遅いチームの間の摩擦が、企業の標準化を促すのです。エージェントの出力を 1 行ずつレビューするという大変な作業により、チームはワークフローを再設計することになります。痛みを無視すると、次に進む理由を失います。痛みが生じ始めている現在のステージほど、次のステージに進む理由となるものはありません。
そして、ずっと懐疑的であり続けてください。これに関する誇大宣伝は大騒ぎです。証拠は、はるかに静かです。したがって、私たちの約束も含め、ここでの約束はすべて慎重に扱ってください。
質問

重要なのはAIを採用するかどうかではない。ほとんどのチームはすでに持っています。本当の問題は、その下のシステムが拡張する価値があるほど十分に優れているかどうかです。
それがAIなのですから。私はそれを確信しています、それはアンプです。それは、良い習慣も悪い習慣も同様に、すでに存在しているものを加速させます。私たちがこれまで観察してきた組織は、品質を損なうことなく速度を上げていますが、明確なサービス所有権、包括的なテスト、文書化されたサービス、自動化された標準という同じ基盤を共有しています。これはどれも新しいものではありません。私たちはこれらのベストプラクティスを何年も説教してきました。 AI によって、彼らの不在が無視できなくなっただけです。
これが目指すのは自律型工場です。最も難しいのはエージェントそのものではありません。エージェントを個々のラップトップから取り外し、共有インフラストラクチャ上で実行できるほど十分に信頼できるかどうかです。いつもと同じ方法で、証拠を使って構築します。 1 つの評価、1 つの繰り返しタスク、一度に 1 つの検証済みデプロイメント。
最新の更新情報やニュースを入手するには、毎月のニュースレターを購読してください。
あなたの最高の作品
ちょうど地平線上にあります
会社 トラスト センター イベント 私たちについて 求人 サステナビリティ オープンソース FAQ 比較 Vercel の代替 Amazee の代替 Heroku の代替 Pantheon の代替 マネージド ホスティングの代替 Fly.io の代替 レンダーの代替 AWS の代替 Acquia の代替 DigitalOcean の代替 製品概要 サービス サポート プレビュー環境 マルチクラウドおよびエッジ Git ベースの自動化 可観測性とプロファイリング セキュリティとコンプライアンス スケーリングとパフォーマンス バックアップとデータのリカバリ チームとアクセスの管理 CLI、コンソール、API 統合と Webhook 価格価格計算ツール ユースケース アプリケーションのモダナイゼーション eコマースと CMS ホスティング マルチサイト管理 コンプライアンスとガバナンス DevOps とプラットフォーム エンジニアリング AI と自動化

業界 金融サービス 政府 SaaS 製造業 高等教育 旅行とホスピタリティ 小売業 ヘルスケア パートナー 戦略的実施 代理店パートナーになる パートナー ポータル リソース ケーススタディ ブログ 開発者センター ドキュメント コミュニティ フォーラム クラウド リージョン サポート お問い合わせ 月刊ニュースレターにご参加ください

## Original Extract

Most teams already use AI. The question is whether your org is moving together. A framework for taking your whole team from chaos to autonomous agents.

8 Stages of AI Engineering Maturity for Teams | Upsun Formerly Platform.sh
Solutions Solutions Application modernization
DevOps and platform engineering
Partners Expertise Strategic Implementation Tap upsun's implementation architects for the clarity and technical pathways that turn complex projects into confident launches. See how we can help
Company Community Developer center Hub for the latest Upsun dev resources: find demos, docs, and articles for your chosen frameworks and languages Check it out
Watch a demo Free trial Blog Blog Blog Product Case studies News Insights Blog The 8 stages of AI engineering maturity: a framework for teams
Fabien Potencier Chief Product and Technology Officer Share A few months ago, Steve Yegge published his 8 levels of AI-assisted development, and it clicked the moment I read it, because I had lived that exact progression myself, moving from autocomplete to running agents one step at a time. Framed as an AI trust gradient, it finally gave the industry a vocabulary for something most of us were already going through without a name for it. If you haven’t read it, save it for later.
Yegge's levels focus on the individual developer. And that’s the thing: when you’re leading an engineering team, you’re not managing one trust gradient, you’re managing a whole distribution of them. You’ve got someone running parallel agents all day, sitting next to someone who still thinks of AI as a fancier autocomplete, and both of them are shipping code. This is why the 10x productivity boost everyone keeps talking about almost never shows up at the org level. The hard part is getting the whole organization to move together, without the fast movers creating chaos and the slow adopters quietly falling behind.
So we took the same principles and looked at them from another angle: not the individual developer Yegge describes, but the team and the organization the team lives in. Eight stages of maturity, from “nobody has decided anything” to “the factory runs itself”.
A quick word on vocabulary, because we lean on it throughout. When we say team , we mean a group of people working together toward the same outcome: engineers, yes, but also product, design, QA, whoever it takes to ship software. When we say organization , we mean the entire collection of those teams, plus the leadership, budgets, and governance that sit above them.
And here’s the part that makes this an SDLC story and not just a developer story: the new software development life cycle isn’t just about developers. It’s about the whole team. The old boundaries between product, design, and engineering are getting blurry. A product manager can now vibe-code a working prototype instead of writing a three-page spec nobody reads. A designer can ship real HTML and CSS, not just a design system and a Figma file to hand off. That shift is the whole point, and it’s why thinking in terms of individual developers stops being enough.
Let’s lay out the whole map before we walk through it.
Leadership hasn’t taken a real position; at most, it bought a batch of licenses and stopped there. Developers form habits without guardrails: pasting code into whatever chat window is open, no shared context files, no agent setup, no managed API keys. This is a learning phase, and I think that’s fine: people are building intuition for what these models can and can't do. That intuition is the real return right now. Don’t mistake silence for inaction. Your developers have already decided for themselves.
Stage 1 was about what the organization hadn’t decided; Stage 2 is about what individuals have decided on their own. One engineer quietly has agents handling a big chunk of their output, running off a personal stash of prompts, custom skills, and a carefully tuned AGENTS.md they keep on their own machine. The person beside them hasn’t changed anything in two years, and nobody flags it because individual variation has always looked normal. This is the last stage where inaction is free. Once the drift becomes visible, and it will, it turns into a team-level problem.
The gap that used to run between individuals now runs between whole teams, and it shows up on the delivery calendar where everyone can see it. One team has shared AGENTS.md files, wired up MCP servers, and built reusable skills. Its throughput jumps. The team next door still writes code like it's two years ago.. What starts as a tooling gap hardens into resentment, and that's much harder to repair.
Stage 4: The standardization bet
The first stage that requires real commitment from leadership is that AI becomes a capability you build deliberately. Three things matter here. Context engineering becomes explicit work, with shared AGENTS.md files and a curated skill and prompt library in the repos, so the team encodes its knowledge once instead of every developer teaching the AI the same lessons. Security and governance come before scale: SSO and SCIM, secret scanning, PR gates that run on agent output, audit logs, and an approved list of models and tools behind a gateway, built before people run at full speed. And training is ongoing, not a one-off workshop, usually built around the champions from Stage 2.
Stage 5: The workflow redesign
Stage 4 upgrades the tools; Stage 5 redesigns the factory floor, changing how the work actually gets done. Spec-first becomes the default: the spec lives in the repo as the agent’s entry point, and an agent can’t work from a vague ticket any better than a junior developer can. Code review shifts from line-by-line to risk-based questions: Does this match the spec? Do the tests prove it? What’s the blast radius if it’s wrong? CI gates treat agent-opened PRs exactly like human ones, and evals start running right next to the unit tests. Expect a productivity dip while the role moves from writing code to reviewing and orchestrating; that’s normal. Watch quality, not just velocity: more code shipped at lower quality isn’t speed, it’s incidents arriving sooner.
Stage 5 changed how a team works; Stage 6 changes what a team is made of and how teams coordinate. A sprint might hold three engineers and five agent slots, and the engineer now looks more like a product manager than a programmer: writing specs, arguing with the AI about them, checking the tests. TDD becomes a dependency, not a nice-to-have: agents optimize for passing tests, so a weak suite gets gamed, and you won’t find out until production does. Parallel agents running in isolated sandboxes can burn more compute in a single sprint than a quarter of the hosting cost, so token budgets and per-run observability stop being optional. And shared context becomes real infrastructure (project memory, skills, prompt libraries, MCP servers), maintained as team assets and coordinated across teams.
The line from Stage 6 is who holds the pen. The agents now write and ship whole units of work with barely any human authorship, and the human steps back from driver to supervisor. Most teams are still babysitting: agents kicked off from a terminal on someone’s laptop, looping locally and opening PRs from a dev machine, with no shared runtime and no central record of what ran or why. That local-machine detail is the one thing separating this stage from the next.
Stage 8: The autonomous factory
Agents move off laptops and onto shared infrastructure, with scheduled runs in sandboxed environments and centralized logs and traces, so a migration that used to need babysitting runs overnight and reports its results in the morning. Recurring jobs become first-class (dependency updates, security patches, test expansion) with defined success criteria and automatic escalation when something fails. Eval becomes the product: the knowledge that lived in people's heads now lives in agent configs, skills, and eval suites that gate every merge, encoded once and enforced automatically. Your standards no longer live in someone's head. They live in the system
No organization sits cleanly on a single stage. You’ll spot some teams at Stage 6 and another still firmly in Stage 1. The number is a center of gravity, not a label you pin on everyone.
Also, you shouldn’t skip stages. Go straight for autonomous agents without the governance, testing, and shared context underneath, and you just become another canceled project. But there’s a subtler reason to go one stage at a time: each stage hurts in its own particular way, and that pain is the lesson. The friction between a fast team and a slow one is what pushes a company to standardize. The grind of reviewing agent output line by line is what convinces a team to redesign the workflow. Skip the pain, and you lose the reason to move on: nothing makes the case for the next stage like the current one starting to hurt.
And stay skeptical the whole way. The hype around this is loud; the evidence, much quieter. So treat every promise here, including ours, with care.
The question isn't whether to adopt AI. Most teams already have. The real question is whether the system underneath is good enough to be worth amplifying.
Because that's what AI is. I'm convinced of it: an amplifier. It accelerates what’s already there, the good practices and the bad ones alike. The organizations we’ve watched gain velocity without losing quality all share the same foundation: clear service ownership, comprehensive testing, documented services, and automated standards. None of this is new. We’ve been preaching these best practices for years. AI just made their absence impossible to ignore.
The autonomous factory is where this is heading. The hardest part won’t be the agents themselves: it’ll be trusting them enough to take them off individual laptops and run them on shared infrastructure. You build it the same way you always have, with evidence. One eval, one recurring task, one verified deployment at a time.
Subscribe to our monthly newsletter for the latest updates and news.
Your greatest work
is just on the horizon
Company Trust center Events About us Jobs Sustainability Open source FAQ Compare Vercel alternative Amazee alternative Heroku alternative Pantheon alternative Managed hosting alternative Fly.io alternative Render alternative AWS alternative Acquia alternative DigitalOcean alternative Product Overview Services Support Preview environments Multi-cloud and edge Git-driven automation Observability and profiling Security and compliance Scaling and performance Backups and data recovery Team and access management CLI, console, and API Integrations and webhooks Pricing Pricing calculator Use cases Application modernization eCommerce and CMS hosting Multi site management Compliance and governance DevOps and platform engineering AI and automation Industries Financial services Government SaaS Manufacturing Higher education Travel and hospitality Retail Healthcare Partners Strategic implementation Become an agency partner Partner portal Resources Case studies Blog Developer center Documentation Community forum Cloud regions Support Contact us Join our monthly newsletter
