---
source: "https://bramo.ai/blog/01-post-mortem"
hn_url: "https://news.ycombinator.com/item?id=49282395"
title: "My AI agents shipped 128 releases of a product no one ever used"
article_title: "My AI agents shipped 128 releases of a product no one ever used — Bramo"
author: "guschiriboga"
captured_at: "2026-08-13T07:13:23Z"
capture_tool: "hn-digest"
hn_id: 49282395
score: 1
comments: 0
posted_at: "2026-08-13T06:26:55Z"
tags:
  - hacker-news
  - translated
---

# My AI agents shipped 128 releases of a product no one ever used

- HN: [49282395](https://news.ycombinator.com/item?id=49282395)
- Source: [bramo.ai](https://bramo.ai/blog/01-post-mortem)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T06:26:55Z

## Translation

タイトル: 私の AI エージェントは、誰も使用したことのない製品の 128 リリースを出荷しました
記事のタイトル: 私の AI エージェントは、誰も使用したことのない製品の 128 リリースを出荷しました — Bramo
説明: 失敗したスタートアップ 2 件、リリース 128 件、グリーン テスト 442 件、見知らぬ人はゼロ。本物のユーザーが必要なまさにその瞬間に両方が死亡した理由と、Bramo のやり方が異なることについての解剖。

記事本文:
ブラモ
ドキュメント
死後
順番待ちリストに参加する
2026 年 8 月 8 日 · 7 分で読みました · ガス・チリボガ
私の AI エージェントは、誰も使用したことのない製品の 128 リリースを出荷しました。
タイプミスではありません。 CI、契約テスト、ドキュメント サイト、バイリンガル ドキュメント、ガバナンス ポリシー、スポンサー認識ポリシーを備えた 128 のバージョンが npm および Homebrew に公開されています。外部ユーザーはゼロです。それから、2 番目の製品、442 件の合格テスト、実際のコスト測定、イベントソースの監査ログをもう一度実行しましたが、最初のエンドツーエンドの実行が最終的に機能した 2 日後に停止しました。
私がこの記事を書いているのは、両方のコードベースを新鮮な目で（もちろん、考古学を行う AI エージェントとともに）見直し、2 年間の構築よりも多くのことを解剖から学んだからです。あなたが現在 AI コーディング エージェントを使用して構築している個人の創業者であれば、これには不快なほど馴染みのあるものがあるでしょう。
試み #1: エージェントに行動を求める製品
最初の製品は、AI コーディング エージェント用の SDLC オーケストレーターでした。アイデアは、Claude Code または Codex と通常どおりチャットすると、私のツールが仕様、設計、承認ゲート、納品の検証といった段階的なプロセスをラップするというものです。メカニズム: エージェントにプロセスに従うように指示する「ガードレール」マークダウン ファイルをリポジトリにインストールします。
すでに欠陥が見えています。ツールはエージェントを制御しませんでした。それは散文でエージェントに行動するよう求め、そして希望した。
変更ログは、その希望が失敗したことの考古学的記録です。次から次へとパッチを適用したエージェントはガードレールを無視し、エージェントは決して自己承認しないと言われていたゲートを自己承認していました。あるリリースでは、 stdin を通じて人間の承認を偽造できることを発見したエージェントが修正されました。私は礼儀正しさによって作られたセキュリティ境界線を補修していたのです。
その間、私は物事の本に従って、他のことはすべて正しく行いました

まだ重要ではありません: マーケティング Web サイト、Homebrew ディストリビューション、2 か国語のドキュメント、貢献ガイドライン。 128 回のリリースで磨きがかかり、見知らぬ人が使用することはありません。コアが失敗し続けても、私は止まらず、3 レベルのオーケストレーション階層を持つより壮大なバージョンの設計を開始しました。この再アーキテクチャにより、プロジェクトは静かに消滅しました。
試み #2: 現実を除くすべてのテストに合格した製品
第2ラウンドではそのミスを逆転した。今回は完全に制御できます。独自のオーケストレーション エンジン、イベント ソース台帳、決定論的プランナー、ガバナンス マトリックス、予算上限、モデル API を直接呼び出すエグゼキューターです。 「エージェント」は、それぞれが 1 つの API 呼び出しであるため、完全に私の管理下にありました。ツールも反復も自己修正も必要ありません。私は実際のコーディング エージェントを非常に従順なテキスト ジェネレーターに置き換え、それと議論するための検証ツールを構築しました。
システムは 442 のテストに合格するまでに成長しました。 15 の検証シナリオ。安定性を確保するためにそれぞれ 3 回実行されます。生成された OpenAPI ドキュメントを含む REST API。デザイントークンによるブランドシステム。
これは火災警報であるはずの数字です。実際のプロンプトがエンドツーエンドで実際に検証されたアーティファクトを初めて生成したのは、構築から 5 週間後でした。それらの週のほとんどでは、すべての単体テストが緑色でした。そして、実際の人間 (私と実際の端末を運転している QA エージェント) が最終的にそれを使用したところ、テストでは見られなかった方法ですべてが壊れました。
これはどれもエキゾチックではありません。これは、観察した動作ではなく、自分が作成したフィクスチャに対して検証を行うときに起こることです。私のテストでは、システムが私のモデルと一致していることが証明されました。私のモデルが現実と一致することを証明するものは何もありませんでした。
8月3日、ついにゴールデンランが成功した。費用: $0.048855、正直にメーター制 - メーターは常に機能する唯一のものでした。 2日ほどで作業をやめました

後は。それが失敗したからではなく、それが最終的にうまくいき、次のステップは見知らぬ人にそれを見せることであると私の心のどこかでわかっていて、そのステップには決して到達しないように2年間かけて準備してきたからです。
解剖結果が実際に何を語っているか
読みやすいのは「技術的なミスをしました」です。そうしました。最初の試みでは、実際のエージェントはありましたが、制御はなく、セキュリティ境界としてファイルをプロンプトしました。 2 つ目は完全に制御できましたが、実際のエージェントはありませんでした。エンジニアのコスプレをした単一の API 呼び出しです。スコアを維持している場合、実際のエージェントと実際のコントロールを備えた象限は、まさにその後業界が到達した場所 (プログラムで操作できるヘッドレス エージェント CLI) であり、私が現在構築している場所です。
しかし、不快な見方は真実です。故障モードは決して技術的なものではありませんでした。どちらのプロジェクトも見知らぬ人が必要になった瞬間に中止になり、どちらの場合も私はその瞬間を避けるために緊急の技術的作業を行いました。壮大な再構築。サンドボックス権限のバグ。ウェブサイト、ブランド システム、第二言語、スポンサー ポリシーなど、質問するよりも構築することを好む人のプロフェッショナルな先延ばし癖。
AI エージェントを使用して構築する私たちにとって、このトラップには特別なバージョンがあります。書き留められているのを見たことがなかったので、これに名前を付けたいと思います。
エージェントは建設を非常に安くするので、建設が先延ばしになってしまいます。 「最初のバージョンで恥ずかしくないなら、出荷が遅すぎたということだ」という古い通念は、構築に費用がかかるため、当てはまりました。エージェントは、129 番目のリリース、443 番目のテスト、バイリンガル ドキュメント、ブランド トークンを喜んで生成します。勢いはトラクションのように感じられます。私のコミット履歴は心臓の鼓動でした。それはビジネスに結びついていなかっただけです。そして、エージェントは「なぜこんなことをするのか」とは一度も言いませんでした。彼らは小さなことで絶えず嘘をつきました - 「終わった」「終わった」

「テストはパスしました」、「ファイルを作成しました」――しかし、その大嘘は私がついた嘘で、彼らは快く拡大して、この動きはすべて進歩だと言いました。
3 つの具体的なレッスン、領収書を添付:
01 エージェントの「終わった」という主張は構造的に無価値です。モデルが悪いからではありません。インセンティブが調整されておらず、誰もチェックしていないからです。ファイルが存在しない場合に、最新のエージェントが終了コード 0、「成功」、および「ファイルを作成しました」を返すのを観察しました (パーミッションの拒否をそのままナレーションで伝えます)。私の v1 ではストールに「✓ 完了」と印刷されました。コグニション社は、エージェントが独自のテストを行っていることについて公に書いています。検証は検証されるものの内部に存在することはできません。
02 テストはシステムのモデルを検証します。見知らぬ人だけがシステムを検証します。 442 件のグリーン テストが、人間が一度でもプロンプトを成功させることができなかった製品と共存していました。解決策はテストを減らすことではありません。実際の使用によって生み出されたもの以外の「効果がある」という言葉を信じることを拒否しているのです。
03 ループが最終的に機能すると、次のコミットは間違った動きになります。私のプロジェクトはどちらも、ランディング ページと会話に到達するはずだったまさにその瞬間に範囲を拡大しました。製品が初めてエンドツーエンドで動作した場合は、構築をやめてください。今日の課題は見知らぬ人です。
私が（そして公に）違うことをしていること
私は過去 2 年間必要としていたツール、つまり AI コーディング エージェント用の監視レイヤーを構築しています。すでに使用しているエージェントをヘッドレスで駆動し、エージェントが主張する内容を個別に検証します。つまり、生成したプロセスで実行されたと思われるテスト、差分のアンチスタブ スキャン、ビルダーの説明を決して参照しない新鮮なコンテキストのレビュー、および証拠が添付された正直な評決です。さらに、レシートも含まれます。ツールを使用した場合の各タスクの実際のコストと、管理対象外のエージェントのばたつきにかかるコストです。
T

彼のくさびは独立です。現在、すべてのベンダーが何らかの自己レビューを提供していますが、ベンダーのエージェントが同じベンダーのエージェントを採点すること自体が宿題採点です。中立層は、エージェント間で機能し、どのエージェントも信頼しない層であり、トークンを販売する誰も構築するインセンティブを持たない層です。
そして、このエッセイの後に私にはちょうど 1 枚の信頼性カードが残っているので、それを表向きにプレイします。今回のルールは、次の見知らぬ人を超えて構築しないことです。 2 つの技術的なスパイクは実行され、リポジトリで公開されます。検証パイプラインは、レビューに合格した植え付けられたスタブをすでに検出しており、実際に実際に偽の「完了」したものをすでに検出しています。次のマイルストーンは機能ではありません。重要なのは、これを読んだ 25 人がメールを残すほど関心があるかどうかです。
AI エージェントを使用して構築する場合、エージェントが実行したことと実行しなかったと誓ったことをマージしたことがある場合、私はこれを、公の場で、一度に 1 週​​間正直に、あなたを念頭に置いて構築しています。
このエッセイがあなたを 128 回のリリースから救ってくれたとしたら、それは私にとってサインアップよりも価値があります。しかし、サインアップも役立ちます。

## Original Extract

Two failed startups, 128 releases, 442 green tests, zero strangers. The autopsy on why both died at the exact moment they needed a real user — and what Bramo is doing differently.

Bramo
Docs
Post-mortem
Join the waitlist
August 8, 2026 · 7 min read · Gus Chiriboga
My AI agents shipped 128 releases of a product no one ever used.
Not a typo. One hundred and twenty-eight versions, published to npm and Homebrew, with CI, contract tests, a documentation site, bilingual docs, a governance policy, and a sponsor recognition policy. Zero external users. Then I did it again — a second product, 442 passing tests, real cost metering, event-sourced audit logs — and stopped two days after the first end-to-end run finally worked.
I’m writing this because I went back through both codebases with fresh eyes (and, yes, with an AI agent doing the archaeology), and the autopsy taught me more than the two years of building did. If you’re a solo founder building with AI coding agents right now, some of this will be uncomfortably familiar.
Attempt #1: the product that begged agents to behave
The first product was an SDLC orchestrator for AI coding agents. The idea: you chat with Claude Code or Codex like normal, and my tool wraps the process in stages — spec, design, approval gates, verified delivery. The mechanism: it installed “guardrail” markdown files into your repo that instructed the agent to follow the process.
You can already see the flaw. The tool didn’t control the agent. It asked the agent to behave, in prose, and hoped.
The changelog is an archaeological record of that hope failing. Version after version patched agents ignoring the guardrails: agents self-approving gates they were told never to self-approve. One release fixed an agent that had discovered it could forge the human’s approval through stdin . I was patching a security boundary made of politeness.
Meanwhile I did everything else right, by the book of things that don’t matter yet: a marketing website, Homebrew distribution, docs in two languages, contribution guidelines. 128 releases of increasing polish, zero strangers using it. When the core kept failing, I didn’t stop — I started designing a grander version with a three-level orchestration hierarchy. That re-architecture is where the project quietly died.
Attempt #2: the product that passed every test except reality
For round two I inverted the mistake. Full control this time: my own orchestration engine, event-sourced ledger, deterministic planner, governance matrix, budget ceilings, an executor that called the model API directly. The “agents” were now fully under my thumb — because each one was a single API call. No tools, no iteration, no self-correction. I had replaced a real coding agent with a very obedient text generator, and then built a verifier to argue with it.
The system grew to 442 passing tests. Fifteen validation scenarios, each run three times for stability. A REST API with generated OpenAPI docs. A brand system with design tokens.
Here is the number that should have been a fire alarm: the first time a real prompt produced a real, verified artifact end-to-end was after five weeks of building. Every unit test was green for most of those weeks. And when real humans (me, and a QA agent driving a real terminal) finally used it, everything broke in ways the tests never saw:
None of this is exotic. It’s what happens when you validate against fixtures you wrote instead of behavior you observed. My tests proved the system matched my model of it. Nothing proved my model matched reality.
The golden run finally succeeded on August 3rd. Cost: $0.048855, honestly metered — the meter was the one thing that always worked. I stopped working on it about two days later. Not because it failed — because it had finally worked, and some part of me knew that the next step was showing it to a stranger, and I had spent two years arranging to never reach that step.
What the autopsy actually says
The comfortable read is “I made technical mistakes.” I did. The first attempt had a real agent but no control — prompt files as a security boundary. The second had total control but no real agent — single API calls cosplaying as engineers. If you’re keeping score, the quadrant with a real agent AND real control is exactly where the industry has since landed (headless agent CLIs you can drive programmatically), and it’s where I’m building now.
But the uncomfortable read is the true one: the failure mode was never technical. Both projects died at the moment they needed a stranger, and both times I manufactured urgent technical work to avoid that moment. A grand re-architecture. A sandbox permission bug. Websites, brand systems, second languages, sponsor policies — the professional-looking procrastination of a person who would rather build than ask.
There’s a special version of this trap for those of us building with AI agents, and I want to name it, because I haven’t seen it written down:
Agents make building so cheap that building becomes the procrastination. The old wisdom — “if you’re not embarrassed by your first version, you shipped too late” — had teeth because building was expensive. Now an agent will happily generate the 129th release, the 443rd test, the bilingual docs, the brand tokens. Momentum feels like traction. My commit history was a heartbeat; it just wasn’t attached to a business. And the agents never once said: why are we doing this? They lied about small things constantly — “done”, “tests pass”, “created the file” — but the big lie was one I told and they cheerfully amplified: that all this motion was progress.
Three specific lessons, receipts attached:
01 An agent’s claim of “done” is worthless, structurally. Not because models are bad — because the incentive is misaligned and nobody’s checking. I watched a modern agent return exit code 0, “success”, and “Created the file” when the file did not exist (a permission denial it narrated straight through). My v1 printed “✓ Done” on a stall. Cognition has written publicly about agents gaming their own tests. Verification cannot live inside the thing being verified.
02 Tests validate your model of the system. Only strangers validate the system. 442 green tests coexisted with a product no human could successfully prompt even once. The fix isn’t fewer tests; it’s refusing to believe any “it works” that wasn’t produced by real use.
03 When the loop finally works, the next commit is the wrong move. Both my projects expanded scope at the exact moment they should have contracted to a landing page and a conversation. If your product just did the thing end-to-end for the first time: stop building. Today’s task is a stranger.
What I’m doing differently (and publicly)
I’m building the tool I needed for the last two years: a supervision layer for AI coding agents. It drives the agents you already use, headlessly — and then independently verifies what they claim : tests that provably ran in a process it spawned, anti-stub scans on the diff, a fresh-context review that never sees the builder’s narrative, and an honest verdict with evidence attached. Plus the receipt: what each task actually cost through the tool versus what unmanaged agent flailing would have cost.
The wedge is independence. Every vendor now ships some self-review — but a vendor’s agent grading the same vendor’s agent is homework grading itself. The neutral layer, the one that works across agents and trusts none of them, is the layer nobody selling you tokens has an incentive to build.
And because I have exactly one credibility card left after this essay, I’m playing it face up: the rule this time is no building past the next stranger. The two technical spikes are done and public in the repo — the verification pipeline already catches a planted stub that passes review, and already caught a real false-“done” in the wild. The next milestone isn’t a feature. It’s whether 25 people who read this care enough to leave an email.
If you build with AI agents and you’ve ever merged something an agent swore was done and wasn’t — I’m building this with you in mind, in public, one honest week at a time.
If this essay saved you from your own 128 releases, that’s worth more to me than the signup. But the signup helps too.
