---
source: "https://github.com/tiny-cloud-ventures/amnesia"
hn_url: "https://news.ycombinator.com/item?id=48981446"
title: "Show HN: Amnesia – audit Claude Code's memory for contradictions"
article_title: "GitHub - tiny-cloud-ventures/amnesia: Give your agent selective amnesia — see, search, and clean Claude Code's memory across all your repos · GitHub"
author: "suttles"
captured_at: "2026-07-20T17:23:44Z"
capture_tool: "hn-digest"
hn_id: 48981446
score: 2
comments: 0
posted_at: "2026-07-20T16:53:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Amnesia – audit Claude Code's memory for contradictions

- HN: [48981446](https://news.ycombinator.com/item?id=48981446)
- Source: [github.com](https://github.com/tiny-cloud-ventures/amnesia)
- Score: 2
- Comments: 0
- Posted: 2026-07-20T16:53:51Z

## Translation

タイトル: Show HN: Amnesia – クロード・コードの記憶の矛盾を監査する
記事のタイトル: GitHub - tiny-cloud-ventures/amnesia: エージェントに選択的記憶喪失を与える — すべてのリポジトリにわたってクロード コードのメモリを参照、検索、および消去する · GitHub
説明: エージェントに選択的記憶喪失を与えます — すべてのリポジトリでクロード コードのメモリを参照、検索、および消去します - tiny-cloud-ventures/amnesia

記事本文:
GitHub - tiny-cloud-ventures/amnesia: エージェントに選択的記憶喪失を与えます — すべてのリポジトリにわたってクロード コードのメモリを参照、検索、および消去します · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。解放してください

このページをご覧ください。
小さなクラウドベンチャー
/
健忘症
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット .claude-plugin .claude-plugin コマンド コマンド docs docs .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md amnesia.py amnesia.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントに選択的記憶喪失を与えます。クロード コードが保存したあなたに関するすべての記憶を見て、検索し、消去します。そして、クロードに自身の記憶の矛盾を監査させます。
カウントはライブであり、固定されていません。 Amnesia は、~/.claude/projects/*/memory/ に保存されているすべてのメモリを検出し、ストアの実際の合計を報告します。以下のスクリーンショットは 1 つの実際の店舗であり、制限ではありません。
すべての点は 1 つの記憶であり、色はプロジェクトであり、線はそれらの間の参照です。スキャン後、矛盾は赤く光ります。組み込み UI でライブ、ドラッグ、クリック可能。
Claude Code は、 ~/.claude/projects/*/memory/ の下にプロジェクト ディレクトリごとにメモリ ファイルを静かに蓄積します。何か月も経つと腐敗が蓄積され、事実は古くなり（「サービス X は稼働中」 - 数週間前に廃止されました）、同じルールが 3 つのプロジェクトで 3 回保存され、あるプロジェクトの記憶が別のプロジェクトのセッションに漏洩します。エージェントの記憶に関する研究では、これを 記憶汚染 と呼び、記憶が汚染されていると、エージェントの状態はまったく記憶がない場合よりも明らかに悪化します。
記憶喪失は 1 つの文、1 つのボタン、一度に 1 つの質問です。
開いてみてください。Amnesia はあなたのストアをライブでカウントし、見つかった思い出とプロジェクトの数を正確に教えてくれます。ボタン 1 つ: スキャン 。
スキャンはストア全体を独自の claude CLI にフィードし、矛盾、古い事実、重複にフラグを立てます。

食べたもの、そして間違ってファイルされた記憶。
レビューでは、わかりやすい言葉で一度に 1 つずつフラグを確認し、「これらは両方とも当てはまらない」「これは間違ったプロジェクトに登録されているようだ」と、一言で答えます。忘れる、移動する、結合する、両方保持する。決定は、次のスキャンが成功するまでリロードされても存続し、新しい証拠から再度質問されます。お急ぎですか？確信を持って統合するたびにすべてのプレビューを修正し、それらを 1 つの取り消し可能なバッチとして適用します。
すべての変更は元に戻すことができます。Amnesia は、影響を受けるメモリとインデックスを変更する前にスナップショットを作成します。元に戻すは、トーストの [最近の変更] の下にあります。後で編集したファイルの上書きは拒否されます。
検索、プロジェクト フィルターを備えたすべてのビュー、日付/サイズの並べ替え、およびマップ (依存関係のない最大 100 行のキャンバスに手動で展開された物理学上のライブ フォース グラフ) は、それぞれ 1 クリックで表示され、デフォルトではありません。
Claude Code プラグインとして、必要なものはすべてすでに揃っています。
/plugin マーケットプレイスに tiny-cloud-ventures/amnesia を追加
/プラグインインストールamnesia@amnesia
次に、/amnesia:open は UI を開き、/amnesia:scan はセッションを離れることなく記憶を監査します。
CLI として (Python 3.9 以降が必要):
uvx --from git+https://github.com/tiny-cloud-ventures/amnesiaアムネジア
または、ツールを使用しない方法 — 単一のファイル (stdlib のみ) です。
カール -O https://raw.githubusercontent.com/tiny-cloud-ventures/amnesia/main/amnesia.py
python3 amnesia.py
使用する
amnesia # UI (http://localhost:8780) — スキャン、レビュー、完了
記憶喪失分析 # 同じ監査、端末から
amnesia 9000 # 別のポート
または、プラグインを使用: /amnesia:open および /amnesia:scan 。
マシンからは何も残りません。 UI はローカルホストのみにバインドされます。アカウント、API キー、テレメトリ、サーバーはありません。変更には、同じセッション、同じオリジンの JSON リクエストが必要です。
あなたのメモを読むのは唯一の人

ries は、すでに使用している Claude アカウントです。スキャンは、独自の claude CLI (Claude Code が実行されるのと同じサブスクリプション) を介してストアにパイプします。
何も削除されることはありません。忘れられてマージされたソース ファイルは ~/.claude/memory-trash/ に移動します。すべての忘却、移動、マージ、およびすべて修正のバッチでは、 ~/.claude/amnesia/history/ の下にローカル回復スナップショットも取得されます。
スコープのない古いメモリは負の値です。古いポート番号や廃止されたサービスを「記憶」しているエージェントは、自信を持ってその番号を今日の作業に適用します。修正は、より多くのことを記憶することではなく、忘れて正しくフェンシングすることです。 amnesia はそのための最小の誠実なツールです。完全な可視性、安全な削除、ストアの内部一貫性の LLM 監査です。
アナライザーは、統合操作も提案します。つまり、メモリを実際に属しているプロジェクト ディレクトリに移動し、複製を正規コピーにマージします。それぞれはワンクリックで適用され、ゴミ箱に戻されます。人間の承認を得たリポジトリ間の統合が中心的なアイデアです。すべての主流メモリ システム (ChatGPT、Claude.ai、Mem0、Zep) はサイレントに統合されます。 none を使用すると、マージを確認できます。この設計の基礎となっている研究については、 docs/ を参照してください。
アーカイブ状態 (メモリをソフト無効にし、削除する前に何か破損していないか確認します)
Supersede-don't-delete: 古いファクトを削除せずに、superseded: マーカーを使用して再書き込みします。
Gated PROMOTE オペレーション: 2 つ以上のリポジトリで観察された場合にのみ、出所とともにファクトをグローバル スコープにホイストします
より多くのエージェントに対するメモリのサポート (カーソル ルール、プレーン CLAUDE.md ファイル)
問題や PR は歓迎です — COTRIBUTING.md を参照してください。何よりも重要なルールは、依存関係がゼロであることが製品であるということです。
エージェントに選択的記憶喪失を与えます — すべてのリポジトリでクロード コードのメモリを参照、検索、消去します
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
リポ

rtリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Give your agent selective amnesia — see, search, and clean Claude Code's memory across all your repos - tiny-cloud-ventures/amnesia

GitHub - tiny-cloud-ventures/amnesia: Give your agent selective amnesia — see, search, and clean Claude Code's memory across all your repos · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
tiny-cloud-ventures
/
amnesia
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .claude-plugin .claude-plugin commands commands docs docs .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md amnesia.py amnesia.py pyproject.toml pyproject.toml View all files Repository files navigation
Give your agent selective amnesia. See, search, and clean every memory Claude Code has saved about you — and let Claude audit its own memory for contradictions.
The counts are live, not fixed. Amnesia discovers every saved memory under ~/.claude/projects/*/memory/ and reports the actual total for your store. The screenshot below is one real store, not a limit.
Every dot is one memory, colors are projects, lines are the references between them — contradictions glow red after a scan. Live, draggable, and clickable in the built-in UI.
Claude Code quietly accumulates memory files per project directory under ~/.claude/projects/*/memory/ . Over months that store rots: facts go stale ("service X is live" — it was retired weeks ago), the same rule gets saved three times in three projects, and memories from one project leak into sessions for another. Research on agent memory calls this memory contamination , and polluted memory measurably makes agents worse than no memory at all.
amnesia is one sentence, one button, one question at a time:
Open it: Amnesia counts your store live and tells you exactly how many memories and projects it found. One button: Scan .
The scan feeds your whole store to your own claude CLI and flags contradictions , stale facts , duplicates , and misfiled memories .
Review walks you through the flags one at a time, in plain language — "These can't both be true" , "This seems filed in the wrong project" — with one-word answers: Forget, Move it, Combine them, Keep both. Decisions survive reloads until the next successful scan, which asks again from fresh evidence. In a hurry? Fix all previews every confident consolidation, then applies them as one undoable batch.
Every change is reversible: Amnesia snapshots each affected memory and index before changing it. Undo is right there in the toast and under Recent changes ; it refuses to overwrite a file you edited afterward.
Search, a browse-everything view with project filters, date/size sorting, and the map — the live force-graph above, physics hand-rolled in ~100 lines of dependency-free canvas — are each one click away, never the default.
As a Claude Code plugin — you already have everything it needs:
/plugin marketplace add tiny-cloud-ventures/amnesia
/plugin install amnesia@amnesia
Then /amnesia:open opens the UI, and /amnesia:scan audits your memory without leaving the session.
As a CLI (requires Python 3.9+):
uvx --from git+https://github.com/tiny-cloud-ventures/amnesia amnesia
Or the zero-tooling way — it's a single file, stdlib only:
curl -O https://raw.githubusercontent.com/tiny-cloud-ventures/amnesia/main/amnesia.py
python3 amnesia.py
Use
amnesia # UI at http://localhost:8780 — scan, review, done
amnesia analyze # same audit, from the terminal
amnesia 9000 # different port
Or, with the plugin: /amnesia:open and /amnesia:scan .
Nothing leaves your machine. The UI binds to localhost only. No accounts, no API keys, no telemetry, no server. Changes require a same-session, same-origin JSON request.
The only thing that ever reads your memories is the Claude account you already use. The scan pipes your store through your own claude CLI — the same subscription Claude Code runs on.
Nothing is ever deleted. Forgotten and merged source files move to ~/.claude/memory-trash/ ; every forget, move, merge, and Fix all batch also gets a local recovery snapshot under ~/.claude/amnesia/history/ .
Unscoped, stale memory is negative-value: an agent that "remembers" your old port number or a retired service confidently applies it to today's work. The fix isn't remembering more — it's forgetting and fencing correctly. amnesia is the smallest honest tool for that: full visibility, safe deletion, and an LLM audit of the store's internal consistency.
The analyzer also proposes consolidation ops — MOVE a memory to the project directory it actually belongs to, MERGE a duplicate into its canonical copy — each applied with one click, trash-backed. Cross-repo consolidation with human approval is the core idea: every mainstream memory system (ChatGPT, Claude.ai, Mem0, Zep) consolidates silently; none lets you review the merge. See docs/ for the research this design is based on.
Archive state (soft-disable a memory and see if anything breaks before deleting)
Supersede-don't-delete: rewrite stale facts in place with a superseded: marker instead of trashing
Gated PROMOTE op: hoist a fact to global scope only when observed in 2+ repos, with provenance
Memory support for more agents (Cursor rules, plain CLAUDE.md files)
Issues and PRs welcome — see CONTRIBUTING.md . One rule above all: zero dependencies is the product.
Give your agent selective amnesia — see, search, and clean Claude Code's memory across all your repos
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
