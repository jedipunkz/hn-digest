---
source: "https://github.com/glebmish/claude-code-replay"
hn_url: "https://news.ycombinator.com/item?id=48322770"
title: "Show HN: Claude-code-replay – get lost project code back from Claude Code logs"
article_title: "GitHub - glebmish/claude-code-replay: Replay Claude Code session logs to reconstruct lost project files, commit by commit. · GitHub"
author: "glebmish"
captured_at: "2026-05-30T11:48:12Z"
capture_tool: "hn-digest"
hn_id: 48322770
score: 2
comments: 0
posted_at: "2026-05-29T13:20:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude-code-replay – get lost project code back from Claude Code logs

- HN: [48322770](https://news.ycombinator.com/item?id=48322770)
- Source: [github.com](https://github.com/glebmish/claude-code-replay)
- Score: 2
- Comments: 0
- Posted: 2026-05-29T13:20:11Z

## Translation

タイトル: Show HN: Claude-code-replay – クロード コード ログから失われたプロジェクト コードを取得する
記事のタイトル: GitHub - glebmish/claude-code-replay: クロード コードのセッション ログを再生して、コミットごとに失われたプロジェクト ファイルを再構築します。 · GitHub
説明: クロード コードのセッション ログを再生して、コミットごとに失われたプロジェクト ファイルを再構築します。 - glebmish/クロードコードリプレイ

記事本文:
GitHub - glebmish/claude-code-replay: クロード コードのセッション ログを再生して、コミットごとに失われたプロジェクト ファイルを再構築します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
斑点
/
クロードコードリプレイ
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
72 コミット 72 コミット .github/ workflows .github/ workflows docs docs src src testing testing .gitignore .gitignore ライセンス ライセンス README.md README.md esbuild.config.mjs esbuild.config.mjs package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードのセッション ログ ( *.jsonl ) を再生して、失われたプロジェクトを再構築します。
状態 — ファイルごと、コミットごとに、イベントが発生した順序で。
破壊的なコマンドによってツリーが消去された場合の最後の手段のツール。
2 つの再生レイヤーがあり、2 番目はオプトインです。
確定的リプレイ — --logs-dir の下にあるすべての *.jsonl を調べます
( <session>/subagents/ の下にあるサブエージェント JSONL を含む)
時系列順にファイルの書き込みを適用します。
クロード分類子 (オプトイン、 --enable-llm-classifier ) — すべての Bash
それ以外の場合、イベントはスキップされます。分類子をオンにすると、各 Bash
イベントはクロード (Sonnet 4.6) に送信され、クロードが決定します。
イベントごとに理由を付けて実行またはスキップします。参照
分類子の仕組み。
Claude Code CLI がインストールされ、認証されている (claude login) — のみ
分類子に必要です ( --enable-llm-classifier )。分類子
Claude Agent SDK を介してその認証を再利用するため、別個の Anthropic API は必要ありません
キーが必要です。
npx claude-code-replay --target … --source-root … [フラグ]
または、グローバルにインストールします。
npm install -g クロードコードリプレイ
クロードコードリプレイ --target … --source-root … [フラグ]
または、ソースからビルド — クローン、npm install のいずれかを実行します。
npm run restart -- <flags> を直接、または npm リンクして次のように公開します
PATH 上の claude-code-replay 。
クロードコードリプレイ \
--target /tmp/myrepo-re

覆われた／＼
--source-root /Users/you/projects/myrepo \
--enable-llm-classifier
標準出力に表示される内容 (実際の 304 イベントのリプレイから):
/Users/you/.claude/projects/-Users-you-projects-myrepo からイベントを収集する INFO
INFO は 304 件のイベントを収集しました
INFO ファイル履歴スナップショット エントリからスナップショット インデックスを構築
INFO スナップショット インデックスは 43 のパスをカバーします
INFO 分類子: 208 個のペイロード イベント (117 個の Bash、91 個のコンテキスト) にわたる 4 つのバッチ。サイズ=[71,64,59,14]
INFO 分類子モデル = クロード-ソネット-4-6、モード = ベース、ソースルート = 1
INFO 分類子バッチ 1/4 キャッシュ ヒット
INFO 分類子バッチ 2/4 キャッシュ ヒット
INFO 分類子バッチ 3/4 キャッシュ ヒット
INFO 分類子バッチ 4/4 キャッシュ ヒット
INFO 分類子は 208 件の決定を返しました
=== クロードコードリプレイの概要 ===
イベント合計: 304
再生回数: 64 (64 件中)
スキップ: 240
bash の実行: 25/25
分類子バッチ: 4 (4 キャッシュ、0 ライブ)
停止: いいえ
経過: 3.70秒
対象ファイル: 732 (合計 8519381 バイト)
この要約では、通常の実行ではゼロになる行が省略されます (オーバーライド、
cwd でフィルターされた Bash、スナップショットの修復、寛大な読み取りスキップ)。イベントごと
CLASSIFY / APPLY / CHECK トレースと詳細な分類子
診断は --debug の背後でゲートされます。実際のエラー (argv 解析
失敗、分類子 API エラー) は stderr に送られます。この実行ログは次の場所に保存されます
stdout を使用して、診断を失わずにパイプできるようにします。
終了コード: 0 成功、2 argv エラー、10 はコマンド失敗時に停止。
--target <パス> — リプレイが書き込まれるディレクトリ。区別する必要があります
すべてのログ ディレクトリから (どちらの方向でも) rm -rf を再生しました。できるかもしれない
それ以外の場合は、実行中にログが破棄されます。
--source-root <path> — セッションからの元の絶対 CWD。
ログ内のevent.cwdと逐語的に比較されるため、一致する必要があります
文字ごとに (相対パスやシンボリックリンク解決されたパスはありません)。
ルートを越えて移動したセッションでも繰り返し可能。
--logs-dir <パス>

— セッション *.jsonl ファイルを含むディレクトリ
(および任意の <session>/subagents/JSONL)。オプション、繰り返し可能。によって
デフォルトでは、各 --source-root から 1 つのログ ディレクトリが次のように推測されます。
~/.claude/projects/<encoded-source-root> (絶対パス内のすべての /
source-root は - ) に置き換えられます。存在しない推測されたディレクトリ
ディスクはサイレントにスキップされます。明示的な --logs-dir 値が追加されます
推論されたセットの先頭にあり、存在する必要があります。
--cutoff <iso-ts> — この ISO 8601 タイムスタンプ以降のイベントを削除します
解析時。セッションの後のイベントに以下が含まれる場合に使用します。
破壊的な手術から回復中です。
--start <iso-ts> — タイムスタンプが一致する最初のイベントで再生を開始します。
これ以降です。 --cutoff で構成してウィンドウを定義します。の
ターゲット ディレクトリは、以前の状態イベントをすでに反映していると信頼されています。
--start が生成されます。
--from-index <N> — イベント インデックス N で再生を開始します (イベント 0..N-1
分類または適用されていません）。 --start で作成します。どちらでも
後で土地が勝ちます。停止と再開のプリミティブ: K で停止したら、修正します
原因を調べて、 --from-index K で再開します。
--strict — 両方の修復層 (スナップショット修復と読み取りの適用) を無効にします。
治ります）。読み取りの不一致またはターゲットの欠落がある場合は、直ちに停止します。役に立つ
リプレイのどの程度の修復が必要かを測定するとき (例:
分類子の評価 — デフォルト モードのヒール カウントは、
分類子はテーブル上に残されました）。
--strict-reads — 最初に失敗した読み取りチェックポイントで停止するのではなく、
デフォルト (ログ + 続行)。どのイベントをデバッグするのに役立ちます
ファイル欠落シナリオをトリガーしました。デフォルトでオンになっている寛大な動作は、
長いリプレイが分類器のたびに停止しないのはなぜですか
生成する Bash チェーンを正しく省略します (次のカスケード ルールを参照)
docs/classifier-prompt.md )。
--enable-llm-classifier — LLM 呼び出しをオプトインします。使用するために必要な

の
まったく分類子。基本プロンプトには常に git に焦点を当てたプロンプトが含まれます。
git add / git commit / git Branch / を呼び出す補足
git checkout / git merge / git rebase / git revert /
git replace / git tag / git filter-repo (すべてを網羅しているわけではありません。同じです)
ロジックは、同等の状態変化コマンドに拡張され、
heredoc/sed はその内容を書き込み、後のコミットでキャプチャします)。復元中
元の git 履歴は、現実世界の主要なユースケースです。
再生するため、オプトイン フラグではなくデフォルトとして出荷されます。
--custom-intent "<intent>" — 自然言語インテントを追加します
リプレイで達成すべきことを説明します。行動に使用する
組み込みの git フォーカスを超えたもの、例:
「すべての依存関係のインストール (npm/pip) を保持し、node_modules にデータが取り込まれるようにします。」
または、「docker/podman コマンドをスキップします。このリプレイはデーモンなしで実行されます」。
再現可能。各値は改行で結合されます。
--override-classifier-cache — 分類子キャッシュからの読み取りをスキップします
新しい LLM 呼び出しを強制しますが、それでも新しい応答を書き戻します。
キャッシュ (既存のエントリを上書き)。
--skip-uncached-tail — キャッシュされた実行の last_event_ts が低下した場合
現在のログ内で、それ以降のタイムスタンプを持つすべてのイベントを削除します
分類子がそれらを認識する前に。その後、分類子がフルヒットします。
キャッシュすると、ランタイムはすでにキャッシュされているものだけを再生します。
「今日のリプレイに対して昨日のリプレイを少しだけ再実行する」ことを目的としています。
新しい尾の代金を支払わずに丸太を育てました。」キャッシュが存在しない場合、
フラグは警告し、切り捨てられずに続行します。注意: イベント
キャップを超えると未分類になります - 追加された末尾に
破壊的なコマンドは表示されません。
--override-skip <INDEX> — 繰り返し可能。イベント INDEX を強制的にスキップします。
ルールベースまたは LLM 分類に関係なく。あらゆるイベントに対応
タイプ ( Bash 、 Read 、 Edit 、 Write 、チェックポイント)。
--ov

erride-execute <INDEX>[=CMD] — 反復可能。強制イベント INDEX
(Bash のみ) を実行します。裸のフォームはイベントの元のコマンドを実行します。
=CMD は代わりに部分文字列 CMD を実行します (リテラル部分文字列である必要があります)
イベントの元のコマンドの - LLM と同じ制約
分類子の決定.コマンド)。同様の対象となります
cwd -inside-source-roots は、classifier-approved の実行時にチェックします。
--dry-run — 分類のみで、実行はありません。イベントストリームを歩き、
概要を出力しますが、書き込み/編集は適用されません。読み取りを確認します。
チェックポイントを削除するか、承認された Bash を実行します。 --debug と組み合わせて確認してください
すべてのイベントのイベントごとの CLASSIFY 行。
--debug — イベントごとの CLASSIFY / APPLY / CHECK をオンにします
トレース (イベントごとに 1 行) と詳細な分類器の計測。
デフォルトの実行では少数の情報が保持されるため、デフォルトではオフになります。
設定行と最終的な要約。
すべてのリクエストは Claude Agent SDK ( claude-sonnet-4-6 、
マルチターン ストリーミング、ツールは使用しません)。クロードコードの既存のものを再利用します。
auth — 個別の API キーは必要ありません。デフォルトのモデル ID は、
200k コンテキストのバリアント。 1M コンテキスト バリアント ( [1m] への切り替え)
接尾辞) には、Anthropic の「使用クレジット」オプトインが必要であり、現在は
src/llm-classifier/sdk.ts のソースレベルの切り替え。
Bash ペイロードは 50 ～ 100 個のイベントのバッチに分割され、最初のイベントで切り取られます。
git commit がしきい値を超えました。各バッチは 1 回のユーザー ターンになります。
単一の会話なので、システム プロンプトと以前のバッチは
後続のターンでキャッシュが提供されます。
バッチごとの応答は次の場所にキャッシュされます。
$XDG_CACHE_HOME/claude-code-replay/<エンコードされたターゲット>/batch-NNNN.json
加えて 1 つの meta.json 。エンコーディングはクロード コードのエンコーディングを反映しています。
project-dir スキーム: 絶対 --target パス内のすべての / は、
- なので、/tmp/myrepo-recovered は次の場所にあります。
$XDG_CACHE_HOME/claude-code-replay/-tmp-myr

