---
source: "https://github.com/lenamonj/jeffy-loop"
hn_url: "https://news.ycombinator.com/item?id=49063529"
title: "Jeffy Loop, an autonomous code improvement loop for Claude Code"
article_title: "GitHub - lenamonj/jeffy-loop: An autonomous improvement loop for Claude Code: audit, fix, verify, converge - one verified task per iteration, checkpointed in local commits it never pushes. · GitHub"
author: "lenamonj"
captured_at: "2026-07-26T23:53:14Z"
capture_tool: "hn-digest"
hn_id: 49063529
score: 2
comments: 0
posted_at: "2026-07-26T23:45:50Z"
tags:
  - hacker-news
  - translated
---

# Jeffy Loop, an autonomous code improvement loop for Claude Code

- HN: [49063529](https://news.ycombinator.com/item?id=49063529)
- Source: [github.com](https://github.com/lenamonj/jeffy-loop)
- Score: 2
- Comments: 0
- Posted: 2026-07-26T23:45:50Z

## Translation

タイトル: Jeffy Loop、Claude Code の自律的なコード改善ループ
記事のタイトル: GitHub - lenamonj/jeffy-loop: クロード コードの自律的改善ループ: 監査、修正、検証、収束 - 反復ごとに 1 つの検証済みタスク、プッシュされないローカル コミットでチェックポイントが設定されます。 · GitHub
説明: クロード コードの自律的な改善ループ: 監査、修正、検証、収束 - 反復ごとに 1 つの検証済みタスク、ローカル コミットでチェックポイントが設定され、プッシュされることはありません。 - レナモンジ/ジェフィーループ

記事本文:
GitHub - lenamonj/jeffy-loop: クロード コードの自律的な改善ループ: 監査、修正、検証、収束 - 反復ごとに 1 つの検証済みタスク、プッシュされないローカル コミットでチェックポイントが設定されます。 · GitHub
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
アラートを閉じる
{{ メッセージ }}
レナ

もんじ
/
ジェフィーループ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
23 コミット 23 コミット .github .github evals evals メディア メディア スクリプト スクリプト スキル スキル .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスREADME.md README.md SECURITY.md SECURITY.md install.ps1 install.ps1 install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
それをプロジェクトに向けます。予算を与えてください。より良いコードベースとレポートに戻ってください。
Jeffy Loop は、Claude Code の自律的な改善ループであり、規律ある主任エンジニアが行うようにコードベースで機能します。まず監査し、実際の影響に基づいて優先順位を付け、検証済みのタスクを一度に 1 つずつ修正し、ジョブが実際に完了したら停止します。
任意のプロジェクトで /jeffy 10 を実行し、終了します。 Jeffy は、適用されるすべての品質ディメンションを監査し、すべてのタスクが実行可能な受け入れチェックを行うバックログを作成し、それを焼き尽くします。反復ごとに 1 つのタスクが実行され、それぞれが検証され、それぞれがチェックポイントされます。完了すると、何が変更されたのか、何ができなくなったのか、そして何が決定を必要としているのかが正確に表示されます。
Claude Code (一度インストールしてサインイン) と git が必要です。インストーラーは、 jq を含むその他すべてを処理します。
git クローン https://github.com/lenamonj/jeffy-loop.git
CD ジェフィーループ
./install.sh # Windows PowerShell: .\install.ps1
次に、改善したいプロジェクトの Claude Code を開き、 /jeffy 10 と入力します。これは、シェル コマンドではなく、クロード コード セッション内のスラッシュ コマンドです。 Windows 実行ポリシーのメモを含む詳細については、「インストール」を参照してください。
1 つのコマンドが実行される仕組み。実線の矢印は制御フローです。破線の矢印は

ファイルはそれを制御する読み取りと書き込みを行います。実行外では、Stop フックは即座に終了します (ライブ状態ファイルや動作はありません)。また、/cancel-jeffy はいつでも実行を終了します。図のソース: media/flowchart.mmd 。
直接対決と生のプロンプト ループ。すべての行は、コード内で検証できることを保証します。エンジンは skill/jeffy/hooks/stop-hook.sh 、分野は skill/jeffy/references/iteration-prompt.txt で、実際の実行からのレシートは evals/ の下に保存されます。 HD で視聴します。
Jeffy Loop は、コーディング エージェントがループ内の 1 つのプロンプトを再入力するという Geoffrey Huntley の洞察である Ralph テクニックから派生し、実際の作業に組み込まれます。上記の比較はエンジンとメソッドです。生のループは Jeffy が構築されるエンジン パターンであり、Jeffy はそれにラップされたエンジニアリング メソッドです。このメソッド自体は、ループを大規模に実行している人々が公開しているもの (Anthropic の Claude Code ベスト プラクティスと Boris Cherny のパブリック ワークフロー) を抽出したものです。エージェントに一度に 1 タスクずつ実行できるかどうかのチェックを与え、苦労して得たすべてのレッスンを次の実行で読み取るファイルにプロモートし、1 回の長い実行よりも新鮮なコンテキストの小さな実行を優先します。
リンターの判断ではなく、エンジニアの判断です。すべての実行は、アーキテクチャ、正確性、セキュリティ、テスト、エラー処理、パフォーマンス、アクセシビリティ、開発者エクスペリエンスなどの実際の監査から始まり、すべての発見事項は、具体的な受け入れチェックを伴う優先順位の高いタスクになります。アサーションよりも証拠: ループがそれを指すことができる場合にのみ、発見が存在します。
リポジトリを破壊することはできません。すべての反復はローカル チェックポイントのコミットで終了します。リポジトリレベルの検証ゲートは反復ごとに実行され、プロジェクトを中断する反復はその場で元に戻されます。何もプッシュされず、ブランチも作成されません。 git log で確認し、単一の反復を元に戻し、満足したら実行を破棄します。
入ってないよ

通気口の問題。重大度は、宣言された動作範囲、つまり想像上の攻撃者ではなく、プロジェクトの実際の入力面に基づいて判断されます。範囲外の発見によってバックログが膨らむことはなく、範囲を広げることができるのはあなただけです。ループは提案をファイルして次に進みます。
Doneは完了という意味です。ループは、完全な新しい監査で高および中検出がゼロで、バックログが空の場合にのみ収束します。すべての低は修正されるか、理由を付けて明示的に拒否されます。 「小さな」問題が静かに残される痕跡はありません。そして、宣言には副署名が付いています。自分の作業を採点するエージェントがそれを賞賛するため、ループが収束を主張する前に、新しいコンテキストの懐疑的な評価者 (実行の自己説得をまったく持たないサブエージェント) が検証ゲートとクローズされたタスクの受け入れチェックを再実行し、実行の変更で見逃した結果を探し、PASS を返さなければなりません。拒否された場合は、その証拠が新しいタスクとしてファイルされ、実行が続行されます。そして、Promise 自体はマシンチェックされます。Stop フック (モデルではなくプレーンなシェル) は、タスク台帳が空でない限り収束停止を拒否し、記録された Converged commit がツリーを証明し (ループ状態がそれ以降変更された以外は何もありません)、フックがタイムアウトで再実行すると、プロジェクト自身の verify コマンドは緑色で終了します。チェックに失敗すると、実行を終了させるのではなく、ループに証拠を再供給します。
いつ停止すべきかを知っています。予算が消費された、収束に達した、進行が停滞した、またはあなただけが下せる決定 - ループは自動的に終了し、予算を消費する代わりにその理由を示します。
すべての実行はレポートで終了します。使用された反復、重大度によって閉じられたタスク、実行の差分統計、ブロックされたもの、待機中の決定。追加専用ジャーナルとチェックポイント コミットには、grep 可能な完全なレコードが保持されます。
クロード コード セッションで /jeffy を実行します。
ブートストラップ thr

プロジェクト ルートにある ee ファイル: PLAN.md (目標、動作エンベロープ、メソッド、検証コマンド、完了の定義)、BACKLOG.md (優先順位付けされたタスク、最悪の重大度が最初)、JOURNAL.md (追加のみの反復ログ)。これらは実行間で持続します。これらはループのメモリです。
予算付きループを開始します。各反復では、プロジェクトを監査するか、バックログ タスクを 1 つだけ実行してその受け入れチェックを検証します。すべての反復で verify コマンドが実行され、ローカル チェックポイントのコミットで終了します。
予算、収束時 (クリーンな監査、空のバックログ)、停止時、またはキャンセル時に停止し、実行レポートで終了します。
Jeffy は、それ自体で実行してこのリポジトリを構築しました。それが書いた開発ジャーナルは公開されたツリーから外れています - 状態ファイルは製品ではなくループのメモリです - しかし、2 つの要約されたエントリは実行の形状を示し、書かれたものとして示されています (ジャーナルの見出しはそれ以来、ループが現在使用しているパイプ区切りの文法に強化されています)。まず、バックログを生成した開始監査:
## 反復 1 - 2026-07-03 - 監査 (改善モード)
監査スコア (該当するディメンションごとの最高の検出重大度):
- テスト: 中。インストーラーやスキルの自動検証はゼロです。構文
SKILL.md フロントマターが壊れているか欠落している場合は、サイレントに出荷されます (M2)。
- Git の衛生状態: 中。 .gitignore はループの一時的なものを除外しません
Jeffy を実行するたびに作成される、セッション スコープの状態ファイル (M1)。
- セキュリティ: なし。信頼できる CLI を超えたネットワークフェッチはありません。秘密はありません。
(さらに 10 個の次元が低、なし、または該当なしとスコア付けされました)
BACKLOG.md に書き込まれる結果: 0 高、2 中、4 低。
次に: M1 (ループ状態ファイルの gitignore) をブロックされていない項目の先頭として実行します。
次に、後の実行で終了した完全な新証拠監査:
## 反復 1 (実行 6、予算 5) - 2026-07-05 - 完全な監査 (収束チェック)
この反復で収集された証拠

(新鮮):
- バリデータ: bash scripts/validate.sh は 0 で終了し、すべてのチェックは緑色になります。
- チェック 6 に歯がある: スクラッチ コピーでのネガティブ パス テスト
「## 動作エンベロープ」マーカーが壊れているため、ビルドに失敗します。サイレントノーオペレーションではありません。
(さらに 7 つの証拠行)
結果: ゼロ高、ゼロ中。完了の定義は純粋に、
検証可能な真実。完全なコミットハッシュを含む収束行を記録しました
## の下で BACKLOG.md に収束されるため、今後は変更されていないツリーで再起動されます
再監査する代わりに O(1) をラチェットします。
このコンバージェンスはアーカイブされるのではなく、再獲得されるものです。このリポジトリで Jeffy を新たに実行するたびに、新しい証拠とともに再び到達する必要があります。 Jeffy がプロジェクトに収束すると、チェックポイントが git ログに記録され、ループのバックログの ## Converged の下に記録されるため、変更されていないツリーで再起動され、再監査ではなく再検証されます。自分のプロジェクトで /jeffy を実行し、プロジェクトが残した日記を読んでください。
見知らぬ人のコードベースで実証済み
セルフランはイージーモードです。そこで Jeffy は、関連性のない、有名な 4 つの実際のプロジェクトを指摘されました。それぞれがローカル クローン内にあり、上流にプッシュされるものは何もなく、すべての実行が同じルールの下で収束します。つまり、ファイルする前の証拠、反復ごとに 1 つの検証済みタスク、チェックポイント、修正とともにオペレーターの間違いを記録するジャーナルです。
kennethreitz/レコード (約 7.2,000 個の星)。上流の HEAD では、pytest は 31 が合格したと報告します。同じ HEAD では、INSERT は静かにデータを失い、transaction() はすべての例外を飲み込み、すべてのクエリはプールされた接続をリークします。 Jeffy は、緑色のテスト スイートの背後に隠れていた 4 つの重大度の高いバグを再現し、共有する境界で 1 つの構造修正で 3 つを閉じ、アップストリームが作成されたその日に元に戻した修正を復元し、古いコードで失敗することが証明された回帰スイートを残しました。repro.py はアップストリームの HEAD のバグを表示し、fixes.patch はそれらのバグが修正されたことを示します。アル

そのうちの 4 件は、kennethreitz/records#236 で再現と PR オファーとともに上流で開示されました。何かをマージするのはメンテナの要求です。
janl/mustache.js (約 16.7,000 個のスター)。現在のノードのアップストリーム HEAD では、テスト スイートを開始できません。放棄された esm シムは単一のアサーションの前にクラッシュし、bin/mustache は完全にクラッシュします。 Jeffy は、3 つのローディング サイトすべてで 1 つの構造修正を行ってゲートを復活させ、CLI で再現された 2 つ目の正確性バグを回帰テストで修正し、機能しなくなったブラウザ テスト スタックを削除し、ツールチェーンを最新化して、npm Audit の脆弱性 107 件 (重大 24 件) をスイートの 297 件で最低 2 件に改善しました (公式の Mustache 仕様への準拠も含まれます)。その後、最終監査は、実行自体の以前の作業 (ドキュメントはまだ削除されたスタックを指している) に対してメディアをファイルし、収束を宣言する前に修正しました。 4 つの調査結果はすべて、 janl/mustache.js#848 の再現と PR オファーとともに上流で開示されました。
sivel/speedtest-cli (~13,000 スター)。対照的なケースは、2021 年以降休止状態にあるものの、基本的には健全であり、正直な結果は小さな発見と自制にとどまる。 Jeffy は、プロジェクト独自の lint ゲートを修正しました (8 つの Python 2 組み込み誤検知による変更されていないコードを赤で示しました)。

[切り捨てられた]

## Original Extract

An autonomous improvement loop for Claude Code: audit, fix, verify, converge - one verified task per iteration, checkpointed in local commits it never pushes. - lenamonj/jeffy-loop

GitHub - lenamonj/jeffy-loop: An autonomous improvement loop for Claude Code: audit, fix, verify, converge - one verified task per iteration, checkpointed in local commits it never pushes. · GitHub
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
lenamonj
/
jeffy-loop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
23 Commits 23 Commits .github .github evals evals media media scripts scripts skills skills .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
Point it at a project. Give it a budget. Come back to a better codebase and a report.
Jeffy Loop is an autonomous improvement loop for Claude Code that works on your codebase the way a disciplined principal engineer would: audit first, prioritize by real impact, fix one verified task at a time, and stop when the job is actually done.
Run /jeffy 10 in any project and walk away. Jeffy audits every quality dimension that applies, writes a backlog where every task carries a runnable acceptance check, then burns through it - one task per iteration, each one verified, each one checkpointed. When it finishes, it tells you exactly what changed, what it couldn't do, and what needs your decision.
You need Claude Code (installed and signed in once) and git. The installer handles everything else, including jq .
git clone https://github.com/lenamonj/jeffy-loop.git
cd jeffy-loop
./install.sh # Windows PowerShell: .\install.ps1
Then open Claude Code in the project you want to improve and type /jeffy 10 . It is a slash command inside the Claude Code session, not a shell command. Details, including the Windows execution-policy note, are under Install .
How one command becomes a run. Solid arrows are control flow; dashed arrows are the file reads and writes that steer it. Outside a run the Stop hook exits instantly - no live state file, no behavior - and /cancel-jeffy ends a run at any time. Diagram source: media/flowchart.mmd .
The head-to-head vs a raw prompt loop. Every row is a guarantee you can verify in the code: the engine is skills/jeffy/hooks/stop-hook.sh , the discipline is skills/jeffy/references/iteration-prompt.txt , and the receipts from real runs live under evals/ . Watch in HD .
Jeffy Loop descends from the Ralph technique , Geoffrey Huntley's insight that a coding agent re-fed one prompt in a loop compounds into real work. The comparison above is engine versus method: the raw loop is the engine pattern Jeffy is built on, and Jeffy is the engineering method wrapped around it. The method itself distills what the people running loops at scale have published - Anthropic's Claude Code best practices and Boris Cherny's public workflow : give the agent a check it can run, one task at a time, promote every hard-won lesson into a file the next run reads, and prefer small fresh-context runs over one long one.
An engineer's judgment, not a linter's. Every run starts from a real audit - architecture, correctness, security, testing, error handling, performance, accessibility, developer experience, and more - and every finding becomes a prioritized task with a concrete acceptance check. Evidence over assertion: a finding exists only if the loop can point at it.
It cannot wreck your repo. Every iteration ends in a local checkpoint commit. A repo-level verify gate runs every iteration, and an iteration that breaks the project is reverted on the spot. Nothing is ever pushed, no branches are created. Review with git log , revert any single iteration, squash the run when you're happy.
It doesn't invent problems. Severity is judged against a declared operating envelope - your project's real input surfaces, not imagined attackers. Out-of-envelope findings can't inflate the backlog, and only you can widen the envelope: the loop files a proposal and moves on.
Done means done. The loop converges only when a full fresh audit finds zero High and zero Medium findings and the backlog is empty - every Low either fixed or explicitly declined with a reason. No trail of "minor" issues quietly left behind. And the declaration is countersigned: an agent grading its own work praises it, so before the loop may claim convergence, a fresh-context skeptical evaluator - a sub-agent carrying none of the run's self-persuasion - re-runs the verify gate and the closed tasks' acceptance checks, hunts for missed findings in the run's changes, and must return PASS. A rejection files its evidence as new tasks and the run continues. And the promise itself is machine-checked: the Stop hook - plain shell, not a model - refuses the converged stop unless the task ledger is empty, the recorded Converged commit certifies the tree (nothing but loop state changed since it), and the project's own verify command exits green when the hook re-runs it under a timeout. A failed check re-feeds the loop with the evidence instead of letting the run end.
It knows when to stop. Budget spent, convergence reached, progress stalled, or a decision only you can make - the loop ends itself and says why, instead of burning budget spinning.
Every run ends with a report. Iterations used, tasks closed with severities, the run's diffstat, anything blocked, and decisions waiting on you. An append-only journal and the checkpoint commits hold the full, greppable record.
Running /jeffy in a Claude Code session:
Bootstraps three files at the project root: PLAN.md (goal, operating envelope, method, verify command, definition of done), BACKLOG.md (prioritized tasks, worst severity first), JOURNAL.md (append-only iteration log). They persist between runs - they are the loop's memory.
Launches a budgeted loop. Each iteration either audits the project or executes exactly one backlog task and verifies its acceptance check; every iteration runs the verify command and ends in a local checkpoint commit.
Stops at the budget, at convergence (clean audit, empty backlog), on a stall, or when you cancel - and closes with the run report.
Jeffy built this repository by running on itself. The dev journal it wrote stays out of the published tree - state files are the loop's memory, not the product - but two abridged entries show the shape of a run, shown as written (journal headings have since tightened to the pipe-delimited grammar the loop uses today). First, the opening audit that generated the backlog:
## Iteration 1 - 2026-07-03 - Audit (Improvement mode)
Audit scores (highest finding severity per applicable dimension):
- Testing: Medium. Zero automated validation of installers or skills; a syntax
break or missing SKILL.md frontmatter would ship silently (M2).
- Git hygiene: Medium. .gitignore does not exclude the loop's transient
session-scoped state file, which every Jeffy run creates (M1).
- Security: None. No network fetch beyond trusted CLIs; no secrets.
(10 more dimensions scored Low, None, or N/A)
Findings written to BACKLOG.md: 0 High, 2 Medium, 4 Low.
Next: Execute M1 (gitignore the loop state file) as the top unblocked item.
Then, the full fresh-evidence audit that ended a later run:
## Iteration 1 (run 6, budget 5) - 2026-07-05 - Full audit (convergence check)
Evidence gathered this iteration (fresh):
- Validator: bash scripts/validate.sh exits 0, every check green.
- Check 6 has teeth: negative-path test on a scratch copy with the
"## Operating envelope" marker mangled fails the build. Not a silent no-op.
(7 more evidence lines)
Result: zero High, zero Medium. The Definition of done is genuinely and
verifiably true. Recorded a Converged line with the full commit hash
under ## Converged in BACKLOG.md so future relaunches on an unchanged tree
ratchet in O(1) instead of re-auditing.
That convergence is re-earned, not archived: every fresh run of Jeffy on this repo has to reach it again with fresh evidence. When Jeffy converges on your project, the checkpoint is recorded in your git log and under ## Converged in the loop's backlog, so relaunches on an unchanged tree re-verify instead of re-auditing. Run /jeffy on your own project and read the journal it leaves behind.
Proven on strangers' codebases
Self-runs are easy mode. So Jeffy was pointed at four real, famous, unaffiliated projects - each in a local clone, nothing pushed upstream, every run converging under the same rules: evidence before filing, one verified task per iteration, checkpoints, and a journal that records the operator's mistakes alongside the fixes.
kennethreitz/records (~7.2k stars). At upstream HEAD, pytest says 31 passed. At the same HEAD, INSERT s silently lose data, transaction() swallows every exception, and every query leaks a pooled connection. Jeffy reproduced four High-severity bugs hiding behind a green test suite , closed three with one structural fix at the boundary they share, restored a fix upstream had reverted the same day it was made, and left a regression suite proven to fail on the old code - repro.py shows the bugs on upstream HEAD, fixes.patch makes it show them fixed. All four were disclosed upstream with repros and a PR offer in kennethreitz/records#236 ; merging anything is the maintainers' call.
janl/mustache.js (~16.7k stars). At upstream HEAD on current Node, the test suite cannot start - the abandoned esm shim crashes before a single assertion - and bin/mustache crashes outright. Jeffy revived the gate with one structural fix across all three loading sites, fixed a second reproduced correctness bug in the CLI with a regression test, deleted the dead browser-test stack, and modernized the toolchain, taking npm audit from 107 vulnerabilities (24 critical) to 2 lows with the suite at 297 passing, official Mustache spec compliance included. The closing audit then filed a Medium against the run's own earlier work - docs still pointing at the deleted stack - and fixed it before declaring convergence. All four findings were disclosed upstream with repros and a PR offer in janl/mustache.js#848 .
sivel/speedtest-cli (~13k stars). The contrasting case: dormant since 2021 but fundamentally sound, where the honest outcome is small findings and restraint. Jeffy fixed the project's own lint gate - red on unchanged code from eight Python 2 builtin false positives - with

[truncated]
