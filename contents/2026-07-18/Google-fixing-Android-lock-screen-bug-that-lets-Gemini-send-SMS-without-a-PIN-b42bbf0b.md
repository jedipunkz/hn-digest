---
source: "https://www.theregister.com/security/2026/07/17/google-fixing-android-lock-screen-bug-that-lets-gemini-send-sms-without-a-pin/5273027"
hn_url: "https://news.ycombinator.com/item?id=48959272"
title: "Google fixing Android lock screen bug that lets Gemini send SMS without a PIN"
article_title: "Google fixing Android lock screen bug that lets Gemini send SMS without a PIN"
author: "Bender"
captured_at: "2026-07-18T16:44:48Z"
capture_tool: "hn-digest"
hn_id: 48959272
score: 2
comments: 0
posted_at: "2026-07-18T15:59:23Z"
tags:
  - hacker-news
  - translated
---

# Google fixing Android lock screen bug that lets Gemini send SMS without a PIN

- HN: [48959272](https://news.ycombinator.com/item?id=48959272)
- Source: [www.theregister.com](https://www.theregister.com/security/2026/07/17/google-fixing-android-lock-screen-bug-that-lets-gemini-send-sms-without-a-pin/5273027)
- Score: 2
- Comments: 0
- Posted: 2026-07-18T15:59:23Z

## Translation

タイトル: Google、Gemini が PIN なしで SMS を送信できる Android ロック画面のバグを修正
説明: 特定のマルチタッチ ジェスチャは認証プロンプトをバイパスし、誰でもメッセージを送信できるようにします。

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
Google、Gemini が PIN なしで SMS を送信できる Android ロック画面のバグを修正
特定のマルチタッチ ジェスチャは認証プロンプトをバイパスし、誰でもメッセージを送信できるようにします。
これをイメージしてください。誰かがあなたの Android スマートフォンを手に入れ、あなたの PIN を知らないにもかかわらず、ロック画面から Gemini を使用してあなたとして SMS または WhatsApp メッセージを送信することができます。これは本物のバグであり、Google によると、今週中に修正が行われる予定です。
5 月以来、The Register は、ロック画面からの Gemini アクセスを可能にする Android 16 デバイス上でユーザーがデバイス認証をバイパスしているという複数の報告を受け取りました。
これらは、2025 年 9 月以降に広まった同様の Gemini ベースの Android ロック画面バイパス バグとは異なります。
私たちに報告されたバグの 1 つは、Android デバイスに物理的にアクセスできる認証されていないユーザーが、特定のマルチタッチ ジェスチャを使用して、ロック画面上で Gemini 経由で電話、テキスト、WhatsApp などの機能を有効にすることができたものです。
デバイス所有者がメッセージなどの特定のアプリへの Gemini のアクセスを取り消し、その後誰かがロック画面で Gemini 経由で SMS を送信しようとすると、チャットボットはユーザーに関連するアプリを開くように促します。 「続行」を選択すると、メッセージにアクセスするための正しい PIN を入力するようユーザーに求められます。
ただし、Gemini の [添付ファイルを追加] ボタンと同時に [続行] を押すと、デバイスは認証されていないユーザーが PIN を入力せずに Gemini 経由で SMS を送信できるようになります。
そこから、ユーザーは次のコマンドを呼び出すことで、以前はユーザーの設定で Gemini から切断されていた他のアプリへの Gemini のアクセスを有効にすることができます。

関連するプロンプトを表示します。
たとえば、Gemini に WhatsApp へのアクセスを許可するには、ユーザーは Gemini テキスト ウィンドウに「@WhatsApp」と入力します。 PIN は必要ありません。
正しい PIN を入力した後、ユーザー設定に戻ると、これが機能していることを確認できます。期待された認証手順を完了せずに WhatsApp が Gemini に接続されていることが表示されます。
この欠陥を悪用するには、デバイスへの物理的なアクセスが必要です。現実のシナリオでは実行が困難な場合が多いため、ほとんどの場合、この種の脆弱性に対してあまりにも多くの通信時間を与えることは避けられます。
サイドローディング Android アプリの不正開発者、Google の検証体制からの解放を懇願
Android キーボードはキーを完全に廃止し、ユーザーの意図を予測します
AppleのアップデートはロックアウトされたiPhoneユーザーのチェコ人仲間のようだ
Googleは悪者と戦うためにさらに多くのAIセキュリティエージェントを解放します
たとえば、Windows に、何らかの理由で攻撃者が Windows マシンに接続されているキーボードに物理的にアクセスする必要がある、私を管理者にするバグがあった場合、そのマシンの所有者は、脆弱性自体よりも大きな問題を抱えていることになります。
しかし、特に英国における携帯電話盗難犯罪の現状と、可能性の 1 つを挙げると、偽の誘拐詐欺の一環として説得力のある SMS メッセージが送信される可能性を考慮すると、この件はある程度の注目に値すると考えられます。
Googleの広報担当者は、これは既知のバグであり、今週完全な展開が予定されている修正をすでに実装していると述べた。
また、サムスンの端末では再現できないとの声もあったが、どのメーカー、モデル、バージョンが脆弱なのかを特定するには至らなかったため、このバグはPixel固有のものではないとも述べた。 ®
ソフトウェア
Mozilla、Firefox のリリーススケジュールを隔週に短縮
そして次の ESR リリースが近づくにつれて Firefox 153 に何が期待されるか

請求ソフトウェアのエラーにより AWS に数十億ドルの見積もりが送信される
Amazon、バグ修正に向けてパニックにならないようユーザーに呼びかけ
Gobi X: AI のためのエネルギーを社会から奪うのではなく、より多くのエネルギーを生み出す
パートナーコンテンツ: エンビジョンは、その逆ではなく、コンピューティングが豊富な砂漠の電力を追い求めることで、データセンターの戦略をどのように逆転させているか
AI スパム フィルターは、昔ながらのテキスト ソルティングに悩まされています
数十年前の電子メールのトリックが、一部の LLM を利用した電子メール フィルターに対して依然として機能することが判明
PaaS + IaaS
アマゾン ウェブ サービスの最も声高な顧客は現在 EC2 を実行しています
小売財団のリーダー、デイブ・トレッドウェル氏が上級リーダーに就任し、19年間の退役軍人であるデイブ・ブラウン氏が未知の牧草地へ出発
Microsoft、管理者に Exchange Online の休憩スペースを提供
PowerShell の廃止 - Credential パラメーターは 2026 年末に延期
OSプラットフォーム
Torvalds 氏は、Linux をフォークするよう嫌いな人々に挑戦しました。誰かが「ビールを我慢して」と言った
aiとml
ライナス・トーバルズ氏、AI嫌いの人たちにフォークをやめるよう指示
AI と ML
研究者が無差別級 AI モデルを 100 ドル未満で毒殺
AIとml
Grok Buildがリポジトリ全体をクラウドに送信していたことが発覚した後、マスク氏は粛清を約束
PAAS と IAAS
アマゾン ウェブ サービスの最も声高な顧客は現在 EC2 を実行しています
データのパージは、新興企業が回避しようと取り組んでいる「不整合な動作」の一例とみなされる
モデルは検証を提供せずに信頼を要求します
プレスリリースを掲載しているファブメーカーに注意してください
Thinking Machines の最初のオープンウェイト モデルは、中国の LLM に代わる 9,750 億パラメータです
ワンツーパンチは、低精度 AI が高精度シミュレーションをどのように補完できるかを垣間見せます
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
投票村のレポートは非常に成功しているとジェフ・モス氏は言う

s、DEF CON全体が含まれるようになりました
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
x86-32 での Debian の最終リリースに黙祷を捧げてください
新しい Debian バージョンが 13.6 および 12.15 の形で FOSSland に登場
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
KDE Plasma ユーザーは恐ろしい変化の予兆に直面しています: 6.6.6 の登場
現在は 6.7 が最新版であり、6.8 では好むと好まざるにかかわらず Wayland を入手できるようになります。
FOSS クラウドオフィススイート間の競争が激化する中、Collabora が CODE 26.04 をリリース
Markdown のサポートと、よりスマートな数式エラー処理に加え、AI が統合されました (デフォルトではオフになっています)。
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
状況出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

A specific multi-touch gesture bypasses an authentication prompt, allowing anyone to send messages

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Google fixing Android lock screen bug that lets Gemini send SMS without a PIN
A specific multi-touch gesture bypasses an authentication prompt, allowing anyone to send messages
Picture this. Someone gets hold of your Android phone and, despite not knowing your PIN, they can use Gemini from the lock screen to send SMS or WhatsApp messages as you. This is a real bug and Google says a fix is coming as soon as this week.
Since May, The Register has received multiple reports of users bypassing device authentication on Android 16 devices that enable Gemini access from the lock screen.
These are distinct from the similar Gemini-based Android lock screen bypass bugs that have made the rounds since September 2025.
One of the bugs reported to us allowed unauthenticated users with physical access to an Android device to enable functionality such as phone, texts, and WhatsApp via Gemini on the lock screen using a specific multi-touch gesture.
When device owners revoke Gemini’s access to certain apps, like Messages, and someone later tries to send an SMS via Gemini on the lock screen, the chatbot will prompt the user to open the relevant app. Selecting "Continue" prompts the user to enter the correct PIN to access messages.
However, when “Continue” is pressed simultaneously with Gemini’s “Add attachment” button, the device will then allow unauthenticated users to send that SMS via Gemini, without needing to enter a PIN.
From there, users can enable Gemini’s access to other apps, which were previously disconnected from Gemini in the user’s Settings, by invoking the relevant prompt.
For example, to allow Gemini access to WhatsApp , users can enter “@WhatsApp” in the Gemini text window. No PIN needed.
You can then check this has worked by going back into user Settings, after entering a correct PIN, and it will show that WhatsApp is connected to Gemini without completing the expected authentication step.
Exploiting the flaw requires physical access to a device. In most circumstances, we steer clear of giving this type of vulnerability too much airtime since it is often difficult to pull off in real-world scenarios.
Rogue devs of sideloaded Android apps beg for freedom from Google’s verification regime
Android keyboard ditches keys entirely, predicts what you mean
Apple update looks like Czech mate for locked-out iPhone user
Google unleashes even more AI security agents to fight the baddies
If Windows, for example, had a make-me-admin bug that required the attacker, for whatever reason, to have physical access to the keyboard connected to the Windows machine, then the owner of said machine has bigger problems than the vulnerability itself.
However, given the state of phone theft crime, especially in the UK , and the potential to send convincing SMS messages as part of fake kidnapping scams , to name one possibility, we think this one deserves some attention.
A Google spokesperson told us this is a known bug and it has already implemented a fix that was scheduled for a full deployment this week.
They also said the bug is not Pixel-specific, after some claimed that they could not reproduce it on Samsung devices, but fell short of specifying which manufacturers, models, or versions are vulnerable. ®
SOFTWARE
Mozilla speeds Firefox release schedule to biweekly
And what to expect in Firefox 153 as the next ESR release gets close
Billing software error sends billion-dollar AWS estimates
Amazon asks users not to panic as it works to fix the bug
Gobi X: Creating more energy for AI, not taking it from society
PARTNER CONTENT: How Envision is reversing the datacenter playbook by making computing chase abundant desert power, not the other way around
AI spam filters are getting suckered by old-school text salting
Turns out decades-old email tricks still work against some LLM-powered email filters
PaaS + IaaS
Amazon Web Services' most vocal customer now runs EC2
Retail foundation leader Dave Treadwell takes over as senior leader and 19-year vet Dave Brown departs for pastures unknown
Microsoft gives admins Exchange Online breathing room
Retirement of PowerShell -Credential parameter pushed back to the end of 2026
OS PLATFORMS
Torvalds challenged the haters to fork Linux. Someone said 'hold my beer'
ai and ml
Linus Torvalds tells AI haters to fork off
AI and ML
Researcher poisons open-weight AI model for under $100
AI and ml
Musk promises purge after Grok Build caught sending entire repos to the cloud
PAAS AND IAAS
Amazon Web Services' most vocal customer now runs EC2
Data purges deemed an example of 'misaligned behavior' that upstart is working to avoid
Models demand trust without offering verification
Beware of fab makers bearing press releases
Thinking Machines' first open weights model is a 975 billion parameter alternative to Chinese LLMs
One-two punch offers a glimpse of how low-precision AI can complement high-precision simulations
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
A moment of silence, please, for the final release of Debian on x86-32
New Debian versions hit FOSSland in the form of 13.6 and 12.15
Baddies caught exploiting extensions bugs with perfect 10 scores on vulnerable Joomla websites
Flaws in iCagenda, Balbooa Forms extensions can impact open source CMS that powers a million sites worldwide
Frame: A new X11 server – implemented directly in assembly
Joins yserver, Phoenix, and of course XLibre – and outlier Arcan
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
KDE Plasma users face a dire omen of change: 6.6.6 arrives
6.7 is now current, and in 6.8 you're getting Wayland whether you like it or not
Collabora releases CODE 26.04 as rivalry between FOSS cloudy office suites heats up
Now with Markdown support and smarter formula error handling – plus integrated AI, though it's off by default
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
