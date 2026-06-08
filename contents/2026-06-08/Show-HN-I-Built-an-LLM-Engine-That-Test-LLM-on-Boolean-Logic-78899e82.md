---
source: "https://github.com/Shrivastava-Aditya/boolean-algebra-engine"
hn_url: "https://news.ycombinator.com/item?id=48450471"
title: "Show HN: I Built an LLM Engine, That Test LLM on Boolean Logic"
article_title: "GitHub - Shrivastava-Aditya/boolean-algebra-engine: Deterministic boolean algebra engine — evaluates expressions, detects contradictions, audits logic rules. MCP server, NL layer, REST API, CLI, Streamlit UI. · GitHub"
author: "shrvx"
captured_at: "2026-06-08T21:39:10Z"
capture_tool: "hn-digest"
hn_id: 48450471
score: 1
comments: 0
posted_at: "2026-06-08T19:33:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I Built an LLM Engine, That Test LLM on Boolean Logic

- HN: [48450471](https://news.ycombinator.com/item?id=48450471)
- Source: [github.com](https://github.com/Shrivastava-Aditya/boolean-algebra-engine)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T19:33:48Z

## Translation

タイトル: Show HN: ブール論理で LLM をテストする LLM エンジンを構築しました
記事のタイトル: GitHub - Shrivastava-Aditya/boolean-algebra-engine: 決定論的ブール代数エンジン — 式を評価し、矛盾を検出し、ロジック ルールを監査します。 MCP サーバー、NL レイヤー、REST API、CLI、Streamlit UI。 · GitHub
説明: 決定論的ブール代数エンジン — 式を評価し、矛盾を検出し、論理ルールを監査します。 MCP サーバー、NL レイヤー、REST API、CLI、Streamlit UI。 - Shrivastava-Aditya/ブール代数エンジン
HN テキスト: 私は、Quine-McClusky 法を通じて数学的に解決できる基本的な論理的推論上の SOTA モデルの幻覚を追跡する LLM エンジンを構築しました。さまざまなユースケースを構築するためにそれを使用しています。何か試してほしいことがあれば言ってください。

記事本文:
GitHub - Shrivastava-Aditya/boolean-algebra-engine: 決定論的ブール代数エンジン — 式を評価し、矛盾を検出し、ロジック ルールを監査します。 MCP サーバー、NL レイヤー、REST API、CLI、Streamlit UI。 · GitHub
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
シュリヴァスタヴァ アディティヤ
/

ブール代数エンジン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
Shrivastava-Aditya/ブール代数エンジン
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
176 コミット 176 コミット .github/ workflows .github/ workflows archive/ memo archive/ memo boolean_algebra_engine boolean_algebra_engine core core docs docs image image mcp_server mcp_server テスト テスト ui ui .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md Dockerfile Dockerfile LICENSE LICENSE PUBLISHING.md PUBLISHING.md README.md README.md Analytics_notify.py Analytics_notify.py Analytics_watch.py Analytics_watch.py Benchmark.py Benchmark.py Curve.py Curve.py Curve_plot.py Curve_plot.pyダッシュボード.py ダッシュボード.py デモ.py デモ.py docker-compose.gpu.yml docker-compose.gpu.yml docker-compose.yml docker-compose.yml Index.html Index.html llms-full.txt llms-full.txt llms.txt llms.txt public.sh public.sh pyproject.toml pyproject.toml results_visualiser.ipynb results_visualiser.ipynb thanks.html thanks.html waitlist.html waitlist.html すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェント: 機械可読ドキュメントは、llms.txt (概要) および llms-full.txt (完全な API リファレンス) にあります。
AI エージェントの決定ルールは互いに矛盾する可能性があります。 LLM はそれを確実に捕捉できません。このエンジンは、おそらく 10 ミリ秒未満で実行できます。
意思決定ロジック (融資承認、コンプライアンスチェック、アクセス制御、ポリシー施行) を備えた AI エージェントは、人間が作成したブール ルールに基づいて実行されます。エージェントがルールに基づいて動作する前に、それらのルールが矛盾していないかどうかを検証する人は誰もいません。このエンジンはそうなります。
ベンチマークは、なぜそれが必要なのかを示しています。70B モデルでさえ、ブール論理の質問の最大 20% が間違っています。 LLM に次のことを尋ねることはできません

あなたのルールは矛盾しており、その答えを信頼しています。それを計算する決定論的な層が必要です。
90 件のテストに合格 · 10ms 未満の評価 · 依存関係なし · サンプリングではなく徹底的な列挙
pip インストール ブール代数エンジン
boolean_algebra_engine からインポート評価、合成
＃矛盾は存在しますか？
table , _ = 評価 ( "A.!A" )
print ( table .satisfiable ) # False — 常に矛盾
# 2 つのルールが同時に真になることはありますか?
table , _ = 評価 ( "(A.B).(!A)" )
print ( table .satisfiable ) # False — A と !A を両方とも保持することはできません
# 完全な真理値表
table , _ = 評価 ( "A.(B+C)" )
print ( table . variables ) # ['A', 'B', 'C']
print ( table . minterms ) # [5, 6, 7]
print ( table .satisfiable ) # True
# 最小限の形式に簡略化する
最小、_ = 合成 (テーブル)
印刷 (最小限) # A.C+A.B
ワンライナーチェック:
python -c " from boolean_algebra_engine import Evaluate; t,_ = Evaluate('A.!A'); print(t.satisfiable) "
# 偽
オプションの追加物
pip install " boolean-algebra-engine[cli] " # CLI / REPL
pip install " boolean-algebra-engine[mcp] " # MCP サーバー (クロード デスクトップ)
pip install " boolean-algebra-engine[api] " # REST API
pip install " boolean-algebra-engine[nl-anthropic] " # Anthropic 経由の NL 層
pip install " boolean-algebra-engine[nl-openai] " # OpenAI 経由の NL 層
問題
6つのルール。 3 つの変数。 4人が半年かけて執筆。
フィンテック AI エージェントは、これらのルールに基づいてローン申請を自動承認または拒否します。誰もそれらを一緒に検証したことはありません。エンジンは、すべてのルールのすべての入力の組み合わせをすべての組み合わせでチェックします。
# pip インストール ブール代数エンジン[mcp]
boolean_algebra_engine から。 mcp 。サーバーインポート check_prompt_logic
結果 = check_prompt_logic ([
"A.B" 、 # 承認: 良好な信用と収入が確認されました
"!A" 、 # 拒否: 悪い

クレジット
"C" , # 承認: 担保が存在します
"!C" , # 拒否: 担保なし
])
print (結果 [ "概要" ])
# {'合計': 4, '矛盾': 0, 'トートロジー': 0,
# 'equivalent_pairs': 0, 'conflicting_pairs': 2}
print ([( p [ "rule1" ], p [ "rule2" ]) for p in result [ "pairwise" ] if p [ "always_conflict" ]])
# [('A.B', '!A'), ('C', '!C')]
見つかったこと:
A.B と !A の競合 - A=1 の場合、良好な信用承認と不良信用拒否が同時に発生します。エージェントが任意に勝者を選択します。
C と !C の競合 - 担保の承認と担保なしの拒否は、定義上、相互に排他的です。両方のルールを同時に適用することはできません。
ルールを読んでこれらを発見した人は誰もいませんでした。エンジンはあらゆる組み合わせをチェックしてそれらを捕らえました。
テストされたすべてのモデルはブール論理に基づいて幻覚を起こしますが、その方法はサイズとアーキテクチャに応じて異なります。
小規模モデルはまったく論理的ではなく、デフォルトを選択し、それに固執しました。大型のモデルは論理的に考えていますが、依然として幻覚を持っています。 llama-3.3-70b のスコアは 20% ですが、発生するエラーは 1 種類のみです。ルールが競合していない場合でも競合するとみなします (競合の見逃し 0%、互換性の見逃し 40%)。
変動曲線 — 複雑さは状況を悪化させるのでしょうか?
qwen3-32b は、3 ～ 10 (真理値表の行数 8 ～ 1,024) の変数カウントで、それぞれ 100 件実行されました。幻覚率は全体を通して 16 ～ 19% で横ばいでした。複雑さによって性能が低下することはありません。エラーは一貫したベースラインであり、より複雑なロジックによって引き起こされたものではありません。
方法論: 正しい答え (満足できるかどうか) が正確にわかっているブール式のペアを生成します。 LLMに聞いてください。比較する。曖昧さ、人間によるラベル付け、解釈はありません。エンジンはオラクルです。グランド トゥルースは推測ではなく、徹底的な列挙によって計算されます。 LLM の不一致はすべて、証明可能な幻覚です。
python3ベンチマーク.py

--プロバイダー ollama --モデル tinyllama --ケース 20
python3 benchmark.py --provider ollam --model llama3.2:3b --cases 20
python3 benchmark.py --provider ollama --model gemma3:4b --cases 20
タイニーラマ — 1.1B
正解：10/20・幻覚率：50.0%
競合の見逃し: 10/10 (100%) · 互換性の見逃し: 0/10 (0%)
パターン: 常に「はい」を出力します - ケースバイケースの推論はありません
ラマ3.2:3b — 3B
正解：10/20・幻覚率：50.0%
競合の見逃し: 0/10 (0%) · 互換性の見逃し: 10/10 (100%)
パターン: 常に "no" を出力します。デフォルトとは反対で、同じ推論が失敗します。
ジェマ3:4b — 4B
正解：13/20・幻覚率：35.0%
競合の見逃し: 4/10 (40%) · 互換性の見逃し: 3/10 (30%)
パターン: 各ケースを個別に扱うが、3 分の 1 が間違っている
qwen3-32b — 32B
正解: ~83/100 · 幻覚率: ~17%
3 ～ 10 個の変数にわたってフラット — 誤差は一貫したベースラインであり、複雑さによって引き起こされるものではありません
ラマ-3.3-70b — 70B
正解: ~80/100 · 幻覚率: ~20%
競合の見逃し: 0% · 互換性の見逃し: ~40%
パターン: 慎重すぎる — 存在しない競合にはフラグを立てませんが、互換性のある 5 つのペアのうち 2 つを見逃します
エンジンコラムは本物です。 llm との不一致はすべて証明可能な幻覚であり、意見ではありません。
boolean_algebra_engine からインポート評価、合成
# 前方: 式 → 真理値表
table , _ = 評価 ( "A.(B+C)" )
print ( table . variables ) # ['A', 'B', 'C']
print ( table . minterms ) # [5, 6, 7]
print ( table .satisfiable ) # True
# 逆: 真理値表 → 最小式
最小、_ = 合成 (テーブル)
印刷 (最小限) # A.C+A.B
# 等価性と充足性 (MCP サーバー関数経由 - HTTP なし、直接呼び出し)
# pip インストール ブール代数エンジン[mcp]
boolean_algebra_engine から。 mcp 。サーバーインポート同等、満足

有能な
print ( 等価 ( "A.(B+C)" , "A.B+A.C" )[ "等価" ]) # True — 分配法則
print (satisfiable ( "A.!A" )[ "satisfiable" ]) # False — 矛盾
コア モジュールには外部依存性がありません。それを任意の Python プロジェクトにインポートします。
エンジンを Claude Desktop に接続すると、Claude はブール ロジックの予測を停止します。それを計算してくれるのです。
{
"mcpサーバー": {
"ブール代数エンジン" : {
"コマンド" : " Python " ,
"args" : [ " -m " 、 " boolean_algebra_engine.mcp.server " ]
}
}
}
クロードが会話中に呼び出せる 5 つのツール:
評価 — 式 → 真理値表
単純化 — 式 → 最小形式
同等 — 2 つの式は同一ですか?
満足可能 — これを真にする入力はありますか?
check_prompt_logic — ルールセット全体の矛盾、トートロジー、衝突、重複を監査する
LLM が翻訳します。エンジンが決める。ロジックは常に正確です。
単純な英語のルールを受け入れるほとんどのツールは、ルールを LLM に渡し、返されたものをすべて信頼します。それは要約に役立ちます。これはロジックでは機能しません。このリポジトリのベンチマークは、テストしたすべてのモデルで幻覚率が 17% から 50% であることを証明しています。
NL 層は異なるアプローチを採用します。 LLM は、実際に得意とすること、つまり人間の言語をシンボルにマッピングするために LLM を使用します。文が式になると、LLM は完了します。エンジンが引き継ぎ、考えられるすべての入力の組み合わせを評価し、正しいと証明された結果を返します。
「ユーザーが管理者であるか、認証されていて一時停止されていない場合にアクセスが許可されます」
「管理者である場合、または検証済みで一時停止されていない場合にアクセスが許可されます」
│
▼ LLM — 単語を変数にマッピングし、式を返します
A+(V.!S)
│
▼ コアエンジン — すべての入力の組み合わせを評価します
真理値表: 8 行 · 6 行 (出力 = 1)
最小形式: A+V.!S
満足できる: はい · ピンと張った

論理: いいえ · 矛盾: いいえ
│
▼ LLM — 結果を平易な英語に戻します
「8 件中 6 件でアクセスが許可されました。一時停止が確認されました」
ユーザーは、管理者のステータスに関係なく、常に拒否されます。」
LLM は 2 回関与します。2 回ともあいまいなこと (言語) を実行し、正確なこと (ロジック) を実行することはありません。
最速のパスは Ollama です。ローカルで実行され、API キーは不要で、無料です。
#1. Ollama をインストールする
カール -fsSL https://ollama.com/install.sh |しー
#2. モデルをプルする
ollama pull deepseek-r1:最新 # 5.2 GB、最高品質
# ollama pull deepseek-r1:1.5b # 1.1 GB、低メモリマシン用
# 3. パッケージをインストールする
pip インストール ブール代数エンジン
boolean_algebra_engine から。 nl 。 nl インポート依頼
result = ask ( "ドアまたは窓が開いている場合はアラームが鳴りますが、システムが無効になっている場合は鳴らされません" )
print (結果.式) # (D+W).!S
print (結果の最小値) # D.!S+W.!S
print ( result .satisfiable ) # True
print ( result . variables ) # {'D': 'ドアが開いています', 'W': '窓が開いています', 'S': 'システムが無効です'}
print（結果と説明） #わかりやすい英語まとめ
プロバイダーは自動検出されます。Ollama が実行されている場合は、自動的に使用されます。設定は必要ありません。
boolcalc ask 「ドアまたは窓が開いている場合はアラームが作動しますが、システムが無効になっている場合は作動しません」
boolcalc check-rules " 管理者の場合はアクセス " " 検証された場合はアクセス " " 一時停止の場合はアクセスなし "
提供の選択

[切り捨てられた]

## Original Extract

Deterministic boolean algebra engine — evaluates expressions, detects contradictions, audits logic rules. MCP server, NL layer, REST API, CLI, Streamlit UI. - Shrivastava-Aditya/boolean-algebra-engine

I have built an LLM Engine, That tracks SOTA Models hallucinating over basic Logical Reasoning that can be solved through Quine-McClusky method mathematically. I am using it to build diverse use cases. Let me know if you want me to try it out on something.

GitHub - Shrivastava-Aditya/boolean-algebra-engine: Deterministic boolean algebra engine — evaluates expressions, detects contradictions, audits logic rules. MCP server, NL layer, REST API, CLI, Streamlit UI. · GitHub
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
Shrivastava-Aditya
/
boolean-algebra-engine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Shrivastava-Aditya/boolean-algebra-engine
master Branches Tags Go to file Code Open more actions menu Folders and files
176 Commits 176 Commits .github/ workflows .github/ workflows archive/ memo archive/ memo boolean_algebra_engine boolean_algebra_engine core core docs docs images images mcp_server mcp_server tests tests ui ui .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md Dockerfile Dockerfile LICENSE LICENSE PUBLISHING.md PUBLISHING.md README.md README.md analytics_notify.py analytics_notify.py analytics_watch.py analytics_watch.py benchmark.py benchmark.py curve.py curve.py curve_plot.py curve_plot.py dashboard.py dashboard.py demo.py demo.py docker-compose.gpu.yml docker-compose.gpu.yml docker-compose.yml docker-compose.yml index.html index.html llms-full.txt llms-full.txt llms.txt llms.txt publish.sh publish.sh pyproject.toml pyproject.toml results_visualiser.ipynb results_visualiser.ipynb thanks.html thanks.html waitlist.html waitlist.html View all files Repository files navigation
Coding agents: machine-readable docs are in llms.txt (summary) and llms-full.txt (full API reference).
Your AI agent's decision rules might contradict each other. LLMs can't reliably catch that. This engine can — provably, in under 10ms.
AI agents with decision logic — loan approval, compliance checks, access control, policy enforcement — run on boolean rules written by humans. Nobody verifies those rules don't conflict before the agent acts on them. This engine does.
The benchmark shows why you need it: even a 70B model gets ~20% of boolean logic questions wrong. You can't ask an LLM if your rules conflict and trust the answer. You need a deterministic layer that computes it.
90 tests passing · <10ms evaluation · zero dependencies · exhaustive enumeration, not sampling
pip install boolean-algebra-engine
from boolean_algebra_engine import evaluate , synthesize
# Does a contradiction exist?
table , _ = evaluate ( "A.!A" )
print ( table . satisfiable ) # False — always a contradiction
# Can two rules both be true simultaneously?
table , _ = evaluate ( "(A.B).(!A)" )
print ( table . satisfiable ) # False — A and !A can't both hold
# Full truth table
table , _ = evaluate ( "A.(B+C)" )
print ( table . variables ) # ['A', 'B', 'C']
print ( table . minterms ) # [5, 6, 7]
print ( table . satisfiable ) # True
# Simplify to minimal form
minimal , _ = synthesize ( table )
print ( minimal ) # A.C+A.B
One-liner check:
python -c " from boolean_algebra_engine import evaluate; t,_ = evaluate('A.!A'); print(t.satisfiable) "
# False
Optional extras
pip install " boolean-algebra-engine[cli] " # CLI / REPL
pip install " boolean-algebra-engine[mcp] " # MCP server (Claude Desktop)
pip install " boolean-algebra-engine[api] " # REST API
pip install " boolean-algebra-engine[nl-anthropic] " # NL layer via Anthropic
pip install " boolean-algebra-engine[nl-openai] " # NL layer via OpenAI
The problem
Six rules. Three variables. Written by four people over six months.
A fintech AI agent auto-approves or rejects loan applications based on these rules — nobody ever verified them together. The engine checks all input combinations for every rule, in every combination:
# pip install boolean-algebra-engine[mcp]
from boolean_algebra_engine . mcp . server import check_prompt_logic
result = check_prompt_logic ([
"A.B" , # approve: good credit AND income verified
"!A" , # reject: bad credit
"C" , # approve: collateral exists
"!C" , # reject: no collateral
])
print ( result [ "summary" ])
# {'total': 4, 'contradictions': 0, 'tautologies': 0,
# 'equivalent_pairs': 0, 'conflicting_pairs': 2}
print ([( p [ "rule1" ], p [ "rule2" ]) for p in result [ "pairwise" ] if p [ "always_conflict" ]])
# [('A.B', '!A'), ('C', '!C')]
What it found:
A.B and !A conflict — good credit approval and bad credit rejection fire simultaneously when A=1 . The agent picks a winner arbitrarily.
C and !C conflict — collateral approval and no-collateral rejection are mutually exclusive by definition. Both rules can never apply at the same time.
Nobody caught these by reading the rules. The engine caught them by checking every combination.
Every model tested hallucinates on boolean logic — but in different ways depending on size and architecture.
The small models aren't reasoning at all — they picked a default and stuck to it. The larger models reason but still hallucinate. llama-3.3-70b scores 20% but makes only one type of error: it assumes rules conflict when they don't (0% missed conflicts, 40% missed compatibles).
Variable curve — does complexity make it worse?
qwen3-32b was run across variable counts from 3 to 10 (8 to 1,024 truth table rows), 100 cases each. The hallucination rate stayed flat at 16–19% throughout. Complexity doesn't degrade it — the errors are a consistent baseline, not caused by harder logic.
Methodology: generate pairs of boolean expressions where the correct answer (satisfiable or not) is known exactly. Ask the LLM. Compare. No ambiguity, no human labeling, no interpretation. The engine is the oracle — ground truth is computed by exhaustive enumeration, not guessed. Every LLM disagreement is a provable hallucination.
python3 benchmark.py --provider ollama --model tinyllama --cases 20
python3 benchmark.py --provider ollama --model llama3.2:3b --cases 20
python3 benchmark.py --provider ollama --model gemma3:4b --cases 20
tinyllama — 1.1B
correct: 10/20 · hallucination rate: 50.0%
missed conflicts: 10/10 (100%) · missed compatibles: 0/10 (0%)
pattern: always outputs "yes" — no case-by-case reasoning
llama3.2:3b — 3B
correct: 10/20 · hallucination rate: 50.0%
missed conflicts: 0/10 (0%) · missed compatibles: 10/10 (100%)
pattern: always outputs "no" — opposite default, same reasoning failure
gemma3:4b — 4B
correct: 13/20 · hallucination rate: 35.0%
missed conflicts: 4/10 (40%) · missed compatibles: 3/10 (30%)
pattern: engages with each case individually — but wrong 1 in 3
qwen3-32b — 32B
correct: ~83/100 · hallucination rate: ~17%
flat across 3–10 variables — errors are a consistent baseline, not complexity-driven
llama-3.3-70b — 70B
correct: ~80/100 · hallucination rate: ~20%
missed conflicts: 0% · missed compatibles: ~40%
pattern: over-cautious — never flags a conflict that isn't there, but misses 2 in 5 compatible pairs
The engine column is ground truth. Every mismatch with llm is a provable hallucination — not an opinion.
from boolean_algebra_engine import evaluate , synthesize
# Forward: expression → truth table
table , _ = evaluate ( "A.(B+C)" )
print ( table . variables ) # ['A', 'B', 'C']
print ( table . minterms ) # [5, 6, 7]
print ( table . satisfiable ) # True
# Inverse: truth table → minimal expression
minimal , _ = synthesize ( table )
print ( minimal ) # A.C+A.B
# Equivalence and satisfiability (via MCP server functions — no HTTP, direct call)
# pip install boolean-algebra-engine[mcp]
from boolean_algebra_engine . mcp . server import equivalent , satisfiable
print ( equivalent ( "A.(B+C)" , "A.B+A.C" )[ "equivalent" ]) # True — distributive law
print ( satisfiable ( "A.!A" )[ "satisfiable" ]) # False — contradiction
The core module has zero external dependencies. Import it into any Python project.
Wire the engine into Claude Desktop and Claude stops predicting boolean logic. It computes it.
{
"mcpServers" : {
"boolean-algebra-engine" : {
"command" : " python " ,
"args" : [ " -m " , " boolean_algebra_engine.mcp.server " ]
}
}
}
Five tools Claude can call mid-conversation:
evaluate — expression → truth table
simplify — expression → minimal form
equivalent — are two expressions identical?
satisfiable — does any input make this true?
check_prompt_logic — audit a full rule set for contradictions, tautologies, conflicts, duplicates
The LLM translates. The engine decides. The logic is always exact.
Most tools that accept plain English rules pass them to an LLM and trust whatever comes back. That works for summarisation. It doesn't work for logic — the benchmark in this repo proves it, with hallucination rates from 17% to 50% across every model tested.
The NL layer takes a different approach. It uses the LLM for the one thing it's actually good at: mapping human language to symbols. Once the sentence becomes an expression, the LLM is done. The engine takes over, evaluates every possible input combination, and returns a result that is provably correct.
"access granted if the user is an admin, or if they're verified and not suspended"
"access granted if admin, or verified and not suspended"
│
▼ LLM — maps words to variables, returns an expression
A+(V.!S)
│
▼ core engine — evaluates every input combination
truth table: 8 rows · 6 rows where output = 1
minimal form: A+V.!S
satisfiable: yes · tautology: no · contradiction: no
│
▼ LLM — turns the result back into plain English
"Access is granted in 6 of 8 cases. A suspended verified
user is always denied, regardless of admin status."
The LLM is involved twice — both times doing something fuzzy (language), never something exact (logic).
The fastest path is Ollama — runs locally, no API key, free.
# 1. Install Ollama
curl -fsSL https://ollama.com/install.sh | sh
# 2. Pull a model
ollama pull deepseek-r1:latest # 5.2 GB, best quality
# ollama pull deepseek-r1:1.5b # 1.1 GB, for low-memory machines
# 3. Install the package
pip install boolean-algebra-engine
from boolean_algebra_engine . nl . nl import ask
result = ask ( "alarm on if door open or window open, but not if system disabled" )
print ( result . expression ) # (D+W).!S
print ( result . minimal ) # D.!S+W.!S
print ( result . satisfiable ) # True
print ( result . variables ) # {'D': 'door is open', 'W': 'window is open', 'S': 'system is disabled'}
print ( result . explanation ) # plain English summary
The provider is auto-detected: if Ollama is running, it is used automatically. No configuration needed.
boolcalc ask " alarm on if door open or window open, but not if system disabled "
boolcalc check-rules " access if admin " " access if verified " " no access if suspended "
Choosing a provid

[truncated]
