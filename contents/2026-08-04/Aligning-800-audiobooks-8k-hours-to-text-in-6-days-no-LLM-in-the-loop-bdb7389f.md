---
source: "https://tale.fyi/@sam/aligning-800-audiobooks-8-000-hours-to-text-in-6-days-no-llm-in-the-loop?from=hn"
hn_url: "https://news.ycombinator.com/item?id=49168259"
title: "Aligning 800 audiobooks (8k hours) to text in 6 days, no LLM in the loop"
article_title: "Aligning 800 audiobooks (8,000 hours) to text in 6 days, no LLM in the loop by samuel cole · tale.fyi"
author: "samuelcole"
captured_at: "2026-08-04T13:51:58Z"
capture_tool: "hn-digest"
hn_id: 49168259
score: 2
comments: 0
posted_at: "2026-08-04T12:58:58Z"
tags:
  - hacker-news
  - translated
---

# Aligning 800 audiobooks (8k hours) to text in 6 days, no LLM in the loop

- HN: [49168259](https://news.ycombinator.com/item?id=49168259)
- Source: [tale.fyi](https://tale.fyi/@sam/aligning-800-audiobooks-8-000-hours-to-text-in-6-days-no-llm-in-the-loop?from=hn)
- Score: 2
- Comments: 0
- Posted: 2026-08-04T12:58:58Z

## Translation

タイトル: 6 日間で 800 冊のオーディオブック (8,000 時間) をテキストに調整、ループ内に LLM なし
記事のタイトル: 800 オーディオブック (8,000 時間) を 6 日間でテキストに調整、ループに LLM を使用しない by samuel cole · Tale.fyi
説明: 800 オーディオブック (8,000 時間) を 6 日間でテキストに調整します。サミュエル・コールによるループ内に LLM は使用せず、1 つの長く読みやすいページとして作成します。

記事本文:
6 日間で 800 冊のオーディオブック (8,000 時間) をテキストに調整、ループ内に LLM なし by samuel cole · Tale.fyi Tale.fyi
物語を読むとこんな感じになります。
自動 800 オーディオブック (8,000 時間) を 6 日間でテキストに調整します。ループ内に LLM はありません
サミュエル・コール ノンフィクションをフォローする
7 月 14 日に、 Tale という git リポジトリを開始しました。数か月前に、私は大学時代に発掘した「つぎはぎの心臓を持つ男」と呼ばれる古い原稿の簡単なプロトタイプを作成していました。私が 2009 年に書いたときも、2026 年の今日も、この論文はシンプルでした。「文学はウェブ上で機能するのか?」ほんのわずかしか持っていない自分の著作に加えて、Standard Ebook のコーパスには何千ものパブリック ドメインの著作物が含まれていることがわかっていました。それらをすべて Web サイト上で適切に動作させ、美しい読書体験を提供できれば、ブラウザが文学を読むのに最適な場所であることを証明できるでしょう。
次の 8 日間、私は Lemontree で開発したのと同じエージェント コーディングの規律を適用し、6 月に別のコードベースで洗練させて、エレガントなシングル ページ リーダーを開発しました。このリーダーは、アカウント不要でシンプルで読みやすい URL で誰でもコピーできる、匿名優先の記法でデバイス間で進行状況を保存しました。
そして、7 月 22 日の真夜中頃に、私は、tale.fyi の最初の本当に魔法の機能になったと思うアイデアを思いつきました。それは、読書と同時にオーディオブックを聞くことができ、車の中でオーディオブックを聞くことから昼休みに読書にシームレスに切り替えることができたらどうなるでしょうか。ちょっとGoogleで調べた後、LibriVoxで驚くべきパフォーマンスを発見しました。これらはすべてパブリックドメインとしてリリースされており、これは私たちの最初の概念実証、つまりオープンウェブのささやき同期を形成する可能性があります。エージェントと私は 2 つのルールを決めました。オーディオブックは、単語ごとにテキストに合わせて配置する必要があります。

読み上げ版を出荷する前に、調整が確実であるという確信が持てます。
そう決めたので、エージェントに /goal を設定し、就寝しました。エージェントは当初、測定されたテストの実行が速かったため (MacBook の CPU でリアルタイムで 380 倍、強制 CTC アライメントよりも 60 倍速かった)、aeneas ライブラリが答えになるだろうと考えていましたが、大量の本を読んだ後、午前 3 時 12 分に、それらが私たちの ASR グラウンド トゥルースと一致していないと判断しました。ドラキュラの最初の章の段落 27 は、実際の ~1490 秒 (+582 秒) に対して 2072 秒に配置されていました。この章の後半3分の1は146秒に短縮されました。さらに重要なのは、aeneas は信頼度スコアをエクスポートしないため、信頼性が失われている箇所にフラグを立てることさえできませんでした。午後 2 時 24 分、キーボードの前に戻った私は、より現代的で信頼性の高いソリューション、つまり Meta の大規模多言語音声モデルを使用した CTC ニューラル アライメントにたどり着きました。最初は遅すぎると拒否していましたが、処理を GPU に移行することで劇的に高速化されました (106 倍リアルタイム)。速度は 3 倍 (60 倍ではありません) でしたが、より正確で、最も重要なことに、信頼度スコア (確率) が生成されました。
転写してから整列してみてはいかがでしょうか？ Transcribe-then-align は常に何かを発行します。不一致のエディションでは、何かが微妙に間違っています。強制配置では、テキストを音響に直接配置し、1 つのモデルと 1 つのエラーソースを配置します。ワードタイミングは 20ms に解決され、ロックできない場合は棄権されます。 「私たちは棄権を好みます」は、このプロジェクト全体の指針となった README の正確な文言です。
最終的に、Evals は同期スクリプトを改良するために重要になりました。オプションの前付けが付いた書籍が問題を引き起こしました。 『ベーオウルフ』では、ナレーターが訳者の序文と巻末注を飛ばしたため、ベーオウルフは開始 3 分で中央値のみのゲートを通過した。イソップ寓話は 380 段落のうち 39 段落のみに配置され、自信満々に塗りつぶされていた

オーディオ全体にわたって (スパン カバレージ 100%、中央値 0.93)、壊れた状態で出荷されました。このようなことが二度と起こらないように、新しい全冊ゲートを追加しました。
メモリ不足エラーにより、長い本のアライナーがクラッシュしました。私はディスクからのストリーミング技術を開発しました。これはクロック時間がかかるため、15 時間を超える書籍にのみ適用されます。
まだクラッシュがあり、「大きな本のメモリが不足している」という考えは間違っていたことが判明しました。1.6 時間のタオ・テー・チン号は死亡しましたが、17.9 時間のダーバーヴィルのテス号は航行しました。本当の原因は、整列がまばらなことだった。本の長さではなく、固定された文章の間のギャップによって記憶が爆発したのだ。その差を300秒で詰めた。
ついに、真夜中のアイデアから 6 日後の 7 月 28 日、私はクラッシュすることなく、ミラーリングされた 816 冊のオーディオブックすべて、約 8,000 時間のナレーションを実行し、それぞれに判定を付けました。 553 は読み物を出荷しました。門は残りを拒否しました。次の数日間、私たちは最初に拒否された 100 冊以上の書籍を救出する作業を繰り返しました。
これで、7 月 22 日深夜のアイデアから生まれた 668 冊のオーディオブックを、すべて整列して、 Tale.fyi/audiobooks のテキストまたはオーディオから再開できます。費用は？コンピューティングは 0 ドル (すべて私の MacBook 上で実行されました)、ストレージは月額 6 ドルです。
著者によって出版されたレポート

## Original Extract

Aligning 800 audiobooks (8,000 hours) to text in 6 days, no LLM in the loop by samuel cole, as one long readable page.

Aligning 800 audiobooks (8,000 hours) to text in 6 days, no LLM in the loop by samuel cole · tale.fyi tale.fyi
This is how the story will feel to read.
auto Aligning 800 audiobooks (8,000 hours) to text in 6 days, no LLM in the loop
samuel cole follow non-fiction
On July 14th, I started a git repo called tale . Several months before I had built a simple prototype for an old manuscript from college that I unearthed called the man with the patchwork heart . The thesis was simple, both when I wrote it in 2009, and today in 2026: could literature work on the web? In addition to my own writing, which I only have so much of, I knew the corpus of Standard Ebooks could fill in thousands of public domain works, and if I could get them all working well on a website, with a beautiful reading experience, I could prove that the browser could be a great place to read literature.
Over the next 8 days, I applied the same agentic coding discipline that I had developed at Lemontree , and refined on another codebase in June, to develop an elegant single page reader, which stored your progress across devices with anonymous-first notations anyone could copy with simple, readable urls, no account required.
Then on July 22nd around midnight I had an idea which I think became the first truly magical feature of tale.fyi: what if we could listen to an audiobook along with our reading, so that we could seamlessly switch from listening to an audiobook in the car, to reading during our lunch break. After a quick Google, I discovered the incredible performances on LibriVox, all released as public domain, which could form our first proof-of-concept: a whisper-sync for the open web. My agents and I made two rules: the audiobooks must be aligned word-by-word to the text, and there must be high confidence that the alignment is solid before I ship the read-along.
With that decided, I set a /goal on my agents and went to bed. The agents originally thought the aeneas library would be the answer, because our measured test runs were fast (380x realtime on my MacBook's CPU, 60x faster than forced CTC alignment), but after running through a bunch of books, they decided at 3:12am that they just weren't matching our ASR ground truth: paragraph 27 of Dracula's first chapter placed at 2072s against the actual ~1490s (+582 seconds), the chapter's back third crushed into 146 seconds. More critically: aeneas does not export a confidence score, so I couldn't even flag where it was losing confidence. At 2:24pm, with me back at the keyboard, I landed on a more modern, higher confidence solution: a CTC neural alignment using Meta's Massively Multilingual Speech model. What I first rejected for being too slow was sped up dramatically by moving the processing to GPU: 106x realtime. It was 3x slower (not 60x), but more accurate, and most critically, emitted confidence scores (probabilities).
Why not transcribe, then align? Transcribe-then-align always emits something: on a mismatched edition, something subtly wrong; forced alignment aligns text to acoustics directly, one model and one error source. Word timing resolves to 20ms, and when it can't lock, it abstains. "We prefer abstaining" is the exact wording in the README, which has guided this whole project.
Evals ended up being critical for refining the sync script. Books with optional front matter caused trouble. In Beowulf , the narrator skips the translator's preface and endnotes, so Beowulf passed the median-only gate while beginning three minutes in. Aesop's Fables placed only 39 of 380 paragraphs, smeared confidently across the whole audio (span coverage 100%, median 0.93) and it shipped broken; I added a new whole-book gate to prevent that from happening again.
Out of memory errors crashed the aligner on long books. I developed a streaming from disk technique, which only applies on books longer than 15 hours, since it costs clock time.
I still had crashes, and "big books run out of memory" turned out to be the wrong tree: the 1.6-hour Tao Te Ching died while the 17.9-hour Tess of the d'Urbervilles sailed through. The real culprit was sparse alignment: memory blew up with the gap between anchored passages, not the length of the book. I capped the gap at 300 seconds.
Finally, on July 28th, six days after the midnight idea, I had a completed run, without crashes: all 816 mirrored audiobooks, about 8,000 hours of narration, each with a verdict. 553 shipped a read-along; the gates rejected the rest. Over the next few days, we kept iterating to rescue over a hundred of the initially rejected books.
Now you can enjoy 668 audiobooks, all aligned, all resumable from either text or audio on tale.fyi/audiobooks , all from a midnight idea on July 22nd. The cost? $0 compute (it all ran on my MacBook) and $6 per month in storage.
published by its author · report
