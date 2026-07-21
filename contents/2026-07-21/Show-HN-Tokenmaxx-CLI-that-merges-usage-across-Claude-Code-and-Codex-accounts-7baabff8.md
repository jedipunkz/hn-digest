---
source: "https://github.com/RubricLab/tokenmaxx"
hn_url: "https://news.ycombinator.com/item?id=48995775"
title: "Show HN: Tokenmaxx – CLI that merges usage across Claude Code and Codex accounts"
article_title: "GitHub - RubricLab/tokenmaxx: A local proxy that aggregates usage across your Codex and Claude Code accounts — with live token-throughput analytics. · GitHub"
author: "sarimmalik"
captured_at: "2026-07-21T18:10:10Z"
capture_tool: "hn-digest"
hn_id: 48995775
score: 4
comments: 0
posted_at: "2026-07-21T17:56:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tokenmaxx – CLI that merges usage across Claude Code and Codex accounts

- HN: [48995775](https://news.ycombinator.com/item?id=48995775)
- Source: [github.com](https://github.com/RubricLab/tokenmaxx)
- Score: 4
- Comments: 0
- Posted: 2026-07-21T17:56:28Z

## Translation

タイトル: Show HN: Tokenmaxx – Claude Code アカウントと Codex アカウント間の使用法をマージする CLI
記事のタイトル: GitHub - RubricLab/tokenmaxx: ライブ トークン スループット分析を使用して、Codex アカウントと Claude Code アカウント全体の使用状況を集約するローカル プロキシ。 · GitHub
説明: ライブ トークン スループット分析を使用して、Codex アカウントと Claude Code アカウント全体の使用状況を集約するローカル プロキシ。 - RubricLab/tokenmaxx

記事本文:
GitHub - RubricLab/tokenmaxx: ライブ トークン スループット分析を使用して、Codex アカウントと Claude Code アカウント全体の使用状況を集約するローカル プロキシ。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。 P

このページをリロードしてください。
ルーブリックラボ
/
トークンマックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
55 コミット 55 コミット .github/ workflows .github/ workflows bin bin media media src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md biome.json biome.json bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すべての Codex アカウントと Claude Code アカウントに 1 つのダッシュボードが対応します。それらを簡単に切り替えて、使用状況を監視します。
macOS · Bun · Rubric Labs プロジェクト · OpenAI や Anthropic とは無関係
bun add -g tokenmaxx
tokenmaxx # ダッシュボードを開始します
何をするのか
複数の Codex または Claude アカウントを使用して、一連のコーディング エージェントを実行します。
tokenmaxx では、それらすべてがローカルでサインインされた状態に保たれ、クライアントがどれを使用するかを選択できるようになります。マシン上の小さなプロキシがアクティブなアカウントの資格情報を各リクエストに添付するため、切り替えは次のリクエスト (ターンの途中であっても) で有効になります。認証情報は macOS キーチェーン内に存在し、認証情報を発行したプロバイダー以外の場所に送信されることはありません。
tokenmaxx を実行します。 [アカウント] には、すべてのアカウントとそのライブ レート制限ウィンドウが圧力によって色分けされ、プラン階層とリセット カウントダウンがインラインで表示されます。行は圧力によって並べ替えられます。 ● は現在交通が進んでいる場所を示します。
分析は、すべてのアカウントと両方のプロバイダーの合計トークン スループットと、API リスト レートでのその使用量のコストを合わせたものです。トークンは応答のストリームとして測定され、バッファリングされることはありません。そのため、すべての数値をクライアント自身のセッション ログと照合できます。 m を押すと、モデルごとの詳細な価格の内訳が表示されます。
設定はプロバイダーごとにマスターのオン/オフを保持します。

次に、自動回転、スイッチのしきい値、クールダウンがライブで適用されます。
これをオンにすると、tokenmaxx はアクティブなアカウントのレート制限ウィンドウを監視します。最大のアカウントがしきい値を超えると、アカウントの中で最も余裕のあるアカウントに切り替わります。デフォルトのしきい値は 90% で、後で必要な場合に備えて、各ウィンドウの最後の部分をそのまま残します。リクエストの途中でアカウントがハードリミットに達した場合、プロキシは余裕を持って次のアカウントでそのリクエストを再試行します。
--threshold 90 の両方で tokenmaxx auto # または: codex | codex |クロード …オフ
仕組み
127.0.0.1:8459 上の単一のループバック プロキシと、すでに使用しているクライアント。
リクエストごとのアカウント。プロキシは、リクエストごとにどのアカウントがアクティブであるかを読み取り、その資格情報を添付します。次のリクエストで土地を切り替えます。
無料で読める圧力。どちらのプロバイダーも、すべての応答でレート制限の状態を報告します。プロキシはトラフィック ストリームとしてそれを読み取るため、アクティブなアカウントがどの程度満たされているかを常に把握しており、追加のリクエストはありません。
公式アプリのみ。あなたのサブスクリプション ログインは、Claude Code および Codex 用です。カスタムのものを構築している場合は、プロバイダーからの API キーを使用します。 tokenmaxx はサブスクリプションを API プランに変換しません。
tokenmaxx ライブ ダッシュボード
tokenmaxx ログイン <codex|claude> サインイン;孤立した冪等
tokenmaxx インストール ルート ネイティブ コーデックスとクロード
tokenmaxx アンインストール ネイティブ設定を復元
tokenmaxx スイッチ <codex|claude> <email> アカウントをアクティブにする
tokenmaxx auto <両方|コーデックス|クロード> <オン|オフ> [--しきい値 N]
tokenmaxx リスト |ステータス |リフレッシュ |医者
環境: TOKENMAXX_HOME 、 TOKENMAXX_PROXY_PORT 、 TOKENMAXX_THEME 。
tokenmaxx は、自分で支払うアカウントを持つ 1 人用です。アカウントの制限が大きくなったり、制限がバイパスされたりすることはなく、資格情報はキーチェーンとプロバイダーの間に留まります。アカウントを共有しないでください

それらをプールしたり、アクセスを転売したりしないでください。プロバイダーの規約は変更されるため、この種の切り替えが許可されているかどうかを確認するのはユーザーの責任です。ソフトウェアは現状のまま提供され、保証はありません。
公式ではなく、独立した Rubric Labs プロジェクト
OpenAI または Anthropic の製品、OpenAI または Anthropic と提携、または承認された製品。インスピレーションを得た
コーデックスアカウントスイッチャー 。
ライブ トークン スループット分析を使用して、Codex アカウントと Claude Code アカウント全体の使用状況を集約するローカル プロキシ。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A local proxy that aggregates usage across your Codex and Claude Code accounts — with live token-throughput analytics. - RubricLab/tokenmaxx

GitHub - RubricLab/tokenmaxx: A local proxy that aggregates usage across your Codex and Claude Code accounts — with live token-throughput analytics. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
RubricLab
/
tokenmaxx
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
55 Commits 55 Commits .github/ workflows .github/ workflows bin bin media media src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md biome.json biome.json bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
One dashboard for all of your Codex and Claude Code accounts. Easily switch between them, and monitor usage.
macOS · Bun · a Rubric Labs project · not affiliated with OpenAI or Anthropic
bun add -g tokenmaxx
tokenmaxx # starts the dashboard
What it does
You run a fleet of coding agents using multiple Codex or Claude accounts:
tokenmaxx keeps them all signed in locally and lets you choose which one your clients use. A small proxy on your machine attaches the active account's credential to each request, so a switch takes effect on the very next request, even mid-turn. Your credentials live in the macOS Keychain and never go anywhere except to the provider that issued them.
Run tokenmaxx . Accounts shows every account and its live rate-limit windows, colored by pressure, with plan tier and reset countdowns inline. Rows sort by pressure; the ● marks where traffic is going right now.
Analytics is combined token throughput across all accounts and both providers, with the ≈ cost of that usage at API list rates. Tokens are metered as responses stream by, never buffered, so every number is cross-checkable against your clients' own session logs. Press m for the full pricing breakdown per model.
Settings holds the master on/off per provider, then auto-rotation, the switch threshold, and cooldown, applied live.
Turn it on and tokenmaxx watches the active account's rate-limit windows. When the fullest one crosses your threshold, it switches to whichever of your accounts has the most room. The default threshold is 90%, which leaves the last stretch of every window alone in case you want it later. If an account hits a hard limit in the middle of a request, the proxy retries that request on your next account with room.
tokenmaxx auto both on --threshold 90 # or: codex | claude … off
How it works
A single loopback proxy on 127.0.0.1:8459 , and the clients you already use.
Account per request. The proxy reads which account is active for each request and attaches its credential. Switching lands on the next request.
Pressure read for free. Both providers report rate-limit state on every response; the proxy reads it as traffic streams by, so it always knows how full the active account is, with zero extra requests.
Official apps only. Your subscription login is for Claude Code and Codex. If you're building something custom, use an API key from the provider. tokenmaxx doesn't turn a subscription into an API plan.
tokenmaxx live dashboard
tokenmaxx login <codex|claude> sign in; isolated, idempotent
tokenmaxx install route native codex & claude
tokenmaxx uninstall restore native config
tokenmaxx switch <codex|claude> <email> make an account active
tokenmaxx auto <both|codex|claude> <on|off> [--threshold N]
tokenmaxx list | status | refresh | doctor
Env: TOKENMAXX_HOME , TOKENMAXX_PROXY_PORT , TOKENMAXX_THEME .
tokenmaxx is for one person with accounts they pay for themselves. No account gets bigger limits, no limit gets bypassed, and your credentials stay between your Keychain and the provider. Don't share accounts, don't pool them, don't resell access. Provider terms change, and it's on you to check that yours allow this kind of switching. The software is provided as is, with no warranty.
An independent Rubric Labs project, not an official
product of, affiliated with, or endorsed by OpenAI or Anthropic. Inspired by
codex-account-switcher .
A local proxy that aggregates usage across your Codex and Claude Code accounts — with live token-throughput analytics.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
