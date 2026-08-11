---
source: "https://www.jrzs.dev/blog/claude-watermarking-ai-text"
hn_url: "https://news.ycombinator.com/item?id=49255920"
title: "Claude Could Watermark Text Without Invisible Characters"
article_title: "How Claude Could Watermark Text Without Invisible Characters | James O'Reilly"
author: "jrhey"
captured_at: "2026-08-11T10:42:11Z"
capture_tool: "hn-digest"
hn_id: 49255920
score: 1
comments: 0
posted_at: "2026-08-11T10:26:57Z"
tags:
  - hacker-news
  - translated
---

# Claude Could Watermark Text Without Invisible Characters

- HN: [49255920](https://news.ycombinator.com/item?id=49255920)
- Source: [www.jrzs.dev](https://www.jrzs.dev/blog/claude-watermarking-ai-text)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T10:26:57Z

## Translation

タイトル: クロードは非表示文字なしでテキストに透かしを入れることができます
記事のタイトル: クロードはどのようにして非表示文字を使用せずにテキストに透かしを入れることができるのか |ジェームズ・オライリー
説明: クロードが生成する人為的な透かしテキストを開始しました。そのようなものが実際にどのように機能するかについては2セントです。

記事本文:
クロードはどのようにして不可視文字を使わずにテキストに透かしを入れることができるのか |ジェームズ・オライリー
ジェームズ・オライリー
ブログ
/
仕事
/
製品
ブログに戻る
2026 年 8 月 11 日
•
3 分で読む
クロード氏は、透明性を確保するため、EU AI法の一環として、生成されるテキストに透かしを入れる予定だと述べた。単純な「AI 生成」ラベルではなく、単語の間に隠された目に見えない文字でもありません。ウォーターマークはテキスト自体に焼き付けられると思います。
LLM はテキストを生成するたびに、考えられる次のトークンの束の中から選択します。おおよそ:
猫 9%
犬 7%
子猫 3%
動物 2%
理論的には、透かしがそれらの選択を促す可能性があります。モデル (およびモデルを作成した人) だけが知っている秘密のパターンに基づいて、特定のトークンに小さなエッジを与えます。単一の選択肢が間違っていることに気づくことはありませんが、数百または数千のトークンをつなぎ合わせると、統計的なフィンガープリントが得られます。基本的に、統計的だが知覚できない単語の選択が透かしです。
Anthropic は具体的にどのようにやっているのかを明らかにしていないので、上記の内容は「クロードがやっているやり方」というよりも、「これがうまくいくかもしれない 1 つの方法」として受け止めてください。彼らのドキュメントもマーケティング版よりもさらに深くは書かれていません。
私が印象に残っているのは、コピーアンドペーストが生き残っているということです。オンラインの人々はすでに「テキスト エディターにコピー/ペーストして、見えない文字を削除すればよい」と考えていますが、信号が単語そのものであれば、幸運を祈ります。 Anthropic は、多少の編集にも耐えられると述べています。
防弾ではないけどね。テキストを十分に変更すると、信号がすぐに崩れてしまう可能性があります。単語の選択において非常に特殊な統計パターンに依存しているのであれば、これは当然のことです。
私が実際に関心があるのはコードですが、それがどのように機能するかはわかりません

そこで働きます。散文には同じことを言う方法がたくさんあり、モデルが気付かないうちにあるトークンを別のトークンよりも静かに好む余地が十分にあります。コードではそんな贅沢はできません。
if user.isAuthenticated {
syncData()を待つ
}
そこにはあまり自由な余地はありません。クロードは、構文が壊れたり動作が変わったりする場合、フィンガープリントのために単にトークンを交換することはできません。 Anthropic 氏は、ウォーターマーケティングにはクロード コードと API の使用が含まれると言っていますが、それがコードのような制約のあるもので実際にどのように機能するのか、私にはわかりません。そこはもっと見ていきたい部分です。
AI を使用した場合と使用しない場合のソフトウェア構築に関する時々のメモ。
スパムはありません。いつでも購読を解除してください。
ボタンダウンを搭載。

## Original Extract

Anthropic started watermarking text Claude generates. My two cents on how something like that could actually work.

How Claude Could Watermark Text Without Invisible Characters | James O'Reilly
James O'Reilly
blog
/
work
/
products
Back to blog
Aug 11, 2026
•
3 min read How Claude Could Watermark Text Without Invisible Characters
Claude said they’ll be watermarking the text it generates as part of the EU AI Act for transparancy. Not a simple “AI generated” label, and not some invisible characters hidden between words either. I think the watermark will be baked into the text itself.
Every time an LLM generates text, it’s choosing between a bunch of possible next tokens. Roughly:
cat 9%
dog 7%
kitten 3%
animal 2%
A watermark could in theory nudge those choices. Give certain tokens a tiny edge based on some secret pattern only the model (and whoever built it) knows about. You’d never notice a single choice being off, but string together a few hundred or a few thousand tokens and you get a statistical fingerprint. Essentially, statisical but imperceptible word choices is the watermark.
Anthropic hasn’t said exactly how they’re doing it, so take the above as “one way this could work” rather than “here’s what Claude does.” Their docs don’t go much deeper than the marketing version either.
The thing that sticks out to me is that it survives copy and paste. People online are already thinking “just copy/paste it into a text editor and delete the invisible characters” but if the signal is the words themselves, good luck. Anthropic says it can survive some editing as well.
It’s not bullet proof though. Modify the text enough and the signal probably falls apart pretty quickly. Which makes sense if you’re relying on a very specific statistical pattern in word selection.
The bit I actually care about is code, and it isn’t clear to me how it’ll work there. Prose has a ton of ways to say the same thing, plenty of room for a model to quietly prefer one token over another without you ever noticing. Code doesn’t give you that luxury.
if user.isAuthenticated {
await syncData()
}
Not a lot of wiggle room there. Claude can’t just swap tokens around for the sake of a fingerprint if doing so breaks the syntax or changes behaviour. Anthropic says watermarketing includes Claude Code and API usage, but how that actually holds up in something as constrained as code, I have no idea. That’s the part I want to see more on.
Occasional notes about building software with and without AI.
No spam. Unsubscribe anytime.
Powered by Buttondown.
