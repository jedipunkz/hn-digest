---
source: "https://bun.com/docs/runtime/templating/init"
hn_url: "https://news.ycombinator.com/item?id=49089156"
title: "`bun init` automatically creates a Claude.md file by default"
article_title: "bun init - Bun"
author: "achristmascarl"
captured_at: "2026-07-28T21:00:31Z"
capture_tool: "hn-digest"
hn_id: 49089156
score: 3
comments: 1
posted_at: "2026-07-28T20:02:27Z"
tags:
  - hacker-news
  - translated
---

# `bun init` automatically creates a Claude.md file by default

- HN: [49089156](https://news.ycombinator.com/item?id=49089156)
- Source: [bun.com](https://bun.com/docs/runtime/templating/init)
- Score: 3
- Comments: 1
- Posted: 2026-07-28T20:02:27Z

## Translation

タイトル: `bun init` はデフォルトで Claude.md ファイルを自動的に作成します
記事タイトル: bun init - Bun
説明: インタラクティブな bun init コマンドを使用して、空の Bun プロジェクトをスキャフォールディングします。

記事本文:
bun init - Bun ドキュメントインデックス
/docs/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
検索... ナビゲーション はじめに bun init ランタイム パッケージ マネージャー バンドラー テスト ランナー ガイド リファレンス ブログ フィードバック はじめに
CLI 使用法の初期化オプション
ページをコピーする ページをコピーする インタラクティブな bun init コマンドを使用して空の Bun プロジェクトをスキャフォールディングする
ページをコピー ページをコピー bun init を使用して新しい Bun プロジェクトをスキャフォールディングします。
ターミナル bun init my-app
?プロジェクト テンプレートを選択します。Return キーを押して送信します。
❯空白
反応する
図書館
✓ プロジェクト テンプレートを選択します: 空白
+ .gitignore
+ クロード.md
+ .cursor/rules/use-bun-instead-of-node-vite-npm-pnpm.mdc -> CLAUDE.md
+index.ts
+ tsconfig.json (エディターのオートコンプリート用)
+ README.md
Enter キーを押して各プロンプトのデフォルトの回答を受け入れるか、-y フラグを渡してデフォルトを自動的に受け入れます。
bun init は適切なデフォルトで設定を推論し、複数回実行しても非破壊的です。
それは以下を作成します:
現在のディレクトリ名がデフォルトの名前を持つ package.json ファイル
エントリ ポイントが TypeScript ファイルであるかどうかに応じて、tsconfig.json または jsconfig.json ファイル
エントリ ポイント。index.{tsx, jsx, js, mts, mjs} のいずれかが存在するか、package.json がモジュールまたはメイン フィールドを指定しない限り、デフォルトは Index.ts になります。
Claude CLI が検出された場合の CLAUDE.md ファイル (CLAUDE_CODE_AGENT_RULE_DISABLED 環境変数で無効化)
Cursor が検出された場合の .cursor/rules/*.mdc ファイル。これにより、Cursor AI に Node.js と npm の代わりに Bun を使用するよう指示されます。
bun init <フォルダー ?>
初期化オプション
--yes boolean 質問せずにデフォルトのプロンプトをすべて受け入れます。別名: -y
--minimal boolean 型定義のみを初期化します (アプリのスキャフォールディングをスキップします)。別名: -m
プロジェクトテンプレート
--react 文字列|ブール値 Scaff

古い React プロジェクト。値を指定せずに使用すると、ベースライン React アプリが作成されます。
プリセットの値を受け入れます: tailwind – Tailwind CSS で事前構成された React アプリ
shadcn – @shadcn/ui と Tailwind CSS を使用した React アプリ
bun init —反応 bun init —react=tailwind bun init —react=shadcn
出力とファイル
(結果) info 選択したオプションのプロジェクト ファイルと構成を初期化します。正確なファイルはテンプレートによって異なります。
ヘルプ
--help boolean このヘルプ メニューを印刷します。別名: -h
例
すべてのデフォルトを受け入れます
ターミナル bun init -y
反応する
ターミナル bun init --react
React + Tailwind CSS
ターミナル bun init --react=tailwind
反応 + @shadcn/ui
ターミナル bun init --react=shadcn

## Original Extract

Scaffold an empty Bun project with the interactive bun init command

bun init - Bun Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Search... Navigation Get Started bun init Runtime Package Manager Bundler Test Runner Guides Reference Blog Feedback Get Started
CLI Usage Initialization Options
Copy page Copy page Scaffold an empty Bun project with the interactive bun init command
Copy page Copy page Scaffold a new Bun project with bun init .
terminal bun init my-app
? Select a project template - Press return to submit.
❯ Blank
React
Library
✓ Select a project template: Blank
+ .gitignore
+ CLAUDE.md
+ .cursor/rules/use-bun-instead-of-node-vite-npm-pnpm.mdc -> CLAUDE.md
+ index.ts
+ tsconfig.json (for editor autocomplete)
+ README.md
Press enter to accept the default answer for each prompt, or pass the -y flag to auto-accept the defaults.
bun init infers settings with sane defaults and is non-destructive when run multiple times.
It creates:
a package.json file with a name that defaults to the current directory name
a tsconfig.json or jsconfig.json file, depending on whether the entry point is a TypeScript file
an entry point, which defaults to index.ts unless any of index.{tsx, jsx, js, mts, mjs} exist or the package.json specifies a module or main field
a CLAUDE.md file when Claude CLI is detected (disable with CLAUDE_CODE_AGENT_RULE_DISABLED env var)
a .cursor/rules/*.mdc file when Cursor is detected, which tells Cursor AI to use Bun instead of Node.js and npm
bun init < folder ?>
​ Initialization Options
--yes boolean Accept all default prompts without asking questions. Alias: -y
​ --minimal boolean Only initialize type definitions (skip app scaffolding). Alias: -m
​ Project Templates
--react string|boolean Scaffold a React project. When used without a value, creates a baseline React app.
Accepts values for presets: tailwind – React app preconfigured with Tailwind CSS
shadcn – React app with @shadcn/ui and Tailwind CSS
bun init —react bun init —react=tailwind bun init —react=shadcn
​ Output & Files
(result) info Initializes project files and configuration for the chosen options. Exact files vary by template.
​ Help
--help boolean Print this help menu. Alias: -h
​ Examples
Accept all defaults
terminal bun init -y
React
terminal bun init --react
React + Tailwind CSS
terminal bun init --react=tailwind
React + @shadcn/ui
terminal bun init --react=shadcn
