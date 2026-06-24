---
source: "https://undes.app/blog/cheaper-ai-code-generation-engineering-cost"
hn_url: "https://news.ycombinator.com/item?id=48663557"
title: "AI coding agents need evidence-first review, not just cheaper routing"
article_title: "Cheaper AI Code Generation and Engineering Cost | Undes"
author: "CalmAngler"
captured_at: "2026-06-24T18:33:58Z"
capture_tool: "hn-digest"
hn_id: 48663557
score: 2
comments: 1
posted_at: "2026-06-24T18:06:39Z"
tags:
  - hacker-news
  - translated
---

# AI coding agents need evidence-first review, not just cheaper routing

- HN: [48663557](https://news.ycombinator.com/item?id=48663557)
- Source: [undes.app](https://undes.app/blog/cheaper-ai-code-generation-engineering-cost)
- Score: 2
- Comments: 1
- Posted: 2026-06-24T18:06:39Z

## Translation

タイトル: AI コーディング エージェントには、単に安価なルーティングだけでなく、証拠優先のレビューが必要です
記事のタイトル: AI コードの生成とエンジニアリング コストの削減 |ウンデス
説明: モデルのルーティング、取得、複数モデルの検討、自動チェック、および AI 生成コードの検証にかかる全コストの定量的な比較。

記事本文:
製品
例
ドキュメント
ブログ
価格設定
サインイン
始めましょう
製品
例
ドキュメント
ブログ
価格設定
サインイン
始めましょう
← アンデスのブログ
AI エンジニアリング経済学
Why Cheaper AI Code Generation Does Not Necessarily Reduce Engineering Cost
モデルのルーティングにより推論コストを削減できます。それで自動的にレビューが減るわけではありませんが、
手戻り、またはエラー回避のリスク。 A quantitative comparison of the control layers around AI code.
In many AI-assisted workflows, code generation is no longer the only bottleneck.
Assistants read repositories, edit files, run commands, and write tests.エージェント
systems plan, call tools, retrieve more context, and assemble an answer over several
ステップまたは複数のモデル。
What was actually checked, what did the model merely assume, and how much of this
マージ前に結果を信頼できますか?
妥当なコードの作成が安価になりました。その基盤を確認していない
必然的に続く。 Comparing AI tools only by token price, generation speed, or
agent count misses the engineering decision that matters: the path from a request to
正当なマージ決定。
この記事では次の 3 つの質問をします。
Does AI reduce total decision cost once calls, review, rework, and escaped-error risk are counted?
Which part of that cost is targeted by routing, retrieval, multi-model deliberation, and automated checks?
What should a verification layer produce, and how can its value be falsified rather than merely claimed?
生産性に関する証拠はまちまちです。 METR は、以下のランダム化比較試験を実施しました。
16 experienced open-source developers performing 246 real tasks in mature
repositories they knew well, using early-2025 tooling. AI の導入により、タスクに必要な作業が軽減されました
平均して 19% 長くなります [1] 。
In February 2026, METR reported that newer data probably shows a larger uplift,
しかし、信号を明示的に un と呼びました。

信頼性のある。返品の生の見積もり
開発者は完了時間の変化が -18% だったことに自信を持っています。
[-38%, +9%] の間隔;新しく採用された開発者にとっては、
-4% で [-15%, +9%] 、負の値は高速化を意味します。
どちらの間隔にもゼロ効果が含まれます [2] 。
正直な結論は、「AI は常に開発者のスピードを向上させる」でも「AI は常に開発者を高速化する」でもありません。
彼らの速度を低下させます。」生産性はツールの成熟度、リポジトリの習熟度、
タスクの形状、コンテキストの取得、結果を確認するコスト。
2025 年の DORA レポートは、約 5,000 件の異なる観察的見解を提供します。
テクノロジー専門家: 90% が仕事で AI を使用しています。
80% が生産性の向上を感じていますが、30% は生産性の向上を感じています
AI が生成したコードはほとんど、またはまったく信頼されていません。 AI の導入は次のこととプラスの関連性があります
配送スループットと製品パフォーマンスに悪影響を及ぼします。
配送の安定性 [9] 。これは因果関係のある推定ではありません。それは
システム仮説と一致: より高速なローカル生成が増加する可能性がある
テストと配信の制御が変更量に合わせて拡張できない場合、下流の負荷がかかります。
7 つの Google 調査を総合した結果、外部の
開発者は GenAI の出力品質をわずかしか信頼していないか、まったく信頼していません。認識される厳しさ
レビューとテストの実施、および AI が使用される場所に対する開発者の制御が肯定的に評価されました。
信頼に関連付けられています [7] 。
レビュー自体は欠陥を見つけるだけではありません。 Bacchelli と Bird の 200 件の研究では
Microsoft レビュー スレッドと 570 件のコメント、コードの改善が考慮されました
コメントと欠陥の 29%、14%。著者たち
状況と変化の理解がレビューと記録の中心となることを特定する
知識の伝達はそれ自体の成果として行われます [3] 。
レビュー負荷モデルの例
チームが週に 20 件の PR を処理し、平均レビューに 30 分かかると仮定します。
20 PR × 0.5 時間 = 10 レビュー

勤務時間/週
AI によってスループットが 2 倍になり、PR ごとのレビューコストが固定された場合:
40 PR × 0.5 時間 = 20 レビュー担当時間/週
AI 支援 PR の範囲が広がり、レビュー時間が 25% 増加した場合:
40 PR × 0.625 時間 = 25 レビュー担当時間/週
シナリオ PR/週レビュー/PR レビュー負荷
AI前 20 30分 10時間
2× スループット 40 30 分 20 時間
2 倍のスループット + より広い PR 40 37.5 分 25 時間
これは感応度モデルであり、市場統計ではありません。メカニズムを示します。
より高速な世代では、作業を削除するのではなく、書き込みからチェックに移す可能性があります。
2. エンジニアリング上の意思決定にかかる総コスト
トークンの請求額は総コストではありません。 1 つの決定にかかる予想コストを定義します。
C_total = C_model + C_tools + R_hour × (T_review + T_rework) + P_escape × L_escape
C_model : モデル呼び出し;
C_tools : CI、サンドボックス、取得、その他のコンピューティング。
R_hour : 1 エンジニアリング時間の内部コスト。
T_review : 適用/レビュー/拒否の決定を行うまでの時間。
T_rework : マージ前に見つかった問題を修正するのにかかる予想時間。
P_escape : 重要なエラーがレビューに合格する確率。
L_escape : このようなエスケープによる予想損失。
ベースラインの例を示します: C_model = $5 、レビューには 60 分かかります。
そして R_hour = $80 。ツールを設定し、作業をやり直し、リスクを一時的に脇に置きます。
C_total = $5 + $80 = $85
純粋なモデル請求最適化の上限
モデル呼び出しが分数 f = C_model / C_total の場合、最適化
ワークロード、品質、レビュー、やり直し、リスクを固定したまま、モデルの請求のみを行う
C_total を最大 f だけ下げます。参照番号:
f = 5 / 85 = 5.9%
これはルーティングの全体的な効果の上限ではありません。低価格モデルは値上がりする可能性がある
再試行、 T_rework 、および P_escape ;優れたルーターでは切断される可能性があります
遅延と失敗した通話。これは会計上の観察です。モデル請求書が
全体のごく一部であり、その行を最適化するだけではレビューを解決できません。

縛られた
ボトルネック。
レビューを 60 分から 40 分に短縮すると、異なるスケールの変化が生じます。
C_total = $5 + $80 × (40/60) = $58.33
節約 = (85 ドル - 58.33 ドル) / 85 ドル = 31.4%
機種変更レビュー C_total Saving
ベースライン $5.00 $80.00 $85.00 —
モデルコールが半減 $2.50 $80.00 $82.50 2.9%
レビュー 60→40分 $5.00 $53.33 $58.33 31.4%
両方 $2.50 $53.33 $55.83 34.3%
人間の監視がほとんどない自律的なエージェント ループでは、f は次のようになります。
大規模なルーティングが主な経済手段となる可能性があります。制約を受けるワークフローでは、
人間によるレビューにはコストがかかり、f は低くなります。関連する質問はどの用語ですか
実際には総コストの大部分を占めます。
3. 異なるシステムがコストの異なる部分を制御する
現代の AI システムは、エージェント、オーケストレーション、検索、裁判官、
そして合成。似た形状は同じ仕事を意味するわけではありません。
ルーティング: Kilo Gateway と RouteLLM
Kilo は、OpenAI 互換エンドポイント、多くのモデルへのアクセス、BYOK、使用法を公開します
追跡、支出制限、組織管理 [11] 。
ByteByteGo は、既知のモード (計画、コーディング、デバッグ) でのルーティングを説明します。
ユーザーが選択した層とサーバーが更新したモデル マップ。報告されたキロ数 —
リクエストの平均コストが約 3 分の 1 低くなり、リクエストの 80 ～ 90% は必要ありません。
フロンティア モデル、10 倍を超える階層ギャップ、四半期あたり推定 87,000 ドル
ルーチン トラフィックの誤ったルーティングによる過剰支出 — ベンダーから報告されていますが、そうではありません
独自に検証済み [8] 。
理想化されたモデルは、潜在的なスケールを示します。
相対コスト = 0.15 × 1 + 0.85 × 0.10 = 0.235
相対的な減少 = 1 - 0.235 = 76.5%
RouteLLM は、トレードオフに関する主要な研究証拠を提供します: 3.66 倍のコスト削減
GPT-4/Mixtral-8×7B ペアの GPT-4 MT-Bench スコアの 95% に相当する比率。
72.7% の相対コスト削減 [12] 。そのコストモデルは短いものを使用します
シングルターンプロンプトと

品質としてのベンチマークスコア。コーディングエージェントループではありません
またはリポジトリの変更が安全であるという証拠。
エージェント RAG: 十分なコンテキスト
Google では、専用の Sufficient Context Agent を備えたマルチエージェント RAG について説明しています。それ
クエリ、取得したスニペット、ドラフトを比較し、不足している情報に名前を付けます。
別の取得パスをトリガーできます。 Google は最大 34% 高い精度を報告
事実データセットに対する標準的な RAG よりも優れています [4]。
十分なコンテキストの調査により、より広範な故障モードが明らかになります。モデルは多くの場合、次のような答えを出します。
文脈が不十分な場合に棄権するのではなく、間違って発言します。誘導された棄権
Gemini、GPT、Gemma について、回答されたケースの正しさが 2 ～ 10% 向上しました
[5] 。
これは十分なコンテキストのループをサポートしますが、測定された量の削減ではありません。
ソフトウェア開発用の T_rework または P_escape。コードベース
単なる文書コーパスではありません。実行時の動作、呼び出し元、不変条件が含まれます。
そして移住。
マルチモデルの審議: コンセンサスは証拠ではない
OpenRouter Fusion は 1 ～ 8 モデルの並列パネルを実行します。裁判官は構造化された判決を返す
コンセンサス、矛盾、部分的な報道、独自の洞察の比較、
盲点。最終モデルが答えを書きます。ドキュメントでは、
パイプラインですが、独立した有効性ベンチマークは提供しません
[10] 。
Google Research では 180 のエージェント構成を比較しました。独立したトポロジーの増幅
一元的な調整を維持しながら、エラーは最大 17.2 倍に増加
4.4倍までの増幅。マルチエージェントにより並列化可能性が向上
Finance-Agent の結果は 80.9% ですが、すべてのマルチエージェント バリアント
PlanCraft の連続結果が 39 ～ 70% 低下しました。著者らの
予測モデルは、目に見えない構成の 87% に対して最適なアーキテクチャを選択しました
[6] 。
この評価にはリポジトリのコードレビューは含まれていませんでした。より狭いエンジニアリング
仮説

それは、値がトポロジ、タスクの分解可能性、集中管理に依存することです。
ゲートと証拠の引き渡し - エージェントの数だけではありません。
SAST、DAST、CodeQL、Semgrep、単体テスト、および突然変異テストにより、反復可能なチェックが可能
制御された入力、構成、および制御の下で明示的にエンコードされたプロパティの
環境。その品質は、カバレッジ、偽陽性、偽陰性、
そして薄片状。
これらは必要ですが、モデルが関連するファイルを開いていないことを常に明らかにするわけではありません。
ファイルを作成したり、誤った仮定に基づいて結論を構築したり、実装の詳細をテストしたりする
システム不変式の代わりに。緑色のチェックは完全な意図を証明するものではありません。
これらのシステムは必ずしも直接の競合相手ではありません。ルーティングはモデル呼び出しを管理します
コスト。 Agentic RAG はコンテキストの十分性をテストします。マルチモデルの検討面
意見の相違。テストは形式化されたプロパティをチェックします。検証アーティファクトは、
これらのシグナルを、候補者がどの程度支持されているかに関する決定に結び付けます。
5. 信託債務と隠れたチェック作業
工学的な回答に一連の重要な主張が含まれているとします。
C = {c1, c2, ..., cn}
査読者はそれぞれの主張について、それが証拠によって裏付けられているかどうかを知る必要があります。
矛盾しているか、まだ仮定です。大まかな診断指標は次のとおりです。
証拠の適用範囲 = サポートされるクレーム数 / マテリアルの総クレーム数
回答に 20 の重要な主張が含まれており、12 について十分な証拠が存在する場合:
証拠範囲 = 12 / 20 = 60%
残りの 40% は必ずしも間違っているわけではありません。彼らはまだレビュアーの領域です
検査する必要がある。ツールがその領域を公開しない場合、エンジニアはまず次のことを行う必要があります。
それを発見してから初めて検証してください。それは隠れた検証作業です。
検証層の目標は、答えが完全に正しいと宣言することではありません。それは次のとおりです。
重要な主張をチェック可能な証拠に結び付ける。
関連するターゲットを公開する

検査されたことも検査されていないこともある。
仮定を裏付けられた結論から分離する。
批判と否定された仮説を保持します。
オープンプロダクションとPRのリスクが表面化。
不確実性を隠すことなく手動検索範囲を狭めます。
レビューが残っています。探索範囲が小さくなるはずです。
6. 追加の検証が報われる場合
リスクを一時的に無視すれば、ΔC の費用がかかる追加小切手は元が取れます。
少なくとも T_break_even = ΔC / R_hour を節約した場合。で
R_時間 = $80 :
P_escape を 0.1 パーセント ポイント削減 (1.0% から 0.9% に)
L_escape = $10,000 の利回り:
(0.010 - 0.009) × 10,000 ドル = 実行ごとに 10 ドルの節約が期待される
L_escape 保存/実行 100 回実行時の保存/月
$1,000 $1 $100
10,000ドル 10ドル 1,000ドル
100,000ドル 100ドル 10,000ドル
1,000,000ドル 1,000ドル 100,000ドル
これは予想損失モデルであり、測定された製品の結果や文字通りの結果ではありません
保険。小規模な場合には、高価な検証も経済的に合理的である可能性があります。
故障確率の低減により、大きな損失が防止されます。
7. 仮説をテストするために使用される 1 つの実装
私たちが構築および評価している実装の 1 つは、
ウンデス。複数のモデル、批評、コンセンサス、
合成はメカニズムです。テストされている製品オブジェクトはレビュー可能なアーティファクトです
それは以下を保存することを目的としています。
提案されたソリューションまたはコード候補。
関連する

[切り捨てられた]

## Original Extract

A quantitative comparison of model routing, retrieval, multi-model deliberation, automated checks, and the full cost of verifying AI-generated code.

Product
Examples
Docs
Blog
Pricing
Sign in
Get started
Product
Examples
Docs
Blog
Pricing
Sign in
Get started
← Undes Blog
AI engineering economics
Why Cheaper AI Code Generation Does Not Necessarily Reduce Engineering Cost
Model routing can reduce inference cost. That does not automatically reduce review,
rework, or escaped-error risk. A quantitative comparison of the control layers around AI code.
In many AI-assisted workflows, code generation is no longer the only bottleneck.
Assistants read repositories, edit files, run commands, and write tests. Agentic
systems plan, call tools, retrieve more context, and assemble an answer over several
steps or several models.
What was actually checked, what did the model merely assume, and how much of this
result can I rely on before merge?
Producing plausible code has become cheaper. Checking its foundations has not
necessarily followed. Comparing AI tools only by token price, generation speed, or
agent count misses the engineering decision that matters: the path from a request to
a justified merge decision.
This article asks three questions:
Does AI reduce total decision cost once calls, review, rework, and escaped-error risk are counted?
Which part of that cost is targeted by routing, retrieval, multi-model deliberation, and automated checks?
What should a verification layer produce, and how can its value be falsified rather than merely claimed?
The productivity evidence is mixed. METR ran a randomized controlled trial with
16 experienced open-source developers performing 246 real tasks in mature
repositories they knew well, using early-2025 tooling. With AI, tasks took
19% longer on average [1] .
In February 2026, METR reported that newer data probably shows a larger uplift,
but explicitly called the signal unreliable. The raw estimate for returning
developers was -18% change in completion time with a confidence
interval of [-38%, +9%] ; for newly recruited developers it was
-4% with [-15%, +9%] , where negative means speedup.
Both intervals include zero effect [2] .
The honest conclusion is neither “AI always speeds developers up” nor “AI always
slows them down.” Productivity depends on tool maturity, repository familiarity,
task shape, context acquisition, and the cost of checking the result.
The 2025 DORA report provides a different, observational view of nearly 5,000
technology professionals: 90% use AI at work, more than
80% perceive a productivity gain, but 30% have
little or no trust in AI-generated code. AI adoption is positively associated with
delivery throughput and product performance and negatively associated with
delivery stability [9] . This is not a causal estimate. It is
consistent with a systems hypothesis: faster local generation may increase
downstream load if testing and delivery controls do not scale with change volume.
A synthesis of seven Google studies found that 39% of external
developers trust GenAI output quality only slightly or not at all. Perceived rigor
of review and testing, and developer control over where AI is used, were positively
associated with trust [7] .
Review itself is not only defect-finding. In Bacchelli and Bird’s study of 200
Microsoft review threads and 570 comments, code improvements accounted for
29% of comments and defects for 14% . The authors
identify understanding the context and the change as central to review and record
knowledge transfer as an outcome in its own right [3] .
An illustrative review-load model
Assume a team handles 20 PRs per week and an average review takes 30 minutes:
20 PR × 0.5 h = 10 reviewer-hours / week
If AI doubles throughput while review cost per PR stays fixed:
40 PR × 0.5 h = 20 reviewer-hours / week
If AI-assisted PRs become wider and review time rises by 25%:
40 PR × 0.625 h = 25 reviewer-hours / week
Scenario PR/wk Review/PR Review load
Pre-AI 20 30 min 10 h
2× throughput 40 30 min 20 h
2× throughput + wider PRs 40 37.5 min 25 h
This is a sensitivity model, not a market statistic. It shows the mechanism:
faster generation may move work from writing to checking rather than remove it.
2. The total cost of an engineering decision
The token bill is not the total cost. Define the expected cost of one decision:
C_total = C_model + C_tools + R_hour × (T_review + T_rework) + P_escape × L_escape
C_model : model calls;
C_tools : CI, sandbox, retrieval, and other compute;
R_hour : internal cost of one engineering hour;
T_review : time to an apply/review/reject decision;
T_rework : expected time to fix issues found before merge;
P_escape : probability that a material error passes review;
L_escape : expected loss from such an escape.
Take an illustrative baseline: C_model = $5 , review takes 60 minutes,
and R_hour = $80 . Set tools, rework, and risk aside temporarily:
C_total = $5 + $80 = $85
The ceiling on pure model-bill optimization
If model calls are a fraction f = C_model / C_total , then optimizing
only the model bill while holding workload, quality, review, rework, and risk fixed
lowers C_total by at most f . At the reference numbers:
f = 5 / 85 = 5.9%
This is not a ceiling on routing’s total effect. A weaker cheap model may raise
retries, T_rework , and P_escape ; a good router may cut
latency and failed calls. It is an accounting observation: when the model bill is
a small part of the total, optimizing that line alone cannot solve a review-bound
bottleneck.
Cutting review from 60 to 40 minutes produces a different scale of change:
C_total = $5 + $80 × (40/60) = $58.33
Saving = ($85 - $58.33) / $85 = 31.4%
Change Model Review C_total Saving
Baseline $5.00 $80.00 $85.00 —
Model calls halved $2.50 $80.00 $82.50 2.9%
Review 60→40 min $5.00 $53.33 $58.33 31.4%
Both $2.50 $53.33 $55.83 34.3%
In autonomous agentic loops with little human oversight, f may be
large and routing can become the main economic lever. In workflows constrained by
costly human review, f is lower. The relevant question is which term
actually dominates the total cost.
3. Different systems control different parts of the cost
Modern AI systems often look similar: agents, orchestration, retrieval, a judge,
and synthesis. Similar shape does not imply the same job.
Routing: Kilo Gateway and RouteLLM
Kilo exposes an OpenAI-compatible endpoint, access to many models, BYOK, usage
tracking, spend limits, and organization controls [11] .
ByteByteGo describes routing on a known mode — planning, coding, debugging — with
user-selected tiers and a server-updated model map. The reported Kilo figures —
roughly one-third lower average request cost, 80–90% of requests not requiring
frontier models, a greater-than-10× tier gap, and an estimated $87K quarterly
overspend from misrouting routine traffic — are vendor-reported and not
independently verified [8] .
An idealized model shows the potential scale:
relative_cost = 0.15 × 1 + 0.85 × 0.10 = 0.235
relative reduction = 1 - 0.235 = 76.5%
RouteLLM provides primary research evidence for the trade-off: a 3.66× cost-saving
ratio at 95% of GPT-4’s MT-Bench score for a GPT-4/Mixtral-8×7B pair, equivalent to
72.7% relative cost reduction [12] . Its cost model uses short
single-turn prompts and benchmark score as quality. It is not a coding-agent loop
or evidence that a repository change is safe.
Agentic RAG: sufficient context
Google describes a multi-agent RAG with a dedicated Sufficient Context Agent. It
compares the query, retrieved snippets, and a draft, names missing information,
and can trigger another retrieval pass. Google reports up to 34% higher accuracy
than standard RAG on factuality datasets [4] .
The Sufficient Context research exposes a broader failure mode: models often answer
incorrectly rather than abstain when context is insufficient. Guided abstention
improved correctness among answered cases by 2–10% for Gemini, GPT, and Gemma
[5] .
This supports a sufficient-context loop, but it is not a measured reduction in
T_rework or P_escape for software development. A codebase
is not merely a document corpus; it contains runtime behavior, callers, invariants,
and migrations.
Multi-model deliberation: consensus is not proof
OpenRouter Fusion runs a parallel panel of 1–8 models. A judge returns a structured
comparison of consensus, contradictions, partial coverage, unique insights, and
blind spots; a final model writes the answer. The documentation describes the
pipeline but does not provide an independent effectiveness benchmark
[10] .
Google Research compared 180 agent configurations. Independent topology amplified
errors by up to 17.2× , while centralized coordination held
amplification to 4.4× . Multi-agent improved the parallelizable
Finance-Agent result by 80.9% , but every multi-agent variant
degraded the sequential PlanCraft result by 39–70% . The authors’
predictive model selected the optimal architecture for 87% of unseen configurations
[6] .
This evaluation did not contain repository code review. The narrower engineering
hypothesis is that value depends on topology, task decomposability, a centralized
gate, and evidence handoffs — not on agent count alone.
SAST, DAST, CodeQL, Semgrep, unit tests, and mutation tests provide repeatable checks
of explicitly encoded properties under controlled inputs, configuration, and
environment. Their quality is bounded by coverage, false positives, false negatives,
and flakiness.
They are necessary, but do not always reveal that a model never opened the relevant
file, built a conclusion on a false assumption, or tested an implementation detail
instead of a system invariant. Green checks are not proof of complete intent.
These systems are not necessarily direct competitors. Routing manages model-call
cost. Agentic RAG tests context sufficiency. Multi-model deliberation surfaces
disagreement. Tests check formalized properties. A verification artifact should
connect those signals to a decision about how far a candidate is supported.
5. Trust debt and hidden checking work
Suppose an engineering answer contains a set of material claims:
C = {c1, c2, ..., cn}
For each claim, a reviewer needs to know whether it is supported by evidence,
contradicted, or still an assumption. A rough diagnostic metric is:
evidence_coverage = supported_claims / total_material_claims
If an answer contains 20 material claims and sufficient evidence exists for 12:
evidence_coverage = 12 / 20 = 60%
The remaining 40% are not necessarily wrong. They are the area a reviewer still
needs to inspect. If a tool does not expose that area, the engineer first has to
discover it and only then verify it. That is hidden verification work.
The goal of a verification layer is not to declare an answer absolutely correct. It is to:
connect material claims to checkable evidence;
expose relevant targets that were and were not inspected;
separate assumptions from supported conclusions;
preserve critique and rejected hypotheses;
surface open production and PR risks;
narrow the manual search area without hiding uncertainty.
Review remains. The search area should become smaller.
6. When extra verification pays for itself
Ignoring risk for a moment, an extra check costing ΔC pays for itself
when it saves at least T_break_even = ΔC / R_hour . At
R_hour = $80 :
Reducing P_escape by 0.1 percentage point — from 1.0% to 0.9% — at
L_escape = $10,000 yields:
(0.010 - 0.009) × $10,000 = $10 expected saving per run
L_escape Saving/run Saving/month at 100 runs
$1,000 $1 $100
$10,000 $10 $1,000
$100,000 $100 $10,000
$1,000,000 $1,000 $100,000
This is an expected-loss model, not a measured product outcome and not literal
insurance. Expensive verification can still be economically rational when a small
reduction in failure probability protects against a large loss.
7. One implementation used to test the hypothesis
One implementation we are building and evaluating is
Undes . Multiple models, critique, consensus, and
synthesis are mechanisms. The product object being tested is a reviewable artifact
that aims to preserve:
the proposed solution or code candidate;
relevant

[truncated]
