---
source: "https://github.com/yashmahajan10/llm-differential-privacy-gateway"
hn_url: "https://news.ycombinator.com/item?id=49119049"
title: "Noisegate: A differential-privacy gateway for untrusted AI agents"
article_title: "GitHub - yashmahajan10/llm-differential-privacy-gateway: Noisegate: a differential privacy gateway that lets an untrusted LLM agent query sensitive data over MCP (Model Context Protocol), with a formal guarantee no individual's record can leak even if the agent is adversarial - enforcement lives in\n[truncated]"
author: "handfuloflight"
captured_at: "2026-07-31T05:26:36Z"
capture_tool: "hn-digest"
hn_id: 49119049
score: 1
comments: 0
posted_at: "2026-07-31T04:36:02Z"
tags:
  - hacker-news
  - translated
---

# Noisegate: A differential-privacy gateway for untrusted AI agents

- HN: [49119049](https://news.ycombinator.com/item?id=49119049)
- Source: [github.com](https://github.com/yashmahajan10/llm-differential-privacy-gateway)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T04:36:02Z

## Translation

タイトル: Noisegate: 信頼できない AI エージェント用の差分プライバシー ゲートウェイ
記事のタイトル: GitHub - yashmahajan10/llm-Difference-privacy-gateway: Noisegate: 信頼できない LLM エージェントが MCP (モデル コンテキスト プロトコル) 経由で機密データをクエリできるようにする差分プライバシー ゲートウェイ。エージェントが敵対的であっても個人の記録が漏洩しないという正式な保証があり、強制力が生きています。
[切り捨てられた]
説明: ノイズゲート: 信頼できない LLM エージェントが MCP (モデル コンテキスト プロトコル) 経由で機密データをクエリできるようにする差分プライバシー ゲートウェイで、エージェントが敵対的であっても個人の記録が漏洩しないことを正式に保証します。強制はモデルの下の信頼できるコード内に存在し、実行可能な att によって検証されます。
[切り捨てられた]

記事本文:
GitHub - yashmahajan10/llm-Difference-privacy-gateway: Noisegate: 信頼できない LLM エージェントが MCP (モデル コンテキスト プロトコル) 経由で機密データをクエリできるようにする差分プライバシー ゲートウェイ。エージェントが敵対的であっても個人の記録が漏洩しないことが正式に保証されています。適用はモデルの下の信頼できるコード内に存在し、実行可能な攻撃ギャラリーによって検証されています。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
あなたは署名します

別のタブまたはウィンドウに表示されます。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ヤシュマハジャン10
/
llm-差分-プライバシー-ゲートウェイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
134 コミット 134 コミット .github/ workflows .github/ workflows api api 攻撃 攻撃 ベンチマーク ベンチマーク コンパイラ コンパイラ データセット データセット ドキュメント ドキュメント エンジン エンジン ID アイデンティティ mcp_server mcp_server スクリプト スクリプト サービス サービス テスト テスト UI UI 検証 検証 .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff DESIGN.md DESIGN.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md SETUP.md SETUP.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml pytest.ini pytest.ini 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Noisegate: 信頼できない AI エージェント用の差分プライバシー ゲートウェイ
AI エージェントのクエリに機密データへのアクセスを許可します。その際、エージェントが間違っていたり、操作されたり、敵対的であったりしても、個人の記録が漏洩しないという数学的保証が与えられます。
記録されたクロード デスクトップ セッション (返信はトリミングされています。グラフ カードはセッション独自のものです)。 AI エージェントが 20 人の患者を診断によって分類すると、±12 のノイズがすべてのビンに発生します。ノイズを消すことはできないことを思い出させてください。ゲートが静かな答えの代わりに拒否を返すまで、3 つの答えの予算を使い果たします。 32,561 行の国勢調査では、狭すぎるスライスは信頼境界で拒否されますが、教育全体の内訳は大規模にクリーンに返されます。拒否と再

ジェクションは、ライブゲートウェイの実際の強制であり、 python scripts/render_demo_gif.py によって再現されます。
主張ではなく、有効な攻撃です。 3 つの古典的なプライバシー攻撃 (差分、メンバーシップ推論、再識別による選別) は、システム独自のエンジンに対して実行されます。それぞれがプライバシーをオフにして成功し、プライバシーをオンにして敗北し、防御が静かに腐らないようにCIに固定されていることが示されています。
独立して検証された数学。ノイズ メカニズムはゼロから構築されており、35 のノイズ スケール チェックすべてにおいて 1e-9 以内で業界リファレンス実装である OpenDP と一致します。
会計の厳格化からさらに多くの質問が。デプロイメントのクエリーごとの ε では、ハイブリッド zCDP 構成では、同じバジェットに対して 308 個のクエリーが許可されます。これに対して、高度な構成では 268 個、単純合計ε では 100 個が許可されます。保証されるのは、純粋な ε ではなく (ε, δ)-DP です。
AI エージェント向けに構築されています。 Claude Desktop の MCP サーバーとして実行します。接続エージェントは設計上信頼されておらず、すべてのプライバシー プロパティがその下で強制されます。
スタック: Python、DuckDB、FastAPI、Streamlit、MCP SDK、Docker、GitHub Actions、CI の 250 以上のテスト スイート。
機密データセットについて平易な英語で質問します。 LLM は、各質問を小さな制約されたクエリにコンパイルします。差分プライバシー エンジンは、追跡されたプライバシー バジェットに基づいてそれを実行し、指定された信頼区間で意図的にノイズの多い回答を返します。名前がぴったりです。オーディオの同名のように、ゲートウェイはすべての信号をノイズ フロアの下で設定されたしきい値以下に保ちます。個人の寄与はかき消され、人口規模の信号はほとんど影響を受けずに通過します。何を質問されたとしても、またそれがどれほど巧みに表現されたとしても、回答から個々の記録を再構成することはできません。
興味深いのは、LLM がクエリを作成できるということではありません。それは

プライバシーの保証は、LLM が信頼できるかどうかに依存しません。モデルはクエリを提案する便宜的なものです。何も強制しません。すべてのプライバシー プロパティは、人間が手動でクエリを入力した場合と同じように動作するコンポーネントによって下流で強制されます。これは、本番システムの信頼できない入力に適用される信頼境界の規律であり、ここでは AI エージェントに適用されます。
1. 攻撃を実行します。API キー、データ取得、サーバーは使用しません。
攻撃ギャラリーは、実際の DP エンジンに対してインプロセスで実行されます。
pip install -e 。
python -m Attacks.patients_alice # プライバシーをオフにしてアリスを再識別してから監視します
# 警備員と騒音と予算がそれを倒す
2. AIエージェントを接続する
ゲートウェイは、Claude Desktop の MCP stdio サーバーとして実行されます。エージェントは信頼できないクエリ作成者となり、引数スキーマがデータセット ポリシーから生成される構造化ツール ( count 、 sum 、 Average 、 histogram 、 get_budget ) のみを取得します。接続エージェントはインテリジェンスであるため、API キーはどこにも必要ありません。
3. 完全な自然言語 UI
ローカルのシングルテナントのデモ。 API キーは、信頼できない NL→クエリ コンパイラにのみ必要です。
import ANTHROPIC_API_KEY=... # 信頼できない NL→クエリ コンパイラによってのみ使用されます
docker compose up # エンジン、API、UI を起動します
# http://localhost:8501 を開きます
その HTTP + Streamlit サーフェスは、ローカルのシングルテナントのデモです。 ID はスプーフィング可能な X-Identity ヘッダーから取得されるため、パブリック展開ではなく、自分のマシン上の 1 人の信頼できるオペレーターを対象としています (これらのサーフェスがどのようなもので、そうでないのかを参照してください)。ローカルの非 Docker セットアップ、テストの実行、および構成ノブについては、 SETUP.md を参照してください。
見出し: 攻撃ギャラリー
誰でもプライバシーを主張できます。このリポジトリは、主張を破るエクスプロイトを同梱し、それを独自のエンジンに対して実行します。

d 結果を CI に固定します。ゲートウェイが何を保証しているかを理解する最も早い方法は、単純な「データベースのクエリ」システムを破壊する 3 つの古典的な攻撃をゲートウェイが打ち破る様子を観察することです。
攻撃 1 — 差分攻撃
差分攻撃では、その人ごとに異なる 2 つの質問をまとめて質問することで、1 人の人を隔離します。
クエリ A: 「X 部門の 100 人全員の合計収入」 → 7,240,000ドル
クエリ B: 「部門 X のアリスを除くすべての人の収入の合計」 → 7,135,000ドル
攻撃者は次のように計算します: A − B = 105,000 ドル ← アリスの正確な給与、漏洩。
どちらのクエリも「単なる集計」です。どちらも単一の行に名前を付けません。しかし、それらは一緒になって個人を暴露します。ギャラリー ( Attacks/difference.py ) は、プライバシーを無効にしてこの攻撃が成功したことを示しています。ターゲットのプライベート値は正確に復元されています。 (上記の給与スケッチは説明用です。実際の UCI アダルト データでは、「アリス」はグループの最大キャピタル ゲインの唯一の所有者です。) 次に、DP がオンになった後に同じ攻撃が敗北したことを示しています。各回答の調整されたノイズにより減算が無意味になり、予算会計担当者は両方のクエリを独立したものとして扱うのではなく、両方のクエリにわたって公開された情報に対して請求します。
攻撃 2 — メンバーシップの推論
メンバーシップ推論攻撃は、特定の個人がデータセット内に存在するかどうかを判断します。多くのデータセット (医学研究、債務不履行者のリスト) では、その事実自体が重要です。クエリ アクセスのみを持つ攻撃者は、「この人物がデータ内に存在するのか?」と判断しようとします。
ギャラリーは、この攻撃を広範囲のプライバシー バジェット (ε - 回答の精度をプライバシーと引き換えにするダイヤル) に対して実行し、結果をプロットします。
単一のノイズの多い COUNT から 1 人のメンバーシップを決定する最適な (Neyman-Pearson) 攻撃者が、ε にわたって実際のエンジンに対して実行されます。経験的な成功 (青、95)

% ウィルソン間隔) は解析的なラプラス曲線に沿っており、最悪の場合の DP 上限 (破線) を下回っています。 ε が縮小するにつれて、確実性 (DP オフ) から 0.5 コイン投げに向かって崩壊します。緑色のユーティリティ曲線 (右軸) は、同じスイープにおける集計クエリの相対誤差を示しており、攻撃が無効になったところではほとんど凹みません。 Python -m Attacks.membership によって生成されます (ε あたり 10,000 回の試行)。
ε が縮小する (プライバシーが強化される) と、攻撃者の成功はコイントスに向けて崩壊します。ユーティリティのオーバーレイは、支払った代償を示しています。1 人の攻撃を無効にする同じノイズが、人口規模の集合体をほとんど動かさないのです。プライバシーは無料ではありません。グラフは、プライバシーと引き換えに何を犠牲にするかを正確に示しています。スイープ範囲は ε = 8 から 0.5 までで、全範囲をカバーします。バンドルされたデプロイメントの料金は、このグラフの左端にあるクエリごとに ε = 0.05 であり、単一クエリ攻撃はすでに偶然と区別できません。
攻撃 3 — 再識別による特定
ラターニャ・スウィーニーは 2002 年に、郵便番号 + 生年月日 + 性別でアメリカ人の約 87% を一意に識別できることを示しました。データセット内で誰かを見つけるのに名前は必要ありません。いくつかの無害な属性で十分です。ギャラリーの 3 回目の攻撃 ( Attacks/patients_alice.py ) は、20 人の合成患者でその構造を再現しています。性別と年齢だけを考慮すると、コホートはちょうど 1 人、64 歳以上の唯一の女性であるアリスに絞り込まれます。
プライバシーをオフにして、2 つのまったく普通の人口統計ヒストグラムを差し引くと、アリスの診断が正確に復元されます。外れ値は必要ありません。再識別可能であれば十分なので、これがギャラリー内で最も強力な攻撃になります。プライバシーがオンの場合、3 つの独立した防御によってプライバシーが終了します。フィルター ガードは明らかな狭いクエリを完全に拒否します (何も費やしません)。調整されたノイズは 2 つのクエリの減算をかき消します (信号対ノイズ ≈ 0.13 なので、reco

確かな「診断」は基本的にランダムな抽選です)。そして予算は、平均エスカレーションが機能するずっと前に平均エスカレーションを拒否します (最大 256 回の繰り返しが必要で、10 回は手頃な価格)。これは、このページの上部にあるデモ GIF と、Claude Desktop のウォークスルー、ライブ再生の攻撃です。結果は、 testing/test_attach_patients.py で回帰テストされます。
システム自体のエンジンに対して実行されるこれらの実験は、保証が本物で理解されており、ライブラリから輸入されたものや信仰に基づいたものではないことを示す中心的な証拠です。
この攻撃は、私たち自身の計算に照らして保証を検証します。自己一貫性はあるものの間違った実装を防ぐために、このメカニズムは独立した参照として OpenDP に対してクロスチェックも行われます (攻撃/crosscheck_opendp.py、opendp 0.15.1)。
つまり、ゼロから構築されたメカニズムは小数点以下 9 桁までの業界基準と一致しており、意図的に誤って調整されたバージョンは同じテストに合格しないため、この一致には意味があるのです。一致は 2 つの方法でチェックされます。OpenDP の認定限界に対する感度とノイズ スケール、および OpenDP に対するサンプラーの実際の分布と、テストが失敗する可能性があることを証明するポジティブ コントロールです。測定値は DESIGN.md セクション 6.3 にあります。
OpenDP はリファレンスであり、決して実装ではありません。

[切り捨てられた]

## Original Extract

Noisegate: a differential privacy gateway that lets an untrusted LLM agent query sensitive data over MCP (Model Context Protocol), with a formal guarantee no individual's record can leak even if the agent is adversarial - enforcement lives in trusted code below the model, validated by a runnable att
[truncated]

GitHub - yashmahajan10/llm-differential-privacy-gateway: Noisegate: a differential privacy gateway that lets an untrusted LLM agent query sensitive data over MCP (Model Context Protocol), with a formal guarantee no individual's record can leak even if the agent is adversarial - enforcement lives in trusted code below the model, validated by a runnable attack gallery. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
yashmahajan10
/
llm-differential-privacy-gateway
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
134 Commits 134 Commits .github/ workflows .github/ workflows api api attacks attacks benchmarks benchmarks compiler compiler datasets datasets docs docs engine engine identity identity mcp_server mcp_server scripts scripts service service tests tests ui ui validation validation .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff DESIGN.md DESIGN.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SETUP.md SETUP.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml pytest.ini pytest.ini requirements.txt requirements.txt View all files Repository files navigation
Noisegate: a differential-privacy gateway for untrusted AI agents
Give an AI agent query access to sensitive data, with a mathematical guarantee that no individual's record can leak — even if the agent is wrong, manipulated, or adversarial.
A recorded Claude Desktop session (replies trimmed; the chart cards are the session's own). An AI agent breaks 20 patients down by diagnosis, and the ±12 noise swamps every bin. Reminded that it cannot turn the noise off, it drains a three-answer budget until the gate returns a refusal instead of a quieter answer. On the 32,561-row census, a too-narrow slice is rejected at the trust boundary , while a full education breakdown comes back clean at scale. The refusal and the rejection are the live gateway's real enforcement, reproduced by python scripts/render_demo_gif.py .
Working attacks, not claims. Three classic privacy attacks — differencing, membership inference, and singling out by re-identification — run against the system's own engine. Each one is shown succeeding with privacy off, defeated with privacy on, and pinned in CI so the defense cannot quietly rot.
Independently verified math. The noise mechanism is built from scratch, and it matches OpenDP , the industry reference implementation, in all 35 noise-scale checks to within 1e-9.
More questions from tighter accounting. At the deployment's per-query ε, hybrid zCDP composition admits 308 queries against the same budget, against 268 under advanced composition and 100 under naive sum-of-ε. The guarantee is (ε, δ)-DP rather than pure ε.
Built for AI agents. Runs as an MCP server for Claude Desktop. The connecting agent is untrusted by design, and every privacy property is enforced below it.
Stack: Python · DuckDB · FastAPI · Streamlit · MCP SDK · Docker · GitHub Actions, with a 250+ test suite in CI.
You ask questions about a sensitive dataset in plain English. An LLM compiles each question into a small, constrained query. A differential-privacy engine executes it under a tracked privacy budget and returns a deliberately noisy answer with a stated confidence interval. The name fits. Like its audio namesake , the gateway keeps every signal below a set threshold under the noise floor: any one individual's contribution is drowned out, while population-scale signals pass through nearly untouched. No individual record can be reconstructed from the answers, no matter what is asked or how cleverly it is phrased.
The interesting part isn't that an LLM can write queries. It's that the privacy guarantee does not depend on the LLM being trustworthy. The model is a convenience that proposes a query. It enforces nothing. Every privacy property is enforced downstream, by components that would behave the same way if a human typed the query by hand. This is the trust-boundary discipline you'd apply to any untrusted input in a production system, applied here to an AI agent.
1. Run the attacks — no API key, no data fetch, no server
The attack gallery runs in-process against the real DP engine:
pip install -e .
python -m attacks.patients_alice # re-identify Alice with privacy off, then watch
# the guard, the noise, and the budget defeat it
2. Connect an AI agent
The gateway runs as an MCP stdio server for Claude Desktop. The agent becomes the untrusted query author, and it gets only structured tools ( count , sum , average , histogram , get_budget ) whose argument schemas are generated from the dataset policy. No API key is needed anywhere, because the connecting agent is the intelligence.
3. The full natural-language UI
A local single-tenant demo. An API key is needed only for the untrusted NL→query compiler:
export ANTHROPIC_API_KEY=... # used only by the untrusted NL→query compiler
docker compose up # brings up the engine, API, and UI
# open http://localhost:8501
That HTTP + Streamlit surface is a local, single-tenant demo . Identity comes from a spoofable X-Identity header, so it is meant for one trusted operator on their own machine, not a public deployment (see what these surfaces are, and are not ). For local non-Docker setup, running the tests, and configuration knobs, see SETUP.md .
The headline: an attack gallery
Anyone can claim privacy. This repository ships the exploits that would break the claim, runs them against its own engine, and pins the outcomes in CI. The fastest way to understand what the gateway guarantees is to watch it defeat three classic attacks that break naive "query a database" systems.
Attack 1 — The differencing attack
A differencing attack isolates one person by asking two aggregate questions that differ by exactly that person.
Query A: "Total income of all 100 people in department X." → $7,240,000
Query B: "Total income of all people in department X except Alice." → $7,135,000
Attacker computes: A − B = $105,000 ← Alice's exact salary, leaked.
Both queries are "just aggregates." Neither names a single row. Yet together they expose an individual. The gallery ( attacks/differencing.py ) shows this attack succeeding with privacy disabled : the target's private value is recovered exactly. (The salary sketch above is illustrative; on the real UCI Adult data, "Alice" is the unique holder of her group's maximum capital gain.) It then shows the same attack defeated once DP is on : the calibrated noise on each answer makes the subtraction useless, and the budget accountant charges for the information released across both queries rather than treating them as independent.
Attack 2 — Membership inference
A membership-inference attack determines whether a specific individual is in the dataset at all. For many datasets (a medical study, a list of defaulters), that fact is itself sensitive. An attacker with only query access tries to decide: "is this exact person in the data?"
The gallery runs this attack across a sweep of privacy budgets (ε — the dial that trades answer accuracy for privacy) and plots the result:
An optimal (Neyman–Pearson) attacker deciding one person's membership from a single noisy COUNT , run against the real engine across ε. Empirical success (blue, 95% Wilson intervals) hugs the analytic Laplace curve and stays below the worst-case DP ceiling (dashed); it collapses from certainty (DP off) toward the 0.5 coin-flip as ε shrinks. The green utility curve (right axis) shows an aggregate query's relative error over the same sweep, barely dented where the attack is defeated. Generated by python -m attacks.membership (10,000 trials per ε).
As ε shrinks (stronger privacy), the attacker's success collapses toward a coin flip. The utility overlay shows the price paid: the same noise that defeats the one-person attack barely moves a population-scale aggregate. Privacy is not free, and the chart shows exactly what you trade for it. The sweep spans ε = 8 down to 0.5 to cover the whole range; the bundled deployment charges ε = 0.05 per query , off the left edge of this chart, where the single-query attack is already indistinguishable from chance.
Attack 3 — Singling out by re-identification
Latanya Sweeney showed in 2002 that ZIP code + birth date + sex uniquely identify roughly 87% of Americans. You do not need someone's name to find them in a dataset; a few innocuous attributes will do. The gallery's third attack ( attacks/patients_alice.py ) reproduces that structure on 20 synthetic patients: sex and age alone narrow the cohort to exactly one person, Alice, the only woman over 64.
With privacy off, subtracting two perfectly ordinary demographic histograms recovers Alice's diagnosis exactly. No outlier value is required; being re-identifiable is enough, which makes this the strongest attack in the gallery. With privacy on, three independent defenses end it: the filter guard rejects the obvious narrow query outright (spending nothing); calibrated noise drowns the two-query subtraction (signal-to-noise ≈ 0.13, so the recovered "diagnosis" is essentially a random draw); and the budget refuses the averaging escalation long before it could work (~256 repetitions needed, 10 affordable). This is the attack the demo GIF at the top of this page, and the Claude Desktop walkthrough , replay live. Outcomes are regression-tested in tests/test_attack_patients.py .
These experiments, run against the system's own engine, are the core evidence that the guarantee is real and understood, not imported from a library and taken on faith.
The attacks validate the guarantee against our own math. To guard against a self-consistent-but-wrong implementation, the mechanism is also cross-checked against OpenDP as an independent reference ( attacks/crosscheck_opendp.py , opendp 0.15.1).
In short: a mechanism built from scratch agrees with the industry reference to nine decimal places, and a version that is intentionally mis-calibrated fails the same tests, so the agreement means something. Agreement is checked two ways — the sensitivity and noise scale against OpenDP's certified bound, and the sampler's actual distribution against OpenDP's, with a positive control proving the test can fail. The measured figures are in DESIGN.md section 6.3.
OpenDP is the reference, never the implementation:

[truncated]
