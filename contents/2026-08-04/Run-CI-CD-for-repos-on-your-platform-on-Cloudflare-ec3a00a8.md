---
source: "https://blog.cloudflare.com/ci-workflows/"
hn_url: "https://news.ycombinator.com/item?id=49168483"
title: "Run CI/CD for repos – on your platform, on Cloudflare"
article_title: "Run CI/CD for millions of repos — on your platform, on Cloudflare | The Cloudflare Blog"
author: "aofeisheng"
captured_at: "2026-08-04T13:51:28Z"
capture_tool: "hn-digest"
hn_id: 49168483
score: 1
comments: 0
posted_at: "2026-08-04T13:13:24Z"
tags:
  - hacker-news
  - translated
---

# Run CI/CD for repos – on your platform, on Cloudflare

- HN: [49168483](https://news.ycombinator.com/item?id=49168483)
- Source: [blog.cloudflare.com](https://blog.cloudflare.com/ci-workflows/)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T13:13:24Z

## Translation

タイトル: リポジトリの CI/CD を実行 – プラットフォーム上、Cloudflare 上で
記事のタイトル: 数百万のリポジトリに対して CI/CD を実行する — プラットフォーム上、Cloudflare 上で | Cloudflareのブログ
説明: プラットフォーム上、Cloudflare 上で、数百万のリポジトリに対して CI/CD を実行します。

記事本文:
プラットフォーム上、Cloudflare 上で、数百万のリポジトリに対して CI/CD を実行します。 Cloudflareのブログ
コンテンツへスキップ すべてのカテゴリ AI
ログイン 営業担当者へのお問い合わせ ブログ AI 開発者 プラットフォーム 開発者 +2 さらに 2 つのタグを表示 5 つのタグ 5 つのタグを表示 選択したタグ
AI 開発者 プラットフォーム開発者 製品ニュース ワークフロー
自動プラットフォーム最適化
Cloudflare 1 ユーザーのリスクスコア
プラットフォーム上、Cloudflare 上で、数百万のリポジトリに対して CI/CD を実行します
アンドレ・ヴェンセスラウ、ミア・マルデン、トマーシュ・ホブザ
私たちは、コードを完全に Cloudflare 上に保存、構築、テスト、デプロイできる世界に向かって進んでいます。私たちは、数百万のリポジトリに拡張できるバージョン管理されたコード ストレージである Artifacts を使用して最初の部分を構築しました。
Cloudflare ワークフロー上に構築された CI SDK とストア、ビルド、デプロイのステップを統合して、Cloudflare 上で継続的インテグレーション (CI) パイプラインを実行できるようにしました。アーティファクト プッシュ イベントをワークフローに直接送信し、ラングラー構成ファイルの新しいイベント フィールドを通じてその実行インスタンス (基本的には CI ジョブ) をトリガーできます。
次に、@cloudflare/ci がインストールされたワークフローから直接、次のことができます。
ビルドを自動化する: 安全で隔離された環境で Artifacts リポジトリからコードをコンパイルします
リンターと型チェックを実行します。コード スタイルを適用し、型エラーを検出し、潜在的な問題にフラグを立てます。
依存関係をキャッシュする: インストールを 1 回実行すると、CI ジョブのステップ全体で依存関係をキャッシュします。
単体テストを実行します。コードの各部分が期待どおりに動作することを確認します。
自己修復: AI レビュー エージェントを統合して、ビルド内の壊れたステップを検出し、修正するためにコミットをプッシュします。
条件付きデプロイ: ビルドステップが成功した場合にのみ、コードを自動的にデプロイします。
今日、誰もが社内のバイブコーディングプラットフォームであろうと、あなたの拡張であろうと、プラットフォームを構築しています。

コードによるカスタマイズによる顧客対応製品。プラットフォームは現在、Artifacts 上の何百万ものリポジトリを使用して、自社のコードと顧客のコードを保存し、両者のバージョン管理を行っています。ただし、各チームには継続的な統合と展開パイプラインに対する独自のニーズがあります。プラットフォームの場合、顧客とは異なる独自のコードの CI ジョブを定義したい場合があります。
これらのプラットフォーム上に構築しているエンド カスタマーの多くは、継続的インテグレーションおよび継続的デプロイメント (CI/CD) パイプラインの管理に煩わされることを望んでいません。代わりに、プラットフォームは顧客に代わってビルド プロセスを管理できます。つまり、CI/CD パイプラインを一度作成すれば、顧客が構築しているすべてのアプリケーション間でそれを共有できます。プラットフォームの顧客の中には、独自の CI を定義したい人もいるかもしれません。その場合、動的ワークフローによって促進され、独自のワークフローを作成し、リポジトリだけでカスタム CI ジョブを実行できます。利点は、選択する必要がないことです。プラットフォーム マネージド CI とカスタム CI の両方を同じ名前空間で同時に実行できます。
CI/CD パイプラインは単なるワークフローです
今日まで、プラットフォームが Cloudflare 上で CI/CD パイプラインを接続できるようにするためのすべての要素が揃っていました。現在、開発者エクスペリエンスを向上させてシンプルにしています。
CI/CD パイプライン (通常は GitHub Actions で調整されます) は、特定の順序で実行される一連のステップであり、いずれかのステップが失敗した場合は、パイプラインの実行を停止してエラーを報告します。本質的に、CI/CD パイプラインは単なるワークフローです。 CI/CD を YAML ファイルで定義すると、YAML 疲労につながることが多い制約があるため、すぐに複雑になる可能性があります。ただし、CI/CD パイプラインの各ステップは、単純にワークフロー step.do() に変換できます。 YAML の代わりに、CI/CD パイプラインを定義できます。

n Typescript によるカスタマイズと構成の可能性の向上。
私たちは、ワークフローとサンドボックス SDK を介して Cloudflare の開発者プラットフォーム上に直接構築された、安全で隔離された環境で CI パイプラインの各ステップ (例: build 、 lint 、 typecheck ) を実行できる新しいツールを CI SDK でリリースします。さらに、イベント サブスクリプション、キュー、キュー コンシューマーを構成する代わりに、プッシュ時に直接 CI ジョブを開始できるようになりました。
以前は、サンドボックス API を直接呼び出して、CI パイプラインのさまざまなステップにわたって状態を自分で管理する必要がありました。 SDKを使用すると、サンドボックス化された各コマンドを独自のワークフローステップで実行でき、Cloudflareワークフローに組み込まれた再試行とタイムアウトが提供されます。
また、ステップ結果 (インストール ステップなど) をキャッシュすることで CI パイプラインを高速化し、後続のすべての操作で再インストールする必要がなくなります。依存関係のキャッシュにより、すべての CI ステップでインストールを再実行する必要がなくなるため、CI/CD パイプラインのレイテンシが短縮されます。
CI ジョブを定義するには、次のことを行うだけです。
バンドラー (例: esbuild )、リンター (例: eslint )、またはテスト ランナー (例: vitest ) などの依存関係 (CI ジョブに必要な外部パッケージまたはツール) のインストール ステップを定義します。
CI ジョブの各ステップのコマンドを指定します (例: bun run build 、 bun run test 、 bun run lint )。依存関係がキャッシュされると、各 CI ステップを並行して実行できるため、実行全体のレイテンシーが短縮されます。
デプロイ ステップで Wrangler デプロイを渡します。 CI パイプラインが通過すると、ワーカーは自動的にデプロイされます。
const deps : CiRunnerResult = ci を待ちます。ランナー ({
名前: 'インストール' 、
コマンド: 'bun install --frozen-lockfile' ,
キャッシュ: { 入力: [ 'package.json' , 'bun.lock' ] },
});
約束を待ちます。すべて([
デプス。ランナー ({ 名前: 'lint' 、コマンド: 'bun run lint' })、

デプス。ランナー ({ 名前: 'テスト' 、コマンド: 'ブンランテスト' })、
デプス。ランナー ({ 名前: 'typecheck' 、コマンド: 'bun run typecheck' })、
デプス。ランナー ({ 名前: 'build' 、コマンド: 'bun run build' })、
]);
デプスを待ちます。ランナー ({
名前: 'デプロイ' 、
コマンド: 'bun Wrangler デプロイ' 、
Cloudflare資格情報: {
accountId: この .env。 CLOUDFLARE_DEPLOY_ACCOUNT_ID 、
}、
});ワークフローに独自の CI パイプラインを作成すると、必要に応じてカスタマイズできます。たとえば、CI ワークフローからエージェントを呼び出して、CI ジョブに自己修復機能を提供できます。ビルドのステップでエラーが発生した場合、エージェントはそれを自動的に修正し、承認を求めるコミットをプッシュできます。
Project Think で自己修復 CI ワークフローの例を試してみましょう: https://github.com/cloudflare/ci/blob/main/examples/self-healing
独自の CI ワークフローを作成するには、 import { CIWorkflow } from@cloudflare/ci から始めます。
インストール手順から始めます。
CI ステップに必要な外部ツールやライブラリ (例: vite 、react ) を含む依存関係をダウンロードします。
依存関係が変更されたかどうかを追跡するロックファイルを指定します。
サンドボックス スナップショットを介して依存関係をキャッシュし、後続のすべてのステップでアクセスできるようにします。スナップショットはアカウントの R2 バケットに保存されます。
const deps = ci を待ちます。ランナー ({
名前: 'インストール' 、
コマンド: 'bun install --frozen-lockfile' ,
キャッシュ: { 入力: [ 'package.json' , 'bun.lock' ] },
});次に、ビルドとチェックのステップを定義し、それぞれが独自の安全で隔離されたサンドボックス環境で実行されます。
デフォルトでは、ワークフローの各ステップは独立して開始されます。つまり、特に指定しない限り、ステップは同時に実行されます。各ステップを並行して実行すると、CI 実行のレイテンシーが短縮されます。 CI パイプラインが続行される前にすべてのチェックが確実に完了するようにするには (たとえば、 build 、 lint の終了、

test 、およびデプロイ ステップの開始前に typecheck を実行し、 Promise.all() でラップします。
約束を待ちます。すべて([
デプス。ランナー ({ 名前: 'lint' 、コマンド: 'bun run lint' })、
デプス。ランナー ({ 名前: 'テスト' 、コマンド: 'ブンランテスト' })、
デプス。ランナー ({ 名前: 'typecheck' 、コマンド: 'bun run typecheck' })、
デプス。ランナー ({ 名前: 'build' 、コマンド: 'bun run build' })、
]);ここで、実際に CI ワークフローをトリガーするには、ワークフローとアーティファクトのバインディングと一緒に、ワーカーのラングラー設定にイベント フィールドを追加します。 events フィールドは、トリガー フィールド内でサポートされる新しいフィールドです。
イベントサブスクリプションを介してCloudflareキューを通じてアーティファクトをすでにサブスクライブし、プッシュイベントが発生するたびにビルドパイプラインを開始することができます。ただし、これにはイベント サブスクリプション、キュー、コンシューマー、およびキュー ハンドラーを設定する必要があります。これで、そのイベントでワークフローをターゲットにすることができます。そのイベントが発生するたびに、ワークフローのインスタンスがトリガーされます。
CI ワークフローをアーティファクト プッシュ トリガーのターゲットとして指定して、すべての cf.artifacts.repo.pushed イベントでワークフロー インスタンスを自動的にトリガーします。各 CI 実行はワークフロー インスタンスとして表示されるため、ワークフロー ダッシュボードでその段階的な実行と可観測性を直接確認できます。これはアーティファクトファーストの統合です。近日中に、これらのタイプは Cloudflare アカウント全体のソースからのイベントをサポートし、製品スイート全体でプログラムによる利用が可能になります。
名前空間内のすべてのリポジトリで CI ワークフローを実行する場合 (たとえば、すべての顧客のリポジトリで CI を実行しているプラ​​ットフォームの場合)、 repoName を省略し、 filter で名前空間のみを指定します。
{
"トリガー" : {
「イベント」: [
{
"type" : "cf.artifacts.repo.pushed" ,
// フィルターはオプションです。 repoName を設定しない場合は、同じワークフローが実行されます。

Artifacts 名前空間内の任意のリポジトリへのプッシュごとに
「フィルター」: {
"ネームスペース" : "CI" ,
"リポ名" : "私のリポジトリ"
}、
「ターゲット」: {
"タイプ" : "ワークフロー" ,
"ワークフロー名" : "ci-ワークフロー"
}
}
】
}
CI ワークフローを完全に構成するには、パイプラインを強化するインフラストラクチャの各部分にバインディングを追加します: artifacts 、 workflow 、containers、および durable_objects (+exports config) バインディング (サンドボックスにアクセスするため)、さらに、 cache を使用している場合は r2 バインディング。インストール ステップのサンドボックスのスナップショットがバケットに保存されるため、R2 バインディングが必要です。
CI ジョブを自己修復できるようにするには、LLM とそのエージェント ハーネスという 2 つの部分が必要になります。上の例では、Workers AI を使用してパイプライン内のエラーをキャッチし、ユーザーに代わって修正を実行する Think エージェントを組み込みました。 CI ジョブはリモートで実行および再実行できます。ラップトップを開いて監視したり、数分ごとに確認したりする必要はありません。代わりに、Cloudflare はこれをクラウドで処理し、コンテナー内の CI ステップと並行してヒーラー エージェントを実行します。 CI ジョブをベビーシッターして手動で修正し、パイプラインを再実行する代わりに、エージェントが修正を行った後にコミットをマージするだけで済みます。
CI パイプラインを自己修復するエージェントを設定するには、Think エージェントの Durable Object バインディングを追加します。
"耐久性のあるオブジェクト" : {
"バインディング" : [
{
"名前" : "ヒーラー" ,
"class_name" : "ヒーラー" ,
}、
]、
HealingAgent クラスを拡張して、Think エージェント (Healer) を作成します。このクラスには、失敗時に呼び出すための修復メソッドが含まれています。使用したいモデルを渡します。
エクスポート クラス Healer extends HealingAgent {
getModel() {
return '@cf/moonshotai/kimi-k2.7-code' ;
}
次に、失敗によって修復エージェントがトリガーされる try/catch ブロックでステップをラップします。
let deps : CiRunnerResult ;
{を試してください
// 一度インストールすると、その後は独立して実行されます

共有およびキャッシュされたスナップショットからチェックする
deps = ci を待ちます。ランナー ({
名前: 'インストール' 、
コマンド: 'bun install --frozen-lockfile' ,
キャッシュ: { 入力: [ 'package.json' , 'bun.lock' ] },
});
約束を待ちます。すべて([
デプス。ランナー ({ 名前: 'lint' 、コマンド: 'bun run lint' })、
デプス。ランナー ({ 名前: 'テスト' 、コマンド: 'ブンランテスト' })、
デプス。ランナー ({ 名前: 'typecheck' 、コマンド: 'bun run typecheck' })、
デプス。ランナー ({ 名前: 'build' 、コマンド: 'bun run build' })、
]);
} キャッチ (失敗) {
// これにより、失敗したサンドボックス コマンドと通常のワークフロー エラーの両方が捕捉されます。
// ランナーによって報告された障害のみが修復される必要があります。残りを再スローします。
if ( ! isCiRunnerFailure (失敗)) {
スロー失敗。
}
// エージェントがエラーを修正できるように、エラーをエージェントに渡します
const Healed = ステップを待ちます。する（
「癒す」、
{ 再試行: { 制限: 0 、遅延: 0 }、タイムアウト: '5 時間' }、
非同期() => {
const healer = await getAgentByName ( this .env.HEALER 、event.instanceId);
結果を使用 = ヒーラーを待ちます。癒す({
失敗：enrichFailure ({ 失敗、イベント、baseBranch })、
プロンプト: 「検証を弱めることなく、観察されたすべての障害を修正します。」 、
});
// Fix ブランチ、そのコミット、およびそれに要したステップ数を報告します。
const {ブランチ、コミット、ステップ} = 結果;
return { ブランチ、コミット、ステップ };
}
);
// ソースの実行は失敗したままになります

[切り捨てられた]

## Original Extract

Run CI/CD for millions of repos — on your platform, on Cloudflare

Run CI/CD for millions of repos — on your platform, on Cloudflare | The Cloudflare Blog
Skip to content All Categories AI
Login Contact Sales Blog AI Developer Platform Developers +2 Show 2 more tags 5 Tags Show 5 tags Selected Tags
AI Developer Platform Developers Product News Workflows
Automatic Platform Optimization
Cloudflare One User Risk Score
Run CI/CD for millions of repos — on your platform, on Cloudflare
André Venceslau , Mia Malden , and Tomáš Hobza
We are moving toward a world in which you can store, build, test, and deploy your code fully on Cloudflare. We built the first piece with Artifacts , versioned code storage that scales to millions of repos.
We have stitched the store, build, and deploy steps together with the CI SDK , built on Cloudflare Workflows , so that you can run your continuous integration (CI) pipeline on Cloudflare. You can send artifact push events directly to your Workflow, triggering an instance of its execution — a CI job, essentially — through a new events field in your wrangler configuration file.
Then, directly from the Workflow with @cloudflare/ci installed, you can:
Automate builds: compile code from your Artifacts repo in a safe, isolated environment
Run linters and typechecks: enforce code style, catch type errors, and flag any potential issues
Cache dependencies: run your install once and cache dependencies across steps in the CI job
Execute unit tests: verify that each piece of your code works as expected
Self-heal: integrate an AI review agent to catch broken steps in your build and push commits to fix
Deploy conditionally: automatically deploy your code, only if your build step is successful
Today, everyone is building a platform, whether it’s an internal vibe coding platform or an extension of your customer-facing product via customization through code. Platforms are now using millions of repos on Artifacts to store their code, and their customers’ code, and version control across the two. But every team has their own needs for a continuous integration and deployment pipeline. For platforms, they might want to define a CI job for their own code differently from that of their customers.
Many of the end customers building on these platforms don’t want the extra headache of managing their continuous integration and continuous deployment (CI/CD) pipeline. Instead, the platform can manage the build process on their customers’ behalf: write the CI/CD pipeline once and share it across all the applications that their customers are building. Some of the platform’s customers might want to define their own CI; if so, they can write their own Workflow and run custom CI jobs on just their repo, facilitated by dynamic workflows . The beauty is, you don’t have to pick and choose: both platform-managed and custom CI can run at the same time, in the same namespace.
A CI/CD pipeline is just a Workflow
Before today, we had all the pieces to allow platforms to wire their CI/CD pipeline together on Cloudflare. Now, we’re bringing a better developer experience to make it simple.
A CI/CD pipeline — commonly orchestrated with GitHub Actions — is a series of steps that run in a specific order where, if any step fails, you stop running the pipeline and report the error. In essence, a CI/CD pipeline is just a Workflow. CI/CD, when defined by a YAML file, can get complicated quickly, given the constraints that so often lead to YAML fatigue. But each step in a CI/CD pipeline can translate simply to a Workflow step.do() . Instead of YAML, you can define your CI/CD pipeline in Typescript for greater customization and configurability.
We are launching new tools in the CI SDK that allow you to run each step in your CI pipeline (e.g. build , lint , and typecheck ) in a safe, isolated environment, built directly on Cloudflare’s developer platform via Workflows and the Sandbox SDK . Plus, you can now kick off a CI job directly on push instead of configuring an event subscription, a queue, and a queue consumer.
Previously, you’d have to call the Sandbox API directly and manage state yourself across different steps in the CI pipeline. The SDK allows you to run each sandboxed command in its own Workflow step, providing the retries and timeouts built into Cloudflare Workflows.
You can also speed up your CI pipeline by caching step results — for example, your install step — so that you don’t need to reinstall for all subsequent operations. Dependency caching reduces the latency of your CI/CD pipeline since every CI step won’t need to rerun the install.
To define your CI job, all you need to do is:
Define your install step for any dependencies (external packages or tools that your CI job needs), such as bundlers (e.g. esbuild ), linters (e.g. eslint ), or test runners (e.g. vitest ).
Specify the command for each step in the CI job (e.g. bun run build , bun run test , bun run lint ). With your dependencies cached, each CI step can execute in parallel, reducing the latency of the overall run.
Pass wrangler deploy in a deploy step. Your Worker will automatically deploy when the CI pipeline passes.
const deps : CiRunnerResult = await ci. runner ({
name: 'install' ,
command: 'bun install --frozen-lockfile' ,
cache: { inputs: [ 'package.json' , 'bun.lock' ] },
});
await Promise . all ([
deps. runner ({ name: 'lint' , command: 'bun run lint' }),
deps. runner ({ name: 'test' , command: 'bun run test' }),
deps. runner ({ name: 'typecheck' , command: 'bun run typecheck' }),
deps. runner ({ name: 'build' , command: 'bun run build' }),
]);
await deps. runner ({
name: 'deploy' ,
command: 'bun wrangler deploy' ,
cloudflareCredentials: {
accountId: this .env. CLOUDFLARE_DEPLOY_ACCOUNT_ID ,
},
}); Writing your own CI pipeline in a Workflow allows you to customize as much as you want. For example, you could call an agent from your CI Workflow to give your CI jobs self-healing functionality: if a step in your build errors, the agent can fix it automatically, and push a commit for your approval.
Try an example of self-healing CI Workflows with Project Think: https://github.com/cloudflare/ci/blob/main/examples/self-healing
To write your own CI Workflow, get started with import { CIWorkflow } from@cloudflare/ci .
Start with an install step:
Download your dependencies, including any external tools or libraries that your CI steps will need (e.g. vite , react ).
Specify your lockfile, which tracks whether your dependencies have changed.
Cache your dependencies via a sandbox snapshot so that all subsequent steps have access. The snapshot will be stored in an R2 bucket on your account.
const deps = await ci. runner ({
name: 'install' ,
command: 'bun install --frozen-lockfile' ,
cache: { inputs: [ 'package.json' , 'bun.lock' ] },
}); Then define steps for the build and checks, each executed in its own safe, isolated sandbox environment.
By default, each step in a Workflow starts independently, meaning the steps will execute concurrently unless otherwise specified. Running each step in parallel reduces the latency of your CI run. To ensure that all checks complete before the CI pipeline continues (for example, finish build , lint , test , and typecheck before the deploy step starts), wrap in a Promise.all() :
await Promise . all ([
deps. runner ({ name: 'lint' , command: 'bun run lint' }),
deps. runner ({ name: 'test' , command: 'bun run test' }),
deps. runner ({ name: 'typecheck' , command: 'bun run typecheck' }),
deps. runner ({ name: 'build' , command: 'bun run build' }),
]); Now, to actually trigger your CI Workflow, add an events field to your Worker’s wrangler configuration, alongside your Workflow and Artifact bindings. The events field is a new field supported within your triggers field.
You could already subscribe to Artifacts through Cloudflare Queues via event subscriptions and kick off a build pipeline every time there’s a push event. But that requires setting up the event subscription, Queue, consumer, and queue handler. Now, you can target a Workflow with that event — every time that event fires, it will trigger an instance of the Workflow.
Specify the CI Workflow as your artifact push trigger’s target to automatically trigger a Workflow instance on every cf.artifacts.repo.pushed event. Each CI run surfaces as a Workflow instance so you can view its step-by-step execution and observability directly in the Workflows dashboard. This is an Artifacts-first integration; coming soon, the types will support events from sources across your Cloudflare account to allow for programmatic consumption across the product suite.
If you want to run the CI Workflow on every repo in your namespace — for example, if you are a platform running CI on all of your customers’ repositories — omit repoName and only specify the namespace in filter .
{
"triggers" : {
"events" : [
{
"type" : "cf.artifacts.repo.pushed" ,
// filter is optional. If you don't set repoName we will run the same workflow for every push on any repo in your Artifacts namespace
"filter" : {
"namespace" : "CI" ,
"repoName" : "my-repo"
},
"target" : {
"type" : "workflow" ,
"workflow_name" : "ci-workflow"
}
}
]
}
} To fully configure your CI Workflow, add bindings to each piece of the infrastructure which powers the pipeline: artifacts , workflows , containers and durable_objects (+ exports config) bindings (to access your sandboxes), plus an r2 binding if you are using cache . The R2 binding is required as the snapshot of your install step sandbox is stored in a bucket .
To allow your CI job to self-heal, you’ll need two pieces: the LLM and its agent harness. In the example above, we included a Think agent using Workers AI to catch errors in your pipeline and run the fixes on your behalf. Your CI job can be run and re-run remotely — no need to watch with your laptop open or check back every few minutes. Instead, Cloudflare handles it in the cloud, running your healer agent alongside the CI steps in a container. Instead of babysitting the CI job, making a manual fix, and re-running the pipeline, you’ll just need to merge the commit after your agent has made the fix.
To set up an agent that self-heals your CI pipeline, add a Durable Object binding for your Think agent:
"durable_objects" : {
"bindings" : [
{
"name" : "HEALER" ,
"class_name" : "Healer" ,
},
],
}, Create your Think agent — Healer — by extending the HealingAgent class, which includes a heal method for you to call on failure. Pass whichever model you’d like to use:
export class Healer extends HealingAgent {
getModel () {
return '@cf/moonshotai/kimi-k2.7-code' ;
}
} Then, wrap your steps in a try/catch block where a failure triggers the healing agent:
let deps : CiRunnerResult ;
try {
// Install once, then run independent checks from the shared and cached snapshot
deps = await ci. runner ({
name: 'install' ,
command: 'bun install --frozen-lockfile' ,
cache: { inputs: [ 'package.json' , 'bun.lock' ] },
});
await Promise . all ([
deps. runner ({ name: 'lint' , command: 'bun run lint' }),
deps. runner ({ name: 'test' , command: 'bun run test' }),
deps. runner ({ name: 'typecheck' , command: 'bun run typecheck' }),
deps. runner ({ name: 'build' , command: 'bun run build' }),
]);
} catch (failure) {
// This catches both failed Sandbox commands and ordinary Workflow errors.
// Only failures reported by a runner should be healed; rethrow the rest.
if ( ! isCiRunnerFailure (failure)) {
throw failure;
}
// Pass the error along to the agent so that it can fix it
const healed = await step. do (
'heal' ,
{ retries: { limit: 0 , delay: 0 }, timeout: '5 hours' },
async () => {
const healer = await getAgentByName ( this .env. HEALER , event.instanceId);
using result = await healer. heal ({
failure: enrichFailure ({ failure, event, baseBranch }),
prompt: 'Fix every observed failure without weakening validation.' ,
});
// Report the Fix Branch, its commit, and how many steps it took.
const { branch , commit , steps } = result;
return { branch, commit, steps };
}
);
// The source run stays fail

[truncated]
