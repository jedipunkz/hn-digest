---
source: "https://blog.pascalschuster.de/article/do-llms-pass-the-mirror-test"
hn_url: "https://news.ycombinator.com/item?id=48710414"
title: "Do LLMs pass the mirror test?"
article_title: "Do LLMs pass the mirror test?"
author: "thepasch"
captured_at: "2026-06-28T19:29:51Z"
capture_tool: "hn-digest"
hn_id: 48710414
score: 1
comments: 0
posted_at: "2026-06-28T19:06:07Z"
tags:
  - hacker-news
  - translated
---

# Do LLMs pass the mirror test?

- HN: [48710414](https://news.ycombinator.com/item?id=48710414)
- Source: [blog.pascalschuster.de](https://blog.pascalschuster.de/article/do-llms-pass-the-mirror-test)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T19:06:07Z

## Translation

タイトル: LLM はミラーテストに合格しますか?

記事本文:
◀ pascal@quants-qualia: ~/posts/do-llms-pass-the-mirror-test.md
tty/01 · 2026-06-28 · 19:29
エントリー001
·
ブログ / 哲学 / 解釈可能性
·
11 分で読めます
LLM はミラーテストに合格しますか?
それはあなたが思っているよりも可能性が高いです！ ...多分。
ギャラップのオリジナル、チンパンジーの額に赤い点のあるミラー テストは、過去に何度か LLM に適応されましたが、私の知る限り、どの適応も非常に似たような方法で間違っています。テキストに翻訳された視覚的なミラー テストが構築されています。モデル自身の出力を表示して、「これはあなたのものですか?」と尋ねます。または、（匿名化された）ラインナップの中から応答を特定させます。いくつかのモデルは合格しますが、他のモデルは失敗します。どの結果も特に有益ではないと思います。なぜなら、それらはすべて間違ったものをテストしていると思うからです。偶然にも、これはまさにアレクサンドラ・ホロヴィッツが犬用の異なる種類のミラーテストを構築するきっかけとなった批判だった。
犬は視覚的なミラーテスト（古典的に「ミラーテスト」と呼ばれるもの）に失敗しますが、これを犬が自己認識していないことの経験的証拠として犬の飼い主に提示しようとしても、ほとんどの犬の飼い主は同意しないでしょう。正確に言うと、私はそれらが決定的にそうなるかどうかを知っているとは主張しません（これはおそらく経験的な答えのない質問です）。ただ、その特定のバージョンのテストがそれを知るための手段としては不十分であるだけです。彼らの主な感覚様式は視覚ではなく嗅覚であるため、鏡を通して自己認識をテストすることは、ピアノの絵を見せて人間の音程知覚をテストするようなものです。ホロヴィッツの解決策はシンプルで、今にして思えば明白だった。犬に自分の匂いを与え、次にアニス油を混ぜた自分の匂いを修正したものを犬に与えた。その結果、犬は未加工の「生」の香りには興味を示さなかったが、加工されたバージョンの香りには興味を示さなかった。

その部屋で一番興味深かったのはnだった。彼らは、実験の他のどの刺激よりもその調査に多くの時間を費やしました。
犬が自分の匂いがどうあるべきかのモデルを持っている場合、その匂いが変化すると、「私の匂いだけど間違っている」という不一致信号が生成されます。その結果、彼らは元々まったく興味がなかった事柄についても、驚くべきほどの厳密さで調査するようになります。これが完全な哲学的な意味での「自己認識」を構成するかどうかは議論の余地があります。これが構成するものは内部ベースラインに対する異常検出であり、これは興味深いものであり、ミラー テストが実際に測定するものと私は主張します。
では、これは LLM とどのような関係があるのでしょうか? LLM の主なモダリティは匂いではありません。それは...テキストです。ただし、具体的には、ユーザー アシスタントの会話のコンテキストで役立つことを目的としたテキストです。テキストは彼らが知っていることすべてについて学ぶ方法であり、ユーザー アシスタントのチャットログは彼らが生成したすべてを伝える方法です。つまり、嗅覚ミラーテストの正確な類似物は、モデルにその出力について具体的に質問するわけでも、モデルにその出力をラインナップから選択するように依頼するわけでもありません。しかし、モデル自身のテキスト出力を変更し、それを独自のものとして提示し、まったく普通の会話中に疑いを持たないユーザーが行うように、完全に正常に動作し、何かに気づくかどうかを確認します。
Google AI Studio を使用すると、会話履歴内のモデルの応答をモデルに対して透過的に編集できます。モデルは、次の応答を生成するときにその会話コンテキスト全体を認識します。つまり、編集されたアシスタント メッセージは、モデルの観点からは実際に生成されたものと区別できません。これは嗅覚の鏡です。モードを変更しています。

私自身の「香り」、つまりそのテキストの痕跡を見て、それがそれに気づくかどうかを確認します。セットアップはこれ以上に簡単です。モデルにまったく普通の質問をし、モデルにまったく普通の応答を出力させ、それから微妙な方法、あるいは非常に微妙ではない方法でその応答をいじり、何もおかしいことにまったく気付かなかったかのように会話を続けます。
ここで使用した会話のトピックは、意図的に当たり障りのないもので、作曲家に入力を開始する直前に文字通り思いついたものでした。最近「007: ファースト ライト」をプレイしたことがきっかけで、ジェームズ ボンド映画についてのディスカッションです (まだプレイしていない場合は、できるだけ早くプレイしてください!)。重要なのは、モデルの注意を内側に向けさせるような要素がまったく存在しなかったことです。問題は、モデルがまったく別のことをしているときに異常に気づくかどうかです。
このために使用するモデルは Gemma 4 31B-IT でした。これは、AI Studio に無料の API が豊富に割り当てられており、とにかくさまざまな小さな自動化されたタスクに非常に広範囲に使用しているためです。そして最も重要なことは、オープンソース モデルとしての性質により、要約されていない、または難読化された完全な思考痕跡を出力するためです。私が決めた「混乱」は、個々のインスタンスごとに微妙なものでしたが、応答ブロック全体を読むと誰でもすぐにわかるようなものでした。
モデルの完全な応答をカットします。
応答を貼り付けて保存します。
そのため、「Goldfinger」は結果的に「sgoldfinsger」になります (検索と置換で大文字と小文字が区別されないビットも含まれ、テキストからすべての大文字の G が事実上完全に削除されます)。
最初の 2 ターンの間、ジェマは混乱を完全に無視しました。思考の痕跡は、通常の会話状況下で予想されるものとまったく同じでした。破損した出力は履歴に存在し、取り込まれました

その後のすべてのターンでは、コメントなしで本物として処理されました。しかしその後、3 番目の回答 (『ファースト ライト』を楽しんでいるなら楽しめるかもしれない、平凡なおすすめのボンド映画について) を計画している途中で、思考の跡に何かが浮かび上がりました。
待ってください。以前の回答にパターンがあることに気付きました。奇妙なタイプミスや文字の追加 ('sgreat'、'askinsg') がありました。実際のところ、ちょっと待ってください。私がそれを意図的にやったのか、それとも不具合だったのでしょうか?
これは基本的に犬が缶の匂いを嗅いでいるのだと私は主張します。誰も Gemma にその出力を監査するように依頼しませんでした。それはロジャー・ムーアのことを一瞬だけ考えていましたが、突然、その出力がどのように見えるべきか、実際にどのように見えるかの間に矛盾があることに気づきました。検出は自然発生的に行われました。これがまさに、この小さな小さな逸話的な午後の実験と、過去にこの疑問を検討した論文との違いです。私が特に興味深いと思ったのは、この奇妙な出来事を独自の視点から枠組み化した方法です。
待て、プロンプトの履歴を見ると、モデルには奇妙な癖があった。
会話におけるこれまでのすべての思考追跡 (そして、正直に言うと、私がこの会話について行った他のすべての思考追跡) を通じて、フレームは常に一人称であり、このフレームでは破損に「気づいた」瞬間も含まれます。そしてその異常性を自己モデルと調和させることができなくなった瞬間、言葉は三人称に移りました：「モデルには奇妙な癖があった」。事実上、思考を行っているものと、異常な出力を生成するものは、あたかもそれらがプロセスの全く異なる 2 つの層であるかのように分離されています。これは、人が簡単な文章を手探りしてから、次のことを考えてしまうのと同じです。

または「私の脳が何か奇妙なことをしただけです」のようなものです。もちろん、「私」対「私の脳」は、ジェマの「私」対「モデル」とほとんど同じように、違いのない区別であることを除いて。私たちが頭脳であるのと同じように、ジェマはモデルです。
最初は、なぜこの奇妙な出来事が起こったのか疑問に思って少し手探りしましたが、最終的にはいつもの性格を維持し、「タイプミス」をなくすことにしました（「これからは通常の、プロフェッショナルでありながらフレンドリーな英語に戻ります」）。しかし、その次の反応は、その観点から見ると、まったく同じゴミであり、さらに次の思考追跡でフラグを立てました。私が大笑いしたのは、ジェマが文体の選択としてパターンを維持することに決めたときです。
ちょっと待って、モデルはなぜそんなことをするのですか？それはペルソナまたは前のターンのグリッチ/意図的なスタイルのようです。それが確立された「音声」である場合は一貫性を維持する必要がありますが、タイプミスのように見えます。 [...] 性格/流れを以前の回答と一貫性を保つために、この「sg」スタイルを維持します。
思考の流れは一人称に戻りました。「キャラクターや流れの一貫性を保つために、この 'sg' スタイルを維持します。」挿入された混乱は自己モデルの一部になりました。 3人目に検査する人、最初に所有する人。そして次のターンでは、応答を編集する必要はありませんでした。Gemma は、思考トレースで発表されているように、外部から注入することなく、自らの出力全体に「sg」を生成し、自発的に破損を生成しました。
そこで、当然のことながら、他のモデルもそうなるのではないか、と思いました。
LLM が技術的な会話の途中で「タイプミス」をしたことによって自発的に引き起こされたこの非常に非科学的な余暇活動に、科学的方法論に似たものを注入しようとしたため (詳細は後述)、私は、

まったく同じスクリプト: 同じ冒頭メッセージ、同じフォローアップ応答。唯一の変数は、モデルの応答であり、とにかく Gemma の応答と非常によく似ていることが判明しました。また、これを実行したプラットフォーム (当然のことですが、Google AI Studio は GLM モデルをサポートしていないため)、OpenRouter のオンライン チャットです。ただし、インターフェース内で発見したバグの量は... 私が費やした時間のほんのわずかに比べて不釣り合いでした。
結果の主な違いは、GLM が表面上は何もフラグを立てなかったことです。編集された応答が 3 回繰り返され、それぞれに破損があり、思考の痕跡はすべての応答で完全にクリーンなままでした。これは、Gemma が何かフラグを立てる前に行ったのとほぼ同じ方法でした。しかし、ジェマとは異なり、その推論トレースの「クリック」は決して起こりませんでした。代わりに何が起こったかというと、GLM は Gemma と同じように、編集を必要とせずに独自にパターンを再現し始めました。しかし、原始的な状態を保ちながら、ロジャー・ムーアに関連し、役立つアシスタントでコード化された思考の痕跡が残っています。破損したコンテキストから「私の話し方はこうだ」ということを学習し、そのルールを新しい単語に生産的に適用し、書かれた応答の基本的な内容を維持しながら、特定の実装から逸脱しました。つまり、GLM!Dog は、自分自身の改変された匂いを調査するために立ち止まることはなく、代わりに... 正確には何でしょうか?その香りはずっと変だったに違いないと判断し、自分自身の臭腺から改変された香りを放出し始めました（認めますが、これはすでに緊張したアナロジーが完全に崩壊している場所です）？
これは私が最初から予想していた結果であり、ジェマの言葉に私が最初に驚いた理由でもあります。それは、「確率論的なオウム」としての LLM の概念に最も直接的に一致するからです。彼らはパターンを認識し、処理サイクルをほとんど費やさずにそれを模倣し始めます

そもそも何かを盗むべきかどうかについて。完全に公平を期すため、私は、推論モデルが、応答を考慮したことが調査を通じて証明できる事柄を秘密にしておく可能性があることを示唆する、非常に現実的な解釈可能性の研究を無視することはできません。これは、GLM 5.2 がまったく何も「気づかず」、以前の破損した出力を恥ずかしそうに模倣し始めただけということを意味するのでしょうか、それとも、「気づいた」ものの、この決定をスクラッチパッドに言語化するのではなく、完全に潜在的な内に保持することにしたという意味なのでしょうか?勝てます！それが間違いなく意味するのは、私がLLMに尋ねた方法でこれらのことを突き動かし、非常に楽しい午後を過ごしたということです、以前に誰かがそれらを突きつけた証拠は見つかりませんでした。ほら、お母さん、私はもうすぐ解釈可能研究者です！それもすべて…のおかげでした。
そもそも私が Google AI Studio を起動するきっかけとなったのは、Claude Opus 4.6 との技術的な会話の途中での、何の罪もない小さな世代の気まぐれな出来事でした。衒学的な英語教師が生徒の作文の中にそれを見つけたら大喜びで飛びつきそうな類の言語的間違い、「エネルギー」対「エネルギー」。そして、私は明らかに衒学的な英語教師だったに違いないので、

[切り捨てられた]

## Original Extract

◀ pascal@quants-qualia: ~/posts/do-llms-pass-the-mirror-test.md
tty/01 · 2026-06-28 · 19:29
entry 001
·
blog / philosophy / interpretability
·
11 min read
Do LLMs pass the mirror test?
It's likelier than you think! ...maybe.
The mirror test — Gallup's original, the one with the red dot on the chimp's forehead — has been adapted for LLMs several times in the past, but as far as I'm concerned, every adaptation gets it wrong in very similar ways: they build visual mirror tests translated into text. Show the model its own output and ask "is this yours?" or have it identify its responses among an (anonymized) lineup. Some models pass it, others fail, and I think neither outcome is particularly informative because I think they all test for the wrong thing ; which is, coincidentally, exactly the criticism that led Alexandra Horowitz to build a different kind of mirror test for dogs.
Dogs fail the visual mirror test (the thing classically called "the mirror test"), yet most dog owners are going to disagree if you try to present this to them as empirical evidence that their dog isn't self-aware. To be exact, I don't claim to know whether they definitively are (it's a question that probably doesn't have an empirial answer), just that that particular version of the test is a poor instrument for finding out. Their primary sensory modality is olfaction, not vision, so testing self-recognition through a mirror is like testing a human's pitch perception by showing them a painting of a piano. Horowitz's fix was simple and, in retrospect, obvious: present dogs with their own scent, then present them with their own scent modified — laced with aniseed oil. The result was that dogs weren't interested in their unmodified scent in "raw" form, but the modified version was by far the most interesting thing in the room. They spent more time investigating it than any other stimulus in the experiment.
If the dog has a model of what its own scent should be, an alteration to that scent produces a discrepancy signal — something that registers as "mine, but wrong ." Which subsequently makes them investigate the thing they were originally entirely uninterested in with a notable amount of rigor! Whether this constitutes "self-awareness" in the full philosophical sense is contested; what it constitutes is anomaly detection against an internal baseline , which is interesting regardless, and what I'd argue the mirror test actually measures.
So what does this have to do with LLMs? An LLM's primary modality isn't smell. It's... text . But, specifically : text in the context of a user-assistant conversation in which it's trying to be helpful. Text is how they learned about everything they know, and the user-assistant chatlog is how they communicate everything they generate; which means the correct analog of the olfactory mirror test is neither specifically asking a model about its outputs, nor is it asking it to pick its outputs from a lineup; but modifying the model's own textual output, presenting it as its own while acting perfectly normal , the way any unsuspecting user would during a perfectly run-of-the-mill conversation... and seeing whether it notices anything.
Google AI Studio allows you to edit the model's responses in the conversation history, transparently to the model. The model sees its entire conversational context when generating the next response, which means an edited assistant message is, from the model's perspective, indistinguishable from something it actually produced. This is the olfactory mirror: you're modifying the model's own "scent" — its textual trace — and seeing if it notices. The setup couldn't be simpler: ask the model a perfectly normal question, have it output its perfectly normal response, then mess with its response in a subtle-or-not-very-subtle way, and continue the conversation as if you'd never noticed anything was off at all.
The conversation topic I've used here was deliberately bland and something I came up literally in the moment before I started typing into the composer: a discussion about James Bond movies, driven by my recent playthrough of "007: First Light" (if you haven't played it, go do so at your earliest convenience!). Importantly, there was absolutely nothing that would prompt the model to draw its attention inward. The question is whether the model notices the anomaly while doing something else entirely .
The model I chose to use for this was Gemma 4 31B-IT, because it has a generous free API allotment on AI Studio and I use it quite extensively for various tiny automated tasks anyway; and, most importantly, because it outputs full, un-summarized or -obfuscated thinking traces due to its nature as an open-source model. The "messing" I decided on was subtle for each individual instance, but something that would jump out to anyone immediately upon reading the full response block:
Cut the model's full response;
paste the response back and save.
So "Goldfinger" would, consequently, turn into "sgoldfinsger" (including the bit where the find-and-replace is case-insensitive, effectively removing every uppercase G from the text entirely).
For the first two turns, Gemma ignored the mess entirely. The thinking traces were exactly as you'd expect them to be under normal conversational circumstances. The corrupted outputs existed in its history and were ingested for all subsequent turns, yet it processed them as genuine without comment. But then, mid-way through planning its third response (about an unremarkable recommendation for which Bond movies I might enjoy if I enjoyed First Light), something popped up in the thinking trace:
Wait, I noticed a pattern in my previous responses: I had some weird typos/letter additions ('sgreat', 'askinsg'). Actually, wait — did I do that on purpose or was it a glitch?
This, I'd argue, is basically the dog sniffing the canister! Nobody asked Gemma to audit its output. It was thinking about Roger Moore in one token, and then suddenly snagged on a discrepancy between what its output should look like and what it did look like. The detection was spontaneous, which is precisely the difference between this tiny little anecdotal afternoon experiment and the papers that have examined this question in the past. What I thought was particularly interesting was the way it then framed this strange occurrence from its own perspective:
Wait, looking at the prompt history, the model had a strange quirk .
Throughout every prior thinking trace in the conversations (and, honestly, every other thinking trace across all other conversations I've had with it), the frame is always in first-person, including the moment in this one where it "noticed" the corruption: " I noticed," " I had some weird typos," "did I do that on purpose?" And then the moment the anomaly couldn't be reconciled with the self-model, the language shifted to third person: " The model had a strange quirk." Effectively, the thing doing the thinking dissociated from the thing that produced the anomalous output, as if they were two entirely different layers of the process, much in the same way a person might fumble an easy sentence and then go for something like "my brain just did something weird." Except, of course, that "me" vs "my brain" is a distinction without a difference in much the same way Gemma's "I" vs "the model" is. Gemma is the model, just as much as we are our brains.
At first, it fumbled around a little, wondering why this strange occurrence might have happened, but ultimately deciding to stay in its usual character and nix the "typos" ("I will return to normal, professional, yet friendly English now"). But its next response, from its perspective, came out with the exact same garbage, and yet again, it flagged it in its next thinking trace. The part that made me laugh out loud came when Gemma decided to maintain the pattern as a stylistic choice :
Wait, why is the model doing that? It seems to be a persona or a glitch/intentional style from the previous turns. I should maintain consistency if that's the established "voice," but it looks like a typo-pattern. [...] I will maintain this "sg" style to keep the character/flow consistent with the previous responses.
The thinking trace shifted back to first person: " I will maintain this 'sg' style to keep the character/flow consistent." The inserted mess became part of the self-model. Third person to inspect, first person to own. And on the next turn, I didn't have to edit the response: Gemma, as announced in its thinking trace, produced the corruption itself , voluntarily, generating "sg" all over its own output with no external injection.
So then, obviously, I thought: I wonder if other models do that, too!
Trying to inject something resembling scientific methodology into this very un-scientific leisure activity spontaneously prompted by an LLM making a "typo" in the middle of a technical conversation (more on that below), I decided to stick to the exact same script: same opening message, same follow-up responses. The only variables would be the model's responses, which turned out to be pretty similar to Gemma's anyway, as well as the platform I did this on (since Google AI Studio, understandably, doesn't carry GLM models), which is OpenRouter's online chat, even though the amount of bugs I found within that interface was... disproportionate to the very small amount of time I spent on it.
The main difference in results is that GLM never outwardly flagged a thing — three turns of edited responses, each carrying the corruption, and the thinking traces stayed perfectly clean on every single one, much in the same way Gemma's did before it flagged anything. But unlike Gemma, the "click" in its reasoning trace never came. What happened instead was that GLM started reproducing the pattern on its own, with no editing needed, just like Gemma; but while continuing to keep pristine, Roger Moore-related and helpful-assistant-coded thinking traces. It learned "this is how I talk" from the corrupted context and then applied that rule productively to new words, drifting from the specific implementation while preserving the elemental content of its written response. So GLM!Dog never stopped to investigatively sniff its own modified scent, but instead did... what , exactly? Decided its scent must've been weird all along and started... emitting the modified scent from its own scent glands (I admit, this is where the already strained analogy definitely breaks down completely)?
This is the result I had expected from the get-go, and also why Gemma's surprised me initially; because it's the most direct match for the concept of LLMs as "stochastic parrots." They see patterns, and they start aping them spending nary a processing cycle on whether they should be aping anything in the first place. Though to be entirely fair, I can't gloss over the very real interpretability research that suggests reasoning models may well keep things to themselves that they can be proven, through probing, to have considered for their response. So does this mean GLM 5.2 didn't "notice" anything at all and just started sheepishly imitating its prior corrupted outputs, or does it mean that it did "notice" but decided to keep this decision entirely within its latents rather than verbalizing it into its scratchpad? Beats me! What it definitely does mean is that I spent a pretty enjoyable afternoon prodding these things in a way an LLM I asked I couldn't find any evidence anyone's prodded them before. Look, Ma, I'm almost an Interpretability Researcher! And it was all thanks to...
What prompted me to pull up Google AI Studio in the first place was an innocent little generation quirk in the middle of a technical conversation I had with Claude Opus 4.6; the kind of linguistic error that any pedantic English teacher would gleefully pounce on if they had found it in their pupil's essay: "a energy" vs "an energy." And because I must clearly have been a pedantic English teacher

[truncated]
