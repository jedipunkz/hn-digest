---
source: "https://github.com/SupermodularAI/local_bench"
hn_url: "https://news.ycombinator.com/item?id=49208099"
title: "Show HN: Benchmark local LLMs fit for your device specs"
article_title: "GitHub - SupermodularAI/local_bench: Benchmark the local LLMs you already have — speed, memory, energy, and quality — as a live terminal leaderboard · GitHub"
author: "andrelago"
captured_at: "2026-08-07T10:42:56Z"
capture_tool: "hn-digest"
hn_id: 49208099
score: 1
comments: 0
posted_at: "2026-08-07T09:57:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Benchmark local LLMs fit for your device specs

- HN: [49208099](https://news.ycombinator.com/item?id=49208099)
- Source: [github.com](https://github.com/SupermodularAI/local_bench)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T09:57:40Z

## Translation

タイトル: HN の表示: デバイスの仕様に合わせたローカル LLM のベンチマーク
記事のタイトル: GitHub - SupermodularAI/local_bench: ライブ ターミナル リーダーボードとして、すでに所有しているローカル LLM (速度、メモリ、エネルギー、品質) をベンチマークする · GitHub
説明: すでに所有しているローカル LLM (速度、メモリ、エネルギー、品質) をライブ ターミナル リーダーボードとしてベンチマークします - SupermodularAI/local_bench
HN テキスト: Ollama および同様のプロバイダーを介して、さまざまなメトリクスでカスタム セットのタスクを使用してローカル LLM をベンチマークします。決定論的な評価基準またはカスタム命令による LLM ジャッジを使用できます。これにより、HuggingFace へのクエリもサポートされ、トレンドのモデルとデバイスの仕様を比較できるため、セットアップでどのモデルをテストする価値があるかがわかります。

記事本文:
GitHub - SupermodularAI/local_bench: すでに所有しているローカル LLM (速度、メモリ、エネルギー、品質) をライブ ターミナル リーダーボードとしてベンチマークします · GitHub
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
スーパーモジュラーAI
/
ローカルベンチ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
45 コミット 45 コミット .github/ workflows .github/ workflows example

s 例 src/ local_bench src/ local_bench テスト テスト tools/ Pack_viewer tools/ Pack_viewer .gitignore .gitignore .gitlab-ci.yml .gitlab-ci.yml .pre-commit-config.yaml .pre-commit-config.yaml ライセンス ライセンス MANIFEST.in MANIFEST.in NOTICE NOTICE README.md README.md RELEASING.md RELEASING.md conftest.py conftest.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すでに所有しているローカル LLM (速度、メモリ、エネルギー、品質) をライブ ターミナル リーダーボードとしてベンチマークします。
local_bench は、ローカル ランナー ( Ollama 、 LM Studio 、 llama.cpp 、 vLLM 、または任意の OpenAI 互換サーバー) にインストールされているモデルを検出し、厳選された品質スイートを実行し、トークン/秒、最初のトークンまでの時間、メモリ フットプリント、およびオプションで実際のマシン上のトークンごとのジュールを測定し、ライブ比較リーダーボードをレンダリングする単一コマンド TUI です。
git clone https://github.com/SupermodularAI/local_bench.git
cd ローカルベンチ
pip インストール 。
ローカルベンチ
それだけです。構成も API キーもクラウドもありません。
この問題の半分を解決する優れたツールはありますが、両方を解決するローカルファーストのツールはありません。
llama-bench (llama.cpp 内) は速度のみを測定します。
lm-evaluation-harness は品質を測定しますが、洗練されたラップトップ UX がなく、ほとんどの人が実際にローカルで使用するモデルランナーを中心に構築されていません。
local_bench は、ローカルファースト、ゼロ構成、UX 主導のギャップを埋めます。クローンを作成して実行し、既に取得したモデルを指定すると、「ローカル モデルのどれが実際に優れているか、そしてこのラップトップではどれくらい速いか?」に対する一目で答えが得られます。
メトリック
どのように
トク/秒
出力トークン ÷ 生成時間。 Ollama はサーバー側の評価タイミングを報告します。 OpenAI 互換のバックエンドは、トークン ストリームからクライアント側でタイミングが測定されます。プロンプト処理とモデルのロードは除きます。
TTFT
最初にストリーミングされたトークンまでの実時間 (分

ランナーが報告するモデルの読み込み時間)。
記憶
ランナーが公開するときの常駐モデルのサイズ (Ollama /api/ps 、LM Studio /api/v0 )、およびバックエンドのプロセスのベストエフォートのピーク RSS サンプル。
エネルギー
--energy を使用した出力トークンあたりのジュール。速度プローブ中にプラットフォーム パワー メーター (macOS powermetrics 、Linux RAPL) からサンプリングされ、実行ごとに 1 回測定されるアイドル ベースラインが差し引かれます。 root が必要です — Energy を参照してください。
品質
数学、推論、事実の想起、指示に従う/構造化された出力、抽出、コード理解にわたる、決定的に等級付けされた 31 のタスク。オプションの LLM-as-judge では、自由形式のタスク (要約、電子メール、俳句、説明) を追加します。
インストール
パッケージ インデックスに公開されていません - クローンからインストールします。
git clone https://github.com/SupermodularAI/local_bench.git
cd ローカルベンチ
pip インストール 。 # 次に実行: local_bench
隔離しておきたいですか?最初に virtualenv を作成するか、次のようにインストールします。
同じディレクトリから pipx を取得します。
pipx をインストールします。
オプションの追加機能 (すべてクローンから同じ方法でインストールされます):
local_bench # 高速デフォルト: 3 つの最小モデル、クイック スイート (TUI)
local_bench --all # 検出されたすべてのモデルをベンチマークします
local_bench --full # 完全な品質のスイート (高速サブセットだけでなく) を実行します
local_bench --no-tui # プレーンなライブ レンダラー (パイプ/CI に最適)
local_bench -m llama3.2,qwen3:8b # これらのモデルのみ
local_bench --limit 3 # モデルの数を制限します
local_bench --provider lmstudio # 自動検出の代わりに LM Studio を使用します
local_bench --provider llamacpp # llama.cpp サーバー (llama-server)
local_bench --provider vllm # vLLM
local_bench --provider openai --host http://localhost:5000 # OpenAI 互換サーバー
local_bench --refresh-cache # キャッシュされた応答を再利用する代わりに再計算します
local_bench --energy # トークンあたりのジュールも測定します (「エネルギー」を参照)
ローカルベン

ch --no-quality # 速度 + メモリのみ (高速)
local_bench --no-speed # 品質のみ
local_bench --judge qwen3:8b # LLM-as-judge を有効にする (オープンエンドタスクを追加)
local_bench --judge-provider openai --judge-host https://gateway # ジャッジオフボックス
local_bench --tasks mypack.yaml # 組み込みスイートの代わりにカスタム タスク パックを使用します
local_bench --only math # 1 つのカテゴリを実行するか、タスク ID に名前を付けます
local_bench --num-ctx 16384 # コンテキスト ウィンドウを表示します (長いプロンプトに必要)
local_bench --add-tasks mypack.yaml # 組み込みスイートの上にパックを追加します
local_bench --label " チューニング前 " # 後で比較するためにこの実行にタグを付けます
local_bench --md results.md # Markdown レポートもエクスポートします
local_bench --json results.json # 生の JSON もエクスポートします
local_bench list # 検出されたモデルをリストするだけ
local_bench タスク # 品質スイートを表示します (パックをプレビューするには --tasks を追加します)
local_bench History # 過去の実行のリスト (自動的に保存)
local_bench diff # 最近の 2 つの実行の差分
local_bench diff 3 1 # 実行 #3 (ベース) と実行 #1 (新しい) の差分
local_bench スループット # バッチ スループット スイープ (同時実行 1、2、4、8)
local_bench スループット --concurrency 1,8,16 --provider vllm
local_bench fit # どの人気モデルがあなたのハードウェアに適合しますか?
ローカルで実行_
[切り捨てられた]
Ollama 経由で Apple M4 上で実行される実際のクイックスイート:
最終リーダーボード
┏━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━┳━━━━━━┳━━━━━━━┳━━━━━━━━┳━━━━━━━━━━┳━━━━━━━━┓
┃ # �� モデル �� パラメータ �� 品質 �� パス �� tok/s �� TTFT �� 平均タスク �� メモリ �� 
┡━━━╇━━━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━╇━━━━━━╇━

━━━━━━╇━━━━━━━━╇━━━━━━━━━━╇━━━━━━━━┩
│ 1 │ gemma4:最新 │ 8.0B │ 100% │ 8/8 │ 27.7 │ 205 ミリ秒 │ 1.29 秒 │ 8.8 GB │
│ 2 │ gemma3:最新 │ 4.3B │ 62% │ 5/8 │ 34.6 │ 160 ミリ秒 │ 1.23 秒 │ 3.5 GB │
━───┴─────────┴───┴─────┴─────┴───────┴───┴─────┴───┘
平均タスクは、1 つの高品質タスクを完了するまでの平均実時間です。カバーします
実際に生成されたタスクのみ — キャッシュされた応答は機能しないため、
完全にキャッシュされた実行時間は、測定されたことのないゼロに近い時間ではなく表示されます。
--energy を使用すると、J/tok 列がボードに追加されます。
(数値はその時点でのラップトップのものです。「制限事項」を参照してください。)
少なくとも 1 つのローカル モデル ランナーが到達可能である必要があります。
自動検出では、Ollama → LM Studio → llama.cpp → vLLM が試行されます (汎用の openai プロバイダーは明示的専用です)。 --provider を使用して強制的に 1 つを指定します。 --host または上記の環境変数でホストをオーバーライドします。
このスイートは意図的に小さくなっています。モデルを分割するのに十分な数のカテゴリにわたるタスクがあり、ラップトップですべてのモデルを数分で実行できるほど少ないものです。各タスクは決定的に評価されます (数値の正確な一致、複数選択の文字、部分文字列、有効な JSON、正規表現)。温度は 0 で、再現性を高めるために固定シードが使用されます。リストについては、local_bench タスクを参照してください。
オプションの --judge MODEL フラグは、参照回答に対してオープンエンド タスク 1 ～ 5 のスコアを付ける LLM-as-judge (任意のローカル モデル) をオンにします。それは神託ではなく合図だ。
フルスイートのすべてのモデルのベンチマークをラップトップで行うには時間がかかるため、デフォルトは一目でわかるように調整されています。
3秒

デフォルトで最も小さいモデル (最初に小さいため、結果が速く表示されます) — すべての場合は --all、選択する場合は -m。
高速品質のサブセット (すべてのカテゴリにわたって最大 8 つのタスク) — --31 件すべてで完全。
応答のキャッシュ: 品質実行では温度 0 + 固定シードが使用されるため、応答は決定的であり、 ~/.local_bench の下にキャッシュされます。再実行では、新しいモデル/タスクのみが再生成されます (変更されていないものは、ミリ秒以内にキャッシュから再グレーディングされます)。 --refresh-cache は再計算を強制し、--no-cache は再計算を無効にします。
実際には、これにより最初の実行は約 15 ～ 25 分 (すべてのモデル、フル スイート) から約 1 ～ 2 分に短縮され、再実行は数秒に短縮されます。完全なパス (CI、最終数値) を行うには、 local_bench --all --full を使用します。
JSON または YAML パックを使用して独自の eval を持ち込みます。Python は必要ありません。 --tasks は組み込みスイートを置き換えます。 --add-tasks を追加します。 YAML にはオプションの追加 ( pip install ".[yaml]" ) が必要です。 JSON はそのまま使用できます。
# mypack.yaml — local_bench --tasks mypack.yaml
名前 : マイパック
タスク:
- id : Capital_japan
カテゴリー : 事実
プロンプト:「日本の首都はどこですか?都市名だけを答えてください。」
採点者 : {タイプ: contains_any、値: ["東京"]}
問い合わせ先：東京
- ID : 追加
カテゴリー : 数学
プロンプト: 「12 + 30 とは何ですか? 答えを 1 行で終了してください。」
採点者 : {タイプ: 正確な番号、値: 42}
- id : Explain # 採点者なし -> 自由形式、--judge でのみ採点
カテゴリー : オープン
プロンプト:「光合成を一文で説明してください。」
参考：「植物は太陽光、水、CO2をブドウ糖と酸素に変換します。」
Grader タイプの値:exact_number ( value 、 tol )、multiple_choice ( value )、contains_any ( value )、regex ( pattern 、ignorecase )、valid_json ( key )、valid_json_array ( length )。審査員のみのタスクの場合は採点者を省略します。 local_bench タスク --tasks mypack.yaml を使用してパックをプレビューします。
2 つの実行可能なパックが出荷される

Examples/ では、sample_pack.yaml は形式を調べる 4 つのタスクであり、extended_pa​​ck.yaml はすべてのカテゴリとすべての採点者のタイプをカバーする 21 のタスクのテンプレートです。これをコピーして、独自のプロンプトに置き換えます。
パックを誠実に保つ 2 つのルールは、両方とも出荷されたもののテストによって適用されます。
すべての決定論的タスクの参照は、独自のグレーダーを通過する必要があります。参考資料は、実際に実行された正解の例です。グレーダーをクリアできない場合は何もクリアできず、結果の 0 はオーサリングの問題ではなく、モデルの失敗のように見えます。
max_tokens にはヘッドルームが必要です。予算は回答全体をカバーします。最後の数字を読み取るexact_numberを使用すると、正しい答えに到達し、その後話し続けるモデルは文の途中で打ち切られ、最後に着いた数字でスコアが付けられます。最後に答えを聞いて、そこでやめてくださいと言いなさい。
パックを実行する前に確認する
local_bench タスク --tasks mypack.yaml はパックをターミナルに出力します。より大きなものとしては、小さなブラウザ UI である tools/pack_viewer もあります。これは、すべてのタスクを判断する正確なグレーダー仕様の横に表示し、モデルが実際にロードされるコンテキスト ウィンドウに対するプロンプト サイズをグラフで表示し、上記の両方のルールをパックに適用します。つまり、独自のグレーダーを使用して各リファレンスを再グレーディングし、コンテキストからオーバーフローするプロンプトにフラグを立てます。
pip install -r tools/pack_viewer/requirements.txt
streamlit 実行 tools/pack_viewer/app.py -- mypack.yaml
パッケージの一部ではなくスタンドアロンのユーティリティ: src/ には何もインポートされず、Streamlit は local_bench 自体の依存関係から外れます。
別のプロバイダーで判断する
デフォルトでは、ジャッジはテスト対象のモデルと同じプロバイダー上で実行されるため、プロバイダーが提供するモデルのみをジャッジできます。地元のランナーが 1 人いる場合、それは通常、審査員も候補者であることを意味し、その後、独自の得点を獲得します。

無制限のタスクに取り組みます。

[切り捨てられた]

## Original Extract

Benchmark the local LLMs you already have — speed, memory, energy, and quality — as a live terminal leaderboard - SupermodularAI/local_bench

Benchmark local LLMs with custom sets of tasks via Ollama and similar providers on a wide variety of metrics. You can use deterministic evaluation criteria or LLM judges with custom instructions. This also supports querying HuggingFace to compare trending models with your device's specs so you know which models might be worth testing in your setup.

GitHub - SupermodularAI/local_bench: Benchmark the local LLMs you already have — speed, memory, energy, and quality — as a live terminal leaderboard · GitHub
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
SupermodularAI
/
local_bench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
45 Commits 45 Commits .github/ workflows .github/ workflows examples examples src/ local_bench src/ local_bench tests tests tools/ pack_viewer tools/ pack_viewer .gitignore .gitignore .gitlab-ci.yml .gitlab-ci.yml .pre-commit-config.yaml .pre-commit-config.yaml LICENSE LICENSE MANIFEST.in MANIFEST.in NOTICE NOTICE README.md README.md RELEASING.md RELEASING.md conftest.py conftest.py pyproject.toml pyproject.toml View all files Repository files navigation
Benchmark the local LLMs you already have — speed, memory, energy, and quality — as a live terminal leaderboard.
local_bench is a single-command TUI that discovers the models installed in your local runner ( Ollama , LM Studio , llama.cpp , vLLM , or any OpenAI-compatible server), runs a curated quality suite, measures tokens/sec , time-to-first-token , memory footprint , and optionally joules per token on your actual machine , and renders a live comparison leaderboard.
git clone https://github.com/SupermodularAI/local_bench.git
cd local_bench
pip install .
local_bench
That's it. No config, no API keys, no cloud.
There are great tools for one half of this problem, but nothing local-first that does both:
llama-bench (inside llama.cpp) measures speed only .
lm-evaluation-harness measures quality but has no polished laptop UX and isn't built around the model runners most people actually use locally.
local_bench fills the gap: local-first, zero-config, UX-driven. Clone-and-run, point it at the models you already pulled, and get an at-a-glance answer to "which of my local models is actually good, and how fast is it on this laptop?"
Metric
How
tok/s
Output tokens ÷ generation time. Ollama reports server-side eval timing; OpenAI-compatible backends are timed client-side from the token stream. Excludes prompt processing and model load.
TTFT
Wall-clock time to the first streamed token (minus model-load time where the runner reports it).
Memory
Resident model size when the runner exposes it (Ollama /api/ps , LM Studio /api/v0 ), plus a best-effort peak-RSS sample of the backend's processes.
Energy
Joules per output token, with --energy . Sampled from the platform power meter (macOS powermetrics , Linux RAPL) during the speed probe, minus an idle baseline measured once per run. Needs root — see Energy .
Quality
31 deterministically-graded tasks across math, reasoning, factual recall, instruction-following/structured-output, extraction, and code understanding. Optional LLM-as-judge adds open-ended tasks (summaries, email, haiku, explanations).
Install
Not published to a package index — install from a clone:
git clone https://github.com/SupermodularAI/local_bench.git
cd local_bench
pip install . # then run: local_bench
Prefer to keep it isolated? Create a virtualenv first, or install with
pipx from the same directory:
pipx install .
Optional extras (all installed the same way, from the clone):
local_bench # fast default: 3 smallest models, quick suite (TUI)
local_bench --all # benchmark every discovered model
local_bench --full # run the full quality suite (not just the fast subset)
local_bench --no-tui # plain live renderer (great for piping / CI)
local_bench -m llama3.2,qwen3:8b # only these models
local_bench --limit 3 # cap the number of models
local_bench --provider lmstudio # use LM Studio instead of auto-detect
local_bench --provider llamacpp # llama.cpp server (llama-server)
local_bench --provider vllm # vLLM
local_bench --provider openai --host http://localhost:5000 # any OpenAI-compatible server
local_bench --refresh-cache # recompute instead of reusing cached responses
local_bench --energy # also measure joules per token (see Energy)
local_bench --no-quality # speed + memory only (fast)
local_bench --no-speed # quality only
local_bench --judge qwen3:8b # enable LLM-as-judge (adds open-ended tasks)
local_bench --judge-provider openai --judge-host https://gateway # judge off-box
local_bench --tasks mypack.yaml # use a custom task pack instead of the built-in suite
local_bench --only math # run one category, or name task ids
local_bench --num-ctx 16384 # raise the context window (needed for long prompts)
local_bench --add-tasks mypack.yaml # add a pack on top of the built-in suite
local_bench --label " before tuning " # tag this run for later diffing
local_bench --md results.md # also export a Markdown report
local_bench --json results.json # also export raw JSON
local_bench list # just list discovered models
local_bench tasks # show the quality suite (add --tasks to preview a pack)
local_bench history # list past runs (saved automatically)
local_bench diff # diff the two most recent runs
local_bench diff 3 1 # diff run #3 (base) against run #1 (newer)
local_bench throughput # batch-throughput sweep (concurrency 1,2,4,8)
local_bench throughput --concurrency 1,8,16 --provider vllm
local_bench fit # which popular models fit YOUR hardware?
Run local_
[truncated]
A real quick-suite run on an Apple M4, via Ollama:
Final leaderboard
┏━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━┳━━━━━━┳━━━━━━━┳━━━━━━━━┳━━━━━━━━━━┳━━━━━━━━┓
┃ # ┃ Model ┃ Params ┃ Quality ┃ Pass ┃ tok/s ┃ TTFT ┃ Avg task ┃ Memory ┃
┡━━━╇━━━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━╇━━━━━━╇━━━━━━━╇━━━━━━━━╇━━━━━━━━━━╇━━━━━━━━┩
│ 1 │ gemma4:latest │ 8.0B │ 100% │ 8/8 │ 27.7 │ 205 ms │ 1.29 s │ 8.8 GB │
│ 2 │ gemma3:latest │ 4.3B │ 62% │ 5/8 │ 34.6 │ 160 ms │ 1.23 s │ 3.5 GB │
└───┴───────────────┴────────┴─────────┴──────┴───────┴────────┴──────────┴────────┘
Avg task is the mean wall-clock time to complete one quality task. It covers
only tasks that were actually generated — a cached response does no work, so a
fully cached run shows – rather than a near-zero time it never measured.
With --energy , a J/tok column joins the board.
(Numbers are for that laptop at that moment — see Limitations .)
At least one local model runner must be reachable:
Auto-detection tries Ollama → LM Studio → llama.cpp → vLLM (the generic openai provider is explicit-only). Force one with --provider . Override host with --host or the env var above.
The suite is small on purpose — enough tasks across categories to separate models, few enough that every model runs in a couple of minutes on a laptop. Each task is graded deterministically (exact numeric match, multiple-choice letter, substring, valid-JSON, regex). Temperature is 0 and a fixed seed is used for reproducibility. See local_bench tasks for the list.
The optional --judge MODEL flag turns on an LLM-as-judge (any local model) that scores open-ended tasks 1–5 against a reference answer. It's a signal, not an oracle.
Benchmarking every model on the full suite takes a while on a laptop, so the defaults are tuned for a quick first look:
3 smallest models by default (smallest first, so results appear fast) — --all for everything, -m to choose.
A fast quality subset (~8 tasks across all categories) — --full for all 31.
Response caching : quality runs use temperature 0 + a fixed seed, so responses are deterministic and cached under ~/.local_bench . Re-running only regenerates new models/tasks (unchanged ones are re-graded from cache in milliseconds); --refresh-cache forces recompute, --no-cache disables it.
In practice this turns a first run from ~15–25 min (all models, full suite) into ~1–2 min, and a re-run into seconds. For a thorough pass (CI, final numbers) use local_bench --all --full .
Bring your own evals with a JSON or YAML pack — no Python required. --tasks replaces the built-in suite; --add-tasks appends to it. YAML needs the optional extra ( pip install ".[yaml]" ); JSON works out of the box.
# mypack.yaml — local_bench --tasks mypack.yaml
name : my-pack
tasks :
- id : capital_japan
category : factual
prompt : " What is the capital of Japan? Answer with just the city name. "
grader : {type: contains_any, values: ["Tokyo"]}
reference : Tokyo
- id : add
category : math
prompt : " What is 12 + 30? End with the answer on its own line. "
grader : {type: exact_number, value: 42}
- id : explain # no grader -> open-ended, scored only with --judge
category : open
prompt : " Explain photosynthesis in one sentence. "
reference : " Plants convert sunlight, water, and CO2 into glucose and oxygen. "
Grader type values: exact_number ( value , tol ), multiple_choice ( value ), contains_any ( values ), regex ( pattern , ignorecase ), valid_json ( keys ), valid_json_array ( length ). Omit grader for a judge-only task. Preview any pack with local_bench tasks --tasks mypack.yaml .
Two runnable packs ship in examples/ : sample_pack.yaml is a four-task look at the format, and extended_pack.yaml is a 21-task template covering every category and every grader type — copy it and swap in your own prompts.
Two rules keep a pack honest, both enforced by the tests for the shipped ones:
Every deterministic task's reference must pass its own grader. The reference is the worked example of a correct answer; if it can't clear the grader, nothing can, and the resulting 0 looks like a model failure rather than an authoring one.
max_tokens must have headroom. The budget covers the whole reply. With exact_number — which reads the last number — a model that reaches the right answer and then keeps talking gets cut off mid-sentence and scored on whatever number landed last. Ask for the answer last, and say to stop there.
Reviewing a pack before you run it
local_bench tasks --tasks mypack.yaml prints a pack to the terminal. For a bigger one there's also a small browser UI, tools/pack_viewer , which shows every task next to the exact grader spec that will judge it, charts prompt sizes against the context window the model will actually be loaded with, and applies both rules above to your pack — re-grading each reference with its own grader, and flagging prompts that will overflow the context:
pip install -r tools/pack_viewer/requirements.txt
streamlit run tools/pack_viewer/app.py -- mypack.yaml
A standalone utility rather than part of the package: nothing in src/ imports it, and Streamlit stays out of local_bench's own dependencies.
Judging with a different provider
By default the judge runs on the same provider as the models under test, so it can only be a model that provider serves. On one local runner that usually means the judge is also a candidate — and then it scores its own answers on the open-ended tasks.

[truncated]
