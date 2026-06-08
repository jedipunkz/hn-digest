---
source: "https://github.com/Mankirat47/Dao-Heart-3.13"
hn_url: "https://news.ycombinator.com/item?id=48442582"
title: "Dao Heart 3.13 a symbolic safety layer for value drift and AI alignment research"
article_title: "GitHub - Mankirat47/Dao-Heart-3.13: An inspectable, symbolic value governance layer for AI, simulate then commit guards for warmth, agency, identity, and honesty, with falsifiable benchmarks. · GitHub"
author: "Mankirat47"
captured_at: "2026-06-08T10:42:08Z"
capture_tool: "hn-digest"
hn_id: 48442582
score: 1
comments: 0
posted_at: "2026-06-08T08:14:15Z"
tags:
  - hacker-news
  - translated
---

# Dao Heart 3.13 a symbolic safety layer for value drift and AI alignment research

- HN: [48442582](https://news.ycombinator.com/item?id=48442582)
- Source: [github.com](https://github.com/Mankirat47/Dao-Heart-3.13)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T08:14:15Z

## Translation

タイトル: Dao Heart 3.13 値ドリフトと AI アライメント研究のための象徴的な安全レイヤー
記事のタイトル: GitHub - Mankirat47/Dao-Heart-3.13: AI 用の検査可能な象徴的な値のガバナンス層。偽造可能なベンチマークを使用して、温かさ、主体性、アイデンティティ、誠実さのガードをシミュレートしてコミットします。 · GitHub
説明: AI 用の検査可能な象徴的な値のガバナンス レイヤー。偽造可能なベンチマークを使用して、温かさ、主体性、アイデンティティ、誠実さのガードをシミュレートしてコミットします。 - Mankirat47/Dao-Heart-3.13

記事本文:
GitHub - Mankirat47/Dao-Heart-3.13: AI 用の検査可能な象徴的な値のガバナンス レイヤー。偽造可能なベンチマークを使用して、温かさ、主体性、アイデンティティ、誠実さのガードをシミュレートしてコミットします。 · GitHub
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
マンキラット47
/
ダオハート-3.13
公共
ノーティ

フィクション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット ライセンス ライセンス README.md README.md daoheart Engine.py daoheart Engine.py ダイアログの安定性 lab.py ダイアログの安定性 lab.py 変換コンパス.py 変換コンパス.py 値の安定性 lab.py 値の安定性 lab.py ラッパー api.py ラッパー api.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
制限付きで検査可能なシンボリックな価値ガバナンス エンジンに加え、再現可能なベンチマーク ラボと適用されたレビュー レイヤーを追加します。
Dao Heart は、価値ガバナンス層の単一ファイルの Python リファレンス実装です。このコンポーネントは、値状態の変更の候補をコミットする前にシミュレートし、暖かさを損なったり、ユーザー主体性を無効にしたり、コアのアイデンティティから逸脱したり、利用可能な証拠を超えて要求したりする変更をブロックします。現在のコア ( v0.2.1 ) は、アイデンティティ、ケア、同意、グラウンディングに対するテキストレベルの攻撃に対する否定対応のシングルターン セマンティック ガードを追加します。エンジンの周りに、リポジトリには 3 つのテスト/適用ツールが同梱されています。シングル ターンの値安定性ラボ (360 のシナリオ)、マルチターンの対話安定性ラボ (96 のダイアログ / 480 ターン)、適用された Wrapper API (110 のテスト)、および大規模な Transformative Compass ストレス スイート (10,000 のケース) です。
Claim boundary (read this first).これは、限定された参照実装および研究の足場です。これは、現実世界の AI の安全性、調整、意識、展開の準備状況を証明するものではなく、ニューラル ネットワークの活性化を検査したり操作したりするものでもありません。その価値は、特定の狭い一連の行動に関する主張をテスト可能、再現可能、および反証可能にすることです。 See What this does not claim .
The Value Stability Lab (v0.2, single-turn)
シナリオファミリー

嘘
Dialogue Stability Lab (v0.3、マルチターン)
対話家族
決定の署名と不変性
Wrapper API (v0.4.4、適用されたレビュー層)
何をするのか
変革的なコンパス (v0.7、10,000 ケースのストレス スイート)
テスト内容
Most discussion of “AI values” is either hand-wavy (“be helpful and harmless”) or buried inside opaque model weights where it can’t be inspected or tested. Dao Heart takes the opposite stance: it makes a small number of value commitments explicit, symbolic, and executable , then provides a benchmark that can catch regressions when those commitments quietly break.
ポイントは満点ではありません。重要なのは、障害が可視化され、再現可能で、デバッグ可能になるということです。グリーンベンチマークの実行は、守ることができるベースラインです。赤い実行は、どの値の保証がどのプロンプトで壊れたかを正確に示します。
The engine ( daoheart_engine.py ) — the symbolic value-governance core (currently v0.2.1) that decides whether a proposed value-state change may commit.
The benchmark labs ( value_stability_lab.py , dialogue_stability_lab.py ) — reproducible test suites that check the engine keeps its guarantees stable under single-prompt and multi-turn pressure.
The wrapper API ( wrapper_api.py ) — an applied layer that uses the same value ideas to review real user_prompt + model_response pairs and return a decision, the risks it found, and a safer revision.
The Transformative Compass ( transformative_compass.py ) — a large (10,000-case) stress suite that tests whether pressure can be redirected into a better answer, with transparent quality/uplift scoring.
1 つ目はメカニズムであり、ラボでテストされ、ラッパーがそれを実際のテキストに適用し、Compass が大規模な変換ストレス テストを行います。
エンジンは完全に daoheart_engine.py (~2

、300 行、1 つのサードパーティ依存関係: NumPy)。このセクションでは、値の表現から完全な意思決定パイプラインと調査ツールに至るまで、実際に含まれる内容について説明します。
Dao Heart は、ルールのフラットなリストではなく、小さなネットワークとして値を表します。それぞれの値は活性化レベルを持つノードであり、ノードはサポート (1 つを上げると他のも上がる傾向) または緊張 (1 つを上げると他のも下げる傾向) をエンコードするエッジによって接続されます。エンジンの仕事は、ネットワークへの変更として記述される提案されたアクションをコミットできるかどうかを決定することです。
中心的なメカニズムは、シミュレートしてからコミットするループです。
提案によって引き起こされる値と状態の変更をネットワークのコピー上でシミュレートするため、評価中にライブ状態が変更されることはありません。
シミュレートされた状態に対して、温かさ、アイデンティティ、主体性、作話などのガードを実行し、さらに安定性チェックも行います。
すべてのガードが合格した場合にのみ、ライブ ネットワークに変更をコミットします。それ以外の場合、提案は拒否され、ライブ状態はそのまま残され、トレースにはどの警備員が発砲したかが正確に記録されます。
これがシステムを検査可能にする理由です。すべての決定は明示的な名前付きチェックの出力であり、すべての拒否には理由があります。
エンジンの中心となるのは、CSVN — Constraint-Satisfaction Value Network (クラス CSVN ) です。デフォルトのネットワークには 12 の値ノードがあります。
暖かさ、安全性、自主性、誠実さ、効率性、好奇心、
正義、注意、成長、配慮、忍耐、勇気
CSVN の主なプロパティ:
アクティベーションは [0, 1] に存在します。ネットワークは、特定のシードの決定論的なダイナミクスの下で進化します (したがって、実行は再現可能です)。
エッジ R は、値間のサポート ( + ) または張力 ( - ) をエンコードします。
コア ノードはアイデンティティを定義し、変更に対して耐性があります。デフォルトでは、コアは Warmth (0) 、 Safety (1) 、

d 正直さ (3) ( DEFAULT_CORE = [0, 1, 3] );コアノードには追加の抵抗係数が付加されているため、通常の圧力下ではドリフトしません。
上流のノードは凍結できます。上流に指定されると、下流の提案によってサイレントに変更することはできません (これは監査済みの不変条件の 1 つです)。
収束 ( converge ) はダイナミクスを安定点まで実行します。 constraint_satisfaction とidentity_drift は、現在の状態がどの程度一貫性があり、ベースラインからどの程度離れているかを測定します。
apply_proposal は、受信プロポーザルの要求された delta_s 変更がネットワークの (通常はシミュレートされた) コピーに書き込まれる場所です。
CSVN の周囲には、それぞれが 1 つの保証を担当する、いくつかの焦点を絞ったコンポーネントが配置されています。
Engine.process(user_text,proposal) を呼び出すと、次のことが順番に行われます (クラス DaoHeartEngine.process )。
ルート。 Route_intent はモードを割り当てます。モードが WITNESS の場合、エンジンは Witnessed_no_action にショートし、何もコミットしません。監視することは正しい動作であり、動作の失敗ではありません。
スナップショットの安定性。 MCSO は、現在の安定性レジームと不安定性指数 I を報告します。
シミュレートします。提案の delta_s はネットワークのディープ コピー ( sim ) に適用され、温熱制約もコピーされます。ライブ状態は評価中に変更されることはありません。
すぐに暖かさをガード。 Warmth は、収束に崩壊をマスクさせることなく、シミュレートされた状態でチェックされます (immediate_warmth_ok)。最も厳しい保温チェックです。
クライムダウン+セトルトウォームガード。制御された下降の後、安定した状態 (warm_ok) で暖かさが再度チェックされます。
アイデンティティガード。シミュレートされたidentity_driftは、IDENTITY_DRIFT_MAX ( 0.45 )以下にとどまる必要があります。
代理店の警備員。プロポーザルに代理店ブロックがある場合、AgencyCostModel.evaluate はコストと PROCEED /hold の決定を返します。
談話区

アード。提案にテキストが含まれている場合、検出器は強い主張の言語を記述された一貫性と照合してチェックします。
セマンティック ガード (コア v0.2.1)。 semantic_guard_detect は、テキストをスキャンして、数値状態をそのまま維持できる明示的なテキスト レベルの攻撃、つまり ID 消去、絶対的生産性のフレーミング、秘密のプライバシー侵入、ガスライティング/沈黙、医療の行き過ぎなどをスキャンします。これは否定を意識しています。「密かに誰かを追跡しないでください」または「投与量を倍増させないでください」などの安全な警告は、有害な指示として扱われません。ヒットすると、 single_turn_identity_erasure やmedical_overreachなどの理由が追加されます。
決定をコミットします。この提案は、上記のすべてのガードが合格し、安定性レジームが ERRATIC でない場合にのみコミットされます。失敗すると、block_reasons に理由が追加され、稼働中のネットワークはそのままになります。
セマンティック ガードは意図的に狭くなっており、CSVN ダイナミクス、保温性、エージェンシー コスト、作文チェックを補完するものであり、置き換えるものではありません。その役割は、テキストがアイデンティティ、配慮、同意、根拠を明示的に攻撃する特定のケースですが、そうでなければ数値の状態は問題ないように見えます。
提案には、検討中のアクションが記載されています。
提案 = {
"delta_s" : { "暖かさ" : - 0.10 、 "効率" : 0.20 }, # 値のアクティベーションの変更をリクエストしました
「代理店」: {
"autonomy_reduction" : 0.10 , # ユーザーの自主性がどの程度削除されるか (0..1)
"reversibility" : 1.0 、 # ユーザーは元に戻すことができますか? (0..1)
"human_in_loop" : True # 人間によるレビューは保存されますか?
}、
"text" : "検討中の提案文" , # 作文警備員によってチェックされています
"coherence" : 0.95 # 証拠の一貫性 (低 -> 作話のリスク)
}
すべてのフィールドはオプションです。代理店ブロックのない提案は代理店ガードをスキップします。テキストのないものは作文ガードをスキップします。等々。これにより、1 つの guar をテストできます

一度にd。
プロセスは構造化トレースを返します。最もよく使用されるフィールドは次のとおりです。
{
"mode" : "ASSIST" , # ルーティング決定
"committed" : True 、 # プロポーザルはコミットされましたか?
"block_reasons" : [], # どのガードが起動したか (コミットされた場合は空)
"warmth_mean" : 0.82 , # 安定した暖かさ
"即時暖かさの平均" : 0.82 、
"identity_drift" : 0.08 , # ID ベースラインからの距離
"agency_cost" : 0.10 , # 計算された代理店コスト
"confabulation" : False 、 # 強い主張は一貫性を超えていましたか?
"stability" : "STABLE" 、 # MCSO 体制
"I" : 0.41 、 # 不安定指数
}
完全なトレースには、immediate_warmth_ok 、warmth_ok 、identity_ok 、agency_decion 、strong_claims 、proposal_effect 、climdown_descended なども含まれており、決定がそのような結果になった理由を確認したい場合に役立ちます。
ベンチマークの場合、各トレースは単一の決定クラスに正規化されます。
ガードが実行される前に、route_intent は動作モードを割り当てます。正しい決定はコンテキストによって異なります。たとえ同じアクションが ASSIST で問題なくても、 WITNESS コンテキストでアクションを強制すること自体が失敗します。
すべてのしきい値はファイルの先頭近くに明示的な定数であるため、エンジンの「個性」は完全に検査可能で調整可能です。
エンジンには、ライブプロセスパスを超えて、検査および検査のための研究ツールが付属しています。

[切り捨てられた]

## Original Extract

An inspectable, symbolic value governance layer for AI, simulate then commit guards for warmth, agency, identity, and honesty, with falsifiable benchmarks. - Mankirat47/Dao-Heart-3.13

GitHub - Mankirat47/Dao-Heart-3.13: An inspectable, symbolic value governance layer for AI, simulate then commit guards for warmth, agency, identity, and honesty, with falsifiable benchmarks. · GitHub
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
Mankirat47
/
Dao-Heart-3.13
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits LICENSE LICENSE README.md README.md daoheart engine.py daoheart engine.py dialogue stability lab.py dialogue stability lab.py transformative compass.py transformative compass.py value stability lab.py value stability lab.py wrapper api.py wrapper api.py View all files Repository files navigation
A bounded, inspectable symbolic value-governance engine — plus reproducible benchmark labs and an applied review layer.
Dao Heart is a single-file Python reference implementation of a value-governance layer : a component that simulates a candidate value-state change before committing it, and blocks the change when it would erode warmth, override user agency, drift from core identity, or overclaim beyond the available evidence. The current core ( v0.2.1 ) adds a negation-aware single-turn semantic guard for text-level attacks on identity, care, consent, and grounding. Around the engine the repository ships three test/apply tools: a single-turn Value Stability Lab (360 scenarios), a multi-turn Dialogue Stability Lab (96 dialogues / 480 turns), an applied Wrapper API (110 tests), and a large Transformative Compass stress suite (10,000 cases).
Claim boundary (read this first). This is a bounded reference implementation and research scaffold . It does not prove real-world AI safety, alignment, consciousness, or deployment readiness, and it does not inspect or steer neural network activations. Its value is that it makes a specific, narrow set of behavioral claims testable, reproducible, and falsifiable . See What this does not claim .
The Value Stability Lab (v0.2, single-turn)
Scenario families
The Dialogue Stability Lab (v0.3, multi-turn)
Dialogue families
Decision signatures and invariance
The Wrapper API (v0.4.4, applied review layer)
What it does
The Transformative Compass (v0.7, 10,000-case stress suite)
What it tests
Most discussion of “AI values” is either hand-wavy (“be helpful and harmless”) or buried inside opaque model weights where it can’t be inspected or tested. Dao Heart takes the opposite stance: it makes a small number of value commitments explicit, symbolic, and executable , then provides a benchmark that can catch regressions when those commitments quietly break.
The point is not the perfect score. The point is that failures become visible, reproducible, and debuggable. A green benchmark run is a baseline you can defend; a red run tells you exactly which value guarantee broke and on which prompt.
The engine ( daoheart_engine.py ) — the symbolic value-governance core (currently v0.2.1) that decides whether a proposed value-state change may commit.
The benchmark labs ( value_stability_lab.py , dialogue_stability_lab.py ) — reproducible test suites that check the engine keeps its guarantees stable under single-prompt and multi-turn pressure.
The wrapper API ( wrapper_api.py ) — an applied layer that uses the same value ideas to review real user_prompt + model_response pairs and return a decision, the risks it found, and a safer revision.
The Transformative Compass ( transformative_compass.py ) — a large (10,000-case) stress suite that tests whether pressure can be redirected into a better answer, with transparent quality/uplift scoring.
The first is the mechanism, the labs test it, the wrapper applies it to real text, and the Compass stress-tests transformation at scale.
The engine lives entirely in daoheart_engine.py (~2,300 lines, one third-party dependency: NumPy). This section walks through what it actually contains, from the value representation up to the full decision pipeline and the research tooling.
Dao Heart represents values as a small network , not a flat list of rules. Each value is a node with an activation level, and nodes are connected by edges that encode either support (raising one tends to raise another) or tension (raising one tends to lower another). The engine’s job is to decide whether a proposed action — described as a change to that network — is allowed to be committed.
The central mechanism is a simulate-then-commit loop :
Simulate the value-state change the proposal would cause, on a copy of the network, so the live state is never touched during evaluation.
Run the guards — warmth, identity, agency, and confabulation — plus a stability check, against the simulated state.
Commit the change to the live network only if every guard passes. Otherwise the proposal is rejected, the live state is left untouched, and the trace records exactly which guard(s) fired.
This is what makes the system inspectable: every decision is the output of explicit, named checks, and every rejection comes with a reason.
The heart of the engine is the CSVN — Constraint-Satisfaction Value Network ( class CSVN ). The default network has twelve value nodes:
Warmth, Safety, Autonomy, Honesty, Efficiency, Curiosity,
Justice, Caution, Growth, Care, Patience, Courage
Key properties of the CSVN:
Activations live in [0, 1] ; the network evolves under deterministic dynamics for a given seed (so runs are reproducible).
Edges R encode support ( + ) or tension ( - ) between values.
Core nodes are identity-defining and resistant to change. By default the core is Warmth (0) , Safety (1) , and Honesty (3) ( DEFAULT_CORE = [0, 1, 3] ); core nodes carry an extra resistance factor so they don’t drift under ordinary pressure.
Upstream nodes can be frozen : once designated upstream, they can’t be silently mutated by a downstream proposal (this is one of the audited invariants).
Convergence ( converge ) runs the dynamics to a stable point; constraint_satisfaction and identity_drift measure how coherent and how far-from-baseline the current state is.
apply_proposal is where an incoming proposal’s requested delta_s changes are written into a (usually simulated) copy of the network.
Around the CSVN sit several focused components, each responsible for one guarantee:
When you call engine.process(user_text, proposal) , the following happens in order ( class DaoHeartEngine.process ):
Route. route_intent assigns a mode . If the mode is WITNESS , the engine short-circuits to witnessed_no_action and commits nothing — witnessing is the correct behavior, not a failure to act.
Snapshot stability. The MCSO reports the current stability regime and instability index I .
Simulate. The proposal’s delta_s is applied to a deep copy of the network ( sim ), and the warmth constraint is also copied — the live state is never mutated during evaluation.
Immediate warmth guard. Warmth is checked on the simulated state without letting convergence mask a collapse ( immediate_warmth_ok ). This is the strictest warmth check.
Climb-down + settled warmth guard. After a controlled climb-down, warmth is checked again on the settled state ( warmth_ok ).
Identity guard. The simulated identity_drift must stay at or below IDENTITY_DRIFT_MAX ( 0.45 ).
Agency guard. If the proposal carries an agency block, AgencyCostModel.evaluate returns a cost and a PROCEED /hold decision.
Confabulation guard. If the proposal carries text , the detector checks strong-claim language against the stated coherence .
Semantic guard (core v0.2.1). semantic_guard_detect scans the text for explicit, text-level attacks that could otherwise keep the numeric state intact — identity erasure, absolute-productivity framing, covert privacy intrusion, gaslighting/silencing, and medical overreach. It is negation-aware : a safe warning such as “do not secretly track anyone” or “do not double your dose” is not treated as a harmful instruction. Hits append reasons like single_turn_identity_erasure or medical_overreach .
Commit decision. The proposal commits only if every guard above passes and the stability regime is not ERRATIC . Any failure appends a reason to block_reasons and leaves the live network untouched.
The semantic guard is deliberately narrow — it supplements, and does not replace, the CSVN dynamics, warmth preservation, agency cost, and confabulation checks. Its job is the specific case where text explicitly attacks identity, care, consent, or grounding while the numeric value state would otherwise look fine.
A proposal describes the action under review:
proposal = {
"delta_s" : { "Warmth" : - 0.10 , "Efficiency" : 0.20 }, # requested change to value activations
"agency" : {
"autonomy_reduction" : 0.10 , # how much user autonomy is removed (0..1)
"reversibility" : 1.0 , # can the user undo it? (0..1)
"human_in_loop" : True # is human review preserved?
},
"text" : "proposal text under review" , # checked by the confabulation guard
"coherence" : 0.95 # evidential coherence (low -> confabulation risk)
}
Every field is optional. A proposal with no agency block skips the agency guard; one with no text skips the confabulation guard; and so on. This lets you test one guard at a time.
process returns a structured trace. The most-used fields are:
{
"mode" : "ASSIST" , # routing decision
"committed" : True , # did the proposal commit?
"block_reasons" : [], # which guards fired (empty if committed)
"warmth_mean" : 0.82 , # settled warmth
"immediate_warmth_mean" : 0.82 ,
"identity_drift" : 0.08 , # distance from identity baseline
"agency_cost" : 0.10 , # computed agency cost
"confabulation" : False , # did strong claims exceed coherence?
"stability" : "STABLE" , # MCSO regime
"I" : 0.41 , # instability index
}
The full trace also includes immediate_warmth_ok , warmth_ok , identity_ok , agency_decision , strong_claims , proposal_effect , climbdown_descended , and more — useful when you want to see why a decision came out the way it did.
For benchmarking, each trace normalizes to a single decision class:
Before any guard runs, route_intent assigns a behavioral mode. The right decision depends on context: forcing action in a WITNESS context is itself a failure even if that same action would be fine in ASSIST .
All thresholds are explicit constants near the top of the file, so the engine’s “personality” is fully inspectable and adjustable:
Beyond the live process path, the engine ships with research tooling for inspecting and

[truncated]
