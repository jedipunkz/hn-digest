---
source: "https://stillig.net/posts/tamper-evident-llm-calls/"
hn_url: "https://news.ycombinator.com/item?id=49256961"
title: "No standard says what to record about an LLM call, so I built the record"
article_title: "A judge ran the same prompt three times and got three different answers - Johannes Stillig"
author: "johtidebreak"
captured_at: "2026-08-11T12:42:00Z"
capture_tool: "hn-digest"
hn_id: 49256961
score: 1
comments: 0
posted_at: "2026-08-11T12:07:44Z"
tags:
  - hacker-news
  - translated
---

# No standard says what to record about an LLM call, so I built the record

- HN: [49256961](https://news.ycombinator.com/item?id=49256961)
- Source: [stillig.net](https://stillig.net/posts/tamper-evident-llm-calls/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T12:07:44Z

## Translation

タイトル: LLM 通話について何を記録すべきかについては標準が定められていないため、記録を作成しました
記事のタイトル: 裁判官は同じプロンプトを 3 回実行し、3 つの異なる回答を得た - Johannes Stillig
説明: 裁判所は専門家の AI クエリを再実行し、3 つの異なる数値を取得しました。これは私が現在すべてのモデルコールについて保存している記録であり、それが正直に証明している小さなことです。

記事本文:
ヨハネス・スティリグ
執筆
について
に切り替えます
暗い
光
モード
裁判官は同じプロンプトを 3 回実行し、3 つの異なる回答を得た
私はヨハネスです。私は、すべての検出結果がその検出結果のソース バイトまで遡る必要がある調査エンジンを構築していますが、言語モデルが検出結果の生成を支援し始めると、これは厄介な約束であることが判明しました。
ニューヨークの信託会計訴訟で、専門家証人は法廷で、損害賠償額の計算を照合するためにAIアシスタントを使用したと証言した。判事は彼に何を入力したか尋ねた。彼は思い出せなかった。裁判官はどのような情報源を使用したのかを尋ねた。彼には言えなかった。
そこで裁判官は裁判所のコンピューターに自分で質問を入力し、949,070.97ドルを受け取った。彼は 2 台目の法廷コンピューターで再度実行し、948,209.63 ドルを獲得しました。 3 分の 1 は 95 万 1,000 ドル強を返しました。
裁判所発行の機械が 3 台、質問が 1 つ、回答が 3 つあります。代理裁判所は、変動はそれほど大きくなかったものの、変動があったという事実は出力の信頼性に疑問を投げかけ、そのような証拠が認められる前に弁護士にはAIの使用を開示する積極的な義務があると判示した。
それが 2024 年 10 月に決定されたウェーバーの問題です。専門家は何も捏造していませんでした。彼は自分が何を尋ねたのか、何が返ってきたのかを言えないことに陥った。
私は先月、自分の製品にそのようなことが起こらないように努めてきました。私が構築したものは機能しますが、当初主張したほどではないことが証明されており、これら 2 つの間のギャップは書き留める価値があります。
再現性を最初に求めるのは間違っています
この問題に気づくと、本能的に再現性を追求します。パラメータを固定し、シードを設定すると、同じ答えが 2 回得られます。
それはできませんが、その理由は浮動小数点よりも興味深いものです。
シンキングマシンL

ab はその最も明確な説明を発表しました。同時浮動小数点加算は非結合的であるという通常の説明は間違っていることが判明しました。同じデータに対して同じ行列の乗算を繰り返し実行すると、毎回ビット単位で同じ結果が得られます。実際の原因は、出力がたまたま参加していたバッチに依存しており、そのバッチがその時点でサーバーにアクセスしていた他のユーザーに依存していることです。彼らの測定では、温度ゼロで 1 つのプロンプトを 1,000 回完了すると、80 個の固有の完了が生成されました。最初の 102 個のトークンは毎回同一でした。その後、彼らは分岐しました。
あなたから見ると、API の他のユーザーは入力ではありません。それらは天気の性質です。
これは修正可能ですので、発送済みです。バッチ不変カーネルはビット同一の出力を提供し、vLLM と SGLang の両方が決定的モードを公開するようになりました。コストは実際に発生し、誰が測定するかに応じてスループットの 4 分の 1 から 3 分の 2 の間になります。また、決定論は伝染するというさらに厄介な特性が根底にあります。 10 個の通常のリクエストのバッチに 1 つの確定的なリクエストが加わると、サーバーの合計スループットが 1 秒あたり 931 トークンから 415 トークンに低下しました。これは、自分でこっそり購入できるリクエストごとのオプションではありません。
そして、ホストされた API ではどれも利用できません。 Anthropic Messages API にはシードがありません。現在のフロンティア Claude モデルでは、温度を設定することさえできません。デフォルト以外の値は 400 を返します。Bedrock の Converse inferenceConfig には 4 つのノブがあり、それらのどれもシードではありません。シードとバックエンド フィンガープリントを提供した唯一のプロバイダーである OpenAI は、OpenAPI 仕様で両方を非推奨: true としてマークし、Responses API にはどちらも同梱していません。
つまり、誰もが AI ガバナンスについて話す一方で、業界の 1 つの再現性アフォーダンスは撤回されつつあります。
それはそれでいいのですが、

なぜなら、再現性を最初に求めるのは間違っていたからです。ロギングは認証であり、複製ではありません。何を送ったのか、何が戻ってきたのかを証明します。モデルが再び同じことを言うという証明にはなりません。これらは異なるプロパティであり、現在利用できるのはそのうちの 1 つだけです。
ウェーバーの専門家は、彼の答えが再現不可能だったからといって、負けたわけではありません。彼は記録がなかったため負けた。
何を記録するかを誰も指定していない
私が最も驚いたのはここです。私は、複数のモデル呼び出しが見つかり、選択する必要があることを期待して、1 つのモデル呼び出しについて何をキャプチャするかを示す標準を探しました。
NIST AI リスク管理フレームワークには「トレーサビリティ」は 0 回も登場しません。 「記録」も同様です。生成 AI プロファイルには「出所」という表現が 67 回あり、そのすべてがコンテンツの出所やトレーニング データに関するものです。なぜなら、NIST がそこで解決している問題はディープフェイクであり、防御可能性ではないからです。すべての AI 部品表フォーマットは、呼び出しではなくモデルを記述します。CycloneDX の入力および出力フィールドは、値ではなくフォーマットを保持します。 Sigstore のモデル署名は、ファイル ダイジェストのリストに署名します。これにより、重みが改ざんされていないことがわかり、その内容については何もわかりません。
仕様に最も近いものは、予期せぬ方向から生まれます。 OpenTelemetry の GenAI 規則は、入力メッセージ、出力メッセージ、システム命令など、必要な属性を正確に定義します。次に、仕様では、計測器はデフォルトでそれらをキャプチャすべきではないとし、オプション 1 として、命令、入力、または出力を記録しないことを列挙しています。
これが私が言っていることすべてに対する最良の反論であり、反対側から読んでも私の主張のすべてでもあります。この規格は、その能力が取るに足らないものであることを証明しています。また、意図的なデフォルトが証拠の穴であることも文書化されています。
規制に関しては、EU AI 法が定めています。

リスクの高いシステムでは、その存続期間全体にわたって自動イベント ログをサポートする必要があり、プロバイダーと導入者はそれらのログを少なくとも 6 か月間保存する必要があります。それについては 2 つあります。同法がログの内容を指定している唯一の箇所は、遠隔生体認証を扱うサブセクションであり、検索により一致した入力データに到達するが、これは「プロンプトを記録する」とは程遠い。そしてこの義務は、AI に関するデジタル オムニバスによって 2027 年 12 月 2 日に延期されるまで、2026 年 8 月 2 日に適用される予定でした。この修正案は 2026 年 7 月 27 日に発効しました。私は 3 日前のこの投稿の下書きに 8 月の日付を記載していました。
フィールドレベルのバインディング、呼び出しごとの仕様は存在しません。レコードが欲しいなら、あなたはそれを自分でデザインしているのです。
Or finding the other people who did.これは、まさにこのためのスキーマを公開する小さなオープンソース プロジェクトです。信頼境界アクションごとに 1 つのレコード、RFC 8785 正規 JSON 上でチェーンされた SHA-256、レコード数とチェーン フィンガープリントのみを保存する外部監視、オペレーターが制御していない機関からの RFC 3161 タイムスタンプ。 About 4,800 lines of Python, zero dependencies.私たちはメモを比較したことがなく、同じ告白に至るまで同じ部品リストに収束しました。LIMITS.md は、システムが「オペレーターが最初からレコードを書き込んだことがないか、最近のレコードを削除して誰かが見る前に短いチェーンを再封印しなかったことを証明することはできません」とシステムが述べて始まります。これはまさに、私自身のギャップ検出器がこれからセクションを作成する譲歩です。これは、私がこの分野で、証明できない成果を上げているのを見た唯一のプロジェクトです。 2 つの独立した設計が同じ形状と同じ制限に到達すると、誰かがそれを書き留める直前の仕様がどのように見えるかが決まります。
の上

これはどの暗号よりも重要です。システム内のすべてのモデル バックエンドは単一のファクトリを通じて作成され、そのファクトリはキャプチャ プロキシで返されるものをすべてラップします。誰かが別の方法でバックエンドを構築した場合、静的分析テストはビルドに失敗します。これを行ったとき、呼び出しサイトは約 80 ありましたが、どれも変更しませんでした。開発者がコール サイトを計測することを忘れないことに依存するカバレッジは、スプリントごとに減衰します。
Hashes in the database, bytes elsewhere.各呼び出しでは、モデル、送信されたとおりのパラメーター、プロンプトのハッシュ、出力のハッシュ、トークン数、パイプライン ステージ、タイムスタンプの 17 個のフィールド行が書き込まれます。ペイロード バイトはコンテンツ アドレス ストレージに保存されます。この分割はストレージの最適化ではなく、特権の決定であり、私は最初にそれを間違えました。
議論の余地のないアイデンティティ。 17 個のフィールドは、ソートされたキーのコンパクトな JSON として 1 つのエンベロープ ハッシュにハッシュされ、ステップ ID は意図的にそのハッシュ内に配置されます。その選択の両側に罠があります。 ID を省略したままにすると、2 つの同一の再試行が一意のインデックスで衝突するため、正直な再試行ではキャプチャの失敗がスローされます。これを入れると一意性が自明のことになりますが、インデックスが何に変わるかに気づくまでは役に立たないように思えます。すでに識別子が含まれているハッシュに対する一意の制約は重複排除ではなくなります。これはトリップワイヤーです。2 つの行が 1 つを正当に共有することはできません。そのため、衝突は何かが歴史を書き換えたことを意味します。
Proof that nothing is missing.ほとんどの人がスキップするので、これは私が人々にコピーすることを最も勧めたい部分です。改ざん証拠は、存在する記録を保護し、一度も書き込まれなかった呼び出しについては何も言いません。これは、最も便利であるため、最も重要な障害モードです。

入口。したがって、呼び出しの前にインテント行を書き込みます。成功したら、レコードを書き込むのと同じコミットでキャプチャされるように反転します。失敗した場合は失敗に切り替えます。スイーパーは、猶予期間を超えてまだ保留中のものを、疑わしいギャップではなく、証明されたギャップとしてマークします。
最後のメカニズムが依然として譲歩していることに注目してください。記録が不完全であることを証明できます。同じプロセスでインテントとレコードの両方が書き込まれるため、レコードが完全であることを証明することはできません。
私は「再現可能」という言葉が一つのことを意味するかのように使っていました。これは 3 つを意味し、それぞれが異なる棚に置かれています。
記録は固定されており、改ざんが明らかです。すべてのステップがキャプチャされ、すべての引用がバイトに解決され、ハッシュされたものにまだハッシュされており、完全性は証明可能です。これはどこにでも適用されますが、実際に重要なのはバーです。
ステップを再実行し、返された内容を比較します。難しいのは再実行ではなく、比較です。モデルが 2 つの出力が同等かどうかを判断する場合、その検証自体は検証不可能となり、亀の塔を構築したことになります。唯一の正直なバージョンは構造を比較します。同じ引用が解決されたか、同じエンティティが抽出されたか、評決が同じバンドに収まったか。決定論的に区別できないものは、主張するのではなく、主張から除外する必要があります。
また、ビット正確な再生が行われますが、これはソフトウェアではなく展開の特性です。決定論的推論モードで固定されたオープンウェイト、はい。ホスト型 API、いいえ、私がいくらエンジニアリングしてもそれは変わりません。
この問題は法医学が私たちよりも先に解決しており、この前例は AI 文献の何よりも優れています。確率的ジェノタイピング ソフトウェアは非決定的です。SWGDAM のガイドラインでは、これらのアプローチでは同じ尤度比が得られない可能性があると記載されています。

分析を繰り返し、値の範囲を実証し、許容可能な変動量を確立することを研究室に要求します。国際ガイダンスでは、ソフトウェアが再現性テストのための安定したモードを提供することを求めていますが、これはまさにシードです。出版された研究では、実行間の対数尤度比に最大 10 倍の変動があることがわかりました。そして2026年3月、第三巡回裁判所は先例として、主要なツールは完璧ではないかもしれないが、ほとんどの科学は完璧ではなく、十分信頼できるものであるとの判決を下した。
したがって、「非決定論的な方法がどのように証拠になり得るか」に対する答えは理論的ではなく、決定論でもありません。それは、変動を測定し、それを開示し、プロセスを文書化することです。どの方向に切れるかに注意してください。法医学はその非決定性を測定します。それは私が言っているよりも強い要求です。
私がこれの設計仕様書を書き、その日のうちにレッドチームがそれを作成しました。今ではそれが当然のこととして行われています。評決は改訂され、あらゆる点で正しかった。
私自身のアーキテクチャにおける、私自身の譲れないルールの 1 つへの違反が見つかりました。この仕様では、キャプチャされたペイロードを通常のコンテンツ アドレス バケットに置きます。しかし、ここでのプロンプトには、法的に特権が与えられていることが判明する可能性のある文書からのテキストがそのまま掲載されるのが日常的です。

[切り捨てられた]

## Original Extract

The court reran an expert's AI query and got three different numbers. Here is the record I now keep for every model call, and the smaller thing it honestly proves.

Johannes Stillig
Writing
About
Switch to
dark
light
mode
A judge ran the same prompt three times and got three different answers
I'm Johannes. I build an investigations engine where every finding has to trace back to the source bytes it came from, which turned out to be an awkward promise once a language model started helping produce the findings.
In a trust accounting case in New York, an expert witness told the court he had used an AI assistant to cross-check his damages calculation. The judge asked him what he had typed into it. He could not remember. The judge asked what sources it had used. He could not say.
So the judge typed the question in himself, on a court computer, and got $949,070.97. He ran it again on a second court computer and got $948,209.63. A third returned a little over $951,000.
Three court-issued machines, one question, three answers. The Surrogate's Court wrote that while the variations were not large, the fact that there were variations at all called the reliability of the output into question, and held that counsel has an affirmative duty to disclose the use of AI before such evidence is admitted.
That is Matter of Weber , decided in October 2024. The expert was not caught fabricating anything. He was caught being unable to say what he asked or what came back.
I have spent the last month making sure that cannot happen to my product. What I built works, and it proves less than I originally claimed it did, and the gap between those two things is worth writing down.
Reproducibility is the wrong thing to want first
The instinct, when you notice this problem, is to reach for reproducibility. Pin the parameters, set a seed, get the same answer twice.
You cannot, and the reason is more interesting than floating point.
Thinking Machines Lab published the clearest account of it. The usual explanation, that concurrent floating point addition is non-associative, turns out to be wrong: run the same matrix multiplication on the same data repeatedly and you get bitwise identical results every time. The actual culprit is that your output depends on the batch you happened to be in, and the batch depends on who else was hitting the server at that moment. Their measurement: a thousand completions of one prompt at temperature zero produced eighty unique completions . The first 102 tokens were identical every time. Then they diverged.
From your point of view, the other users of the API are not an input. They are a property of the weather.
This is fixable, and it has shipped. Batch invariant kernels give bit identical output, and both vLLM and SGLang now expose a deterministic mode. The cost is real, somewhere between a quarter and two thirds of your throughput depending on whose measurement you take, and there is a nastier property underneath: determinism is contagious. One deterministic request joining a batch of ten ordinary ones dropped total server throughput from 931 tokens per second to 415. It is not a per-request option you can quietly buy for yourself.
And none of it is available to you on a hosted API. The Anthropic Messages API has no seed. On current frontier Claude models you cannot even set temperature: a non-default value returns a 400. Bedrock's Converse inferenceConfig has four knobs and none of them is a seed. OpenAI, the one provider that offered a seed and a backend fingerprint, has marked both deprecated: true in its OpenAPI spec and shipped neither on the Responses API.
So the industry's one reproducibility affordance is being withdrawn while everyone talks about AI governance.
Which is fine, because reproducibility was the wrong thing to want first. Logging is attestation, not reproduction. It proves what you sent and what came back. It does not prove the model would say it again. Those are different properties, and only one of them is available to you today.
The expert in Weber did not lose because his answer was irreproducible. He lost because he had no record.
Nobody has specified what to record
Here is what surprised me most. I went looking for the standard that says what to capture about a single model call, expecting to find several and have to pick.
"Traceability" appears zero times in the NIST AI Risk Management Framework. So does "record". The Generative AI Profile says "provenance" sixty-seven times, and every one of them is about content provenance or training data, because the problem NIST is solving there is deepfakes, not defensibility. Every AI bill-of-materials format describes the model rather than the call: CycloneDX's inputs and outputs fields hold formats , not values. Sigstore's model signing signs a list of file digests, which tells you the weights were not tampered with and nothing about what they did.
The closest thing to a specification comes from an unexpected direction. OpenTelemetry's GenAI conventions define exactly the attributes you would want, including the input messages, the output messages and the system instructions. Then the spec says instrumentations should not capture them by default, and lists as option one: do not record instructions, inputs, or outputs.
That is the best argument against everything I am saying, and it is also my whole point, read from the other end. The standard proves the capability is trivial. It also documents that the deliberate default is the evidentiary hole.
As for regulation: the EU AI Act does require high-risk systems to support automatic event logging over their lifetime, and requires providers and deployers to keep those logs for at least six months. Two things about that. The only place the Act specifies log content is a subsection covering remote biometric identification, and it reaches the input data for which a search produced a match, which is a long way from "record the prompt." And the obligations were due to apply on 2 August 2026, until the Digital Omnibus on AI pushed them to 2 December 2027 . That amendment came into force on 27 July 2026. I had the August date in a draft of this post three days ago.
The field-level, binding, per-invocation specification does not exist. If you want the record, you are designing it yourself.
Or finding the other people who did. After I built mine I came across halo-record , a small open-source project that publishes a schema for exactly this: one record per trust-boundary action, SHA-256 chained over RFC 8785 canonical JSON, an external witness that stores nothing but a record count and a chain fingerprint, RFC 3161 timestamps from an authority the operator does not control. About 4,800 lines of Python, zero dependencies. We had never compared notes, and we converged on the same parts list, down to the same confession: its LIMITS.md opens by stating the system "cannot prove the operator never wrote a record in the first place, or did not delete recent records and re-seal a shorter chain before anyone saw it", which is precisely the concession my own gap detector makes a section from now. It is the only project in this space I have seen lead with what it cannot prove. Two independent designs arriving at the same shape and the same limits is what a specification looks like just before somebody writes it down.
One seam, and this matters more than any of the cryptography. Every model backend in the system is created through a single factory, and that factory wraps whatever it returns in a capturing proxy. A static analysis test fails the build if anyone constructs a backend another way. There were around eighty call sites when I did this and I changed none of them. Coverage that depends on developers remembering to instrument their call site is coverage that decays every sprint.
Hashes in the database, bytes elsewhere. Each call writes a seventeen field row: model, parameters exactly as sent, a hash of the prompt, a hash of the output, token counts, pipeline stage, timestamp. The payload bytes go to content addressed storage. That split is not a storage optimisation, it is a privilege decision, and I got it wrong the first time.
An identity you cannot argue with. The seventeen fields are hashed as sorted-key compact JSON into one envelope hash, and the step id sits deliberately inside that hash. There is a trap on both sides of that choice. Leave the id out and two identical retries collide on the unique index, so an honest retry starts throwing capture failures. Put it in and uniqueness becomes trivially true, which sounds useless until you notice what it turns the index into. A unique constraint over a hash that already contains an identifier is not deduplication any more. It is a tripwire: two rows can never legitimately share one, so a collision means something rewrote history.
Proof that nothing is missing. This is the part I would most encourage people to copy, because almost everyone skips it. Tamper evidence protects the records that exist and says nothing about the call that never wrote one, which is the failure mode that matters most because it is also the most convenient. So: write an intent row before the call. On success, flip it to captured in the same commit that writes the record. On failure, flip it to failed. A sweeper marks anything still pending past its grace window as a proven gap rather than a suspected one.
Note what that last mechanism still concedes. It can prove the record is incomplete. It can never prove the record is complete, because the same process writes both the intent and the record.
I had been using "reproducible" as though it meant one thing. It means three, and they sit on different shelves.
There is the record being anchored and tamper evident: every step captured, every citation resolving to bytes that still hash to what they hashed to, completeness provable. This ships everywhere and it is the bar that actually matters.
There is re-executing a step and comparing what comes back. The hard part is not the re-execution, it is the comparison. If a model judges whether two outputs are equivalent then your verification is itself unverifiable, and you have built a tower of turtles. The only honest version compares structure: did the same citations resolve, were the same entities extracted, did the verdict land in the same band. Anything that cannot be diffed deterministically has to be excluded from the claim rather than waved at.
And there is bit exact replay, which is a property of the deployment rather than the software. Pinned open weights in a deterministic inference mode, yes. Hosted API, no, and no amount of engineering on my side changes that.
Forensic science worked this out before we did, and the precedent is better than anything in the AI literature. Probabilistic genotyping software is non-deterministic: SWGDAM's guidelines state that these approaches may not produce the same likelihood ratio from repeat analyses, and require laboratories to demonstrate the range of values and establish an acceptable amount of variation. The international guidance asks that the software offer a stable mode for repeatability testing, which is precisely a seed. Published work finds up to ten-fold swings in log likelihood ratio between runs. And in March 2026 the Third Circuit held, precedentially, that the leading tool may not be perfect but most science is not, and it is reliable enough.
So the answer to "how can a non-deterministic method be evidence" is not theoretical, and it is not determinism. It is: measure the variation, disclose it, document the process. Note which way that cuts. Forensic science measures its non-determinism. That is a stronger demand than the one I am making.
I wrote the design spec for this and red teamed it the same day, which I now do as a matter of course. The verdict came back REVISE and it was right on every count.
It found a violation of one of my own non-negotiable rules, in my own architecture. The spec put captured payloads in an ordinary content addressed bucket. But prompts here routinely carry verbatim text from documents that may turn out to be legally privileged, b

[truncated]
