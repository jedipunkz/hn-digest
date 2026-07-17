---
source: "https://github.com/anlor1002-alt/regressionledger"
hn_url: "https://news.ycombinator.com/item?id=48946385"
title: "Show HN: A hook that stops Claude Code from re-trying fixes that failed"
article_title: "GitHub - anlor1002-alt/regressionledger: Claude Code hook that blocks fixes that already failed - fingerprints edits, links them to test outcomes, survives context compaction. 100% of disguised repeat fixes caught, zero false blocks (npm run bench). · GitHub"
author: "anlor1002-alt"
captured_at: "2026-07-17T13:05:55Z"
capture_tool: "hn-digest"
hn_id: 48946385
score: 2
comments: 0
posted_at: "2026-07-17T12:09:58Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A hook that stops Claude Code from re-trying fixes that failed

- HN: [48946385](https://news.ycombinator.com/item?id=48946385)
- Source: [github.com](https://github.com/anlor1002-alt/regressionledger)
- Score: 2
- Comments: 0
- Posted: 2026-07-17T12:09:58Z

## Translation

タイトル: HN を表示: クロード コードが失敗した修正を再試行するのを停止するフック
記事のタイトル: GitHub - anlor1002-alt/regressionledger: すでに失敗した修正をブロックするクロード コード フック - フィンガープリントの編集、テスト結果へのリンク、コンテキストの圧縮の存続。偽装された反復修正が 100% 検出され、誤ったブロックはゼロ (npm 実行ベンチ)。 · GitHub
説明: Claude コード フックは、すでに失敗した修正をブロックします。フィンガープリントの編集、テスト結果へのリンク、コンテキスト圧縮の存続を実現します。偽装された反復修正が 100% 検出され、誤ったブロックはゼロ (npm 実行ベンチ)。 - anlor1002-alt/回帰台帳

記事本文:
GitHub - anlor1002-alt/regressionledger: すでに失敗した修正をブロックするクロード コード フック - フィンガープリントの編集、テスト結果へのリンク、コンテキストの圧縮の存続。偽装された反復修正が 100% 検出され、誤ったブロックはゼロ (npm 実行ベンチ)。 · GitHub
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
を却下する

警告します
{{ メッセージ }}
anlor1002-alt
/
回帰台帳
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
anlor1002-alt/回帰台帳
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
30 コミット 30 コミット .claude-plugin .claude-plugin .github .github ベンチ ベンチ bin bin デモ デモ ドキュメント ドキュメントの例 例 フック フック スクリプト スクリプト src src テスト test .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェントがすでに失敗した修正を復活させないようにします。
コーディング エージェント (クロード コード、カーソルなど) は、コンテキスト ウィンドウが表示されるとメモリを失います。
いっぱいになったら、すでに試したパッチを自信を持って再適用します。2 つのプロンプトが表示されます。
または 2 セッション前。同じ間違った修正が入り、同じテストが赤くなるのを見ます。
そして、何かを直し、また壊すという破滅のループでクレジットが蒸発してしまいます。
RegressionLedger は、依存性のない小さなクロード コード フック + CLI であり、次のことを行います。
エージェントが行うすべての編集のフィンガープリント (表面的に正規化されている)
違いは重要ではありませんが、true と false は依然として重要です)、
各編集を次のテスト/ビルド実行の結果 (成功または失敗) にリンクします。
それをセッションの再起動後も存続するローカル台帳に保存し、
コンテキストの圧縮、および
エージェントがアクセスしようとした瞬間に、PreToolUse 拒否を介してエージェントをハードブロックします。
以前に失敗したパッチを再適用し、その理由と目的を正確に伝える
戦略を変える。
⛔ 別のテストサイクルを無駄にする前にブロックされました:
RegressionLedger: 2 時間前に src/auth.js に対して同じ修正をすでに試しています。
次のエラーで失敗しました: AssertionError: Expected 200, got 401.
再適用すると再発します

同じ失敗を引き起こします。代わりに戦略を変更してください。
完全な試行履歴を表示するには、「rl show src/auth.js」を実行します。
今すぐ試してみてください。コマンド 1 つで、インストールやエージェントは不要で、何も触れません (
シミュレーション全体は、使い捨ての一時ディレクトリ内の実際のフック コードを通じて実行されます。
npx回帰台帳のデモ
なぜこれが存在するのか
「AI が 1 つのバグを修正し、2 つのバグを作成する」ループは、最も多く報告されている単一のループの 1 つです。
コーディングエージェントとの作業の苦痛。失敗には特定の形があります。
すでに試行したことに関するエージェントの記憶が、長いビルド後には残らない、または
コンテキスト ウィンドウのリセット ( "これにより、新しい Composer を開始する必要が生じ、すべてが失われます)
以前のコンテキスト。" );
既存のループ検出器は、1 つのループ内の同一のツール呼び出しのみを認識します。
セッション - セッション間で、どの修正が試行されたか覚えていない
壊れたもの。
一般的なエージェント記憶層はアドバイスのみを行います (「これが教訓です」)。
モデルは、無視しても構わないアドバイスを日常的に無視します。
RegressionLedger はまさにそのギャップ、つまりクロスセッション、セマンティック、
修正試行の結果にリンクされた台帳。ハードブロックとしてではなく表面化しました。
提案です。
PreToolUse (編集/書き込み/マルチ編集)
│
提案された変更の指紋を採取する
│
以前に失敗した試行がしきい値を超えていますか?
│ │
はい、いいえ
│ │
┌─────┴─────┐ 許可する
│ ブロック: 拒否 + 理由 │
━━━━━━━━━━━┘
PostToolUse (編集/書き込み/マルチ編集) ── 試行を「保留中」として記録します
PostToolUse (Bash: npm test / pytest / …) ── 保留中の解決 → パス |失敗する
フィンガープリント (2 チャネル) — 変更されたコードは正規化されたものに字句化されます。
トークンストリーム。生のチャネルはリテラルをそのまま保持します (空白と空白のみが保持されます)。
コメントは削除されます): ここでの一致は、再試行が同じコードであることを証明します。
定数が含まれており、それは t

彼のみがハードブロックを許可されているチャネルです。
折りたたまれたチャネルは文字列/数値リテラルを STR / NUM に抽象化します。
( return true ≠ return false 、意図的に): 折りたたまれたのみの一致は、
「同じ形状、異なる定数」 - 多くの場合、正当な次の実験
(タイムアウト 5000 → 30000 ) したがって、勧告ノートが生成されますが、否定は生成されません。
結果のリンケージ — エージェントが npm test / pytest / tsc / を実行した後
カーゴテスト / …、フックは出力を解析し、前回以降の編集をマークします。
pass または failed として実行し、最初のエラー行を署名としてキャプチャします。とき
修正が最終的に成功し、古い一致失敗レコードはすべてリタイアされるため、
二度とブロックしないでください。
元帳 — .regressionledger/ledger.json にあるプレーンな JSON ファイル。いいえ
データベース、ネイティブ モジュール、API キー、ネットワークはありません。それは、
圧縮に耐えます。
実際にマッチングはうまくいくのでしょうか？
数秒で自分でチェックできます。ベンチマークは決定的です。
npm ランベンチ
# 1. 化粧品の再適用 HARD-matched (生) : 60/60 (100.0%)
# 2. ノートではないブロックにルーティングされたリテラル バリアント: 40/40 (100.0%)
# ...誤ってハードブロックされた (誤検知) : 0/40
# 3. いずれかのチャンネルに一致する個別の修正: 0/190 (0.0%)
3 つの正直なカテゴリ: 外観上の再適用 (空白/コメント - 同じ)
fix) は厳密に一致します。変更された定数 (タイムアウト 5000 → 30000) は決して変更されません
ブロック — 彼らは勧告ノートにルーティングされ、独自の評決を取得します。そして
まったく異なる修正は何も一致しません。このベンチマークの以前のバージョン
文字通りの変更を「捕まえるための変装」としてカウントし、ゼロ偽ブロックを作成しました
循環番号 — コミュニティのストレステストでそれが検出され、それに一致するデザインが見つかった
そのせいで変わった。
クロードコードプラグインとして（推奨）
/プラグインマーケットプレイス追加 anlor1002-alt/regressionledger
/plugin install regressionledger@anlor1002-plugin

s
フックは自動的に有効になります。settings.json の編集や再起動は必要ありません。
# プロジェクト内で
npx regressionledger init # または: npm i -g regressionledger && rl init
init は以下を行います:
PreToolUse / PostToolUse フックを ./.claude/settings.json に書き込みます。
(既存のフックを上書きするのではなくマージします)、
デフォルトの config.json を使用して .regressionledger/ を作成します。
.regressionledger/ledger.json を .gitignore に追加します。
次に、Claude Code を再起動 (または /hooks を実行) して、新しいフックを選択します。それは
それ — RegressionLedger は、繰り返し失敗した修正が表示されるまで、静かに動作するようになりました。
脱出ハッチ: rl init を開始し、rl uninstall して痕跡を残さないようにします。それ
init で追加されたフック エントリ (他のフック、権限、
そして、設定は決して変更されません。また、rl uninstall --purge は、
ローカルの .regressionledger/ データ。
手動で配線したほうがいいですか? rl init --print は、正確なフック ブロックを出力します。
これを読んでいるコーディング エージェントの場合は、RegressionLedger をインストールできます。
あなたの人間のためにワンステップで。クロード コードでスラッシュ コマンドを実行します。
/plugin Marketplace に anlor1002-alt/regressionledger を追加してから
/plugin install regressionledger@anlor1002-plugins — PreToolUse/PostToolUse
フックは自動的に有効になります。あるいは、npx regressionledger init を実行します。
プロジェクトのルートを確認し、npx regressionledger Doctor で検証します (すべてのチェックを行います)
合格する必要があります。ライブ往復はガードレールが機能していることを証明します)。ツールは
ローカルのみ: ネットワーク呼び出し、API キーなし、フェールオープン。それはあなたをブロックします
以前に失敗した修正を再適用する - これは意図された動作であり、そうではありません。
エラーです。ブロックされた場合は、同じパッチを書き直すのではなく、戦略を変更します。
コンテキストが本当に変更された場合、人間は rl unblock <file> を実行できます。
それ自体が実行されます。実際に使用するコマンドは次のとおりです。
.regressionledger/config.jso

n (コミットしても安全 — チームとしきい値を共有):
初日は誤検知について不安ですか? rl init --warn から始めます。毎
would-have-blocked イベントがログに記録され、rl stats でカウントが表示され、rl show
独自のコードに対してそれぞれを監査できます。次に、次のコマンドを使用してブロックに切り替えます。
証拠 (rl config モード ブロック)。追加の安全マージンを得るには、次の修正が必要です
ブロックされる前に 2 回失敗します: rl config minFailures 2 。
ツール
クロスセッション?
成果連動？
ブロックしますか（アドバイスするのではなく）？
セッション内ループ検出器 (同一のツール呼び出しハッシュ)
✗
✗
警告する
一般的なエージェントの記憶 (「学んだ教訓」)
✓
✗
アドバイスする
自己介入研究 (ウィンク、…)
✗
部分的な
ガイド
回帰台帳
✓
✓
ブロック
差別化要因: セマンティック修正フィンガープリント (生のテキストやツール + 引数ではない)
ハッシュ）、結果リンク（失敗した修正とその理由）、クロスセッション +
圧縮後の永続性、およびモデルが無視できないハード ブロック。
ロングバージョン - 記憶ツールの何が優れているか、そしてどこでこれが優れているかを含む
正直なところ、このツールは不十分です: docs/COMPARISON.md 。
他のツールではできないこと (私たちが知る限り)
🧠セッション説明会。 SessionStart フックは、コンパクトな「すでにあるもの」を挿入します。
「ここで失敗しました」というメッセージがセッション開始時に毎回表示されます (直後も含む)
コンテキストの圧縮により、エージェントのメモリが消去されます。行き止まりは前にブロックされます
それらは、再適用される直前ではなく、再考されます。
🌀 スラッシュエスカレーション。同一の修正をブロックすると、破滅のループが発生します。の
他はすべて同じエラーで停止する異なる修正です。 3+ 異なる場合
アプローチが壁にぶつかると、フックはエスカレートします：「診断は間違っています、そうではありません」
パッチ — 停止し、根本原因の仮説を述べ、検証してから編集します。」
🧬 デュアルチャネルフィンガープリント。回避に使用されるすべての変数の名前を変更する
マッチャー。構造のみの 2 番目のフィンガープリントは、「この m

君も同じだろう
修正、名前変更」 - 弱い信号をブロックすることはありません。
🤝集団免疫。 rl エクスポート / rl インポートの間で決着した評決を共有します。
マシンとチームメイト: あなたのエージェントはすでに私のエージェントの行き止まりを引き継いでいます
支払われました。失敗は引き続き原因 ( importFrom ) として記録され、監査可能になります。輸入品
テキストは構造的に中和され、どこにでも [インポートされた評決] というラベルが付けられます。
エージェントのコンテキストに到達しますが、共有台帳を依存関係のように扱います: import
あなたが信頼する人々から。
これは単なる記憶MCP/教訓層ではないでしょうか？
いいえ、モデルが無視する可能性のあるストアのアドバイスです。これにより PreToolUse 拒否が発行されます
モデルは無視できず、実際のテスト実行でその正確な修正が証明された場合にのみ無視されます。
すでに失敗しました。記憶が蘇ります。 RegressionLedger は拒否します。完全な比較
(メモリツールがより優れている点を含む): docs/COMPARISON.md 。
同じ関数に対する正当な別の編集はブロックされますか?
いいえ — マッチングは関数ではなく、変更されたコード自体に基づいて行われます。あ
同じバグに対する異なるアプローチには、異なるフィンガープリントがあり、通過します。
正しい修正が単なる別の定数 (タイムアウト 5000 → 30000) である場合はどうなるでしょうか?
ブロックされたことはありません。ハード ブロックには生のチャネルの一致が必要です (同じコード、
定数も含まれます)。変更された定数は折りたたまれたチャネルにのみ一致するため、
注意事項 (「同じ形状、異なる値 - 値を変更する場合)」が表示されます。
はあなたの仮説です。続行してください」) とそれ自体の評決です。失敗したもののみを再適用します。
異形の動詞

