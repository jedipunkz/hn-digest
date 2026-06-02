---
source: "https://theaileverageweekly.com/posts/8-concrete-ways-i-use-ai-during-a-normal-engineering-workday-week-4-ro.html"
hn_url: "https://news.ycombinator.com/item?id=48354296"
title: "Concrete Ways I Use AI During a Normal Engineering Workday (Week 4 Roundup)"
article_title: "8 Concrete Ways I Use AI During a Normal Engineering Workday (Week 4 Roundup) — The AI Leverage Weekly"
author: "talvardi7"
captured_at: "2026-06-02T00:43:29Z"
capture_tool: "hn-digest"
hn_id: 48354296
score: 1
comments: 0
posted_at: "2026-06-01T09:00:57Z"
tags:
  - hacker-news
  - translated
---

# Concrete Ways I Use AI During a Normal Engineering Workday (Week 4 Roundup)

- HN: [48354296](https://news.ycombinator.com/item?id=48354296)
- Source: [theaileverageweekly.com](https://theaileverageweekly.com/posts/8-concrete-ways-i-use-ai-during-a-normal-engineering-workday-week-4-ro.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-01T09:00:57Z

## Translation

タイトル: 通常のエンジニアリング業務中に AI を使用する具体的な方法 (第 4 週のまとめ)
記事のタイトル: 通常のエンジニアリング業務中に AI を使用する 8 つの具体的な方法 (第 4 週のまとめ) — The AI Leverage Weekly
説明: AI の生産性に関するアドバイスのほとんどは抽象的なままです。 「AI を使用してより良いコードを作成します。」素晴らしい。どうやって？ここでは、私が実際に使用した、コピー＆ペースト可能な 8 つの具体的なパターンを紹介します。

記事本文:
AI Leverage Weekly エンジニアのための実践的な AI ワークフロー
無料プロンプト 購読する
← すべての投稿
通常のエンジニアリング業務中に AI を使用する 8 つの具体的な方法 (第 4 週のまとめ)
AI の生産性に関するアドバイスのほとんどは抽象的なままです。 「AI を使用してより良いコードを作成します。」素晴らしい。どうやって？ここでは、私が先週実際に使用した、コピー＆ペーストが可能な 8 つの具体的なパターンを紹介します。それぞれが自己完結型で、実際に期限内に成果が得られます。
技術的な同期を行う前に、関連する PR の差分またはチケットの説明を貼り付けて、「この変更における未解決の質問とおそらく争点を要約してください」と尋ねます。 30秒かかります。つまり、私は、誰かがエッジケースについて尋ねたときに、ぼんやりと見つめているような人間ではありません。
2. 強制機能によるラバーダックデバッグ
バグを説明する代わりに、私が予想していることと実際に何が起こっているのかを説明し、モデルに、私が間違っている可能性のあるすべての仮定をリストするように依頼します。そのため、正確に言う必要があり、応答が到着する前にバグが表面化することがよくあります。
3. テスト行列の生成 自分で書くのは省略します
関数を書いた後、それを貼り付けて次のように尋ねます。
この関数が与えられた場合、エッジ ケースを含む、カバーする価値のあるすべてのテスト ケースをリストします。
境界条件と故障モード。番号付きリストとしてフォーマットします。
まだコードはなく、ケースだけです。
全画面モードに入る
全画面モードを終了する
おそらく 2 ～ 3 つのテストを書きます。これで 7 ～ 10 が表面化します。重要なものを選びます。
4. コードレビューのコメントをパターンに変える
2 回以上残したレビュー コメントを残すときは、それをモデルに貼り付けて、「これをチーム用に文書化できる再利用可能なコード レビュー ヒューリスティックに変換してください」と尋ねます。それに関する会議をスケジュールせずに、チーム固有のスタイル ガイドをゆっくりと作成します。
5. 緻密なドキュメントを実用的なサンプルに変換する
ライブラリの技術的なドキュメントにアクセスしたとき

確かに正しいですが、実際には役に立たないので、関連するセクションを貼り付けて質問します。
[特定のこと]を行うために[ライブラリ]を使用しています。関連するドキュメントセクションは次のとおりです。
[貼り付け]
正確に [X] を実行する最小限の動作するコード例を教えてください。実際の値を使用し、
プレースホルダーではありません。ドキュメントに記載されていない問題があれば指摘してください。
全画面モードに入る
全画面モードを終了する
これは、認証フロー、SDK の癖、および非同期動作に関係するものに特に適しています。
6. 差分からの最初のドラフトのコミットメッセージ
ステージングされた diff を貼り付け、1 行の概要と 3 つの箇条書きの本文を含む従来のコミット形式のメッセージを要求します。編集しているのですが、毎回カーソルが点滅しているところから始めるのが面倒です。
7. アーキテクチャに関する決定を書き出す前にプレッシャーテストを行う
設計ドキュメントのセクションを書く前に、私はアプローチを説明し、「懐疑的な上級エンジニアがこれに対して提起する最も強力な反対は何ですか?」と尋ねます。モデルはスチールマンニングが得意です。ドキュメントレビューよりもここで反対意見を得たいと思います。
このパターンは、私が『AI Leverage Playbook: 50 Prompts & Workflows for Engineers』にパッケージ化したものの 1 つですが、以下のバージョンだけでも価値を得るのに十分です。
私は次のようなアーキテクチャの決定を提案しています: [2 ～ 4 文で説明してください]。
運用の複雑さを気にする懐疑的な上級エンジニアの役割を果たします。
保守性と故障モード。重大度順にランク付けして、上位 5 つの反対意見をリストします。
一般的なものではなく、具体的なものにしてください。
全画面モードに入る
全画面モードを終了する
出力により、より明確なドキュメントが得られ、明確に擁護する価値のある仮定が明らかになりました。
8. 一日の終わりの知識の獲得
習慣を築くのに最も時間がかかったのは、エディターを閉じる前に、学んだことや決定したことを要約するのに協力してくれるよう 3 分間モデルに尋ねることです。

関連するコンテキストを貼り付けて使用します。
「この作業に基づいて、次回再導出する必要がないように、書き留める価値のあることを 2 ～ 3 つ挙げてください。」
これらは実行中のメモ ファイルに保存されます。今から 6 か月後、そのファイルはどのチュートリアルよりも価値があります。
これらはどれも判断に代わるものではありません。これらはすべて、判断を適用するコストを削減します。つまり、「これを行うべきだ」という考えと実際にそれを行うことの間の摩擦が減ります。それは引く価値のあるレバーです。
印象に残る柄は派手なものではありません。これらは、すでに行っている作業に適合し、セットアップに 1 分もかからず、次の 1 時間の煩わしさが大幅に軽減されます。
The AI Leverage Weekly では、このようなワークフローを毎週 1 つずつ詳しく解説しています。実践的で、飾り気のない、無料のワークフローです。購読: https://theaileverageweekly.beehiiv.com/subscribe?utm_source=devto&utm_medium=article&utm_campaign=long_w4
次のメールを受信箱で受け取る
エンジニア向けの実践的な AI ワークフロー。週に 1 冊、綿毛はありません。
20 分間の AI 監査で 3 人の上級開発者が見逃したバグをどのように発見したか (第 3 週のまとめ)
AI ツールはあなたの判断力によって決まります - それが重要です
盗む価値のある AI 支援エンジニアリングの 7 つの習慣 (第 2 週のまとめ)

## Original Extract

Most AI productivity advice stays abstract. "Use AI to write better code." Great. How? Here are 8 specific, copy-paste-ready patterns I actually used this…

The AI Leverage Weekly Practical AI workflows for engineers
Free prompts Subscribe
← All posts
8 Concrete Ways I Use AI During a Normal Engineering Workday (Week 4 Roundup)
Most AI productivity advice stays abstract. "Use AI to write better code." Great. How? Here are 8 specific, copy-paste-ready patterns I actually used this past week — each one self-contained, each one with a real return on time.
Before any technical sync, I paste the relevant PR diff or ticket description and ask: "Summarize the open questions and likely points of contention in this change." Takes 30 seconds. Means I'm never the person staring blankly when someone asks about edge cases.
2. Rubber-duck debugging with a forcing function
Instead of describing the bug, I describe what I expect to happen and what's actually happening — then ask the model to list every assumption I'm making that could be wrong. This forces me to be precise, which often surfaces the bug before the response even arrives.
3. Generating the test matrix I'd skip writing myself
After writing a function, I paste it and ask:
Given this function, list every test case worth covering — including edge cases,
boundary conditions, and failure modes. Format as a numbered list.
No code yet, just the cases.
Enter fullscreen mode
Exit fullscreen mode
I probably write 2–3 tests. This surfaces 7–10. I pick the ones that matter.
4. Turning code review comments into patterns
When I leave a review comment I've left more than twice, I paste it into the model and ask: "Turn this into a reusable code review heuristic I can document for my team." Slowly building a team-specific style guide without scheduling a meeting about it.
5. Translating dense documentation into working examples
When I hit library docs that are technically correct but practically useless, I paste the relevant section and ask:
I'm using [library] to do [specific thing]. Here's the relevant docs section:
[paste]
Give me a minimal, working code example that does exactly [X]. Use real values,
not placeholders. Point out any gotchas the docs don't mention.
Enter fullscreen mode
Exit fullscreen mode
This works especially well for authentication flows, SDK quirks, and anything involving async behavior.
6. First-draft commit messages from diffs
I paste the staged diff and ask for a conventional-commit-format message with a one-line summary and a 3-bullet body. I edit it, but starting from something beats starting from a blinking cursor every time.
7. Pressure-testing architecture decisions before you write them up
Before I write a design doc section, I describe the approach and ask: "What are the strongest objections a skeptical senior engineer would raise against this?" The model is good at steelmanning. I'd rather get the objections here than in the doc review.
This pattern is one of the ones I've packaged into The AI Leverage Playbook: 50 Prompts & Workflows for Engineers — but the version below is enough to get value on its own.
I'm proposing the following architecture decision: [describe it in 2-4 sentences].
Play the role of a skeptical senior engineer who cares about operational complexity,
maintainability, and failure modes. List your top 5 objections, ranked by severity.
Be specific, not generic.
Enter fullscreen mode
Exit fullscreen mode
The output gives me a sharper doc and surfaces assumptions worth defending explicitly.
8. End-of-day knowledge capture
The one that took me longest to build into a habit: before I close my editor, I spend 3 minutes asking the model to help me summarize what I learned or decided. I paste any relevant context and use:
"Based on this work, what are 2–3 things worth writing down so I don't have to re-derive them next time?"
These go into a running notes file. Six months from now, that file is worth more than any tutorial.
None of these replace judgment. They all reduce the cost of applying judgment — less friction between "I should do this thing" and actually doing it. That's the lever worth pulling.
The patterns that stick aren't the flashy ones. They're the ones that fit inside work you're already doing, take under a minute to set up, and make the next hour measurably less annoying.
I break down one workflow like this every week in The AI Leverage Weekly — practical, no fluff, free. Subscribe: https://theaileverageweekly.beehiiv.com/subscribe?utm_source=devto&utm_medium=article&utm_campaign=long_w4
Get the next one in your inbox
Practical AI workflows for engineers. One issue a week, no fluff.
How a 20-Minute AI Audit Caught a Bug That 3 Senior Devs Missed (Week 3 Roundup)
Your AI Tools Are Only as Good as Your Judgment — And That's the Point
7 AI-Assisted Engineering Habits Worth Stealing (Week 2 Roundup)
