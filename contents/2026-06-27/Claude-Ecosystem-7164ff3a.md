---
source: "https://claude.com/ecosystem"
hn_url: "https://news.ycombinator.com/item?id=48701356"
title: "Claude Ecosystem"
article_title: "Claude ecosystem | Claude by Anthropic"
author: "geoffbp"
captured_at: "2026-06-27T20:22:42Z"
capture_tool: "hn-digest"
hn_id: 48701356
score: 1
comments: 0
posted_at: "2026-06-27T20:17:37Z"
tags:
  - hacker-news
  - translated
---

# Claude Ecosystem

- HN: [48701356](https://news.ycombinator.com/item?id=48701356)
- Source: [claude.com](https://claude.com/ecosystem)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T20:17:37Z

## Translation

タイトル: クロードの生態系
記事タイトル: クロードの生態系 |クロード by Anthropic
説明: チームがすでに依存しているツールをクロードに接続します。ディレクトリとマーケットプレイスを見つけてください。

記事本文:
クロードの生態系 |クロード by Anthropic
クロード製品のご紹介 クロード
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
クロードの生態系 クロードの生態系
クロードはあなたが働いている場所であなたに会います
クロードはあなたが働いている場所であなたに会います
Claude エコシステムに基づいて構築したり、Claude エコシステムとともに成長したり、そこから購入したりできます。お気に入りのサービスに接続したり、Claude に基づいて構築されたパートナー ソリューションを見つけたりしてください。
クロードに接続 クロードに接続 パートナーを表示 パートナーを表示 パートナーを表示 最良の結果を得るためにクロードをカスタマイズ
コネクタはアプリとデータを Claude に取り込み、タスクを完了します。プラグインはコネクタとスキルを組み合わせてワークフローを完成させます。
Slack、Google Workspace、Microsoft Office などを接続します。
ツールは、Web、デスクトップ、モバイル上の Claude Code、Cowork、および Claude に接続できます。
管理者は、各チームが使用できるツールと、それらのツールで何ができるかを制御します。
後付けの必要はありません。毎日使うツールを接続します。
デザイン 開発 プロジェクト管理 営業 法務 財務 アクティブ プロンプト この Figma フレームをレスポンシブ サイトとして構築します。
デザインをコード化された Web サイトに変える
設計をエンジニアリングに引き渡すには時間がかかり、翻訳中に何かが失われる余地が残ります。 Figma を Claude に接続すると、フレームから応答性の高い本番環境に対応したコードまで 1 ステップで移行できます。
最新の Sentry エラーを修正し、PR を開きます。
$ claude "最新の Sentry エラーを修正し、PR を開きます"
C

ロード・クロード・ソネット-4-6 · ~/projects/acme-web
● Sentry の最新号を取得中…
✔ セントリーの問題リスト --project acme-web --limit 1 --status unresolved
ISSUE-4821 TypeError: 未定義のプロパティを読み取れません (「データ」を読み取っています)
プロジェクト：acme-web · 環境：production
最初に確認された日: 2026-06-15 22:14 UTC · 最後に確認された日: 2026-06-16 07:03 UTC
イベント: 1,847 · 影響を受けるユーザー: 312
犯人: src/components/Dashboard/MetricsPanel.tsx 94 行目
● ISSUE-4821 の完全なスタックトレースを取得中…
TypeError: 未定義のプロパティを読み取れません (「データ」を読み取っています)
MetricsPanel.tsx:94:28 で — processChartData(応答)
MetricsPanel.tsx:61:12 で — useEffect(コールバック)
反応-dom/cjs/react-dom.development.js:22794 で
最新のパンくずリスト:
→ GET /api/v2/metrics?range=7d → 200 OK (response.metrics は null でした)
● src/components/Dashboard/MetricsPanel.tsx を読み取り中…
90 ～ 98 行目:
90 │ const processChartData = (応答: ApiResponse) => {
91 │ const { メトリック } = 応答
92 │ // 選択した範囲にデータがない場合、メトリクスは null になる可能性があります
93 │ return metrics.data.map((m: メトリック) => ({
94 │ レーベル: m.label,
95 │ 値: m.値 ?? 0、
96 │ }))
97 │ }
● 根本原因が特定されました。
選択した範囲にデータがない場合、API は { metrics: null } を返します。 93行目
null ガードなしで metrics.data を逆参照し、実行時にスローします。
main ブランチ上 · コミットするものは何もなく、作業ツリーはクリーンです
$ git checkout -b fix/metrics-panel-null-data-sentry-4821
新しいブランチ「fix/metrics-panel-null-data-sentry-4821」に切り替えました
● 修正を src/components/Dashboard/MetricsPanel.tsx に適用しています…
@@ -90,8 +90,8 @@
const processChartData = (応答: ApiResponse) => {
const { メトリクス } = レスポンス
- return metrics.data.map((m: メトリック) => ({
+ if (!metrics?.data) return []
+ return metrics.data.map((m: メトリック) => ({
ラベル:

m.ラベル、
値: m.値 ?? 0、
}))
● src/components/Dashboard/MetricsPanel.tsx レンダー ガードも更新中…
@@ -61,6 +61,7 @@
useEffect(() => {
fetchMetrics(range).then((response) => {
+ if (!response?.metrics) { setChartData([]);戻る }
setChartData(processChartData(応答))
})
}, [範囲])
$ pnpm test --testPathPattern=MetricsPanel
PASS src/components/Dashboard/__tests__/MetricsPanel.test.tsx
✓ 読み込み状態をレンダリングします (12ms)
✓ データが存在する場合にチャートをレンダリングします (34 ミリ秒)
✓ メトリクスが null の場合、空の状態をレンダリングします (8 ミリ秒)
✓ metrics.data が空の配列の場合、空の状態をレンダリングします (6ms)
テストスイート: 1 つ合格、合計 1 つ
テスト: 4 回合格、合計 4 回
$ git add src/components/Dashboard/MetricsPanel.tsx
$ git commit -m "fix(dashboard): MetricsPanel の null metrics.data を防ぐ
Sentry ISSUE-4821 (1,847 イベント、312 ユーザー) を修正しました。
データがない場合、/api/v2/metrics エンドポイントは { metrics: null } を返します
要求された範囲に存在します。 processChartData() が逆参照されました
null チェックを行わない metrics.data により、実行時に TypeError が発生します。
- processChartData() に早期リターン ガードを追加します: if !metrics?.data return []
- useEffect にレスポンス ガードを追加して setChartData 呼び出しを短絡する
- 影響を受けるユーザーには、空白のパネルがクラッシュする可能性があります。空の状態をレンダリングするようになりました
[fix/metrics-panel-null-data-sentry-4821 a3f9c12] 修正(ダッシュボード): null metrics.data を防ぐ...
1 ファイル変更、2 挿入 (+)、1 削除 (-)
● ブランチのプッシュと PR の開始…
$ git Push -u Origin fix/metrics-panel-null-data-sentry-4821
オブジェクトの列挙: 9、完了。
オブジェクトのカウント: 100% (9/9)、完了。
オブジェクトの書き込み: 100% (5/5)、1.02 KiB | 1.02 MiB/秒、完了。
ブランチ「fix/metrics-panel-null-data-sentry-4821」がリモートを追跡するように設定されました。
$ gh pr 作成 \
--title "修正 (ダッシュボード): MetricsPanel の null metrics.data を防ぐ" \
--本文 "..." \

--ベースメイン\
--ラベル「バグ、セントリー」
✓ プルリクエストが作成されました:
https://github.com/acme-org/acme-web/pull/1138
─────────────────
完了しました。 PR #1138 がオープンされました · 2 行が変更されました · 4 つのテストに合格しました
Sentry ISSUE-4821 は、デプロイが完了すると自動解決されます。
制作上のエラーをすぐにレビューできる PR に変える
本番エラーのトリアージとは、修正を 1 行も書く前に、エラー モニター、コードベース、バージョン管理の間を行き来することを意味します。 Sentry と GitHub を Claude に接続し、エラーからプル リクエストに直接進みます。
この Slack スレッドのアクション アイテムを Asana タスクに変換します。
スレッドをアクションアイテムに変える
重要な決定や次のステップは Slack スレッドに埋もれてしまいます。 Slack と Asana を Claude に接続すると、スレッドから離れることなく、あらゆる会話を構造化されたタスク リストに変換できます。
前回の Granola 通話から HubSpot を更新し、フォローアップの草案を作成します。
前回の Granola 通話から HubSpot を更新し、フォローアップの下書きを求めます。
前回の Granola 通話から HubSpot を更新し、フォローアップの下書きを求めます。
前回の Granola 通話から HubSpot を更新し、フォローアップの下書きを求めます。
前回の Granola 通話から HubSpot を更新し、フォローアップの下書きを求めます。
スレッドをアクションアイテムに変える
各電話の後、担当者は取引に集中するのではなく、手動でメモを記録し、フォローアップの下書きに時間を費やします。 Granola、HubSpot、Gmail をクロードに接続し、数分で通話を終了します。
90 日以内に期限切れになる契約を見つけて、更新のリマインダーを追加します。
期限切れの契約を見つけて更新リマインダーを設定する
契約は静かに期限切れになります。誰かがそれを発見したときには、再交渉の余地はすでに閉ざされていることがよくあります。 Ironclad と Google カレンダーをクロードに接続し、更新のたびに先を行きましょう

食べた。
先月の Brex の支出と NetSuite を照合し、不一致にフラグを立てます。
帳簿に照らしてカードの支出を調整し、不一致にフラグを立てます
月末の調整とは、カード取引と会計システムを行ごとに手動で比較することを意味します。 Brex と NetSuite をクロードに接続し、1 つのパスで不一致を表面化します。
Claude を活用したパートナーから購入する
チームがすでに使用しているツール全体に、既存の Anthropic への取り組みを適用します。
Claude を活用した複数のパートナー ツールにわたって、既存の Anthropic コミットメントを活用します。
エンタープライズ チーム向けに構築された Claude を利用したツールを参照してください。評価に費やす時間を減らし、構築に多くの時間を費やします。
ニーズの進化に応じてパートナーを追加してください。あなたのコミットメントは組織に合わせて変化します。
ディレクトリとマーケットプレイスの違いは何ですか?
このディレクトリにより、ツールが Claude に取り込まれます。マーケットプレイスは、コミットメントされた支出を使用して、パートナー ソリューション内にクロードをもたらします。多くのチームが両方を使用しています。
コネクタとプラグインの違いは何ですか?
コネクタを使用すると、Claude は単一のアプリに直接アクセスできます。プラグインはコネクタと命令を組み合わせるため、Claude は毎回設定しなくても、複数のツールにわたる完全なワークフローを処理できます。
前へ 前へ 次へ 次へ ホームページ ホームページ 次へ 次へ ありがとうございます!あなたの提出物は受理されました！おっと！フォームの送信中に問題が発生しました。書き込みボタン テキスト ボタン テキスト 学習ボタン テキスト ボタン テキスト コード ボタン テキスト ボタン テキスト 書き込み 聴衆向けにユニークな声を開発するのを手伝ってください こんにちはクロード!聴衆に合わせたユニークな声を開発するのを手伝ってくれませんか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。 Google ドライブ、ウェブ検索など、アクセス権のあるツールを使用できます。

—このタスクをより適切に達成するのに役立つのであれば。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
文章のスタイルを改善する こんにちはクロード!私の書き方を改善してもらえませんか？私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
創造的なアイデアをブレインストーミングします、こんにちはクロード!創造的なアイデアをブレインストーミングしていただけますか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。

ご協力いただきありがとうございます！
Learn 複雑なトピックを簡単に説明する こんにちはクロード!複雑なトピックを簡単に説明してもらえますか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
これらのアイデアを理解するのを手伝ってください、こんにちはクロード!これらの考えを理解するのを手伝ってもらえますか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔、会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どのような種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) が考えられるかを検討してください。

[切り捨てられた]

## Original Extract

Connect the tools your team already relies on to Claude. Discover the directory and marketplace.

Claude ecosystem | Claude by Anthropic
Meet Claude Products Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Claude ecosystem Claude ecosystem
Claude meets you where you work
Claude meets you where you work
Build on, grow with, or buy from the Claude ecosystem. Connect your favorite services or find partner solutions built on Claude.
Connect to Claude Connect to Claude View partners View partners View partners Customize Claude for the best results
Connectors bring your apps and data into Claude to complete tasks. Plugins combine connectors and skills to complete workflows.
Connect Slack, Google Workspace, Microsoft Office, and more.
Your tools can be connected to Claude Code, Cowork, and Claude on web, desktop, and mobile.
Admins control which tools each team can use, and what those tools can do.
No retrofitting required. Connect the tools you use everyday.
Design Development Project management Sales Legal Finance Active Prompt Build this Figma frame as a responsive site.
Turn designs into coded websites
Handing off designs to engineering takes time and leaves room for things to get lost in translation. Connect Figma to Claude and go from frame to responsive, production-ready code in one step.
Fix the latest Sentry error and open a PR.
$ claude "Fix the latest Sentry error and open a PR"
Claude claude-sonnet-4-6 · ~/projects/acme-web
● Fetching latest Sentry issues…
✔ sentry issues list --project acme-web --limit 1 --status unresolved
ISSUE-4821 TypeError: Cannot read properties of undefined (reading 'data')
Project: acme-web · Environment: production
First seen: 2026-06-15 22:14 UTC · Last seen: 2026-06-16 07:03 UTC
Events: 1,847 · Users affected: 312
Culprit: src/components/Dashboard/MetricsPanel.tsx line 94
● Fetching full stacktrace for ISSUE-4821…
TypeError: Cannot read properties of undefined (reading 'data')
at MetricsPanel.tsx:94:28 — processChartData(response)
at MetricsPanel.tsx:61:12 — useEffect(callback)
at react-dom/cjs/react-dom.development.js:22794
Most recent breadcrumb:
→ GET /api/v2/metrics?range=7d → 200 OK (response.metrics was null)
● Reading src/components/Dashboard/MetricsPanel.tsx…
Line 90–98:
90 │ const processChartData = (response: ApiResponse) => {
91 │ const { metrics } = response
92 │ // metrics can be null when no data for the selected range
93 │ return metrics.data.map((m: Metric) => ({
94 │ label: m.label,
95 │ value: m.value ?? 0,
96 │ }))
97 │ }
● Root cause identified.
API returns { metrics: null } when the selected range has no data. Line 93
dereferences metrics.data without a null guard, throwing at runtime.
On branch main · nothing to commit, working tree clean
$ git checkout -b fix/metrics-panel-null-data-sentry-4821
Switched to a new branch 'fix/metrics-panel-null-data-sentry-4821'
● Applying fix to src/components/Dashboard/MetricsPanel.tsx…
@@ -90,8 +90,8 @@
const processChartData = (response: ApiResponse) => {
const { metrics } = response
- return metrics.data.map((m: Metric) => ({
+ if (!metrics?.data) return []
+ return metrics.data.map((m: Metric) => ({
label: m.label,
value: m.value ?? 0,
}))
● Also updating src/components/Dashboard/MetricsPanel.tsx render guard…
@@ -61,6 +61,7 @@
useEffect(() => {
fetchMetrics(range).then((response) => {
+ if (!response?.metrics) { setChartData([]); return }
setChartData(processChartData(response))
})
}, [range])
$ pnpm test --testPathPattern=MetricsPanel
PASS src/components/Dashboard/__tests__/MetricsPanel.test.tsx
✓ renders loading state (12ms)
✓ renders chart when data present (34ms)
✓ renders empty state when metrics is null (8ms)
✓ renders empty state when metrics.data is empty array (6ms)
Test Suites: 1 passed, 1 total
Tests: 4 passed, 4 total
$ git add src/components/Dashboard/MetricsPanel.tsx
$ git commit -m "fix(dashboard): guard against null metrics.data in MetricsPanel
Fixes Sentry ISSUE-4821 (1,847 events, 312 users).
The /api/v2/metrics endpoint returns { metrics: null } when no data
exists for the requested range. processChartData() dereferenced
metrics.data without a null check, causing a TypeError at runtime.
- Add early return guard in processChartData(): if !metrics?.data return []
- Add response guard in useEffect to short-circuit setChartData call
- Affected users would see a blank panel crash; now renders empty state
[fix/metrics-panel-null-data-sentry-4821 a3f9c12] fix(dashboard): guard against null metrics.data...
1 file changed, 2 insertions(+), 1 deletion(-)
● Pushing branch and opening PR…
$ git push -u origin fix/metrics-panel-null-data-sentry-4821
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Writing objects: 100% (5/5), 1.02 KiB | 1.02 MiB/s, done.
Branch 'fix/metrics-panel-null-data-sentry-4821' set up to track remote.
$ gh pr create \
--title "fix(dashboard): guard against null metrics.data in MetricsPanel" \
--body "..." \
--base main \
--label "bug,sentry"
✓ Pull request created:
https://github.com/acme-org/acme-web/pull/1138
──────────────────────────────────────────────────
Done. PR #1138 opened · 2 lines changed · 4 tests passing
Sentry ISSUE-4821 will auto-resolve once the deploy goes out.
Turn a production error into a ready-to-review PR
Triaging a production error means jumping between your error monitor, your codebase, and your version control before you can write a single line of fix. Connect Sentry and GitHub to Claude and go straight from error to pull request.
Turn this Slack thread's action items into Asana tasks.
Turn threads into action items
Important decisions and next steps get buried in Slack threads. Connect Slack and Asana to Claude and turn any conversation into a structured task list without leaving the thread.
Update HubSpot from my last Granola call and draft a follow-up.
Prompt Update HubSpot from my last Granola call and draft a follow-up.
Prompt Update HubSpot from my last Granola call and draft a follow-up.
Prompt Update HubSpot from my last Granola call and draft a follow-up.
Prompt Update HubSpot from my last Granola call and draft a follow-up.
Turn threads into action items
After every call, reps spend time manually logging notes and drafting follow-ups instead of focusing on the deal. Connect Granola, HubSpot, and Gmail to Claude and close out the call in minutes.
Find contracts expiring in 90 days and add renewal reminders.
Find expiring contracts and set renewal reminders
Contracts expire quietly. By the time someone catches it, the window to renegotiate has often closed. Connect Ironclad and Google Calendar to Claude and stay ahead of every renewal date.
Reconcile last month's Brex spend against NetSuite and flag mismatches.
Reconcile card spend against the books and flag mismatches
Month-end reconciliation means manually comparing card transactions against your accounting system, line by line. Connect Brex and NetSuite to Claude and surface mismatches in one pass.
Purchase from Claude-powered partners
Apply your existing Anthropic commitment across the tools your team already uses.
Use your existing Anthropic commitment across multiple Claude-powered partner tools.
Browse Claude-powered tools built for enterprise teams. Spend less time evaluating, more time building.
Add partners as your needs evolve. Your commitment flexes with your organization.
What's the difference between the directory and the marketplace?
The directory brings your tools into Claude. The marketplace brings Claude to you inside partner solutions, using your committed spend. Many teams use both.
What's the difference between connectors and plugins?
Connectors give Claude direct access to a single app. Plugins combine connectors with instructions, so Claude can handle complete workflows across multiple tools without you having to set it up each time.
Prev Prev Next Next Homepage Homepage Next Next Thank you! Your submission has been received! Oops! Something went wrong while submitting the form. Write Button Text Button Text Learn Button Text Button Text Code Button Text Button Text Write Help me develop a unique voice for an audience Hi Claude! Could you help me develop a unique voice for an audience? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Improve my writing style Hi Claude! Could you improve my writing style? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Brainstorm creative ideas Hi Claude! Could you brainstorm creative ideas? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Learn Explain a complex topic simply Hi Claude! Could you explain a complex topic simply? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Help me make sense of these ideas Hi Claude! Could you help me make sense of these ideas? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be

[truncated]
