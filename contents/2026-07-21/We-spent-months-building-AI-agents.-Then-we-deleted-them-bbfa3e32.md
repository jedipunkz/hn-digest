---
source: "https://runnit.io/articles/we-spent-months-building-ai-agents-into-runnit-then-we-deleted-them"
hn_url: "https://news.ycombinator.com/item?id=48991202"
title: "We spent months building AI agents. Then we deleted them"
article_title: "We spent months building AI agents. Then we deleted them. · Runnit"
author: "marktolson"
captured_at: "2026-07-21T12:19:48Z"
capture_tool: "hn-digest"
hn_id: 48991202
score: 2
comments: 0
posted_at: "2026-07-21T12:10:05Z"
tags:
  - hacker-news
  - translated
---

# We spent months building AI agents. Then we deleted them

- HN: [48991202](https://news.ycombinator.com/item?id=48991202)
- Source: [runnit.io](https://runnit.io/articles/we-spent-months-building-ai-agents-into-runnit-then-we-deleted-them)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T12:10:05Z

## Translation

タイトル: AI エージェントの構築に数か月を費やしました。それからそれらを削除しました
記事のタイトル: AI エージェントの構築に数か月を費やしました。その後、それらを削除しました。 · ランニット
説明: 私たちは何か月もかけて、特化した AI エージェントを Runnit に構築しました。その後、それらのほとんどを削除しました。それは、それらが機能しなかったからではなく、モデルが十分に進化して、アーキテクチャがもはや意味をなさなくなったからです。

記事本文:
ランニット
仕組み
なぜランニットなのか
リソース
ドキュメント
早期アクセスをリクエストする
← すべてのリソース
記事
私たちは AI エージェントの構築に数か月を費やしました。その後、それらを削除しました。
過去数か月間、私たちは AI エージェントを Runnit に組み込むことに多くの時間を費やしてきました。
1 つや 2 つではありません。たくさんあります。
それぞれに特定の責任がありました。 1 人は計画に、もう 1 人はリサーチに、もう 1 人はスケジュール管理に、そしてもう 1 人は執筆に焦点を当てていました。彼らは皆、注意深く作成されたプロンプト、独特の動作、そして独自の専門分野を持っていました。
当時はそれが正しいアーキテクチャだと感じました。
私たちが構築していたオリジナルの LLM (つい昨日のことのようでした) のコンテキスト ウィンドウ (64k ～ 128k) ははるかに小さかったため、作業を特殊なエージェントに分割することはクリーンなだけでなく、半正確な結果を生成する必要もありました。各エージェントは 1 つの問題を理解するだけで済みます。これは、利用可能なコンテキストをより効果的に使用できることを意味します。
しかし、モデルが改善されるにつれて、私たちはまだ昨日の制限を考慮して設計していることに徐々に気づきました。
より大きなコンテキストで新しいモデルをテストすればするほど、完全に異なるアシスタントになる必要がなく、異なる種類の作業を自然に切り替えることができることがわかりました。適切な情報を適切なタイミングで入手できれば、同じ会話の中で計画、執筆、調査、推論を行うことができます。
そこで私たちは単純な質問をせざるを得ませんでした。
これらは本当に個別のエージェントである必要があるのでしょうか?
結局のところ、私たちの答えはノーでした。そこで、それらを削除しました。
増え続ける専門エージェントのコレクションを維持する代わりに、必要なときに機能をロードする単一のインテリジェンス ( Ru の導入) を中心にアーキテクチャを再構築しました。
それは能力と指示です。これらは、タスクのためにロードしたり、再度アンロードしたりできる専門知識の一部です。

ja 仕事は終わりました。
この変更は小さいように思えますが、ほぼすべてが簡素化されました。
あなたのビジネスについての一貫した理解が 1 つあります。会話がひとつ。一つの思い出。組織の知識が 1 か所に存在します。
速度が重要な場合は、同じインテリジェンスがそれ自身の並列インスタンスを作成するだけでよく、各インスタンスは、すべてを元に戻す前に、独自の作業に必要なディレクティブをロードします。外側から見ると、まだ並行して動作していますが、その下では同じインテリジェンスが問題のさまざまな部分を解決しています。
また、驚くほど多くの複雑さを削除していることにも気づきました。数十の個別のプロンプトを管理したり、別々のエージェント間の引き継ぎを調整したり、異なるエージェントが進化するにつれて徐々に離れていくことを心配したりする必要はなくなりました。
これは、専門エージェントが時代遅れであることを意味するものではありません。作業を分離したり、別のモデルを使用したり、権限を分けたりする十分な理由はまだあります。
しかし、私たちの多くは、もう存在しないモデルに対して行われたアーキテクチャ上の決定を引き継いでいると思います。
もし私たちが今日再び Runnit を始めるとしたら、「どのエージェントを構築すべきか?」という問いから始めることはないと思います。
別の質問から始めます。
「単一のインテリジェンスにはどのような能力が必要ですか?」
場合によっては、別のレイヤーを追加することが進歩ではない場合があります。
もう必要ないことに気づくこともあります。
Runnit の AI オペレーター Ru の紹介
Ru は Runnit 内のオペレーショナル インテリジェンスであり、組織、クライアント、および実際の仕事の進め方を理解するために構築されています。ニーズの変化に応じて、質問に答え、行動を起こし、チームを調整し、作業を自動化し、システムを拡張できます。
Runnit は 1 つの質問から始まりました。自動化によって代理店に何ができるでしょうか?これはcreatのAIネイティブ作業管理プラットフォームとなるプロトタイプの物語です

私たちのチーム。
Runnit が Web サイト訪問者、早期アクセス申請者、製品ユーザーの個人情報をどのように収集、使用、保存、保護するか。
Runnit は、オーストラリアが所有および運営する会社である Null Theory によって ❤️ を使用して構築されています。

## Original Extract

We spent months building specialised AI agents into Runnit. Then we deleted most of them. Not because they didn't work, but because the models had evolved enough that the architecture no longer made sense.

RUNNIT
How it works
Why Runnit
Resources
Documentation
Request early access
← All resources
Article
We spent months building AI agents. Then we deleted them.
Over the last few months, we've spent a lot of time building AI agents into Runnit.
Not one or two of them. Lots of them.
Each had a specific responsibility. One focused on planning, another on research, another on scheduling, another on writing. They all had carefully crafted prompts, distinct behaviours and their own area of expertise.
At the time, it felt like the right architecture.
The original LLMs we were building on (which seemed like only yesterday) had much smaller context windows (64k-128k), so breaking work into specialised agents wasn't just cleaner, it was necessary to produce even semi-accurate results. Each agent only needed to understand one problem, which meant it could use its available context more effectively.
As the models improved though, we slowly realised we were still designing around yesterday's limitations.
The more we tested newer models with larger contexts, the more we found they could naturally switch between different kinds of work without needing to become a completely different assistant. They could plan, write, research and reason within the same conversation, provided they had the right information available at the right time.
That forced us to ask a simple question.
Do these really need to be separate agents?
In the end, our answer was no. So we removed them.
Instead of maintaining a growing collection of specialised agents, we rebuilt the architecture around a single intelligence ( introducing Ru ) that loads capabilities when it needs them.
They're capabilities and directives. They're pieces of expertise that can be loaded for a task and then unloaded again when the work is finished.
The change sounds small, but it simplified almost everything.
There's one consistent understanding of your business. One conversation. One memory. One place for organisational knowledge to live.
When speed matters, that same intelligence can simply create parallel instances of itself, each loading the directives required for its own piece of work before bringing everything back together. From the outside it still works in parallel, but underneath it's the same intelligence solving different parts of the problem.
We also found ourselves deleting a surprising amount of complexity. We no longer had to maintain dozens of individual prompts, coordinate hand-offs between separate agents, or worry about different agents gradually drifting apart as they evolved.
None of this means specialised agents are obsolete. There are still good reasons to isolate work, use different models or separate permissions.
But I do think many of us are carrying architectural decisions that were made for models that no longer exist.
If we were starting Runnit again today, I don't think we'd begin by asking, "What agents should we build?"
I'd start with a different question.
"What capabilities does a single intelligence need?"
Sometimes progress isn't adding another layer.
Sometimes it's realising you don't need it anymore.
Introducing Ru - Runnit's ai operator
Ru is the operational intelligence inside Runnit, built to understand your organisation, your clients and the way work actually gets done. It can answer questions, take action, coordinate teams, automate work and extend the system as your needs change.
Runnit started with one question: what could automation do for an agency? This is the story of the prototype that became an AI-native work management platform for creative teams.
How Runnit collects, uses, stores and protects personal information for website visitors, early access applicants and product users.
Runnit is built with ❤️ by Null Theory , an Australian owned and operated company.