epo-recovered/ 。キャッシュ
ディレクトリは意図的に --target の外側にあるため、再生されます
git add 。掃き込むことはできません。
キャッシュの無効化。全か無か: 1 つの共有キーですべてをカバー
バッチなので、入力を変更すると実行全体が一度に無効になります。で
不一致の場合、古いエントリは消去され、分類器は次から再計算します。
スクラッチ、および実行ログにはどの入力が変更されたかが記録されます (例:
INFO 分類子キャッシュ ミス: 入力が変更されました (セッション ログ)。古いエントリを消去し、4 つのバッチを再計算します)。以下の変更点
無効にする:
システム プロンプト ( src/llm-classifier/prompts.ts ) を編集します。
--custom-intent の追加、削除、または変更。
--source-root セットを変更します。
クロード コードはセットの変更をログに記録します — 新しいセッション JSONL が表示され、
既存のものは拡大するか、--cutoff (解析時に適用) によって切り取られます。
セット。異なる --from-index / を使用して同じログを再開する
--start は無効にしません。
--override-classifier-cache — なしで新しい呼び出しを強制します
キャッシュを参照する。結果はまだ書き戻されます。
別の --target を指している — キャッシュのサブディレクトリは
エンコードされたターゲット パスなので、別のターゲットは別のキャッシュになります
名前空間 (古いものは削除されずに孤立したままになります)。
リテラルのシステム プロンプト — ファイル依存関係カスケードを含む
ルール — に住んでいます
src/llm-classifier/pro

