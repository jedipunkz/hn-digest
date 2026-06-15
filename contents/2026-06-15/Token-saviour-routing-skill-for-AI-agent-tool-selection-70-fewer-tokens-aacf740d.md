---
source: "https://github.com/vagkaratzas/skills/blob/main/token-saviour/SKILL.md"
hn_url: "https://news.ycombinator.com/item?id=48542411"
title: "Token-saviour – routing skill for AI agent tool selection (~70% fewer tokens)"
article_title: "skills/token-saviour/SKILL.md at main · vagkaratzas/skills · GitHub"
author: "vagkaratzas"
captured_at: "2026-06-15T16:43:20Z"
capture_tool: "hn-digest"
hn_id: 48542411
score: 3
comments: 0
posted_at: "2026-06-15T15:08:19Z"
tags:
  - hacker-news
  - translated
---

# Token-saviour – routing skill for AI agent tool selection (~70% fewer tokens)

- HN: [48542411](https://news.ycombinator.com/item?id=48542411)
- Source: [github.com](https://github.com/vagkaratzas/skills/blob/main/token-saviour/SKILL.md)
- Score: 3
- Comments: 0
- Posted: 2026-06-15T15:08:19Z

## Translation

タイトル: トークン救世主 – AI エージェント ツール選択のためのルーティング スキル (トークンの最大 70% 削減)
記事タイトル：skills/token-saviour/SKILL.md at main · vagkaratzas/skills · GitHub
説明: 日々のプログラミングに使用されるエージェント スキル。 GitHub でアカウントを作成して、vagkaratzas/スキル開発に貢献してください。

記事本文:
メインのスキル/トークンセイバー/SKILL.md · vagkaratzas/skills · GitHub
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
ヴァグカラツァス
/
スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする 責任を負う さらに多くのファイル アクションを責任を負う

その他のファイルアクション 最新のコミット
履歴 履歴 151 行 (118 loc) · 8.13 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション
名前
トークンセイバー
説明
ファイル全体を反射的に読み取ったり、生のコマンド出力をコンテキストにダンプしたりするのではなく、コーディング タスクに最もトークン効率の高いツールを選択してください。このスキルは、コードベースの探索や説明、シンボル/定義/呼び出し元の特定、呼び出しパスのトレース、アーキテクチャのマッピング、複数のレイヤーにわたる機能の計画、または冗長コマンド (テスト、ビルド、git、grep、ディレクトリ一覧) の実行前に、コンテキスト/トークンの予算、コスト、または「トークンの使用量を減らす」という問題が発生したときに使用します。具体的なコマンドと、使用するか回避するかの組み合わせを使用して、シナリオを serena、graphify、rtk、caveman、または単純なツールにルーティングします。ユーザーがツールに名前を付けていない場合でも、それに手を伸ばしてください。質問に答えるためにいくつかのファイルを「猫」/読み取りしようとしている場合、それがトリガーになります。
トークンの救世主: 重要な場所にトークンを使う
狭い質問に答えるためにファイル全体を読むことは、エージェントにとって最も無駄なことです
そうです。約 30 ファイルの Python アプリのベンチマークで、ファイル全体の読み取りをセマンティックに交換する
取得により合計トークンが最大 66% 削減されました。他のツールはそれぞれ、より狭いスライスを所有します。このスキル
コンテキストの予算を使い果たす前に、適切なものに到達し、その後、元の状態に戻るのに役立ちます。
実際のタスク。
メンタル モデル: トークン コストには 4 つの独立したレイヤーがあり、異なるツールがそれぞれを所有します。
タスクが実際に重点を置くレイヤーにツールを合わせます。
まず、空き状況を確認します。これらはオプションのサードパーティ ツールです。関連するものを実行します
--help / --version 1 回;ツールがインストールされていない場合は、次善のオプションにフォールバックします。
その行 (最終的にはプレーンな Read/Grep/Bash)。ツールが実行されたふりをしないでください。正常に機能を低下させます。
仕事

トップダウン。行が一致した瞬間に、そのツールを使用して停止します。
「X はどこですか / 何が X を呼び出しますか / A はどのようにして B に到達しますか / X を変更すると何が壊れますか?」
→ グラフ化 (クエリ可能なコード グラフ)。ナビゲーションとインパクトで輝きます:発信者、
コールパストレース、クロスモジュールフロー。測定値: 4 層コール パスのトレース コスト 65
トークン対ファイル読み取り 1,633 個 (-89%)。
「このモジュール/クラスを説明し、そのメソッドをリストし、アーキテクチャを要約します」 — または、あなたは
シンボル → セレナ (ライブ LSP シンボル + セマンティック編集) によってコードを編集しようとしています。それ
幅広い理解と編集に優れています: 多くのファイルにわたるシンボルの概要。
グラフ化ではできないリファクタリングや名前の変更。測定済み: アーキテクチャ全体のコストを説明する
243 トークン vs 3,250 (-85%); 1 つのクラスのメソッド、55 対 458 (-71%)。
ノイズの多いものを実行する — テスト スイート、ビルド、リンター、 git 、 grep 、大きなディレクトリ
リスト → rtk (コマンド出力を圧縮)。測定値: テスト実行 -65%、
構造リスト -38% 。勝利は冗長性によって決まるため、テストの不合格時に最大の効果が得られます
ダンプと長いビルド ログ。
長くて散文の多い回答 (説明、追記、状況) を書こうとしている → 検討する
穴居人 (簡潔な「穴居人」応答スタイル)。フィラー/ヘッジ/プレサントリーをトリミングします。小さい
すでに簡潔なテキストですが、マークダウンを追加することで簡潔なリスト/構造化された出力を増やすことができます。
短い事実に基づく回答や正確なものではなく、本当におしゃべりな出力に使用してください。
文言/順序に関する事項 (警告、取り消し不能な手順、順序付けられた手順)。
上記のどれでもない / 1 ファイル、1 行のルックアップ → 単純な Read/Grep/Bash を使用してください。の
ツールにはセットアップと呼び出しのオーバーヘッドがあります。小さなタスクではオーバーヘッドの方がコストが高くなる可能性があります
節約します。つまらない作業をツール化しないでください。
セレナ vs グラフィファイ — 重複しているため、どちらかを選択してください
どちらも同じレイヤー (コード読み取り入力) をカットするため、両方をインストールすることは

ほとんどが冗長 — 選択基準
エージェントの主な動作方法:
作業がナビゲーションに重点を置いている場合はグラフ化します。「誰がこれを呼び出しているか」、呼び出しパス、衝撃/爆発
半径、クロスリポジトリ構造。読み取り専用。 1 回限りのグラフの構築とその後の再構築が必要
コードが変更されます。
作業が理解と編集に重点を置いている場合は serena : 複数ファイルの広範な理解
さらに、セマンティックな編集/名前変更 (graphify は編集できません)。言語サーバーが必要です。インデックスを一度に作成する
セッション開始。
ツールが 1 つしかない場合は、それをコード検索ツール (serena または grificify) にします。
これは、主要な (コード読み取り) コストを動かす唯一のレバーです。
最適な組み合わせと避けるべきもの
レイヤーは独立しているため、レイヤーごとに 1 つを結合すると加算的になります。
🥇 推奨スタック: (セレナまたはグラフィファイ) + rtk + 穴居人。それぞれが所有している
個別のスライス。 rtk と穴居人のコストは、追加するものは何もありません。合計測定値: -70% 。
❌ Serena とgraphify を一緒に実行しないでください。2 つの統合で 1 つのメリットが得られます。 1 つ選んでください。
❌ 理解コストを修正するために rtk を使用しないでください。コストは ~0% です。用途:
コマンドを多用するループ。
❌ 穴居人が入力の多いコードベースの作業に関して法案を動かすことを期待しないでください。それは縮小するだけです
大量のコードを読んでいるとき、これは小さなシェアです。
ツールのクイックリファレンス (コマンド)
これらを再帰的な Read / Grep / Bash の代わりに使用してください。
serena (MCP サーバー、LSP) — コードを理解するためにファイルを読む代わりに
プロジェクトの serena MCP サーバーに接続し、ファイルを読み取るのではなくツールを呼び出します。
get_symbols_overview(relative_path) — ファイル内のトップレベルのシンボル (メソッドの場合は深さ 1 を追加)。
find_symbol(name_path_pattern,relative_path,include_body=true) — シンボルの本体。
find_referencing_symbols(name_path,relative_path) — すべての呼び出し元とコンテキスト。
また、 replace_symbol_body 、 insert_after_symbol も編集します。

、rename_symbol 。
ヘッドレス (MCP クライアントなし): serena start-mcp-server --project <dir> --transport stdio および
改行区切りの JSON-RPC を話します (最小限のクライアントについては、このリポジトリの Benchmark/serena_client.py を参照してください)。
graphify (CLI + コードグラフ) — ナビゲートするための grep/読み取りの代わりに
graphify update < dir > --no-cluster # グラフを一度構築します (ツリーシッター、LLM/API キーなし)
graphify Explain " ClassName " # ノード + そのメソッド + 近隣ノード
グラフ化クエリ「誰が通知を呼び出しますか?」 # 質問のグラフに対する BFS (--budget N)
graphify path " create_task " " Database " # 2 つのシンボル間の最短パス
影響を受けるグラフ化 " ClassName " # 逆走査: 変更の影響
グラフは <dir>/graphify-out/graph.json にあります。編集後にgraphify updateで再構築します。
rtk (Rust CLI プロキシ) — 生の詳細コマンドの代わり
rtk テスト < テスト cmd > # 例: rtk テスト pytest → 失敗のみ + まとめ
rtk read < files... > # フィルタリングを使用した cat
rtk grep < pattern > < p > # コンパクト grep (グループ、切り捨て); -t py でタイプ別にスコープを指定します
rtk find < find args > # コンパクトなファイルリスト
rtk ls / rtk ツリー / rtk git / rtk diff / rtk ログ ...
穴居人 — 長くおしゃべりな返事の代わりに
次の場合は、簡潔モード (「穴居人モード」 / /caveman 、レベル lite|full|ultra と言います) をアクティブにします。
冗長な説明を発しようとしています。コード ブロック、コマンド、API 名、エラー文字列、および
順序付けられた / 不可逆的な手順を逐語的に — 穴居人は内容ではなく散文を圧縮します。
これが機能する理由 (調整できるようにするため)
エージェントがコードベースで作業する場合、入力が請求額の大部分を占めます (ベンチマークでは最大 88% でした)。
その入力のほとんどはオーバーリードです。ファイル全体を取得して、シンボルの検索に応答します。
または、グラフ クエリの応答はトークンの一部になります。セマンティック検索が優れているのは、
関連する構造のみを返します。 rtkと穴居人

全体として小さく見えるのは、単に理由だけです
狭いスライスを対象としていますが、安いので積み重ねてください。 1 つの罠はこれらを扱うことです
交換可能: それぞれがそのレーン内では優れていますが、レーン外では役に立たないため、シナリオごとにルートを決定します
お気に入りに手を伸ばすのではなく。疑問がある場合は、「このタスクがどの層に費やしているのか」を尋ねてください。
トークンはオンですか？」そしてその行を選択します。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agentic skills used for my day-to-day programming. Contribute to vagkaratzas/skills development by creating an account on GitHub.

skills/token-saviour/SKILL.md at main · vagkaratzas/skills · GitHub
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
vagkaratzas
/
skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 151 lines (118 loc) · 8.13 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions
name
token-saviour
description
Pick the most token-efficient tool for a coding task instead of reflexively reading whole files or dumping raw command output into context. Use this skill BEFORE you explore or explain a codebase, locate a symbol/definition/callers, trace a call path, map an architecture, plan a feature across layers, or run verbose commands (tests, builds, git, grep, directory listings) — and whenever context/token budget, cost, or "make this use fewer tokens" comes up. It routes the scenario to serena, graphify, rtk, caveman, or plain tools, with concrete commands and the combinations to use vs. avoid. Reach for it even when the user doesn't name a tool: if you're about to `cat`/Read several files to answer a question, that's the trigger.
token-saviour: spend tokens where they matter
Reading whole files to answer a narrow question is the single most wasteful thing an agent
does. On a benchmark over a ~30-file Python app, swapping whole-file reads for semantic
retrieval cut total tokens ~66% ; the other tools each own a narrower slice. This skill
helps you reach for the right one before you blow the context budget — then get back to the
actual task.
The mental model: token cost has four independent layers , and a different tool owns each.
Match the tool to the layer the task actually stresses.
First, check availability. These are optional third-party tools. Run the relevant
--help / --version once; if a tool isn't installed, fall back to the next-best option in
its row (ultimately plain Read/Grep/Bash). Never pretend a tool ran — degrade gracefully.
Work top-down. The moment a row matches, use that tool and stop.
"Where is X / what calls X / how does A reach B / what breaks if I change X?"
→ graphify (a queryable code graph). It shines at navigation and impact : callers,
call-path tracing, cross-module flow. Measured: tracing a 4-layer call path cost 65
tokens vs 1,633 reading the files (−89%).
"Explain this module/class, list its methods, summarize the architecture" — or you're
about to edit code by symbol → serena (live LSP symbols + semantic edits). It
shines at broad comprehension and editing : symbol overviews across many files, and
refactors/renames that graphify can't do. Measured: explaining the whole architecture cost
243 tokens vs 3,250 (−85%); one class's methods, 55 vs 458 (−71%).
Running something noisy — test suite, build, linter, git , grep , a big directory
listing → rtk (compresses command output). Measured: the test run −65% , the
structure listing −38% . Its win scales with verbosity, so it's biggest on failing test
dumps and long build logs.
About to write a long, prose-heavy answer (explanations, write-ups, status) → consider
caveman (terse "caveman" reply style). It trims filler/hedging/pleasantries. Small on
already-concise text and it can grow terse list/structured output by adding markdown — so
use it for genuinely chatty output, not for short factual answers or anything where exact
wording/order matters (warnings, irreversible steps, ordered procedures).
None of the above / a one-file, one-line lookup → just use plain Read/Grep/Bash. The
tools have setup and call overhead; on tiny tasks that overhead can cost more than it
saves. Don't tool-ify trivial work.
serena vs graphify — they overlap, so pick ONE
Both cut the same layer (code-read input), so installing both is mostly redundant — choose by
how the agent mostly works:
graphify if the work is navigation-heavy : "who calls this", call paths, impact/blast
radius, cross-repo structure. Read-only; needs a one-time graph build and a rebuild after
code changes.
serena if the work is comprehension- and edit-heavy : broad multi-file understanding
plus semantic edits/renames (graphify can't edit). Needs a language server; indexes once at
session start.
If you can only have one tool at all, make it the code-retrieval one (serena or graphify) —
it's the only lever that moves the dominant (code-reading) cost.
Best combination, and what to avoid
The layers are independent, so combining one-per-layer is additive:
🥇 Recommended stack: (serena or graphify) + rtk + caveman . Each owns a
distinct slice; rtk and caveman cost ~nothing to add. Measured combined: −70% .
❌ Don't run serena and graphify together — two integrations, one benefit. Pick one.
❌ Don't reach for rtk to fix comprehension cost — it's ~0% there. Use it for
command-heavy loops.
❌ Don't expect caveman to move the bill on input-heavy codebase work — it only shrinks
your output, which is a small share when you're reading lots of code.
Tool quick-reference (commands)
Use these in place of the reflexive Read / Grep / Bash equivalents.
serena (MCP server, LSP) — instead of reading files to understand code
Connect to the serena MCP server for the project, then call tools rather than reading files:
get_symbols_overview(relative_path) — top-level symbols in a file (add depth: 1 for methods).
find_symbol(name_path_pattern, relative_path, include_body=true) — a symbol's body.
find_referencing_symbols(name_path, relative_path) — every caller, with context.
It also edits: replace_symbol_body , insert_after_symbol , rename_symbol .
Headless (no MCP client): serena start-mcp-server --project <dir> --transport stdio and
speak newline-delimited JSON-RPC (see benchmark/serena_client.py in this repo for a minimal client).
graphify (CLI + code graph) — instead of grepping/reading to navigate
graphify update < dir > --no-cluster # build the graph once (tree-sitter, no LLM/API key)
graphify explain " ClassName " # a node + its methods + neighbors
graphify query " who calls notify? " # BFS over the graph for a question (--budget N)
graphify path " create_task " " Database " # shortest path between two symbols
graphify affected " ClassName " # reverse traversal: impact of a change
Graph lives at <dir>/graphify-out/graph.json ; rebuild with graphify update after edits.
rtk (Rust CLI proxy) — instead of raw verbose commands
rtk test < test cmd > # e.g. rtk test pytest → only failures + summary
rtk read < files... > # cat with filtering
rtk grep < pattern > < p > # compact grep (groups, truncates); -t py to scope by type
rtk find < find args > # compact file listing
rtk ls / rtk tree / rtk git / rtk diff / rtk log ...
caveman — instead of a long, chatty reply
Activate terse mode (say "caveman mode" / /caveman , levels lite|full|ultra ) when you're
about to emit a verbose explanation. Keep code blocks, commands, API names, error strings, and
ordered/irreversible steps verbatim — caveman compresses prose, not substance.
Why this works (so you can adapt it)
Input dominates the bill when an agent works in a codebase (it was ~88% in the benchmark), and
most of that input is over-reading — pulling entire files to answer something a symbol lookup
or a graph query answers in a fraction of the tokens. Semantic retrieval wins because it
returns just the relevant structure . rtk and caveman look small in aggregate only because
they target narrow slices — but they're cheap, so stack them. The one trap is treating these
as interchangeable: each is excellent in its lane and ~useless outside it, so route by scenario
rather than reaching for a favorite. When in doubt, ask "which layer is this task spending
tokens on?" and pick that row.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
