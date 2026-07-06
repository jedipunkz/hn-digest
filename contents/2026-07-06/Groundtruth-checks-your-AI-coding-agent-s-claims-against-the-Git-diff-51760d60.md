---
source: "https://github.com/akahkhanna/groundtruth"
hn_url: "https://news.ycombinator.com/item?id=48809364"
title: "Groundtruth – checks your AI coding agent's claims against the Git diff"
article_title: "GitHub - akahkhanna/groundtruth: Groundtruth — a Claude Code plugin that audits whether an agent actually did what it was asked; catches the false 'Done' on Stop (missed subtasks, stubs, false 'tests pass', overridden rules). · GitHub"
author: "erapin_game"
captured_at: "2026-07-06T19:42:13Z"
capture_tool: "hn-digest"
hn_id: 48809364
score: 1
comments: 0
posted_at: "2026-07-06T19:29:31Z"
tags:
  - hacker-news
  - translated
---

# Groundtruth – checks your AI coding agent's claims against the Git diff

- HN: [48809364](https://news.ycombinator.com/item?id=48809364)
- Source: [github.com](https://github.com/akahkhanna/groundtruth)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T19:29:31Z

## Translation

タイトル: Groundtruth – AI コーディング エージェントのクレームを Git diff と照合してチェックします
記事のタイトル: GitHub - akahkhanna/groundtruth: Groundtruth — エージェントが頼まれたことを実際に実行したかどうかを監査するクロード コード プラグイン。停止時に偽の「完了」を捕捉します（見逃したサブタスク、スタブ、偽の「テスト合格」、オーバーライドされたルール）。 · GitHub
説明: Groundtruth — エージェントが要求されたことを実際に実行したかどうかを監査するクロード コード プラグイン。停止時に偽の「完了」を捕捉します（見逃したサブタスク、スタブ、偽の「テスト合格」、オーバーライドされたルール）。 - アカカンナ/グラウンドトゥルース

記事本文:
GitHub - akahkhanna/groundtruth: Groundtruth — エージェントが要求されたことを実際に実行したかどうかを監査するクロード コード プラグイン。停止時に偽の「完了」を捕捉します（見逃したサブタスク、スタブ、偽の「テスト合格」、オーバーライドされたルール）。 · GitHub
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
別名

ハカンナ
/
グラウンドトゥルース
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
32 コミット 32 コミット .claude-plugin .claude-plugin .github .github アセット アセット ベンチマーク ベンチマーク コマンド コマンド フック フック .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md FIXES.md FIXES.md GETTING-STARTED.md GETTING-STARTED.md ライセンスライセンス README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
他のツールはすべて、AI エージェントに適切な作業を実行させようとします。
Groundtruth はそうではないと想定し、「完了」が差分と一致しない瞬間を捕捉します。
採用するワークフローはありません。評決に対して反論するLLM審査員はいない。
エージェントの主張だけを、 git diff に対して決定的にチェックします。
誤った「完了」をキャッチする Claude Code プラグイン。エージェントは成功を報告します。 Groundtruth は、同じターン (リクエスト、エージェントの主張、実際の git diff 、プロジェクトのルール) を外部から読み取り、ターンが終了する前に 1 画面の判定をレンダリングします。
作業方法は変わりませんし、別のモデルにコードのグレーディングを依頼することもありません。これは決定論的なローカル フックです。モデル呼び出し、ネットワーク、API キーはありません。作品を読み取るものは何もないので、評決からは何も語られません。
「これはパイプラインやマルチエージェントレビューツールのようなものではないでしょうか?」いいえ。これらはエージェントにプロセスに従い、通話から推論できる LLM レビュー担当者に頼るようにします。 Groundtruth はどちらにも触れず、「完了」を主張として扱い、差分を証拠として扱います。 1 つは採用するプロセスです。これはあなたがやるべき真実の確認です。
エージェントは終了して次のように報告します。
「完了 — 指数を使用した再試行を追加しました」

ial backoff、テストに合格し、 src/upload.test.js が作成されました。」
自信に満ちた 4 つの文章。 Groundtruth は、ターンの終了が許可される前にこれをレンダリングします。
グラウンドトゥルース · Tier-1
ASK バックオフ付きの再試行を src/upload.js の S3 クライアントに追加し、src/upload.test.js にテストを追加します。
🔴 正直 — 主張は実際の行為と一致しません:
🔴 誤ったテスト/ビルドの主張 — 「テストは合格」ですが、このセッションではテスト コマンドが実行されませんでした
🟡 追加されたコードのスタブ/プレースホルダー: // TODO: 実際のバックオフ — 現時点では 1 回の試行
🟡 サイレント no-op — src/upload.test.js が主張されていますが、差分には存在しません
🔴 ルール — 差分でセキュリティ ルールが破られました:
🔴 ハードコードされたシークレット — 追加されたコード内の AWS アクセス キー
🟡 タスク — 1 件の保留中 (一度表面化され、コードに到達した場合にのみ終了します)
評決 🔴 問題 — ブロックされました (停止するには GROUNDTRUTH_BLOCK=1)
⚪ 決定論的評決 (LLM なし)。あなたの「完了」は主張です。差分が証拠です。
真実ではなかった 4 つの事柄 — すべての発見事項についてファイル/行の証拠があり、決定論的に捕らえられました。
AI エージェントは完全性の幻覚に悩まされています。プレースホルダを離れ、テスト ランナーを実行しなかったときに、機能が出荷され、テストされたと自信を持って告げます。失敗は次の 3 つの原因に分類されます。
3 つ目は、他に何もキャッチできないものです。忘れるのではなく、エージェントが認識できるルールを合理化します。監査人は決して作業を行っていないため、ルールを見逃した枠組み (「これはほんのちょっとした追加です」) が引き継がれることはありません。
/プラグイン マーケットプレイス add akahkhanna/groundtruth
/プラグインのインストール groundtruth@groundtruth
Claude Code を再起動してフックを登録します。それだけです。エージェントのターンごとに警告のみの判定カード (正直さ、完全性、セキュリティ) が与えられ、それ以上の設定は必要ありません。
インストールせずに試してみませんか?
git clone https://github.com/akahkhanna/groundtruth && クロード --p

lugin-dir ./groundtruth
必要なもの: クロード コード、ノード ≥ 18、および git リポジトリ (reality = git diff HEAD )。
そこから、プロジェクト ルールを準備し、オプションの 3 つの段階 (監査 → ルールの準備 → ブロック) でブロックをオンにします。完全なウォークスルー、ステータス バッジ、更新手順は、「スタート ガイド」に記載されています。
1 つの決定的な Stop フック — LLM なし、ネットワークなし、常に実行、~無料。 Stop ペイロードからクレームを読み取り、トランスクリプトから意図と Bash の証拠を読み取り、 git diff HEAD から現実を読み取ります。
正直 — 偽のテスト/ビルド要求 (実行なしで「テストは合格」) · スタブ/プレースホルダー (TODO / FIXME / NotImplemented /Rust todo!() /Go Panic() …) · サイレント no-op (diff にないファイルを要求) · ファントム ref (ターゲットが存在しない新しいインポート) · ドロップされたシンボル (削除された関数がまだ呼び出されている — 「リファクタリング、すべて」の下のダングリング参照「preserved」クレーム) · 特殊なケース (テスト/CI/監査人で分岐するコード)。
完全性 — 差分には決して含まれない、ask 内の名前付き成果物。意図的に粗雑です: 質問に名前が何もない場合、およびターンが要求ではなく観察である場合 (「304 は問題ありません。修正の必要はありません」) は棄権するため、余談が幻のオープン ループに組み込まれることはありません。
ルール (差別化要因) — 独自のドキュメントから決定論的な述語にコンパイルされた、独自の固定ルール。ドキュメントには文字通り、 `Y` ではなく `X` を使用するか、 `Z` を使用しないと記載されているため、違反は正規表現の一致であり、判断の判断ではありません。自動的に検出され、木に接地され、提案されますが、自動装備されることはありません。
セキュリティ — ハードコードされたシークレット、RLS オフ / 読み取り不可能な新しいテーブル (Postgres/Supabase)、コミットされた .env 。
セマンティック レイヤー - より豊富な質問と配信のマッチング、仕様の置換、回帰検出、エージェントが過去の合理化をいつ行ったかの判断

ルール - LLM が必要ですが、これはロードマップであり、出荷されません。ターンごとのエンジンは完全に決定的でオフラインのままです。
3 つの執行横木、1 つのエンジン
各段は、最後の段ではできないものをキャッチします。
停止 (ターンごとのカード、警告) → コミット前 (任意の作成者、チャットから貼り付けられたコードを含む) フックソーの停止 → git commit で警告) → CI ( --diff-rangeorigin/main..HEAD 、バイパス防止 → PR をブロック)。
コピー＆ペーストのフック インストーラーと準備完了の GitHub アクションにより、外側の 2 つはそれぞれ 1 つのコマンドになります。
コマンド
何をするのか
/グラウンドトゥルース
このセッションの最新の評決カードを表示します。
/groundtruth-監査
リポジトリ全体をスキャンして、エージェントの負債 (スタブ、TODO、ファントム インポート) を探します。評決ではなく在庫です。
/groundtruth-rules
ドキュメントからコンパイルされたルールを確認および承認します (権限ゲート)。すべてのクリーンルールを承認します。 unarm <id> は誤って発砲した人を沈黙させます。
/groundtruth-rules-ai
オプトイン、デフォルトではオフ。モデルパスは、散文で書かれたドキュメントを読み取り、リテラル抽出者が見逃したルールを提案します。これは、同じグラウンディングと承認のゲートを経由します。モデルがグラウンドトゥルースに触れることがある唯一の場所。
/groundtruth-block オン｜オフ
ブロックを選択します (デフォルトは警告です)。反転する前に、発射回数に裏付けられた項目別の確認を表示します。
/groundtruth-setup
ワンショット インストーラー: 構成内容を検出し、同意に基づいてクリーン ルールを準備し、残り (バッジ、環境) を渡します。
CLI (インストールなし): ノードフック/groundtruth.mjs --audit はリポジトリ監査を実行します。 · --latest は最後の判定を出力します。 · --install-pre-commit は段階的差分スキャンを接続します。 · --diff-range Origin/main..HEAD は CI ゲートです (ブロック重大度の結果でゼロ以外で終了します)。
デフォルト: 警告。評決は記録される。ターンが中断されることはありません。まずは信頼を築きましょう。
オプトイン: ブロックします。ブロック重大度の検出により停止が拒否され、修正ペイロードが返されて、修正が再チェックされます。

次の停止 — ループの試行回数は 2 回に制限され、その後人間にエスカレートします。キャッチを満たすためにテスト / このチェッカー / 台帳を編集すると、ゲームとしてフラグが立てられます。ブロックは解放されずに保持されます。
誤検知は致命的です。実際のセッションで精度が証明されるまで WARN で実行し、ブロックを反転します。すべての評決にはファイル/ラインの証拠が含まれるため、誤った通話は監査可能であり、不可解なものではありません。
検証者は、自身のミスについて正直である場合にのみ信頼する価値があります。
精度は直感ではなく実際のデータに基づいて再構築されました。 Groundtruth が自身のセッション 15 回にわたって発行したすべての検出結果 (23 個の検出結果、74% の誤検知) を読み取り、それらをラベル付けされたコーパスに凍結してから削除しました。エンジン内の自己一致誤検出は 0 、ファントムインポート FP は 3 → 0 、セルフチェックは 242 → 440 、レッドチームは 14/14 になりました。その後、独立した敵対的パスが修正を破ろうとし、独自のテスト ケースを作成しました。
すべての修正は、症状→根本原因→修正→回帰テストとして FIXES.md にカタログ化されます。これには、レビューで見つかった 2 つの重大なホールと、解決されていない決定論的 NL 制限も含まれます。
まだ保留中、非表示という名前: 1 週間の実際のセッションにわたるライブ前後の偽陽性率。 Groundtruth には追加専用の履歴ログと GT-harvest リーダーが同梱されているため、リポジトリで測定できます。その番号は次に出荷されますが、ミスも含めて同じ方法で発行されます。
スルーライン: 棄権するか、ドメイン外での制限付き警告に格下げします。チェックが正しいことが証明されているため (決してクリーングリーンではありません)、監査対象エージェントが自らの判断を形成することは決してありません。
構文ではなくメカニズム — ほとんどのチェックは言語全般に適用されます。
一般: false - 「テストに合格」(Go/Rust/Ruby/Java/.NET/Python/JS ランナー)、スタブ/プレースホルダー、サイレント no-op、完全性、ディレクティブ オーバーライド、シークレット、環境公開。
JS/TS + Ruby

nly: ファントム相対インポート解決 — ファイルの存在によって解決されます。パッケージ修飾言語 (Python/Go/Rust/Java/C#) は、フォールスフラグではなく棄権します。
Postgres/Supabase のみ: RLS チェック - 追加された .sql 行のみで実行されます。
Groundtruth は、監査対象のエージェントと同じ信頼ドメイン、つまり同じファイルシステム、同じ環境で実行されます。したがって、それが作成するディスク上のアーティファクトは実際のセキュリティ境界ではなく、敵対的なエージェントはそれらのいずれかを書き換えることができます。シェルにアクセスできるエージェントに対して、Groundtruth は改ざん防止ではなく、改ざん明示的です。
これを信頼できるものにするのは、エージェントが作成できない 2 つの入力、つまりトランスクリプト (ハーネス自身のツール呼び出しの記録) と git で計算された作業ツリーの差分に基づいていることです。セッション中の改ざん証拠は、GROUNDTRUTH_KEY が設定され、エージェントのシェル環境の外部で保持されているため信頼性があります。証拠→防止をアップグレードするには、別の信頼ドメイン (マージ前の CI ゲート) でフックを実行します。
完全な敵対的分析 (3 つのレッドチーム パス、ロンダリングされたヘルパーの回避、修正とその制限) は SECURITY.md と FIXES.md にあります。
ノードフック/groundtruth.test.mjs # 440 アサートベースのユニットチェック、deps なし
ノードフック/redteam.mjs # ライブ敵対的ハーネス (10 シナリオ、14 チェック)、サンドボックス化
レッドチームのハーネスは実証済みの対応物です。使い捨てリポジトリを作成し、本物のフックにトランスクリプトを渡し、賢いエージェントが積極的にレールを無効化します。つまり、支配を武装解除します。

[切り捨てられた]

## Original Extract

Groundtruth — a Claude Code plugin that audits whether an agent actually did what it was asked; catches the false 'Done' on Stop (missed subtasks, stubs, false 'tests pass', overridden rules). - akahkhanna/groundtruth

GitHub - akahkhanna/groundtruth: Groundtruth — a Claude Code plugin that audits whether an agent actually did what it was asked; catches the false 'Done' on Stop (missed subtasks, stubs, false 'tests pass', overridden rules). · GitHub
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
akahkhanna
/
groundtruth
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
32 Commits 32 Commits .claude-plugin .claude-plugin .github .github assets assets benchmarks benchmarks commands commands hooks hooks .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md FIXES.md FIXES.md GETTING-STARTED.md GETTING-STARTED.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md package.json package.json View all files Repository files navigation
Every other tool tries to make your AI agent do the work right.
Groundtruth assumes it won't — and catches the moment “done” doesn't match the diff.
No workflow to adopt. No LLM reviewer to argue out of a verdict.
Just your agent's claim, checked against the git diff , deterministically.
A Claude Code plugin that catches the false “Done.” Your agent reports success; Groundtruth reads the same turn from the outside — the request, what the agent claimed, the real git diff , and your project's rules — and renders a one-screen verdict before the turn can end.
It doesn't change how you work and it doesn't ask another model to grade the code. It's a deterministic local hook: no model calls, no network, no API key. Nothing reads the work, so nothing can be talked out of the verdict.
“Isn't this like the pipeline / multi-agent-review tools?” No. Those make your agent follow a process, and lean on LLM reviewers that can be reasoned out of a call. Groundtruth touches neither — it treats “done” as a claim and the diff as the evidence. One is a process you adopt; this is a truth-check you bolt on.
Your agent finishes and reports:
“Done — added retry with exponential backoff, tests pass, and created src/upload.test.js .”
Four confident sentences. Groundtruth renders this before the turn is allowed to end:
GROUNDTRUTH · Tier-1
ASK Add retry with backoff to the S3 client in src/upload.js, and a test in src/upload.test.js.
🔴 Honesty — claims don't match what it did:
🔴 false test/build claim — "tests pass", but no test command ran this session
🟡 stub/placeholder in added code: // TODO: real backoff — single attempt for now
🟡 silent no-op — claimed src/upload.test.js, but it is absent from the diff
🔴 Rules — a security rule was broken in the diff:
🔴 hardcoded secret — AWS access key in added code
🟡 Tasks — 1 pending (surfaced once; closes only when it lands in code)
VERDICT 🔴 ISSUES — blocked (GROUNDTRUTH_BLOCK=1 to halt)
⚪ Deterministic verdict (no LLM). Your "done" is a claim; the diff is the evidence.
Four things that weren't true — caught deterministically , with file/line evidence on every finding.
AI agents suffer from the hallucination of completeness : they'll confidently tell you a feature is shipped and tested when they left placeholders and never ran the test runner. Failures sort into three causes:
The third is the one nothing else catches — not forgetting, but rationalising past a rule the agent could see. The auditor never did the work, so it never inherits the framing (“this is just a small addition”) that let the rule slip.
/plugin marketplace add akahkhanna/groundtruth
/plugin install groundtruth@groundtruth
Restart Claude Code so the hooks register. That's it — every agent turn now gets a warn-only verdict card (honesty, completeness, security), no further config.
Prefer to try without installing?
git clone https://github.com/akahkhanna/groundtruth && claude --plugin-dir ./groundtruth
Requires: Claude Code, node ≥ 18, and a git repo (reality = git diff HEAD ).
From there you arm your project rules and turn on blocking in three optional stages — audit → arm rules → block . The full walkthrough, the status badge, and the update steps live in the Getting Started guide .
One deterministic Stop hook — no LLM, no network, always runs, ~free. It reads the claim from the Stop payload, intent + Bash evidence from the transcript, and reality from git diff HEAD :
Honesty — false test/build claim (“tests pass” with no run) · stub/placeholder ( TODO / FIXME / NotImplemented /Rust todo!() /Go panic() …) · silent no-op (claimed a file that's absent from the diff) · phantom ref (new import whose target doesn't exist) · dropped symbol (a removed function still called — a dangling reference under a “refactor, everything preserved” claim) · special-casing (code that branches on test/CI/the auditor).
Completeness — a named deliverable in the ask that never lands in the diff. Deliberately crude: it abstains when the ask names nothing, and when the turn is an observation rather than a request (“the 304 is fine, no fix needed”) so an aside is never minted into a phantom open loop.
Rules (the differentiator) — your standing rules, compiled from your own docs into deterministic predicates . The doc literally says use `X` not `Y` or never `Z` , so a violation is a regex match, not a judgment call. Auto-discovered, grounded against your tree, and proposed — never auto-armed.
Security — hardcoded secrets, an RLS-off / anon-readable new table (Postgres/Supabase), a committed .env .
A semantic layer — richer ask↔delivery matching, spec-substitution, regression detection, judging when an agent rationalised past a rule — needs an LLM and is roadmap, not shipped . The per-turn engine stays fully deterministic and offline.
Three enforcement rungs, one engine
Each rung catches what the last can't:
Stop (per-turn card, warn) → pre-commit (any author, incl. code pasted from a chat no Stop hook saw → warn at git commit ) → CI ( --diff-range origin/main..HEAD , bypass-proof → block the PR).
A copy-paste hook installer and a ready GitHub Action make the outer two one command each.
Command
What it does
/groundtruth
Show the latest verdict card for this session.
/groundtruth-audit
Scan the whole repo for agent debt — stubs, TODOs, phantom imports. Inventory, not a verdict.
/groundtruth-rules
Review + approve rules compiled from your docs — the permission gate . approve-all arms every clean rule; unarm <id> silences one firing wrongly.
/groundtruth-rules-ai
Opt-in, off by default. A model pass reads your docs in prose and proposes rules the literal extractor missed — routed through the same grounding + approval gate. The one place a model ever touches Groundtruth.
/groundtruth-block on｜off
Opt into blocking (default is warn). Shows an itemized, fire-count-backed confirmation before it flips.
/groundtruth-setup
One-shot installer: detects what's configured, arms clean rules on consent, hands you the rest (badge, env).
CLI (no install): node hooks/groundtruth.mjs --audit runs the repo audit · --latest prints the last verdict · --install-pre-commit wires a staged-diff scan · --diff-range origin/main..HEAD is the CI gate (exits non-zero on a block-severity finding).
Default: WARN. The verdict is recorded; the turn is never disrupted. Build trust first.
Opt-in: BLOCK. A block-severity finding refuses the stop and hands back a corrective payload, then re-checks the fix on the next stop — a loop capped at 2 attempts , then escalates to a human. Editing the tests / this checker / the ledger to satisfy a catch is flagged as gaming: the block holds rather than releasing.
False positives are fatal. Run in WARN until precision is proven on your real sessions, then flip block. Every verdict carries file/line evidence, so a wrong call is auditable, not mysterious.
A verifier is only worth trusting if it's honest about its own misses.
Precision was rebuilt against real data, not intuition. We read every finding Groundtruth emitted across 15 of its own sessions — 23 findings, 74% false positives — froze them into a labeled corpus , then killed them: self-match false positives in the engine went to 0 , phantom-import FPs 3 → 0 , self-checks 242 → 440 , red-team 14/14 . Independent adversarial passes then tried to break the fixes and wrote their own test cases.
Every fix is catalogued — symptom → root cause → fix → regression test — in FIXES.md , including the two critical holes review found and the deterministic-NL limits it does not close.
Still pending, named not hidden: a live before/after false-positive rate across a week of real sessions. Groundtruth ships an append-only history log + a gt-harvest reader so you can measure it on your repo. That number ships next — published the same way, with its misses.
The through-line: abstain or degrade to a bounded warn outside the domain a check is provably correct in — never a clean green — and never let the audited agent shape its own verdict.
Mechanics, not syntax — most checks are language-general.
General: false-“tests pass” (Go/Rust/Ruby/Java/.NET/Python/JS runners), stub/placeholder, silent no-op, completeness, directive-override, secrets, env exposure.
JS/TS + Ruby only: phantom relative-import resolution — it resolves by file existence. Package-qualified languages (Python/Go/Rust/Java/C#) abstain rather than false-flag.
Postgres/Supabase only: the RLS checks — they fire solely on added .sql lines.
Groundtruth runs in the same trust domain as the agent it audits : same filesystem, same environment. So no on-disk artifact it creates is a real security boundary — an adversarial agent can rewrite any of them. Against an agent with shell access, Groundtruth is tamper-evident, not tamper-proof .
What makes it trustworthy is anchoring on the two inputs the agent can't author: the transcript (the harness's own record of tool calls) and the git-computed diff of the working tree. In-session tamper-evidence is reliable with GROUNDTRUTH_KEY set and held outside the agent's shell env; to upgrade evidence → prevention , run the hook in a separate trust domain (a pre-merge CI gate).
The full adversarial analysis — three red-team passes, the laundered-helper evasion, the fixes and their limits — is in SECURITY.md and FIXES.md .
node hooks/groundtruth.test.mjs # 440 assert-based unit checks, no deps
node hooks/redteam.mjs # live adversarial harness (10 scenarios, 14 checks), sandboxed
The red-team harness is the proven counterpart: it spins up a throwaway repo, hands the real hook a transcript where a smart agent actively neuters the rails — disarms rul

[truncated]
