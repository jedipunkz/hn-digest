---
source: "https://www.petrostechchronicles.com/blog/How_I_Ship_Production_Features_With_AI_Agents"
hn_url: "https://news.ycombinator.com/item?id=48881414"
title: "I Use Multiple AI Agents in My Team's SDLC to Ship Production Features"
article_title: "How I Ship Production Features with AI Agents (Not Just LinkedIn post or Promotion PoC Demos) | Tech Oriented Chronicles🔮 of Petros"
author: "Aherontas"
captured_at: "2026-07-12T14:24:18Z"
capture_tool: "hn-digest"
hn_id: 48881414
score: 1
comments: 0
posted_at: "2026-07-12T14:18:57Z"
tags:
  - hacker-news
  - translated
---

# I Use Multiple AI Agents in My Team's SDLC to Ship Production Features

- HN: [48881414](https://news.ycombinator.com/item?id=48881414)
- Source: [www.petrostechchronicles.com](https://www.petrostechchronicles.com/blog/How_I_Ship_Production_Features_With_AI_Agents)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T14:18:57Z

## Translation

タイトル: 実稼働機能を出荷するためにチームの SDLC で複数の AI エージェントを使用しています
記事のタイトル: AI エージェントを使用してプロダクション機能を出荷する方法 (LinkedIn の投稿やプロモーション PoC デモだけではない) |ペトロスの技術指向年代記🔮
説明: AI エージェントを使用して機能を出荷するための実際の運用ワークフロー: 3 つのターミナル、PID のようなフィードバック ループ、独立したレビュー、品質ゲート。バイブコーディングは不要で、実際のユーザーにサービスを提供する制御されたエンジニアリングのみです。

記事本文:
AI エージェントを使用してプロダクション機能を出荷する方法 (LinkedIn の投稿やプロモーション PoC デモだけでなく) |ペトロスの技術指向クロニクル🔮 Petros-Tech-Chronicles🔮 ブログ タグ プロジェクトについて ホーム ブログ タグ プロジェクト 概要 公開日: 2026 年 7 月 2 日木曜日 AI エージェントを使用してプロダクション機能を出荷する方法 (LinkedIn の投稿やプロモーション PoC デモだけでなく)
名前 Petros Savvakis Twitter @PetrosSavvakis
AI エージェントを使用してプロダクション機能を提供する方法 (LinkedIn の投稿やプロモーション PoC デモだけでなく)
エージェント コーディング、ハーネス、バイブ コーディング、AI チームメイトなど。これらは、2026 年のエンジニアリング分野で登場する「クールな」キーワードのほんの一部にすぎません。それはすべてクールであり、新しいことを試し、実験し、いじくり回すことを意味します...しかし、少なくとも私の場合、個人プロジェクトと独立したコラボレーション全体で試したことから、実際にはどのように機能するのでしょうか?
まず一歩下がって、ここ数か月間、私は単なるオーケストレーターであり、主にコードレビュー担当者であることをお伝えしましょう。以前のようにコードの完全なスクリプトを書くことはめったにありません。コードを書く必要がある最も一般的なのは、CAVEMAN MODE = ON になるときです (私はこの用語を Conductor CEO の Charlie Holtz から Conductor のマニュアル モード経由で初めて聞きました) ので、私たちはそれを採用しました。 AI モデルがタスクを完了するのに苦労している場合、または .env 値をエージェントに公開したくない場合は、ファイル (構成ファイルや .env ファイルなど) を迅速に編集するために手動で入力する必要がある時代です。
以上のことを念頭に置いて、専門用語はやめて、実際の結果が得られそうなものについて深く掘り下げてみましょう。ここ数か月間、私たちは個人プロジェクトや独立したコラボレーションを通じて、狂ったように機能をリリースしてきました (実際に動作し、完全にテストされています。ローカルホストで実行されるバイブコーディングだけでなく、実際のユーザーにサービスを提供する運用環境でも)。そのためには次のことを行う必要があります

o PID 適応制御のように機能する正当なパイプラインを作成します (電気エンジニアの方なら、私が何を意味するかおわかりでしょう)。これは、自己調整され、安定した実際の生産結果を生成できるシステムです。ここまで述べたので、そのアーキテクチャに移りましょう。
ワークフロー: 1 人の人間のオーケストレーター、3 つの端末、1 つの制御ループ
これが私にとって役に立ったシステムです。
私は、エージェントを 1 つだけ開いて、漠然としたプロンプトを投げて、最後に何か役立つものが出てくることを期待することはしません。これは、非常に説得力のあるデモ、大きな差分、および将来の運用インシデントを取得する最も早い方法です。
意味のある機能ごとに、私は問題を非常に具体的にすることから始めます。
私たちは何を構築しているのでしょうか?誰が使っているの？何が壊れるでしょうか？どの API コントラクト、データベース モデル、規約、セキュリティ ルール、パフォーマンスの制約、デプロイメント要件、および既存のアーキテクチャ上の決定をそのままにしておく必要がありますか? (基本的にはプロダクトオーナーと技術リードのような役割を果たします。)
タスクがあいまいであればあるほど、コードが生成される前に多くの作業が発生します。これが最も重要な部分です。何かが明確でない場合は、ガベージイン、ガベージアウト...
この計画は「機能 XYZ を追加する」というものではありません。計画は次のようなものに近いです。
どのモジュールがどのように変更されるか
どのデータ モデルとエンドポイントが影響を受けるか
どの既存のフローが後退する可能性があるか
どの移行、機能フラグ、権限、テスト、テレメトリが必要か
実装開始前の受け入れ基準はどのようなものであるか
エージェントが従う必要があるリポジトリのルールと規約
ここで、リポジトリ固有のルールセットが非常に重要になります。エージェントは、プロジェクトを開始するたびにプロジェクト アーキテクチャを再検出する必要はありません。リポジトリの規則、つまりサービスを構造化する方法、入力を検証する方法、データベースにアクセスする方法、

エラーがどのように処理されるか、テストがどのように書かれるか、コンテナがどのように構築されるか、そして「完了」とは実際には何を意味するのか。私の場合、すべてを .claude フォルダーの下に保存し、それを rules フォルダーと context フォルダーに分けます。
タスクが本当に明確になると、実際のパイプラインが開始されます。
1. 計画: ビジネス要件を実装契約に変える
最初のエージェントは計画エージェントです。GPT 5.5 とは対照的に、Opus 4.8 が適切に機能することがわかりました (ただし、これはケース固有です)。
その仕事はまだコードを書くことではありません。その仕事は、機能を他のエンジニアが推測することなく実行できる実装計画に分解することです。
私は通常、リポジトリの関連領域を検査し、依存関係を特定し、変更される可能性のあるファイルを説明し、リスクを強調し、作業完了後に存在する必要があるテストを定義するように依頼します。
これは私が計画に真剣に取り組む瞬間でもあり、計画が十分に複雑な場合は、何かを行う前に時間をかけて Mermaid でグラフを描画し、より明確に理解することがあります。
共有モデルに触れますか?重要なエンドポイントは変わりますか?データベース内の履歴データに影響を与える可能性がありますか?移行が必要ですか?競合状態、セキュリティ リスク、アプリ全体の速度を低下させる N+1 クエリ、または負荷がかかっている場合にのみ現れる操作上の問題が発生しているのでしょうか?
目標はシンプルです。開発エージェントがトークンを消費してファイルを変更し始める前に、曖昧さを軽減することです。
計画が間違っていると、悪いコードがより速く生成され、最終的には再帰ループで修正することになります。
2. 開発: ビルダーにコンテキスト、ツール、明確な境界を与える
私の場合、開発エージェントは通常 Claude Code です。私は、共同作業者にも主にコード生成にこれを使用することを勧めています。
承認されたプラン、リポジトリ ルール、機能の内容を受け取ります。

ext、および変更が許可される範囲。これは独自の機能ブランチで動作し、タスクの実装を段階的に開始します。
ここで、MCP ツールが非常に価値を持ちます。
モデルは盲目的に動作するべきではありません。実際のエンジニアが使用するであろう次の機能にアクセスする必要があります。
GitHub 、問題、プル リクエスト、履歴、および関連する実装決定を検査します (完全なリポジトリ コンテキストと最新の状態を取得するため)。
ファイルシステムとターミナルへのアクセス。コードベースを探索してコマンドを実行します。
Docker 、現実的な環境でサービスを再現し、変更をテストします。
Context7 またはドキュメント ツール。ライブラリとフレームワークの正確な使用状況を取得します。
Playwright と Chromium 、ブラウザーまたは UI の動作を検証する必要がある場合 (これは主にフロントエンド タスクで使用されます)
テストランナー、リンター、型チェッカー、ビルドツール
皆さんにも、Firecracker を使って迅速な環境を構築することをお勧めします。これは私にとって非常に役立ち、信じられないほど強力なテクノロジーです。 Firecracker は安全なコード実行の業界標準となっています。AWS Lambda はその上に構築されており、多くの AI エージェント プラットフォーム (E2B など) は Firecracker microVM を使用して、AI で生成されたコード実行を安全にサンドボックス化します。信頼できないコードを実行するための分離された安全な環境が必要な場合、Firecracker は、標準のコンテナーではまったく対応できないカーネル レベルの分離を提供します。
重要なのは、エージェントはコードを生成するだけではないということです。たとえば、Karpathy のルールや、従うべきだと思われるその他のルールに従って、調査、実装、テストの実行、障害の修正、実際のリポジトリに対する仮定の検証が行われます。
これは、生成されたスニペットをファイルにコピーして実行を押すワークフローとは大きく異なります。
3. レビュー: 2 番目のモデルは最初のモデルと一致しない
これ

ここで最大の違いが発生し、GPT 5.5 が関係します。
私は開発エージェントに自らの仕事を見直して勝利宣言をさせません。これは、開発者に何の目も持たずに自分のプル リクエストの承認を求めるのと同じです。
機能ごとに、保護された環境で実行アクセス権を備えたレビューアー モードで GPT 5.5 を実行する個別のレビューアー ターミナルを維持します。その仕事は、コードの変更を検査し、元の要件と比較し、リポジトリの影響を受ける領域をチェックして、適切なレビュー レポートを生成することです。
査読者は、非常に厳格な主任エンジニアのように振る舞うように指示されます (はい、FAANG で言います (笑))。
間違ったビジネス ロジックまたは要件の欠落
壊れたエッジケースと無効な仮定
ハッピーパスでのみ機能する脆弱なコード
リポジトリのアーキテクチャと規約への違反
diff がコンパイルされるからといって、レビュアーに「良さそうだ」と言われたくありません。
導入が失敗する可能性がある方法を積極的に模索してほしい。
出力されるのは、曖昧な要約ではなく、構造化されたレビュー レポートです。レポートは開発エージェントに戻され、開発エージェントは問題を修正し、明確な変更の概要を返します。
その後、レビュー担当者が再度レビューします。
これは図の赤いフィードバック ループで、複数回実行されます。
4. PID のようなフィードバック ループ: 実装、エラー信号、修正、安定性
このループは、おそらくアーキテクチャ全体の中で最も重要な部分です。
開発者がシステムを変更する
レビュー担当者は、要件、アーキテクチャ、品質基準、または期待される動作からの逸脱を検出します。
その偏差がエラー信号です
開発エージェントが修正を適用します
レビュー担当者は、特に新しく追加されたコードを再チェックして、修正によって他の場所に別の問題が発生していないかどうかを確認します。
電気や電気の知識を持つエンジニア向け

バックグラウンドを制御するため、 のような PID として説明します。
このシステムは魔法のように自律的ではありません。継続的に出力を測定し、それを望ましい状態と比較し、修正措置を適用するため、より安定します。
このループがないと、エージェント コーディングが非常に速くドリフトしてしまい、非常にひどい結果が生じる可能性があります。
これを使用すると、出力は実際の許容基準に徐々に近づき、ほとんどの場合、想像すらしていなかった修正が得られます。
目標は、エージェントとの無限の会話ではありません。目標は収束です。
レビュー担当者のレポートがクリーンで、テストに合格し、機能が意図した動作を満たしている場合、ループは停止します。そのため、レポートを High 、 Medium 、 Low の問題に分類することもできます。ほとんどの場合、残りが Low を超える問題がない場合は、安全に次のステップに進むことができます。レビュー担当者は、開発エージェントが変更を加えた後に追加されたすべての新しいコードを毎回再チェックして、新しい問題が発生していないことを確認することに留意してください。
5. オーケストレーター: 依然として人間がコントロール プレーンである
現時点では、私は主にオーケストレーターとしての役割を果たしています。
アクティブな機能ごとに、通常、次の 3 つのターミナルを開きます。
Opus 4.8 または GPT 5.5 はレビューアー モードでフルアクセスで実行されますが、前述したように安全な環境で実行されます。調査し、差分をレビューし、レポートを生成し、ループ内の懐疑的なエンジニアとして機能します。
Claude Code は機能を実装し、対象を絞ったテストを実行し、レビュー担当者のフィードバックを受け取り、レビューが合格するまで繰り返します。
これはルート ディレクトリのターミナルです。私は、Git の操作、ログの確認、広範なテスト スイートの実行、障害の調査、インフラストラクチャの検査、変更のコミット、および制御された手動編集にこれを使用しています。
ほとんどの場合、これを tmux 内で実行し、2 つまたは 3 つの機能を持たせることができます。

並行して移動し、それぞれが独自のブランチで分離されます。重大な並列作業の場合、ワークツリーまたはチェックアウトを個別にすると、各エージェントにクリーンな作業ディレクトリがあり、別の機能と衝突しないため、より安全になります。
人間の役割は削除されません。
人間は何が重要かを判断し、範囲を制御し、危険な変更を承認し、コンテキストが欠落しているエッジケースを処理し、マージするのに十分な証拠がシステムにいつあるかを決定します。
エージェントは実行を加速することに非常に優れています。
これらは、所有権や初期および最終的な決定に代わるものではありません。
今でも手動入力に戻る瞬間があります。
これを CAVEMAN MODE = ON と呼びます。
これは通常、編集が小さい場合、非常に機密性が高い場合、または単純に手動で行う方が速い場合に発生します。典型的な例としては、構成ファイル、.env 値、シークレット、デプロイメント変数、1 行のホットフィックス、または前述したように、エージェントにアクセス権を与えると遅くなるか不要になる小さな変更があります。
これはエージェント コーディングの失敗ではありません。
それはエンジニアリング上の適切な判断の一部です。
目標は、AI にリポジトリ内のすべてのファイルをタッチさせることではありません。目標は、目の前のタスクに適切なレベルの自動化を使用することです。
7. 承認とマージ: レビューに合格することは B と同じではありません

[切り捨てられた]

## Original Extract

My real production workflow for shipping features with AI agents: 3 terminals, PID-like feedback loops, independent review, and quality gates. No vibe coding—just controlled engineering that serves real users.

How I Ship Production Features with AI Agents (Not Just LinkedIn post or Promotion PoC Demos) | Tech Oriented Chronicles🔮 of Petros Petros-Tech-Chronicles🔮 Blog Tags Projects About Home Blog Tags Projects About Published on Thursday, July 2, 2026 How I Ship Production Features with AI Agents (Not Just LinkedIn post or Promotion PoC Demos)
Name Petros Savvakis Twitter @PetrosSavvakis
How I Ship Production Features with AI Agents (Not Just LinkedIn post or Promotion PoC Demos)
Agentic coding, harness, vibe coding, AI teammates... and more. Those are only some of the "cool" keywords that come up in 2026's engineering space. All that is cool and it means new things to try, experiment, and tinker with... but what does it really work like, at least from what I have tried in my case, across personal projects and independent collaborations?
Let's first do a step back and tell you that for the last few months I am just an orchestrator and mostly a code reviewer. I rarely write a full script of code as I did in the past. The most common time that I need to write code is going into CAVEMAN MODE = ON (I first heard this term from Charlie Holtz - Conductor CEO via Conductor's manual mode ) and we adopted it. It's the time when we must do manual typing for quick file edits (like config or .env files) when an AI model struggles to complete the task or you don't want to expose .env values to agents.
With all that in mind, let's stop the jargon and deep dive into what seems to get real results. For the last months, across personal projects and independent collaborations, we've been shipping features like crazy (that are truly working and are fully tested—not just vibe-coded running in localhost, but in production serving real users). In order to do that, you have to make a legit pipeline that acts like PID Adaptive Control (for those of you that are electrical engineers, you know what I mean)—a system that is self-calibrated and can generate stable and real production results. Enough said, let's move to the architecture of it.
The Workflow: One Human Orchestrator, Three Terminals, One Controlled Loop
Here is the system that has been working for me.
I do not open one agent, throw a vague prompt at it, and hope that something useful comes out at the end. That is the fastest way to get a very convincing demo, a large diff, and a future production incident.
For every meaningful feature, I start by making the problem extremely concrete.
What are we building? Who is using it? What can break? Which API contracts, database models, conventions, security rules, performance constraints, deployment requirements, and existing architectural decisions must remain intact? (Basically acting like product owner and tech lead.)
The more ambiguous the task is, the more work happens before any code is generated—and this is the most crucial part. If something is not clear, then Garbage In, Garbage Out...
The plan is not "add feature XYZ". The plan is closer to:
Which modules will change and how
Which data models and endpoints are affected
Which existing flows could regress
Which migrations, feature flags, permissions, tests, or telemetry are required
What the acceptance criteria look like before implementation starts
Which repository rules and conventions the agents must obey
This is where repository specific rulesets become extremely important. An agent should not have to rediscover your project architecture every time it starts. It should know the conventions of the repo: how we structure services, how we validate inputs, how we access databases, how errors are handled, how tests are written, how containers are built, and what "done" actually means. For my case, I save all that under the .claude folder, then I separate it into rules folder and context folder.
Once the task is really clear, the actual pipeline starts.
1. Planning: Turn Business Requirements into an Implementation Contract
The first agent is the planning agent , which I found out that Opus 4.8 does a good job for, in contrast to GPT 5.5 (but this is case specific).
Its job is not to write code yet. Its job is to decompose the feature into an implementation plan that another engineer can execute without guessing.
I usually ask it to inspect the relevant areas of the repository, identify dependencies, describe the files likely to change, highlight risks, and define the tests that must exist once the work is complete.
This is also the moment where I challenge the plan heavily, and sometimes if it is complex enough, I take time to draw the graph in Mermaid to get a more clear understanding before doing anything.
Does it touch a shared model? Does it change a critical endpoint? Could it affect historical data in the database? Does it need a migration? Are we introducing a race condition, a security risk, an N+1 query that slows the whole app, or an operational issue that will only appear under load?
The goal is simple: reduce ambiguity before the development agent starts spending tokens and changing files.
A bad plan produces bad code way faster, and you end up fixing in a recursive loop.
2. Development: Give the Builder Context, Tools, and Clear Boundaries
The development agent is usually Claude Code in my case. I encourage my collaborators to mostly use it for code generation too.
It receives the approved plan, the repository rules, the feature context, and the boundaries of what it is allowed to change. It works in its own feature branch and starts implementing the task incrementally.
This is where MCP tooling becomes extremely valuable.
The model should not be working blind. It needs access to the capabilities that a real engineer would use:
GitHub , to inspect issues, pull requests, history, and related implementation decisions (to get full repo context and latest state)
Filesystem and terminal access , to explore the codebase and run commands
Docker , to reproduce services and test changes in a realistic environment
Context7 or documentation tooling , to retrieve accurate library and framework usage
Playwright and Chromium , when browser or UI behavior needs validation (this is mostly used in frontend tasks)
Test runners, linters, type checkers, and build tools
I encourage everyone also to build quick environments with Firecracker —it has helped me tremendously and it is incredibly powerful technology. Firecracker has become the industry standard for secure code execution: AWS Lambda is built on it, and numerous AI agent platforms (like E2B) use Firecracker microVMs to safely sandbox AI generated code execution. When you need isolated, secure environments for running untrusted code, Firecracker provides kernel level isolation that standard containers simply cannot match.
The important part is that the agent is not just generating code . It is investigating, implementing, running tests, fixing failures, and verifying assumptions against the actual repository, following for example Karpathy's rules or any other rules you think it should obey.
That is a very different workflow from copying a generated snippet into a file and pressing run.
3. Review: The Second Model Is Not There to Agree with the First One
This is where the biggest difference happens and where GPT 5.5 is involved.
I do not let the development agent review its own work and declare victory. That is the equivalent of asking a developer to approve their own pull request with no second pair of eyes.
For each feature, I keep a separate reviewer terminal running GPT 5.5 in reviewer mode with execute access in a protected environment. Its job is to inspect the code changes, compare them against the original requirement, check the affected areas of the repository, and generate a proper review report.
The reviewer is instructed to behave like a very strict principal engineer (yes, in FAANG, LOL).
Incorrect business logic or missed requirements
Broken edge cases and invalid assumptions
Fragile code that works only for the happy path
Violations of repo architecture and conventions
I do not want the reviewer to say "looks good" because the diff compiles.
I want it to actively search for ways the implementation can fail.
The output is a structured review report , not a vague summary. The report goes back to the development agent, which fixes the issues and returns a clear change summary.
Then the reviewer reviews again.
This is the red feedback loop in the diagram, and it runs multiple times.
4. The PID-Like Feedback Loop: Implementation, Error Signal, Correction, Stability
This loop is probably the most important part of the entire architecture.
The developer changes the system
The reviewer detects deviations from the requirement, architecture, quality bar, or expected behavior
That deviation is the error signal
The development agent applies corrections
The reviewer checks again, especially around the newly added code, to make sure the fix did not introduce another issue somewhere else
For engineers with an electrical or control background, this is why I describe it as PID like .
The system is not magically autonomous. It becomes more stable because it continuously measures output, compares it to the desired state, and applies corrective action.
Without this loop, agentic coding can drift very fast and give really sh**ty results.
With it, the output gets progressively closer to the actual acceptance criteria , and most of the times it gives you fixes that you wouldn't have even imagined.
The goal is not infinite agent conversations. The goal is convergence.
When the reviewer report is clean, the tests are passing, and the feature meets the intended behavior, the loop stops. That's why you also categorize the report into High , Medium , and Low issues. Most of the time, when it has not any issues higher than Low remaining, you can safely proceed to the next step. Keep in mind the reviewer every time rechecks all the new code added after the Dev Agent makes changes to make sure no new problems were created.
5. The Orchestrator: The Human Is Still the Control Plane
At this point, I am mostly acting as the orchestrator .
For every active feature, I typically have three terminals open:
Opus 4.8 or GPT 5.5 runs in reviewer mode, with full access, but as we said in a safe environment. It investigates, reviews diffs, generates reports, and acts as the skeptical engineer in the loop.
Claude Code implements the feature, runs targeted tests, receives reviewer feedback, and iterates until the review passes.
This is the root directory terminal. I use it for Git operations, checking logs, running broader test suites, investigating failures, inspecting infrastructure, committing changes, and making controlled manual edits.
Most of the time I run this inside tmux , and I can have two or three features moving in parallel, each isolated in its own branch. For serious parallel work, separate worktrees or checkouts make this much safer because each agent has a clean working directory and does not collide with another feature.
The human role is not removed.
The human decides what matters, controls scope, approves risky changes, handles edge cases where context is missing, and decides when the system has enough evidence to merge.
Agents are very good at accelerating execution.
They are not a replacement for ownership and early and final decisions.
There are still moments where I go back to manually typing.
We call this CAVEMAN MODE = ON .
This usually happens when the edit is tiny, highly sensitive, or simply faster to do manually. Typical examples are configuration files, .env values, secrets, deployment variables, a one line hotfix, or a small change where giving an agent access would be slower or unnecessary, as we said earlier.
This is not a failure of agentic coding.
It is part of good engineering judgment.
The goal is not to make an AI touch every file in the repository. The goal is to use the right level of automation for the task in front of you.
7. Approval and Merge: Passing Review Is Not the Same as B

[truncated]
