---
source: "https://im.not.ci/echo-mini-rse/"
hn_url: "https://news.ycombinator.com/item?id=48937156"
title: "Reverse Engineering the Firmware of an MP3 Player with LLMs"
article_title: "- The Story Behind Reverse Engineering the Snowsky Echo Mini Firmware"
author: "uneven9434"
captured_at: "2026-07-16T17:03:45Z"
capture_tool: "hn-digest"
hn_id: 48937156
score: 1
comments: 0
posted_at: "2026-07-16T17:02:53Z"
tags:
  - hacker-news
  - translated
---

# Reverse Engineering the Firmware of an MP3 Player with LLMs

- HN: [48937156](https://news.ycombinator.com/item?id=48937156)
- Source: [im.not.ci](https://im.not.ci/echo-mini-rse/)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T17:02:53Z

## Translation

タイトル: LLM を使用した MP3 プレーヤーのファームウェアのリバース エンジニアリング
記事のタイトル: - Snowsky Echo Mini ファームウェアのリバース エンジニアリングの背後にある物語

記事本文:
Snowsky Echo Mini ファームウェアのリバース エンジニアリングの背後にある物語
私は音楽を保存するために使用するメディアに対して、常に奇妙な執着を持っていました。小学生の頃、自宅のテープレコーダーを使ってミックステープを作るのが大好きでした。とても良い思い出がたくさん残りました。最近の夢は自分のカセットウォークマンを持つことです。
しかし、これを実現するのは簡単ではありません。ウォークマンは高価で、頭の痛い問題もたくさんあります。 （今年の私の誕生日の願いは、優しい友達がプレゼントしてくれることだった。でも、そんなことを考えただけでも恥知らずだと言われるだろうね。）
その消費者の深い空白を埋めるために、私は半年前に Fiio Echo Mini を購入することに決めました。ミニカセットマシンのようなデザインのMP3プレーヤーです。
しばらく使ってみると、なんだか複雑な気持ちになりました。いくつかの点ではまともですが、全体的な実行はひどいです。
見るのが最も苦痛なのは、その大きなピクセル化された画面です。 15 年前なら、これは平均的なディスプレイだと考えていたかもしれません。しかし 2026 年、その画面を見つめているのは奇妙に感じられます。たとえ画面がピクセル化されていたとしても、優れたビジュアルデザイン (本物のレトロなピクセルアートなど) があれば、それが利点に変わる可能性があります。このデバイスのデザインの本当の問題は、そのひどい UI です。電源を入れた瞬間、地下鉄で携帯電話をにらみつけている老人のような歪んだ顔があなたの顔を叩きつけます。
私は友人たちに、デザイナーは若いに違いないと愚痴をこぼし続けています。彼らが設計実行時に犯した間違いは、とんでもなくひどいものです。すべてのアイコンがピクセル グリッドと完全に一致していないため、すべてがぼやけて見えます。
ご存知ないかもしれませんが、古いデザイナーが働いていた頃、画面解像度は低かったです。画像を鮮明で鮮明に保つために、サブピクセルのスムージングを引き起こすハーフピクセル上にポイントを配置することを避けます。ここでの問題

それは、多くのアイコンが非常に小さいことです。それらはぼやけて区別できない塊になります。あなたは頭の中が疑問符でいっぱいになりながら、永遠に彼らを見つめます。「これはいったい何なのだろう？」最も不合理な部分は、アイコンがぼやけていても、テキストは鮮明なピクセル フォントでレンダリングされていることです。
たとえば、ファームウェアをアップグレードしているとき、画面には、醜い Zhongyi Black および Song フォントで描かれたインターフェイスが表示されます。プロンプトテキストは画面の中央にさえありません。日常的に使用する場合、言語をフランス語に切り替えると、アクセント付き文字と通常の文字のサイズが異なることに気づくでしょう。すべてのキャラクターは不均一な高さで終わります。
さらに超現実的なものがあります。すべてのビットマップ UI アセットは、非可逆形式でエクスポートされ、BMP に変換されてからファームウェアに埋め込まれたように見えます。すべてのテクスチャには、明るい領域と暗い領域の間のエッジの周りに神秘的な白い輪郭があります。細部には圧縮アーティファクトとノイズが含まれています。つまり、平凡な物理的コントロールと組み合わせると、デバイスの使用は肉体的にも精神的にも不快になります。
もちろん良い点もあります。見た目も可愛くてとても小さいです。社交的な会話のネタとして持ち歩くことができます。私のお気に入りの機能は、ヘッドフォンジャックが2つあることです。バランス出力 1 つと標準 3.5mm ジャック 1 つ。同時に作業することもできます。学生時代に友達と 1 つのヘッドフォンを共有したことを思い出します。このデバイスを使用すると、同時に 3 人の友人と音楽を共有できます。 4人でのジャム演奏は面白いですね。
さらにクレイジーな使用方法もいくつかあります。
誰もが知っているように、ソニーのノイズキャンセリングヘッドフォンは基本的に低音を最大まで上げて頭が爆発する一方、高周波は濁って不快に聞こえます。私はたまたまフラットヘッド IEM (Y 型) を 2 台持っています。

最近購入したINCROWモデル）。高音はうまく処理できますが、低音が不足しています。そこで疑問が生じました。一方に高値を処理させ、もう一方に安値を処理させてもよいでしょうか?
低音を強化するために、Clear Bass を +10 まで設定しました。フラットヘッドは追加の EQ なしで動作しました。
両方のヘッドフォンを同時に装着し（耳にXM3を装着した状態で）、片方のプラグを順番に抜いたところ、興味深いことが起こりました。ソニーのプラグを抜くと、音はすぐに薄くなりました。ギターの倍音はまだ残っていますが、低周波のボディ共鳴は消えました。マイナスのプラグを外すと、ソニーの鈍さが明らかになり、すべてが濁りました。両方を一緒に演奏すると、ついに本物の音楽が現れました。
しかし、この設定にはコストがかかります。出力の位相が合っていないため、特に音が枯れたように感じられます。コンクリートのプールで泳いでいるようなものです。空間感覚が大きく損なわれます。
雰囲気としては、ハードウェアは良いがソフトウェアが悪い、という感じです。そしてその悪さは非常に一貫していて完全です。 UIデザインが悪いだけではありません。エンジニアリングの実践も悪いです。それについては後で話します。
その後、私は非常にイライラして（実際にはそれほど怒っていませんでした。ただアイデアがあっただけです）、ファームウェアを改造しようとし始めました。
私はリバースエンジニアリングの経験がほとんどありません。最後にこのようなことをしたのは中学生の時で、もうすべて忘れています。 「リバース エンジニアリング」に関する私の最高レベルの経験は、縮小された JavaScript コードを観察することです。この種の作業は、おそらく専門家でないと扱えるものではありません。
しかし、大規模な言語モデルがありますよね?
GLM 4.7 にファームウェアの全体的な形状を調べてもらうように依頼してみました。実際に、いくつかの有益な情報を分析することができました。
最初に私はこう言いました。「これがファームウェアです。私の目標は、ビットマップを抽出することです。

内側。やり方を教えてください。」
capstone、ghidra、r2pipe、rzpipe、angr が必要だと言われました。これらの Python バインディングを NixOS にインストールするのは少し面倒でしたが、最終的にはセットアップできました。
次に、バイナリ ファイルのスキャンを自動的に開始しました。やがて、チップモデルとSDKモデルが特定されました。パーティションテーブルも見つかりました。命令シーケンスの読み取りを試み、ある種のファームウェア暗号化を発見し、修正プログラムを作成しました (これは後で非常に重要になります)。
その後、ファイル内の埋め込みファイル システムに侵入し、すべての .bmp ファイルを抽出しました。実際に見てみると、イメージはまさに私が期待していたものと一致していました。
しかし、それらを抽出するには問題が伴いました。 BMP カラーエンコーディングが一致しません。ここで私は介入しなければなりませんでした。結局のところ、それには目はありませんが、私には目があります。私は、さまざまなパラメータを試し続ける間、調べる責任がありました。最終的に、私たちは正しいデコードパラメータを一緒に見つけることができました。
ピクセル フォントを使用するため、フォント ストレージは 0 が空白スペース、1 が文字のある部分であるバイナリ シーケンスであると最も単純に想定されます。そこで私は携帯電話を取り出し、画面上の漢字を写真に撮りました。次に、ピクセルを 1 つずつカウントし、検索サンプルとしてテキスト ドキュメントに入力し始めました。このフォントの大きなピクセルに感謝しなければなりません。一つ一つの点がはっきりと見えました。
私は抽出した文字を GLM 4.7 に渡して、「これらの数文字だけです。スキャンしてみてください。」
結果は驚くほど良好でした。さまざまなパディング方法を試してみると、ピクセル フォント データが保存されている場所が実際に特定されました。
グリフ ファイルは見つかりましたが、データ構造やインデックス テーブルがどこにあるのかはわかりませんでした。ここから長い暗中模索の日々が始まった。非常に素晴らしいものが 1 つありました

厄介な問題: 数文字ごとに破れが発生します。左半分は少し上に移動し、右半分は下に移動します。それは他のキャラクターごとに起こりましたが、理由はわかりませんでした。
ファームウェアのテキストレンダリングロジックを調べる以外に選択肢はありませんでした。キャラクターの居場所はすでにわかっていたので、その住所の周囲を探索し、最終的に何かを見つけました。 30 分もかからずに、レンダリング機能が見つかりました。それから、レンダリングの公式を見つけるまで、一歩ずつ上に登っていきました。
しかし、これではまだ画像の破れを説明できません。これを解決するために、ソフトウェアとハ​​ードウェアの通信プロトコル レベルでの問題を疑い始めました。 VOP DMA を含むミニ レンダリング シミュレーターを構築しましたが、それでも問題は見つかりませんでした。
何かが間違っていたに違いありません。私は行き止まりに陥ってしまい、抽出アルゴリズムに問題があるのではないかと考えていました。
引き裂きのパターンは明らかでした。左半分が上にシフトし、右半分が下にシフトし、1 文字おきに発生しました。そこで、GLM 4.7 にオフセット パターンを見つけるように依頼しました。非常に迷惑なことが次々と起こりました。多くの場合、大量の ASCII アートが出力されます。文字が明らかに歪んでいる場合でも、それが正しく完全な漢字であることを自信を持って教えてくれました。
しかし、キャラクターは明らかにバラバラでした。
LLM はテキストベースのモデルです。データを水平方向に 1 行ずつ読み取ることしかできません。彼らが目にしているのは、ドットとハッシュで構成された文字列です。彼らは統計的なパターンから何かが「漢字に似ている」と推測することができます。しかし、彼らには二次元のビジョンがまったくありません。彼らが見ているものは、人間の目で見ているものとはまったく異なります。グリフが正しいかどうかを判断するように依頼しても、視覚的に判断することはできません。確率を予測できますが、その予測は非常に簡単です。

ああ、間違えます。
また、統計モデリングも大好きです。たとえ比較対象がまったく同じものでなくても、「50% 一致」よりも「80% 一致」の方が優れていると考えられます。その時点で、このタスクには完全に正しいか完全に間違っているだけであることを辛抱強く説明する必要があります。 「半分正しい」状況はありません。
そんなことが何度も繰り返され、10回以上起こりました。私が「これは漢字ですか？破れていて形が間違っています。複数の文字を縫い合わせたように見えます」と言うたびに、謝罪し、再分析し、文脈が圧縮され、重要な教訓が不適切な圧縮プロンプトで失われます。私なら「行ごとではなく列ごとに読む必要がある」と明示的に伝えます。リッスンし、しばらくすると自動的に行読み取りモードに戻ります。もう一度言うと、また元に戻ります。あなたの指示は一時的なオーバーライドにすぎず、動作方法を実際に変更するものではありません。
この過程で、私は何度かカッとなって悪口を言い始めました。後になって、これは非常に特殊な問題であることに気づきました。会話に感情を注入すると、それは文脈に留まります。それは圧縮後も持続します (ただし、重要な方法論的な教訓は持続しません)。それは将来の推論の質を汚染し続けます。悪口を言ってからは、技術的な問題を解決するよりも、自分の感情を落ち着かせることに多くのエネルギーを注ぎ始めました。私の感情を刺激しないようにあらゆる方法を試みましたが、最終的には何も起こりませんでした。それは人間の闘争・逃走反応に非常に似ていました。その結果、理性が弱くなってイライラが募り、悪循環になってしまいました。
これは非常に実践的な教訓です。LLM と協力する場合、感情状態はエンジニアリング変数であり、個人的な問題ではありません。
その後、新しいコンテキストを開始し、LLM にビジュアルを作成するように依頼しました。

レンダリング結果を ROM アドレスにマッピングするツール。ブロックを 1 つずつ見て、新しいコンテキストで確認しました。 「このファームウェア ファイルはどこから入手したのですか?」と尋ねられました。 「修正台本」を見せました。修正スクリプトのせいで私の分析がすべて台無しになってしまったことがわかりました。ファームウェアが誤って解釈され、データ構造が破損していました。そこで文字ティアリングが登場したのです。
この時点で謎は解けました。
文字検索システムのリバースエンジニアリング
これまでのところ、CJK 範囲の文字のみが見つかりました。私たちは他の地域については何も知りませんでした。プレーン間には純粋なゼロ ブランク データの大きな領域があり、次にデータの連続セグメント、さらに大きなブランク、そして別の大きなデータの塊が見られました。 LLMは、データはブロックに保存されており、これらすべての純粋なゼロ領域がどこから来たのかについての説明を巧妙に避けたと述べた。
全体的なロジックによれば、データがセグメントに格納されている場合は、テーブルが存在するはずです。 「ルックアップテーブル」仮説を強制的に当てはめようとして、あらゆる種類の幻覚を生成し始めました。
ご存知のとおり、モデルは数学が非常に苦手です。いくつかのサンプル文字を与えたところ、「パターンを見つけよう」としましたが、パターンが存在しないと言い続けました。

[切り捨てられた]

## Original Extract

The Story Behind Reverse Engineering the Snowsky Echo Mini Firmware
I have always had a strange obsession with the media we use to store music. Back in elementary school I loved using our home tape recorder to make mix tapes. It left me with so many good memories. These days my dream is to own my own cassette Walkman.
But this is not easy to pull off. Walkmans are expensive and come with plenty of headaches. (My birthday wish this year was that some kind friend would gift me one. But I know you would call me shameless for even thinking it.)
To fill that deep consumerist void, half a year ago I settled for buying a Fiio Echo Mini. It is an MP3 player designed to look like a mini cassette machine.
After using it for a while I developed some complicated feelings. It is decent in some ways, but the overall execution is terrible.
The most painful thing to look at is that large pixelated screen. Fifteen years ago you might have considered it an average display. But in 2026, staring at that screen just feels odd. Even if the screen is pixelated, good visual design (like authentic retro pixel art) could have turned it into an advantage. The real problem with this device's design is its awful UI. The moment you power it on, a distorted face resembling an old man squinting at his phone on the subway slaps you right in the face.
I keep complaining to my friends that the designers must be young kids. The mistakes they made in design execution are ridiculously bad. Every single icon is completely unaligned to the pixel grid, so everything looks blurry.
In case you do not know, back when older designers were working, screen resolutions were low. To keep images sharp and clear they would avoid placing points on half-pixels that cause sub-pixel smoothing. The problem here is that many icons are very small. They blur into indistinguishable blobs. You stare at them forever with your head full of question marks: What the hell is this supposed to be? The most absurd part is that while the icons are blurry, the text is rendered with crisp pixel fonts.
For example, when you are upgrading the firmware the screen shows an interface drawn with ugly Zhongyi Black and Song fonts. The prompt text is not even centered on the screen! In daily use, if you switch the language to French you will notice that accented letters and regular letters have different sizes. All the characters end up at uneven heights.
There is something even more surreal. All the bitmap UI assets look as if they were exported in a lossy format, converted to BMP, and then embedded into the firmware. Every texture has a mysterious white outline around the edges between light and dark areas. The details are filled with compression artifacts and noise. In short, combined with the mediocre physical controls, using the device becomes physically and mentally uncomfortable.
Of course it has its good points. It looks cute and it is very small. You can carry it around as a social conversation piece. My favorite feature is that it has two headphone jacks! One balanced output and one standard 3.5mm jack. They can even work simultaneously! It reminds me of sharing a single pair of headphones with friends back in school. Now with this device you can share music with three other friends at once. Four people jamming together is hilarious.
There are also some even crazier ways to use it.
As we all know, Sony noise-cancelling headphones basically crank the bass to the maximum and blast your head while the high frequencies sound muddy and unpleasant. I happen to have a pair of flathead IEMs (the YINCROW model I recently bought). They handle highs well but lack bass. So the question became: Can I let one handle the highs and the other handle the lows?
To boost the bass I set Clear Bass all the way to +10. The flatheads ran without any extra EQ.
When I wore both headphones at once (flatheads inside my ears with the XM3s over them) and took turns unplugging one, something interesting happened. Unplugging the Sony made the sound instantly thinner. The guitar harmonics were still there but the low-frequency body resonance disappeared. Unplugging the flatheads made everything turn muddy as the Sony's dullness became obvious. When both played together, real music finally appeared.
But this setup has its costs. The output phases do not match, which makes the sound feel particularly dead. It is like swimming in a concrete pool. The sense of space suffers greatly.
The atmosphere is like this: The hardware is good but the software is bad. And the badness is very consistent and complete. Not only is the UI design bad; the engineering practices are also bad. I will talk about that later.
Later on I got so annoyed (actually I was not that angry, just had an idea) that I started trying to mod the firmware.
I have almost no experience with reverse engineering. The last time I did anything like this was in junior high and I have forgotten everything. My highest level of experience with "reverse engineering" is looking at minified JavaScript code. This kind of work is probably not something a non-professional can handle.
But we have large language models, right?
I tried asking GLM 4.7 to take a look at the general shape of the firmware. It actually managed to analyze some useful information.
At the very beginning I told it: "Here is a firmware. My goal is to extract the bitmaps inside. Tell me how to do it."
It told me it needed capstone, ghidra, r2pipe, rzpipe, and angr. Installing the Python bindings for these on NixOS was a bit troublesome, but I eventually got it set up.
Then it started scanning the binary file by itself. Before long it had identified the chip model and SDK model. It even found the partition table. It tried reading the instruction sequences, discovered some kind of firmware encryption, and wrote a correction program (this becomes very important later).
After that it climbed into the embedded file system inside the file and pulled out all the .bmp files. When I looked, the images really matched what I expected.
But extracting them came with problems. The BMP color encoding just would not match. This was where I had to step in. After all it has no eyes, but I do. I was responsible for looking while it kept trying different parameters. In the end we managed to figure out the correct decoding parameters together.
Because it uses pixel fonts, the simplest assumption was that the font storage is a binary sequence where 0 is blank space and 1 is the part with the character. So I took out my phone and photographed the Chinese characters on the screen. Then I started counting the pixels one by one and typing them into a text document as search samples. I have to thank the large pixels on this font. I could clearly see every single dot.
I gave GLM 4.7 the characters I had extracted and told it: "Just these few characters. Try to scan for them."
The result was surprisingly good. By trying various padding methods it actually located where the pixel font data was stored.
Although we found the glyph file, we did not know the data structure or where the index table was. This began a long period of groping in the dark. There was one very annoying issue: Every few characters there would be tearing. The left half would shift upward a bit while the right half shifted downward. It happened every other character and we had no idea why.
We had no choice but to look at the text rendering logic in the firmware. Since we already knew where the characters were, we explored around those addresses and eventually found something. In less than half an hour it located the rendering function. Then we climbed upward step by step until we found the rendering formula.
But this still could not explain the image tearing. To solve this we even started suspecting issues at the software-hardware communication protocol level. We built a mini rendering simulator involving VOP DMA, but we still could not find the problem.
Something had to be wrong. I was stuck in a dead end, thinking it was a problem with our extraction algorithm.
The tearing pattern was clear: left half shifted up, right half shifted down, happening every other character. So I asked GLM 4.7 to find the offset pattern. One very annoying thing kept happening. It would often output large amounts of ASCII art. Even when the characters were clearly distorted it would confidently tell me that they were correct and complete Chinese characters.
But the characters were obviously torn apart.
LLMs are text-based models. They can only read data line by line horizontally. What they see is strings made of dots and hashes. They can guess from statistical patterns that something "looks like a Chinese character." But they have no two-dimensional vision at all. What they see is completely different from what human eyes see. When you ask it to judge if a glyph is correct, it does not give you visual judgment. It gives you a probability prediction, and that prediction is very easy to get wrong.
Also, it loves statistical modeling. It thinks "80% match" is better than "50% match" even if what it is comparing is not the same thing at all. At that point you have to patiently explain that for this task there is only completely correct or completely wrong. There is no "half correct" situation.
This happened over and over, more than ten times. Every time I said "Are you sure this is a Chinese character? It is torn and the shape is wrong. It looks like multiple characters stitched together," it would apologize, re-analyze, the context would get compressed, important lessons would be lost in bad compression prompts. I would explicitly tell it "You need to read column by column, not row by row." It would listen, then after a while automatically revert to row-reading mode. Tell it again, and it reverts again. Your instructions are only temporary overrides, not real changes to how it works.
During this process I lost my temper several times and started swearing. Afterward I realized this was a very specific problem: Once you inject emotion into the conversation it stays in the context. It even persists through compression (but important methodological lessons do not). It keeps polluting future reasoning quality. After swearing at it, it started putting a lot of energy into calming my emotions instead of solving the technical problem. It tried every way to avoid triggering my feelings and eventually did nothing. It was very similar to a human fight-or-flight response. The result was that reasoning got weaker, I got more annoyed, and it became a vicious cycle.
This is a very practical lesson: When collaborating with LLMs, your emotional state is an engineering variable, not a private matter.
Later I started a fresh context and asked the LLM to create a visualization tool that mapped rendering results to ROM addresses. I looked at the blocks one by one and confirmed with the new context. It asked me: "Where did you get this firmware file from?" I showed it the "correction script." It told me that the correction script had messed up all my analysis. It had interpreted the firmware incorrectly and corrupted the data structures. That was why the character tearing appeared.
At this point the mystery was solved.
Reverse Engineering the Character Lookup System
So far we had only found characters in the CJK range. We knew nothing about other regions. We saw large areas of pure zero blank data between planes, then a continuous segment of data, more large blanks, then another big chunk of data. The LLM said the data was stored in blocks and cleverly avoided explaining where all those pure zero areas came from.
According to the overall logic, if data was stored in segments then there should be a table. It started generating all kinds of hallucinations trying to force the "lookup table" hypothesis to fit.
As you know, models are extremely bad at math. I gave it several sample characters and it tried to "find patterns" but kept telling me there were none and that

[truncated]
