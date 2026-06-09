---
source: "https://1ps0.info/storytime/"
hn_url: "https://news.ycombinator.com/item?id=48457980"
title: "Show HN: Storytime – Continuity for Claude Code (and other ideas)"
article_title: "Storytime — Continuity System for Claude Code"
author: "oriel"
captured_at: "2026-06-09T10:19:15Z"
capture_tool: "hn-digest"
hn_id: 48457980
score: 1
comments: 0
posted_at: "2026-06-09T07:50:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Storytime – Continuity for Claude Code (and other ideas)

- HN: [48457980](https://news.ycombinator.com/item?id=48457980)
- Source: [1ps0.info](https://1ps0.info/storytime/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T07:50:59Z

## Translation

タイトル: Show HN: Storytime – Continuity for Claude Code (および他のアイデア)
記事のタイトル: Storytime — クロード コードの継続システム
HN テキスト: LLM ハーネス (クロード コードを含む) は急速に進歩しているため、すべての主張の検証を待つよりもこれを公開したほうが良いと考えました。たくさんのアイデアをここに詰め込みました！ HTML にはいくつかの核となるアイデアが概説されていますが、最初のアイデアはスレッド内のペルソナとしてのドメイン レンズでした。メイン パスの /storytime cmd はプロセスを開始し、完全なセットの増分ドキュメントを生成します。これは私にとっても非常に役立ちました。ブレークアウト、@roles などの他のアイデアは、エージェント チームへの最近の動きに組み込まれていますが、このフレーバーはそのアプローチと形式において独特だと思います。残念ながら、他のアイデアの多くは CC の独自のシステムやより構築されたシステムと衝突しており、それらが製品化に値するかどうかを判断するのは困難です。とはいえ、私はこれを毎日のドライバーとしておそらく 4 か月間使用しているので、できれば他の人にとって役に立ったり、インスピレーションを与えたりすることができます。

記事本文:
ストーリータイム v1.0.1
ホーム
ガイド
ウォークスルー
参考資料
ストーリータイム
LLM とハーネスのコラボレーションのための継続システム。
構造化されたペルソナの会話を通じて技術仕様を構築します —
そして、圧縮、セッション、時間にわたってセッション状態を保持します。
問題について説明します。 Storytime はコードベースを調査し、チームを編成します。
ドメインの専門家であるペルソナを配置し、構造化された会話を実行して、
計画 — 実際のコードに基づいて、引用、決定事項、
視覚補助。仕様ワークフローの下には統合ループがあります
これにより、圧縮とセッション間の連続性が維持されます。
空の場合、フェーズは崩壊します。すべての走行ですべてのギアが使用されるわけではありません。
完全なチュートリアルを見る →
/storytime "私たちのパブリック API にはレート制限がないため、悪用を受けています"
Storytime はコードを調査し、Express ミドルウェア チェーンに
スロットル層。 @owner [アンカー]、@systems [格子]、をアセンブルします。
@批評家[鍛冶]、@オペレーター[潮流]。彼らは以下を調査し、生成します。
spec/.storytime/sessions/rate-limiting/001/
§── Survey.md コードベースコンテキスト + フィンガープリント
§──team.md ペルソナ定義
§──icebreaker.md 現状議論
§── Breakout-algo.md トークンバケットとスライディングウィンドウ
§── Breakout-store.md カウンター用の Redis とインメモリの比較
━── plan.md ASCII スライドデッキ + 実装手順
主要な概念
「API にレート制限が必要です」
「このサービスは分割したほうがいいでしょうか？」
「Webhook が静かに失敗します」
「認証ミドルウェアの書き換えが必要です」
「このシステムはどのようにしてここにたどり着いたのでしょうか?」
「その計画は私たちが構築したものと一致しましたか？」
ストーリータイム v1.0.1 — github.com/1ps0/storytime
Alex Evers による Claude Code プラグイン

## Original Extract

Since LLM harness (claude code included) are moving fast, I figured it would be better to put this out than wait to validate each and every claim. I crammed a lot of ideas in here! Theres a few core ideas outlined on the html, but the incepting one was domain lenses as personas within a thread. The main path /storytime cmd kicks off the process and yields a full set incremental documents, and thats been very useful for me as well. Other ideas like breakouts, @roles, and others have been subsumed by recent moves towards agentic teams, but I think this flavor is unique in its approach and format. Unfortunately many of the other ideas have collided with CC's opinionated/more built-out systems, and its hard to say if they're production worthy. That said, I've been using this as my daily driver for maybe 4 months now, so hopefully its useful or inspiring to someone else.

storytime v1.0.1
Home
Guide
Walkthrough
Reference
Storytime
A continuity system for LLM–harness collaboration.
Builds technical specifications through structured persona conversations —
and carries session state across compactions, sessions, and time.
Describe a problem. Storytime surveys your codebase, assembles a team of
domain-expert personas, and runs a structured conversation that produces a
plan — grounded in your actual code, with citations, decisions, and
visual aids. Underneath the spec workflow is a consolidation loop
that preserves continuity across compactions and sessions.
Phases collapse when empty. Not every run uses every gear.
See a full walkthrough →
/storytime "our public API has no rate limiting and we're getting abuse"
Storytime surveys your code, finds the Express middleware chain has no
throttling layer. Assembles @owner [anchor], @systems [lattice],
@critic [forge], @operator [tide]. They investigate, and produce:
specs/.storytime/sessions/rate-limiting/001/
├── survey.md Codebase context + fingerprint
├── team.md Persona definitions
├── icebreaker.md Status quo discussion
├── breakout-algo.md Token bucket vs sliding window
├── breakout-store.md Redis vs in-memory for counters
└── plan.md ASCII slide deck + implementation steps
Key Concepts
"We need rate limiting on the API"
"Should we split this service?"
"Webhooks are silently failing"
"The auth middleware needs a rewrite"
"How did this system get here?"
"Did the plan match what we built?"
Storytime v1.0.1 — github.com/1ps0/storytime
A Claude Code plugin by Alex Evers
