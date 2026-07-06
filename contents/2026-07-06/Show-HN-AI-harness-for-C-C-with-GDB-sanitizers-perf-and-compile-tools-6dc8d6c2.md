---
source: "https://byteask.ai/"
hn_url: "https://news.ycombinator.com/item?id=48805309"
title: "Show HN: AI harness for C/C++ with GDB, sanitizers, perf and compile tools"
article_title: "ByteAsk. The AI coding agent for C and C++ that verifies its own work"
author: "anirudhak47"
captured_at: "2026-07-06T15:01:12Z"
capture_tool: "hn-digest"
hn_id: 48805309
score: 2
comments: 2
posted_at: "2026-07-06T14:38:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI harness for C/C++ with GDB, sanitizers, perf and compile tools

- HN: [48805309](https://news.ycombinator.com/item?id=48805309)
- Source: [byteask.ai](https://byteask.ai/)
- Score: 2
- Comments: 2
- Posted: 2026-07-06T14:38:49Z

## Translation

タイトル: HN の表示: GDB、サニタイザー、パフォーマンスおよびコンパイル ツールを使用した C/C++ 用 AI ハーネス
記事のタイトル: ByteAsk。自身の動作を検証する C および C++ 用の AI コーディング エージェント
説明: ByteAsk はターミナルからリポジトリを編集し、差分が表示される前にサニタイザー、デバッガー、およびテスト スイートを実行します。 LLVM、GCC、gdb、Valgrind、CMake などの組み込みサポート。
HN テキスト: 皆さん、私のサイド プロジェクトの 1 つを紹介したいと思います。このアイデアは、C/C++ 開発者のワークフロー向けにネイティブに設計された (モデルに依存しない) コーディング ハーネスです。私は C++ 開発者ですが、クロード コードが gdb や perf などの C++ ツールチェーンで適切に動作するとは思いません。
現在のバージョンには、gdb、clang-tidy、cppcheck、サニタイザー、perf、ベンチマーク、コンパイル DB ナビゲーション、Godbolt、シンボル化、バイナリ インスペクション、および逆コンパイルの統合が含まれています。 Anthropic、OpenAI、Gemini、セルフホスト モデルをサポートしています。 VS Code、CLion、emacs、neovim、cursor 用のエディター ワークフローがあります。

記事本文:
ByteAsk。自身の動作を検証する C および C++ 用の AI コーディング エージェント

## Original Extract

ByteAsk edits your repo from the terminal, then drives the sanitizers, the debugger, and your test suite before you see a diff. Built-in support for LLVM, GCC, gdb, Valgrind, CMake, and more.

hey guys, i wanted to show one of my side projects. The idea is a coding harness (independent of models) natively designed for C/C++ developer workflows. I'm a C++ dev and do not find claude code work well with C++ toolchain like gdb and perf.
The current version has integrations for gdb, clang-tidy, cppcheck, sanitizers, perf, benchmarking, compile DB navigation, Godbolt, symbolization, binary inspection, and decompilation. It supports Anthropic, OpenAI, Gemini, and self-hosted models. There are editor workflows for VS Code, CLion, emacs, neovim, and cursor.

ByteAsk. The AI coding agent for C and C++ that verifies its own work
