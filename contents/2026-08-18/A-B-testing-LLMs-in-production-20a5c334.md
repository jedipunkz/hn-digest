---
source: "https://www.together.ai/blog/a-b-test-models-in-production"
hn_url: "https://news.ycombinator.com/item?id=49341617"
title: "A/B testing LLMs in production"
article_title: "A/B test models in production"
image: "https://cdn.prod.website-files.com/69654e88dce9154b5f12070c/6a83d4464bad620dc1d00390_OG-image_20260817_a-b-testing.png"
author: "zagwdt"
captured_at: "2026-08-18T06:25:47Z"
capture_tool: "hn-digest"
hn_id: 49341617
score: 2
comments: 0
posted_at: "2026-08-18T05:28:23Z"
tags:
  - hacker-news
  - translated
---

# A/B testing LLMs in production

- HN: [49341617](https://news.ycombinator.com/item?id=49341617)
- Source: [www.together.ai](https://www.together.ai/blog/a-b-test-models-in-production)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T05:28:23Z

## Translation

タイトル: 本番環境での LLM の A/B テスト
記事のタイトル: 運用中の A/B テスト モデル
説明: シャドウ トラフィックは、候補者が運用上健全であることを証明します。ユーザーがそれを気に入っているかどうかはわかりません。アプリコード内ではなく、エンドポイントで分割を実行します。

記事本文:
本番環境での A/B テスト モデル Webflow 追跡ブリッジの分析/最適化 -->
💰 シリーズ C を発表します。インテリジェンスは高価ではなく豊富であるべきです →
🤝 Together AI & Y Combinator が初の専用 YC GPU クラスターを提供するパートナーシップを発表 →
⚡ オンデマンド B200 が Together GPU クラスターで利用可能になりました →
🚀 効率的な推論のために MiniMax-M3 を提供中 →
API としての高パフォーマンス推論
SLA を備えたトークンベースの容量
上位のオープンソース モデルを探索する
大規模な信頼性の高い GPU クラスター
フロンティア規模のカスタムインフラストラクチャ
AIの開発環境を構築する
モデルの重みとデータを安全に保存
強化学習から完全制御へ
トップのオープンソース モデルを微調整する
プロダクションAIのシステム研究
Together AI の技術ドキュメント
実践的な実装ガイド
実稼働用の音声エージェントを構築する
質問に対する答えを見つける
API としての高パフォーマンス推論
SLA を備えたトークンベースの容量
上位のオープンソース モデルを探索する
大規模な信頼性の高い GPU クラスター
フロンティア規模のカスタムインフラストラクチャ
AIの開発環境を構築する
モデルの重みとデータを安全に保存
強化学習から完全制御へ
トップのオープンソース モデルを微調整する
プロダクションAIのシステム研究
Together AI の技術ドキュメント
実践的な実装ガイド
実稼働用の音声エージェントを構築する
質問に対する答えを見つける
すべてのブログ投稿 推論 2026 年 8 月 17 日公開
A/B テスト導入ガイド
ゼイン・ハサン、ザルニ・ピョ、ニキータ・スーリヤデバラ、テッド・キュイ
40 個以上のモデルが量産用に選択されました...40 個以上のモデルが量産用に選択されました...40 個以上のモデルが量産用に選択されました...
A/B テストでは、エンドポイントのライブ トラフィックを 1 つのコントロールと最大 20 のバリアントからなる固定コホートに分割し、それぞれのトラフィックの割合を分割できます。これにより、どのように変化するかを測定できます。

候補モデルは、選択した露出レベルで実際のユーザーに対して実行されます。バリアントの強化も 1 回の呼び出しで行うことができ、青と緑のロールアウトを使用して勝者を昇格させることができます。実験を削除すると、後でアンワインドするためにクライアント側またはルーティング ロジックを変更する必要がなく、トラフィックの 100% がコントロールに返されます。以下では、ライブ エンドポイントで実験を実行します。この実験では、95%/5% で作成し、80%/20% および 50%/50% に増加させてから、各段階で観察されたトラフィック シェアを削除して確認します。
実稼働環境での LLM の A/B テストの実装
遅かれ早かれ、どのチームも同じ質問に答えたいと考えています。新しいモデルは、現在のモデルと比較して実際にユーザーにとって優れているのでしょうか?ベンチマークでは優れていませんが、製品の実際の測定値が何であれ、リテンション、親指アップ率、タスク完了などの点で優れています。
シャドウ トラフィックではその質問に答えることはできません。シャドウイングにより、候補が遅延、エラー、スループットに関して動作的には正常であることがわかりますが、その応答は破棄されます。ユーザーがそれらに基づいて行動することはありません。質の高い質問をするには、一部のユーザーがモデル B を取得し、何が起こるかを比較するというエンド ユーザーへの実際の露出が必要です。
通常、チームは次の組み合わせを使用して、アプリケーション層でこれを自分たちで構築します。
クライアント コード内のユーザー ID の機能フラグまたは hash-mod-100。
クライアントが切り替える 2 つのエンドポイント (またはハードコードされた 2 つのモデル文字列)。
グループ A と B の意味を説明するどこかのスプレッドシート。
それは機能しますが、後で問題になる形で実験とインフラストラクチャが絡み合います。ルーティング ロジックはアプリケーションに同梱されており、クライアントが決定をキャッシュするにつれてコホート分割がドリフトする可能性があり、実験が「終了」した後でも、削除しても安全か誰も確信できないため、分岐コードはずっと後まで残り続けます。
Together AI プラットフォームを使用すると、次のことが可能になります。

A/B 実験ロジックをエンドポイント レベルで実行します。
A/B 実験はエンドポイントに接続し、厳密に 1 つのコントロールと 1 つ以上のバリアントを持つメンバーを宣言します。各メンバーはデプロイメントを指し、それぞれにパーセント設定があり、合計が 100 になる必要があり、トラフィック ルーティングを制御します。
エンドポイント ルーターがどのように機能するかというと、ベース トラフィックがコントロールにリクエストを送信するたびに、実験によってリクエストがアーム間で再サンプリングされ、95% がコントロールに残り、5% がバリアントに送られるように再分配されます。
メカニズムについて正確に言うと、実験ではベース トラフィック分割におけるコントロールのシェアを細分化します。ルーティングでは、まずウェイト分割を通じてリクエストを解決します。勝者が A/B 実験のコントロールである場合、リクエストは実験のアーム間でパーセントに基づいて再サンプリングされます。したがって、コントロールを分割メンバーの唯一のエントリポイントとして使用すると、絶対的なトラフィック シェアがパーセントになります。また、分割重みが 0 のコントロールでは実験に細分化するものが何も与えられず、結果として実験全体がトラフィックを受信しないことも注目に値します。
重要なことは、バリアント展開はエンドポイントのトラフィック分割に含めてはいけないことです。プラットフォームでは、バリアントがゼロの重みを保持する必要があります。コントロールのみがベース分割に存在します。実験では、バリアントにルーティングされたトラフィック全体を所有します。そのパーセンテージがトラフィックシェアです。バリアントが分割から容量加重トラフィックを引き出す可能性がある場合、測定値は静かに間違っていることになります。これについて考える 1 つの方法は、バリアントをシャドウ デプロイメントのように設定する必要があることです。つまり、作成済み、 READY 、重み 0 を設定し、実験のパーセンテージ設定をそこにルーティングします。
ここでのもう 1 つの重要な点は、A/B パーセントは合計が 100% になる真の固定トラフィック シェアであり、レプリカ数には依存しないということです。これを tra とは意図的に異なるものにしました

ffi-split 重み (レディレプリカごとであり、容量に従います)。実験は測定器です。測定中は分割が一定であり、自動スケーリングによって変動しないようにする必要があります。
tg ベータ エンドポイント ab my-org/candidate-model --control $CONTROL_DEPLOYMENT --percent 5
表面上は同じエンドポイント名、API、キーが存続するため、クライアントはこの実験に気づきません。バックエンドでは、リクエストの 5% がバリアント候補によって応答されるようになります。
内部: ランピング、測定、終了
ランピングはメンバー セットを再送信しています
個別の「ランプ」API はありません。更新により完全なメンバー リストが置き換えられます。これにより、メンタル モデルがシンプルに保たれ (実験は常にメンバーの発言どおりになります)、すべてのランプが明示的にレビュー可能な変更になります。
# 2 週目: 候補者は 5% で良さそうです —> 20% に進みます
client.beta.endpoints.ab_experiments.update(
id=実験ID、
エンドポイントID=エンドポイントID、
update_mask="メンバー",
etag=experiment.etag, # チームメイトの同時ランプは上書きされずに拒否される
メンバー=[
{"deployment_id": control_dep、"percent": 80、"role": "AB_EXPERIMENT_MEMBER_ROLE_CONTROL"},
{"deployment_id":variant_dep、"percent": 20、"role": "AB_EXPERIMENT_MEMBER_ROLE_VARIANT"},
]、
)
あなたがアップデートを作成している間にチームメイトが実験を開始した場合、あなたのアップデートは彼らのアップデートを黙って上書きするのではなく拒否されるため、アップデートは etag によって保護されています。
この API 設計では、グループ B にルーティングするトラフィックの量について一般的なエクスポージャーの選択を行う必要があります。
最大 20 個のバリアント メンバーを使用して、多方向テストを実行することもできます。たとえば、フル精度のエンドポイントと他の 3 つの量子化バリアント (V1、V2、V3) を試したいとします。これは、パーセントの合計が 100% であり、コントロールが 1 つだけ存在する限り、期待どおりに機能します。
自分

保証: プラットフォームのメトリクスを製品のメトリクスと結合する
すべてのリクエストは特定のデプロイメントによって処理され、すべてのプラットフォーム メトリクスはデプロイメントごとに利用可能です。そのため、比較のインフラ側 (レイテンシー、エラー、コホートごとのスループット) はプロジェクトではなくフィルターです。製品側はあなたのものです。評価、再試行、タスクの完了などの品質シグナルとともに、どのデプロイメントが各応答を提供したかを記録し (応答メタデータにあります)、結合キーは単なるデプロイメント ID です。プラットフォームは意図的に品質指標を推測しません。これにより、帰属が簡単になり、分析が判断できるようになります。
実験の結果、バリアントが勝つことが示されたとします。終了には 2 段階のプロセスがあります。
ロールアウトを通じてプロモーションします。コントロール デプロイメントからバリアント デプロイメントへの Blue-Green ロールアウトを実行します。ヘルス ゲート、伝播待機、ロールバックの安全性がすべて適用されます。
実験を削除します。すべての実験ルーティングが消え、トラフィックの 100% が、ロールアウト後に勝者を示すエンドポイントの基本トラフィック分割に従います。
バリアントが負けた場合は、実験を削除するだけで、トラフィックは完全にコントロールに戻ります。バックエンドでは、バリアントのデプロイメントがゼロにスケールされるか、削除されます。
tg beta endpoints rm abx_abc123 # 実験を削除します。 100%は基本分割に戻ります
tg beta endpoints rm dep_variant123 # またはバリアントを削除します — 自動巻き戻し実験 + 分割
エッジケース
変異体は実験の途中で劣化します。爆発範囲はどれくらいですか？
そのコホートのみ。デプロイメントは個別に監視され、個別に自動スケーリングされます。つまり、問題が発生したバリアントによって制御が低下することはありません。これを修正するには、問題のあるバリアントを含まないメンバー セットを再送信すると、そのユーザーは一定の伝播時間内に制御状態に戻ります。これも始める理由です

5% にするのが良いでしょう。
観察されたシェアは実際に設定されたパーセントと一致しますか?
意味のあるボリュームを超えています、はい！実際の例については、以下の実験をご覧ください。ウィンドウが小さい場合、サンプリング ノイズが発生することが予想されます。1,000 リクエストの 5% のシェアは小さなサンプルです。観察されたシェアが大幅にずれており、ずれたままである場合は、上記の設定ルールを確認してください。
A/B テストとロールアウトは同じエンドポイントで実行できますか?
はい。構成順序は次のように定義されます。ルーティングは最初にベース分割を解決し、次に A/B 実験 (コントロールの共有を細分化)、次にロールアウト (ソース デプロイメントとそのロールアウト ターゲットの間の再サンプリング) を解決します。プラットフォームは引き続きエンドポイントごとに 1 つのアクティブなロールアウトを強制しますが、ステージは重複する場合に構成するように設計されています。
割り当てでは、リクエストのサンプリング キー (たとえば、リクエスト本文のトップレベルの prompt_cache_key またはユーザー フィールド) を使用するため、同じキー ルートを運ぶリクエストは一貫して行われ、特定のユーザーはセッション全体で 1 つのテスト アームにとどまることができます。キーのないリクエストは、リクエストごとに実質的にランダムに割り当てられます (以下の実験測定ではキーのないトラフィックを使用しました。これが、観察されたシェアがパーセントと非常によく一致する理由です)。ユーザーごとの一貫性が研究にとって重要な場合、特に複数ターンの品質比較がある場合は、安定したユーザー フィールドを送信します。
分割時の自動スケーリングはどうなりますか?
各メンバーの展開は、トラフィックの独自のシェアに基づいてサイズ設定された独自のポリシーに基づいて拡張されます。境界 1 ～ 2 の 5% のバリアントと境界 2 ～ 8 の 95% のコントロールは完全に正常な形状です。各コホートのレプリカを個別に監視する必要があります。これは、以下のグラフに示されているものです。
1 つの A/B テストを開始から終了まで表示
ライブ エンドポイントに対してライフサイクル全体を実行し、95/5 で作成し、80/20 に増加し、50/50 に増加してから削除しました。私たちはそうします

これにより、安定した 3 RPS ストリームが維持され、すべてのリクエストはそのリクエストを処理するデプロイメントに起因します。以下のグラフは、バリアントで発生したトラフィックを示しており、緑色の点はバリアントのトラフィック シェアを示しています。
各段階での設定済みトラフィック シェアと観測されたトラフィック シェアを次に示します。
この実行で注目すべき 3 つの詳細:
各ランプは実際には 1 回の呼び出しであり、現在の etag を使用して完全なメンバー セットを再送信できます。 etag は 2 つのランプを 1 → 2 → 3 と進みます。古い etag は、チームメイトの変更を黙って上書きするのではなく、拒否されるでしょう。
伝播は速いですが、瞬時ではありません。各更新後、測定する前に約 75 秒待機しました。ルーティング層は、トラフィック分割の変更と同じ 30 ～ 60 秒のタイムスケールで実験の変更を取得します。
削除: 実験を削除した後、360 件のリクエストを連続して送信し、それらはすべてコントロールに到達しました。巻き戻すために残った削除以外に、残ったコホート ロジックはありません。
エンドポイントの [トラフィック テスト] タブにある、コンソールに表示される実験を次に示します (A/B テストとシャドウ テストがページを共有します)。
トラフィックを処理する制御デプロイメントを備えたエンドポイントと、候補デプロイメント (トラフィック分割ではなく作成済み、 READY ) が必要です。次に:
95/5 で実験を作成します。

[切り捨てられた]

## Original Extract

Shadow traffic proves a candidate is operationally sound. It can't tell you if users like it better. Run the split at the endpoint instead of in your app code.

A/B test models in production Webflow Analyze/Optimize tracking bridge -->
💰 Announcing our Series C. Intelligence should be abundant, not expensive →
🤝 Together AI & Y Combinator announce partnership to deliver the first dedicated YC GPU cluster →
⚡ On-demand B200s now available on Together GPU Clusters →
🚀 Now serving MiniMax-M3 for efficient inference →
High-performance inference as APIs
Token-based capacity with SLAs
Explore the top open-source models
Reliable GPU clusters at scale
Custom infrastructure at frontier scale
Build development environments for AI
Store model weights & data securely
From reinforcement learning to full control
Fine-tune top open-source models
Systems research for production AI
Technical docs for Together AI
Practical implementation guides
Build voice agents for production
Find answers to your questions
High-performance inference as APIs
Token-based capacity with SLAs
Explore the top open-source models
Reliable GPU clusters at scale
Custom infrastructure at frontier scale
Build development environments for AI
Store model weights & data securely
From reinforcement learning to full control
Fine-tune top open-source models
Systems research for production AI
Technical docs for Together AI
Practical implementation guides
Build voice agents for production
Find answers to your questions
All blog posts Inference Published 8/17/2026
A guide to implementing A/B testing
Zain Hasan, Zarni Phyo, Nikitha Suryadevara, Ted Cui
40+ Models Chosen for Production...40+ Models Chosen for Production...40+ Models Chosen for Production...
A/B experiments allow you to split an endpoint's live traffic into fixed cohorts of one control and up to 20 variants, each with a percentage split of the traffic. This allows you to measure how a candidate model performs with real users at an exposure level you choose. Ramping up a variant can also be done with one call where you can promote the winner using a blue-green rollout. Deleting the experiment returns 100% of traffic to the control without the need to make any client-side or routing logic changes to unwind afterwards. Below we'll run an experiment on a live endpoint where we create at 95%/5%, ramp to 80%/20% and 50%/50%, then delete and check the observed traffic shares at every stage.
Implementing A/B testing for LLMs in production
Sooner or later every team wants to answer the same question: is the new model actually better for our users compared to the current model? Not better on a benchmark but rather better on retention, thumbs-up rate, task completion, whatever your product actually measures.
Shadow traffic can't answer that question. Shadowing tells you the candidate is operationally sound with respect to latency, errors, throughput, but its responses are discarded; no user ever acts on them. Quality questions need real exposure to end users where a cut of your users get model B, and you compare what happens.
Typically teams build this themselves in the application layer using some combination of:
A feature flag or a hash-mod-100 on user ID in the client code.
Two endpoints (or two hardcoded model strings) the client switches between.
A spreadsheet somewhere explaining what group A vs B means.
It works, but it entangles your experiment with your infrastructure in ways that hurt later: the routing logic ships with your application, the cohort split can drift as clients cache decisions, and even after the experiment "ends" the branching code lives on long afterward because nobody's sure it's safe to remove.
The Together AI platform allows you to run A/B experiment logic at the endpoint level.
An A/B experiment attaches to an endpoint and declares members with exactly one control and one or more variants , each pointing at a deployment, each with a percent setting, that must sum to 100, controlling traffic routing.
How the endpoint router works is that whenever the base traffic sends a request to the control the experiment re-samples it among the arms and redistributes such that 95% stays on the control, 5% goes to the variant.
To be precise about the mechanism: the experiment subdivides the control's share of the base traffic split. Routing first resolves a request through the weight split; when the winner is the control of an A/B experiment, the request is re-sampled among the experiment's arms by their percents. With the control as the only entrypoint in the split member percents therefore are absolute traffic shares. Also worth noting is that a control whose split weight is zero gives the experiment nothing to subdivide and as a result the whole experiment receives no traffic.
Importantly variant deployments must not be in the endpoint's traffic split, the platform requires variants to carry zero weight; only the control lives in the base split. The experiment will own traffic routed to the variant entirely; its percentage is its traffic share. If a variant could also draw capacity-weighted traffic from the split, your measurements would be quietly wrong. One way to think about it is that you should set up the variant like a shadow deployment: created, READY , weight zero and then let the experiment percentage setting route to it.
Another important point here is that A/B percents are true fixed traffic shares summing to 100% and are independent of replica counts. We made this deliberately different from traffic-split weights (which are per-ready-replica and follow capacity). An experiment is a measurement instrument ; you want the split to be constant while you measure and not drift with autoscaling.
tg beta endpoints ab my-org/candidate-model --control $CONTROL_DEPLOYMENT --percent 5
Your clients won’t notice this experiment because on the surface the same endpoint name, API and keys persist. On the backend 5% of requests will now be answered by the variant candidate.
Under the hood: ramping, measuring, ending
Ramping is resending the member set
There's no separate "ramp" API, an update will replace the full member list , which keeps the mental model simple (the experiment is always exactly what its members say) and makes every ramp an explicit reviewable change:
# Week 2: candidate looks good at 5% —> go to 20%
client.beta.endpoints.ab_experiments.update(
id=experiment_id,
endpoint_id=endpoint_id,
update_mask="members",
etag=experiment.etag, # a teammate's concurrent ramp gets rejected, not overwritten
members=[
{"deployment_id": control_dep, "percent": 80, "role": "AB_EXPERIMENT_MEMBER_ROLE_CONTROL"},
{"deployment_id": variant_dep, "percent": 20, "role": "AB_EXPERIMENT_MEMBER_ROLE_VARIANT"},
],
)
Updates are guarded by an etag because if a teammate ramped the experiment while you were composing your update, yours will be rejected instead of silently overwriting theirs.
With this API design you still need to make the common exposure choice of how much traffic to route to group B:
With up to 20 variant members you can also run multi-way tests, lets say for example you want to try out a full-precision endpoint along with three other quantized variants (V1, V2, V3). This will work as expected as long as the percents still sum to 100% and there's exactly one control.
Measuring: join platform metrics with your product metrics
Every request is served by a specific deployment, and every platform metric is available per deployment — so the infra side of the comparison (latency, errors, throughput per cohort) is a filter, not a project. The product side is yours: log which deployment served each response (it's in the response metadata) alongside your quality signals — ratings, retries, task completions — and the join key is just the deployment ID. The platform deliberately doesn't guess at your quality metrics; it makes the attribution trivial so your analytics can do the judging.
Suppose the experiment shows the variant wins. Ending it is a two-step process:
Promote via rollout. Run a blue-green rollout from the control deployment to the variant deployment. Health gates, propagation waits, and rollback safety all apply.
Delete the experiment. All experiment routing disappears and 100% of the traffic then follows the endpoint's base traffic split which, post-rollout, points to your winner.
And if the variant loses, you can just delete the experiment and traffic will return entirely to the control. On the backend the variant deployment scales to zero or gets deleted.
tg beta endpoints rm abx_abc123 # delete the experiment; 100% returns to the base split
tg beta endpoints rm dep_variant123 # or delete the variant — auto-unwinds experiment + split
Edge cases
The variant degrades mid-experiment. What's the blast radius?
Only its cohort. Deployments are independently monitored and independently autoscaled, which means that a struggling variant won’t drag the control down. To fix this you can resend the member set without the sick variant and its users will be back on control within a certain propagation time. This is also why starting at 5% is a good idea.
Do the observed shares actually match the configured percents?
Over meaningful volume, yes! For an example of this in action check out the experiment below. Over small windows you can expect sampling noise: a 5% share of 1,000 requests is a small sample. If your observed share is off by a lot and stays off, check the setup rule above.
Can an A/B experiment and a rollout run on the same endpoint?
Yes, and the composition order is defined as: routing resolves the base split first, then A/B experiments (subdividing the control's share), then rollouts (re-sampling between a source deployment and its rollout target). The platform still enforces one active rollout per endpoint but the stages are designed to compose if they overlap.
Assignment uses the request's sampling key , for example a top-level prompt_cache_key or user field in the request body, so requests carrying the same key route consistently and a given user can stay in one test arm across a session. Requests without a key are assigned effectively at random per request (our experiment measurements below used key-less traffic, which is why observed shares match the percents so closely). If per-user consistency matters to your study, especially if you have multi-turn quality comparisons, send a stable user field.
What happens to autoscaling under a split?
Each member deployment scales on its own policy, sized by its own share of traffic. A variant at 5% with bounds 1-2 and a control at 95% with bounds 2-8 is a perfectly normal shape. You should watch each cohort's replicas independently, this is what we capture in the chart below.
Showing one A/B experiment, start to finish
We ran the full lifecycle against a live endpoint where we create at 95/5, ramp to 80/20, ramp to 50/50 then delete. We do this while maintaining a steady 3 RPS stream, with every request attributed to the deployment that served it. The graph below shows traffic seen at the variant with green dots capturing variant traffic share:
Here are the configured vs observed traffic shares at every stage:
Three details from the run worth calling out:
Each ramp really was just one call and you can resend the full member set with the current etag . The etag advances 1 → 2 → 3 across the two ramps; a stale etag would have been rejected rather than it silently overwriting a teammate's change.
Propagation is fast but not instant. We waited ~75 seconds after each update before measuring; the routing layer picks up experiment changes on the same 30–60s timescale as traffic-split changes.
Delete: After removing the experiment we sent 360 consecutive requests and they all landed on the control. There is no left over cohort logic besides the delete left over to unwind.
Here's the experiment as the console shows it, on the endpoint's Traffic tests tab (A/B tests and shadow tests share the page):
You need an endpoint with a control deployment serving traffic, plus a candidate deployment (created, READY , not in the traffic split). Then:
Create the experiment at 95/5 .

[truncated]
