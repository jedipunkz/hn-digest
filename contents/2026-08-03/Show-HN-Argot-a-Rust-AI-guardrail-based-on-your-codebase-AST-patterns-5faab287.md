---
source: "https://argot.tmonier.com/"
hn_url: "https://news.ycombinator.com/item?id=49154714"
title: "Show HN: Argot, a Rust AI guardrail based on your codebase AST patterns"
article_title: "argot — find code your repository would question"
author: "damienmeur"
captured_at: "2026-08-03T12:50:08Z"
capture_tool: "hn-digest"
hn_id: 49154714
score: 3
comments: 0
posted_at: "2026-08-03T12:12:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Argot, a Rust AI guardrail based on your codebase AST patterns

- HN: [49154714](https://news.ycombinator.com/item?id=49154714)
- Source: [argot.tmonier.com](https://argot.tmonier.com/)
- Score: 3
- Comments: 0
- Posted: 2026-08-03T12:12:51Z

## Translation

タイトル: ショー HN: Argot、コードベース AST パターンに基づく Rust AI ガードレール
記事のタイトル: argot — リポジトリが疑問視するコードを見つける
説明: リポジトリから構築された統計コード アナライザー
HN テキスト: 新しいフロンティア LLM のベンチマーク (SWE ベンチマークなど) を多数用意し、「最良のコード」でスコアを獲得しています。大規模なコードベースでは、コードベース自身の意見や慣例と一致するコードが最良のコードとなります。それがチームのメンタル モデルだからです。完全なコードベースの明示的/暗黙的パターンをエージェント ウィンドウに取り込んでコードベースと一致させるのは機能しません: コンテキストが多すぎる、幻覚など... そこで私はこの無料/オープン ソース プロジェクトを構築しました。コードベースは独自の音声を生成する AST コード パターンの宝庫であり、Argot はそれらを学習し、積極的または防御的にそれらを強制できるようにします。

記事本文:
メイン コンテンツにスキップ ar g ot 監査デモの仕組み ベンチマーク ドキュメント ナビゲーションを開く Tmonier · fr GitHub 監査デモ 仕組み ベンチマーク ドキュメントはここにはインポートされていません g g AI が作成したコードのハーネス · 統計、2 番目の LLM ではありません · 100% ローカル 書き留めなかったルールを Lint します。
AIがコードを書きます。 argot は、幻覚を起こせない唯一のもの、つまりリポジトリ自体の歴史を利用してそれを利用します。決定的、測定済み、ローカル。
フィルムを見る · 0:45 flag ⟺ token Surprise t max log P repo ( t ) P Baseline ( t ) + call-receiver min ( ∑ c w ( c ) , cap ) > サイズ認識しきい値 θ + β log n 90 n 97.3% の異質パターンをキャッチ ·実際の編集の 0.25% にフラグが付けられました
型チェッカーはコンパイルできるかどうかを尋ねます。アルゴットはそれがあなたのものかと尋ねます。
クリーンで型が正しい PR は、依然としてリポジトリにとって異物である可能性があります。これが実際の出力です。
有効な Python — ただし、このリポジトリがインポートしたことのないフレームワークです。証拠は、それが何を目的としているかを示しています。このリポジトリは、数か月前にリクエストを httpx に置き換えました。 argot は移行中のコミットを引用し、残りをもう 1 つ追加する前に警告します。リポジトリにはすでにこの機能があります。 argot はオリジナルと類似点に名前を付けます。正しいコード、間違ったホーム - その最も近いピアはすべてコア/ダウンローダーに存在します。ここで、cli はコアをインポートします。その逆は決して行われません。この 1 つのインポートにより、アーキテクチャが反転します。緑色は修正されずにスキップされたためです。 argot は、テストとその対象となるコードに名前を付けます。
コードを修正できないエージェントはテストを「修正」します。
緑色のチェックは弱化されたテストを隠すことができます。 argot は、それを変更されたコードと組み合わせて、両方に名前を付けます。
スキップしてください @pytest.mark.skip("flaky")
腸のアサーションは削除され、テストは維持されました
期待値 429 をリターゲット → 200 になる
削除するとテストは消えますが、コードは残ります
154/164 (93.9%) — 検出器固有のフィクスチャ

キャッチ;コントロールと受け入れられた履歴の結果は別のものです。検出器のスコープとコントロールを読んでください。
受け入れられた変更を習慣化する前に監査します。
argot Audit は、受け入れられた変更をその前のリポジトリ履歴と比較します。所見は検査を促すものであり、欠陥の判定を示すものではありません。
次に、argot init を実行し、定期的なチェック パスを選択します。
正直な数値、構造により漏れはありません。
620/637 (97.3%) — 検出器固有のフィクスチャのリコール。製品全体の精度を主張するものではありません。治具のリコール。変更されたシンボルは差分に表示されます。
36 コーパス / 12 言語 / 除外ホストのパージ後に外国インポートおよび外国 API フィクスチャ シンボルが表示される · d1007f50; generated_at=2026-07-28
264/272 (97.1%) — 検出器固有の実際の再現率。制御結果は別です。 25 のコーパスと 12 の言語にわたって内部階層化違反を作成しました。
25 コーパス / 12 言語。内部階層化違反を作成しました · 7886a7967c6fe5b822a249083c7871894db1efcf; generated_at=2026-07-20
154/164 (93.9%) — 検出器固有のフィクスチャ キャッチ。コントロールと受け入れられた履歴の結果は別のものです。制作に適したチェック用のゲーム設備を作成しました。
23 コーパス / 12 言語。テスト ゲーム フィクスチャを作成、本番環境に適合→チェック · 1b0de32d;整合性検証のみ、2026-07-28
1 つの静的バイナリ。 12 の言語 - それぞれに独自のツリーシッター アダプターと独自の学習モデルがあります。
実際のファイルに埋め込まれたパターンを思い出してください。一時的なホールドアウトでの誤警報。構造的な盲点、つまり隠蔽された異物さえも、隠されるのではなく公開されるのです。
リポジトリごとの完全な数 → 実際に見つかったものを確認する →
監査から定期的なチェックまで、お客様が選択できます。
argot init を実行し、CLI、スキル、コミットフック、または GitHub アクションを選択します。 Claude プラグインは、完全な受け入れ時間チェックではなく、狭い事前書き込みプロンプトを追加します。
7 つのオンデマンド スキル

互換性のあるホスト:
/argot-setup はツリーを読み取り、argot.toml を書き込み、キャッチを検証します
/argot-refresh はスコープとミュートを確認し、学習したスナップショットを更新します
/argot-check は各差分をスコアリングし、外部のものにフラグを立てます。決してブロックしません
/argot-review-pr は、チェックアウトなしで、リポジトリの声に照らして 1 つの PR をレビューします
/argot-setup-ci すべての PR のノンブロッキング音声スコア
/argot-write-rule は、指定した規則をテスト済みのルールに変換します。
/argot-suggest-rules は規則を見つけて成文化します。
Claude プラグインは、オプションの MCP コンテキストと狭いフェールオープン事前書き込みプロンプトを追加します。エージェントはいつ Argot に電話するかを決定します。プラグインを入手 →
または CLI を手動で操作 → CI ガイド ↗
$ argot 初期化
ステップ 1/2: 音声モデルのトレーニング …
ステップ 2/2: しきい値を調整する …
コーパス
1129 ファイル · 503 学習 · 626 テスト/ドキュメントスキップ
評決：準備完了
スナップショット → .argot/ · 音声 · セマンティックインデックス · ヘルス
レビュー + コミット → argot.toml · .argot/
次へ: argot check 学習されたスナップショットがコミットされます。キャッシュはローカルに残ります。 CI はベース ブランチのコピーを読み取りますが、適合しません。
モデル全体を一目で確認
一度学んでください。 1 つのベースラインを共有します。リポジトリが実際に移動した場合にのみ更新してください。
Argot の学習された状態は、ホストされたサービスではなく、Git でレビューされたスナップショットです。すべての開発者、エージェント、PR は同じリポジトリ メモリと比較します。
argot init は、マシン上の音声、セマンティックネイバー、アーキテクチャ、およびテスト信号を学習します。
argot.toml + .argot/ を確認してコミットします。これは、チェック用のロックファイルのような、リポジトリ固有の学習された状態です。
ローカル ツール、エージェント、CI はそのスナップショットを読み取ります。 PR はベース ブランチに対して判断されるため、PR 自体を学習することはできません。
マテリアルがドリフトを受け入れた後、ステータスは /argot-refresh を推奨します。変更されたスコープとミュートを確認し、ローカルに適合させ、再コミットします。
ワークフローで構成された PR またはプッシュ信号。
の

GitHub Action はコミットされたベース スナップショットを読み取り、デフォルトでは非ブロッキングです。意図的な相違は依然として人間の判断によるものであり、監査証跡として記録されます。
ここから開始します。レビュー キューの各行を開きます。
レアトークン 1 個 · 未知の呼び出し先 1 個
証拠: オプション、強調表示、移動 — 馴染みのない用語です。
復習：意図しない場合はリポジトリの確立された形式で書き直す。
証拠: Lazy::new は、ここでの同様のコードでは使用されていません。
復習: 呼び出しをリポジトリの確立された API と比較します。
💬 欠陥の証明ではなく、レビューを求めるプロンプト。
調査結果は PR に反映されます。スナップショットの健全性は同じ概要に表示され続けます。
意味的な理解。コアに生成 LLM はありません。
4 つのローカル エンジン、1 つの静的 Rust バイナリ — モデルは含まれていますが、何もフェッチされていません — すべてが git 履歴に基づいています。
バイナリ内のコード埋め込みモデル
コード エンコーダーから抽出された 15.6 MB のテーブル argot は、すべての関数をベクトルに変換します。argot は、あなたがこれをすでに書いたことをどのようにして認識するのでしょうか。ダウンロードするものはなく、GPU やクラウドも必要ありません。コンパイルされて出荷され、エアギャップで動作します。
2 つの頻度表と呼び出し先クラスタリング - リポジトリが実際に使用するインポート、呼び出し先、およびトークン形状。
モジュール依存関係のトポロジ。確立された方向を反転する新しいエッジには、切断された方向のフラグが立てられます。
ツリーシッターは、すべてのテストで何がアサートされるかを追跡します。製品変更の横でスキップ、削除、または削除されたテストは、ペアになって名前が付けられます。
FastAPI、ラップトップ CPU で測定。単一の静的バイナリ — Python、ノード、モデルのダウンロード、GPU はありません。
自分の規約を確認し、それを強制します。
argot 規則は、共有 API とコードが属する場所を見つけます。 1 つの規則を、テスト可能な小さなルールに変えます。
推測ではなく発見された隠語規約には、リポジトリがすでに実行していること、つまり共有 API と各種類のコードが存在する場所がリストされます。

s — つまり、ルールは argot が見つけたものから始まります。
両面 ts_query_old は、削除された変更を確認します。これは、従来のリンターでは表現することさえできないルールです。
歴史を意識した import_attested("moment") は、「これを使用したことがありますか?」と尋ねます。 — 他のリンターではできません。
テスト駆動のアーゴット ルールにより、フィクスチャ (赤/緑のオーサリング ループ) がテスト実行されます。
改ざん防止ロック = true はルールを凍結します。ミュート、ダウングレード、または書き換えを行う diff は、ルール改ざんをトリップします。固定エラー、抑制不可能、大音量の PR アノテーションです。
それぞれの種類のコードが存在する場所 — レイアウトから学習
// check.rhai — 以前に存在したルート
// この変更は現在は静かに消えています
それでは = [];
for m in ts_query(ROUTES) { now.push(m.text); }
for m in ts_query_old(ROUTES) {
if !now.contains(m.text) {
report(m.line, m.text + " 削除 —
最初に非推奨にする (docs/api-lifecycle.md)");
}
$ argot rules test no-dropped-endpoints ok fires-on-removal · ok Quiet-on-refactor CI が欠落しているレイヤーを追加します。
MIT ライセンスのオープンソース。最初に監査を行ってから、ワークフローに合った定期的なチェックを選択します。

## Original Extract

A statistical code analyzer built from your repository

We have lots of benchmarks for new frontier LLMs (SWE benchmarks etc) to make them score on the "best code". In large codebase, the best code is the one that matches the codebase own voice and conventions, because that's the mental model of the team. Ingesting a full codebase explicit/implict patterns in an agent windows to make them match the codebase doesnt work: too much context, hallucinations, ... So I built this free/open source project, your codebase is a golden mine of AST code patterns that makes its unique voice, Argot learns them and then allows you to enforce them proactively or defensively.

Skip to main content ar g ot Audit Demo How it works Benchmarks Docs Open navigation Tmonier · fr GitHub Audit Demo How it works Benchmarks Docs never imported here g g The harness for AI-written code · statistics, not a second LLM · 100% local Lint the rules you never wrote down.
AI writes the code. argot harnesses it with the one thing that can’t hallucinate: your repo’s own history . Deterministic, measured, local.
Watch the film · 0:45 flag ⟺ token surprise t max ​ lo g P repo ​ ( t ) P baseline ​ ( t ) ​ ​ ​ + call-receiver min ( ∑ c ​ w ( c ) , cap ) ​ ​ > size-aware threshold θ + β lo g n 90 ​ n ​ ​ ​ 97.3% of foreign patterns caught · 0.25% of real edits flagged
Type checkers ask if it compiles. argot asks if it’s yours.
Clean, type-correct PRs can still be foreign to your repository. This is its real output.
Valid Python — but a framework this repo has never imported. The evidence shows what it reaches for instead. This repo replaced requests with httpx months ago. argot cites the migrating commits — and warns before you add one more leftover. The repo already has this function. argot names the original and the similarity. Right code, wrong home — its nearest peers all live in core/downloader. Here, cli imports core — never the reverse. This one import flips the architecture. Green because it was skipped, not fixed. argot names the test and the code it covers.
An agent that can’t fix the code will “fix” the test.
A green check can hide a weakened test. argot pairs it with the changed code and names both.
skip it @pytest.mark.skip("flaky")
gut it assertions removed, test kept
retarget it expected 429 → becomes 200
delete it test gone, code stays
154/164 (93.9%) — detector-specific fixture catch; controls and accepted-history results are separate. Read the detector scope and controls.
Audit accepted changes before making it a habit.
argot audit compares accepted changes with the repository history before them. Findings are prompts to inspect, not defect verdicts.
Then run argot init and choose a recurring check path.
Honest numbers, leak-free by construction.
620/637 (97.3%) — detector-specific fixture recall; not a product-wide accuracy claim. Fixture recall; the changed symbol is visible in the diff.
36 corpora / 12 languages / visible foreign-import and foreign-api fixture symbols after excluded-host purge · d1007f50; generated_at=2026-07-28
264/272 (97.1%) — detector-specific real recall; control result is separate. Authored internal layering violations across 25 corpora and 12 languages.
25 corpora / 12 languages; authored internal layering violations · 7886a7967c6fe5b822a249083c7871894db1efcf; generated_at=2026-07-20
154/164 (93.9%) — detector-specific fixture catch; controls and accepted-history results are separate. Authored production fit-to-check gaming fixtures.
23 corpora / 12 languages; authored test-gaming fixtures, production fit→check · 1b0de32d; just integrity-verify, 2026-07-28
One static binary . Twelve languages — each with its own tree-sitter adapter and its own learned model:
Recall on patterns planted in real files; false alarms on a temporal holdout. Even the structural blind spot — masked foreign — is published, not hidden.
Full per-repo numbers → See it caught in the wild →
From audit to a recurring check you choose.
Run argot init , then choose the CLI, skills, a commit hook, or a GitHub Action. The Claude plugin adds a narrow pre-write prompt — not a full acceptance-time check.
seven on-demand skills for compatible hosts:
/argot-setup reads your tree, writes argot.toml, verifies the catch
/argot-refresh reviews scope and mutes, then refreshes the learned snapshot
/argot-check scores each diff, flags what’s foreign — never blocks
/argot-review-pr reviews one PR against your repo’s voice, no checkout
/argot-setup-ci a non-blocking voice score on every PR
/argot-write-rule turns a convention you state into a tested rule
/argot-suggest-rules finds your conventions, codifies one
The Claude plugin adds optional MCP context and a narrow, fail-open pre-write prompt; agents still decide when to call Argot. Get the plugin →
Or drive the CLI by hand → the CI guide ↗
$ argot init
Step 1/2: training voice model …
Step 2/2: calibrating threshold …
Corpus
1129 files · 503 learned · 626 tests/docs skipped
Verdict: Ready
Snapshot → .argot/ · voice · semantic index · health
Review + commit → argot.toml · .argot/
Next: argot check The learned snapshot is committed; caches stay local. CI reads the base branch copy and never fits.
The whole model, in one glance
Learn once. Share one baseline. Refresh only when the repo truly moves.
Argot’s learned state is a reviewed snapshot in Git , not a hosted service. Every developer, agent, and PR compares against the same repository memory.
argot init learns the voice, semantic neighbours, architecture, and test signals on your machine.
Review and commit argot.toml + .argot/ . It is repository-specific learned state, like a lockfile for checks.
Local tools, agents, and CI read that snapshot. A PR is judged against the base branch , so it cannot teach itself.
After material accepted drift, status recommends /argot-refresh : review changed scope and mutes, fit locally, recommit.
A workflow-configured PR or push signal.
The GitHub Action reads the committed base snapshot and is non-blocking by default. Intentional divergence remains a human decision, recorded as an audit trail.
Start here: open each row in the review queue.
1 rare-tokens · 1 unfamiliar-callee
Evidence: option, highlighted, Move — unfamiliar vocabulary here.
Review: rewrite in the repository’s established form if unintended.
Evidence: Lazy::new is not used by similar code here.
Review: compare the call with the repository’s established API.
💬 Prompts for review, not proof of defects.
Findings land on the PR. Snapshot health stays visible in the same summary.
Semantic understanding. No generative LLM in the core.
Four local engines, one static Rust binary — model included, nothing fetched — all grounded in your git history.
A code-embedding model inside the binary
A 15.6 MB table argot distilled from a code encoder turns every function into a vector — how argot knows you already wrote this . Nothing to download, no GPU, no cloud: it ships compiled in and works air-gapped.
Two frequency tables and a callee clustering — the imports, callees, and token shapes your repo actually uses .
Your module-dependency topology. A new edge that reverses the established direction is flagged with the direction it breaks.
tree-sitter tracks what every test asserts. A test skipped, gutted, or deleted beside a prod change gets paired and named.
Measured on FastAPI, laptop CPU. Single static binary — no Python, no Node, no model download, no GPU.
See your conventions — then enforce them.
argot conventions finds the shared API and where code belongs. Turn one convention into a small, testable rule.
Discovered, not guessed argot conventions lists what your repo already does — its shared API, and where each kind of code lives — so a rule starts from what argot found.
Two-sided ts_query_old sees what a change removed — a rule no classic linter can even express.
History-aware import_attested("moment") asks “have we ever used this?” — no other linter can.
Test-driven argot rules test runs your fixtures — the red/green authoring loop .
Tamper-evident locked = true freezes a rule; a diff that mutes, downgrades, or rewrites it trips rule-tampered — pinned error, unsuppressable, a loud PR annotation.
where each kind of code lives — learned from your layout
// check.rhai — a route that existed before
// this change, and is silently gone now
let now = [];
for m in ts_query(ROUTES) { now.push(m.text); }
for m in ts_query_old(ROUTES) {
if !now.contains(m.text) {
report(m.line, m.text + " removed —
deprecate first (docs/api-lifecycle.md)");
}
} $ argot rules test no-dropped-endpoints ok fires-on-removal · ok quiet-on-refactor Add the layer your CI is missing.
MIT-licensed open source. Audit first, then choose the recurring check that fits your workflow.
