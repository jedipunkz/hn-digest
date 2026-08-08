---
source: "https://castform.com/blog/unslop/"
hn_url: "https://news.ycombinator.com/item?id=49218397"
title: "I fought the slop and I won (unslopping LLM output)"
article_title: "i rl-finetuned an llm to unslop my writing - castform blog"
author: "BoredomIsFun"
captured_at: "2026-08-08T02:50:40Z"
capture_tool: "hn-digest"
hn_id: 49218397
score: 1
comments: 0
posted_at: "2026-08-08T02:27:40Z"
tags:
  - hacker-news
  - translated
---

# I fought the slop and I won (unslopping LLM output)

- HN: [49218397](https://news.ycombinator.com/item?id=49218397)
- Source: [castform.com](https://castform.com/blog/unslop/)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T02:27:40Z

## Translation

タイトル: スロップと戦って勝利しました (スロップのない LLM 出力)
記事のタイトル: 執筆の遅れをなくすために LLM を rl 微調整しました - Castform ブログ
説明: 強化学習を使用して 4b モデルをトレーニングし、AI のように聞こえるテキストを人間の言葉に戻すように書き換えます。報酬として AI 検出器を使用し、モデルがAI 検出器をハッキングすることを学習したときに何が起こったかを確認します。

記事本文:
執筆の遅れをなくすために LLM を rl 微調整しました - Castform ブログ
価格設定に関するドキュメント ブログを学ぶ キャリアを開始 トレーニングブック デモ
価格設定のドキュメント ブログを学ぶ キャリアを学ぶ トレーニングブックのデモを開始する 執筆の遅れをなくすために LLM を rl 微調整しました
tl;dr: 週末のプロジェクトとして、強化学習を使用して 4b モデルをトレーニングし、AI 検出器に勝つためにテキストを書き換えて「緩みをなくす」ようにしました。私はクロードにこの完全なブログ投稿を生成してもらい、それから上記の微調整されたモデルを取得して書き直しました。これがトレーニングの様子です。
私は最近たくさんの執筆を行っており、多くの人と同じように、AI モデルは私のワークフローの定期的な一部です。空白のページのブロックを乗り越えて、つながりのないアイデアを具体化するのに役立ちます。
しかし問題は、文章がaiのように聞こえることです。
それは必ずしも悪い文章ではありません。それは単なる一般的な書き方です。一般にaiスロップと呼ばれます。
そこで私は疑問に思いました。テキストの傾斜を解消するように AI モデルをトレーニングできるでしょうか。つまり、AI モデルから一般的な「AI 傾斜」パターンを削除し、元のコンテンツ、構造、トーンを維持しながらこれを行うことで、より人間らしく聞こえる方法でテキストを書き直すことができるでしょうか。
これは素晴らしい強化学習問題 (そして楽しい週末プロジェクト) のように感じました。報酬関数として AI 検出器を使用し、それに対して最適化します。ここに私が試したことについてのメモがあります。
私は最初に AI スロップの合成データセットを作成しました。パイプラインには 2 つのステップがありました。1 つの LLM が最初にシナリオ (スラック メッセージ、上司への電子メール、学生の作文など) を生成しました。次に、別の llm がそのシナリオを採用し、実際にテキストを具体化しました。私の目標は、このテキストをより人間らしい方法で書き換えることができる LLM をトレーニングすることでした。
私の最初のバージョンでは、報酬関数として 1 つの AI 検出器 (slop-guard) から始めました。スロップガードは「AI による定型的な文章パターンに対してテキストに 0 ～ 100 のスコアを付ける、ルールベースの散文リンター」です。

そうだね。」それは純粋にプログラム的であり、「200 以上のリテラルおよび構造ヒューリスティックに裏付けられた 24 の構成可能なルール」を備えています。どのルールに違反したかに基づいて数値スコアを返します。つまり、違反が少なければ少ないほど、報酬は高くなります。
残念ながら、この報酬は非常にハッキング可能でした。モデルは崩壊して、元のコンテンツをまったく保持せずに単純な短い文を生成するだけでした。
オリジナルコンテンツ:
> 何年にもわたって入学者数が減少し続けた後、大学は入学者数を減少させると発表した。
> 3つの人文科学部を閉鎖し、40人の教員ポストを削減し、
> 資金をエンジニアリングとコンピューターサイエンスに振り向ける。学生たちは抗議した
> この決定は、以下のような低所得の学生に不当に害を及ぼすことになるだろうということ。
> それらのプログラムに依存していました。
報酬ハッキングによる書き換え:
> 大学は人文科学の 3 つの学部を閉鎖します。
これを修正するために、コンテンツの保存を測定する別の報酬関数を追加しました。具体的には、情報がどの程度保存されているかを採点するために、元のコンテンツと書き換えられたコンテンツを LLM 審査員に渡しました。これは意味の保存に非常に役立ちました。
確率的AI検出器への拡張
スロップガードは良いスタートですが、プログラムのルールに依存しすぎます。そこで、他のテクニックを追加することにしました。私は、huggingface から微調整されたバート モデル ( modernbert-ai-detection-raid-mage ) から始めました。これを使ったトレーニングは概ね良好に見えましたが、たとえソースが正式なものであっても、モデルは非常にカジュアルなスタイルでテキストを書き換える傾向があることに気づきました。これを回避するために、別のトーン保存報酬を追加しました → 特に別の LLM ジャッジを使用してトーンの逸脱を検出しました。
オリジナルコンテンツ:
> 委員会は、提案された政策はかなりの負担を課すことになると結論付けた
> 対応する改善をもたらさない管理コスト
> 公共

結果。そのため実施を延期するよう勧告した
>さらなる検討が保留されています。
書き換えます:
> それで委員会は基本的に「これは非常に面倒なことになるだろう」という感じでした。
> なんとかなりますが、おそらくあまり役に立ちません。」それで彼らは言いました、「我慢しましょう」
> 電源を切って、もう一度見てみましょう。」
キッチンのシンクを投げる
上記の作業が完了したので、キッチンのシンクに問題を解決することにしました。私はオープンソースの世界にある大量の AI 検出器を集約し、それらを組み合わせた出力を報酬関数として使用しました。これには、diveye、tmr、bert-tiny-raid などが含まれます。
上記のモデルは、オープンソースの AI 検出器が人間であると判断した方法でテキストを書くことに関してはかなりまともですが、パングラムなどのクローズドな方法ではあまりうまくいかないことに気づきました (彼らに敬意を表します!)
オープンソースの検出器には強いですが、パングラムのようなクローズドな検出器には弱いです。
35b モデルにスケールアップしても、検出器間の一般化には役に立ちませんでした。ここでは、クローズドソース検出器への一般化可能性の向上に関して、さらなる作業を行う必要があります。ここで私が考えているいくつかのアイデア:
単純なオプション: 多額の $$ を費やして、閉じた検出器で直接トレーニングするだけです
閉じた検出器からの検出を別のモデルに抽出し、それを報酬メトリクスとして使用します
トレーニング全体を通じて検出器の重み、しきい値、プロンプト、サンプリング設定を変更する
ガンのようなループでトレーニングします。リライターは継続的に更新される検出器を騙す方法を学習しますが、検出器はリライターの最新の出力で再トレーニングされます。
rl は非常に反復的です。通常、最初の報酬はハッキングされるため、作業のほとんどは出力の検査、障害モードの発見、報酬関数の改善に行われます。より良いモデルは、より良い報酬、より良いデータ、より良い評価よりも役に立ちます 🙂。
rlft で新しいガイドや製品の更新情報を入手してください。
© 2026 キャストフォーム。無断転載を禁じます。

## Original Extract

training a 4b model with reinforcement learning to rewrite ai-sounding text so it reads human again — using ai detectors as the reward, and what happened when the model learned to hack them.

i rl-finetuned an llm to unslop my writing - castform blog
c a s t f o r m pricing docs learn blog careers start training book demo
pricing docs learn blog careers start training book demo i rl-finetuned an llm to unslop my writing
tl;dr: as a weekend project, i trained a 4b model with reinforcement learning to rewrite & “unslop” text to beat ai detectors. i got claude to generate this full blogpost and then got the above finetuned model to rewrite it. here’s the training run .
i have been doing a lot of writing recently and like many people, ai models are a regular part of my workflow. it’s useful to get past the blank page block & flesh out disconnected ideas.
the problem though is that the writing sounds like ai.
it is not necessarily bad writing. it is just generic writing. commonly referred to as ai slop.
so i wondered: could i train an ai model to unslopify text i.e., rewrite text in a way that sounded more human by removing the common “ai slop” patterns from ai models and doing this while preserving the original content, structure and tone.
this felt like a nice reinforcement learning problem (and a fun weekend project): use an ai detector as a reward function and optimize against it. here are some notes on the things i tried.
i first created a synthetic dataset of ai slop. the pipeline had two steps: one llm first generated a scenario (e.g., slack message, email to superior, student essay, etc.). second, another llm took that scenario and actually fleshed out the text. my goal was to train an llm that can rewrite this text in a more human way.
for my first version, i started with a single ai-detector ( slop-guard ) as the reward function. slop-guard is “a rule-based prose linter that scores text 0–100 for formulaic ai writing patterns.” it’s purely programmatic, with “24 configurable rules backed by 200+ literal and structural heuristics.” it returns a numeric score based on which rules were violated. so effectively, the fewer the violations, the higher the reward.
unfortunately, this reward was quite hackable → the model just collapsed to producing simple short sentences while not preserving any of the original content.
original content:
> After years of declining enrollment, the university announced that it would
> close three humanities departments, eliminate 40 faculty positions, and
> redirect funding toward engineering and computer science. Students protested
> that the decision would disproportionately harm low-income students who
> relied on those programs.
reward-hacked rewrite:
> The university is closing three humanities departments.
to fix this, i added another reward function to measure content preservation. specifically, i passed the original content and the rewritten content to an llm judge to grade how well information was preserved. that helped a lot with semantic preservation.
extending to probabilistic ai detectors
slop-guard is a good start, but it’s too dependent on programmatic rules. so i decided to add other techniques. i started with a fine-tuned bert model off huggingface ( modernbert-ai-detection-raid-mage ). training with it generally looked good, but i noticed that the model tended to rewrite text in a very casual style, even if the source was formal. to get around this, i added another tone preservation reward → specifically using another llm judge to detect any deviations in tone.
original content:
> The committee concluded that the proposed policy would impose substantial
> administrative costs without producing a corresponding improvement in
> public outcomes. It therefore recommended that implementation be postponed
> pending further review.
rewrite:
> so the committee was basically like, "This is going to be a huge pain to
> manage, and it probably won't even help much." So they said, "Let's hold
> off and take another look."
throwing the kitchen sink
with the above working, i decided to throw the kitchen sink at the problem. i aggregated a whole bunch of ai detectors in the open-source world and used their combined output as a reward function. this included diveye, tmr, bert-tiny-raid, etc.
while the above model is pretty decent at writing text in a way that open-source ai detectors said was human, i did notice that it wasn’t so good with closed ones like pangram, etc. (kudos to them!)
strong against open-source detectors, weaker against closed ones like pangram.
scaling up to a 35b model didn’t help with cross-detector generalization either. more work needs to be done here around improving generalizability to closed-source detectors. some ideas i have here:
the naive option: just spend lots of $$s and train directly with closed detectors
distill detection from closed detectors into a separate model and use that as a reward metric
vary detector weights, thresholds, prompts, and sampling settings throughout training
train in a gan-like loop where the rewriter learns to fool a continually updated detector, while the detector is retrained on the rewriter’s latest outputs
rl is very iterative. the first reward usually gets hacked, so most of the work is in inspecting outputs, finding failure modes, and improving the reward functions. better models helped less than better rewards, better data, and better evals 🙂.
stay updated with new guides and product updates on rlft.
© 2026 castform. all rights reserved.
