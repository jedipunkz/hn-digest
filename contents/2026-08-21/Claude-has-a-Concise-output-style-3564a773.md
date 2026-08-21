---
source: "https://code.claude.com/docs/en/output-styles"
hn_url: "https://news.ycombinator.com/item?id=49392644"
title: "Claude has a \"Concise\" output style"
article_title: "Output styles - Claude Code Docs"
image: "https://claude-code.mintlify.app/_next/image?url=%2F_mintlify%2Fapi%2Fog%3Fdivision%3DModel%2Band%2Bresponses%26appearance%3Dsystem%26title%3DOutput%2Bstyles%26description%3DAdapt%2BClaude%2BCode%2Bfor%2Buses%2Bbeyond%2Bsoftware%2Bengineering%26logoLight%3Dhttps%253A%252F%252Fmintcdn.com%252Fclaude-code%252Fc5r9_6tjPMzFdDDT%252Flogo%252Flight.svg%253Ffit%253Dmax%2526auto%253Dformat%2526n%253Dc5r9_6tjPMzFdDDT%2526q%253D85%2526s%253D78fd01ff4f4340295a4f66e2ea54903c%26logoDark%3Dhttps%253A%252F%252Fmintcdn.com%252Fclaude-code%252Fc5r9_6tjPMzFdDDT%252Flogo%252Fdark.svg%253Ffit%253Dmax%2526auto%253Dformat%2526n%253Dc5r9_6tjPMzFdDDT%2526q%253D85%2526s%253D1298a0c3b3a1da603b190d0de0e31712%26primaryColor%3D%25230E0E0E%26lightColor%3D%2523D4A27F%26backgroundLight%3D%2523FDFDF7%26backgroundDark%3D%252309090B&w=1200&q=100"
author: "ball_of_lint"
captured_at: "2026-08-21T19:19:17Z"
capture_tool: "hn-digest"
hn_id: 49392644
score: 1
comments: 1
posted_at: "2026-08-21T19:15:26Z"
tags:
  - hacker-news
  - translated
---

# Claude has a "Concise" output style

