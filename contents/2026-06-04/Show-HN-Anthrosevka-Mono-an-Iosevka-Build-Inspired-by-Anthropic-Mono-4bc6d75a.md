---
source: "https://github.com/nanxstats/anthrosevka"
hn_url: "https://news.ycombinator.com/item?id=48400820"
title: "Show HN: Anthrosevka Mono, an Iosevka Build Inspired by Anthropic Mono"
article_title: "GitHub - nanxstats/anthrosevka: An Iosevka custom build inspired by Anthropic Mono · GitHub"
author: "road2stat"
captured_at: "2026-06-04T18:55:26Z"
capture_tool: "hn-digest"
hn_id: 48400820
score: 1
comments: 0
posted_at: "2026-06-04T16:18:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Anthrosevka Mono, an Iosevka Build Inspired by Anthropic Mono

- HN: [48400820](https://news.ycombinator.com/item?id=48400820)
- Source: [github.com](https://github.com/nanxstats/anthrosevka)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T16:18:40Z

## Translation

タイトル: Show HN: Anthrosevka Mono、Anthropic Mono からインスピレーションを得た Iosevka ビルド
記事のタイトル: GitHub - nanxstats/anthrosevka: Anthropic Mono · GitHub に触発された Iosevka カスタム ビルド
説明: Anthropic Mono からインスピレーションを得た Iosevka カスタム ビルド - nanxstats/anthrosevka
HN テキスト: Anthropic Mono のルック アンド フィールが気に入ったのでこれを作成しましたが、Iosevka から構築された再現可能なオープン ソース コーディング/ターミナル フォントが欲しかったです。

記事本文:
GitHub - nanxstats/anthrosevka: Anthropic Mono からインスピレーションを得た Iosevka カスタム ビルド · GitHub
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
nanxstats
/
アントロセフカ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット Assets docs docs .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md private-build-plans.toml private-build-plans.toml release.sh release.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
アントロエフカ・モノはヨセフカです
Anthropic Mono のルック アンド フィールからインスピレーションを得たカスタム ビルド。
事前構築済みフォントは次の場所で利用できます。
GitHub がリリースします。
アイデアは、すぐに使用できるエクスペリエンスを向上させる、賢明で独自のデフォルトを出荷することです。
間隔を「ターミナル」に設定します。これにより、矢印などの特殊な記号が強制的に適合します。
厳密で狭い 1 列レイアウトであり、端末でのレンダリングの問題が修正されています。
デフォルトの重みは 400 と 700 です。
デフォルトでは幅 600。圧縮バージョンはありません。
高解像度ディスプレイを使用する価値があるため、ヒントなしの TTF。
ソースから Iosevka をビルドする手順をたどる
ビルド環境をセットアップします。
次に、private-build-plans.toml を Iosevka クローンのルートにコピーします。走る
npm run build -- ttf-unhinted::AnthrosevkaMono
免責事項
Anthrosevka Mono は、によって承認されていない独立したサードパーティ プロジェクトです。
Anthropic PBC と提携、または Anthropic PBC によってサポートされています。
Anthrosevka Mono は、SIL Open Font License、バージョン 1.1 に基づいてライセンスされています。
Anthropic Mono からインスピレーションを得た Iosevka カスタム ビルド
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
アントロエフカ モノ 0.1.0
最新の
2026 年 6 月 3 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An Iosevka custom build inspired by Anthropic Mono - nanxstats/anthrosevka

I made this because I like the look and feel of Anthropic Mono, but wanted a reproducible open source coding/terminal font built from Iosevka.

GitHub - nanxstats/anthrosevka: An Iosevka custom build inspired by Anthropic Mono · GitHub
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
nanxstats
/
anthrosevka
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits assets assets docs docs .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md private-build-plans.toml private-build-plans.toml release.sh release.sh View all files Repository files navigation
Anthrosevka Mono is an Iosevka
custom build inspired by the look and feel of Anthropic Mono.
Prebuilt fonts are available in the
GitHub releases .
The idea is to ship sensible, opinionated defaults for a better experience out of the box:
Set spacing to "terminal". This forces special symbols such as arrows to fit
a strict, narrow one-column layout and fixes rendering issues in terminals.
Weights 400 and 700 as the defaults.
Width 600 as the default; no condensed version.
Unhinted TTF, because you deserve to use a high-resolution display.
Follow building Iosevka from source
to set up the build environment.
Next, copy private-build-plans.toml to the root of the Iosevka clone. Run
npm run build -- ttf-unhinted::AnthrosevkaMono
Disclaimer
Anthrosevka Mono is an independent third-party project not endorsed by,
affiliated with, or supported by Anthropic PBC.
Anthrosevka Mono is licensed under the SIL Open Font License, Version 1.1 .
An Iosevka custom build inspired by Anthropic Mono
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Anthrosevka Mono 0.1.0
Latest
Jun 3, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