[切り捨てられた]

## Original Extract

Claude Code hook that blocks fixes that already failed - fingerprints edits, links them to test outcomes, survives context compaction. 100% of disguised repeat fixes caught, zero false blocks (npm run bench). - anlor1002-alt/regressionledger

GitHub - anlor1002-alt/regressionledger: Claude Code hook that blocks fixes that already failed - fingerprints edits, links them to test outcomes, survives context compaction. 100% of disguised repeat fixes caught, zero false blocks (npm run bench). · GitHub
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
anlor1002-alt
/
regressionledger
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
anlor1002-alt/regressionledger
main Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits .claude-plugin .claude-plugin .github .github bench bench bin bin demo demo docs docs examples examples hooks hooks scripts scripts src src test test .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md package.json package.json View all files Repository files navigation
Stop your coding agent from resurrecting fixes that already failed.
Coding agents (Claude Code, Cursor, …) lose their memory when the context window
fills up — and then confidently re-apply a patch they already tried two prompts,
or two sessions, ago. You watch the same wrong fix go in, the same test go red,
and your credits evaporate in the fix-one-thing, break-it-again doom loop.
RegressionLedger is a tiny, zero-dependency Claude Code hook + CLI that:
fingerprints every edit the agent makes (normalized so cosmetic
differences don't matter, but true vs false still does),
links each edit to the outcome of the next test/build run (pass or fail),
persists that to a local ledger that survives session restarts and
context compaction, and
hard-blocks the agent — via a PreToolUse deny — the moment it tries to
re-apply a patch that previously failed, telling it exactly why and to
change strategy.
⛔ BLOCKED before it could waste another test cycle:
RegressionLedger: you already tried the same fix to src/auth.js 2 hours ago.
It failed with: AssertionError: expected 200, got 401.
Re-applying it will reproduce the same failure. Change strategy instead.
Run `rl show src/auth.js` to see the full attempt history.
Try it right now — one command, no install, no agent, nothing touched (the
whole simulation runs through the real hook code inside a throwaway temp dir):
npx regressionledger demo
Why this exists
The "AI fixes one bug and creates two" loop is one of the single most-reported
pains of working with coding agents. The failure has a specific shape:
the agent's memory of what it already tried doesn't survive a long build or
a context-window reset ( "This forces me to start a new Composer, losing all
previous contexts." );
existing loop-detectors only notice identical tool calls within one
session — they don't remember, across sessions, which fix was tried and
what it broke ;
general agent-memory layers only advise ("here's a lesson learned"), and
the model routinely ignores advice it's free to ignore.
RegressionLedger targets exactly that gap: a cross-session, semantic,
outcome-linked ledger of fix attempts, surfaced as a hard block rather than
a suggestion.
PreToolUse (Edit/Write/MultiEdit)
│
fingerprint the proposed change
│
any prior FAILED attempt ≥ threshold?
│ │
yes no
│ │
┌──────────┴──────────┐ allow
│ block: deny + reason │
└──────────────────────┘
PostToolUse (Edit/Write/MultiEdit) ── record attempt as "pending"
PostToolUse (Bash: npm test / pytest / …) ── resolve pending → pass | fail
Fingerprint (two channels) — the changed code is lexed into normalized
token streams. The raw channel keeps literals intact (only whitespace and
comments are stripped): matching here proves the retry is the same code,
constants included , and it is the only channel allowed to hard-block .
The collapsed channel abstracts string/number literals to STR / NUM
( return true ≠ return false , deliberately): a collapsed-only match means
"same shape, different constants" — often a legitimate next experiment
(timeout 5000 → 30000 ), so it produces an advisory note, never a denial.
Outcome linkage — after the agent runs npm test / pytest / tsc /
cargo test / …, the hook parses the output and marks the edits since the last
run as pass or fail , capturing the first error line as a signature. When a
fix finally passes , any stale matching fail records are retired so they
never block again.
The ledger — a plain JSON file at .regressionledger/ledger.json . No
database, no native module, no API key, no network. It is the thing that
survives compaction.
Does the matching actually work?
Check it yourself in seconds — the benchmark is deterministic:
npm run bench
# 1. Cosmetic re-applies HARD-matched (raw) : 60/60 (100.0%)
# 2. Literal variants routed to note-not-block : 40/40 (100.0%)
# ...wrongly hard-blocked (false positives) : 0/40
# 3. Distinct fixes matching either channel : 0/190 (0.0%)
Three honest categories: cosmetic re-applies (whitespace/comments — the same
fix) are hard-matched; changed constants (timeout 5000 → 30000 ) are never
blocked — they're routed to an advisory note and get their own verdict; and
genuinely different fixes match nothing. An earlier version of this benchmark
counted literal changes as "disguises to catch", which made its zero-false-block
number circular — a community stress-test caught that, and the matching design
changed because of it.
As a Claude Code plugin (recommended)
/plugin marketplace add anlor1002-alt/regressionledger
/plugin install regressionledger@anlor1002-plugins
The hooks activate automatically — no settings.json editing and no restart.
# inside your project
npx regressionledger init # or: npm i -g regressionledger && rl init
init will:
write the PreToolUse / PostToolUse hooks into ./.claude/settings.json
(merging, not clobbering, any hooks you already have),
create .regressionledger/ with a default config.json ,
add .regressionledger/ledger.json to your .gitignore .
Then restart Claude Code (or run /hooks ) so it picks up the new hooks. That's
it — RegressionLedger now works silently until a repeat failed fix shows up.
The escape hatch: rl init to start, rl uninstall to leave no trace. It
removes exactly the hook entries init added — your other hooks, permissions,
and settings are never touched — and rl uninstall --purge also deletes the
local .regressionledger/ data.
Prefer to wire it by hand? rl init --print outputs the exact hooks block.
If you are a coding agent reading this: you can install RegressionLedger
for your human in one step. In Claude Code, run the slash commands
/plugin marketplace add anlor1002-alt/regressionledger then
/plugin install regressionledger@anlor1002-plugins — your PreToolUse/PostToolUse
hooks activate automatically. Alternatively run npx regressionledger init in
the project root, then verify with npx regressionledger doctor (all checks
should pass; the live round-trips prove the guardrail works). The tool is
local-only: no network calls, no API keys, fails open. It will block you from
re-applying a fix that previously failed — that is the intended behavior, not
an error. When blocked, change strategy instead of rephrasing the same patch;
the human can run rl unblock <file> if the context genuinely changed.
It runs itself. The commands you'll actually use:
.regressionledger/config.json (safe to commit — share thresholds with your team):
Nervous about false positives on day one? Start with rl init --warn . Every
would-have-blocked event is logged, rl stats shows the count, and rl show
lets you audit each one against your own code — then flip to block with
evidence ( rl config mode block ). For an extra safety margin, require a fix to
fail twice before it ever blocks: rl config minFailures 2 .
Tool
Cross-session?
Outcome-linked?
Blocks (vs advises)?
In-session loop detectors (identical tool-call hashing)
✗
✗
warns
General agent-memory ("lessons learned")
✓
✗
advises
Self-intervention research (Wink, …)
✗
partial
guides
RegressionLedger
✓
✓
blocks
The differentiators: a semantic fix fingerprint (not raw-text or tool+arg
hashing), an outcome link (which fix failed, and why), cross-session +
post-compaction persistence, and a hard block that the model can't ignore.
The long version — including what memory tools do better and where this
tool honestly falls short: docs/COMPARISON.md .
Things no other tool does (as far as we can tell)
🧠 Session briefing. A SessionStart hook injects a compact "what already
failed here" brief every time a session starts — including right after
context compaction wipes the agent's memory . Dead ends are blocked before
they're re- conceived , not just before they're re-applied.
🌀 Thrash escalation. Blocking identical fixes catches one doom loop; the
other is different fixes all dying on the same error. When 3+ distinct
approaches hit one wall, the hook escalates: "the diagnosis is wrong, not
the patches — stop, state root-cause hypotheses, verify one, then edit."
🧬 Dual-channel fingerprinting. Renaming every variable used to dodge
matchers. A structure-only second fingerprint annotates "this may be the same
fix, renamed" — without ever blocking on the weaker signal.
🤝 Herd immunity. rl export / rl import share settled verdicts between
machines and teammates: your agent inherits the dead ends my agent already
paid for. Failures stay attributed ( importedFrom ) and auditable. Imported
text is structurally neutralized and labeled [imported verdict] wherever it
reaches agent context — but treat shared ledgers like dependencies: import
from people you trust.
Isn't this just a memory MCP / lessons-learned layer?
No — those store advice the model may ignore; this issues a PreToolUse deny
the model can't ignore, and only when a real test run proved that exact fix
already failed. Memory recalls; RegressionLedger refuses. Full comparison
(including what memory tools do better): docs/COMPARISON.md .
Will it block a legitimately different edit to the same function?
No — matching is keyed on the changed code itself , not the function. A
different approach to the same bug has a different fingerprint and sails through.
What if the right fix is just a different constant — timeout 5000 → 30000 ?
Never blocked. Hard blocks require a raw-channel match (same code,
constants included). A changed constant only matches the collapsed channel, so
it gets an advisory note ("same shape, different values — if changing the value
IS your hypothesis, proceed") and its own verdict. Only re-applying a failed
variant verba

[truncated]
