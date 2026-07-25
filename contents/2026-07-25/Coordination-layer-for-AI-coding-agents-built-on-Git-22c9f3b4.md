---
source: "https://github.com/zyads/loom-vcs"
hn_url: "https://news.ycombinator.com/item?id=49044517"
title: "Coordination layer for AI coding agents, built on Git"
article_title: "GitHub - zyads/loom-vcs: Version control for many hands moving at once — intent leases, stitches, green-by-construction fabric above git · GitHub"
author: "aether-zyads"
captured_at: "2026-07-25T04:57:19Z"
capture_tool: "hn-digest"
hn_id: 49044517
score: 2
comments: 0
posted_at: "2026-07-25T04:25:57Z"
tags:
  - hacker-news
  - translated
---

# Coordination layer for AI coding agents, built on Git

- HN: [49044517](https://news.ycombinator.com/item?id=49044517)
- Source: [github.com](https://github.com/zyads/loom-vcs)
- Score: 2
- Comments: 0
- Posted: 2026-07-25T04:25:57Z

## Translation

タイトル: Git 上に構築された AI コーディング エージェントの調整層
記事のタイトル: GitHub - zyads/loom-vcs: 同時に移動する多数のハンドのバージョン管理 — インテント リース、ステッチ、git 上のグリーン バイ コンストラクション ファブリック · GitHub
説明: 一度に動く多くの人のバージョン管理 - インテントリース、ステッチ、git 上のグリーンバイコンストラクションファブリック - zyads/loom-vcs

記事本文:
GitHub - zyads/loom-vcs: 一度に動く多数の人のバージョン管理 — インテントリース、ステッチ、git 上のグリーンバイコンストラクションファブリック · GitHub
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
ザド
/
ルームVC
公共
通知
あなたは署名しているに違いありません

通知設定を変更するには編集してください
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github/ workflows .github/ workflows docs docs src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
一度に動く多くの手をバージョン管理します。
Git はファイルシステムです。織機は航空管制です。
織機なし。リポジトリで 2 つの Claude Code セッションをポイントします。エージェントA
新しいセッションで src/auth/ を書き換えるのに 40 分と約 200,000 のトークンを費やします
モデル。 「ログイン フローをクリーンアップする」ように指示されたエージェント B は、次のコマンドで同じファイルにアクセスします。
逆の仮定。どちらももう一方の存在を知りません - git には保存する場所がありません
その知識を入れてください。 B が最後に終了するため、B の編集が勝ちます。 Aさんの仕事は黙々と行われます
木から消えてしまいました。テスト スイートはコミット時に赤色になり、
夕方（そしてさらに数十万トークン）で3回目のセッションを行う
実際には 2 つの AI の半分の書き換えがインターリーブされた破損であると診断します。の
エージェントが行う最もコストのかかることは、重複する 2 つの意味を意味的に調整することです
事後的に書き直します。このワークフローにより、それが日常的に行われます。
ルームと一緒に。 A は「認証を新しいものに移動する」という目的で src/auth/** をリースします。
セッションモデル」。 B のセッションが Loom に何に取り組んでいるかを尋ねたとき — 1 人の MCP
call — リース、アタッチされたゴール、および自己パーティションを確認します。
提案された素のスコープまたは待機。衝突は次の場所で報告されています。
start 、それを処理するときは、 end ではなく、両方の後のツール呼び出しが 1 回かかります。
予算が使われてしまう。とにかく両方とも続行すると、それぞれが独自のワークツリーで動作します。
何も上書きされず、2 番目のランディングではどのファイルが正確に通知されます。
一度移動してリベースすると、コストのかかる調整が発生します。

代わりに珍しい
速い。
これがすべてです。git はトークンが消費された後に衝突を報告します。
Loom は以前にそれらを報告しています。
1 つのリポジトリで複数のコーディング エージェントを実行します。お互いのファイルを上書きし、
赤の上に赤が着地し、タスクの途中で死んでしまい、誰も所有していない汚れたワークツリーが残ります。
誰が何をしているのかが見えないため、すべての衝突は最後に発見されます。
— マージ時または CI — 唯一の修正が高価な再作業である場合。 Git にはありません
インテントの概念: ブランチ名は文字列であるため、エージェントは「ブランチ名は誰ですか」と尋ねることはできません。
すでにこれに取り組んでいますか？」そしてお互いを迂回します。
分離 — すべてのタスクは独自の git ワークツリーを取得します。 2つのスレッド
物理的にお互いの編集を妨害することはできません。
調整 - 編集前に、スレッドはインテント リースを宣言します。
(機械可読な目標 + ファイル グロブ);重複は宣言時に警告されます。
これにより、エージェントは人間の交通警官なしで MCP 上で自己分割して作業できるようになります。
リースは警告しますが、決してブロックしません。エージェントは決して停止しません。唯一の門は
着陸時。
統合 — メインライン (ファブリック) はグリーンのみを通過します。
正直なファイルレベルのマージを使用して、実行と人間/同意ゲートを検証します。
ファブリックとスレッドの両方でファイルが変更された場合、ランディングが拒否されます (「ファブリック」
どちらかの側を上書きするのではなく、自分の下に移動しました — リベース」）。決して赤くない
main は、見知らぬ人の破損を診断するフリート全体のトークンの書き込みがないことを意味します。
回復 — 死んだエージェントのスレッドは孤立します : 目標、受け入れ
基準と数秒前のチェックポイントが添付されており、誰でも要求できます。それ
これはスケジュール可能な作業であり、謎のぶら下がりブランチではありません。
正直な譲歩: 慎重に作業する 2 人の人間には、これは必要ありません。の
値はライターの数と編集頻度に応じて変化します。これはフリートに対して存在します。
エージェント（および彼らと同じように働く人間）。
N 人のエージェントが同じタスクを競っています。 Loom はさまざまなタスクを調整します
に

レポ。 1 つのタスクで重複して試行する場合は、ジャッジ/トーナメントを使用します。
ハーネス - その後、勝者の踊り場をリースします。
gitを置き換えます。織機が上にあります。 git には履歴、リモート、および
ツールを確認します。オプションの橋プロジェクトは、それぞれの陸地の地域に織り込まれます。
git History — デフォルトでは Weave ごとに 1 つのコミット、希望する場合はチェックポイントレベル
(下記のbridge_modeを参照) — そして決してプッシュしません。
普通にPRをする人間。ゆっくりと調整された帯域外の作業が必要です
git+GitHub はすでにうまく機能しています。
災害は自分のマシンで回避されます。 1 つのリポジトリ、2 つのターミナル
(「エージェント」A および B):
カーゴインストール --path 。 # 1 つのバイナリ: 織機
cd あなたのリポジトリ
loom init --verify " Cargo check " # または緑色で 0 で終了するコマンド
答え:
織機リース「フランス語で挨拶」「greeting.txt」
# WORK IN: ~/.loom/<repo>/worktrees/<thread-A> ← A 自身のワークツリー
Bさんは同時にこう言いました。
織機リース「もっと大きな声で挨拶」「greeting.txt」
# TOE-STEP: 「大きな声で挨拶」が「フランス語で挨拶」と重なる
# (リースは警告しますが、ブロックすることはありません - 調整または続行)
# 作業場所: ~/.loom/<repo>/worktrees/<thread-B> ← 別のワークツリー
どちらも同じファイルを独自のワークツリー内で編集するため、何も破壊されません。あ
最初に終了します:
# A (端末を共有する場合は、リース出力から --lease/--thread ID を使用します):
織機ステッチ # チェックポイント (コンテンツに対処; 削除は追跡)
loom提案 # スクラッチコピーで実行を検証します。緑色ではy/Nを尋ねます
# → 織物: 1 ファイル適用
B は次に提案し、サイレント上書きの代わりに真実を取得します。
織機ステッチ && 織機提案
# → 緑色を確認します…しかし着陸は拒否されます:
# ファブリックがgreeting.txtであなたの下に移動しました — スレッドをリベースして再提案します
織機リベース
# 競合 — ファブリックとこのスレッドの両方が変更されました:
#greeting.txt (バージョンはワークツリーに保存されます。ファブリックはリポジトリにあります)
# ファイルを調整する

B のワークツリーでは、次のようになります。
織機ステッチ && 織機提案
# → 織り — B の着地は A の作業を消去する代わりに含めます
タスクの途中でセッションを強制終了すると（またはリース TTL が経過すると）、ステータスが表示されます
スレッドは孤児として — 採用 <thread-id> が次のエージェントに渡されます
その目標、基準、ワークツリー、チェックポイントはそのままです。
エージェント (クロード コードまたは任意の MCP クライアント)
クロード mcp 織機を追加 -- 織機 mcp
エージェントが従うツール フローは次のとおりです。
loom_status — 誰が何に取り組んでいるか (スレッド、ライブリース、孤立)。
loom_lease — 目標と範囲を宣言します。応答の working_dir は
スレッドのワークツリー: そこに cd して、その中ですべての編集を行います。トーステップ
警告が同じ応答で返されます。これはスコープを再ネゴシエートする瞬間です。
work → loom_stitch 数回​​の編集ごと (チェックポイント + リースのハートビート)。
loom_propose — スクラッチ コピーで検証を実行し、緑/赤をレポートします。
申請は常に MCP 経由で拒否されます: 着陸にはターミナルで人間が必要です
( loom提案 )または承認キューのあるホスト。提案は検証します。
それは決して着陸しません。
「ファブリックの移動」: loom_rebase で、ワークツリー内の競合を調整します。
ステッチして、再提案します。 loom_adopt は孤立者 (ローカルまたは同期されたもの) を選択します。
同僚のもの）。
友達と 5 分で仕事ができる
2 人、1 つのプライベート GitHub リポジトリ。 Loom は非表示を介して同期します
refs/loom/* refs — サーバー、デーモン、ホストに新しいものはありません。
# 二人とも、自分のクローンに一度:
loom init --verify " npm test " # プロジェクトにとって緑が意味するものは何でも
loom sync --remote Origin # このリポジトリのリモートを記憶します
次に、クイックスタートとまったく同じように作業します。リース、ワークツリーでの編集、
ステッチ、提案 - 状態を交換したいときはいつでも、loom sync を実行します
(または、loom sync --auto を選択して、ステッチ/プロポーズごとに同期します):
あなた > ルームリース " ダークモードスタイル " ' style.css ' # + 編集、ステッチ
友達＞私

oom リース " コンパクトヘッダー " ' style.css ' # + 編集、ステッチ
あなた > 提案する && 同期する # 土地、公開する
友人 > loom sync # あなたの織りを彼らのツリーに取り込みます
友人 > 織機が # グリーンを提案する — しかし着陸すると言う: 布地があなたの下に移動しました
友人 > loom rebase # 変更を早送りします。本当の対立は続いた
友人 > 織機ステッチ && 織機提案 && 織機同期 # 結合された結果の土地
あなた > loom sync # 両方のツリーが同一になり、両方の変更が存在します
2 つの文でマージの恐れはゼロです。どちらも他方を上書きすることはできません。
編集はスレッドごとのワークツリー内に存在し、共有ラインは前進するだけであるためです。
比較と交換の ref プッシュを介して、それがあなたの下に移動したときに拒否されます。
最悪のケースはツリーが壊れることではなく、ファイル名によって次のように伝えられます。
仕事が始まる前に何を見るべきか。
同期が共有するものは、明確に言うと、織機のメタデータ (目標、範囲、所有者) と
ステッチされたスコープのファイルコンテンツはそのリモートに移動します - 同じです
そこに枝を突き刺すような露出。 --auto はリポジトリごとにオプトインされます。死んだ
マシンのスレッドは養子縁組可能な孤児として表示されます。主張は先押し勝ちです
git ref で、敗者には誰が勝ったかが通知されます。
なぜ単に git merge/rebase をしないのでしょうか?
コンクリートの破損が1件あります。エージェント A と B は両方ともブランチ上の config.rs を編集します。
マージ時間: git は同じ行が触れられていることを確認し、競合を宣言し、マージします。
最後にマージした人に送信されます。どちらの記憶もない新しいエージェント セッションです。
意図し、2 つの書き換えをトークンごとに調整します。さらに悪いことに、彼らは接触した
異なる行、git はサイレントに自動マージし、メインは破損して赤になります
どちらの作者も単独で複製することはできません。その後、すべてのエージェントが赤のメインをプルして燃えます
ワークフローによって生成された障害に関するトークン。
Loom の答えは、ポイントごとに次のとおりです。
重複は少なくとも時点ではわかっていました。 git にはそれを記録する場所がありませんでした。織機
いずれかのエージェントより先に報告する

何でも費やします。
和解が必要な場合、和解は生き残った著者（誰が行うのか）に行われます。
コンテキストを持っています)、スコープは名前付きファイル、グリーン ツリーに対してです。
赤いものに対してマージ時に見知らぬ人。
リベースは高速ではなくまれに行われます: リース時間での自己パーティショニングは意味します
ほとんどの衝突は決して起こりません。
着陸は同意ゲート制です。緑色の検証は証拠であり、アクションではありません。
端末は y/N を尋ねます。 MCP は常に適用を拒否します。ホストパークの埋め込み
承認。 「はい」がなければリポジトリには何も到達しません。
削除は第一級です。ワークツリー内で削除したファイルは、
墓石として記録され、削除として着陸し、削除と編集の衝突が発生します。
他の紛争と同様に拒否してください。
大きなファイルは大音量でスキップされます。 8 MiB を超えるファイルはスナップショットが作成されません。
ステッチ結果ではそれぞれに名前が付けられます。でキャップを上げます
必要に応じて LOOM_MAX_FILE_MB になります。リポジトリのルートに .loomignore を配置することを推奨します
(gitignore-lite: リテラル ディレクトリと単純なグロブ、1 行に 1 つ) — 拡張します
組み込みの .git/target/node_modules は除外し、無視を解除できません
彼ら。
織機クリーンは、織られた糸のワークツリーを除去します。
ワークツリーには、ステッチに取り込まれていないものはすべて保持されます。
ハード タイムアウトを使用して、スクラッチ コピー内のリポジトリ全体であることを確認します。レッドネバー
土地。小切手を実行することさえできないゲートは、沈黙ではなく赤を報告します。
Git ヒストリー

[切り捨てられた]

## Original Extract

Version control for many hands moving at once — intent leases, stitches, green-by-construction fabric above git - zyads/loom-vcs

GitHub - zyads/loom-vcs: Version control for many hands moving at once — intent leases, stitches, green-by-construction fabric above git · GitHub
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
zyads
/
loom-vcs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows docs docs src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
Version control for many hands moving at once.
Git is the filesystem; Loom is air traffic control.
Without Loom. You point two Claude Code sessions at your repo. Agent A
spends 40 minutes and ~200k tokens rewriting src/auth/ around a new session
model. Agent B, told to "clean up the login flow", touches the same files with
the opposite assumption. Neither knows the other exists — git has nowhere to
put that knowledge. B finishes last, so B's edits win; A's work is silently
gone from the tree. The test suite goes red at commit time, and you spend the
evening (and another few hundred thousand tokens) having a third session
diagnose a breakage that is really two AIs' half-rewrites interleaved. The
most expensive thing agents do is semantically reconcile two overlapping
rewrites after the fact — and this workflow makes it routine.
With Loom. A leases src/auth/** with the goal "move auth to the new
session model". When B's session asks Loom what's being worked on — one MCP
call — it sees that lease, goal attached, and self-partitions: it either takes
the suggested disjoint scope or waits. The collision is reported at the
start , when handling it costs one tool call, not at the end , after both
budgets are spent. If both proceed anyway, each works in its own worktree,
nothing is overwritten, and the second landing is told exactly which files
moved and rebases once — the expensive reconciliation becomes rare instead of
fast.
That is the entire bet: git reports collisions after the tokens are spent;
Loom reports them before.
You run several coding agents on one repo. They overwrite each other's files,
land red on top of red, and die mid-task leaving dirty worktrees nobody owns.
You can't see who is doing what, so every collision is discovered at the end
— at merge or in CI — when the only fix is expensive re-work. Git has no
concept of intent : a branch name is a string, so agents can't ask "who is
already working on this?" and route around each other.
Isolation — every task gets its own git worktree; two threads
physically cannot clobber each other's edits.
Coordination — before editing, a thread declares an intent lease
(machine-readable goal + file globs); overlaps warn at declaration time,
which lets agents self-partition work over MCP with no human traffic cop.
Leases warn, they never block — an agent is never stopped; the only gate is
at landing.
Integration — the mainline ( fabric ) advances only through a green
verify run plus a human/consent gate, with an honest file-level merge:
a file changed in both the fabric and the thread refuses to land ("fabric
moved under you — rebase") instead of overwriting either side. Never-red
main means no fleet-wide token burn diagnosing a stranger's breakage.
Recovery — a dead agent's thread becomes an orphan : goal, acceptance
criteria, and a seconds-old checkpoint attached, claimable by anyone. That
is schedulable work, not a mystery dangling branch.
An honest concession: two careful humans working slowly don't need this. The
value scales with writer count and edit frequency — it exists for fleets of
agents (and humans who work like them).
N agents racing the SAME task. Loom coordinates different tasks on
one repo. For redundant attempts at one task, use a judge/tournament
harness — then lease the winner's landing.
Replacing git. Loom sits on top; git stays your history, remotes, and
review tooling. The optional bridge projects each landed weave into local
git history — one commit per weave by default, checkpoint-level if you ask
(see bridge_mode below) — and never pushes.
Humans doing normal PRs. Slow, coordinated-out-of-band work is what
git+GitHub already does well.
The disaster being prevented, on your own machine. One repo, two terminals
("agents" A and B):
cargo install --path . # one binary: loom
cd your-repo
loom init --verify " cargo check " # or any command that exits 0 on green
A:
loom lease " greet in french " ' greeting.txt '
# WORK IN: ~/.loom/<repo>/worktrees/<thread-A> ← A's own worktree
B, at the same time:
loom lease " greet louder " ' greeting.txt '
# TOE-STEP: your 'greet louder' overlaps 'greet in french'
# (a lease warns, it never blocks — coordinate or continue)
# WORK IN: ~/.loom/<repo>/worktrees/<thread-B> ← a DIFFERENT worktree
Both edit the same file — in their own worktrees, so nothing clobbers. A
finishes first:
# A (use --lease/--thread ids from the lease output when sharing a terminal):
loom stitch # checkpoint (content-addressed; deletions tracked)
loom propose # verify runs in a scratch copy; on green it asks y/N
# → woven: 1 files applied
B proposes next, and gets the truth instead of a silent overwrite:
loom stitch && loom propose
# → verify green … but landing refuses:
# fabric moved under you on greeting.txt — rebase the thread and re-propose
loom rebase
# CONFLICTS — both the fabric and this thread changed:
# greeting.txt (your version kept in the worktree; the fabric's is in the repo tree)
# reconcile the file in B's worktree, then:
loom stitch && loom propose
# → woven — B's landing includes A's work instead of erasing it
Kill a session mid-task (or let its lease TTL lapse) and loom status shows
the thread as an orphan — loom adopt <thread-id> hands the next agent
its goal, criteria, and worktree, checkpoint intact.
Agents (Claude Code or any MCP client)
claude mcp add loom -- loom mcp
The tool flow an agent follows:
loom_status — who's working on what (threads, live leases, orphans).
loom_lease — declare goal + scope. The response's working_dir is the
thread's worktree: cd there and make all edits in it. Toe-step
warnings come back in the same response — the moment to renegotiate scope.
work → loom_stitch every few edits (checkpoints + heartbeats the lease).
loom_propose — runs the verify in a scratch copy and reports green/red.
The apply is always refused over MCP: landing takes a human at a terminal
( loom propose ) or a host with an approvals queue. Proposing verifies;
it never lands.
On "fabric moved": loom_rebase , reconcile any conflicts in the worktree,
stitch, re-propose. loom_adopt picks up orphans (local or a synced
peer's).
Work with a friend in 5 minutes
Two people, one private GitHub repo. Loom syncs through it over hidden
refs/loom/* refs — no server, no daemon, nothing new to host.
# Both of you, once, in your own clone:
loom init --verify " npm test " # whatever green means for the project
loom sync --remote origin # remembers the remote for this repo
Then work exactly as in the quickstart — lease, edit in your worktree,
stitch, propose — and run loom sync whenever you want to exchange state
(or opt in to loom sync --auto to sync after every stitch/propose):
you > loom lease " dark mode styles " ' style.css ' # + edit, stitch
friend > loom lease " compact header " ' style.css ' # + edit, stitch
you > loom propose && loom sync # lands, publishes
friend > loom sync # pulls your weave into their tree
friend > loom propose # green — but landing says: fabric moved under you
friend > loom rebase # your change fast-forwards in; real conflicts kept
friend > loom stitch && loom propose && loom sync # merged result lands
you > loom sync # both trees now identical, both changes present
Zero merge fear in two sentences: neither of you can overwrite the other,
because edits live in per-thread worktrees and the shared line only advances
through a compare-and-swap ref push that refuses when it moved under you.
The worst case is not a broken tree — it is being told, by name of file,
what to look at before your work lands.
What sync shares, plainly: your loom metadata (goals, scopes, holders) and
the file content of your stitched scope go to that remote — the same
exposure as pushing a branch there. --auto is opt-in per repo. Dead
machines' threads show up as adoptable orphans; claims are first-push-wins
on a git ref, and the loser is told who won.
Why not just git merge/rebase?
One concrete breakage. Agents A and B both edit config.rs on branches.
Merge time: git sees the same lines touched, declares a conflict, and hands
it to whoever merges last — a fresh agent session with no memory of either
intent, reconciling two rewrites token-by-token. Or worse: they touched
different lines, git auto-merges silently, and main is red with a breakage
neither author can reproduce alone. Every agent then pulls red main and burns
tokens on a failure that was manufactured by the workflow.
Loom's answer, point by point:
The overlap was knowable at lease time; git had nowhere to record it. Loom
reports it before either agent spends anything.
Reconciliation, when it must happen, goes to the surviving author (who
has the context), scoped to named files, against a green tree — not to a
stranger at merge time against a red one.
Rebases are made rare , not fast: self-partitioning at lease time means
most collisions never happen.
Landing is consent-gated. A green verify is evidence, not an action.
The terminal asks y/N; MCP always refuses the apply; embedding hosts park
an approval. Nothing reaches the repo tree without a yes.
Deletions are first-class. A file you delete in your worktree is
recorded as a tombstone, lands as a deletion, and delete-vs-edit collisions
refuse like any other conflict.
Big files are skipped, loudly. Files over 8 MiB are not snapshotted;
each one is named in the stitch result. Raise the cap with
LOOM_MAX_FILE_MB if you must; prefer a .loomignore at the repo root
(gitignore-lite: literal dirs and simple globs, one per line) — it extends
the built-in .git / target / node_modules excludes and cannot un-ignore
them.
loom clean removes worktrees of woven threads — and refuses if the
worktree holds anything not captured in a stitch.
Verify is whole-repo, in a scratch copy, with a hard timeout. Red never
lands; a gate that can't even stage its check reports red, not silence.
Git histor

[truncated]
