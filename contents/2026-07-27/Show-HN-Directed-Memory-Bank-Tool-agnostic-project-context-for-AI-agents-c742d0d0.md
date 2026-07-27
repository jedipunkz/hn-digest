---
source: "https://github.com/pmikutel/directed-memory-bank"
hn_url: "https://news.ycombinator.com/item?id=49072254"
title: "Show HN: Directed Memory Bank – Tool-agnostic project context for AI agents"
article_title: "GitHub - pmikutel/directed-memory-bank: Stop re-explaining your project every AI session. Tool-agnostic project context in plain markdown — works with Claude Code, Cursor, Codex, Gemini, any agent that reads files. · GitHub"
author: "pmikutel"
captured_at: "2026-07-27T17:31:16Z"
capture_tool: "hn-digest"
hn_id: 49072254
score: 1
comments: 0
posted_at: "2026-07-27T16:45:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Directed Memory Bank – Tool-agnostic project context for AI agents

- HN: [49072254](https://news.ycombinator.com/item?id=49072254)
- Source: [github.com](https://github.com/pmikutel/directed-memory-bank)
- Score: 1
- Comments: 0
- Posted: 2026-07-27T16:45:53Z

## Translation

タイトル: Show HN: Directed Memory Bank – AI エージェント向けのツールに依存しないプロジェクト コンテキスト
記事タイトル: GitHub - pmikutel/directed-memory-bank: AI セッションごとにプロジェクトを再説明するのはやめてください。プレーンなマークダウンでのツールに依存しないプロジェクト コンテキスト — クロード コード、カーソル、コーデックス、Gemini、ファイルを読み取るあらゆるエージェントで動作します。 · GitHub
説明: AI セッションごとにプロジェクトを再説明するのをやめます。プレーンなマークダウンでのツールに依存しないプロジェクト コンテキスト — クロード コード、カーソル、コーデックス、Gemini、ファイルを読み取るあらゆるエージェントで動作します。 - pmikutel/directed-memory-bank

記事本文:
GitHub - pmikutel/directed-memory-bank: AI セッションのたびにプロジェクトを再説明するのはやめてください。プレーンなマークダウンでのツールに依存しないプロジェクト コンテキスト — クロード コード、カーソル、コーデックス、Gemini、ファイルを読み取るあらゆるエージェントで動作します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
却下する

警告
{{ メッセージ }}
ピミクテル
/
ダイレクトメモリバンク
パブリックテンプレート
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 Commits 4 Commits cli cli docs docs examples examples template template CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MIGRATION.md MIGRATION.md README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
AI コーディング エージェントのための永続的なプロジェクトの理解。
AI エージェントは何も知らずにすべてのセッションを開始します。 Directed Memory Bank (DMB) は、構造化されたプロジェクトの知識と作業状態を、あらゆるツールで動作するプレーンなマークダウン ファイルで提供します。
Cline Memory Bank パターンからインスピレーションを得たもの。以下の「インスピレーション」を参照してください。
あなたは儀式を知っています。 Coming back after the weekend, you open a new AI session and burn the first ten minutes re-explaining the stack, the auth flow, the deprecation plan, why library X is off the table — before you can finally ask the actual question. DMB では、その儀式はなくなります。 AI はすでにプロジェクトを認識しています。ただ続けてください。
セッションは再開されますが、再開されません。新しいクロード コードまたはカーソル セッションを開くと、プロジェクト コンテキストがすでにスコープ内にあります。 AI は、役立つようになる前にコードベースに再導入する必要はありません。
ツールの切り替えには費用はかかりません。クロード コードでリファクタリングを開始し、カーソルで終了します。どちらも同じメモリバンク/ファイルを読み取ります。両方ともスタックと規約について同じ認識を持っています。
インターネットではなく、AI による推奨事項がチームに適合します。 Technical/stack.md が Postgres-as-cache で標準化していると示すと、モデルは Redis の推奨を停止します。 FastAPI にコミットすると、NestJS の提案が停止されます。一度下した決断はそのまま残ります。
知識は入れ替わりしても存続します。 W

チームメイトが去っても、システムに対する理解は残りません。 「なぜこのように構築したか」は、8 か月前の Slack スレッドではなく、決定事項の隣にあるtechnical/architecture.md に記載されています。
飛行中の作業のクロスセッション記憶。緊急のバグに対処するために機能を一時停止しました。 1 週間後、あなたはそれを再び手に取ります。AI はどこで中断したか、何がまだ残っているか、何が決定されたかを知っています。これは、tasks/work/<slug>.md がその仕事を行っていることです。
コンテキストの喪失 — セッションごとにプロジェクトを再説明するのはやめましょう。
AI の提案は的外れです。チームの実際の決定を無視する提案と戦うのはやめてください。
発見は消える — 集中して作業しているときに気づいたものを失うのを防ぎます。
ツールの断片化 — Claude、Cursor、Gemini 構成間での規則の重複を停止します。
意思決定の論理的根拠は失われます。「なぜこのようになったのか?」と考えるのはやめましょう。数か月後。
ファイル規約。サーバーもビルドステップもロックインもありません。 AI エージェントが必要なときに必要なものを見つけられるように、マークダウン ファイルを整理するだけです。
ツールに依存しない — Claude Code、Cursor、Gemini、Codex、またはファイルを読み取ることができる LLM ベースのエージェントで動作します。
必要なものを採用してください。ナレッジ層と RAM 層は独立して価値があります。フレームワーク全体を使用するか、問題を解決する部分だけを使用します。
2つの角度。まず、DMB を使用した場合と使用しない場合の 1 日の様子。 2 つ目 - DMB は、すでに使用しているコンテキスト ツールの隣にあります。
DMBなし
DMBあり
新しい AI セッションを開始する
スタック、認証フロー、制約について再説明します (5 ～ 15 分)
プロジェクトのコンテキストはすでにスコープ内にあります
タスクの途中でクロード コードからカーソルに切り替える
新しいツールでコンテキストを再構築する
どちらのツールも同じメモリバンク/ファイルを読み取ります
チームメイトが去る
彼らの理解は彼らに残ります
「なぜ」はtechnical/architecture.mdにあります
AIが図書館を提案する
W

訓練データで最も一般的な憎悪
実際に使用しているとtechnical/stack.mdに記載されているもの
ボットは PR ログに書き込みます
人間が共有ファイルに追加するとマージが競合する
ボットは独自のファイルを task/log/ に書き込みます。
一時停止した機能を 2 週間後に再開する
「待って、私は X について何を決めたのですか?」
task/work/<slug>.md を読んで続行します
DMB とすでに持っているものとの比較
単一の CLAUDE.md / カーソル ルール / AGENTS.md
ツールの内蔵自動メモリ
MCP メモリ サーバー (mem0、Letta、…)
DMB
AI ツール全体で機能する
いいえ — 1 つのツールのみ
いいえ — ユーザーごと、ツールごと
部分的 — MCP サポートに依存します
はい - ファイルを読み取るすべてのツール
バージョン管理下にある
はい
いいえ
いいえ (サーバー内)
はい
1 つの大きなファイルを超えて拡張可能
いいえ — 膨満します
該当なし
はい
はい - 変化率で分割
操作スコープのロード
いいえ
いいえ
ツールに依存する
はい — 統合レイヤーはインテントごとにルートします
マルチライターセーフ (人間 + ボット)
いいえ — 競合を追加します
いいえ
限定
はい — アイテムごとのファイル
ランタイム依存性なし
はい
はい (ツールのみ)
いいえ - サーバーが必要です
はい
DMB はこれらのいずれにも代わるものではありません。薄い CLAUDE.md または Cursor ルールは DMB を参照します。ツールの自動メモリは個人の好みをキャプチャします。 MCP メモリ サーバーを実行している場合、DMB ファイルはそれを通じて公開される可能性があります。
隣接: ナレッジベース Wiki
Karpathy 氏は最近、LLM Wiki パターンについて説明しました。同じ知的系統 (ファイルベースのマークダウン、メンテナとしての LLM、Memex スタイルの連想知識) が知識の蓄積に適用されます。研究ノート、書籍のコンパニオン、デュー デリジェンス、Slack によって提供されるチーム Wiki、および会議の記録です。
DMB は、サブストレートをコード プロジェクト コンテキストに適用します。問題が異なれば、設計の選択肢も異なります。コードベースが真実のソースです (取り込む「生のソース」レイヤーはありません)。単一の追加専用 log.md の代わりに PR ごとのファイルを作成することで、人間と CI ボットがコンテキストを書き込むことができます。

現在;別個の統合レイヤーにより、Claude Code、Cursor、Gemini などで同じメモリバンクが機能します。ナレッジと RAM が分割されるため、実行中の作業がセッション間で持続します。
競合するのではなく、補完的です。実際のチームは両方を実行できます。カルパシーの形をした研究/スパイク用の Wiki。出荷元のリポジトリの DMB。スパイクが終了すると、ロックインされた技術上の決定が Wiki から Memory-bank/technical/stack.md に移動します。
いくつかの傾向により、2026 年にはファイルベースのコンテキスト レイヤーが真に重要になります。
コンテキスト プロトコルは収束しています。クロード コード スキル、カーソル ルール、コーデックス エージェント スキル、Gemini CLI フック、MCP サーバーはすべて、読み取り可能なプロジェクト ナレッジ レイヤーが存在することを前提としています。 DMB は、あらゆるツールが理解できるものを提供します。
AI チームはマルチライターです。人間とオートメーションがプロジェクトの状態 (進捗ログ、PR ステータス、可観測性フィードバック) を書き込みます。この標準は、マージ競合が絶えず発生することなく、複数のライターが共存できるように設計されています。
Agentic CI が台頭しています — レビューアーとしての LLM、仕様主導型エンジニアリング、自律型トリアージ ボット。すべてのユーザーは、構造化され、クエリ可能なプロジェクト コンテキストを読み取り、慎重に書き込むことができるというメリットを享受できます。
標準はシンプルなままです (マークダウン、ツールなし)。人間と自動化が出会う継ぎ目に細心の注意を払うだけです。
npx ダイレクトメモリバンク初期化
ノードがありませんか?手動でクローンを作成してコピーします。
git clone https://github.com/pmikutel/directed-memory-bank.git
cp -r direct-memory-bank/template/memory-bank/ your-project/memory-bank/
プロジェクトで選択した AI ツール (Claude Code / Cursor / Codex / Gemini) を開き、次のように言います。
「memory-bank/INSTALL.md に従って DMB を設定してください。」
質問に答える。準備ができていないものはスキップし、可能な場合は入力する代わりにドキュメントを貼り付けます。進捗状況は、memory-bank/_adoption.md で追跡されます。いつでも戻ってきて、「con」と言ってください。

引き続き DMB インストールを実行してください。」; AI は中断したところから再開します。
それだけです。詳細については、 docs/ を参照してください。
_index.md は、memory-bank/ の内容と、開発、デバッグ、展開などの各種作業に重要なファイルをカタログ化します。これは 3 つの役割を果たします。
インベントリ — DMB を閲覧する人間と AI エージェントは、インベントリを読み取って移動します。
統合レイヤーの信頼できる情報源の作成 — あなた (またはあなたの代わりに動作する AI エージェント) がクロード コード スキル、カーソル ルール、フック、またはその他のツールごとの統合アーティファクトを作成または更新するときは、最初に _index.md を読み取り、操作にどのメモリ バンク ファイルが重要であるかを確認します。次に、統合層ファイルはそれらのメモリバンク ファイルを直接参照します。
フォールバック ランタイム ルーター — スキル、フック、またはパス スコープのルール (Codex、Gemini CLI、カスタム エージェント) を持たないツールの場合、_index.md をランタイム ルーティング テーブルとしてロードできます。
強力なハーネス メカニズムを備えたツールの場合、ルーティングが行われるのは統合レイヤーです。スキル、フック、ルールはメモリ バンク ファイルに直接名前を付けます。 _index.md は、実行時にそれらと知識の間に立ち入ることなく、一貫性を保ちます。
メモリバンク内の 2 つの層/
レイヤー
何
変更点
オプションですか？
知識
アーキテクチャ、ドメイン、規約、スタックの決定
ゆっくりと - プロジェクトが進化するたびに
コア
RAM
現在の焦点、進行中のトピック、完了した作業ログ
セッションごとに
はい
プロジェクト ドキュメントの対象読者は 3 名
DMB は、ドキュメントの 3 つの保存先のうちの 1 つです。分割:
解決テスト: 「このリポジトリでコードの作成、レビュー、デバッグ中に AI エージェントにこのファイルをロードさせたいですか?」はい → メモリバンク。いいえ → 外部ドキュメント。不明 → メモリバンク (デフォルト)。
完全なガイド: docs/what-goes-where.md 。
DMB は、リポジトリ内の 2 つの宛先に存在します。
さらに、個人的な設定を保持するツールの自動メモリ (terse-vs-ve)

rbose、ワークフロー調整) — ツールによって管理され、コミットされることはなく、DMB の範囲外です。
クロードコード カーソルコーデックス / ジェミニ / 他
│ │ │
▼ ▼ ▼
┌───────────────────┐
│ 統合ファイル (ツールごと、リポジトリ内) │
│ .claude/skills/, .cursor/rules/, │
│ クロード.md、ジェミニ.md、エージェント.md │
━━━━━━━━━━━━━━━━━━┘
│ 内のファイルを参照します
▼
┌───────────────────┐
│ メモリバンク/ (リポジトリ内、ツールに依存しない) │
│ │
│ ┌─────────┐ ┌───────────┐ │
│ │ 知識 │ │ RAM │ │
│ │ プロジェクト/ │ │ タスク/仕事/ │ │
│ │ 技術/ │ │ タスク/ログ/ │ │
│ ━━━━━━━┘ ━━━━━━━━┘ │
│ │
│ _index.md — ここにあるものの一覧表 │
━━━━━━━━━━━━━━━━━━━┘
AI ツールは変わらず上位に表示されます。ツールごとの統合ファイルは、ツールごとに作成する唯一のものであり、薄いポインターです。 Memory-bank/ は、すべてのツールが読み取る共有ストアです。
統合レイヤーの原則: 現在機能するものを使用する。最新の AI ツールはどれも、コンテキスト (命令ファイル、パススコープのルール、スキル、フック、MCP サーバーなど) にフックするための複数のメカニズムを提供します。各ツールについて、現在最も信頼性の高いメカニズムを使用してください。 D

信頼性の低いものをリスクの低いルーティングに移します。ツールの進化に合わせてバランスを再調整します。記憶禁止

[切り捨てられた]

## Original Extract

Stop re-explaining your project every AI session. Tool-agnostic project context in plain markdown — works with Claude Code, Cursor, Codex, Gemini, any agent that reads files. - pmikutel/directed-memory-bank

GitHub - pmikutel/directed-memory-bank: Stop re-explaining your project every AI session. Tool-agnostic project context in plain markdown — works with Claude Code, Cursor, Codex, Gemini, any agent that reads files. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
pmikutel
/
directed-memory-bank
Public template
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits cli cli docs docs examples examples template template CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MIGRATION.md MIGRATION.md README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Persistent project understanding for AI coding agents.
Your AI agent starts every session knowing nothing. Directed Memory Bank (DMB) gives it structured project knowledge and working state — in plain markdown files that work with any tool.
Inspired by the Cline Memory Bank pattern . See Inspiration below.
You know the ritual. Coming back after the weekend, you open a new AI session and burn the first ten minutes re-explaining the stack, the auth flow, the deprecation plan, why library X is off the table — before you can finally ask the actual question. With DMB, that ritual is gone. The AI already knows the project. You just continue.
Sessions resume, they don't restart. Open a new Claude Code or Cursor session and your project context is already in scope. The AI doesn't need to be re-introduced to the codebase before it can help.
Switching tools costs nothing. Start a refactor in Claude Code, finish it in Cursor. Both read the same memory-bank/ files; both are on the same page about your stack and your conventions.
AI recommendations match your team, not the internet. The model stops suggesting Redis when technical/stack.md says you've standardised on Postgres-as-cache. Stops suggesting NestJS when you've committed to FastAPI. Decisions you made once stay made.
Knowledge survives turnover. When a teammate leaves, their understanding of the system doesn't leave with them. The "why we built it this way" lives in technical/architecture.md next to the decision, not in a Slack thread from eight months ago.
Cross-session memory of in-flight work. You paused a feature to handle an urgent bug. A week later you pick it back up — the AI knows where you left off, what's still open, what was decided. That's tasks/work/<slug>.md doing its job.
Context loss — stop re-explaining your project every session.
AI suggestions miss the mark — stop fighting recommendations that ignore your team's actual decisions.
Discoveries vanish — stop losing things you noticed during focused work.
Tool fragmentation — stop duplicating conventions across Claude, Cursor, Gemini configs.
Decision rationale decays — stop wondering "why was this done this way?" months later.
A file convention. No server, no build step, no lock-in. Just markdown files organised so AI agents find what they need, when they need it.
Tool-agnostic — works with Claude Code, Cursor, Gemini, Codex, or any LLM-based agent that can read files.
Adopt what you need — the knowledge layer and RAM layer are independently valuable. Use the full framework or just the parts that solve your problem.
Two angles. First — what your day looks like with DMB vs without. Second — where DMB sits next to context tools you might already use.
Without DMB
With DMB
Start a new AI session
Re-explain the stack, the auth flow, the constraints (5–15 min)
Project context already in scope
Switch from Claude Code to Cursor mid-task
Rebuild context in the new tool
Both tools read the same memory-bank/ files
A teammate leaves
Their understanding leaves with them
The "why" lives in technical/architecture.md
AI suggests a library
Whatever's most common in training data
What technical/stack.md says you actually use
Bot writes to a PR log
Merge conflicts with humans appending to a shared file
Bot writes its own file in tasks/log/
Pick up a paused feature two weeks later
"Wait, what did I decide about X?"
Read tasks/work/<slug>.md and continue
DMB vs the things you already have
Single CLAUDE.md / Cursor rules / AGENTS.md
Tool's built-in auto-memory
MCP memory servers (mem0, Letta, …)
DMB
Works across AI tools
No — one tool only
No — per-user, per-tool
Partial — depends on MCP support
Yes — every tool that reads files
Lives in version control
Yes
No
No (in server)
Yes
Scales past one big file
No — gets bloated
N/A
Yes
Yes — split by rate of change
Operation-scoped loading
No
No
Tool-dependent
Yes — integration layer routes by intent
Multi-writer safe (humans + bots)
No — append conflicts
No
Limited
Yes — file-per-item
No runtime dependency
Yes
Yes (tool only)
No — needs server
Yes
DMB doesn't replace any of these. It composes with them: a thin CLAUDE.md or Cursor rule references DMB; the tool's auto-memory still captures personal preferences; if you run an MCP memory server, your DMB files can be exposed through it.
Adjacent: knowledge-base wikis
Karpathy recently described an LLM Wiki pattern — same intellectual lineage (file-based markdown, LLM as maintainer, Memex-style associative knowledge) applied to knowledge accumulation : research notes, book companions, due diligence, team wikis fed by Slack and meeting transcripts.
DMB applies the substrate to code project context . Different problem, different design choices: the codebase is the source of truth (no "raw sources" layer to ingest); file-per-PR log instead of single append-only log.md so humans + CI bots can write concurrently; a separate integration layer so the same memory-bank/ works across Claude Code, Cursor, Gemini and others; a knowledge/RAM split so in-flight work persists across sessions.
Complementary, not competing. A real team can run both — Karpathy-shaped wiki for research/spikes; DMB for the repo they ship from. When a spike concludes, the locked-in tech decision graduates from the wiki into memory-bank/technical/stack.md .
A few trends make a file-based context layer genuinely important in 2026:
Context protocols are converging — Claude Code skills, Cursor rules, Codex Agent Skills, Gemini CLI hooks, MCP servers all assume a readable project knowledge layer exists. DMB gives you one that every tool understands.
AI teams are multi-writer — humans and automations write project state now (progress logs, PR status, observability feedback). The standard is designed so multiple writers can coexist without constant merge conflicts.
Agentic CI is rising — LLM-as-reviewer, spec-driven engineering, autonomous triage bots. All of them benefit from a structured, queryable project context they can read and, carefully, write.
The standard stays simple (markdown, no tooling). It just pays close attention to the seams where humans and automations meet.
npx directed-memory-bank init
No Node? Clone and copy manually:
git clone https://github.com/pmikutel/directed-memory-bank.git
cp -r directed-memory-bank/template/memory-bank/ your-project/memory-bank/
Open your AI tool of choice (Claude Code / Cursor / Codex / Gemini) in your project, and say:
"Follow memory-bank/INSTALL.md to set up my DMB."
Answer the questions. Skip what you're not ready for, paste docs instead of typing where you can. Your progress is tracked in memory-bank/_adoption.md — come back any time and say "continue DMB install" ; your AI will pick up where you left off.
That's it. For more depth see docs/ .
_index.md catalogs what's in memory-bank/ and which files matter for each kind of work — development, debugging, deployment, etc. It plays three roles:
Inventory — humans and AI agents browsing DMB read it to navigate.
Authoring source-of-truth for the integration layer — when you (or an AI agent acting on your behalf) create or update a Claude Code skill, a Cursor rule, a hook, or any other per-tool integration artefact, you read _index.md first to see which memory-bank files matter for the operation. The integration-layer file then references those memory-bank files directly.
Fallback runtime router — for tools without skills, hooks, or path-scoped rules (Codex, Gemini CLI, custom agents), _index.md can be loaded as the runtime routing table.
For tools with strong harness mechanisms, the integration layer is where routing happens — skills, hooks, and rules name the memory-bank files directly. _index.md keeps them coherent without standing between them and the knowledge at runtime.
Two layers inside memory-bank/
Layer
What
Changes
Optional?
Knowledge
Architecture, domain, conventions, stack decisions
Slowly — when project evolves
Core
RAM
Current focus, in-flight topics, completed-work log
Every session
Yes
Three audiences for project docs
DMB is one of three destinations for documentation. The split:
The resolving test: "Would I want an AI agent to load this file while writing / reviewing / debugging code in this repo?" Yes → memory-bank. No → external-docs. Unsure → memory-bank (the default).
Full guide: docs/what-goes-where.md .
DMB lives in two destinations inside your repo:
Plus your tool's auto-memory, which holds personal preferences (terse-vs-verbose, workflow tweaks) — managed by the tool, never committed, out of DMB's scope.
Claude Code Cursor Codex / Gemini / others
│ │ │
▼ ▼ ▼
┌──────────────────────────────────────────┐
│ Integration files (per-tool, in repo) │
│ .claude/skills/, .cursor/rules/, │
│ CLAUDE.md, GEMINI.md, AGENTS.md │
└──────────────────┬───────────────────────┘
│ references files in
▼
┌──────────────────────────────────────────┐
│ memory-bank/ (in repo, tool-agnostic) │
│ │
│ ┌──────────────┐ ┌────────────────────┐ │
│ │ Knowledge │ │ RAM │ │
│ │ project/ │ │ tasks/work/ │ │
│ │ technical/ │ │ tasks/log/ │ │
│ └──────────────┘ └────────────────────┘ │
│ │
│ _index.md — inventory of what's in here │
└──────────────────────────────────────────┘
Your AI tools stay unchanged at the top. Per-tool integration files are the only thing you write per-tool — and they're thin pointers. memory-bank/ is the shared store every tool reads.
The principle for the integration layer: use what works today . Every modern AI tool offers more than one mechanism for hooking into context (instruction files, path-scoped rules, skills, hooks, MCP servers, …). For each tool, lead with the mechanism that is most reliable today. Demote the less-reliable ones to lower-stakes routing. Rebalance as the tools evolve. memory-ban

[truncated]
