---
source: "https://nanonets.com/blog/what-is-a-context-graph/"
hn_url: "https://news.ycombinator.com/item?id=48798442"
title: "Context graphs: how AI agents can store and use past decisions"
article_title: "Context graphs: how AI agents remember why decisions were made"
author: "vitaelabitur"
captured_at: "2026-07-05T22:52:53Z"
capture_tool: "hn-digest"
hn_id: 48798442
score: 5
comments: 0
posted_at: "2026-07-05T22:16:31Z"
tags:
  - hacker-news
  - translated
---

# Context graphs: how AI agents can store and use past decisions

- HN: [48798442](https://news.ycombinator.com/item?id=48798442)
- Source: [nanonets.com](https://nanonets.com/blog/what-is-a-context-graph/)
- Score: 5
- Comments: 0
- Posted: 2026-07-05T22:16:31Z

## Translation

タイトル: コンテキスト グラフ: AI エージェントが過去の決定を保存および使用する方法
記事のタイトル: コンテキスト グラフ: AI エージェントが意思決定が行われた理由をどのように記憶するか
説明: エージェントが意思決定を取得する必要がある理由、コンテキスト グラフがこれを実現する方法、およびエージェントがこれらのグラフを使用する方法について説明します。

記事本文:
ナノネット
製品
事前構築されたエージェント
APエージェント
調整エージェント
注文管理エージェント
現金申請代理人
サプライヤーオンボーディングエージェント
月末締め担当者
RCMエージェント
データ抽出エージェント 開発者向け
エージェントプラットフォーム
Agent Builder カスタム エージェントの構築
コンテキストグラフ
ドキュメントインテリジェンス
ドキュメントの生成
例外管理
分析
ワークフローオーケストレーション
統合
ソリューション
チーム別
金融
会計
運営
サプライチェーン
RCM
ソフトウェア
業界別
製造業
日用品雑貨
ヘルスケア
金融サービス
物流とサプライチェーン
お客様
価格設定
リソース
私たちの使命、チーム、投資家について。
Nanonets チームによるブログ ガイド、研究、製品の最新情報。
Docs Platform のドキュメント、クイックスタート、統合ガイド。
API リファレンス Nanonet 上に構築するための完全な API リファレンス。
Nanonets チームからのリサーチ モデルのリリース、ベンチマーク、およびリサーチ ノート。
仕事を完成させる AI を構築するチームに参加してください。
パートナーシップ 当社のグローバル パートナー プログラムに参加して、IDP 実践を拡大してください。
デモを予約する
始めましょう
製品
事前構築されたエージェント
APエージェント
調整エージェント
注文管理エージェント
現金申請代理人
サプライヤーオンボーディングエージェント
月末締め担当者
RCMエージェント
データ抽出エージェント
エージェントプラットフォーム
エージェントビルダー
コンテキストグラフ
ドキュメントインテリジェンス
ドキュメントの生成
例外管理
分析
ワークフローオーケストレーション
統合
ソリューション
チーム別
金融
会計
運営
サプライチェーン
RCM
ソフトウェア
業界別
製造業
日用品雑貨
ヘルスケア
金融サービス
物流とサプライチェーン
お客様
価格設定
リソース
について
ブログ
ドキュメント
APIリファレンス
研究
キャリア
パートナーシップ
デモを予約する
始めましょう
コンテキスト グラフ: AI エージェントが意思決定が行われた理由をどのように記憶するか
によって
カラン
2026 年 7 月 1 日
リンクがコピーされました!
コピーに失敗しました!
目次
シェアする
リンクコピー

エド！
コピーに失敗しました!
AI エージェントでは、コンテキスト グラフは意思決定をキャプチャするエージェント メモリの一部です。
この投稿では、エージェントが意思決定を取得する必要がある理由、コンテキスト グラフがこれを実現する方法、およびエージェントがこれらのグラフを使用する方法について説明します。
四半期の最後の週です。更新エージェントは 48 万ドルのアカウントを運用しています。顧客は 20% オフを希望し、そうでなければ歩きます。エージェントの指示には、10 万ドルを超えるアカウントは解約してはならないと記載されていますが、エージェントのポリシーでは更新の上限が 10% に設定されています。さて、何でしょうか？
人間がこれに対処する場合、おそらく経験と記憶を使って解決するでしょう。
まさにこれと同じことを Globex で前四半期に行いませんでしたか?それも似たような話でした。彼らは解約する恐れがあったため、CEO がフォーチュン 500 のロゴを保持したいことと、30 万ドルのアカウントをリスクを取る価値があるという理由で、誰かが 20% でサインオフしました。それはうまくいき、すぐにGlobexは更新されました。
人がそれをどう扱うか
担当者
担当者はどの程度の割引を承認する必要がありますか?
10％キャップはクリアです。イニテックが例外を正当化するかどうかは判断次第だ。
「フォーチュン 500 のアカウントをすべて保持します。更新割引でアカウントを失わないようにしてください。」
最後はCEOから総力を挙げて。
Globex の Zoom 通話
CC 「これほど大きなアカウントにはリスクを負う価値があり、通常はお金を支払います。」
スレッド # の更新
9月12日
MC マヤ・チェン 2:41 PM Globex は 10% の上限を超えなければ撤退すると脅しています。
3 件の返信
DR Devang Rao 2:44 PM 30 万ドルの口座、フォーチュン 500、そして副社長との素晴らしい関係。
PN Priya Nair 2:45 PM 保持する価値があります。彼らのために上限を超えてみましょう。私は財務チームを担当します。
SO サム・オカフォー 2:47 PM 20% で承認されました。
9月29日
SO サム・オカフォー 2:47 PM Globex が更新されました。
返信…
決定: 20% 割引して更新を終了します。
イニテックは 5 日以内に更新します。
意思決定を行うこの一連の推論は、エージェントが読める場所には書かれていません。エージェントは Globex の情報を見つけます。

ただし、Salesforce では、その番号が例外であること、誰が承認したのか、なぜ承認されたのか、現在の状況が同一であるかどうかは通知されません。
財務チームが 30 万ドルの口座にリスクを負う価値があると認めている古い緩いスレッド。
Zoom での通話では、営業のベテランがこの種の取引先が最終的に支払う金額について言及します。
Fortune-500 のロゴを保持することが重要であるという CEO からのメール。
これらは意思決定を行うために必要な重要な情報ですが、エージェントはこれらにアクセスできません。
したがって、エージェントは 2 つのうちの 1 つを実行します -
ポリシーの上限が 10% であることを通知する電子メールが送信され、アカウントが失われます。
それは人間にまでエスカレートし、その人間は会社が一度下した決定を再構築するために Slack の考古学に 24 時間を費やします。
更新は約1日停止します
電子メール: 「政策の上限は 10% です。これ以上のことはできません。」
イニテックのチャーン
いずれにせよ、組織は AI エージェントを導入しても利益を得ることはできませんでした。
これは普遍的な問題です。組織は毎年数十億ドルの損失を抱えています -
同じソリューションを再発明し、
以前に解決した問題に時間を浪費し、
新入社員の新人研修が信じられないほど遅い、
AI ネイティブの世界でコンプライアンスと監査のギャップに苦しんでいます。
これは、私たちが何が起こったのかを記録するのが非常に上手になっているためですが、なぜそれが起こったのかを体系的に放棄しているためです。これは、エージェントがここで必要としていたものの 1 つです。
コンテキスト グラフは、エージェントのメモリに理由を保存します。
Foundation Capital は、これを「1 兆ドル規模の AI の機会」と呼びました。これは、人々が戻るボタンに手を伸ばすような言葉です。しかし、とにかくそこに留まってください。この用語は新しいもので、少々大げさですが、現実的なものを指しています。エージェントを使用または構築すると、すぐにコンテキスト グラフを使用することになります。
AI エージェントにコンテキストを与える悪い方法は、単に与えることです。

すべてのデータ、レコード、ドキュメント、ルール、ポリシーがフラットなコンテキスト ウィンドウに表示されます。
請求書処理エージェントがいるとします。請求書、発注書、ベンダー記録、契約書、ポリシー文書をコンテキストにダンプします。次に、エージェントが失敗するのを観察します。なぜ？
エージェントが「請求書 #842 を支払ってもいいですか?」と自問すると、答えるには、ホップする必要があります。この請求書はどの PO を参照しているのでしょうか?その発注書にはまだ予算がありますか?配達物は受け取られましたか？ベンダーは支払いを保留していますか? $12,400 は承認基準を超えていますか?契約の支払い条件には何と記載されていますか?ポリシー文書は正確で最新ですか?この請求書支払いに関連する文書化されていないニュアンスが現在 Slack メッセージや Zoom 通話記録に残っていますか?
フラット検索では、この大量の情報を、切断されたチャンクの山として LLM モデルに渡そうとします。ここに請求書のテキスト、そこに注文書のテキスト、ランダムなドキュメントやスラック チャネルからの言語がすべて、平らなテキストの壁として表示されます。
Acme は常に四半期の終わりに支払います。追いかけないでください。
エージェントは、各ターンでこれらのデータ間の接続をすべて最初から再導出する必要があります。ビジネス データを緩いテキストにフラット化すると、必要な構造が正確に破壊されます。
そして、コンテキストが大きくなるにつれて、LLM はそのサイズに対処するのに苦労し、タスクに失敗し始めます。彼らは指示に従わず、ルールをランダムに削除し、遠く離れた 2 つの情報間の関係を誤解し、コンテキストの中間のデータを無視し、ルールと制約を順序を間違えて適用します。
Surge AI は、これを命令追従ベンチマークで文書化しています。最良のフロンティア モデルは、このような複雑なタスクの 41% 未満を解決します。
最初の例で見たように、AI エージェントは、人間が前例、経験、組織の記憶を使って毎日解決しているのと同じ曖昧さに遭遇します。でもそれはできません

これらをフラットなコンテキスト ウィンドウでエージェントに渡します。
部族の知識。 「私たちは物流会社に対して常に 5,000 ドルの入社費用を免除しますが、それは会社がスケジュールを先に延期した場合に限ります。」それはCRMには含まれていません。それは内部の会話を通じて受け継がれた部族の知識です。
過去の決断。 「私たちはアカウント X に対して支払いを分割払いにする取引を組み立てました。この同様のアカウント Y にも同様の提案をすべきです。」 Y の契約書がこのように作成された理由を伝えるために 2 つの契約をリンクするシステムはありません。
記録システム全体のコンテキスト。アカウント マネージャーは、製品ダッシュボードの使用量スライド、NetSuite の未払い請求書、コールド 1 行メールを確認します。 CRM でアカウントに「解約リスク」のフラグを立てます。この推論は彼らの頭の中で起こったものですが、CRM 記録は単に「解約リスク」を示しているだけです。
手動承認。 VP が Zoom 通話の割引を承認します。 Hubspot レコードには、変更された価格が表示されます。なぜこの決定がなされたのかは示されていない。
データ、意思決定、アクションの背後にある推論は、フラットなコンテキスト ウィンドウではキャプチャされません。
あなたが開発者である場合、この概念はさらに難しくなります。 2019 年にあのキューではなくこのキューを選択したのはなぜですか?再試行パスに sleep(200) があり、削除するとすべてが中断されるのはなぜですか?誰が書いたかは明らかですが、その情報は今では失われています。
アーキテクチャの決定記録を覚えていますか?これらはまさにこれを解決するために 2011 年に発明されました。しかし、ほとんどの ADR フォルダーは 3 つのエントリで消滅します。これは、フォルダーを書き込むのに手間がかかり、後から誰も読み取らないためです。
これは普遍的な問題です。企業は何が起こったのかを保存するのは得意ですが、なぜそれが起こったのかを保存するのは苦手です。これは、Why が構造化されておらず、システム全体に分散しており、保存したとしても誰も読まないためです。
コンテキストの腐敗と決定トレースの欠如という両方の問題は、コンテキスト グラフによって解決されます。
コンテキストグラフとは、

エージェントの記憶をグラフとして構造化する方法。ノードは情報の断片を保持し、エッジは情報間の関係を保持します。人間が閲覧するためではなく、エージェントが読み取るために最適化されています。
現在、ほとんどのエージェントのメモリはフラットです。 AI エージェントはデータを埋め込み、それらをチャンクに分割し、進行中のタスクに最も似ているいくつかのチャンクを返します。 LLM は、これらのチャンクがどのように相互に接続されているかを認識せずに、テキストの山を取得します。これは、現在 AI エージェントで使用されている標準メモリである Vector RAG です。
コンテキスト グラフはそれらの接続を維持します。 「ここに 5 つの同様の段落があります」の代わりに、「サービス A はサービス B に依存します」、「このリリースが原因でその機能停止が発生しました」、または「この請求書はそのポリシーに従います」などとすることができます。エッジには意味があり、モデルはエッジを横切ることができます。
類似性は関連性ではないため、これは重要です。 2 つのチャンクはタスクと単語を共有できますが、実際のタスクとは何の関係もありません。他の 2 つのチャンクはタスクと単語を共有していませんが、意味的にはタスクに関連しています。
コンテキストグラフを作成するにはどうすればよいですか?
コンテキスト グラフは、保存する内容と方法を変更することで、両方の失敗、コンテキストの腐敗、意思決定トレースの欠如を追跡します。
まずは「保管方法」から始めましょう。ビジネスをグラフとして保存します。請求書、発注書、アカウント、ベンダー、契約、ポリシー、承認者などのエンティティはすべて、コンテキスト グラフ内のノードです。それらの間の関係はエッジです。この請求書は、その注文書を参照します。この発注書はその予算を利用しています。この予算は、その人によって承認されました。このベンダー - によって管理される -> その契約。
これらの各リンクは、必要になるたびにモデルから離れてテキストのチャンクから再導出するのではなく、一度保存します。タスク/サブタスクごとに、エージェントは小さなサブグラフを取得して終了します。

残りの 1 万件のレコードは窓の外にあります。ウィンドウが小さく、適切な位置を保つため、コンテキストの腐敗はなくなります。
次に「保存するもの」です。また、各決定をコンテキスト グラフに保存できるようになりました。このコンテキストグラフの単位が判定トレースとなります。横ばいの記録は「Initech が 20% で更新」という結果で止まります。意思決定の痕跡は、その背後にあるストーリーを保持します。それを引き起こした問題、考慮された選択肢、拒否された選択肢が負けた理由、制約、例外、誰が決定したか、そしてその理由。
これは社員も頭の中に入れていることです。しかし、意思決定トレースのコンテキスト グラフを使用すると、エージェントはそれを読み取ることができます。
記録と時間のシステム全体で作成されたエンティティとその関係、さらに意思決定とエンティティやその他の意思決定との関係。 Foundation Capital の一言は「意思決定の記録システム」です。ほとんどのシステムには、現在の状態がすでに保存されています。コンテキスト グラフには、状態がどのようにしてそのようになったかが保存されます。
ユースケースに最適な意思決定トレースのスキーマを使用します。
さて後半では、古い決定をエージェントが頼れるものに変えています。
コンテキスト グラフを使用する場合の実装の詳細はほとんどありません -
コンテキスト グラフを備えたエージェントは低摩擦である必要があります。そうしないと、誰もそれを望まなくなります。

[切り捨てられた]

## Original Extract

We explain why agents need to capture decisions, how context graphs achieve this, and how agents can use these graphs.

Nanonets
Products
Pre-built Agents
AP Agent
Reconciliation Agent
Order Management Agent
Cash Application Agent
Supplier Onboarding Agent
Month End Closing Agent
RCM Agent
Data Extraction Agent For developers
Agent Platform
Agent Builder Build Custom Agents
Context Graph
Document Intelligence
Document Generation
Exception Management
Analytics
Workflow Orchestration
Integrations
Solutions
By Team
Finance
Accounting
Operations
Supply Chain
RCM
Software
By Industry
Manufacturing
CPG
Healthcare
Financial Services
Logistics & Supply Chain
Customers
Pricing
Resources
About Our mission, team, and investors.
Blog Guides, research, and product updates from the Nanonets team.
Docs Platform documentation, quickstarts, and integration guides.
API Reference Full API reference for building on Nanonets.
Research Model releases, benchmarks, and research notes from the Nanonets team.
Careers Join the team building AI that finishes the work.
Partnerships Join our global partner program and grow your IDP practice.
Book a demo
Get Started
Products
Pre-built Agents
AP Agent
Reconciliation Agent
Order Management Agent
Cash Application Agent
Supplier Onboarding Agent
Month End Closing Agent
RCM Agent
Data Extraction Agent
Agent Platform
Agent Builder
Context Graph
Document Intelligence
Document Generation
Exception Management
Analytics
Workflow Orchestration
Integrations
Solutions
By Team
Finance
Accounting
Operations
Supply Chain
RCM
Software
By Industry
Manufacturing
CPG
Healthcare
Financial Services
Logistics & Supply Chain
Customers
Pricing
Resources
About
Blog
Docs
API Reference
Research
Careers
Partnerships
Book a demo
Get Started
Context graphs: how AI agents remember why decisions were made
By
Karan
July 01, 2026
Link copied!
Copy failed!
Table of contents
Share
Link copied!
Copy failed!
In AI agents, a context graph is the part of agent memory that captures decisions.
This post explains why agents need to capture decisions, how context graphs achieve this, and how agents can use these graphs.
It's the last week of the quarter. A renewal agent is working a $480k account. The customer wants 20% off or they walk. The agent's instructions state >$100k accounts should not churn, but the agent's policy caps renewals at 10%. Now what?
If a human was handling this, they'll probably use experience and memory to resolve.
Didn't we do this exact thing with Globex last quarter? It was a similar story. They were threatening to churn, and someone signed off on 20% because the CEO wanted to retain Fortune 500 logos and the risk was worth taking on a $300k account. It worked and Globex renewed shortly after.
How a person handles it
The rep
How big a discount should the rep approve?
10% cap is clear. Whether Initech justifies an exception is a judgment call.
"Retain all our Fortune 500 accounts. Don't lose one over a renewal discount."
From the CEO, at the last all-hands.
Zoom call on Globex
CC "Accounts this big are worth the risk, and they usually pay."
Thread # renewals
September 12
MC Maya Chen 2:41 PM Globex is threatening to walk unless we beat the 10% cap.
3 replies
DR Devang Rao 2:44 PM $300k account, Fortune 500, and a great relationship with their VP.
PN Priya Nair 2:45 PM Worth retaining. Let's go above the cap for them. I'll handle finance team.
SO Sam Okafor 2:47 PM Approved at 20%.
September 29
SO Sam Okafor 2:47 PM Globex renewed.
Reply…
Decision: Give 20% discount and close renewal.
Initech renews within 5 days.
This reasoning chain that makes the decision is not written down anywhere your agent can read. The agent will find Globex's exception case in Salesforce, but Salesforce will not tell it that the number was an exception, who approved it, why it was approved, whether the current situation is identical or not.
Old slack threads where finance team admits a $300k account is worth the risk.
Zoom calls where sales veterans mention these kind of accounts pay eventually.
Emails from the CEO saying retaining Fortune-500 logos is critical.
These are critical pieces of information needed to make the decision, but the agent cannot access them.
So your agent does one of two things -
It sends an email informing the policy caps at 10%, and you lose the account.
It escalates to a human, who spends 24 hours doing Slack archaeology to reconstruct a decision the company already made once.
The renewal stalls for about a day
Email: “Policy caps at 10%. We won’t be able to do better.”
Initech churns
Either way, the organization failed to benefit by adopting an AI agent.
This is a universal problem. Organizations lose billions each year -
reinventing the same solutions,
wasting time on previously solved problems,
being incredibly slow at onboarding new employees,
struggling with compliance and audit gaps in an AI-native world.
This is because we have gotten extremely good at recording what happened, but we systematically throw away why it happened, which is the one thing your agent needed here.
Context graphs store the why in an agent's memory.
Foundation Capital called it " a trillion-dollar AI opportunity ", which is the kind of phrase that sends people reaching for the back button. But stick around anyway. The term is new and a little overloaded, but it points at something real. If you use or build agents, you'll end up using a context graph soon.
The bad way to give your AI agents context is to just give them all the data, records, documents, rules, policies in a flat context window.
Say you have an invoice processing agent. You dump the invoice, PO, vendor record, contract, policy document into the context. Then watch the agent fail. Why?
When the agent asks itself “can I pay invoice #842?” To answer, it has to hop - which PO does this invoice reference? does that PO still have budget? was the delivery received? is the vendor on payment hold? does $12,400 cross the approval threshold? what do the contract's payment terms say? is the policy doc accurate and up-to-date? are there undocumented nuances related to this invoice payment currently living on Slack messages and Zoom call transcripts?
Flat retrieval tries to hand this massive amount of information to the LLM model in a pile of disconnected chunks. Some invoice text here, some PO text there, some language from random docs and slack channels, all given as a flat wall of text.
Acme always pays end of quarter — don’t chase them.
The agent is forced to re-derive every one of those connections between these pieces of data from scratch on each turn. Flattening business data into loose text destroys exactly the structure it needs.
And as context becomes large, LLMs struggle to cope up with the size and start failing in their tasks. They fail to follow instructions, drop rules randomly, misunderstand the relation between two pieces of information far apart, ignore middle-of-context data, apply rules and constraints out of order.
Surge AI documents this in their instruction-following benchmark . The best frontier model solves <41% of such complex tasks.
Like we saw in our first example, AI Agents run into the same ambiguity humans resolve every day with precedents, experiences, organizational memory. But you can't give these things to an agent in a flat context window.
Tribal knowledge. "We always waive the $5k onboarding fee for logistics companies but only if they push back on the timeline first." That's not in the CRM. It's tribal knowledge passed down through internal conversations.
Past decisions. "We structured a deal for account X where they split payments into installments. We should offer this similar account Y the same." No system links the two deals to convey why Y's contract was drafted this way.
Context across systems of record. An account manager sees usage sliding in the product dashboard, an unpaid invoice in NetSuite, a cold one-line email. They flag the account as "churn risk" in the CRM. The reasoning happened in their head, but the CRM record just shows "churn risk".
Manual approvals. A VP approves a discount on a Zoom call. The Hubspot record shows the changed price. It doesn't show why this decision was made.
Reasoning behind data, decisions, actions isn't captured in a flat context window.
If you are a developer, this concept hits even harder. Why did we pick this queue over that one in 2019? Why is there a sleep(200) in the retry path that breaks everything when you remove it? It was obvious to whoever wrote it, but that information is gone now.
Remember Architecture Decision Records? They were invented back in 2011 to fix exactly this. But most ADR folders die at three entries, because writing them is friction and nobody reads them later.
This is a universal problem. Companies are good at storing what happened, and bad at storing why they happened. This is because the why is unstructured, spread across systems, and nobody reads it even if you store it.
Both problems, context rot and lack of decision traces, are solved by context graphs.
A context graph is a way of structuring an agent's memory as a graph, where nodes hold pieces of information and edges hold the relationships between them. It's optimized for the agent to read, not for a human to browse.
Most agent memory today is flat. AI agents embed your data, split them into chunks, and return the few chunks that look most similar to the ongoing task. The LLM gets a pile of text with no sense of how these chunks connect to one another. This is vector RAG, the standard memory used in AI agents today.
A context graph keeps those connections. Instead of "here are five similar paragraphs," it can say "Service A –depends on–> on Service B," "this release –caused–> that outage," or "this invoice –follows–> that policy." The edges carry meaning, and the model can traverse them.
This matters because similarity is not relevance. Two chunks can share words with your task and still have nothing to do with your actual task. Two other chunks can share no words with your task and still relate to your task semantically.
How to create a context graph?
A context graph goes after both failures, context rot and lack of decision traces, by changing what you store and how.
Start with "how you store". You store your business as a graph. Entities, for example, the invoice, the PO, the account, the vendor, the contract, the policy, the approver, are all nodes in the context graph. The relationships between them are edges. This invoice –references–> that PO. This PO –draws on–> that budget. This budget –approved by–> that person. This vendor –governed by–> that contract.
You store each of those links once, instead of leaving the model to re-derive them from chunks of text every time it needs them. For each task/subtask, the agent pulls a small subgraph and leaves the other ten thousand records out of the window. Context rot goes away, because the window stays small and on-point.
Now the "what you store". You now also store each decision in a context graph. The unit of this context graph will be a decision trace. A flat record stops at the outcome "Initech renewed at 20%". A decision trace keeps the story behind it. The problem that triggered it, the options weighed, why the rejected ones lost, the constraints, the exceptions, who decided, and the reasoning.
This is also what an employee keeps in their head. But with a context graph of decision traces, an agent can read it.
Entities and their relationships, plus decisions and their relationships to entities and other decisions, created across systems of record and time. Foundation Capital's one-liner for it is a " system of record for decisions ". Most of your systems already store the current state of things. A context graph stores how the state got that way.
You use a schema for the decision trace that is ideal for your use case.
Now the second half is turning old decisions into something the agent can lean on.
Few implementation details when using context graphs -
Agents with context graphs need to be low friction, otherwise no one will want to

[truncated]
