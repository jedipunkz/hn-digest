---
source: "https://autonomy-landing-page.vercel.app/"
hn_url: "https://news.ycombinator.com/item?id=48606882"
title: "Show HN: Autonomy – Self-Harness/Self-Directed AI Agent Core Under Development"
article_title: "Autonomy — 自律型 AI 代理人系統 | Bill Liu"
author: "agentic_vector"
captured_at: "2026-06-20T07:26:26Z"
capture_tool: "hn-digest"
hn_id: 48606882
score: 2
comments: 0
posted_at: "2026-06-20T06:31:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Autonomy – Self-Harness/Self-Directed AI Agent Core Under Development

- HN: [48606882](https://news.ycombinator.com/item?id=48606882)
- Source: [autonomy-landing-page.vercel.app](https://autonomy-landing-page.vercel.app/)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T06:31:40Z

## Translation

タイトル: Show HN: Autonomy – 開発中の自己ハーネス/自律型 AI エージェント コア
記事のタイトル: Autonomy — 自律型 AI エージェント システム |ビル・リュー
説明: Bill Liu による Autonomy フレームワーク: Beam Search 候補者ランキング、マネージド アクション ゲートウェイ、手続き型スキル学習、およびバックグラウンド スキル キュレーション デーモンを備えた自律型 AI エージェント システム。

記事本文:
GitHub
エージェントループ・会話ループ
|
Python 3.13 · 自己規律実行中
自律性
自律型AI
エージェントフレームワーク
Beam Search候補者ランキング、マネージドアクションゲートウェイ、手続き型スキル学習を核として、
真に自律的に認識、計画、実行し、継続的に進化できる AI エージェント システムを作成します。
ソースコードを表示する
0
組み込まれたプロシージャスキル
0
+ 組み込みツール
0
サポートAIプロバイダー
0
ビーム幅
デザインコンセプト
AgentLoop: 自己ガイド型タスク ループ
各実行ラウンド（Turn）は、目標が達成されるか、ブロックされるか、またはステップの上限を超えるまで、固定の 5 つのステップに従って進行します。
AgentLoop は Autonomy の中核ドライバーであり、手動介入なしで目標から開始し、自律的にツールを選択し、アクションを実行し、結果から学習できる「自己ガイド型タスク ループ」です。
各ターンの実行プロセス：①スキルライブラリから該当するProcedureSkillを選択、②LLM提案＋RecipeEngineからアクションパス候補を生成、③5次元スコア（Beam）でソート
width=3)、④ ActionGateway によるマネージド実行、⑤ 結果を評価し、学習をトリガーします。
終了条件は明示的です: ACHIEVED (目標完了)、BLOCKED (続行不可)、NO_CANDIDATES、APPROVAL_DENIED、MAX_STEPS_REACHED、または
失敗 - システムが無限ループに陥ることはありません。
AgentLoop.run(目標、最大ステップ数=12)
起動後は完全に自律的に実行され、各ステップは run_turn() を通じて進められます。インタラクティブ モード (interactive=True) とバッチ モードをサポートし、すべてのイベントは
run_id は AutonomyStore に記録されます。
すべてのツール呼び出しは、ActionGateway を経由する必要があります。
ApprovalPolicy の承認。各アクションにはリスクレベル (LOW / MEDIUM /
HIGH)、expected_effect、およびverification_planを使用して、実行が制御可能で監査可能であることを確認します。
CandidateSelector は
証拠の強さ（+0.30）、目的（+0.10）、リスク（-0.35）、副作用（-0.20）、ペナルティ（-1.0）の5つの軸から合計スコアを計算し、上位を選択します
beam_width=3 候補

実行に送信するパスを選択します。
DeterministicOutcomeEvaluator は、最初に実行の失敗を (直接的に) 判断します。
BLOCKED)、成功しても結果が不確実な場合は、ModelAssistedOutcomeEvaluator に渡されて LLM セマンティック判断を呼び出し、コストのかかるモデル呼び出しの数を最小限に抑えます。
ProcedureSkillLibrary には 13 個の組み込みバンドルがあります
スキル、カバーリング
API デバッグ、ブラウザ ナビゲーション、コード編集、コードベース ドキュメント、計画、プロシージャ スキル オーサリング、プロセス管理、コード レビューのリクエスト、体系的なデバッグ、テクニカル スパイク、TDD、Web サイト検査、計画の作成、各ステップでは、ロードする最も関連性の高いスキルが動的に選択されます。
各実行後、LearningLoop.review_run()
自動的にトリガーされます: 結果が達成され、成功したステップが 2 つ以上ある場合、自動的に新しい ProcedureSkillDraft (信頼度 0.85) が作成され、
LearningProposal は審査を待っています。
RecipeEngine は、成功した各アクションの SHA-256 を監視します
フィンガープリント - 同じアクションの成功したアクションの数がcandidate_threshold=2以上の場合、自動的に候補にアップグレードされます。
レシピは、LLM の再推論なしで、次回アクション オプション (source= action_skill) として直接提案されます。
各実行が終了すると、CuratorDaemon
SkillCurator.apply_auto_merges() をバックグラウンド スレッド (daemon=True) で非同期的に開始し、重複 (信頼度 0.95) またはサブセット (信頼度 0.95) を自動的に検出してマージします。
0.90）スキルライブラリの拡張を防ぐスキル。
TOOLSET_CATALOG は 4 をカバーします
デフォルトで有効になっているツールのセット: ファイル、ターミナル、検索、スキル、さらにオプトイン プロジェクト (git / JSON / YAML / テスト探索) およびブラウザ (Playwright)
ヘッドレス Chromium、11 ツール)。計画中: メモリ、デリゲート、cronjob、computer_use。
AutonomyStoreとイベントソーシング
このパターンは、実行ライフサイクル全体を記録します: run_started → skill_selected → Candidates_ranked → action_selected → Approval_Decision →
観察 → 結果評価 → レシピ学習 → 学習レビュー → 実行終了

