---
source: "https://github.com/amita-seal/wa-metaai"
hn_url: "https://news.ycombinator.com/item?id=49138536"
title: "Using Meta AI via WhatsApp for OpenCode"
article_title: "GitHub - amita-seal/wa-metaai: Use Meta AI (inside WhatsApp) as an OpenAI-compatible LLM provider for opencode, with synthesized tool calling · GitHub"
author: "amita2"
captured_at: "2026-08-01T21:44:31Z"
capture_tool: "hn-digest"
hn_id: 49138536
score: 1
comments: 0
posted_at: "2026-08-01T21:10:44Z"
tags:
  - hacker-news
  - translated
---

# Using Meta AI via WhatsApp for OpenCode

- HN: [49138536](https://news.ycombinator.com/item?id=49138536)
- Source: [github.com](https://github.com/amita-seal/wa-metaai)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T21:10:44Z

## Translation

タイトル: WhatsApp 経由で OpenCode にメタ AI を使用する
記事タイトル: GitHub - amita-seal/wa-metaai: 合成ツール呼び出しを使用して、オープンコード用の OpenAI 互換 LLM プロバイダーとして Meta AI (WhatsApp 内) を使用する · GitHub
説明: Meta AI (WhatsApp 内) をオープンコード用の OpenAI 互換 LLM プロバイダーとして使用し、合成ツール呼び出し - amita-seal/wa-metaai

記事本文:
GitHub - amita-seal/wa-metaai: 合成ツール呼び出しを使用して、オープンコード用の OpenAI 互換 LLM プロバイダーとして Meta AI (WhatsApp 内) を使用する · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
阿弥陀印
/
わめたあい
公共
通知
あなたはきっとそうでしょう

サインインして通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .gitignore .gitignore README.md README.md go.mod go.mod go.sum go.sum main.go main.go main_test.go main_test.go すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Meta AI によってサポートされる OpenAI 互換のエンドポイントは、WhatsApp 経由で到達します。オープンコード (または任意のコード) を実行しましょう
OpenAI 互換クライアント) は、Meta AI をモデルとして使用します。
opencode --> POST localhost:8788/v1/chat/completions --> whatsmeow --> WhatsApp
<-- メタ AI (867051314767696@bot)
なぜベイリーズではなくワッツミャウなのか
メタ AI は電話番号による連絡先ではありません。これはボット JID : 867051314767696@bot です。ボットへの配信
送信スタンザには 2 つのことが必要です。
スタンザコンテンツに追加された <bot> ノード
HKDF 由来の BotMessageSecret (メッセージ シークレットに BotMessageHKDF を適用)
Baileys はどちらも実装していません。送信パス全体がボット JID を一度参照し、発行をスキップするだけです。
TCトークン。観察可能な症状は正確であり、誤解を招くものです。WhatsApp サーバーがメッセージを確認します。
( status: 2 ) その後、決して配信されないため、永久に 1 つの灰色のティックのままになり、メタ AI は決して配信しません
と答える。同じコードパス上の人間の JID へのコントロール送信は保留→ 2 → 3 → 4 (読み取り) になりました。
トランスポートが正常であり、欠落しているのはボットのサポートだけであることが証明されました。
それは何ですか：types.NewMetaAIJID は正確に 867051314767696@bot 、send.go は isBotMode を設定します
ターゲット IsBot() の場合、ボット シークレットを取得し、 <bot> ノードをアタッチします。
Baileys の META_AI_JID = 13135550002@c.us はレガシーでルーティング不可能です。これに送信すると、
USync フェッチでは、保留中の PN の結果は得られませんでした。これは、その番号がユーザーではないことを WhatsApp が示しています。
アカウント自体は実際の WhatsApp モバイル アプリ内に存在する必要があります (Android エミュレーターは動作します)。このサービス

はリンクされたデバイスとして接続され、単独で登録を保持することはできません。
CGO_ENABLED=1 go build -o wametaai 。
WA_PHONE= < 数字、国コードが最初 > ./wametaai
最初の実行ではペアリング コードが出力されます。それを [WhatsApp] > [リンクされたデバイス] > [電話とのリンク] に入力します。
番号。セッションは wametaai.db に保持されるため、後の実行にはコードは必要ありません。
~/.config/opencode/opencode.json :
{
「プロバイダー」: {
"WhatsApp" : {
"npm" : " @ai-sdk/openai-compatibility " ,
"名前" : " WhatsApp (メタ AI) " ,
"options" : { "baseURL" : " http://localhost:8788/v1 " , "apiKey" : " 未使用 " },
「モデル」: {
"meta-ai" : { "name" : "メタ AI (WhatsApp) " 、 "tool_call" : true 、 "limit" : { "context" : 8000 、 "output" : 4000 } }
}
}
}
}
次に、opencode を実行して --model whatsapp/meta-ai "..." します。
メタ AI にはネイティブ関数呼び出しがないため、シムがそれを提供します。リクエストにツールが含まれる場合、
toolPrompt は、各ツールの名前、説明、JSON スキーマをプロンプトにレンダリングし、
単一のフェンスで囲まれたオブジェクト:
{ "ツール" : " <名前> " 、 "引数" : { ... }}
parseToolCall はそれを抽出します — フェンス、言語タグ、周囲のおしゃべりを許容します
Meta AI は追加を好みます。そして、応答は実際の OpenAI tools_calls メッセージとして返されます。
finish_reason: "tool_calls" 、JSON パスと SSE パスの両方。以前のターンは次のように再生されます
[ASSISTANT run tools] と [TOOL RESULT] で、すでに実行されたものを確認できます。
オープンコードの実際のエージェント ループで動作することを確認しました。
$ opencode run --model whatsapp/meta-ai "notes.txt ファイルを読んで、秘密の言葉を教えてください。"
> ビルド・メタアイ
→notes.txtを読む
ザクロ
観察: opencode の完全なシステム プロンプトとそのツール スキーマは、最大 37,000 文字 (最大 9,000 文字) になります。
トークン) をリクエストごとに生成し、Meta AI は問題なく処理し、相対パスを正しく解決しました。
絶対的なものに。 2 往復、それぞれ最大 6 秒。
セイ

