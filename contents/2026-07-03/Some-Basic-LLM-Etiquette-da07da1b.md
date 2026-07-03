---
source: "https://steenbok.space/blog/ai-etiquette/"
hn_url: "https://news.ycombinator.com/item?id=48769662"
title: "Some Basic LLM Etiquette"
article_title: "spork: LLM Etiquette"
author: "sporkl_l"
captured_at: "2026-07-03T01:58:35Z"
capture_tool: "hn-digest"
hn_id: 48769662
score: 1
comments: 0
posted_at: "2026-07-03T01:39:37Z"
tags:
  - hacker-news
  - translated
---

# Some Basic LLM Etiquette

- HN: [48769662](https://news.ycombinator.com/item?id=48769662)
- Source: [steenbok.space](https://steenbok.space/blog/ai-etiquette/)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T01:39:37Z

## Translation

タイトル: LLM の基本的なエチケット
記事タイトル: スポーク: LLM エチケット
説明: AI ユーザーは、他の人々とのやり取りの補助として AI を使用する場合、より敬意を払う必要があります。

記事本文:
あなたが私と同じように判断力の強い女王であれば、おそらく、うんざりすることを防ぐために、LinkedIn フィードをスクロールすることを避けているでしょう。このサイトのコンテンツの 90% (引用しないでください) は常に、不誠実で絵文字だらけの、フォーム コンテンツの奇妙なコレクションであったと思います。私は常々、精神に何らかの病理学的障害がある場合を除けば、こうした投稿の著者にとって、このような戯言を書くのは骨の折れる仕事に違いないと思っている。 AI の出現以来、この不誠実なコンテンツの膨大な部分が自動化されました。当然のことですが、ほとんどの人は実際にはこれらの投稿を価値の高いものとは考えていないので、なぜそれをマシンにオフロードしないのでしょうか?
LLM ユーザーが相互に対話する際にこれらのツールをどのように使用するかについてのルールを策定し始める時期が来たと思います。今のところ、私は彼らが再び独自の LinkedIn 投稿を書き始めることを期待していません。それはいいことですが、AI がすべてを書き始めるずっと前に、その場所は失われた大義でした。そうは言っても、私は日常のやりとりの中で LLM チャットボットの使用がより頻繁に発生していることに気づきました。それは芽を摘み取る必要があるものです。
仮定の状況を見てみましょう。
Holcomb 氏と Truman 氏はどちらも、Middling Millennial Software Company のデータ エンジニアです。マーケティング部門の関係者である Eudora は、会社の BI ツールに表示される指標の一部が CRM で取得するレポートと一致していないという懸念を Holcomb に提起しました。ホルコム氏はダッシュボードを見て、DBT モデルを搭載していると判断しました。 Holcomb はあまりコードを読むのが好きではないので、Claude を起動し、モデルをレビューするために送信します。
クロードが戻ってきて、モデルのロジックに重大な欠陥があり、すぐに修正する必要があると報告しました。ホルコム氏はこれを非常に憂慮すべきだと感じ、彼に与えられた返答をコピーしました。

クロード、トルーマンと緩んだスレッドに貼り付けます。トルーマンは、メッセージに散りばめられた対照的な修辞的枠組みと全角ダッシュに非常に腹を立てている。ホルコムが問題の理解やメッセージの作成に何の努力もしていないことは、あからさまに明らかです。
トルーマンはこの問題を調査し続け、彼の疑い通り、クロードが間違っていたことを発見した。モデルのロジックには問題ありません。同じロジックが BI ツールと CRM のデータに電力を供給します。クロードには、CRM がデータをどこから取得するのかというコンテキストがまったくありませんでした。実はこれはタイミングの問題でした。 BI ツールはデータ ウェアハウスにライブ クエリを実行しますが、CRM はバッチ ジョブによって更新されます。トルーマンはユードラに連絡を取り、遅れについて知らせ、彼女は満足した。
この仮説の中で、ホルコムは、LLM ユーザーに守ってほしいいくつかのルールを破りました。
LLM の出力を検証します。
自分の言葉で情報を伝えることに努めてください。
コピーして貼り付ける必要がある場合は、出力が LLM からのものであることを示します。
LLM が頻繁に間違っていることは誰もが知っています。 LLM の出力を検証する時間を取らないと、相手に 2 つのことのいずれかを伝えていることになります。つまり、LLM が何を言っているのか実際にはわかっていないか、共有している情報が正しいかどうかを確認するのが面倒であるかのどちらかです。
自分が何を言っているのかを理解していることを示すために、提供された出力を自分の言葉で置き換えることができます。最初のステップを通過したことを証明できるだけでなく、AI 応答の迷惑な兆候をすべて取り除くこともできます。
さて、何のことを言っているのかよく分からない場合は、そう言えば大丈夫です。コンテンツを自分のものとして伝えようとしないでください。またしても、あなたがルーブのように見えることになるからです。 「ねえ、クロードにxyzについて尋ねたら、これが返ってきたものです。」と単純に言うとします。

それが何を意味するか知っていますか？チャットしている相手の知性を尊重していることを伝える返信を貼り付ける前に。
AI ユーザーとして、これらは自分自身に耐えられないように見せるためにできるいくつかのことにすぎません。理想的には、AI をまったく使用せず、クソドキュメントを参照するか、IDE の機能を使用してコードベースを操作するだけです。でも、あまり多くは聞きたくないので、これだけは言っておきます。あなたはとても特別な少年で、どんなことをしてもあなたが悪者になることはありません。

## Original Extract

AI users should be more respectful when they use it as an aide to their interactions with other people.

If you are as much of a judgy queen as I am, you probably avoid scrolling through your LinkedIn feed as a prophylactic against cringe. I’d say 90% (don’t quote me) of the content on this site has always been this bizarre collection of disingenuous, emoji-laden, form-content. I’ve always thought that, barring some pathological disorder in their psyche, writing such drivel must be an exhausting task for the authors of these posts. Since the advent of AI, a vast swath of this disingenuous content has become automated. It makes sense, most folks don’t actually view these posts as high-value so why not offload it to the machine?
I think it is time that LLM users start developing rules around how they use those tools when interacting with each other. Now, I’m not expecting them to start writing their own LinkedIn posts again. Though it would be nice, that place was a lost cause well before AI started writing everything. That being said, I have noticed the use of LLM chatbots more frequently cropping up in my day-to-day interactions. That is something that needs to be nipped in the bud.
Lets walk through a hypothetical situation.
Holcomb and Truman are both data engineers at Middling Millennial Software Company. Eudora, a stakeholder from the marketing department raises a concern to Holcomb that some metrics displayed in the company’s BI tool aren’t aligned with the reporting she gets in the CRM. Holcomb looks at the dashboard and determines it is powered by a DBT model. Holcomb doesn’t like to read very much code, so he spins up Claude and sends it off to review the model.
Claude comes back and reports a critical flaw in the model’s logic that must be remediated immediately. Holcomb finds this very alarming, copies the response given to him by Claude, and pastes into a slack thread with Truman. Truman is very annoyed by the contrastive rhetorical framing and em-dashes littering the message. It is blatantly obvious that Holcomb put no effort into understanding the problem or writing the message.
Truman goes on to investigate the issue and finds, as he suspected, that Claude was wrong. There is no problem in the model logic. The same logic powers the data in the BI tool and the CRM. Claude simply didn’t have the context of where the CRM gets it data. This was actually a timing issue. The BI tool live queries the data warehouse whereas the CRM gets updated by a batch job. Truman reached out to Eudora and let her know about the delay, and she was satisfied.
In this hypothetical, Holcomb broke a few rules that I wish LLM users would follow:
Validate the output of the LLM.
Put effort into relaying the information in your own words.
If you must copy and paste, denote that the output came from an LLM.
Everyone knows that LLMs are frequently wrong. When you don’t take the time to validate the output of the LLM you tell the other person you’re interacting with one of two things: either you don’t actually know what it is talking about or you are too lazy to make sure the information you are sharing is correct.
To show you know what you’re talking about you can put the provided output in your own words. Not only will you prove you went through the first step, but you’re also going to get rid of all the annoying telltale signs of an AI response.
Now, if you don’t really know what it is talking about, you can just say that. Don’t try and pass the content off as your own because, once again, it makes you look like a rube. If you simply say, “Hey, I asked Claude about xyz and this is what it came back with. Do you know what it means?” before you paste in the response that will communicate that you respect the intelligence of the person you’re chatting with.
As an AI-user, these are just a few things you can do to make yourself seem less insufferable. Ideally, you wouldn’t be using AI at all and you would just consult the fucking documentation or use the features of your IDE to navigate the codebase. But hey, I don’t want to ask too much of you, so I’ll leave you with this: you’re a very special boy and nothing you could ever do would make you a bad person.
