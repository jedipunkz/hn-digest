---
source: "https://cameronmpalmer.medium.com/should-you-even-use-an-llm-b4f3b7914f4d"
hn_url: "https://news.ycombinator.com/item?id=49008624"
title: "Six questions before you add an LLM"
article_title: "Medium"
author: "cameronmpalmer"
captured_at: "2026-07-22T16:16:17Z"
capture_tool: "hn-digest"
hn_id: 49008624
score: 4
comments: 1
posted_at: "2026-07-22T15:43:45Z"
tags:
  - hacker-news
  - translated
---

# Six questions before you add an LLM

- HN: [49008624](https://news.ycombinator.com/item?id=49008624)
- Source: [cameronmpalmer.medium.com](https://cameronmpalmer.medium.com/should-you-even-use-an-llm-b4f3b7914f4d)
- Score: 4
- Comments: 1
- Posted: 2026-07-22T15:43:45Z

## Translation

タイトル: LLM を追加する前の 6 つの質問
記事タイトル: 中
説明: LLM を使用する必要がありますか?やりすぎたエージェント 私は問題を抱えていました。AI 導入コンサルタントとして、リードを発見して評価し、リードを自分のプロジェクトに挿入するための自動化されたパイプラインがありませんでした。

記事本文:
中 LLM を使用する必要がありますか?決定のための 6 つの質問の枠組み… |キャメロン・パーマー著 | 2026 年 7 月 |中 サイトマップ アプリで開く サインアップ
「AIはどこで使えるの？」逆向きです
LLM の取引決定主義による柔軟性の実現
LLM を追加する前の 6 つの質問
フレームワークを私の探査システムに適用する
あいまいさが存在する場合は LLM を使用する
私には問題がありました。AI 導入コンサルタントとして、リードを発見して評価し、セルフホスト型の Twenty CRM インスタンスに挿入するための自動パイプラインがありませんでした。私は、LLM を使用してこの問題全体を解決できると考えました。
私の見込み客発見システムの最初のバージョンでは、1 人の LLM エージェントに、Web 検索、重複排除、検証、データベース挿入などの完全なスカウト パイプラインを提供しました。見つかった見込み顧客の約 10 ～ 20% は、無関係であるか、重複しているか、CRM に不適切に挿入されていました。さらに悪いことに、エージェントの各インスタンスが異なるアプローチを採用し、独自のプロセスやツールの速度向上に遭遇したため、システムのデバッグが非常に困難でした。結局、CRM 内の 800 を超える見込み客のリスト全体を手動で検索し、重複がないか確認し、見込み客が関連性があるかどうかを検証する必要がありました。明らかに、このソリューションはそのままでは機能しません。
「AIはどこで使えるの？」逆向きです
経営陣は可能な限りあらゆる形で AI を導入することに執着しているようです。 AI 導入の進捗状況は、トークンの使用量や書かれたコード行などの任意の指標を使用して測定されており、AI が有益な方法で使用されているか有害な方法で使用されているかを測定するものではない、とよく聞きます。大規模言語モデル (LLM) は、基本的には単なるツールです。 CEO が AI をハンマーであると考え、そのハンマーをあらゆることに使用したいと考えている場合、すべてが釘のように見えてしまいます。これは、AI ソリューション アーキテクチャに対する間違ったアプローチです。
g

「当社は AI ファーストの企業です」または「当社の製品は AI を使用しています」と言うべきではありません。目標は、適切なツールを使用して最も緊急で関連性のある問題を解決することです。 LLM はそれらのツールの 1 つです。最初に尋ねられる質問は、「AI をどこに適用できるか?」ではなく、「どのような問題を解決するのか?」であるべきです。解決すべき問題が決まれば、初めて解決策を特定することができます。ソリューションにはどのような機能が必要ですか?現在のソリューションに問題がある場合、どこで問題が発生しますか?そして、LLM をそのソリューションに統合することは実際に必要なのでしょうか?
LLM の取引決定主義による柔軟性の実現
ソフトウェア開発ツールの各ツールには長所と短所があり、LLM も例外ではありません。問題は、LLM が有用かどうかではなく、LLM がもたらすトレードオフに見合う機能があるかどうかです。 LLM は、決定性を犠牲にすることで柔軟性を提供します。これらは、自然言語の解釈、豊富または多様な情報の合成、プロセスのルールを開始前に指定できない段階的な推論が必要な作業に役立ちます。
これらの柔軟性の向上により、再現性と決定性が失われます。 LLM はさまざまな出力を生成します。つまり、同じモデルに同じプロンプトを 2 回与えると、明らかに異なる結果が生成されます。このため、成功基準は主観的なものが多いため、テストと障害分析が従来のコードよりもはるかに困難になります。さらに、LLM によって実行されるプロセスは通常、プレーン コードを使用して実行される同じワークフローよりもはるかに時間がかかり、コストが高くなります。これらの制限は、AI が役に立つと盲目的に考える前に、ソリューションにおける AI の必要性を徹底的に評価する必要があることを意味します。
LLM を追加する前の 6 つの質問
LLM が役立つかどうかを評価する場合、

具体的な質問を使用して決定を検討するのに役立ちます。
ソリューションを実装するワークフローを事前に完全に指定できますか? 「はい」の場合、LLM は必要ない可能性があります。たとえば、典型的な CI/CD パイプラインを考えてみましょう。コードがリモートにプッシュされ、lint を生成し、コード テストを実行し、開発環境にデプロイするワークフローがトリガーされます。このワークフローは開始前に完全に指定でき、毎回同じように実行されます。
同一の入力は同一の出力を生成しなければなりませんか? LLM は非決定的です。ワークフローが同一の入力に対して同じ出力をもたらす必要がある場合は、LLM を除外することを検討することをお勧めします。たとえば、支払システムは同じ請求書、税務管轄区域、および再試行リクエストを 2 回受け取ります。同じ合計を計算し、顧客に二度請求することを避ける必要があります。
解決策には、実行中に生じる曖昧さを解釈する必要がありますか? LLM は、実行中のあいまいさに優れています。ファイルが指定されたフォルダーにない場合、次に解決策はどこを調べればよいでしょうか?決定論的ルールは、特定のレベルまでのあいまいさを処理できますが、従来のルールの実装では、ソリューションの実行が幸せな道から外れる可能性があるすべての方法を事前に検討する必要があります。これは、「実行されていないソリューションに曖昧さがあるかどうか」とは異なることに注意してください。ソリューションに実行開始前に解決できるあいまいさがある場合は、あいまいさを解決してからプレーン コードを使用します。
ソリューションの結果を安価かつ正確に検証できるか?コードを明示的な基準に対して安価にテストできるため、誰もがコードをバイブレーションできるようになりました。必要な機能が存在し、テストに合格し、サーバーが安定している場合、コードはおそらく実行可能であると考えられます。一方で、医学的診断が安全かどうかを検証する

患者に推奨するには時間がかかり、専門知識が必要です。一般に、検証に大量の作業や人間の特別な専門知識が必要な出力 (この法的アドバイスは正しいか?)、または検証に期限がある (このビジネス戦略は成功につながるか?) 出力は、LLM の使用には適さないことがよくあります。
出力が間違っている場合はどうなりますか? 「もし」ではなく「いつ」。システムには障害が発生するものであり、私たちはそれに備えなければなりません。ソリューションの出力が間違っている場合、その影響はどれほど深刻ですか?その影響は軽減できるでしょうか?誤って顧客に 15 ドルの割引を与えるソリューションは、顧客に 5000 ドルの割引を与えるソリューションよりもリスクがはるかに低くなります。結果の発生点からソリューションを遠ざけることで、誤った出力を軽減できます。顧客に直接割引を与えるボットの代わりに、提供すべき割引について人間のサポート エージェントにアドバイスするチャット ボットを検討してください。
LLM ベースのソリューションは、より単純な代替ソリューションよりも大幅にパフォーマンスが優れていますか?ここでは KISS 手法をお勧めします (Keep It Simple, Stupid)。LLM ベースのソリューションの品質 (エラー率、人的介入の数、コストなど) の 95% を達成できるコードベースの代替手段がある場合は、コードベースのソリューションを使用してください。
フレームワークを私の探査システムに適用する
私のプロスペクティング システムでは、当初、LLM にプロスペクティング パイプライン全体 (検索、重複排除、検証、DB 挿入) を実行させていました。このモデルは、検索クエリを生成し、MCP Web 検索ツールでクエリを実行し、返された見込み客と現在 DB にある見込み客との重複を除去し、理想的な顧客の説明と照らし合わせて見込み客を検証し、生き残ったリストを CRM に書き込みます。段差が多かったのでよく混乱しました。
を使ってみましょう

彼の例を参考にして、前のセクションで概説した質問をじっくり考えてみましょう。
ワークフローを事前に指定することはできますか?はい、上で完全なワークフローの概要を説明しました。
同一の入力は同一の出力をもたらす必要がありますか?はい。プロスペクト ファイルが 1 回の実行で有効であれば、次の実行でも有効であるはずです。
解決策には、実行中に生じる曖昧さを解釈する必要がありますか?はい。ただし、すべてのステップではそうではありません。クエリを生成するには、CRM にすでに存在する見込み客のタイプを解釈し、ギャップを埋めるクエリを作成する必要があります。ただし、Decodo 検索によるクエリの実行、既存のエントリからの結果の重複排除、CRM データベースへの挿入など、他の手順にはあいまいさはありません。
ソリューションの結果を安価かつ正確に検証できるか?はい - 名前、会社、役職を 10 秒以内に読み取ることで、データベース内の見込み顧客が私の理想的な顧客の説明に一致するかどうかを確認できます。
出力が間違っている場合はどうなりますか?私の場合、誰も亡くなったり、多額のお金を失ったりすることはありません。おそらく、誤ってプロスペクトにメッセージを 2 回送信したか、無効なプロスペクトにメッセージを送信してしまった可能性がありますが、これは修正できます。
LLM ベースのソリューションは、より単純な代替ソリューションよりも大幅にパフォーマンスが優れていますか?はい。別の解決策は、自分でクエリを考え、それを重複排除検索スクリプトに渡し、理想的な顧客の説明と照らし合わせて見込み客を手動で検証し、挿入に最も有望なものをマークすることです。それか、ある種の決定論的なクエリ生成のギャップ埋めシステムとハックな正規表現キーワード マッチング検証システムを実装するかのどちらかですが、どちらもあまり効果的でなく、時間効率も良くありません。
あいまいさが存在する場合は LLM を使用する
LLM を使用するかどうかの決定は簡単ではありません。柔軟性と非効率性のトレードオフを行う必要があります。

終端主義、適応性と精度、コストとパフォーマンス。私の比較的単純な見込み客システムでも、検証の質問に対する答えが必ずしも明確な「はい」か「いいえ」であるとは限りません。これはよくあることだと私は感じています。通常、LLM を含むソリューションには、決定論的検証コードとプロセス自動化コードも含まれています。最適なソリューションは、LLM の長所と決定論的コードの長所を組み合わせたものであり、どちらか一方を選択するというものではありません。
多くの問題は決定論的な解決策を使用して解決でき、大規模な言語モデルを追加するメリットはありませんが、私たちは以前はコンピューターで解決できなかった多くの問題が解決できる時代に生きています。これらのソリューションでは、エンジニアが毎日数百ドルをトークンに費やす必要はありません。チームが当面の問題を評価し、適切な仕事に適切なツールを使用することを決定する必要があります。
LLM を決定論的コードに置き換えることでシステムはどこに改善されましたか?また、それによってどのような新たなトレードオフが生じましたか?

## Original Extract

Should you even use an LLM? The Agent That Did Too Much I had a problem: as an AI implementation consultant, I had no automated pipeline to discover and qualify leads and insert them into my …

Medium Should you even use an LLM?. A six-question framework for deciding… | by Cameron Palmer | Jul, 2026 | Medium Sitemap Open in app Sign up
“Where Can We Use AI?” Is Backwards
LLMs Trade Determinism for Flexibility
Six Questions Before You Add an LLM
Applying the Framework to My Prospecting System
Use the LLM Where Ambiguity Lives
I had a problem: as an AI implementation consultant, I had no automated pipeline to discover and qualify leads and insert them into my self-hosted Twenty CRM instance. I thought I could solve this whole problem using LLMs.
The first version of my prospect discovery system gave one LLM agent the full scouting pipeline: web search, deduplication, validation, and database insertion. About 10–20% of the prospects found were irrelevant, duplicated, or improperly inserted into the CRM. To make matters worse, the system was extremely difficult to debug, as each instance of the agent took a different approach and ran into its own unique process and tooling speed bumps. I ended up having to search through the entire list of 800+ prospects in the CRM manually myself, checking for duplicates and validating if the prospect was relevant. Clearly this solution wasn’t going to work as-is.
“Where Can We Use AI?” Is Backwards
Executives seem obsessed with implementing AI in any form possible. I often hear that AI adoption progress is measured using arbitrary metrics such as token usage or lines of code written, which don’t measure if AI is being used in a helpful or a harmful way. Large language models (LLMs) are fundamentally just a tool. If your CEO thinks AI is a hammer, and she wants to use that hammer for everything, then everything ends up looking like a nail. This is the wrong approach to AI solution architecture.
The goal should not be to say “we’re an AI-first company” or “our product uses AI”. The goal should be to solve the most pressing and relevant problems using the appropriate tools. LLMs are one of those tools. The first question asked should be “what problem are we solving?”, not “where can we apply AI?” When the problem to be solved is decided, only then can the solution be specified. What capabilities does the solution require? Where does the current solution, if any, fail? And is integrating an LLM into that solution actually necessary?
LLMs Trade Determinism for Flexibility
Every tool in the software development tool belt has strengths and weaknesses, and LLMs are no different. The question is not whether LLMs are useful, but whether the capabilities justify the trade offs they bring with them. LLMs provide flexibility by sacrificing determinism. They are useful when the work involved requires interpretation of natural language, synthesis across abundant or varied information, and step-by-step reasoning where the rules of a process can’t be specified before it begins.
These gains in flexibility result in losses in repeatability and determinism. LLMs produce variable output — that is, the same prompt given to the same model twice will produce tangibly different results. This makes testing and failure analysis much more difficult than for traditional code because success criteria are often subjective. Additionally, a process executed by an LLM will typically take much longer and cost more than the same workflow executed using plain code. These limitations mean that we need to evaluate the need for AI in a solution thoroughly before we blindly assume it will be helpful.
Six Questions Before You Add an LLM
When evaluating if an LLM will be useful, it’s helpful to think through the decision using concrete questions.
Can the workflow that implements the solution be specified completely and in advance? If yes, you may not need an LLM. Take, for example, a typical CI/CD pipeline: code is pushed to the remote, which triggers a workflow that lints, runs code tests, and deploys to a development environment. This workflow can be completely specified before it has commenced and runs the same way every time.
Must identical inputs produce identical outputs? LLMs are nondeterministic. If your workflow needs to lead to the same output across identical inputs, you may want to consider excluding LLMs. For example: a payment system receives the same invoice, tax jurisdiction, and retry request twice. It must calculate the same total and avoid charging the customer twice.
Does the solution require interpreting ambiguity that arises during execution? LLMs excel at ambiguity during execution. If a file isn’t located in the given folder, where should the solution look next? Deterministic rules can handle ambiguity up to a certain level, but traditional rule implementation requires thinking through in advance all the ways in which the solution’s execution could stray from the happy path. Note that this is different from “does the unexecuted solution have some ambiguity?” If the solution has ambiguity that can be resolved before execution begins, resolve the ambiguity and then use plain code.
Can the result of the solution be verified cheaply and accurately? Everyone can vibe code now because code can be cheaply tested against explicit criteria. If the required functionality exists, the tests pass, and the server remains stable, the code can probably be deemed workable. On the other hand, verifying if a medical diagnosis is safe to recommend to a patient is time-intensive and requires specialized expertise. Generally, outputs that require large amounts of work and/or special human expertise to verify (is this legal advice correct?) or whose verification is time-bound (will this business strategy lead to success?) are often not good candidates for LLM use.
What happens when the output is wrong? “When”, not “if”. Systems fail, and we must prepare for that. When the solution’s output is wrong, how severe are the consequences, and can the consequences be mitigated? A solution that mistakenly gives a customer a $15 discount is a much lower risk than one that gives a customer a $5000 discount. Wrong outputs can be mitigated by placing the solution further from the point of consequence: consider a chat bot that advises a human support agent about discounts they should provide instead of a bot that gives discounts directly to customers.
Does the LLM-based solution substantially outperform the simpler alternative? I recommend the KISS methodology here (Keep It Simple, Stupid) — if there’s a code-based alternative that can achieve 95% of the quality (error rate, number of human interventions, cost, etc.) of an LLM-based solution, go with the code-based solution.
Applying the Framework to My Prospecting System
In my prospecting system, I originally had the LLM execute the entire prospecting pipeline, searching, deduplicating, validating, and DB insertion. The model would generate the search queries, execute the queries with an MCP web search tool, deduplicate the returned prospects against the ones currently in the DB, validate the prospects against my ideal client description, and then write the surviving list into the CRM. Because there were so many steps, it messed up often.
Let’s use this example to think through the questions I outlined in the previous section.
Can the workflow be specified in advance? Yes — we just outlined the full workflow above.
Must identical inputs lead to identical outputs? Yes. If a prospect file is valid in one run, it should be valid in the next.
Does the solution require interpreting ambiguity that arises during execution? Yes — but not in all of the steps. Generating queries requires interpreting the types of prospects that already exist in the CRM and writing queries that will fill gaps. But other steps have no ambiguity: executing queries via Decodo search, deduplicating results from existing entries, and insertion into the CRM database.
Can the result of the solution be verified cheaply and accurately? Yes — I can verify if a prospect in the database fits my ideal client description by reading the name, company, and role title in ten seconds or less.
What happens when the output is wrong? In my case, nobody dies or loses large amounts of money. Maybe I accidentally send a message to a prospect twice, or I send a message to an invalid prospect, which can be corrected.
Does the LLM-based solution substantially outperform the simpler alternative? Yes — the alternative solution is that I think of queries myself, pass those to a deduplicating search script, manually validate prospects against my ideal client description and mark the most promising for insertion. Either that or I implement some sort of deterministic query-generating gap-filling system and a hacky regex keyword-matching validation system — neither of which would be very effective or time efficient.
Use the LLM Where Ambiguity Lives
The decision of whether or not to use an LLM is not straightforward. Trade offs have to be made: flexibility versus determinism, adaptability versus accuracy, cost versus performance. Even in my relatively simple prospecting system, the answers to the validation questions are not always a clear yes or no. I’ve found this often to be the case: usually a solution that contains an LLM also contains deterministic validation and process-automation code. The best solutions combine the strengths of LLMs with the strengths of deterministic code — it’s not a one-or-the-other decision.
Though many problems can be solved using deterministic solutions and would not benefit from the addition of a large language model, we live in a time where many problems that previously could not be solved with a computer now can be. These solutions don’t require your engineers to spend hundreds of dollars in tokens every day — they require your team to evaluate the problem at hand and decide to use the right tool for the right job.
Where did replacing an LLM with deterministic code improve your system, and what new trade offs did that introduce?
