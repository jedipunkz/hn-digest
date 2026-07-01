---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48746066"
title: "Show HN: Ox – an AI agent that catches tech debt before it's committed"
article_title: ""
author: "riggo"
captured_at: "2026-07-01T13:48:05Z"
capture_tool: "hn-digest"
hn_id: 48746066
score: 2
comments: 1
posted_at: "2026-07-01T13:03:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ox – an AI agent that catches tech debt before it's committed

- HN: [48746066](https://news.ycombinator.com/item?id=48746066)
- Score: 2
- Comments: 1
- Posted: 2026-07-01T13:03:38Z

## Translation

タイトル: 番組 HN: Ox – 技術的負債をコミット前に把握する AI エージェント
HN テキスト: こんにちは、HN。私の名前は Craig です。Ox ( https://try-ox.com ) の作成者です。 Ox は、コードベースに導入される前に技術的負債を回避します。 IBM では、私はソフトウェア エンジニアとして、エンジニアがコードベースを理解するのに役立つツールを構築していました。このツールでは、常に静的分析を利用してソースをマッピングし、ベースラインの理解を提供していました。既知の技術的負債を削減することを優先するために経営陣と何年も争った後、私は大規模な全面見直しではなく、クライアントが単純に根本的な進歩を遂げられるように支援するものを作りたいと心から思っていました。介入がなければ、AI が生成したコードの爆発的な増加は、技術的負債の爆発につながるでしょう。この問題は導入されたときに修正するのが最も安価であるため、Ox はコード変更が発生するたびに (CLI REPL セッションを介して git プッシュする直前に、または GitHub でプル リクエストが開かれたときに自動的に) レビューし、作成者またはレビュー担当者が簡単に受け入れることができる具体的な修正案とともに、状況に応じた具体的な調査結果を投稿します。 Ox は、包括的でありながら開発者のスピードを優先し、導入と使用が簡単になるように設計されました。 Python 用に構築されているため、どの言語でも動作するはずです (ただし、「どの言語」もまだテストしていません)。もちろん、最もクリーンなコードが最もパフォーマンスの高いコードを保証するとは限りません (Shipilev の曲線を参照)。 Ox はクリーンさと保守性を向上させるために構築されていますが、学習するためにも構築されています。Ox のエンジンはインストール時にリポジトリに合わせてカスタマイズし、コードベースの規則、ルールセット、例外をすぐに理解して、PR ごとに学習を続けます。試してみたい場合は、ここで汎用のデモ アプリをセットアップしました ( https://github.com/CraigRiggins/taskflow-ox-demo )。楽しみのために、私たちは

最もクリーンなコードと最もダーティなコードをスキャンするミニアプリのオプトイン リーダーボードは、こちら ( https://try-ox.com/score.html ) で確認できます。ご意見やご質問があれば、よろしくお願いします。コメントで私を見つけることができます。対面でのデモをご希望の場合は、craig@try-ox.com までご連絡ください。

## Original Extract

Hello HN, My name is Craig, and I’m the creator of Ox ( https://try-ox.com ). Ox averts technical debt before it’s introduced into your codebase. At IBM, I was a software engineer who built tools to help engineers understand their codebases which always leveraged static analysis to map the source and provide a baseline understanding. After many years of fighting with executives for priority to reduce known tech debt, I really wanted to make something that simply helped the client make simple fundamental progress rather than massive overhauls. Without intervention, the explosion of AI-generated code will lead to an explosion of tech debt. The problem is cheapest to fix when it’s introduced, so Ox reviews code changes as they happen — either just before a git push via a CLI REPL session, or automatically when a pull request is opened on GitHub — and posts specific, contextual findings with concrete suggested fixes for the author or reviewer to simply accept. Ox was designed to be simple to adopt and use, prioritizing developer speed while remaining comprehensive. Built for Python, it should work with any language (but I haven’t tested “any” yet). Of course, the cleanest code does not necessarily guarantee the most performant code (see Shipilev’s curve). Ox is built to improve cleanliness and maintainability, but it’s also built to learn: Ox’s engine customizes itself to the repository on install, understanding your codebase’s conventions, rulesets, and exceptions right out of the box, and it continues to learn with each PR. If you want to give it a run, I set up a generic demo app here ( https://github.com/CraigRiggins/taskflow-ox-demo ). For fun, we created an opt-in leaderboard for a mini-app that scans for the cleanest and dirtiest code you can check out here ( https://try-ox.com/score.html ) I’d deeply appreciate any thoughts or questions; you can find me in the comments. If you'd like to reach out for an in person demo you can reach me at craig@try-ox.com.

