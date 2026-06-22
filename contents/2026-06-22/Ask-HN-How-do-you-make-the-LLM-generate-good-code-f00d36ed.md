---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48637538"
title: "Ask HN: How do you make the LLM generate good code?"
article_title: ""
author: "bjourne"
captured_at: "2026-06-22T23:28:56Z"
capture_tool: "hn-digest"
hn_id: 48637538
score: 1
comments: 0
posted_at: "2026-06-22T22:51:35Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: How do you make the LLM generate good code?

- HN: [48637538](https://news.ycombinator.com/item?id=48637538)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T22:51:35Z

## Translation

タイトル: HN に質問: LLM に適切なコードを生成させるにはどうすればよいですか?
HN テキスト: 最近、音楽コレクションの整理に役立つボットを入手できないかどうかを知りたいと思っていました。具体的には、曲のオーディオ ファイルへのパスを取得し、その曲の正規 (つまり、最初の) アルバムの名前を返す Python 関数をコード化したいと考えています。リマスター、ライブ録音、ブートレッグなどを無視すると、これは簡単なヒューリスティックな全射マッピングです: スモーク・オン・ザ・ウォーター -> マシーン・ヘッド、カム・アズ・ユー・アー -> ネヴァーマインド、フールズ・ゴールド -> ザ・ストーン・ローゼズ。私は MusicBrainz API を使用してこの関数をすばやくコーディングできます。他のほとんどの HN 読者も同様にコーディングできると思います。しかし、何時間も経っても、ボットにどのように指示しても、ボットに実行させることができません。これは ChatGPT が生成するものです --- 他のボットも同様のゴミで応答します: https://chatgpt.com/share/6a396ea9-44d4-83eb-bdfd-216dfcc87e99 彼らは、問題がいかに曖昧であるか (曖昧ではない)、または MusicBrainz がどのようにデータを持っていないのか (存在します) について不満を述べています。ボットに間違ったプロンプトを送っているのでしょうか、それとも十分強力なコーディング モデルを使用していませんか?結果は非常に期待外れです。

## Original Extract

Lately, I wanted to see if I could get the bot to help me organize my music collection. Specifically, I want it to code a Python function that takes the path to an audio file of a song and returns the name of the canonical (i.e., first) album of that song. If one ignores remasters, live recordings, bootlegs, etc., this is an easy heuristic surjective mapping: Smoke on the Water -> Machine Head, Come as You Are -> Nevermind, Fools Gold -> The Stone Roses. I can code this function quickly using the MusicBrainz API and I believe most other HN readers can too. Yet many hours later, no matter how I prompt the bot I can't get it to do it. This is what ChatGPT generates---other bots reply with similar garbage: https://chatgpt.com/share/6a396ea9-44d4-83eb-bdfd-216dfcc87e99 They complain about how the problem is ambiguous (is not) or how MusicBrainz doesn't have the data (it does). Am I prompting the bot incorrectly or not using powerful enough coding models? The results are very underwhelming.

