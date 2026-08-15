---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49308317"
title: "Ask HN: Could Anthropic's watermark be much simpler than we think?"
article_title: ""
author: "CITIZENDOT"
captured_at: "2026-08-15T07:21:52Z"
capture_tool: "hn-digest"
hn_id: 49308317
score: 1
comments: 0
posted_at: "2026-08-15T06:45:10Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: Could Anthropic's watermark be much simpler than we think?

- HN: [49308317](https://news.ycombinator.com/item?id=49308317)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T06:45:10Z

## Translation

タイトル: HN に聞く: Anthropic の透かしは私たちが思っているよりもはるかに単純なのでしょうか?
HN テキスト: Anthropic は最近、Claude 出力に透かしを入れることに取り組んでいると発表しましたが、生成品質には影響しないとも述べています。単なるハッシュフィンガープリンティングなのかどうか疑問に思っています。たとえば、生成されたテキストを取得し、それを重複するチャンクに分割します。「その会社は力強い成長を報告しました...」
「収益の大幅な増加が報告されました...」
「第 2 四半期の収益は大幅に増加しました...」
...
各チャンクをハッシュし、ハッシュを保存します。テキストが検出のために送信されると、同じことを行い、データベースにすでに存在するチャンク ハッシュの数を数えます。誰かがいくつかの単語を編集したとしても、多くの重複する部分が一致する可能性があります。検索自体は特に問題ありません。 256 ビット ハッシュでは 2^256 の空間を扱いますが、実際に保存したハッシュのみを検索します。二分探索では、256 回の反復であらゆるハッシュが検索されます。これは Anthropic の要件も満たします: *トークン生成中に何も変更する必要はありません*。したがって、品質のトレードオフはありません: https://x.com/i/status/2088343978873966687 明らかな疑問は、その規模での誤検知率の動作をどのように処理するかです。これで彼らのアプローチが説明できるでしょうか、それとも私に何か見落としがあるのでしょうか?

## Original Extract

Anthropic recently said they're working on watermarking Claude output, while also saying it won't interfere with generation quality. I'm wondering if is just hash-fingerprinting. For example, take the generated text and split it into overlapping chunks: "The company reported strong growth..."
"reported strong growth in revenue..."
"strong growth in revenue during Q2..."
...
Hash each chunk and store the hashes. When text is submitted for detection, do the same thing and count how many chunk hashes are already in the database. Even if someone edits a few words, many overlapping chunks could still match. The search itself isn't really a problem. With 256-bit hashes you're dealing with a 2^256 space, but you only search the hashes you've actually stored. Binary search would search any hash in 256 iterations. This also satisfies the Anthropic requirements: *nothing needs to be changed during token generation*, so there's no quality tradeoff: https://x.com/i/status/2088343978873966687 The obvious question is how they handle false-positive rate works at their scale. Could this explain their approach, or is there something I'm missing?

