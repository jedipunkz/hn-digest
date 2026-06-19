---
source: "https://technicalstrat.com/articles/enterprise-app-vibecode-recipe"
hn_url: "https://news.ycombinator.com/item?id=48601721"
title: "Two production Next.js apps, built solo with AI, $13,945"
article_title: "How I Vibe-Coded Two Enterprise Apps in 8 Months for $13,945 | TechnicalStrat | TechnicalStrat"
author: "mpistole"
captured_at: "2026-06-19T18:42:12Z"
capture_tool: "hn-digest"
hn_id: 48601721
score: 1
comments: 0
posted_at: "2026-06-19T18:39:27Z"
tags:
  - hacker-news
  - translated
---

# Two production Next.js apps, built solo with AI, $13,945

- HN: [48601721](https://news.ycombinator.com/item?id=48601721)
- Source: [technicalstrat.com](https://technicalstrat.com/articles/enterprise-app-vibecode-recipe)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T18:39:27Z

## Translation

タイトル: AI を使用して単独で構築された 2 つの本番 Next.js アプリ、13,945 ドル
記事のタイトル: 8 か月で 13,945 ドルで 2 つのエンタープライズ アプリを Vibe コーディングした方法 |テクニカルストラト |テクニカルストラト
説明: Azure 上の 2 つの本番 Next.js アプリ、約 350,000 行の TypeScript、Cursor および Anthropic のフロンティア モデルを使用して単独で構築されました。スタック、支出、レシピ。

記事本文:
8 か月で 13,945 ドルで 2 つのエンタープライズ アプリを Vibe コーディングした方法 |テクニカルストラト | TechnicalStrat コンテンツにスキップ ts 。 TechnicalStrat 記事 Toggle テーマ Vibe コーディングについて
8 か月で 13,945 ドルで 2 つのエンタープライズ アプリを Vibe コーディングした方法
Azure 上の 2 つの本番 Next.js アプリ、約 350,000 行の TypeScript、Cursor および Anthropic のフロンティア モデルを単独で構築しました。スタック、支出、レシピ。
LinkedIn X / Twitter を共有する リンクをコピーする 始める前に
この記事は私の経験の証言であり、専門的なアドバイスではありません。私のバックグラウンドは IT なので、Azure アーキテクチャには長年の運用上の自信があります。アプリケーション コードは Cursor を使用してバイブコーディングされており、私の専門の操舵室ではありません。ターンキー生産システムではなく、レシピを適応の出発点として使用してください。プロビジョニングする前に Azure 予算アラートを設定し、長いエージェント セッションの前にカーソルの支出を監視します。クラウドの請求書、AI の請求書、またはビルドからの生産結果はすべてお客様のものです。 TechnicalStrat と私はそれらに対して一切の責任を負いません。
2 つの運用グレードのエンタープライズ アプリ (どちらも Azure Container Apps 上の Next.js 15) は、Cursor と Anthropic のフロンティア モデルを使用して 8 か月で単独で構築されました。単独開発者の支出: 34,875 のカーソル イベントと 450 のアクティブなコーディング時間で 13,945 ドル。以下のレシピは、それらを作成した正確なスタックとロールアウト計画です。ダウンロードしてカーソルに貼り付け、独自のものを作成します。
2 つのアプリ、1 つのスタック、1 つのロールアウト パターンにわたる TypeScript の最大 350,000 行
モデル支出の大部分は、人類オーパス思考モデルに費やされました。フレームワークよりもフロンティアモデルが重要
正直な機能あたりのコストは、AI が間違っていた部分も含めて、アクティブなコーディング時間あたり 1 時間あたり約 31 ドルでした。
以下: アーキテクチャ、露出モデルの決定、何が機能したのか、何が機能しなかったのか、そして

カーソルに直接入力できるレシピ
私は過去 8 か月間、Cursor と Anthropic のフロンティア モデルを使用して 2 つのエンタープライズ アプリケーションを単独で出荷してきました。どちらも現在生産中です。一つは商業ローン組成制度です。もう 1 つは、シングル サインオンの背後に多数のバックオフィス ツールを統合するマルチアプリの内部プラットフォームです。問題のドメインは異なりますが、スタックはほぼ同じです。
私はこれを、「バイブコーディング」を聞いたことがある、それが本物なのか、そのコストはどれくらいなのか、裏側の実際のアーキテクチャはどのようなものなのか知りたいと考えているエンジニアまたは技術創設者に向けて書いています。誇大広告はありません。数字、スタック、レシピだけです。
ビルド ウィンドウ上のアカウントのカーソル使用状況レポートから直接取得したもの:
8 か月のカレンダー時間、約 450 時間の集中的なコーディング、AI 支出は 13,945 ドル。 2 つのアプリが出荷されました。比較のために、時給 150 ドルの中級レベルの契約エンジニア 1 名がいた場合、同じ 450 時間のコストは 67,500 ドルかかり、その時間枠で 350,000 行のコードを記述することはできません。
それが見出し番号です。正直なバージョンは以下です。
たどり着いたスタック (両方のアプリ)
どちらのアプリケーションも同じアーキテクチャに統合されました。その収束は計画されていなかった。 AI に妥当なデフォルトを選択させ、スケールする最も単純なものに向かって軌道修正し、フレームワークとの戦いをやめると、このようなことが起こります。
両方のアプリでバックグラウンドで作業するためのスケジュールされたコンテナ アプリ ジョブ
アプリごとのアプリケーション登録を備えた 1 つの共有 Microsoft Entra ID テナント
Kubernetes クラスター 0、予想外の請求は 0
あらゆるアーキテクチャ上の選択は 3 つの原則に基づいて行われました。
境界線は退屈、途中からは面白い。接続 (認証、データベース、デプロイ、シークレット、テレメトリ) はすべて、ファーストパーティの Microsoft またはファーストパーティのフレームワークのデフォルトです。興味深いコードはビジネス ログにあります

ic は、カスタムで記述する価値のある唯一のコードでもあります。デフォルト以外のサービスを結合するのに費やされる 1 時間は、実際の問題に費やされない 1 時間になります。
エンドツーエンドの静的型付け。型がデータベース行から Prisma クライアント、API ハンドラー、Zod スキーマを経由して React フォームに流れると、AI は劇的に向上します。 Prisma モデルに 1 つの変更を加えると、壊れた箇所がすべて表面化します。コンパイラは、AI がキャッチできないリファクタリングの自動操縦です。
長続きする秘密はどこにもありません。 GitHub は OIDC 経由で Azure にフェデレーションします。コンテナー アプリは、マネージド ID を介して Key Vault からプルします。 GitHub の唯一のシークレットは、移行ジョブに必要な DATABASE_URL であり、それさえもアプリではなく移行ランナーに限定されています。資格情報のローテーションは、デプロイではなく、30 秒間の Key Vault 操作です。
複合効果: 前回どの代替パスを使用したかを覚えていなくても、AI がエンドツーエンドでナビゲートできるスタックになります。
インターネットに接続されたセキュリティのための 2 つのパス
2 つのアプリには異なる公開要件があり、それがスタック自体の次に最も重大な決定を引き起こしました。
アプリ A には、テナント以外のユーザーがパブリック インターネットからアクセスできる必要があるクライアント ポータルと、パブリック URL に到達する必要があるサードパーティの金融サービスからの Webhook があります。アプリ プロキシは適していませんでした。代わりに、私はサードパーティの専門的なセキュリティレビューに料金を支払いました。つまり、上級開発者が認証フロー、マジックリンクポータル、レート制限、入力検証、および起動前の監査体制について承認しました。コンテナ アプリは、管理された証明書を通じてパブリック トラフィックを直接受け取ります。
アプリ B は内部専用です。すべてのユーザーは、アプリケーションを表示する前に、同じ Microsoft Entra ID テナントに対して認証を行います。 Microsoft Entra Application Proxy をその前に置きました。ザ・コンテナ

r アプリ環境は内部イングレスで実行され、唯一のパブリック サーフェスはアプリ プロキシ URL であり、アプリ プロキシはリクエストがコードに到達する前に事前認証、条件付きアクセス、およびトラフィック検査を強制します。これは、ユーザーに VPN クライアントをインストールさせる必要のない、VPN のセキュリティ体制です。料金は、ユーザーごとの月額料金と、パッチを適用し続けるための 1 つのコネクタ VM です。テナントの外部から決してアクセスできないアプリの場合、これはスタック内で最も簡単な呼び出しです。
どちらの姿勢も有効です。どちらを選択するかは、誰がアプリにアクセスする必要があるかによって決まります。
内部ユーザーのみ、Entra テナント内のすべてのユーザー: Azure App Proxy を前面に置きます。ユーザーごとの定期的なコストにより、VPN グレードの分離、エッジでの条件付きアクセス、および最小限の攻撃対象領域が実現します。アプリのセキュリティ グレードがまったく不明な場合は、これがより安全なデフォルトです。
テナント以外のユーザー、カスタマー ポータル、またはパートナー アクセス: コンテナ アプリをパブリック Ingress で立ち上げ、リリース前に有料の実際のセキュリティ レビューの予算を設定します。レビューを受ける余裕がない場合は、アプリを一般公開しないでください。
以下のレシピは両方のパスをカバーしており、選択した露出モデルに応じて変化するステップを示しています。
バイブコーディングは実際にどんな感じだったのか
この用語は「AI が私の for ループを自動補完した」から「GPT-4 が私のビジネス全体を書いた」まで、あらゆる意味に拡大解釈されているため、この点については正確に言いたいと思います。
それは私にとって実際にどのように見えたか:
1 日の労働時間は 3 ～ 5 時間の集中的なカーソル セッションでした。私は機能仕様書 (通常は docs/specs/ のマークダウン ファイルに入力した段落) を開き、モデル (ほとんどの場合、拡張思考モードの Claude Opus) に計画を提案するように依頼しました。私は計画を読み、同意できない部分を押し戻し、それから実行を依頼しました。私はそうするだろう

差分を確認し、型チェッカーを実行し、テストを実行し、反復します。
モデルがタイピングをしてくれました。決断は私が行いました。
ほとんどの機能は、着陸するまでに 2 ～ 4 回の反復を要しました。通常、最初のパスは、誤った仮定が組み込まれた機能するスケルトンでした。2 番目のパスでは、その仮定が修正されました。 3 番目のパスでは型を強化し、エラー処理を追加し、テストを作成しました。場合によっては、テスト スイートで検出された問題を修正するために 4 回目のパスを実行することもあります。
正直な要約: 私は、非常に高速で非常に文字通りの若手エンジニアからのプル リクエストをレビューする技術リーダーとして活動していました。彼は決して疲れず、決して守りに入らず、時には Prisma クライアントに存在しないメソッドの幻覚を見せます。
レバレッジは本物です。監督については交渉の余地はありません。
AIが強いところとそうでないところ
このパターンは両方のアプリで一貫していました。
リポジトリ内の既存の規則に一致する新しい API ルートをスキャフォールディングする
新しい Prisma モデルとそれに伴う移行の設定
明確な英語の仕様を、動作するサーバー コンポーネントとそのデータ ローダーに翻訳する
人間が書きたくない種類のグルー コードの作成 - 再試行ロジック、指数関数的バックオフ、冪等性キー、監査ログ エントリ
予想される動作の記述からテストを生成し、それを実行して障害を修正する
Cursor の Composer またはエージェント モードを使用した流暢な複数ファイルのリファクタリング
正しい答えがリポジトリの外のコンテキスト (コンプライアンス、コストの上限、ベンダーの好み) に依存するアーキテクチャの決定。防御可能ではあるが最適ではないものを喜んで選択します。
システム間の不変条件に関係するあらゆるもの: データベース行との同期を維持する必要がある Webhook、スケジュールされたジョブの競合状態、分散ロック。差分を注意深く読んで本当のバグを見つけました。
いつ止めるべきかを知ること。放っておいても、モデルは

実際には単なる再配置にすぎない「改善」を追加しました。範囲について明確にする必要がありました。
約 1,500 行を超えるファイルの長いコンテキストのリファクタリング。モデルは床に微妙なものを落とします。ファイルを最初に分割することは、とにかく AI が好む強制的な機能でした。
唯一最大のロック解除は、リポジトリ ルートにある CLAUDE.md (または AGENTS.md ) ファイルでした。スタック、規則、内部ヘルパーの名前、AI が決して実行すべきではないことについて説明する 2 ページのドキュメント。すべてのセッションはそのコンテキストを無償で継承しました。それがなければ、すべてのチャットはゼロから始まりました。
この法案が興味深いのは、フロンティアモデルを目新しいものとしてではなく、実際に日常のドライバーとして使用した場合にどのくらいの費用がかかるかを示しているためです。
私の支出 13,945 ドルの大部分は、人類オーパス思考モデルに使われました。 Sonnet は、ボイラープレート、コミット メッセージ、小さな型付きリファクタリングなど、より深い推論を必要としない作業を処理しました。いくつかの安価な非思考モデルがトリビアを取り上げました。コードベースが増大し、より多くのセッションがマルチファイル エージェントの作業に移ったため、支出はビルド ウィンドウの後半に集中しました。
正直なトレードオフ: フロンティア モデルには実際のお金がかかります。 Sonnet のみのビルドは、スキーマの再設計、複数ファイルのリファクタリング、厄介な Prisma と Postgres の相互作用のデバッグ、セキュリティに敏感なコードなどのハード部分で大幅にコストが安くなり、著しく遅くなったでしょう。オーパスの考え方のプレミアムはそこで元が取れました。
この比率でもう一度やろうと思います。エンタープライズ グレードのビルドでは、Sonnet のみの比率では実行しません。
2 番目のビルドで省略しないこと
後になってこの 2 つのアプリを振り返ってみると、最も重要なのはテクノロジーではありませんでした。それらは規律でした。
初日の CLAUDE.md。 200ワード版でも。 AI は、リポジトリのルールを認識すると、劇的に優れた性能を発揮します。
ライブラリ/許可

初日から。ハンドラー全体に散在する承認は、後でクリーンアップするのに最もコストがかかる間違いです。 2 番目のハンドラーが存在する前に集中化します。
ヘルスエンドポイントと初日からの監査ログ。それらなしで発送することはできます。それらなしでは操作できません。
初日から GitHub から Azure への OIDC。後で静的資格情報から OIDC に切り替えるのは面倒です。開始時の OIDC のセットアップには 20 分かかります。
最初のデプロイ前のエクスポージャ モデルの決定。内部アプリ: Azure アプリ プロキシの予算。パブリック アプリ: 有料のセキュリティ レビューの予算。どちらでも構いません。両方をスキップすることはできません。
施行前のレポート専用の条件付きアクセス ポリシー。ユーザーあたり 50 ドルのハードウェアで SMB Microsoft 365 テナントをロックダウンする方法についての記事をすべてまとめました。同じパターンが、あなたが自社の製品のために立ち上げた Entra アプリにも当てはまります。
スキップするもの: 2 か月目に構築したオーダーメイドの抽象化のほとんどは、6 か月目に削除しました。 AI は抽象化を構築するのが大好きです。必要な 3 番目の場所まで抵抗してください。
この作品の要点は以下のレシピです。これは、このスタックでアプリをブートストラップするために初日に Cursor にフィードする正確な段階的な計画です。 60 以上の番号が付けられたステップ、オルガニ

[切り捨てられた]

## Original Extract

Two production Next.js apps on Azure, ~350,000 lines of TypeScript, built solo with Cursor and Anthropic's frontier models. The stack, the spend, and the recipe.

How I Vibe-Coded Two Enterprise Apps in 8 Months for $13,945 | TechnicalStrat | TechnicalStrat Skip to content ts . TechnicalStrat Articles About Toggle theme Vibe Coding
How I Vibe-Coded Two Enterprise Apps in 8 Months for $13,945
Two production Next.js apps on Azure, ~350,000 lines of TypeScript, built solo with Cursor and Anthropic's frontier models. The stack, the spend, and the recipe.
Share LinkedIn X / Twitter Copy link Before you start
This article is a testimony of my experience, not professional advice. My background is in IT, so the Azure architecture is where I have years of operational confidence. The application code was vibe-coded with Cursor and is not my professional wheelhouse. Use the recipe as a starting point to adapt, not a turnkey production system. Set Azure budget alerts before you provision and watch your Cursor spend before any long agent session. Any cloud bills, AI bills, or production outcomes from your build are yours to own; TechnicalStrat and I assume no liability for them.
Two production-grade enterprise apps, both Next.js 15 on Azure Container Apps, built solo in eight months with Cursor and Anthropic's frontier models. Solo developer spend: $13,945 across 34,875 Cursor events and 450 active coding hours. The recipe below is the exact stack and rollout plan that produced them. Download it, paste it into Cursor, build your own.
~350,000 lines of TypeScript across two apps, one stack, one rollout pattern
The majority of model spend was on Anthropic Opus-thinking models. The frontier model matters more than the framework
The honest cost-per-feature was about $31/hour of active coding time, all-in, including the parts where the AI was wrong
Below: the architecture, the exposure-model decision, what worked, what didn't, and the recipe you can feed straight into Cursor
I spent the last eight months shipping two enterprise applications, solo, using Cursor and Anthropic's frontier models. Both are in production now. One is a commercial loan origination system. The other is a multi-app internal platform that consolidates a dozen back-office tools behind a single sign-on. Different problem domains, almost identical stack.
I am writing this for the engineer or technical founder who has heard "vibe coding" and wants to know whether it is real, what it costs, and what the actual architecture looks like on the other side. No hype. Just the numbers, the stack, and the recipe.
Pulled directly from the Cursor usage report for my account over the build window:
Eight months of calendar time, roughly 450 hours of focused coding, $13,945 in AI spend. Two apps shipped. For comparison: a single mid-level contract engineer at $150/hour would have cost $67,500 for the same 450 hours, and would not have written 350,000 lines of code in that window.
That is the headline number. The honest version is below.
The stack I landed on (both apps)
Both applications converged on the same architecture. That convergence was not planned. It is what happens when you let the AI pick reasonable defaults, course-correct toward the simplest thing that scales, and stop fighting the framework.
Scheduled Container Apps Jobs for background work on both apps
1 shared Microsoft Entra ID tenant with per-app application registrations
0 Kubernetes clusters, 0 surprise bills
Three principles drove every architectural choice.
Boring at the boundary, interesting in the middle. The plumbing - auth, database, deploy, secrets, telemetry - is all first-party Microsoft or first-party framework defaults. The interesting code is in the business logic, which is also the only code worth writing custom. Every hour spent gluing a non-default service together is an hour not spent on the actual problem.
Static typing end to end. The AI is dramatically better when types flow from the database row through the Prisma client through the API handler through the Zod schema into the React form. One change to a Prisma model surfaces every place it breaks. The compiler is the autopilot for refactors the AI does not catch.
No long-lived secrets anywhere. GitHub federates to Azure via OIDC. Container Apps pull from Key Vault via managed identity. The only secret in GitHub is the DATABASE_URL needed by the migration job, and even that is scoped to the migration runner, not the app. Rotating credentials is a 30-second Key Vault operation, not a deploy.
The combined effect: a stack the AI can navigate end-to-end without you having to remember which alternative path you took last time.
Two paths for internet-facing security
The two apps had different exposure requirements, and that drove the single most consequential decision after the stack itself.
App A has a client portal that must be reachable from the public internet for non-tenant users, plus webhooks from third-party financial services that have to land on a public URL. App Proxy was not a fit. Instead, I paid for a third-party professional security review: a senior developer signed off on the auth flows, the magic-link portal, the rate limits, the input validation, and the audit posture before launch. The Container App takes public traffic directly through its managed certificate.
App B is internal-only. Every user authenticates against the same Microsoft Entra ID tenant before they ever see the application. I put Microsoft Entra Application Proxy in front of it. The Container Apps Environment runs with internal ingress, the only public surface is the App Proxy URL, and the App Proxy enforces pre-authentication, conditional access, and traffic inspection before any request reaches my code. It is the security posture of a VPN, without making users install a VPN client. The cost is a per-user monthly charge and one connector VM to keep patched. For an app that should never be reachable from outside the tenant, that is the easiest call in the stack.
Both postures are valid. Which one you pick is a function of who needs to reach the app:
Internal users only, all inside your Entra tenant : put Azure App Proxy in front. The recurring per-user cost buys you VPN-grade isolation, conditional access at the edge, and the lowest possible attack surface. If you are at all uncertain about your app's security grade, this is the safer default.
Any non-tenant users, customer portals, or partner access : stand the Container App up with public ingress and budget for a real, paid security review before launch. If you cannot afford the review, the app should not launch publicly.
The recipe below covers both paths and marks the steps that change depending on which exposure model you pick.
What vibe coding actually felt like
I want to be precise about this because the term has been stretched to mean everything from "AI autocompleted my for-loop" to "GPT-4 wrote my entire business."
What it actually looked like for me:
A working day was three to five hours of focused Cursor sessions. I would open a feature spec - usually a paragraph I had typed into a markdown file in docs/specs/ - and ask the model (almost always Claude Opus in extended thinking mode) to propose a plan. I would read the plan, push back on parts I disagreed with, then ask it to execute. I would watch the diffs, run the type checker, run the tests, and iterate.
The model did the typing. I did the deciding.
Most features took two to four iterations to land. The first pass was usually a working skeleton with a wrong assumption baked in. The second pass corrected the assumption. The third pass tightened the types, added error handling, wrote the tests. Sometimes a fourth pass to fix something the test suite caught.
The honest summary : I was acting as a technical lead reviewing pull requests from an extremely fast, extremely literal junior engineer who never tires, never gets defensive, and occasionally hallucinates a method that does not exist on the Prisma client.
The leverage is real. The supervision is non-negotiable.
Where the AI was strong, and where it wasn't
The pattern was consistent across both apps.
Scaffolding new API routes that matched the existing conventions in the repo
Wiring up a new Prisma model and the migration that goes with it
Translating a clear English spec into a working server component plus its data loader
Writing the kind of glue code humans hate to write - retry logic, exponential backoff, idempotency keys, audit-log entries
Generating tests from a description of expected behavior, then running them and fixing the failures
Fluent multi-file refactors with Cursor's Composer or agent mode
Architecture decisions where the right answer depended on context outside the repo (compliance, cost ceilings, vendor preference). It would happily pick something defensible but suboptimal.
Anything involving cross-system invariants: webhooks that needed to remain in lockstep with a database row, race conditions in scheduled jobs, distributed locks. I caught real bugs by reading the diff carefully.
Knowing when to stop. Left alone, the model would keep adding "improvements" that were really just rearrangements. I had to be explicit about scope.
Long-context refactors on files past about 1,500 lines. The model would drop subtle things on the floor. Splitting files first was a forcing function the AI liked anyway.
The single biggest unlock was the CLAUDE.md (or AGENTS.md ) file at the repo root. A two-page document describing the stack, the conventions, the names of internal helpers, and the things the AI should never do. Every session inherited that context for free. Without it, every chat started from zero.
The bill is interesting because it shows what frontier models cost when you actually use them as a daily driver, not as a novelty.
The majority of my $13,945 spend went to Anthropic Opus-thinking models. Sonnet handled the work that did not need deeper reasoning: boilerplate, commit messages, small typed refactors. A handful of cheaper non-thinking models picked up trivia. Spend was concentrated in the second half of the build window as the codebases grew and more sessions moved into multi-file agent work.
The honest tradeoff: frontier models cost real money. A Sonnet-only build would have been meaningfully cheaper and noticeably slower on the hard parts: schema redesigns, multi-file refactors, debugging gnarly Prisma and Postgres interactions, security-sensitive code. The Opus-thinking premium paid for itself there.
I would do it again at this ratio. I would not do it at a Sonnet-only ratio for an enterprise-grade build.
What I would not skip on a second build
Looking at the two apps with hindsight, the things that mattered most were not the technologies. They were the disciplines.
A CLAUDE.md from day one. Even a 200-word version. The AI is dramatically better when it knows the rules of your repo.
A lib/permissions.ts from day one. Authorization sprinkled across handlers is the single most expensive mistake to clean up later. Centralize it before the second handler exists.
A health endpoint and an audit log from day one. You can ship without them. You cannot operate without them.
OIDC from GitHub to Azure from day one. Switching from static credentials to OIDC later is annoying. Setting up OIDC at the start is twenty minutes.
The exposure-model decision before the first deploy. Internal app: budget for Azure App Proxy. Public app: budget for a paid security review. Either is fine. Skipping both is not.
Report-only Conditional Access policies before any enforcement. I have an entire piece on how we lock down SMB Microsoft 365 tenants for $50 of hardware per user - the same pattern applies to the Entra apps you stand up for your own product.
What I would skip: most of the bespoke abstractions I built in month two that I deleted in month six. The AI loves building abstractions. Resist until the third place you need them.
The whole point of this piece is the recipe below. It is the exact step-by-step plan I would feed into Cursor on day one to bootstrap an app on this stack. Sixty-plus numbered steps, organi

[truncated]
