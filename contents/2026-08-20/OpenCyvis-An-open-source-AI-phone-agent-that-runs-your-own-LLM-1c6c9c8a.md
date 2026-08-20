---
source: "https://github.com/opencyvis/opencyvis-phone"
hn_url: "https://news.ycombinator.com/item?id=49381419"
title: "OpenCyvis: An open-source AI phone agent that runs your own LLM"
article_title: "GitHub - opencyvis/opencyvis-phone · GitHub"
image: "https://opengraph.githubassets.com/0e64c42a829202e5dcda5f11b5085a31a99d79511f8eba4681ded060712140dc/opencyvis/opencyvis-phone"
author: "Alephinitesimal"
captured_at: "2026-08-20T23:16:43Z"
capture_tool: "hn-digest"
hn_id: 49381419
score: 1
comments: 0
posted_at: "2026-08-20T22:57:56Z"
tags:
  - hacker-news
  - translated
---

# OpenCyvis: An open-source AI phone agent that runs your own LLM

- HN: [49381419](https://news.ycombinator.com/item?id=49381419)
- Source: [github.com](https://github.com/opencyvis/opencyvis-phone)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T22:57:56Z

## Translation

タイトル: OpenCyvis: 独自の LLM を実行するオープンソース AI 電話エージェント
記事タイトル: GitHub - opencyvis/opencyvis-phone · GitHub
説明: GitHub でアカウントを作成して、opencyvis/opencyvis-phone の開発に貢献します。

記事本文:
GitHub - opencyvis/opencyvis-phone · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
オープンシービス
/
opencyvis-phone
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
14 コミット 14 コミット フォルダーとファイル
.github .github android android aosp-integration/ OpenCyvis aosp-integration/ OpenCyvis docs docs scripts/ ci スクリプト/ ci テスト テスト .gitignore .gitignore CH

ANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 README.md README.md README_CN.md README_CN.md github-banner-dark.png github-banner-dark.png github-banner-light.png github-banner-light.png github-logo.png github-logo.png github-social-preview.png github-social-preview.png logo.png logo.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープンソースの AI 電話。
商用 AI 電話はブラックボックスです。これは違います。
Open Cyber Jar vis
ドキュメント •
中国語 •
はじめに •
ロードマップ •
貢献する
OpenCyvis は Android を AI 電話に変えます。自然言語でタスクを与えます。ユーザーと同じように画面を見て、UI を理解し、アプリを操作します。
「WeChat で「寿司を食べに行きましょう」と返信します」 — WeChat を開き、会話を見つけて、次のように入力して送信します。
「近くの最高評価のコーヒー ショップを見つけて道順を確認」 — 地図を開き、検索し、評価で並べ替え、タップしてナビゲートします。
「午前 7 時にアラームを設定し、おやすみモードをオンにして、ダーク モードに切り替える」 — 時計、設定、ディスプレイを一度に連鎖させます。
ほとんどの AI ツールは、動作中に画面をロックします。 OpenCyvis は、仮想ディスプレイ、つまり分離された背景画面上で動作します。 Twitter をスクロールしている間に、AI がフライトを予約します。
┌─────────────┐ ┌─────────────┐
│ あなたの画面 │ │ 仮想ディスプレイ │
│ │ │ (ここで AI が働きます) │
│ 閲覧、チャット、 │ │
│ ビデオを見る — │ │ フライトの予約、 │
│ 電話はあなたのものです │ │ メッセージを送信します │
│ │ │ ご注文 │
━━━━━━━━━━━━━━━━━━┘ ━━━━━━━━━━━━━━━━━┘
あなたはこれを使います AIはこれを使います
AI の動作をいつでも観察できます。何かあった場合は引き継いでください

間違っています。終わったら返してください。
企業が「AI 電話」を出荷すると、その企業は画面、アプリ、メッセージに完全にアクセスできるようになります。しかし、どのモデルが実行されているかを確認することも、デバイスから送信されるデータを確認することも、代替手段を選択することもできません。
少なくとも選択の余地はあるはずです。
OpenCyvis はオープンソースの代替手段です。コードのすべての行を確認し、AI モデルを選択し、データの保存先を決定します。ローカル モデルでは、デバイスから何も出ていくことはありません。
ほとんどのユーザーにとって。カスタムROMもルートもコンピューターもありません。
アプリを開き、セットアップ ウィザードに従って ADB ワイヤレス ペアリングを完了します
LLM バックエンド (クラウドまたはローカル) を選択し、使用を開始します
ペアリング プロセス全体がデバイス上で完了します。 Android 11以降をサポートします。
ウィザードは、「ワイヤレス デバッグ」→「ペアリング コードでデバイスをペアリング」に進みます。 6 桁のコードが表示されたら、通知シェードをプルダウンして、OpenCyvis 通知にコードを直接入力します。アプリに戻る必要はありません。
一部のベンダー ROM (ColorOS / OnePlus、MIUI など) では、アプリがバックグラウンドに移行するとフリーズします。ペアリング中にバックグラウンド アクティビティを許可するか、バッテリーの最適化を無視するかをシステムが要求した場合は、[許可] を選択すると、よりスムーズなエクスペリエンスが得られます。
入力フィールドに通知が表示されない場合は、シェードを一度引き下げて確認してください。
開発者とパワーユーザー向け。
AOSP システム イメージをフラッシュします。アプリは、完全なプラットフォーム署名権限を持つシステム アプリケーションとして実行されます。スクリーンショットは SurfaceControl を直接使用します - 可能な限り高速です。システム API を介した完全な仮想ディスプレイ タスク管理。
同じ AI エンジン、同じ LLM バックエンド、同じ UI、同じ機能。唯一の違いは、アプリがシステム権限を取得する方法です。標準モードでは、ADB シェル権限を使用します。システム アプリ モードではプラットフォーム署名を使用します。日常的な作業では、違いに気づくことはありません。

IM アプリのボットにメッセージを送信して、携帯電話の AI をリモートで制御します。現在、Feishu と Telegram をサポートしています。
使用例: 親の電話に OpenCyvis をインストールします。お母さんが「テキストが小さすぎる」と言います。あなたは IM で「フォント サイズを最大に設定してください」と送信します。 AI がそれを実行し、確認のスクリーンショットを送り返します。双方が同時に画面を見る必要はありません。
サポート: コマンドの送信、進行状況の受信、スクリーンショットの表示、AI の質問への回答、タスクの停止。ペアリングには6桁のコードを使用します。
頻繁に行う操作を保存し、スケジュールに基づいて、またはワンタップで実行します。
例: 「毎朝午前 8 時にカレンダー、天気、未読メールをチェック」 — AI が自動的に実行され、概要がチャットにプッシュされます。ジオフェンシングもサポートしており、オフィスに到着したときに自動で出勤します。
複数の AI 設定 (クラウド Qwen、ローカル Gemma 4、Claude など) を保存し、ワンタップでそれらを切り替えます。毎回 API URL とキーを再入力する必要はありません。
完全な昼/夜のテーマ。システム設定に従うか、手動で設定します。ホーム、チャット、設定、時計モードなどのすべての画面に、一致するダーク バリアントがあります。
標準モードでは、MIUI、ColorOS、OriginOS、およびその他のベンダーの ROM がサポートされるようになりました。メーカーが異なれば、ワイヤレス デバッグ エントリ ポイントは大きく異なります。OemHelper がその違いを処理します。
商用AIフォン
クラウドフォン
電話制御スクリプト
オープンサイビス
オープンソース
❌
❌
⚠️
✅
AI モデルを選択してください
❌
❌
⚠️
✅
データはデバイス上に残ります
❌
❌
⚠️
✅
AIが動作している間も電話は使用可能
⚠️
✅
❌
✅
どのアプリでも動作します
⚠️
⚠️
⚠️
✅
コンピューターのセットアップがありません
⚠️
⚠️
❌
✅
日常の携帯電話で動作します
✅
⚠️
❌
✅
対応機種
OpenCyvis はモデルに依存しません。独自の AI アカウントを使用したり、プライベート サービスに接続したり、ローカル モデルを実行したりできます。
モデル
ステップごとのレイテンシ
合格率
注意事項
クウェン３。

5プラス
4～6秒
4/4
安定している、推奨される
クロード 作品4
4-8秒
4/4
最高の推論品質
MiMo v2.5
2.3～4.5秒
4/4
最速
GPT-4o
3-6秒
3/4
場合によっては、tool_choice が無視される
ローカルモデル (Ollama経由)
モデル
サイズ
速度
合格率
ジェマ 4 26B-A4B Q4
17GB
63トーク/秒
4/4
ジェマ 4 E2B Q4
1.8GB
41トーク/秒
4/4
クウェン 3.5 35B-A3B Q4
22GB
47トーク/秒
3/4
ジェマ 4 E4B Q4
3GB
61トーク/秒
3/4
推奨: Gemma 4 26B-A4B — 速度、品質、メモリの最適なバランス。
最小: Gemma 4 E2B — わずか 1.8 GB、それでも 4 つのテストすべてに合格します。
どちらのインストール モードでも、すべての上位層コードを共有します。違いは、PrivilegeBackend インターフェイスの背後に分離された特権層のみです。
バックエンドは実行時に自動的に選択されます。
電話に完全にアクセスできる AI エージェントは、実行できるソフトウェアの中で最も特権のあるものの 1 つです。ここは「私たちを信頼してください」という場所ではありません。
AI サービスをホスト型、プライベート、またはローカルから選択します。
テレメトリ、分析、テレホンホームは不要 - トラッキングコードはゼロ
オープンソース — 誰でも監査できる
ローカル モデル オプション — デバイスから何も残らない
リリース ページでは 2 つの APK が入手可能です。
opencyvis-standard-release.apkをダウンロードしてインストールします
アプリを開き、セットアップ ウィザードに従ってワイヤレス ペアリングを完了します
[設定] で LLM プロバイダーを構成します
rootもコンピューターもカスタムROMも必要ありません。ペアリング中に、OpenCyvis 通知に 6 桁のコードを入力し (シェードをプルダウンします)、ROM が要求した場合はバックグラウンド アクティビティを許可します。
カスタム AOSP イメージを構築する開発者向け:
git clone https://github.com/opencyvis/opencyvis-phone.git
cd opencyvis-phone/android
./gradlew AssemblySystemRelease
AOSP の展開とプラットフォーム キーの署名については、android/README-AOSP.md を参照してください。
アプリ内またはディープリンク経由でプロバイダーを設定します。
# ローカルオラマ（完全プライベート）
adb シェル am start -a android.intent.action.VIEW \
-d " オプション

encyvis://config?provider=ollama&base_url=http://localhost:11434&model=gemma4:26b "
# クラウドAPI
adb シェル am start -a android.intent.action.VIEW \
-d " opencyvis://config?provider=openai&base_url=https://api.example.com/v1&api_key=YOUR_KEY&model=qwen-vl-max "
ロードマップ
より便利な特典取得方法を検討する
ローカルモデルのサポートをさらに最適化
クロスデバイス調整 (電話 + デスクトップ)
COTRIBUTING.md を参照してください。コード、バグレポート、セキュリティ監査、翻訳、ドキュメントを歓迎します。
Sherpa-ONNX — オンデバイス音声認識 (Apache 2.0)
Readme Apache-2.0 ライセンス
貢献活動 カスタム プロパティ スター
57 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to opencyvis/opencyvis-phone development by creating an account on GitHub.

GitHub - opencyvis/opencyvis-phone · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
opencyvis
/
opencyvis-phone
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
.github .github android android aosp-integration/ OpenCyvis aosp-integration/ OpenCyvis docs docs scripts/ ci scripts/ ci tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md README_CN.md README_CN.md github-banner-dark.png github-banner-dark.png github-banner-light.png github-banner-light.png github-logo.png github-logo.png github-social-preview.png github-social-preview.png logo.png logo.png View all files Repository files navigation
The open-source AI phone.
Commercial AI phones are black boxes. This one isn't.
Open Cy ber Jar vis
Docs •
中文 •
Getting Started •
Roadmap •
Contributing
OpenCyvis turns Android into an AI phone. Give it a task in natural language — it sees your screen, understands the UI, and operates apps just like you would.
"Reply 'let's go eat sushi' in WeChat" — opens WeChat, finds the conversation, types and sends:
"Find the best-rated coffee shop nearby and get directions" — opens Maps, searches, sorts by rating, taps navigate.
"Set a 7am alarm, turn on Do Not Disturb, and switch to dark mode" — chains Clock, Settings, and Display in one go.
Most AI tools lock your screen while they work. OpenCyvis operates on a virtual display — an isolated background screen. The AI books your flight while you scroll Twitter.
┌─────────────────────┐ ┌─────────────────────┐
│ Your screen │ │ Virtual display │
│ │ │ (AI works here) │
│ Browse, chat, │ │ │
│ watch videos — │ │ Booking flights, │
│ phone is yours │ │ sending messages, │
│ │ │ placing orders │
└─────────────────────┘ └─────────────────────┘
You use this AI uses this
Watch the AI work anytime. Take over if something looks wrong. Hand it back when you're done.
When a company ships an "AI phone," they get full access to your screen, your apps, your messages — and you can't see what model is running, can't verify what data leaves your device, can't choose an alternative.
You should at least have the choice.
OpenCyvis is the open-source alternative: you see every line of code, you pick the AI model, you decide where your data goes. With a local model, nothing ever leaves your device.
For most users. No custom ROM, no root, no computer.
Open the app, follow the setup wizard to complete ADB wireless pairing
Choose your LLM backend (cloud or local), start using
The entire pairing process completes on-device. Supports Android 11+.
The wizard sends you to Wireless debugging → Pair device with pairing code . Once you see the 6-digit code, pull down the notification shade and type the code straight into the OpenCyvis notification — no need to switch back to the app.
Some vendor ROMs (ColorOS / OnePlus, MIUI, …) freeze apps once they go to the background. If the system asks you to allow background activity / ignore battery optimization during pairing, choose Allow for a smoother experience.
If you don't see the notification with the input field, just pull down the shade once to find it.
For developers and power users.
Flash an AOSP system image. The app runs as a system application with full platform signing privileges. Screenshots use SurfaceControl directly — fastest possible. Full virtual display task management via system APIs.
Same AI engine, same LLM backends, same UI, same capabilities. The only difference is how the app obtains system permissions. Standard mode uses ADB shell privileges; System App mode uses platform signing. For everyday tasks, you won't notice the difference.
Send messages to a bot in your IM app to control the phone's AI remotely. Currently supports Feishu and Telegram .
Use case: Install OpenCyvis on a parent's phone. Mom says "the text is too small" — you send "set font size to largest" in IM. The AI does it and sends back a confirmation screenshot. No need for both parties to watch the screen simultaneously.
Supports: sending commands, receiving progress, viewing screenshots, answering the AI's questions, stopping tasks. Pairing uses a 6-digit code.
Save frequent operations and run them on a schedule or with one tap.
Example: "Check calendar, weather, and unread emails every morning at 8am" — the AI runs automatically and pushes a summary to chat. Also supports geofencing — auto clock-in when arriving at the office.
Save multiple AI configurations (e.g., cloud Qwen, local Gemma 4, Claude) and switch between them with one tap. No need to re-enter API URLs and keys each time.
Full day/night theme. Follows system settings or set manually. All screens — home, chat, settings, watch mode — have matching dark variants.
Standard mode now supports MIUI, ColorOS, OriginOS, and other vendor ROMs. Different manufacturers have vastly different wireless debugging entry points — OemHelper handles the differences.
Commercial AI Phones
Cloud Phones
Phone-control Scripts
OpenCyvis
Open source
❌
❌
⚠️
✅
Choose your AI model
❌
❌
⚠️
✅
Data stays on device
❌
❌
⚠️
✅
Phone usable while AI works
⚠️
✅
❌
✅
Works with any app
⚠️
⚠️
⚠️
✅
No computer setup
⚠️
⚠️
❌
✅
Works on everyday phones
✅
⚠️
❌
✅
Supported Models
OpenCyvis is model-agnostic. Bring your own AI account, connect a private service, or run a local model.
Model
Latency per step
Pass Rate
Notes
Qwen 3.5 Plus
4-6s
4/4
Stable, recommended
Claude Opus 4
4-8s
4/4
Highest reasoning quality
MiMo v2.5
2.3-4.5s
4/4
Fastest
GPT-4o
3-6s
3/4
Occasionally ignores tool_choice
Local Models (via Ollama)
Model
Size
Speed
Pass Rate
Gemma 4 26B-A4B Q4
17 GB
63 tok/s
4/4
Gemma 4 E2B Q4
1.8 GB
41 tok/s
4/4
Qwen 3.5 35B-A3B Q4
22 GB
47 tok/s
3/4
Gemma 4 E4B Q4
3 GB
61 tok/s
3/4
Recommended: Gemma 4 26B-A4B — best balance of speed, quality, and memory.
Minimal: Gemma 4 E2B — just 1.8 GB, still passes all 4 tests.
Both install modes share all upper-layer code. The difference is only in the privilege layer, isolated behind a PrivilegeBackend interface:
The backend is selected automatically at runtime.
An AI agent with full phone access is one of the most privileged pieces of software you can run. This is not a place for "trust us."
You choose the AI service — hosted, private, or local
No telemetry, no analytics, no phone-home — zero tracking code
Open source — anyone can audit
Local model option — nothing leaves your device
Two APKs are available on the Releases page:
Download opencyvis-standard-release.apk and install
Open the app, follow the setup wizard to complete wireless pairing
Configure your LLM provider in Settings
No root, no computer, no custom ROM required. During pairing, type the 6-digit code into the OpenCyvis notification (pull down the shade), and allow background activity if your ROM prompts for it.
For developers building a custom AOSP image:
git clone https://github.com/opencyvis/opencyvis-phone.git
cd opencyvis-phone/android
./gradlew assembleSystemRelease
See android/README-AOSP.md for AOSP deployment and platform key signing.
Set your provider in-app, or via deeplink:
# Local Ollama (fully private)
adb shell am start -a android.intent.action.VIEW \
-d " opencyvis://config?provider=ollama&base_url=http://localhost:11434&model=gemma4:26b "
# Cloud API
adb shell am start -a android.intent.action.VIEW \
-d " opencyvis://config?provider=openai&base_url=https://api.example.com/v1&api_key=YOUR_KEY&model=qwen-vl-max "
Roadmap
Explore more convenient privilege acquisition methods
Further optimize local model support
Cross-device coordination (phone + desktop)
See CONTRIBUTING.md . We welcome code, bug reports, security audits, translations, and documentation.
Sherpa-ONNX — on-device speech recognition (Apache 2.0)
Readme Apache-2.0 license Contributing
Contributing Activity Custom properties Stars
57 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
