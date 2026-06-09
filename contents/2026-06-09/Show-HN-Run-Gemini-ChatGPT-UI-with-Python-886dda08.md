---
source: "https://github.com/pseudo-usama/hermex"
hn_url: "https://news.ycombinator.com/item?id=48461184"
title: "Show HN: Run Gemini & ChatGPT UI with Python"
article_title: "GitHub - pseudo-usama/hermex: Drive ChatGPT and Gemini from Python — no API keys, no billing, just the free web UI. · GitHub"
author: "pseudo-usama"
captured_at: "2026-06-09T16:06:13Z"
capture_tool: "hn-digest"
hn_id: 48461184
score: 3
comments: 1
posted_at: "2026-06-09T13:52:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Run Gemini & ChatGPT UI with Python

- HN: [48461184](https://news.ycombinator.com/item?id=48461184)
- Source: [github.com](https://github.com/pseudo-usama/hermex)
- Score: 3
- Comments: 1
- Posted: 2026-06-09T13:52:48Z

## Translation

タイトル: HN を表示: Python を使用して Gemini と ChatGPT UI を実行する
記事のタイトル: GitHub - pseudo-usama/hermex: Python から ChatGPT と Gemini を駆動する — API キーなし、課金なし、無料の Web UI のみ。 · GitHub
説明: Python から ChatGPT と Gemini を駆動します。API キーや課金はなく、無料の Web UI だけを使用します。 - 擬似うさま/ヘルメックス

記事本文:
GitHub - pseudo-usama/hermex: Python から ChatGPT と Gemini を駆動します。API キーや課金はなく、無料の Web UI のみです。 · GitHub
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
似非うさま
/
ヘルメックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
149 コミット 149 コミット .github/ workflows .github/ workflows アセット アセット docs docs hermex hermex .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md pyproject.toml pyproject.toml 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Python から ChatGPT と Gemini を駆動します。API キーや課金はなく、無料の Web UI だけを使用します。
ChatGPT と Gemini は信じられないほど有能ですが、公式 API は高価であり、多くのタスクではそれらを必要としません。画像に対して OCR を実行したり、アートワークを生成したり、スクリーンショットからテキストを抽出したり、スクリプトで簡単な質問をしたりしたい場合、無料の Web UI で同じことができるのに、API アクセスに対してトークンごとに料金を支払うのは過剰です。
Hermex では、ChatGPT と Gemini を Python で自動化できます。API キー、課金、ペイウォールは必要ありません。人間と同じように、本物の Chrome ブラウザを開き、メッセージを入力し、ファイルをアップロードし、応答を待って、それを Python オブジェクトとして返します。
hermexインポートChatGPTから
応答 = ChatGPT 。 simple_query ( "この領収書には何と書いてありますか?" ,attachments = [ "receipt.jpg" ])
print (応答のテキスト)
内部で undetected-chromedriver を使用してボットの検出を回避し、永続的なブラウザー プロファイルを再利用するので、ログイン セッションは実行後も存続します。
pip インストール hermex
Python 3.11以降およびGoogle Chrome 130以降が必要です
Hermex は永続的な Chrome プロファイルを再利用するため、ログインする必要があるのは 1 回だけです。
hermex輸入ジェミニから
ジェミニ。 setup () # ブラウザを開きます — ログインし、少し閲覧してからウィンドウを閉じます
セットアップ後は、今後のすべての実行で、保存されたセッションが自動的に再利用されます。セッションの有効期限が切れた場合は、これを繰り返します。

ゲスト モード (ログインなし) は、Gemini での基本的なテキスト クエリには機能しますが、ファイルのアップロードにはログイン セッションが必要です。 ChatGPT はテキスト クエリやファイルのアップロードにはログインなしで機能しますが、画像の生成にはログインしたセッションが必要です。
hermex import Gemini 、ChatGPT から
#ジェミニ
ジェミニ = ジェミニ ()
ジェミニ。 open_url ()
応答 = ジェミニ 。クエリ (「インターネットの歴史を要約してください。」)
print (応答のテキスト)
ジェミニ。閉じる（）
#ChatGPT
chatgpt = ChatGPT ()
チャットポイント 。 open_url ()
応答 = chatgpt 。クエリ (「インターネットの歴史を要約してください。」)
print (応答のテキスト)
チャットポイント 。閉じる（）
ファイルを添付する
応答 = ジェミニ 。クエリ (
「この画像に何が写っているのか説明してください。」 、
添付ファイル = [ "photo.jpg" ],
)
print (応答のテキスト)
サポートされている形式: .jpg 、 .jpeg 、 .png 、 .gif 、 .webp 、 .pdf 、 .csv 、 .txt 、 .json 。各プラットフォームは、 Gemini.SUPPORTED_ATTACHMENTS および ChatGPT.SUPPORTED_ATTACHMENTS を介して、許可されているタイプを公開します。
hermex import Gemini 、ChatGPT から
応答 = ジェミニ。 simple_query ( "フランスの首都はどこですか?" )
print (応答のテキスト)
応答 = ChatGPT 。 simple_query ( "フランスの首都はどこですか?" )
print (応答のテキスト)
#添付ファイルあり
応答 = ジェミニ。 simple_query ( "この画像について説明します。" 、attachments = [ "photo.jpg" ])
print (応答のテキスト)
AssistantMessage オブジェクト
query() と get_last_response() は AssistantMessage データクラスを返します。
@データクラス
クラス AssistantMessage :
テキスト: str |なし # プレーンテキスト (または get_markdown=True の場合はマークダウン)
画像 : パス |なし # ダウンロードされたイメージへのパス、またはなし
範囲と制限事項
制作用ではなく、趣味、スクリプト作成、研究用に構築されています。実際のブラウザ UI を駆動するため、公式 API よりも遅く、堅牢性も劣ります。
チャット UI が変更されると壊れる可能性があります。 Hermex は ChatGPT と Gem のページ構造に依存しています

イニ;彼らの側で再設計が行われると、ここでの更新が必要になる場合があります。
各プラットフォームの利用規約を尊重してください。 Web UI の自動化は、あらゆるユースケースに適しているわけではありません。責任を持って、自己責任で使用してください。
プラットフォーム独自の使用制限は引き続き適用されます。無料の Web UI で適用される制限 (メッセージと画像の生成制限、クールダウン) は引き続き有効です。
Gemini と ChatGPT はどちらも同じインターフェイスを共有します。特に記載がない限り、以下のすべてのメソッドが両方に適用されます。
Gemini と ChatGPT の詳細なガイドについては、完全なドキュメントを参照してください。
ジェミニ (
chrome_version = None 、 # インストールされている Chrome から自動検出
download_dir = パス ( "." )、 # 生成された画像が保存される場所
headless = False 、
タイピング_遅延 = 0.025 、キーストローク間の # 秒
disable_web_security = True 、
)
# ChatGPT は同じパラメータを受け入れます
ウォーターマークの除去
Gemini は、生成された画像に透かしを入れます。削除するには、remove_watermark=True を渡します。
応答 = ジェミニ 。 query ( "夕日の画像を生成します。" 、remove_watermark = True )
注意事項
Bot detection is mitigated through per-character typing delays, fake typing before paste, a persistent browser profile, and a spoofed user agent.機密性の高いセッションではヘッドレスでの実行を避けてください。
ブラウザーのプロファイルとセッション データは、プラットフォーム データ ディレクトリ (macOS では ~/Library/Application Support/hermex) に保存されます。
完全なリリース履歴については、CHANGELOG.md を参照してください。
Python から ChatGPT と Gemini を駆動します。API キーや課金はなく、無料の Web UI だけを使用します。
読み込み中にエラーが発生しました。このページをリロードしてください。
7
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Drive ChatGPT and Gemini from Python — no API keys, no billing, just the free web UI. - pseudo-usama/hermex

GitHub - pseudo-usama/hermex: Drive ChatGPT and Gemini from Python — no API keys, no billing, just the free web UI. · GitHub
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
pseudo-usama
/
hermex
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
149 Commits 149 Commits .github/ workflows .github/ workflows assets assets docs docs hermex hermex .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
Drive ChatGPT and Gemini from Python — no API keys, no billing, just the free web UI.
ChatGPT and Gemini are incredibly capable — but their official APIs are expensive, and for many tasks you simply don't need them. If you want to run OCR on an image, generate artwork, extract text from a screenshot, or just ask a quick question in a script, paying per-token for API access is overkill when the free web UI can do the same thing.
Hermex lets you automate ChatGPT and Gemini with Python — no API keys, no billing, no paywalls. It opens a real Chrome browser, types your message, uploads your files, waits for the response, and hands it back to you as a Python object, just like a human would.
from hermex import ChatGPT
response = ChatGPT . simple_query ( "What does this receipt say?" , attachments = [ "receipt.jpg" ])
print ( response . text )
It uses undetected-chromedriver under the hood to avoid bot detection, and reuses a persistent browser profile so your login session survives across runs.
pip install hermex
Requires Python 3.11+ and Google Chrome 130+
Hermex reuses a persistent Chrome profile so you only need to log in once:
from hermex import Gemini
Gemini . setup () # opens a browser — log in, browse briefly, then close the window
After setup, all future runs reuse the saved session automatically. Repeat this if your session expires.
Guest mode (no login) works for basic text queries on Gemini but file upload requires a logged-in session. ChatGPT works without login for text queries and file upload, but image generation requires a logged-in session.
from hermex import Gemini , ChatGPT
# Gemini
gemini = Gemini ()
gemini . open_url ()
response = gemini . query ( "Summarize the history of the internet." )
print ( response . text )
gemini . close ()
# ChatGPT
chatgpt = ChatGPT ()
chatgpt . open_url ()
response = chatgpt . query ( "Summarize the history of the internet." )
print ( response . text )
chatgpt . close ()
Attaching files
response = gemini . query (
"Describe what's in this image." ,
attachments = [ "photo.jpg" ],
)
print ( response . text )
Supported formats: .jpg , .jpeg , .png , .gif , .webp , .pdf , .csv , .txt , .json . Each platform exposes its allowed types via Gemini.SUPPORTED_ATTACHMENTS and ChatGPT.SUPPORTED_ATTACHMENTS .
from hermex import Gemini , ChatGPT
response = Gemini . simple_query ( "What is the capital of France?" )
print ( response . text )
response = ChatGPT . simple_query ( "What is the capital of France?" )
print ( response . text )
# With an attachment
response = Gemini . simple_query ( "Describe this image." , attachments = [ "photo.jpg" ])
print ( response . text )
AssistantMessage object
query() and get_last_response() return an AssistantMessage dataclass:
@ dataclass
class AssistantMessage :
text : str | None # plain text (or markdown if get_markdown=True)
image : Path | None # path to downloaded image, or None
Scope & limitations
Built for hobby, scripting, and research — not production. It drives a real browser UI, so it's slower and less robust than an official API.
It can break when the chat UIs change. Hermex depends on the page structure of ChatGPT and Gemini; a redesign on their end may require an update here.
Respect each platform's Terms of Service. Automating the web UI isn't appropriate for every use case — use it responsibly and at your own risk.
The platforms' own usage limits still apply. Any caps the free web UI enforces (message and image-generation limits, cooldowns) are still in effect.
Both Gemini and ChatGPT share the same interface — all methods below apply to both unless noted.
See the full documentation for detailed guides on Gemini and ChatGPT.
Gemini (
chrome_version = None , # auto-detected from installed Chrome
download_dir = Path ( "." ), # where generated images are saved
headless = False ,
typing_delay = 0.025 , # seconds between keystrokes
disable_web_security = True ,
)
# ChatGPT accepts the same parameters
Watermark removal
Gemini watermarks its generated images. Pass remove_watermark=True to strip it:
response = gemini . query ( "Generate an image of a sunset." , remove_watermark = True )
Notes
Bot detection is mitigated through per-character typing delays, fake typing before paste, a persistent browser profile, and a spoofed user agent. Avoid running headless for sensitive sessions.
Browser profile and session data are stored in the platform data directory ( ~/Library/Application Support/hermex on macOS).
See CHANGELOG.md for the full release history.
Drive ChatGPT and Gemini from Python — no API keys, no billing, just the free web UI.
There was an error while loading. Please reload this page .
7
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
