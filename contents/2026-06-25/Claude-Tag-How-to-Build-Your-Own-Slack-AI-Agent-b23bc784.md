---
source: "https://www.arcade.dev/blog/claude-tag-build-slack-ai-agent/"
hn_url: "https://news.ycombinator.com/item?id=48677297"
title: "Claude Tag: How to Build Your Own Slack AI Agent"
article_title: "Claude Tag: Build Your Own Slack AI Agent"
author: "manveerc"
captured_at: "2026-06-25T18:42:41Z"
capture_tool: "hn-digest"
hn_id: 48677297
score: 1
comments: 0
posted_at: "2026-06-25T18:18:57Z"
tags:
  - hacker-news
  - translated
---

# Claude Tag: How to Build Your Own Slack AI Agent

- HN: [48677297](https://news.ycombinator.com/item?id=48677297)
- Source: [www.arcade.dev](https://www.arcade.dev/blog/claude-tag-build-slack-ai-agent/)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T18:18:57Z

## Translation

タイトル: Claude タグ: 独自の Slack AI エージェントを構築する方法
記事のタイトル: Claude タグ: 独自の Slack AI エージェントを構築する
説明: Anthropic は、Slack 用の共有 AI エージェントである Claude Tag を開始しました。その仕組みを学び、Python、OpenAI、Arcade を使用して Claude Tag スタイルのエージェントを構築します。

記事本文:
Claude Tag: 独自の Slack AI エージェントを構築する
製品価格 構築を開始 カタログを参照 ビルド ツール パターン MCP フレームワーク LangChain フリート リソース ドキュメント ブログ ホワイト ペーパー オンサイト ワークショップ キャリア サインアップ ダッシュボード デモを予約 製品価格 構築を開始 カタログを参照 ビルド ツール パターン MCP フレームワーク LangChain フリート リソース ドキュメント ブログ ホワイト ペーパー オンサイト ワークショップ キャリア サインアップ ダッシュボード デモを予約 クロード タグ: Arcade.dev で独自の Slack AI エージェントを構築する方法
「現在、当社の製品チームのコードの 65% は、内部バージョンの Claude Tag によって作成されています。」
それが Anthropic であり、自社のエンジニアリング チームについて話しています。これは、コードのオートコンプリートやスニペットを個別に生成するチャットボットではありません。 Claude Tag は Slack 内の共有エージェントであり、チームメイトが名前を挙げてバグを調査し、メトリクスを取得し、サポート チケットを作業し、長時間実行されるタスクを完了します。スレッドのコンテキストを読み取り、承認されたツールとコードベースに接続し、同じ会話に結果をポストします。
問題は、クロード・タグが印象的かどうかではありません。それは、あなたのチームに委任者がいるとしたら何を委任するでしょうか?
それを知るために Anthropic の製品全体を再作成する必要はありません。このチュートリアルは、Anthropic の独自製品ではなく、Claude Tag のコア インタラクション パターンを再作成します。価値の高い 1 つの Slack ワークフローから開始し、エージェントに小さなツールセットを提供し、ツールの接続、承認、外部システムへのアクセス制御などのアクション レイヤーに Arcade.dev を使用します。
重要なポイント: Claude Tag と独自の Slack AI エージェントの構築
Claude Tag は、Anthropic の Slack 用共有 AI エージェントです。これにより、チームは選択したチャネルで @Claude にメンションし、会話コンテキスト、接続されたツール、コードベースを使用して複数のステップの作業を完了できます。
Claude Tag は Slack をエージェント インターフェイスに変えます。できる

関連するチャネル コンテキストを記憶し、非同期で動作し、専用の ID を使用して、リクエストが開始されたスレッドで結果を返します。
コアとなるクロード タグ パターンを再作成できます。このチュートリアルでは、Python、Slack Bolt、OpenAI、Arcade を使用して Claude Tag スタイルの Slack AI エージェントを構築します。
Arcade は安全なツールへのアクセスを提供します。この例では、エージェントを読み取り専用の GitHub、Datadog、PagerDuty ツールに接続し、Arcade が承認、資格情報、ツールの実行、およびアクセス制御を処理します。
1 つの限定されたワークフローから始めます。インシデントのトリアージは、複数のシステムを横断し、レビュー可能な証拠を生成し、取り消し可能なアクションを必要としないため、最初の使用例としては強力です。
実稼働エージェントには明示的な保護手段が必要です。エージェントを承認された Slack チャネルに制限し、専用の ID またはユーザーごとの ID を使用し、結果的な書き込みには人間の承認を必要とし、アクションをログに記録し、キル スイッチを維持します。
クロード タグとは何ですか?なぜあなたのチームはそれを望んでいますか?
Anthropic は、2026 年 6 月 23 日に Enterprise および Team 顧客向けのベータ版として Claude Tag を開始しました。運用モデルはシンプルです。クロードは、選択した Slack チャネルにチームメイトとして参加します。チャンネル内の誰もがリクエストで @Claude をタグ付けできます。タスクを段階に分割し、接続されたツールを使用してそれらの段階を処理し、生成された内容をスレッド内で返信します。スレッドがアクティブになると、そこにいる誰もがエージェントを再度言及することなくスレッドを操作できます。
これが個人用チャットボットと異なるのは、作業が公開で行われることです。チャネルはインターフェイス、コンテキスト、監査証跡です。単一の共有 Claude インスタンスがチャネル全体にサービスを提供し、それに応じて永続メモリを構築します。非同期で動作し、独自のフォローアップ タスクをスケジュールし、Slack スレッド、Google ドライブ ドキュメント、チケット発行システム、データ ウェアハウスからのコンテキストを組み合わせることができます。

たった一つの答えに。
根底にある洞察は AI の機能に関するものではありません。それは仕事がどこから始まるかについてです。ほとんどの部門横断的なタスクは Slack メッセージとして始まります。誰かが質問したり、問題にフラグを立てたり、3 つのシステムにまたがる情報を要求したりします。共有エージェントの真の価値は、作業がすでに始まっている場所で有用な作業を実行できる場合にあります。
AI 従業員を育成しないでください。ワークフローを 1 つ選択します。
エージェント プロジェクトを遅らせる最も早い方法は、エージェント プロジェクトを「何でもできる AI」としてスコープすることです。 1 つのワークフローから始めます。次のようなものを選択してください。
頻繁に。チームはそれを毎週、理想的には毎日行っています。
クロスシステム。 2 つ以上のツール (Slack、GitHub、ダッシュボード、CRM) からコンテキストを取得する必要があります。
手動で調査するのは面倒です。誰かがタブ間でコピー＆ペーストし、調査結果を要約し、更新を投稿する必要があります。
人間にとってレビューしやすい。エージェントは、最終的な取り消し不能なアクションではなく、概要または推奨事項を作成します。
価値の高い出発点をいくつか示します。
Slack、GitHub、および可観測性ツールにわたるインシデントのトリアージ。デプロイ後にエラーが急増すると、エージェントは最近のコミットを取得し、Datadog にエラー率とレイテンシをクエリし、PagerDuty で関連インシデントをチェックし、証拠リンクを含む構造化された概要を投稿します。
チケット発行システム、CRM、社内ドキュメントを使用してエスカレーションの概要をサポートします。エンジニアがエスカレーションされたチケットのコンテキストの再構築に 15 分を費やす代わりに、エージェントが数秒でそれを実行し、エスカレーション チャネルに概要を投稿します。
Slack スレッドを読み取り、コア リクエストを抽出し、Linear または Jira で重複をチェックし、リンクされた元のスレッドで適切にタグ付けされた課題を作成する製品フィードバックのトリアージ。
CRM データ、最近の電子メール スレッド、製品使用状況の指標、保管前の内部メモをまとめたアカウント調査

マーコール。
狭く始めてください。集中力のあるエージェントは、幅広い能力を持つエージェントよりも早く信頼を獲得します。
Claude Tag スタイルの Slack エージェントはどのように機能しますか?
Claude Tag スタイルのエージェントの背後にあるアーキテクチャには 4 つの層があります。
Slack はインターフェースです。ユーザーはスレッド内のエージェントにタグを付けます。 Slack はトリガーとなるイベントを配信します。アプリケーションは API 経由でスレッド コンテキストを取得し、結果を表示します。
モデルは推論層です。リクエストを理解し、どのような情報が必要かを判断し、応答を合成します。スタックに適合する LLM およびエージェント フレームワークを使用してください。
アーケードはアクション層です。エージェントを承認されたツールに接続し、認可とトークン管理を処理し、アクセス ポリシーを適用します。モデルは資格情報を決して参照しません。
アプリはオーケストレーションを処理します。タスクの状態、再試行、非同期ジョブ処理、および Slack への更新の投稿。
各層は独立して交換可能です。モデルを交換し、フレームワークを変更し、ツールを追加します。境界線はきれいなままです。
私たちが構築しているのは共有エージェントであり、マルチユーザー エージェントではありません。すべてのツール呼び出しは、誰がボットにタグを付けたかに関係なく、単一のサービス ID の下で実行されます。ステップ 4 では、ユースケースで必要な場合にユーザーごとの承認を追加する方法について説明します。
このプロトタイプは、言及された場合にのみ実行を開始します。 Claude Tag の制作経験は、アクティブなスレッド内での言及されていないフォローアップをサポートします。この動作を追加するには、 message.channels と message.groups をサブスクライブし、アクティブなスレッド ID を追跡し、ボットによって生成されたメッセージをフィルターで除外します。これは、このウォークスルーの範囲を超えた運用上の拡張機能です。
Arcade を使用して Claude Tag スタイルの Slack エージェントを構築する方法
このチュートリアルでは、Python と Slack の Bolt フレームワークおよび Arcade Python SDK を使用します。同じパターンは、MCP または Arcade の REST API をサポートする任意の言語またはエージェント フレームワークで機能します。
Python 3.8 以降が必要です、許可します

Slack アプリ、Arcade アカウントと API キー、OpenAI API キーを作成してインストールします。ローカルの Slack イベント API テストの場合は、ngrok CLI または別のパブリック HTTPS トンネルもインストールして認証します。
python3 -m venv .venv
ソース .venv/bin/activate
python -m pip install slack-bolt archivepy openai
ステップ 1: Slack アプリとイベント トリガーを作成する
api.slack.com/apps で Slack アプリを作成します。 [OAuth & Permissions] で、ボット スコープ app_mentions:read 、 chat:write 、 Channels:history 、および groups:history を追加します。アプリをワークスペースにインストールし、アプリ設定からボット ユーザー OAuth トークン ( xoxb-... ) と署名シークレットをコピーします。
これで、環境変数を設定するために必要なものがすべて揃いました。
import SLACK_BOT_TOKEN = "xoxb-..."
エクスポート SLACK_SIGNING_SECRET = "..."
エクスポート ARCADE_API_KEY = "..."
ARCADE_USER_ID = "you@company.com" をエクスポート
エクスポート OPENAI_API_KEY = "..."
エクスポート SLACK_ALLOWED_CHANNEL_IDS = "C0123456789"
ARCADE_USER_ID には、Arcade アカウントに関連付けられた電子メールを使用します。 Arcade のデフォルトの開発検証ツールはその ID を期待します。これは、すべてのツール呼び出しが実行される単一の共有 ID です。承認されたすべてのチャネルのすべてのメンションは、この 1 つのアカウントに解決されます。 GitHub または PagerDuty サービス アカウントを独自に作成することはありません。エージェントが専用のダウンストリーム ID で動作する必要がある場合は、ステップ 2 の OAuth フロー中に専用アカウントを使用します。
C0123456789 を実際の Slack チャネル ID に置き換えます。 Slack の Web アプリまたはデスクトップ アプリでチャネルを開き、その URL ( https://app.slack.com/client/T.../C... ) の C... 部分をコピーします。詳細については、ID を見つけるための Slack のガイドを参照してください。
SLACK_ALLOWED_CHANNEL_IDS は、エージェントを特定のチャネルに制限し、Claude Tag が使用するチャネルごとのスコープを強制します。カンマで区切られた複数のチャネル ID。違ったらチャ

nnels には異なる権限またはツールセットが必要なので、channel_id から ID へのマッピングまたは個別のデプロイメントが必要になります。
Slack の 3 秒ルールは、実装の重要な詳細です。エンドポイントは 3 秒以内に HTTP 200 を返す必要があり、そうでない場合、Slack は配信を失敗としてマークし、最大 3 回再試行します。標準のデコレータ パターンを使用する場合、Bolt は確認応答を自動的に処理します。エージェントの処理に時間がかかる運用ワークロードの場合は、作業をタスク キューにオフロードします。作業をキューに入れる前に、Slack のトップレベルのevent_id で重複を排除します。そうしないと、再試行で同じツールが 2 回実行される可能性があります。
インポートログ
OSをインポートする
lack_boltインポートアプリから
from Agent import run_agent # ステップ 3
ALLOWED_CHANNEL_IDS = {
値。ストリップ（）
os.environ[ "SLACK_ALLOWED_CHANNEL_IDS" ] の値。分割 ( "," )
値の場合。ストリップ（）
}
アプリ = アプリ (
トークン = os.environ[ "SLACK_BOT_TOKEN" ],
signed_secret = os.environ[ "SLACK_SIGNING_SECRET" ],
）
@app 。イベント ( "app_mention" )
def handle_mention (event 、 client 、say 、 context 、 logger ):
イベント[ "channel" ] が ALLOWED_CHANNEL_IDS にない場合:
ロガー。警告 ( "承認されていないチャネル %s からのメンションを無視します" 、イベント[ "チャネル" ])
戻る
# ループを防ぐためにボット (このメッセージを含む) からのメッセージを無視します
イベントなら。 get ( "bot_id" ):
戻る
thread_ts = イベント。 get ( "thread_ts" ) またはevent[ "ts" ]
試してみてください:
# スレッド コンテキストのメッセージを最大 50 件取得します。
# 実稼働実装は以下に従う必要があります
# 長いスレッドの場合は、response_metadata.next_cursor。
返信 = クライアント。会話_返信 (
チャンネル = イベント[ "チャンネル" ],
ts = thread_ts、
制限 = 50 、
）
bot_user_id = コンテキスト。 get ( "bot_user_id" )
トランスクリプト = []
返信のメッセージ用。 get ( "メッセージ" , []):
テキスト = メッセージ。 get ( "テキスト" , "" )
bot_user_id の場合:
テキスト＝テキスト。 replace ( f "<@ { bot_user_id } >" , "

" ).ストリップ()
テキストの場合:
スピーカー = メッセージ。 get ( "user" ) またはメッセージ。 get ( "bot_id" , "unknown" )
転写。 append ( f " { スピーカー } : { テキスト } " )
Say ( "準備中です。コンテキストを収集中..." , thread_ts = thread_ts)
結果 = run_agent (
os.environ[ "ARCADE_USER_ID" ],
" \n " 。参加（転写）、
）
# Slack ではメッセージを 4,000 文字以下にすることを推奨しています。
# 運用環境で長い応答を切り詰めるか、チャンクします。
言う (結果、thread_ts = thread_ts)
例外を除く:
ロガー。例外 (「エージェントが失敗しました」)
言う（
「その調査を完了できませんでした。アプリケーション ログを確認してください。」 、
スレッド_ts = スレッド_ts、
）
if __name__ == "__main__" :
ロギング。 BasicConfig (レベル = ロギング。情報)
# これは、Bolt の組み込み開発サーバーです。生産にあたっては、
# サポートされている Web フレームワーク アダプターを介してデプロイします (例: Flask + Gunic)
[切り捨てられた]
ステップ 2: GitHub、Datadog、PagerDuty を Arcade に接続する
Arcade は、厳選されたツール セットを通じてエージェントを外部システムに接続します。インシデントのトリアージには、GitHub、Datadog、PagerDuty の読み取り専用ツールが必要です。ツールキット全体をロードするのではなく、特定のツールを選択します。ツールキットには、読み取り専用エージェントのスコープと矛盾する書き込み操作が含まれており、ツール リストが狭いため、モデルはより信頼性の高い適切なツールを選択することができます。

[切り捨てられた]

## Original Extract

Anthropic launched Claude Tag, a shared AI agent for Slack. Learn how it works and build a Claude Tag-style agent with Python, OpenAI, and Arcade.

Claude Tag: Build Your Own Slack AI Agent
Product Pricing Start Building Browse the Catalog Build Tool Patterns MCP Framework LangChain Fleet Resources Docs Blog White Papers Onsite Workshop Careers Sign up Dashboard Book a demo Product Pricing Start Building Browse the Catalog Build Tool Patterns MCP Framework LangChain Fleet Resources Docs Blog White Papers Onsite Workshop Careers Sign up Dashboard Book a demo Claude Tag: How to Build Your Own Slack AI Agent with Arcade.dev
“Today, 65% of our product team’s code is created by our internal version of Claude Tag.”
That’s Anthropic, talking about its own engineering team. And this is not code autocomplete or a chatbot generating snippets in isolation. Claude Tag is a shared agent inside Slack that teammates mention by name to investigate bugs, pull metrics, work support tickets, and complete longer-running tasks. It reads thread context, connects to approved tools and codebases, and posts results back in the same conversation.
The question is not whether Claude Tag is impressive. It is: what would your team delegate if you had one?
You do not need to recreate Anthropic’s entire product to find out. This tutorial recreates Claude Tag’s core interaction pattern, not Anthropic’s proprietary product. Start with one high-value Slack workflow, give the agent a small toolset, and use Arcade.dev for the action layer: tool connectivity, authorization, and controlled access to external systems.
Key takeaways: Claude Tag and building your own Slack AI agent
Claude Tag is Anthropic’s shared AI agent for Slack . It lets teams mention @Claude in selected channels to complete multi-step work using conversation context, connected tools, and codebases.
Claude Tag turns Slack into the agent interface . It can remember relevant channel context, work asynchronously, use a dedicated identity, and return results in the thread where the request began.
You can recreate the core Claude Tag pattern. This tutorial builds a Claude Tag-style Slack AI agent with Python, Slack Bolt, OpenAI, and Arcade.
Arcade provides secure tool access. The example connects the agent to read-only GitHub, Datadog, and PagerDuty tools while Arcade handles authorization, credentials, tool execution, and access controls.
Start with one bounded workflow . Incident triage is a strong first use case because it crosses multiple systems, produces reviewable evidence, and does not require irreversible actions.
Production agents need explicit safeguards. Restrict the agent to approved Slack channels, use dedicated or per-user identities, require human approval for consequential writes, log its actions, and maintain a kill switch.
What is Claude Tag and why does your team want it?
Anthropic launched Claude Tag on June 23, 2026 as a beta for Enterprise and Team customers. The operating model is simple: Claude joins selected Slack channels as a teammate. Anyone in the channel can tag @Claude with a request. It breaks the task into stages, works through them using connected tools, and replies in-thread with what it produced. Once a thread is active, anyone there can steer it without re-mentioning the agent.
What makes this different from a personal chatbot is that the work happens in public. The channel is the interface, the context, and the audit trail. A single shared Claude instance serves an entire channel, building persistent memory as it follows along. It can work asynchronously, schedule its own follow-up tasks, and combine context from Slack threads, Google Drive docs, ticketing systems, and data warehouses into a single answer.
The underlying insight is not about AI capabilities. It is about where work starts. Most cross-functional tasks begin as a Slack message. Someone asks a question, flags a problem, or requests information that lives across three systems. The true value of shared agents is when it can do useful work in a place where that work already begins.
Do not build an AI employee. Pick one workflow.
The fastest way to stall an agent project is to scope it as “an AI that can do anything.” Start with one workflow. Choose something that is:
Frequent. The team does it every week, ideally every day.
Cross-system. It requires pulling context from two or more tools (Slack, GitHub, a dashboard, a CRM).
Tedious to investigate manually. Someone has to copy-paste between tabs, summarize findings, and post an update.
Easy for a human to review. The agent produces a summary or recommendation, not a final irreversible action.
Some high-value starting points:
Incident triage across Slack, GitHub, and observability tools. When errors spike after a deployment, the agent pulls recent commits, queries Datadog for error rates and latency, checks PagerDuty for related incidents, and posts a structured summary with evidence links.
Support escalation summaries using your ticketing system, CRM, and internal docs. Instead of an engineer spending 15 minutes rebuilding context on an escalated ticket, the agent does it in seconds and posts the summary in the escalation channel.
Product-feedback triage that reads a Slack thread, extracts the core request, checks for duplicates in Linear or Jira, and creates a properly tagged issue with the original thread linked.
Account research that pulls together CRM data, recent email threads, product usage metrics, and internal notes before a customer call.
Start narrow. A focused agent earns trust faster than a broadly capable one.
How does a Claude Tag-style Slack agent work?
The architecture behind a Claude Tag-style agent has four layers:
Slack is the interface. Users tag the agent in a thread. Slack delivers the triggering event; your application retrieves thread context via the API and displays results.
The model is the reasoning layer. It understands the request, decides what information it needs, and synthesizes a response. Use whatever LLM and agent framework fits your stack.
Arcade is the action layer. It connects the agent to approved tools, handles authorization and token management, and enforces access policy. The model never sees credentials.
Your app handles orchestration. Task state, retries, async job processing, and posting updates back to Slack.
Each layer is independently replaceable. Swap the model, change the framework, add tools. The boundaries stay clean.
What we are building is a shared agent, not a multi-user agent. Every tool call runs under a single service identity regardless of who tagged the bot. Step 4 covers how to add per-user authorization if your use case requires it.
This prototype starts a run only when mentioned. Claude Tag’s production experience supports unmentioned follow-ups within an active thread. To add that behavior, subscribe to message.channels and message.groups , track active thread IDs, and filter out bot-generated messages. That is a production extension beyond the scope of this walkthrough.
How to build a Claude Tag-style Slack agent with Arcade
This walkthrough uses Python with Slack’s Bolt framework and the Arcade Python SDK. The same pattern works with any language or agent framework that supports MCP or Arcade’s REST API.
You need Python 3.8+, permission to create and install a Slack app, an Arcade account and API key , and an OpenAI API key . For local Slack Events API testing, also install and authenticate the ngrok CLI or another public HTTPS tunnel.
python3 -m venv .venv
source .venv/bin/activate
python -m pip install slack-bolt arcadepy openai
Step 1: Create the Slack app and event trigger
Create a Slack app at api.slack.com/apps . Under OAuth & Permissions , add the bot scopes app_mentions:read , chat:write , channels:history , and groups:history . Install the app to your workspace, then copy the Bot User OAuth Token ( xoxb-... ) and Signing Secret from the app settings.
You now have everything needed to set the environment variables:
export SLACK_BOT_TOKEN = "xoxb-..."
export SLACK_SIGNING_SECRET = "..."
export ARCADE_API_KEY = "..."
export ARCADE_USER_ID = "you@company.com"
export OPENAI_API_KEY = "..."
export SLACK_ALLOWED_CHANNEL_IDS = "C0123456789"
For ARCADE_USER_ID , use the email associated with your Arcade account. Arcade’s default development verifier expects that identity. This is the single shared identity under which every tool call executes. All mentions in all approved channels resolve to this one account. It does not create GitHub or PagerDuty service accounts on its own. If the agent must act under a dedicated downstream identity, use dedicated accounts during the OAuth flows in Step 2.
Replace C0123456789 with your actual Slack channel ID. Open the channel in Slack’s web or desktop app and copy the C... portion of its URL ( https://app.slack.com/client/T.../C... ). See Slack’s guide to locating IDs for details.
SLACK_ALLOWED_CHANNEL_IDS restricts the agent to specific channels, enforcing the per-channel scoping that Claude Tag uses. Comma-separate multiple channel IDs. If different channels need different permissions or toolsets, you will need a channel_id -to-identity mapping or separate deployments.
Slack’s three-second rule is the critical implementation detail. Your endpoint must return HTTP 200 within three seconds or Slack marks delivery as failed and retries up to three times. Bolt handles acknowledgement automatically when you use the standard decorator pattern. For production workloads where agent processing takes longer, offload work to a task queue. Deduplicate on Slack’s top-level event_id before enqueueing work, otherwise retries can execute the same tools twice.
import logging
import os
from slack_bolt import App
from agent import run_agent # Step 3
ALLOWED_CHANNEL_IDS = {
value. strip ()
for value in os.environ[ "SLACK_ALLOWED_CHANNEL_IDS" ]. split ( "," )
if value. strip ()
}
app = App (
token = os.environ[ "SLACK_BOT_TOKEN" ],
signing_secret = os.environ[ "SLACK_SIGNING_SECRET" ],
)
@app . event ( "app_mention" )
def handle_mention ( event , client , say , context , logger ):
if event[ "channel" ] not in ALLOWED_CHANNEL_IDS :
logger. warning ( "Ignoring mention from unauthorized channel %s " , event[ "channel" ])
return
# Ignore messages from bots (including this one) to prevent loops
if event. get ( "bot_id" ):
return
thread_ts = event. get ( "thread_ts" ) or event[ "ts" ]
try :
# Retrieve up to 50 messages of thread context.
# Production implementations should follow
# response_metadata.next_cursor for longer threads.
replies = client. conversations_replies (
channel = event[ "channel" ],
ts = thread_ts,
limit = 50 ,
)
bot_user_id = context. get ( "bot_user_id" )
transcript = []
for message in replies. get ( "messages" , []):
text = message. get ( "text" , "" )
if bot_user_id:
text = text. replace ( f "<@ { bot_user_id } >" , "" ). strip ()
if text:
speaker = message. get ( "user" ) or message. get ( "bot_id" , "unknown" )
transcript. append ( f " { speaker } : { text } " )
say ( "On it. Gathering context..." , thread_ts = thread_ts)
result = run_agent (
os.environ[ "ARCADE_USER_ID" ],
" \n " . join (transcript),
)
# Slack recommends keeping messages under 4,000 characters.
# Truncate or chunk longer responses in production.
say (result, thread_ts = thread_ts)
except Exception:
logger. exception ( "Agent failed" )
say (
"I couldn't complete that investigation. Check the application logs." ,
thread_ts = thread_ts,
)
if __name__ == "__main__" :
logging. basicConfig ( level = logging. INFO )
# This is Bolt's built-in development server. For production,
# deploy through a supported web-framework adapter (e.g. Flask + Gunic
[truncated]
Step 2: Connect GitHub, Datadog, and PagerDuty with Arcade
Arcade connects your agent to external systems through a curated set of tools. For incident triage, you need read-only tools from GitHub, Datadog, and PagerDuty. Select specific tools rather than loading entire toolkits. Toolkits include write operations that contradict a read-only agent’s scope, and a narrower tool list helps the model pick the right tool more reliabl

[truncated]
