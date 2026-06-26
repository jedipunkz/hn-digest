---
source: "https://github.com/maxoliverbr/vcupid-plugin"
hn_url: "https://news.ycombinator.com/item?id=48685860"
title: "VCupid Skills – AI Fundraising Toolkit for Founders"
article_title: "GitHub - maxoliverbr/vcupid-plugin · GitHub"
author: "maxoliverbr"
captured_at: "2026-06-26T12:47:35Z"
capture_tool: "hn-digest"
hn_id: 48685860
score: 3
comments: 0
posted_at: "2026-06-26T12:28:16Z"
tags:
  - hacker-news
  - translated
---

# VCupid Skills – AI Fundraising Toolkit for Founders

- HN: [48685860](https://news.ycombinator.com/item?id=48685860)
- Source: [github.com](https://github.com/maxoliverbr/vcupid-plugin)
- Score: 3
- Comments: 0
- Posted: 2026-06-26T12:28:16Z

## Translation

タイトル: VCupid Skills – 創業者向け AI 資金調達ツールキット
記事のタイトル: GitHub - maxoliverbr/vcupid-plugin · GitHub
説明: GitHub でアカウントを作成して、maxoliverbr/vcupid-plugin の開発に貢献します。

記事本文:
GitHub - maxoliverbr/vcupid-plugin · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
マクオリバーbr
/
vcupid プラグイン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く

フォルダーとファイル
13 コミット 13 コミット .claude-plugin .claude-plugin アセット アセット スキル スキル .gitignore .gitignore ライセンス ライセンス README.md README.md STARTUP_PROFILE.md STARTUP_PROFILE.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
VCupid Skills — 創業者向け AI 資金調達ツールキット
スタートアップ プロフィールを完全な VC 資金調達ワークフローに変える一連のスキル: パイプライン調査、資金マッチング、パートナー プロファイリング、アウトリーチ草案作成、会議準備、会議後の報告会、ターム シート分析、戦略。
資金調達キャンペーンを完了するには、次の順序でコマンドを実行します。
# 0. STARTUP_PROFILE.md を作成します。
# 情報を含むサンプル ファイルを使用します。
#1. 戦略を立てる
/vクレイズ
# 2. 景観を調査する
/vclist
# 3. 各 Tier 1 ファンドについて、最初に正当性を確認します。
/vcposer < ファンド名 > # このファンドは本物でアクティブですか?スコアが 40 未満の場合はドロップします。
#4. vcposer に合格した各ファンド (スコア 60+):
/vcmatch < ファンド名 > # 彼らの使命はあなたのスタートアップに適合しますか?
/vcperks < ファンド名 > # 小切手以外に何を提供しているのでしょうか?
#5. あなたが追求している資金に対してあなたの提案をストレステストします。
/vcdevil # 10 の致命的な質問 — 何を擁護すべきかを知る
/vcangel # 10 の最強のシグナル — 何をリードすべきかを知る
#6. vcmatch の追求またはウォームアップの判定が行われた各ファンドについて:
/vcpartner < ファンド > < パートナー名 >
/vcintro vcmatch- < ファンド > .md # アウトリーチを送信
/vclp vcmatch- < 資金 > .md # 概要資料を添付します
#7. 会議が予約されている場合:
/vcprep vcmatch- < ファンド > .md
#8. 会議の後:
/vcdebrief vcmatch- < ファンド > .md # シグナル読み取り、フォローアップ電子メール、次のステップ
# 9. タームシートを受け取ったら:
/vcterm termsheet- < Fund > .md # 条項分析、スコア、交渉優先順位リスト
# いつでも — パイプライン全体を確認します。
/vctrack
ファイルの命名規則:
これらのコマンド

は、vcupid Claude Code プラグインとして利用できます。単一のコマンドでインストール — Claude Code Web または CLI で動作します。
カール -fsSL https://raw.githubusercontent.com/maxoliverbr/vcupid-plugin/main/install.sh |バッシュ
このスクリプトは、プラグインを ~/.claude/plugins/vcupid/ に複製し、 ~/.claude/plugins/installed_plugins.json に登録し、 ~/.claude/settings.json で有効にします。これは冪等であるため、再実行して更新しても安全です。
クロードコードを再起動します。すべての /vc* コマンドは、 STARTUP_PROFILE.md が含まれる任意のディレクトリで使用できます。
リポジトリをローカルにクローンした場合は、クローンから直接インストールできます。
cd ~ /dev/vcupid-plugin && bash install.sh
これにより、ローカル クローン パスが登録されるため、インストーラーを再実行しなくても、スキル ファイルへの編集がすぐに有効になります。
~/.claude/plugins/installed_plugins.json に登録します。
"vcupid@local" : [{
"スコープ" : "ユーザー" ,
"installPath" : " /home/<you>/dev/vcupid " ,
"バージョン" : " 1.0.0 " 、
"installedAt" : " <ISO タイムスタンプ> " ,
"lastUpdated" : " <ISO タイムスタンプ> "
}]
~/.claude/settings.json で有効にします。
"有効なプラグイン" : {
"vcupid@local" : true
}
ヒント
STARTUP_PROFILE.md を最新の状態に保ちます。新しいトラクションマイルストーン（LOIの署名、新しいパートナーシップ、最初の収益）ごとに、アプローチできる資金や発言できる内容が変わります。更新して /vcraise を毎月再実行します。
アウトリーチに時間を費やす前に /vcmatch を実行してください。適合スコアが 40/100 の場合はダメです。メールは書かないでください。
ウォーム パスが存在する場合は常に、コールド電子メール上でバリアント B ( /vcintro ) を使用します。転送されたイントロは、コールド アウトリーチの 5 ～ 10 倍のレートでコンバージョンします。
vcprep はリハーサル用であり、部屋用ではありません。印刷してリハーサルし、そのまま残してください。会議は発表会ではなく会話です。
会議と同じ日に /vcdebrief を実行します。記憶は急速に薄れ、フォローアップメールは 24 時間後には影響力を失います。する

翌朝まで待てません。
/vcterm の出力を弁護士の代わりに共有するのではなく、弁護士と共有してください。タームシートの分析は、何のために戦うべきかを教えてくれます。あなたの弁護士は、契約を破棄せずにそれのために戦う方法を教えてくれます。
毎週のレビューの前に /vctrack を実行します。 5 つ以上のファンドを並行して運用している場合、パイプライン テーブルは、何が古くなり、何がナッジを必要としているのかを確認する唯一の方法です。
スキルを追加または変更したいですか?以下の「貢献」を参照してください。
すべてのコマンドは、現在の作業ディレクトリから STARTUP_PROFILE.md を読み取ります。コマンドを実行する前にこのファイルを作成してください。
プロジェクト ルートに次の構造のマークダウン ファイルを作成します。詳細を指定するほど、すべてのコマンド出力がより具体的かつ正確になります。曖昧な入力は一般的な出力を生成します。
# スタートアッププロフィール: [会社名]
## I. 会社概要
|フィールド |詳細 |
| ------- | -------- |
| ** スタートアップ名 ** | [ 法人名 ] |
| ** ワンライナー ** | [ 一文: 誰のために何をするか ] |
| ** 核心的な問題 ** | [ 解決する具体的な問題 — 規模や緊急性など ] |
| ** 解決策 ** | [ 何を構築し、どのように機能するか - 関連する場合は技術的な内容を記載してください ] |
| ** キー値プロパティ ** | [ 成果の見出し: 時間の節約、コストの削減など ] |
| ** カテゴリー ** | [ 例: ディープテック / インフラストラクチャ / 気候変動 / SaaS ] |
| ** 所在地と設立 ** | [ 都市、州、国。月、年。 ] |
| ** ウェブサイト ** | [ URL ] |
## II.市場とビジネスモデル
|フィールド |詳細 |
| ------- | -------- |
| ** 対象顧客 ** | [ 特定の購入者 — 「企業」ではなく、「電力協同組合の送電事業者」 ] |
| ** 市場機会 ** | [ 橋頭堡の TAM と情報源。拡張パス。 ] |
| ** ビジネスモデル ** | [ 請求方法: SaaS、価値ベース、使用量、マーケットプレイスなど ] |
＃＃III。トラクションとマイルストーン
|フィ

フィールド |詳細 |
| ------- | -------- |
| ** 製品段階 ** | [ アイデア / MVP / MVBP / 収益 / 成長 ] |
| ** 製品発表 ** | [ 対象の日付または "[日付] 時点で有効" ] |
| ** トラクション ** | [ 具体的に: デザイン パートナー、LOI、収益、試験運用、補助金、政府パートナーシップ ] |
| ** 紹介 ** | [ VC のドアを開けてくれる心温まる連絡先 — 名前、メールアドレス、関係性 ] |
## IV.チーム
|役割 |名前 |主なハイライト |
| ------ | ------ | --------------- |
| ** CEO ** | [ 名前 ] | [ 投資家にとって最も関連性の高い資格情報 3 ～ 5 ] |
| ** CTO ** | [ 名前 ] | [ 投資家にとって最も関連性の高い資格情報 3 ～ 5 ] |
| [ その他 ] | [ 名前 ] | [資格情報] |
## V. 募金活動
|フィールド |詳細 |
| ------- | -------- |
| ** 調達額 ** | [ $XM ] |
| ** 楽器 ** | [ 安全 / シード価格 / 未定 ] |
| ** 資金の使途 ** | [ 上位 3 項目 ] |
| ** この資金調達のマイルストーン ** | [ お金がなくなったら何を達成するか ] |
強力なプロフィールを作成するためのヒント:
「500 万ドルの DOE 提案」、「38,000 人のメンバーを持つ LinkedIn グループ」、「3,020 万ドルの送信プロジェクト」など、どこでも特定の数字を使用します。「大規模な DOE 提案」ではありません。
紹介者に連絡先情報を付けて名前を付けます。アウトリーチ コマンドではそれらの情報が使用されます。
あなたが構築したソリューションだけでなく、VC が感じる緊急性を含めて問題を説明してください
完全な履歴書ではなく、投資家にとって重要な資格情報をリストアップします
/vcraise — 募金戦略メモ
ここから始めましょう。プロファイルと既存の vcmatch/vclist ファイルを完全な戦略に合成します。
/vクレイズ
読み取り: STARTUP_PROFILE.md + 現在のディレクトリ内の任意の vclist.md および vcmatch-*.md
保存: vcraise.md
現在の状態 — あなたが持っているものと欠けているものについての正直な評価
引き上げパラメータ — 推奨額、手段、評価上限、タイムライン
シーケンス戦略 — 誰が最初に、並列か逐次か、どのように作成するか

緊急です
マイルストーン マップ — どの証拠ポイントが投資家のどの層のロックを解除するか
月ごとのタイムライン — 今日から年末まで
リスク要因と不測の事態 — それぞれの昇給とフォールバックの妨げとなるもの
今週 — 次の 3 つのアクション — 具体的、具体的、7 日間で実行可能
VC の状況を調査し、ターゲットとなる 15 ～ 20 のファンドのランク付けリストを作成します。
/vclist
読み取り: STARTUP_PROFILE.md
保存: vclist.md
出力: 根拠とファンドごとの推奨接触角を含む 3 層のファンド:
Tier 1 — 今すぐ追求: 強いフィット感、正しい段階、暖かいパスが存在します
Tier 2 — 最初にウォームアップ: 論文の適合性は良好ですが、証明ポイントや導入部分が欠落しています
Tier 3 — モニター: ステージの不一致または不明瞭なパス — 後で再確認します
/vcmatch の前にこれを実行して、どのファンドが詳細な分析に値するかを確認します。
完全な一致分析に取り組む前に、パイプラインから資金を事前にフィルタリングします。ファンドが実際に活動しており、小切手を書いているかどうかについて、論文があなたのスタートアップに適合するかどうかに関係なく、8 つの証拠に基づくチェックを実行します。
/vcposer <VC ファンド名>
例: /vcposer "Slow Ventures" または /vcposer a16z
読み取り: STARTUP_PROFILE.md (ステージ/セクターのキャリブレーション用)
保存: vcposer-<fund>.md
出力 — 8 つのチェックにわたる Poser スコア (0 ～ 100):
ファンドの活力 — 最終投資日。 18 か月以上沈黙 = ゾンビのリスク。
ファンドの健全性 — ドライパウダーを使用したアクティブファンド?主要パートナーの離脱?
論文の信頼性 — 彼らのポートフォリオは、実際に述べられたセクターの論文と一致していますか?
リード vs. フォロー — 単なる共同投資ではなく、ラウンドをリードしている証拠。
ステージの正直さ — 最近の投資の記載されたステージと実際のエントリーステージ。
サイズの現実性を確認する — 記載された範囲と観察可能な取引サイズ。
Signal-to-Check Ratio — カンファレンス/コンテンツのボリュームと実際の新規取引。
創設者の感情 — パブリックフィードバックパターン — ゴースティング、遅い意思決定

、反逆する。
合法 (80–100): アクティブなファンド、本物の小切手。 /vcmatch に進みます。
可能性の高い (60 ～ 79): 黄色のフラグ — 適合性が高い場合に優先順位を付け、未解決の疑問を明確にします。
ウォッチ リスト (40 ～ 59): マテリアル ポーザー信号。温かみのあるイントロが存在しない限り、優先順位を下げます。
Poser の可能性 (0–39): サイクルを無駄にしないでください。次のファンドに移ります。
vcposer ≠ vcmatch。 Poser Score は、「このファンドは本物でアクティブですか?」という質問に答えます。 Vcmatch は、「彼らの使命は私たちのスタートアップに適合しますか?」と答えます。両方必要です。ファンドは完全に合法であっても、あなたの段階では間違っている可能性があります。また、ファンドは完璧な論文を主張していても、ゾンビである可能性があります。スコアが 60 以上のファンドに対してのみ、最初に /vcposer を実行し、次に /vcmatch を実行します。
あなたのスタートアッププロフィールに対して単一のファンドを徹底的に調査します。 4 つの並列ドメイン調査サブエージェントを生成し、結果を構造化レポートに合成し、8 つの次元にわたる一致を合計して /100 になるスコアを付けます。
/vcmatch <VC ファンド名>
例: /vcmatch a16z または /vcmatch "Lowercarbon Capital"
読み取り: STARTUP_PROFILE.md
保存: vcmatch-<fund>.md
仕組み: 4 つのサブエージェント (論文と委任、ポートフォリオ インテリジェンス、人材とネットワーク、取引とプロセス) が並行して実行され、それぞれがドメイン固有の検索クエリを使用します。メイン エージェントは 4 つの戻り値すべてを合成し、適合度をスコア付けします。
企業スナップショット — 設立、モデル、主要 GP、標準的な取引、ファンドの規模、注目すべきポートフォリオ
投資計画

[切り捨てられた]

## Original Extract

Contribute to maxoliverbr/vcupid-plugin development by creating an account on GitHub.

GitHub - maxoliverbr/vcupid-plugin · GitHub
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
maxoliverbr
/
vcupid-plugin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
13 Commits 13 Commits .claude-plugin .claude-plugin assets assets skills skills .gitignore .gitignore LICENSE LICENSE README.md README.md STARTUP_PROFILE.md STARTUP_PROFILE.md install.sh install.sh View all files Repository files navigation
VCupid Skills — AI Fundraising Toolkit for Founders
A set of Skills that turn your startup profile into a full VC fundraising workflow: pipeline research, fund matching, partner profiling, outreach drafting, meeting prep, post-meeting debrief, term sheet analysis, and strategy.
Run commands in this order for a complete fundraising campaign:
# 0. Create your STARTUP_PROFILE.md
# Use the example file with your information.
# 1. Build the strategy
/vcraise
# 2. Research the landscape
/vclist
# 3. For each Tier 1 fund — legitimacy check first:
/vcposer < fund name > # Is this fund real and active? Drop if score < 40.
# 4. For each fund that passes vcposer (score 60+):
/vcmatch < fund name > # Does their mandate fit your startup?
/vcperks < fund name > # What do they offer beyond the check?
# 5. Stress-test your pitch against the funds you're pursuing:
/vcdevil # 10 lethal questions — know what to defend
/vcangel # 10 strongest signals — know what to lead with
# 6. For each fund with a vcmatch Pursue or Warm Up verdict:
/vcpartner < fund > < partner name >
/vcintro vcmatch- < fund > .md # send outreach
/vclp vcmatch- < fund > .md # attach the one-pager
# 7. When a meeting is booked:
/vcprep vcmatch- < fund > .md
# 8. After the meeting:
/vcdebrief vcmatch- < fund > .md # signal read, follow-up email, next step
# 9. When you receive a term sheet:
/vcterm termsheet- < fund > .md # clause analysis, score, negotiation priority list
# Anytime — review your full pipeline:
/vctrack
File naming convention:
These commands are available as the vcupid Claude Code plugin. Install with a single command — works in Claude Code web or CLI:
curl -fsSL https://raw.githubusercontent.com/maxoliverbr/vcupid-plugin/main/install.sh | bash
The script clones the plugin to ~/.claude/plugins/vcupid/ , registers it in ~/.claude/plugins/installed_plugins.json , and enables it in ~/.claude/settings.json . It is idempotent — safe to re-run to update.
Restart Claude Code. All /vc* commands will be available in any directory that contains a STARTUP_PROFILE.md .
If you've cloned the repo locally, you can install from the clone directly:
cd ~ /dev/vcupid-plugin && bash install.sh
This registers the local clone path so edits to skill files take effect immediately without re-running the installer.
Register in ~/.claude/plugins/installed_plugins.json :
"vcupid@local" : [{
"scope" : " user " ,
"installPath" : " /home/<you>/dev/vcupid " ,
"version" : " 1.0.0 " ,
"installedAt" : " <ISO timestamp> " ,
"lastUpdated" : " <ISO timestamp> "
}]
Enable in ~/.claude/settings.json :
"enabledPlugins" : {
"vcupid@local" : true
}
Tips
Keep STARTUP_PROFILE.md current. Every new traction milestone (signed LOI, new partnership, first revenue) changes what funds you can approach and what you can say. Update it and re-run /vcraise monthly.
Run /vcmatch before spending time on outreach. A 40/100 fit score is a no-go — don't write the email.
Use Variant B ( /vcintro ) over cold email whenever a warm path exists. A forwarded intro converts at 5–10x the rate of cold outreach.
The vcprep is for rehearsal, not for the room. Print it, rehearse it, then leave it behind — the meeting is a conversation, not a recital.
Run /vcdebrief the same day as the meeting. Memory fades fast and the follow-up email loses impact after 24 hours. Don't wait until the next morning.
Share /vcterm output with your attorney, not instead of one. The term sheet analysis tells you what to fight for — your lawyer tells you how to fight for it without blowing the deal.
Run /vctrack before every weekly review. If you're working 5+ funds in parallel, the pipeline table is the only way to see what's stale and what needs a nudge.
Want to add or change a skill? See Contributing below.
All commands read STARTUP_PROFILE.md from your current working directory. Create this file before running any command.
Create a markdown file in your project root with the following structure. The more detail you provide, the more specific and accurate every command output will be. Vague inputs produce generic outputs.
# Startup Profile: [ Company Name ]
## I. Company Overview
| Field | Detail |
| ------- | -------- |
| ** Startup Name ** | [ Legal company name ] |
| ** One-Liner ** | [ One sentence: what you do and for whom ] |
| ** Core Problem ** | [ The specific problem you solve — include scale/urgency ] |
| ** Solution ** | [ What you built and how it works — be technical if relevant ] |
| ** Key Value Prop ** | [ The headline outcome you deliver: time saved, cost reduced, etc. ] |
| ** Category ** | [ e.g., Deep Tech / Infrastructure / Climate / SaaS ] |
| ** Location & Founded ** | [ City, State, Country. Month Year. ] |
| ** Website ** | [ URL ] |
## II. Market & Business Model
| Field | Detail |
| ------- | -------- |
| ** Target Customer ** | [ Specific buyer — not "enterprises", but "grid operators at electric cooperatives" ] |
| ** Market Opportunity ** | [ Beachhead TAM with source. Expansion path. ] |
| ** Business Model ** | [ How you charge: SaaS, value-based, usage, marketplace, etc. ] |
## III. Traction & Milestones
| Field | Detail |
| ------- | -------- |
| ** Product Stage ** | [ Idea / MVP / MVBP / Revenue / Growth ] |
| ** Product Launch ** | [ Target date or "Live as of [ date ] " ] |
| ** Traction ** | [ Be specific: design partners, LOIs, revenue, pilots, grants, gov partnerships ] |
| ** Referral ** | [ Any warm contacts who can open VC doors — Name, email, relationship ] |
## IV. Team
| Role | Name | Key Highlights |
| ------ | ------ | --------------- |
| ** CEO ** | [ Name ] | [ 3–5 most relevant credentials for investors ] |
| ** CTO ** | [ Name ] | [ 3–5 most relevant credentials for investors ] |
| [ Other ] | [ Name ] | [ Credentials ] |
## V. Fundraise
| Field | Detail |
| ------- | -------- |
| ** Raise Amount ** | [ $XM ] |
| ** Instrument ** | [ SAFE / Priced Seed / TBD ] |
| ** Use of Funds ** | [ Top 3 line items ] |
| ** Milestone this raise funds ** | [ What you'll have accomplished when the money runs out ] |
Tips for a strong profile:
Use specific numbers everywhere: "$5M DOE proposal", "38,000-member LinkedIn group", "$30.2M transmission project" — not "large DOE proposal"
Name your referrals with contact info — the outreach commands use them
Describe the problem with the urgency a VC would feel, not just the solution you built
List credentials that matter to investors, not your full resume
/vcraise — Fundraising Strategy Memo
Start here. Synthesizes your profile and any existing vcmatch/vclist files into a full strategy.
/vcraise
Reads: STARTUP_PROFILE.md + any vclist.md and vcmatch-*.md in the current directory
Saves: vcraise.md
Current State — honest assessment of what you have and what's missing
Raise Parameters — recommended amount, instrument, valuation cap, timeline
Sequencing Strategy — who first, parallel vs. sequential, how to create urgency
Milestone Map — which proof points unlock which tier of investors
Month-by-Month Timeline — from today through close
Risk Factors & Contingencies — what derails the raise and the fallback for each
This Week — Next 3 Actions — concrete, specific, doable in 7 days
Research the VC landscape and produce a ranked list of 15–20 funds to target.
/vclist
Reads: STARTUP_PROFILE.md
Saves: vclist.md
Output: Three tiers of funds with rationale and suggested contact angle per fund:
Tier 1 — Pursue Now: Strong fit, correct stage, warm path exists
Tier 2 — Warm Up First: Good thesis fit, missing proof point or intro
Tier 3 — Monitor: Stage mismatch or unclear path — revisit later
Run this before /vcmatch to know which funds are worth the deep analysis.
Pre-filter funds from your pipeline before committing to a full match analysis. Runs 8 evidence-based checks on whether a fund is actually active and writing checks — independent of whether their thesis fits your startup.
/vcposer <VC Fund Name>
Example: /vcposer "Slow Ventures" or /vcposer a16z
Reads: STARTUP_PROFILE.md (for stage/sector calibration)
Saves: vcposer-<fund>.md
Output — Poser Score (0–100) across 8 checks:
Fund Vitality — Last investment date. >18 months silent = zombie risk.
Fund Health — Active fund with dry powder? Key partner departures?
Thesis Authenticity — Does their portfolio actually match their stated sector thesis?
Lead vs. Follow — Evidence of leading rounds, not just co-investing.
Stage Honesty — Stated stage vs. actual entry stage of recent investments.
Check Size Reality — Stated range vs. observable deal sizes.
Signal-to-Check Ratio — Conference/content volume vs. actual new deals.
Founder Sentiment — Public feedback patterns — ghosting, slow decisions, reneging.
Legit (80–100): Active fund, real checks. Proceed to /vcmatch .
Probable (60–79): Yellow flags — prioritize if fit is strong, clarify open questions.
Watch List (40–59): Material poser signals. De-prioritize unless a warm intro exists.
Likely Poser (0–39): Don't waste cycles. Move to the next fund.
vcposer ≠ vcmatch. Poser Score answers "is this fund real and active?" Vcmatch answers "does their mandate fit our startup?" You need both. A fund can be fully legit but wrong for your stage — and a fund can claim a perfect thesis but be a zombie. Run /vcposer first, then /vcmatch only for funds that score 60+.
Deep research on a single fund against your startup profile. Spawns 4 parallel domain research sub-agents, synthesizes the results into a structured report, and scores the match across 8 dimensions summing to /100.
/vcmatch <VC Fund Name>
Example: /vcmatch a16z or /vcmatch "Lowercarbon Capital"
Reads: STARTUP_PROFILE.md
Saves: vcmatch-<fund>.md
How it works: Four sub-agents run in parallel — Thesis & Mandate, Portfolio Intelligence, People & Network, and Deal & Process — each with domain-specific search queries. The main agent synthesizes all four returns and scores the fit.
Firm Snapshot — founded, model, key GPs, standard deal, fund size, notable portfolio
Investment Th

[truncated]
