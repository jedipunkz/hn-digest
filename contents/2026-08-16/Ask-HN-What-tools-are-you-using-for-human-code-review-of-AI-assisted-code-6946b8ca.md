---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49321400"
title: "Ask HN: What tools are you using for human code review of AI-assisted code?"
article_title: ""
author: "dafelst"
captured_at: "2026-08-16T17:11:45Z"
capture_tool: "hn-digest"
hn_id: 49321400
score: 1
comments: 0
posted_at: "2026-08-16T16:20:45Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: What tools are you using for human code review of AI-assisted code?

- HN: [49321400](https://news.ycombinator.com/item?id=49321400)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T16:20:45Z

## Translation

タイトル: HN に聞く: AI 支援コードの人によるコード レビューにはどのようなツールを使用していますか?
HN テキスト: 私たちとその同僚のかなりの割合が現在、信じられないほどの速度でエージェント支援コードを量産しています。その中には実際に優れたコードもあれば、あまり良くないコードもたくさんあります。私個人としては、私たちのプロジェクトの本当の品質の関門は、生成されたコードが人間によってどれだけ徹底的にレビューされ、それが正しいだけでなく、アーキテクチャ的に合理的であることを確認するかどうかであると感じています。 coderabbit や copilot などの AI コード レビュー ツール、あるいは PR でクロード コードを指定するツールは、一般にバグやスタイルの欠陥を見つけるのにはかなり優れていますが、たとえそうするよう求められたとしても、重複コード、モジュールの相互結合、問題の分離などを見つけるのはあまり得意ではありません。 github の PR インターフェイスは私にとってあまり向いていないことに気づきました。レビューが小さいときでも不安定でしたが、今ではこの規模になると管理できなくなりました。これに、エージェントのレビューが混入する余分なノイズや、コピー＆ペーストされたエージェントの出力を「肉プロキシ」する人々が加わると、かなりノイズが多くなり、ナビゲートするのが難しくなります。 AI 支援コードの人によるレビューを合理化するのに効果的なことは何だと思いましたか?ツールやプロセスに関する提案は大歓迎です。

## Original Extract

A good proportion of us and our colleagues are now churning out agent-assisted code at an incredible rate, with some of it that is actually good, and a lot that is not so good. I'm personally finding that the real quality gate for our projects is now how thoroughly the generated code was human reviewed to ensure that it is not just correct, but architecturally sensible. AI code review tools like coderabbit and copilot, or even pointing claude code at a PR are all generally pretty good at finding bugs and style nits, but less good at finding duplicate code, module cross coupling, bad separation of concerns, and so on, even if prompted to do so. I'm finding that github's PR interface is not really cutting it for me, it was janky even when the reviews were small, but now at the size they're at, it is becoming unmanageable. Add to that the extra noise of mixing in agent reviews, and people "meat-proxying" in copy-pasted agent output, and it's getting pretty noisy and difficult to navigate. What have you all found that works well for streamlining human review of AI assisted code? Tools and process suggestions are welcome.

