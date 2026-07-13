---
source: "https://github.com/e35zhang/property-testing-skill"
hn_url: "https://news.ycombinator.com/item?id=48894904"
title: "Show HN: A prompt to make LLMs write property-based tests"
article_title: "GitHub - e35zhang/property-testing-skill · GitHub"
author: "e35zhang"
captured_at: "2026-07-13T17:08:07Z"
capture_tool: "hn-digest"
hn_id: 48894904
score: 2
comments: 1
posted_at: "2026-07-13T16:17:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A prompt to make LLMs write property-based tests

- HN: [48894904](https://news.ycombinator.com/item?id=48894904)
- Source: [github.com](https://github.com/e35zhang/property-testing-skill)
- Score: 2
- Comments: 1
- Posted: 2026-07-13T16:17:55Z

## Translation

タイトル: HN の表示: LLM にプロパティベースのテストを作成させるためのプロンプト
記事のタイトル: GitHub - e35zhang/property-testing-skill · GitHub
説明: GitHub でアカウントを作成して、e35zhang/property-testing-skill の開発に貢献します。

記事本文:
GitHub - e35zhang/property-testing-skill · GitHub
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
e35zhang
/
プロパティテストスキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
e35zhang/プロパティ-テスト-スキル
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
名前
プロパティファーストテスト
説明
設計テスト (および保護されるアーキテクチャ) は、機能ごとではなく形式的な分解によってテストされます。テストの作成、テスト戦略の設計、コンポーネントの正確性のモデル化、テスト対象の決定、またはテスト スイートのレビューを求められた場合に使用します。 「X のテストの作成」をワークフローに変換します。X を直交プロパティに分解し、プロパティごとに 1 つの基底テストを作成して、それを個別に証明します。次に、TLA+ スタイルのスパン テストまたはプロパティ テスト ライブラリを作成して、到達可能なすべての構成が不変条件を保持していることを証明します。 LLM がプロパティの代わりに出力をアサートするハッピー パス、データ フロー、特徴形状のテストを発行するデフォルトの障害モードに対抗します。
プロパティファーストのテスト
設計では、問題を分解してアーキテクチャを設計する方法をテストします。
直交するプロパティ、それらのプロパティを動作に組み込む、そして
両方の層のテスト - 基底プロパティごとに 1 つのテスト (ベクトルの証明)、および
TLA+ スタイルのテストは、到達可能な状態空間にわたってテストします (スパンを証明します)。
このスキルは、1 つの特定の障害モードを克服するために存在します。 「テストを書く」ように依頼されました
X」の言語モデルは、デフォルトで機能型、データフロー、ハッピーパス テストになります。
関数を呼び出し、厳選されたいくつかの入力に対して戻り値をアサートします。
これらのテストは、プロパティではなくサンプルをエンコードします。システムが作動している間、彼らは通過します
深く壊れており、コードがたまたま持っていた動作をすべてロックしてしまいます。
それらが書かれた日に。このスキルは、その反射神経を規律ある行動に置き換えます。
思考の流れ。
核心的な主張: ソフトウェアエンジニアリングは分解と合成である
問題。正しさは例のリストではありません。それはプロパティのセットです
それはOを保持します

状態空間です。テストは証明であり、証明には構造があります。
メンタルモデル: 建築はベクトル空間である
十分に複雑なドメインはすべて、少数の独立したドメインに分解されます。
変動軸 — ドメインの基底ベクトル。各軸には、
他のものでは合成できない固有の特性。この直交性は、
デザイン上の選択ではありません。それはあなたが発見したドメインに関する事実です。
数学に例えるとそれが具体的になります。 3 つの関数を「基底ベクトル」として考えます。
どのようなパラメータであっても、線形結合 a·x + b·e^x は周期的ではありません。の
プロパティ「periodic」は、x と e^x が単純に持たない基底ベクトルに基づいて存在します。
スパン。それが直交性の意味です。つまり、それから構成できない特性です。
他のコンポーネント。
ソフトウェアも同じように動作します。サブシステムの正確性のプロパティ
(「ラベルはライブ中に再利用されることはありません」、「カウントはチャーン全体で維持されます」、
「選択は到着順に依存しません」) は特定の軸上に存在します。テストしたら
特徴ごとに、空間内の点をサンプリングして希望します。プロパティごとにテストすると、
基底ベクトルを証明してから、そのスパンを証明します。
テストの 2 層は、このモデルから直接続きます。
最初の層のみを持つスイートは部分を証明しますが、全体を証明しません。スイート
2 番目のみが全体を証明しますが、壊れたときに何もローカライズされません。
両方欲しいですよね。これらは冗長ではなく、直交するリスクをカバーします。
タスクが次のいずれかである場合は常にこれを使用します。
「このモジュール/関数/サブシステムのテストを作成します。」
「テスト戦略を設計する」または「テスト カバレッジを改善する」。
「X の正しさをモデル化する」/「X の不変条件は何ですか?」
「これらのテストを確認する」 - プロパティまたはサンプルがエンコードされているかどうかを監査します。
表面が大きく、例が恣意的であると感じられる場合に、何をテストするかを決定します。
Assert f(2, 3) == 5 を書こうとしていることに気付いた場合

、停止して実行します。
ワークフロー。そのアサーションはサンプリングされたポイントです。サンプルとなるプロパティを見つけます。
分解 合成 ベースのテスト スパンのテスト
(軸を特定する) → (代数を述べる) → (各ベクトルを証明する) → (スパンを証明する)
A相 B相 C相 D相
フェーズ A — 分解 (システムの識別)
根拠を見つけてください。これは難しい 10% であり、人間/LLM 推論によるステップであり、そうではありません。
コーディングステップ。
境界に名前を付けます。サブシステムをマップとして記述します。その入力は何か、その入力は何か、
アウトプットと、それが橋渡しする 2 つの世界。一文を書いてください:
「X は ⟨入力⟩ を ⟨出力⟩ に変換し、⟨状態⟩ を維持します。」
行動を列挙します。サブシステムが行うすべてのことをリストします。整理しないでください
まだ — 列挙してみましょう。それぞれの動作はデータ ポイントです。
変化の軸を見つけます。行動のペアごとに次のように尋ねます。「もし私が変わったら
Aさん、Bさんは変わらないといけないでしょうか？」 「いいえ」の場合、A と B は異なる軸上にあります。軸とは、
独立した変化の方向性。
- 軸 1 に沿った変更には、軸 2 に沿った変更は必要ありません。
- 軸 1 のプロパティは、軸 2 を参照せずに記述および検証できます。
「固有基準」に換算します。最小の、全域にわたる独立したセットを見つけます。
すべての動作を生成するプロパティ:
- スパンニング — すべての動作は基本プロパティの構成です。
- 独立 — 他の基底プロパティから導出可能な基底プロパティはありません。
- 最小限 — どれか 1 つを削除すると、一部の動作は表現できなくなります。
各軸のプロパティを記述します。すべての基底ベクトルについて、
不変条件、またはそれを単独で特徴付ける法則。この文は になります
フェーズ C の基礎テスト。正確に名前を付けるには、以下の分類法を使用します。
フェーズ A の出力: それぞれ 1 行の正式なプロパティを持つ軸の短いリスト。
フェーズ B — 構成 (代数の説明)
システムを構成として表現し、作成するステートマシンを書き留めます

私
後でチェックしてください。
特徴＝構成。外部から見える動作ごとに、次のように記述します。
動作 = compose(prop_axis_i, prop_axis_j, ...) 。全体で共有されるプロパティ
多くの動作は共有コンポーネントです (これはあなたのアーキテクチャでもあります: 直交
軸 → 相互インポートのないモジュール、DAG として階層化されます)。
フェーズ D でモデルチェックするステート マシンを定義します。
- 変数 — 不変条件が語る最小の状態。
- Init — クリーンな開始状態 (すべてのテスト動作はここから始まります)。
- 次へ — アクションの論理和: 演算の有限アルファベット
それが遷移 (作成、削除、再構成、クラッシュなど) を引き起こします。
- 不変条件 — フェーズ A のプロパティ。今度は述語として使用されます。
到達可能なすべての状態で保持する必要がある変数。
フェーズ B の出力: アクション アルファベット + 不変述語。
フェーズ C — 基礎をテストする (各ベクトルを証明する)
軸ごとに、その固有の特性を単独で証明するテストを作成します。
実際のコンポーネントを実行し、境界を越えてのみモックします。交換のみ
特定された境界の向こう側に何が存在するか（永続性、
ダウンストリーム サブシステム、クロック)。実際の運用コードのパスを実行します。を縮小する
テストを高速化するためのリソース (128K テーブル → 256 エントリ) は、次の条件を満たす限り問題ありません。
同じ製品コードが実行されます。
出力ではなくプロパティをアサートします。建物全体を歩いて確認してください。
戻り値だけでなく、すべての要素に対して不変です。 (例: 「
USED リストには IN-TREE フラグも含まれています - すべてのノードにわたってチェックされます。)
プロパティごとにホワイトボックスまたはブラックボックスを選択します。構造的不変条件
(リスト/ツリーの整合性、保全) ホワイトボックスの状態検査が必要。行動的な
不変条件 (出力選択) は、発行されたメッセージ上でブラックボックスになる可能性があります。ピック
物件に必要なもの。
テストごとに 1 つのプロパティ。テストアッサーの場合

2 つの無関係なプロパティです。
2 つの軸を同時にサンプリングするため、位置を特定するのが困難になります。
フェーズ C の出力: 基本プロパティごとに 1 つの集中テスト。
フェーズ D — スパンのテスト (TLA+ スタイル)
組成を証明してください。これは、最も適合する言語またはプロパティ テスト ライブラリで実行される、境界付きモデル チェックです。基本事項には次のものが含まれます。
到達可能な状態空間を列挙します。すべてのインターリーブを生成します。
制限された長さまでのアクションアルファベット (例: 長さ 1..3 のすべてのシーケンス、および
削除、再構成、再追加など）。これは TLC の仕事であり、手作業で行われます。
すべてのアクションをあらゆる順序で適用することで、到達可能な状態を探索します。
到達可能なすべての状態で合成された不変式をアサートします (各状態の後で)。
最後だけでなくアクションも。
フォールトをファーストクラスのアクションとして挿入します。クラッシュ/再起動は別のことです
アルファベットのアクション。その後、回復不変条件をアサートします (状態が復元され、
何も失われず、何も複製されません）。
スパンが大きすぎないか確認してください。違法国家であると主張する
到達不能 — 「単調基底からの周期性がない」チェック。アクションの場合
シーケンスが、コードまたはモデルのいずれかが、不変式が禁止する状態に達しました。
間違っています。
フェーズ D の出力: 順列ハーネス + 不変チェック (+ オプションのベースライン)。
物件に正確に名前を付けることが最も重要です - 適切な言葉がそれを物語ります
どのようなテストを書くか。アサーションを作成する前に、次の点に注意してください。
プロパティに名前を付けることができない場合は、フェーズ A が完了していません。戻ってください。
基礎テスト (フェーズ C) — 1 つのベクトルを単独で証明する
test_<軸>_<プロパティ>():
# 配置: 実際のコンポーネント、境界を越えたモックのみ
init_real_component(size=SMALL) # リソースを縮小、同じ製品コード
mock_the_thing_on_the_other_side() # 永続性 / ダウンストリーム / クロック
# ACT: 単軸を駆動します
...これのみを実行する操作を適用します

財産 ...
# ASSERT: 構造全体にわたる固有の不変式
walk(構造) の要素の場合:
assert invariant_holds(element) # 戻り値だけではない
assert conserved_quantity == Expected # 例:使用済み + 無料 == 合計
スパン ハーネス (フェーズ D) — 到達可能な状態にわたる構成を証明する
ALPHABET = [作成_a、作成_b、削除_a、削除_b、再構成、クラッシュ]
for seq inbounded_interleavings(ALPHABET, max_len=3): # + 除去/フラップ
状態 = 初期化()
シーケンス内のアクションの場合:
state = apply(アクション, 状態)
assert generated_invariant(state) # すべてのアクションの後にチェックされます
TLA+ スケッチ (フェーズ B) — ハーネスが近似するモデル
VARIABLES \* 不変式が語る最小の状態
Init == s = <クリーンな初期状態>
次 == \/ CreateA(s) \/ CreateB(s)
\/ A を削除します \/ B を削除します
\/ 再構成 \/ クラッシュ \* アクションのアルファベット
Inv == /\ Conserved(s) \* カウント不変
/\ FlagsConsistent(s) \* 構造不変式
/\ NoDangling(s) \* 参照整合性
\* Spec == Init /\ [][Next]_s ; Inv が Spec の不変式であることを確認してください
メリットを得るために TLC を実行する必要はありません。 4 つの部分を書く —
VARIABLES / Init / Next / Inv — は設計行為です。の順列ハーネス
フェーズ D は、 Init /\ [][Next]_s ⇒ []Inv の実行可能な近似です。
実用的な例 (要約 — 境界のあるレウサ)

[切り捨てられた]

## Original Extract

Contribute to e35zhang/property-testing-skill development by creating an account on GitHub.

GitHub - e35zhang/property-testing-skill · GitHub
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
e35zhang
/
property-testing-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
e35zhang/property-testing-skill
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits LICENSE LICENSE README.md README.md View all files Repository files navigation
name
property-first-testing
description
Design tests (and the architecture they protect) by formal decomposition instead of by feature. Use when asked to write tests, design a test strategy, model a component's correctness, decide what to test, or review a test suite. Turns "write tests for X" into a workflow — decompose X into orthogonal properties, write one basis test per property proving it in isolation, then write TLA+-style span tests or property testing libraries proving every reachable composition preserves the invariants. Counters the default failure mode where an LLM emits happy-path, data-flow, feature-shaped tests that assert outputs instead of properties.
Property-First Testing
Design tests the way you design architecture: by decomposing a problem into
orthogonal properties, composing those properties into behavior, and then
testing both tiers — one test per basis property (proving the vector), and
TLA+-style tests over the reachable state space (proving the span).
This skill exists to defeat one specific failure mode. Asked to "write tests for
X," a language model defaults to feature-shaped, data-flow, happy-path tests:
call the function, assert the return value on a couple of hand-picked inputs.
Those tests encode examples , not properties . They pass while the system
is deeply broken, and they lock in whatever behavior the code happened to have
on the day they were written. This skill replaces that reflex with a disciplined
thought-flow.
The core claim: software engineering is a decomposition-and-composition
problem. Correctness is not a list of examples; it is a set of properties
that hold over a state space. Tests are the proof, and a proof has structure.
Mental Model: Architecture Is a Vector Space
Every sufficiently complex domain decomposes into a small number of independent
axes of variation — the basis vectors of the domain. Each axis carries an
intrinsic property that the others cannot synthesize. This orthogonality is
not a design choice; it is a fact about the domain that you discover .
The math analogy makes it concrete. Consider three functions as "basis vectors":
No linear combination a·x + b·e^x — for any parameters — is periodic. The
property "periodic" lives on a basis vector that x and e^x simply do not
span. That is what orthogonality means : a property you cannot compose out of
the other components.
Software works the same way. A subsystem's correctness properties
("labels are never reused while live," "count is conserved across churn,"
"selection is independent of arrival order") live on specific axes. If you test
by feature, you sample points in the space and hope. If you test by property,
you prove the basis vectors and then prove their span.
Two tiers of test follow directly from this model:
A suite that has only the first tier proves the parts but not the whole. A suite
that has only the second proves the whole but localizes nothing when it breaks.
You want both. They are not redundant — they cover orthogonal risks.
Use it whenever the task is any of:
"Write tests for this module / function / subsystem."
"Design a test strategy" or "improve test coverage."
"Model the correctness of X" / "what are the invariants of X?"
"Review these tests" — audit whether they encode properties or examples.
Deciding what to test when the surface is large and examples feel arbitrary.
If you catch yourself about to write assert f(2, 3) == 5 , stop and run the
workflow. That assertion is a sampled point; find the property it is a sample of.
DECOMPOSE COMPOSE TEST THE BASIS TEST THE SPAN
(identify axes) → (state the algebra) → (prove each vector) → (prove the span)
Phase A Phase B Phase C Phase D
Phase A — Decompose (System Identification)
Find the basis. This is the hard 10% and it is a human/LLM-reasoning step, not
a coding step.
Name the boundary. State the subsystem as a map: what are its inputs, its
outputs, and the two worlds it bridges. Write one sentence:
"X turns ⟨inputs⟩ into ⟨outputs⟩, maintaining ⟨state⟩."
Enumerate behaviors. List everything the subsystem does. Don't organize
yet — just enumerate. Each behavior is a data point.
Find the axes of variation. For each pair of behaviors ask: "If I changed
A, would B have to change?" If no, A and B lie on different axes. An axis is
an independent direction of change.
- A change along axis 1 does not require a change along axis 2.
- A property on axis 1 can be stated and verified without reference to axis 2.
Reduce to the “eigen-basis”. Find the minimal, spanning, independent set of
properties that generate all behaviors:
- Spanning — every behavior is a composition of basis properties.
- Independent — no basis property is derivable from the others.
- Minimal — remove any one and some behavior can no longer be expressed.
Write the property for each axis. For every basis vector, state the
invariant or law that characterizes it in isolation. This sentence becomes a
basis test in Phase C. Use the taxonomy below to name it precisely.
Output of Phase A: a short list of axes, each with a one-line formal property.
Phase B — Compose (State the Algebra)
Express the system as composition, and write down the state machine you will
later check.
Feature = composition. For each externally visible behavior, write it as
behavior = compose(prop_axis_i, prop_axis_j, ...) . Shared properties across
many behaviors are shared components (this is also your architecture: orthogonal
axes → modules with no cross-imports, layered as a DAG).
Define the state machine you will model-check in Phase D:
- VARIABLES — the minimal state the invariants talk about.
- Init — the clean starting state (every test behavior begins here).
- Next — a disjunction of actions : the finite alphabet of operations
that drive transitions ( create , remove , reconfigure , crash , …).
- Invariants — the properties from Phase A, now as predicates over
VARIABLES that must hold in every reachable state .
Output of Phase B: the action alphabet + the invariant predicates.
Phase C — Test the Basis (Prove Each Vector)
For each axis, write a test that proves its intrinsic property in isolation .
Exercise the real component, mock only across the boundary. Replace only
what lives on the other side of the identified boundary (persistence,
downstream subsystem, clock). Run the real production code path. Shrinking a
resource (a 128K table → 256 entries) to make the test fast is fine as long as
the same production code runs .
Assert the property, not an output. Walk the whole structure and check the
invariant on every element, not just the return value. (e.g. "every node on the
USED list also carries the IN-TREE flag" — checked across all nodes.)
Choose white-box or black-box per property. Structural invariants
(list/tree integrity, conservation) need white-box state inspection; behavioral
invariants (output selection) can be black-box on the emitted messages. Pick
what the property requires.
One property per test. If a test asserts two unrelated properties, it is
sampling two axes at once and will be hard to localize.
Output of Phase C: one focused test per basis property.
Phase D — Test the Span (TLA+-Style)
Prove the composition. This is bounded model checking, done in whatever language or property testing libraries that fits the most. The fundamentals contains:
Enumerate the reachable state space. Generate all interleavings of the
action alphabet up to a bounded length (e.g. all sequences of length 1..3, plus
removals, reconfigurations, and re-additions). This is TLC's job, hand-rolled:
explore reachable states by applying every action in every order.
Assert the composed invariant in every reachable state — after each
action, not only at the end.
Inject faults as first-class actions. crash / restart is just another
action in the alphabet. After it, assert the recovery invariant (state restored,
nothing lost, nothing duplicated).
Check the span is not too big. Assert that illegal states are
unreachable — the "no periodicity from monotone bases" check. If an action
sequence reaches a state the invariant forbids, either the code or the model is
wrong.
Output of Phase D: a permutation harness + invariant checks (+ optional baseline).
Naming the property precisely is most of the battle — the right word tells you
what test to write. Reach for these before writing any assertion:
If you can't name the property, you have not finished Phase A. Go back.
Basis test (Phase C) — prove one vector in isolation
test_<axis>_<property>():
# ARRANGE: real component, mock ONLY across the boundary
init_real_component(size=SMALL) # shrink resource, same prod code
mock_the_thing_on_the_other_side() # persistence / downstream / clock
# ACT: drive the single axis
... apply operations that exercise ONLY this property ...
# ASSERT: the intrinsic invariant, over the WHOLE structure
for element in walk(structure):
assert invariant_holds(element) # not just the return value
assert conserved_quantity == expected # e.g. used + free == TOTAL
Span harness (Phase D) — prove the composition over reachable states
ALPHABET = [create_a, create_b, remove_a, remove_b, reconfigure, crash]
for seq in bounded_interleavings(ALPHABET, max_len=3): # + removals/flaps
state = Init()
for action in seq:
state = apply(action, state)
assert composed_invariant(state) # checked after EVERY action
TLA+ sketch (Phase B) — the model the harness approximates
VARIABLES s \* the minimal state the invariants talk about
Init == s = <clean initial state>
Next == \/ CreateA(s) \/ CreateB(s)
\/ RemoveA(s) \/ RemoveB(s)
\/ Reconfigure(s) \/ Crash(s) \* the action alphabet
Inv == /\ Conserved(s) \* count invariant
/\ FlagsConsistent(s) \* structural invariant
/\ NoDangling(s) \* referential integrity
\* Spec == Init /\ [][Next]_s ; check Inv is an invariant of Spec
You do not have to run TLC to benefit. Writing the four parts —
VARIABLES / Init / Next / Inv — is the design act. The permutation harness in
Phase D is the executable approximation of Init /\ [][Next]_s ⇒ []Inv .
Worked Example (Abstract — a Bounded Reusa

[truncated]
