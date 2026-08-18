---
source: "https://github.com/Steel-predictor-project/steel-llm-eval"
hn_url: "https://news.ycombinator.com/item?id=49346442"
title: "Show HN: I benchmarked LLMs on predicting knife steel properties"
article_title: "GitHub - Steel-predictor-project/steel-llm-eval: Open benchmark: how well can LLMs predict knife-steel properties (edge retention, toughness) from chemical composition, scored against CATRA/Charpy lab measurements. · GitHub"
image: "https://opengraph.githubassets.com/9d85fca86df833329324c3c89207d82b5056bc950c515d2f5327a2b52e80a412/Steel-predictor-project/steel-llm-eval"
author: "p-s-v"
captured_at: "2026-08-18T15:22:39Z"
capture_tool: "hn-digest"
hn_id: 49346442
score: 1
comments: 0
posted_at: "2026-08-18T14:49:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I benchmarked LLMs on predicting knife steel properties

- HN: [49346442](https://news.ycombinator.com/item?id=49346442)
- Source: [github.com](https://github.com/Steel-predictor-project/steel-llm-eval)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T14:49:15Z

## Translation

タイトル: HN を表示: ナイフ鋼の特性の予測について LLM のベンチマークを実施しました
記事のタイトル: GitHub - Steel-predictor-project/steel-llm-eval: オープン ベンチマーク: LLM が化学組成からナイフ鋼の特性 (エッジ保持力、靱性) をどの程度予測できるか (CATRA/シャルピー ラボ測定に対してスコア化)。 · GitHub
説明: オープンベンチマーク: LLM が化学組成からナイフ鋼の特性 (エッジ保持力、靱性) をどの程度予測できるか (CATRA/Charpy ラボ測定に対してスコア化)。 - Steel-predictor-project/steel-llm-eval

記事本文:
GitHub - Steel-predictor-project/steel-llm-eval: オープン ベンチマーク: LLM が化学組成からナイフ鋼の特性 (エッジ保持力、靱性) をどの程度予測できるか (CATRA/Charpy ラボ測定に対してスコア化)。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
鉄鋼予測プロジェクト
/
スチール-llm-評価
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
3コム

mit 3 フォルダーとファイルをコミットします
データ データ ドキュメント ドキュメント ハーネス ハーネス 結果 結果 .gitignore .gitignore ライセンス ライセンス通知 通知 README.md README.md 要件.txt 要件.txt run_benchmark.sh run_benchmark.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
大規模な言語モデルは、化学組成のみからナイフ鋼の特性をどの程度正確に予測できるでしょうか?
LLM に鋼の組成のみ (例: C=1.45%、Cr=20%、V=4%、Mo=1%、粉末冶金: はい) を与え、2 つの特性を 1 ～ 10 のスケールで評価するよう依頼し、客観的な実験室測定に対してそれらの評価をスコア化する、オープンで再現可能なベンチマーク。
エッジ保持 ← CATRA 標準化機械切断テスト (カードストックの合計カット、mm) — 48 鋼
靭性 ← シャルピー衝撃エネルギー (ft-lbs) — 12 鋼
スコアリングにはスケールフリー (ランク相関 + ペアごとのランク付け精度) があるため、モデルは 1 ～ 10 のスケールをどのように調整するかではなく、鋼材を正しく注文したかどうかのみで判断されます。
📊 チャートと分析のあるサイト: https://steel-predictor-project.github.io/steel-llm-eval/ · ドキュメント: 方法論 · 結果と分析
平均スピアマン順位相関 (ρ) と測定値によってランク付けされます。高いほど良いです。 1.0 = 完全な順序付け、0.0 = ランダム。
エッジ保持率 n=48 (CATRA)、靭性 n=12 (シャルピー)。ゼロショット、温度 0、鋼あたり 1 つのサンプル。
† 公平性に関する重要な注意事項: 参照 ML モデル ( Steel-predictor ) は、これらの同じ CATRA/Charpy 測定でトレーニングされたため、ここでのスコアは大部分がサンプル内であり、ゼロショット LLM との公平な直接比較としてではなく、上部参照バーとして表示されます。モデルの正直なサンプル外パフォーマンスは、そのリポジトリで報告されている LOOCV MAE (0.391) です。対照的に、LLM はこのラベル付きセットを見たことがありません。
LLM はエッジ保持のランク付けが真に得意です (ρ ≈ 0)

.85〜0.92）。耐摩耗性は組成 (炭化物形成元素 - C、V、Cr、W、Mo) で強力かつ読みやすくエンコードされており、フロンティア モデルはその化学反応を明確に「認識」しています。
タフネスは彼らが苦労するところです (ρ 0.38–0.84)。それは、組成文字列からは明らかではない、より微妙な要因 (炭化物のサイズ/分布、粉末冶金処理、マトリックスの状態) に依存し、モデル間のばらつきは大きくなります。
フロンティア＞小さい。クロード・ソネットとジェミニがリード。小型/安価なモデルは、エッジ保持力で競争力を維持しながらも、靭性が大幅に低下します。
git clone https://github.com/Steel-predictor-project/steel-llm-eval.git
cd スチール-llm-eval
import OPENROUTER_API_KEY=sk-or-... # 1 つのキー → OpenAI、Anthropic、Google、Meta、DeepSeek など
./run_benchmark.sh # すべてのモデルを実行し、リーダーボードを再構築します
単一モデルを実行するか、API キーを使用せずに簡単なオフライン健全性チェックを実行します。
python harness/run_eval.py --model anthropic/claude-sonnet-5
python harness/run_eval.py --provider モック # 決定論的ヒューリスティック、キーは必要ありません
Pythonハーネス/score.py
スチールごとの生の応答は results/raw_<model>.csv に書き込まれます。スコアは results/scores.csv および results/leaderboard.md に保存されます。
プロンプト ( harness/prompts.py ) — 固定システム + ユーザー プロンプトは、モデルに組成、PM フラグ、およびテスト硬度 (既知の場合) を与え、JSON: {"edge_retention": n, "toughness": n} を要求します。どのモデルでも同じです。
実行 ( harness/run_eval.py ) — OpenRouter 経由で 51 個すべての鋼材のモデルをクエリし、JSON を解析します。
スコア ( harness/score.py ) — 測定値との比較:
スピアマン ρ とケンダル τ の順位相関 (見出し、スケールフリー)。
ペアごとの精度 — すべての鋼ペアにわたって、モデルが測定と同じ方法でそれらを順序付ける頻度 (同点は除外)。
正規化された MAE — キャリブレーション健全性チェック

測定値を 1 ～ 10 にスケーリングする term min-max (セカンダリ。scores.csv を参照)。
ベースライン — 専用の ML モデル (上の参照、上のサンプル内の警告) と定数予測子 (下限)。
方法論の注意事項と制限事項
グラウンド トゥルースは客観的な測定のみです (CATRA、Charpy)。専門家による主観的な 1 ～ 10 の評価は、採点の際には使用されません。
構成のみ。モデルには熱処理プロトコルや形状 (記録される場合は硬度のみ) が指示されないため、これは化学のみが意味するもの、つまり参照モデルが動作するのと同じ制約を測定します。
小さい靭性セット (n=12)。靱性 ρ は決定的なものではなく、指標として扱います。温度 0 で鋼あたり 1 つのサンプル (自己一貫性/複数サンプルの平均化はまだ行われていません)。
モデルが異なれば 1 ～ 10 のスケールが異なる方法で調整されるため、ランク メトリクスが主要になります。ランキングは比較可能であり、意思決定に関連するものです。
ベンチマーク ( data/benchmark.csv ) は、Steel-predictor プロジェクトの処理されたデータセットから派生します。 2 つのグラウンド トゥルース測定値は次のものから得られます。
エッジ保持 (CATRA、48 スチール) — Larrin Thomas、「48 ナイフ スチールのエッジ保持のテスト」(2020)、KnifeSteelNerds.com 。このベンチマークのエッジ保持に関するグラウンド トゥルース全体は、Larrin Thomas が公開した CATRA 測定値に基づいて構築されています。これは彼の功績です。
靭性 (シャルピー、12 鋼) — Crucible Industries が公開したデータシート (CPM シリーズ)。
組成と硬度テスト - メーカーのデータシート (Crucible、Böhler-Uddeholm/voestalpine、Carpenter、Alleima、Hitachi/Proterial) および出版された文献。
個々のソースはすべて、Steel-predictor リポジトリの DATA_SOURCES.md 内のリンクで列挙されます。基礎となる事実の測定値は、元の出版社の所有物のままです。このリポジトリは、独自の正規化されたコンパイルと派生のみを再配布します。

dの特徴。
コード: Apache-2.0 (ライセンス)。厳選されたベンチマーク データ + リファレンス モデルの出力: CC BY 4.0 ( data/LICENSE )、このプロジェクトのコンパイル/派生機能のみをカバーします。帰属要求: 「Steel Property Predictor Project」 (このリポジトリへのリンク付き)。
PR は、モデル ( run_benchmark.sh のリストを拡張)、プロンプトのバリアント (少数のショット、思考の連鎖、自己一貫性)、または追加の測定された鋼材 (引用された公開ソース付き) を追加することを歓迎します。主観的な評価データセットをグラウンド トゥルースとして追加しないでください。
オープンベンチマーク: LLM が化学組成からナイフ鋼の特性 (エッジ保持力、靱性) をどの程度予測できるか (CATRA/Charpy 実験室測定値に対してスコア化)。
github.com/Steel-predictor-project/Steel-predictor トピック
Readme Apache-2.0 ライセンス アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open benchmark: how well can LLMs predict knife-steel properties (edge retention, toughness) from chemical composition, scored against CATRA/Charpy lab measurements. - Steel-predictor-project/steel-llm-eval

GitHub - Steel-predictor-project/steel-llm-eval: Open benchmark: how well can LLMs predict knife-steel properties (edge retention, toughness) from chemical composition, scored against CATRA/Charpy lab measurements. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Steel-predictor-project
/
steel-llm-eval
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
data data docs docs harness harness results results .gitignore .gitignore LICENSE LICENSE NOTICE NOTICE README.md README.md requirements.txt requirements.txt run_benchmark.sh run_benchmark.sh View all files Repository files navigation
How well can large language models predict knife-steel properties from chemical composition alone?
An open, reproducible benchmark that gives an LLM only a steel's composition (e.g. C=1.45%, Cr=20%, V=4%, Mo=1%, powder-metallurgy: yes ) and asks it to rate two properties on a 1–10 scale, then scores those ratings against objective laboratory measurements :
Edge retention ← CATRA standardized machine-cutting test (total card stock cut, mm) — 48 steels
Toughness ← Charpy impact energy (ft-lbs) — 12 steels
Scoring is scale-free (rank correlation + pairwise ranking accuracy), so a model is judged purely on whether it orders steels correctly, not on how it calibrates the 1–10 scale.
📊 Site with charts & analysis: https://steel-predictor-project.github.io/steel-llm-eval/ · Docs: methodology · results & analysis
Ranked by mean Spearman rank correlation (ρ) vs. the measurements. Higher is better; 1.0 = perfect ordering, 0.0 = random.
Edge retention n=48 (CATRA), toughness n=12 (Charpy). Zero-shot, temperature 0, one sample per steel.
† Important fairness caveat: the reference ML model ( Steel-predictor ) was trained on these same CATRA/Charpy measurements , so its scores here are largely in-sample and are shown as an upper-reference bar, not as a fair head-to-head with the zero-shot LLMs. The model's honest out-of-sample performance is its LOOCV MAE (0.391), reported in that repo. The LLMs, by contrast, have never seen this labeled set.
LLMs are genuinely good at ranking edge retention (ρ ≈ 0.85–0.92). Wear resistance is strongly and legibly encoded in composition (carbide-forming elements — C, V, Cr, W, Mo), and frontier models clearly "know" that chemistry.
Toughness is where they struggle (ρ 0.38–0.84). It depends on subtler factors (carbide size/distribution, powder-metallurgy processing, matrix state) that aren't obvious from a composition string, and the spread across models is large.
Frontier > small. Claude Sonnet and Gemini lead; the smaller/cheaper models drop off sharply on toughness while staying competitive on edge retention.
git clone https://github.com/Steel-predictor-project/steel-llm-eval.git
cd steel-llm-eval
export OPENROUTER_API_KEY=sk-or-... # one key → OpenAI, Anthropic, Google, Meta, DeepSeek, ...
./run_benchmark.sh # runs every model and rebuilds the leaderboard
Run a single model, or a quick offline sanity check with no API key:
python harness/run_eval.py --model anthropic/claude-sonnet-5
python harness/run_eval.py --provider mock # deterministic heuristic, no key needed
python harness/score.py
Raw per-steel responses are written to results/raw_<model>.csv ; scores to results/scores.csv and results/leaderboard.md .
Prompt ( harness/prompts.py ) — a fixed system + user prompt gives the model the composition, PM flag, and test hardness (when known) and asks for JSON: {"edge_retention": n, "toughness": n} . Identical for every model.
Run ( harness/run_eval.py ) — queries a model for all 51 steels via OpenRouter and parses the JSON.
Score ( harness/score.py ) — vs. the measurements:
Spearman ρ and Kendall τ rank correlation (headline; scale-free).
Pairwise accuracy — over all steel pairs, how often the model orders them the same way the measurement does (ties excluded).
Normalized MAE — a calibration sanity check after min-max scaling the measurement to 1–10 (secondary; see scores.csv ).
Baselines — the purpose-built ML model (upper reference, in-sample caveat above) and a constant predictor (floor).
Methodology notes & limitations
Ground truth is objective measurement only (CATRA, Charpy). No subjective 1–10 expert ratings are used anywhere in scoring.
Composition-only. Models are not told heat-treat protocol or geometry (only hardness where recorded), so this measures what chemistry alone implies — the same constraint the reference model operates under.
Small toughness set (n=12). Treat toughness ρ as indicative, not definitive; single sample per steel at temperature 0 (no self-consistency / multi-sample averaging yet).
Rank metrics are primary precisely because different models calibrate the 1–10 scale differently; ranking is what's comparable and decision-relevant.
The benchmark ( data/benchmark.csv ) is derived from the processed dataset of the Steel-predictor project. The two ground-truth measurements come from:
Edge retention (CATRA, 48 steels) — Larrin Thomas, "Testing the Edge Retention of 48 Knife Steels" (2020), KnifeSteelNerds.com . This benchmark's entire edge-retention ground truth is built on Larrin Thomas's published CATRA measurements — full credit to him.
Toughness (Charpy, 12 steels) — Crucible Industries published datasheets (CPM series).
Compositions & test hardness — manufacturer datasheets (Crucible, Böhler-Uddeholm/voestalpine, Carpenter, Alleima, Hitachi/Proterial) plus published literature.
Every individual source is enumerated with links in the Steel-predictor repo's DATA_SOURCES.md . Underlying factual measurements remain the property of their original publishers; this repo redistributes only its own normalized compilation and derived features.
Code: Apache-2.0 ( LICENSE ). Curated benchmark data + reference model outputs: CC BY 4.0 ( data/LICENSE ), covering only this project's compilation/derived features. Attribution requested: "Steel Property Predictor Project" with a link to this repo.
PRs welcome to add models (extend the list in run_benchmark.sh ), prompt variants (few-shot, chain-of-thought, self-consistency), or additional measured steels (with cited public sources). Please don't add subjective-rating datasets as ground truth.
Open benchmark: how well can LLMs predict knife-steel properties (edge retention, toughness) from chemical composition, scored against CATRA/Charpy lab measurements.
github.com/Steel-predictor-project/Steel-predictor Topics
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
