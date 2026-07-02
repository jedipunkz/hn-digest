---
source: "https://github.com/anthropics/claude-code/issues/73125"
hn_url: "https://news.ycombinator.com/item?id=48765630"
title: "Claude's AskUserQuestion: \"No response after 60s – continued without an answer\""
article_title: "[BUG] **EXTREME DANGER**: AskUserQuestion: \"No response after 60s — continued without an answer\" · Issue #73125 · anthropics/claude-code · GitHub"
author: "tubignaaso"
captured_at: "2026-07-02T19:10:18Z"
capture_tool: "hn-digest"
hn_id: 48765630
score: 2
comments: 0
posted_at: "2026-07-02T18:37:57Z"
tags:
  - hacker-news
  - translated
---

# Claude's AskUserQuestion: "No response after 60s – continued without an answer"

- HN: [48765630](https://news.ycombinator.com/item?id=48765630)
- Source: [github.com](https://github.com/anthropics/claude-code/issues/73125)
- Score: 2
- Comments: 0
- Posted: 2026-07-02T18:37:57Z

## Translation

タイトル: クロードのAskUserQuestion: 「60秒を過ぎても応答なし – 応答なしで継続」
記事のタイトル: [バグ] **非常に危険**: AskUserQuestion: 「60 秒を過ぎると応答なし — 応答なしで続行」 · Issue #73125 · anthropics/claude-code · GitHub
説明: プリフライト チェックリスト 既存の問題を検索しましたが、まだ報告されていません これは 1 つのバグ レポートです (バグごとに個別のレポートを提出してください) 最新バージョンの Claude Code を使用しています 何が問題ですか? ●クロードは尋ねた…

記事本文:
[バグ] **非常に危険**: AskUserQuestion: 「60 秒を過ぎても応答なし — 応答なしで継続」 · 問題 #73125 · anthropics/claude-code · GitHub
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
クロード・コッド

e
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
[バグ] **非常に危険**: AskUserQuestion: 「60 秒を過ぎても応答なし — 応答なしで続行」 #73125
リンクをコピー 新しい問題 リンクをコピー 開く 開く [バグ] **非常に危険**: AskUserQuestion: "60 秒後に応答なし — 応答なしで続行" #73125 リンクをコピー ラベル api:bedrock 問題は AWS Bedrock に関連します 問題は AWS Bedrock に関連します エリア:ツールエリア:tui バグ 何かが動作していません 何かが動作していません プラットフォーム:linux 問題は特に Linux で発生します 問題は特に Linux プラットフォームで発生します:vscode 問題は特に VS Code で発生します問題は特に VS コードで発生します。
発行本体アクションのプリフライト チェックリスト
既存の問題を検索しましたが、まだ報告されていません
これは 1 つのバグ レポートです (バグごとに別々のレポートを提出してください)
最新バージョンのクロードコードを使用しています
●クロードはこう尋ねた。
⎿ ·[...]
● 60 秒を超えると応答なし — 応答なしで継続
[...]
❯ これは何ですか 60 秒以降応答なし — 応答なしで継続??????????
● 直接の答え: この行はあなたからではなく、AskUserQuestion ツール自体からのものです。私が 3 つの部分からなる質問 ([...]) を投げかけたとき、ツールの呼び出しは 60 秒間応答されず、ハーネスは次の内容をそのまま返しました。
▎ 「60 秒経過しても応答がありません。ユーザーがキーボードから離れている可能性があります。これまでの状況に基づいて最善の判断を行ってください。次のことが可能です。
▎ まだ関連性がある場合は、後でこの質問を再質問してください。」
そのメカニズムは次のとおりです。AskUserQuestion には 60 秒のタイムアウトがあります。選択が到着しない場合は、「最善の判断で続行する」というメッセージが自動的に返されます。
ブロックする代わりに結果を返します。私は何も捏造したり省略したりしていません。それがツールから返されたものです。
[...]
❯ これはツールパラメータですか？

● いいえ、そうではありません。与えられた AskUserQuestion スキーマを確認しました。そのパラメーターは質問、回答、注釈、メタデータです。
タイムアウトパラメータや「60秒」、応答タイミングに関するものはありません。そのようなパラメーターは渡しませんでした。渡すものは存在しません。
何が起こるべきでしょうか?
これじゃないよ！！！ AskUserQuestion は明示的な安全性チェックであり、適切な質問が行われ、私の入力なしに恐ろしいことが起こらないようにするためのルールとフックがあります。クロードは突然「最善の判断をしてください」と告げられ？何？！
はい、これは以前のバージョンでは機能していました
https://code.claude.com/docs/en/changelog にはこれについて警告するものさえありません。
リアクションは現在利用できません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Preflight Checklist I have searched existing issues and this hasn't been reported yet This is a single bug report (please file separate reports for different bugs) I am using the latest version of Claude Code What's Wrong? ● Claude asked...

[BUG] **EXTREME DANGER**: AskUserQuestion: "No response after 60s — continued without an answer" · Issue #73125 · anthropics/claude-code · GitHub
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
[BUG] **EXTREME DANGER**: AskUserQuestion: "No response after 60s — continued without an answer" #73125
Copy link New issue Copy link Open Open [BUG] **EXTREME DANGER**: AskUserQuestion: "No response after 60s — continued without an answer" #73125 Copy link Labels api:bedrock Issue relates to AWS Bedrock Issue relates to AWS Bedrock area:tools area:tui bug Something isn't working Something isn't working platform:linux Issue specifically occurs on Linux Issue specifically occurs on Linux platform:vscode Issue specifically occurs in VS Code Issue specifically occurs in VS Code Description
Issue body actions Preflight Checklist
I have searched existing issues and this hasn't been reported yet
This is a single bug report (please file separate reports for different bugs)
I am using the latest version of Claude Code
● Claude asked:
⎿ ·[...]
● No response after 60s — continued without an answer
[...]
❯ What is this No response after 60s — continued without an answer??????????
● Direct answer: that line came from the AskUserQuestion tool itself, not from you. When I posed the three-part question ([...]), the tool call sat unanswered for 60 seconds and the harness returned this verbatim:
▎ "No response after 60s — the user may be away from keyboard. Proceed using your best judgment based on the context so far; you can
▎ re-ask this question later if it's still relevant."
So the mechanism is: AskUserQuestion has a 60-second timeout; if no selection arrives, it auto-returns that "proceed with best judgment"
result instead of blocking. I did not fabricate or skip anything — that's what the tool handed back.
[...]
❯ Is this a tool parameter?
● No — it is not. I've just checked the AskUserQuestion schema I was given: its parameters are questions, answers, annotations, metadata.
There is no timeout parameter, no "60s", nothing about response timing. I passed no such parameter, and none exists to pass.
What Should Happen?
Not this!!! AskUserQuestion is an explicit safety check and I have rules and hooks to make sure the right questions are asked and that scary things do not happen without my input. Suddenly Claude is told to "proceed with best judgment"? What?!
Yes, this worked in a previous version
Nothing in https://code.claude.com/docs/en/changelog to even alert about this!!!
Reactions are currently unavailable Metadata
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
