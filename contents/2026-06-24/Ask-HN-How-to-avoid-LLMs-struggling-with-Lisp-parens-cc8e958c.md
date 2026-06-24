---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48657691"
title: "Ask HN: How to avoid LLMs struggling with Lisp parens?"
article_title: ""
author: "chriswarbo"
captured_at: "2026-06-24T11:00:16Z"
capture_tool: "hn-digest"
hn_id: 48657691
score: 2
comments: 1
posted_at: "2026-06-24T10:22:37Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: How to avoid LLMs struggling with Lisp parens?

- HN: [48657691](https://news.ycombinator.com/item?id=48657691)
- Score: 2
- Comments: 1
- Posted: 2026-06-24T10:22:37Z

## Translation

タイトル: HN に聞く: LLM が Lisp の括弧で苦労するのを避けるには?
HN テキスト: LLM は特定の言語 (Python、Bash など) を好むようですが、Lisp (Racket や Emacs Lisp など) にはすべて苦労しているようです。 Claude のさまざまな反復や、DeepSeekV4 などの安価なモデルを試しましたが、パターンは同じでした。いくつかの編集は成功しますが、最終的にいくつかの括弧が少し間違ってしまい、終わりのないループで「手動で」括弧を数えて一致させることで構文エラーを修正しようとするため、狂気のスパイラルに陥ります。これには 2 つの理由でイライラさせられます。まず、LLM は文字を数えるのが苦手であることで有名です (たとえば、「strawberry」の「r」の数)。そのため、文字を生成してカウントするこのアプローチがあまりうまく機能しないのも不思議ではありません。第 2 に、括弧のバランスをとることは、従来の非 LLM アルゴリズムでは簡単です。したがって、（大型で高価なモデルに頼ることなく）完全に回避可能な問題のように感じられます。 Lispy プロジェクトで LLM をうまく使っている人はいますか?もしそうなら、どのようなワークフローやツールなどがうまく機能すると感じましたか?私は、「手動で」数えるのではなく、Emacs の「check-parens」を使用するように彼らを指導しようとしました。しかし、インデントから推論するほうがうまくいくかもしれません。おそらく、ツリーベースの生成/ツールは、そもそもそのような問題の発生を回避するのでしょうか?

## Original Extract

LLMs seem to love certain languages (Python, Bash, etc.), but they all seem to struggle with Lisp (e.g. Racket or Emacs Lisp). I've tried various iterations of Claude, as well as cheaper models like DeepSeekV4, etc. and the pattern is the same: they'll make a few successful edits, but eventually they'll get some parentheses slightly wrong, then spiral into madness as they attempt to fix the syntax errors by counting and matching-up parentheses "manually" in a never-ending loop. This is frustrating for two reasons: Firstly, LLMs are famously bad at counting characters (e.g. the number of "r"s in "strawberry"), so it's no wonder this approach of generating and counting characters doesn't work very well. Secondly, balancing parentheses is trivial for traditional, non-LLM algorithms; so it feels like an entirely avoidable problem (without resorting to larger, more-expensive models). Is anyone using LLMs successfully on Lispy projects? If so, what workflows, tooling, etc. have you found to work well? I've tried guiding them to use Emacs `check-parens` rather than counting "manually"; but maybe inferring from indentation might work better? Perhaps tree-based generation/tools would avoid introducing such problems in the first place?

