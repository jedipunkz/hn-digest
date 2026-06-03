---
source: "https://firefox-source-docs.mozilla.org/contributing/ai-coding.html#firefox-ai-coding-policy"
hn_url: "https://news.ycombinator.com/item?id=48377056"
title: "Firefox AI Coding Policy"
article_title: "Firefox AI Coding Policy — Firefox Source Docs documentation"
author: "Topfi"
captured_at: "2026-06-03T00:35:15Z"
capture_tool: "hn-digest"
hn_id: 48377056
score: 2
comments: 1
posted_at: "2026-06-02T22:13:26Z"
tags:
  - hacker-news
  - translated
---

# Firefox AI Coding Policy

- HN: [48377056](https://news.ycombinator.com/item?id=48377056)
- Source: [firefox-source-docs.mozilla.org](https://firefox-source-docs.mozilla.org/contributing/ai-coding.html#firefox-ai-coding-policy)
- Score: 2
- Comments: 1
- Posted: 2026-06-02T22:13:26Z

## Translation

タイトル: Firefox AI コーディング ポリシー
記事のタイトル: Firefox AI コーディング ポリシー — Firefox ソース ドキュメント ドキュメント

記事本文:
Firefox ソースドキュメント
クイック検索
Mozilla アプリケーションのクイック ガイド
Firefox コードベースで作業するためのセットアップを行う
Firefox での作業
Firefox 貢献者向けクイック リファレンス
パッチのスタックの操作クイックリファレンス
ポケット ガイド: Firefox の出荷
Firefox AI コーディング ポリシー
哲学
責任
AIツールを利用する場合
Pernosco を使用した Firefox のデバッグ
Valgrind を使用した Firefox のデバッグ
TreeHerder スタックをローカルでシンボリック化する
Windowsタスクマネージャーでプロセスダンプを取得する方法
バグレポートのスタックトレースを取得する方法
WinDbg でスタックトレースを取得する方法
開発証明書の構成
Firefox のソースコードのディレクトリ構造
サポートされているビルド ホストとターゲット
Firefox DevTools 貢献者ドキュメント
データベース バインディング (SQLite、KV など)
ネットワーク セキュリティ サービス (NSS)
サードパーティコンポーネントの販売
Treeherder の結果を理解する
断続的な障害の回避
新しい構成の Firefox テストを有効にする
断続的なテストの失敗のデバッグ
Firefox 開発用 AI エージェント ツール
ローカリゼーションと国際化
サードパーティの Python パッケージの使用
問題を報告する
/
ページのソースを表示
このポリシーは、Firefox プロジェクトに貢献するすべての人に適用されます。これは、ブラウザー用の実稼働コードを作成する際の AI または AI 関連ツールの使用に対する期待と要件を定義します。これらのガイドラインは、Mozilla コミュニティが経験を積むにつれて進化する可能性があります。
AI 支援ツールは急速に進化していますが、開発者がコードを探索、理解、作成するために使用すると、機会とリスクの両方が生じることは明らかです。これらは作業を加速し、問題を早期に発見するのに役立ちますが、正しくなくても正しく見えるコードを生成することも簡単にします。
これらのツールを使用することを選択した寄稿者に対しては、私たちがこれまで彼らを信頼してきたのと同じように、慎重に使用することを信頼しています。

Firefox の開発の他の部分。 AI は支援できますが、責任は常に変化の背後にある人間にあります。
AI 支援による貢献は、他のパッチと同じ正確性、セキュリティ、保守性の基準を満たしている必要があります。
品質と範囲を維持します。使用したツールに関係なく、自分の仕事の技術的な優秀さに対して責任を負います。パッチを小さく、焦点を絞ったものにして、レビューしやすく、正当性を明確に示します。
提出する内容を理解してください。行うすべての変更を理解し、説明できることが期待されます。レビュー担当者の役割は、ツールの出力ではなく、人間の作業を再確認することです。
提出する前にセルフレビューしてください。パッチを完全に理解するには、ピアレビューをリクエストする前に、自分でコードを完全にレビューする必要があります。これに加えて、AI レビューをローカルで実行して、見逃している可能性のある問題を見つけることを検討してください。
機密データを保護します。外部 AI ツールへのプロンプトには、個人情報、セキュリティ上重要な情報、またはその他の機密情報を含めないでください。
結局のところ、使用するツールに関係なく、送信したすべての変更に対して責任を負います。
「Good First」および「Good Next」のバグは、新しい寄稿者が Firefox コードベースを学習するのに役立つように存在します。これらは、学習して貢献を続けるつもりの場合にのみお読みください。これらの多くを AI ツールだけで修正することは、その目標に対して逆効果です。

## Original Extract

Firefox Source Docs
Quick search
A Quick Guide to Mozilla Applications
Getting Set Up To Work On The Firefox Codebase
Working on Firefox
Firefox Contributors’ Quick Reference
Working with stack of patches Quick Reference
Pocket Guide: Shipping Firefox
Firefox AI Coding Policy
Philosophy
Responsibilities
When using AI tools
Debugging Firefox with Pernosco
Debugging Firefox with Valgrind
Symbolicating TreeHerder stacks locally
How to get a process dump with Windows Task Manager
How to get a stacktrace for a bug report
How to get a stacktrace with WinDbg
Configure Development Certificate
Firefox Source Code Directory Structure
Supported Build Hosts and Targets
Firefox DevTools Contributor Docs
Database bindings (SQLite, KV, …)
Network Security Services (NSS)
Vendoring Third Party Components
Understanding Treeherder Results
Sheriffed intermittent failures
Turning on Firefox tests for a new configuration
Debugging Intermittent Test Failures
AI Agent Tools for Firefox Development
Localization & Internationalization
Using third-party Python packages
Report an issue
/
View page source
This policy applies to anyone contributing to the Firefox project. It defines expectations and requirements for any usage of AI or AI-adjacent tools when authoring production code for the browser. These guidelines may evolve as the Mozilla community gains experience.
AI-assisted tools are evolving quickly, but it’s clear that they present both opportunities and risks when used by developers to explore, understand, and create code. They can accelerate work and help catch issues earlier, but they also make it easy to generate code that looks right without being right.
For contributors who choose to use these tools, we trust them to do so with care, just as we trust them with every other part of Firefox’s development. AI can assist, but responsibility always stays with the human behind the change.
AI-assisted contributions must meet the same standards of correctness, security, and maintainability as any other patch.
Maintain quality and scope. You are responsible for the technical excellence of your work, regardless of the tools you used. Keep patches small and focused so they’re easy to review and clearly justified.
Understand what you submit. You’re expected to understand and be able to explain every change you make. The role of the reviewer is to double-check the work of a human, not the output of a tool.
Self-review before submitting. To fully understand the patch you must fully review the code yourself, before requesting peer review. In addition to this, consider running an AI review locally to catch issues you may have missed.
Protect sensitive data. Do not include private, security-sensitive, or otherwise confidential information in prompts to external AI tools.
At the end of the day you are accountable for all changes you submit, regardless of the tools you use.
“Good First” and “Good Next” bugs exist to help new contributors learn the Firefox codebase. Take these only if you intend to learn and continue contributing. Fixing a large number of these with AI tools alone is counterproductive to that goal.
