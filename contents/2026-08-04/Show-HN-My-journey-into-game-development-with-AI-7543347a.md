---
source: "https://playdowntime.com/"
hn_url: "https://news.ycombinator.com/item?id=49168567"
title: "Show HN: My journey into game development with AI"
article_title: "Downtime - Offline Puzzle Games for Android"
author: "gennadiygryva"
captured_at: "2026-08-04T13:50:54Z"
capture_tool: "hn-digest"
hn_id: 49168567
score: 1
comments: 0
posted_at: "2026-08-04T13:19:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: My journey into game development with AI

- HN: [49168567](https://news.ycombinator.com/item?id=49168567)
- Source: [playdowntime.com](https://playdowntime.com/)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T13:19:54Z

## Translation

タイトル: Show HN: AI を使用したゲーム開発への私の旅
記事のタイトル: ダウンタイム - Android 用オフライン パズル ゲーム
説明: インターネットがなくても動作する Android 用の 2 つの完全なパズル ゲーム。広告もアカウントもデータも収集されません。
HN text: 私はもともと開発の経歴があり、20 年以上前に別の分野に移りましたが、定期的に何かを作成しようとする試みに懐かしさを感じることがよくあります。最後に共有したいのは、AI のスキルアップの一環として 2 か月前に始めたものです。何を開発するかを選択するのは簡単でした。オフィスへの通勤に多くの時間を費やしていたので、広告なしで 5 ～ 10 分以内にオフラインでプレイできるシンプルなゲームが必要でした。私はモバイル ゲームを開発したことがありませんでした。これが実験のポイントでした。どこまで到達できるか? :) ステージ 0 (フル コントロール & 信頼なし) - VS Code + Android Studio から始めました。はい、クロードは何かを作成しましたが、コンパイルとエラー修正は昔ながらの面倒な作業でした。最初はグーグルで検索し、AIに「エラーをチェックして修正してください」と依頼しました。その後、VS Code も Android Studio も必要ないことに気づきました。コードをレビューする必要がある場合、プロジェクトフォルダーで直接レビューできます。ステージ 1 (信頼の構築) - Claude Code CLI に移行しました。私たちはクロードと共同でプロセスとチーム構造 (コア フレームワーク、ゲーム、QA テストを担当するエージェント、コード レビューを実行する技術リーダー) を定義し、クロードが環境をセットアップしました (バイバイ VS Code)。この段階ではまだコードをいくつかレビューしていましたが、技術リーダーによるレビューを上書きするという信頼を築き始めました。ステージ 2 (AI チームの指示と監督) - 適切な計画を立てて要件 (コンセプト、高レベルのアーキテクチャ、デザイン) を定義し、エージェントがゲームを開発し、自動化された QA テストを追加して、技術リーダーがゲームを開発できるようにすることがより価値があることに気づきました。

コードレビューを行います。この旅で定義した道路上のいくつかのルールと役立つプロセス:
- 主要なアーキテクチャ上の決定は、初期計画中に決定されます。 AI がそれらを回避することは許可されていないため、私の承認を得る必要があります。ほとんどの場合、私たちは最初の境界内でより良い方法を見つけました。初期実装と QA 後に凍結されたコア アプリ フレームワークにも同じことが当てはまります。
- ハード モジュールの所有権 - エージェントがお互いのコードを静かに変更することを防ぎます。技術責任者は依然としてレビュー中にそのようなケースにフラグを立てますが、現在ではこれらはまれな例外です。
- UI デザイン - Claude は UI を行うことができますが、最も得意な分野ではなく、最初は苦労しました...すべての UI 変更が最初にローカル Web サーバーに表示されるときのワークフローを確立しました。適切なレベルに到達するまでには依然として複数の反復が必要ですが、現在ははるかに簡単です。
- デバイス テスト マトリックス - さまざまな画面 (解像度、密度、アスペクト) でどのように見えるかを確認できます。ローエンドからハイエンドの携帯電話やタブレットまで12種類になりました。 1 つのスクリプトはエミュレータ上でアプリを実行し、スクリーンショットを撮ります。別のスクリプトはアプリ画面の問題を特定します。最後は、ローカル Web サーバー上のすべてのスクリーンショット (デバイス別、ポートレート/ランドスケープ、タッチ/ボタン、ダーク/ライト テーマに基づくレイアウト) を表示します。すべて問題がなければ、スクリーンショットはゴールデンとして昇格され、後の QA テストで使用されます。
- トリアージ フラグ - 外出中にアイデアや考えがあるときに役立ちます。私が Github モバイルで問題を提起し、トリアージ フラグを追加すると、クロードが次のセッションでそれを選択します。
- ライブデバイスでの手動テストのための詳細なテスト計画。
- マーケティング資料の場合、AI の焦点は低レベルの詳細に移されます。いくつかのポイントを提案したり、人間が書いた草案の感覚をチェックしたりすると役立ちます。ただし、適切な焦点を当てたストーリーが必要な場合は、手動で行ってください... Key co

プロセス内のコントロールは次のとおりです。
- すべての変更について技術主任エージェントによるピアレビュー。
- デバイスマトリックスに基づく自動化された QA テストと UI テスト。
- 標準化されたリリース プロセス (QA テストの検証、AAB リリースの最終テスト、Play ストアのリストとウェブサイトの更新、およびその他の手順)。
- Play ストアのアップロードは依然として手動の手順であり、すべてのリリース プロセスのチェックと手順に満足した場合に私が行います。そして、アプリケーション自体について少し説明します。これは無料でオフラインで、現在 2 つのゲーム (遊び方を学ぶ必要のない簡単なクラシック ゲーム) が含まれており、さらに多くのゲームが登場する予定です :) Web ページ: https://playdowntime.com

