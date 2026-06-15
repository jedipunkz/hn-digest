---
source: "https://github.com/T-Chartrand/tuningfork"
hn_url: "https://news.ycombinator.com/item?id=48539346"
title: "Tuningfork – LLM agent grounding rules derived from human reality-testing"
article_title: "GitHub - T-Chartrand/tuningfork: Grounding rules for LLM agents, derived from human reality-testing · GitHub"
author: "T-Chartrand"
captured_at: "2026-06-15T11:10:41Z"
capture_tool: "hn-digest"
hn_id: 48539346
score: 1
comments: 0
posted_at: "2026-06-15T10:40:39Z"
tags:
  - hacker-news
  - translated
---

# Tuningfork – LLM agent grounding rules derived from human reality-testing

- HN: [48539346](https://news.ycombinator.com/item?id=48539346)
- Source: [github.com](https://github.com/T-Chartrand/tuningfork)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T10:40:39Z

## Translation

タイトル: Tuningfork – 人間の現実テストから派生した LLM エージェントのグラウンディング ルール
記事のタイトル: GitHub - T-Chartrand/tuningfork: 人間の現実テストから派生した LLM エージェントのグラウンディング ルール · GitHub
説明: 人間の現実テストから派生した LLM エージェントのグラウンディング ルール - T-Chartrand/音叉

記事本文:
GitHub - T-Chartrand/音叉: 人間の現実テストから派生した LLM エージェントのグラウンディング ルール · GitHub
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
Tチャートランド
/
音叉
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
C

頌歌
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット docs docs サンプル サンプル src/ Tuningfork src/ Tuningfork テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
人間の現実テストから派生した、LLM エージェントの基本ルール。
現実の認識と説得力のある内部の認識を日常的に区別しなければならない人間
製造業者は、そのための実用的なチェックを洗練するために数十年を費やしてきました。それらの小切手
エージェントの幻覚問題に直接マッピングされることが判明 - 多くの場合、それ以上
ML 文献のフレームよりもきれいです。音叉はそのマッピングです、
9 つのルールとして記述され、依存関係のない小さな Python として出荷されます。
リファレンス実装。
この名前は、人間のテクニックの 1 つである物理的な音叉に由来しています。
耳には独立したチャンネルを通じて幻聴が遮断されます。
それは偽の信号とは議論しません - それは状態を壊します。それが、
このライブラリ全体の設計原則。
検証者が疑わしいシステムの外にいる場合、チェックは終了します。
自身の出力を再読み取りするモデルは、独自の故障モードを共有します。
自らの捏造をすらすらと確認する。 grep、パーサー、チェックサム、終了
コードではできません。独立したチャネルからの決定的な確認の 1 つは、
最終;同じモデルを 100 回再チェックする必要はありません。ここにあるものはすべて以下から導かれます
それは、環境が真実の源であり、モデルの記憶が真実であるということです。
キャッシュが古い可能性があります。
ルール
フェーズ
ワンライナー
G0 非対称信頼
すべてを統治する
コンテンツは有罪判決を下すことはあっても、無罪判決を下すことは決してない — 信頼フローはソース追跡のみからのもの
G1 アサート前検証
予見する
ツールチェックできるクレームは、記載される前にチェックされる必要があります。
G2 閉ループ実行
認識する
レポート

結果は観察されず、コマンドは発行されませんでした。読み取り専用の観察は最終的なものです
G3 不一致の三角測量
認識する
ツールは記憶に勝ります。サプライズに関する独立したチェックが 1 つあります。 1 つの決定的な確認が最終的なものとなる
G4 負のスペースのプローブ
予見する
記憶されているエンティティに依存する前に、存在を調べます。既知の製造痕跡のカタログを保管する
G5 再現性スナップショット
スナップアウト
修正後は、ツールの出力のみから状態を再構築します。壊れたナラティブは何も引き継がれません。
G6 コスト段階予算
継続的な
生成前に決定される爆発範囲による階層検証。疑わしいほど完璧な主張のランクが上がる
G7 パッシブ独立バリデーター
継続的な
安価な決定論的モニターはあらゆるもので実行され、ジェネレーターの許可を求めることはありません
G8 ソースの再帰属
判決後
検証済みの偽の出力は、ジェネレーターに関する証拠です。つまり、ジェネレーターをマイニングします。信念と行動は切り離されている
根拠を含む全文: docs/framework.md · その背後にあるストーリー: docs/essay.md
pip install -e 。
音叉インポートから (GroundedAgent 、 ValidatorBank 、
CitationValidator 、 PathValidator 、 JsonBlockValidator )
銀行 = ValidatorBank ([
CitationValidator ( valid_source_ids = [ "1" , "2" , "3" ])、
PathValidator (証拠パス = ツール戻りパス)、
JsonBlockValidator()、
])
エージェント = グラウンデッドエージェント (生成 = my_llm_callable 、バンク = 銀行)
結果 = エージェント。 run ( "ソース [1] ～ [3] を要約し、関連する構成ファイルをリストします。" )
print (結果 . 層 . 根拠 ) # 生成前にクレームの価格がどのように設定されたか
print (結果 . レポート . 概要 ()) # 独立チャネルが観察した内容
print ( result . trustworthy ) # モデルの判断ではなくバリデーターの判断
ハーネスでは、バリデーター障害時に 1 回だけの再生成パスが許可されます。
APOLではなくバリデーターの証拠を提供した

ああ、プロンプト。 ２つ目の失敗は、
同じチャネルを再試行すると再チェックされるため、未解決として報告されます。
チェックしてください。
v0.3.0 では、オーバーレイをオンにした小さな実行可能なエージェントが追加されます。
ツール ループ (stdlib urllib 経由の Anthropic Messages API — SDK なし)
すべてのアシスタントの発話は、そこから構築された証拠に照らして検証されます。
ツールの ACTUAL が返され、存在しないツール呼び出しは代わりに拒否されます。
即興で、証拠に基づいた修正ターンが 1 回許可され、
拒否はセッション間で台帳ファイルに保持されます。つまり、
既知の製造はリセットされる代わりに蓄積されます。
API キーや有料アカウントは必要ありません。トランスポートはプラグイン可能であり、
OpenAI 互換アダプターは、Ollama (ローカル、無料)、Groq、OpenRouter、
LM スタジオと vLLM:
チューニングフォークインポートから ChildAgent 、 OpenAISupportLLM 、builtin_fs_tools
llm = OpenAIContinuousLLM (model = "qwen2.5:7b" ) # ローカルホスト上の Ollama
エージェント = ChildAgent ( llm 、builtin_fs_tools ( "." ))
結果 = エージェント。 run ( "./docs 内のどのファイルが 'echo' に言及していますか? パスを引用してください。" )
print (結果が信頼できる、結果が回答)
MCP サーバーは、最小限の stdlib クライアントを介してファーストクラスのツールとして接続されます。
(改行区切りの JSON-RPC over stdio):
音叉インポート MCPServer 、 mcp_tools から
srv = MCPServer ([ "python3" , "my_server.py" ], name = "files" )
サーバー。開始()
エージェント = ChildAgent ( AnthropicLLM ()、mcp_tools ( srv ))
実行可能なバージョンについては、examples/agent_demo.py を参照してください。
モデルを「より正直」にするラッパーではありません。モデルは変更されていません。
評価スイートではありません。ランタイムハーネスです。
新しいコンポーネントではありません - 検証してからアサートする、閉ループの実行、および
出力ガードレールはすべて以前の作業に存在します (検証チェーン、
SelfCheckGPT、ReAct/Reflexion、ガードレール フレームワーク)。ここでの新機能は
統一フレームと終了原則、

ほとんどのフレームワークは
欠落: 検証をまったく行わないか、停止できないかのどちらかです。
v0.2.0 — リファレンス実装、完全なテスト スイートに合格、stdlib のみ。 EchoValidator (構造的な先行指標としての繰り返し) と RejectionLedger (マイニングされた拒否から蓄積される G4 カタログ) が含まれます。
ロードマップ: カバレッジバリデータ (応答が無視されたことを証明)、非同期
バリデーターバンク、一般的なエージェントフレームワーク用のアダプター。
人間の現実テストから派生した LLM エージェントの基本ルール
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

Grounding rules for LLM agents, derived from human reality-testing - T-Chartrand/tuningfork

GitHub - T-Chartrand/tuningfork: Grounding rules for LLM agents, derived from human reality-testing · GitHub
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
T-Chartrand
/
tuningfork
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits docs docs examples examples src/ tuningfork src/ tuningfork tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Grounding rules for LLM agents, derived from human reality-testing.
Humans who must routinely distinguish real perception from convincing internal
fabrication have spent decades refining practical checks for it. Those checks
turn out to map directly onto the agent hallucination problem — often more
cleanly than the framings in the ML literature. tuningfork is that mapping,
written down as nine rules and shipped as a small, dependency-free Python
reference implementation.
The name comes from one of those human techniques: a physical tuning fork held
to the ear interrupts auditory hallucination through an independent channel.
It doesn't argue with the false signal — it breaks the state. That is the
design principle of this entire library.
A check terminates when the verifier sits outside the system being doubted.
A model re-reading its own output shares its own failure modes — it can
fluently confirm its own fabrication. A grep, a parser, a checksum, an exit
code cannot. One deterministic confirmation from an independent channel is
final ; a hundred same-model re-checks are not. Everything here follows from
that: the environment is the source of truth, and the model's memory is a
cache that may be stale.
Rule
Phase
One-liner
G0 Asymmetric Trust
governs all
Content can convict, but never acquit — trust flows from source-tracing only
G1 Verify-Before-Assert
foresee
A claim that could be tool-checked must be, before it's stated
G2 Closed-Loop Execution
recognize
Report observed results, never issued commands. Read-only observations are terminal
G3 Disagreement Triangulation
recognize
Tool beats memory; one independent check on surprises; one deterministic confirmation is final
G4 Negative-Space Probing
foresee
Probe for existence before relying on remembered entities; keep a catalog of known fabrication signatures
G5 Reproducibility Snapshot
snap out
After a correction, rebuild state from tool output only — nothing from the broken narrative carries over
G6 Cost-Tiered Budget
continuous
Tier verification by blast radius, decided before generation; suspiciously perfect claims get their tier raised
G7 Passive Independent Validators
continuous
Cheap deterministic monitors run on everything and never ask the generator's permission
G8 Source Re-attribution
after the verdict
A verified-false output is evidence about the generator — mine it; belief and action are decoupled
Full text with rationale: docs/framework.md · The story behind it: docs/essay.md
pip install -e .
from tuningfork import ( GroundedAgent , ValidatorBank ,
CitationValidator , PathValidator , JsonBlockValidator )
bank = ValidatorBank ([
CitationValidator ( valid_source_ids = [ "1" , "2" , "3" ]),
PathValidator ( evidence_paths = tool_returned_paths ),
JsonBlockValidator (),
])
agent = GroundedAgent ( generate = my_llm_callable , bank = bank )
result = agent . run ( "Summarize sources [1]-[3] and list the config files involved." )
print ( result . tier . rationale ) # how the claim was priced before generation
print ( result . report . summary ()) # what the independent channels observed
print ( result . trustworthy ) # validators' verdict, not the model's
The harness permits exactly one regeneration pass on validator failure —
fed the validator evidence, not an apology prompt. A second failure is
reported as unresolved, because retrying the same channel is re-checking the
check.
v0.3.0 adds a small runnable agent with the overlay on: an ordinary
tool loop (Anthropic Messages API via stdlib urllib — no SDK) where
every assistant utterance is validated against evidence built from the
tools' ACTUAL returns, nonexistent tool calls are refused instead of
improvised, one evidence-fed correction turn is permitted, and
rejections persist to a ledger file across sessions — the catalog of
known fabrications accumulates instead of resetting.
No API key or paid account required — the transport is pluggable, and an
OpenAI-compatible adapter covers Ollama (local, free), Groq, OpenRouter,
LM Studio, and vLLM:
from tuningfork import ChildAgent , OpenAICompatibleLLM , builtin_fs_tools
llm = OpenAICompatibleLLM ( model = "qwen2.5:7b" ) # Ollama on localhost
agent = ChildAgent ( llm , builtin_fs_tools ( "." ))
result = agent . run ( "Which files in ./docs mention 'echo'? Cite paths." )
print ( result . trustworthy , result . answer )
MCP servers wire in as first-class tools via a minimal stdlib client
(newline-delimited JSON-RPC over stdio):
from tuningfork import MCPServer , mcp_tools
srv = MCPServer ([ "python3" , "my_server.py" ], name = "files" )
srv . start ()
agent = ChildAgent ( AnthropicLLM (), mcp_tools ( srv ))
See examples/agent_demo.py for the runnable version.
Not a wrapper that makes a model "more honest." The model is unchanged.
Not an eval suite. It's a runtime harness.
Not novel components — verify-then-assert, closed-loop execution, and
output guardrails all exist in prior work (Chain-of-Verification,
SelfCheckGPT, ReAct/Reflexion, guardrails frameworks). What's new here is
the unifying frame and the termination principle, which most frameworks
lack: they either never verify or can't stop.
v0.2.0 — reference implementation, full test suite passing, stdlib only. Includes EchoValidator (repetition as a structural leading indicator) and RejectionLedger (the G4 catalog accumulates from mined rejections).
Roadmap: coverage validator (evidence the response ignored ), async
validator bank, adapters for popular agent frameworks.
Grounding rules for LLM agents, derived from human reality-testing
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
