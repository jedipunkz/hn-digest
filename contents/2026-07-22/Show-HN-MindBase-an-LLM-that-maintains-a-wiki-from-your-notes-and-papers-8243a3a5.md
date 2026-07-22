---
source: "https://github.com/frankchu91/mindbase"
hn_url: "https://news.ycombinator.com/item?id=49001892"
title: "Show HN: MindBase – an LLM that maintains a wiki from your notes and papers"
article_title: "GitHub - frankchu91/mindbase · GitHub"
author: "haobing0304"
captured_at: "2026-07-22T04:59:29Z"
capture_tool: "hn-digest"
hn_id: 49001892
score: 1
comments: 0
posted_at: "2026-07-22T04:27:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: MindBase – an LLM that maintains a wiki from your notes and papers

- HN: [49001892](https://news.ycombinator.com/item?id=49001892)
- Source: [github.com](https://github.com/frankchu91/mindbase)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T04:27:11Z

## Translation

タイトル: Show HN: MindBase – メモや論文から wiki を管理する LLM
記事タイトル: GitHub - Frankchu91/mindbase · GitHub
説明: GitHub でアカウントを作成して、frankchu91/mindbase の開発に貢献します。

記事本文:
GitHub - Frankchu91/mindbase · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
フランチュ91
/
マインドベース
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 Co

de 「その他のアクション」メニューを開く フォルダーとファイル
39 コミット 39 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows apps apps docs docs package/ core package/ core schema schema .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md Mindbase.config.example.json Mindbase.config.example.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ソースから Wiki を構築および維持する AI 研究アシスタント。 RAG-in-a-vector-DB ではありません。ディスク上にある本物のマークダウン Wiki で、会話の合間に LLM が自動的に管理します。
MindBase は、Andrej Karpathy の LLM-Wiki パターンを製品として実装しています。あなたはソース（論文、記事、考え）をそれにフィードします。 LLM は、構造化された Wiki ページの読み取り、相互参照、矛盾のフラグ付け、書き込みを行います。後で質問すると、Wiki にはすでに合成された回答が用意されており、クエリ時にベクトル検索を再導出する必要はありません。
ステータス: 早期アクセス。フィードバック用のベータ版 (2026-Q3)。データ モデル + コア ループは安定しています。一部の UI 機能はまだ v1 → v2 に移行中です。
インストール — AI エディターを選択します
クロード コード (フラッグシップ エクスペリエンス)
その他の MCP 互換クライアント
複数のプロジェクトで作業する
よく読んでいますね。論文、記事、ツイート、ドキュメント。あなたはそれらを覚えておき、それらを結びつけ、それらから意見を形成したいと考えています。現在、悪い選択肢が 2 つあります。
Notion / Obsidian / Roam: パッシブコンテナ。整理整頓はすべてあなたが行います。 AI 機能は、メンテナンスではなく、ボルトオンで生成されます。
NotebookLM / Perplexity Pages / ChatGPT 検索: RAG ベース。何も溜まらない。すべての質問は、生の情報源から答えを再導き出します。
MindBase は 3 番目のオプションです: LLM 法

ソースを入力すると、永続的で構造化された Wiki が維持されます。知識が複合します。あなたが貢献するたびに、context.md はより鮮明になります。 AI はセッションをまたいであなたを覚えています。あなたの信念はマークダウン ファイルに記録され、要約されるチャット履歴には保存されないからです。
これは、AI インターンがあなたのために書き、最新の状態に保たれ、相互参照され、知らないことについて正直に書かれる個人的なウィキペディアと考えてください。
ディスク上の 3 つの物理層 (Karpathy のモデル):
レイヤ 1 — ソース/これらを追加します。追加のみ。 LLM は読み取りを行いますが、書き換えは行いません。
§─ 寄稿者/ ユーザーごとの日付が記載されたあなたの考え (実験ノートのような)
§─ LLM が作成する研究/Wiki ページ (概念、エンティティ、合成)
└─ raw/PDF、URL キャプチャ、ペースト — オリジナルのバイナリ
レイヤ 2 — README.md ウィキのルール。あなたが編集します。 LLM はすべての操作で読み取ります。
context.md 合成された現在の真実。 LLMは書いています。読みましたね。
Index.yaml 自動生成されたファイル カタログ。 LLMはこう主張する。
レイヤ 3 — ログ/すべての操作の時系列ログ。 LLM が追加されます。
アーティファクト/生成された出力: 日次概要、エクスポート、lint レポート。
AI が実行できる 3 つの操作:
これらはループを形成します。良い回答と糸くずの発見が新しい投稿として戻ってくるため、Wiki は次のように複雑になります。
フローチャート LR
C["貢献<br/><small>考え · PDF · URL</small>"] --> S[("sources/<br/><small>追加のみ</small>")]
S --> B["ビルド<br/><small>LLM が合成</small>"]
B --> W[("context.md<br/>research/*.md")]
W --> Q[「質問<br/><small>引用された回答</small>」]
W --> L["糸くず<br/><small>矛盾 · 孤児 · ギャップ</small>"]
Ｑ－。 「良い回答がファイルバックされました」 .-> C
L-。 「フォローアップノート」 .-> C
読み込み中
それが製品全体です。それ以外はすべてUXです。
vs. Notion: データ (独自の DB ではなくマークダウン ファイル) を所有しているのはあなたです。 AI はメインタです

iner 、ボルトオン発電機ではありません。 context.md は進化します。 Notion ページにはありません。
vs. Obsidian: 同じローカル ファーストのマークダウン ボールトが得られますが、AI がウィキを積極的に作成するため、すべてのメモを手作りする必要はありません。 [[wikilinks]] の相互参照が付属しています。
vs. NotebookLM / Perplexity: 知識の複合体。 「RAG についての私の見解は何ですか?」と尋ねてください。そして答えはすでに context.md に存在します。 NotebookLM は毎回生のソースからそれを再派生します。何も蓄積されません。
他の RAG ツールとの比較: RAG はステートレスな取得です。 MindBase はステートフル合成です。 Wiki は取り込むたびに内容が充実していきます。時間の経過とともに、回答はより速く、より一貫性のあるものになります。
Node.js 20+ (Cursor / Windsurf / Cline / Continue については以上です。MCP サーバーは npx 経由で自体をインストールします)
次のいずれか: Claude Code、Cursor、Windsurf、Cline、Continue.dev、または別の MCP 互換 AI エディター
pnpm 10+ は、Claude Code プラグインまたは Web UI をソースからインストールする場合のみ
(オプション) Web UI が必要な場合は、最新のブラウザー
MindBase は完全にマシン上で実行されます。クラウドへのアップロード、サインアップ、テレメトリはありません。
インストール — AI エディターを選択します
MindBase は MCP (Model Context Protocol) サーバーとして出荷され、 npm で mindbase-mcp として公開されます。すべての MCP 互換 AI エディターは、1 行の構成でツールを使用できます。クローンもビルドもありません。 Claude Code は、プラグインを通じてさらに洗練されたレイヤー (スラッシュ コマンド、サブエージェント) を取得します。
クロード コード (フラッグシップ エクスペリエンス)
サブエージェント、スラッシュ コマンド、自動コンテキスト、およびエージェントごとのツール境界を使用して、Karpathy の 8 ステップの完全な取り込みを取得します。
プラグインをインストールします。Claude Code 内に 2 つのコマンドがあり、クローンやビルドはありません。
/プラグイン マーケットプレイス追加 Frankchu91/mindbase
/プラグインのインストール mb@mindbase
プロンプトが表示されたら、Claude Code を再起動します。 (プラグインは npx -y Mindbase-mcp 経由で MCP サーバーを起動するため、npm はサーバーに

自動的に。)
git clone https://github.com/frankchu91/mindbase.git
cd マインドベース && pnpm インストール
クロード --plugin-dir アプリ/プラグイン
確認: クロード コードを開き、 / と入力します。 /mb:contribute 、 /mb:ask 、 /mb:build 、 /mb:status 、およびその他 8 つが表示されるはずです。 /mcp を実行します — 49 個のツールに接続された mb が表示されるはずです。
セキュリティ境界を持つ 5 つのサブエージェント ( contributor 、 builder 、 curator 、 Researcher 、 migrator ) — それぞれに厳格なツール許可リストがあります
SessionStart フック — 新しい Claude Code セッションごとに、プロジェクトの README + コンテキスト + 最近のログ エントリが自動挿入されます。
Karpathy の完全な 8 ステップの取り込み: 読む → 要点について議論する → 計画を提案する → 承認を待つ → 実行する
~/.cursor/mcp.json に追加します (ファイルがない場合は作成します)。
{
"mcpサーバー": {
"マインドベース" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " mindbase-mcp " ]
}
}
}
カーソルを再開します。カーソル チャットを開く → ツール ピッカーには、mindbase_contribute 、mindbase_ask_wiki 、およびその他 47 個のリストが表示されます。
オプションですが推奨: ~/.cursor/rules.md (またはプロジェクトごとの .cursorrules ) に追加して、Cursor の LLM に MindBase の規則を教えます。
# MindBase の規約
ユーザーが「mindbase」、「mb」、「my wiki」、「记一下」、または「记忆」に言及すると、
次の MCP ツールが利用可能です:mindbase_contribute、mindbase_ask_wiki、
minbase_status、mindbase_load_project など。
ルール:
1. ユーザーが「mindbase X に追加」または「记一下 X」と言ったときは、常に Mindbase_contribute を呼び出します。
text=X の場合。ツールを呼び出さずに承認だけをしないでください。
2. ユーザーが「mindbase X に質問する」または「问问我的 wiki」と言ったら、mindbase_ask_wiki を呼び出します。
クエリ=Xの場合。
3. PDF/URL を取り込む場合は、8 ステップの Karpathy ループに従います。
読む → 3 つの重要なポイントについてユーザーと話し合う → 計画を提案する → 承認を待つ
→ 次に、概要を指定してmindbase_contributeを呼び出します。
4. デフォルトのプロジェクトは config.json から取得されます

