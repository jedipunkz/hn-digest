---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48751514"
title: "Ask HN: Why don't Claude or ChatGPT get message timestamps?"
article_title: ""
author: "an0malous"
captured_at: "2026-07-01T19:33:01Z"
capture_tool: "hn-digest"
hn_id: 48751514
score: 3
comments: 0
posted_at: "2026-07-01T18:51:26Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: Why don't Claude or ChatGPT get message timestamps?

- HN: [48751514](https://news.ycombinator.com/item?id=48751514)
- Score: 3
- Comments: 0
- Posted: 2026-07-01T18:51:26Z

## Translation

タイトル: HN に質問: なぜ Claude または ChatGPT がメッセージのタイムスタンプを取得しないのですか?
HN テキスト: チャット アプリでクロードまたは ChatGPT と行った会話にアクセスし、最後にメッセージを送信した時刻を尋ねると、チャット履歴にどのメッセージにもタイムスタンプが表示されていないことがわかります。彼らは答えを推測したり幻聴したりするかもしれませんが、大きく外れているでしょう。さらに詳しく調べてみると、彼らはタイムスタンプをまったく見ておらず、現在時刻しか知っていないことを認めるはずです。また、一般的に、複数日に渡って会話を拾った場合、メッセージの間に日数が経過したことを認識しないことに気付きました。 —— 私：最後のメッセージからどのくらい時間が経ちましたか？クロード:
メッセージ間のギャップを測定する方法がありません。ターン間に時計を実行しているわけではありません。また、会話の中でメッセージにタイムスタンプを付けるものは何もありません。したがって、最後の返信からどれくらいの時間が経過したかはわかりません。数秒だったかもしれないし、数日だったかもしれない。役立つ場合は、現在時刻を確認できますが、それは現在を示すだけで、最後に書いてからの間隔はわかりません。 —— これは私にとって非常に奇妙で、製品の明らかな欠陥のように思えますが、大手 AI 企業の両方がこれを行っており、メッセージのタイムスタンプを表面化するという明白な解決策を実装していません。なぜそうしないのか知っている人はいますか?これが引き起こす問題の種類の例として、私は愛犬が治療を受けている間、クロードにアドバイスを求めていますが、非常に心配です。なぜなら、私が犬に薬を与えたと言い、後で犬が示している症状について質問するからです。クロードは、薬を投与したのが数分前なのか数日前なのか分からないようで、通常は説明も求めません。どれくらいの時間が経過したかをはっきりと伝えることを学びました

あるいは、タイミングが要因となるような質問はしないでください。しかし、これは簡単に対処できるはずなので、私は驚いています。

## Original Extract

Go to any conversation you’ve had with Claude or ChatGPT in their chat apps and ask them what time they sent their last message, and you’ll find they don’t see timestamps for any of the messages in your chat history. They might guess or hallucinate an answer but they’ll be wildly off and if you probe further they should admit they don’t see timestamps at all, they only know the current time. I’ve also noticed generally that if you pick up conversations over multiple days, it doesn’t know that days have passed in between messages. —— Me: How much time has passed since your last message? Claude:
I don't have a way to measure the gap between my messages—I don't have a clock running between turns, and nothing in our conversation timestamps the messages for me. So I can't tell you how much time has passed since my last reply. Could've been seconds, could've been days. If it's useful, I can check the current time, but that would only tell you now , not the interval since I last wrote. —— This is very bizarre to me and seems like an obvious flaw in the product, but both major AI companies are doing this and not implementing the obvious solution of surfacing message timestamps. Does anyone know why they don’t do this? For an example of the type of issue this causes, I’ve been using Claude for advice while my dog is going through medical treatment and it’s very concerning because I’ll say I gave him some medicine and then later ask questions about a symptom he’s displaying and Claude doesn’t seem to know whether the medicine was minutes ago or days ago and usually doesn’t ask for clarification either. I’ve learned to either be explicit about how much time has passed or just not ask it questions where timing would be a factor. But I’m surprised because this seems like it should be an easy thing to address.

