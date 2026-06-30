---
source: "https://techstackups.com/articles/useful-llm-prompts-for-editing-your-own-technical-writing/"
hn_url: "https://news.ycombinator.com/item?id=48736393"
title: "Useful LLM Prompts for Editing Your Own Technical Writing"
article_title: "Useful LLM Prompts for Editing Your Own Technical Writing | Tech Stackups"
author: "ritzaco"
captured_at: "2026-06-30T18:35:45Z"
capture_tool: "hn-digest"
hn_id: 48736393
score: 1
comments: 0
posted_at: "2026-06-30T17:47:25Z"
tags:
  - hacker-news
  - translated
---

# Useful LLM Prompts for Editing Your Own Technical Writing

- HN: [48736393](https://news.ycombinator.com/item?id=48736393)
- Source: [techstackups.com](https://techstackups.com/articles/useful-llm-prompts-for-editing-your-own-technical-writing/)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T17:47:25Z

## Translation

タイトル: 独自のテクニカル ライティングを編集するための便利な LLM プロンプト
記事のタイトル: 独自のテクニカル ライティングを編集するための便利な LLM プロンプト |技術の積み重ね
説明: テクニカル ライターが自分の作品を自己編集するために使用できる LLM プロンプトの実用的なセットです。編集者がドラフトを見る前に、構造的、論理的、および行レベルの問題を発見します。

記事本文:
メイン コンテンツにスキップ 技術スタック ホーム トピック 比較 ガイド 記事 ニュース AX 独自のテクニカル ライティングを編集するための便利な LLM プロンプト
テクニカル ライティングの仕事は、ここ数年、特に生成 AI が一般大衆に導入されて以来、多くの進化を遂げてきました。技術文書の性質上、専門家と LLM の両方が読むコードと専門用語の両方が含まれているため、その文書の編集は、より一般的な文書の場合よりもはるかに複雑なプロセスになる可能性があります。
このため、ライターが技術編集者に送信する前に作業をスムーズにするために手元にクロード コードがあると便利です (いずれにしても、特に LLM を使用して最初の草稿を作成している場合)。
ここでは、機能と最大の効果を得るために実行する必要がある順序によって分類された一連のプロンプトを示します。私はこのプロセスの要約版を「テクニカル ライターが自分の作品を評価するためにクロード オーパス 4.8 を使用できますか?」で実行しました。結果を知りたい場合は。
すべてのプロンプトを実行する必要はないことに注意してください。むしろ、最も重要なもの (3 つのアスタリスク***) から、スキップ可能だが、おそらくとにかく実行する必要がある (1 つのアスタリスク *) まで、それらの重要性をランク付けしました。これを厳密なマニュアルではなく、参照できるライブラリとして扱います。
ステージ 1: 最初の読み取り / トリアージ
最初にこれらを下書き全体に対して実行して、編集時間をどこに配置するかを決定します。
この記事の最悪の点を 8 語で何ですか?***
この記事の全体像で最も悪い点は何ですか?***
読者はどこで読むのをやめてタブを閉じるでしょうか?行を引用します。**
これはスロープですか？イエスかノーか、証拠とともに。 (ステップ 5 で再実行)**
それはまったく機能しますか？主張、構造、完全性?
単一セッションで各サブパスを実行します

順番に; 「その他すべてを無視する」フレーミングにより、クロードは一度に 1 つのスコープに留まります。
論理的なエラーは無視してください。ただし、構造的な問題がある場合はこの記事で特定してください。***
見出しとトピックセンテンスのみからこの記事の概要を説明して、進行が成り立つかどうかを確認します。***
イントロは体が伝えられない何かを約束しているのでしょうか? それともその逆ですか?***
トランジションが必要な場所を提案してください。突然のジャンプを引用します。**
各セクションの最後に、私が答えていない質問が読者から寄せられるでしょうか?**
読者が必要とする、どのステップ、前提条件、または仮定をスキップしましたか?**
最も損失が少なく削除できる単一セクションはどれですか?**
これを 30% 削減しなければならない場合、何を優先しますか?**
体が十分に機能していないことを約束するタイトルは何ですか?**
構造上のエラーは無視してください。ただし、論理的なエラーがある場合はこの記事で特定してください。***
懐疑的な上級エンジニアとして行動してください。私が証拠や例で裏付けていないすべての主張をリストします。***
ドメインの専門家が最も強く反対するものは何でしょうか?最も強い反対をしてください。**
「なぜ」を主張しても実際には説明しないのはどこですか?**
認めるべきなのに認めていない明らかな反論はありますか?**
すべてのコード ブロックとコマンドを確認します。書かれたとおりに実行されないもの、または文章と一致しないものにはフラグを立ててください。***
古い可能性があるクレーム、バージョン固有のクレーム、または「as of」修飾子が必要なクレームにフラグを立てます。**
間違っているほど単純化しすぎた記述はありますか?**
私の用語は一貫して正しく使用されていますか?**
私の例は実際に要点を説明しているのでしょうか、それとも装飾的なものでしょうか?**
私は知っているふりをしながら実際に実証していないことは何ですか?**
変数名、ファイル パス、出力はすべてのコード サンプルで一貫していますか?*
記事内のスペルと文法の間違いを特定します。 ***
ステージ 3: 定性的な質問
どのような結末

読者はあなたの文章から得られる結果は何ですか？
3a.それは読者に何かを教えてくれますか？ ​
この記事には明確なストーリーの全体像がありますか?***
読者は、この記事で実行された実験やタスクを再現できますか?***
この記事には最終的なスタンス (または最終的なスタンスがないことの説明) が含まれていますか?**
→ これらの結論に基づいて、この記事は読者に何かを教えてくれますか?***
この記事は、読者がおそらく間違っている (または正しい) と何を教えていますか?***
読者の旅を最初から最後まで 100 語で要約すると、どのようなものになるでしょうか?**
著者は主張の連続性を維持していますか?つまり、読者は著者が自分の考えを知っていると信頼できますか?**
この記事の説得力を 1 (弱い) ～ 5 (非常に説得力がある) で評価します。**
→ これらの結論に基づいて、この記事は読者に影響を与えますか?記憶力、説得から行動、行動の変化を評価し、必要な修正を提案します。***
3c.読者を惹きつけるだけの説得力があるだろうか？ ​
この記事はどれほど説得力がありますか?どこに注意を示しましたか?**
最も退屈な段落は何ですか?**
読者が一週間後に覚えていることは何でしょうか?何もない場合は、そう言ってください。**
結末は得られるものなのでしょうか、それともただ終わるのでしょうか?**
→ これらの回答に基づいて、この記事は最初から最後まで読者の注意を引き付けることができますか?また、読者を失う危険性が最も高いのはどこですか?**
3D。聴衆にとって適切なピッチになっていますか？ ​
私の視聴者は[説明]です。基本的すぎるものや、彼らが持っていない知識を前提としているものにはフラグを立ててください。***
未定義の頭字語や内部関係者への言及をすべて見つけて、視聴者がそれを知っているかどうかを教えてください。**
どこでその口調が見下したり、頭上を越えたりするのでしょうか?**
ドメインの専門家が目を丸くするのは何でしょうか?**
読者が最もそう思う可能性が高い主張は何ですか

同意しませんか?**
読者について私が間違っている可能性がある最大の仮定は何ですか?**
英語を母国語としない読者がこれに遭遇した場合、どの文やイディオムにつまずくでしょうか?*
競合他社がこの記事のテイクダウンを書いた場合、冒頭の一文は何ですか?*
→ これらの回答に基づくと、この記事は [読者] にとって正しく提案されていますか? 基本的すぎるのか、高度すぎるのか、それとも正しいですか?**
3e.機械ではなく人間と読みますか？ ​
最も強い LLM フラグは何ですか?***
あらゆるトピックに関するあらゆる記事に登場する可能性のある、最も一般的な文は何ですか?**
フラグフィラーフレーズとより厳密な置換によるヘッジ。**
私がコミットを拒否する、最も過度にヘッジされているステートメントは何ですか?**
ここで最も怠惰な単語の選択は何ですか?*
声と口調をいくつかの言葉で説明してください – 一貫性は保たれていますか?*
たくさんのことを言っているが、何も意味していない文章はありますか?*
→ これらの回答に基づくと、この記事は人間が書いたものと考えられますか?また、よりそうであると思われる上位 3 つの変更は何ですか?**
次に、文章の明瞭さと経済性です。
この記事で最も弱い文は何ですか?***
最も効果が少ない文はどれですか?**
読者に何かを見せるのではなく、どこで何かを伝えているのでしょうか?**
解析に複数回のパスが必要なすべての文とその理由にフラグを立てます。**
最も長く、最も負荷の高い文を 3 つ見つけて、分割を提案します。**
受動態、名詞化、埋もれた主語を書き換えて特定します。**
同じ概念を異なる言葉で 2 回説明しているのはどこでしょうか?**
一度に頭の中に留めておくには長すぎる段落にマークを付けて、休憩を提案します。*
ステージ 5: コピー編集と校正
機械的な一貫性。変更するテキストを再チェックしないように最後に実行します。
すべての一貫性の問題をリストします: 見出しの大文字化、コードの書式設定、リストの句読点、製品名のスペル、数値の書式設定、ストレート

t と中引用符、全角ダッシュ スペース、オックスフォード カンマ、ダブル スペース。***
一般的なエラーは無視します。英国と米国のスペルや非並列構造 (例: 一部の箇条書きの後にコロンがあり、他の箇条書きにはコロンがない) などの不一致のみを確認します。***
すべてのハイパーリンクと相互参照を確認します: リンク テキストが指す場所と一致していますか?**
すべての図、表、およびコード ブロックが散文内で参照され、同じスタイルでキャプションが付けられていることを確認してください。**
同音異義語の誤用にフラグを立ててください – それ/それ、影響/効果、あなた/あなた*。
以下は、AI が生成した文章にフラグを立てる可能性のある単語のリストです。
次の単語の出現にフラグを立てて、置換を提案します。***
新しいプロンプトはありません。洗練されたドラフトでステージ 1 からこれら 2 つを再実行します。最終的なテキストに満足するまで繰り返します。
この記事の最悪の点を 8 語で何ですか?***
これはスロープですか？イエスかノーか、証拠を添えて。**
(最後の質問も、LLM にとっては興味深い実存的な質問です。)
もちろん最終作業は自分で行う必要がありますが、この編集ガイドを使用すれば、作業を引き渡す前に最も厄介なエラーを自分で見つけることができるはずです。
AP Punnoose は、Ritza のソフトウェア エンジニア兼テクニカル ライターであり、TechStackups.com の寄稿者です。
ステージ 2: 目標パス 2a。構造的
ステージ 3: 定性的質問 3a.それは読者に何かを教えてくれますか？
3c.読者を惹きつけるだけの説得力があるだろうか？
3D。聴衆にとって適切なピッチになっていますか？
3e.機械ではなく人間と読みますか？

## Original Extract

A practical set of LLM prompts technical writers can use to self-edit their work — catching structural, logical, and line-level issues before an editor sees the draft.

Skip to main content Tech Stackups Home Topics Comparisons Guides Articles News AX Useful LLM Prompts for Editing Your Own Technical Writing
The work of technical writing has seen many evolutions over the last few years – particularly since the introduction of generative AI to the general public. Because of the nature of technical documentation, which contains both code and jargon that will be read by both specialists and LLMs, editing that writing can be a far more complex process than would be the case for more generalist writing.
For this reason, it can be useful for a writer to have Claude Code on hand to smooth out their work before they send it to a technical editor (particularly if they’re using the LLM to write their first draft, in any case).
Here are a set of prompts, categorised by function and order in which they should be run for maximum effect. I ran an abridged version of this process in “Can Claude Opus 4.8 Be Used by Technical Writers to Evaluate Their Own Work?” if you’re interested in seeing the results.
Note that you don’t have to run every single prompt; rather, I’ve ranked them in importance from most critical (3 asterisks***) to skippable-but-you-should-probably-do-it-anyway (1 asterisk*). Treat this as a library you can draw from, rather than a strict manual.
Stage 1: First Read / Triage ​
Run these first, on the whole draft, to decide where your editing time should go.
What's the worst thing about this article in 8 words?***
What's the worst big-picture thing about this article?***
Where would a reader stop reading and close the tab? Quote the line.**
Is this slop? Yes or no, with the evidence. (re-run in Step 5)**
Does the thing work at all? Argument, structure, completeness?
In a single session, run each sub-pass in order; the "ignore everything else" framing will keep Claude on one scope at a time.
Ignore logical errors, but identify any structural problems with this article.***
Outline this article from headings and topic sentences only, so I can see whether the progression holds.***
Does the intro promise something the body doesn't deliver – or vice versa?***
Suggest where I need transitions; quote the abrupt jumps.**
What questions will a reader have at the end of each section that I haven't answered?**
What step, prerequisite, or assumption did I skip that a reader needs?**
Which single section could be deleted with the least loss?**
If you had to cut this by 30%, what goes first?**
What's the title promising that the body underdelivers on?**
Ignore structural errors, but identify any logical errors in this article.***
Act as a skeptical senior engineer. List every claim I haven't backed with evidence or an example.***
What would a domain expert push back on hardest? Give me the strongest objections.**
Where do I assert a 'why' but never actually explain it?**
Is there an obvious counterargument I should acknowledge but didn't?**
Review every code block and command. Flag anything that wouldn't run as written or doesn't match the prose.***
Flag any claim that may be outdated, version-specific, or that needs an 'as of' qualifier.**
Are any statements oversimplified to the point of being wrong?**
Is my terminology used consistently and correctly throughout?**
Do my examples actually illustrate the point, or are they decorative?**
What am I pretending to know but never actually demonstrate?**
Are variable names, file paths, and outputs consistent across all code samples?*
Identify any spelling and grammar errors in the article. ***
Stage 3: Qualitative Questions ​
What end result does the reader gain from your writing?
3a. Does it teach the reader something? ​
Does this article have a clear narrative throughline?***
Can the reader replicate any experiments or tasks performed in this article?***
Does this article contain a final stance (or an explanation for the absence of one)?**
→ Based on these conclusions, does this article teach the reader something?***
What is this article teaching the reader that they are probably wrong (or right) about?***
What does the reader's journey look like from start to finish, summarised in 100 words?**
Does the author maintain continuity in the points they make i.e. can the reader trust the author to know their own mind?**
Grade the persuasiveness of this article from 1 (weak) to 5 (very persuasive).**
→ Based on these conclusions, does this article make an impact on the reader? Evaluate memorability, persuasion-to-action, and behavior change, and suggest any necessary fixes.***
3c. Is it compelling enough to hold a reader? ​
How compelling is this article? Where did your attention flag?**
What's the single most boring paragraph?**
What's the one thing a reader will remember a week later? If nothing, say so.**
Is the ending earned, or does it just stop?**
→ Based on these answers, will this article hold a reader's attention start to finish, and where is it most at risk of losing them?**
3d. Is it pitched right for the audience? ​
My audience is [describe]. Flag anything too basic, and anything that assumes knowledge they won't have.***
Find every undefined acronym or insider reference and tell me whether my audience would know it.**
Where does the tone slip into condescending or over-their-head?**
What would make a domain expert roll their eyes?**
What's the one claim a reader is most likely to disagree with?**
What's the biggest assumption I'm making about the reader that might be wrong?**
If a non-native English reader hit this, which sentences or idioms would trip them up?*
If a competitor wrote a takedown of this article, what's their opening line?*
→ Based on these answers, is this article pitched correctly for [audience] – too basic, too advanced, or right?**
3e. Does it read as human, not machine? ​
What's the strongest LLM-flag?***
What's the most generic sentence, one that could appear in any article on any topic?**
Flag filler phrases and hedging with tighter replacements.**
What's the most over-hedged statement, where I refuse to commit?**
What's the laziest word choice here?*
Describe the voice and tone in a few words – does it stay consistent?*
Are there any sentences that say a lot but mean nothing?*
→ Based on these answers, does this article read as human-written, and what are the top three changes that would make it more so?**
Now the sentences – clarity and economy.
What's the weakest sentence in this article?***
What sentence is doing the least work?**
Where am I telling the reader something instead of showing them?**
Flag every sentence that took more than one pass to parse, and why.**
Find the three longest, most overloaded sentences and propose splits.**
Identify passive voice, nominalizations, and buried subjects, with rewrites.**
Where am I explaining the same concept twice in different words?**
Mark any paragraph too long to hold in my head at once and suggest a break.*
Stage 5: Copyedit & Proofread ​
Mechanical consistency, run last so you're not re-checking text you'll still change.
List every consistency issue: heading capitalization, code formatting, list punctuation, product-name spelling, number formatting, straight vs curly quotes, em-dash spacing, Oxford commas, double spacing.***
Ignore general errors – look only at inconsistencies like British vs US spelling and non-parallel structures (e.g. colons after some bullets but not others).***
Check every hyperlink and cross-reference: does the link text match where it points?**
Check that every figure, table, and code block is referenced in the prose and captioned in the same style.**
Flag any misused homophones – its/it's, affect/effect, your/you're.*
Below are a list of words that could flag AI-generated writing.
Flag any instances of the following words, and suggest replacements.***
No new prompts – re-run these two from Stage 1 on the polished draft. Repeat until you’re happy with the final text.
What's the worst thing about this article in 8 words?***
Is this slop? Yes or no, with the evidence.**
(That last one’s also a fun existential question for an LLM.)
You’ll obviously need to go through the final work yourself, but with this editing guide, you should be able to pick out the gnarliest errors yourself before you have to pass your work along.
AP Punnoose is a software engineer and technical writer for Ritza and contributor to TechStackups.com
Stage 2: Objective Pass 2a. Structural
Stage 3: Qualitative Questions 3a. Does it teach the reader something?
3c. Is it compelling enough to hold a reader?
3d. Is it pitched right for the audience?
3e. Does it read as human, not machine?
