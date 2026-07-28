---
source: "https://github.com/benchwire/labwire"
hn_url: "https://news.ycombinator.com/item?id=49087756"
title: "Labwire, an open protocol for AI agents to control lab instruments"
article_title: "GitHub - benchwire/labwire · GitHub"
author: "RoboSy"
captured_at: "2026-07-28T19:08:50Z"
capture_tool: "hn-digest"
hn_id: 49087756
score: 1
comments: 0
posted_at: "2026-07-28T18:11:36Z"
tags:
  - hacker-news
  - translated
---

# Labwire, an open protocol for AI agents to control lab instruments

- HN: [49087756](https://news.ycombinator.com/item?id=49087756)
- Source: [github.com](https://github.com/benchwire/labwire)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T18:11:36Z

## Translation

タイトル: Labwire、AI エージェントが実験器具を制御するためのオープン プロトコル
記事タイトル: GitHub - ベンチワイヤー/ラボワイヤー · GitHub
説明: GitHub でアカウントを作成して、ベンチワイヤー/ラボワイヤーの開発に貢献します。

記事本文:
GitHub - ベンチワイヤー/ラボワイヤー · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ベンチワイヤー
/
実験用ワイヤー
公共
通知
通知設定を変更するにはサインインする必要があります
追加

アルナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
41 コミット 41 コミット .github .github docs docs サンプル サンプル パッケージ パッケージ 仕様 仕様 テスト テスト .gitignore .gitignore .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE Makefile Makefile 注意通知 PRIOR_ART.md PRIOR_ART.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SPEC-FINDINGS.md SPEC-FINDINGS.md pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json uv.lock uv.lock すべてのファイルを表示リポジトリ ファイルのナビゲーション
AI 制御の実験器具のためのオープン プロトコル。 「MCP」を考えてみましょう
実験装置」: AI エージェントが発見するための普遍的な方法の 1 つ
機器の機能、コマンド、測定値のストリーミング、
何が行われたかを示す暗号署名された証拠を持って立ち去ります。
仮タイトル、プロトコル v0.3 ドラフト。ワイヤプロトコルは前に変更されます
1.0。フィードバックや従来技術の修正は大歓迎です。参照してください
貢献.md 。
自動運転ラボには、エージェントが安全に操作できる機器が必要です。
聴覚的に。現在、各ベンダーは異なる方言を話しており、
統合はオーダーメイドです。 Labwire の賭けは、欠けている部分は小さく、
今すぐ構築可能:
AI エージェントネイティブ。 MCP をモデルとした機能検出: 手段
コマンドを JSON スキーマとして記述しているため、任意のエージェント フレームワークでそれを駆動できます。
グルー コードはゼロです。バンドルされている MCP アダプターがそれを証明しています。
署名された結果。実行するたびに、ed25519 で署名されたマニフェストが生成される可能性があります。
記録された正確なテレメトリー: ポータブルで改ざんが明らかな証拠
機器が何をしたかを 1 つの CLI コマンドで検証しました。
プロトコルにおける安全性と物理的なタイピング。必須の UCUM ユニット
すべての数量、S0 ～ S3 安全クラス、ここで irrev

ersible actions take an
オペレータの確認と危険なものはオペレータの承諾を得る
mint はできず、正確なパラメータにバインドされ、インターロック、キャンセル、および
再試行可能な型付きエラー。ベンダーのアドオンではなく、すべて指定されています。
量だけではない、物。 v0.3 はリソース (型付き、読み取り可能) を追加します
液体ハンドラーのデッキなどの機器の状態) と型付き参照
(井戸またはサイトに名前を付けるパラメータ。現在の状態に対して検証され、
エラーが発生すると、回復する読み取りがエージェントに渡されます)。 Designed so that
発見だけがエージェントをデッキに導き、迅速なコーチングはありません。 CI
前提条件を強制し、デモで動作を主張します。
見知らぬ人でも5分以内に走れます。ゼロハードウェア: リファレンス
この実装には、3 つの現実的なシミュレートされた計測器が同梱されています。
PyPI (Python 3.12+) からは、チェックアウトは必要ありません。
pip install labwire
これにより、SDK、3 つのシミュレートされたインストゥルメント、それらのドライバー、および
labwire CLI.インストゥルメントを宣言し、エンドツーエンドで直接実行します。
a Python file:
非同期をインポートする
from labwire . core import (
PROTOCOL_VERSION ,
CommandContext ,
アイデンティティ情報 、
Instrument ,
インストゥルメントサーバー 、
LabwireClient ,
MemoryTransport ,
command ,
)
pydantic import BaseModel 、 ConfigDict から
クラス MassReading ( BaseModel ):
model_config = ConfigDict ( extra = "forbid" ) # 閉じられたスキーマ: ユニットウォーカーがそれを要求します
mass_g : float
クラス バランス (楽器):
"""ワンコマンドの手段。単位は必須です: "g" を省略すると宣言が拒否されます。"""
アイデンティティ = アイデンティティ情報 (
メーカー = "あなた" 、モデル = "Balance-1" 、シリアル番号 = "B-1" 、ファームウェア バージョン = "1.0"
)
@ コマンド (returns_units = { "mass_g" : "g" })
async def meter (self , ctx : CommandContext ) -> MassReading :
"""定住した集団を報告してください。"""
MassReading を返します (mass_g = 12.3456)
として

ync def main () -> なし :
サーバー = InstrumentServer ( Balance ())
client_end 、server_end = MemoryTransport 。ペア（）
サーバー。アタッチ (サーバーエンド)
LabwireClient と非同期。アタッチ ( client_end ) をクライアントとして:
記述子 = クライアントを待ちます。説明する()
print (記述子.アイデンティティ.モデル、「Labwireプロトコルを話す」、PROTOCOL_VERSION)
ハンドル = クライアントを待ちます。送信 ( "測定" , {})
result = ハンドルを待ちます。結果 (タイムアウト = 5.0)
print ( "mass:" , result [ "mass_g" ] , "g" )
非同期 。実行 (メイン ())
テレメトリーストリーミング、安全確認、および ed25519 署名付き実行バンドル
あと数行です。例/クイックスタート.py
それらを示します。
完全なデモと例についてはソースから:
git clone https://github.com/benchwire/labwire.git && cd labwire
make setup # uv は Python 3.12 とすべてをインストールします
uv run Examples/quickstart.py # 60 秒: シミュレートされたバランスをエンドツーエンドで駆動する
uv run example/streaming.py # テレメトリ、キャンセル、インターロック回復
デモを作成 # 閉ループ最適化 + 署名された証拠
make Demon は完全な自律型実験キャンペーンを実行します: スクリプト化されたオプティマイザー
3 つのシミュレートされたヒーター電圧と試薬流量を調整します
計測器を使用し、隠れた最適収量に収束し、
ウイニングランのサイン入りバンドル:
安全性: ポンプのディスペンスはクラス S2 (不可逆的) です。オペレータの継続的な助成金の下で実行されます
実行 13 V= 15.0 V -> T= 68.0 ℃ q= 127 μL/分 収率= 87.2% 最良= 87.2%
収束: 14 回の実験で 15.0 V (68.0 ℃)、127 uL/min で最高収率 87.2%
署名された証拠：demo_runs/d3b15e9f-...
labwire 検証: OK - 本物です
make Demon-claude は、実際の Claude エージェントの計画と同じループを実行します。
機器のツール スキーマによる実験 (必要な
ANTHROPIC_API_KEY ;スクリプト化されたオプティマイザに正常に機能を低下させます。
それ）。
フローチャート LR
部分グラフの年齢

nts [エージェント]
クロード[クロード / 任意の MCP クライアント]
スクリプト[オプティマイザー / LabwireClient]
終わり
アダプター["labwire-mcp<br>(MCPアダプター)"]
サブグラフサーバー [計測サーバー - labwire-core]
psu[電源ドライバー]
ポンプ[シリンジポンプドライバー]
bal[バランスドライバー]
終わり
サブグラフ デバイス [ネイティブ ワイヤ プロトコル - labwire-sim]
scpi["SimPSU-3005<br>SCPI over TCP"]
シリアル["SimPump-200<br>シリアルスタイルのライン"]
stream["SimBalance-120<br>ストリーミング測定値"]
終わり
verify["labwire verify<br>(ed25519 + RFC 8785)"]
クロード -->|MCP ツール|アダプター
アダプター -->|JSON-RPC / WebSocket|サーバー
スクリプト -->|検出/コマンド/ストリーム|サーバー
psu --> scpi
ポンプ --> シリアル
バル --> ストリーム
サーバー -->|署名された実行バンドル|検証する
読み込み中
プロトコルは WebSocket (stdio 指定) 経由の JSON-RPC 2.0 です。
MCP からインスピレーションを得た初期化/機能ハンドシェイク、プッシュファースト コマンド ライフサイクル、
シーケンスされたテレメトリー、プロトコルレベルの安全インターロック、および署名された規範
マニフェストを実行します。完全な仕様は spec/SPEC.md にあります。
その中のすべての JSON サンプルは、実装に対して機械検証されます。
CIで。
パッケージ
それは何ですか
ラボワイヤーコア
サーバー + クライアント SDK、トランスポート、セッション層、署名、JCS
ラボワイヤーシム
ネイティブ ワイヤ プロトコルを話す 3 つのリアルなシミュレートされた機器
ラボワイヤードライバー
これらのネイティブ プロトコルを Labwire 計測器としてラップするドライバー
ラボワイヤー-MCP
MCP アダプター: すべての機器コマンドが MCP ツールになります
ラボワイヤー-cli
labwire verify <bundle> : 署名された実行証拠を認証します
ラボワイヤーオフィド
ブリッジ: 任意の ophyd (Bluesky) デバイスを Labwire 機器として機能させます
labwire-pylabrobot
ブリッジ: PyLabRobot リキッド ハンドラーを Labwire 機器として機能させます
スペック/
プロトコル仕様 (v0.2 ドラフト)
例/
クイックスタート、ストリーミング/リカバリ、クローズドループデモ
独自のデバイスをラップするのはクラスとデコレータです。
私のクラス

ポンプ (器具):
アイデンティティ = IdentityInfo (メーカー = "あなた" 、モデル = "ポンプ-1" 、
シリアル番号 = "001" 、ファームウェアバージョン = "1.0" )
flow = channel ( "flow_rate" 、unit = "uL/min" ) # UCUM コードは必須です
@コマンド(
単位 = { "volume_ul" : "uL" , "rate_ul_min" : "uL/min" },
returns_units = { "dispensed_ul" : "uL" },
safety_class = "S2" , # 不可逆: 確認が必要
)
async def dispense ( self 、 ctx : CommandContext 、
volume_ul : float 、 rate_ul_min : float ) -> dict [ str , float ]:
"""制御された流量で一定量を吐出します。"""
...
クロードからドライブ（MCP）
1 つの端末でシミュレートされた機器を提供します。
UV 実行の例/serve_pump.py
次に、それを別の MCP クライアントに公開します。
uv 実行 labwire-mcp ws://127.0.0.1:9520
宣言されたすべてのコマンドは、そのスキーマ、ユニット、および
アイデンティティを持っているため、クロードはハードウェアをネイティブに検出して駆動します。参照
クロード スタイル MCP の例/mcp-config.json
サーバーエントリー。
3 つの機器はオリジナルのシミュレートされたデバイス モデルであり、
現実的な遅延、ノイズ、ドリフト、故障モード、安全インターロック。彼らは
実際のベンダーのハードウェアのエミュレーションではなく、互換性もありません。
本物の楽器が主張されています。クローズドループのデモでは、次のような化学反応が起こります。
デバイスはデモ ハーネスによって計算されます。 S2/S3の安全確認
コマンドは、オペレータの身元ではなく展開ポリシーを証明します。暗号化
オペレーター バインディングは v0.2 ではなくロードマップにあります。非目標
現在: フリート制御、Web UI、スタブ API キーを超えた認証、実際のハードウェア
ドライバー、クラウドホスティング。
自分の楽器を持参してください（オフィドブリッジ）
Labwire は、何千ものドライバーを再実装することを目的としていません。の
ophyd ブリッジはあらゆる古典を公開します
ophyd デバイス、その下のハードウェア層
Labwireとして放射光施設で広く使われているBluesky
命令

コメント:
uv run labwire-ophyd annotate ophyd.sim:motor -o labwire-ophyd.yaml
demo-ophyd # ブリッジされた ophyd.sim デバイス上でのピーク検出スキャンを実行する
ophyd はデバイスの構造を知っています。それには単位もリスクの概念もありません。
小さな YAML アノテーション ファイルがそれらを提供しますが、ブリッジはアノテーションの提供を拒否します。
数量に UCUM 単位がなく、作動が S2 に分類されるデバイス。
エージェントは何かを移動するにはオペレーターの確認を提示する必要があります。確認済み
シミュレートされたデバイスとチャネル アクセス上のソフト EPICS IOC に対して、
物理ハードウェアに対しては決して反対しないでください。パッケージの制限セクションは次のとおりです
それが何を意味し、何を意味しないのかを明確にします。
PyLabRobot ブリッジは液体に対しても同じことを行います
ophyd デバイスは、特定のものをテストするために構築されました。
signal-shape 、このプロトコルが設計された形状であるため、
それらを橋渡しすることは、見た目よりもうまくいかないことが判明しました。リキッドハンドラーのコマンドが機能する
物の上にあり、その状態がツリーです。
署名された証拠を使用して、demo-pylabrobot # 段階希釈を作成します
それは機能し、緊張した場所は滑らかになるのではなく書き留められます
終わった。 SPEC-FINDINGS.md は、そのうちの 8 つの正直なリストです。
v0.3 の具体的な推奨事項を示します。短いバージョン: v0.2 モデルのアクション

[切り捨てられた]

## Original Extract

Contribute to benchwire/labwire development by creating an account on GitHub.

GitHub - benchwire/labwire · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
benchwire
/
labwire
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
41 Commits 41 Commits .github .github docs docs examples examples packages packages spec spec tests tests .gitignore .gitignore .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE PRIOR_ART.md PRIOR_ART.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SPEC-FINDINGS.md SPEC-FINDINGS.md pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json uv.lock uv.lock View all files Repository files navigation
An open protocol for AI-controlled laboratory instruments. Think "MCP for
lab equipment": one universal way for AI agents to discover an
instrument's capabilities, command it, stream its measurements, and
walk away with cryptographically signed proof of what was done.
Working title, protocol v0.3 draft. The wire protocol will change before
1.0. Feedback and prior-art corrections are very welcome; see
CONTRIBUTING.md .
Self-driving labs need instruments that agents can operate safely and
auditably. Today every vendor speaks a different dialect and every
integration is bespoke. Labwire's bet is that the missing piece is small and
buildable now:
AI-agent-native. Capability discovery modeled on MCP: an instrument
describes its commands as JSON Schema, so any agent framework can drive it
with zero glue code: the bundled MCP adapter proves it.
Signed results. Every run can produce an ed25519-signed manifest over
the exact telemetry recorded: portable, tamper-evident evidence of what
instrument did what, verified by one CLI command.
Safety and physical typing in the protocol. Mandatory UCUM units on
every quantity, S0-S3 safety classes where irreversible actions take an
operator confirmation and hazardous ones take an operator grant an agent
cannot mint , bound to the exact parameters, interlocks, cancellation, and
typed errors with retryability. All specified, not vendor add-ons.
Things, not only quantities. v0.3 adds resources (typed, readable
instrument state, like a liquid handler's deck) and typed references
(parameters that name a well or a site, validated against current state,
with errors that hand the agent the read that recovers). Designed so that
discovery alone leads an agent to the deck, with no prompt coaching; CI
enforces the preconditions, and the demo asserts the behaviour.
Runnable by a stranger in 5 minutes. Zero hardware: the reference
implementation ships three realistic simulated instruments.
From PyPI (Python 3.12+), no checkout needed:
pip install labwire
That installs the SDK, three simulated instruments, their drivers, and the
labwire CLI. Declare an instrument and drive it end to end, straight from
a Python file:
import asyncio
from labwire . core import (
PROTOCOL_VERSION ,
CommandContext ,
IdentityInfo ,
Instrument ,
InstrumentServer ,
LabwireClient ,
MemoryTransport ,
command ,
)
from pydantic import BaseModel , ConfigDict
class MassReading ( BaseModel ):
model_config = ConfigDict ( extra = "forbid" ) # closed schema: the unit walker demands it
mass_g : float
class Balance ( Instrument ):
"""A one-command instrument. Units are mandatory: omit "g" and this refuses to declare."""
identity = IdentityInfo (
manufacturer = "You" , model = "Balance-1" , serial_number = "B-1" , firmware_version = "1.0"
)
@ command ( returns_units = { "mass_g" : "g" })
async def measure ( self , ctx : CommandContext ) -> MassReading :
"""Report the settled mass."""
return MassReading ( mass_g = 12.3456 )
async def main () -> None :
server = InstrumentServer ( Balance ())
client_end , server_end = MemoryTransport . pair ()
server . attach ( server_end )
async with LabwireClient . attach ( client_end ) as client :
descriptor = await client . describe ()
print ( descriptor . identity . model , "speaks Labwire protocol" , PROTOCOL_VERSION )
handle = await client . submit ( "measure" , {})
result = await handle . result ( timeout = 5.0 )
print ( "mass:" , result [ "mass_g" ], "g" )
asyncio . run ( main ())
Telemetry streaming, safety confirmations, and ed25519-signed run bundles
are a few lines more; examples/quickstart.py
shows them.
From source, for the full demos and examples:
git clone https://github.com/benchwire/labwire.git && cd labwire
make setup # uv installs Python 3.12 + everything
uv run examples/quickstart.py # 60 s: drive a simulated balance end to end
uv run examples/streaming.py # telemetry, cancellation, interlock recovery
make demo # closed-loop optimization + signed evidence
make demo runs a full autonomous experiment campaign: a scripted optimizer
tunes heater voltage and reagent flow rate across three simulated
instruments, converges on the hidden yield optimum, and ends by verifying the
winning run's signed bundle:
safety: pump dispense is class S2 (irreversible); running under the operator standing grant
run 13 V= 15.0 V -> T= 68.0 degC q= 127 uL/min yield= 87.2% best= 87.2%
converged: best yield 87.2% at 15.0 V (68.0 degC), 127 uL/min in 14 experiments
signed evidence: demo_runs/d3b15e9f-...
labwire verify: OK - authentic
make demo-claude runs the same loop with a real Claude agent planning
the experiments through the instruments' tool schemas (needs
ANTHROPIC_API_KEY ; degrades gracefully to the scripted optimizer without
it).
flowchart LR
subgraph agents [Agents]
claude[Claude / any MCP client]
script[Optimizer / LabwireClient]
end
adapter["labwire-mcp<br>(MCP adapter)"]
subgraph servers [Instrument Servers - labwire-core]
psu[PowerSupply driver]
pump[SyringePump driver]
bal[Balance driver]
end
subgraph devices [Native wire protocols - labwire-sim]
scpi["SimPSU-3005<br>SCPI over TCP"]
serial["SimPump-200<br>serial-style lines"]
stream["SimBalance-120<br>streaming readings"]
end
verify["labwire verify<br>(ed25519 + RFC 8785)"]
claude -->|MCP tools| adapter
adapter -->|JSON-RPC / WebSocket| servers
script -->|discover / command / stream| servers
psu --> scpi
pump --> serial
bal --> stream
servers -->|signed run bundles| verify
Loading
The protocol is JSON-RPC 2.0 over WebSocket (stdio specified), with an
MCP-inspired initialize/capability handshake, a push-first command lifecycle,
sequenced telemetry, protocol-level safety interlocks, and normative signed
run manifests. The full specification lives at spec/SPEC.md ,
and every JSON example in it is machine-validated against the implementation
in CI.
Package
What it is
labwire-core
Server + client SDKs, transports, session layer, signing, JCS
labwire-sim
Three realistic simulated instruments speaking native wire protocols
labwire-drivers
Drivers wrapping those native protocols as Labwire instruments
labwire-mcp
MCP adapter: every instrument command becomes an MCP tool
labwire-cli
labwire verify <bundle> : authenticate signed run evidence
labwire-ophyd
Bridge: serve any ophyd (Bluesky) device as a Labwire instrument
labwire-pylabrobot
Bridge: serve a PyLabRobot liquid handler as a Labwire instrument
spec/
The protocol specification (v0.2 draft)
examples/
Quickstart, streaming/recovery, and the closed-loop demo
Wrapping your own device is a class and a decorator:
class MyPump ( Instrument ):
identity = IdentityInfo ( manufacturer = "You" , model = "Pump-1" ,
serial_number = "001" , firmware_version = "1.0" )
flow = channel ( "flow_rate" , unit = "uL/min" ) # UCUM codes are mandatory
@ command (
units = { "volume_ul" : "uL" , "rate_ul_min" : "uL/min" },
returns_units = { "dispensed_ul" : "uL" },
safety_class = "S2" , # irreversible: needs confirmation
)
async def dispense ( self , ctx : CommandContext ,
volume_ul : float , rate_ul_min : float ) -> dict [ str , float ]:
"""Dispense a volume at a controlled flow rate."""
...
Drive it from Claude (MCP)
Serve a simulated instrument in one terminal:
uv run examples/serve_pump.py
Then expose it to any MCP client from another:
uv run labwire-mcp ws://127.0.0.1:9520
Every declared command appears as an MCP tool with its schema, units, and
identity, so Claude discovers and drives the hardware natively. See
examples/mcp-config.json for a Claude-style MCP
server entry.
The three instruments are original simulated device models , with
realistic latency, noise, drift, failure modes, and safety interlocks. They
are not emulations of any real vendor's hardware, and no compatibility with
real instruments is claimed . In the closed-loop demo, the chemistry between
devices is computed by the demo harness. Safety confirmation for S2/S3
commands proves deployment policy, not operator identity; cryptographic
operator binding is on the roadmap , not in v0.2. Non-goals for
now: fleet control, web UI, auth beyond a stub API key, real hardware
drivers, cloud hosting.
Bring your own instruments (ophyd bridge)
Labwire does not aim to reimplement thousands of drivers. The
ophyd bridge exposes any classic
ophyd device, the hardware layer under
Bluesky that is widely used at synchrotron facilities, as a Labwire
instrument:
uv run labwire-ophyd annotate ophyd.sim:motor -o labwire-ophyd.yaml
make demo-ophyd # a peak-finding scan over bridged ophyd.sim devices
ophyd knows a device's structure; it carries no units and no notion of risk.
A small YAML annotation file supplies those, the bridge refuses to serve a
device whose quantities have no UCUM unit, and actuation is classified S2 so
an agent must present an operator confirmation to move anything. Verified
against simulated devices and a soft EPICS IOC over Channel Access,
never against physical hardware ; the package's LIMITATIONS section is
explicit about what that does and does not mean.
The PyLabRobot bridge does the same for liquid
handling, and was built to test something specific: ophyd devices are
signal-shaped , which is the shape this protocol was designed around, so
bridging them proved less than it looked like. A liquid handler's commands act
on things, and its state is a tree.
make demo-pylabrobot # a serial dilution, with signed evidence
It works, and the places it strained are written down rather than smoothed
over. SPEC-FINDINGS.md is an honest list of eight of them,
with concrete recommendations for v0.3. The short version: v0.2 models actions

[truncated]
