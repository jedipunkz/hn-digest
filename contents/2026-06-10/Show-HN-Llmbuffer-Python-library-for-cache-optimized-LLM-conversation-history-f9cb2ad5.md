---
source: "https://github.com/scottpurdy/llmbuffer"
hn_url: "https://news.ycombinator.com/item?id=48481750"
title: "Show HN: Llmbuffer – Python library for cache-optimized LLM conversation history"
article_title: "GitHub - scottpurdy/llmbuffer: LLM conversation buffer with cache optimization and dynamic context. · GitHub"
author: "scottmp10"
captured_at: "2026-06-10T21:45:24Z"
capture_tool: "hn-digest"
hn_id: 48481750
score: 1
comments: 1
posted_at: "2026-06-10T19:57:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Llmbuffer – Python library for cache-optimized LLM conversation history

- HN: [48481750](https://news.ycombinator.com/item?id=48481750)
- Source: [github.com](https://github.com/scottpurdy/llmbuffer)
- Score: 1
- Comments: 1
- Posted: 2026-06-10T19:57:26Z

## Translation

タイトル: Show HN: Llmbuffer – キャッシュに最適化された LLM 会話履歴用の Python ライブラリ
記事のタイトル: GitHub - scottpurdy/llmbuffer: キャッシュの最適化と動的コンテキストを備えた LLM 会話バッファー。 · GitHub
説明: キャッシュ最適化と動的コンテキストを備えた LLM 会話バッファー。 - スコットパーディ/llmbuffer
HN テキスト: LLMbuffer のリリース - LLM プロンプト キャッシュ ヒットを最大化する Python ライブラリ。柔軟なフックを介して、動的なコンテキスト、圧縮、ツール出力の切り捨てや要約を処理します。通常の使用法では 10 倍のコスト削減が期待できます。以下をサポートします。
- ステートフル (PromptManager) およびステートレス/機能的 (純粋な辞書から辞書へ) インターフェイス
- メッセージを短期メッセージ履歴から長期メッセージ履歴に移動するための 3 つの移行モード: none、manual、agent_cycle。たとえば、agent_cycle モードを使用すると、エージェント ループの終了後にメッセージを長期履歴に移動するときに、ツールの出力を切り捨て、要約、または削除できます。
- 履歴が長くなった場合のプラグイン可能な圧縮フック
- Anthropic (cache_control マーカー)、OpenAI (自動プレフィックス キャッシュ)、Gemini、および HF トークナイザーを介したローカル モデル用のプロバイダー アダプター
- --compare モードを備えたベンチマーク スイートは、llmbuffer と naive を並べて表示し、各プロバイダーの使用状況メタデータから実際のキャッシュ ヒット カウントを読み取ります。私の使用状況では、頻繁に動的コンテキストが変更されるにもかかわらず、90% を超えるトークンがキャッシュされていることがわかります。 MITライセンス。
GitHub: https://github.com/scottpurdy/llmbuffer
PyPI: https://pypi.org/project/llmbuffer/

記事本文:
GitHub - scottpurdy/llmbuffer: キャッシュ最適化と動的コンテキストを備えた LLM 会話バッファー。 · GitHub
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
スコットパーディ
/
llmbuffer
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
タラ

e
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット src/ llmbuffer src/ llmbuffer テスト testing .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイル ナビゲーション
キャッシュに最適化された LLM 会話履歴管理。
ほとんどの LLM アプリケーションは、システム プロンプト、会話履歴、および動的コンテキストを単純に 1 つのメッセージ リストに連結し、毎ターン最初から再構築します。これは機能しますが、プロバイダーのプロンプト キャッシュを常に無効にするため、多額の費用と待ち時間が発生します。
llmbuffer は、キャッシュの再利用を最大化する順序でメッセージを組み立て、安定したコンテンツと変化するコンテンツの境界を管理し、履歴が長くなりすぎた場合の圧縮を処理します。これらすべてをユーザーが考える必要はありません。
[静的システム プロンプト] → [長期履歴] → [動的コンテキスト] → [最近のメッセージ]
キャッシュ済み ✓ キャッシュ済み ✓ キャッシュなし キャッシュなし
静的なシステム プロンプトとコミットされた会話履歴は、ターンをまたいで変更されたり並べ替えられたりすることのないバイト安定プレフィックスを形成します。頻繁に変更される部分 (RAG 結果、タイムスタンプ、実行中のツール呼び出し) は、プレフィックスを無効にできない最後に存在します。
pip インストール llmbuffer
ライブベンチマーク用のオプションの追加機能:
pip install " llmbuffer[anthropic] " # Anthropic プロンプト キャッシュ
pip install " llmbuffer[openai] " # OpenAI プレフィックス キャッシュ
llmbuffer には必要な依存関係はなく、Python 3.9 以降のみです。
llmbuffer インポート PromptManager 、 AnthropicAdapter から
マネージャー = プロンプトマネージャー (
static_system_prompt = "あなたは上級ソフトウェア エンジニアリング アシスタントです..." ,
transition_mode = "agent_cycle" , # 自動コミットは安定したプレフィックスに変わります
アダプター = AnthropicAdapter (), #

cache_control マーカーを挿入する
max_tokens = 8_000 , # これを超える長期にわたる履歴をコンパクトにする
）
# 各ターン:
マネージャー。 append ({ "ロール" : "ユーザー" , "コンテンツ" : user_message })
メッセージ = マネージャー。 build_messages (dynamic_system_prompt = rag_context)
返信 = anthropic_client 。メッセージ。作成 (メッセージ = メッセージ , ...)
マネージャー。 append ({ "役割" : "アシスタント" , "コンテンツ" : 返信 })
ステートレス（Webアプリ/サーバーレス）
JSON でシリアル化可能な状態辞書を介した純粋な関数 — リクエスト間の任意の場所に保持します。
llmbuffer インポート機能、 new_state 、ダンプ、ロードから
SYSTEM = 「あなたは上級ソフトウェア エンジニアリング アシスタントです...」
# DB/セッションから状態をロード
state =loads ( row .conversation_json ) if row else new_state ()
# メッセージを構築し、LLM を呼び出し、更新された状態を保存します
状態 = 機能的。 append_message ( state , { "role" : "user" , "content" : text },
遷移モード = "手動")
メッセージ = 機能的。 build_messages ( state 、 static_system_prompt = SYSTEM 、
動的システムプロンプト = rag_context )
# ... LLM に電話します ...
状態 = 機能的。 append_message (state、reply、transition_mode = "manual")
状態 = 機能的。 Compact ( state , max_tokens = 8_000 ) # 関数 API で明示的
行。会話_json = ダンプ (状態)
各関数は、使用する設定のみを受け取ります。スレッドスルーする構成オブジェクトはありません。圧縮は、関数 API での明示的な Compact() 呼び出しです (ステートフルな PromptManager が自動的に実行します)。
build_messages() は常に次の正確な順序でメッセージを出力します。
メッセージが短期的なメッセージから安定した長期的な履歴に移行するタイミングを制御します。
メッセージが短期的なものから長期的な (キャッシュされた) 履歴に移動する前に、オプションのtransition_hook でメッセージを書き換えることができます。ツールの詳細な出力をトリミングしたり、作成したコンテンツを削除したりするのに役立ちます。

安定したプレフィックスに永久にロックされたくないからです。
def Trim_tool_outputs (メッセージ):
"""ツール出力が長期にわたる履歴に入る前に、最後の 20 行のみを保持してください。"""
結果 = []
メッセージ内のメッセージの場合:
メッセージがあれば。 get ( "役割" ) == "ツール" :
コンテンツ = メッセージ . get ( "コンテンツ" , "" )
行 = コンテンツ 。分割線()
len (行数) > 20 の場合:
= " \n " を維持します。結合 (行 [ - 20 :])
msg = { ** msg , "コンテンツ" : f"[… { len ( 行 ) - 20 } 行が切り捨てられました] \n { 保持 } " }
結果。追加 ( メッセージ )
結果を返す
マネージャー = プロンプトマネージャー (
トランジションモード = "エージェントサイクル" ,
トランジションフック = トリムツール出力 、
）
# Functional API: フックを直接渡す
# state = function.append_message(state, msg,transition_mode="agent_cycle",
#transition_hook=trim_tool_outputs)
フックは、コミットされている短期メッセージのリストを受け取り、長期履歴に実際に存在するものを返します。メッセージを完全に削除し、要約し、バイナリ BLOB を説明に置き換えます。返されたリストがキャッシュされます。
会話中に変化するコンテキストには 2 つの種類があり、異なる配置が必要です。
揮発性コンテキストは呼び出しごとに渡され、保存されることはありません。キャッシュされたプレフィックスの後に置かれ、何も無効にすることができません。
永続コンテキストはキー付きシステム メッセージとして追加され、通常の遷移パスに乗って時間的順序が維持されます。ミッドターン更新は、プレフィックスに埋もれず、注目度の高いリストの最後に配置されます。
マネージャー = プロンプトマネージャー (
static_system_prompt = システム 、
Initial_context =Initial_world_state , # 作成時に安定したプレフィックスをシードします
context_key = "ワールド" 、
max_tokens = 8_000 、
）
# 後で、何かが変わったとき:
マネージャー。 append_context ( "更新: インベントリには 3 つのキーが含まれています。" )
圧縮時に、初期コンテキストとすべての

そのデルタは統合されます。ContextConsolidationHook はキーのすべてのメッセージを受信し (最初のブロック、デルタの順)、新しく完全に書き換えられたコンテキストを返します。このコンテキストは、圧縮された履歴の先頭 (静的システム プロンプトの直後) に配置されます。デフォルトのフックはロスレスに連結します。独自のものを指定して diff を適用したり、LLM で要約したりできます。
def consolidate ( key ,messages ):
return rewrite_world_state (base = メッセージ [ 0 ][ "コンテンツ" ],
デルタ = [ m [ "コンテンツ" ] メッセージ内の m の場合 [ 1 :]])
manager = PromptManager (..., context_consolidation_hook = consolidate )
キー付きメッセージは非可逆圧縮フックを通過することはありません。統合と非可逆圧縮は別のフェーズであり、圧縮がトリガーされると両方とも常に実行されます (プレフィックスはとにかく書き換えられるため、最後まで圧縮します)。
状態の外に存在するもの: 静的システム プロンプトと呼び出しごとの揮発性動的プロンプトという正確に 2 つのものがシリアル化された状態では保持されません。永続コンテキストとそのデルタを含むその他すべては、 dumps() /loads() を介してラウンドトリップされます。ステートレス パターンでは、再水和状態 + 一定の静的プロンプト = 完全な会話になります。
長期間の履歴が max_tokens を超えると、圧縮フックによって max_tokens // 2 (構成可能) に減らされます。デフォルトのフックは古いものから順に切り捨てられます。代わりに独自の要約を提供してください。
def summary (messages 、 target_tokens 、adapter ):
summary = call_llm_to_summarise (メッセージ)
return [{ "役割" : "システム" , "コンテンツ" : 概要 }]
manager = PromptManager (max_tokens = 8_000、compaction_hook = Summary)
# Functional API: コンパクションは明示的な呼び出しです
# state = function.compact(state, max_tokens=8_000, Compaction_hook=summarise)
境界メタデータ
with_metadata=True を build_mess に渡します

age() を使用して、予測されたキャッシュ可能なプレフィックス レイアウトも取得します。これは、独自のテストでのロギング、デバッグ、またはプレフィックスの安定性の確認に役立ちます。
メッセージ、メタ = マネージャー。 build_messages (dynamic_system_prompt = rag、with_metadata = True)
# メタ == {"境界": [0, 12], "プレフィックスメッセージ数": 13,
# "prefix_tokens": 4203、"suffix_tokens": 310、"total_tokens": 4513}
これらはプレフィックスの安定性からの予測です。実際のキャッシュ ヒットは、プロバイダーの応答使用状況メタデータでのみ報告されます。
Compact() は、長期にわたる履歴を個別に予算化します。実際の制約がリクエスト全体 (静的システム + 履歴 + 動的コンテキスト ≤ コンテキスト ウィンドウ) である場合は、 Compact_for_request() を使用します。
状態 = 機能的。リクエスト用のコンパクト (
状態、
request_budget = 128_000 , # リクエスト全体のトークン予算
static_system_prompt = SYSTEM , # 測定済み (契約により安定しています)
reserved_tokens = 8_000 、 # 動的 + 短期コンテンツ用の固定ヘッドルーム
）
# または: manager.compact_for_request(request_budget=128_000,reserved_tokens=8_000)
reserved_tokens は、現在のターンの測定値ではなく、意図的に宣言された定数です。バジェットが変動する動的コンテンツを追跡している場合、圧縮はあるターンではトリガーされ、次のターンではトリガーされない可能性があります。つまり、長期存続するプレフィックスを書き換えてキャッシュを無効にする可能性があります。最悪のケースを予約しておくと、導き出される予算は確定的なままになります。
アダプター
キャッシュマーカー
トークンカウンティング
OpenAIAdapter (デフォルト)
不要 - 自動プレフィックス キャッシュ
~4 文字/トークン
人間アダプター
cache_control: {type: ephemeral} がプレフィックス境界に挿入されました
~4 文字/トークン
トランスフォーマーアダプター(tok)
なし
HF トークナイザーによる正確性
新しいプロバイダーを追加するための ProviderAdapter をサブクラス化します。 count_tokens() および/または apply_cache_markers() をオーバーライドします。
ベンチマーク スイートは、両方の llmbuffer を通じてマルチターン会話を実行します。

素朴なアプローチで、プロバイダー自身の使用状況メタデータからキャッシュ ヒットを報告します。
単純なアプローチでは、すべてのメッセージ リストの先頭に静的および動的システム プロンプトがまとめられ、コンテキスト制限に達すると最も古いメッセージが削除されます。これは、今日のほとんどの LLM アプリケーションのデフォルト パターンです。
結果 (シミュレーション、15 ターン、人為的価格設定)
シミュレートされたプロバイダーは、プロバイダー プレフィックス キャッシュを正確にモデル化します。つまり、メッセージ リストが以前に見たターンとプレフィックスを共有する場合、ターンはキャッシュ ヒットとなります。ライブ数値に対して --provider anthropic または --provider openai を実行します。
結論: llmbuffer の入力コストは、単純なアプローチよりも最大 43% 低くなります (同じ 15 ターンの会話の場合、$0.016 対 $0.028)。
キャッシュ ヒット率は、キャッシュから提供される入力トークンの割合です。キャッシュされたトークンは無料ではないため、1:1 は節約にはつながりません。Anthropic の価格設定では、キャッシュ読み取りにはキャッシュされていないレートの 10% のコストがかかります。したがって、節約とまったくキャッシュを行わない実行の比較 = ヒット率 × 90%: llmbuffer は 76.7% を節約し、naive は 59.5% を節約します。
llmbuffer と naive の比較はドルでの数字です: $0.016 / $0.028 ≈ 57%、つまり ~43% 安いです。ギャップはミス ターンから生じます。動的コンテキストがローテーションするたび (ターン 4、7、10、13)、単純なアプローチでは完全なキャッシュ ミスが発生します。

[切り捨てられた]

## Original Extract

LLM conversation buffer with cache optimization and dynamic context. - scottpurdy/llmbuffer

Releasing llmbuffer - a Python library that maximizes LLM prompt cache hits. It handles dynamic context, compaction, and tool output truncation or summarization via flexible hooks. Expect 10x cost savings in typical usage. It supports:
- Stateful (PromptManager) and stateless/functional (pure dict to dict) interfaces
- Three transition modes for moving messages from short term to long term message histories: none, manual, and agent_cycle. The agent_cycle mode, for instance, let's you truncate, summarize, or strip tool outputs when moving messages to long term history after an agent loop finishes
- Pluggable compaction hooks when history gets long
- Provider adapters for Anthropic (cache_control markers), OpenAI (automatic prefix caching), Gemini, and local models via HF tokenizers
- A benchmark suite with --compare mode showing llmbuffer vs naive side-by-side, reading real cache hit counts from each provider's usage metadata. In my usage, I see >90% of tokens are cached despite frequent dynamic context changing. MIT license.
GitHub: https://github.com/scottpurdy/llmbuffer
PyPI: https://pypi.org/project/llmbuffer/

GitHub - scottpurdy/llmbuffer: LLM conversation buffer with cache optimization and dynamic context. · GitHub
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
scottpurdy
/
llmbuffer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits src/ llmbuffer src/ llmbuffer tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Cache-optimized LLM conversation history management.
Most LLM applications naively concatenate their system prompt, conversation history, and any dynamic context into a single message list — and rebuild it from scratch every turn. This works, but it leaves significant money and latency on the table by constantly invalidating the provider's prompt cache.
llmbuffer assembles your messages in the order that maximises cache reuse, manages the boundary between stable and changing content, and handles compaction when history grows too long — all without you having to think about it.
[Static System Prompt] → [Long-Lived History] → [Dynamic Context] → [Recent Messages]
cached ✓ cached ✓ not cached not cached
The static system prompt and committed conversation history form a byte-stable prefix that is never mutated or re-ordered across turns. The frequently-changing parts — RAG results, timestamps, in-flight tool calls — live at the end where they can't invalidate the prefix.
pip install llmbuffer
Optional extras for live benchmarking:
pip install " llmbuffer[anthropic] " # Anthropic prompt caching
pip install " llmbuffer[openai] " # OpenAI prefix caching
llmbuffer has zero required dependencies — just Python 3.9+.
from llmbuffer import PromptManager , AnthropicAdapter
manager = PromptManager (
static_system_prompt = "You are a senior software engineering assistant..." ,
transition_mode = "agent_cycle" , # auto-commit turns to the stable prefix
adapter = AnthropicAdapter (), # inject cache_control markers
max_tokens = 8_000 , # compact long-lived history beyond this
)
# Each turn:
manager . append ({ "role" : "user" , "content" : user_message })
messages = manager . build_messages ( dynamic_system_prompt = rag_context )
reply = anthropic_client . messages . create ( messages = messages , ...)
manager . append ({ "role" : "assistant" , "content" : reply })
Stateless (web app / serverless)
Pure functions over a JSON-serializable state dict — persist it anywhere between requests:
from llmbuffer import functional , new_state , dumps , loads
SYSTEM = "You are a senior software engineering assistant..."
# Load state from DB / session
state = loads ( row . conversation_json ) if row else new_state ()
# Build messages, call LLM, store updated state
state = functional . append_message ( state , { "role" : "user" , "content" : text },
transition_mode = "manual" )
messages = functional . build_messages ( state , static_system_prompt = SYSTEM ,
dynamic_system_prompt = rag_context )
# ... call your LLM ...
state = functional . append_message ( state , reply , transition_mode = "manual" )
state = functional . compact ( state , max_tokens = 8_000 ) # explicit in the functional API
row . conversation_json = dumps ( state )
Each function takes only the settings it uses — there's no config object to thread through. Compaction is an explicit compact() call in the functional API (the stateful PromptManager runs it automatically).
build_messages() always emits messages in this exact order:
Control when messages graduate from short-term into the stable long-lived history:
Before messages move from short-term into the long-lived (cached) history, an optional transition_hook can rewrite them — useful for trimming verbose tool outputs or stripping content you don't want locked into the stable prefix forever.
def trim_tool_outputs ( messages ):
"""Keep only the last 20 lines of any tool output before it enters long-lived history."""
result = []
for msg in messages :
if msg . get ( "role" ) == "tool" :
content = msg . get ( "content" , "" )
lines = content . splitlines ()
if len ( lines ) > 20 :
kept = " \n " . join ( lines [ - 20 :])
msg = { ** msg , "content" : f"[… { len ( lines ) - 20 } lines truncated] \n { kept } " }
result . append ( msg )
return result
manager = PromptManager (
transition_mode = "agent_cycle" ,
transition_hook = trim_tool_outputs ,
)
# Functional API: pass the hook directly
# state = functional.append_message(state, msg, transition_mode="agent_cycle",
# transition_hook=trim_tool_outputs)
The hook receives the list of short-term messages being committed and returns whatever should actually land in long-lived history. Drop messages entirely, summarise them, replace binary blobs with descriptions — the returned list is what gets cached.
Context that changes during a conversation comes in two flavours, and they want different placement:
Volatile context is passed per-call and never stored — it sits after the cached prefix where it can't invalidate anything.
Durable context is appended as a keyed system message, riding the normal transition path so temporal ordering is preserved — a mid-turn update lands at the high-attention end of the list, not buried in the prefix:
manager = PromptManager (
static_system_prompt = SYSTEM ,
initial_context = initial_world_state , # seeds the stable prefix at creation
context_key = "world" ,
max_tokens = 8_000 ,
)
# later, when something changes:
manager . append_context ( "Update: the inventory now contains 3 keys." )
At compaction time, the initial context and all its deltas are consolidated : a ContextConsolidationHook receives every message for a key (initial block first, deltas in order) and returns the new, fully rewritten context, which is placed at the front of the compacted history — right after the static system prompt. The default hook concatenates losslessly; supply your own to apply diffs or summarise with an LLM:
def consolidate ( key , messages ):
return rewrite_world_state ( base = messages [ 0 ][ "content" ],
deltas = [ m [ "content" ] for m in messages [ 1 :]])
manager = PromptManager (..., context_consolidation_hook = consolidate )
Keyed messages never pass through the lossy compaction hook — consolidation and lossy compaction are separate phases, and both always run once compaction triggers (the prefix is being rewritten anyway, so compact all the way down).
What lives outside the state: exactly two things are not carried in the serialized state — the static system prompt and the per-call volatile dynamic prompt . Everything else, including durable context and its deltas, round-trips through dumps() / loads() . In the stateless pattern: rehydrated state + your constant static prompt = the complete conversation.
When the long-lived history exceeds max_tokens , a compaction hook reduces it to max_tokens // 2 (configurable). The default hook truncates oldest-first; supply your own to summarise instead:
def summarise ( messages , target_tokens , adapter ):
summary = call_llm_to_summarise ( messages )
return [{ "role" : "system" , "content" : summary }]
manager = PromptManager ( max_tokens = 8_000 , compaction_hook = summarise )
# Functional API: compaction is an explicit call
# state = functional.compact(state, max_tokens=8_000, compaction_hook=summarise)
Boundary metadata
Pass with_metadata=True to build_messages() to also get the predicted cacheable-prefix layout — useful for logging, debugging, or asserting prefix stability in your own tests:
messages , meta = manager . build_messages ( dynamic_system_prompt = rag , with_metadata = True )
# meta == {"boundaries": [0, 12], "prefix_message_count": 13,
# "prefix_tokens": 4203, "suffix_tokens": 310, "total_tokens": 4513}
These are predictions from prefix stability; actual cache hits are only reported in the provider's response usage metadata.
compact() budgets the long-lived history in isolation. When your real constraint is the whole request (static system + history + dynamic context ≤ context window), use compact_for_request() :
state = functional . compact_for_request (
state ,
request_budget = 128_000 , # whole-request token budget
static_system_prompt = SYSTEM , # measured (it's stable by contract)
reserved_tokens = 8_000 , # fixed headroom for dynamic + short-term content
)
# or: manager.compact_for_request(request_budget=128_000, reserved_tokens=8_000)
reserved_tokens is deliberately a declared constant, not a measurement of the current turn: if the budget tracked the fluctuating dynamic content, compaction could trigger on one turn and not the next — rewriting the long-lived prefix and invalidating the cache. Reserve your worst case and the derived budget stays deterministic.
Adapter
Cache markers
Token counting
OpenAIAdapter (default)
None needed — automatic prefix caching
~4 chars/token
AnthropicAdapter
cache_control: {type: ephemeral} injected at prefix boundaries
~4 chars/token
TransformersAdapter(tok)
None
Exact via HF tokenizer
Subclass ProviderAdapter to add a new provider — override count_tokens() and/or apply_cache_markers() .
The benchmark suite runs a multi-turn conversation through both llmbuffer and a naive approach, and reports cache hits from the provider's own usage metadata.
The naive approach puts the static and dynamic system prompts together at the start of every message list and drops the oldest messages when the context limit is hit — this is the default pattern in most LLM applications today.
Results (simulated, 15 turns, Anthropic pricing)
The simulated provider models provider prefix caching exactly: a turn is a cache hit when its message list shares a prefix with a previously-seen turn. Run --provider anthropic or --provider openai for live numbers.
Bottom line: llmbuffer's input costs are ~43% lower than the naive approach ($0.016 vs $0.028 for the same 15-turn conversation).
The cache hit ratio is the fraction of input tokens served from cache. It doesn't translate 1:1 into savings because cached tokens aren't free — on Anthropic pricing, cache reads cost 10% of the uncached rate. So savings vs. running with no caching at all = hit ratio × 90%: llmbuffer saves 76.7%, naive saves 59.5%.
The llmbuffer-vs-naive comparison is the dollar figures: $0.016 / $0.028 ≈ 57%, i.e. ~43% cheaper. The gap comes from the miss turns: every time the dynamic context rotates (turns 4, 7, 10, 13) the naive approach takes a full cache miss —

[truncated]
