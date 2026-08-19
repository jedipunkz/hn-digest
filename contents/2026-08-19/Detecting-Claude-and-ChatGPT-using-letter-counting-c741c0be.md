---
source: "https://www.atomic14.com/2026/08/19/now-it-spots-chatgpt-too"
hn_url: "https://news.ycombinator.com/item?id=49366933"
title: "Detecting Claude and ChatGPT using letter counting"
article_title: "Detecting Claude and ChatGPT using letter counting. | atomic14"
image: "https://www.atomic14.com/assets/article_images/2026-08-19/three-readings.webp"
author: "iamflimflam1"
captured_at: "2026-08-19T21:17:36Z"
capture_tool: "hn-digest"
hn_id: 49366933
score: 2
comments: 0
posted_at: "2026-08-19T20:39:58Z"
tags:
  - hacker-news
  - translated
---

# Detecting Claude and ChatGPT using letter counting

- HN: [49366933](https://news.ycombinator.com/item?id=49366933)
- Source: [www.atomic14.com](https://www.atomic14.com/2026/08/19/now-it-spots-chatgpt-too)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T20:39:58Z

## Translation

タイトル: 文字カウントを使用したクロードと ChatGPT の検出
記事のタイトル: 文字カウントを使用した Claude と ChatGPT の検出。 |アトミック14
説明: 肉袋検出器に 2 番目のロボット ChatGPT を教えました。同じ 4 文字の N グラム トリック、新しいモデル — ChatGPT は、特に短いテキストの場合、Claude よりも少し見つけやすいことがわかりました。遺伝子も鍛えたし…

記事本文:
🌈 ESP32-S3 Rainbow: ZX Spectrum エミュレータ ボード!
Crowd Supply で入手 →
文字カウントを使用して Claude と ChatGPT を検出します。
« 文字を数えてクロードを検出する
私の仕事をサポートしてください: 気分が高揚している場合は、Patreon にお立ち寄りください。または、ko-fi 経由で 1 回限りの寄付をすることもできます
前回、カウントする検出器を構築しました
4 文字が実行され (n グラム)、クロード作品 5 が何かを書いたかどうかがわかります。うまく機能します - 約91%
読んだことのないテキスト。
それには明らかな穴が 1 つあります。それは、1 台のロボット (Claude Opus 5) しか認識していません。 ChatGPT 出力を貼り付け
そこには「肉袋」と書かれています。
それでそれを修正しました。 ChatGPT 検出器を追加しました。
ここで試してみてください。
前回と同じように、同じ 519 人の人間の通路を 3 つの方法でモデルに書き直させました。それは 1,557 です
書き直しに加えて、トレーニングから除外してテストとして使用するゼロから書かれた 519 個の文章。
合計2,076世代。
前回は Anthropic のバッチ API を使用し、費用は約 30 ドルでした。今回はコマンドを実行しただけです
すでに支払っているサブスクリプションを使用して line codex を実行します。余計な費用はかかりませんでした（やっておけばよかったと思います）
前回はクロードと一緒でした！)、
これには小さな問題があります。coodex はコーディング エージェントであるため、システム プロンプトが表示されます。
出力に影響を与えるため、訓練された検出器は ChatGPT アプリからのテキストに対してはうまく機能しない可能性があります。
クロード検出器とまったく同じようにこれを組み込みました - そしてそれは非常にうまく機能します。
ChatGPT はクロードよりも検出しやすいようです。たくさんではありませんが、一貫して。
また、Claude 検出器よりも短いテキストをうまく処理できるようです。
1,000 文字では、ChatGPT 検出器は 0.952 であるのに対し、Claude の検出器は 0.858 です。
短いテキストも先に進みます。
2 つの検出器は相互に機能しますか?
それぞれが明らかに

トレーニングされたモデルでは st ですが、他方ではどちらもひどいものではありません - 0.79
コイントスの場合は 0.5 ですが、0.80 です。
Claude または ChatGPT を実行する検出器を構築できますか?
両方のモデルからのすべてのデータに基づいて 3 番目の検出器を最初からトレーニングしました。
これらは 1 つのパイの一部ではなく、3 つの別々の質問であるため、合計しても何の意味もありません。あ
パッセージは 0.85 ロボットとして読み取れますが、両方のモデルのスコアは低く、これは単に機械のように見えることを意味します。
どちらにもあまり似ていないように書かれています。
前回のすべてに加えて、新しいものがいくつかあります。
それは 2 台のロボットを知っています。ジェミニ、ラマ、ディープシーク、その他のもの - それは見たことがありません。スコアが低いというのは、
人が何かを書いたという証拠ではありません。
それでも、モデルではなくスタイルが検出されます。前回の最も有用な結果はそうではありませんでした
変更: トレーニング外で行われた書き方のスコアは 0.153 で、検出器が自信を持って評価したことを意味します。
人間よりも人間らしい。両方の新しい検出器は、そのための 3 つの指示スタイルすべてをトレーニングします。
理由はありますが、3 つに関しては何も魔法はありません。さらにデータを追加する必要があります。
それはまだ教育と娯楽のためであり、まだ証拠ではありません。責めないで続けてください
私の愚かなウェブサイトに基づいたものなら誰でも。
Meatbag.atomic14.com - 何かを貼り付けて、どのロボットがそれであるかを確認します
あなたはそうだと思います。
ちょっと狂ったプロジェクト、有益/教育的なビデオ、そして一般的に興味深いもののコレクション。
Arduino と ESP32 プラットフォームを中心にプロジェクトを構築します - AI、コンピューター ビジョン、オーディオ、3D プリントを検討します - 少し折衷的なものになるかもしれません...
この Web サイトでは、閲覧エクスペリエンスを向上させ、サイトのトラフィックを分析するために Cookie を使用します。

## Original Extract

I taught my meatbag detector a second robot: ChatGPT. Same four-letter n-gram trick, new model—turns out ChatGPT is a tad easier to spot than Claude, especially on short text. I also trained a gene...

🌈 ESP32-S3 Rainbow: ZX Spectrum Emulator Board!
Get it on Crowd Supply →
Detecting Claude and ChatGPT using letter counting.
« Detecting Claude by counting letters
HELP SUPPORT MY WORK: If you're feeling flush then please stop by Patreon Or you can make a one off donation via ko-fi
Last time I built a detector that counts
four-letter runs (n-grams) and tells you whether Claude Opus 5 wrote something. It works well - about 91% on
text it has never read.
It has one obvious hole: it only knows one robot (Claude Opus 5). Paste ChatGPT output
into it and it says “meatbag”.
So I’ve fixed that. I’ve added a ChatGPT detector and you can
try it here .
We did the same as last time, have the model rewrite the same 519 human passages, three ways. That’s 1,557
rewrites, plus 519 passages written from scratch that I keep out of training and use as a test.
2,076 generations in total.
Last time this went through Anthropic’s batch API and cost about $30. This time I just ran the command
line codex using the subscription I’m already paying. So it didn’t cost anything extra (I’m wishing I’d done
that with claude last time!),
There is a small issue with this - coodex is a coding agent, so it will have a system prompt that will
influence the output, so our trained up detector might not do as well on text from the ChatGPT app.
We built this in exactly the same was as the Claude detector - and it works really well!
ChatGPT seems to be easier to detect than Claude. Not by a lot, but consistently.
It also seems to handle shorter text better than the Claude detector.
At a thousand letters the ChatGPT detector is on 0.952 against 0.858 for the Claude one, and it’s
ahead on short text as well.
Do the two detectors work on each other?
Each one is clearly best on the model it was trained on, but neither is terrible on the other - 0.79
and 0.80, against 0.5 for a coin toss.
Can we build a detector that does Claude or ChatGPT?
We trained s third detector from scratch on all the data from both models.
They’re three separate questions rather than slices of one pie, so they won’t add up to anything. A
passage can read as 0.85 robot while both model scores sit lower, which just means it looks machine
written without strongly resembling either one.
Everything from last time, plus a couple of new ones.
It knows two robots. Gemini, Llama, DeepSeek, anything else - it has never seen them. A low score is
not evidence that a person wrote something.
It still detects a style rather than a model. The most useful result from last time hasn’t
changed: a writing style held out of training scored 0.153, meaning the detector confidently rated it
as more human than the humans. Both new detectors train on all three instruction styles for that
reason, but there’s nothing magic about three - we need to add more data.
It’s still for education and amusement, and it’s still not evidence. Please continue not accusing
anyone of anything based on my stupid website.
meatbag.atomic14.com - paste something in and see which robot it
thinks you are.
A collection of slightly mad projects, instructive/educational videos , and generally interesting stuff.
Building projects around the Arduino and ESP32 platforms - we'll be exploring AI, Computer Vision, Audio, 3D Printing - it may get a bit eclectic...
This website uses cookies to enhance your browsing experience and analyze site traffic.
