---
source: "https://github.com/danielealbano/android-remote-control-mcp/"
hn_url: "https://news.ycombinator.com/item?id=49362047"
title: "Show HN: MCP app for Android, drive apps via AI (no root, PII redacted locally)"
article_title: "GitHub - danielealbano/android-remote-control-mcp: An MCP Server for Android running on the phone, optmized for token usage, supports also files downloads and cloudflare and ngrok automated tunnelling. · GitHub"
image: "https://opengraph.githubassets.com/2cfdac1520898191285cd57aebf42639cc1fa5af16f65d48837e342351e6c8c5/danielealbano/android-remote-control-mcp"
author: "daniele_dll"
captured_at: "2026-08-19T15:23:04Z"
capture_tool: "hn-digest"
hn_id: 49362047
score: 4
comments: 0
posted_at: "2026-08-19T14:23:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: MCP app for Android, drive apps via AI (no root, PII redacted locally)

- HN: [49362047](https://news.ycombinator.com/item?id=49362047)
- Source: [github.com](https://github.com/danielealbano/android-remote-control-mcp/)
- Score: 4
- Comments: 0
- Posted: 2026-08-19T14:23:13Z

## Translation

タイトル: HN を表示: Android 用 MCP アプリ、AI 経由でアプリを駆動 (ルートなし、PII はローカルで編集)
記事のタイトル: GitHub - danielealbano/android-remote-control-mcp: 携帯電話上で動作する Android 用 MCP サーバーは、トークンの使用に最適化されており、ファイルのダウンロード、cloudflare および ngrok 自動トンネリングもサポートしています。 · GitHub
説明: 電話機上で実行される Android 用 MCP サーバーは、トークンの使用に最適化されており、ファイルのダウンロード、cloudflare および ngrok の自動トンネリングもサポートしています。 - ダニエルバノ/android-remote-control-mcp
HN テキスト: こんにちは、HN、著者です。 Android Remote Control MCP は、Android スマートフォン上で直接実行される MCP サーバー (ルート、ADB、中間のコンピューターは不要) で、AI エージェントが人間と同じように実際のアプリを操作できるようにします。つまり、アクセシビリティ ツリーを通じて画面を読み取り、タップ、入力、スクロールし、オプションでスクリーンショットも取得します。ツールの使用とトークン消費の最適化に多くの時間を費やし、ローカル ハーネス経由だけでなく、Claude.ai / Claude Desktop および chatgpt.com 経由でも機能するようにしました (適切なアカウントを持っている場合、アプリは独自の OAuth サーバーとして機能し、電話上のコードで接続を承認します)。最新の部分はプライバシー モードです。これが存在するのは、以前のリリースの後、誰かが私に、当然のことながら、LLM プロバイダーに画面上のすべてを見せたくないため、決して使用しないと明言したからです。
そこで現在は、小規模なローカル モデルと決定論的検出器の組み合わせにより、個人情報 (電子メール、電話番号、クレジット カード、IBAN、国民 ID、英語名など) を特定し、何かが携帯電話から送信される前にデバイス上で秘匿化します。ベンチマーク検出率は約 87% です (ベンチマークはリポジトリにあります。英語以外の名前は現在の弱点ですが、現在取り組んでいます)。なぜ私がそれを建てたのか?エージェントが使用できるようにしたい

携帯電話を使って、検索したり、予約したり、時には退屈な作業をしたりすることができます。すべてのデータをサービス プロバイダーと共有したり、ローカル LLM を実行したりする必要はありません。トレードオフ: このサービスは自身をアクセシビリティ アプリとして宣言しているため、アプリを読み取ることができますが、このアプリを Google Play ストアで配布することもできなくなります。そのため、現時点では GitHub 経由で配布されています (標準ビルドと、Google Play Services を使用しない FOSS ビルドがあり、間もなく F-Droid で公開されます)。また、私はプライバシーと同じくらいプロンプト インジェクションを重視しているため、エージェントに返されるデータには、コンテンツがサード パーティのソースから提供され、信頼できない入力であることを明確にするために、かなり強い表現を使用したメッセージがプレフィックスとして付けられます。潜在的な攻撃を大幅に軽減します。完璧ですか？いいえ、特にアプリのカバー範囲に関しては改善すべき点がたくさんありますが、十分な詳細を提供すれば、小規模なモデル (Haiku など) でも電話を操作し、目に見えない状況に対処できます。次は何に取り組んでいますか?大きく分けて 3 つのこと:
- 無料のリバース トンネル、エンドツーエンドで暗号化され、Let's Encrypt 証明書を使用すると、Claude.ai / Claude Desktop / Chatgpt.com などでアプリを無料で使用できるようになり、ngrok や Cloudflare を使用するよりもはるかにプライバシーが高くなります (e2e 暗号化は提供されず、エッジへの暗号化のみが提供されるため、すべてが表示されます)
- アプリのスキル データベースにより、エージェントは何をすべきかをより迅速に判断できます
- 特に英語以外の名前に対するプライバシーの立場を改善するためのカスタム モデル さらに、アプリの使用を簡単に開始できるように、新しいガイド付きセットアップと改訂された UI が追加されました。

記事本文:
GitHub - danielealbano/android-remote-control-mcp: 携帯電話上で実行される Android 用 MCP サーバーは、トークンの使用に最適化されており、ファイルのダウンロード、cloudflare および ngrok 自動トンネリングもサポートしています。 · GitHub
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
ダニエルバノ
/
アンドロイド-リモートコントロール-mcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
933 コミット 933 コミット フォルダーとファイル
.claude-プラグイン .claude-plugin .claude/ ag

ents .claude/ エージェント .github/ ワークフロー .github/ ワークフロー アプリ アプリ チャネル プラグイン チャネル プラグイン compose-test-app compose-test-app docs docs e2e-tests e2e-tests gradle gradle プライバシー-ベンチマーク プライバシー-ベンチマーク プライバシー プライバシー スクリプト スクリプト ベンダー ベンダー .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .gitmodules .gitmodules CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE.md LICENSE.md Makefile Makefile README.md README.md build.gradle.kts build.gradle.kts gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat keystore.properties.example keystore.properties.example settings.gradle.kts settings.gradle.kts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
MCP (Model Context Protocol) サーバーとして実行される Android アプリケーション。これにより、AI モデルがアクセシビリティ サービスとスクリーンショット キャプチャを使用して Android デバイスをリモートで完全に制御できるようになります。
アプリは Android デバイス (またはエミュレーター) 上で直接実行され、MCP プロトコルを実装する HTTP サーバー (オプションの HTTPS あり) を公開します。クロードのような AI モデルはそれに接続し、UI 要素の読み取り、ボタンのタップ、テキストの入力、スワイプ、スクリーンショットのキャプチャ、ファイルの管理、アプリの起動など、デバイス上のあらゆるアプリと対話できます。
警告: このソフトウェアは、研究および教育目的のみに、いかなる種類の保証もなく「現状のまま」提供されます。作者は、違法、無許可、または非倫理的な活動にこのツールを使用することを容認しません。ユーザーは、その使用が適用されるすべての法律および規制に準拠していることを確認する単独の責任を負います。このソフトウェアを使用することにより、お客様は責任を持って、ご自身の責任で使用することに同意したものとみなされます。
Reddit への投稿 (高速化、オリジナルは約 6 分)
Booking.comアプリのインストール
スカイスキャナーでのフライト検索、チューリッヒ → イスタンブール (高速化、~4

最小オリジナル)
MCP サーバー経由で Android デバイスを制御する AI モデル — 任意のクリップをクリックするとフル解像度が表示されます。
Android (Ktor + Netty) 上で直接実行される HTTP サーバー (オプションの HTTPS あり)
/mcp でのストリーミング可能な HTTP トランスポート (MCP 仕様準拠、JSON のみ、SSE なし)
複合認証: 静的ベアラー トークンおよび/または自己完結型 OAuth 2.1 サーバー (Claude.ai / Claude Desktop カスタム コネクタ、デバイス上の承認付き)
自動生成された自己署名 TLS 証明書 (またはカスタム証明書のアップロード)
構成可能なバインディング: localhost (127.0.0.1) または network (0.0.0.0)
Cloudflare Quick Tunnels または ngrok (パブリック HTTPS URL) を介したリモート アクセス トンネル
14 カテゴリにわたる 57 の MCP ツール
画面のイントロスペクション、システム アクション、タッチ アクション、ジェスチャ、ノード アクション、テキスト入力、ユーティリティ、ファイル操作、アプリ管理、カメラ、インテント、通知、位置情報、共有。
すべてのツール名はデフォルトで android_ 接頭辞を使用します (例: android_tap )。デバイス スラグが設定されている場合 (例:Pixel7 )、プレフィックスは android_pixel7_ (例: android_pixel7_tap ) になります。
入出力スキーマと例を含む完全なツール リファレンスについては、docs/MCP_TOOLS.md を参照してください。
タブ付きレイアウト (サーバー / 設定 / バージョン情報) とダーク モードを備えたマテリアル デザイン 3 UI
サーバーのステータス監視 (実行中/停止中) と権限警告バナー
接続情報表示（IP、ポート、トークン、トンネルURL）
ツールごとおよびパラメータごとの権限 (個々の MCP ツールの有効化/無効化)
権限管理 (アクセシビリティ、通知、カメラ、マイク)
リモートアクセストンネル構成 (Cloudflare / ngrok)
保管場所の管理 (自動保管場所 + ファイルツール用の SAF 認証)
サーバー ログ ビューア (MCP ツール呼び出し、トンネル イベント)
ADB 経由のヘッドレス セットアップ (UI を使用しない構成、権限の付与、サーバーの起動/停止)
特徴
これ

プロジェクト
モバイルMCP
Android-MCP
アンドロイド-mcp-サーバー
adb-mcp
ドロイドラン-mcp
MCPツール
57
21
11
5
10
11
電話機上で実行 (ADB なし)
✅
❌
❌
❌
❌
❌
アクションのレイテンシ
10～100ミリ秒
1～4秒
1～4秒
1～4秒
1～4秒
1～4秒
インターネット経由で動作します
✅
❌
❌
❌
❌
❌
トークン効率の高い画面状態
✅
❌
✅
❌
❌
✅
注釈付きのスクリーンショット
✅
❌
✅
❌
❌
❌
設定可能なスクリーンショットの解像度/品質
✅
❌
❌
❌
❌
❌
ツールごとの有効化/無効化
✅
❌
❌
❌
❌
❌
マルチデバイスのサポート
✅
✅
❌
❌
❌
❌
カメラ、クリップボード、ファイル、ダウンロード
✅
❌
❌
❌
❌
❌
iOSのサポート
❌
✅
❌
❌
❌
❌
ほとんどの代替手段は、ホスト マシン上で実行される ADB に依存します。これは、USB ケーブルまたはローカル ネットワーク接続と電話機の隣にあるコンピュータを意味します。このプロジェクトは完全にデバイス自体で実行されるため、トンネルを通じて MCP エンドポイントを公開し、どこからでも電話を制御できます。
トークン効率の面では、ADB ベースのツールは通常、生の uiauTomator XML ダンプを返しますが、これはここで使用されているコンパクトな表現よりも簡単に 10 ～ 50 倍冗長になる可能性があります。番号付きのスクリーンショットの注釈、設定可能な画質、不要なツールを無効にする機能 (すべてのツール定義は毎ターントークンを消費します) と組み合わせることで、エージェント ループのインタラクションごとのコストが大幅に削減されます。
どのビルドをダウンロードするか: GMS または FOSS
各リリースには 2 つの種類の APK が提供されています。お使いのデバイスに一致するものを選択してください。
GMS ( …-gms-release.apk ) — 完全なビルド。電話機に Google モバイル サービス (Google Play サービス) がインストールされている必要があります。 Google サービスに同梱されている標準デバイス/エミュレータではこれを選択してください。
FOSS ( …-foss-release.apk ) — Google モバイル サービスに依存しない Google フリーのビルド。 Google から除外された ROM を含む、あらゆる Android スマートフォンで動作します。

Google サービスが利用できないデバイス。
よくわからない場合は、Play ストアを備えた一般的な携帯電話には GMS ビルドが適切な選択です。
オプション A: 携帯電話にダウンロード (最も簡単)
携帯電話のブラウザでリリース ページを開きます
最新リリースから APK をダウンロードする
ダウンロードした APK を開き、プロンプトに従ってインストールします (不明なソースからのインストールを許可する必要がある場合があります)
オプション B: PC にダウンロードし、ADB 経由でインストールする
リリースページからAPKをダウンロードします。
USB 経由で携帯電話を接続します (USB デバッグが有効な場合)
APK をインストールします (ダウンロードした GMS または FOSS ファイルを使用します)。
adb install android-remote-control-mcp-<バージョン> -gms-release.apk
オプション C: ソースからビルドする
git clone https://github.com/danielealbano/android-remote-control-mcp.git
cd アンドロイド-リモート コントロール-mcp
ビルドする
make install # 接続されたデバイス/エミュレータにインストールします
完全なビルド要件と手順については、CONTRIBUTING.md を参照してください。
アプリを開きます。 「サーバー」タブには、許可に関する警告バナーが表示されます。そこから、または [設定] > [権限] を介して権限を付与します。
アクセシビリティ サービス — 必須。 UI のイントロスペクション、アクションの実行、スクリーンショットを強化します。それなしでは何も機能しません。
トグルがグレー表示になっている場合 (Android 13 以降、Samsung/One UI で APK からインストールされたアプリに共通): これは Android の「制限付き設定」保護です。ダイアログに「制限された設定 — セキ​​ュリティのため、この設定は現在利用できません。」と表示されます。ブロックを解除するには、 [設定] → [アプリ] → [Android Remote Control MCP] → [アプリ情報] に移動し、 [⋮] メニュー (右上) をタップして [制限付き設定を許可] を選択し、PIN/パターン/生体認証で確認してから、ユーザー補助サービスを有効にします。 (Play ストア / F-Droid のインストールには影響しません。)
通知 — オプション (OAuth の場合に推奨)。 OAuth サインインの承認 (Claude.ai / Claude Desktop、

ChatGPT またはサードパーティの OAuth クライアント) は、2 桁のコードを確認して承認するためにタップすると、ヘッドアップ通知として到着します。ただし、OAuth ログインには必要なくなりました。この権限を付与しない場合でも、保留中のリクエストはアプリのサーバー画面にカードとして表示され、タップして確認して承認/拒否します。イベント チャネル プッシュ機能 — 例:デバイス イベントを Claude Code にプッシュする — Android 13 以降では、この権限を通じてステータス通知を投稿するフォアグラウンド サービスとして実行されます。
その他の権限 (オプション) - カメラ、マイク、位置情報は、対応する MCP ツールを有効にします。アプリはそれらがなくても動作します。依存する機能のみが許可されるまで無効になります。完全なリストについては、以下の権限リファレンスを参照してください。
ストレージの場所 — ファイル ツールを使用する場合は、[設定] > [ストレージ] で構成します (ダウンロードなどの自動の場所、および SAF によるカスタムの場所)。
サーバーを起動します。 「サーバー」タブに戻り、「開始」をタップします。
サーバーはデフォルトで http://127.0.0.1:8080 で起動します。接続情報 (IP、ポート、トークン、URL) が [サーバー] タブに表示されます。
注: 127.0.0.1 は、コンピュータではなく電話のローカルホストを指します。コンピュータから接続するには、 adb ポート転送を使用するか、0.0.0.0 (ネットワーク モード) にバインドするか、リモート アクセス トンネルを有効にします。
プライバシー モードは、電子メール、クレジット カードと IBAN、資格情報、電話などの個人データを検出して隠します。
番号、国民 ID、住所、名前など、MCP ツールが返すすべてのもの (画面のコンテンツ、
通知、クリップボード、ファイルなど) がデバイスから AI プロバイダーに送られる前に。検出
完全にデバイス上で実行: 高速なルールベースのチェック (チェックサム、フォーマット、フィールド コンテキスト) と、
ローカル NER モデル - クラウド呼び出しはありません。検出された値は、一貫性のある仮名化のいずれかです。
プレースホルダー ( EMAIL#a

1b2c 、つまり AI ツールは画面間で動作し続ける)、または完全に編集されます。検出
これはベスト エフォート型の緩和策であり、保証するものではありません。
モデルは Ai4Privacy 多言語匿名化ツールです
(MIT ライセンス、Llama で構築)、最初の有効化時に 1 回 (~150 MB) をピン留めされたファイルから直接ダウンロード
デバイス上でハグフェイスのリビジョンとチェックサムが検証されました。
インリポジトリで測定された、パイプライン全体によってカテゴリごとに検出された個人データ項目のシェア
8 つの言語にわたるリアルな UI スタイルのコンテンツに対する有効性ベンチマーク ( make Privacy-benchmark ):
名前は、オンデバイス モデルの既知の弱点です (西洋の一般的な名前は、
ベンチマーク内の世界的に多様な名前);微調整されたモデルによる名前検出の改善がオンです
ロードマップ。実際の結果はコンテンツや言語によって異なります。
メインのサーバー画面から: [プライバシー モード] カードの [プライバシー モードのセットアップ] をタップします (表示中に表示されます)。
モードはオフです）。
または、[設定] → [プライバシー] で、[プライバシー モードを有効にする] をオンにし、ワンタイム モデルを確認します。
ダウンロード。ダウンロードとウォームアップの後、デバイス上の短いベンチマークにより、予想されるパフォーマンスが測定されます。
画面ごとのオーバーヘッドを軽減し、同じ画面上に表示します。
保護されたカテゴリ、仮名化対墨消し、およびプレースホルダー形式は、次の場所で構成できます。
設定 → プライバシー 。
サーバーを .mcp.json 構成ファイルに追加します。
{
"mcpサーバー": {
"アンドロイドフォン" : {
"type" : " http " ,
"url" : "ht

[切り捨てられた]

## Original Extract

An MCP Server for Android running on the phone, optmized for token usage, supports also files downloads and cloudflare and ngrok automated tunnelling. - danielealbano/android-remote-control-mcp

Hi HN, author here. Android Remote Control MCP is an MCP server that runs directly on your Android phone (no root, no ADB, no computer in the middle) and lets an AI agent drive real apps the way a human would: it reads the screen through the accessibility tree, taps, types, scrolls, optionally also screenshots. I spent a lot of time optimizing tool usage and token consumption, and making it work not only via local harnesses but also via Claude.ai / Claude Desktop and chatgpt.com (if you have the proper account, the app acts as its own OAuth server, you approve connections with a code on the phone). The newest part is Privacy Mode, and it exists because after an earlier release someone told me in plain terms they'd never use it because, rightfully, they didn't want the LLM provider to see everything on their screen!
So now a combination of a small local model plus deterministic detectors identify personal information (emails, phone numbers, credit cards, IBANs, national IDs, English names, etc.) and redact it on-device before anything leaves the phone, with a benchmarked detection rate of about 87% (the benchmark is in the repo; non-English names, are the current weak spot but I am working on it). Why did I build it? I want my agents to be able to use my phone to do searches, book things for me and do sometimes boring stuff ... without sharing all my data with service providers and without having to run a local LLM! Tradeoffs: the service declares itself as an accessibility app so it can read apps which also makes it impossible to distribute this app on the Google Play store, hence it's distributed via GitHub at the moment (there's a standard build and a FOSS build without Google Play Services which soon will be published on F-Droid). Also, because I care about prompt injection as much as privacy, the data returned to the agent is prefixed with a message, which uses quite some strong wording, to make it clear that the content is supplied by third party sources and is an untrusted input! Mitigates the potential attacks a lot. Is it perfect? Nope, there's plenty to improve, especially around apps coverage but even small models (like Haiku) can drive the phone and handle unseen situations if you give them enough detail! What I am working on next? Three major things:
- a free reverse tunnel, encrypted end to end, with Let's Encrypt certificates to be able to use the app with Claude.ai / Claude Desktop / Chatgpt.com and such for free and with much more privacy than using ngrok or cloudflare (which don't provide e2e encryption but only encryption to the edge, so they see everything)
- a skills database for the apps, so your agent can work out what to do much faster
- a custom model to improve the privacy standing, especially for non-English names And, additionally, a new guided setup and a revised UI to make it easier to start to use the app!

GitHub - danielealbano/android-remote-control-mcp: An MCP Server for Android running on the phone, optmized for token usage, supports also files downloads and cloudflare and ngrok automated tunnelling. · GitHub
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
danielealbano
/
android-remote-control-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
933 Commits 933 Commits Folders and files
.claude-plugin .claude-plugin .claude/ agents .claude/ agents .github/ workflows .github/ workflows app app channel-plugin channel-plugin compose-test-app compose-test-app docs docs e2e-tests e2e-tests gradle gradle privacy-benchmark privacy-benchmark privacy privacy scripts scripts vendor vendor .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .gitmodules .gitmodules CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE.md LICENSE.md Makefile Makefile README.md README.md build.gradle.kts build.gradle.kts gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat keystore.properties.example keystore.properties.example settings.gradle.kts settings.gradle.kts View all files Repository files navigation
An Android application that runs as an MCP (Model Context Protocol) server , enabling AI models to fully control an Android device remotely using accessibility services and screenshot capture.
The app runs directly on your Android device (or emulator) and exposes an HTTP server (with optional HTTPS) implementing the MCP protocol. AI models like Claude can connect to it and interact with any app on the device — reading UI elements, tapping buttons, typing text, swiping, capturing screenshots, managing files, launching apps, and more.
Warning: This software is provided "as-is" without warranty of any kind, for research and educational purposes only . The authors do not condone the use of this tool for any illegal, unauthorized, or unethical activities. Users are solely responsible for ensuring their use complies with all applicable laws and regulations. By using this software, you agree to use it responsibly and at your own risk.
Posting on Reddit (sped up, ~6 min original)
Installing the booking.com app
Flight search on Skyscanner, Zurich → Istanbul (sped up, ~4 min original)
An AI model controlling an Android device through the MCP server — click any clip for full resolution.
HTTP server running directly on Android (Ktor + Netty), with optional HTTPS
Streamable HTTP transport at /mcp (MCP specification compliant, JSON-only, no SSE)
Combined authentication: static bearer token and/or a self-contained OAuth 2.1 server (Claude.ai / Claude Desktop custom connectors, with on-device approval)
Auto-generated self-signed TLS certificates (or custom certificate upload)
Configurable binding: localhost (127.0.0.1) or network (0.0.0.0)
Remote access tunnels via Cloudflare Quick Tunnels or ngrok (public HTTPS URL)
57 MCP Tools across 14 Categories
Screen introspection, system actions, touch actions, gestures, node actions, text input, utilities, file operations, app management, camera, intents, notifications, location, and sharing.
All tool names use the android_ prefix by default (e.g., android_tap ). When a device slug is configured (e.g., pixel7 ), the prefix becomes android_pixel7_ (e.g., android_pixel7_tap ).
See docs/MCP_TOOLS.md for the full tool reference with input/output schemas and examples.
Material Design 3 UI with tabbed layout (Server / Settings / About) and dark mode
Server status monitoring (running/stopped) with permission warning banner
Connection info display (IP, port, token, tunnel URL)
Per-tool and per-parameter permissions (enable/disable individual MCP tools)
Permission management (Accessibility, Notifications, Camera, Microphone)
Remote access tunnel configuration (Cloudflare / ngrok)
Storage location management (automatic locations + SAF authorization for file tools)
Server log viewer (MCP tool calls, tunnel events)
Headless setup via ADB (configure, grant permissions, start/stop server without UI)
Feature
This project
mobile-mcp
Android-MCP
android-mcp-server
adb-mcp
droidrun-mcp
MCP tools
57
21
11
5
10
11
Runs on the phone (no ADB)
✅
❌
❌
❌
❌
❌
Action latency
10-100 ms
1-4 s
1-4 s
1-4 s
1-4 s
1-4 s
Works over the internet
✅
❌
❌
❌
❌
❌
Token-efficient screen state
✅
❌
✅
❌
❌
✅
Annotated screenshots
✅
❌
✅
❌
❌
❌
Configurable screenshot resolution/quality
✅
❌
❌
❌
❌
❌
Per-tool enable/disable
✅
❌
❌
❌
❌
❌
Multi-device support
✅
✅
❌
❌
❌
❌
Camera, clipboard, files, downloads
✅
❌
❌
❌
❌
❌
iOS support
❌
✅
❌
❌
❌
❌
Most alternatives rely on ADB running on a host machine, which means a USB cable or local network connection and a computer sitting next to the phone. This project runs entirely on the device itself, so you can expose the MCP endpoint through a tunnel and control your phone from anywhere.
On the token efficiency side, ADB-based tools typically return raw uiautomator XML dumps which can easily be 10-50x more verbose than the compact representation used here. Combined with numbered screenshot annotations, configurable image quality, and the ability to disable tools you don't need (every tool definition costs tokens on every turn), this significantly reduces the per-interaction cost in agentic loops.
Which build to download: GMS or FOSS
Each release provides two flavors of the APK — pick the one that matches your device:
GMS ( …-gms-release.apk ) — the full build. Requires Google Mobile Services (Google Play Services) installed on the phone. Choose this on standard devices/emulators that ship with Google services.
FOSS ( …-foss-release.apk ) — a Google-free build with no dependency on Google Mobile Services . Works on any Android phone, including de-Googled ROMs and devices where Google services are unavailable.
If unsure, the GMS build is the right choice for a typical phone with the Play Store.
Option A: Download on your phone (easiest)
Open the Releases page on your phone's browser
Download the APK from the latest release
Open the downloaded APK and follow the prompts to install it (you may need to allow installation from unknown sources)
Option B: Download on your PC and install via ADB
Download the APK from the Releases page
Connect your phone via USB (with USB Debugging enabled)
Install the APK (use the GMS or FOSS file you downloaded):
adb install android-remote-control-mcp- < version > -gms-release.apk
Option C: Build from sources
git clone https://github.com/danielealbano/android-remote-control-mcp.git
cd android-remote-control-mcp
make build
make install # installs on connected device/emulator
See CONTRIBUTING.md for full build requirements and instructions.
Open the app. The Server tab shows a permission warning banner; grant permissions from there or via Settings > Permissions.
Accessibility Service — required. Powers UI introspection, action execution, and screenshots; nothing works without it.
If the toggle is greyed out (Android 13+, common on Samsung/One UI for apps installed from an APK): this is Android's "Restricted settings" protection. A dialog reads "Restricted setting — For your security, this setting is currently unavailable." To unblock it, go to Settings → Apps → Android Remote Control MCP → App info , tap the ⋮ menu (top-right), choose Allow restricted settings , confirm with your PIN/pattern/biometric, then enable the Accessibility Service. (Play Store / F-Droid installs are not affected.)
Notifications — optional (recommended for OAuth). OAuth sign-in approvals (Claude.ai / Claude Desktop, ChatGPT, or any third-party OAuth client) arrive as heads-up notifications you tap to confirm the 2-digit code and approve. It is no longer required for OAuth login, though: if you don't grant this permission, pending requests still appear as a card on the app's Server screen that you tap to review and approve/deny. The Event Channel push feature — e.g. pushing device events to Claude Code — runs as a foreground service that posts its status notification through this permission on Android 13+.
Other permissions (optional) — Camera, Microphone, and Location enable their corresponding MCP tools. The app works without them; only the dependent features are disabled until granted. See the Permissions Reference below for the complete list.
Storage locations — configure in Settings > Storage if you plan to use the file tools (automatic locations like Downloads, plus custom locations via SAF).
Start the server. Go back to the Server tab and tap Start .
The server starts on http://127.0.0.1:8080 by default. The connection info (IP, port, token, URL) is displayed on the Server tab.
Note : 127.0.0.1 refers to the phone's localhost, not your computer. To connect from your computer, use adb port forwarding , bind to 0.0.0.0 (network mode), or enable a remote access tunnel .
Privacy Mode detects and hides personal data — emails, credit cards and IBANs, credentials, phone
numbers, national IDs, addresses, and names — in everything the MCP tools return (screen content,
notifications, clipboard, files, …) before it leaves the device for the AI provider. Detection
runs fully on-device: fast rule-based checks (checksums, formats, field context) combined with a
local NER model — no cloud calls. Detected values are either pseudonymized with consistent
placeholders ( EMAIL#a1b2c , so AI tools keep working across screens) or fully redacted. Detection
is best-effort mitigation, not a guarantee.
The model is the Ai4Privacy multilingual anonymiser
(MIT license, Built with Llama), downloaded once on first enable (~150 MB) directly from the pinned
Hugging Face revision and checksum-verified on device.
Share of personal data items detected per category by the full pipeline, measured with the in-repo
effectiveness benchmark ( make privacy-benchmark ) on realistic UI-style content across 8 languages:
Names are the on-device model's known weak spot (it recognizes common Western names far better than
the globally diverse names in the benchmark); improving name detection with a fine-tuned model is on
the roadmap. Actual results vary with content and language.
From the main Server screen : tap Set up Privacy Mode on the Privacy Mode card (shown while
the mode is off).
Or via Settings → Privacy : turn on Enable Privacy Mode and confirm the one-time model
download. After the download and warm-up, a short on-device benchmark measures the expected
per-screen overhead and shows it on the same screen.
Protected categories, pseudonymize-vs-redact, and the placeholder format are configurable in
Settings → Privacy .
Add the server to your .mcp.json configuration file:
{
"mcpServers" : {
"android-phone" : {
"type" : " http " ,
"url" : " ht

[truncated]
