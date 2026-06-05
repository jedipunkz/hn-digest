---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48405221"
title: "I'm tired of LLM skill slop, so I built mine with regression tests"
article_title: ""
author: "iliaov"
captured_at: "2026-06-05T00:58:58Z"
capture_tool: "hn-digest"
hn_id: 48405221
score: 3
comments: 0
posted_at: "2026-06-04T22:02:25Z"
tags:
  - hacker-news
  - translated
---

# I'm tired of LLM skill slop, so I built mine with regression tests

- HN: [48405221](https://news.ycombinator.com/item?id=48405221)
- Score: 3
- Comments: 0
- Posted: 2026-06-04T22:02:25Z

## Translation

タイトル: LLM スキルの低下にうんざりしたので、回帰テストを使用して自分のものを構築しました
HN テキスト: 私は最近、Garry Tan の GStack のようなスキルを試して 1 週間過ごしましたが、いくつかの欠陥があることに気付きました (これについては別途投稿します)。これが私の問題です。スキルやプロンプトが適切かどうか (例: GStack の /office-hours) をどうやって知ることができますか?類似したスキル (例: 異なる「詳細調査」スキル) を比較するにはどうすればよいですか?壊れたソフトウェアを見つけるのは (比較的) 簡単です。クラッシュしたり、エラーが表示されたりします。壊れたスキルはそうではありません。完璧に磨かれ、自信に満ちているように聞こえるスキルは、日常的に私を誤解させ、時間を無駄にさせてしまうため、LLM をまったく使用しなければよかったと思うほどです。 AI スキルはソフトウェアであり、回帰テストが付属している必要があります。 LLM チームには大量の即時回帰テストがあります。 LLM ラッパー SaaS 企業は、大量の即時回帰テストを行っています。しかし、オープンソースのスキルに関して言えば、SKILL.md は合理的ですが、出荷時にはテストがありません (たとえば、GStack の /office-hours には、この記事の執筆時点ではテストがありません)。 Garry Tan、私の声が聞こえたら、/office-hours、/plan-ceo-review、/plan-eng-review などの回帰テストの配布を検討してください。回帰テストでは次のことを行う必要があります: 1. スキルが正しく機能することを証明する 2. 正しい使い方と間違った使い方を実証する 3. スキルの価値を証明する 4. スキルのベンチマークを可能にするスコアルーブリックが付属する 5. 最後のテストは、類似したスキルを相互にベンチマークできるため、最も価値があります。それで私はこれを自分でやり始めました。これは進行中の例です。 plan-cmo-review は、この記事の執筆時点では GStack にマーケティング レビューが欠けているため、GStack を補完するスキルです。私はマーケティング担当者ではありません。このスキルを共有するポイントは、回帰セットアップの概要を説明することです。簡単に言うと、私の探求がどのように進んだのかを以下に示します。 - いくつかの製品で GStack を使用し、結果として得られるデジリレーションを実現しました。

gn_document.md は主にマーケティング面で私を失敗に導いていました。 - Claude Opus 4.8 の助けを借りてスキルの失敗を手動で調査し、最終的に正しい解決策を見つけました。 - Claude に plan-cmo-review スキルを構築するように依頼し、それを実行したところ、欠陥のあるソリューション (GStack の出力と同様) に到達しました。 - 私はクロードに、分析してスコアリング ルーブリックを使用した回帰フィクスチャとして追加するための正しい (手動) ソリューションを提供しました。 - クロードは (ブラインド) 回帰を実行しました - 失敗しました。私たちは何度か繰り返し、重要な問題を発見しました。それは、クロードが私のプロンプトを究極の真実として暗黙的に信頼していたことです。クロードは GStack が何をしているのか知っていると信じていました。 GStack は、私が何をしているのか分かっていると信じていました。しかし、私は製品やスタートアップのリサーチをしていました。定義上、「リサーチ」とは自分が何をしているのかわからないときに行うものです。その信頼の連鎖がスキルを破壊したのです。 - 信頼性の問題を修正し、回帰テストに合格しました。さらにいくつか追加しました。彼らは合格した。 - クロードに回帰を複数回実行してもらいました - 亀裂が生じました。クロードはスキルを繰り返した。今では彼らは通り過ぎていきます。 - この方法論にはまだ欠陥があります。さまざまな LLM の実行、モデル間判定、さらに多くの回帰テストを試してみたいと考えています。スキル github.com/remakeai/plan-cmo-review 。 iliaov.substack.com のメモ。

## Original Extract

I've recently tried skills like Garry Tan's GStack, spent a week with it, and realized it has some flaws (I'll post separately about that). Here's my problem: how do I know if a skill or prompt is any good (e.g. GStack's /office-hours)? How do I compare similar skills (e.g. different "deep research" skills)? Spotting broken software is (relatively) easy — it crashes, prints errors. Broken skills don't. Perfectly polished, confident-sounding skills routinely mislead me and waste my time, to the point where I wish I weren't using an LLM at all. AI skills are software — and they should come with regression tests. LLM teams have tons of prompt regression tests. LLM-wrapper SaaS companies have tons of prompt regression tests. But when it comes to open-source skills, SKILL.md reads reasonable, yet ships with zero tests (e.g. GStack's /office-hours has none at the time of writing). Garry Tan, if you hear me — please consider shipping regression tests for your /office-hours, /plan-ceo-review, /plan-eng-review, and so on. Regression tests should: 1. Prove the skill works correctly 2. Demonstrate correct and incorrect usage 3. Prove the skill's value 4. Come with a scoring rubric to allow skill benchmarking 5. The last one is the most valuable, because it lets you benchmark similar skills against each other. So I started doing this myself. Here's a work-in-progress example: plan-cmo-review, a skill to complement GStack since GStack is missing a marketing review at the time of writing. I'm not a marketing guy; the point of sharing this skill is to outline its regression setup. Briefly, here's how my exploration progressed: - I used GStack on a couple of products and realized the resulting design_document.md was leading me to failure, mainly marketing-wise. - I dug into the skill's failures manually with Claude Opus 4.8's help and ended up finding the correct solution. - I asked Claude to build a plan-cmo-review skill, ran it, and it arrived at a flawed solution (similar to GStack's output). - I gave Claude the correct (manual) solution to analyze and add as a regression fixture with a scoring rubric. - Claude ran the (blind) regression — it failed. We iterated several times and found the key problem: Claude was trusting my prompts implicitly as the ultimate truth. Claude believed GStack knew what it was doing. GStack believed I knew what I was doing. But I was doing product/startup research — and by definition, "research" is what you do when you don't know what you're doing. That trust chain is what broke the skills. - We fixed the trust problem and the regression test passed. We added a few more. They passed. - I had Claude run the regressions multiple times — cracks appeared. Claude iterated the skill. Now they pass. - This methodology is still flawed. I'd like to try running different LLMs, cross-model judging, and a lot more regression tests. Skill github.com/remakeai/plan-cmo-review . Notes at iliaov.substack.com .

