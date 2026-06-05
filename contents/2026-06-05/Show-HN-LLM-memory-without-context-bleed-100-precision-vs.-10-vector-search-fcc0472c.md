---
source: "https://tenureai.dev/"
hn_url: "https://news.ycombinator.com/item?id=48409678"
title: "Show HN: LLM memory without context bleed; 100% precision vs. <10% vector search"
article_title: "Tenure | Local AI Memory for Developers - 1.0 Precision, No Cloud, Not RAG"
author: "decorner"
captured_at: "2026-06-05T10:25:54Z"
capture_tool: "hn-digest"
hn_id: 48409678
score: 3
comments: 1
posted_at: "2026-06-05T08:33:58Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LLM memory without context bleed; 100% precision vs. <10% vector search

- HN: [48409678](https://news.ycombinator.com/item?id=48409678)
- Source: [tenureai.dev](https://tenureai.dev/)
- Score: 3
- Comments: 1
- Posted: 2026-06-05T08:33:58Z

## Translation

タイトル: HN を表示: コンテキスト ブリードなしの LLM メモリ。 100% の精度と 10% 未満のベクトル検索
記事タイトル: 在職期間 |開発者向けのローカル AI メモリ - 精度 1.0、クラウドなし、RAG なし
説明: Tenure は、ローカルホスト上で実行される完全にローカルな AI メモリ レイヤーです。すべてのツールおよびすべてのセッションにわたって、永続的で構造化された記憶が続きます。 1.0の検索精度。雲ゼロ。ワークフローの変更はゼロです。

記事本文:
在職期間
ユースケース ▼ モバイル VSCode VSCodium ドキュメント GitHub ライティング
研究 ▼ ベンチマーク ペーパーを無料でインストール Mobile VSCode VSCodium Docs GitHub Writing Benchmark Paper を無料でインストール
完全にローカルホスト上で実行されます。あなたのマシンから何も離れることはありません。
チャット クライアントでブレインストーミングを行います。 IDE を開きます。それはすでに知っています。
ラグではありません。別のメモリラッパーではありません。その前に学習内容を観察する
常に反応を変えます。
あらゆるツールやあらゆる環境であなたを追いかける永続的な記憶
セッション。事前ではなく、準備ができたら注入をオンにします。
// 既存のツールからインストールします
ターミナルの方がいいですか？
カール -fsSL
https://raw.githubusercontent.com/tenurehq/tenure/main/scripts/install.sh
|バッシュ
コピー
精度1.0。スマートなモデルであれば、面倒な掃除をする必要がないからです。
回収。
他の記憶ツールは不正行為を行います。彼らはあなたのチャット履歴全体をダンプするか、流出させます
ベクトル クラスタをコンテキスト ウィンドウに取り込んで、有能なダウンストリームを許可します。
モデルは干し草の山から針を見つけます。彼らは答えに対してSOTAを主張します
正確性は向上しますが、レイテンシとコンテキストにかかる税金を支払う必要があります。
データベースのクエリとフィルタリングの違いは次のとおりです。
アプリケーションコード。どちらが本番環境に属するかは誰もが知っています。
文脈における無関係な信念はすべて、
あなたが支払っているトークンと
待機中のレイテンシー。
結果は再現可能です。 HuggingFace のデータセット。
自分で実行してください
どれくらいの費用がかかるのかを正確に確認できます。
ほとんどのツールは無駄を隠します。在職期間がそれを測定し、それを示します
減少しています。
示されている数字は代表的なセッションの例です。あなたの
ダッシュボードには初日からの実際の数値が表示されます。
コミットする前に試してみてください。
抽出をオンにし、注入をオフにして 1 ～ 2 週間実行します。参照
働き方が変わる前に Tenure が学んだこととまったく同じです
単一の応答。リスクはありません。行動の変化はありません。驚くことはありません。
在職期間はあなたのものを見守ります

静かにセッション。決定事項を抽出し、
好み、事実、質問、阻害要因を構造化した
信念は保存されますが、何も注入されません。 AI の応答は完全に
変わらない。
パネルを開けます。それが何を知っているかを見てください。
インジェクションをオンにする前に、VS Code のサイド パネルを開いて、
在職期間が獲得したすべての信念を読んでください。何でも編集してください。何でも削除してください。
最も重要なことをピン留めします。モデルが何を知るかを決めるのはあなたです。
MCP に依存するメモリ ツールには根本的な欠陥があります。
モデルは彼らに電話をかけるかどうかを決定する必要があります。つまり、あなたの
コンテキストは、モデルがコンテキストに到達しようと考えたときにのみ取得されます。
そうでないこともよくあります。最終的には時々機能する記憶が得られますが、
特定の質問に関して、予想外に。
リクエストに応じて回収します。時々。
メモリがツール呼び出しである場合、いつフェッチするかをモデルが決定します。
コンテキスト。事前情報が必要であると認識しない場合は、
それは問いません。モデルが作成されている間、メモリはアイドル状態になります。
あなたがすでに下した決定を幻覚します。
メモリは自動的に挿入されます。いつも。
Tenure はクライアントとプロバイダーの間にあります。あらゆるリクエスト
モデルに到達する前に強化されます。トリガーするツール呼び出しはありません。
迅速なエンジニアリングは必要ありません。あなたの信念は文脈の中にあります
あらゆるリクエスト。
LLM プロバイダーとして登録されました。
あなたは何も変わりません。
Tenure は、ネイティブ LLM プロバイダーとして VS Code 内に直接登録します。
独自の API キーを持参してください。既存のツール — Copilot、Continue、その他
OpenAI 互換クライアントは、メモリがあれば正常に動作します。
VSコード
ネイティブ LLM プロバイダー · BYOK
誰かをラップするプラグインではなく、プロバイダーレベルで登録されています
他人の電話。在職期間は連鎖の中にあります。あなたのキー、あなたのモデル、あなたの
記憶。オンボーディングは IDE 内で行われます。
どこでも考えてください。それはあなたを追いかけます。
チャット クライアント · モバイル · ゼロ コールド スタート
あなたは散歩中です。認証フローがクリックされます。 OpenClaw に次のように伝えます
ワッツアップ:

「Postgres ではなくセッション用の Redis。TTL ベースの有効期限。」
明日の朝、IDE はすでにそれを知っています。コピー＆ペーストはできません。いいえ
再説明。言った瞬間に決断が下りました。
あなたの信念。サイドバーにあります。
Tenure サイドパネルには、モデルが知っていることを正確に表示します。
ファイルとグローバルに、VS Code 内で直接。信念をクリックして確認してください
毎ターン、それが挿入され、それを表面化した正確なクエリ。
● 信念をクリックすると、それが注入されたすべてのターンが表示されます。 ●
それを表面化した正確なクエリは、それとともに永続化されます。 ●
パネルから直接編集、固定、または削除を行います。 ● ファイルごとおよび
グローバルな信念は範囲内に留まります。プロジェクト A がプロジェクト B に影響することはありません。
何か問題が発生しました。
その理由を調べる方法は次のとおりです。
他のすべてのシステムでは、信念のリストが表示されます。在職期間が示すのは、
歴史。ターンをクリックすると、何が注入されたかを正確に確認できます。信念をクリックして、
それがアクティブだったすべてのターンを確認してください。それをトリガーしたクエリは永続化されます
ログから再構成されたものではなく、発生当時のものです。
ターンごとにどの信念が文脈に含まれていたかを正確に確認します。そうではない
推測される。起こったことをそのまま記録しました。
信念をクリックすると、それが注入されたすべてのセッションが表示されます。
毎回浮上するクエリ。
モデルに指定した正確なクエリは、
注射イベント。記録は完了です。
ターンごとの噴射を追跡するシステムは他にありません。
すべての競合他社は、あなたの信念を静的なリストとして示します。在職期間が与えるもの
あなたは完全な紙の痕跡 — どの信念がいつ存在したか、
どのクエリに対して。なぜモデルなのかを知りたい開発者向け
その違いがすべてです。
必要なときにコントロールします。
そうでないときは見えません。
在職期間は自動的に学習されます。しかし、それが知っているものはすべて目に見えます、
編集可能、修正可能です。
生のチャット履歴ではありません。型付けされ、スコープが設定され、バージョン管理された結論
モデルは指示に従って動作できます

から再導出することなく、
転写。
エンジニアリングの信念はエンジニアリング セッションに残ります。プロジェクトAは決してない
プロジェクト B に流れ込みます。すべての範囲が厳格な構造境界です。
確率フィルターではありません。
決定を変更すると、古い決定は削除されずに廃止されます。
再度提案されることはありませんが、記録は監査のために残ります。
種類
!抽出してください
どのセッションでも。あなたの既存の記憶はまだ機能します。在職期間だけ
メモを取るのをやめる。
開く
/信念
または、VS Code サイド パネルを使用して、Tenure が知っているすべてを確認します。編集
何でも。重要なものをピン留めします。間違っているものは削除してください。
ハードスコープフィルタリングを備えたエイリアス重み付けBM25。埋め込みモデルはありません。
リランキングパスも待ち時間もありません。店舗になるほど精度が向上
劣化するのではなく、成長します。
あなたとプロバイダーの間に位置します。
クライアントに localhost:5757 を指定します。プロバイダーへのテニュアルート、
会話から学び、あらゆる場面で知っていることを注入します。
リクエスト。
ワンクリック。それが唯一の手動ステップです。
VS Code、Windsurf、またはContinueから直接インストールします。または、
スクリプトをインストールします。ローカルの Docker コンテナとして実行されます。あなたの持ち手
トークンは最後に出力されます。 30秒かかります。
拡張機能はポート 5757 でローカル デーモンを安全に起動します。
そしてプロキシ層を接続します。デーモンが自動的に起動します。いいえ
端末が必要です。
キーは IDE のシークレット ストアに保存されます。設定ファイルではありません。
Tenure はローカルベアラートークンを生成し、自動的に
それを IDE のネイティブ シークレット ストアに直接挿入します。ゼロ
プレーンテキストのキーがグローバル設定内に浮遊したままになる
ファイル。
在職期間はバックグラウンドで学習します。最初のセッションはいいですね。の
10番目の方が明らかに優れています。 50番目までに、それはあなたを知っています。毎
セッションコンパウンド。ストアの精度は低下するのではなく、より正確になります。
成長します。
データがマシンから離れることはありません。
雲はありません。アカウントがありません。テレメトリーはありません。在職期間はずっと続きます

頼りになる
ローカルホスト。信念は保存時に暗号化されます。全体をエクスポートする
メモリを暗号化されたアーカイブとして保存し、どこにでも復元できます。
あなたはあなたの記憶を所有しています。全部。永遠に。
インストールには 30 秒かかります。最初のセッションはすでに良くなりました。ターン
準備ができたら、事前ではなく注射を行ってください。
// 既存のツールからインストールします
ターミナルの方がいいですか？
カール -fsSL
https://raw.githubusercontent.com/tenurehq/tenure/main/scripts/install.sh
|バッシュ
コピー
Docker・MITライセンスが必要です
· 完全にローカル · アカウントなし
必要な

## Original Extract

Tenure is a fully local AI memory layer that runs on localhost. Persistent, structured memory that follows you across every tool and every session. 1.0 retrieval precision. Zero cloud. Zero workflow changes.

tenure
Use Cases ▼ Mobile VSCode VSCodium Docs GitHub Writing
Research ▼ Benchmark Paper Install Free Mobile VSCode VSCodium Docs GitHub Writing Benchmark Paper Install Free
Runs entirely on localhost. Nothing ever leaves your machine.
Brainstorm in your chat client. Open your IDE. It already knows.
Not RAG. Not another memory wrapper. Watch what it learns before it
ever changes a response.
Persistent memory that follows you across every tool and every
session. Turn injection on when you're ready, not before.
// install from your existing tools
Prefer the terminal?
curl -fsSL
https://raw.githubusercontent.com/tenurehq/tenure/main/scripts/install.sh
| bash
Copy
1.0 precision. Because a smart model shouldn't have to clean up messy
retrieval.
Other memory tools cheat. They dump your entire chat history or loose
vector clusters into the context window and let a capable downstream
model find the needle in the haystack. They claim SOTA on answer
accuracy but you pay the latency and context tax.
It's the difference between querying your database and filtering in
application code. Everyone knows which one belongs in production.
Every irrelevant belief in context is
tokens you're paying for and
latency you're waiting on.
Results are reproducible. Dataset on HuggingFace.
Run it yourself
You can see exactly what it's costing you.
Most tools hide the waste. Tenure measures it and shows you it
decreasing.
Numbers shown are illustrative from a representative session. Your
dashboard shows your actual numbers from day one.
Try it before you commit to it.
Run with extraction on and injection off for a week or two. See
exactly what Tenure learned about how you work before it ever changes
a single response. No risk. No behavior change. No surprises.
Tenure watches your sessions silently. It extracts decisions,
preferences, facts, questions, and blockers into a structured
belief store but injects nothing. Your AI responses are completely
unchanged.
Open the panel. See what it knows.
Before you ever turn injection on, open the VS Code side panel and
read every belief Tenure captured. Edit anything. Delete anything.
Pin what matters most. You decide what the model gets to know.
Memory tools that rely on MCP have a fundamental flaw.
The model has to decide to call them. That means your
context is only retrieved when the model thinks to reach for it and it
often does not. You end up with memory that works sometimes,
unpredictably, on certain questions.
Retrieval on request. Sometimes.
When memory is a tool call, the model decides when to fetch
context. If it does not recognize that it needs prior information,
it does not ask. Your memory sits idle while the model
hallucinates decisions you already made.
Memory injected automatically. Always.
Tenure sits between your client and your provider. Every request
is enriched before it reaches the model. No tool call to trigger.
No prompt engineering required. Your beliefs are in context on
every single request.
Registered as your LLM provider.
You don't change a thing.
Tenure registers directly inside VS Code as a native LLM provider.
Bring your own API key. Your existing tools — Copilot, Continue, any
OpenAI-compatible client just works, with memory.
VS Code
Native LLM Provider · BYOK
Registered at the provider level, not a plugin wrapping someone
else's call. Tenure is in the chain. Your key, your model, your
memory. Onboarding happens inside the IDE.
Think anywhere. It follows you.
Chat client · Mobile · Zero cold starts
You're on a walk. The auth flow clicks. You tell OpenClaw on
WhatsApp: "Redis for sessions, not Postgres. TTL-based expiry."
Tomorrow morning, your IDE already knows. No copy-paste. No
re-explanation. The decision landed the moment you said it.
Your beliefs. Right in the sidebar.
The Tenure side panel shows you exactly what the model knows, per
file and globally, directly inside VS Code. Click any belief to see
every turn it was injected and the exact query that surfaced it.
● Click any belief to see every turn it was injected. ●
The exact query that surfaced it is persisted alongside it. ●
Edit, pin, or delete directly from the panel. ● Per-file and
global beliefs stay scoped — project A never bleeds into project B.
Something went wrong.
Here's how you find out why.
Every other system shows you a list of beliefs. Tenure shows you a
history. Click a turn, see exactly what was injected. Click a belief,
see every turn it was active. The query that triggered it is persisted
at the time it happened, not reconstructed from logs.
See exactly which beliefs were in context for every turn. Not
inferred. Recorded as it happened.
Click any belief to see every session it was injected, and the
query that surfaced it each time.
The exact query you gave the model is stored alongside the
injection event. The record is complete.
No other system tracks injection per turn.
Every competitor shows you beliefs as a static list. Tenure gives
you a complete paper trail — which beliefs were present, when,
against which query. For developers who care about why the model
responded the way it did, that difference is everything.
Control when you want it.
Invisible when you do not.
Tenure learns automatically. But everything it knows is visible,
editable, and correctable.
Not raw chat history. Typed, scoped, versioned conclusions the
model can act on directly without re-deriving them from a
transcript.
Engineering beliefs stay in engineering sessions. Project A never
bleeds into Project B. Every scope is a hard structural boundary
not a probabilistic filter.
When you change a decision, the old one is retired — not deleted.
It never gets suggested again, but the record stays for audit.
Type
!extract off
in any session. Your existing memory still works. Tenure just
stops taking notes.
Open
/beliefs
or use the VS Code side panel to see everything Tenure knows. Edit
anything. Pin what matters. Delete what is wrong.
Alias-weighted BM25 with hard scope filtering. No embedding model,
no reranking pass, no waiting. Precision improves as the store
grows — not degrades.
Sits between you and your provider.
Point your client at localhost:5757. Tenure routes to your provider,
learns from the conversation, and injects what it knows on every
request.
One click. That's the only manual step.
Install directly from VS Code, Windsurf, or Continue. Or run the
install script. Runs as a local Docker container. Your bearer
token is printed at the end. Takes thirty seconds.
The extension securely spins up the local daemon on port 5757
and hooks up the proxy layer. The daemon starts itself. No
terminal required.
Your key goes into the IDE's secret store. Not a config file.
Tenure generates your local bearer token and automatically
injects it directly into your IDE's native secret store. Zero
plain-text keys left floating around in your global config
files.
Tenure learns in the background. The first session is good. The
tenth is noticeably better. By the fiftieth, it knows you. Every
session compounds. The store gets more precise, not less, as it
grows.
Your data never leaves your machine.
No cloud. No accounts. No telemetry. Tenure runs entirely on
localhost. Beliefs are encrypted at rest. Export your entire
memory as an encrypted archive and restore it anywhere.
You own your memory. All of it. Forever.
Thirty seconds to install. First session already better. Turn
injection on when you are ready, not before.
// install from your existing tools
Prefer the terminal?
curl -fsSL
https://raw.githubusercontent.com/tenurehq/tenure/main/scripts/install.sh
| bash
Copy
Requires Docker · MIT licensed
· Fully local · No account
needed
