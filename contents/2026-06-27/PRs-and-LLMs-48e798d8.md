---
source: "https://gerdzellweger.com/engineering/2026/06/27/prs-and-llms.html"
hn_url: "https://news.ycombinator.com/item?id=48700392"
title: "PRs and LLMs"
article_title: "PRs and LLMs"
author: "gz09"
captured_at: "2026-06-27T18:22:44Z"
capture_tool: "hn-digest"
hn_id: 48700392
score: 2
comments: 0
posted_at: "2026-06-27T18:10:06Z"
tags:
  - hacker-news
  - translated
---

# PRs and LLMs

- HN: [48700392](https://news.ycombinator.com/item?id=48700392)
- Source: [gerdzellweger.com](https://gerdzellweger.com/engineering/2026/06/27/prs-and-llms.html)
- Score: 2
- Comments: 0
- Posted: 2026-06-27T18:10:06Z

## Translation

タイトル: PR と LLM

記事本文:
PR と LLM
ガズ
ゲルト・ゼルウィガー
ブログ
履歴書
GitHub
リンクトイン
ポスト
Feldera では人材を採用しているので、多くのソフトウェア エンジニアと話をしています。これらの問い合わせは、「ソフトウェア エンジニアリングのメタ」がどこにあるのかを明らかにする傾向があります。多くの場合、それは良い兆候がなく、人々がさまざまな理由で新しい機会を探しているということです。
今回は違います。面接を受けていた上級エンジニアが同じ不満を言い続けました。
これらのエンジニアは現在、コーディング エージェントによって作成された膨大な量の PR をレビューしています。彼らは 3 つの理由でそれに苦労しています。
コードの作成に費やしていた時間のほとんどがレビューに費やされるようになりました。
AI が生成したコードが「十分に優れている」とは認識していません
短期間に大規模な変更を製品にマージすることを心配している
これに対する一般的な解決策があるとは主張しませんが、Feldera で私たちが取り組み始めていることは次のとおりです。
レビューには、コードの作成に費やしていた時間が消費されます
よく引用される回答の 1 つは、人間によるレビューを要求するのをやめることです。これは、正確性の問題が低コストで済む製品では機能する可能性があります。しかし、それが私たちにとって良い解決策であるとは思えません。
したがって、フェルデラにおける私たちのアプローチは、PR の準備ができているという「立証責任」を著者にさらに移すことです。
私たちの古いルールでは、各 PR は既存の単体テストを追加または更新する必要があり、主要な機能には統合テストが含まれている必要がありました。
これではもう十分ではないと思います。
現在実験中のアイデアをいくつか紹介します。
新しいリポジトリのカバレッジ目標を 99% に設定します。これにより、エージェントが作成したコードにテストが強制的に組み込まれ、そうでない場合は受け入れられなくなります。通常、LLM は、llvm-cov などのカバレッジ ツールがギャップを特定すると、不足しているテストを作成できます。
CLAUDE/AGENTS.md に、変更を元に戻すことで新しいテストが失敗することが検証されるというルールを追加します。

e テストは実際の問題を実際に捉えます。
実質的な PR はそれぞれ、自律的なエンドツーエンド テストを実行する必要があります。エージェントは製品をエンドツーエンドで実行します (PR 作成者によって開始されます)。エージェントの仕事は、さまざまな構成、環境、障害ケースで新機能/バグ修正を検証することです。テストの概要は PR の説明に含める必要があります。
エージェントが作成したコードを制約するテストに投資します。クライアントが使用する API には、実行可能な仕様または Oracle が含まれている必要があります。プロパティベースのテストは、仕様と実装を通じて同じ生成された入力を実行します。レビュー担当者は仕様を注意深く読み、実装をざっと読んでください。 PR 作成者がテストを作成/生成します。
実行可能な仕様と実装は互換性があるように設計します。実装をその仕様に置き換えたり、仕様を実装に置き換えたりするのは簡単である必要があります。これにより、個々のコンポーネントのバグを分離しやすくなります。
すべての API とコンポーネントの実行可能仕様を維持するのは非常に労力がかかるため、ほとんどの製品ソフトウェアはこのように構造化されていません。しかし、コードを書くのが無料の世界では、これは素晴らしいアイデアであることがわかります。
エンジニアは AI が生成したコードを信頼していない
LLM コードがどれほど優れているかについてはさまざまな意見があります。経験の浅いエンジニアはこれに驚くかもしれませんが、15 年の経験を持つエンジニアは愕然とするかもしれません。
その不一致の一部は味です。一部は測定可能です。レビューの主観を少なくするために、生成されたコードを次の 5 つのポイントに沿って評価するようにしています。
「適切な」データ構造を使用し、「適切な」アルゴリズムの複雑さを備えています
「適切な」API サーフェスを公開して使用します
実装はクリーン コード原則 (DRY など) を使用して記述されています。
重要な場合: 十分に高速です
(1) ほとんどの場合、誰かが次のことを行う必要があります。

予想される動作を明確に述べ、LLM に十分なコンテキストを提供する代表的なテスト コーパスを提供します。多くのアプリケーションでは、かなり遠くまで到達できます。 (2) と (3) は通常、人間による慎重なレビューが必要ですが、(4) は多くの場合、 CLAUDE.md または AGENTS.md で適切なコーディング規約をエンコードすることで改善できます。
(5) は、私たちのようなエンジニアリング チームが最も多くの時間を費やすと思われる部分です。
コードベースが吸収できるよりも速く変更をマージする
上級エンジニアは、コードベースのどの領域を変更すると安全か、危険か、または完全に危険かを知っています。これは、新人エンジニア (LLM) が持っていない知識です。それは、午前 3 時のページの多さから学びました。
その制度上の記憶は貴重です。 Feldera では、この分野で上級エンジニアが反発し、物事の速度を遅らせることが期待されています。
そうは言っても、「危険」または「完全に危険」の領域に該当するコードが多すぎることは、現在ではさらに大きな責任となっています。非常に複雑なコードを読み取ったり変更したりしても、エージェントの速度が低下することはありません。しかし、フットガンだらけのモジュールを考慮すると、エージェントのパフォーマンスは人間よりもさらに悪いことがよくあります。
そのコードは通常、いくつかのカテゴリのいずれかに分類されます。分割すべき神モジュールに成長したのかもしれません。複雑さが蓄積されている可能性があるため、見直しが必要です。あるいは、仕様が不十分で、テストが不十分で、現在必要な基準を下回っている可能性があります。幸いなことに、AI はコードをその状態に残すコストも変化させます。
これについて強い意見がある場合は、それについて教えてください。 Rust バックエンドおよびプロダクト エンジニアを募集しています。

## Original Extract

PRs and LLMs
gz
Gerd Zellweger
Blog
CV
GitHub
LinkedIn
Post
We are hiring at Feldera, so I have been talking to many software engineers. These calls tend to reveal where the “meta of software engineering” is at. Often the signal is that there isn’t a good signal and people are looking for new opportunities for various reasons.
This time it’s different: Senior engineers who were interviewing kept raising the same complaint.
These engineers now review an unholy amount of PRs written by coding agents. They struggle with it for three reasons.
Reviewing now takes up most of the time they used to spend writing code.
They do not perceive AI-generated code as ‘good enough’
They worry about merging large changes into their product in such a short time
I don’t claim to have a general solution for this, but here is what we are starting to do at Feldera:
Reviewing consumes the time I used to spend writing code
One often quoted response is to stop requiring human reviews. That may work for products where correctness issues have a low cost. But I don’t think it’s a good solution for us.
Hence, our approach at feldera is to shift more of the “burden of proof” that a PR is ready to the author .
Our old rule was that each PR must either add or update existing unit tests, and major features had to include integration tests.
I believe this isn’t enough anymore.
Here are some ideas we are experimenting with at the moment:
Set a 99% coverage target for new repositories. This forces agent-written code to include tests, otherwise it won’t get accepted. LLMs can usually write the missing tests once coverage tooling such as llvm-cov identifies the gaps.
Add a rule to CLAUDE/AGENTS.md that new tests are validated to fail by reverting the change to make sure the tests actually capture a real issue.
Each substantial PR must run autonomous end-to-end testing. An agent runs the product end-to-end (started by the PR author). The agent’s job is to validate the new feature/bugfix in different configurations, environments, and failure cases. The test summary should be included in the PR description.
Invest in tests that constrain agent-written code: APIs used by clients must include an executable specification or oracle. Property-based testing runs the same generated inputs through the spec and the implementation. Reviewers should read the spec carefully and skim the implementation. The PR author writes/generates the tests.
Design executable specifications and implementations so they are interchangeable. It should be easy to replace an implementation with its specification, or a specification with the implementation. This makes it easier to isolate bugs in individual components.
Most production software is not structured this way because maintaining executable specifications for every API and component is very labor-intensive. But in a world where writing code is free it turns out to be a great idea.
The engineer does not trust AI-generated code
There are different opinions about how good LLM code is. One engineer with little experience might be amazed by it, an engineer with 15 years of experience might be appalled.
Some of that disagreement is taste. Some of it is measurable. To make reviews less subjective, we try to evaluate generated code along five points:
It uses ‘the right’ data structures and has ‘the right’ algorithmic complexity
It exposes and uses ‘the right’ API surface
The implementation is written using clean code principles (DRY, etc.)
If it matters: It is fast enough
(1) mostly requires someone to state the expected behavior clearly and provide a representative test corpus that gives an LLM enough context. For many applications that gets you quite far. (2) and (3) usually need a careful human review, (4) can often be improved by encoding the right coding conventions in your CLAUDE.md or AGENTS.md .
(5) is where I expect engineering teams like ours to spend most of their time.
Merging changes faster than the codebase can absorb
A senior engineer knows which areas of the codebase are safe, risky, or outright dangerous to change. This is knowledge a new engineer (or LLM) does not have. It’s learned from too many 3 a.m. pages.
That institutional memory is valuable. At Feldera, this is where senior engineers are expected to push back and slow things down.
That being said, having too much code that falls in the ‘risky’ or ‘outright dangerous’ territory has now become an even bigger liability. Agents are not slowed down by reading or changing very complicated code. But given a module riddled with footguns an agent often performs even worse than a human.
That code usually falls into one of a few categories. It may have grown into a god module that should be split up. It may have accumulated complexity that needs review. Or it may be underspecified, undertested, and below the standard we now need. Luckily, AI also changes the cost of leaving code in that state.
If you have strong opinions about this, tell me about it. We are hiring Rust backend and product engineers .
