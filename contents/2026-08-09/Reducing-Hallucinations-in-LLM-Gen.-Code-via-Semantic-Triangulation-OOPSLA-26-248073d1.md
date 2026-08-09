---
source: "https://github.com/msv-lab/just-tri-it"
hn_url: "https://news.ycombinator.com/item?id=49230199"
title: "Reducing Hallucinations in LLM-Gen. Code via Semantic Triangulation (OOPSLA 26)"
article_title: "GitHub - msv-lab/just-tri-it: Reducing Hallucinations in LLM-Generated Code via Semantic Triangulation (OOPSLA'26) · GitHub"
author: "mechtaev"
captured_at: "2026-08-09T11:21:18Z"
capture_tool: "hn-digest"
hn_id: 49230199
score: 2
comments: 0
posted_at: "2026-08-09T10:42:39Z"
tags:
  - hacker-news
  - translated
---

# Reducing Hallucinations in LLM-Gen. Code via Semantic Triangulation (OOPSLA 26)

- HN: [49230199](https://news.ycombinator.com/item?id=49230199)
- Source: [github.com](https://github.com/msv-lab/just-tri-it)
- Score: 2
- Comments: 0
- Posted: 2026-08-09T10:42:39Z

## Translation

タイトル: LLM 世代における幻覚の軽減セマンティック三角形分割によるコード (OOPSLA 26)
記事のタイトル: GitHub - msv-lab/just-tri-it: セマンティック三角形分割による LLM 生成コードの幻覚の軽減 (OOPSLA'26) · GitHub
説明: セマンティック三角形分割による LLM 生成コードの幻覚の低減 (OOPSLA'26) - msv-lab/just-tri-it

記事本文:
GitHub - msv-lab/just-tri-it: セマンティック三角形分割による LLM 生成コードの幻覚の軽減 (OOPSLA'26) · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
msv-lab
/
ジャストトライイット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
415 コミット 415 コミット データセット データセット ドキュメント ドキュメント 証明 証明 スクリプト スクリプト src/ just_tri_it src/ just_tri_i

t テスト テスト .gitignore .gitignore .python-version .python-version README.md README.md pyproject.toml pyproject.toml test_requirements.txt test_requirements.txt uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM で生成されたコードにおける幻覚の軽減
セマンティック三角測量経由
ダイ・イーハン、リャン・シジエ、シュー・ハオティアン、謝ペイチュー、セルゲイ・メフタエフ
arXiv:2511.12288
LLM で生成されたコードには幻覚的なバグが含まれることが多く、予期される動作が正式に指定されることはほとんどないため、自動的に検出するのは困難です。サンプリングされたプログラムのうちどれが正しいかを特定することは、警察の刑事が容疑者に尋問するのと似ています。 LLM は相関関係のある誤りを犯すため、ほとんどの容疑者が同じ偽のアリバイで共謀しているため、複数 (多数決) の投票では真実が特定されません。それは彼らが共有する欺瞞を増幅させるだけです。
以前の方法では、LLM によって生成されたテスト、または問題の説明から自動的に形式化された仕様 (例: Hoare スタイルの事後条件) など、追加の証人が必要になります。しかし、これらの証人は同じLLMによって作成され、容疑者の欠陥のある論理を共有しています。偏った証人は指を組んで宣誓し、偽りのアリバイを裏付けます。
ジャスト・トライイットは検査官を演じ、信頼できる証人を得るために予期せぬ角度から問題を尋問することで嘘を暴きます。これはセマンティック三角測量と呼ばれます。
意味的三角形分割 (τ, φ) は、解離的な問題変換 τ と、プログラムの意味的同値クラス間の全単射を誘発するプログラムのペアに関する関係 (超プロパティ) φ で構成され、問題 d の正しい解を変換された問題 τ( d ) の正しい解にマッピングします。
調査中の問題 d を考慮すると、検査官は予期せぬ調査項目 d' = τ( d ) を開始し、bo に対する解決策をサンプルします。

th は独立して問題を生成し、φ でクロステストします。
各要件は、尋問において独自の役割を果たします。
解離性: τ( d ) を解くには根本的に異なるアルゴリズムが必要であるため、容疑者が耐える単なる言い換えとは異なり、質問は容疑者のリハーサルされたアリバイとはまったく無関係の角度から出てきます。実装された変換には、部分反転 (入力と出力の交換)、回答の列挙 (有効な回答をすべて出力)、および問題の分解が含まれます。
全単射誘発 : d の解における明確なエラーは、 d' の解における明確なエラーにマッピングされるため、容疑者の話の微妙な矛盾さえも検出されます。
正しさ結合 : 全単射は正しい解決策を正しい解決策にマッピングするため、真実の説明は真実の証人によって裏付けられますが、2 つの独立した嘘が偶然に一致することはほとんどありません。
正しい番組のストーリーは、どのような質問にも耐えられます。幻覚プログラムは、不意を突かれた嘘つきと同じように、矛盾によって自らを裏切ります。私たちの LLM 幻覚の数学的モデルに基づいて、そのような三角測量の目撃者との一致は、複数投票よりも厳密に高い正しさの信頼を生み出すことを証明します。選択されたプログラムは、偽の統計的相関ではなく、正確な一般化を反映しています。 CodeElo および LiveCodeBench の問題の詳細な図、完全な理論、および評価は、論文に記載されています。
環境変数 AI302_API_KEY を介して 302.ai API キーを設定します。依存関係は uv によって管理されます。
UVランラフチェック。
uvx mypy src/
UV で pytest を実行
使用法
コード生成データセットのツール構成を比較し、基本的な統計を計算します。
uv 実行ベンチマーク --dataset DATASET [--task TASK_ID] --selector TOOL_CONFIG --model MODEL
たとえば:
uv 実行ベンチマーク --dataset データセット/tes

t.json --selector 複数形 --model gpt-4o
LiveCodeBench v6 の場合は、まずデータセットを解凍し ( unzip datasets/lcb_part6.json.zip )、次に次の操作を行います。
uv 実行ベンチマーク --dataset datasets/lcb_part6.json --selector CodeT_IO --model gpt-4o --task atcoder_abc387_b
利用可能な構成: Plurality 、 MaxTest_Assert 、 MaxTest_IO 、 CodeT_Assert 、 CodeT_IO 、 Syntactic 、 OffByOne 、 Postcondition 、 FWD_INV 、 FWD_SINV 。各三角測量スキームの呼び出し例は doc/Examples.md にあります。
包括的な測定値を収集します ( data_dir に追加されます)。
uv 実行実験 --dataset datasets/lcb_part6.json --model gpt-4o --data data_dir
特定のタスクを実行するには、--only atcoder_abc387_b を使用します。次に、メジャーを計算してプロットを生成します。
UV 実行分析 --data data_dir --report report_dir
分離されたテストの実行
生成されたプログラムを分離されたサブプロセスで実行するには、専用の環境を作成します。
uv venv --no-project --seed --python 3.13 test_venv
ソース test_venv/bin/activate
pip install -r test_requirements.txt
非アクティブ化する
次に、上記のコマンドに --test-venv test_venv/ を追加します。
--cache-root DIR — LLM キャッシュを設定します (デフォルト: ~/.just_tri_it_cache/ )
--export-cache DIR — 実行中に使用されたすべてのキャッシュされたサンプルを別のディレクトリにエクスポートします
--replicate — キャッシュのみを使用します。キャッシュミスで失敗する
このリポジトリのコミット ハッシュとそのルートから実行される bash コマンドがあれば、実験は再現可能です。これを複製するには、LLM キャッシュのコミット ハッシュを追加して提供します。キャッシュはからダウンロードできます
https://github.com/msv-lab/just-tri-it-cache-USER/archive/COMMIT.zip
ここで、 USER は yihan 、 haotian 、 sijie 、 sergey のいずれかです。 LLM キャッシュから自明に導出できないものはすべて just-tri-it-data に保存されます。
セマンティック三角形分割による LLM 生成コードの幻覚の低減 (OOPSLA'26)
arxiv.org/ab

s/2511.12288 リソース
Readme アクティビティ カスタム プロパティ スター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Reducing Hallucinations in LLM-Generated Code via Semantic Triangulation (OOPSLA'26) - msv-lab/just-tri-it

GitHub - msv-lab/just-tri-it: Reducing Hallucinations in LLM-Generated Code via Semantic Triangulation (OOPSLA'26) · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
msv-lab
/
just-tri-it
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
415 Commits 415 Commits datasets datasets doc doc proofs proofs scripts scripts src/ just_tri_it src/ just_tri_it tests tests .gitignore .gitignore .python-version .python-version README.md README.md pyproject.toml pyproject.toml test_requirements.txt test_requirements.txt uv.lock uv.lock View all files Repository files navigation
Reducing Hallucinations in LLM-Generated Code
via Semantic Triangulation
Yihan Dai, Sijie Liang, Haotian Xu, Peichu Xie, Sergey Mechtaev
arXiv:2511.12288
LLM-generated code often contains hallucinated bugs, and since expected behavior is rarely formally specified, they are hard to detect automatically. Identifying which, if any, of the sampled programs are correct is akin to a police detective questioning suspects. Because LLMs make correlated errors , most suspects have colluded on the same fake alibi — so plurality (majority) voting does not identify the truth; it merely amplifies their shared deception.
Previous methods bring in extra witnesses: LLM-generated tests, or specifications auto-formalized from the problem description (e.g., Hoare-style postconditions). But these witnesses are produced by the same LLM and share the suspects' flawed logic — a biased witness who swears an oath with crossed fingers and corroborates the false alibi:
just-tri-it plays the inspector, exposing lies by questioning the problem from an unexpected angle to obtain a reliable witness . This is called semantic triangulation :
A semantic triangulation (τ, φ) consists of a dissociative problem transformation τ and a relation over pairs of programs (a hyperproperty) φ that induces a bijection between semantic equivalence classes of programs, mapping correct solutions of a problem d to correct solutions of the transformed problem τ( d ).
Given the problem d under investigation, the inspector opens an unexpected line of inquiry d′ = τ( d ), samples solutions to both problems independently, and cross-examines them with φ:
Each requirement plays its own role in the interrogation:
Dissociative : solving τ( d ) requires a fundamentally different algorithm, so the question comes from an angle wholly unrelated to the suspects' rehearsed alibi — unlike mere paraphrasing, which they withstand. Implemented transformations include partial inversion (swap input and output), answer enumeration (output all valid answers), and problem decomposition.
Bijection-inducing : distinct errors in solutions to d map to distinct errors in solutions to d′ , so even subtle inconsistencies in the suspects' stories are detected.
Correctness-coupling : the bijection maps correct solutions to correct solutions, so a truthful account is corroborated by a truthful witness, while two independent lies rarely match by coincidence.
A correct program's story holds up under any line of questioning; hallucinated programs, like liars caught off guard, betray themselves through contradiction. Under our mathematical model of LLM hallucinations, we prove that agreement with such a triangulated witness yields strictly higher confidence of correctness than plurality voting — the selected program reflects accurate generalization rather than spurious statistical correlations. Detailed illustrations on CodeElo and LiveCodeBench problems, the full theory, and the evaluation are in the paper .
Set your 302.ai API key via the environment variable AI302_API_KEY . Dependencies are managed by uv .
uv run ruff check .
uvx mypy src/
uv run pytest
Usage
Compare tool configurations on code generation datasets and compute basic statistics:
uv run benchmark --dataset DATASET [--task TASK_ID] --selector TOOL_CONFIG --model MODEL
For example:
uv run benchmark --dataset datasets/test.json --selector Plurality --model gpt-4o
For LiveCodeBench v6, first decompress the dataset ( unzip datasets/lcb_part6.json.zip ), then:
uv run benchmark --dataset datasets/lcb_part6.json --selector CodeT_IO --model gpt-4o --task atcoder_abc387_b
Available configurations: Plurality , MaxTest_Assert , MaxTest_IO , CodeT_Assert , CodeT_IO , Syntactic , OffByOne , Postcondition , FWD_INV , FWD_SINV . Example invocations for each triangulation scheme are in doc/Examples.md .
Collect comprehensive measurements (appended into data_dir ):
uv run experiment --dataset datasets/lcb_part6.json --model gpt-4o --data data_dir
Use --only atcoder_abc387_b to run a specific task. Then compute measures and generate plots:
uv run analyze --data data_dir --report report_dir
Isolated test execution
To execute generated programs in isolated subprocesses, create a dedicated environment:
uv venv --no-project --seed --python 3.13 test_venv
source test_venv/bin/activate
pip install -r test_requirements.txt
deactivate
Then add --test-venv test_venv/ to the above commands.
--cache-root DIR — set LLM cache (default: ~/.just_tri_it_cache/ )
--export-cache DIR — export all cached samples used during the run to a different directory
--replicate — use only cache; fail on cache misses
An experiment is reproducible given a commit hash of this repository and a bash command executed from its root; to replicate it, additionally provide the commit hash of your LLM cache. Caches can be downloaded from
https://github.com/msv-lab/just-tri-it-cache-USER/archive/COMMIT.zip
where USER is one of yihan , haotian , sijie , sergey . Everything not trivially derivable from LLM cache is stored in just-tri-it-data .
Reducing Hallucinations in LLM-Generated Code via Semantic Triangulation (OOPSLA'26)
arxiv.org/abs/2511.12288 Resources
Readme Activity Custom properties Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
