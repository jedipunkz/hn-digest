---
source: "https://github.com/zzvimercm-git/mirofish-calibration"
hn_url: "https://news.ycombinator.com/item?id=48635080"
title: "Show HN: A local rig to test if AI social simulation predicts reality"
article_title: "GitHub - zzvimercm-git/mirofish-calibration · GitHub"
author: "zzvimercm"
captured_at: "2026-06-22T20:02:36Z"
capture_tool: "hn-digest"
hn_id: 48635080
score: 1
comments: 0
posted_at: "2026-06-22T19:45:43Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A local rig to test if AI social simulation predicts reality

- HN: [48635080](https://news.ycombinator.com/item?id=48635080)
- Source: [github.com](https://github.com/zzvimercm-git/mirofish-calibration)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T19:45:43Z

## Translation

タイトル: Show HN: AI ソーシャル シミュレーションが現実を予測するかどうかをテストするローカル リグ
記事タイトル: GitHub - zzvimercm-git/mirofish-calibration · GitHub
説明: GitHub でアカウントを作成して、zzvimercm-git/mirofish-calibration の開発に貢献します。

記事本文:
GitHub - zzvimercm-git/mirofish-calibration · GitHub
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
zzvimercm-git
/
mirofish-キャリブレーション
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
zzvimercm-git/mirofish-calibration
メイン

ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット ケース ケース ハーネス ハーネス .env.example .env.example .gitignore .gitignore ライセンス ライセンス README.md README.md 要件.txt 要件.txt run.py run.py すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI 社会シミュレーションは実際に現実を予測するのでしょうか? — キャリブレーションリグ
マルチエージェント「ソーシャル シミュレーション」エンジン (MiroFish — 16k★、OASIS/CAMEL-AI) は、ドキュメントを入力し、何百もの AI ペルソナを生成し、出荷前に大衆の反応を予測することを約束します。このカテゴリーは人気があり、資金も豊富です。
問題の 1 つは、誰もキャリブレーションを公開していないことです。デモでは、1 つのケースで 1 つの印象的な実行を示し、「ほら、予測したよ!」と言います。シミュレーションは実際に単一の LLM を要求するだけで優れていますか?誰もそれを測定しません。
これはそれを測定する、小さくて正直なリグです。 Ollama 上で 100% ローカルで実行されます (ソブリン、クラウドなし)。
⚠️ 調査結果の前に制限事項をお読みください。これはリハーサルであり、評決ではない。以下を参照してください。
TL;DR (暫定 — n=5 合成ケース、ローカル qwen2.5:7b)
人々が何を言うか (感情の方向性): 単一の LLM が粗雑なマルチエージェントの群れを結びつけます。ハードケースではどちらも平凡 (~60%)。
どのような反対意見が表面化するかについては、単一の LLM が明らかに勝ちます (~98% 対 ~70% を思い出してください)。
「魔法の」シグナルの集合体 (バイラル度の大きさ、分極化) については、シミュレーションが得意であるはずのものです。このスケールでは数値はノイズです。スピアマン ρ は実行間で符号を反転します (+0.71 ↔ −0.71; +0.82 ↔ +0.10)。 n=5 では、ρ≈±0.7 は有意ですらありません。
エージェント インタラクション ラウンド (MiroFish の核となる論文) を追加しても、この粗雑な形式では役に立ちませんでした。
結論: 小規模では、「予測マジック」はコイントスと区別がつきません。それは MiroFish を反証するものではなく、立証責任を転換するものです

カテゴリに追加すると、デモを信頼するのではなく、実際にテストするためのリグが提供されます。
ヘッドラインの結果 (5 倍の平均、ローカル qwen2.5:7b)
予測者
センチメント監督。
異議申し立てのリコール
異議申し立て
マグニチュード（ランク）
二極化（ランク）
mini_swarm (対話なし)
64%
71%
62%
+0.10
−0.47
single_llm (1 つのゼロショット コール)
52%
84%
71%
+0.22
+0.05
愚か（常に「混合」）
40%
0%
0%
該当なし
該当なし
単一の LLM が突破すべきハードルです。粗雑な群れはそうではありません。
⚠️ 制限事項 (前面と中央 - これが重要です)
n=5、症例は合成（手書き、説明）です。これは方法論のリハーサルであり、現実世界に関する証拠ではありません。
ここでの swarm は未加工のプロキシであり、MiroFish ではありません。本物の MiroFish には、より多くのエージェントとより豊富なインタラクション ダイナミクスがあります。このリグは、単純なペルソナ平均化とおもちゃのインタラクション ラウンドをテストします。実際の MiroFish は (まだ) テストしていません。
1 つの小さなローカル モデル (qwen2.5:7b)。より大きなモデルや別のモデルではすべてが変わる可能性があります。
5 ポイントの順位相関は統計的に意味がありません。ここでは、振幅/偏光を信号ではなくノイズの図として扱います。
→ 本当の答えを得るには、文書化されたグラウンド トゥルース、複数のシード、実際の MiroFish エンジンを備えた数十の実際のケースが必要です。それがオープンワークです。
Cases (cases/*.yaml): 実際の刺激 + その既知の反応 (グラウンド トゥルース)。
プレディクター (交換可能): mirofish (実際のシム - 実装するアダプター スタブ)、mini_swarm / swarm_x (未加工の群れ、インタラクションなし/インタラクションあり)、single_llm (ビートするベースライン)、dam (健全性)。
メトリック: センチメントの方向性、反対意見の再現率/精度 (セマンティック LLM 判定)、マグニチュードと二極化のランク相関。
レポート: --runs N を使用して実行間のノイズを平均化した正直な比較。
pip install -r required.txt # または: python -m venv .venv && .venv/bin/pip install -r 要件。

テキスト
cp .env.example .env # デフォルトではローカルの Ollama をポイントします
オラマは qwen2.5:7b をプルします
python run.py --predictors single_llm,dumb # ベースライン、高速
python run.py --predictors swarm_x,mini_swarm,single_llm --runs 5 # 実際の比較
未解決の質問/貢献
このリグの優れている点は、ケースと SIM アダプターだけです。 PR の方は大歓迎です:
文書化されたグラウンドトゥルースを含む実際のケースを追加します (case/case_01_template.yaml)。カットオフ後のイベントを優先します (そうでない場合、LLM は予測ではなく記憶します)。
MiroFish アダプター ( harness/adapters/mirofish.py ) を実装します。これを実際のエンジンでの判定に変える実際の統合です。
複数のシードを使用して N≥30 で実行し、集約信号がノイズ フロアを乗り越えるかどうかを報告します。
MiroFish および OASIS / CAMEL-AI の一連の作業の背後にある前提をストレステストするために構築されました。これらのプロジェクトに多大な敬意を表します。このリグは、デモではなくメソッドを使用して、カテゴリ自体を証明するために存在します。
私は実際のエージェント システムを構築するインフラ/DevOps エンジニアです。エージェント AI スペースには印象的なデモが満載ですが、測定についてはほとんどありません。私は、不快な真実をお世辞にするデモよりも、不快な真実を伝えるリグを出荷したいと思っています。主張ではなく証拠。
読み込み中にエラーが発生しました。このページをリロードしてください。
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

Contribute to zzvimercm-git/mirofish-calibration development by creating an account on GitHub.

GitHub - zzvimercm-git/mirofish-calibration · GitHub
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
zzvimercm-git
/
mirofish-calibration
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
zzvimercm-git/mirofish-calibration
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit cases cases harness harness .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md requirements.txt requirements.txt run.py run.py View all files Repository files navigation
Does AI social simulation actually predict reality? — a calibration rig
Multi-agent "social simulation" engines (à la MiroFish — 16k★, OASIS/CAMEL-AI) promise: feed in a document, spawn hundreds of AI personas, and predict how the public will react — before you ship. The category is hot and well-funded.
One problem: nobody publishes the calibration. The demos show one impressive run on one case and say "look, it predicted!". Does the simulation actually beat just asking a single LLM ? Nobody measures it.
This is a small, honest rig that measures it. Runs 100% locally on Ollama (sovereign, no cloud).
⚠️ Read the limitations before the findings. This is a rehearsal, not a verdict. See below.
TL;DR (preliminary — n=5 synthetic cases, local qwen2.5:7b)
On what people will say (sentiment direction): a single LLM ties a crude multi-agent swarm. Both mediocre on hard cases (~60%).
On which objections will surface : a single LLM wins clearly (recall ~98% vs ~70%).
On the aggregate "magic" signals (virality magnitude, polarization) — the things simulation is supposed to be good at: the numbers are noise at this scale. Spearman ρ flips sign between runs (+0.71 ↔ −0.71; +0.82 ↔ +0.10). At n=5, ρ≈±0.7 isn't even significant.
Adding an agent-interaction round (the core MiroFish thesis) did not help in this crude form.
Conclusion: at small scale the "predictive magic" is indistinguishable from a coin flip. That doesn't disprove MiroFish — it shifts the burden of proof onto the category , and gives you a rig to actually test it instead of trusting a demo.
Headline result (5× averaged, local qwen2.5:7b)
Predictor
Sentiment dir.
Objection recall
Objection prec.
Magnitude (rank)
Polarization (rank)
mini_swarm (no interaction)
64%
71%
62%
+0.10
−0.47
single_llm (one zero-shot call)
52%
84%
71%
+0.22
+0.05
dumb (always "mixed")
40%
0%
0%
n/a
n/a
The single LLM is the bar to beat. A crude swarm doesn't.
⚠️ Limitations (front and center — this is the whole point)
n=5, and the cases are synthetic (hand-written, illustrative). This is a methodology rehearsal, not evidence about the real world.
The swarm here is a crude proxy, NOT MiroFish. Real MiroFish has many more agents and richer interaction dynamics. This rig tests naive persona-averaging and a toy interaction round — it does not (yet) test real MiroFish.
One small local model (qwen2.5:7b). A bigger/different model may change everything.
5-point rank correlations are not statistically meaningful. Treat magnitude/polarization here as noise illustration , not signal.
→ To get a real answer you need: dozens of real cases with documented ground truth, multiple seeds, and the actual MiroFish engine. That's the open work.
Cases ( cases/*.yaml ): a real stimulus + its known reaction (ground truth).
Predictors (interchangeable): mirofish (the real sim — adapter stub to implement), mini_swarm / swarm_x (crude swarm, no/with interaction), single_llm (the baseline to beat), dumb (sanity).
Metrics : sentiment direction, objection recall/precision (semantic LLM-judge), magnitude & polarization rank correlation.
Report : honest comparison, with --runs N to average away run-to-run noise.
pip install -r requirements.txt # or: python -m venv .venv && .venv/bin/pip install -r requirements.txt
cp .env.example .env # points at local Ollama by default
ollama pull qwen2.5:7b
python run.py --predictors single_llm,dumb # baselines, fast
python run.py --predictors swarm_x,mini_swarm,single_llm --runs 5 # the real comparison
Open questions / contributing
This rig is only as good as its cases and its sim adapter. PRs very welcome:
Add real cases with documented ground truth ( cases/case_01_template.yaml ). Prefer post-cutoff events (else the LLM remembers instead of predicting).
Implement the MiroFish adapter ( harness/adapters/mirofish.py ) — the one real integration that turns this into a verdict on the actual engine.
Run at N≥30 with multiple seeds and report whether the aggregate signals survive the noise floor.
Built to stress-test the premise behind MiroFish and the OASIS / CAMEL-AI line of work. Huge respect to those projects — this rig exists to help the category prove itself , with method instead of demos.
I'm an infra/DevOps engineer who builds real agentic systems. The agentic-AI space is full of impressive demos and thin on measurement. I'd rather ship a rig that tells the uncomfortable truth than a demo that flatters it. Proof, not claims.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
