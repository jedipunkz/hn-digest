---
source: "https://carette.xyz/posts/the_five_pillars_of_post_ai_interview/"
hn_url: "https://news.ycombinator.com/item?id=48515383"
title: "The five pillars of the post-AI interview"
article_title: "The five pillars of the post-AI interview | A journey into a wild pointer"
author: "weird_trousers"
captured_at: "2026-06-13T10:03:03Z"
capture_tool: "hn-digest"
hn_id: 48515383
score: 1
comments: 0
posted_at: "2026-06-13T09:44:05Z"
tags:
  - hacker-news
  - translated
---

# The five pillars of the post-AI interview

- HN: [48515383](https://news.ycombinator.com/item?id=48515383)
- Source: [carette.xyz](https://carette.xyz/posts/the_five_pillars_of_post_ai_interview/)
- Score: 1
- Comments: 0
- Posted: 2026-06-13T09:44:05Z

## Translation

タイトル: ポスト AI インタビューの 5 つの柱
記事タイトル: ポスト AI インタビューの 5 つの柱 |ワイルドポインターへの旅
説明: プログラミング テストは、プロファイルを分離するための主要なフィルターでした。
ソフトウェア エンジニアリングでは、歴史的に 3 つの異なるものを区別してきました。
根本的な問題を解決するのはエンジニアです。適切なパターンを適用し、アーキテクチャを調整する開発者は、それぞれに応じて異なります。
[切り捨てられた]

記事本文:
ポスト AI インタビューの 5 つの柱 |ワイルドポインターへの旅
ワイルドポインターへの旅
AI 後の面接の 5 つの柱
プログラミング テストは、プロファイルを分離するための主要なフィルターでした。
ソフトウェア エンジニアリングでは、歴史的に 3 つの異なるものを区別してきました。
エンジニア、つまり根本的な問題を解決する人たちです。
開発者は、ニーズに応じて適切なパターンを適用し、アーキテクチャを調整します。
コードを書く人、コードをひたすら書く人たち。
私たちはかつて候補者に、文章、コミュニケーション、「構文の想起」、そして彼らの考えがチームの文化に適合しているかどうかなど、すべてをチェックするための面倒な手順を課していました。
LLM ブームが起こる前から、ほとんどの場合、ライブ コーディングの面接は「暗記のパフォーマンス」であったため、すでに疑わしいものでした。
しかし、今日ではそれも無意味になってしまいました。
カルチャーフィットは不動のものであり続けます。有毒な天才やコミュニケーションの取れないチームメイトは、使用するツールに関係なく、依然として悲惨な採用となります。そして、すでにご存知かもしれませんが、AI は人間の行動を修正しません。
しかし、コーディング テスト自体が廃止されたのは、コードを書くことが役に立たないからではなく (私は、何かを作ることがそれを理解する最良の方法であると常に考える人です)、テストを不正行為することがこれまでよりも簡単になったからです。私たちは、Builder エージェントが数秒で 500 行のプル リクエスト (ほとんどがスロップ) を生成し、それをスロップの海に直接送信できる時代に入りつつあります。
ワークフローを改善するために、人間の主な仕事が変わりつつあります。私たちの中には、もはや単なるクリエイターではなく、（残念なことに、私の意見では）編集者や監査人になっている人もいます。
時代遅れの儀式を取り去れば、新たな現実が残ります。
この現実を反映するために、私は

個人的には、従来のテストを 5 つの異なるチェック、または柱に置き換えます。
アルゴリズムのコア: 疑似コードのみでアルゴリズムを設計できない場合は、スケーラブルで保守不可能なソフトウェアを作成することになります。さらに悪いことに、LLM にそれを書かせても、根本的に欠陥があることを直感的に認識することさえできなくなります。
アーキテクチャの感覚 : 今日の最も難しい問題は、分離された関数を記述することではなく、データ フローの管理、非同期イベントの処理、厳密な API コントラクトの定義、分散アーキテクチャの混乱の中での生き残りです。これは今でも貴重なスキルです (Anthropic にも所有されていないそうです…)。
コードレビュー: コードをレビューできなければ、AI ツールを指示することはできません。チームを率いることはできません。自分の仕事さえも信頼できなくなります。
コードのデバッグ : LLM が定型文を書くという重労働を行っている場合、人間には最も難しい部分であるデバッグが残されます。現代の優れた面接テストでは、AI が生成した微妙なロジックの欠陥が満載のコードベースが候補者に渡され、バグを見つけて修正するよう求められます。
最終的に、候補者がコードを批評、デバッグ、検証できない場合、その候補者は資産ではなく負債になります。
以前は、3 つのプロファイルすべてを受け入れる余地があり、テストはこれらの特定のスキルを明らかにするように設計されていました。
今日は落とし穴があります。
この傾向は否定できません。LLM は究極のプログラマーです。彼らは、定型文を無限に、瞬時に、そして疲労することなく生成します。私たちはプログラマーを置き去りにし、必然的に完全にエンジニアによって構築される未来に戻りつつあります。
提供元
ヒューゴ
そして
トムフラン/タイプミス

## Original Extract

Programming tests used to be our primary filter to separate profiles.
In software engineering, they historically differentiated three distinct ones:
the engineers, the ones who solve the fundamental problems. the developers, the ones who apply the right patterns and orchestrate the architecture depe
[truncated]

The five pillars of the post-AI interview | A journey into a wild pointer
A journey into a wild pointer
The five pillars of the post-AI interview
Programming tests used to be our primary filter to separate profiles.
In software engineering, they historically differentiated three distinct ones:
the engineers , the ones who solve the fundamental problems.
the developers , the ones who apply the right patterns and orchestrate the architecture depending on the needs.
the coders , the ones who just write code, code, and code.
We used to put candidates through exhausting steps to check everything: their writing, their communication, their “syntax recall”, and if their mind fit the culture of the team.
Before the LLM boom, the live coding interview was already questionable as, most of the time, it was a “performance of memorization”.
However, today, it became pointless.
Culture fit remains the immovable constant. A toxic genius or an uncommunicative teammate is still a catastrophic hire, regardless of the tools they use. And, as you might already know, AI doesn’t fix human behaviours.
But the coding test itself is dead , not because writing code isn’t useful (I will always be the one who considers that making a thing is the best way to understand that thing), but because cheating the test is easier than ever . We are entering an era where a Builder agent can generate a 500-line pull request in a matter of seconds, mostly slop, and submit it directly in an ocean of slops.
In order to improve the workflow, the primary job of the human is shifting. Some of us are no longer just creators, but (unfortunately, in my own opinion) editors and auditors.
If we strip away the obsolete rituals, we are left with a new reality.
In order to reflect this reality, I personally would replace the traditional tests with five distinct checks, or pillars:
the algorithmic core : if you can’t design an algorithm with nothing but pseudo-code, you will write unscalable, unmaintainable software. Worse: you will let an LLM write it for you, and you won’t even get the intuition to realize it is fundamentally flawed.
the architectural feeling : the hardest problems today are not about writing isolated functions but managing data flow, handling asynchronous events, defining strict API contracts, and surviving the chaos of distributed architectures. This is still a valuable skill ( which is not even owned at Anthropic apparently… ).
the code review : if you can’t review code, you can’t direct AI tools. You cannot lead a team. You cannot even trust your own work.
debugging code : if LLMs are doing the heavy lifting of writing the boilerplate, the human is left with the hardest part: debugging. A great modern interview test would hand a candidate a codebase riddled with subtle, AI-generated logic flaws and ask them to find and fix the bugs.
At the end, if a candidate cannot critique, debug, and verify code, they are not an asset but a liability.
Previously, we had room for all three profiles, and our tests were designed to expose those specific skills.
Today, there is a catch.
The trend is undeniable: LLMs are the ultimate coders. They produce boilerplate infinitely, instantly, and without fatigue. We are leaving the coders behind, returning, by necessity, to a future built entirely by engineers.
Powered by
Hugo
and
tomfran/typo
