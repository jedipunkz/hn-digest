---
source: "https://research.roundtable.ai/captchas-detect-ai/"
hn_url: "https://news.ycombinator.com/item?id=48324910"
title: "CAPTCHAs can still detect AI agents"
article_title: "CAPTCHAs can still detect AI agents | Roundtable Research"
author: "timshell"
captured_at: "2026-05-30T11:44:41Z"
capture_tool: "hn-digest"
hn_id: 48324910
score: 75
comments: 62
posted_at: "2026-05-29T15:57:37Z"
tags:
  - hacker-news
  - translated
---

# CAPTCHAs can still detect AI agents

- HN: [48324910](https://news.ycombinator.com/item?id=48324910)
- Source: [research.roundtable.ai](https://research.roundtable.ai/captchas-detect-ai/)
- Score: 75
- Comments: 62
- Posted: 2026-05-29T15:57:37Z

## Translation

タイトル: CAPTCHA は依然として AI エージェントを検出できる
記事のタイトル: CAPTCHA は依然として AI エージェントを検出できる |ラウンドテーブルリサーチ
説明: AI システムは現在、多くのタスクに関して人間の解決策と一致していますが、それらの解決策にはかなり異なるプロセスを介して到達します。このギャップを利用して、AI エージェントやオンライン ボットを検出できます。

記事本文:
CAPTCHA は引き続き AI エージェントを検出できます。ラウンドテーブルリサーチ
メインサイト
CAPTCHA は引き続き AI エージェントを検出できる
AI システムは現在、多くのタスクにおいて人間に匹敵し、人間を上回っていますが、明らかに異なる認知プロセスを通じて動作します。
このギャップを利用して、AI エージェントやオンライン ボットを検出できます。
「最近の CAPTCHA は壊れています。」 AI は静的なグリッド内のすべての信号機を簡単に識別できます。だからキャプチャ
貴重な人間の信号を提供していないですよね?
はい、ビジョン ランゲージ モデル (VLM) は煙突、消火栓、信号機などの画像を認識できるためです。
ディープラーニングは、2010 年代初頭に CAPTCHA スタイルの画像分類を「解決」しました。
いいえ、AI は人間のように CAPTCHA を完了しないからです。人間とAIが完成させたあらゆるデータを見渡せば
CAPTCHA を使用すると、エラー パターンなどの機能の違いに気づき始めます。私たちの最近の論文では、統計的に有意な差があることがわかりました。
連続クリック パターン、方向変更、および過剰選択動作 - 参加者がどのように操作するかを定義する機能
エージェントまたは人間が CAPTCHA の問題を解決します。言い換えれば、AI は CAPTCHA を解決できますが、解決することはできません。
人間のように。
チューリング テスト (もともと 1950 年にアラン チューリングによって提案されたもの) は、機械知能の単純な基準を提供します。もし
裁判官が機械の応答と人間の応答を確実に区別できない場合、機械は考慮される可能性があります。
インテリジェント。
チューリングは、この行動基準は譲歩であり、人間対機械の最終的な決定ではないと理解していました。
知性。彼は認めざるを得ませんでした。質問は難しすぎ、抽象的で、内容が多すぎます。行動的
区別がつかないことにより、より扱いやすい状態が得られ、1950 年代には良好な北極星のように見えました。
ハム音を分離できる敵対的に堅牢な識別器を定義したチューリングの足跡をたどる

ボットからのメッセージ、
CogCAPTCHA30を設計しました。これは、チューリング テストよりも 1 レベル深く、出力を調査することから始まります (何を
人間とエージェントができること）処理すること（どのように行うことができるか）。 CogCAPTCHA30 は、オリジナルの CAPTCHA と 29 を結合します。
30 タスクバッテリーの古典的な認知心理学のタスク。
これらのタスクを実行するために人間の参加者を募集し、AI エージェントも導入しました。 CAPTCHAの実験
人間とエージェントは同様のパフォーマンス (出力) レベルでパフォーマンスを発揮できるが、パフォーマンスが異なることを実証しました。
プロセス。次に、出力の同等性 - どのように (回答がどの程度類似しているか) を測定しました。
30 タスクのパラダイム全体にわたってプロセスの等価性 (どのようにして答えに到達したか) を調べたところ、相関がないことがわかりました。
古典的なチューリング テストでは、機械が人間と区別できない出力を生成するかどうかを測定しますが、
機械が人間と区別できないプロセスを生成するかどうかを測定するプロセス チューリング テストを提案します。
私たちの結果は 2 つの疑問を引き起こします。それは、言語モデルが存在する場合、どのタイプが人間に似ているのか、そしてどのように敵対的であるのかということです。
この差別プロセスは堅牢ですか?
最初の質問に答えるために、人間と最先端のフロンティアモデルとの距離を比較しました。
(OpenAI の GPT、Anthropic の Claude、Google DeepMind の Gemini) および Qwen (オープンソースの 1.5 億基盤)
モデル) と Centaur (人間の認知のオープンソース 70B パラメーター基礎モデル)。
最先端のフロンティア モデル (Claude、GPT、Gemini) には人間のプロセスの特徴があまり類似していないことがわかりました。
小型モデル (Qwen、Centaur) と比較して。 「 AI の能力は人間らしさではない 」で私たちが主張したように、
フロンティアモデルは時間の経過とともにより強力になっていますが、必ずしもより人間らしくなっているわけではありません。コンテンポラリー
人工知能の進歩は人間のシミュレーションの進歩とは無関係です。
クウェン、

小規模なオープンソース モデルは、大型の Claude、GPT、Gemini よりも人間らしいモデルです。そして、嬉しいこととして、
検証の結果、Centaur はヒューマン プロセス特徴空間との類似性において他のモデルよりも優れています。私たちはこれを仮説とします
これは、大規模な出力の微調整、特に 160 の認知実験にわたる 1,000 万以上の人間の選択によるものです。
これにより、2 番目の質問が生じます。人間とエージェントを区別するプロセスは、敵対的にどの程度堅牢ですか?
2 つを区別するために使用される動作の特徴は、それ自体が最適化のターゲットになる可能性があります。したがって、
既製のエージェントに対して成功した検出器は、現在の攻撃者の下でのみ動作ギャップを確立します
モデル - AI が現在どのように存在し、動作しているか。それが耐久性のある人間による検証信号になるかどうかはわかります
未来のテクノロジーのために。これにより、より強力なテストが行われます。つまり、エージェントは人間と人間のプロセス間のギャップを埋めることができるかということです。
そしてエージェントはタスクを完了します - 人間のデータへの直接アクセスがますます与えられるようになったとき?
Qwen2.5 Instruct モデルを微調整して人間に近づけました。完全な情報が与えられたとき - 観察された
特徴と識別子の目的関数 - 人間とエージェントの間のギャップがなくなります。ただし、ギャップは
特徴空間の一部が省略された場合に再表示され、エージェントが一般化する必要がある場合に完全に戻ります
クロスタスク。言い換えれば、プロセス チューリング テストは、AI が完全なアクセス権を持たない場合でも堅牢です。
識別子と特徴セット (つまり、モデルはそれがどのように評価されるかを知りません)。
プロセス チューリング テストが提起する課題は、AI が人間の認知機能のすべてを継続的に複製できるかどうかです。
心理学。時間の経過とともにモデルの能力が向上しているという不安にもかかわらず、経験的にはモデルの能力が向上しているわけではありません。
より人間らしく。パスワード、CAPTCHA、

文書の識別とデバイス
フィンガープリンティングと同様に、プロセス チューリング テストは人間による検証のステップアップ機能を提供します。人間を模倣する
認知心理学は、飛躍的に困難な課題です。
Roundtable は、インターネットの認証レイヤーを構築する人間認証会社です。
当社の Proof of Human システムは、人間のユーザー エクスペリエンスを維持しながら、AI エージェントとボットを目に見えずに検出します。
詳細については、roundtable.ai をご覧ください。
Mayank Agrawal、Milena Rmus、Mathew Hardy は Roundtable Technologies Inc. で働いており、そこで
Proof of Human、Web 用の目に見えない認証システム。
以前、彼らはプリンストン大学 (マヤンクとマット) と同大学で認知科学の博士号を取得しました。
カリフォルニア州バークレー（ミレーナ）出身。

## Original Extract

AI systems now match human solutions on many tasks, but reach those solutions via measurably different processes. This gap can be exploited to detect AI agents and online bots.

CAPTCHAs can still detect AI agents | Roundtable Research
Main site
CAPTCHAs can still detect AI agents
AI systems now match and exceed humans on many tasks, but behave through measurably different cognitive processes.
This gap can be exploited to detect AI agents and online bots.
"CAPTCHAs are broken these days." AI can easily identify all the traffic lights in a static grid. So CAPTCHAs
don't provide a valuable human signal, right?
Yes, because vision language models (VLMs) can recognize images like chimneys, fire hydrants, and traffic lights.
Deep learning "solved" CAPTCHA-style image classification in the early 2010s.
No, because AI does not complete CAPTCHAs like humans. If you look across all the data of humans and AI completing
CAPTCHAs, you start noticing differences in features like error patterns. Our recent paper found statistically significant differences across
sequential click patterns, direction changes, and overselection behavior - features that define how a participant,
agent or human, would solve the CAPTCHA problem. In other words, AI can solve CAPTCHAs, but they don't solve them
like humans.
The Turing Test - originally proposed in 1950 by Alan Turing - offers a simple criterion for machine intelligence. If
a judge cannot reliably distinguish a machine's responses from a human's, the machine can be considered
intelligent .
Turing understood this behavioral criterion was a concession and not the end-all-be-all of human vs. machine
intelligence. He had to concede: the question is too difficult, abstract, and loaded. Behavioral
indistinguishability provided a more tractable condition, and one that seemed like a good North Star in the 1950s.
Following Turing's footsteps of defining an adversarially robust discriminator that can separate humans from bots,
we designed CogCAPTCHA30. This goes one level deeper than the Turing Test, from exploring output (what
humans and agents can do) to process (how it can do it). CogCAPTCHA30 combines the original CAPTCHA with 29
classic cognitive psychology tasks for a 30-task battery.
We recruited human participants and also deployed AI agents to perform these tasks. The CAPTCHA experiment
demonstrated that humans and agents can perform at similar performance ( output ) levels, but with different
processes . We then measured output equivalence - how (how similar their answers were)
and process equivalence (how they arrived at their answers) across the whole 30-task paradigm and found that they were uncorrelated:
While the classic Turing test measures whether a machine produces output indistinguishable from a human, we
propose a Process Turing Test measuring whether machines produce a process indistinguishable from humans.
Our results raise two questions: what types of language models - if any - are like humans, and how adversarially
robust is this discrimination process?
To answer the first question, we compared the distance between humans and state-of-the-art frontier models
(OpenAI's GPT, Anthropic's Claude, Google DeepMind's Gemini) as well as Qwen (an open-source 1.5B foundation
model) and Centaur (an open-source 70B-parameter foundation model of human cognition).
We found that state-of-the-art frontier models (Claude, GPT, Gemini) have less similar human process features
compared to smaller models (Qwen, Centaur). As we argued in AI Capability isn't Humanness , while
frontier models are becoming more powerful over time, they are not necessarily becoming more human. Contemporary
progress in artificial intelligence is independent of progress in human simulation.
Qwen, a smaller open-source model, is more humanlike than the larger Claude, GPT, and Gemini. And, as a nice
validation, Centaur outperforms the other models in similarity to human process feature space. We hypothesize this
is due to large-scale output fine-tuning, specifically 10M+ human choices across 160 cognitive experiments.
This introduces the second question: how adversarially robust is the process to discriminate humans from agents?
Any behavioral feature used to distinguish the two may itself become a target for optimization. Accordingly, a
detector that succeeds against off-the-shelf agents establishes a behavioral gap only under the current attacker
model - how AI exists and operates now. It's to be seen whether it can become a durable human-verification signal
for the future technologies. This motivates a stronger test: can an agent close the process gap - between how humans
and agents complete tasks - when given increasingly direct access to human data?
We fine-tuned a Qwen2.5 Instruct model to bring it closer to humans. When given full information - the observed
features and the discriminator's objective function - the gap between humans and agents disappears. However, the gap
reappears when parts of the feature space are left out and fully returns when agents have to generalize
cross-task. In other words, the Process Turing Test is robust when the AI does not have full access to the
discriminator and the feature set (i.e., the model does not know how it will be evaluated).
The challenge the Process Turing Test poses is whether AI can continuously replicate all of human cognitive
psychology. Despite the anxiety that models are becoming more capable over time, they are empirically not becoming
more humanlike . Compared to one-time checks like passwords, CAPTCHAs, document identification, and device
fingerprinting, the Process Turing Test provides a step-up function in human verification. Simulating human
cognitive psychology is an exponentially more challenging task.
Roundtable is a human verification company building the authentication layer for the internet.
Our Proof of Human system invisibly detects AI agents and bots while preserving the user experience for humans.
Learn more at roundtable.ai .
Mayank Agrawal, Milena Rmus, and Mathew Hardy work at Roundtable Technologies Inc., where they are building
Proof of Human, an invisible authentication system for the web.
Previously, they completed PhDs in cognitive science at Princeton University (Mayank and Matt) and the University
of California, Berkeley (Milena).
