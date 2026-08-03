---
source: "https://bymorning.ai/journal/ai-style-guide"
hn_url: "https://news.ycombinator.com/item?id=49161229"
title: "Can an AI Model Obey a Strict Style Guide?"
article_title: "Can an AI Model Obey a Strict Style Guide? I Tested Ten of Them | Bymorning"
author: "johnrising"
captured_at: "2026-08-03T20:55:06Z"
capture_tool: "hn-digest"
hn_id: 49161229
score: 1
comments: 1
posted_at: "2026-08-03T20:49:04Z"
tags:
  - hacker-news
  - translated
---

# Can an AI Model Obey a Strict Style Guide?

- HN: [49161229](https://news.ycombinator.com/item?id=49161229)
- Source: [bymorning.ai](https://bymorning.ai/journal/ai-style-guide)
- Score: 1
- Comments: 1
- Posted: 2026-08-03T20:49:04Z

## Translation

タイトル: AI モデルは厳格なスタイル ガイドに従うことができますか?
記事のタイトル: AI モデルは厳格なスタイル ガイドに従うことができますか?そのうちの 10 個をテストしました |午前中までに
説明: ASD-STE100 簡易技術英語に対して 10 個の LLM をテストしました。プロンプトはほとんどの場合に役立ちます。完全なコンプライアンスにはチェッカーが関与します。

記事本文:
AI モデルは厳格なスタイル ガイドに従うことができますか?そのうちの 10 個をテストしました | ByMorning Docs Journal Team デモを見る ← Journal Engineering John Rising · 2026 年 8 月 1 日 AI モデルは厳格なスタイル ガイドに従うことができますか?そのうちの 10 個をテストしました
大規模な言語モデルは流暢なテキストを記述しますが、ルールにはうまく従いません。明確な技術的な散文を要求すると、いくつかの小さな変更を加えた、モデルが好むスタイルのテキストが得られます。これは、自動化された LinkedIn のスロップ投稿には問題ありませんが、整備士が各文から 1 つのことしか理解しなければならない航空機整備マニュアルには適していません。
航空業界は、AI が登場するずっと前に、約 53 のルールと 900 の承認済み単語からなるルールブックである ASD-STE100 簡易技術英語 (STE) を使用して、この問題を解決しました。 STE は、自然な英語の原動力であり、AI ライティングで普及している受動態、「-ing」動詞形、ソフト法助動詞を禁止しています。 AI モデルはそのようなルールブックに従うことができるでしょうか?
余談ですが、いくつかの意見の相違があるため、さまざまなアプローチとモデルをテストする実験を構築しました。簡単に言うと、プロンプトを使用してある程度の方法は取得できますが、完全にコンプライアンスに準拠するには、外部ツールに対するチェックが必要です。
このテストには 4 つの部分があります。ルール違反をカウントするチェッカー、固定された一連の書き込みタスク、段階的に作業を追加する 4 つのプロンプト条件、およびチェッカーの結果をモデルに送り返すフィードバック ループです。各実行は、チェッカーによってカウントされた 100 単語ごとの違反によってスコア付けされます。低いほど良いです。
チェッカーはテキストを取り込み、ASD-STE100 からの 10 ルール ファミリの違反数をカウントします。
文内の複数の指示
未承認の単語のリストとその承認済みの置換
各モデルには、航空機のメンテナンスからネットワーク ハードウェアまで、6 つのドメインをカバーする同じ 10 のタスクが割り当てられました。 10 個のタスクのうち、6 個が手順です

4 は説明です。 6 件は不適切なソース テキストの書き直しで、4 件は白紙のページから始まります。
ソース テキストを、エンジン駆動の油圧ポンプの番号付きの段階的な取り外し手順として書き換えます。
各タスクは 4 つのプロンプト条件の下で実行されたため、各モデルは合計 40 のプロンプトに回答しました。条件によって段階的に努力が追加されます。
C0: 「明確で正確な技術的な文章を書く」。プロンプトには STE という名前が表示されません。
C1: ASD-STE100 の名前を追加した文が 1 つあります。これはモデルがすでに知っていることを測定します。
C2: プロンプト内の STE ルールの 1 ページのダイジェスト。
C3: ダイジェストと 3 つの前後の例。
最後の部分では、パイプラインの「修復」ステップをテストします。モデルは、違反を削除し、すべての技術コンテンツを保持するように指示された、独自のテキストと違反のリストをチェッカーから取得します。各ドキュメントは最大 3 ラウンドまで行われ、それぞれが新しいセッションで行われ、プロンプトにはルール ダイジェストは表示されません。ループは C1 出力から始まります。
この表は、条件ごとの 100 単語あたりの平均違反を示しています。低いほど良いです。並べ替えは、基本的なクロード スキルまたはプロンプトの最適なプロキシである C2 によって行われます。
一文が大きな利益をもたらします。プロンプトで標準 (C1) のみが指定されている場合、すべてのモデルが改善され、25 ～ 73 パーセントの向上が見られます。したがって、すべてのモデル家族は教育によって STE について何らかの知識を持っていますが、その深さは 3 倍異なります。 1 ページのルール ダイジェストにより、ほとんどのモデルで残りが 3 分の 1 から 2 分の 1 に削減され、迅速な作業が増加するにつれて、モデル間の差異が縮小します。
Claude Sonnet 5 は、すべてのフロンティア モデルのパフォーマンスが非常に似ていますが、名簿上で最高の STE ライターです。上位 4 つのモデルのうち 3 つは AWS GovCloud で利用できるため、航空宇宙チームや防衛チームは輸出規制されたデータに簡単にアクセスできます。
より迅速な対応が常に良いとは限りませんが、

うーん。前後の 3 つの例を追加すると、9 つのモデルが改善されましたが、GPT のフラッグシップ モデルは悪化し、1.96 から 2.19 になりました。私の推測では、サンプルはルールの実行から離れて、表面スタイルの模倣に向けてモデルを引っ張っているのではないかと思います。
残りの違反はルール全体に均等に分散されていません。 「-ing」の禁止は強制するのが非常に困難です。40 個のモデルと条件のセルのうち 37 個では、「-ing」動詞の形式が違反のトップとなっています。ただし、より明確なトリガーワードを備えた同様のルールは、適切なプロンプトの下では崩壊します。たとえば受動態は、5 つの最高のモデルのうち 4 つ (Sonnet 5、GPT フラッグシップ、Opus 4.8、Sonnet 4.6、ただし 6 つを維持した Sonnet 4.5 は除く) で C3 までにゼロになりますが、「-ing」は残ります。
ただし、この指標だけでは、出力が適切であるかどうかはわかりません。テキストが規格の基準にどの程度準拠しているかだけです。実際、小さなモデルは切断することで「ごまかします」が、大きなモデルはそうではありません。小型のラマモデルは最悪でした。 C3 スコアは一見良好であるにもかかわらず、テキストが理解できないほど縮小されました。
フィードバック ループは、実際に完全なコンプライアンスが達成される場所です。プロンプトのみのメソッドは頭打ちになります。完全な C3 プロンプトは各世代で約 1,200 ワードを消費しますが、依然として違反が残ります。修復ループはドキュメントごとに 1 つの短い違反リストを使用し、最良のモデルでは 10 個のドキュメントのうち 8 個がチェッカー独自のメトリクスでゼロまで追い込まれました。これは他のアプリケーションでも見られるものと一致します。LLM は目標までの約 80% を達成できますが、自明ではない目標を完全に達成するにはツールを操作する必要があります。
モデル全体で、コンバージェンスはリーダーボードを大まかに追跡しました。最も良く修復されたのはフロンティア モデルです。Sonnet 5、Opus 4.8、GPT mini はそれぞれ 10 タスク中 8 タスクをゼロにしましたが、GPT フラッグシップモデルは 7 タスクかかりました。ソネット 4

.6も7になりました。 Haiku 4.5 は 10 点中 6 点を獲得し、1 ラウンド遅くなりました。 Sonnet 4.5 は 1 つのパスのみを獲得し、違反を 81% 削減し、3 つのタスクでゼロに達しました。
ミストラルは一度収束し、その後振動しました。 6 つのタスクでは、ある段階でカウントが増加しました。 2 匹のラマは使用可能なリビジョンをまったく生成しませんでした。
ただし、収束だけでも不正な指標です。すべての技術コンテンツを保持するというループの指示は単なるリクエストであり、チェッカーはモデルが何を削除するかを確認できません。そのため、修復されたすべての文書は、オリジナルとタスク仕様に対する単語ごとの監査という忠実度の監視も通過しました (付録 C)。ガードが順位を変える。 Sonnet 4.6 は 7 つのタスクでゼロに達しましたが、完全に忠実であったのはそのうちの 4 つだけでした。ソネット 4.5 は 3 つでゼロに到達し、3 つすべてを忠実に保ちました。
2 つのメトリクスを相互にプロットすると、古いモデルと中層モデルの間にパレート フロンティアが現れます。つまり、コンプライアンスを高めると忠実度がある程度犠牲になります。 Sonnet 4.5、Haiku 4.5、Sonnet 4.6 はそれぞれ、トレードオフの異なる点でそのフロンティアに位置しています。ミストラルはその中にあり、忠実度ではソネット 4.5 に匹敵しますが、コンバージェンスでは負けています。しかし、フロンティアモデルはそのトレードオフを打破しました。 Sonnet 5 は 8 つのタスクでゼロに到達し、8 つすべてで忠実を維持しました。これは、忠実度のコストがかからずに準拠した唯一のモデルです。Opus 4.8 と両方の GPT モデルは、そのすぐ下に集まっています。最近テストされた 4 つのフロンティア モデルでは、いずれも大きな損失は 1 つも発生しませんでした。
0
2
4
6
8
10
0
25
50
75
100
タスクがゼロ (10 個中)
忠実なタスクの割合がゼロ (%)
ミストラルラージ3
クロード・ソネット 4.5
クロード俳句 4.5
クロード・ソネット 4.6
クロード・ソネット 5
クロード オーパス 4.8、GPT Terra
GPT-5.6ソル
これをどうするか
標準に書き込むパイプラインを構築する場合、次のような教訓が得られます。
努力を続けるのではなく、ループに取り組みましょう

プロンプトが表示されません。標準に名前を付け、ルールを一度述べ、残りの労力を LLM の外部の評価フレームワークに費やします。
多ければ多いほど良いというわけではありません。より多くの例が、ある旗艦を悪化させました。
チェッカーの精度を安全特性として扱います。チェッカーをループに配線する前に、チェッカーの偽フラグを監査してください。
AI は、定められた目標に違反する場合でも指標に合わせて最適化することが多いため、AI が作成したテキストを盲目的に受け入れないでください。
A. 一文から得られる利益。このグラフは、プロンプトで標準名を指定し、他に何も追加しない場合の、C0 から C1 への違反の減少をパーセントで示しています。
xychart-ベータ版
タイトル「違反件数が C0 から C1 に減少（パーセント）」
x 軸 ["Sonnet 5"、"Opus 4.8"、"Sonnet 4.5"、"GPT Sol"、"Mistral"、"Llama 70B"、"GPT Terra"、"Haiku 4.5"、"Sonnet 4.6"、"Llama 8B"]
y 軸「パーセント」 0 --> 80
bar [73, 66, 65, 61, 54, 50, 39, 38, 25, 25] B. 修復ループ。各モデルは、違反リストとともに最大 3 ラウンドの 10 個の C1 出力を取得しました。この表は、10 個のタスクの前後における違反の合計を示しています。
2 匹のラマは表にありません。彼らのループでは使用可能なリビジョンが生成されませんでした。
C. 忠実性の監査。私は修復された各文書を一語一語、その文書と照らし合わせて監査しました。
オリジナルとそのタスク仕様。この表はドキュメント上のみで各モデルをスコア付けしています。
完全なコンプライアンス（違反ゼロ）に達した企業: そのうち、何社が遵守を続けましたか
ソースは？ 「クリーンコンプライアンス」では、違反ゼロで終了したタスクと、違反なしで終了したタスクをカウントします。
完全に忠実 — 本番パイプラインが受け入れる必要がある唯一の結果。
エンジニアリング バイモーニングの仕組み

## Original Extract

I tested ten LLMs against ASD-STE100 Simplified Technical English. Prompting gets most of the way there; full compliance takes a checker in the loop.

Can an AI Model Obey a Strict Style Guide? I Tested Ten of Them | Bymorning Docs Journal Team See a demo ← Journal Engineering John Rising · August 1, 2026 Can an AI Model Obey a Strict Style Guide? I Tested Ten of Them
Large language models write fluent text, but they do not obey rules well. Ask one for clear technical prose and you get text in the style the model prefers with a few small changes. That is fine for an automated LinkedIn slop post, but it is not fine for an aircraft maintenance manual, where a mechanic must understand only one thing from each sentence.
The aviation industry solved this long before AI with ASD-STE100 Simplified Technical English (STE), a rulebook of about 53 rules and 900 approved words. STE bans passive voice, "-ing" verb forms, and soft modal verbs, which are the fuel of natural English and prevalent in AI writing. Can an AI model obey a rulebook like that?
Anecdotally, there is some disagreement , so I built an experiment testing different approaches and models. The short answer is they can get some of the way with prompts, but full compliance requires checking against external tools.
The test has four parts: a checker that counts rule violations, a fixed set of writing tasks, four prompt conditions that add effort step by step, and a feedback loop that sends the checker's findings back to the model. Each run is scored by violations per 100 words, as counted by the checker. Lower is better.
The checker ingests text and counts the number of violations of 10 rule families from ASD-STE100:
more than one instruction in a sentence
a list of unapproved words, with their approved replacements
Each model got the same 10 tasks covering six domains, from aircraft maintenance to network hardware. Of the 10 tasks, 6 are procedures and 4 are descriptions. Six are rewrites of bad source text and four start from a blank page.
Rewrite the source text as a numbered, step-by-step removal procedure for the engine-driven hydraulic pump.
Each task ran under four prompt conditions, so each model answered 40 prompts in total. The conditions add effort step by step:
C0: "write clear, accurate technical prose". The prompt does not name STE.
C1: one added sentence that names ASD-STE100. This measures what the model already knows.
C2: a one-page digest of the STE rules, in the prompt.
C3: the digest plus three before-and-after examples.
The last part tests the "repair" step of the pipeline. The model gets its own text back plus the list of violations from the checker, with an instruction to remove the violations and keep all technical content. Each document gets up to three rounds, each in a fresh session, with no rules digest in the prompt. The loop starts from the C1 outputs.
The table shows mean violations per 100 words, by condition. Lower is better. The sort is by C2, the best proxy for a basic Claude Skill or prompt.
One sentence pays a large dividend. When the prompt only names the standard (C1), every model improves, with gains from 25 to 73 percent. So every model family knows something about STE from its education, but the depth varies by a factor of three. The one-page rules digest then cuts what remains by a third to a half for most models, and as prompt effort increases, the differences between the models shrink.
Claude Sonnet 5 is the best STE writer on the roster, though all frontier models perform very similarly. Three of the top four models are available on AWS GovCloud, and thus easily accessible to aerospace and defense teams with export controlled data.
More prompt is not always better, though. Nine models improved when I added the three before-and-after examples, but the GPT flagship model got worse: 1.96 to 2.19. My guess is that examples pull a model toward imitation of surface style, away from execution of rules.
The residual violations are not evenly spread across the rules. The "-ing" ban is very difficult to enforce: in 37 of the 40 model-and-condition cells, "-ing" verb forms are the top violation. However similar rules with a clearer trigger word collapse under a good prompt. Passive voice, for example, falls to zero by C3 for four of the five best models (Sonnet 5, the GPT flagship, Opus 4.8, and Sonnet 4.6 — though not Sonnet 4.5, which kept six), but "-ing" persists.
However, this metric alone does not tell us whether the output is any good, just how compliant the text is with the measures in the standard. In fact, small models would "cheat" by cutting meaning that the larger models did not. The small Llama model was the worst; it shrank text to incomprehensibility despite a seemingly good C3 score.
The feedback loop is where full compliance actually arrives. The prompt-only method plateaus: the full C3 prompt spends about 1,200 words on every generation and still leaves violations behind. The repair loop spends one short list of violations per document, and the best models drove eight of their ten documents all the way to zero on the checker's own metric. This matches what we see in other applications as well - the LLM can get you about 80% of the way to a goal, but fully achieving a non-trivial goal requires interacting with tools.
Across models, convergence roughly tracked the leaderboard. The frontier models repaired best: Sonnet 5, Opus 4.8, and the GPT mini each took eight of ten tasks to zero, and the GPT flagship took seven. Sonnet 4.6 also reached seven. Haiku 4.5 got six of ten, one round slower. Sonnet 4.5 got one pass only, cut its violations by 81 percent, and reached zero on three tasks.
Mistral converged once and then oscillated; on six tasks its count went up at some step. The two Llamas produced no usable revisions at all.
Convergence alone is a rigged metric, though. The loop's instruction to keep all technical content is only a request, and the checker cannot see what a model deletes. So every repaired document also passed through a fidelity guard: a word-by-word audit against the original and the task specification (Appendix C). The guard changes the standings. Sonnet 4.6 reached zero on seven tasks but stayed fully faithful on only four of them. Sonnet 4.5 reached zero on three and kept all three faithful.
Plot the two metrics against each other and, among the older and mid-tier models, a Pareto frontier appears: more compliance costs some fidelity. Sonnet 4.5, Haiku 4.5, and Sonnet 4.6 each sit on that frontier at a different point of the tradeoff; Mistral sits inside it, matched on fidelity by Sonnet 4.5 and beaten on convergence. But the frontier models broke the tradeoff. Sonnet 5 reached zero on eight tasks and stayed faithful on all eight — the only model whose compliance came at no fidelity cost — and Opus 4.8 and both GPT models cluster just below it. None of the four late-tested frontier models produced a single major loss.
0
2
4
6
8
10
0
25
50
75
100
Tasks at zero (of 10)
Faithful share of tasks at zero (%)
Mistral Large 3
Claude Sonnet 4.5
Claude Haiku 4.5
Claude Sonnet 4.6
Claude Sonnet 5
Claude Opus 4.8, GPT Terra
GPT-5.6 Sol
What to do with this
If you build a pipeline that writes to a standard, the transferable lessons are these:
Put your effort in the loop, not a good prompt. Name the standard, state the rules once, and spend the rest of your effort on an evaluation framework outside of the LLM.
More is not always better. More examples made one flagship worse.
Treat checker precision as a safety property. Audit your checker's false flags before you wire it into a loop.
Do not accept AI-written text blindly, as AI will often optimize for metrics even when they violate your stated goals.
A. The gain from one sentence. The chart shows the decrease in violations from C0 to C1, in percent, when the prompt names the standard and adds nothing else.
xychart-beta
title "Decrease in violations from C0 to C1 (percent)"
x-axis ["Sonnet 5", "Opus 4.8", "Sonnet 4.5", "GPT Sol", "Mistral", "Llama 70B", "GPT Terra", "Haiku 4.5", "Sonnet 4.6", "Llama 8B"]
y-axis "Percent" 0 --> 80
bar [73, 66, 65, 61, 54, 50, 39, 38, 25, 25] B. The repair loop. Each model got its ten C1 outputs back, with the violation list, for up to three rounds. The table shows total violations across the ten tasks, before and after.
The two Llamas are not in the table. Their loops produced no usable revisions.
C. The fidelity audit. I audited each repaired document, word by word, against its
original and its task specification. The table scores each model only on the documents
that reached full compliance (zero violations): of those, how many stayed faithful to
the source? "Clean compliance" counts tasks that ended both at zero violations and
fully faithful — the only outcome a production pipeline should accept.
Engineering How bymorning works
