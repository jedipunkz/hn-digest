---
source: "https://jo-lang.org/"
hn_url: "https://news.ycombinator.com/item?id=48399905"
title: "The Jo Language: Secure Programming for the AI Era"
article_title: "Home | Jo Programming Language"
author: "theanonymousone"
captured_at: "2026-06-04T16:12:23Z"
capture_tool: "hn-digest"
hn_id: 48399905
score: 2
comments: 0
posted_at: "2026-06-04T15:14:13Z"
tags:
  - hacker-news
  - translated
---

# The Jo Language: Secure Programming for the AI Era

- HN: [48399905](https://news.ycombinator.com/item?id=48399905)
- Source: [jo-lang.org](https://jo-lang.org/)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T15:14:13Z

## Translation

タイトル: Jo 言語: AI 時代の安全なプログラミング
記事のタイトル: ホーム |ジョープログラミング言語
説明: AI 時代の安全なプログラミング

記事本文:
コンテンツへスキップ Jo 言語検索 K メイン ナビゲーション 概要 AI セキュリティ ビルド ツール 言語リファレンス ブログの外観
ジョー AI 時代のセキュア プログラミング
AI エージェントには必要な機能のみを提供します。これは、ランタイム サンドボックスではなく型システムによって保証されます。
コードは、宣言して許可されたものにのみアクセスできます。それを超えるものはコンパイル エラーとして拒否されます。
AI エージェントを宣言された機能に限定します。明示的に許可されている範囲を超えて、ネットワーク、ファイルシステム、またはデータにアクセスすることはできません。
コールスタックを通じて機能を暗黙的に渡します。ボイラープレート、グローバル、フレームワークは必要ありません。
ランタイム型エラーはなく、コンパイラは他の言語が見逃したものをキャッチします。
論理演算子を使用して再利用可能なパターン述語を作成します。徹底的なマッチング — コンパイラーは大文字と小文字の区別が見落とされないことを保証します。
1 つの言語、複数のランタイム。すべてのプラットフォームで同じ型を保証して、Python または Ruby にコンパイルします。
Jo のご紹介 — AI 時代の安全なプログラミング · 2026 年 6 月
Jo は TypeScope によって開発および保守されています。

## Original Extract

Secure programming for the AI era

Skip to content Jo Language Search K Main Navigation Overview AI Security Build Tool Language Reference Blog Appearance
Jo Secure Programming for the AI Era
Give AI agents only the capabilities they need — guaranteed by the type system, not a runtime sandbox.
Code can only access what it declares and is granted. Anything beyond that is rejected as compiling errors.
Confine AI agents to declared capabilities. No network, filesystem, or data access beyond what you explicitly allow.
Pass capabilities through the call stack implicitly — no boilerplate, no globals, no frameworks.
No runtime type errors, the compiler catches what other languages let slip through.
Compose reusable pattern predicates with logical operators. Exhaustive matching — the compiler guarantees no case is missed.
One language, multiple runtimes. Compile to Python or Ruby with the same type guarantees on every platform.
Introducing Jo — Secure Programming for the AI Era · June 2026
Jo is developed and maintained by TypeScope .
