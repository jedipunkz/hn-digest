---
source: "https://projan.ai/blog/should-you-use-ai-to-write-user-stories"
hn_url: "https://news.ycombinator.com/item?id=49362701"
title: "Should you use AI to write user stories"
article_title: "Should You Use AI to Write User Stories? | Projan Blog"
image: "https://projan.ai/og-default.png"
author: "davec271"
captured_at: "2026-08-19T15:22:04Z"
capture_tool: "hn-digest"
hn_id: 49362701
score: 1
comments: 0
posted_at: "2026-08-19T15:12:45Z"
tags:
  - hacker-news
  - translated
---

# Should you use AI to write user stories

- HN: [49362701](https://news.ycombinator.com/item?id=49362701)
- Source: [projan.ai](https://projan.ai/blog/should-you-use-ai-to-write-user-stories)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T15:12:45Z

## Translation

タイトル: AI を使用してユーザー ストーリーを作成する必要がありますか
記事のタイトル: AI を使用してユーザー ストーリーを作成する必要がありますか? |プロジャンのブログ
説明: AI を使用してユーザー ストーリーを作成する必要がありますか?何がうまく下書きされているか、どこで文脈が発明されているか、そして議論に値するストーリーを維持する方法について、正直に考察します。

記事本文:
Projan の使用例 エンジニアリングとテクノロジー 製品 クリエイティブと戦略 ビジネス開発 コンサルティング 創設者 教育 公共部門と非営利の統合 価格設定 ブログ ログイン サインアップ 使用例
2026 年 8 月 19 日 |デイブ・クリソルド | 9 分で読めます
AI を使用してユーザー ストーリーを作成する必要がありますか?
AI を使用してユーザー ストーリーを作成する必要がありますか?何がうまく下書きされているか、どこで文脈が発明されているか、そして議論に値するストーリーを維持する方法について、正直に考察します。
AI を使用して、判断ではなく、仕事の機械的な部分のユーザー ストーリーを作成します。モデルは、私が望むように、その形状を適切に生成し、その形状は決して難しいことではありませんでした。どのユーザー、どの問題、そしてなぜ今なのかを、モデルがこれまでに交わしたことのない会話から判断します。
これは、バックログを超えてすでにモデルを実行している、または実行していない理由を尋ねられているプロダクト マネージャー、ビジネス アナリスト、技術リーダーを対象としています。ここでは、AI がどのような点で優れているのか、生成されたストーリーがどのように失敗するのか、そしてその境界線はどこにあるのかについて説明します。
はい、簡単に。モデルに機能名を与えると、十数のストーリーが正しい形式で返されます。それぞれに役割、要望、根拠があり、そのほとんどは文法的で、いくつかはもっともらしいものです。テンプレートは空白を埋める演習であり、空白を埋めることはまさにこれらのシステムが行うように構築されているものです。
AI が書いたストーリーに対する反対意見は、出力が間違っているように見えるということではないので、このことははっきりと認める価値があります。見た目も良いですね。どのユーザーがこれを要求したのか、ビルドしないと何が壊れるのか、それが機能したことを後からどうやって知ることができるのかを誰かが語れるときに、ストーリーが成立します。残りは、エンジニアが対応する必要がある作業のフォーマット規則です。
したがって、能力の問題は解決されており、興味がありません。重要な問題は、部屋の外から来るはずだった物語の部分がどうなるかということです。
W

AI はバックログで本当に得意です
AI は、文字起こし、再構築、完全性チェックといったストーリー執筆の部分でその役割を果たしています。具体的には:
すでに持っているメモを再フォーマットします。発見メモのページ、記録された決定事項、書きかけのチケットが入ります。一貫した物語の形が出てきます。すべてがすでにそこにあったので、何も発明されません。
明らかに大きすぎるストーリーを分割する。ワークフロー ステップ、ユーザー タイプ、またはルールのバリエーションごとに分割を要求すると、適切なセットの中から選択できます。どちらの分割が実際であるかは、サイジング ルールによって決まります。
受け入れ基準に最初に合格し、その後議論します。ドラフトが正しいことはほとんどありません。空のリストに直面するよりも、間違ったリストを攻撃する方がはるかに速いです。ここでは、どの基準形式で議論しているかが重要です。
構造的なギャップを捉える。アクターが欠けているストーリー、値句が欠けているストーリー、またはユーザーがまったく含まれていないソリューションとして書かれたストーリー。これは校正であり、モデルは適切に校正されます。
退屈なバリエーションを作成します。エラー状態、空の状態、許可のケース、有効期限、2 回目の送信で何が起こるか。人間は金曜日の午後にはこれらをスキップします。モデルは飽きません。
下の模様は5枚とも同じです。 AI は、答えがすでにわかっていて、それを一貫して書き留めておきたい場合に役立ちます。答えがわからない場合は危険です。なぜなら、とにかく答えを提供し、提供された答えは実際のものとまったく同じであるからです。
AI が書いたストーリーが確実に失敗する場所
彼らは、誰もチェックしていないことについて自信を持って具体的に言うことで失敗します。障害はランダムなノイズではなく、クラスター化されています。
発明したユーザー。生成されたバックログは、コンプライアンス担当者、パワー ユーザー、および初めての訪問者に喜んで対応します。誰もそれらの人々を調査しなかったとしても、ペルソナ向けにストーリーが書かれたことになります。

at はプロンプトのエコー内にのみ存在します。
制約するのではなく、改めて述べる受け入れ基準。 「ユーザーがログインしている場合、保存をクリックするとレコードが保存される」ということは、開発者がまだ想定していないことを何も伝えません。実際の基準は、コードレビューで議論されるであろうケースを特定します。
他人のドメインからのありそうなエッジケース。モデルは、汎用システムのエッジケースに適しています。あなたのものには、トレーニング データのどこにも現れない年間プランの一部返金に関するルールがあり、生成されたリストにはそれが含まれていませんが、適用されない 6 つが自信を持って含まれています。
洗練を打ち破るボリューム。人は朝に8つの物語を書くことができます。モデルは 80 個を書くことができます。 80 のストーリーを適切に洗練する人はいないため、バッチは流し読みされて無視され、洗練はひっそりとレビューではなくなります。すべてが承認された状態でセッションが早期に終了し始めた場合、それが症状です。
責任のギャップ。最もダメージを与えるもの。なぜストーリーがスプリントにあるのかを尋ねると、正直な答えは肩をすくめることになります。誰も一行ずつ擁護できないバックログは、薄いバックログよりも悪いです。なぜなら、薄いバックログはそれを埋める会話を目に見えて促しますが、完全なバックログは会話を閉じてしまうからです。
AI PRD ジェネレーター vs 自分で作成する
どちらの純粋な選択肢も正しいわけではなく、初稿のスピードと文脈の由来との間でトレードが行われます。アーティファクトがストーリー、叙事詩、またはレビューに耐えなければならない完全な要件ドキュメントであるかどうかに関係なく、比較は同じです。
右側の列はノスタルジーではありません。その利点は、遅さが効果を発揮することです。「リピート客として」を手書きで書くには、リピート客が実際にそのように行動するかどうかを少し考えずに書くことはできません。その一時停止こそがこの演習の価値のすべてです。
の

3 番目の形状は、要求されるまで製図を行わないツールです。 Projan はそのように機能します。質問は最初に Slack または Microsoft Teams のチームに送られ、ストーリーがどのユーザーに提供されるのか、レビュー担当者が拒否する理由は何なのかを迫られます。ストーリーはプロンプトからではなくそれらの回答に基づいて作成され、その後 Jira または Linear にエクスポートされます。ユーザーについてはまだ何も知りません。違いは、モデルが推測するのではなく、質問に答える人によってギャップが埋められることです。
AIで物語を書いても大丈夫ですか？
はい、指名された人物が各ストーリーを所有しており、モデルに言及することなくそれを擁護できる限り、可能です。これがテストの全体であり、スプリントに相当する作業を生成したり、誰も反対しなかったものを採用したりすることが除外されるため、これは思っているよりも厳格です。
誰かが尋ねる前に、2 つの実践的な規範に同意する価値があります。
所有権は移転しません。ストーリーをボードに載せた人は誰でも、誰も解決できなかった問題を解決したことが判明したときに、改良、レビュー、そしてレトロでそのストーリーに答えます。
クレームが誰から入力されたかではなく、どこから来たのかを述べてください。チケットに開示バナーを付ける必要はありません。 「ユーザーは確認メールを期待している」がサポート チケットから来たのか、それともソフトウェアの動作に関するモデルの一般的な感覚から来たのかを知る必要があります。この区別によって、その請求に異議を申し立てることができるかどうかが決まります。
チームは、半分専門的で半分領土的な理由で、これについて神経質になります。プロの半分は合法です。ストーリーライティングは、製品担当者がユーザーを理解していることを証明する必要がある数少ないポイントの 1 つであり、それをジェネレーターに渡すと、労力とともに証拠が削除されます。
AI を使用してユーザー ストーリーを作成する必要がありますか?
答えがすでにどこかに存在し、形を整える必要がある場合に使用します。答えが必要な場合は使用しないでください

モデルは答えの形で何かを生成し、置換に気付かないため、es はまだ存在しません。
ソース資料はメモ、チケット、トランスクリプト、または記録された決定の中に存在します
タスクは構造的です: 分割、再フォーマット、欠落しているアクターまたは値句のチェック
すでに検討したケースのバリエーションを生成している
特定の誰かがボードに届く前にすべての行を読みます
次の場合は、自分で書くか、むしろ他の人と一緒に書いてください。
推測せずにユーザーに名前を付けることはできません
このストーリーは、スコープの決定やチーム間のトレードオフをエンコードしています。
受け入れ基準は、まだ起こっていない議論を解決します
この機能は、誰かの頭の中にだけ存在するドメイン ルールに触れる
一線を越えたことは、一度探せば簡単に分かります。プロンプト内のどの文が出力の特定の詳細を生成したかを特定できない場合、その詳細は発明されたものであり、それを構築しようとしているところです。
一般的ではないユーザー ストーリーを AI に求めるにはどうすればよいでしょうか?
話題ではなく生の素材を与えてください。インタビューメモ、サポートチケット、エンジニアが提起した制約、チームがすでに使用している受け入れ基準パターンを貼り付けます。機能に名前を付けるプロンプトにより、インターネット上のすべての同様の機能の妥当な平均が得られます。詳細を示すプロンプトにより、詳細が整理されて表示されます。
AIはストーリーだけでなく合格基準も書けるのでしょうか？
ドラフトされる可能性があり、最初のパスは通常柔らかすぎます。生成された基準は、ストーリーを制約するのではなく、衣服の「与えられた、いつ、その後」で言い換える傾向があります。それらを議論するためのリストとして扱います。それぞれについて、文言を満たしながらも意図を満たさない開発者が何を出荷できるかを尋ねます。
AI はエピックを小さなユーザー ストーリーに分割できますか?
はい、そして

分割は主にパターン マッチングのジョブであるため、これはより良い用途の 1 つです。ワークフロー ステップごと、ユーザー タイプごと、ルールのバリエーションごと、ハッピー パスとエラー パスごとに分割を依頼し、単独では価値を提供しないものは破棄します。どのスライスが最初に出荷されるかは、あなたに判断が委ねられています。
AI が起草したストーリーを一度にいくつまで洗練に取り込む必要がありますか?
チームが実際に読み取ることができる数は、モデルが生成できる数よりもはるかに少ないです。ボード上で生成された 60 のストーリーは洗練されるのではなく、流し読みされて承認されます。予約したセッションに合わせてバッチを切り出し、残りはバックログに放置せずに削除します。
ユーザー ストーリーの形式はその中で最もコストがかからない部分であり、AI によって確実に改善される唯一の部分です。実際にサービスを提供しているユーザーについての議論に時間を費やしてください。完了したように見えるバックログは、その部屋の誰もが同意するバックログと同じではありません。
協力すると物事がより良くなります
一度書いてみましょう。全員に同意してもらいます。
Projan は、エンジニアが行う前に、測定不可能な成功基準とテストされていない前提を明らかにします。
14 日間の無料トライアル。クレジットカードは必要ありません。
Projan はエンジニアと同じように仕様を読み取り、スプリントを開始する前にギャップについて質問します。
© 2026 Projan AI Ltd.全著作権所有。
Projan AI Ltd. イングランドおよびウェールズで登録 (17196385)。

## Original Extract

Should you use AI to write user stories? An honest look at what it drafts well, where it invents context, and how to keep a story worth arguing about.

Projan Use cases Engineering & Technology Product Creative & Strategy Business Development Consulting Founder Education Public Sector & Non-profit Integrations Pricing Blog Log in Sign Up Use cases
19 August 2026 | Dave Clissold | 9 min read
Should You Use AI to Write User Stories?
Should you use AI to write user stories? An honest look at what it drafts well, where it invents context, and how to keep a story worth arguing about.
Use AI to write user stories for the mechanical part of the job and not for the judgement. A model produces the As a, I want, so that shape competently, and that shape was never the hard bit. Deciding which user, which problem and why now comes from conversations the model has not had.
This is for product managers, business analysts and tech leads already running a model over the backlog, or being asked why they are not. It covers what AI does well here, how generated stories go wrong, and where the line sits.
Yes, easily. Give a model a feature name and it returns a dozen stories in correct form, each with a role, a want and a rationale, most of them grammatical and several of them plausible. The template is a fill-in-the-blanks exercise, and filling in blanks is exactly what these systems are built to do.
It is worth conceding that plainly, because the objection to AI-written stories is not that the output looks wrong. It looks good. A story earns its place when someone can say which user asked for this, what breaks if it is never built, and how you will know afterwards that it worked. The rest is a formatting convention for work an engineer has to act on .
So the capability question is settled and uninteresting. The question that matters is what happens to the parts of a story that were supposed to come from outside the room.
What AI is genuinely good at in a backlog
AI earns its keep on the parts of story writing that are transcription, restructuring or completeness checking. Specifically:
Reformatting notes you already have. A page of discovery notes, a recorded decision and a half-written ticket go in; consistent story shape comes out. Nothing is invented because everything was already there.
Splitting a story that is plainly too big. Ask for splits by workflow step, user type or rule variation and you get a reasonable set to choose from. Sizing rules still decide which of those splits is real .
A first pass at acceptance criteria you then argue with. The draft is rarely right. It is much faster to attack a wrong list than to face an empty one. Which criteria format you are arguing in matters here .
Catching structural gaps. Stories missing an actor, missing a value clause, or written as a solution with no user in them at all. This is proofreading, and models proofread well.
Producing the boring variants. Error states, empty states, permission cases, expiry, what happens on a second submission. Humans skip these on a Friday afternoon. The model does not get tired.
The pattern underneath all five is the same. AI is useful where you already know the answer and want it written down consistently. It is dangerous where you do not know the answer, because it will supply one anyway and the supplied answer will read exactly like the real ones.
Where AI-written stories reliably fail
They fail by being confidently specific about things nobody checked. The failures are not random noise, they cluster:
Invented users. A generated backlog will happily address the compliance officer, the power user and the first-time visitor. If nobody researched those people, you now have stories written for personas that exist only in the prompt’s echo.
Acceptance criteria that restate rather than constrain. “Given the user is logged in, when they click save, then the record is saved” tells a developer nothing they did not already assume. Real criteria pin down the case that would otherwise be argued about in code review.
Plausible edge cases from someone else’s domain. Models are good at the edge cases of a generic system. Yours has a rule about part-refunds on annual plans that appears in no training data anywhere, and the generated list will not include it while confidently including six that do not apply.
Volume that defeats refinement. A person can write eight stories in a morning. A model can write eighty. Nobody refines eighty stories properly, so the batch gets skimmed and waved through, and refinement quietly stops being a review. If your sessions have started ending early with everything approved, that is the symptom.
The accountability gap. The most damaging one. Ask why a story is in the sprint and the honest answer becomes a shrug. A backlog nobody can defend line by line is worse than a thin one, because a thin backlog visibly prompts the conversation that fills it, while a full one closes the conversation down.
AI PRD generator vs writing it yourself
Neither pure option is right, and the trade is between speed of first draft and where the context comes from. The comparison is the same whether the artefact is a story, an epic or a full requirements document that has to survive review .
The right-hand column is not nostalgia. Its advantage is that the slowness is doing work: you cannot write “as a returning customer” by hand without briefly wondering whether returning customers actually behave that way, and that pause is the entire value of the exercise.
The third shape is a tool that will not draft until it has asked. Projan works that way: the questions go to the team in Slack or Microsoft Teams first, pressing on which user a story serves and what would make a reviewer reject it, and the stories are written from those answers rather than from the prompt, then exported to Jira or Linear. It still knows nothing about your users; the difference is that the gap gets filled by a person answering a question instead of by the model guessing.
Is it okay to write stories with AI?
Yes, as long as a named person owns each story and can defend it without mentioning the model. That is the whole test, and it is stricter than it sounds, because it rules out generating a sprint’s worth of work and adopting whichever ones nobody objected to.
Two practical norms are worth agreeing before anyone asks:
Ownership does not transfer. Whoever puts the story on the board answers for it in refinement, in review and in the retro when it turns out to have solved a problem nobody had.
Say where a claim came from, not who typed it. Nobody needs a disclosure banner on the ticket. They do need to know whether “users expect an email confirmation” came from a support ticket or from the model’s general sense of how software behaves. That distinction decides whether the claim can be challenged.
Teams get twitchy about this for reasons that are half professional and half territorial. The professional half is legitimate. Story writing is one of the few points where a product person is forced to demonstrate that they understand the users, and handing that to a generator removes the evidence along with the effort.
Should you use AI to write user stories?
Use it when the answer already exists somewhere and needs shaping. Do not use it when the answer does not exist yet, because the model will produce something in the shape of an answer and you will not notice the substitution.
The source material exists in notes, tickets, transcripts or a recorded decision
The task is structural: splitting, reformatting, checking for missing actors or value clauses
You are generating variants of a case you have already thought through
Someone specific will read every line before it reaches the board
Write it yourself, or rather write it with other people, when:
You cannot name the user without guessing
The story encodes a scope decision or a trade-off between teams
The acceptance criteria will settle an argument that has not happened yet
The feature touches a domain rule that lives only in someone’s head
The tell that you have crossed the line is easy to spot once you look for it. If you cannot say which sentence in the prompt produced a given detail in the output, that detail was invented, and you are about to build it.
How do you prompt AI for a user story that is not generic?
Give it the raw material rather than the topic. Paste the interview notes, the support ticket, the constraint an engineer raised, and the acceptance criteria pattern your team already uses. A prompt that names a feature gets you a plausible average of every similar feature on the internet. A prompt carrying your specifics gets you your specifics, tidied.
Can AI write acceptance criteria as well as the story?
It can draft them, and the first pass is usually too soft. Generated criteria tend to restate the story in Given, When, Then clothing rather than constrain it. Treat them as a list to argue with: for each one, ask what a developer could ship that satisfies the wording and still fails the intent.
Can AI split an epic into smaller user stories?
Yes, and this is one of its better uses, because splitting is mostly a pattern-matching job. Ask for splits by workflow step, by user type, by rule variation and by happy path versus error path, then throw out the ones that do not deliver value alone. The judgement left to you is which slice ships first.
How many AI-drafted stories should you take into refinement at once?
As many as the team can genuinely read, which is far fewer than a model can produce. Sixty generated stories on a board do not get refined, they get skimmed and approved. Cut the batch to what fits the session you have booked, and delete the rest rather than leaving them to rot in the backlog.
The format of a user story is the cheapest part of it and the only part AI reliably improves. Spend the time it gives back on the argument about which user you are actually serving. A backlog that looks finished is not the same thing as a backlog anyone in the room agrees with.
Things are made better when we collaborate
Write it once. Have everyone agree.
Projan surfaces the unmeasurable success criteria and the untested assumptions, before your engineers do.
14-day free trial. No credit card required.
Projan reads a spec the way your engineers will, and asks about the gaps before the sprint starts.
© 2026 Projan AI Ltd. All rights reserved.
Projan AI Ltd. Registered in England & Wales (17196385).
