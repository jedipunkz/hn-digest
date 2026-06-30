---
source: "https://www.moderntreasury.com/journal/reducing-feedback-latency-with-local-ci-for-developers-and-ai-agents"
hn_url: "https://news.ycombinator.com/item?id=48735105"
title: "Reducing Feedback Latency with Local CI for Developers and AI Agents"
article_title: "Reducing Feedback Latency with Local CI for Developers and AI Agents"
author: "mattmarcus"
captured_at: "2026-06-30T16:42:54Z"
capture_tool: "hn-digest"
hn_id: 48735105
score: 3
comments: 0
posted_at: "2026-06-30T16:29:16Z"
tags:
  - hacker-news
  - translated
---

# Reducing Feedback Latency with Local CI for Developers and AI Agents

- HN: [48735105](https://news.ycombinator.com/item?id=48735105)
- Source: [www.moderntreasury.com](https://www.moderntreasury.com/journal/reducing-feedback-latency-with-local-ci-for-developers-and-ai-agents)
- Score: 3
- Comments: 0
- Posted: 2026-06-30T16:29:16Z

## Translation

タイトル: 開発者と AI エージェント向けのローカル CI によるフィードバックの遅延の削減
説明: クラウド CI の高速化により、コードをより迅速にマージできるようになりましたが、開発中にフィードバックを待つという問題は解決されませんでした。 CI を開発者マシンに戻すことで、コンテキストの切り替えを減らし、エージェントのワークフローを改善し、反復を大幅に高速化することができました。

記事本文:
開発者と AI エージェント向けのローカル CI によるフィードバックの遅延の削減 最新の財務製品
6 月 25 日の Spritz にご参加ください — 法定通貨とステーブルコインが実際の製品でどのようにシームレスに機能するかをご覧ください。登録はこちら→
開発者と AI エージェント向けのローカル CI によるフィードバックの遅延の削減
クラウド CI の高速化により、コードをより迅速にマージできるようになりましたが、開発中にフィードバックを待つという問題は解決されませんでした。 CI を開発者マシンに戻すことで、コンテキストの切り替えを減らし、エージェントのワークフローを改善し、反復を大幅に高速化することができました。
問題: クラウド CI の待ち時間が長い 解決策: ローカル CI の高速化 フィードバックの高速化 コンテキストの切り替えの削減 LLM ハーネスの実現 これを実現する bin/setup によるローカル環境の標準化 クラウドとローカルで同じように動作するように個別の手順を統合する CI ランナーの選択 bin/ci DSL のロールアウト フィルターの実行をサポートするための CI DSL の拡張 クラウドでの bin/ci の導入 bin/autofix による事前の自動修正 警告 要点AI OpenAI で開く ↗ クロードで開く ↗ マークダウンとしてコピー コピーされたトピック エンジニアリング 私たちは、クラウドの継続的インテグレーション (CI) の実行を高速化するために多くの時間を費やしてきました。キャッシュの改善、マシンの高速化、DAG の並べ替え、遅いステップの排除などの改善により、フィードバックの取得と新しいコードのマージにかかる時間が約 5 ～ 7 分に短縮されました。しかし、これだけの作業を行った後でも、ローカル開発中のフィードバック ループは依然として遅すぎると感じました。 Cloud CI には依然として、各プル リクエストの環境のプロビジョニング、パイプライン ステージ間の一時停止、レポート インフラストラクチャ、分散システムの調整コストによるオーバーヘッドがかかり、結果に基づいて行動できるようになるまでに各レイヤーでレイテンシが追加されます。
プル リクエストを作成した後、開発者は多くの場合、クラウド CI に切り替えるまでに十分な待ち時間が必要でした。

実行中に別のタスクが発生しました。 CI が完了しても、通常はすぐには戻りません。代わりに、新しいタスクが自然な停止点に達するまで待ってから、コンテキストを元に切り替えて、CI によって発生したエラーに対処します。
実際には、これは、大部分の作業が完了した後でも、すべてのエンジニアの頭の中にいくつかの異なるコードの変更が残ることを意味しました。 LLM が関与している場合、各変更間の待ち時間が長くなるため、反復速度は非常に遅くなります。
最近の開発者マシンは非常に強力になっています。たとえば、私がこれを書いている M4 Macbook Pro には 14 コアと 48 ギガバイトの RAM が搭載されており、単一のコミットに使用される CI クラスターと同等のリソースが得られます。これは最終的に、小規模プロジェクトのテスト スイート全体を実行し、メイン モノリスのほぼすべてのステップを非常に迅速に実行できるほど強力になりました。
これにより、さまざまな領域のオーバーヘッドが削減されました。
セットアップ: CI ループをローカルに導入すると、開発者が機能構築の一環としてすべてをすでにインストールして実行しているため、セットアップのオーバーヘッドがなくなり、実際のチェックの実行に進むことができます。
ネットワーク: ローカルで実行することで、クラウド CI に固有のネットワークと調整のオーバーヘッドが排除されました。コードを GitHub にプッシュし、GitHub が CI プロバイダーをトリガーするのを待ち、結果が複数のサービスを介して伝播されるのを待つのではなく、開発者はマシン上で直接フィードバックを受け取りました。
ワーカーの割り当て: クラウド CI では、各ステップは、共有クラスターによってプロビジョニングおよび調整された短期間のワーカー上で実行されました。これにより、ステップ間に数秒のオーバーヘッドが発生し、同じインフラストラクチャを競合する他のビルドからのリソース競合によって発生する時折の遅延も発生しました。
迅速なローカル フィードバックにより、自信を持ってチャットできるようになるまでクラウド CI を待つ必要がなくなりました。

ンジ。ローカル CI が合格した場合、通常はクラウド CI も合格します。この高速フィードバックにより、開発者と LLM の両方にとって、より緊密な反復ループが可能になりました。以前は、誰かが変更を加えた後は関連するテストのみが実行されていましたが、現在では変更のたびに CI 全体が実行されるようになり、問題をより迅速に検出できるようになりました。
CI を実行するだけで結果がすぐに確認できるということは、クラウドの実行が完了するのを待たずに問題をすぐに修正できることを意味します。フィードバックを即座に取得できるため、開発者にとっても、LLM が関与するワークフローにとっても、イテレーションが高速化されました。
クラウド CI の最適化に多大なエンジニアリング労力を費やしたにもかかわらず、競合するリソース、コンテナーの初期化、コンピューティング割り当てなどによるオーバーヘッドが削減されたため、ローカル実行のほうが結果的に高速化されました。
その時点で、場合によっては、広範なテスト スイート (RSpec および Jest で記述された) に対して LLM をハーネス ループとして動作させることができましたが、コードが正常に動作しているように見えても、自動生成された型が欠落しているか、リンター ルールが欠落しているか、ライセンス チェックが失敗したために、後でクラウド CI に失敗するという状況が依然として発生していました。
ローカル CI を使用することで、失敗する可能性のあるほぼすべてのステップにハーネスを拡張できるようになり、ビルド全体が成功するまで LLM を反復できるようになりました。これは正しいコードを持っていることを保証するものではありませんが、動作する、準拠したコードがあることは保証しました。
より複雑な権限を持つ MCP や API を通じてこれを有効にすることもできますが、エージェントがスクリプトを実行し、出力を解析し、結果を反復処理できるようにする方が、統合がはるかに簡単でした。
これらのステップの多くは、DevX または AgentX の改善に向けてすでに実行されていたため、その即時的なプラスの影響は、完全なローカル CI から生じるメリットと組み合わされました。
地域環境の標準化

ビン/セットアップ付き
Ruby on Rails には、新規または既存の環境に対して 1 つのコマンドで一貫した開発環境を提供する優れたコマンド bin/setup が備わっています。
まず、開発者がすべてのローカル開発依存関係を管理する $ setup を毎日実行するだけで済むようにすることを目標に、既存のオンボーディング スクリプトを冪等のステップを含む bin/setup パターンに変換しました。
次に、クラウド内にのみ存在するチェッカーをローカルの依存関係として徐々に追加し、すべてがローカルでも利用できるようにしました。開発者が bin/setup を時々実行すると、ローカルの開発環境はローカル CI を機能させるために必要なものをすべて取得しました。
結果の bin/setup スクリプトでは、Gemfile、package.json、Brewfile、および特定の Docker イメージ上に構築するとともに、プロジェクト固有の依存関係のリストを組み合わせて使用​​し、すべてのマシンで一貫した環境を確保しました。これは、ローカル開発セットアップとオンボーディングを管理する開発者の煩わしさを軽減するために作成されましたが、すべてのマシンに対するこの一貫した環境により、ローカル CI の前提条件が確立されました。
個々の手順を統合してクラウドとローカルで同じように作業する
クラウド CI ステップの多くはクラウド ビルド環境向けにのみ書かれていますが、これを成功させるには、すべてのステップがローカルでも機能する必要がありました。ローカル環境がクラウド環境と統合されるにつれて、ローカル コマンドの作成または書き換えを開始しました。
既存のスクリプトの場合: クラウド ステップは、必要に応じてシステム固有の癖に対応したブランチを備えた、より汎用的な *nix 形式で書き直す必要がありました。ほとんどの場合、これらは非常に迅速に完了しました。
CI 拡張機能の場合: チェックを実行するサードパーティ拡張機能を介して実行される多くの手順がありました。これらをすべてのシステムで動作する単純なスクリプトに書き直しました。
当初の動機は、開発者の再開発を支援することでした。

クラウド CI 障害をローカルで生成し、より迅速に反復します。その過程で、ほぼすべての CI ステップをローカルで利用できるようにし、新しい方法でそれらを構成するための基盤を作成しました。
ローカル設定に一貫性があれば、DSL を実行する方法を決定することができます。私たちが最初に直感したのは、クラウド CI 構成にローカル ランナーを使用することでした。しかし、ファーストパーティのサポートは初期段階を超えてパイプラインを現状のままサポートすることはなく、そのためのサードパーティ ライブラリも行き過ぎることはありませんでした。そこで私たちは彼らの生態系の外を探索することにしました。私たちが探したのは:
最小限のセットアップで、理想的には箱から出してすぐに動作します
概要レポート - 迅速な反復のために失敗と成功の概要を表示します。
状況は確実に異なるため、評価の一部として次の点を必ず確認してください。
タスクファイル - 既存のクラウド CI 構成と比較すると、構文と機能は非常によく知られており、クロスプラットフォームのサポートは良好に見えましたが、新しいローカル依存関係を導入するのは好きではありませんでした。同様に、ローカルの Gitlab CI および GitHub Action ランナーも除外しました。
すでに利用可能なものに目を向ける
シェル スクリプト - これらは開発者にとって馴染みがあり、各ステップで結果が出力され (set -x)、非常に簡単に始めることができるため、見栄えがよくなります。私たちの CI ロジックについては、かなり複雑になり始め、すぐに使える優れた概要レポートがありませんでしたが、最初は単純なものを構築しようとしました。ビルドが非常に簡単な場合、これは素晴らしい選択です。たとえば、 dhh/signoff を確認してください。
Rake - Ruby の Make-like タスク ランナーは、すでに Rails にバンドルされているため非常にうまく機能し、大規模なシェル スクリプトよりも読みやすいシンプルな構成を作成できましたが、それでも独自のレポートの概要を追加する必要がありました。

コードを作成し、依存関係を調整します。プロジェクトが完了した後、このケースをサポートするためのきちんとした構文で Rake を拡張する HellRok/LocalCI という素晴らしいコミュニティ プロジェクトを見つけました。
Rails bin/ci コマンド - 最近、Rails 8.1 では、レポートのニーズをカバーし、きちんとした bin/ci コマンドでほぼすべてのユースケースをカバーする構文を備えた、すぐに使える DSL が導入されました。最初に見た他のオプションほど具体化されていませんでしたが、私たちのケースをカバーする際に、箱から出してすぐに使用できました。
CI.実行する
ステップ「スタイル：Ruby」、「bin/rubocop」
ステップ「チェック: TODO がありません」、「if grep -r TODO app/; then exit 1; fi」
実行するには、次のとおりです。
$ ci bin/ci DSL の展開
CI ランナーの選択を最終的に決定した後、移行するプロジェクトがいくつかありました。私たちは、単純な CI チェック パイプラインから始まったばかりの最近の Rails プロジェクトの 1 つから始めました。 Rails 8.1 は最新だったので、最新の DSL で構成を作成しました。
初期フィードバックは良好で、いくつかの小さなバグが修正され、安定しているように見えたので、他のプロジェクトへの展開を検討し始めました。古いプロジェクトの多くはまだ最新の Rails 8.1 ではなかったため、バージョンをアップグレードせずにアクセスできるように、Rails から ActiveSupport::ContinuousIntegration ファイルを内部的に組み込みました。私たちは一度に 1 つのプロジェクトを実行し、最小のものから最大のものまで作業を進めました。ただし、最大のものは、そのサイズのために追加のサポートが必要でした。
フィルタの実行をサポートするための CI DSL の拡張
メインのモノリスを移行する段階になったとき、多数のテストとチェックのため、プラットフォーム CI の実行に適応する必要があり、数時間かかりました。ローカル CI はクラウド システムに存在するオーバーヘッドの多くを排除し、大幅に高速に実行されましたが、それでも十分な速度ではありませんでした。できる方法が必要でした

一度に 1 つのサブセットを実行します。
ContinuousIntegration を組み込むことで得られた嬉しい利点は、DSL を自分たちで拡張する余地が得られたことです。この目的のために、グループの簡単な構文と、特定のグループを実行するための --group/-g コマンド ライン オプションを追加しました。
CI.実行する
グループ「バックエンド」が行うこと
ステップ「スタイル：Ruby」、「bin/rubocop」
終わり
グループ「フロントエンド」が行うこと
ステップ「Biome: TypeScript」、「bin/biome」
終わり
一般に、開発者/LLM はどのサブセットを実行するかを判断でき、完全な分散クラウド CI の実行により、すべてが包括的に実行されます。開発者が変更に関連する CI の部分を実行して迅速に反復できるため、それでもかなりのアップグレードでしたが、クラウド構築が失敗する可能性は依然としてほぼゼロになりました。
これは非常に一般的に便利だったので、これをアップストリームしようとしました。
当初、クラウド CI 構成を定義する YAML ファイルはそのままにして、ローカル CI をオプションの追加ワークフローにしました。ただし、これは CI ロジックに関して 2 つの信頼できる情報源を維持することを意味し、開発者が一方の構成を更新しても他方の構成を更新しないというリスクが生じ、時間の経過とともに構成がずれてしまいます。また、統合により、開発者と LLM エージェントは、ローカル CI を通過すればクラウド CI も通過できるという確信を得ることができました。
プロジェクトごとに、クラウド CI の手順を置き換えました。

[切り捨てられた]

## Original Extract

Faster cloud CI helped us merge code more quickly, but it didn't solve the problem of waiting for feedback while developing. By bringing CI back to developer machines, we were able to reduce context switching, improve agent workflows, and make iteration substantially faster.

Reducing Feedback Latency with Local CI for Developers and AI Agents Modern Treasury Products
Join us June 25th with Spritz — see how fiat and stablecoins work seamlessly in a real product. Register Here →
Reducing Feedback Latency with Local CI for Developers and AI Agents
Faster cloud CI helped us merge code more quickly, but it didn't solve the problem of waiting for feedback while developing. By bringing CI back to developer machines, we were able to reduce context switching, improve agent workflows, and make iteration substantially faster.
Problem: Long Cloud CI Waits Solution: Local CI Faster Feedback Less Context Switching LLM Harness Making This Happen Standardizing Local Environment with bin/setup Unifying Individual Steps to Work the same in cloud and local Selecting the CI Runner Rolling out the bin/ci DSL Extending CI DSL to support filter running Adopting bin/ci in the Cloud Auto fixing upfront with bin/autofix Caveats Takeaways Explore With AI Open in OpenAI ↗ Open in Claude ↗ Copy as Markdown Copied Topics Engineering We’ve spent a lot of time speeding up our cloud Continuous Integration (CI) runs. Improvements like better caching, faster machines, DAG reordering, and eliminating slow steps reduced the time to get feedback and merge new code to roughly 5–7 minutes. Yet even after all of that work, the feedback loop still felt too slow during local development. Cloud CI was still carried overhead from provisioning environments for each pull request, pauses between pipeline stages, reporting infrastructure, and the coordination costs of a distributed system, with each layer adding latency before we can act on the results.
After making any Pull Request, developers often faced a long enough wait for cloud CI that they would switch to another task while it ran. Even when CI completed, they typically wouldn't return immediately; instead, they would wait until the new task reached a natural stopping point before context-switching back to address any errors CI had surfaced.
In practice, this meant having a few different code changes lingering in every engineer’s mind even after the bulk of work was done. If LLMs were involved, their rate of iteration was quite slow as it waited longer between each change.
Recent developer machines have gotten very powerful For example, the M4 Macbook Pro I’m writing this on has 14 cores and 48 Gigabytes of RAM, giving it comparable resources to our CI cluster used for a single commit. This ended up powerful enough to run the entire test suite for our smaller projects, and run almost all our steps of our main monolith quite quickly.
This reduced overhead across various areas:
Setup: Bringing the CI loop locally removed setup overhead as a developer already had everything installed and running as part of the building out the feature, letting us skip to running the actual checks.
Network: Running locally eliminated the network and coordination overhead inherent in cloud CI. Rather than pushing code to GitHub, waiting for GitHub to trigger our CI provider, and then waiting for results to propagate back through multiple services, developers received feedback directly on their machines.
Worker allocation: In cloud CI, each step ran on a short-lived worker provisioned and coordinated by a shared cluster. This introduced several seconds of overhead between steps, along with occasional delays caused by resource contention from other builds competing for the same infrastructure.
Fast local feedback meant we no longer had to wait for cloud CI before feeling confident in a change. If local CI passed, cloud CI would typically pass as well. This faster feedback allowed for tighter iteration loops for developers and LLMs alike. In the past, only the relevant test would run after someone made a change, but now the entire CI was run after every change, catching issues faster.
Being able to just run CI and see results right away meant being able to fix issues immediately instead of waiting for a cloud run to finish. Getting immediate feedback led to faster iterations, both for developers and for workflows involving LLMs.
Despite spending considerable engineering effort optimizing our cloud CI, local runs still ended up being faster due to reduced overhead from competing resources, container initialization, compute allocation, and so on.
At that point, we were already having LLMs work against our extensive test suite (written in RSpec and Jest) as a harness loop in some cases, but we would still encounter situations where code appeared to work fine, only to later fail cloud CI because an autogenerated type was missing, a linter rule missed, or a license check had failed.
Having a local CI allowed us to expand our harness to practically every step that could fail, allowing the LLM to iterate till an entire build passed. While this didn't guarantee we had the correct code, it did guarantee we had working and compliant code.
While it's also possible to enable this through MCPs and APIs with more complex permissions, allowing an agent to simply run a script, parse the output, and iterate on the results was much easier to integrate.
Many of these steps were already in motion for their DevX or AgentX improvements, so their immediate positive impacts were also combined with the benefits that arose from a complete local CI.
Standardizing Local Environment with bin/setup
Ruby on Rails has an excellent command bin/setup that provides a consistent development environment in a single command for a new or pre-existing environment.
First, we converted our pre-existing onboarding scripts to the bin/setup pattern with idempotent steps, with a goal that developers simply just run $ setup daily that manages all the local development dependencies.
Then we progressively added our checkers that were only in the cloud as local dependencies so that everything was also available locally. As developers ran bin/setup occasionally, the local dev environment gained everything required to make local CI work.
The resulting bin/setup script used a mix of a list of project specific dependencies alongside building on top of Gemfile, package.json, Brewfile , and specific Docker images that ensured a consistent environment on every machine. While we created this to reduce developer annoyance for managing their local development set up and onboarding, this consistent environment per all machines established the pre-requisite for local CI.
Unifying Individual Steps to Work the same in cloud and local
Many of our Cloud CI steps were only written for our cloud build environment, but for this to succeed we needed every step to also work locally. As the local environment unified with our cloud environment, we started creating or rewriting local commands:
For existing scripts: Cloud steps had to be rewritten in a more generic *nix format, with branches for system specific quirks if needed. In most cases these were quite quick to run through.
For CI extensions: We had many steps that were run through a third party extension that ran the check. We rewrote these into simple scripts that worked on all systems.
The initial motivation was to help developers reproduce cloud CI failures locally and iterate more quickly. In the process, we made nearly every CI step available locally, which created a foundation for composing them in new ways.
Once the local set up was consistent, we had the choice to figure out how to run our DSL. Our first instinct was to use a local runner for our cloud CI configuration, but first-party support never got too far beyond early stages and didn’t support our pipeline as-is, and third-party libraries for it never got too far either. So we decided to explore outside their ecosystem. We looked for:
Minimal setup, ideally works out of the box
Summary Reporting - showing a summary of failures and successes for quick iteration
Your circumstances will definitely vary, so definitely check these out as part of your evaluation:
Taskfile - The syntax and functionality looked quite familiar compared to our existing cloud CI configuration, and the cross platform support looked good, but we didn’t like introducing a new local dependency. Similarly, we also ruled out local Gitlab CI and GitHub Action runners.
Turning to what we had available already
Shell scripts - These look great because they are familiar to developers, output each step with results (set -x), and quite easy to start with. For our CI logic, it started getting quite complicated and it didn’t have good out of the box summary reporting, though we initially did try building something simple. This is a great choice if your build is quite straightforward. For example, check out dhh/signoff .
Rake - Ruby’s make-like task runner could work quite well since it already came bundled with rails and we were able to create a simple configuration that was more readable than the large shell script, but we still had to bolt on our own reporting summary code and aligning the dependencies. After we finished the project, we did find a nice community project called HellRok/LocalCI that extends Rake with some neat syntax to support this case.
Rails bin/ci command - Recently, Rails 8.1 introduced an out of the box DSL that covered our needs for reporting and had a syntax that covered almost all our use cases with a neat bin/ci command. While it wasn’t as fleshed out as some of the other options we saw initially, it came out of the box while covering our cases.
CI.run do
step "Style: Ruby", "bin/rubocop"
step "Check: No TODOs", "if grep -r TODO app/; then exit 1; fi"
end To run, it was simply:
$ ci Rolling out the bin/ci DSL
After finalizing our choice of CI runner, we had several projects to migrate. We started with one of our recent Rails Projects that had just started with a simple CI check pipeline. Since it was up to date for Rails 8.1, we wrote our configuration in the recent DSL.
We got good initial feedback and fixed a few small bugs, and once that looked stable, we started to explore rolling it out to other projects. Many older projects were not at the latest Rails 8.1 yet, so we internally included the ActiveSupport::ContinuousIntegration file from Rails so that we could access it without upgrading the version. We did one project at a time, working up from the smallest to the biggest one. The biggest one, however, needed some extra support due to its size.
Extending CI DSL to support filter running
When we got to migrating our main monolith, we had to adapt around the platform CI run taking multiple hours due to the large number of tests and checks. Although local CI eliminated much of the overhead present in our cloud system and ran significantly faster, it still wasn't fast enough. We needed a way to be able to run one subset at a time.
A welcome bonus of including ContinuousIntegration was getting room to extend the DSL ourselves. For our purposes, we added a simple syntax for groups and a --group/-g command line option to run specific groups.
CI.run do
group "backend" do
step "Style: Ruby", "bin/rubocop"
end
group "frontend" do
step "Biome: TypeScript", "bin/biome"
end
end Generally a developer/llm could figure out what subset to run, and the full distributed cloud CI run would run everything comprehensively. It was still quite an upgrade since a developer could run the parts of the CI relevant to their change to iterate quickly, so the odds of a cloud build failure still got to almost zero.
This was quite generally useful, so we took an attempt to upstream this.
Initially, we left the YAML files that defined our cloud CI configuration untouched, making local CI an optional, additive workflow. However, this meant maintaining two sources of truth for CI logic, which created the risk of developers updating one configuration but not the other, causing them to drift over time. Unification also gave developers and LLM agents more confidence that passing local CI will pass cloud CI.
Project by project, we replaced cloud CI steps w

[truncated]
