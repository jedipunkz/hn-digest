---
source: "https://github.com/Wavesonics/C2PAVerify"
hn_url: "https://news.ycombinator.com/item?id=48606789"
title: "I created an open source app to check if a photo is from an AI or a Camera"
article_title: "GitHub - Wavesonics/C2PAVerify: An Android app for Viewing and Verifying the C2PA data embeded in photos. · GitHub"
author: "Wavesonics"
captured_at: "2026-06-20T07:26:32Z"
capture_tool: "hn-digest"
hn_id: 48606789
score: 2
comments: 1
posted_at: "2026-06-20T06:12:23Z"
tags:
  - hacker-news
  - translated
---

# I created an open source app to check if a photo is from an AI or a Camera

- HN: [48606789](https://news.ycombinator.com/item?id=48606789)
- Source: [github.com](https://github.com/Wavesonics/C2PAVerify)
- Score: 2
- Comments: 1
- Posted: 2026-06-20T06:12:23Z

## Translation

タイトル: 写真が AI かカメラからのものかを確認するオープンソース アプリを作成しました
記事タイトル: GitHub - Wavesonics/C2PAVerify: 写真に埋め込まれた C2PA データを表示および検証するための Android アプリ。 · GitHub
説明: 写真に埋め込まれた C2PA データを表示および検証するための Android アプリ。 - ウェーブソニクス/C2PAVerify

記事本文:
GitHub - Wavesonics/C2PAVerify: 写真に埋め込まれた C2PA データを表示および検証するための Android アプリ。 · GitHub
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
ウェーブソニクス
/
C2PA検証
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション最適化

オン
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
48 コミット 48 コミット .github/ workflows .github/ workflows app app docs docs fastlane fastlane gradle gradle scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md Gemfile Gemfile LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md build.gradle.kts build.gradle.kts c2pa_logo.svg c2pa_logo.svg gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat settings.gradle.kts settings.gradle.kts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
写真を作成した人、およびその写真が信頼できるかどうかを確認します。
C2PA Verify は、コンテンツ認証情報を読み取る Android アプリです。
(C2PA) が写真に埋め込まれており、その画像の出所、つまり誰が署名したかが一目で分かります。
それを作成したツール、編集されたか AI によって生成されたか、署名者が
信頼できるもの。
C2PA は、AI がますます傾斜する世界で何が真実かを知るための戦いにおいて重要な標準となる可能性があります。
写真が撮影されたことを証明するための署名をサポートする一部のカメラが市場に登場します
本物のカメラを使って。一部の画像編集ソフトウェアは、すでに C2PA データの処理をサポートしています。しかし、一つあります
ギャップ、言ってみればラストワンマイル。写真に完全な C2PA チェーンが含まれていても、誰も読み取れない場合はどうなるのでしょうか。
それは？
この情報は、便利​​でアクセスしやすい方法でユーザーに表示される必要があります。そこでこのアプリは、
それは初めての試み。将来的には、このようなツールが Web ブラウザに直接組み込まれるようになるのではないかと思います
と画像ビューアがありますが、それができるまでは、このツールが何人かの人々に役立つことを願っています。
Dark Rock Studios は、フリーのオープンソース ソフトウェアを構築することに重点を置いています。
🐛 バグを見つけましたか?
💡 ご提案はありますか?
📚 翻訳を手伝ってみませんか?
🎮 他のアプリにも興味がありますか?
👉 Open So のコミュニティに参加してください

Discord の熱狂的なファン!
./gradlew アセンブルデバッグ
Android Studio でプロジェクトを開き、デバイスまたはエミュレーターでアプリ構成を実行します。
写真に埋め込まれた C2PA データを表示および検証するための Android アプリ。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
5
v0.2.3
最新の
2026 年 6 月 11 日
+ 4 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An Android app for Viewing and Verifying the C2PA data embeded in photos. - Wavesonics/C2PAVerify

GitHub - Wavesonics/C2PAVerify: An Android app for Viewing and Verifying the C2PA data embeded in photos. · GitHub
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
Wavesonics
/
C2PAVerify
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
48 Commits 48 Commits .github/ workflows .github/ workflows app app docs docs fastlane fastlane gradle gradle scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md Gemfile Gemfile LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md build.gradle.kts build.gradle.kts c2pa_logo.svg c2pa_logo.svg gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat settings.gradle.kts settings.gradle.kts View all files Repository files navigation
See who made a photo, and whether you can trust it.
C2PA Verify is an Android app that reads the Content Credentials
(C2PA) embedded in a photo and tells you, at a glance, where the image came from: who signed it,
the tool that created it, whether it has been edited or generated by AI, and whether the signer is
one you can trust.
C2PA could be an important standard in the fight to know whats real in our increasingly AI slopped world.
Some cameras are coming on the market that supported signing their photos to prove they were taken
with a real camera. Some image editing software already supports C2PA data handling. But there is one
gap, the last mile if you will. What does it matter if a photo has a full C2PA chain, if no one can read
it?
We need to surface this information to users in a useful and easy to access manner. So this app is
a first attempt at that. I imagine in the future tools like this will be build right into web-browsers
and image viewers, but until they are, I hope this tool might be useful to a couple people.
Dark Rock Studios is all about building Free and Open Source Software .
🐛 Found bugs?
💡 Have suggestions?
📚 Want to help translate?
🎮 Interested in our other apps?
👉 Join our community of Open Source enthusiasts on Discord !
./gradlew assembleDebug
Open the project in Android Studio and run the app configuration on a device or emulator.
An Android app for Viewing and Verifying the C2PA data embeded in photos.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
5
v0.2.3
Latest
Jun 11, 2026
+ 4 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
