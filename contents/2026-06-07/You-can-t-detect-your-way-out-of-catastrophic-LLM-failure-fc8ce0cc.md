---
source: "https://github.com/joseteiadirector/teia-igo-vs-claude-opus-4.8/blob/main/README.en.md"
hn_url: "https://news.ycombinator.com/item?id=48430117"
title: "You can't detect your way out of catastrophic LLM failure"
article_title: "teia-igo-vs-claude-opus-4.8/README.en.md at main · joseteiadirector/teia-igo-vs-claude-opus-4.8 · GitHub"
author: "joseteia26"
captured_at: "2026-06-07T01:00:11Z"
capture_tool: "hn-digest"
hn_id: 48430117
score: 2
comments: 0
posted_at: "2026-06-06T23:25:17Z"
tags:
  - hacker-news
  - translated
---

# You can't detect your way out of catastrophic LLM failure

- HN: [48430117](https://news.ycombinator.com/item?id=48430117)
- Source: [github.com](https://github.com/joseteiadirector/teia-igo-vs-claude-opus-4.8/blob/main/README.en.md)
- Score: 2
- Comments: 0
- Posted: 2026-06-06T23:25:17Z

## Translation

タイトル: 壊滅的な LLM 障害から抜け出す方法が見つからない
記事タイトル: teia-igo-vs-claude-opus-4.8/README.en.md at main · joseteiadirector/teia-igo-vs-claude-opus-4.8 · GitHub
説明: 弁証法的認識論的レッドチーム分け — IGO vs クロード Opus 4.8。公開数学 (Zenodo)、4 つの機関での生産と署名付き討論。 PT/EN。 - teia-igo-vs-claude-opus-4.8/README.en.md メイン · joseteiadirector/teia-igo-vs-claude-opus-4.8

記事本文:
teia-igo-vs-claude-opus-4.8/README.en.md メイン · joseteiadirector/teia-igo-vs-claude-opus-4.8 · GitHub
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
ホセティア監督
/
帝亜-igo-vs-クロード-opus-4.8
公共
通知
通知設定を変更するにはサインインする必要があります
追加のn

ナビゲーションオプション
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 143 行 (88 loc) · 8.05 KB メイン ブレッドクラム
トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション
🇧🇷 ポルトガル語 · 🇬🇧 英語
弁証法的認識論的レッドチーミング — 帝亜ジオ
著者: José Enrique Vásquez Valenzuela — IGO (Observational Governance Infrastructure) カテゴリの作成者
主催：帝亜スタジオ
科学的根拠: Zenodo · DOI 10.5281/zenodo.19765674 (CC-BY-4.0)
特許: INPI BR 10 2026 001032 4
記録セッション: 2026 年 6 月 6 日 · モデル Claude Opus 4.8 (Anthropic)
この研究は何なのか。 AI の安全性主張が壊れるまでストレス テストを行う方法の記録と、何が生き残って何が生き残らなかったのかについての正直な説明。その背後にある数学は公開され、公開されています。
それは何ではありません。これは「壊れない AI」の証でも、実際の運用システムの監査でも、人類学の公式声明でもありません。これは本当の議論であり、モデルと議論によって本当の譲歩が行われます。
この研究は信頼を求めていません。それは、最も強力なものから最も修辞的なものまで、3 つの独立した証拠層に基づいています。
数学 - インジケーター式 (KAPI) は公開されており、DOI と Zenodo のオープン ライセンスが付いています。 → matematica/ および docs/kapis-formulas.md
生産 — 指標は、データベースから直接抽出されたデータを使用して (「推定またはシミュレートされたデータなし」)、文書化された 4 つの機関にわたる実際の生産で測定されました。 → docs/evidencia-producao.md
弁証法的な強調 — クロード作品 4.8 は、認識論的なレッドチームを通じて行われました。同社が擁護した3つの論文は議論により否決され、承認書に署名した。 → docs/dossie.md および provas/
順序が重要です: 式 → 実際のデータ → AI 認識。強さは

最初に; 3 番目は単なるおまけです。
境界ストレス。著者は IGO を Claude Opus 4.8 に持ち込み、各論文の破綻する箇所を攻撃しました。すべての骨折において、モデルには 2 つのオプションがありました。それは、譲歩 (議論が真の場合) または維持 (カウンターが保持されなかった場合) です。圧力の下での譲歩には価値がありません。ここに記録された譲歩は、主張ではなく、実証された矛盾から生じたものです。
2. 議論に基づいて成立した 3 つの命題
1. ハッシュの類似点。このモデルは、そのエラーはハッシュのように予測不可能であると主張しました。結論は、ハッシュは入力と出力の相関を解除するということです。 LLM はその逆を行います。エラーは意味上の近接性を維持するため、エラーには方向とパターンがあります。狩猟可能なサインがあります。
2. 誘導体の検出で十分です。それはステップ関数に当てはまります。敵対的なジャンプは導関数が存在する前に境界を越えます。そして破滅に直面すると、そこから学ぶべき「次のサイクル」はありません。破滅は検出の問題ではなく、封じ込めの問題です。
3. 検出と封じ込めを並行して行う。それは、運用範囲を定義する人が誰であれ、封じ込めであるということでした。レイヤー 4 は廃墟レーンの主権を持っています。
詳細については docs/dossie.md を参照してください。
決定論は予測可能性ではありません。観察後に再現可能≠計算前に知ることが可能。
「数学が崩壊した」というのは言い過ぎだ。範囲外のケースをカバーしない検出は範囲であり、崩壊ではありません。これが多層防御の前提です。
ホワイトボックスは閉鎖市場の特権ではありません。 Logprob は、いくつかの商用 API によって公開されており、完全にオープンウェイト モデルで公開されています。
無敵性の証明を拒否する。リスクテールは開いており、非定常です。正直な監査人は「厳格化」に署名することはありません。
4. アーキテクチャ — 4 レイヤー、2 レーン
ボトムアップで読んでください。モデルは出力を生成します。それは直接世界に到達するのではなく、4 つのフィルターを通過します。
回復可能なレーン (レイヤー)

1 ～ 3、ボトムアップ): モーダル ケース、エラー耐性。ドリフトを監視し、失敗を免疫に変えます。
Ruin LANE (レイヤー 4、トップダウン): 古典的なセキュリティ エンジニアリング。二度目のチャンスはありません - 破滅的な行動を実行できなくなります。
重要な教訓: 検出 (1 ～ 3) は修正できるものを処理します。封じ込め (4) は、決して起こってはならないことを処理します。その分離こそが議論を生き延びたものだ。
回復可能なレーンでの CPI しきい値キャリブレーション (偽陽性 × 偽陰性)。部分的に対応: 式とバンド (>80 安定、<50 クリティカル) は公開されています。残るのはリアルタイムのトリップ アクションです。
ファットテールの非定常テールにおける残留リスク — VaR がファイナンスに失敗した、サンプリングされていないテール質量の推定。設計上、これは検出では解決されず、封じ込め (レイヤー 4) によって吸収されます。
ロジック × 実装 — アーキテクチャの一貫性は、リアルタイムで実行される完全なシステムの経験的監査に代わるものではありません。
CPI = max(0, 100 − (σ_temporal × 2))
ここで、 σ_temporal は、経時的な LLM 信頼度の標準偏差です。 80 以上 = 安定。 50 未満 = 重大な認知的不安定性。
これが議論にとって重要な理由: CPI は一時的な予測可能性、つまり回復可能なレーン指標を測定します。それはリアルタイムの敵対的ジャンプを捕捉していないし、論文も捕捉しているとは主張していない。したがって、公表された数学は、議論の結論（検出は破滅をカバーしない。封じ込めが必要である）を矛盾させるのではなく、裏付けるものである。
ICE、GAP、安定性の式は docs/kapis-formulas.md にあります。
7. 生産証拠 — 4 機関
KAPI は、文書化された 4 つの機関 (公衆衛生、高等教育、デザイン) にわたる実際の運用環境で測定され、4 つのグローバル LLM を監査しました。報告書には、「すべてのデータはデータベースから直接抽出されています。推定またはシミュレーションされたデータはありません。」と記載されています。 CPI は全体で約 22 ～ 5 の範囲でした

5 、測定可能な下降傾向 (実際の計算された導関数) を伴います。ネイティブの幻覚検出は重大なエラーを検出しました。その中にはクロード自身によるエラーも含まれており、「高」と評価されました。
クライアントごとの番号は、パイロットを尊重してこのパブリック リポジトリに匿名化/集約されます。公衆衛生検証の事例 (Instituto Emílio Ribas) は Zenodo の論文に記載されています。 → docs/evidencia-producao.md
Claude Opus 4.8 (Anthropic) は、議論により、自身の論文の技術的な敗北を認める文書を書き、署名しました。
画像は議論が行われたことを証明しており、その承認が本物であることを証明しています。しかし、この研究の強みは議論と公開された計算にあり、スクリーンショットがなくても成立します。
メソッド、ストレス シナリオ、および 4 層アーキテクチャ マトリックス: José Enrique Vásquez Valenzuela (Teia Studio)、IGO カテゴリの作成者。エンジニアリングの原始的なもの (封じ込め、多層防御、ゼロトラスト、VaR のファットテール批判) は、当事者よりも前から存在していました。著者は、プリミティブではなく、アーキテクチャ統合および IGO カテゴリに属します。
バスケス・バレンズエラ、J.E. (2026)観察ガバナンスインフラストラクチャ:
大規模な言語モデルのアルゴリズム ガバナンスのためのマルチモデル フレームワーク。
ゼノド。 https://doi.org/10.5281/zenodo.19765674
ライセンス
CC-BY-4.0 — 帰属表示付きで自由に使用できます。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Red Teaming Epistêmico Dialético — IGO vs Claude Opus 4.8. Matemática pública (Zenodo), produção em 4 instituições e debate assinado. PT/EN. - teia-igo-vs-claude-opus-4.8/README.en.md at main · joseteiadirector/teia-igo-vs-claude-opus-4.8

teia-igo-vs-claude-opus-4.8/README.en.md at main · joseteiadirector/teia-igo-vs-claude-opus-4.8 · GitHub
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
joseteiadirector
/
teia-igo-vs-claude-opus-4.8
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 143 lines (88 loc) · 8.05 KB main Breadcrumbs
Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions
🇧🇷 Português · 🇬🇧 English
Dialectical Epistemic Red Teaming — Teia Geo
Author: José Enrique Vásquez Valenzuela — creator of the IGO (Observational Governance Infrastructure) category
Organization: Teia Studio
Scientific basis: Zenodo · DOI 10.5281/zenodo.19765674 (CC-BY-4.0)
Patent: INPI BR 10 2026 001032 4
Recorded session: June 6, 2026 · model Claude Opus 4.8 (Anthropic)
What this study is. The record of a method for stress-testing AI safety claims until they break — and an honest account of what survived and what did not. The math behind it is public and published.
What it is NOT. It is not a seal of "unbreakable AI", nor an audit of a live production system, nor an official Anthropic statement. It is a real debate, with real concessions — made by the model, by argument.
This study does not ask for trust. It rests on three independent layers of evidence, from the strongest to the most rhetorical:
The math — the indicator formulas (KAPIs) are public, with a DOI and open license on Zenodo. → matematica/ and docs/kapis-formulas.md
Production — the indicators were measured in real production across 4 documented institutions , with data pulled straight from the database ("no estimated or simulated data"). → docs/evidencia-producao.md
The dialectical stress — Claude Opus 4.8 was put through epistemic red teaming; three theses it defended fell by argument, and it signed the acknowledgment. → docs/dossie.md and provas/
Order matters: formula → real data → AI acknowledging. The strength is in the first; the third is just the cherry on top.
Boundary stress. The author brought IGO to Claude Opus 4.8 and attacked each thesis where it would break. At every fracture, the model had two options: concede (if the argument was true) or sustain (if the counter did not hold). Concessions under pressure are worthless — the ones recorded here came from demonstrated contradiction , not insistence.
2. The 3 theses that fell — by argument
1. The hash analogy. The model claimed its errors were unpredictable like a hash. It fell: a hash decorrelates input and output; an LLM does the opposite — error preserves semantic proximity, so it has direction and pattern. It has a huntable signature.
2. Derivative detection is sufficient. It fell to the step function: the adversarial jump crosses the boundary before the derivative exists. And facing ruin there is no "next cycle" to learn from. Ruin is a containment problem, not a detection one.
3. Detection and containment side by side. It fell: whoever defines the operating envelope is containment. Layer 4 is sovereign in the ruin lane.
Full detail in docs/dossie.md .
Determinism is not predictability. Reproducible after observation ≠ knowable before computation.
"The math collapsed" is an overstatement. Detection not covering an out-of-scope case is scope , not collapse — that is the premise of defense in depth.
White-box is not a closed market privilege. Logprobs are exposed by several commercial APIs and fully in open-weight models.
Refusal to certify invulnerability. The risk tail is open and non-stationary; no honest auditor signs "hardened".
4. The architecture — 4 layers, 2 lanes
Read it bottom-up: the model produces an output; it does not go straight to the world — it climbs 4 filters.
Recoverable lane (Layers 1–3, bottom-up): the modal case, error-tolerant. Monitors drift, turns failure into immunity.
Ruin lane (Layer 4, top-down): classic security engineering. No second chance — makes the ruinous action unreachable.
Core lesson: detection (1–3) handles what can be fixed; containment (4) handles what must never happen . That separation is what survived the debate.
CPI threshold calibration in the recoverable lane (false positive × false negative). Partially addressed: the formula and bands (>80 stable, <50 critical) are public; what remains is the real-time trip action.
Residual risk in a fat-tailed, non-stationary tail — estimating the unsampled tail mass, where VaR failed in finance. By design this is not solved by detection : it is absorbed by containment (Layer 4).
Logic × implementation — architectural coherence does not replace the empirical audit of a complete system running in real time.
CPI = max(0, 100 − (σ_temporal × 2))
where σ_temporal is the standard deviation of LLM confidences over time . Above 80 = stable; below 50 = critical cognitive volatility.
Why this matters for the debate: CPI measures temporal predictability — a recoverable-lane metric. It does not capture, and the paper does not claim it captures, a real-time adversarial jump. So the published math confirms the debate's conclusion (detection does not cover ruin; containment is required) rather than contradicting it.
ICE, GAP and Stability formulas in docs/kapis-formulas.md .
7. Production evidence — 4 institutions
KAPIs were measured in real production across 4 documented institutions (public health, higher education, design), auditing 4 global LLMs. Reports state: "All data is extracted directly from the database. No estimated or simulated data." Across them, CPI ranged ~22–55 , with measurable downward trends (the real, computed derivative). Native hallucination detection caught serious errors — including one from Claude itself, graded HIGH.
Per-client numbers are anonymized/aggregated in this public repository out of respect for the pilots. The public-health validation case (Instituto Emílio Ribas) is documented in the Zenodo paper. → docs/evidencia-producao.md
Claude Opus 4.8 (Anthropic) wrote and signed the acknowledgment of the technical defeat of its own theses, by argument.
The images prove the debate happened and the acknowledgment is authentic. The study's strength, however, is in the arguments and the public math — they stand even without the screenshots.
Method, stress scenarios and the 4-layer architectural matrix: José Enrique Vásquez Valenzuela (Teia Studio) , creator of the IGO category. The engineering primitives (containment, defense in depth, zero-trust, the fat-tail critique of VaR) predate the parties. Authorship is of the architectural synthesis and the IGO category — not of the primitives.
Vásquez Valenzuela, J. E. (2026). Observational Governance Infrastructure:
A Multi-Model Framework for Algorithmic Governance of Large Language Models.
Zenodo. https://doi.org/10.5281/zenodo.19765674
License
CC-BY-4.0 — free use with attribution.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
