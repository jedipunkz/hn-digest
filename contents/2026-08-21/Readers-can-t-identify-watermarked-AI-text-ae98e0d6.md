---
source: "https://www.seangoedecke.com/readers-cant-identify-watermarked-ai-text/"
hn_url: "https://news.ycombinator.com/item?id=49392819"
title: "Readers can't identify watermarked AI text"
article_title: "Readers can't identify watermarked AI text"
image: ""
author: "dmarto"
captured_at: "2026-08-21T20:14:43Z"
capture_tool: "hn-digest"
hn_id: 49392819
score: 5
comments: 0
posted_at: "2026-08-21T19:30:09Z"
tags:
  - hacker-news
  - translated
---

# Readers can't identify watermarked AI text

- HN: [49392819](https://news.ycombinator.com/item?id=49392819)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/readers-cant-identify-watermarked-ai-text/)
- Score: 5
- Comments: 0
- Posted: 2026-08-21T19:30:09Z

## Translation

タイトル: 読者は透かし入りの AI テキストを識別できない
説明: ここ数週間、私は AI 透かし入れについてみんなが間違っていると不満を言ってきました。AI 透かし入れは実際には反消費者ではないし、出力も何の役にも立ちません…

記事本文:
読者は透かし入りの AI テキストを識別できない
読者は透かし入りの AI テキストを識別できない
ここ数週間、私は AI 透かし入れについてみんなが間違っていると不満を言い続けてきました。AI 透かし入れは実際には反消費者ではないし、出力を悪化させるわけでもありません。透かしを入れた論文はこれが真実であることを証明していますが、実際にテストしてみるのも面白いかもしれないと思いました。同じプロンプトに対する透かし入りの回答と透かしなしの回答の例を与えられた場合、読者はどちらがどれであるかを判断できますか?
それを知るために、読者にクイズを出題する静的サイト 2 https://sgoedecke.github.io/watermark-quiz/ を作成しました。レンタルした H200 で Qwen3-30B-A3B-Instruct-2507 を使用して 30 の応答を生成しました。質問ごとに 3 つの応答があり、そのうちの 1 つは秘密裏に SynthID-Text の透かしが入っていました。レンタルした GPU の費用は約 2 ドルでした。結果を測定するために、スコアごとにユーザーを異なるページに送り、分析でページごとの訪問者数を集計しました 3 。誰かがそうすることに十分な注意を払っていれば、これは簡単になりすますことができますが、カジュアルなテストであれば許容できると思います。
私がクイズに参加した最初のトラフィック (参加者 278 人) では、次のような少し不可解な結果が得られました。
純粋にランダムに選択した場合、平均スコアは 3.33/10 になります。ただし、ここでの平均スコアは 3.92 です。予想どおり、確かに 3/10 付近にスパイクがありますが、6/10 にも 2 番目の奇妙なスパイクもあります。何故ですか？ SynthID の応答は 10 問中 6 問で選択肢 A だったことが判明したため、すべての質問に対して最初の回答を選択したユーザーは 6/10 を獲得することになります。おっと。
質問を再シャッフルしたところ、次のような結果が得られました。
現在、平均は 3.4/10 で、予想される 3.333 にかなり近づきます。 6 のあたりに急増はありません。質問をシャッフルした後、クイズに参加したのは 73 人だけでした。ほとんどの人がそれを見てクイズに答えてくれました。

私がそれを LinkedIn と Hacker News に投稿した直後でしたが、これまでの結果を考慮すると、人々が単に無作為に推測しているだけであると確信するには、それでも十分だと思います。
つまり、人々は AI 透かしの存在を識別することができません。明らかに、これは厳密には科学的な研究ではありませんでしたが、それでもかなり示唆に富んでいます。もし透かしが実際にモデルが決して選ばないようなランダムな単語を選んでいるのなら、3 つの応答を並べて表示することで、どれが奇妙な透かしの入った道をたどったかがわかることもありますよね?また、このようなものが説得力のあるツールとして役立つことを願っています。透かしがどのような影響を与えるか心配していて、数学的な説明に直感が動かされない場合は、透かしの入った応答と透かしの入っていない応答を読んで、実際には品質に違いがないことを納得できるかもしれません。
その理由を一文で説明すると、AI モデルはすでに少数の上位トークンからランダムに選択されており、ウォーターマークはそのランダムな選択を、同等の「ランダム」でありながら予測可能なバイアスに置き換えるだけです。簡単な例として、「上位 3 つのトークンからランダムに選択する」代わりに、「前の 10 個のトークン内の文字を数えて、mod 3 を取得し、そのトークンを選択する」ことができます。
雰囲気からのいくつかのメモ: GPT-5.6-Sol は、ページ全体に無関係なテキストを配置し、私はそれを削除する必要があり、今では非常によく知られているスタイルを選択したので、私はそれを取り除かなければなりませんでした。また、私が手作業で構築するクロスリンクされた純粋な HTML の代わりに、ある種の奇妙な Javascript 駆動の静的サイトを構築しました。所要時間は 1 時間ほどでした（実際の作業は 10 分程度でした）。
Umami、PikaPods でホストされています。私のブログでは、JS ベースの分析では 50% 以上のエラーが発生するため、Netlify 分析にも料金を払っています。

技術的なユーザーですが、このようなものにはうまみは問題ありません。
この投稿を気に入っていただけた場合は、私の新しい投稿に関する更新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
AI テキスト透かしは大した問題ではない
クロード モデルの出力に隠し透かしを含める予定であるという Anthropic の最近の発表について、人々はかなり不満を抱いています。これは人類モデルからの大量流出につながるでしょうか?ウォーターマークの導入はユーザーにとって有意義な変化となるでしょうか?
いいえ、AI テキスト透かしは大したことではありません。テキストが悪化するわけでもなく、AI 出力が実際に検出可能になるわけでもなく、ユーザーのプライバシーを侵害するわけでもなく、関係なく、2027 年までに誰もがそうすることになるでしょう。
続きを読む...
購読する │ About │ ポッドキャスト │ 人気 │ タグ │ RSS

## Original Extract

In the last few weeks, I’ve been complaining that everyone is wrong about AI watermarking: it isn’t really anti-consumer and it doesn’t make the outputs any…

Readers can't identify watermarked AI text sean goedecke
Readers can't identify watermarked AI text
In the last few weeks, I’ve been complaining that everyone is wrong about AI watermarking: it isn’t really anti-consumer and it doesn’t make the outputs any worse. The watermarking papers demonstrate 1 that this is true, but I thought it might be interesting to put it to a practical test. Given examples of watermarked and unwatermarked answers to the same prompt, could readers tell which is which?
To find out, I vibed up 2 https://sgoedecke.github.io/watermark-quiz/ , a static site that quizzes readers. I used Qwen3-30B-A3B-Instruct-2507 on a rented H200 to generate thirty responses: three responses per question, one of which was secretly watermarked with SynthID-Text. The rented GPU cost around two dollars. To measure results, I just sent users to a different page for each score, and aggregated visitors-per-page in my analytics 3 . This would be easily spoofable if anyone cared enough to do so, but for a casual test I think it’s acceptable.
The first round of traffic I got to the quiz (278 participants) had these slightly puzzling results:
Pure random choice would lead to an average score of 3.33/10. However, the mean score here is 3.92. There is indeed a spike around 3/10, as expected, but there’s also a second weird spike at 6/10. Why is that? It turned out that the SynthID response was option A in six of the ten questions, so users who just selected the first answer for every question would get 6/10. Oops.
I re-shuffled the questions and got these results:
Now the mean is 3.4/10, much closer to the expected 3.333. There’s no spike around 6. We only had 73 people take the quiz after I shuffled the questions — most people saw it and took it immediately after I posted it to my LinkedIn and Hacker News — but given the previous results, I think that’s still enough to feel confident that people were just guessing randomly.
So no, people can’t identify the presence of AI watermarks . Obviously this wasn’t exactly a scientific study, but it’s still pretty suggestive. If watermarks were really choosing random words that the model would never pick, you’d be able to sometimes tell from three side-by-side responses which one went down the weird watermarked road, right? I also hope that something like this can serve as a persuasive tool: if you’re worrying about what impact watermarking is going to have, and your intuition is unmoved by the mathematical explanations, having a read of the watermarked and unwatermarked responses might convince you that there’s really no difference in quality.
The one-sentence explanation for why is that AI models already randomly select from a handful of top tokens, and watermarking just replaces that random choice with a bias that is predictable while still being equivalently “random”: as a simple example, instead of “pick randomly from the top three tokens”, you could do “count the letters in the previous ten tokens, take mod three, then pick that token”.
Some notes from the vibing: GPT-5.6-Sol put extraneous text all over the page I had to get it to remove, it chose the now-very-recognizable styling that I had to rip out, and it built some kind of weird Javascript-driven static site instead of just the cross-linked pure HTML thing I would have built by hand. It took me about an hour (although I did maybe ten minutes of actual work).
Umami, hosted on PikaPods. For my blog, I do also pay for Netlify analytics because I find JS-based analytics misses >50% of technical users, but for stuff like this Umami is fine.
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
AI text watermarking is not a big deal
People are pretty unhappy about Anthropic’s recent announcement that they’re planning to include a hidden watermark in Claude model outputs. Will this lead to a mass exodus from Anthropic models? Will the introduction of watermarking be a meaningful change for users?
No. AI text watermarking is not a big deal. It doesn’t make the text worse, it doesn’t make AI outputs more detectable in practice, it doesn’t violate user privacy, and everyone’s going to be doing it by 2027 regardless.
Continue reading...
subscribe │ about │ podcasts │ popular │ tags │ rss
