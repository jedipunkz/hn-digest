---
source: "https://marketplace.visualstudio.com/items?itemName=sdevries.prompt-foundry"
hn_url: "https://news.ycombinator.com/item?id=48588396"
title: "Show HN: writing tasks with focused context + instructions = sharper AI"
article_title: "Prompt Foundry - Visual Studio Marketplace"
author: "espresso_dev"
captured_at: "2026-06-18T18:55:58Z"
capture_tool: "hn-digest"
hn_id: 48588396
score: 1
comments: 0
posted_at: "2026-06-18T17:08:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: writing tasks with focused context + instructions = sharper AI

- HN: [48588396](https://news.ycombinator.com/item?id=48588396)
- Source: [marketplace.visualstudio.com](https://marketplace.visualstudio.com/items?itemName=sdevries.prompt-foundry)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T17:08:33Z

## Translation

タイトル: Show HN: 焦点を絞ったコンテキスト + 指示を使用してタスクを作成する = よりシャープな AI
記事のタイトル: Prompt Foundry - Visual Studio Marketplace
説明: Visual Studio Code の拡張機能 - 進化するブロックのライブラリから複雑な AI プロンプトと仕様を作成します。
HN テキスト: AI にタスクをうまく実行させるには、AI に適切なコンテキストと適切なタスクの説明を与える必要があります。 AI を既存のパターンとアーキテクチャに従わせるために、大規模なコードベースで作業するのに多くの苦労を経験しました。
そこで私は、特定のタスクに関連するサブプロンプト/指示を素早く結合できる VS Code / Cursor 拡張機能の構築を開始しました。この価値を理解した後、リキッドシンタックスエンジンを使用することでサブプロンプトをさらにカスタマイズできることに気付きました。プロジェクトが進行するにつれて元のコンテキスト ブロックが古くなるまで、数日間は問題なく動作しました。私はこれらのブロックを更新してフィードバック ループを作成できる MCP サーバーを作成しました。数日後、AI が進行中にログと意思決定記録を書き込めるように、修正オプションも追加しました。この時点で、ここには多くのチャンスがあることに気付きました。AI が最初の 5 分を調査に費やすことがなくなり、コンテキスト ウィンドウをより適切に制御できるようになり、リポジトリを切り替えるときに Agents.md ファイルから漏れる無関係なデータが少なくなり、別個のワークフロー用にプロンプ​​ト グループを作成できるようになりました。
私にとって最後の障害は、実際に Claude Code を使用し、VS Code をあまり使用しなくなったことです。そこで、外部エディターのコマンドをフックして、CC から直接 TUI を開き、そこからプロンプトを記述する方法を考え出しました。過去 5 か月かけて完成させた後、ようやく公開できることを嬉しく思います。ぜひ試してみて、ご意見をお聞かせください。たとえそれがあなたにとって興味のないものであっても、なぜですか?詳細については、GH を参照してください: https://github.com/simondevri

es/prompt-foundry vscode をインストールします: https://marketplace.visualstudio.com/items?itemName=sdevries... OpenVsx をインストールします: https://open-vsx.org/extension/sdevries/prompt-foundry

記事本文:
Prompt Foundry - Visual Studio マーケットプレイス
コンテンツにスキップ
|マーケットプレイス
サインイン
Visual Studio Code > プログラミング言語 > Prompt Foundry Visual Studio Code は初めてですか?今すぐ入手してください。プロンプトファウンドリ
ジャンプ先: デモ |セットアップ手順
Prompt Foundry は、大規模なコードベースで AI の一貫性と効率性を高める VS Code 拡張機能です。セッションごとにコンテキストを書き換えるのではなく、再利用可能なプロンプト ブロック* (アーキテクチャ ノート、コーディング標準、動作規約) の個人ライブラリを構築し、各タスクの前にそれらをスナップします。バンドルされた MCP サーバーを使用すると、AI が検出結果をそのライブラリに書き戻すことができるため、手動でメンテナンスを行うことなく、時間の経過とともにプロンプ​​トが改善されます。
*プロンプト ブロックは、他のブロックおよび AI に送信する前に完全なプロンプトを形成するための主な命令とともにコンパイルされる、ライ​​ブラリ内のテンプレート プロンプトを表すマークダウン ファイルです。
AI は、Prompt Foundry の TUI バージョンを構築するように依頼されました (これはバンドルされた機能になりました!)。実行 #1 はベースライン プロンプトです。実行 #2 は Prompt Foundry を使用します。どちらも Gemini Flash Lite 3.1 を使用します。
主な違いは、より優れたコード構造と、はるかに少ないハンドリングです。
実行 #2 もプロンプト ブロックの指示に従い、定義されたスタイルに従ってコードを生成しました。
詳細: ベースライン |鋳造工場の稼働
どちらの実行でも、インク (React ベースの TUI) と Clipboardy を使用しました。この TUI がコア バックエンドを VS Code 拡張機能と共有する場合、これが意図された方向であるため、アーキテクチャの違いが最も重要になります。
ベースライン (#1) では、command と meow が追加されました。これは、同じ仕事を行う 2 つの CLI 引数パーサーです。また、 ink とともにターミナル スタイル用のチョークも追加され、ink の宣言型 React コンポーネント モデルと命令型 ANSI エスケープ コードという 2 つの競合するアプローチが作成されました。この種の混合パラダイム依存関係セットは、コードベースが成長するにつれて摩擦を引き起こします。

また、テスト インフラストラクチャも含まれていませんでした。
実行 #2 では、両方の冗長な依存関係を削除し、ink のコンポーネント モデル内でスタイルを維持し、 ink-select-input でコンポーネント セットを拡張し、 jest 、 ts-jest 、および ink-testing-library を追加しました。コア ビジネス ロジック ( LibraryManager 、 SessionManager 、 PromptCompiler ) には、設計上すでに VS Code への依存関係がありません。実行 #2 のテスト セットアップは、VS Code インスタンスやライブ端末をスピンアップせずに、TUI コンポーネントとコア ロジックの両方を分離して単体テストできることを意味します。
つまり、実行 #1 は機能しますが、負債が蓄積します。実行 #2 は、プロジェクトがどこへ向かっているのかについての理解を反映しています。
プロンプト ブロック (指示と情報) を選択します。
最も重要なものを目標としてマークする
MCP でループを閉じる: セッションの終了時に、ライブラリを更新するように AI に指示します。 *「今決定した内容でプロンプト ファウンドリ mcp を使用して認証アーキテクチャ ブロックを更新します。」次のセッションでは、その知識はすでに存在しています。
MCP サーバーは、プロンプト ブロック ライブラリを読み取りおよび書き込み可能なファイルとして公開します。セッション中、AI は次のことができます。
タスク中に任意のブロックを読み取り、最新のコンテキストを取得します
ブロックに追加 - 決定の実行ログを保持するのに役立ちます
変更されたコンテンツでブロックを上書きします。コミットする前に内容を確認します。
これはコンテキスト複合体を意味します。アーキテクチャ上の決定、注意点、命名規則など、覚えておく価値のあるものはすべて適切な場所に書き戻され、次のセッションに備えられます。
出力プロンプト構造の例:
VS Code Marketplace から拡張機能をインストールします。これは MCP サーバーと TUI にバンドルされています。この拡張機能には、すぐに使用できるデフォルトのブロック ライブラリが含まれています。独自のプロンプト ライブラリ ディレクトリを作成するまで、ブロックは読み取り専用です。プロンプト ブロックの横にある歯車アイコンまたは VS Code 設定を通じて場所を設定します。
ない

e: エディターから編集する場合は、VS Code でプロンプト ライブラリ フォルダーを開き、[信頼] をクリックする必要があります。このフォルダーには、.md ファイルとプロンプト設定ファイルのみが含まれます。
注: MCP サーバーでは、マシンに Node.js がインストールされている必要があります。
「プロンプト ブロック ライブラリ」の横にある歯車をクリックします
[機能と権限] > [MCP サーバー] をクリックします。
生成されたJSONスニペットをコピーします
AI の MCP 構成に追加します。
注: MCP サーバーは、VS Code 拡張フォルダー内の Node.js プロセスとしてローカルで実行されます。 AI は、MCP ツール呼び出しを通じて、指定されたプロンプト ライブラリ フォルダーのコンテンツを読み取り、変更できます。
TUI は、プロンプト ブロックを別のアプリのプロンプトにすばやく添付する方法を提供します。たとえば、Claude Code では、Ctrl+G を押して外部エディターを開き、このツールを使用するように構成できます。
「プロンプト ブロック ライブラリ」の横にある歯車をクリックします
クロード コードの指示に従って、または外部エディタを設定します。
上部の指示ボックスにメインの指示プロンプトを入力します。
ライブ フォーカス (⚡) ボタンを使用すると、IDE でファイルと行を選択し、それらの場所をプロンプトに追加できます。コードベースをナビゲートしながらディクテーションするのに便利です。ナビゲートすると、誰かがあなたの画面を見ながら説明するのと同じように、コンテキスト ファイル タグが表示されます。
プロンプト ブロックは、再利用可能な命令や情報を含むマークダウン ファイルです。独自のプロンプト ライブラリがありますが、カスタム ディレクトリまたはワークスペース ディレクトリをアタッチすることもできます。簡単な例を次に示します。
# コードスタイル
- TypeScript 厳密モードを使用する
- デフォルトのエクスポートよりも名前付きエクスポートを優先します
- すべてのパブリック関数に JSDoc コメントを追加します
ブロックは、フォルダーごとに 1 つのカテゴリーに加えて、一連の特別なカテゴリーに分類されます。必要に応じて、クロードまたはカーソルのスキルも追加します。
ブロックには短い参照、つまり pr 内の特定の位置にコンパイルされたリマインダーを含めることができます。

オムプト。これにより、すべての命令を 1 か所にダンプするのではなく、AI がワークフロー内のどこでそれを認識するかを制御します。
ReferenceLocation は、コンパイルされたプロンプト内のどこにリマインダーが表示されるかを制御します。
ブロックにスターを付けてゴールとしてマークします。その参照テキストは、プロンプトの最後にある専用の # 「このタスクを完了する際の主要な目標:」セクションに取り込まれ、AI が作業を開始する前に最も重要なことの明確な概要を提供します。
プロンプト ブロックは Liquid 構文をサポートします。ブロックを書き換えずにカスタム変数をブロックに追加する場合に便利です。
{% コメント %}
変数:
選択例:
タイプ: 選択
オプション: [
「ひとつ」、
「二つ」、
】
テキスト例:
タイプ: テキスト
{% 終了コメント %}
オプションを選択: {{ selectExample }}
テキストの例: {{ textExample }}
特別なカテゴリー
AI 契約 (編集可能なテンプレート): 役割、コメント スタイル、その他の期待される行動を定義します。この拡張機能は、AI が契約を遵守するよう促すプロンプトを構造化します。
ツール:
Git コミット: 特定の git コミットを追加します。
Git diff: ブランチまたはコミット ハッシュに対する差分を追加します。
IDE 診断: IDE MCP がセットアップされていない場合は、詳細なエラーを共有します。
アクティブなシンボル: コンテキストのファイル概要を追加します。
クロードのスキルとコマンド: グローバル クロード フォルダーからクロードのスキルとコマンドを一覧表示し、AI にそれらを使用するよう促します。
同じプロンプト ブロックを頻繁に追加する必要がある場合は、グループ機能を使用します。グループを作成するには、プロンプトにいくつかのブロックを追加し、グループの横にある「+」ボタンを押します。アクティブなプロンプトはグループに保存されます。
プライバシー: 100% ローカル。テレメトリもデータ収集もありません。プロンプトがマシンから離れることはありません。
テレメトリーがないということは、私があなたに依存していることを意味します。提案や機能リクエストを歓迎します:
https://form.typeform.com/to/hAc2CQ6A

## Original Extract

Extension for Visual Studio Code - Forge complex AI prompts and specs from a library of evolving blocks.

To get an AI to do a task well, you need to give it the right context and the right task description. I've had a lot of struggles working in a large codebase getting AI to follow the existing patterns and architecture.
So I started building a VS Code / Cursor extension that lets you quickly snap together sub prompts/instructions relevant to a specific task. After seeing the value of this, I realized you could make those sub prompts more customizable by using a liquid syntax engine. That worked fine for a few days, until the original context blocks got stale as the project progressed. I wrote an MCP server that can update those blocks, creating a feedback loop. A few days later I added an amend option too, so the AI can write a log and decision record as it goes. At this point I realized there were a lot of opportunities here: my AI isn't spending the first 5 minutes doing research, I've got better control over my context window, there's less irrelevant data leaking in from my agents.md file when I switch repos, I can write prompt groups for separate workflows, and so on.
The last blocker for me was that I actually use Claude Code and not VS Code much anymore. So I figured out how to hook into the external editor command to open a TUI directly from CC and write a prompt from there. Now I'm happy to finally share it after spending the last 5 months perfecting it. Would love for you to try it and let me know what you think, even if its not something for you - why? More details see GH: https://github.com/simondevries/prompt-foundry install vscode: https://marketplace.visualstudio.com/items?itemName=sdevries... Install OpenVsx: https://open-vsx.org/extension/sdevries/prompt-foundry

Prompt Foundry - Visual Studio Marketplace
Skip to content
| Marketplace
Sign in
Visual Studio Code > Programming Languages > Prompt Foundry New to Visual Studio Code? Get it now. Prompt Foundry
Jump to: Demo | Setup instructions
Prompt Foundry is a VS Code extension that makes AI more consistent and effective in large codebases. Instead of rewriting context every session, you build a personal library of reusable prompt blocks* — architecture notes, coding standards, behavioural contracts — and snap them together before each task. A bundled MCP server lets the AI write discoveries back into that library, so your prompts improve over time without manual upkeep.
*A prompt block a markdown file that represents template prompt in your library which is compiled with other blocks and your main instruction to form a complete prompt before sending to the AI.
The AI was asked to build a TUI version of Prompt Foundry (this is now a bundled feature!). Run #1 is a baseline prompt. Run #2 uses Prompt Foundry. Both use Gemini Flash Lite 3.1.
The key differences: better code structure and far less handholding.
Run #2 also followed instructions in the prompt blocks, generating code according to the defined style.
Details: Baseline | Foundry run
Both runs used ink (React-based TUI) and clipboardy . The architectural differences matter most if this TUI is to share a core backend with the VS Code extension, which is the intended direction.
Baseline (#1) added commander and meow — two CLI argument parsers doing the same job. It also added chalk for terminal styling alongside ink , creating two competing approaches: ink's declarative React component model vs. imperative ANSI escape codes. This kind of mixed-paradigm dependency set creates friction as the codebase grows. It also included no test infrastructure.
Run #2 dropped both redundant dependencies, kept styling within ink's component model, extended the component set with ink-select-input , and added jest , ts-jest , and ink-testing-library . The core business logic ( LibraryManager , SessionManager , PromptCompiler ) already has zero VS Code dependencies by design. Run #2 's test setup means TUI components and core logic can both be unit-tested in isolation, without spinning up a VS Code instance or a live terminal.
In short: Run #1 works but accumulates debt. Run #2 reflects an understanding of where the project is heading.
Select your prompt blocks (instructions and information)
Mark the most important ones as goals
Close the loop with MCP: At the end of a session, tell the AI to update your library — e.g. *"update the auth-architecture block using the prompt foundry mcp with what we just decided." Next session, that knowledge is already there.
The MCP server exposes your prompt block library as readable and writable files. During a session the AI can:
Read any block to get up-to-date context mid-task
Append to a block — useful for keeping a running log of decisions
Overwrite a block with revised content, which you review before committing
This means context compounds. Architectural decisions, gotchas, naming conventions — anything worth remembering gets written back in, in the right place, ready for the next session.
Example output prompt structure:
Install the extension from the VS Code Marketplace . This bundles with the MCP server and TUI. The extension includes a default block library to get you started. Blocks are read-only until you create your own prompt library directory — set the location via the gear icon next to prompt blocks or through VS Code settings.
Note: When editing from the editor you need to open the prompt library folder in VS Code and click 'Trust'. The folder only contains .md files and prompt settings files.
Note: The MCP server requires Node.js to be installed on your machine.
Click the gear next to Prompt Block Library
Click Features & Permissions > MCP Server
Copy the generated JSON snippet
Add it to your AI's MCP config
Note: The MCP server runs locally as a Node.js process in the VS Code extension folder. The AI can read and modify the content of the specified prompt library folder via MCP tool calls.
The TUI provides a way to quickly attach prompt blocks to a prompt in another app. For instance, in Claude Code you can press Ctrl+G to open an external editor, which can be configured to use this tool.
Click the gear next to Prompt Block Library
Follow the instructions for Claude Code or to set the external editor
Enter your main instructional prompt into the top instruction box.
The live focus (⚡) button lets you select files and lines in the IDE and adds those locations to the prompt. Useful for dictating while navigating a codebase — you end up with contextual file tags as you navigate, similar to someone watching your screen as you explain.
A prompt block is a Markdown file containing reusable instructions or information. You have your own prompt library but can also attach a custom directory or workspace directory. Here is a simple example:
# Code style
- Use TypeScript strict mode
- Prefer named exports over default exports
- Add JSDoc comments to all public functions
Blocks are organised into categories, one per folder, plus a set of special categories. Optionally add your Claude or Cursor skills too.
A block can include a short reference — a reminder compiled into a specific position in the prompt. This controls where in the workflow the AI sees it, rather than dumping all instructions in one place.
referenceLocation controls where in the compiled prompt the reminder appears:
Star a block to mark it as a goal. Its reference text gets pulled into a dedicated # Key goals while completing this task: section at the end of the prompt, giving the AI a clear summary of what matters most before it starts work.
Prompt blocks support Liquid syntax. Useful when you want to add custom variables to a block without rewriting it:
{% comment %}
vars:
selectExample:
type: select
options: [
"One",
"Two",
]
textExample:
type: text
{% endcomment %}
Select option: {{ selectExample }}
Text example: {{ textExample }}
Special categories
AI Contract (editable template): Define role, comment style, and other behavioural expectations. The extension structures the prompt to encourage the AI to stick to the contract.
Tools:
Git commit: Add a specific git commit.
Git diff: Add a diff against a branch or commit hash.
IDE diagnostics: Share detailed errors if you don't have IDE MCP set up.
Active symbols: Add file summaries for context.
Claude skills and commands: Lists Claude skills and commands from your global Claude folder and prompts the AI to use them.
When you frequently need to add the same prompt blocks, use the group feature. To create a group, add some blocks to the prompt and then press the '+' button next to group. The active prompts will be saved to the group.
Privacy: 100% local. No telemetry, no data collection. Your prompts never leave your machine.
No telemetry means I rely on you. Suggestions and feature requests welcome:
https://form.typeform.com/to/hAc2CQ6A
