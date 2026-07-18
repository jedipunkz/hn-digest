---
source: "https://github.com/adgapar/teya"
hn_url: "https://news.ycombinator.com/item?id=48958769"
title: "Show HN: Teya – Open‑source AI family agent on a wall‑mounted Android phone"
article_title: "GitHub - adgapar/teya: personal ai agent at home · GitHub"
author: "AdiletGaparov"
captured_at: "2026-07-18T15:45:13Z"
capture_tool: "hn-digest"
hn_id: 48958769
score: 1
comments: 0
posted_at: "2026-07-18T15:04:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Teya – Open‑source AI family agent on a wall‑mounted Android phone

- HN: [48958769](https://news.ycombinator.com/item?id=48958769)
- Source: [github.com](https://github.com/adgapar/teya)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T15:04:12Z

## Translation

タイトル: 表示 HN: Teya – 壁掛け Android フォン上のオープンソース AI ファミリー エージェント
記事タイトル: GitHub - adgapar/teya: 自宅のパーソナル AI エージェント · GitHub
説明: 自宅のパーソナル AI エージェント。 GitHub でアカウントを作成して、adgapar/teya の開発に貢献してください。

記事本文:
GitHub - adgapar/teya: 自宅のパーソナル AI エージェント · GitHub
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
アドガパール
/
てや
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く

フォルダーとファイル
161 コミット 161 コミット .claude/ skill/ add-visualization .claude/ skill/ add-visualization .github/ workflows .github/ workflows app app docs docs gradle gradle thought/共有思考/shared .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSEライセンス README.md README.md THIRD_PARTY_MODELS.md THIRD_PARTY_MODELS.md build.gradle.kts build.gradle.kts gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat settings.gradle.kts settings.gradle.kts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
家庭の家族代理人 — 耳を傾け、理解し、記憶し、行動する、温かく知的な存在。予備の Android スマートフォン上に構築されました。
家はあなたのことを知っているので、あなたの面倒を見てくれていると感じるべきです。
teya-demo-compressed.mp4
すべてを忘れる機械
アレクサにタイマーをセットしてもらうと、それが機能します。誰が子供たちをサッカーに迎えに行くのかと尋ねても、そこには何もありません。カレンダーが 2 つあるという概念もなく、毎週火曜日にこの話題が来るという記憶もなく、質問が理解できたとしても答えに基づいて行動する方法もありません。音声アシスタントは、シンプルなコマンド、つまり派手な if-then システムを中心に 2010 年代初頭に設計されました。彼らはコンテキストを理解せず、セッション間ではすべてを忘れ、アプリ間で有意義な方法で行動することができません。家族はタイマーをセットしたり音楽を再生したりするためにそれらを使用し、残りは無視します。
一方、ChatGPT と Claude の背後にある AI は真にインテリジェントです。ただし、チャット ウィンドウの背後にあり、実際の家庭の物理的およびデジタル インフラストラクチャからは切り離されています。同じインテリジェンスを、1 人のチャット ウィンドウではなく、世帯全体が共有するデバイスに配置すると、世帯運営を実際に支援するスマート ファミリー エージェント、つまり、話を聞き、理解し、記憶し、実行するエージェントが得られます。
Th

これを構築可能にするアイデアは、安価な Android フォンがスーパーチャージされた Arduino であり、すぐに消費者に提供できるようになるということです。すでにディスプレイ、4G/5G、マイクとスピーカー、GPS、カレンダーとタイマーを備えており、すでに構築されている何百万ものアプリに加えて、AI モデルを直接呼び出すことができます。
安価な電話機を壁に取り付けて差し込むと、目に見えないホーム ネットワークの単一の目に見えるノードになります。スマート電球、錠前、カレンダー、メッセージング アプリ、それらはすべて隠されたインフラストラクチャです。壁にある顔は、家族が見る必要がある唯一のインターフェースです。
その背後にサーバーやローカル エージェント スタックはありません。今日の AI エージェントの波は激しくなっており、人々はコーディング エージェントやオーケストレーション フレームワークをローカルでホストするために Mac Mini を購入しています。 Teya には何も必要ありません。インテリジェンスはクラウドに存在し、ターンごとに数回の API 呼び出しで到達できるため、ハードウェアはそれらの呼び出しを行い、画面を維持し、デバイスを駆動する安価な電話です。最新の Android スマートフォン、Android 8.0 (Oreo) 以降が動作します。
ただし、マイクとスピーカーはこのために作られたものではありません。これらは、腕が届く距離で近くで使用できるように調整されており、部屋の向こう側でウェイクワードや明瞭な音声を拾うことには、それ以上のことが求められます。初期のテストは有望ですが限界があり、より大きな部屋では外部マイク アレ​​イまたはスピーカーが必要になる場合があります。
Play ストアやアプリ ストアのレビューはありません。APK をリリースから直接取得してサイドロードします。
電話機で、ブラウザの「不明なアプリのインストール」を有効にします (初回試行時に設定のプロンプトが表示されます)。
上記のリリース ページから最新の .apk をダウンロードします。
ダウンロードしたファイルをタップしてインストールします。
Teya を開いてオンボーディングに従います。世帯の Mistral API キーと必要な権限が求められます。その後、ウェイク ワードの調整と個人ごとの音声登録が管理画面に反映されます (画面を長押しします)。
新しい v の場合、自動アップデーターはありません。

問題が解決したら、「リリース」に戻り、新しい APK をダウンロードして、もう一度タップします。同じキーで署名されている限り、最初にアンインストールする必要がなく、既存のアプリの上にインストールされます (そのため、家庭のデータ、設定、メモリは保持されます)。
買い物リスト — 「牛乳がなくなりました」は二度尋ねることなく追加されます。
カレンダー — 定期的かどうかにかかわらず、音声でイベントを追加すると、残りの家族を電子メールで自動的に招待します。
タイマーとアラーム — 設定しておくと、時間が来ると自分の声で大声でアナウンスします。
リマインダー — 「20 分以内に配管工に電話するようリマインドする」または「学ランにカップケーキを持っていくようにリマインドする」は、タイマーまたは静かなカレンダーのエントリとして、実際に適したものになります。
支出 — 「果物代 12 ユーロ」はその場で記録され、分類されます。 「今月いくら使いましたか？」と尋ねると、決して推測することなく、正確に数字が合計されます。
電話 - ダイヤラーを操作できない 5 歳児が話す「おばあちゃんに電話して」は、通常のハンズフリー携帯電話をかけることができますが、承認された家族の連絡先許可リストに登録されている相手にのみ通話できます。未知の番号、任意の番号、またはプレミアム番号にダイヤルする方法はありません。
最大の受益者は、家の中で精神的な負担を負っている人です。約束、食事、物流、学校の管理者など誰も追跡しません。 Teya は家庭の第 2 の頭脳となり、バックグラウンドで静かに動作します。今後の内容を含む完全なリストは docs/roadmap.md にあります。
設計によりロックダウンされています。箱入り家電は所定の位置に固定され、家族の個人的な Google、ソーシャル、銀行へのログインではなく、それを操作するためだけに作成された新しいアカウントで実行されます。ハイジャックしたり盗んだりする個人的なものは何もありません。通話許可リストと組み合わせることで、子供やゲストの手の届く壁に安全に置いておくことができます。
プライバシーも同様に機能します

つまり、アーキテクチャ自体に組み込まれています。世帯名簿、個人ごとのメモリ、連絡先の許可リストはデバイス上に保存されます。生の会話のトランスクリプトは書き込まれずに残り、ディスクに触れることがありません。ターンで考えたり話したりする必要があるもの、または毎晩の夢で記憶を強化するために必要なものだけがモデルに送られます。
あなたは、家族があなたを認識し、あなたについてのことを覚えていてくれることを期待しています。それが Teya の機能です。Teya は記憶し、誰が話しているのかを認識します。
何かを思い出してくださいと頼むと、「私がピーナッツアレルギーであることを覚えておいてください」は、具体的にあなたに関する事実として保存されます。また、それが 1 人に関するものではない場合は、その家族に関する事実として保存されます。忘れるように頼めば、その記憶は消えてしまいます。保存されたものはすべて手動で確認および編集することもできます。
しかし、テヤが覚えていることのほとんどは、誰もそれを保存するように頼んだものではありません。すべての会話の終わりに、モデルは実際に保持する価値のある内容を 1 つの短いメモに抽出するか、保持する価値がまったくないと判断します。 Teya は会話の記録を保存する代わりにそのメモを保存するため、デバイス上に話された内容が完全に記録されることはありません。
それは夢です。毎晩、家が眠っている間に、それらのばらばらのメモはその家庭に関する永続的な事実に統合され、補強されていないものは永遠に積み重なるのではなく、忘却曲線で消えていきます。それは人の記憶と同じように老化し、重要なものは保持され、残りは静かにされます。保管されているものは何でも、調べるように頼まれなくても、勝手に戻ってきます。
誰が話しているのかもわかります。世帯の各メンバーの声が一度登録されると、それ以降は、クラウド スピーカー ID サービスを必要とせず、デバイス上の小さなモデルがそれらのサンプルに対して誰が話しているのかを照合します。静かな試合は二人を区別するだけです。自信のある人は、テヤに次のように挨拶させます。

名前。これはウェイク ワードとバージイン検出と並行して実行され、どちらもデバイス上でも実行されるため、ほとんどのリスニングはクラウドに到達する前に電話で行われます。
画面全体がコントロールです。短くタップすると、「Hey Teya」と言うのと同じように彼女が目覚めます。長押しすると「管理」が開きます。
携帯電話は粒子の視覚化では水平方向に、顔 (2 つの目と口) では垂直方向に動作します。管理者に切り替えると (画面を長押しするか、Teya に尋ねてください。他の設定と同じように、彼女は方法を知っています)、選択はオンボーディングと管理者自身の背景にも反映されます。クロード コード スキル ( /add-visualization ) は 3 番目のスキルを追加する足場になります。
Teya のエージェント ループは、LLM、TTS、STT、ウェイク ワード、および VAD モデルを電話機自体に接続し、モデルが決定した内容を Android のネイティブ SDK に関連付けられたツール呼び出しに変換します。つまり、通話の発信、ショッピング リストの項目の追加、メモリの保存などです。
オンボーディングをシンプルに保ちたかったので、複数のプロバイダーを混在させるのではなく、1 つのプロバイダーにこだわりました。それは安価で多言語である必要がありました。私自身の家族は家で複数の言語を話すので、Mistral は非常に適しており、それが今日の推論 (Mistral Small) に加えて、音声合成および音声合成 (Voxtral) を実行しているものです。モデル自体は交換可能で、後でさらにプロバイダーを追加できます。オンボーディングでは、世帯が独自の API キーを入力します。
これらの背後にはバックエンド サーバーはありません。電話とマイク/スピーカーは Android ローカル API であるため、クラウド サーバーはいずれにせよアクセスできないため、家族データとデバイス制御はデフォルトでローカルに留まります。
このアプリはネイティブ Kotlin です。これは、通知の読み取り、他のアプリ内での動作、通話の発信、常時オンのフォアグラウンド サービスとしての実行など、Android 独自の API への深いアクセスが必要であるためです。ハードウェアのコストと、その種のバックグラウンドと権限アクセスに対する iOS 独自のロックダウンが、そのプラットフォームを支配しました

うーん、アウト。また、アプリは Play ストアや App Store を経由することはありません。そのため、クロスプラットフォームの本当のセールスポイントである 1 つのコードベースが両方のストアに出荷されるということは、ここではまったく関係ありませんでした。
電話がかかってくるとそのまま放置されます。子供が「お父さんに電話して」と言うと、アプリは通常の携帯電話にダイヤルして脇に進みます。会話に参加することはなく、一度ダイヤルすると通話の管理を停止します。 SIM 上の単純なセルラー通話: VoIP、オーディオ キャプチャ、ルートはありません。
オンデバイス (ローカル、ネットワーク往復なし): ウェイク ワード (自己トレーニングされた hey_teya モデル)、バージイン/VAD および文途中中断のためのエコー キャンセル、話者ごとの音声 ID、アニメーション化された画面上のプレゼンス、買い物リスト、支出ログ、および家族のすべての記憶、世帯名簿、エイリアス、および連絡先の許可リスト。
クラウド、ミストラル経由: 推論 (ツール使用ループ)、音声からテキストへの変換、およびテキストから音声への変換。これは、システム内で唯一必須のクラウド依存関係です。
電話 API: テレフォニー (発信通話)、カレンダーおよびアラーム/タイマー (同梱)、さらには、他のアプリの通知の読み取り、他のアプリ内での動作、BLE/Matter 上のスマートホーム制御も予定されています。
これらはどれも、紙の上では印象深いものである必要はありません。家の中の誰もそれについて考えないように、静かに動作する必要があります。
ドキュメント/ロードマップ。

[切り捨てられた]

## Original Extract

personal ai agent at home. Contribute to adgapar/teya development by creating an account on GitHub.

GitHub - adgapar/teya: personal ai agent at home · GitHub
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
adgapar
/
teya
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
161 Commits 161 Commits .claude/ skills/ add-visualization .claude/ skills/ add-visualization .github/ workflows .github/ workflows app app docs docs gradle gradle thoughts/ shared thoughts/ shared .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md THIRD_PARTY_MODELS.md THIRD_PARTY_MODELS.md build.gradle.kts build.gradle.kts gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat settings.gradle.kts settings.gradle.kts View all files Repository files navigation
A family agent for the home — a warm, intelligent presence that listens, understands, remembers, and does. Built on a spare Android phone.
A home should feel like it's looking after you, because it knows you.
teya-demo-compressed.mp4
a machine that forgets everything
Ask Alexa to set a timer and it works. Ask who's picking the kids up from football, and there's nothing there: no notion of two separate calendars, no memory that this comes up every Tuesday, no way to act on the answer even if it understood the question. Voice assistants were designed in the early 2010s around simple commands — fancy if-then systems. They don't understand context, they forget everything between sessions, and they can't act across apps in any meaningful way. Families use them to set timers and play music, then ignore the rest.
Meanwhile the AI behind ChatGPT and Claude is genuinely intelligent. It sits behind a chat window, though, disconnected from the physical and digital infrastructure of an actual home. Put that same intelligence on a device the whole household shares, instead of one person's chat window, and you get a smart family agent that actually helps run the household: an agent that listens, understands, remembers, and does .
The idea that makes this buildable at all: a cheap Android phone is a supercharged Arduino, consumer-ready out of the box. It already has a display, 4G/5G, a mic and speaker, GPS, a calendar and timer, and it can call an AI model directly, on top of the millions of apps already built for it.
Wall-mount any cheap phone, plug it in, and it becomes the single visible node of an invisible home network. Smart bulbs, locks, calendars, messaging apps: all of it hidden infrastructure. The face on the wall is the only interface the family needs to see.
There's no server behind it, no local agent stack. Today's AI-agent wave runs heavy — people buy Mac Minis to host coding agents and orchestration frameworks locally. Teya needs none of it: the intelligence lives in the cloud, reached with a few API calls per turn, so the hardware is a cheap phone that makes those calls, keeps the screen alive, and drives the device. Any modern Android phone, Android 8.0 (Oreo) or newer, works.
The mic and speaker weren't built for this, though. They're tuned for close, arm's-reach use, and picking up a wake word or clear speech across a room asks more of them than that. Early testing is encouraging but limited, and a bigger room may need an external mic array or speaker.
No Play Store, no app store review: grab the APK straight from Releases and sideload it.
On the phone, enable "install unknown apps" for your browser (Settings prompts you the first time you try).
Download the latest .apk from the Releases page above.
Tap the downloaded file to install.
Open Teya and follow onboarding — it asks for the household's Mistral API key and the permissions it needs. Wake word tuning and per-person voice enrollment live in Admin (long-press the screen) afterward.
There's no auto-updater — for a new version, come back to Releases, download the new APK, and tap it again. As long as it's signed with the same key, it installs over the existing app without needing an uninstall first (so household data, settings, and memory are kept).
Shopping list — "we're out of milk" gets added without being asked twice.
Calendar — add an event by voice, recurring or not, and it invites the rest of the household by email automatically.
Timers & alarms — set one, and it announces out loud, in its own voice, when time's up.
Reminders — "remind me to call the plumber in twenty minutes" or "remind me to bring cupcakes to the school run" becomes a timer or a quiet calendar entry, whichever actually fits.
Expenses — "12 euros for fruit" gets logged and categorized on the spot; ask "how much have we spent this month" and it adds the numbers up exactly, never guessing.
Calls — "Call Grandma," spoken by a five-year-old who can't navigate a dialer, places a normal, hands-free cellular call, but only to someone on an approved family contacts allowlist: no path to dialing an unknown, arbitrary, or premium number.
The biggest beneficiary is whoever in the house carries the mental load: the appointments, the meals, the logistics, the school admin nobody else tracks. Teya becomes a second brain for the household, running quietly in the background. The fuller list, including what's still ahead, lives in docs/roadmap.md .
It's locked down by design: a boxed home appliance fixed in place, running on fresh accounts created solely to operate it, never the family's personal Google, social, or banking logins. There's nothing personal on it to hijack or steal. Combined with the calling allowlist, that's what makes it safe to leave on a wall within reach of kids and guests.
Privacy works the same way: it's built into the architecture itself. The household roster, per-person memory, and the contacts allowlist live on the device. Raw conversation transcripts stay unwritten, never touching disk. Only what a turn needs to think and speak, or what the nightly dream needs to consolidate memories, goes to the model.
You expect a family member to recognize you and remember things about you. That's what Teya does: it remembers, and it recognizes who's talking.
Ask it to remember something, and it will: "remember I'm allergic to peanuts" gets saved as a fact about you specifically, or about the household if it isn't about one person. Ask it to forget, and that memory is gone. Everything it's saved can also be reviewed and edited by hand.
Most of what Teya remembers, though, nobody asked it to save. At the end of every conversation, the model distills what was actually worth keeping into one short note, or decides there's nothing worth keeping at all. Instead of saving a transcript of the conversation, Teya saves that note, so there's never a full record of what was said sitting on the device.
It dreams. Every night, while the house is asleep, those loose notes get consolidated into durable facts about the household, and anything that hasn't been reinforced fades on a forgetting curve instead of piling up forever. It ages the way a person's memory does, keeping what mattered and letting the rest go quiet. Whatever's been kept comes back on its own, without ever being asked to look it up.
It also knows who's talking. Each household member's voice gets enrolled once, and from then on a small on-device model matches who's speaking against those samples, no cloud speaker-ID service involved. A quiet match just tells two people apart; a confident one lets Teya greet someone by name. That runs alongside the wake word and barge-in detection, both on-device too, so most of the listening happens on the phone before anything reaches the cloud.
The whole screen is the control: a short tap wakes her, same as saying "Hey Teya"; a long press opens Admin.
The phone runs horizontal with the particles visualization, vertical with the face — two eyes and a mouth. Switched in Admin (long-press the screen, or just ask Teya — she knows how, same as any other setting), and the pick carries through onboarding and Admin's own background too. A Claude Code skill ( /add-visualization ) scaffolds adding a third.
The agent loop of Teya connects an LLM, TTS, STT, wake word, and VAD models to the phone itself, turning what the model decides into a tool call tied to Android's native SDK: placing a call, adding a shopping-list item, saving a memory.
I wanted onboarding to stay simple, so I stuck to one provider instead of mixing several. It needed to be cheap and multilingual — my own family speaks several languages at home, so Mistral was a very good fit, and it's what runs the reasoning (Mistral Small) plus the speech-to-text and text-to-speech (Voxtral) today. The model itself is swappable, more providers can be added later. Onboarding is where the household enters its own API key.
There's no backend server behind any of this: family data and device control stay local by default, since telephony and the mic/speaker are Android-local APIs a cloud server couldn't reach anyway.
The app is native Kotlin, because it needs deep access to Android's own APIs: reading notifications, acting inside other apps, placing calls, running as an always-on foreground service. Hardware cost and iOS's own lockdown on that kind of background and permission access ruled that platform out. And the app never goes through Play Store or the App Store, so cross-platform's real selling point, one codebase shipped to both stores, was never relevant here.
A call gets placed and then left alone: when a kid says "call Dad," the app dials a normal cellular call and steps aside. It never joins the conversation, and stops managing the call once dialed. A plain cellular call on the SIM: no VoIP, no audio capture, no root.
On-device (local, no network round-trip): the wake word (a self-trained hey_teya model), barge-in/VAD and echo cancellation for mid-sentence interruption, per-speaker voice ID, the animated on-screen presence, the shopping list, the expense log, and all of family memory, the household roster, aliases, and the contacts allowlist.
Cloud, via Mistral: reasoning (the tool-use loop), speech-to-text, and text-to-speech. This is the only mandatory cloud dependency in the system.
Phone APIs: telephony (outbound calls), calendar and alarms/timers (shipped), and, still ahead, reading other apps' notifications, acting inside other apps, and smart-home control over BLE/Matter.
None of this needs to be impressive on paper. It needs to work quietly enough that nobody in the house thinks about it.
docs/roadmap.

[truncated]
