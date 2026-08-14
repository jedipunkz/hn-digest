---
source: "https://www.anthropic.com/news/claude-text-watermark"
hn_url: "https://news.ycombinator.com/item?id=49303350"
title: "How Claude's text watermarking works"
article_title: "How Claude's text watermarking works \\ Anthropic"
author: "surprisetalk"
captured_at: "2026-08-14T19:42:16Z"
capture_tool: "hn-digest"
hn_id: 49303350
score: 3
comments: 0
posted_at: "2026-08-14T19:15:51Z"
tags:
  - hacker-news
  - translated
---

# How Claude's text watermarking works

- HN: [49303350](https://news.ycombinator.com/item?id=49303350)
- Source: [www.anthropic.com](https://www.anthropic.com/news/claude-text-watermark)
- Score: 3
- Comments: 0
- Posted: 2026-08-14T19:15:51Z

## Translation

タイトル: クロードのテキスト透かしの仕組み
記事のタイトル: クロードのテキスト透かしの仕組み \ Anthropic
説明: 将来のクロード モデルは、透かしを含むテキストを生成します。これは、クロードがテキストの執筆に関与した可能性を判断する方法であり、私たちは他のいくつかの主要な AI プロバイダーとともに、EU AI 法に準拠するためにこの変更を実装しています。この記事では、その答えを共有します
[切り捨てられた]

記事本文:
クロードのテキスト透かしの仕組み \ Anthropic メインコンテンツにスキップ フッターにスキップ 研究
クロードのテキスト透かしの仕組み
将来のクロード モデルは、透かしを含むテキストを生成します。これは、クロードがテキストの執筆に関与した可能性を判断する方法であり、私たちは他のいくつかの主要な AI プロバイダーとともに、EU AI 法に準拠するためにこの変更を実装しています。
この記事では、選択した透かし手法がどのように機能するか、それがクロードの出力に影響するかどうか、およびこの変更を行う理由について、寄せられた質問の一部に対する回答を共有します。要約すると:
クロードの出力の品質や内容に実質的な影響を与えない透かしの方法を使用します。
透かしの入ったテキストと透かしの入っていないテキストの違いは、読者には区別できません。
テキストには何も追加されず、隠し文字もありません。
透かし入れには追加のトークンは必要なく、費用もかかりません。
透かしには識別情報が含まれず、特定の個人、組織、またはチャットを追跡することはできません。
透かしはクロードに固有のものではありません。 8 月 2 日の時点で、EU は市場にサービスを提供する AI プロバイダーに対し、AI で生成されたコンテンツをマークすることを義務付けています。他の主要なモデル開発者も同じ実施規範に署名しており、独自のウォーターマークを実装する予定です。
クロードのような大規模な言語モデルは、一度に 1 つの単語を生成することで機能します。モデルは次の単語を決定するたびに、潜在的な候補のリストから選択し、最終的には前のテキストに基づいて最も合理的または可能性の高い単語を選択します。 「今日の天気は寒かったし…」という文を考えてみましょう。次の単語が「甘い」である可能性は非常に低いです。しかし、「曇り」または「灰色」になる可能性が非常に高いです。ほとんどのサーキットの下で

おそらく、モデルが最後の 2 つの単語のどちらを最終的に選択するかは、読者にとってはあまり重要ではありません。どちらの場合でも、文の意味はほぼ同じです。このような場合、選択は乱数によって決定されます。
透かしは、このような危険性の低い選択肢を使用して (生成されたテキストの一部に何度も発生します)、クロードの応答にパターンを残します。そのパターンはリーダーには検出できませんが、それをエンコードするキーを持っている人には検出可能です。透かしを使用する場合でも、選択はランダムに行われますが、ランダム性の原因は異なります。次の単語を選択するために任意の乱数ジェネレーターを使用する代わりに、ウォーターメイキングでは、キーとその前にあるいくつかの単語を使用して、モデルがどの単語を選択するかを決定します。つまり、クロードが選択する単語はまだランダムですが、単語の順序をチェックして、キーを使用した場合にクロードが行うであろう選択と一致しているかどうかを確認できるようになりました。そうであれば、そのテキストがクロードによって生成された可能性を割り当てることができます。
重要なのは、モデルが常に曇りまたは灰色に偏るわけではないということです。透かしのないテキストと同様に、前の単語に応じて、ある文では曇りが選択され、次の文では灰色が選択される場合があります。そして、透かし入れの方法がクロードに、考えもしなかった単語の選択を強制するわけではありません（たとえば、クロードが「nubilous」のような単語を選択することはありません。これは、クロードが通常の状況ではほぼ間違いなく使用しない曇りまたは灰色のあいまいな同義語です）。
クロードの出力にはどのような影響がありますか?
透かしはクロードの出力の品質には影響しません。読者にとって、透かし入りの応答は透かしなしの応答と区別がつきません（このように、AI 透かしはサブメッセージと異なります）

紙幣、その他の物理的オブジェクト、および肉眼で見える一部のデジタル文書に記載されている同名の名前から明らかです)。
内部テストでは、クロードのテキストの内容、創造性のレベル、読みやすさに対する透かしの影響は確認されていません。私たちが使用している技術を紹介した SynthID-Text 論文 で、Google DeepMind は、透かしを使用したモデルを Gemini トラフィックの一部に提供し、高評価と低評価を比較することで、この影響をテストしました。彼らは、透かしを入れていないモデルと統計的に有意な差を発見しませんでした。また、対照研究では、人間の評価者が透かし入りの回答と透かしなしの回答を並べて比較したところ、品質に違いは見られませんでした。
わかりやすい例として、モノポリーのようなゲームをプレイしていると想像してください。各ターンで、各プレイヤーはサイコロの目に従ってボードの周りのランダムな数のスペースを移動します。このランダム性を得るためにサイコロを振る代わりに、円周率の桁の本を使用することにしたとします。 2 ランダムに選ばれた数字 (たとえば、小数点以下 1,012,845 番目、たまたま 6) から開始し、その時点から各プレイヤーは単純にシーケンス内の次の数字を次の「ロール」として使用します。
どう見ても、動きはランダムであることに変わりはありません。ランダム性が毎回の円周率によるものなのか、サイコロの出目によるものなのかは、プレイヤーにとって、あるいはゲームの結果にとっては何の違いもありません。しかし、ゲーム後にすべての動きの順序を見ることができれば (そして円周率の値がわかっていれば)、これが手を決定するために円周率を使用する可能性が高いゲームであるかどうかを判断することができます。円周率を使用したゲームは、ある意味「透かし入り」です。
クロードが生成したテキストについても同様です。透かしは、それを読む人の意味や経験を変えるものではありませんが、透かしを入れたい場合は、

ウォーターマークを使用すると、テキストがクロードによって生成された可能性が高いかどうかを事後的に確認できます。
透かしを入れる具体的な方法はどれですか?
クロードのテキスト ウォーターマークは、2024 年に Google DeepMind が Nature 論文で発表した SynthID-Text アプローチのバージョンです。これは、2022 年の Scott Aaronson による提案に遡る一連のアプローチに属しており、そのすべてが上で説明した同じ設計原則を共有しています。ウォーターマークは、単語の中から選択するために使用されるランダム性のソースを変更するだけです。
透かしの有効性には制限があります。私たちのキーを使用すると、「これが部分的にクロードによって書かれた可能性はどれくらいですか?」という質問に答えることしかできません。テキストが人間によって書かれたものかどうかは確認できませんし、そのテキストが別の AI によって書かれたかどうかもわかりません (たとえ他の AI が透かしを使用しているとしても、キーは異なるでしょう。また、まったく別の透かし方法を使用している可能性もあります)。透かしの検出は、単語の選択肢が少なく、続く情報も少ない小さなサンプルではうまく機能しません。文章が長くなるにつれて、クロードの関与に対する確信も高まります。
テキストの精度を低下させることなく実行できる選択肢が少ない事実に基づく文章では、透かしはよりまばらになります。たとえば、「アイザック ニュートンの最も有名な作品はプリンキピアと呼ばれていました…」という文を考えてみましょう。次の単語が「 Mathematica 」であるかどうかが非常に重要なので (これが唯一の正解です)、ウォーターマークは何も作用しません。校正についても同様です。クロードに文章を渡し、文法と句読点だけを編集して他には何も編集しないように依頼した場合、透かしは少数の修正箇所のみに存在し、登録するには少なすぎる可能性があります。
なんて腹筋

クロードが人間の文章を校正または編集したケースはありますか？
透かしはクロードが選択した単語にのみ適用されます。クロードが人によって書かれたテキストを校正すると、返される内容は通常、軽く編集されただけです。ほぼすべての言葉がその人の言葉なので、透かしを付ける部分は（あったとしても）ほとんどありません。テキストの長さとクロードがどの程度編集したかによっては、それらの変更だけではクロードの関与を検出できるほど十分ではない可能性があります。クロードが書けば書くほど、より多くの決定を下す必要があり、透かしを入れるスペースが増えます。
上で述べたように、AI 透かしは、どちらの単語を選択しても同様に適切であるという決定を利用します。正確な出力が必要な場合、つまり選択の余地がなく、何かが事実上間違っている場合や、別の用語が選択されるとコードの一部が破損する場合には、ウォーターマークは適用されません。
たとえば、モデルが「2 + 2 =」を書き込むと、次のトークンには明確な最適な選択肢が存在します (モデルが合計を計算している場合、「4」と同じくらい良い答えはありません。ジョージ オーウェルの 1984 について話している場合、「5」と同じくらい良い答えはありません)。ウォーターマークの「ナッジ」はここでは適用されません。同じ理由で、コード (多くの場合、正確でなければなりません) は、他の形式のテキストよりも透かしが少ないのが一般的です。
そうは言っても、コード内のコメントなど、コード内の特定の単語や用語を任意に選択できる領域では、ウォーターマークを使用できます。ただし、定義上、生成される実際のコードに与える影響はごくわずかです。
これはユーザーにとって何を意味しますか?
これによりモデルの速度が遅くなりますか、それとも価格が高くなりますか?
いいえ、透かしには n があります

モデルの速度に明らかな影響があり、追加のトークンが生成されないため、モデルの提供と使用の価格は同じです。
ウォーターマークは私または私の組織に遡ることができますか?
いいえ。透かしはクロードとその出力に適用されます。個々のユーザーとの関係を特定するものではありません。透かしやそのキーには、ユーザー、その組織、またはクロードとのチャットに関する情報を復元できるようなものは何もありません。
なぜクロードの出力に透かしを入れるのですか?
EU AI 法に準拠するためにウォーターマーキングを実装しています。 Anthropic は、他の主要な AI モデル プロバイダー数社および合計約 190 社の署名者とともに、2026 年 7 月に AI 生成コンテンツの透明性に関する EU 実践規範に署名しました。これにより、AI システム プロバイダーは AI 生成テキストを「マーキング」する方法を使用することが求められます。地域ごとに透かしを適用する永続的な方法がまだないため、リリース時に透かしをグローバルに適用しています。ただし、今後もさまざまなアプローチの評価を継続し、最新情報が得られたら共有します。
テキストがクロードによって書かれたものであるかどうかを確認するにはどうすればよいですか?
間もなく透かし検出 API を提供する予定です。現在、実装の詳細を検討中です。
画像やその他のファイルはどうなりますか?
Claude がサポートされているタイプのファイル (.png、.jpg、または .svg など) を生成すると、そのファイルのメタデータに、そのファイルが Claude で作成または処理されたことを示す、暗号的に署名された小さなメモの形式でコンテンツ資格情報が添付されます。これは C2PA と呼ばれるオープン業界標準であり、カメラ メーカーや写真編集ソフトウェアで画像の出所を記録するために使用されているものと同じです。 C2PA 対応ツールならどれでも読み取ることができます。ファイルをドロップして確認できる独自のファイルを提供する予定です。
このメタデータ ラベルは ve です

ウォーターマークとは異なります。ファイル内の内容は何も変更されません。ファイルは埋め込まれたり、非表示になったりしません。テキストと同様に、資格情報には、クロードがファイルの作成に関与したことが記載されているだけです。識別情報は含まれません。
誰かがテキストを編集するだけで透かしを回避できないのでしょうか?
ある程度はそうです。軽い編集ではおそらくウォーターマークを完全に削除することはできません。すべての単語が置き換えられる完全な書き換えが行われます。もちろん、後者の場合、そのテキストがもはや AI によって生成されたものと言えるかどうかには議論の余地があります。
透かしは実際に何を証明するのでしょうか?
透かしは、クロードがある時点でコンテンツに関与した可能性があることを判断することしかできません。 「クロードがこれを書いた」と「クロードがこれを大幅に編集した」を区別することはできません。
透かしは翻訳に適用されますか?
はい。クロードが作成した翻訳には透かしが入っています。これは、この場合、すべての単語がクロードによって選択されているためです。
古いクロードモデルはどうですか？
EU 法には、2026 年 8 月 2 日より前に発売される Anthropic モデルに対する移行期間が含まれており、私たちはこれらのモデルにもウォーターマークを追加するよう取り組んでいます。これは今後数か月以内に展開される予定です。
これは Pangram のような AI 検出ソフトウェアとどう違うのですか?
AI 検出ソフトウェアは、それを提供する企業が私たちのキーを持っていないため、別の方法を使用します。とりわけ、これらのサービスは、テキストの微妙な（そしてそれほど微妙ではない）「内容」などの側面に注目します。

[切り捨てられた]

## Original Extract

Future Claude models will generate text that contains a watermark. This is a way of determining the likelihood that Claude was involved in writing the text, and we, along with several other major AI providers, are implementing this change to comply with the EU AI Act. In this article, we share answe
[truncated]

How Claude's text watermarking works \ Anthropic Skip to main content Skip to footer Research
Announcements How Claude’s text watermark works
Future Claude models will generate text that contains a watermark. This is a way of determining the likelihood that Claude was involved in writing the text, and we, along with several other major AI providers, are implementing this change to comply with the EU AI Act.
In this article, we share answers to some of the questions we’ve received about how our chosen watermarking method works, whether it affects Claude’s outputs, and why we’re making this change. To summarize:
We use a method of watermarking that does not have any practical impact on the quality or content of Claude’s outputs;
The difference between watermarked and un-watermarked text will not be distinguishable to readers;
Nothing is added to the text and there are no hidden characters;
Watermarking doesn’t require extra tokens, and will not be more expensive;
Watermarking carries no identifying information and can’t be traced to a specific person, organization, or chat;
Watermarking won’t be specific to Claude. As of August 2, the EU requires AI providers serving its market to mark AI-generated content. Other major model developers have signed the same Code of Practice and will be implementing their own watermarks.
Large language models like Claude work by generating one word at a time. Each time the model decides on the next word, it chooses among a list of potential candidates, ultimately selecting the most sensible or likely based on the preceding text. Take the sentence “The weather today was cold and…”. The next word is very unlikely to be “sugary.” But it is quite likely to be “overcast” or “grey.” Under most circumstances, it doesn’t matter much to the reader which of these latter two words the model ultimately chooses—the meaning of the sentence is largely the same either way. In cases like this, the choice is settled by a random number.
Watermarking uses low-stakes choices like these—which occur many times over a piece of generated text—to leave a pattern in Claude’s responses. That pattern is undetectable to the reader, but is detectable to anyone who has a key that encodes it. When watermarking is used, choices are still made at random, but the source of the randomness is different. Instead of using an arbitrary random number generator to pick the next word, watermaking uses the key and a few words that come before to settle what word the model should pick. That is, the words that Claude picks are still random, but now, one can check the sequence of words and see if it’s consistent with the choices Claude would make if it was using the key. If it is, one can assign a probability that the text was generated by Claude.
Importantly, it isn’t that the model will now always be biased toward overcast or grey. Just as with non-watermarked text, overcast might be selected in one sentence, grey in the next, depending on the words that came before. And it’s not the case that the watermarking method pushes Claude to choose a word it wouldn’t have considered anyway (for instance, it wouldn’t make Claude pick a word like “nubilous”—an obscure 1 synonym for overcast or grey that Claude almost certainly wouldn’t use under normal circumstances).
How does affect Claude’s outputs?
Watermarking does not impact the quality of Claude’s output. To a reader, a watermarked response is indistinguishable from an unwatermarked one (in this way, AI watermarks differ substantially from their namesakes on banknotes, other physical objects, and some digital documents, which are visible to the naked eye).
In internal testing, we’ve seen no impact of watermarking on the content, level of creativity, or readability of Claude’s text. In the SynthID-Text paper , which introduced the technique we use, Google DeepMind tested this impact by serving a model that used watermarking to a portion of their Gemini traffic and comparing thumbs-up and thumbs-down ratings. They found no statistically significant differences from the unwatermarked model. And in a controlled study, human raters comparing watermarked and unwatermarked answers side-by-side saw no difference in quality.
A useful analogy is to imagine you’re playing a game like Monopoly. On each turn, each player moves a random number of spaces around the board according to the roll of a die. Suppose that, instead of rolling the die to get this randomness, we decided to use a book of the digits of pi. 2 We start from a randomly-chosen digit (say, the 1,012,845th after the decimal place, which happens to be a 6), and from that point on each player simply uses the next digit in the sequence as their next “roll”.
For all intents and purposes, the moves are still random: it makes no difference to the players—or to the outcome of the game—whether the randomness comes from pi or from dice rolls each time. But if we could see the sequence of all the moves after the game (and we knew the value of pi), we could work out whether this was a game that likely used pi to determine its moves. The game that used pi is, in a sense, “watermarked”.
It’s the same for Claude-generated text. Watermarking doesn’t change the meaning or experience for the person reading it, but if you wanted to check after the fact whether the text was likely generated by Claude, the watermark allows you to do so.
Which specific method of watermarking do you use?
Claude’s text watermark is a version of the SynthID-Text approach published by Google DeepMind in a Nature paper in 2024. It belongs to a family of approaches that go back to a proposal by Scott Aaronson in 2022, all of which share the same design principle that we described above—the watermark only changes the source of the randomness used to pick among words.
There are limitations to the effectiveness of watermarking. Using our key, one can only answer the question “What is the likelihood this was partly written by Claude?” It doesn’t confirm whether the text was human-written, and it can’t tell whether the text was written by a different AI (even if that other AI uses watermarking, it would have a different key; it might also use a different watermarking method altogether). Detecting a watermarking also doesn’t work well on small samples, where there are fewer word choices and thus less information to go on. As a passage increases in length, confidence about Claude’s involvement increases too.
Watermarking is sparser on factual passages where there are fewer choices that can be made without decreasing the accuracy of the text. For example, take the sentence “Isaac Newton’s most famous work was called Principia …”. It really matters whether the next word is “ Mathematica ” (it’s the only right answer), so the watermark would have nothing to act on. The same is true for proofreading. If you hand Claude a piece of writing and ask it to edit only the grammar and punctuation and nothing else, the watermark can only live in the handful of corrections, which might be too few to register.
What about cases where Claude has proofread or edited human text?
The watermark only applies to words Claude chooses. When Claude proofreads text written by a person, what it gives back has generally only been lightly edited; because nearly all the words are the person’s, there’s very little (if anything) for the watermark to attach to. Depending on the length of the text and how heavily Claude has edited it, those changes might not be enough to make Claude’s involvement detectable. The more Claude writes, the more decisions it has to make, and the more space there is for a watermark.
As we noted above, AI watermarking takes advantage of decisions where either choice of a word would be equally good. Where an exact output is required—where there isn’t a choice, and something would be factually wrong or a piece of code would break if a different term was chosen—the watermark isn’t applied.
For example, once the model has written “2 + 2 =”, there is a very clear best choice for the next token (if the model is completing the sum, there isn’t an answer that’s equally as good as “4”; if it’s talking about George Orwell’s Nineteen Eighty-Four , there isn’t an answer that’s equally as good as “5”). The “nudge” of the watermark wouldn’t be applied here. For the same reason, code—which in very many cases has to be exact—has generally less watermarking than some other forms of text.
Having said that, in areas where there is an arbitrary choice between particular words or terms within the code, the watermark can be used, such as comments within code. But by definition, it will have a negligible effect on the actual code produced.
What does this mean for users?
Does this slow the model down, or make it more expensive?
No. Watermarking has a negligible impact on the speed of models, and because it produces no extra tokens, the model is the same price to serve and use.
Can a watermark be traced back to me or my organization?
No. The watermarking applies to Claude and its outputs. It doesn’t identify anything to do with individual users. There’s nothing in the watermark, or its key, that would allow anyone to recover any information about the user, their organization, or their chats with Claude.
Why are you watermarking Claude’s outputs?
We’re implementing watermarking to comply with the EU AI Act. Anthropic, along with several other major AI model providers and around 190 total signatories , signed the EU Code of Practice on Transparency of AI-Generated Content in July 2026. This requires AI system providers to use methods of “marking” AI-generated text. We’re applying watermarking globally at launch because we don't yet have a durable way to scope it by region. However, we will continue to evaluate different approaches, and will share updates when we have them.
How do I check if a piece of text was written by Claude?
We will soon be offering a watermark detection API. We’re in the process of working out the details of its implementation.
What about images and other files?
When Claude produces a file of a supported type (such as a .png, .jpg, or .svg), it will attach a content credential in the form of a small, cryptographically signed note in the file’s metadata, saying that the file was made or processed with Claude. This is an open industry standard called C2PA —the same used by camera manufacturers and in photo-editing software to record where an image came from. Any C2PA-aware tool can read it; we’ll be providing our own where you can drop a file and check.
This metadata label is very different from a watermark. Nothing in the file changes—it is not embedded or hidden. As with text, the credential only says Claude was involved in producing the file; it doesn’t include any identifying information.
Can’t someone just edit the text to get around the watermarking?
To some extent, yes. Light editing probably won’t remove the watermark completely; a complete rewrite where every word is replaced will. In the latter case, of course, it’s arguable whether the text can any longer be described as AI-generated.
What does a watermark actually prove?
A watermark can only determine that Claude was likely involved with the content at some point. It cannot distinguish “Claude wrote this” from “Claude heavily edited this.”
Do watermarks apply to translations?
Yes. A translation produced by Claude carries a watermark, because in this case every word is chosen by Claude.
What about older Claude models?
The EU law includes a transition period for Anthropic models launched before August 2, 2026, and we’re working to add watermarking for those models as well. This will be rolled out over the coming months.
How does this differ from AI detection software, like Pangram?
AI detection software uses a different method, because the companies that provide it don’t have our key. Among other things, those services look at aspects of the text like the subtle (and not-so-subtle) “tells” tha

[truncated]
