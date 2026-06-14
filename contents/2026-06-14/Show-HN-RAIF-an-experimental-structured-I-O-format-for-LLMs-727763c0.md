---
source: "https://github.com/skrrt-sh/raif-standard"
hn_url: "https://news.ycombinator.com/item?id=48532748"
title: "Show HN: RAIF – an experimental structured I/O format for LLMs"
article_title: "GitHub - skrrt-sh/raif-standard: RAIF — Repairable AI Interchange Format: a wire format for LLM-emitted JSON that round-trips losslessly, self-repairs syntax errors, and costs fewer tokens than JSON. · GitHub"
author: "truehazker"
captured_at: "2026-06-14T21:32:12Z"
capture_tool: "hn-digest"
hn_id: 48532748
score: 1
comments: 1
posted_at: "2026-06-14T21:08:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RAIF – an experimental structured I/O format for LLMs

- HN: [48532748](https://news.ycombinator.com/item?id=48532748)
- Source: [github.com](https://github.com/skrrt-sh/raif-standard)
- Score: 1
- Comments: 1
- Posted: 2026-06-14T21:08:55Z

## Translation

タイトル: HN を表示: RAIF – LLM 用の実験的な構造化 I/O フォーマット
記事のタイトル: GitHub - skrrt-sh/raif-standard: RAIF — Repairable AI Interchange Format: LLM が発行する JSON のワイヤ形式。ロスレスでラウンドトリップし、構文エラーを自己修復し、JSON よりもトークンのコストが低くなります。 · GitHub
説明: RAIF — Repairable AI Interchange Format: LLM が発行する JSON のワイヤ形式で、ロスレスでラウンドトリップし、構文エラーを自己修復し、JSON よりもトークンのコストが低くなります。 - skrrt-sh/raif-standard

記事本文:
GitHub - skrrt-sh/raif-standard: RAIF — Repairable AI Interchange Format: LLM が発行する JSON のワイヤ形式。ロスレスでラウンドトリップし、構文エラーを自己修復し、JSON よりもトークンのコストが低くなります。 · GitHub
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
skrrt-sh
/
raif標準
出版

ic
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .github/ workflows .github/ workflows 資産 アセット 適合性 適合性 docs docs パッケージ パッケージ .gitignore .gitignore CONTEXT.md HANDOFF.md HANDOFF.md LICENSE LICENSE README.md README.md misse.toml misse.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
修復可能な AI 交換フォーマット
JSON 言語モデルのドロップイン層が生成する - 構造化出力、JSON モード、
厳密なオブジェクト、ツール引数も同様です。独自の構文エラー、ラウンドトリップを修復します。
ロスなく JSON に変換され、トークンのコストが最大 14% 削減されます。
JSON は決定論的なライター向けに設計されました。 LLM は 1 つではありません。クロージングを落とします。
中括弧を使用したり、出力をマークダウン フェンスで囲んだり、ストリームの途中で切断されたりします。今日は毎日
実稼働スタックは、正規表現、再試行、および jsonrepair を使用してこれにパッチを適用します。
RAIF はその仮定を覆します。作家はモデルです。読者は通訳です
を修復、検証、正規化できます。形式は行指向です。
値優先であり、生成された出力の一般的な故障モードが次のとおりであるように構築されています。
建設により回収可能。
特定の API や機能に関連付けられていません。 RAIF はあらゆるもののドロップイン層です
モデルが JSON を生成するようにします — 構造化出力、JSON/厳密モード、応答
スキーマ、関数の引数 — そして、JSON が軽量になり、自己修復され、
無損失。ツール呼び出しは 1 つの使用例であり、重要な点ではありません。
import { encode , decode } from "raif-format" ;
encode ( { ユーザー : { 名前 : "Ada" 、メール : "ada@x.io" } 、アクティブ : true } ) ;
// アクティブ=true
// ユーザー={email=ada @x .io,name=Ada}
デコード ( raif ) 。価値 ; // → 正確な JSON オブジェクトを返す
なぜ JSON だけではいけないのか
JSON
RAIF
トークンコスト

ベースライン
−14%
構文エラーの回復
外部 ( jsonrepair )、サイレント
内蔵、すべての修理が報告される
切り捨てられた出力
全か無かの解析
葉ごとの回収率 (46% vs 41% 葉)
ロスレス往復
該当なし
バイト正確、正規、5,000 シードのファズで実証済み
スキーマ型のデコード
該当なし
オプション — 「null」は null ではなく文字列のままです
トークンコストと他のフォーマットの比較
低いほど良いです。 cl100k と o200k を使用して 18 形状のコーパス全体を測定
トークナイザー。 bun Compare で再現します。
TOON と YAML は LLM 入力形式です。 RAIF は LLM 出力をターゲットとします。上記の数字
同一ペイロードのトークンコストのみを比較 - 修復と回復の保証
以下は RAIF に固有のものです。
自己修復デコード
マークダウン フェンス、モード マーカー、スリップした : → = セパレーターを自動修正します。すべての修理を報告します。曖昧なものは拒否します。値を決して書き換えません。
トランケーションリカバリ
decodeLenient は、カットオフ ストリームの無傷のリーフとリーフごとのエラー リストを返します。等しいトークン バジェットでリーフの回復率は 46% でしたが、JSON + jsonrepair の場合は 41% でした。
ロスレス往復
decode(encode(x)) は x と等しく、正規の UTF-8 ソートと冪等の出力を使用します。 5,000 シードのファズテストで実証済み。
スキーマ型のデコード
オプションのスキーマは型を固定します。文字列フィールドの下の裸の null は文字列 "null" のままになります。リテラル文字列忠実度のトラップは構築によりなくなりました。
依存性ライト
純粋な TypeScript エンコーダーおよびデコーダー。コアパスにはランタイム依存関係はありません。
デコード ( "```\nactive=true\nuser.name: Ada\n```" ) ;
// ok: true、修復: [markdown_tripped, separator_coerced]
decodeLenient ( "<raif>\ncity=Oslo\nlat" ) ; // ストリームが切断されました
// { 値: { 都市: "オスロ" } 、切り捨て: true、エラー: [{ line: 2, … }] }
API
encode ( obj , opts ? ) // JSON オブジェクト → RAIF
decode ( raif , schema ? ) // → { ok, value, Repair } (修復してから解析します)
デコードレン

ient ( raif , schema ? ) // → { 値、エラー、切り捨て、修復 } (スローしません)
fix ( raif , schema ? ) // → 正規の RAIF
validate ( raif , schema ? ) // 読み取り専用の正規性チェック
実装
2 つの言語で同じ 4 つの機能を備えた API を共有し、
リファレンスから生成された言語に依存しない適合コーパス
エンコーダ。どちらも依存関係が軽い (コア パスにランタイム依存関係がありません)。
raif インポートからエンコード、デコード
encode ({ "ユーザー" : { "名前" : "エイダ" }, "アクティブ" : True })
decode ( raif )[ "value" ] # → 正確な JSON オブジェクトを返す
クイックスタート
# TypeScript (パッケージ/js)
cd パッケージ/js && bun install
bun テスト # 247 テスト: プロパティ スイート + 共有適合性コーパス
bun run ビルド # デュアル ESM+CJS + タイプ
# Python (パッケージ/py)
cd パッケージ/py && UV 同期
uv run pytest # ユニット + 適合性 + (開発のみ) 差分と TS リファレンス
ツールチェーンは、mise.toml (mise install) に固定されています。
JSON の代わりに RAIF をネイティブに出力する Llama-3.2-3B LoRA が公開されています。
Hugging Face — トークンの節約と切り捨ての回復をわずかに抑えます。
ローカルおよび自己ホスト型の推論:
モデル: skrrt-sh/raif-llama-3.2-3b-lora — v0.5 承認ゲートをクリア (100% / 95%)
エージェント グレード: skrrt-sh/raif-qwen3-4b-lora — 実際のセルフホスト エージェント用の Qwen3-4B (~14 GB VRAM)。ホールドアウト 98% パース / 95% 忠実度、3/4 ゲート メトリクスをマージンでクリア
小型: skrrt-sh/raif-qwen2.5-0.5b-lora — 6 倍小さいベースでの同じレシピ (97% 解析、81% ホールドアウト忠実度)。小さなローカルモデルをどこまで拡張できるかに関する研究
トレーニング ワークストリーム: raif-lora
これらのモデルは JSON ではなく RAIF を出力します。上記のパッケージを使用して出力をデコードします。
( pip install raif-format → from raif import decode 、または npm の raif-format)。
RAIF は、LLM 出力の単一の JSON オブジェクト (文字列、数値) をカバーします。

rs、ブール値、null、
それらの配列とネストされたオブジェクト。一般的な交換形式ではありませんが、
圧縮、スキーマ言語、LLM 入力形式ではありません。
docs/raif_v0.3_spec.md 基本仕様。 v0.5 サーフェス = このドキュメント + ADR
docs/adr/0001…0019 設計決定 (v0.5 修正)、ファイルごとに 1 つ
適合性/言語に依存しないテスト コーパス (クロス IPL コントラクト)
package/js/src/raif.ts TypeScript リファレンス エンコーダ + デコーダ (純粋)
パッケージ/py/src/raif/ Python 実装 (エンコード/デコード/修正/検証)
misse.toml の固定ツールチェーン (node、bun、python、uv)
CONTEXT.md 用語集 - 最初にお読みください
HANDOFF.md プロジェクトの完全な状態と調査結果
ライセンス
Apache-2.0 。特許付与は、他の人が行う可能性があるワイヤ形式に関係します。
実装 — RAIF 上に構築するすべての人が対象となります。
RAIF — Repairable AI Interchange Format: LLM が発行する JSON のワイヤ形式で、ロスレスでラウンドトリップし、構文エラーを自己修復し、JSON よりもトークンのコストが低くなります。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

RAIF — Repairable AI Interchange Format: a wire format for LLM-emitted JSON that round-trips losslessly, self-repairs syntax errors, and costs fewer tokens than JSON. - skrrt-sh/raif-standard

GitHub - skrrt-sh/raif-standard: RAIF — Repairable AI Interchange Format: a wire format for LLM-emitted JSON that round-trips losslessly, self-repairs syntax errors, and costs fewer tokens than JSON. · GitHub
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
skrrt-sh
/
raif-standard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .github/ workflows .github/ workflows assets assets conformance conformance docs docs packages packages .gitignore .gitignore CONTEXT.md CONTEXT.md HANDOFF.md HANDOFF.md LICENSE LICENSE README.md README.md mise.toml mise.toml View all files Repository files navigation
Repairable AI Interchange Format
A drop-in layer for the JSON language models produce — structured outputs, JSON mode,
strict objects, tool arguments alike. It repairs its own syntax errors, round-trips
losslessly to JSON, and costs ~14% fewer tokens.
JSON was designed for a deterministic writer. An LLM is not one: it drops a closing
brace, wraps output in a markdown fence, or gets cut off mid-stream. Today every
production stack patches over this with regex, retries, and jsonrepair .
RAIF inverts the assumption. The writer is a model; the reader is an interpreter
that can repair, validate, and canonicalize . The format is line-oriented,
value-first, and built so that the common failure modes of generated output are
recoverable by construction.
It isn't tied to any one API or feature. RAIF is a drop-in layer for whatever
makes a model produce JSON — structured outputs, JSON/strict mode, response
schemas, function arguments — and it makes that JSON lighter, self-repairing, and
lossless. Tool calls are one use case, not the point.
import { encode , decode } from "raif-format" ;
encode ( { user : { name : "Ada" , email : "ada@x.io" } , active : true } ) ;
// active=true
// user={email=ada @x .io,name=Ada}
decode ( raif ) . value ; // → the exact JSON object back
Why not just JSON
JSON
RAIF
Token cost
baseline
−14%
Syntax-error recovery
external ( jsonrepair ), silent
built-in, every repair reported
Truncated output
all-or-nothing parse
per-leaf recovery ( 46% vs 41% leaves)
Lossless round-trip
n/a
byte-exact, canonical, 5,000-seed fuzz-proven
Schema-typed decode
n/a
optional — "null" stays the string, not null
Token cost vs other formats
Lower is better. Measured across an 18-shape corpus with the cl100k and o200k
tokenizers. Reproduce with bun compare .
TOON and YAML are LLM- input formats; RAIF targets LLM output . The numbers above
compare only token cost on identical payloads — the repair and recovery guarantees
below are unique to RAIF.
Self-healing decode
Auto-fixes markdown fences, mode markers, and slipped : → = separators. Reports every repair; refuses ambiguous ones. Never rewrites values.
Truncation recovery
decodeLenient returns the intact leaves of a cut-off stream plus a per-leaf error list — 46% leaf recovery at equal token budget, vs 41% for JSON + jsonrepair .
Lossless round-trip
decode(encode(x)) equals x , with canonical UTF-8 sort and idempotent output. Proven under a 5,000-seed fuzz test.
Schema-typed decode
An optional schema pins types: a bare null under a string field stays the string "null" . The literal-string fidelity trap is gone by construction.
Dependency-light
Pure TypeScript encoder and decoder. No runtime dependencies in the core path.
decode ( "```\nactive=true\nuser.name: Ada\n```" ) ;
// ok: true, repairs: [markdown_stripped, separator_coerced]
decodeLenient ( "<raif>\ncity=Oslo\nlat" ) ; // stream cut off
// { value: { city: "Oslo" } , truncated: true, errors: [{ line: 2, … }] }
API
encode ( obj , opts ? ) // JSON object → RAIF
decode ( raif , schema ? ) // → { ok, value, repairs } (repairs, then parses)
decodeLenient ( raif , schema ? ) // → { value, errors, truncated, repairs } (never throws)
fix ( raif , schema ? ) // → canonical RAIF
validate ( raif , schema ? ) // read-only canonicality check
Implementations
The same four-function API in two languages, kept in lockstep by a shared,
language-agnostic conformance corpus generated from the reference
encoder. Both are dependency-light (zero runtime dependencies in the core path).
from raif import encode , decode
encode ({ "user" : { "name" : "Ada" }, "active" : True })
decode ( raif )[ "value" ] # → the exact JSON object back
Quick start
# TypeScript (packages/js)
cd packages/js && bun install
bun test # 247 tests: property suite + shared conformance corpus
bun run build # dual ESM+CJS + types
# Python (packages/py)
cd packages/py && uv sync
uv run pytest # unit + conformance + (dev-only) differential vs the TS reference
Toolchains are pinned in mise.toml ( mise install ).
A Llama-3.2-3B LoRA that natively emits RAIF instead of JSON is published on
Hugging Face — it brings the token savings and truncation recovery to small,
local, and self-hosted inference:
Model: skrrt-sh/raif-llama-3.2-3b-lora — clears the v0.5 acceptance gate (100% / 95%)
Agent-grade: skrrt-sh/raif-qwen3-4b-lora — Qwen3-4B for real self-hosted agents (~14 GB VRAM); holdout 98% parse / 95% fidelity, clears 3/4 gate metrics with margin
Tiny: skrrt-sh/raif-qwen2.5-0.5b-lora — the same recipe on a 6×-smaller base (97% parse, 81% holdout fidelity); a study in how far a tiny local model can be pushed
Training workstream: raif-lora
These models emit RAIF, not JSON — decode their output with the package above
( pip install raif-format → from raif import decode , or raif-format on npm).
RAIF covers a single JSON object of LLM output: strings, numbers, booleans, nulls,
arrays of those, and nested objects. It is not a general interchange format,
not compression, not a schema language, and not an LLM- input format.
docs/raif_v0.3_spec.md base specification; the v0.5 surface = this doc + the ADRs
docs/adr/0001…0019 design decisions (the v0.5 amendments), one per file
conformance/ language-agnostic test corpus (the cross-impl contract)
packages/js/src/raif.ts TypeScript reference encoder + decoder (pure)
packages/py/src/raif/ Python implementation (encode/decode/fix/validate)
mise.toml pinned toolchains (node, bun, python, uv)
CONTEXT.md glossary — read first
HANDOFF.md full project state and findings
License
Apache-2.0 . The patent grant matters for a wire format others may
implement — anyone building on RAIF is covered.
RAIF — Repairable AI Interchange Format: a wire format for LLM-emitted JSON that round-trips losslessly, self-repairs syntax errors, and costs fewer tokens than JSON.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
