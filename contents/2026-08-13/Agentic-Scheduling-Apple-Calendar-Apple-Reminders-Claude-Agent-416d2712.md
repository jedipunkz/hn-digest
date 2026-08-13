---
source: "https://github.com/hunterZh37/agentic-scheduling"
hn_url: "https://news.ycombinator.com/item?id=49292383"
title: "Agentic Scheduling = Apple Calendar + Apple Reminders + Claude Agent"
article_title: "GitHub - hunterZh37/agentic-scheduling: Your calendar, run by an AI agent you can text. Self-hostable scheduling: Google/Outlook sync, a booking page with only real openings, WhatsApp control and reminders. · GitHub"
author: "hunterzh37"
captured_at: "2026-08-13T22:31:11Z"
capture_tool: "hn-digest"
hn_id: 49292383
score: 1
comments: 0
posted_at: "2026-08-13T22:00:59Z"
tags:
  - hacker-news
  - translated
---

# Agentic Scheduling = Apple Calendar + Apple Reminders + Claude Agent

- HN: [49292383](https://news.ycombinator.com/item?id=49292383)
- Source: [github.com](https://github.com/hunterZh37/agentic-scheduling)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T22:00:59Z

## Translation

タイトル: Agentic Scheduling = Apple Calendar + Apple Reminders + Claude Agent
記事のタイトル: GitHub - HunterZh37/agentic-scheduling: AI エージェントによって実行されるカレンダーで、テキスト メッセージを送信できます。自己ホスト可能なスケジュール設定: Google/Outlook 同期、実際の空き状況のみを表示する予約ページ、WhatsApp コントロール、リマインダー。 · GitHub
説明: テキストメッセージを送信できる AI エージェントによって実行されるカレンダー。自己ホスト可能なスケジュール設定: Google/Outlook 同期、実際の空き状況のみを表示する予約ページ、WhatsApp コントロール、リマインダー。 - HunterZh37/エージェントスケジューリング

記事本文:
GitHub - HunterZh37/agentic-scheduling: テキストメッセージを送信できる AI エージェントによって実行されるカレンダー。自己ホスト可能なスケジュール設定: Google/Outlook 同期、実際の空き状況のみを表示する予約ページ、WhatsApp コントロール、リマインダー。 · GitHub
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
ハンターZh37
/
エージェントのスケジュール設定
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .githooks .githooks .github .github calendly-extension ca

lenly-extension docs docs e2e e2e prisma prisma public public scripts scripts src src website website .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md docker-compose.yml docker-compose.yml eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json playwright.config.ts playwright.config.ts postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json vercel.json vercel.json vitest.config.mts vitest.config.mts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントによって実行されるカレンダー。メッセージを送信できます。
自己ホスト可能な単一所有者のスケジューリング プラットフォーム。 Google と Outlook を同期します
カレンダー、空いているスロットのみを表示する予約ページを訪問者に提供し、
クロード エージェントが担当し、Web または WhatsApp から制御可能。
デモ: 訪問者がアシスタントに営業時間を尋ねると、Chrome 拡張機能がオーバーレイされます。
他の人の Calendly ページに自分の空き時間情報を表示します。
(フル品質ビデオ)
一般予約ページ。訪問者は長さと日付を選択します。
日を選択すると、オープン時間が開きます。表示されているすべてのスロットは完全に無料です
つながったカレンダー。
同じページ、エージェント モード。訪問者がわかりやすい言葉で質問し、アシスタントが答える
ライブ可用性から。選択したスロットを予約することもできます。
他の人の Calendly ページにある Chrome 拡張機能 (Calendly の公開デモ)。毎
スロットには独自の空き時間情報がバッジとして付けられ、デプロイメントからライブで取得されるため、
自分自身を紛争に巻き込んでください。
電話でも同様の予約フローです。
両方のチャンネルでリマインダー。左: 出席者に送信される確認とリマインダーのメール
受け取ります。右: オーナーが受け取る WhatsApp メッセージ、その瞬間の予約アラート
誰かが会議の前に予約とリマインダーをくれた

g.実際のテンプレートからレンダリングされたもの
サンプルデータ。
任意の数の Google アカウントと Outlook アカウントを接続できます (7 つのカレンダー、1 つの場所)。の
空き状況エンジンは空き時間情報を 1 つのビューに統合します。予約ページでは、
実際の空き状況と確認された予約は、選択したカレンダーに書き戻されます。
%%{init: {"themeVariables": {"fontSize": "12px"}}}%%
フローチャートTB
サブグラフ cal["あなたのカレンダー、任意の数"]
LR方向
G1[「グーグル 1」]
G2[「グーグル2」]
G3[「グーグルN」]
O1[「Outlook 1」]
O2[「Outlook N」]
終わり
cal --> AGG["空き時間情報の集計<br>UTC にマージ"]
AGG --> AV["可用性エンジン<br>空きスロット、予約時に再検証"]
AV --> BP[「公開予約ページ」]
AV --> AG[「クロードエージェント」]
AG --> WA["WhatsApp<br>コントロール、アラート、リマインダー"]
AG --> EM["メール<br>確認、リマインダー"]
BP --> DEST[「目的地カレンダー<br>イベントと招待状の書き戻し」]
読み込み中
特長
プロバイダー間でのカレンダーの同期。任意の数の Google から空き時間情報を集計します
および Microsoft (Outlook) アカウント。それぞれ OAuth によって接続されます。予約は次の宛先に書き込まれます
あなたが指定したカレンダー。
実際の空き情報のみを提供する予約ページ。エージェントが継続的にあなたのことをチェックします
開いているスロットのカレンダーを統合しました。空室状況は予約時に再確認されますので、
取得したばかりのスロットをダブルブッキングすることはできません。
WhatsApp の直接コミュニケーション。イベントを作成、移動、またはキャンセルするためにエージェントにテキストメッセージを送信します。
空き状況を確認したり、リマインダーを設定したりできます。新規のご予約、キャンセル、日程変更について
発生した瞬間に WhatsApp やメールでプッシュ通知されます。
リマインダーメールと WhatsApp リマインダー。出席者とオーナーがミーティングを行う
電子メールによるリマインダーと承認済みの WhatsApp テンプレート。クレーム/デッドレター担当者
失敗した送信をドロップする代わりに再試行します。
予約ページでアシスタントにお尋ねください。訪問者は公的エージェントとチャットできます。

ああ
時間を見つけてください。スコープで囲まれたツールが 2 つだけ渡されます。専用ツールは、
したがって、安全性はプロンプトではなくアーキテクチャによって決まります。
毎日の朝の簡単な説明。その日の会議とタスクの DST セーフな WhatsApp の概要。
セルフサービスでスケジュール変更とキャンセルを行います。すべての確認には署名されたリンクが含まれます。
参加者はアカウントなしで使用できます。
Chromeの拡張機能。自分の実際の空き状況を他の人のカレンダーに重ねて表示します
ページ。
予約、エージェント、リマインダー、カレンダーの追加機能 (個人ブロック、誕生日、
フォローアップ) は src/app/api/ にあるプレーンな JSON API なので、どこからでも再利用できます。
クライアント。
エージェント間のスケジューリング。訪問者のエージェントとあなたのエージェントが会議について交渉します
共有プロトコル上で時間を測定し、結果を予約します。
git clone https://github.com/hunterZh37/agentic-scheduling.git && cd Agentic-scheduling
npmインストール
cp .env.example .env # デフォルトはそのまま使用できます
docker compose up -d # ローカル Postgres
npx prisma 移行開発 && npx prisma db シード
npm run dev # http://localhost:3000
アプリはプレースホルダー構成で完全に起動するため、UI をクリックしてから起動することができます。
現実のものをすべて接続します。ライブの準備ができたら、
docs/SETUP.md は、各認証情報 (Google、Microsoft、
Anthropic、Twilio、Resend) を一度に 1 つずつ、任意の順序で実行します。
すべての ID は環境変数から取得されます: 所有者名、電子メール、タイムゾーン、電話番号
番号、およびカレンダーの資格情報 ( .env.example を参照)。 1つのファイル
ブランドのデフォルトが焼き付けられているのは、
src/lib/booking/publicConfig.ts (実践名、
ドメイン、コンサルティング分野、研究リンク）。独自のインスタンスをデプロイするときに編集します。
すべては環境変数です ( .env.example を参照)。個人的なものはありません
データはハードコードされています。
docs/SETUP.md : 認証情報、各環境変数にマッピング
docs/DEPLOY.md : Vercel へのデプロイ
docs/SMS.md : WhatsApp/SMS コントロール

チャンネル
docs/MORNING_BRIEF.md : WhatsApp の毎日の概要
Vercel 上の Next.js (App Router)、React 19、および TypeScript。 Prisma 経由の Postgres。クロード
エージェントを呼び出すツール。 SMS/WhatsApp 用の Twilio、電子メール用の再送信、Web 用の VAPID
押します。時間と繰り返しのためのルクソンとルール。テスト用のヴィテスト。
問題やPRを歓迎します。送信する前に npm test と npm run lint を実行します。
AI エージェントによって実行されるカレンダー。メッセージを送信できます。自己ホスト可能なスケジュール設定: Google/Outlook 同期、実際の空き状況のみを表示する予約ページ、WhatsApp コントロール、リマインダー。
bookwithhunter.com/book トピックス
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Your calendar, run by an AI agent you can text. Self-hostable scheduling: Google/Outlook sync, a booking page with only real openings, WhatsApp control and reminders. - hunterZh37/agentic-scheduling

GitHub - hunterZh37/agentic-scheduling: Your calendar, run by an AI agent you can text. Self-hostable scheduling: Google/Outlook sync, a booking page with only real openings, WhatsApp control and reminders. · GitHub
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
hunterZh37
/
agentic-scheduling
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .githooks .githooks .github .github calendly-extension calendly-extension docs docs e2e e2e prisma prisma public public scripts scripts src src website website .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json playwright.config.ts playwright.config.ts postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json vercel.json vercel.json vitest.config.mts vitest.config.mts View all files Repository files navigation
Your calendar, run by an AI agent you can text.
A self-hostable, single-owner scheduling platform. It syncs your Google and Outlook
calendars, gives visitors a booking page that shows only your open slots, and puts a
Claude agent in charge, controllable from the web or WhatsApp.
Demo: a visitor asks the assistant for open times, then the Chrome extension overlays
your own free/busy on someone else's Calendly page.
( full-quality video )
The public booking page. Visitors pick a length and a day.
Picking a day opens the open times. Every slot shown is genuinely free across all
connected calendars.
The same page, agent mode. Visitors ask in plain language and the assistant answers
from live availability. It can also book the chosen slot.
The Chrome extension on someone else's Calendly page (Calendly's public demo). Every
slot is badged with your own free/busy, pulled live from your deployment, so you never
book yourself into a conflict.
The same booking flow on a phone.
Reminders on both channels. Left: the confirmation and reminder email an attendee
receives. Right: the WhatsApp messages the owner receives, a booking alert the moment
someone books and a reminder before the meeting. Rendered from the real templates with
sample data.
Connect any number of Google and Outlook accounts (seven calendars, one place). The
availability engine merges their free/busy into one view, the booking page offers only
real openings, and confirmed bookings are written back to the calendar you choose.
%%{init: {"themeVariables": {"fontSize": "12px"}}}%%
flowchart TB
subgraph cal["Your calendars, any number"]
direction LR
G1["Google 1"]
G2["Google 2"]
G3["Google N"]
O1["Outlook 1"]
O2["Outlook N"]
end
cal --> AGG["Free/busy aggregation<br>merged in UTC"]
AGG --> AV["Availability engine<br>open slots, revalidated at booking time"]
AV --> BP["Public booking page"]
AV --> AG["Claude agent"]
AG --> WA["WhatsApp<br>control, alerts, reminders"]
AG --> EM["Email<br>confirmations, reminders"]
BP --> DEST["Destination calendar<br>event and invite written back"]
Loading
Features
Calendar sync across providers. Aggregates free/busy from any number of Google
and Microsoft (Outlook) accounts, each connected by OAuth. Bookings are written to
the calendar you designate.
A booking page that only offers real openings. An agent continuously checks your
merged calendars for open slots. Availability is revalidated at booking time, so a
slot that was just taken cannot be double-booked.
WhatsApp direct communication. Text the agent to create, move, or cancel events,
check availability, or set reminders. New bookings, cancellations, and reschedules
are pushed to you on WhatsApp and email the moment they happen.
Reminder emails and WhatsApp reminders. Attendees and the owner get meeting
reminders over email and approved WhatsApp templates. A claim/dead-letter worker
retries failed sends instead of dropping them.
Ask-the-assistant on the booking page. Visitors can chat with a public agent to
find a time. It is handed exactly two scope-fenced tools. The private tools are
never constructed for it, so safety comes from architecture, not from prompts.
Daily morning brief. A DST-safe WhatsApp summary of the day's meetings and tasks.
Self-serve reschedule and cancel. Every confirmation carries a signed link the
attendee can use without an account.
Chrome extension. Overlays your real availability on other people's Calendly
pages.
Booking, agents, reminders, and calendar extras (personal blocks, birthdays,
follow-ups) are plain JSON APIs under src/app/api/ , so you can reuse them from any
client.
Agent-to-agent scheduling. A visitor's agent and your agent negotiate a meeting
time over a shared protocol and book the result.
git clone https://github.com/hunterZh37/agentic-scheduling.git && cd agentic-scheduling
npm install
cp .env.example .env # defaults work out of the box
docker compose up -d # local Postgres
npx prisma migrate dev && npx prisma db seed
npm run dev # http://localhost:3000
The app boots fully on placeholder config, so you can click through the UI before
connecting anything real. When you are ready to go live,
docs/SETUP.md walks through each credential (Google, Microsoft,
Anthropic, Twilio, Resend) one at a time, in any order.
All identity comes from environment variables: owner name, email, timezone, phone
numbers, and calendar credentials (see .env.example ). The one file
with brand defaults baked in is
src/lib/booking/publicConfig.ts (practice name,
domain, consulting areas, research link). Edit it when deploying your own instance.
Everything is environment variables (see .env.example ). No personal
data is hardcoded.
docs/SETUP.md : credentials, mapped to each env var
docs/DEPLOY.md : deploying to Vercel
docs/SMS.md : the WhatsApp/SMS control channel
docs/MORNING_BRIEF.md : the daily WhatsApp brief
Next.js (App Router), React 19, and TypeScript on Vercel. Postgres via Prisma. Claude
tool-calling for the agents. Twilio for SMS/WhatsApp, Resend for email, VAPID for web
push. Luxon and rrule for time and recurrence. Vitest for tests.
Issues and PRs welcome. Run npm test and npm run lint before submitting.
Your calendar, run by an AI agent you can text. Self-hostable scheduling: Google/Outlook sync, a booking page with only real openings, WhatsApp control and reminders.
bookwithhunter.com/book Topics
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
