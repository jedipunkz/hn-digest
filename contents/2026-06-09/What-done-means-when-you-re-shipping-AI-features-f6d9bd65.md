---
source: "https://jeffgothelf.com/blog/what-done-means-when-youre-shipping-ai-features/"
hn_url: "https://news.ycombinator.com/item?id=48460711"
title: "What \"done\" means when you're shipping AI features"
article_title: "What \"done\" means when you're shipping AI features"
author: "speckx"
captured_at: "2026-06-09T16:07:31Z"
capture_tool: "hn-digest"
hn_id: 48460711
score: 1
comments: 0
posted_at: "2026-06-09T13:13:47Z"
tags:
  - hacker-news
  - translated
---

# What "done" means when you're shipping AI features

- HN: [48460711](https://news.ycombinator.com/item?id=48460711)
- Source: [jeffgothelf.com](https://jeffgothelf.com/blog/what-done-means-when-youre-shipping-ai-features/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T13:13:47Z

## Translation

タイトル: AI 機能を出荷する際の「完了」とは何を意味するか
説明: AI は「完了」の古い定義を打ち破りました。あなたはもはや自動販売機を出荷しているのではなく、確率的な動作分布を出荷しているのです。

記事本文:
AI 機能を出荷する際の「完了」とは何を意味しますか
コンテンツにスキップ
ジェフ・ゴセルフフ
Sense & Respond Learning チーム向けの世界クラスの製品管理トレーニング オプション。
0
AI 機能を出荷するときの「完了」とは
ああ、スプリント レビューの古き良き時代。エンジニアリング部門は [何か] を出荷したと言っています。すべてのテストに合格しました。 P0 のバグはありません。デモは魅力的に機能しました。 「設計どおりに動作する」が再び実現されました。悲しいことに、このシンプルさを享受できる余裕はもうありません（かつてあったとしても）。
「完了」の定義は常に問題です。 AIの場合は完全に書き換える必要がある。
私たちは決定論的なソフトウェアの完了の定義を構築しました。
入力が入力され、出力が出力され、火曜日の出力は月曜日と同じになります。それが要点でした。私たちは予測可能なユーザー エクスペリエンスを構築していました。 QA がエッジをテストし、エンジニアリングが単体テストを作成し、仕様が動作を記述し、この 3 つがすべて揃った瞬間が「完了」となります。ソフトウェアが自動販売機のように動作したため、これが機能しました。お金を入れてボタンを押すと、ソーダが出てきます。
AI 機能は自動販売機のように動作しません。
同じプロンプトが、ユーザー、セッション、モデルの更新、QA チームが想像していなかったコンテキストにわたって異なる出力を生成します。あなたが出荷したものは、固定され、一貫性があり、予測可能な物体ではありません。それは行動の分布です。そのうちのいくつかはあなたも見たことがあり、予測できるものです。そうでない人もいます。驚きのほとんどは、あなたが知らないものの中に存在します。テスト スイートは、コードのプロパティを検証するように設計されています。これは、オークランドで午後 11 時に顧客が、仕様が予期していなかった質問で体験していることを検証するように設計されたものではありません。
「すべてのテストに合格する」ということは今でも重要であり、これまでと同じことを意味します。それはただ

もはや完了を意味しません。
AI を使用すると、仕様の遵守に関する二元的な結果ではなく、出力の許容可能な変動、さらに重要なことに動作に関する調整が完了します。機能が十分に適切に動作する条件の範囲、許容することを決定した障害モード、まだ実行していないものに対して導入した監視、実行の準備が整い実際にリハーサルしたロールバックによって完了を定義します。
最後の文には多くの内容が含まれているため、AI 対応の完了の定義を作成するために必要な 3 つのステップに分けて説明します。
ステップ 1: 受け入れ基準をアサーションではなく分布として記述する
「ユーザーが X を尋ねるか入力すると、システムは Y を返します。」それは主張です。バイナリなので検証可能です。 AI 機能の場合、それは多くの場合無意味です。
代わりに、次の枠組みを試してみてください。「カテゴリ X の入力の 80% について、システムは品質バー Y を満たす応答を返します。残りの 20% については、故障モードは低下していますが、恥ずかしいことではありません。」それは配布物です。書くのもテストするのも難しくなり、実際に出荷するものについては限りなく正直になります。
すべての基準が確率的である必要はありません。どの成功結果がどれであるかについて、自分自身とチームに対して正直になる必要があります。そして最も重要なのは、確率的な結果が決定的であるかのように振る舞うのをやめる必要があります（つまり、固定的で予測可能です）。私がこれを正しく実現しているチームは、同じ機能に関して 2 つの基準を設定していると見ています。
自動販売機のように動作するシステム部分の決定的なもの (認証、請求、ナビゲーションなど)
そうでない部分の分布的 (または確率的) もの (システム フィードバック、コンテキスト固有の応答、カスタム ユーザー リクエストなど)
ステップ 2: 障害トリアージ能力を構築する

打ち上げ後ではなく、打ち上げに参加する
従来、チームは機能を出荷し、ダッシュボードを監視し、対応する価値のあるレベルで苦情が表面化したときに優先順位を付けていました。 AI 機能の場合、ワークフローが遅すぎます。苦情がエンジニアリングに届くまでに、ユーザーは製品についての意見を形成しています。
AI 製品をサポートするには、AI 機能が開始される前にトリアージ プレイブックを作成します。モデルの品質の問題、UX の問題、コンテンツの問題、そしておそらくはモデルが誤って引き起こす可能性のある PR 災害の責任者を明確にしてください。リリースは単に機能が公開されるだけではありません。立ち上げは、結果に対処するチームが実際に対処する準備ができた瞬間です。これは、私が一緒に働いているほとんどのチーム、特に企業ではまだ当てはまりません。
機能がリリースされても完了ではありません。機能の下流にある人々が、機能が不正な動作をした場合に何をすべきかを理解すれば、作業は完了です。 AI なのでそうなります。
ステップ 3: 「これは結局終わっていない」というシグナルを定義し、ロールバックのリハーサルを行う
すべての AI 機能にはトリップワイヤーが付属している必要があります。製品とコンテキスト (およびブランド) にとって最も意味のある指標を選択できます。それは、エラー率、1,000 セッションあたりのオフトーンの苦情、人間によるレビューへのエスカレーション、サンプル監査でフラグが付けられた幻覚率など、最適化されていないユーザー エクスペリエンスの指標であるあらゆるものである可能性があります。次に、しきい値を設定し、しきい値を超えた場合に何をするかを起動前に決定します。
ワイヤーが切れたらどうするのですか？理論上ではなく、実際に。もちろん、書き留めることもできます。ただ書くだけではなく、ロールバックのリハーサルをしてみてはいかがでしょうか。理論上はそうではありません。実際に実行してみます。なぜなら、ロールバックが機能しないことがわかった瞬間が、ロールバックが最も必要な瞬間である可能性が高いからです。
AI 時代に向けて更新されたリーン UX の議論
すごいですね

私は、AI が登場した現在でも、リーン UX のどれだけが常に最新の状態で残っているかを知りたいと思っています。リーン UX の古い議論は、アウトプットは結果ではないというものでした。言い換えれば、機能をリリースすることと、顧客がその機能で成功することは同じではありません。 AI時代のバージョンはさらに難しいです。機能を出荷することは、機能を出荷することでさえありません。なぜなら、出荷したものは、展開が完了した後も予期せぬ事態を引き起こし続ける予測不可能なシステム動作の配布だからです。
「完了」はチェックボックスではありません。それはスタンスです。あなたは、公の場で、私はこれらの差異を受け入れることに決め、これらの失敗モードを計画し、予想していなかった差異が現れたときにどのように対応するかをリハーサルしたと言っているのです。
これを正しく理解しているチームは、間違っているチームよりも速いというわけではありません。それらはより意図的であり、最終的には出荷される機能が少なくなります (これは悪いことではありません)。また、リリースごとに生成されるサポート チケットが少なくなり、インシデントごとにブランドへのダメージが少なくなり、導入ごとに信頼性が高まります。これを誤解したチームは、より多く、より速く出荷し、先週完了したと宣言したことがなぜ今壊れているのかを説明するのに残りの週を費やしています。
「完了」が構築文化の言語であるとすれば、「調整済み」は学習文化の言語です。 2番目のものを選択してください。 1 つ目は、応答するための帯域幅を持たないチケットに埋もれてしまうことです。
今日これを試してみてください。ロードマップ上の次の機能を引き上げます。合格基準を声に出して読み上げます。それらがすべて自動販売機についての主張のように聞こえる場合、あなたはすでに何をすべきか知っています。
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
X で共有 (新しいウィンドウで開きます)
×
リンクを友人にメールで送信します (新しいウィンドウで開きます)
電子メール
Reddit で共有 (新しいウィンドウで開きます)
レディット
WhatsApp で共有 (新しいウィンドウで開きます)
ワッツアップ
LinkedIn で共有 (開く)

新しいウィンドウで)
リンクトイン
Jeff Gothelf の書籍は、革新的な洞察を提供し、ユーザー エクスペリエンス、アジャイルな方法論、個人のキャリア戦略といったダイナミックな領域をナビゲートするよう読者を導きます。
リーン vs アジャイル vs デザイン思考
「AI 機能を出荷するときに「完了」とは何を意味しますか」に対する 1 件の回答
Amazon AI トークンの問題は Amazon の問題ではありません。
2026 年 6 月 8 日午後 3 時 48 分
[…] 暗黙的ではありましたが、完全に明らかではありませんでした。仕事は方向性を監督することであり、活動をカウントすることではありません。同じロジックが、AI 機能の「完了」の定義方法にも当てはまります。 […]ではなく、結果の分布を指定します。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
フォローアップコメントを電子メールで通知します。
新しい投稿をメールで通知します。
無料のニュースレターを購読してください
© Copyright 2026 Jeff Gothelf、全著作権所有。 WordPressでデザインされています。主催はPresable。
保存
小計
$0.00
送料と割引はチェックアウト時に計算されます。
カートを見る
チェックアウトに行く
あなたのカートは現在空です!

## Original Extract

AI broke the old definition of “done.” You’re no longer shipping vending machines—you’re shipping probabilistic behavior distributions.

What "done" means when you're shipping AI features
Skip to content
Jeff Gothelf
Sense & Respond Learning World class product management training options for your teams.
0
What “done” means when you’re shipping AI features
Ah, the good old days of sprint reviews.Engineering says they shipped [something]. All tests passed. No P0 bugs. The demo worked like a charm. “Works as designed” had been achieved once again. Sadly, we don’t have the luxury of this simplicity any more (if we ever did).
The definition of “done” has always been the problem. For AI, it has to be completely rewritten.
We built our definitions of done for deterministic software.
Inputs go in, outputs come out, and the output is the same on Tuesday as it was on Monday. That was the whole point. We were building predictable user experiences. QA tested the edges, engineering wrote the unit tests, the spec described the behavior, and “done” was the moment all three lined up. That worked because the software behaved like a vending machine. Put the money in, press the button, and get the soda.
AI features don’t behave like vending machines.
The same prompt produces different outputs across users, sessions, model updates, and contexts your QA team never imagined. The thing you shipped isn’t a fixed, consistent and predictable object . It’s a distribution of behaviors. Some of them you’ve seen and can predict. Others, you haven’t. Most of the surprises live in the ones you haven’t. Your test suite was designed to verify the properties of your code. It was not designed to verify the experience your customer is having at 11pm in Auckland with a question your spec didn’t anticipate.
“All tests pass” is still important and means the same thing it always did. It just no longer means done.
Now, with AI, done is a calibration about an acceptable variance in output and more importantly, behavior, not a binary result about specification adherence. You define done by the range of conditions under which the feature behaves well enough , the failure modes you’ve decided to tolerate, the monitoring you’ve put in place for the ones you haven’t, and the rollback you’re ready to execute and that you’ve actually rehearsed.
There’s a lot in that last sentence so let me break it down into three steps to take to create an AI-ready definition of done.
Step 1: Write acceptance criteria as distributions, not assertions
“When the user asks or inputs X, the system returns Y.” That’s an assertion. It’s verifiable because it’s binary. For an AI feature it’s often meaningless.
Try this framing instead: “For 80% of inputs in category X, the system returns a response that meets quality bar Y; for the remaining 20%, the failure mode is degraded but not embarrassing.” That’s a distribution. It’s harder to write, harder to test, and infinitely more honest about what you’re actually shipping.
You don’t need every criterion to be probabilistic. You need to be honest with yourself and the team about which success outcomes are which and most importantly, stop pretending that the probabilistic ones are deterministic (i.e., fixed and predictable). The teams I see getting this right have two sets of criteria on the same feature:
The deterministic ones for the parts of the system that still behave like vending machines (e.g., auth, billing, navigation)
The distributional (or probabilistic) ones for the parts that don’t (e.g., system feedback, context-specific responses, custom user requests)
Step 2: Build the failure triage capacity into the launch, not after
Traditionally teams shipped the feature, watched the dashboard, and triaged when complaints surfaced at a level worth reacting to. For AI features that workflow is too slow. By the time the complaint reaches engineering, the user has formed an opinion about your product.
To support AI products, write the triage playbook before the AI feature launches. Be clear who owns model-quality issues, UX issues content issues and, perhaps, any PR disasters the model may inadvertently trigger. The launch isn’t just the feature going live. The launch is the moment the team that will field the consequences is actually ready to handle them. This is still not the case in most teams I work with, especially in the enterprise.
You are not done when the feature ships. You are done when the people downstream of the feature know what to do when it misbehaves. And it will, because it’s AI.
Step 3: Define the signal that means “this isn’t done after all” and rehearse the rollback
Every AI feature should ship with a tripwire. You can choose whatever metric makes the most sense for your product and context (and brand). It could be error rate, off-tone complaints per thousand sessions, escalations to human review, hallucination rate flagged by a sample audit, anything that is an indicator of a sub-optimal user experience. Then, set the threshold, and decide before launch what you do when it’s crossed.
Once the wire is tripped, what do you do? Not theoretically, but in practice. You can write it down, sure. How about, instead of just writing it, you rehearse the rollback. Not in theory. Actually run it. Because the moment you find out your rollback won’t work is likely the moment you most need it to.
The Lean UX argument, updated for the AI era
It amazes me how much of Lean UX remains evergreen, even today with AI. The old Lean UX argument was that outputs aren’t outcomes. In other words, shipping the feature isn’t the same as the customer succeeding with it. The AI-era version is harder. Shipping the feature isn’t even shipping the feature, because what you shipped is a distribution of unpredictable, system behaviors that will keep producing surprises after the deploy completes.
Done isn’t a checkbox. It’s a stance. You are saying, in public, I have decided to accept these variances, I have planned for these failure modes, and I have rehearsed how I’ll respond when the variances I didn’t anticipate show up anyway.
The teams getting this right are not faster than the ones getting it wrong. They’re more deliberate, and, ultimately, they ship fewer features (which is not a bad thing). They also generate fewer support tickets per launch, less brand damage per incident, and more confidence per deployment. The teams getting this wrong are shipping more, faster, and spending the rest of their week explaining why the thing they declared done last week is broken now.
If “done” was the language of a build culture, “calibrated” is the language of a learning culture. Pick the second one. The first one is going to bury you in tickets you don’t have the bandwidth to answer.
Try this today. Pull up the next feature on your roadmap. Read the acceptance criteria out loud. If they all sound like assertions about a vending machine, you already know what to do.
Share on Facebook (Opens in new window)
Facebook
Share on X (Opens in new window)
X
Email a link to a friend (Opens in new window)
Email
Share on Reddit (Opens in new window)
Reddit
Share on WhatsApp (Opens in new window)
WhatsApp
Share on LinkedIn (Opens in new window)
LinkedIn
Jeff Gothelf’s books provide transformative insights, guiding readers to navigate the dynamic realms of user experience, agile methodologies, and personal career strategies.
Lean vs. Agile vs. Design Thinking
One response to “What “done” means when you’re shipping AI features”
The Amazon AI-tokens problem isn't an Amazon problem.
June 8, 2026 at 3:48 pm
[…] was implicit but not totally obvious. The job is to supervise direction, not to count activity. The same logic applies to how “done” gets defined for AI features now. You specify an outcome distribution, not a […]
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
Notify me of follow-up comments by email.
Notify me of new posts by email.
Subscribe to my free newsletter
© Copyright 2026 Jeff Gothelf, All Rights Reserved. Designed with WordPress. Hosted by Pressable.
Save
Subtotal
$0.00
Shipping and discounts calculated at checkout.
View my cart
Go to checkout
Your cart is currently empty!
