---
source: "https://github.com/morganwilliscloud/billai-bass"
hn_url: "https://news.ycombinator.com/item?id=48896599"
title: "Show HN: BillAI Bass, an AI-Powered Big Mouth Billy Bass Using Strands Agents"
article_title: "GitHub - morganwilliscloud/billai-bass: Turn a Big Mouth Billy Bass into a real-time voice assistant powered by Strands BidiAgent + Amazon Nova 2 Sonic on a Raspberry Pi 5 · GitHub"
author: "mtw14"
captured_at: "2026-07-13T19:13:10Z"
capture_tool: "hn-digest"
hn_id: 48896599
score: 11
comments: 3
posted_at: "2026-07-13T18:18:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: BillAI Bass, an AI-Powered Big Mouth Billy Bass Using Strands Agents

- HN: [48896599](https://news.ycombinator.com/item?id=48896599)
- Source: [github.com](https://github.com/morganwilliscloud/billai-bass)
- Score: 11
- Comments: 3
- Posted: 2026-07-13T18:18:23Z

## Translation

タイトル: ショー HN: BillAI Bass、ストランド エージェントを使用した AI 搭載のビッグマウス ビリー バス
記事タイトル: GitHub - morganwilliscloud/bilrai-bass: Raspberry Pi 5 上の Strands BidiAgent + Amazon Nova 2 Sonic を利用したリアルタイム音声アシスタントにビッグマウス ビリー ベースを変える · GitHub
説明: Big Mouth Billy Bass を Raspberry Pi 5 上の Strands BidiAgent + Amazon Nova 2 Sonic を利用したリアルタイム音声アシスタントに変える - morganwilliscloud/bilrai-bass

記事本文:
GitHub - morganwilliscloud/billai-bass: Raspberry Pi 5 上の Strands BidiAgent + Amazon Nova 2 Sonic を利用したリアルタイム音声アシスタントに Big Mouth Billy Bass を変身させる · GitHub
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
モーガンウィリスクラウド
/
ビライベース
公共
通知
あなたはむ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット iot-identity iot-identity .gitignore .gitignore README.md README.md asoundrc.example asoundrc.example billy.py billy.py motors.py motors.py 要件-frozen.txt 要件-frozen.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Strands Agents 双方向ストリーミング + Amazon Nova 2 Sonic を搭載
ビライベースデモストランド.mp4
Big Mouth Billy Bass をリアルタイムの音声アシスタントに変えましょう。あなたが話すと、魚がしゃべり返します。頭を回転させ、口を自分の声に合わせて口パクし、尻尾を羽ばたかせて強調します。 Fish は Raspberry Pi 5 上で Strands Agents 双方向ストリーミング エージェント ( BidiAgent ) を実行し、Amazon Bedrock 上の Amazon Nova 2 Sonic との間でライブ オーディオをストリーミングします。
ロボット工学の経験は必要ありません。事前のはんだ付け経験は必要ありません。このガイドのベースとなっている人物は、これまで Raspberry Pi に接続したことがなかったのですが、週末にしゃべる魚を手に入れました。
優れた billy-b-assistant プロジェクトからインスピレーションを得て、Strands 上に再構築されました。
ファイル
それは何ですか
README.md
このガイド — ここから始めて、上から下まで進めてください
ビリー・パイ
最終的に動く Python — 話す魚
モーター.py
スタンドアロン モーター テスト リグ - billy.py の前にこれを実行して配線を確認します
asoundrc.example
USB マイク + スピーカーをデフォルトとして固定する ALSA 構成
要件-凍結.txt
既知の良好な Pi ビルドからの正確な依存関係バージョン - 動作保証されたセットアップの場合は pip install -r
iot-アイデンティティ/
オプションの実稼働グレードの認証情報: CloudFormation テンプレート + アクセス キーの代わりに X.509 証明書認証のチュートリアル
AIアシスタントとの併用方法
この README を Claude (または別の AI アシスタント) に貼り付けて、次のように言います。「これを作成しています。手順を 1 つ説明してください」

一度にステップずつ — 次のステップを実行する前に、私が各ステップを確認するのを待ちます。」 これがこのように構築された方法であり、最もよく構築される方法です。行き詰まったときに送信する内容については、以下の「AI アシスタントを使用してこれを構築する」セクションを参照してください。
🤝 AI アシスタントを使用してこれを構築
このプロジェクトは Claude Code の助けを借りて作成されており、任意の AI アシスタントを使用して作業を進めることができます。
このガイド全体を AI に貼り付けて、「これを構築しています。一度に 1 ステップずつ説明します。次のステップを説明する前に、私が各ステップを確認するのを待ちます。」
AI はビルドのペースを調整し、正確なエラー メッセージをデバッグし、ハードウェアがこのガイドとわずかに異なる場合に適応します。このビルド中に AI が優れている点:
正確なエラー メッセージを貼り付けてください。恐ろしい 50 行の Python トレースバックには通常、意味のある行が 1 行あり、AI がそれを見つけます。
写真を送信 — モーターの特定に困っていませんか?どのピンがどれであるかわかりませんか?写真を撮って聞いてください。 (Mac 上の Claude Code の場合: 画像ファイルをターミナルにドロップするか、「~/Downloads/IMG_1234.jpg を見て」と言います。)
「X について何も知りません」と言えば、AI がバックアップして説明します。愚かな質問はありません。
不明な点がある場合は、切断/はんだ付けする前に質問してください。
このガイドでカバーされていないことに遭遇したとしても、それはあなたの失敗ではありません。おそらくガイドの隙間です。 AI とエラー メッセージで解決します。
キットを入手した場合: この表にあるものはすべてすでに揃っています。ゼロから構築: これは Amazon の最初の注文です (総額約 240 ドル、2026 年 1 月の価格)。
SD カードをフラッシュするためのパーソナル コンピューター (Mac/Windows/Linux) — ⚠️ 企業のラップトップは USB ストレージをブロックすることがよくあります。家庭用マシンを使用する
microSD をコンピュータに接続する方法 (ほとんどのカードには SD アダプタが含まれています)
支払い方法を備えた AWS アカウント (雇用主のものではなく、個人用) — アカウントを作成します。

以下の前提条件にある、ロックダウンされたフィッシュ専用 IAM ユーザー
テープ (電気的には理想的ですが、スコッチで動作します)、小さなプラスドライバー
コンピューターとPiの両方が参加できるホームWiFi
✅ 前提条件 (ビルド日の前にこれらを実行してください)
1. AWS / Amazon Bedrock のセットアップ (最小限の権限 - 正しく行う)
💸 AWS アカウントをまだお持ちですか?このプロジェクトは無料利用枠でカバーされます。新しい AWS アカウントには 100 ドルのクレジット (アクティビティを含めると最大 200 ドル) が付与され、Amazon Bedrock は Nova の使用量に適用されるクレジット付きの無料プランで使用できます。ボイスチャットは安価で、数分間の会話でおよそ 1 ペニーなので、100 時間以上の雑談でクレジットが買えます。無料プランの 2 つの注意点:
新しいアカウントは、Bedrock 呼び出しクォータがゼロに設定された状態で開始されることがあります (支払い履歴はまだありません)。モデルへのアクセスは許可されているが、すべての呼び出しが ValidationException: Operation not allowed で失敗する場合、それがクォータです。Nova Sonic オンデマンド クォータを求める (無料) AWS サポート ケースを提出するか、有料プランにアップグレードしてください (残りのクレジットは引き継がれます)。
無料プランのアカウントは 6 か月後 (またはクレジットがなくなると) 自動的に閉鎖されます。デモフィッシュとしては問題ありません。ビリーが永続的な同居人になる場合は、期限前に有料にアップグレードしてください。未使用のクレジットはアップグレード後も残ります。
AWS キーは、Fish 内の Pi 上の平文ファイル内に存在します。魚はデモされ、貸し出され、机の上に放置されます。 SDカードが飛び出します。したがって、キーは 1 つのことだけを実行できる必要があります。それは、Nova Sonic と話すことです。漏洩した場合、最悪のケースは、あなたのアカウントで暗号通貨をマイニングしている誰かではなく、誰かがあなたのお金のために魚とチャットしていることです。
AWS コンソールで、リージョンを us-east-1 (バージニア北部) に切り替えます。Nova Sonic はそこに存在します。
「Amazon Bedrock」→「Model access」に移動し、Nova 2 Sonic (amazon.nova-2-sonic-v1:0) へのアクセスをリクエストします。通常はインスタントです。 (v1 ではありません — オリジナルの Nova Sonic は終了に達します

-ライフ 2026 年 9 月)
ステップ 2 — 専用の IAM ユーザーを作成します。
コンソール → IAM → ユーザー → ユーザーの作成 。 billy-bass のような正直な名前を付けます。
「AWS マネジメントコンソールへのユーザーアクセスを提供する」にはチェックを入れないでください。このユーザーはロボットフィッシュです。何もログインすることはありません。
権限画面で、「ポリシーを直接アタッチする」を選択しますが、AWS 管理のポリシーは選択しません (AmazonBedrockFullAccess は選択しません。魚が必要とするものよりもはるかに多くの権限を付与します)。クリックして、権限がゼロのユーザーを作成するだけです。
ステップ 3 — Nova Sonic のみを許可するポリシーをアタッチします。
新しいユーザー→権限→権限の追加→インラインポリシーの作成→JSONタブを開きます。
貼り付け:
{
"バージョン" : " 2012-10-17 " 、
「ステートメント」: [
{
"シド" : " BillyTalksToNovaSonicOnly " ,
"効果" : "許可" 、
"アクション" : " bedrock:InvokeModelWithBidirectionStream " 、
"リソース" : " arn:aws:bedrock:us-east-1::foundation-model/amazon.nova-2-sonic-v1:0 "
}
]
}
billy-nova-sonic-only という名前を付けて保存します。
これは、1 つのリージョンの 1 つのモデルに対する 1 つのアクションです。 S3、EC2、その他の Bedrock モデルはなく、アカウント内のその他のものを表示または変更することはできません。 (後で、AWS に関連する Billy ツールに、たとえば、DynamoDB テーブルの読み取りを与える場合は、今は広範なアクセスではなく、その特定のアクセス許可をこのポリシーに追加します。)
ステップ 4 — アクセスキーを作成します。
ユーザー → セキュリティ認証情報 → アクセスキーの作成 → 「AWS の外で実行されているアプリケーション」を選択します。
アクセス キー ID とシークレットを安全な場所にコピーします。シークレットは 1 回だけ表示されます。これらは後で Pi に適用されます (§1.6)。
ステップ 5 — オプションですが賢明です: 請求アラームを設定します。コンソール → 請求 → 予算 → 電子メール アラート付きで月額 10 ドルの予算を作成します。 Nova Sonic のカジュアルなチャットには数セントかかります。アラームが付いているので、常に聞いている魚が立ち往生したり（またはキーが漏れたり）、驚かれることはありません。
kの場合

情報が漏洩したり、キットフィッシュが行方不明になったりした場合: IAM → ユーザー billy-bass → セキュリティ認証情報 → アクセス キーを非アクティブ化します。ワンクリックでフィッシュレンガが完成し、アカウントは安全です。
2. ペーストの 2 つの注意点を理解する (必ずこれらに遭遇します)
SSH 経由でコードをターミナル/nano に貼り付けると、次の 2 つの問題が発生しやすくなります。
追加のインデント: 貼り付けられたすべての行の先頭にスペースが追加されます。 ALSA 設定は肩をすくめます。 Python が爆発します (IndentationError: 1 行目の予期しないインデントが署名です)。修正: sed -i 's/^ //' yourfile.py 、または慎重に nano に再貼り付けします。
文字の脱落/破損: コンマの欠落、または引用符が「スマート引用符」に変わりました。
あなたを救う習慣: Python ファイルを作成した後、実行します。
python -m py_compile yourfile.py && echo OK
実際に実行する前に。 OK と表示されない場合は、エラーをクロードに貼り付けます。
パート 1 — 脳 (Raspberry Pi、まだフィッシュなし)
戦略メモ: 最初に音声アシスタント全体を裸の Pi で動作させてから、初めてフィッシュを開きます。ソフトウェアの問題とハードウェアの問題は、互いに重なり合うことができない場合、デバッグがはるかに簡単になります。
raspberrypi.com/software から Raspberry Pi Imager をダウンロードし、microSD を挿入します。
次を選択します: デバイス = Raspberry Pi 5 · OS = Raspberry Pi OS Lite (64 ビット) · ストレージ = カード。
⚠️ OS バージョンが重要です: Python 3.12 以降のイメージが必要です (Nova Sonic 要件)。現在の「Trixie」ベースの Raspberry Pi OS は 3.13 ✅ で出荷されます。古い「本の虫」イメージは 3.11 ❌ に出荷されます。
「ライト」 = デスクトップなし。正解です。モニターを接続することはありません。
⚠️「Lite」を慎重に選択してください。 OS リストで、「Raspberry Pi OS (other)」を展開し、リストの一番上にある通常の「Raspberry Pi OS (64 ビット)」(デスクトップ エディション) ではなく、Raspberry Pi OS Lite (64 ビット) を選択します。デスクトップ版には、ログインや侵入のたびに ~/.asoundrc を積極的に書き換える Wayfire パネルが同梱されています。

明らかな原因もなく、オーディオ設定を繰り返し実行します。 Lite には GUI もパネルもありません。
間違ってデスクトップをすでにインストールしてしまいましたか?これは再フラッシュせずに修正可能です — sudo systemctl set-default multi-user.target && sudo systemctl disable lightdm.service && sudo reboot により、デスクトップが Lite のように起動します。ただし、最初から Lite の方がクリーンです。
OS のカスタマイズについて尋ねられたら、「はい」と答えてください。これは魔法のステップです。
ユーザー名 + パスワード: 選択して覚えておいてください
WiFi: ネットワーク名 + パスワード、正確に入力してください (ここでのタイプミス = 謎のノーショー)
「サービス」タブ: SSH (パスワード認証) を有効にする
書き込み、待機、イジェクト。 (「Raspberry Pi Connect」プロンプト: スキップしてください。これをどこかでデモした場合は、後で 1 つのコマンドでリモート アクセスを追加できます。)
MicroSDをPiの下側のスロットに差し込みます。 USB-C 電源を接続します。電源ボタンはありません。プラグイン = オン。
2 ～ 3 分待ちます (最初の起動でサイズが変更され、WiFi に接続されます)。
コンピュータの端末から:
ssh あなたのユーザー名@billy.local
指紋プロンプトで「yes」（小文字）と入力し、次にパスワードを入力します（パスワードの入力中にカーソルは移動しません - 通常）。
指紋認証プロンプトの直後に接続がリセットされる → Pi はまだ初回起動セットアップを完了中です。 1 ～ 2 分待ってから、もう一度試してください。 (リモート ホスト ID が変更されましたと表示される場合: ssh-keygen -R billy.local を実行して再試行します。Pi はその ID を再生成しました。

[切り捨てられた]

## Original Extract

Turn a Big Mouth Billy Bass into a real-time voice assistant powered by Strands BidiAgent + Amazon Nova 2 Sonic on a Raspberry Pi 5 - morganwilliscloud/billai-bass

GitHub - morganwilliscloud/billai-bass: Turn a Big Mouth Billy Bass into a real-time voice assistant powered by Strands BidiAgent + Amazon Nova 2 Sonic on a Raspberry Pi 5 · GitHub
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
morganwilliscloud
/
billai-bass
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits iot-identity iot-identity .gitignore .gitignore README.md README.md asoundrc.example asoundrc.example billy.py billy.py motors.py motors.py requirements-frozen.txt requirements-frozen.txt View all files Repository files navigation
Powered by Strands Agents Bidirectional Streaming + Amazon Nova 2 Sonic
billai-bass-demo-strands.mp4
Turn a Big Mouth Billy Bass into a real-time voice assistant: you talk, the fish talks back — head swiveling, mouth lip-syncing to its own voice, tail flapping for emphasis. The fish runs a Strands Agents bidirectional streaming agent ( BidiAgent ) on a Raspberry Pi 5, streaming live audio to and from Amazon Nova 2 Sonic on Amazon Bedrock.
No robotics experience needed. No prior soldering experience needed. The person this guide is based on had never plugged in a Raspberry Pi before and got a talking fish in a weekend.
Inspired by the excellent billy-b-assistant project, rebuilt on Strands.
File
What it is
README.md
This guide — start here, work top to bottom
billy.py
The final working Python — the talking fish
motors.py
Standalone motor test rig — run this before billy.py to verify wiring
asoundrc.example
The ALSA config that pins the USB mic + speaker as defaults
requirements-frozen.txt
Exact dependency versions from a known-good Pi build — pip install -r for guaranteed-working setup
iot-identity/
Optional production-grade credentials: CloudFormation template + walkthrough for X.509-cert auth instead of access keys
How to use this with an AI assistant
Paste this README into Claude (or another AI assistant) and say: "I'm building this. Walk me through it one step at a time — wait for me to confirm each step before giving the next." That's the way it was built; that's the way it builds best. See the Build this WITH an AI Assistant section below for what to send when stuck.
🤝 Build this WITH an AI Assistant
This project was created with help from Claude Code and you can use any AI assistant to help you get through.
Paste this entire guide into your AI and say: "I'm building this. Walk me through it one step at a time — wait for me to confirm each step before giving the next."
Your AI will pace the build to you, debug your exact error messages, and adapt when your hardware differs slightly from this guide. Things AI is great at during this build:
Paste your exact error message — the scary 50-line Python traceback usually has one meaningful line, and AI will find it.
Send photos — stuck identifying a motor? Can't tell which pin is which? Photograph it and ask. (In Claude Code on a Mac: drop the image file into the terminal, or say "look at ~/Downloads/IMG_1234.jpg".)
Say "I know nothing about X" — AI will back up and explain. There are no dumb questions.
Ask before cutting/soldering anything you're unsure about.
If you hit something this guide doesn't cover, that's not a failure of you. It's probably a gap in the guide. AI + your error message will get you through.
If you got a kit: you have everything in this table already. Building from scratch: this is the original Amazon order (~$240 total, January 2026 prices).
A personal computer (Mac/Windows/Linux) to flash the SD card — ⚠️ corporate laptops often block USB storage; use a home machine
A way to plug a microSD into that computer (most cards include an SD adapter)
An AWS account (personal, not your employer's) with a payment method — you'll create a locked-down, fish-only IAM user in the prerequisites below
Tape (electrical ideally, Scotch works), a small Phillips screwdriver
Home WiFi that both your computer and the Pi can join
✅ Prerequisites (do these before build day)
1. AWS / Amazon Bedrock setup (least privilege — do it right)
💸 No AWS account yet? The free tier covers this project. New AWS accounts get $100 in credits (up to $200 with activities), and Amazon Bedrock is usable on the free plan with credits applied to Nova usage. Voice chat is cheap — roughly a penny per few minutes of conversation — so credits buy on the order of 100+ hours of fish banter. Two free-plan caveats:
Brand-new accounts sometimes start with Bedrock invocation quotas set to zero (no payment history yet). If model access is granted but every call fails with ValidationException: Operation not allowed , that's the quota — file a (free) AWS support case asking for Nova Sonic on-demand quota, or upgrade to the paid plan (your remaining credits carry over).
Free-plan accounts auto-close after 6 months (or when credits run out). Fine for a demo fish; if Billy's becoming a permanent housemate, upgrade to paid before the deadline — unused credits survive the upgrade.
Your AWS keys will live in a plaintext file on a Pi inside a fish. Fish get demoed, lent out, and left on desks; SD cards pop out. So the keys must be able to do exactly one thing : talk to Nova Sonic. If they leak, the worst case is someone chats with a fish on your dime — not someone mining crypto in your account.
In the AWS console, switch region to us-east-1 (N. Virginia) — Nova Sonic lives there.
Go to Amazon Bedrock → Model access and request access to Nova 2 Sonic ( amazon.nova-2-sonic-v1:0 ). Usually instant. (Not v1 — original Nova Sonic hits end-of-life September 2026.)
Step 2 — Create a dedicated IAM user:
Console → IAM → Users → Create user . Name it something honest like billy-bass .
Do NOT check "Provide user access to the AWS Management Console" — this user is a robot fish; it never logs into anything.
On the permissions screen, choose "Attach policies directly" but don't select any of the AWS-managed policies (no AmazonBedrockFullAccess — that grants far more than the fish needs). Just click through and create the user with zero permissions.
Step 3 — Attach a policy that allows only Nova Sonic:
Open the new user → Permissions → Add permissions → Create inline policy → JSON tab.
Paste:
{
"Version" : " 2012-10-17 " ,
"Statement" : [
{
"Sid" : " BillyTalksToNovaSonicOnly " ,
"Effect" : " Allow " ,
"Action" : " bedrock:InvokeModelWithBidirectionalStream " ,
"Resource" : " arn:aws:bedrock:us-east-1::foundation-model/amazon.nova-2-sonic-v1:0 "
}
]
}
Name it billy-nova-sonic-only , save.
That's one action on one model in one region. No S3, no EC2, no other Bedrock models, no ability to see or change anything else in your account. (If you later give Billy tools that touch AWS — say, reading a DynamoDB table — add that specific permission to this policy then, not broad access now.)
Step 4 — Create the access key:
User → Security credentials → Create access key → choose "Application running outside AWS".
Copy the access key ID + secret somewhere safe — the secret is shown exactly once. These go on the Pi later (§1.6).
Step 5 — Optional but smart: set a billing alarm. Console → Billing → Budgets → create a $10/month budget with an email alert. Casual Nova Sonic chat costs cents; the alarm is there so a stuck always-listening fish (or leaked key) can't surprise you.
If the keys ever leak or a kit fish goes missing: IAM → user billy-bass → Security credentials → deactivate the access key . One click, fish bricked, account safe.
2. Know the two paste gotchas (you WILL hit these)
When you paste code into a terminal/nano over SSH, two things love to go wrong:
Extra indentation : every pasted line gains leading spaces. ALSA config shrugs; Python explodes ( IndentationError: unexpected indent on line 1 is the signature). Fix: sed -i 's/^ //' yourfile.py , or re-paste into nano carefully.
Dropped/mangled characters : a missing comma, or quotes turned into “smart quotes.”
The habit that saves you : after creating any Python file, run
python -m py_compile yourfile.py && echo OK
before running it for real. If it doesn't say OK, paste the error to Claude.
Part 1 — The Brain (Raspberry Pi, no fish yet)
Strategy note: we get the entire voice assistant working on the bare Pi first, and only then open the fish. Software problems and hardware problems are much easier to debug when they can't be each other.
Download Raspberry Pi Imager from raspberrypi.com/software , insert the microSD.
Choose: Device = Raspberry Pi 5 · OS = Raspberry Pi OS Lite (64-bit) · Storage = your card.
⚠️ OS version matters : you need an image with Python 3.12+ (Nova Sonic requirement). The current "Trixie"-based Raspberry Pi OS ships 3.13 ✅. Old "Bookworm" images ship 3.11 ❌.
"Lite" = no desktop. Correct — you'll never plug in a monitor.
⚠️ Pick "Lite" carefully. In the OS list, expand "Raspberry Pi OS (other)" and choose Raspberry Pi OS Lite (64-bit) — not the regular "Raspberry Pi OS (64-bit)" at the top of the list, which is the Desktop edition. The Desktop edition ships a Wayfire panel that aggressively rewrites ~/.asoundrc on every login, breaking your audio config repeatedly with no obvious cause. Lite has no GUI and no panel.
Already installed Desktop by mistake? It's fixable without re-flashing — sudo systemctl set-default multi-user.target && sudo systemctl disable lightdm.service && sudo reboot makes Desktop boot like Lite. But Lite from the start is cleaner.
When asked about OS customisation, say yes — this is the magic step:
Username + password: pick and remember them
WiFi: your network name + password, exactly right (a typo here = mystery no-show later)
Services tab: enable SSH (password authentication)
Write, wait, eject. ("Raspberry Pi Connect" prompt: skip it — you can add remote access later with one command if you ever demo this somewhere.)
MicroSD into the slot on the Pi's underside. Plug in the USB-C power supply — there is no power button; plugging in = on.
Wait 2–3 minutes (first boot resizes things and joins WiFi).
From your computer's terminal:
ssh yourusername@billy.local
Type yes (lowercase) at the fingerprint prompt, then your password (the cursor won't move while typing a password — normal).
Connection reset right after the fingerprint prompt → the Pi is still finishing first-boot setup. Wait 1–2 minutes, try again. (If it then complains REMOTE HOST IDENTIFICATION HAS CHANGED : run ssh-keygen -R billy.local and retry — the Pi regenerated its

[truncated]
