---
source: "https://www.autodidacts.io/usable-local-ai-handwriting-recognition/"
hn_url: "https://news.ycombinator.com/item?id=48370959"
title: "Local AI handwriting recognition is usable now"
article_title: "PSA: local AI handwriting recognition is usable now — The Autodidacts"
author: "surprisetalk"
captured_at: "2026-06-03T00:43:17Z"
capture_tool: "hn-digest"
hn_id: 48370959
score: 3
comments: 0
posted_at: "2026-06-02T14:45:16Z"
tags:
  - hacker-news
  - translated
---

# Local AI handwriting recognition is usable now

- HN: [48370959](https://news.ycombinator.com/item?id=48370959)
- Source: [www.autodidacts.io](https://www.autodidacts.io/usable-local-ai-handwriting-recognition/)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T14:45:16Z

## Translation

タイトル: ローカル AI 手書き認識が使用可能になりました
記事のタイトル: PSA: ローカル AI 手書き認識が使用可能になりました — The Autodidacts
説明: 嬉しいことに、そして驚いたことに、ローカルの Qwen3-VL:8b WER は、私の不可解な手書きでも許容されます。

記事本文:
宇宙を内側から外側まで探求する
PSA: ローカル AI 手書き認識が使用可能になりました
嬉しいことに、そして驚いたことに、ローカルの Qwen3-VL:8b WER は、私の不可解な手書きでも許容されます。
私は最近、タイプライター原稿の OCR にローカル LLM (別名「AI」) を使用することについて書きました。数日後、真夜中に書いたエッセイで同じ設定を試してみたところ、結果の良さに驚きました。
具体的には、ゾンダーに関する私のエッセイの初稿は、生産性不眠症に関する私のエッセイで説明したテクニックに従って、ドストエフスキーが言うように「ほとんど狂乱の状態で」半分眠っている間に書かれました。 OCR を実行する予定も、下書きを誰かに見せる予定もなかったので、読みやすさではなく速度を最適化していました。
私の手書き文字は…読みやすいことで有名ではありません。私はかなりの量を書きますが、速くて雑な傾向があります。いずれ OCR で私の手書き文字を処理できるようになることを期待していましたが、これまでのところ、それに近いものはありませんでした。
タイプ打ちした文書に使用したのと同じコマンドとプロンプトを使用したところ、単語のエラー率はタイプ打ちした文書よりもはるかに悪かったですが、間違いなく価値がありました。前回と同様に、600dpiでスキャンしたJPGファイルを処理していました。 4 ページを処理するのに約 20 ～ 30 分かかりました。その後、別のエッセイで deepseek-ocr を使用して試してみましたが、処理時間ははるかに速く、精度は比較的同等であるように見えました。
Qwen3-VL には、意図しない場合でも自分の考えを出力に含めてしまう悪い癖があります。 「マイケル効果」を「マシュー効果」に訂正する勇気さえありました。 (もちろん、それは正しかったです。私は名前を間違って覚えていたのではないかと思い、実際にチェックしているときに気づいたでしょう。) これは一線を越えました。たとえテキストが改善されるとしても、OCR でテキストを編集したくないのです。これは特に不快でした。

私の同意なしに修正されていた手書きのテキストは、私の今後のエッセイ「AI を嫌う過小評価の理由」と、目立つように取り上げられた非決定論でした。
Boris Smus のプロンプトを使用して Qwen を軌道上に留めて、再度実行しました。
時間 ls **.jpg |ソート -n | xargs -I{} ollama run qwen3-vl --hide Thinking "./{}\n$(cat path/to/dotfiles/prompts/qwen-handwriting-ocr.txt)" >> Output.txt
品質は他の面で少し後退しましたが (Tesseract は Tessera になりました)、これにより Qwen は独り言を言って私を修正するのをやめました。自分のタイプミスをじっくり考えることができました。
もちろん、クロード・オーパスの方が優れています。ただし、書いた内容をデバイス上に残しておきたい場合は、ローカル ビジョン モデルが民生用ハードウェア上で手書き文字を適切に OCR できるようになりました。そして、Qwen3-VL:8B は現在、OCR Arena リーダーボードのトップの無差別級モデルです。
素晴らしい！受信箱を確認し、リンクをクリックしてサブスクリプションを確認します。
有効なメールアドレスを入力してください！
私は物を書いたり作ったりします。
カナダの湿った西海岸
http://curiositry.com
@curiositry
素晴らしい！受信箱を確認し、リンクをクリックしてサブスクリプションを確認します。

## Original Extract

To my delight and surprise, local Qwen3-VL:8b WER is acceptable, even on my cryptic handwriting

Exploring the universe from the inside out
PSA: local AI handwriting recognition is usable now
To my delight and surprise, local Qwen3-VL:8b WER is acceptable, even on my cryptic handwriting
I recently wrote about using local LLMs (aka “AI”) for OCR of typewritten manuscripts . A few days later, I tried the same set-up with an essay that I’d written in the middle of the night, and was stunned at how good the results were.
Specifically, the first draft of my essay on Sonder , which was written “almost in a frenzy”, as Dostoyevsky might say, while half asleep, following the technique described in my essay on productive insomnia . Since I had no plans to run OCR on it, or show the draft to anyone, I was optimizing for speed, not legibility.
My handwriting is … not renowned for its legibility. I write quite a bit, and tend toward fast and sloppy. I have hoped that eventually an OCR would be able to process my handwriting, but until now, nothing was even close.
Using the same command and prompt as I used for typewritten documents, the word error rate was far worse than typewritten, but definitely worthwhile. As before, I was processing JPG files scanned at 600dpi. It took ~20–30 mins to process 4 pages. I later tried it on another essay, with deepseek-ocr , and processing time was much faster, and accuracy seemed relatively equivalent.
Qwen3-VL has a bad habit of including its thinking in the output even when it’s not supposed to. It even had the gall to correct “michael effect” to “ Matthew effect ”. (It was right, of course; I suspected I had mis-remembered the name, and would have caught it in fact checking.) This crossed a line: I don’t want OCR editing my text, even if it makes it better! This was especially offensive because the handwritten text it was correcting without my consent was my forthcoming essay, Underrated reasons to dislike AI, and non-determinism featured prominently.
I ran it again, using Boris Smus’s prompt to keep Qwen on the rails:
time ls **.jpg | sort -n | xargs -I{} ollama run qwen3-vl --hidethinking "./{}\n$(cat path/to/dotfiles/prompts/qwen-handwriting-ocr.txt)" >> output.txt
The quality regressed a tad in other ways (Tesseract became Tessera), but this got Qwen to stop talking to itself and correcting me. I could bask in my typos.
Claude Opus is, of course, better. But if you want your writing to stay on your device, local vision models can now OCR handwriting decently well on consumer hardware. And Qwen3-VL:8B is currently the top open-weights model on the OCR Arena leaderboard .
Great! Check your inbox and click the link to confirm your subscription.
Please enter a valid email address!
I write and build things.
Canada’s wet west coast
http://curiositry.com
@curiositry
Great! Check your inbox and click the link to confirm your subscription.
