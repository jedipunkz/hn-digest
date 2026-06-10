---
source: "https://steel.dev/blog/claude-code-deep-research-autopsy"
hn_url: "https://news.ycombinator.com/item?id=48476316"
title: "An autopsy of Claude Code's deep research"
article_title: "An autopsy of Claude Code's deep research - Steel | Open-source Headless Browser API"
author: "nkko"
captured_at: "2026-06-10T16:22:54Z"
capture_tool: "hn-digest"
hn_id: 48476316
score: 3
comments: 0
posted_at: "2026-06-10T13:54:00Z"
tags:
  - hacker-news
  - translated
---

# An autopsy of Claude Code's deep research

- HN: [48476316](https://news.ycombinator.com/item?id=48476316)
- Source: [steel.dev](https://steel.dev/blog/claude-code-deep-research-autopsy)
- Score: 3
- Comments: 0
- Posted: 2026-06-10T13:54:00Z

## Translation

タイトル: クロード・コードの深い研究の解剖
記事のタイトル: クロード・コードの深い研究の解剖 - スチール |オープンソースのヘッドレス ブラウザ API
説明: Steel は、AI エージェント専用に構築されたオープンソースのブラウザ API です。

記事本文:
クロード・コードの深い研究の解剖
クロード・コードの深い研究の解剖
私は、Claude Code にバイナリから独自のディープリサーチ ワークフローを取り出してもらい、そのワークフローにそれ自体に関する疑問を投げかけました。結論は、広範囲を検索し、倍返しすることはありません。システムも似ています。ゲーム全体はその 2 番目のホップです。
クロード コードは、その頭脳が完全に閉じられた単一のコンパイル済みバイナリとして出荷されます。内部のどこかには、Web の調査を依頼したときに実行されるワークフローがあります。つまり、質問の範囲を絞り、検索を広げ、読み取り、投票し、書き込みます。
読み取り可能なファイルとしては出荷されないため、研究目的で、Claude Code に独自のバイナリ内からワークフローを再構築してもらいました。数分後、ツールは自ら手術を行い、臓器を私に手渡しました。
私には持論がありました。これらの「深い調査」エージェントは深いのではなく、範囲が広いのです。彼らはあなたの質問を受け取り、それをいくつかの並行検索にスプレーし、結果を積み上げ、ボックスの奥深くに単語を刻印します。そこで私は、抽出されたワークフローに対して、深層研究ハーネスが実際にどのように機能するのか、幅が広いのか深いのか、という 1 つの質問を提示し、独自の解剖を調査させました。
私がそれを観察したところ、その理論は裏付けられ、同じ実行でそれが破られました。
これは動的ワークフローであり、Web の調査を依頼したときにクロード コードが実行されるようなものです。 /deep-research を実行するたびに、自分のマシン上で実行されます。この記事のすべては、Claude Code v2.1.170 に同梱されているワークフローを説明しています。後のビルドでは変更される可能性があります。
ヘッダーのコメントには、「bughunter アーキテクチャから移植された」と書かれており、 git と grep を WebSearch と WebFetch に置き換えています。誰かがバグハンティング エージェントを構築し、同じスケルトンが null ポインタの逆参照を見つけるのと同じくらい簡単に世界に関する事実を見つけられることに気づきました。
参照コードはパターンがどのように広がるかです。人々はそれを読みます

形状を学習してから、その形状を出荷します。したがって、重要なのは、それが何をするかということよりも、それをコピーする人に何を教えるかということです。
ディープリサーチのワークフローの仕組み
5つのフェーズ。上から下へ。一度。
あるエージェントは質問を、広範、技術的、最近の、逆張り、実務家の 5 つの角度に分割します。
5 人のエージェントが、角度ごとに 1 人ずつ並行して実行され、各エージェントは他のエージェントに対してブラインドになります。
URL の重複排除、ソース数の上限は 15 です。各情報源から 2 ～ 5 の反証可能な主張が得られ、それぞれに直接の引用と情報源の品質グレードが付けられます。
それぞれの主張に対して3人の懐疑論者がおり、それぞれが反論するよう指示された。 3 件中 2 件が拒否され、請求は無効になります。
あるエージェントは生存者を統合し、信頼度によってランク付けし、何が殺害されたかをリストしたメモを含む報告書を作成します。
抽出ステップでは Web ページを信頼しません。チェック可能なステートメントとそれを裏付ける引用が必要です。検証は意図的に敵対的です。研究エージェントがどのように団結しているかを誰かに教えたい場合は、このファイルを渡すよりもはるかに悪いことをすることもできます。
見ずにはいられなかったもの
それから私はハーネスが各捜索者に手渡すプロンプトを読みました。
すべての検索者は、元の質問との関連性によって結果をランク付けするように指示されます。指示はそのままです:「検索クエリだけでなく、元の質問との関連性によってランク付けします。」それらはすべて、スコープ エージェントが 2 番目のゼロに書き込んだのと同じプロンプトから始まります。
検索者が何を見つけても、検索される内容が変わることはありません。エージェントは結果を読んで、それが何かを暗示するような緊張感を感じて、そこからより鋭い質問を形成することはありません。
オーケストレーターはループしません。スコープは 1 回だけ発生します。検索は 1 回だけ行われます。この報告書は、その単一の掃海が海底から引きずり出されたものすべてに基づいて作成されています。
それが広さと深さのギャップで、一枚の絵に収まります。
重要なことを実際にどのように研究するか考えてみましょう。検索すればわかる

d、そしてその読み取りによって次の質問が書き換えられます。ホップ 1 への答えはホップ 2 への入力になります。
系図学者が教区の記録で姓のスペルミスを発見し、そのスペルミスが次のクエリになります。記者は日付が合わないことに気づき、日付を追う。
事前に 5 つのクエリを作成して停止する人はいません。
リファレンス ハーネスはスレッドに従うことができません。それが不足している 2 番目のホップです。
次に、大手ホスト製品がより優れたブランディングの下で​​同じことを行うかどうかを確認しました。
そうではありません。私は彼ら自身のエンジニアリングに関する記事を読みましたが、ほとんどすべての本格的な記事はハイブリッドです。つまり、ラウンド内での並列ファンアウトと、ラウンド間の真の反復です。
Anthropic のマルチエージェント研究システムは、この分野で最も広範なシステムのように見えます。つまり、リード エージェントが 3 ～ 5 人のサブエージェントを逐次ではなく並行して起動します。しかし、次の文を読んでください。リードエージェントは「これらの結果を総合して、さらなる調査が必要かどうかを判断し、必要な場合は追加のサブエージェントを作成したり、戦略を洗練したりすることができる。」それがループです。最もファンのように見えるシステムでさえ、最上位で反復されています。
残りも同じように並べます。
OpenAI Deep Research は 1 つの推論モデルであり、o1 と同じレシピでトレーニングされ、ブラウジングし、「遭遇した情報に反応して必要に応じてピボット」します。 2 番目のホップはウェイトにベイクされます。
Gemini は明示的な「計画、検索、読み取り、反復、出力」ループを提供し、それ自体を「シングルパス モデルではなく、反復的なマルチステップ システム」と呼んでいます。
Perplexity は「次に何をすべきかを検討し、さらに学ぶにつれて研究計画を洗練させます。」
オープンソースも同じ流れで分割されます。 GPT Researcher はデフォルトで幅広いファンアウトを実行し、ツリーに対して別の再帰的な「ディープリサーチ」モードを提供します。 dzhng/deep-research は、幅と深さをノブとしてそこに配置し、各レベルの発見をフィードします。

次のレベルのクエリに進みます。 LangChain の open_deep_research は逆に、スーパーバイザーに「ほとんどのクエリに対して 1 つのサブエージェントから開始する」ように指示します。
したがって、「広くて深くない」というのは、この分野の説明としては間違っています。何が生き残るかはより正確です:
同梱のリファレンス ハーネスはシングルパス幅です。表面的に似ているシステムはハイブリッドです。ラウンド内では広く、ラウンド全体では奥行きがあります。違い全体は 2 番目のホップであり、2 番目のホップがコストがかかる部分です。
最後の条項は私が何度も思い出す条項です。
参照が浅いままである理由
なぜループではなく 1 つのパスなのでしょうか?理由は 3 つありますが、どれも怠惰ではありません。
深さは並列化されません。 2 番目のホップは最初のホップを待つ必要があるため、ファンアウトを魅力的なものにする実時間の勝利を失うことになります。
ウサギの穴のような深さ。何も停止を指示されないため、エージェントは反復的に放置され、十数のクエリを通じて接線を追跡します。
幅は読みやすいです。パイプライン全体を頭の中に収めることができ、まさにそれが人々が模倣する形状である理由です。
そして深さは高価です。 Anthropic の報告によると、エージェントはチャットのトークンの約 4 倍、マルチエージェント システムでは約 15 倍のトークンを消費します。彼らのマルチエージェント設定は、内部評価で単一エージェントを 90.2% 上回っていましたが、それが価値の高い作業にのみ効果があることは認めています。 (自己報告、内部的、幅優先。事実ではなく方向性として捉えてください。)
リファレンスは理解できるように構築されています。シングルパスで幅広の形状は最も読みやすい形状であり、それが伝播する形状である理由です。
敵対的検証の限界
このハーネスは、浅い調査を補うために検証に重点を置いています。3 人の懐疑論者がすべての主張に投票し、2 人の拒否者がそれを無効にします。これはパイプラインの最も洗練された部分です。
しかし、多数決は、3 人の懐疑論者が独立して失敗した場合にのみ機能します。自己一貫性により精度が向上します

まさにその仮定の下で、さまざまな推論パスをサンプルし、独立したエラーをキャンセルします。深層研究エージェントに関する最近の研究では、その逆のことが判明しました。それらのエラーは相関しています。「ある実行で発生したエラーは、他の実行でも再発する傾向があります。」 3 人の懐疑論者が 1 つの盲点を共有する場合、3 票は 1 票が 3 回請求されることになります。
したがって、2 つの弱点は実際には 1 つです。探索は浅く、間違いを見つけるために構築された網には探索と同じ穴があります。
私はそれを本当の質問に向けました
形状の変化を観察するために、ハーネスに純正のハーネスを与えました。LoRA は 7B ～ 70B 範囲のオープン モデルの完全な微調整に適合しますか?約 13 分間実行され、1 行の算術演算から読み取ることができる形式で 98 のエージェントが生成されました。つまり、1 つのスコープ、5 つの検索、16 のフェッチ、75 のベリファイア、1 つの合成です。
演奏はOpus 4.8でした。入出力の分割を記録したことはないため、請求額は概算ですが、定価 (100 万トークンあたり 5 ドルの入金、25 ドルのアウト) で、これらの 236 万トークンのコストは 10 ドルから 60 ドルの間になります。一つ質問です。
検証者は収入を得ました。彼らは、「新しい亜種はギャップをほぼ埋める」という主張、つまり懐疑論者が一つの矛盾する情報源を使って反論できる簡単で独立した失敗を潰した。ふるい分けの全体を 1 つのグラフで見ることができます。情報源が息を吹き込み、主張が吐き出し、ギロチンが残りをとります。
しかし、この報告書は自身のギャップも指摘している。70Bの証拠は乏しく、最も有力な主張は査読を受けていないブログに依拠しているというものだ。穴が見えたが、置く場所がなかった。
ディープハーネスを使用すると、その警告が次のクエリに変わったはずです。これは、次に作成するクエリがないため、直接合成に進みました。実際の質問では、欠落している 2 番目のホップがそれです。
次に、それを打ち破るために構築された質問、Perplexity の DRACO ベンチからの針でそれを実行しました。

hmark : カリフォルニアの古いグレート アメリカ パークで行われる 5K レースのタイトルに「風船ガム」を付けてください。答えは同音異義語の中に隠されています。フォレスト ガンプのエビ ブランドである Bubba Gump は、「Run Forrest Run」5K を運営しています。これはまさに、単一のディープ チェーンが見つけて自動的に抜け出す傾向がある、疑わしいほど巧妙な接続です。
ハーネスは構造的な理由でそうなりました。 5 回の並行検索のうち 4 回で独立して同音異義語が明らかになったので、孤独な推論者が幸運な予感を信じる必要はありませんでした。その後、敵対的な検証者は、エビのブランドが風船ガム「である」とシステムが過剰に主張するのを防ぎました。書かれた内容は慎重でした。前提には同音異義語の間違いがあります。ここに本当の人種があり、信頼度は中程度です。
その慎重な答えが、彼らが互いに戦うときに生み出す幅と懐疑性です。場合によっては、ファンがまさに正しい形状であることもあります。
何かを構築している人に言いたいこと
どのパターンから始めたかを把握してください。ほとんどの研究エージェントは、コピーされたリファレンス実装として開始されます。クロード コードの詳細な研究の形状を変更せずに出荷すると、シングルパス幅で出荷することになります。つまり、安価、高速、読みやすく、再確認が必要なものは何も見えなくなります。
トポロジを難易度に合わせます。角度が独立していて、答えが並行する事実の総合である場合には、より広い範囲に進みます。次の質問が最後の回答に本当に依存している場合にのみ、深さに対して料金を支払います。
2 番目のホップの予算。深く掘り下げると、シリアル レイテンシー、最大 15 倍のトークン請求、および独自の投票が相関していると想定する必要がある検証レイヤーが予想されます。
フロンティアの幅はこれ以上ありません。それは規律ある第二のホップです。いつそれを受け入れるべきか、いつやめるべきか、そして自分が見つけたものをどのように信頼するかを知っています。それをきれいに解決した人は誰もいません。それは私たちが目指して構築している部分です。
ハーネスは私にとってそれ自体が限界であることがわかりました。私は、それがまさにこの質問を調査し、私のプロンプトを五方向に煽り、決して顔色を変えないのを見ました。

それが何を見つけたかを確認してください。
そこで私はディープ バージョンを構築しました。つまり、検索する前にルーティングし、ループする前にブラウズし、自身の主張を検証するリサーチ エージェントです。良くなる前にさらに悪くなりました。私が最も誇りに思っていた検証ツールによってレポートの品質が低下し、その理由を解明するのに時間がかかりました。それが次の部分です: 耐久性のある研究者: 失敗しても生き残る研究エージェントを構築する 。
セカンドホップについてはまだ考え中です。しかし今、私はそれを持っています。
LLM をオンラインで利用するためのより良い方法。

## Original Extract

Steel is an open-source browser API purpose-built for AI agents.

An autopsy of Claude Code's deep research
An autopsy of Claude Code's deep research
I had Claude Code pry its own deep-research workflow out of its binary, then pointed that workflow at a question about itself. The verdict: it searches wide and never doubles back. The systems it resembles do. The whole game is that second hop.
Claude Code ships as a single compiled binary with its brain welded shut. Somewhere inside is the workflow it runs when you ask it to research the web: scope the question, fan out searches, read, vote, write.
It does not ship as a readable file, so, for research purposes, I had Claude Code reconstruct the workflow from inside its own binary. A few minutes later the tool had performed surgery on itself and handed me the organ.
I had a theory: these "deep research" agents are not deep, they are wide. They take your question, spray it across a handful of parallel searches, pile up the results, and stamp the word deep on the box. So I pointed the extracted workflow at one question, how do deep-research harnesses actually work, are they wide or deep , and let it research its own autopsy.
What I watched it do confirmed the theory and broke it in the same run.
This is a Dynamic workflow , the kind Claude Code runs when you ask it to research the web. It executes on your own machine every time you run /deep-research . Everything in this piece describes the workflow as it ships in Claude Code v2.1.170; later builds may change it.
The header comment says it was "ported from a bughunter architecture," swapping git and grep for WebSearch and WebFetch . Someone built a bug-hunting agent, then noticed the same skeleton finds facts about the world as easily as it finds null-pointer dereferences.
Reference code is how patterns spread. People read it to learn the shape, then they ship the shape. So what matters is less what it does than what it teaches everyone who copies it.
How the deep-research workflow works
Five phases. Top to bottom. Once.
One agent splits the question into 5 angles: broad, technical, recent, contrarian, practitioner.
Five agents run in parallel, one per angle, each blind to the others.
Dedup URLs, cap at 15 sources. Each source yields 2-5 falsifiable claims, each with a direct quote and a source-quality grade.
Three skeptics per claim, each told to refute it. Two rejections out of three and the claim dies.
One agent merges survivors, ranks by confidence, writes the report with a note listing what got killed.
The extraction step does not trust a webpage; it demands a checkable statement plus the quote that backs it. Verification is adversarial on purpose. If you wanted to teach someone how a research agent hangs together, you could do far worse than handing them this file.
The thing I couldn't stop looking at
Then I read the prompt the harness hands each searcher.
Every searcher is told to rank its results by relevance to the original question . The instruction is verbatim: "Rank by relevance to the ORIGINAL question, not just the search query." Every one of them starts from the same prompt the scoping agent wrote at second zero.
Nothing any searcher finds ever changes what gets searched. No agent reads a result, feels the tug of wait, that implies something , and forms a sharper question from it.
The orchestrator does not loop. Scope happens once. Search happens once. The report is built from whatever that single sweep dragged up off the seabed.
That is the gap between wide and deep, and it fits in one picture.
Think about how you actually research something that matters. You search, you read, and the reading rewrites the next question. The answer to hop one becomes the input to hop two.
A genealogist hits a misspelled surname in a parish record and that misspelling becomes the next query. A reporter notices the dates don't line up and chases the dates.
Nobody writes five queries in advance and stops.
The reference harness cannot follow a thread. That is the missing second hop.
Then I checked whether the big hosted products do the same thing under nicer branding.
They don't. I read their own engineering writeups, and almost every serious one is a hybrid: parallel fan-out inside a round, genuine iteration across rounds.
Anthropic's multi-agent research system looks like the most wide-open thing in the field: a lead agent spinning up three to five subagents in parallel rather than serially . But read the next sentence. The lead agent "synthesizes these results and decides whether more research is needed, and if so, it can create additional subagents or refine its strategy." That is a loop. Even the system that looks the most like a fan is iterating at the top.
The rest line up the same way:
OpenAI Deep Research is one reasoning model, trained with the same recipe as o1, browsing and "pivoting as needed in reaction to information it encounters." The second hop is baked into the weights.
Gemini ships an explicit "Plan, Search, Read, Iterate, Output" loop and calls itself "an iterative, multi-step system, not a single-pass model."
Perplexity "reasons about what to do next, refining its research plan as it learns more."
Open source splits down the same seam. GPT Researcher runs wide fan-out by default, then offers a separate recursive "Deep Research" mode for the tree. dzhng/deep-research puts breadth and depth right there as knobs and feeds each level's findings into the next level's queries. LangChain's open_deep_research goes the other way and tells its supervisor to "start with one sub-agent for most queries."
So "wide, not deep" is wrong as a description of the field. What survives is more precise:
The bundled reference harness is single-pass wide. The systems it superficially resembles are hybrid: wide within a round, deep across rounds. The whole difference is the second hop, and the second hop is the part that costs.
That last clause is the one I keep coming back to.
Why the reference stays shallow
Why one pass and not a loop? Three reasons, none of them laziness:
Depth doesn't parallelize. A second hop has to wait on the first, so you forfeit the wall-clock win that makes fan-out attractive.
Depth rabbit-holes. Left to iterate, an agent chases a tangent through a dozen queries because nothing tells it to stop.
Width is legible. You can hold the whole pipeline in your head, which is exactly why it is the shape people copy.
And depth is expensive. Anthropic reports that an agent burns roughly 4× the tokens of a chat, and a multi-agent system about 15× . Their multi-agent setup beat a single agent by 90.2% on an internal eval, but they concede it only pays off on high-value work. (Self-reported, internal, breadth-first. Take it as direction, not fact.)
A reference is built to be understood. Single-pass and wide is the most legible shape there is, which is why it is the one that propagates.
The limits of adversarial verification
The harness leans hard on verification to make up for shallow search: three skeptics vote on every claim, two rejections kill it. It is the most sophisticated part of the pipeline.
But majority voting only works if the three skeptics fail independently . Self-consistency earns its accuracy gains exactly under that assumption: sample diverse reasoning paths, let independent errors cancel. Recent work on deep-research agents finds the opposite. Their errors are correlated: "Errors arising in one run also tend to recur in other runs." When three skeptics share one blind spot, three votes is one vote billed three times.
So the two weaknesses are really one. The search is shallow, and the net built to catch its mistakes has the same holes the search does.
I pointed it at real questions
To watch the shape move, I gave the harness a genuine one: does LoRA match full fine-tuning for open models in the 7B-70B range? It ran for about thirteen minutes and spawned 98 agents in a shape you can read off a single line of arithmetic: one scope, five searches, sixteen fetches, seventy-five verifiers, one synthesis.
The run was on Opus 4.8. We never logged the input/output split, so the bill is an estimate, but at list prices ($5 per million tokens in, $25 out) those 2.36 million tokens cost somewhere between ten and sixty dollars. For one question.
The verifiers earned their keep. They killed the "new variant nearly closes the gap" claims, the easy, independent failures a skeptic can refute with one contradicting source. You can watch the whole winnowing in one chart: sources breathe in, claims breathe out, the guillotine takes the rest.
But the report also flagged its own gap: the 70B evidence is thin, and the strongest claims lean on a non-peer-reviewed blog. It saw the hole and had nowhere to put it.
A deep harness would have turned that caveat into the next query. This one moved straight to synthesis, because it has no next query to make. That is the missing second hop, on a real question.
Then I ran it on a question built to break it, a needle from Perplexity's DRACO benchmark : name the 5K race at California's old Great America park with "bubble gum" in its title. The answer hides in a homophone. Bubba Gump, the Forrest Gump shrimp brand, runs a "Run Forrest Run" 5K, exactly the suspiciously neat connection a single deep chain tends to find and then talk itself out of.
The harness got it, and for a structural reason. Four of the five parallel searches surfaced the homophone independently, so no lone reasoner had to trust a lucky hunch. The adversarial verifiers then kept the system from overclaiming that a shrimp brand "is" bubble gum. What it wrote was careful: the premise has a homophone error, here is the real race, medium confidence.
That careful answer is what width and skepticism produce when they fight each other. Sometimes the fan is exactly the right shape.
What I'd tell someone building one
Know which pattern you started from. Most research agents begin life as a copied reference implementation. Ship the Claude Code deep-research shape unchanged and you ship single-pass width: cheap, fast, legible, and blind to anything that needs a second look.
Match topology to difficulty. Go wide when the angles are independent and the answer is a synthesis of parallel facts. Pay for depth only when the next question genuinely depends on the last answer.
Budget for the second hop. When you go deep, expect serial latency, a ~15× token bill, and a verification layer that must assume its own votes are correlated.
The frontier is not more width. It is a disciplined second hop: knowing when to take it, when to quit, and how to trust what you find. Nobody has cleanly solved that. It is the part we are building toward.
The harness proved its own ceiling for me. I watched it research this exact question, fan my prompt five ways, and never look back at what it found.
So I built the deep version: a research agent that routes before it searches, browses before it loops, and verifies its own claims. It got worse before it got better. The verifier I was proudest of dropped my report quality, and it took me a while to work out why. That is the next piece: Durable Researcher: building a research agent that survives its own failures .
I'm still thinking about the second hop. But now I have one.
A better way to take your LLMs online.
