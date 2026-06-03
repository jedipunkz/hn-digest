---
source: "https://github.com/openai/imagegencam"
hn_url: "https://news.ycombinator.com/item?id=48372591"
title: "Turn the World into Cheese – OpenAI RPI Image Gen Camera"
article_title: "GitHub - openai/imagegencam: A digital camera you can build yourself with Codex. · GitHub"
author: "smoser"
captured_at: "2026-06-03T00:40:50Z"
capture_tool: "hn-digest"
hn_id: 48372591
score: 2
comments: 0
posted_at: "2026-06-02T16:37:28Z"
tags:
  - hacker-news
  - translated
---

# Turn the World into Cheese – OpenAI RPI Image Gen Camera

- HN: [48372591](https://news.ycombinator.com/item?id=48372591)
- Source: [github.com](https://github.com/openai/imagegencam)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T16:37:28Z

## Translation

タイトル: 世界をチーズに変える – OpenAI RPI Image Gen Camera
記事タイトル: GitHub - openai/imagegencam: Codex を使用して自分で構築できるデジタル カメラ。 · GitHub
説明: Codex を使用して自分で構築できるデジタル カメラ。 - openai/imagegencam

記事本文:
GitHub - openai/imagegencam: Codex を使用して自分で構築できるデジタル カメラ。 · GitHub
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
オープンナイ
/
イメージジェンカム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主な支店のタグ

ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .githooks .githooks 3d モデル 3d モデル アセット アセット ドキュメント ドキュメント スクリプト スクリプト ソフトウェア ソフトウェア third_party/ ライセンス third_party/ ライセンス .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ImageGenCam は、Codex を使用して自分で構築できるデジタル カメラです。
基本的なメーカー部品と 3D プリントされたシェルを使用する ImageGenCam は、
あなた自身のスタイル、興味、アイデアを反映するように設計されたカスタマイズ可能なプロジェクト。
写真を撮り、画像生成で変換し、調整を続けます。
自分のものだと感じるまで体験してください。コンパニオン Web アプリが携帯電話上で実行されます。
写真をダウンロードしたり、プロンプトを更新したりできます。
このプロジェクトは週末のビルドとして設計されていますが、すでに慣れている場合は、
Codex、Raspberry Pi、および 3D プリンティングは、次の環境で実行できる可能性があります。
1時間。私たちは、熱意のある高揚感から誰でも親しみやすいものにすることを目指しました
学生から職人、アーティスト、エンジニアまで。これを作りたいなら、私たちは
確かに（Codex の支援があれば）それは可能です。
以下の各パーツが 1 つ必要になります。 Codex がインストールされた Mac も必要です
インストールされたデスクトップ アプリ、信頼性の高い Wi-Fi 接続、および OpenAI アカウント。もし
ChatGPT を使用している場合は、すでに ChatGPT を持っています。
Raspberry Pi Zero 2 W (ヘッダー付き)
Raspberry Pi Zero用スパイカメラ
注: カメラを家の外で動作させたい場合は、
自宅の Wi-Fi から携帯電話のモバイル ホットスポットに切り替えることをお勧めします。
チュートリアルの終わり。
プロジェクトについてよく知るために、このドキュメントを読んでください。いつ
組み立てる準備ができました。組み立てを始める前に、開くだけです。
コーデックス デスクトップ o

Mac で次のように入力します。
ImageGenCam の作成を手伝ってください https://github.com/openai/imagegencam
次に、このビルドの残りの部分を Codex で直接完了します。コーデックスはこれを読みます
リポジトリにアクセスすると、最初のセットアップ手順に進み、そこからビルドを順を追って説明します。
そこに。
すべてが組み立てられると、カメラは非常に小さなカメラのように動作します。
奇妙なオートフォーカス。プロンプトを選択して写真を撮ると、ChatGPT Images 2.0 が
あなたのイメージを変えます。
シャッター / 電源: 短く押して写真を撮ります。長押しすると電源が切れます。へ
電源オン: 短く押して放し、長押しして放します。
マジック ボタン: 特別なリミックス ボタン。 Codex に何でもできるように依頼してください
あなたが欲しいのです。
右上/右下: 上/下。
ライブ プレビューで右上をトリプルタップ: Wi-Fi 設定を開きます。
右上をトリプルクリック: コンパニオン アプリの QR コードを表示します。
写真を撮ると、ビューファインダーが一瞬フリーズし、その後フェードバックします。
ライブプレビュー。画像の生成はバックグラウンドで継続されるため、
カメラが動作している間に撮影します。画像の準備が完了すると、アルバムアイコンが
輝きます。
プロンプトを試してください。変更してください。より具体的、より混沌とした、より有用なものにする、または
さらに呪われた。モバイル アプリを使用してプロンプト エディターを開いていろいろ試してみます。
それ。組み込みプロンプトには次のものが含まれます。
Pathetic Scribble - 画像を不器用で低品質のマウス描画として再描画します。
落書き。
チーズに向きを変える - 全員が向きを変える以外は、この写真については何も変更しません
チーズに。
ゴブリン モード - 被写体をキュートで気難しいファンタジー キャラクターに変えます。
手作りのインディーズウェブコミックスタイル。
アニメポートレート - 被写体を明るく表情豊かに変えます。
様式化されたアニメ/漫画のポートレート。
リマインダー: アクセスするには、携帯電話がカメラと同じ Wi-Fi ネットワーク上にある必要があります
モバイルアプリ。
基本が機能すれば、プログラム可能な小さなカメラ プラットフォームが完成します。
役に立ちますよ、マック

それは呪われました、それを美しくし、それをあなたのものにしてください。
チャームでスタイリングしたり、完全にカスタムのカメラケースをデザインしたり、
Codex にコードのカスタマイズを依頼したり、独自に作成できるものはたくさんあります。起動画面を変更する、UI のスタイルを変更する、マジック ボタンに機能を追加する、または
何か賢くて、創造的で、あなた自身のものをしてください。
3D モデリングのスキルがある場合は、3D プリント カメラ ケースを使用して、選択したモデリング ツールに .step ファイルをインポートして、新しい形状や機能をモデル化することもできます。
これはどのプリンターでも機能しますか?
PETG と PLA を使用したさまざまな一般的な 3D プリンターでこれをテストしました。
フィラメント。カメラの筐体は折りたたみ式であるため、PETG をお勧めします。
これを作りたいのですが3Dプリンターを持っていないのですがどうすればいいでしょうか？
3Dプリンターがなくても問題ありません。 ChatGPT にオンライン 3D プリントを勧めてもらうだけです
サービス。このファイルに .stl ファイルをアップロードできるサービスは数多くあります。
リポジトリを印刷して配送します。地元でも確認できます
図書館に行くか、プリンターを持っている友人に手伝ってもらいましょう。印刷をスキップすることもできます
ケースは完全に覆われていますが、電子機器はカバーなしで放置すると非常に壊れやすくなります。
このプロジェクトは、Apache License バージョン 2.0 に基づいてライセンスされています。 「ライセンスと通知」を参照してください。
Codex を使用して自分で構築できるデジタル カメラ。
Apache-2.0ライセンス
セキュリティポリシー
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
2
フォーク
レポートリポジトリ
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A digital camera you can build yourself with Codex. - openai/imagegencam

GitHub - openai/imagegencam: A digital camera you can build yourself with Codex. · GitHub
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
openai
/
imagegencam
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .githooks .githooks 3d model 3d model assets assets docs docs scripts scripts software software third_party/ licenses third_party/ licenses .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
ImageGenCam is a digital camera you can build yourself with Codex.
Using basic maker parts and a 3D-printed shell, ImageGenCam is a highly
customizable project designed to reflect your own style, interests, and ideas.
Take photos, transform them with image generation, and keep tuning the
experience until it feels like yours. A companion web app runs on your phone,
letting you download photos and update prompts.
This project is designed as a weekend build, but if you’re already familiar with
Codex, Raspberry Pi, and 3D printing, you may be able to get it running in under
an hour. We’ve aimed to make it approachable for anyone from an eager high
schooler to craftspeople, artists, and engineers. If you want to make this, we’re
pretty sure (with an assist from Codex) you can.
You'll need one of each part below. You'll also need a Mac with the Codex
Desktop app installed, a reliable Wi-Fi connection, and an OpenAI account. If
you use ChatGPT, you already have one.
Raspberry Pi Zero 2 W with headers
Spy Camera for Raspberry Pi Zero
Note: If you want your camera to work outside your home, we highly
recommend switching it from your home Wi-Fi to your phone's mobile hotspot at
the end of the tutorial.
Feel free to read through this document to get familiar with the project. When
you're ready to build, before you even get started assembling anything, just open
Codex Desktop on your Mac and type:
Help me make ImageGenCam https://github.com/openai/imagegencam
Then complete the rest of this build directly in Codex. Codex will read this
repo, take you to the first setup step, and walk you through the build from
there.
Once everything is assembled, your camera works a lot like a very small, very
weird point-and-shoot. Pick a prompt, snap a photo, and ChatGPT Images 2.0 will
transform your image.
Shutter / Power: short press to take a photo. Long press to power off. To
power on: short press, release, long press, release.
Magic Button: your special remix button. Ask Codex to make it do whatever
you want.
Top-right / Bottom-right: up / down.
Triple-tap top-right from live preview: open Wi-Fi settings.
Triple-click top-right: show a QR code for the companion app.
When you take a photo, the viewfinder freezes for a moment, then fades back to
live preview. Image generation continues in the background, so you can keep
shooting while the camera does its thing. When an image is ready, the album icon
will sparkle.
Try a prompt. Change it. Make it more specific, more chaotic, more useful, or
more cursed. Use the mobile app to open the prompt editor and play around with
it. Built-in prompts include:
Pathetic Scribble - Redraws the image as a clumsy, low-quality mouse-drawn
scribble.
Turn to Cheese - Change nothing about this photo except everyone is turned
into cheese.
Goblin Mode - Turns the subject into a cute scrappy fantasy character in a
handmade indie webcomic style.
Anime Portrait - Converts the subject into a bright, expressive,
stylized anime/cartoon portrait.
Reminder: your phone must be on the same Wi-Fi network as the camera to access
the mobile app.
Once the basics are working, you've got a tiny programmable camera platform.
Make it useful, make it cursed, make it beautiful, make it yours.
Whether you're styling it with charms, designing an entirely custom camera case,
or asking Codex to customize code, there's a lot you can make your own. Try asking to: change the boot up screen, restyle the UI, add a feature to the Magic Button, or
do something clever, creative and all your own.
If you've got skills in 3d modeling, you can even take the 3D Printed Camera Case and model new shapes and features by importing the .step file into the modeling tool of your choosing.
Will this work on any printer?
We've tested this with a variety of common 3D printers using PETG and PLA
filament. We recommend PETG due to the folding nature of the camera enclosure.
I want to make this but I don't have a 3D printer, what can I do?
No 3D printer, no problem. Just ask ChatGPT to recommend an online 3D printing
service. There are many services where you can upload the .stl files in this
repository to be printed and shipped to you. You can also check your local
library, or ask a friend with a printer to help. You can also skip the printed
case entirely, but the electronics are quite fragile if left uncovered.
This project is licensed under the Apache License, Version 2.0. See LICENSE and NOTICE .
A digital camera you can build yourself with Codex.
Apache-2.0 license
Security policy
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
2
forks
Report repository
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