- HN: [49392644](https://news.ycombinator.com/item?id=49392644)
- Source: [code.claude.com](https://code.claude.com/docs/en/output-styles)
- Score: 1
- Comments: 1
- Posted: 2026-08-21T19:15:26Z

## Translation

タイトル: クロードの出力スタイルは「簡潔」です
記事のタイトル: 出力スタイル - Claude Code Docs
説明: クロード コードをソフトウェア エンジニアリングを超えた用途に適応させる

記事本文:
出力スタイル - Claude Code Docs Documentation Index
/docs/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
メイン コンテンツにスキップ Claude Code Docs ホーム ページ 英語 検索... ⌘ K Ask Assistant Claude 開発者プラットフォーム
検索... ナビゲーション モデルと応答 出力スタイル はじめに クロード コードを使用したビルド 管理 構成リファレンス エージェント SDK 新機能 リソース 設定と権限
高速モードで応答を高速化
アドバイザ ツールを使用して難しい意思決定をエスカレーションする
カスタム出力スタイル Frontmatter を作成する
関連機能との比較
ページをコピー ページをコピー クロード コードをソフトウェア エンジニアリングを超えた用途に適応させる
ページをコピー ページをコピー 出力スタイルは、クロードが何を知っているかではなく、クロードがどのように反応するかを変更します。システム プロンプトを変更して、役割、トーン、出力形式を設定します。毎ターン同じ音声やフォーマットを再プロンプトし続ける場合、またはクロードにソフトウェア エンジニア以外の役割を果たしてもらいたい場合に使用します。
カスタム出力スタイルにより、システム プロンプトに指示が追加され、Claude Code の組み込みソフトウェア エンジニアリング指示を保持するかどうかを選択できます。クロードのコミュニケーション方法を変更しながらもコーディングを続ける場合、たとえば常に図で答える場合などに、これらを保持してください。クロードがライティングアシスタントやデータアナリストなど、ソフトウェアエンジニアリングをまったく行っていない場合は、それらを省いてください。
プロジェクト、規約、またはコードベースに関する指示については、代わりに CLAUDE.md を使用してください。
組み込みの出力スタイル
プロアクティブ : クロードは即座に実行し、日常的な決定のために一時停止するのではなく合理的な仮定を立て、計画よりも行動を好みます。これは、自動モードが適用されるよりも強力な自律実行ガイダンスであり、許可モードを変更せずに機能します。

そのため、許可モードによって、ユーザーに尋ねることなく何が実行されるかが決定されます。
Concise : クロードは、デフォルト スタイルと同じくらい徹底的にエンジニアリング作業を行いながら、結果を導き、前文とナレーションをスキップし、デフォルトで応答を短く保ちます。説明や詳細を求めると、クロードはすべてに答えます。クロードは、エラー レポート、セキュリティ警告、破壊的なアクションの確認の完全な内容を常に保管しています。クロード コード v2.1.237 以降が必要です。
説明 : ソフトウェア エンジニアリング タスクを完了するのに役立つ、教育的な「洞察」を提供します。実装の選択とコードベースのパターンを理解するのに役立ちます。
学習 : クロードがコーディング中に「洞察」を共有するだけでなく、小規模で戦略的なコードを自分で提供するよう求める、共同作業型の実践学習モードです。 Claude Code は、実装するための TODO(人間) マーカーをコードに追加します。
ターミナル: /config を実行し、[出力スタイル] を選択してメニューからスタイルを選択します。 Claude Code は、選択内容をローカル プロジェクト レベルで .claude/settings.local.json に保存します。
デスクトップ アプリ: 設定ファイル (たとえば、ターミナル メニューが書き込むファイル .claude/settings.local.json ) に OutputStyle フィールドを設定します。そこで /config を実行すると、Claude Code はメニューではなく [設定] > [Claude Code] を開きます。
スタンドアロンの /output-style コマンドは v2.1.73 で非推奨となり、v2.1.91 で削除されました。 /config を使用するか、outputStyle 設定を直接編集します。
メニューを使用せずにスタイルを設定するには、設定ファイルで OutputStyle フィールドを直接編集します。
{
"outputStyle" : "説明"
}
出力スタイルはシステム プロンプトの一部であり、クロード コードはセッション開始時に一度読み取ります。変更は、/clear または新しいセッションの後に有効になります。出力スタイルの変更が c に与える影響については、「クロード コードがプロンプト キャッシュを使用する方法」を参照してください。

痛み。
カスタム出力スタイルを作成する
プロジェクト: .claude/output-styles
管理ポリシー: 管理設定ディレクトリ内の .claude/output-styles
前付と説明を追加する
---
名前 : 図が最初
description : すべての説明を図で説明します
keep-coding-instructions : true
---
コード、アーキテクチャ、データ フローを説明するときは、構造を示す人魚図から始めて、散文で説明します。
## 図の規則
制御フローには「フローチャート TD」を使用し、リクエスト パスには「シーケンス図」を使用します。図は 15 ノード以下に保ちます。
3 自分のスタイルに切り替えます
Claude Code は、各出力スタイルのカスタム命令をシステム プロンプトの最後に追加します。
すべての出力スタイルは、会話中に出力スタイルの指示に従うようにクロードにリマインダーをトリガーします。
カスタム出力スタイルでは、 keep-coding-instructions が true に設定されていない限り、変更のスコープの設定、コメントの記述、作業の検証方法など、Claude Code の組み込みソフトウェア エンジニアリング命令が省略されます。
関連機能との比較
設定 : OutputStyle フィールドが存在する場所と設定の優先順位の仕組み
権限モード: プロアクティブ スタイルと自動モードの比較
プラグイン: スキル、フック、エージェントとともに出力スタイルをパッケージ化して配布します。
構成をデバッグする: 出力スタイルが有効にならない理由を診断する
はい いいえ アドバイザ ツールを使用して難しい意思決定をエスカレーションする 端末構成 ⌘ I Claude Code Docs ホームページ x linkedin Company

## Original Extract

Adapt Claude Code for uses beyond software engineering

Output styles - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant Claude Developer Platform
Search... Navigation Model and responses Output styles Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Settings and permissions
Speed up responses with fast mode
Escalate hard decisions with the advisor tool
Create a custom output style Frontmatter
Comparisons to related features
Copy page Copy page Adapt Claude Code for uses beyond software engineering
Copy page Copy page Output styles change how Claude responds, not what Claude knows. They modify the system prompt to set role, tone, and output format. Use one when you keep re-prompting for the same voice or format every turn, or when you want Claude to act as something other than a software engineer.
A custom output style adds your instructions to the system prompt and lets you choose whether to keep Claude Code’s built-in software engineering instructions. Keep them when you’re changing how Claude communicates but still coding, like always answering with a diagram. Leave them out when Claude isn’t doing software engineering at all, like a writing assistant or data analyst.
For instructions about your project, conventions, or codebase, use CLAUDE.md instead.
​ Built-in output styles
Proactive : Claude executes immediately, makes reasonable assumptions instead of pausing for routine decisions, and prefers action over planning. This is stronger autonomous-execution guidance than auto mode applies, and it works without changing your permission mode, so your permission mode still decides what runs without asking you.
Concise : Claude leads with the result, skips preamble and narration, and keeps responses short by default, while doing the engineering work as thoroughly as in the Default style. When you ask for an explanation or more detail, Claude answers in full. Claude always keeps the complete content of error reports, security warnings, and confirmations for destructive actions. Requires Claude Code v2.1.237 or later.
Explanatory : Provides educational “Insights” in between helping you complete software engineering tasks. Helps you understand implementation choices and codebase patterns.
Learning : Collaborative, learn-by-doing mode where Claude will not only share “Insights” while coding, but also ask you to contribute small, strategic pieces of code yourself. Claude Code will add TODO(human) markers in your code for you to implement.
Terminal : run /config and select Output style to pick a style from a menu. Claude Code saves your selection to .claude/settings.local.json at the local project level .
Desktop app : set the outputStyle field in a settings file, for example .claude/settings.local.json , the file the terminal menu writes. When you run /config there, Claude Code opens Settings > Claude Code rather than a menu.
The standalone /output-style command was deprecated in v2.1.73 and removed in v2.1.91. Use /config or edit the outputStyle setting directly.
To set a style without the menu, edit the outputStyle field directly in a settings file:
{
"outputStyle" : "Explanatory"
}
Output style is part of the system prompt, which Claude Code reads once at session start. Changes take effect after /clear or a new session. See How Claude Code uses prompt caching for what an output style change does to the cache.
​ Create a custom output style
Project: .claude/output-styles
Managed policy: .claude/output-styles inside the managed settings directory
Add frontmatter and instructions
---
name : Diagrams first
description : Lead every explanation with a diagram
keep-coding-instructions : true
---
When explaining code, architecture, or data flow, start with a Mermaid diagram showing the structure, then explain in prose.
## Diagram conventions
Use `flowchart TD` for control flow and `sequenceDiagram` for request paths. Keep diagrams under 15 nodes.
3 Switch to your style
Claude Code adds each output style’s custom instructions to the end of the system prompt.
All output styles trigger reminders for Claude to adhere to the output style instructions during the conversation.
Custom output styles leave out Claude Code’s built-in software engineering instructions, such as how to scope changes, write comments, and verify work, unless keep-coding-instructions is set to true .
​ Comparisons to related features
Settings : where the outputStyle field lives and how settings precedence works
Permission modes : how the Proactive style compares to auto mode
Plugins : package and distribute output styles alongside skills, hooks, and agents
Debug your configuration : diagnose why an output style isn’t taking effect
Yes No Escalate hard decisions with the advisor tool Terminal configuration ⌘ I Claude Code Docs home page x linkedin Company
