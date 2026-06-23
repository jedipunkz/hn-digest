---
source: "https://docs.cloud.google.com/recaptcha/docs/hand-gesture-verification"
hn_url: "https://news.ycombinator.com/item?id=48649795"
title: "Hand Gesture Verification – Google Cloud Fraud Defense"
article_title: "Hand gesture verification | Google Cloud Fraud Defense | Google Cloud Documentation"
author: "josephcsible"
captured_at: "2026-06-23T19:35:54Z"
capture_tool: "hn-digest"
hn_id: 48649795
score: 2
comments: 2
posted_at: "2026-06-23T19:05:32Z"
tags:
  - hacker-news
  - translated
---

# Hand Gesture Verification – Google Cloud Fraud Defense

- HN: [48649795](https://news.ycombinator.com/item?id=48649795)
- Source: [docs.cloud.google.com](https://docs.cloud.google.com/recaptcha/docs/hand-gesture-verification)
- Score: 2
- Comments: 2
- Posted: 2026-06-23T19:05:32Z

## Translation

タイトル: 手のジェスチャー検証 – Google Cloud の不正行為防御
記事タイトル: ハンドジェスチャー認証 | Google Cloud の不正行為防御 | Google Cloud ドキュメント

記事本文:
ハンドジェスチャー認証 | Google Cloud の不正行為防御 | Google Cloud ドキュメント
メインコンテンツにスキップ
技術分野
閉じる
AI と ML
分散型、ハイブリッド、マルチクラウド
アクセスとリソースの管理
SDK、言語、フレームワーク、ツール
デモ Web サイトで reCAPTCHA をテストする
reCAPTCHA がプライベート アクセス トークンをどのように使用するかを理解する
Web サイトで reCAPTCHA を設定する
Web サイトのセットアップの概要
Web ページにスコアベースのキーをインストールする
Web ページにチェックボックス キーをインストールする
ポリシーベースのチャレンジ キーを Web ページにインストールする
Web サイトの評価を作成する
Web サイトの評価を解釈する
IP アドレス許可リストを構成する
モバイル アプリケーションで reCAPTCHA を設定する
モバイル アプリケーションのセットアップの概要
モバイルアプリケーション用のキーを作成する
モバイル SDK の非推奨およびシャットダウンのポリシー
iOS
Apple のプライバシーの詳細を確認する
reCAPTCHA を iOS アプリと統合する
アンドロイド
Google Play のデータ開示要件に備える
reCAPTCHA を Android アプリと統合する
モバイルアプリケーションの評価を作成する
モバイルアプリケーションの評価を解釈する
Google Cloud Armor との統合
Google Cloud Armor のセットアップの概要
Google Cloud Armor との統合のための機能
Google Cloud Armor の reCAPTCHA トークン属性
ウェブサイト用の Google Cloud Armor との統合
モバイル アプリケーション向けに Google Cloud Armor と統合する
Google Cloud Armor との統合例
API エンドポイントの reCAPTCHA Express をセットアップする
ユーザーアカウント
ユーザーアカウント保護機能
多要素認証を構成する
パスワードの漏洩と認証情報の侵害を検出する
Web サイト上の不正行為を検出する
Web サイト上で関連する動作を示すユーザー アカウントを特定する
モバイルアプリケーションでの不正行為を検出
アカウント乗っ取りを検出して防止する
攻撃を特定して調査する
自動化された脅威
ベストプラクティス

自動化された脅威からの保護のため
支払い取引
トランザクション防御を設定する
トランザクション データで評価に注釈を付ける
reCAPTCHA キーの操作を制限する
reCAPTCHA Classic からの移行
移行後に不正防止機能を使用する
Looker Studio でダッシュボードを作成する
reCAPTCHA 統合の問題のトラブルシューティング
すべての製品のすべてのコード サンプル
分散型、ハイブリッド、マルチクラウド
アクセスとリソースの管理
SDK、言語、フレームワーク、ツール
ハンドジェスチャー認証
コレクションを整理整頓
好みに基づいてコンテンツを保存し、分類します。
このドキュメントでは、データ収集、カメラの許可、
手のジェスチャー認証のためのアクセシビリティ。
ハンド ジェスチャ機能が有効になっている場合、reCAPTCHA は次のデータを収集します。
Google は、ユーザーがさまざまな動作をしているときの手を撮影した 1 つ以上のビデオを分析します。
動作やジェスチャー。ビデオを処理して手のランドマークデータを抽出し、
これには以下が含まれます
21 ハンドナックル座標。
動画はユーザーの身元と関連付けられることはなく、ユーザーの身元が判明した後に削除されます。
検証プロセス。音声は録音されません。
Google は、ユーザーの手のジェスチャーの画像や動画を保存しません。
検証プロセスに使用したり、他の目的にデータを使用したりすることはできません。動画や
画像はチャレンジが完了すると自動的に削除されます。
Google が収集する情報は、次の規則に従って使用および保存されます。
Google プライバシー ポリシー 。
ハンド ジェスチャー チャレンジでは、ユーザーのカメラにアクセスする許可が必要です。その後
ユーザーはハンドジェスチャーを実行することに同意し、Google はカメラの許可を受け取ります。
ユーザーはブラウザの設定でいつでもこれらの権限を管理できます。 Google
セキュリティ検証のみを目的としてハンドジェスチャービデオを処理し、
関連データや権限を第三者に転送しません

パーティー。
アクセシビリティが必要で、ハンドジェスチャーを使用して完了できないユーザー向け
reCAPTCHA は引き続きビジュアルとオーディオを提供します。
課題を解決し、よりアクセスしやすく安全な代替手段を開発します。
特に明記されていない限り、このページのコンテンツはクリエイティブ コモンズ表示 4.0 ライセンスに基づいてライセンスされており、コード サンプルは Apache 2.0 ライセンスに基づいてライセンスされています。詳細については、Google デベロッパー サイト ポリシーを参照してください。 Java は、Oracle および/またはその関連会社の登録商標です。
Google Cloud の使用を開始する
気候変動対策の 30 年間: 参加しましょう
Google Cloud ニュースレターに登録する
購読する

## Original Extract

Hand gesture verification | Google Cloud Fraud Defense | Google Cloud Documentation
Skip to main content
Technology areas
close
AI and ML
Distributed, hybrid, and multicloud
Access and resources management
SDK, languages, frameworks, and tools
Test reCAPTCHA in a demo website
Understand how reCAPTCHA uses Private Access Tokens
Set up reCAPTCHA on websites
Setup overview for websites
Install score-based keys on web pages
Install checkbox keys on web pages
Install policy-based challenge keys on web pages
Create assessments for websites
Interpret assessments for websites
Configure an IP address allowlist
Set up reCAPTCHA on mobile applications
Setup overview for mobile applications
Create keys for mobile applications
Deprecation and shutdown policy for mobile SDKs
iOS
Review Apple privacy details
Integrate reCAPTCHA with iOS apps
Android
Prepare for Google Play data disclosure requirements
Integrate reCAPTCHA with Android apps
Create assessments for mobile applications
Interpret assessments for mobile applications
Integration with Google Cloud Armor
Setup overview for Google Cloud Armor
Features for integration with Google Cloud Armor
reCAPTCHA token attributes for Google Cloud Armor
Integrate with Google Cloud Armor for websites
Integrate with Google Cloud Armor for mobile applications
Examples of integration with Google Cloud Armor
Set up reCAPTCHA express for API endpoints
User accounts
User accounts protection features
Configure Multi-factor authentication
Detect password leaks and breached credentials
Detect fraudulent activities on websites
Identify user accounts that show related behaviors on websites
Detect fraudulent activities on mobile applications
Detect and prevent account takeovers
Identify and investigate attacks
Automated threats
Best practices for protection from automated threats
Payment transactions
Set up Transaction defense
Annotate assesments with transaction data
Restrict operations on reCAPTCHA keys
Migrate from reCAPTCHA Classic
Use Fraud Defense features after migration
Create dashboards with Looker Studio
Troubleshoot reCAPTCHA integration issues
All code samples for all products
Distributed, hybrid, and multicloud
Access and resources management
SDK, languages, frameworks, and tools
Hand gesture verification
Stay organized with collections
Save and categorize content based on your preferences.
This document provides information about data collection, camera permissions,
and accessibility for hand gesture verification.
When the hand gesture feature is enabled, reCAPTCHA collects the following data:
Google analyzes one or more videos of a user's hand as they perform various
actions or gestures. The video is processed to extract hand landmark data,
which includes
21 hand-knuckle coordinates .
The videos are never associated with a user's identity and are deleted after
the verification process. Audio is never recorded.
Google does not retain any images or videos of a user's hand gestures beyond
the verification process or use the data for any other purpose. Videos or
images are automatically deleted after the challenge is complete.
The information Google collects is used and stored in accordance with the
Google Privacy Policy .
Hand gesture challenges require permission to access a user's camera. After the
user consents to perform hand gestures, Google receives camera permissions, and
users can manage these permissions in their browser settings at any time. Google
processes hand gesture videos for the sole purpose of security verification and
does not transfer any related data or permissions to third parties.
For users with accessibility needs who cannot use hand gestures to complete the
challenge, reCAPTCHA continues to provide visual and audio
challenges, and develop more accessible and secure alternatives.
Except as otherwise noted, the content of this page is licensed under the Creative Commons Attribution 4.0 License , and code samples are licensed under the Apache 2.0 License . For details, see the Google Developers Site Policies . Java is a registered trademark of Oracle and/or its affiliates.
Getting Started with Google Cloud
Our third decade of climate action: join us
Sign up for the Google Cloud newsletter
Subscribe
