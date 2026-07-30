---
source: "https://github.com/MaorBril/agentic"
hn_url: "https://news.ycombinator.com/item?id=49112732"
title: "Show HN: Agentic – Run Claude Code on different models with routing and budgets"
article_title: "GitHub - MaorBril/agentic: Run Claude Code on any model, with a budget — multi-provider router, LLM-triaged tier routing, spend tracking · GitHub"
author: "maorbril"
captured_at: "2026-07-30T17:15:35Z"
capture_tool: "hn-digest"
hn_id: 49112732
score: 1
comments: 0
posted_at: "2026-07-30T17:04:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agentic – Run Claude Code on different models with routing and budgets

- HN: [49112732](https://news.ycombinator.com/item?id=49112732)
- Source: [github.com](https://github.com/MaorBril/agentic)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T17:04:00Z

## Translation

タイトル: HN の表示: Agentic – ルーティングと予算を使用してさまざまなモデルでクロード コードを実行する
記事のタイトル: GitHub - MaorBril/agentic: 予算内で任意のモデルでクロード コードを実行する — マルチプロバイダー ルーター、LLM トリアージ層ルーティング、支出追跡 · GitHub
説明: 予算内で任意のモデルでクロード コードを実行 - マルチプロバイダー ルーター、LLM トリアージ層ルーティング、支出追跡 - MaorBril/agentic

記事本文:
GitHub - MaorBril/agentic: 予算内で任意のモデルでクロード コードを実行 - マルチプロバイダー ルーター、LLM トリアージ層ルーティング、支出追跡 · GitHub
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
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
マオールブリル
/
エージェント的な
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
30 コミット 30 コミット .github/ workflows .github/ workflows cmd cmd docs docs 例 例 内部 内部 .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh main.go main.go すべてのファイルを表示 リポジトリ ファイル ナビゲーション
予算に応じて、任意のモデルで Claude Code を実行します。
Agentic は、クロード コードをシン ローカル ルーターにラップします。セッションの見た目と感触はクロードとまったく同じです (同じ TUI、同じツール、同じアップデート) が、その背後にあるモデルは Anthropic、OpenAI、xAI、または OpenAI と互換性のあるもの (Ollama、vLLM、OpenRouter、DeepSeek、Groq) にすることができます。すべてのトークンは測定され、価格が設定され、設定した予算に照らしてチェックされます。
エージェント # クロード コード、デフォルトのプロファイルで追跡されます
エージェント -p 安い # 同じセッション、安価なモデル
Agentic --model grok # 1 回限りのモデル オーバーライド
代理店のコスト # 今日の $4.31 はどこに消えましたか?
なぜ
Claude Code は優れたハーネスであり、改良され続けています。フォークすることは、それを失うことを意味します。しかし、それは 1 つのプロバイダーとのみ通信し、最終的に尋ねる 2 つの質問には答えません。「そのセッションの費用はいくらですか?」そして、安い部品を安いモデルで実行できますか？
エージェントは、クロード コード自体には触れずに両方に応答します。 Claude Code は、 ANTHROPIC_BASE_URL を介したゲートウェイでのポイントを正式にサポートしています。エージェントとは、そのゲートウェイとその周囲の CLI です。
「Claude コードが必要だが、1 つのプロバイダーに固定されていない」という問題を解決する方法は 3 つあります。
短いバージョン: Agentic は、Claude Code より優れたハーネスになろうとはしていません。Agentic は、Anthropic が出荷したものとまったく同じように Claude Code を保持し、 ANTHROPIC_BASE_URL の背後にあるものを交換するだけです。まったく別のエージェント ループが必要な場合は、Op

代わりに、enCode/Crush/Goose/Aider が適切なレイヤーです。チームのために展開して管理するゲートウェイが必要な場合は、LiteLLM がより成熟した選択肢となります。 Agentic は、予算があり、1 つのコマンドでインストールされ、何も操作せずに実行できる安価なモデルの避難ハッチを備えた claude を変更せずに実行したいと考えている 1 人の開発者向けです。
Agentic (ランチャー) ──▶ クロード (無修正、自動更新)
│ ANTHROPIC_BASE_URL
▼
ローカルルーター (127.0.0.1)
§─ anthropic: バイト忠実なパススルー
§─ openai 方言: 完全なリクエスト/ストリーム翻訳
§─ 利用ログ（SQLite）＋価格
└─ バジェットゲート
▼
Anthropic · OpenAI · xAI · Ollama · vLLM · OpenRouter · ...
デーモンはありません。最初のエージェント セッションはルーター ポートをバインドし、全員にサービスを提供します。終了すると、別の実行中のセッションが数秒以内に引き継ぎます。最後のセッションが終わると照明が消えます。
モデル名はユーザーが定義するエイリアスです。クロード コードはモデル ID を不透明な文字列として扱うため、ANTHROPIC_MODEL=grok がそのまま流れ、ルーターがそれを解決します。 claude- で始まるものはすべて Anthropic untouched に渡されます。メイン モデルがまったく別のものであっても、バックグラウンド タスクは動作し続けます。
カール -fsSL https://raw.githubusercontent.com/maorbril/agentic/main/install.sh |しー
エージェントのセットアップ
またはソースから: go install github.com/maorbril/agentic@latest
後で更新するには、エージェント更新を実行します (または、エージェント更新 -- インストールせずに利用できるかどうかを確認します)。これにより、エージェント自体が更新されます。クロードコードはそれにもかかわらず、自動的に自動更新を続けます。
すべては ~/.agentic/config.yaml に存在し、ターミナルからすべて編集可能です。
エージェントプロバイダーは openai --type openai --base-url https://api.openai.com/v1 \ を追加します
--key-env OPENAI_API_KEY --max-tokens-param max_completion_tokens

エージェント モデルは gpt --provider openai --id gpt-5.2 --reasoning experience --max-output 16384 を追加します。
エージェント モデル テスト gpt # 1-token プローブ: 正しく構成されましたか?
エージェントの予算セット -- 毎日 25
編集はライブ セッションに直ちに適用されます。CLI は実行中のルーターをホットリロードします。
プロファイルには、メイン モデル、バックグラウンド タスク用の小規模/高速モデル、階層マッピング (プロファイル内で /model opus が解決されるように)、およびオプションのバジェットがバンドルされています。
プロフィール:
main : {model: ソネット、small_fast: haiku、tiers: {opus: opus、sonnet: ソネット、haiku: haiku}}
安い : {モデル: gpt、小型高速: gpt、予算: {日: 5.00}}
ローカル : {モデル: qwen、small_fast: qwen}
subscription : {passthrough: true} # プレーン クロード、サブスクリプション請求、追跡なし
動的ルーティング
手動でモデルを選択する代わりに、安価な LLM にすべてのタスクを優先順位付けさせます。
エージェント ルーティング セット auto --classifier haiku \
--ディープ オーパス --スタンダード ソネット --ライト クウェン
auto はモデル ( /model auto 、またはプロファイル: {model: auto} ) のように動作するようになりました。新しいユーザーのターンごとに、分類子はリクエストを読み取り、階層を割り当てます。計画とハード デバッグは詳細に、通常のコーディングは標準に、機械的な編集と検証は簡易に行われます。この決定はターンの残りの間維持される (ツールの結果によって再トリガーされることはない) ため、タスクが飛行中にモデルを反転することはありません。分類の失敗は --default (標準) にフォールバックし、すべての決定がログに記録されます。
$ grep autoroute ~/.agentic/router.log
... エイリアス = 自動層 = ディープ モデル = opus
... エイリアス=自動層=ライトモデル=qwen
エージェント費用 -- モデル別では、実際に費用が階層間でどのように配分されているかが示されます。各分類には、分類子モデルへの 1 つの小さなリクエストがかかります (俳句の場合は ~0.0005 ドル)。
auto のような routing: alias が有効になっているときは常に、2 番目の独立した分類子パスが新しいユーザーの順番を調べて、異なる質問をします。

n: 「どの層ですか?」ではありません。しかし、「これには単一の応答ではなく永続的なループが必要なようですか?」 — 長いビルドを監視し、条件が満たされるまで再試行し、デプロイをベビーシッターし、外部状態をポーリングします。
答えが「はい」の場合、エージェントはループ自体を開始しません (そして開始できません)。ルーターは要求と応答の本文のみを参照し、クロード コード プロセス内に実行コンテキストはありません。代わりに、ハーネス独自のメカニズムを直接指定するシステム リマインダーをリクエストに追加します。
<システムリマインダー>
エージェント的: このタスクは、目標ループではなく、繰り返しの目標ループに適しているように見えます。
単一の応答 (長いビルドのポーリング)。永続的なループが役立つとしたら —
進行状況を再度チェックし、条件が満たされるまで再試行し、
長時間実行プロセス - プロンプトを使用して ScheduleWakeup を呼び出す
「<<autonomous-loop-dynamic>>」と理由を指定するか、/loop スキルを呼び出します。これ
これは提案であり、要件ではありません。1 回で完了するタスクの場合は無視してください。
パスします。
</システムリマインダー>
クロード・コードは、それに基づいて行動するかどうかを決定します。リマインダーは命令ではなく、小言です。決定はログに記録され ( grep autogoal ~/.agentic/router.log )、層の決定と同様にステータスラインに表示されます ( ⟳ゴール (長いビルドのポーリング) )。
これは動的ルーティングと同じ分類子エイリアスに基づいているため、routing: auto が設定されている場合はどこでも実行されますが、層呼び出しと並行して新しいターンごとに 1 つの追加の安価な分類子呼び出しがかかります。
Claude Code は、対話していると思われる Claude モデルの ~200K ウィンドウに対して自動コンパクトのサイズを設定します。ルーティングされたモデルがこれに一致することはほとんどありません。ローカルの qwen は 32K を保持し、GPT は 400K を保持し、多くのモデルは、宣伝されている制限よりかなり前に信頼性が低くなります。モデルが実際に保持しているものを宣言すると、ルーターはレポートするすべてのトークン数をスケーリングするため、クライアントのコンテキスト ゲージ、つまり自動コンパックが行われます。

t — 実際のウィンドウを追跡します。
エージェント モデルは qwen --provider local --id qwen3-coder-30b --context-window 32768 を追加します
エージェント モデルは glm --provider z --id glm-4.7 --context-window 200000 --Effective-context 60000 を追加します。
効果的なコンテキストは注意ノブです。ウィンドウが名目上 200K であっても、クライアントは 60K の実際のトークンを圧縮し、モデルを一貫した範囲に保ちます。価格設定と予算には常に実際の使用状況が記録されます。エージェント コンテキストは、これらの数値を調整するためのセッションの実際の軌跡と報告された軌跡を示します。詳細と調査方法: docs/context-scaling.md 。
エージェント評価は、同じコーディング タスクでベースライン モデルとテスト対象のモデルを比較します。各アームは非インタラクティブなクロード コードを分離して実行し、独自のセッション ID でルーターの使用状況とルート決定を記録し、パッチを生成します。任意の裁判官はブラインドパッチと検証者の証拠を閲覧します。モデル名、コスト、実行順序は決して表示されません。
エグゼキュータには 2 つのタイプがあります。ローカル マニフェストは、リポジトリ、セットアップ コマンド、およびベリファイアを直接提供します。
バージョン : 1
名前 : ローカルサンプル
タスク:
- ID : ジャンゴ-11001
リポジトリ: /path/to/prepared/django
ベース：メイン
プロンプト : このタスクで説明されている問題を修正します。関連するテストを実行します。
検証者:
実行: [python、-m、pytest、tests/example_test.py]
タイムアウト：10分
SWE ベンチ マニフェストは、リポジトリのチェックアウト、依存関係のセットアップ、公式タスク イメージ、テスト パッチ適用、および FAIL_TO_PASS/PASS_TO_PASS グレーディングを固定された公式ハーネスに委任します。
バージョン : 1
名前：スウェベンチスモーク
データセット:
タイプ : スウェベンチ
出典：princeton-nlp/SWE-bench_Verified
分割 : テスト
タスク:
- astropy__astropy-14309
サンドボックス:
タイプ: ドッカー
同じマニフェストが example/swebench-smoke.yaml にあります。 SWE ベンチの実行には、サポートされている正確なパッケージ バージョンを備えた Docker および Python 3.10 以降が必要です。
Python3 -m ven

v ~ /.agentic/swebench-venv
~ /.agentic/swebench-venv/bin/pip インストール ' swebench==4.1.0 '
エージェント評価実行の例/swebench-smoke.yaml \
--python ~ /.agentic/swebench-venv/bin/python \
--baseline opus --mut auto --judge Sonnet \
--attempts 1 --timeout 45m --output ~ /.agentic/evals/swebench-smoke
エージェント評価レポート ~ /.agentic/evals/swebench-smoke
実際に 1 つのタスクを実行して、kimi-k3 と生成された opus を比較します。
swebench-smoke: ベースライン = opus mut = kimi-k3 ジャッジ = ソネット ペア = 1
勝利: ベースライン 1 · ミュート 0 · 同点 0 · ジャッジエラー 0 · インフラペア 0
検証成功: ベースライン 1 · mut 1 — 実行失敗: ベースライン 0 · mut 0 · インフラ失敗 0
タスク試行勝者のベースライン MUT
astropy__astropy-14309 1 ベースライン完了/パス完了/パス
どちらのパッチも公式の SWE ベンチグレーダーに合格しました。盲検ジャッジは、上流の修正と正確に一致し、幅が狭かったため、ベースラインを好みました。一方、kimi-k3 は、より広いがそれでも正しい防御修正を使用しました。これは 1 つの煙テスト データ ポイントであり、統計的に意味のあるモデルのランキングではありません。複数のタスクを使用して、実行する予定の比較を試みます。
アダプターは、モデル要求を行う前に、Python、正確な SWE ベンチ API、および Docker をチェックします。公式ハーネスに各インスタンス イメージを構築または再利用するよう依頼し、新しい候補を開始します

[切り捨てられた]

## Original Extract

Run Claude Code on any model, with a budget — multi-provider router, LLM-triaged tier routing, spend tracking - MaorBril/agentic

GitHub - MaorBril/agentic: Run Claude Code on any model, with a budget — multi-provider router, LLM-triaged tier routing, spend tracking · GitHub
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
MaorBril
/
agentic
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits .github/ workflows .github/ workflows cmd cmd docs docs examples examples internal internal .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh main.go main.go View all files Repository files navigation
Run Claude Code on any model, with a budget.
agentic wraps Claude Code in a thin local router. Your sessions look and feel exactly like claude — same TUI, same tools, same updates — but the model behind them can be Anthropic, OpenAI, xAI, or anything OpenAI-compatible (Ollama, vLLM, OpenRouter, DeepSeek, Groq). Every token is metered, priced, and checked against budgets you set.
agentic # Claude Code, tracked, on your default profile
agentic -p cheap # same session, cheaper models
agentic --model grok # one-off model override
agentic cost # where did today's $4.31 go?
Why
Claude Code is a great harness, and it keeps getting better — forking it means losing that. But it only talks to one provider, and it doesn't answer two questions you eventually ask: how much did that session cost? and can I run the cheap parts on a cheap model?
agentic answers both without touching Claude Code itself. Claude Code officially supports pointing at a gateway via ANTHROPIC_BASE_URL ; agentic is that gateway, plus the CLI around it.
There are three ways people solve "I want Claude Code but not locked to one provider":
The short version: agentic doesn't try to be a better harness than Claude Code — it keeps Claude Code exactly as Anthropic ships it and only swaps what's behind ANTHROPIC_BASE_URL . If you want a different agent loop altogether, OpenCode/Crush/Goose/Aider are the right layer to look at instead. If you want a gateway you deploy and administer for a team, LiteLLM is a more mature choice for that. agentic is for a single developer who wants claude , unmodified, with a budget and a cheap-model escape hatch, installed in one command and running with nothing to operate.
agentic (launcher) ──▶ claude (unmodified, auto-updating)
│ ANTHROPIC_BASE_URL
▼
local router (127.0.0.1)
├─ anthropic: byte-faithful passthrough
├─ openai dialect: full request/stream translation
├─ usage log (SQLite) + pricing
└─ budget gate
▼
Anthropic · OpenAI · xAI · Ollama · vLLM · OpenRouter · ...
There is no daemon. The first agentic session binds the router port and serves everyone; when it exits, another running session takes over within a couple of seconds. The last session out turns off the lights.
Model names are aliases you define. Claude Code treats model IDs as opaque strings, so ANTHROPIC_MODEL=grok flows straight through and the router resolves it. Anything starting with claude- passes through to Anthropic untouched — background tasks keep working even when your main model is something else entirely.
curl -fsSL https://raw.githubusercontent.com/maorbril/agentic/main/install.sh | sh
agentic setup
Or from source: go install github.com/maorbril/agentic@latest
To update later, run agentic update (or agentic update --check to see if one's available without installing it). This updates agentic itself; Claude Code keeps auto-updating on its own regardless.
Everything lives in ~/.agentic/config.yaml , and everything is editable from the terminal:
agentic providers add openai --type openai --base-url https://api.openai.com/v1 \
--key-env OPENAI_API_KEY --max-tokens-param max_completion_tokens
agentic models add gpt --provider openai --id gpt-5.2 --reasoning effort --max-output 16384
agentic models test gpt # 1-token probe: did I configure it right?
agentic budget set --daily 25
Edits apply to live sessions immediately — the CLI hot-reloads the running router.
A profile bundles a main model, a small/fast model for background tasks, tier mappings (so /model opus resolves inside the profile), and optional budgets:
profiles :
main : {model: sonnet, small_fast: haiku, tiers: {opus: opus, sonnet: sonnet, haiku: haiku}}
cheap : {model: gpt, small_fast: gpt, budget: {daily: 5.00}}
local : {model: qwen, small_fast: qwen}
subscription : {passthrough: true} # plain claude, subscription billing, no tracking
Dynamic routing
Instead of picking models by hand, let a cheap LLM triage every task:
agentic routing set auto --classifier haiku \
--deep opus --standard sonnet --light qwen
auto now behaves like a model ( /model auto , or profiles: {model: auto} ). On each new user turn, the classifier reads the request and assigns a tier — planning and hard debugging go deep , ordinary coding goes standard , mechanical edits and verification go light . The decision sticks for the rest of the turn (tool results don't re-trigger it), so a task never flips models mid-flight. Classification failures fall back to --default (standard), and every decision is logged:
$ grep autoroute ~/.agentic/router.log
... alias=auto tier=deep model=opus
... alias=auto tier=light model=qwen
agentic cost --by model then shows how spend actually distributed across tiers. Each classification costs one tiny request to the classifier model (~$0.0005 with haiku).
Whenever a routing: alias like auto is in play, a second, independent classifier pass looks at each new user turn and asks a different question: not "which tier?" but "does this look like it needs a persistent loop rather than a single reply?" — monitoring a long build, retrying until a condition holds, babysitting a deploy, polling for external state.
When the answer is yes, agentic doesn't (and can't) start a loop itself — the router only ever sees request and response bodies, it has no execution context inside the Claude Code process. Instead it appends a system-reminder to the request naming the harness's own mechanisms directly:
<system-reminder>
agentic: this task looks well suited to a recurring goal loop rather than a
single reply (polling a long build). If a persistent loop would help —
checking back on progress, retrying until a condition holds, babysitting a
long-running process — call ScheduleWakeup with prompt
"<<autonomous-loop-dynamic>>" and a reason, or invoke the /loop skill. This
is a suggestion, not a requirement: ignore it for tasks that finish in one
pass.
</system-reminder>
Claude Code decides whether to act on it — the reminder is a nudge, not a command. Decisions are logged ( grep autogoal ~/.agentic/router.log ) and, like tier decisions, surfaced in the statusline ( ⟳ goal (polling a long build) ).
This rides on the same classifier alias as dynamic routing, so it's on wherever routing: auto is configured, at the cost of one extra cheap classifier call per new turn alongside the tier call.
Claude Code sizes its auto-compact against the ~200K window of the Claude model it thinks it's talking to. Routed models rarely match that: a local qwen holds 32K, GPT holds 400K, and many models get unreliable well before their advertised limit. Declare what a model really holds and the router scales every token count it reports so the client's context gauge — and therefore auto-compact — tracks the real window:
agentic models add qwen --provider local --id qwen3-coder-30b --context-window 32768
agentic models add glm --provider z --id glm-4.7 --context-window 200000 --effective-context 60000
effective_context is the attention knob: the client compacts at 60K real tokens even though the window is nominally 200K, keeping the model in its coherent range. Pricing and budgets always record true usage; agentic context shows a session's true-vs-reported trajectory for tuning these numbers. Details and research methodology: docs/context-scaling.md .
agentic eval compares a baseline model with a model under test on the same coding tasks. Each arm runs non-interactive Claude Code in isolation, records router usage and route decisions under its own session ID, and produces a patch. An optional judge sees blinded patches and verifier evidence; it never sees model names, cost, or execution order.
There are two executor types. A local manifest supplies its repository, setup command, and verifier directly:
version : 1
name : local-sample
tasks :
- id : django-11001
repo : /path/to/prepared/django
base : main
prompt : Fix the issue described in this task. Run relevant tests.
verifier :
run : [python, -m, pytest, tests/example_test.py]
timeout : 10m
A SWE-bench manifest delegates repository checkout, dependency setup, official task images, test-patch application, and FAIL_TO_PASS/PASS_TO_PASS grading to the pinned official harness:
version : 1
name : swebench-smoke
dataset :
type : swebench
source : princeton-nlp/SWE-bench_Verified
split : test
tasks :
- astropy__astropy-14309
sandbox :
type : docker
The same manifest is in examples/swebench-smoke.yaml . SWE-bench runs require Docker and Python 3.10 or newer with the exact supported package version:
python3 -m venv ~ /.agentic/swebench-venv
~ /.agentic/swebench-venv/bin/pip install ' swebench==4.1.0 '
agentic eval run examples/swebench-smoke.yaml \
--python ~ /.agentic/swebench-venv/bin/python \
--baseline opus --mut auto --judge sonnet \
--attempts 1 --timeout 45m --output ~ /.agentic/evals/swebench-smoke
agentic eval report ~ /.agentic/evals/swebench-smoke
A real one-task run comparing kimi-k3 against opus produced:
swebench-smoke: baseline=opus mut=kimi-k3 judge=sonnet pairs=1
wins: baseline 1 · mut 0 · ties 0 · judge errors 0 · infra pairs 0
verifier passes: baseline 1 · mut 1 — run failures: baseline 0 · mut 0 · infra failures 0
TASK ATTEMPT WINNER BASELINE MUT
astropy__astropy-14309 1 baseline complete/pass complete/pass
Both patches passed the official SWE-bench grader. The blinded judge preferred the baseline because it exactly matched the upstream fix and was narrower, while kimi-k3 used a broader but still correct defensive fix. This is one smoke-test data point, not a statistically meaningful model ranking; use multiple tasks and attempts for comparisons you intend to act on.
The adapter checks Python, the exact SWE-bench API, and Docker before making a model request. It asks the official harness to build or reuse each instance image, starts a fresh candidate

[truncated]
