---
source: "https://github.com/MobileReality/mdma"
hn_url: "https://news.ycombinator.com/item?id=48368005"
title: "MDMA – Turn LLM Responses into Interactive UI via MCP"
article_title: "GitHub - MobileReality/mdma: Interactive documents from Markdown. Extends MD with forms, approvals, webhooks, and more — built for next gen apps · GitHub"
author: "mattsadowsky"
captured_at: "2026-06-03T00:50:01Z"
capture_tool: "hn-digest"
hn_id: 48368005
score: 5
comments: 2
posted_at: "2026-06-02T09:46:45Z"
tags:
  - hacker-news
  - translated
---

# MDMA – Turn LLM Responses into Interactive UI via MCP

- HN: [48368005](https://news.ycombinator.com/item?id=48368005)
- Source: [github.com](https://github.com/MobileReality/mdma)
- Score: 5
- Comments: 2
- Posted: 2026-06-02T09:46:45Z

## Translation

タイトル: MDMA – MCP を介して LLM 応答をインタラクティブ UI に変える
記事のタイトル: GitHub - MobileReality/mdma: Markdown からのインタラクティブ ドキュメント。フォーム、承認、Webhook などで MD を拡張 — 次世代アプリ用に構築 · GitHub
説明: Markdown からのインタラクティブなドキュメント。フォーム、承認、Webhook などで MD を拡張 - 次世代アプリ用に構築 - MobileReality/mdma

記事本文:
GitHub - MobileReality/mdma: Markdown からのインタラクティブなドキュメント。フォーム、承認、Webhook などで MD を拡張 — 次世代アプリ用に構築 · GitHub
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
モバイル現実
/
MDMA
公共
通知
通知を変更するにはサインインする必要があります

設定上
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
226 コミット 226 コミット .changeset .changeset .github .github アセット アセット ブループリント ブループリント デモ デモ docs docs evals evals 例 例 パッケージ パッケージ スキル/ mdma-integration スキル/ mdma-integration .gitignore .gitignore .npmrc .npmrc COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md SECURITY.md SECURITY.md biome.json biome.json Glama.json Glama.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.jsonturbo.jsonturbo.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
マウントされたアプリケーションを含むマークダウン ドキュメント
Markdown からのインタラクティブなドキュメント。次世代アプリ向けに構築
🚀 ライブデモ
·
📖 ドキュメント
·
💬不協和音
今日の AI の会話はプレーン テキストであり、ユーザーは応答を読み、それに手動で対応します。 MDMA はそれを変えます。 LLM が MDMA 仕様を認識している場合、テキストだけでなく対話型コンポーネント (フォーム、テーブル、承認ゲート) で応答できます。会話は実用的なものになります。ユーザーはフォームに記入し、ステップを承認し、構造化データをレビューします。これらはすべてインラインで行われ、アプリがレンダリングおよび処理する方法をすでに知っている予測可能なスキーマを使用します。
ユースケースごとのカスタム UI はありません。自由形式のテキストは解析されません。 AI は構造化され検証されたコンポーネントを生成し、フロントエンドはそれらを即座にレンダリングします。
MDMA は、フェンスで囲まれた MDMA コード ブロックで定義されたインタラクティブ コンポーネントを使用して Markdown を拡張します。通常の Markdown ファイルが対話型アプリケーションになります。
# 患者摂取量
「」MDMA
タイプ: フォーム
id: インテークフォーム
フィールド:
- 名前: 患者名
タイプ: テキスト
ラベル: 「フルネーム」
必須: true
センシティブ: true
- 名前: メールアドレス
タイプ: 電子メール

ラベル: 「メール」
必須: true
センシティブ: true
- 名前: 理由
タイプ: テキストエリア
ラベル: 「訪問の理由」
必須: true
「」
「」MDMA
タイプ: ボタン
id: 送信ボタン
テキスト: 「受付フォームを送信」
バリアント: プライマリ
onAction: 送信
「」
MDMA_AUTHOR プロンプト マトリックス
各セルは、リストされた評価スイートのモデルに特化した MDMA_AUTHOR プロンプト バリアントの合格率を示します。
🟡 スイートで 80 ～ 99% のスコア。
🔴 スイートのスコアが 80% 未満。
あなたのモデルが見つかりませんか? package/prompt-pack/src/prompts/mdma-author/<vendor>/ にプロンプ​​ト バリアントを追加し、PR を開きます。eval スイートを実行して、このテーブルに追加します。
† gpt-5.4 の断続的な重複バグ — gpt-5.4 はワンショット評価を確実に渡しますが、マルチターン評価、カスタムプロンプト評価、およびフロー評価では非決定的な出力重複を示します (実行の約 7 ～ 15%)。モデルは完全で正しい応答を生成し、すぐに出力全体をそのまま再送信するため、[duplicate-ids] 検証エラーが発生します。これは、プロンプト バリアントとは関係のない、モデル レベルの既知の問題です。詳細については、OpenAI コミュニティ スレッドを参照してください。これがユースケースに影響を与える場合は、 gpt-5.5 または gpt-5.2 をお勧めします。
‡ gemini-3.1-pro-preview 確率的プリアンブル ループ — flow-eval 実行の約 7 ～ 15% で、モデルは ```mdma ブロックを開く代わりに、目に見える Markdown 散文として思考連鎖を出力します (例: **製造エラーの調査** を 3 ～ 5 回繰り返します)。エラー。 Google の公式 Gemini 3 プロンプト ガイドによると、これは温度/サンプリングによって引き起こされるモデル レベルの動作です。プロンプト レベルの修正は、ループを排除するのではなく、テスト ループを変更します。決定的なフロー出力が重要な場合は、実稼働マルチステップ フローには gemini-2.5-pro をお勧めします。
* あらゆるラボの小規模/低層モデル (OpenAI m

ini · nano、Anthropic Haiku、Google Gemini Flash など) は、短く構造化されたテスト ケースを実行する評価スイートに合格します。現実世界での長い会話では、幻覚を見たり、前のターンを忘れたり、仕様から逸脱したりする傾向があります。マルチターンダイアログやステートフルフローを伴う運用環境での使用の場合は、同じファミリーのフラッグシップ層モデルをお勧めします。
[i] 応答時間が著しく遅い — 通常、1 ターンの応答には数十秒かかりますが、完全な評価の実行には数分かかります。
各セルには、単一ブロック フィクサー評価 (構造修正、バインディング、PII、フォーム、表/チャート、承認をカバーする 15 のテスト) におけるモデルに特化した MDMA_FIXER プロンプト バリアントの合格率が表示されます。フィクサーは、 validate() に失敗した LLM 出力の自動修復を強化するものです。サポートされているすべてのモデルは、モデルに合わせたインライン ガード (先頭区切り文字なし、入力構造の保持、テーブル キーの方向、すべてのプレースホルダーの置換、リストにあるすべてのエラーの修正など) を介して ✅ に到達します。
✅ シングルブロックフィクサー評価で 100% (15/15)。
推論トークンのリーク抑制 — 推論風味の Gemini Pro バリアントおよび Grok 4.3 の場合、修正者は、すべての応答の前に目に見える「思考: トピック」の散文を表示します。 eval 構成は、 passthrough.reasoning.exclude: true (デモの usePreviewValidation はプロバイダーごとに同じことを行います) を設定して、プロンプト層ではなく API 層で応答本文から推論トークンを削除します。
9 つの組み込みコンポーネント タイプ。すべて @mobile-reality/mdma-renderer-react によってすぐにレンダリングされます。
さらに、標準の Markdown コンテンツ (見出し、段落、リスト、コード ブロック、画像、リンク、表など) がコンポーネント間でインラインでレンダリングされます。
組み込みのチャート レンダラーは、データを意図的にプレーン テーブルとしてレンダリングするため、ライブラリは軽量に保たれます。実際のチャートを取得するには、reg

カスタムレンダラーをスターします:
import { MdmaDocument } から '@mobile-reality/mdma-renderer-react' ;
import { MyRechartsRenderer } から './MyRechartsRenderer' ;
関数 App ( { ast , ストア } ) {
戻る (
< MdmaDocument
ast = { ast }
ストア = { ストア }
カスタマイズ = { {
コンポーネント: {
チャート: MyRechartsRenderer 、
} 、
} }
/>
) ;
}
このパターンは、任意の組み込みコンポーネントをオーバーライドする場合に機能します。つまり、customizations.components.<type> の下にカスタム React コンポーネントを渡します。
# コア — MDMA ドキュメントを解析して実行する
npm install @mobile-reality/mdma-parser @mobile-reality/mdma-runtime
# レンダリングに反応する
npm install @mobile-reality/mdma-renderer-react
# AI オーサリング — LLM ベースの生成のためのシステム プロンプト
npm install @mobile-reality/mdma-prompt-pack
# 検証 — MDMA ドキュメントの静的分析
npm install @mobile-reality/mdma-validator
# CLI — 対話型プロンプトビルダー + ドキュメント検証
npx @mobile-reality/mdma-cli
すべてのパッケージは @mobile-reality npm 組織の下で公開されます。
'統合' からインポート {統合} ;
「remark-parse」からremarkParseをインポートします。
import {remarkMdma} から '@mobile-reality/mdma-parser' ;
import { createDocumentStore } から '@mobile-reality/mdma-runtime' ;
'@mobile-reality/mdma-spec' からタイプ { MdmaRoot } をインポートします。
// 1. マークダウンを AST に解析します
const プロセッサ = 統合された () 。 (remarkParse) を使用します。使用 (remarkMdma) ;
const ツリー = プロセッサ。解析 (マークダウン) ;
const ast = (プロセッサを待機します。実行(ツリー)) MdmaRoot ;
// 2. リアクティブドキュメントストアを作成します
const store = createDocumentStore ( ast , {
documentId : 'my-doc' 、
セッション ID : 暗号。ランダムUUID ( ) 、
} ) ;
// 3. 状態の変更をサブスクライブします
店舗も購読 ( ( 状態 ) => {
コンソール。 log ( 'Bindings:' , state . bindings ) ;
} ) ;
// 4. ユーザーアクションをディスパッチします
店舗も発送 ( {
タイプ: 'FIELD_C

絞首刑」、
コンポーネントId : 'intake-form' 、
フィールド: '患者名' 、
値: 'ジェーン・ドウ' 、
} ) ;
チャット中
import { buildSystemPrompt , getAuthorPromptVariant } から '@mobile-reality/mdma-prompt-pack' ;
// モデルに合わせて調整されたプロンプト バリアントを選択します (不明な場合はデフォルトに戻ります)
const { プロンプト : authorPrompt } = getAuthorPromptVariant ( 'google/gemini-2.5-pro' ) ;
// 必要に応じて、ドメイン固有の生成のためにカスタム プロンプトを最上位に重ねます
const systemPrompt = buildSystemPrompt ( {
著者プロンプト 、
CustomPrompt : `あなたはバグ追跡アシスタントです。ユーザーがバグを報告すると、
常に、この正確な構造に一致する単一のフォーム コンポーネントを生成します。
\`\`\`mdma
タイプ: フォーム
ID: バグレポート
フィールド:
- 名前: タイトル
タイプ: テキスト
ラベル: 「バグのタイトル」
必須: true
- 名前: 重大度
タイプ: 選択
ラベル: 「重大度」
オプション:
- { ラベル: クリティカル、値: クリティカル }
- { ラベル: 高、値: 高 }
- { ラベル: 中、値: 中 }
- { ラベル: 低、値: 低 }
- 名前: ステップ
タイプ: テキストエリア
ラベル: 「再現手順」
必須: true
- 名前: 期待されています
タイプ: テキストエリア
ラベル: 「期待される動作」
- 名前: 実際
タイプ: テキストエリア
ラベル: 「実際の行動」
\`\`\`` 、
} ) ;
// OpenAI 互換の API に送信します
const response = await fetch ( 'https://api.openai.com/v1/chat/completions' , {
メソッド: 'POST' 、
ヘッダー: { 認可: `Bearer ${ apiKey } ` }、
本文: JSON 。 stringify ( {
モデル: 'gemini-2.5-pro' 、
メッセージ: [
{ 役割 : 'システム' 、コンテンツ : systemPrompt } 、
{ role : 'user' 、 content : 'パスワードを入力するとログイン ページがクラッシュします。' } 、
]、
} )、
} ) ;
// LLM は ```mdma ブロックを含む通常のマークダウンで応答します
// 上記のように解析して AST + ストアを作成します
反応する
import { MdmaDocument } から '@mobile-reality/mdma-renderer-react' ;
'@mobile-reality/mdma-renderer-react/styles.css' をインポートします。 //

