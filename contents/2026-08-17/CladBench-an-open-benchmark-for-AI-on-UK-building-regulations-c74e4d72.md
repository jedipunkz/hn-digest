---
source: "https://github.com/cladbrain/cladbench"
hn_url: "https://news.ycombinator.com/item?id=49328898"
title: "CladBench – an open benchmark for AI on UK building regulations"
article_title: "GitHub - cladbrain/cladbench: An open evaluation benchmark for large language models on the UK and EU built environment — 536 questions across twelve categories. · GitHub"
author: "tamilselvan77"
captured_at: "2026-08-17T11:17:24Z"
capture_tool: "hn-digest"
hn_id: 49328898
score: 1
comments: 0
posted_at: "2026-08-17T10:44:09Z"
tags:
  - hacker-news
  - translated
---

# CladBench – an open benchmark for AI on UK building regulations

- HN: [49328898](https://news.ycombinator.com/item?id=49328898)
- Source: [github.com](https://github.com/cladbrain/cladbench)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T10:44:09Z

## Translation

タイトル: CladBench – 英国の建築規制に関する AI のオープン ベンチマーク
記事のタイトル: GitHub - cladbrain/cladbench: 英国および EU 構築環境における大規模言語モデルのオープン評価ベンチマーク — 12 のカテゴリーにわたる 536 の質問。 · GitHub
説明: 英国および EU の構築環境における大規模言語モデルのオープン評価ベンチマーク — 12 のカテゴリにわたる 536 の質問。 - クラッドブレイン/クラッドベンチ

記事本文:
GitHub - cladbrain/cladbench: 英国および EU の構築環境における大規模言語モデルのオープン評価ベンチマーク — 12 のカテゴリにわたる 536 の質問。 · GitHub
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
クラドブレイン
/
クラッドベンチ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット cladbench-hf cladbench-hf ドキュメント ドキュメント ペーパー ペーパー r

結果 結果 src/ cladbench src/ cladbench .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml required.txt requirements.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
英国および EU の構築環境における大規模言語モデルのオープン評価ベンチマーク。
12 のカテゴリーにわたる 536 の質問があり、すべての模範解答とスコアが公開されています。
📄 論文を読む — 12 ページ、PDF
測量士、エネルギー評価士、改修コーディネーター、持続可能性コンサルタントは、
実際の作業に言語モデルを使用する: 改造後の EPC バンドの推定、チェック
BREEAM クレジットが達成可能かどうか、MEES の期限に必要なものを検討する
特定の建物。一般的なベンチマークでは、モデルがどのような点で優れているかどうかを測定することはできません。
それ。クラッドベンチはそうです。
対象範囲は主に英国の慣行であり、厳選された EU の規制内容も含まれます。
7 つのモデル。すべての解答はルーブリックで採点され、本人ではない 2 人の審査員によって採点されます。
評価済み (DeepSeek Chat および Grok 4)。 10,000 ブートストラップからの 95% 信頼区間
再サンプル。
このベンチマークでは、Gemini 2.5 Pro と GPT-5 は区別されません。彼らの違いは、
+0.009 (95% 間隔は [-0.022, +0.039])。むしろ比較可能なものとして説明してください
ランキングするよりも。
全体の数値は、ここでは最も興味深い結果ではありません。すべてのモデルのスコアは少なくとも 0.75
BMS 異常分類について。規制の崖っぷちの問題についても、同じ 7 つの範囲にまたがる
0.901～0.126。タスクがどのカテゴリに分類されるかは、どのモデルがタスクを実行するかよりも重要です。
使用する前に回答キーを読んでください
参照の回答の強さは均一ではなく、データセットにはどれがどれであるかが示されています。
相互検証は、より弱い根拠です。カテゴリ 1 では、55 問すべてに合格し、そのうち 45 問が合格しました。
どの

承認された文書が実際に開かれた後に修正が必要でした。重み付けしてください
したがって、厳密なサブセットが必要な場合は、primary_source_verified を優先します。ただし、注意してください。
そのサブセットは 4 つのカテゴリに集中しており、12 つのカテゴリのバランスが取れたサンプルではありません。
55 の質問は、変更される可能性のある規制上の立場に依存します (metadata.policy_dependency: live )。
これらは 2026 年 8 月 13 日時点で有効です。再利用する前に再確認してください。
git clone https://github.com/cladbrain/cladbench
CDクラッドベンチ
pip install -e 。
API キーを使用せずに公開されたスコアを複製する
すべてのモデル応答がリリースされるため、評価済みモデルにアクセスする必要はありません。
数字を確認してください:
python -m cladbench スコア --input results/responses/full_opus47.jsonl
これにより、536 行のうち 243 行が再計算されます (すべての行は決定論的方法でグレーディングされます)。
保存されている 243 のスコアをすべて正確に再現します。残りの 293 はルーブリックで採点されます。追加する
--judge anthropic (および ANTHROPIC_API_KEY ) を使用して、それらも再判断します。
echo " ANTHROPIC_API_KEY=<your-key> " >> .env # 必要なプロバイダー
python -m cladbench 評価 --model anthropic:claude-opus-4-7 \
--split public --output my_run.jsonl
モデル仕様は、プロバイダー:モデルの形式をとり、プロバイダーは anthropic 、 openai 、
グーグル、一緒にそしてhf。実行ごとに、データセットのハッシュを記録するマニフェストが書き込まれます。
ハーネスのバージョンとトークンの使用状況。
数字を引用する前にこれらをお読みください。
質問はモデルによって生成され、一次情報源と照合されます。
名前。答えの鍵は検証でサポートされるものです。フレージングは単一のものを反映しています
発電機。
暫定的に 2 つのカテゴリーが採点されます。カテゴリ 2 では、50 件中 2 件の回答が検証されました
一次情報源に対して、カテゴリ 8 には 9/50 が含まれます。CIBSE 資料はライセンスが付与されており、
合法的なコピーが入手されたため、回答のうち 29 件はまだ公開されていません。

ソースと照合してチェックされます。
オプションの長さのキューが存在します。正しい選択肢は 37% の中で最も長くなります。
多肢選択問題の確率は 25% 近く、カテゴリー 6 では 69% です。
カテゴリ 6 の結果の一部は、BREEAM の知識ではなく、その手がかりである可能性があります。
2 つのオープンウェイトの結果は、あるプロバイダーのモデルではなく、あるプロバイダーの FP8 エンドポイントを説明します。
完全な精度。
デコードは均一ではありません。 GPT-5 および Claude Opus 4.x ファミリは、
温度が固定されているため、7 つのモデルのうち 3 つとすべての裁判官の評決がむしろサンプリングされます。
貪欲よりも。繰り返し実行の分散が測定され、論文で報告されます。
テキストのみ。図面、BIM モデル、PDF レポートはテストされていません。
カテゴリごとの n は 30–55 であるため、単一カテゴリの比較は参考値です。
このリポジトリにないもの
120 の質問からなる非公開のホールドアウトが存在しますが、意図的に公開されていません。今まで一度もなかった
任意のモデルに送信され、パブリック セットに関する汚染に関する質問ができるように存在します。
議論するのではなく、実行することで解決します。そのジェネレーターも保留されています。
シードジェネレーターが質問です。
src/cladbench/ パッケージ: CLI、スコアラー、アダプター、schema.json
src/cladbench/data/ 536 の質問、カテゴリごとに 1 つの JSONL
結果/応答/すべてのモデルの応答とスコア、モデルごとに 1 つのファイル
結果/ジャッジ/回答ごとの 2 人の中立ジャッジの採点
結果/マニフェスト/データセット ハッシュ、ハーネス バージョン、および実行ごとのトークン使用量
ドキュメント/スキーマリファレンス、カテゴリ仕様、元帳、結果テーブル
論文/PDF、Markdown、LaTeX としての論文
引用
セルバン、R.T. (2026)。 CladBench v1: 大規模言語モデルの 12 カテゴリーのベンチマーク
英国とEUの建築環境について。 https://doi.org/10.5281/zenodo.21951911
その DOI は常に最新バージョンに解決されます。この特定のリリースを引用するには、次を使用します
10.5281/ゼノド.219519

１２．機械可読メタデータは CITATION.cff にあります。
質問は承認された文書 (クラウン著作権、オープンガバメントライセンス)、BREEAM UK を引用しています。
New Construction 2018 (SD5078、BRE 著作権)、IFC4 EXPRESS スキーマ (buildingSMART)、および
CIBSE ガイダンス。これらは、それらの情報源を参照するオリジナルのテキストであり、複製ではありません。
彼ら: 536 の質問すべてにおいて、ソース文書と共有される最長の文章は 22 です。
単語であり、その長さの一致のほとんどは出版物のタイトルです。
回答キーに誤りを見つけた場合は、問題を開いてください。での検証履歴は、
論文が存在するのは、エラーが発見され修正されたためです。そのプロセスは止まらない
出版物。
英国および EU の構築環境における大規模言語モデルのオープン評価ベンチマーク — 12 のカテゴリにわたる 536 の質問。
Readme Apache-2.0 ライセンス
貢献 このリポジトリを引用する アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An open evaluation benchmark for large language models on the UK and EU built environment — 536 questions across twelve categories. - cladbrain/cladbench

GitHub - cladbrain/cladbench: An open evaluation benchmark for large language models on the UK and EU built environment — 536 questions across twelve categories. · GitHub
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
cladbrain
/
cladbench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits cladbench-hf cladbench-hf docs docs paper paper results results src/ cladbench src/ cladbench .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
An open evaluation benchmark for large language models on the UK and EU built environment.
536 questions across twelve categories, with every model response and score released.
📄 Read the paper — 12 pages, PDF
Surveyors, energy assessors, retrofit coordinators and sustainability consultants are
using language models for real work: estimating an EPC band after a retrofit, checking
whether a BREEAM credit is achievable, working out what a MEES deadline requires of a
particular building. General benchmarks do not measure whether a model is any good at
that. CladBench does.
Coverage is primarily UK practice, with selected EU regulatory content.
Seven models, all rubric-graded answers marked by two judges that are not themselves
evaluated (DeepSeek Chat and Grok 4). 95% confidence intervals from 10,000 bootstrap
resamples.
Gemini 2.5 Pro and GPT-5 are not separated by this benchmark. Their difference is
+0.009 with a 95% interval of [−0.022, +0.039]. Please describe them as comparable rather
than ranking them.
The overall number is the least interesting result here. Every model scores at least 0.75
on BMS anomaly classification; on regulatory cliff-edge questions the same seven span
0.901 to 0.126. Which category a task falls in matters more than which model runs it.
Read the answer key before you use it
Reference answers are not uniform in strength, and the dataset says which is which:
Cross-validation is the weaker warrant. In Category 1 it passed all 55 questions, 45 of
which needed correction once the Approved Documents were actually opened. Weight it
accordingly, and prefer primary_source_verified if you need a strict subset — but note
that subset is concentrated in four categories and is not a balanced sample of the twelve.
Fifty-five questions depend on regulatory positions that can move ( metadata.policy_dependency: live ).
They are valid as at 13 August 2026. Re-verify before reusing them.
git clone https://github.com/cladbrain/cladbench
cd cladbench
pip install -e .
Reproduce the published scores without an API key
Every model response is released, so you do not need access to any evaluated model to
check the numbers:
python -m cladbench score --input results/responses/full_opus47.jsonl
This recomputes 243 of the 536 rows — every row graded by a deterministic method — and
reproduces all 243 stored scores exactly. The remaining 293 are rubric-graded; add
--judge anthropic (and an ANTHROPIC_API_KEY ) to re-judge those too.
echo " ANTHROPIC_API_KEY=<your-key> " >> .env # whichever providers you need
python -m cladbench evaluate --model anthropic:claude-opus-4-7 \
--split public --output my_run.jsonl
Model specs take the form provider:model , with providers anthropic , openai ,
google , together and hf . Every run writes a manifest recording the dataset hash,
harness version and token usage.
Please read these before quoting a number.
The questions were model-generated , then checked against the primary sources they
name. The answer key is what the verification supports; the phrasing reflects a single
generator.
Two categories are provisionally scored. Category 2 has 2 of 50 answers verified
against a primary source, Category 8 has 9 of 50. CIBSE material is licensed and no
lawful copy was obtained, so 29 of those answers have not been checked against a source.
An option-length cue is present. The correct option is the longest in 37% of
multiple-choice questions against a chance rate near 25%, and in 69% of Category 6.
Part of the Category 6 result may be that cue rather than BREEAM knowledge.
Two open-weights results describe FP8 endpoints at one provider, not the models at
full precision.
Decoding is not uniform. GPT-5 and the Claude Opus 4.x family do not accept a
pinned temperature, so three of seven models and every judge verdict are sampled rather
than greedy. Repeat-run variance is measured and reported in the paper.
Text only. Drawings, BIM models and PDF reports are not tested.
Per-category n is 30–55 , so single-category comparisons are indicative.
What is not in this repository
A 120-question private holdout exists and is deliberately not published. It has never been
sent to any model, and exists so that a contamination question about the public set can be
settled by running it rather than argued about. Its generators are withheld too, because a
seeded generator is the questions.
src/cladbench/ the package: CLI, scorers, adapters, schema.json
src/cladbench/data/ the 536 questions, one JSONL per category
results/responses/ every model response and score, one file per model
results/judges/ the two neutral judges' marks, per answer
results/manifests/ dataset hash, harness version and token usage per run
docs/ schema reference, category specifications, ledger, result tables
paper/ the paper as PDF, Markdown and LaTeX
Citing
Selvan, R. T. (2026). CladBench v1: A Twelve-Category Benchmark for Large Language Models
on the UK and EU Built Environment. https://doi.org/10.5281/zenodo.21951911
That DOI always resolves to the latest version. To cite this specific release, use
10.5281/zenodo.21951912 . Machine-readable metadata is in CITATION.cff .
Questions cite Approved Documents (Crown copyright, Open Government Licence), BREEAM UK
New Construction 2018 (SD5078, BRE copyright), the IFC4 EXPRESS schema (buildingSMART) and
CIBSE guidance. They are original text referencing those sources, not reproductions of
them: across all 536 questions the longest passage shared with any source document is 22
words, and most of the matches at that length are publication titles.
If you find an error in the answer key, please open an issue. The verification history in
the paper exists because errors were found and fixed; that process does not stop at
publication.
An open evaluation benchmark for large language models on the UK and EU built environment — 536 questions across twelve categories.
Readme Apache-2.0 license Contributing
Contributing Cite this repository Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
