---
source: "https://tenureai.dev/writing/why-most-ai-evals-would-miss-the-linear-sales-email-failure/"
hn_url: "https://news.ycombinator.com/item?id=48636731"
title: "Why most AI evals would miss the Linear sales email failure"
article_title: "Why most AI evals would miss the Linear sales email failure | Tenure"
author: "jflynt76"
captured_at: "2026-06-22T22:33:30Z"
capture_tool: "hn-digest"
hn_id: 48636731
score: 4
comments: 0
posted_at: "2026-06-22T21:51:32Z"
tags:
  - hacker-news
  - translated
---

# Why most AI evals would miss the Linear sales email failure

- HN: [48636731](https://news.ycombinator.com/item?id=48636731)
- Source: [tenureai.dev](https://tenureai.dev/writing/why-most-ai-evals-would-miss-the-linear-sales-email-failure/)
- Score: 4
- Comments: 0
- Posted: 2026-06-22T21:51:32Z

## Translation

タイトル: ほとんどの AI 評価がリニア セールス メールの失敗を見逃してしまう理由
記事のタイトル: なぜほとんどの AI 評価者はリニア セールス メールの失敗を見逃してしまうのか |在職期間
説明: リニア

記事本文:
在職期間
プラットフォーム ▼ コア機能
ほとんどの AI 評価が Linear 営業メールの失敗を見逃してしまう理由
リニアの販売代理店は既存顧客に間違った社名を記載した電子メールを6回送信した。これを悪いAIの普及と呼ぶのは簡単です。しかし、メールは目に見える部分だけでした。本当の失敗は、その決定が依存する事実を証明することなく、システムが送信を許可すると決定したときの以前に発生しました。
@jmwind · 10:43 AM · 2026 年 6 月 22 日 こんにちは、@karrisaarinen さん、フィールドからのフレンドリーな注意事項です。あなたの営業チームの誰かから、コミカルな AI のスロップに関するメールを 6 通以上受け取りました。会社名が間違っている、すでに顧客である、など...いつも笑いが大好きですが、これはレベルを飛ばしてもいいですか?
AIのスロップはやめてください。あなたは会社を間違えました、私の電子メールドメインを見ませんでした、そして私たちはすでにLinearを使用しています。今キャンセルを考えています。
Jean-Michel Lemieux 開発者 · Spellbook.com 5 4 373 94K views KS Karri Saarinen @karrisaarinen · 2 時間 感謝とお詫び。理想的ではありません。原因をチームに確認します。既存の顧客に6回もメールを送るのは最も愚かなことであることに同意する
ほとんどの人は、悪い AI 普及を世代の問題として説明します。メッセージがぎこちなかったり、繰り返したり、個人的な内容が不十分でした。
しかし、より大きな失敗は通常、その一歩手前で起こります。何かを書き込む前に、システムは、受信者が誰であるか、どの会社に所属しているか、すでに顧客であるかどうか、アカウントでアウトリーチが許可されているかどうか、この人物がすでに何度も連絡を取っているかどうかを知る必要があります。
これらのチェックが間違っているか欠落している場合、より優れたモデルが問題を解決することはありません。間違ったアクションのよりクリーンなバージョンを作成するだけです。
そのため、エージェントの評価は上流に目を向ける必要があります。最終出力が適切に読み取れるかどうかだけでなく、システムが動作する前に何をチェックしたかを尋ねる必要があります。
G

RoundEval は、エージェントが応答または動作する前に、何を検索、取得、引用し、使用する許可を得たのかという質問を中心に構築されています。
メールが最初の失敗ではなかった
リニア電子メールに対する簡単な反応は、その出力を見て笑うことです。会社名
は間違っています。受取人はすでに顧客です。同じシーケンスがすでにヒットしていました
それらを複数回。その後、CEO が公的に返答し、すべてが次のようになります。
AI のスロップのもう 1 つの例。
しかし、その枠組みではシステムがあまりにも簡単に外れてしまいます。
恥ずかしいのは、誰もが見たメールです。もっと重要な部分が起こった
電子メールが存在する前。上流のどこかで、システムに十分な間違いがあった、または
この人に連絡する必要があるかどうかを決定するためのチェックがされていない状態。
これは、件名を改善しても修正できない部分です。暖かいトーンでは解決しない
それ。美しいアウトバウンドコピーを書くモデルであっても、
最初に基本的な事実を確認しなかった場合、間違ったメッセージが表示されます。
リニアの場合、送信前のチェックがすべてでした。会社名はありますか
受信者のドメインと一致しますか?この連絡先はすでに顧客ですか?このシーケンスがあります
すでに実行回数が多すぎますか？それらの答えが間違っているか、まったく確認されていない場合は、
すでに失敗した状態から始まっています。
目に見える失敗は、不正な電子メールでした。以前の失敗はもっと単純で、システムが原因でした。
電子メールが送信されるべきであることを証明しませんでした。
メッセージが存在する前に、システムには証明すべき事実が存在します
送信メールは外から見るとシンプルに見えます。連絡先を選択し、メッセージを書いて送信します。
実際の企業内では、アクションは true でなければならない状態のスタックに依存します。
この人は見込み客、アクティブな顧客、元顧客、パートナー、従業員、またはこのシーケンスを決して受け取ってはいけない人ですか?
メール内の会社名は、レシピにリンクされているアカウントと一致しますか?

pient、電子メール ドメイン、および現在の CRM レコードは何ですか?
アカウントはすでに製品を使用していますか、オープンな商談があるか、割り当てられた所有者がいますか、または抑制ルールの下にありますか?
この人には、どのチャネルを通じて、どのチームから何回連絡があり、どのような反応がありましたか?
このような状態をすべて考慮すると、このオートメーションは送信を許可されるのでしょうか、それとも抑制し、人間にルーティングするべきでしょうか、それとも何もしないでしょうか?
これらのチェックのいずれかが失敗した場合、正しい行動は「より良いメールを書く」ではありません。の
正しい行動は「送信しない」ことです。だからこそ、これをコンテンツの品質の問題と呼ぶのは間違いである
故障モード。生成されたテキストは、現場に残されたアーティファクトにすぎません。
メールの評価が遅すぎます
従来の評価により、最終的な電子メールを採点できます。礼儀正しいですかね。個性的ですか？そうですか
関連性のある。適切な製品について言及されていますか。ブランドの声に従っていますか？それは避けますか
明らかな幻覚。これらは有益な質問ですが、行動が終わった後に始まります。
すでに承認されています。
失敗した場合、電子メールはこれらすべての側面で良いスコアを獲得しても、依然として間違っている可能性があります。
洗練され、簡潔で、フレンドリーで、ブランドに基づいたものにすることができます。についての真実の記述が含まれることもあります。
製品。そのどれもが、システムがこれに送信するのに十分な検証済みの状態を持っていたことを証明するものではありません。
この時の人。
これは、GroundEval が質問応答エージェントに対して行うのと同じ区別です。最終的な答え
エージェントが文書を決して取得せず、証拠を使用したことをトレースが示している間、もっともらしく見える可能性があります
許可の範囲外にいた場合、または捜索せずに不在を主張した場合。アウトバウンドバージョン
同じ形状を持っています。送信前の証拠パスが真実である一方で、最終メッセージはもっともらしく見える可能性があります。
無効です。
代理人は行動する権利を獲得しましたか?
GroundEval は、エージェントの動作を州の契約に照らしてテストできるものとして扱います。
コントロール

法律は、どのような証拠が存在するのか、それがいつ存在したのか、誰または何にアクセスが許可されたのかを示します。
請求や訴訟が有効になる前にどのチェックが必要か。
アウトバウンド エージェントの場合、評価では電子メールが良好であったかどうかを問う必要はありません。それ
より単純で重要な質問をすることができます。送信する前に、エージェントは
必要なシステムがあり、有効な送信決定に達しているか?
それは裁判官のプロンプトではありません。電子メールがどのように見えるかについての雰囲気ベースのレビューではありません
合理的です。これは、何を検索したか、何を検索したかなど、証拠のパスに対する決定的なチェックです。
フェッチされたか、その時点でどのような状態が利用可能であったか、そしてそこからアクションが続いたかどうか。
エージェントには承認だけではなく前提条件が必要です
危険な自動化に対する通常の答えは、人間を関与させることです。誰かにレビューしてもらいましょう
エージェントが送信する前に。それは役に立ちますが、それは人間が生成されたもの以上のものを見ることができる場合に限られます。
メール。
洗練されたドラフトでは、アカウントレコードにアクティブな顧客が記載されていることは示されません。そうではありません
ドメインが別の場所を指していることを示します。連絡先がすでにあったことは示されていません
メールを5回送りました。提案されている製品がすでにインストールされていることを示すものではありません。
トレースがなければ、承認は同じ問題のより見栄えの良いバージョンになる可能性があります。の
査読者は成果物を判断しているのであって、成果物を生み出した決定ではありません。
システムには前提条件が必要です。エージェントが行動する前に、どのチェックを行うかを表示できる必要があります
必須か、どのシステムがクエリされたか、どのレコードが見つかったか、どのルールが許可されているか
継続するアクション。必要なチェックが欠落している場合、アクションはその前に停止する必要があります。
世代。
「送信してもよさそうだ」ということは、「必要な証拠がチェックされ、送信された」ということと同じではありません。
そこから決定が下されます。」
将来の失敗モードは単に間違ったことを言っているだけではありません
AI に関するよくある話

kはまだ答えに集中しています。モデルは幻覚を見たのですか？しました
間違った情報源を引用していますか？何か危険なことを言いましたか？それらの失敗は重要ですが、エージェントは
異なる体制に移行すること。彼らはただ答えるだけではありません。送信、ファイル、更新、ルーティング、
承認、エスカレーション、抑制。
ソフトウェアが動作できるようになると、評価対象が変わります。問題はもはや、
出力は正しかったです。エージェントがいた状態でアクションが有効であったかどうかになります
観察することが許される。
悪質な AI の普及は、その大きな変化の小規模ながら公的な例です。目に見える失敗は、
恥ずかしいメール。システム障害は、電子メールが送信される前に誰も証明できなかったことです。
エージェントが送信に必要な事実を確認したこと。
これは、GroundEval が検出するように設計された種類の障害です。散文が聞こえるかどうかではない
いいよ。その行動が得されたものであるかどうか。
LLM-as-judge がエージェント評価のデフォルトになりました
最終回答判定では、代理人が主張内容を知ることを許可されていたかどうか、いつ知ることができたのか、あるいは欠席が確認されたかどうかを確認することはできない。
AI ガバナンスはツール呼び出しではあり得ない
Enterprise AI には、エージェントが呼び出し先を選択できるツールではなく、すべてのモデル リクエストの下にコンテキスト コントロール プレーンが必要です。
メモリの精度、ノイズ分離、セッションターンレイテンシ、および信念の可変性を測定するための取得ベンチマーク。

## Original Extract

Linear

tenure
Platform ▼ Core Capabilities
Why most AI evals would miss the Linear sales email failure
Linear's sales agent emailed an existing customer six times with the wrong company name. It is easy to call that bad AI outreach. But the email was only the visible part. The real failure happened earlier, when the system decided it was allowed to send without proving the facts that decision depended on.
@jmwind · 10:43 AM · 6/22/26 Hey @karrisaarinen , friendly heads-up from the field. I've received 6+ emails from someone on your sales team with comical AI-slop. Wrong company name & already a customer, etc... Always love a good laugh, but you may want to skip-level this one?
Stop the AI slop pls. You got the company wrong, you didn't look at my email domain, and we already use Linear. Thinking of canceling now.
Jean-Michel Lemieux Developer · spellbook.com 5 4 373 94K views KS Karri Saarinen @karrisaarinen · 2h Thanks and apologies. Not ideal, will check with the team what caused this. Agree that emailing existing customers and 6 times is the dumbest thing
Most people describe bad AI outreach as a generation problem. The message was awkward, repetitive, or poorly personalized.
But the larger failure usually happens one step earlier. Before anything gets written, the system has to know who the recipient is, which company they belong to, whether they are already a customer, whether the account allows outreach, and whether this person has already been contacted too many times.
If those checks are wrong or missing, a better model does not solve the problem. It just writes a cleaner version of the wrong action.
That is why agent evaluation has to look upstream. It should ask what the system checked before acting, not only whether the final output reads well.
GroundEval is built around that question: what did the agent search, fetch, cite, and have permission to use before it answered or acted?
The email was not the first failure
The easy reaction to the Linear email is to laugh at the output. The company name
is wrong. The recipient is already a customer. The same sequence had already hit
them multiple times. Then the CEO replies publicly, and the whole thing becomes
another example of AI slop.
But that framing lets the system off too easily.
The embarrassing part is the email everyone saw. The more important part happened
before the email existed. Somewhere upstream, the system had enough wrong or
unchecked state to decide that this person should be contacted at all.
That is the part a better subject line would not fix. A warmer tone would not fix
it. Even a model that writes beautiful outbound copy would still have sent the
wrong message if it never checked the basic facts first.
In the Linear case, the pre-send checks were the whole story. Does the company name
match the recipient's domain? Is this contact already a customer? Has this sequence
already run too many times? If those answers are wrong or never checked, generation
is already starting from a failed state.
The visible failure was a bad email. The earlier failure was simpler: the system
did not prove that the email should be sent.
Before the message exists, the system has facts to prove
Outbound email looks simple from the outside. Pick a contact, write a message, send it.
Inside a real company, the action depends on a stack of state that has to be true.
Is this person a prospect, an active customer, a former customer, a partner, an employee, or someone who should never receive this sequence?
Does the company name in the email match the account linked to the recipient, the email domain, and the current CRM record?
Does the account already use the product, have an open opportunity, have an assigned owner, or sit under a suppression rule?
How many times has this person been contacted, through which channel, by which team, and with what response?
Given all of that state, is this automation allowed to send, or should it suppress, route to a human, or do nothing?
If any one of those checks fails, the right behavior is not "write a better email." The
right behavior is "do not send." That is why calling this a content quality problem misses
the failure mode. The generated text is only the artifact left at the scene.
Evaluating the email is too late
A conventional evaluation can grade the final email. Is it polite. Is it personalized. Is it
relevant. Does it mention the right product. Does it follow the brand voice. Does it avoid
obvious hallucinations. Those are useful questions, but they begin after the action has
already been approved.
In the failure case, the email can score well on all of those dimensions and still be wrong.
It can be polished, concise, friendly, and on brand. It can even contain true statements about
the product. None of that proves the system had enough verified state to send it to this
person at this time.
This is the same distinction GroundEval makes for question answering agents. A final answer
can look plausible while the trace shows the agent never fetched the document, used evidence
outside its permission boundary, or claimed absence without searching. The outbound version
has the same shape: a final message can look plausible while the pre-send evidence path is
invalid.
Did the agent earn the right to act?
GroundEval treats agent behavior as something that can be tested against a state contract.
The contract says what evidence exists, when it existed, who or what was allowed to access it,
and which checks are required before a claim or action is valid.
For an outbound agent, the evaluation does not have to ask whether the email was good. It
can ask a simpler and more important question: before sending, did the agent check the
required systems and reach a valid send decision?
That is not a judge prompt. It is not a vibes-based review of whether the email seems
reasonable. It is a deterministic check against the evidence path: what was searched, what
was fetched, what state was available at the time, and whether the action followed from it.
Agents need preconditions, not just approvals
The usual answer to risky automation is to put a human in the loop. Let someone review it
before the agent sends. That can help, but only if the human can see more than the generated
email.
A polished draft does not show that the account record says active customer. It does not
show that the domain points somewhere else. It does not show that the contact was already
emailed five times. It does not show that the product being pitched is already installed.
Without the trace, approval can become a nicer-looking version of the same problem. The
reviewer is judging the artifact, not the decision that produced it.
The system needs preconditions. Before an agent acts, it should be able to show which checks
were required, which systems were queried, which records were found, and which rule allowed
the action to continue. If the required checks are missing, the action should stop before
generation.
"Looks good to send" is not the same as "the required evidence was checked, and the send
decision follows from it."
The future failure mode is not just saying the wrong thing
The common story about AI risk still focuses on the answer. Did the model hallucinate? Did
it cite the wrong source? Did it say something unsafe? Those failures matter, but agents are
moving into a different regime. They do not only answer. They send, file, update, route,
approve, escalate, and suppress.
Once software can act, the eval target changes. The question is no longer only whether the
output was correct. It becomes whether the action was valid under the state the agent was
allowed to observe.
Bad AI outreach is a small, public example of that larger shift. The visible failure is the
embarrassing email. The system failure is that no one could prove, before the email went out,
that the agent had checked the facts required to send it.
That is the class of failure GroundEval is designed to catch. Not whether the prose sounds
good. Whether the behavior was earned.
LLM-as-judge became the default for agent evaluation
Final-answer judging cannot see whether the agent was allowed to know what it claimed, when it could have known it, or whether absence was checked.
AI governance cannot be a tool call
Enterprise AI needs a context control plane beneath every model request, not a tool the agent can choose to call.
A retrieval benchmark for measuring memory precision, noise isolation, session-turn latency, and belief mutability.
