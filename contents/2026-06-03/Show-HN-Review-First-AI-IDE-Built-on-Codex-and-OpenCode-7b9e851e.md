---
source: "https://handler.team/"
hn_url: "https://news.ycombinator.com/item?id=48367682"
title: "Show HN: Review-First AI IDE, Built on Codex and OpenCode"
article_title: "Don't ship code you can't explain."
author: "vignesh_warar"
captured_at: "2026-06-03T00:50:45Z"
capture_tool: "hn-digest"
hn_id: 48367682
score: 4
comments: 0
posted_at: "2026-06-02T08:58:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Review-First AI IDE, Built on Codex and OpenCode

- HN: [48367682](https://news.ycombinator.com/item?id=48367682)
- Source: [handler.team](https://handler.team/)
- Score: 4
- Comments: 0
- Posted: 2026-06-02T08:58:24Z

## Translation

タイトル: Show HN: Codex と OpenCode に基づいて構築されたレビューファースト AI IDE
記事のタイトル: 説明できないコードを出荷しないでください。
説明: エージェントが提案するすべての編集には、組み込みの説明と独自のチャットが含まれます。理由を尋ねたり、何に触れるのかを尋ねたり、やり直すように指示したりしてください。どれもメインエージェントの邪魔をするものではありません。
HN テキスト: こんにちは、HN、私はソロ開発者の Vignesh です。 Handler は、エージェントがコードを生成している間にレビュー レイヤーを追加する Codex および OpenCode 用の Mac アプリです。すべての編集には、何が変更されたか、変更された理由、変更内容などの短い説明が付いています。 1 つの編集にさらに詳しいコンテキストが必要な場合は、メイン エージェントの実行を中断することなく、その変更についてのみサイド チャットを開いてフォローアップを求めることができます。目標は、コードが生成されるまで待って 1 つの巨大な diff コールドをレビューするのではなく、コードが生成されるときにメンタル モデルを構築することです。これは、積み重ねられた PR をレビューしやすくするのと同じ考え方です。つまり、より小さく、わかりやすい単位になります。ハンドラーは、変更が 1 つの巨大な差分になる前に、AI 生成時にそれを適用します。今年はコードを手書きで書くことをほとんどやめました。私は今、エージェントを運転していますが、その方が速いです。しかし、数か月経って問題に気づきました。エージェントは 400 行の変更を加えて戻ってきて、テストは緑色で、差分は妥当であるように見えましたが、最初の数チャンクの後、私はメンタル モデルをほとんど失いました。その時点では、もう本当にレビューしていませんでした。私は変化の形をざっと読み取って先に進みました。それがエージェントを動かす理想的な方法ではないことは承知しています。しかし、2022 年のように全体の速度を落としてすべての行を手動で検査することも望んでいません。これらのツールを使用する最大のポイントは速度です。そのため、ボトルネックはコードを書くことでなくなりました。コードを早く理解できるようになり、自分のプロジェクトを管理できるようになりました。ハンドラーは、それを自分で修正するために私が構築したものです。 - エージェントが提案するすべての編集が反映されます。

その特定の変更が何を行うのか、そしてその理由についてのわかりやすい説明が付いています。
- 編集には説明が添付されているため、最後に膨大な概要を読むのではなく、段階的にコンテキストを構築できます。
- 単一の編集をサイド チャットに取り込んで、メイン エージェントのコンテキストを混乱させることなく、「なぜこのアプローチをするのですか?」、「これは他に何に影響しますか?」、または「この部分を (すぐに) やり直してください」と尋ねることができます。
- 一枚一枚の差分を見つめて期待するのではなく、編集ごとに編集を進め、承認または拒否します。
- ワークツリーを使用して既存の Codex および OpenCode セットアップ上で実行されるため、並列実行が衝突しません。
- 通常のターミナルよりも Chrome DevTools コンソールに近い組み込みターミナルも備えています。コマンド出力は検査が容易で、エージェントはターミナルの状態を読み取ることができ、別のツールにコピー＆ペーストすることなく構造化データを操作できるネイティブ JSON ビューア/ビルダーもあります。そうでないもの: CodeRabbit や Greptile のようなレビュー ボット。私はそれらを使用していますが、実際のバグを捕らえます。しかし、彼らは事後的にコードをレビューします。ハンドラーはワークフローの前半でポイントします。これは、エージェントが変更を作成している間にレビューするのに役立ち、完成した PR から後からメンタル モデルを再構築するのではなく、メンタル モデルを保持することができます。私はこれを毎日使用していますが、合併エージェントの変更を盲目的に行うことには戻れません。 AI によって生成された大規模な差分で同じ問題を感じたことがある場合は、ぜひ試して、どこで問題が発生するかを教えてください。

記事本文:
説明できないコードを送信しないでください。ハンドラーの価格設定 説明できないコードを出荷しないでください。
エージェントが提案するすべての編集には、組み込みの説明と独自のチャットが含まれます。理由を尋ねたり、何に触れるのかを尋ねたり、やり直すように指示したりしてください。どれもメインエージェントの邪魔をするものではありません。
- 20.8 MB 問題 600 行が一度に着地します。流し読みして、説明できないことを受け入れます。
すべての編集がそれ自体で説明されるようになりました。主要エージェントの話を逸らさずに誰にでも質問してください。
コードベースから AI を排除する機能
発送する商品が到着する前に理解する
すべての変更はそれ自体で説明され、メインのコンテキストを損なうことなく、サイドチャットで変更に質問することができます。
エージェントが端末自体を読み取ります
何が壊れたかを説明するためにログを貼り付ける必要はもうありません。これは JSON を展開し、展開した正確な行を指します。
コンテキストを失わずにモデルを切り替える
画面上のあらゆるものをスクリーンショットし、コマンド バーから直接エージェントに渡します。
こちらは早割価格です。統合メモリ層を出荷すると増加します。つまり、実行するすべてのエージェントで 1 つの共有メモリが使用されるため、最初から開始するものはありません。
- 20.8 MB 今すぐ購入 コードベースに反映される前にすべての変更を確認します
エージェントがメッセージを書く前にアプローチを計画する
編集内容については、メインスレッドとは別のサイドチャットで質問してください
何でもスクリーンショットしてエージェントに直接渡します
エージェントが読み取る端末、JSON ビューアなど
コンテキストを失わずにオープンコードとコーデックスを切り替える
エージェントが衝突しないように分離されたワークツリーをスピンアップする
1 つのクリーンなタブ UI で複数のエージェントを一度に実行
エージェント間の統合メモリ層 (近日提供予定) サブスクリプション

## Original Extract

Every edit the agent proposes arrives with a built-in explanation and its own chat. Ask why, what it touches, or tell it to redo it. None of it derails the main agent.

Hey HN, I’m Vignesh, solo dev. Handler is a Mac app for Codex and OpenCode that adds a review layer while the agent is generating code. Every edit comes with a short explanation: what changed, why it changed, and what it touches. If one edit needs more context, you can open a side chat on just that change and ask follow-ups without derailing the main agent run. The goal is to build a mental model as the code is generated, instead of waiting until the end and reviewing one giant diff cold. It’s the same idea that makes stacked PRs easier to review: smaller, understandable units. Handler applies that at AI generation time, before the change turns into one giant diff. This year I mostly stopped writing code by hand. I drive agents now, and I’m faster for it. But a few months in I noticed a problem. The agent would come back with a 400-line change, tests green, diff looked reasonable, and after the first few chunks I’d pretty much lose the mental model. At that point I wasn’t really reviewing anymore. I was skimming the shape of the change and moving on. I know that’s not the ideal way to drive an agent. But I also don’t want to slow the whole thing down and manually inspect every line like it’s 2022. The whole point of using these tools is speed. So the bottleneck stopped being writing code. It became understanding the code fast enough to stay in control of my own project. Handler is what I built to fix that for myself: - Every edit the agent proposes comes with a plain explanation of what that specific change does and why.
- The explanation is attached to the edit, so you build context incrementally instead of reading one huge summary at the end.
- You can pull any single edit into a side chat and ask “why this approach?”, “what else does this affect?”, or “redo this part (soon)” without blowing up the main agent context.
- You go edit by edit and accept or reject, instead of staring at one wall of diff and hoping.
- It runs on your existing Codex and OpenCode setups, with worktrees so parallel runs don’t collide.
- It also has a built-in terminal that feels closer to a Chrome DevTools console than a plain terminal: command output is easier to inspect, the agent can read the terminal state, and there’s a native JSON viewer/builder for working with structured data without copy-pasting into another tool. What it’s not: a review bot like CodeRabbit or Greptile. I use those, and they catch real bugs. But they review the code after the fact. Handler points earlier in the workflow. It helps you review while the agent is producing the change, so you can keep the mental model instead of trying to reconstruct it later from a finished PR. I’ve been using it daily and can’t go back to merging agent changes blind. If you’ve felt the same problem with large AI-generated diffs, I’d love for you to try it and tell me where it falls over.

Don't ship code you can't explain. Handler Pricing Don't ship code you can't explain.
Every edit the agent proposes arrives with a built-in explanation and its own chat. Ask why, what it touches, or tell it to redo it. None of it derails the main agent.
- 20.8 MB Problem 600 lines land at once. You skim, then accept what you can't explain.
Now every edit explains itself. Question any one without derailing the main agent.
FEATURES THAT KEEP AI SLOP OUT OF YOUR CODEBASE
Understand what you're shipping before it lands
Every change explains itself, and you can question any one in a side chat without polluting the main context.
Your agent reads the terminal itself
No more pasting logs to explain what broke. It expands the JSON and points to the exact row that did.
Switch models without losing context
Screenshot anything on screen and hand it straight to the agent, right from the command bar.
This is the early bird price. It goes up when we ship the unified memory layer: one shared memory across every agent you run, so none of them start from scratch.
- 20.8 MB Buy now Review every change before it lands in your codebase
Plan the approach before the agent writes a line
Question any edit in a side chat, off the main thread
Screenshot anything and hand it straight to the agent
A terminal your agent reads, JSON viewer and all
Switch between opencode and codex without losing context
Spin up isolated worktrees so agents never collide
Run several agents at once in one clean tab UI
Unified memory layer across agents (coming soon) Subscription
