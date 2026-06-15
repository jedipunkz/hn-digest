---
source: "https://github.com/softcane/ccverdict"
hn_url: "https://news.ycombinator.com/item?id=48545671"
title: "Show HN: Behavioral gauge for Claude Code sessions"
article_title: "GitHub - softcane/ccverdict: A stop-loss statusline for Claude Code. · GitHub"
author: "pradeep1177"
captured_at: "2026-06-15T19:24:14Z"
capture_tool: "hn-digest"
hn_id: 48545671
score: 1
comments: 0
posted_at: "2026-06-15T19:09:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Behavioral gauge for Claude Code sessions

- HN: [48545671](https://news.ycombinator.com/item?id=48545671)
- Source: [github.com](https://github.com/softcane/ccverdict)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T19:09:16Z

## Translation

タイトル: HN を表示: クロード コード セッションの行動ゲージ
記事のタイトル: GitHub - Softcane/ccverdict: Claude Code のストップロス ステータスライン。 · GitHub
説明: クロード コードのストップロス ステータスライン。 GitHub でアカウントを作成して、softcane/ccverdict の開発に貢献してください。

記事本文:
GitHub - Softcane/ccverdict: Claude Code のストップロス ステータスライン。 · GitHub
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
ソフトケーン
/
ccverdict
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイル C に移動

ode 「その他のアクション」メニューを開く フォルダーとファイル
67 コミット 67 コミット .github/ workflows .github/ workflows アセット アセット src src test test .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コード セッションのローカル動作ゲージ。
クロード コードは漂流している間、忙しそうに見えることがあります。失敗した同じチェックを再試行したり、検証せずにファイルを編集したり、すでに確認したファイルを再読み込みしたり、コンテキスト ウィンドウを押し上げすぎたりする可能性があります。
多くのステータスラインには、モデル、ブランチ、コスト、時間などの事実が表示されます。 ccverdict は動作判定を追加します。
進行中 · 漂流中 · 介入が必要 · 信号なし
派生したローカル メタデータを監視し、次の 1 つの質問に答えます。
エージェントは現在何をしていますか? その動作は正常に見えますか?
フックをインストールする前に試してみてください。
npx --yes ccverdict Audit --project 。
クロードコードのステータスラインに判決を反映させたい場合にインストールしてください。
ライブ ゲージには 1 つの短い行が出力されます。
<ドット> · <動詞> · <証拠> · <ファイル> · <ctx> · <コスト>
● 編集 · 1 ファイル、1 つは未チェック (auth.ts…) · ctx 42% · $0.18
◐ 編集 · 3 ファイル、2 つは未チェック (auth.ts…) · ctx 42%
■ テストの再試行 · 3 回失敗、実行間で修正なし
■ 探索・config.ts の再読み込み 3 回
○ 信号がない・トランスクリプトが読めない
● アイドル状態 · まだアクティビティがありません
ccverdict はその行に命令を出力しません。ドット、現在のアクティビティ、およびドットの背後にある証拠が表示されます。
ファイル セグメントには、カウントと 1 つのベース名のヒント、つまり最新の未チェックのファイルが表示されます。すべてのファイルがリストされるわけではありません。
デモでは、ステータスラインの状態の例を循環します: 健全な編集、チェックなし

ドリフト、繰り返しの検証失敗、繰り返しのファイル読み取り、高度なコンテキスト、および読み取り不能な証拠。
ドットには色と形状の両方があるため、 NO_COLOR でも機能します。
アクティビティ動詞は、以下のいくつかが当てはまる場合にこの優先順位を使用します。
再試行 > テスト > 編集 > 探索 > アイドル
コンテキストとコストは事実です。彼らは自分自身でドットを変更しません。例外は 92% 以上のコンテキストで、セッションが満杯に近いため赤色に変わります。
80% ～ 91% では、ccverdict は ctx セグメントのみを強調表示します。色を無効にすると、マーカーが使用されます。
●探検・ctx85%！
何を捕まえるのか
修正を行わないと、同じ検証チェックが再び失敗します。
繰り返されるテスト、lint、タイプチェック、またはビルドの失敗。
合格チェックを受けずに溜まっていく編集。
回復シグナルのない読み取り負荷の高いセッション。
同じ変更されていないファイルが再度読み取られます。
多くの入力トークンを追加する大規模なツールの結果。
プロンプトキャッシュの再利用は、動作していた後に削除されました。
セッションがいっぱいになる前のコンテキストプレッシャー。
検証がまだ失敗している間にコストや時間が増加します。
新たな目標チェックが必要な圧縮境界。
許可プロンプトは別のものです。ツール権限プロンプトを拒否した場合、ccverdict はそれをコマンドの失敗または再試行ループとしてカウントしません。
ステータスラインをサポートするクロードコード
npx --yes ccverdict install --scope local
プロジェクト内の Claude Code を再起動します。ステータスラインが下部に表示されます。
デフォルトのインストールではコーチ モードを使用し、可能な場合は小さなローカル ベースラインを構築します。 --replace を渡さない限り、既存のクロード コードの statusLine が保持されます。
npx --yes ccverdict install --scope local --replace
ゲージとローカル履歴のみが必要な場合は、observ-only を使用します。
npx --yes ccverdict install --scope local --observe-only
コーチからのフィードバックと、信頼性の高い検証の再試行の繰り返しを厳密に拒否したい場合は、ガードを使用します。

s:
npx --yes ccverdict install --scope local --guard
ベースラインの作成とレッスンの記憶をスキップします。
npx --yes ccverdict install --scope local --no-learn
アンインストール:
npx --yes ccverdict uninstall --scope local
サポートされているスコープは、ローカル、プロジェクト、およびユーザーです。プロジェクト レベルまたはユーザー レベルのクロード設定を編集する必要がない限り、1 つのリポジトリにはローカルを使用します。
npm install -g ccverdict
ccverdict install --scope local
フィードバックモード
Observ-Only はステータスラインとローカル派生イベント データを保持します。クロードにはフィードバックは送信されません。
コーチがデフォルトです。 ccverdict が未チェックの編集、繰り返しの検証失敗、終了時の未解決の検証リスクなどの危険なパターンを検出した場合、クロードに向けた短いメモを送信できます。
ガードにはコーチのフィードバックが含まれます。編集やチェックに合格せずに失敗を繰り返した後に同じテスト、lint、タイプチェック、またはビルド カテゴリを再実行するなど、信頼性の高い繰り返しの Bash 検証の再試行を拒否できます。通常の読み取り、編集、または任意のコマンドはブロックされません。
ccverdict は、安全なフィードバック結果も記録します。コーチのフィードバックが検証を要求し、後でクロードが合格テストを実行すると、監査でループが解決されたことが示されます。
最近の ccverdict ループ:
1. コーチのフィードバック: 検証が必要な編集。
2. クロードはテストを実行しました。
3. テストに合格しました。
4. 結果: 解決されました。
インストールする前に試してください
最近のローカル クロード コード履歴に対して監査を実行します。
npx --yes ccverdict Audit --project 。
npx --yes ccverdict Audit --all-projects --recent 200
Audit はステータスラインやフックをインストールしません。 --all-projects は、ローカルのクロード プロジェクト全体の最新のトランスクリプトを検査する場合にのみ使用します。
Audit はゲージのオフライン ビューです。 3 つのセクションが印刷されます。
現在のプロジェクトのみに関する最新の ccverdict 決定を表示します: ドット、ライト、経過時間、すべての結果、元帳の編集、およびコーチまたはガードのフィードバック出力

オメス。プロジェクトに ccverdict 履歴がない場合、監査はそのように判断します。別のリポジトリのセッションは表示されません。
レポートには、ブラインド再試行、未チェック編集ストリーク、繰り返し読み取り、変更後のリカバリ、圧縮リスク、セッション終了リスクなどのローカル パターンが集約されます。
CLAUDE.md を読み取り、命令行を最近検出されたカテゴリと比較します。空ではないサブセクションのみを出力します。
削除候補 : 最近見つかったカテゴリに一致しなかった行。
どうやらフォローされているようです: カテゴリが既に高い準拠性を持っている行。
ギャップ : どの指示行も対応していないカテゴリが繰り返し見つかります。
レポートでは「一致した」および「一致しなかった」という表現が使用されています。指示が結果を引き起こしたとは主張しません。
Plain Audit は、ccverdict 自身のストアの外には何も書き込みません。
Audit --apply は統合された diff を出力し、ccverdict のマークされた CLAUDE.md ブロック内にのみ書き込みます。まずバックアップします。実行ごとに追加される汎用行は最大 2 行で、生のプロンプト、コマンド、パス、ツール出力は決して書き込まれません。
ccverdict は、ユーザーが作成した CLAUDE.md 行を削除または編集しません。削除の提案は、差分にコメントされた提案として表示されます。 Audit --cleanup は、バックアップ後にマークされたブロックを削除します。
プロジェクト固有のレッスンは ./CLAUDE.md にルーティングされます。 --global または --all-projects を使用する場合、プロジェクト間の習慣は ~/.claude/CLAUDE.md にルーティングされます。 Audit は AGENTS.md を編集しません。
ccverdict 監査 --project 。
ccverdict 監査 --all-projects --最近の 200
ccverdict 監査 --apply --project 。
ccverdict 監査 --cleanup --project 。
医師の評決
ccverdict 医師 --ベースライン
ccverdict uninstall --scope local
ccverdict アンインストール --purge
Audit --json は、3 つの監査セクションすべてについて機械可読な出力を返します。
ドクターは、ノード、クロード設定、オプションのフック、トランスクリプトへのアクセス、価格設定キャッシュ、ストアのバージョン、およびベースライン診断をチェックします。
博士

tor --baseline は、安全な集計ベースラインの事実を示します。
アンインストールは、ccverdict が所有するステータスラインとフックを削除します。有効なバックアップが存在する場合、以前のクロード コードのステータスラインが復元されます。 uninstall --purge は、学習されたベースライン、レッスン メモリ、派生イベント ストアも削除します。
ccverdict は、クロード コードの実行チェックを監視します。テスト、lint、typecheck、またはビルド コマンドを単独で実行することはありません。
一般的な Bash 検証コマンドを認識し、それらを安全なカテゴリにグループ化します。
ccverdict は、再試行ループの検出、回復ベースライン、コーチ フィードバック、およびガード モードの再試行拒否にこれらのカテゴリを使用します。
プロジェクトでカスタム検証コマンドを使用する場合は、 .ccverdict.json を追加します。
{
"検証コマンド" : {
"テスト" : [ "テストを作成する" ],
"lint" : [ "lint を作成する" ],
"ビルド" : [ "ビルドを作成する " ]
}
}
このファイルはオプションです。 ccverdict はこれを分類に使用しますが、それらの生のコマンドをイベント履歴にコピーしません。
● 探索 · ctx 24% · $0.12
● 編集中 · 2 ファイル、1 つは未チェック (router.ts…) · ctx 31%
◐ 編集 · 4 ファイル、3 つは未チェック (router.ts…) · ctx 44%
◐ テスト · テストは 2 回失敗しました · ctx 52%
◐ 探索 · キャッシュの再利用が 68% から 29% に低下
◐ アイドル状態、圧縮境界が開いている状態
■ テストの再試行 · 3 回失敗、実行間で修正なし
■ 探索・config.ts の再読み込み 3 回
■探索中・ctx 93%、ほぼフル
○ 信号なし・トランスクリプトセッションの不一致
幅の狭い端末では、行はコンパクトなカウントに縮小されます。
◐ 編集 · 3✎? · 44%
◐ 3✎? 44%
ドットは表示されたままになります。
ccverdict はローカルファーストです。クラウド バックエンド、SaaS ダッシュボード、トランスクリプト アップロード、プロキシ、ゲートウェイ、メッセージ ルーターはありません。
現在のゲージ レコードは派生メタデータのみを使用します: ライト、アクティビティ、検索カテゴリ、信頼ラベル、カウント、レート、パーセンタイル、解決済みまたは無視された、安全などのフィードバック結果

テスト、トークンとコストのフィールド、コンテキスト フィールド、タイムスタンプ、ハッシュされたファイル ID、ハッシュされたセッション キー、ハッシュされたプロジェクト キーなどの検証カテゴリ。
イベント ストアには、プロンプト、アシスタント テキスト、ツール出力、シェル出力、コマンド引数、ファイル コンテンツ、トランスクリプト パス、ワークスペース パス、API キー、生のクロード セッション ID、または生の MCP 名は保存されません。
繰り返し読み取りの警告の場合、 ccverdict は auth.ts などのベース名のヒントを表示する場合があります。完全なローカル パスは保存または出力されません。
CCVERDICT_HOME または別のオーバーライドを設定しない限り、ローカル ファイルはデフォルトで ~/.claude/ccverdict の下に存在します。
events.json には、派生したローカルの決定、フック イベント、およびフィードバックの結果が保存されます。
Baseline.json には、個人の集計ベースラインが保存されます。
project-baselines/<hashed-project>.json には、集計されたプロジェクト ベースラインが格納されます。
project-lessons/<hashed-project>.json には、減衰するプロジェクト レッスン メモリが保存されます。
litellm-pricing.json は、更新時に公開価格データをキャッシュします。
バックアップ/アンインストールで使用されるクロード設定のバックアップを保存します。
インストール バックアップはローカル設定のスナップショットです。これらには、クロード設定にすでに存在するステータスライン、フック コマンド、またはパスが含まれる場合があります。 ccverdict はそれらをアップロードしません。
LiteLLM は、コスト見積もりの​​ための公開価格データとしてのみ使用されます。 ccverdict は LiteLLM プロキシまたは rou を実行しません

[切り捨てられた]

## Original Extract

A stop-loss statusline for Claude Code. Contribute to softcane/ccverdict development by creating an account on GitHub.

GitHub - softcane/ccverdict: A stop-loss statusline for Claude Code. · GitHub
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
softcane
/
ccverdict
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
67 Commits 67 Commits .github/ workflows .github/ workflows assets assets src src test test .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A local behavior gauge for Claude Code sessions.
Claude Code can look busy while it drifts. It can retry the same failed check, edit files without validation, reread a file it already saw, or push the context window too high.
Many statuslines show facts like model, branch, cost, or time. ccverdict adds a behavior verdict:
progressing · drifting · needs intervention · no signal
It watches derived local metadata and answers one question:
What is the agent doing right now, and does its behavior look healthy?
Try it before installing hooks:
npx --yes ccverdict audit --project .
Install it when you want the verdict live in your Claude Code statusline.
The live gauge prints one short line:
<dot> · <verb> · <evidence> · <files> · <ctx> · <cost>
● editing · 1 file, 1 unchecked (auth.ts…) · ctx 42% · $0.18
◐ editing · 3 files, 2 unchecked (auth.ts…) · ctx 42%
■ retrying tests · 3 fails, no fix between runs
■ exploring · reread config.ts 3x
○ no signal · transcript unreadable
● idle · no activity yet
ccverdict prints no instructions on the line. It shows the dot, the current activity, and the evidence behind the dot.
The file segment shows counts plus one basename hint: the latest unchecked file. It does not list every file.
The demo cycles through example statusline states: healthy editing, unchecked-edit drift, repeated validation failure, repeated file reads, high context, and unreadable evidence.
The dot has both a color and a shape, so it still works with NO_COLOR .
Activity verbs use this priority when several apply:
retrying > testing > editing > exploring > idle
Context and cost are facts. They do not change the dot by themselves. The exception is context at 92% or higher, which turns red because the session is close to full.
At 80% through 91% , ccverdict highlights only the ctx segment. With color disabled, it uses a marker:
● exploring · ctx 85%!
What It Catches
The same validation check failing again without a fix.
Test, lint, typecheck, or build failures that repeat.
Edits that pile up without a passing check.
Read-heavy sessions with no recovery signal.
The same unchanged file being read again.
A large tool result that adds many input tokens.
Prompt-cache reuse dropping after it was working.
Context pressure before the session gets too full.
Cost or time climbing while validation still fails.
Compaction boundaries that need a fresh goal check.
Permission prompts are separate. If you deny a tool permission prompt, ccverdict does not count that as a command failure or retry loop.
Claude Code with status line support
npx --yes ccverdict install --scope local
Restart Claude Code in the project. The statusline appears at the bottom.
Default install uses coach mode and builds a small local baseline when it can. It preserves an existing Claude Code statusLine unless you pass --replace .
npx --yes ccverdict install --scope local --replace
Use observe-only when you only want the gauge and local history:
npx --yes ccverdict install --scope local --observe-only
Use guard when you want coach feedback plus strict denial of high-confidence repeated validation retries:
npx --yes ccverdict install --scope local --guard
Skip baseline creation and lesson memory:
npx --yes ccverdict install --scope local --no-learn
Uninstall:
npx --yes ccverdict uninstall --scope local
Supported scopes are local , project , and user . Use local for one repo unless you need to edit project-level or user-level Claude settings.
npm install -g ccverdict
ccverdict install --scope local
Feedback Modes
observe-only keeps the statusline and local derived event data. It sends no feedback to Claude.
coach is the default. It can send short Claude-facing notes when ccverdict sees a risky pattern, such as unchecked edits, repeated validation failure, or unresolved validation risk at finish.
guard includes coach feedback. It can deny a high-confidence repeated Bash validation retry, such as rerunning the same test, lint, typecheck, or build category after repeated failures without an edit or passing check. It does not block normal reads, edits, or arbitrary commands.
ccverdict also records safe feedback outcomes. If coach feedback asks for validation and Claude later runs a passing test, audit can show that the loop resolved:
Recent ccverdict loop:
1. Coach feedback: edits needed validation.
2. Claude ran tests.
3. Tests passed.
4. Outcome: resolved.
Try Before Installing
Run an audit against recent local Claude Code history:
npx --yes ccverdict audit --project .
npx --yes ccverdict audit --all-projects --recent 200
audit does not install a statusline or hooks. Use --all-projects only when you want to inspect the newest transcripts across local Claude projects.
audit is the offline view of the gauge. It prints three sections.
Shows the latest ccverdict decision for the current project only: dot, light, age, all findings, edit ledger, and coach or guard feedback outcomes. If the project has no ccverdict history, audit says so. It does not show another repo's session.
Reports aggregate local patterns: blind retries, unchecked-edit streaks, repeated reads, recovery after change, compaction risk, and session-end risk.
Reads CLAUDE.md and compares instruction lines against recent finding categories. It prints only non-empty subsections:
Candidates for removal : lines that matched no recent finding category.
Apparently followed : lines whose category already has high compliance.
Gaps : recurring finding categories that no instruction line addresses.
The report uses "matched" and "didn't match" language. It does not claim that an instruction caused an outcome.
Plain audit writes nothing outside ccverdict's own store.
audit --apply prints a unified diff, then writes only inside ccverdict's marked CLAUDE.md block. It backs up first. It adds at most two generic lines per run, and it never writes raw prompts, commands, paths, or tool output.
ccverdict does not delete or edit user-written CLAUDE.md lines. Removal suggestions appear as commented proposals in the diff. audit --cleanup removes the marked block after a backup.
Project-specific lessons route to ./CLAUDE.md . Cross-project habits route to ~/.claude/CLAUDE.md when you use --global or --all-projects . audit does not edit AGENTS.md .
ccverdict audit --project .
ccverdict audit --all-projects --recent 200
ccverdict audit --apply --project .
ccverdict audit --cleanup --project .
ccverdict doctor
ccverdict doctor --baseline
ccverdict uninstall --scope local
ccverdict uninstall --purge
audit --json returns machine-readable output for all three audit sections.
doctor checks Node, Claude settings, optional hooks, transcript access, pricing cache, store version, and baseline diagnostics.
doctor --baseline shows safe aggregate baseline facts.
uninstall removes ccverdict-owned statusline and hooks. When a valid backup exists, it restores the previous Claude Code statusline. uninstall --purge also deletes learned baselines, lesson memory, and the derived event store.
ccverdict observes checks Claude Code runs. It does not run tests, lint, typecheck, or build commands by itself.
It recognizes common Bash validation commands and groups them into safe categories:
ccverdict uses those categories for retry-loop detection, recovery baselines, coach feedback, and guard-mode retry denial.
If your project uses custom validation commands, add .ccverdict.json :
{
"validationCommands" : {
"tests" : [ " make test " ],
"lint" : [ " make lint " ],
"build" : [ " make build " ]
}
}
This file is optional. ccverdict uses it for classification, but it does not copy those raw commands into event history.
● exploring · ctx 24% · $0.12
● editing · 2 files, 1 unchecked (router.ts…) · ctx 31%
◐ editing · 4 files, 3 unchecked (router.ts…) · ctx 44%
◐ testing · tests failed twice · ctx 52%
◐ exploring · cache reuse dropped from 68% to 29%
◐ idle · compaction boundary open
■ retrying tests · 3 fails, no fix between runs
■ exploring · reread config.ts 3x
■ exploring · ctx 93%, nearly full
○ no signal · transcript session mismatch
On narrow terminals the line shrinks to compact counts:
◐ editing · 3✎? · 44%
◐ 3✎? 44%
The dot stays visible.
ccverdict is local-first. It has no cloud backend, SaaS dashboard, transcript upload, proxy, gateway, or message router.
Current gauge records use derived metadata only: light, activity, finding categories, confidence labels, counts, rates, percentiles, feedback outcomes such as resolved or ignored , safe validation categories such as tests , token and cost fields, context fields, timestamps, hashed file identities, hashed session keys, and hashed project keys.
The event store does not store prompts, assistant text, tool output, shell output, command arguments, file contents, transcript paths, workspace paths, API keys, raw Claude session ids, or raw MCP names.
For repeated-read warnings, ccverdict may show a basename hint such as auth.ts . It does not store or print the full local path.
Local files live under ~/.claude/ccverdict by default, unless you set CCVERDICT_HOME or another override:
events.json stores derived local decisions, hook events, and feedback outcomes.
baseline.json stores the personal aggregate baseline.
project-baselines/<hashed-project>.json stores aggregate project baselines.
project-lessons/<hashed-project>.json stores decaying project lesson memory.
litellm-pricing.json caches public pricing data when refreshed.
backups/ stores Claude settings backups used by uninstall.
Install backups are local settings snapshots. They may contain whatever statusline, hook commands, or paths already existed in those Claude settings. ccverdict does not upload them.
LiteLLM is used only as public pricing data for cost estimates. ccverdict does not run a LiteLLM proxy or rou

[truncated]
