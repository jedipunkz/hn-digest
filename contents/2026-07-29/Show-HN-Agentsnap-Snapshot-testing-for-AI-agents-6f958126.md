---
source: "https://github.com/iamfaham/AgentSnap"
hn_url: "https://news.ycombinator.com/item?id=49099464"
title: "Show HN: Agentsnap – Snapshot testing for AI agents"
article_title: "GitHub - iamfaham/AgentSnap: Deterministic snapshot testing for AI agents · GitHub"
author: "iamfaham"
captured_at: "2026-07-29T17:05:55Z"
capture_tool: "hn-digest"
hn_id: 49099464
score: 4
comments: 0
posted_at: "2026-07-29T16:18:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agentsnap – Snapshot testing for AI agents

- HN: [49099464](https://news.ycombinator.com/item?id=49099464)
- Source: [github.com](https://github.com/iamfaham/AgentSnap)
- Score: 4
- Comments: 0
- Posted: 2026-07-29T16:18:08Z

## Translation

タイトル: Show HN: Agentsnap – AI エージェントのスナップショット テスト
記事のタイトル: GitHub - iamfaham/AgentSnap: AI エージェントの決定論的スナップショット テスト · GitHub
説明: AI エージェントの決定論的スナップショット テスト。 GitHub でアカウントを作成して、iamfaham/AgentSnap の開発に貢献してください。

記事本文:
GitHub - iamfaham/AgentSnap: AI エージェントの決定論的スナップショット テスト · GitHub
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
イアムファハム
/
エージェントスナップ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

ns
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
147 コミット 147 コミット .github .github Agentsnap Agentsnap アセット アセットの例 例 サイトドキュメント サイトドキュメントのテスト テスト .gitignore .gitignore .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md LICENSEライセンス README.md README.md USAGE.md USAGE.md conftest.py conftest.py main.py main.py mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントの決定論的スナップショット テスト。
ガイドや API リファレンスを含む完全なドキュメントは iamfaham.github.io/AgentSnap にあります。
Agentsnap は、ゴールデン ラン中のエージェントの LLM 呼び出しとツール呼び出しを記録し、コミットされたスナップショット ファイルを生成します。後続の実行ごとに同じ入力が再生され、新しいトレースとスナップショットが 4 つの次元にわたって比較されます。
いずれかのディメンションがしきい値を超えた場合、agentsnap は構造化された差分レポートで AgentRegressionError を生成します。
エージェントは静かに退行します。プロンプトの微調整、モデルの交換、間違った引数に接続されたツールなど、何も例外がスローされず、CI が失敗することもありません。エージェントが静かにさらに悪い答えを出し始めたときに本番環境で気づきます。
Agentsnap では、2 つの異なるジョブに対して 2 つのモードが提供されます。
すべての PR で再生 - 実際の API を呼び出す代わりに、記録された応答が再生されます。決定的でコストゼロで、コード回帰 (プロンプト編集、ツール配線の破損、呼び出し回数の変更) を捕捉します。
毎晩のライブ — 現在のモデルに対する実際の API 呼び出しで、モデル自体が変更された場合にのみ現れるドリフトをキャッチします。
リプレイ モードでキャッチされたプロンプト編集。API 呼び出しは必要ありません。
「リプレイ」におけるエージェントの回帰
============================
[ARGS] llm_call[0].messages:
メッセージ: [{'コンテンツ'

: '簡潔に答えてください: Python とは何ですか?', ...}] ->
[{'role': 'user', 'content': 'あなたは海賊です。答え: ...'}]
[セマンティック] llm_call[0]: 100% 合格
[SEMANTIC] 出力: 100% PASS
失敗したチェック: ['llm_requests']
3 分間のクイックスタート
pip インストールエージェントスナップ
2 — セットアップを実行する
エージェントスナップの初期化
セマンティック比較バックエンドを選択するよう求められます。
ウィザードは、選択内容を pyproject.toml に保存し、API キー (存在する場合) を .env に保存します。キーが pyproject.toml に書き込まれることはありません。また、 __agent_snapshots__/.last_run/ を .gitignore に追加し (必要に応じてファイルを作成します)、 testing/test_agentsnap_example.py でスナップショット テストの例をスキャフォールディングすることを提案します。
Agentsnap check # いつでも設定を確認します
3 — エージェントを記録します (コードの変更は必要ありません)
PatchSet は、インストールされているすべての LLM SDK にクラス レベルでパッチを適用します。どこで作成された生のクライアントも自動的にキャプチャされます。クライアントをラップする必要はありません。
Agentsnap インポート PatchSet 、 AgentRecorder から
輸入人間
# 既存のエージェント — そのまま
def my_agent (質問):
クライアント = 人間。 Anthropic () # 生のクライアント、ラッパーは必要ありません
クライアントを返す。メッセージ。作成する （...）。コンテンツ [ 0 ]。テキスト
# 最初の実行: ゴールデン スナップショットを記録します
PatchSet () を使用:
AgentRecorder ( "my_agent" ) を rec として使用:
result = my_agent (「Python とは何ですか?」)
記録。出力 = 結果
# __agent_snapshots__/my_agent.json を書き込みます — このファイルをコミットします
4 — 今後の実行についてアサートする
Agentsnap import PatchSet 、 AgentAsserter から
PatchSet () を使用:
AgentAsserter ( "my_agent" ) を次のように使用します。
result = my_agent (「Python とは何ですか?」)
。出力 = 結果
# 動作がドリフトした場合は AgentRegressionError を発生させます
5 — pytest フィクスチャを使用する (最も単純)
snapshot.run() は最初の呼び出し時に自動記録し、その後の実行ごとに自動アサートします。切り替えは必要ありません。 PatchSet オートマチックをアクティブ化するために Agentsnap_instrument を追加します

呼んで:
def test_my_agent (スナップショット、agentsnap_instrument):
スナップショット付き。 ( "my_agent" ) を s として実行します。
result = my_agent ( "Python とは何ですか?" ) # 生のクライアント — 自動的にキャプチャされます
s 。出力 = 結果
pytest
# または、セッション内のすべてのテストに対して PatchSet を有効にします。
pytest --agentsnap-instrument
リプレイモードとライブモード
すべてのアサートは、次の 2 つのモードのいずれかで実行できます。
再生モードでは、各 LLM コールに対して記録された応答がフィードバックされます。
エージェント — API キーもコストもフレークもありません。比較はリクエストに切り替わります
サイド: コードが異なるプロンプトを送信する場合、agentsnap はテストに失敗します。
LLM 呼び出しの数が異なるか、ツールのシーケンスが変更されます。
テストごとの#
AgentAsserter ( "my_agent" 、 mode = "replay" ) を次のように使用します。
# スイート全体
pytest - - Agentsnap -replay # 強制リプレイ
pytest - - Agentsnap - ライブ # ライブを強制する
「ツール。エージェントスナップ]
mode = "replay" #プロジェクトのリプレイをデフォルトにします
ツール呼び出しは、リプレイ モードでも実際に実行されます。にreplay_tools=Trueを渡します
録音からもスタブします (副作用はまったくありません)。
再生には、agentsnap >= 0.2.0 で記録されたスナップショットが必要です (スナップショットには、
raw_response )。古いスナップショットでは SnapshotFormatError が発生する - 再記録
pytest --agentsnap-record を使用します。
Replay は現在、Anthropic、OpenAI、Groq、OpenRouter をサポートしています。
他のプロバイダーは ReplayError を発生させます。それらのテストにはライブ モードを使用してください。
シナリオの場合、再生モードで明示的に scenario= を渡します (自動ハッシュを入力します)
テスト本体が実行される前にスナップショットが読み取られるため、使用できません)。
再生された最終出力がゴールデンとバイト同一でない場合、スコアリングします。
セマンティック バックエンドが必要 — 埋め込みバックエンドをインストールして構成する
( pip install Agentsnap[offline] 、次に Agentsnap init オプション 2) またはジャッジを設定する
( AGENTSNAP_JUDGE_API_KEY )。
非同期クライアント ( AsyncAnthropic 、 AsyncOpenAI )

