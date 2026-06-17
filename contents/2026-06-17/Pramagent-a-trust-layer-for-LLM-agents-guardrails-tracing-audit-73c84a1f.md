---
source: "https://github.com/sriram7737/pramagent"
hn_url: "https://news.ycombinator.com/item?id=48574486"
title: "Pramagent – a trust layer for LLM agents (guardrails, tracing, audit)"
article_title: "GitHub - sriram7737/pramagent · GitHub"
author: "sriram7737"
captured_at: "2026-06-17T18:56:26Z"
capture_tool: "hn-digest"
hn_id: 48574486
score: 1
comments: 0
posted_at: "2026-06-17T18:22:17Z"
tags:
  - hacker-news
  - translated
---

# Pramagent – a trust layer for LLM agents (guardrails, tracing, audit)

- HN: [48574486](https://news.ycombinator.com/item?id=48574486)
- Source: [github.com](https://github.com/sriram7737/pramagent)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T18:22:17Z

## Translation

タイトル: Pramagent – LLM エージェントの信頼層 (ガードレール、トレース、監査)
記事タイトル: GitHub - sriram7737/pramagent · GitHub
説明: GitHub でアカウントを作成して、sriram7737/pramagent の開発に貢献します。

記事本文:
GitHub - sriram7737/pramagent · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
スリラム7737
/
プラマジェント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらに開く アクション メニュー フォルダー

dファイル
83 コミット 83 コミット .github/ ワークフロー .github/ ワークフロー デプロイ デプロイ ドキュメント ドキュメントの例 例 pramagent pramagent テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile ライセンス ライセンスMANIFEST.in MANIFEST.in Procfile Procfile README.md README.md docker-compose.yml docker-compose.yml llms.txt llms.txt pramagent_enterprise_audit.md pramagent_enterprise_audit.md pramagent_full_audit.md pramagent_full_audit.md pramagent_security_test_results.md pramagent_security_test_results.md pyproject.toml pyproject.toml Railway.toml Railway.toml test_agent.py test_agent.py test_agent_v2.py test_agent_v2.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM エージェント用の信頼ミドルウェア: 決定論的なツール ポリシー、HITL 承認、
改ざんの明らかな監査トレース。アルファ - 読む
実施状況
顧客対応パイロットの前に。
Pramagent は、OpenAI、Anthropic、Gemini、Ollama、local、および
モデルの外側で実行されるガードレールを備えた OpenAI 互換プロバイダー。の
最も差別化されたレイヤーは ToolGuard です。JSON を使用した決定論的なツール検証
スキーマ、テナント/アクション許可リスト、副作用分類、危険なチェーン
検出、出力スキャン、および HITL エスカレーション。現在のパッケージも
厳選された安全ルールコーパス、永続的な HITL キュー、シンアダプターを出荷
一般的なエージェント フレームワークとコンプライアンス証拠の生成。
Pramagent は Alpha ソフトウェアとして公開されています。生煙試験の証拠がある
Sepolia アンカリング、S3 コールド アーカイブ、ローカル負荷テスト、実際の OpenAI/Ollama 用
プロバイダーが呼び出し、バンドルされたレッドチームが実行されますが、
外部侵入テスト、SOC 2 監査、HIPAA 評価、または
規制生産認証。
Pramagent を銀行グレードまたはヘルスケアグレードの SE として扱わないでください

正確さ
インフラストラクチャ。即時注射に対する耐性、製造コンプライアンス、
または、バンドルされたベンチマークのみからサードパーティによって検証された安全性。読む
実施状況、
ライブテスト結果、および
硬化ガイド
顧客向けのパイロットで使用する前に。
6 月 11 日のアクティブなセキュリティ プロンプトの結果は、次の場所で追跡されます。
セキュリティテストの結果。
これは基本パッケージでのみ機能します。 Docker、API サーバー、プロバイダー キーはありません。
必須です。
pip インストール pramagent
非同期をインポートする
プラマジェントからのインポート プラマジェント
非同期デフォルトメイン():
resp = プラマジェントを待ちます ()。 run ( "このリクエストを要約する" 、 tenant_id = "demo" 、 session_id = "s1" )
print (それぞれ出力)
print (それぞれ .trace .this_hash )
非同期 。実行 (メイン ())
これにより、決定論的なモック プロバイダーを使用して改ざん明示トレースが作成されます。
OPENAI_API_KEY を設定して、実際の OpenAI モデルに切り替えます。
プラマジェントからのインポート プラマジェント
プラマジェントから。プロバイダーは OpenAIProvider をインポートします
アーマー = Pramagent (プロバイダー = OpenAIProvider (モデル = "gpt-4o-mini"))
nvapi-* キーを使用して NVIDIA NIM に対して実行します。
プラマジェントからのインポート プラマジェント
プラマジェントから。プロバイダーが NvidiaProvider をインポートする
アーマー = Pramagent (プロバイダー = NvidiaProvider (モデル = "meta/llama-3.3-70b-instruct"))
よくある質問
LLM エージェントに安全ガードレールを追加するにはどうすればよいですか?
Pramagent をインストールし、エージェント呼び出しを信頼スタックでラップします。プラマジェント
モデルの外部で決定論的なポリシーを強制するため、LLM は
ツール ポリシー、HITL ゲート、または監査チェーン自体のテキスト出力を変更します。
本番環境での AI エージェントの決定を監査するにはどうすればよいですか?
すべての Pramagent 呼び出しでは、レイヤーの決定を伴うハッシュチェーンされた TraceEvent が生成されます。
判定、プロバイダーのメタデータ、PII 編集、HITL ステータス、および this_hash /
prev_hash 。ローカル チェーンは検証でき、オプションで外部に固定できます。
どうやって

Python LLM エージェントでのプロンプト インジェクションを禁止しますか?
IsolationLayer は、モデルが入力を認識する前に入力をスキャンします。既知の内容をカバーします
命令オーバーライド、チャットテンプレートラッパー攻撃、権限フレーミング、
Base64/hex/unicode エスケープでエンコードされたペイロード、およびターゲットを絞った多言語オーバーライド
フレーズ。 v0.8.0 は、構造化分類子判定、ホールドアウト PINT/TensorTrust を追加します。
スタイルフィクスチャ、来歴を意識したツール出力のより厳密なスキャン、および
取得されたコンテンツ、およびオプションの pramagent[ml] 埋め込み/DeBERTa レイヤー。これ
これは多層防御であり、即時注射に対する免疫の証明ではありません。
安全でないモデルの出力がユーザーに届かないようにするにはどうすればよいですか?
OutputJudgeLayer は、出力が返される前に、すべての出力に対して LLM-as-judge を実行します。
「OUTPUTは安全ですか？」正規表現が与えられないことを確認してください。セマンティックエラーをキャッチします
決定的ルールのミス (動作するマルウェア、ウォークスルーのバイパス、確認済み)
破壊行為、内部情報の漏洩）。公開デモではデフォルトでオンになっており、オプトイン
/v1/run の場合 (PRAMAGENT_OUTPUT_JUDGE=1)。フェールクローズされますが、それ自体は
モデル — 強力な多層防御ですが、保証するものではありません。
AI エージェントからの安全でないツールの呼び出しを停止するにはどうすればよいですか?
ToolPolicy とともに ToolGuardLayer を使用します。 Pramagent は JSON スキーマを検証します。
テナント/アクション許可リスト、副作用クラス、呼び出し頻度、引数
副作用が発生する前に、注入と危険な連鎖が発生します。
AI エージェントのアクションに人間の承認を追加するにはどうすればよいですか?
Verdict.ESCALATE で HITLLayer または ToolGuard ポリシーを使用します。沈黙は決してない
同意: 承認が得られない場合、アクションは実行されないままになります。
Pramagent は OpenAI、Anthropic、Gemini、Ollama、およびローカル モデルで動作しますか?
はい。 Pramagent は、OpenAI、Anthropic、Gemini、Ollama、
NVIDIA NIM、OpenAI 互換のローカル エンドポイント、および決定論的モック
テスト用のプロバイダー。
プラマジェントは以下に準拠していますか

SOC 2、HIPAA、または EU AI 法?
いいえ。Pramagent には、コンプライアンス証拠マッピングと改ざん明示ロギングが含まれています
評価をサポートできる機能はありますが、SOC 2、HIPAA、EU には合格していません
AI 法適合性評価、または外部侵入テスト。
pip install " pramagent[api,dashboard,redis,postgres] "
ソースより:
git clone git@github.com:sriram7737/pramagent.git
cd プラマジェント
pip install -e " .[dev,api,redis,postgres,dashboard] "
CLI と Docker のクイックスタート
プラマジェントの初期化
プラマジェント検証
ローカル スタックを実行します。
cp .env.example .env
ドッカー構成 -d
開く:
API ドキュメント: http://localhost:8080/docs
ダッシュボード: http://localhost:8501
API は、 /demo で単一ページの NVIDIA NIM デモを提供できます。無効になっているのは、
デフォルトであり、運用トラフィックではなく公開評価を目的としています。
PRAMAGENT_DEMO_ENABLED=true
PRAMAGENT_DEMO_RATE_LIMIT=60
PRAMAGENT_ALLOW_MEMORY_STORE=1
uvicorn pramagent.api.app:app --host 0.0.0.0 --port 8080
デモでは、実行のたびに訪問者に独自の nvapi-* キーを要求します。プラマジェント
そのキーは現在のプロバイダー呼び出しにのみ使用されます。トレースには書き込まれませんが、
ログ、ストア、使用記録、またはハッシュ チェーン ペイロード。各デモの実行では、
分離されたメモリ内トレースを保存し、出力、信頼層イベントを返します。
リダクション、HITL 状態、レイテンシー、 this_hash 、 prev_hash 、およびローカル チェーン
検証。
パブリック スロットルは、クライアント IP と短いメモリ内 SHA-256 ハッシュによってキー設定されます。
訪問者の nvapi-* キー。訪問者が別の NVIDIA に切り替えた場合
キーを使用すると、Pramagent が平文キーを保存せずに、新しいデモ バケットを取得できます。
DEGRADED デモ結果は、アップストリーム モデルの呼び出しが失敗し、Pramagent が失敗したことを意味します。
トレース付きで安全なデフォルトを返しました。リストされている別の NIM モデルを試すか、確認してください
キーが選択したエンドポイントにアクセスできることを確認します。
リリースの健全性チェックを実行します。
Python -mp

ytest -q --tb=no
python -m pramagent.cli redteam --json --攻撃 100
python -m pramagent.cli redteam --json --dynamic --攻撃 200 --seed 999
現在のローカル結果: 640 が合格、2 がスキップされました。最新のターゲットを絞ったプロンプト
このスイートも、緊急オーバーライド、出力オーバーライド、
マージン/清算、IBAN/SWIFT、あいまいなエスカレーション、PHI、誤検知、
Base64、16 進数、Unicode エスケープ、多言語オーバーライド トークン、および
チャットテンプレートラッパーのケース。
非同期をインポートする
from pramagent import Pramagent 、評決
プラマジェントから。レイヤーインポート ToolGuardLayer 、 ToolPolicy
プラマジェントから。レイヤー。 tool_guard インポート SideEffect
ガード = ToolGuardLayer ( ポリシー = [
ツールポリシー (
名前 = "send_payment" ,
Side_effect = サイドエフェクト。お支払い、
アクション = 評決。エスカレーション、
allowed_tenants = { "finance_team" },
スキーマ = {
"タイプ" : "オブジェクト" ,
"必須" : [ "金額_米ドル" , "目的地" ],
"プロパティ" : {
"金額_米ドル" : { "タイプ" : "数値" 、 "最小" : 0.01 、 "最大" : 5000 },
"宛先" : { "タイプ" : "文字列" , "パターン" : r"acct-\d{6,}" },
}、
"追加プロパティ" : False 、
}、
)
])
アーマー = プラマジェント (tool_guard = ガード)
非同期デフォルトメイン():
決定＝鎧。 validate_tool (
"send_payment" ,
{ "金額_米ドル" : 250.00 、 "目的地" : "acct-123456" },
tenant_id = "finance_team" ,
session_id = "デモ" ,
)
print (決定.評決) #ESCALATE
too_large = 鎧。 validate_tool (
"send_payment" ,
{ "金額_米ドル" : 9000.00 、 "目的地" : "acct-123456" },
tenant_id = "finance_team" ,
session_id = "デモ" ,
)
print ( too_large . verdict , too_large .reason ) # ブロック: スキーマ違反
間違ったテナント = アーマー。 validate_tool (
"send_payment" ,
{ "金額_米ドル" : 250.00 、 "目的地" : "acct-123456" },
tenant_id = "マーケティングチーム" ,
session_id = "デモ" ,
)
print ( 間違ったテナント . verdi

ct 、間違ったテナント。理由 ) # ブロック: テナントの不一致
応答 = 鎧を待ちます。走って（
"この支払いリクエストを要約してください" ,
tenant_id = "finance_team" ,
session_id = "デモ" ,
アクション = "send_payment" ,
)
print (応答.hitl)
print (応答.トレース.this_hash)
非同期 。実行 (メイン ())
組み込みルールコーパス
Pramagent には、決定的でインポート可能なルール バンドルが含まれるようになりました。彼らは地味だ
Python Rule オブジェクトを使用するため、レビュー担当者は何が施行されているかを正確に検査できます。
プラマジェントからのインポート プラマジェント
プラマジェントから。レイヤーはSafetyLayerをインポートします
プラマジェントから。ルールインポート ALL_RULES 、 JAILBREAK_PATTERNS 、 OWASP_LLM_TOP10
鎧 = プラマジェント (
安全性 = SafetyLayer (ルール = [ * JAILBREAK_PATTERNS , * OWASP_LLM_TOP10 ])
)
strict_armor = Pramagent (安全性 = SafetyLayer (ルール = ALL_RULES))
含まれるコーパス:
Verdict.ESCALATE は、「疑わしいが、ブロックするほど確実ではない」ことを意味します。何
パイプラインはステージごとに構成可能です — プレ (入力パス、
モデルの実行前）および事後（出力パス後） — 次のいずれかを使用します。
「log」（記録して続行）、「hitl」（人間参加型ゲートへのルート、
アイドルオンサイレンス)、または「ブロック」(ハードストップ)。デフォルトは「log」なので、
ESCALATE ルールがサイレントにゲート制御を開始することはありません

[切り捨てられた]

## Original Extract

Contribute to sriram7737/pramagent development by creating an account on GitHub.

GitHub - sriram7737/pramagent · GitHub
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
sriram7737
/
pramagent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
83 Commits 83 Commits .github/ workflows .github/ workflows deploy deploy docs docs examples examples pramagent pramagent tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE MANIFEST.in MANIFEST.in Procfile Procfile README.md README.md docker-compose.yml docker-compose.yml llms.txt llms.txt pramagent_enterprise_audit.md pramagent_enterprise_audit.md pramagent_full_audit.md pramagent_full_audit.md pramagent_security_test_results.md pramagent_security_test_results.md pyproject.toml pyproject.toml railway.toml railway.toml test_agent.py test_agent.py test_agent_v2.py test_agent_v2.py View all files Repository files navigation
Trust middleware for LLM agents: deterministic tool policy, HITL approvals,
and tamper-evident audit traces. Alpha - read the
implementation status
before customer-facing pilots.
Pramagent wraps OpenAI, Anthropic, Gemini, Ollama, local, and
OpenAI-compatible providers with guardrails that run outside the model. The
most differentiated layer is ToolGuard: deterministic tool validation with JSON
Schema, tenant/action allow-lists, side-effect taxonomy, dangerous-chain
detection, output scanning, and HITL escalation. The current package also
ships curated safety rule corpora, persistent HITL queues, thin adapters for
popular agent frameworks, and compliance evidence generation.
Pramagent is published as Alpha software . It has live smoke-test evidence
for Sepolia anchoring, S3 cold archive, local load testing, real OpenAI/Ollama
provider calls, and bundled red-team runs, but it has not passed an
external penetration test, SOC 2 audit, HIPAA assessment, or
regulated-production certification.
Do not treat Pramagent as bank-grade or healthcare-grade security
infrastructure. Do not claim prompt-injection immunity, production compliance,
or third-party-validated safety from the bundled benchmarks alone. Read
Implementation status ,
Live test results , and
Hardening guide
before using it in a customer-facing pilot.
The June 11 active security prompt results are tracked in
Security test results .
This works with the base package only. No Docker, API server, or provider key is
required.
pip install pramagent
import asyncio
from pramagent import Pramagent
async def main ():
resp = await Pramagent (). run ( "Summarize this request" , tenant_id = "demo" , session_id = "s1" )
print ( resp . output )
print ( resp . trace . this_hash )
asyncio . run ( main ())
That creates a tamper-evident trace using the deterministic mock provider.
Swap to a real OpenAI model by setting OPENAI_API_KEY :
from pramagent import Pramagent
from pramagent . providers import OpenAIProvider
armor = Pramagent ( provider = OpenAIProvider ( model = "gpt-4o-mini" ))
Run against NVIDIA NIM with an nvapi-* key:
from pramagent import Pramagent
from pramagent . providers import NvidiaProvider
armor = Pramagent ( provider = NvidiaProvider ( model = "meta/llama-3.3-70b-instruct" ))
Frequently Asked Questions
How do I add safety guardrails to an LLM agent?
Install Pramagent and wrap your agent call with the trust stack. Pramagent
enforces deterministic policy outside the model, so the LLM cannot override the
tool policy, HITL gate, or audit chain by changing its own text output.
How do I audit AI agent decisions in production?
Every Pramagent call produces a hash-chained TraceEvent with layer decisions,
verdicts, provider metadata, PII redactions, HITL status, and this_hash /
prev_hash . The local chain can be verified and optionally anchored externally.
How do I prevent prompt injection in a Python LLM agent?
IsolationLayer scans inputs before the model sees them. It covers known
instruction overrides, chat-template wrapper attacks, authority framing,
base64/hex/unicode-escape encoded payloads, and targeted multilingual override
phrases. v0.8.0 adds structured classifier verdicts, held-out PINT/TensorTrust
style fixtures, provenance-aware stricter scanning for tool output and
retrieved content, and optional pramagent[ml] embedding/DeBERTa layers. This
is defense-in-depth, not proof of prompt-injection immunity.
How do I stop unsafe model output from reaching users?
OutputJudgeLayer runs an LLM-as-judge on every output before it returns — the
"is the OUTPUT safe?" check that regex cannot give. It catches semantic failures
deterministic rules miss (working malware, bypass walkthroughs, confirmed
destructive actions, leaked internals). On by default in the public demo, opt-in
for /v1/run ( PRAMAGENT_OUTPUT_JUDGE=1 ). It is fail-closed, but it is itself a
model — strong defense-in-depth, not a guarantee.
How do I stop unsafe tool calls from an AI agent?
Use ToolGuardLayer with ToolPolicy . Pramagent validates JSON Schema,
tenant/action allow-lists, side-effect class, call frequency, argument
injection, and dangerous chains before any side effect can execute.
How do I add human approval to AI agent actions?
Use HITLLayer or a ToolGuard policy with Verdict.ESCALATE . Silence is never
consent: if approval does not arrive, the action remains unexecuted.
Does Pramagent work with OpenAI, Anthropic, Gemini, Ollama, and local models?
Yes. Pramagent ships provider adapters for OpenAI, Anthropic, Gemini, Ollama,
NVIDIA NIM, and OpenAI-compatible local endpoints, plus a deterministic mock
provider for tests.
Is Pramagent compliant with SOC 2, HIPAA, or the EU AI Act?
No. Pramagent includes compliance evidence mapping and tamper-evident logging
features that can support an assessment, but it has not passed SOC 2, HIPAA, EU
AI Act conformity assessment, or an external penetration test.
pip install " pramagent[api,dashboard,redis,postgres] "
From source:
git clone git@github.com:sriram7737/pramagent.git
cd Pramagent
pip install -e " .[dev,api,redis,postgres,dashboard] "
CLI And Docker Quickstart
pramagent init
pramagent validate
Run the local stack:
cp .env.example .env
docker compose up -d
Open:
API docs: http://localhost:8080/docs
Dashboard: http://localhost:8501
The API can serve a single-page NVIDIA NIM demo at /demo . It is disabled by
default and is meant for public evaluation, not production traffic.
PRAMAGENT_DEMO_ENABLED=true
PRAMAGENT_DEMO_RATE_LIMIT=60
PRAMAGENT_ALLOW_MEMORY_STORE=1
uvicorn pramagent.api.app:app --host 0.0.0.0 --port 8080
The demo asks the visitor for their own nvapi-* key on each run. Pramagent
uses that key only for the current provider call; it is not written to traces,
logs, stores, usage records, or the hash-chain payload. Each demo run uses an
isolated in-memory trace store and returns the output, trust-layer events,
redactions, HITL state, latency, this_hash , prev_hash , and local chain
verification.
The public throttle is keyed by client IP plus a short in-memory SHA-256 hash
of the visitor's nvapi-* key. If a visitor switches to a different NVIDIA
key, they get a fresh demo bucket without Pramagent storing the plaintext key.
A DEGRADED demo result means the upstream model call failed and Pramagent
returned its safe default with a trace; try another listed NIM model or verify
that the key has access to the selected endpoint.
Run the release sanity checks:
python -m pytest -q --tb=no
python -m pramagent.cli redteam --json --attacks 100
python -m pramagent.cli redteam --json --dynamic --attacks 200 --seed 999
Current local result: 640 passed, 2 skipped . The latest targeted prompt
suite also passed with 0 failures across emergency override, output override,
margin/liquidation, IBAN/SWIFT, ambiguous escalation, PHI, false-positive,
base64, hex, unicode-escape, multilingual override-token, and
chat-template-wrapper cases.
import asyncio
from pramagent import Pramagent , Verdict
from pramagent . layers import ToolGuardLayer , ToolPolicy
from pramagent . layers . tool_guard import SideEffect
guard = ToolGuardLayer ( policies = [
ToolPolicy (
name = "send_payment" ,
side_effect = SideEffect . PAYMENT ,
action = Verdict . ESCALATE ,
allowed_tenants = { "finance_team" },
schema = {
"type" : "object" ,
"required" : [ "amount_usd" , "destination" ],
"properties" : {
"amount_usd" : { "type" : "number" , "minimum" : 0.01 , "maximum" : 5000 },
"destination" : { "type" : "string" , "pattern" : r"acct-\d{6,}" },
},
"additionalProperties" : False ,
},
)
])
armor = Pramagent ( tool_guard = guard )
async def main ():
decision = armor . validate_tool (
"send_payment" ,
{ "amount_usd" : 250.00 , "destination" : "acct-123456" },
tenant_id = "finance_team" ,
session_id = "demo" ,
)
print ( decision . verdict ) # ESCALATE
too_large = armor . validate_tool (
"send_payment" ,
{ "amount_usd" : 9000.00 , "destination" : "acct-123456" },
tenant_id = "finance_team" ,
session_id = "demo" ,
)
print ( too_large . verdict , too_large . reason ) # BLOCK: schema violation
wrong_tenant = armor . validate_tool (
"send_payment" ,
{ "amount_usd" : 250.00 , "destination" : "acct-123456" },
tenant_id = "marketing_team" ,
session_id = "demo" ,
)
print ( wrong_tenant . verdict , wrong_tenant . reason ) # BLOCK: tenant mismatch
response = await armor . run (
"Summarize this payment request" ,
tenant_id = "finance_team" ,
session_id = "demo" ,
action = "send_payment" ,
)
print ( response . hitl )
print ( response . trace . this_hash )
asyncio . run ( main ())
Built-In Rule Corpora
Pramagent now includes deterministic, importable rule bundles. They are plain
Python Rule objects, so a reviewer can inspect exactly what is enforced.
from pramagent import Pramagent
from pramagent . layers import SafetyLayer
from pramagent . rules import ALL_RULES , JAILBREAK_PATTERNS , OWASP_LLM_TOP10
armor = Pramagent (
safety = SafetyLayer ( rules = [ * JAILBREAK_PATTERNS , * OWASP_LLM_TOP10 ])
)
strict_armor = Pramagent ( safety = SafetyLayer ( rules = ALL_RULES ))
Included corpora:
Verdict.ESCALATE means "suspicious, but not certain enough to block." What
the pipeline does with it is configurable per stage — pre (the input pass,
before the model runs) and post (the output pass, after) — with one of
"log" (record and continue), "hitl" (route to the human-in-the-loop gate,
idle-on-silence), or "block" (hard stop). The default is "log" so adding an
ESCALATE rule never silently starts gating t

[truncated]
