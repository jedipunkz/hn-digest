---
source: "https://github.com/assafkip/investor-recon"
hn_url: "https://news.ycombinator.com/item?id=48494224"
title: "Show HN: Investor-recon – read a VC's real thesis before you pitch (Claude Code)"
article_title: "GitHub - assafkip/investor-recon: A threat-intelligence approach to VC meetings. Six prompts, ~25 min per investor. Built with Claude Code. · GitHub"
author: "Assafkip"
captured_at: "2026-06-11T19:05:29Z"
capture_tool: "hn-digest"
hn_id: 48494224
score: 1
comments: 0
posted_at: "2026-06-11T18:12:19Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Investor-recon – read a VC's real thesis before you pitch (Claude Code)

- HN: [48494224](https://news.ycombinator.com/item?id=48494224)
- Source: [github.com](https://github.com/assafkip/investor-recon)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T18:12:19Z

## Translation

タイトル: Show HN: Investor-recon – 売り込む前に VC の本当の論文を読んでください (Claude Code)
記事のタイトル: GitHub - assafkip/investor-recon: VC ミーティングへの脅威インテリジェンス アプローチ。 6 つのプロンプト、投資家あたり最大 25 分。クロードコードで構築。 · GitHub
説明: VC 会議への脅威インテリジェンス アプローチ。 6 つのプロンプト、投資家あたり最大 25 分。クロードコードで構築。 - assafkip/投資家偵察

記事本文:
GitHub - assafkip/investor-recon: VC ミーティングへの脅威インテリジェンス アプローチ。 6 つのプロンプト、投資家あたり最大 25 分。クロードコードで構築。 · GitHub
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
アサフキップ
/
投資家偵察
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .claude-plugin .claude-plugin skill/ investor-recon skill/ investor-recon .gitignore .gitignore ライセンス ライセンス PASTE-THIS.md PASTE-THIS.md README.md README.md example-brief.md example-brief.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
VC ミーティングへの脅威インテリジェンス アプローチ。クロード・コードで実行し、基金の実際の資料を調査し、でっち上げを拒否します。
この概要の対話型バージョン · claudedaddy.io
私は 12 年間、脅威インテリジェンスに取り組みました。仕事内容は簡単に説明できました。敵があなたが探していることに気づく前に、敵のイメージを構築します。彼らのパターン、ギャップ、次の動き。
VC ミーティングも同じ問題です。
ほとんどの創業者は、ピッチのリハーサルをして準備をします。それは後ろ向きです。ピッチは決定が行われる場所ではありません。基金は、あなたが入社する前にあなたについての論文を作成します。会議はそれを確認する場所です。
だからリハーサルはやめてください。偵察を実行します。
これは、実際のインターネットで基金の資料を検索し、すべての主張を引用しています。何かが見つからない場合は、「見つかりません」と表示されます。ファンドの取引内容や論文、小切手の規模などを名前から推測することは決してない。自信を持って間違ったブリーフを作成すると、間違った部屋に誘導されます。それは準備書面がないよりも悪いです。
これが、チャット ボックスではなくクロード コードで実行される理由です。 Claude Code は実際に検索してフェッチすることができます。単純なチャットでは推測できますが、わかりません。アースは商品です。
論文概要。この基金が実際に支援しているものを、彼ら自身の資料からリンク付きで引用します。彼らの言葉は、あなたが彼らの取引について読んだものとは切り離されていました。
ポートフォリオマップ。あなたがどこに当てはまるか、どのポートフォリオ企業と衝突するか、あなたが埋めるギャップ。
質問セット。引用された論文とあなたのステージから、このパートナーが尋ねる質問

。
ためらい。あなたの会社のたった 1 つの点が、この特定のファンドを停止させています。
カウンターです。彼らが質問する前に、どのように答えるか。
ソースとギャップ。使用されたすべてのリンクと、確認できなかったものの正直なリスト。
最も簡単で、インストールは必要ありません。 Claude Code を開き、PASTE-THIS.md 内のブロックをコピーして貼り付けます。2 つの質問が表示され、実行されます。完全なセットアップ手順 (クロード コードが要求するフォルダー ステップを含む) は、そのファイルに含まれています。
コマンドとして。プラグインをインストールします。
/プラグインのインストール github:assafkip/investor-recon
次に、任意のクロード コード セッションで /investor-recon を実行します。
完成した概要については、example-brief.md を参照してください。
私は創業者や小規模チーム向けにカスタム AI システムを構築しています。このリポジトリは、私が有料で行った作業の無料の軽量バージョンです。実際の昇給に結び付けたい場合、または AI システムが処理すべき別の運用上の問題を解決したい場合は、 assaf@askconsulting.io までご連絡ください。
私は 12 年間、脅威インテリジェンスに携わり、チームが同じ障害を 4 回発見して修正するのを見てきました。学習が行き詰ることはありませんでした。私はそれを定着させるツールを構築します。
これは無料版です。有料キットは claudedaddy.io にあります。
これをチームのリポジトリに接続したいですか、それともより重い仕様とレビューのパイプラインに接続しますか?それがコンサルティングです。電話を予約します。
VC ミーティングへの脅威インテリジェンス アプローチ。 6 つのプロンプト、投資家あたり最大 25 分。クロードコードで構築。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A threat-intelligence approach to VC meetings. Six prompts, ~25 min per investor. Built with Claude Code. - assafkip/investor-recon

GitHub - assafkip/investor-recon: A threat-intelligence approach to VC meetings. Six prompts, ~25 min per investor. Built with Claude Code. · GitHub
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
assafkip
/
investor-recon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .claude-plugin .claude-plugin skills/ investor-recon skills/ investor-recon .gitignore .gitignore LICENSE LICENSE PASTE-THIS.md PASTE-THIS.md README.md README.md example-brief.md example-brief.md View all files Repository files navigation
A threat-intelligence approach to VC meetings. Runs in Claude Code, searches the fund's real material, and refuses to make things up.
Interactive version of this brief · claudedaddy.io
I spent 12 years doing threat intelligence. The job was simple to describe. Build a picture of an adversary before they ever know you are looking. Their patterns, their gaps, their next move.
A VC meeting is the same problem.
Most founders prep by rehearsing their pitch. That is backwards. The pitch is not where the decision happens. The fund builds a thesis about you before you walk in. The meeting is where they confirm it.
So stop rehearsing. Run recon.
This searches the real internet for the fund's material and cites every claim. If it cannot find something, it says "Not found." It never guesses the fund's deals, thesis, or check size from the name. A confident wrong brief walks you into the wrong room. That is worse than no brief.
This is why it runs in Claude Code, not a chat box. Claude Code can actually search and fetch. A plain chat guesses, and you cannot tell. The grounding is the product.
Thesis brief. What this fund actually backs, quoted from their own material, with links. Their words separated from your read of their deals.
Portfolio map. Where you fit, which portfolio company you collide with, the gap you fill.
Question set. The questions this partner will ask, from their cited thesis and your stage.
The hesitation. The single thing about your company that makes this specific fund pause.
The counter. How you answer it before they raise it.
Sources and gaps. Every link used, and an honest list of what it could not confirm.
Easiest, no install. Open Claude Code, then copy the block in PASTE-THIS.md and paste it in. It asks you two questions and runs. Full setup steps (including the folder step Claude Code asks for) are in that file.
As a command. Install the plugin:
/plugin install github:assafkip/investor-recon
Then run /investor-recon in any Claude Code session.
See example-brief.md for what a finished brief looks like.
I build custom AI systems for founders and small teams. This repo is the free, lighter version of the work I do paid. If you want it wired into your actual raise, or a different operational problem an AI system should be handling, reach me at assaf@askconsulting.io .
I spent 12 years in threat intelligence watching teams find the same failure and fix it four times. The learning never stuck. I build tools that make it stick.
This is the free version. The paid kits live at claudedaddy.io .
Want this wired into your team's repo, or a heavier spec-and-review pipeline? That's the consulting. Book a call.
A threat-intelligence approach to VC meetings. Six prompts, ~25 min per investor. Built with Claude Code.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
