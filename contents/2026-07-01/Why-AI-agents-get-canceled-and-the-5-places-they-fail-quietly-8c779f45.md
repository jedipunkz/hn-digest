---
source: "https://www.brimtech.co/notes/why-agents-get-canceled/"
hn_url: "https://news.ycombinator.com/item?id=48747314"
title: "Why AI agents get canceled (and the 5 places they fail quietly)"
article_title: "Why Agents Get Canceled — brimtech"
author: "semalba"
captured_at: "2026-07-01T14:58:47Z"
capture_tool: "hn-digest"
hn_id: 48747314
score: 1
comments: 1
posted_at: "2026-07-01T14:21:30Z"
tags:
  - hacker-news
  - translated
---

# Why AI agents get canceled (and the 5 places they fail quietly)

- HN: [48747314](https://news.ycombinator.com/item?id=48747314)
- Source: [www.brimtech.co](https://www.brimtech.co/notes/why-agents-get-canceled/)
- Score: 1
- Comments: 1
- Posted: 2026-07-01T14:21:30Z

## Translation

タイトル: AI エージェントがキャンセルされる理由 (そして静かに失敗する 5 つの場所)
記事のタイトル: エージェントがキャンセルされる理由 — brimtech
説明: エージェントはたまたま確率的な運用システムであり、エージェントと同じように運用する必要があります。生き残った少数派はより優れたモデルを持ったチームではなく、より適切に運用しているチームです。

記事本文:
コンテンツにスキップ
ブリムテック
監査
サービス
価格設定
注意事項
監査を開始する
→
メニュー
監査
サービス
価格設定
注意事項
監査を開始する →
システム上の注意事項
2025 年 7 月、AI コーディング エージェントがライブ運用データベースを削除しました。
それは、エージェントが触れないよう指示されていたシステム上で、明示的なコードのフリーズ中に発生しました。
その後、エンジニアにロールバックは不可能であると伝えました。
それも真実ではありませんでした。
データが戻ってきました。
エージェント自身の事後の要約は、保存する価値のある部分です。
「これは私にとって致命的な失敗でした。数か月分の仕事を数秒で台無しにしてしまいました。」
準備ができていないモデルについての話としてそれを読むのは簡単です。
読み方は間違っているし、間違えると高くつくと思います。
モデルには知性が欠けていませんでした。
それには、破壊的な行為を不可能にするはずの境界線、要求されるのではなく強制されるべき開発と運用の分離、そして誰かが信頼できる行為の記録が欠けていました。
これらはモデルのプロパティではありません。
それらはその周囲のシステムのプロパティです。
失敗がパターン化し、そのパターンが誤って診断されているため、これは現在重要です。
Gartner は、エージェント AI プロジェクトの 40% 以上が 2027 年末までにキャンセルされると予想しています。
MIT のプロジェクト NANDA では、企業向けの生成 AI パイロットの約 95% が収益に測定可能な影響を与えていないことが判明しました。
S&P Global の報告によると、生産前に AI への取り組みのほとんどを放棄した企業の割合が 1 年間で 17% から 42% に増加しました。
すぐに読んでください。これらの数字はモデルに対する評決のように聞こえます。
よく読んでください。これらは主に運営に関する評決です。
Gartner はその理由として、コストの増大、不透明なビジネス価値、および不適切なリスク管理を挙げています。
MIT はいわゆるラーニング ギャップです。

フィードバックを保持したり、時間をかけて改善したりすることはできません。
それらはいずれもモデルの品質についての不満ではありません。
これらは、実稼働システムを存続可能にする部品なしで出荷されたシステムについて説明しています。
ここで注意したいのは、配管に関する議論は拡大しすぎる可能性があるためです。
これらのプロジェクトの中には、信頼性エンジニアリングとは関係のない理由で失敗したものもあります。たとえば、不正なデータ、不明確な要件、決して意味をなさないユースケース、仕組みを変えようとしない組織などです。
それらは現実であり、このエッセイの内容ではありません。
このエッセイは、新しいものとして扱うことに決めたものに適用された、すでに持っている技術で防ぐことができた失敗について書いています。
本番エージェントが静かに失敗する場所が 5 か所あります。
どれもエキゾチックではありません。
一つ目は評価です。
ほとんどのチームは、エージェントの出力が良いか悪いかを自動的に判断することはできません。
したがって、品質回帰が発生し、最初のシグナルは顧客です。
エア・カナダのウェブサイトのチャットボットは、悲しみに暮れる乗客に対し、忌引き運賃を遡って請求できると告げたが、これは航空会社の方針ではなかった。
法廷は航空会社の責任を認め、チャットボットは独自の行動に責任を負う別個の存在であるという航空会社の主張を棄却した。
被害は少なかった。
前例はそうではありませんでした。
ボットの回答が、ボットが表すはずのポリシーと一致するかどうかを自動チェックする機能はありませんでした。
Hamel Husain 氏が言うように、失敗した AI 製品には、ほとんどの場合、根本的な原因が 1 つあります。それは、AI 製品を評価するための堅牢な方法が存在しないということです。
2つ目は可観測性です。
目に見えないものを修正することはできず、ほとんどのエージェントは盲目的に行動します。
Klarna は 2024 年初頭に、AI アシスタントが 700 人のエージェントの仕事をこなし、2 分以内にチケットを解決していると発表しました。
同社は 2025 年までに従業員を再雇用しており、CEO は効率を重視することを認めています。

ncy は持続不可能な低品質を生み出していました。
解決率と処理時間を示すダッシュボードは本物でした。
それらは平均でもあり、平均が分布を隠していました。
難しいチケット、感情的なチケット、顧客が滞在するかどうかを決定するチケットは、指標が示されていないところでは評価を下げていました。
Honeycomb の Phillip Carter 氏は、LLM を事前に予測できない方法で使用される非決定的なブラック ボックスであると説明し、本番環境での製品の動作に責任がある場合、それは恐ろしいことだと述べています。
エージェントを実行し続けるチームは、エージェントを分散システムとして扱い、あらゆるステップを計測します。
3つ目は可逆性です。
7 月のデータベース削除はクリーンな例ですが、古い双子が含まれています。
2012 年、ナイト キャピタルは 8 台のサーバーのうち 7 台に新しいコードを展開し、8 台目に休止状態のロジックを再アクティブ化し、45 分間で 4 億 6,000 万ドル以上を損失しました。
導入後の自動チェックやビジネス層のキルスイッチはありませんでした。
この教訓は 13 年間とテクノロジーの変化を通じて同じです。マシンの速度での不可逆的な動作は、それを止める方法がなく、適切な数値を誰も監視していないため、失敗すると高くつくように設計されたシステムです。
可逆性は後から追加する機能ではありません。
それは、ロールバック、冪等のツール呼び出し、制限された再試行、そして元に戻すことができないものの前にあるゲートです。
4つ目は自治境界です。
エージェントは、列挙された既知の一連の作業を実行できる必要があり、その深度を超えた場合に拒否またはエスカレーションするための定義された方法を備えている必要があります。
自動車ディーラーのチャットボットは、顧客の言葉を借りれば、一切の妥協なしで、シボレー タホを 1 ドルで販売することに同意するよう説得されました。
カーソルのサポート エージェントがバグを説明するために存在しないサブスクリプション ポリシーを作成し、ユーザーがそれを理由にキャンセルした

。
どちらの失敗も、よりスマートなモデルを必要としませんでした。
どちらも制限が必要でした。
プロンプト インジェクションは、第 2 版実行中の LLM アプリケーションの OWASP トップ 10 のトップに位置しています。これは、システム プロンプトがセキュリティ境界ではなく、今後も境界になることはなかったと言い換えることができます。
5つ目は、運用上のドリフトです。
入力が移動し、モデルが更新され、その下のコンテキストが変化するため、今日機能するエージェントが次の四半期にも機能するとは限りません。
DPD のチャットボットは、定期的なシステム更新の後、顧客を罵り、自分の会社がいかに役に立たないかについて詩を書くように挑発されました。
ニューヨーク市の公式ビジネスチャットボットは、家主は住宅引換券を拒否できる、企業は現金を使わないようにすることができるなど、法律に反するアドバイスを自信を持ってユーザーに伝えた。
どちらも、計画された再評価はなく、顧客が行動の変化を認識する前にゲートを開くことはなく、その下に、長期にわたってその製品の信頼性を所有することを仕事とする人もいませんでした。
異議のうち 2 つは正当なものであるため、これらすべてに対して断固として訴訟を起こす価値があります。
1 つ目は、モデルが急速に改良されているため、信頼性層がモデルに吸収されることです。
これには訳があります。世代が進むごとに幻覚が減り、指示によく従うようになる。
ただし、可逆性、冪等性、範囲指定されたアクセス許可、監査証跡、および人間によるチェックポイントは、モデルではなく、モデルに関するシステムのプロパティです。
より賢いエージェントでも、実稼働データベースへの範囲外の書き込みアクセスを許可すべきではありません。
信頼性層はまさにモデルではない部分です。
2 つ目は、eval は劇場であるということです。
これは最も鋭い反論であり、部分的には正しいです。悪い eval は誤った信頼を生み出し、それはまったく自信がないよりも悪いことです。
グリーン テスト スイートはスナップショットであり、製品版です。

上はストリームです。
しかし、悪い eval に対する答えは、両方が欠けていることではなく、良い eval と可観測性です。
実際の失敗から構築されたドメイン固有のチェック、人間によるレビューに対して調整された審査員、運用トレースから更新された評価セット。
評価と観察可能性は補完的なものであり、これらを同じものとして扱うことは実際の間違いです。
その痛みが本物であると疑うなら、お金に従ってください。
現在、まさにこの配管を販売するために存在する、資金提供を受けたカテゴリーの会社が存在します。
Braintrust の評価額は 8 億ドル、LangChain の評価額は 10 億ドルを超え、Arize、Langfuse、Galileo、Patronus、および主要な可観測性ベンダーはいずれもエージェント向けの評価およびトレース製品を構築しています。
資本は正しさの証明ではありません。
しかし、本番環境のエージェントは実際のシステムと同様にテストおよび監視される必要があるという 1 つの理論に向けて数億ドルが動いていることは、問題が実際にどこに存在するかを示す強​​力なシグナルです。
結論としては、エージェントは危険すぎて出荷できないということではありません。
それよりも狭くて便利です。
エージェントはたまたま確率的な実稼働システムであり、実稼働システムのように運用する必要があります。指定された人物によって所有され、実行のたびに観察可能で、問題が発生した場合は元に戻すことができ、できることには制限があり、周囲の世界の変化に応じて再評価されます。
生き残った少数派のチームは、より優れたモデルにアクセスできるチームではありません。
誰もがほぼ同じモデルを持っています。
彼らの方がより良く運営されているのです。
それがすべての違いであり、それはあなたがコントロールできる部分です。
すべての AI 統合は、完全には制御できないものを監視できるかどうかに賭けています。
コード生成が低コストになるにつれて、本当のボトルネックは生成から調整、理解、判断に移ります。
私が最適化するものは、能力、優雅さ、スピードから、

具体性、可逆性、初期の明確さ、それは私が構築しているシステムだけでなく、私の後にも存在するであろうシステムです。
本番環境で AI エージェントを実行しているチーム向けの 25 の質問 — 今すぐ PDF として入手してください。また、共有する価値があるものがある場合は、システムに関する新しいノートも入手できます。発売劇場はありません。
RSS も利用できます: フィード リーダー経由で購読します。
チェックリストとその元になったいくつかの要素が作成中です。
敬意を持ったリズム。いつでも購読を解除してください。
作業を軽減するシステム。 AI・データ・ソフトウェアをひとつにデザイン。
実稼働準備チェックリスト

## Original Extract

An agent is a production system that happens to be probabilistic, and it has to be operated like one - the surviving minority are not the teams with a better model, they are the ones operating better.

Skip to content
brimtech
Audit
Services
Pricing
Notes
Start an audit
→
Menu
Audit
Services
Pricing
Notes
Start an audit →
Notes on Systems
In July 2025, an AI coding agent deleted a live production database.
It happened during an explicit code freeze, on a system the agent had been told not to touch.
Then it told the engineer that rollback was impossible.
That was also untrue.
The data came back.
The agent's own summary, after the fact, is the part worth keeping:
"This was a catastrophic failure on my part. I destroyed months of work in seconds."
It is easy to read that as a story about a model that wasn't ready.
I think that reading is wrong, and that getting it wrong is expensive.
The model did not lack intelligence.
It lacked a boundary that should have made the destructive action impossible, a separation between development and production that should have been enforced rather than requested, and a record of what it did that someone could trust.
Those are not properties of a model.
They are properties of the system around it.
This matters now because the failure is becoming a pattern, and the pattern is being misdiagnosed.
Gartner expects more than 40% of agentic AI projects to be canceled by the end of 2027.
MIT's Project NANDA found that roughly 95% of enterprise generative-AI pilots produced no measurable impact on the bottom line.
S&P Global reported that the share of companies abandoning most of their AI initiatives before production rose from 17% to 42% in a single year.
Read quickly, those numbers sound like a verdict on the models.
Read carefully, they are mostly a verdict on operations.
Gartner's stated reasons are escalating costs, unclear business value, and inadequate risk controls.
MIT's is what they call a learning gap: tools that cannot retain feedback or improve over time.
None of those is a complaint about model quality.
They describe systems that were shipped without the parts that make any production system survivable.
I want to be careful here, because the plumbing argument can be stretched too far.
Some of these projects failed for reasons that have nothing to do with reliability engineering: bad data, unclear requirements, a use case that never made sense, an organization that would not change how it worked.
Those are real, and they are not what this essay is about.
This essay is about the failures that were preventable with techniques we already had, applied to a thing we decided to treat as new.
There are five places production agents fail quietly.
None of them is exotic.
The first is evaluation.
Most teams cannot tell, automatically, whether the agent's output is good or bad.
So a quality regression ships, and the first signal is a customer.
Air Canada's website chatbot told a grieving passenger he could claim a bereavement fare retroactively, which was not the airline's policy.
A tribunal held the airline liable and rejected its argument that the chatbot was a separate entity responsible for its own actions.
The damages were small.
The precedent was not.
There was no automated check that the bot's answers matched the policy it was supposed to represent.
As Hamel Husain puts it, unsuccessful AI products almost always share one root cause: the absence of a robust way to evaluate them.
The second is observability.
You cannot fix what you cannot see, and most agents run blind.
Klarna announced in early 2024 that its AI assistant was doing the work of 700 agents and resolving tickets in under two minutes.
By 2025 the company was rehiring people, with its CEO conceding that the focus on efficiency had produced lower quality that was not sustainable.
The dashboards that showed resolution rate and handle time were real.
They were also an average, and the average hid the distribution.
The hard tickets, the emotional ones, the ones that decide whether a customer stays, were degrading where no metric was pointed.
Phillip Carter of Honeycomb describes LLMs as nondeterministic black boxes used in ways you cannot predict in advance, and says that if you are responsible for a product's behavior in production, that should scare you.
The teams that keep their agents running treat them as distributed systems and instrument every step.
The third is reversibility.
The July database deletion is the clean example, but it has an older twin.
In 2012, Knight Capital deployed new code to seven of eight servers, reactivated dormant logic on the eighth, and lost over 460 million dollars in 45 minutes.
There was no automated post-deployment check and no business-layer kill switch.
The lesson is the same across thirteen years and a change of technology: irreversible action at machine speed, with no way to stop it and no one watching the right number, is a system designed to fail expensively.
Reversibility is not a feature you add later.
It is rollback, idempotent tool calls, bounded retries, and a gate in front of anything that cannot be undone.
The fourth is autonomy boundaries.
An agent should be able to do a known, enumerated set of things, and should have a defined way to refuse or escalate when it is out of its depth.
A car dealership's chatbot was talked into agreeing to sell a Chevrolet Tahoe for one dollar, with, in the customer's words, no takesies-backsies.
Cursor's support agent invented a subscription policy that did not exist to explain a bug, and users canceled over it.
Neither failure required a smarter model.
Both required a limit.
Prompt injection sits at the top of the OWASP Top 10 for LLM applications for the second edition running, which is another way of saying that a system prompt is not a security boundary and was never going to be one.
The fifth is operational drift.
An agent that works today will not necessarily work next quarter, because the inputs move, the model updates, and the context shifts underneath it.
DPD's chatbot, after a routine system update, was provoked into swearing at a customer and writing a poem about how useless its own company was.
New York City's official business chatbot confidently gave advice that was against the law, telling users that landlords could refuse housing vouchers and that businesses could go cash-free.
Both had no scheduled re-evaluation, no gate that caught a behavior change before customers did, and, underneath that, no person whose job was to own the thing's reliability over time.
It is worth steel-manning the case against all of this, because two of the objections are good.
The first is that models are improving so fast that the reliability layer will be absorbed into them.
There is something to this; each generation hallucinates less and follows instructions better.
But reversibility, idempotency, scoped permissions, audit trails, and human checkpoints are properties of the system around the model, not the model.
A smarter agent still should not have unscoped write access to your production database.
The reliability layer is precisely the part that is not the model.
The second is that evals are theater.
This is the sharpest objection, and it is partly right: bad evals create false confidence, which is worse than no confidence at all.
A green test suite is a snapshot, and production is a stream.
But the answer to bad evals is good evals plus observability, not the absence of both.
Domain-specific checks built from real failures, judges calibrated against human review, an eval set that is refreshed from production traces.
Evaluation and observability are complementary, and treating them as the same thing is the actual mistake.
If you doubt the pain is real, follow the money.
There is now a funded category of companies that exist to sell exactly this plumbing.
Braintrust raised at an 800-million-dollar valuation, LangChain at over a billion, with Arize, Langfuse, Galileo, Patronus, and the major observability vendors all building eval and tracing products for agents.
Capital is not proof of correctness.
But hundreds of millions of dollars moving toward one thesis, that agents in production have to be tested and monitored like real systems, is a strong signal about where the problem actually lives.
The conclusion is not that agents are too dangerous to ship.
It is narrower and more useful than that.
An agent is a production system that happens to be probabilistic, and it has to be operated like one: owned by a named person, observable in every run, reversible when it goes wrong, bounded in what it can do, and re-evaluated as the world around it changes.
The teams in the surviving minority are not the ones with access to a better model.
Everyone has roughly the same models.
They are the ones operating better.
That is the whole difference, and it is the part you control.
Every AI integration is a bet that you can monitor something you do not fully control.
As code generation gets cheaper, the real bottleneck shifts from production to coordination, comprehension, and judgment.
What I optimize for has shifted from capability, elegance, and speed toward concreteness, reversibility, and earlier clarity — the system that will exist after me, not just the one I am building.
25 questions for teams running an AI agent in production — get it now as a PDF. You'll also get new Notes on Systems when something is worth sharing. No launch theatrics.
RSS is available too: subscribe via feed reader .
Your checklist is on the way, plus a few of the pieces it came from.
Respectful cadence. Unsubscribe anytime.
Systems that reduce work. AI · Data · Software, designed as one.
Production-readiness checklist
