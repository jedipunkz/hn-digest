---
source: "https://cognition.ai/blog/ai-productivity"
hn_url: "https://news.ycombinator.com/item?id=48541302"
title: "Estimating the Productivity of an Autonomous AI Software Engineer"
article_title: "Estimating the Productivity of an Autonomous AI Software Engineer | Cognition"
author: "weltview"
captured_at: "2026-06-15T14:13:55Z"
capture_tool: "hn-digest"
hn_id: 48541302
score: 1
comments: 0
posted_at: "2026-06-15T13:52:40Z"
tags:
  - hacker-news
  - translated
---

# Estimating the Productivity of an Autonomous AI Software Engineer

- HN: [48541302](https://news.ycombinator.com/item?id=48541302)
- Source: [cognition.ai](https://cognition.ai/blog/ai-productivity)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T13:52:40Z

## Translation

タイトル: 自律型 AI ソフトウェア エンジニアの生産性の推定
記事のタイトル: 自律型 AI ソフトウェア エンジニアの生産性の推定 |認知
説明: エンジニアリングのリーダーは、AI が実際にどれだけの価値を提供しているかを知りたいと考えています。私たちは、デビンの人間の労働時間に相当する時間を測定するシステムを構築しました。

記事本文:
自律型 AI ソフトウェア エンジニアの生産性の推定 |認知メニュー ホームを閉じる
自律型 AI ソフトウェア エンジニアの生産性の推定
6 か月前、すべての CTO は、チームが十分なトークンを使用していないことを懸念していました。トークンの使用量とAI支出が急増するにつれて、この傾向は逆転しました。すべてのトークンが真の価値を提供するわけではないため、エンジニアリングのリーダーたちは現在、実際の出力を測定する方法を見つけようとしています。エンジニアリング時間を節約してプロジェクトを加速する人もいます。無駄なセッションや不適切なプロンプトに無駄に費やされる人もいます。
十分な規模の組織の場合、数千のセッションと数十億のトークンにわたる価値を測定することは本質的に不可能です。私たちはこれを AI で自動化することに着手しました。
実際の ROI を予測することは困難です。その理由については後ほど説明します。そこで私たちは、各 Devin セッションが生産的なエンジニアリング時間にどれだけの価値があるかを見積もるという重要なサブタスクに焦点を当てました。
LLM は時間の見積もりが苦手であることで有名なので、これには苦労するだろうと予想していました。ただし、システムを注意深く調整した後、モデルの r l o g r_{log} r l o g は 0.74 0.74 0.74 となり、偏りがないように見えます。個々の予測は完璧ではありませんが、このモデルは集計された合計を推定するために使用するには十分です。また、エンジニアの給与を使用して金額に換算できるため、ビジネス価値に近づくことができます。
私たちのシステムでは、エージェントが完了した各 Devin セッションをレビューします。まず、有用な出力が生成されたかどうかを分類し、次に人間のエンジニアが同じ作業を生成するのにかかる時間を推定します。私たちは人間のエンジニアに同じタスクにどのくらいの時間を費やすかを尋ねることによってそれを検証しました。
このシステムは現在、顧客とともに稼働しています。私たちの知る限り、これは生産現場で AI エンジニアリングの生産性を測定する初の自動化システムです。
どうやって

生産的なエンジニアリング時間を基準にしているのでしょうか?私たちが答える必要がある最初の質問は、何を推定するかということでした。理想的には、出荷された機能による収益やバグ修正によって回避されたコストなど、金額への影響を直接測定することになります。実際には、これは私たちの分野ではまだ未解決の問題です。エンジニアにとって、先週出荷された PR によって何ドルのビジネス価値が生み出されたかを把握するのは非常に困難です。
スペクトルの対極では、コード行、コミット、PR、消費されたトークンなど、生のアクティビティを測定できます。これらは集めるのは簡単ですが、努力に見合ったものではありません。機械的なリファクタリングでは、午後に何千行もの作業を行うことがあります。 2 行のバグ修正は数時間の調査に相当する場合があります。多くの貴重なタスク (バグの優先順位付け、分析クエリの実行、コードのレビュー) では、コードはまったく生成されません。
私たちが測定することに決めた中間点は、人間工学の時間です。つまり、人間のエンジニアが同じ成果を生み出すのにどれくらいの時間がかかるでしょうか?組織がエンジニアリング作業を評価する方法はすでに時間によって決まり、給与や請負業者の料金は時間によって決まります。リーダーが投資を評価するときは、時間とコストの節約の観点から考えます。時間はビジネスの状況に関係なく組織全体で標準化されており、エンジニアリングレートによってドルに換算できます。
しかし、すべての時間が同じというわけではありません。たとえば、セッションによって作成されたすべての PR が閉じられた場合、その PR は価値がなくなる可能性があります。私たちは生産的なエンジニアリング時間のみを測定したいと考えていました。そのため、セッションが実際に生産的であったかどうかを分類するシステムも構築する必要がありました。
私たちは、Devin ユーザーに最近の代表的なセッションをレビューしてもらい、Devin なしで完了した各セッションにかかる時間を推定するよう依頼して、真実のデータセットを収集しました。私たちのデータセットは 258 258 258 のセッションで構成されています。

さまざまな企業顧客にわたる 126 126 126 人のユーザー。ライブインタビューとアンケートを通じてデータを収集しました。
すべての Devin セッションには、ユーザーのリクエスト、実行されたすべてのアクション、結果のコード、コードベースのコンテキストなど、完全な実行トレースがあります。これにより、調査、集計アクティビティ指標、またはオープンソース ベンチマークだけでは取得するのが難しい詳細レベルでの生産エンジニアリング作業の記録が得られます。
以下のグラフでは、データセットを分析しています。私たちのデータセットは、さまざまな言語、フレームワーク、セッション タイプ、推定時間にわたる、実際の企業のワークロードを表す分布で構成されています。
データセットを確認したところ、すべてのセッションが有用な作業に対応しているわけではないこと、非生産的なセッションを除外するための分類器を構築する必要があることがわかりました。
PR を含むセッションの場合、これは比較的簡単です。セッションからの PR がマージされる場合は、推定値が含まれます。そうでない場合は破棄します。これには少し損失があります。すべてのクローズされた PR を使用したセッションでも生産的な作業を提供できますが、私たちは保守的な誤りを犯したかったのです。
PR のないセッションはさらに複雑になります。非生産的なセッションをフィルターで除外するための分類器を構築しました。これにより、顧客に応じてセッションの約 1 − 20 % 1-20\% 1 − 20% が削除されます。 PR 以外のセッションの多くは真に生産的であり、たとえば、未使用の依存関係の発見、セキュリティ脆弱性のスキャン、プル リクエストのレビュー、分析クエリの実行、バグのトリアージなどであり、私たちはこれらのセッションを保持します。エージェントがタスクを実行するためのアクセス権を持たなかったセッション、エージェントが説明を求めたがユーザーが応答しなかったセッション、および Devin が有意義にタスクを進めることができなかったその他のシナリオは破棄されます。
どのセッションが生産的であったかを特定した後、実際にそのセッションを推定する必要がありました。

同等のエンジニアリング時間。これを行うために、コンテキストとプロンプトという 2 つの主要なコンポーネントを備えた推定エージェントを構築しました。
コンテキスト側では、ユーザーのメッセージ、生成された PR (該当する場合)、完全なエージェント トレース (ログの表示、コードのトレース、lint の修正など)、および DeepWiki からの追加のコードベース コンテキストなど、セッションに関する可能な限り多くの情報を含めました。
迅速な対応として、反復のための開発セットとして 25 25 25 セッションを確保しました。これらのセッションでエージェントの実行を手動でトリアージし、障害モードについて批判的に推論することで、次の設計原則に到達しました。それは、Devin が実際に保存した作業のみをクレジットにし、Devin を保守的な人間の基準と比較するというものです。具体的には、次のことを行います。
エージェントのパスではなく人間のパスに関する理由。エージェント自身の軌跡は、人間の努力の代理としては不十分な場合があります。エージェントは回り道をし、環境やセットアップの失敗から回復し、単独のエンジニアでは作成できない概要レポートなどの成果物を作成します。推定ツールは、再試行、環境設定、エージェントが途中で生成したアーティファクトなどを考慮せずに、人間が同じ出力に到達するまでにたどったであろう経路を推論します。
ユーザーが指定しなかった作品のみをクレジットします。ユーザーが Devin に質問する前にどれだけ理解していたかに応じて、同じ出力でも非常に異なる作業量を表す可能性があります。推定ツールは、ユーザーが実際に言ったことを見て、Devin が自分で解決しなければならなかった問題の量を判断します。たとえば、ユーザーがバグレポートを持ってきて修正案がない場合は、バグのトリアージにかかる時間を含めますが、ユーザーが実装計画を持ってきた場合は実装時間のみをカウントします。
コードベースに精通していることを考慮してください。エンジニアがどのような知識を持っているかに応じて、同じタスクに数分かかる場合もあれば、1 日かかる場合もあります。

オード。これが Devin の最も明白な利点の 1 つであるとユーザーは語ってくれました。開発者にとって立ち上げに 1 日かかるような、不慣れなコードベースやレガシー コードベースでのタスクが、多くの場合、迅速に提供されるということです。推定者は、セッションで明らかになった内容から、関連する親密度のレベルを推測します。ユーザーがシステムの各部分がどのように機能するかを尋ねる場合、その質問には、コード内での方向調整に必要な調査時間が含まれます。どちらにしても信号がない場合は、典型的な知識、つまり高レベルのアーキテクチャは知っているがすべての機能を覚えていないエンジニアを想定します。
関連する専門知識を想定します。私たちのインタビューで繰り返し話題になったのは、エージェントを使用すると、自分たちだけではできなかったこと、つまり以前はフロントエンドやデータ サイエンスの同僚が必要だった作業をバックエンド エンジニアが実行できるようになるということでした。専門分野を超えた影響力を評価すると見積もりが膨らむことになるため、リファレンス エンジニアはタスクに必要な専門知識をすでに持っていると控えめに想定しています。これは、人間が最初になじみのない言語やフレームワークを学ばなければならなかった多くの場合、その労力を過小評価するものです。
233 233 233 セッションのホールドアウト評価セットでは、推定器の r l o g r_{log} r l o g は 0.74 0.74 0.74 、r l o g 2 r_{log}^2 r l o g 2 は 0.54 0.54 0.54 になります。相関は統計的に非常に有意です ( F ( 1 , 231 ) = 279.9 F(1,231) = 279.9 F ( 1 , 231 ) = 279.9 、 p < 10 − 5 p < 10^{-5} p < 1 0 − 5 )。約 50 % 50\% 50% のセッションは、実際の推定値の 2 2 2 の係数内に収まります。個々の推定値にはノイズが多く、どちらの方向でも 2 2 2 – 3 3 3 × の誤差がよく見られます。ただし、誤差はほぼ偏りがなく独立しているため、セッション数が増加するにつれて誤差は相殺され、集計値は人間が報告した合計値に収束します。
ノイズの多くは分散から生じます

ユーザー間では、推定方法と速度の実際の違いの両方が異なります。残りの不一致のおよそ半分は、ユーザー自身のセッション内ではなくユーザー間にあります ( I C C = 0.58 ICC =0.58 I C C = 0.58 )。たとえば、製品内でユーザーにブートストラップの見積もりをいくつか与えるよう促すなど、ユーザーごとの調整を検討しました。簡素化のため、また、集計の目的としては、「平均的な」ユーザーを基準にして推定するだけで十分であるため、これは行わないことにしました。
私たちの初期の未校正モデルは一貫して過小評価されていました。これを修正するために、対数空間で線形回帰を近似します: h = 2.28 × m 0.923 h = 2.28 \times m^{0.923} h = 2.28 × m 0.923 。これは定数乗数に近く、傾きは 1 をわずかに下回ります。傾きを 1 (単一の乗定数 2.08 2.08 2.08 ) に制約すると、すべてのメトリックが最大 0.01 0.01 0.01 変化します。バケットごとの残差は、キャリブレーション後に体系的な傾向を示しません。
この修正後でも、人間の推定値の合計は、1.4 × 1.4\× 1.4 × 修正されたモデル推定値の合計のままです。このギャップは予期されたものです。対数空間で不偏の推定量は、その予測が線形空間で合計されると偏りを持ち、体系的に合計を過小評価します。その理由を確認するには、どちらの方向にも 2 倍ずれている可能性が等しい単純化された推定量を考えてみましょう。 2 2 2 時間を予測する場合、真の値は 1 1 1 時間または 4 4 4 時間である可能性が高く、期待値は 1 + 4 2 = 2.5 \frac{1+4}{2} = 2.5 となります。
2 1 + 4 = 2.5 — 25 % 25\% 予測より 25% 上回りました。したがって、そのような予測を合計すると、真の合計が過小評価されます。これにさらなる補正を適用するのではなく、未調整の数値を意図的に保守的に過小評価したものとして報告します。
より単純な予測子もテストしました。

o 信号のどのくらいが最終的なコード変更からのものなのか、それとも完全な Devin セッションからのものなのかを理解します。 1 つ目は、人間の推定に対して、単一のスカラーを回帰し、合計行数 (セッション内のすべての PR で合計された追加 + 削除) が変更されました。パフォーマンスは悪く、R l o g 2 R^2_{log} でした。
R log 2 は 0.27 0.27 0.27 であり、コード量がエンジニアリング作業の弱い代用であることを確認しています。次に、ユーザー メッセージやその他のセッション アクティビティを使用せず、エージェントの編集ツール呼び出しのトレースのみをコンテキストとして与えて、推定エージェントを評価しました。パフォーマンスは向上しましたが、それでも完全な推定値には及ばず、重要な信号が差分外に存在することを示唆しています。
これらの結果は、エンジニアリング作業が実際にどのように行われるかに一致します。セッションの労力は、多くの場合、調査、診断、環境セットアップ、トレードオフによる推論、またはコード以外の有用な出力の生成から生じます。これらのシグナルはセッション トレース (ユーザー メッセージ、実行されたアクション、中間観察、コードベース コンテキスト) に表示されますが、必ずしも最終的なコード変更に表示されるわけではありません。
以下の回帰結果は、行数データ (セッションのプル リクエスト全体で追加および削除された行の数) がある評価データセット内の 129 129 129 セッションに関するものです。合計 170 170 170 セッション

[切り捨てられた]

## Original Extract

Engineering leaders want to know how much value AI is actually providing. We built a system that measures the number of human-equivalent hours of Devin's work.

Estimating the Productivity of an Autonomous AI Software Engineer | Cognition Menu Close Home
Estimating the Productivity of an Autonomous AI Software Engineer
Six months ago, every CTO was concerned their team wasn't using enough tokens. That trend has reversed as token usage and AI spend have skyrocketed. Engineering leaders are now trying to figure out how to measure actual output, because not every token delivers real value. Some save engineering hours and accelerate projects; others are wasted on useless sessions and bad prompting.
For any organization of sufficient scale, it's essentially impossible to measure value over thousands of sessions & billions of tokens. We set out to automate this with AI.
Predicting real ROI is hard, for reasons we’ll go into later. So we focused on a key sub-task: estimating how many productive engineering hours each Devin session is worth.
LLMs are notoriously bad at time estimates — so we expected this to be a struggle. After carefully tuning the system, however, our model has an r l o g r_{log} r l o g ​ of 0.74 0.74 0.74 and appears to be unbiased. Individual predictions aren’t perfect, but the model is good enough to be used for estimating aggregated totals. They’re also convertible to dollar amounts using engineering salaries, getting us closer to business value.
In our system, an agent reviews each completed Devin session — first classifying whether it produced useful output, then estimating how long a human engineer would have taken to produce the same work. We validated it by asking human engineers how long they would have spent on the same tasks.
The system is now running with customers. To our knowledge, this is the first automated system measuring AI engineering productivity in production.
How did we land on productive engineering hours as our metric? The first question we needed to answer was what to estimate. Ideally, we'd measure dollar impact directly, such as revenue attributable to features shipped or costs avoided by bugs fixed. In practice, this is still an unsolved problem in our field. It’s incredibly hard for an engineer to know how many dollars of business value they created through the PRs shipped last week.
On the other end of the spectrum, we could measure raw activity: lines of code, commits, PRs, tokens consumed. These are easy to collect but don't correspond to effort. A mechanical refactor can touch thousands of lines in an afternoon; a two-line bug fix can represent hours of investigation. Many valuable tasks — triaging bugs, running analytics queries, reviewing code — produce no code at all.
The middle ground we decided to measure is human engineering hours: how long would a human engineer have taken to produce the same output? Hours are already how organizations value engineering work — salaries and contractor rates are denominated in time. When leadership evaluates an investment, they think in terms of time and cost savings. Hours are standardized across organizations, independent of business context, and convertible to dollars via engineering rates.
But not all hours are equal. For example, if all PRs created by a session were closed, it likely wasn’t valuable. We wanted to measure only productive engineering hours. So, we also had to build a system to classify whether the sessions were actually productive.
We collected a ground-truth dataset by asking Devin users to review recent representative sessions, and estimate how long each completed session would have taken without Devin. Our dataset consists of 258 258 258 sessions from 126 126 126 users across a diverse set of enterprise customers. We collected the data via live interviews and a survey.
Every Devin session has a full execution trace: the user's request, every action taken, the resulting code, codebase context. This gives us a record of production engineering work at a level of detail that is difficult to obtain from surveys, aggregate activity metrics, or open-source benchmarks alone.
In the charts below we analyze our dataset. Our dataset consists of a distribution representing real enterprise workloads, spanning a variety of languages, frameworks, session types, and hour estimates.
As we reviewed our dataset, we realized that not all sessions correspond to useful work, and that we would need to build a classifier to filter out unproductive sessions.
For sessions with a PR, this is relatively straightforward: if any PR from the session is merged, we include the estimate; if not, we discard it. This is slightly lossy; sessions with all closed PRs can still have delivered productive work, but we wanted to err conservative.
Sessions without a PR are more complicated. We built a classifier to filter out unproductive sessions, which removes around 1 − 20 % 1-20\% 1 − 20% of sessions, depending on the customer. Many non-PR sessions are genuinely productive — e.g., finding unused dependencies, scanning for security vulnerabilities, reviewing a pull request, running analytics queries, triaging a bug — and we retain those. We discard sessions where the agent lacked access to carry out the task, sessions where the agent asked for clarification and the user never replied, and other scenarios where Devin was unable to meaningfully advance the task.
After we identified which sessions were productive, we needed to actually estimate their equivalent engineering hours. To do this, we built an estimator agent with two key components: context and prompts.
On the context side, we included as much information about the session as possible — the user's messages, the PR produced (if applicable), the full agent trace (viewing logs, tracing through code, fixing lint, etc.), and additional codebase context from DeepWiki .
On the prompt side, we set aside 25 25 25 sessions as a development set for iteration. By manually triaging agent runs on these sessions and reasoning critically about failure modes, we arrived at the following design principles: credit only the work Devin actually saved and compare Devin against a conservative human reference. Concretely, we:
Reason about the human's path instead of the agent's . The agent's own trajectory is sometimes a poor proxy for human effort. Agents take detours, recover from environment and setup failures, and produce artifacts like summary reports that a solo engineer wouldn't. The estimator reasons about the path a human would have taken to the same output, discounting things like retries, environment setup, and artifacts the agent produced along the way.
Credit only the work the user did not specify. The same output can represent very different amounts of effort depending on how much the user already figured out before asking Devin. The estimator looks at what the user actually said to determine how much of the problem Devin had to solve on its own. For example, if a user comes with a bug report and no proposed fix, we include the time to triage the bug, whereas if the user comes with an implementation plan, we only count the implementation time.
Account for codebase familiarity. The same task can take minutes or a day depending on how well the engineer knows the code. Users told us this is one of Devin’s clearest advantages: tasks in unfamiliar or legacy codebases that would cost a developer a day of ramp-up are often delivered quickly. The estimator infers the relevant level of familiarity from what the session reveals. When the user asks how parts of the system work, it includes the exploration time that orienting in the code would have required. When there's no signal either way, it assumes typical familiarity: an engineer who knows the high-level architecture but hasn't memorized every function.
Assume relevant expertise . A recurring theme in our interviews was that the agent lets people do things they could not have done on their own — a backend engineer delivering work that would previously have required frontend and data-science colleagues. Crediting that cross-disciplinary reach would inflate estimates, so we conservatively assume the reference engineer already has the expertise the task demands. This understates the effort in many cases where a human would first have had to learn an unfamiliar language or framework.
On the held-out evaluation set of 233 233 233 sessions, our estimator has an r l o g r_{log} r l o g ​ of 0.74 0.74 0.74 and r l o g 2 r_{log}^2 r l o g 2 ​ of 0.54 0.54 0.54 . The correlation is highly statistically significant ( F ( 1 , 231 ) = 279.9 F(1,231) = 279.9 F ( 1 , 231 ) = 279.9 , p < 10 − 5 p < 10^{-5} p < 1 0 − 5 ). Around 50 % 50\% 50% of sessions fall within a factor of 2 2 2 of the true estimate. Individual estimates are noisy — 2 2 2 – 3 3 3 × errors in either direction are common — but because errors are roughly unbiased and independent, they cancel as session count grows and the aggregate converges toward the human-reported total.
A lot of the noise comes from variance between users, both in how they estimate and in genuine differences in speed. Roughly half the residual disagreement lies between users rather than within a user's own sessions ( I C C = 0.58 ICC =0.58 I C C = 0.58 ). We considered per-user calibration, for example, prompting users in-product to give a few estimates for bootstrapping. We decided against it for simplicity and since, for our purpose of aggregation, estimating relative to an "average" user is sufficient.
Our initial, uncalibrated model consistently underestimated. To correct this, we fit a linear regression in log-space: h = 2.28 × m 0.923 h = 2.28 \times m^{0.923} h = 2.28 × m 0.923 . This is close to a constant multiplier, with a slope slightly below one. Constraining the slope to one (a single multiplicative constant of 2.08 2.08 2.08 ) changes every metric by at most 0.01 0.01 0.01 . Residuals by bucket show no systematic trend after calibration.
Even after this correction, the total of the human estimates remains 1.4 × 1.4\times 1.4 × the total of the corrected model estimates. This gap is expected: an estimator that is unbiased in log-space becomes biased once its predictions are summed in linear space, systematically underestimating the total. To see why, consider a simplified estimator that is equally likely to be off by a factor of two in either direction. When it predicts 2 2 2 hours, the true value is equally likely to be 1 1 1 hour or 4 4 4 hours, giving an expected value of 1 + 4 2 = 2.5 \frac{1+4}{2} = 2.5
2 1 + 4 ​ = 2.5 — 25 % 25\% 25% above the prediction. Summing such predictions therefore underestimates the true total. Rather than apply a further correction for this, we report the unadjusted figure as a deliberately conservative underestimate.
We also tested simpler predictors to understand how much of the signal comes from the final code change versus the full Devin session. The first regresses a single scalar, the total lines changed (additions + deletions summed across all PRs in the session), against our human estimates. It performed poorly, with an R l o g 2 R^2_{log}
R l o g 2 ​ of 0.27 0.27 0.27 , confirming that code volume is a weak proxy for engineering effort. We then evaluated an estimator agent given only the trace of the agent's edit tool calls as context, with no user messages or other session activity. It performed better, but still lagged the full estimator, suggesting that important signal lives outside the diff.
These results match how engineering work actually happens. The effort in a session often comes from investigation, diagnosis, environment setup, reasoning through tradeoffs, or producing useful non-code outputs. Those signals are visible in the session trace (user messages, actions taken, intermediate observations, and codebase context) but not always in the final code change.
The regression results below are for the 129 129 129 sessions in our evaluation dataset for which we have line-count data (the number of lines added and deleted across the session's pull requests). A total of 170 170 170 sessio

[truncated]
