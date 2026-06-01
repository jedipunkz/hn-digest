---
source: "https://00f.net/2026/05/31/ai-slop-is-hard-to-fork/"
hn_url: "https://news.ycombinator.com/item?id=48343898"
title: "AI slop is hard to fork"
article_title: "AI slop is hard to fork - Frank DENIS random thoughts."
author: "jedisct1"
captured_at: "2026-06-01T00:34:39Z"
capture_tool: "hn-digest"
hn_id: 48343898
score: 2
comments: 0
posted_at: "2026-05-31T08:05:12Z"
tags:
  - hacker-news
  - translated
---

# AI slop is hard to fork

- HN: [48343898](https://news.ycombinator.com/item?id=48343898)
- Source: [00f.net](https://00f.net/2026/05/31/ai-slop-is-hard-to-fork/)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T08:05:12Z

## Translation

タイトル: AI スロップはフォークするのが難しい
記事のタイトル: AI スロップはフォークするのが難しい - Frank DENIS のランダムな考え。
説明: プロジェクトのフォーク、または着陸するまでに永遠にかかるプル リクエストを保守したことがあるなら、あなたはそこにいたはずです。

記事本文:
メインコンテンツにスキップ
フランク・デニスのランダムな考え。
プロジェクトのフォークや、着陸するまでに永遠にかかるプル リクエストを保守したことがあるなら、あなたはそこにいたはずです。
ある時点で、上流に移動します。次に、 git rebase または git merge を実行すると、突然、分離された素晴らしい変更がマージ競合に覆われてしまいます。
場合によってはこれで対処できることもあります。近くの機能が変更されました。ファイルが移動されました。誰かがタイプの名前を変更しました。十分簡単です。数分かけて問題を修正し、テストを再度実行して、次に進みます。
メンテナはプロジェクト全体を再フォーマットしました。または、すべてのモジュールを再編成しました。または、変更が依存する内部構造を書き換えます。これで、あなたのフォークはもはや実際にはフォークではなくなりました。それは考古学プロジェクトです。
そして、これは常に苦痛でした。
それは単なる退屈な機械作業ではありません。危険を伴う作業でもあります。最初に作成したバージョンは、特定のアップストリーム状態に対してテストされました。その周りのコードは理解できました。おそらく、すべての行に一貫した理由があったでしょう。
3 回の痛みを伴うリベースの後、目標は静かに変化します。
元の変更の品質を維持しようとする必要はもうありません。あなたはただ、このいまいましいものを再度コンパイルしようとしているだけです。あなたはテストをパスさせようとしています。 5 つのわずかに異なるファイルで同じ概念的な矛盾を解決しながら、正気を失わないように努めています。
そして、元々存在しなかったバグや脆弱性が突然導入されてしまいます。リベースのたびにフォークの品質が低下します。
厄介な点は、以前はこれが比較的まれだったということです。
上流での大幅な大幅な変更が発生しましたが、それには実際の人的時間がかかりました。メンテナは、大規模なリファクタリングには苦労する価値があると判断する必要がありました。彼らはその仕事をしなければならなかった。そのため、ほとんどのプロジェクトには自然な摩擦が生じていました。
しかし、AI はその摩擦の多くを取り除きます。
安価な実験は便利です。しかし、それも

o コミットの形状を変更します。
バイブコーディングされたプロジェクトでは、コミットが巨大になることがよくあります。それがパターンになってしまうこともよくあります。プロンプトは、多くの部分に関わる差分を生成します。メンテナは結果を見てテストを実行し、方向性を気に入ってコミットします。単一のコミット、大きな変更。
その差分を作成するコストはわずかでした。他の人がそれに統合するコストはかかりませんでした。
AI によって生成された大規模なリファクタリングは、Enter キーを押す人にとっては安価かもしれませんが、ローカルの変更、ダウンストリームのパッチ セット、または長時間実行されるプル リクエストを維持する人にとっては非常に高価になる可能性があります。
このプロジェクトは単に行動を変えただけではありません。形が変わってしまいました。
フォークは単なるコードのコピーではありません。これは、物がどこに存在するか、機能がどのように分割されるか、どの内部境界が構築に十分安定しているか、どのファイルが他の世界に影響を与えることなく変更できるかについての一連の仮定です。
アップストリームのコミットが 1 つおきにそのシェイプを再シャッフルすると、フォークはアンカーを失います。これは、重要ではあるがすぐにマージできない変更の場合に特に悪影響を及ぼします。
おそらくメンテナはそのアイデアには同意するが、別の API を望んでいるかもしれません。おそらく変更にはさらにテストが必要です。おそらく、これは 1 つのデプロイメントには便利ですが、アップストリームには限定的すぎるかもしれません。もしかしたら、みんな忙しいからプロジェクトのレビューがゆっくり進むのかもしれません。
また、上流のプロジェクトがプロンプトによって常に書き換えられている場合、正常なマージのためのウィンドウははるかに小さくなります。あなたの作品はすぐに着地するか、すぐに腐り始めるかのどちらかです。ロジックが間違っていたからではありません。周囲のコードが別の形にかき混ぜられたからです。
しばらくすると、フォークを維持することが事実上不可能になります。
アップストリームからの更新を停止できます。つまり、フォークは徐々に独自のプロジェクトになります。
リベースを続けることができます。つまり、より多くの費用と費用がかかります。

無関係なグローバル編集によって生じた損傷を修復するのに時間がかかりました。
あるいは、諦めて AI に独自の競合バージョンをゼロから生成するよう依頼することもできます。
最後の選択肢は最悪ですが、まさにインセンティブが指し示すところです。上流側がコードを、気分が変わるたびにグローバルに再生成できる使い捨てテキストとして扱うと、下流側のユーザーも最終的には上流側を同じように扱うようになるでしょう。
安定した形状を維持できないものを、なぜ慎重にフォークし続ける必要があるのでしょうか?
意図的な大きな変更と偶然のチャーンとは違います。
人間のリファクターは通常、何らかの瘢痕組織を持っています。管理者が被害を最小限に抑えようとしているのがわかります。差分には境界があります。コミット メッセージには、なぜこれが発生しなければならなかったのかが説明されています。互換レイヤーが表示されます。古い道はしばらく残ります。査読者は、これが下流ユーザーに損害を与えるかどうかを尋ねます。
AI のスロップは当然、そんなことは気にしません。
現在のチェックアウトのプロンプトを満たすように最適化されます。誰かのフォークにどのパッチ シリーズが存在するかはわかりません。小さな関数の名前変更が 10 件のオープン プル リクエストで競合を引き起こすことは気にしません。他の人に仕事をやり直させることによる社会的コストを感じません。
フォーク可能性はプロジェクトの品質です (だった?)。 AI の登場により、これまで以上に簡単に破壊できるようになりました。
2026 年 5 月 27 日
自分のものではないプル リクエストにプッシュする
2026 年 5 月 17 日
バンの問題は公然と発展している可能性がある
2026 年 4 月 26 日
aHash は PRF ではありません
2026 年 4 月 13 日
Swival は私が本当に欲しかった AI エージェントです
2026 年 4 月 11 日
構成フラグはソフトウェアが腐敗する場所です

## Original Extract

If you’ve ever maintained a fork of a project, or a pull request that takes forever to land, you’ve been there.

Skip to main content
Frank DENIS random thoughts.
If you’ve ever maintained a fork of a project, or a pull request that takes forever to land, you’ve been there.
At some point, upstream moves. Then you run git rebase or git merge , and suddenly your nice isolated change is covered in merge conflicts
Sometimes this is manageable. A nearby function changed. A file moved. Someone renamed a type. Easy enough. You spend a few minutes fixing things, run the tests again, and move on.
The maintainer reformatted the entire project. Or reorganized every module. Or rewrote internal structures your change depends on. Now your fork isn’t really a fork any more. It’s an archaeology project.
And this has always been painful.
It’s not just boring mechanical work. It’s also risky work. The version you originally wrote was tested against a specific upstream state. You understood the code around it. You probably had a coherent reason for every line.
After three painful rebases, the goal quietly changes.
You’re no longer trying to preserve the quality of the original change. You’re just trying to make the damn thing compile again. You’re trying to make tests pass. You’re trying not to lose your mind while resolving the same conceptual conflict in five slightly different files.
And boom, bugs and vulnerabilities that didn’t originally exist get introduced. The quality of the fork degrades after every rebase.
The annoying part is that this used to be relatively rare.
Big, sweeping upstream changes happened, but they cost real human time. A maintainer had to decide that a massive refactor was worth the pain. They had to do the work. So most projects had some natural friction.
But AI removes a lot of that friction.
Cheap experimentation is useful. But it also changes the shape of commits.
In vibe-coded projects, commits often get huge. Often enough that it becomes a pattern. A prompt produces a diff that touches a lot of surface. The maintainer looks at the result, runs the tests, likes the direction, and commits it. Single commit, large changes.
The cost of producing that diff was tiny. The cost of everyone else integrating with it was not.
A large AI-generated refactor may be cheap for the person pressing enter, but it can be extremely expensive for anybody maintaining local changes, a downstream patch set, or a long-running pull request.
The project didn’t just change behavior. It changed shape.
A fork isn’t only a copy of code. It’s a set of assumptions about where things live, how functions are split, which internal boundaries are stable enough to build on, and which files can be changed without touching the rest of the world.
When every other upstream commit reshuffles that shape, the fork loses its anchor. This is especially bad for changes that are important but not immediately mergeable.
Maybe the maintainer agrees with the idea but wants a different API. Maybe the change needs more testing. Maybe it’s useful for one deployment but too specific for upstream. Maybe the project moves slowly on review because everybody is busy.
And if the upstream project is constantly being rewritten by prompts, the window for a sane merge gets much smaller. Either your work lands quickly, or it starts rotting immediately. Not because the logic became wrong. Because the surrounding code was churned into another shape.
After a while, maintaining the fork becomes virtually impossible.
You can stop updating from upstream, which means the fork slowly becomes its own project.
You can keep rebasing, which means spending more and more time repairing damage caused by unrelated global edits.
Or you can give up and ask an AI to generate your own competing version from scratch.
That last option sucks, but it’s exactly where the incentives point. If upstream treats code as disposable text that can be globally regenerated whenever the mood changes, downstream users will eventually treat upstream the same way.
Why maintain a careful fork of something that refuses to keep a stable shape?
There’s a difference between intentional large changes and casual churn.
A human refactor usually carries some scar tissue. You can see the maintainer trying to minimize damage. The diff has boundaries. The commit message explains why this had to happen. Compatibility layers appear. Old paths survive for a while. Reviewers ask whether this will hurt downstream users.
AI slop doesn’t naturally care about any of that.
It optimizes for satisfying the prompt in the current checkout. It doesn’t know which patch series exists in someone’s fork. It doesn’t care that a small function rename creates conflicts in ten open pull requests. It doesn’t feel the social cost of making everybody else redo work.
Forkability is (was?) a project quality. AI makes it easier than ever to destroy it.
May 27, 2026
Pushing to a pull request that isn't yours
May 17, 2026
Bun's problem may be developing in the open
Apr 26, 2026
aHash is not a PRF
Apr 13, 2026
Swival is the AI agent I actually wanted
Apr 11, 2026
Configuration flags are where software goes to rot
