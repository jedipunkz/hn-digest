---
source: "https://guibai.dev/a/7657392618506764326/"
hn_url: "https://news.ycombinator.com/item?id=48763230"
title: "AI Refactored a 3-Year Codebase in 20 Minutes–and Nearly Torched a Team"
article_title: "AI Refactored a Colleague's Three-Year Codebase in 20 Minutes—and Nearly Torched a Team — Guibai"
author: "Soarez"
captured_at: "2026-07-02T16:00:32Z"
capture_tool: "hn-digest"
hn_id: 48763230
score: 1
comments: 0
posted_at: "2026-07-02T15:38:56Z"
tags:
  - hacker-news
  - translated
---

# AI Refactored a 3-Year Codebase in 20 Minutes–and Nearly Torched a Team

- HN: [48763230](https://news.ycombinator.com/item?id=48763230)
- Source: [guibai.dev](https://guibai.dev/a/7657392618506764326/)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T15:38:56Z

## Translation

タイトル: AI が 3 年分のコードベースを 20 分でリファクタリング – チームはほぼ白熱状態に
記事のタイトル: AI が同僚の 3 年間のコードベースを 20 分でリファクタリングし、チームをほぼ炎上させた - Guibai
説明: ある開発者は、週末をかけて Codex を使用して、無秩序に広がる 3,000 行のフォーム エンジンをリファクタリングしましたが、AI 主導の変更のスピード自体がプロフェッショナルを軽視する行為になっていることがわかりました。

記事本文:
跪拜 Guibai ☾ ← すべての記事 フロントエンド · JavaScript · AI プログラミング AI が同僚の 3 年間のコードベースを 20 分でリファクタリング - チームをほぼ炎上させた
AI ツールは、大規模なコード変更にかかる時間コストを数週間から数分に短縮します。これにより、努力は尊敬に等しいという社会的シグナルも崩壊します。この新しいスピードを中心にコミュニケーション規範を意図的に再構築しないチームは、技術的負債よりも早く対人的負債を蓄積することになります。
Codex を使用した週末の実験により、モノリシックで型指定されていない JavaScript フォーム エンジンが 6 つのクリーンな TypeScript モジュールに変わりました。技術的な結果は客観的に優れていて、明確な境界、完全な型、合格したテスト、そして突然追加するのが簡単になった新機能などです。 But the PR landed like a bomb.オリジナルの作成者は、3 年間の作業が 20 分でツールに置き換えられるのを目の当たりにし、チーム リーダーは本当の失敗、つまり他人のドメインを書き換える変更前のコミュニケーションがゼロであると警告しました。
The core problem wasn't the code quality. AI の速度により、リファクタリングにかかる​​通常のコストは数週間から数時間に短縮され、真剣な取り組みのシグナルが尊重する暗黙のシグナルが消去されました。他の人の仕事を変更することがほぼ自由になると、「私がすべきだ」と「私がやった」の間の心理的な障壁がなくなり、チームを結び付ける人間の契約がほころび始めます。
The resolution required killing the giant PR, splitting the work into small, reviewed chunks, and spending a week on what AI did in minutes.コードは結局同じ場所に落ちましたが、チームは生き残りました。
AI doesn't just accelerate coding; it removes the natural braking mechanism that manual effort provides.リファクタリングに 2 週間かかる場合は、開始する前によく考えてください。 When it takes twenty minutes, impulse wins.
「効率の暴力」という言葉が何かを捉えている

real: 他人の長年の仕事を数分で代替可能に見せるツールには、意図的かどうかにかかわらず、暗黙の判断が伴います。
チーム内でのコードの所有権は、部分的には社会契約です。新しい規範が明示的に構築されない限り、誰でも他人のモジュールを書き換えることができる AI ツールは、その契約を即座に脅かします。
この修正は技術的なものではなく、手順的なものでした。小規模な PR、共同レビュー、および遅いペースでは、同じコード結果が生成されましたが、より困難でより価値のある資産である関係が維持されました。
AI 時代のコラボレーションには、最初に元の作成者とリファクタリングについて話し合うこと、レビュー中の同僚を待ち伏せするために AI 出力を決して使用しないことなど、明確なルールが必要です。これまでは、膨大な労力コストによって強制されていました。
中心的な反対意見は、自分の範囲外のコード、特に AI に触れるのは無謀であり、職業的に侮辱であるということです。オリジナルの作成者だけがメンテナンスの負担を負い、ビジネス ロジックを理解しているため、変更の速さは美点ではなく欠点とみなされます。少数派の意見では、AI がレビューとテストを処理できる可能性があると示唆されていますが、支配的な立場では、一方的なリファクタリングは雇用の安全に対する直接の脅威であり、所有権の侵害であると見なされています。
自分の業務に関係のないコードには触れないでください。そこに地雷を落としたのがあなたであれAIであれ、責任を追及するのは難しい。
まさに、善意がトラブルを引き起こすことになるのです。
この話を話半分に聞いてください。他人が所有する大きなモジュール。部外者がすべてのビジネス ロジックを知ることはできず、ましてやコードのすべての行をレビューしてすべてのテスト ケースに合格することはできません。ばかばかしい。
コードレビューとテストケースの実行を AI に任せるだけですよね?
シャオ・チャンがあなたと問題を抱えている主な理由は、あなたが彼の仕事を脅したかもしれないということです。人の生活を脅かすことは、両親を殺すようなものです...
ネ

ヴァーはあえてそれをもう一度やります。

## Original Extract

A developer used Codex to refactor a sprawling 3,000-line form engine over a weekend, only to discover that the speed of AI-driven change itself became an act of professional disrespect.

跪拜 Guibai ☾ ← All articles Frontend · JavaScript · AI Programming AI Refactored a Colleague's Three-Year Codebase in 20 Minutes—and Nearly Torched a Team
AI tools collapse the time-cost of large-scale code changes from weeks to minutes, which also collapses the social signal that effort equals respect. Teams that don't deliberately rebuild communication norms around this new speed will accumulate interpersonal debt faster than technical debt.
A weekend experiment with Codex turned a monolithic, untyped JavaScript form engine into six clean TypeScript modules. The technical result was objectively better: clear boundaries, full types, passing tests, and new features that were suddenly trivial to add. But the PR landed like a bomb. The original author saw three years of work replaced in twenty minutes by a tool, and the team lead flagged the real failure—zero communication before a change that rewrote someone else's domain.
The core problem wasn't the code quality. AI's speed collapsed the normal cost of refactoring from weeks to hours, erasing the implicit signal that serious effort signals respect. When changing someone else's work becomes nearly free, the psychological barrier between "should I" and "I did" disappears, and the human contract that holds teams together starts to fray.
The resolution required killing the giant PR, splitting the work into small, reviewed chunks, and spending a week on what AI did in minutes. The code ended up in the same place, but the team survived.
AI doesn't just accelerate coding; it removes the natural braking mechanism that manual effort provides. When a refactor takes two weeks, you think twice before starting. When it takes twenty minutes, impulse wins.
The phrase 'efficiency violence' captures something real: a tool that makes someone else's years of work look replaceable in minutes carries an implicit judgement, whether intended or not.
Code ownership in teams is partly a social contract. AI tools that let anyone rewrite anyone else's module instantly threaten that contract unless new norms are explicitly built.
The fix wasn't technical—it was procedural. Small PRs, co-review, and a slower pace produced the same code outcome but preserved the relationship, which is the harder and more valuable asset.
AI-era collaboration needs explicit rules—like discussing refactors with original authors first and never using AI output to ambush a colleague in review—that were previously enforced by sheer effort cost.
The core objection is that touching code outside one's own scope—especially with AI—is reckless and professionally insulting. The speed of the change is seen as a liability, not a virtue, because the original author alone bears the maintenance burden and understands the business logic. A minority view suggests AI could handle the review and testing, but the dominant stance treats unsolicited refactoring as a direct threat to job security and a violation of ownership.
Don't touch code that isn't part of your own business!!! Whether you or AI slip a landmine in there, it's hard to pin the blame.
Exactly, good intentions end up causing trouble.
Take this story with a grain of salt. A large module someone else owns—there's no way an outsider can know all the business logic, let alone review every line of code and pass every test case. Laughable.
Just hand the code review and test-case running to AI, no?
The main reason Xiao Zhang has a problem with you is that you might have threatened his job. Threatening someone's livelihood is like killing their parents...
Never daring to do that again.