[切り捨てられた]

## Original Extract

Replay Claude Code session logs to reconstruct lost project files, commit by commit. - glebmish/claude-code-replay

GitHub - glebmish/claude-code-replay: Replay Claude Code session logs to reconstruct lost project files, commit by commit. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
glebmish
/
claude-code-replay
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
72 Commits 72 Commits .github/ workflows .github/ workflows docs docs src src tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md esbuild.config.mjs esbuild.config.mjs package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Replay Claude Code session logs ( *.jsonl ) to reconstruct the lost project
state — file by file, commit by commit, in the order events happened.
The tool of last resort when a destructive command wiped the tree.
There are two replay layers, and the second is opt-in:
Deterministic replay — walks every *.jsonl under --logs-dir
(including subagent JSONLs under <session>/subagents/ ) in strict
chronological order and applies file writes.
Claude classifier (opt-in, --enable-llm-classifier ) — every Bash
event would otherwise be skipped. With the classifier on, each Bash
event is sent to Claude (Sonnet 4.6) which decides
execute or skip per event, with reasons. See
How the classifier works .
Claude Code CLI installed and authenticated ( claude login ) — only
needed for the classifier ( --enable-llm-classifier ). The classifier
reuses that auth via the Claude Agent SDK, so no separate Anthropic API
key is needed.
npx claude-code-replay --target … --source-root … [flags]
Or install globally:
npm install -g claude-code-replay
claude-code-replay --target … --source-root … [flags]
Or, build from source — clone, npm install , then either
npm run replay -- <flags> directly or npm link to expose it as
claude-code-replay on your PATH .
claude-code-replay \
--target /tmp/myrepo-recovered/ \
--source-root /Users/you/projects/myrepo \
--enable-llm-classifier
What you'll see on stdout (from a real 304-event replay):
INFO collecting events from /Users/you/.claude/projects/-Users-you-projects-myrepo
INFO collected 304 events
INFO building snapshot index from file-history-snapshot entries
INFO snapshot index covers 43 paths
INFO classifier: 4 batch(es) over 208 payload events (117 Bash, 91 context); sizes=[71,64,59,14]
INFO classifier model=claude-sonnet-4-6, mode=base, source-roots=1
INFO classifier batch 1/4 cache hit
INFO classifier batch 2/4 cache hit
INFO classifier batch 3/4 cache hit
INFO classifier batch 4/4 cache hit
INFO classifier returned 208 decisions
=== claude-code-replay summary ===
events total: 304
replayed: 64 (of 64)
skipped: 240
bash executed: 25 of 25
classifier batches: 4 (4 cached, 0 live)
halted: no
elapsed: 3.70s
target files: 732 (8519381 bytes total)
The summary omits rows that would be zero on a typical run (overrides,
cwd-filtered Bash, snapshot heals, lenient-read skips). Per-event
CLASSIFY / APPLY / CHECK traces and detailed classifier
diagnostics are gated behind --debug . Real errors (argv parse
failures, classifier API errors) go to stderr; this run log goes to
stdout so you can pipe it without losing diagnostics.
Exit codes: 0 success, 2 argv error, 10 halted on command failure.
--target <path> — directory the replay writes into. Must be distinct
from every logs dir (in either direction); replayed rm -rf . could
otherwise destroy the logs mid-run.
--source-root <path> — original absolute cwd from the session.
Compared verbatim against event.cwd in the logs, so it must match
character-for-character (no relative paths, no symlink-resolved paths).
Repeatable for sessions that moved across roots.
--logs-dir <path> — directory containing the session *.jsonl files
(and any <session>/subagents/ JSONLs). Optional, repeatable. By
default, one logs dir is inferred from each --source-root as
~/.claude/projects/<encoded-source-root> (every / in the absolute
source-root is replaced with - ). Inferred dirs that don't exist on
disk are silently skipped; explicit --logs-dir values are added on
top of the inferred set and must exist.
--cutoff <iso-ts> — drop events at or after this ISO 8601 timestamp
at parse time. Use when the session's later events include the
destructive operation you're recovering from.
--start <iso-ts> — start replay at the first event whose timestamp
is at or after this. Composes with --cutoff to define a window. The
target dir is trusted to already reflect the state events before
--start would have produced.
--from-index <N> — start replay at event index N (events 0..N-1
are not classified or applied). Composes with --start ; whichever
lands later wins. The halt-and-resume primitive: on a halt at K , fix
the cause and resume with --from-index K .
--strict — disable both heal layers (snapshot heal and apply-reads
heal). Any Read mismatch or missing target halts immediately. Useful
when measuring how much of a replay needs healing (e.g. when
evaluating a classifier — heal counts in default mode signal what the
classifier left on the table).
--strict-reads — halt on the first failed Read checkpoint instead of
the default (log + continue). Useful for debugging which event
triggered a missing-file scenario; the default-on lenient behaviour is
what keeps long replays from stopping every time the classifier
correctly omits a producing Bash chain (see the cascade rule in
docs/classifier-prompt.md ).
--enable-llm-classifier — opt in to LLM calls. Required to use the
classifier at all. The base prompt always includes a git-focused
supplement that calls out git add / git commit / git branch /
git checkout / git merge / git rebase / git revert /
git reset / git tag / git filter-repo (non-exhaustive; the same
logic extends to any equivalent state-mutating command, and to
heredoc/sed writes whose content a later commit captures). Restoring
the original git history is the dominant real-world use case for
replay, so it ships as the default rather than an opt-in flag.
--custom-intent "<intent>" — append a natural-language intent
describing what the replay should accomplish. Use for behaviour
beyond the built-in git focus, e.g.
"keep all dependency installs (npm/pip) so node_modules ends up populated"
or "skip any docker/podman commands; this replay runs without a daemon" .
Repeatable; each value is joined with a newline.
--override-classifier-cache — skip reading from the classifier cache
and force a fresh LLM call, but still write the new response back to
the cache (overwriting any existing entry).
--skip-uncached-tail — if the cached run's last_event_ts falls
inside the current logs, drop every event with a later timestamp
before the classifier sees them. The classifier then full-hits the
cache and the runtime replays only what was already cached.
Intended for "re-run yesterday's replay against today's slightly
grown logs without paying for the new tail." If no cache exists, the
flag warns and proceeds without truncation. Caveat: events
past the cap go unclassified — if the appended tail contains a
destructive command, you won't see it.
--override-skip <INDEX> — repeatable. Force event INDEX to skip,
regardless of any rule-based or LLM classification. Works on any event
type ( Bash , Read , Edit , Write , checkpoint).
--override-execute <INDEX>[=CMD] — repeatable. Force event INDEX
( Bash only) to execute. Bare form runs the event's original command;
=CMD runs the substring CMD instead (must be a literal substring
of the event's original command — same constraint as the LLM
classifier's decision.command ). Subject to the same
cwd -inside-source-roots check as classifier-approved executes.
--dry-run — classify only, no execution. Walks the event stream,
prints the summary, but does not apply Writes/Edits, verify Read
checkpoints, or execute approved Bash. Combine with --debug to see
the per-event CLASSIFY line for every event.
--debug — turn on the per-event CLASSIFY / APPLY / CHECK
trace (one line per event) plus verbose classifier instrumentation.
Off by default because the default run keeps to a handful of INFO
setup lines and the final summary.
All requests go through the Claude Agent SDK ( claude-sonnet-4-6 ,
multi-turn streaming, no tool use). It reuses Claude Code's existing
auth — no separate API key needed. The default model id targets the
200k-context variant; switching to the 1M-context variant ( [1m]
suffix) requires Anthropic "Usage credits" opt-in and is currently a
source-level toggle in src/llm-classifier/sdk.ts .
The Bash payload is split into batches of 50–100 events, cut at the first
git commit past the threshold. Each batch becomes one user turn in a
single conversation, so the system prompt and earlier batches are
cache-served on subsequent turns.
Per-batch responses are cached at
$XDG_CACHE_HOME/claude-code-replay/<encoded-target>/batch-NNNN.json
plus a single meta.json . The encoding mirrors Claude Code's own
project-dir scheme: every / in the absolute --target path becomes
- , so /tmp/myrepo-recovered lives at
$XDG_CACHE_HOME/claude-code-replay/-tmp-myrepo-recovered/ . The cache
directory is intentionally outside --target so a replayed
git add . can't sweep it in.
Cache invalidation. All-or-nothing: one shared key covers every
batch, so any input change invalidates the entire run at once. On a
mismatch the stale entries are wiped, the classifier recomputes from
scratch, and the run log states which input changed (e.g.
INFO classifier cache miss: inputs changed (session logs); wiping stale entries and recomputing 4 batch(es) ). The following changes
invalidate:
Editing the system prompt ( src/llm-classifier/prompts.ts ).
Adding, removing, or changing any --custom-intent .
Changing the --source-root set .
Claude Code logs set changes — a new session JSONL appears, an
existing one grows, or --cutoff (applied at parse time) crops the
set. Resuming the same logs with a different --from-index /
--start does NOT invalidate.
--override-classifier-cache — forces a fresh call without
consulting the cache; results are still written back.
Pointing at a different --target — the cache subdir is the
encoded target path, so a different target is a different cache
namespace (the old one is left orphaned, not deleted).
The literal system prompt — including the file-dependency cascade
rule — lives in
src/llm-classifier/pro

[truncated]
