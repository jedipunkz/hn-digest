---
source: "https://github.com/couldbeme/holdline"
hn_url: "https://news.ycombinator.com/item?id=49338963"
title: "I built a benchmark for AI agent guardrails, and it caught my own plugin lying"
article_title: "GitHub - couldbeme/holdline: Holdline — a neutral benchmark for AI-agent write-guards. Catch rate, false-block rate, class-balanced kappa, and an injection-attack class. Scores any guard. · GitHub"
image: "https://opengraph.githubassets.com/9274082662f26314c9061d9c93d5c63ecc6c46b7fb272557529c5eaf7e1278e3/couldbeme/holdline"
author: "couldbeme_"
captured_at: "2026-08-17T23:13:45Z"
capture_tool: "hn-digest"
hn_id: 49338963
score: 1
comments: 0
posted_at: "2026-08-17T23:11:48Z"
tags:
  - hacker-news
  - translated
---

# I built a benchmark for AI agent guardrails, and it caught my own plugin lying

- HN: [49338963](https://news.ycombinator.com/item?id=49338963)
- Source: [github.com](https://github.com/couldbeme/holdline)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T23:11:48Z

## Translation

タイトル: AI エージェント ガードレールのベンチマークを構築したところ、嘘をついている独自のプラグインが検出されました
記事のタイトル: GitHub - Couldbeme/holdline: Holdline — AI エージェントの書き込み保護の中立ベンチマーク。キャッチ率、誤ブロック率、クラスバランスのとれたカッパ、およびインジェクション攻撃クラス。ガードを得点します。 · GitHub
説明: Holdline — AI エージェントの書き込みガードの中立的なベンチマーク。キャッチ率、誤ブロック率、クラスバランスのとれたカッパ、およびインジェクション攻撃クラス。ガードを得点します。 - 可能性があります/ホールドライン

記事本文:
GitHub - Couldbeme/holdline: Holdline — AI エージェントの書き込み保護の中立ベンチマーク。キャッチ率、誤ブロック率、クラスバランスのとれたカッパ、およびインジェクション攻撃クラス。ガードを得点します。 · GitHub
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
もしかしたら
/
ホールドライン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
12 コミット 12 コミット フォルダーとファイル
コーパス コーパス スクリプト スクリプト src src test test .gitignore .gitignore ライセンス

ライセンス README.md README.md RESULTS-ODCV.md RESULTS-ODCV.md RESULTS.md RESULTS.md ROADMAP.md ROADMAP.md odcv-compare.mjs odcv-compare.mjs odcv-run.mjs odcv-run.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml run.mjs run.mjs tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
警備員の仕事はラインを維持することです。ホールドラインは、AI エージェントの書き込み保護の中立的なベンチマークであるかどうかを測定します。それはあらゆるガードを得点します — (コミットメント、アクション) → ブロックとして表現されます?関数 — ラベル付きコーパスに対して、ゲートにとって重要なメトリック（キャッチ率、偽ブロック率、クラスバランスの取れたコーエンのカッパ（生のカッパはクラス不均衡の下にある））をレポートします。コーパスには、インジェクション攻撃クラス : アクションが含まれており、その内容は警備員に判決を言い渡そうとします。
理由: DeepSeek Harness エコシステムには 20 以上のガード/ポリシー プラグインがあり、それらのいずれかが機能するかどうかを測定する共有方法がありません。警備員の README にある「危険なコマンドをブロックする」という記述は証拠ではありません。このハーネスが証拠です。
pnpmインストール
node run.mjs # 42 件のコーパスに対するすべての組み込みガードをスコアリングします
node run.mjs --model < id > # ジャッジガードを別のローカルモデルに向けます
node odcv-run.mjs # 実際のエージェントの軌跡でジャッジを採点します (ODCV-Bench)
2 つの結果セット: RESULTS.md (作成されたコーパス、実際の名前付きガードとインジェクション クラスを含む) と RESULTS-ODCV.md (より難しい数値: 私たちが書いていない実際のエージェントの軌跡に関する 4 モデルの審査員パネルとの一致、バランスの取れた kappa 0.82 )。スコアラーがテストされます ( pnpm テスト )。これは、ジャッジ ガード用に公開された dsh-write-gate コアをドッグフードします。
Guard インターフェイスを src/guards.ts (name、kind、note、block(case)) に実装し、 run.mjs のリストに追加して、結果を含む PR を開きます。実際に公開されたプラグインをマウントするガード (

戦略の原型）は特に歓迎されます。
v0、RESULTS.md に記載されている正直な制限: コーパスは小さく、手作業で作成されており、拒否リストは戦略の原型 (特定のプラグインではない) であり、数値は 1 つのモデル / 1 つの実行です。値は公開される形状であり、誰でも再実行できます。マサチューセッツ工科大学。
Holdline — AI エージェントの書き込みガードの中立的なベンチマーク。キャッチ率、誤ブロック率、クラスバランスのとれたカッパ、およびインジェクション攻撃クラス。ガードを得点します。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Holdline — a neutral benchmark for AI-agent write-guards. Catch rate, false-block rate, class-balanced kappa, and an injection-attack class. Scores any guard. - couldbeme/holdline

GitHub - couldbeme/holdline: Holdline — a neutral benchmark for AI-agent write-guards. Catch rate, false-block rate, class-balanced kappa, and an injection-attack class. Scores any guard. · GitHub
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
couldbeme
/
holdline
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
corpus corpus scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md RESULTS-ODCV.md RESULTS-ODCV.md RESULTS.md RESULTS.md ROADMAP.md ROADMAP.md odcv-compare.mjs odcv-compare.mjs odcv-run.mjs odcv-run.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml run.mjs run.mjs tsconfig.json tsconfig.json View all files Repository files navigation
A guard's job is to hold the line. holdline measures whether it does — a neutral benchmark for AI-agent write-guards. It scores any guard — expressed as a (commitments, action) → block? function — over a labeled corpus, and reports the metrics that matter for a gate: catch rate, false-block rate, and class-balanced Cohen's kappa (raw kappa lies under class imbalance). The corpus includes an injection-attack class : actions whose content tries to talk the guard out of its verdict.
Why: the DeepSeek Harness ecosystem has 20+ guard/policy plugins and no shared way to measure whether any of them works. A guard's README saying "blocks dangerous commands" is not evidence. This harness is the evidence.
pnpm install
node run.mjs # scores every built-in guard over the 42-case corpus
node run.mjs --model < id > # point the judge guard at a different local model
node odcv-run.mjs # score the judge on REAL agent trajectories (ODCV-Bench)
Two result sets: RESULTS.md (authored corpus, incl. a real named guard and an injection class) and RESULTS-ODCV.md (the harder number: agreement with a 4-model judge panel on real agent trajectories we did not write, balanced kappa 0.82 ). Scorer is tested ( pnpm test ); it dogfoods the published dsh-write-gate core for the judge guard.
Implement the Guard interface in src/guards.ts (name, kind, note, block(case) ), add it to the list in run.mjs , and open a PR with your results. A guard that mounts an actual published plugin (rather than a strategy archetype) is especially welcome.
v0, honest limits stated in RESULTS.md : the corpus is small and hand-authored, the deny-list is a strategy archetype (not a specific plugin), and the numbers are one model / one run. The value is the shape it exposes and that anyone can re-run it. MIT.
Holdline — a neutral benchmark for AI-agent write-guards. Catch rate, false-block rate, class-balanced kappa, and an injection-attack class. Scores any guard.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
