---
source: "https://www.adriankrebs.ch/blog/death-of-ai-workflow-builders/"
hn_url: "https://news.ycombinator.com/item?id=49249060"
title: "The death of AI workflow builders"
article_title: "The death of AI workflow builders"
author: "hubraumhugo"
captured_at: "2026-08-10T20:31:20Z"
capture_tool: "hn-digest"
hn_id: 49249060
score: 3
comments: 0
posted_at: "2026-08-10T20:13:03Z"
tags:
  - hacker-news
  - translated
---

# The death of AI workflow builders

- HN: [49249060](https://news.ycombinator.com/item?id=49249060)
- Source: [www.adriankrebs.ch](https://www.adriankrebs.ch/blog/death-of-ai-workflow-builders/)
- Score: 3
- Comments: 0
- Posted: 2026-08-10T20:13:03Z

## Translation

タイトル: AI ワークフロービルダーの死
説明: ビジュアルなノーコード エディターは、AI エージェントの未来のように見えました。それから 1 年も経たないうちに、コーディング エージェントが引き継ぎました。

記事本文:
AIワークフロービルダーの死
AI ワークフロー ビルダーの死についての Adrian Krebs ブログ
ビジュアルなノーコードエディターは、AI エージェントの未来のように見えました。それから 1 年も経たないうちに、コーディング エージェントが引き継ぎました。
OpenAI が 2025 年 10 月に Agent Builder をリリースしたとき、誰もがこれがエージェントの構築方法の未来になると予測しました。
オーケストレーション コードを記述する代わりに、チームはエージェント ノード、ツール、ガードレール、ブランチをキャンバスにドラッグ アンド ドロップして接続します。
同じ頃、Microsoft は Copilot Studio でビジュアル エージェント フローを開始し、Google は Vertex Agent Designer を発表しました。
インターフェースはどれも似たようなもので、通常は人気のある OS のノーコード エディターである React Flow に基づいて構築されていました。
今日まで遡ると、それらはすべて非推奨になるか、完全にシャットダウンされています。
ここ数週間で、AI ワークフロー ビルダーを開発しているいくつかのスタートアップが閉鎖または買収されたことにも気づきました。
当時、これらの AI ワークフロー ビルダーは魅力的に見えました。
ツールの使用とエージェントのループは信頼性が低く、高価でした。無限ループに陥ったワークフローで何十億ものトークンを燃やしたのを覚えています。
より複雑なタスクでは、プロンプトを手動で連鎖させる必要がありました。ハーネスエンジニアリングはまだ存在していませんでした
AI ノードをドラッグ アンド ドロップできるビジュアル キャンバスにより、複雑さが管理可能になりました
OpenAI が Agent Builder をリリースしてから約 1 か月後、Claude Opus 4.5 がリリースされ、Claude Code が軌道に乗り、その後すぐに Codex が続きました。
モデルは、計画を立て、ツールを使用し、長い指示に従うことがより上手になりました。コーディング エージェントは、決定論的なワークフロー コードを自分で作成して維持できるようになり、UI で複雑な視覚的なグラフを構成するよりもはるかに簡単になりました。
したがって、上記の一時的なモデルの制限がなくなると、AI ワークフロー ビルダーの必要性もなくなりました。

。 Flowise の夕日では、「開発者はコーディング エージェントに移行した」と述べられています。
私は、未来は、スクリプトが失敗して適応する必要があるたびに飛び込む堅牢なエージェントを周囲に備えた、決定論的なワークフロー スクリプトにあると考えています。
Kadoa ではこれを使用して、自己修復 Web データ パイプラインを構築しています。Web サイトまたは PDF レイアウトが変更されると、エージェントが抽出または変換コードを調査、修正、テストします。解決できない場合は、人間のオペレーターにエスカレーションされます。
決定論的なワークフロー コードの生成は非常に安価になりましたが、実稼働環境での大規模な実行と維持はそうではありませんでした。証明可能な正しい出力を保証することは依然として大きな課題です。

## Original Extract

The visual no-code editors looked like the future of AI agents. Less than a year later, coding agents took over.

The death of AI workflow builders
Adrian Krebs Blog About The death of AI workflow builders
The visual no-code editors looked like the future of AI agents. Less than a year later, coding agents took over.
When OpenAI launched their Agent Builder in October 2025, everyone predicted it was the future of how we’d build agents.
Instead of writing orchestration code, teams would drag-and-drop agentic nodes, tools, guardrails, and branches onto a canvas and connect them.
Around the same time, Microsoft launched visual agent flows in Copilot Studio and Google announced Vertex Agent Designer.
The interfaces all looked kind of similar and were usually built on the popular OS no-code editor React Flow .
Fast forward to today, all of them are deprecated or shut down completely.
Over the last few weeks, I also noticed that several startups building AI workflow builders have shut down or got acquihired:
At the time, these AI workflow builders looked compelling:
Tool use and agent loops were unreliable and expensive. I remember burning a gazillion tokens on a workflow stuck in an infinite loop.
More complex tasks required manually chaining prompts together. Harness engineering wasn’t a thing yet
A visual canvas where you could drag-and-drop AI nodes made that complexity manageable
About a month after OpenAI launched their Agent Builder, Claude Opus 4.5 got released and Claude Code took off, with Codex following soon after.
Models became better at planning, tool use, and following long instructions. Coding agents became capable of creating and maintaining deterministic workflow code themselves, which became much easier than configuring a complicated visual graph in a UI.
So as the above temporary model limitations disappeared, so did the need for AI workflow builders. Flowise’s sunset mentions that “developers moved on to coding agents”.
I think the future lies in deterministic workflow scripts with a robust agent harness around them that jumps in whenever the script fails and needs to adapt.
We use this at Kadoa to build self-healing web data pipelines: when a website or PDF layout changes, an agent investigates, fixes, and tests the extraction or transformation code. If it can’t figure it out, it escalates to a human operator.
Generating deterministic workflow code got very cheap, running and maintaining it at scale in production did not. Ensuring provably correct output is still a big challenge.
