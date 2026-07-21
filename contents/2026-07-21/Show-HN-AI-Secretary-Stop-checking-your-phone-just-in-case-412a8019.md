---
source: "https://github.com/mathigatti/telegram-ai-secretary"
hn_url: "https://news.ycombinator.com/item?id=48986826"
title: "Show HN: AI Secretary – Stop checking your phone \"just in case\""
article_title: "GitHub - mathigatti/telegram-ai-secretary: Self-hosted AI notification filter for Telegram that sends only important alerts via ntfy · GitHub"
author: "mathigatti"
captured_at: "2026-07-21T01:45:23Z"
capture_tool: "hn-digest"
hn_id: 48986826
score: 1
comments: 0
posted_at: "2026-07-21T00:50:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Secretary – Stop checking your phone "just in case"

- HN: [48986826](https://news.ycombinator.com/item?id=48986826)
- Source: [github.com](https://github.com/mathigatti/telegram-ai-secretary)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T00:50:41Z

## Translation

タイトル: HN を表示: AI 秘書 – 「念のため」携帯電話をチェックするのはやめてください
記事タイトル: GitHub - mathigatti/telegram-ai-secretary: ntfy 経由で重要なアラートのみを送信する Telegram 用の自己ホスト型 AI 通知フィルター · GitHub
説明: ntfy 経由で重要なアラートのみを送信する Telegram 用の自己ホスト型 AI 通知フィルター - mathigatti/telegram-ai-secretary

記事本文:
GitHub - mathigatti/telegram-ai-secretary: ntfy 経由で重要なアラートのみを送信する Telegram 用の自己ホスト型 AI 通知フィルター · GitHub
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
マティガッティ
/
電報-ai-秘書
公共
通知
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
mathigatti/電報-ai-秘書
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット systemd systemd .env.example .env.example .gitignore .gitignore ライセンス ライセンス README.md README.md notification_agent.py notification_agent.py notification_rules.py notification_rules.py notification_state.py notification_state.py要件.txt要件.txttelegram_notification_listener.py telegram_notification_listener.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
重要なことを見逃すことなく、Telegram をミュートにしておきます。
Telegram AI セクレタリーは、Telegram 用の自己ホスト型 AI 通知フィルターです。それ
Telethon で Telegram ユーザー アカウントをリッスンし、騒々しいチャットをフィルタリングし、質問します。
LLM はコンテキストが重要な場合に、重要なアラートのみを携帯電話に送信します。
NTFY。
デフォルトで通知を見逃さずにオフにしたい人向けに設計されています
緊急のメッセージ、電話、同日の予定、または玄関で待っている人など。
携帯電話は私たちとのつながりを保ってくれますが、騒音が非常に大きくなり、多くの人が
通知をオンにしておくことができません。それらをオフにすると中断が解決されます
問題はありますが、別の問題が発生します。数分ごとに電話をチェックすることです。
万が一に備えて」誰かが重要なことを書いた。
このループは集中力に悪影響を及ぼし、電話の中毒性を低下させるどころか、さらに高めてしまいます。
秘書は長い間、忙しい人々が個人的に連絡をとらなくても連絡が取れるよう支援してきました。
すべての割り込みを処理します。彼らは無関係なメッセージをフィルタリングし、理解します
コンテキストを把握し、実際に注意が必要なときにすぐに通知します。 LLM
その軽量の個人版を誰でも利用できるようにします。
Telegram AI 秘書は、電話を静かに保つという方向の実験です
デフォルトでは、インポートを許可します

nt メッセージがブレークスルーします。
緊急の DM を受信して​​いる間は、Telegram をミュートにしておきます。
誰かが「家にいますか？」と尋ねたときに通知を受け取ります。または「今日会えますか？」
通話や当日調整をバイパスして邪魔にならないようにします。
言及されない限り、グループのおしゃべりは無視してください。
誰かが外で待っているときに重要な ntfy アラートを送信します。
最近のプライベート チャットのコンテキストを使用して、「整理整頓したので、13 時に通ります」とします。
前の「教えてください」の後に重要であると理解できます。
ユーザー アカウントからの受信 Telegram メッセージと通話イベントをリッスンします。
あなたに言及しない限り、騒々しいグループ メッセージを削除します。
呼び出しと一致する高速ルールを直接 ntfy に送信します。
最近の会話のコンテキストを含むプライベートな人間のメッセージをクロードに送信します。
現在のプライベート会話セグメントからのメッセージを最大 20 件保持します。
連続するメッセージが 24 時間以上離れている場合、セグメントはカットされます。
緊急でない場合でも、同日の調整を通知に値するものとして扱います。
物理的に直接存在するメッセージを重要なアラートの候補としてマークします。
CLI からの一時的な高速ルールと応答不可状態をサポートします。
デフォルトでは音声/オーディオ メッセージを無視しますが、オプションで DND 自動応答を使用できます。
https://my.telegram.org/apps の Telegram API ID とハッシュ
クロード CLI がインストールされ、認証された
ntfy トピック (例: https://ntfy.sh)
git clone https://github.com/mathigatti/telegram-ai-secretary.git
cd電報-ai-秘書
python3 -m venv .venv
.venv/bin/pip install -rrequirements.txt
cp .env.example .env
Telegram API 資格情報と ntfy トピックを使用して .env を編集します。
リスナーを対話的に 1 回実行して Telethon セッションを作成します。
.venv/bin/python3 telegram_notification_listener.py
Telegram は初回に電話のログイン コードを尋ねます。
.venv/bin/python3 telegram_notification_listener.py
別のところで

端末で意思決定エージェントをテストできます。
echo ' {"sender_name":"Alex","sender_is_bot":false,"is_private":true,"event_type":"message","text":"帰宅していますか? 今日は寄れます"} ' \
| .venv/bin/python3 notification_agent.py --dry-run
システムド
テンプレートをコピーし、パス/ユーザーを編集します。
sudo cp systemd/telegram-ai-secretary-listener.service /etc/systemd/system/
sudo systemctl デーモン-リロード
sudo systemctl Enable --now telegram-ai-secretary-listener.service
systemctl status telegram-ai-secretary-listener.service --no-pager
テンプレートは、アプリが /opt/telegram-ai-secretary に存在することを前提としています。場合は調整してください
別の場所にデプロイします。
高速ルールは、明らかなケースでのモデル呼び出しを回避します。彼らはに住んでいます
notification-agent/rules.json であり、次のもので管理できます。
.venv/bin/python3 notification_rules.py リスト
.venv/bin/python3 notification_rules.py dnd ステータス
.venv/bin/python3 notification_rules.py dnd on --duration 2h --reason " フォーカス時間 "
.venv/bin/python3 notification_rules.py dnd オフ
.venv/bin/python3 notification_rules.py add --description " 通話時に通知する " --action Notice --call --duration four --urgency high --notification-text " {sender}: {event_label} "
.venv/bin/python3 notification_rules.py add --description " グループのチャットを 2 時間無視する " --action destroy --condition-json ' {"private": false} ' --duration 2h
.venv/bin/python3 notification_rules.py <rule_id> を削除します
期間は、30 分、2 時間、1 日、今日、または永久に設定できます。
アプリは、以下を含むローカル notification-agent/ ディレクトリを作成します。
rules.json : 決定的な高速ルールと DND 状態。
settings.json : 自然言語の通知設定。
settings.events.jsonl : 追加専用の設定/ルールの履歴。
Decision.jsonl : 追加専用の決定とモデルのメタデータ。
onsers.jsonl : 軽量のプライベート メッセージ履歴に使用されます。

コンテクスト。
audio_auto_replies.json : 音声自動応答のクールダウンを停止します。
これらのファイルには、プライベート メッセージ、名前、テレグラム ID、および ntfy の詳細が含まれる場合があります。
これらは git によって無視されます。
リスナーは Telegram イベントを受信し、トリガーを構築します。
トリガーがプライベートな人間の DM であり、それを決定する高速なルールがない場合、エージェントは
クロード プロンプトに最近の会話コンテキストが含まれます。
クロードは、notify 、 urgency 、reason 、 notification_text を含む JSON を返します。
そして自信。
Notice=true の場合、エージェントはテキストを ntfy に送信します。
フローチャート LR
A[テレグラムユーザーアカウント] --> B[テレグラムリスナー]
B --> C[ファストルール]
C -->|破棄| D[通知なし]
C -->|通知| E[プッシュしない]
C -->|判断が必要| F[最近の状況を含むクロード]
F -->|notify=true| E
F -->|notify=false| D
読み込み中
緊急度は ntfy 優先度にマップされます。
実際の電話の音は、ntfy アプリとオペレーティング システムによって制御されます。
通知チャンネルの設定。
これは電報転送業者とどう違うのですか?
Telegram AI 秘書は転送ツールではありません。遮断フィルターです。
すべてのメッセージを別の場所にコピーしようとするわけではありません。それは、
メッセージはあなたの邪魔をするに値します。
ほとんどの Telegram フォワーダーは、チャネルに基づいてメッセージをルーティングするように最適化されています。
キーワードまたは正規表現。このプロジェクトは個人向けに最適化されています
注意管理: ダイレクト メッセージ、コンテキスト、緊急性、同日調整、
電話通知の優先順位。
このアプリはプライベート Telegram メッセージをローカルで処理し、選択したメッセージを送信することがあります
メッセージの内容をクロードとntfyに送信します。 notification_agent.py のプロンプトを確認します。
機密ファイルで実行する前に、notification-agent/ 内のデータ ファイルを削除します。
アカウント。
Telegram AI 長官は、いくつかのより広範なアイデアとツールに取り組んでいます。
秘書、このツールの人間の役割
から借りています。
注目の経済、
経済的インセンティブ

多くの気を散らすインターフェイスの背後にあります。
認知戦争、さらなる
認識と注意をめぐる競争の極端な枠組み。
YIHAD COTRA LAS CORPORACIONES DE LA DISTRACSION とその
懲戒ページ。
デジタル自動決定のサイバーセキュリティを強化します。
ニュースフィード撲滅者、
ソーシャル ネットワークから自動フィードを削除するブラウザ拡張機能。
意図的に開いて検索するツールのように動作します。
複数の LLM プロバイダーをサポートします。
決定事項を確認し、ルールを調整するための小さな Web UI を追加します。
電話の音の重要な通知チャネルを改善し、より臨場感を高めます。
通常のプッシュ通知よりも着信音が変わります。
ntfy 経由で重要なアラートのみを送信する Telegram 用の自己ホスト型 AI 通知フィルター
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

Self-hosted AI notification filter for Telegram that sends only important alerts via ntfy - mathigatti/telegram-ai-secretary

GitHub - mathigatti/telegram-ai-secretary: Self-hosted AI notification filter for Telegram that sends only important alerts via ntfy · GitHub
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
mathigatti
/
telegram-ai-secretary
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
mathigatti/telegram-ai-secretary
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits systemd systemd .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md notification_agent.py notification_agent.py notification_rules.py notification_rules.py notification_state.py notification_state.py requirements.txt requirements.txt telegram_notification_listener.py telegram_notification_listener.py View all files Repository files navigation
Keep Telegram muted without missing what matters.
Telegram AI Secretary is a self-hosted AI notification filter for Telegram. It
listens to a Telegram user account with Telethon, filters noisy chats, asks an
LLM when context matters, and sends only important alerts to your phone through
ntfy.
It is designed for people who want notifications off by default without missing
urgent messages, calls, same-day plans, or someone waiting at the door.
Mobile phones keep us connected, but they have become so noisy that many people
cannot keep notifications turned on. Turning them off solves the interruption
problem, but creates a different one: checking the phone every few minutes "just
in case" someone wrote something important.
That loop is bad for focus and makes the phone more addictive, not less.
Secretaries have long helped busy people stay reachable without personally
processing every interruption. They filter irrelevant messages, understand
context, and notify immediately when something actually needs attention. LLMs
make a lightweight personal version of that possible for everyone.
Telegram AI Secretary is an experiment in that direction: keep the phone quiet
by default, but still let important messages break through.
Keep Telegram muted while still receiving urgent DMs.
Get notified when someone asks "are you home?" or "can we meet today?"
Let calls and same-day coordination bypass do-not-disturb.
Ignore group chatter unless you are mentioned.
Send critical ntfy alerts when someone is waiting outside.
Use recent private-chat context so "I organized myself, I will pass by at 13"
can be understood as important after a previous "please let me know".
Listens to incoming Telegram messages and call events from a user account.
Drops noisy group messages unless they mention you.
Sends calls and matching fast rules directly to ntfy.
Sends private human messages to Claude with recent conversation context.
Keeps up to 20 messages from the current private conversation segment, where a
segment is cut when consecutive messages are more than 24 hours apart.
Treats same-day coordination as notify-worthy, even when it is not an emergency.
Marks immediate physical-presence messages as candidates for critical alerts.
Supports temporary fast rules and do-not-disturb state from the CLI.
Ignores voice/audio messages by default, with optional DND auto-replies.
A Telegram API ID and hash from https://my.telegram.org/apps
The claude CLI installed and authenticated
An ntfy topic, for example on https://ntfy.sh
git clone https://github.com/mathigatti/telegram-ai-secretary.git
cd telegram-ai-secretary
python3 -m venv .venv
.venv/bin/pip install -r requirements.txt
cp .env.example .env
Edit .env with your Telegram API credentials and ntfy topic.
Run the listener once interactively to create the Telethon session:
.venv/bin/python3 telegram_notification_listener.py
Telegram will ask for your phone login code the first time.
.venv/bin/python3 telegram_notification_listener.py
In another terminal you can test the decision agent:
echo ' {"sender_name":"Alex","sender_is_bot":false,"is_private":true,"event_type":"message","text":"Are you home? I can stop by today"} ' \
| .venv/bin/python3 notification_agent.py --dry-run
systemd
Copy the template and edit the paths/user:
sudo cp systemd/telegram-ai-secretary-listener.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now telegram-ai-secretary-listener.service
systemctl status telegram-ai-secretary-listener.service --no-pager
The template assumes the app lives at /opt/telegram-ai-secretary ; adjust it if
you deploy elsewhere.
Fast rules avoid model calls for obvious cases. They live in
notification-agent/rules.json and can be managed with:
.venv/bin/python3 notification_rules.py list
.venv/bin/python3 notification_rules.py dnd status
.venv/bin/python3 notification_rules.py dnd on --duration 2h --reason " focus time "
.venv/bin/python3 notification_rules.py dnd off
.venv/bin/python3 notification_rules.py add --description " Notify on calls " --action notify --call --duration forever --urgency high --notification-text " {sender}: {event_label} "
.venv/bin/python3 notification_rules.py add --description " Ignore group chatter for 2h " --action discard --condition-json ' {"private": false} ' --duration 2h
.venv/bin/python3 notification_rules.py remove < rule_id >
Durations can be 30m , 2h , 1d , today , or forever .
The app creates a local notification-agent/ directory containing:
rules.json : deterministic fast rules and DND state.
preferences.json : natural-language notification preferences.
preferences.events.jsonl : append-only preference/rule history.
decisions.jsonl : append-only decisions and model metadata.
conversations.jsonl : lightweight private-message history used for context.
audio_auto_replies.json : DND audio auto-reply cooldowns.
These files may contain private messages, names, Telegram IDs, and ntfy details.
They are ignored by git.
The listener receives a Telegram event and builds a trigger.
If the trigger is a private human DM and no fast rule decides it, the agent
includes recent conversation context in the Claude prompt.
Claude returns JSON with notify , urgency , reason , notification_text ,
and confidence .
If notify=true , the agent sends the text to ntfy.
flowchart LR
A[Telegram user account] --> B[Telethon listener]
B --> C[Fast rules]
C -->|discard| D[No notification]
C -->|notify| E[ntfy push]
C -->|needs judgment| F[Claude with recent context]
F -->|notify=true| E
F -->|notify=false| D
Loading
Urgency maps to ntfy priority:
The actual phone sound is controlled by the ntfy app and your operating system's
notification channel settings.
How Is This Different From A Telegram Forwarder?
Telegram AI Secretary is not a forwarding tool. It is an interruption filter.
It does not try to copy every message somewhere else; it decides whether a
message deserves to interrupt you.
Most Telegram forwarders are optimized for routing messages based on channels,
keywords, or regular expressions. This project is optimized for personal
attention management: direct messages, context, urgency, same-day coordination,
and phone notification priority.
This app processes private Telegram messages locally and may send selected
message contents to Claude and ntfy. Review the prompt in notification_agent.py
and the data files in notification-agent/ before running it on sensitive
accounts.
Telegram AI Secretary sits near a few broader ideas and tools:
Secretary , the human role this tool
is borrowing from.
Attention economy , the
economic incentive behind many distracting interfaces.
Cognitive warfare , a more
extreme framing of competition over perception and attention.
YIHAD CONTRA LAS CORPORACIONES DE LA DISTRACCION and its
disciplina page.
Guia cyberciruja para la autodeterminacion digital .
News Feed Eradicator ,
a browser extension that removes automatic feeds from social networks so they
behave more like tools you intentionally open and search.
Support multiple LLM providers.
Add a small web UI for reviewing decisions and tuning rules.
Improve critical notification channels for phone sounds that feel closer to a
ringtone than a normal push notification.
Self-hosted AI notification filter for Telegram that sends only important alerts via ntfy
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
