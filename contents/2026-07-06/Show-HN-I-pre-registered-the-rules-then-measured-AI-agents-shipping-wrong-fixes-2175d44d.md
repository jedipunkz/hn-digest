---
source: "https://github.com/kolesnikov-arch/patchward"
hn_url: "https://news.ycombinator.com/item?id=48804907"
title: "Show HN: I pre-registered the rules,then measured AI agents shipping wrong fixes"
article_title: "GitHub - kolesnikov-arch/patchward: Evidence home for the verdict-layer approach: pre-registered, reproducible trust evaluation of AI-generated code changes - silent wrong ships 17/50 → 0/50. · GitHub"
author: "kolesnikov-arch"
captured_at: "2026-07-06T15:02:06Z"
capture_tool: "hn-digest"
hn_id: 48804907
score: 1
comments: 0
posted_at: "2026-07-06T14:12:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I pre-registered the rules,then measured AI agents shipping wrong fixes

- HN: [48804907](https://news.ycombinator.com/item?id=48804907)
- Source: [github.com](https://github.com/kolesnikov-arch/patchward)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T14:12:47Z

## Translation

タイトル: HN を表示: ルールを事前登録し、AI エージェントが間違った修正を送信することを測定しました
記事のタイトル: GitHub - kolesnikov-arch/patchward: 評決層アプローチの証拠の本拠地: AI 生成コード変更の事前登録済み、再現可能な信頼評価 - サイレント間違った船 17/50 → 0/50。 · GitHub
説明: 評決層アプローチの証拠の拠点: AI によって生成されたコード変更の事前登録された再現可能な信頼評価 - サイレント間違った出荷 17/50 → 0/50。 - kolesnikov-arch/patchward
HN テキスト: SWE-bench Lite での非ゲート AI パッチ適用とゲート AI パッチ適用の制御された比較。リポジトリ内の完全な方法論、結果、再現可能なスクリプト。フラッシュモデルの画期的な進歩。

記事本文:
GitHub - kolesnikov-arch/patchward: 評決層アプローチの証拠の本拠地: AI によって生成されたコード変更の事前登録された再現可能な信頼評価 - サイレント間違った船 17/50 → 0/50。 · GitHub
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
コレスニコフアーチ
/
パッチワード
公共
いいえ

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット 評価アーティファクト 評価アーティファクト .gitignore .gitignore CURRENT_SCOPE_AND_LIMITATIONS.md CURRENT_SCOPE_AND_LIMITATIONS.md CURRENT_SCOPE_AND_LIMITATIONS_zh.md CURRENT_SCOPE_AND_LIMITATIONS_zh.mdライセンス ライセンス PREREGISTRATION.md PREREGISTRATION.md PREREGISTRATION_zh.md PREREGISTRATION_zh.md README.md README.md README_zh.md README_zh.md RESULTS.md RESULTS.md RESULTS_zh.md RESULTS_zh.md すべてのファイルを表示 リポジトリファイルナビゲーション
AI コーディング エージェント用の判定レイヤー - そのため、エージェントは壊れた修正を黙って出荷することができません。
📊 最初に保留された結果が公開されます: RESULTS.md — 凍結された、
事前に登録された 50 のタスクのセット、同じモデルが 17/50 の間違った修正をサイレントに出荷
ゲートなしと 0/50 ゲートあり。から再計算可能なすべての数値
評価成果物/ 。強化されたエンジンは非公開です
デザイン。ここで公開されているのは証拠なので、主張を自分で確認できます。
ベンチマークは能力を測定します。 CI には信頼が必要です。
AI コーディング エージェントは、解決率、つまり修正したバグの数に基づいてスコア付けされます。それは
能力 、そして最も広く使用されているエージェント ベンチマークはそれを測定します。その数字は、
実際にエージェントを CI に含めることができるかどうかを決定するのは、まったく異なります。
あまり報告されていないもの: 誤受理率 - エージェントが自信を持って出荷する頻度
一見正しく見えても、実は正しくない修正。
エージェントが自分で宿題を採点するため、事態はさらに悪化します。
エージェントが修正を作成する → エージェントがそれに対するテストを作成する → テストが合格する
(エージェント自身の誤解をコード化します) → エージェントは成功を報告します。
緑色のチェックマーク、隠れたバグ。勝てるように見えるベンチマーク上。本番環境では、
fi という回帰

負荷がかかっている状態。
アプローチ: 生成と判断を分離する
確率的 LLM は次のように提案します。独立した決定論的な評決レイヤー
決定します。エージェント自体ではなく小切手が評決を下します。
二値的な合格/不合格ではなく、3 つの正直な結果:
🟢 検証済み — 独立して確認済み (例: 既存のテストが分離されたテストで合格するなど)
コンテナ。変更は範囲内であり、後退しません)。
🟡 未検証 — もっともらしいが、独立して確認されていない → 人間によるフラグが立てられている
レビュー。このシステムは、ゴム印ではなく、チェックできなかったことについて正直に報告します。
🔴 ブロック済み — ゲートによってキャッチされました (宣言された範囲外の編集、チェックの失敗、回帰)。
重要なのは説明責任です。説明責任のない AI パッチを、責任を負う AI パッチに変えるのです。
テストの実行の有無にかかわらず、判定と監査証跡。
独立した決定論的なチェック — スコープ、証拠、回帰 —
代理人ではなく評決を下す。それらのほとんどはコードを実行する必要がないため、
オフライン、エアギャップ、あらゆる言語で作業できます。個別のテスト実行は追加です
ランタイムが利用可能な場合の確認レイヤー。
この背後にある理由 - なぜ独立した評決が代理人の自己報告を上回るのか
そしてなぜその決定が二者択一ではなく三者状態であるのかが文書化されています（要約、
コンパニオンでは実装名とツール名が削除されています)
判定レイヤーフレームワーク。
それを実装する特定の調整されたエンジンは非公開のままです。
何が公開されていて、何が公開されていないのか
コンセプトと測定方法。
RESULTS.md — 保留された結果: 見出し数、正確な CI、全文
性質テーブル、悲観的な感度行、根本原因のあるすべての誤拒否、および
実行整合性ログ。
事前登録された採点契約 - 採点ルール、レポート
約束とその答え

