---
source: "https://vercel.com/blog/building-a-software-factory-for-ai-sdk"
hn_url: "https://news.ycombinator.com/item?id=49296646"
title: "Building a software factory for AI SDK"
article_title: "Building a software factory for AI SDK - Vercel"
author: "lgrammel"
captured_at: "2026-08-14T10:55:55Z"
capture_tool: "hn-digest"
hn_id: 49296646
score: 1
comments: 0
posted_at: "2026-08-14T09:58:42Z"
tags:
  - hacker-news
  - translated
---

# Building a software factory for AI SDK

- HN: [49296646](https://news.ycombinator.com/item?id=49296646)
- Source: [vercel.com](https://vercel.com/blog/building-a-software-factory-for-ai-sdk)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T09:58:42Z

## Translation

タイトル: AI SDK のソフトウェア ファクトリの構築
記事のタイトル: AI SDK のソフトウェア ファクトリの構築 - Vercel
説明: AI SDK の問題と PR を自律的に処理するソフトウェア ファクトリを構築し、すべてのマージを人間が制御します。 4 週間後、マージされた PR の 25 ～ 40% が作成されました。

記事本文:
AI SDK のソフトウェア ファクトリーの構築 - Vercel コンテンツへスキップ ワードマークをコピー ロゴをコピー ブランド資産をダウンロード ブランド ガイドライン 製品 エージェント スタック AI SDK
AI SDKのソフトウェアファクトリーの構築
AI SDK は、世界で最も人気のあるオープンソース AI プロジェクトの 1 つです。週に 2,000 万を超える npm ダウンロードがあり、リポジトリには 26,000 を超えるスターが付いています。コードベースを維持するということは、4 つの移動ターゲットを同時に追跡することを意味します。
モデルプロバイダー: 新しいプロバイダー、新しい機能、新しいバグ
UI フレームワーク: React、Next.js、Svelte、Vue などのバインディング
サンドボックス: エージェントがコードを実行する実行環境
ハーネス: Codex、Claude Code、Pi などのアダプター
数年間の成長を経て、リポジトリには毎月 100 以上の新しい号が発行されるようになり、Anthropic の Opus 4.6 モデルがリリースされたとき、PR は変曲点に達しました。 6 月下旬までに、この複合的な問題により 1,000 を超える未解決の問題と約 800 のプル リクエストが蓄積されました。
その未処理は規律の問題ではありません。どんなに優秀なメンテナであっても、一生懸命働いてそのギャップを埋めることはできません。また、コードの生成は安価であるため、コードは増大する一方です。
私たちは規模を拡大しようとする代わりに、ソフトウェア ファクトリーを構築しました。 4 週間後、マージした PR の 25 ～ 35% が作成され、問題の 70 ～ 80% がクローズされました。
リンクを見出しにコピーする どのタイプの工場を建設するかを決定する
何かを構築する前に、次の 3 つの質問に答える必要がありました。
エージェントを使用した既存のアプローチが十分ではなかった理由
AI SDK のようなプロジェクトに適合する自動化レベルはどれか
自動化と人的労力をリスクに合わせて調整する方法
見出しへのリンクをコピーします エージェントをさらに追加してみませんか?
優秀なメンテナはすでにエージェントを積極的に使用しています。ミッチェル・ハシモトは、エージェントが常に動作するという目標を持って Ghostty を実行し、エージェントの失敗をすべて AGENTS.md にエンコードして、失敗が繰り返されないようにします。サイモン・ウィル

ison は 4 つのコーディング エージェントを並行して実行し、Vercel Agent や CodeRabbit などのレビュー ボットは数百万のリポジトリに常駐します。 Daniel Stenberg のような他のメンテナーは、AI によって生成されたカールへの送信をブロックすることを選択しました。
これらはすべて役に立ちますが、核となる制約を解決するものはありません。これらのソリューションはいずれも、依然としてすべての変更を 1 人の人間の注意を介してルーティングしています。私たちは人間の責任が依然としてエージェント エンジニアリングの信頼の中核であると信じているため、工場では第一原則としてレビュー担当者の効率性を解決する必要があることを認識していました。
見出しへのリンクをコピーします AI SDK に適合する自動化はどのようなものですか?
ソフトウェア工場はさまざまな領域に存在します。一方の端は完全な自動化であり、人間がコードを読むことなく、エージェントが記述、出荷、展開します。中央には Codex や Claude Code のようなハーネスがあり、ここで人がエージェントまたはエージェントの艦隊を操縦します。最後にあるのは、ペースメーカーや自動運転車のファームウェアなど、リスクが高すぎるためほとんど自動化したくないソフトウェアです。
AI SDK は慎重な最終段階に近いところで動作する必要があります。これは基盤となる AI インフラストラクチャであり、その上に何百万ものアプリケーションが構築されているため、品質とセキュリティには交渉の余地がなく、何を出荷するかを人間が制御する必要があります。私たちの工場では、人間を撤去することなく、人間の周囲のライフサイクルを高度に自動化する必要がありました。
見出しへのリンクをコピーします 自動化と取り組みをリスクにどのように調整するか?
特定の変更に対して必要な人間によるレビューの深さは、リスクに応じて変化します。
詳細な機能仕様を含む一元化されたロードマップにより、事前にリスクを定義し、ソフトウェア ファクトリに導入される作業を直接形作ることができます。しかし、AI SDK のようなオープンソース プロジェクトもコミュニティから問題やプル リクエストを受け取りますが、それらがプロジェクトの目標と一致しているという保証や、安全に変更を加えられるという保証はありません。リスクが高いということは人間のことを意味します

判断はシステムのさらに重要な部分です。
その判断に合わせてファクトリーを最適化するには、エージェントがリクエストに基づいてコードを生成する以上の作業を行う必要があることを私たちは認識していました。プロジェクト全体のコンテキストで作業単位全体を評価する必要がありました。
私たちの目標は、工場が一連の文書化された証拠を含め、適合性とリスクの両方について各変更の包括的な評価を生成し、レビュー担当者が適切な量の労力を適用しやすくすることでした。
ドキュメントの修正を確認するために一目で確認できる
明確に定義されたプロバイダーの変更は集中的に検証されます
新しいパブリック API が徹底的にレビューされる
見出しへのリンクをコピーします。
ai-sdk-factory は、受信した問題と AI SDK のプル リクエストを自律的に処理するソフトウェア ファクトリです。工場のエージェントは、バグの再現、機能の実装、古い SDK バージョンのバックポートの作成など、レビュー可能な特定のタスクを実行します。すべての変更をマージするなど、プロセス全体を通じて人間が制御を維持します。
私たちは工場を一度に出荷したわけではありません。私たちはそれを段階的に構築し、問題をバグ、機能、またはドキュメントの更新として分類するプロセスの開始から始めました。この最初のステップにより、形状バックログの可視性が向上しただけでなく、後で構築した他の特化したエージェントに役立つコンテキストも渡されました。
私たちは、それぞれの新しいタイプの機能を構築するための複数の方法をプロトタイプ化し、その過程で、生産に入る工場のアーキテクチャを形作る一連の指針を開発しました。
リンクを見出し「タスクごとに 1 つのエージェントを構築する」にコピーします
分類エージェントが高レベルの精度に達すると、バグの再現、修正、それらの修正のレビューの自動化に重点を置きました。私たちは各ステップのスキルを備えた単一のエージェントの構築を検討しましたが、すぐに次のことに気づきました。

時間の経過とともに、メンテナンスとトラブルシューティングの負担が増大する可能性があります。
代わりに、特定のタスクごとに単一のエージェントを構築し、すべての新機能の推論、分離テスト、デバッグを容易にしました。それぞれのスコープは特定のジョブに限定され、独自のプロンプト、コンテキスト、および評価が行われます。
現在、工場にはフローの各ステップを担当する専任のエージェントがいます。
見出しへのリンクをコピーします 最初からすべてを保護します
バグの再現はファクトリが制御していないコンテンツに基づいてコードを実行する最初のステップであったため、2 番目のエージェントを使用してセキュリティを実装しました。
パブリック リポジトリで動作するファクトリは、攻撃者が制御する入力を想定する必要があります。すべての課題、プル リクエスト、コメント、およびそれらに含まれるリンクは信頼できません。成功したオープンソース プロジェクトは価値の高い標的であるため、脅威は悪意のあるコードの変更やサプライ チェーン攻撃から、リソースの枯渇、API キーの漏洩、即時漏洩に至るまで多岐にわたります。
サンドボックスは防御の基盤です。 ai-sdk-factory のすべてのエージェントは、コード、ランタイム、およびエージェントの特定のタスクに必要なシークレットのみを含む分離された Vercel サンドボックス内で実行されます。これらのガードレールを使用すると、信頼できないコンテンツがエージェントの提案を形作る可能性がありますが、1 つのタスクが与える可能性のある損害はサンドボックスに抑えられます。
また、サンドボックスの周囲にシールド層を構築し、エージェントがネットワーク経由でアクセスできる範囲を制御し、攻撃者が隔離された環境から秘密を引き出すために使用するパスをブロックしました。
最後の防御線は人間によるレビューです。AI SDK チームの人間の承認がなければ、何もマージされません。
見出しへのリンクをコピーします。ローカルで価値を構築してからクラウドに移動します。
工場用に構築した最初のいくつかのエージェントは、ローカル CLI を通じて実行されました。そのおかげで、私たちのチームは不正確な点に気づき、面倒に感じたときに迅速に反復できるようになりました。

を考え、さまざまなアイデアを試作しました。
CLI を介して複数のステップが確実に実行されるようになると、システムを管理対象インフラストラクチャに移行する準備が整いました。 ai-sdk-factory は以下を使用します。
API、ワーカー、Webhook Ingress 用の Vercel 関数
タスク実行用の Vercel キュー
エージェント ワークスペース用の Vercel サンドボックス
工場データ用の Neon Postgres
現在、GitHub Webhook は問題キューにフィードし、到着するとすぐにワーカーが自動的にそれを取得し、サンドボックスでエージェントの実行を開始します。また、すべての実行を並行して追跡し、レビュー担当者チームのキューを視覚化する監視 UI も構築しました。
見出しへのリンクをコピーします 当社の工場を通じてソフトウェアを出荷
7 月 24 日、コミュニティ メンバーが OpenAI Web 検索でのブロック ドメインのサポートを要求し、その要求は問題 #17898 になりました。次のセクションでは、ソフトウェア ファクトリが問題を処理し、機能の実装を含むプル リクエストをオープンし、最終的にマージされた機能を SDK の v5 および v6 にバックポートするために実行した各ステップについて説明します。
見出しへのリンクをコピーします。 分類 ファクトリで実行される最初のエージェントは、問題とプル リクエストを分類します。この場合、ai-sdk-factory ボットは問題についてコメントしてラベルを適用し、タイプを高い信頼度で機能として識別しました。代理人はコメントに分類の根拠を含めた。
見出しへのリンクをコピーします。 分析 分類後、分析エージェントが実行されます。工場のエージェントは機能リクエストやバグの技術的妥当性について何の仮定も立てないため、分析エージェントはブロックされたドメインがサポートされていないことを確認するプローブを作成しました。 issue-17898-type-probe.ts が生成されて実行されましたが、ブロックされたドメインを探しても見つからず、エラーが発生して失敗しました。
プローブの失敗により、その機能が main に欠落していることが証明されました。

エージェントはそれを証拠として分析に含めました。
次に、分析エージェントは調査結果を使用して機能の仕様を構築しました。オプションのblockedDomainsフィルターを既存のWeb検索ツールに追加し、それをプロバイダーのblocked_domainsフィールドにマップします。
また、エージェントは、仕様が SDK のプロバイダー アダプター アーキテクチャに適合しており、下位互換性があり、ドキュメントの変更範囲も限定されていることも確認しました。
リンクを見出し「実装」にコピーします。その後、別のエージェントが仕様を実装し、プル リクエストをオープンしました。実装エージェントはライブ エンドツーエンド テストを実行し、wikipedia.org がブロックされた状態で OpenAI Web 検索を実行し、ドメインに到達できないことを確認しました。テストはプル リクエストの追加証拠として含まれていました。
見出しへのリンクをコピーします。 自動レビュー 次に、レビュー エージェントが変更を採点し、懸念事項が見つからなかったら承認しました。エージェントは、この機能が完全に実装されていると評価し、次のように評価しました。
下位互換性のリスク: 低い
見出しへのリンクをコピー 人間によるレビュー 最後に、ラーズはエージェントから一連の証拠を読み、コードの変更をレビューし、PR #18033 を main にマージしました。
バックポートの見出しへのリンクをコピー Lars が最初の PR をマージすると、ai-sdk-factory は 2 つのバックポート (v6 の場合は #18035、v5 の場合は #18036) に対して追加の PR を開きました。 v5 バックポートは適切に適用されなかったため、ファクトリ エージェントは競合状態にラベルを付けてコミットし、修正を特定して検証し、17 分後にプッシュしました。レビューの後、Lars は両方のバックポート PR をマージしました。
ソフトウェア工場を本番稼働させてから 4 週間あまりが経ちました。結果は次のとおりです。
私たちが毎週マージする PR の 25 ～ 35% は、現在、ai-sdk-factory エージェントによって作成されています。
ファクトリー PR は、v6 リリースラインへの毎週のマージの 50% を超えており、v5 も同様に見えます。バックポートは以前は私たちの仕事でした

マージ競合に対処するのは労力に値しないため、スキップされました。 v5 と v6 のサポートが大幅に改善されました。
7 月には、クローズされた問題の 75% 以上が工場によってクローズされました。
未解決の問題は 6 月下旬の 1,022 件のピークから、8 月初旬には 844 件に減少し、未解決のバグは約 25% 減少しました。
ai-sdk-factory はパブリックで実行されるため、リポジトリ上で作成されたすべてのプル リクエストを確認できます。
見出しへのリンクをコピーする 工場の改善が仕事になる
工場の運営で最も興味深いのは、工場が失敗したときに何が起こるかということです。すべての実行は、成功、欠陥、ブロック、または手動の 4 つの方法のいずれかで終了します。成功のみが出荷されるため、残りはフィードバックとしてシステムに再入力される信号になります。
実行に欠陥があるということは、エージェントが間違ったものを作成したことを意味し、修正はより良いプロンプト、より良いコンテキスト、または新しい評価ケースであり、次回同じ間違いが自動的に検出されるようになります。
実行がブロックされたということは、環境に資格情報、サービス、依存関係などの何かが欠落しており、修正によりそれがプロビジョニングされていることを意味します。
手動による稼働は、私たちが意図的に引いた境界線を示しており、工場の改善によって境界線を取り除くことが正当化されるかどうかを問うことになります。
これらの修正のそれぞれが自動化の境界を拡張し、工場は毎週、必要な作業を処理できるようになります。

[切り捨てられた]

## Original Extract

We built a software factory that autonomously processes issues and PRs for the AI SDK, with humans in control of every merge. Four weeks in, it authors 25-40% of merged PRs.

Building a software factory for AI SDK - Vercel Skip to content Copy Wordmark Copy Logo Download Brand Assets Brand Guidelines Products Agent Stack AI SDK
Building a software factory for AI SDK
The AI SDK is one of the most popular open-source AI projects in the world. It serves over 20 million npm downloads a week and the repo has over 26,000 stars. Maintaining the codebase means tracking four moving targets at once:
Model providers: new providers, new capabilities, and new bugs
UI frameworks: bindings for React, Next.js, Svelte, Vue, and others
Sandboxes: the execution environments agents run code in
Harnesses: adapters for Codex, Claude Code, Pi, and others
After multiple years of growth, the repo was getting 100+ new issues every month, and when Anthropic's Opus 4.6 model was released, PRs hit an inflection point. By late June, that compounding had accumulated over 1,000 open issues and almost 800 pull requests.
That backlog is not a discipline problem. No maintainer, however good, can close that gap by working harder, and because generating code is cheap, it will only grow.
Instead of trying to scale ourselves, we built a software factory. Four weeks in, it authors between 25 and 35% of PRs we merge and closes 70-80% of issues.
Copy link to heading Deciding what type of factory to build
Before we built anything, we had to answer three questions:
Why our existing approach using agents wasn't enough
What level of automation fit a project like the AI SDK
How to align automation and human effort to risk
Copy link to heading Why not add more agents?
The best maintainers are already using agents aggressively. Mitchell Hashimoto runs Ghostty with the goal of an agent always working, and encodes every agent failure in AGENTS.md so it never repeats. Simon Willison runs four coding agents in parallel, and review bots like Vercel Agent and CodeRabbit sit on millions of repos. Other maintainers like Daniel Stenberg have opted to block AI-generated submissions to curl.
All of it helps, but none of it solves for the core constraint: every one of these solutions still routes every change through one human's attention. We believe that human accountability is still the core of trust in agentic engineering, so we knew our factory needed to solve for reviewer efficiency as a first principle.
Copy link to heading What kind of automation fits AI SDK?
Software factories sit on a spectrum. At one end is full automation, where agents write, ship, and deploy without a human ever reading the code. In the middle are harnesses like Codex and Claude Code, where a person steers an agent or fleet of agents. At the far end is software you barely want to automate because the risk is too high, like firmware in a pacemaker or self-driving vehicle.
The AI SDK needs to operate closer to the careful end. It is foundational AI infrastructure with millions of applications built on top, so quality and security are non-negotiable, and a human needs to have control over what ships. We needed our factory to heavily automate the lifecycle around the human, without removing them.
Copy link to heading How do we align automation and effort to risk?
For a given change, the depth of human review needed scales with risk.
Centralized roadmaps with detailed feature specs can define risk up front and directly shape the work going into a software factory. But open-source projects like AI SDK also get issues and pull requests from the community, and there is no guarantee they align with the goals of the project, or that changes are safe to make. Higher risk means human judgment is an even more critical part of the system.
To optimize our factory for that judgment, we knew agents had to go beyond generating code based on a request; they needed to evaluate full units of work in the context of the entire project.
Our goal was for the factory generate a comprehensive assessment of each change for both fit and risk, including a full chain of documented evidence, making it easy for reviewers to apply the right amount of effort:
Docs fixes get a quick glance for verification
Well-defined provider changes get focused validation
A new public API gets deep review
Copy link to heading What we built
ai-sdk-factory is a software factory that autonomously processes incoming issues and pull requests for AI SDK. Agents in the factory perform specific, reviewable tasks, like reproducing bugs, implementing features, and creating backports for older SDK versions. A human stays in control throughout the entire process, including merging every change.
We didn't ship the factory in one swing. We built it incrementally, starting at the beginning of the process with classification of issues as bugs, features, or documentation updates. That first step not only gave us more visibility into the shape backlog, but also passed helpful context to the other specialized agents we built later.
We prototyped multiple ways to build each new type of functionality, and in the process developed a set of guiding principles that shaped the architecture of the factory that went into production.
Copy link to heading Build one agent per task
Once our classification agent reached a high level of accuracy, we focused on automating bug reproduction, fixes, and review of those fixes. We explored building a single agent equipped with skills for each step, but quickly realized that would come with a higher maintenance and troubleshooting burden over time.
Instead, we built a single agent for each specific task, making every new capability easier to reason about, test in isolation, and debug. Each one is scoped to a specific job, with its own prompts, context, and evals.
Today the factory has dedicated agents for every step in the flow:
Copy link to heading Secure everything from the start
We implemented security with the second agent, because bug reproduction was the first step where the factory executed code based on content it didn't control.
A factory operating on a public repository has to assume attacker-controlled input: every issue, pull request, comment, and the links inside them are untrusted. Because successful open-source projects are high-value targets, threats range from malicious code changes and supply chain attacks to resource exhaustion, API key exfiltration, and prompt exfiltration.
Sandboxes are the foundation of our defense. Every agent in ai-sdk-factory runs inside an isolated Vercel Sandbox containing its code, its runtime, and only the secrets the agent's specific task needs. With those guardrails, untrusted content can shape what an agent proposes, but the damage any one task can do is contained to the sandbox.
We also built a shielding layer around the sandbox that controls what agents can reach over the network, blocking the paths an attacker would use to pull secrets out of the isolated environment.
The last line of defense is human review: nothing is merged without approval from a human on the AI SDK team.
Copy link to heading Build value locally, then move to the cloud
The first several agents we built for the factory ran through a local CLI. That enabled our team to iterate quickly as we noticed inaccuracies, felt friction, and prototyped different ideas.
Once multiple steps were running reliably through the CLI, we were ready to move the system onto managed infrastructure. ai-sdk-factory uses:
Vercel Functions for the API, workers, and webhook ingress
Vercel Queues for task execution
Vercel Sandbox for the agent workspaces
Neon Postgres for factory data
Today, GitHub webhooks feed the issue queue, and as soon as they arrive workers automatically pull them and kick off agent runs in sandboxes. We also built a monitoring UI that tracks every run in parallel, and visualizes the queue for the team of reviewers.
Copy link to heading Shipping software through our factory
On July 24, a community member asked for blocked-domain support in OpenAI web search, and the request became issue #17898 . The following section explains every step the software factory went through to process the issue, open a pull request with an implementation of the feature, and ultimately backport the merged feature to v5 and v6 of the SDK.
Copy link to heading Classification The first agent that runs in the factory classifies issues and pull requests. In this case, the ai-sdk-factory bot commented on the issue and applied a label, identifying the type as Feature with high confidence. The agent included its rationale for the classification in the comment.
Copy link to heading Analysis After classification, the analysis agent runs. Agents in the factory don't make any assumptions about the technical validity of a feature request or bug, so the analysis agent wrote a probe to confirm the absence of support for blocked domains. issue-17898-type-probe.ts was generated and run, and failed with an error when it looked for blocked-domains and didn't find it.
The failing probe proved the feature was missing on main , and the agent included it in the analysis as evidence.
The analysis agent then used the results of its investigation to build out a spec for the feature: add an optional blockedDomains filter to the existing web-search tool and map it to the provider's blocked_domains field.
The agent also confirmed the spec fit the SDK's provider-adapter architecture and was backward compatible, and even scoped documentation changes.
Copy link to heading Implementation Another agent then implemented the spec and opened pull request . The implementation agent ran a live end-to-end test, executing an OpenAI web search with wikipedia.org blocked and confirming the domain wassn't reachable. The test was included as additional evidence on the pull request.
Copy link to heading Automated review Next, a review agent scored the change , and when it didn't find any concerns, approved it. The agent rated the feature as fully implemented, with:
Backwards-compatibility risk: low
Copy link to heading Human review Finally, Lars read the chain of evidence from the agents, reviewed the code changes, and merged PR #18033 into main.
Copy link to heading Backporting Once Lars merged the initial PR, ai-sdk-factory opened additional PRs for two backports, #18035 for v6 and #18036 for v5. The v5 backport did not apply cleanly, so the factory agent labeled and committed the conflicted state, identified and validated a fix, and pushed it seventeen minutes later. After review, Lars merged both backport PRs.
We are just over four weeks into running the software factory in production. Here are the results:
Between 25-35% of the PRs we merge on a weekly basis are now authored by ai-sdk-factory agents.
Factory PRs are above 50% of weekly merges to the v6 release line, and v5 looks similar. Backports used to be work we skipped because dealing with merge conflicts wasn't worth the effort. v5 and v6 get far better support now.
In July, over 75% of closed issues were closed by the factory.
Open issues fell from a peak of 1,022 in late June to 844 by early August, and open bugs are down roughly 25%.
The ai-sdk-factory runs in public, so you can see every pull request it has authored on the repo.
Copy link to heading Improving the factory becomes the job
The most interesting part of running the factory is what happens when it fails. Every run ends one of four ways: success, flawed, blocked, or manual. Only success ships, so the rest are signal that re-enters the system as feedback.
A flawed run means an agent produced the wrong thing, and the fix is better prompts, better context, or a new eval case so the same mistake gets caught automatically next time
A blocked run means the environment was missing something, like a credential, a service, or a dependency, and the fix is provisioning it
A manual run marks a boundary we drew on purpose, and forces us to ask whether improvements in the factory justify removing it
Each of those fixes expands the automation boundary, and every week the factory can handle work that it

[truncated]
