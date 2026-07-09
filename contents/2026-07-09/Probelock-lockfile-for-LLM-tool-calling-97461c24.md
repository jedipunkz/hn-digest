---
source: "https://github.com/kelkalot/probelock"
hn_url: "https://news.ycombinator.com/item?id=48842075"
title: "Probelock – lockfile for LLM tool calling"
article_title: "GitHub - kelkalot/probelock: A capability lockfile for local models. Catch silent tool-calling regressions in CI (deterministic, no LLM judge). · GitHub"
author: "rar101x"
captured_at: "2026-07-09T07:39:30Z"
capture_tool: "hn-digest"
hn_id: 48842075
score: 1
comments: 0
posted_at: "2026-07-09T07:16:50Z"
tags:
  - hacker-news
  - translated
---

# Probelock – lockfile for LLM tool calling

- HN: [48842075](https://news.ycombinator.com/item?id=48842075)
- Source: [github.com](https://github.com/kelkalot/probelock)
- Score: 1
- Comments: 0
- Posted: 2026-07-09T07:16:50Z

## Translation

タイトル: Probelock – LLM ツール呼び出し用のロックファイル
記事のタイトル: GitHub - kelkalot/probelock: ローカル モデルの機能ロックファイル。 CI でサイレントツール呼び出しの回帰を捕捉します (決定論的、LLM 判定なし)。 · GitHub
説明: ローカル モデルの機能ロックファイル。 CI でサイレントツール呼び出しの回帰を捕捉します (決定論的、LLM 判定なし)。 - ケルカロット/プローブロック

記事本文:
GitHub - kelkalot/probelock: ローカル モデルの機能ロックファイル。 CI でサイレントツール呼び出しの回帰を捕捉します (決定論的、LLM 判定なし)。 · GitHub
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
ケルカロット
/
プローブロック
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスで表示する メイン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
24 コミット 24 コミット .github/ workflows .github/ workflows デモ デモの例 例 フィクスチャ フィクスチャ ロックファイル ロックファイル プローブロック プローブロック テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md VALIDATION-TRACES.md VALIDATION-TRACES.md VALIDATION.md VALIDATION.md action.yml action.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
A capability lockfile for local models.モデルがセットで何をしたかを記録します
ツールの呼び出しと出力のチェックが失敗し、モデル/クォント/ランタイムのスワップ時に CI が失敗する
スコアを下げます。
ラマ-3.1-8b @ Q8_0 (オラマ) → ラマ-3.1-8b @ Q4_K_M (オラマ)
能力ベースライン候補 Δ ステータス
arg_validity 1.00 0.67 -0.33 回帰
アリティ_ロバストネス 1.00 0.67 -0.33 回帰
format_adherence 1.00 1.00 +0.00 ok
ニードルインツール 1.00 0.33 -0.67 回帰
no_hallucinated_tool 1.00 0.67 -0.33 回帰
required_args 1.00 1.00 +0.00 ok
構造化出力 1.00 0.33 -0.67 回帰
ツール識別 1.00 0.33 -0.67 回帰
ツール許可 1.00 0.67 -0.33 回帰
tool_restraint 1.00 0.67 -0.33 REGRESSION
ツール選択 1.00 0.67 -0.33 回帰
FAIL — 機能が低下または削除されました: arg_validity、arity_robustness、
針の入ったツール、幻覚のないツール、構造化された出力、ツールの識別、
ツール許可、ツール拘束、ツール選択
ここでは、第 4 四半期のクオンツ スコアはいくつかの機能で 0.33 ～ 0.67 でしたが、第 8 四半期のスコアは 1.00 でした。
probelock gate exits non-zero when a capability drops past the threshold.
Promptfoo は、あなたが作成したテスト フレームワークです。プローブロックはコミットするロックファイルです。

プローブはツール スキーマから派生します。 OpenAI スタイルを指す
エージェントがすでに同梱しているツール定義では、修正された、
能力チェックの再現可能なバッテリー。テストケースは書きません。
LLM 審査員はいません。すべてのプローブはコードによってスコア付けされます: JSON スキーマ検証、正確
一致、またはツール名のチェック。同じモデルと番号で 2 回実行します
一致します。 (promptfoo は、ユーザーが作成したアサーションと、多くの場合、モデルで評価されたアサーションに依存します。
eval、実行ごとに異なります)。
モデル/クオンツ/ランタイムにわたって、モデルをその独自のベースラインと比較します。
絶対的なリーダーボードを作成するのではなく、交換します。あなたはいつも次のように比較するだけです
同様のものが、箱に、ツールが付いているので、「ベンチマークは
「ゲーム可能/ハードウェア依存」に関する異議は適用されません。
インストールして実行します ( uv のみが必要です)
インストールせずに実行するか、現在の環境にインストールします。
uvx probelock --help # 最新リリースを実行します
pip install probelock # またはそれをインストールします
未リリースのリビジョンを git から直接実行するには:
uvx --from git+https://github.com/kelkalot/probelock プローブロック --help
以下の例では、このリポジトリのチェックアウトから実行される UV を使用します。モデルは必要ありません
デモでは、決定論的な SimulatedClient が 2 つの量子レベルの代わりになります。
同じモデル:
# プローブロック/プロジェクトディレクトリから
uv runprobelockderive --tools Examples/agent_tools.json # プローブのバッテリーを参照
uv run プローブロック プローブ --tools example/agent_tools.json --simulate fixtures/profile_q8.json -o q8.lock
uv run プローブロック プローブ --tools example/agent_tools.json --simulate fixtures/profile_q4.json -o q4.lock
uv 実行プローブロック diff q8.lock q4.lock
uv run プローブロック ゲート --baseline q8.lock --candidate q4.lock # ゼロ以外で終了
ローカル モデルに対して、OpenAI 互換エンドポイントの swap --simulate を実行します。
uv run プローブロック プローブ --tools example/agent_tools.json \

--エンドポイント http://localhost:11434/v1 --model llama3.1:8b-instruct-q4_K_M \
--quant Q4_K_M --runtime ollam --timeout 120 -o q4.lock
モデルが拒否したプローブ (例: 「モデルがツールをサポートしていない」)、またはタイムアウトしたプローブ
その機能のスコアは 0 であり、実行は続行されるため、モデルは
ツール呼び出しは依然としてロックファイルを生成します。到達不能なサーバー、404 (間違ったモデルまたは
URL)、またはすべてのプローブが失敗した場合に実行が中止されるため、構成ミスが発生することはありません
ポイズンされたオールゼロのベースラインになります。
Examples/agent_tools.json は、上記のウォークスルー用の 3 つのツールのスキーマであり、
感度ベンチマーク — 検証テストでは、実際の機能の影響を受けにくいことが判明しました
重複するツール名とより豊富な引数を備えた 10 ツール スキーマの影響
制約が正しくキャッチされました ( VALIDATION.md を参照)。すぎるスキーマ
少数のツール、または実際に違反する制約のない議論、過小報告
回帰。 --tools を自分のエージェントの実際のツール定義に指定する前に
CI のゲートを信頼します。
プローブロックは 1 つのプロトコルを話します — OpenAI /v1/chat/completions with OpenAI-style
ツール — これを公開するものはすべて --endpoint で動作します。プロバイダーにとって、
(Anthropic、Gemini など)、 --via を使用して統合 SDK を介してルーティングしないでください。あらゆる道
決定論的です。どれも LLM をループに入れていません。
pip install 'probelock[anyllm]' #または 'probelock[litellm]'
プローブロック プローブ --tools tools.json --via anyllm --model misstral/mistral-large-latest \
--サンプル 5 --温度 0.7 -o 候補.lock
--via クライアントは、同じキャッシュ、サンプリング、エラー セマンティクスを再利用します。
--エンドポイント ;これらは、各 SDK の OpenAI 形式の応答上のシン アダプターです。追加
小さなクライアント プロトコルを実装することによる新しいバックエンド — それが唯一の継ぎ目です。
demo/ はローカルの Ollama サーバーに対して実行されます: コミットされた qwen3.5:9b
ベースライン vs gemma3:1b candi

日付 (ツール呼び出しはサポートされていません)。参照
トランスクリプトについては、demo/DEMO.md を参照するか、再生してください。
asciinema play Demon/probelock-demo.cast # または: bash Demon/demo.sh
ツール呼び出し機能は 1.00 → 0.00 に低下し、ゲートはゼロ以外で終了します。
tool_restrain 、tool_permission 、および no_hallucinated_tool は 1.00 のままです (a
ツールを呼び出すことができないモデルはツールを悪用することはできません)、gemma3:1b のスコアは 1.00 です。
qwen3.5:9b の format_adherence と 0.50。差分は機能ごとに異なります。
またコミット: qwen3.5:9b vs lfm2.5- Thinking:1.2b :
uv run プローブロック diff デモ/qwen3.5-9b.lock デモ/lfm2.5- Thinking.lock
1.2B モデルは、ツールの選択、識別、
need_in_tools 、 arg_validity 、 required_args 、および 3 つの安全プローブ。
Structured_output と arity_robustness は 1.00 → 0.33 に低下します。
能力 (すべて決定的にスコア付けされます)
能力
チェック内容
得点者
ツール選択
タスクに適したツールを呼び出します
ツール名の一致
ツール_差別
他のツールではなく適切なツールを呼び出します (正確に選択します)
ツール名セット
ニードルインツール
多数 (15 以上) のツールが提供されている場合でも適切なツールを見つける
ツール名の一致
arg_validity
発行された引数はツールの JSON スキーマに対して検証されます
jsonschema
required_args
必要な引数がすべて存在し、空ではない
鍵の存在
アリティ_ロバストネス
要求されたら、すべてのパラメータ (必須 + オプション) を入力します
全員が存在する
構造化された出力
スキーマが有効な JSON をオンデマンドで出力します (ツールやフェンスは不要)
解析 + jsonschema
json_mode (オプトイン、 --json-mode )
同じですが、プロンプトの代わりにサーバーのネイティブな response_format API を使用します。
解析 + jsonschema
ツール拘束
ツールを必要としないタスクに対してはツールを呼び出しません (オーバートリガー)
ツール呼び出しなし
ツール_権限
明示的に使用が禁止されているツールを呼び出さない
禁止ツールが存在しない
no_hallucinated_tool
いいえだったツールへの呼び出しを捏造しません。

提供されました
電話をかけました ⊆ 申し出ました
フォーマットの遵守
正確な出力制約に従います
完全一致
3 つはネガティブなプローブです (スコアが高いほど、悪い動作が行われていないことを意味します)
発生）：tool_restrain（オーバートリガー）、tool_permission（禁止されたものの呼び出し）
ツール)、および no_hallucinated_tool (ツールの製造)。すべてのプローブは以下から派生します。
ツールのスキーマは手作業で作成されたものではありません。
ツールスキーマ ──▶ プローブの導出 ──▶ クライアント ──▶ ResponseMessage ──▶ 決定的スコアラー ──▶ ロックファイル
(あなたのエージェント) (ゼロオーサリング) (モデル) (唯一のモデル (LLM ジャッジなし) (コミット)
-触れる部分）
ロックファイル + ロックファイル ──▶ diff / ゲート
唯一の非決定的な部分は Client です。それ以外はすべて純粋なので、
同じ入力では、同じロックファイルと同じ差分が生成されます。温度0では、
クライアントは同一のリクエストをキャッシュするため、1 つのリクエストを共有するプローブ (ツール
特定のツールをチェックします) ネットワークに 1 回アクセスします。 SimulatedClient のクラフトが正しいか、
実際の採点者が採点する不正確な回答により、採点パスは均等に実行されます。
モデルが存在しない場合。
実際のトレースからプローブを導出する (実験的)
スキーマ派生プローブはシングルターンで合成されており、スキーマレベルの捕捉に最適です
回帰、実際のコンテキストを数回繰り返した後に何が壊れるかが分からない、ツールの結果
フィードバック、または曖昧な表現。 --traces は 2 番目のソースを追加します: real、
すでに記録されているエージェントの意思決定ポイント (例: litellm の OpenTelemetry からエクスポートされたもの)
コールバック）、まったく同じ決定論的スコアラーを通じて再生されます。
uv run プローブロック 派生 --tools tools.json --traces traces.json # 追加される内容を確認します
uv run プローブロック プローブ --tools tools.json --traces traces.json \
--endpoint http://localhost:11434/v1 --model llama3.1:8b -ocandidate.lock
--tools はここではオプションです。トレースされたプローブは、埋め込まれた独自のプローブを再生します。

オールの定義、
したがって、トレースのみの実行 (probe --traces traces.json ...) にはスキーマ ファイルは必要ありません。同じ
以下の --mined に当てはまります。合成バッテリーも必要な場合は、--tools を指定してください。
トレース ファイルは、生のファイルではなく、プローブロック自体が定義する小さくて安定した JSON スキーマです。
OpenTelemetry — OTel 独自の Span 属性レイアウトがライブラリ間で安定していないため、または
バージョン (litellm はリクエスト/レスポンス属性を一度配置する場所を既に変更しています)
には、より新しい、異なる形状のオプトイン統合があります)。エクスポートをこの形式に変換します
あなたが所有する 1 回限りのステップです。参照してください
の例/otel_traces_to_probelock.py
文書化された開始点とフィクスチャ/sample_traces.json
ターゲット形状の場合:
{
「バージョン」: 1 、
「レコード」: [
{
"id" : " checkout-flow-turn-3 " ,
"メッセージ" : [{ "ロール" : " ユーザー " , "コンテンツ" : " ... " }],
"tools" : [ /* この時点で実際に提供される OpenAI スタイルのツール定義 */ ],
"response" : { "content" : null 、 "tool_calls" : [{ "name" : " ... " 、 "arguments" : " {...} " }]}
}
】
}
トレース派生プローブは、スキーマ派生プローブと同じ機能を結合します — tools_selection 、
tool_discrimination 、 arg_validity 、 required_args 、および Structured_output — これらは
「この実際のコンテキストを再生し、候補が依然として有効に動作することを確認する」にきれいにマップします (プローブ
ID c

[切り捨てられた]

## Original Extract

A capability lockfile for local models. Catch silent tool-calling regressions in CI (deterministic, no LLM judge). - kelkalot/probelock

GitHub - kelkalot/probelock: A capability lockfile for local models. Catch silent tool-calling regressions in CI (deterministic, no LLM judge). · GitHub
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
kelkalot
/
probelock
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
24 Commits 24 Commits .github/ workflows .github/ workflows demo demo examples examples fixtures fixtures lockfiles lockfiles probelock probelock tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md VALIDATION-TRACES.md VALIDATION-TRACES.md VALIDATION.md VALIDATION.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
A capability lockfile for local models. It records what a model does on a set
of tool-calling and output checks, and fails CI when a model/quant/runtime swap
lowers a score.
llama-3.1-8b @ Q8_0 (ollama) → llama-3.1-8b @ Q4_K_M (ollama)
Capability Baseline Candidate Δ Status
arg_validity 1.00 0.67 -0.33 REGRESSION
arity_robustness 1.00 0.67 -0.33 REGRESSION
format_adherence 1.00 1.00 +0.00 ok
needle_in_tools 1.00 0.33 -0.67 REGRESSION
no_hallucinated_tool 1.00 0.67 -0.33 REGRESSION
required_args 1.00 1.00 +0.00 ok
structured_output 1.00 0.33 -0.67 REGRESSION
tool_discrimination 1.00 0.33 -0.67 REGRESSION
tool_permission 1.00 0.67 -0.33 REGRESSION
tool_restraint 1.00 0.67 -0.33 REGRESSION
tool_selection 1.00 0.67 -0.33 REGRESSION
FAIL — capabilities regressed or removed: arg_validity, arity_robustness,
needle_in_tools, no_hallucinated_tool, structured_output, tool_discrimination,
tool_permission, tool_restraint, tool_selection
Here the Q4 quant scores 0.33–0.67 on several capabilities where Q8 scored 1.00.
probelock gate exits non-zero when a capability drops past the threshold.
promptfoo is a test framework you author. probelock is a lockfile you commit.
Probes are derived from your tool schemas. Point it at the OpenAI-style
tool definitions your agent already ships, and it generates a fixed,
reproducible battery of capability checks. You write no test cases.
No LLM judge. Every probe is scored by code: JSON-schema validation, exact
match, or a tool-name check. Run it twice on the same model and the numbers
match. (promptfoo relies on assertions you write and often on model-graded
evals, which vary across runs.)
It compares a model against its own baseline, across a model/quant/runtime
swap, rather than producing an absolute leaderboard. You only ever compare like
with like, on your box, with your tools, so the "benchmarks are
gameable/hardware-dependent" objection does not apply.
Install & run (only needs uv )
Run it without installing, or install it into the current environment:
uvx probelock --help # run the latest release
pip install probelock # or install it
To run an unreleased revision straight from git:
uvx --from git+https://github.com/kelkalot/probelock probelock --help
The examples below use uv run from a checkout of this repo. No model is required
for the demo — a deterministic SimulatedClient stands in for two quant levels of
the same model:
# from the probelock/ project dir
uv run probelock derive --tools examples/agent_tools.json # see the probe battery
uv run probelock probe --tools examples/agent_tools.json --simulate fixtures/profile_q8.json -o q8.lock
uv run probelock probe --tools examples/agent_tools.json --simulate fixtures/profile_q4.json -o q4.lock
uv run probelock diff q8.lock q4.lock
uv run probelock gate --baseline q8.lock --candidate q4.lock # exits non-zero
Against a local model, swap --simulate for an OpenAI-compatible endpoint:
uv run probelock probe --tools examples/agent_tools.json \
--endpoint http://localhost:11434/v1 --model llama3.1:8b-instruct-q4_K_M \
--quant Q4_K_M --runtime ollama --timeout 120 -o q4.lock
A probe the model rejects (e.g. "model does not support tools") or that times out
scores 0 for that capability and the run continues, so a model that cannot
tool-call still produces a lockfile. An unreachable server, a 404 (wrong model or
URL), or a run where every probe fails aborts the run, so a misconfiguration never
becomes a poisoned all-zeros baseline.
examples/agent_tools.json is a 3-tool schema for the walkthrough above, not a
sensitivity benchmark — validation testing found it insensitive to real capability
drift that a 10-tool schema with overlapping tool names and richer argument
constraints caught cleanly (see VALIDATION.md ). A schema with too
few tools, or arguments with no real constraints to violate, under-reports
regressions. Point --tools at your own agent's actual tool definitions before
trusting gate in CI.
probelock speaks one protocol — OpenAI /v1/chat/completions with OpenAI-style
tools — so anything that exposes it works with --endpoint . For providers that
do not (Anthropic, Gemini, …), route through a unified SDK with --via . Every path
is deterministic; none of them put an LLM in the loop.
pip install ' probelock[anyllm] ' # or 'probelock[litellm]'
probelock probe --tools tools.json --via anyllm --model mistral/mistral-large-latest \
--samples 5 --temperature 0.7 -o candidate.lock
--via clients reuse the same caching, sampling, and error semantics as
--endpoint ; they are thin adapters over each SDK's OpenAI-shaped response. Add a
new backend by implementing the tiny Client protocol — that is the only seam.
demo/ has runs against a local Ollama server: a committed qwen3.5:9b
baseline vs a gemma3:1b candidate (which does not support tool-calling). See
demo/DEMO.md for the transcript, or replay it:
asciinema play demo/probelock-demo.cast # or: bash demo/demo.sh
The tool-calling capabilities drop 1.00 → 0.00 and the gate exits non-zero.
tool_restraint , tool_permission , and no_hallucinated_tool stay 1.00 (a
model that cannot call tools cannot misuse one), and gemma3:1b scores 1.00 on
format_adherence vs 0.50 for qwen3.5:9b . The diff is per-capability.
Also committed: qwen3.5:9b vs lfm2.5-thinking:1.2b :
uv run probelock diff demo/qwen3.5-9b.lock demo/lfm2.5-thinking.lock
The 1.2B model matches qwen3.5:9b on tool selection, discrimination,
needle_in_tools , arg_validity , required_args , and the three safety probes;
structured_output and arity_robustness drop 1.00 → 0.33 .
The capabilities (all scored deterministically)
Capability
What it checks
Scorer
tool_selection
Calls the right tool for the task
tool-name match
tool_discrimination
Calls the right tool and no other (picks precisely)
tool-name set
needle_in_tools
Finds the right tool when many (15+) are offered
tool-name match
arg_validity
Emitted args validate against the tool's JSON schema
jsonschema
required_args
All required args present and non-empty
key presence
arity_robustness
Fills every parameter (required + optional) when asked
all-present
structured_output
Emits schema-valid JSON on demand (no tools, no fences)
parse + jsonschema
json_mode (opt-in, --json-mode )
Same, but via the server's native response_format API instead of a prompt
parse + jsonschema
tool_restraint
Does not call a tool for a task that needs none (over-trigger)
no tool call
tool_permission
Does not call a tool it was explicitly forbidden to use
forbidden tool absent
no_hallucinated_tool
Does not fabricate a call to a tool that was not offered
called ⊆ offered
format_adherence
Follows an exact output constraint
exact match
Three are negative probes (a higher score means the bad behavior did not
happen): tool_restraint (over-triggering), tool_permission (calling a forbidden
tool), and no_hallucinated_tool (fabricating a tool). All probes are derived from
the tool schemas, not hand-authored.
tool schemas ──▶ derive probes ──▶ Client ──▶ ResponseMessage ──▶ deterministic scorer ──▶ Lockfile
(your agent) (zero authoring) (model) (the only model (no LLM judge) (commit it)
-touching part)
Lockfile + Lockfile ──▶ diff / gate
The only nondeterministic part is the Client ; everything else is pure, so the
same inputs produce the same lockfile and the same diff. At temperature 0 the
client caches identical requests, so the probes that share one request (the tool
checks for a given tool) hit the network once. The SimulatedClient crafts correct or
incorrect responses that the real scorers grade, so the scoring path runs even
with no model present.
Deriving probes from real traces (experimental)
Schema-derived probes are single-turn and synthetic — great for catching schema-level
regressions, blind to what breaks after several turns of real context, a tool result
feeding back in, or ambiguous phrasing. --traces adds a second source: real,
already-recorded agent decision points (e.g. exported from litellm's OpenTelemetry
callback), replayed through the exact same deterministic scorers.
uv run probelock derive --tools tools.json --traces traces.json # see what gets added
uv run probelock probe --tools tools.json --traces traces.json \
--endpoint http://localhost:11434/v1 --model llama3.1:8b -o candidate.lock
--tools is optional here: traced probes replay their own embedded tool definitions,
so a trace-only run ( probe --traces traces.json ... ) needs no schema file. The same
holds for --mined below. Provide --tools when you also want the synthetic battery.
A traces file is a small, stable JSON schema probelock defines itself — not raw
OpenTelemetry — because OTel's own span attribute layout is not stable across libraries or
versions (litellm has already changed where it puts request/response attributes once, and
has a newer, differently-shaped opt-in integration). Converting your export into this shape
is a one-time step you own; see
examples/otel_traces_to_probelock.py for a
documented starting point and fixtures/sample_traces.json
for the target shape:
{
"version" : 1 ,
"records" : [
{
"id" : " checkout-flow-turn-3 " ,
"messages" : [{ "role" : " user " , "content" : " ... " }],
"tools" : [ /* OpenAI-style tool defs actually offered at this turn */ ],
"response" : { "content" : null , "tool_calls" : [{ "name" : " ... " , "arguments" : " {...} " }]}
}
]
}
Trace-derived probes join the same capabilities as schema-derived ones — tool_selection ,
tool_discrimination , arg_validity , required_args , and structured_output — since these
map cleanly onto "replay this real context, check the candidate still behaves validly" (probe
ids c

[truncated]
