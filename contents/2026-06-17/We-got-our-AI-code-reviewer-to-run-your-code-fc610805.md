---
source: "https://www.greptile.com/blog/trex-code-execution"
hn_url: "https://news.ycombinator.com/item?id=48571851"
title: "We got our AI code reviewer to run your code"
article_title: "Building TREX: Code Execution and Artifact Generation for AI Code Review | Greptile"
author: "dakshgupta"
captured_at: "2026-06-17T16:21:28Z"
capture_tool: "hn-digest"
hn_id: 48571851
score: 2
comments: 0
posted_at: "2026-06-17T15:29:01Z"
tags:
  - hacker-news
  - translated
---

# We got our AI code reviewer to run your code

- HN: [48571851](https://news.ycombinator.com/item?id=48571851)
- Source: [www.greptile.com](https://www.greptile.com/blog/trex-code-execution)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T15:29:01Z

## Translation

タイトル: AI コードレビューアにコードを実行してもらいました
記事のタイトル: TREX の構築: AI コードレビューのためのコード実行とアーティファクト生成 |グレプタイル
説明: Greptile が、サンドボックス環境でコードを実行し、ランタイム バグを捕捉するマルチモーダル アーティファクトを生成する AI コード レビュー用の実行レイヤーである TREX をどのように構築したか。

記事本文:
TREX の構築: AI コードレビューのためのコード実行とアーティファクト生成 | Greptile TREX の紹介: Greptile がコードを実行できるようになりました。 TREX がコードを実行します。
例 価格設定 機能 エンタープライズ ブログ リソース 営業担当へのお問い合わせ サインアップ サインアップ TREX: テスト生成 パブリック ベータ版 Greptile エージェントの独立性 学習とカスタム コンテキスト Greptile CLI 統合 ドキュメント 包括的なドキュメントとガイド Greptile の例 「Greptile が大規模な OS リポジトリでバグをキャッチする」を参照してください。営業へのお問い合わせ 当社チームとのデモのスケジュール お客様 成功事例とケーススタディ ベンチマーク パフォーマンス指標と比較 セキュリティ セキュリティ慣行とコンプライアンス TREX の構築: AI コードレビューのためのコード実行とアーティファクト生成
[ シュロック・メロトラ | 2026-06-17 ]
コンテキストを無駄にすることなくエージェントを調整する
作品を見せるためのマルチモーダルな成果物のデザイン
モデルに依存しない評価ハーネスの構築
ファーストクラスのレビュープリミティブとしてのコード実行
ブログ構築 TREX: コードの実行と成果物... 目次 目次
コンテキストを無駄にすることなくエージェントを調整する
作品を見せるためのマルチモーダルな成果物のデザイン
モデルに依存しない評価ハーネスの構築
ファーストクラスのレビュープリミティブとしてのコード実行
私は Shlok、Greptile のソフトウェア エンジニアです。私たちは最近、プル リクエストのレビューに加えて、実際にコードを実行して何が問題だったかを示すコード レビューアーを構築しました。
1976 年、Michael Fagan は IBM での正式なコード検査を紹介する論文を発表しました。開発者はリストを印刷し、同じ部屋に座り、コードを一行ずつ読んでいきました。
今でも画面上で差分を読みます。 AI ツールの登場によりこれが高速化されましたが、そのほとんどは依然としてコードを読み取るだけです。このアプローチは、コード内で明確に通知されるバグなど、多くのバグに対して機能します。
問題

つまり、コードにはまったく現れないバグのカテゴリ全体が存在するということです。これらはプログラムの実行中に存在します。特定の状態シーケンスを必要とするロジック エラー、ページの読み込み後に発生する UI 回帰、または実際のリクエストを必要とする競合状態を考えてください。 diff を完全に読み取ることができても、この種のバグを完全に見逃すことはできます。
静的コードレビューには上限があります。コードの内容を推論することができます。それが何をするのかは説明できません。 TREX (「Test、Run、Execute」の略) は、その上限に対する Greptile の対応策であり、コード レビューに直接組み込まれた実行層です。
コンテキストを無駄にすることなくエージェントを調整する
TREX は、テストを生成して実行するスタンドアロン エージェントとして、Greptile とは完全に別の製品としてスタートしました。その結果、バグが表面化することを期待していました。彼らはそうしませんでした。テストの生成は、バグの発見と同じ作業ではありませんでした。別の TREX エージェントがテストを作成しようとしたとき、そのテストはユーザが行おうとしていたものとは無関係でした。これにより、不要なノイズが発生し、エッジケースも見逃されました。今にして思えば当然のことのように思えますが、この教訓を学ぶには予想以上に時間がかかりました。
私たちは、各エージェントに独自のコンテキスト ウィンドウを与えることを前提として、これらのエージェントを別個に構築しました。また、両方のエージェントが知識を共有せずに別々に実行されることも意味しました。これらは重複することが多く、どちらのエージェントももう一方のエージェントがすでに見つけたことを知らないまま、コードベースの同じ部分を 2 回探索することになり、最終的には無駄なコンピューティングが発生しました。
明らかな修正は、それらを 1 つのエージェントに統合することのように思えました。私たちはそれを試みましたが、別の問題に遭遇しました。完全なレビューを処理する単一のエージェントが過負荷になってしまいました。サービスの起動、スクリーンショットの取得、テストの実行の間には、1 人のエージェントがきれいに管理するにはコンテキストが多すぎました。
解決策

TREX を完全に別の製品として存在させるのではなく、メインの Greptile レビュアーと同じコンテキストを共有させることでした。エージェント内からエージェントを管理するのは初めてでした。 2 つの独立したエージェントとは異なり、これは TREX がゼロから開始するわけではないことを意味します。これは、Greptile レビュー担当エージェントがすでに見つけたものを継承し、独自のコンテキスト ウィンドウを持ち、調査を依頼された特定の問題に範囲を限定します。
Greptile レビューアー エージェントはオーケストレーターとして機能します。差分を読み取り、調査する価値のある問題を特定し、問題ごとに専用の TREX エージェントを起動し、すべてが並行して実行されます。 TREX エージェントは、オーケストレータ エージェントの自由、計算能力、および知識を持っています。
この好例は、認証ゲートの背後に隠された UI 機能です。ローカルでテストするということは、環境をセットアップし、認証を処理し、機能フラグを正しい状態に取得することを意味します。サブエージェントはすべてを独自に把握し、レンダリングされたフィーチャのスクリーンショットを返します。
作品を見せるためのマルチモーダルな成果物のデザイン
TREX の最初のバージョンでは、テストされた内容と何が起こったかを箇条書きとして結果が出力されました。これは適切な出発点でしたが、十分な情報が得られませんでした。
エージェントや人間のレビュー担当者が「チェックアウト フローをテストしたところ、失敗が見つかりました」のような箇条書きを読んでも、あまり役に立ちません。彼らは、プロセスのどこで何か問題が起こったのかを知ることができません。テストが失敗した場合、それはセットアップにありましたか?主張は？環境の問題？私たちは、このエージェントの初期バージョンが、何かをどれほど徹底的にテストしたかについて時々幻覚を見せ、試していないものを試したと主張することが判明しました。箇条書きでは確認する方法がありませんでした。
修正は、各 TRE に設定されたマルチモーダル アーティファクトを使用して箇条書きリストをバックアップすることでした。

X の検出: スクリーンショット、ログ、API トレース、実行スクリプト。各モダリティはストーリーの異なる部分をカバーします。実際に重要なのは、特定の問題に対してテストされたすべての内容の包括的な全体像を把握することです。
私たちに「すごい」と言わせた最初の成果物はビデオでした。アニメーションの変更をプッシュすると、TREX はその再生のビデオをキャプチャします。ローカル環境を開かなくても、アニメーションがどのように見えるかを正確に確認できます。
アーティファクトも信頼できるものである必要があります。すべてのアーティファクトは、レビュー担当者自身が実行を検証するのに十分な情報を提供する必要があります。スクリーンショット、ログ、トレース、スクリプトはすべてそこにあるため、担当者または下流エージェントは何が起こったのかを正確に見て確認できます。悪い証拠は、証拠がないより悪いです。
特に下流のエージェントにとって成果物が重要である理由は、教師が生徒に自分の作品を見せるよう要求する理由と同じです。これは小学校の算数に似ています。手順を示すまで、答えがどこで間違っていたかわかりません。エージェントも同様です。成果物が、結論に達するためにテストが行​​ったすべてのことの証拠を示している場合、エージェントはどのステップが間違っていたのかを正確に特定できます。その痕跡がなければ、答えだけがあり、その答えはどこを修正すればよいかについては何も教えてくれません。
TREX がバグを発見した場合、それは PR にコメントされます。機能が実行され、すべてが機能した場合は、変更が実際にテストされた証拠として概要に記載されます。すべての実行で問題を見つける必要があるわけではありません。
モデルに依存しない評価ハーネスの構築
モデルプロバイダー間のフロンティア競争は急速に進んでいます。ある月にコードタスクでリードしていたモデルが、次の月には遅れる可能性があります。単一プロバイダーの API を中心に緊密に構築するということは、ランキングが変化したときに再構築することを意味します。それは実行可能な長期戦略ではありません。
私たちは最初から、モデルに依存しないハーネスを中心に TREX を設計しました。

再構築せずにフロンティア モデル間のホットスワップが可能です。柔軟性はほとんどの人が予想するよりも深く、メイン エージェントとサブエージェントは異なるプロバイダーを使用できます。同じレビュー内で複数のモデルを実行できます。これにより、内部評価に基づいて、任意の時点で最適なモデルを簡単に選択できるようになります。
現在の評価には、再現率 (実際のバグがどれだけ捕捉されるか、オープンソースの PR やコメントが対処された顧客データに対して測定されるかなど) と精度 (実行間の一貫性: 同じ PR を 2 回レビューした場合、ほぼ同じ一連の問題が見つかるかなど) の測定が含まれます。
評価ではレイテンシを意図的に優先順位を下げています。レビューを待っている開発者は、信頼できない回答をすぐに得るよりも、もう少し待って正確な結果を得るほうを好みます。
私たちが使用するオープンソース評価ハーネスは、ネイティブ プロバイダー ハーネスと同等のパフォーマンスを発揮します。モデルに依存しないことによる重大な品質上のペナルティはありません。テストする前に推測するように頼まれても、私は自信を持てなかったでしょう。
TREX の差別化は、どのモデルが実行されているかではありません。これは、コードベースのインデックス作成、オーケストレーション、アーティファクトの生成、評価フレームワークなど、モデル周辺のインフラストラクチャです。知能は継続的に改善されるものであるため、差別化要因として知能を気にする必要はありません。 Greptile の仕事は、実際に違いをもたらすハーネス、アーキテクチャ、アーティファクト パイプラインを構築することです。
ファーストクラスのレビュープリミティブとしてのコード実行
TREX レビューごとに、使い捨てのサンドボックス環境が起動されます。つまり、レビューごとに分離されたコンピューティング インスタンスがミリ秒単位で新たに開始され、実行が完了すると破棄されます。これらの環境は起動が速く、実行が適切に分離され、実行可能です。

実際のプロジェクト。モックに対する単体テストだけでなく、実際の依存関係を持つ実際のサービスもテストします。多くのバグはエンドツーエンドでのみ発生します。
毎回何もないところから始めるのでは遅すぎます。そのため、再利用可能な基本イメージとリポジトリごとのスナップショットに依存しています。リポジトリは一度複製し、キャプチャして再開できます。各レビューでは、実行が開始される前に正確な PR コミットをフェッチし、認証情報をローテーションします。キャッシュにより、古い状態が環境にフリーズされることなく、セットアップが高速化されます。キャッシュに含まれる内容が少なすぎると速度が遅くなります。キャッシュに含まれる内容が多すぎると、取り憑かれてしまいます。 TREX には、素早く動けるほど暖かく、信頼できるほど新鮮な、便利な種類が必要です。
サンドボックスはアーティファクトの信頼性を高めるものです。 TREX が壊れたレイアウトでレンダリングされたページを報告すると、そのコードが実際の環境で実際に実行されたことがわかります。これは、「コードが壊れたレイアウトでレンダリングされる可能性がある」という主張とは異なります。
これらの部分 (サブエージェント アーキテクチャ、アーティファクト検証バー、サンドボックス実行環境、評価ハーネス) は、別個の機能ではありません。それらは 1 つのシステムです。オーケストレーターは、実行する価値のある問題を特定します。サンドボックスを使用すると、安全かつ高速に実行できます。アーティファクト パイプラインにより、結果に基づいて行動できるほど信頼できるものになります。評価ハーネスは、適切なモデルが機能していることを確認します。それぞれが他のものの価値を高めます。作成されたレビューは、証拠を添付した再現可能な実験です。
私たちのビジョンは、現在のコードレビューよりも大きいです。目指すはバグのない世界。そこに到達するために、私たちはもはや自分たちをコードレビューツールとは考えていません。私たちは検証スイート、つまりエンジニアリング チームが何十年にもわたって行ってきたことを模倣しながら、自動化され、すべての PR で実行されるエンドツーエンドのレイヤーになりたいと考えています。 TREX はその vi に向けた新たな一歩です

シオン。
製品 Greptile 2.0 の紹介
コードベースを理解する AI プラットフォームである Greptile 2.0 をご覧ください。スマートなコード レビュー、JIRA 統合、カスタム ワークフローで開発を促進します。
プログラミング openclaw/openclaw で開かれた PR の統計的研究
OpenClaw は、ほぼ一夜にして GitHub 史上最も急成長するリポジトリになりました。 PR データは、オープンソース コントリビューションの将来がどうなるかをプレビューするものです。
プログラミング コードベースは意味的に検索するのが独特に難しい
コードベースでのセマンティック検索がなぜ難しいのか、そして自然言語処理と詳細なチャンキングが AI コードの理解をどのように向上させるのかを探ります。

## Original Extract

How Greptile built TREX, an execution layer for AI code review that runs code in sandboxed environments and generates multi-modal artifacts to catch runtime bugs.

Building TREX: Code Execution and Artifact Generation for AI Code Review | Greptile Introducing TREX: Greptile Now Runs Your Code. TREX Runs Your Code.
Examples Pricing Features Enterprise Blog Resources Contact Sales Sign up Sign up TREX: Test Generation Public Beta Greptile Agent Independence Learning & Custom Context Greptile CLI Integrations Docs Comprehensive documentation and guides Greptile Examples See Greptile catching bugs in large OS repos. Contact Sales Schedule a demo with our team Customers Success stories and case studies Benchmarks Performance metrics and comparisons Security Security practices and compliance Building TREX: Code execution and artifact generation for AI code review
[ Shlok Mehrotra | 2026-06-17 ]
Orchestrating agents without wasting context
Designing multi-modal artifacts to show the work
Building a model-agnostic evaluation harness
Code execution as a first-class review primitive
Blog Building TREX: Code execution and artifa... Table of Contents Table of Contents
Orchestrating agents without wasting context
Designing multi-modal artifacts to show the work
Building a model-agnostic evaluation harness
Code execution as a first-class review primitive
I'm Shlok, a software engineer at Greptile. We recently built a code reviewer that, in addition to reviewing pull requests, actually runs the code and shows you what went wrong.
In 1976, Michael Fagan published a paper introducing formal code inspection at IBM. Developers would print out listings, sit in a room together, and read through the code line by line.
Today we still read a diff on a screen. AI tools have made that faster, though most of them are still just reading the code. This approach works for a lot of bugs, the ones that announce themselves plainly in code.
The problem is there's a whole category of bugs that don't show up in code at all; they exist when the program is running. Think of the logic error that needs a specific sequence of state, the UI regression that appears after the page loads, or the race condition that needs a real request. You can read the diff perfectly and still miss these types of bugs completely.
Static code review has a ceiling. It can reason about what the code says. It can't tell you what it does. TREX (which stands for "Test, Run, Execute") is Greptile's response to that ceiling: an execution layer built directly into code review.
Orchestrating agents without wasting context
TREX started as a completely separate product from Greptile, as a standalone agent that generated and ran tests. We hoped that bugs would surface as a result. They didn't. Generating tests wasn't the same activity as finding bugs. When the separate TREX agent tried to write tests, the tests weren't relevant to what the user was trying to do. This created unnecessary noise, and it also missed edge cases. This sounds obvious in hindsight, but it took us more time than expected to learn this lesson.
We'd built these agents to be separate with the assumption it would give each agent its own context window. It also meant both agents ran separately without sharing knowledge. They often overlapped, exploring the same parts of the codebase twice without either agent knowing what the other had already found, ultimately leading to wasted compute.
The obvious fix seemed like combining them into one agent. We tried that, and ran into a different problem: a single agent handling the full review got overloaded. Between spinning up services, taking screenshots, running tests, there was too much context for one agent to manage cleanly.
The solution was to make TREX share the same context as the main Greptile reviewer rather than having it exist entirely as a separate product. It was the first time we were managing agents from within an agent. Unlike two independent agents, this means TREX doesn't start from scratch. It inherits what the Greptile reviewer agent already found, has its own context window, and is scoped to the specific problem it's been asked to investigate.
The Greptile reviewer agent acts as an orchestrator. It reads the diff, identifies issues worth investigating, and spins up a dedicated TREX agent per issue, all running in parallel. The TREX agents have the liberty, the compute, and the knowledge of the orchestrator agent.
A good example of this is a UI feature hidden behind an auth gate. Testing it locally means setting up the environment, handling authentication, getting the feature flag in the right state. A subagent figures all of that out on its own and comes back with a screenshot of the rendered feature.
Designing multi-modal artifacts to show the work
The first version of TREX output findings as bullet points listing out what was tested and what happened. This was a reasonable starting point, but it didn't provide sufficient information.
An agent or a human reviewer reading a bullet point like, "Tested the checkout flow, found failure" wouldn't find it very useful. They wouldn't be able to tell where in the process something went wrong. If the test failed, was it the setup? The assertion? An environment issue? We found an early version of the agent would sometimes hallucinate about how thoroughly it had tested something, claiming to have tried something it hadn't. Bullet points gave us no way to verify.
The fix was to back the bullet point list with a multi-modal artifact set for each TREX finding: screenshots, logs, API traces, execution scripts. Each modality covers a different part of the story. Having a comprehensive picture of everything that was tested for a specific issue is what actually matters.
The first artifact that made us say "Wow" was video. If you push an animation change, TREX captures a video of it playing. You can see exactly what the animation looks like without opening a local environment.
Artifacts also need to be trustworthy. Every artifact has to give the reviewer enough to verify the run themselves. The screenshots, logs, traces, and scripts are all there so a person or downstream agent can look at exactly what happened and confirm it. Bad evidence is worse than no evidence.
The reason artifacts matter, especially for agents downstream, is the same reason teachers require students to show their work. It's analogous to grade school math; you don't know where your answer was wrong until you show the steps. Agents are the same way. If an artifact shows proof of everything the test did to reach a conclusion, the agent can identify exactly which step went wrong. Without that trace, all it has is the answer, and the answer tells you nothing about where to fix it.
If TREX finds a bug, it becomes a comment on the PR. If it runs a feature and everything works, that goes in the summary as proof the change was actually tested. Not every run needs to find something wrong to be useful.
Building a model-agnostic evaluation harness
The frontier race between model providers moves fast. A model that leads on code tasks one month can be behind the next. Building tightly around any single provider's API means rebuilding when rankings shift. That's not a viable long-term strategy.
From the start, we designed TREX around a model-agnostic harness that allows hot-swapping between frontier models without rebuilding. The flexibility goes deeper than most people expect: the main agent and the subagents can use different providers. We can have multiple models running within the same review. This makes it easy for us to pick the best model at any given point, based on internal evals.
Our current evaluation involves measuring recall (e.g., how many real bugs are caught, measured against open-source PRs or customer data where comments were addressed) and precision (e.g., consistency across runs: if you review the same PR twice, are you finding roughly the same set of issues?).
We intentionally deprioritize latency in our evaluation. A developer waiting on a review would rather wait a little longer and get something accurate than get a fast answer they can't trust.
The open source evaluation harness we use performs on par with native provider harnesses. There's no meaningful quality penalty for being model-agnostic which, if you'd asked me to guess before we tested it, I wouldn't have been confident about.
TREX's differentiation is not which model it's running. It's the infrastructure around the model: the codebase indexing, the orchestration, the artifact generation, the evaluation framework. We don't need to care about intelligence as a differentiator, because intelligence is something that will continuously improve. Greptile's job is to build the harness, the architecture, and the artifact pipeline that makes the difference in practice.
Code execution as a first-class review primitive
Every TREX review spins up a disposable sandboxed environment: an isolated compute instance per review, started fresh in milliseconds, thrown away when the run is done. These environments start fast, isolate execution properly, and can run real projects; not just unit tests against a mock, but actual services with actual dependencies. Many bugs only appear end-to-end.
Starting from nothing every time would be too slow. So we rely on reusable base images and per-repository snapshots. A repository can be cloned once, captured, and resumed. Each review still fetches the exact PR commits and rotates credentials before execution begins. The cache speeds up setup without freezing stale state into the environment. A cache that includes too little is slow. A cache that includes too much becomes haunted. TREX needs the useful kind: warm enough to move quickly, fresh enough to trust.
The sandbox is what makes the artifacts trustworthy. When TREX reports that a page rendered with a broken layout, you know the code actually ran in a real environment. That's a different claim than, "The code looks like it might render with a broken layout."
These pieces — the subagent architecture, the artifact verification bar, the sandboxed execution environment, and the evaluation harness — aren't separate features. They're one system. The orchestrator identifies issues worth running. The sandbox makes running safe and fast. The artifact pipeline makes the results trustworthy enough to act on. The evaluation harness makes sure the right model is doing the work. Each one makes the others more valuable. The review it comes up with is a reproducible experiment with attached evidence.
Our vision is bigger than code review as it exists today. The goal is a world with no bugs. To get there, we're no longer thinking of ourselves as a code review tool. We want to be a validation suite: an end-to-end layer that mimics what engineering teams have done for decades, but automated and running on every PR. TREX is another step toward that vision.
product Introducing Greptile 2.0
Discover Greptile 2.0, the AI platform that understands your codebase—boost development with smart code reviews, JIRA integration, and custom workflows.
programming A statistical study of PRs opened on openclaw/openclaw
OpenClaw became the fastest-growing repo in GitHub history almost overnight. The PR data offers a preview of what the future of open source contribution may look like.
programming Codebases are uniquely hard to search semantically
Explore why semantic search in codebases is challenging and how natural language processing and granular chunking improve AI code understanding.
