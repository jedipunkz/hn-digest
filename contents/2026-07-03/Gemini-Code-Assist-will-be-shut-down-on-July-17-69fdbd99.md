---
source: "https://docs.cloud.google.com/gemini/docs/code-review/review-repo-code"
hn_url: "https://news.ycombinator.com/item?id=48774429"
title: "Gemini Code Assist will be shut down on July 17"
article_title: "Review GitHub code using Gemini Code Assist | Gemini for Google Cloud | Google Cloud Documentation"
author: "ushakov"
captured_at: "2026-07-03T13:43:39Z"
capture_tool: "hn-digest"
hn_id: 48774429
score: 2
comments: 0
posted_at: "2026-07-03T12:52:48Z"
tags:
  - hacker-news
  - translated
---

# Gemini Code Assist will be shut down on July 17

- HN: [48774429](https://news.ycombinator.com/item?id=48774429)
- Source: [docs.cloud.google.com](https://docs.cloud.google.com/gemini/docs/code-review/review-repo-code)
- Score: 2
- Comments: 0
- Posted: 2026-07-03T12:52:48Z

## Translation

タイトル: Gemini Code Assist は 7 月 17 日に終了します
記事のタイトル: Gemini Code Assist を使用して GitHub コードをレビューする | Google Cloud 向け Gemini | Google Cloud ドキュメント
説明: Gemini Code Assist を使用してプル リクエストをレビューします。

記事本文:
Gemini Code Assist を使用して GitHub コードをレビューする | Google Cloud 向け Gemini | Google Cloud ドキュメント
メインコンテンツにスキップ
技術分野
閉じる
AI と ML
分散型、ハイブリッド、マルチクラウド
アクセスとリソースの管理
SDK、言語、フレームワーク、ツール
製品の提供
ジェミニコードアシスト
概要
Gemini コードアシストを備えた Gemini 3
サポートされている言語、IDE、およびインターフェイス
リリースノート
Gemini for Google Cloud リリースノート
Gemini Code Assist リリースノート
Gemini for Google Cloud 管理設定の概要
Google Cloud プロダクトに対して Gemini をオフにする
Gemini コードアシストの構成
Gemini Code Assist リリース チャネルの構成
Gemini Code Assist の使用からファイルを除外する
ローカルコードベース認識を構成する
Gemini Code Assist のログ記録を構成する
Gemini Code Assist サブスクリプションを追加または変更する
Gemini コードアシスト ライセンス
Gemini Code Assist ライセンスをリクエストする
Gemini Code Assist ライセンスの管理
組織を越えたライセンスの使用を防止する
Gemini Code Assist へのアクセスを個人向けに制御する
Gemini Code Assist for VS Code のプレリリース機能を使用する
VPC サービス コントロールの構成
ユーザードメイン制限によるネットワークアクセスの制御
Gemini Code Assist を使用したコード作成
コード機能の概要
Gemini Code Assist とチャットする
チャット機能の概要
Gemini Code Assist チャットを使用する
Gemini Code Assist エージェント モードを使用する
コードのカスタマイズ
コードのカスタマイズの概要
顧客管理の暗号化キーを使用してデータを暗号化する
Gemini Code Assist を使用して GitHub コードをレビューする
GitHub 用の Gemini Code Assist をセットアップする
GitHub に Gemini コード アシストを使用する
GitHub で Gemini Code Assist の動作をカスタマイズする
Gemini Code Assist の使用状況を監視する
Gemini for Google Cloud のログを表示する
Google Cloud 監査ロギング用の Gemini
ビジネス AI コードの監査ログ
Gemini Code Assist 機能へのアクセスのトラブルシューティング
分散型、ハイブリッド、マルチクラウド
あ

アクセスとリソースの管理
SDK、言語、フレームワーク、ツール
Gemini Code Assist を使用して GitHub コードをレビューする
コレクションを整理整頓
好みに基づいてコンテンツを保存し、分類します。
GitHub 上の Gemini Code Assist は、次の機能を提供します。
Gemini は、コードレビューアとして機能してプルリクエストプロセスに参加します。
GitHub 上の Gemini Code Assist は、
プルを自動的に要約する Gemini 搭載エージェント
詳細なコードレビューを要求して提供し、レビューを迅速化し、
コードの品質を向上させます。
GitHub で Gemini Code Assist をセットアップしたら、
Gemini Code Assist はプルのどの段階でも呼び出すことができます
コードのレビューをリクエストします。と対話できます
Gemini Code Assist は、プル リクエストに直接コメントします。
レビューに関して明確な質問をすることで、
Gemini コードアシストが作成します。
/gemini タグを追加して Gemini コード アシストを要求する
プル リクエストのコンテキストで質問するには、コメントを入力してください。
Gemini Code Assist は役立つ情報を自動的に取得します。
リポジトリからの情報とそのタスクを実行するためのプル リクエスト。
このドキュメントは、あらゆるスキル レベルの開発者を対象としています。それは次のことを前提としています
GitHub に関する実用的な知識があること。
コンシューマ版とエンタープライズ版
GitHub の Gemini Code Assist は企業で利用可能です
このバージョンは、Google Cloud を通じてインストールします。コンシューマ版も
存在します。ただし、コンシューマ バージョンへのリクエストの処理は中止されています。
このバージョンはインストールしないでください。
次の表は、コンシューマ バージョンとコンシューマ バージョンの違いをまとめたものです。
エンタープライズ版:
設定ページを使用してアカウントに関連付けられたすべてのリポジトリ
Google Cloud を使用した複数のリポジトリ間
Google Cloud を使用した複数のリポジトリ間
GitHub 上の Gemini Code Assist が生成されない
概要またはコードの提案

内にあるファイルのイオン
.github/workflows ディレクトリ。この除外は侵入を防ぐのに役立ちます
安全でない可能性のある構成をリポジトリに追加します。
エンタープライズ バージョンでは、
開発者接続接続
GitHub リポジトリを Google Cloud に接続します。
この Developer Connect 接続は常にリージョン内に作成されます。
us-east1 。
この Developer Connect 接続は、
Code Assist のソース コード管理セクションは次のとおりです。
Gemini Code Assist エージェントとツール。手順については、
GitHub の Gemini Code Assist のセットアップを参照してください。
GitHub で Gemini Code Assist をセットアップします。
GitHub の Gemini Code Assist を使用します。
方法を学ぶ
Gemini Code Assist on GitHub の動作をカスタマイズします。
特に明記されていない限り、このページのコンテンツはクリエイティブ コモンズ表示 4.0 ライセンスに基づいてライセンスされており、コード サンプルは Apache 2.0 ライセンスに基づいてライセンスされています。詳細については、Google デベロッパー サイト ポリシーを参照してください。 Java は、Oracle および/またはその関連会社の登録商標です。
Google Cloud の使用を開始する
気候変動対策の 30 年目: 参加しましょう
Google Cloud ニュースレターに登録する
購読する

## Original Extract

Review pull requests using Gemini Code Assist.

Review GitHub code using Gemini Code Assist | Gemini for Google Cloud | Google Cloud Documentation
Skip to main content
Technology areas
close
AI and ML
Distributed, hybrid, and multicloud
Access and resources management
SDK, languages, frameworks, and tools
Product offerings
Gemini Code Assist
Overview
Gemini 3 with Gemini Code Assist
Supported languages, IDEs, and interfaces
Release notes
Gemini for Google Cloud release notes
Gemini Code Assist release notes
Gemini for Google Cloud admin settings overview
Turn off Gemini for Google Cloud products
Configure Gemini Code Assist
Configure Gemini Code Assist release channels
Exclude files from Gemini Code Assist use
Configure local codebase awareness
Configure Gemini Code Assist logging
Add or change Gemini Code Assist subscriptions
Gemini Code Assist licenses
Request a Gemini Code Assist license
Manage Gemini Code Assist licenses
Prevent cross-organization license usage
Control access to Gemini Code Assist for individuals
Use pre-release features in Gemini Code Assist for VS Code
Configure VPC Service Controls
Control Network Access with User Domain Restrictions
Code with Gemini Code Assist
Code features overview
Chat with Gemini Code Assist
Chat features overview
Use the Gemini Code Assist chat
Use the Gemini Code Assist agent mode
Code customization
Code customization overview
Encrypt data with customer-managed encryption keys
Review GitHub code with Gemini Code Assist
Set up Gemini Code Assist for GitHub
Use Gemini Code Assist for GitHub
Customize Gemini Code Assist behavior in GitHub
Monitor Gemini Code Assist usage
View Gemini for Google Cloud logs
Gemini for Google Cloud audit logging
Business AI Code audit logging
Troubleshoot access to Gemini Code Assist features
Distributed, hybrid, and multicloud
Access and resources management
SDK, languages, frameworks, and tools
Review GitHub code using Gemini Code Assist
Stay organized with collections
Save and categorize content based on your preferences.
Gemini Code Assist on GitHub brings the power of
Gemini to the pull request process by acting as a code reviewer.
Gemini Code Assist on GitHub uses a
Gemini-powered agent that automatically summarizes pull
requests and provides in-depth code reviews, speeding up reviews and
increasing the quality of code.
Once you've set up Gemini Code Assist on GitHub ,
you can invoke Gemini Code Assist at any stage of the pull
request to review the code. You can interact with
Gemini Code Assist in the pull request comments directly by:
Asking clarifying questions on the review that
Gemini Code Assist creates.
Prompting Gemini Code Assist by adding the /gemini tag to
your comments to ask questions in the context of the pull request.
Gemini Code Assist will automatically retrieve helpful
information from the repository and pull request to perform its tasks.
This document is intended for developers of all skill levels. It assumes that
you have a working knowledge of GitHub.
Consumer version and enterprise version
Gemini Code Assist on GitHub is available in an enterprise
version, which you install through Google Cloud. A consumer version also
exists; however, serving requests to the consumer version is being discontinued,
and you shouldn't install this version.
The following table summarizes the differences between the consumer version and
enterprise version:
All repositories associated with an account using the settings page
Across multiple repositories using Google Cloud
Across multiple repositories using Google Cloud
Gemini Code Assist on GitHub does not generate
summaries or code suggestions for any files located within the
.github/workflows directory. This exclusion helps prevent the introduction
of potentially insecure configurations to the repository.
The enterprise version uses a
Developer Connect connection
to connect your GitHub repositories to Google Cloud.
This Developer Connect connection is always created in the region
us-east1 .
This Developer Connect connection must be created using the
Code Assist Source Code Management section found in
Gemini Code Assist Agents & Tools . For instructions,
see Set up Gemini Code Assist on GitHub .
Set up Gemini Code Assist on GitHub .
Use Gemini Code Assist on GitHub .
Learn how to
customize Gemini Code Assist on GitHub behavior .
Except as otherwise noted, the content of this page is licensed under the Creative Commons Attribution 4.0 License , and code samples are licensed under the Apache 2.0 License . For details, see the Google Developers Site Policies . Java is a registered trademark of Oracle and/or its affiliates.
Getting Started with Google Cloud
Our third decade of climate action: join us
Sign up for the Google Cloud newsletter
Subscribe
