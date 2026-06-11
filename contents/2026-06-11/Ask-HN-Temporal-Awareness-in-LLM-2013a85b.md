---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48487573"
title: "Ask HN: Temporal Awareness in LLM?"
article_title: ""
author: "Pamar"
captured_at: "2026-06-11T10:36:58Z"
capture_tool: "hn-digest"
hn_id: 48487573
score: 1
comments: 0
posted_at: "2026-06-11T07:55:43Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: Temporal Awareness in LLM?

- HN: [48487573](https://news.ycombinator.com/item?id=48487573)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T07:55:43Z

## Translation

タイトル: HN に聞く: LLM における時間的認識?
HN テキスト: ChatGPT に一種の「機能リクエスト」を送信しました (ただし、私の質問はほとんどの LLM に一般的に当てはまります)。以下は ChatGPT 自体によって書かれたテキストです: 機能リクエスト: 会話/プロジェクトごとのオプションの時間認識。モデルがメッセージ間の相対的な経過時間を受信できるようにします (例: +5 分、+4 時間、+2 日)。これは、経過時間自体が意味論的な意味を持つ、長時間にわたる物語、日記、プロジェクト追跡、関係指向の会話に特に役立ちます。私は主に日記に使っていますが、1 時間操作して次のメッセージが来ても、システムはまったく認識していないのです... 6 時間後、その間に何かが起こりました。もしかしたら居眠りしてしまったのかもしれない。あるいは、もっと予兆的なことが起こったのかもしれません。どのクライアントでも TZ で時刻を送信するか、相対時間のみが重要であると判断して他のメカニズムで時刻を提供できることを考慮すると、これを実装しなかったのは意識的な決定だったと思います。では、その理由は何でしょうか? - プライバシーに関する懸念はありますか?
- 効率（Google から得られるものよりも完全な回答を得るために、またはバイブコーディングを行うために主に LLM を使用する場合...その場合、それはほとんど役に立ちません） あるいは、この時間機能が回答にあまりにも大きな影響を与えるか、より多くのリソースを消費するためかもしれません（たとえこれが小さくても、すべてのユーザーからのすべてのメッセージのペイロードまたはトークンの使用量がわずかに大きい場合、負担になる可能性があると確信しています）。両方の理由から、これをユーザー プロファイル/プロジェクト/スレッド レベルで構成可能にします。しかし、おそらく私は他の要素を見逃しているかもしれません（これにはすでに72回回答しており、まだ木曜日であることも含まれます）、それが私がここで質問している理由です。

## Original Extract

I have just sent a sort of "feature request" to ChatGPT (but my question is generally applicable to most LLMs). Here is the text, written by ChatGPT itself: Feature request: Optional temporal awareness per conversation/project. Allow the model to receive relative elapsed time between messages (e.g. +5 min, +4 h, +2 days). This is particularly useful for long-running narrative, journaling, project-tracking, and relationship-oriented conversations where elapsed time itself carries semantic meaning. I mostly use it for journaling, and the fact that the system is completely unaware that if I have interacted with it for 1 hour and the next message comes ... 6 hours later, something happened in between. Maybe I dozed off. Or maybe something more portentous has happened. I suppose that not implementing this was a conscious decision, considering that any client can just send either the time with TZ or decide that only relative time is important and provide it by some other mechanism. So what could be the reason? - privacy concerns?
- efficiency (if I use an LLM mostly to get more complete answers than I can get from Google, or to do vibecoding... then it is pretty much useless) Or maybe it is because this time feature would impact the answers too much/consume more resources (even if this is tiny, if EVERY MESSAGE FROM EVERY USER has a slightly bigger payload or token usage, I am sure it could become a burden). For both reason I would make this configurable at the user profile/project/thread level. But maybe I am missing other elements (including: we have already answered this 72 times and it is only thursday) and this is why I am asking here.