記事本文:
飛行機や通勤のためのオフライン パズル ゲーム。
このアプリケーションの全体的なアイデアは、出張中の私自身の経験に基づいています
旅行や定期的なオフィスへの通勤に。私が試したゲームのほとんどはオンラインを必要としました
接続できなかったり、広告を待ったり視聴したりして時間を無駄にしたり、かなりの時間を必要としたりする
学習とスキルアップのための投資、そして全体的により長い時間の取り組み。
その結果、最小限の学習曲線で、複数のミニゲームを備えた独自のアプリケーションを開発しました。
短いセッションでゲームをすばやく開いてプレイすることができ、完全にオフラインでプレイすることもできます (
機内モード)、あなたに関するデータは収集されません。
最初は 2 つのゲームのみが含まれています:
Tight Fit - ブロックが落ちてくる伝統的なペースの速いゲームです。
Carry On - ボックスをマークされた場所に押し込む必要があります。 99個あります
合計レベルは、そのほとんどが David W. Skinner の古典的な Microban セットからのものです。
画面をスワイプするか、画面上のボタンを使用して、暗い状態と明るい状態を切り替えて再生できます。
テーマ、ポートレートまたはランドスケープ - より使いやすいものを選択してください。
通勤時間や通勤時間などの暇つぶしにぜひお役立てください。
オフグリッド時代。改善とさらなる開発のためのフィードバックやアイデアを共有してください。
gennadiygryva@gmail.com 。
Google Play および Google Play ロゴは Google LLC の商標です。

## Original Extract

Two complete puzzle games for Android that work with no internet at all. No adverts, no account, no data collected.

I had a development background initially, moved out to a different field 20+ years ago, but often have a nostalgia with periodic attempts to create something. Wanted to share the last one, started 2 months ago as part of my upskilling with AI. Choice what to develop was easy - spent lots of time commuting to the office and wanted a simple game playable offline, without ads, within 5-10 minutes. I have never developed mobile games and this was an point of the experiment - how far I can reach? :) Stage 0 (full control & no trust) - I started with VS Code + Android Studio. Yes, Claude created something, however compilation and error correction was an old-style pain. Initially googled, then asked AI "check and correct the error". Then I realised that I don't need neither VS Code nor Android Studio. When I need to review a code, i can do it in project folder directly. Stage 1 (building the trust) - moved to Claude Code CLI. Jointly with Claude we defined the processes and the team structure (agents responsible for core framework, games, QA tests, and the tech lead performing the code review), and then Claude set up the environment (bye-bye VS Code). At this stage, I still reviewed some code but started to build trust into overriding reviews by tech lead. Stage 2 (direction and supervision of the AI team) - I realised that it's more valuable to have a proper planning to define the requirements (concept, high-level architecture and design), and then allow agents to develop the game, add automated QA tests, and then tech lead to do a code review. Some rules on the road and helpful processes that we defined on this journey:
- Key architectural decisions are fixed during the initial planning. AI is not allowed to circumvent them and must ask my approval, and in most cases, we found better ways within the initial boundaries. The same applies to the core app framework that was frozen after the initial implementation and QA.
- Hard module ownership - prevents agents to make quiet changes to each other's code. Technical lead still flags such cases during review, but these are rare exceptions now.
- UI design - Claude can do UI, but it's not the strongest area, and initially it was a pain... We established a workflow when all UI changes are first shown on local web server. It still takes multiple iterations to get to appropriate level, but much easier now.
- Device testing matrix - it allows to see how it looks like on different screens (resolution, density, aspect). We ended up in 12 types from low-end to high-end phones and tablets. One script runs app on emulator and takes screenshots. Another script identifies issues in app screens. The last presents all screenshots (by devices, layouts based on portrait/landscape, touch/buttons, dark/light theme) on local web server. Once all ok, the screenshots are promoted as golden and used in later QA tests.
- Triage flag - helpful when I'm away and have some ideas or thoughts. I simply raise an issue on Github mobile, add a triage flag, and Claude picks it in our next session.
- Detailed test plans for manual testing on live devices.
- For marketing materials, AI's focus is shifted towards low-level details. It's helpful to propose some points or do a sense check of human-written draft. But if you want a story with a correct focus - do it manually... Key controls in the process are:
- peer review by tech lead agent for all changes;
- automated QA tests and UI tests based on the device matrix;
- standardised release process (validation of QA tests, final testing of AAB release, updates to Play Store listing and website, and some other steps);
- Play Store upload is still a manual step, done by me when happy with all release process checks and steps. And a little about the application itself - it's free, offline, currently contains 2 games (easy classic games where you don't need to learn how to play), and more games are coming :) Web-page: https://playdowntime.com

Offline puzzle games for flights and commutes.
The overall idea for this application was based on my own experience when travelling on business
trips and regularly commuting to the office. Most of the games I tried did require an online
connection, or wasted my time waiting for and watching ads, or would require a significant time
investment to learn and skill up, and overall longer time commitment.
As a result, I developed my own application with multiple mini-games - minimal learning curve,
you can quickly open and play the games in short sessions, and can be played fully offline (even in
airplane mode) without any data collected about you.
At the start, it only includes 2 games:
Tight Fit - a traditional fast-paced game with falling blocks.
Carry On - where you need to push boxes into their marked spots. There are 99
levels in total, most of them from David W. Skinner's classic Microban set.
You can play it by swiping on screen, or using on-screen buttons, change between dark and light
themes, portrait or landscape - choose what is more convenient for you.
I hope you will enjoy the application, and it will help you to fill your time during commutes or
off-grid times. And please share your feedback and ideas for improvement and further development to
gennadiygryva@gmail.com .
Google Play and the Google Play logo are trademarks of Google LLC.
