---
source: "https://github.com/SakanaAI/fugu"
hn_url: "https://news.ycombinator.com/item?id=48797562"
title: "Fugu – A multi-agent LLM orchestrator delivered as a single API"
article_title: "GitHub - SakanaAI/fugu · GitHub"
author: "terminalchai"
captured_at: "2026-07-05T20:47:16Z"
capture_tool: "hn-digest"
hn_id: 48797562
score: 4
comments: 0
posted_at: "2026-07-05T20:13:11Z"
tags:
  - hacker-news
  - translated
---

# Fugu – A multi-agent LLM orchestrator delivered as a single API

- HN: [48797562](https://news.ycombinator.com/item?id=48797562)
- Source: [github.com](https://github.com/SakanaAI/fugu)
- Score: 4
- Comments: 0
- Posted: 2026-07-05T20:13:11Z

## Translation

タイトル: Fugu – 単一の API として提供されるマルチエージェント LLM オーケストレーター
記事タイトル：GitHub - SakenaAI/fugu · GitHub
説明: GitHub にアカウントを作成して、SakanaAI/fugu の開発に貢献します。

記事本文:
GitHub - SakenaAI/fugu · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
さかなあい
/
ふぐ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
86 コミット 86 コミット アセット/数値 アセット/数値 構成 構成 デモ デモ ドキュメント ドキュメント ノート ノート スクリプト スクリプト Fugu_technical_report.pdf Fugu_technical_report.pdf README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
さかなふぐは、1 つのモデルとして提供されるマルチエージェント システムです。 Fugu はフロンティア モデルを動的に調整して、複雑な複数ステップのタスクに取り組みます。チャット完了と応答エンドポイントの両方をサポートする、Sakana API を介して単一の LLM としてマルチエージェント システムにアクセスできます。
すぐに始めるために、単一のコマンドで Fugu を Codex にインストールできます。
カール -fsSL https://sakana.ai/fugu/install |バッシュ
次に、次のように起動します。
コーデックスフグ
追加のフラグとオプションについては、コマンド リファレンスを参照してください。 1 行インストールでは、Ubuntu と macOS がサポートされます。 Windows の場合、またはインストールが完了しない場合は、こちらのガイドに従ってください。
インテリジェントな調整による優れたパフォーマンス
Sakana Fugu は、強力なモデルの多様なプールを動的に調整およびオーケストレーションすることで、優れたパフォーマンスを実現します。評価の詳細についてはテクニカルレポートをご確認ください。
これらの結果は、2026 年 6 月の評価を反映しています。新しいフロンティア モデルがリリースされるたびに、Fugu のパフォーマンス上の優位性を維持するためにモデル プールを継続的に更新し、コーディネーターを再トレーニングします。
これらの例では、さかなふぐモデルを 3 つのフロンティア ベースライン、Gemini 3.1 Pro (高)、Opus 4.8 (最大)、および GPT 5.5 (xhigh) と比較しています。ブランドごとの帰属ではなく行動に重点を置くため、各説明ではベースラインがモデル A、モデル B、およびモデル C として匿名化されています。マッピングはサンプル間で意図的に固定されていません。
『さかなふぐ』は、ICLR 2026 に掲載された 2 つの論文に基づいています。
公開以来、いくつかの機能強化を行ってきました。の

完全な技術レポートはここから入手できます。
Sakana Fuguの使用中に問題やバグが発生した場合は、https://fugu.sakana.ai までお問い合わせください。
研究にさかなふぐを使用する場合は、当社の技術レポートを引用してください。
@misc { fugu2026sakana ,
title = { さかなふぐテクニカルレポート } ,
author = { {ふぐチーム、さかなAI} } ,
年 = { 2026 } 、
eprint = { 2606.21228 } 、
archivePrefix = { arXiv } 、
プライマリクラス = { cs.LG } 、
URL = { https://arxiv.org/abs/2606.21228 } 、
}
について
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
89
フォーク
レポートリポジトリ
リリース
2
ユーザーからインスピレーションを得たデモ
最新の
2026 年 6 月 26 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to SakanaAI/fugu development by creating an account on GitHub.

GitHub - SakanaAI/fugu · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
SakanaAI
/
fugu
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
86 Commits 86 Commits assets/ figures assets/ figures configs configs demos demos docs docs notes notes scripts scripts Fugu_technical_report.pdf Fugu_technical_report.pdf README.md README.md View all files Repository files navigation
Sakana Fugu is a multi-agent system delivered as one model. Fugu dynamically orchestrates frontier models to tackle complex, multi-step tasks. You can access the multi-agent system as a single LLM through the Sakana API , which supports both Chat Completions and Responses endpoints.
To quickly get started, you can install Fugu into Codex with a single command:
curl -fsSL https://sakana.ai/fugu/install | bash
Then launch it with:
codex-fugu
See the command reference for additional flags and options. The one-line install supports Ubuntu and macOS. On Windows, or if the install does not complete, follow the guide here .
Superior performance via intelligent coordination
Sakana Fugu achieves superior performance by dynamically coordinating and orchestrating a diverse pool of powerful models. For evaluation details, check our technical report .
These results reflect our June 2026 evaluation. As new frontier models are released, we continuously update our model pool and retrain our coordinators to maintain Fugu's performance advantage.
These examples compare Sakana Fugu models with three frontier baselines: Gemini 3.1 Pro (high), Opus 4.8 (max), and GPT 5.5 (xhigh). To keep the focus on behavior rather than brand-by-brand attribution, the baselines are anonymized as Model A, Model B, and Model C in each description. The mapping is intentionally not fixed across examples.
Sakana Fugu is based on two papers published in ICLR 2026.
Since publication, we have made several enhancements. The full technical report is available here .
Please contact us at https://fugu.sakana.ai for issues or bugs while using Sakana Fugu.
If you use Sakana Fugu in your research, please cite our technical report:
@misc { fugu2026sakana ,
title = { Sakana Fugu Technical Report } ,
author = { {Fugu Team, Sakana AI} } ,
year = { 2026 } ,
eprint = { 2606.21228 } ,
archivePrefix = { arXiv } ,
primaryClass = { cs.LG } ,
url = { https://arxiv.org/abs/2606.21228 } ,
}
About
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
89
forks
Report repository
Releases
2
User-inspired demos
Latest
Jun 26, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
