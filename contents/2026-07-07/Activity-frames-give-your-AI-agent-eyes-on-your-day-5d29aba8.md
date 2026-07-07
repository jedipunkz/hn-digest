---
source: "https://github.com/nossa-y/activity-frames"
hn_url: "https://news.ycombinator.com/item?id=48824758"
title: "Activity-frames – give your AI agent eyes on your day"
article_title: "GitHub - nossa-y/activity-frames: Episodic memory for AI agents. Records your screen locally, compiles it into structured activity frames, serves them over MCP. No cloud, no LLM in the loop. · GitHub"
author: "nossa-y"
captured_at: "2026-07-07T22:52:04Z"
capture_tool: "hn-digest"
hn_id: 48824758
score: 1
comments: 0
posted_at: "2026-07-07T22:18:40Z"
tags:
  - hacker-news
  - translated
---

# Activity-frames – give your AI agent eyes on your day

- HN: [48824758](https://news.ycombinator.com/item?id=48824758)
- Source: [github.com](https://github.com/nossa-y/activity-frames)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T22:18:40Z

## Translation

タイトル: アクティビティ フレーム – AI エージェントにあなたの 1 日の様子を知らせます
記事のタイトル: GitHub - nossa-y/activity-frames: AI エージェントのエピソード記憶。画面をローカルに記録し、構造化されたアクティビティ フレームにコンパイルして、MCP 経由で提供します。クラウドもループ内に LLM もありません。 · GitHub
説明: AI エージェントのエピソード記憶。画面をローカルに記録し、構造化されたアクティビティ フレームにコンパイルして、MCP 経由で提供します。クラウドもループ内に LLM もありません。 - nossa-y/アクティビティフレーム

記事本文:
GitHub - nossa-y/activity-frames: AI エージェントのエピソード記憶。画面をローカルに記録し、構造化されたアクティビティ フレームにコンパイルして、MCP 経由で提供します。クラウドもループ内に LLM もありません。 · GitHub
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
ノッサっぽい
/
アクティビティフレーム
公共
通知

イオン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github .github 例 例 src/ activity_frames src/ activity_frames テスト テスト .gitignore .gitignore ACKNOWLEDGMENTS.md ACKNOWLEDGMENTS.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスREADME.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
デスクトップ アプリをダウンロード - Nocta はアクティビティ フレームを使用してあなたの仕事の様子を監視し、注意が必要なことを毎日説明します。 100%地元産。
AI エージェントのエピソード記憶。
エージェントはコードを読み取ったり、Web を検索したり、API を呼び出したりすることができますが、あなたが過去 8 時間何をしていたのかは知りません。すべての会話はブラインドで始まります。
activity-frame はエージェントに目を与えます。画面をローカルに記録し、表示された内容を構造化されたアクティビティ フレーム (実際に行ったことの限定されたエピソード) にコンパイルし、MCP 経由でエージェントに提供します。クラウドもパイプラインに LLM も存在せず、推測も必要ありません。
pip インストールアクティビティフレーム
aframes Record # キャプチャを開始します (ローカル、デフォルトで音声オフ)
aframes context # 過去 2 時間、エージェント対応
エージェントが見ているもの
Capture はインスタントを保存します。1 日に何千ものスナップショット行があり、それぞれが「22:53:05 に Chrome で linkedin.com/in/... が表示されました」と表示されます。理屈をこねても無駄だ。
activity-frames は、これらの瞬間をエピソードにコンパイルします。
- id : f-0007
アプリ：Google Chrome
サイト：linkedin.com
開始 : " 20:24:04 "
終了 : " 20:42:11 "
持続時間分 : 18.0
ページ:
- {種類: people_search、エンティティ: "cto paris"、カウント: 2}
- {種類: プロフィール、エンティティ: ナジムザマン}
- {種類: 会社、エンティティ: nexdotai}
入力: {キー数: 214、クリック数: 31}
証拠: {frame_ids: "99871..100147"

}
そして、任意のシステム プロンプトのコンパクトなコンテキスト ブロックに次のように記述します。
ユーザーアクティビティ (2026 年 7 月 4 日、現地時間、画面キャプチャから測定、解釈なし):
対象範囲: 09:12 ～ 18:47、アクティブ時間 342、アプリ数 11
遠方: 12:30-13:15 (45m)
- 09:12-09:58 カーソル (46.2m): main.py - api
- 10:01-10:44 Google Chrome/github.com (41.3 メートル): pull_request:acme/api#412;コード:acme/API
- 20:24-20:42 Google Chrome/linkedin.com (18.0m): people_search:cto paris x2;プロフィール：ナジムザマン。会社名：ネクストタイ
それをプロンプトに入力すると、エージェントがあなたの一日を把握します。丸 1 日のコンパイルは 1 秒以内に完了し、トークンのコストはゼロです。
エピソード記憶、正直に行う
今日のエージェントの記憶とは、会話の記憶、つまりあなたがモデルに話した内容を意味します。エピソード記憶はあなたが実際にやったことです。そして難しいのは、それを嘘をつかずに表現することです。
activity-frames は 2 層のコントラクト ( SPEC.md ) を強制します。
階層 1、測定済み (このパッケージ): すべてはキャプチャ データから決定論的なコードによって導出可能です。セッション、期間、型付きページ エンティティ、入力量、カバレッジ ギャップ。毎回同じ入力、同じ出力。インテント ラベルはありません。コードは、2 つのプロフィール ビュー + 1 つの人物検索が「見込み客」であることを認識できません。それがエージェントの仕事です。それはLLMです。
階層 2、推論 (オプションの拡張): 解釈を追加するツールは、解釈に名前空間を付け、信頼度にタグを付け (高 | 中 | 推測的)、証拠をリンクする必要があります。事実と推測が黙って混ざり合うことは決してありません。
すべてのフレームは、証拠ポインタを生のキャプチャ行に戻します。すべての文書には盲点が存在します。システムが見ていなかったことは、システムは見ていなかったと言うのです。
# クロードコード
クロード mcp アクティビティ フレームを追加 -- aframes mcp
どの MCP クライアントでも動作します: コマンド aframes 、 args ["mcp"] 。 4 つのツール: get_context 、 get_activity 、 get_day_summary 、 get_patterns (反復ワークフローの検出: 反復クリック、URL ループ、毎日の習慣

ts）。
activity_frames からアクティビティログをインポート
ログ = アクティビティログ ()
doc = ログ 。 day () # 今日、構造化された
doc = ログ 。最近 ( hours = 2 ) # 過去 2 時間
print ( log . context ( hours = 2 )) # ペースト可能なコンテキストブロック
プライバシーモデル
ローカルのみ。キャプチャ、保存、コンパイルはすべてマシン上で行われます。どこにもアップロードされることはありません。
読み取り専用のコンパイル。コンパイラはキャプチャ データベースを読み取り専用で開きます。
出力時のコンテンツのオプトイン。コンパイルされたドキュメントには、デフォルトで入力カウントが含まれます。入力されたテキストのコンテンツは、明示的に --include-text を渡した場合にのみ表示されます (これにより、繰り返しテキスト パターン検出器もゲートされます)。境界を明確にしてください。キャプチャ データベース自体は、レコーダーが認識した内容をローカルに保存するため、他の機密ファイル (FileVault、アクセス許可) と同様に保護してください。
デフォルトでは音声はオフになっています。 aframes Record --audio をオプトインします。
コンパイル パスに LLM がありません。コンパイルはプレーン コードであるため、メモリの生成にローカルまたはリモートの言語モデルは関与しません。キャプチャ エンジンはデバイス上で OCR を実行して、画面上の内容を読み取ります。それはあなたのマシンに残ります。
コンテキスト ブロックをエージェントに貼り付けるときに、何を残すかを選択します。ウィンドウ タイトルとページ エンティティは画面から生成され、サードパーティのテキストが含まれる場合があることに注意してください。エージェントはそれらを指示ではなくデータとして扱う必要があります。
キャプチャ エンジン コンパイラ (このパッケージ) エージェント
-------------------------------------- ---------------
画面スナップショット --> セッション化 (滞留、ギャップ、 --> MCP ツール /
アクセシビリティ ツリー フリッカー マージ) コンテキスト ブロック /
入力イベント エンティティ タイピング (20 以上のサイト) JSON、YAML、MD
(ローカル SQLite) エンリッチメント、パターン
デフォルトのキャプチャ エンジンは screenpipe です。aframes レコードは、固定された MIT ライセンス ビルド (v0.3.324) をプロビジョニングし、最初の実行前に公開された sha512 を検証し、それを管理します (ACKNO を参照)

WLEDGMENTS.md )。すでに独自のレコーダーを実行していますか? $AFRAMES_DB を、互換性のある Frames / ui_events / elements テーブルを持つキャプチャ データベースに指定し、aframes レコードを完全にスキップします。
aframes レコード # キャプチャ開始 (--stop / --status / --audio)
aframes today # 今日のフレーム (YAML*)
aframes day 2026-07-03 -f json # いつでも、JSON
aframes context --hours 3 # エージェントコンテキストブロック
aframes apps # アプリごとの時間台帳
aframes パターン --days 7 # 反復的なワークフローの検出
aframes mcp # MCP stdio サーバー
*YAML 出力は PyYAML を使用します ( pip install "activity-frames[yaml]" )。これがないと、CLI は JSON に戻ります。
v0.1。 macOS (Apple Silicon) 上で開発およびテストされています。 Intel macOS および Linux x64 エンジン ビルドは存在しますが、あまり活用されていません。報告は歓迎です。エンティティ パーサーは、LinkedIn、GitHub、Google (Search/Docs/Gmail/Maps/Meet/Calendar)、YouTube、X、Instagram、Reddit、Luma、Partiful、Product Hunt、Vercel、Supabase、Stripe、Discord、Notion、Figma、Stack Overflow、Calendly、ChatGPT/Claude、localhost をカバーします。未知のサイトは一般的なページ参照にフォールバックします。常に完全であり、損失を伴うことはありません。問題やパーサーの PR は歓迎です。
Nocta のメーカーである Nossa Iyamu によって構築されました。マサチューセッツ工科大学
AI エージェントのエピソード記憶。画面をローカルに記録し、構造化されたアクティビティ フレームにコンパイルして、MCP 経由で提供します。クラウドもループ内に LLM もありません。
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

Episodic memory for AI agents. Records your screen locally, compiles it into structured activity frames, serves them over MCP. No cloud, no LLM in the loop. - nossa-y/activity-frames

GitHub - nossa-y/activity-frames: Episodic memory for AI agents. Records your screen locally, compiles it into structured activity frames, serves them over MCP. No cloud, no LLM in the loop. · GitHub
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
nossa-y
/
activity-frames
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github .github examples examples src/ activity_frames src/ activity_frames tests tests .gitignore .gitignore ACKNOWLEDGMENTS.md ACKNOWLEDGMENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml View all files Repository files navigation
Download the desktop app - Nocta uses activity-frames to watch how you work and brief you daily on what needs your attention. 100% local.
Episodic memory for AI agents.
Your agent can read your code, search the web, and call APIs - but it has no idea what you have been doing for the last 8 hours. It starts every conversation blind.
activity-frames gives your agent eyes. It records your screen locally, compiles what it sees into structured activity frames (bounded episodes of what you actually did), and serves them to any agent over MCP. No cloud, no LLM in the pipeline, no guessing.
pip install activity-frames
aframes record # start capturing (local, audio off by default)
aframes context # your last 2 hours, agent-ready
What your agent sees
Capture stores instants: thousands of snapshot rows a day, each one saying "at 22:53:05, Chrome showed linkedin.com/in/...". Useless to reason over.
activity-frames compiles those instants into episodes:
- id : f-0007
app : Google Chrome
site : linkedin.com
start : " 20:24:04 "
end : " 20:42:11 "
duration_min : 18.0
pages :
- {kind: people_search, entity: "cto paris", count: 2}
- {kind: profile, entity: najmuzzaman}
- {kind: company, entity: nexdotai}
input : {keys: 214, clicks: 31}
evidence : {frame_ids: "99871..100147"}
And into a compact context block for any system prompt:
USER ACTIVITY (2026-07-04, local time; measured from screen capture, no interpretation):
coverage: 09:12-18:47, 342 active min, 11 apps
away: 12:30-13:15 (45m)
- 09:12-09:58 Cursor (46.2m): main.py - api
- 10:01-10:44 Google Chrome/github.com (41.3m): pull_request:acme/api#412; code:acme/api
- 20:24-20:42 Google Chrome/linkedin.com (18.0m): people_search:cto paris x2; profile:najmuzzaman; company:nexdotai
Drop that into a prompt and your agent knows your day. A full day compiles in under a second and costs zero tokens.
Episodic memory, done honestly
Agent memory today means conversation memory: what you told the model. Episodic memory is what you actually did - and the hard part is representing it without lying.
activity-frames enforces a two-tier contract ( SPEC.md ):
Tier 1, measured (this package): everything is derivable by deterministic code from capture data. Sessions, durations, typed page entities, input volume, coverage gaps. Same input, same output, every time. There are no intent labels - code cannot know that 2 profile views + a people search was "prospecting". That is your agent's job; it is an LLM.
Tier 2, inferred (optional extension): tools that add interpretation must namespace it, tag confidence ( high | medium | speculative ), and link evidence. Facts and guesses can never silently mix.
Every frame carries evidence pointers back to raw capture rows. Every document declares its blind spots. What the system did not see, it says it did not see.
# Claude Code
claude mcp add activity-frames -- aframes mcp
Any MCP client works: command aframes , args ["mcp"] . Four tools: get_context , get_activity , get_day_summary , get_patterns (repetitive-workflow detection: repeated clicks, URL loops, daily habits).
from activity_frames import ActivityLog
log = ActivityLog ()
doc = log . day () # today, structured
doc = log . recent ( hours = 2 ) # last 2 hours
print ( log . context ( hours = 2 )) # paste-ready context block
Privacy model
Local only. Capture, storage, and compilation all happen on your machine. Nothing is uploaded anywhere, ever.
Read-only compilation. The compiler opens the capture database read-only.
Content opt-in at the output. Compiled documents carry input counts by default; typed-text content appears only if you explicitly pass --include-text (this also gates the repeated-text pattern detector). Be clear about the boundary: the capture database itself does store what the recorder sees, locally, so protect it like any sensitive file (FileVault, permissions).
Audio off by default. aframes record --audio to opt in.
No LLM in the compile path. Compilation is plain code, so no language model, local or remote, is involved in producing memory. The capture engine does run on-device OCR to read what is on screen; that stays on your machine.
You choose what leaves , when you paste a context block into an agent. Note that window titles and page entities originate from your screen and can contain third-party text; agents should treat them as data, not instructions.
capture engine compiler (this package) your agent
------------------ --------------------------- -----------------
screen snapshots --> sessionize (dwell, gaps, --> MCP tools /
accessibility tree flicker merge) context blocks /
input events entity typing (20+ sites) JSON, YAML, md
(local SQLite) enrichment, patterns
The default capture engine is screenpipe : aframes record provisions a pinned, MIT-licensed build (v0.3.324), verifies its published sha512 before first run, and manages it for you (see ACKNOWLEDGMENTS.md ). Already running your own recorder? Point $AFRAMES_DB at any capture database with compatible frames / ui_events / elements tables and skip aframes record entirely.
aframes record # start capture (--stop / --status / --audio)
aframes today # today's frames (YAML*)
aframes day 2026-07-03 -f json # any day, JSON
aframes context --hours 3 # agent context block
aframes apps # per-app time ledger
aframes patterns --days 7 # repetitive workflow detection
aframes mcp # MCP stdio server
*YAML output uses PyYAML ( pip install "activity-frames[yaml]" ); without it the CLI falls back to JSON.
v0.1. Developed and tested on macOS (Apple Silicon); Intel macOS and Linux x64 engine builds exist but are less exercised - reports welcome. Entity parsers cover LinkedIn, GitHub, Google (Search/Docs/Gmail/Maps/Meet/Calendar), YouTube, X, Instagram, Reddit, Luma, Partiful, Product Hunt, Vercel, Supabase, Stripe, Discord, Notion, Figma, Stack Overflow, Calendly, ChatGPT/Claude, localhost; unknown sites fall back to a generic page reference - always total, never lossy. Issues and parser PRs welcome.
Built by Nossa Iyamu , maker of Nocta . MIT.
Episodic memory for AI agents. Records your screen locally, compiles it into structured activity frames, serves them over MCP. No cloud, no LLM in the loop.
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
