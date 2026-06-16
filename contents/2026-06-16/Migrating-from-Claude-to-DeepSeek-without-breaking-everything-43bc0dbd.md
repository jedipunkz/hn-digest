---
source: "https://blog.firetiger.com/migrating-from-claude-to-deepseek-without-breaking-everything/"
hn_url: "https://news.ycombinator.com/item?id=48559587"
title: "Migrating from Claude to DeepSeek without breaking everything"
article_title: "Migrating from Claude to DeepSeek without breaking everything"
author: "matsur"
captured_at: "2026-06-16T19:19:31Z"
capture_tool: "hn-digest"
hn_id: 48559587
score: 5
comments: 0
posted_at: "2026-06-16T18:16:45Z"
tags:
  - hacker-news
  - translated
---

# Migrating from Claude to DeepSeek without breaking everything

- HN: [48559587](https://news.ycombinator.com/item?id=48559587)
- Source: [blog.firetiger.com](https://blog.firetiger.com/migrating-from-claude-to-deepseek-without-breaking-everything/)
- Score: 5
- Comments: 0
- Posted: 2026-06-16T18:16:45Z

## Translation

タイトル: すべてを壊すことなく Claude から DeepSeek に移行する
説明: Firetiger は、お客様に代わってインシデントを調査し、展開を監視し、テレメトリを調査する一連のエージェントを実行しています。これらのエージェントはどれも、本質的には LLM を巡るループです。つまり、最大の売上原価は推論です。私たちが選んだモデルは、

記事本文:
購読する
すべてを壊さずに Claude から DeepSeek に移行する
Firetiger は、お客様に代わってインシデントを調査し、展開を監視し、テレメトリを調査する一連のエージェントを運営しています。これらのエージェントはどれも、本質的には LLM を巡るループです。つまり、最大の売上原価は推論です。エージェントを強化するために私たちが選択したモデルは、提供できる製品エクスペリエンス、顧客に請求する必要がある金額、増加した負荷と規模への対処方法など、ビジネス全体に影響を与えます。
さまざまなフレーバーとヴィンテージのクロードが、主に AWS Bedrock を通じて提供され、過去 1 年間私たちが選んだモデルでした。最近、採用（およびコスト）が増加しているため、オープンソースの低コスト モデルへの移行をより真剣に検討しました。トークンエコノミクスを改善する方法を見つけることで、変更モニターやサービスモニターなどの成功を基にして、さらに多くの製品エクスペリエンスを構築できるようになります。
本番環境で Claude から DeepSeek に切り替えると、移行された最初の 3 つのエージェント タイプの実質支出額が 62% 削減され、Anthropic での年間約 606,000 ドルから、現在の DeepSeek での年間 231,000 ドルに増加しました。
最初の調査の後、私たちは安価な日付として DeepSeek v4 モデル ファミリに焦点を当てました。
モデル s/claude/deepseek/g を交換するだけです
この変更に対する単純な計画: API クライアントを別のエンドポイントに向け、多額の費用を節約し、家に帰ります。
シンプルなチャット製品であれば、その計画を回避できるかもしれません。ああ、ああ。当社のエージェントは、(時間的にも範囲的にも) 長期間にわたり、数十回のツール呼び出しを伴う多段階の調査を実行します。モデル間の小さな動作の違いは時間の経過とともに増大し、エージェントの全体的な軌跡と製品の品質に大きな影響を与えます。
私たちは実験に必要な 3 つの仮説を立てました

正常に移行できることが証明されました。
DeepSeek モデルは、膨大な労力を費やすことなく、私たちのタスクを実行できます。重要な指標: タスク完了の精度。
DeepSeek モデルは、タスクあたりのコストに基づいて測定すると、Claude よりも安価になります。
DeepSeek を利用したエージェントはタスクを実行し、Claude を利用したエージェントと同様の方法で動作します。重要な指標: 完了までのステップ、完了までの時間、エラー率。
小規模に始めるために、私たちは最初の実験の範囲を 3 つの代表的なタスクに絞りました。独自の製品の調査を中心としたユーザー定義タスク、コード変更を監視するための計画生成 (「この PR を踏まえて、デプロイメントを監視する計画を立て、それが想定どおりに動作し、何も壊さないことを確認する」)、そして事後の根本原因分析です。
測定精度には良好な評価が必要でした。私たちのものには2つの味があります。静的評価は保存されたデータに対してローカルで実行され、主にモデルが苦労する特定の機能をターゲットとしていますが、ライブ評価データセットは毎週更新されます。古いケースの有効期限が切れると、興味深いエージェント セッションが検出、分析され、データセットに昇格されます。
質問と指標を用意した後、ワークロードのサブセットに対してモデル文字列と API エンドポイントを反転しました。
ステップ 1: 品質ギャップを測定して埋める
まずモデルを交換し、タスク完了評価を実行しました。 DeepSeek 4.0 Pro のスコアは、ベースライン Sonnet 4.6 の 94% に対して、理屈なしで 65%、理屈ありで 80% でした。これらはドロップイン交換品としては十分なものでしたが、出荷可能とは程遠いものでした。初期のデータを考慮すると、DeepSeek で推論を使用することは成功には譲れないと判断しました。
ここから、プロンプトとツールの説明を変更する自己改善ループを実行しました。各ターンと特定された問題を理解することで、DeepSeek とクロードのどこが苦戦しているのかがわかりました。ディー

pSeek は明らかに有能なモデルですが、いくつかのパターンが現れました。
1 つ目: 運用環境に移行する PR の変更を監視する計画を作成する際、コードがもたらす二次的および三次的な影響を見つけるのに苦労しました。これを「ローカル アーティファクトからの非ローカル依存関係の推論」と呼びます。エージェントが差分を確認すると、次のことがわかります。
コードを呼び出すもの (grep を使用)。
彼らには見えていないが、推論する必要があるもの:
このコードが生成するデータを読む人。
このコードが制御するタイミングに誰が依存するのか。
他のシステムがこのコードの動作に関して仮定する不変条件。
クロードはその 2 番目のリストについて何も言われずにアイデアを考える傾向がありましたが、DeepSeek には 2 次効果について注意深く考えるように指示する必要がありました。ループによって計画プロンプトに加えられた実際の変更は次のとおりです。
@@ 「変更点を理解する」セクション
次に:
- 影響を受けるファイルを読んでコンテキストを理解する
- コード パスをトレースして、影響を受けるサービスまたはコンポーネントを特定します
- PRのタイトル、説明文、ユーザーのリクエストに注意する
特定の懸念事項に対して
+
+ テレメトリを調査する前に、次の 2 段階のトレースを完了してください。
+ 1. この変更により何が生成または変更されるかをリストします。
+ 2. それぞれについて、差分に現れず、それに依存するものに名前を付けます。
+ これらのページ外の依存関係が計画です。候補者がチェックした場合
+ diff が接触するもののみを参照します。監視しています
+ 変化ではなくプロデューサー。
2 番目のパターン: DeepSeek は、他のエージェントやユーザーから与えられた追加のコンテキストを無視して、最初から作業をやり直し、タスクのサブセットの優先順位を再設定することがよくありました。あるケースでは、PR の説明にはどの変更が重要であるかが正確に記載されており、DeepSeek は代わりにセクションごとに変更された行に基づいて優先順位を再設定しました。モデルは分析しているアーティファクトのコンテキストを無視し、代わりに粗雑なヒューリスティックを使用することを選択しました。
修正に関係するのは、

コンテキストがどのように提示されるかを構成する。私たちの元のプロンプトはヘッジされていました。それはモデルノートが「権威ではない」と告げ、DeepSeek はそのヘッジを完全に破棄する許可として読みました。修正:
@@ 「利用可能なテレメトリの調査」セクション
- - `read_notes` を使用して、「一般」および「デプロイ」のメモを読んでください。
- システム知識。メモは一般的な学習であり、
- この PR の信頼できる定義。
+ - フロー、メモ、および付属のインジケーターには方向性があります
+ コンテキスト — 人間と以前のエージェントがすでに持っているもの
+ このコードベースについて確立されました。それらをアンカーします。正確な
+ 数値がずれている可能性があります。彼らが描写するものの形
+ 通常はそうではありません。この PR にとって何が重要かを確認しますが、そうではありません
+ すでにコンテキストが何であるかを一から再調査する
+が成立します。
+ - 例外: メモに *異なる* PR が記述されている場合
+ 類似のキーワードは無視してください。注意事項は一般的なものです。このPRは
+具体的。
3 番目: 調査中に、モデルが極小値で行き詰まり、特定の穴に深く潜り、そこで迷子になる一方で、実際の根本原因は別の場所にありました。
これらはどれも特殊な失敗ではありませんが、テスト方法が少数の応答に注目するだけで終わってしまうと、簡単に見落とされてしまうでしょう。
これらの修正を適用すると、DeepSeek v4 Pro の推論機能は評価セットで 92% のスコアを獲得し、Claude Sonnet 4.6 のベースライン 94% と同等になりました。重要なことに、プロンプトの変更によって Sonnet 4.6 のパフォーマンスは変化せず、特別なケースや IF MODEL == X DO Y タイプのロジックを使用せずにモデル ファミリ間を移行できるようになりました。
この時点で、DeepSeek が必要なことを確実に達成できることが証明されました。次に、どのようにして、どのような費用がかかり、どのような経路でそこに到達するのでしょうか?
タスクあたりのコストのデータを詳しく調べてみると、モデル間のトークンあたりの価格設定ほどは数値が下がっていませんでした。

s.次の 2 つの理由が明らかになりました。
DeepSeek は途中で自身の出力を膨らませる傾向があり、自分自身にメモを書き、大規模な計画とより大きな結論を生成しました。
トークンの消費および発行方法がモデルごとに異なることに加えて、プロンプトを最適化する過程で、入力トークンと出力トークンの両方の使用量が増加しました。
プロンプト最適化のもう 1 つのラウンドと、主要な出力のトークン数の検証により、数値が当初想定していた支出曲線と一致するように戻りました。
次に、タスクを完了するために必要な時間とステップ数に焦点を当てました。ある例では、RCA タスクに関する DeepSeek 調査は正しい結論 (評価で満点) に達しましたが、時間は 4 倍、ステップ数は 2 倍かかりました。ツールの使用状況を見ると、DeepSeek は Claude よりもはるかに頻繁にインターネット検索にアクセスしており、この場合はウサギの穴の中で Go ドキュメントを読むのに 10 分を費やしましたが、そのほとんどは最終的な根本原因とは何の関係もありませんでした。評価スコアだけを見ていたら、著しく劣った製品エクスペリエンスを出荷していたでしょう。
エラー率を比較すると、幻覚のようなツールの引数、使用されるツールの奇妙な選択、クロードから見たことのないタスク戦略など、さらに多くの問題が見つかりました。
モデルとツールの呼び出し動作の変更により、エージェントの認証戦略とポリシーも再検討する必要がありました。以前は、クロードの行動や傾向に暗黙の依存関係を設定していました。モデルを切り替えることにより、サンドボックス化と分離のプリミティブに多層防御を追加する必要がありました。
最終チェックとして、DeepSeek でユーザー セッションのサブセットを再生し、品質ではなく動作の違いを分析して、他に問題が発生しないことを確認しました。
私たちのテストベンチでは、モデルの動作を予測可能に保ち、品質を高く保ち、コストを大幅に削減しました。

冷静に、私たちは実験が成功したと判断しました。つまり、運用エージェントのオペレーションの一部を DeepSeek に移行し、品質、コスト、その他の指標を継続的に測定することができました。
まず、キャッシュ ヒット率が低いエージェントをクロード上で移行することから始めました。これは、非常に安価なキャッシュ トークンの読み取りが行われるヒット率の高いエージェントと比較して、トークンあたりの価格設定のスプレッドがより意味があるためです。エージェントの Anthropic プロンプト トークンの少なくとも 30% がキャッシュ書き込みであり、同じプロンプトが生成するトークンが V4 Pro の 2 倍未満である場合は、移行の適切な候補となります。
運用環境では、移行された最初の 3 つのエージェント タイプで実質支出が 62% 削減され、Anthropic での年間約 606,000 ドルから、現在 DeepSeek での年間 231,000 ドルまで削減されました。
私たちが予想していなかった問題が 1 つあります。それは、DeepSeek を使用したキャッシュのチューニングが、本番環境で私たちが思っていたよりも複雑だったということです。これら 3 つのエージェントは、Bedrock + Claude でキャッシュ ヒット率が最も低かったため、適切な候補でした。 DeepSeek では、現在 30 ～ 40% のヒット率が確認されており、キャッシュが Anthropic で確認されたものと一致する場合、同じワークロードは 70% 削減の 181,000 ドル/年で実行されることになります。私たちはそのギャップを埋めるために、推論プロバイダーである Baseten と協力しています。
この移行を自分で検討している場合、私たちのアドバイスを一文で表すと次のようになります。評価スコアは作業の始まりであり、終わりではありません。正しい答えに到達すると、移行が可能であることが確実にわかります。自分が理解した道をたどり、効率的にそこに到達することがゴールです。
前の投稿
Firetiger の新しくシンプルな価格設定
Firetiger ブログを購読する
jamie@example.com
購読する
ファイヤータイガーのブログ © 2026
ホーム

## Original Extract

Firetiger runs a fleet of agents that investigate incidents, monitor deployments, and dig through telemetry on behalf of our customers. Every one of those agents is, at its core, a loop around an LLM, which means our single biggest cost of goods sold is inference. The models we choose to

Subscribe
Migrating from Claude to DeepSeek without breaking everything
Firetiger runs a fleet of agents that investigate incidents, monitor deployments, and dig through telemetry on behalf of our customers. Every one of those agents is, at its core, a loop around an LLM, which means our single biggest cost of goods sold is inference. The models we choose to power our agents have impacts across our business: on the product experiences we can deliver, how much we need to charge customers, and how we think about handling increased load and scale.
Claudes of various flavors and vintages have been our models of choice for the past year, primarily served through AWS Bedrock. Recently, with adoption (and costs) growing, we more seriously investigated migrating to an open source, lower cost model. Finding a way to improve token economics allows us to build even more product experiences, building on the success of things like Change Monitors and Service Monitors.
Switching from Claude to DeepSeek in production, we realized 62% reductions in real dollar spend for the first three agent types migrated, going from roughly $606K/yr on Anthropic to $231K/yr on DeepSeek today.
After initial explorations, we put our focus on the DeepSeek v4 model family as our cheap(er) date.
Just swap the model s/claude/deepseek/g
A naive plan for this change: point our API client at a different endpoint, save a bunch of money, go home.
In a simple chat product you might even get away with that plan! Alas. Our agents run long (both in time and scope), multi-step investigations with dozens of tool calls. Small behavioral differences between models compound over time, and would have huge impact on overall agent trajectories and product quality.
We formed three hypotheses the experiment needed to prove to successfully migrate:
DeepSeek models are capable of our tasks without a heroic amount of effort. The metrics that matter: Task completion accuracy.
DeepSeek models will be cheaper than Claude, measured on a cost per task basis.
DeepSeek powered agents will accomplish their tasks and behave in ways similar to Claude powered agents. The metrics that matter: Steps to completion, time to completion, and error rate.
To start small, we scoped our initial experiment to three representative tasks: user-defined tasks centered on exploring their own product, plan generation for monitoring code changes ("given this PR, come up with a plan to monitor the deployment and make sure it does what it's supposed to and doesn't break anything"), and post-hoc root cause analysis.
Measuring accuracy required good evals. Ours come in two flavors. Static evals run locally against stored data and mostly target specific capabilities where we've seen models struggle, while living eval datasets update every week: interesting agent sessions get detected, analyzed, and promoted into the dataset as older cases expire.
With questions and metrics in place, we flipped our model string and API endpoint for a subset of workloads.
Step one: measuring and closing the quality gap
We started by swapping models and runing task completion evals. DeepSeek 4.0 Pro scored 65% without reasoning and 80% with it, against our baseline Sonnet 4.6's 94%. These were respectable for a drop-in replacement, but nowhere near shippable. Given the early data, we decided using reasoning with DeepSeek was non-negotiable for success.
From here, we ran a self-improvement loop that modified the prompt and tool descriptions. Understanding each turn and identified issue told us where DeepSeek struggles vs Claude. DeepSeek is clearly a capable model, but a few patterns showed up.
First: when creating a plan to monitor PR changes going to production, it had trouble finding the secondary and tertiary effects the code would have. Call it "inferring non-local dependencies from a local artifact". When our agents look at a diff, they see:
What calls the code (with grep).
What they don't see, but need to reason about:
Who reads the data this code produces.
Who depends on the timing this code controls.
What invariants other systems assume about this code's behavior.
Claude tended to ideate about that second list unprompted, while DeepSeek needed to be told to think carefully about second order effects. Here's the actual change the loop made to the planning prompt:
@@ "Understand the changes" section
Then:
- Read affected files to understand context
- Trace code paths to identify what services or components are affected
- Pay attention to the PR title, description, and user's request
for specific concerns
+
+ Before researching telemetry, complete this two-step trace:
+ 1. List what this change produces or alters.
+ 2. For each, name what depends on it that does NOT appear in the diff.
+ Those off-page dependents are the plan. If a candidate check
+ only references things the diff touches, you are monitoring
+ the producer, not the change.
Second pattern: DeepSeek would often ignore additional context it was given, from other agents or from the user, and start from scratch, redoing work and reprioritizing subsets of the task. In one case, the PR description spelled out exactly which changes mattered, and DeepSeek reprioritized based on lines changed per section instead. The model ignored context in the artifact it was analyzing, and chose to use a crude heuristic instead.
The fix involved reframing how context is presented. Our original prompt was hedged: it told the model notes were "NOT authoritative", and DeepSeek read that hedge as permission to discard them entirely. The fix:
@@ "Research available telemetry" section
- - Use `read_notes` to read the "general" and "deploy" notes for
- system knowledge. Notes are general learnings, NOT the
- authoritative definition of this PR.
+ - Flows, notes, and attached indicators are directional
+ context — what humans and prior agents have already
+ established about this codebase. Anchor on them. The exact
+ numbers may have drifted; the shape of what they describe
+ usually hasn't. Verify what matters for this PR, but don't
+ re-investigate from scratch what the context already
+ establishes.
+ - The exception: if a note describes a *different* PR with
+ similar keywords, ignore it. Notes are general; this PR is
+ specific.
Third: during investigations, the model would get stuck in local minima, diving deep into a particular hole and getting lost there while the actual root cause sat elsewhere.
None of these are exotic failures, but they’d be easy to miss if the test methodology stopped at eyeballing a few responses.
With these fixes in place, DeepSeek v4 Pro with reasoning scored 92% on our eval set, on par with Claude Sonnet 4.6's 94% baseline. Critically, the prompt changes did not change Sonnet 4.6's performance, allowing us to transition between model families without further special cases or IF MODEL == X DO Y type logic.
At this point, we’d proven to ourselves that DeepSeek can reliably accomplish what we need it to. Next: how does it get there, at what cost, and by what path?
Digging into cost per task data, the numbers hadn't dropped nearly as far as the per-token pricing spread between the models. Two reasons became apparent:
DeepSeek had a tendency to inflate its own output along the way, writing notes to itself and producing large plans and larger conclusions.
In addition to per-model differences in how tokens are consumed and emitted, in the process of optimizing our prompts, we'd increased both input and output token usage.
Another round of prompt optimization, plus validation on token count for key outputs, brought numbers back in line with the spend curve we originally envisioned.
We next focused on time and number of steps required to complete tasks. In one example, a DeepSeek investigation on an RCA task arrived at the right conclusion (with full marks from the eval) but took 4x as long and 2x as many steps. Looking at tool usage, we found that DeepSeek reached for internet lookups far more often than Claude did, and in this case spent ten minutes in a rabbit hole reading Go documentation, most of which had nothing to do with the ultimate root cause. If we'd only looked at eval score, we would have shipped a markedly poorer product experience.
Comparing error rates caught more issues: hallucinated tool arguments, odd choices in which tools got used, and task strategies we'd never seen from Claude.
Changing models and tool calling behavior also caused us to re-examine our agent authorization strategy and policy. Previously, we'd taken implicit dependencies on Claude's behaviors and tendencies. Switching models forced us to add defense in depth to our sandboxing and isolation primitives.
As a final check, we replayed a subset of user sessions on DeepSeek and analyzed differences in behavior, rather than quality, to make sure nothing else would cause issues.
In our test benches, we’d kept model behavior predictable, kept quality high, and cut cost significantly, so we called the experiment a success. That meant we were good to move a subset of our production agent operations to DeepSeek, continuing to measure quality, cost, and other metrics as we went.
We started by migrating agents that ran with poor cache hit rates on Claude, because the per token pricing spread is more meaningful there vs agents with high hit rates that see very cheap cached token reads. If at least 30% of an agent's Anthropic prompt tokens are cache writes, and the same prompts produce fewer than 2x the tokens on V4 Pro, it's a good candidate to migrate.
In production, we realized 62% reductions in real dollar spend for the first three agent types migrated, going from roughly $606K/yr on Anthropic to $231K/yr on DeepSeek today.
One wrinkle we didn't anticipate: cache tuning with DeepSeek was more complicated in prod than we’d thought. These three agents were the right candidates partly because they had our lowest cache hit rates on Bedrock + Claude. With DeepSeek, we're currently seeing 30-40% hit rates, and if caching matched what we saw on Anthropic, the same workloads would run at $181K/yr, a 70% reduction. We're working with our inference provider, Baseten, to close that gap.
If you're considering this migration yourself, the one-sentence version of our advice is this: the eval score is the start of the work, not the end. Getting to the right answer reliably tells you the migration is possible; getting there efficiently, by a path you understand, is the finish line.
Previous post
New, simpler pricing for Firetiger
Subscribe to The Firetiger Blog
jamie@example.com
Subscribe
The Firetiger Blog © 2026
Home
