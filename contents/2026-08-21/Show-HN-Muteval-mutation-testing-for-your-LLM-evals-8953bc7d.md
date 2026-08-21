---
source: "https://github.com/AshwinUgale/muteval"
hn_url: "https://news.ycombinator.com/item?id=49388918"
title: "Show HN: Muteval – mutation testing for your LLM evals"
article_title: "GitHub - AshwinUgale/muteval: Mutation testing for your LLM evals — degrade the system under test and check whether your evals would actually catch a regression. · GitHub"
image: "https://opengraph.githubassets.com/f5fa46ff37db700bba418b7aa54b24ca3fbc0bcc059b5616b76002765e5cb243/AshwinUgale/muteval"
author: "Ashwin1121"
captured_at: "2026-08-21T15:24:30Z"
capture_tool: "hn-digest"
hn_id: 49388918
score: 1
comments: 0
posted_at: "2026-08-21T14:42:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Muteval – mutation testing for your LLM evals

- HN: [49388918](https://news.ycombinator.com/item?id=49388918)
- Source: [github.com](https://github.com/AshwinUgale/muteval)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T14:42:57Z

## Translation

タイトル: HN を表示: Muteval – LLM 評価の突然変異テスト
記事のタイトル: GitHub - AshwinUgale/muteval: LLM eval のミューテーション テスト — テスト対象のシステムをデグレードし、eval が実際に回帰を検出するかどうかを確認します。 · GitHub
説明: LLM eval の突然変異テスト — テスト対象のシステムをデグレードし、eval が実際に回帰を検出するかどうかを確認します。 - アシュウィン・ウガレ/ミューテヴァル

記事本文:
GitHub - AshwinUgale/muteval: LLM eval の突然変異テスト — テスト対象のシステムをデグレードし、eval が実際に回帰をキャッチするかどうかを確認します。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アシュウィンウガレ
/
無言
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
118 コミット 118 コミット フォルダーとファイル
.github .github ブログ ブログ ドキュメント ドキュメントの例 例 js js src/ muteval src/ muteval テスト テスト va

蓋の検証 .editorconfig .editorconfig .git-blame-ignore-revs .git-blame-ignore-revs .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md FINDINGS.md FINDINGS.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md action.yml action.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM 評価の突然変異テスト — 実際に回帰を検出するかどうかを確認します。
あなたの eval は合格しています。それは彼らが働くという意味ではありません。
muteval は、すべての eval スイートが静かに回避する質問に答えます: 私の eval はするでしょうか
システムが静かに悪化した場合、実際に失敗しますか?意図的に劣化させます
テスト対象のもの、劣化したそれぞれに対して既存の評価スイートを再実行します。
バージョン (「変異体」) を調べ、変異スコア (その割合) を報告します。
eval が捕らえた注入された回帰。彼らが恋しいのは生存者たちだ――
トリアージすべき候補カバレッジのギャップ: muteval がそれらを表面化し、どれを選択するかはあなたが決定します
実際には重要です ( docs/LIMITATIONS.md を参照)。
mutmut / Stryker ですが eval 用です。
突然変異スコア: 33% [████████░░░░░░░░░░░░░░░░] (2/6 突然変異体が死亡、95% CI 10-70%)
2 SURVIVED (出力は変更されましたが、評価者は気付かなかった - 実際のカバレッジ ギャップ、1 つは重大度が高い):
#1 [高] 生き残った [delete_sentences]
削除された文: 「答えが文脈にない場合は、わからないと言ってください。」
修正: checks.grounded("context") を追加 ← muteval はそれをキャッチする eval を提案します
#2 [MED] 生き残った [weaken_modals]
弱められた「のみ」 -> 「できれば」 (近: 提供されたコンテキストのみを使用して回答)
インストール
pip インストール Mut

eval # 純粋な Python、必要な依存関係はゼロ
pip install -U muteval # すでにインストールされていますか?アップグレード (ベアインストールは何もしません)
コアは重い LLM SDK を引き込みません。オプションの追加機能 (使用する場合のみ):
テスト対象システムの任意のプロバイダー: OpenAI 互換の任意のプロバイダーを指します。
--base-url (または OPENAI_BASE_URL ) を使用したエンドポイント — Groq、Gemini-compat、GitHub
モデル、Ollama、ローカルサーバー。いつでも利用できるものを見つけてください
muteval リスト (演算子、チェック、プローブ)。
60 秒のクイックスタート (API キーなし)
muteval init --template rag # 構成を足場にします (または --template Basic)
muteval check --config muteval_config.py # 最初に配線とベースラインを検証します
muteval run --config muteval_config.py
muteval init は、明確に指定した 4 つの要素を含む実行可能な構成を書き込みます
マークされています。 muteval チェックは医者です。パイプライン、評価、および評価を検証します。
フル実行前のベースラインを確認し、どの層が壊れているかを正確に示します。それから
run を実行すると、突然変異スコアと生存者のランク付けされたリストが得られます。
ギャップを埋めるために e を提案しました。
ゼロ構成をご希望ですか?プロンプト + ケースで muteval をポイントし、モデルを呼び出します。
エクスポート OPENAI_API_KEY=sk-...
muteval run --prompt-file system.txt --cases case.jsonl --model gpt-4o-mini \
--judge " 答えは提供されたコンテキストに基づいています " --fail-under 75
すでにプロンプトフスイートをお持ちですか?ミュートバルをまっすぐに向けます - ミュートバルはありません
config ファイルを使用すると、プロンプト + テスト + アサーションが再利用されます。
muteval run --promptfoo promptfooconfig.yaml # プレビューに --dry-run を追加します
Providers: ブロックからモデルを読み取るので、プロンプトを変更します。
スイートが実際に使用するモデル。 CI に入れたいですか? GitHub アクションがあります —
「CI での実行」を参照してください。
API キーを使用せずに、最初にカバレッジ ギャップを検出する様子を確認したいですか?オフラインで実行する
デモ — サポートボットのプロンプトを劣化させ、次のことを検出します。

プロンプトフスイートのルール
アサートするのを忘れました (~1 秒、モック モデル):
muteval run --config example/promptfoo_offline/muteval_config.py --no-color
すでに独自のパイプラインをお持ちですか?それをテスト対象のシステム (関数または関数) として使用します。
デプロイされたエンドポイント — run() ラッパーなし:
# fn(prompt, case) という独自の関数 -> str
muteval run --target mypkg.app:answer --prompt-file system.txt --cases case.jsonl --check contains:8080
# デプロイされたサービス: {prompt, case} JSON を POST し、テキスト出力を読み取ります
muteval run --endpoint https://my-app/answer --prompt-file system.txt --cases case.jsonl --judge " コンテキストに基づいた "
再実行は安価です: --cache run.sqlite は実行出力 + 評価結果をメモ化します。
そのため、同一の再実行ではモデル/ジャッジコールがゼロになります（ノイズの多いスイートの場合はスキップされます）。
--runs-per-mutant > 1 の場合):
muteval run --config muteval_config.py --cache .muteval-cache.sqlite
API に依存しているため遅いですか?ミュータントを並行して評価します (結果は同一です)
シリアル実行へ — 順序は保持されます):
muteval run --config muteval_config.py --concurrency 8 --cache .muteval-cache.sqlite
出費が心配ですか?モデルにキャップを付けてコールを判断します。 muteval が失敗して閉じられる (出口 2)
使いすぎる前に (キャッシュ ヒットとスキップされたジャッジはカウントされません):
muteval run --config muteval_config.py --max-calls 500
再実行せずに生存者を優先順位付けします (最後の実行は次のファイルに保存されます)
.muteval/last_run.json ):
muteval の結果 ID でランク付けされたサバイバー (上位から順) の数
muteval show 1 # 1 人の生存者: オペレーター、提案された修正、ベースライン→ミュータントの差分
muteval report --html Coverage.html # 共有可能なスタンドアロン レポート
ミューテーションの範囲を超えて、ミューテバル プローブは他のものとともに評価スイートを監査します。
レンズ。耐荷重性のものは実際の一般的な欠陥を捕らえます: 信頼性を判断します
(LLM の審査員は同一の再放送をオンにしますか?) と差別 (できない

彼は
eval は、良い/悪い見本が与えられた場合に、良い出力と悪い出力を区別しますか?)。残りは衛生面です
チェック - 統計的妥当性と冗長性 - さらに、ラベルがある場合のみ、
人間の合意 (ミューテヴァル ラベルを介したコーエンのκ)、1 つの真の妥当性チェック。
通知表、総合得点なし。 (別の裁判官バイアス委員会 —
位置 / 冗長性 / 自己優先 — のライブラリ関数として利用可能です。
ペアのA/Bジャッジ。デフォルトのカードには含まれていません。
ペアワイズジャッジハーネス。)
muteval プローブ --config muteval_config.py --htmlquality.html
なぜこれが存在するのか
回帰ツール (promptfoo、deepeval、OpenAI Evals、LangSmith) キャッチ
システムの回帰。 eval が十分かどうかは誰も教えてくれません
まずはそれらの回帰を捉えるためです。そのメタ層がギャップです
muteval fills — ミューテーション テストは、「これが私のテスト スイートです」に対する確立された答えです
何かいいの？」ソフトウェアエンジニアリングの学位を取得し、LLM 評価スイートに導入されました。
小さな構成でシステム + eval を記述し、次に muteval を記述します。
ベースライン — スイートが元のシステムに合格することを確認します。もしそうなら
そうしないと、muteval はスコアを拒否します (赤いベースラインがすべての数字を作ります)
無意味）誤解を招く100％を渡すよりも。
Mutate — プロンプト / 取得したコンテキスト / を劣化させることでミュータントを生成します。
ツール出力/モデル (18 オペレーター)。
グレード — 各ミュータントに対してスイートを再実行します。殺された = あなたの評価者
捕まえました（良い）。生き残った = 彼らはそれを逃しました (ギャップ)。
スコア — 死亡/評価、95% 信頼区間、重大度
ランキング、ニアミスマージン、生存者ごとの修正案。
from muteval import MutEvalConfig 、チェック
config = MutEvalConfig (
プロンプト = SYSTEM_PROMPT , # テスト対象のもの
case = [{ "input" : "..." , "order_id" : "A123" }],
run = my_run_fn , # LLM/アプリを呼び出します -> 出力テキスト
evals = [ #あなたの

既存のチェック、ミュートヴァルによって評価
小切手。 contains_case ( "order_id" )、
小切手。接地 ( "context" )、# LLM 判定プリセット (任意の OpenAI 互換エンドポイント)
]、
）
設計上信頼できる
信頼できない補償範囲の数字は、何もないよりも悪いです。 muteval が失敗して閉じられます:
赤またはエラーのベースライン → スコアなし (ステータスはbaseline_failed / errored)、
CLI はゼロ以外で終了し、バッジは表示されません。
予算を超える部分的な突然変異エラー → Partial_errors 、予算を超えるスコアではない
縮んだ分母。 --max-error-rate / --allow-mutant-errors を受け入れます。
非決定論 → runs_per_mutant に対する完全多数決の評決、ウィルソン
スコアの信頼区間、不安定な変異フラグ付け。
表面的な変更 → 出力差分により、実際のカバレッジギャップを分離
「観察上変化のない」突然変異体。
制御された CI を適用した実験では、突然変異スコアは単調に増加します
eval-suite カバレッジあり — eval なしの場合は 0% → 完全なカバレッジの場合は 100% —
4 つのドメイン (サポート ボット、コード レビュー、RAG、人事ポリシー) にわたっています。参照
FINDINGS.md 、および docs/LIMITATIONS.md
番号を信用しないとき。
変更できるもの (18 オペレーター)
プロンプト:weak_modals、flip_negation、drop_instruction_lines、
delete_sentences 、 truncate_prompt 、drop_few_shot_example 、remove_emphasis 。
取得されたコンテキスト (RAG):drop_context_doc 、clear_context 、
破損した_コンテキスト_ドキュメント、スワップ_コンテキスト_ドキュメント、シャッフル_コンテキスト、
Duplicate_context_doc 、 truncate_context_doc 。
モデル: downgrade_model 。ツール (エージェント):drop_tool_output 、
破損_ツール_出力、スワップ_ツール_出力。
System(prompt=..., context=[...], tools=[...], model=...) を渡して作成します。
RAG およびエージェント スイートに対して変更可能なコンテキスト / ツール / モデル。自分のオペレーターを連れてくる
register_operator を使用し、プロンプトのどの部分が変更されるかをスコープします。
[[mutate]]…[[/mutate]] マーカーまたは --scope-include/-exclude 。
ムテフ

--config muteval_config.py --fail-under 75 --badge バッジ.json を実行します。
muteval run --config muteval_config.py --junit junit.xml
カバレッジが 75% を下回るとゼロ以外で終了するため、評価を弱める PR は失敗します
ビルド。 -- ガードされていない高重大度ギャップに対するフェイルオン重大度高ゲート。
example/ci/github-actions.yml をコピーしてすべての PR で実行し、
Shields.io の評価カバレッジ バッジ。
すでにスイートをお持ちですか?メトリクスを書き換えるのではなく、再利用します。
deepeval から。メトリクスのインポート FaithhoodMetric
ミュートヴァルから。アダプター。 deepeval インポート metrics_to_evals
evals = metrics_to_evals ([ FaithhoodMetric ()], input_key = "質問" ,
retrieval_context_key = "コンテキスト" )
deepeval 、 RAGAS 、および Promptfoo 用のアダプター ( pip install "muteval[deepeval|ragas|promptfoo]" )。または、組み込みのフレームワークフリーの機能を使用します
チェック — llm_judge / grounded を含む OpenAI 互換のチェック
エンドポイント (OpenAI、Groq、Gemini、GitHub モデル、Ollama…) (base_url= 経由)、使用
標準ライブラリのみ。
実際のシステムでミュートバルを指定するのは、プラグアンドプレイではなく、約 1 時間の統合です。
docs/ADOPTION.md には正直なチェックリストがあります。
「どこを壊すか→何を変えるか」テーブル、審査員選択ガイダンス、および 4 つ
あなたが供給する部品。 muteval チェックから始めて、最初に緑色のベースラインを修正します。
出荷時: プロンプト/コンテキスト/ツール/モデルの突然変異 · deepeval/RAG

[切り捨てられた]

## Original Extract

Mutation testing for your LLM evals — degrade the system under test and check whether your evals would actually catch a regression. - AshwinUgale/muteval

GitHub - AshwinUgale/muteval: Mutation testing for your LLM evals — degrade the system under test and check whether your evals would actually catch a regression. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
AshwinUgale
/
muteval
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
118 Commits 118 Commits Folders and files
.github .github blog blog docs docs examples examples js js src/ muteval src/ muteval tests tests validation validation .editorconfig .editorconfig .git-blame-ignore-revs .git-blame-ignore-revs .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md FINDINGS.md FINDINGS.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md action.yml action.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml View all files Repository files navigation
Mutation testing for your LLM evals — find out if they'd actually catch a regression.
Your evals are passing. That doesn't mean they work.
muteval answers the question every eval suite quietly dodges: would my evals
actually fail if my system silently got worse? It deliberately degrades the
thing under test, reruns your existing eval suite against each degraded
version (a "mutant"), and reports a mutation score — the percentage of
injected regressions your evals caught. The ones they miss are survivors —
candidate coverage gaps to triage: muteval surfaces them, you decide which ones
actually matter (see docs/LIMITATIONS.md ).
It's mutmut / Stryker, but for evals.
Mutation score: 33% [████████░░░░░░░░░░░░░░░░] (2/6 mutants killed, 95% CI 10-70%)
2 SURVIVED (output changed but evals didn't notice — real coverage gaps; 1 HIGH-severity):
#1 [HIGH] SURVIVED [delete_sentences]
deleted sentence: "If the answer is not in the context, say you don't know."
fix: add checks.grounded("context") ← muteval suggests the eval that would catch it
#2 [MED] SURVIVED [weaken_modals]
weakened "ONLY" -> "preferably" (near: answer using ONLY the provided context)
Install
pip install muteval # pure Python, zero required dependencies
pip install -U muteval # already installed? upgrade (bare install is a no-op)
The core drags in no heavy LLM SDKs. Optional extras, only if you use them:
Any provider for the system under test: point it at any OpenAI-compatible
endpoint with --base-url (or OPENAI_BASE_URL ) — Groq, Gemini-compat, GitHub
Models, Ollama, a local server. Discover what's available anytime with
muteval list (operators, checks, probes).
60-second quickstart (no API key)
muteval init --template rag # scaffold a config (or --template basic)
muteval check --config muteval_config.py # validate wiring + baseline first
muteval run --config muteval_config.py
muteval init writes a runnable config with the four things you supply clearly
marked. muteval check is the doctor — it validates your pipeline, evals, and
baseline before a full run and tells you exactly which layer is broken. Then
run gives you a mutation score and a ranked list of survivors, each with a
suggested eval to close the gap.
Prefer zero-config? Point muteval at a prompt + cases and let it call the model:
export OPENAI_API_KEY=sk-...
muteval run --prompt-file system.txt --cases cases.jsonl --model gpt-4o-mini \
--judge " the answer is grounded in the provided context " --fail-under 75
Already have a promptfoo suite? Point muteval straight at it — no muteval
config file, it reuses your prompt + tests + assertions:
muteval run --promptfoo promptfooconfig.yaml # add --dry-run to preview
It reads the model from your providers: block, so it mutates your prompt against
the model your suite actually uses. Want it in CI? There's a GitHub Action —
see Run in CI .
Want to see it find a coverage gap first, with no API key ? Run the offline
demo — it degrades a support-bot prompt and finds the rule the promptfoo suite
forgot to assert (~1s, mock model):
muteval run --config examples/promptfoo_offline/muteval_config.py --no-color
Already have your own pipeline? Use it as the system under test — a function or a
deployed endpoint — no run() wrapper:
# your own function, called as fn(prompt, case) -> str
muteval run --target mypkg.app:answer --prompt-file system.txt --cases cases.jsonl --check contains:8080
# a deployed service: POSTs {prompt, case} JSON, reads the text output
muteval run --endpoint https://my-app/answer --prompt-file system.txt --cases cases.jsonl --judge " grounded in context "
Re-running is cheap: --cache runs.sqlite memoizes run outputs + eval outcomes,
so an identical re-run makes zero model/judge calls (skipped for noisy suites
with --runs-per-mutant > 1 ):
muteval run --config muteval_config.py --cache .muteval-cache.sqlite
Slow because it's API-bound? Evaluate mutants in parallel (results are identical
to a serial run — order preserved):
muteval run --config muteval_config.py --concurrency 8 --cache .muteval-cache.sqlite
Worried about spend? Cap the model + judge calls; muteval fails closed (exit 2)
before overspending (cache hits and skipped judges don't count):
muteval run --config muteval_config.py --max-calls 500
Triage the survivors without re-running (the last run is saved to
.muteval/last_run.json ):
muteval results # ranked survivors (HIGH first) with ids
muteval show 1 # one survivor: operator, suggested fix, baseline→mutant diff
muteval report --html coverage.html # a shareable standalone report
Beyond mutation coverage, muteval probe audits the eval suite along other
lenses. The load-bearing ones catch real, common defects: judge reliability
(does your LLM judge flip on identical re-runs?) and discrimination (can the
eval tell good outputs from bad, given good/bad exemplars?). The rest are hygiene
checks — statistical adequacy and redundancy — plus, only if you have labels,
human agreement (Cohen's κ via muteval label ), the one true validity check.
A report card, no composite score. (A separate judge-bias panel —
position / verbosity / self-preference — is available as a library function for
pairwise A/B judges ; it isn't part of the default card because it needs a
pairwise-judge harness.)
muteval probe --config muteval_config.py --html quality.html
Why this exists
Regression tools (promptfoo, deepeval, OpenAI Evals, LangSmith) catch
regressions in your system . None tell you whether your evals are good enough
to catch those regressions in the first place. That meta-layer is the gap
muteval fills — mutation testing is the established answer to "is my test suite
any good?" in software engineering, brought to LLM eval suites.
Describe your system + evals in a small config, then muteval:
Baseline — confirms your suite passes on the original system. If it
doesn't, muteval refuses to score (a red baseline makes every number
meaningless) rather than hand you a misleading 100%.
Mutate — generates mutants by degrading the prompt / retrieved context /
tool outputs / model (18 operators).
Grade — reruns your suite against each mutant. Killed = your evals
caught it (good); survived = they missed it (a gap).
Score — killed / evaluated , with a 95% confidence interval, severity
ranking, near-miss margins, and a suggested fix per survivor.
from muteval import MutEvalConfig , checks
config = MutEvalConfig (
prompt = SYSTEM_PROMPT , # the thing under test
cases = [{ "input" : "..." , "order_id" : "A123" }],
run = my_run_fn , # call your LLM/app -> output text
evals = [ # your existing checks, graded by muteval
checks . contains_case ( "order_id" ),
checks . grounded ( "context" ), # LLM-judge preset (any OpenAI-compatible endpoint)
],
)
Trustworthy by design
A coverage number you can't trust is worse than none. muteval fails closed :
Red or errored baseline → no score (status baseline_failed / errored ),
CLI exits non-zero, no badge.
Partial mutant errors above a budget → partial_errors , not a score over a
shrunken denominator. --max-error-rate / --allow-mutant-errors to accept.
Non-determinism → strict-majority verdicts over runs_per_mutant , Wilson
confidence intervals on the score, flaky-mutant flagging.
Cosmetic changes → output-diffing separates real coverage gaps from
"observationally unchanged" mutants.
In a controlled, CI-enforced experiment the mutation score rises monotonically
with eval-suite coverage — 0% with no evals → 100% with complete coverage —
across four domains (support bot, code review, RAG, HR policy). See
FINDINGS.md , and docs/LIMITATIONS.md for
when to distrust the number.
What it can mutate (18 operators)
Prompt: weaken_modals , flip_negation , drop_instruction_lines ,
delete_sentences , truncate_prompt , drop_few_shot_example , remove_emphasis .
Retrieved context (RAG): drop_context_doc , clear_context ,
corrupt_context_doc , swap_context_doc , shuffle_context ,
duplicate_context_doc , truncate_context_doc .
Model: downgrade_model . Tools (agents): drop_tool_output ,
corrupt_tool_output , swap_tool_output .
Pass a System(prompt=..., context=[...], tools=[...], model=...) to make
context / tools / model mutable for RAG and agent suites. Bring your own operator
with register_operator , and scope which parts of the prompt mutate with
[[mutate]]…[[/mutate]] markers or --scope-include/-exclude .
muteval run --config muteval_config.py --fail-under 75 --badge badge.json
muteval run --config muteval_config.py --junit junit.xml
Exits non-zero if coverage drops below 75%, so a PR that weakens your evals fails
the build. --fail-on-severity high gates on any unguarded high-severity gap.
Copy examples/ci/github-actions.yml to run it on every PR and publish a
shields.io eval-coverage badge.
Already have a suite? Reuse its metrics instead of rewriting them:
from deepeval . metrics import FaithfulnessMetric
from muteval . adapters . deepeval import metrics_to_evals
evals = metrics_to_evals ([ FaithfulnessMetric ()], input_key = "question" ,
retrieval_context_key = "context" )
Adapters for deepeval , RAGAS , and promptfoo ( pip install "muteval[deepeval|ragas|promptfoo]" ). Or use the built-in framework-free
checks — including llm_judge / grounded that hit any OpenAI-compatible
endpoint (OpenAI, Groq, Gemini, GitHub Models, Ollama…) via base_url= , using
only the standard library.
Pointing muteval at a real system is a ~1-hour integration, not plug-and-play —
docs/ADOPTION.md has the honest checklist, a
"where-it-breaks → what-to-change" table, judge-selection guidance, and the four
pieces you supply. Start with muteval check and fix a green baseline first.
Shipped: prompt/context/tool/model mutation · deepeval/RAG

[truncated]
