---
source: "https://github.com/stefbuilds/cross-examine"
hn_url: "https://news.ycombinator.com/item?id=49270997"
title: "Cross-Examine – replay base behavior against an AI-generated PR"
article_title: "GitHub - stefbuilds/cross-examine · GitHub"
author: "stefanospalyvos"
captured_at: "2026-08-12T12:45:54Z"
capture_tool: "hn-digest"
hn_id: 49270997
score: 1
comments: 0
posted_at: "2026-08-12T11:56:28Z"
tags:
  - hacker-news
  - translated
---

# Cross-Examine – replay base behavior against an AI-generated PR

- HN: [49270997](https://news.ycombinator.com/item?id=49270997)
- Source: [github.com](https://github.com/stefbuilds/cross-examine)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T11:56:28Z

## Translation

タイトル: 反対調査 – AI が生成した PR に対して基本的な動作を再現
記事タイトル: GitHub - stefbuilds/cross-examine · GitHub
説明: GitHub でアカウントを作成して、stefbuilds/クロステスト開発に貢献します。

記事本文:
GitHub - stefbuilds/反対尋問 · GitHub
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
ステフビルド
/
反対尋問
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
118 コミット 118 コミット .github/ workflows .github/ workflows api api artifacts/ product-audit-2026-07-18 artifacts/ product-audit-2026-07-18 docs docs フロントエンド フロントエンド スクリプト スクリプト src/cross_examine src/cross_examine テスト テスト .gitignore .gitigno

re .python-version .python-version AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml要件.txt要件.txt uv.lock uv.lock vercel.json vercel.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Codex がコードを記述します。反対尋問はそれをスタンドに置きます。
OpenAI Build Week 2026 · トラック: 開発者ツール (テスト · エージェント ワークフロー)
Git ワークツリー → GPT-5.6 Sol クレーム → 信頼できる入力ベース/ヘッド実行 → 純粋な集計() → FastAPI/React レポート。
問題。エージェントが作成したコードは、存在するテストに合格します。かどうかはまだ何もチェックされていません。
置き換えられた動作はまだ保持されています。したがって、モデルは 1 つのバグを修正し、別のバグを導入し、
スイート全体が緑のままです。 Codex で作成されたプル リクエストをマージする人は誰でも、
diff には、diff がサイレントに変更された動作に関する証拠はありません。
ツール。 Cross-Examine は、Codex で作成された Python 用の独立した検証ハーネスです。
変化します。基本リビジョンの動作をキャプチャし、ヘッド リビジョンを実行します。
同じ入力を入力し、敵対的な境界を探索します。パイプラインを通過する新しく実行されたレポート
検証は、すべての VERIFIED または REFUTED の背後にある正確なコマンドとキャプチャされた出力を表示します。
発見。棄権は、代わりに試みられた証拠または確定的な診断を示します。
領収書を捏造すること。
キャッチは製品です。もっともらしい最適化では、空のリストに対して None が返されます。
既存のハッピーパス テストは緑色のままで、クロス検査は [] を使用して BROKEN を返します。
入力の再生 — 以下のコマンドにより、クリーンなチェックアウトから 60 秒で再生できます。
ジャッジクイックスタート: 60 秒でキャッチを確認
Codex と GPT-5.6 の使用方法
このリポジトリにも含まれています: 要件、ディレクトリ マップ、Windows セットアップ、実際のリポジトリの実行、テスト、ビデオ アウトライン
ジャッジのクイックスタート: キャッチを参照してください。

60秒
macOS または Linux では、新しいワークスペースを割り当て、アンビエント モデルをクリアし、ストレージを実行します。
変数を使用し、チェックインされた特性評価フィクスチャを強制します。発見はまだ続く
実際のローカル パイプラインから:
hero_workspace= $( mktemp -d )
env -u OPENAI_API_KEY -u CROSS_EXAMINE_DB -u CROSS_EXAMINE_RUNS CROSS_EXAMINE_DEMO_CHARACTERIZER=フィクスチャ \
uv run --isorated --no-editable 反対尋問デモ --no-open \
--ワークスペース " $hero_workspace "
新しいワークスペースでの最初の実行では、次のレポートがレポートされます。
特徴付け: 決定論的なヒーローの特徴
評決: 壊れた
コーパス: +2 この実行 · 合計 2
反論された主張: 空の保存
入力を再生しています: []
同じ資格情報をクリアしたコマンドを、同じ hero_workspace を使用して再度実行します。の
評決は依然として壊れたままである。コーパス出力は今回の実行で+0・合計2になります。新しいワークスペース
これにより、宣伝されている初回実行 +2 が正確になります。
製品 UI で同じ証拠を検査するには:
env -u OPENAI_API_KEY \
CROSS_EXAMINE_DB= " $hero_workspace /cross-examine.db " \
CROSS_EXAMINE_RUNS= " $hero_workspace /runs " \
UV 実行反対検査サーブ
ターミナル コマンドによって出力された実行 URL を開き、反論された結果を展開します。これ
サーバーは同じワークスペースローカルデータベースを読み取り、root を実行するため、正確なコマンド、base
そこから出力、ヘッド出力、期待値、実際値、再生入力が得られます。
パイプライン検証済みの永続レポート。
何も構築せずにテストしてください。生きた証拠の探索者
デプロイされたデモ インスタンスです。クローン、インストール、API キーはありません。それは明示的に機能します
ラベル付きでチェックインされた証拠フィクスチャなので、レポート UI、正確なコマンドの受信、および
判定面はブラウザで直接検査できます。 Vercel 関数では、
リポジトリの実行には Git およびローカル ランタイム機能が必要なので、任意のリポジトリ
分析はインテルです

理論的にはローカルのみ — 上記のクイックスタートは、実際の 5 段階のパイプラインを実行します。
サポートされているプラ​​ットフォーム。 macOS、Linux、Windows。上記のコマンドは macOS/Linux のものです。の
同等の Windows PowerShell セットアップは以下のとおりです。 CI演習
3 つすべてで Python 3.12。
Codex と GPT-5.6 の使用方法
Codex が作業を加速した場所。 Codex がアプリケーション全体を作成して反復しました。
Python パイプライン、スキーマと検証レイヤー、実行制御、SQLite
永続性、FastAPI サービス、React 証拠エクスプローラー、CLI、パッケージ化、
クロスプラットフォーム検証スクリプトとテスト スイート。簡単な作業も行いました
過小評価する — Windows cp1252 子エンコーディングの失敗、pytest-cache の診断
切り離されたワークツリーでの拒否の名前変更、および依存関係に基づく誤検出の文書化
docs/trials.md 内 — それぞれが実行ポリシーをむしろ変更しました
単なるコードではなく。 Devpost で提供される日付付きの Git 履歴と Codex セッション
提出物はその進行状況を示します。
重要な決定が行われた場所。人間は製品の権威をずっと保持していました。コーデックス
実装を選択しました。分裂は意図的であり、それが評決の理由である
信頼できる: 左側のすべての教義は、右側のコードで許可される内容を制限します。
結論として。
実行時に GPT-5.6 がどのように使用されるか。 GPT-5.6 Sol ( gpt-5.6-sol ) は、有界 diff と
ソース コンテキストを取得し、スキーマに制約された Claim とオプションの ProbePlan を発行します。決して発しない
結果または評決。不正な形式、重複した、未知のターゲット、および禁止された構造
フィールドは拒否され、提案テキストは信頼できないままになります。モデルは意図的に
裁判官ではなく拘束された構成要素 — 行動的な主張を提案する一方で、
モデルフリーの実行により証拠が提供され、純粋な決定論的な集計() によって決定が行われます。
の

製品の評決。
スキルは評価されるシステムの一部です。容疑者に陪審員になってもらうことはできません。
クロスエグザミンは、別個の状態ストアを持つ別個のプロセスです。提案および実行されます。
をチェックし、決定論的な判定関数を適用します。コーパス v1 は検証済みのレイヤー A を保持します
フィクスチャを作成し、リポジトリ ロケーターとシンボルによってそれらを再生します。
スキーマ制約のあるクレームは信頼できない提案であり、オラクルではありません。特性評価
オプションで信頼できない ProbePlan を提案する場合もあります。どちらも結果をもたらすことができない、または
評決。散文的な主張ではなく、実行された基本的な動作と決定論的なポリシーが、
保存状態の発見。以下の意図的変更の棄権規則は、同じものに基づいています。
境界線。
---
構成:
テーマ: ベース
レイアウト:ダグレ
テーマ変数:
フォントサイズ: 14px
線の色: '#9ca3af'
プライマリテキストカラー: '#111827'
エッジラベル背景: '#ffffff'
tertiaryTextColor: '#4b5563'
フローチャート:
曲線: 基底
ノード間隔: 44
ランク間隔: 60
パディング: 16
htmlラベル: true
---
フローチャートTB
PR["<b>Python の差分</b><br/>ベースからヘッドまで"]:::input
サブグラフ U [「信頼できない提案」]
LR方向
I["<b>1 · 取り込み</b><br/>Git ワークツリー<br/>変更されたファイル候補"]:::untrusted
C["<b>2 · 特徴付け</b><br/>GPT-5.6 ソル<br/>主張、決して評決なし"]:::untrusted
私 --> C
終わり
サブグラフ EX [「3・CROSS-EXAMINE」]
LR方向
LA["<b>レイヤー A</b><br/>ベースキャプチャ<br/>ヘッドリプレイ"]:::grounded
LB["<b>レイヤー B</b><br/>有界仮説<br/>そして縮小"]:::grounded
RT["<b>リポジトリ テスト</b><br/>発見されたコマンド"]:::grounded
LA --> LB --> RT
終わり
AG[["<b>4 · 集約()</b><br/>純粋 · I/O なし · モデルなし"]]:::純粋
CORPUS["<b>コーパス v1</b><br/>検証済みレイヤー A フィクスチャ"]:::corpus
R["<b>5 · レンダリング</b><br/>SQLite とグラウンデッド UI"]:::レポート
広報 --> 私
C -- 「主張、証明されていない」 --> LA
RT ==> AG
ラ～。 "ピン対象フィクスチャ" .->

コーパス
コルパス - 。 「次の実行をリプレイ」 .-> LA
AG -- "批判的反論の保持" --> BROKEN(["<b>BROKEN</b>"]):::broken
AG -- 「その他の反論、批判的棄権」 --> RISKY(["<b>RISKY</b>"]):::risky
AG -- "上記のどれでもない" --> SAFE(["<b>SAFE · 制限付き</b>"]):::safe
壊れた --> R
リスキー --> R
安全 --> R
classDef 入力の塗りつぶし:#ffffff、ストローク:#9ca3af、ストローク幅:1px、色:#374151
classDef untrusted fill:#ffffff、ストローク:#b45309、ストローク幅:1px、color:#78350f
classDef 接地塗りつぶし:#ffffff、ストローク:#475569、ストローク幅:1px、色:#1e293b
classDef 純粋な塗りつぶし:#111827、ストローク:#111827、ストローク幅:1px、カラー:#ffffff
classDef コーパスの塗りつぶし:#ffffff、ストローク:#9ca3af、ストローク幅:1px、色:#4b5563
classDef レポートの塗りつぶし:#ffffff、ストローク:#374151、ストローク幅:1px、色:#111827
classDef 壊れた塗りつぶし:#ffffff、ストローク:#991b1b、ストローク幅:2px、色:#7f1d1d
classDef 危険な塗りつぶし:#ffffff、ストローク:#b45309、ストローク幅:2px、色:#7835
[切り捨てられた]
インジェストはベースを解決し、デタッチされた Git ワークツリーとカタログ クラスに進みます。
変更された Python ファイル内の関数、非同期、およびネストされた候補の定義。これは
変更された行の精度ではなく、ファイルレベルの検出。
Characterize は GPT-5.6 Sol に厳密なクレームとオプションの ProbePlan を要求します。どちらも
信頼できない提案。オフライン ヒーローでは、ラベル付きのチェックインされた Claim フィクスチャが置き換えられます。
モデルコール。
反対尋問は、キャプチャされた基本的な動作を頭部に対してリプレイし、制限付きのアクションを実行します。
エッジケースの仮説探索。 JSON 互換の同期呼び出し可能オブジェクトをプローブします。
入力。設定された外部のものは推測ではなく棄権されます。
集計は純粋な関数です。保守に批判的な反論は壊れています。その他
反論、批判的な棄権、または重要な主張の欠落は危険です。
Render は、永続化された Report を読み取ります。検証済みおよび反駁済みの調査結果は、
正確なコマンドと捕獲

出力;棄権は確定的な診断を示します。
境界と障害の動作については、docs/architecture.md を参照してください。
V1 は、提案に次のような変更がない限り、意図的な変更の正確性を意図的に控えます。
独立した実行可能なオラクル。模範的な散文は決して神託ではないので、表現されたものは
意図的変更の主張が何もない場合、レポートは少なくとも RISKY に保たれます。
SAFE は、限界があり、証明されていないことを意味します。同紙は、関係者間で何も反論されなかったと報告している。
プルリクエストが正しいかどうかをチェックするのではなく、実際に実行されたかどうかをチェックします。
クロスエグザインはターゲット リポジトリのコードを実行するため、対象となるリポジトリのみを対象とします。
信頼する。コマンドは、shell=False (実行可能な許可リスト) を指定した引数ベクトルを介して実行されます。
シークレット形式の変数、期限、2 MB の出力を取り除いた最小限の子環境
キャップとレシートの編集。これらのコントロールはターゲットをコーディングするのではなく、ハーネス自体をバインドします
スポーン — 運用環境での使用には実際の分離が必要です。 127.0.0.1 でサービスを提供します。
公的な証拠探索者はまさにそれです。
レポート UI とそのレシートを検査できるように、ラベル付きのチェックインされたフィクスチャを提供します
何もインストールせずに。リポジトリの実行には Git とローカル ランタイムが必要なので、
そのパスは設計によりローカルで実行されます。
要件
注意事項
パイソン
3.12 はテスト済みです。パッケージのメタデータでは現在、>= が許可されています

[切り捨てられた]

## Original Extract

Contribute to stefbuilds/cross-examine development by creating an account on GitHub.

GitHub - stefbuilds/cross-examine · GitHub
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
stefbuilds
/
cross-examine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
118 Commits 118 Commits .github/ workflows .github/ workflows api api artifacts/ product-audit-2026-07-18 artifacts/ product-audit-2026-07-18 docs docs frontend frontend scripts scripts src/ cross_examine src/ cross_examine tests tests .gitignore .gitignore .python-version .python-version AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt uv.lock uv.lock vercel.json vercel.json View all files Repository files navigation
Codex writes the code. Cross-Examine puts it on the stand.
OpenAI Build Week 2026 · Track: Developer tools (testing · agentic workflows)
Git worktrees → GPT-5.6 Sol claims → trusted-input base/head execution → pure aggregate() → FastAPI/React report.
The problem. Agent-authored code passes the tests that exist. Nothing yet checks whether the
behavior it replaced still holds. So, the model fixes one bug, introduces another, and the
suite stays green throughout. Anyone merging Codex-authored pull requests is reviewing a
diff with no evidence about the behavior that diff silently changed.
The tool. Cross-Examine is an independent verification harness for Codex-authored Python
changes. It captures the base revision's behavior, executes the head revision against the
same inputs, and hunts adversarial boundaries. Newly executed reports that pass pipeline
validation show the exact command and captured output behind every VERIFIED or REFUTED
finding. Abstentions show attempted evidence or a deterministic diagnostic instead of
fabricating a receipt.
The catch is the product. A plausible optimization returns None for an empty list, the
existing happy-path test stays green, and Cross-Examine returns BROKEN with [] as the
reproducing input — reproducible in 60 seconds from a clean checkout by the command below.
Judge quickstart: see the catch in 60 seconds
How Codex and GPT-5.6 were used
Also in this repo: requirements · directory map · Windows setup · real repository runs · tests · video outline
Judge quickstart: see the catch in 60 seconds
On macOS or Linux, allocate a fresh workspace, clear ambient model and run-storage
variables, and force the checked-in characterization fixture. The findings still come
from the real local pipeline:
hero_workspace= $( mktemp -d )
env -u OPENAI_API_KEY -u CROSS_EXAMINE_DB -u CROSS_EXAMINE_RUNS CROSS_EXAMINE_DEMO_CHARACTERIZER=fixture \
uv run --isolated --no-editable cross-examine demo --no-open \
--workspace " $hero_workspace "
The first run in that new workspace reports:
Characterization: deterministic hero fixture
Verdict: BROKEN
Corpus: +2 this run · 2 total
Refuted claim: preserve-empty
Reproducing input: []
Run the same credential-cleared command again with the same hero_workspace . The
verdict remains BROKEN ; corpus output becomes +0 this run · 2 total . A new workspace
is what makes the advertised first-run +2 exact.
To inspect the same evidence in the product UI:
env -u OPENAI_API_KEY \
CROSS_EXAMINE_DB= " $hero_workspace /cross-examine.db " \
CROSS_EXAMINE_RUNS= " $hero_workspace /runs " \
uv run cross-examine serve
Open the run URL printed by the terminal command, then expand the refuted finding. This
server reads the same workspace-local database and run root, so the exact command, base
output, head output, expected value, actual value, and reproducing input come from that
pipeline-validated persisted report.
Test it without building anything. The live evidence explorer
is a deployed demo instance — no clone, no install, no API key. It serves an explicitly
labeled, checked-in evidence fixture so the report UI, the exact-command receipts, and the
verdict surface can be inspected directly in a browser. Vercel Functions do not provide the
Git and local-runtime capabilities required to execute repositories, so arbitrary repository
analysis is intentionally local-only — the quickstart above runs the real five-stage pipeline.
Supported platforms. macOS, Linux, and Windows. The commands above are macOS/Linux; the
equivalent Windows PowerShell setup is below. CI exercises
Python 3.12 on all three.
How Codex and GPT-5.6 were used
Where Codex accelerated the work. Codex authored and iterated the whole application:
the Python pipeline, the schema and validation layer, execution controls, SQLite
persistence, the FastAPI service, the React evidence explorer, the CLI, packaging, the
cross-platform verification scripts, and the test suite. It also did the work that is easy
to underestimate — diagnosing a Windows cp1252 child-encoding failure, a pytest-cache
rename denial in detached worktrees, and the dependency-shaped false positives documented
in docs/trials.md — each of which changed the execution policy rather
than just the code. The dated Git history and the Codex session supplied with the Devpost
submission show that progression.
Where the key decisions were made. The human held product authority throughout; Codex
chose the implementation. The split was deliberate and is the reason the verdict is
trustworthy: every doctrine on the left constrains what the code on the right is allowed
to conclude.
How GPT-5.6 is used at run time. GPT-5.6 Sol ( gpt-5.6-sol ) reads bounded diff and
source context and emits schema-constrained Claims plus optional ProbePlans. It never emits
an outcome or a verdict. Malformed, duplicate, unknown-target, and forbidden structured
fields are rejected, and proposal text stays untrusted. The model is a deliberately
constrained component rather than the judge — it proposes behavioral claims, while
model-free execution supplies the evidence and a pure deterministic aggregate() decides
the product verdict.
A skill is part of the system being judged. You cannot ask the suspect to be the jury.
Cross-Examine is a separate process with a separate state store: it proposes and executes
checks, then applies a deterministic verdict function. Corpus v1 persists verified Layer-A
fixtures and replays them by repository locator and symbol.
A schema-constrained Claim is an untrusted proposal, not an oracle. Characterization
may also propose an optional untrusted ProbePlan ; neither can carry an outcome or
verdict. Executed base behavior and deterministic policy, not claim prose, decide a
preservation finding. The intended-change abstention rule below follows from that same
boundary.
---
config:
theme: base
layout: dagre
themeVariables:
fontSize: 14px
lineColor: '#9ca3af'
primaryTextColor: '#111827'
edgeLabelBackground: '#ffffff'
tertiaryTextColor: '#4b5563'
flowchart:
curve: basis
nodeSpacing: 44
rankSpacing: 60
padding: 16
htmlLabels: true
---
flowchart TB
PR["<b>Python diff</b><br/>base to head"]:::input
subgraph U ["UNTRUSTED PROPOSAL"]
direction LR
I["<b>1 · Ingest</b><br/>Git worktrees<br/>changed-file candidates"]:::untrusted
C["<b>2 · Characterize</b><br/>GPT-5.6 Sol<br/>claims, never a verdict"]:::untrusted
I --> C
end
subgraph EX ["3 · CROSS-EXAMINE"]
direction LR
LA["<b>Layer A</b><br/>base capture<br/>head replay"]:::grounded
LB["<b>Layer B</b><br/>bounded Hypothesis<br/>and shrink"]:::grounded
RT["<b>Repository tests</b><br/>discovered command"]:::grounded
LA --> LB --> RT
end
AG[["<b>4 · aggregate()</b><br/>pure · no I/O · no model"]]:::pure
CORPUS["<b>Corpus v1</b><br/>verified Layer-A fixtures"]:::corpus
R["<b>5 · Render</b><br/>SQLite and grounded UI"]:::report
PR --> I
C -- "claims, unproven" --> LA
RT ==> AG
LA -. "pins eligible fixtures" .-> CORPUS
CORPUS -. "replays next run" .-> LA
AG -- "preserve-critical refutation" --> BROKEN(["<b>BROKEN</b>"]):::broken
AG -- "other refutation, critical abstain" --> RISKY(["<b>RISKY</b>"]):::risky
AG -- "none of the above" --> SAFE(["<b>SAFE · bounded</b>"]):::safe
BROKEN --> R
RISKY --> R
SAFE --> R
classDef input fill:#ffffff,stroke:#9ca3af,stroke-width:1px,color:#374151
classDef untrusted fill:#ffffff,stroke:#b45309,stroke-width:1px,color:#78350f
classDef grounded fill:#ffffff,stroke:#475569,stroke-width:1px,color:#1e293b
classDef pure fill:#111827,stroke:#111827,stroke-width:1px,color:#ffffff
classDef corpus fill:#ffffff,stroke:#9ca3af,stroke-width:1px,color:#4b5563
classDef report fill:#ffffff,stroke:#374151,stroke-width:1px,color:#111827
classDef broken fill:#ffffff,stroke:#991b1b,stroke-width:2px,color:#7f1d1d
classDef risky fill:#ffffff,stroke:#b45309,stroke-width:2px,color:#7835
[truncated]
Ingest resolves base and head into detached Git worktrees and catalogues class,
function, async, and nested candidate definitions in changed Python files. This is
file-level discovery, not changed-line precision.
Characterize asks GPT-5.6 Sol for strict Claims and optional ProbePlans. Both are
untrusted proposals. In the offline hero, a labeled checked-in Claim fixture replaces
the model call.
Cross-examine replays captured base behavior against head, then runs a bounded
Hypothesis search for edge cases. It probes synchronous callables with JSON-compatible
inputs; anything outside that set abstains rather than guesses.
Aggregate is a pure function. A preserve-critical refutation is BROKEN ; other
refutations, critical abstentions, or missing critical claims are RISKY .
Render reads the persisted Report . VERIFIED and REFUTED findings open to an
exact command and captured output; abstentions show a deterministic diagnostic.
See docs/architecture.md for boundaries and failure behavior.
V1 deliberately abstains on intended-change correctness unless the proposal has an
independent executable oracle. Since model prose is never an oracle, a represented
intended-change claim without one keeps the report at least RISKY .
SAFE means bounded, not proven. It reports that nothing was refuted among the
checks that actually ran — not that the pull request is correct.
Cross-Examine executes the target repository's code, so point it only at repositories you
trust. Commands run through argument vectors with shell=False , an executable allowlist,
a minimal child environment that strips secret-shaped variables, deadlines, a 2 MB output
cap, and receipt redaction. Those controls bound the harness itself, not code the target
spawns — production use needs real isolation. Serve on 127.0.0.1 .
The public evidence explorer is exactly that: it
serves a labeled, checked-in fixture so the report UI and its receipts can be inspected
without installing anything. Executing a repository needs Git and a local runtime, so
that path runs locally by design.
Requirement
Notes
Python
3.12 is tested; package metadata currently permits >=

[truncated]
