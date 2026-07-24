---
source: "https://www.robinlinacre.com/letterpaths_llms_good_even_when_bad/"
hn_url: "https://news.ycombinator.com/item?id=49030093"
title: "Letterpaths: Or how LLMs can be good even when they're bad"
article_title: "letterpaths: or how LLMs can be good even when they're bad | Robin Linacre's blog"
author: "rzk"
captured_at: "2026-07-24T01:47:54Z"
capture_tool: "hn-digest"
hn_id: 49030093
score: 1
comments: 0
posted_at: "2026-07-24T00:58:35Z"
tags:
  - hacker-news
  - translated
---

# Letterpaths: Or how LLMs can be good even when they're bad

- HN: [49030093](https://news.ycombinator.com/item?id=49030093)
- Source: [www.robinlinacre.com](https://www.robinlinacre.com/letterpaths_llms_good_even_when_bad/)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T00:58:35Z

## Translation

タイトル: レターパス: または LLM が悪い場合でもどのように優れているのか
記事のタイトル: レターパス: または LLM が悪い場合でもどのように優れているのか |ロビン・リナカーのブログ
説明: 視覚化と使い捨て GUI を使用して、LLM に迅速なフィードバックを提供します。

記事本文:
>robinlinacre プロジェクトについて レコード連携 Twitter GitHub LinkedIn Bluesky 初投稿日: 2026-06-15.このページのソースコードはここでご覧ください。
レターパス: または LLM が悪い場合でもどのように優れているのか
Letterpaths は、英国の学童を支援する筆記体手書きを生成するためのライブラリです。その価値は、文字と文字間の結合の品質によって決まります。
開発は LLM 主導で行われました。しかし、フロンティア モデル 1 では、自然な結合を生成することはおろか、正確な文字の形状を生成するなどの基本的な視覚的タスクさえもできませんでした。したがって、文字と接合部は手作業で作成する必要がありました。
それでも、LLM の助けを借りて、おそらく 20 時間の作業でライブラリ全体を構築することができました。
エージェント ループが失敗するこのような状況では、目標は、できるだけ摩擦の少ない方法で人間をループに参加させることです (多くの場合、モデルに GUI を構築させることによって)。このように、AI は主要なタスクには不向きであるにもかかわらず、人間の作業を劇的に加速することができました。
この投稿では、ライブラリを機能させるために私が使用したいくつかのトリックをレビューします。
まず、各文字のジオメトリを表す JSON スキーマを定義しました。
Frontier LLM は、文字の画像をスキーマに変換できないことが判明しました。そこで私は LLM にレター エディターの構築を依頼しました。進捗状況を簡単に確認できるように、簡単なレター ギャラリーも作成しました。
ここで便利なトリックの 1 つは、ブラウザーのファイル システム アクセス API を使用して、ブラウザー アプリが JSON ファイルをディスクに直接保存できるようにすることです。
結合アルゴリズムはライブラリの最も重要な機能です。結合は自然に見え、学校での結合の教え方に非常に近いものである必要があります。
これを言葉で説明するのは難しいことが判明し、モデルでは結果のスクリーンショットを見て反復処理を行ってもアルゴリズムを改善できないことが判明しました。
で

代わりに、さまざまなパラメーター化された結合アルゴリズムと、パラメーターを微調整できる結果のインタラクティブな視覚化を生成するようにモデルに依頼しました。
多くの反復を経て、最終的なアルゴリズムでは、水平距離と垂直距離、最大曲率、合計曲率が考慮され、それらに基づいてカーニング (間隔) も選択されました。
しかし、それはまだ完全に正しいとは言えませんでした。エッジケースが多すぎました。間違っていると思われる特定の文字の組み合わせ。
最終的に、私はカーニングと結合のエディターを構築しました。これにより、すべてのペアの組み合わせを英語での一般的な順に並べ替えて確認、編集、保存できるようになりました。
多くはデフォルトのアルゴリズム設定を使用して保存できますが、これにより、間違っていると思われるペアを迅速にターゲットにすることができました。
フォント (.woff2、.otf) の構築
Letterpaths には、文字を出力し、ジオメトリをフォントとして結合するためのスクリプトが含まれています。
これを正しく理解するのは特に難しく、まだ完璧ではありません。しかし、繰り返しになりますが、LLM にさまざまなアルゴリズムのアプローチを比較するための視覚的なインターフェイスを作成させることが役に立ちました。
フォント ビルダーは私がレターパスに追加した最後の部分であり、その時点では GPT-5.5 と Opus 4.8 が利用可能でした。興味深いことに、Opus 4.8 だけが、文字間に明らかな不連続性を持たずにフォントを生成するというまずまずの作業を行うことができました。それでも、アルゴリズムがどのように機能するかについてのガイダンスを与える必要がありました。
コード変更時のGUIのライブリロード
コア JavaScript ライブラリと関連する GUI を開発するための適切なパターンを見つけるまでにしばらく時間がかかりました。メインのライブラリ コードが使い捨ての GUI やその他のコードから確実に分離されるようにしたいと考えました。しかし同時に、コアのレターパス ライブラリの更新を GUI に即座に反映させる (ライブ リロード) 必要もありました。つまり、たとえば、

つまり、文字のジオメトリの更新は、保存時にどこにでもすぐに表示されます。
レターパス リポジトリは、うまく機能したパターンの例です。これは、さまざまなアプリを含む package/ ディレクトリを持つ小さな pnpm モノリポジトリとして設定されます。これにより、GUI でレターパス ライブラリを直接使用できるようになります。アプリは pnpm dev で実行されます。
2 番目のトリックは、ライブ リロードによって GUI の状態がリセットされないように、GUI の状態 (どの文字が読み込まれているか、設定、マシンのどのフォルダーが開いているか) が URL エンコードされていることを確認することでした。
LLM は、フィードバック ループが確立できる場合に最も効果的に機能します。多くの場合、使い捨て GUI がこのループに役立ちます。
これは、教育アプリ、データ パイプライン、またはその他のソフトウェアを構築しているかどうかに関係なく、非常に一般的なパターンのように感じられます。現在最適化しようとしているものは何か、それを高い情報密度で視覚化するにはどうすればよいでしょうか。最終的に LLM がそこに到達する可能性がある場合でも、現在の問題をより正確に診断できれば、多くの場合、LLM をより迅速に制御できるようになります。
ライブラリ作成時の GPT 5.4 および Opus 4.6。 ↩
このサイトは Observable HQ を使用して構築されており、
アストロ。ソースコード
ここで。
感謝を伝えてください！

## Original Extract

Using visualisations and disposable GUIs to give faster feedback to LLMs

>robinlinacre About Projects Record linkage Twitter GitHub LinkedIn Bluesky Originally posted: 2026-06-15. View source code for this page here.
letterpaths: or how LLMs can be good even when they’re bad
Letterpaths is a library for generating cursive handwriting to help schoolchildren in the UK. Its value rests on the quality of the letters and the joins between them.
Development was LLM-driven. But frontier models 1 were incapable of even basic visual tasks such as generating accurate letter shapes, let alone generating natural joins. Letters and joins therefore had to be built by hand.
And yet, with the help of LLMs I was able to build the whole library in perhaps 20 hours of work.
In situations like this where the agentic loop fails, the goal is to put a human in the loop in the lowest-friction way possible - often by getting the model to build a GUI. Thus, despite being ill suited to the main task, AI was able to dramatically accelerate human work.
This post reviews some of the tricks I used to get the library working.
I started by defining a JSON schema to represent the geometry of each letter.
Frontier LLMs proved unable to convert an image of a letter into the schema. So I asked the LLM to build me the letter editor . I also created a simple letter gallery so I could easily see progress.
One useful trick here was to allow the browser app to save the JSON files straight back to disk for me using the browser’s File System Access API.
The join algorithm is the most important feature of the library. The joins need to look natural and closely approximate how joins are taught in school.
This turned out to be difficult to describe in words, and models proved incapable of improving the algorithm by looking at screenshots of the results and iterating.
Instead, I asked the model to generate various different parameterised join algorithms and an interactive visualisation of the results that allowed me to tweak the parameters.
After many iterations, the final algorithm accounted for horizontal and vertical distance, maximum curvature and total curvature, and also chose kerning (spacing) on their basis.
But it was still not quite right: there were too many edge cases; specific letter combinations that looked wrong.
In the end I built a kerning and join editor that allowed me to review and edit and save all pairwise combinations, sorted in order of how common they were in the English language.
Many could be saved using their default algorithmic settings, but this allowed me to target any pairs that looked wrong rapidly.
Building a font (.woff2, .otf)
Letterpaths contains scripts to output the letter and join geometry as a font.
This was particularly difficult to get right - and is still not perfect. But again, what helped was getting the LLM to create a visual interface to compare different algorithmic approaches:
The font builder was the last part that I added to letterpaths, at which point GPT-5.5 and Opus 4.8 were available. Interestingly, only Opus 4.8 was able to do a passable job of generating the font without obvious discontinuities between letters. Even then, it had to be given guidance about how the algorithm should work.
Live reloading the GUIs when code was changed
It took me a while to figure out a good pattern for developing the core JavaScript library and associated GUIs. I wanted to ensure the main library code was separate from the disposable GUIs and other code. But at the same time, I needed updates to the core letterpaths library to be immediately reflected in the GUI (live reload). This meant, for instance, that an update to the geometry of a letter would be immediately shown everywhere on save.
The letterpaths repo is an example of a pattern that worked well. It’s set up as a small pnpm monorepo with a packages/ directory containing the various apps. This allows the GUIs to use the letterpaths library directly. The apps are run with pnpm dev .
A secondary trick was to ensure the state of the GUI (which letter is loaded, any settings, which folder from the machine is open) was URL-encoded, so that live reload did not reset the state of the GUI.
LLMs work best when a feedback loop can be established. Disposable GUIs can often help with this loop.
This feels like quite a general pattern whether you’re building an education app, a data pipeline, or really any other software: what is the thing you’re currently trying to optimise, and how can you visualise it with high information density. Even when the LLM may get there eventually, you’ll often be able to steer it faster if you can more precisely diagnose the current problem.
GPT 5.4 and Opus 4.6 at the time the library was written. ↩
This site is built using Observable HQ and
Astro . Source code
here .
Say thanks!