デフォルトのスタイル
関数 App ( { ast , ストア } ) {
return < MdmaDocument ast = { ast } ストア = { ストア } /> ;
}
注:styles.css インポートは、すべての MDMA コンポーネント (フォーム、テーブル、コールアウト、アニメーションなど) にデフォルトのスタイルを提供します。これはオプションです。代わりに、.mdma-* CSS クラスをターゲットとする独自のスタイルを作成できます。
パッケージ
説明
@モバイル現実/mdma-spec
MDMA エコシステムの基盤 — Zod スキーマ、TypeScript タイプ、および 9 つのコンポーネント タイプすべての AST 定義。他のすべてのパッケージは、検証とタイプ セーフティの仕様に依存します。
@mobile-reality/mdma-parser
標準の Markdown を MDMA 拡張 AST に変換するリマーク プラグイン。 mdma コード ブロックを抽出し、コンポーネント スキーマに対して YAML を検証し、バインディング依存関係グラフを構築します。
@モバイルリアリティ/mdma-runtime
MDMA ドキュメント用のヘッドレス状態管理エンジン — インタラクティブなドキュメントに特化したミニ状態のようなもの。リアクティブ バインディングを管理し、アクションをディスパッチし、環境ポリシーを適用し、自動 PII 編集によりすべてのイベントを改ざん明示監査ログに書き込みます。
@mobile-reality/mdma-attachables-core
9 つのコンポーネント タイプのうち 7 つのハンドラー - 状態を管理するハンドラー (フォーム、ボタン、タスクリスト、テーブル、コールアウト、承認ゲート、Webhook)。チャートと思考は表示専用であり、

