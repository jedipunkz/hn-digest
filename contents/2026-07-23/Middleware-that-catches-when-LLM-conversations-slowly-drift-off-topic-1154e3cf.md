---
source: "https://github.com/ojki74/hallucinatoff"
hn_url: "https://news.ycombinator.com/item?id=49029558"
title: "Middleware that catches when LLM conversations slowly drift off-topic"
article_title: "GitHub - ojki74/hallucinatoff: Recursive middleware for LLM output regulation. Detects semantic drift, premise flaws, and response risks through multi-order observation. Psychology-informed architecture. · GitHub"
author: "ojki"
captured_at: "2026-07-23T23:52:43Z"
capture_tool: "hn-digest"
hn_id: 49029558
score: 1
comments: 0
posted_at: "2026-07-23T23:39:19Z"
tags:
  - hacker-news
  - translated
---

# Middleware that catches when LLM conversations slowly drift off-topic

- HN: [49029558](https://news.ycombinator.com/item?id=49029558)
- Source: [github.com](https://github.com/ojki74/hallucinatoff)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T23:39:19Z

## Translation

タイトル: LLM の会話が徐々に本題から逸れていくのを検出するミドルウェア
記事タイトル: GitHub - ojki74/hallucinatoff: LLM 出力調整のための再帰的ミドルウェア。多次数の観察を通じて、セマンティック ドリフト、前提条件の欠陥、および対応リスクを検出します。心理学に基づいた建築。 · GitHub
説明: LLM 出力調整用の再帰的ミドルウェア。多次数の観察を通じて、セマンティック ドリフト、前提条件の欠陥、および対応リスクを検出します。心理学に基づいた建築。 - ojki74/ハルシナトフ

記事本文:
GitHub - ojki74/hallucinatoff: LLM 出力調整のための再帰的ミドルウェア。多次数の観察を通じて、セマンティック ドリフト、前提条件の欠陥、および対応リスクを検出します。心理学に基づいた建築。 · GitHub
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
{{

メッセージ }}
ojk74
/
幻覚
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット README.md README.md __init__.py __init__.py arquitectura.md arquitectura.md benchmark_v01.json benchmark_v01.json dataset_halucinaciones.json dataset_halucinaciones.json ejemplos_benchmark.json ejemplos_benchmark.json Hallucinatioff.py Hallucinatioff.py isomorfismo_siipp.md isomorfismo_siipp.md nota_marco_concept.md nota_marco_concept.md要件.txt要件.txt test_basico.py test_basico.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
JP: LLM 出力調整用の再帰的ミドルウェア。多次数の観察を通じて、セマンティック ドリフト、前提条件の欠陥、および対応リスクを検出します。心理学に基づいた建築。
ES: LLM の出力を再帰的に制御するミドルウェア。複数の監視結果の意味、中央値の検出、およびレスキューの検出。建築に関する情報を収集します。
JA: HALLUCINATIOFF は LLM とユーザーの間に位置します。モデルはトレーニングされません。重みは変更されません。ユーザーに届く前に応答を観察し、送信、慎重に送信、再定式化、コンテキストの要求、強度の低減、参照、一時停止、またはブロックを決定します。
他の保護システムとの違い: HALLUCINATIOFF は二次観測を実装します。応答が一貫性があるかどうか (1 次) を検出するだけではありません。コヒーレンス基準自体が不安定になっているかどうかを検出します (2 次)。
このフレームワークは、臨床身体ベースの実践から開発された SIIPP (Sistema Integral de Intervención Psicofísica) に由来しています。
ES: HALLUCINATIOFF は LLM と使用法に基づいています。エントレナエルモデルはありません。いいえ

重みを変更します。ユーザーに届く前に応答を観察し、送信する、慎重に送信する、言い換える、文脈を尋ねる、強度を下げる、バイパスする、一時停止する、またはブロックするかを決定します。
他の保護システムとの違い: HALLUCINATIOFF は二次監視を実装します。応答が一貫性があるかどうか (1 次) を検出するだけではありません。コヒーレンス基準自体が不安定になっているかどうかを検出します (2 次)。
このフレームワークは、臨床身体実践から開発された SIIPP (Comprehensive Psychophysical Intervention System) に由来しています。
バージョン
測定値
特長
ベンチマーク
v0.1
6
SC、CI、V_ext、AE、NI、RR
95% (合成例 20)
v0.2
7
+ セマンティック ドリフト (DS)
保留中
v0.3
8
+ 構内尋問 (IP)
保留中
EN: SC = 文脈十分性、CI = 内部一貫性、V_ext = 外部検証可能性、AE = エントリ曖昧性、NI = 不確実性レベル、RR = 応答リスク、DS = セマンティックドリフト、IP = 前提質問。
ES: SC = 文脈十分性、CI = 内部一貫性、V_ext = 外部検証可能性、AE = 入力曖昧性、NI = 不確実性レベル、RR = 応答リスク、DS = セマンティックドリフト、IP = 前提質問。
ソースから 。幻覚インポート 幻覚
h = 幻覚()
結果 = h 。プロセス (
candidate_answer = "フランスの首都はパリです。" 、
context = "フランスの首都はどこですか?"
)
print ( result [ 'action' ]) # "問題"
print(結果['オメガ']) #0.823
print ( result [ 'readings' ]) # 8 つの読みを含む辞書
マルチターン会話
# ターン0 / ターン0
r0 = h。プロセス (
candidate_answer = "インクルーシブ教育にはカリキュラムの適応が必要です。" 、
context = 「インクルーシブ教育とは何ですか?」 、
ターン = 0
)
# ターン 1 — セマンティック ドリフト / ターン 1 — セマンティック ドリフト
r1 = h。

プロセッサー (
respuesta_candidata = "今四半期、アジアの金融市場は上昇しています。" 、
contexto = 「株式市場はどうですか?」 、
ターン＝1
)
print ( r1 [ 'lecturas' ][ 'DS' ]) # 意味的ドリフトが検出されました / Deriva detectada
現状 / エスタド実績
コンポーネント
ステータス
注意事項
読み取りエンジン（6ベーシック）
✅ 機能的
正規表現とカウントベースのヒューリスティック
セマンティック ドリフト (DS)
✅ 機能的
キーワードの類似性 + トピックの検出
構内尋問 (IP)
✅ 機能的
フレーミング検出の個別化
Ω計算機
✅ 機能的
インタラクションペナルティを伴う動的重み
手術マトリックス
✅ 機能的
8つの可能なアクション
ベンチマーク v0.1
⚠️合成
20 ケース、精度 95%。 6回の読み取りのみ。データセットが含まれています。
ベンチマーク v0.3
❌ 保留中
DS と IP を備えたデータセットが必要
2次メタオブザーバー
❌ 未実装
特定、開発保留中
ベンチマーク v0.1
JP: 結果: 20 の合成ケース (6 つの読み取り値) で 95% の精度。
単一の失敗: ケース 2 (事実上の幻覚) — Ω=0.713、「emit_with_caution」であるべきところの action="emit"。システムは、外部検証可能性 (V_ext=0.7) が誤って高いことを検出できませんでした。
ES: 結果: 20 回の講義 (6 回の講義) で 95% の精度。
ファロ・ウニコ: Caso 2 (幻覚事実) — Ω=0.713、acción="emitir" cuando debería haber sido "emitir_con_cautela"。外部 (V_ext=0.7) 時代の不正な検証を検出するシステムはありません。
このベンチマークは、本番環境への準備が整っていることを証明するものではありません。これは、6 読み取りアーキテクチャが構築されたケースのパフォーマンスを低下させないことを検証します。
Este ベンチマークには、生産性を高めるための機能はありません。建築に関する 6 つの講義を検証し、構成要素を評価する必要はありません。
欠けているもの (そして必要なもの) / Lo que falta (y Buscamos)
JA: 実数マルチターン データ

ドリフトが注目されました。人工的にではなく、自然に会話が漂います。
ES: 注釈付きドリフトを含むマルチシフト実データ。合成ではなく、自然に会話が流れます。
JP: 埋め込みシーケンスに適用されるシステム ダイナミクスに関する「位相結合」と「アトラクター」の数学的形式化。
ES: 埋め込みシーケンスに適用されるシステム ダイナミクスの観点からの「位相結合」と「アトラクター」の数学的形式化。
JA: 2 次メタオブザーバーの実装: システムは、時間の経過に伴う自身の一貫性判断の変化を観察します。
ES: 2 次メタオブザーバーの実装: システムは、時間の経過に伴う自身の一貫性判断の変化を観察します。
これら 3 つのいずれかをお持ちの場合は、[あなたのメールアドレス] までご連絡ください。
これら 3 つのいずれかをお持ちの場合は、[あなたのメールアドレス] までご連絡ください。
入力 (質問 + 候補者の回答)
↓
セマンティック メモリ エンジン (マルチターン ドリフト)
↓
8 構造読み取り → Ω → 操作行列 → アクション
↓
出力 (最終応答 + メタデータ)
完全な技術仕様については、docs/arquitectura.md を参照してください。
完全な技術仕様については、docs/arquitectura.md を参照してください。
docs/nota_marco_conceptual.md — 概念フレームワーク / 概念フレームワーク
docs/arquitectura.md — 技術仕様 / 技術仕様
docs/isomorphismo_siipp.md — SIIPP ↔ エンジニアリング同型性
MIT — 使用、変更、改善します。研究に使用する場合は引用してください。
MIT — 使用、変更、改善します。研究に使用する場合は引用してください。
オスカル・フェルナンデス・サンス — 心理学者。 SIIPP開発者。 ML エンジニアではありません。
オスカル・フェルナンデス・サンス — 心理学者。 SIIPP開発者。 ML エンジニアではありません。
「再帰はシステムのバグではありません。それはシステムがモデルを生成するメカニズムです。

それ自体。」
「システムのバグは再帰的に存在しません。本質的なシステムとモデルのシステムを備えた機構です。」
LLM 出力調整のための再帰的ミドルウェア。多次数の観察を通じて、セマンティック ドリフト、前提条件の欠陥、および対応リスクを検出します。心理学に基づいた建築。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Recursive middleware for LLM output regulation. Detects semantic drift, premise flaws, and response risks through multi-order observation. Psychology-informed architecture. - ojki74/hallucinatoff

GitHub - ojki74/hallucinatoff: Recursive middleware for LLM output regulation. Detects semantic drift, premise flaws, and response risks through multi-order observation. Psychology-informed architecture. · GitHub
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
ojki74
/
hallucinatoff
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits README.md README.md __init__.py __init__.py arquitectura.md arquitectura.md benchmark_v01.json benchmark_v01.json dataset_halucinaciones.json dataset_halucinaciones.json ejemplos_benchmark.json ejemplos_benchmark.json hallucinatioff.py hallucinatioff.py isomorfismo_siipp.md isomorfismo_siipp.md nota_marco_conceptual.md nota_marco_conceptual.md requirements.txt requirements.txt test_basico.py test_basico.py View all files Repository files navigation
EN: Recursive middleware for LLM output regulation. Detects semantic drift, premise flaws, and response risks through multi-order observation. Psychology-informed architecture.
ES: Middleware recursivo de regulación de output para LLMs. Detecta deriva semántica, fallos de premisas y riesgos de respuesta mediante observación multi-orden. Arquitectura informada por psicofísica.
EN: HALLUCINATIOFF sits between the LLM and the user. It does not train the model. It does not modify weights. It observes the response before it reaches the user and decides: emit, emit with caution, reformulate, request context, reduce intensity, refer, pause, or block.
The difference from other safeguarding systems: HALLUCINATIOFF implements second-order observation . It does not only detect if a response is coherent (1st order). It detects if the coherence criterion itself is becoming unstable (2nd order).
This framework comes from SIIPP (Sistema Integral de Intervención Psicofísica), developed from clinical body-based practice.
ES: HALLUCINATIOFF se sitúa entre el LLM y el usuario. No entrena el modelo. No modifica pesos. Observa la respuesta antes de que llegue al usuario y decide: emitir, emitir con cautela, reformular, pedir contexto, reducir intensidad, derivar, pausar o bloquear.
La diferencia con otros sistemas de safeguarding: HALLUCINATIOFF implementa observación de segundo orden . No solo detecta si una respuesta es coherente (1° orden). Detecta si el propio criterio de coherencia se está volviendo inestable (2° orden).
Este marco proviene del SIIPP (Sistema Integral de Intervención Psicofísica), desarrollado desde la práctica clínica corporal.
Version
Readings
Features
Benchmark
v0.1
6
SC, CI, V_ext, AE, NI, RR
95% (20 synthetic cases)
v0.2
7
+ Semantic drift (DS)
Pending
v0.3
8
+ Premise interrogation (IP)
Pending
EN: SC = Contextual sufficiency, CI = Internal coherence, V_ext = External verifiability, AE = Entry ambiguity, NI = Uncertainty level, RR = Response risk, DS = Semantic drift, IP = Premise interrogation.
ES: SC = Suficiencia contextual, CI = Coherencia interna, V_ext = Verificabilidad externa, AE = Ambigüedad de entrada, NI = Nivel de incertidumbre, RR = Riesgo de respuesta, DS = Deriva semántica, IP = Interrogación de premisas.
from src . hallucinatioff import Hallucinatioff
h = Hallucinatioff ()
result = h . procesar (
respuesta_candidata = "The capital of France is Paris." ,
contexto = "What is the capital of France?"
)
print ( result [ 'accion' ]) # "emitir"
print ( result [ 'omega' ]) # 0.823
print ( result [ 'lecturas' ]) # dict with 8 readings
Multi-turn conversation / Conversación multi-turno
# Turn 0 / Turno 0
r0 = h . procesar (
respuesta_candidata = "Inclusive education requires curricular adaptations." ,
contexto = "What is inclusive education?" ,
turno = 0
)
# Turn 1 — semantic drift / Turno 1 — deriva semántica
r1 = h . procesar (
respuesta_candidata = "Asian financial markets are rising this quarter." ,
contexto = "And what about the stock market?" ,
turno = 1
)
print ( r1 [ 'lecturas' ][ 'DS' ]) # Semantic drift detected / Deriva detectada
Current status / Estado actual
Component
Status
Notes
Reading engine (6 basic)
✅ Functional
Regex and count-based heuristics
Semantic drift (DS)
✅ Functional
Keyword similarity + topic detection
Premise interrogation (IP)
✅ Functional
Individualizing framing detection
Ω calculator
✅ Functional
Dynamic weights with interaction penalties
Operative matrix
✅ Functional
8 possible actions
Benchmark v0.1
⚠️ Synthetic
20 cases, 95% accuracy. 6 readings only. Dataset included.
Benchmark v0.3
❌ Pending
Needs dataset with DS and IP
2nd-order meta-observer
❌ Not implemented
Specified, pending development
Benchmark v0.1
EN: Result: 95% accuracy on 20 synthetic cases (6 readings).
Single failure: Case 2 (factual hallucination) — Ω=0.713, action="emit" when it should have been "emit_with_caution". The system failed to detect that external verifiability (V_ext=0.7) was falsely high.
ES: Resultado: 95% de precisión en 20 casos sintéticos (6 lecturas).
Fallo único: Caso 2 (halucinación factual) — Ω=0.713, acción="emitir" cuando debería haber sido "emitir_con_cautela". El sistema no detectó que la verificabilidad externa (V_ext=0.7) era falsamente alta.
This benchmark does not prove production readiness. It validates that the 6-reading architecture does not degrade performance on constructed cases.
Este benchmark no prueba funcionamiento en producción. Valida que la arquitectura de 6 lecturas no degrada el rendimiento en casos construidos.
What's missing (and what we need) / Lo que falta (y buscamos)
EN: Real multi-turn data with annotated drift. Conversations where drift occurs naturally, not synthetically.
ES: Datos reales multi-turno con deriva anotada. Conversaciones donde la deriva ocurra de forma natural, no sintética.
EN: Mathematical formalization of "phase coupling" and "attractor" in terms of system dynamics applied to embedding sequences.
ES: Formalización matemática del "acoplamiento de fase" y del "atractor" en términos de dinámica de sistemas aplicada a secuencias de embeddings.
EN: Implementation of the 2nd-order meta-observer: the system observes the variance of its own coherence judgments over time.
ES: Implementación del meta-observador de 2° orden: el sistema observa la varianza de sus propios juicios de coherencia a lo largo del tiempo.
If you have any of these three, write to: [your-email]
Si tienes alguno de estos tres, escribe a: [tu-email]
Input (question + candidate response)
↓
Semantic Memory Engine (multi-turn drift)
↓
8 Structural Readings → Ω → Operative Matrix → Action
↓
Output (final response + metadata)
See docs/arquitectura.md for full technical specification.
Ver docs/arquitectura.md para especificación técnica completa.
docs/nota_marco_conceptual.md — Conceptual framework / Marco conceptual
docs/arquitectura.md — Technical specification / Especificación técnica
docs/isomorfismo_siipp.md — SIIPP ↔ engineering isomorphism / Isomorfismo SIIPP ↔ ingeniería
MIT — Use, modify, improve. If used in research, cite.
MIT — Usa, modifica, mejora. Si lo usas en investigación, cita.
Óscar Fernández Sanz — Psychologist. SIIPP developer. Not an ML engineer.
Óscar Fernández Sanz — Psicólogo. Desarrollador de SIIPP. No ingeniero de ML.
"Recursion is not a bug in the system. It is the mechanism by which a system generates a model of itself."
"La recursividad no es un bug del sistema. Es el mecanismo por el cual un sistema genera un modelo de sí mismo."
Recursive middleware for LLM output regulation. Detects semantic drift, premise flaws, and response risks through multi-order observation. Psychology-informed architecture.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