eも傍受されました—
リプレイのネットワークなし保証は、同期クライアントと非同期クライアントの両方をカバーします。
非同期ストリーム。残る 1 つの穴は、ストリーミングされた OpenAI Response API です。
(response.create(stream=True))、未記録のまま通過します。参照
例/async_agents.py 。
AnthropicAdapter および OpenAIAdapter ティー stream=True 呼び出し (Groq および
OpenRouter はこれを OpenAIAdapter から継承します): チャンクは
組み立てられた応答が再生のために記録される間、エージェントは変更されません。
raw_response={"__stream__": True、"chunks": [...]} 。
再生モードでは、記録されたチャンクが実際の SDK チャンク/イベントに再構築されます。
オブジェクトを生成し、段階的に返還します - エージェントはそれらを正確に消費します
ライブ ストリームのようなもので、API 呼び出しはありません。ストリーミングからの録音
呼び出しは非ストリーミング リクエストとして再生できません (またはその逆)。
ReplayError と明確な「形状の不一致」メッセージ。
まだサポートされていません: client.messages.stream() コンテキスト マネージャー ヘルパー、
OpenAI Responses-API 呼び出しのストリーミング。ミストラルは依然として stream=False を強制します
すべての通話で。完全な実行可能ファイルについては、examples/streaming.py を参照してください。
ウォークスルー、および非同期クライアント バージョンの example/async_agents.py。
決して反復されず、決して閉じられないストリームは自動的にファイナライズされます
Recorder/Asserter の終了時ですが、すぐに消費または終了することはまだです
イベントが呼び出し順に表示されるようにすることをお勧めします。
コードが実際に実行されるツール以外にも、agentsnap はどのツールをキャプチャするかも行います。
モデルが呼び出すことを決定したツール。すべての非ストリーミング Anthropic/OpenAI
llm_call イベントは、tool_requests リストを記録します。tool_use は、
返されるモデルは、それぞれ {"name": ..., "args": {...}} として返されます。 Groq と OpenRouter
OpenAIAdapter をサブクラス化しているため、これを無料で入手できます。
アサート時に、agentsnap はモデルの要求されたツール シーケンス (単に
あなたは何ですか

コードが実行されます)、変更された場合、model_tools が失敗します。または
同じツールが異なる引数でリクエストされた場合は、model_tool_args —
レポートでは [MODEL TOOLS] ... として表示されます。こちらは静かにモデルを捕まえます
コードが独自のものであっても、ゴールデン ランとは異なるツールを選択する
ツール呼び出しロジックは変更されません (モデルの更新、プロンプト挿入、
プロバイダー側の回帰）。
[モデルツール] モデルが要求したツールシーケンスが変更されました (編集距離 1): ['search'] -> ['delete_file']
[ARGS] モデルツール:検索->ファイルの削除[0]:
引数: {'クエリ': 'フランスの首都'} -> {'パス': '/etc/passwd'}
失敗したチェック: ['model_tools'、'model_tool_args']
下位互換性: llm_call が実行されるたびにのみ比較が行われます。
diff の両側のイベントには、tool_requests が含まれます。この門に注目してください
イベントごとではなくトレース全体です: 単一のストリーミング呼び出しまたは非 Anthropic/OpenAI
トレース内の任意の場所で呼び出しを行うと、実行全体のモデル ツール チェックが無効になります。
古いゴールデン (この機能の前に記録されたもの) は、新しい表面から失敗することはありません。
現在の範囲: 非ストリーミング Anthropic 呼び出しと OpenAI 呼び出し、および Groq/OpenRouter
継承を介して。ストリーミングされた tool_use アセンブリはまだキャプチャされていません。
実行可能な完全なチュートリアルについては、examples/model_tools.py を参照してください。
プロバイダー
アダプター
インターセプト
人間的
人間アダプター
.messages.create()
OpenAI
OpenAIAdapter
.chat.completions.create()
Google ジェミニ
ジェミニアダプター
.models.generate_content()
コヒア
Cohereアダプター
.chat()
ミストラル
ミストラルアダプター
.chat.complete()
グロク
Groqアダプター
.chat.completions.create()
オープンルーター
オープンルーターアダプター
.chat.completions.create()
ランググラフ
ランググラフアダプター
.invoke() + コールバックを介したノードレベルの LLM/ツール イベント
任意の呼び出し可能
ツールアダプター
直接電話
必要に応じてプロバイダー SDK をインストールします。
pip install Agentsnap[google] # google-genai
pip install Agentsnap[cohere] # cohere
pip のインストール年齢

ntsnap[ミストラル] #ミストラライ
pip install Agentsnap[groq] # groq
pip install Agentsnap[すべてのプロバイダー]
フレームワークと連携して動作します
フレームワークは独自の SDK クライアントを内部で構築するため、何もする必要はありません。
ラップ — PatchSet は基礎となる SDK クラス (同期および非同期) にパッチを適用します。
Anthropic/OpenAI チャット、および OpenAI Responses API)、任意のフレームワーク
それらの上に構築されたものは自動的にキャプチャされます。
普遍的なパターン — フレームワーク呼び出しをラップし、他には何も変更しません。
エージェントスナップからパッチセットをインポート
エージェントスナップから。コア。アサータインポート AgentAsserter
PatchSet () を使用:
AgentAsserter ( "my_framework_agent" ) を次のように使用します。
。出力 = my_pydantic_ai_agent 。 run_sync (「Python とは何ですか?」)。出力
注意事項:
ストリーミングされた OpenAI Responses-API の実行 (response.create(stream=True)) パス
記録されていないこの反復を通じて、非ストリーミング応答呼び出しとすべての
チャット完了ストリーミング (同期 + 非同期) が記録され、再生可能です。
モデル ツール チェック (以下を参照) はトレース全体でゲートされます。
トレースはストリーミング呼び出しまたは非 Anthropic/OpenAI プロバイダー、全体
実行のmodel_tools / model_tool_argsの比較はスキップされます。
実際のフレームワーク検証テストは、tests/frameworks/ (マーカー
フレームワーク、pytest.importorskip -guarded、別の CI ジョブ経由で実行

[切り捨てられた]

## Original Extract

Deterministic snapshot testing for AI agents. Contribute to iamfaham/AgentSnap development by creating an account on GitHub.

GitHub - iamfaham/AgentSnap: Deterministic snapshot testing for AI agents · GitHub
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
iamfaham
/
AgentSnap
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
147 Commits 147 Commits .github .github agentsnap agentsnap assets assets examples examples site-docs site-docs tests tests .gitignore .gitignore .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md USAGE.md USAGE.md conftest.py conftest.py main.py main.py mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Deterministic snapshot testing for AI agents.
Full documentation, including guides and API reference, lives at iamfaham.github.io/AgentSnap .
agentsnap records your agent's LLM and tool calls during a golden run and produces a committed snapshot file. On every subsequent run it replays the same inputs and compares the new trace against the snapshot across four dimensions:
If any dimension drifts beyond its threshold, agentsnap raises AgentRegressionError with a structured diff report.
Agents regress silently. A prompt tweak, a model swap, a tool wired to the wrong argument — nothing throws an exception, nothing fails CI, and you find out in production when the agent quietly starts giving worse answers.
agentsnap gives you two modes for two different jobs:
Replay, on every PR — recorded responses are replayed instead of calling a real API. Deterministic, zero cost, catches code regressions (prompt edits, broken tool wiring, changed call counts).
Live, nightly — real API calls against the current model, catching drift that only shows up when the model itself changes.
A prompt edit caught by replay mode, no API call required:
Agent regression in 'replay'
============================
[ARGS] llm_call[0].messages:
messages: [{'content': 'Answer concisely: What is Python?', ...}] ->
[{'role': 'user', 'content': 'You are a pirate. Answer: ...'}]
[SEMANTIC] llm_call[0]: 100% PASS
[SEMANTIC] output: 100% PASS
Failed checks: ['llm_requests']
3-minute quickstart
pip install agentsnap
2 — Run setup
agentsnap init
Asks you to choose a semantic comparison backend:
The wizard saves your choice to pyproject.toml and your API key (if any) to .env . Keys are never written to pyproject.toml . It also adds __agent_snapshots__/.last_run/ to .gitignore (creating the file if needed) and offers to scaffold an example snapshot test at tests/test_agentsnap_example.py .
agentsnap check # verify your setup at any time
3 — Record your agent (no code changes needed)
PatchSet patches all installed LLM SDKs at the class level — any raw client created anywhere is captured automatically. No need to wrap your clients:
from agentsnap import PatchSet , AgentRecorder
import anthropic
# your existing agent — untouched
def my_agent ( question ):
client = anthropic . Anthropic () # raw client, no wrapper needed
return client . messages . create (...). content [ 0 ]. text
# First run: records the golden snapshot
with PatchSet ():
with AgentRecorder ( "my_agent" ) as rec :
result = my_agent ( "What is Python?" )
rec . output = result
# Writes __agent_snapshots__/my_agent.json — commit this file
4 — Assert on future runs
from agentsnap import PatchSet , AgentAsserter
with PatchSet ():
with AgentAsserter ( "my_agent" ) as a :
result = my_agent ( "What is Python?" )
a . output = result
# Raises AgentRegressionError if behavior drifted
5 — Use the pytest fixture (simplest)
snapshot.run() auto-records on first call and auto-asserts on every run after — no switching needed. Add agentsnap_instrument to activate PatchSet automatically:
def test_my_agent ( snapshot , agentsnap_instrument ):
with snapshot . run ( "my_agent" ) as s :
result = my_agent ( "What is Python?" ) # raw client — captured automatically
s . output = result
pytest
# or enable PatchSet for every test in a session:
pytest --agentsnap-instrument
Replay vs live mode
Every assert can run in one of two modes:
In replay mode the recorded response for each LLM call is fed back to your
agent — no API key, no cost, no flakes. The comparison flips to the request
side : agentsnap fails the test if your code sends different prompts, makes a
different number of LLM calls, or changes the tool sequence.
# per test
with AgentAsserter ( "my_agent" , mode = "replay" ) as a : ...
# whole suite
pytest - - agentsnap - replay # force replay
pytest - - agentsnap - live # force live
[ tool . agentsnap ]
mode = " replay " # make replay the project default
Tool calls still execute for real in replay mode. Pass replay_tools=True to
stub them from the recording too (no side effects at all).
Replay needs snapshots recorded with agentsnap >= 0.2.0 (they include
raw_response ). Older snapshots raise SnapshotFormatError — re-record
with pytest --agentsnap-record .
Replay currently supports Anthropic, OpenAI, Groq, and OpenRouter.
Other providers raise ReplayError — use live mode for those tests.
With scenarios, pass scenario= explicitly in replay mode (input auto-hash
is not available because the snapshot is read before the test body runs).
If the replayed final output isn't byte-identical to the golden, scoring it
needs a semantic backend — install and configure the embeddings backend
( pip install agentsnap[offline] , then agentsnap init option 2) or configure a judge
( AGENTSNAP_JUDGE_API_KEY ).
Async clients ( AsyncAnthropic , AsyncOpenAI ) are intercepted too —
replay's no-network guarantee covers both sync and async clients, including
async streams. The one remaining hole is the streamed OpenAI Responses API
( responses.create(stream=True) ), which passes through unrecorded. See
examples/async_agents.py .
AnthropicAdapter and OpenAIAdapter tee stream=True calls (Groq and
OpenRouter inherit this from OpenAIAdapter ): chunks are forwarded to your
agent unmodified while the assembled response is recorded for replay, with
raw_response={"__stream__": True, "chunks": [...]} .
In replay mode the recorded chunks are rebuilt into real SDK chunk/event
objects and yielded back incrementally — your agent consumes them exactly
like a live stream, with zero API calls. A recording made from a streaming
call cannot replay as a non-streaming request (or vice versa) — that raises
ReplayError with a clear "shape mismatch" message.
Not yet supported: the client.messages.stream() context-manager helper,
and streamed OpenAI Responses-API calls. Mistral still forces stream=False
on every call. See examples/streaming.py for a full runnable
walkthrough, and examples/async_agents.py for the async-client version.
A stream that is never iterated and never closed is finalized automatically
at recorder/asserter exit, but consuming or closing it promptly is still
recommended so events appear in call order.
Beyond the tools your code actually executes, agentsnap also captures which
tool the model decided to call. Every non-streaming Anthropic/OpenAI
llm_call event records a tool_requests list — the tool_use blocks the
model returned, each as {"name": ..., "args": {...}} . Groq and OpenRouter
get this for free since they subclass OpenAIAdapter .
On assert, agentsnap compares the model's requested tool sequence (not just
what your code executed) and fails model_tools if it changed, or
model_tool_args if the same tool was requested with different arguments —
surfaced in the report as [MODEL TOOLS] ... . This catches a model quietly
choosing a different tool than the golden run even when your code's own
tool-calling logic is untouched (a model update, a prompt injection, a
provider-side regression).
[MODEL TOOLS] Model-requested tool sequence changed (edit distance 1): ['search'] -> ['delete_file']
[ARGS] model_tool:search->delete_file[0]:
args: {'query': 'capital of France'} -> {'path': '/etc/passwd'}
Failed checks: ['model_tools', 'model_tool_args']
Backward compatible: the comparison only engages when every llm_call
event on both sides of the diff carries tool_requests . Note this gate
is trace-wide, not per-event: a single streamed call or non-Anthropic/OpenAI
call anywhere in the trace disables the model-tools check for the whole run.
Old goldens (recorded before this feature) never fail from the new surface.
Scope today: non-streaming Anthropic and OpenAI calls, plus Groq/OpenRouter
via inheritance. Streamed tool_use assembly is not captured yet.
See examples/model_tools.py for a full runnable walkthrough.
Provider
Adapter
Intercepts
Anthropic
AnthropicAdapter
.messages.create()
OpenAI
OpenAIAdapter
.chat.completions.create()
Google Gemini
GeminiAdapter
.models.generate_content()
Cohere
CohereAdapter
.chat()
Mistral
MistralAdapter
.chat.complete()
Groq
GroqAdapter
.chat.completions.create()
OpenRouter
OpenRouterAdapter
.chat.completions.create()
LangGraph
LangGraphAdapter
.invoke() + node-level LLM/tool events via callbacks
Any callable
ToolAdapter
direct call
Install provider SDKs as needed:
pip install agentsnap[google] # google-genai
pip install agentsnap[cohere] # cohere
pip install agentsnap[mistral] # mistralai
pip install agentsnap[groq] # groq
pip install agentsnap[all-providers]
Works with your framework
Frameworks build their own SDK clients internally, so there's nothing to
wrap — PatchSet patches the underlying SDK classes (sync and async
Anthropic/OpenAI chat, plus the OpenAI Responses API), so any framework
built on top of them is captured automatically.
The universal pattern — wrap the framework call, nothing else changes:
from agentsnap import PatchSet
from agentsnap . core . asserter import AgentAsserter
with PatchSet ():
with AgentAsserter ( "my_framework_agent" ) as a :
a . output = my_pydantic_ai_agent . run_sync ( "What is Python?" ). output
Caveats:
Streamed OpenAI Responses-API runs ( responses.create(stream=True) ) pass
through unrecorded this iteration — non-streaming Responses calls and all
chat-completions streaming (sync + async) are recorded and replayable.
The model-tools check (see below) is gated trace-wide: if any call in the
trace is a streamed call or a non-Anthropic/OpenAI provider, the whole
run's model_tools / model_tool_args comparison is skipped.
Real-framework verification tests live in tests/frameworks/ (marker
frameworks , pytest.importorskip -guarded, run via a separate CI job wi

[truncated]
