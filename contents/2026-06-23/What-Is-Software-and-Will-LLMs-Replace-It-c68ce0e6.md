---
source: "https://tomassetti.me/what-is-software-llms-interface-layer/"
hn_url: "https://news.ycombinator.com/item?id=48642685"
title: "What Is Software, and Will LLMs Replace It?"
article_title: "What Is Software, and Will LLMs Replace It? - Strumenta"
author: "ftomassetti"
captured_at: "2026-06-23T11:02:25Z"
capture_tool: "hn-digest"
hn_id: 48642685
score: 2
comments: 0
posted_at: "2026-06-23T10:02:49Z"
tags:
  - hacker-news
  - translated
---

# What Is Software, and Will LLMs Replace It?

- HN: [48642685](https://news.ycombinator.com/item?id=48642685)
- Source: [tomassetti.me](https://tomassetti.me/what-is-software-llms-interface-layer/)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T10:02:49Z

## Translation

タイトル: ソフトウェアとは何か、LLM はそれを置き換えるのか?
記事のタイトル: ソフトウェアとは何か、LLM はそれを置き換えるのか? - ストルメンタ
説明: ソフトウェアは LLM に置き換えられるのではなく、LLM が前面に押し出されており、決定論的なコア (スキーマ、制約、プロセス) はこれまでと同様に重要なままです。

記事本文:
-->
ソフトウェアとは何ですか?LLM はソフトウェアに取って代わるものですか? - ストルメンタ
コンテンツにスキップ
ソリューション
ソリューションの概要
移行サービス
移行ブループリント
すぐに使える RPG から Python への移行
DSL とエディター
ドメイン固有言語
パーサーとトランスパイラー
カスタムパーサー
移行サービス
移行ブループリント
すぐに使える RPG から Python への移行
DSL とエディター
ドメイン固有言語
パーサーとトランスパイラー
カスタムパーサー
ソフトウェアとは何ですか?LLM はソフトウェアに取って代わるものですか?
AI 、言語エンジニアリング、LLM、リフレクション、ソフトウェア開発
×
リンクトイン
フェイスブック
スレッド
電子メール
レディット
ブルースカイ
目次
私たちは皆、しばらくの間 LLM を使用しており、そのことに感銘を受けています。ある時点で、「これは本当ですか?」という質問をするのは自然なことです。これはソフトウェアに代わるものなのでしょうか?これからは、私たちはただコンピューターに話しかけて、欲しいものを説明し、それを表示させて、その間のすべてをスキップするだけなのでしょうか?
私はそうは思いませんが、信じたくなります。 「過去 5 年間の売上を表示」と入力すると、グラフが表示されます。スライドデッキを要求すると、スライドデッキが提供されます。終わり。もうSaaSを必要とする人はいないでしょうか?
しかし、これではソフトウェアが私たちにずっとしてくれてきたことが欠けています。私の意見では、これは 4 つの点で依然として関連性があります。
データは整理され正規化されています。
パターンを理解するのに役立つ方法で視覚化されたもの。
プロセスは段階的にガイドされます。仕事を正しく行う方法について長年蓄積されたノウハウをキャプチャし、実行可能にしました。
はい、LLM は、インターフェイスが私たちが想像していたよりもはるかに柔軟になり得るという、本当に役立つことを私たちに示してくれています。しかし、それはソフトウェアが会話の中に消えていくのと同じではないと思います。その理由を見てみましょう。
ソフトウェアは実際に何をするのでしょうか?
これを具体的なもの、つまり CRM に基づいて考えてみましょう。はい、CRM はまったく退屈ですが、どの B2B 企業も

一つのものであり、私たちはそれらをよく知っています。そして、正直に言うと、私たちが使用するソフトウェアのほとんどは魅力的である必要はありません。それは私たちにとって退屈な仕事をしなければならないだけです。
データを構造化され、クエリ可能な正規化された形式に整理します。
商談はテキストのかたまりではありません。これは、電話番号と電子メールによる連絡先、マーケティング帰属をフィードするリードソースフィールド、および過去の契約に遡るチェーンを含む企業にリンクされたレコードです。この構造により、「過去 6 か月間で紹介から生じた商談が 50,000 ユーロを超えて成約したのはどれですか?」と尋ねることができます。もっともらしい段落ではなく、ミリ秒単位で信頼できる答えが得られます。
はい、自由形式のメモを書くだけの方が便利です。しかし、組織化されたデータの規律がなければ、ほとんどの分析を行う可能性が失われます。また、同僚が情報を標準形式に変換できるのは嬉しいことですよね?
これにより、一貫性と整合性が強化されます。
商談が属する会社を最初に作成しない限り、商談を作成することはできません。未処理の契約がまだある会社を削除することはできません。システムはユーザーを停止するか、定義された順序でリンクされたすべてのレコードの削除をカスケードします。これらは、6 か月後にデータがゴミにならないようにするためのルールです。より良い未来のために目先の満足を放棄することです。
視覚化とフィルタリングが可能になります。
純粋にテキスト形式のインターフェイスは、すべてに対してうまく機能するわけではありません。データを確認し、グラフ内の外れ値に気づき、一目でパターンを特定する必要があります。そのためには、視覚化とフィルタリングが必要です。
LLM に売上グラフを依頼すると、喜んでグラフを作成してくれます。しかし、その数字はどこから来たのでしょうか?来月も同じ答えが得られるでしょうか?
会社の用語、承認の順序、および

正しく理解するのに何年もかかったドメインのサヴォアフェールの小さな断片。たとえば、「契約が締結されるまで、商談を落札としてマークすることはできません。」それは摩擦ではありません。それがルールです。あなたはそれらが好きではないかもしれませんが、同僚からいくつかのタスクを引き継ぐ必要がある場合は、間違いなく同僚にフォローしてもらいたいと思うでしょう。
歴史的な余談: SQL と平易な英語の夢
このすべての背後には古い夢があります。 1970 年代、IBM 研究者の Donald Chamberlin 氏と Raymond Boyce 氏が SQL を設計したとき、その提案の一部は、プログラマーでなくても直接使用できるほど単純であるということでした。当時の講演や論文で繰り返された考え方は、開発者がユーザーとデータの間に立つ必要はなく、管理者は必要なことを平易な英語に近い言葉で入力するだけで答えが得られるというものでした。
よく引用されるイラストは、家庭料理人が手元にある材料を入力して、作れるレシピのリストを取得するようなものでした。それは邪魔をする「技術層」の終わりのように聞こえました。 50 年経った今でも、私たちはその約束を実現しようと努力しており、難しい部分は決してインターフェイスだけではなかったことを発見しています。
LLM だけではビジネス ソフトウェアを置き換えることができないのはなぜですか?
チャット インターフェイスを取り除いて、LLM はデータをどこに保存するのかを考えてみましょう。そうではありません。モデルにはスキーマ、外部キー、トランザクション、制約はありません。
多くのエージェント設定では、「メモリ」は、エージェントがセッション間で読み書きするファイルまたはメモのコレクションです。これは便利ですが、データベースではありません。メモは、商談に会社があることを強制するものではありません。メモは削除をカスケードしません。メモは、今日取得する番号が明日取得する番号と同じであることを保証しません。
そして答えが見つからない場合は

それらのメモから確実に掘り出されたでしょうか？まあ、それでも LLM は答えます。そして、その答えは納得できるでしょう。ただし、これらの数字を IRS に提出してみて、どうなったかを教えてください。
エア・カナダのチャットボット事件は何を示していますか?
これは仮定のリスクではありません。 2024 年 2 月、カナダの民事解決裁判所は、モファット対エア・カナダ (CRT 149) でエア・カナダに対する判決を下しました。航空会社のチャットボットが顧客に対し、航空券を正規料金で予約し、その後忌引割引を適用できると伝えた後です。これは、同社が公表したルールと矛盾していました。エア・カナダの弁護は、本質的にはチャットボットが自らの発言に責任を負っているというものだった。法廷はその主張を受け入れず、航空会社に補償金、利息、手数料の支払いを命じた。
これが問題のすべてを一文で表したものです。決定論的なポリシー テーブルでは、正確に 1 つの答えが返されることになります。そのテーブルがなかったチャット インターフェイスが即興でチャット インターフェイスを作成したため、会社はその即興に夢中になりました。
MCP とソフトウェアへの橋渡し
それでは、モデルにはない構造が必要な場合はどうすればよいでしょうか?データベースに重みを詰め込もうとする必要はありません。既存のデータベースへのブリッジを構築します。
当初はツール呼び出しを通じて行われていましたが、現在ではモデル コンテキスト プロトコル、つまり「AI アプリケーション用 USB-C」が存在します。この比喩は実際に機能しているということです。USB-C はハード ドライブ、モニター、キーボードに取って代わるのではなく、それらすべてに 1 つの標準プラグを提供しました。 LLM はその機能を補完する他のソフトウェアと対話する必要があるため、MCP は実際の構造化された状態を保持するモデルとシステムに対して同じことを行います。
Salesforce は CRM をチャットボットとして再構築しませんでした。 MCP サポートが Agentforce に組み込まれたため、モデルはすでにスキーマと CA を適用している CRM に到達できます。

スケードルール。アトラシアンは、Jira と Confluence を会話の中に溶解しませんでした。リモート MCP サーバーを構築したので、モデルは、常に存在していた場所に正確に存在し、構造化され、制約されたままのデータに対してクエリを実行し、操作できるようになります。
動きの方向に注目してください。これらの企業は自社製品をチャットボットに統合するつもりはありません。これらは、既にデータを保存しているシステムにアクセスし、権限を強制し、どのアクションが有効かを定義するための制御された方法をモデルに提供します。ソフトウェアが LLM に吸収されるわけではありません。それは、LLM がソフトウェアに組み込まれていることです。
問題は、LLM だけを使用することから始めたということです。その後、LLM はソフトウェアを使用し始めました。次に起こることは、最終的にはソフトウェアが再び原動力となり、比率が逆転すると思います。
最終的には、少しの決定論的なソフトウェアを備えた LLM アプリケーションは作成できなくなります。 LLM をコンポーネントの 1 つとして使用する決定論的なソフトウェアを作成します。
ソフトウェアが LLM に置き換えられることはありません。彼らが先頭に立って戦うことになるだろう。正規化され、よく整理されたデータが依然として必要です。まだ制約が必要です。実行するには検証済みのプロセスがまだ必要です。私たちは依然としてソフトウェアによって動かされる規律を必要としています。 LLM は、私たちをそれらから解放してくれるかのような錯覚を与えます。変わるのは、そのすべてに到達するために通過するドアです。
本当に新しいもの (そしてその欠点)
わかりやすい言葉で質問を入力することは、クエリ構文を学習したり、適切なフィルターを見つけるために 6 つのメニューを探したりするよりも簡単です。これにより、レポートを作成するためだけにトレーニング セッションが必要となる人々にとって、参入障壁が低くなります。これはアクセシビリティにとって真の勝利です。
技術者として、それは憤りを感じます。 「機械」と通信する優れた能力のおかげで、私はある種の即席の貴族の一部になっているように感じました。でもそれはvです

貴重な。
問題は、「許す」ということは双方向に影響を与えるということです。曖昧な入力を受け入れ、常に自信を持って聞こえる何かで応答するシステムは、遅かれ早かれ、本当の答えのない質問にも自信を持って応答するでしょう。それがエア・カナダに損害賠償を伴う法廷で起こったことだ。
インターフェースを快適なものにする柔軟性は、その背後に構造化されたグラウンドトゥルースが存在しない場合、インターフェースを危険なものにしてしまいます。したがって、私たちは完全な自由を手に入れることはできません。決定論的ソフトウェアとも呼ばれるガードレールが必要です。実際のビジネス システムでは、ガードレールがおそらくシステムの最大の部分となり、LLM は明確に定義された領域に制限されます。 ChatGPT を使用して休暇を計画することもできます。しかし、実際の作業を行うには、より制約されたものが必要です。
すべてがうまくいくので、LLM との仕事はとても気持ちいいと思います。従来のソフトウェアは、規定された方法で物事を行うことを強制するのに対し、それらはまったく規律を強制しません。
それは仕事のように感じられますが、良いことです。ソフトウェアは厳格な親のようなものですが、LLM を使用すると、宿題をせずにジェラートを好きなだけ食べさせてくれるおばあちゃんを訪ねているような気分になります。私たちは皆、おばあちゃんたちと過ごした思い出が大好きですが、おそらくそれが私たちの生産性のピークではなかったのでしょう。
有用なシステムを作成するという仕事の本質的な部分は消えません。誰かが依然としてドメイン分析を行う必要があります。つまり、実際の商談が何であるか、いつ契約にできるか、何がチャーンとしてカウントされるか、データの形式は何か、絶対に破ってはいけないルールはどれかを定義します。その作業はいつも大変でしたし、今でもそうです。
新しいのは、システム自体の内部に線を引く必要があることです。つまり、どの部分が決定性、厳格性、証明可能性を維持する必要があり、どの部分を誰かに引き渡すことができるかということです。

より柔軟になります。
データベース、スキーマ、オープンな契約を持つ会社の削除を阻止するカスケード ルール。これは一貫性を保証するものであるため、堅固であり続ける必要がある部分です。 LLM はその上にあり、人々が最初に内部構造を学ばなくても LLM に到達できるようにします。
これは本当の変化であり、有益なものです。以前は専用の自動化を構築したり、複数ステップの手動プロセスをガイドする必要があったタスクは、現在では、決定論的コアと正しく通信する方法を知っている LLM によって処理できることが多くなりました。
しかし、それが依存しているものに注目してください。LLM が置き換えるものではなく、LLM が通信する決定論的なコアです。境界を正しく設定すると、LLM はインターフェイスの真に役立つ部分になります。誤解すると、はるかに少ないトレーニングで、データの正直さを保つためのルールを回避するための自信に満ちた方法を、はるかに多くの人々に提供したことになります。
AI 、言語エンジニアリング、LLM、リフレクション、ソフトウェア開発
AI コーディングは信頼性の低いコンパイラーを使用しているように感じます
Visual Basic 6 と VBA パーサーの使用方法
検索
購読してください
ニュースレター
移行サービス
移行ブループリント
すぐに使える RPG から Python への移行
DSL とエディター
ドメイン固有言語
パーサーとトランスパイラー
カスタムパーサー
移行サービス
移行ブループリント
すぐに使える RPG

[切り捨てられた]

## Original Extract

Software isn't being replaced by LLMs, it's being fronted by them, with the deterministic core (schemas, constraints, processes) staying as essential as ever.

-->
What Is Software, and Will LLMs Replace It? - Strumenta
Skip to content
Solutions
Solutions overview
Migration Services
Migration Blueprint
Ready-to-go RPG to Python Migration
DSLs and Editors
Domain specific languages
Parsers and Transpilers
Parsers custom
Migration Services
Migration Blueprint
Ready-to-go RPG to Python Migration
DSLs and Editors
Domain specific languages
Parsers and Transpilers
Parsers custom
What Is Software, and Will LLMs Replace It?
AI , Language Engineering , LLMs , Reflections , Software Development
X
LinkedIn
Facebook
Threads
Email
Reddit
BlueSky
Table of contents
We’ve all been using LLMs for a while now, and we’ve all been impressed by them. At some point it is natural to ask the question: is this it? Is this what is going to replace software? Are we just going to talk to computers from now on, describe what we want, have it appear, and skip everything in between?
I don’t think so, but it’s tempting to believe it. Type “show me sales for the last five years” and you get a chart. Ask for the slide deck and you get the slide deck. Done. Who needs SaaS anymore?
But that misses what software has been doing for us all along. In my opinion, it remains relevant in four ways:
Data organized and normalized.
Things visualized in ways that help us see patterns.
Processes guided step by step. Years of accumulated know-how about how to do a job right, captured and made executable.
So yes, LLMs are showing us something genuinely useful: that interfaces can be far more flexible than we assumed. But I don’t think that’s the same as software disappearing into conversation. Let’s look at why.
What does software actually do?
Let’s ground this in something concrete: a CRM. Yes, CRMs are boring as hell, but every B2B company has one and we are familiar with them. And, let’s face it, most of the software we use does not have to be glamorous. It just has to do some boring job for us.
It organizes data into a structured, queryable, normalized form.
An opportunity is not a blob of text. It is a record linked to a company, which has contacts with phone numbers and emails, a lead-source field that feeds your marketing attribution, and a chain back to past contracts. That structure is what lets you ask “which opportunities came from referrals and closed above €50k in the last six months?” and get an answer you can trust in milliseconds, not a paragraph that sounds plausible.
Yes, just writing free-form notes would be more convenient. But without the discipline of organized data, we would lose the possibility of doing most analysis. Also, we are glad our colleagues have to put information into a standard format, right?
It enforces consistency and integrity.
You cannot create an opportunity without first creating the company it belongs to. You cannot delete a company that still has open contracts. The system stops you, or cascades the deletion through every linked record in a defined order. These are the rules that keep your data from turning into garbage six months from now. You know, giving up immediate gratification for a better future.
It enables visualization and filtering.
Purely textual interfaces do not work very well for everything. We need to see data, notice outliers in a graph, and spot patterns at a glance. For that we need visualizations, and also filtering.
Ask an LLM for a sales chart and it’ll happily draw you one. But where did the numbers come from, and will it give you the same answer next month?
It encodes your company’s terminology, its sequence of approvals, and the little pieces of domain savoir-faire that took years to get right. “You can’t mark an opportunity as won until the contract is attached,” for example. That’s not friction. That’s rules. You may not like them, but you would certainly like your colleagues to follow them if you need to take over some tasks from them.
A historical side note: SQL and the dream of plain English
There is an old dream behind all this. Back in the 1970s, when IBM researchers Donald Chamberlin and Raymond Boyce designed SQL, part of the pitch was that it would be simple enough for non-programmers to use directly. The idea, repeated in talks and papers of the era, was that you wouldn’t need a developer standing between you and your data: a manager could just type out what they wanted in something close to plain English and get an answer back.
The often-cited illustration was something like a home cook typing in the ingredients they had on hand and getting back a list of recipes they could make. It sounded like the end of “the technical layer” getting in the way. Fifty years later, we are still trying to make that promise true, and still discovering that the hard part was never only the interface.
Why can’t an LLM alone replace business software?
Strip away the chat interface and ask: where does an LLM keep its data? It doesn’t. A model has no schema, no foreign keys, no transactions, no constraints.
In many agent setups, the “memory” is a collection of files or notes that the agent reads and writes between sessions. That can be useful, but it is not a database. Notes don’t enforce that an opportunity has a company. Notes don’t cascade deletes. Notes don’t guarantee that the number you get today is the same number you’ll get tomorrow.
And if an answer cannot be determined reliably from those notes? Well, the LLM will answer nevertheless. And the answer will be plausible. But try submitting those numbers to the IRS and let me know how it goes.
What does the Air Canada chatbot case show?
This isn’t a hypothetical risk. In February 2024, Canada’s Civil Resolution Tribunal ruled against Air Canada in Moffatt v. Air Canada (CRT 149), after the airline’s chatbot told a customer he could book a flight at full price and apply for a bereavement discount afterward. That contradicted the company’s own published rules. Air Canada’s defense was, essentially, that the chatbot was responsible for its own words. The tribunal did not accept that argument, and ordered the airline to pay compensation, interest, and fees.
That’s the whole problem in one sentence: a deterministic policy table would have returned exactly one answer. A chat interface, lacking that table, improvised one, and the company was on the hook for the improvisation.
MCP and the bridge back to software
So what do you do when your model needs structure it doesn’t have? You don’t try to cram a database into its weights. You build a bridge to the database that already exists.
Initially it was done through tool calling and now there is the Model Context Protocol , the “USB-C for AI applications”. The metaphor is doing real work: USB-C didn’t replace your hard drive, your monitor, or your keyboard—it gave them all one standard plug. MCP does the same for models and the systems that hold actual structured state, because LLMs need to interact with other software that complements their capabilities.
Salesforce didn’t rebuild its CRM as a chatbot. It built MCP support into Agentforce, so the model can reach into the CRM that already enforces the schema and the cascade rules. Atlassian didn’t dissolve Jira and Confluence into conversation. It built a remote MCP server so a model can query and act on data that still lives, structured and constrained, exactly where it always lived.
Notice the direction of motion. These companies are not dissolving their products into chatbots. They are giving models controlled ways to reach into the systems that already store the data, enforce permissions, and define what actions are valid. That’s not software being absorbed by LLMs. That’s LLMs being absorbed into software.
The thing is that we started by using just LLMs. Then LLMs started to use software. What I think will happen next is that software will eventually be the driving force again, inverting the proportion.
In the end, we will not have an LLM application with a little bit of deterministic software. We will have deterministic software that uses LLMs as one of its components.
Software is not going to be replaced by LLMs. It’s going to be fronted by them. We still need normalized, well organized data. We still need constraints. We still need validated processes to be executed. We still need the discipline that is nudged by software. LLMs give us the illusion of freeing us from those. What changes is the door you walk through to get to all of that.
What’s genuinely new (and the catch)
Typing a question in plain language is more forgiving than learning a query syntax or hunting through six menus for the right filter. It lowers the barrier to entry for people who would otherwise need a training session just to pull a report. That’s a genuine win for accessibility.
As a technician, I can resent that. I felt part of some sort of improvised nobility because of my superior ability to communicate with “the machine”. But it is valuable.
The catch is that “forgiving” cuts both ways. A system that accepts vague input and always responds with something confident-sounding will, sooner or later, respond confidently to a question it has no real answer for. That’s what happened to Air Canada, in front of a tribunal, with a damages award attached.
The same flexibility that makes the interface pleasant is what makes it dangerous when there is no structured ground truth behind it. So we cannot have absolute freedom. We need guardrails, also known as deterministic software. In real business systems, the guardrails will probably be the biggest part of the system, with the LLM restricted to a well-defined area. You can still use ChatGPT to plan your holiday. But to do real work, you need something more constrained.
I think working with LLMs feels great because everything goes. They enforce exactly zero discipline, while traditional software forces us to do things in a prescribed way.
And while that feels like work, it is a good thing . Software is like a strict parent, while using LLMs feels like visiting the grandma who lets you eat as much gelato as you want while doing zero homework. We all love the memories we spent with our grandmas, but perhaps that was not the peak of our productivity.
The real part of the job in creating useful systems does not disappear. Someone still has to do the domain analysis: define what an opportunity actually is, when it is allowed to become a contract, what counts as churn, what format the data takes, and which rules can never be broken. That work was always the hard part, and it still is.
What’s new is that there is now a line to draw inside the system itself: which parts should stay deterministic, rigid, and provable, and which parts can be handed over to something more flexible.
The database, the schema, the cascade rule that stops you from deleting a company with open contracts: that’s the part that has to stay solid, because it is what guarantees consistency. The LLM sits on top of that, helping people reach into it without having to learn its internals first.
That is a real shift, and a useful one. Tasks that used to require building a dedicated automation, or walking someone through a multi-step manual process, can now often be handled by an LLM that knows how to talk to the deterministic core correctly.
But notice what that depends on: a deterministic core that the LLM talks to, not one it replaces. Get the boundary right, and the LLM becomes a genuinely useful part of the interface. Get it wrong, and you have given a much larger number of people, with much less training, a confident-sounding way to bypass the rules that were keeping your data honest.
AI , Language Engineering , LLMs , Reflections , Software Development
AI Coding Feels Like Using an Unreliable Compiler
How to use the Visual Basic 6 and VBA Parser
Search
Subscribe to our
newsletter
Migration Services
Migration Blueprint
Ready-to-go RPG to Python Migration
DSLs and Editors
Domain specific languages
Parsers and Transpilers
Parsers custom
Migration Services
Migration Blueprint
Ready-to-go RPG to

[truncated]
