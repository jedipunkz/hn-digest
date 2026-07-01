---
source: "https://github.com/torontodeveloper/job-application-agent"
hn_url: "https://news.ycombinator.com/item?id=48752969"
title: "Show HN:An AI agent that applies to jobs for me (Playwright,GPT5.4 form filling)"
article_title: "GitHub - torontodeveloper/job-application-agent: Job application agent, ai applies jobs and I review and Submit using Python, gpt 5.4, PLAYWRIGHT · GitHub"
author: "torontodev007"
captured_at: "2026-07-01T21:31:30Z"
capture_tool: "hn-digest"
hn_id: 48752969
score: 2
comments: 3
posted_at: "2026-07-01T20:53:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN:An AI agent that applies to jobs for me (Playwright,GPT5.4 form filling)

- HN: [48752969](https://news.ycombinator.com/item?id=48752969)
- Source: [github.com](https://github.com/torontodeveloper/job-application-agent)
- Score: 2
- Comments: 3
- Posted: 2026-07-01T20:53:47Z

## Translation

タイトル: Show HN:An AI Agent that apply to jobs for me (Playwright、GPT5.4 フォーム入力)
記事タイトル: GitHub - torontodeveloper/job-application-agent: 求人応募エージェント、AI が求人を応募し、Python、gpt 5.4、PLAYWRIGHT を使用してレビューして送信する · GitHub
説明: 求人応募エージェント、AI が求人を適用し、Python、gpt 5.4、PLAYWRIGHT を使用してレビューして送信します - torontodeveloper/job-application-agent

記事本文:
GitHub - torontodeveloper/job-application-agent: 求人応募エージェント、AI が求人を適用し、Python、gpt 5.4、PLAYWRIGHT を使用してレビューして送信する · GitHub
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
トロント開発者
/
求人応募エージェント
公共
通知
サインインする必要があります

o 通知設定を変更する
追加のナビゲーション オプション
コード
トロント開発者/求人応募エージェント
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .gitignore .gitignore README.md README.md apply.sh apply.sh apply.py Discover.py find_jobs.sh find_jobs.sh followup.py followup.py gen_pdf.py gen_pdf.py gmail_job_audit.py gmail_job_audit.py gmail_web_audit.py gmail_web_audit.pylogin.pylogin.pymain.pymain.pymake_cv.pymake_cv.pyoutreach.pyoutreach.pyoutreach.shoutreach.shrequirements.txtrequirements.txt simplify_capture.py simplify_debug.py simplify_debug.py tailor.py tailor.py yc_discover.py yc_discover.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
求人と応募のループ全体を実行する AI エージェント: 空いている役割を発見し、
職務内容ごとに履歴書を調整し、PDF を生成し、必要事項を記入します。
ATS アプリケーション フォーム (Greenhouse、Ashby、Lever など) — ドロップダウン、コンボボックス、
そして「なぜここで働きたいのか」のボックスも含まれています。それからそれは止まり、
確認して「送信」をクリックするのを待ちます。単独で提出することはありません。
▶️ 完全なデモを見る - カメラ上の実際のジョブに適用されます
求人掲示板全体で空き職種を発見します ( Discover.py 、 yc_discover.py )
履歴書を各職務内容に合わせて調整し、新しい PDF を生成します
( tailor.py 、 gen_pdf.py 、 make_cv.py )
実際のアプリケーションを独自のブラウザ ウィンドウに表示します。つまり、メイン
拡張機能ベースの自動入力 ( main.py ) とは異なり、ブラウザは無料のままです
送信するたびに人間によるレビューのために一時停止します。これはアシスタントではなく、アシスタントです。
スパム砲
クラッシュしても生き残る : キューの処理は中断したところから再開されます
フォローアップ: 送信されたアプリケーションに対するフォローアップ メールの下書き ( followup.py )
profile.yaml には事実と音声メモが含まれます (唯一の信頼できる情報源)
遊ぶ

wright は、永続的なジョブ URL を専用の Chromium で開きます。
プロファイル (ログイン/Cookie は実行間で存続します)
フォームは DOM から抽出されます。 LLM はすべてのフィールドを答えにマッピングします。
プロフィールの事実をそのまま、あなたの声で書かれた自由回答の質問、
不明なものはスキップされました
エージェントはフォームに入力し、その後一時停止します。確認して [送信] をクリックします。
あなた自身
python -m venv .venv && ソース .venv/bin/activate
pip install -r 要件.txt
Playwright が Chrome をインストールする
# 詳細情報 (名前、連絡先、職歴、音声メモ) を含む profile.yaml を作成します
export LLM_API_KEY=... # およびオプションで LLM_BASE_URL / LLM_MODEL
走る
python main.py https://boards.greenhouse.io/acme/jobs/12345
python main.py --queue jobs.txt # 1 行に 1 つの URL
./apply.sh # ランチャー スクリプトによる完全なパイプライン
注意事項
量よりも質: 送信する前にすべての申請書を確認してください
これを LinkedIn Easy apply (ToS) または CAPTCHA で囲まれたフローに向けないでください。
Workday: エージェントのブラウザ ウィンドウに 1 回ログインします。永続的なプロファイル
覚えています
ブラウザのプロフィール、履歴書、アプリケーションデータは無視されます - そのままにしておきます
そのように;セッション Cookie と個人情報が含まれています
求人応募エージェント、AI が求人を応募し、Python、gpt 5.4、PLAYWRIGHT を使用してレビューして送信します
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

Job application agent, ai applies jobs and I review and Submit using Python, gpt 5.4, PLAYWRIGHT - torontodeveloper/job-application-agent

GitHub - torontodeveloper/job-application-agent: Job application agent, ai applies jobs and I review and Submit using Python, gpt 5.4, PLAYWRIGHT · GitHub
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
torontodeveloper
/
job-application-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
torontodeveloper/job-application-agent
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .gitignore .gitignore README.md README.md apply.sh apply.sh discover.py discover.py find_jobs.sh find_jobs.sh followup.py followup.py gen_pdf.py gen_pdf.py gmail_job_audit.py gmail_job_audit.py gmail_web_audit.py gmail_web_audit.py login.py login.py main.py main.py make_cv.py make_cv.py outreach.py outreach.py outreach.sh outreach.sh requirements.txt requirements.txt simplify_capture.py simplify_capture.py simplify_debug.py simplify_debug.py tailor.py tailor.py yc_discover.py yc_discover.py View all files Repository files navigation
An AI agent that does the whole job-application loop: discovers open roles,
tailors your resume per job description, generates a PDF, and fills out the
ATS application form (Greenhouse, Ashby, Lever, ...) — dropdowns, comboboxes,
and the "why do you want to work here" boxes included. Then it stops and
waits for you to review and click Submit. It never submits on its own.
▶️ Watch the full demo — it applies to real jobs on camera
Discovers open roles across job boards ( discover.py , yc_discover.py )
Tailors your resume to each job description and generates a fresh PDF
( tailor.py , gen_pdf.py , make_cv.py )
Fills the actual application in its own browser window — so your main
browser stays free, unlike extension-based autofill ( main.py )
Pauses for human review before every submit — it's an assistant, not a
spam cannon
Survives crashes : queue processing picks up where it left off
Follows up : drafts follow-up emails for submitted applications ( followup.py )
profile.yaml holds your facts + voice notes (single source of truth)
Playwright opens the job URL in a dedicated Chromium with a persistent
profile (logins/cookies survive between runs)
The form is extracted from the DOM; an LLM maps every field to an answer —
profile facts verbatim, open-ended questions drafted in your voice,
unknowns skipped
The agent fills the form, then pauses — you review and click Submit
yourself
python -m venv .venv && source .venv/bin/activate
pip install -r requirements.txt
playwright install chromium
# create profile.yaml with your details (name, contact, work history, voice notes)
export LLM_API_KEY=... # and optionally LLM_BASE_URL / LLM_MODEL
Run
python main.py https://boards.greenhouse.io/acme/jobs/12345
python main.py --queue jobs.txt # one URL per line
./apply.sh # full pipeline via launcher script
Notes
Quality over volume: review every application before submit
Don't point this at LinkedIn Easy Apply (ToS) or CAPTCHA-walled flows
Workday: log in once in the agent's browser window; the persistent profile
remembers
Browser profiles, resumes, and application data are gitignored — keep it
that way; they contain session cookies and personal info
Job application agent, ai applies jobs and I review and Submit using Python, gpt 5.4, PLAYWRIGHT
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
