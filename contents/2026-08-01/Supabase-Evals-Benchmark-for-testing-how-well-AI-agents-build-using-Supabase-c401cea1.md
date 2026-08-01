---
source: "https://supabase.com/blog/introducing-supabase-evals"
hn_url: "https://news.ycombinator.com/item?id=49136928"
title: "Supabase Evals: Benchmark for testing how well AI agents build using Supabase"
article_title: "Introducing Supabase Evals"
author: "maxloh"
captured_at: "2026-08-01T18:55:14Z"
capture_tool: "hn-digest"
hn_id: 49136928
score: 1
comments: 0
posted_at: "2026-08-01T18:19:41Z"
tags:
  - hacker-news
  - translated
---

# Supabase Evals: Benchmark for testing how well AI agents build using Supabase

- HN: [49136928](https://news.ycombinator.com/item?id=49136928)
- Source: [supabase.com](https://supabase.com/blog/introducing-supabase-evals)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T18:19:41Z

## Translation

タイトル: Supabase Evals: Supabase を使用して AI エージェントがどの程度うまく構築されるかをテストするためのベンチマーク
記事のタイトル: Supabase Evals の紹介
説明: AI コーディング エージェントが Supabase でどの程度うまく構築されるかを示すオープンソース ベンチマーク。

記事本文:
Supabase Evals の紹介 コンテンツへスキップ 製品
107.3K サインイン プロジェクトを開始 メイン メニューを開く ブログ / 製品 Supabase Evals の紹介
本日、私たちは supabase/evals をオープンソース化します。これは、Supabase を使用して AI エージェントがどの程度適切に構築されるかをテストするためのベンチマークおよびフレームワークです。実際の Supabase タスク (たとえば、スキーマの構築、失敗したエッジ関数のデバッグ、壊れた RLS ポリシーの修正など) に対して、Claude Code、Codex、OpenCode などのコーディング エージェントを実行し、それらのパフォーマンスをスコア化します。これは、公開されたベンチマークと、毎日監視している内部回帰スイートの両方を強化します。 Supabase プロジェクトを手動ではなくエージェントを通じて出荷する人が増えているため、私たちは推測ではなくそのエクスペリエンスを測定する方法を必要としていました。
ここで引用した結果は、執筆時点のスナップショットを反映しています。 AI ツールは急速に変化するため、現在の結果についてはライブ ページを確認してください。
エージェントは、人々が Supabase で構築する主な方法になりつつあり、CLI、MCP サーバー、エージェント スキル、ドキュメントを通じて対話します。私たちは、1 つのツールだけを単独で使用するのではなく、これらすべての Supabase サーフェスにわたってエージェントがどのように動作するかを理解する方法を必要としていました。エージェントがつまずいた箇所を追跡することで、意図的な修正を加え、それが後退するのではなく維持されていることを確認し、出荷前に自信を持って新機能をテストできます。
私たちはまず、カバーしたい次のような次元を定義しました。
データベースや認証などの製品分野
SDK や可観測性など、エージェントがあらゆる製品にわたって適切に処理する必要があるトピック
アプリの構築や問題の解決など、Supabase ビルダーの作業の段階
次に、サポート チケット、バグ レポート、GitHub の問題などの実際の問題に基づいて、各側面に少なくとも 1 回は影響する最小のシナリオ セットを選択しました。
シナリオをベンチマークと回帰の 2 つのスイートに分割します。
ベンチマーク

シナリオは幅を広げることを目指しており、実際の Supabase ユーザー ジャーニーをカバーする、小さいながらも多様なシナリオのセットとして機能します。これらを、より高機能なモデルと小規模なモデルの組み合わせを含むいくつかのハーネス構成に対して実行し、結果をサイトで公開します。
回帰シナリオは、ベンチマーク スコアに影響を与えることなく、より頻繁に監視する特定の既知の障害モードをカバーすることで、深度を高めることを目的としています。これにより、公開された結果を歪めることなく、バグレポートを優先順位付けしたり、新機能を実験したりできるスペースが得られます。
すべてのシナリオは実際の Supabase 環境に対して実行され、エージェントが実際に実際に行う動作を測定します。私たちは、ホスト型の Supabase スタックとコンテナ内のローカル CLI プロジェクトの両方を起動するフレームワークを構築しました。そのため、エージェントは実際の MCP サーバーと CLI を呼び出します。決定論的なチェック (ユーザーが特定のデータにアクセスできるかどうか、またはエッジ関数が予期した結果を返すかどうか) と、セマンティックな判断が必要なものに対する LLM を判断者として組み合わせて、ユーザーの動作をスコアリングします。持続可能な実行を維持しながら誤検知を減らすために、エージェントは失敗後に評価する前に 1 回再試行できるようになります。新しい変更やハーネスを評価するときにベンチマーク シナリオを実行し、回帰結果を毎日更新します。
ベンチマーク結果は Web アプリで視覚化されるため、誰でもスコアを参照してフィルタリングしたり、各実行の詳細を確認したりできます。
以下は、ベンチマーク開発中にエージェントがつまずくことが確認されたいくつかの領域と、それらにどのように対処しているかを示しています。
私たちのベンチマーク全体で、エージェントはスキルをまったくロードせずにほとんどのシナリオをすでに通過しています。たとえば、Build ステージでは、Opus 5 と Kim K3 は両方ともスキルがロードされていない状態で 100% のスコアを獲得しました。スキルが残りの差を縮めました。Sonnet 5 は 78% から 100%、GPT-5.6 Sol は 89% から 100%、GPT-5.4 mini は 78% から 89% になりました。つまり、エージェントは合理的に ca

すぐに広範な Supabase タスクに対応できます。私たちが目にした最大の効果は、スキルを積んだエージェントが Supabase ドキュメントをより一貫して徹底的にチェックしたことです。これにより、モデルが古いトレーニング前の知識を取り除く必要があるエッジケースでスキルを維持できるようになりました。当社のスキルを使用して、エージェントが Supabase のベスト プラクティスと重大な変更に関する最新情報を確実に入手できるようにすることをお勧めします。その例を以下に示します。
宣言型スキーマは、複数の移行ファイルにわたって推論するのではなく、データベースの形状を 1 か所で記述する開発者に優しい方法です。私たちは、エージェントが一連の移行からスキーマをつなぎ合わせるのではなく、その単一の信頼できる情報源からスキーマを管理することを好むことを期待しています。しかし、すでに宣言型スキーマを使用しているプロジェクトであっても、エージェントが代わりに移行を手書きで作成しようとしていることに気付きました。その結果、スキル ガイダンスを更新して、各ワークフローを選択するタイミングをより明確にし、評価で動作の修正を検証しました。
私たちは最近、安全なエッジ関数の作成の定型文を簡素化するために @supabase/server をリリースしました。ただし、エッジ関数を記述するタスクを任された場合、エージェントは依然として supabase-js を使用して手動で認証を検証することを選択します。その後、エージェントの参考として、各パッケージをいつ使用するかを明確にする「どのパッケージを選択するか」ガイドを公開しました。
エージェントがセッション中にロードしたスキルと、利用可能なスキルを比較して追跡します。私たちのメインのスーパーベース スキルは、利用可能なすべてのシナリオにロードされます。 Postgres のベスト プラクティスをカバーする以前のスキルは、当初は同じセッションの約 10 分の 1 にしかロードされませんでした。私たちはそのスキルの説明を書き直して、より明確なトリガーでリードするようにし、ベンチマーク全体でアクティベーションを 60% に増やしましたが、それでも OpenAI モデルの方がアクティベートの信頼性が高いと感じています。

スキル。
実行中、エージェントがいつドキュメントを読んだかを、Supabase MCP search_docs ツールまたはネイティブ Web ツールを介して追跡します。 Codex ベースのエージェントは、Claude Code エージェントよりも頻繁に文書をチェックします。より有能な OpenAI モデルはドキュメントを最も一貫してチェックし、Codex / GPT-5.6 は最も多くのページを読み取ります。シナリオごとに約 8 ページですが、クロード コードの場合は約 2 ページです。 Claude Code は、スキルがロードされている場合でも、シナリオの 40% 未満でドキュメントをチェックします。私たちは、エージェントが必要な状況でドキュメントを確認し、適切な意思決定を行うために必要な情報をすべて見つけられるように取り組んでいます。
supabase.com/evals でベンチマーク結果を参照し、製品、ステージ、またはエージェントごとにグループ化して、どこのエージェントが強力で、どこが弱いかを確認します。または、ベンチマークと回帰シナリオの詳細については、GitHub リポジトリを参照してください。
これが出発点です。この問題領域内のエッジケースのカバー範囲を拡大し、エージェントや製品の変更に応じて新しいシナリオを追加します。また、スコアをより厳密かつ安定させるために取り組んでいます。特定の回帰シナリオに自信が持てるようになると、それらのシナリオを公開ベンチマークに段階的に組み込む可能性があります。最後に、私たちは、エージェントが困難に直面したときにフィードバックを送信できるようにする新しい CLI コマンドと MCP ツールの開発に取り組んでいます。これは、次の作業の優先順位付けに役立ちます。
週末に構築して数百万人にスケールアップ
Twitter GitHub Discord Youtube TikTok Instagram Supabase から製品の最新情報やニュースを入手してください。

## Original Extract

Our open-source benchmark for how well AI coding agents build with Supabase.

Introducing Supabase Evals Skip to content Product
107.3K Sign in Start your project Open main menu Blog / product Introducing Supabase Evals
Today we're open sourcing supabase/evals , our benchmark and framework for testing how well AI agents build using Supabase. It runs coding agents including Claude Code, Codex, and OpenCode against real Supabase tasks, for example, building a schema, debugging a failed Edge Function, or fixing a broken RLS policy, and then scores how well they performed. It powers both our published benchmark and an internal regression suite we monitor daily. As more people ship Supabase projects through an agent instead of by hand, we wanted a way to measure that experience instead of guessing.
The results cited here reflect a snapshot at the time of writing. AI tooling changes rapidly, so check the live page for current results.
Agents are becoming a primary way people build with Supabase, interacting through our CLI, MCP server, agent skills, and docs. We needed a way to understand how agents perform across all of these Supabase surfaces, not just one tool in isolation. By tracking where agents stumble, we can make intentional fixes and verify they hold instead of regressing, and test new features confidently before shipping.
We started by defining the dimensions we wanted to cover, including:
Product areas like Database and Auth
Topics agents need to handle well across any product, like the SDK and observability
Stages of a Supabase builder's journey, such as building an app or resolving issues
Then we picked the smallest set of scenarios that touched each dimension at least once, grounded in real problems like support tickets, bug reports, or GitHub issues.
We split scenarios into two suites, benchmark and regression:
Benchmark scenarios aim for breadth, acting as a small but diverse set of scenarios that cover the real Supabase user journey. We run these against several harness configurations, including a mixture of more capable and smaller models, and publish results on our site .
Regression scenarios aim for depth, covering specific known failure modes that we monitor more frequently without influencing benchmark scores. This gives us a space to triage bug reports or experiment with new features without skewing our published results.
Every scenario runs against a real Supabase environment to measure what an agent would actually do in the wild. We built a framework that spins up both a hosted-like Supabase stack and a local CLI project in containers, so agents invoke our actual MCP server and CLI. We score their behavior with a combination of deterministic checks (whether a user can access certain data, or an Edge Function returns an expected result) and LLM-as-a-judge for anything that needs semantic judgment. To reduce false negatives while keeping runs sustainable, we let agents retry once after a failure before grading them. We run benchmark scenarios when assessing new changes or harnesses, and refresh regression results daily.
Benchmark results are visualized in a web app, so anyone can browse and filter the scores or see details of each run.
Below are some areas we saw agents tripping up during benchmark development, and how we're addressing them:
Across our benchmark, agents already pass most scenarios with no skill loaded at all. In the Build stage, for example, Opus 5 and Kimi K3 both scored 100% with no skill loaded. Skills closed the gap for the rest: Sonnet 5 went from 78% to 100%, GPT-5.6 Sol from 89% to 100%, and GPT-5.4 mini from 78% to 89%. That means agents are reasonably capable at broad Supabase tasks out of the box. The biggest impact we saw is that agents with our skills loaded checked Supabase docs more consistently and thoroughly, which helped skills earn their keep in the edge cases where models need to unlearn outdated pre-training knowledge. We still recommend using our skill to ensure your agents have the most up-to-date information on Supabase best practices and breaking changes, examples of which follow below.
Declarative schemas are a developer-friendly way of describing the shape of your database in one place instead of reasoning across several migration files. We expect agents to prefer managing schemas from that single source of truth instead of piecing it together from a chain of migrations. But even in a project that already uses declarative schemas, we noticed agents try to hand-write migrations instead. As a result, we updated our skill guidance to make it clearer when to pick each workflow and verified the behavior correction with our evals.
We recently released @supabase/server to simplify boilerplate of writing secure Edge Functions. When tasked with writing edge functions though, agents still choose to verify auth by hand with supabase-js . We've since published a "Which package to choose" guide to clarify when to use each package, as a reference for agents.
We track which skills agents loaded during a session compared to what was available. Our main supabase skill loads in every scenario where it's available. Our earlier skill, covering Postgres best practices , originally loaded in only about one in ten of those same sessions. We rewrote our description of that skill to lead with clearer triggers, increasing activation to 60% across our benchmarks, though we still find OpenAI models more reliable at activating the skill.
During runs we track when agents read our docs, either via the Supabase MCP search_docs tool or their native web tools. Codex-based agents check docs more than Claude Code agents do. The more capable OpenAI model checks docs most consistently, and Codex / GPT-5.6 reads the most pages, around 8 per scenario versus about 2 for Claude Code. Claude Code checks our docs in under 40% of scenarios even with skills loaded. We're working to make sure agents are checking docs in the situations they need to, and that they find all the information they need to make good decisions.
Browse the benchmark results at supabase.com/evals and group by product, stage, or agent to see where agents are strong and where they're not. Or see the GitHub repo for details on the benchmark and regression scenarios.
This is a starting point. We'll expand coverage of edge cases within this problem space, and add new scenarios as agents and our product change. We're also working to make our scoring more rigorous and stable. As we build confidence in specific regression scenarios, we may graduate them into the published benchmark. Lastly, we're working on a new CLI command and MCP tool that will allow agents to submit feedback when they struggle, which will help us prioritize what's next.
Build in a weekend, scale to millions
Twitter GitHub Discord Youtube TikTok Instagram Get product updates and news from Supabase.
