---
source: "https://github.com/wrobelda/goodix-fp-spi-linux"
hn_url: "https://news.ycombinator.com/item?id=49209348"
title: "Show HN: I used LLM to reverse-engineer secure-enclaved fingerprint scanner"
article_title: "GitHub - wrobelda/goodix-fp-spi-linux: Driving Goodix SPI fingerprint sensors on Linux, including versions running in Qualcomm QSEE enclave. · GitHub"
author: "wrobelda"
captured_at: "2026-08-07T12:40:35Z"
capture_tool: "hn-digest"
hn_id: 49209348
score: 1
comments: 0
posted_at: "2026-08-07T12:26:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I used LLM to reverse-engineer secure-enclaved fingerprint scanner

- HN: [49209348](https://news.ycombinator.com/item?id=49209348)
- Source: [github.com](https://github.com/wrobelda/goodix-fp-spi-linux)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T12:26:07Z

## Translation

タイトル: HN を表示: LLM を使用して、セキュアにエンクレーブされた指紋スキャナーをリバース エンジニアリングしました
記事のタイトル: GitHub - wrobelda/goodix-fp-spi-linux: Linux で Goodix SPI 指紋センサーを駆動する (Qualcomm QSEE エンクレーブで実行されているバージョンを含む)。 · GitHub
説明: Linux 上で Goodix SPI 指紋センサーを駆動します (Qualcomm QSEE エンクレーブで実行されているバージョンを含む)。 - GitHub - wrobelda/goodix-fp-spi-linux: Qualcomm QSEE エンクレーブで実行されているバージョンを含む、Linux 上で Goodix SPI 指紋センサーを駆動します。
HN テキスト: 私は Xiaomi Mi Pad 5 Pro 5G Android タブレットを所有していますが、これはほぼ完璧な Linux タブレットでもあります。5G サポート、素晴らしい画面、素晴らしいフォリオ キーボード、8 つのスピーカー、ペンがあり、軽量で比較的強力です。現時点ではすべてのハードウェアがサポートされていますが、まだすべてがメインライン化されているわけではありません。また、Goodix SPI 指紋スキャナーも備えていますが、スキャナーを処理する Qualcomm の QSEE セキュア エンクレーブ内に存在するアプリを使用して実装されており、Android は文書化されていない HAL およびプロトコルを使用して通信します。この 2021 年の記事では、この問題について詳しく説明しています: https://emainline.gitlab.io/2021/12/12/fingerprint_P1.html Opus 5 と Kim K3 を使用したリバース エンジニアリングの集中的な作業に約 1 週間かかり、ドキュメントの磨き上げにさらに 5 日間かかりましたが、これで完了したと言っても過言ではありません。 Fable はセキュリティのテーマなので当然ドロップアウトしたため、Opus 5 を使用し、コードのレビューには Kimi K3 を使用しました。私は、途中の問題を解明するために、Frida を使用してライブ Android セッションを逆アセンブルし、トレースすることにしました。ご質問がございましたら喜んでお答えいたします。また、Macbook 上の安全に保護された指紋スキャナーのリバース エンジニアリングにも同じアプローチが使用できると思います。私は M1 で Asahi を実行しているので、次回の指紋スキャナーの実行時に検討してみます。

通常の仕事からの休憩。

記事本文:
GitHub - wrobelda/goodix-fp-spi-linux: Qualcomm QSEE エンクレーブで実行されているバージョンを含む、Linux 上で Goodix SPI 指紋センサーを駆動します。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ワロベルダ
/
Goodix-fp-spi-linux
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github/ workflows .github/ workflows docs docs ハーネス ハーネス トレース トレース .gitignore .gitignore ライセンス ライセンスを読む

ME.md README.md _typos.toml _typos.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Goodix SPI 指紋センサー
Goodix SPI 指紋センサー — ラップトップによくある無関係な種類の USB ではありません —
2つの方法で駆動できます。
「従来の」ものでは、リッチ実行環境 (REE)、オペレーティング システムの
カーネル ドライバーは SPI バス プロトコルを介してすべてのハードウェアを処理し、ユーザー空間で
センサーに直接話しかけてイメージングを実行したり、
イメージング、テンプレート、マッチング ロジック。
ただし、これは重大なセキュリティ リスクを伴うため、現在ではかなりまれです。
画像とテンプレートはオペレーティング システムによって処理されるため、
カーネルまたはユーザー空間スタックを侵害するものはすべて読み取ることができます。
それらを置き換えます。
別の方法は、信頼された実行でフィンガープリント ロジックを実行することです。
環境 (TEE)。オペレーティング システムがイメージングや
まったく一致しない — 代わりに、信頼された、独自の、署名された、
SoC の安全な世界で実行されるクローズド ソース アプリケーション。カーネルドライバー
必要最小限のハードウェアのみを処理する必要があり、
スキャナが接続されている SPI バス。
Goodix スキャナーに関して言えば、そのような TEE の 2 つの既知の実装
アプリケーションが存在し、1 つは MediaTek の MicroTrust TEE 用、もう 1 つは Qualcomm の
QSEE であり、デバイスに付属のオペレーティング システムに同梱されています。
Linux の TEE モードで Goodix をサポートする
TEE モードは、ユーザーにとって 2 つのモードよりもはるかに安全ですが、かなり安全でもあります。
通常は Android デバイスの OEM がハードウェアを維持することを保証します。
Abstraction Layer (HAL) ライブラリは閉じられ、そのライブラリが
オペレーティング システムの権限サブシステム間の通信全体
そして信頼できるAP

ひだ。唯一開いている部分はカーネルセンサーです
ドライバー。スキャナーを起動する最低限の機能を実行します。
残念ながら、これは Linux でこれらのスキャナのサポートを追加する唯一の方法であることを意味します
これは、Android 実装のリバース エンジニアリングによるものです。 2021年の記事、
タキシードの Android スマートフォンに指紋センサーを搭載: 不可能ですか?
(アーカイブ済み)、
ポートが何をしなければならないのか、そしてなぜ誰もそれをしなかったのかをよく説明しています。
このプロジェクトは、Claude Opus 5 と Kim K3 を使用して 10 日間にわたって開発されました。
Claude Fable 5 はおそらく処理を少しスピードアップするでしょうが、セキュリティに関連するものは何でも
Anthropic によってフラグが立てられ、モデルは定期的に Opus 5 に落とされていました。
したがって、ここでは Kim K3 を使用します。
これは未検証の AI の失敗ではありません。アーキテクチャの決定とコード レイアウトのほとんどが問題です。
分解/リバースエンジニアリングのアプローチとそれを駆動するだけでなく、自分で作成したものです。
ここでの作業は私のハードウェアでテストされており、さらなる改良を期待して、より多くのハードウェアでテストしています
デバイスをインストールし、Linux カーネルへのメインラインの準備を整えます。
Qualcomm SoC 上のこれらのスキャナーの TEE サポートは完了しました。
このプロジェクトは、完全なユーザー空間の実装ではなく、概念実証として機能します。このプロジェクトの内容
提供するのはプロトコル、ハードウェア側と TEE 側でプロトコルに到達するために必要な 2 つのカーネル ドライバー、および
実際のベースとなるシーケンスを示すリファレンス クライアント
fprintd ドライバー。
REE モード — オペレーティング システム自体が SPI 経由でイメージングを実行します。それ
移植することでセンサードライバーに追加可能
MediaTek-lineage gf_spi_tee.c の SPI 転送および画像スキャン コード (背後にあります)
SUPPORT_REE_SPI
クアルコムの TEE モード — 信頼できるアプリケーションに到達する
QSEECOM;ここに記載されているものはすべてこのモードで実行されました。
信頼できるアプリケーションをロードして通信する

エーション
再起動後の持続性 —
QSEE で保護されたオブジェクト。リスナー サービスを通じて保存されます。
MediaTek の TEE モード — プロトコルとセンサー ドライバーは、
終わった;欠けている半分は、 drivers/tee/ の MicroTrust バックエンドです。
MediaTek プラットフォームは、OP-TEE を通じて上流の TEE に到達します。必要なもの
リバースエンジニアリング
信頼できるアプリケーションのプロトコル —
Qualcomm のものではなく Goodix の汎用品であるため、これらは次の製品に引き継がれる必要があります。
メディアテックポート:
登録、最初の指
認証 (照合および拒否)
校正 —
gfenu 自体によってオンザフライで生成されるため、処理する必要はありません
ロックアウトポリシー —
信頼できるアプリケーションによって強制されることは観察されない
ゲートキーパー署名付き認証トークン —
実装されていない
登録された指のリスト —
ENUMERATE ペイロード +100 でのカウントのみ。コマンドは ID を報告しません
サンプルごとの品質フィードバック —
デコードされていない。エラー語彙はわかっていますが、それを含むフィールドは不明です
どのテンプレートが一致したか —
報告されていない。評決には ID が含まれていないため、ID を構築できません。
DUMP_TEMPLATE は未調査です
タッチでウェイクを表示します - ではありません
デコードされた。 SCREEN_ON / SCREEN_OFF (1017、1018) はクライアントが送信するものであり、送信するものではありません。
センサードライバーの
GW 表示ファミリー —
未テスト。容量性コピーとの主な違いは、
ボードの電源とピン
カーネル側のアップストリーム — どちらのドライバーもビルドして実行しますが、どちらも実行しません
が投稿されました。これらはまずメンテナーと解決する必要があります。
qcom-scm-blocked-listener-warn —
カーネル内の呼び出し元が実行できない qcom_scm_qseecom_call() 内の 2 つの WARN_ON()
到達しますが、通常の呼び出しでも可能であるため、panic_on_warn はそれらを
サービスの拒否。 TEE シリーズが依存するスタンドアロンのパッチ
TEE_IMPL_ID_QSEECOM = 5 は新しい uapi です — クライアントがこのドライバーに指示します
qcomtee から読んで
TEE_IOC_VERSION からの値であるため、番号は ABI であり、TEE のメンテナンスが必要です

インナーの
同じ投稿で確認する
セッションは UUID ではなく名前によって開かれます — QSEE は文字列に一致します
UUID をレンダリングするものがないため、名前は
パラメータと UUID はゼロでなければなりません。 amdtee はその UUID を
ファームウェアのファイル名。 qcomtee も同じ壁にぶつかり、objref パラメータを使用しました。
このドライバーは TEE_GEN_CAP_GP を要求しませんが、TEE サブシステムは
バックエンドをオーバーロードする各バックエンドよりも、独自の汎用的な「名前によるセッション」を優先します。
パラメータ
特権デバイスの open_session はポリモーフィックであることがわかります
「アプリケーションのロード」から「リスナーの登録」はパラメーター0が
値またはメモリ参照。 2 つの関数の値はそれ自体を文書化します
特権デバイス上の CAP_SYS_ADMIN —
他の TEE バックエンドは機能をチェックしません。 OP-TEE は権限に依存します
/dev/teepriv0 の。守るよりもドロップする価値がある
呼び出しごとの TZ メモリ プール - 各コマンドのバウンス バッファ
独自のプールを取得します。 IRQコマンドの場合
これは、割り込みレートでの大きな割り当てであり、キャプチャ中に -ENOMEM を使用します。
断片化の下で
mdt_loader 拡張機能には soc/qcom ack が必要です。これにより 2 つの ack が追加されます。
このプロジェクトが所有していないファイルに対する関数と Kconfig プロンプト。参照してください
カーネルの残りの部分が成長しなければならなかったもの
ウェッジされたサプリカントはすべての QSEECOM ユーザーをブロックします。
カーネル内のものも含まれます - 特権を誰に与えるかを決定します
デバイス
これは概念実証であり、このリストは意図的にその一部として含まれています。どれも
これらは作業ライフサイクルをブロックします。それらはすべてプロダクションにとって重要です
運転手。
これらに共通するのは、ダウンストリーム コードを読み取るか移植すると、
解決しないでください。テストするためのデバイスや、回答するためのベンダーが必要です。
完全にこのプロジェクトの外部からの証拠です。
資格情報が存在するとゲートキーパーのパスが変更されるかどうか。署名されていない
トークンはACです

プロビジョニングされたことがないデバイス ( [私たちのデバイス] ) で受信されました。
プロビジョニング後に何が起こるかはテストされていません。
信頼できるアプリケーションが独自のロックアウトを強制するかどうか。そうではない
観察された (ステータス);欠席を未確認として扱い、実装する
関係なくロックアウト。
他のリスナー サービスが必要になるかどうか。私たちは決して
すべてのフローにわたって、gfenu がそれらのいずれかに対してリクエストを発行していることを観察しました。
実行されました ([私たちのデバイス])。実行されていないパスおよび他のファームウェア バージョンは、
それによってカバーされません。
信頼できるアプリケーション イメージがモデルまたはベンダーにバインドされているかどうか、および
誰がそれに署名するのか。ここで使用されている画像は、モデルの標準 OS から取得したものです。
実行されますが、他の場所では試したことはありません - 同じモデルの別のユニット
これは興味深いケースではありません。別のモデルまたはメーカーが共有している
センサーとSoCです。安全な世界は不正な画像を拒否します。
整合性をチェックしていることは示されていますが、デバイスのバインディングや
サインチェーン。特に、信頼できるアプリケーションが実行できるかどうかはわかりません。
他の Qualcomm ファームウェアと同じように再配布されるか、または他の Qualcomm ファームウェアが再配布されるかどうか
クアルコムまたは OEM のみが保持するキーで署名されています。かどうか誰もチェックしていない
linux-firmware はすでに信頼できるアプリケーションを搭載しており、それに応答します。
例として、最も安価に開始できる場所です。その質問によって、これが可能かどうかが決まります
デバイスごとに抽出されるのではなく、ユーザー向けにパッケージ化されます。
センサー自身の画像データが意味するもの。リファレンスクライアントは生のまま保存可能
フレームですが、ここではフレームを解釈するものはありません。すべての画像処理は内部で行われます
安全な世界。
5 つのピースがあり、それぞれ独立して役立ちます。最初の 4 つは、次のいずれかに存在する必要がありました。
それは機能します。
順番に読んでください。 TA プロトコルはリスナー サービスがなければ意味がありません
実行中のため、リスナー サービスを登録できません

カーネルを注意してください
運転手。
2 つのカーネル ドライバー、それぞれ 1 つのブランチ、両方ともメインライン v7.1 からオフ
ロベルダ/Linux :
これらは独立したシリーズであり、互いに依存しませんが、シーケンスは
以下の場合は両方に加えて、センサー用のデバイス ツリー ノードが必要です。ボード サポートは次のとおりです。
どちらのブランチにもありません。
cc -O2 -o gfharness ハーネス/gfharness.c
# アプリケーションがロードされる前にファイル サービスが実行されている必要があります。
# 初期化中にファイルを要求するため
./gfharness --supp 0 サーブ &
./gfharness --load gfenu # カーネルフェッチ gfenu.mdt + .bNN
./gfharness --capture 300 --enroll # 300 秒;指を繰り返し提示する
./gfharness --capture 120 --auth # 120 秒;それをもう一度提示する
./gfharness --0x1234abcd を削除
--supp 0 は、アプリケーションが 1 つのプロセスから使用するすべてのリスナーを提供します。
必須 — カーネル ドライバーは、1 つのサプリカント キューを 1 つずつ保持します。
デバイス。
ここに記載されているものはすべて、次の場所で検証されました。
デバイス Xiaomi Pad 5 Pro 5G (「enuma」)、ポストマーケットOS
センサー Goodix GF3626 (GF_CHIP_3626ZS1、部品 A005203)、背面
電源ボタン
SoC Qualcomm SM8250、QSEE 信頼できる実行環境
アプリケーション gfenu 、 /vendor/firmware_mnt/image/ からロードされる
センサードライバー
drivers/input/misc/goodix_fp_spi.c (レギュレータを所有する)、リセット行
そして
中断、レポート中断

[切り捨てられた]

## Original Extract

Driving Goodix SPI fingerprint sensors on Linux, including versions running in Qualcomm QSEE enclave. - GitHub - wrobelda/goodix-fp-spi-linux: Driving Goodix SPI fingerprint sensors on Linux, including versions running in Qualcomm QSEE enclave.

I own a Xiaomi Mi Pad 5 Pro 5G Android tablet, which is also pretty-much a perfect Linux tablet: has 5G support, great screen, nice folio keyboard, 8 speakers, a pen, is lightweight and relatively powerful. All of its hardware is supported at this point, although not all has been mainlined yet. It also features a Goodix SPI fingerprint scanner, except it is implemented using an app living in Qualcomm's QSEE secure enclave which handles the scanner, and which Android communicates with using a HAL and a protocol that are not documented. This 2021 article explains the problem in details: https://emainline.gitlab.io/2021/12/12/fingerprint_P1.html It took me about a week of intense work of reverse-engineering it with Opus 5 and Kimi K3, and further 5 days on polishing the documentation, but I can safely say this is now done. I used Opus 5, since Fable naturally dropped out as this is security subject, and Kimi K3 for reviewing the code. I resorted to both disassembling and tracing live Android sessions using Frida to figure out the issues along the way. Happy to answer any questions. Also, I expect the same approach can be used to reverse-engineer the secure-enclaved fingerprint scanners on Macbooks and since I run Asahi on my M1, I will look into it when I next take a break from regular work.

GitHub - wrobelda/goodix-fp-spi-linux: Driving Goodix SPI fingerprint sensors on Linux, including versions running in Qualcomm QSEE enclave. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
wrobelda
/
goodix-fp-spi-linux
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows docs docs harness harness tracing tracing .gitignore .gitignore LICENSE LICENSE README.md README.md _typos.toml _typos.toml View all files Repository files navigation
Goodix SPI fingerprint sensors
A Goodix SPI fingerprint sensor — not the unrelated USB kind common in laptops —
can be driven in two ways.
In the "traditional" one, the Rich Execution Environment (REE), the operating system's
kernel driver handles all the hardware over SPI bus protocol, allowing the user space to
talk to the sensor directly to do the imaging, to perform the
imaging, the templating and the matching logic.
However, this is rather rare nowadays as it involves substantial security risk,
since the images and the templates are handled by the operating system, where
anything that compromises the kernel or the userspace stack can read or
substitute them.
The alternative is running the fingerprint logic in a Trusted Execution
Environment (TEE), such that the operating system never touches the imaging or the
matching at all — which instead happen inside a trusted, proprietary, signed and
closed-source application running in the SoC's secure world. The kernel driver
is left with handling only the bare minimum of the hardware, without access to
the SPI bus the scanner is connected to.
When it comes to Goodix scanners, two known implementations of such TEE
applications exist, one for MediaTek's MicroTrust TEE, the other for Qualcomm's
QSEE, and they ship in the operating system the device comes with.
Supporting the Goodix in TEE mode on Linux
TEE mode is much safer of the two for the user, but it also pretty much
guarantees that the OEM of the — typically an Android — device keeps the Hardware
Abstraction Layer (HAL) library closed, and it is that library that implements the
entirety of the communication between the Operating System's permissions subsystem
and the trusted application. The only open piece is the kernel sensor
driver, which does the bare minimum of bringing the scanner up.
This unfortunately means that the only way to add support for these scanners in Linux
is by reverse-engineering the Android implementation. A 2021 article,
Fingerprint sensors on tuxified Android phones: Impossible?
( archived ),
describes well what a port would have to do and why nobody had done it.
The project was developed using Claude Opus 5 and Kimi K3 over a span of 10 days.
Claude Fable 5 would likely speed things up a bit, but anything security-adjacent
gets flagged by Anthropic and the model was getting routinely dropped to Opus 5,
hence the Kimi K3 usage here.
This is not unverified AI slop — most of the architecture decisions and code layout
were made by myself, as well as the disassembly/reverse-engineering approach and driving it.
The work here is tested on my hardware, with the hope of further refinement, testing on more
devices and readying for mainlining into the Linux kernel.
The TEE support for these scanners on Qualcomm SoCs is now done.
This project serves as a proof of concept rather than a complete userspace implementation: what this project
provides is the protocol, the two kernel drivers needed to reach it on the hardware side and on the TEE side, and a
reference client that demonstrates the sequence, as the basis for a real
fprintd driver.
REE mode — the operating system does the imaging itself over SPI. It
can be added to the sensor driver by porting
the SPI transfer and image scanning code from the MediaTek-lineage gf_spi_tee.c , where it sits behind
SUPPORT_REE_SPI
TEE mode on Qualcomm — reaching the trusted application over
QSEECOM; everything documented here was exercised in this mode:
load and talk to the trusted application
persistence across reboot —
QSEE-sealed objects, stored through the listener services
TEE mode on MediaTek — the protocol and the sensor driver carry
over; the missing half is a MicroTrust backend in drivers/tee/ , where
MediaTek platforms that reach a TEE upstream do so through OP-TEE. Requires
reverse-engineering
The trusted application's protocol —
Goodix-generic rather than Qualcomm's, so these should carry over to a
MediaTek port:
enrolment, first finger
authentication (match and reject)
calibration —
generated on-the-fly by gfenu itself, nothing to handle
lockout policy —
not observed being enforced by the trusted application
gatekeeper-signed auth token —
not implemented
listing enrolled fingers —
only a count, at ENUMERATE payload +100; no command reports ids
per-sample quality feedback —
not decoded; the error vocabulary is known, the field carrying it is not
which template matched —
not reported; the verdict carries no id, so identify cannot be built, and
DUMP_TEMPLATE is unexplored
display wake on touch — not
decoded; SCREEN_ON / SCREEN_OFF (1017, 1018) are a client's to send, not
the sensor driver's
GW in-display family —
untested; differs from the capacitive copies mostly in
board power and pins
Upstreaming the kernel side — both drivers build and run, but neither
has been posted. These have to be settled with the maintainers first:
qcom-scm-blocked-listener-warn —
two WARN_ON() s in qcom_scm_qseecom_call() that no in-kernel caller can
reach but an ordinary invoke can, so panic_on_warn turns them into a
denial of service. A standalone patch the TEE series depends on
TEE_IMPL_ID_QSEECOM = 5 is new uapi — a client tells this driver
from qcomtee by reading it
from TEE_IOC_VERSION , so the number is ABI and needs the TEE maintainer's
ack in the same posting
Sessions are opened by name, not by UUID — QSEE matches on a string
and there is nothing to render a UUID into, so the name arrives in a
parameter and the UUID must be zero. amdtee renders its UUID into a
firmware filename; qcomtee hit the same wall and used objref parameters.
This driver does not claim TEE_GEN_CAP_GP , but the TEE subsystem may
prefer a generic "session by name" of its own to each backend overloading a
parameter
The privileged device's open_session is polymorphic — it tells
"register a listener" from "load an application" by whether parameter 0 is
a value or a memref. Two func values would document themselves
CAP_SYS_ADMIN on the privileged device —
no other TEE backend checks a capability; OP-TEE relies on the permissions
of /dev/teepriv0 . Worth offering to drop rather than defending
A TZ memory pool per invoke — the bounce buffer for each command
gets its own pool. For the IRQ command
that is a large allocation at interrupt rates, and -ENOMEM mid-capture
under fragmentation
The mdt_loader extension needs a soc/qcom ack — it adds two
functions and a Kconfig prompt to a file this project does not own; see
what the rest of the kernel had to grow
A wedged supplicant blocks every QSEECOM user ,
in-kernel ones included — which decides who may be granted the privileged
device
This is a proof of concept, and this list is deliberately part of it. None of
these block the working lifecycle; all of them would matter to a production
driver.
What these have in common is that reading or porting the downstream code does
not settle them. They need a device to test against, a vendor to answer, or
evidence from outside this project entirely.
Whether the gatekeeper path changes once a credential exists. An unsigned
token is accepted on a device that never had one provisioned ( [our device] ).
What happens after provisioning is untested.
Whether the trusted application enforces any lockout of its own. Not
observed ( status ); treat the absence as unverified and implement
lockout regardless.
Whether the other listener services are ever needed. We never
observed gfenu raising a request on any of them, across every flow we
exercised ( [our device] ). Unexercised paths, and other firmware versions, are
not covered by that.
Whether a trusted application image is bound to a model or a vendor, and
who signs it. The image used here came from the stock OS of the model it
runs on, and we never tried it anywhere else — another unit of the same model
is not the interesting case; a different model or manufacturer sharing the
sensor and SoC is. The secure world rejects a malformed image, which
shows it checks integrity, but says nothing about device binding or about the
signing chain. In particular we do not know whether trusted applications can
be redistributed the way other Qualcomm firmware is, or whether they are
signed with keys held only by Qualcomm or the OEM. Nobody has checked whether
linux-firmware already carries a trusted application, which would answer it
by example and is the cheapest place to start. That question decides whether any of this can be
packaged for users rather than extracted per device.
What the sensor's own image data means. The reference client can save raw
frames, but nothing here interprets them; all image processing happens inside
the secure world.
Five pieces, each independently useful; the first four had to exist for any of
it to work.
Read them in order. The TA protocol is meaningless without a listener service
running, and the listener service cannot be registered without the kernel
driver.
Two kernel drivers, one branch each, both off mainline v7.1 at
wrobelda/linux :
They are independent series and do not depend on each other, but the sequence
below needs both, plus a device tree node for the sensor — the board support is
not on either branch.
cc -O2 -o gfharness harness/gfharness.c
# the file service must be running before the application is loaded,
# because it asks for its files during initialisation
./gfharness --supp 0 serve &
./gfharness --load gfenu # kernel fetches gfenu.mdt + .bNN
./gfharness --capture 300 --enroll # 300 seconds; present a finger repeatedly
./gfharness --capture 120 --auth # 120 seconds; present it again
./gfharness --remove 0x1234abcd
--supp 0 serves every listener the application uses from one process, which is
required — the kernel driver keeps one supplicant queue per
device .
Everything documented here was verified on:
Device Xiaomi Pad 5 Pro 5G ("enuma"), postmarketOS
Sensor Goodix GF3626 ( GF_CHIP_3626ZS1 , part A005203 ), behind the
power button
SoC Qualcomm SM8250, QSEE trusted execution environment
Application gfenu , loaded from /vendor/firmware_mnt/image/
Sensor driver
drivers/input/misc/goodix_fp_spi.c , which owns the regulator, reset line
and
interrupt, and reports interrup

[truncated]
