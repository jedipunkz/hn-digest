---
source: "https://navels.dev/blog/neal/"
hn_url: "https://news.ycombinator.com/item?id=49212155"
title: "Neal: Codex and Claude in a loop shipped a 549-commit migration"
article_title: "neal: coordinating different models on complex coding projects"
author: "navels"
captured_at: "2026-08-07T15:44:23Z"
capture_tool: "hn-digest"
hn_id: 49212155
score: 1
comments: 0
posted_at: "2026-08-07T15:40:19Z"
tags:
  - hacker-news
  - translated
---

# Neal: Codex and Claude in a loop shipped a 549-commit migration

- HN: [49212155](https://news.ycombinator.com/item?id=49212155)
- Source: [navels.dev](https://navels.dev/blog/neal/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T15:40:19Z

## Translation

タイトル: ニール: ループ内のコーデックスとクロードが 549 コミットの移行を出荷
記事のタイトル: ニール: 複雑なコーディング プロジェクトでのさまざまなモデルの調整
説明: neal は、各シート (プランナー、コーダー、レビュー担当者) に異なるモデルを配置し、スコープ付きループでそれらを調整するオープンソースのコーディング エージェントです。

記事本文:
ニール: 複雑なコーディング プロジェクトでさまざまなモデルを調整する navels.dev
Lee Nave on LinkedIn 2026 年 8 月 4 日 ニール: 複雑なコーディング プロジェクトでさまざまなモデルを調整する
neal の最初のバージョンは計画ファイルであり、次のプロンプトが表示されました。
@plans/EMBER_MIGRATION.mdを実行します。続けて。ブロックされない限り止まらないでください。
仕事でフロントエンド フレームワークの大規模なアップグレードを行っているときに、それを Codex に入力しました。当時の現行モデルはGPT-5.4でした。私は、数か月にわたる退屈な作業を任せて、定期的にチェックインし、それ以外の場合はそのまま動かし続けたかったのです。アレンジが崩れるたびに何かを追加していきました。それは最終的に、計画に基づいてコーディング エージェントを調整するローカル CLI である neal になりました。
私たちのフロントエンド コードベースは、長い間 Ember.js v3 から v5 (ああ… v6) へのアップグレードが予定されていました。これは、数千のファイル、大幅な構文の変更、非推奨のパターンの置き換えなど、事実上すべてのコードに影響を与える可能性があります。編集のほとんどは繰り返しでしたが、簡単にスクリプト化できるものではありませんでした。一部の共有コンポーネントは完全に作り直す必要があります。手作業による作業は 300 時間以上かかると見積もられました。かなり包括的なテスト スイートを使用したとしても、このプロジェクトには重大なリスクが伴いました。
無駄のないエンジニアリングチームでは、それは不可能な仕事のように思えました。コーディングエージェントが登場するまでは。 2025 年半ば、私は興奮して Claude Sonnet 4 をこのタスクに投入しましたが、それがテスト スイートで延々とモグラたたきをし、同じバグを何度も修正したり引き起こしたりするのを見るだけでした。当時必要とされたベビーシッターの量は、平凡な結果を正当化するには多すぎました。移行は棚上げされました。
今年の初め、私たちは再び挑戦する時期が来たと感じていました。 Codex モデルと Claude モデルは、計画、コーディング、および自律的な作業において大幅に向上しました。詳細な移行計画のドキュメントがあればよかったのですが、

フロンティア コーディング エージェントの 1 人が単独で段階的に移行作業を進め、私は途中で進捗状況を確認することができました。
移行計画には、Ember アップグレード リソースへのリンク、テスト スイートを実行するためのコマンド、および一度に 10 ～ 20 個の関連ファイルを処理するための指示が含まれていました。また、まだ移行する必要があるファイルのリストも保持しているため、Codex は中断されたときに中断したところから再開できます。
しばらくの間、裸のプロンプトは驚くほどうまく機能しました。 Codex は長期間にわたって単独で動作する可能性があります。再起動しても、計画には、何が完了したか、次に何が行われるか、残りの作業にどの制約が適用されるかが表示されました。
問題はそれを続けることでした。 Codex はバッチを終了し、明確な理由もなく停止していました。 「なぜ移行を繰り返すのをやめたのですか?」と私は尋ねます。何時間も何もしていなかったにもかかわらず、返事は「しませんでした」ということが多かった。
長時間のセッションにより、2 つ目の問題が明らかになりました。 Codex は、以前に従った指示から徐々に離れていきました。 「ファイルの次のバッチを開始する前に、この移行計画を再読する」ことを計画に追加しましたが、それだけでは必ずしも十分ではありませんでした。私はこの動作に「コンテキスト腐敗」という名前を見つけました。長いセッションでは、コンテキストが拡大または圧縮されるにつれて、古い指示や決定は影響力を失います。
私は Codex 内で両方の問題を解決しようとしました。私はプロンプトを、Stop フック、ファイルのバッチが完了したか、ブロッカーが発生したことを示すテキスト マーカー、および各バッチの前に /new 経由でエージェント コンテキストをリセットする指示を備えた $work-autonomously スキルに変更しました。このスキルは役に立ちましたが、Codex はまだ新しいコンテキストを開始することを常に覚えているわけではありませんでした。
作業のすべての部分で新しいコンテキストを開始するという指示に常に従うとは限りません。指示するノード アプリを構築することはどの程度実現可能ですか?

コーデックスで大量の作業を行うには?
これは、Neal の前身です。Codex SDK を使用して移行をループで実行するノード スクリプトです。
増分移行コミットを確認すると、別の問題が明らかになりました。コードの品質は必ずしも最高ではありませんでした。私は、Codex に何かを実装させ、Claude にレビューしてもらい、Codex が Claude のフィードバックに応答するという、私の共通のワークフローを組み込むことにしました。両方のモデルが同意するまでそれを繰り返し、その後、変更を自分でレビューします。
最新のコミットを確認するために、Claude SDK を使用してこれをノード スクリプトに追加しました。スクリプトはクロードのフィードバックを収集し、Codex に送信しました。この時点で、オーケストレーターに「neal」という名前を付けることにしました。
この名前は、内部応力が安定するまで制御されたサイクルで材料を加熱および冷却するプロセスであるアニールに由来しています。 「an」を省略してプロセスを続行しました。
neal は、計画ドキュメントに関してプランナー、コーダー、およびレビュー担当者の役割を調整するローカル CLI です。役割ごとにプロバイダーとモデルを選択します。最近では、Anthropic の Fable をコーダーとして、OpenAI の Sol をレビューアーとして使用しています。
neal run は通常のワークフローです。プランナーとレビュー担当者を通じて大まかな計画を送信することから始まります。彼らはそれに実行の形を与え、より大きな作業をスコープに分割し、実装アプローチ、検証、成功条件を記入します。ファイルレベルの検出とローカル実装の選択肢はコーディング段階に残ります。実装を開始する前に洗練されたプランを読み取ったり編集したりする場合は、neal plan と nealexecute を別々に使用できます。
実行中、コーダーは新しいコンテキストで各スコープを開始し、作業を実装して検証し、それをコミットします。読み取り専用レビューアはスコープ全体でコンテキストを保持し、そのスコープに対して各コミットをチェックします。調査結果はコーダーに戻されます

査読者が作品を受け入れるまで。
実行が行き詰まった場合、ニールはレビュー担当者モデルに制限付きのコンサルタント ターンを依頼できます。これは、ブロックされたコーダー、ストールしたコーダー/レビューアー ループ、または不適切なスコープ分割をカバーします。コンサルタントは問題を診断し、コーダーに具体的な指示を与えることはできますが、コードを編集したり、検証を放棄したりすることはできません。安全に進む方法がない場合、ニールは人間の指示を求めます。スコープが大きすぎることが判明した場合は、サブプランに分割することもできます。
すべてのスコープが受け入れられた後、コーダーとレビュー担当者は、完全な実装と計画に対して最終パスを 1 回実行します。残りの調査結果は、同じ実装とレビューのループを経て戻ります。 neal は、ユーザーが指示しない限り、スコープのコミットを潰します。実行状態、トランスクリプト、およびレビュー成果物は .neal/ の下に存在するため、中断された実行は最後に記録されたフェーズから再開できます。
neal には、Codex と Claude 用のネイティブ アダプターに加えて、OpenRouter などのサービスや Ollama や vLLM などのローカル エンドポイント用の OpenAI 互換アダプターがあります。ネイティブ アダプターは、既存の Codex および Claude サブスクリプション認証を使用できます。私がこの方法で構築したのは、プランナー/コーダー/レビューアーの長いループが、トークンごとに支払うよりもサブスクリプション プランの方がはるかに安くなる可能性があるためです。また、幸運なことに、私は Codex プランと Claude Max プランの両方を払い戻してくれる会社で働くことができました ;-)
ニールは移行を完了しました。移行ブランチは 549 のコミットで着陸し、3,000 を超えるファイルに触れました。また、約 13,000 行のテストコードも追加されました。顧客の受け入れテスト環境で数週間にわたって調理した後、先週の日曜日に本番環境に出荷しました。
数日かかったと言いたいところですが、移行を進めながらオーケストレーターを構築していたので、プロジェクトの実際の実行時間はよくわかりません。それも

完成までに約 1 か月かかりましたが、私の作業のほとんどはニールに関するものでした。移住そのものよりも、それがずっと楽しかったと言えます。
また、Codex+Claude セットアップが標準のコーディング ベンチマークに対してどのように動作するかについても興味がありました。まだ信号があり、カスタム オーケストレーターをサポートできる認識されたベンチマークを見つけるために、少しウサギの穴に落ち込んだ後、最終的に SWE-bench Pro に落ち着きました。
Codex (GPT-5.5) だけでは解決できなかった 105 件のケースから始めて、3 つの役割すべてで Codex を使用してニールを実行しました。その結果、これまで不合格だった8件のケースが合格しました。次に、コーデックスの計画/コーディングとクロード (Opus 4.8) のレビューを伴うニールを通じて同じケースを実行しました。その結果、15件が合格となりました。私はもっ​​と良い結果を期待していましたが、この実験は私が個人的に便利だと思うニールの部分の価値を示しました。コード作成者とレビュー担当者の間の実装ループ。そしてレビューアーの役割には別のモデルがいます。ハーネス、方法論、およびケースごとの結果は、neal-swebench リポジトリにあります。
neal で非常に多くの反復 (すべてのベンチマークと Ember 移行) を取得したことによる素晴らしい副産物の 1 つは、実行アーティファクトの収集でした。 LLM チャットの記録、計画の改訂、各フェーズで必要なループの数などの記録がありました。定期的に (もちろんコーディング エージェントを使用して) 最近の実行で失敗や非効率性を分析することができました。これらの発見は、neel 自身の設計を改善するためにフィードバックされました。
モデルはプロンプトに答えても、ニールの役割の 1 つを実行できない場合があります。 neal は、構造化された出力、ツール呼び出し、ファイル検査、および特定のレビュー判定を期待します。それ以外の点では有能なモデルの多くは、これらのプロトコルのいずれかで失敗します。
neal compat は、コーダーに対して小規模で決定論的なタスクを実行します。

r、およびプランナーの役割と、合格または不合格のレポート。合格とは、モデルがニールのプロトコルを理解できることのみを意味し、計画、コーディング、またはレビューが得意であることを意味するものではありません。
このプロジェクトの一環として、各候補の既知の良好なパートナーとして GPT-5.5 を備えた Codex を使用して、90 の OpenRouter モデルに対する互換性チェックを実行しました。 44人がすべての役職に合格した。日付付きの結果と正確な失敗数は、リポジトリの互換性リストに記載されています。
neal compat は、neal 自体のバグも捕らえました。私のバリデーターの厳しすぎるルールにより、5 つの主力モデルがまったく同じ方法で失敗しました。もう 1 つのレビュー担当者の失敗グループは、Codex の読み取り専用サンドボックスをブロックする Ubuntu 24.04 AppArmor によって発生しました。どちらの場合も、モデルではなくハーネスが間違っていました。
npm install -g @navels/neal
ソースとドキュメントは GitHub にあります。役に立ったと思われる場合は、MusiCares への寄付をご検討ください。これは、経済的、個人的、または医療的危機の際にミュージシャンに頼れる場所を提供する非営利団体です。 LinkedIn でも私を見つけることができます。

## Original Extract

neal is an open-source coding agent that puts a different model in each seat (planner, coder, reviewer) and coordinates them in a scoped loop.

neal: coordinating different models on complex coding projects navels.dev
Lee Nave on LinkedIn Aug 4, 2026 neal: coordinating different models on complex coding projects
The first version of neal was a plan file and this prompt:
Execute @plans/EMBER_MIGRATION.md. keep going. don't stop unless you are blocked.
I typed that into Codex during a large frontend framework upgrade at work. GPT-5.4 was the current model at the time. I wanted to hand it several months of tedious work, check in periodically, and otherwise let it keep moving. Each time the arrangement broke down, I added something. That eventually became neal, a local CLI that coordinates coding agents around a plan.
Our frontend codebase had been due for an upgrade from Ember.js v3 to v5 (ahem… v6) for a long time. This would impact virtually all of our code: thousands of files, significant syntax changes, and replacement of deprecated patterns. Most of the edits were repetitive, but not easily scriptable. Some shared components would have to be completely reworked. We estimated the manual work to be over 300 engineering hours. Even with our reasonably comprehensive test suite, the project carried significant risk.
With a lean engineering team, it seemed like an impossible task. Until coding agents came along. In mid-2025, I excitedly put Claude Sonnet 4 on the task, only to watch it play whack-a-mole endlessly with our test suite, fixing and causing the same bugs over and over. The amount of babysitting required at the time was too much to justify the mediocre results. The migration was shelved.
Earlier this year, we felt like it was time to try again. Codex and Claude models had improved significantly at planning, coding, and working autonomously. I hoped that with a detailed migration plan doc, one of the frontier coding agents could work through the migration incrementally on its own, with me checking in to review progress along the way.
The migration plan had links to Ember upgrade resources, commands for running the test suite, and an instruction to work on 10 to 20 related files at a time. It also kept a list of files that still needed to be migrated, so Codex could pick up where it left off when interrupted.
For a while, the bare prompt worked surprisingly well. Codex could work on its own for long stretches. When I restarted it, the plan still showed what was finished, what came next, and which constraints applied to the remaining work.
The problem was keeping it going. Codex would finish a batch and stop for no apparent reason. I would ask, “Why did you stop iterating on the migration?” The reply was often, “I didn’t,” despite having been idle for hours.
Long sessions uncovered a second problem. Codex gradually drifted from the instructions it had followed earlier. I added “Re-read this migration plan before starting on the next batch of files” to the plan, but that wasn’t always enough. I found a name for this behavior: context rot. Over a long session, old instructions and decisions lose influence as the context grows or is compacted.
I tried to solve both problems inside Codex. I turned the prompt into a $work-autonomously skill with a Stop hook, textual markers to signify that a batch of files had been completed or a blocker encountered, and an instruction to reset the agent context via /new before each batch. The skill helped, but Codex still didn’t always remember to start a new context.
you don't always follow the instruction to start a new context with every chunk of work. how feasible would it be to build a node app to direct codex to do a chunk of work?
This was the precursor to neal: a node script using the Codex SDK to execute the migration in a loop.
Reviewing the incremental migration commits exposed another issue. Code quality was not always the greatest. I decided to incorporate a common workflow of mine: have Codex implement a thing, have Claude review it, and have Codex respond to Claude’s feedback. I would repeat that until both models agreed, then review the change myself.
I added this to the node script using the Claude SDK to review the latest commit. The script captured Claude’s feedback and sent it to Codex. At this point I decided to give my orchestrator a name: neal.
The name comes from anneal , the process of heating and cooling a material in controlled cycles until its internal stresses settle. I dropped the “an” and kept the process.
neal is a local CLI that coordinates planner, coder, and reviewer roles around a plan document. You choose the provider and model for each role. These days I am using Anthropic’s Fable as the coder and OpenAI’s Sol as the reviewer.
neal run is the normal workflow. It starts by sending a rough plan through the planner and reviewer. They give it an execution shape, split larger work into scopes, and fill in the implementation approach, verification, and success conditions. They leave file-level discovery and local implementation choices for the coding phase. You can use neal plan and neal execute separately when you want to read or edit the refined plan before starting the implementation.
During execution, the coder starts each scope with a fresh context, implements and verifies the work, then commits it. The read-only reviewer keeps its context across scopes and checks each commit against that scope. Findings go back to the coder until the reviewer accepts the work.
When a run gets stuck, neal can ask the reviewer model for a bounded consultant turn. This covers a blocked coder, a stalled coder/reviewer loop, or a bad scope split. The consultant can diagnose the problem and give the coder specific direction, but it can’t edit the code or waive verification. If there isn’t a safe way forward, neal asks for human direction. A scope that turns out to be too large can also split into a sub-plan.
After every scope is accepted, the coder and reviewer make one final pass over the complete implementation and plan. Any remaining findings go back through the same implementation and review loop. neal then squashes the scope commits unless you tell it not to. Run state, transcripts, and review artifacts live under .neal/ , so an interrupted run can resume from the last recorded phase.
neal has native adapters for Codex and Claude, plus an OpenAI-compatible adapter for services such as OpenRouter and local endpoints such as Ollama or vLLM. The native adapters can use existing Codex and Claude subscription authentication. I built it this way because a long planner/coder/reviewer loop can be much cheaper on subscription plans than paying per token. Also, I am fortunate enough to work for a company that reimburses both my Codex and Claude Max plans ;-)
neal finished the migration. The migration branch landed with 549 commits and touched over 3,000 files. It also added around 13,000 lines of test code. After weeks of cooking in our customers’ acceptance testing environment, we shipped it to production last Sunday.
I would like to be able to say it took a few days, but since I was building out the orchestrator as the migration was ongoing, I don’t really know the true execution time of the project. It took around a month to complete, but most of my work was on neal. I can say that was a lot more fun than the migration itself.
I was also curious about how a Codex+Claude setup would perform against standard coding benchmarks. After going down a bit of a rabbit hole trying to find recognized benchmarks that still had signal and that could support a custom orchestrator, I eventually settled on SWE-bench Pro .
Starting with 105 cases that Codex (GPT-5.5) was unable to solve on its own, I ran neal with Codex in all three roles. That resulted in 8 cases passing that had previously failed. Then I ran the same cases through neal with Codex planning/coding and Claude (Opus 4.8) reviewing. This resulted in 15 cases passing. I had hoped for better results, but the experiment did show the value of the parts of neal I personally find useful: a planning loop that generates a scoped, reviewed plan; an implementation loop between coder and reviewer; and a different model in the reviewer role. The harness, methodology, and per-case results are in the neal-swebench repo.
One nice by-product of getting so many iterations on neal (all the benchmarking and the Ember migration) was the collection of run artifacts. I had LLM chat transcripts, plan revisions, a record of how many loops were needed in each phase, etc. I was able to periodically (using a coding agent, of course) analyze recent runs for failures and inefficiencies. These findings were fed back through neal to improve its own design.
A model can answer a prompt and still be unable to run one of neal’s roles. neal expects structured output, tool calls, file inspection, and specific review verdicts. Plenty of otherwise capable models fail one of those protocols.
neal compat runs small, deterministic tasks against the coder, reviewer, and planner roles and reports PASS or FAIL. A pass means only that the model can speak neal’s protocol, not that it’s good at planning, coding, or reviewing.
As part of this project, I ran the compatibility check against 90 OpenRouter models, using Codex with GPT-5.5 as the known-good partner for each candidate. Forty-four passed every role. The dated results and exact failure counts live in the repo’s compatibility list .
neal compat also caught bugs in neal itself. Five flagship models failed in exactly the same way because of an overly strict rule in my validator. Another group of reviewer failures came from Ubuntu 24.04 AppArmor blocking Codex’s read-only sandbox. In both cases, the harness was wrong rather than the models.
npm install -g @navels/neal
The source and docs are on GitHub. If you find it useful, please consider donating to MusiCares . It’s a non-profit giving musicians a place to turn to in times of financial, personal, or medical crisis. You can also find me on LinkedIn .
