---
source: "https://stackoverflow.blog/2026/06/17/ai-agents-expose-the-security-checks-you-never-actually-wrote/"
hn_url: "https://news.ycombinator.com/item?id=48598825"
title: "AI agents are a confused deputy with the keys to your kingdom"
article_title: "AI agents are a confused deputy with the keys to your kingdom - Stack Overflow"
author: "fabsalvadori"
captured_at: "2026-06-19T16:04:54Z"
capture_tool: "hn-digest"
hn_id: 48598825
score: 3
comments: 0
posted_at: "2026-06-19T14:10:43Z"
tags:
  - hacker-news
  - translated
---

# AI agents are a confused deputy with the keys to your kingdom

- HN: [48598825](https://news.ycombinator.com/item?id=48598825)
- Source: [stackoverflow.blog](https://stackoverflow.blog/2026/06/17/ai-agents-expose-the-security-checks-you-never-actually-wrote/)
- Score: 3
- Comments: 0
- Posted: 2026-06-19T14:10:43Z

## Translation

タイトル: AI エージェントは王国の鍵を手にした混乱した代理人です
記事のタイトル:AIエージェントはあなたの王国への鍵を持った混乱した代理人です - コードログ

記事本文:
2026 年 6 月 17 日 AI エージェントは王国への鍵を手にした混乱した代理人です
攻撃者はどのようにして Meta の AI に丁寧に質問して 2 万の Instagram アカウントを乗っ取ったのか、そしてなぜその失敗が一般的になりつつあるのか。
クレジット: Alexandra Francis 6 月初旬、攻撃者はエクスプロイトを作成したりパスワードを 1 つも推測することなく、休眠状態にあったオバマ時代のホワイトハウス アカウントを含む 20,000 以上の Instagram アカウントを制御しました。彼らは Meta の AI サポート アシスタントとのチャットを開始し、自分が管理している電子メール アドレスを自分が所有していないアカウントに添付するよう依頼し、そのアドレスのパスワードをリセットするよう依頼しました。 Meta は後にログがすでに示していることを確認しました。アシスタントは設計どおりに動作しましたが、システムの別の部分が電子メールがアカウントに属していることを検証することになっており、そのチェックは実行されませんでした。
これを AI のミスと呼ぶのは、何が起こったのかを見逃しています。アシスタントは、話し相手に対して許可された一連の有効な操作を実行しました。この攻撃を止めたのは人物だった。見知らぬ人が有名人の回復メールのルートを変更しているのを見て、何かがおかしいと感じ、拒否したサポートスタッフだった。
現実世界の認証の大部分は、ソフトウェアとしてまったく作成されていませんでした。代わりに、それはリクエストとシステムの間に立つ人の裁量に依存し、その背後にあるものはすべて、裁量が常に存在することを前提として構築されていました。エージェントをその席に座らせると、裁量権はなくなり、下流では何も、そして誰も気づきません。エージェントはセキュリティ モデルをバイパスするのではなく、人である部分を公開するだけです。
チャットウィンドウを見て混乱する議員
セキュリティ業界では、メタが攻撃したものを正確に表す用語があります。それは「混乱した副官」です。実際の特権を保持しているプロセスは、特権の低い当事者によってその権限を使用するよう説得されます。

代わりに権限を与えます。ボスが送ってきたと電話をかけてきた人のために金庫の鍵を開けるのは夜警だ。ボスが鍵を持っている、いい話があると言う。正規の 1988 年のケースは、保護された請求ファイルに書き込むことができるコンパイラでした。そこに書き込むことができないユーザーが、コンパイラーに代わりにそれを行うように要求すると、コンパイラーはそれに応じました。コンパイラーには権限があり、誰のリクエストに対応するかを決して尋ねなかったからです。
LLM エージェントは、その構造上、これらの 1 つです。そのインターフェイスは自然言語であり、誰が何をする権限を持っているかという概念はなく、モデルの仕事全体は、もっともらしく聞こえる文をツール呼び出しに変えることです。直接 API リクエストには、少なくとも呼び出し元の ID が伴います。文にはそうではないため、通話が開始される前にその ID が再接続されない限り、エージェントは独自の権限に基づいて動作し、要求者の許可が関与することはありません。
また、エージェントは指示をデータから確実に分離することもできません。コンテキスト ウィンドウ内のすべてのものは、ユーザーのメッセージ、取得された文書、エージェントに要約するよう求められた電子メールの本文など、潜在的な指示として読み取られます。説得力のあるチャットで指示されたためにパスワードをリセットするサポート ボットは、処理するために渡されたファイル内に隠されたコマンドにもすぐに従います。 Meta の地理位置情報チェックを破った VPN のトリックは、これの粗雑なバージョンです。エージェントが取り込むコンテンツを通じて悪意のある命令が密かに持ち込まれる洗練されたバージョンは、エージェント攻撃の主要なクラスとしてすでに文書化されています。
爆発範囲は倍増しようとしています
Instagram ボットはパスワードをリセットする可能性があり、これは重大な侵害ではありますが、制限されたものです。現在出荷されているエージェントにはそのような制限はありません。同じ週に、Meta はサポート ツールを無効にし、ビジネス エージェントを開始しました。

予定を決め、リードを評価し、販売を完了し、支払いを受け取り、Shopify や Zendesk などのシステムに接続して会社に代わって行動します。支払い API と CRM を通じて同じ混乱した代理ロジックを実行すると、失敗はアカウントの盗難ではなくなります。それは、間違った相手に送られた返金、注文のルート変更、価格の上書き、顧客記録の編集であり、それぞれが依頼者に対してエージェントが実行する権限を与えられた正当な操作です。
市場はセキュリティモデルを追い越しています。 Gartner は、エンタープライズ アプリケーションの 40% に、2026 年末までにタスク固有の AI エージェントが含まれるようになると予測しています。これは、当初の 5% 未満から増加しています。これらのデプロイメントのほとんどは、Meta の場合と同じ前提を継承します。つまり、特権アクションの遠端にあるものには判断力があるということです。
より優れたモデルがこの問題を解決しない理由
同じワークフローの背後にあるより有能なモデルであれば、より優れた文法を使用して同じアカウントを引き渡すはずです。これがまさに、モデルが攻撃者が制御する部分である認証が存在する場所になり得ない理由です。アクションを許可するかどうかの決定は、アクションの外部で、何かが実行される前に実際にセッションの背後に誰がいるかをチェックするポリシー層によって行われる必要があります。 Meta のアシスタントは、回復メールを返送するまで、通話相手がアカウントを所有していることを確認しませんでした。
数行のコードで、出荷されるものは次のようになります。エージェントは関数を呼び出すことができ、それを呼び出すことができることが承認の全体でした。
def add_recovery_email(アカウント, new_email):
account.recovery_email = new_email # ここには発信者と関係するものは何もありません
send_reset_link(new_email) この修正は、よりスマートなモデルやより優れたプロンプトではありません。欠けていたのは主なチェックであり、チャットが影響を与える可能性のあるもの以外で決定されました。
pythondef add_recovery_email(アカウント

t、new_email、プリンシパル):
そうでない場合は、principal.owns(account): # 実際に質問しているのは誰か、検証済み
raise Unauthorized("セッションがアカウント所有者として認証されていません")
account.recovery_email = new_email
send_reset_link(新しいメール)
攻撃者は会話を制御していましたが、プリンシパルはチャットではなく認証されたセッションから取得されたものであるため、説得力のあるメッセージのシーケンスがその境界を満たすことはできません。
エージェントは、永続的なアクセスではなく、限定された短期間の権限を保持する必要があります。顧客のオープンチケットを要約するために発行されたトークンは、最後の注文の払い戻しには役に立たないはずです。これは最小限の特権ですが、アクションごと、リソースごとに適用する必要があり、セッションが開いてそれ以降信頼されるときに一度だけ付与するのではなく、エージェントはその資格情報で許可されているすべての情報にアクセスするように話しかけられるためです。
取り返しのつかないことには、エージェントが通過できないゲートが必要です。モデルが適切な単語を生成することで満足できる確認は、コントロールではなく、単なる綿毛です。支払い、削除、権限の変更、およびアカウントの回復は、人間の承認または厳格なポリシー ルールの背後にあり、日常的な検索と同じ道をたどるのではなく、与えることができる損害の程度によって分類されます。
エージェントが実行するすべてのアクションには、そのアクションを生成したプリンシパル、セッション、およびプロンプトを意味する出所が含まれている必要があるため、何が起こったかを監査し、何かが悪用された場合にアクションを取り消すことができます。 Instagram 攻撃が約 6 週間続いたことを考えると、インシデントの封じ込めと 2 万件のアカウントの盗難との間の距離は、多くの場合、共通点のないアカウントに対して 1 つの特権アクションが繰り返し実行されているのを、ほぼリアルタイムで誰かが確認できたかどうかにかかっています。
判断をコードに落とし込む
これはエージェントを実際のシステムから遠ざける理由にはなりません。構築する価値があります

そしてメタでうまくいかなかったのは通常のエンジニアリングのギャップであり、私たちが恐れるべきAIの性質ではありません。ここでのすべての修正は、チームがすでにその方法を知っているものです。資格情報のスコープを設定し、プリンシパルを検証し、元に戻せないアクションをゲートし、実行内容を記録します。
一つの習慣が彼らを結びつけます。エージェントを重要な事柄に接続する前に、そのループにいる人が何をチェックしていたのかを尋ねてください。その判断は実際の作業であり、エージェントが即興で判断してくれるわけではないため、コードとして存在する必要があります。そうすれば、従順は責任ではなくなります。許可されている内容を正確に実行するエージェントは、許可されている内容を決定する作業をユーザーが行っている限り、まさにユーザーの希望どおりです。
AI システム アーキテクトおよびインディー ビルダーは、エージェント AI、信頼できる自動化、開発者インフラストラクチャに重点を置いています。
私は、影響力の高い AI エージェントのアクションのためのオープンソース ガバナンス層である PIC Standard と、… セキュリティ 実行中のソフトウェア 最近の記事 2026 年 6 月 19 日 オライリーからの発信: 能力から責任へのプロトコル指向のアプローチである NCP の作成者です。
2026 年 6 月 15 日 Selenium vs. Cypress vs Playwright: テスト自動化フレームワークの選択
2026 年 6 月 12 日 CherryScript の設計: カスタム Python ベースのインタープリターによるデータ駆動型ワークフローの最適化
2026 年 6 月 12 日 ポケベル慈善活動?リーダーにチームを無理に動かさないようにするにはどうすればよいでしょうか?
2026 年 6 月 19 日 あなたは自分で思っているほど DNS を理解していません
ディスカッションに参加するには、stackoverflow.com アカウントでログインしてください。私たちのスタック

## Original Extract

June 17, 2026 AI agents are a confused deputy with the keys to your kingdom
How attackers took twenty thousand Instagram accounts by asking Meta's AI politely, and why that failure is about to become common.
Credit: Alexandra Francis Earlier in June, attackers took control of more than twenty thousand Instagram accounts, including the dormant Obama-era White House account, without writing an exploit or guessing a single password. They opened a chat with Meta's AI support assistant, asked it to attach an email address they controlled to an account they did not own, and requested a password reset to that address. Meta later confirmed what the logs already showed: the assistant behaved exactly as designed, while a separate part of the system was supposed to verify that the email belonged to the account, and that check never ran.
Calling this an AI mistake misses what happened. The assistant carried out a valid sequence of permitted operations for whoever was talking to it. What would have stopped the attack was a person: a support worker who saw a stranger rerouting a celebrity's recovery email, sensed something was wrong, and refused.
A large share of real-world authorization was never written as software at all. Instead, it lived in the discretion of whoever stood between a request and the system, and everything behind them was built assuming that discretion would always be there. Put an agent in that seat and discretion vanishes, while nothing and nobody downstream notices. The agent does not bypass your security model, it just exposes the part of it that was a person.
A confused deputy with a chat window
Security has a precise term for what Meta hit: the confused deputy . A process holding real privileges is talked by a less-privileged party into using those privileges on its behalf. It's the night guard who unlocks the vault for anyone who calls and says the boss sent them: he's got the keys, they've got a good story. The canonical 1988 case was a compiler that could write to a protected billing file. A user who could not write there asked the compiler to do it for them, and it complied, because it had the authority and never asked whose request it was serving.
An LLM agent is one of these by construction. Its interface is natural language, which carries no notion of who is authorized to do what, and the model's whole job is to turn a plausible-sounding sentence into a tool call. A direct API request at least brings the caller's identity along with it. A sentence does not, so unless that identity is reattached before the call fires, the agent acts on its own authority and the requester's permissions never enter the picture.
Agents also cannot reliably separate instructions from data. Everything in the context window reads as potential instruction: the user's message, a retrieved document, the body of an email the agent was asked to summarize. A support bot that resets a password because a convincing chat told it to will just as readily follow a command hidden inside a file it was handed to process. The VPN trick that defeated Meta's geolocation check is the crude version of this. The sophisticated versions, where the malicious instruction is smuggled in through content the agent ingests, are already being documented as the dominant class of agent attack.
The blast radius is about to multiply
The Instagram bot could reset passwords, and that's a serious breach, but a bounded one. The agents shipping now are not bounded that way. In the same week Meta disabled the support tool, it launched its Business Agent, which books appointments, qualifies leads, closes sales, takes payments, and connects to systems like Shopify and Zendesk to act on a company's behalf. Run the same confused-deputy logic through a payment API and a CRM and the failure is no longer a stolen account. It is a refund sent to the wrong party, an order rerouted, a price overridden, a customer record edited, each one a legitimate operation the agent was authorized to perform for whoever asked.
The market is outrunning the security model. Gartner has projected that 40% of enterprise applications will include task-specific AI agents by the end of 2026 , up from under 5% at the start of it. Most of those deployments will inherit the same assumption Meta's did, which is that whatever sits at the far end of a privileged action has judgment.
Why a better model does not fix this
A more capable model behind the same workflow would have handed over the same accounts with better grammar, which is exactly why the model cannot be the place authorization lives, being the part an attacker controls. The decision to allow an action has to be made outside it, by a policy layer that checks who is actually behind the session before anything runs. Meta's assistant never established that the person it was talking to owned the account before it rebound their recovery email.
In a couple lines of code, what shipped looked something like this. The agent could call the function, and being able to call it was the whole authorization:
def add_recovery_email(account, new_email):
account.recovery_email = new_email # nothing here ties to the caller
send_reset_link(new_email) The fix is not a smarter model or a better prompt. It is the principal check that was missing, decided outside anything the chat can influence:
pythondef add_recovery_email(account, new_email, principal):
if not principal.owns(account): # who is actually asking, verified
raise Unauthorized("session not authenticated as the account owner")
account.recovery_email = new_email
send_reset_link(new_email)
The attacker controlled the conversation, but principal comes from the authenticated session, not the chat, so no sequence of convincing messages can satisfy that line.
Agents should hold scoped, short-lived authority instead of standing access. A token minted to summarize a customer's open tickets should be useless for refunding their last order. That is least privilege, but it has to be enforced per action and per resource, not granted once when the session opens and trusted from then on, because an agent will be talked into reaching for everything its credentials permit.
Anything irreversible needs a gate the agent cannot drive through. A confirmation the model can satisfy by generating the right words is not a control, just fluff. Payments, deletions, permission changes, and account recovery belong behind a human approval or a hard policy rule, classified by how much damage they can do rather than waved through on the same path as a routine lookup.
Every action an agent takes should carry its provenance, meaning the principal, the session, and the prompt that produced it, so you can audit what happened and revoke it when something is exploited. Considering that the Instagram attack ran for roughly six weeks, the distance between a contained incident and twenty thousand stolen accounts is often just whether anyone could see, close to real time, that one privileged action was firing again and again for accounts with nothing in common.
Putting the judgment into code
None of this is a reason to keep agents away from real systems. They are worth building, and what went wrong at Meta was an ordinary engineering gap, not some property of AI we have to fear. Every fix here is something teams already know how to do: scope the credentials, verify the principal, gate the actions you cannot undo, and keep a record of what ran.
One habit ties them together. Before you connect an agent to anything that matters, ask what the person in that loop used to check. That judgment was real work, and now it has to exist as code, because the agent will not improvise it for you. Do that, and obedience stops being the liability. An agent that does exactly what it is allowed to do is precisely what you want, as long as you have done the work of deciding what it is allowed to do.
AI systems architect and indie builder focused on agentic AI, trustworthy automation, and developer infrastructure.
I’m the creator of PIC Standard, an open-source governance layer for high-impact AI agent actions, and NCP, a protocol-oriented approach to … security Running software Recent articles June 19, 2026 Dispatches from O'Reilly: From capabilities to responsibilities
June 15, 2026 Selenium vs. Cypress vs Playwright: Choosing your test automation framework
June 12, 2026 Designing CherryScript: Optimizing Data-Driven Workflows via Custom Python-Based Interpreters
June 12, 2026 Paging Charity? How do I get my leaders to stop running teams Into the ground?
June 19, 2026 You don’t understand DNS like you think you do
Login with your stackoverflow.com account to take part in the discussion. Our Stack
