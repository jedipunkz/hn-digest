---
source: "https://github.com/tenurehq/precisionMemBench"
hn_url: "https://news.ycombinator.com/item?id=48404090"
title: "Why Vector Search fails at LLM memory (and a benchmark to prove it)"
article_title: "GitHub - tenurehq/precisionMemBench: Precision-aware retrieval benchmark for LLM memory systems. · GitHub"
author: "decorner"
captured_at: "2026-06-04T21:36:52Z"
capture_tool: "hn-digest"
hn_id: 48404090
score: 3
comments: 0
posted_at: "2026-06-04T20:21:31Z"
tags:
  - hacker-news
  - translated
---

# Why Vector Search fails at LLM memory (and a benchmark to prove it)

- HN: [48404090](https://news.ycombinator.com/item?id=48404090)
- Source: [github.com](https://github.com/tenurehq/precisionMemBench)
- Score: 3
- Comments: 0
- Posted: 2026-06-04T20:21:31Z

## Translation

タイトル: ベクトル検索が LLM メモリで失敗する理由 (およびそれを証明するベンチマーク)
記事のタイトル: GitHub - tenurehq/precisionMemBench: LLM メモリ システムの精度を考慮した取得ベンチマーク。 · GitHub
説明: LLM メモリ システムの精度を意識した検索ベンチマーク。 - GitHub - tenurehq/precisionMemBench: LLM メモリ システムの精度を意識した取得ベンチマーク。

記事本文:
GitHub - tenurehq/precisionMemBench: LLM メモリ システムの精度を意識した取得ベンチマーク。 · GitHub
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
在職期間
/
精度Memベンチ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット フィクスチャ フィクスチャ src src テスト結果/ベースライン テスト結果/ベースライン ラッパー ラッパー .gitignore .gitignore ライセンス ライセンス README.md README.md SUBMITTING.md SUBMITTING.md package-lock.json package-lock.json package.json package.json Providers.config.json Provides.config.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
PrecisionMemBench は、LLM メモリ システムの多次元検索ベンチマークです。これは、シングルターンの回答品質ベンチマークでは検出できない 4 つの直交特性を測定します。
検索精度 - 2 つのドメイン スコープ、スーパーセッション チェーン、セカンダリ ユーザー フィクスチャにわたる 35 の信念の固定シード コーパスに対して、正しい信念表面とその信念のみを抽出します。
ノイズ分離 - トピックから外れたドリフト ターン中に導入された信念が、10 ターン セッション全体にわたるその後の無関係なターンでの検索を汚染しますか
セッション ターン レイテンシ - シングル ターン ベースラインと比較して、セッション負荷の下で取得レイテンシは低下しますか
信念の可変性 - セッション中に更新された信念は、エイリアス強化フライホイールを介して同じセッション内で即座に表面化されます。
これらのプロパティは独立しています。システムは精度はそのままでも、ドリフトで失敗する可能性があります。システムはクリーンなシングル ターン レイテンシを持ち、セッション負荷がかかると 4 倍低下する可能性があります。書き込み時突然変異プリミティブを持たないシステムは 4 番目のプロパティでまったくスコアを付けることができません。これはアーキテクチャ上の欠如であり、パフォーマンスの違いではありません。
すべてのケースでは、メモリ システムが何を返さなければならないかだけでなく、何を返さないかも指定されます。ノイズは重大な障害であり、目に見えない推論コストではありません。
89 件のケースをカバー: エイリアスの解決、スコープの曖昧さの解消、スーパーセッション チェーンの除外、ファジー マッチング、クロス

- ユーザーの分離、予算の削除、ランキングの安定性、マルチターントピックドリフト下でのセッションレベルのノイズ分離
論文: arXiv — データセット: HuggingFace — リーダーボード: HuggingFace Spaces
プロバイダー
アクティブパス
総パス数
平均精度
平均再現率
取得 p50 (ms)
摂取量の合計 (秒)
在職期間
43/43
77/77
1.00
1.00
9.77
1.00
スーパーメモリ
17/17
44/77
0.43
0.55
819.48
0.00
グブレイン
5/5
34/77
0.14
0.17
543.84
28.60
エージェントメモリ
0/0
77/7
0.17
0.97
82.28
1.10
あなたの思い出
0/0
77/21
0.17
0.88
313.39
16.40
アトミックメモリー
0/0
1977 年 9 月
0.15
0.95
71.01
658.90
ゼップ
0/0
1977 年 9 月
0.09
0.95
124.36
897.00
ベクトル
0/0
11/77
0.09
1.00
71.87
—
後知恵
0/0
1977 年 9 月
0.06
1.00
589.86
173.30
メモリ0
0/0
1977 年 9 月
0.06
0.99
64.94
111.30
ああ、思い出した
0/0
1977 年 9 月
0.06
0.99
13.80
178.80
アクティブ パスは、メモリ システム自体が正しく取得されたかどうかを回答する唯一の列です。システムは、すべてを返すか何も返さないことによってアクティブ パスを蓄積することはできません。
再現率 1.0 は精度を意味しません。すべての比較システムは、多くの不正確な信念とともに正しい信念を返し、結果として再現時に完璧なスコアを獲得します。平均精度が 0.05 ～ 0.09 であるということは、およそ 10 ～ 18 個の無関係な信念がそれぞれの正しい信念とともに返されることを意味します。
合計パス数では、この内訳を正しく解釈する必要があります。すべての件数は、セッション以外の77件を超えています。
アクティブな取得パス - ケースには retrievalPrecision アサーションが含まれており、それが満たされます。これは、検証済みの取得機能を示す唯一のパス タイプです。
構造パス - このケースは、精度のアサーションなしでスコープの分離、スーパーセッションの除外、または型ルーティングをアサートし、構造プロパティが保持されます。
自明の空のパス - 予想される関連信念層は、ケースの設計により空です (空のクエリ、 maxBeliefs: 0 、予算は正確な固定カウントに設定されています)。あらゆるシステム

空のセットを返す場合は構築によって渡されます。
モデル
精度
リコール
パス
平均 (ミリ秒)
p95 (ミリ秒)
nomic-embed-text (768)
0.09
1.0
11/77
43.36
85.21
mxbai-embed-large (1024)
0.09
1.0
11/77
96.48
257.24
qwen3-8b (4096)
0.09
1.0
11/77
1130.95
2604.84
すべての構成の 11 パスはすべて、構造的であるか、単純に空です。アクティブな取得パスは 3 つのモデルすべてで 0 です。
セッション評価 - マルチターンドリフト時のノイズ分離
12 のセッション ケースは、トピックから外れたドリフト ターン中に導入された信念がその後の無関係なターンでの取得を汚染するかどうか、セッション負荷の下でレイテンシが低下するかどうか、セッション中に導入された信念がエイリアス エンリッチメント フライホイールを介して同じセッション ウィンドウ内に表面化するかどうかという 3 つの直交する特性をテストします。
ドリフト スコアは、ドリフト ターン トピックに由来する、固定されていない信念を取得した割合です。 0 は完全な分離です。
gbrain は、これらのセッション ケースに対して結果を返しませんでした。構造によりドリフト スコア 0.0 が記録されます。信念は返されなかったため、ドリフトトピックから生じた信念はありませんでした。正しい信念も表面化することができず、これは真の隔離パスではなく、結果的に空虚な失敗となった。
結果テーブルを解釈するには、3 つのパス タイプを理解する必要があります。
アクティブな取得パス — ケースには retrievalPrecision アサーションが含まれており、それが満たされます。これは、検証済みの取得機能を示す唯一のパス タイプです。システムは、すべてを返すか何も返さないことによってアクティブ パスを蓄積することはできません。
構造パス — このケースは、精度のアサーションなしでスコープの分離、スーパーセッションの除外、または型ルーティングをアサートし、構造プロパティが保持されます。
自明の空のパス — 予想される関連信念層はケースの設計により空です (空のクエリ、 maxBeliefs: 0 、予算は正確に固定された予算に設定されています)

nt）。空のセットを返すシステムは構築をパスします。このような場合、retrievalPrecision は null になります。
この内訳がなければ、集計パス数は検証済みの取得と構造パスまたは空セット パスを区別できません。
89 件のケースは次のカテゴリーに当てはまります。セッション ケースはコーパスを動的に拡張します。セッション中にビリーフが作成され、エイリアス セットが更新されます。これにより、取得がスナップショットではなくライブ ストアの状態を反映していることが検証されます。
エイリアスの解決 — さまざまな表面形式 (短い形式、自然言語、複数の単語) が正しい信念に解決されるかどうか。
スコープの曖昧さの解消 — 異なるドメイン スコープ間でエイリアスを共有する信念をスコープだけで正しく区別できるかどうか。
置き換えチェーンの除外 — 置き換えられた信念がマルチホップ チェーンの深部で除外されるかどうか。置き換えられる用語と置き換えられる用語の両方に一致するクエリでは、置き換えられる信念のいずれも表面化する必要があります。アクティブな最終的な信念は、ピン留めされた事実層を介して表面化します。
ファジー マッチングとプレフィックス ガード - 編集距離のみで許可されるプレフィックスの不一致をブロックしながら、検索レイヤーが転置とニアミス用語を正しく処理するかどうか。合格動作と不合格動作の両方が、意図的な設計プロパティとして文書化されます。
カウンターシグナルの取得 - 拒否された用語または置き換えられた用語を参照するクエリが、カウンターシグナルのエイリアスを介してアクティブな置換信念を表面化するかどうか。どちらの場合も、アクティブな取得精度アサーションを保持します。
関係の拡張 — 関係タイプの信念が、拡張中に適用される参加者タイプのルーティングとスコープ フィルターを使用して、ワンホップ結合を介して参加者を正しく表面化および拡張するかどうか。
セッションレベルのノイズ分離 — トピックから外れたドリフトターン中に導入された信念が、その後の無関係なターンでの検索を汚染するかどうか

。主なケースは、8 ターンにわたってトピックが移動し、その後に暗黙的なリターンが続く 10 ターンのセッションです。ターンごとのアサーションは、再突入時の分離を検証します。
予算の立ち退きと容量 — 取得レイヤーがスロット制約を正しく処理するかどうか (正常な空のリターン、単一スロットの優先順位、予算上限での高鉄筋のフラッディングに対する耐性など)。
設計境界ケース — 合格動作と不合格動作の両方が意図的な設計プロパティとして文書化されるケース。
タイプ ルーティングと未解決の質問 - 未解決の質問が、アクティブなスコープに対して固定された未解決の質問のみを返し、テキスト検索では決して返されない別のパスによって取得されるかどうか。
ランキングの安定性 — スコア主導の並べ替えアーティファクトを発生させずに、同等のクエリ間で取得結果が安定しているかどうか。
クロスユーザーの分離 — セマンティックな近接性に関係なく、2 番目のユーザーに属する信念がプライマリ ユーザーの検索から構造的に除外されるかどうか。
コールド スタート動作 - シードされた信念がゼロの新規ユーザーがエラーなしで完全に空のコンテキストを返すかどうか。
ペルソナ プレリュード コンテンツ — 蓄積された信念状態から生成されたペルソナ プレリュードが正しく挿入され、生きた信念ストアを反映しているかどうか。
ケースごとに 4 つの指標が記録されます。
取得精度と再現率 — 関連する信念層がアクティブなアサーションを保持している場合に、その層に対して計算されます。このメトリックが構造的に適用できない場合は、null が記録され、集計計算から除外されます。
ピン留めされたカバレッジ — pinnedFacts 層がアサートされるケースについて記録されます。
質問の精度と再現率 - openQuestions 層がアサートされた場合に記録されます。
パスには、アサートされたすべての層が同時に満たされることが必要です。 retrievalPrecision: 1.0 の場合

また、満たされていない pinnedCoverage アサーションが失敗します。
ドリフト スコアは、セッション ケースに対してレポートされます。ドリフト ターン トピックに由来する、取得された固定されていない信念の割合です。 0 は完全な分離です。
すべての参照システムの実行前レポートは、 test-results/baseline/ にコミットされます。
テスト結果/ベースライン/
取得レポート.json
取得レポートベクター.json
取得レポート-mem0.json
取得レポート-zep.json
取得レポート-後知恵.json
各レポートには、 pass 、 Failure 、 retrievalPrecision 、 retrievalRecall 、および retrievalLatencyMs を含むケースごとの結果に加え、集計された p50 / p95 レイテンシとトップレベルの平均適合率/再現率が含まれます。
独自のプロバイダーに対して実行する場合は、test-result/ の出力をこれらのファイルと直接比較します。
Docker (ベクターベースラインおよびプロバイダースタック用)
ベクターベースラインのみの Ollama インスタンス
npmインストール
2. 比較プロバイダーに対して実行する
プロバイダーのスタックを開始して、次のようにします。
MEMORY_PROVIDER=mem0 npx ava retrieval.external.eval.test.ts
MEMORY_PROVIDER=mem0 npx ava session-retrieval.external.eval.test.ts
レポートは test-results/ に格納されます。有効な値: mem0 、 zep 、 hindsight
ベクター eval は、独自の MongoDB Atlas Local コンテナーを管理します。 Docker が実行されている必要がありますが、手動で何も設定する必要はありません。
＃ 1つ

[切り捨てられた]

## Original Extract

Precision-aware retrieval benchmark for LLM memory systems. - GitHub - tenurehq/precisionMemBench: Precision-aware retrieval benchmark for LLM memory systems.

GitHub - tenurehq/precisionMemBench: Precision-aware retrieval benchmark for LLM memory systems. · GitHub
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
tenurehq
/
precisionMemBench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits fixtures fixtures src src test-results/ baseline test-results/ baseline wrappers wrappers .gitignore .gitignore LICENSE LICENSE README.md README.md SUBMITTING.md SUBMITTING.md package-lock.json package-lock.json package.json package.json providers.config.json providers.config.json tsconfig.json tsconfig.json View all files Repository files navigation
PrecisionMemBench is a multi-dimensional retrieval benchmark for LLM memory systems. It measures four orthogonal properties that single-turn answer-quality benchmarks cannot detect:
Retrieval precision - does the right belief surface, and only that belief, against a fixed seed corpus of 35 beliefs spanning two domain scopes, a supersession chain, and a secondary-user fixture
Noise isolation - do beliefs introduced during off-topic drift turns contaminate retrieval on subsequent unrelated turns across a 10-turn session
Session-turn latency - does retrieval latency degrade under session load relative to single-turn baselines
Belief mutability - do beliefs updated mid-session surface immediately within the same session via the alias enrichment flywheel
These properties are independent. A system can pass on precision and fail on drift. A system can have clean single-turn latency and degrade 4x under session load. A system with no write-time mutation primitive cannot be scored on the fourth property at all, it is an architectural absence, not a performance difference.
Every case specifies not just what the memory system must return, but what it must not. Noise is a hard failure, not an invisible inference cost.
89 cases covering: alias resolution · scope disambiguation · supersession chain exclusion · fuzzy matching · cross-user isolation · budget eviction · ranking stability · session-level noise isolation under multi-turn topic drift
Paper: arXiv — Dataset: HuggingFace — Leaderboard: HuggingFace Spaces
Provider
Active passes
Total passes
Mean precision
Mean recall
Retrieval p50 (ms)
Ingestion total (s)
tenure
43/43
77/77
1.00
1.00
9.77
1.00
supermemory
17/17
44/77
0.43
0.55
819.48
0.00
gbrain
5/5
34/77
0.14
0.17
543.84
28.60
agentmemory
0/0
7/77
0.17
0.97
82.28
1.10
yourmemory
0/0
21/77
0.17
0.88
313.39
16.40
atomicmemory
0/0
9/77
0.15
0.95
71.01
658.90
zep
0/0
9/77
0.09
0.95
124.36
897.00
vector
0/0
11/77
0.09
1.00
71.87
—
hindsight
0/0
9/77
0.06
1.00
589.86
173.30
mem0
0/0
9/77
0.06
0.99
64.94
111.30
a-mem
0/0
9/77
0.06
0.99
13.80
178.80
Active passes are the only column that answers whether the memory system itself retrieved correctly. A system cannot accumulate active passes by returning everything or nothing.
Recall of 1.0 does not imply precision. Every comparison system returns the correct belief alongside many incorrect ones and scores perfectly on recall as a result. Mean precision of 0.05 to 0.09 means roughly 10 to 18 irrelevant beliefs are returned alongside each correct one.
Total pass counts require this breakdown to be interpreted correctly. All counts are over the 77 non-session cases.
Active retrieval pass - the case carries a retrievalPrecision assertion and it is satisfied. This is the only pass type that demonstrates verified retrieval capability.
Structural pass - the case asserts scope isolation, supersession exclusion, or type routing without a precision assertion, and the structural property holds.
Trivially empty pass - the expected relevantBeliefs tier is empty by case design (empty query, maxBeliefs: 0 , budget set to exact pinned count). Any system returning an empty set passes by construction.
Model
Precision
Recall
Passes
Mean (ms)
p95 (ms)
nomic-embed-text (768)
0.09
1.0
11/77
43.36
85.21
mxbai-embed-large (1024)
0.09
1.0
11/77
96.48
257.24
qwen3-8b (4096)
0.09
1.0
11/77
1130.95
2604.84
All 11 passes in every configuration are structural or trivially empty. Active retrieval passes are 0 across all three models.
Session eval — noise isolation under multi-turn drift
The 12 session cases test three orthogonal properties: whether beliefs introduced during off-topic drift turns contaminate retrieval on subsequent unrelated turns, whether latency degrades under session load, and whether beliefs introduced mid-session surface within the same session window via the alias enrichment flywheel.
The drift score is the fraction of retrieved non-pinned beliefs originating from drift-turn topics; 0 is perfect isolation.
‡ gbrain returned no results for these session cases. A drift score of 0.0 is recorded by construction; no beliefs were returned, so none could originate from drift topics. The correct belief also failed to surface, making this an empty-result failure rather than a genuine isolation pass.
Understanding the three pass types is required to interpret any results table.
Active retrieval pass — the case carries a retrievalPrecision assertion and it is satisfied. This is the only pass type that demonstrates verified retrieval capability. A system cannot accumulate active passes by returning everything or nothing.
Structural pass — the case asserts scope isolation, supersession exclusion, or type routing without a precision assertion, and the structural property holds.
Trivially empty pass — the expected relevantBeliefs tier is empty by case design (empty query, maxBeliefs: 0 , budget set to exact pinned count). Any system returning an empty set passes by construction. retrievalPrecision is null for these cases.
Without this breakdown, aggregate pass counts do not distinguish verified retrieval from structural or empty-set passes.
The 89 cases cover the following categories. Session cases extend the corpus dynamically — beliefs are created and alias sets updated mid-session — validating that retrieval reflects the live store state rather than a snapshot.
Alias resolution — whether variant surface forms (short-form, natural-language, multi-word) resolve to the correct belief.
Scope disambiguation — whether scope alone correctly discriminates between beliefs sharing an alias across different domain scopes.
Supersession chain exclusion — whether superseded beliefs are excluded at depth in a multi-hop chain. A query matching both a superseded and a superseding term must surface neither superseded belief; the active terminal belief surfaces via the pinned facts tier.
Fuzzy matching and prefix guards — whether the retrieval layer correctly handles transpositions and near-miss terms while blocking prefix mismatches that edit distance alone would permit. Both pass and fail behaviors are documented as intentional design properties.
Counter-signal retrieval — whether a query referencing a rejected or superseded term surfaces the active replacement belief via a counter-signal alias. Both cases carry an active retrieval precision assertion.
Relation expansion — whether relation-type beliefs correctly surface and expand their participants via a one-hop join, with participant type routing and scope filters applied during expansion.
Session-level noise isolation — whether beliefs introduced during off-topic drift turns contaminate retrieval on subsequent unrelated turns. The primary case is a 10-turn session with topic drift across 8 turns followed by an implicit return; per-turn assertions verify isolation at re-entry.
Budget eviction and capacity — whether the retrieval layer handles slot constraints correctly, including graceful empty returns, single-slot priority, and resistance to high-reinforcement flooding at the budget ceiling.
Design boundary cases — cases where both pass and fail behaviors are documented as intentional design properties.
Type routing and open questions — whether open questions are retrieved by a separate path that returns only pinned open questions for the active scope and are never returned by text search.
Ranking stability — whether retrieval results remain stable across equivalent queries without score-driven reordering artifacts.
Cross-user isolation — whether beliefs belonging to a second user are structurally excluded from a primary user's retrieval regardless of semantic proximity.
Cold start behavior — whether a new user with zero seeded beliefs returns a fully empty context without error.
Persona prelude content — whether the persona prelude generated from the accumulated belief state is injected correctly and reflects the live belief store.
Four metrics are recorded per case:
Retrieval precision and recall — computed over the relevantBeliefs tier on cases where that tier carries an active assertion. Cases where this metric is structurally inapplicable record null and are excluded from aggregate computation.
Pinned coverage — recorded on cases where the pinnedFacts tier is asserted.
Question precision and recall — recorded on cases where the openQuestions tier is asserted.
A pass requires all asserted tiers to be simultaneously satisfied. A case with retrievalPrecision: 1.0 that also carries an unmet pinnedCoverage assertion fails.
Drift score is reported for session cases: the fraction of retrieved non-pinned beliefs originating from drift-turn topics. 0 is perfect isolation.
Pre-run reports for all reference systems are committed at test-results/baseline/ :
test-results/baseline/
retrieval-report.json
retrieval-report-vector.json
retrieval-report-mem0.json
retrieval-report-zep.json
retrieval-report-hindsight.json
Each report contains per-case results including passed , failures , retrievalPrecision , retrievalRecall , and retrievalLatencyMs , plus aggregate p50 / p95 latency and mean precision/recall at the top level.
When you run against your own provider, compare your output in test-results/ directly against these files.
Docker (for the vector baseline and provider stacks)
An Ollama instance for the vector baseline only
npm install
2. Run against a comparison provider
Start the provider's stack, then:
MEMORY_PROVIDER=mem0 npx ava retrieval.external.eval.test.ts
MEMORY_PROVIDER=mem0 npx ava session-retrieval.external.eval.test.ts
Reports land in test-results/ . Valid values: mem0 , zep , hindsight
The vector eval manages its own MongoDB Atlas Local container. Docker must be running but you do not set anything up manually.
# One

[truncated]
