---
source: "https://schipper.ai/posts/technical-leaders-should-have-the-largest-ai-exhaust/"
hn_url: "https://news.ycombinator.com/item?id=49368389"
title: "Technical leaders should have the largest AI exhaust"
article_title: "Manuel Schipper"
image: "https://schipper.ai/og-card.png"
author: "schipperai"
captured_at: "2026-08-19T23:14:07Z"
capture_tool: "hn-digest"
hn_id: 49368389
score: 1
comments: 0
posted_at: "2026-08-19T23:12:38Z"
tags:
  - hacker-news
  - translated
---

# Technical leaders should have the largest AI exhaust

- HN: [49368389](https://news.ycombinator.com/item?id=49368389)
- Source: [schipper.ai](https://schipper.ai/posts/technical-leaders-should-have-the-largest-ai-exhaust/)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T23:12:38Z

## Translation

タイトル: 技術リーダーは AI を最大限に活用する必要がある
記事のタイトル: マヌエル・シッパー
説明: 私は現在、Snowflake のスタッフ AI エンジニアとして GTM チームのエージェントに取り組んでおり、nah の作者です。このサイトには、エージェントと開発者ツールに関する私の公開作品が集められています。

記事本文:
マヌエル・シッパー
について
仕事
明るい/暗い
≡
内容
エージェントを使用してコーディングする正しい方法を誰も知りません
年功序列はより多くの排気ガスを生み出すべきである
AIの排出は貢献ではない
技術リーダーは AI を最大限に活用する必要がある
AI の排出量 (トークンの書き込みやコード行) によってエンジニアを評価するのは悪い考えです。生成されたコードと消費されたトークンは、影響と同じものではありません。スタッフ エンジニアは、コーディング以外の活動 (ビジョン ドキュメントの仕上げ、複雑なプル リクエストのレビュー、他のチームへの影響、他のチームの指導など) にほとんどの時間を費やしても、2 倍の PR を作成したエンジニアよりも大きな影響力を与えることができます。
アンディ・グローブ氏は、「高出力管理」の中で、これをレバレッジと表現しました。リーダーの成果は、リーダーが生み出す仕事を超えて、組織全体で実現できることに直接広がります。グローブ氏は、トレーニングはリーダーが実行できる最も大きな力を発揮する活動の 1 つであると述べています。
エンジニアの上級職になるほど、同僚からのフィードバックや、自分が主導したり貢献した仕事の影響力に基づいて評価が重くなり、影響力のある PR の数に基づく評価は低くなります。歴史的には、これにより年功序列と直接的な技術成果の間には逆の関係が生まれました。エンジニアが上級になるにつれて、コーディング以外の活用が期待されるようになりました。
しかし、AI は私たちの実践を変えました。これはもはや当てはまらないと思います。
エージェントを使用してコーディングする正しい方法を誰も知りません
長い間、ソフトウェア エンジニアリングの実践は十分に安定していたので、上級エンジニアはソフトウェアの作り方を完全に失うことなく、主にコーディングから離れることができました。確かに、言​​語、インフラ、フレームワークは時々変わりましたが、基本原則は変わりませんでした。人間とソフトウェアの間の主なインターフェイスは、IDE と端末でした。エンジニアデス

システムに署名し、コードを作成し、プル リクエストをレビューし、テストし、出荷し、CI/CD を構築しました。
しかし現在、私たちは人間の意図とソフトウェアの間のインターフェースについて積極的に再訴訟を起こす時期に入っています。コーディング エージェントを使用して構築するための確立された方法はなく、技術リーダーには新しい制作手段の直接の経験が必要です。
以下は、私が個人的に取り組んでいる未解決の質問のリストです。
エンジニアはコードをすべて読む必要がありますか?
主にテストとその結果の動作をレビューする必要がありますか?
コード変更の説明を AI に頼っても大丈夫でしょうか?
AI コード レビューは 1 つで十分ですか? それとも複数のゲートが必要ですか?
効果的なコンテキスト ウィンドウとは何でしょうか?
エージェントが独自にスキルを呼び出すことを許可されるべきですか?
LLM はコードベースの Wiki を保守する必要がありますか?
Markdown プランはリポジトリに保存する必要がありますか?
コードベースのキュレーションと編成は依然として重要ですか?
エージェントを副操縦士として協力すべきでしょうか、それとも監督なしで働くべきでしょうか?
スペック駆動開発を採用すべきでしょうか?
エージェントは独自の ID を持つ必要がありますか、それともユーザーの代わりに行動する必要がありますか?
エージェントはコーディングのみに集中すべきでしょうか、それとも GitHub や Jira などの外部サービスに接続すべきでしょうか?
一方で、新しいツールは毎日登場します。私が試してきたいくつかのツール カテゴリを次に示します (これもすべてではありません)。
実稼働ソフトウェアを構築する標準的な方法については、まだ合意に近づいているとは思えません。これらの質問は、工学原理についての推論では答えることができません。それらは経験的な質問です。
実際のコードベースに対してエージェントを実行し、エージェントを失敗に追い込む方法を見つける必要があります。好みのスタックを立ち上げ、独自のカスタム ツールを構築し、すべてを捨てて 3 か月後に再び開始する準備をしなければなりません。そうでなければ、知的な立場を築くことはできません。中古のレポート、デモ、または記事は、

十分ではありません。
年功序列はより多くの排気ガスを生み出すべきである
エンジニアはコーディング エージェントを使用してタスクを完了します。技術リーダーは現在、コーディング エージェントを共同で活用する効果的な方法を見つけ出すという、より広範な任務を負っています。
エージェントはどこでうまく機能しますか?どこで失敗するのでしょうか？彼らはどのようなコンテキストを必要としているのでしょうか?人間は何を見直すべきでしょうか？どのコントロールを決定論的にする必要がありますか?チーム全体で何を標準化し、何をエンジニア個人に任せるべきなのか。
これらの決定に影響を与えるには、多くの直接の経験が必要です。スタッフと主任エンジニアは、実験を実行し、新しいツールを試し、実際の問題に対してモデルを最大限にプッシュする必要があります。そうなると、焼かれたトークン、コード行、失敗したプロトタイプ、放棄されたブランチなど、AI による排出のかなりの痕跡が必然的に残るはずです。
実験を通じて私が到達したいくつかの結論は次のとおりです。
コンテキスト ウィンドウの場合: Fable や Sol などの最も有能なモデルでも、依然として大量の間違いを犯し、指示を無視します。これは、コンテキスト ウィンドウが長い場合に特に顕著です。したがって、エージェントには効果的なコンテキスト ウィンドウがあると私は結論付けました。それが何なのかはまだわかりませんが、これは厳密な数字ではなく、タスクに大きく依存しているのではないかと思います。それにもかかわらず、私は現在コンテキスト ウィンドウの上限を 50% に設定しており、20% を超えると警告がトリガーされます。
自律性について: 最近のモデルは非常にエージェント的です。彼らは PR をファイルし、メインにプッシュし、要求されていない機能を構築します。これは、エージェントの指示があいまいな場合に特に当てはまります。私は最近スタック全体を削除してゼロから始めたときにこれを発見しました。私はすべてのスキルと AGENTS.md をクリアし、仕様駆動開発 (SDD) フレームワークの使用をやめました。私は現在、指示ガードレールと決定論的なフックとワークフローを使用してハーネスを再構築しています。明確に定義された

事前にタスクを実行することで、予期せぬ事態が最小限に抑えられ、エージェントはより予測しやすくなります。
コードレビューについて: 人間がコードをレビューすることには依然として多くの価値があります。仕様や指示に決定事項が明記されていない場合、エージェントは空白を埋めることがよくあり、その結果、過剰に設計されたソリューション、要求されていないテスト、またはまったくひどいコードが作成されることがよくあります。また、質問しない限り教えてくれないこともよくあります。そもそも仕様書に記載するのを忘れていたため、質問することができないかもしれません。何かが重要な場合は、PR を送信する前にコードをスキャンすることが依然として絶対に必要です。人間とエージェントの理解においては多くの革新が起こっており、私はこの分野を注意深く観察しています。
コードベースの保守について: エージェントは grep を使用してコードベースのコンテキストを構築します。コードベースが grepable ではない場合、エージェントはタスクに必要なすべてのコンテキストを読み込むのに苦労します。ドメインの単語を知っているエージェントが、ファイル全体を読まなくても検索によって概念とその配線を見つけることができる場合、コードベースは grepable です。私はよくエージェントに、「名前はその内容を意味する」、「副作用には明らかな所有者がいる」、「ディレクトリには単一の国家責任がある」などの原則に基づいてコードベースを監査してもらいます。 Markdown ファイル ( spec.md など) は、最終的に grep で検出され、不要なコンテキストをロードしてしまうため、リポジトリには保存しません。
AI レビューについて: エージェントは仕様を実装する最初の段階でバグを導入します。独立した QA エージェントによるゲート実装作業は、バグを表面化するのに効果的です。実装者のモデルとは異なるモデル ファミリの QA エージェントは、より良い結果を生み出す傾向があります。しかし、QA エージェントは細かいことを指摘する傾向があるため、修正される内容については依然として人間によるゲートが必要です。そうしないと、エージェントが過度に批判的になり、「多層防御」を導入してしまうからです。

必要ありません。
これらの結論の多くは、私自身の得意なプロジェクトを実行することで収集したことに注意してください。私の専門的な仕事では、私が所有しているものの多くは重要なものであるため、エージェントをしっかりと縛り付けておきます。
AIの排出は貢献ではない
トークンの消費量が多いこと自体は何の証明にもならず、パフォーマンスの目標となるべきではありません。私が言いたいのは、真剣な実験では目に見える排気が生じるということです。フロンティア モデルをテストし、ハーネスを構築し、ワークフローを比較し、エージェントの失敗を調査し、エージェントに困難な問題を突きつけている場合、トークンを消費してコードを生成することになります。 AI の排出量が多いからといって、その人が技術的なリーダーであることを証明するものではありませんが、技術的な方向性を設定する人による AI の排出量が少ないと、疑問が生じるはずです。
以前は、基礎となる慣行がよく理解されていたため、上級エンジニアは組織内での影響力を構築する方法として委任に頼ることができました。今日の世界では、実践そのものがデザインされています。したがって、実験を委任するということは、自分自身の判断のソースを委任することを意味します。新たな間違いは、年功序列がソフトウェアの製造方法に直接触れる必要がなくなったことを意味すると仮定することです。おそらく私たちの練習は安定するでしょう、そしてそれは再び真実になるかもしれません。

## Original Extract

I'm currently a Staff AI Engineer at Snowflake working on agents for GTM teams, and the author of nah. This site collects my public work on agents and developer tooling.

Manuel Schipper
about
work
light/dark
≡
Contents
Nobody knows the right way to code with agents
Seniority should produce more exhaust
The AI exhaust is not the contribution
Technical leaders should have the largest AI exhaust
Measuring engineers by their AI exhaust (token burn and lines of code) is a bad idea. Code generated and tokens spent are not the same thing as impact. A staff engineer can spend most of their time on non-coding activities (polishing a vision document, reviewing complex pull requests, influencing other teams, mentoring others) and still have more impact than an engineer who authored twice as many PRs.
In High Output Management , Andy Grove described this as leverage. The output of a leader extends beyond the work they produce directly into what they enable across the organization. Grove calls training one of the highest-leverage activities a leader can perform.
The more senior an engineer is, they are evaluated more heavily based on peer feedback and the impact of the work they led or contributed to, and less so based on the number of impactful PRs. Historically, this created an inverse relationship between seniority and direct technical output. As engineers became more senior, they were expected to find leverage other than coding.
But AI has changed our practice, and I think this is no longer the case.
Nobody knows the right way to code with agents
For a long time, the practice of software engineering was stable enough that a senior engineer could move away from primarily coding without completely losing touch with how software was made. Sure, languages, infra, and frameworks changed every now and then, but the basic principles remained true. The primary interface between humans and software was an IDE and a terminal. Engineers designed systems, wrote code, reviewed pull requests, tested, shipped, and built CI/CD.
Yet now, we have entered a period of actively re-litigating the interface between human intent and software. There is no settled way to build with coding agents, and technical leaders need firsthand experience with the new means of production.
Here’s a non-exhaustive list of open questions I am personally wrestling with:
Should engineers read all of the code?
Should we mostly review the tests and resulting behavior?
Is it okay to rely on AI to explain the code changes?
Is one AI code review sufficient, or do we need multiple gates?
What’s an effective context window?
Should agents be allowed to invoke skills on their own?
Should an LLM maintain a wiki of the codebase?
Should Markdown plans be stored in the repo?
Does codebase curation and organization still matter?
Should we work with agents as copilots, or should they work unsupervised?
Should we adopt spec-driven development?
Should agents have their own identity or act on behalf of the user?
Should agents focus solely on coding or be connected to external services like GitHub and Jira?
Meanwhile, new tools pop up every day. Here are some tool categories (again, non-exhaustive) that I’ve been experimenting with:
I don’t think we are close to agreeing on a standard way to build production software. These questions cannot be answered by reasoning about engineering principles. They are empirical questions.
You have to run agents against real codebases and find ways to push them into failing. You have to stand up your preferred stack, build your own custom tooling, and be ready to throw it all away and start again three months later. Otherwise, you cannot develop an intelligent position. Secondhand reports, demos, or articles will not suffice.
Seniority should produce more exhaust
An engineer uses coding agents to complete a task. A technical leader now has a broader mandate: figuring out effective ways to leverage coding agents collectively.
Where do agents work well? Where do they fail? What context do they need? What should humans review? Which controls should be deterministic? What should be standardized across the team, and what should be left to individual engineers?
To influence these decisions, you need a lot of firsthand experience. Staff and principal engineers should be running experiments, trying new tools, and pushing models the hardest against real problems. That should inevitably leave a sizable trace of AI exhaust: burned tokens, lines of code, failed prototypes, and abandoned branches.
Here are some conclusions I’ve reached through experimentation:
On context windows: Even the most capable models, such as Fable and Sol, still make a ton of mistakes and ignore instructions. This is especially apparent at longer context windows. Thus, I have concluded that agents have an effective context window. I am still not sure what that is, and I suspect it is not a hard number and instead depends a lot on the task. Nonetheless, I am now capping my context windows at 50%, with a warning triggered after I pass 20%.
On autonomy: Models are extremely agentic these days. They will file a PR, push to main, and build features you didn’t ask for. This is especially true when agent instructions are vague. I discovered this when I recently deleted my whole stack to start from zero. I cleared all my skills and my AGENTS.md , and I stopped using my spec-driven development (SDD) framework. I am now rebuilding my harness with instruction guardrails and deterministic hooks and workflows. Well-defined tasks upfront minimize surprises and keep agents more predictable.
On code review: There’s still a lot of value in humans reviewing code. Agents will often fill in the blank when the spec or instructions do not specify a decision, which often results in over-engineered solutions, tests you didn’t ask for, or just outright terrible code. They also often will not tell you about it unless you ask. You might never ask because you forgot to put it in the spec in the first place. If something is critical, scanning the code before shipping a PR is absolutely still necessary. There’s a ton of innovation happening in human-agent understanding, and I’m watching this space closely.
On maintaining codebases: Agents use grep to build context for a codebase. If your codebase is not greppable, agents will struggle to load all the necessary context for the task. A codebase is greppable when an agent that knows the domain words can find a concept and its wiring by searching without having to read whole files. I often have agents audit my codebase against principles like “names mean what they say,” “side effects have obvious owners,” and “directories have one state responsibility.” I don’t store Markdown files (like spec.md ) in my repo because they end up being discovered with grep and then load unwanted context.
On AI reviews: Agents will introduce bugs in their first pass at implementing a spec. Having an independent QA agent gate implementation work is effective at surfacing bugs. A QA agent from a different model family than the implementer’s model tends to produce better results. But QA agents tend to nitpick, so a human gate on what gets fixed is still needed because otherwise agents will be overly critical and introduce “defense in depth” that you don’t need.
Note that I gathered a lot of these conclusions by doing my own pet projects. In my professional work, I keep agents on a tight leash because much of what I own is critical.
The AI exhaust is not the contribution
High token consumption in itself proves nothing and should never be a performance target. The point I’m making is that serious experimentation produces visible exhaust. If you are testing frontier models, building harnesses, comparing workflows, studying agent failures, and pushing agents against difficult problems, you will consume tokens and generate code. High AI exhaust does not prove that someone is a technical leader, but low AI exhaust from those setting technical direction should raise questions.
In the past, a senior engineer could rely on delegation as a way to build influence in an organization because the underlying practice was well understood. In today’s world, the practice itself is the thing being designed. Thus, delegating experimentation means delegating the source of your own judgment. A new mistake would be assuming that seniority means you no longer need as much firsthand contact with how software is made. Perhaps our practice will stabilize, and that might be true again.
