---
source: "https://github.com/joey-io/gauge"
hn_url: "https://news.ycombinator.com/item?id=49149111"
title: "Gauge – see where your Claude Code subscription goes"
article_title: "GitHub - joey-io/gauge · GitHub"
author: "joey-io"
captured_at: "2026-08-02T22:45:35Z"
capture_tool: "hn-digest"
hn_id: 49149111
score: 1
comments: 0
posted_at: "2026-08-02T22:34:01Z"
tags:
  - hacker-news
  - translated
---

# Gauge – see where your Claude Code subscription goes

- HN: [49149111](https://news.ycombinator.com/item?id=49149111)
- Source: [github.com](https://github.com/joey-io/gauge)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T22:34:01Z

## Translation

タイトル: ゲージ – クロード コードのサブスクリプションがどこに行くのかを確認します
記事タイトル: GitHub - joey-io/gauge · GitHub
説明: GitHub でアカウントを作成して、joey-io/gauge の開発に貢献します。

記事本文:
GitHub - joey-io/gauge · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ジョーイ・イオ
/
ゲージ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらにアクトを開く

イオンメニュー フォルダーとファイル
8 コミット 8 コミット bin bin lib lib site site site .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コード プランのローカル メーター。月額定額料金を支払います。ゲージ
5 時間枠、日、モデル、プロジェクトごとに、それがどこに向かうのかを示します。
npx github:joey-io/gauge
これでインストールは完了です。セッショントランスクリプトのクロードコードをすでに読み取ります
~/.claude/projects に書き込み、実際の API メッセージまで重複を排除します。
1 つの画面をレンダリングします。
これまでの最大の 5 時間枠に対する過去 5 時間の支出
過去 14 日間の日足
どのモデルとどのプロジェクトが許容量を消費しているか
マシンからは何も残りません。アカウント、テレメトリ、ネットワーク通話はありません。
ドルは API リストに相当します - この使用量にかかる費用
Anthropic が公開している API レート。サブスクリプションプランは割り当てを公開しません
したがって、これが唯一の正直な共通通貨です。それはあなたの請求書ではありません。
クロード コードは、ストリーミング チャンクごとに 1 つのトランスクリプト行を書き込み、すべてに
同じメッセージ ID を使用し、セッションが再開されるとメッセージを新しいファイルに再生します。
メッセージ ID によって重複排除を測定します。単純な行合計は大幅にカウントしすぎます。
サブエージェントのトランスクリプトは、いくつかのディレクトリの深さにネストされます。ゲージが彼らを歩かせる —
テストでは、彼らは支出の約半分を保持していました。
「ピーク 5 時間ウィンドウ」は、観測された天井であり、そのようにラベル付けされています。人間主義はそうではありません
実際のプランの制限を明らかにします。
ノード 22.5 以降 (組み込みの sqlite を使用 - npm 依存関係なし)。
~/.claude/projects にあるクロード コードの履歴。
ゲージ --days 30 # モデル別/プロジェクト別のセクションを拡大します
ゲージ activate <key> # ライセンス キーを登録します (オフラインで検証します)
GAUGE_DATA_DIR=... # キャッシュ データベースを移動します (デフォルト ~/.cache/gauge)
最初の実行ではスキャンします

あなたの全履歴 (トランスクリプト 1 GB あたり約 1 分)。
その後は段階的に実行され、1 ～ 2 秒かかります。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to joey-io/gauge development by creating an account on GitHub.

GitHub - joey-io/gauge · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
joey-io
/
gauge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits bin bin lib lib site site .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
A local meter for your Claude Code plan. You pay a flat monthly price; gauge
shows you where it goes — by 5-hour window, day, model, and project.
npx github:joey-io/gauge
That's the whole install. It reads the session transcripts Claude Code already
writes to ~/.claude/projects , dedupes them down to real API messages, and
renders one screen:
your last 5 hours of spend against the biggest 5-hour window you've ever had
daily bars for the last 14 days
which models and which projects are eating the allowance
Nothing leaves your machine. No account, no telemetry, no network calls.
Dollars are API-list-equivalent — what this usage would cost at
Anthropic's published API rates. Subscription plans don't publish their quota
mechanics, so this is the only honest common currency. It is not your bill.
Claude Code writes one transcript line per streamed chunk, all carrying the
same message id, and replays messages into new files when sessions resume.
gauge dedupes by message id; naive line-summing over-counts badly.
Subagent transcripts nest several directories deep. gauge walks them —
in testing they held roughly half the spend.
"Peak 5h window" is your observed ceiling, labeled as such. Anthropic doesn't
expose the actual plan limit.
Node 22.5 or newer (uses the built-in sqlite — zero npm dependencies).
A Claude Code history at ~/.claude/projects .
gauge --days 30 # widen the by-model / by-project sections
gauge activate <key> # register a license key (verifies offline)
GAUGE_DATA_DIR=... # move the cache db (default ~/.cache/gauge)
The first run scans your full history (about a minute per GB of transcripts).
After that it's incremental and takes a second or two.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
