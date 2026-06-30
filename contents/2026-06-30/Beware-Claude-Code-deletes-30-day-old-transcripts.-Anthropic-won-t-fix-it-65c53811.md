---
source: "https://github.com/anthropics/claude-code/issues/62476"
hn_url: "https://news.ycombinator.com/item?id=48732846"
title: "Beware, Claude Code deletes >30 day old transcripts. Anthropic won't fix it"
article_title: "[BUG] Claude Code silently deletes conversation transcripts after 30 days by default · Issue #62476 · anthropics/claude-code · GitHub"
author: "ojura"
captured_at: "2026-06-30T14:51:47Z"
capture_tool: "hn-digest"
hn_id: 48732846
score: 4
comments: 1
posted_at: "2026-06-30T13:58:21Z"
tags:
  - hacker-news
  - translated
---

# Beware, Claude Code deletes >30 day old transcripts. Anthropic won't fix it

- HN: [48732846](https://news.ycombinator.com/item?id=48732846)
- Source: [github.com](https://github.com/anthropics/claude-code/issues/62476)
- Score: 4
- Comments: 1
- Posted: 2026-06-30T13:58:21Z

## Translation

タイトル: 注意してください、Claude Code は 30 日以上前のトランスクリプトを削除します。人間学では解決しないよ
記事のタイトル: [バグ] クロード コードは、デフォルトで 30 日後に会話の記録をサイレントに削除します · 問題 #62476 · anthropics/claude-code · GitHub
説明: プリフライト チェックリスト 既存の問題を検索しましたが、まだ報告されていません これは 1 つのバグ レポートです (バグごとに個別のレポートを提出してください) 最新バージョンの Claude Code を使用しています 何が問題ですか?クリーンアップ期間...

記事本文:
[バグ] Claude Code はデフォルトで 30 日後に会話トランスクリプトをサイレントに削除します · 問題 #62476 · anthropics/claude-code · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
人類学
/
クロードコード
公共
通知

イケーション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
[BUG] Claude Code はデフォルトで 30 日後に会話トランスクリプトをサイレントに削除します #62476
リンクをコピー 新しい問題 リンクをコピー 開く 開く [バグ] クロード コードはデフォルトで 30 日後に会話トランスクリプトをサイレントに削除します #62476 リンクをコピー ラベルのバグ 何かが機能していません 何かが機能していません 説明
発行本体アクションのプリフライト チェックリスト
既存の問題を検索しましたが、まだ報告されていません
これは 1 つのバグ レポートです (バグごとに別々のレポートを提出してください)
最新バージョンのクロードコードを使用しています
cleanupPeriodDays のデフォルトは 30 で、クロード コードはサイレントに削除します。
~/.claude/projects//.jsonl ファイルは 30 日以上経過しています
スタートアップ。初回公開、削除前の警告はありません。
設定は /config には表示されません。数か月分の会話履歴を失った
こうなっていることに気づく前に。
再現: cleanupPeriodDays を設定せずにクロード コードを 30 日以上使用します。古い
トランスクリプトが消えます。
私のマシン上の証拠: ~/.claude/history.jsonl は 14 セッション / 1,315 を示しています
3 月から 4 月にかけて 1 つのプロジェクトのプロンプトが表示されます。現在のセッションの .jsonl のみ
~/.claude/projects/ に保存されます。カットオフはデフォルトの 30 日と一致します
まさに。私の settings.json はデフォルトをオーバーライドしません。
影響: コードと git 履歴は残りますが、推論の痕跡 - 設計は残ります。
ディスカッション、デバッグコンテキスト、分析はなくなりました。研究活動のため
コンテキストは成果物です。
デフォルトを非破壊 (無効または非常に長い保持) に変更します。
最初の実行時に開示します。自動削除にはオプトインが必要です。
unlink() の代わりにゴミ箱フォルダーにソフト削除します。
/config の設定を表示します。
他の人のための回避策: "cleanupPeriodDays": 3650 を追加します。
~/.claude/settings.json i

すぐに。
cleanupPeriodDays を設定せずに、30 日を超えてクロード コードを使用します。古いトランスクリプトが消えます。
リアクションは現在利用できません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Preflight Checklist I have searched existing issues and this hasn't been reported yet This is a single bug report (please file separate reports for different bugs) I am using the latest version of Claude Code What's Wrong? cleanupPeriodD...

[BUG] Claude Code silently deletes conversation transcripts after 30 days by default · Issue #62476 · anthropics/claude-code · GitHub
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
[BUG] Claude Code silently deletes conversation transcripts after 30 days by default #62476
Copy link New issue Copy link Open Open [BUG] Claude Code silently deletes conversation transcripts after 30 days by default #62476 Copy link Labels bug Something isn't working Something isn't working Description
Issue body actions Preflight Checklist
I have searched existing issues and this hasn't been reported yet
This is a single bug report (please file separate reports for different bugs)
I am using the latest version of Claude Code
cleanupPeriodDays defaults to 30, causing Claude Code to silently delete
~/.claude/projects//.jsonl files older than 30 days on
startup. There is no first-run disclosure, no warning before deletion, and the
setting is not surfaced in /config. I lost months of conversation history
before realizing this was happening.
Repro: Use Claude Code for >30 days without setting cleanupPeriodDays. Old
transcripts disappear.
Evidence on my machine: ~/.claude/history.jsonl shows 14 sessions / 1,315
prompts for one project from March–April. Only the current session's .jsonl
survives in ~/.claude/projects/. The cutoff matches the 30-day default
exactly. My settings.json does not override the default.
Impact: The code and git history remain, but the reasoning trail — design
discussions, debugging context, analysis — is gone. For research work that
context is the artifact.
Change the default to non-destructive (disabled, or very long retention).
Disclose at first run; require opt-in for auto-deletion.
Soft-delete to a trash folder instead of unlink().
Surface the setting in /config.
Workaround for others: add "cleanupPeriodDays": 3650 to
~/.claude/settings.json immediately.
Use Claude Code for >30 days without setting cleanupPeriodDays. Old transcripts disappear.
Reactions are currently unavailable Metadata
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
