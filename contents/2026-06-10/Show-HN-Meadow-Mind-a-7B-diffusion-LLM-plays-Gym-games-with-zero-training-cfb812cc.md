---
source: "https://github.com/Hey-Meadow/meadow-mind"
hn_url: "https://news.ycombinator.com/item?id=48479887"
title: "Show HN: Meadow Mind – a 7B diffusion LLM plays Gym games with zero training"
article_title: "GitHub - Hey-Meadow/meadow-mind: Zero training, second-level reactions (~400ms). A language-rule decision mind on a local 7B diffusion LM. · GitHub"
author: "akaiHuang"
captured_at: "2026-06-10T19:00:34Z"
capture_tool: "hn-digest"
hn_id: 48479887
score: 2
comments: 0
posted_at: "2026-06-10T17:41:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Meadow Mind – a 7B diffusion LLM plays Gym games with zero training

- HN: [48479887](https://news.ycombinator.com/item?id=48479887)
- Source: [github.com](https://github.com/Hey-Meadow/meadow-mind)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T17:41:53Z

## Translation

タイトル: Show HN: Meadow Mind – 7B 拡散 LLM がトレーニングなしでジム ゲームをプレイ
記事のタイトル: GitHub - Hey-Meadow/meadow-mind: トレーニングなし、第 2 レベルの反応 (~400 ミリ秒)。ローカル 7B 拡散 LM 上の言語ルール決定心。 · GitHub
説明: トレーニングなし、第 2 レベルの反応 (約 400 ミリ秒)。ローカル 7B 拡散 LM 上の言語ルール決定心。 - Hey-Meadow/メドウ-マインド

記事本文:
GitHub - Hey-Meadow/meadow-mind: トレーニングなし、第 2 レベルの反応 (~400 ミリ秒)。ローカル 7B 拡散 LM 上の言語ルール決定心。 · GitHub
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
ヘイ、メドウ
/
草原の心
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット アセット アセット meadow_mind meadow_mind .gitignore .gitignore ライセンス ライセンス README.md README.md README.zh-TW.md README.zh-TW.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
トレーニングゼロ。第 2 レベルの反応 (~400 ミリ秒)。
言語ルールの意思決定心: ポリシーを 1 つの文として記述し、状態を 1 つの文として記述します。ローカル 7B モデルは、約 0.4 秒ごとに実際の決定を行います。 RL、報酬エンジニアリング、勾配、サンプルはありません。
🌐 デモ サイト : meadow-mind.pages.dev (中文) · English · 繁體中文 README
pip install meadow-mind # 初回使用時に重みを自動ダウンロード
from meadow_mind import MeadowMind 、タスク
min = MeadowMind () # 1 回ロードされ、デバイス上で実行されます
タスク = タスク 。登山車（）
心。チェック (タスク) # 正気ゲート: デシジョンテーブル試験
行動、情報＝心。 Decide ( task , obs ) # obs in、env action out (~0.4s)
結果
すべては公式の体育館環境で、物理学に触れず、トレーニングは不要です。以下の各フレームは、1 つの実際のモデルの決定に対応します。スクリプト化されたポリシーや編集された高速化はありません。
MountainCar のポリシーは、「ブランコのようにエネルギーを送り出すために、車が動いている方向と同じ方向に押す」という直感に反する文で、RL 報酬曲線全体を置き換えるものです。
リアルタイム反射 (ターンベースではなく壁時計)
モデルはスレッド内で実行され、障害物はリアルタイムで落下します。障害物が着地したときにまだ考えていると、本当にクラッシュします。
漏斗迷路では、両方が同じ行き止まりのポケットに落ちます。反応性（左）は永遠に口をついて歩きます。 Task(memory=True) (右) では、もがき、後退し、迂回しながら 22 ステップでゴールに到達します。唯一の違いは、認識メッセージの 5 つの単語です

そうだね。
意思決定の待ち時間: 従来の LLM と Meadow Mind の比較
従来の LLM エージェントは、動作する前に完全な応答を生成する必要があり、応答の長さに応じて待ち時間が長くなります。 Meadow Mind はルールと状況を読み取り、人間の反応速度 (0.3 ～ 0.4 秒) で、1 回の固定遅延パスで決定します。
従来の LLM エージェント Meadow Mind
───────────────
状態→長いプロンプト状態→一文（知覚者）
→ 答えを生成 → 一文ルール（ポリシー）
トークンごと (1.2 ～ 3.9 秒、→ 1 つの決定パス、固定 ~0.4 秒)
長さに応じて成長します) → アクション文字 (アクチュエーター)
→ フリーテキストを解析 → 導入前に試験ゲートを行う
なぜその下に拡散 LLM があるのか
Meadow Mind は、自己回帰モデルではなく、拡散言語モデル (MeadowCoder-7B) に基づいて構築されています。重要な違い:
そのうちの 2 つは、Meadow Mind を可能にするものです。1 つは複数段階の自己修正 (作業中に自分自身の下書きを修正できる)、もう 1 つは Σ による全体的なタスク認識 (答えにコミットする前に、何を求められているか、そして理解しているかどうかを認識する) です。
┌───────────────────────┐
│ ① コードを認識する: 数字 -> 一文 │
│「ポールが右に傾き、回転が速い。」 │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ ② ルール一文＝方針 │
│ 単語を編集して動作を編集する │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ ③心

7B オンデバイス モデルはルール + 状態を読み取ります、│
│ アクションレターに 1 回で回答します │
│ 決定パス、固定 ~0.4 秒 │
━━━━━━━━━━━━━━━━━━━━━━━┤
│ ④ アクチュエータ文字 -> 環境アクション │
━━━━━━━━━━━━━━━━━━━━━┘
ループには報酬はありません。 env スコアは単なる成績表です。改善は結果のフィードバックによって起こります。エピソード トレースはどの文が間違っていたかを示し、それを編集します。 (LunarLander は、知覚装置に 1 本のタッチダウン クッション ラインを追加することで、+27.5 の衝突から +251 の着陸に到達しました。10 秒。)
タスクを理解し、インプットとアウトプットを探ります。変数、アクション、勝敗条件。反応期限は約 0.4 秒よりも緩くする必要があります。すべてのアクションをリストし、その効果を観察します。
認識の言葉を構築します。現在の状況を説明する一文。バケットの連続値 (小さい/大きい、速い/遅い)。常に速度/トレンド項を含めてください。
ルールを刷り込みます。効果を「状況 X でアクション B を行う」に反転します。キーワード → 文字、1 層マッピング、複数選択のみ。
記憶力で決める。 「同じ状態に戻るのは失敗の合図ですか?」と尋ねてください。はい (迷路、探索、行き止まり) → タスク (記憶=True) 。いいえ (バランス、着地、追跡 - 繰り返しが仕事です) → やめてください。注釈は規制タスクに明らかに悪影響を及ぼします (CartPole の健全性 7/8 → 6/8)。わからない → やめます。ランナーはループを検出するとヒントを出力します。
試験を受けてください。あらゆる状況を予想される文字とともに列挙します。 min.check(task) は最大 1 回のミスでパスします。失敗した場合は、表現が不完全であることを意味します。言い換えて再確認しますが、トレーニングは必要ありません

NG。
または、5 つすべてをスキップします。meadow_mind.ai_prompt() とゲームの説明を任意のコード エージェントに渡すと、タスクが割り当てられます。試験のスコアを確認するだけです。
重みの解決: MEADOW_MIND_MODEL 環境 → 明示的なパス → ローカル キャッシュ ( ~/.meadow-mind/models/ ) → 自動ダウンロード。
フィールド
知覚(obs) -> str (またはメモリ付きの知覚(obs, task))
知覚層
ルール / オプションテキスト / オプション / act_text
一文ポリシーとその複数選択アクション
正気
試験: [(状況文、期待される文字)]
メモリ / mem_key
ワーキングメモリスイッチ (デフォルトオフ) + ステートキー fn
env_id / env_kwargs / max_steps / ジャッジ
環境配線とレポートカード
Memory=True の場合、ランナーは task.visited を自動追跡します。注釈を付けるには、知覚内で task.seen(key) を使用します。 （安全です、すでに訪問済みです）。
meadow-mind carpole # sanity Gate -> 1 エピソードを再生 -> ビデオ + 評決
正直な限界
リアクションフロアは 1 つの決定パス (~0.4 秒 ≈ 2 Hz) です。より厳しい締め切り（1 m ポール、ポンの軌道予測）は、今日では手が届きません。
状況が文章で言え、方針がルールに適合するタスクに適しています。継続的な高精度制御はできません。
知覚者は人間によって設計されています (または ai_prompt() を介して AI によって生成されます)。モデルの仕事はルールを読んで決定することです。
v0.2 — 早期の行動を伴う階層化された認識: ネットワークを通じて信頼を蓄積し、閾値を超えたときに行動します。簡単な状況では、約 0.15 秒付近に着地するはずです。
ルール学習ループ: 失敗したエピソードから、MountainCar のスイング トリックなどのルールを自動的に発見します (勾配なし - 学習された成果物は読みやすい文です)。
トレーニングなし、第 2 レベルの反応 (約 400 ミリ秒)。ローカル 7B 拡散 LM 上の言語ルール決定心。
meadow-mind.pages.dev
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ

y
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Zero training, second-level reactions (~400ms). A language-rule decision mind on a local 7B diffusion LM. - Hey-Meadow/meadow-mind

GitHub - Hey-Meadow/meadow-mind: Zero training, second-level reactions (~400ms). A language-rule decision mind on a local 7B diffusion LM. · GitHub
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
Hey-Meadow
/
meadow-mind
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits assets assets meadow_mind meadow_mind .gitignore .gitignore LICENSE LICENSE README.md README.md README.zh-TW.md README.zh-TW.md pyproject.toml pyproject.toml View all files Repository files navigation
Zero training. Second-level reactions (~400 ms).
A language-rule decision mind: write the policy as one sentence, describe the state as one sentence, and a local 7B model makes a real decision every ~0.4 s. No RL, no reward engineering, no gradients, no samples.
🌐 Demo site : meadow-mind.pages.dev (中文) · English · 繁體中文 README
pip install meadow-mind # weights auto-download on first use
from meadow_mind import MeadowMind , tasks
mind = MeadowMind () # loads once, runs on-device
task = tasks . mountaincar ()
mind . check ( task ) # sanity gate: decision-table exam
action , info = mind . decide ( task , obs ) # obs in, env action out (~0.4s)
Results
All on official Gymnasium environments, untouched physics, zero training . Every frame below corresponds to one real model decision; no scripted policy, no edited speed-ups.
The MountainCar policy is one counterintuitive sentence — "push in the same direction the car is moving, to pump energy like a swing" — which replaces an entire RL reward curve.
Real-time reflex (wall-clock, not turn-based)
The model runs in a thread while obstacles fall in real time. If it is still thinking when the obstacle lands, it really crashes.
A funnel maze forces both runs into the same dead-end pocket. Reactive (left) paces at its mouth forever; with Task(memory=True) (right) it struggles, backs out, and detours to the goal in 22 steps. The only difference is five words in the perception sentence.
Decision latency: traditional LLM vs Meadow Mind
A traditional LLM agent must generate its full answer before acting — and latency grows with answer length. Meadow Mind reads the rule and the situation and decides in one fixed-latency pass , right at human reaction speed (0.3–0.4 s):
Traditional LLM agent Meadow Mind
───────────────────── ───────────
state → long prompt state → one sentence (Perceiver)
→ generate the answer → one sentence rule (Policy)
token by token (1.2–3.9 s, → ONE decision pass, fixed ~0.4 s
grows with length) → action letter (Actuator)
→ parse free text → act exam-gated before deployment
Why a diffusion LLM underneath
Meadow Mind is built on a diffusion language model (MeadowCoder-7B), not an autoregressive one. The differences that matter:
Two of these are what make Meadow Mind possible: multi-step self-correction (it can fix its own draft while working) and global task perception with Σ (it knows what it is being asked — and whether it understands — before committing to an answer).
┌────────────────────────────────────────────────────┐
│ ① Perceiver your code: numbers -> one sentence │
│ "The pole tilts right, fast spin." │
├────────────────────────────────────────────────────┤
│ ② Rule one sentence = the policy │
│ edit behavior by editing words │
├────────────────────────────────────────────────────┤
│ ③ Mind 7B on-device model reads rule+state, │
│ answers an action letter in a single │
│ decision pass, fixed ~0.4 s │
├────────────────────────────────────────────────────┤
│ ④ Actuator letter -> env action │
└────────────────────────────────────────────────────┘
There is no reward in the loop. The env score is only a report card; improvement happens by outcome feedback : the episode trace shows which sentence was wrong, and you edit it. (LunarLander went from a +27.5 crash to a +251 landing by adding one touchdown-cushion line to the perceiver. Ten seconds.)
Understand the task, explore input-output. Variables, actions, win/lose conditions; the reaction deadline must be looser than ~0.4 s. List every action and watch its effect.
Build perception words. One sentence describing the current situation. Bucket continuous values (small/big, fast/slow); always include a velocity/trend term.
Imprint the rule. Invert the effects into "on situation X do action B". Keyword → letter, one-layer mapping, multiple choice only.
Decide on memory. Ask: "is revisiting the same state a failure signal?" Yes (maze, exploration, dead ends) → Task(memory=True) . No (balance, landing, tracking — repetition IS the job) → keep it off; annotations measurably hurt regulation tasks (CartPole sanity 7/8 → 6/8). Unsure → leave off; the runner prints a hint when it detects looping.
Take the exam. Enumerate every situation with its expected letter; mind.check(task) passes with at most 1 miss. Failures mean the wording is incomplete — rephrase and re-check, no training.
Or skip all five: hand meadow_mind.ai_prompt() plus your game description to any code agent, and it wires the task for you. You only review the exam score.
Weight resolution: MEADOW_MIND_MODEL env → explicit path → local cache ( ~/.meadow-mind/models/ ) → auto-download.
Field
perceive(obs) -> str (or perceive(obs, task) with memory)
perception layer
rule / option_text / options / act_text
the one-sentence policy and its multiple-choice actions
sanity
the exam: [(status sentence, expected letter)]
memory / mem_key
working-memory switch (default off) + state key fn
env_id / env_kwargs / max_steps / judge
environment wiring and report card
With memory=True the runner auto-tracks task.visited ; use task.seen(key) inside perceive to annotate, e.g. (safe, already visited) .
meadow-mind cartpole # sanity gate -> play one episode -> video + verdict
Honest limits
Reaction floor is one decision pass (~0.4 s ≈ 2 Hz). Tighter deadlines (1 m pole, Pong trajectory prediction) are out of reach today.
Suited to tasks whose situations can be said in a sentence and whose policy fits a rule. Continuous high-precision control is not.
The perceiver is human-designed (or AI-generated via ai_prompt() ); the model's job is reading the rule and deciding.
v0.2 — layered perception with early action: accumulate confidence through the network and act when it crosses a threshold; easy situations should land near ~0.15 s.
Rule-learning loop: discover rules like MountainCar's swing trick from failed episodes automatically (no gradients — the learned artifact is a readable sentence).
Zero training, second-level reactions (~400ms). A language-rule decision mind on a local 7B diffusion LM.
meadow-mind.pages.dev
Resources
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
