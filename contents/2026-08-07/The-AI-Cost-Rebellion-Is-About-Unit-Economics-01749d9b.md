---
source: "https://www.voxelperfect.com/writing/the-ai-cost-rebellion-is-really-about-unit-economics"
hn_url: "https://news.ycombinator.com/item?id=49208768"
title: "The AI Cost Rebellion Is About Unit Economics"
article_title: "The AI Cost Rebellion Is Really About Unit Economics | Kostas Karolemeas"
author: "voxelperfect"
captured_at: "2026-08-07T11:38:21Z"
capture_tool: "hn-digest"
hn_id: 49208768
score: 1
comments: 0
posted_at: "2026-08-07T11:27:19Z"
tags:
  - hacker-news
  - translated
---

# The AI Cost Rebellion Is About Unit Economics

- HN: [49208768](https://news.ycombinator.com/item?id=49208768)
- Source: [www.voxelperfect.com](https://www.voxelperfect.com/writing/the-ai-cost-rebellion-is-really-about-unit-economics)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T11:27:19Z

## Translation

タイトル: AI コストの反乱はユニットエコノミクスに関するものです
記事のタイトル: AI コストの反乱は実際にはユニットエコノミクスに関するもの |コスタス・カロレメアス
説明: トークン紙幣は価値を証明しないというアレックス・カープの指摘は正しい。しかし、トークン価格設定を廃止しても、アトリビューションの問題は解決されません。企業は、ワークフローごとの総コストと実現価値を測定し、フロンティア、小規模、オープンウェイト、および人的能力を越えて作業をルーティングする必要があります。

記事本文:
AI のコスト反乱の正体はユニットエコノミクス | Kostas Karolemeas コンテンツへスキップ Kostas Karolemeas 仕事 執筆 個人的な連絡先について話す 執筆アーカイブ 2026 年 8 月 5 日 · AI FinOps
AI のコスト反乱の正体はユニットエコノミクス
アレックス・カープ氏は、トークン紙幣は価値を証明しないという指摘は正しい。しかし、トークン価格設定を廃止しても、アトリビューションの問題は解決されません。企業は、ワークフローごとの総コストと実現価値を測定し、フロンティア、小規模、オープンウェイト、および人的能力を越えて作業をルーティングする必要があります。
アレックス・カープは最も挑発的な方法で正しい質問をしました。
最先端の AI 企業がこれほど多くの価値を生み出すのであれば、なぜ成果を共有せずにトークンに課金するのでしょうか?
質問は商業的なものです。 Palantir は、制御された展開、独自のデータ、運用ワークフローを中心に構築された、エンタープライズ AI の代替ビジョンを販売しています。しかし、多くの企業顧客がすでにこの質問の、あまり劇場版ではないものを求めていたため、この質問は的中しました。
私たちは実際に何に対してお金を払っているのでしょうか？
Handelsblatt 氏は、OpenAI と Anthropic の価格設定に対する顧客の抵抗が増大しており、同時に制御と結果に対する不満も高まっていると報告しています。コスト管理に関するフォローアップにより、運用上の問題が明らかになりました。AI への支出は急速に増加していますが、何に支払っているのかを正確に把握している企業はほとんどありません。
報告に貢献した Vanessa Cann は、この移行について詳しく説明しています。パイロット段階では、消費量が小さすぎてビジネスケースを支配できませんでした。企業が有用なシステム、特に 1 つのユーザー要求で多くのモデル呼び出しをトリガーできるエージェントを拡張すると、隠れたコスト構造が可視化されました。新たな問題は、もはやパイロットが機能するかどうかではありません。特定のワークフローが繰り返し運用されることでどのような価値が生み出されるのかを指します。
これは AI の拒絶サイクルではありません。
私

これは、AI のユニット経済サイクルの始まりです。
トークンは実際のメーターであり、管理単位が間違っています
トークンは想像上のものではありません。これは、モデルが実行する作業の一部を測定し、エンジニアに消費量を理解する方法を提供し、プロバイダーに変数の使用量を請求する方法を提供します。
しかし、ビジネスオーナーにとってトークンはほとんど意味がありません。
保険請求担当者はトークンを望んでいません。彼女は請求が正しく処理されることを望んでいます。サポート リーダーは解決されたケースを望んでいます。エンジニアリングマネージャーは、安全な変更をマージしたいと考えています。コンプライアンス担当者は、監査可能な決定を伴う調査済みの例外を望んでいます。
トークンは、それらの結果の下のいくつかの層に位置します。
また、単純な比較単位としても機能しません。小規模な分類器からの 1 つのトークン、フロンティア推論モデルからの 1 つ、および長いエージェント ループ内で生成された 1 つのトークンは、経済的に同等ではありません。プロバイダーの価格リストでは、入力、出力、キャッシュされた入力、長いコンテキスト、サービス層、ツール、バッチ処理が区別されるようになりました。 OpenAI 、 Anthropic 、 DeepSeek の公開カタログを見ると、価格と製品範囲がいかに広くなっているのかがわかります。
これらのオプションは最適化に役立ちます。彼らはワークフローがその地位を獲得するかどうかには答えていない。
FinOps Foundation のトークン経済学に関する研究では、この区別が明確にされています。トークンコストは、インフラストラクチャ、データ、ネットワーキング、エンジニアリング、可観測性、評価、ガバナンス、労働力を含む幅広いスタックの 1 つのレイヤーにすぎません。また、推論と複数のツール呼び出しを伴う取得パイプラインは、より小規模なモデルへの直接リクエストよりも 1 ～ 2 桁多くのトークンを消費する可能性があることにも注意してください。
プロバイダーには生産メーターが必要です。
企業には経済台帳が必要です。
両者を混同することが反乱の根源である。
スケールはパイロットが隠したものを明らかにする
パイロットの経済学は寛大です。
ユーザーグループは小さいです。仕事

kflow は監視されています。失敗は学習として扱われます。統合とガバナンスの労働力は、多くの場合中央から資金提供されます。消費はトライアル、高額なサブスクリプション、またはまだ誰も挑戦していない予算の範囲内にある可能性があります。
本番環境ではこれらの保護が削除されます。
導入が成功するたびに、ボリュームが増加します。エージェントは、計画、取得、ツール呼び出し、再試行、検証、およびメモリを追加します。より多くのユーザーがより多くのエッジケースを作成します。自律性を高めるには、より適切な評価、監視、アクセス制御、インシデント対応、人間によるエスカレーションが必要です。インタラクションごとに安価に見えるシステムでも、完了したプロセスごとに高価になる可能性があります。
証拠はまだ出てきていますが、警告の兆候は強いです。マッキンゼーは、資格のある企業回答者 75 名を対象とした 2026 年 5 月の調査で、組織が個別のユースケースから企業導入に移行するにつれて支出がほぼ 4 倍に増加し、93 パーセントが AI 予算を超過していると報告したことを発見しました。同じ分析では、エージェントが同じタスクを実行する場合、トークンの使用には最大 30 倍の変動があることが判明した調査結果を引用しています。
これは小規模な調査であり、普遍的な市場推定ではありません。しかし、これは構造的な問題を説明しています。エージェントの需要は非線形であるのに対し、ほとんどの予算はユーザー、リクエスト、コスト間の安定した関係を前提としています。
これは現在、経営上の主要な懸念事項となっています。 State of FinOps 2026 レポートによると、FinOps チームの 98% が AI 支出を管理しており、2 年前の 31% から増加しており、主なスキルギャップとして AI コスト管理を挙げています。
マーケットは「Can we build it?」から移行しました。 to 経済的に運用できるか？
Karp の価値観は正しいが、結果の価格設定は魔法ではない
Karp の暗黙の代替案は魅力的です。ベンダーがビジネス価値を生み出すと主張する場合、トークンの代わりにその価値を課金させます。
市場の一部はすでにその方向に動きつつある。
インターホンになりました

結果によって Fin エージェントをライスします。サポートの解決、手順の引き継ぎ、または失格には 0.99 ドルの費用がかかります。販売資格の取得には 9.99 ドルかかります。 Zendesk も同様に、AI エージェントの請求単位として自動解決を使用しています。
これは、内部推論のたびに顧客に請求するよりも適切です。これにより、ベンダーはより多くのパフォーマンスリスクを負うことになり、購入者には作業に似た単位が与えられます。
また、結果の価格設定がなぜ難しいのかも明らかになります。
結果は自然なものではありません。それは定義を超えた契約です。
Intercom は、顧客が成功を確認した場合、または最後の回答後に追加のサポートを要求しなかった場合に、解決策としてカウントします。これは合理的な運用上の定義ですが、顧客の根本的な問題が解決されたこと、答えが依然として正しいこと、または後から下流コストが発生しないことを証明することと同じではありません。
カスタマーサポートの外では難易度が上がります。
AI が提案書を作成し、営業担当者が取引を成立させた場合、誰が評価されるのでしょうか?
マージされたコード変更は、1 週間後に運用インシデントを作成した場合に結果として影響しますか?
組織がキャパシティを再配置しない場合、時間の節約には真の価値があるのでしょうか?
リスクを回避したり、サイクルタイムを短縮したり、より適切な決断を下した場合、ベンダーにはどのような報酬が支払われるべきでしょうか?
どのベースラインが利益を決定しますか? 以前の人的プロセス、最良の代替モデル、または何もしないことです。
成果の価格設定により、インセンティブを調整できます。帰属、遅延効果、ゲーム、または反事実をめぐる意見の相違を排除することはできません。
トークンは、サプライヤーの質問「モデルの容量はどのくらい消費されましたか?」に答えます。
成果価格設定は、どのようなイベントが支払いをトリガーするのかという商業的な質問に答えます。
どちらも、それ自体では、「このワークフローは永続的な正味価値を生み出したのか?」という経営上の疑問には答えられません。
企業にはワークフローレベルのユニットエコノミクスが必要です
右側のユニットが間にあります

n トークンと年次 ROI スライド。
それはビジネス イベントです。つまり、事件の解決、請求の処理、変更の確認、注文の回収、調査の完了、申請の承認などです。
組織はイベントごとに、完全なコストと結果の記録を必要とします。
基本的な管理方程式はシンプルです。
ワークフローの合計コスト / 受け入れられた結果 = 成功した結果ごとのコスト
難しい作業は、「受け入れられた」を正直に定義し、分子全体を測定することです。
これが、「節約された時間」が通常、最終的な指標ではなく中間的な指標である理由です。節約された時間は、それによって目に見える何かが変化した場合にのみ経済的価値となります。つまり、完了する作業の増加、収益の迅速化、請負業者の支出の削減、雇用の回避、未処理の削減、品質の向上、従業員の能力の真の再配置などです。
同じ規律が品質にも当てはまります。レビュー時間が 2 倍になったら、より安価なモデルが安くなるわけではありません。フロンティア モデルは、重大な結果をもたらすエラーを防ぐのであれば、それほど高価ではありません。自動化に多くの監視と再作業が必要になり、プロセス全体のコストが高くなる場合、人間の作業は非効率的ではありません。
ユニットエコノミクス台帳を使用すると、これらの比較が可能になります。
フロンティア、小規模、オープン、そして人間的なポートフォリオの選択
コストの反動により、より小型でオープンウェイトのシステムを含め、モデルの代替が加速するでしょう。それは健全ですが、「より安価なモデルを使用する」ということは戦略ではありません。
すべてのワークフローは、品質、遅延、制御、リスクの要件を満たす最も安価な実行パスにルーティングされる必要があります。
より強力な推論が結果を大きく変えるような、複雑、曖昧、または結果の大きい作業のためのフロンティア モデル。
測定可能な許容基準を備えた、制限された反復的な大量のタスク向けの小規模なホスト モデル。
スケール、データ制御、カスタマイズ、または継続性によってインフラストラクチャが正当化されるオープンウェイト モデル

構造も操作負担も。
例外、論争のある決定、人間関係に敏感な作業、またはレビューや失敗のコストによって自動化の利点が失われるタスクには人間が対応します。
オープンウェイトの移動コスト。彼らはそれを消滅させません。プロバイダーのマージンは、アクセラレータの容量、プラットフォーム エンジニアリング、セキュリティ、評価、運用によって置き換えられる可能性があります。正しい比較は、100 万トークンあたりの API 価格ではなく、承認された結果あたりの総コストです。
これにより、フロンティアモデルの役割も変わります。これは、あらゆる要約、分類、エージェント ループの背後にあるデフォルト エンジンではなく、意図的に使用される希少な機能になります。
コントロール プレーンが経済記録システムになる
私は以前、AI トークンのコストには、ワークロードのルーティング、帰属、予算、承認、例外がランタイム近くに適用される運用モデルが必要であると主張しました。
さらに、同じコントロール プレーンで消費と結果を結び付ける必要があります。
すべての本番ワークフローについて、以下を記録する必要があります。
使用されるモデル、ツール、データ、
トークン、インフラストラクチャ、サードパーティのコスト、
レイテンシー、再試行、フォールバック、失敗したループ、
人間によるレビューと修正の時間、
評価または合格結果、
そして次に利用可能なモデルまたは人間の代替品。
この層がなければ、企業は価値を破壊しながらトークンを削減する可能性があります。また、生産性を向上させる一方で、レビューの労力、運用リスク、ベンダーへの依存を静かに増大させることもできます。
これにより、コスト管理が動的になります。このシステムは、日常的な作業を下方にルーティングし、困難なケースを上方にエスカレーションし、安定したコンテキストをキャッシュし、暴走ループを阻止し、予算のしきい値を強制し、一般的なベンチマークではなく企業独自のタスクを使用してプロバイダーを比較できます。
これは FinOps のレポート以上のものです。
これは、インテリジェンスの経済学のためのオペレーティング システムです。
調達が牽引する

ardハイブリッド契約
私はトークンの価格設定がなくなるとは思っていません。
API、インフラストラクチャ、実験、およびプロバイダーが顧客の最終的なビジネス結果を観察できないワークロードには引き続き役立ちます。純粋な結果価格設定は、イベントが範囲が狭く、頻度が高く、原因が特定でき、すぐに検証できる場合に最も効果的です。サポート解決がその良い例です。
したがって、ほとんどの企業契約はハイブリッドになるでしょう。
予測可能な基本需要に対する容量コミットメント、
変動使用量に対する消費価格設定、
レイテンシと可用性のためのサービス層、
成功を定義できる結果の要素、
合意されたパフォーマンスが達成されなかった場合の商用クレジットまたはリスク分担。
購入者の影響力は、計測器と移植性によってもたらされます。企業は、エクスポート可能な使用状況データ、ワークロードレベルの帰属、明確な結果の定義、監査可能な品質尺度、コスト上限、モデル代替の権利、別のプロバイダーまたはオープンウェイト展開へのテスト済みのパスを必要とします。
測定して切り替える能力が、ベンダーの競争を企業の交渉力に変えるのです。
反乱は成熟のシグナルである
企業は AI に価値がないことに気づいていません。
彼らは、養子縁組は経済的利益を伴わないものであることを発見しつつあります。

[切り捨てられた]

## Original Extract

Alex Karp is right that token bills do not prove value. But abolishing token pricing will not fix the attribution problem. Enterprises need to measure total cost and realized value per workflow, then route work across frontier, smaller, open-weight, and human capacity.

The AI Cost Rebellion Is Really About Unit Economics | Kostas Karolemeas Skip to content Kostas Karolemeas Work Writing Speaking About Personal Contact Writing archive Aug 5, 2026 · AI FinOps
The AI Cost Rebellion Is Really About Unit Economics
Alex Karp is right that token bills do not prove value. But abolishing token pricing will not fix the attribution problem. Enterprises need to measure total cost and realized value per workflow, then route work across frontier, smaller, open-weight, and human capacity.
Alex Karp asked the right question in the most provocative way.
If frontier AI companies create so much value, why do they charge for tokens instead of sharing in the outcome?
The question is commercially interested. Palantir sells an alternative vision of enterprise AI, built around controlled deployment, proprietary data, and operational workflows. But the question landed because many enterprise customers were already asking a less theatrical version of it:
What are we actually paying for?
Handelsblatt reports growing customer resistance to the pricing of OpenAI and Anthropic, alongside frustration about control and results. Its follow-up on cost management makes the operational problem explicit: AI spend is rising quickly, but few companies know precisely what they are paying for.
Vanessa Cann, who contributed to the reporting, describes the transition well . During the pilot phase, consumption was too small to dominate the business case. Once companies scaled useful systems—especially agents, where one user request can trigger many model calls—the hidden cost structure became visible. The new question is no longer whether a pilot works. It is what value a specific workflow creates when operated repeatedly.
This is not an AI rejection cycle.
It is the beginning of an AI unit-economics cycle.
The Token Is a Real Meter and the Wrong Management Unit
A token is not imaginary. It measures part of the work a model performs, gives engineers a way to understand consumption, and gives providers a way to invoice variable usage.
But a token has almost no meaning to a business owner.
A claims executive does not want tokens. She wants a correctly processed claim. A support leader wants a resolved case. An engineering manager wants a safe change merged. A compliance officer wants an investigated exception with an auditable decision.
Tokens sit several layers below those outcomes.
They also fail as a simple comparison unit. One token from a small classifier, one from a frontier reasoning model, and one produced inside a long agent loop are not economically equivalent. Provider price lists now distinguish input, output, cached input, long context, service tiers, tools, and batch processing. The public catalogs from OpenAI , Anthropic , and DeepSeek show how wide the price and product range has become.
Those options are useful for optimization. They do not answer whether the workflow earns its place.
The FinOps Foundation's work on token economics makes the distinction clearly. Token cost is only one layer in a wider stack that includes infrastructure, data, networking, engineering, observability, evaluation, governance, and labor. It also notes that a retrieval pipeline with reasoning and multiple tool calls can consume one or two orders of magnitude more tokens than a direct request to a smaller model.
The provider needs a production meter.
The enterprise needs an economic ledger.
Confusing the two is the root of the rebellion.
Scale Exposes What the Pilot Hid
Pilot economics are forgiving.
The user group is small. The workflow is supervised. Failures are treated as learning. Integration and governance labor are often funded centrally. Consumption may sit inside a trial, a generous subscription, or a budget no one has yet challenged.
Production removes those protections.
Every successful deployment increases volume. Agents add planning, retrieval, tool calls, retries, verification, and memory. More users create more edge cases. Higher autonomy requires better evaluation, monitoring, access control, incident response, and human escalation. A system that looked cheap per interaction can become expensive per completed process.
The evidence is still emerging, but the warning signs are strong. In a May 2026 survey with 75 qualified enterprise respondents, McKinsey found that spend rose nearly fourfold as organizations moved from isolated use cases toward enterprise adoption, while 93 percent reported exceeding their AI budgets. The same analysis cites research finding up to 30-fold variation in token use when an agent executes the same task.
That is a small survey, not a universal market estimate. But it describes a structural problem: agentic demand is nonlinear, while most budgets assume a stable relationship between users, requests, and cost.
This is now a mainstream operating concern. The State of FinOps 2026 report says 98 percent of FinOps teams manage AI spend, up from 31 percent two years earlier, and names AI cost management as the leading skill gap.
The market has moved from Can we build it? to Can we operate it economically?
Karp Is Right About Value, but Outcome Pricing Is Not Magic
Karp's implied alternative is attractive: if a vendor claims to create business value, let it charge for the value instead of the tokens.
Parts of the market are already moving that way.
Intercom now prices its Fin agent by outcomes . A support resolution, procedure handoff, or disqualification costs $0.99; a sales qualification costs $9.99. Zendesk similarly uses automated resolutions as the billing unit for its AI agents.
That is better aligned than charging a customer for every internal inference. It makes the vendor carry more performance risk, and it gives the buyer a unit that resembles work.
It also reveals why outcome pricing is difficult.
An outcome is not a natural object. It is a contract over a definition.
Intercom counts a resolution when the customer confirms success or does not request more help after the last answer. That is a reasonable operational definition, but it is not the same as proving that the customer's underlying problem was solved, that the answer remained correct, or that no downstream cost appeared later.
The difficulty increases outside customer support:
Who gets credit when AI drafts a proposal but a salesperson closes the deal?
Is a merged code change an outcome if it creates a production incident a week later?
Is time saved real value if the organization does not redeploy the capacity?
How should a vendor be paid for risk avoided, cycle time reduced, or a better decision?
Which baseline determines the gain: the previous human process, the best alternative model, or doing nothing?
Outcome pricing can align incentives. It cannot eliminate attribution, delayed effects, gaming, or disagreement over the counterfactual.
Tokens answer a supplier question: how much model capacity was consumed?
Outcome pricing answers a commercial question: what event triggers payment?
Neither, on its own, answers the management question: did this workflow create durable net value?
The Enterprise Needs Workflow-Level Unit Economics
The right unit sits between the token and the annual ROI slide.
It is the business event: a case resolved, claim processed, change reviewed, order recovered, investigation completed, or application approved.
For each event, the organization needs a complete cost and outcome record.
The basic management equation is simple:
Total workflow cost / accepted outcomes = cost per successful outcome
The difficult work is defining “accepted” honestly and measuring the full numerator.
This is why “hours saved” is usually an intermediate metric, not a final one. Saved time becomes economic value only when it changes something observable: more work completed, faster revenue, lower contractor spend, avoided hiring, reduced backlog, better quality, or genuinely redeployed employee capacity.
The same discipline applies to quality. A cheaper model is not cheaper if it doubles review time. A frontier model is not expensive if it prevents a high-consequence error. A human is not inefficient if the automation requires so much monitoring and rework that the total process costs more.
The unit-economics ledger makes those comparisons possible.
Frontier, Smaller, Open, and Human Are Portfolio Choices
The cost backlash will accelerate model substitution, including toward smaller and open-weight systems. That is healthy, but “use a cheaper model” is not a strategy.
Every workflow should be routed to the least expensive execution path that meets its quality, latency, control, and risk requirements.
Frontier models for complex, ambiguous, or high-consequence work where stronger reasoning materially changes the result.
Smaller hosted models for bounded, repetitive, high-volume tasks with measurable acceptance criteria.
Open-weight models where scale, data control, customization, or continuity justifies the infrastructure and operating burden.
Humans for exceptions, contested decisions, relationship-sensitive work, or tasks where review and failure costs erase the automation advantage.
Open weights move cost; they do not make it disappear. Provider margin may be replaced by accelerator capacity, platform engineering, security, evaluation, and operations. The correct comparison is total cost per accepted outcome, not API price per million tokens.
This also changes the role of the frontier model. It becomes a scarce capability used deliberately—not the default engine behind every summarization, classification, and agent loop.
The Control Plane Becomes the Economic System of Record
I previously argued that AI token costs need an operating model , with workload routing, attribution, budgets, approvals, and exceptions enforced close to runtime.
The follow-on is that the same control plane must connect consumption to outcomes.
It should record, for every production workflow:
the models, tools, and data used,
token, infrastructure, and third-party cost,
latency, retries, fallbacks, and failed loops,
human review and correction time,
the evaluation or acceptance result,
and the model or human alternative available next.
Without this layer, a company can cut tokens while destroying value. It can also celebrate productivity while quietly increasing review labor, operational risk, or vendor dependence.
With it, cost control becomes dynamic. The system can route routine work downward, escalate difficult cases upward, cache stable context, stop runaway loops, enforce budget thresholds, and compare providers using the company's own tasks rather than generic benchmarks.
That is more than FinOps reporting.
It is an operating system for the economics of intelligence.
Procurement Will Move Toward Hybrid Contracts
I do not expect token pricing to disappear.
It remains useful for APIs, infrastructure, experimentation, and any workload where the provider cannot observe the customer's final business result. Pure outcome pricing will work best where the event is narrow, frequent, attributable, and quickly verifiable—support resolution is a good example.
Most enterprise agreements will therefore become hybrids:
capacity commitments for predictable base demand,
consumption pricing for variable usage,
service tiers for latency and availability,
outcome components where success can be defined,
and commercial credits or risk sharing when agreed performance is missed.
The buyer's leverage will come from instrumentation and portability. Enterprises should require exportable usage data, workload-level attribution, clear outcome definitions, auditable quality measures, cost ceilings, model-substitution rights, and a tested path to another provider or an open-weight deployment.
The ability to measure and switch is what turns vendor competition into enterprise bargaining power.
The Rebellion Is a Maturity Signal
Enterprises are not discovering that AI has no value.
They are discovering that adoption without economic i

[truncated]