現在のプロジェクトID。ユーザーがこう言うなら
「for project X」または「在 X 项目里」では、projectId=X をツール呼び出しに渡します。
最初のプロジェクトにスキップします。
Windsurf は同じ MCP 構成フォーマットを使用します。カスケード設定→MCP→追加を開きます：
{
"mcpサーバー": {
"マインドベース" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " mindbase-mcp " ]
}
}
}
ウィンドサーフィンを再起動します。 Cursor と同じルール ファイル アプローチです。MindBase 規則ブロックをカスケード ルールに配置します。
VSCode → Cline 拡張設定 → MCP サーバー を開きます。追加:
{
"mcpサーバー": {
"マインドベース" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " mindbase-mcp " ]
}
}
}
クラインは自動検出します。すべてのツール呼び出しには、デフォルトで確認ダイアログが表示されます (透明性が優先)。最初のプロジェクトにスキップします。
~/.Continue/config.json を開きます。 Experimental.modelContextProtocolServers の下に追加します。
{
"実験的" : {
"modelContextProtocolServers" : [
{
"輸送" : {
"タイプ" : " stdio " ,
"コマンド" : " npx " ,
"args" : [ " -y " 、 " mindbase-mcp " ]
}
}
】
}
}
VSCodeを再起動します。 [チャットを続行] では、MCP ツールが @ の下に表示されます。最初のプロジェクトにスキップします。
その他の MCP 互換クライアント
エディターが MCP (Zed、Aider、Goose、ChatGPT デスクトップの将来の MCP サポート、カスタム Claude Agent SDK アプリ) をサポートしている場合、パターンは同じです。
npx -y マインドベース-mcp
標準入出力サーバーとして。正確な構成の場所については、クライアントの MCP ドキュメントを参照してください。
エディターが MCP サーバーに接続したら、最初のプロジェクトを作成します。
/mb:init my-research --mission "トランスフォーマーの注意メカニズムを研究する"
Claude は、 ~/mindbase-data/projects/my-research/ を README.md 、 context.md 、 Index.yaml 、および空のsources/ 、 logs/ 、 artifacts/ ディレクトリでスキャフォールディングし、それを現在のプロジェクトとして設定します。
カーソル / ウィンドサーフィン / クライン / 継続
my-research aboutTransformer という新しいマインドベース プロジェクトを作成します。
注意のメカニズム。
の

LLM は、mindbase_init_project({ name: "my-research", Mission: "..." }) を呼び出して確認します。
クロードコードの /mb:status #
または、他の IDE に尋ねます。
私のマインドベースはどうなっているでしょうか？
以下が表示されるはずです:
MindBase — プロジェクト: my-research
場所: ~/mindbase-data/projects/my-research
README 28 行
context.md 8 行 (最後にビルドされたもの: なし)
Index.yaml 5行
寄稿者: 0 ユーザー、0 エントリ
出典: 0 研究、0 生
ログ: 1 日
アーティファクト: 0 出力
準備は完了です。
毎日やるべき4つのこと。
短い観察、決定、または注目の意見を追加します。儀式は行わずに、今日の貢献者ファイルに直接入ります。
/mb:contribute "注意は基本的にソフト辞書検索です - 各クエリ
キーとの類似性に基づいて値の加重平均を取得します。」
他の IDE:
私のマインドベースに追加: 注意は基本的にソフトな辞書検索です - それぞれ
クエリは、キーとの類似性に基づいて値の加重平均を取得します。
結果: ~/mindbase-data/projects/my-research/sources/contributors/<you>/2026-07-11.md に追加され、操作が記録されます。
2. 実質的なソース（論文、記事、PDF）を摂取する
実際のソースについては、MindBase が Karpathy の 8 ステップの取り込みを詳しく説明しています。
クロード・コード (フラッグシップ — サブエージェント付き):
/mb:contribute /ダウンロード/attention-is-all-you-need.pdf
貢献者のサブエージェントは次のことを行います。
3 つの重要なポイントをお知らせします
提案してください

[切り捨てられた]

## Original Extract

Contribute to frankchu91/mindbase development by creating an account on GitHub.

GitHub - frankchu91/mindbase · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
frankchu91
/
mindbase
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
39 Commits 39 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows apps apps docs docs packages/ core packages/ core schema schema .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md mindbase.config.example.json mindbase.config.example.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
An AI research assistant that builds and maintains a wiki from your sources. Not RAG-in-a-vector-DB. A real markdown wiki on your disk, that an LLM gardens for you between conversations.
MindBase implements Andrej Karpathy's LLM-Wiki pattern as a product. You feed it sources (papers, articles, thoughts). The LLM reads, cross-references, flags contradictions, and writes structured wiki pages. Later, when you ask a question, the wiki already has the synthesized answer — no vector search re-derivation at query time.
Status: Early access. Beta for feedback (2026-Q3). Data model + core loop stable; some UI features still in v1 → v2 migration.
Install — choose your AI editor
Claude Code (flagship experience)
Any other MCP-compatible client
Working with multiple projects
You read a lot. Papers, articles, tweets, docs. You want to remember them, connect them, form opinions from them. Today you have two bad options:
Notion / Obsidian / Roam: Passive containers. You do all the organizing. AI features are bolted-on generation, not maintenance.
NotebookLM / Perplexity Pages / ChatGPT search: RAG-based. Nothing accumulates. Every question re-derives the answer from raw sources.
MindBase is the third option: the LLM actively maintains a persistent, structured wiki as you feed it sources. Knowledge compounds. Your context.md gets sharper every time you contribute. The AI remembers you across sessions because your beliefs are written down in markdown files — not stored in a chat history that gets summarized away.
Think of it as a personal Wikipedia that an AI intern writes for you , kept up to date, cross-referenced, and honest about what it doesn't know.
Three physical layers on disk (Karpathy's model):
Layer 1 — sources/ You add these. Append-only. LLM reads, never rewrites.
├─ contributors/ Your thoughts, dated per-user (like a lab notebook)
├─ research/ Wiki pages the LLM writes (concepts, entities, syntheses)
└─ raw/ PDFs, URL captures, pastes — original binaries
Layer 2 — README.md The rules of your wiki. You edit. LLM reads at every op.
context.md The synthesized current truth. LLM writes. You read.
index.yaml Auto-generated file catalog. LLM maintains.
Layer 3 — logs/ Chronological log of every operation. LLM appends.
artifacts/ Generated outputs: daily briefs, exports, lint reports.
Three operations the AI can run:
They form a loop — good answers and lint findings flow back in as new contributions, so the wiki compounds:
flowchart LR
C["contribute<br/><small>thought · PDF · URL</small>"] --> S[("sources/<br/><small>append-only</small>")]
S --> B["build<br/><small>LLM synthesizes</small>"]
B --> W[("context.md<br/>research/*.md")]
W --> Q["ask<br/><small>cited answers</small>"]
W --> L["lint<br/><small>contradictions · orphans · gaps</small>"]
Q -. "good answers filed back" .-> C
L -. "follow-up notes" .-> C
Loading
That's the entire product. Everything else is UX.
vs. Notion: You own the data (markdown files, not a proprietary DB). The AI is a maintainer , not a bolt-on generator. context.md evolves; Notion pages don't.
vs. Obsidian: You get the same local-first markdown vault, but with an AI that actively writes the wiki — you don't have to hand-craft every note. Comes with [[wikilinks]] cross-referencing done for you.
vs. NotebookLM / Perplexity: Knowledge compounds . Ask "what's my view on RAG?" and the answer already exists in context.md . NotebookLM re-derives it from raw sources every time — nothing accumulates.
vs. any RAG tool: RAG is stateless retrieval. MindBase is stateful synthesis . The wiki gets richer with every ingest. Answers get faster and more consistent over time.
Node.js 20+ (that's it for Cursor / Windsurf / Cline / Continue — the MCP server installs itself via npx )
One of: Claude Code, Cursor, Windsurf, Cline, Continue.dev, or another MCP-compatible AI editor
pnpm 10+ only if you install the Claude Code plugin or the web UI from source
(Optional) A modern browser, if you want the web UI
MindBase runs entirely on your machine . No cloud upload, no signup, no telemetry.
Install — choose your AI editor
MindBase ships as an MCP (Model Context Protocol) server — published on npm as mindbase-mcp . Every MCP-compatible AI editor can use its tools with a one-line config; no clone, no build. Claude Code gets an extra layer of polish (slash commands, sub-agents) via the plugin.
Claude Code (flagship experience)
Get the full Karpathy 8-step ingest with sub-agents, slash commands, auto-context, and per-agent tool boundaries.
Install the plugin — two commands inside Claude Code, no clone, no build:
/plugin marketplace add frankchu91/mindbase
/plugin install mb@mindbase
Restart Claude Code when prompted. (The plugin launches its MCP server via npx -y mindbase-mcp , so npm delivers the server automatically.)
git clone https://github.com/frankchu91/mindbase.git
cd mindbase && pnpm install
claude --plugin-dir apps/plugin
Verify: Open Claude Code, type / . You should see /mb:contribute , /mb:ask , /mb:build , /mb:status , and 8 others. Run /mcp — you should see mb connected with 49 tools.
5 sub-agents with security boundaries ( contributor , builder , curator , researcher , migrator ) — each has a strict tool allowlist
SessionStart hook — every new Claude Code session auto-injects your project's README + context + recent log entries
The full Karpathy 8-step ingest: read → discuss takeaways → propose plan → wait for approval → execute
Add to ~/.cursor/mcp.json (create the file if missing):
{
"mcpServers" : {
"mindbase" : {
"command" : " npx " ,
"args" : [ " -y " , " mindbase-mcp " ]
}
}
}
Restart Cursor. Open Cursor Chat → the tool picker should list mindbase_contribute , mindbase_ask_wiki , and 47 others.
Optional but recommended: teach Cursor's LLM about MindBase conventions by adding to ~/.cursor/rules.md (or per-project .cursorrules ):
# MindBase conventions
When the user mentions "mindbase", "mb", "my wiki", "记一下", or "记忆",
the following MCP tools are available: mindbase_contribute, mindbase_ask_wiki,
mindbase_status, mindbase_load_project, and more.
Rules:
1. When user says "add to mindbase X" or "记一下 X", ALWAYS call mindbase_contribute
with text=X. Never just acknowledge without calling the tool.
2. When user says "ask mindbase X" or "问问我的 wiki", call mindbase_ask_wiki
with query=X.
3. When ingesting a PDF/URL, follow the 8-step Karpathy loop:
read → discuss 3 key takeaways with user → propose plan → wait for approval
→ then call mindbase_contribute with the summary.
4. Default project comes from config.json currentProjectId. If the user says
"for project X" or "在 X 项目里", pass projectId=X to the tool call.
Skip to First project .
Windsurf uses the same MCP config format. Open Cascade settings → MCP → add:
{
"mcpServers" : {
"mindbase" : {
"command" : " npx " ,
"args" : [ " -y " , " mindbase-mcp " ]
}
}
}
Restart Windsurf. Same rules-file approach as Cursor works — put a MindBase conventions block in your Cascade rules.
Open VSCode → Cline extension settings → MCP Servers . Add:
{
"mcpServers" : {
"mindbase" : {
"command" : " npx " ,
"args" : [ " -y " , " mindbase-mcp " ]
}
}
}
Cline auto-detects. Every tool call gets a confirmation dialog by default (transparency win). Skip to First project .
Open ~/.continue/config.json . Add under experimental.modelContextProtocolServers :
{
"experimental" : {
"modelContextProtocolServers" : [
{
"transport" : {
"type" : " stdio " ,
"command" : " npx " ,
"args" : [ " -y " , " mindbase-mcp " ]
}
}
]
}
}
Restart VSCode. In Continue chat, MCP tools appear under @ . Skip to First project .
Any other MCP-compatible client
If your editor supports MCP (Zed, Aider, Goose, ChatGPT Desktop's future MCP support, custom Claude Agent SDK apps), the pattern is identical: point it at
npx -y mindbase-mcp
as a stdio server. See your client's MCP docs for the exact config location.
Once your editor has the MCP server connected, create your first project.
/mb:init my-research --mission "Study transformer attention mechanisms"
Claude will scaffold ~/mindbase-data/projects/my-research/ with README.md , context.md , index.yaml , and empty sources/ , logs/ , artifacts/ directories, and set it as your current project.
Cursor / Windsurf / Cline / Continue
Create a new mindbase project called my-research about transformer
attention mechanisms.
The LLM will call mindbase_init_project({ name: "my-research", mission: "..." }) and confirm.
/mb:status # in Claude Code
Or ask any other IDE:
what's the status of my mindbase?
You should see:
MindBase — Project: my-research
Location: ~/mindbase-data/projects/my-research
README 28 lines
context.md 8 lines (last built: never)
index.yaml 5 lines
Contributors: 0 users, 0 entries
Sources: 0 research, 0 raw
Logs: 1 day
Artifacts: 0 outputs
You're ready.
The four things you'll do every day.
Add a short observation, a decision, or a hot take. Goes straight into today's contributor file, no ceremony.
/mb:contribute "Attention is basically a soft dictionary lookup — each query
retrieves a weighted average of values based on similarity to keys."
Any other IDE:
Add to my mindbase: Attention is basically a soft dictionary lookup — each
query retrieves a weighted average of values based on similarity to keys.
Result: appends to ~/mindbase-data/projects/my-research/sources/contributors/<you>/2026-07-11.md and logs the operation.
2. Ingest a substantive source (paper, article, PDF)
For a real source, MindBase walks through Karpathy's 8-step ingest.
Claude Code (flagship — with sub-agent):
/mb:contribute /Downloads/attention-is-all-you-need.pdf
The contributor sub-agent will:
Come back to you with 3 key takeaways
Propose wh

[truncated]