[切り捨てられた]

## Original Extract

Interactive documents from Markdown. Extends MD with forms, approvals, webhooks, and more — built for next gen apps - MobileReality/mdma

GitHub - MobileReality/mdma: Interactive documents from Markdown. Extends MD with forms, approvals, webhooks, and more — built for next gen apps · GitHub
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
MobileReality
/
mdma
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
226 Commits 226 Commits .changeset .changeset .github .github assets assets blueprints blueprints demo demo docs docs evals evals examples examples packages packages skills/ mdma-integration skills/ mdma-integration .gitignore .gitignore .npmrc .npmrc CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md biome.json biome.json glama.json glama.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json turbo.json turbo.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Markdown Document with Mounted Applications
Interactive documents from Markdown. Built for next gen-apps
🚀 Live Demo
·
📖 Docs
·
💬 Discord
AI conversations today are plain text — the user reads a response and manually acts on it. MDMA changes that. When an LLM knows the MDMA spec, it can respond with interactive components (forms, tables, approval gates) instead of just text. The conversation becomes actionable: the user fills out a form, approves a step, or reviews structured data — all inline, with a predictable schema that your app already knows how to render and process.
No custom UI per use case. No parsing free-form text. The AI generates structured, validated components and your frontend renders them instantly.
MDMA extends Markdown with interactive components defined in fenced mdma code blocks. A regular Markdown file becomes an interactive application:
# Patient Intake
``` mdma
type: form
id: intake-form
fields:
- name: patient-name
type: text
label: "Full Name"
required: true
sensitive: true
- name: email
type: email
label: "Email"
required: true
sensitive: true
- name: reason
type: textarea
label: "Reason for Visit"
required: true
```
``` mdma
type: button
id: submit-btn
text: "Submit Intake Form"
variant: primary
onAction: submit
```
MDMA_AUTHOR prompt matrix
Each cell shows the pass rate of the model-specialized MDMA_AUTHOR prompt variant on the listed eval suite.
🟡 Scoring between 80–99% on the suite.
🔴 Scoring below 80% on the suite.
Don't see your model? Add a prompt variant under packages/prompt-pack/src/prompts/mdma-author/<vendor>/ and open a PR — we'll run the eval suite and add it to this table.
† gpt-5.4 intermittent duplication bug — gpt-5.4 passes one-shot evals reliably but shows a non-deterministic output duplication in multi-turn, custom-prompt, and flow evals (~7–15% of runs). The model generates a complete, correct response and then immediately re-emits the entire output verbatim, causing [duplicate-ids] validation errors. This is a known model-level issue unrelated to the prompt variant. See the OpenAI community thread for details. If this affects your use case, prefer gpt-5.5 or gpt-5.2 .
‡ gemini-3.1-pro-preview stochastic preamble loop — on ~7–15% of flow-eval runs, the model emits a chain-of-thought as visible Markdown prose (e.g. **Investigating Production Errors** repeated 3–5 times) instead of opening a ```mdma block, producing either [yaml-correctness: outside fenced block] or [duplicate-ids] errors. Per Google's official Gemini 3 prompting guide, this is a model-level behavior driven by temperature/sampling — prompt-level fixes shift which test loops rather than eliminating the loops. If deterministic flow output matters, prefer gemini-2.5-pro for production multi-step flows.
* Smaller / lower-tier models from any lab (OpenAI mini · nano, Anthropic Haiku, Google Gemini Flash, etc.) pass our eval suites, which exercise short, structured test cases. In longer real-world conversations they tend to hallucinate, forget earlier turns, or drift from the spec. For production use that involves multi-turn dialogue or stateful flows, prefer the flagship-tier model from the same family.
[i] Noticeably slow response times — single-turn responses commonly take tens of seconds and full eval runs measure in minutes.
Each cell shows the pass rate of the model-specialized MDMA_FIXER prompt variant on the single-block fixer eval (15 tests covering structural fixes, bindings, PII, forms, tables/charts, approvals). The fixer is what powers automatic repair of LLM output that fails validate() — every supported model lands at ✅ via model-tailored inline guards (no-leading-separator, preserve-input-structure, table-key-direction, replace-all-placeholders, fix-all-listed-errors, etc.).
✅ 100% on the single-block fixer eval (15/15).
‡ Reasoning-token leak suppression — for reasoning-flavoured Gemini Pro variants and Grok 4.3, the fixer would otherwise see visible "Thinking: Topic " prose prepended to every response. The eval config sets passthrough.reasoning.exclude: true (and the demo's usePreviewValidation does the same per-provider) to strip reasoning tokens from the response body at the API layer rather than at the prompt layer.
9 built-in component types, all rendered out of the box by @mobile-reality/mdma-renderer-react :
Additionally, standard Markdown content (headings, paragraphs, lists, code blocks, images, links, tables, etc.) is rendered inline between components.
The built-in chart renderer intentionally renders data as a plain table so the library stays lightweight. To get actual charts, register a custom renderer:
import { MdmaDocument } from '@mobile-reality/mdma-renderer-react' ;
import { MyRechartsRenderer } from './MyRechartsRenderer' ;
function App ( { ast , store } ) {
return (
< MdmaDocument
ast = { ast }
store = { store }
customizations = { {
components : {
chart : MyRechartsRenderer ,
} ,
} }
/>
) ;
}
This pattern works for overriding any built-in component — pass a custom React component under customizations.components.<type> .
# Core — parse and run MDMA documents
npm install @mobile-reality/mdma-parser @mobile-reality/mdma-runtime
# React rendering
npm install @mobile-reality/mdma-renderer-react
# AI authoring — system prompts for LLM-based generation
npm install @mobile-reality/mdma-prompt-pack
# Validation — static analysis for MDMA documents
npm install @mobile-reality/mdma-validator
# CLI — interactive prompt builder + document validation
npx @mobile-reality/mdma-cli
All packages are published under the @mobile-reality npm org.
import { unified } from 'unified' ;
import remarkParse from 'remark-parse' ;
import { remarkMdma } from '@mobile-reality/mdma-parser' ;
import { createDocumentStore } from '@mobile-reality/mdma-runtime' ;
import type { MdmaRoot } from '@mobile-reality/mdma-spec' ;
// 1. Parse markdown into AST
const processor = unified ( ) . use ( remarkParse ) . use ( remarkMdma ) ;
const tree = processor . parse ( markdown ) ;
const ast = ( await processor . run ( tree ) ) as MdmaRoot ;
// 2. Create a reactive document store
const store = createDocumentStore ( ast , {
documentId : 'my-doc' ,
sessionId : crypto . randomUUID ( ) ,
} ) ;
// 3. Subscribe to state changes
store . subscribe ( ( state ) => {
console . log ( 'Bindings:' , state . bindings ) ;
} ) ;
// 4. Dispatch user actions
store . dispatch ( {
type : 'FIELD_CHANGED' ,
componentId : 'intake-form' ,
field : 'patient-name' ,
value : 'Jane Doe' ,
} ) ;
In a Chat
import { buildSystemPrompt , getAuthorPromptVariant } from '@mobile-reality/mdma-prompt-pack' ;
// Pick the prompt variant tuned for your model (falls back to default if unknown)
const { prompt : authorPrompt } = getAuthorPromptVariant ( 'google/gemini-2.5-pro' ) ;
// Optionally layer a custom prompt on top for domain-specific generation
const systemPrompt = buildSystemPrompt ( {
authorPrompt ,
customPrompt : `You are a bug tracking assistant. When a user reports a bug,
always generate a single form component matching this exact structure:
\`\`\`mdma
type: form
id: bug-report
fields:
- name: title
type: text
label: "Bug Title"
required: true
- name: severity
type: select
label: "Severity"
options:
- { label: Critical, value: critical }
- { label: High, value: high }
- { label: Medium, value: medium }
- { label: Low, value: low }
- name: steps
type: textarea
label: "Steps to Reproduce"
required: true
- name: expected
type: textarea
label: "Expected Behavior"
- name: actual
type: textarea
label: "Actual Behavior"
\`\`\`` ,
} ) ;
// Send to any OpenAI-compatible API
const response = await fetch ( 'https://api.openai.com/v1/chat/completions' , {
method : 'POST' ,
headers : { Authorization : `Bearer ${ apiKey } ` } ,
body : JSON . stringify ( {
model : 'gemini-2.5-pro' ,
messages : [
{ role : 'system' , content : systemPrompt } ,
{ role : 'user' , content : 'The login page crashes after entering my password.' } ,
] ,
} ) ,
} ) ;
// The LLM responds with regular markdown containing ```mdma blocks
// Parse it into an AST + store as shown above
React
import { MdmaDocument } from '@mobile-reality/mdma-renderer-react' ;
import '@mobile-reality/mdma-renderer-react/styles.css' ; // default styles
function App ( { ast , store } ) {
return < MdmaDocument ast = { ast } store = { store } /> ;
}
Note: The styles.css import provides default styling for all MDMA components (forms, tables, callouts, animations, etc.). It's optional — you can write your own styles targeting the .mdma-* CSS classes instead.
Package
Description
@mobile-reality/mdma-spec
The foundation of the MDMA ecosystem — Zod schemas, TypeScript types, and AST definitions for all 9 component types. Every other package depends on spec for validation and type safety.
@mobile-reality/mdma-parser
A remark plugin that transforms standard Markdown into an MDMA-extended AST. Extracts mdma code blocks, validates YAML against component schemas, and builds a binding dependency graph.
@mobile-reality/mdma-runtime
Headless state management engine for MDMA documents — like a mini state specialized for interactive documents. Manages reactive bindings, dispatches actions, enforces environment policies, and writes every event to a tamper-evident audit log with automatic PII redaction.
@mobile-reality/mdma-attachables-core
Handlers for 7 of the 9 component types — the ones that manage state (form, button, tasklist, table, callout, approval-gate, webhook). Chart and thinking are display-only and

[truncated]
