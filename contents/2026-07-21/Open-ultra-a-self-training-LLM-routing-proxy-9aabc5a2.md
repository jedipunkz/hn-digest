---
source: "https://github.com/numinous-technology/open-ultra"
hn_url: "https://news.ycombinator.com/item?id=48993540"
title: "Open-ultra: a self-training LLM routing proxy"
article_title: "GitHub - numinous-technology/open-ultra: Self-training LLM routing proxy for agentic CLIs: match your frontier model's quality at a fraction of the cost · GitHub"
author: "joshkolo"
captured_at: "2026-07-21T16:17:54Z"
capture_tool: "hn-digest"
hn_id: 48993540
score: 2
comments: 0
posted_at: "2026-07-21T15:26:10Z"
tags:
  - hacker-news
  - translated
---

# Open-ultra: a self-training LLM routing proxy

- HN: [48993540](https://news.ycombinator.com/item?id=48993540)
- Source: [github.com](https://github.com/numinous-technology/open-ultra)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T15:26:10Z

## Translation

タイトル: Open-ultra: 自己トレーニング LLM ルーティング プロキシ
記事のタイトル: GitHub - numinous-technology/open-ultra: エージェント CLI 用のセルフトレーニング LLM ルーティング プロキシ: わずかなコストでフロンティア モデルの品質に匹敵する · GitHub
説明: エージェント CLI 用の自己トレーニング LLM ルーティング プロキシ: わずかなコストでフロンティア モデルの品質に匹敵します - numinous-technology/open-ultra

記事本文:
GitHub - numinous-technology/open-ultra: エージェント CLI 用の自己トレーニング LLM ルーティング プロキシ: わずかなコストでフロンティア モデルの品質に匹敵する · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ああ、ああ！
読み込み中にエラーが発生しました。

このページをリロードしてください。
数え切れないほどのテクノロジー
/
オープンウルトラ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
numinous-technology/オープンウルトラ
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット プール プール src/ open_ultra src/ open_ultra テスト テスト .gitignore .gitignore .python-version .python-version ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイル ナビゲーション
エージェント CLI 用の自己トレーニング LLM ルーティング プロキシ。
わずかなコストでフロンティア モデルの出力品質に匹敵します。
open-ultra を通じてコーディング エージェント トラフィックを実行します。隠されたシャドウ レーンが安価なオープン モデルでリクエストを再実行し、エージェント グレーダーがフロンティア モデルの実際の動作と比較してリクエストをスコア付けします。安価なモデルがある種のタスクでフロンティアと一致できることが証明されると、ルーターは代わりにそのトラフィックを安価なモデルに送信し始めます。安価なモデルがうまく処理する作業については同じ品質を維持し、実際に必要な作業に対してのみフロンティア価格を支払います。ベンチマークや合成データはありません。実際の作業に合わせて調整されます。
トレーニングされたルーターが評価に合格するまでルーティングはオフのままで、安価なモデルは少なくともフロンティアと同じスコアを獲得した場所でのみトラフィックを獲得します。それまでは、open-ultra は証拠を静かに収集する透明なパススルーです。
エージェント ──► プロキシ (フロンティア モデルへのバイト パススルー、追加遅延ゼロ)
│
§─► ストア: プライマリ リクエスト/レスポンス ペア
│
└─► シャドウ レーン (サービス中は常にオンになり、ユーザーには表示されません)
§─ 応答専用モード: 同一のリクエスト バイト → シャドウ モデル
└─ worktree-exec モード: シャドウ モデルはエージェント的に動作します。
使い捨ての git ワークツリー、ツール co

そこに閉じ込められ、差分がキャプチャされました
│
エージェント採点者: シャドウとプライマリのスコア (解答と解答、
またはワークツリーの差分と実際の差分）→pairs.jsonl
│
train ──► ルーター (モデルごとの機能ヘッド) ──► ルートオン
プライマリ レーン : リクエストは認証ヘッダーを含めてバイトごとに転送されます。エージェントは以前とまったく同じように動作します。
シャドウ レーン : 同じリクエストが、ホット パスから離れたプール内の安価なモデルに送信されます。 worktree-exec モードでは、シャドウ モデルはリポジトリとエージェント ループの実際の (分離された) コピーを取得するため、単なる散文ではなくエージェントの能力が等級付けされます。
Grader : 構成可能な判定モデルは、シャドウ出力を実際に提供されたフロンティア出力と比較します。実際のルーティング決定と同じコンテキスト、同じ粒度。
Router : プール モデルごとに独立した機能ヘッドを持つ 1 つの共有分類器。適格 = 閾値を上回る。プロキシは、現在稼働している最も安価な対象モデルを選択します。それ以外の場合はフロンティアを選択します。価格と在庫状況は構成であり、重みではありません。
CLI
セットアップ
確認済み
クロード・コード
openultra クロード (または ANTHROPIC_BASE_URL=http://localhost:8484 )
ライブセッションを含むストリーミング + ルート ターン
コーデックス
openultra codex (構成フラグを介して Responses-API プロバイダーを挿入します)
ライブコーデックス実行セッション
オープンコード
openultraattach opencode はプロバイダー設定を出力します。次に、ウルトラオープンコードを開きます
ライブオープンコード実行セッション
プロキシは、Anthropic メッセージと OpenAI (チャット完了 + 応答) を受信して読み上げます。 Outbound は OpenAI のみと互換性があり、OpenRouter、vLLM、Ollama、SGLang、Togetter、DeepSeek、および事実上すべてのサービング スタックをカバーします。シャドウ モデル、グレーダー、およびフロンティアはそれぞれ、 pool.json 内の単なるエンドポイント エントリです。
open-ultra は UV ツール (PyPI 上) です。 CLI をインストールします (PATH に openultra として配置されます)。
#PyPIから
UVツールのインストールopen-ultra
# または

リポジトリから直接
uvツールインストール「git+https://github.com/numinous-technology/open-ultra」
# またはローカル チェックアウトから (編集可能、開発用)
git clone https://github.com/numinous-technology/open-ultra
CDオープンウルトラ
uv ツールのインストール --editable 。
uv ツール upgrade open-ultra でアップグレードし、uv ツール アンインストール open-ultra で削除します。
openultra init # ワンタイム: API キー、プール プリセットの選択、接続チェック
openultra claude # デーモンを自動起動し、添付されたクロード コードを起動します
# ... しばらくは正常に動作します ...
openultra レポート # モデルごとの良品率、予測される節約額
openultra train # 蓄積されたペアにルーターを適合させる
openultra Route on # 対象となるトラフィックを安価なモデルにルーティングし始める
openultra init は ~/.openultra/{env,config.json,pool.json} を書き込みます。プリセット: 無料 (試用できる料金制限付きの 0 ドルのモデル)、低予算 (DeepSeek V4 Flash)、バランス型 (Flash + MiniMax M3)。
openultra up はデーモンを実行しますが、自動的には何も接続されません。CLI は、そのベース URL がプロキシを指している場合にのみ接続されます。 3 つのレベル:
セッションごと: openultra claude / openultra codex / openultra opencode。必要に応じてデーモンを起動し、そのプロセスのみに挿入されたベース URL を使用して CLI を起動するシン ラッパー。プレーンクロードはまだバイパスします。
永続的: openultraattach claude は、ANTHROPIC_BASE_URL を ~/.claude/settings.json (Codex config.toml 、OpenCode config と同様) に書き込みます。今後の実行はすべて添付されます。 openultra detach クロードはそれを元に戻します。
手動/任意のアプリ: openultra up を実行し、http://localhost:8484 にあるものを指定します。デーモンは /v1/messages (Anthropic) および /v1/chat/completions + /v1/responses (OpenAI) を話すため、エージェント CLI だけでなく、base_url オーバーライドを持つあらゆる SDK が機能します。
動詞
何をするのか
クロード / コーデックス / オープンコード
プロキシに接続されているエージェントを起動します (デーモンを自動起動します)。
上へ

/ダウン
デーモンを実行/停止します (パススルーの場合のみ --no-shadow)
ステータス
レーン、リクエスト、ルーティング数、収集されたペア
モデルは <slug> を追加します
影モデルを追加します。 OpenRouter は価格/コンテキストを自動入力しません。 --endpoint 経由の OpenAI 互換エンドポイント
モデルリスト / rm / 有効 / 無効
プールエントリを管理します。 --laneshadow|route|両方ともデータ収集とライブルーティングを個別に切り替えます
報告書
受領額: ターンタイプごとのモデルごとの良品率、予測される節約額
電車
モデルごとの機能ヘッド、時間分割評価、高精度ゲートに適合します。各実行はバージョンです ( --name 、 --no-activate を除くと自動アクティブ化されます)
ルーターリスト/使用/rm
ルーターのバージョンを管理します。アクティブなホットスワップを使用します (即時ロールバック)
ルートのオン/オフ
ライブ ルーティングを有効にします (バージョンを選択するには --router <name>)。モデルがゲートを通過するまで拒否されます。失敗したルーティングされたコールは透過的にフロンティアにフォールバックします
影のオン/オフ
シャドウレーンを切り替える
<エージェント> のアタッチ / デタッチ
永続的なアタッチ (クロード: 自動設定編集、コーデックス/オープンコード: 正確な構成を出力)
すべてのコマンドとサブコマンドは -h/--help をサポートしています。
~/.openultra/
pool.json # フロンティア + シャドウ モデル: エンドポイント、価格、有効化
Traffic.jsonl # プライマリ レーンのリクエスト/レスポンス ログ
pairs.jsonl # {リクエストハッシュ、シャドウモデル、モード、グレード、マージン、コスト}
router/ # トレーニングされた重み、しきい値、評価レポート
設計の不変条件
バイト正確なトレーニング データ: ルーターは、ルーティングするリテラル リクエストに基づいてトレーニングします。再構築や合成プロンプトはありません。
デフォルトはフロンティアです: 安価なモデルは、ワークロードの段階的な証拠を通じてトラフィックを獲得します。最悪の場合は現状維持です。
プールは構成であり、コードを記述する必要はありません。価格、エンドポイント、可用性は pool.json にあります。モデルを追加するということは、モデルをしばらくシャドウイングしてから再トレーニングすることを意味します。
目に見えないシャドウ : シャドウ レーンがセッションやファイル (ワークツリー) に触れることはありません。

は使い捨てで隔離されており、デフォルトではネットワークがオフになっています）、または遅延。
ローカルファースト: トラフィックは構成したエンドポイントにのみ送信されます。ログはマシン上に残ります。
シャドウ レーンはリクエスト量を約 2 倍にしますが、モデルの価格はフロンティア価格の 1 ～ 3% です。 Worktree-exec モードでは、ローカル CPU とディスクが追加されます。 pool.json の支出上限はシャドウ バジェットを制限します。キャップに当たるとレーンが一時停止します。
src/open_ultra/
cli.py # すべての動詞
init_wizard.py #openultra init
paths.py # ~/.openultra の場所 (OPENULTRA_HOME オーバーライド)
代理/
server.py # インバウンド Anthropic + OpenAI 方言、パススルー + キャプチャ + ルーティング
Route_exec.py # 呼び出し元の方言で安価なモデルからルーティングされたリクエストを処理します
parse.py # キャプチャされた応答を正規化します (プレーン + SSE、すべての方言)
shadow/runner.py # キャプチャされたリクエストをシャドウ モデルで再生します
Grader/judge.py # ブラインド A/B 採点 ->pairs.jsonl
client/translate.py # Anthropic / 応答 -> チャット完了
render.py # 共有リクエストのレンダリング (採点者とルーターが同じテキストを参照)
ルーター/
train.py # モデルごとの機能ヘッド、会話グループ化の分割、評価ゲート
infer.py # バージョン管理されたルーターの読み込み + ルーティングの決定
eval/report.py # 件のレシート
control.py # デーモンのライフサイクル + 設定の編集
testing/e2e.py # オフライン 23 チェック エンドツーエンド スイート
開発
UV同期
uv run python testing/e2e.py # オフラインのエンドツーエンド テスト: アップストリームのモック、分離された OPENULTRA_HOME
uv build # ビルドホイール + sdist を dist/
データ ディレクトリを再配置するには OPENULTRA_HOME を設定し、デーモンがデフォルト以外のアドレスで実行される場合は OPENULTRA_URL を設定します。
公開 (メンテナー): uv ビルド && uv パブリッシュ ( UV_PUBLISH_TOKEN の PyPI トークンが必要)。
フロンティアへのバイト正確なパススルー、追加遅延ゼロ
人間的なメッセージの受信 (クロード・コード)
OpenAI チャット完了インバウンド
OpenAI 応答受信 (Codex)
ストリーミング (S

SE) クライアントをバッファリングせずにすべての方言をキャプチャします
Traffic.jsonl へのリクエスト/レスポンスのキャプチャ (ヘッダーは永続化されません)
Localhost のみのバインド、0600 の ~/.openultra/env
バックグラウンド ワーカーによって排出される常時オンのシャドウ キュー
Anthropic → OpenAI 互換リクエスト変換 (システム、ツール、tool_use/tool_result)
応答 → チャット完了の翻訳
モデルごとのコスト見積もり、試行制限付きの一時的なエラーの再試行
レーンごとの有効化フラグ (シャドウとルート、独立)
Worktree-exec グレーディング モード (シャドウ モデルは分離された git ワークツリーでエージェント的に実行されます)
レーンを一時停止するシャドウ支出上限
ブラインド、順序シャッフルによる A/B 判定 →pairs.jsonl
ツール呼び出しを認識したスコアリング (散文だけでなく、提案されたアクションを比較)
次のターンのグラウンド トゥルースを使用する遅延グレーディング (テストは失敗し、ユーザーが修正)
モデルごとの機能ヘッド (モデルごとに 1 つの分類子、独立したシグモイド)
会話をグループ化した時間分割 (ほぼ重複したリークなし)
精度ゲート: モデルは eval 精度をクリアした場合にのみルーティング可能です
最も安価で適格かつ利用可能なルーティング、重みではなく構成からのコスト順
バージョン管理されたルーター ( train --name 、ルーター list/show/note/use/rm )、ホットスワップ + ロールバック
障害時の透過的なフロンティアフォールバックを備えたライブルーティング
機能強化・小L

[切り捨てられた]

## Original Extract

Self-training LLM routing proxy for agentic CLIs: match your frontier model's quality at a fraction of the cost - numinous-technology/open-ultra

GitHub - numinous-technology/open-ultra: Self-training LLM routing proxy for agentic CLIs: match your frontier model's quality at a fraction of the cost · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
numinous-technology
/
open-ultra
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
numinous-technology/open-ultra
master Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits pools pools src/ open_ultra src/ open_ultra tests tests .gitignore .gitignore .python-version .python-version LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A self-training LLM routing proxy for agentic CLIs.
Match your frontier model's output quality at a fraction of the cost.
Run your coding-agent traffic through open-ultra. A hidden shadow lane re-runs your requests on cheap open models, an agentic grader scores them against what your frontier model actually did, and once a cheap model has proven it can match the frontier on a kind of task, the router starts sending that traffic to the cheap model instead. You keep the same quality on the work cheap models handle well, and only pay frontier prices for the work that actually needs it. No benchmarks, no synthetic data: it calibrates to your live work.
Routing stays off until a trained router passes evaluation, and a cheap model only earns traffic where it scored at least as well as the frontier. Until then open-ultra is a transparent passthrough that quietly collects evidence.
your agent ──► proxy (byte passthrough to your frontier model, zero added latency)
│
├─► store: primary request/response pairs
│
└─► shadow lane (always on while serving, invisible to you)
├─ response-only mode: identical request bytes → shadow model(s)
└─ worktree-exec mode: shadow model works agentically in a
throwaway git worktree, tools confined there, diff captured
│
agentic grader: scores shadow vs primary (answer vs answer,
or worktree diff vs your real diff) → pairs.jsonl
│
train ──► router (per-model capability heads) ──► route on
Primary lane : your request is forwarded byte-for-byte, auth headers included. Your agent behaves exactly as before.
Shadow lane : the same request goes to the cheap models in your pool, off the hot path. In worktree-exec mode the shadow model gets a real (isolated) copy of your repo and an agent loop, so agentic ability is graded, not just prose.
Grader : a configurable judge model compares shadow output against the frontier output that actually served you. Same context, same granularity as real routing decisions.
Router : one shared classifier with an independent capability head per pool model. Eligible = heads above threshold; the proxy picks the cheapest eligible model that is currently up, otherwise the frontier. Prices and availability are config, never weights.
CLI
Setup
Verified
Claude Code
openultra claude (or ANTHROPIC_BASE_URL=http://localhost:8484 )
live session incl. streamed + routed turns
Codex
openultra codex (injects a Responses-API provider via config flags)
live codex exec session
OpenCode
openultra attach opencode prints the provider config; then openultra opencode
live opencode run session
The proxy speaks Anthropic Messages and OpenAI (Chat Completions + Responses) inbound. Outbound is OpenAI-compatible only, which covers OpenRouter, vLLM, Ollama, SGLang, Together, DeepSeek, and effectively every serving stack. Shadow models, the grader, and the frontier are each just an endpoint entry in pool.json .
open-ultra is a uv tool ( on PyPI ). Install the CLI (it lands on your PATH as openultra ):
# from PyPI
uv tool install open-ultra
# or straight from the repo
uv tool install " git+https://github.com/numinous-technology/open-ultra "
# or from a local checkout (editable, for development)
git clone https://github.com/numinous-technology/open-ultra
cd open-ultra
uv tool install --editable .
Upgrade with uv tool upgrade open-ultra , remove with uv tool uninstall open-ultra .
openultra init # one-time: API key, pick a pool preset, connectivity check
openultra claude # auto-starts the daemon, launches Claude Code attached
# ... work normally for a while ...
openultra report # per-model as-good rates, projected savings
openultra train # fit the router on accumulated pairs
openultra route on # start routing eligible traffic to cheap models
openultra init writes ~/.openultra/{env,config.json,pool.json} . Presets: free (rate-limited $0 models to try it out), budget (DeepSeek V4 Flash), balanced (Flash + MiniMax M3).
openultra up runs the daemon, but nothing attaches automatically: a CLI is attached only when its base URL points at the proxy. Three levels:
Per session : openultra claude / openultra codex / openultra opencode . Thin wrappers that start the daemon if needed and launch the CLI with the base URL injected for that process only. Plain claude still bypasses.
Persistent : openultra attach claude writes ANTHROPIC_BASE_URL into ~/.claude/settings.json (similarly Codex config.toml , OpenCode config). Every future run attaches. openultra detach claude undoes it.
Manual / any app : run openultra up and point anything at http://localhost:8484 . The daemon speaks /v1/messages (Anthropic) and /v1/chat/completions + /v1/responses (OpenAI), so any SDK with a base_url override works, not just agent CLIs.
Verb
What it does
claude / codex / opencode
Launch that agent attached to the proxy (auto-starts the daemon)
up / down
Run / stop the daemon ( --no-shadow for passthrough only)
status
Lanes, requests, routed count, pairs collected
models add <slug>
Add a shadow model. OpenRouter slugs autofill price/context; any OpenAI-compatible endpoint via --endpoint
models list / rm / enable / disable
Manage pool entries. --lane shadow|route|both toggles data collection and live routing independently
report
Receipts: per-model as-good rate by turn type, projected savings
train
Fit per-model capability heads, time-split eval, precision gate. Each run is a version ( --name , auto-activated unless --no-activate )
routers list / use / rm
Manage router versions; use hot-swaps the active one (instant rollback)
route on/off
Enable live routing ( --router <name> to pick a version). Refused until a model passes the gate; failed routed calls fall back to frontier transparently
shadow on/off
Toggle the shadow lane
attach / detach <agent>
Persistent attach (claude: automatic settings edit; codex/opencode: prints exact config)
Every command and subcommand supports -h/--help .
~/.openultra/
pool.json # frontier + shadow models: endpoint, price, enabled
traffic.jsonl # primary lane request/response log
pairs.jsonl # {request_hash, shadow_model, mode, grade, margin, costs}
router/ # trained weights, thresholds, eval report
Design invariants
Byte-exact training data : the router trains on the literal requests it will route. No reconstruction, no synthetic prompts.
Default to frontier : cheap models earn traffic through graded evidence on your workload. Worst case is the status quo.
Pool is config, never code : prices, endpoints, and availability live in pool.json. Adding a model means shadowing it for a while, then retraining.
Invisible shadow : the shadow lane never touches your session, your files (worktrees are throwaway and isolated, network off by default), or your latency.
Local first : traffic goes only to endpoints you configured. Logs stay on your machine.
The shadow lane roughly doubles request volume, but on models that cost 1-3% of frontier prices. Worktree-exec mode adds local CPU and disk. A spend cap in pool.json bounds the shadow budget; the lane pauses when the cap is hit.
src/open_ultra/
cli.py # all verbs
init_wizard.py # openultra init
paths.py # ~/.openultra locations (OPENULTRA_HOME override)
proxy/
server.py # inbound Anthropic + OpenAI dialects, passthrough + capture + routing
route_exec.py # serve a routed request from a cheap model in the caller's dialect
parse.py # normalize captured responses (plain + SSE, all dialects)
shadow/runner.py # replay captured requests on shadow models
grader/judge.py # blinded A/B grading -> pairs.jsonl
clients/translate.py # Anthropic / Responses -> Chat Completions
render.py # shared request rendering (grader + router see the same text)
router/
train.py # per-model capability heads, conversation-grouped split, eval gate
infer.py # versioned router loading + routing decision
eval/report.py # receipts
control.py # daemon lifecycle + config edits
tests/e2e.py # offline 23-check end-to-end suite
Development
uv sync
uv run python tests/e2e.py # offline end-to-end test: mock upstream, isolated OPENULTRA_HOME
uv build # build wheel + sdist into dist/
Set OPENULTRA_HOME to relocate the data directory, OPENULTRA_URL if the daemon runs on a non-default address.
Publishing (maintainers): uv build && uv publish (needs a PyPI token in UV_PUBLISH_TOKEN ).
Byte-exact passthrough to the frontier, zero added latency
Anthropic Messages inbound (Claude Code)
OpenAI Chat Completions inbound
OpenAI Responses inbound (Codex)
Streaming (SSE) capture for every dialect, tee'd without buffering the client
Request/response capture to traffic.jsonl (headers never persisted)
Localhost-only bind, ~/.openultra/env at 0600
Always-on shadow queue drained by a background worker
Anthropic → OpenAI-compatible request translation (system, tools, tool_use/tool_result)
Responses → Chat Completions translation
Per-model cost estimation, transient-error retries with attempt cap
Per-lane enable flags (shadow vs route, independent)
Worktree-exec grading mode (shadow model runs agentically in an isolated git worktree)
Shadow spend cap that pauses the lane
Blinded, order-shuffled A/B judging → pairs.jsonl
Tool-call-aware scoring (compares proposed actions, not just prose)
Delayed grading that uses next-turn ground truth (tests failed, user corrected)
Per-model capability heads (one classifier, independent sigmoid per model)
Conversation-grouped time split (no near-duplicate leakage)
Precision gate: a model is routable only if it clears eval precision
Cheapest-eligible-and-available routing, cost order from config not weights
Versioned routers ( train --name , routers list/show/note/use/rm ), hot-swap + rollback
Live routing with transparent frontier fallback on any failure
Stronger features / small-L

[truncated]