注意: 単一の WhatsApp テキスト メッセージは約 65,000 文字に制限されます。
簡単な 1 ファイルのタスクがすでに 37k を使用しています。より大きなコンテキスト (より多くのファイル、より長い履歴) がそれに該当します。
現在、切り捨てガードがないため、サイズが大きすぎるプロンプトは機能が低下するのではなく失敗します。
ツール呼び出しの信頼性
ネイティブではなく、プロンプトが表示されます。 Meta AI は消費者アシスタントであるため、ツールの呼び出しが必要な場合に散文で応答できます。ループは回復するのではなく、そのターンで停止します。
ストリーミング
応答はバッファリングされ、1 つのチャンクとして送信されます。 Meta AI は、これまでの完全な回答を含む 1 つのメッセージを編集することでストリーミングします。これらの編集は MESSAGE_EDIT タイプのプロトコルメッセージとして到着しますが、これは履歴同期パス上でのみラップを解除するため、ライブ編集は手動でラップを解除する必要があります。そうしないと、応答は最初のフラグメントでスタックしたままになります。完了は、無音の WA_QUIET_MS に加えて、テキストが文の途中で終了しないというヒューリスティックと、制限された追加の待機として WA_GRACE_MS を使用して推測されます。
システムプロンプト
平らな転写物に折りたたまれます。個別のシステム役割はありません
会話状態
独自のメモリを備えた 1 つの WhatsApp スレッドにより、履歴がターンごとに再送信され、Meta AI 自身のリコールがリクエスト全体に波及する可能性があります。
同時実行性
シリアル化 — 単一スレッドは並列リクエストを処理できません
トークンの使用法
トークンあたり最大 4 文字と推定されます。 WhatsApp は何も報告しません
レイテンシ
短い回答の場合は約 7 秒
環境
変数
デフォルト
WA_PORT
8788
HTTPポート
WA_PHONE
—
初回ペアリングの場合のみ必要
WA_QUIET_MS
6000
このくらいの沈黙の後、返信は完了したとみなされます
WA_GRACE_MS
18000
返信がまだ完了していないように見える場合はさらに待機します
WA_TIMEOUT_MS
120000
リクエストごとのハードキャップ
アカウントの耐久性
レンタルした TextVerified 番号に登録されているため、アカウントはレンタル期間中のみ有効です。
解放された番号は、他の人が再レンタルして再登録することができます。

このアカウントを削除したいと思います。
WhatsApp は、プライマリがオフラインのままの場合、リンクされたデバイスも期限切れにするため、エミュレータにアクセスできる状態にしておいてください。
非公式クライアントによる使用は WhatsApp の規約に反しており、番号が禁止されるリスクがあります。
合成されたツール呼び出しを使用して、オープンコード用の OpenAI 互換 LLM プロバイダーとして Meta AI (WhatsApp 内) を使用する
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Use Meta AI (inside WhatsApp) as an OpenAI-compatible LLM provider for opencode, with synthesized tool calling - amita-seal/wa-metaai

