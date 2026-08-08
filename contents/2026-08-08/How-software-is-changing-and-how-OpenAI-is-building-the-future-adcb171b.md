---
source: "https://every.to/p/openai-infrastructure"
hn_url: "https://news.ycombinator.com/item?id=49221375"
title: "How software is changing – and how OpenAI is building the future"
article_title: "Inside OpenAI’s Race to Reinvent Software Development for the Agent Era"
author: "samuel_kx0"
captured_at: "2026-08-08T13:38:26Z"
capture_tool: "hn-digest"
hn_id: 49221375
score: 1
comments: 1
posted_at: "2026-08-08T12:49:53Z"
tags:
  - hacker-news
  - translated
---

# How software is changing – and how OpenAI is building the future

- HN: [49221375](https://news.ycombinator.com/item?id=49221375)
- Source: [every.to](https://every.to/p/openai-infrastructure)
- Score: 1
- Comments: 1
- Posted: 2026-08-08T12:49:53Z

## Translation

タイトル: ソフトウェアはどのように変化しているのか – そして OpenAI はどのように未来を構築しているのか
記事のタイトル: エージェント時代のソフトウェア開発を再発明する OpenAI の競争の内部
説明: OpenAI のインフラストラクチャ チームが AI 生成コードの洪水にどのように適応しているかを独占的に紹介します。

記事本文:
コンテンツへスキップ 次に何が来るのか
未来の構築に役立つ新しいアイデアが毎日受信箱に届きます。
私の個人情報を販売または共有しないでください
このサイトは reCAPTCHA によって保護されており、Google プライバシー ポリシーと利用規約が適用されます。
当社はデフォルトで分析ツールと広告ツールを使用します。これはいつでも更新できます。
オプションの追跡カテゴリを管理します。サイトが機能するために必要な Cookie はオンのままになります。
エージェント時代のソフトウェア開発を再発明する OpenAI の競争の内部
Eロゴ
これは、サンフランシスコのほぼすべての看板の約束です。エージェントの未来が到来し、人間のエンジニアの必要性が自動化されます。
よくあることですが、現実は看板で信じられているよりも複雑です。
真夜中のある時、研究者のラップトップで実行されているコーディング エージェントが、データ エクスポートの失敗に関する Slack メッセージを投稿しました。
研究者は眠っていた。その夜当番だったインフラストラクチャ エンジニアの James Katz も同様でした。
一方、エージェントは遅い時間になっても動じなかった。問題をデバッグするために必要なツールと内部システムにアクセスできました。
そして、応答を待っている間に作業を開始し、インシデントを診断し、修正の可能性を調査し、障害に遭遇するたびに再戦略を立てるという取り組みを熱心に説明しました。
カッツ氏やチームの他の誰も介入することなく、朝までに研究者に必要なデータを渡していた。
これは、サンフランシスコのほぼすべての看板の約束です。エージェントの未来が到来し、人間のエンジニアの必要性が自動化されます。
よくあることですが、現実は看板で信じられているよりも複雑です。
エージェントがすべてのコードを記述する世界に向けてソフトウェア開発を再設計する OpenAI の競争の内部
B

y Laura Entis · サラ・デラゴンとサラ・ジェイ・ハリデイによる写真
サンフランシスコのミッションベイ地区にある、洗練された何の変哲もないオフィスビルの外で、中年女性がインターホンに向かって話している。
「昼食に、奥様？」スピーカーから小さな声が響きわたります。会話がそれ以上進む前に、背の高いひげを生やした警備員がドアのところで取りなしてくれます。 「ここはレストランではありません。」
実際、そうではありません。 OpenAI の本社として知られる 1515 3rd Street では、建物に入るのに招待状が必要です。中に入ると、ちょっとした企業の 5 つ星ホテルのような内装のロビーに出迎えられ、アクセントブックやレモンの輪が付いたウォーター クーラーが完備されています。
OpenAI はレストランではないかもしれませんが、数階上のランチタイムです。同社が掲げる人工知能（AGI）の実現という使命に向かって取り組んでいる従業員たちは、ハーベストボウル、タコス、炒め物、そしてサンフランシスコ最大の直火式ピザ窯から取り出したてのホットスライスを求めて列に並んでいる。
OpenAI がライバルの Anthropic と競争して新しいフロンティア モデルをますます急速にリリースする中、同社のエンジニアは AI 以前の 10 倍、さらには 100 倍の速さでコードを出荷できるようになりました。同社のコーディング エージェントおよびデスクトップ アプリである Codex は、その能力をエンジニアリングを超えて拡張します。プロダクト マネージャー、データ サイエンティスト、研究者、市場投入従業員は、OpenAI の内部システムで実行されるコードの作成、変更、デバッグ、送信にこれを使用します。カフェテリアの多くは携帯電話に釘付けになっており、おそらくテキストメッセージを送信したり、ソーシャルメディアをスクロールしたり、あるいは OpenAI のモバイルアプリを通じてエージェントを調整したりしている可能性が高いでしょう。
新しい製品やツールを構築する人にとって、今は心躍る時期です。しかし、このすべての構築を可能にする内部システムの保守を担当するエンジニアにとっては、その気持ちは非常に厳しいものです。

浮いているようにしてください。
応用インフラストラクチャ データ プラットフォーム チームを率いる Emma Tang 氏は言います。Tang 氏のグループは、ChatGPT や Codex を含むすべての OpenAI 製品の配管を流れるデータを移動、分析、保護するインフラストラクチャを維持しています。チームのプル リクエスト (PR)、または共有コードベースへの変更提案の量は、過去 1 年間で 5 ～ 10 倍に急増しました。
そのコードの重要な部分は、組織の非技術部門の人々から提供されます。 「製品データの連続ストリームを分析するシステムである Apache Flink で構築されたプログラムである Flink ジョブを生成するユーザーがいますが、彼らは Flink のことさえ知りません」と Tang 氏は言います。変化のペースが速くなるにつれて、データ処理の遅延からダッシュボードの読み込み失敗に至るまで、内部の信頼性の問題も増加しています。そして、何かが壊れた場合、作成者は自分が出荷したものを完全には理解していないため、手助けできないことがよくあります。
最近まで、運用部門や市場開拓チームの誰かがカスタムの社内アプリを必要とした場合、それを作成してくれるエンジニアを探す必要がありました。今では、社内の誰もが「週末に作成した実際に動作するプロトタイプを持って月曜日に出勤できる」とタン氏は言います。彼女のエンジニアは教育者や調停者としての役割をますます高めており、同僚が AI によって生成されたツールを信頼できるソフトウェアに変えるのを支援しています。たとえば、バイブコード化されたアプリは数秒ごとに不必要に更新される可能性がありますが、そのアプリが取得する分析システムは 1 日に 1 回しか更新されません。データを真にライブにするということは、データを別の実稼働システムに接続することを意味します。これにより、技術的な背景がない人にはほとんど気づかれない疑問が生じます。たとえば、「どこでデータを実行するのか?」、「サービスの速度を毎回低下させずに情報を取得するにはどうすればよいか」などです。

「他に頼っているのは？」「誰がそれを維持するの？」
特に最後の質問は、タン氏のチームに降りかかることがよくあります。 AI により、社内の誰もが会社の共有データ インフラストラクチャに依存するアプリやシステムを構築できるようになりますが、そのコードが失敗すると、彼女のチームが問題を引き継ぐことになります。
「私たちはすべてをキャッチする存在になります」とタン氏は言います。 「私たち以外にはもう誰もそれを所有していません。」
データ インフラストラクチャ エンジニアリングの責任者、Emma Tang 氏は次のように述べています。写真提供：サラ・デラゴン。
James Katz、技術スタッフのメンバー。写真提供：サラ・ジェイ・ハリデー。
モリー・オコナー、技術スタッフのメンバー。写真提供：サラ・ジェイ・ハリデー。
大洪水に対する応用インフラストラクチャ チームの最初の対応も、それを作成したツールである Codex でした。
GPT-5.5 のリリースから数週間以内に、OpenAI 内での Codex デスクトップ アプリの採用率はほぼ 100% に達しました。エンジニアは現在、アプリ内から作業の大部分を実行し、AI が生成したコード、AI が生成したテスト (コードが正常に動作することを確認する自動チェック)、および AI が提案した修正をレビューしています。 OpenAI の従業員にはトークン制限がないため、エンジニアはコストや効率について深く考えずに、Codex で UI バグのバッチをキューに入れ、複数のエージェントに個別に修正を作成させ、製品のテスト バージョンで実行し、ブラウザで影響を受けるフローをクリックして、結果として生じる問題にフラグを立てることができます。すべて最も強力なモデルを使用しています (多くの従業員はモデルが一般公開される前にモデルへのアクセスが許可されています)。
2024 年 10 月に Molly O'Connor が OpenAI の応用データ プラットフォーム チームに加わったとき、典型的な 1 日は、インシデントの呼び出しを受け、デバッグし、次に進むことが含まれていました。今後は、修正自体を行うとともに、将来的に同様の問題を修正できるように、会社の共有コーデックス手順も更新する可能性があります。

再。オコナー氏は、プル リクエストで繰り返し発生する問題、特にシステムの信頼性を低下させる可能性のある問題を発見したときにも同じことを行います。
OpenAI のデータ インフラストラクチャの信頼性と安全性を確保するという彼女の役割は変わっていません。しかし、現在は、データ パイプラインの構築とトラブルシューティング、大規模な移行の実行、データ処理システムの更新のテストと展開についての知識を Codex が従うことができる指示に変換することが、より多くの仕事に含まれています。
Codex が夜間のエクスポート障害をデバッグしている間に待機していた OpenAI インフラストラクチャ エンジニアである James Katz は、この進化の極端な例です。 GPT-5.5 のリリース後、Katz 氏のチームは実験として、承認されたユーザーに代わって Codex が OpenAI の内部データ エクスポート システムにクエリを実行できるスキルを作成しました。その際、ユーザーはリクエストを SQL で記述された正確にフォーマットされたデータベース命令に変換する方法を知る必要はありません。 (「なんてことだ、ほとんどの場合、これは最初のショットで実際に機能するのだ」と Katz 氏は思ったのを覚えています。) 数週間以内に、新しい Codex スキルがデータ エクスポート システムの総使用量の半分以上を占めました。カッツのチームは不意を突かれた。このツールの即座の人気は実証されましたが、同時に彼はすべてのエッジケースに対処し、誰も文書化していないシステム部分のガイダンスを作成する必要が生じました。
控えめに言っても、それはバランスをとる行為でした。カッツ氏は、自分のチームを「リソースが限られている」と丁寧に表現しています。エクスポート スキルに費やされる時間は、既存のインフラストラクチャ作業には費やされません。 「サポートすべき人員よりも多くの仕事が未処理にあります」と彼は言います。
Bonnie Xu、技術スタッフのメンバー。写真提供：サラ・デラゴン。
リアルタイム データ インフラストラクチャの責任者である Aravind Suresh 氏は次のように述べています。写真提供：サラ

デラゴン。
同様の動きが OpenAI 全体で展開されています。技術スタッフの Bonnie Xu とリアルタイム データ インフラストラクチャの責任者 Aravind Suresh は、OpenAI の内部データ分析エージェントを構築および保守する小規模なチームに所属しています。このエージェントは、同社の従業員の 5,000 人以上がデータ ウェアハウスのクエリに使用しています。
8 か月前、誰かがエージェントを変更したい場合は、Slack 経由で変更をリクエストしていました。 6 か月前、それが「PR をカットしてもいいですか?」に変わりました。 2 か月前、「Codex が作成した PR があるのですが、レビューしてもらえますか?」ということになりました。
現在、社内全体からプル リクエストが殺到しています。オコナー氏とカッツ氏と同様、シュー氏とスレッシュ氏のチームは、共通のコーデックス スキルを作成することで流入を管理しています。 PR がチームに到着するまでに、Codex はその作成者によるテスト、文書化、改良を支援しているはずです。これにより、本当に修復不可能なコード変更を除外するのに長い道のりを歩んできました。
PR の準備が完了すると、別の自動化された Codex ボットが PR をレビューし、最初のエージェントや作成者が見逃していたものを発見することがよくあると Xu 氏は言います。人間のレビュー担当者は今でもすべての PR を調査し、大規模な変更、リスクの高い変更、またはシステムの信頼性にとって重要な変更に特に注意を払っています。これらについて、チームは追加の Codex エージェントに特定の懸念事項を調査するよう依頼します。たとえば、コードの一部を書き換えると別のシステムが中断される可能性があるかどうか、変更がプロジェクトのガイドラインに従っているかどうかなどです。最大の変更では、十数のサブエージェントを使用してプロセス全体が再度実行される場合があります。
Codex はすべてを解決するものではありません。 Tang のチームはまた、ChatGPT Workspace エージェントを利用した Slack ボットを構築し、社内の同僚から寄せられる絶え間ない質問 (「このテーブルを更新するにはどのような権限が必要ですか?」など) を迂回させました。または、「このデータ パイプラインは失敗しました。誰かができますか?」

見て？"
しかし、このパッチワークのようなツールでも十分ではありません。 「解決策があると言う人は、おそらく全体像を知らないだけです」とタン氏は言う。
しかし、一つだけ明らかなことは、これから起こる大洪水に比べれば、今日の大洪水は霧雨のように見えるだろうということだ。
まだ何も見えていません」と OpenAI の応用インフラストラクチャ エンジニアリング担当バイスプレジデントである Venkat Venkataramani 氏は言います。
グリーンとグレーのチェック柄のフランネル ボタンダウン、スニーカー、ブルー ジーンズを着たヴェンカタラマーニのカジュアルな服装と穏やかな口調は、彼のユニークな視点とは裏腹にあります。 AI が生成したコードの氾濫の下で OpenAI のインフラストラクチャを実行し続ける責任者として、彼は会社が直面していること、そしてソフトウェア業界の他の部分が直面しようとしていることに、ほとんどの誰よりも明確な見解を持っています。
今日の AI 生成コードの量は前例のないものです。振り返ってみると、それは趣のあるものになるでしょう、とヴェンカタラマニは言います。近い将来、1 年分のインシデントが 1 週間に圧縮される可能性があります。 「鉄砲水が来ています」と彼は言います。
Venkat Venkataramani 副社長は、応用インフラストラクチャ エンジニアリング担当です。写真提供：サラ・デラゴン。
都市を襲う洪水とは異なり、危険は水の量だけではありません。より高品質なコードは、より有用な、または少なくとも無害なソフトウェアを意味します。危険なのは、洪水がもたらすものです。バグ、セキュリティ問題、その他の内部に隠された問題です。

[切り捨てられた]

## Original Extract

An exclusive look at how OpenAI’s infrastructure team is adapting to a flood of AI-generated code

Skip to content What Comes N ext
New ideas to help you build the future—in your inbox, every day.
Do Not Sell or Share My Personal Information
This site is protected by reCAPTCHA and the Google Privacy Policy and Terms of Service apply.
We use analytics and advertising tools by default. You can update this anytime.
Manage optional tracking categories. Necessary cookies stay on so the site can function.
Inside OpenAI’s Race to Reinvent Software Development for the Agent Era Sign in Subscribe
E logo
This is the promise of virtually every billboard in San Francisco: The agentic future has arrived, automating away the need for human engineers.
The reality, as is typically the case, is more complicated than the billboards would have you believe.
Sometime in the middle of the night, a coding agent running on a researcher’s laptop posted a Slack message about a failed data export.
The researcher was asleep. So was James Katz, the infrastructure engineer on call that evening.
For its part, the agent was unfazed by the late hour. It had access to the tools and internal systems it needed to debug the problem.
And so while it waited for a response, it got to work, diligently narrating its efforts to diagnose the incident, investigate potential fixes, and restrategize whenever it hit a roadblock.
By morning, it had handed the researcher the data they needed, without Katz or anyone else on his team having to step in.
This is the promise of virtually every billboard in San Francisco: The agentic future has arrived, automating away the need for human engineers.
The reality, as is typically the case, is more complicated than the billboards would have you believe.
Inside OpenAI’s race to redesign software development for a world where agents write all of the code
By Laura Entis · Photographs by Sarah Deragon and Sarah Jay Halliday
A middle-aged woman speaks into an intercom outside a sleek, nondescript office building in San Francisco’s Mission Bay district.
“For lunch, ma’am?” A tinny voice crackles through the speaker. A tall, bearded security guard intercedes at the door before the conversation can go any further. “This is not a restaurant.”
Indeed, it is not. At 1515 3rd Street, better known as OpenAI’s headquarters, you need an invitation just to enter the building. Once inside, you’re greeted in a lobby appointed like a slightly corporate five-star hotel, complete with accent books and lemon-wheeled water coolers.
OpenAI may not be a restaurant, but a few floors up, it is lunchtime. Employees all working toward the company’s stated mission of achieving Artificial General Intelligence (AGI) line up for harvest bowls, tacos, stir fry, and hot slices fresh out of what I’m told is San Francisco’s largest live-flame pizza oven.
As OpenAI competes with rival Anthropic to release new frontier models at an increasingly rapid clip, its engineers can ship code 10, even 100 times faster than before AI. Codex, the company’s coding agent and desktop app, extends that capacity beyond engineering; product managers, data scientists, researchers, and go-to-market employees use it to write, modify, debug, and submit code that runs on OpenAI’s internal systems. Many in the cafeteria are glued to their phones, perhaps texting, scrolling social media or, more likely, orchestrating agents through OpenAI’s mobile app.
It’s an exhilarating time for anyone building new products or tools. For the engineers responsible for maintaining the internal systems that make all this building possible, however, the feeling is closer to staying afloat.
W e use ‘deluge’ a lot,” says Emma Tang , who heads the applied infrastructure data platform team. Tang’s group maintains the infrastructure that moves, analyzes, and secures the data flowing through the plumbing of every OpenAI product, including ChatGPT and Codex. The team’s volume of pull requests (PRs), or proposed changes to a shared codebase, has jumped five- to tenfold in the past year.
A meaningful share of that code arrives from people in non-technical parts of the organization. “We have users generating Flink jobs”—programs built with Apache Flink, a system that analyzes continuous streams of product data—“who don’t even know Flink,” Tang says. As the pace of change has increased, so have internal reliability problems, from delayed data processing to dashboards that fail to load. And when something breaks, the author often can’t help because they don’t fully understand what they’ve shipped.
Until recently, if someone in operations or on the go-to-market team wanted a custom internal app, they’d need to find an engineer to make it for them. Now, anyone at the company can “show up on Monday with a working prototype they built over the weekend,” Tang says. Her engineers act increasingly as educators and mediators, helping colleagues turn their AI-generated tools into reliable software. A vibe-coded app, for example, might unnecessarily refresh every few seconds, while the analytics system it pulls from updates only once a day. Making the data truly live means connecting it to a different production system, which raises questions largely invisible to people without a technical background—questions like, “Where will it run?,” “How should it retrieve information without slowing down the service everyone else relies on?,” or “Who will maintain it?”
The last question, in particular, often falls on Tang’s team. AI allows anyone at the company to build apps and systems that depend on the company’s shared data infrastructure, but should that code fail, her team inherits the problem.
“We become the catch-all,” Tang says. “Nobody owns it anymore except us.”
Emma Tang, head of data infrastructure engineering. Photo courtesy Sarah Deragon.
James Katz, member of technical staff. Photo courtesy of Sarah Jay Halliday.
Molly O’Connor, member of technical staff. Photo courtesy Sarah Jay Halliday.
T he applied infrastructure team’s first response to the deluge is also the tool that created it: Codex.
Within weeks of the release of GPT-5.5 , adoption of the Codex desktop app within OpenAI reached nearly 100 percent. Engineers now run the bulk of their work from inside the app, reviewing AI-generated code, AI-generated tests (automated checks that confirm the code works as it should), and AI-proposed fixes. Because OpenAI employees have no token limits, an engineer can queue up a batch of UI bugs in Codex and have multiple agents independently write fixes, run them in test versions of the product, click through the affected flows in a browser, and flag resulting issues—all using the most powerful models (many employees are granted access to models before they’re available to the public), without thinking twice about cost or efficiency.
When Molly O’Connor joined OpenAI’s applied data platform team in October 2024, a typical day included getting paged for an incident, debugging it, and then moving on. Now, along with making the fix itself, she might also update the company’s shared Codex instructions so that it can correct similar problems in the future. O’Connor does the same when she spots recurring problems in pull requests—especially ones that could make the system less reliable.
Her role—to make sure OpenAI’s data infrastructure remains reliable and secure—hasn’t changed. But more of the work now involves turning what she knows about building and troubleshooting data pipelines, executing large-scale migrations, and testing and rolling out updates to data-processing systems into instructions Codex can follow.
James Katz , the OpenAI infrastructure engineer on call while Codex debugged the overnight export failure, is an extreme example of this evolution. Following the release of GPT-5.5, Katz’s team created, as an experiment, a skill that lets Codex query OpenAI’s internal data export system on behalf of approved users, without requiring them to know how to translate their requests into precisely formatted database instructions written in SQL. (“Holy crap, this is really working on the first shot most of the time,” Katz remembers thinking.) Within a couple of weeks, the new Codex skill accounted for more than half of the data export system’s total usage. Katz’s team was caught off guard; the tool’s immediate popularity was validating, but it also left him to deal with all of the edge cases and write guidance for the parts of the system no one had documented.
It’s been a balancing act, to say the least. Katz politely describes his team as “resource constrained”; any time devoted to the export skill is time not spent on existing infrastructure work. “There’s always more work on the backlog than we have people to support,” he says.
Bonnie Xu, member of technical staff. Photo courtesy of Sarah Deragon.
Aravind Suresh, head of realtime data infrastructure. Photo courtesy of Sarah Deragon.
A similar dynamic is playing out throughout OpenAI. Technical staff member Bonnie Xu and head of realtime data infrastructure Aravind Suresh are on a small team that built and maintains OpenAI’s internal data analytics agent, which is used by more than 5,000 of the company’s employees to query its data warehouse.
Eight months ago, if someone wanted to modify the agent, they’d request the change via Slack. Six months ago, that morphed into “Hey, can I cut a PR?” Two months ago, it became, “I have a PR that Codex created for me. Can you review it?”
Pull requests now pour in from across the company. Like O’Connor and Katz, Xu and Suresh’s team manage the influx by creating shared Codex skills. By the time a PR reaches their team, Codex should have helped its author test, document, and refine it—which has gone a long way to filtering out the truly unsalvageable code changes.
Once a PR is ready, a separate, automated Codex bot reviews it, often catching things the first agent or author missed, Xu says. Human reviewers still look at every PR, with particular attention to changes that are large, risky, or crucial to the system’s reliability. For those, the team asks additional Codex agents to investigate specific concerns—for example, whether rewriting part of the code could disrupt another system or whether the change follows a project’s guidelines. The largest changes may go through the entire process again with more than a dozen subagents.
Codex isn’t a fix for everything. Tang’s team also built a Slack bot powered by a ChatGPT Workspace agent to divert the constant stream of questions from colleagues across the company—questions like “What permission do I need to update this table?” or “This data pipeline failed, can someone look?”
But even this patchwork of tools isn’t enough. “If anybody says they have a solution, they probably just don’t know the whole picture,” Tang says.
One thing is clear, however: Compared to what’s coming, today’s deluge will look like a drizzle.
W e ain’t seen nothing yet,” says Venkat Venkataramani , OpenAI’s vice president of applied infrastructure engineering.
Dressed in a green-and-gray plaid flannel button-down, sneakers, and blue jeans, Venkataramani’s casual clothes and soft-spoken cadence belie his unique vantage point. As the person responsible for keeping OpenAI’s infrastructure running under the inundation of AI-generated code, he has a clearer view than almost anyone of what the company is facing—and what the rest of the software industry is about to face.
Today’s volume of AI generated code is unprecedented; in retrospect, says Venkataramani, it will look quaint. In the near future, a year’s worth of incidents could be compressed into a single week. “A flash flood is coming,” he says.
Venkat Venkataramani, vice president, applied infrastructure engineering. Photo courtesy of Sarah Deragon.
Unlike a flood that hits a city, the danger isn’t in the volume of the water alone. More high-quality code just means more useful—or at the very least, harmless—software. What’s dangerous is what the flood carries with it: bugs, security issues, and other problems hidden within t

[truncated]
