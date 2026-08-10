---
source: "https://github.com/anthropics/claude-code/commit/66edf5358349356774812264b75b8ea792f0d0a3"
hn_url: "https://news.ycombinator.com/item?id=49248636"
title: "Claude Code no longer limits sessions to 200 spawned subagents"
article_title: "chore: Update CHANGELOG.md and feed.xml · anthropics/claude-code@66edf53 · GitHub"
author: "bakigul"
captured_at: "2026-08-10T19:47:51Z"
capture_tool: "hn-digest"
hn_id: 49248636
score: 1
comments: 0
posted_at: "2026-08-10T19:39:40Z"
tags:
  - hacker-news
  - translated
---

# Claude Code no longer limits sessions to 200 spawned subagents

- HN: [49248636](https://news.ycombinator.com/item?id=49248636)
- Source: [github.com](https://github.com/anthropics/claude-code/commit/66edf5358349356774812264b75b8ea792f0d0a3)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T19:39:40Z

## Translation

タイトル: クロード コードは、セッションを生成されるサブエージェント 200 に制限しなくなりました
記事のタイトル: chore: CHANGELOG.md と feed.xml を更新する · anthropics/claude-code@66edf53 · GitHub
説明: Claude Code は、端末内に常駐し、コードベースを理解し、自然言語コマンドを通じて日常的なタスクの実行、複雑なコードの説明、git ワークフローの処理を行うことで、コーディングを高速化するのに役立つエージェント コーディング ツールです。 - 雑用: CHANGELOG.md と feed.xml を更新 · anthropics/clau
[切り捨てられた]

記事本文:
雑務: CHANGELOG.md と feed.xml を更新する · anthropics/claude-code@66edf53 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
人類学
/
クロードコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ファイルを参照 履歴のこの時点のリポジトリを参照 ファイルを参照 アクション - ユーザーがコミットした作業: CHANGELOG.md および feed.xml を更新 1 つの親 5cf69b1 コミット 66edf53 66edf53 の完全な SHA をコピー 2 ファイルが変更されました
ファイルを展開する

ツリー ファイル ツリーを折りたたむ 差分表示設定を開く フィルター オプション CHANGELOG.md
ファイルツリーを展開する ファイルツリーを折りたたむ 差分表示設定を開く ファイルを折りたたむ CHANGELOG.md
ファイル名をクリップボードにコピー 全行展開：CHANGELOG.md + 34 変更行：追加34行、削除0行 ソースの差分を表示
元のファイルの行番号 差分行番号 差分行の変更 @@ -1,5 +1,39 @@ 1 1 # 変更ログ 2 2
3 + ## 2.1.224 4 + 5 + - セルフホスト環境の追加: ` claude self-hosted-runner ` は、独自のマシンまたはコンテナを、Team および Enterprise プランで、Claude Code の Web、モバイル、デスクトップ セッションを実行できる場所に変えます 6 + - ` archive ` プラグイン ソースを追加: git または npm を使用せず、HTTPS 経由で zip からプラグインをインストールし、オプションの SHA-256 ピンニングを使用します 7 + - を追加しました使用できないペーストを削除するとコマンドのテキストが変更されるキャンセルと確認のステップ 8 + - ` AWS_REGION ` から派生したものよりも特定のクロスリージョン推論プロファイルを優先するために、Bedrock の ` ANTHROPIC_BEDROCK_REGION_PREFIX ` 環境変数を追加しました 9 + - `crossSessionInbound ` および `dialogExpiry ` 設定を追加しました: バイパスされたアクセス許可で実行されているセッションにクロスセッション メッセージが送信されます承認のために保持され、他のセッションへのメッセージは自動配信 10 + - サンドボックス認証情報マスキング オプションを追加しました: 構造化環境値の `extract ` および `onExtractNoMatch `、JWT 対応マスキングの `maskClaims ` を使用した `decode: "jwt" `、および AWS SigV4 の `awsPairs ` / ` sigv4 `再署名。これらには ` network.tlsTerminate ` が必要で、ユーザー設定、管理対象設定、または ` --settings ` 設定からのみ受け入れられます 11 + - クロスセッション ` SendMessage ` を追加しました。クロード コード セッションは、どのマシンでも、` ListAgents ` を使用して相互にメッセージを送信できるようになり、それらを検出できるようになりました (macOS および Linux) 12 + - 別のプロジェクトの s に解決される長い (>200 文字) プロジェクト パスを修正しました

共有サニタイズされたプレフィックスの下の ession ディレクトリ。セッション リスト、名前変更、フォーク、削除、および ` /resume ` がプロジェクトをまたがることはなくなりました 13 + - チームメイトの受信箱への書き込みが実際に失敗したときに「メッセージが送信されました」と報告する ` SendMessage ` を修正しました。失敗した配信はエラーとして報告されるようになりました 14 + - Linux および macOS でサイレントにバイパスできる、末尾のスラッシュで書かれたサンドボックス ファイルシステムの拒否エントリ (例: `denyRead: "~/.aws/" ` ) を修正しました 15 + - Bash ツールにサンドボックス違反の詳細が表示されない問題を修正しました
[切り捨てられた]
ファイル名をクリップボードにコピー すべての行を展開: feed.xml + 38 - 44 変更された行: 38 追加 & 44 削除 元のファイル行番号 差分行番号 差分行変更 @@ -6,7 +6,44 @@ 6 6 < author >< name >Anthropic</ name ></ author > 7 7 < link rel = " alter " type = " text/html " href = " https://github.com/anthropics/claude-code/blob/main/CHANGELOG.md " /> 8 8 < link rel = " self " type = " application/atom+xml " href = " https://raw.githubusercontent.com/anthropics/claude-code/main/feed.xml " /> 9 - < updated >2026-08-06T00:52:31Z</更新済み > 9 + < 更新済み >2026-08-07T04:00:51Z</ 更新済み > 10 + < エントリー > 11 + < id >https://github.com/anthropics/claude-code/releases/tag/v2.1.224</ id > 12 + < title >クロード コード v2.1.224</ title > 13 + < link rel = " alter " type = " text/html " href = " https://github.com/anthropics/claude-code/releases/tag/v2.1.224 " /> 14 + < updated >2026-08-07T04:00:51Z</ updated > 15 + < content type = " html " > < p > • セルフホスト環境の追加: claude self-hosted-runner は、自分のマシンまたはコンテナを場所に変えます Claudeコード Web、モバイル、およびデスクトップ セッションは、チーム プランおよびエンタープライズ プランで実行できます。 < /p > 16 + < p > • アーカイブ プラグイン ソースの追加: git o を使用せずに、HTTPS 経由で zip からプラグインをインストールします。

r npm、オプションの SHA-256 ピンニングあり </p > 17 + < p > • 使用できないペーストを削除するとコマンドのテキストが変更される場合のキャンセルと確認のステップを追加しました </p > 18 + < p > • AWS_REGION 由来の推論プロファイルよりも特定のクロスリージョン推論プロファイルを優先するために、Bedrock の ANTHROPIC_BEDROCK_REGION_PREFIX 環境変数を追加しました </p > 19 + < p > • 追加しましたCrossSessionInbound および DialogExpiry 設定: バイパスされたアクセス許可で実行されているセッションに送信されたセッション間メッセージは承認のために保留され、他のセッションへのメッセージは自動配信されます。 < /p > 20 + < p > • 追加されたサンドボックス資格情報マスク オプション: 構造化された環境値の extract および onExtractNoMatch、デコード: mas 付きの "jwt"
[切り捨てられた]
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Code is an agentic coding tool that lives in your terminal, understands your codebase, and helps you code faster by executing routine tasks, explaining complex code, and handling git workflows - all through natural language commands. - chore: Update CHANGELOG.md and feed.xml · anthropics/clau
[truncated]

chore: Update CHANGELOG.md and feed.xml · anthropics/claude-code@66edf53 · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
anthropics
/
claude-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Browse files Browse the repository at this point in the history Browse files actions-user committed chore: Update CHANGELOG.md and feed.xml 1 parent 5cf69b1 commit 66edf53 Copy full SHA for 66edf53 2 file s changed
Expand file tree Collapse file tree Open diff view settings Filter options CHANGELOG.md
Expand file tree Collapse file tree Open diff view settings Collapse file ‎ CHANGELOG.md ‎
Copy file name to clipboard Expand all lines: CHANGELOG.md + 34 Lines changed: 34 additions & 0 deletions Display the source diff
Original file line number Diff line number Diff line change @@ -1,5 +1,39 @@ 1 1 # Changelog 2 2
3 + ## 2.1.224 4 + 5 + - Added self-hosted environments: ` claude self-hosted-runner ` turns your own machines or containers into a place Claude Code web, mobile, and desktop sessions can run, on Team and Enterprise plans 6 + - Added ` archive ` plugin source: install plugins from a zip over HTTPS without git or npm, with optional SHA-256 pinning 7 + - Added a cancel-and-confirm step when removing an unavailable paste changes a command's text 8 + - Added ` ANTHROPIC_BEDROCK_REGION_PREFIX ` env var for Bedrock to prefer a specific cross-region inference profile over the ` AWS_REGION ` -derived one 9 + - Added ` crossSessionInbound ` and ` dialogExpiry ` settings: cross-session messages sent to a session running with bypassed permissions are held for your approval, and messages to other sessions auto-deliver 10 + - Added sandbox credential-masking options: ` extract ` and ` onExtractNoMatch ` for structured env values, ` decode: "jwt" ` with ` maskClaims ` for JWT-aware masking, and ` awsPairs ` / ` sigv4 ` for AWS SigV4 re-signing; these need ` network.tlsTerminate ` and are honored only from user, managed, or ` --settings ` settings 11 + - Added cross-session ` SendMessage ` : Claude Code sessions can now message each other, on any of your machines, with ` ListAgents ` to discover them (macOS and Linux) 12 + - Fixed long (>200 char) project paths resolving to another project's session directory under a shared sanitized prefix; session list, rename, fork, delete and ` /resume ` no longer cross projects 13 + - Fixed ` SendMessage ` reporting "Message sent" when the write to a teammate's inbox had actually failed; failed deliveries are now reported as errors 14 + - Fixed sandbox filesystem deny entries written with a trailing slash (e.g. ` denyRead: "~/.aws/" ` ) being silently bypassable on Linux and macOS 15 + - Fixed sandbox violation details never appearing in Bash tool
[truncated]
Copy file name to clipboard Expand all lines: feed.xml + 38 - 44 Lines changed: 38 additions & 44 deletions Original file line number Diff line number Diff line change @@ -6,7 +6,44 @@ 6 6 < author >< name >Anthropic</ name ></ author > 7 7 < link rel = " alternate " type = " text/html " href = " https://github.com/anthropics/claude-code/blob/main/CHANGELOG.md " /> 8 8 < link rel = " self " type = " application/atom+xml " href = " https://raw.githubusercontent.com/anthropics/claude-code/main/feed.xml " /> 9 - < updated >2026-08-06T00:52:31Z</ updated > 9 + < updated >2026-08-07T04:00:51Z</ updated > 10 + < entry > 11 + < id >https://github.com/anthropics/claude-code/releases/tag/v2.1.224</ id > 12 + < title >Claude Code v2.1.224</ title > 13 + < link rel = " alternate " type = " text/html " href = " https://github.com/anthropics/claude-code/releases/tag/v2.1.224 " /> 14 + < updated >2026-08-07T04:00:51Z</ updated > 15 + < content type = " html " > &lt; p &gt; • Added self-hosted environments: claude self-hosted-runner turns your own machines or containers into a place Claude Code web, mobile, and desktop sessions can run, on Team and Enterprise plans &lt; /p &gt; 16 + &lt; p &gt; • Added archive plugin source: install plugins from a zip over HTTPS without git or npm, with optional SHA-256 pinning &lt; /p &gt; 17 + &lt; p &gt; • Added a cancel-and-confirm step when removing an unavailable paste changes a command's text &lt; /p &gt; 18 + &lt; p &gt; • Added ANTHROPIC_BEDROCK_REGION_PREFIX env var for Bedrock to prefer a specific cross-region inference profile over the AWS_REGION-derived one &lt; /p &gt; 19 + &lt; p &gt; • Added crossSessionInbound and dialogExpiry settings: cross-session messages sent to a session running with bypassed permissions are held for your approval, and messages to other sessions auto-deliver &lt; /p &gt; 20 + &lt; p &gt; • Added sandbox credential-masking options: extract and onExtractNoMatch for structured env values, decode: "jwt" with mas
[truncated]
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
