---
source: "https://www.seangoedecke.com/text-ai-watermarks/"
hn_url: "https://news.ycombinator.com/item?id=49251153"
title: "Text AI watermarks will always be trivial to remove"
article_title: "Text AI watermarks will always be trivial to remove"
author: "gfysfm"
captured_at: "2026-08-10T23:25:43Z"
capture_tool: "hn-digest"
hn_id: 49251153
score: 2
comments: 0
posted_at: "2026-08-10T23:12:00Z"
tags:
  - hacker-news
  - translated
---

# Text AI watermarks will always be trivial to remove

- HN: [49251153](https://news.ycombinator.com/item?id=49251153)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/text-ai-watermarks/)
- Score: 2
- Comments: 0
- Posted: 2026-08-10T23:12:00Z

## Translation

タイトル: テキスト AI ウォーターマークの削除は常に簡単です

記事本文:
テキスト AI ウォーターマークを削除するのは常に簡単です Sean Goedeck
テキスト AI 透かしは常に簡単に削除できます
欧州連合 AI 法は、今から 1 か月後の 2026 年 8 月に施行が開始されます 1 。最大の新しい要件の 1 つは第 50 条で、すべての AI 出力が「人工的に生成されたものとして検出可能」であることを要求しています。言い換えれば、LLM プロバイダーが EU でビジネスを行いたい場合は、出力 2 にウォーターマーク (AI コンテンツを識別するために使用できる隠された署名) を適用する必要があります。
LLM テキストの透かしは興味深い問題です。最高のエンジニアリング問題と同様、この問題も理論的には完全に解決するのは困難ですが、部分的な解決策は複数あります。たとえば、Google の SynthID や、(これから主張するように) OpenAI や Anthropic による静かな Unicode の策略などです。 AI ラボが年末までにこれらのトレードオフをどのように乗り越えるかを見るのは興味深いでしょう。
私は昨年末に「AI 検出ツールはテキストが AI によって生成されたものであることを証明できない」で AI 透かしについて書きました。デジタル画像には人間の目には見えないノイズが多く含まれているため、画像に透かしを入れるのは簡単です。たとえば、「これらの正確なスポットのこれら 20 ピクセルは常に同じ色を共有します」のような透かしを適用できます。テキストはずっとずっと難しいです。画像とは異なり、テキストは非常に圧縮されたメディアです。人間が気付かないような変更を文に加えることができません (1 つの例外を除き、これについては後ほど説明します)。では、どうやって透かしを入れるのでしょうか？
これは基本的にテキスト ステガノグラフィー (秘密コードの隠蔽) の問題ですが、平文は恣意的に操作できないため、より困難になります。ウォーターマークを適用するために変更を加えると、出力の品質が低下します。たとえば、「5 文字ごとに「e」が表示されます。

「」は良い透かしですが、単純に適用すると AI 出力がタイプミスだらけになってしまいます。透かしを適合させる方法をモデルに理解させられますか? 強力な AI モデルは、この種の制約をうまく処理できるほど賢いものです 3 が、それでもユーザーの問題に費やすべき推論時間が消費され、モデルの能力が実際よりもはるかに低いように見えます 4 。
AI コンテンツを検出するにはウォーターマークが必要ですか?
本当に透かしが必要ですか?あなたが人間的で、モデルが特定のテキスト ブロックを生成したかどうかを検証する必要がある場合、単に各モデルにテキストを実行し、モデルの予測されたトークンがテキストの各トークンとどの程度一致するかを測定しながら実行することはできないでしょうか?
そうではありません。 「質問に対するクロード ソネットの可能なすべての回答」のスペースは、「質問に対する透かし付きの可能なすべての回答」のスペースよりもはるかに大きいです。言い換えれば、AI が書いたかのように読める人間のテキストでは、誤検知が多すぎるということです。人間がうっかり透かしを再現してしまうよりも、人間がうっかりクロードのように書いてしまう可能性のほうがはるかに高いのです。
また、テキストに透かしを入れるために、テキストに対してすべての Anthropic モデルを実行すると、法外なコストがかかります。 EU AI 法は最終的に、Anthropic のような研究所に対し、すべての EU 国民に無料の透かしサービスを提供することを義務付けることになります (コミットメント 2 を参照)。 「モデルを実行する」アプローチではこれを行うことはできません。
私の知る限り、テキスト出力に透かしを入れると述べている唯一の AI プロバイダーは Google です。Google は SynthID と呼ばれるツールを使用しています。その仕組みは次のとおりです。
LLM がテキストを生成するとき、一連のトークン (単語または単語の塊) が生成されます。各ステップで、モデル自体は 1 つのトークンを出力するのではなく、(たとえば) 100,000 個のトークンすべての完全なリストを出力します。

その語彙には、そのトークンが次のトークンになる確率が注釈として付けられています。 ChatGPT や Claude Code などのツールは、出力を取得するために最も可能性の高いオプションから半ランダムに選択します。この半ランダム サンプリング プロセスは、検出可能な方法で影響を受ける可能性があります。
たとえば、「2 番目に可能性の高いトークンを選択し、次に最初のトークンを選択し、次に 2 番目のトークンを選択し、次に最初のトークンを選択する、というようなサンプリング戦略を選択できます。」これでも高品質の出力が生成されますが、生成されたテキストに対してモデルを再実行して、パターンが保持されていることを確認することができます。ただし、これでは検証に非常にコストがかかり、出力をわずかに調整するとパターンが崩れ、フィンガープリントが破壊されてしまいます。もっと良い方法はありますか?
はい。 SynthID は、以前のトークンに基づいて各トークンに「スコア」を割り当てるプロセスです (たとえば、トークンの ID とその前の 3 つのトークンの ID を合計し、mod 5 を取得します) 5。ウォーターマークを適用するために、モデルは「最も可能性の高い上位 5 つのトークンから、SynthID スコアが最も高いものを選択する」ようなサンプリング戦略を採用します 6 。透かしは、テキスト ブロックの集計 SynthID スコアを計算することで検出できます。疑わしいほど高い場合は、AI によって生成された可能性が非常に高くなります。
これは基本的に、 em-dash を使用して LLM を識別できるという一般的なアドバイスのバージョンです。ただし、キーワードのリストではなく、人間が識別できない単語間の微妙な数学的関係に依存している点が異なります。スコアを割り当てるプロセスは簡単であるため、透かし検出の実行は非常に安価です。
ホモグリフによる Unicode 透かし
Google には、SynthID がモデルを愚かにしない理由について、複雑な数学的根拠があります。おそらく、SynthID のスコアリングは十分にランダムであるため、

通常の擬似ランダム トークン サンプラーのように動作し、出力に検出可能な指紋を残すだけです。しかし、もちろんこれは疑わしい。たとえば、温度をゼロに設定して推論を行うのが一般的で、これにより常にモデルの最も可能性の高い次のトークンが選択されます。その場合、指紋をまったく残すことはできません (または、ユーザーの好みを無視して、とにかく 2 番目または 3 番目の選択肢を選択する必要があります)。
モデルの出力を変更できない場合でも、コンテンツのフィンガープリントを取得できますか?そうですね。 OpenAI と Anthropic が時々派手な Unicode トリックを適用していることは間違いありません。たとえば、通常の ” ” スペース (unicode U+0020 ) を 3 つずつの ” ” スペース (unicode U+2004 ) または CJK 表意文字 ” ” スペース (unicode U+3000 ) に置き換えることができます。これらは「ホモグリフ」と呼ばれ、ここでさらに詳しく見つけることができます。
もちろん、人間が作成したテキストの多くでは同形文字が使用されています。しかし、実際には発生する可能性がはるかに低い同形文字のパターン (たとえば、「3 つおきのスペースが 3 つずつになる」) をエンコードするのは簡単です。 SynthID ウォーターマークと同様、ホモグリフベースのウォーターマークは非常に安価に検出できます。ホモグリフベースのウォーターマークは、SynthID よりも適用コストが低く、完全にクライアント上で実行することもできます。
これは陰謀論ではないと思います。 Claude Code は、中国人ユーザーからの不審なリクエストにタグを付けるためにこれを行っていたのは間違いありません (「今日の日付」の ' の同形文字を悪用していましたが、その後、撤回しています)。ここ数年、ChatGPT からテキストのブロックをコピーして VSCode に貼り付けると、VSCode が一部またはすべてのスペースを異常な Unicode 文字としてマークする場合があることに気づきました 7 。 OpenAI と Anthropic は AI が生成した透かしとしてホモグリフを使用していますか?よくわからない。でも、

彼らは間違いなく同形文字を使用しています。
テキストの透かしは簡単に削除できます
AI 法 (具体的には、それに関連する実施規範) では、透かしを「コンテンツから分離することが困難な方法でコンテンツ内に埋め込む」ことを義務付けています。ただし、テキストの透かしは簡単に削除できます。
Unicode ホモグリフの透かしを削除するには、すべてのホモグリフを「実際の」文字に相当するものに置き換えるだけです。比較的弱い透かしの入っていない LLM 8 にさえアクセスできる場合は、その LLM にテキスト コンテンツの言い換えを依頼することで、SynthID の透かしを取り除くことができます。透かしは微妙な語彙の選択に固有のものであるため、コンテンツを言い換えると透かしが削除されます。手動で行うこともできますが、その時点では、実際には AI によって生成されたコンテンツではなくなります。無料のパブリックウォーターマークテストツールのようなものがあるので、陰性が返されるまで微調整を続けることができます。
さらに、AI 法では、透かし技術は「技術的に可能な限り相互運用可能」であることが求められています。つまり、AIプロバイダーは透かし入れのプロセスを公開する必要があり、場合によっては同じ種類の透かしの適用を標準化しようとすることさえあるだろう。これが、LLM テキスト透かしが依存する種類の隠蔽によるセキュリティとどのように互換性があるのか​​わかりません。画像やビデオの透かしとは異なり、テキストの透かしは常に簡単に削除できます。
AI 法と実施規範では、「デジタル署名されたメタデータ」について多くのことが語られています。ここでの考え方は、理想的には改ざんできない方法 (たとえば、ファイルの内容のハッシュに署名するなど) で、ファイルのメタデータ自体に AI 開示を含めることができるということです。この署名付きメタデータ プロセスは、基本的には C2PA コンテンツ資格情報です。できるうちに

C2PA メタデータを削除すると、(理論的には) 偽造できないため、「人間によって作成された」メタデータを含むファイルは信頼でき、メタデータがまったく含まれていないファイルは疑わしいと考えられます。
この投稿はすでに長すぎて、C2PA について私が考えていることを説明するには時間がかかりますが、C2PA はテキストの透かしの代替ではない、ということだけは言っておきたいと思います。実際に適用されるのはファイルのみです。実践規範の言葉を借りれば、これは「メタデータ (オーディオ、画像、ビデオ、コンテナ化されたテキストなど) の添付をサポートするデータ形式」です。チャット ツールの出力 (および AI エージェントの出力のほとんど) はコンテナ化されたテキストではなく、古い通常のテキストであるため、署名できません。 ChatGPT 出力に署名するとどうなるでしょうか?渡すアーティファクトはありません。
Claude Code が生成する HTML ファイルや PDF に C2PA 署名する必要があるかどうかは、興味深い問題だと思います。それを正しく理解するのはちょっと難しいようです。しかし、いずれにせよ、AI 法では、ある種の実際の透かし入れも義務付けられています。
さて、今年は何が起こるでしょうか？推測する必要があるとすれば、各 AI プロバイダー (OpenAI や Anthropic などのラボだけでなく、Fireworks や Groq などのサードパーティ プロバイダー) は、推論スタックの前に SynthID トークン サンプラーを貼り付けることになるでしょう。これは EU 内のユーザーに限定される可能性がありますが、SynthID は通常の上位 K トークン サンプリング アプローチと少なくとも同じくらい優れているため、そうではない可能性があります。
AI プロバイダーは、ユーザーが提供したテキストを再トークン化し、スコアリングを実行し、それが特定のしきい値を超えているかどうかをチェックする「透かしのチェック」ページを提供します。相互運用性条項をどの程度真剣に受け止めるかによっては、プロバイダーが同じ SynthID 設定を標準化する可能性さえあります。その場合、EU がホストする単一の「このテキストに透かしを入れる」ページが存在する可能性があります。
Unicodeベースのwatermaではないと思います

rking は AI 法に準拠しているとみなされますが、SynthID を設定したくないプロバイダーもこれを試みる可能性があります。いずれにせよ、技術者ユーザーはウォーターマークを自由に取り除くことができ、非技術者ユーザーがこの目的で使用できるツールが多数存在することになります。
さて、新しいシステムの場合です。既存のものは12月までです。
第 50 条の平文ではこれが要求されていないと思いますが、第 133 条と実施規定を見ると、透かしを求めていることは明らかです。
非常に高度に考えても、GPT-5.5 は SynthID が 5 文字ごとに「e」であることを私に説明することはできませんでしたが、GPT-5.5-Pro は次の不可解な公案を生成しました: 「これらの隠されたコードはモデルで作られた画像、音声、映画、散文にラベルを付けます。トレースを調べる: おそらくモデルで作られた作品。おそらくトレースを消去します。おそらくトレースを残します。したがって、トレースだけでしょうか? いいえ。」
AI の安全ガードレールとの類似点については、読者の演習として残しておきます。
それはおもちゃの例です。実際には、ランダム シードなど、複数の異なる (ただし数学的には単純な) スコアリング方法が組み合わされています。なぜ種子が含まれるのでしょうか？そうしないと、ウォーターマークは同じトークンのセットに偏ることになります。
トークンは相互に複数ラウンドのノックアウトで得点されますが、それは実装の詳細であり、SynthID が機能する理由の背後にある核となる直感を理解するためには必要ないと思います。
これが公知になったとき、OpenAI は、これは単なるモデルであると主張しました。

[切り捨てられた]

## Original Extract

Text AI watermarks will always be trivial to remove sean goedecke
Text AI watermarks will always be trivial to remove
The European Union AI Act will begin to be enforceable in August 2026, one month from now 1 . One of the biggest new requirements is Article 50 , which requires all AI outputs to be “detectable as artificially generated”. In other words, if LLM providers want to do business in the EU, they will have to apply a watermark to their outputs 2 : some hidden signature that can be used to identify AI content.
LLM text watermarking is a fascinating problem. Like the best engineering problems, it is theoretically hard to solve perfectly, but has multiple partial solutions: for instance, Google’s SynthID , and (as I’ll argue) some quiet Unicode trickery from OpenAI and Anthropic. It will be interesting to see how the AI labs navigate these tradeoffs before the end of the year.
I wrote about AI watermarking at the end of last year in AI detection tools cannot prove that text is AI-generated . It’s easy to watermark an image, because digital images contain lots of noise that the human eye can’t really see. For instance, you could apply a watermark like “these twenty pixels in these exact spots will always share a color”. Text is much, much harder. Unlike images, text is a very compressed medium: you cannot make any change to a sentence that a human wouldn’t notice (with one exception, which we’ll get to later). So how are you supposed to watermark it?
It’s basically a text steganography problem (concealing a secret code), made more difficult because the plaintext cannot be arbitrarily manipulated. Any changes you make to apply the watermark will compromise the quality of the output. For instance, “every fifth letter is an ‘e’” would be a good watermark, but applied naively would make the AI output full of typos. Could you just let the model figure out how to fit the watermark? Strong AI models are smart enough to juggle this kind of constraint 3 , but it’d still consume reasoning time that would be better spent on the user’s problem, and make the model sound much less capable than it is 4 .
Do we need watermarks to detect AI content?
Do you really need a watermark? If you’re Anthropic, and you’re required to be able to verify whether your models produced a particular block of text, can’t you simply run the text through each model, measuring as you go how closely the model’s predicted tokens match each token from the text?
Not really. The space of “all possible Claude Sonnet answers to a question” is way larger than the space of “all possible watermarked answers to a question”. In other words, you’d get too many false positives for human text that reads like it was AI-written. It’s way more likely for a human to accidentally write like Claude than it is for a human to accidentally reproduce a watermark.
It would also be prohibitively expensive to run every Anthropic model against a piece of text in order to watermark it. The EU AI Act will eventually require labs like Anthropic to offer free watermarking services to every EU citizen (see Commitment 2). You couldn’t do that with the “run the model” approach.
As far as I know, the only AI provider to say they watermark text output is Google, who use a tool called SynthID . Here’s how it works.
When an LLM generates text, it’s generating a series of tokens (words or chunks of words). At each step, the model itself doesn’t output a single token, but instead outputs a full list of all (say) 100,000 tokens in its vocabulary, each annotated with the probability that that token will be the next one. Tools like ChatGPT or Claude Code will pick semi-randomly from the most likely options in order to get their outputs. This semi-random sampling process can be influenced in a detectable way.
For instance, we could choose a sampling strategy like “we pick the second most likely token, then the first, then the second, then the first, and so on”. That would still produce high-quality output, but you’d be able to re-run the model against the generated text to verify that the pattern holds. However, that’d make verification really expensive, and any slight tweaks to the output would break the pattern and thus break the fingerprint. Is there a better way?
Yes. SynthID is a process for assigning each token a “score” based on its previous tokens (for instance, sum the token’s ID with the IDs of its previous three tokens then take mod 5) 5 . To apply the watermark, the model adopts a sampling strategy like “out of the top five most likely tokens, pick the one with the top SynthID score” 6 . The watermark can then be detected by calculating the aggregate SynthID score of a block of text. If it’s suspiciously high, it’s very likely to have been AI-generated.
This is basically a version of the common advice that you can identify LLMs by use of the em-dash , except that instead of a list of keywords, it relies on subtle mathematical relationships between words that humans can’t identify. Because the process for assigning the score is trivial, it’s very cheap to run watermark detection.
Unicode watermarks via homoglyphs
Google have a complicated mathematical rationale for why SynthID doesn’t make the model dumber: supposedly the SynthID scoring is random enough to act like a normal pseudo-random token sampler, just one that leaves a detectable fingerprint on the outputs. But of course this is suspicious. For instance, it’s common to do inference setting temperature to zero, which always picks the model’s most likely next token. In that case, you can’t leave a fingerprint at all (or you have to ignore the user’s preference and pick the second or third choice anyway).
If you can’t alter the model outputs, can you still fingerprint the content? Well, kind of. I’m pretty sure OpenAI and Anthropic are sometimes applying fancy Unicode tricks. For instance, you might go through and replace your normal ” ” spaces (unicode U+0020 ) with a three-per-em ” ” space (unicode U+2004 ), or a CJK ideographic ” ” space (unicode U+3000 ). These are called “homoglyphs”, and you can find more of them here .
Of course, lots of human-generated text uses homoglyphs. But it’s trivial to encode a pattern of homoglyphs (say, “every third space becomes a three-per-em”) that is much less likely to occur in the wild. Like the SynthID watermark, a homoglyph-based watermark can be detected very cheaply. A homoglyph-based watermark is cheaper to apply than SynthID: you could even do it entirely on the client.
I don’t think this is a conspiracy theory. Claude Code was definitely doing this to tag suspicious requests from Chinese users (exploiting homoglyphs for the ’ character in “Today’s date”, though they’ve since walked that back). In the last few years, I’ve noticed that when I copy blocks of text from ChatGPT and paste them into VSCode, sometimes VSCode marks some or all of the spaces as unusual Unicode characters 7 . Are OpenAI and Anthropic using homoglyphs as an AI-generated watermark? I’m not sure. But they’re definitely using homoglyphs.
Text watermarks can be trivially removed
The AI Act (specifically, its associated Code of Practice ) requires watermarking to be “embedded within the content in a manner that is difficult for it to be separated from the content”. However, text watermarks can be trivially removed.
To remove unicode homoglyph watermarking, you simply have to replace all the homoglyphs with their “real” character equivalents. If you have access to even a relatively weak un-watermarked LLM 8 , you can strip out SynthID watermarking by asking that LLM to paraphrase the text content. Because the watermark is inherent to subtle vocabulary choices, re-wording the content will remove the watermark. You could even do it by hand, although at that point it’s not really AI-generated content anymore. Since there will be some kind of free public watermark testing tool, you can just keep tweaking until it comes back negative.
Moreover, the AI Act requires watermarking techniques to be “interoperable… as far as this is technically feasible”. That means AI providers would have to publish their watermarking process, and potentially even attempt to standardize on applying the same kind of watermarks. I just don’t see how this is compatible with the kind of security-by-obscurity that LLM text watermarking depends on. Unlike image and video watermarks, text watermarks will always be trivial to remove.
The AI Act and Code of Practice talk a lot about “digitally signed metadata”. The idea here is that you can include an AI disclosure in the file’s metadata itself, ideally in a way that cannot be tampered with (for instance, by signing a hash of the file’s contents). This signed-metadata process is basically C2PA Content Credentials . While you can remove C2PA metadata, you (theoretically) can’t fake it, so a file with “created by a human” metadata can be trusted, and files with no metadata at all can be held in suspicion.
This post is already too long to get into what I think about C2PA, but I do want to say that C2PA is not a substitute for text watermarking . It only really applies to files . In the words of the Code of Practice, that’s “a data format that supports attaching metadata (e.g., an audio, image, video, or containerised text)“. The output of chat tools (and most of the output of AI agents) is not containerized text, but plain old regular text, and so can’t be signed. What would it even look like to sign ChatGPT outputs? There’s no artifact to pass around.
I think it’s a fascinating question whether Claude Code has to C2PA-sign any HTML files or PDFs it generates for you. That seems kind of tricky to get right. But in any case, the AI Act also mandates some kind of actual watermarking as well.
So what’s going to happen this year? If I had to guess, I’d say that each AI provider (not just labs like OpenAI or Anthropic, but third-party providers like Fireworks or Groq) will stick a SynthID token sampler in front of their inference stacks. This might be limited to users in the EU, but it might not be, since SynthID is at least as good as a normal top-k token sampling approach.
AI providers will then offer a “check for watermark” page that re-tokenizes user-provided text, runs the scoring, and checks whether it’s above a certain threshold. Depending on how seriously the interoperability clause is taken, providers might even standardize on the same SynthID setup, in which case there could be a single EU-hosted “watermark this text” page.
I don’t think unicode-based watermarking is going to be considered compliant with the AI Act, but some providers which don’t want to set up SynthID might try it. Either way, technical users will be able to strip out the watermark at will, and there will be a plethora of tools that non-technical users will use for this purpose.
Well, for new systems; existing ones get until December.
I don’t think the plain text of Article 50 requires this, but Recital 133 and the Code of Practice makes it pretty clear that they’re looking for watermarks.
Even with extra high thinking, GPT-5.5 could not explain SynthID to me with every fifth letter being an “e”, but GPT-5.5-Pro produced this puzzling koan: “These hidden codes label model-made image, voice, movie, prose. Probe trace: maybe a model-made piece. Maybe erase trace; maybe leave trace. Hence trace alone? No.”
I leave the analogy with AI safety guardrails as an exercise for the reader.
That’s a toy example. In practice there are multiple different (but still mathematically simple) scoring methods that get combined together, including a random seed. Why include the seed? Otherwise the watermark would bias towards the same set of tokens.
The tokens are scored in a multi-round knockout against each other, but I think that’s more of an implementation detail and not required to get the core intuition behind why SynthID works.
When this became public knowledge , OpenAI claimed it was just a model q

[truncated]
