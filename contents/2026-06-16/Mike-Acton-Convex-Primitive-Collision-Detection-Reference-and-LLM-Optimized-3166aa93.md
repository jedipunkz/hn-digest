---
source: "https://github.com/macton/differentiable-collisions-optc"
hn_url: "https://news.ycombinator.com/item?id=48552679"
title: "Mike Acton: Convex Primitive Collision Detection – Reference and LLM-Optimized"
article_title: "GitHub - macton/differentiable-collisions-optc · GitHub"
author: "thdr"
captured_at: "2026-06-16T10:41:28Z"
capture_tool: "hn-digest"
hn_id: 48552679
score: 2
comments: 0
posted_at: "2026-06-16T09:25:33Z"
tags:
  - hacker-news
  - translated
---

# Mike Acton: Convex Primitive Collision Detection – Reference and LLM-Optimized

- HN: [48552679](https://news.ycombinator.com/item?id=48552679)
- Source: [github.com](https://github.com/macton/differentiable-collisions-optc)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T09:25:33Z

## Translation

タイトル: Mike Acton: 凸プリミティブ衝突検出 – リファレンスおよび LLM 最適化
記事のタイトル: GitHub - macton/Differenceiable-collisions-optc · GitHub
説明: GitHub でアカウントを作成して、macton/Differenceiable-collisions-optc の開発に貢献します。

記事本文:
GitHub - macton/Differenceiable-collisions-optc · GitHub
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
マクトン
/
微分可能な衝突オプティカル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マクトン/微分可能衝突-o

ptc
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット ビルド ビルド ドキュメント ドキュメント パフォーマンス テスト最適化 パフォーマンス テスト最適化 パフォーマンス テスト パフォーマンス テスト プロンプト プロンプト src-optimized src-optimized src src test test viz viz .gitignore .gitignore LICENSE LICENSE Makefile Makefile Makefile.optimized Makefile.optimized README.md README.md bin_format.h bin_format.h robe-optimized-harness.sh give-optimized-harness.sh run-performance-test run-performance-test run-performance-test-optimized run-performance-test-optimized すべてのファイルを表示 リポジトリ ファイルのナビゲーション
凸プリミティブ衝突検出 — リファレンスおよび LLM 最適化
このリポジトリは、K. Tracy、T. A. Howell、および
Z. Manchester、「凸集合の微分可能な衝突検出」
プリミティブ」 (arXiv:2207.00669、documents/2207.00669.pdf )。一対の
凸プリミティブ — 球、ボックス、カプセル、または凸多面体 — を計算します。
両方の形状に適用する必要がある最小均一スケーリング α
touch (紙の問題 (10))、および式からの接触点。 (24)。 α < 1
α > 1 はそれらが重なっていることを意味し、α > 1 はそれらが分離されていることを意味します。
これは狭位相ソルバーです。呼び出し元がすでに安価なサービスを実行していることを前提としています。
ワールド AABB が重複しないブロードフェーズおよび破棄されたペアのみ
AABB が重複するペアは常にクエリされます。コミットされたベンチマークはそれを反映しています
仮定 - その 1000 ペアはすべて AABB オーバーラップ (近接接触または
貫通している）ので、タイミングは実際の狭い位相の仕事を測定します。
遠く離れた形状の些細な拒否。
ここには 2 つの実装があります。
src/ — 論文に直接従うリファレンス C 実装。
src-optimized/ — 最適化された単精度実装

それ
同じ衝突フラグと同じ距離 (指定された範囲内) が生成されます。
許容差)、コミットされた 1000 ペアのベンチマークを約 102 倍高速に実行します
リファレンスより: リファレンスの中央値 ≈ 0.276 秒、最適化された中央値 ≈ 0.0027 秒
(私のマシンでは中央値 5、シングル スレッド — gcc 11、x86-64、WSL2)。
この 102 倍は、コミットされたベンチマークに対して設定した 100 倍の目標を超えました。それも
そのベンチマークを維持します。代替シード入力では 97.6 ～ 101.7 倍と測定されます。
(シードは 4 つ)、すべて合格です。一律100倍とは言いません —
4 つの種子のうち 2 つは真下に着地します。つまり、私は「コミットされた種子の 100 倍」と主張します。
ベンチマーク、一般に ~98 ～ 102 倍」、そしてそれ以上はありません。数値と注意事項は次のとおりです
結果と限界。
同様に重要な 2 つの理由:
最適化された衝突ルーチンを提供する。 src-optimized/ は実数です、
構築して使用できるテスト済みのコード。独立した機関によって参照が保持されます。
ハーネス。
LLM を使用して最適化を行う方法を具体的に示します。
再現可能に。このプロジェクトのすべてのフェーズは言語モデルによって生成されました
私が作成した指示文書に基づいて作成され、すべての結果は担当者によってチェックされました。
モデルが歩き回ることができないハーネス。メソッドはこうであってほしい
検証可能ですが、信じて受け入れなければならない話ではありません。
ここでテストしたモデルは GPT-5.5 です。これは 1 つのモデル、1 つの実行、つまりケースです
ベンチマークではなく、最適化問題で LLM を駆動する方法の研究
モデルの比較。
4 つの役割を明確に分離するのが最も明確だと思います。
100 倍の目標の出典: リファレンス コードを読んで、
このハードウェアで何が達成できると私が考えたかについての判断。それはではありません
導出された限界または上限の証明 — それはエンジニアの推定値であり、私は
それを一つとして述べます。それはおおよそ正しい大きさであることが判明しました
強く押しつける。

私のアプローチ自体は、 context/data-owned-design.md に書き留められています。それら
運用ルール — 実際のデータから開始し、コストを明示し、事前に作業を削除します。
より高速に実行し、一般的なケースを直線的に処理します。
最適化に関する会話。つまり、「私が貢献したこと」だけが目標ではなく、
プロンプト;それはモデルが従うように作られた方法です。
構造化された状態とターンごとの証明は儀式ではありません。 LLM が残したもの
最適化自体はドリフトする傾向があります。前回の最適化の記憶から推論します。
新しい測定値の代わりに結果が得られると、決して得られなかった良好な変化が失われる可能性があります。
コミットされ、実際には実行されなかった結果が報告される可能性があります。を維持する
作業状態を検査可能なファイルに保存し、保持されているすべてのゲインを即座にコミットします。
毎ターン実際のゲートとスピードアップのステータスを注入することが「モデルを変える」ものです
「より速くて正しい」を「より速く測定され、ゲートを通過し、コミットされた」に変換します。
それがデモと結果の違いです。
方法 — 4 つの文書、4 つのフェーズ
各フェーズは指示文書とそれによって生成された成果物です。の
ドキュメントはプロンプト/ に存在します。
プロンプト/create-reference.md → 参照 ( src/ )。忠実なC11
論文のポート: α 解と式からの接触点。 (24)、付き
明示的な入力検証。これが正しさのアンカーです。他のものはすべてそうです
に対して測定されます。
Prompts/create-optimized-test-harness.md → ハーネス。これは指定します
テスト、比較、測定の足場とその制約: 固定、
1000 ペアの入力をコミット。参照対最適化コンパレータ。の
どちらのソルバーともコードを共有しない独立したバリデータ。連絡先
認証者;決定性チェック。中央値 5 のタイミング プロトコル。重要なことは、
ハーネスは作成され、次の ID コピーに対して証明されました。

参照
最適化が存在する前 (「
Performance-test-optimized/HARNESS-BASELINE.md )、測定パイプライン
それ自体は、何かを判断するために使用される前に信頼されていました。
プロンプト/create-optimized.md → 最適化されたソルバー ( src-optimized/ )。
これは、モデルが反復する命令に変換された私の最適化アプローチです。
on: サイクルの行き先をプロファイルし、利益によって候補者をランク付けし、削除を優先します。
作業し、単純化パスを実行し、一般的なケースのブランチを最小限に抑え、
データのレイアウトとバッチ処理を最上級のものとして扱います。モデルは内部でこのループを実行しました
nagent ( https://github.com/macton/nagent ) — データ指向のエージェント ループ
ここでは、作業状態はプレーン ファイル内に存在し、モデルはそれを通じてのみ動作します。
構造化タグの固定セット。証明ハーネスは、1 回につき 1 回実行されるように配線されました。
ターン：
nagent --read プロンプト/create-optimized.md \
--hook-per-run ./prove-optimized-harness.sh \
「目標の 100 倍に達するまで続けます。」
したがって、すべてのターンは実際に測定されたゲート ステータスを注入することから始まりました。
会話 — モデルの記憶ではありません。
プロンプト/create-visualizer.md → ビジュアライザー ( viz/ )。説明済み
以下。
何が最適化され、何が拒否されたのか
測定と保持/元に戻す決定を含む、仮説ごとの完全な履歴は次のとおりです。
src-optimized/OPTIMIZATION-LOG.md 内。 git 履歴はそれを反映しています: 1 つのコミット
保持された変更ごとに、拒否された各トライアルを記録するコミットも含まれます。の形状
進歩はいかなる単一のステップよりも重要です。それは漸進的であり、測定され、そして
可逆的であり、行き止まりは隠されるのではなく書き留められました。
参照の対数バリア ニュートン ソルブをサポート/GJK + 二分法で置き換えます。
α の計算 — 単一の最大の勝利。
タイプごとの特殊化: ボックスボックスとボックスの分離軸 (SAT) パス
ボックスポリトープの非対称 SAT。シフトされたGJKパット

hs は球体/カプセル-ポリトープを表します。
シェイプごとの作業を、ビルド段階の事前計算に移動します。
時間指定ソルブ (ランタイムはフラットな事前計算テーブルからソルバを実行します)。
全体的に単精度で、各ペアをメーターに再センタリングすることで安全性が向上
解決する前にスケールします。
グローバルな多面体ハーフスペースを前もって構築するのはやめてください。いくつかの軸を計算します
ペアが実際に必要とするポリトープの一意のハル エッジを事前計算します。
ボックスポリSAT。
アクティブパスのビルド状態を圧縮します。ホットなものを特化して強制的にインライン化する
サポート機能。
半径形状ファミリの閉じた形式 (解析) 接触証人
(球/カプセル、ボックス-カプセル、球/カプセル-ポリトープ)、GJK を回避します。
ジオメトリがそれを可能にする場所を目撃してください。
追加のステップが発生しなかった場合の二等分/リファインメントの反復回数を削減します。
許容範囲内で結果を変更します。
拒否されました (非表示ではなく記録されました): ボックスポリシフト GJK パスとボックスポリ
SAT パスが退行したか、許容範囲/フラグ契約を破った。いくつかの
インライン化、ブラケット化、イテレーションキャップのトライアルは目に見えて役に立ちませんでした。ある
ソルブラッパー内のコピーと削除。そしてさまざまな証人簿記の変更。それぞれ
は 1 行のコミットと、削除された理由を含むログ エントリです。
ログには、各仮説のコスト (時計とトークン) も記録されます。
そのため、結果だけでなく、演習全体のコストが表示されます。
一致規約 — 「より速い」は「ビットが同一」ではない
最適化されたソルバーは、リファレンスとビットごとに同一ではなく、
そうあるべきだ。次の場合にのみ受け入れられます。
衝突フラグは同一です。つまり、衝突フラグとまったく同じペアにフラグを立てます。
基準として衝突します。そして
すべての距離は以内に一致します
|Δ| ≤ 1 mm + 0.1%·|d_ref| + 5e-4·(|c1−c2|/α²) ( build/compare_results )。
1 mm フロアは文書化された解像度です。リラ

有効な用語は広範囲をカバーします
別居。最後のは条件付け項です。固定 α 誤差は次のようにスケールされます。
|c1−c2|/α² なので、単精度の場合、極度の浸透でのみ増加します。
本当に深さを解決できず、値は最も実用的ではありません。
接触点の有効性は証明されていますが、一致していません: 面またはエッジ
contact には同様に有効な証人ポイントが多数あるため、build/validate_contacts
放出された各ポイントが両方のサーフェス上に存在し、適切であることを個別にチェックします。
に等しいことが要求されるのではなく、報告された距離によって分離されます。
参考人の選択。
viz/ は、自己完結型の小さな Web ツール (prompts/create-visualizer.md) です。
一度に 1 つのクエリ ペアをレンダリングします: 2 つのプリミティブ、放出されたコンタクト ポイント
参照ソルバーと最適化ソルバーの両方、およびそれらの間の分離によって。
結果が幾何学的に正常であるかどうかは、このようにして判断されます。
公差番号。この README 内の画像は、それによって作成されました。
# `make -f Makefile.optimized 最適化` を実行して、
# 2 つの結果ファイル:
CD ビジュアライゼーション
python3 -m http.server 8000 # ES モジュールにはオリジンが必要です
# http://localhost:8000/index.html を開く
私のマシン (gcc 11.4.0、x86-64、WSL2) で測定し、1000 ペアの入力をコミットしました。
中央値 5、シングルスレッド。
スピードアップ: コミットされた入力で ≈ 102 倍 (基準中央値 ≈

[切り捨てられた]

## Original Extract

Contribute to macton/differentiable-collisions-optc development by creating an account on GitHub.

GitHub - macton/differentiable-collisions-optc · GitHub
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
macton
/
differentiable-collisions-optc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
macton/differentiable-collisions-optc
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits build build documents documents performance-test-optimized performance-test-optimized performance-test performance-test prompts prompts src-optimized src-optimized src src test test viz viz .gitignore .gitignore LICENSE LICENSE Makefile Makefile Makefile.optimized Makefile.optimized README.md README.md bin_format.h bin_format.h prove-optimized-harness.sh prove-optimized-harness.sh run-performance-test run-performance-test run-performance-test-optimized run-performance-test-optimized View all files Repository files navigation
Convex Primitive Collision Detection — reference and LLM-optimized
This repository implements the collision query from K. Tracy, T. A. Howell, and
Z. Manchester, "Differentiable Collision Detection for a Set of Convex
Primitives" (arXiv:2207.00669, documents/2207.00669.pdf ). For a pair of
convex primitives — sphere, box, capsule, or convex polytope — it computes the
minimum uniform scaling α that must be applied to both shapes for them to
touch (the paper's problem (10)), and the contact points from eq. (24). α < 1
means they overlap, α > 1 means they are separated.
This is a narrow-phase solver. It assumes the caller has already run a cheap
broadphase and discarded pairs whose world AABBs do not overlap, so only
AABB-overlapping pairs are ever queried. The committed benchmark reflects that
assumption — its 1000 pairs are all AABB-overlapping (near-contact or
penetrating), so the timing measures real narrow-phase work rather than the
trivial rejection of far-apart shapes.
There are two implementations here:
src/ — a reference C implementation that follows the paper directly.
src-optimized/ — an optimized single-precision implementation that
produces the same collision flags and the same distances (within a stated
tolerance) and runs the committed 1000-pair benchmark about 102× faster
than the reference: reference median ≈ 0.276 s, optimized median ≈ 0.0027 s
(median-of-5, single thread, on my machine — gcc 11, x86-64, WSL2).
That 102× crossed the 100× target I set for the committed benchmark. It also
holds up off that benchmark: on alternate-seed inputs it measures 97.6–101.7×
(four seeds), all passing correctness. I would not call it a uniform 100× —
two of the four seeds land just under — so I claim "100× on the committed
benchmark, ~98–102× generally," and no more. Numbers and caveats are in
Results and limits .
Two reasons, equally important:
To provide the optimized collision routines. src-optimized/ is real,
tested code you can build and use, held to the reference by an independent
harness.
To show how an LLM was used to do the optimization — concretely and
reproducibly. Every phase of this project was generated by a language model
from an instruction document I wrote, and every result was checked by a
harness that the model could not talk its way around. I want the method to be
inspectable, not a story you have to take on faith.
The model under test here was GPT-5.5 . This is one model, one run — a case
study in how to drive an LLM at an optimization problem, not a benchmark
comparing models.
I find it clearest to separate the four roles explicitly.
Where the 100× target came from: I read the reference code and made a
judgment call about what I thought was achievable on this hardware. It is not a
derived bound or a proof of a ceiling — it is an engineer's estimate, and I
state it as one. It turned out to be roughly the right order of magnitude to
push hard against.
My approach is itself written down, in context/data-oriented-design.md . Those
operating rules — start from the real data, state the cost, remove work before
doing it faster, handle the common case straight-line — were injected into every
optimization conversation. So "what I contributed" is not just the target and
the prompts; it is the method the model was made to follow.
The structured state and the per-turn proof are not ceremony. An LLM left to
optimize on its own tends to drift: it reasons from its recollection of the last
result instead of a fresh measurement, it can lose a good change that was never
committed, and it can report a result it did not actually run. Keeping the
working state in inspectable files, committing every kept gain immediately, and
injecting the real gate-and-speedup status every turn are what turn "the model
says it is faster and correct" into "measured faster, gates pass, committed."
That is the difference between a demo and a result.
Method — four documents, four phases
Each phase is an instruction document and the artifact it produced. The
documents live in prompts/ .
prompts/create-reference.md → the reference ( src/ ). A faithful C11
port of the paper: the α solve and the contact points from eq. (24), with
explicit input validation. This is the correctness anchor everything else is
measured against.
prompts/create-optimized-test-harness.md → the harness. This specifies
the test, comparison, and measurement scaffold and its constraints: a fixed,
committed 1000-pair input; a reference-vs-optimized comparator; an
independent validator that shares no code with either solver; a contact-point
certifier; a determinism check; and a median-of-5 timing protocol. Crucially,
the harness was built and proven against an identity copy of the reference
before any optimization existed (see
performance-test-optimized/HARNESS-BASELINE.md ), so the measurement pipeline
itself was trusted before it was used to judge anything.
prompts/create-optimized.md → the optimized solver ( src-optimized/ ).
This is my optimization approach turned into instructions the model iterates
on: profile where the cycles go, rank candidates by payoff, prefer removing
work, run a simplification pass, keep the common case branch-minimal, and
treat data layout and batching as first-class. The model ran this loop inside
nagent ( https://github.com/macton/nagent ) — a data-oriented agent loop
where the working state lives in plain files and the model acts only through
a fixed set of structured tags. The proof harness was wired to run once per
turn :
nagent --read prompts/create-optimized.md \
--hook-per-run ./prove-optimized-harness.sh \
"Continue until 100x target reached."
so every turn began with the real, measured gate status injected into the
conversation — not the model's memory of it.
prompts/create-visualizer.md → the visualizer ( viz/ ). Described
below.
What was optimized — and what was rejected
The full per-hypothesis history, with measurements and keep/revert decisions, is
in src-optimized/OPTIMIZATION-LOG.md . The git history mirrors it: one commit
per kept change, plus a commit recording each rejected trial. The shape of the
progress matters more than any single step — it was incremental, measured, and
reversible, and the dead ends were written down rather than hidden.
Replace the reference's log-barrier Newton solve with a support/GJK + bisection
computation of α — the single largest win.
Per-type specializations: separating-axis (SAT) paths for box-box and an
asymmetric SAT for box-polytope; shifted GJK paths for sphere/capsule-polytope.
Move per-shape work into a build-stage precompute that is excluded from the
timed solve (the runtime solves from a flat precomputed table).
Single precision throughout, made safe by re-centering each pair to metre
scale before solving.
Stop building global polytope half-spaces up front; compute the few axes a
pair actually needs, and precompute the polytope's unique hull edges for the
box-poly SAT.
Compact the active-path build state; specialize and force-inline the hot
support function.
Closed-form (analytic) contact witnesses for the radius-shape families
(sphere/capsule, box-capsule, sphere/capsule-polytope), avoiding GJK for the
witness where the geometry allows it.
Reduce bisection/refinement iteration counts where the extra steps did not
change the result within tolerance.
Rejected (recorded, not hidden): a box-poly shifted-GJK path and a box-poly
SAT path that either regressed or broke the tolerance/flag contract; several
inlining, bracketing, and iteration-cap trials that did not measurably help; a
copy-removal in the solve wrapper; and assorted witness-bookkeeping changes. Each
is a one-line commit and a log entry with the reason it was dropped.
The log also records the cost of each hypothesis — wall-clock and tokens —
so the price of the whole exercise is visible, not just the result.
The match contract — "faster" is not "bit-identical"
The optimized solver is not bit-for-bit identical to the reference, and it is not
supposed to be. It is accepted only when:
the collision flags are identical — it flags exactly the same pairs as
colliding as the reference; and
every distance agrees within
|Δ| ≤ 1 mm + 0.1%·|d_ref| + 5e-4·(|c1−c2|/α²) ( build/compare_results ).
The 1 mm floor is the documented resolution; the relative term covers large
separations; the last is a conditioning term — a fixed α error scales by
|c1−c2|/α² , so it grows only at extreme penetration, where single precision
genuinely cannot resolve the depth and the value is least actionable.
Contact points are certified for validity, not matched : a face or edge
contact has many equally valid witness points, so build/validate_contacts
independently checks that each emitted point lies on both surfaces and is
separated by the reported distance, rather than requiring it to equal the
reference's choice.
viz/ is a small, self-contained web tool ( prompts/create-visualizer.md ) that
renders one query pair at a time: the two primitives, the contact points emitted
by both the reference and the optimized solver, and the separation between them.
It is how I eyeball that a result is geometrically sane, not just within a
tolerance number. The images in this README were produced by it.
# after `make -f Makefile.optimized optimized` and a run that produced the
# two results files:
cd viz
python3 -m http.server 8000 # ES modules need an origin
# open http://localhost:8000/index.html
Measured on my machine (gcc 11.4.0, x86-64, WSL2), committed 1000-pair input,
median-of-5, single thread.
Speedup: ≈ 102× on the committed input (reference median ≈

[truncated]
