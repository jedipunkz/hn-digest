---
source: "https://venturebeat.com/orchestration/nvidias-switchyard-router-reshuffles-ai-models-mid-task-cutting-task-costs-to-a-third-in-its-own-tests"
hn_url: "https://news.ycombinator.com/item?id=49388335"
title: "Nvidia's Switchyard router reshuffles AI models mid-task to cut task costs"
article_title: "Nvidia's Switchyard router reshuffles AI models mid-task, cutting task costs to a third in its own tests | VentureBeat"
image: "https://images.ctfassets.net/jdtwqhzvc2n1/6qbSNxUnm1pPACzx2jzAAQ/18b1af901de9038b19d1069310092946/lightning-nemotron-smk1.jpg?w=800&q=75"
author: "gmays"
captured_at: "2026-08-21T14:25:25Z"
capture_tool: "hn-digest"
hn_id: 49388335
score: 2
comments: 0
posted_at: "2026-08-21T14:10:29Z"
tags:
  - hacker-news
  - translated
---

# Nvidia's Switchyard router reshuffles AI models mid-task to cut task costs

- HN: [49388335](https://news.ycombinator.com/item?id=49388335)
- Source: [venturebeat.com](https://venturebeat.com/orchestration/nvidias-switchyard-router-reshuffles-ai-models-mid-task-cutting-task-costs-to-a-third-in-its-own-tests)
- Score: 2
- Comments: 0
- Posted: 2026-08-21T14:10:29Z

## Translation

タイトル: Nvidia の操車場ルーターはタスクの途中で AI モデルを再シャッフルし、タスクのコストを削減します
記事のタイトル: Nvidia の Switchyard ルーター、タスクの途中で AI モデルを再シャッフルし、独自のテストでタスクのコストを 3 分の 1 に削減 |ベンチャービート
説明: Nvidia の Nemotron 3.5 Lightning は、NeMo Switchyard ルーターと組み合わせて、タスクの途中でモデルを再割り当てし、Nvidia 独自のテストでタスクのコストを 3 分の 1 に削減します。

記事本文:
Nvidia の Switchyard ルーターはタスクの途中で AI モデルを再シャッフルし、独自のテストでタスクのコストを 3 分の 1 に削減 | VentureBeat オーケストレーション
Nvidia の操車場ルーターはタスクの途中で AI モデルを再シャッフルし、独自のテストでタスクのコストを 3 分の 1 に削減
クレジット: FLUX-2-Pro を使用して VentureBeat によって生成された画像
常時稼働の AI エージェントを実行している企業は、同じトレードオフに直面し続けています。すべてのタスクをフロンティア モデルに送信すると、請求額が急速に上昇します。カスタム ルーティング ロジックを構築して簡単なタスクを安価なモデルに送信すると、それが独自のエンジニアリング プロジェクトとなり、ワークフローが変更されるたびに維持する必要があります。
Nvidia は、この問題の両端に同時に触れる修正を提案しています。同社は火曜日に、エージェントのワークフローの各ステップを最適なモデルにルーティングするオープンソース ライブラリである NeMo Switchyard と並行して、大量の特殊なエージェント タスク向けに構築された 300 億パラメータのオープンな専門家混合モデルである Nemotron 3.5 Lightning を発表しました。
見出しの数字: Nvidia によると、Lightning は同クラスの同等モデルよりも最大 4 倍高速な出力を実現し、マッチング精度で Qwen3.6-35B よりも約 30% 速くエージェント タスクを完了します。 Switchyard を介して組み合わせることにより、Opus 4.8 を単独で実行する場合の約 3 分の 1 にベンチマーク コストを削減しながら、フロンティア レベルのタスクの完了を維持できると Nvidia は述べています。
このタイミングにより、Nvidia は業界がここ数カ月で最も多忙な無差別級試合の真っ只中にいることになります。 Alibaba、Moonshot、Zhipu、DeepSeekはいずれも、春以降、競争力のあるオープンモデルを中国から出荷しており、そのいくつかは規模や価格で米国の研究所を下回りながら、最先端かそれに近い性能に達している。 Meta は、独自の 300 億パラメータのオープン エージェント モデル Muse Glimmer をリリースすることで、その圧力をさらに高めました。オープンウェイトは、

数か月以内に差別化要因を獲得し、Nvidia のリリースはその変化に先んじてではなく、まさにその中にあります。
ペアリングがポイントです。モデルだけではコストの問題は解決できませんし、ルーターだけでは効率的なルーティング先が何もありません。 Nvidia は、モデル層とルーティング層の両方に適用されるオープンソースこそが、実際にエージェント AI のコストを大きく左右するものであり、単一の安価なモデルや、他人のスタックにボルトで取り付けられたよりスマートなルーターではないと賭けています。
Switchyard の本当のライバルは他のオープン モデルではありません。すでに OpenRouter の Auto モードを強化している Not Diamond と、UC Berkeley と LMSYS のオープンソース フレームワークである RouteLLM です。どちらも独自のモデルは出荷しません。 Nvidia の賭けは、1 つのオープン ライセンスの下で決定の両側を所有することは、ルーターのみまたはモデルのみの競合他社が太刀打ちできないことです。
Nvidia の生成 AI 担当バイスプレジデントであるカリ ブリスキー氏は、「これがモデル システムの力であり、ワークフローの各ステップに適切なモデルを適合させます」とブリーフィングで述べました。
ルーターが実際にワークフローをどのように変えるか
モデル ルーティングは新しいカテゴリではありません。 OpenRouter、LiteLLM、およびいくつかのスタンドアロン ルーティングのスタートアップはすでに、開発者が複数のプロバイダー間でトラフィックをポイントできるようにしています。開閉所は、それらを完全に交換するのではなく、それらのいくつかに接続します。
Switchyard が解決する中心的な問題は、エージェントがタスクを進めるにつれて適切なモデルが変化するということです。ツールが結果を返したり、エラーが発生したり、ステップが複雑ではなく日常的なものになったりすると、エージェントの状態は変化します。固定されたモデルの選択はそのいずれにも適応できません。
Briski 氏は、静的なタスク カテゴリではなく、その変化する状態に対応するルーティング戦略について説明しました。
「さまざまな種類のルーティング戦略がある」とブリスキー氏は言う。 「ランダムなルーターを使用することもできますが、それはそれほど優れたものではありません。

または、エージェント状態ルートまたは分類子ルートを使用することもできます。ルーティング戦略に応じて、最適なモデルを選択する必要があります。場合によっては、本当に効率的なタスクを実現するために Lightning のようなモデルを使用する必要があり、モデルのプールに設定されている場合、ルーターは実際に Lightning を選択します。」
コストは後付けではなく、ルーティングの決定に直接影響します。 VentureBeat からの質問に答えて、Briski 氏は、Switchyard はモデルの冗長性、つまり特定のモデルがタスクに対して生成する傾向にあるトークンの数を評価し、その予測を使用して、呼び出しが行われる前に、より安価なオプションに作業を誘導できると述べました。
これが独自の統合プロジェクトにならないようにしているのは、Switchyard が置かれている場所です。 Nvidia はパートナーを 2 つのグループに分割しました。Cognition、LangChain、Nous Research などの Switchyard を直接呼び出すエージェント フレームワークと、Kong、LiteLLM、OpenRouter などの自社製品に Switchyard サポートを組み込んだ LLM ゲートウェイです。 Kong は、Switchyard を Kong AI ゲートウェイ内にネイティブに搭載しています。 Briski 氏は、ライブラリが既存のルーティング エコシステムにどのように適合するかを説明する際に、同じゲートウェイ パートナーのリストを指摘しました。
「私たちはエコシステムを愛する者であり、確実に統合されるようにしたいと考えています」とブリスキー氏は語った。 「私たちは OpenRouter、LiteLLM、Kong と提携しており、彼らはすでに私たちのルーティング アルゴリズムを統合しているので、すでに最高のツールを使用している場所でそのアルゴリズムを利用することができます。」
Nvidia は、Switchyard をテストした 9 社の結果を共有しました。そのうちのいくつかには具体的な数値が添付されていました。 LangChain は、6% の精度のトレードオフで、わずか 7% のコールをフロンティア モデルにルーティングすることにより、145 のマルチターン Deep Agent タスク全体で 74% のコスト削減を報告しました。 Ramp は、Ramp SWE-Bench でのフロンティア モデルのパフォーマンスに匹敵し、コストを 58%、ランタイムを 33% 削減したと述べています。コグニット

イオンは、内部使用のために Switchyard のステージド ルーターを Devin Desktop に統合し、すべてを単一のフロンティア モデルにルーティングする場合と比較して平均コストを 28% 削減しながら、FrontierCode Main でフロンティアに近いパフォーマンスを報告しました。
Lightning のアーキテクチャとパフォーマンスの向上
Nemotron 3.5 Lightning は、それ自体がスタンドアロンのオープン モデルであり、汎用用途ではなく、大量の特殊なエージェント タスク向けに構築されています。
これは、Nvidia が 2025 年 12 月に Nemotron 3 ファミリで導入した潜在的な専門家混合アーキテクチャであるハイブリッド Mamba-Transformer を拡張します。これは、Nvidia がトレーニング後の比較で Lightning 独自のベースラインとして使用する、Nemotron 3 Super の背後にある同じラインです。 Switchyard のようなルーティング設定内に配置され、フロンティアエンドではなく、高速かつ安価な意思決定のエンドに位置するように構築されていますが、どのルーターからも独立して実行および出荷されます。
9 つの評価にわたる一般的な能力ベンチマークである Artificial Analysis Intelligence Index によると、Lightning スコアは 24 で、gpt-oss-120b と同点で、Nemotron 3 Super、Gemma 4 31B、Claude 4.5 Haiku、Mistral Medium 3.5 のすべて 30 に次ぐスコアです。 Lightning は、その規模クラスでは汎用インテリジェンスのリーダーではなく、NVIDIA はそれを主張していませんです。
実際の主張はもっと狭いです。Nvidia が提供した PinchBench データによると、コーディング、研究、ファイル管理にまたがる現実世界のエージェント タスクのベンチマークである PinchBench では、Lightning は Qwen3.6-35B の精度に約 30% 速く匹敵し、同様の完了時間で Gemma 4 26B の精度を上回っています。これは速度と精度のトレードオフであり、機能を優先するものではありません。
Nvidia によれば、トレーニング後はより大きな効果が現れるとのことです。同社は、早期アクセス パートナー 4 社からのビフォーアフターの数値を共有しました: Nemotron 3 Super ベースライン、CodeRab に対する CrowdStrike の悪意のあるコンテンツのリコール

bit のコーディング ルーターは GPT 5.4 Nano ベースラインに対して、Harvey and Trajectory の法的タスクの完了は Opus 4.6 ベースラインに対して、Lila Sciences のエネルギー シミュレーションは Opus 4.8 ベースラインに対して動作します。 CodeRabbit のケースは最も具体的です。Nvidia によれば、標準 NeMo Auto モデル レシピは 1 エポックでトレーニングされ、約 2 時間で 85 ドルで動作するルーター エージェントに組み込まれます。
これが企業にとって何を意味するか
成長を続けるオープン モデル市場では、競争力のある製品が不足することはありません。新しい Nemotron Lightning リリースは、組織が検討すべきもう 1 つの選択肢となるでしょう。
モデル側では、Lightning 独自のベンチマーク チャートでは、直接の比較ポイントとして Qwen3.6-35B が選択されています。 VentureBeat から直接、Lightning を中国のモデルとより広範に比較する方法を尋ねられたところ、Briski 氏は直接のベンチマークを提供せず、代わりにオープン性とカスタマイズ性が差別化要因であると指摘しました。
「私たちの価値提案はオープンなだけではなく、非常にカスタマイズ可能です」とブリスキー氏は語った。
エージェント インフラストラクチャを構築している企業には、次の 3 つの傾向が際立っています。
ルーティングの決定は、静的ではなく動的になりつつあります。単一のデフォルト モデルを中心にエージェント パイプラインを構築した企業は、設計時に設定された固定割り当てではなく、エージェントの状態やトークン コストなどのライブ シグナルに基づくステップごとのルーティングを推進しています。
オープンソースは現在、1 層ではなく 2 層のコストレバーになっています。ベンダーがエンドツーエンドで制御するオープン モデルとオープン ルーターを組み合わせるのは、より安価な重みだけを使用するよりも新しい議論であり、他のラボが同じパターンに従うかどうかに注目する価値があります。
競争上の課題は、最高のモデルから最高のシステムに移ります。ルーティング ライブラリが成熟するにつれて、差別化要因は、企業がデフォルトで使用するモデルから、ルーティング層が実稼働環境のタスクにモデルをどの程度適切に一致させるかという方向に移行し、より困難になります。

ベンチマークを行う必要がありますが、市場に出すのはさらに困難です。
私の個人情報を販売または共有しないでください
私の機密個人情報の使用を制限する
© 2026 ベンチャービート。無断転載を禁じます。

## Original Extract

Nvidia's Nemotron 3.5 Lightning pairs with its NeMo Switchyard router, which reassigns models mid-task and cuts task costs to a third in Nvidia's own tests.

Nvidia's Switchyard router reshuffles AI models mid-task, cutting task costs to a third in its own tests | VentureBeat Orchestration
Newsletters Nvidia's Switchyard router reshuffles AI models mid-task, cutting task costs to a third in its own tests
Credit: Image generated by VentureBeat with FLUX-2-Pro
Enterprises running always-on AI agents keep hitting the same tradeoff. Send every task to a frontier model and the bill climbs fast. Build custom routing logic to send easy tasks to cheaper models and that becomes its own engineering project, one that has to be maintained every time a workflow changes.
Nvidia is proposing a fix that touches both ends of that problem at once. The company is out on Tuesday with Nemotron 3.5 Lightning, a 30-billion-parameter open mixture-of-experts model built for high-volume, specialized agent tasks, alongside NeMo Switchyard, an open-source library that routes each step of an agent workflow to whichever model fits it best.
The headline numbers: According to Nvidia, Lightning delivers up to 4x faster output than comparable models in its class, completing agentic tasks roughly 30% faster than Qwen3.6-35B at matching accuracy. Paired through Switchyard, Nvidia says the combination holds frontier-level task completion while cutting benchmark costs to roughly a third of running Opus 4.8 alone.
The timing puts Nvidia in the middle of the busiest open-weight stretch the industry has seen in months. Alibaba, Moonshot, Zhipu and DeepSeek have all shipped competitive open models out of China since the spring, several landing at or near frontier performance while undercutting US labs on size or price. Meta added to that pressure by releasing its own 30-billion-parameter open agentic model, Muse Glimmer . Open weights have gone from a differentiator to table stakes in a matter of months, and Nvidia's release lands squarely inside that shift rather than ahead of it.
The pairing is the point. A model alone doesn't solve the cost problem, and a router alone has nothing efficient to route to. Nvidia is betting that open source, applied at both the model layer and the routing layer, is what actually moves the cost needle on agentic AI, not a single cheaper model and not a smarter router bolted onto someone else's stack.
Switchyard's real rivals aren't other open models — they're Not Diamond, which already powers OpenRouter's Auto mode, and RouteLLM, the open-source framework from UC Berkeley and LMSYS. Neither ships its own model. Nvidia's bet is that owning both sides of the decision, under one open license, is what a router-only or model-only competitor can't match.
"That is the power of a system of models, matching the right model to each step of the workflow," Kari Briski, vice president of generative AI at Nvidia, said in a briefing.
How the router actually changes the workflow
Model routing isn't a new category. OpenRouter, LiteLLM and a handful of standalone routing startups already let developers point traffic across multiple providers. Switchyard plugs into several of them rather than replacing them outright.
The core problem Switchyard solves is that the right model changes as an agent moves through a task. An agent's state shifts as tools return results, errors show up, or a step turns out to be routine rather than complex, and a fixed model choice can't adapt to any of that.
Briski described routing strategies that respond to that shifting state rather than a static task category.
"It has many types of routing strategies," Briski said. "You can have a random router, which is not that great, or you can have an agent state route or a classifier route. Depending on your routing strategy, it wants to choose the best model. In some cases you want to go with a model like Lightning for really efficient tasks, and the router will actually choose Lightning if it's set up in your pool of models."
Cost enters the routing decision directly, not as an afterthought. In response to a question from VentureBeat , Briski said Switchyard can evaluate model verbosity, meaning how many tokens a given model tends to produce for a task, and use that prediction to steer work toward the cheaper option before the call is made.
The part that keeps this from becoming its own integration project is where Switchyard sits. Nvidia split its partners into two groups: agent frameworks that call Switchyard directly, including Cognition, LangChain and Nous Research, and LLM gateways that have built Switchyard support into their own products, including Kong, LiteLLM and OpenRouter. Kong ships Switchyard natively inside Kong AI Gateway. Briski pointed to that same list of gateway partners when describing how the library fits into the existing routing ecosystem.
"We are an ecosystem lover, and we want to make sure that we are integrated," Briski said. "We've partnered with OpenRouter, LiteLLM and Kong, and they've already integrated our routing algorithm, so you can pick it up right where you're already using the best tools."
Nvidia shared results from nine companies testing Switchyard, several with specific figures attached. LangChain reported a 74% cost reduction across 145 multi-turn Deep Agents tasks by routing just 7% of calls to a frontier model, at a 6% accuracy tradeoff. Ramp said it matched a frontier model's performance on Ramp SWE-Bench while cutting costs 58% and runtime 33%. Cognition integrated Switchyard's staged router into Devin Desktop for internal use and reported near-frontier performance on FrontierCode Main while cutting mean cost 28% relative to routing everything to a single frontier model.
Lightning's architecture and performance gains
Nemotron 3.5 Lightning is a standalone open model in its own right, built for high-volume, specialized agent tasks rather than general-purpose use.
It extends the hybrid Mamba-Transformer, latent mixture-of-experts architecture Nvidia introduced with the Nemotron 3 family in December 2025 , the same line behind Nemotron 3 Super , which Nvidia uses as Lightning's own baseline in its post-training comparisons. Positioned within a routing setup like Switchyard, it's built to sit at the fast, cheap end of the decision rather than the frontier end, but it runs and ships independent of any router.
According to the Artificial Analysis Intelligence Index , a general capability benchmark spanning nine evaluations, Lightning scores 24, tied with gpt-oss-120b and behind Nemotron 3 Super, Gemma 4 31B, Claude 4.5 Haiku and Mistral Medium 3.5, all at 30. Lightning isn't a general-intelligence leader in its size class, and Nvidia isn't claiming it is.
The actual claim is narrower: according to PinchBench data supplied by Nvidia, Lightning matches Qwen3.6-35B's accuracy roughly 30% faster and beats Gemma 4 26B's accuracy at a similar completion time on PinchBench, a real-world agent task benchmark spanning coding, research and file management. That's a speed-to-accuracy tradeoff, not a capability win.
Post-training is where Nvidia says the bigger gains show up. The company shared before-and-after figures from four early-access partners: CrowdStrike's malicious-content recall against a Nemotron 3 Super baseline, CodeRabbit's coding router against a GPT 5.4 Nano baseline, Harvey and Trajectory's legal task completion against an Opus 4.6 baseline, and Lila Sciences' energy simulation work against an Opus 4.8 baseline. CodeRabbit's case is the most specific: Nvidia says the standard NeMo Auto model recipe, trained for one epoch, built into a working router agent for $85 in about two hours.
What this means for enterprises
There is no shortage of competitive offerings in the growing market for open models. The new Nemotron Lightning release will be yet another option for organizations to consider.
On the model side, Lightning's own benchmark chart picks Qwen3.6-35B as its direct comparison point. Asked by VentureBeat directly how Lightning compares to Chinese models more broadly, Briski didn't offer a head-to-head benchmark, pointing instead to openness and customizability as the differentiator.
"Our value proposition is not just open and it's very customizable," Briski said.
For enterprises building agentic infrastructure, three trends stand out:
The routing decision is becoming dynamic instead of static. Enterprises that built agent pipelines around a single default model are being pushed toward per-step routing based on live signals like agent state and token cost, not a fixed assignment set at design time.
Open source is now a cost lever at two layers, not one. Pairing an open model with an open router a vendor controls end to end is a newer argument than cheaper weights alone, and worth watching for whether other labs follow the same pattern.
The competitive question shifts from best model to best system. As routing libraries mature, the differentiator moves from which model an enterprise defaults to, toward how well its routing layer matches models to tasks in production, a harder thing to benchmark and a harder thing to market.
Do Not Sell or Share My Personal Information
Limit the Use Of My Sensitive Personal Information
© 2026 VentureBeat. All rights reserved.
