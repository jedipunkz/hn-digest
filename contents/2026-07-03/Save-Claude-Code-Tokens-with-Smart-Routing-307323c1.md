---
source: "https://github.com/regolo-ai/brick-SR1"
hn_url: "https://news.ycombinator.com/item?id=48780858"
title: "Save Claude Code Tokens with Smart Routing"
article_title: "GitHub - regolo-ai/brick-SR1: brick is a smart AI Models router, based on complexity & capabilities extraction from the query to the models via proprietary spatial embedding algorythm · GitHub"
author: "FrancescoMassa"
captured_at: "2026-07-03T22:57:33Z"
capture_tool: "hn-digest"
hn_id: 48780858
score: 2
comments: 0
posted_at: "2026-07-03T22:39:24Z"
tags:
  - hacker-news
  - translated
---

# Save Claude Code Tokens with Smart Routing

- HN: [48780858](https://news.ycombinator.com/item?id=48780858)
- Source: [github.com](https://github.com/regolo-ai/brick-SR1)
- Score: 2
- Comments: 0
- Posted: 2026-07-03T22:39:24Z

## Translation

タイトル: スマート ルーティングによるクロード コード トークンの保存
記事のタイトル: GitHub - regolo-ai/brick-SR1: ブリックは、独自の空間埋め込みアルゴリズムを介したクエリからモデルへの複雑さと機能の抽出に基づいたスマート AI モデル ルーター · GitHub
説明: ブリックは、独自の空間埋め込みアルゴリズムによるモデルへのクエリからの複雑さと機能の抽出に基づいたスマート AI モデル ルーターです - regolo-ai/brick-SR1

記事本文:
GitHub - regolo-ai/brick-SR1: ブリックは、独自の空間埋め込みアルゴリズムを介したクエリからモデルへの複雑さと機能の抽出に基づいた、スマート AI モデル ルーターです · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。お願いします

このページをリロードします。
レゴロアイ
/
レンガ-SR1
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ ワークフロー .github/ ワークフロー アプリ アプリ デプロイ デプロイ docs docs evals evals パッケージ パッケージ スクリプト スクリプト .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml ライセンス ライセンス Makefile Makefile NOTICE NOTICE README.md README.md config.yaml config.yaml package-lock.json package-lock.json package.json package.json 価格設定.yaml 価格設定.yaml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1 つのクエリ、1 つのエンドポイント、地球上のすべての LLM。
Brick は、Mixture-of-Models (MoM) ルーティング ゲートウェイです。各プロンプトを読み取ります
機能と複雑さを考慮して、プール内の最適なバックエンドにルーティングします。
オープンウェイトおよびクローズウェイト LLM は、最強の単一モデルの品質に匹敵します。
そのコストの一部。カスケードはありません。無駄な通話はありません。ドロップイン モデル: "レンガ" 。
Brick をいつ使用するか · クイックスタート · Brick を使用する理由 · Claude Code · Codex · FAQ · ベンチマーク · 仕組み · 論文
Brick は、複数のモデルに対して実行する人、または 1 つの強力なモデルに対して定額料金を支払う人向けです。よくある 3 つのケース:
モデルのプールがあり、各クエリが適切なクエリに到達するようにしたいと考えています。
安価なプロンプトが最も高価なモデルを焼き尽くしてはいけませんし、ハードなプロンプトが小規模なモデルで枯渇してはいけません。 Brick はクエリごとに機能と複雑さを読み取り、それに応じてディスパッチするため、プールは手動で選択するのではなく、1 つの段階的なシステムとして機能します。
品質を損なうことなく、クロード コード/コーデックスのコストを削減したいと考えています。
Brick をコーディング エージェントと ev の前に置きます

すべてのリクエストは、実際にジョブを実行できる最も安価なモデルにルーティングされ、タスクで必要な場合にのみエスカレーションされます。同じ UX を維持し、簡単なターンではなく、困難なターンに対して料金を支払います。
さまざまなモデルを 1 つのツールで統合したいと考えています。
単一の OpenAI 互換エンドポイントを介して、Claude Code または Codex 内から OpenAI モデル、GLM、DeepSeek、Kimi、Qwen などを使用します。 config.yaml でプールを一度定義し、どこでも model: "brick" を呼び出します。
現在、最も高速に動作するパスは CLI です。CLI はルーターを自己ホストし、ルーターに接続します。
クロード・コードをあなたに。ノード 18 以上と Docker が必要です。
git クローン https://github.com/regolo-ai/brick-SR1.git
cd render-SR1/apps/cli && npm install && npm run build && npm link
rick claude on # ルーターを起動 + ANTHROPIC_BASE_URL を ~/.claude/settings.json に配線
次に、新しいクロード コード セッションを開き、/model ピッカーで ブリック クロードを選択します。
すべてのリクエストは、機能と複雑さによって俳句/ソネット/作品にルーティングされるようになりました。参照
モード、エフォートピッカー、ライブ用の Brick + Claude コード
ブリック クロード ステータス ダッシュボード。
Docker イメージが公開されると (「配布チャネル」を参照)、
ゲートウェイを直接実行できるようになります。
docker run --rm -p 18000:18000 \
-e REGOLO_API_KEY= $REGOLO_API_KEY \
ghcr.io/regolo-ai/brick:latest # 次の v2.1.0 タグで公開されます
次に、他の OpenAI エンドポイントと同じように呼び出して、 "model": "brick" を設定するだけです。
カール http://localhost:18000/v1/chat/completions \
-H " 認可: ベアラー $REGOLO_API_KEY " \
-H " Content-Type: application/json " \
-d ' {"モデル":"ブリック","メッセージ":[{"ロール":"ユーザー","コンテンツ":"sqrt(2) が非合理であることを証明する"}]} '
x-selected-model 応答ヘッダーは、Brick がどのバックエンドを選択したかを示します。
その数学プロンプトは推論モデルにルーティングされます。 「こんにちは」は最も安いものへのルートです。
それまでは、ブリック サーブ (上記の CLI から) r

同じルータをソースからローカルにアンインストールします。
単一モデル
ルートLLM
FrugalGPT / カスケード
レンガ
クエリごとに 1 つの呼び出し (カスケードの無駄なし)
✅
✅
❌
✅
能力を意識した (6 次元)
該当なし
❌ バイナリ
❌
✅
複雑さを意識した
該当なし
部分的な
✅
✅
N 個のオープン + クローズド モデルのプール
該当なし
2
少ない
✅
継続的なコスト ↔ 品質のつまみ
❌
❌
しきい値
✅ r ∈ [-1, 1]
ネイティブ マルチモーダル (画像/音声)
変化する
❌
❌
✅
ドロップインOpenAI対応
該当なし
該当なし
該当なし
✅
カスケード ルーター (FrugalGPT、Cascade Routing) は、モデルが完了するまでモデルを次々と呼び出します。
信頼性チェックに合格すると、トークンと遅延のすべてのミスに対して料金が発生します。レンガが作る
クエリごとに単一の前方決定を行うため、無駄なものはありません。
ゴスミレーター.mp4
OpenAI/Anthropic 互換のエンドポイントを 1 つクロード コードの前に置くと、Brick は機能と複雑さに基づいてすべてのリクエストを haiku 、 Sonnet 、または opus にルーティングします。クロード コードの UX を維持します。 Brick は、その仕事を実行できる最も安価なモデルを選択します。
~/.claude/settings.json の # 個のワイヤー ANTHROPIC_BASE_URL 上のブリック クロード、ルーターを自動起動します
次に:
新しいクロード コード セッションを開きます (現在のセッションは影響を受けません)。
/model ピッカーで、brick-claude を選択します (組み込みの opus/sonnet/haiku エイリアスと並んで配置されますが、置き換えられません)。
rick claude off # ANTHROPIC_BASE_URL を復元し、オプションでルーターを停止します
自動起動ルーターの代わりにすでに正常なルーターを要求するには、brick claude on --no-start を使用します。また、プロンプトなしでルーターを制御するには、brick claude off --stop / --keep を使用します。
5 つのモード: コストと品質のトレードオフを選択してください
モードとは、ブリックにいくら使うかを伝える方法です。それぞれ、簡単/中程度/難しいクエリを、最も安い ( eco 、常に Haiku) から最強 ( max 、常に opus) までのモデル層にマッピングし、その間に lite 、mid、pro があります。 1 つを選択すると、Brick が内部でクエリごとのルーティングを処理します。

それ。
2026-07-03.23-55-05.mp4
Claude Code の /model ピッカーの思考努力スライダーからモードを直接切り替えます: low は、 eco 、 middle lite 、 high mid 、 xhigh pro 、および max max を選択します。したがって、工数制御は思考予算を設定するのではなく、モデル層を選択します。ブリック クロード モードまたはブリック クロード <モード> を使用して明示的に切り替えることもできます。
ミッドがデフォルトです。 Haiku には 1M のバリアントがないため、1M コンテキストのリクエストではマップが上にシフトします。ソネットへの解決は容易および中程度で、作品への解決は困難です。
層を選択すると、ルーター自身の信号 (クエリの難易度と選択したモデルのヘッドルーム) からのリクエストごとに、思考の難易度が自動的に決定されます。
ネイティブモデルはルーターをバイパスします
ピッカーで opus 、 Sonnet 、または haiku を明示的に選択すると、Brick が完全にスキップされます。リクエストは、スキルのルーティングや労力のオーバーライドなしで、その正確なモデルにそのまま転送されます。ブリック・クロードのみがルーターを実行します。
ブリック クロード ステータス # ライブ ダッシュボード (対話型ターミナルのデフォルト)
ブリック・クロード・ステータス --once # 静的なワンショット・ビュー
ダッシュボードには、ルーターの最後の再起動以降のレポートが表示されます。
モデルごとにルーティング: モデルごとの数とパーセント。
モデルごとの労力の分布: 各モデル内で推論の労力がどのように分散されるか。
難易度の組み合わせ: ルーティングされたリクエスト全体での分類子の簡単/中/難しい判定。
エコノミー : ルーティングされたリクエスト数に対する全オペレーションと比較して、推定 ~X% の節約 (実際のトークン数とキャッシュを除いた、リクエストの組み合わせからの相対的な推定値)。
また、接続/配線状態、分類子の遅延 (平均、p50、p95)、およびフォールバック レートも表示されます。
ワークフローおよびサブエージェントと連携
ブリック ルーティングはリクエストごとに行われます。クロード コードのワークフローとサブエージェントでは、エージェントが ブリック クロード を使用している限り、各エージェントの呼び出しは独立してルーティングされるため、安価なサブエージェント タスクは Haiku に到達する一方で、ハードなサブエージェント タスクはエスカラに到達する可能性があります。

tes から opus までを同じ実行で実行します。
OpenAI Codex の背後にある同じ考え方: Brick は Codex の前に配置され、各リクエストをモデル プール全体にルーティングするため、簡単なターンでコストを削減し、1 つの OpenAI 互換エンドポイントを通じて非 OpenAI モデルで Codex を駆動できます。
ブリックコーデックスオン # ~/.codex/config.toml でモデル/モデルプロバイダーをブリックに設定し、ルーターを自動起動します
これにより、専用の Codex プロファイル (OpenAI プール スキル ルーター) が具体化され、ローカル ルーターを指すマネージド プロバイダーが追加されます。新しい Codex セッションを開始すると、Brick 経由でルーティングされるようになります。
ブリックコーデックスオフ # 以前のコーデックスモデル/プロバイダーを復元します
Codex は、Claude Code と同じ 5 つのモードとステータス ビューを公開します。
ブリック コーデックス モード # または: ブリック コーデックス エコ |ライト |ミッド |プロ |最大
ブリック コーデックス ステータス # ライブ ルーティング ダッシュボード
--no-start でブリック コーデックスを使用して、自動起動ルーターの代わりにすでに正常なルーターを要求します。 Claude ルーター スタックと Codex ルーター スタックはホスト ポート 8000 を共有するため、一度にサービスを提供できるのは 1 つだけです。配線する前にもう一方を停止してください。
コーディングエージェントは必要ありません。 Brick は、任意のクライアント、スクリプト、アプリから呼び出すことができるプレーンな OpenAI 互換ゲートウェイです。
ブリック サーブ # docker は http://localhost:18000 で構成します
ブリック チャット # ローカル ルーターに対する TUI チャット
ブリックルート " 2+2 とは何ですか? " # プロンプトのルーティング決定を出力します。呼び出しは行われません
他の OpenAI エンドポイントと同様に呼び出して、 "model": "brick" を設定するだけです。
カール http://localhost:18000/v1/chat/completions \
-H " 認可: ベアラー $REGOLO_API_KEY " \
-H " Content-Type: application/json " \
-d ' {"モデル":"ブリック","メッセージ":[{"ロール":"ユーザー","コンテンツ":"sqrt(2) が非合理であることを証明する"}]} '
x-selected-model 応答ヘッダーは、Brick がどのバックエンドを選択したかを示します。その数学プロンプトは推論モデルにルーティングされます。 「こんにちは」は最も安いものへのルートです。
config.yamlでプールを構成します。
今まで

Brick が決定することは config.yaml から得られます。コア ブロックは skill_router で、ここでプール、各モデルのスキル ベクトル、およびそのコストの重みを宣言します。
スキルルーター :
有効 : true
機能 : # すべてのクエリとモデルが存在する 6 つのディメンション
- コーディング
- クリエイティブ合成
- 指示_フォロー中
- 数学的推論
- 計画エージェント
- 世界の知識
モデル:
- モデル：「qwen3.5-9b」
skill_vector : [0.71, 0.51, 0.81, 0.91, 0.58, 0.18] # 次元ごとの能力
use_reasoning : false
cost_weight : 0.10 # 相対価格、コストバイアスを駆動します
- モデル: " deepseek-v4-flash "
スキルベクトル : [0.82, 0.66, 0.86, 0.93, 0.62, 0.49]
use_reasoning : false
コストの重み : 0.40
- モデル：「kimi2.6」
スキルベクトル : [0.90, 0.75, 0.87, 0.94, 0.64, 0.34]
use_reasoning : true
reasoning_effort : " 中 "
コストの重み : 0.60
ここに OpenAI 互換のバックエンドを追加または交換します。バックエンド自体は、provider_profiles / model_config の下で宣言されます (出荷された構成では、バックエンドはすべて Regolo を指します)。さらに 2 つのブロックを使用すると、計算に触れることなく配線を調整できます。
キーワードルール:
- name : "force_coder " # ハード オーバーライド: これらのプロンプトを特定のモデルに送信します
モード: "オーバーライド"
モデル：「kimi2.6」
演算子:「OR」
キーワード : ["デバッグ"、"リファクタリング"、"コンパイル"、"書き込み"

[切り捨てられた]

## Original Extract

brick is a smart AI Models router, based on complexity & capabilities extraction from the query to the models via proprietary spatial embedding algorythm - regolo-ai/brick-SR1

GitHub - regolo-ai/brick-SR1: brick is a smart AI Models router, based on complexity & capabilities extraction from the query to the models via proprietary spatial embedding algorythm · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
regolo-ai
/
brick-SR1
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows apps apps deploy deploy docs docs evals evals packages packages scripts scripts .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md config.yaml config.yaml package-lock.json package-lock.json package.json package.json pricing.yaml pricing.yaml pyproject.toml pyproject.toml View all files Repository files navigation
One Query, One Endpoint, Every LLM on Earth.
Brick is a Mixture-of-Models (MoM) routing gateway . It reads each prompt's
capability and complexity , then routes it to the best backend in a pool of
open- and closed-weight LLMs, matching the strongest single model's quality at a
fraction of its cost. No cascades. No wasted calls. Drop-in model: "brick" .
When to use Brick · Quickstart · Why Brick · Claude Code · Codex · FAQ · Benchmarks · How it works · Paper
Brick is for anyone running against more than one model, or paying flat rate for a single strong one. Three common cases:
You have a pool of models and want each query to reach the right one.
Cheap prompts should not burn your most expensive model, and hard prompts should not be starved on a small one. Brick reads capability and complexity per query and dispatches accordingly, so the pool works as one graded system instead of a manual pick.
You want to cut Claude Code / Codex costs without losing quality.
Put Brick in front of your coding agent and every request is routed to the cheapest model that can actually do the job, escalating only when the task needs it. You keep the same UX and pay for the hard turns, not the easy ones.
You want to unify different models behind one tool.
Use OpenAI models, GLM, DeepSeek, Kimi, Qwen and others from inside Claude Code or Codex through a single OpenAI-compatible endpoint. Define the pool once in config.yaml and call model: "brick" everywhere.
The fastest working path today is the CLI, which self-hosts the router and wires it into
Claude Code for you. Requires Node >= 18 and Docker.
git clone https://github.com/regolo-ai/brick-SR1.git
cd brick-SR1/apps/cli && npm install && npm run build && npm link
brick claude on # starts the router + wires ANTHROPIC_BASE_URL in ~/.claude/settings.json
Then open a new Claude Code session and pick brick-claude in the /model picker.
Every request now routes to haiku / sonnet / opus by capability and complexity. See
Brick + Claude Code for modes, the effort picker, and the live
brick claude status dashboard.
Once the Docker image is published (see Distribution channels ), you'll
be able to run the gateway directly:
docker run --rm -p 18000:18000 \
-e REGOLO_API_KEY= $REGOLO_API_KEY \
ghcr.io/regolo-ai/brick:latest # published at the next v2.1.0 tag
Then call it like any OpenAI endpoint, just set "model": "brick" :
curl http://localhost:18000/v1/chat/completions \
-H " Authorization: Bearer $REGOLO_API_KEY " \
-H " Content-Type: application/json " \
-d ' {"model":"brick","messages":[{"role":"user","content":"Prove that sqrt(2) is irrational"}]} '
The x-selected-model response header tells you which backend Brick picked.
That math prompt routes to a reasoning model; "Hello" routes to the cheapest one.
Until then, brick serve (from the CLI above) runs the same router locally from source.
Single model
RouteLLM
FrugalGPT / Cascade
Brick
One call per query (no cascade waste)
✅
✅
❌
✅
Capability-aware (6 dimensions)
n/a
❌ binary
❌
✅
Complexity-aware
n/a
partial
✅
✅
Pool of N open + closed models
n/a
2
few
✅
Continuous cost ↔ quality knob
❌
❌
threshold
✅ r ∈ [-1, 1]
Native multimodal (image / audio)
varies
❌
❌
✅
Drop-in OpenAI-compatible
n/a
n/a
n/a
✅
Cascade routers (FrugalGPT, Cascade Routing) call models one after another until a
confidence check passes, paying for every miss in tokens and latency. Brick makes a
single forward decision per query, so there is nothing to waste.
gosmiulator.mp4
Put one OpenAI/Anthropic-compatible endpoint in front of Claude Code, and Brick routes every request to haiku , sonnet , or opus based on capability and complexity. You keep the Claude Code UX; Brick picks the cheapest model that can do the job.
brick claude on # wires ANTHROPIC_BASE_URL in ~/.claude/settings.json, auto-starts the router
Then:
Open a new Claude Code session (your current session is unaffected).
In the /model picker, select brick-claude (it sits alongside the built-in opus/sonnet/haiku aliases, which it does not replace).
brick claude off # restores ANTHROPIC_BASE_URL, optionally stops the router
Use brick claude on --no-start to require an already-healthy router instead of auto-starting one, and brick claude off --stop / --keep to control the router without a prompt.
The 5 modes: pick your cost/quality trade-off
A mode is how you tell Brick how much to spend. Each one maps easy/medium/hard queries to a model tier, from cheapest ( eco , always haiku) to strongest ( max , always opus), with lite , mid and pro in between. Pick one and Brick handles the per-query routing inside it.
2026-07-03.23-55-05.mp4
You switch mode straight from the thinking effort slider in Claude Code's /model picker: low picks eco , medium lite , high mid , xhigh pro , and max max . So the effort control does not set a thinking budget, it selects the model tier. You can also switch explicitly with brick claude mode or brick claude <mode> .
mid is the default. On 1M-context requests the map shifts up since Haiku has no 1M variant: easy and medium resolve to sonnet, hard to opus.
Once you have picked the tier, how hard to think is decided autonomously per request from the router's own signals (query difficulty plus the chosen model's headroom).
Native models bypass the router
Selecting opus , sonnet , or haiku explicitly in the picker skips Brick entirely: the request is forwarded verbatim to that exact model, with no skill routing and no effort override. Only brick-claude runs the router.
brick claude status # live dashboard (default in an interactive terminal)
brick claude status --once # static one-shot view
The dashboard reports, since the last router restart:
Routed by model : count and percent per model.
Per-model effort distribution : how reasoning effort spread out within each model.
Difficulty mix : the classifier's easy/medium/hard verdicts across routed requests.
Economy : an estimated saved ~X% vs all-opus over the routed request count (a relative estimate from request mix, excluding real token counts and caching).
It also shows connection/wiring state, classifier latency (avg, p50, p95), and fallback rate.
Works with workflows and subagents
Brick routing is per request. In Claude Code workflows and subagents, each agent's call is routed independently as long as that agent uses brick-claude , so a cheap subagent task can land on haiku while a hard one escalates to opus in the same run.
The same idea behind OpenAI Codex: Brick sits in front of Codex and routes each request across your model pool, so you cut cost on easy turns and can drive Codex with non-OpenAI models through one OpenAI-compatible endpoint.
brick codex on # sets model/model_provider to brick in ~/.codex/config.toml, auto-starts the router
This materializes a dedicated Codex profile (the OpenAI-pool skill router) and adds a managed provider pointing at the local router. Start a new Codex session and it now routes through Brick.
brick codex off # restores your previous Codex model/provider
Codex exposes the same 5 modes and status view as Claude Code:
brick codex mode # or: brick codex eco | lite | mid | pro | max
brick codex status # live routing dashboard
Use brick codex on --no-start to require an already-healthy router instead of auto-starting one. The Claude and Codex router stacks share host port 8000, so only one can serve at a time; stop the other before wiring.
You do not need a coding agent. Brick is a plain OpenAI-compatible gateway you can call from any client, script, or app.
brick serve # docker compose up on http://localhost:18000
brick chat # TUI chat against the local router
brick route " what is 2+2? " # print the routing decision for a prompt, no call made
Call it like any OpenAI endpoint, just set "model": "brick" :
curl http://localhost:18000/v1/chat/completions \
-H " Authorization: Bearer $REGOLO_API_KEY " \
-H " Content-Type: application/json " \
-d ' {"model":"brick","messages":[{"role":"user","content":"Prove that sqrt(2) is irrational"}]} '
The x-selected-model response header tells you which backend Brick picked. That math prompt routes to a reasoning model; "Hello" routes to the cheapest one.
Configure the pool in config.yaml
Everything Brick decides comes from config.yaml . The core block is skill_router , where you declare the pool, each model's skill vector, and its cost weight:
skill_router :
enabled : true
capabilities : # the 6 dimensions every query and model live in
- coding
- creative_synthesis
- instruction_following
- math_reasoning
- planning_agentic
- world_knowledge
models :
- model : " qwen3.5-9b "
skill_vector : [0.71, 0.51, 0.81, 0.91, 0.58, 0.18] # capability per dimension
use_reasoning : false
cost_weight : 0.10 # relative price, drives the cost bias
- model : " deepseek-v4-flash "
skill_vector : [0.82, 0.66, 0.86, 0.93, 0.62, 0.49]
use_reasoning : false
cost_weight : 0.40
- model : " kimi2.6 "
skill_vector : [0.90, 0.75, 0.87, 0.94, 0.64, 0.34]
use_reasoning : true
reasoning_effort : " medium "
cost_weight : 0.60
Add or swap any OpenAI-compatible backend here; the backends themselves are declared under provider_profiles / model_config (the shipped config points them all at Regolo). Two more blocks let you nudge routing without touching the math:
keyword_rules :
- name : " force_coder " # hard override: send these prompts to a specific model
mode : " override "
model : " kimi2.6 "
operator : " OR "
keywords : ["debug", "refactor", "compile", "write

[truncated]
