---
source: "https://crlf.link/log/entries/260525-1/"
hn_url: "https://news.ycombinator.com/item?id=48487829"
title: "Making a Vintage LLM from Scratch"
article_title: "Making a vintage LLM from scratch - Cr;Lf;"
author: "croqaz"
captured_at: "2026-06-11T10:36:46Z"
capture_tool: "hn-digest"
hn_id: 48487829
score: 2
comments: 1
posted_at: "2026-06-11T08:38:00Z"
tags:
  - hacker-news
  - translated
---

# Making a Vintage LLM from Scratch

- HN: [48487829](https://news.ycombinator.com/item?id=48487829)
- Source: [crlf.link](https://crlf.link/log/entries/260525-1/)
- Score: 2
- Comments: 1
- Posted: 2026-06-11T08:38:00Z

## Translation

タイトル: ビンテージ LLM をゼロから作成する
記事のタイトル: ビンテージ LLM をゼロから作る - Cr;Lf;

記事本文:
記事
&;&;
について
RSS
ヴィンテージLLMを一から作る
このブログ投稿では、古いテキストのみを使用して (ほぼ) ゼロから独自の LLM を作成する際の冒険を共有します。
私は独自のベーストレーニングおよび微調整スクリプト、データ処理パイプライン、カスタム データセットを作成しました。
(「ほぼゼロから」とは、AI を「ゼロから」構築する他の人と同じように、既存のプログラミング言語とライブラリを使用し、アセンブリで記述しなかったことを意味します...)
モデルは HuggingFace で見つけることができます: https://huggingface.co/croqaz/vintage-LLM-340m-v1-base ;
すべてのコードは https://github.com/croqaz/vintage-LLM でオープンソースです。
より大きなビンテージ モデルを確認したい場合は、私の以前の投稿を参照してください: ヴィンテージ LLM モデル 。
3 か月前の 2 月末、私は Hayk Grigorian による Reddit の投稿をいくつか発見しました。そこで彼は時間ゲート言語モデルの作成について説明していました。すっかり魅了されてしまいました。
1800 年代のロンドンのテキスト、90GB データセットのみで LLM をトレーニングする:
https://reddit.com/r/LocalLLaMA/comments/1pkpsee/training_an_llm_only_on_1800s_london_texts_90gb
1800 年代のロンドンのテキストだけをもとにゼロから訓練された LLM は、1834 年からの実際の抗議活動を持ち出します。
https://reddit.com/r/LocalLLaMA/comments/1mvnmjo/my_llm_trained_from_scratch_on_only_1800s_london
明らかに、独自の LLM を作成した他の人の投稿を読みましたが、おそらく自分でそれを行う準備ができていないか、彼らが取り組んでいたモデルがそれほど興味深いものではなかったのでしょう。とにかく、自分のビクトリア様式のチャット ボットを持つという考えは...クソ壮大です !!
それ以来、私は毎日自分の「Vintage LLM」の開発に取り組みました。例外なく。病気のときも。
その間に、 Violet-1B4-Chat 、 Mr. Chatterbox 、 GPT-1900 、 Talkie 、 TypewriterLM-base など、さらに多くの歴史的な LLM がリリースされました。
何？
これはタイムロック LLM/歴史的 LLM です (英語)

h のみであり、その知識の限界は 1900 年です。
(特定の年に限定するとエラーが発生しやすくなりますが、最善の努力をしました)。
Llama アーキテクチャに基づいており、340M (0.3B) のパラメータがあります。
なぜ？
なぜなら、自分でやってみないと学ぶことができないし、とても楽しいプロジェクトだからです。
どこで、どのようにして？
私は独自のデータセット、独自の処理およびトレーニング コードを作成しました。
コードは、VS-Code と PI (OpenRouter モデル) で使用した LLM を使用してセミバイブコーディングされています。
すべての関数をチェックして検証し、すべてのコード ファイルが何を行っているかを深く理解しました。
データセットの処理に最も時間がかかり、あらゆる種類のことを試しましたがうまくいかず、膨大な時間を無駄にしました。複雑な解決策は最悪です...
私はすべてのデータを自分の PC で処理し、LLM の小さいバージョンを自分の PC (Cachy OS Linux、AMD Ryzen 7 9700X CPU、64GB RAM、Radeon RX 9070 16GB VRAM) でトレーニングしました。
より大きな 340M モデルに関しては、RunPod、ThunderCompute、および Vast.ai でトレーニングしました。私のPCでは永遠に時間がかかったでしょう。
このプロジェクトの総コストは、GPU コストのみで約 80 ドルでした。
それはデータを処理するためのまともな PC を持っているからです。もっと多くの RAM があれば、特にメモリ内のテキストの重複を除去する場合に、一部のデータをより速く処理できたでしょう。
免責事項 : これはおもちゃ/ホビー LLM です (ただし、私は非常に真剣に扱っています)。
それは幻覚を引き起こし、当時は正常であると考えられていましたが、今日の基準では有毒で攻撃的で安全ではないと考えられる歴史的な半正確なコンテンツを生成します。調整を行っていないため、これは想定内のことです。モデルの調整 (または打ち切り) には多大な労力が必要であり、これまでの精度が損なわれてしまいます。
また、（最善を尽くしたとしても）私のモデルが 1900 年に厳密に限定されることは保証できません。たとえば、「アルバート アインシュタイン テスト」を実行する場合などです。
私は毎日仕事で AI を使用していますが、

それがどのように機能するかは理解していますが、LLM を自分で構築したことはありません。私は仕事で特定の AI トレーニングと微調整パイプラインを実行し、過去には C と Python で小さなニューラル ネットワークを構築しましたが、このプロジェクトを開始したとき、人々が通常どのように LLM を構築しているのか知りませんでした。
私は 1 週間検索し、さまざまな視点を得るために複数のボットとチャットしました (トピックを調査するときにいつもそうしているように)。
つまり、LLM を構築するには 4 つのものが必要です。
データ -- LLM には識別力も理解力もありません。良いことも悪いことも、あなたが話しかけることから学習します。これが最も長いプロセスです。
トークン化 -- Tokenizer は、単語または文字を数値 (トークン) に変換する小さなプログラムです。 LLM は単語を理解せず、数値のみを理解します。
事前トレーニング -- これは紛らわしい表現ですが、LLM がテキストのオートコンプリートを学習する「基本トレーニング」を意味します。 300m 以上のパラメータを使用する場合、これが最も高価なプロセスになります。
微調整 -- LLM が順番にチャットしたり、質疑応答したりする方法を学習します。
これらの簡単な手順以外にももう少し詳しく説明しますが、この記事ではこれ以上深くは説明しません。
それでは、各ステップを詳しく見てみましょう。
言及する価値があるのは、私は多くの間違いを犯し、「大きな」モデルに落ち着く前にいくつかのデータセットとモデル アーキテクチャを実験したことです。
(引用符で囲まれた「大きい」というのは、トーキー 13B やタイプライター LM-7.24B のような大きなモデルと比較すると、私のモデルは単なるおもちゃにすぎないからです)
自分の PC でトレーニングした v2 おもちゃ EleutherAI/pythia-14m の詳細:
https://github.com/croqaz/vintage-LLM/tree/e272b94fcf96316f874babbed549d20809fe5a39/models/m-v2
検証損失とパープレキシティ SVG ファイルを見ると、大きなジャンプが見られます。これは、ファイル チャンクをランダム化しておらず、データセット ファイルがアルファベット順にトークン化されているためです。

最初はクリーンなブック ファイルだったわけではなく、モデルが Time-Capsule データセットにさらされ始めるとすぐに、データセットが奇妙な OCR アーティファクト、壊れた単語や文章などを多く含んでいたため、徐々に状態が悪化していきました。
間違いもあった...しかし私はそこから学びました。
不良文書を除外し始めるまで、しばらく行き詰まっていました。
データ処理はこれまでで最も長く、最も退屈なプロセスでしたが、その理由を理解していただけたと思います...
インターネットからデータを収集した最新の高品質なデータセットはたくさんありますが、LLM にコンピューター、原爆、宇宙船について学ばせたくなかったので、独自のデータセットを作成する以外に選択肢はありませんでした。
幸いなことに、利用可能なデータセットはいくつかありますが、それらはかなり品質が悪く、私の作業のほとんどは、重複を除去し、本当に悪いテキストをフィルタリングし、既存のテキストの一部を強化することでした。
歴史的なデータセットは非常に限られており、さらに古い本を発見して誰かがスキャンしない限り、私たちが持っているのは古い本だけなので、持っているものを使用する必要があります。
言及する価値のあるデータセットとしては、Project Gutenberg、Oxford Text Archive、Internet Archive 書籍、TheBritishLibrary/blbooks、storytracer/LoC-PD-Books、dell-research-harvard/AmericanStories、dell-research-harvard/NewsWire、Heritage Made Digital Newspapers (HMD) があります。
各データセットの年と言語を見つけるために最善を尽くしたので、1900 年以前のものをすべて英語に制限できます。
たとえ品質が良くても、年が明記されていない文書や本文中に日付が見つからない文書は、念のため完全に無視しました。
サイド プロジェクトとして、たくさんの古い本、タイトル、著者、書籍 ID、ソースを含む Book-Metadata HF データセットを作成しました: https://huggingface.co/datasets/croqaz/book-metadata ;私の目標は、すべてのグーテンベルク本の出版年を調べることでしたが、私が知っているのはそれだけです。

そして私がその年について100％確信している5300冊の本。
繰り返しますが、これにはまったく永遠に時間がかかり、このブログ記事を書いている時点でもまだ完全には終わっていません。
別の LLM をトレーニングすることがあれば、次回はより多くのより良いデータを取得できるでしょう。
当初、私は MinHash や埋め込みベクトルの類似性など、いくつかの重複除外メソッドを必要としていました。意味が分からなくても、詳しくは説明しませんのでご安心ください。あまりにも遅くて高価だったので、断念せざるを得ませんでした。これがどれほど遅いかを体験してもらうために、強力な DEV サーバーに短いテキスト データセットのエンベディングを計算させ、昼夜を問わずサーバーを実行して 1 週間でデータセットの 10% を計算しました。サーバーには RTX 3090 GPU があり、それを共有しました。
最終的に、正規化されたテキスト (すべてのスペースが削除された小文字のテキスト) に基づいて重複を排除しました。基本的に、テキスト「hello world」はテキスト「Hello World」と同一であり (スペースとタイトルの大文字小文字に注意してください)、テキストはデータセットに 1 回だけ保存されます。
最初から、データが最も重要であることはわかっていました。つまり、ガベージ イン -> ガベージ アウトです。私は非常に多くの実験と反復を行い、データセットを保存するために Qdrant、Zvec、Lance、ValKey、LevelDB などの DB を試しました。
多くのエントリを追加し始める前から DB ディスク サイズが巨大だったので、Qdrant を削除しました。
Zvec を削除したのは、DB エントリを循環させる方法がなかったためです。基本的に、一度エントリを保存すると、DB を探索する方法がありません。これに関して問題を作成しました。 Zvec も非常に新しいので、おそらく成熟するまでもっと時間を与える必要があります。
私が Lance を削除したのは、バージョン管理のためです。数百万を超えるエントリを追加し始めると、DB はますます遅くなります。これは私のせいかもしれませんが、もっと良い方法を見つけることができると確信しています。
何かを注入した後、メモリが不足したためValKeyを削除しました

レコードが 1,000 万件あると、サーバーが私の PC で OOM クラッシュし始めましたが、それでもさらに多くのデータを取り込む必要がありました。それ以外は、ValKey は本当に素晴らしかったです。
私は最終的に LevelDB (ビットコインとイーサリアムのトランザクションを保存するためにローカル ウォレット アプリで使用されていた) を使用したので、PC 上で拡張できることがわかりました。 CPU または RAM の使用量を最小限に抑え、問題なく 1,200 万行を挿入しました。 LevelDB は時々遅くなることがありますが、一貫して信頼性があります。
もっと優れた PC やスーパー コンピューターを持っていたら、おそらくずっと ValKey を使っていたでしょう。
テキストの品質を把握するために、私はまず各文書の長さと固有の文字を調べました。最初の段階では短いテキスト (最大 32k の長さ) を使用し、2 番目の段階では最大 10 MB の長いテキストを使用することにしました。英語では通常、30 ～ 50 個を超える記号を使用すべきではありません。テキストの塊に 100 個以上の固有の記号が含まれている場合、それは英語ではないため、それらは破棄しました。また、テキストに固有の記号が 8 個しかない場合は意味をなさないので、それらも削除しました。
超簡単な指標、ZLIBの圧縮率。短すぎて多様なテキストは大きな価値を持ちますが、過度に繰り返されるテキストは非常に小さな価値を持ちます。
# ZLIB 圧縮率
# 適切なウィンドウは 0.5...0.7 です。
def 圧縮率(テキスト) -> 浮動小数点:
raw = text.encode("utf-8")
圧縮 = zlib.compress(生)
len(圧縮) / len(生)を返します
Compression_ratio("ロレム・イプサム・ドーロ・シット・アメット")
#1.3
圧縮率("その他とその他とその他" * 100)
# 0.01 -- 非常に繰り返し
Compression_ratio("大統領は、トーマス・ジョンソン、ウィリアム・クランチ、チャールズ・シムズをコロンビア特別区判事に\n指名した。\n\n先週土曜日、\n現在アメリカ合衆国副大統領であり上院議長でもあるトーマス・ジェファーソンは\nその会を休み、\n次の演説を行った:\n\n

元老院の紳士諸君、\n\nいつもの機会を与えてください。")
# 0.64 -- 通常のテキスト
シャノンのエントロピーも次のようになります。
#シャノンエントロピー
# 印刷された英語の推定エントロピー レートは約 4.2...5.5 です。
def char_entropy(text) -> float:
カウント = カウンター(テキスト)
合計 = len(テキスト)
エントロピー = 0.0
counts.values() のカウントの場合:
p = カウント / 合計
エントロピー -= p * log2(p)
戻りエントロピー
char_entropy(("a " * 10 + "!"))
# 1.22 -- 低すぎる
char_entropy("ロレム・イプサム・ドーロ・シット・アメット")
# 3.6 -- 少し低い
char_entropy("リッチモンド地区の大公法高等法院では\nヘンリー・バンクス原告、\nそして\nナサニエル・アンダーソン、ロバート・ポラードの間。")
# 4.5 -- 通常の英語
char_entropy(''.join(chr(i) for i in range(200)))
# 7.6 -- 超高エントロピー
そして独自の品質検出フィルター。これは、奇妙な記号がたくさん含まれる非常に悪い OCR テキストを識別するのに役立ちました。
_LETTER_RE = re.compile(r'[a-zα-ωàââäçèéêëîïôöùûüüÿæœß]$', re.I)
_DIGIT_SPACE_RE = re.compile(r'[0-9 \n]$')
_PUNCT_RE = re.compile(r'[.,;!?\'"_\-]$')
# Cro のカスタム品質スコア
# 通常の文字のスコアは 2 です。
# 数字とスペースのスコアは 1 になります。
# 句読点は 0.5 を取得します。
# 他の文字は -0.5 を取得します。
defquality_score(text: str) -> float:
スコア = 0.0
フォ

[切り捨てられた]

## Original Extract

Articles
&;&;
About
RSS
Making a vintage LLM from scratch
In this blog post, I will share the adventures I had creating my own LLM, from (almost) scratch, trained only on old texts.
I made my own base-training and fine-tuning scripts, data processing pipelines and custom datasets.
("almost from scratch" means I did use existing programming languages and libraries, I didn't write in Assembly, just like anyone else who builds an AI "from scratch"...)
The model can be found on HuggingFace: https://huggingface.co/croqaz/vintage-LLM-340m-v1-base ;
All the code is open source at: https://github.com/croqaz/vintage-LLM ;
If you want to check bigger Vintage models, see my previous post: Vintage LLM models .
Three months ago at the end of February I discovered a few Reddit posts by Hayk Grigorian , where he described creating his temporal gated language model. I was absolutely fascinated.
Training an LLM only on 1800s London texts, 90GB dataset:
https://reddit.com/r/LocalLLaMA/comments/1pkpsee/training_an_llm_only_on_1800s_london_texts_90gb
LLM trained from scratch on only 1800s London texts brings up a real protest from 1834:
https://reddit.com/r/LocalLLaMA/comments/1mvnmjo/my_llm_trained_from_scratch_on_only_1800s_london
Obviously I read other posts from other people that made their own LLMs, but maybe I wasn't ready to do it myself, or the model they were working on wasn't that interesting. Anyway, the thought of having my own Victorian chat bot... fuckin' epic !!
Since then, I worked on my own "Vintage LLM" every single day. Without exceptions. Even when I was sick.
In the meantime, a lot more historic LLMs have been released like: Violet-1B4-Chat , Mr. Chatterbox , GPT-1900 , Talkie and TypewriterLM-base .
What?
This is a time-locked LLM/ historical LLM, English only, and its knowledge cutoff is year 1900.
(Limiting to a specific year is error prone, but I did my best effort).
It is based on Llama architecture and has 340M (0.3B) params.
Why?
Because I can only learn if I do it myself and it's a super fun project.
Where and how?
I made my own dataset, my own processing and training code.
The code is semi-vibe-coded with whatever LLM I had with VS-Code and PI (OpenRouter models).
I checked and validated every single function and I deeply understand what every single code file is doing.
The dataset processing took the most and I tried all sorts of things that didn't work, and I wasted a ton of time. Complicated solutions are the worst...
I processed all the data on my own PC and I trained smaller versions of the LLM on my PC ( Cachy OS Linux, AMD Ryzen 7 9700X CPU, 64GB RAM, Radeon RX 9070 16GB VRAM ).
As for the larger 340M model, I trained it on RunPod, ThunderCompute and Vast.ai . It would have taken forever on my PC.
The total cost of this project was: ~$80, GPU costs only.
That's because I have a decent PC to process the data. If I had more RAM, I could have processed some of the data much faster, especially when it comes to de-duplicating texts in memory.
Disclaimer : This is a toy/ hobby LLM (but I treat it very seriously).
It will hallucinate and generate historic semi-accurate content which, at the time was considered normal but by today's standards is considered: toxic, offensive and unsafe. This is expected, because I didn't do any alignment. Aligning (or censoring) the model requires significant effort and it would ruin the historic accuracy.
Also, I can't guarantee that my model is strictly limited to the year 1900 (even if I did my best) eg: as to perform the "Albert Einstein test".
I use AI everyday at work and I understand how it works, but I never built an LLM myself. I ran specific AI training and fine-tuning pipelines at work, I built tiny neural networks in C and Python in the past, but when I started this project I didn't know how people are usually building LLMs.
I searched for a week and I chatted with multiple bots to get different points of view (like I always do when I research a topic).
In short, to build an LLM you need 4 things:
the data -- an LLM has no discernment or understanding. It will learn from anything you tell it to, good or bad. This is the longest process.
tokenization -- the Tokenizer is a little program that converts words or letters into numbers (tokens). LLMs don't understand words, they only understand numbers.
pre-training -- it's a confusing expression and it means "base-training", where the LLM learns to autocomplete text. If you're going for a 300m+ params, this is the most expensive process.
fine-tuning -- where the LLM learns how to chat in turns, question & answer.
Well, there's a bit more to it than these simple steps, but I won't go super deep in this article.
Now, let's look at each step in more detail.
It's worth mentioning that I made lots of mistakes and I experimented with some datasets and model architectures before I settled on the "big" model.
("big" in quotes here, because you know, compared to larger models like Talkie-13B or TypewriterLM-7.24B, my model is just a toy)
Some details about my v2 toy EleutherAI/pythia-14m that I trained on my own PC:
https://github.com/croqaz/vintage-LLM/tree/e272b94fcf96316f874babbed549d20809fe5a39/models/m-v2
If you look at the validation loss and perplexity SVG files, you'll see huge jumps, because I didn't randomize the file chunks and the dataset files were tokenized in alphabetic order, and it happened that the clean book files were first and as soon as the model started to get exposed to the Time-Capsule dataset, it gradually become worse and worse because that dataset was bad with lots of weird OCR artefacts, broken words & sentences and so on.
Mistakes were made ... but I learned from them.
I was stuck for a while, until I started to filter out the bad documents.
The data processing was the longest and most boring process by far, and I hope you understand why...
There are plenty of modern, high quality datasets with data scraped from the internet, but I didn't want my LLM to learn about computers, atomic bombs and space-ships, so I had no choice but to make my own.
Luckily there are some datasets available, but they are pretty bad quality and most of my work was to de-duplicate, filter the really bad texts and enhance some of the existing texts.
The historic datasets are quite limited, the old books are all we have unless we discover more old books and someone scans them, so we have to use what we have.
A few datasets worth mentioning are: Project Gutenberg, Oxford Text Archive, Internet Archive books, TheBritishLibrary/blbooks, storytracer/LoC-PD-Books, dell-research-harvard/AmericanStories, dell-research-harvard/NewsWire, Heritage Made Digital Newspapers (HMD).
I did my best to find out the year and language of each dataset, so I can limit everything to English before year 1900.
I completely ignored documents that don't specify a year, or I can't find date in the texts, even if they had good quality, just to be safe.
As a side-project, I created the Book-Metadata HF dataset with lots of old books, title, author, book ID and source: https://huggingface.co/datasets/croqaz/book-metadata ; My goal was find out the year of all Gutenberg books, but I only found 5300 books where I'm 100% sure of the year.
Again, this took absolutely forever and as I am writing this blog post, I'm still not completely finished.
If I'll ever train another LLM, I'll have more and better data next time.
Initially I wanted several de-duplication methods, including MinHash and embeddings vector similarity. If you don't know what it means, don't worry, I won't go into details. It was way too slow and expensive and I had to abandon it. Just to give you a taste of how slow this was, I had a beefy DEV server calculate embeddings for my short text datasets and I did 10% of the dataset in a week with the server running day & night. The server had a RTX 3090 GPU that I shared and so on.
In the end, I de-duplicated based on the normalized text (lower-case text with all spacing removed). Basically the text "hello world" is identical to the text " Hello World " (notice the spaces and title case), and the text is only saved once in my dataset.
From the start, I knew that the data is the most important: garbage in -> garbage out. I did lots and lots of experiments and iterations, I played with DBs like Qdrant, Zvec, Lance, ValKey, and LevelDB for storing the datasets.
I dropped Qdrant because the DB disk size was huge even before I started adding many entries.
I dropped Zvec because I had no way to cycle the DB entries, basically once you save them, there's no way to explore the DB. I created an issue for this . Zvec is also super new and I should probably give it more time to mature.
I dropped Lance because of the versioning, the DB becomes slower and slower once you start adding more than a few million entries. This could be my fault, I'm sure I can find out a way to do this better.
I dropped ValKey because I ran out of memory, after injesting something like 10 mil records, the server started to OOM crash on my PC, but I still had to injest way more data. Other than that, ValKey was really great.
I ended up using LevelDB (which was used by local wallet apps to store BitCoin and Ethereum transactions), so I know it can scale on my PC. I injested 12 mil rows without any issues and with minimal CPU or RAM usage. LevelDB can be slow at times, but it's consistently reliable.
If I had a better PC, or a super computer, I would probably have used ValKey all the way.
To get a sense of the quality of the texts, I was first looking at the length and unique characters of each document. For the first stage I decided to use short texts (up to 32k long) and for the second, long texts up to 10MB. English shouldn't use more than 30-50 symbols normally; if a chunk of text has 100 or more unique symbols, it's not English, so I discarded those. And if a text has only 8 unique symbols, it doesn't make any sense, so I removed those too.
Super easy metric, the compression ratio of ZLIB. Text that is too short and diverse will have a big value and text that is super repeated will be really small value.
# ZLIB compression ratio
# A good window is 0.5...0.7;
def compression_ratio(text) -> float:
raw = text.encode("utf-8")
compressed = zlib.compress(raw)
return len(compressed) / len(raw)
compression_ratio("Lorem ipsum dolor sit amet")
# 1.3
compression_ratio("other, and other and other" * 100)
# 0.01 -- very repeated
compression_ratio("The President has nominated Thomas Johnson, William Cranch, and Charles\nSimms, Judges of the district of Columbia.\n\nOn Saturday last, Thomas Jefferson, at\npresent Vice President of the United States,\nand President of the Senate, took leave of\nthat body on which occasion he delivered\nthe following address:\n\nGentlemen of the Senate,\n\nTo give the usual opportunity")
# 0.64 -- Regular text
Also the Shannon entropy:
# Shannon Entropy
# The estimated entropy rate for printed English is approximately 4.2...5.5;
def char_entropy(text) -> float:
counts = Counter(text)
total = len(text)
entropy = 0.0
for count in counts.values():
p = count / total
entropy -= p * log2(p)
return entropy
char_entropy(("a " * 10 + "!"))
# 1.22 -- Too low
char_entropy("Lorem ipsum dolor sit amet")
# 3.6 -- A bit low
char_entropy("IN the High court of Chancery for the Rich\nmond District,\nBetween\nHenry Banks plaintiff,\nAnd\nNathaniel Anderson, Robert Pollard.")
# 4.5 -- Regular english
char_entropy(''.join(chr(i) for i in range(200)))
# 7.6 -- Super high entropy
And my own quality detection filter. This helped me identify very bad OCR texts, with lots of weird symbols:
_LETTER_RE = re.compile(r'[a-zα-ωàâäçèéêëîïôöùûüüÿæœß]$', re.I)
_DIGIT_SPACE_RE = re.compile(r'[0-9 \n]$')
_PUNCT_RE = re.compile(r'[.,;!?\'"_\-]$')
# Cro's custom quality score
# Regular letters get a score of 2;
# Digits and spaces get a score of 1;
# Punctuation gets 0.5;
# Any other characters get -0.5.
def quality_score(text: str) -> float:
score = 0.0
fo

[truncated]
