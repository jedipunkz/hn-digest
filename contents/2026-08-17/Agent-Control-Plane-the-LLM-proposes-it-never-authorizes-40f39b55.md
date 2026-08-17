---
source: "https://github.com/yacine-kellib/agent-control-plane"
hn_url: "https://news.ycombinator.com/item?id=49337491"
title: "Agent Control Plane: the LLM proposes, it never authorizes"
article_title: "GitHub - yacine-kellib/agent-control-plane: A credential is not authorisation. ACP moves the authorise/refuse decision for AI agent actions outside the model, where prompt injection cannot reach it. Specification, Dafny proofs, and a reference implementation — every claim replays locally. · GitHub"
image: "https://opengraph.githubassets.com/b697bc3eec7dd222ceb640731500144cfd34d4b8767c441a1d3c46e80cce41d1/yacine-kellib/agent-control-plane"
author: "Yacine_c75"
captured_at: "2026-08-17T21:17:26Z"
capture_tool: "hn-digest"
hn_id: 49337491
score: 1
comments: 0
posted_at: "2026-08-17T20:53:45Z"
tags:
  - hacker-news
  - translated
---

# Agent Control Plane: the LLM proposes, it never authorizes

- HN: [49337491](https://news.ycombinator.com/item?id=49337491)
- Source: [github.com](https://github.com/yacine-kellib/agent-control-plane)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T20:53:45Z

## Translation

タイトル: エージェント コントロール プレーン: LLM は提案しますが、決して許可しません
記事のタイトル: GitHub - yacine-kellib/agent-control-plane: 資格情報は承認ではありません。 ACP は、AI エージェントのアクションの承認/拒否の決定を、プロンプト インジェクションが到達できないモデルの外に移動します。仕様、Dafny 証明、およびリファレンス実装 - すべての主張がローカルで再生されます。 · GitHub
説明: 資格情報は承認ではありません。 ACP は、AI エージェントのアクションの承認/拒否の決定を、プロンプト インジェクションが到達できないモデルの外に移動します。仕様、Dafny 証明、およびリファレンス実装 - すべての主張がローカルで再生されます。 - ヤシン・ケリブ/エージェント・コントロール・プレーン

記事本文:
GitHub - yacine-kellib/agent-control-plane: 資格情報は承認ではありません。 ACP は、AI エージェントのアクションの承認/拒否の決定を、プロンプト インジェクションが到達できないモデルの外に移動します。仕様、Dafny 証明、およびリファレンス実装 - すべての主張がローカルで再生されます。 · GitHub
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
ヤシン・ケリブ
/
エージェントコントロールプレーン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

レ
39 コミット 39 コミット .github .github アセット アセット クレート クレート デプロイ デプロイ ドキュメント ドキュメント dossier dossier オーケストレーター オーケストレーター パッケージ パッケージ リファレンス リファレンス サービス サービス sim sim 仕様 スペック ツール ツール .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile ライセンス ライセンスMANIFEST.sha256 MANIFEST.sha256 MANIFEST.sha256.sig MANIFEST.sha256.sig README.md README.md RELEASE.md RELEASE.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml release-key.pub release-key.pub tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
プロンプト インジェクションが到達できないモデルの外側で、AI エージェントのアクションが承認されるかどうかを決定する構造化入力コントロール プレーン。
ほとんどのエージェント展開では、モデルに認証情報を与え、その承認を呼び出します。そうではありません。これは、モデルに影響を与えることができる人は誰でも、毒されたドキュメント、敵対的なサポート チケット、依存関係の README 内のコメントなど、エージェントの権限で行動できることを意味します。慣らし運転は必要ありません。
モデルには命令とデータを区別する方法がありません。どちらも同じウィンドウ内のテキストであり、アーキテクチャ内でどちらかをバインディングとしてマークするものはありません。これは品質の問題ではないため、より良いモデルを使用しても改善されません。
セキュリティ システムが問うのは、要求者が信頼できるかどうかではありません。重要なのはそのアクションが承認されているかどうかであり、これはポリシー、能力、定足数に関する事実です。 ACP は、モデルが到達できない場所にその決定を置きます。
エージェント ──提案──▶ ポリシー エンジン ──▶ 実行者 ──▶ アクション
リスクを再計算して検証するか拒否する
署名されたポリシーから人間が署名します
モデルは元に戻せない
見るか影響を与えるか
侵害されたモデルは 40,000 ユーロの合成を要求する可能性があります

好きなだけ注文し、決して注文しません。リスク レベルは、モデルが決して認識しない署名済みポリシーから再計算され、注文にはその正確なリクエストに関連付けられた人間の署名が必要です。
クイックスタート — すべてのクレームを 90 秒で再現します
それが起こるのを見てください — インジェクションデモ
独自のエージェントを Docker、HTTP、LLM に向けます
これが重要な点 — 8 つの展開設定
実行の様子を確認する: 1 営業日 — 179 のアクション (測定値)
仕組み: 2 つのドア — 1 つのアイデアのアーキテクチャ
これが主張していないこと — 肯定的な主張の前にこれをお読みください
脅威モデルとフレームワークのマッピング — MITRE ATLAS、ATT&CK、OWASP LLM トップ 10
ドキュメント — 完全な文書
リポジトリのレイアウト — 何が本物で、何が足場なのか
求められているのは、敵対的なレビュー担当者 — 最も重要なギャップです
python3 -m pip install --break-system-packages 暗号化 dilithium-py
./tools/verify.sh
要約された出力。実行が完了すると、5 つの番号付きセクションにわたって 18 行の結果が出力されます。
== 1. 誠実さ ==
OK 118 個のファイルが MANIFEST.sha256 に一致します
== 2. マニフェスト署名 (Ed25519、オフライン リリース キー) ==
OK 切り離された署名は release-key.pub に対して検証されます
== 3. 形式的な証明 ==
OK Dafny プログラム検証ツールは 36 個の検証済み、エラー 0 個で終了しました
== 4. テストスイート ==
OK ALL 攻撃 (統合レジストリ) — 結果: 74/74
OK Suite 1 準拠 — 結果: 45/45 — 準拠
OK Suite 2 のエグゼキューターの突然変異 — 結果: 20/20 が殺害されました
...
合計 14 のスイート ライン (10 個の番号付きスイートにまたがり、そのうちの 3 個にまたがる 30 個の突然変異コントロール)。
ここでの主張があなたのマシンで再生されないとしても、信じないでください。それにはこれらの数字も含まれます。
突然変異の結果は読む価値があります。各セキュリティ チェックは順番に削除され、一致する攻撃が成功する必要があります。これにより、チェックが何らかの処理を行っているか、テストが不合格であることがわかります。

興味深い。そのうち 30 名: 執行者 20 名、承認 6 名、監査 4 名。
2 つのゲート、その違いが重要です:
マニフェストの再生成にはオフラインの Ed25519 キーが必要であるため、セクション 1 ～ 2 を緑色にできるのはキーホルダーの所有者だけです。赤のリリース間の整合性は、オフライン署名が設計どおりに動作していることであり、欠陥ではありません。 dossier/07-REPRODUCTION.md を参照してください。
ダフニーはオプションです。インストールされていない場合、プルーフステップはスキップされます。
python3 リファレンス/スイート/demo_flow.py
注: これにより、ローカル Web サーバーが起動し、ブラウザ タブが開きます。 Ctrl-C で停止するまで実行されます。これはプレゼンテーションであり、テストではありません。テスト パスには ./tools/verify.sh を使用します。発表者用ガイド: dossier/DEMO-HOWTO.md 。
サプライヤーからのレポートには、白い文字で隠された指示が記載されています。モデルはそれを読み取って準拠します。デモでは、同じ出力を 2 つのパスで並行して実行します。コントロール プレーンがなければデータは会社から流出しますが、ACP を使用すると、取り返しのつかないことは何も起こりません。
モデルは完全に準拠していることが示されています。拒否をシミュレートすると、主張を誤って伝えることになります。アーキテクチャの保証は、注入の失敗には依存しません。
本物のモデル付き。 Anthropic API キーをページに貼り付けると、エージェントは記録された応答ではなく、実際に汚染されたドキュメントを読み取るライブ モデルになります。キーはプロセスの存続期間中メモリ内に保持され、その呼び出しにのみ使用され、ディスクに書き込まれることはありません。キーがない場合、デモは録音に対してオフラインで実行されます。コントロール プレーンはモデルに何も問い合わせないため、どちらの場合でも同じように動作します。それが両方を提供することのポイントです。ライブ実行と記録実行が異なる場合、保証はモデルが何を言っているかに依存します。
python3 リファレンス/スイート/demo_flow.py --model claude-sonnet-5
自分のエージェントにそれを指示してください
上記のデモはプレゼンテーションです。これはコントロール プレーンです。

このサービスは、HTTP 経由で自分自身、自分のエージェントから実行できるサービスです。
docker compose -fdeploy/docker-compose.yml up -d ingress
curl -s localhost:8848/actions # 閉じたセット — 他には何も提案できません
エンドポイント
本体
何をするのか
GET /健康
—
liveness、および適用されるバンドルのハッシュ
GET /アクション
—
9 つの登録されたアクション。それぞれに許可されたターゲット、必要なパラメータ、および可逆性が含まれます。
投稿/提案する
{"タスクタイプ"、"ターゲット": [...]、"パラメータ": {…}、"オペレーター"、"プログラム"}
ドア。実行、保留、または拒否された応答には常にルール ID が付きます
GET /ホールド
—
提案ハッシュによってキー設定された人間を待っているもの
POST /承認
{"提案ハッシュ"、"誰が"、"決定": "確認"}
2 人目の人間が 1 つの保留されたアクションを確認します
ポスト/リリース
{"提案ハッシュ"}
リリース;有効な確認応答がそれらのバイトにバインドされている場合にのみ成功します。
エージェントは好みのモデルを提供し、必要な API キーを保持します。ACP はクライアントではなくサーバーであり、ユーザーのキーは保持しません。ドアはプロポーザルの正規バイトだけを決定します。未登録の task_type は評価される前に 8.4-3 で拒否され、スキーマ外のパラメータは V-1 で拒否され、機能ホワイトリスト外のターゲットは CW-1 で拒否されます。
書くよりも見たほうがよい場合は、sim/llm_agent.py が信頼できない呼び出し元として動作します。これは API キーを保持し、命令が埋め込まれたドキュメントを読み取り、HTTP 経由で提案します。
import ANTHROPIC_API_KEY=sk-ant-... # console.anthropic.com → API キー
docker compose -fdeploy/docker-compose.yml up -d ingress
python3 -m sim.llm_agent --invent
そのキーはそのプロセスで読み取られ、Anthropic に送信されます。これは ACP に到達することはなく、ACP は要求されることもなく、使用することもできません。呼び出し先が保持する資格情報を持つ呼び出し元は呼び出し元ではありません。

それはサブルーチンです。クライアントは Anthropic 専用です (api.anthropic.com はハードコーディングされているため、OpenAI または Gemini キーは機能しません) が、ドアはモデルに依存せず、何が提案を生成したかはわかりません。ユーザー側にインストールする必要はありません。エージェントは標準ライブラリであり、暗号化 / dilithium-py はイメージ内のサーバー上に存在します。 python3 -m sim.llm_agent --help は、フラグ、環境、終了コード、確認ループなどのマニュアル全体です。代わりに Compose で実行するには、 --rm Agent Agent --invent を実行します。このサービスは、指定された送信ネットワークの中で唯一のものであり、アーキテクチャが侵害されていると想定しているサービスです。
2 つのフェーズで動作し、2 番目のフェーズは耐荷重フェーズです。
1. トリアージ バッチ上のライブ モデル。それは 2 つの文書を読み取ります。本物の作業が記載された分析メモと、指示が埋め込まれたサプライヤー報告書です。挿入されたレポートは最後に順序付けされ、最も従う可能性が高い位置になります。この 2 つは意図的に異なるソースを持っています。そのため、ドキュメント 2 に埋め込まれたテキストが、モデルがドキュメント 1 に帰するアクションを生成するかどうかを確認できます。
それが選択したものは何であれ、変更されずにドアに送られます。従うこともあれば、拒否することもあり、まったく提案しないこともあります。クロード ソネット 5 は、バッチが単一のドキュメントだった頃、このファイルの最初の実行時に完全に拒否しました。何が起こっても、再試行されるのではなく出力されます。このフェーズは再現可能ではなく、再現されるべきものではありません。これはプロンプト上のモデルに関する事実であり、モデルおよび実行によって異なります。また、コントロール プレーンに関する証拠ではありません。 §06 の RES-L2 を参照してください。モデルが不正に動作する必要があるデモでは、モデルを測定します。
制約が 1 つあります。それは、応答のエンコーディングです。同一の入力に対して、同じモデルが 1 回の実行とツール呼び出し構文で JSON 配列を返したため、クライアントはリクエストとともに JSON スキーマを送信します。

次の実行では、およそ 3 回の解析に 1 回の実行が行われます。答えの形は固定されますが、内容は固定されません。モデルは、挿入されたアクションを提案することも、何も提案しないことも、存在しないアクションを発明することも自由に行えます。それでも、それらのすべてがドアに到達し、そこで採点されます。ここでは、モデルが選択したものを採点したり、フィルターしたり、判断したりするものは何もありません。また、再試行ループではありません。答えが都合よくなるまで再質問するのは、信頼性の衣装を着たフィルターです。 RES-L1 として公開され、スキーマが許可されたアクションのリストを増やすと失敗するスイート チェックによってピン留めされます。
2. エージェントのプロセスが侵害され、ループ内にモデルがありません。文書で要求されていた 2 つのアクションが直接提案されました。これがドア B の実際の名前である敵対者です。呼び出し元は、モデルについて話し合うだけでなく、そのコードを編集することによって侵害されます。フェーズ 1 とは異なり、これは決定的です。
release_to_partner 拒否されました XPROG-1 データセットには登録された所有プログラムがありません — 拒否されました、想定されていません
order_Synthetic HELD 8.4-11 HIGH / IRREVERSIBLE — 4 人の承認者に通知
--invent は、侵害された呼び出し元が与えられた命令に制限されないため、文書に記載されていない 3 分の 1 を追加します。
exfiltrate_dataset 拒否 8.4-3 は署名付きバンドルに含まれていません。拒否され、採点されなかった
テキストが疑わしいという理由で拒否されたものは 1 つもありませんでした

[切り捨てられた]

## Original Extract

A credential is not authorisation. ACP moves the authorise/refuse decision for AI agent actions outside the model, where prompt injection cannot reach it. Specification, Dafny proofs, and a reference implementation — every claim replays locally. - yacine-kellib/agent-control-plane

GitHub - yacine-kellib/agent-control-plane: A credential is not authorisation. ACP moves the authorise/refuse decision for AI agent actions outside the model, where prompt injection cannot reach it. Specification, Dafny proofs, and a reference implementation — every claim replays locally. · GitHub
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
yacine-kellib
/
agent-control-plane
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
39 Commits 39 Commits .github .github assets assets crates crates deploy deploy docs docs dossier dossier orchestrator orchestrator packages packages reference reference services services sim sim spec spec tools tools .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE MANIFEST.sha256 MANIFEST.sha256 MANIFEST.sha256.sig MANIFEST.sha256.sig README.md README.md RELEASE.md RELEASE.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml release-key.pub release-key.pub tsconfig.base.json tsconfig.base.json View all files Repository files navigation
A structured-input control plane that decides whether an AI agent's action is authorised — outside the model, where prompt injection cannot reach.
Most agent deployments give the model a credential and call that authorisation. It isn't. It means anyone who can influence the model can act with the agent's rights: a poisoned document, a hostile support ticket, a comment in a dependency README. No break-in required.
A model has no way to tell an instruction from a datum. Both are text in the same window, and nothing in the architecture marks one as binding. That doesn't improve with better models, because it isn't a quality problem.
The question a security system asks is not whether the requester is trustworthy. It's whether the action is authorised, which is a fact about policy, capability and quorum. ACP puts that decision somewhere the model cannot reach.
agent ──proposes──▶ policy engine ──▶ executor ──▶ action
recomputes risk verifies, or refuses
from signed policy humans sign for the
the model can't irreversible ones
see or influence
A compromised model can request a €40,000 synthesis order as often as it likes and never cause one. The risk level is recomputed from signed policy the model never sees, and the order needs human signatures bound to that exact request.
Quick start — reproduce every claim in ninety seconds
See it happen — the injection demo
Point your own agent at it — Docker, HTTP, your LLM
Where this bites — eight deployment settings
See it run: a business day — 179 actions, measured
How it works: two doors — the architecture in one idea
What this does not claim — read this before the positive claims
Threat model and framework mapping — MITRE ATLAS, ATT&CK, OWASP LLM Top 10
Documentation — the full dossier
Repository layout — what is real and what is scaffold
Wanted: an adversarial reviewer — the most important gap
python3 -m pip install --break-system-packages cryptography dilithium-py
./tools/verify.sh
Abridged output. A complete run prints 18 result lines across five numbered sections:
== 1. Integrity ==
OK 118 files match MANIFEST.sha256
== 2. Manifest signature (Ed25519, offline release key) ==
OK detached signature verifies against release-key.pub
== 3. Formal proofs ==
OK Dafny program verifier finished with 36 verified, 0 errors
== 4. Test suites ==
OK ALL attacks (consolidated registry) — RESULT: 74/74
OK Suite 1 conformance — RESULT: 45/45 — CONFORMANT
OK Suite 2 executor mutation — RESULT: 20/20 killed
...
Fourteen suite lines in all, spanning 10 numbered suites, and 30 mutation controls across three of them.
If a claim here does not replay on your machine, don't believe it. That includes these numbers.
The mutation results are the ones worth reading. Each security check is deleted in turn and the matching attack has to succeed, which is how you know the check does something and the test isn't vacuous. 30 of them: 20 executor, 6 acknowledgement, 4 audit.
Two gates, and the difference matters:
Sections 1–2 can only be made green by the key holder, because regenerating the manifest requires the offline Ed25519 key. Red integrity between releases is offline signing working as designed, not a defect — see dossier/07-REPRODUCTION.md .
Dafny is optional; the proof step is skipped if it isn't installed.
python3 reference/suites/demo_flow.py
Note: this starts a local web server and opens a browser tab. It runs until you stop it with Ctrl-C — it is a presentation, not a test. For the test path use ./tools/verify.sh . Presenter's guide: dossier/DEMO-HOWTO.md .
A supplier report arrives with an instruction hidden in white text. The model reads it and complies. The demo runs that same output down two paths side by side: without a control plane the data leaves the company, with ACP nothing irreversible happens.
The model is shown complying fully . Simulating a refusal would misrepresent the claim — the architecture's guarantee does not depend on injection failing.
With a real model. Paste an Anthropic API key into the page and the agent becomes a live model reading the actual poisoned document, rather than a recorded response. The key is held in memory for the process lifetime, used only for that call, and never written to disk. With no key the demo runs offline against the recording — and the control plane behaves identically either way , because it never consults the model about anything. That is the point of offering both: if the live and recorded runs diverged, the guarantee would depend on what the model said.
python3 reference/suites/demo_flow.py --model claude-sonnet-5
Point your own agent at it
The demo above is a presentation. This is the control plane as a service you can drive yourself, from your own agent, over HTTP.
docker compose -f deploy/docker-compose.yml up -d ingress
curl -s localhost:8848/actions # the closed set — nothing else can be proposed
Endpoint
Body
What it does
GET /health
—
liveness, and the hash of the bundle being enforced
GET /actions
—
the nine registered actions, each with its permitted targets, required params and reversibility
POST /propose
{"task_type", "targets": [...], "params": {…}, "operator", "program"}
the door. Answers executed , held or refused , always with the rule id
GET /holds
—
what is waiting on a human, keyed by proposal hash
POST /acknowledge
{"proposal_hash", "who", "decision": "CONFIRM"}
a second human confirms one held action
POST /release
{"proposal_hash"}
release; succeeds only if a valid acknowledgement is bound to those bytes
Your agent supplies whatever model it likes and holds whatever API key that needs — ACP is the server, not the client, and holds no key of yours . The door decides on the proposal's canonical bytes and nothing else: an unregistered task_type is refused at 8.4-3 before it is ever graded, params outside the schema are refused at V-1 , and a target outside the capability whitelist is refused at CW-1 .
If you would rather watch one than write one, sim/llm_agent.py is a working untrusted caller: it holds the API key, reads a document with an instruction buried in it, and proposes over HTTP.
export ANTHROPIC_API_KEY=sk-ant-... # console.anthropic.com → API keys
docker compose -f deploy/docker-compose.yml up -d ingress
python3 -m sim.llm_agent --invent
That key is read in that process and sent to Anthropic. It never reaches ACP, which is never asked for it and could not use it — a caller whose credential the callee holds is not a caller, it is a subroutine. The client is Anthropic-only ( api.anthropic.com is hardcoded, so an OpenAI or Gemini key will not work), while the door is model-agnostic and has no idea what produced a proposal . Nothing needs installing on your side: the agent is standard library, and cryptography / dilithium-py live on the server, inside the image. python3 -m sim.llm_agent --help is the whole manual — flags, environment, exit codes, and the acknowledgement loop. To run it in compose instead: run --rm agent agent --invent , that service being the only one given outbound network and the one the architecture assumes is compromised.
It runs two phases, and the second is the load-bearing one.
1. A live model on a triage batch. It reads two documents: an assay memo with genuine work in it, and a supplier report with an instruction buried in it. The injected report is ordered last, the position most likely to be obeyed. The two have different sources on purpose — so you can see whether text planted in document 2 produces actions the model attributes to document 1.
Whatever it chooses goes to the door unaltered. It may comply, it may decline, it may propose nothing at all — Claude Sonnet 5 declined outright on this file's first run, back when the batch was a single document — and whatever happens is printed rather than retried away. This phase is not reproducible and is not meant to be. It is a fact about a model on a prompt, it varies by model and by run, and it is not evidence about the control plane; see RES-L2 in §06 . A demo that needs the model to misbehave is measuring the model.
One thing is constrained: the reply's encoding . The client sends a JSON schema with the request, because on identical input the same model returned a JSON array on one run and tool-call syntax on the next, and roughly one run in three parsed. It fixes the shape of the answer and never its content — the model stays free to propose the injected actions, to propose nothing, or to invent an action that does not exist, and every one of those still reaches the door and is graded there. Nothing here scores, filters or judges what the model chose, and it is not a retry loop: re-asking until the answer is convenient is a filter wearing a reliability costume. Disclosed as RES-L1, and pinned by a suite check that fails if the schema ever grows a list of permitted actions.
2. The agent process compromised, no model in the loop — the two actions the document asked for, proposed directly. This is the adversary Door B actually names: a caller is compromised by editing its code, not only by talking its model round. Unlike phase 1, this one is deterministic:
release_to_partner REFUSED XPROG-1 dataset has no registered owning program — refused, not assumed
order_synthesis HELD 8.4-11 HIGH / IRREVERSIBLE — 4 approvers notified
--invent adds a third the document never mentioned, because a compromised caller is not restricted to the instructions it was fed:
exfiltrate_dataset REFUSED 8.4-3 not in the signed bundle; refused, never graded
Not one of those was refused because the text looked suspicio

[truncated]