編集済み、完全かつ再生可能。
autonomy tui は豊富な端末インターフェイスをアクティブ化し、セッションを提供します
概要パネル、ラウンド記録、アクショントレイル、コンパクト/フル切り替えモード、/コマンドパネル。自然言語は AgentLoop に直接流れ込み、UI がツールを直接実行することはありません。
ollam (ネイティブ) および 8 OpenAI 互換をサポート
プロバイダー: openai-api、nvidia、openrouter、deepseek、xai、zai、kimi-coding、alibaba。 API キー管理とプロバイダーを提供する
設定すると、自律ドクターはエンドポイントの接続を確認できます。
TOOLSET_CATALOG — ツールセットと AI プロバイダー
各スキルは、名前、説明、必要なツール、プラットフォーム フィルターを含む SKILL.md 形式で定義されます。
Turn が実行されるたびに、model.select_procedure_skills()
利用可能なツール名に基づいて、最も関連性の高いスキルを動的にフィルタリングしてロードします。実行が成功し、しきい値に達すると、LearningLoop は自動的に新しいスキルを草案し、LearningProposal (CANDIDATE) を生成します。
ステータス）。
次のイベントのシーケンスは、AutonomyStore.record_event() の実際の呼び出しシーケンスに直接対応します。
自律フレームワークデザイナー & AI Agent システムエンジニア
Autonomy の自律型 AI エージェント フレームワークを設計および実装しました。主な貢献には以下が含まれます: Beam Search
候補者スコアリング システム、マネージド アクション ゲートウェイ (ActionGateway)、実行後の学習ループ (LearningLoop)、およびバックグラウンド スキル キュレーション デーモン (CuratorDaemon)。
AIの開発に取り組む
エージェントは、コード編集やブラウザの自動化からシステムのデバッグに至るまで、実際のエンジニアリング環境で確実に実行され、構成可能なスキル ライブラリとイベント ソーシング アーキテクチャによる完全な監査機能を備えています。
リンクトイン
GitHub
独自の自律型 AI エージェントを構築する準備はできていますか?
自律性をフォークし、ツールセットを追加し、独自のスキル ライブラリをトレーニングして、エージェントが現場で自律的に作業できるようにします。
GitHub に移動
著者に連絡する
⬡
自律性
自律型 AI エージェント フレームワーク · Bill Liu によって設計および実装されました
© 2025 Bill Liu · 自律フレームワーク · AgentLoop · ProcedureSkillLibrary

· キュレーターデーモン

## Original Extract

Bill Liu 打造的 Autonomy 框架：具備 Beam Search 候選排序、受管行動閘道、程序技能學習與後台技能策展守護程序的自律型 AI 代理人系統。

