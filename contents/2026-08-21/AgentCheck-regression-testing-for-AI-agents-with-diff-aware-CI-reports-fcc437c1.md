---
source: "https://github.com/rez-99/agentcheck"
hn_url: "https://news.ycombinator.com/item?id=49393322"
title: "AgentCheck – regression testing for AI agents, with diff-aware CI reports"
article_title: "GitHub - rez-99/agentcheck: Regression testing for AI agents — pytest-style YAML tests, LLM-as-judge scoring, and diff-aware CI reports that show what actually got worse. · GitHub"
image: "https://opengraph.githubassets.com/c3a7b49cc6da1496ae4d9545858bf104123d2591fee0bcb46c6e37ea3db41135/rez-99/agentcheck"
author: "zz99"
captured_at: "2026-08-21T21:15:03Z"
capture_tool: "hn-digest"
hn_id: 49393322
score: 1
comments: 0
posted_at: "2026-08-21T20:21:15Z"
tags:
  - hacker-news
  - translated
---

# AgentCheck – regression testing for AI agents, with diff-aware CI reports

- HN: [49393322](https://news.ycombinator.com/item?id=49393322)
- Source: [github.com](https://github.com/rez-99/agentcheck)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T20:21:15Z

## Translation

タイトル: AgentCheck – 差分認識 CI レポートによる AI エージェントの回帰テスト
記事のタイトル: GitHub - rez-99/agentcheck: AI エージェントの回帰テスト — pytest スタイルの YAML テスト、LLM-as-judge スコアリング、および実際に何が悪化したかを示す diff 対応 CI レポート。 · GitHub
説明: AI エージェントの回帰テスト — pytest スタイルの YAML テスト、LLM-as-judge スコアリング、および実際に何が悪化したかを示す差分認識 CI レポート。 - rez-99/エージェントチェック

記事本文:
GitHub - rez-99/agentcheck: AI エージェントの回帰テスト — pytest スタイルの YAML テスト、LLM-as-judge スコアリング、および実際に何が悪化したかを示す差分認識 CI レポート。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
レズ99
/
エージェントチェック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
4 コミット 4 コミット フォルダーとファイル
.github/ workflows .github/ workflows 例 例 結果 結果 src/ Agentcheck src/ agen

tcheck テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントの回帰テスト。エージェントが行うべきことをプレーン YAML で定義します。
実際のエージェント (CLI コマンドまたは HTTP エンドポイント) に対して実行し、LLM に判断させます。
指定した基準に照らしてすべての回答を採点し、合格/不合格とその理由を付けます。に配線します
CI を使用するため、迅速な変更、ツールの交換、またはモデルのアップグレードによって、動作が静かに破壊されることはありません。
ユーザーは依存します。
これは意図的に狭い範囲です。実稼働環境の可観測性プラットフォームではありません (つまり、
Langfuse / Braintrust / Arize の領域にあり、資金も豊富です。正面から競争しないでください)。
これは、まだほとんど誰もうまく構築できていないものです。高速で開発者に優しい事前デプロイメントです。
pytest と同じ方法で、GitHub Actions ステップに適合するかどうかを確認します。
完全な可観測性/評価プラットフォームは、昨年 5,000 万ドルから 8,000 万ドルのラウンドを調達しており、
「実稼働エージェントのトラフィックのログと監視」スペースを積極的に混雑させます。
それらのほとんどは、単独で軽量の git ネイティブ回帰スイートとして構築されていません。
開発者は 5 分で CI を追加できます。このギャップがくさびです。
配布はセルフサービス/PLG (オープンソース CLI、開発者対象) ではなく、
エンタープライズ営業 — 個人の AI 創設者が一貫して最も苦手とすること。
pip install -e 。
エクスポート ANTHROPIC_API_KEY=sk-...
Agentcheck の実行例/tests.yaml
テストファイル形式
example/tests.yaml を参照してください。各テスト ケースは入力、つまり平易な英語を指定します。
正しい応答がどのようなものかについての説明と、あなたの意見に到達するための 2 つの方法のうちの 1 つ
エージェント:
コマンド: "..." — サブプロセスとして実行します。入力は標準入力にパイプされ、標準出力は
出力としてキャプチャされます。どの言語でも動作します。
エージェント: "module.path:function_name" — そのモジュールをインポートして関数を呼び出す
仕掛品

s 唯一の引数として入力を指定します。その戻り値は出力です。役に立つ
LangGraph/CrewAI/Claude-Agent-SDK スタイルのエージェントの場合は、むしろ Python 呼び出し可能です
スタンドアロン CLI スクリプトよりも — example/inprocess_agent.py を参照してください。
ケースごとに 2 つのうち 1 つだけが必要です。いずれの場合も、出力にはスコアが付けられます
LLM 審査員による 1 行の理由で合否が判定されます。脆弱な文字列一致はありません。
Agentcheck run testing.yaml --json-out results.json は、次のような JSON レポートを書き込みます。
ビルド アーティファクトとしてアップロードします (examples/.github/workflows/agentcheck.yml を参照)。
--post-pr-comment を追加し、プルリクエストで GITHUB_TOKEN を設定して実行します (ジョブ
権限が必要です: プルリクエスト: write )、agentcheck がマークダウン概要テーブルを投稿します
PR コメントとして、新しいコメントを蓄積するのではなく、繰り返し実行時に同じコメントを更新します。
もの。それ以外の場所ではサイレントノーオペレーション (プッシュ、ローカル実行) なので、安全に終了できます。
すべての CI 呼び出しでオンになります。
ベースラインとの比較 (「この変更により良くなったのか、悪くなったのか」)
一定のパス数 (「18/20 パス」) だけでは、変更が役に立ったか悪かったかはわかりません。
表を読みに行かなければなりません。 --baseline は、現在の実行との差分を取得することで問題を修正します。
テスト名をキーとした以前の --json-out レポートと比較します。
Agentcheck run testing.yaml --json-out results.json --baselinebaseline.json
バンドルされたサンプルに対してローカルで試してみます。
Agentcheck run 例/tests.yaml --baseline 例/baseline.json
すべてのテストは 1 つのバケットに収まります: 変更なし (ベースラインと同じ合格/不合格)、
回帰 (ベースラインは合格しましたが、今度は失敗します。これが重要です)、
改善された (ベースラインは失敗したが、現在は合格)、新規 (ベースラインにない)、または
削除されました (ベースラインではありますが、この実行では削除されていません - おそらくテスト ケースが削除されています)
一見の価値あり）。コンソールには 1 行の概要と回帰の表が出力されます。
および改善。行方不明または

読めないベースライン (ベースラインはまだありません)
リポジトリの最初の実行) は警告を出力し、単純な合格/失敗レポートに戻ります。
実行全体を失敗するのではなく。
--baseline も設定されている場合、--post-pr-comment はこれを自動的に取得します。
PR コメントは「ベースラインとの比較: 2 つは変化なし、1 つは改善、1 つは後退」で始まり、次のように呼びかけています。
完全な結果表を単に再説明するのではなく、回帰を具体的に取り除きます。
これを実際に CI に接続するには、ベースラインの取得元となる場所が必要です。
通常のパターンは次のとおりです: デフォルトのブランチにプッシュするたびに、agentcheck を実行します。
--json-out results/baseline.json を選択し、そのファイルをリポジトリにコミットして戻します。あらゆるPRにおいて、
そのファイルのベース ブランチのコピーを読み取ります ( git showorigin/main:results/baseline.json )
それを --baseline として渡します。詳細については、examples/.github/workflows/agentcheck.yml を参照してください。
それの作業バージョン。
pip install -e " .[dev] "
pytest
ロードマップ (順番に実行し、先に飛ばさないでください)
今週: この CLI を独自のおもちゃエージェントに対してエンドツーエンドで動作させます。ドッグフードそれ。
第 2 週: オープンソース化します。エージェントビルダーが実際にたむろする場所 (r/LocalLLaMA、
LangChain/LlamaIndex Discord、Hacker News「Show HN」、関連する X スレッド）。目標
それはバイラル性ではありません — まさにこの問題点に達し、それを伝える 10 ～ 20 人を見つけることです
足りないものはあなたです。
3 ～ 4 週目: 最も求められているものを追加します。考えられる候補: ホストされたダッシュボード
実行履歴、Slack/GitHub PR コメント レポート、または特定の人気エージェントのサポート
フレームワーク (LangGraph、CrewAI、Claude Agent SDK) をファーストクラスの統合として提供します。
無料の CLI を定期的に使用するようになった場合のみ、有料のホスト層を導入します
(実行履歴、チーム共有、傾向グラフ) — 収益化を実現する前に収益化を構築しないでください。
無料ユーザーは、それがなくなったら見逃してしまうでしょう。
将軍になろうとするな

可観測性プラットフォーム — それは、混雑していて資金が豊富なプラットフォームです
レーン。 「高速な git ネイティブ回帰チェック」ツールを使い続けてください。 1 つの狭い仕事の深さが勝る
資金提供を受けた競合他社に対する幅広さ。
AI エージェントの回帰テスト — pytest スタイルの YAML テスト、ジャッジとしての LLM スコアリング、および実際に何が悪化したかを示す差分認識 CI レポート。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Regression testing for AI agents — pytest-style YAML tests, LLM-as-judge scoring, and diff-aware CI reports that show what actually got worse. - rez-99/agentcheck

GitHub - rez-99/agentcheck: Regression testing for AI agents — pytest-style YAML tests, LLM-as-judge scoring, and diff-aware CI reports that show what actually got worse. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
rez-99
/
agentcheck
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.github/ workflows .github/ workflows examples examples results results src/ agentcheck src/ agentcheck tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Regression testing for AI agents. Define what your agent should do in plain YAML,
run it against your actual agent (any CLI command or HTTP endpoint), and let an LLM judge
score every response against your stated criteria — pass/fail, with a reason. Wire it into
CI so a prompt change, a tool swap, or a model upgrade can't silently break behavior your
users depend on.
This is deliberately narrow: it is not a production-observability platform (that's
Langfuse / Braintrust / Arize territory, and they're well-funded — don't compete head-on).
It's the thing almost nobody has built well yet: a fast, dev-friendly pre-deployment
check that fits in a GitHub Actions step the same way pytest does.
Full observability/eval platforms have raised $50–80M rounds in the last year and are
actively crowding the "log + monitor production agent traffic" space.
Almost none of them are built as a lightweight, git-native regression suite a solo
developer can add to CI in five minutes — that gap is the wedge.
Distribution is self-serve/PLG (open-source CLI, developer audience) rather than
enterprise sales — the thing solo AI founders are consistently worst at.
pip install -e .
export ANTHROPIC_API_KEY=sk-...
agentcheck run examples/tests.yaml
Test file format
See examples/tests.yaml . Each test case specifies an input, a plain-English
description of what a correct response looks like, and one of two ways to reach your
agent:
command: "..." — run it as a subprocess; input is piped to stdin, stdout is
captured as the output. Works with any language.
agent: "module.path:function_name" — import that module and call the function
in-process with input as its only argument; its return value is the output. Useful
for LangGraph/CrewAI/Claude-Agent-SDK-style agents that are Python callables rather
than standalone CLI scripts — see examples/inprocess_agent.py .
Exactly one of the two is required per case. Either way, the output is scored
pass/fail with a one-line reason by an LLM judge — no brittle string matching.
agentcheck run tests.yaml --json-out results.json writes a JSON report you can
upload as a build artifact (see examples/.github/workflows/agentcheck.yml ).
Add --post-pr-comment and, on a pull-request run with GITHUB_TOKEN set (the job
needs permissions: pull-requests: write ), agentcheck posts a markdown summary table
as a PR comment, updating the same comment on repeat runs instead of piling up new
ones. It's a silent no-op everywhere else (pushes, local runs), so it's safe to leave
on in every CI invocation.
Comparing against a baseline ("did this change make it better or worse")
A flat pass count ("18/20 passed") doesn't tell you whether a change helped or hurt —
you have to go read the table. --baseline fixes that by diffing the current run
against a previous --json-out report, keyed by test name:
agentcheck run tests.yaml --json-out results.json --baseline baseline.json
Try it locally against the bundled example:
agentcheck run examples/tests.yaml --baseline examples/baseline.json
Every test lands in one bucket: unchanged (same pass/fail as the baseline),
regressed (baseline passed, now fails — this is the one you care about),
improved (baseline failed, now passes), new (not in the baseline), or
removed (in the baseline but not in this run — probably a deleted test case,
worth a glance). The console prints a one-line summary plus a table of regressions
and improvements; a missing or unreadable baseline (there's no baseline yet on a
repo's first run) prints a warning and falls back to the plain pass/fail report
instead of failing the whole run.
--post-pr-comment picks this up automatically when --baseline is also set, so the
PR comment leads with "vs baseline: 2 unchanged, 1 improved, 1 regressed" and calls
out the regressions specifically, instead of just restating the full results table.
To actually wire this into CI you need somewhere for the baseline to come from — the
usual pattern is: on every push to your default branch, run agentcheck with
--json-out results/baseline.json and commit that file back to the repo; on every PR,
read the base branch's copy of that file ( git show origin/main:results/baseline.json )
and pass it as --baseline . See examples/.github/workflows/agentcheck.yml for a full
working version of that.
pip install -e " .[dev] "
pytest
Roadmap (do this in order, don't skip ahead)
This week: get this CLI working end to end against your own toy agent. Dogfood it.
Week 2: open-source it. Post it where agent builders actually hang out (r/LocalLLaMA,
the LangChain/LlamaIndex Discords, Hacker News "Show HN", relevant X threads). The goal
isn't virality — it's finding 10-20 people who hit this exact pain point and will tell
you what's missing.
Week 3-4: add the thing they ask for most. Likely candidates: a hosted dashboard for
run history, Slack/GitHub PR-comment reporting, or support for a specific popular agent
framework (LangGraph, CrewAI, the Claude Agent SDK) as a first-class integration.
Only once people are using the free CLI regularly: introduce a paid hosted tier
(run history, team sharing, trend charts) — don't build monetization before you have
free users who'd miss it if it disappeared.
Don't try to become a general observability platform — that's the crowded, well-funded
lane. Stay the "fast, git-native regression check" tool. Depth in one narrow job beats
breadth against funded competitors.
Regression testing for AI agents — pytest-style YAML tests, LLM-as-judge scoring, and diff-aware CI reports that show what actually got worse.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
