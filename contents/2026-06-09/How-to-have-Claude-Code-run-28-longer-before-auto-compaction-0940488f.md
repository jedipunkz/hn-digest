---
source: "https://github.com/wiztek-llc/context-ledger/tree/main"
hn_url: "https://news.ycombinator.com/item?id=48455012"
title: "How to have Claude Code run 28× longer before auto-compaction"
article_title: "GitHub - wiztek-llc/context-ledger: Commit-boundary context compaction for long-horizon coding agents — ties full-context retention at ~5x lower cost, proven Pareto-dominant · GitHub"
author: "newtechwiz"
captured_at: "2026-06-09T04:30:21Z"
capture_tool: "hn-digest"
hn_id: 48455012
score: 3
comments: 0
posted_at: "2026-06-09T01:30:34Z"
tags:
  - hacker-news
  - translated
---

# How to have Claude Code run 28× longer before auto-compaction

- HN: [48455012](https://news.ycombinator.com/item?id=48455012)
- Source: [github.com](https://github.com/wiztek-llc/context-ledger/tree/main)
- Score: 3
- Comments: 0
- Posted: 2026-06-09T01:30:34Z

## Translation

タイトル: 自動圧縮前にクロード コードを 28 倍長く実行する方法
記事のタイトル: GitHub - wiztek-llc/context-ledger: 長期的なコーディング エージェント向けのコミット境界コンテキストの圧縮 — フルコンテキストの保持を最大 5 倍の低コストで結び付ける、パレート支配であることが証明 · GitHub
説明: ロングホライズンコーディングエージェント向けのコミット境界コンテキスト圧縮 — フルコンテキストの保持を最大 5 倍の低コストで結び付け、パレート優位であることが証明されています - wiztek-llc/context-ledger

記事本文:
GitHub - wiztek-llc/context-ledger: 長期コーディング エージェント向けのコミット境界コンテキスト圧縮 — フルコンテキストの保持を最大 5 倍の低コストで結び付ける、パレート優位であることが証明 · GitHub
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
ウィズテックLLC
/
コンテキスト台帳
公共
通知
あなたはむ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット アーティファクト アーティファクト ベンチ ベンチ cl cl コーパス コーパス フック フック テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md RESEARCH.md RESEARCH.md RESULTS.md RESULTS.md ctxledger ctxledger install.sh install.sh run_all.sh run_all.sh uninstall.sh uninstall.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ロングホライズンコーディングエージェントのコミット境界コンテキスト圧縮。
製品を機能ごとに構築するエージェントは、ヒットするまでコンテキストを蓄積します。
モデルのウィンドウが表示され、タスクの途中で失敗するか、率直に要約されます (負けます)
詳細は回復できません）。 Context Ledger はこれを自然な継ぎ目で修正します。
git コミット 。機能がコミットされると、そのビルド コンテキストの最大 10,000 トークン
(ファイルの読み取り、推論、ツールの出力、差分) はコンパクトなものに置き換えられます。
決定事項、公開インターフェース、および
— 非常に重要なのは、正確なバイトへの git ポインターです。ポインタは、
エビクションセーフ : 後の作業でエビクションされた詳細が必要になった場合、エージェントは
オンデマンドで git からリハイドレートします。
コミットされたコード + git 履歴は外部メモリです。コミットはそうではありません
ただ「完了」しただけです — コンテキストを安全に削除できるロスレスストアです
それを生み出したのが。
これは、一時的に適用される「復元可能な圧縮」（コンテンツをドロップし、ハンドルを保持する）です。
意味上の境界。 RESEARCH.md との関連性については、RESEARCH.md を参照してください。
最先端 (人間のコンテキスト編集 + 記憶ツール、MemGPT/Letta、
マナス、クライン・メモリー・バンク、MEM1 /「アクションとしてのメモリー」）。
実際の 4 機能ビルド (このプロジェクトの兄弟リポジトリ) では、Context Ledger
を維持しながら、以前の機能の事実の完全なコンテキスト保持と一致します。
働いている

コンテキストは約 30 倍小さく、実行されるセッションの長さ全体に投影されます
素朴な場合、ロスレスで 200k ウィンドウに到達するまでに最大 28 倍の時間がかかります
切り捨ては忘れ、単純な要約は回復できない詳細を失います。
ピーク時の作業コンテキストとセッションの長さ (4 つの実際の機能を循環)。フル
(圧縮なし) N≈24 のフィーチャとダイで 200k ウィンドウを横切ります。台帳が残る
平ら（約 277 トーク/フィーチャ）で、N≈673 でのみ壁に到達します。
パレートフロンティア — リテンションとコンテキストコスト
保持率 = 以前の機能に関する 14 件の事実調査のうち、正しく回答した割合
保持されたコンテキストのみから (同じモデルがすべての戦略を判断します):
Ledger は Full の保持を 5 分の 1 のコストで結び付け、損失の多いものを打ち負かします
ベースラインは 29 ～ 36 ポイント、フロンティアの左上です。ベースライン
復元できないためアクセスできません: Truncate は最も古いものを忘れます
特徴; Rollingsummary の散文では詳細が省​​略されています。元帳は git を保管します
ポインタを作成してそれらを回復します。
機能 0 ビルド (≈13k tok) ── コミット ──▶ 元帳エントリ 0 (≈300 tok) + git SHA
機能 1 ビルド (≈ 7k tok) ── コミット ──▶ 台帳エントリ 1 (≈ 300 tok) + git SHA
機能 2 ビルド (≈ 9k tok) ── コミット ──▶ 台帳エントリ 2 (≈300 tok) + git SHA
...
作業コンテキスト = システム + Σ 元帳エントリ (コンパクト) + 現在の機能
プローブが削除する必要がある場合の詳細: 関連するエントリを取得 → `git show <sha>` → 回答
台帳エントリには以下が保持されます: 人間が作成したコミット メッセージ (意思決定と
Gotchas live)、公開インターフェースの署名 (後の機能)
に依存します)、変更されたファイルのマップ、およびリハイドレート ポインター。それ以外はすべてそうです
git から回復できるため、削除しても安全です。
git clone https://github.com/wiztek-llc/context-ledger
cd context-ledger && ./install.sh
ctxledger CLI をインストールし、ライブ

クロードコードのフック。元に戻す
./uninstall.sh を使用します。非対話型: ./install.sh --with-hooks /
--no-hooks / --with-bench 。
ライブ統合 — 実際のビルド上で実行します
フックがインストールされていると、Context Ledger は自動的に動作します。
git コミットのたびに → コンパクトな台帳エントリが記録されます
<project>/.claude/context-ledger.md がコンテキストに挿入されます (
PostToolUse フック on git commit )、エージェントはその機能のビルドを認識します
細部まで安全にコンパクトになりました。
セッション開始時 → 台帳が再挿入される (SessionStart フック)
圧縮されたフィーチャ メモリは、圧縮および新しいセッションを維持します。
ctxledger ドクター # ここに配線されていますか?何が記録されていますか？次に何をすべきか
ctxledger install-git-hook # Bulletproof: 「git commit」だけでなく、すべてのコミット (make/deploy/IDE) を記録します。
ctxledger Record # HEADコミットを抽出 → 台帳エントリを追加
ctxledger show # プロジェクト台帳を印刷します
ctxledger re水和物 < sha > [パス] # git から削除された正確な詳細を復元します
ctxledger stats 記録された機能数 · 保存されたトークン · 圧縮
「それは機能していますか？」任意のプロジェクトで ctxledger Doctor を実行します。
git リポジトリ、インストールされているフック、記録されている機能/トークンの数、
そして次のステップを教えてくれます。フックはセッション開始時にロードされるため、
セッションがすでに開いている間にそれらをインストールした場合は、一度セッションを再起動してください。
これはベンチマークが証明するメカニズムの展開可能な形式です: 台帳
ファイルは耐久性があり、復元可能なメモリです。 git はロスレス ストアです。
./run_all.sh # venv + テスト + スケーリング + リテンション ベンチマーク + チャート + スコアボード
ベンチマークは、バンドルされた実際のコーパス ( corpus/statusline.bundle 、
最初の実行時に .corpus/ に実体化されるため、公開された数値が再現されます。
新しいクローン作成後も同様です。個人用のパスやネットワークはありません。
決定的な部分。ポ

CL_REPO=/path/to/repo を持つ任意のリポジトリの int 。
決定的な部分（トークンカーブ、スケーリング、圧縮、リハイドレーション）は、
正確でネットワークは必要ありません。
保持ベンチマークは、ローカル クロード -p を 1 回につき 1 回呼び出します。
(戦略、プローブ)、ディスク上にキャッシュされるため、再実行は即座に行われ、LLM ごとに実行されます
出力はベンチ/キャッシュ/で検査可能です。すべての場合に同じモデルが答えます
戦略なので、比較は公平です。
.venv/bin/python testing/test_engine.py # 16 の決定論的エンジン チェック
.venv/bin/python bench/demo.py # live: Truncate は忘れたが、Ledger は回復するという事実
.venv/bin/python bench/summary.py # スコアボード
何が測定されるのか (そして何が測定されないのか)
保持率 = 以前の機能に関する事実調査のうち正しく回答した割合、
セッション終了時に戦略が保持したコンテキストのみを使用する
(最も困難なケース - 最も古い機能が最も圧縮されます)。事実は真実です
ビルドの詳細。ドキュメント ( .md ) は再構築されたビルドから除外されます
コンテキストを変更することで、後の機能の README から以前の機能の事実 (つまり、
不当に最新ベースのベースラインを支援することになります）。
コスト = 応答するために必要なコンテキストのトークン (Ledger の場合、これには
水分補給はオンデマンドで支払います）。
正直な制限は RESEARCH.md : retrieval にあります
見逃す可能性があります。リハイドレーションを呼び出すとトークンがかかります。コミットは良いですが、
不完全な境界線。
cl/エンジン: トークン · セッション (git から) · 台帳 · 戦略 · llm (キャッシュ)
ベンチ/プローブ · 実行(保持) · スケーリング · プロット · サマリー · デモ
テスト/決定論的エンジンテスト
アーティファクト/ results.json · scaling.json · *.png
について
ロングホライズンコーディングエージェント向けのコミット境界コンテキスト圧縮 — 最大 5 分の 1 の低コストでフルコンテキストの保持を実現し、パレート優位性が証明されています
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました

ディン。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Commit-boundary context compaction for long-horizon coding agents — ties full-context retention at ~5x lower cost, proven Pareto-dominant - wiztek-llc/context-ledger

GitHub - wiztek-llc/context-ledger: Commit-boundary context compaction for long-horizon coding agents — ties full-context retention at ~5x lower cost, proven Pareto-dominant · GitHub
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
wiztek-llc
/
context-ledger
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits artifacts artifacts bench bench cl cl corpus corpus hooks hooks tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md RESEARCH.md RESEARCH.md RESULTS.md RESULTS.md ctxledger ctxledger install.sh install.sh run_all.sh run_all.sh uninstall.sh uninstall.sh View all files Repository files navigation
Commit-boundary context compaction for long-horizon coding agents.
An agent building a product feature-by-feature accumulates context until it hits
the model's window and either fails mid-task or gets bluntly summarized (losing
detail it can't recover). Context Ledger fixes this at the natural seam — the
git commit . When a feature is committed, its ~10k tokens of build context
(file reads, reasoning, tool output, diffs) are replaced with a compact,
structured ledger entry that keeps the decisions, the public interfaces, and
— crucially — a git pointer to the exact bytes. The pointer is what makes
eviction safe : when later work needs an evicted detail, the agent
rehydrates it from git on demand.
The committed code + git history is the external memory. The commit isn't
just "done" — it's the lossless store that makes it safe to evict the context
that produced it.
This is "restorable compression" (drop the content, keep the handle) applied at a
semantic boundary. See RESEARCH.md for how it relates to the
state of the art (Anthropic context-editing + memory tool, MemGPT/Letta,
Manus, Cline Memory Bank, MEM1 / "Memory as Action").
On a real 4-feature build (this project's sibling repo), Context Ledger
matches full-context retention of earlier-feature facts while keeping the
working context ~30× smaller — and projected over session length it runs
~28× longer before hitting the 200k window , losslessly , where naive
truncation forgets and naive summarization loses detail it cannot recover.
Peak working-context vs session length (cycling the 4 real features). Full
(no compaction) crosses the 200k window at N≈24 features and dies; Ledger stays
flat (~277 tok/feature) and reaches the wall only at N≈673.
Pareto frontier — retention vs context cost
Retention = % of 14 factual probes about earlier features answered correctly
from only the retained context (same model judges every strategy):
Ledger ties Full's retention at 5× lower cost and beats the lossy
baselines by 29–36 points — the top-left of the frontier. The baselines
can't reach it because they aren't restorable: Truncate forgets the oldest
feature; RollingSummary's prose drops the specifics; Ledger keeps a git
pointer and recovers them.
feature 0 build (≈13k tok) ── commit ──▶ ledger entry 0 (≈300 tok) + git SHA
feature 1 build (≈ 7k tok) ── commit ──▶ ledger entry 1 (≈300 tok) + git SHA
feature 2 build (≈ 9k tok) ── commit ──▶ ledger entry 2 (≈300 tok) + git SHA
...
working context = system + Σ ledger entries (compact) + current feature
when a probe needs evicted detail: retrieve relevant entry → `git show <sha>` → answer
A ledger entry keeps: the human-authored commit message (where decisions and
gotchas live), the public interface signatures touched (what later features
depend on), the changed-file map, and the rehydrate pointer. Everything else is
recoverable from git and therefore safe to drop.
git clone https://github.com/wiztek-llc/context-ledger
cd context-ledger && ./install.sh
Installs the ctxledger CLI and (opt-in) the live Claude Code hooks . Revert
with ./uninstall.sh . Non-interactive: ./install.sh --with-hooks /
--no-hooks / --with-bench .
Live integration — runs on your real builds
With the hooks installed, Context Ledger works automatically:
after every git commit → a compact ledger entry is recorded to
<project>/.claude/context-ledger.md and injected into context (a
PostToolUse hook on git commit ), so the agent knows that feature's build
detail is now safe to compact.
at session start → the ledger is re-injected (a SessionStart hook), so
the compacted feature memory survives compaction and new sessions .
ctxledger doctor # is it wired up here? what's recorded? what to do next
ctxledger install-git-hook # bulletproof: record EVERY commit (make/deploy/IDE), not just `git commit`
ctxledger record # extract HEAD commit → append a ledger entry
ctxledger show # print the project ledger
ctxledger rehydrate < sha > [path] # recover exact evicted detail from git
ctxledger stats # features recorded · tokens saved · compression
"Is it even working?" Run ctxledger doctor in any project — it checks the
git repo, the installed hooks, and how many features/tokens have been recorded,
and tells you the next step. The hooks load at session start , so if you
installed them while a session was already open, restart it once.
This is the deployable form of the mechanism the benchmark proves: the ledger
file is the durable, restorable memory; git is the lossless store.
./run_all.sh # venv + tests + scaling + retention benchmark + charts + scoreboard
The benchmark runs on a bundled real corpus ( corpus/statusline.bundle ,
materialized into .corpus/ on first run) so the published numbers reproduce
identically after a fresh clone — no personal paths, no network for the
deterministic parts. Point at any repo with CL_REPO=/path/to/repo .
Deterministic parts (token curves, scaling, compression, rehydration) are
exact and need no network.
The retention benchmark calls the local claude -p once per
(strategy, probe), cached on disk so re-runs are instant and every LLM
output is inspectable in bench/cache/ . The same model answers for every
strategy, so the comparison is fair.
.venv/bin/python tests/test_engine.py # 16 deterministic engine checks
.venv/bin/python bench/demo.py # live: a fact Truncate forgets but Ledger recovers
.venv/bin/python bench/summary.py # the scoreboard
What's measured (and what isn't)
Retention = % of factual probes about earlier features answered correctly,
using only the context that strategy retained at the end of the session
(the hardest case — the oldest feature is the most compacted). Facts are real
details of the build; docs ( .md ) are excluded from the reconstructed build
context so a later feature's README can't leak earlier features' facts (that
would unfairly help recency-based baselines).
Cost = tokens of context needed to answer (for Ledger this includes the
rehydration it pays on demand).
Honest limitations are in RESEARCH.md : retrieval
can miss; rehydration costs tokens when invoked; the commit is a good but
imperfect boundary.
cl/ engine: tokens · session(from git) · ledger · strategies · llm(cached)
bench/ probes · run(retention) · scaling · plot · summary · demo
tests/ deterministic engine tests
artifacts/ results.json · scaling.json · *.png
About
Commit-boundary context compaction for long-horizon coding agents — ties full-context retention at ~5x lower cost, proven Pareto-dominant
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