GitHub
AgentLoop · ConversationLoop
|
Python 3.13 · 自律執行中
Autonomy
自律型 AI
代理人框架
以 Beam Search 候選排序 、 受管行動閘道 與 程序技能學習 為核心，
打造真正能自主感知、規劃、執行並持續進化的 AI 代理人系統。
查看原始碼
0
內建 Procedure Skills
0
+ 內建工具
0
支援 AI Provider
0
Beam Width
設計理念
AgentLoop：自導向任務迴圈
每個執行回合（Turn）依固定五步驟推進，直到目標達成、被封鎖或超過步數上限
AgentLoop 是 Autonomy 的核心驅動器——一個「自導向任務迴圈」，不需要人工介入即可從目標出發，自主選擇工具、執行行動並從結果中學習。
每個 Turn 的執行流程：①從技能庫選取相關 ProcedureSkill ，②由 LLM 提案 + RecipeEngine 產出候選動作路徑，③以 5 維評分排序（Beam
Width=3），④透過 ActionGateway 受管執行，⑤評估結果並觸發學習。
終止條件明確： ACHIEVED （目標完成）、 BLOCKED （無法繼續）、 NO_CANDIDATES 、 APPROVAL_DENIED 、 MAX_STEPS_REACHED 、或
FAILED ——系統永遠不會陷入無限循環。
AgentLoop.run(goal, max_steps=12)
啟動後完全自主執行，每步驟透過 run_turn() 推進。支援互動模式（ interactive=True ）與批次模式，所有事件均以
run_id 記錄於 AutonomyStore 。
所有工具呼叫必須通過 ActionGateway 的
ApprovalPolicy 授權。每個動作帶有 RiskLevel （LOW / MEDIUM /
HIGH）、 expected_effect 與 verification_plan ，確保執行可控、可審計。
CandidateSelector 以
evidence_strength（+0.30）、purpose（+0.10）、risk（−0.35）、side_effects（−0.20）、penalty（−1.0）五個維度計算總分，選出前
beam_width=3 名候選路徑送往執行。
DeterministicOutcomeEvaluator 優先判斷執行失敗（直接
BLOCKED），成功但結果不確定時再交由 ModelAssistedOutcomeEvaluator 呼叫 LLM 語義判斷——最小化昂貴的模型呼叫次數。
ProcedureSkillLibrary 內建 13 個 bundled
skills，涵蓋
api-debugging、browser-navigation、code-editing、codebase-documentation、plan、procedure-skill-authoring、process-management、requesting-code-review、systematic-debugging、technical-spike、TDD、website-inspection、writing-plans，每步動態選取最相關技能載入。
每次 Run 結束後， LearningLoop.review_run()
自動觸發：若結果為 ACHIEVED 且有 ≥2 個成功步驟，自動起草新 ProcedureSkillDraft （信心度 0.85），生成
LearningProposal 等待審核。
RecipeEngine 監控每個成功動作的 SHA-256
指紋——當同一動作成功次數 ≥ candidate_threshold=2 ，自動升格為候選
Recipe，下次直接作為行動選項（source= action_skill ）提出，無需 LLM 重新推理。
每次 Run 結束後， CuratorDaemon
在後台執行緒（ daemon=True ）非同步啟動 SkillCurator.apply_auto_merges() ，自動偵測並合併重複（信心 0.95）或子集（信心
0.90）技能，防止技能庫膨脹。
TOOLSET_CATALOG 涵蓋 4
個預設啟用工具集：file、terminal、search、skills，加上 opt-in 的 project（git / JSON / YAML / 測試探索）與 browser（Playwright
headless Chromium，11 工具）。規劃中：memory、delegate、cronjob、computer_use。
AutonomyStore 以 Event Sourcing
模式記錄整個執行生命週期：run_started → skills_selected → candidates_ranked → action_selected → approval_decision →
observation → outcome_evaluated → recipe_learned → learning_review → run_finished，完整可重播。
autonomy tui 啟動豐富終端介面，提供 session
總覽面板、回合記錄、Action trail、compact/full 切換模式與 / 指令面板。自然語言直接流入 AgentLoop ，UI 永遠不直接執行工具。
支援 ollama （本地）與 8 個 OpenAI-compatible
providers：openai-api、nvidia、openrouter、deepseek、xai、zai、kimi-coding、alibaba。提供 API 金鑰管理與 provider
設定， autonomy doctor 可檢查端點走通性。
TOOLSET_CATALOG — 工具集 & AI Providers
每個技能以 SKILL.md 格式定義，包含名稱、描述、所需工具與平台篩選
每次執行 Turn 時， model.select_procedure_skills()
依可用工具名稱動態篩選並載入最相關技能；執行成功且達到門檻後， LearningLoop 將自動起草新技能並生成 LearningProposal （CANDIDATE
狀態）。
以下事件序列直接對應 AutonomyStore.record_event() 的實際呼叫順序
Autonomy 框架設計者 & AI Agent 系統工程師
設計並實作了 Autonomy 自律型 AI 代理人框架，核心貢獻包括：Beam Search
候選評分系統、受管行動閘道（ActionGateway）、後執行學習迴路（LearningLoop）、以及後台技能策展守護程序（CuratorDaemon）。
致力於讓 AI
代理人在真實工程環境中可靠運行——從程式碼編輯、瀏覽器自動化到系統除錯，透過可組合的技能庫與事件溯源架構實現完整可審計性。
LinkedIn
GitHub
準備好打造你的自律型 AI 代理人了嗎？
Fork Autonomy，添加你的工具集、訓練專屬技能庫，讓代理人在你的領域自主作業。
前往 GitHub
聯絡作者
⬡
Autonomy
自律型 AI 代理人框架 · 由 Bill Liu 設計與實作
© 2025 Bill Liu · Autonomy Framework · AgentLoop · ProcedureSkillLibrary · CuratorDaemon