GitHub - amita-seal/wa-metaai: Use Meta AI (inside WhatsApp) as an OpenAI-compatible LLM provider for opencode, with synthesized tool calling · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
amita-seal
/
wa-metaai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .gitignore .gitignore README.md README.md go.mod go.mod go.sum go.sum main.go main.go main_test.go main_test.go View all files Repository files navigation
OpenAI-compatible endpoint backed by Meta AI, reached over WhatsApp. Lets opencode (or any
OpenAI-compatible client) use Meta AI as a model.
opencode --> POST localhost:8788/v1/chat/completions --> whatsmeow --> WhatsApp
<-- Meta AI (867051314767696@bot)
Why whatsmeow and not Baileys
Meta AI is not a phone-number contact. It is a bot JID : 867051314767696@bot . Delivery to a bot
requires two things on the outgoing stanza:
a <bot> node appended to the stanza content
an HKDF-derived BotMessageSecret ( applyBotMessageHKDF over the message secret)
Baileys implements neither — its entire send path references bot JIDs once, only to skip issuing a
TC token. The observable symptom is precise and misleading: WhatsApp server-acks the message
( status: 2 ) and then never delivers it, so it sits at one grey tick forever and Meta AI never
replies. A control send to a human JID on the same code path went PENDING → 2 → 3 → 4 (read),
proving the transport was fine and only bot support was missing.
whatsmeow has it: types.NewMetaAIJID is exactly 867051314767696@bot , send.go sets isBotMode
when the target IsBot() , derives the bot secret, and attaches the <bot> node.
Baileys' META_AI_JID = 13135550002@c.us is legacy and unroutable — sending to it produces
USync fetch yielded no results for pending PNs , WhatsApp's way of saying that number isn't a user.
The account itself must live in a real WhatsApp mobile app (an Android emulator works); this service
attaches as a linked device and cannot hold the registration on its own.
CGO_ENABLED=1 go build -o wametaai .
WA_PHONE= < digits, country code first > ./wametaai
First run prints a pairing code — enter it under WhatsApp > Linked devices > Link with phone
number . The session persists in wametaai.db , so later runs need no code.
~/.config/opencode/opencode.json :
{
"provider" : {
"whatsapp" : {
"npm" : " @ai-sdk/openai-compatible " ,
"name" : " WhatsApp (Meta AI) " ,
"options" : { "baseURL" : " http://localhost:8788/v1 " , "apiKey" : " unused " },
"models" : {
"meta-ai" : { "name" : " Meta AI (WhatsApp) " , "tool_call" : true , "limit" : { "context" : 8000 , "output" : 4000 } }
}
}
}
}
Then opencode run --model whatsapp/meta-ai "..." .
Meta AI has no native function calling, so the shim supplies it. When a request carries tools ,
toolPrompt renders each tool's name, description and JSON Schema into the prompt and asks for a
single fenced object:
{ "tool" : " <name> " , "args" : { ... }}
parseToolCall extracts that — tolerating the fence, a language tag, and the surrounding chatter
Meta AI likes to add — and the reply is returned as a real OpenAI tool_calls message with
finish_reason: "tool_calls" , in both the JSON and SSE paths. Prior turns are replayed as
[ASSISTANT ran tool] and [TOOL RESULT] so it can see what already ran.
Verified working through opencode's real agent loop:
$ opencode run --model whatsapp/meta-ai "Read the file notes.txt and tell me the secret word."
> build · meta-ai
→ Read notes.txt
pomegranate
Observed: opencode's full system prompt plus its tool schemas came to ~37,000 characters (~9k
tokens) per request, and Meta AI handled it without complaint, correctly resolving a relative path
to an absolute one. Two round trips, ~6s each.
The ceiling to watch: a single WhatsApp text message caps out around 65k characters, and a
trivial one-file task already used 37k. Larger context — more files, longer histories — will hit that
wall, and there is currently no truncation guard, so an oversized prompt fails rather than degrades.
tool-call reliability
prompted, not native. Meta AI is a consumer assistant, so it can answer in prose where a tool call was wanted; the loop stalls on that turn rather than recovering.
streaming
reply is buffered then emitted as one chunk. Meta AI streams by editing one message with the full answer so far; those edits arrive as a protocolMessage of type MESSAGE_EDIT that whatsmeow only unwraps on its history-sync path, so live edits must be unwrapped by hand or the reply stays stuck at the first fragment. Completion is inferred from WA_QUIET_MS of silence plus a heuristic that the text does not end mid-sentence, with WA_GRACE_MS as a bounded extra wait.
system prompt
folded into a flattened transcript; there is no separate system role
conversation state
one WhatsApp thread with its own memory, so history is resent each turn and Meta AI's own recall can bleed across requests
concurrency
serialized — a single thread cannot serve parallel requests
token usage
estimated at ~4 chars/token; WhatsApp reports none
latency
~7s for a short answer
Env
var
default
WA_PORT
8788
HTTP port
WA_PHONE
—
required only for first-time pairing
WA_QUIET_MS
6000
reply considered complete after this much silence
WA_GRACE_MS
18000
extra wait when the reply still looks unfinished
WA_TIMEOUT_MS
120000
hard cap per request
Account durability
Registered on a rented TextVerified number, so the account lives only as long as that rental — a
released number can be re-rented and re-registered by someone else, which would evict this account.
WhatsApp also expires linked devices when the primary stays offline, so keep the emulator reachable.
Unofficial-client use is against WhatsApp's terms and carries a ban risk for the number.
Use Meta AI (inside WhatsApp) as an OpenAI-compatible LLM provider for opencode, with synthesized tool calling
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
