---
source: "https://github.com/lu-wo/apple-contacts-mcp"
hn_url: "https://news.ycombinator.com/item?id=48420790"
title: "Show HN: Apple Contacts MCP – Local AI Access to macOS Contacts"
article_title: "GitHub - lu-wo/apple-contacts-mcp: Local-first MCP server for safely searching and editing Apple Contacts on macOS · GitHub"
author: "luwo"
captured_at: "2026-06-06T04:27:07Z"
capture_tool: "hn-digest"
hn_id: 48420790
score: 1
comments: 0
posted_at: "2026-06-06T02:21:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Apple Contacts MCP – Local AI Access to macOS Contacts

- HN: [48420790](https://news.ycombinator.com/item?id=48420790)
- Source: [github.com](https://github.com/lu-wo/apple-contacts-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T02:21:33Z

## Translation

タイトル: HN を表示: Apple Contacts MCP – macOS 連絡先へのローカル AI アクセス
記事のタイトル: GitHub - lu-wo/apple-contacts-mcp: macOS で Apple の連絡先を安全に検索および編集するためのローカルファースト MCP サーバー · GitHub
説明: macOS 上で Apple の連絡先を安全に検索および編集するためのローカルファースト MCP サーバー - lu-wo/apple-contacts-mcp

記事本文:
GitHub - lu-wo/apple-contacts-mcp: macOS 上で Apple 連絡先を安全に検索および編集するためのローカルファースト MCP サーバー · GitHub
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
ルーウォ
/
アップルの連絡先-mcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビ

イゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .codex-plugin .codex-plugin bin bin docs docs scripts scripts skill/ apple-contacts skill/ apple-contacts testing testing .gitignore .gitignore .mcp.json .mcp.json LICENSE LICENSE README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
macOS 上で Apple Contacts メモを安全に検索、編集、維持するためのローカルファースト MCP サーバー。
現在、サーバーは AppleScript による Contacts.app オートメーションを使用しています。これにより、インストールが簡単になり、クラウド認証情報の代わりに macOS プライバシー プロンプトが使用されます。書き込みはデフォルトで予行演習され、明示的な確認が必要です。
contacts_status : Contacts.app へのアクセスを確認し、集計数を返します。
search_contacts : 名前、組織、役職、電子メール、または電話でローカルの連絡先を検索します。
create_contact : 連絡先を作成します。デフォルトでは予行演習が行われます。
update_contact : スカラー フィールド、メモを更新するか、電子メール/電話番号の値を追加します。デフォルトでは予行演習が行われます。
append_contact_note : 日付の付いたインタラクション ログ エントリを連絡先メモに追加します。デフォルトでは予行演習が行われます。
delete_contact : 連絡先を削除します。デフォルトでは予行演習が行われ、確認フレーズが必要です。
test_roundtrip : 1 つのダミー連絡先を作成、編集、検証、削除します。
macOS のプロンプトが表示されたときの連絡先/オートメーションのアクセス許可
git clone < リポジトリ URL >
cd apple-contacts-mcp
npmテスト
npm 煙を実行する
npm run steam は、連絡先に触れずに MCP ハンドシェイクとツール リストを検証します。
ライブダミー作成/編集/メモ/削除のラウンドトリップを実行するには:
npm 実行スモーク:ライブ
macOS では、連絡先またはオートメーションのアクセス許可を求めるプロンプトが表示される場合があります。生煙テストでは、ダミーの連絡先を 1 つ作成し、それを編集し、メモを追加し、変更を確認して、ダミーを削除します。
これは標準入出力 MCP サーバーです。 MCP 対応エージェントには同じコマンドが必要です。
ノード /

絶対/パス/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs
絶対パスを使用してください。相対パスは、エージェントが別の作業ディレクトリから MCP サーバーを起動するときに簡単に壊れます。
codex mcp add apple-contacts -- ノード /absolute/path/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs
次に、Codex を再起動するか、新しい Codex スレッドを開始して、MCP ツールがロードされるようにします。
サーバーが登録されていることを確認するには:
apple-contacts を使用して contacts_status を実行します。
クロードデスクトップにインストールする
macOS で Claude デスクトップ構成ファイルを開きます。
「 $HOME /Library/Application Support/Claude/claude_desktop_config.json 」を開きます
これを最上位の mcpServers オブジェクトの下に追加し、Claude Desktop を再起動します。
{
"mcpサーバー": {
"リンゴの連絡先" : {
"コマンド" : "ノード" ,
"args" : [ " /absolute/path/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs " ]
}
}
}
ファイルにすでに他の MCP サーバーが含まれている場合は、既存の mcpServers オブジェクト内に apple-contacts エントリのみを追加します。
Apple Contacts MCP サーバーを使用して contacts_status を実行します。接触値は表示しません。
他のエージェントにインストールする
{
"mcpサーバー": {
"リンゴの連絡先" : {
"コマンド" : "ノード" ,
"args" : [ " /absolute/path/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs " ]
}
}
}
エージェントに MCP CLI がある場合は、次のものと同等のものを使用します。
< エージェント > mcp add apple-contacts -- ノード /absolute/path/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs
どのエージェントにとっても適切な最初のプロンプト:
このリポジトリを apple-contacts という名前のローカル MCP サーバーとしてインストールします。 bin/apple-contacts-mcp.cjs への絶対パスを使用し、MCP ツールを再起動またはリロードして contacts_status を実行します。連絡先データを個人データとして扱い、私が明示的に承認しない限り、書き込みは予行演習を続けます。
権限
このサーバーは Contacts.app をローカルで自動化します。最初のライブ通話では、連絡先や連絡先に対する macOS の許可プロンプトがトリガーされる場合があります。

オートメーション。 MCP サーバーを起動している端末またはアプリに対するこれらのプロンプトを承認します。
Contacts.app が実行されていないために通話が失敗した場合は、連絡先を開いて再試行してください。
作成、更新、削除の操作はデフォルトで予行演習されます。実際の書き込みでは、次の両方を通過する必要があります。
{
"dryRun" : false 、
「確認」 : true
}
削除には以下も必要です。
{
"confirmPhrase" : " 連絡先を削除 "
}
これにより、エージェントは自然な 2 段階のフローになります。最初に変更を提案し、次にユーザーの承認後にのみ適用します。
CRM スタイルのインタラクション ログには、メモ フィールド全体を上書きする代わりに、append_contact_note を使用します。
{
"contactId" : "search_contacts からの連絡先 ID " ,
"日付" : " 2026-06-04 " ,
"summary" : " AI 創設者とのディナーで会いました。彼らはローカルファーストのエージェント ツールに興味を持っています。" ,
"openThreads" : [
" GitHub リポジトリを送信する" 、
「来週のデモについてフォローアップします」
]、
"ドライラン" : true
}
追加されたメモのエントリでは次のものが使用されます。
- 2026-06-04 - AI 創設者のディナーで会いました。彼らはローカルファーストのエージェントツールに興味を持っています。
スレッドを開く: GitHub リポジトリを送信します。来週のデモについてフォローアップします
実際の追加には次のものが必要です。
{
"dryRun" : false 、
「確認」 : true
}
プライバシー
このサーバーは Mac 上でローカルに実行されます。クラウド API を呼び出したり、連絡先を独自にアップロードしたりすることはありません。エージェントは、MCP サーバーに返すように要求した連絡先データを引き続き参照するため、可能な場合はフィールド フィルターと編集を使用してください。
連絡先は、iCloud、Google、Exchange、または別の構成済みアカウントを通じて同期できます。ローカル書き込みはそれらのサービスに伝播する可能性があります。
最初のバックエンドは、 Contacts.app の AppleScript 自動化です。将来のバックエンドでは、より構造化されたアクセスを実現するために、Apple の Contacts.framework の周囲に署名された Swift ヘルパーが使用される可能性があります。
~/Library/Application Support/AddressBook への SQLite の直接書き込みは意図的にサポートされていません。
安全な検索と電子メールを実現するローカルファースト MCP サーバー

macOS で Apple の連絡先を編集する
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

Local-first MCP server for safely searching and editing Apple Contacts on macOS - lu-wo/apple-contacts-mcp

GitHub - lu-wo/apple-contacts-mcp: Local-first MCP server for safely searching and editing Apple Contacts on macOS · GitHub
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
lu-wo
/
apple-contacts-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .codex-plugin .codex-plugin bin bin docs docs scripts scripts skills/ apple-contacts skills/ apple-contacts tests tests .gitignore .gitignore .mcp.json .mcp.json LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
Local-first MCP server for safely searching, editing, and maintaining Apple Contacts notes on macOS.
The server uses Contacts.app automation through AppleScript today. That keeps install simple and uses macOS privacy prompts instead of cloud credentials. Writes are dry-run by default and require explicit confirmation.
contacts_status : check Contacts.app access and return aggregate counts.
search_contacts : search local contacts by name, organization, job title, email, or phone.
create_contact : create a contact. Dry-run by default.
update_contact : update scalar fields, notes, or add email/phone values. Dry-run by default.
append_contact_note : append a dated interaction log entry to a contact note. Dry-run by default.
delete_contact : delete a contact. Dry-run by default and requires a confirmation phrase.
test_roundtrip : create, edit, verify, and delete one dummy contact.
Contacts/Automation permissions when macOS prompts
git clone < repo-url >
cd apple-contacts-mcp
npm test
npm run smoke
npm run smoke verifies the MCP handshake and tool list without touching Contacts.
To run a live dummy create/edit/note/delete roundtrip:
npm run smoke:live
macOS may prompt for Contacts or Automation permissions. The live smoke test creates one dummy contact, edits it, appends a note, verifies the change, and deletes the dummy.
This is a stdio MCP server. Any MCP-capable agent needs the same command:
node /absolute/path/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs
Use an absolute path. Relative paths are easy to break when an agent launches MCP servers from another working directory.
codex mcp add apple-contacts -- node /absolute/path/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs
Then restart Codex or start a new Codex thread so the MCP tools are loaded.
To confirm the server is registered:
Use apple-contacts to run contacts_status.
Install In Claude Desktop
Open the Claude Desktop config file on macOS:
open " $HOME /Library/Application Support/Claude/claude_desktop_config.json "
Add this under the top-level mcpServers object, then restart Claude Desktop:
{
"mcpServers" : {
"apple-contacts" : {
"command" : " node " ,
"args" : [ " /absolute/path/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs " ]
}
}
}
If the file already has other MCP servers, add only the apple-contacts entry inside the existing mcpServers object.
Use the Apple Contacts MCP server to run contacts_status. Do not show any contact values.
Install In Other Agents
{
"mcpServers" : {
"apple-contacts" : {
"command" : " node " ,
"args" : [ " /absolute/path/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs " ]
}
}
}
If your agent has an MCP CLI, use its equivalent of:
< agent > mcp add apple-contacts -- node /absolute/path/to/apple-contacts-mcp/bin/apple-contacts-mcp.cjs
Good first prompt for any agent:
Install this repository as a local MCP server named apple-contacts. Use the absolute path to bin/apple-contacts-mcp.cjs, then restart or reload your MCP tools and run contacts_status. Treat contact data as personal data and keep writes dry-run unless I explicitly approve them.
Permissions
This server automates Contacts.app locally. The first live call may trigger macOS permission prompts for Contacts and/or Automation. Approve those prompts for the terminal or app that is launching the MCP server.
If a call fails because Contacts.app is not running, open Contacts and retry:
Create, update, and delete operations are dry-run by default. An actual write must pass both:
{
"dryRun" : false ,
"confirm" : true
}
Delete also requires:
{
"confirmPhrase" : " delete contact "
}
This gives agents a natural two-step flow: propose the change first, then apply only after user approval.
Use append_contact_note for CRM-style interaction logs instead of overwriting the full note field.
{
"contactId" : " CONTACT-ID-FROM-search_contacts " ,
"date" : " 2026-06-04 " ,
"summary" : " Met at an AI founder dinner. They are interested in local-first agent tooling. " ,
"openThreads" : [
" Send the GitHub repo " ,
" Follow up about a demo next week "
],
"dryRun" : true
}
The appended note entry uses:
- 2026-06-04 - Met at an AI founder dinner. They are interested in local-first agent tooling.
Open threads: Send the GitHub repo; Follow up about a demo next week
An actual append requires:
{
"dryRun" : false ,
"confirm" : true
}
Privacy
This server runs locally on your Mac. It does not call a cloud API or upload contacts on its own. Your agent will still see whatever contact data you ask the MCP server to return, so use field filters and redaction when possible.
Contacts may sync through iCloud, Google, Exchange, or another configured account. A local write can propagate to those services.
The first backend is AppleScript automation of Contacts.app . A future backend may use a signed Swift helper around Apple's Contacts.framework for more structured access.
Direct SQLite writes to ~/Library/Application Support/AddressBook are intentionally not supported.
Local-first MCP server for safely searching and editing Apple Contacts on macOS
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
