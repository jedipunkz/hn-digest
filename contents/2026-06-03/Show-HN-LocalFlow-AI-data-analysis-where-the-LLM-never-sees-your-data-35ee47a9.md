---
source: "https://github.com/localflow-ai/localflow-core"
hn_url: "https://news.ycombinator.com/item?id=48370348"
title: "Show HN: LocalFlow – AI data analysis where the LLM never sees your data"
article_title: "GitHub - localflow-ai/localflow-core: The core API to enable metadata-first AI in your projects. · GitHub"
author: "renaudpawlak"
captured_at: "2026-06-03T00:44:53Z"
capture_tool: "hn-digest"
hn_id: 48370348
score: 2
comments: 0
posted_at: "2026-06-02T13:58:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LocalFlow – AI data analysis where the LLM never sees your data

- HN: [48370348](https://news.ycombinator.com/item?id=48370348)
- Source: [github.com](https://github.com/localflow-ai/localflow-core)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T13:58:42Z

## Translation

タイトル: Show HN: LocalFlow – LLM がデータを決して参照しない AI データ分析
記事のタイトル: GitHub - localflow-ai/localflow-core: プロジェクトでメタデータファースト AI を有効にするコア API。 · GitHub
説明: プロジェクトでメタデータファースト AI を有効にするコア API。 - ローカルフロー-ai/ローカルフロー-コア
HN テキスト: こんにちは、HN。私は JSweet プロジェクト (Java から TypeScript へのトランスパイラー) の作成者ですが、今日は新しいもの、LocalFlow を紹介したいと思います。 AI を使用してデータを分析するには、一般的な方法が 2 つあります (「データと対話する」)。
1. クラウド上で強力なフロンティア モデルを使用しますが、データを引き渡し、トークンごとに支払います
2. オンプレミスまたはローカル モデルを使用します – ただし、弱い結果と大幅なハードウェア コストを受け入れます LocalFlow は、私が「メタデータ ファースト AI」と呼んでいる 3 番目のアプローチを提案しています。LLM にスキーマのみを送信し (列名、型、いくつかの統計情報。実際の値は決して送信しません)、データセット全体でローカルに実行されるコードを生成するように要求します (追加のトークン コストをゼロで数百万行にスケールします)。あなたがどう思うか興味があります :) ライブデモ (CSV をドロップイン): https://apps.localflow.fr/demo/
GitHub: https://github.com/localflow-ai/localflow-core

記事本文:
GitHub - localflow-ai/localflow-core: プロジェクトでメタデータファースト AI を有効にするコア API。 · GitHub
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
ローカルフロー-ai
/
ローカルフローコア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
53 コミット 53 コミット src src .gitignore .gitignore ライセンス ライセンス README.md README.md package.json package.json proxy.js proxy.js tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LocalFlow コア — メタデータファーストの AI フレームワーク
メタデータファーストAIとは何ですか?
メタデータの境界
アーキテクチャの概要
ブラウザから何が離れるのでしょうか?
APIリファレンス
ローカルアシスタント
コンストラクター
プロキシ / ローカルプロキシ / プロキシクライアント
メタデータファースト AI では、データに関するメタデータ (列名、統計サンプル、ドキュメント構造) のみが LLM に到達します。実際の行、値、ドキュメントはマシン上に残ります。モデルはコード ジェネレーターとして機能します。データの形状の説明が与えられると、実際のデータのサンドボックスでローカルに実行される分析コードが作成されます。
これは、通常議論される 2 つのアプローチとは異なる軸です。
古典的なクラウド AI — 生データをモデルに送信します。強力で柔軟ですが、データは環境から離れ、すべての結果は新鮮な推論であり、本質的に非決定的です。
ローカル モデル AI (Ollama、llama.cpp など) — デバイス上でモデルを実行するため、データはローカルに残りますが、ハードウェアに適合するものに制約され、結果は非決定的なままになります。
メタデータファースト AI は制約を明示的にします。メタデータのみが推論境界を越え、生データは決して越えません。クラウド LLM を使用すると、メタデータ (列名、統計、ドキュメント構造) がマシンから流出しますが、これはほとんどの組織にとって許容可能な漏洩です。最も厳しいプライバシー要件については、メタデータファースト AI と自己ホスト型 LLM を組み合わせることで、それさえも排除され、インフラストラクチャの境界を越えるものは何もありません。生成されたコードは、実際のブラウザーで決定的に実行されます。

データ。
一般的な使用例は、機密性の高いエンタープライズ スプレッドシートでの「データへの対話」 (モデルに値を公開せずに自然言語クエリを実行する場合) から、大規模な地理空間分析、機密 PDF 上のドキュメント インテリジェンス、決定論的で再現可能な結果が重要な分析パイプラインまで多岐にわたります。
LocalFlow は、各データ型のメタデータを構成するものを正確に定義します。
構造化データ (CSV、Excel、CRM): 列ヘッダーと統計サンプル。LLM が正しい分析コードを作成するのに十分です。生の行は決して送信されません。
ドキュメント (PDF): LLM がドキュメントの構造を理解し、信頼できるパーサーを作成できるように、抽出されたテキストが必要です。ユーザーは、難読化されたドキュメントまたはテンプレート ドキュメントを操作して数式を生成し、それを実際のドキュメント上でローカルに実行できます。LLM には構造のみが必要で、実際の値は必要ありません。
コード生成 (ステップ 1) とローカル実行 (ステップ 2) はユーザーには見えません。アシスタントは通常の AI アシスタントと同様に動作します。違いは、生成された分析を保存し、別の LLM 呼び出しを行わずに互換性のあるデータセットで再実行できることです。
完全な AI パワー — 利用可能な最高の LLM を使用して、複雑で異種データを分析します
データの安全性 — 生データが LLM に到達することはありません。メタデータ (スキーマ、統計) はコード生成のみに使用されます。自己ホスト型 LLM はそのような危険性さえも排除します
幻覚のような結果はありません — 出力はモデルによって推論されるのではなく、決定論的なコードによって実際のデータから計算されます。
スケーラブル — 一度生成されたら、追加の AI トークンを消費せずに、大規模なデータセットに対して同じ分析を必要に応じて何度でも実行できます。
説明可能 — 生成されたコードは完全に検査可能です。 AI なら、式が機能する理由を説明したり、式が失敗する理由をデバッグしたりできます。
グリーンで持続可能 — AI はコード生成のみに使用されます。

分析ごとのst。以降の実行では AI 推論がまったく消費されないため、エネルギーを大量に消費するインフラストラクチャへの依存が軽減されます。
┌─────────────────────────┐
│ ブラウザ（クライアント） │
│ │
│ ┌─────────┐ ┌───────────┐ │
│ │ ホストアプリ / UI │ │ サンドボックス │ │
│ │ (React またはその他) │ │ 式の実行 │ │
│ │ │ │ 図表・地図 │ │
│ │ LocalAssistant │◄────►│ フェッチ→プロキシ中継 │ │
│ │ (バニラ JS) │ └─────────────┘ │
│ └───┬───┘ │
│ │ HTTPS (列統計 + 生成されたコード) │
━───────┼───────────────────┘
│
┌───▼───┐
│ LocalFlow プロキシ │ (キー、認証、
│ │ ホワイトリスト API、
└───┬───┘ エッジサービス: PDF、OCR...)
│
━━━━━━━━━━━━━━━━━━━━━━━┐
│ │ │
┌───▼───┐ ┌───────▼───────┐ ┌───────▼──────┐
│ LLM API │ │ ホワイトリストに登録 │ │ あなたの CRM / ERP / DB │
━━━━━━━┘ │ 外部API │ └─────────

───┘
━━━━━━━━┘
ブラウザから何が離れるのでしょうか?
操作
送信されたデータ
どこで
表形式の分析 — コード生成
列ヘッダー + 統計
🟠LLM
PDF 抽出
未加工の PDF バイト
🔵 プロキシ
PDF 分析 — コード生成
抽出された文書テキスト
🟠LLM
分析の実行
実際のデータ
🟢 ブラウザ
外部 API 呼び出し (オプション)
クエリパラメータのみ
🔵 プロキシ
🟢 ブラウザ — ブラウザ内に留まります · 🔵 プロキシ — サーバーのみにアクセスし、API をプロキシします · 🟠 LLM — プロキシ経由で AI モデルに転送します
LLM はメタデータのみを参照し、実際の行は参照しないため、メタデータ優先のアプローチだけでは実行できないタスクがあります。 Inference は、自然言語の理解、自由形式のコンテンツの翻訳、文書の要約、または生のテキスト全体のパターンの検出に優れています。これらの機能をデータセット全体で実行する必要がある場合、従来の AI パイプライン (データをモデルに送信するパイプライン) が引き続き適切なツールとなります。
したがって、LocalFlow は従来の AI を補完するものであり、置き換えるものではありません。これにより、古典的な AI では手の届かないユースケース、つまり、プライバシーにさらされることのない大規模なデータ分析、決定論的で再現可能な結果、限界 AI コストゼロでのスケーラブルな実行が開かれる一方で、真に必要とされる古典的なアプローチの余地も残されています。
その境界は固定されていません。プロキシは、管理者またはユーザーによって定義および制御され、慎重に範囲が限定されたデータのサブセットに対して動作するツール (LLM を利用したツールを含む) を公開できます。たとえば、数式は、LLM が完全なデータセットを参照することなく、特定のフィールドを要約したり、列を変換したりするプロキシ ホスト型サービスを呼び出すことができます。この種の拡張機能には、プロキシ環境の意図的な構成が必要です

nment とその利用可能なツールを利用して、特定のユースケースと許容可能なデータ共有境界に合わせてセットアップを調整します。
実際に動作しているところを見てみたいですか? LocalFlow オンライン アシスタントの例を試してください。インストールは必要ありません。ソース: localflow-examples 。
📄 これらの概念についてさらに詳しく知りたい場合は、LocalFlow ホワイト ペーパーをお読みください。
クイック スタート — LocalAssistant をアプリに埋め込む
サーバーは必要ありません — LocalProxy は完全にブラウザ内で実行され、LLM を直接呼び出します。すぐに実行できる完全な React アプリとバニラ JS アプリについては、localflow-examples を参照してください。
npm インストール @localflow/core
2. アシスタントを作成する
import { LocalProxy , LocalAssistant } から '@localflow/core'
// サーバーは必要ありません — LocalProxy はブラウザから LLM を直接呼び出します
const proxy = new LocalProxy ()
const Assistant = 新しい LocalAssistant ( {
代理、
llm : {
タイプ: 'ジェミニ' 、
model : 'gemini-3-flash-preview' , // オプション、これがデフォルトです
} 、
// 数式の結果を表示する div をポイントします。
// HTMLElement、CSS セレクター文字列、またはファクトリー関数を受け入れます。
resultContainer : '#result' 、
})
// ユーザーの Gemini API キーを渡します — ローカルに保存され、第三者に送信されることはありません
// 注: 実際のプロキシを使用すると自動的に暗号化されます
アシスタントを待ちます。 setLlmApiKey ( 'AIza...' )
// LLM 構成が変更されるたびに永続化します (ユーザーが新しいキー、モデルなどを設定する)
アシスタント。 on ( 'llm:change' , ( llm ) => {
if (llm . apiKey ) localStorage . setItem ( 'llm-key' , llm . apiKey )
})
ページの読み込み全体でキーを復元するには、構築時にキーを渡します。
const Assistant = 新しい LocalAssistant ( {
代理、
llm : { タイプ : 'gemini' 、apiKey : localStorage 。 getItem ( 'llm-key' ) ?? '' } 、
resultContainer : '#result' 、
})
3. データをロードします
// 表形式のデータ — あらゆるソースからのデータ: CSV 解析、DB クエリ、API 応答など。
お尻

すぐに。 addDataset ( 'portfolio' , portfolioRows ) // 行: Record<文字列, 不明>[]
アシスタント。 addDataset ( 'market' 、marketRows )
// どのデータセットが「アクティブ」であるかをマークします (式の `data` 変数)
アシスタント。 setActiveDataset ( 'ポートフォリオ' )
4. メッセージを送信して結果をレンダリングする
resultContainer が構成されている場合、アシスタントは iframe の作成、サンドボックス権限の設定、プロキシされた API 呼び出しの中継など、すべてを処理します。定型文は必要ありません。
アシスタント。 on ( 'メッセージ' , ( 応答 ) => {
// チャット UI に response.answer を表示します
appendChatBubble (応答.回答)
})
アシスタント。 on ( 'formula:done' , ( { data } ) => {
// オプション: 式の出力データに反応します
コンソール。 log ( '分析結果:' , data )
})
// メッセージを送信します — LLM が数式を生成し、アシスタントが
// 設定された resultContainer 内で自動的にレンダリングされます。
const 応答 = アシスタントを待ちます。プロンプト (「資産クラス別の配分を表示」)
結果コンテナはいつでも変更できます。
アシスタント。 resultContainer = ドキュメント 。 getElementById ( '結果' )
// または:assistant.resultContainer = '#result'
// または:assistant.resultContainer = () => document.querySelector('.panel.active')
5. 数式の保存と再生 (オプション)
生成された式は決定的です。追加の LLM 呼び出しを行わずに、それらを保存し、互換性のあるデータセットで再実行できます。
// 最後の応答から式を取得します
const 式 = アシスタント 。 getLastFormula ()
// カタログに保存します...
saveToCatalog ({ 式 , タイトル : 応答 . タイトル , 説明 : 応答 . 説明 } )
// ...後で再実行します
アシスタント。 executeFormula (savedFormula)
6. 意味解析一致フックの登録（オプション）
これにより、アシスタントは関連する過去の分析をシステム プロンプトとして挿入できるようになります。

LLM を呼び出す前にサンプルを実行し、出力品質を向上させます。
アシスタント。 setAnalysisMatchHook ( async ( query , ctx ) => {
const Analysis = categoryLoad ( ) // カタログ ストア
return findBestMatch (クエリ、分析、ctx .activeColumns)
})
実際のプロキシを使用する
LocalProxy は、ローカル開発と迅速なプロトタイピング用に設計されています。本番環境向け - セッション管理、API ガバナンス、PDF 抽出、BYOK キー暗号化、RAT が必要な場合

[切り捨てられた]

## Original Extract

The core API to enable metadata-first AI in your projects. - localflow-ai/localflow-core

Hi HN, I'm the author of the JSweet project (Java-to-TypeScript transpiler), but today I'd like to show you something new: LocalFlow. There are two common ways to analyse your data with AI ("Talk to your Data"):
1. Use powerful frontier models on the cloud — but hand over your data and pay per token
2. Use on-premise or local models — but accept weaker results and significant hardware costs LocalFlow proposes a third approach I call "metadata-first AI": send only the schema to the LLM (column names, types, a few stats — never the actual values), and ask it to generate code that runs locally on your full dataset (scale it to millions of rows at zero additional token cost). I'm curious what you think :) Live demo (drop in any CSV): https://apps.localflow.fr/demo/
GitHub: https://github.com/localflow-ai/localflow-core

GitHub - localflow-ai/localflow-core: The core API to enable metadata-first AI in your projects. · GitHub
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
localflow-ai
/
localflow-core
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
53 Commits 53 Commits src src .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json proxy.js proxy.js tsconfig.json tsconfig.json View all files Repository files navigation
LocalFlow Core — Metadata-first AI Framework
What is Metadata-first AI?
The metadata boundary
Architecture overview
What leaves the browser?
API Reference
LocalAssistant
Constructor
Proxy / LocalProxy / ProxyClient
In metadata-first AI , only metadata about your data ever reaches the LLM — column names, statistical samples, document structure. The actual rows, values, and documents stay on your machine. The model acts as a code generator : given a description of the data's shape, it writes analysis code that executes locally in a sandbox on your real data.
This is a different axis than the two approaches usually discussed:
Classical cloud AI — sends raw data to the model. Powerful and flexible, but your data leaves your environment and every result is a fresh inference, non-deterministic by nature.
Local-model AI (Ollama, llama.cpp, etc.) — runs the model on your device so data stays local, but you are constrained by what fits on your hardware, and results remain non-deterministic.
Metadata-first AI makes the constraint explicit: only metadata crosses the inference boundary — raw data never does. When using a cloud LLM, metadata (column names, statistics, document structure) does leave your machine, which is an acceptable exposure for most organisations. For the strictest privacy requirements, combining metadata-first AI with a self-hosted LLM eliminates even that: nothing crosses your infrastructure boundary. The generated code runs deterministically in your browser on your real data.
Typical use cases range from "Talk to your Data" on sensitive enterprise spreadsheets — where you want natural language querying without exposing values to the model — to large-scale geospatial analysis, document intelligence on confidential PDFs, and any analytical pipeline where deterministic, repeatable results matter.
LocalFlow defines precisely what constitutes metadata for each data type:
Structured data (CSV, Excel, CRM): column headers and statistical samples — enough for the LLM to write correct analysis code. Raw rows are never sent.
Documents (PDF): the extracted text is needed so the LLM understands the document's structure and can write a reliable parser. Users can work with obfuscated or template documents to generate formulas, then run them locally on real documents — the LLM only needs the structure, not the actual values.
Code generation (step 1) and local execution (step 2) are invisible to the user — the assistant behaves like a regular AI assistant. The difference is that generated analyses can be saved and re-run on any compatible dataset without making another LLM call.
Full AI power — use the best available LLM to analyse complex, heterogeneous data
Data safety — raw data never reaches the LLM. Metadata (schema, statistics) leaves only for code generation; a self-hosted LLM removes even that exposure
No hallucinated results — outputs are computed from real data by deterministic code, not inferred by the model
Scalable — once generated, run the same analysis on large datasets as many times as needed, consuming no additional AI tokens
Explainable — the generated code is fully inspectable; any AI can explain why a formula works or debug why it fails
Green and sustainable — AI is used only for code generation, a one-time cost per analysis. Subsequent runs consume no AI inference at all, reducing dependence on energy-intensive infrastructure
┌─────────────────────────────────────────────────────┐
│ Browser (client) │
│ │
│ ┌─────────────────┐ ┌──────────────────────┐ │
│ │ Host App / UI │ │ Sandbox │ │
│ │ (React or any) │ │ formula execution │ │
│ │ │ │ charts / maps │ │
│ │ LocalAssistant │◄────►│ fetch → proxy relay │ │
│ │ (vanilla JS) │ └──────────────────────┘ │
│ └────────┬────────┘ │
│ │ HTTPS (column stats + generated code) │
└───────────┼─────────────────────────────────────────┘
│
┌────────▼────────┐
│ LocalFlow Proxy │ (manages keys, auth,
│ │ whitelists APIs,
└────────┬────────┘ edge services: PDF, OCR...)
│
├──────────────────┐─────────────────────┐
│ │ │
┌────────▼────────┐ ┌───────▼───────┐ ┌───────────▼───────────┐
│ LLM API │ │ Whitelisted │ │ Your CRM / ERP / DB │
└─────────────────┘ │ external APIs │ └───────────────────────┘
└───────────────┘
What leaves the browser?
Operation
Data sent
Where
Tabular analysis — code generation
Column headers + statistics
🟠 LLM
PDF extraction
Raw PDF bytes
🔵 Proxy
PDF analysis — code generation
Extracted document text
🟠 LLM
Analysis execution
Actual data
🟢 Browser
External API calls (optional)
Query parameters only
🔵 Proxy
🟢 Browser — stays in your browser · 🔵 Proxy — goes to your server only and proxied APIs · 🟠 LLM — forwarded to the AI model via your proxy
Because the LLM only ever sees metadata — never the actual rows — there are tasks that a metadata-first approach cannot perform on its own. Inference excels at understanding natural language, translating free-form content, summarising documents, or spotting patterns across raw text. When those capabilities need to operate on the full dataset, a classical AI pipeline (one that sends the data to the model) remains the right tool.
LocalFlow is therefore a complement to classical AI, not a replacement . It opens up use cases that are out of reach for classical AI — large-scale data analysis without privacy exposure, deterministic and repeatable results, scalable execution at zero marginal AI cost — while leaving room for classical approaches where they are genuinely needed.
That boundary is not fixed. The proxy can expose tools — including LLM-powered ones — that operate on a carefully scoped subset of data, defined and controlled by the administrator or the user. A formula could, for example, call a proxy-hosted service that summarises a specific field or translates a column, without the LLM ever seeing the full dataset. This kind of extension requires intentional configuration of your proxy environment and its available tools, tailoring the setup to your specific use cases and acceptable data-sharing boundaries.
Want to see it in action? Try the LocalFlow online assistant example — no installation needed. Source: localflow-examples .
📄 For a deeper dive into these concepts, read the LocalFlow white paper .
Quick start — embedding LocalAssistant in your app
No server required — LocalProxy runs entirely in the browser and calls the LLM directly. See localflow-examples for complete React and vanilla JS apps you can run immediately.
npm install @localflow/core
2. Create the assistant
import { LocalProxy , LocalAssistant } from '@localflow/core'
// No server needed — LocalProxy calls the LLM directly from the browser
const proxy = new LocalProxy ( )
const assistant = new LocalAssistant ( {
proxy ,
llm : {
type : 'gemini' ,
model : 'gemini-3-flash-preview' , // optional, this is the default
} ,
// Point to the div where formula results should be rendered.
// Accepts an HTMLElement, a CSS selector string, or a factory function.
resultContainer : '#result' ,
} )
// Pass the user's Gemini API key — stored locally, never sent to any third party
// NOTE: automatically encryted when using an actual proxy
await assistant . setLlmApiKey ( 'AIza...' )
// Persist LLM config whenever it changes (user sets a new key, model, etc.)
assistant . on ( 'llm:change' , ( llm ) => {
if ( llm . apiKey ) localStorage . setItem ( 'llm-key' , llm . apiKey )
} )
To restore a key across page loads, pass it at construction:
const assistant = new LocalAssistant ( {
proxy ,
llm : { type : 'gemini' , apiKey : localStorage . getItem ( 'llm-key' ) ?? '' } ,
resultContainer : '#result' ,
} )
3. Load your data
// Tabular data — from any source: CSV parse, DB query, API response, etc.
assistant . addDataset ( 'portfolio' , portfolioRows ) // rows: Record<string, unknown>[]
assistant . addDataset ( 'market' , marketRows )
// Mark which dataset is the "active" one (the `data` variable in formulas)
assistant . setActiveDataset ( 'portfolio' )
4. Send a message and render the result
With resultContainer configured, the assistant takes care of everything — creating the iframe, setting sandbox permissions, and relaying proxied API calls. No boilerplate needed.
assistant . on ( 'message' , ( response ) => {
// Show response.answer in your chat UI
appendChatBubble ( response . answer )
} )
assistant . on ( 'formula:done' , ( { data } ) => {
// Optional: react to the formula's output data
console . log ( 'Analysis result:' , data )
} )
// Send a message — the LLM generates a formula, and the assistant
// renders it automatically in the configured resultContainer.
const response = await assistant . prompt ( 'Show me the allocation by asset class' )
You can change the result container at any time:
assistant . resultContainer = document . getElementById ( 'result' )
// or: assistant.resultContainer = '#result'
// or: assistant.resultContainer = () => document.querySelector('.panel.active')
5. Save and replay formulas (optional)
Generated formulas are deterministic — you can save them and re-run on any compatible dataset without an additional LLM call.
// Get the formula from the last response
const formula = assistant . getLastFormula ( )
// Save it to your catalog...
saveToCatalog ( { formula , title : response . title , description : response . description } )
// ...and replay it later
assistant . executeFormula ( savedFormula )
6. Register a semantic analysis-match hook (optional)
This lets the assistant inject a relevant past analysis as a system-prompt example before calling the LLM, improving output quality.
assistant . setAnalysisMatchHook ( async ( query , ctx ) => {
const analyses = catalogLoad ( ) // your catalog store
return findBestMatch ( query , analyses , ctx . activeColumns )
} )
Use an actual proxy
LocalProxy is designed for local development and quick prototyping. For production — where you need session management, API governance, PDF extraction, BYOK key encryption, rat

[truncated]
