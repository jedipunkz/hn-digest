---
source: "https://github.com/janbalangue/async-bulkhead-llm"
hn_url: "https://news.ycombinator.com/item?id=49278841"
title: "Show HN: Keep batch LLM jobs from starving interactive traffic (TypeScript)"
article_title: "GitHub - janbalangue/async-bulkhead-llm: Fail-fast admission control for LLM workloads, with concurrency limits, token budgeting, deduplication, streaming support, and overload protection. · GitHub"
author: "janbalangue"
captured_at: "2026-08-12T21:35:24Z"
capture_tool: "hn-digest"
hn_id: 49278841
score: 1
comments: 0
posted_at: "2026-08-12T21:30:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Keep batch LLM jobs from starving interactive traffic (TypeScript)

- HN: [49278841](https://news.ycombinator.com/item?id=49278841)
- Source: [github.com](https://github.com/janbalangue/async-bulkhead-llm)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T21:30:56Z

## Translation

タイトル: HN を表示: バッチ LLM ジョブが対話型トラフィックを枯渇させないようにする (TypeScript)
記事のタイトル: GitHub - janbalangue/async-bulkhead-llm: 同時実行制限、トークン バジェット、重複排除、ストリーミング サポート、過負荷保護を備えた LLM ワークロードのフェイルファスト アドミッション コントロール。 · GitHub
説明: 同時実行制限、トークン バジェット、重複排除、ストリーミング サポート、および過負荷保護を備えた、LLM ワークロードのフェイルファスト アドミッション コントロール。 - ジャンバランゲ/async-bulkhead-llm

記事本文:
GitHub - janbalangue/async-bulkhead-llm: 同時実行制限、トークン バジェット、重複排除、ストリーミング サポート、および過負荷保護を備えた、LLM ワークロードのフェイルファスト アドミッション コントロール。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ジャンバランゲ
/
非同期バルクヘッド-llm
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
41 コミット 41 コミット .github/ workflows .github/ workflows src src test test .gitignore .giti

無視 CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json rename-cjs.mjs rename-cjs.mjs tsconfig.cjs.json tsconfig.cjs.json tsconfig.eslint.json tsconfig.eslint.json tsconfig.esm.json tsconfig.esm.json tsconfig.json tsconfig.json tsconfig.types.json tsconfig.types.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリファイルナビゲーション
async-bulkhead-ts に基づいて構築された LLM ワークロードのフェイルファースト アドミッション コントロール。
リクエストのファンアウト前、プロバイダーのレート制限に達する前、飽和カスケードに達する前に、コスト上限、同時実行制限、LLM 呼び出しの境界でバックプレッシャーを適用する必要があるサービス向けに設計されています。
✅ ハードマックスのインフライト同時実行数 ( maxConcurrent )
✅ アトミックなバージョン制限の更新 — 同時実行性、キュー、トークンバジェット、優先予約、および入場クラスの下限/上限を 1 つの改訂されたスナップショットとして適用します
✅ 保護されたアドミッションクラスフロア — リクエストが共有の残りから借用される前に、クラスごとの同時実行性とトークン容量を予約します。
✅ 制限されたアドミッション クラス — クラスごとのハード同時実行性とインフライト トークンの上限を使用してポリシー クラスを分離します。
✅ 承認に応じて線形化されたリビジョン — 承認またはバイパスされたすべてのライフサイクルは、決定を行った正確な制限リビジョンを保持します。
✅ トークン対応アドミッション - 推定入力 + 最大出力トークンに対する予備
✅ プログレッシブ調整 — ライブストリーム中に処理された入力と生成された出力の予約を解放します。
✅ トークンの払い戻し — 完了後の実際の使用量から未使用の予算容量を回収します。
✅ モデルを意識した推定 — 既知のプロバイダーのモデルごとの文字比率
✅ リクエストごとのモデル — 単一のルーティングを介した混合モデルのルーティング

隔壁
✅ マルチモーダルコンテンツ — テキストブロックがカウントされます。不透明なブロックはデフォルトで無視されるか、構成可能なブロックごとの予約 ( opaqueBlockTokens ) が課金されます。
✅ フルリクエストの推定 — システムプロンプトはネイティブにカウントされます。 extraInputTokens は呼び出し元が計算したコスト (ツール スキーマ、プロバイダー価格のメディア) を負担します。
✅ コールごとの予約オーバーライド — 独自の推定値を持つゲートウェイは、acquire() / run() / wouldAdmit() の予約を介して推定器をバイパスできます。
✅ デフォルトでフェイルファスト — ロードを早期に遮断し、サイレントキューに入れません
✅ オブザーブ/シャドウ モード — 選択された拒否される可能性のあるコールを実行しながら、キャパシティ ポリシーへの影響を測定します。
✅ 独自のプロファイル — エスケープハッチを備えた「インタラクティブ」および「バッチ」プリセット
✅ 実行中の重複排除 — 同一のリクエストが 1 つの LLM 呼び出しを共有します。ハッシュされたキー、リクエスト全体の等価性
✅ カスタム重複排除キー + テナントごとのスコープ — 独自の等価関数を導入します。 dedupScope はテナントを分離します
✅ ストリーミングセーフな重複排除 — 単一消費者の結果がサイレントに共有されることはありません。 shareResult ファンアウト フック + コールごとの重複排除: false オプトアウト
✅ 正確な予約プレビュー —estimate() は使用する予約受付を公開し、その結果を通話ごとの予約としてそのまま返すことができます。
✅ 適応推定 — createAdaptiveTokenEstimator() は、観察された使用状況からモデルごとの入力推定を自己調整します
✅ アドバイザリー容量スナップショット — wouldAdmit(request, {detail: true }) は、拒否された場合と同じ容量番号を返します。
✅ 安定したアドミッション ID — ゲートウェイのリクエスト、トレース、使用状況、リリースを関連付けます。
✅ 順序付けされた使用状況イベント — 外部コーディネーター向けのシーケンス番号付きの保留変更
✅ イベントフック - 許可されたライフサイクルとバイパスされたライフサイクルには個別のテレメトリ イベントがあります
✅ 正常なシャットダウン — close() + dry() 、opti 付き

限定された待機: ドレイン({ timeoutMs })
✅ オプションの AbortSignal とタイムアウトのキャンセル
✅Bulkhead.run(request, fn) — 取得と解放が自動的に処理されます
✅ async-bulkhead-ts を超える依存関係はゼロ
❌ プロバイダー SDK なし - 独自のクライアントを使用
❌ コスト計算なし — トークンの推定は負荷制限のみを目的としています
競合マトリックス (LLM ワークロード)
機能 / ライブラリ
非同期バルクヘッド-llm
ラングチェーン / ラマインデックス
OpenAI SDK (生)
pリミット/ボトルネック
オカメインコ / ポリー
主な目標
LLM アドミッション コントロール (コスト + 同時実行性)
オーケストレーション/パイプライン
APIクライアント
同時実行性/スケジューリング
回復力のパターン
デフォルトでフェイルファスト
✅ はい
❌ いいえ
❌ いいえ
❌ いいえ
⚠️場合による
トークンを意識したアドミッション
✅ はい（入学前予算）
❌ いいえ
❌ いいえ
❌ いいえ
❌ いいえ
トークンの払い戻し（コール後の修正）
✅ はい
❌ いいえ
❌ いいえ
❌ いいえ
❌ いいえ
コスト上限の強制
✅ はい ( tokenBudget )
❌ いいえ
❌ いいえ
❌ いいえ
❌ いいえ
同時実行の制限
✅ はい
⚠️間接的
❌ いいえ
✅ はい
⚠️間接的
有界キュー (オプション)
✅ はい
⚠️内部
❌ いいえ
✅ はい
⚠️間接的
フェイルファスト過負荷処理
✅ コア機能
❌ いいえ
❌ いいえ
❌ いいえ
⚠️間接的
観察/シャドウ ロールアウト モード
✅ はい
❌ いいえ
❌ いいえ
❌ いいえ
❌ いいえ
飛行中の重複排除
✅ はい
⚠️ 部分キャッシュ
❌ いいえ
❌ いいえ
❌ いいえ
カスタム重複排除キー
✅ はい
❌ いいえ
❌ いいえ
❌ いいえ
❌ いいえ
モデルを意識した推定
✅ はい
❌ いいえ
❌ いいえ
❌ いいえ
❌ いいえ
リクエストごとのモデルルーティング
✅ はい
✅ はい
✅ はい
❌ いいえ
❌ いいえ
マルチモーダルを考慮した推定
✅ はい (テキストのみカウントされます)
❌ いいえ
❌ いいえ
❌ いいえ
❌ いいえ
アボート/タイムアウト（受付）
✅ はい
⚠️部分的
⚠️ SDKレベル
⚠️部分的
✅ はい
イベントフック (メトリクス/ロギング)
✅ はい
❌ いいえ
❌ いいえ
❌ いいえ
⚠️限定
正常なシャットダウン (ドレイン/クローズ)
✅ はい
❌ いいえ
❌ いいえ
❌ いいえ
❌ いいえ
再試行/フォールバック

k
❌ いいえ
⚠️はい
❌ いいえ
❌ いいえ
✅ はい
LLM オーケストレーション (チェーン/エージェント)
❌ いいえ
✅ はい
❌ いいえ
❌ いいえ
❌ いいえ
素早い位置決め
LangChain / LlamaIndex → 何を実行するか (オーケストレーション)
OpenAI SDK → プロバイダーの呼び出し方法
p-limit / ボトルネック → 実行されるタスクの数
オカメインコ / ポリー → 失敗したらどうなるか
async-bulkhead-llm → リクエストを実行する必要があるかどうか
LLM パイプラインを構築する場合は、LangChain を使用します。
負荷がかかっている状態でシステムと予算を保護したい場合は、async-bulkhead-llm を使用します。
ほとんどの LLM ツールは実行を最適化します。
async-bulkhead-llm は、負荷がかかった状態での生存を最適化します。
リクエストがプロバイダーに届く前に。
npm インストール async-bulkhead-llm
クイックスタート
'async-bulkhead-llm' から { createLLMBulkhead } をインポートします。
const バルクヘッド = createLLMBulkhead ( {
モデル: 'クロード-ソネット-4' 、
maxConcurrent : 10 、
} ) ;
const リクエスト = {
メッセージ : [ { 役割 : 'ユーザー' 、内容 : 'この文書の要約...' } ] 、
max_tokens : 1024 、
} ;
const result = バルクヘッドを待ちます。 run ( request , async ( ) => {
callYourLLMProvider (リクエスト) を返します。
} ) ;
v3.15 の新機能 — 保護された入場クラスのフロア
v3.15 では、静的に設定された各アドミッション クラスで厳密なローカル クラスを予約できます。
同時実行フロア、およびトークンアドミッションが有効な場合は厳密なインフライト
トークンフロア。すべてのフロアの収容人数は、どのクラスよりも共有される残りの部分を形成します。
は、そのハードクラスの上限とそれに含まれるグローバル制限に従って借入することができます。
const バルクヘッド = createLLMBulkhead ( {
モデル：「gpt-4o」、
maxConcurrent : 12 、
tokenBudget : { 予算 : 24_000 } 、
入学クラス : {
デフォルトクラス : "標準" 、
クラス: {
プレミアム: {
保護された同時実行: 6 、
maxConcurrent : 10 、
protectedInFlightTokens : 12_000 、
maxInFlightTokens : 20_000 、
} 、
標準: {
保護された同時実行: 2 、
maxConcurrent : 6 、
プロ

tectedInFlightTokens : 4_000 、
maxInFlightTokens : 10_000 、
} 、
} 、
} 、
} ) ;
この例では、8 つの同時実行スロットと 16,000 のトークンが保護されます。の
残りの 4 つのスロットと 8,000 トークンは共有されます。プレミアムは下限を超える可能性があります
その共有された剰余から借用することによって、ただし 2 つのスロットを消費することはできません
または標準で保護された 4,000 トークン。
フロアは厳密な予約であり、空き容量の貸し出しではありません。未使用の保護済み
容量は、その所有クラスのみが利用可能なままになります。クラスを越えた融資
アイドル フロアには、より高いリビジョンをインストールするためのゲートウェイまたはコントロール プレーンが必要です
スナップショット。下限を増やしてもアクティブな作業が取り消されることはありません: 新たな借用は一時停止されます
完了するまでは減耗により保証が回復されます。
stats().admissionClasses は、共有剰余に各クラスの剰余を加えたものを報告するようになりました。
保護および借用された使用法。別のクラスの保存によって引き起こされる拒否
フロア使用制約:「admission_class_protection」、共有を含む
容量のスナップショット。両方のフロアフィールドを省略すると、v3.14 のアドミッションが保持されます。
行動。
v3.14 の新機能 — 制限付き入場クラス
v3.14 では、1 つのバルクヘッド内に有界階層アドミッション レイヤが追加されます。あ
信頼できるゲートウェイは、テナントまたはアプリケーション ID を小さな静的セットにマッピングできます。
ポリシー クラスの管理を行い、ハード同時実行性とインフライト トークンの上限を適用します。
周囲のグローバル制限を維持しながら、各クラスに対して。
const バルクヘッド = createLLMBulkhead ( {
モデル：「gpt-4o」、
maxConcurrent : 12 、
tokenBudget : { 予算 : 24_000 } 、
入学クラス : {
デフォルトクラス : "標準" 、
クラス: {
プレミアム: { maxConcurrent : 8 、 maxInFlightTokens : 18_000 } 、
標準: { maxConcurrent : 4 、 maxInFlightTokens : 6_000 } 、
} 、
} 、
} ) ;
バルクヘッドを待ちます。 run ( request , callProvider , {
入場クラス : "プレミアム" 、
dedupScope :trustedTenantId 、
} ) ;
c

ラステーブルは建設時に固定されます。不明な ID は作成されずにスローされます
無制限の状態またはメトリック カーディナリティ。クラスの上限はフェイルファストであり、
グローバル同時実行性、優先度調整されたトークン バジェット、および
シャットダウン状態。 applyLimits() はクラス番号をアトミックに更新できますが、
スナップショットには構築時のクラス キーが正確に含まれている必要があります。
インフライト重複排除が有効になっている場合、選択されたアドミッション クラスは
自動重複排除パーティション。プレミアムと
したがって、標準は最初のファイルを共有するのではなく、独立して実行されます。
クラスの入場決定またはキューの位置。 dedupScope は引き続き必要です
クラス内のテナントレベルの分離用。
アドミッション クラスは、生のテナント ID ではなく、意図的にポリシー バケットです。の
ライブラリは呼び出し元の認証、ID の割り当て、重み付けの実装を行いません。
公平性、貸し出し可能フロアの予約、またはプロセス全体にわたるキャパシティの調整。
これらの責任はゲートウェイとコントロール プレーンに残ります。いつ
飛行中の重複排除が有効になっており、admissionClass も置き換えられません
dedupスコープ ;結果が必要な場合、信頼できるテナント ID を dedupScope に含める
テナント間で共有することはできません。
ライフサイクル結果、実行コンテキスト、トークン、使用状況レポート、容量の詳細、
テレメトリ イベントは、構成時に admissionClass を伝えるようになりました。拒否の詳細
また、バインディング制約がグローバルかどうかも識別します。
admission_class 、および stats().admissionClasses は、クラスごとに制限されたものを公開します
会計。
v3.13 の新機能 — プログレッシブ調整
受信するゲートウェイ

[切り捨てられた]

## Original Extract

Fail-fast admission control for LLM workloads, with concurrency limits, token budgeting, deduplication, streaming support, and overload protection. - janbalangue/async-bulkhead-llm

GitHub - janbalangue/async-bulkhead-llm: Fail-fast admission control for LLM workloads, with concurrency limits, token budgeting, deduplication, streaming support, and overload protection. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
janbalangue
/
async-bulkhead-llm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
41 Commits 41 Commits .github/ workflows .github/ workflows src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json rename-cjs.mjs rename-cjs.mjs tsconfig.cjs.json tsconfig.cjs.json tsconfig.eslint.json tsconfig.eslint.json tsconfig.esm.json tsconfig.esm.json tsconfig.json tsconfig.json tsconfig.types.json tsconfig.types.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Fail-fast admission control for LLM workloads , built on async-bulkhead-ts .
Designed for services that need to enforce cost ceilings, concurrency limits, and backpressure at the boundary of their LLM calls — before request fan-out, before hitting provider rate limits, before saturation cascades.
✅ Hard max in-flight concurrency ( maxConcurrent )
✅ Atomic versioned limit updates — apply concurrency, queue, token budget, priority reserve, and admission-class floors/ceilings as one revisioned snapshot
✅ Protected admission-class floors — reserve per-class concurrency and token capacity before requests borrow from the shared remainder
✅ Bounded admission classes — isolate policy classes with hard per-class concurrency and in-flight token ceilings
✅ Admission-linearized revisions — every admitted or bypassed lifecycle retains the exact limit revision that made the decision
✅ Token-aware admission — reserves against estimated input + max output tokens
✅ Progressive reconciliation — release processed input and generated-output reservation during live streams
✅ Token refund — reclaim unused budget capacity from actual usage post-completion
✅ Model-aware estimation — per-model character ratios for known providers
✅ Per-request model — mixed-model routing through a single bulkhead
✅ Multimodal content — text blocks counted; opaque blocks ignored by default or charged a configurable per-block reservation ( opaqueBlockTokens )
✅ Full-request estimation — system prompts counted natively; extraInputTokens carries caller-computed costs (tool schemas, provider-priced media)
✅ Per-call reservation override — gateways with their own estimate can bypass the estimator via reservation on acquire() / run() / wouldAdmit()
✅ Fail-fast by default — shed load early, never silently queue
✅ Observe/shadow mode — measure capacity-policy impact while still executing selected would-be-rejected calls
✅ Opinionated profiles — 'interactive' and 'batch' presets with escape hatch
✅ In-flight deduplication — identical requests share one LLM call; hashed keys, whole-request equality
✅ Custom dedup key + per-tenant scope — bring your own equivalence function; dedupScope isolates tenants
✅ Streaming-safe deduplication — single-consumer results are never silently shared; shareResult fan-out hook + per-call dedup: false opt-out
✅ Exact reservation preview — estimate() exposes the reservation admission will use, and its result can be passed back verbatim as the per-call reservation
✅ Adaptive estimation — createAdaptiveTokenEstimator() self-calibrates per-model input estimates from observed usage
✅ Advisory capacity snapshots — wouldAdmit(request, { detail: true }) returns the same capacity numbers rejections carry
✅ Stable admission IDs — correlate gateway requests, traces, usage, and release
✅ Ordered usage events — sequence-numbered hold changes for external coordinators
✅ Event hooks — admitted and bypassed lifecycles have separate telemetry events
✅ Graceful shutdown — close() + drain() , with an optional bounded wait: drain({ timeoutMs })
✅ Optional AbortSignal and timeout cancellation
✅ bulkhead.run(request, fn) — acquire + release handled automatically
✅ Zero dependencies beyond async-bulkhead-ts
❌ No provider SDK — bring your own client
❌ No cost accounting — token estimation is for load-shedding only
Competitive Matrix (LLM Workloads)
Capability / Library
async-bulkhead-llm
LangChain / LlamaIndex
OpenAI SDK (raw)
p-limit / Bottleneck
cockatiel / polly
Primary goal
LLM admission control (cost + concurrency)
Orchestration / pipelines
API client
Concurrency / scheduling
Resilience patterns
Fail-fast by default
✅ Yes
❌ No
❌ No
❌ No
⚠️ Depends
Token-aware admission
✅ Yes (pre-admission budget)
❌ No
❌ No
❌ No
❌ No
Token refund (post-call correction)
✅ Yes
❌ No
❌ No
❌ No
❌ No
Cost ceiling enforcement
✅ Yes ( tokenBudget )
❌ No
❌ No
❌ No
❌ No
Concurrency limits
✅ Yes
⚠️ Indirect
❌ No
✅ Yes
⚠️ Indirect
Bounded queue (optional)
✅ Yes
⚠️ Internal
❌ No
✅ Yes
⚠️ Indirect
Fail-fast overload handling
✅ Core feature
❌ No
❌ No
❌ No
⚠️ Indirect
Observe/shadow rollout mode
✅ Yes
❌ No
❌ No
❌ No
❌ No
In-flight deduplication
✅ Yes
⚠️ Partial caching
❌ No
❌ No
❌ No
Custom dedup key
✅ Yes
❌ No
❌ No
❌ No
❌ No
Model-aware estimation
✅ Yes
❌ No
❌ No
❌ No
❌ No
Per-request model routing
✅ Yes
✅ Yes
✅ Yes
❌ No
❌ No
Multimodal-aware estimation
✅ Yes (text-only counted)
❌ No
❌ No
❌ No
❌ No
Abort / timeout (admission)
✅ Yes
⚠️ Partial
⚠️ SDK-level
⚠️ Partial
✅ Yes
Event hooks (metrics/logging)
✅ Yes
❌ No
❌ No
❌ No
⚠️ Limited
Graceful shutdown (drain/close)
✅ Yes
❌ No
❌ No
❌ No
❌ No
Retries / fallback
❌ No
⚠️ Yes
❌ No
❌ No
✅ Yes
LLM orchestration (chains/agents)
❌ No
✅ Yes
❌ No
❌ No
❌ No
Quick positioning
LangChain / LlamaIndex → what to run (orchestration)
OpenAI SDK → how to call the provider
p-limit / Bottleneck → how many tasks run
cockatiel / polly → what happens after failure
async-bulkhead-llm → whether a request should run at all
If you want to build LLM pipelines , use LangChain.
If you want to protect your system and budget under load , use async-bulkhead-llm.
Most LLM tooling optimizes execution.
async-bulkhead-llm optimizes survival under load.
before a request ever reaches your provider.
npm install async-bulkhead-llm
Quick Start
import { createLLMBulkhead } from 'async-bulkhead-llm' ;
const bulkhead = createLLMBulkhead ( {
model : 'claude-sonnet-4' ,
maxConcurrent : 10 ,
} ) ;
const request = {
messages : [ { role : 'user' , content : 'Summarise this document...' } ] ,
max_tokens : 1024 ,
} ;
const result = await bulkhead . run ( request , async ( ) => {
return callYourLLMProvider ( request ) ;
} ) ;
What's New in v3.15 — Protected admission-class floors
v3.15 lets each statically configured admission class reserve a strict local
concurrency floor and, when token admission is enabled, a strict in-flight
token floor. Capacity above all floors forms a shared remainder that any class
may borrow, subject to its hard class ceilings and the enclosing global limits.
const bulkhead = createLLMBulkhead ( {
model : "gpt-4o" ,
maxConcurrent : 12 ,
tokenBudget : { budget : 24_000 } ,
admissionClasses : {
defaultClass : "standard" ,
classes : {
premium : {
protectedConcurrent : 6 ,
maxConcurrent : 10 ,
protectedInFlightTokens : 12_000 ,
maxInFlightTokens : 20_000 ,
} ,
standard : {
protectedConcurrent : 2 ,
maxConcurrent : 6 ,
protectedInFlightTokens : 4_000 ,
maxInFlightTokens : 10_000 ,
} ,
} ,
} ,
} ) ;
In this example, eight concurrency slots and 16,000 tokens are protected. The
remaining four slots and 8,000 tokens are shared. Premium may exceed its floor
by borrowing from that shared remainder, but it cannot consume the two slots
or 4,000 tokens protected for standard.
Floors are strict reservations, not idle-capacity lending. Unused protected
capacity remains available only to its owning class. Cross-class lending of an
idle floor requires a gateway or control plane to install a higher-revision
snapshot. Increasing a floor never revokes active work: new borrowing pauses
until completions restore the guarantee by attrition.
stats().admissionClasses now reports the shared remainder plus each class's
protected and borrowed usage. Rejections caused by preserving another class's
floor use constraint: "admission_class_protection" and include the shared
capacity snapshot. Omitting both floor fields preserves v3.14 admission
behavior.
What's New in v3.14 — Bounded admission classes
v3.14 adds a bounded hierarchical admission layer inside one bulkhead. A
trusted gateway can map tenant or application identity to a small, static set
of policy classes, then enforce hard concurrency and in-flight token ceilings
for each class while retaining the enclosing global limits.
const bulkhead = createLLMBulkhead ( {
model : "gpt-4o" ,
maxConcurrent : 12 ,
tokenBudget : { budget : 24_000 } ,
admissionClasses : {
defaultClass : "standard" ,
classes : {
premium : { maxConcurrent : 8 , maxInFlightTokens : 18_000 } ,
standard : { maxConcurrent : 4 , maxInFlightTokens : 6_000 } ,
} ,
} ,
} ) ;
await bulkhead . run ( request , callProvider , {
admissionClass : "premium" ,
dedupScope : trustedTenantId ,
} ) ;
The class table is fixed at construction. Unknown IDs throw instead of creating
unbounded state or metric cardinality. Class ceilings are fail-fast and are
checked together with global concurrency, priority-adjusted token budget, and
shutdown state. applyLimits() may update class numbers atomically, but every
snapshot must contain exactly the construction-time class keys.
When in-flight deduplication is enabled, the selected admission class is an
automatic deduplication partition. Identical requests in premium and
standard therefore execute independently rather than sharing the first
class's admission decision or queue position. dedupScope remains necessary
for tenant-level isolation inside a class.
Admission classes are intentionally policy buckets, not raw tenant IDs. The
library does not authenticate callers, assign identities, implement weighted
fairness, reserve borrowable floors, or coordinate capacity across processes.
Those responsibilities remain in the gateway and control plane. When
in-flight deduplication is enabled, admissionClass also does not replace
dedupScope ; include trusted tenant identity in dedupScope when results must
not be shared across tenants.
Lifecycle results, run contexts, tokens, usage reports, capacity details, and
telemetry events now carry admissionClass when configured. Rejection detail
also identifies whether the binding constraint was global or
admission_class , and stats().admissionClasses exposes bounded per-class
accounting.
What's New in v3.13 — Progressive reconciliation
Gateways that receiv

[truncated]