o 厳しい反対意見には結果の前に日付が記入されている
存在していたので、ルールが結果に適合していなかった可能性があります。
現在の範囲と制限 - 測定内容
あるのか、まだ示されていないのか（構造的には正直です）。
評価アーティファクト/ — 再現可能な証明キット: 両腕'
評価されたとおりの正確な予測、生のインスタンスごとのスコアリングレポート、ペアの結果テーブル、
シードされたタスクの選択と、すべての見出しの数値を再計算する stdlib 専用のスクリプト。
動作中の判定ロジックのインタラクティブなシミュレーション (
コンパニオンレポ）。
設計によるプライベート: 調整されたエンジン — ゲート、プロンプト、および
彼らが調整している失敗記憶コーパス。何百もの音から抽出されたそのチューニング
実際に走る、それが仕事です。私たちはエンジンではなく証拠を公開します - したがって、主張は
堀を渡さずにチェック可能。
評価 #1 が公開されました (2026-07-05): RESULTS.md 。採点ルール
結果が存在する前に事前登録されていた ( PREREGISTRATION.md 、
2026-07-03)、すべての数値は次から再計算可能です。
評価-アーティファクト/ — 誇大広告よりも誠実さを構築する。
コンセプトと信頼性の理論はコンパニオン リポジトリにあります。
評決層フレームワーク 。
評価実行時のフィールド ノート:
AI デリバリーを信頼してください (ニュースレター)。
コンセプト、ドキュメント、および結果: CC BY-NC 4.0 (
評決層フレームワーク)。
デモ/サンプルコード (追加時): MIT。 「ライセンス」を参照してください。
評決層アプローチの証拠の拠点: AI によって生成されたコード変更の事前登録された再現可能な信頼評価 - サイレント間違った出荷 17/50 → 0/50。
kolesnikov-arch.github.io/verdict-layer-framework/architecture_logic_sim/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
そこw

ロード中のエラーとして。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Evidence home for the verdict-layer approach: pre-registered, reproducible trust evaluation of AI-generated code changes - silent wrong ships 17/50 → 0/50. - kolesnikov-arch/patchward

A controlled comparison of ungated vs gated AI patching on SWE-bench Lite. Full methodology, results, and reproducible scripts in the repo. Breakthrough for flash models.

GitHub - kolesnikov-arch/patchward: Evidence home for the verdict-layer approach: pre-registered, reproducible trust evaluation of AI-generated code changes - silent wrong ships 17/50 → 0/50. · GitHub
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
kolesnikov-arch
/
patchward
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits evaluation-artifacts evaluation-artifacts .gitignore .gitignore CURRENT_SCOPE_AND_LIMITATIONS.md CURRENT_SCOPE_AND_LIMITATIONS.md CURRENT_SCOPE_AND_LIMITATIONS_zh.md CURRENT_SCOPE_AND_LIMITATIONS_zh.md LICENSE LICENSE PREREGISTRATION.md PREREGISTRATION.md PREREGISTRATION_zh.md PREREGISTRATION_zh.md README.md README.md README_zh.md README_zh.md RESULTS.md RESULTS.md RESULTS_zh.md RESULTS_zh.md View all files Repository files navigation
A verdict layer for AI coding agents — so an agent can't quietly ship a broken fix.
📊 First held-out results are published: RESULTS.md — on a frozen,
pre-registered set of 50 tasks, the same model silently shipped 17/50 wrong fixes
ungated vs 0/50 gated. Every number recomputable from
evaluation-artifacts/ . The hardened engine is private by
design; what's public here is the evidence , so you can check the claims yourself.
Benchmarks measure capability. CI needs trust.
AI coding agents are scored on resolve rate — how many bugs they fix. That's
capability , and most widely-used agent benchmarks measure it. The number that
actually decides whether you can let an agent into your CI is a different, far
less-reported one: the false-accept rate — how often an agent confidently ships
a fix that looks right and silently isn't.
It gets worse, because agents grade their own homework:
the agent writes the fix → the agent writes a test for it → the test passes
(it encodes the agent's own misunderstanding) → the agent reports success.
Green checkmark, hidden bug. On a benchmark that looks like a win. In production it's
the regression that fires under load.
The approach: separate generation from judgement
A probabilistic LLM proposes ; an independent, deterministic verdict layer
decides . The checks, not the agent itself, render the verdict — and there are
three honest outcomes, not a binary pass/fail:
🟢 Verified — independently confirmed (e.g. existing tests pass in an isolated
container; the change is in scope and doesn't regress).
🟡 Unverified — plausible, but not independently confirmed → flagged for human
review. The system is honest about what it couldn't check, instead of rubber-stamping.
🔴 Blocked — caught by a gate (edit outside declared scope, failed check, regression).
The point is accountability : turn an unaccountable AI patch into one that carries
a verdict and an audit trail — with or without test execution.
Independent, deterministic checks — for scope , evidence , and regression —
render the verdict, not the agent. Most of them don't require executing code, so they
work offline, air-gapped, in any language; an isolated test run is an additional
confirmation layer when a runtime is available.
The reasoning behind this — why an independent verdict beats an agent's self-report,
and why the decision is three-state rather than binary — is documented (abstracted,
with implementation and tool names stripped) in the companion
Verdict Layer Framework .
The specific tuned engine that implements it stays private.
What's public — and what isn't
the concept and the measurement methodology;
RESULTS.md — the held-out results: headline counts, exact CIs, the full
disposition table, the pessimistic sensitivity row, every false-reject with root cause, and
the run-integrity log;
Pre-registered Scoring Contract — the scoring rules, the reporting
commitments, and the answers to the hard objections, date-stamped before the results
existed — so the rules provably weren't fitted around the outcome;
Current Scope & Limitations — what the measurement
does and doesn't yet show (honest by construction);
evaluation-artifacts/ — the reproducible proof-kit: both arms'
predictions exactly as evaluated, raw per-instance scoring reports, the paired results table,
the seeded task selection, and a stdlib-only script that recomputes every headline figure;
an interactive sim of the verdict logic in action (in the
companion repo ) .
Private, by design: the tuned engine — the gates, the prompts, and the
failure-memory corpus they're tuned against. That tuning, distilled from hundreds of
real runs, is the work. We open the evidence , not the engine — so the claims
are checkable without handing over the moat.
Evaluation #1 published (2026-07-05): RESULTS.md . The scoring rules
were pre-registered before the outcome existed ( PREREGISTRATION.md ,
2026-07-03) and every number is recomputable from
evaluation-artifacts/ — honesty over hype, by construction.
The concept and the trust thesis live in the companion repo:
verdict-layer-framework .
Field notes from the evaluation as it runs:
Trust in AI Delivery (newsletter).
Concept, documentation, and results: CC BY-NC 4.0 (matching the
Verdict Layer Framework ).
Demo / sample code (when added): MIT. See LICENSE .
Evidence home for the verdict-layer approach: pre-registered, reproducible trust evaluation of AI-generated code changes - silent wrong ships 17/50 → 0/50.
kolesnikov-arch.github.io/verdict-layer-framework/architecture_logic_sim/
Topics
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
