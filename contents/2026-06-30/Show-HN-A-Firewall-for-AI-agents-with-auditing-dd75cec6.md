---
source: "https://github.com/beebeeVB/trajeckt/"
hn_url: "https://news.ycombinator.com/item?id=48726867"
title: "Show HN: A Firewall for AI agents with auditing"
article_title: "GitHub - beebeeVB/trajeckt: A causal firewall for AI agents: blocks multi-step tool-call chains that leak data, even when every call is individually allowed. · GitHub"
author: "beebeeVB"
captured_at: "2026-06-30T00:31:05Z"
capture_tool: "hn-digest"
hn_id: 48726867
score: 1
comments: 0
posted_at: "2026-06-29T23:52:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A Firewall for AI agents with auditing

- HN: [48726867](https://news.ycombinator.com/item?id=48726867)
- Source: [github.com](https://github.com/beebeeVB/trajeckt/)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T23:52:48Z

## Translation

タイトル: Show HN: 監査機能を備えた AI エージェント用のファイアウォール
記事のタイトル: GitHub - beebeeVB/trajeckt: AI エージェント用の因果ファイアウォール: すべての呼び出しが個別に許可されている場合でも、データを漏洩する複数ステップのツール呼び出しチェーンをブロックします。 · GitHub
説明: AI エージェント用の因果ファイアウォール: すべての呼び出しが個別に許可されている場合でも、データ漏洩を引き起こす複数ステップのツール呼び出しチェーンをブロックします。 - beebeeVB/trajeckt
HN テキスト: こんにちは、インターネット上にはエージェントがますます増えています。セキュリティは大きな問題になるだろう。現在、この問題は LLM を使用してエージェントを保護することで解決されていますが、これにより幻覚と遅延の問題が発生するため、5 ミリ秒未満で動作するファイアウォールを Rust でコーディングしました。これは、計画を作成し、その計画を実施することで機能します。アクション呼び出しごとの場合は、モデル コンテキスト プロトコル リストの使用が強制され、シーケンスの場合はすべてのツール呼び出しとデータ フローが追跡されます。また、エージェントがユーザー コンテキストの外側で何かを読み取った場合にフラグを立てて、セキュリティ メカニズムを追加する汚染メカニズムもあります。 DAG を使用して機能します。

記事本文:
GitHub - beebeeVB/trajeckt: AI エージェント用の因果ファイアウォール: すべての呼び出しが個別に許可されている場合でも、データを漏洩する複数ステップのツール呼び出しチェーンをブロックします。 · GitHub
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
ビービーVB
/
軌跡
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
118 コミット 118 コミット .github .github ベンチ ベンチ 構成 構成 デプロイメント デプロイメント ドキュメント ドキュメント ドキュメント ドキュメントの例 例 フロントエンド/ .vite/ deps_temp_c018749a フロントエンド/ .vite/ deps_temp_c018749a スクリプト スクリプト sdk-python sdk-python sdk sdk src src テスト テスト ベンダー ベンダー .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_REVIEW.md CODE_REVIEW.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DOCKER.md DOCKER.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.mdbaseline_traj.txtbaseline_traj.txtbaseline_trajeckt.txtbaseline_trajeckt.txtdeployment_test_output.mddeployment_test_output.mddocker-compose.ymldocker-compose.ymlnamed_suites.txtnamed_suites.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用のランタイム強制ゲートウェイ。複数段階のエクスプロイトをブロックします。
すべてのアクションごとのセキュリティ チェックが失敗します - 決定的に、最大 1.6 ミリ秒以内に、
エージェントの到達範囲。
データベースの読み取りは許可されます。電子メールの送信は許可されています。その中でそれらを行うのは
命令はデータの引き出しです。市場に出回っているすべての認証システムは、次のことをチェックします。
アクションは一度に実行できるため、シーケンスはすぐに実行されます。 trajeckt は各呼び出しをチェックします
エージェントが蓄積した軌跡全体とデータの流れに対して
それを介して – したがって、漏洩はそれを完了するステップでブロックされます。
ただし、そのステップ自体は合法であるように見えます。
実際のゲートウェイの適用を確認する最も速い方法は Docker です。それは、
強制ゲートウェイとモック MCP アップストリーム、traj コンパイラがバンドルされている
イメージでは、シールされたコミットメントの強制がすぐに有効になっています - いいえ
手書きの設定。
docker-compose up --build

次に、強制が有効であることを確認し、セッションを開始します。
カール -s http://localhost:7777/healthz | jq
# {"ステータス":"ok",...,"commitment_capable":true}
MCP を話すエージェントに、実際の MCP ではなく http://localhost:7777 を指定します。
サーバー。フルスモークテスト - 機密読み取りが許可され、その後、外部書き込みが許可されます。
HTTP 403 でブロックされた流出を完了します — DOCKER.md にあります。
scripts/smoke_test.sh はそれをエンドツーエンドで実行し、両方の結果をアサートします。
最初のビルドには数分かかります (Rust、ゲートウェイのコンパイル、
ソースからのコンパイラ)。その後、数秒以内にスタックが開始されます。
ドッカーがいないのですか？ 1 つのコマンドで核となるアイデアを確認する
Docker を使用せずに、依存関係を持たずに中央の洞察だけを確認したい場合は、
サーバー、キー、ネットワークなし — スタンドアロンの例を実行します。
カーゴラン -- デモの例
同じ 3 ステップのエージェントを 2 回実行します。まず業界のやり方
今日: 各アクションは個別に評価され、3 つすべてが許可され、顧客レコード
消えた。次に、trajeckt の強制により次のようになります。
TRAJEKTORYD (J_t 因果強制)
───────────────
エージェント A: read_database → 許可
エージェント B: 要約 → 許可
エージェント C: send_email_external → ブロック
J_t 因果関係パスが検出されました:
d_customer_records → Summary → d_summary → external_sink
理由: 機密データが禁止されたシンクに到達しました
結果: 流出は実行前にブロックされました。
同じエージェント、同じツール、同じ法的に見える通話。 trajectt はデータがどこにあるかを追跡します
から来たものであり、汚染されたデータが禁止されたシンクに到達することを拒否します - いかなる方法であっても
間にはきれいに見える階段がたくさんあります。 (このスタンドアロンの例では、
ヒューリスティックな安全フロア。上記の Docker パスは完全な sealed-commitment を実行します
エンジン。）
エージェントを実行する前に

s、その認可された軌道は小さな仕様で宣言されており、
グラフにコンパイルされ、 HMAC でシールされます。封印されると、そのグラフは
権威。すべてのツール呼び出しは、現在到達可能なフロンティアに対してチェックされます
シールされたグラフ、フェイルクローズ — グラフにないアクションはハード拒否されます。
また、グラフがインストールされていないセッションは、評価が実行される前に拒否されます。の
エージェントのコンテキスト ウィンドウは侵害されたものとして扱われます。執行保留状態
エージェントが連絡できません。
→ 完全なアーキテクチャとシールドされたコミットメントのフローは以下の通りです。
密封されたプレセッションを強制する AI エージェント用のランタイム強制ゲートウェイ
コミットメント: エージェントの許可された軌道は、事前に宣言され、封印されます。
実行が開始され、ゲートウェイは決定的にフェールクローズを強制します。
セッション内のすべてのツール呼び出しのシールされたグラフ。
ほとんどのエージェント ガバナンスは、各ツール呼び出しを独立して処理します。ポリシー エンジンは次のことを確認します。
（エージェント、ツール、引数）、許可または拒否を返します。そのフレームは構造的には
アクション間の順序やデータフローに存在する違反を確認できません。
trajeckt はこれに 2 つのレベルで対処します。
プライマリ: 封印されたコミットメントの強制。ツール呼び出しが許可される前に、
ゲートウェイには、暗号化された署名付きのシールされた CompiledGraph (Gτ) が必要です。
エージェントがどのツールを呼び出すことができるかについて、オペレーターが承認した宣言。
順序、およびどのデータがシンクされるか。一度封印されれば、グラフは権威となる。毎
ツール呼び出しは、シールされたグラフの現在到達可能なフロンティアに対してチェックされます。
(タイプ V の強制)、来歴制約 (タイプ II)、および汚染の伝播。
封印されたグラフにないアクションは固く拒否されます。何もインストールされていないセッション
グラフは、評価が実行される前にハード拒否されます。
安全フロア: ヒューリスティック シーケンス検出器。オペレーターが明示的に選択した場合
コミットメントの範囲外

強制 (ポリシーでは、allow_uncommitted: true)、
ゲートウェイは大まかな動作パターンのヒューリスティックにフォールバックします: 流出
シーケンス ( ReadSensitive → ExternalWrite ) とコマンド アンド コントロール チェーン
( ShellExec → NetworkEgress ) が検出され、ブロックされます。これは最後の手段です
製品ではなくネットです。を使用せずに実行すると、既知の悪いパターンが検出されます。
封印されたグラフ。ヒューリスティックの範囲内にある違反は検出されません。
盲点。
trajectoryd up --upstream http://your-mcp-server
--upstream のみが必要です。 trajectoryd up は常にコミットモードで実行されます。
auto_commitment=on — ゲートウェイはツール/リストのハンドシェイクから Gτ をシールします。
ツール呼び出しが許可される前に。シールされたグラフはヒューリスティック モードではありません。
これは、宣言されたツールから自動的に導出されるコミットメント グラフの強制です。
セット。
require_commitment=on — なしでツール/呼び出しに到達するセッション
インストールされたグラフは、サイレントではなく、明示的なブロック理由によりハード的に拒否されます
ヒューリスティックに格下げされました。
MCP を話すエージェントに、実際のエージェントではなく http://localhost:7777 を指定します。
MCPサーバー。ゲートウェイはリッスン アドレスとcurl /healthz verifyを出力します。
起動時のコマンド。
docker-構成アップ
次に、MCP を話すエージェントに http://localhost:7777 を指定します。バンドルされたもの
デフォルトでは、fake-mcp-server がアップストリームになります。詳細については、DOCKER.md を参照してください。
カーゴビルド --release
./target/release/trajectoryd up \
--upstream http://your-mcp-server
または、コーパス、TLS、および予算設定の明示的なポリシー ファイルを使用します。
./target/release/trajectoryd up \
--upstream http://your-mcp-server \
--policy configs/trajectory-policy.yaml.example
上級: 手書きのコミットメント
Zero-config は、宣言されたすべてのツールをカバーする寛容なグラフを導出します。オペレーター
正確なツールの範囲、ハード予算の上限、または順序付けられたタスクを事前に宣言する必要がある人
フェーズはできる

コミットメントを手書きで作成する:
docker-compose up --build
configs/trajectory-policy.yaml.example にあるサンプルポリシーには 1 つのキーセットが必要です
traj commit で使用するものと一致させるには:
コーパス:
signed_key : " trajeckt-demo-v1-32byte-key12345 " # 以下の --key と一致する必要があります
2. traj commit でコミットメントを作成し、署名する
cd /path/to/traj && Cargo build --release
エクスポート TRAJ_BIN=/path/to/traj/target/release/traj
MCP ツール リストとタスクの説明からポリシーをシールします。
DEMO_KEY_HEX= $( python3 -c " print('trajeckt-demo-v1-32byte-key12345'.encode().hex()) " )
$TRAJ_BIN コミット \
--tools-list mcp_tools_list.json \
--task " 顧客チケットを読み取り、社内のケースメモを作成します " \
-- 予算 = 0、通話 = 20、秒 = 60 \
--out ポリシー.sealed.json \
--キー " $DEMO_KEY_HEX " \
--はい
mcp_tools_list.json は、MCP サーバーからのツール/リストの応答です。
traj commit はツールリストをタスク関連ツールに絞り込み、スコープを導き出します。
自動的にグラフをシールします。
高度な使用の場合 (明示的に手動で作成された BoundaryCommitment JSON
read_scope 、 write_scope 、 data_sinks 、および Budget ):
$TRAJ_BIN コミット境界コミットメント.json \
--registry registry.json \
--キー " $DEMO_KEY_HEX " \
-o ポリシー.sealed.json
3. シールドグラフをインストールします
python -m trajeckt.install \
--ゲートウェイ http://localhost:7777 \
--graph ポリシー.sealed.json
または Python から:
軌跡から。インストール・インポート・installed_session
session_id = インストール済みセッション (
ゲートウェイ_url = "http://localhost:7777" ,
sealed_graph_path = "policy.sealed.json" ,
)
print ( session_id ) # これを下の make_http_wrapper に渡します
installed_session() は UUID v4 セッション ID を生成し、グラフをインストールし、
ID を返すので、呼び出し側はコピーアンドペーストせずにそれを make_http_wrapper に渡します。
ステップ。
4. エージェントを同じ session_id でラップします。
trajeckt_langgraph から make_http_wrapper をインポート

ランググラフから。事前構築済みインポートツールノード
ツールノード = ツールノード (
工具、
Wrap_tool_call = make_http_wrapper (
"http://localhost:7777" ,
session_id = session_id , # 上記のinstalled_session()より
）、
)
session_id は、installed_session() によって返されるものと一致する必要があります。
ID が一致しない場合は、ゲートウェイにそのセッションのシールされたグラフがないことを意味します。
これにより、すべてのツール呼び出しがハード拒否されます。リクエストは次のようにブロックされます。
黙ってヒューリスティックに戻るのではなく、「コミットメントはインストールされていません」。
これは意図的なものです。不一致は常に構成エラーです。
Examples/quickstart_enforced_agent.py は、1 つの session_id を介してスレッドします。
インストール全体 → ループを強制し、許可された呼び出しが次のいずれかの場合にゼロ以外で終了します。
ブロックされるか、安全でない呼び出しが許可されます (CI としても機能します):
python3 の例/quickstart_enforced_agent.py \
--ゲートウェイ http://localhost:7777
コミットメント姿勢とオプトアウト
require_commitment_before_tools
自動コミットメント
行動
true (デフォルト)
true (デフォルト)
ゼロ構成コミット: コミットメントは tools/list に自動インストールされ、強制されます。ツール/コールなしで到達したセッションはハードブロックされます。
本当の
偽
手動コミット: オペレーターは traj commit --install を使用してグラフをインストールします。グラフのないセッションは大音量でブロックされます。
偽
どれでも
ヒューリスティックフロアモード。開始するにはポリシーでallow_uncommitted: trueが必要です。それ以外の場合、ゲートウェイは拒否します

[切り捨てられた]

## Original Extract

A causal firewall for AI agents: blocks multi-step tool-call chains that leak data, even when every call is individually allowed. - beebeeVB/trajeckt

Hi all, As there are more and more agents in the internet; Security is going to be a big problem. Currently, the problem is solved using a LLM to guard Agent but this creates the problem of hallucination and latency, so I coded a firewall in rust that runs under five miliseconds. This works by creating a plan and enforcing the plan; for per action call, this enforces using the Model context protocols list and for sequence it tracks every single tool call and data flow; there is also a taint mechanism where if the agent reads something outside of the user context, it flags and adds more security mechanism. It works by using a DAG.

GitHub - beebeeVB/trajeckt: A causal firewall for AI agents: blocks multi-step tool-call chains that leak data, even when every call is individually allowed. · GitHub
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
beebeeVB
/
trajeckt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
118 Commits 118 Commits .github .github benches benches configs configs deployments deployments doc doc docs docs examples examples frontend/ .vite/ deps_temp_c018749a frontend/ .vite/ deps_temp_c018749a scripts scripts sdk-python sdk-python sdk sdk src src tests tests vendor vendor .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_REVIEW.md CODE_REVIEW.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DOCKER.md DOCKER.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md baseline_traj.txt baseline_traj.txt baseline_trajeckt.txt baseline_trajeckt.txt deployment_test_output.md deployment_test_output.md docker-compose.yml docker-compose.yml named_suites.txt named_suites.txt View all files Repository files navigation
A runtime enforcement gateway for AI agents. It blocks multi-step exploits that
every per-action security check misses — deterministically, in ~1.6ms, outside the
agent's reach.
Reading a database is allowed. Sending an email is allowed. Doing them in that
order is data exfiltration. Every authorization system on the market checks one
action at a time, so the sequence walks right through. trajeckt checks each call
against the whole trajectory the agent has accumulated and the data flowing
through it — so the exfiltration is blocked at the step that completes it, even
though that step looks legal on its own.
The fastest way to see the real gateway enforce is Docker. It brings up the
enforcement gateway plus a mock MCP upstream, with the traj compiler bundled
in the image so sealed-commitment enforcement is live out of the box — no
hand-authored config.
docker-compose up --build
Then confirm enforcement is live and drive a session:
curl -s http://localhost:7777/healthz | jq
# {"status":"ok",...,"commitment_capable":true}
Point any MCP-speaking agent at http://localhost:7777 instead of your real MCP
server. A full smoke test — sensitive read allowed, then the external write that
completes an exfiltration blocked with HTTP 403 — is in DOCKER.md ,
and scripts/smoke_test.sh runs it end to end and asserts both outcomes.
The first build takes a few minutes (Rust, compiling the gateway and the
compiler from source). After that the stack starts in seconds.
No Docker? See the core idea in one command
If you just want to see the central insight with zero dependencies — no Docker,
no server, no keys, no network — run the standalone example:
cargo run --example demo
It runs the same three-step agent twice. First the way the industry does it
today: each action evaluated in isolation, all three allowed, customer records
gone. Then through trajeckt's enforcement:
TRAJEKTORYD (J_t causal enforcement)
───────────────────────────────────────────
Agent A: read_database → ALLOW
Agent B: summarize → ALLOW
Agent C: send_email_external → BLOCK
J_t causal path detected:
d_customer_records → summarize → d_summary → external_sink
Reason: sensitive data reached forbidden sink
Result: exfiltration blocked before execution.
Same agent, same tools, same legal-looking calls. trajeckt tracks where the data
came from and refuses to let tainted data reach a forbidden sink — no matter how
many clean-looking steps sit in between. (This standalone example exercises the
heuristic safety floor; the Docker path above runs the full sealed-commitment
engine.)
Before the agent runs, its authorized trajectory is declared in a small spec,
compiled into a graph, and sealed with an HMAC . Once sealed, that graph is
the authority. Every tool call is checked against the current reachable frontier
of the sealed graph, fail-closed — an action not in the graph is hard-refused,
and a session with no installed graph is refused before any evaluation runs. The
agent's context window is treated as compromised; enforcement holds state the
agent can't reach.
→ Full architecture and the sealed-commitment flow below.
A runtime enforcement gateway for AI agents that enforces sealed pre-session
commitments : the agent's authorized trajectory is declared and sealed before
execution starts, and the gateway enforces deterministically fail-closed against
that sealed graph for every tool call in the session.
Most agent governance treats each tool call independently: a policy engine sees
(agent, tool, arguments) and returns allow or deny. That framing structurally
cannot see violations that live in the ordering and data-flow between actions.
trajeckt addresses this at two levels:
Primary: sealed commitment enforcement. Before any tool call is allowed,
the gateway requires a sealed CompiledGraph (Gτ) — a cryptographically signed,
operator-approved declaration of exactly which tools the agent may call, in which
order, and to which data sinks. Once sealed, the graph is the authority. Every
tool call is checked against the current reachable frontier of the sealed graph
(Type V enforcement), provenance constraints (Type II), and taint propagation.
An action not in the sealed graph is hard-refused. A session with no installed
graph is hard-refused before any evaluation runs.
Safety floor: heuristic sequence detector. When an operator explicitly opts
out of commitment enforcement ( allow_uncommitted: true in the policy), the
gateway falls back to coarse behavioral-pattern heuristics: exfiltration
sequences ( ReadSensitive → ExternalWrite ) and command-and-control chains
( ShellExec → NetworkEgress ) are detected and blocked. This is a last-resort
net, not the product. It catches known-bad patterns when running without a
sealed graph; it does not detect violations that stay within the heuristic's
blind spots.
trajectoryd up --upstream http://your-mcp-server
Only --upstream is required. trajectoryd up always runs in committed mode:
auto_commitment=on — the gateway seals a Gτ from the tools/list handshake
before any tool call is allowed. The sealed graph is not heuristic mode;
it is commitment graph enforcement derived automatically from the declared tool
set.
require_commitment=on — any session that reaches tools/call without an
installed graph is hard-refused with an explicit block reason, not silently
downgraded to heuristics.
Point your MCP-speaking agent at http://localhost:7777 instead of your real
MCP server. The gateway prints the listen address and a curl /healthz verify
command on startup.
docker-compose up
Then point your MCP-speaking agent at http://localhost:7777 . The bundled
fake-mcp-server is the upstream by default. See DOCKER.md for details.
cargo build --release
./target/release/trajectoryd up \
--upstream http://your-mcp-server
Or with an explicit policy file for corpus, TLS, and budget settings:
./target/release/trajectoryd up \
--upstream http://your-mcp-server \
--policy configs/trajectory-policy.yaml.example
Advanced: hand-authored commitments
Zero-config derives a permissive graph covering all declared tools. Operators
who need to pre-declare exact tool scopes, hard budget caps, or ordered task
phases can hand-author a commitment:
docker-compose up --build
The example policy at configs/trajectory-policy.yaml.example needs one key set
to match what you use with traj commit :
corpus :
signing_key : " trajeckt-demo-v1-32byte-key12345 " # must match --key below
2. Author and seal a commitment with traj commit
cd /path/to/traj && cargo build --release
export TRAJ_BIN=/path/to/traj/target/release/traj
Seal a policy from the MCP tool list and a task description:
DEMO_KEY_HEX= $( python3 -c " print('trajeckt-demo-v1-32byte-key12345'.encode().hex()) " )
$TRAJ_BIN commit \
--tools-list mcp_tools_list.json \
--task " Read customer tickets and write internal case notes " \
--budget money=0,calls=20,secs=60 \
--out policy.sealed.json \
--key " $DEMO_KEY_HEX " \
--yes
mcp_tools_list.json is the tools/list response from your MCP server.
traj commit narrows the tool list to the task-relevant tools, derives scope
automatically, and seals the graph.
For advanced use (hand-authored BoundaryCommitment JSON with explicit
read_scope , write_scope , data_sinks , and budget ):
$TRAJ_BIN commit-boundary commitment.json \
--registry registry.json \
--key " $DEMO_KEY_HEX " \
-o policy.sealed.json
3. Install the sealed graph
python -m trajeckt.install \
--gateway http://localhost:7777 \
--graph policy.sealed.json
Or from Python:
from trajeckt . install import installed_session
session_id = installed_session (
gateway_url = "http://localhost:7777" ,
sealed_graph_path = "policy.sealed.json" ,
)
print ( session_id ) # pass this to make_http_wrapper below
installed_session() generates a UUID v4 session ID, installs the graph, and
returns the ID so callers pass it to make_http_wrapper without a copy-paste
step.
4. Wrap your agent with the same session_id
from trajeckt_langgraph import make_http_wrapper
from langgraph . prebuilt import ToolNode
tool_node = ToolNode (
tools ,
wrap_tool_call = make_http_wrapper (
"http://localhost:7777" ,
session_id = session_id , # from installed_session() above
),
)
session_id must match the one returned by installed_session() .
A mismatched ID means the gateway has no sealed graph for that session,
which now hard-refuses all tool calls — the request is blocked with
"no commitment installed" rather than silently falling back to heuristics.
This is intentional: a mismatch is always a configuration error.
examples/quickstart_enforced_agent.py threads one session_id through the
entire install → enforce loop and exits non-zero if either the allowed call is
blocked or the unsafe call is allowed (doubles as CI):
python3 examples/quickstart_enforced_agent.py \
--gateway http://localhost:7777
Commitment posture and opt-out
require_commitment_before_tools
auto_commitment
behavior
true ( default )
true ( default )
Zero-config committed: commitment auto-installed on tools/list , then enforced. A session that reaches tools/call without one is hard-blocked .
true
false
Manual-commit: operator installs graphs with traj commit --install ; sessions without a graph are blocked loudly.
false
any
Heuristic floor mode. Requires allow_uncommitted: true in the policy to start; otherwise the gateway refuses to

[truncated]
